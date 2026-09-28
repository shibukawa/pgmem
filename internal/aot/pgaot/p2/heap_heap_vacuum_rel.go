package p2

import (
	base "github.com/shibukawa/pgmem/internal/aot/pgaot/base"
	"sync/atomic"
	"unsafe"
)

func F_heap_vacuum_rel(m *base.Module, l0 int32, l1 int32, l2 int32) {
	mBase := m.M
	_ = mBase
	var v4 int32
	_ = v4
	var v40 int64
	_ = v40
	var v51 int32
	_ = v51
	var v53 int32
	_ = v53
	var v56 int64
	_ = v56
	var v59 int64
	_ = v59
	var v62 int64
	_ = v62
	var v65 int64
	_ = v65
	var v68 int64
	_ = v68
	var v75 int32
	_ = v75
	var v77 int32
	_ = v77
	var v82 int32
	_ = v82
	var v86 int32
	_ = v86
	var v95 int32
	_ = v95
	var v98 int32
	_ = v98
	var v102 int64
	_ = v102
	var v104 int64
	_ = v104
	var v105 int32
	_ = v105
	var v106 int64
	_ = v106
	var v107 int64
	_ = v107
	var v109 int32
	_ = v109
	var v113 int32
	_ = v113
	var v114 int32
	_ = v114
	var v115 int32
	_ = v115
	var v118 int64
	_ = v118
	var v119 int64
	_ = v119
	var v127 int64
	_ = v127
	var v129 int32
	_ = v129
	var v132 int32
	_ = v132
	var v136 int32
	_ = v136
	var v141 int32
	_ = v141
	var v143 int32
	_ = v143
	var v144 int32
	_ = v144
	var v147 int32
	_ = v147
	var v151 int32
	_ = v151
	var v153 int32
	_ = v153
	var v154 int32
	_ = v154
	var v165 int32
	_ = v165
	var v166 int32
	_ = v166
	var v172 int32
	_ = v172
	var v178 int32
	_ = v178
	var v183 int32
	_ = v183
	var v184 int64
	_ = v184
	var v186 int64
	_ = v186
	var v189 int32
	_ = v189
	var v193 int32
	_ = v193
	var v198 int32
	_ = v198
	var v200 int32
	_ = v200
	var v201 int32
	_ = v201
	var v204 int32
	_ = v204
	var v208 int32
	_ = v208
	var v210 int32
	_ = v210
	var v211 int32
	_ = v211
	var v219 int32
	_ = v219
	var v220 int32
	_ = v220
	var v226 int32
	_ = v226
	var v231 int32
	_ = v231
	var v232 int32
	_ = v232
	var v234 int32
	_ = v234
	var v235 int32
	_ = v235
	var v236 int32
	_ = v236
	var v238 int32
	_ = v238
	var v239 int32
	_ = v239
	var v240 int32
	_ = v240
	var v241 int32
	_ = v241
	var v243 int32
	_ = v243
	var v246 int32
	_ = v246
	var v247 int32
	_ = v247
	var v249 int32
	_ = v249
	var v257 int32
	_ = v257
	var v258 int32
	_ = v258
	var v267 int32
	_ = v267
	var v269 int32
	_ = v269
	var v271 int32
	_ = v271
	var v275 int32
	_ = v275
	var v279 int32
	_ = v279
	var v280 int32
	_ = v280
	var v281 int32
	_ = v281
	var v287 int32
	_ = v287
	var v336 int32
	_ = v336
	var v338 int32
	_ = v338
	var v340 int32
	_ = v340
	var v341 int32
	_ = v341
	var v344 int32
	_ = v344
	var v345 int32
	_ = v345
	var v348 int32
	_ = v348
	var v349 int32
	_ = v349
	var v378 int32
	_ = v378
	var v402 int32
	_ = v402
	var v404 int32
	_ = v404
	var v406 int32
	_ = v406
	var v408 int32
	_ = v408
	var v409 int32
	_ = v409
	var v412 int32
	_ = v412
	var v415 int32
	_ = v415
	var v417 int32
	_ = v417
	var v419 int64
	_ = v419
	var v431 int32
	_ = v431
	var v434 int32
	_ = v434
	var v435 int32
	_ = v435
	var v436 int32
	_ = v436
	var v442 int32
	_ = v442
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
	var v455 int32
	_ = v455
	var v458 int32
	_ = v458
	var v459 int32
	_ = v459
	var v462 int64
	_ = v462
	var v464 int32
	_ = v464
	var v466 int32
	_ = v466
	var v467 int32
	_ = v467
	var v476 float64
	_ = v476
	var v479 int32
	_ = v479
	var v480 int32
	_ = v480
	var v483 int32
	_ = v483
	var v486 int32
	_ = v486
	var v493 int32
	_ = v493
	var v496 int32
	_ = v496
	var v502 int32
	_ = v502
	var v508 int32
	_ = v508
	var v509 int32
	_ = v509
	var v510 int32
	_ = v510
	var v515 int32
	_ = v515
	var v520 int32
	_ = v520
	var v522 int32
	_ = v522
	var v524 float64
	_ = v524
	var v527 int32
	_ = v527
	var v543 int32
	_ = v543
	var v544 int64
	_ = v544
	var v547 int32
	_ = v547
	var v551 int32
	_ = v551
	var v556 int32
	_ = v556
	var v558 int32
	_ = v558
	var v559 int32
	_ = v559
	var v562 int32
	_ = v562
	var v566 int32
	_ = v566
	var v568 int32
	_ = v568
	var v569 int32
	_ = v569
	var v577 int32
	_ = v577
	var v578 int32
	_ = v578
	var v584 int32
	_ = v584
	var v590 int32
	_ = v590
	var v593 int32
	_ = v593
	var v594 int32
	_ = v594
	var v597 int64
	_ = v597
	var v598 int32
	_ = v598
	var v604 int32
	_ = v604
	var v605 int32
	_ = v605
	var v609 int32
	_ = v609
	var v613 int32
	_ = v613
	var v616 int32
	_ = v616
	var v619 int32
	_ = v619
	var v620 int32
	_ = v620
	var v622 int32
	_ = v622
	var v624 int32
	_ = v624
	var v627 int32
	_ = v627
	var v629 int32
	_ = v629
	var v632 int32
	_ = v632
	var v633 int32
	_ = v633
	var v636 int32
	_ = v636
	var v639 int32
	_ = v639
	var v642 int32
	_ = v642
	var v643 int32
	_ = v643
	var v644 int32
	_ = v644
	var v651 int32
	_ = v651
	var v652 int32
	_ = v652
	var v655 int32
	_ = v655
	var v661 int32
	_ = v661
	var v666 int32
	_ = v666
	var v667 int32
	_ = v667
	var v670 int32
	_ = v670
	var v671 int32
	_ = v671
	var v672 int32
	_ = v672
	var v674 int32
	_ = v674
	var v675 int32
	_ = v675
	var v679 int32
	_ = v679
	var v687 int32
	_ = v687
	var v690 int32
	_ = v690
	var v691 int32
	_ = v691
	var v700 int32
	_ = v700
	var v701 int32
	_ = v701
	var v702 int32
	_ = v702
	var v747 int32
	_ = v747
	var v748 int32
	_ = v748
	var v749 int32
	_ = v749
	var v753 int32
	_ = v753
	var v754 int32
	_ = v754
	var v756 int32
	_ = v756
	var v759 int32
	_ = v759
	var v769 int32
	_ = v769
	var v770 int32
	_ = v770
	var v772 int32
	_ = v772
	var v775 int32
	_ = v775
	var v777 int32
	_ = v777
	var v781 int32
	_ = v781
	var v784 int32
	_ = v784
	var v786 int32
	_ = v786
	var v840 int32
	_ = v840
	var v843 int32
	_ = v843
	var v844 int32
	_ = v844
	var v852 int32
	_ = v852
	var v853 int32
	_ = v853
	var v859 int32
	_ = v859
	var v860 int32
	_ = v860
	var v863 int32
	_ = v863
	var v864 int32
	_ = v864
	var v865 int32
	_ = v865
	var v870 int32
	_ = v870
	var v871 int32
	_ = v871
	var v873 int32
	_ = v873
	var v875 int32
	_ = v875
	var v876 int32
	_ = v876
	var v878 int32
	_ = v878
	var v880 int32
	_ = v880
	var v881 int32
	_ = v881
	var v883 int32
	_ = v883
	var v885 int32
	_ = v885
	var v886 int32
	_ = v886
	var v888 int32
	_ = v888
	var v890 int32
	_ = v890
	var v891 int32
	_ = v891
	var v892 int32
	_ = v892
	var v897 int32
	_ = v897
	var v898 int32
	_ = v898
	var v900 int32
	_ = v900
	var v902 int32
	_ = v902
	var v903 int32
	_ = v903
	var v905 int32
	_ = v905
	var v907 int32
	_ = v907
	var v908 int32
	_ = v908
	var v909 int32
	_ = v909
	var v914 int32
	_ = v914
	var v915 int32
	_ = v915
	var v917 int32
	_ = v917
	var v919 int32
	_ = v919
	var v920 int32
	_ = v920
	var v923 int32
	_ = v923
	var v924 int32
	_ = v924
	var v925 int32
	_ = v925
	var v930 int32
	_ = v930
	var v931 int32
	_ = v931
	var v933 int32
	_ = v933
	var v935 int32
	_ = v935
	var v936 int32
	_ = v936
	var v938 int32
	_ = v938
	var v940 int32
	_ = v940
	var v943 int32
	_ = v943
	var v944 int32
	_ = v944
	var v945 int32
	_ = v945
	var v958 int32
	_ = v958
	var v960 int32
	_ = v960
	var v962 int32
	_ = v962
	var v967 int32
	_ = v967
	var v978 int32
	_ = v978
	var v981 int32
	_ = v981
	var v982 int32
	_ = v982
	var v989 int32
	_ = v989
	var v991 int32
	_ = v991
	var v1035 int32
	_ = v1035
	var v1041 int32
	_ = v1041
	var v1042 int32
	_ = v1042
	var v1043 int32
	_ = v1043
	var v1044 int32
	_ = v1044
	var v1047 int32
	_ = v1047
	var v1053 int32
	_ = v1053
	var v1057 int32
	_ = v1057
	var v1062 int32
	_ = v1062
	var v1066 int32
	_ = v1066
	var v1070 int32
	_ = v1070
	var v1072 int32
	_ = v1072
	var v1075 int32
	_ = v1075
	var v1077 int32
	_ = v1077
	var v1079 int32
	_ = v1079
	var v1080 int32
	_ = v1080
	var v1081 int32
	_ = v1081
	var v1084 int32
	_ = v1084
	var v1089 int32
	_ = v1089
	var v1093 int64
	_ = v1093
	var v1094 int64
	_ = v1094
	var v1099 int32
	_ = v1099
	var v1100 int32
	_ = v1100
	var v1101 int32
	_ = v1101
	var v1104 int32
	_ = v1104
	var v1106 int32
	_ = v1106
	var v1107 int32
	_ = v1107
	var v1109 int32
	_ = v1109
	var v1110 int32
	_ = v1110
	var v1111 int32
	_ = v1111
	var v1113 int32
	_ = v1113
	var v1114 int32
	_ = v1114
	var v1115 int32
	_ = v1115
	var v1120 int32
	_ = v1120
	var v1121 int32
	_ = v1121
	var v1122 int32
	_ = v1122
	var v1130 int32
	_ = v1130
	var v1132 int32
	_ = v1132
	var v1137 float64
	_ = v1137
	var v1140 int32
	_ = v1140
	var v1143 int32
	_ = v1143
	var v1146 int32
	_ = v1146
	var v1149 int32
	_ = v1149
	var v1153 int32
	_ = v1153
	var v1160 int32
	_ = v1160
	var v1164 int32
	_ = v1164
	var v1166 int32
	_ = v1166
	var v1169 int32
	_ = v1169
	var v1171 int32
	_ = v1171
	var v1173 int32
	_ = v1173
	var v1174 int32
	_ = v1174
	var v1175 int32
	_ = v1175
	var v1176 int32
	_ = v1176
	var v1177 int32
	_ = v1177
	var v1178 int32
	_ = v1178
	var v1181 int32
	_ = v1181
	var v1183 int32
	_ = v1183
	var v1185 int32
	_ = v1185
	var v1186 int32
	_ = v1186
	var v1187 int32
	_ = v1187
	var v1188 int32
	_ = v1188
	var v1189 int32
	_ = v1189
	var v1190 int32
	_ = v1190
	var v1193 int32
	_ = v1193
	var v1196 int32
	_ = v1196
	var v1197 int32
	_ = v1197
	var v1199 int32
	_ = v1199
	var v1200 int32
	_ = v1200
	var v1201 int32
	_ = v1201
	var v1203 int32
	_ = v1203
	var v1206 int32
	_ = v1206
	var v1208 int32
	_ = v1208
	var v1211 int32
	_ = v1211
	var v1264 int32
	_ = v1264
	var v1316 int32
	_ = v1316
	var v1321 int32
	_ = v1321
	var v1325 int32
	_ = v1325
	var v1377 int32
	_ = v1377
	var v1378 int32
	_ = v1378
	var v1385 int32
	_ = v1385
	var v1386 int32
	_ = v1386
	var v1387 int32
	_ = v1387
	var v1438 int32
	_ = v1438
	var v1440 int32
	_ = v1440
	var v1442 int32
	_ = v1442
	var v1444 int32
	_ = v1444
	var v1448 int32
	_ = v1448
	var v1450 int32
	_ = v1450
	var v1451 int32
	_ = v1451
	var v1452 int32
	_ = v1452
	var v1455 int32
	_ = v1455
	var v1458 int64
	_ = v1458
	var v1460 int64
	_ = v1460
	var v1464 int32
	_ = v1464
	var v1465 int64
	_ = v1465
	var v1481 int32
	_ = v1481
	var v1485 int32
	_ = v1485
	var v1490 int32
	_ = v1490
	var v1492 int32
	_ = v1492
	var v1493 int32
	_ = v1493
	var v1496 int32
	_ = v1496
	var v1500 int32
	_ = v1500
	var v1503 int32
	_ = v1503
	var v1595 int32
	_ = v1595
	var v1598 int32
	_ = v1598
	var v1607 int32
	_ = v1607
	var v1608 int32
	_ = v1608
	var v1614 int64
	_ = v1614
	var v1616 int32
	_ = v1616
	var v1619 int32
	_ = v1619
	var v1630 int32
	_ = v1630
	var v1633 int32
	_ = v1633
	var v1634 int32
	_ = v1634
	var v1635 int32
	_ = v1635
	var v1638 int32
	_ = v1638
	var v1640 int32
	_ = v1640
	var v1653 int32
	_ = v1653
	var v1660 int32
	_ = v1660
	var v1662 int32
	_ = v1662
	var v1663 int32
	_ = v1663
	var v1664 int32
	_ = v1664
	var v1665 int32
	_ = v1665
	var v1669 int32
	_ = v1669
	var v1670 int32
	_ = v1670
	var v1677 int32
	_ = v1677
	var v1681 int32
	_ = v1681
	var v1703 int32
	_ = v1703
	var v1722 int32
	_ = v1722
	var v1726 int32
	_ = v1726
	var v1727 int32
	_ = v1727
	var v1728 int32
	_ = v1728
	var v1735 int32
	_ = v1735
	var v1736 int32
	_ = v1736
	var v1737 int32
	_ = v1737
	var v1738 int64
	_ = v1738
	var v1741 int32
	_ = v1741
	var v1742 int32
	_ = v1742
	var v1743 int32
	_ = v1743
	var v1744 int32
	_ = v1744
	var v1745 int32
	_ = v1745
	var v1747 int32
	_ = v1747
	var v1749 int32
	_ = v1749
	var v1752 int32
	_ = v1752
	var v1755 int32
	_ = v1755
	var v1756 int32
	_ = v1756
	var v1760 int32
	_ = v1760
	var v1765 int32
	_ = v1765
	var v1769 int32
	_ = v1769
	var v1774 int32
	_ = v1774
	var v1776 int32
	_ = v1776
	var v1777 int32
	_ = v1777
	var v1780 int32
	_ = v1780
	var v1784 int32
	_ = v1784
	var v1786 int32
	_ = v1786
	var v1787 int32
	_ = v1787
	var v1795 int32
	_ = v1795
	var v1796 int32
	_ = v1796
	var v1802 int32
	_ = v1802
	var v1807 int32
	_ = v1807
	var v1809 int32
	_ = v1809
	var v1814 int32
	_ = v1814
	var v1815 int32
	_ = v1815
	var v1823 int32
	_ = v1823
	var v1868 int32
	_ = v1868
	var v1875 int32
	_ = v1875
	var v1876 int32
	_ = v1876
	var v2010 int32
	_ = v2010
	var v2031 int32
	_ = v2031
	var v2032 int32
	_ = v2032
	var v2033 int32
	_ = v2033
	var v2034 int32
	_ = v2034
	var v2036 int32
	_ = v2036
	var v2040 int32
	_ = v2040
	var v2046 int32
	_ = v2046
	var v2048 int32
	_ = v2048
	var v2054 int32
	_ = v2054
	var v2058 int32
	_ = v2058
	var v2064 int32
	_ = v2064
	var v2066 int32
	_ = v2066
	var v2067 int32
	_ = v2067
	var v2072 int32
	_ = v2072
	var v2073 int32
	_ = v2073
	var v2074 int32
	_ = v2074
	var v2075 int32
	_ = v2075
	var v2080 int32
	_ = v2080
	var v2088 int32
	_ = v2088
	var v2092 int32
	_ = v2092
	var v2097 int32
	_ = v2097
	var v2099 int32
	_ = v2099
	var v2100 int32
	_ = v2100
	var v2103 int32
	_ = v2103
	var v2107 int32
	_ = v2107
	var v2109 int32
	_ = v2109
	var v2110 int32
	_ = v2110
	var v2118 int32
	_ = v2118
	var v2119 int32
	_ = v2119
	var v2125 int32
	_ = v2125
	var v2131 int32
	_ = v2131
	var v2134 int32
	_ = v2134
	var v2138 int32
	_ = v2138
	var v2139 int32
	_ = v2139
	var v2140 int32
	_ = v2140
	var v2145 int32
	_ = v2145
	var v2146 int32
	_ = v2146
	var v2150 int32
	_ = v2150
	var v2151 int32
	_ = v2151
	var v2152 int32
	_ = v2152
	var v2153 int32
	_ = v2153
	var v2155 int32
	_ = v2155
	var v2156 int32
	_ = v2156
	var v2162 int32
	_ = v2162
	var v2165 int32
	_ = v2165
	var v2166 int32
	_ = v2166
	var v2170 int32
	_ = v2170
	var v2177 int32
	_ = v2177
	var v2178 int32
	_ = v2178
	var v2180 int32
	_ = v2180
	var v2185 int32
	_ = v2185
	var v2188 int32
	_ = v2188
	var v2190 int32
	_ = v2190
	var v2192 int32
	_ = v2192
	var v2193 int32
	_ = v2193
	var v2195 int64
	_ = v2195
	var v2198 int32
	_ = v2198
	var v2199 int32
	_ = v2199
	var v2200 int32
	_ = v2200
	var v2201 int32
	_ = v2201
	var v2202 int32
	_ = v2202
	var v2206 int32
	_ = v2206
	var v2209 int32
	_ = v2209
	var v2210 int32
	_ = v2210
	var v2212 int32
	_ = v2212
	var v2224 int32
	_ = v2224
	var v2225 int32
	_ = v2225
	var v2227 int32
	_ = v2227
	var v2232 int32
	_ = v2232
	var v2233 int32
	_ = v2233
	var v2234 int32
	_ = v2234
	var v2237 int32
	_ = v2237
	var v2245 int32
	_ = v2245
	var v2246 int32
	_ = v2246
	var v2247 int32
	_ = v2247
	var v2248 int32
	_ = v2248
	var v2251 int32
	_ = v2251
	var v2253 int32
	_ = v2253
	var v2264 int32
	_ = v2264
	var v2269 int32
	_ = v2269
	var v2278 int32
	_ = v2278
	var v2287 int32
	_ = v2287
	var v2293 int32
	_ = v2293
	var v2294 int32
	_ = v2294
	var v2308 int32
	_ = v2308
	var v2310 int32
	_ = v2310
	var v2311 int32
	_ = v2311
	var v2314 int32
	_ = v2314
	var v2316 int32
	_ = v2316
	var v2317 int32
	_ = v2317
	var v2318 int32
	_ = v2318
	var v2320 int32
	_ = v2320
	var v2327 int32
	_ = v2327
	var v2329 int32
	_ = v2329
	var v2332 int32
	_ = v2332
	var v2340 int32
	_ = v2340
	var v2344 int32
	_ = v2344
	var v2348 int32
	_ = v2348
	var v2349 int32
	_ = v2349
	var v2355 int32
	_ = v2355
	var v2357 int32
	_ = v2357
	var v2393 int32
	_ = v2393
	var v2394 int32
	_ = v2394
	var v2403 int32
	_ = v2403
	var v2416 int32
	_ = v2416
	var v2417 int32
	_ = v2417
	var v2418 int32
	_ = v2418
	var v2422 int32
	_ = v2422
	var v2430 int32
	_ = v2430
	var v2431 int32
	_ = v2431
	var v2433 int32
	_ = v2433
	var v2436 int32
	_ = v2436
	var v2437 int32
	_ = v2437
	var v2438 int32
	_ = v2438
	var v2448 int32
	_ = v2448
	var v2452 int32
	_ = v2452
	var v2457 int32
	_ = v2457
	var v2459 int32
	_ = v2459
	var v2460 int32
	_ = v2460
	var v2461 int32
	_ = v2461
	var v2462 int32
	_ = v2462
	var v2463 int32
	_ = v2463
	var v2465 int32
	_ = v2465
	var v2469 int32
	_ = v2469
	var v2470 int32
	_ = v2470
	var v2471 int32
	_ = v2471
	var v2475 int64
	_ = v2475
	var v2476 int64
	_ = v2476
	var v2477 int32
	_ = v2477
	var v2478 int32
	_ = v2478
	var v2481 int64
	_ = v2481
	var v2486 int32
	_ = v2486
	var v2489 int32
	_ = v2489
	var v2497 int32
	_ = v2497
	var v2498 int32
	_ = v2498
	var v2526 int64
	_ = v2526
	var v2530 int64
	_ = v2530
	var v2537 int32
	_ = v2537
	var v2538 int32
	_ = v2538
	var v2541 int32
	_ = v2541
	var v2546 int32
	_ = v2546
	var v2547 int32
	_ = v2547
	var v2553 int32
	_ = v2553
	var v2557 int32
	_ = v2557
	var v2558 int32
	_ = v2558
	var v2559 int64
	_ = v2559
	var v2560 int64
	_ = v2560
	var v2563 int32
	_ = v2563
	var v2564 int64
	_ = v2564
	var v2566 int32
	_ = v2566
	var v2567 int32
	_ = v2567
	var v2568 int32
	_ = v2568
	var v2585 int32
	_ = v2585
	var v2589 int32
	_ = v2589
	var v2594 int32
	_ = v2594
	var v2596 int32
	_ = v2596
	var v2597 int32
	_ = v2597
	var v2600 int32
	_ = v2600
	var v2604 int32
	_ = v2604
	var v2607 int32
	_ = v2607
	var v2699 int32
	_ = v2699
	var v2702 int32
	_ = v2702
	var v2711 int32
	_ = v2711
	var v2712 int32
	_ = v2712
	var v2718 int64
	_ = v2718
	var v2720 int32
	_ = v2720
	var v2723 int32
	_ = v2723
	var v2734 int32
	_ = v2734
	var v2737 int32
	_ = v2737
	var v2738 int32
	_ = v2738
	var v2739 int32
	_ = v2739
	var v2742 int32
	_ = v2742
	var v2744 int32
	_ = v2744
	var v2757 int64
	_ = v2757
	var v2762 int32
	_ = v2762
	var v2766 int32
	_ = v2766
	var v2770 int32
	_ = v2770
	var v2799 int64
	_ = v2799
	var v2803 int64
	_ = v2803
	var v2810 int64
	_ = v2810
	var v2813 int64
	_ = v2813
	var v2816 int64
	_ = v2816
	var v2822 int32
	_ = v2822
	var v2831 int32
	_ = v2831
	var v2833 int32
	_ = v2833
	var v2836 int32
	_ = v2836
	var v2838 int32
	_ = v2838
	var v2839 int32
	_ = v2839
	var v2846 int32
	_ = v2846
	var v2890 int32
	_ = v2890
	var v2896 int32
	_ = v2896
	var v2900 int32
	_ = v2900
	var v2906 int32
	_ = v2906
	var v2907 int32
	_ = v2907
	var v2916 int32
	_ = v2916
	var v2917 int32
	_ = v2917
	var v2920 int32
	_ = v2920
	var v2924 int32
	_ = v2924
	var v2927 int32
	_ = v2927
	var v2934 int32
	_ = v2934
	var v2935 int32
	_ = v2935
	var v2938 int32
	_ = v2938
	var v2940 int32
	_ = v2940
	var v2941 int32
	_ = v2941
	var v2942 int64
	_ = v2942
	var v2946 int32
	_ = v2946
	var v2947 int64
	_ = v2947
	var v2949 int32
	_ = v2949
	var v2950 int32
	_ = v2950
	var v2951 int32
	_ = v2951
	var v2968 int32
	_ = v2968
	var v2972 int32
	_ = v2972
	var v2977 int32
	_ = v2977
	var v2979 int32
	_ = v2979
	var v2980 int32
	_ = v2980
	var v2983 int32
	_ = v2983
	var v2987 int32
	_ = v2987
	var v2990 int32
	_ = v2990
	var v3082 int32
	_ = v3082
	var v3085 int32
	_ = v3085
	var v3094 int32
	_ = v3094
	var v3095 int32
	_ = v3095
	var v3101 int64
	_ = v3101
	var v3103 int32
	_ = v3103
	var v3106 int32
	_ = v3106
	var v3117 int32
	_ = v3117
	var v3120 int32
	_ = v3120
	var v3121 int32
	_ = v3121
	var v3122 int32
	_ = v3122
	var v3125 int32
	_ = v3125
	var v3127 int32
	_ = v3127
	var v3142 int32
	_ = v3142
	var v3145 int32
	_ = v3145
	var v3149 int32
	_ = v3149
	var v3152 int32
	_ = v3152
	var v3156 int32
	_ = v3156
	var v3159 int32
	_ = v3159
	var v3163 int64
	_ = v3163
	var v3164 int32
	_ = v3164
	var v3168 int64
	_ = v3168
	var v3169 int64
	_ = v3169
	var v3172 int64
	_ = v3172
	var v3173 int32
	_ = v3173
	var v3177 int64
	_ = v3177
	var v3178 int64
	_ = v3178
	var v3181 int64
	_ = v3181
	var v3182 int64
	_ = v3182
	var v3185 int32
	_ = v3185
	var v3191 int32
	_ = v3191
	var v3192 int32
	_ = v3192
	var v3194 int32
	_ = v3194
	var v3195 int32
	_ = v3195
	var v3201 int32
	_ = v3201
	var v3203 int32
	_ = v3203
	var v3206 int32
	_ = v3206
	var v3211 int32
	_ = v3211
	var v3212 int32
	_ = v3212
	var v3214 int32
	_ = v3214
	var v3215 int32
	_ = v3215
	var v3218 int64
	_ = v3218
	var v3219 int32
	_ = v3219
	var v3227 int32
	_ = v3227
	var v3232 int32
	_ = v3232
	var v3238 int32
	_ = v3238
	var v3246 int32
	_ = v3246
	var v3250 int32
	_ = v3250
	var v3254 int32
	_ = v3254
	var v3294 int32
	_ = v3294
	var v3295 int32
	_ = v3295
	var v3302 int32
	_ = v3302
	var v3303 int32
	_ = v3303
	var v3304 int32
	_ = v3304
	var v3305 int32
	_ = v3305
	var v3308 int32
	_ = v3308
	var v3310 int32
	_ = v3310
	var v3321 int32
	_ = v3321
	var v3326 int32
	_ = v3326
	var v3335 int32
	_ = v3335
	var v3344 int32
	_ = v3344
	var v3350 int32
	_ = v3350
	var v3351 int32
	_ = v3351
	var v3365 int32
	_ = v3365
	var v3367 int32
	_ = v3367
	var v3368 int32
	_ = v3368
	var v3370 int32
	_ = v3370
	var v3373 int32
	_ = v3373
	var v3374 int32
	_ = v3374
	var v3383 int32
	_ = v3383
	var v3385 int32
	_ = v3385
	var v3387 int32
	_ = v3387
	var v3390 int32
	_ = v3390
	var v3392 int32
	_ = v3392
	var v3396 int32
	_ = v3396
	var v3400 int32
	_ = v3400
	var v3405 int32
	_ = v3405
	var v3407 int32
	_ = v3407
	var v3408 int32
	_ = v3408
	var v3411 int32
	_ = v3411
	var v3415 int32
	_ = v3415
	var v3417 int32
	_ = v3417
	var v3418 int32
	_ = v3418
	var v3426 int32
	_ = v3426
	var v3427 int32
	_ = v3427
	var v3433 int32
	_ = v3433
	var v3437 int32
	_ = v3437
	var v3438 int32
	_ = v3438
	var v3439 int64
	_ = v3439
	var v3440 float64
	_ = v3440
	var v3445 int32
	_ = v3445
	var v3446 float32
	_ = v3446
	var v3447 float64
	_ = v3447
	var v3448 int32
	_ = v3448
	var v3461 float64
	_ = v3461
	var v3465 int32
	_ = v3465
	var v3485 float64
	_ = v3485
	var v3490 float64
	_ = v3490
	var v3492 float64
	_ = v3492
	var v3495 float64
	_ = v3495
	var v3496 int64
	_ = v3496
	var v3499 int64
	_ = v3499
	var v3504 int32
	_ = v3504
	var v3505 int32
	_ = v3505
	var v3506 int64
	_ = v3506
	var v3510 int32
	_ = v3510
	var v3512 int32
	_ = v3512
	var v3514 int32
	_ = v3514
	var v3518 int32
	_ = v3518
	var v3522 int32
	_ = v3522
	var v3527 int32
	_ = v3527
	var v3529 int32
	_ = v3529
	var v3530 int32
	_ = v3530
	var v3533 int32
	_ = v3533
	var v3537 int32
	_ = v3537
	var v3539 int32
	_ = v3539
	var v3540 int32
	_ = v3540
	var v3548 int32
	_ = v3548
	var v3549 int32
	_ = v3549
	var v3555 int32
	_ = v3555
	var v3559 int32
	_ = v3559
	var v3562 int32
	_ = v3562
	var v3565 int32
	_ = v3565
	var v3566 int32
	_ = v3566
	var v3567 float64
	_ = v3567
	var v3576 int64
	_ = v3576
	var v3594 int32
	_ = v3594
	var v3598 int32
	_ = v3598
	var v3603 int32
	_ = v3603
	var v3605 int32
	_ = v3605
	var v3606 int32
	_ = v3606
	var v3609 int32
	_ = v3609
	var v3613 int32
	_ = v3613
	var v3616 int32
	_ = v3616
	var v3708 int32
	_ = v3708
	var v3711 int32
	_ = v3711
	var v3720 int32
	_ = v3720
	var v3721 int32
	_ = v3721
	var v3727 int64
	_ = v3727
	var v3729 int32
	_ = v3729
	var v3732 int32
	_ = v3732
	var v3743 int32
	_ = v3743
	var v3746 int32
	_ = v3746
	var v3747 int32
	_ = v3747
	var v3748 int32
	_ = v3748
	var v3751 int32
	_ = v3751
	var v3753 int32
	_ = v3753
	var v3766 int32
	_ = v3766
	var v3769 int32
	_ = v3769
	var v3813 int64
	_ = v3813
	var v3826 int32
	_ = v3826
	var v3827 int32
	_ = v3827
	var v3829 int32
	_ = v3829
	var v3830 int32
	_ = v3830
	var v3832 int32
	_ = v3832
	var v3834 int32
	_ = v3834
	var v3839 int32
	_ = v3839
	var v3842 int32
	_ = v3842
	var v3844 int32
	_ = v3844
	var v3847 int32
	_ = v3847
	var v3848 int32
	_ = v3848
	var v3850 int32
	_ = v3850
	var v3851 int32
	_ = v3851
	var v3853 int32
	_ = v3853
	var v3856 int32
	_ = v3856
	var v3861 int32
	_ = v3861
	var v3862 int32
	_ = v3862
	var v3866 int32
	_ = v3866
	var v3868 int32
	_ = v3868
	var v3869 int32
	_ = v3869
	var v3871 int32
	_ = v3871
	var v3876 int64
	_ = v3876
	var v3879 int32
	_ = v3879
	var v3883 int32
	_ = v3883
	var v3888 int32
	_ = v3888
	var v3890 int32
	_ = v3890
	var v3891 int32
	_ = v3891
	var v3894 int32
	_ = v3894
	var v3898 int32
	_ = v3898
	var v3900 int32
	_ = v3900
	var v3901 int32
	_ = v3901
	var v3909 int32
	_ = v3909
	var v3910 int32
	_ = v3910
	var v3916 int32
	_ = v3916
	var v3920 int64
	_ = v3920
	var v3922 int32
	_ = v3922
	var v3923 int32
	_ = v3923
	var v3927 int32
	_ = v3927
	var v3932 int32
	_ = v3932
	var v3997 int32
	_ = v3997
	var v4001 int32
	_ = v4001
	var v4006 int32
	_ = v4006
	var v4008 int32
	_ = v4008
	var v4009 int32
	_ = v4009
	var v4012 int32
	_ = v4012
	var v4016 int32
	_ = v4016
	var v4019 int32
	_ = v4019
	var v4111 int32
	_ = v4111
	var v4114 int32
	_ = v4114
	var v4123 int32
	_ = v4123
	var v4124 int32
	_ = v4124
	var v4130 int64
	_ = v4130
	var v4132 int32
	_ = v4132
	var v4135 int32
	_ = v4135
	var v4146 int32
	_ = v4146
	var v4149 int32
	_ = v4149
	var v4150 int32
	_ = v4150
	var v4151 int32
	_ = v4151
	var v4154 int32
	_ = v4154
	var v4156 int32
	_ = v4156
	var v4219 int32
	_ = v4219
	var v4220 int32
	_ = v4220
	var v4221 int32
	_ = v4221
	var v4222 int32
	_ = v4222
	var v4223 int32
	_ = v4223
	var v4224 int32
	_ = v4224
	var v4227 int32
	_ = v4227
	var v4228 int32
	_ = v4228
	var v4229 int32
	_ = v4229
	var v4230 int32
	_ = v4230
	var v4239 int32
	_ = v4239
	var v4283 int32
	_ = v4283
	var v4286 int32
	_ = v4286
	var v4287 int32
	_ = v4287
	var v4294 int32
	_ = v4294
	var v4295 int32
	_ = v4295
	var v4297 int64
	_ = v4297
	var v4299 int64
	_ = v4299
	var v4301 int64
	_ = v4301
	var v4303 int64
	_ = v4303
	var v4305 int64
	_ = v4305
	var v4314 int32
	_ = v4314
	var v4315 int32
	_ = v4315
	var v4367 int32
	_ = v4367
	var v4369 int32
	_ = v4369
	var v4370 int32
	_ = v4370
	var v4372 int32
	_ = v4372
	var v4375 int32
	_ = v4375
	var v4376 int32
	_ = v4376
	var v4381 int32
	_ = v4381
	var v4387 int32
	_ = v4387
	var v4389 int32
	_ = v4389
	var v4391 int32
	_ = v4391
	var v4444 int32
	_ = v4444
	var v4445 int32
	_ = v4445
	var v4448 int32
	_ = v4448
	var v4452 int32
	_ = v4452
	var v4456 int32
	_ = v4456
	var v4505 int32
	_ = v4505
	var v4507 int32
	_ = v4507
	var v4510 int32
	_ = v4510
	var v4512 int32
	_ = v4512
	var v4513 int32
	_ = v4513
	var v4514 float64
	_ = v4514
	var v4515 int32
	_ = v4515
	var v4524 int32
	_ = v4524
	var v4526 int32
	_ = v4526
	var v4528 int32
	_ = v4528
	var v4529 int32
	_ = v4529
	var v4540 int32
	_ = v4540
	var v4580 int32
	_ = v4580
	var v4583 int32
	_ = v4583
	var v4584 int32
	_ = v4584
	var v4588 int32
	_ = v4588
	var v4591 int32
	_ = v4591
	var v4592 int32
	_ = v4592
	var v4594 int32
	_ = v4594
	var v4605 int32
	_ = v4605
	var v4609 int32
	_ = v4609
	var v4614 int32
	_ = v4614
	var v4616 int32
	_ = v4616
	var v4617 int32
	_ = v4617
	var v4620 int32
	_ = v4620
	var v4624 int32
	_ = v4624
	var v4626 int32
	_ = v4626
	var v4627 int32
	_ = v4627
	var v4635 int32
	_ = v4635
	var v4636 int32
	_ = v4636
	var v4642 int32
	_ = v4642
	var v4648 int32
	_ = v4648
	var v4650 int32
	_ = v4650
	var v4662 int32
	_ = v4662
	var v4703 int32
	_ = v4703
	var v4704 int32
	_ = v4704
	var v4705 int32
	_ = v4705
	var v4710 int32
	_ = v4710
	var v4759 int32
	_ = v4759
	var v4761 int32
	_ = v4761
	var v4766 int32
	_ = v4766
	var v4767 int32
	_ = v4767
	var v4769 int32
	_ = v4769
	var v4770 int32
	_ = v4770
	var v4773 int32
	_ = v4773
	var v4779 int32
	_ = v4779
	var v4784 int32
	_ = v4784
	var v4786 int32
	_ = v4786
	var v4790 int32
	_ = v4790
	var v4791 int32
	_ = v4791
	var v4793 int32
	_ = v4793
	var v4794 int32
	_ = v4794
	var v4799 int32
	_ = v4799
	var v4802 int32
	_ = v4802
	var v4803 int32
	_ = v4803
	var v4804 int32
	_ = v4804
	var v4857 int32
	_ = v4857
	var v4859 int32
	_ = v4859
	var v4860 int32
	_ = v4860
	var v4862 int32
	_ = v4862
	var v4864 int32
	_ = v4864
	var v4869 int32
	_ = v4869
	var v4870 int32
	_ = v4870
	var v4871 int32
	_ = v4871
	var v4873 int64
	_ = v4873
	var v4874 int64
	_ = v4874
	var v4884 int32
	_ = v4884
	var v4891 int32
	_ = v4891
	var v4918 int64
	_ = v4918
	var v4935 int64
	_ = v4935
	var v4936 int64
	_ = v4936
	var v4939 int64
	_ = v4939
	var v4943 int32
	_ = v4943
	var v4944 int32
	_ = v4944
	var v4946 int32
	_ = v4946
	var v4948 int32
	_ = v4948
	var v4950 int32
	_ = v4950
	var v4954 int32
	_ = v4954
	var v4956 int32
	_ = v4956
	var v4958 int32
	_ = v4958
	var v4967 int32
	_ = v4967
	var v4968 int32
	_ = v4968
	var v4971 int64
	_ = v4971
	var v4973 int64
	_ = v4973
	var v4977 int32
	_ = v4977
	var v4979 int32
	_ = v4979
	var v4984 int32
	_ = v4984
	var v4985 int32
	_ = v4985
	var v4986 int64
	_ = v4986
	var v4991 int32
	_ = v4991
	var v4992 int32
	_ = v4992
	var v4995 int32
	_ = v4995
	var v4996 int32
	_ = v4996
	var v5002 int32
	_ = v5002
	var v5007 int32
	_ = v5007
	var v5009 int32
	_ = v5009
	var v5010 int32
	_ = v5010
	var v5017 int32
	_ = v5017
	var v5019 int32
	_ = v5019
	var v5020 int32
	_ = v5020
	var v5021 int32
	_ = v5021
	var v5022 int32
	_ = v5022
	var v5030 int32
	_ = v5030
	var v5033 int32
	_ = v5033
	var v5034 int32
	_ = v5034
	var v5035 int32
	_ = v5035
	var v5036 int32
	_ = v5036
	var v5042 int32
	_ = v5042
	var v5047 int32
	_ = v5047
	var v5049 int32
	_ = v5049
	var v5051 int32
	_ = v5051
	var v5052 int32
	_ = v5052
	var v5053 int32
	_ = v5053
	var v5054 int32
	_ = v5054
	var v5056 int32
	_ = v5056
	var v5061 int32
	_ = v5061
	var v5068 int32
	_ = v5068
	var v5072 int32
	_ = v5072
	var v5077 int32
	_ = v5077
	var v5081 int32
	_ = v5081
	var v5088 int32
	_ = v5088
	var v5093 int32
	_ = v5093
	var v5099 int32
	_ = v5099
	var v5102 int32
	_ = v5102
	var v5103 int32
	_ = v5103
	var v5105 int32
	_ = v5105
	var v5106 int32
	_ = v5106
	var v5109 int32
	_ = v5109
	var v5115 int32
	_ = v5115
	var v5120 int32
	_ = v5120
	var v5126 int64
	_ = v5126
	var v5129 int32
	_ = v5129
	var v5131 int32
	_ = v5131
	var v5133 int32
	_ = v5133
	var v5136 int32
	_ = v5136
	var v5139 int32
	_ = v5139
	var v5189 int32
	_ = v5189
	var v5192 int32
	_ = v5192
	var v5194 int32
	_ = v5194
	var v5196 int32
	_ = v5196
	var v5212 int32
	_ = v5212
	var v5250 int32
	_ = v5250
	var v5251 int32
	_ = v5251
	var v5253 int32
	_ = v5253
	var v5254 int32
	_ = v5254
	var v5255 int32
	_ = v5255
	var v5258 int32
	_ = v5258
	var v5262 int32
	_ = v5262
	var v5268 int32
	_ = v5268
	var v5270 int32
	_ = v5270
	var v5276 int32
	_ = v5276
	var v5277 int32
	_ = v5277
	var v5280 int32
	_ = v5280
	var v5288 int32
	_ = v5288
	var v5296 int32
	_ = v5296
	var v5349 int32
	_ = v5349
	var v5355 int32
	_ = v5355
	var v5360 int32
	_ = v5360
	var v5412 int32
	_ = v5412
	var v5413 int32
	_ = v5413
	var v5421 int32
	_ = v5421
	var v5426 int32
	_ = v5426
	var v5466 int32
	_ = v5466
	var v5469 int32
	_ = v5469
	var v5471 int32
	_ = v5471
	var v5472 int32
	_ = v5472
	var v5474 int32
	_ = v5474
	var v5476 int32
	_ = v5476
	var v5482 int32
	_ = v5482
	var v5483 int32
	_ = v5483
	var v5485 int32
	_ = v5485
	var v5486 int32
	_ = v5486
	var v5487 int32
	_ = v5487
	var v5495 int32
	_ = v5495
	var v5500 int32
	_ = v5500
	var v5502 int32
	_ = v5502
	var v5556 int32
	_ = v5556
	var v5562 int32
	_ = v5562
	var v5566 int32
	_ = v5566
	var v5571 int32
	_ = v5571
	var v5573 int32
	_ = v5573
	var v5574 int32
	_ = v5574
	var v5577 int32
	_ = v5577
	var v5581 int32
	_ = v5581
	var v5583 int32
	_ = v5583
	var v5584 int32
	_ = v5584
	var v5592 int32
	_ = v5592
	var v5593 int32
	_ = v5593
	var v5599 int32
	_ = v5599
	var v5603 int32
	_ = v5603
	var v5606 int32
	_ = v5606
	var v5610 int32
	_ = v5610
	var v5616 int32
	_ = v5616
	var v5617 int32
	_ = v5617
	var v5620 int32
	_ = v5620
	var v5621 int32
	_ = v5621
	var v5624 int32
	_ = v5624
	var v5625 int32
	_ = v5625
	var v5626 float64
	_ = v5626
	var v5627 int32
	_ = v5627
	var v5630 int32
	_ = v5630
	var v5631 int32
	_ = v5631
	var v5638 int32
	_ = v5638
	var v5639 float64
	_ = v5639
	var v5640 float64
	_ = v5640
	var v5643 float64
	_ = v5643
	var v5645 int64
	_ = v5645
	var v5646 int64
	_ = v5646
	var v5649 int32
	_ = v5649
	var v5653 int32
	_ = v5653
	var v5654 int32
	_ = v5654
	var v5655 int32
	_ = v5655
	var v5659 int32
	_ = v5659
	var v5660 int32
	_ = v5660
	var v5661 int32
	_ = v5661
	var v5664 int64
	_ = v5664
	var v5665 int64
	_ = v5665
	var v5673 int64
	_ = v5673
	var v5679 int64
	_ = v5679
	var v5688 int64
	_ = v5688
	var v5691 int32
	_ = v5691
	var v5694 int32
	_ = v5694
	var v5695 int64
	_ = v5695
	var v5697 int32
	_ = v5697
	var v5698 int32
	_ = v5698
	var v5699 int32
	_ = v5699
	var v5707 int32
	_ = v5707
	var v5709 int32
	_ = v5709
	var v5710 int32
	_ = v5710
	var v5715 int32
	_ = v5715
	var v5716 int32
	_ = v5716
	var v5717 int64
	_ = v5717
	var v5723 int32
	_ = v5723
	var v5724 int32
	_ = v5724
	var v5725 int64
	_ = v5725
	var v5730 int32
	_ = v5730
	var v5733 int32
	_ = v5733
	var v5736 int32
	_ = v5736
	var v5737 int32
	_ = v5737
	var v5746 int32
	_ = v5746
	var v5750 int32
	_ = v5750
	var v5755 int32
	_ = v5755
	var v5758 int32
	_ = v5758
	var v5760 int32
	_ = v5760
	var v5761 int32
	_ = v5761
	var v5764 int32
	_ = v5764
	var v5768 int32
	_ = v5768
	var v5770 int32
	_ = v5770
	var v5771 int32
	_ = v5771
	var v5779 int32
	_ = v5779
	var v5780 int32
	_ = v5780
	var v5786 int32
	_ = v5786
	var v5795 int32
	_ = v5795
	var v5796 int32
	_ = v5796
	var v5797 int32
	_ = v5797
	var v5800 int64
	_ = v5800
	var v5801 int64
	_ = v5801
	var v5809 int64
	_ = v5809
	var v5810 int32
	_ = v5810
	var v5827 int64
	_ = v5827
	var v5831 int64
	_ = v5831
	var v5832 int64
	_ = v5832
	var v5839 int32
	_ = v5839
	var v5840 int32
	_ = v5840
	var v5843 int64
	_ = v5843
	var v5854 int32
	_ = v5854
	var v5856 int32
	_ = v5856
	var v5857 int64
	_ = v5857
	var v5859 int64
	_ = v5859
	var v5860 int64
	_ = v5860
	var v5864 int64
	_ = v5864
	var v5866 int64
	_ = v5866
	var v5867 int64
	_ = v5867
	var v5871 int64
	_ = v5871
	var v5873 int64
	_ = v5873
	var v5874 int64
	_ = v5874
	var v5878 int64
	_ = v5878
	var v5880 int64
	_ = v5880
	var v5881 int64
	_ = v5881
	var v5885 int64
	_ = v5885
	var v5887 int64
	_ = v5887
	var v5888 int64
	_ = v5888
	var v5893 int32
	_ = v5893
	var v5898 int32
	_ = v5898
	var v5899 int64
	_ = v5899
	var v5901 int64
	_ = v5901
	var v5902 int64
	_ = v5902
	var v5906 int64
	_ = v5906
	var v5908 int64
	_ = v5908
	var v5909 int64
	_ = v5909
	var v5913 int64
	_ = v5913
	var v5915 int64
	_ = v5915
	var v5916 int64
	_ = v5916
	var v5920 int64
	_ = v5920
	var v5922 int64
	_ = v5922
	var v5923 int64
	_ = v5923
	var v5927 int64
	_ = v5927
	var v5929 int64
	_ = v5929
	var v5930 int64
	_ = v5930
	var v5934 int64
	_ = v5934
	var v5936 int64
	_ = v5936
	var v5937 int64
	_ = v5937
	var v5941 int64
	_ = v5941
	var v5943 int64
	_ = v5943
	var v5944 int64
	_ = v5944
	var v5948 int64
	_ = v5948
	var v5950 int64
	_ = v5950
	var v5951 int64
	_ = v5951
	var v5955 int64
	_ = v5955
	var v5957 int64
	_ = v5957
	var v5958 int64
	_ = v5958
	var v5962 int64
	_ = v5962
	var v5964 int64
	_ = v5964
	var v5965 int64
	_ = v5965
	var v5969 int64
	_ = v5969
	var v5971 int64
	_ = v5971
	var v5972 int64
	_ = v5972
	var v5976 int64
	_ = v5976
	var v5978 int64
	_ = v5978
	var v5979 int64
	_ = v5979
	var v5983 int64
	_ = v5983
	var v5985 int64
	_ = v5985
	var v5986 int64
	_ = v5986
	var v5990 int64
	_ = v5990
	var v5992 int64
	_ = v5992
	var v5993 int64
	_ = v5993
	var v5997 int64
	_ = v5997
	var v5999 int64
	_ = v5999
	var v6000 int64
	_ = v6000
	var v6004 int64
	_ = v6004
	var v6006 int64
	_ = v6006
	var v6007 int64
	_ = v6007
	var v6011 int64
	_ = v6011
	var v6012 int64
	_ = v6012
	var v6013 int64
	_ = v6013
	var v6014 int64
	_ = v6014
	var v6015 int64
	_ = v6015
	var v6016 int64
	_ = v6016
	var v6020 int32
	_ = v6020
	var v6024 int32
	_ = v6024
	var v6027 int32
	_ = v6027
	var v6028 int32
	_ = v6028
	var v6035 int32
	_ = v6035
	var v6037 int32
	_ = v6037
	var v6038 int32
	_ = v6038
	var v6039 int64
	_ = v6039
	var v6040 int32
	_ = v6040
	var v6045 int32
	_ = v6045
	var v6049 int32
	_ = v6049
	var v6050 int32
	_ = v6050
	var v6051 int32
	_ = v6051
	var v6052 int32
	_ = v6052
	var v6055 float64
	_ = v6055
	var v6060 float64
	_ = v6060
	var v6069 int32
	_ = v6069
	var v6070 float64
	_ = v6070
	var v6071 int64
	_ = v6071
	var v6072 int64
	_ = v6072
	var v6081 int32
	_ = v6081
	var v6082 int64
	_ = v6082
	var v6085 int32
	_ = v6085
	var v6092 int32
	_ = v6092
	var v6093 int64
	_ = v6093
	var v6094 int32
	_ = v6094
	var v6095 int32
	_ = v6095
	var v6101 int32
	_ = v6101
	var v6106 int32
	_ = v6106
	var v6107 int32
	_ = v6107
	var v6110 int32
	_ = v6110
	var v6111 int32
	_ = v6111
	var v6119 int32
	_ = v6119
	var v6122 int32
	_ = v6122
	var v6125 int32
	_ = v6125
	var v6126 int32
	_ = v6126
	var v6136 int32
	_ = v6136
	var v6139 int32
	_ = v6139
	var v6140 int64
	_ = v6140
	var v6143 float64
	_ = v6143
	var v6148 float64
	_ = v6148
	var v6152 int32
	_ = v6152
	var v6157 int32
	_ = v6157
	var v6158 int32
	_ = v6158
	var v6159 int32
	_ = v6159
	var v6160 int32
	_ = v6160
	var v6169 int32
	_ = v6169
	var v6170 int32
	_ = v6170
	var v6173 int32
	_ = v6173
	var v6175 int32
	_ = v6175
	var v6180 int32
	_ = v6180
	var v6181 int32
	_ = v6181
	var v6186 int32
	_ = v6186
	var v6187 int32
	_ = v6187
	var v6188 int32
	_ = v6188
	var v6189 int32
	_ = v6189
	var v6191 int32
	_ = v6191
	var v6192 int32
	_ = v6192
	var v6193 int64
	_ = v6193
	var v6196 float64
	_ = v6196
	var v6201 float64
	_ = v6201
	var v6205 int32
	_ = v6205
	var v6209 int32
	_ = v6209
	var v6210 int32
	_ = v6210
	var v6213 int32
	_ = v6213
	var v6220 int32
	_ = v6220
	var v6221 int32
	_ = v6221
	var v6224 int32
	_ = v6224
	var v6233 int32
	_ = v6233
	var v6234 int32
	_ = v6234
	var v6235 int32
	_ = v6235
	var v6245 int32
	_ = v6245
	var v6246 int32
	_ = v6246
	var v6291 int32
	_ = v6291
	var v6292 int32
	_ = v6292
	var v6294 int32
	_ = v6294
	var v6296 int32
	_ = v6296
	var v6297 int64
	_ = v6297
	var v6298 int32
	_ = v6298
	var v6299 int32
	_ = v6299
	var v6310 int32
	_ = v6310
	var v6311 int32
	_ = v6311
	var v6312 int32
	_ = v6312
	var v6316 int32
	_ = v6316
	var v6369 int32
	_ = v6369
	var v6371 int32
	_ = v6371
	var v6372 int64
	_ = v6372
	var v6383 int32
	_ = v6383
	var v6385 int32
	_ = v6385
	var v6389 int64
	_ = v6389
	var v6392 float64
	_ = v6392
	var v6396 int64
	_ = v6396
	var v6408 int32
	_ = v6408
	var v6409 int64
	_ = v6409
	var v6410 int64
	_ = v6410
	var v6412 int32
	_ = v6412
	var v6413 int32
	_ = v6413
	var v6420 float64
	_ = v6420
	var v6422 float64
	_ = v6422
	var v6428 float64
	_ = v6428
	var v6437 float64
	_ = v6437
	var v6438 float64
	_ = v6438
	var v6442 int32
	_ = v6442
	var v6447 int32
	_ = v6447
	var v6455 int32
	_ = v6455
	var v6456 int64
	_ = v6456
	var v6458 int64
	_ = v6458
	var v6460 int64
	_ = v6460
	var v6462 int64
	_ = v6462
	var v6464 int64
	_ = v6464
	var v6470 int32
	_ = v6470
	var v6471 int32
	_ = v6471
	var v6472 int32
	_ = v6472
	var v6474 float64
	_ = v6474
	var v6486 int32
	_ = v6486
	var v6490 int32
	_ = v6490
	var v6493 int32
	_ = v6493
	var v6494 int32
	_ = v6494
	var v6500 int32
	_ = v6500
	var v6503 int32
	_ = v6503
	var v6505 int32
	_ = v6505
	var v6506 int32
	_ = v6506
	var v6507 int32
	_ = v6507
	var v6511 int32
	_ = v6511
	var v6516 int32
	_ = v6516
	var v6517 int32
	_ = v6517
	var v6519 int32
	_ = v6519
	var v6570 int32
	_ = v6570
	var v6575 int32
	_ = v6575
	var v6624 int32
	_ = v6624
	var v6625 int32
	_ = v6625
	var v6627 int32
	_ = v6627
	var v6629 int32
	_ = v6629
	var v6631 int32
	_ = v6631
	var v6633 int32
	_ = v6633
	var v6635 int32
	_ = v6635
	var v6636 int32
	_ = v6636
	v4 = int32(0)
	v40 = int64(0)
	v51 = m.G0
	v53 = v51 - int32(1696)
	m.G0 = v53
	v56 = *(*int64)(unsafe.Add(mBase, _c_F_heap_vacuum_rel[0]))
	*(*int64)(unsafe.Add(mBase, uint32(v53)+808)) = v56
	v59 = *(*int64)(unsafe.Add(mBase, _c_F_heap_vacuum_rel[1]))
	*(*int64)(unsafe.Add(mBase, uint32(v53)+800)) = v59
	v62 = *(*int64)(unsafe.Add(mBase, _c_F_heap_vacuum_rel[2]))
	*(*int64)(unsafe.Add(mBase, uint32(v53)+792)) = v62
	v65 = *(*int64)(unsafe.Add(mBase, _c_F_heap_vacuum_rel[3]))
	*(*int64)(unsafe.Add(mBase, uint32(v53)+784)) = v65
	v68 = *(*int64)(unsafe.Add(mBase, _c_F_heap_vacuum_rel[4]))
	*(*int64)(unsafe.Add(mBase, uint32(v53)+776)) = v68
	base.MemoryCopy(m, v53+int32(648), int32(_a_F_heap_vacuum_rel_0), int32(128))
	v75 = *(*int32)(unsafe.Add(mBase, uint32(l1)))
	v77 = v75 & int32(4)
	if v77 == v4 {
		goto L2
	} else {
		goto L3
	}
L1:
	;
	v109 = int32(base.Ui32(v77) >> (uint(int32(2)) % 32))
	v113 = m.G0
	v114 = int32(16)
	v115 = v113 - v114
	m.G0 = v115
	F_gettimeofday(m, v115)
	mBase = m.M
	v118 = *(*int64)(unsafe.Add(mBase, uint32(v115)))
	v119 = int64(*(*int32)(unsafe.Add(mBase, uint32(v115)+8)))
	m.G0 = v115 + v114
	v127 = v119 + v118*int64(1000000) - int64(946684800000000)
	goto L9
L2:
	;
	v82 = *(*int32)(unsafe.Add(mBase, _c_F_heap_vacuum_rel[5]))
	if v82 != int32(4) {
		v105 = v4
		v106 = v40
		v107 = int64(0)
		goto L1
	} else {
		goto L5
	}
L3:
	;
	goto L4
L4:
	;
	F_getrusage(m, v53+int32(832))
	mBase = m.M
	F_gettimeofday(m, v53+int32(816))
	mBase = m.M
	goto L7
L5:
	;
	v86 = *(*int32)(unsafe.Add(mBase, uint32(l1)+24))
	if v86 < int32(0) {
		v105 = v4
		v106 = v40
		v107 = int64(0)
		goto L1
	} else {
		goto L6
	}
L6:
	;
	goto L4
L7:
	;
	v95 = int32(1)
	v98 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_heap_vacuum_rel[6])))
	if v98 != v95 {
		v105 = v95
		v106 = v40
		v107 = int64(0)
		goto L1
	} else {
		goto L8
	}
L8:
	;
	v102 = *(*int64)(unsafe.Add(mBase, _c_F_heap_vacuum_rel[7]))
	v104 = *(*int64)(unsafe.Add(mBase, _c_F_heap_vacuum_rel[8]))
	v105 = v95
	v106 = v102
	v107 = v104
	goto L1
L9:
	;
	v129 = *(*int32)(unsafe.Add(mBase, uint32(l0)+56))
	v132 = *(*int32)(unsafe.Add(mBase, _c_F_heap_vacuum_rel[9]))
	if v132 == int32(0) {
		goto L11
	} else {
		goto L12
	}
L10:
	;
	v178 = *(*int32)(unsafe.Add(mBase, _c_F_heap_vacuum_rel[5]))
	if v178 == int32(4) {
		goto L14
	} else {
		goto L15
	}
L11:
	;
	goto L10
L12:
	;
	v136 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_heap_vacuum_rel[10])))
	if v136&int32(1) == int32(0) {
		goto L11
	} else {
		goto L13
	}
L13:
	;
	v141 = int32(_a_F_heap_vacuum_rel_1)
	v143 = *(*int32)(unsafe.Add(mBase, _c_F_heap_vacuum_rel[11]))
	v144 = int32(1)
	*(*int32)(unsafe.Add(mBase, _c_F_heap_vacuum_rel[11])) = v143 + v144
	v147 = *(*int32)(unsafe.Add(mBase, uint32(v132)))
	*(*int32)(unsafe.Add(mBase, uint32(v132))) = v147 + v144
	v151 = int32(0)
	v153 = int32(_a_F_heap_vacuum_rel_2)
	v154 = base.AtomicRmwOr32(m, v151, v153, v151)
	*(*int32)(unsafe.Add(mBase, uint32(v132)+220)) = int32(1)
	*(*int32)(unsafe.Add(mBase, uint32(v132)+224)) = v129
	base.MemoryFill(m, v132+int32(232), v151, int32(160))
	v165 = base.AtomicRmwOr32(m, v151, v153, v151)
	v166 = *(*int32)(unsafe.Add(mBase, uint32(v132)))
	*(*int32)(unsafe.Add(mBase, uint32(v132))) = v166 + v144
	v172 = *(*int32)(unsafe.Add(mBase, _c_F_heap_vacuum_rel[11]))
	*(*int32)(unsafe.Add(mBase, _c_F_heap_vacuum_rel[11])) = v172 - v144
	goto L11
L14:
	;
	v183 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l1)+20)))
	if v183 != 0 {
		goto L17
	} else {
		goto L18
	}
L15:
	;
	v186 = int64(1)
	goto L16
L16:
	;
	v189 = *(*int32)(unsafe.Add(mBase, _c_F_heap_vacuum_rel[9]))
	if v189 == int32(0) {
		goto L21
	} else {
		goto L22
	}
L17:
	;
	v184 = int64(3)
	goto L19
L18:
	;
	v184 = int64(2)
	goto L19
L19:
	;
	v186 = v184
	goto L16
L20:
	;
	v231 = F_palloc0(m, int32(280))
	mBase = m.M
	v232 = m.ExcPending
	if v232 != 0 {
		goto L24
	} else {
		goto L25
	}
L21:
	;
	goto L20
L22:
	;
	v193 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_heap_vacuum_rel[10])))
	if v193&int32(1) == int32(0) {
		goto L21
	} else {
		goto L23
	}
L23:
	;
	v198 = int32(_a_F_heap_vacuum_rel_1)
	v200 = *(*int32)(unsafe.Add(mBase, _c_F_heap_vacuum_rel[11]))
	v201 = int32(1)
	*(*int32)(unsafe.Add(mBase, _c_F_heap_vacuum_rel[11])) = v200 + v201
	v204 = *(*int32)(unsafe.Add(mBase, uint32(v189)))
	*(*int32)(unsafe.Add(mBase, uint32(v189))) = v204 + v201
	v208 = int32(0)
	v210 = int32(_a_F_heap_vacuum_rel_2)
	v211 = base.AtomicRmwOr32(m, v208, v210, v208)
	*(*int64)(unsafe.Add(mBase, uint32(v189+int32(96))+232)) = v186
	v219 = base.AtomicRmwOr32(m, v208, v210, v208)
	v220 = *(*int32)(unsafe.Add(mBase, uint32(v189)))
	*(*int32)(unsafe.Add(mBase, uint32(v189))) = v220 + v201
	v226 = *(*int32)(unsafe.Add(mBase, _c_F_heap_vacuum_rel[11]))
	*(*int32)(unsafe.Add(mBase, _c_F_heap_vacuum_rel[11])) = v226 - v201
	goto L21
L24:
	;
	return
L25:
	;
	v234 = *(*int32)(unsafe.Add(mBase, _c_F_heap_vacuum_rel[12]))
	v235 = F_get_database_name(m, v234)
	mBase = m.M
	v236 = m.ExcPending
	if v236 != 0 {
		goto L24
	} else {
		goto L26
	}
L26:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v231)+68)) = v235
	v238 = *(*int32)(unsafe.Add(mBase, uint32(l0)+48))
	v239 = *(*int32)(unsafe.Add(mBase, uint32(v238)+68))
	v240 = F_get_namespace_name(m, v239)
	mBase = m.M
	v241 = m.ExcPending
	if v241 != 0 {
		goto L24
	} else {
		goto L27
	}
L27:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v231)+72)) = v240
	v243 = *(*int32)(unsafe.Add(mBase, uint32(l0)+48))
	v246 = F_pstrdup(m, v243+int32(4))
	mBase = m.M
	v247 = m.ExcPending
	if v247 != 0 {
		goto L24
	} else {
		goto L28
	}
L28:
	;
	*(*uint8)(unsafe.Add(mBase, uint32(v231)+96)) = uint8(v109)
	v249 = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v231)+92)) = v249
	*(*int32)(unsafe.Add(mBase, uint32(v231)+80)) = v249
	*(*int32)(unsafe.Add(mBase, uint32(v231)+76)) = v246
	*(*int32)(unsafe.Add(mBase, uint32(v53)+644)) = v231
	*(*int32)(unsafe.Add(mBase, uint32(v53)+640)) = int32(189)
	v257 = int32(_a_F_heap_vacuum_rel_3)
	v258 = *(*int32)(unsafe.Add(mBase, _c_F_heap_vacuum_rel[13]))
	*(*int32)(unsafe.Add(mBase, _c_F_heap_vacuum_rel[13])) = v53 + int32(636)
	*(*int32)(unsafe.Add(mBase, uint32(v53)+636)) = v258
	*(*int32)(unsafe.Add(mBase, uint32(v231))) = l0
	v267 = v231 + int32(8)
	v269 = v231 + int32(4)
	F_vac_open_indexes(m, l0, int32(3), v267, v269)
	mBase = m.M
	v271 = m.ExcPending
	if v271 != 0 {
		goto L24
	} else {
		goto L29
	}
L29:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v231)+12)) = l2
	if v105 == int32(0) {
		v378 = v4
		goto L30
	} else {
		goto L31
	}
L30:
	;
	v402 = int32(0)
	*(*uint8)(unsafe.Add(mBase, _c_F_heap_vacuum_rel[14])) = uint8(v402)
	v404 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(v231)+24)) = uint8(v404)
	v406 = int32(257)
	*(*uint16)(unsafe.Add(mBase, uint32(v231)+22)) = uint16(v406)
	v408 = *(*int32)(unsafe.Add(mBase, uint32(l1)+36))
	v409 = int32(2)
	*(*uint8)(unsafe.Add(mBase, uint32(v231)+25)) = uint8(base.B2i32(v408 != v409))
	v412 = *(*int32)(unsafe.Add(mBase, uint32(l1)+32))
	switch v412 - v409 {
	case 0:
		goto L41
	case 1:
		goto L40
	default:
		goto L39
	}
L31:
	;
	v275 = *(*int32)(unsafe.Add(mBase, uint32(v267)))
	if v275 <= int32(0) {
		v378 = v4
		goto L30
	} else {
		goto L32
	}
L32:
	;
	v279 = F_palloc_mul(m, int32(4), v275)
	mBase = m.M
	v280 = m.ExcPending
	if v280 != 0 {
		goto L24
	} else {
		goto L33
	}
L33:
	;
	v281 = *(*int32)(unsafe.Add(mBase, uint32(v267)))
	if v281 <= int32(0) {
		v378 = v279
		goto L30
	} else {
		goto L34
	}
L34:
	;
	v287 = int32(0)
	goto L35
L35:
	;
	v336 = v287 << (uint(int32(2)) % 32)
	v338 = *(*int32)(unsafe.Add(mBase, uint32(v269)))
	v340 = *(*int32)(unsafe.Add(mBase, uint32(v338+v336)))
	v341 = *(*int32)(unsafe.Add(mBase, uint32(v340)+48))
	v344 = F_pstrdup(m, v341+int32(4))
	mBase = m.M
	v345 = m.ExcPending
	if v345 != 0 {
		goto L24
	} else {
		goto L37
	}
L36:
	;
	v378 = v279
	goto L30
L37:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v279+v336))) = v344
	v348 = v287 + int32(1)
	v349 = *(*int32)(unsafe.Add(mBase, uint32(v267)))
	if v348 < v349 {
		v287 = v348
		goto L35
	} else {
		goto L38
	}
L38:
	;
	goto L36
L39:
	;
	v419 = int64(0)
	*(*int64)(unsafe.Add(mBase, uint32(v231)+120)) = v419
	*(*int64)(unsafe.Add(mBase, uint32(v231)+112)) = v419
	*(*int64)(unsafe.Add(mBase, uint32(v231)+140)) = v419
	*(*int64)(unsafe.Add(mBase, uint32(v231)+148)) = v419
	*(*int64)(unsafe.Add(mBase, uint32(v231)+156)) = v419
	*(*int32)(unsafe.Add(mBase, uint32(v231)+164)) = int32(0)
	v431 = *(*int32)(unsafe.Add(mBase, uint32(v231)+8))
	v434 = F_palloc0(m, v431<<(uint(int32(2))%32))
	mBase = m.M
	v435 = m.ExcPending
	if v435 != 0 {
		goto L24
	} else {
		goto L42
	}
L40:
	;
	v417 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v231)+22)) = uint8(v417)
	goto L39
L41:
	;
	v415 = int32(0)
	*(*uint16)(unsafe.Add(mBase, uint32(v231)+23)) = uint16(v415)
	goto L39
L42:
	;
	v436 = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v231)+136)) = v436
	*(*int64)(unsafe.Add(mBase, uint32(v231)+128)) = int64(0)
	*(*int32)(unsafe.Add(mBase, uint32(v231)+168)) = v434
	v442 = v231 + int32(172)
	base.MemoryFill(m, v442, v436, int32(76))
	v447 = v231 + int32(28)
	v448 = F_vacuum_get_cutoffs(m, l0, l1, v447)
	mBase = m.M
	v449 = m.ExcPending
	if v449 != 0 {
		goto L24
	} else {
		goto L43
	}
L43:
	;
	*(*uint8)(unsafe.Add(mBase, uint32(v231)+20)) = uint8(v448)
	v452 = F_RelationGetNumberOfBlocksInFork(m, l0, int32(0))
	mBase = m.M
	v453 = m.ExcPending
	if v453 != 0 {
		goto L24
	} else {
		goto L44
	}
L44:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v231)+108)) = v452
	v455 = F_GlobalVisHorizonKindForRel(m, l0)
	mBase = m.M
	v458 = *(*int32)(unsafe.Add(mBase, uint32(v455<<(uint(int32(2))%32))+uint32(_c_F_heap_vacuum_rel[15])))
	goto L45
L45:
	;
	v459 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v231)+64)) = uint8(v459)
	*(*int32)(unsafe.Add(mBase, uint32(v231)+52)) = v458
	v462 = *(*int64)(unsafe.Add(mBase, uint32(v231)+36))
	*(*int64)(unsafe.Add(mBase, uint32(v231)+56)) = v462
	v464 = *(*int32)(unsafe.Add(mBase, uint32(l1)))
	v466 = v464 & int32(256)
	if v466 != 0 {
		goto L46
	} else {
		goto L47
	}
L46:
	;
	v467 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(v231)+20)) = uint8(v467)
	goto L48
L47:
	;
	goto L48
L48:
	;
	*(*uint8)(unsafe.Add(mBase, uint32(v231)+21)) = uint8(base.B2i32(v466 == int32(0)))
	*(*int64)(unsafe.Add(mBase, uint32(v231)+272)) = int64(0)
	*(*int64)(unsafe.Add(mBase, uint32(v231)+264)) = int64(4294967295)
	v476 = *(*float64)(unsafe.Add(mBase, uint32(l1)+48))
	if base.F64_eq(v476, float64(0)) != 0 {
		goto L49
	} else {
		goto L50
	}
L49:
	;
	v543 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v231)+20)))
	if v543 != 0 {
		goto L63
	} else {
		goto L64
	}
L50:
	;
	v479 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v231)+20)))
	if v479 != 0 {
		goto L49
	} else {
		goto L51
	}
L51:
	;
	v480 = *(*int32)(unsafe.Add(mBase, uint32(v231)+108))
	if base.Ui32(v480) < base.Ui32(int32(_a_F_heap_vacuum_rel_4)) {
		goto L49
	} else {
		goto L52
	}
L52:
	;
	v483 = *(*int32)(unsafe.Add(mBase, uint32(v447)))
	if base.Ui32(v483) < base.Ui32(int32(3)) {
		goto L54
	} else {
		goto L55
	}
L53:
	;
	v502 = *(*int32)(unsafe.Add(mBase, uint32(v231)))
	F_visibilitymap_count(m, v502, v53+int32(1040), v53+int32(592))
	mBase = m.M
	v508 = m.ExcPending
	if v508 != 0 {
		goto L24
	} else {
		goto L60
	}
L54:
	;
	v493 = *(*int32)(unsafe.Add(mBase, uint32(v231)+32))
	if v493 == int32(0) {
		goto L49
	} else {
		goto L58
	}
L55:
	;
	v486 = *(*int32)(unsafe.Add(mBase, uint32(v231)+44))
	if base.Ui32(v486) < base.Ui32(int32(3)) {
		goto L54
	} else {
		goto L56
	}
L56:
	;
	if v483-v486 < int32(0) {
		goto L53
	} else {
		goto L57
	}
L57:
	;
	goto L54
L58:
	;
	v496 = *(*int32)(unsafe.Add(mBase, uint32(v231)+48))
	if int32(0) <= v493-v496 {
		goto L49
	} else {
		goto L59
	}
L59:
	;
	goto L53
L60:
	;
	v509 = *(*int32)(unsafe.Add(mBase, uint32(v53)+1040))
	v510 = *(*int32)(unsafe.Add(mBase, uint32(v53)+592))
	v515 = base.I32_trunc_sat_f64_u(base.F64_mul(base.F64_convert_i32_u(v509-v510), float64(0.2)))
	*(*int32)(unsafe.Add(mBase, uint32(v231)+268)) = v515
	if v515 == int32(0) {
		goto L49
	} else {
		goto L61
	}
L61:
	;
	v520 = Fn14349(m, int64(32))
	mBase = m.M
	goto L62
L62:
	;
	v522 = v520 & int32(4095)
	*(*int32)(unsafe.Add(mBase, uint32(v231)+264)) = v522
	v524 = *(*float64)(unsafe.Add(mBase, uint32(l1)+48))
	v527 = base.I32_trunc_sat_f64_u(base.F64_mul(v524, float64(4096)))
	*(*int32)(unsafe.Add(mBase, uint32(v231)+272)) = v527
	*(*int32)(unsafe.Add(mBase, uint32(v231)+276)) = base.I32_trunc_sat_f32_u(base.F32_mul(base.F32_add(base.F32_mul(base.F32_convert_i32_u(v522), float32(-0.00024414062)), float32(1)), base.F32_convert_i32_u(v527)))
	goto L49
L63:
	;
	v544 = int64(2)
	goto L65
L64:
	;
	v544 = int64(1)
	goto L65
L65:
	;
	v547 = *(*int32)(unsafe.Add(mBase, _c_F_heap_vacuum_rel[9]))
	if v547 == int32(0) {
		goto L67
	} else {
		goto L68
	}
L66:
	;
	if v77 == int32(0) {
		goto L70
	} else {
		goto L71
	}
L67:
	;
	goto L66
L68:
	;
	v551 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_heap_vacuum_rel[10])))
	if v551&int32(1) == int32(0) {
		goto L67
	} else {
		goto L69
	}
L69:
	;
	v556 = int32(_a_F_heap_vacuum_rel_1)
	v558 = *(*int32)(unsafe.Add(mBase, _c_F_heap_vacuum_rel[11]))
	v559 = int32(1)
	*(*int32)(unsafe.Add(mBase, _c_F_heap_vacuum_rel[11])) = v558 + v559
	v562 = *(*int32)(unsafe.Add(mBase, uint32(v547)))
	*(*int32)(unsafe.Add(mBase, uint32(v547))) = v562 + v559
	v566 = int32(0)
	v568 = int32(_a_F_heap_vacuum_rel_2)
	v569 = base.AtomicRmwOr32(m, v566, v568, v566)
	*(*int64)(unsafe.Add(mBase, uint32(v547+int32(88))+232)) = v544
	v577 = base.AtomicRmwOr32(m, v566, v568, v566)
	v578 = *(*int32)(unsafe.Add(mBase, uint32(v547)))
	*(*int32)(unsafe.Add(mBase, uint32(v547))) = v578 + v559
	v584 = *(*int32)(unsafe.Add(mBase, _c_F_heap_vacuum_rel[11]))
	*(*int32)(unsafe.Add(mBase, _c_F_heap_vacuum_rel[11])) = v584 - v559
	goto L67
L70:
	;
	v619 = F_lazy_check_wraparound_failsafe(m, v231)
	mBase = m.M
	v620 = m.ExcPending
	if v620 != 0 {
		goto L24
	} else {
		goto L82
	}
L71:
	;
	v590 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v231)+20)))
	v593 = F_errstart(m, int32(17), int32(0))
	mBase = m.M
	v594 = m.ExcPending
	if v594 != 0 {
		goto L24
	} else {
		goto L72
	}
L72:
	;
	if v593 == int32(0) {
		goto L70
	} else {
		goto L73
	}
L73:
	;
	v597 = *(*int64)(unsafe.Add(mBase, uint32(v231)+68))
	v598 = *(*int32)(unsafe.Add(mBase, uint32(v231)+76))
	*(*int32)(unsafe.Add(mBase, uint32(v53)+584)) = v598
	*(*int64)(unsafe.Add(mBase, uint32(v53)+576)) = v597
	v604 = v590 & int32(1)
	if v604 != 0 {
		goto L74
	} else {
		goto L75
	}
L74:
	;
	v605 = int32(_a_F_heap_vacuum_rel_5)
	goto L76
L75:
	;
	v605 = int32(_a_F_heap_vacuum_rel_6)
	goto L76
L76:
	;
	F_errmsg(m, v605, v53+int32(576))
	mBase = m.M
	v609 = m.ExcPending
	if v609 != 0 {
		goto L24
	} else {
		goto L77
	}
L77:
	;
	if v604 != 0 {
		goto L78
	} else {
		goto L79
	}
L78:
	;
	v613 = int32(848)
	goto L80
L79:
	;
	v613 = int32(853)
	goto L80
L80:
	;
	F_errfinish(m, int32(_a_F_heap_vacuum_rel_7), v613, int32(_a_F_heap_vacuum_rel_8))
	mBase = m.M
	v616 = m.ExcPending
	if v616 != 0 {
		goto L24
	} else {
		goto L81
	}
L81:
	;
	goto L70
L82:
	;
	v622 = *(*int32)(unsafe.Add(mBase, _c_F_heap_vacuum_rel[16]))
	v624 = *(*int32)(unsafe.Add(mBase, _c_F_heap_vacuum_rel[17]))
	if v622 != int32(-1) {
		goto L83
	} else {
		goto L84
	}
L83:
	;
	v627 = v622
	goto L85
L84:
	;
	v627 = v624
	goto L85
L85:
	;
	v629 = *(*int32)(unsafe.Add(mBase, _c_F_heap_vacuum_rel[5]))
	if v629 == int32(4) {
		goto L86
	} else {
		goto L87
	}
L86:
	;
	v632 = v627
	goto L88
L87:
	;
	v632 = v624
	goto L88
L88:
	;
	v633 = *(*int32)(unsafe.Add(mBase, uint32(l1)+56))
	if v633 < int32(0) {
		goto L90
	} else {
		goto L91
	}
L89:
	;
	v1440 = v231 + int32(60)
	v1442 = v231 + int32(56)
	v1444 = v231 + int32(192)
	v1448 = v231 + int32(112)
	*(*int32)(unsafe.Add(mBase, uint32(v231)+100)) = v1438
	v1450 = *(*int32)(unsafe.Add(mBase, uint32(v231)+268))
	v1451 = *(*int32)(unsafe.Add(mBase, uint32(v231)+108))
	v1452 = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v53)+1036)) = v1452
	v1455 = *(*int32)(unsafe.Add(mBase, _c_F_heap_vacuum_rel[18]))
	*(*int32)(unsafe.Add(mBase, uint32(v53)+1032)) = v1455
	v1458 = *(*int64)(unsafe.Add(mBase, _c_F_heap_vacuum_rel[19]))
	*(*int64)(unsafe.Add(mBase, uint32(v53)+1024)) = v1458
	v1460 = base.I64_extend_i32_u(v1451)
	*(*int64)(unsafe.Add(mBase, uint32(v53)+1000)) = v1460
	*(*int64)(unsafe.Add(mBase, uint32(v53)+992)) = int64(1)
	v1464 = *(*int32)(unsafe.Add(mBase, uint32(v231)+104))
	v1465 = int64(*(*uint32)(unsafe.Add(mBase, uint32(v1464))))
	*(*int64)(unsafe.Add(mBase, uint32(v53)+1008)) = v1465
	goto L228
L90:
	;
	v1377 = F_palloc(m, int32(16))
	mBase = m.M
	v1378 = m.ExcPending
	if v1378 != 0 {
		goto L24
	} else {
		goto L224
	}
L91:
	;
	v636 = *(*int32)(unsafe.Add(mBase, uint32(v267)))
	if v636 < int32(2) {
		goto L90
	} else {
		goto L92
	}
L92:
	;
	v639 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v231)+23)))
	if v639 != int32(1) {
		goto L90
	} else {
		goto L93
	}
L93:
	;
	v642 = *(*int32)(unsafe.Add(mBase, uint32(v231)))
	v643 = *(*int32)(unsafe.Add(mBase, uint32(v642)+48))
	v644 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v643)+118)))
	if v644 == int32(116) {
		goto L95
	} else {
		goto L96
	}
L94:
	;
	v1316 = *(*int32)(unsafe.Add(mBase, uint32(v231)+16))
	if v1316 == int32(0) {
		goto L90
	} else {
		goto L222
	}
L95:
	;
	if v633 == int32(0) {
		goto L94
	} else {
		goto L98
	}
L96:
	;
	goto L97
L97:
	;
	v667 = *(*int32)(unsafe.Add(mBase, uint32(v231)+4))
	v670 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v231)+96)))
	if v670 != 0 {
		goto L104
	} else {
		goto L105
	}
L98:
	;
	v651 = F_errstart(m, int32(19), int32(0))
	mBase = m.M
	v652 = m.ExcPending
	if v652 != 0 {
		goto L24
	} else {
		goto L99
	}
L99:
	;
	if v651 == int32(0) {
		goto L94
	} else {
		goto L100
	}
L100:
	;
	v655 = *(*int32)(unsafe.Add(mBase, uint32(v231)+76))
	*(*int32)(unsafe.Add(mBase, uint32(v53)+560)) = v655
	F_errmsg(m, int32(_a_F_heap_vacuum_rel_9), v53+int32(560))
	mBase = m.M
	v661 = m.ExcPending
	if v661 != 0 {
		goto L24
	} else {
		goto L101
	}
L101:
	;
	F_errfinish(m, int32(_a_F_heap_vacuum_rel_7), int32(3467), int32(_a_F_heap_vacuum_rel_10))
	mBase = m.M
	v666 = m.ExcPending
	if v666 != 0 {
		goto L24
	} else {
		goto L102
	}
L102:
	;
	goto L94
L103:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v231)+16)) = v1264
	goto L94
L104:
	;
	v671 = int32(17)
	goto L106
L105:
	;
	v671 = int32(13)
	goto L106
L106:
	;
	v672 = *(*int32)(unsafe.Add(mBase, uint32(v231)+12))
	v674 = F_palloc0_mul(m, int32(1), v636)
	mBase = m.M
	v675 = m.ExcPending
	if v675 != 0 {
		goto L24
	} else {
		goto L107
	}
L107:
	;
	if v636 <= int32(0) {
		goto L109
	} else {
		goto L110
	}
L108:
	;
	v843 = F_palloc0(m, int32(72))
	mBase = m.M
	v844 = m.ExcPending
	if v844 != 0 {
		goto L24
	} else {
		goto L138
	}
L109:
	;
	F_pfree(m, v674)
	mBase = m.M
	v840 = m.ExcPending
	if v840 != 0 {
		goto L24
	} else {
		goto L137
	}
L110:
	;
	v679 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_heap_vacuum_rel[20])))
	if v679&int32(1) == int32(0) {
		goto L109
	} else {
		goto L111
	}
L111:
	;
	v687 = *(*int32)(unsafe.Add(mBase, _c_F_heap_vacuum_rel[5]))
	if v687 == int32(4) {
		goto L112
	} else {
		goto L113
	}
L112:
	;
	v690 = int32(_a_F_heap_vacuum_rel_11)
	goto L114
L113:
	;
	v690 = int32(_a_F_heap_vacuum_rel_12)
	goto L114
L114:
	;
	v691 = *(*int32)(unsafe.Add(mBase, uint32(v690)))
	if v691 == int32(0) {
		goto L109
	} else {
		goto L115
	}
L115:
	;
	v700 = v4
	v701 = v4
	v702 = v4
	goto L116
L116:
	;
	v747 = *(*int32)(unsafe.Add(mBase, uint32(v667+v701<<(uint(int32(2))%32))))
	v748 = *(*int32)(unsafe.Add(mBase, uint32(v747)+204))
	v749 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v748)+29)))
	if v749 == int32(0) {
		v769 = v700
		v770 = v702
		goto L118
	} else {
		goto L119
	}
L117:
	;
	if v770 < v769 {
		goto L123
	} else {
		goto L124
	}
L118:
	;
	v772 = v701 + int32(1)
	if v772 != v636 {
		v700 = v769
		v701 = v772
		v702 = v770
		goto L116
	} else {
		goto L122
	}
L119:
	;
	v753 = F_RelationGetNumberOfBlocksInFork(m, v747, int32(0))
	mBase = m.M
	v754 = m.ExcPending
	if v754 != 0 {
		goto L24
	} else {
		goto L120
	}
L120:
	;
	v756 = *(*int32)(unsafe.Add(mBase, _c_F_heap_vacuum_rel[21]))
	if base.Ui32(v753) < base.Ui32(v756) {
		v769 = v700
		v770 = v702
		goto L118
	} else {
		goto L121
	}
L121:
	;
	v759 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(v701+v674))) = uint8(v759)
	v769 = v749&v759 + v700
	v770 = v702 + base.B2i32(v749&int32(6) != int32(0))
	goto L118
L122:
	;
	goto L117
L123:
	;
	v775 = v769
	goto L125
L124:
	;
	v775 = v770
	goto L125
L125:
	;
	v777 = v775 - int32(1)
	if v777 <= int32(0) {
		goto L109
	} else {
		goto L126
	}
L126:
	;
	if base.Ui32(v633) < base.Ui32(v777) {
		goto L127
	} else {
		goto L128
	}
L127:
	;
	v781 = v633
	goto L129
L128:
	;
	v781 = v777
	goto L129
L129:
	;
	if int32(0) < v633 {
		goto L130
	} else {
		goto L131
	}
L130:
	;
	v784 = v781
	goto L132
L131:
	;
	v784 = v777
	goto L132
L132:
	;
	if v784 < v691 {
		goto L133
	} else {
		goto L134
	}
L133:
	;
	v786 = v784
	goto L135
L134:
	;
	v786 = v691
	goto L135
L135:
	;
	if int32(0) < v786 {
		goto L108
	} else {
		goto L136
	}
L136:
	;
	goto L109
L137:
	;
	v1264 = int32(0)
	goto L103
L138:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v843)+52)) = v672
	*(*int32)(unsafe.Add(mBase, uint32(v843)+36)) = v674
	*(*int32)(unsafe.Add(mBase, uint32(v843)+12)) = v636
	*(*int32)(unsafe.Add(mBase, uint32(v843)+8)) = v667
	*(*int32)(unsafe.Add(mBase, uint32(v843)+4)) = v642
	v852 = *(*int32)(unsafe.Add(mBase, _c_F_heap_vacuum_rel[22]))
	v853 = *(*int32)(unsafe.Add(mBase, uint32(v852)+72))
	*(*int32)(unsafe.Add(mBase, uint32(v852)+72)) = v853 + int32(1)
	goto L139
L139:
	;
	v859 = F_CreateParallelContext(m, int32(_a_F_heap_vacuum_rel_13), int32(_a_F_heap_vacuum_rel_14), v786)
	mBase = m.M
	v860 = m.ExcPending
	if v860 != 0 {
		goto L24
	} else {
		goto L140
	}
L140:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v843))) = v859
	v863 = F_mul_size(m, int32(48), v636)
	mBase = m.M
	v864 = m.ExcPending
	if v864 != 0 {
		goto L24
	} else {
		goto L141
	}
L141:
	;
	v865 = *(*int32)(unsafe.Add(mBase, uint32(v859)+36))
	v870 = F_add_size(m, v865, (v863+int32(31))&int32(-32))
	mBase = m.M
	v871 = m.ExcPending
	if v871 != 0 {
		goto L24
	} else {
		goto L142
	}
L142:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v859)+36)) = v870
	v873 = *(*int32)(unsafe.Add(mBase, uint32(v859)+40))
	v875 = F_add_size(m, v873, int32(1))
	mBase = m.M
	v876 = m.ExcPending
	if v876 != 0 {
		goto L24
	} else {
		goto L143
	}
L143:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v859)+40)) = v875
	v878 = *(*int32)(unsafe.Add(mBase, uint32(v859)+36))
	v880 = F_add_size(m, v878, int32(128))
	mBase = m.M
	v881 = m.ExcPending
	if v881 != 0 {
		goto L24
	} else {
		goto L144
	}
L144:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v859)+36)) = v880
	v883 = *(*int32)(unsafe.Add(mBase, uint32(v859)+40))
	v885 = F_add_size(m, v883, int32(1))
	mBase = m.M
	v886 = m.ExcPending
	if v886 != 0 {
		goto L24
	} else {
		goto L145
	}
L145:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v859)+40)) = v885
	v888 = *(*int32)(unsafe.Add(mBase, uint32(v859)+36))
	v890 = *(*int32)(unsafe.Add(mBase, uint32(v859)+12))
	v891 = F_mul_size(m, int32(128), v890)
	mBase = m.M
	v892 = m.ExcPending
	if v892 != 0 {
		goto L24
	} else {
		goto L146
	}
L146:
	;
	v897 = F_add_size(m, v888, (v891+int32(31))&int32(-32))
	mBase = m.M
	v898 = m.ExcPending
	if v898 != 0 {
		goto L24
	} else {
		goto L147
	}
L147:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v859)+36)) = v897
	v900 = *(*int32)(unsafe.Add(mBase, uint32(v859)+40))
	v902 = F_add_size(m, v900, int32(1))
	mBase = m.M
	v903 = m.ExcPending
	if v903 != 0 {
		goto L24
	} else {
		goto L148
	}
L148:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v859)+40)) = v902
	v905 = *(*int32)(unsafe.Add(mBase, uint32(v859)+36))
	v907 = *(*int32)(unsafe.Add(mBase, uint32(v859)+12))
	v908 = F_mul_size(m, int32(40), v907)
	mBase = m.M
	v909 = m.ExcPending
	if v909 != 0 {
		goto L24
	} else {
		goto L149
	}
L149:
	;
	v914 = F_add_size(m, v905, (v908+int32(31))&int32(-32))
	mBase = m.M
	v915 = m.ExcPending
	if v915 != 0 {
		goto L24
	} else {
		goto L150
	}
L150:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v859)+36)) = v914
	v917 = *(*int32)(unsafe.Add(mBase, uint32(v859)+40))
	v919 = F_add_size(m, v917, int32(1))
	mBase = m.M
	v920 = m.ExcPending
	if v920 != 0 {
		goto L24
	} else {
		goto L151
	}
L151:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v859)+40)) = v919
	v923 = *(*int32)(unsafe.Add(mBase, _c_F_heap_vacuum_rel[23]))
	if v923 != 0 {
		goto L152
	} else {
		goto L153
	}
L152:
	;
	v924 = *(*int32)(unsafe.Add(mBase, uint32(v859)+36))
	v925 = F_strlen(m, v923)
	mBase = m.M
	v930 = F_add_size(m, v924, v925&int32(-32)+int32(32))
	mBase = m.M
	v931 = m.ExcPending
	if v931 != 0 {
		goto L24
	} else {
		goto L155
	}
L153:
	;
	v938 = v4
	goto L154
L154:
	;
	F_InitializeParallelDSM(m, v859)
	mBase = m.M
	v940 = m.ExcPending
	if v940 != 0 {
		goto L24
	} else {
		goto L157
	}
L155:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v859)+36)) = v930
	v933 = *(*int32)(unsafe.Add(mBase, uint32(v859)+40))
	v935 = F_add_size(m, v933, int32(1))
	mBase = m.M
	v936 = m.ExcPending
	if v936 != 0 {
		goto L24
	} else {
		goto L156
	}
L156:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v859)+40)) = v935
	v938 = v925
	goto L154
L157:
	;
	v943 = *(*int32)(unsafe.Add(mBase, uint32(v859)+52))
	v944 = F_shm_toc_allocate(m, v943, v863)
	mBase = m.M
	v945 = m.ExcPending
	if v945 != 0 {
		goto L24
	} else {
		goto L159
	}
L158:
	;
	v978 = int32(1)
	if v636 <= v978 {
		goto L169
	} else {
		goto L170
	}
L159:
	;
	if v863&int32(3)|(v944&int32(3)|base.B2i32(base.Ui32(int32(1024)) < base.Ui32(v863))) == int32(0) {
		goto L160
	} else {
		goto L161
	}
L160:
	;
	if v863 == int32(0) {
		goto L158
	} else {
		goto L163
	}
L161:
	;
	goto L162
L162:
	;
	if v863 == int32(0) {
		goto L158
	} else {
		goto L168
	}
L163:
	;
	v958 = v863 + v944
	v960 = v944 + int32(4)
	if base.Ui32(v960) < base.Ui32(v958) {
		goto L164
	} else {
		goto L165
	}
L164:
	;
	v962 = v958
	goto L166
L165:
	;
	v962 = v960
	goto L166
L166:
	;
	v967 = (v944^int32(-1)+v962)&int32(-4) + int32(4)
	if v967 == int32(0) {
		goto L158
	} else {
		goto L167
	}
L167:
	;
	base.MemoryFill(m, v944, int32(0), v967)
	goto L158
L168:
	;
	base.MemoryFill(m, v944, int32(0), v863)
	goto L158
L169:
	;
	v981 = v978
	goto L171
L170:
	;
	v981 = v636
	goto L171
L171:
	;
	v982 = int32(0)
	v989 = v982
	v991 = v982
	goto L172
L172:
	;
	v1035 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v991+v674))))
	if v1035 != int32(1) {
		v1066 = v989
		goto L174
	} else {
		goto L175
	}
L173:
	;
	v1072 = *(*int32)(unsafe.Add(mBase, uint32(v859)+52))
	F_shm_toc_insert(m, v1072, int64(5), v944)
	mBase = m.M
	v1075 = m.ExcPending
	if v1075 != 0 {
		goto L24
	} else {
		goto L184
	}
L174:
	;
	v1070 = v991 + int32(1)
	if v1070 != v981 {
		v989 = v1066
		v991 = v1070
		goto L172
	} else {
		goto L183
	}
L175:
	;
	v1041 = *(*int32)(unsafe.Add(mBase, uint32(v667+v991<<(uint(int32(2))%32))))
	v1042 = *(*int32)(unsafe.Add(mBase, uint32(v1041)+204))
	v1043 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1042)+27)))
	v1044 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1042)+29)))
	if v1044&int32(1) != 0 {
		goto L176
	} else {
		goto L177
	}
L176:
	;
	v1047 = *(*int32)(unsafe.Add(mBase, uint32(v843)+40))
	*(*int32)(unsafe.Add(mBase, uint32(v843)+40)) = v1047 + int32(1)
	goto L178
L177:
	;
	goto L178
L178:
	;
	if v1044&int32(4) != 0 {
		goto L179
	} else {
		goto L180
	}
L179:
	;
	v1053 = *(*int32)(unsafe.Add(mBase, uint32(v843)+44))
	*(*int32)(unsafe.Add(mBase, uint32(v843)+44)) = v1053 + int32(1)
	goto L181
L180:
	;
	goto L181
L181:
	;
	v1057 = v989 + v1043
	if v1044&int32(2) == int32(0) {
		v1066 = v1057
		goto L174
	} else {
		goto L182
	}
L182:
	;
	v1062 = *(*int32)(unsafe.Add(mBase, uint32(v843)+48))
	*(*int32)(unsafe.Add(mBase, uint32(v843)+48)) = v1062 + int32(1)
	v1066 = v1057
	goto L174
L183:
	;
	goto L173
L184:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v843)+20)) = v944
	v1077 = *(*int32)(unsafe.Add(mBase, uint32(v859)+52))
	v1079 = F_shm_toc_allocate(m, v1077, int32(112))
	mBase = m.M
	v1080 = m.ExcPending
	if v1080 != 0 {
		goto L24
	} else {
		goto L185
	}
L185:
	;
	v1081 = int32(0)
	base.MemoryFill(m, v1079, v1081, int32(112))
	v1084 = *(*int32)(unsafe.Add(mBase, uint32(v642)+56))
	*(*int32)(unsafe.Add(mBase, uint32(v1079)+4)) = v671
	*(*int32)(unsafe.Add(mBase, uint32(v1079))) = v1084
	v1089 = *(*int32)(unsafe.Add(mBase, _c_F_heap_vacuum_rel[9]))
	if v1089 == v1081 {
		goto L187
	} else {
		goto L188
	}
L186:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v1079)+8)) = v1094
	if int32(0) < v1066 {
		goto L190
	} else {
		goto L191
	}
L187:
	;
	v1094 = int64(0)
	goto L186
L188:
	;
	goto L189
L189:
	;
	v1093 = *(*int64)(unsafe.Add(mBase, uint32(v1089)+392))
	v1094 = v1093
	goto L186
L190:
	;
	if v786 < v1066 {
		goto L193
	} else {
		goto L194
	}
L191:
	;
	v1101 = v632
	goto L192
L192:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v1079)+28)) = v1101
	v1104 = v632 << (uint(int32(10)) % 32)
	*(*int32)(unsafe.Add(mBase, uint32(v1079)+56)) = v1104
	v1106 = F_TidStoreCreateShared(m, v1104)
	mBase = m.M
	v1107 = m.ExcPending
	if v1107 != 0 {
		goto L24
	} else {
		goto L196
	}
L193:
	;
	v1099 = v786
	goto L195
L194:
	;
	v1099 = v1066
	goto L195
L195:
	;
	v1100 = base.I32_div_s(v632, v1099)
	v1101 = v1100
	goto L192
L196:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v843)+24)) = v1106
	v1109 = *(*int32)(unsafe.Add(mBase, uint32(v1106)+4))
	v1110 = *(*int32)(unsafe.Add(mBase, uint32(v1109)))
	v1111 = *(*int32)(unsafe.Add(mBase, uint32(v1110)))
	goto L197
L197:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v1079)+52)) = v1111
	v1113 = *(*int32)(unsafe.Add(mBase, uint32(v1106)+8))
	v1114 = *(*int32)(unsafe.Add(mBase, uint32(v1113)))
	v1115 = *(*int32)(unsafe.Add(mBase, uint32(v1114)+28))
	goto L198
L198:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v1079)+48)) = v1115
	if v672 == int32(0) {
		goto L200
	} else {
		goto L201
	}
L199:
	;
	v1122 = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v1079)+36)) = v1122
	*(*int32)(unsafe.Add(mBase, uint32(v1079)+40)) = v1122
	*(*int32)(unsafe.Add(mBase, uint32(v1079)+32)) = v1121
	*(*int32)(unsafe.Add(mBase, uint32(v1079)+44)) = v1122
	v1130 = *(*int32)(unsafe.Add(mBase, _c_F_heap_vacuum_rel[5]))
	v1132 = base.B2i32(v1130 == int32(4))
	*(*uint8)(unsafe.Add(mBase, uint32(v1079)+72)) = uint8(v1132)
	if v1130 == int32(4) {
		goto L203
	} else {
		goto L204
	}
L200:
	;
	v1121 = int32(0)
	goto L199
L201:
	;
	goto L202
L202:
	;
	v1120 = *(*int32)(unsafe.Add(mBase, uint32(v672)+4))
	v1121 = v1120
	goto L199
L203:
	;
	v1137 = *(*float64)(unsafe.Add(mBase, _c_F_heap_vacuum_rel[24]))
	*(*float64)(unsafe.Add(mBase, uint32(v1079)+88)) = v1137
	v1140 = *(*int32)(unsafe.Add(mBase, _c_F_heap_vacuum_rel[25]))
	*(*int32)(unsafe.Add(mBase, uint32(v1079)+96)) = v1140
	v1143 = *(*int32)(unsafe.Add(mBase, _c_F_heap_vacuum_rel[26]))
	*(*int32)(unsafe.Add(mBase, uint32(v1079)+100)) = v1143
	v1146 = *(*int32)(unsafe.Add(mBase, _c_F_heap_vacuum_rel[27]))
	*(*int32)(unsafe.Add(mBase, uint32(v1079)+104)) = v1146
	v1149 = *(*int32)(unsafe.Add(mBase, _c_F_heap_vacuum_rel[28]))
	*(*int32)(unsafe.Add(mBase, uint32(v1079)+80)) = int32(1)
	*(*int32)(unsafe.Add(mBase, uint32(v1079)+108)) = v1149
	v1153 = int32(0)
	atomic.StoreUint32((*uint32)(unsafe.Add(mBase, uint32(v1079)+84)), uint32(v1153))
	*(*int32)(unsafe.Add(mBase, _c_F_heap_vacuum_rel[29])) = v1079 + int32(80)
	v1160 = *(*int32)(unsafe.Add(mBase, uint32(v859)+44))
	F_on_dsm_detach(m, v1160, int32(629), int64(0))
	mBase = m.M
	v1164 = m.ExcPending
	if v1164 != 0 {
		goto L24
	} else {
		goto L206
	}
L204:
	;
	goto L205
L205:
	;
	v1166 = *(*int32)(unsafe.Add(mBase, uint32(v859)+52))
	F_shm_toc_insert(m, v1166, int64(1), v1079)
	mBase = m.M
	v1169 = m.ExcPending
	if v1169 != 0 {
		goto L24
	} else {
		goto L207
	}
L206:
	;
	goto L205
L207:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v843)+16)) = v1079
	v1171 = *(*int32)(unsafe.Add(mBase, uint32(v859)+52))
	v1173 = *(*int32)(unsafe.Add(mBase, uint32(v859)+12))
	v1174 = F_mul_size(m, int32(128), v1173)
	mBase = m.M
	v1175 = m.ExcPending
	if v1175 != 0 {
		goto L24
	} else {
		goto L208
	}
L208:
	;
	v1176 = F_shm_toc_allocate(m, v1171, v1174)
	mBase = m.M
	v1177 = m.ExcPending
	if v1177 != 0 {
		goto L24
	} else {
		goto L209
	}
L209:
	;
	v1178 = *(*int32)(unsafe.Add(mBase, uint32(v859)+52))
	F_shm_toc_insert(m, v1178, int64(3), v1176)
	mBase = m.M
	v1181 = m.ExcPending
	if v1181 != 0 {
		goto L24
	} else {
		goto L210
	}
L210:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v843)+28)) = v1176
	v1183 = *(*int32)(unsafe.Add(mBase, uint32(v859)+52))
	v1185 = *(*int32)(unsafe.Add(mBase, uint32(v859)+12))
	v1186 = F_mul_size(m, int32(40), v1185)
	mBase = m.M
	v1187 = m.ExcPending
	if v1187 != 0 {
		goto L24
	} else {
		goto L211
	}
L211:
	;
	v1188 = F_shm_toc_allocate(m, v1183, v1186)
	mBase = m.M
	v1189 = m.ExcPending
	if v1189 != 0 {
		goto L24
	} else {
		goto L212
	}
L212:
	;
	v1190 = *(*int32)(unsafe.Add(mBase, uint32(v859)+52))
	F_shm_toc_insert(m, v1190, int64(4), v1188)
	mBase = m.M
	v1193 = m.ExcPending
	if v1193 != 0 {
		goto L24
	} else {
		goto L213
	}
L213:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v843)+32)) = v1188
	v1196 = *(*int32)(unsafe.Add(mBase, _c_F_heap_vacuum_rel[23]))
	if v1196 != 0 {
		goto L214
	} else {
		goto L215
	}
L214:
	;
	v1197 = *(*int32)(unsafe.Add(mBase, uint32(v859)+52))
	v1199 = v938 + int32(1)
	v1200 = F_shm_toc_allocate(m, v1197, v1199)
	mBase = m.M
	v1201 = m.ExcPending
	if v1201 != 0 {
		goto L24
	} else {
		goto L217
	}
L215:
	;
	goto L216
L216:
	;
	v1264 = v843
	goto L103
L217:
	;
	if v1199 != 0 {
		goto L218
	} else {
		goto L219
	}
L218:
	;
	v1203 = *(*int32)(unsafe.Add(mBase, _c_F_heap_vacuum_rel[23]))
	base.MemoryCopy(m, v1200, v1203, v1199)
	goto L220
L219:
	;
	goto L220
L220:
	;
	v1206 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v1200+v938))) = uint8(v1206)
	v1208 = *(*int32)(unsafe.Add(mBase, uint32(v859)+52))
	F_shm_toc_insert(m, v1208, int64(2), v1200)
	mBase = m.M
	v1211 = m.ExcPending
	if v1211 != 0 {
		goto L24
	} else {
		goto L221
	}
L221:
	;
	goto L216
L222:
	;
	v1321 = *(*int32)(unsafe.Add(mBase, uint32(v1316)+16))
	*(*int32)(unsafe.Add(mBase, uint32(v231+int32(104)))) = v1321 + int32(56)
	v1325 = *(*int32)(unsafe.Add(mBase, uint32(v1316)+24))
	goto L223
L223:
	;
	v1438 = v1325
	goto L89
L224:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v1377)+8)) = int64(0)
	*(*int32)(unsafe.Add(mBase, uint32(v1377))) = v632 << (uint(int32(10)) % 32)
	*(*int32)(unsafe.Add(mBase, uint32(v231)+104)) = v1377
	v1385 = *(*int32)(unsafe.Add(mBase, uint32(v1377)))
	v1386 = F_TidStoreCreateLocal(m, v1385)
	mBase = m.M
	v1387 = m.ExcPending
	if v1387 != 0 {
		goto L24
	} else {
		goto L225
	}
L225:
	;
	v1438 = v1386
	goto L89
L226:
	;
	v1653 = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v231)+260)) = v1653
	*(*uint8)(unsafe.Add(mBase, uint32(v231)+256)) = uint8(v1653)
	*(*int64)(unsafe.Add(mBase, uint32(v231)+248)) = int64(-1)
	v1660 = v231 + int32(88)
	v1662 = v53 + int32(1068)
	v1663 = int32(1)
	v1664 = *(*int32)(unsafe.Add(mBase, uint32(v231)+12))
	v1665 = *(*int32)(unsafe.Add(mBase, uint32(v231)))
	v1669 = F_read_stream_begin_relation(m, v1663, v1664, v1665, v1653, int32(190), v231, v1663)
	mBase = m.M
	v1670 = m.ExcPending
	if v1670 != 0 {
		goto L24
	} else {
		goto L243
	}
L227:
	;
	goto L226
L228:
	;
	v1481 = *(*int32)(unsafe.Add(mBase, _c_F_heap_vacuum_rel[9]))
	if v1481 == int32(0) {
		goto L227
	} else {
		goto L229
	}
L229:
	;
	v1485 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_heap_vacuum_rel[10])))
	if v1485&int32(1) == int32(0) {
		goto L227
	} else {
		goto L230
	}
L230:
	;
	v1490 = int32(_a_F_heap_vacuum_rel_1)
	v1492 = *(*int32)(unsafe.Add(mBase, _c_F_heap_vacuum_rel[11]))
	v1493 = int32(1)
	*(*int32)(unsafe.Add(mBase, _c_F_heap_vacuum_rel[11])) = v1492 + v1493
	v1496 = *(*int32)(unsafe.Add(mBase, uint32(v1481)))
	*(*int32)(unsafe.Add(mBase, uint32(v1481))) = v1496 + v1493
	v1500 = int32(0)
	v1503 = base.AtomicRmwOr32(m, v1500, int32(_a_F_heap_vacuum_rel_2), v1500)
	goto L232
L231:
	;
	v1630 = int32(0)
	v1633 = base.AtomicRmwOr32(m, v1630, int32(_a_F_heap_vacuum_rel_2), v1630)
	v1634 = *(*int32)(unsafe.Add(mBase, uint32(v1481)))
	v1635 = int32(1)
	*(*int32)(unsafe.Add(mBase, uint32(v1481))) = v1634 + v1635
	v1638 = int32(_a_F_heap_vacuum_rel_1)
	v1640 = *(*int32)(unsafe.Add(mBase, _c_F_heap_vacuum_rel[11]))
	*(*int32)(unsafe.Add(mBase, _c_F_heap_vacuum_rel[11])) = v1640 - v1635
	goto L227
L232:
	;
	goto L234
L234:
	;
	goto L235
L235:
	;
	v1595 = int32(0)
	v1598 = v1452
	goto L240
L240:
	;
	v1607 = *(*int32)(unsafe.Add(mBase, uint32(v53+int32(1024)+v1598<<(uint(int32(2))%32))))
	v1608 = int32(3)
	v1614 = *(*int64)(unsafe.Add(mBase, uint32(v53+int32(992)+v1598<<(uint(v1608)%32))))
	*(*int64)(unsafe.Add(mBase, uint32(v1481+int32(232)+v1607<<(uint(v1608)%32)))) = v1614
	v1616 = int32(1)
	v1619 = v1595 + v1616
	if v1619 != int32(3) {
		v1595 = v1619
		v1598 = v1598 + v1616
		goto L240
	} else {
		goto L242
	}
L241:
	;
	goto L231
L242:
	;
	goto L241
L243:
	;
	v1677 = int32(0)
	v1681 = v4
	v1703 = v4
	goto L244
L244:
	;
	v1722 = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v53)+988)) = v1722
	F_vacuum_delay_point(m, v1722)
	mBase = m.M
	v1726 = m.ExcPending
	if v1726 != 0 {
		goto L24
	} else {
		goto L246
	}
L245:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v231)+84)) = int32(-1)
	v3390 = *(*int32)(unsafe.Add(mBase, uint32(v53)+1036))
	if v3390 != 0 {
		goto L513
	} else {
		goto L514
	}
L246:
	;
	v1727 = *(*int32)(unsafe.Add(mBase, uint32(v1448)))
	v1728 = int32(0)
	if base.B2i32(v1727 == v1728)|v1727&int32(_a_F_heap_vacuum_rel_15) == v1728 {
		goto L247
	} else {
		goto L248
	}
L247:
	;
	v1735 = F_lazy_check_wraparound_failsafe(m, v231)
	mBase = m.M
	v1736 = m.ExcPending
	if v1736 != 0 {
		goto L24
	} else {
		goto L250
	}
L248:
	;
	goto L249
L249:
	;
	v1737 = *(*int32)(unsafe.Add(mBase, uint32(v231)+104))
	v1738 = *(*int64)(unsafe.Add(mBase, uint32(v1737)+8))
	if v1738 <= int64(0) {
		v1807 = v1681
		goto L251
	} else {
		goto L252
	}
L250:
	;
	goto L249
L251:
	;
	v1809 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_heap_vacuum_rel[14])))
	if v1809 == int32(1) {
		goto L265
	} else {
		goto L266
	}
L252:
	;
	v1741 = *(*int32)(unsafe.Add(mBase, uint32(v231)+100))
	v1742 = F_TidStoreMemoryUsage(m, v1741)
	mBase = m.M
	v1743 = m.ExcPending
	if v1743 != 0 {
		goto L24
	} else {
		goto L253
	}
L253:
	;
	v1744 = *(*int32)(unsafe.Add(mBase, uint32(v231)+104))
	v1745 = *(*int32)(unsafe.Add(mBase, uint32(v1744)))
	if base.Ui32(v1742) <= base.Ui32(v1745) {
		v1807 = v1681
		goto L251
	} else {
		goto L254
	}
L254:
	;
	v1747 = *(*int32)(unsafe.Add(mBase, uint32(v53)+1036))
	if v1747 != 0 {
		goto L255
	} else {
		goto L256
	}
L255:
	;
	F_ReleaseBuffer(m, v1747)
	mBase = m.M
	v1749 = m.ExcPending
	if v1749 != 0 {
		goto L24
	} else {
		goto L258
	}
L256:
	;
	goto L257
L257:
	;
	v1752 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v231)+22)) = uint8(v1752)
	F_lazy_vacuum(m, v231)
	mBase = m.M
	v1755 = m.ExcPending
	if v1755 != 0 {
		goto L24
	} else {
		goto L259
	}
L258:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v53)+1036)) = int32(0)
	goto L257
L259:
	;
	v1756 = *(*int32)(unsafe.Add(mBase, uint32(v231)))
	F_FreeSpaceMapVacuumRange(m, v1756, v1681, v1677+int32(1))
	mBase = m.M
	v1760 = m.ExcPending
	if v1760 != 0 {
		goto L24
	} else {
		goto L260
	}
L260:
	;
	v1765 = *(*int32)(unsafe.Add(mBase, _c_F_heap_vacuum_rel[9]))
	if v1765 == int32(0) {
		goto L262
	} else {
		goto L263
	}
L261:
	;
	v1807 = v1677
	goto L251
L262:
	;
	goto L261
L263:
	;
	v1769 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_heap_vacuum_rel[10])))
	if v1769&int32(1) == int32(0) {
		goto L262
	} else {
		goto L264
	}
L264:
	;
	v1774 = int32(_a_F_heap_vacuum_rel_1)
	v1776 = *(*int32)(unsafe.Add(mBase, _c_F_heap_vacuum_rel[11]))
	v1777 = int32(1)
	*(*int32)(unsafe.Add(mBase, _c_F_heap_vacuum_rel[11])) = v1776 + v1777
	v1780 = *(*int32)(unsafe.Add(mBase, uint32(v1765)))
	*(*int32)(unsafe.Add(mBase, uint32(v1765))) = v1780 + v1777
	v1784 = int32(0)
	v1786 = int32(_a_F_heap_vacuum_rel_2)
	v1787 = base.AtomicRmwOr32(m, v1784, v1786, v1784)
	*(*int64)(unsafe.Add(mBase, uint32(v1765+v1784)+232)) = int64(1)
	v1795 = base.AtomicRmwOr32(m, v1784, v1786, v1784)
	v1796 = *(*int32)(unsafe.Add(mBase, uint32(v1765)))
	*(*int32)(unsafe.Add(mBase, uint32(v1765))) = v1796 + v1777
	v1802 = *(*int32)(unsafe.Add(mBase, _c_F_heap_vacuum_rel[11]))
	*(*int32)(unsafe.Add(mBase, _c_F_heap_vacuum_rel[11])) = v1802 - v1777
	goto L262
L265:
	;
	if v1703 == int32(0) {
		goto L268
	} else {
		goto L269
	}
L266:
	;
	v2010 = v1703
	goto L267
L267:
	;
	v2031 = F_read_stream_next_buffer(m, v1669, v53+int32(988))
	mBase = m.M
	v2032 = m.ExcPending
	if v2032 != 0 {
		goto L24
	} else {
		goto L277
	}
L268:
	;
	v1814 = int32(0)
	v1815 = int32(*(*int16)(unsafe.Add(mBase, uint32(v1669))))
	if v1814 < v1815 {
		goto L271
	} else {
		goto L272
	}
L269:
	;
	goto L270
L270:
	;
	v2010 = int32(1)
	goto L267
L271:
	;
	v1823 = v1814
	goto L274
L272:
	;
	goto L273
L273:
	;
	goto L270
L274:
	;
	v1868 = *(*int32)(unsafe.Add(mBase, uint32(v1669)+76))
	*(*int32)(unsafe.Add(mBase, uint32(v1868+v1823*int32(84))+20)) = int32(0)
	v1875 = v1823 + int32(1)
	v1876 = int32(*(*int16)(unsafe.Add(mBase, uint32(v1669))))
	if v1875 < v1876 {
		v1823 = v1875
		goto L274
	} else {
		goto L276
	}
L275:
	;
	goto L273
L276:
	;
	goto L275
L277:
	;
	if v2031 != 0 {
		goto L278
	} else {
		goto L279
	}
L278:
	;
	v2033 = *(*int32)(unsafe.Add(mBase, uint32(v53)+988))
	v2034 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v2033))))
	F_CheckBufferIsPinnedOnce(m, v2031)
	mBase = m.M
	v2036 = m.ExcPending
	if v2036 != 0 {
		goto L24
	} else {
		goto L281
	}
L279:
	;
	goto L280
L280:
	;
	goto L245
L281:
	;
	if v2031 < int32(0) {
		goto L283
	} else {
		goto L284
	}
L282:
	;
	if v2031 < int32(0) {
		goto L287
	} else {
		goto L288
	}
L283:
	;
	v2040 = *(*int32)(unsafe.Add(mBase, _c_F_heap_vacuum_rel[30]))
	v2046 = *(*int32)(unsafe.Add(mBase, uint32(v2040+(v2031^int32(-1))<<(uint(int32(2))%32))))
	v2054 = v2046
	goto L282
L284:
	;
	goto L285
L285:
	;
	v2048 = *(*int32)(unsafe.Add(mBase, _c_F_heap_vacuum_rel[31]))
	v2054 = v2048 + v2031<<(uint(int32(13))%32) + int32(-8192)
	goto L282
L286:
	;
	v2074 = *(*int32)(unsafe.Add(mBase, uint32(v1448)))
	v2075 = int32(1)
	*(*int32)(unsafe.Add(mBase, uint32(v1448))) = v2074 + v2075
	if v2034&v2075 != 0 {
		goto L290
	} else {
		goto L291
	}
L287:
	;
	v2058 = *(*int32)(unsafe.Add(mBase, _c_F_heap_vacuum_rel[32]))
	v2064 = *(*int32)(unsafe.Add(mBase, uint32(v2058+(v2031^int32(-1))*int32(56))+16))
	v2073 = v2064
	goto L286
L288:
	;
	goto L289
L289:
	;
	v2066 = *(*int32)(unsafe.Add(mBase, _c_F_heap_vacuum_rel[33]))
	v2067 = int32(56)
	v2072 = *(*int32)(unsafe.Add(mBase, uint32(v2066+v2031*v2067-v2067)+16))
	v2073 = v2072
	goto L286
L290:
	;
	v2080 = *(*int32)(unsafe.Add(mBase, uint32(v231)+116))
	*(*int32)(unsafe.Add(mBase, uint32(v231)+116)) = v2080 + int32(1)
	goto L292
L291:
	;
	goto L292
L292:
	;
	v2088 = *(*int32)(unsafe.Add(mBase, _c_F_heap_vacuum_rel[9]))
	if v2088 == int32(0) {
		goto L294
	} else {
		goto L295
	}
L293:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v231)+92)) = int32(1)
	v2131 = int32(0)
	*(*uint16)(unsafe.Add(mBase, uint32(v231)+88)) = uint16(v2131)
	*(*int32)(unsafe.Add(mBase, uint32(v231)+84)) = v2073
	v2134 = *(*int32)(unsafe.Add(mBase, uint32(v231)))
	F_visibilitymap_pin(m, v2134, v2073, v53+int32(1036))
	mBase = m.M
	v2138 = m.ExcPending
	if v2138 != 0 {
		goto L24
	} else {
		goto L297
	}
L294:
	;
	goto L293
L295:
	;
	v2092 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_heap_vacuum_rel[10])))
	if v2092&int32(1) == int32(0) {
		goto L294
	} else {
		goto L296
	}
L296:
	;
	v2097 = int32(_a_F_heap_vacuum_rel_1)
	v2099 = *(*int32)(unsafe.Add(mBase, _c_F_heap_vacuum_rel[11]))
	v2100 = int32(1)
	*(*int32)(unsafe.Add(mBase, _c_F_heap_vacuum_rel[11])) = v2099 + v2100
	v2103 = *(*int32)(unsafe.Add(mBase, uint32(v2088)))
	*(*int32)(unsafe.Add(mBase, uint32(v2088))) = v2103 + v2100
	v2107 = int32(0)
	v2109 = int32(_a_F_heap_vacuum_rel_2)
	v2110 = base.AtomicRmwOr32(m, v2107, v2109, v2107)
	*(*int64)(unsafe.Add(mBase, uint32(v2088+int32(16))+232)) = base.I64_extend_i32_u(v2073)
	v2118 = base.AtomicRmwOr32(m, v2107, v2109, v2107)
	v2119 = *(*int32)(unsafe.Add(mBase, uint32(v2088)))
	*(*int32)(unsafe.Add(mBase, uint32(v2088))) = v2119 + v2100
	v2125 = *(*int32)(unsafe.Add(mBase, _c_F_heap_vacuum_rel[11]))
	*(*int32)(unsafe.Add(mBase, _c_F_heap_vacuum_rel[11])) = v2125 - v2100
	goto L294
L297:
	;
	v2139 = F_ConditionalLockBufferForCleanup(m, v2031)
	mBase = m.M
	v2140 = m.ExcPending
	if v2140 != 0 {
		goto L24
	} else {
		goto L298
	}
L298:
	;
	if v2139 == int32(0) {
		goto L299
	} else {
		goto L300
	}
L299:
	;
	F_LockBufferInternal(m, v2031, int32(1))
	mBase = m.M
	v2145 = m.ExcPending
	if v2145 != 0 {
		goto L24
	} else {
		goto L302
	}
L300:
	;
	goto L301
L301:
	;
	v2146 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v2054)+14)))
	if v2146 == int32(0) {
		goto L308
	} else {
		goto L309
	}
L302:
	;
	goto L301
L303:
	;
	v3294 = *(*int32)(unsafe.Add(mBase, uint32(v267)))
	if v3294 != 0 {
		goto L484
	} else {
		goto L485
	}
L304:
	;
	v2890 = *(*int32)(unsafe.Add(mBase, uint32(v231)))
	*(*int64)(unsafe.Add(mBase, uint32(v53)+604)) = int64(42949672961)
	*(*int32)(unsafe.Add(mBase, uint32(v53)+600)) = v2846
	*(*int32)(unsafe.Add(mBase, uint32(v53)+596)) = v2031
	*(*int32)(unsafe.Add(mBase, uint32(v53)+592)) = v2890
	v2896 = *(*int32)(unsafe.Add(mBase, uint32(v231)+52))
	*(*int32)(unsafe.Add(mBase, uint32(v53)+616)) = v447
	*(*int32)(unsafe.Add(mBase, uint32(v53)+612)) = v2896
	v2900 = *(*int32)(unsafe.Add(mBase, uint32(v231)+8))
	if v2900 == int32(0) {
		goto L420
	} else {
		goto L421
	}
L305:
	;
	v2318 = *(*int32)(unsafe.Add(mBase, uint32(v1442)))
	*(*int32)(unsafe.Add(mBase, uint32(v53)+1688)) = v2318
	v2320 = *(*int32)(unsafe.Add(mBase, uint32(v1440)))
	*(*int32)(unsafe.Add(mBase, uint32(v53)+1680)) = v2320
	v2327 = int32(base.Ui32(v2317+int32(_a_F_heap_vacuum_rel_16))>>(uint(int32(2))%32)) & int32(_a_F_heap_vacuum_rel_17)
	if v2327 != 0 {
		goto L362
	} else {
		goto L363
	}
L306:
	;
	if v2139 != 0 {
		v2846 = v2155
		goto L304
	} else {
		goto L357
	}
L307:
	;
	v2314 = *(*int32)(unsafe.Add(mBase, uint32(v231)))
	F_RecordPageWithFreeSpace(m, v2314, v2073, v2311)
	mBase = m.M
	v2316 = m.ExcPending
	if v2316 != 0 {
		goto L24
	} else {
		goto L356
	}
L308:
	;
	F_UnlockReleaseBuffer(m, v2031)
	mBase = m.M
	v2150 = m.ExcPending
	if v2150 != 0 {
		goto L24
	} else {
		goto L311
	}
L309:
	;
	goto L310
L310:
	;
	v2155 = *(*int32)(unsafe.Add(mBase, uint32(v53)+1036))
	v2156 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v2054)+12)))
	if base.Ui32(int32(24)) < base.Ui32(v2156) {
		goto L306
	} else {
		goto L314
	}
L311:
	;
	v2151 = *(*int32)(unsafe.Add(mBase, uint32(v231)))
	v2152 = F_GetRecordedFreeSpace(m, v2151, v2073)
	mBase = m.M
	v2153 = m.ExcPending
	if v2153 != 0 {
		goto L24
	} else {
		goto L312
	}
L312:
	;
	if v2152 != 0 {
		v1677 = v2073
		v1681 = v1807
		v1703 = v2010
		goto L244
	} else {
		goto L313
	}
L313:
	;
	v2311 = int32(_a_F_heap_vacuum_rel_18)
	goto L307
L314:
	;
	if v2139 == int32(0) {
		goto L315
	} else {
		goto L316
	}
L315:
	;
	F_UnlockBuffer(m, v2031)
	mBase = m.M
	v2162 = m.ExcPending
	if v2162 != 0 {
		goto L24
	} else {
		goto L318
	}
L316:
	;
	goto L317
L317:
	;
	v2170 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v2054)+10)))
	if v2170&int32(4) == int32(0) {
		goto L321
	} else {
		goto L322
	}
L318:
	;
	F_LockBufferInternal(m, v2031, int32(3))
	mBase = m.M
	v2165 = m.ExcPending
	if v2165 != 0 {
		goto L24
	} else {
		goto L319
	}
L319:
	;
	v2166 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v2054)+12)))
	if base.Ui32(int32(24)) < base.Ui32(v2166) {
		v2317 = v2166
		goto L305
	} else {
		goto L320
	}
L320:
	;
	goto L317
L321:
	;
	F_LockBufferInternal(m, v2155, int32(3))
	mBase = m.M
	v2177 = m.ExcPending
	if v2177 != 0 {
		goto L24
	} else {
		goto L324
	}
L322:
	;
	goto L323
L323:
	;
	v2245 = int32(4)
	v2246 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v2054)+14)))
	v2247 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v2054)+12)))
	v2248 = v2246 - v2247
	if v2248 <= v2245 {
		goto L337
	} else {
		goto L338
	}
L324:
	;
	v2178 = int32(_a_F_heap_vacuum_rel_1)
	v2180 = *(*int32)(unsafe.Add(mBase, _c_F_heap_vacuum_rel[11]))
	*(*int32)(unsafe.Add(mBase, _c_F_heap_vacuum_rel[11])) = v2180 + int32(1)
	F_MarkBufferDirty(m, v2031)
	mBase = m.M
	v2185 = m.ExcPending
	if v2185 != 0 {
		goto L24
	} else {
		goto L325
	}
L325:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v2054)+20)) = int32(0)
	v2188 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v2054)+10)))
	v2190 = v2188 | int32(4)
	*(*uint16)(unsafe.Add(mBase, uint32(v2054)+10)) = uint16(v2190)
	v2192 = *(*int32)(unsafe.Add(mBase, uint32(v231)))
	v2193 = *(*int32)(unsafe.Add(mBase, uint32(v2192)+8))
	*(*int32)(unsafe.Add(mBase, uint32(v53)+536)) = v2193
	v2195 = *(*int64)(unsafe.Add(mBase, uint32(v2192)))
	*(*int64)(unsafe.Add(mBase, uint32(v53)+528)) = v2195
	v2198 = F_visibilitymap_set(m, v2073, v2155, int32(3))
	mBase = m.M
	v2199 = m.ExcPending
	if v2199 != 0 {
		goto L24
	} else {
		goto L326
	}
L326:
	;
	v2200 = *(*int32)(unsafe.Add(mBase, uint32(v231)))
	v2201 = *(*int32)(unsafe.Add(mBase, uint32(v2200)+48))
	v2202 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v2201)+118)))
	if v2202 != int32(112) {
		goto L327
	} else {
		goto L328
	}
L327:
	;
	v2225 = int32(_a_F_heap_vacuum_rel_1)
	v2227 = *(*int32)(unsafe.Add(mBase, _c_F_heap_vacuum_rel[11]))
	*(*int32)(unsafe.Add(mBase, _c_F_heap_vacuum_rel[11])) = v2227 - int32(1)
	F_UnlockBuffer(m, v2155)
	mBase = m.M
	v2232 = m.ExcPending
	if v2232 != 0 {
		goto L24
	} else {
		goto L335
	}
L328:
	;
	v2206 = *(*int32)(unsafe.Add(mBase, _c_F_heap_vacuum_rel[34]))
	if v2206 <= int32(0) {
		goto L329
	} else {
		goto L330
	}
L329:
	;
	v2209 = *(*int32)(unsafe.Add(mBase, uint32(v2200)+32))
	if v2209 != 0 {
		goto L327
	} else {
		goto L332
	}
L330:
	;
	goto L331
L331:
	;
	v2212 = int32(0)
	F_log_heap_prune_and_freeze(m, v2200, v2031, v2155, int32(3), v2212, v2212, int32(1), v2212, v2212, v2212, v2212, v2212, v2212, v2212, v2212)
	mBase = m.M
	v2224 = m.ExcPending
	if v2224 != 0 {
		goto L24
	} else {
		goto L334
	}
L332:
	;
	v2210 = *(*int32)(unsafe.Add(mBase, uint32(v2200)+40))
	if v2210 != 0 {
		goto L327
	} else {
		goto L333
	}
L333:
	;
	goto L331
L334:
	;
	goto L327
L335:
	;
	v2233 = *(*int32)(unsafe.Add(mBase, uint32(v231)+128))
	v2234 = int32(1)
	*(*int32)(unsafe.Add(mBase, uint32(v231)+128)) = v2233 + v2234
	v2237 = *(*int32)(unsafe.Add(mBase, uint32(v231)+132))
	*(*int32)(unsafe.Add(mBase, uint32(v231)+132)) = v2237 + v2234
	goto L323
L336:
	;
	F_UnlockReleaseBuffer(m, v2031)
	mBase = m.M
	v2310 = m.ExcPending
	if v2310 != 0 {
		goto L24
	} else {
		goto L355
	}
L337:
	;
	v2251 = v2245
	goto L339
L338:
	;
	v2251 = v2248
	goto L339
L339:
	;
	v2253 = v2251 - int32(4)
	if v2253 == int32(0) {
		goto L340
	} else {
		goto L341
	}
L340:
	;
	v2308 = int32(0)
	goto L336
L341:
	;
	goto L342
L342:
	;
	if base.Ui32(int32(25)) <= base.Ui32(v2247) {
		goto L344
	} else {
		goto L345
	}
L343:
	;
	v2308 = v2253
	goto L336
L344:
	;
	v2264 = int32(base.Ui32(v2247+int32(_a_F_heap_vacuum_rel_16)) >> (uint(int32(2)) % 32))
	goto L346
L345:
	;
	v2264 = int32(0)
	goto L346
L346:
	;
	if base.Ui32(v2264&int32(_a_F_heap_vacuum_rel_17)) < base.Ui32(int32(291)) {
		goto L343
	} else {
		goto L347
	}
L347:
	;
	v2269 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v2054)+10)))
	if v2269&int32(1) == int32(0) {
		goto L348
	} else {
		goto L349
	}
L348:
	;
	v2308 = int32(0)
	goto L336
L349:
	;
	goto L350
L350:
	;
	v2278 = int32(1)
	goto L351
L351:
	;
	v2287 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v2054+int32(20)+v2278&int32(_a_F_heap_vacuum_rel_17)<<(uint(int32(2))%32))+1)))
	if v2287&int32(384) == int32(0) {
		goto L343
	} else {
		goto L353
	}
L352:
	;
	v2308 = int32(0)
	goto L336
L353:
	;
	v2293 = v2278 + int32(1)
	v2294 = int32(_a_F_heap_vacuum_rel_17)
	if base.Ui32(v2293&v2294) <= base.Ui32(v2264&v2294) {
		v2278 = v2293
		goto L351
	} else {
		goto L354
	}
L354:
	;
	goto L352
L355:
	;
	v2311 = v2308
	goto L307
L356:
	;
	v1677 = v2073
	v1681 = v1807
	v1703 = v2010
	goto L244
L357:
	;
	v2317 = v2156
	goto L305
L358:
	;
	v2833 = int32(0)
	*(*uint16)(unsafe.Add(mBase, uint32(v1660))) = uint16(v2833)
	F_UnlockBuffer(m, v2031)
	mBase = m.M
	v2836 = m.ExcPending
	if v2836 != 0 {
		goto L24
	} else {
		goto L418
	}
L359:
	;
	v2810 = *(*int64)(unsafe.Add(mBase, uint32(v231)+224))
	*(*int64)(unsafe.Add(mBase, uint32(v231)+224)) = v2810 + v2799
	v2813 = *(*int64)(unsafe.Add(mBase, uint32(v231)+232))
	*(*int64)(unsafe.Add(mBase, uint32(v231)+232)) = v2813 + v2803
	v2816 = *(*int64)(unsafe.Add(mBase, uint32(v231)+240))
	*(*int64)(unsafe.Add(mBase, uint32(v231)+240)) = v2816 + base.I64_extend_i32_s(v2770)
	if int32(0) < v2770 {
		goto L412
	} else {
		goto L413
	}
L360:
	;
	if v2461 <= int32(0) {
		goto L390
	} else {
		goto L391
	}
L361:
	;
	v2537 = int32(0)
	v2538 = base.B2i32(v2537 < v2498)
	if v2537 < v2498 {
		goto L387
	} else {
		goto L388
	}
L362:
	;
	v2329 = int32(base.Ui32(v2073) >> (uint(int32(16)) % 32))
	v2332 = int32(0)
	v2340 = int32(1)
	v2344 = v2332
	v2348 = v2332
	v2349 = v2332
	v2355 = v2332
	v2357 = v2332
	goto L365
L363:
	;
	goto L364
L364:
	;
	v2478 = int32(0)
	*(*uint16)(unsafe.Add(mBase, uint32(v1660))) = uint16(v2478)
	v2481 = int64(0)
	v2486 = *(*int32)(unsafe.Add(mBase, uint32(v267)))
	if v2486 != 0 {
		v2762 = v2478
		v2766 = v2478
		v2770 = v2478
		v2799 = v2481
		v2803 = v2481
		goto L359
	} else {
		goto L386
	}
L365:
	;
	*(*uint16)(unsafe.Add(mBase, uint32(v1660))) = uint16(v2340)
	v2393 = v2054 + int32(20) + v2340&int32(_a_F_heap_vacuum_rel_17)<<(uint(int32(2))%32)
	v2394 = *(*int32)(unsafe.Add(mBase, uint32(v2393)))
	switch int32(base.Ui32(v2394)>>(uint(int32(15))%32))&int32(3) - int32(1) {
	case 0:
		goto L369
	case 1:
		goto L368
	case 2:
		goto L370
	default:
		v2459 = v2344
		v2460 = v2348
		v2461 = v2349
		v2462 = v2355
		v2463 = v2357
		goto L367
	}
L366:
	;
	v2469 = *(*int32)(unsafe.Add(mBase, uint32(v53)+1680))
	v2470 = *(*int32)(unsafe.Add(mBase, uint32(v53)+1688))
	v2471 = int32(0)
	*(*uint16)(unsafe.Add(mBase, uint32(v1660))) = uint16(v2471)
	*(*int32)(unsafe.Add(mBase, uint32(v1442))) = v2470
	*(*int32)(unsafe.Add(mBase, uint32(v1440))) = v2469
	v2475 = base.I64_extend_i32_s(v2462)
	v2476 = base.I64_extend_i32_s(v2463)
	v2477 = *(*int32)(unsafe.Add(mBase, uint32(v267)))
	if v2477 != 0 {
		goto L360
	} else {
		goto L385
	}
L367:
	;
	v2465 = v2340 + int32(1)
	if base.Ui32(v2465&int32(_a_F_heap_vacuum_rel_17)) <= base.Ui32(v2327) {
		v2340 = v2465
		v2344 = v2459
		v2348 = v2460
		v2349 = v2461
		v2355 = v2462
		v2357 = v2463
		goto L365
	} else {
		goto L384
	}
L368:
	;
	v2459 = int32(1)
	v2460 = v2348
	v2461 = v2349
	v2462 = v2355
	v2463 = v2357
	goto L367
L369:
	;
	v2416 = F_heap_tuple_should_freeze(m, v2054+v2394&int32(_a_F_heap_vacuum_rel_19), v447, v53+int32(1688), v53+int32(1680))
	mBase = m.M
	v2417 = m.ExcPending
	if v2417 != 0 {
		goto L24
	} else {
		goto L371
	}
L370:
	;
	v2403 = int32(1)
	*(*uint16)(unsafe.Add(mBase, uint32(v53+int32(1040)+v2349<<(uint(v2403)%32)))) = uint16(v2340)
	v2459 = v2344
	v2460 = v2348
	v2461 = v2349 + v2403
	v2462 = v2355
	v2463 = v2357
	goto L367
L371:
	;
	if v2416 != 0 {
		goto L372
	} else {
		goto L373
	}
L372:
	;
	v2418 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v231)+20)))
	if v2418 != 0 {
		goto L358
	} else {
		goto L375
	}
L373:
	;
	goto L374
L374:
	;
	*(*uint16)(unsafe.Add(mBase, uint32(v53)+600)) = uint16(v2340)
	*(*uint16)(unsafe.Add(mBase, uint32(v53)+598)) = uint16(v2073)
	*(*uint16)(unsafe.Add(mBase, uint32(v53)+596)) = uint16(v2329)
	v2422 = *(*int32)(unsafe.Add(mBase, uint32(v2393)))
	*(*int32)(unsafe.Add(mBase, uint32(v53)+592)) = int32(base.Ui32(v2422) >> (uint(int32(17)) % 32))
	*(*int32)(unsafe.Add(mBase, uint32(v53)+608)) = v2054 + v2422&int32(_a_F_heap_vacuum_rel_19)
	v2430 = *(*int32)(unsafe.Add(mBase, uint32(v231)))
	v2431 = *(*int32)(unsafe.Add(mBase, uint32(v2430)+56))
	*(*int32)(unsafe.Add(mBase, uint32(v53)+604)) = v2431
	v2433 = int32(1)
	v2436 = *(*int32)(unsafe.Add(mBase, uint32(v231)+36))
	v2437 = F_HeapTupleSatisfiesVacuum(m, v53+int32(592), v2436, v2031)
	mBase = m.M
	v2438 = m.ExcPending
	if v2438 != 0 {
		goto L24
	} else {
		goto L380
	}
L375:
	;
	goto L374
L376:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v2448 = m.ExcPending
	if v2448 != 0 {
		goto L24
	} else {
		goto L381
	}
L377:
	;
	v2459 = v2433
	v2460 = v2348
	v2461 = v2349
	v2462 = v2355 + int32(1)
	v2463 = v2357
	goto L367
L378:
	;
	v2459 = v2433
	v2460 = v2348 + int32(1)
	v2461 = v2349
	v2462 = v2355
	v2463 = v2357
	goto L367
L379:
	;
	v2459 = v2433
	v2460 = v2348
	v2461 = v2349
	v2462 = v2355
	v2463 = v2357 + int32(1)
	goto L367
L380:
	;
	switch v2437 {
	case 0:
		goto L378
	case 1, 4:
		goto L379
	case 2:
		goto L377
	case 3:
		v2459 = v2433
		v2460 = v2348
		v2461 = v2349
		v2462 = v2355
		v2463 = v2357
		goto L367
	default:
		goto L376
	}
L381:
	;
	F_errmsg_internal(m, int32(_a_F_heap_vacuum_rel_20), int32(0))
	mBase = m.M
	v2452 = m.ExcPending
	if v2452 != 0 {
		goto L24
	} else {
		goto L382
	}
L382:
	;
	F_errfinish(m, int32(_a_F_heap_vacuum_rel_7), int32(2302), int32(_a_F_heap_vacuum_rel_21))
	mBase = m.M
	v2457 = m.ExcPending
	if v2457 != 0 {
		goto L24
	} else {
		goto L383
	}
L383:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L384:
	;
	goto L366
L385:
	;
	v2489 = v2459
	v2497 = v2460
	v2498 = v2461
	v2526 = v2476
	v2530 = v2475
	goto L361
L386:
	;
	v2489 = v2478
	v2497 = v2478
	v2498 = v2478
	v2526 = v2481
	v2530 = v2481
	goto L361
L387:
	;
	v2541 = v2498
	goto L389
L388:
	;
	v2541 = v2537
	goto L389
L389:
	;
	v2762 = v2538
	v2766 = v2489 | v2538
	v2770 = v2541 + v2497
	v2799 = v2526
	v2803 = v2530
	goto L359
L390:
	;
	v2762 = int32(0)
	v2766 = v2459
	v2770 = v2460
	v2799 = v2476
	v2803 = v2475
	goto L359
L391:
	;
	goto L392
L392:
	;
	v2546 = int32(1)
	v2547 = *(*int32)(unsafe.Add(mBase, uint32(v231)+140))
	*(*int32)(unsafe.Add(mBase, uint32(v231)+140)) = v2547 + v2546
	*(*int64)(unsafe.Add(mBase, uint32(v53)+1664)) = int64(25769803783)
	v2553 = *(*int32)(unsafe.Add(mBase, uint32(v231)+100))
	F_TidStoreSetBlockOffsets(m, v2553, v2073, v53+int32(1040), v2461)
	mBase = m.M
	v2557 = m.ExcPending
	if v2557 != 0 {
		goto L24
	} else {
		goto L393
	}
L393:
	;
	v2558 = *(*int32)(unsafe.Add(mBase, uint32(v231)+104))
	v2559 = base.I64_extend_i32_u(v2461)
	v2560 = *(*int64)(unsafe.Add(mBase, uint32(v2558)+8))
	*(*int64)(unsafe.Add(mBase, uint32(v2558)+8)) = v2559 + v2560
	v2563 = *(*int32)(unsafe.Add(mBase, uint32(v231)+104))
	v2564 = *(*int64)(unsafe.Add(mBase, uint32(v2563)+8))
	*(*int64)(unsafe.Add(mBase, uint32(v53)+592)) = v2564
	v2566 = *(*int32)(unsafe.Add(mBase, uint32(v231)+100))
	v2567 = F_TidStoreMemoryUsage(m, v2566)
	mBase = m.M
	v2568 = m.ExcPending
	if v2568 != 0 {
		goto L24
	} else {
		goto L394
	}
L394:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v53)+600)) = base.I64_extend_i32_u(v2567)
	goto L397
L395:
	;
	v2757 = *(*int64)(unsafe.Add(mBase, uint32(v231)+216))
	*(*int64)(unsafe.Add(mBase, uint32(v231)+216)) = v2757 + v2559
	v2762 = v2546
	v2766 = v2459
	v2770 = v2460
	v2799 = v2476
	v2803 = v2475
	goto L359
L396:
	;
	goto L395
L397:
	;
	v2585 = *(*int32)(unsafe.Add(mBase, _c_F_heap_vacuum_rel[9]))
	if v2585 == int32(0) {
		goto L396
	} else {
		goto L398
	}
L398:
	;
	v2589 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_heap_vacuum_rel[10])))
	if v2589&int32(1) == int32(0) {
		goto L396
	} else {
		goto L399
	}
L399:
	;
	v2594 = int32(_a_F_heap_vacuum_rel_1)
	v2596 = *(*int32)(unsafe.Add(mBase, _c_F_heap_vacuum_rel[11]))
	v2597 = int32(1)
	*(*int32)(unsafe.Add(mBase, _c_F_heap_vacuum_rel[11])) = v2596 + v2597
	v2600 = *(*int32)(unsafe.Add(mBase, uint32(v2585)))
	*(*int32)(unsafe.Add(mBase, uint32(v2585))) = v2600 + v2597
	v2604 = int32(0)
	v2607 = base.AtomicRmwOr32(m, v2604, int32(_a_F_heap_vacuum_rel_2), v2604)
	goto L401
L400:
	;
	v2734 = int32(0)
	v2737 = base.AtomicRmwOr32(m, v2734, int32(_a_F_heap_vacuum_rel_2), v2734)
	v2738 = *(*int32)(unsafe.Add(mBase, uint32(v2585)))
	v2739 = int32(1)
	*(*int32)(unsafe.Add(mBase, uint32(v2585))) = v2738 + v2739
	v2742 = int32(_a_F_heap_vacuum_rel_1)
	v2744 = *(*int32)(unsafe.Add(mBase, _c_F_heap_vacuum_rel[11]))
	*(*int32)(unsafe.Add(mBase, _c_F_heap_vacuum_rel[11])) = v2744 - v2739
	goto L396
L401:
	;
	goto L403
L403:
	;
	goto L404
L404:
	;
	v2699 = int32(0)
	v2702 = int32(0)
	goto L409
L409:
	;
	v2711 = *(*int32)(unsafe.Add(mBase, uint32(v53+int32(1664)+v2702<<(uint(int32(2))%32))))
	v2712 = int32(3)
	v2718 = *(*int64)(unsafe.Add(mBase, uint32(v53+int32(592)+v2702<<(uint(v2712)%32))))
	*(*int64)(unsafe.Add(mBase, uint32(v2585+int32(232)+v2711<<(uint(v2712)%32)))) = v2718
	v2720 = int32(1)
	v2723 = v2699 + v2720
	if v2723 != int32(2) {
		v2699 = v2723
		v2702 = v2702 + v2720
		goto L409
	} else {
		goto L411
	}
L410:
	;
	goto L400
L411:
	;
	goto L410
L412:
	;
	v2822 = *(*int32)(unsafe.Add(mBase, uint32(v231)+144))
	*(*int32)(unsafe.Add(mBase, uint32(v231)+144)) = v2822 + int32(1)
	goto L414
L413:
	;
	goto L414
L414:
	;
	if v2766&int32(1) != 0 {
		goto L415
	} else {
		goto L416
	}
L415:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v231)+148)) = v2073 + int32(1)
	goto L417
L416:
	;
	goto L417
L417:
	;
	v2831 = int32(0)
	v3246 = v2762
	v3250 = v2831
	v3254 = v2831
	goto L303
L418:
	;
	F_LockBufferForCleanup(m, v2031)
	mBase = m.M
	v2838 = m.ExcPending
	if v2838 != 0 {
		goto L24
	} else {
		goto L419
	}
L419:
	;
	v2839 = *(*int32)(unsafe.Add(mBase, uint32(v53)+1036))
	v2846 = v2839
	goto L304
L420:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v53)+608)) = int32(11)
	v2906 = int32(15)
	goto L422
L421:
	;
	v2906 = int32(14)
	goto L422
L422:
	;
	v2907 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v231)+21)))
	if v2907 == int32(1) {
		goto L423
	} else {
		goto L424
	}
L423:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v53)+608)) = v2906
	goto L425
L424:
	;
	goto L425
L425:
	;
	F_heap_page_prune_and_freeze(m, v53+int32(592), v53+int32(1040), v1660, v1442, v1440)
	mBase = m.M
	v2916 = m.ExcPending
	if v2916 != 0 {
		goto L24
	} else {
		goto L426
	}
L426:
	;
	v2917 = *(*int32)(unsafe.Add(mBase, uint32(v53)+1048))
	if int32(0) < v2917 {
		goto L427
	} else {
		goto L428
	}
L427:
	;
	v2920 = *(*int32)(unsafe.Add(mBase, uint32(v231)+124))
	*(*int32)(unsafe.Add(mBase, uint32(v231)+124)) = v2920 + int32(1)
	goto L429
L428:
	;
	goto L429
L429:
	;
	v2924 = *(*int32)(unsafe.Add(mBase, uint32(v53)+1064))
	if int32(0) < v2924 {
		goto L430
	} else {
		goto L431
	}
L430:
	;
	v2927 = *(*int32)(unsafe.Add(mBase, uint32(v231)+140))
	*(*int32)(unsafe.Add(mBase, uint32(v231)+140)) = v2927 + int32(1)
	F_pg_qsort(m, v1662, v2924, int32(2), int32(191))
	mBase = m.M
	v2934 = m.ExcPending
	if v2934 != 0 {
		goto L24
	} else {
		goto L433
	}
L431:
	;
	goto L432
L432:
	;
	v3142 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v53)+1060)))
	if v3142 == int32(1) {
		goto L453
	} else {
		goto L454
	}
L433:
	;
	v2935 = *(*int32)(unsafe.Add(mBase, uint32(v53)+1064))
	*(*int64)(unsafe.Add(mBase, uint32(v53)+1688)) = int64(25769803783)
	v2938 = *(*int32)(unsafe.Add(mBase, uint32(v231)+100))
	F_TidStoreSetBlockOffsets(m, v2938, v2073, v1662, v2935)
	mBase = m.M
	v2940 = m.ExcPending
	if v2940 != 0 {
		goto L24
	} else {
		goto L434
	}
L434:
	;
	v2941 = *(*int32)(unsafe.Add(mBase, uint32(v231)+104))
	v2942 = *(*int64)(unsafe.Add(mBase, uint32(v2941)+8))
	*(*int64)(unsafe.Add(mBase, uint32(v2941)+8)) = v2942 + base.I64_extend_i32_s(v2935)
	v2946 = *(*int32)(unsafe.Add(mBase, uint32(v231)+104))
	v2947 = *(*int64)(unsafe.Add(mBase, uint32(v2946)+8))
	*(*int64)(unsafe.Add(mBase, uint32(v53)+1664)) = v2947
	v2949 = *(*int32)(unsafe.Add(mBase, uint32(v231)+100))
	v2950 = F_TidStoreMemoryUsage(m, v2949)
	mBase = m.M
	v2951 = m.ExcPending
	if v2951 != 0 {
		goto L24
	} else {
		goto L435
	}
L435:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v53)+1672)) = base.I64_extend_i32_u(v2950)
	goto L438
L436:
	;
	goto L432
L437:
	;
	goto L436
L438:
	;
	v2968 = *(*int32)(unsafe.Add(mBase, _c_F_heap_vacuum_rel[9]))
	if v2968 == int32(0) {
		goto L437
	} else {
		goto L439
	}
L439:
	;
	v2972 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_heap_vacuum_rel[10])))
	if v2972&int32(1) == int32(0) {
		goto L437
	} else {
		goto L440
	}
L440:
	;
	v2977 = int32(_a_F_heap_vacuum_rel_1)
	v2979 = *(*int32)(unsafe.Add(mBase, _c_F_heap_vacuum_rel[11]))
	v2980 = int32(1)
	*(*int32)(unsafe.Add(mBase, _c_F_heap_vacuum_rel[11])) = v2979 + v2980
	v2983 = *(*int32)(unsafe.Add(mBase, uint32(v2968)))
	*(*int32)(unsafe.Add(mBase, uint32(v2968))) = v2983 + v2980
	v2987 = int32(0)
	v2990 = base.AtomicRmwOr32(m, v2987, int32(_a_F_heap_vacuum_rel_2), v2987)
	goto L442
L441:
	;
	v3117 = int32(0)
	v3120 = base.AtomicRmwOr32(m, v3117, int32(_a_F_heap_vacuum_rel_2), v3117)
	v3121 = *(*int32)(unsafe.Add(mBase, uint32(v2968)))
	v3122 = int32(1)
	*(*int32)(unsafe.Add(mBase, uint32(v2968))) = v3121 + v3122
	v3125 = int32(_a_F_heap_vacuum_rel_1)
	v3127 = *(*int32)(unsafe.Add(mBase, _c_F_heap_vacuum_rel[11]))
	*(*int32)(unsafe.Add(mBase, _c_F_heap_vacuum_rel[11])) = v3127 - v3122
	goto L437
L442:
	;
	goto L444
L444:
	;
	goto L445
L445:
	;
	v3082 = int32(0)
	v3085 = int32(0)
	goto L450
L450:
	;
	v3094 = *(*int32)(unsafe.Add(mBase, uint32(v53+int32(1688)+v3085<<(uint(int32(2))%32))))
	v3095 = int32(3)
	v3101 = *(*int64)(unsafe.Add(mBase, uint32(v53+int32(1664)+v3085<<(uint(v3095)%32))))
	*(*int64)(unsafe.Add(mBase, uint32(v2968+int32(232)+v3094<<(uint(v3095)%32)))) = v3101
	v3103 = int32(1)
	v3106 = v3082 + v3103
	if v3106 != int32(2) {
		v3082 = v3106
		v3085 = v3085 + v3103
		goto L450
	} else {
		goto L452
	}
L451:
	;
	goto L441
L452:
	;
	goto L451
L453:
	;
	v3145 = *(*int32)(unsafe.Add(mBase, uint32(v231)+128))
	*(*int32)(unsafe.Add(mBase, uint32(v231)+128)) = v3145 + int32(1)
	goto L455
L454:
	;
	goto L455
L455:
	;
	v3149 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v53)+1061)))
	if v3149 == int32(1) {
		goto L456
	} else {
		goto L457
	}
L456:
	;
	v3152 = *(*int32)(unsafe.Add(mBase, uint32(v231)+132))
	*(*int32)(unsafe.Add(mBase, uint32(v231)+132)) = v3152 + int32(1)
	goto L458
L457:
	;
	goto L458
L458:
	;
	v3156 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v53)+1062)))
	if v3156 == int32(1) {
		goto L459
	} else {
		goto L460
	}
L459:
	;
	v3159 = *(*int32)(unsafe.Add(mBase, uint32(v231)+136))
	*(*int32)(unsafe.Add(mBase, uint32(v231)+136)) = v3159 + int32(1)
	goto L461
L460:
	;
	goto L461
L461:
	;
	v3163 = *(*int64)(unsafe.Add(mBase, uint32(v231)+200))
	v3164 = *(*int32)(unsafe.Add(mBase, uint32(v53)+1040))
	*(*int64)(unsafe.Add(mBase, uint32(v231)+200)) = v3163 + base.I64_extend_i32_s(v3164)
	v3168 = *(*int64)(unsafe.Add(mBase, uint32(v231)+208))
	v3169 = int64(*(*int32)(unsafe.Add(mBase, uint32(v53)+1048)))
	*(*int64)(unsafe.Add(mBase, uint32(v231)+208)) = v3168 + v3169
	v3172 = *(*int64)(unsafe.Add(mBase, uint32(v231)+216))
	v3173 = *(*int32)(unsafe.Add(mBase, uint32(v53)+1064))
	*(*int64)(unsafe.Add(mBase, uint32(v231)+216)) = v3172 + base.I64_extend_i32_s(v3173)
	v3177 = *(*int64)(unsafe.Add(mBase, uint32(v231)+224))
	v3178 = int64(*(*int32)(unsafe.Add(mBase, uint32(v53)+1052)))
	*(*int64)(unsafe.Add(mBase, uint32(v231)+224)) = v3177 + v3178
	v3181 = *(*int64)(unsafe.Add(mBase, uint32(v231)+232))
	v3182 = int64(*(*int32)(unsafe.Add(mBase, uint32(v53)+1056)))
	*(*int64)(unsafe.Add(mBase, uint32(v231)+232)) = v3181 + v3182
	v3185 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v53)+1063)))
	if v3185 == int32(1) {
		goto L462
	} else {
		goto L463
	}
L462:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v231)+148)) = v2073 + int32(1)
	goto L464
L463:
	;
	goto L464
L464:
	;
	v3191 = int32(0)
	v3192 = base.B2i32(v3191 < v3164)
	v3194 = base.B2i32(v3191 < v3173)
	v3195 = int32(1)
	if v2034&v3195 == v3191 {
		v3246 = v3194
		v3250 = v3195
		v3254 = v3192
		goto L303
	} else {
		goto L465
	}
L465:
	;
	if v3149|v3156 != 0 {
		goto L466
	} else {
		goto L467
	}
L466:
	;
	v3201 = *(*int32)(unsafe.Add(mBase, uint32(v231)+268))
	if v3201 != 0 {
		goto L469
	} else {
		goto L470
	}
L467:
	;
	goto L468
L468:
	;
	v3238 = *(*int32)(unsafe.Add(mBase, uint32(v231)+276))
	if v3238 == int32(0) {
		v3246 = v3194
		v3250 = v3195
		v3254 = v3192
		goto L303
	} else {
		goto L482
	}
L469:
	;
	v3203 = v3201 - int32(1)
	*(*int32)(unsafe.Add(mBase, uint32(v231)+268)) = v3203
	if v3203 != 0 {
		v3246 = v3194
		v3250 = v3195
		v3254 = v3192
		goto L303
	} else {
		goto L472
	}
L470:
	;
	goto L471
L471:
	;
	v3206 = *(*int32)(unsafe.Add(mBase, uint32(v231)+272))
	if v3206 == int32(0) {
		goto L473
	} else {
		goto L474
	}
L472:
	;
	goto L471
L473:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v231)+264)) = int32(-1)
	*(*int64)(unsafe.Add(mBase, uint32(v231)+272)) = int64(0)
	v3246 = v3194
	v3250 = v3195
	v3254 = v3192
	goto L303
L474:
	;
	v3211 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v231)+96)))
	if v3211 != 0 {
		goto L475
	} else {
		goto L476
	}
L475:
	;
	v3212 = int32(17)
	goto L477
L476:
	;
	v3212 = int32(13)
	goto L477
L477:
	;
	v3214 = F_errstart(m, v3212, int32(0))
	mBase = m.M
	v3215 = m.ExcPending
	if v3215 != 0 {
		goto L24
	} else {
		goto L478
	}
L478:
	;
	if v3214 == int32(0) {
		goto L473
	} else {
		goto L479
	}
L479:
	;
	v3218 = *(*int64)(unsafe.Add(mBase, uint32(v231)+68))
	v3219 = *(*int32)(unsafe.Add(mBase, uint32(v231)+76))
	*(*int32)(unsafe.Add(mBase, uint32(v53)+556)) = v3219
	*(*int64)(unsafe.Add(mBase, uint32(v53)+548)) = v3218
	*(*int32)(unsafe.Add(mBase, uint32(v53)+544)) = v1450
	F_errmsg(m, int32(_a_F_heap_vacuum_rel_22), v53+int32(544))
	mBase = m.M
	v3227 = m.ExcPending
	if v3227 != 0 {
		goto L24
	} else {
		goto L480
	}
L480:
	;
	F_errfinish(m, int32(_a_F_heap_vacuum_rel_7), int32(1525), int32(_a_F_heap_vacuum_rel_23))
	mBase = m.M
	v3232 = m.ExcPending
	if v3232 != 0 {
		goto L24
	} else {
		goto L481
	}
L481:
	;
	goto L473
L482:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v231)+276)) = v3238 - int32(1)
	v3246 = v3194
	v3250 = v3195
	v3254 = v3192
	goto L303
L483:
	;
	F_UnlockReleaseBuffer(m, v2031)
	mBase = m.M
	v3387 = m.ExcPending
	if v3387 != 0 {
		goto L24
	} else {
		goto L512
	}
L484:
	;
	v3295 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v231)+23)))
	if v3295&v3246&int32(1) != 0 {
		goto L483
	} else {
		goto L487
	}
L485:
	;
	goto L486
L486:
	;
	v3302 = int32(4)
	v3303 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v2054)+14)))
	v3304 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v2054)+12)))
	v3305 = v3303 - v3304
	if v3305 <= v3302 {
		goto L489
	} else {
		goto L490
	}
L487:
	;
	goto L486
L488:
	;
	F_UnlockReleaseBuffer(m, v2031)
	mBase = m.M
	v3367 = m.ExcPending
	if v3367 != 0 {
		goto L24
	} else {
		goto L507
	}
L489:
	;
	v3308 = v3302
	goto L491
L490:
	;
	v3308 = v3305
	goto L491
L491:
	;
	v3310 = v3308 - int32(4)
	if v3310 == int32(0) {
		goto L492
	} else {
		goto L493
	}
L492:
	;
	v3365 = int32(0)
	goto L488
L493:
	;
	goto L494
L494:
	;
	if base.Ui32(int32(25)) <= base.Ui32(v3304) {
		goto L496
	} else {
		goto L497
	}
L495:
	;
	v3365 = v3310
	goto L488
L496:
	;
	v3321 = int32(base.Ui32(v3304+int32(_a_F_heap_vacuum_rel_16)) >> (uint(int32(2)) % 32))
	goto L498
L497:
	;
	v3321 = int32(0)
	goto L498
L498:
	;
	if base.Ui32(v3321&int32(_a_F_heap_vacuum_rel_17)) < base.Ui32(int32(291)) {
		goto L495
	} else {
		goto L499
	}
L499:
	;
	v3326 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v2054)+10)))
	if v3326&int32(1) == int32(0) {
		goto L500
	} else {
		goto L501
	}
L500:
	;
	v3365 = int32(0)
	goto L488
L501:
	;
	goto L502
L502:
	;
	v3335 = int32(1)
	goto L503
L503:
	;
	v3344 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v2054+int32(20)+v3335&int32(_a_F_heap_vacuum_rel_17)<<(uint(int32(2))%32))+1)))
	if v3344&int32(384) == int32(0) {
		goto L495
	} else {
		goto L505
	}
L504:
	;
	v3365 = int32(0)
	goto L488
L505:
	;
	v3350 = v3335 + int32(1)
	v3351 = int32(_a_F_heap_vacuum_rel_17)
	if base.Ui32(v3350&v3351) <= base.Ui32(v3321&v3351) {
		v3335 = v3350
		goto L503
	} else {
		goto L506
	}
L506:
	;
	goto L504
L507:
	;
	v3368 = *(*int32)(unsafe.Add(mBase, uint32(v231)))
	F_RecordPageWithFreeSpace(m, v3368, v2073, v3365)
	mBase = m.M
	v3370 = m.ExcPending
	if v3370 != 0 {
		goto L24
	} else {
		goto L508
	}
L508:
	;
	if v3250 == int32(0) {
		v1677 = v2073
		v1681 = v1807
		v1703 = v2010
		goto L244
	} else {
		goto L509
	}
L509:
	;
	v3373 = *(*int32)(unsafe.Add(mBase, uint32(v267)))
	v3374 = int32(0)
	if base.B2i32(base.B2i32(v3373 == v3374)&v3254 == v3374)|base.B2i32(base.Ui32(v2073-v1807) < base.Ui32(int32(_a_F_heap_vacuum_rel_24))) != 0 {
		v1677 = v2073
		v1681 = v1807
		v1703 = v2010
		goto L244
	} else {
		goto L510
	}
L510:
	;
	v3383 = *(*int32)(unsafe.Add(mBase, uint32(v231)))
	F_FreeSpaceMapVacuumRange(m, v3383, v1807, v2073)
	mBase = m.M
	v3385 = m.ExcPending
	if v3385 != 0 {
		goto L24
	} else {
		goto L511
	}
L511:
	;
	v1677 = v2073
	v1681 = v2073
	v1703 = v2010
	goto L244
L512:
	;
	v1677 = v2073
	v1681 = v1807
	v1703 = v2010
	goto L244
L513:
	;
	F_ReleaseBuffer(m, v3390)
	mBase = m.M
	v3392 = m.ExcPending
	if v3392 != 0 {
		goto L24
	} else {
		goto L516
	}
L514:
	;
	goto L515
L515:
	;
	v3396 = *(*int32)(unsafe.Add(mBase, _c_F_heap_vacuum_rel[9]))
	if v3396 == int32(0) {
		goto L518
	} else {
		goto L519
	}
L516:
	;
	goto L515
L517:
	;
	v3437 = *(*int32)(unsafe.Add(mBase, uint32(v231)))
	v3438 = *(*int32)(unsafe.Add(mBase, uint32(v231)+112))
	v3439 = *(*int64)(unsafe.Add(mBase, uint32(v231)+224))
	v3440 = base.F64_convert_i64_s(v3439)
	if base.Ui32(v3438) < base.Ui32(v1451) {
		goto L522
	} else {
		goto L523
	}
L518:
	;
	goto L517
L519:
	;
	v3400 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_heap_vacuum_rel[10])))
	if v3400&int32(1) == int32(0) {
		goto L518
	} else {
		goto L520
	}
L520:
	;
	v3405 = int32(_a_F_heap_vacuum_rel_1)
	v3407 = *(*int32)(unsafe.Add(mBase, _c_F_heap_vacuum_rel[11]))
	v3408 = int32(1)
	*(*int32)(unsafe.Add(mBase, _c_F_heap_vacuum_rel[11])) = v3407 + v3408
	v3411 = *(*int32)(unsafe.Add(mBase, uint32(v3396)))
	*(*int32)(unsafe.Add(mBase, uint32(v3396))) = v3411 + v3408
	v3415 = int32(0)
	v3417 = int32(_a_F_heap_vacuum_rel_2)
	v3418 = base.AtomicRmwOr32(m, v3415, v3417, v3415)
	*(*int64)(unsafe.Add(mBase, uint32(v3396+int32(16))+232)) = v1460
	v3426 = base.AtomicRmwOr32(m, v3415, v3417, v3415)
	v3427 = *(*int32)(unsafe.Add(mBase, uint32(v3396)))
	*(*int32)(unsafe.Add(mBase, uint32(v3396))) = v3427 + v3408
	v3433 = *(*int32)(unsafe.Add(mBase, _c_F_heap_vacuum_rel[11]))
	*(*int32)(unsafe.Add(mBase, _c_F_heap_vacuum_rel[11])) = v3433 - v3408
	goto L518
L521:
	;
	*(*float64)(unsafe.Add(mBase, uint32(v231)+160)) = v3490
	v3492 = float64(0)
	if base.F64_gt(v3490, v3492) != 0 {
		goto L540
	} else {
		goto L541
	}
L522:
	;
	v3445 = *(*int32)(unsafe.Add(mBase, uint32(v3437)+48))
	v3446 = *(*float32)(unsafe.Add(mBase, uint32(v3445)+100))
	v3447 = base.F64_promote_f32(v3446)
	v3448 = *(*int32)(unsafe.Add(mBase, uint32(v3445)+96))
	if v1451 == v3448 {
		goto L526
	} else {
		goto L527
	}
L523:
	;
	v3485 = v3440
	goto L524
L524:
	;
	v3490 = v3485
	goto L521
L525:
	;
	v3461 = base.F64_convert_i32_u(v1451)
	if v3448 != 0 {
		goto L534
	} else {
		goto L535
	}
L526:
	;
	if base.Ui32(v3438) < base.Ui32(int32(2)) {
		goto L529
	} else {
		goto L530
	}
L527:
	;
	goto L528
L528:
	;
	if base.Ui32(int32(2)) <= base.Ui32(v3438) {
		goto L525
	} else {
		goto L533
	}
L529:
	;
	v3490 = v3447
	goto L521
L530:
	;
	goto L531
L531:
	;
	if base.F64_lt(base.F64_convert_i32_u(v3438), base.F64_mul(base.F64_convert_i32_u(v1451), float64(0.02))) == int32(0) {
		goto L525
	} else {
		goto L532
	}
L532:
	;
	v3490 = v3447
	goto L521
L533:
	;
	v3490 = v3447
	goto L521
L534:
	;
	v3465 = base.F32_lt(v3446, float32(0))
	goto L536
L535:
	;
	v3465 = int32(1)
	goto L536
L536:
	;
	if v3465 != 0 {
		goto L537
	} else {
		goto L538
	}
L537:
	;
	v3490 = base.F64_floor(base.F64_add(base.F64_mul(base.F64_div(v3440, base.F64_convert_i32_u(v3438)), v3461), float64(0.5)))
	goto L521
L538:
	;
	goto L539
L539:
	;
	v3485 = base.F64_floor(base.F64_add(base.F64_add(base.F64_mul(base.F64_div(v3447, base.F64_convert_i32_u(v3448)), base.F64_sub(v3461, base.F64_convert_i32_u(v3438))), v3440), float64(0.5)))
	goto L524
L540:
	;
	v3495 = v3490
	goto L542
L541:
	;
	v3495 = v3492
	goto L542
L542:
	;
	v3496 = *(*int64)(unsafe.Add(mBase, uint32(v231)+232))
	v3499 = *(*int64)(unsafe.Add(mBase, uint32(v231)+240))
	*(*float64)(unsafe.Add(mBase, uint32(v231)+152)) = base.F64_add(base.F64_add(v3495, base.F64_convert_i64_s(v3496)), base.F64_convert_i64_s(v3499))
	F_read_stream_end(m, v1669)
	mBase = m.M
	v3504 = m.ExcPending
	if v3504 != 0 {
		goto L24
	} else {
		goto L543
	}
L543:
	;
	v3505 = *(*int32)(unsafe.Add(mBase, uint32(v231)+104))
	v3506 = *(*int64)(unsafe.Add(mBase, uint32(v3505)+8))
	if int64(0) < v3506 {
		goto L544
	} else {
		goto L545
	}
L544:
	;
	F_lazy_vacuum(m, v231)
	mBase = m.M
	v3510 = m.ExcPending
	if v3510 != 0 {
		goto L24
	} else {
		goto L547
	}
L545:
	;
	goto L546
L546:
	;
	if base.Ui32(v1807) < base.Ui32(v1451) {
		goto L548
	} else {
		goto L549
	}
L547:
	;
	goto L546
L548:
	;
	v3512 = *(*int32)(unsafe.Add(mBase, uint32(v231)))
	F_FreeSpaceMapVacuumRange(m, v3512, v1807, v1451)
	mBase = m.M
	v3514 = m.ExcPending
	if v3514 != 0 {
		goto L24
	} else {
		goto L551
	}
L549:
	;
	goto L550
L550:
	;
	v3518 = *(*int32)(unsafe.Add(mBase, _c_F_heap_vacuum_rel[9]))
	if v3518 == int32(0) {
		goto L553
	} else {
		goto L554
	}
L551:
	;
	goto L550
L552:
	;
	v3559 = *(*int32)(unsafe.Add(mBase, uint32(v267)))
	if v3559 <= int32(0) {
		goto L556
	} else {
		goto L557
	}
L553:
	;
	goto L552
L554:
	;
	v3522 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_heap_vacuum_rel[10])))
	if v3522&int32(1) == int32(0) {
		goto L553
	} else {
		goto L555
	}
L555:
	;
	v3527 = int32(_a_F_heap_vacuum_rel_1)
	v3529 = *(*int32)(unsafe.Add(mBase, _c_F_heap_vacuum_rel[11]))
	v3530 = int32(1)
	*(*int32)(unsafe.Add(mBase, _c_F_heap_vacuum_rel[11])) = v3529 + v3530
	v3533 = *(*int32)(unsafe.Add(mBase, uint32(v3518)))
	*(*int32)(unsafe.Add(mBase, uint32(v3518))) = v3533 + v3530
	v3537 = int32(0)
	v3539 = int32(_a_F_heap_vacuum_rel_2)
	v3540 = base.AtomicRmwOr32(m, v3537, v3539, v3537)
	*(*int64)(unsafe.Add(mBase, uint32(v3518+int32(24))+232)) = v1460
	v3548 = base.AtomicRmwOr32(m, v3537, v3539, v3537)
	v3549 = *(*int32)(unsafe.Add(mBase, uint32(v3518)))
	*(*int32)(unsafe.Add(mBase, uint32(v3518))) = v3549 + v3530
	v3555 = *(*int32)(unsafe.Add(mBase, _c_F_heap_vacuum_rel[11]))
	*(*int32)(unsafe.Add(mBase, _c_F_heap_vacuum_rel[11])) = v3555 - v3530
	goto L553
L556:
	;
	v4219 = *(*int32)(unsafe.Add(mBase, uint32(v231)+104))
	v4220 = *(*int32)(unsafe.Add(mBase, uint32(v4219)))
	v4221 = *(*int32)(unsafe.Add(mBase, uint32(v231)+100))
	v4222 = F_TidStoreMemoryUsage(m, v4221)
	mBase = m.M
	v4223 = m.ExcPending
	if v4223 != 0 {
		goto L24
	} else {
		goto L609
	}
L557:
	;
	v3562 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v231)+24)))
	if v3562 != int32(1) {
		goto L556
	} else {
		goto L558
	}
L558:
	;
	v3565 = *(*int32)(unsafe.Add(mBase, uint32(v231)+108))
	v3566 = *(*int32)(unsafe.Add(mBase, uint32(v231)+112))
	v3567 = *(*float64)(unsafe.Add(mBase, uint32(v231)+152))
	*(*int64)(unsafe.Add(mBase, uint32(v53)+1688)) = int64(34359738368)
	*(*int64)(unsafe.Add(mBase, uint32(v53)+1680)) = int64(38654705672)
	*(*int64)(unsafe.Add(mBase, uint32(v53)+600)) = base.I64_extend_i32_u(v3559)
	*(*int64)(unsafe.Add(mBase, uint32(v53)+592)) = int64(4)
	v3576 = int64(0)
	*(*int64)(unsafe.Add(mBase, uint32(v53)+1672)) = v3576
	*(*int64)(unsafe.Add(mBase, uint32(v53)+1664)) = v3576
	goto L561
L559:
	;
	v3766 = *(*int32)(unsafe.Add(mBase, uint32(v231)+16))
	if v3766 == int32(0) {
		goto L577
	} else {
		goto L578
	}
L560:
	;
	goto L559
L561:
	;
	v3594 = *(*int32)(unsafe.Add(mBase, _c_F_heap_vacuum_rel[9]))
	if v3594 == int32(0) {
		goto L560
	} else {
		goto L562
	}
L562:
	;
	v3598 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_heap_vacuum_rel[10])))
	if v3598&int32(1) == int32(0) {
		goto L560
	} else {
		goto L563
	}
L563:
	;
	v3603 = int32(_a_F_heap_vacuum_rel_1)
	v3605 = *(*int32)(unsafe.Add(mBase, _c_F_heap_vacuum_rel[11]))
	v3606 = int32(1)
	*(*int32)(unsafe.Add(mBase, _c_F_heap_vacuum_rel[11])) = v3605 + v3606
	v3609 = *(*int32)(unsafe.Add(mBase, uint32(v3594)))
	*(*int32)(unsafe.Add(mBase, uint32(v3594))) = v3609 + v3606
	v3613 = int32(0)
	v3616 = base.AtomicRmwOr32(m, v3613, int32(_a_F_heap_vacuum_rel_2), v3613)
	goto L565
L564:
	;
	v3743 = int32(0)
	v3746 = base.AtomicRmwOr32(m, v3743, int32(_a_F_heap_vacuum_rel_2), v3743)
	v3747 = *(*int32)(unsafe.Add(mBase, uint32(v3594)))
	v3748 = int32(1)
	*(*int32)(unsafe.Add(mBase, uint32(v3594))) = v3747 + v3748
	v3751 = int32(_a_F_heap_vacuum_rel_1)
	v3753 = *(*int32)(unsafe.Add(mBase, _c_F_heap_vacuum_rel[11]))
	*(*int32)(unsafe.Add(mBase, _c_F_heap_vacuum_rel[11])) = v3753 - v3748
	goto L560
L565:
	;
	goto L567
L567:
	;
	goto L568
L568:
	;
	v3708 = int32(0)
	v3711 = int32(0)
	goto L573
L573:
	;
	v3720 = *(*int32)(unsafe.Add(mBase, uint32(v53+int32(1688)+v3711<<(uint(int32(2))%32))))
	v3721 = int32(3)
	v3727 = *(*int64)(unsafe.Add(mBase, uint32(v53+int32(592)+v3711<<(uint(v3721)%32))))
	*(*int64)(unsafe.Add(mBase, uint32(v3594+int32(232)+v3720<<(uint(v3721)%32)))) = v3727
	v3729 = int32(1)
	v3732 = v3708 + v3729
	if v3732 != int32(2) {
		v3708 = v3732
		v3711 = v3711 + v3729
		goto L573
	} else {
		goto L575
	}
L574:
	;
	goto L564
L575:
	;
	goto L574
L576:
	;
	goto L594
L577:
	;
	v3769 = *(*int32)(unsafe.Add(mBase, uint32(v267)))
	if v3769 <= int32(0) {
		goto L576
	} else {
		goto L580
	}
L578:
	;
	goto L579
L579:
	;
	v3922 = *(*int32)(unsafe.Add(mBase, uint32(v442)))
	v3923 = *(*int32)(unsafe.Add(mBase, uint32(v3766)+16))
	*(*float64)(unsafe.Add(mBase, uint32(v3923)+16)) = base.F64_convert_i32_s(base.I32_trunc_sat_f64_s(v3567))
	v3927 = *(*int32)(unsafe.Add(mBase, uint32(v3766)+16))
	*(*uint8)(unsafe.Add(mBase, uint32(v3927)+24)) = uint8(base.B2i32(base.Ui32(v3566) < base.Ui32(v3565)))
	F_parallel_vacuum_process_all_indexes(m, v3766, v3922, int32(0), v1444)
	mBase = m.M
	v3932 = m.ExcPending
	if v3932 != 0 {
		goto L24
	} else {
		goto L591
	}
L580:
	;
	v3813 = int64(0)
	goto L581
L581:
	;
	v3826 = base.I32_wrap_i64(v3813) << (uint(int32(2)) % 32)
	v3827 = *(*int32)(unsafe.Add(mBase, uint32(v231)+168))
	v3829 = *(*int32)(unsafe.Add(mBase, uint32(v3826+v3827)))
	v3830 = *(*int32)(unsafe.Add(mBase, uint32(v231)+4))
	v3832 = *(*int32)(unsafe.Add(mBase, uint32(v3830+v3826)))
	*(*int32)(unsafe.Add(mBase, uint32(v53)+1040)) = v3832
	v3834 = *(*int32)(unsafe.Add(mBase, uint32(v231)))
	*(*float64)(unsafe.Add(mBase, uint32(v53)+1056)) = v3567
	*(*int32)(unsafe.Add(mBase, uint32(v53)+1052)) = int32(13)
	*(*uint8)(unsafe.Add(mBase, uint32(v53)+1050)) = uint8(base.B2i32(base.Ui32(v3566) < base.Ui32(v3565)))
	v3839 = int32(0)
	*(*uint16)(unsafe.Add(mBase, uint32(v53)+1048)) = uint16(v3839)
	*(*int32)(unsafe.Add(mBase, uint32(v53)+1044)) = v3834
	v3842 = *(*int32)(unsafe.Add(mBase, uint32(v231)+12))
	*(*int32)(unsafe.Add(mBase, uint32(v53)+1064)) = v3842
	v3844 = *(*int32)(unsafe.Add(mBase, uint32(v3832)+48))
	v3847 = F_pstrdup(m, v3844+int32(4))
	mBase = m.M
	v3848 = m.ExcPending
	if v3848 != 0 {
		goto L24
	} else {
		goto L583
	}
L582:
	;
	goto L576
L583:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v231)+80)) = v3847
	v3850 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v231)+88)))
	v3851 = int32(0)
	*(*uint16)(unsafe.Add(mBase, uint32(v231)+88)) = uint16(v3851)
	v3853 = *(*int32)(unsafe.Add(mBase, uint32(v231)+84))
	*(*int32)(unsafe.Add(mBase, uint32(v231)+84)) = int32(-1)
	v3856 = *(*int32)(unsafe.Add(mBase, uint32(v231)+92))
	*(*int32)(unsafe.Add(mBase, uint32(v231)+92)) = int32(4)
	v3861 = F_vac_cleanup_one_index(m, v53+int32(1040), v3829)
	mBase = m.M
	v3862 = m.ExcPending
	if v3862 != 0 {
		goto L24
	} else {
		goto L584
	}
L584:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v231)+92)) = v3856
	*(*uint16)(unsafe.Add(mBase, uint32(v231)+88)) = uint16(v3850)
	*(*int32)(unsafe.Add(mBase, uint32(v231)+84)) = v3853
	v3866 = *(*int32)(unsafe.Add(mBase, uint32(v231)+80))
	F_pfree(m, v3866)
	mBase = m.M
	v3868 = m.ExcPending
	if v3868 != 0 {
		goto L24
	} else {
		goto L585
	}
L585:
	;
	v3869 = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v231)+80)) = v3869
	v3871 = *(*int32)(unsafe.Add(mBase, uint32(v231)+168))
	*(*int32)(unsafe.Add(mBase, uint32(v3871+v3826))) = v3861
	v3876 = v3813 + int64(1)
	v3879 = *(*int32)(unsafe.Add(mBase, _c_F_heap_vacuum_rel[9]))
	if v3879 == v3869 {
		goto L587
	} else {
		goto L588
	}
L586:
	;
	v3920 = int64(*(*int32)(unsafe.Add(mBase, uint32(v231)+8)))
	if v3876 < v3920 {
		v3813 = v3876
		goto L581
	} else {
		goto L590
	}
L587:
	;
	goto L586
L588:
	;
	v3883 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_heap_vacuum_rel[10])))
	if v3883&int32(1) == int32(0) {
		goto L587
	} else {
		goto L589
	}
L589:
	;
	v3888 = int32(_a_F_heap_vacuum_rel_1)
	v3890 = *(*int32)(unsafe.Add(mBase, _c_F_heap_vacuum_rel[11]))
	v3891 = int32(1)
	*(*int32)(unsafe.Add(mBase, _c_F_heap_vacuum_rel[11])) = v3890 + v3891
	v3894 = *(*int32)(unsafe.Add(mBase, uint32(v3879)))
	*(*int32)(unsafe.Add(mBase, uint32(v3879))) = v3894 + v3891
	v3898 = int32(0)
	v3900 = int32(_a_F_heap_vacuum_rel_2)
	v3901 = base.AtomicRmwOr32(m, v3898, v3900, v3898)
	*(*int64)(unsafe.Add(mBase, uint32(v3879+int32(72))+232)) = v3876
	v3909 = base.AtomicRmwOr32(m, v3898, v3900, v3898)
	v3910 = *(*int32)(unsafe.Add(mBase, uint32(v3879)))
	*(*int32)(unsafe.Add(mBase, uint32(v3879))) = v3910 + v3891
	v3916 = *(*int32)(unsafe.Add(mBase, _c_F_heap_vacuum_rel[11]))
	*(*int32)(unsafe.Add(mBase, _c_F_heap_vacuum_rel[11])) = v3916 - v3891
	goto L587
L590:
	;
	goto L582
L591:
	;
	goto L576
L592:
	;
	goto L556
L593:
	;
	goto L592
L594:
	;
	v3997 = *(*int32)(unsafe.Add(mBase, _c_F_heap_vacuum_rel[9]))
	if v3997 == int32(0) {
		goto L593
	} else {
		goto L595
	}
L595:
	;
	v4001 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_heap_vacuum_rel[10])))
	if v4001&int32(1) == int32(0) {
		goto L593
	} else {
		goto L596
	}
L596:
	;
	v4006 = int32(_a_F_heap_vacuum_rel_1)
	v4008 = *(*int32)(unsafe.Add(mBase, _c_F_heap_vacuum_rel[11]))
	v4009 = int32(1)
	*(*int32)(unsafe.Add(mBase, _c_F_heap_vacuum_rel[11])) = v4008 + v4009
	v4012 = *(*int32)(unsafe.Add(mBase, uint32(v3997)))
	*(*int32)(unsafe.Add(mBase, uint32(v3997))) = v4012 + v4009
	v4016 = int32(0)
	v4019 = base.AtomicRmwOr32(m, v4016, int32(_a_F_heap_vacuum_rel_2), v4016)
	goto L598
L597:
	;
	v4146 = int32(0)
	v4149 = base.AtomicRmwOr32(m, v4146, int32(_a_F_heap_vacuum_rel_2), v4146)
	v4150 = *(*int32)(unsafe.Add(mBase, uint32(v3997)))
	v4151 = int32(1)
	*(*int32)(unsafe.Add(mBase, uint32(v3997))) = v4150 + v4151
	v4154 = int32(_a_F_heap_vacuum_rel_1)
	v4156 = *(*int32)(unsafe.Add(mBase, _c_F_heap_vacuum_rel[11]))
	*(*int32)(unsafe.Add(mBase, _c_F_heap_vacuum_rel[11])) = v4156 - v4151
	goto L593
L598:
	;
	goto L600
L600:
	;
	goto L601
L601:
	;
	v4111 = int32(0)
	v4114 = int32(0)
	goto L606
L606:
	;
	v4123 = *(*int32)(unsafe.Add(mBase, uint32(v53+int32(1680)+v4114<<(uint(int32(2))%32))))
	v4124 = int32(3)
	v4130 = *(*int64)(unsafe.Add(mBase, uint32(v53+int32(1664)+v4114<<(uint(v4124)%32))))
	*(*int64)(unsafe.Add(mBase, uint32(v3997+int32(232)+v4123<<(uint(v4124)%32)))) = v4130
	v4132 = int32(1)
	v4135 = v4111 + v4132
	if v4135 != int32(2) {
		v4111 = v4135
		v4114 = v4114 + v4132
		goto L606
	} else {
		goto L608
	}
L607:
	;
	goto L597
L608:
	;
	goto L607
L609:
	;
	v4224 = *(*int32)(unsafe.Add(mBase, uint32(v231)+180))
	*(*int32)(unsafe.Add(mBase, uint32(v231)+180)) = v4222 + v4224
	v4227 = *(*int32)(unsafe.Add(mBase, uint32(v231)+16))
	if v4227 != 0 {
		goto L610
	} else {
		goto L611
	}
L610:
	;
	v4228 = *(*int32)(unsafe.Add(mBase, uint32(v231)+168))
	v4229 = int32(0)
	v4230 = *(*int32)(unsafe.Add(mBase, uint32(v4227)+12))
	if v4229 < v4230 {
		goto L613
	} else {
		goto L614
	}
L611:
	;
	goto L612
L612:
	;
	v4444 = *(*int32)(unsafe.Add(mBase, uint32(v231)+4))
	v4445 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v231)+24)))
	v4448 = *(*int32)(unsafe.Add(mBase, uint32(v231)+8))
	if base.B2i32(v4445 != int32(1))|base.B2i32(v4448 <= int32(0)) != 0 {
		goto L632
	} else {
		goto L633
	}
L613:
	;
	v4239 = v4229
	goto L616
L614:
	;
	goto L615
L615:
	;
	v4367 = *(*int32)(unsafe.Add(mBase, uint32(v4227)+24))
	F_TidStoreDestroy(m, v4367)
	mBase = m.M
	v4369 = m.ExcPending
	if v4369 != 0 {
		goto L24
	} else {
		goto L624
	}
L616:
	;
	v4283 = *(*int32)(unsafe.Add(mBase, uint32(v4227)+20))
	v4286 = v4283 + v4239*int32(48)
	v4287 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v4286)+5)))
	if v4287 == int32(1) {
		goto L619
	} else {
		goto L620
	}
L617:
	;
	goto L615
L618:
	;
	v4314 = v4239 + int32(1)
	v4315 = *(*int32)(unsafe.Add(mBase, uint32(v4227)+12))
	if v4314 < v4315 {
		v4239 = v4314
		goto L616
	} else {
		goto L623
	}
L619:
	;
	v4294 = F_palloc0(m, int32(40))
	mBase = m.M
	v4295 = m.ExcPending
	if v4295 != 0 {
		goto L24
	} else {
		goto L622
	}
L620:
	;
	goto L621
L621:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v4228+v4239<<(uint(int32(2))%32)))) = int32(0)
	goto L618
L622:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v4228+v4239<<(uint(int32(2))%32)))) = v4294
	v4297 = *(*int64)(unsafe.Add(mBase, uint32(v4286)+40))
	*(*int64)(unsafe.Add(mBase, uint32(v4294)+32)) = v4297
	v4299 = *(*int64)(unsafe.Add(mBase, uint32(v4286)+32))
	*(*int64)(unsafe.Add(mBase, uint32(v4294)+24)) = v4299
	v4301 = *(*int64)(unsafe.Add(mBase, uint32(v4286)+24))
	*(*int64)(unsafe.Add(mBase, uint32(v4294)+16)) = v4301
	v4303 = *(*int64)(unsafe.Add(mBase, uint32(v4286)+16))
	*(*int64)(unsafe.Add(mBase, uint32(v4294)+8)) = v4303
	v4305 = *(*int64)(unsafe.Add(mBase, uint32(v4286)+8))
	*(*int64)(unsafe.Add(mBase, uint32(v4294))) = v4305
	goto L618
L623:
	;
	goto L617
L624:
	;
	v4370 = *(*int32)(unsafe.Add(mBase, uint32(v4227)))
	F_DestroyParallelContext(m, v4370)
	mBase = m.M
	v4372 = m.ExcPending
	if v4372 != 0 {
		goto L24
	} else {
		goto L625
	}
L625:
	;
	v4375 = *(*int32)(unsafe.Add(mBase, _c_F_heap_vacuum_rel[22]))
	v4376 = *(*int32)(unsafe.Add(mBase, uint32(v4375)+72))
	*(*int32)(unsafe.Add(mBase, uint32(v4375)+72)) = v4376 - int32(1)
	goto L626
L626:
	;
	v4381 = *(*int32)(unsafe.Add(mBase, _c_F_heap_vacuum_rel[5]))
	if v4381 == int32(4) {
		goto L627
	} else {
		goto L628
	}
L627:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_heap_vacuum_rel[29])) = int32(0)
	goto L629
L628:
	;
	goto L629
L629:
	;
	v4387 = *(*int32)(unsafe.Add(mBase, uint32(v4227)+36))
	F_pfree(m, v4387)
	mBase = m.M
	v4389 = m.ExcPending
	if v4389 != 0 {
		goto L24
	} else {
		goto L630
	}
L630:
	;
	F_pfree(m, v4227)
	mBase = m.M
	v4391 = m.ExcPending
	if v4391 != 0 {
		goto L24
	} else {
		goto L631
	}
L631:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v231)+16)) = int32(0)
	goto L612
L632:
	;
	v4540 = v4444
	v4580 = v4448
	goto L634
L633:
	;
	v4452 = *(*int32)(unsafe.Add(mBase, uint32(v231)+168))
	v4456 = int32(0)
	goto L635
L634:
	;
	F_vac_close_indexes(m, v4580, v4540, int32(0))
	mBase = m.M
	v4583 = m.ExcPending
	if v4583 != 0 {
		goto L24
	} else {
		goto L642
	}
L635:
	;
	v4505 = v4456 << (uint(int32(2)) % 32)
	v4507 = *(*int32)(unsafe.Add(mBase, uint32(v4452+v4505)))
	if v4507 == int32(0) {
		goto L637
	} else {
		goto L638
	}
L636:
	;
	v4528 = *(*int32)(unsafe.Add(mBase, uint32(v269)))
	v4529 = *(*int32)(unsafe.Add(mBase, uint32(v267)))
	v4540 = v4528
	v4580 = v4529
	goto L634
L637:
	;
	v4526 = v4456 + int32(1)
	if v4526 != v4448 {
		v4456 = v4526
		goto L635
	} else {
		goto L641
	}
L638:
	;
	v4510 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v4507)+4)))
	if v4510 != 0 {
		goto L637
	} else {
		goto L639
	}
L639:
	;
	v4512 = *(*int32)(unsafe.Add(mBase, uint32(v4505+v4444)))
	v4513 = *(*int32)(unsafe.Add(mBase, uint32(v4507)))
	v4514 = *(*float64)(unsafe.Add(mBase, uint32(v4507)+8))
	v4515 = int32(0)
	F_vac_update_relstats(m, v4512, v4513, v4514, v4515, v4515, v4515, v4515, v4515, v4515, v4515, v4515)
	mBase = m.M
	v4524 = m.ExcPending
	if v4524 != 0 {
		goto L24
	} else {
		goto L640
	}
L640:
	;
	goto L637
L641:
	;
	goto L636
L642:
	;
	v4584 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v231)+25)))
	if v4584 != int32(1) {
		goto L643
	} else {
		goto L644
	}
L643:
	;
	v5556 = *(*int32)(unsafe.Add(mBase, uint32(v53)+636))
	*(*int32)(unsafe.Add(mBase, _c_F_heap_vacuum_rel[13])) = v5556
	v5562 = *(*int32)(unsafe.Add(mBase, _c_F_heap_vacuum_rel[9]))
	if v5562 == int32(0) {
		goto L785
	} else {
		goto L786
	}
L644:
	;
	v4588 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_heap_vacuum_rel[14])))
	if v4588&int32(1) != 0 {
		goto L643
	} else {
		goto L645
	}
L645:
	;
	v4591 = *(*int32)(unsafe.Add(mBase, uint32(v231)+108))
	v4592 = *(*int32)(unsafe.Add(mBase, uint32(v231)+148))
	if v4591 == v4592 {
		goto L643
	} else {
		goto L646
	}
L646:
	;
	v4594 = v4591 - v4592
	if base.B2i32(base.Ui32(v4594) <= base.Ui32(int32(999)))&base.B2i32(base.Ui32(v4594) < base.Ui32(int32(base.Ui32(v4591)>>(uint(int32(4))%32)))) != 0 {
		goto L643
	} else {
		goto L647
	}
L647:
	;
	v4605 = *(*int32)(unsafe.Add(mBase, _c_F_heap_vacuum_rel[9]))
	if v4605 == int32(0) {
		goto L649
	} else {
		goto L650
	}
L648:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v231)+92)) = int32(5)
	v4648 = int32(0)
	*(*uint16)(unsafe.Add(mBase, uint32(v231)+88)) = uint16(v4648)
	v4650 = *(*int32)(unsafe.Add(mBase, uint32(v231)+148))
	*(*int32)(unsafe.Add(mBase, uint32(v231)+84)) = v4650
	v4662 = v4591
	goto L652
L649:
	;
	goto L648
L650:
	;
	v4609 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_heap_vacuum_rel[10])))
	if v4609&int32(1) == int32(0) {
		goto L649
	} else {
		goto L651
	}
L651:
	;
	v4614 = int32(_a_F_heap_vacuum_rel_1)
	v4616 = *(*int32)(unsafe.Add(mBase, _c_F_heap_vacuum_rel[11]))
	v4617 = int32(1)
	*(*int32)(unsafe.Add(mBase, _c_F_heap_vacuum_rel[11])) = v4616 + v4617
	v4620 = *(*int32)(unsafe.Add(mBase, uint32(v4605)))
	*(*int32)(unsafe.Add(mBase, uint32(v4605))) = v4620 + v4617
	v4624 = int32(0)
	v4626 = int32(_a_F_heap_vacuum_rel_2)
	v4627 = base.AtomicRmwOr32(m, v4624, v4626, v4624)
	*(*int64)(unsafe.Add(mBase, uint32(v4605+v4624)+232)) = int64(5)
	v4635 = base.AtomicRmwOr32(m, v4624, v4626, v4624)
	v4636 = *(*int32)(unsafe.Add(mBase, uint32(v4605)))
	*(*int32)(unsafe.Add(mBase, uint32(v4605))) = v4636 + v4617
	v4642 = *(*int32)(unsafe.Add(mBase, _c_F_heap_vacuum_rel[11]))
	*(*int32)(unsafe.Add(mBase, _c_F_heap_vacuum_rel[11])) = v4642 - v4617
	goto L649
L652:
	;
	v4703 = *(*int32)(unsafe.Add(mBase, uint32(v231)))
	v4704 = F_ConditionalLockRelation(m, v4703)
	mBase = m.M
	v4705 = m.ExcPending
	if v4705 != 0 {
		goto L24
	} else {
		goto L654
	}
L653:
	;
	goto L643
L654:
	;
	if v4704 == int32(0) {
		goto L655
	} else {
		goto L656
	}
L655:
	;
	v4710 = int32(0)
	goto L658
L656:
	;
	goto L657
L657:
	;
	v4857 = *(*int32)(unsafe.Add(mBase, uint32(v231)))
	v4859 = F_RelationGetNumberOfBlocksInFork(m, v4857, int32(0))
	mBase = m.M
	v4860 = m.ExcPending
	if v4860 != 0 {
		goto L24
	} else {
		goto L678
	}
L658:
	;
	v4759 = *(*int32)(unsafe.Add(mBase, _c_F_heap_vacuum_rel[35]))
	if v4759 != 0 {
		goto L660
	} else {
		goto L661
	}
L659:
	;
	goto L657
L660:
	;
	F_ProcessInterrupts(m)
	mBase = m.M
	v4761 = m.ExcPending
	if v4761 != 0 {
		goto L24
	} else {
		goto L663
	}
L661:
	;
	goto L662
L662:
	;
	if v4710 == int32(100) {
		goto L664
	} else {
		goto L665
	}
L663:
	;
	goto L662
L664:
	;
	v4766 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v231)+96)))
	if v4766 != 0 {
		goto L667
	} else {
		goto L668
	}
L665:
	;
	goto L666
L666:
	;
	v4786 = *(*int32)(unsafe.Add(mBase, _c_F_heap_vacuum_rel[36]))
	v4790 = F_WaitLatch(m, v4786, int32(41), int32(50), int32(150994953))
	mBase = m.M
	v4791 = m.ExcPending
	if v4791 != 0 {
		goto L24
	} else {
		goto L674
	}
L667:
	;
	v4767 = int32(17)
	goto L669
L668:
	;
	v4767 = int32(13)
	goto L669
L669:
	;
	v4769 = F_errstart(m, v4767, int32(0))
	mBase = m.M
	v4770 = m.ExcPending
	if v4770 != 0 {
		goto L24
	} else {
		goto L670
	}
L670:
	;
	if v4769 == int32(0) {
		goto L643
	} else {
		goto L671
	}
L671:
	;
	v4773 = *(*int32)(unsafe.Add(mBase, uint32(v231)+76))
	*(*int32)(unsafe.Add(mBase, uint32(v53)+512)) = v4773
	F_errmsg(m, int32(_a_F_heap_vacuum_rel_25), v53+int32(512))
	mBase = m.M
	v4779 = m.ExcPending
	if v4779 != 0 {
		goto L24
	} else {
		goto L672
	}
L672:
	;
	F_errfinish(m, int32(_a_F_heap_vacuum_rel_7), int32(3215), int32(_a_F_heap_vacuum_rel_26))
	mBase = m.M
	v4784 = m.ExcPending
	if v4784 != 0 {
		goto L24
	} else {
		goto L673
	}
L673:
	;
	goto L643
L674:
	;
	v4793 = *(*int32)(unsafe.Add(mBase, _c_F_heap_vacuum_rel[36]))
	v4794 = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v4793))) = v4794
	v4799 = base.AtomicRmwOr32(m, v4794, int32(_a_F_heap_vacuum_rel_27), v4794)
	goto L675
L675:
	;
	v4802 = *(*int32)(unsafe.Add(mBase, uint32(v231)))
	v4803 = F_ConditionalLockRelation(m, v4802)
	mBase = m.M
	v4804 = m.ExcPending
	if v4804 != 0 {
		goto L24
	} else {
		goto L676
	}
L676:
	;
	if v4803 == int32(0) {
		v4710 = v4710 + int32(1)
		goto L658
	} else {
		goto L677
	}
L677:
	;
	goto L659
L678:
	;
	if v4859 != v4662 {
		goto L679
	} else {
		goto L680
	}
L679:
	;
	v4862 = *(*int32)(unsafe.Add(mBase, uint32(v231)))
	F_UnlockRelation(m, v4862)
	mBase = m.M
	v4864 = m.ExcPending
	if v4864 != 0 {
		goto L24
	} else {
		goto L682
	}
L680:
	;
	goto L681
L681:
	;
	F___clock_gettime(m, int32(1), v53+int32(1040))
	mBase = m.M
	v4869 = int32(0)
	v4870 = *(*int32)(unsafe.Add(mBase, uint32(v231)+108))
	v4871 = *(*int32)(unsafe.Add(mBase, uint32(v231)+148))
	if base.Ui32(v4870) <= base.Ui32(v4871) {
		v5421 = v4871
		v5426 = v4869
		goto L683
	} else {
		goto L684
	}
L682:
	;
	goto L643
L683:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v231)+84)) = v5421
	v5466 = *(*int32)(unsafe.Add(mBase, uint32(v231)))
	if base.Ui32(v4662) <= base.Ui32(v5421) {
		goto L768
	} else {
		goto L769
	}
L684:
	;
	v4873 = int64(*(*int32)(unsafe.Add(mBase, uint32(v53)+1048)))
	v4874 = *(*int64)(unsafe.Add(mBase, uint32(v53)+1040))
	v4884 = v4870
	v4891 = int32(-1)
	v4918 = v4873 + v4874*int64(1000000000)
	goto L685
L685:
	;
	if v4884&int32(31) != 0 {
		v5126 = v4918
		goto L687
	} else {
		goto L688
	}
L686:
	;
	v5421 = v5413
	v5426 = v4869
	goto L683
L687:
	;
	v5129 = *(*int32)(unsafe.Add(mBase, _c_F_heap_vacuum_rel[35]))
	if v5129 != 0 {
		goto L734
	} else {
		goto L735
	}
L688:
	;
	F___clock_gettime(m, int32(1), v53+int32(1040))
	mBase = m.M
	v4935 = int64(*(*int32)(unsafe.Add(mBase, uint32(v53)+1048)))
	v4936 = *(*int64)(unsafe.Add(mBase, uint32(v53)+1040))
	v4939 = v4935 + v4936*int64(1000000000)
	if v4939-v4918 < int64(20000000) {
		v5126 = v4918
		goto L687
	} else {
		goto L689
	}
L689:
	;
	v4943 = *(*int32)(unsafe.Add(mBase, uint32(v231)))
	v4944 = m.G0
	v4946 = v4944 - int32(16)
	m.G0 = v4946
	v4948 = *(*int32)(unsafe.Add(mBase, uint32(v4943)+64))
	*(*int32)(unsafe.Add(mBase, uint32(v4946))) = v4948
	v4950 = *(*int32)(unsafe.Add(mBase, uint32(v4943)+60))
	*(*int64)(unsafe.Add(mBase, uint32(v4946)+8)) = int64(72057594037927936)
	*(*int32)(unsafe.Add(mBase, uint32(v4946)+4)) = v4950
	v4954 = m.G0
	v4956 = v4954 - int32(80)
	m.G0 = v4956
	v4958 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v4946)+15)))
	if base.Ui32(int32(253)) < base.Ui32((v4958-int32(3))&int32(255)) {
		goto L692
	} else {
		goto L693
	}
L690:
	;
	m.G0 = v4946 + int32(16)
	if v5061 == int32(0) {
		v5126 = v4939
		goto L687
	} else {
		goto L726
	}
L691:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v5081 = m.ExcPending
	if v5081 != 0 {
		goto L24
	} else {
		goto L723
	}
L692:
	;
	v4967 = *(*int32)(unsafe.Add(mBase, uint32(v4958<<(uint(int32(2))%32))+uint32(_c_F_heap_vacuum_rel[37])))
	v4968 = *(*int32)(unsafe.Add(mBase, uint32(v4967)))
	if v4968 < int32(8) {
		goto L691
	} else {
		goto L695
	}
L693:
	;
	goto L694
L694:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v5068 = m.ExcPending
	if v5068 != 0 {
		goto L24
	} else {
		goto L720
	}
L695:
	;
	v4971 = *(*int64)(unsafe.Add(mBase, uint32(v4946)+8))
	*(*int64)(unsafe.Add(mBase, uint32(v4956)+64)) = v4971
	v4973 = *(*int64)(unsafe.Add(mBase, uint32(v4946)))
	*(*int64)(unsafe.Add(mBase, uint32(v4956)+56)) = v4973
	*(*int32)(unsafe.Add(mBase, uint32(v4956)+72)) = int32(8)
	v4977 = int32(0)
	v4979 = *(*int32)(unsafe.Add(mBase, _c_F_heap_vacuum_rel[38]))
	v4984 = F_hash_search(m, v4979, v4956+int32(56), v4977, v4977)
	mBase = m.M
	v4985 = m.ExcPending
	if v4985 != 0 {
		goto L24
	} else {
		goto L698
	}
L696:
	;
	m.G0 = v4956 + int32(80)
	goto L690
L697:
	;
	v5009 = *(*int32)(unsafe.Add(mBase, _c_F_heap_vacuum_rel[39]))
	v5010 = *(*int32)(unsafe.Add(mBase, uint32(v4984)+20))
	v5017 = v5009 + v5010&int32(15)<<(uint(int32(7))%32) + int32(_a_F_heap_vacuum_rel_28)
	v5019 = F_LWLockAcquire(m, v5017, int32(1))
	mBase = m.M
	v5020 = m.ExcPending
	if v5020 != 0 {
		goto L24
	} else {
		goto L707
	}
L698:
	;
	if v4984 != 0 {
		goto L699
	} else {
		goto L700
	}
L699:
	;
	v4986 = *(*int64)(unsafe.Add(mBase, uint32(v4984)+32))
	if int64(0) < v4986 {
		goto L697
	} else {
		goto L702
	}
L700:
	;
	goto L701
L701:
	;
	v4991 = F_errstart(m, int32(19), int32(0))
	mBase = m.M
	v4992 = m.ExcPending
	if v4992 != 0 {
		goto L24
	} else {
		goto L703
	}
L702:
	;
	goto L701
L703:
	;
	if v4991 == int32(0) {
		v5061 = v4977
		goto L696
	} else {
		goto L704
	}
L704:
	;
	v4995 = *(*int32)(unsafe.Add(mBase, uint32(v4967)+8))
	v4996 = *(*int32)(unsafe.Add(mBase, uint32(v4995)+32))
	*(*int32)(unsafe.Add(mBase, uint32(v4956)+32)) = v4996
	F_errmsg_internal(m, int32(_a_F_heap_vacuum_rel_29), v4956+int32(32))
	mBase = m.M
	v5002 = m.ExcPending
	if v5002 != 0 {
		goto L24
	} else {
		goto L705
	}
L705:
	;
	F_errfinish(m, int32(_a_F_heap_vacuum_rel_30), int32(766), int32(_a_F_heap_vacuum_rel_31))
	mBase = m.M
	v5007 = m.ExcPending
	if v5007 != 0 {
		goto L24
	} else {
		goto L706
	}
L706:
	;
	v5061 = v4977
	goto L696
L707:
	;
	v5021 = *(*int32)(unsafe.Add(mBase, uint32(v4984)+28))
	v5022 = *(*int32)(unsafe.Add(mBase, uint32(v5021)+12))
	if int32(base.Ui32(v5022)>>(uint(int32(8))%32))&int32(1) == int32(0) {
		goto L708
	} else {
		goto L709
	}
L708:
	;
	F_LWLockRelease(m, v5017)
	mBase = m.M
	v5030 = m.ExcPending
	if v5030 != 0 {
		goto L24
	} else {
		goto L711
	}
L709:
	;
	goto L710
L710:
	;
	v5051 = *(*int32)(unsafe.Add(mBase, uint32(v4967)+4))
	v5052 = *(*int32)(unsafe.Add(mBase, uint32(v5051)+32))
	v5053 = *(*int32)(unsafe.Add(mBase, uint32(v4984)+24))
	v5054 = *(*int32)(unsafe.Add(mBase, uint32(v5053)+20))
	F_LWLockRelease(m, v5017)
	mBase = m.M
	v5056 = m.ExcPending
	if v5056 != 0 {
		goto L24
	} else {
		goto L719
	}
L711:
	;
	v5033 = F_errstart(m, int32(19), int32(0))
	mBase = m.M
	v5034 = m.ExcPending
	if v5034 != 0 {
		goto L24
	} else {
		goto L712
	}
L712:
	;
	if v5033 != 0 {
		goto L713
	} else {
		goto L714
	}
L713:
	;
	v5035 = *(*int32)(unsafe.Add(mBase, uint32(v4967)+8))
	v5036 = *(*int32)(unsafe.Add(mBase, uint32(v5035)+32))
	*(*int32)(unsafe.Add(mBase, uint32(v4956)+48)) = v5036
	F_errmsg_internal(m, int32(_a_F_heap_vacuum_rel_29), v4956+int32(48))
	mBase = m.M
	v5042 = m.ExcPending
	if v5042 != 0 {
		goto L24
	} else {
		goto L716
	}
L714:
	;
	goto L715
L715:
	;
	F_RemoveLocalLock(m, v4984)
	mBase = m.M
	v5049 = m.ExcPending
	if v5049 != 0 {
		goto L24
	} else {
		goto L718
	}
L716:
	;
	F_errfinish(m, int32(_a_F_heap_vacuum_rel_30), int32(796), int32(_a_F_heap_vacuum_rel_31))
	mBase = m.M
	v5047 = m.ExcPending
	if v5047 != 0 {
		goto L24
	} else {
		goto L717
	}
L717:
	;
	goto L715
L718:
	;
	v5061 = int32(0)
	goto L696
L719:
	;
	v5061 = base.B2i32(v5054&v5052 != int32(0))
	goto L696
L720:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v4956))) = v4958
	F_errmsg_internal(m, int32(_a_F_heap_vacuum_rel_32), v4956)
	mBase = m.M
	v5072 = m.ExcPending
	if v5072 != 0 {
		goto L24
	} else {
		goto L721
	}
L721:
	;
	F_errfinish(m, int32(_a_F_heap_vacuum_rel_30), int32(737), int32(_a_F_heap_vacuum_rel_31))
	mBase = m.M
	v5077 = m.ExcPending
	if v5077 != 0 {
		goto L24
	} else {
		goto L722
	}
L722:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L723:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v4956)+16)) = int32(8)
	F_errmsg_internal(m, int32(_a_F_heap_vacuum_rel_33), v4956+int32(16))
	mBase = m.M
	v5088 = m.ExcPending
	if v5088 != 0 {
		goto L24
	} else {
		goto L724
	}
L724:
	;
	F_errfinish(m, int32(_a_F_heap_vacuum_rel_30), int32(740), int32(_a_F_heap_vacuum_rel_31))
	mBase = m.M
	v5093 = m.ExcPending
	if v5093 != 0 {
		goto L24
	} else {
		goto L725
	}
L725:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L726:
	;
	v5099 = int32(1)
	v5102 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v231)+96)))
	if v5102 != 0 {
		goto L727
	} else {
		goto L728
	}
L727:
	;
	v5103 = int32(17)
	goto L729
L728:
	;
	v5103 = int32(13)
	goto L729
L729:
	;
	v5105 = F_errstart(m, v5103, int32(0))
	mBase = m.M
	v5106 = m.ExcPending
	if v5106 != 0 {
		goto L24
	} else {
		goto L730
	}
L730:
	;
	if v5105 == int32(0) {
		v5421 = v4884
		v5426 = v5099
		goto L683
	} else {
		goto L731
	}
L731:
	;
	v5109 = *(*int32)(unsafe.Add(mBase, uint32(v231)+76))
	*(*int32)(unsafe.Add(mBase, uint32(v53)+496)) = v5109
	F_errmsg(m, int32(_a_F_heap_vacuum_rel_34), v53+int32(496))
	mBase = m.M
	v5115 = m.ExcPending
	if v5115 != 0 {
		goto L24
	} else {
		goto L732
	}
L732:
	;
	F_errfinish(m, int32(_a_F_heap_vacuum_rel_7), int32(3348), int32(_a_F_heap_vacuum_rel_35))
	mBase = m.M
	v5120 = m.ExcPending
	if v5120 != 0 {
		goto L24
	} else {
		goto L733
	}
L733:
	;
	v5421 = v4884
	v5426 = v5099
	goto L683
L734:
	;
	F_ProcessInterrupts(m)
	mBase = m.M
	v5131 = m.ExcPending
	if v5131 != 0 {
		goto L24
	} else {
		goto L737
	}
L735:
	;
	goto L736
L736:
	;
	v5133 = v4884 - int32(1)
	if base.Ui32(v5133) < base.Ui32(v4891) {
		goto L738
	} else {
		goto L739
	}
L737:
	;
	goto L736
L738:
	;
	v5136 = v5133 & int32(-32)
	v5139 = v5136
	goto L741
L739:
	;
	v5212 = v4891
	goto L740
L740:
	;
	v5250 = *(*int32)(unsafe.Add(mBase, uint32(v231)))
	v5251 = int32(0)
	v5253 = *(*int32)(unsafe.Add(mBase, uint32(v231)+12))
	v5254 = F_ReadBufferExtended(m, v5250, v5251, v5133, v5251, v5253)
	mBase = m.M
	v5255 = m.ExcPending
	if v5255 != 0 {
		goto L24
	} else {
		goto L749
	}
L741:
	;
	v5189 = *(*int32)(unsafe.Add(mBase, uint32(v231)))
	F_PrefetchBuffer(m, v53+int32(1040), v5189, int32(0), v5139)
	mBase = m.M
	v5192 = m.ExcPending
	if v5192 != 0 {
		goto L24
	} else {
		goto L743
	}
L742:
	;
	v5212 = v5136
	goto L740
L743:
	;
	v5194 = *(*int32)(unsafe.Add(mBase, _c_F_heap_vacuum_rel[35]))
	if v5194 != 0 {
		goto L744
	} else {
		goto L745
	}
L744:
	;
	F_ProcessInterrupts(m)
	mBase = m.M
	v5196 = m.ExcPending
	if v5196 != 0 {
		goto L24
	} else {
		goto L747
	}
L745:
	;
	goto L746
L746:
	;
	if base.Ui32(v5139) < base.Ui32(v5133) {
		v5139 = v5139 + int32(1)
		goto L741
	} else {
		goto L748
	}
L747:
	;
	goto L746
L748:
	;
	goto L742
L749:
	;
	F_LockBufferInternal(m, v5254, int32(1))
	mBase = m.M
	v5258 = m.ExcPending
	if v5258 != 0 {
		goto L24
	} else {
		goto L750
	}
L750:
	;
	if v5254 < int32(0) {
		goto L753
	} else {
		goto L754
	}
L751:
	;
	F_UnlockReleaseBuffer(m, v5254)
	mBase = m.M
	v5412 = m.ExcPending
	if v5412 != 0 {
		goto L24
	} else {
		goto L766
	}
L752:
	;
	v5277 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v5276)+14)))
	if v5277 == int32(0) {
		goto L751
	} else {
		goto L756
	}
L753:
	;
	v5262 = *(*int32)(unsafe.Add(mBase, _c_F_heap_vacuum_rel[30]))
	v5268 = *(*int32)(unsafe.Add(mBase, uint32(v5262+(v5254^int32(-1))<<(uint(int32(2))%32))))
	v5276 = v5268
	goto L752
L754:
	;
	goto L755
L755:
	;
	v5270 = *(*int32)(unsafe.Add(mBase, _c_F_heap_vacuum_rel[31]))
	v5276 = v5270 + v5254<<(uint(int32(13))%32) + int32(-8192)
	goto L752
L756:
	;
	v5280 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v5276)+12)))
	if base.Ui32(v5280) < base.Ui32(int32(25)) {
		goto L751
	} else {
		goto L757
	}
L757:
	;
	v5288 = int32(base.Ui32(v5280+int32(_a_F_heap_vacuum_rel_16))>>(uint(int32(2))%32)) & int32(_a_F_heap_vacuum_rel_17)
	if v5288 == int32(0) {
		goto L751
	} else {
		goto L758
	}
L758:
	;
	v5296 = int32(1)
	goto L759
L759:
	;
	v5349 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v5276+int32(20)+v5296&int32(_a_F_heap_vacuum_rel_17)<<(uint(int32(2))%32))+1)))
	if v5349&int32(384) == int32(0) {
		goto L761
	} else {
		goto L762
	}
L760:
	;
	F_UnlockReleaseBuffer(m, v5254)
	mBase = m.M
	v5360 = m.ExcPending
	if v5360 != 0 {
		goto L24
	} else {
		goto L765
	}
L761:
	;
	v5355 = v5296 + int32(1)
	if base.Ui32(v5355&int32(_a_F_heap_vacuum_rel_17)) <= base.Ui32(v5288) {
		v5296 = v5355
		goto L759
	} else {
		goto L764
	}
L762:
	;
	goto L763
L763:
	;
	goto L760
L764:
	;
	goto L751
L765:
	;
	v5421 = v4884
	v5426 = v4869
	goto L683
L766:
	;
	v5413 = *(*int32)(unsafe.Add(mBase, uint32(v231)+148))
	if base.Ui32(v5413) < base.Ui32(v5133) {
		v4884 = v5133
		v4891 = v5212
		v4918 = v5126
		goto L685
	} else {
		goto L767
	}
L767:
	;
	goto L686
L768:
	;
	F_UnlockRelation(m, v5466)
	mBase = m.M
	v5469 = m.ExcPending
	if v5469 != 0 {
		goto L24
	} else {
		goto L771
	}
L769:
	;
	goto L770
L770:
	;
	F_RelationTruncate(m, v5466, v5421)
	mBase = m.M
	v5471 = m.ExcPending
	if v5471 != 0 {
		goto L24
	} else {
		goto L772
	}
L771:
	;
	goto L643
L772:
	;
	v5472 = *(*int32)(unsafe.Add(mBase, uint32(v231)))
	F_UnlockRelation(m, v5472)
	mBase = m.M
	v5474 = m.ExcPending
	if v5474 != 0 {
		goto L24
	} else {
		goto L773
	}
L773:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v231)+108)) = v5421
	v5476 = *(*int32)(unsafe.Add(mBase, uint32(v231)+120))
	*(*int32)(unsafe.Add(mBase, uint32(v231)+120)) = v5476 + (v4662 - v5421)
	v5482 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v231)+96)))
	if v5482 != 0 {
		goto L774
	} else {
		goto L775
	}
L774:
	;
	v5483 = int32(17)
	goto L776
L775:
	;
	v5483 = int32(13)
	goto L776
L776:
	;
	v5485 = F_errstart(m, v5483, int32(0))
	mBase = m.M
	v5486 = m.ExcPending
	if v5486 != 0 {
		goto L24
	} else {
		goto L777
	}
L777:
	;
	if v5485 != 0 {
		goto L778
	} else {
		goto L779
	}
L778:
	;
	v5487 = *(*int32)(unsafe.Add(mBase, uint32(v231)+76))
	*(*int32)(unsafe.Add(mBase, uint32(v53)+488)) = v5421
	*(*int32)(unsafe.Add(mBase, uint32(v53)+484)) = v4662
	*(*int32)(unsafe.Add(mBase, uint32(v53)+480)) = v5487
	F_errmsg(m, int32(_a_F_heap_vacuum_rel_36), v53+int32(480))
	mBase = m.M
	v5495 = m.ExcPending
	if v5495 != 0 {
		goto L24
	} else {
		goto L781
	}
L779:
	;
	goto L780
L780:
	;
	v5502 = *(*int32)(unsafe.Add(mBase, uint32(v231)+148))
	if v5426&base.B2i32(base.Ui32(v5502) < base.Ui32(v5421)) != 0 {
		v4662 = v5421
		goto L652
	} else {
		goto L783
	}
L781:
	;
	F_errfinish(m, int32(_a_F_heap_vacuum_rel_7), int32(3286), int32(_a_F_heap_vacuum_rel_26))
	mBase = m.M
	v5500 = m.ExcPending
	if v5500 != 0 {
		goto L24
	} else {
		goto L782
	}
L782:
	;
	goto L780
L783:
	;
	goto L653
L784:
	;
	v5603 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v231)+64)))
	if v5603 == int32(1) {
		goto L788
	} else {
		goto L789
	}
L785:
	;
	goto L784
L786:
	;
	v5566 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_heap_vacuum_rel[10])))
	if v5566&int32(1) == int32(0) {
		goto L785
	} else {
		goto L787
	}
L787:
	;
	v5571 = int32(_a_F_heap_vacuum_rel_1)
	v5573 = *(*int32)(unsafe.Add(mBase, _c_F_heap_vacuum_rel[11]))
	v5574 = int32(1)
	*(*int32)(unsafe.Add(mBase, _c_F_heap_vacuum_rel[11])) = v5573 + v5574
	v5577 = *(*int32)(unsafe.Add(mBase, uint32(v5562)))
	*(*int32)(unsafe.Add(mBase, uint32(v5562))) = v5577 + v5574
	v5581 = int32(0)
	v5583 = int32(_a_F_heap_vacuum_rel_2)
	v5584 = base.AtomicRmwOr32(m, v5581, v5583, v5581)
	*(*int64)(unsafe.Add(mBase, uint32(v5562+v5581)+232)) = int64(6)
	v5592 = base.AtomicRmwOr32(m, v5581, v5583, v5581)
	v5593 = *(*int32)(unsafe.Add(mBase, uint32(v5562)))
	*(*int32)(unsafe.Add(mBase, uint32(v5562))) = v5593 + v5574
	v5599 = *(*int32)(unsafe.Add(mBase, _c_F_heap_vacuum_rel[11]))
	*(*int32)(unsafe.Add(mBase, _c_F_heap_vacuum_rel[11])) = v5599 - v5574
	goto L785
L788:
	;
	v5606 = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v1442))) = v5606
	*(*int32)(unsafe.Add(mBase, uint32(v1440))) = v5606
	goto L790
L789:
	;
	goto L790
L790:
	;
	v5610 = *(*int32)(unsafe.Add(mBase, uint32(v231)+108))
	F_visibilitymap_count(m, l0, v53+int32(1664), v53+int32(1024))
	mBase = m.M
	v5616 = m.ExcPending
	if v5616 != 0 {
		goto L24
	} else {
		goto L791
	}
L791:
	;
	v5617 = *(*int32)(unsafe.Add(mBase, uint32(v53)+1664))
	if base.Ui32(v5610) < base.Ui32(v5617) {
		goto L792
	} else {
		goto L793
	}
L792:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v53)+1664)) = v5610
	v5620 = v5610
	goto L794
L793:
	;
	v5620 = v5617
	goto L794
L794:
	;
	v5621 = *(*int32)(unsafe.Add(mBase, uint32(v53)+1024))
	if base.Ui32(v5620) < base.Ui32(v5621) {
		goto L795
	} else {
		goto L796
	}
L795:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v53)+1024)) = v5620
	v5624 = v5620
	goto L797
L796:
	;
	v5624 = v5621
	goto L797
L797:
	;
	v5625 = int32(0)
	v5626 = *(*float64)(unsafe.Add(mBase, uint32(v231)+160))
	v5627 = *(*int32)(unsafe.Add(mBase, uint32(v231)+8))
	v5630 = *(*int32)(unsafe.Add(mBase, uint32(v231)+56))
	v5631 = *(*int32)(unsafe.Add(mBase, uint32(v231)+60))
	F_vac_update_relstats(m, l0, v5610, v5626, v5620, v5624, base.B2i32(v5625 < v5627), v5630, v5631, v53+int32(1036), v53+int32(988), v5625)
	mBase = m.M
	v5638 = m.ExcPending
	if v5638 != 0 {
		goto L24
	} else {
		goto L798
	}
L798:
	;
	v5639 = *(*float64)(unsafe.Add(mBase, uint32(v231)+160))
	v5640 = float64(0)
	if base.F64_gt(v5639, v5640) != 0 {
		goto L799
	} else {
		goto L800
	}
L799:
	;
	v5643 = v5639
	goto L801
L800:
	;
	v5643 = v5640
	goto L801
L801:
	;
	v5645 = *(*int64)(unsafe.Add(mBase, uint32(v231)+240))
	v5646 = *(*int64)(unsafe.Add(mBase, uint32(v231)+232))
	v5649 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_heap_vacuum_rel[40])))
	if v5649 == int32(1) {
		goto L802
	} else {
		goto L803
	}
L802:
	;
	v5653 = *(*int32)(unsafe.Add(mBase, _c_F_heap_vacuum_rel[12]))
	v5654 = *(*int32)(unsafe.Add(mBase, uint32(l0)+48))
	v5655 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v5654)+117)))
	v5659 = m.G0
	v5660 = int32(16)
	v5661 = v5659 - v5660
	m.G0 = v5661
	F_gettimeofday(m, v5661)
	mBase = m.M
	v5664 = *(*int64)(unsafe.Add(mBase, uint32(v5661)))
	v5665 = int64(*(*int32)(unsafe.Add(mBase, uint32(v5661)+8)))
	m.G0 = v5661 + v5660
	v5673 = v5665 + v5664*int64(1000000) - int64(946684800000000)
	goto L805
L803:
	;
	goto L804
L804:
	;
	v5746 = *(*int32)(unsafe.Add(mBase, _c_F_heap_vacuum_rel[9]))
	if v5746 == int32(0) {
		goto L827
	} else {
		goto L828
	}
L805:
	;
	if v5673 <= v127 {
		v5691 = int32(0)
		goto L807
	} else {
		goto L808
	}
L806:
	;
	if v5655 != 0 {
		goto L810
	} else {
		goto L811
	}
L807:
	;
	goto L806
L808:
	;
	v5679 = v5673 - v127
	if base.B2i32(int64(0) < v127)^base.B2i32(v5679 < v5673)|base.B2i32(int64(2147483646000) < v5679) != 0 {
		v5691 = int32(2147483647)
		goto L807
	} else {
		goto L809
	}
L809:
	;
	v5688 = base.I64_div_s(v5679+int64(999), int64(1000))
	v5691 = base.I32_wrap_i64(v5688)
	goto L807
L810:
	;
	v5694 = int32(0)
	goto L812
L811:
	;
	v5694 = v5653
	goto L812
L812:
	;
	v5695 = int64(*(*uint32)(unsafe.Add(mBase, uint32(l0)+56)))
	v5697 = F_pgstat_get_entry_ref_locked(m, int32(2), v5694, v5695, int32(0))
	mBase = m.M
	v5698 = m.ExcPending
	if v5698 != 0 {
		goto L24
	} else {
		goto L813
	}
L813:
	;
	v5699 = *(*int32)(unsafe.Add(mBase, uint32(v5697)+4))
	*(*int64)(unsafe.Add(mBase, uint32(v5699)+120)) = int64(0)
	*(*int64)(unsafe.Add(mBase, uint32(v5699)+104)) = v5645 + v5646
	*(*int64)(unsafe.Add(mBase, uint32(v5699)+96)) = base.I64_trunc_sat_f64_s(v5643)
	v5707 = *(*int32)(unsafe.Add(mBase, _c_F_heap_vacuum_rel[5]))
	v5709 = base.B2i32(v5707 == int32(4))
	if v5707 == int32(4) {
		goto L814
	} else {
		goto L815
	}
L814:
	;
	v5710 = int32(160)
	goto L816
L815:
	;
	v5710 = int32(144)
	goto L816
L816:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v5699+v5710))) = v5673
	if v5707 == int32(4) {
		goto L817
	} else {
		goto L818
	}
L817:
	;
	v5715 = int32(168)
	goto L819
L818:
	;
	v5715 = int32(152)
	goto L819
L819:
	;
	v5716 = v5699 + v5715
	v5717 = *(*int64)(unsafe.Add(mBase, uint32(v5716)))
	*(*int64)(unsafe.Add(mBase, uint32(v5716))) = v5717 + int64(1)
	if v5707 == int32(4) {
		goto L820
	} else {
		goto L821
	}
L820:
	;
	v5723 = int32(216)
	goto L822
L821:
	;
	v5723 = int32(208)
	goto L822
L822:
	;
	v5724 = v5699 + v5723
	v5725 = *(*int64)(unsafe.Add(mBase, uint32(v5724)))
	*(*int64)(unsafe.Add(mBase, uint32(v5724))) = v5725 + base.I64_extend_i32_s(v5691)
	F_pgstat_unlock_entry(m, v5697)
	mBase = m.M
	v5730 = m.ExcPending
	if v5730 != 0 {
		goto L24
	} else {
		goto L823
	}
L823:
	;
	F_pgstat_flush_io(m, int32(0))
	mBase = m.M
	v5733 = m.ExcPending
	if v5733 != 0 {
		goto L24
	} else {
		goto L824
	}
L824:
	;
	v5736 = F_pgstat_flush_backend(m, int32(0), int32(1))
	mBase = m.M
	v5737 = m.ExcPending
	if v5737 != 0 {
		goto L24
	} else {
		goto L825
	}
L825:
	;
	goto L804
L826:
	;
	if v105 == int32(0) {
		goto L831
	} else {
		goto L832
	}
L827:
	;
	goto L826
L828:
	;
	v5750 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_heap_vacuum_rel[10])))
	if v5750&int32(1) == int32(0) {
		goto L827
	} else {
		goto L829
	}
L829:
	;
	v5755 = *(*int32)(unsafe.Add(mBase, uint32(v5746)+220))
	if v5755 == int32(0) {
		goto L827
	} else {
		goto L830
	}
L830:
	;
	v5758 = int32(_a_F_heap_vacuum_rel_1)
	v5760 = *(*int32)(unsafe.Add(mBase, _c_F_heap_vacuum_rel[11]))
	v5761 = int32(1)
	*(*int32)(unsafe.Add(mBase, _c_F_heap_vacuum_rel[11])) = v5760 + v5761
	v5764 = *(*int32)(unsafe.Add(mBase, uint32(v5746)))
	*(*int32)(unsafe.Add(mBase, uint32(v5746))) = v5764 + v5761
	v5768 = int32(0)
	v5770 = int32(_a_F_heap_vacuum_rel_2)
	v5771 = base.AtomicRmwOr32(m, v5768, v5770, v5768)
	*(*int32)(unsafe.Add(mBase, uint32(v5746)+220)) = v5768
	*(*int32)(unsafe.Add(mBase, uint32(v5746)+224)) = v5768
	v5779 = base.AtomicRmwOr32(m, v5768, v5770, v5768)
	v5780 = *(*int32)(unsafe.Add(mBase, uint32(v5746)))
	*(*int32)(unsafe.Add(mBase, uint32(v5746))) = v5780 + v5761
	v5786 = *(*int32)(unsafe.Add(mBase, _c_F_heap_vacuum_rel[11]))
	*(*int32)(unsafe.Add(mBase, _c_F_heap_vacuum_rel[11])) = v5786 - v5761
	goto L827
L831:
	;
	v6570 = *(*int32)(unsafe.Add(mBase, uint32(v267)))
	if int32(0) < v6570 {
		goto L946
	} else {
		goto L947
	}
L832:
	;
	v5795 = m.G0
	v5796 = int32(16)
	v5797 = v5795 - v5796
	m.G0 = v5797
	F_gettimeofday(m, v5797)
	mBase = m.M
	v5800 = *(*int64)(unsafe.Add(mBase, uint32(v5797)))
	v5801 = int64(*(*int32)(unsafe.Add(mBase, uint32(v5797)+8)))
	m.G0 = v5797 + v5796
	v5809 = v5801 + v5800*int64(1000000) - int64(946684800000000)
	goto L833
L833:
	;
	if v77 != 0 {
		goto L834
	} else {
		goto L835
	}
L834:
	;
	v5827 = v5809 - v127
	if v5827 <= int64(0) {
		goto L840
	} else {
		goto L841
	}
L835:
	;
	v5810 = *(*int32)(unsafe.Add(mBase, uint32(l1)+24))
	if v5810 == int32(0) {
		goto L834
	} else {
		goto L836
	}
L836:
	;
	goto L837
L837:
	;
	if base.B2i32(base.I64_extend_i32_s(v5810)*int64(1000) <= v5809-v127) == int32(0) {
		goto L831
	} else {
		goto L838
	}
L838:
	;
	goto L834
L839:
	;
	v5843 = int64(0)
	*(*int64)(unsafe.Add(mBase, uint32(v53)+624)) = v5843
	*(*int64)(unsafe.Add(mBase, uint32(v53)+616)) = v5843
	*(*int64)(unsafe.Add(mBase, uint32(v53)+608)) = v5843
	*(*int64)(unsafe.Add(mBase, uint32(v53)+600)) = v5843
	*(*int64)(unsafe.Add(mBase, uint32(v53)+592)) = v5843
	v5854 = v53 + int32(592)
	v5856 = v53 + int32(776)
	v5857 = *(*int64)(unsafe.Add(mBase, uint32(v5854)+16))
	v5859 = *(*int64)(unsafe.Add(mBase, _c_F_heap_vacuum_rel[2]))
	v5860 = *(*int64)(unsafe.Add(mBase, uint32(v5856)+16))
	*(*int64)(unsafe.Add(mBase, uint32(v5854)+16)) = v5857 + (v5859 - v5860)
	v5864 = *(*int64)(unsafe.Add(mBase, uint32(v5854)))
	v5866 = *(*int64)(unsafe.Add(mBase, _c_F_heap_vacuum_rel[4]))
	v5867 = *(*int64)(unsafe.Add(mBase, uint32(v5856)))
	*(*int64)(unsafe.Add(mBase, uint32(v5854))) = v5864 + (v5866 - v5867)
	v5871 = *(*int64)(unsafe.Add(mBase, uint32(v5854)+8))
	v5873 = *(*int64)(unsafe.Add(mBase, _c_F_heap_vacuum_rel[3]))
	v5874 = *(*int64)(unsafe.Add(mBase, uint32(v5856)+8))
	*(*int64)(unsafe.Add(mBase, uint32(v5854)+8)) = v5871 + (v5873 - v5874)
	v5878 = *(*int64)(unsafe.Add(mBase, uint32(v5854)+24))
	v5880 = *(*int64)(unsafe.Add(mBase, _c_F_heap_vacuum_rel[1]))
	v5881 = *(*int64)(unsafe.Add(mBase, uint32(v5856)+24))
	*(*int64)(unsafe.Add(mBase, uint32(v5854)+24)) = v5878 + (v5880 - v5881)
	v5885 = *(*int64)(unsafe.Add(mBase, uint32(v5854)+32))
	v5887 = *(*int64)(unsafe.Add(mBase, _c_F_heap_vacuum_rel[0]))
	v5888 = *(*int64)(unsafe.Add(mBase, uint32(v5856)+32))
	*(*int64)(unsafe.Add(mBase, uint32(v5854)+32)) = v5885 + (v5887 - v5888)
	goto L843
L840:
	;
	v5839 = int32(0)
	v5840 = int32(0)
	goto L842
L841:
	;
	v5831 = int64(1000000)
	v5832 = base.I64_div_u_s(v5827, v5831)
	v5839 = base.I32_wrap_i64(v5832)
	v5840 = base.I32_wrap_i64(v5827 - v5832*v5831)
	goto L842
L842:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v53+int32(1688)))) = v5839
	*(*int32)(unsafe.Add(mBase, uint32(v53+int32(1680)))) = v5840
	goto L839
L843:
	;
	v5893 = v53 + int32(1040)
	base.MemoryFill(m, v5893, int32(0), int32(128))
	v5898 = v53 + int32(648)
	v5899 = *(*int64)(unsafe.Add(mBase, uint32(v5893)))
	v5901 = *(*int64)(unsafe.Add(mBase, _c_F_heap_vacuum_rel[41]))
	v5902 = *(*int64)(unsafe.Add(mBase, uint32(v5898)))
	*(*int64)(unsafe.Add(mBase, uint32(v5893))) = v5899 + (v5901 - v5902)
	v5906 = *(*int64)(unsafe.Add(mBase, uint32(v5893)+8))
	v5908 = *(*int64)(unsafe.Add(mBase, _c_F_heap_vacuum_rel[42]))
	v5909 = *(*int64)(unsafe.Add(mBase, uint32(v5898)+8))
	*(*int64)(unsafe.Add(mBase, uint32(v5893)+8)) = v5906 + (v5908 - v5909)
	v5913 = *(*int64)(unsafe.Add(mBase, uint32(v5893)+16))
	v5915 = *(*int64)(unsafe.Add(mBase, _c_F_heap_vacuum_rel[43]))
	v5916 = *(*int64)(unsafe.Add(mBase, uint32(v5898)+16))
	*(*int64)(unsafe.Add(mBase, uint32(v5893)+16)) = v5913 + (v5915 - v5916)
	v5920 = *(*int64)(unsafe.Add(mBase, uint32(v5893)+24))
	v5922 = *(*int64)(unsafe.Add(mBase, _c_F_heap_vacuum_rel[44]))
	v5923 = *(*int64)(unsafe.Add(mBase, uint32(v5898)+24))
	*(*int64)(unsafe.Add(mBase, uint32(v5893)+24)) = v5920 + (v5922 - v5923)
	v5927 = *(*int64)(unsafe.Add(mBase, uint32(v5893)+32))
	v5929 = *(*int64)(unsafe.Add(mBase, _c_F_heap_vacuum_rel[45]))
	v5930 = *(*int64)(unsafe.Add(mBase, uint32(v5898)+32))
	*(*int64)(unsafe.Add(mBase, uint32(v5893)+32)) = v5927 + (v5929 - v5930)
	v5934 = *(*int64)(unsafe.Add(mBase, uint32(v5893)+40))
	v5936 = *(*int64)(unsafe.Add(mBase, _c_F_heap_vacuum_rel[46]))
	v5937 = *(*int64)(unsafe.Add(mBase, uint32(v5898)+40))
	*(*int64)(unsafe.Add(mBase, uint32(v5893)+40)) = v5934 + (v5936 - v5937)
	v5941 = *(*int64)(unsafe.Add(mBase, uint32(v5893)+48))
	v5943 = *(*int64)(unsafe.Add(mBase, _c_F_heap_vacuum_rel[47]))
	v5944 = *(*int64)(unsafe.Add(mBase, uint32(v5898)+48))
	*(*int64)(unsafe.Add(mBase, uint32(v5893)+48)) = v5941 + (v5943 - v5944)
	v5948 = *(*int64)(unsafe.Add(mBase, uint32(v5893)+56))
	v5950 = *(*int64)(unsafe.Add(mBase, _c_F_heap_vacuum_rel[48]))
	v5951 = *(*int64)(unsafe.Add(mBase, uint32(v5898)+56))
	*(*int64)(unsafe.Add(mBase, uint32(v5893)+56)) = v5948 + (v5950 - v5951)
	v5955 = *(*int64)(unsafe.Add(mBase, uint32(v5893)+64))
	v5957 = *(*int64)(unsafe.Add(mBase, _c_F_heap_vacuum_rel[49]))
	v5958 = *(*int64)(unsafe.Add(mBase, uint32(v5898)+64))
	*(*int64)(unsafe.Add(mBase, uint32(v5893)+64)) = v5955 + (v5957 - v5958)
	v5962 = *(*int64)(unsafe.Add(mBase, uint32(v5893)+72))
	v5964 = *(*int64)(unsafe.Add(mBase, _c_F_heap_vacuum_rel[50]))
	v5965 = *(*int64)(unsafe.Add(mBase, uint32(v5898)+72))
	*(*int64)(unsafe.Add(mBase, uint32(v5893)+72)) = v5962 + (v5964 - v5965)
	v5969 = *(*int64)(unsafe.Add(mBase, uint32(v5893)+80))
	v5971 = *(*int64)(unsafe.Add(mBase, _c_F_heap_vacuum_rel[51]))
	v5972 = *(*int64)(unsafe.Add(mBase, uint32(v5898)+80))
	*(*int64)(unsafe.Add(mBase, uint32(v5893)+80)) = v5969 + (v5971 - v5972)
	v5976 = *(*int64)(unsafe.Add(mBase, uint32(v5893)+88))
	v5978 = *(*int64)(unsafe.Add(mBase, _c_F_heap_vacuum_rel[52]))
	v5979 = *(*int64)(unsafe.Add(mBase, uint32(v5898)+88))
	*(*int64)(unsafe.Add(mBase, uint32(v5893)+88)) = v5976 + (v5978 - v5979)
	v5983 = *(*int64)(unsafe.Add(mBase, uint32(v5893)+96))
	v5985 = *(*int64)(unsafe.Add(mBase, _c_F_heap_vacuum_rel[53]))
	v5986 = *(*int64)(unsafe.Add(mBase, uint32(v5898)+96))
	*(*int64)(unsafe.Add(mBase, uint32(v5893)+96)) = v5983 + (v5985 - v5986)
	v5990 = *(*int64)(unsafe.Add(mBase, uint32(v5893)+104))
	v5992 = *(*int64)(unsafe.Add(mBase, _c_F_heap_vacuum_rel[54]))
	v5993 = *(*int64)(unsafe.Add(mBase, uint32(v5898)+104))
	*(*int64)(unsafe.Add(mBase, uint32(v5893)+104)) = v5990 + (v5992 - v5993)
	v5997 = *(*int64)(unsafe.Add(mBase, uint32(v5893)+112))
	v5999 = *(*int64)(unsafe.Add(mBase, _c_F_heap_vacuum_rel[55]))
	v6000 = *(*int64)(unsafe.Add(mBase, uint32(v5898)+112))
	*(*int64)(unsafe.Add(mBase, uint32(v5893)+112)) = v5997 + (v5999 - v6000)
	v6004 = *(*int64)(unsafe.Add(mBase, uint32(v5893)+120))
	v6006 = *(*int64)(unsafe.Add(mBase, _c_F_heap_vacuum_rel[56]))
	v6007 = *(*int64)(unsafe.Add(mBase, uint32(v5898)+120))
	*(*int64)(unsafe.Add(mBase, uint32(v5893)+120)) = v6004 + (v6006 - v6007)
	goto L844
L844:
	;
	v6011 = *(*int64)(unsafe.Add(mBase, uint32(v53)+1056))
	v6012 = *(*int64)(unsafe.Add(mBase, uint32(v53)+1088))
	v6013 = *(*int64)(unsafe.Add(mBase, uint32(v53)+1048))
	v6014 = *(*int64)(unsafe.Add(mBase, uint32(v53)+1080))
	v6015 = *(*int64)(unsafe.Add(mBase, uint32(v53)+1040))
	v6016 = *(*int64)(unsafe.Add(mBase, uint32(v53)+1072))
	F_initStringInfo(m, v53+int32(992))
	mBase = m.M
	v6020 = m.ExcPending
	if v6020 != 0 {
		goto L24
	} else {
		goto L845
	}
L845:
	;
	if v77 != 0 {
		v6037 = int32(_a_F_heap_vacuum_rel_37)
		goto L846
	} else {
		goto L847
	}
L846:
	;
	v6038 = *(*int32)(unsafe.Add(mBase, uint32(v231)+76))
	v6039 = *(*int64)(unsafe.Add(mBase, uint32(v231)+68))
	v6040 = *(*int32)(unsafe.Add(mBase, uint32(v231)+172))
	*(*int32)(unsafe.Add(mBase, uint32(v53)+476)) = v6040
	*(*int64)(unsafe.Add(mBase, uint32(v53)+464)) = v6039
	*(*int32)(unsafe.Add(mBase, uint32(v53)+472)) = v6038
	v6045 = v53 + int32(992)
	F_appendStringInfo(m, v6045, v6037, v53+int32(464))
	mBase = m.M
	v6049 = m.ExcPending
	if v6049 != 0 {
		goto L24
	} else {
		goto L855
	}
L847:
	;
	v6024 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v231)+20)))
	if v6024&int32(1) != 0 {
		goto L848
	} else {
		goto L849
	}
L848:
	;
	v6027 = int32(_a_F_heap_vacuum_rel_38)
	goto L850
L849:
	;
	v6027 = int32(_a_F_heap_vacuum_rel_39)
	goto L850
L850:
	;
	v6028 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l1)+20)))
	if v6028 == int32(1) {
		v6037 = v6027
		goto L846
	} else {
		goto L851
	}
L851:
	;
	if v6024&int32(1) != 0 {
		goto L852
	} else {
		goto L853
	}
L852:
	;
	v6035 = int32(_a_F_heap_vacuum_rel_40)
	goto L854
L853:
	;
	v6035 = int32(_a_F_heap_vacuum_rel_41)
	goto L854
L854:
	;
	v6037 = v6035
	goto L846
L855:
	;
	v6050 = *(*int32)(unsafe.Add(mBase, uint32(v231)+112))
	v6051 = *(*int32)(unsafe.Add(mBase, uint32(v231)+120))
	v6052 = *(*int32)(unsafe.Add(mBase, uint32(v231)+116))
	*(*int32)(unsafe.Add(mBase, uint32(v53)+456)) = v6052
	v6055 = float64(100)
	if v452 != 0 {
		goto L856
	} else {
		goto L857
	}
L856:
	;
	v6060 = base.F64_div(base.F64_mul(base.F64_convert_i32_u(v6050), v6055), base.F64_convert_i32_u(v452))
	goto L858
L857:
	;
	v6060 = v6055
	goto L858
L858:
	;
	*(*float64)(unsafe.Add(mBase, uint32(v53)+448)) = v6060
	*(*int32)(unsafe.Add(mBase, uint32(v53)+440)) = v6050
	*(*int32)(unsafe.Add(mBase, uint32(v53)+436)) = v5610
	*(*int32)(unsafe.Add(mBase, uint32(v53)+432)) = v6051
	F_appendStringInfo(m, v6045, int32(_a_F_heap_vacuum_rel_42), v53+int32(432))
	mBase = m.M
	v6069 = m.ExcPending
	if v6069 != 0 {
		goto L24
	} else {
		goto L859
	}
L859:
	;
	v6070 = *(*float64)(unsafe.Add(mBase, uint32(v231)+152))
	v6071 = *(*int64)(unsafe.Add(mBase, uint32(v231)+200))
	v6072 = *(*int64)(unsafe.Add(mBase, uint32(v231)+232))
	*(*int64)(unsafe.Add(mBase, uint32(v53)+416)) = v6072
	*(*int64)(unsafe.Add(mBase, uint32(v53)+400)) = v6071
	*(*int64)(unsafe.Add(mBase, uint32(v53)+408)) = base.I64_trunc_sat_f64_s(v6070)
	F_appendStringInfo(m, v6045, int32(_a_F_heap_vacuum_rel_43), v53+int32(400))
	mBase = m.M
	v6081 = m.ExcPending
	if v6081 != 0 {
		goto L24
	} else {
		goto L860
	}
L860:
	;
	v6082 = *(*int64)(unsafe.Add(mBase, uint32(v231)+240))
	if int64(0) < v6082 {
		goto L861
	} else {
		goto L862
	}
L861:
	;
	v6085 = *(*int32)(unsafe.Add(mBase, uint32(v231)+144))
	*(*int32)(unsafe.Add(mBase, uint32(v53)+392)) = v6085
	*(*int64)(unsafe.Add(mBase, uint32(v53)+384)) = v6082
	F_appendStringInfo(m, v6045, int32(_a_F_heap_vacuum_rel_44), v53+int32(384))
	mBase = m.M
	v6092 = m.ExcPending
	if v6092 != 0 {
		goto L24
	} else {
		goto L864
	}
L862:
	;
	goto L863
L863:
	;
	v6093 = F_ReadNextFullTransactionId(m)
	mBase = m.M
	v6094 = m.ExcPending
	if v6094 != 0 {
		goto L24
	} else {
		goto L865
	}
L864:
	;
	goto L863
L865:
	;
	v6095 = *(*int32)(unsafe.Add(mBase, uint32(v231)+36))
	*(*int32)(unsafe.Add(mBase, uint32(v53)+368)) = v6095
	*(*int32)(unsafe.Add(mBase, uint32(v53)+372)) = base.I32_wrap_i64(v6093) - v6095
	v6101 = v53 + int32(992)
	F_appendStringInfo(m, v6101, int32(_a_F_heap_vacuum_rel_45), v53+int32(368))
	mBase = m.M
	v6106 = m.ExcPending
	if v6106 != 0 {
		goto L24
	} else {
		goto L866
	}
L866:
	;
	v6107 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v53)+1036)))
	if v6107 == int32(1) {
		goto L867
	} else {
		goto L868
	}
L867:
	;
	v6110 = *(*int32)(unsafe.Add(mBase, uint32(v447)))
	v6111 = *(*int32)(unsafe.Add(mBase, uint32(v1442)))
	*(*int32)(unsafe.Add(mBase, uint32(v53)+352)) = v6111
	*(*int32)(unsafe.Add(mBase, uint32(v53)+356)) = v6111 - v6110
	F_appendStringInfo(m, v6101, int32(_a_F_heap_vacuum_rel_46), v53+int32(352))
	mBase = m.M
	v6119 = m.ExcPending
	if v6119 != 0 {
		goto L24
	} else {
		goto L870
	}
L868:
	;
	goto L869
L869:
	;
	v6122 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v53)+988)))
	if v6122 == int32(1) {
		goto L871
	} else {
		goto L872
	}
L870:
	;
	goto L869
L871:
	;
	v6125 = *(*int32)(unsafe.Add(mBase, uint32(v231)+32))
	v6126 = *(*int32)(unsafe.Add(mBase, uint32(v231)+60))
	*(*int32)(unsafe.Add(mBase, uint32(v53)+336)) = v6126
	*(*int32)(unsafe.Add(mBase, uint32(v53)+340)) = v6126 - v6125
	F_appendStringInfo(m, v53+int32(992), int32(_a_F_heap_vacuum_rel_47), v53+int32(336))
	mBase = m.M
	v6136 = m.ExcPending
	if v6136 != 0 {
		goto L24
	} else {
		goto L874
	}
L872:
	;
	goto L873
L873:
	;
	v6139 = *(*int32)(unsafe.Add(mBase, uint32(v231)+124))
	v6140 = *(*int64)(unsafe.Add(mBase, uint32(v231)+208))
	*(*int64)(unsafe.Add(mBase, uint32(v53)+320)) = v6140
	v6143 = float64(100)
	if v452 != 0 {
		goto L875
	} else {
		goto L876
	}
L874:
	;
	goto L873
L875:
	;
	v6148 = base.F64_div(base.F64_mul(base.F64_convert_i32_u(v6139), v6143), base.F64_convert_i32_u(v452))
	goto L877
L876:
	;
	v6148 = v6143
	goto L877
L877:
	;
	*(*float64)(unsafe.Add(mBase, uint32(v53)+312)) = v6148
	*(*int32)(unsafe.Add(mBase, uint32(v53)+304)) = v6139
	v6152 = v53 + int32(992)
	F_appendStringInfo(m, v6152, int32(_a_F_heap_vacuum_rel_48), v53+int32(304))
	mBase = m.M
	v6157 = m.ExcPending
	if v6157 != 0 {
		goto L24
	} else {
		goto L878
	}
L878:
	;
	v6158 = *(*int32)(unsafe.Add(mBase, uint32(v231)+132))
	v6159 = *(*int32)(unsafe.Add(mBase, uint32(v231)+128))
	v6160 = *(*int32)(unsafe.Add(mBase, uint32(v231)+136))
	*(*int32)(unsafe.Add(mBase, uint32(v53)+296)) = v6160
	*(*int32)(unsafe.Add(mBase, uint32(v53)+288)) = v6159
	*(*int32)(unsafe.Add(mBase, uint32(v53)+292)) = v6158 + v6160
	F_appendStringInfo(m, v6152, int32(_a_F_heap_vacuum_rel_49), v53+int32(288))
	mBase = m.M
	v6169 = m.ExcPending
	if v6169 != 0 {
		goto L24
	} else {
		goto L879
	}
L879:
	;
	v6170 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v231)+23)))
	if v6170 == int32(1) {
		goto L881
	} else {
		goto L882
	}
L880:
	;
	F_appendStringInfoString(m, v6152, v6189)
	mBase = m.M
	v6191 = m.ExcPending
	if v6191 != 0 {
		goto L24
	} else {
		goto L891
	}
L881:
	;
	v6173 = int32(_a_F_heap_vacuum_rel_50)
	v6175 = *(*int32)(unsafe.Add(mBase, uint32(v267)))
	if v6175 == int32(0) {
		v6188 = v6173
		v6189 = int32(_a_F_heap_vacuum_rel_51)
		goto L880
	} else {
		goto L884
	}
L882:
	;
	goto L883
L883:
	;
	v6186 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_heap_vacuum_rel[14])))
	if v6186 != 0 {
		goto L888
	} else {
		goto L889
	}
L884:
	;
	v6180 = *(*int32)(unsafe.Add(mBase, uint32(v442)))
	if v6180 != 0 {
		goto L885
	} else {
		goto L886
	}
L885:
	;
	v6181 = int32(_a_F_heap_vacuum_rel_52)
	goto L887
L886:
	;
	v6181 = int32(_a_F_heap_vacuum_rel_51)
	goto L887
L887:
	;
	v6188 = v6173
	v6189 = v6181
	goto L880
L888:
	;
	v6187 = int32(_a_F_heap_vacuum_rel_53)
	goto L890
L889:
	;
	v6187 = int32(_a_F_heap_vacuum_rel_54)
	goto L890
L890:
	;
	v6188 = int32(_a_F_heap_vacuum_rel_55)
	v6189 = v6187
	goto L880
L891:
	;
	v6192 = *(*int32)(unsafe.Add(mBase, uint32(v231+int32(140))))
	v6193 = *(*int64)(unsafe.Add(mBase, uint32(v231)+216))
	*(*int64)(unsafe.Add(mBase, uint32(v53)+272)) = v6193
	v6196 = float64(100)
	if v452 != 0 {
		goto L892
	} else {
		goto L893
	}
L892:
	;
	v6201 = base.F64_div(base.F64_mul(base.F64_convert_i32_u(v6192), v6196), base.F64_convert_i32_u(v452))
	goto L894
L893:
	;
	v6201 = v6196
	goto L894
L894:
	;
	*(*float64)(unsafe.Add(mBase, uint32(v53)+264)) = v6201
	*(*int32)(unsafe.Add(mBase, uint32(v53)+256)) = v6192
	v6205 = v53 + int32(992)
	F_appendStringInfo(m, v6205, v6188, v53+int32(256))
	mBase = m.M
	v6209 = m.ExcPending
	if v6209 != 0 {
		goto L24
	} else {
		goto L895
	}
L895:
	;
	v6210 = *(*int32)(unsafe.Add(mBase, uint32(v231)+184))
	if int32(0) < v6210 {
		goto L896
	} else {
		goto L897
	}
L896:
	;
	v6213 = *(*int32)(unsafe.Add(mBase, uint32(v231)+188))
	*(*int32)(unsafe.Add(mBase, uint32(v53)+244)) = v6213
	*(*int32)(unsafe.Add(mBase, uint32(v53)+240)) = v6210
	F_appendStringInfo(m, v6205, int32(_a_F_heap_vacuum_rel_56), v53+int32(240))
	mBase = m.M
	v6220 = m.ExcPending
	if v6220 != 0 {
		goto L24
	} else {
		goto L899
	}
L897:
	;
	goto L898
L898:
	;
	v6221 = *(*int32)(unsafe.Add(mBase, uint32(v1444)))
	if int32(0) < v6221 {
		goto L900
	} else {
		goto L901
	}
L899:
	;
	goto L898
L900:
	;
	v6224 = *(*int32)(unsafe.Add(mBase, uint32(v231)+196))
	*(*int32)(unsafe.Add(mBase, uint32(v53)+228)) = v6224
	*(*int32)(unsafe.Add(mBase, uint32(v53)+224)) = v6221
	F_appendStringInfo(m, v53+int32(992), int32(_a_F_heap_vacuum_rel_57), v53+int32(224))
	mBase = m.M
	v6233 = m.ExcPending
	if v6233 != 0 {
		goto L24
	} else {
		goto L903
	}
L901:
	;
	goto L902
L902:
	;
	v6234 = int32(0)
	v6235 = *(*int32)(unsafe.Add(mBase, uint32(v267)))
	if v6234 < v6235 {
		goto L904
	} else {
		goto L905
	}
L903:
	;
	goto L902
L904:
	;
	v6245 = v6234
	v6246 = v6235
	goto L907
L905:
	;
	goto L906
L906:
	;
	v6369 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_heap_vacuum_rel[57])))
	if v6369 != 0 {
		goto L914
	} else {
		goto L915
	}
L907:
	;
	v6291 = v6245 << (uint(int32(2)) % 32)
	v6292 = *(*int32)(unsafe.Add(mBase, uint32(v231)+168))
	v6294 = *(*int32)(unsafe.Add(mBase, uint32(v6291+v6292)))
	if v6294 != 0 {
		goto L909
	} else {
		goto L910
	}
L908:
	;
	goto L906
L909:
	;
	v6296 = *(*int32)(unsafe.Add(mBase, uint32(v6291+v378)))
	v6297 = *(*int64)(unsafe.Add(mBase, uint32(v6294)+24))
	v6298 = *(*int32)(unsafe.Add(mBase, uint32(v6294)))
	v6299 = *(*int32)(unsafe.Add(mBase, uint32(v6294)+32))
	*(*int32)(unsafe.Add(mBase, uint32(v53+int32(208)))) = v6299
	*(*int32)(unsafe.Add(mBase, uint32(v53)+196)) = v6298
	*(*int64)(unsafe.Add(mBase, uint32(v53)+200)) = v6297
	*(*int32)(unsafe.Add(mBase, uint32(v53)+192)) = v6296
	F_appendStringInfo(m, v53+int32(992), int32(_a_F_heap_vacuum_rel_58), v53+int32(192))
	mBase = m.M
	v6310 = m.ExcPending
	if v6310 != 0 {
		goto L24
	} else {
		goto L912
	}
L910:
	;
	v6312 = v6246
	goto L911
L911:
	;
	v6316 = v6245 + int32(1)
	if v6316 < v6312 {
		v6245 = v6316
		v6246 = v6312
		goto L907
	} else {
		goto L913
	}
L912:
	;
	v6311 = *(*int32)(unsafe.Add(mBase, uint32(v267)))
	v6312 = v6311
	goto L911
L913:
	;
	goto L908
L914:
	;
	v6371 = *(*int32)(unsafe.Add(mBase, _c_F_heap_vacuum_rel[9]))
	v6372 = *(*int64)(unsafe.Add(mBase, uint32(v6371)+312))
	*(*float64)(unsafe.Add(mBase, uint32(v53)+176)) = base.F64_div(base.F64_convert_i64_s(v6372), float64(1e+06))
	F_appendStringInfo(m, v53+int32(992), int32(_a_F_heap_vacuum_rel_59), v53+int32(176))
	mBase = m.M
	v6383 = m.ExcPending
	if v6383 != 0 {
		goto L24
	} else {
		goto L917
	}
L915:
	;
	goto L916
L916:
	;
	v6385 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_heap_vacuum_rel[6])))
	if v6385 == int32(1) {
		goto L918
	} else {
		goto L919
	}
L917:
	;
	goto L916
L918:
	;
	v6389 = *(*int64)(unsafe.Add(mBase, _c_F_heap_vacuum_rel[8]))
	v6392 = float64(1000)
	*(*float64)(unsafe.Add(mBase, uint32(v53)+160)) = base.F64_div(base.F64_convert_i64_s(v6389-v107), v6392)
	v6396 = *(*int64)(unsafe.Add(mBase, _c_F_heap_vacuum_rel[7]))
	*(*float64)(unsafe.Add(mBase, uint32(v53)+168)) = base.F64_div(base.F64_convert_i64_s(v6396-v106), v6392)
	F_appendStringInfo(m, v53+int32(992), int32(_a_F_heap_vacuum_rel_60), v53+int32(160))
	mBase = m.M
	v6408 = m.ExcPending
	if v6408 != 0 {
		goto L24
	} else {
		goto L921
	}
L919:
	;
	goto L920
L920:
	;
	v6409 = v6011 + v6012
	v6410 = v6013 + v6014
	v6412 = *(*int32)(unsafe.Add(mBase, uint32(v53)+1680))
	v6413 = *(*int32)(unsafe.Add(mBase, uint32(v53)+1688))
	if v6413 <= int32(0) {
		goto L923
	} else {
		goto L924
	}
L921:
	;
	goto L920
L922:
	;
	*(*float64)(unsafe.Add(mBase, uint32(v53)+152)) = v6437
	*(*float64)(unsafe.Add(mBase, uint32(v53)+144)) = v6438
	v6442 = v53 + int32(992)
	F_appendStringInfo(m, v6442, int32(_a_F_heap_vacuum_rel_61), v53+int32(144))
	mBase = m.M
	v6447 = m.ExcPending
	if v6447 != 0 {
		goto L24
	} else {
		goto L927
	}
L923:
	;
	if v6412 <= int32(0) {
		v6437 = float64(0)
		v6438 = float64(0)
		goto L922
	} else {
		goto L926
	}
L924:
	;
	goto L925
L925:
	;
	v6420 = float64(8192)
	v6422 = float64(9.5367431640625e-07)
	v6428 = base.F64_add(base.F64_div(base.F64_convert_i32_s(v6412), float64(1e+06)), base.F64_convert_i32_s(v6413))
	v6437 = base.F64_div(base.F64_mul(base.F64_mul(base.F64_convert_i64_s(v6409), v6420), v6422), v6428)
	v6438 = base.F64_div(base.F64_mul(base.F64_mul(base.F64_convert_i64_s(v6410), v6420), v6422), v6428)
	goto L922
L926:
	;
	goto L925
L927:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v53)+128)) = v6409
	*(*int64)(unsafe.Add(mBase, uint32(v53)+120)) = v6410
	*(*int64)(unsafe.Add(mBase, uint32(v53)+112)) = v6015 + v6016
	F_appendStringInfo(m, v6442, int32(_a_F_heap_vacuum_rel_62), v53+int32(112))
	mBase = m.M
	v6455 = m.ExcPending
	if v6455 != 0 {
		goto L24
	} else {
		goto L928
	}
L928:
	;
	v6456 = *(*int64)(unsafe.Add(mBase, uint32(v53)+608))
	*(*int64)(unsafe.Add(mBase, uint32(v53)+80)) = v6456
	v6458 = *(*int64)(unsafe.Add(mBase, uint32(v53)+616))
	*(*int64)(unsafe.Add(mBase, uint32(v53)+88)) = v6458
	v6460 = *(*int64)(unsafe.Add(mBase, uint32(v53)+624))
	*(*int64)(unsafe.Add(mBase, uint32(v53)+96)) = v6460
	v6462 = *(*int64)(unsafe.Add(mBase, uint32(v53)+592))
	*(*int64)(unsafe.Add(mBase, uint32(v53)+64)) = v6462
	v6464 = *(*int64)(unsafe.Add(mBase, uint32(v53)+600))
	*(*int64)(unsafe.Add(mBase, uint32(v53)+72)) = v6464
	F_appendStringInfo(m, v6442, int32(_a_F_heap_vacuum_rel_63), v53-int32(-64))
	mBase = m.M
	v6470 = m.ExcPending
	if v6470 != 0 {
		goto L24
	} else {
		goto L929
	}
L929:
	;
	v6471 = *(*int32)(unsafe.Add(mBase, uint32(v231)+180))
	v6472 = *(*int32)(unsafe.Add(mBase, uint32(v231)+176))
	v6474 = float64(9.5367431640625e-07)
	*(*float64)(unsafe.Add(mBase, uint32(v53)+48)) = base.F64_mul(base.F64_convert_i32_u(v4220), v6474)
	*(*int32)(unsafe.Add(mBase, uint32(v53)+40)) = v6472
	*(*float64)(unsafe.Add(mBase, uint32(v53)+32)) = base.F64_mul(base.F64_convert_i32_u(v6471), v6474)
	if v6472 == int32(1) {
		goto L930
	} else {
		goto L931
	}
L930:
	;
	v6486 = int32(_a_F_heap_vacuum_rel_64)
	goto L932
L931:
	;
	v6486 = int32(_a_F_heap_vacuum_rel_65)
	goto L932
L932:
	;
	F_appendStringInfo(m, v6442, v6486, v53+int32(32))
	mBase = m.M
	v6490 = m.ExcPending
	if v6490 != 0 {
		goto L24
	} else {
		goto L933
	}
L933:
	;
	v6493 = F_pg_rusage_show(m, v53+int32(816))
	mBase = m.M
	v6494 = m.ExcPending
	if v6494 != 0 {
		goto L24
	} else {
		goto L934
	}
L934:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v53)+16)) = v6493
	F_appendStringInfo(m, v6442, int32(_a_F_heap_vacuum_rel_66), v53+int32(16))
	mBase = m.M
	v6500 = m.ExcPending
	if v6500 != 0 {
		goto L24
	} else {
		goto L935
	}
L935:
	;
	if v77 != 0 {
		goto L936
	} else {
		goto L937
	}
L936:
	;
	v6503 = int32(17)
	goto L938
L937:
	;
	v6503 = int32(15)
	goto L938
L938:
	;
	v6505 = F_errstart(m, v6503, int32(0))
	mBase = m.M
	v6506 = m.ExcPending
	if v6506 != 0 {
		goto L24
	} else {
		goto L939
	}
L939:
	;
	if v6505 != 0 {
		goto L940
	} else {
		goto L941
	}
L940:
	;
	v6507 = *(*int32)(unsafe.Add(mBase, uint32(v53)+992))
	*(*int32)(unsafe.Add(mBase, uint32(v53))) = v6507
	F_errmsg_internal(m, int32(_a_F_heap_vacuum_rel_67), v53)
	mBase = m.M
	v6511 = m.ExcPending
	if v6511 != 0 {
		goto L24
	} else {
		goto L943
	}
L941:
	;
	goto L942
L942:
	;
	v6517 = *(*int32)(unsafe.Add(mBase, uint32(v53)+992))
	F_pfree(m, v6517)
	mBase = m.M
	v6519 = m.ExcPending
	if v6519 != 0 {
		goto L24
	} else {
		goto L945
	}
L943:
	;
	F_errfinish(m, int32(_a_F_heap_vacuum_rel_7), int32(1226), int32(_a_F_heap_vacuum_rel_8))
	mBase = m.M
	v6516 = m.ExcPending
	if v6516 != 0 {
		goto L24
	} else {
		goto L944
	}
L944:
	;
	goto L942
L945:
	;
	goto L831
L946:
	;
	v6575 = v5625
	goto L949
L947:
	;
	goto L948
L948:
	;
	m.G0 = v53 + int32(1696)
	return
L949:
	;
	v6624 = v6575 << (uint(int32(2)) % 32)
	v6625 = *(*int32)(unsafe.Add(mBase, uint32(v231)+168))
	v6627 = *(*int32)(unsafe.Add(mBase, uint32(v6624+v6625)))
	if v6627 != 0 {
		goto L951
	} else {
		goto L952
	}
L950:
	;
	goto L948
L951:
	;
	F_pfree(m, v6627)
	mBase = m.M
	v6629 = m.ExcPending
	if v6629 != 0 {
		goto L24
	} else {
		goto L954
	}
L952:
	;
	goto L953
L953:
	;
	if v105 != 0 {
		goto L955
	} else {
		goto L956
	}
L954:
	;
	goto L953
L955:
	;
	v6631 = *(*int32)(unsafe.Add(mBase, uint32(v6624+v378)))
	F_pfree(m, v6631)
	mBase = m.M
	v6633 = m.ExcPending
	if v6633 != 0 {
		goto L24
	} else {
		goto L958
	}
L956:
	;
	goto L957
L957:
	;
	v6635 = v6575 + int32(1)
	v6636 = *(*int32)(unsafe.Add(mBase, uint32(v267)))
	if v6635 < v6636 {
		v6575 = v6635
		goto L949
	} else {
		goto L959
	}
L958:
	;
	goto L957
L959:
	;
	goto L950
}
