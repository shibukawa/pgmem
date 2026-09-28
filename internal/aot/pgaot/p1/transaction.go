package p1

import (
	base "github.com/shibukawa/pgmem/internal/aot/pgaot/base"
	"sync/atomic"
	"unsafe"
)

func F_CommitTransaction(m *base.Module) {
	mBase := m.M
	_ = mBase
	var v26 int32
	_ = v26
	var v28 int32
	_ = v28
	var v31 int32
	_ = v31
	var v32 int32
	_ = v32
	var v35 int32
	_ = v35
	var v39 int32
	_ = v39
	var v46 int32
	_ = v46
	var v49 int32
	_ = v49
	var v62 int32
	_ = v62
	var v66 int32
	_ = v66
	var v70 int32
	_ = v70
	var v76 int32
	_ = v76
	var v79 int32
	_ = v79
	var v83 int32
	_ = v83
	var v85 int32
	_ = v85
	var v86 int32
	_ = v86
	var v91 int32
	_ = v91
	var v92 int32
	_ = v92
	var v95 int32
	_ = v95
	var v100 int32
	_ = v100
	var v102 int32
	_ = v102
	var v108 int32
	_ = v108
	var v113 int32
	_ = v113
	var v141 int32
	_ = v141
	var v143 int32
	_ = v143
	var v144 int32
	_ = v144
	var v146 int32
	_ = v146
	var v148 int32
	_ = v148
	var v151 int32
	_ = v151
	var v153 int32
	_ = v153
	var v177 int32
	_ = v177
	var v178 int32
	_ = v178
	var v179 int32
	_ = v179
	var v181 int32
	_ = v181
	var v209 int32
	_ = v209
	var v210 int32
	_ = v210
	var v217 int32
	_ = v217
	var v218 int32
	_ = v218
	var v227 int32
	_ = v227
	var v228 int32
	_ = v228
	var v233 int32
	_ = v233
	var v234 int32
	_ = v234
	var v235 int32
	_ = v235
	var v238 int32
	_ = v238
	var v242 int32
	_ = v242
	var v245 int32
	_ = v245
	var v247 int32
	_ = v247
	var v250 int32
	_ = v250
	var v252 int32
	_ = v252
	var v255 int32
	_ = v255
	var v256 int32
	_ = v256
	var v258 int32
	_ = v258
	var v261 int32
	_ = v261
	var v263 int32
	_ = v263
	var v268 int32
	_ = v268
	var v273 int32
	_ = v273
	var v274 int32
	_ = v274
	var v280 int32
	_ = v280
	var v285 int32
	_ = v285
	var v287 int32
	_ = v287
	var v288 int32
	_ = v288
	var v290 int32
	_ = v290
	var v292 int32
	_ = v292
	var v296 int32
	_ = v296
	var v297 int32
	_ = v297
	var v298 int32
	_ = v298
	var v299 int32
	_ = v299
	var v302 int32
	_ = v302
	var v305 int32
	_ = v305
	var v306 int32
	_ = v306
	var v313 int32
	_ = v313
	var v314 int32
	_ = v314
	var v317 int32
	_ = v317
	var v319 int32
	_ = v319
	var v321 int32
	_ = v321
	var v324 int32
	_ = v324
	var v327 int32
	_ = v327
	var v328 int32
	_ = v328
	var v331 int32
	_ = v331
	var v332 int32
	_ = v332
	var v333 int32
	_ = v333
	var v335 int32
	_ = v335
	var v338 int32
	_ = v338
	var v339 int32
	_ = v339
	var v340 int32
	_ = v340
	var v342 int32
	_ = v342
	var v345 int32
	_ = v345
	var v347 int32
	_ = v347
	var v348 int32
	_ = v348
	var v349 int32
	_ = v349
	var v352 int32
	_ = v352
	var v355 int32
	_ = v355
	var v358 int32
	_ = v358
	var v359 int32
	_ = v359
	var v361 int32
	_ = v361
	var v362 int32
	_ = v362
	var v368 int32
	_ = v368
	var v372 int32
	_ = v372
	var v376 int32
	_ = v376
	var v380 int32
	_ = v380
	var v391 int32
	_ = v391
	var v392 int32
	_ = v392
	var v395 int32
	_ = v395
	var v396 int32
	_ = v396
	var v398 int32
	_ = v398
	var v402 int32
	_ = v402
	var v406 int32
	_ = v406
	var v407 int64
	_ = v407
	var v409 int64
	_ = v409
	var v413 int32
	_ = v413
	var v414 int32
	_ = v414
	var v417 int32
	_ = v417
	var v419 int32
	_ = v419
	var v420 int32
	_ = v420
	var v423 int32
	_ = v423
	var v427 int32
	_ = v427
	var v433 int32
	_ = v433
	var v453 int32
	_ = v453
	var v457 int32
	_ = v457
	var v458 int32
	_ = v458
	var v460 int32
	_ = v460
	var v462 int32
	_ = v462
	var v467 int32
	_ = v467
	var v468 int32
	_ = v468
	var v472 int32
	_ = v472
	var v478 int32
	_ = v478
	var v483 int32
	_ = v483
	var v485 int32
	_ = v485
	var v491 int32
	_ = v491
	var v493 int32
	_ = v493
	var v496 int32
	_ = v496
	var v500 int32
	_ = v500
	var v501 int32
	_ = v501
	var v502 int32
	_ = v502
	var v504 int32
	_ = v504
	var v506 int32
	_ = v506
	var v507 int32
	_ = v507
	var v508 int32
	_ = v508
	var v509 int64
	_ = v509
	var v510 int32
	_ = v510
	var v511 int64
	_ = v511
	var v512 int32
	_ = v512
	var v518 int32
	_ = v518
	var v519 int32
	_ = v519
	var v523 int32
	_ = v523
	var v525 int32
	_ = v525
	var v528 int32
	_ = v528
	var v540 int64
	_ = v540
	var v545 int32
	_ = v545
	var v546 int32
	_ = v546
	var v547 int32
	_ = v547
	var v549 int32
	_ = v549
	var v550 int64
	_ = v550
	var v552 int32
	_ = v552
	var v554 int32
	_ = v554
	var v556 int32
	_ = v556
	var v557 int64
	_ = v557
	var v558 int32
	_ = v558
	var v560 int32
	_ = v560
	var v561 int32
	_ = v561
	var v563 int64
	_ = v563
	var v566 int32
	_ = v566
	var v568 int32
	_ = v568
	var v575 int32
	_ = v575
	var v577 int32
	_ = v577
	var v580 int32
	_ = v580
	var v592 int64
	_ = v592
	var v596 int32
	_ = v596
	var v598 int32
	_ = v598
	var v602 int32
	_ = v602
	var v603 int32
	_ = v603
	var v604 int32
	_ = v604
	var v605 int32
	_ = v605
	var v612 int32
	_ = v612
	var v615 int32
	_ = v615
	var v617 int32
	_ = v617
	var v622 int32
	_ = v622
	var v625 int32
	_ = v625
	var v627 int32
	_ = v627
	var v631 int32
	_ = v631
	var v634 int32
	_ = v634
	var v636 int32
	_ = v636
	var v643 int32
	_ = v643
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
	var v659 int32
	_ = v659
	var v662 int32
	_ = v662
	var v664 int32
	_ = v664
	var v665 int32
	_ = v665
	var v669 int32
	_ = v669
	var v672 int32
	_ = v672
	var v674 int32
	_ = v674
	var v679 int32
	_ = v679
	var v683 int32
	_ = v683
	var v685 int32
	_ = v685
	var v691 int32
	_ = v691
	var v718 int32
	_ = v718
	var v720 int32
	_ = v720
	var v723 int32
	_ = v723
	var v724 int32
	_ = v724
	var v725 int32
	_ = v725
	var v728 int32
	_ = v728
	var v731 int32
	_ = v731
	var v732 int32
	_ = v732
	var v734 int32
	_ = v734
	var v735 int64
	_ = v735
	var v758 int32
	_ = v758
	var v762 int32
	_ = v762
	var v774 int32
	_ = v774
	var v775 int32
	_ = v775
	var v776 int32
	_ = v776
	var v778 int32
	_ = v778
	var v782 int32
	_ = v782
	var v783 int32
	_ = v783
	var v785 int32
	_ = v785
	var v786 int32
	_ = v786
	var v787 int32
	_ = v787
	var v789 int32
	_ = v789
	var v795 int32
	_ = v795
	var v796 int32
	_ = v796
	var v797 int32
	_ = v797
	var v798 int32
	_ = v798
	var v801 int32
	_ = v801
	var v808 int32
	_ = v808
	var v809 int32
	_ = v809
	var v810 int32
	_ = v810
	var v813 int32
	_ = v813
	var v816 int32
	_ = v816
	var v821 int32
	_ = v821
	var v822 int32
	_ = v822
	var v824 int32
	_ = v824
	var v826 int32
	_ = v826
	var v830 int32
	_ = v830
	var v831 int32
	_ = v831
	var v832 int32
	_ = v832
	var v837 int32
	_ = v837
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
	var v850 int32
	_ = v850
	var v851 int32
	_ = v851
	var v853 int32
	_ = v853
	var v855 int32
	_ = v855
	var v857 int32
	_ = v857
	var v858 int32
	_ = v858
	var v861 int32
	_ = v861
	var v868 int32
	_ = v868
	var v872 int32
	_ = v872
	var v877 int32
	_ = v877
	var v878 int32
	_ = v878
	var v879 int32
	_ = v879
	var v886 int32
	_ = v886
	var v889 int32
	_ = v889
	var v892 int32
	_ = v892
	var v893 int32
	_ = v893
	var v897 int32
	_ = v897
	var v899 int32
	_ = v899
	var v900 int32
	_ = v900
	var v901 int32
	_ = v901
	var v902 int32
	_ = v902
	var v907 int32
	_ = v907
	var v908 int32
	_ = v908
	var v936 int32
	_ = v936
	var v939 int32
	_ = v939
	var v966 int32
	_ = v966
	var v968 int32
	_ = v968
	var v970 int32
	_ = v970
	var v974 int32
	_ = v974
	var v975 int32
	_ = v975
	var v977 int32
	_ = v977
	var v978 int32
	_ = v978
	var v979 int32
	_ = v979
	var v980 int32
	_ = v980
	var v982 int32
	_ = v982
	var v989 int32
	_ = v989
	var v991 int32
	_ = v991
	var v992 int32
	_ = v992
	var v993 int32
	_ = v993
	var v994 int32
	_ = v994
	var v997 int32
	_ = v997
	var v1001 int32
	_ = v1001
	var v1003 int32
	_ = v1003
	var v1007 int32
	_ = v1007
	var v1009 int32
	_ = v1009
	var v1039 int32
	_ = v1039
	var v1041 int32
	_ = v1041
	var v1043 int32
	_ = v1043
	var v1045 int32
	_ = v1045
	var v1046 int32
	_ = v1046
	var v1048 int32
	_ = v1048
	var v1049 int32
	_ = v1049
	var v1053 int32
	_ = v1053
	var v1056 int32
	_ = v1056
	var v1057 int32
	_ = v1057
	var v1061 int32
	_ = v1061
	var v1063 int32
	_ = v1063
	var v1065 int32
	_ = v1065
	var v1066 int32
	_ = v1066
	var v1067 int32
	_ = v1067
	var v1070 int32
	_ = v1070
	var v1096 int32
	_ = v1096
	var v1099 int32
	_ = v1099
	var v1100 int32
	_ = v1100
	var v1105 int32
	_ = v1105
	var v1106 int32
	_ = v1106
	var v1133 int32
	_ = v1133
	var v1134 int32
	_ = v1134
	var v1162 int32
	_ = v1162
	var v1167 int32
	_ = v1167
	var v1169 int32
	_ = v1169
	var v1171 int32
	_ = v1171
	var v1172 int32
	_ = v1172
	var v1173 int32
	_ = v1173
	var v1176 int32
	_ = v1176
	var v1202 int32
	_ = v1202
	var v1203 int32
	_ = v1203
	var v1204 int32
	_ = v1204
	var v1205 int32
	_ = v1205
	var v1207 int32
	_ = v1207
	var v1211 int32
	_ = v1211
	var v1212 int32
	_ = v1212
	var v1213 int32
	_ = v1213
	var v1216 int32
	_ = v1216
	var v1223 int32
	_ = v1223
	var v1225 int32
	_ = v1225
	var v1245 int32
	_ = v1245
	var v1249 int32
	_ = v1249
	var v1251 int32
	_ = v1251
	var v1252 int32
	_ = v1252
	var v1255 int32
	_ = v1255
	var v1258 int32
	_ = v1258
	var v1261 int32
	_ = v1261
	var v1262 int32
	_ = v1262
	var v1264 int32
	_ = v1264
	var v1292 int32
	_ = v1292
	var v1295 int32
	_ = v1295
	var v1298 int32
	_ = v1298
	var v1301 int32
	_ = v1301
	var v1302 int32
	_ = v1302
	var v1305 int32
	_ = v1305
	var v1306 int32
	_ = v1306
	var v1309 int32
	_ = v1309
	var v1316 int32
	_ = v1316
	var v1317 int32
	_ = v1317
	var v1322 int32
	_ = v1322
	var v1349 int32
	_ = v1349
	var v1350 int32
	_ = v1350
	var v1352 int32
	_ = v1352
	var v1359 int32
	_ = v1359
	var v1380 int32
	_ = v1380
	var v1381 int32
	_ = v1381
	var v1409 int32
	_ = v1409
	var v1414 int32
	_ = v1414
	var v1416 int32
	_ = v1416
	var v1419 int32
	_ = v1419
	var v1420 int32
	_ = v1420
	var v1423 int32
	_ = v1423
	var v1428 int32
	_ = v1428
	var v1430 int32
	_ = v1430
	var v1433 int32
	_ = v1433
	var v1434 int32
	_ = v1434
	var v1436 int32
	_ = v1436
	var v1437 int32
	_ = v1437
	var v1442 int32
	_ = v1442
	var v1444 int64
	_ = v1444
	var v1450 int32
	_ = v1450
	var v1456 int32
	_ = v1456
	var v1457 int32
	_ = v1457
	var v1460 int32
	_ = v1460
	var v1469 int32
	_ = v1469
	var v1470 int32
	_ = v1470
	var v1492 int32
	_ = v1492
	var v1496 int32
	_ = v1496
	var v1497 int32
	_ = v1497
	var v1499 int32
	_ = v1499
	var v1503 int64
	_ = v1503
	var v1506 int64
	_ = v1506
	var v1508 int64
	_ = v1508
	var v1509 int64
	_ = v1509
	var v1514 int32
	_ = v1514
	var v1516 float64
	_ = v1516
	var v1522 int32
	_ = v1522
	var v1523 int32
	_ = v1523
	var v1524 int32
	_ = v1524
	var v1527 int64
	_ = v1527
	var v1528 int64
	_ = v1528
	var v1536 int64
	_ = v1536
	var v1538 int32
	_ = v1538
	var v1539 int64
	_ = v1539
	var v1547 int32
	_ = v1547
	var v1550 int32
	_ = v1550
	var v1551 int32
	_ = v1551
	var v1555 int32
	_ = v1555
	var v1556 int64
	_ = v1556
	var v1557 int32
	_ = v1557
	var v1559 int32
	_ = v1559
	var v1563 int32
	_ = v1563
	var v1564 int32
	_ = v1564
	var v1579 int64
	_ = v1579
	var v1584 int32
	_ = v1584
	var v1585 int32
	_ = v1585
	var v1586 int64
	_ = v1586
	var v1588 int32
	_ = v1588
	var v1591 int32
	_ = v1591
	var v1592 int32
	_ = v1592
	var v1596 int32
	_ = v1596
	var v1599 int32
	_ = v1599
	var v1600 int64
	_ = v1600
	var v1602 int32
	_ = v1602
	var v1603 int32
	_ = v1603
	var v1605 int32
	_ = v1605
	var v1606 int64
	_ = v1606
	var v1607 int32
	_ = v1607
	var v1615 int32
	_ = v1615
	var v1637 int32
	_ = v1637
	var v1638 int32
	_ = v1638
	var v1646 int32
	_ = v1646
	var v1651 int32
	_ = v1651
	var v1652 int32
	_ = v1652
	var v1656 int32
	_ = v1656
	var v1661 int32
	_ = v1661
	var v1663 int32
	_ = v1663
	var v1665 int32
	_ = v1665
	var v1691 int64
	_ = v1691
	var v1692 int64
	_ = v1692
	var v1693 int64
	_ = v1693
	var v1696 int64
	_ = v1696
	var v1698 int64
	_ = v1698
	var v1701 int32
	_ = v1701
	var v1702 int32
	_ = v1702
	var v1704 int64
	_ = v1704
	var v1705 int64
	_ = v1705
	var v1709 int32
	_ = v1709
	var v1711 int32
	_ = v1711
	var v1712 int32
	_ = v1712
	var v1715 int32
	_ = v1715
	var v1718 int32
	_ = v1718
	var v1719 int32
	_ = v1719
	var v1724 int32
	_ = v1724
	var v1725 int32
	_ = v1725
	var v1726 int32
	_ = v1726
	var v1728 int32
	_ = v1728
	var v1729 int32
	_ = v1729
	var v1731 int32
	_ = v1731
	var v1737 int32
	_ = v1737
	var v1758 int32
	_ = v1758
	var v1762 int32
	_ = v1762
	var v1763 int32
	_ = v1763
	var v1764 int32
	_ = v1764
	var v1766 int32
	_ = v1766
	var v1768 int32
	_ = v1768
	var v1772 int32
	_ = v1772
	var v1774 int32
	_ = v1774
	var v1775 int32
	_ = v1775
	var v1778 int32
	_ = v1778
	var v1781 int32
	_ = v1781
	var v1785 int32
	_ = v1785
	var v1790 int32
	_ = v1790
	var v1793 int32
	_ = v1793
	var v1794 int32
	_ = v1794
	var v1795 int32
	_ = v1795
	var v1796 int32
	_ = v1796
	var v1801 int32
	_ = v1801
	var v1802 int32
	_ = v1802
	var v1807 int32
	_ = v1807
	var v1809 int32
	_ = v1809
	var v1810 int32
	_ = v1810
	var v1813 int32
	_ = v1813
	var v1814 int32
	_ = v1814
	var v1818 int32
	_ = v1818
	var v1823 int64
	_ = v1823
	var v1824 int32
	_ = v1824
	var v1825 int32
	_ = v1825
	var v1827 int32
	_ = v1827
	var v1829 int32
	_ = v1829
	var v1831 int64
	_ = v1831
	var v1834 int32
	_ = v1834
	var v1839 int32
	_ = v1839
	var v1840 int32
	_ = v1840
	var v1842 int64
	_ = v1842
	var v1843 int64
	_ = v1843
	var v1847 int32
	_ = v1847
	var v1850 int32
	_ = v1850
	var v1852 int32
	_ = v1852
	var v1853 int32
	_ = v1853
	var v1854 int64
	_ = v1854
	var v1855 int32
	_ = v1855
	var v1856 int64
	_ = v1856
	var v1858 int32
	_ = v1858
	var v1859 int32
	_ = v1859
	var v1860 int32
	_ = v1860
	var v1864 int32
	_ = v1864
	var v1866 int32
	_ = v1866
	var v1867 int32
	_ = v1867
	var v1872 int32
	_ = v1872
	var v1873 int64
	_ = v1873
	var v1875 int64
	_ = v1875
	var v1878 int32
	_ = v1878
	var v1881 int32
	_ = v1881
	var v1882 int64
	_ = v1882
	var v1885 int64
	_ = v1885
	var v1888 int32
	_ = v1888
	var v1892 int32
	_ = v1892
	var v1924 int32
	_ = v1924
	var v1927 int32
	_ = v1927
	var v1931 int32
	_ = v1931
	var v1936 int32
	_ = v1936
	var v1940 int32
	_ = v1940
	var v1941 int32
	_ = v1941
	var v1943 int32
	_ = v1943
	var v1944 int32
	_ = v1944
	var v1949 int32
	_ = v1949
	var v1950 int32
	_ = v1950
	var v1957 int32
	_ = v1957
	var v1962 int32
	_ = v1962
	var v1966 int32
	_ = v1966
	var v1967 int32
	_ = v1967
	var v1974 int32
	_ = v1974
	var v1978 int32
	_ = v1978
	var v1984 int32
	_ = v1984
	var v1989 int32
	_ = v1989
	var v1990 int32
	_ = v1990
	var v1993 int32
	_ = v1993
	var v1994 int32
	_ = v1994
	var v1995 int32
	_ = v1995
	var v1996 int32
	_ = v1996
	var v2000 int32
	_ = v2000
	var v2001 int32
	_ = v2001
	var v2003 int32
	_ = v2003
	var v2010 int32
	_ = v2010
	var v2011 int32
	_ = v2011
	var v2012 int32
	_ = v2012
	var v2014 int64
	_ = v2014
	var v2018 int32
	_ = v2018
	var v2019 int32
	_ = v2019
	var v2020 int32
	_ = v2020
	var v2022 int32
	_ = v2022
	var v2028 int32
	_ = v2028
	var v2031 int32
	_ = v2031
	var v2035 int32
	_ = v2035
	var v2038 int32
	_ = v2038
	var v2042 int32
	_ = v2042
	var v2045 int64
	_ = v2045
	var v2046 int32
	_ = v2046
	var v2051 int32
	_ = v2051
	var v2054 int32
	_ = v2054
	var v2056 int32
	_ = v2056
	var v2061 int32
	_ = v2061
	var v2063 int32
	_ = v2063
	var v2064 int32
	_ = v2064
	var v2068 int32
	_ = v2068
	var v2071 int32
	_ = v2071
	var v2073 int64
	_ = v2073
	var v2080 int32
	_ = v2080
	var v2081 int32
	_ = v2081
	var v2082 int32
	_ = v2082
	var v2085 int64
	_ = v2085
	var v2086 int64
	_ = v2086
	var v2094 int64
	_ = v2094
	var v2096 int64
	_ = v2096
	var v2099 int32
	_ = v2099
	var v2100 int32
	_ = v2100
	var v2101 int32
	_ = v2101
	var v2102 int32
	_ = v2102
	var v2104 int32
	_ = v2104
	var v2105 int32
	_ = v2105
	var v2107 int64
	_ = v2107
	var v2108 int32
	_ = v2108
	var v2114 int64
	_ = v2114
	var v2116 int64
	_ = v2116
	var v2118 int32
	_ = v2118
	var v2120 int64
	_ = v2120
	var v2125 int64
	_ = v2125
	var v2132 int32
	_ = v2132
	var v2133 int32
	_ = v2133
	var v2134 int32
	_ = v2134
	var v2137 int64
	_ = v2137
	var v2138 int64
	_ = v2138
	var v2146 int64
	_ = v2146
	var v2148 int64
	_ = v2148
	var v2151 int64
	_ = v2151
	var v2155 int32
	_ = v2155
	var v2157 int32
	_ = v2157
	var v2158 int32
	_ = v2158
	var v2164 int64
	_ = v2164
	var v2166 int32
	_ = v2166
	var v2171 int32
	_ = v2171
	var v2177 int32
	_ = v2177
	var v2181 int32
	_ = v2181
	var v2183 int32
	_ = v2183
	var v2187 int64
	_ = v2187
	var v2189 int32
	_ = v2189
	var v2191 int32
	_ = v2191
	var v2192 int32
	_ = v2192
	var v2196 int32
	_ = v2196
	var v2198 int32
	_ = v2198
	var v2206 int32
	_ = v2206
	var v2211 int32
	_ = v2211
	var v2214 int32
	_ = v2214
	var v2226 int32
	_ = v2226
	var v2229 int32
	_ = v2229
	var v2231 int32
	_ = v2231
	var v2234 int32
	_ = v2234
	var v2235 int32
	_ = v2235
	var v2239 int32
	_ = v2239
	var v2243 int32
	_ = v2243
	var v2244 int32
	_ = v2244
	var v2254 int32
	_ = v2254
	var v2257 int32
	_ = v2257
	var v2269 int32
	_ = v2269
	var v2274 int32
	_ = v2274
	var v2280 int64
	_ = v2280
	var v2283 int32
	_ = v2283
	var v2285 int32
	_ = v2285
	var v2286 int64
	_ = v2286
	var v2291 int32
	_ = v2291
	var v2296 int32
	_ = v2296
	var v2298 int32
	_ = v2298
	var v2301 int32
	_ = v2301
	var v2303 int32
	_ = v2303
	var v2306 int64
	_ = v2306
	var v2308 int32
	_ = v2308
	var v2309 int32
	_ = v2309
	var v2310 int32
	_ = v2310
	var v2313 int32
	_ = v2313
	var v2316 int32
	_ = v2316
	var v2317 int64
	_ = v2317
	var v2320 int32
	_ = v2320
	var v2324 int32
	_ = v2324
	var v2335 int32
	_ = v2335
	var v2337 int32
	_ = v2337
	var v2339 int32
	_ = v2339
	var v2341 int32
	_ = v2341
	var v2343 int32
	_ = v2343
	var v2367 int32
	_ = v2367
	var v2368 int32
	_ = v2368
	var v2369 int32
	_ = v2369
	var v2371 int32
	_ = v2371
	var v2401 int32
	_ = v2401
	var v2402 int32
	_ = v2402
	var v2406 int32
	_ = v2406
	var v2408 int32
	_ = v2408
	var v2411 int32
	_ = v2411
	var v2413 int32
	_ = v2413
	var v2416 int32
	_ = v2416
	var v2418 int32
	_ = v2418
	var v2419 int32
	_ = v2419
	var v2420 int32
	_ = v2420
	var v2421 int32
	_ = v2421
	var v2424 int32
	_ = v2424
	var v2427 int32
	_ = v2427
	var v2429 int32
	_ = v2429
	var v2436 int32
	_ = v2436
	var v2448 int32
	_ = v2448
	var v2450 int32
	_ = v2450
	var v2453 int32
	_ = v2453
	var v2455 int32
	_ = v2455
	var v2457 int32
	_ = v2457
	var v2460 int32
	_ = v2460
	var v2463 int32
	_ = v2463
	var v2464 int32
	_ = v2464
	var v2466 int32
	_ = v2466
	var v2468 int32
	_ = v2468
	var v2471 int32
	_ = v2471
	var v2473 int32
	_ = v2473
	var v2476 int32
	_ = v2476
	var v2481 int32
	_ = v2481
	var v2482 int32
	_ = v2482
	var v2488 int32
	_ = v2488
	var v2493 int32
	_ = v2493
	var v2496 int32
	_ = v2496
	var v2498 int32
	_ = v2498
	var v2505 int32
	_ = v2505
	var v2509 int32
	_ = v2509
	var v2511 int32
	_ = v2511
	var v2512 int32
	_ = v2512
	var v2513 int32
	_ = v2513
	var v2517 int32
	_ = v2517
	var v2545 int32
	_ = v2545
	var v2546 int32
	_ = v2546
	var v2573 int32
	_ = v2573
	var v2577 int32
	_ = v2577
	var v2579 int32
	_ = v2579
	var v2580 int64
	_ = v2580
	var v2581 int64
	_ = v2581
	var v2584 int64
	_ = v2584
	var v2585 int64
	_ = v2585
	var v2586 int64
	_ = v2586
	var v2587 int64
	_ = v2587
	var v2588 int64
	_ = v2588
	var v2589 int64
	_ = v2589
	var v2590 int64
	_ = v2590
	var v2591 int64
	_ = v2591
	var v2592 int64
	_ = v2592
	var v2593 int64
	_ = v2593
	var v2594 int64
	_ = v2594
	var v2595 int64
	_ = v2595
	var v2596 int64
	_ = v2596
	var v2597 int64
	_ = v2597
	var v2598 int64
	_ = v2598
	var v2599 int64
	_ = v2599
	var v2600 int64
	_ = v2600
	var v2601 int64
	_ = v2601
	var v2602 int64
	_ = v2602
	var v2603 int64
	_ = v2603
	var v2604 int64
	_ = v2604
	var v2605 int64
	_ = v2605
	var v2606 int64
	_ = v2606
	var v2607 int64
	_ = v2607
	var v2608 int64
	_ = v2608
	var v2609 int64
	_ = v2609
	var v2610 int64
	_ = v2610
	var v2611 int64
	_ = v2611
	var v2612 int64
	_ = v2612
	var v2613 int64
	_ = v2613
	var v2614 int64
	_ = v2614
	var v2646 int64
	_ = v2646
	var v2650 int32
	_ = v2650
	var v2653 int32
	_ = v2653
	var v2657 int32
	_ = v2657
	var v2661 int32
	_ = v2661
	var v2662 int32
	_ = v2662
	var v2664 int32
	_ = v2664
	var v2665 int32
	_ = v2665
	var v2668 int32
	_ = v2668
	var v2672 int32
	_ = v2672
	var v2677 int32
	_ = v2677
	var v2682 int32
	_ = v2682
	var v2698 int32
	_ = v2698
	var v2702 int32
	_ = v2702
	var v2704 int32
	_ = v2704
	var v2705 int64
	_ = v2705
	var v2728 int32
	_ = v2728
	var v2732 int32
	_ = v2732
	var v2744 int32
	_ = v2744
	var v2745 int32
	_ = v2745
	var v2746 int32
	_ = v2746
	var v2748 int32
	_ = v2748
	var v2752 int32
	_ = v2752
	var v2753 int32
	_ = v2753
	var v2755 int32
	_ = v2755
	var v2756 int32
	_ = v2756
	var v2757 int32
	_ = v2757
	var v2759 int32
	_ = v2759
	var v2765 int32
	_ = v2765
	var v2766 int32
	_ = v2766
	var v2767 int32
	_ = v2767
	var v2768 int32
	_ = v2768
	var v2771 int32
	_ = v2771
	var v2778 int32
	_ = v2778
	var v2779 int32
	_ = v2779
	var v2780 int32
	_ = v2780
	var v2783 int32
	_ = v2783
	var v2786 int32
	_ = v2786
	var v2791 int32
	_ = v2791
	var v2792 int32
	_ = v2792
	var v2794 int32
	_ = v2794
	var v2796 int32
	_ = v2796
	var v2800 int32
	_ = v2800
	var v2801 int32
	_ = v2801
	var v2802 int32
	_ = v2802
	var v2807 int32
	_ = v2807
	var v2808 int32
	_ = v2808
	var v2809 int32
	_ = v2809
	var v2812 int32
	_ = v2812
	var v2813 int32
	_ = v2813
	var v2814 int32
	_ = v2814
	var v2816 int32
	_ = v2816
	var v2820 int32
	_ = v2820
	var v2821 int32
	_ = v2821
	var v2823 int32
	_ = v2823
	var v2825 int32
	_ = v2825
	var v2827 int32
	_ = v2827
	var v2828 int32
	_ = v2828
	var v2831 int32
	_ = v2831
	var v2838 int32
	_ = v2838
	var v2842 int32
	_ = v2842
	var v2846 int32
	_ = v2846
	var v2847 int32
	_ = v2847
	var v2849 int32
	_ = v2849
	var v2850 int32
	_ = v2850
	var v2851 int32
	_ = v2851
	var v2852 int32
	_ = v2852
	var v2853 int32
	_ = v2853
	var v2858 int32
	_ = v2858
	var v2860 int32
	_ = v2860
	var v2862 int32
	_ = v2862
	var v2863 int32
	_ = v2863
	var v2864 int32
	_ = v2864
	var v2867 int32
	_ = v2867
	var v2868 int32
	_ = v2868
	var v2891 int32
	_ = v2891
	var v2894 int32
	_ = v2894
	var v2896 int32
	_ = v2896
	var v2897 int32
	_ = v2897
	var v2899 int32
	_ = v2899
	var v2900 int32
	_ = v2900
	var v2901 int64
	_ = v2901
	var v2902 int64
	_ = v2902
	var v2904 int32
	_ = v2904
	var v2905 int32
	_ = v2905
	var v2907 int32
	_ = v2907
	var v2910 int32
	_ = v2910
	var v2917 int32
	_ = v2917
	var v2919 int32
	_ = v2919
	var v2920 int32
	_ = v2920
	var v2921 int32
	_ = v2921
	var v2922 int32
	_ = v2922
	var v2926 int32
	_ = v2926
	var v2932 int32
	_ = v2932
	var v2954 int32
	_ = v2954
	var v2956 int32
	_ = v2956
	var v2961 int32
	_ = v2961
	var v2983 int32
	_ = v2983
	var v2984 int32
	_ = v2984
	var v2990 int32
	_ = v2990
	var v3013 int32
	_ = v3013
	var v3014 int32
	_ = v3014
	var v3018 int32
	_ = v3018
	var v3020 int32
	_ = v3020
	var v3022 int32
	_ = v3022
	var v3024 int64
	_ = v3024
	var v3026 int32
	_ = v3026
	var v3028 int64
	_ = v3028
	var v3029 int32
	_ = v3029
	var v3033 int32
	_ = v3033
	var v3035 int32
	_ = v3035
	var v3055 int32
	_ = v3055
	var v3056 int32
	_ = v3056
	var v3058 int32
	_ = v3058
	var v3059 int32
	_ = v3059
	var v3061 int32
	_ = v3061
	var v3062 int32
	_ = v3062
	var v3063 int32
	_ = v3063
	var v3064 int64
	_ = v3064
	var v3066 int32
	_ = v3066
	var v3077 int64
	_ = v3077
	var v3080 int64
	_ = v3080
	var v3083 int64
	_ = v3083
	var v3087 int32
	_ = v3087
	var v3090 int32
	_ = v3090
	var v3098 int32
	_ = v3098
	var v3099 int32
	_ = v3099
	var v3100 int32
	_ = v3100
	var v3105 int32
	_ = v3105
	var v3112 int32
	_ = v3112
	var v3134 int32
	_ = v3134
	var v3138 int32
	_ = v3138
	var v3142 int32
	_ = v3142
	var v3167 int32
	_ = v3167
	var v3169 int32
	_ = v3169
	var v3171 int32
	_ = v3171
	var v3173 int32
	_ = v3173
	var v3180 int32
	_ = v3180
	var v3182 int32
	_ = v3182
	var v3183 int32
	_ = v3183
	var v3184 int32
	_ = v3184
	var v3189 int32
	_ = v3189
	var v3190 int32
	_ = v3190
	var v3196 int32
	_ = v3196
	var v3201 int32
	_ = v3201
	var v3203 int32
	_ = v3203
	var v3231 int32
	_ = v3231
	var v3233 int32
	_ = v3233
	var v3236 int32
	_ = v3236
	var v3238 int32
	_ = v3238
	var v3274 int32
	_ = v3274
	var v3277 int32
	_ = v3277
	var v3280 int32
	_ = v3280
	var v3282 int32
	_ = v3282
	var v3289 int32
	_ = v3289
	var v3292 int32
	_ = v3292
	var v3294 int32
	_ = v3294
	var v3297 int32
	_ = v3297
	var v3299 int32
	_ = v3299
	var v3312 int32
	_ = v3312
	var v3314 int32
	_ = v3314
	var v3317 int32
	_ = v3317
	var v3321 int32
	_ = v3321
	var v3326 int32
	_ = v3326
	var v3332 int32
	_ = v3332
	var v3333 int32
	_ = v3333
	var v3337 int32
	_ = v3337
	var v3340 int32
	_ = v3340
	var v3344 int32
	_ = v3344
	var v3346 int32
	_ = v3346
	var v3350 int32
	_ = v3350
	var v3354 int32
	_ = v3354
	var v3357 int32
	_ = v3357
	var v3359 int32
	_ = v3359
	var v3360 int32
	_ = v3360
	var v3363 int32
	_ = v3363
	var v3367 int32
	_ = v3367
	var v3369 int32
	_ = v3369
	var v3370 int32
	_ = v3370
	var v3375 int32
	_ = v3375
	var v3376 int32
	_ = v3376
	var v3382 int32
	_ = v3382
	var v3388 int32
	_ = v3388
	var v3390 int32
	_ = v3390
	var v3391 int32
	_ = v3391
	var v3401 int32
	_ = v3401
	var v3402 int32
	_ = v3402
	var v3405 int32
	_ = v3405
	var v3407 int32
	_ = v3407
	var v3409 int32
	_ = v3409
	var v3415 int64
	_ = v3415
	var v3431 int32
	_ = v3431
	var v3433 int32
	_ = v3433
	var v3443 int32
	_ = v3443
	var v3447 int32
	_ = v3447
	var v3452 int32
	_ = v3452
	v26 = m.G0
	v28 = v26 - int32(48)
	m.G0 = v28
	v31 = *(*int32)(unsafe.Add(mBase, _c_F_CommitTransaction[0]))
	v32 = *(*int32)(unsafe.Add(mBase, uint32(v31)+24))
	if v32 == int32(5) {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	v35 = *(*int32)(unsafe.Add(mBase, uint32(v31)+72))
	*(*int32)(unsafe.Add(mBase, uint32(v31)+72)) = v35 + int32(1)
	goto L3
L2:
	;
	goto L3
L3:
	;
	v39 = int32(10)
	goto L6
L4:
	;
	if v79 != 0 {
		goto L17
	} else {
		goto L18
	}
L5:
	;
	goto L4
L6:
	;
	v46 = *(*int32)(unsafe.Add(mBase, _c_F_CommitTransaction[1]))
	v49 = *(*int32)(unsafe.Add(mBase, uint32(v46<<(uint(int32(2))%32))+uint32(_c_F_CommitTransaction[2])))
	goto L9
L7:
	;
	v62 = int32(0)
	goto L14
L9:
	;
	goto L10
L10:
	;
	if int32(0)|base.B2i32(v49 == int32(15)) != 0 {
		goto L7
	} else {
		goto L12
	}
L12:
	;
	if v49 <= v39 {
		v79 = int32(1)
		goto L5
	} else {
		goto L13
	}
L13:
	;
	goto L7
L14:
	;
	v66 = *(*int32)(unsafe.Add(mBase, _c_F_CommitTransaction[3]))
	if v66 != int32(2) {
		v79 = v62
		goto L5
	} else {
		goto L15
	}
L15:
	;
	v70 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_CommitTransaction[4])))
	if v70&int32(1) != 0 {
		v79 = v62
		goto L5
	} else {
		goto L16
	}
L16:
	;
	v76 = *(*int32)(unsafe.Add(mBase, _c_F_CommitTransaction[5]))
	v79 = int32(0) | base.B2i32(v76 <= v39)
	goto L5
L17:
	;
	v83 = *(*int32)(unsafe.Add(mBase, _c_F_CommitTransaction[0]))
	F_ShowTransactionStateRec(m, int32(_a_F_CommitTransaction_0), v83)
	mBase = m.M
	v85 = m.ExcPending
	if v85 != 0 {
		goto L20
	} else {
		goto L21
	}
L18:
	;
	goto L19
L19:
	;
	v86 = *(*int32)(unsafe.Add(mBase, uint32(v31)+20))
	if v86 == int32(2) {
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
	v91 = F_errstart(m, int32(19), int32(0))
	mBase = m.M
	v92 = m.ExcPending
	if v92 != 0 {
		goto L20
	} else {
		goto L24
	}
L24:
	;
	if v91 == int32(0) {
		goto L22
	} else {
		goto L25
	}
L25:
	;
	v95 = *(*int32)(unsafe.Add(mBase, uint32(v31)+20))
	if base.Ui32(v95) <= base.Ui32(int32(5)) {
		goto L26
	} else {
		goto L27
	}
L26:
	;
	v100 = *(*int32)(unsafe.Add(mBase, uint32(v95<<(uint(int32(2))%32))+uint32(_c_F_CommitTransaction[6])))
	v102 = v100
	goto L28
L27:
	;
	v102 = int32(_a_F_CommitTransaction_1)
	goto L28
L28:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v28)+16)) = v102
	F_errmsg_internal(m, int32(_a_F_CommitTransaction_2), v28+int32(16))
	mBase = m.M
	v108 = m.ExcPending
	if v108 != 0 {
		goto L20
	} else {
		goto L29
	}
L29:
	;
	F_errfinish(m, int32(_a_F_CommitTransaction_3), int32(2289), int32(_a_F_CommitTransaction_0))
	mBase = m.M
	v113 = m.ExcPending
	if v113 != 0 {
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
	v141 = m.ExcPending
	if v141 != 0 {
		goto L20
	} else {
		goto L33
	}
L32:
	;
	v146 = *(*int32)(unsafe.Add(mBase, _c_F_CommitTransaction[7]))
	if v146 != 0 {
		goto L36
	} else {
		goto L37
	}
L33:
	;
	v143 = F_PreCommit_Portals(m, int32(0))
	mBase = m.M
	v144 = m.ExcPending
	if v144 != 0 {
		goto L20
	} else {
		goto L34
	}
L34:
	;
	if v143 != 0 {
		goto L31
	} else {
		goto L35
	}
L35:
	;
	goto L32
L36:
	;
	v148 = int32(5)
	if v32 == v148 {
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
	v209 = m.ExcPending
	if v209 != 0 {
		goto L20
	} else {
		goto L46
	}
L39:
	;
	v151 = int32(6)
	goto L41
L40:
	;
	v151 = v148
	goto L41
L41:
	;
	v153 = v146
	goto L42
L42:
	;
	v177 = *(*int32)(unsafe.Add(mBase, uint32(v153)))
	v178 = *(*int32)(unsafe.Add(mBase, uint32(v153)+8))
	v179 = *(*int32)(unsafe.Add(mBase, uint32(v153)+4))
	m.T0[v179].(func(*base.Module, int32, int32))(m, v151, v178)
	mBase = m.M
	v181 = m.ExcPending
	if v181 != 0 {
		goto L20
	} else {
		goto L44
	}
L43:
	;
	goto L38
L44:
	;
	if v177 != 0 {
		v153 = v177
		goto L42
	} else {
		goto L45
	}
L45:
	;
	goto L43
L46:
	;
	v210 = *(*int32)(unsafe.Add(mBase, uint32(v31)+72))
	if v32 == int32(5) {
		goto L49
	} else {
		goto L50
	}
L47:
	;
	F_AfterTriggerEndXact(m)
	mBase = m.M
	v245 = m.ExcPending
	if v245 != 0 {
		goto L20
	} else {
		goto L60
	}
L48:
	;
	v235 = *(*int32)(unsafe.Add(mBase, uint32(v31)+72))
	*(*int32)(unsafe.Add(mBase, uint32(v28))) = v235
	F_errmsg_internal(m, v234, v28)
	mBase = m.M
	v238 = m.ExcPending
	if v238 != 0 {
		goto L20
	} else {
		goto L58
	}
L49:
	;
	if v210 == int32(1) {
		goto L47
	} else {
		goto L52
	}
L50:
	;
	goto L51
L51:
	;
	if v210 == int32(0) {
		goto L47
	} else {
		goto L55
	}
L52:
	;
	v217 = F_errstart(m, int32(19), int32(0))
	mBase = m.M
	v218 = m.ExcPending
	if v218 != 0 {
		goto L20
	} else {
		goto L53
	}
L53:
	;
	if v217 == int32(0) {
		goto L47
	} else {
		goto L54
	}
L54:
	;
	v233 = int32(2337)
	v234 = int32(_a_F_CommitTransaction_4)
	goto L48
L55:
	;
	v227 = F_errstart(m, int32(19), int32(0))
	mBase = m.M
	v228 = m.ExcPending
	if v228 != 0 {
		goto L20
	} else {
		goto L56
	}
L56:
	;
	if v227 == int32(0) {
		goto L47
	} else {
		goto L57
	}
L57:
	;
	v233 = int32(2343)
	v234 = int32(_a_F_CommitTransaction_5)
	goto L48
L58:
	;
	F_errfinish(m, int32(_a_F_CommitTransaction_3), v233, int32(_a_F_CommitTransaction_0))
	mBase = m.M
	v242 = m.ExcPending
	if v242 != 0 {
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
	v247 = m.ExcPending
	if v247 != 0 {
		goto L20
	} else {
		goto L61
	}
L61:
	;
	v250 = base.B2i32(v32 == int32(5))
	F_smgrDoPendingSyncs(m, int32(1), v250)
	mBase = m.M
	v252 = m.ExcPending
	if v252 != 0 {
		goto L20
	} else {
		goto L62
	}
L62:
	;
	F_AtEOXact_LargeObject(m, int32(1))
	mBase = m.M
	v255 = m.ExcPending
	if v255 != 0 {
		goto L20
	} else {
		goto L63
	}
L63:
	;
	v256 = m.G0
	v258 = v256 - int32(_a_F_CommitTransaction_6)
	m.G0 = v258
	v261 = *(*int32)(unsafe.Add(mBase, _c_F_CommitTransaction[8]))
	v263 = *(*int32)(unsafe.Add(mBase, _c_F_CommitTransaction[9]))
	if v261|v263 == int32(0) {
		goto L66
	} else {
		goto L67
	}
L64:
	;
	if v250 == int32(0) {
		goto L364
	} else {
		goto L365
	}
L65:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v1924 = m.ExcPending
	if v1924 != 0 {
		goto L20
	} else {
		goto L360
	}
L66:
	;
	m.G0 = v258 + int32(_a_F_CommitTransaction_6)
	goto L64
L67:
	;
	v268 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_CommitTransaction[10])))
	if v268 != int32(1) {
		goto L68
	} else {
		goto L69
	}
L68:
	;
	v287 = *(*int32)(unsafe.Add(mBase, _c_F_CommitTransaction[11]))
	v288 = *(*int32)(unsafe.Add(mBase, uint32(v287)+60))
	if v288 != 0 {
		goto L75
	} else {
		goto L76
	}
L69:
	;
	v273 = F_errstart(m, int32(14), int32(0))
	mBase = m.M
	v274 = m.ExcPending
	if v274 != 0 {
		goto L20
	} else {
		goto L70
	}
L70:
	;
	if v273 == int32(0) {
		goto L68
	} else {
		goto L71
	}
L71:
	;
	F_errmsg_internal(m, int32(_a_F_CommitTransaction_7), int32(0))
	mBase = m.M
	v280 = m.ExcPending
	if v280 != 0 {
		goto L20
	} else {
		goto L72
	}
L72:
	;
	F_errfinish(m, int32(_a_F_CommitTransaction_8), int32(1193), int32(_a_F_CommitTransaction_7))
	mBase = m.M
	v285 = m.ExcPending
	if v285 != 0 {
		goto L20
	} else {
		goto L73
	}
L73:
	;
	goto L68
L74:
	;
	v376 = *(*int32)(unsafe.Add(mBase, _c_F_CommitTransaction[8]))
	if v376 == int32(0) {
		goto L95
	} else {
		goto L96
	}
L75:
	;
	v290 = *(*int32)(unsafe.Add(mBase, _c_F_CommitTransaction[12]))
	if v290 != 0 {
		goto L74
	} else {
		goto L78
	}
L76:
	;
	goto L77
L77:
	;
	v292 = *(*int32)(unsafe.Add(mBase, _c_F_CommitTransaction[13]))
	v296 = F_LWLockAcquire(m, v292+int32(3456), int32(0))
	mBase = m.M
	v297 = m.ExcPending
	if v297 != 0 {
		goto L20
	} else {
		goto L79
	}
L78:
	;
	goto L77
L79:
	;
	v298 = int32(_a_F_CommitTransaction_9)
	v299 = *(*int32)(unsafe.Add(mBase, _c_F_CommitTransaction[14]))
	v302 = *(*int32)(unsafe.Add(mBase, _c_F_CommitTransaction[15]))
	*(*int32)(unsafe.Add(mBase, _c_F_CommitTransaction[14])) = v302
	v305 = *(*int32)(unsafe.Add(mBase, _c_F_CommitTransaction[11]))
	v306 = *(*int32)(unsafe.Add(mBase, uint32(v305)+60))
	if v306 == int32(0) {
		goto L81
	} else {
		goto L82
	}
L80:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_CommitTransaction[14])) = v299
	v368 = *(*int32)(unsafe.Add(mBase, _c_F_CommitTransaction[13]))
	F_LWLockRelease(m, v368+int32(3456))
	mBase = m.M
	v372 = m.ExcPending
	if v372 != 0 {
		goto L20
	} else {
		goto L94
	}
L81:
	;
	v313 = F_dsa_create_ext(m, int32(64), int32(_a_F_CommitTransaction_10), int32(134217728))
	mBase = m.M
	v314 = m.ExcPending
	if v314 != 0 {
		goto L20
	} else {
		goto L84
	}
L82:
	;
	goto L83
L83:
	;
	v345 = *(*int32)(unsafe.Add(mBase, _c_F_CommitTransaction[12]))
	if v345 != 0 {
		goto L80
	} else {
		goto L90
	}
L84:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_CommitTransaction[16])) = v313
	F_dsa_pin(m, v313)
	mBase = m.M
	v317 = m.ExcPending
	if v317 != 0 {
		goto L20
	} else {
		goto L85
	}
L85:
	;
	v319 = *(*int32)(unsafe.Add(mBase, _c_F_CommitTransaction[16]))
	F_dsa_pin_mapping(m, v319)
	mBase = m.M
	v321 = m.ExcPending
	if v321 != 0 {
		goto L20
	} else {
		goto L86
	}
L86:
	;
	v324 = *(*int32)(unsafe.Add(mBase, _c_F_CommitTransaction[16]))
	v327 = F_dshash_create(m, v324, int32(_a_F_CommitTransaction_11), int32(0))
	mBase = m.M
	v328 = m.ExcPending
	if v328 != 0 {
		goto L20
	} else {
		goto L87
	}
L87:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_CommitTransaction[12])) = v327
	v331 = *(*int32)(unsafe.Add(mBase, _c_F_CommitTransaction[16]))
	v332 = *(*int32)(unsafe.Add(mBase, uint32(v331)))
	v333 = *(*int32)(unsafe.Add(mBase, uint32(v332)+28))
	goto L88
L88:
	;
	v335 = *(*int32)(unsafe.Add(mBase, _c_F_CommitTransaction[11]))
	*(*int32)(unsafe.Add(mBase, uint32(v335)+56)) = v333
	v338 = *(*int32)(unsafe.Add(mBase, _c_F_CommitTransaction[12]))
	v339 = *(*int32)(unsafe.Add(mBase, uint32(v338)+32))
	v340 = *(*int32)(unsafe.Add(mBase, uint32(v339)))
	goto L89
L89:
	;
	v342 = *(*int32)(unsafe.Add(mBase, _c_F_CommitTransaction[11]))
	*(*int32)(unsafe.Add(mBase, uint32(v342)+60)) = v340
	goto L80
L90:
	;
	v347 = *(*int32)(unsafe.Add(mBase, uint32(v305)+56))
	v348 = F_dsa_attach(m, v347)
	mBase = m.M
	v349 = m.ExcPending
	if v349 != 0 {
		goto L20
	} else {
		goto L91
	}
L91:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_CommitTransaction[16])) = v348
	F_dsa_pin_mapping(m, v348)
	mBase = m.M
	v352 = m.ExcPending
	if v352 != 0 {
		goto L20
	} else {
		goto L92
	}
L92:
	;
	v355 = *(*int32)(unsafe.Add(mBase, _c_F_CommitTransaction[16]))
	v358 = *(*int32)(unsafe.Add(mBase, _c_F_CommitTransaction[11]))
	v359 = *(*int32)(unsafe.Add(mBase, uint32(v358)+60))
	v361 = F_dshash_attach(m, v355, int32(_a_F_CommitTransaction_11), v359, int32(0))
	mBase = m.M
	v362 = m.ExcPending
	if v362 != 0 {
		goto L20
	} else {
		goto L93
	}
L93:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_CommitTransaction[12])) = v361
	goto L80
L94:
	;
	goto L74
L95:
	;
	v1162 = *(*int32)(unsafe.Add(mBase, _c_F_CommitTransaction[9]))
	if v1162 == int32(0) {
		goto L66
	} else {
		goto L226
	}
L96:
	;
	v380 = *(*int32)(unsafe.Add(mBase, _c_F_CommitTransaction[17]))
	if v380 == int32(0) {
		goto L97
	} else {
		goto L98
	}
L97:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v258)+72)) = int64(274877907008)
	v391 = F_hash_create(m, int32(_a_F_CommitTransaction_12), int64(64), v258-int32(-64), int32(24))
	mBase = m.M
	v392 = m.ExcPending
	if v392 != 0 {
		goto L20
	} else {
		goto L100
	}
L98:
	;
	v396 = v376
	goto L99
L99:
	;
	v398 = *(*int32)(unsafe.Add(mBase, _c_F_CommitTransaction[18]))
	if v398 != 0 {
		goto L101
	} else {
		goto L102
	}
L100:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_CommitTransaction[17])) = v391
	v395 = *(*int32)(unsafe.Add(mBase, _c_F_CommitTransaction[8]))
	v396 = v395
	goto L99
L101:
	;
	v419 = v396
	goto L103
L102:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v258)+72)) = int64(292057776192)
	v402 = *(*int32)(unsafe.Add(mBase, _c_F_CommitTransaction[19]))
	*(*int32)(unsafe.Add(mBase, uint32(v258)+100)) = v402
	v406 = *(*int32)(unsafe.Add(mBase, uint32(v396)+4))
	if v406 != 0 {
		goto L104
	} else {
		goto L105
	}
L103:
	;
	v420 = *(*int32)(unsafe.Add(mBase, uint32(v419)+4))
	if v420 == int32(0) {
		goto L95
	} else {
		goto L108
	}
L104:
	;
	v407 = int64(*(*int32)(unsafe.Add(mBase, uint32(v406)+4)))
	v409 = v407
	goto L106
L105:
	;
	v409 = int64(0)
	goto L106
L106:
	;
	v413 = F_hash_create(m, int32(_a_F_CommitTransaction_13), v409, v258-int32(-64), int32(1048))
	mBase = m.M
	v414 = m.ExcPending
	if v414 != 0 {
		goto L20
	} else {
		goto L107
	}
L107:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_CommitTransaction[18])) = v413
	v417 = *(*int32)(unsafe.Add(mBase, _c_F_CommitTransaction[8]))
	v419 = v417
	goto L103
L108:
	;
	v423 = *(*int32)(unsafe.Add(mBase, uint32(v420)+4))
	if v423 <= int32(0) {
		goto L95
	} else {
		goto L109
	}
L109:
	;
	v427 = v258 + int32(68)
	v433 = int32(0)
	goto L110
L110:
	;
	v453 = *(*int32)(unsafe.Add(mBase, uint32(v420)+12))
	v457 = *(*int32)(unsafe.Add(mBase, uint32(v453+v433<<(uint(int32(2))%32))))
	v458 = *(*int32)(unsafe.Add(mBase, uint32(v457)))
	switch v458 {
	case 0:
		goto L115
	case 1:
		goto L114
	case 2:
		goto L113
	default:
		goto L112
	}
L111:
	;
	goto L95
L112:
	;
	v1133 = v433 + int32(1)
	v1134 = *(*int32)(unsafe.Add(mBase, uint32(v420)+4))
	if v1133 < v1134 {
		v433 = v1133
		goto L110
	} else {
		goto L225
	}
L113:
	;
	v1061 = v258 - int32(-64)
	v1063 = *(*int32)(unsafe.Add(mBase, _c_F_CommitTransaction[17]))
	F_hash_seq_init(m, v1061, v1063)
	mBase = m.M
	v1065 = m.ExcPending
	if v1065 != 0 {
		goto L20
	} else {
		goto L217
	}
L114:
	;
	v1043 = *(*int32)(unsafe.Add(mBase, _c_F_CommitTransaction[17]))
	v1045 = v457 + int32(4)
	v1046 = int32(0)
	v1048 = F_hash_search(m, v1043, v1045, v1046, v1046)
	mBase = m.M
	v1049 = m.ExcPending
	if v1049 != 0 {
		goto L20
	} else {
		goto L214
	}
L115:
	;
	v460 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_CommitTransaction[20])))
	if v460 != 0 {
		goto L116
	} else {
		goto L117
	}
L116:
	;
	v718 = *(*int32)(unsafe.Add(mBase, _c_F_CommitTransaction[18]))
	v720 = v457 + int32(4)
	v723 = F_hash_search(m, v718, v720, int32(1), int32(0))
	mBase = m.M
	v724 = m.ExcPending
	if v724 != 0 {
		goto L20
	} else {
		goto L153
	}
L117:
	;
	v462 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_CommitTransaction[10])))
	if v462 != int32(1) {
		goto L118
	} else {
		goto L119
	}
L118:
	;
	v485 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_CommitTransaction[21])))
	if v485 == int32(0) {
		goto L124
	} else {
		goto L125
	}
L119:
	;
	v467 = F_errstart(m, int32(14), int32(0))
	mBase = m.M
	v468 = m.ExcPending
	if v468 != 0 {
		goto L20
	} else {
		goto L120
	}
L120:
	;
	if v467 == int32(0) {
		goto L118
	} else {
		goto L121
	}
L121:
	;
	v472 = *(*int32)(unsafe.Add(mBase, _c_F_CommitTransaction[22]))
	*(*int32)(unsafe.Add(mBase, uint32(v258)+32)) = v472
	F_errmsg_internal(m, int32(_a_F_CommitTransaction_14), v258+int32(32))
	mBase = m.M
	v478 = m.ExcPending
	if v478 != 0 {
		goto L20
	} else {
		goto L122
	}
L122:
	;
	F_errfinish(m, int32(_a_F_CommitTransaction_8), int32(1462), int32(_a_F_CommitTransaction_15))
	mBase = m.M
	v483 = m.ExcPending
	if v483 != 0 {
		goto L20
	} else {
		goto L123
	}
L123:
	;
	goto L118
L124:
	;
	F_before_shmem_exit(m, int32(549), int64(0))
	mBase = m.M
	v491 = m.ExcPending
	if v491 != 0 {
		goto L20
	} else {
		goto L127
	}
L125:
	;
	goto L126
L126:
	;
	v496 = *(*int32)(unsafe.Add(mBase, _c_F_CommitTransaction[13]))
	v500 = F_LWLockAcquire(m, v496+int32(3456), int32(0))
	mBase = m.M
	v501 = m.ExcPending
	if v501 != 0 {
		goto L20
	} else {
		goto L128
	}
L127:
	;
	v493 = int32(1)
	*(*uint8)(unsafe.Add(mBase, _c_F_CommitTransaction[21])) = uint8(v493)
	goto L126
L128:
	;
	v502 = int32(-1)
	v504 = *(*int32)(unsafe.Add(mBase, _c_F_CommitTransaction[23]))
	v506 = *(*int32)(unsafe.Add(mBase, _c_F_CommitTransaction[11]))
	v507 = *(*int32)(unsafe.Add(mBase, uint32(v506)+28))
	v508 = *(*int32)(unsafe.Add(mBase, uint32(v506)+24))
	v509 = *(*int64)(unsafe.Add(mBase, uint32(v506)+16))
	v510 = *(*int32)(unsafe.Add(mBase, uint32(v506)+8))
	v511 = *(*int64)(unsafe.Add(mBase, uint32(v506)))
	v512 = *(*int32)(unsafe.Add(mBase, uint32(v506)+40))
	if v512 != v502 {
		goto L129
	} else {
		goto L130
	}
L129:
	;
	v518 = *(*int32)(unsafe.Add(mBase, _c_F_CommitTransaction[24]))
	v519 = v512
	v523 = v508
	v525 = v502
	v528 = v507
	v540 = v509
	goto L132
L130:
	;
	v575 = v508
	v577 = v502
	v580 = v507
	v592 = v509
	goto L131
L131:
	;
	v596 = int32(40)
	v598 = v506 + v504*v596
	*(*int32)(unsafe.Add(mBase, uint32(v598)+92)) = v580
	*(*int32)(unsafe.Add(mBase, uint32(v598)+88)) = v575
	*(*int64)(unsafe.Add(mBase, uint32(v598)+80)) = v592
	v602 = int32(_a_F_CommitTransaction_16)
	v603 = *(*int32)(unsafe.Add(mBase, _c_F_CommitTransaction[11]))
	v604 = int32(_a_F_CommitTransaction_17)
	v605 = *(*int32)(unsafe.Add(mBase, _c_F_CommitTransaction[23]))
	v612 = *(*int32)(unsafe.Add(mBase, _c_F_CommitTransaction[22]))
	*(*int32)(unsafe.Add(mBase, uint32(v603+v605*v596-int32(-64)))) = v612
	v615 = *(*int32)(unsafe.Add(mBase, _c_F_CommitTransaction[11]))
	v617 = *(*int32)(unsafe.Add(mBase, _c_F_CommitTransaction[23]))
	v622 = *(*int32)(unsafe.Add(mBase, _c_F_CommitTransaction[24]))
	*(*int32)(unsafe.Add(mBase, uint32(v615+v617*v596)+68)) = v622
	v625 = *(*int32)(unsafe.Add(mBase, _c_F_CommitTransaction[11]))
	v627 = *(*int32)(unsafe.Add(mBase, _c_F_CommitTransaction[23]))
	v631 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v625+v627*v596)+96)) = uint8(v631)
	v634 = *(*int32)(unsafe.Add(mBase, _c_F_CommitTransaction[11]))
	v636 = *(*int32)(unsafe.Add(mBase, _c_F_CommitTransaction[23]))
	*(*uint8)(unsafe.Add(mBase, uint32(v634+v636*v596)+97)) = uint8(v631)
	v643 = *(*int32)(unsafe.Add(mBase, _c_F_CommitTransaction[11]))
	if v577 != int32(-1) {
		goto L147
	} else {
		goto L148
	}
L132:
	;
	v545 = v519 * int32(40)
	v546 = v506 - int32(-64) + v545
	v547 = *(*int32)(unsafe.Add(mBase, uint32(v546)+4))
	if v547 != v518 {
		v560 = v523
		v561 = v528
		v563 = v540
		goto L134
	} else {
		goto L135
	}
L133:
	;
	v575 = v560
	v577 = v566
	v580 = v561
	v592 = v563
	goto L131
L134:
	;
	if v519 < v504 {
		goto L142
	} else {
		goto L143
	}
L135:
	;
	v549 = v506 + v545
	v550 = *(*int64)(unsafe.Add(mBase, uint32(v546)+16))
	if v540 < v550 {
		goto L137
	} else {
		goto L138
	}
L136:
	;
	v558 = *(*int32)(unsafe.Add(mBase, uint32(v549)+92))
	v560 = v556
	v561 = v558
	v563 = v557
	goto L134
L137:
	;
	v552 = *(*int32)(unsafe.Add(mBase, uint32(v549)+88))
	v556 = v552
	v557 = v550
	goto L136
L138:
	;
	goto L139
L139:
	;
	if v540 != v550 {
		v560 = v523
		v561 = v528
		v563 = v540
		goto L134
	} else {
		goto L140
	}
L140:
	;
	v554 = *(*int32)(unsafe.Add(mBase, uint32(v549)+88))
	if v554 < v523 {
		v560 = v523
		v561 = v528
		v563 = v540
		goto L134
	} else {
		goto L141
	}
L141:
	;
	v556 = v554
	v557 = v540
	goto L136
L142:
	;
	v566 = v519
	goto L144
L143:
	;
	v566 = v525
	goto L144
L144:
	;
	v568 = *(*int32)(unsafe.Add(mBase, uint32(v506+v545)+72))
	if v568 != int32(-1) {
		v519 = v568
		v523 = v560
		v525 = v566
		v528 = v561
		v540 = v563
		goto L132
	} else {
		goto L145
	}
L145:
	;
	goto L133
L146:
	;
	v679 = *(*int32)(unsafe.Add(mBase, _c_F_CommitTransaction[13]))
	F_LWLockRelease(m, v679+int32(3456))
	mBase = m.M
	v683 = m.ExcPending
	if v683 != 0 {
		goto L20
	} else {
		goto L150
	}
L147:
	;
	v647 = v643 - int32(-64)
	v648 = int32(_a_F_CommitTransaction_17)
	v649 = *(*int32)(unsafe.Add(mBase, _c_F_CommitTransaction[23]))
	v650 = int32(40)
	v654 = v577 * v650
	v656 = *(*int32)(unsafe.Add(mBase, uint32(v647+v654)+8))
	*(*int32)(unsafe.Add(mBase, uint32(v647+v649*v650)+8)) = v656
	v659 = *(*int32)(unsafe.Add(mBase, _c_F_CommitTransaction[11]))
	v662 = *(*int32)(unsafe.Add(mBase, _c_F_CommitTransaction[23]))
	*(*int32)(unsafe.Add(mBase, uint32(v659+v654)+72)) = v662
	goto L146
L148:
	;
	goto L149
L149:
	;
	v664 = int32(_a_F_CommitTransaction_17)
	v665 = *(*int32)(unsafe.Add(mBase, _c_F_CommitTransaction[23]))
	v669 = *(*int32)(unsafe.Add(mBase, uint32(v643)+40))
	*(*int32)(unsafe.Add(mBase, uint32(v643+v665*int32(40))+72)) = v669
	v672 = *(*int32)(unsafe.Add(mBase, _c_F_CommitTransaction[11]))
	v674 = *(*int32)(unsafe.Add(mBase, _c_F_CommitTransaction[23]))
	*(*int32)(unsafe.Add(mBase, uint32(v672)+40)) = v674
	goto L146
L150:
	;
	v685 = int32(1)
	*(*uint8)(unsafe.Add(mBase, _c_F_CommitTransaction[20])) = uint8(v685)
	if base.B2i32(v575 == v510)&base.B2i32(v592 == v511) != 0 {
		goto L116
	} else {
		goto L151
	}
L151:
	;
	F_asyncQueueReadAllNotifications(m)
	mBase = m.M
	v691 = m.ExcPending
	if v691 != 0 {
		goto L20
	} else {
		goto L152
	}
L152:
	;
	goto L116
L153:
	;
	v725 = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v723)+64)) = v725
	v728 = *(*int32)(unsafe.Add(mBase, _c_F_CommitTransaction[17]))
	v731 = F_hash_search(m, v728, v720, int32(1), v725)
	mBase = m.M
	v732 = m.ExcPending
	if v732 != 0 {
		goto L20
	} else {
		goto L154
	}
L154:
	;
	v734 = *(*int32)(unsafe.Add(mBase, _c_F_CommitTransaction[24]))
	v735 = int64(0)
	*(*int64)(unsafe.Add(mBase, uint32(v427)+56)) = v735
	*(*int64)(unsafe.Add(mBase, uint32(v427)+48)) = v735
	*(*int64)(unsafe.Add(mBase, uint32(v427)+40)) = v735
	*(*int64)(unsafe.Add(mBase, uint32(v427)+32)) = v735
	*(*int64)(unsafe.Add(mBase, uint32(v427)+24)) = v735
	*(*int64)(unsafe.Add(mBase, uint32(v427)+16)) = v735
	*(*int64)(unsafe.Add(mBase, uint32(v427)+8)) = v735
	*(*int64)(unsafe.Add(mBase, uint32(v427))) = v735
	*(*int32)(unsafe.Add(mBase, uint32(v258)+64)) = v734
	goto L158
L155:
	;
	v872 = *(*int32)(unsafe.Add(mBase, _c_F_CommitTransaction[12]))
	v877 = F_dshash_find_or_insert_extended(m, v872, v258-int32(-64), v258+int32(48))
	mBase = m.M
	v878 = m.ExcPending
	if v878 != 0 {
		goto L20
	} else {
		goto L186
	}
L156:
	;
	v868 = F_strlen(m, v857)
	mBase = m.M
	goto L155
L158:
	;
	goto L159
L159:
	;
	v758 = int32(63)
	if (v427^v720)&int32(3) != 0 {
		goto L163
	} else {
		goto L164
	}
L160:
	;
	v861 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v858))) = uint8(v861)
	goto L156
L161:
	;
	v842 = v837
	v843 = v838
	v844 = v839
	goto L182
L162:
	;
	if v832 == int32(0) {
		v857 = v830
		v858 = v831
		goto L160
	} else {
		goto L181
	}
L163:
	;
	v830 = v720
	v831 = v427
	v832 = v758
	goto L162
L164:
	;
	goto L165
L165:
	;
	v762 = int32(0)
	if base.B2i32(v720&int32(3) == v762)|int32(0) == v762 {
		goto L167
	} else {
		goto L168
	}
L166:
	;
	if v798 == int32(0) {
		v857 = v795
		v858 = v796
		goto L160
	} else {
		goto L175
	}
L167:
	;
	v774 = v720
	v775 = v427
	v776 = v758
	goto L170
L168:
	;
	goto L169
L169:
	;
	v795 = v720
	v796 = v427
	v797 = v758
	v798 = int32(1)
	goto L166
L170:
	;
	v778 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v774))))
	*(*uint8)(unsafe.Add(mBase, uint32(v775))) = uint8(v778)
	if v778 == int32(0) {
		v837 = v774
		v838 = v775
		v839 = v776
		goto L161
	} else {
		goto L172
	}
L171:
	;
	v795 = v789
	v796 = v783
	v797 = v785
	v798 = v787
	goto L166
L172:
	;
	v782 = int32(1)
	v783 = v775 + v782
	v785 = v776 - v782
	v786 = int32(0)
	v787 = base.B2i32(v785 != v786)
	v789 = v774 + v782
	if v789&int32(3) == v786 {
		v795 = v789
		v796 = v783
		v797 = v785
		v798 = v787
		goto L166
	} else {
		goto L173
	}
L173:
	;
	if v785 != 0 {
		v774 = v789
		v775 = v783
		v776 = v785
		goto L170
	} else {
		goto L174
	}
L174:
	;
	goto L171
L175:
	;
	v801 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v795))))
	if base.B2i32(v801 == int32(0))|base.B2i32(base.Ui32(v797) < base.Ui32(int32(4))) != 0 {
		v830 = v795
		v831 = v796
		v832 = v797
		goto L162
	} else {
		goto L176
	}
L176:
	;
	v808 = v795
	v809 = v796
	v810 = v797
	goto L177
L177:
	;
	v813 = *(*int32)(unsafe.Add(mBase, uint32(v808)))
	v816 = int32(-2139062144)
	if (int32(16843008)-v813|v813)&v816 != v816 {
		v837 = v808
		v838 = v809
		v839 = v810
		goto L161
	} else {
		goto L179
	}
L178:
	;
	v830 = v824
	v831 = v822
	v832 = v826
	goto L162
L179:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v809))) = v813
	v821 = int32(4)
	v822 = v809 + v821
	v824 = v808 + v821
	v826 = v810 - v821
	if base.Ui32(int32(3)) < base.Ui32(v826) {
		v808 = v824
		v809 = v822
		v810 = v826
		goto L177
	} else {
		goto L180
	}
L180:
	;
	goto L178
L181:
	;
	v837 = v830
	v838 = v831
	v839 = v832
	goto L161
L182:
	;
	v846 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v842))))
	*(*uint8)(unsafe.Add(mBase, uint32(v843))) = uint8(v846)
	if v846 == int32(0) {
		v857 = v842
		v858 = v843
		goto L160
	} else {
		goto L184
	}
L183:
	;
	v857 = v853
	v858 = v851
	goto L160
L184:
	;
	v850 = int32(1)
	v851 = v843 + v850
	v853 = v842 + v850
	v855 = v844 - v850
	if v855 != 0 {
		v842 = v853
		v843 = v851
		v844 = v855
		goto L182
	} else {
		goto L185
	}
L185:
	;
	goto L183
L186:
	;
	v879 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v258)+48)))
	if v879 == int32(0) {
		goto L189
	} else {
		goto L190
	}
L187:
	;
	v899 = *(*int32)(unsafe.Add(mBase, _c_F_CommitTransaction[16]))
	v900 = F_dsa_get_address(m, v899, v897)
	mBase = m.M
	v901 = m.ExcPending
	if v901 != 0 {
		goto L20
	} else {
		goto L194
	}
L188:
	;
	v889 = *(*int32)(unsafe.Add(mBase, _c_F_CommitTransaction[16]))
	v892 = F_dsa_allocate_extended(m, v889, int32(32), int32(0))
	mBase = m.M
	v893 = m.ExcPending
	if v893 != 0 {
		goto L20
	} else {
		goto L193
	}
L189:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v877)+76)) = int32(0)
	*(*int64)(unsafe.Add(mBase, uint32(v877)+68)) = int64(0)
	goto L188
L190:
	;
	goto L191
L191:
	;
	v886 = *(*int32)(unsafe.Add(mBase, uint32(v877)+68))
	if v886 != 0 {
		v897 = v886
		goto L187
	} else {
		goto L192
	}
L192:
	;
	goto L188
L193:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v877)+76)) = int32(4)
	*(*int32)(unsafe.Add(mBase, uint32(v877)+68)) = v892
	v897 = v892
	goto L187
L194:
	;
	v902 = *(*int32)(unsafe.Add(mBase, uint32(v877)+72))
	if int32(0) < v902 {
		goto L196
	} else {
		goto L197
	}
L195:
	;
	v1039 = *(*int32)(unsafe.Add(mBase, _c_F_CommitTransaction[12]))
	F_dshash_release_lock(m, v1039, v877)
	mBase = m.M
	v1041 = m.ExcPending
	if v1041 != 0 {
		goto L20
	} else {
		goto L213
	}
L196:
	;
	v907 = *(*int32)(unsafe.Add(mBase, _c_F_CommitTransaction[23]))
	v908 = int32(0)
	goto L199
L197:
	;
	goto L198
L198:
	;
	v966 = *(*int32)(unsafe.Add(mBase, uint32(v877)+76))
	if v902 < v966 {
		goto L204
	} else {
		goto L205
	}
L199:
	;
	v936 = *(*int32)(unsafe.Add(mBase, uint32(v900+v908<<(uint(int32(3))%32))))
	if v936 == v907 {
		goto L195
	} else {
		goto L201
	}
L200:
	;
	goto L198
L201:
	;
	v939 = v908 + int32(1)
	if v939 != v902 {
		v908 = v939
		goto L199
	} else {
		goto L202
	}
L202:
	;
	goto L200
L203:
	;
	v997 = int32(3)
	v1001 = *(*int32)(unsafe.Add(mBase, _c_F_CommitTransaction[23]))
	*(*int32)(unsafe.Add(mBase, uint32(v993+v994<<(uint(v997)%32)))) = v1001
	v1003 = *(*int32)(unsafe.Add(mBase, uint32(v877)+72))
	v1007 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(v993+v1003<<(uint(v997)%32))+4)) = uint8(v1007)
	v1009 = *(*int32)(unsafe.Add(mBase, uint32(v877)+72))
	*(*int32)(unsafe.Add(mBase, uint32(v877)+72)) = v1009 + v1007
	goto L195
L204:
	;
	v993 = v900
	v994 = v902
	goto L203
L205:
	;
	goto L206
L206:
	;
	v968 = *(*int32)(unsafe.Add(mBase, uint32(v877)+68))
	v970 = *(*int32)(unsafe.Add(mBase, _c_F_CommitTransaction[16]))
	v974 = F_dsa_allocate_extended(m, v970, v966<<(uint(int32(4))%32), int32(0))
	mBase = m.M
	v975 = m.ExcPending
	if v975 != 0 {
		goto L20
	} else {
		goto L207
	}
L207:
	;
	v977 = *(*int32)(unsafe.Add(mBase, _c_F_CommitTransaction[16]))
	v978 = F_dsa_get_address(m, v977, v974)
	mBase = m.M
	v979 = m.ExcPending
	if v979 != 0 {
		goto L20
	} else {
		goto L208
	}
L208:
	;
	v980 = *(*int32)(unsafe.Add(mBase, uint32(v877)+72))
	v982 = v980 << (uint(int32(3)) % 32)
	if v982 != 0 {
		goto L209
	} else {
		goto L210
	}
L209:
	;
	base.MemoryCopy(m, v978, v900, v982)
	goto L211
L210:
	;
	goto L211
L211:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v877)+76)) = v966 << (uint(int32(1)) % 32)
	*(*int32)(unsafe.Add(mBase, uint32(v877)+68)) = v974
	v989 = *(*int32)(unsafe.Add(mBase, _c_F_CommitTransaction[16]))
	F_dsa_free(m, v989, v968)
	mBase = m.M
	v991 = m.ExcPending
	if v991 != 0 {
		goto L20
	} else {
		goto L212
	}
L212:
	;
	v992 = *(*int32)(unsafe.Add(mBase, uint32(v877)+72))
	v993 = v978
	v994 = v992
	goto L203
L213:
	;
	goto L112
L214:
	;
	if v1048 == int32(0) {
		goto L112
	} else {
		goto L215
	}
L215:
	;
	v1053 = *(*int32)(unsafe.Add(mBase, _c_F_CommitTransaction[18]))
	v1056 = F_hash_search(m, v1053, v1045, int32(1), int32(0))
	mBase = m.M
	v1057 = m.ExcPending
	if v1057 != 0 {
		goto L20
	} else {
		goto L216
	}
L216:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v1056)+64)) = int32(1)
	goto L112
L217:
	;
	v1066 = F_hash_seq_search(m, v1061)
	mBase = m.M
	v1067 = m.ExcPending
	if v1067 != 0 {
		goto L20
	} else {
		goto L218
	}
L218:
	;
	if v1066 == int32(0) {
		goto L112
	} else {
		goto L219
	}
L219:
	;
	v1070 = v1066
	goto L220
L220:
	;
	v1096 = *(*int32)(unsafe.Add(mBase, _c_F_CommitTransaction[18]))
	v1099 = F_hash_search(m, v1096, v1070, int32(1), int32(0))
	mBase = m.M
	v1100 = m.ExcPending
	if v1100 != 0 {
		goto L20
	} else {
		goto L222
	}
L221:
	;
	goto L112
L222:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v1099)+64)) = int32(1)
	v1105 = F_hash_seq_search(m, v258-int32(-64))
	mBase = m.M
	v1106 = m.ExcPending
	if v1106 != 0 {
		goto L20
	} else {
		goto L223
	}
L223:
	;
	if v1105 != 0 {
		v1070 = v1105
		goto L220
	} else {
		goto L224
	}
L224:
	;
	goto L221
L225:
	;
	goto L111
L226:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v1162)+12)) = int32(0)
	v1167 = *(*int32)(unsafe.Add(mBase, uint32(v1162)+16))
	if v1167 != 0 {
		goto L228
	} else {
		goto L229
	}
L227:
	;
	v1409 = *(*int32)(unsafe.Add(mBase, _c_F_CommitTransaction[25]))
	if v1409 == int32(0) {
		goto L263
	} else {
		goto L264
	}
L228:
	;
	v1169 = v258 - int32(-64)
	F_hash_seq_init(m, v1169, v1167)
	mBase = m.M
	v1171 = m.ExcPending
	if v1171 != 0 {
		goto L20
	} else {
		goto L231
	}
L229:
	;
	goto L230
L230:
	;
	v1213 = *(*int32)(unsafe.Add(mBase, uint32(v1162)+4))
	if v1213 == int32(0) {
		goto L227
	} else {
		goto L239
	}
L231:
	;
	v1172 = F_hash_seq_search(m, v1169)
	mBase = m.M
	v1173 = m.ExcPending
	if v1173 != 0 {
		goto L20
	} else {
		goto L232
	}
L232:
	;
	if v1172 == int32(0) {
		goto L227
	} else {
		goto L233
	}
L233:
	;
	v1176 = v1172
	goto L234
L234:
	;
	v1202 = *(*int32)(unsafe.Add(mBase, _c_F_CommitTransaction[9]))
	v1203 = *(*int32)(unsafe.Add(mBase, uint32(v1202)+12))
	v1204 = F_lappend(m, v1203, v1176)
	mBase = m.M
	v1205 = m.ExcPending
	if v1205 != 0 {
		goto L20
	} else {
		goto L236
	}
L235:
	;
	goto L227
L236:
	;
	v1207 = *(*int32)(unsafe.Add(mBase, _c_F_CommitTransaction[9]))
	*(*int32)(unsafe.Add(mBase, uint32(v1207)+12)) = v1204
	v1211 = F_hash_seq_search(m, v258-int32(-64))
	mBase = m.M
	v1212 = m.ExcPending
	if v1212 != 0 {
		goto L20
	} else {
		goto L237
	}
L237:
	;
	if v1211 != 0 {
		v1176 = v1211
		goto L234
	} else {
		goto L238
	}
L238:
	;
	goto L235
L239:
	;
	v1216 = *(*int32)(unsafe.Add(mBase, uint32(v1213)+4))
	if v1216 <= int32(0) {
		goto L227
	} else {
		goto L240
	}
L240:
	;
	v1223 = int32(0)
	v1225 = v1162
	goto L241
L241:
	;
	v1245 = *(*int32)(unsafe.Add(mBase, uint32(v1213)+12))
	v1249 = *(*int32)(unsafe.Add(mBase, uint32(v1245+v1223<<(uint(int32(2))%32))))
	v1251 = v1249 + int32(4)
	v1252 = *(*int32)(unsafe.Add(mBase, uint32(v1225)+12))
	if v1252 == int32(0) {
		goto L244
	} else {
		goto L245
	}
L242:
	;
	goto L227
L243:
	;
	v1380 = v1223 + int32(1)
	v1381 = *(*int32)(unsafe.Add(mBase, uint32(v1213)+4))
	if v1380 < v1381 {
		v1223 = v1380
		v1225 = v1359
		goto L241
	} else {
		goto L262
	}
L244:
	;
	v1349 = F_lappend(m, v1252, v1251)
	mBase = m.M
	v1350 = m.ExcPending
	if v1350 != 0 {
		goto L20
	} else {
		goto L261
	}
L245:
	;
	v1255 = *(*int32)(unsafe.Add(mBase, uint32(v1252)+4))
	if v1255 <= int32(0) {
		goto L244
	} else {
		goto L246
	}
L246:
	;
	v1258 = int32(0)
	if v1258 < v1255 {
		goto L247
	} else {
		goto L248
	}
L247:
	;
	v1261 = v1255
	goto L249
L248:
	;
	v1261 = v1258
	goto L249
L249:
	;
	v1262 = *(*int32)(unsafe.Add(mBase, uint32(v1252)+12))
	v1264 = int32(0)
	goto L250
L250:
	;
	v1292 = *(*int32)(unsafe.Add(mBase, uint32(v1262+v1264<<(uint(int32(2))%32))))
	v1295 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1292))))
	v1298 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1251))))
	if base.B2i32(v1295 == int32(0))|base.B2i32(v1295 != v1298) != 0 {
		v1316 = v1295
		v1317 = v1298
		goto L253
	} else {
		goto L254
	}
L251:
	;
	goto L244
L252:
	;
	if v1316-v1317 == int32(0) {
		v1359 = v1225
		goto L243
	} else {
		goto L259
	}
L253:
	;
	goto L252
L254:
	;
	v1301 = v1292
	v1302 = v1251
	goto L255
L255:
	;
	v1305 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1302)+1)))
	v1306 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1301)+1)))
	if v1306 == int32(0) {
		v1316 = v1306
		v1317 = v1305
		goto L253
	} else {
		goto L257
	}
L256:
	;
	v1316 = v1306
	v1317 = v1305
	goto L253
L257:
	;
	v1309 = int32(1)
	if v1306 == v1305 {
		v1301 = v1301 + v1309
		v1302 = v1302 + v1309
		goto L255
	} else {
		goto L258
	}
L258:
	;
	goto L256
L259:
	;
	v1322 = v1264 + int32(1)
	if v1322 != v1261 {
		v1264 = v1322
		goto L250
	} else {
		goto L260
	}
L260:
	;
	goto L251
L261:
	;
	v1352 = *(*int32)(unsafe.Add(mBase, _c_F_CommitTransaction[9]))
	*(*int32)(unsafe.Add(mBase, uint32(v1352)+12)) = v1349
	v1359 = v1352
	goto L243
L262:
	;
	goto L242
L263:
	;
	v1414 = *(*int32)(unsafe.Add(mBase, _c_F_CommitTransaction[15]))
	v1416 = *(*int32)(unsafe.Add(mBase, _c_F_CommitTransaction[26]))
	v1419 = F_MemoryContextAlloc(m, v1414, v1416<<(uint(int32(2))%32))
	mBase = m.M
	v1420 = m.ExcPending
	if v1420 != 0 {
		goto L20
	} else {
		goto L266
	}
L264:
	;
	goto L265
L265:
	;
	v1423 = *(*int32)(unsafe.Add(mBase, _c_F_CommitTransaction[27]))
	if v1423 == int32(0) {
		goto L267
	} else {
		goto L268
	}
L266:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_CommitTransaction[25])) = v1419
	goto L265
L267:
	;
	v1428 = *(*int32)(unsafe.Add(mBase, _c_F_CommitTransaction[15]))
	v1430 = *(*int32)(unsafe.Add(mBase, _c_F_CommitTransaction[26]))
	v1433 = F_MemoryContextAlloc(m, v1428, v1430<<(uint(int32(2))%32))
	mBase = m.M
	v1434 = m.ExcPending
	if v1434 != 0 {
		goto L20
	} else {
		goto L270
	}
L268:
	;
	goto L269
L269:
	;
	v1436 = F_GetCurrentTransactionId(m)
	mBase = m.M
	v1437 = m.ExcPending
	if v1437 != 0 {
		goto L20
	} else {
		goto L271
	}
L270:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_CommitTransaction[27])) = v1433
	goto L269
L271:
	;
	F_LockSharedObject(m, int32(1262), int32(0), int32(8))
	mBase = m.M
	v1442 = m.ExcPending
	if v1442 != 0 {
		goto L20
	} else {
		goto L272
	}
L272:
	;
	v1444 = int64(0)
	*(*int64)(unsafe.Add(mBase, _c_F_CommitTransaction[28])) = v1444
	*(*int64)(unsafe.Add(mBase, _c_F_CommitTransaction[29])) = v1444
	v1450 = int32(0)
	*(*int32)(unsafe.Add(mBase, _c_F_CommitTransaction[30])) = v1450
	*(*int32)(unsafe.Add(mBase, _c_F_CommitTransaction[31])) = v1450
	v1456 = *(*int32)(unsafe.Add(mBase, _c_F_CommitTransaction[9]))
	v1457 = *(*int32)(unsafe.Add(mBase, uint32(v1456)+4))
	if v1457 == v1450 {
		goto L66
	} else {
		goto L273
	}
L273:
	;
	v1460 = *(*int32)(unsafe.Add(mBase, uint32(v1457)+12))
	if v1460 == int32(0) {
		goto L66
	} else {
		goto L274
	}
L274:
	;
	v1469 = int32(1)
	v1470 = v1460
	goto L275
L275:
	;
	v1492 = *(*int32)(unsafe.Add(mBase, _c_F_CommitTransaction[13]))
	v1496 = F_LWLockAcquire(m, v1492+int32(3456), int32(0))
	mBase = m.M
	v1497 = m.ExcPending
	if v1497 != 0 {
		goto L20
	} else {
		goto L277
	}
L276:
	;
	goto L66
L277:
	;
	v1499 = *(*int32)(unsafe.Add(mBase, _c_F_CommitTransaction[11]))
	if v1469&int32(1) != 0 {
		goto L278
	} else {
		goto L279
	}
L278:
	;
	v1503 = *(*int64)(unsafe.Add(mBase, uint32(v1499)+8))
	*(*int64)(unsafe.Add(mBase, _c_F_CommitTransaction[30])) = v1503
	v1506 = *(*int64)(unsafe.Add(mBase, uint32(v1499)))
	*(*int64)(unsafe.Add(mBase, _c_F_CommitTransaction[28])) = v1506
	goto L280
L279:
	;
	goto L280
L280:
	;
	v1508 = *(*int64)(unsafe.Add(mBase, uint32(v1499)))
	v1509 = *(*int64)(unsafe.Add(mBase, uint32(v1499)+16))
	if v1508 == v1509 {
		v1665 = v1499
		goto L281
	} else {
		goto L282
	}
L281:
	;
	v1691 = int64(*(*int32)(unsafe.Add(mBase, _c_F_CommitTransaction[32])))
	v1692 = *(*int64)(unsafe.Add(mBase, uint32(v1665)))
	v1693 = *(*int64)(unsafe.Add(mBase, uint32(v1665)+16))
	if v1691 <= v1692-v1693 {
		goto L65
	} else {
		goto L317
	}
L282:
	;
	v1514 = *(*int32)(unsafe.Add(mBase, _c_F_CommitTransaction[32]))
	v1516 = base.F64_div(base.F64_convert_i64_s(v1508-v1509), base.F64_convert_i32_s(v1514))
	if base.F64_lt(v1516, float64(0.5)) != 0 {
		v1665 = v1499
		goto L281
	} else {
		goto L283
	}
L283:
	;
	v1522 = m.G0
	v1523 = int32(16)
	v1524 = v1522 - v1523
	m.G0 = v1524
	F_gettimeofday(m, v1524)
	mBase = m.M
	v1527 = *(*int64)(unsafe.Add(mBase, uint32(v1524)))
	v1528 = int64(*(*int32)(unsafe.Add(mBase, uint32(v1524)+8)))
	m.G0 = v1524 + v1523
	v1536 = v1528 + v1527*int64(1000000) - int64(946684800000000)
	goto L284
L284:
	;
	v1538 = *(*int32)(unsafe.Add(mBase, _c_F_CommitTransaction[11]))
	v1539 = *(*int64)(unsafe.Add(mBase, uint32(v1538)+48))
	goto L285
L285:
	;
	v1547 = *(*int32)(unsafe.Add(mBase, _c_F_CommitTransaction[11]))
	if base.B2i32(base.I64_extend_i32_s(int32(_a_F_CommitTransaction_18))*int64(1000) <= v1536-v1539) == int32(0) {
		v1665 = v1547
		goto L281
	} else {
		goto L286
	}
L286:
	;
	v1550 = int32(-1)
	v1551 = *(*int32)(unsafe.Add(mBase, uint32(v1547)+40))
	if v1551 != v1550 {
		goto L287
	} else {
		goto L288
	}
L287:
	;
	v1555 = v1547 - int32(-64)
	v1556 = *(*int64)(unsafe.Add(mBase, uint32(v1547)))
	v1557 = *(*int32)(unsafe.Add(mBase, uint32(v1547)+8))
	v1559 = v1557
	v1563 = v1550
	v1564 = v1551
	v1579 = v1556
	goto L290
L288:
	;
	v1615 = v1550
	goto L289
L289:
	;
	v1637 = F_errstart(m, int32(19), int32(0))
	mBase = m.M
	v1638 = m.ExcPending
	if v1638 != 0 {
		goto L20
	} else {
		goto L306
	}
L290:
	;
	v1584 = v1564 * int32(40)
	v1585 = v1547 + v1584
	v1586 = *(*int64)(unsafe.Add(mBase, uint32(v1585)+80))
	if v1586 <= v1579 {
		goto L294
	} else {
		goto L295
	}
L291:
	;
	v1615 = v1605
	goto L289
L292:
	;
	v1607 = *(*int32)(unsafe.Add(mBase, uint32(v1585)+72))
	if v1607 != int32(-1) {
		v1559 = v1603
		v1563 = v1605
		v1564 = v1607
		v1579 = v1606
		goto L290
	} else {
		goto L305
	}
L293:
	;
	v1602 = *(*int32)(unsafe.Add(mBase, uint32(v1584+v1555)))
	v1603 = v1599
	v1605 = v1602
	v1606 = v1600
	goto L292
L294:
	;
	v1588 = *(*int32)(unsafe.Add(mBase, uint32(v1585)+88))
	if v1579 != v1586 {
		goto L297
	} else {
		goto L298
	}
L295:
	;
	v1592 = v1559
	goto L296
L296:
	;
	if v1579 != v1586 {
		v1603 = v1592
		v1605 = v1563
		v1606 = v1579
		goto L292
	} else {
		goto L303
	}
L297:
	;
	v1599 = v1588
	v1600 = v1586
	goto L293
L298:
	;
	goto L299
L299:
	;
	if v1559 < v1588 {
		goto L300
	} else {
		goto L301
	}
L300:
	;
	v1591 = v1559
	goto L302
L301:
	;
	v1591 = v1588
	goto L302
L302:
	;
	v1592 = v1591
	goto L296
L303:
	;
	v1596 = *(*int32)(unsafe.Add(mBase, uint32(v1584+v1555)+24))
	if v1592 != v1596 {
		v1603 = v1592
		v1605 = v1563
		v1606 = v1579
		goto L292
	} else {
		goto L304
	}
L304:
	;
	v1599 = v1596
	v1600 = v1579
	goto L293
L305:
	;
	goto L291
L306:
	;
	if v1637 != 0 {
		goto L307
	} else {
		goto L308
	}
L307:
	;
	*(*float64)(unsafe.Add(mBase, uint32(v258)+16)) = base.F64_mul(v1516, float64(100))
	F_errmsg(m, int32(_a_F_CommitTransaction_19), v258+int32(16))
	mBase = m.M
	v1646 = m.ExcPending
	if v1646 != 0 {
		goto L20
	} else {
		goto L310
	}
L308:
	;
	goto L309
L309:
	;
	v1663 = *(*int32)(unsafe.Add(mBase, _c_F_CommitTransaction[11]))
	*(*int64)(unsafe.Add(mBase, uint32(v1663)+48)) = v1536
	v1665 = v1663
	goto L281
L310:
	;
	if v1615 != int32(-1) {
		goto L311
	} else {
		goto L312
	}
L311:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v258))) = v1615
	v1651 = F_errdetail(m, int32(_a_F_CommitTransaction_20), v258)
	mBase = m.M
	v1652 = m.ExcPending
	if v1652 != 0 {
		goto L20
	} else {
		goto L314
	}
L312:
	;
	goto L313
L313:
	;
	F_errfinish(m, int32(_a_F_CommitTransaction_8), int32(2261), int32(_a_F_CommitTransaction_21))
	mBase = m.M
	v1661 = m.ExcPending
	if v1661 != 0 {
		goto L20
	} else {
		goto L316
	}
L314:
	;
	F_errhint(m, int32(_a_F_CommitTransaction_22), int32(0))
	mBase = m.M
	v1656 = m.ExcPending
	if v1656 != 0 {
		goto L20
	} else {
		goto L315
	}
L315:
	;
	goto L313
L316:
	;
	goto L309
L317:
	;
	v1696 = *(*int64)(unsafe.Add(mBase, uint32(v1665)+8))
	*(*int64)(unsafe.Add(mBase, uint32(v258)+56)) = v1696
	v1698 = *(*int64)(unsafe.Add(mBase, uint32(v1665)))
	*(*int64)(unsafe.Add(mBase, uint32(v258)+48)) = v1698
	v1701 = *(*int32)(unsafe.Add(mBase, _c_F_CommitTransaction[33]))
	v1702 = *(*int32)(unsafe.Add(mBase, uint32(v1701)+28))
	v1704 = int64(*(*uint16)(unsafe.Add(mBase, _c_F_CommitTransaction[34])))
	v1705 = base.I64_rem_s(v1698, v1704)
	v1709 = v1702 + base.I32_wrap_i64(v1705)<<(uint(int32(7))%32)
	v1711 = F_LWLockAcquire(m, v1709, int32(0))
	mBase = m.M
	v1712 = m.ExcPending
	if v1712 != 0 {
		goto L20
	} else {
		goto L318
	}
L318:
	;
	if v1698 != int64(0) {
		goto L320
	} else {
		goto L321
	}
L319:
	;
	v1728 = *(*int32)(unsafe.Add(mBase, _c_F_CommitTransaction[33]))
	v1729 = *(*int32)(unsafe.Add(mBase, uint32(v1728)+12))
	v1731 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(v1726+v1729))) = uint8(v1731)
	v1737 = v1470
	goto L326
L320:
	;
	v1724 = F_SimpleLruReadPage(m, int32(_a_F_CommitTransaction_23), v1698, int32(1), v258+int32(48))
	mBase = m.M
	v1725 = m.ExcPending
	if v1725 != 0 {
		goto L20
	} else {
		goto L324
	}
L321:
	;
	v1715 = *(*int32)(unsafe.Add(mBase, uint32(v258)+56))
	if v1715 != 0 {
		goto L320
	} else {
		goto L322
	}
L322:
	;
	v1718 = F_SimpleLruZeroPage(m, int32(_a_F_CommitTransaction_23), int64(0))
	mBase = m.M
	v1719 = m.ExcPending
	if v1719 != 0 {
		goto L20
	} else {
		goto L323
	}
L323:
	;
	v1726 = v1718
	goto L319
L324:
	;
	v1726 = v1724
	goto L319
L325:
	;
	v1872 = *(*int32)(unsafe.Add(mBase, _c_F_CommitTransaction[11]))
	v1873 = *(*int64)(unsafe.Add(mBase, uint32(v258)+56))
	*(*int64)(unsafe.Add(mBase, uint32(v1872)+8)) = v1873
	v1875 = *(*int64)(unsafe.Add(mBase, uint32(v258)+48))
	*(*int64)(unsafe.Add(mBase, uint32(v1872))) = v1875
	F_LWLockRelease(m, v1866)
	mBase = m.M
	v1878 = m.ExcPending
	if v1878 != 0 {
		goto L20
	} else {
		goto L357
	}
L326:
	;
	v1758 = int32(0)
	if v1737 == v1758 {
		goto L328
	} else {
		goto L329
	}
L327:
	;
	v1839 = *(*int32)(unsafe.Add(mBase, _c_F_CommitTransaction[33]))
	v1840 = *(*int32)(unsafe.Add(mBase, uint32(v1839)+28))
	v1842 = int64(*(*uint16)(unsafe.Add(mBase, _c_F_CommitTransaction[34])))
	v1843 = base.I64_rem_s(v1831, v1842)
	v1847 = v1840 + base.I32_wrap_i64(v1843)<<(uint(int32(7))%32)
	if v1709 == v1847 {
		goto L350
	} else {
		goto L351
	}
L328:
	;
	v1866 = v1709
	v1867 = int32(0)
	goto L325
L329:
	;
	goto L330
L330:
	;
	v1762 = *(*int32)(unsafe.Add(mBase, uint32(v1737)))
	v1763 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v1762)+2)))
	v1764 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v1762))))
	v1766 = *(*int32)(unsafe.Add(mBase, _c_F_CommitTransaction[24]))
	*(*int32)(unsafe.Add(mBase, uint32(v258)+68)) = v1766
	v1768 = v1763 + v1764
	v1772 = (v1768 + int32(21)) & int32(_a_F_CommitTransaction_24)
	*(*int32)(unsafe.Add(mBase, uint32(v258)+64)) = v1772
	v1774 = F_GetCurrentTransactionId(m)
	mBase = m.M
	v1775 = m.ExcPending
	if v1775 != 0 {
		goto L20
	} else {
		goto L331
	}
L331:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v258)+72)) = v1774
	v1778 = *(*int32)(unsafe.Add(mBase, _c_F_CommitTransaction[22]))
	*(*int32)(unsafe.Add(mBase, uint32(v258)+76)) = v1778
	v1781 = v1768 + int32(2)
	if v1781 != 0 {
		goto L332
	} else {
		goto L333
	}
L332:
	;
	base.MemoryCopy(m, v258+int32(80), v1762+int32(4), v1781)
	goto L334
L333:
	;
	goto L334
L334:
	;
	v1785 = *(*int32)(unsafe.Add(mBase, uint32(v258)+56))
	if v1785+v1772 <= int32(_a_F_CommitTransaction_25) {
		goto L336
	} else {
		goto L337
	}
L335:
	;
	if v1809 != 0 {
		goto L342
	} else {
		goto L343
	}
L336:
	;
	v1790 = v1737 + int32(4)
	v1793 = *(*int32)(unsafe.Add(mBase, _c_F_CommitTransaction[9]))
	v1794 = *(*int32)(unsafe.Add(mBase, uint32(v1793)+4))
	v1795 = *(*int32)(unsafe.Add(mBase, uint32(v1794)+12))
	v1796 = *(*int32)(unsafe.Add(mBase, uint32(v1794)+4))
	if base.Ui32(v1790) < base.Ui32(v1795+v1796<<(uint(int32(2))%32)) {
		goto L339
	} else {
		goto L340
	}
L337:
	;
	goto L338
L338:
	;
	v1802 = int32(0)
	*(*uint16)(unsafe.Add(mBase, uint32(v258)+80)) = uint16(v1802)
	*(*int64)(unsafe.Add(mBase, uint32(v258)+68)) = int64(0)
	v1807 = int32(_a_F_CommitTransaction_25) - v1785
	*(*int32)(unsafe.Add(mBase, uint32(v258)+64)) = v1807
	v1809 = v1807
	v1810 = v1737
	goto L335
L339:
	;
	v1801 = v1790
	goto L341
L340:
	;
	v1801 = int32(0)
	goto L341
L341:
	;
	v1809 = v1772
	v1810 = v1801
	goto L335
L342:
	;
	v1813 = *(*int32)(unsafe.Add(mBase, _c_F_CommitTransaction[33]))
	v1814 = *(*int32)(unsafe.Add(mBase, uint32(v1813)+4))
	v1818 = *(*int32)(unsafe.Add(mBase, uint32(v1814+v1726<<(uint(int32(2))%32))))
	base.MemoryCopy(m, v1818+v1785, v258-int32(-64), v1809)
	goto L344
L343:
	;
	goto L344
L344:
	;
	v1823 = *(*int64)(unsafe.Add(mBase, uint32(v258)+48))
	v1824 = *(*int32)(unsafe.Add(mBase, uint32(v258)+56))
	v1825 = v1824 + v1809
	v1827 = v1825 - int32(_a_F_CommitTransaction_26)
	v1829 = base.B2i32(base.Ui32(v1827) < base.Ui32(int32(-8193)))
	v1831 = v1823 + base.I64_extend_i32_u(v1829)
	*(*int64)(unsafe.Add(mBase, uint32(v258)+48)) = v1831
	if base.Ui32(v1827) < base.Ui32(int32(-8193)) {
		goto L345
	} else {
		goto L346
	}
L345:
	;
	v1834 = int32(0)
	goto L347
L346:
	;
	v1834 = v1825
	goto L347
L347:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v258)+56)) = v1834
	if base.Ui32(int32(-8194)) < base.Ui32(v1827) {
		v1737 = v1810
		goto L326
	} else {
		goto L348
	}
L348:
	;
	goto L327
L349:
	;
	v1858 = F_SimpleLruZeroPage(m, int32(_a_F_CommitTransaction_23), v1856)
	mBase = m.M
	v1859 = m.ExcPending
	if v1859 != 0 {
		goto L20
	} else {
		goto L355
	}
L350:
	;
	v1855 = v1709
	v1856 = v1831
	goto L349
L351:
	;
	goto L352
L352:
	;
	F_LWLockRelease(m, v1709)
	mBase = m.M
	v1850 = m.ExcPending
	if v1850 != 0 {
		goto L20
	} else {
		goto L353
	}
L353:
	;
	v1852 = F_LWLockAcquire(m, v1847, int32(0))
	mBase = m.M
	v1853 = m.ExcPending
	if v1853 != 0 {
		goto L20
	} else {
		goto L354
	}
L354:
	;
	v1854 = *(*int64)(unsafe.Add(mBase, uint32(v258)+48))
	v1855 = v1847
	v1856 = v1854
	goto L349
L355:
	;
	v1860 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v258)+48)))
	if v1860&int32(3) != 0 {
		v1866 = v1855
		v1867 = v1810
		goto L325
	} else {
		goto L356
	}
L356:
	;
	v1864 = int32(1)
	*(*uint8)(unsafe.Add(mBase, _c_F_CommitTransaction[35])) = uint8(v1864)
	v1866 = v1855
	v1867 = v1810
	goto L325
L357:
	;
	v1881 = *(*int32)(unsafe.Add(mBase, _c_F_CommitTransaction[11]))
	v1882 = *(*int64)(unsafe.Add(mBase, uint32(v1881)))
	*(*int64)(unsafe.Add(mBase, _c_F_CommitTransaction[29])) = v1882
	v1885 = *(*int64)(unsafe.Add(mBase, uint32(v1881)+8))
	*(*int64)(unsafe.Add(mBase, _c_F_CommitTransaction[31])) = v1885
	v1888 = *(*int32)(unsafe.Add(mBase, _c_F_CommitTransaction[13]))
	F_LWLockRelease(m, v1888+int32(3456))
	mBase = m.M
	v1892 = m.ExcPending
	if v1892 != 0 {
		goto L20
	} else {
		goto L358
	}
L358:
	;
	if v1867 != 0 {
		v1469 = v1758
		v1470 = v1867
		goto L275
	} else {
		goto L359
	}
L359:
	;
	goto L276
L360:
	;
	F_errcode(m, int32(261))
	mBase = m.M
	v1927 = m.ExcPending
	if v1927 != 0 {
		goto L20
	} else {
		goto L361
	}
L361:
	;
	F_errmsg(m, int32(_a_F_CommitTransaction_27), int32(0))
	mBase = m.M
	v1931 = m.ExcPending
	if v1931 != 0 {
		goto L20
	} else {
		goto L362
	}
L362:
	;
	F_errfinish(m, int32(_a_F_CommitTransaction_8), int32(1355), int32(_a_F_CommitTransaction_7))
	mBase = m.M
	v1936 = m.ExcPending
	if v1936 != 0 {
		goto L20
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
	F_PreCommit_CheckForSerializationFailure(m)
	mBase = m.M
	v1940 = m.ExcPending
	if v1940 != 0 {
		goto L20
	} else {
		goto L367
	}
L365:
	;
	goto L366
L366:
	;
	v1941 = int32(_a_F_CommitTransaction_28)
	v1943 = *(*int32)(unsafe.Add(mBase, _c_F_CommitTransaction[36]))
	v1944 = int32(1)
	*(*int32)(unsafe.Add(mBase, _c_F_CommitTransaction[36])) = v1943 + v1944
	F_AtEOXact_RelationMap(m, v1944, v250)
	mBase = m.M
	v1949 = m.ExcPending
	if v1949 != 0 {
		goto L20
	} else {
		goto L368
	}
L367:
	;
	goto L366
L368:
	;
	v1950 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v31)+76)) = uint8(v1950)
	*(*int32)(unsafe.Add(mBase, uint32(v31)+72)) = v1950
	*(*int32)(unsafe.Add(mBase, uint32(v31)+20)) = int32(3)
	v1957 = *(*int32)(unsafe.Add(mBase, _c_F_CommitTransaction[37]))
	if v1950 < v1957 {
		goto L369
	} else {
		goto L370
	}
L369:
	;
	F_disable_timeout(m, int32(8))
	mBase = m.M
	v1962 = m.ExcPending
	if v1962 != 0 {
		goto L20
	} else {
		goto L372
	}
L370:
	;
	goto L371
L371:
	;
	if v32 != int32(5) {
		goto L375
	} else {
		goto L376
	}
L372:
	;
	goto L371
L373:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v3443 = m.ExcPending
	if v3443 != 0 {
		goto L20
	} else {
		goto L647
	}
L374:
	;
	v2335 = *(*int32)(unsafe.Add(mBase, _c_F_CommitTransaction[38]))
	F_ProcArrayEndTransaction(m, v2335, v2324)
	mBase = m.M
	v2337 = m.ExcPending
	if v2337 != 0 {
		goto L20
	} else {
		goto L484
	}
L375:
	;
	v1966 = *(*int32)(unsafe.Add(mBase, _c_F_CommitTransaction[39]))
	v1967 = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v28)+40)) = v1967
	*(*int32)(unsafe.Add(mBase, uint32(v28)+36)) = v1967
	*(*uint8)(unsafe.Add(mBase, uint32(v28)+35)) = uint8(v1967)
	v1974 = *(*int32)(unsafe.Add(mBase, _c_F_CommitTransaction[40]))
	if v1974 <= int32(1) {
		goto L379
	} else {
		goto L380
	}
L376:
	;
	goto L377
L377:
	;
	v2306 = *(*int64)(unsafe.Add(mBase, _c_F_CommitTransaction[41]))
	v2308 = *(*int32)(unsafe.Add(mBase, _c_F_CommitTransaction[42]))
	v2309 = int32(68)
	v2310 = v2308 + v2309
	v2313 = base.AtomicRmwXchg32(m, v2308, v2309, int32(1))
	if v2313 != 0 {
		goto L477
	} else {
		goto L478
	}
L378:
	;
	v1989 = F_smgrGetPendingDeletes(m, int32(1), v28+int32(44))
	mBase = m.M
	v1990 = m.ExcPending
	if v1990 != 0 {
		goto L20
	} else {
		goto L384
	}
L379:
	;
	v1978 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_CommitTransaction[43])))
	if v1978&int32(1) == int32(0) {
		goto L378
	} else {
		goto L382
	}
L380:
	;
	goto L381
L381:
	;
	F_LogLogicalInvalidations(m)
	mBase = m.M
	v1984 = m.ExcPending
	if v1984 != 0 {
		goto L20
	} else {
		goto L383
	}
L382:
	;
	goto L381
L383:
	;
	goto L378
L384:
	;
	v1993 = *(*int32)(unsafe.Add(mBase, _c_F_CommitTransaction[0]))
	v1994 = *(*int32)(unsafe.Add(mBase, uint32(v1993)+52))
	if v1994 != 0 {
		goto L385
	} else {
		goto L386
	}
L385:
	;
	v1995 = *(*int32)(unsafe.Add(mBase, uint32(v1993)+48))
	v1996 = v1995
	goto L387
L386:
	;
	v1996 = int32(0)
	goto L387
L387:
	;
	v2000 = F_pgstat_get_transactional_drops(m, int32(1), v28+int32(40))
	mBase = m.M
	v2001 = m.ExcPending
	if v2001 != 0 {
		goto L20
	} else {
		goto L388
	}
L388:
	;
	v2003 = *(*int32)(unsafe.Add(mBase, _c_F_CommitTransaction[40]))
	if int32(0) < v2003 {
		goto L389
	} else {
		goto L390
	}
L389:
	;
	v2010 = F_xactGetCommittedInvalidationMessages(m, v28+int32(36), v28+int32(35))
	mBase = m.M
	v2011 = m.ExcPending
	if v2011 != 0 {
		goto L20
	} else {
		goto L392
	}
L390:
	;
	v2012 = int32(0)
	goto L391
L391:
	;
	v2014 = *(*int64)(unsafe.Add(mBase, _c_F_CommitTransaction[41]))
	if v1966 == int32(0) {
		goto L395
	} else {
		goto L396
	}
L392:
	;
	v2012 = v2010
	goto L391
L393:
	;
	v2296 = *(*int32)(unsafe.Add(mBase, uint32(v28)+44))
	if v2296 != 0 {
		goto L471
	} else {
		goto L472
	}
L394:
	;
	v2164 = *(*int64)(unsafe.Add(mBase, _c_F_CommitTransaction[41]))
	v2166 = *(*int32)(unsafe.Add(mBase, _c_F_CommitTransaction[44]))
	if v2158&base.B2i32(int32(0) < v2166) != 0 {
		goto L426
	} else {
		goto L427
	}
L395:
	;
	if v1989|v2000 != 0 {
		goto L373
	} else {
		goto L398
	}
L396:
	;
	goto L397
L397:
	;
	v2054 = int32(_a_F_CommitTransaction_29)
	v2056 = *(*int32)(unsafe.Add(mBase, _c_F_CommitTransaction[45]))
	*(*int32)(unsafe.Add(mBase, _c_F_CommitTransaction[45])) = v2056 + int32(1)
	v2061 = int32(*(*uint16)(unsafe.Add(mBase, _c_F_CommitTransaction[46])))
	v2063 = *(*int32)(unsafe.Add(mBase, _c_F_CommitTransaction[38]))
	v2064 = *(*int32)(unsafe.Add(mBase, uint32(v2063)+336))
	*(*int32)(unsafe.Add(mBase, uint32(v2063)+336)) = v2064 | int32(5)
	v2068 = int32(0)
	v2071 = base.AtomicRmwOr32(m, v2068, int32(_a_F_CommitTransaction_30), v2068)
	v2073 = *(*int64)(unsafe.Add(mBase, _c_F_CommitTransaction[47]))
	if v2073 == int64(0) {
		goto L407
	} else {
		goto L408
	}
L398:
	;
	if v2012 != 0 {
		goto L399
	} else {
		goto L400
	}
L399:
	;
	v2018 = *(*int32)(unsafe.Add(mBase, uint32(v28)+36))
	v2019 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v28)+35)))
	v2020 = m.G0
	v2022 = v2020 - int32(16)
	m.G0 = v2022
	*(*int32)(unsafe.Add(mBase, uint32(v2022)+8)) = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v2022)+8)) = uint8(v2019)
	v2028 = *(*int32)(unsafe.Add(mBase, _c_F_CommitTransaction[24]))
	*(*int32)(unsafe.Add(mBase, uint32(v2022))) = v2028
	v2031 = *(*int32)(unsafe.Add(mBase, _c_F_CommitTransaction[48]))
	*(*int32)(unsafe.Add(mBase, uint32(v2022)+4)) = v2031
	*(*int32)(unsafe.Add(mBase, uint32(v2022)+12)) = v2012
	F_XLogBeginInsert(m)
	mBase = m.M
	v2035 = m.ExcPending
	if v2035 != 0 {
		goto L20
	} else {
		goto L402
	}
L400:
	;
	goto L401
L401:
	;
	v2051 = int32(0)
	if v2014 != int64(0) {
		v2158 = v2051
		goto L394
	} else {
		goto L406
	}
L402:
	;
	F_XLogRegisterData(m, v2022, int32(16))
	mBase = m.M
	v2038 = m.ExcPending
	if v2038 != 0 {
		goto L20
	} else {
		goto L403
	}
L403:
	;
	F_XLogRegisterData(m, v2018, v2012<<(uint(int32(4))%32))
	mBase = m.M
	v2042 = m.ExcPending
	if v2042 != 0 {
		goto L20
	} else {
		goto L404
	}
L404:
	;
	v2045 = F_XLogInsert(m, int32(8), int32(32))
	mBase = m.M
	v2046 = m.ExcPending
	if v2046 != 0 {
		goto L20
	} else {
		goto L405
	}
L405:
	;
	m.G0 = v2022 + int32(16)
	v2158 = int32(0)
	goto L394
L406:
	;
	v2291 = v2051
	goto L393
L407:
	;
	v2080 = m.G0
	v2081 = int32(16)
	v2082 = v2080 - v2081
	m.G0 = v2082
	F_gettimeofday(m, v2082)
	mBase = m.M
	v2085 = *(*int64)(unsafe.Add(mBase, uint32(v2082)))
	v2086 = int64(*(*int32)(unsafe.Add(mBase, uint32(v2082)+8)))
	m.G0 = v2082 + v2081
	v2094 = v2086 + v2085*int64(1000000) - int64(946684800000000)
	goto L410
L408:
	;
	v2096 = v2073
	goto L409
L409:
	;
	v2099 = *(*int32)(unsafe.Add(mBase, uint32(v28)+44))
	v2100 = *(*int32)(unsafe.Add(mBase, uint32(v28)+40))
	v2101 = *(*int32)(unsafe.Add(mBase, uint32(v28)+36))
	v2102 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v28)+35)))
	v2104 = *(*int32)(unsafe.Add(mBase, _c_F_CommitTransaction[49]))
	v2105 = int32(0)
	v2107 = F_XactLogCommitRecord(m, v2096, v1994, v1996, v1989, v2099, v2000, v2100, v2012, v2101, v2102, v2104, v2105, v2105)
	mBase = m.M
	v2108 = m.ExcPending
	if v2108 != 0 {
		goto L20
	} else {
		goto L411
	}
L410:
	;
	*(*int64)(unsafe.Add(mBase, _c_F_CommitTransaction[47])) = v2094
	v2096 = v2094
	goto L409
L411:
	;
	if base.Ui32(int32(2)) <= base.Ui32((v2061+int32(1))&int32(_a_F_CommitTransaction_31)) {
		goto L413
	} else {
		goto L414
	}
L412:
	;
	v2155 = int32(*(*uint16)(unsafe.Add(mBase, _c_F_CommitTransaction[46])))
	F_TransactionTreeSetCommitTsData(m, v1966, v1994, v1996, v2151, v2155)
	mBase = m.M
	v2157 = m.ExcPending
	if v2157 != 0 {
		goto L20
	} else {
		goto L422
	}
L413:
	;
	v2114 = *(*int64)(unsafe.Add(mBase, _c_F_CommitTransaction[50]))
	v2116 = *(*int64)(unsafe.Add(mBase, _c_F_CommitTransaction[41]))
	F_replorigin_session_advance(m, v2114, v2116)
	mBase = m.M
	v2118 = m.ExcPending
	if v2118 != 0 {
		goto L20
	} else {
		goto L416
	}
L414:
	;
	goto L415
L415:
	;
	v2125 = *(*int64)(unsafe.Add(mBase, _c_F_CommitTransaction[47]))
	if v2125 == int64(0) {
		goto L418
	} else {
		goto L419
	}
L416:
	;
	v2120 = *(*int64)(unsafe.Add(mBase, _c_F_CommitTransaction[51]))
	if v2120 != int64(0) {
		v2151 = v2120
		goto L412
	} else {
		goto L417
	}
L417:
	;
	goto L415
L418:
	;
	v2132 = m.G0
	v2133 = int32(16)
	v2134 = v2132 - v2133
	m.G0 = v2134
	F_gettimeofday(m, v2134)
	mBase = m.M
	v2137 = *(*int64)(unsafe.Add(mBase, uint32(v2134)))
	v2138 = int64(*(*int32)(unsafe.Add(mBase, uint32(v2134)+8)))
	m.G0 = v2134 + v2133
	v2146 = v2138 + v2137*int64(1000000) - int64(946684800000000)
	goto L421
L419:
	;
	v2148 = v2125
	goto L420
L420:
	;
	*(*int64)(unsafe.Add(mBase, _c_F_CommitTransaction[51])) = v2148
	v2151 = v2148
	goto L412
L421:
	;
	*(*int64)(unsafe.Add(mBase, _c_F_CommitTransaction[47])) = v2146
	v2148 = v2146
	goto L420
L422:
	;
	v2158 = base.B2i32(v2014 != int64(0))
	goto L394
L423:
	;
	v2206 = v1994 - int32(1)
	if v2206 < int32(0) {
		v2274 = v1966
		goto L437
	} else {
		goto L438
	}
L424:
	;
	v2191 = *(*int32)(unsafe.Add(mBase, _c_F_CommitTransaction[38]))
	v2192 = *(*int32)(unsafe.Add(mBase, uint32(v2191)+336))
	*(*int32)(unsafe.Add(mBase, uint32(v2191)+336)) = v2192 & int32(-6)
	v2196 = int32(_a_F_CommitTransaction_29)
	v2198 = *(*int32)(unsafe.Add(mBase, _c_F_CommitTransaction[45]))
	*(*int32)(unsafe.Add(mBase, _c_F_CommitTransaction[45])) = v2198 - int32(1)
	goto L423
L425:
	;
	F_XLogSetAsyncXactLSN(m, v2164)
	mBase = m.M
	v2183 = m.ExcPending
	if v2183 != 0 {
		goto L20
	} else {
		goto L433
	}
L426:
	;
	F_XLogFlush(m, v2164)
	mBase = m.M
	v2177 = m.ExcPending
	if v2177 != 0 {
		goto L20
	} else {
		goto L430
	}
L427:
	;
	v2171 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_CommitTransaction[52])))
	if v2171&int32(1) != 0 {
		goto L426
	} else {
		goto L428
	}
L428:
	;
	if v1989 <= int32(0) {
		goto L425
	} else {
		goto L429
	}
L429:
	;
	goto L426
L430:
	;
	if v1966 == int32(0) {
		goto L423
	} else {
		goto L431
	}
L431:
	;
	F_TransactionIdCommitTree(m, v1966, v1994, v1996)
	mBase = m.M
	v2181 = m.ExcPending
	if v2181 != 0 {
		goto L20
	} else {
		goto L432
	}
L432:
	;
	goto L424
L433:
	;
	if v1966 == int32(0) {
		goto L423
	} else {
		goto L434
	}
L434:
	;
	v2187 = *(*int64)(unsafe.Add(mBase, _c_F_CommitTransaction[41]))
	F_TransactionIdAsyncCommitTree(m, v1966, v1994, v1996, v2187)
	mBase = m.M
	v2189 = m.ExcPending
	if v2189 != 0 {
		goto L20
	} else {
		goto L435
	}
L435:
	;
	goto L424
L436:
	;
	if v2158 != 0 {
		goto L467
	} else {
		goto L468
	}
L437:
	;
	goto L436
L438:
	;
	if v1994&int32(1) != 0 {
		goto L439
	} else {
		goto L440
	}
L439:
	;
	v2211 = int32(2)
	v2214 = *(*int32)(unsafe.Add(mBase, uint32(v1996+v2206<<(uint(v2211)%32))))
	if base.B2i32(base.Ui32(v2211) < base.Ui32(v2214))&base.B2i32(base.Ui32(int32(3)) <= base.Ui32(v1966)) == int32(0) {
		goto L444
	} else {
		goto L445
	}
L440:
	;
	v2229 = v1966
	v2231 = v2206
	goto L441
L441:
	;
	if v2206 == int32(0) {
		v2274 = v2229
		goto L437
	} else {
		goto L449
	}
L442:
	;
	v2229 = v2226
	v2231 = v1994 - int32(2)
	goto L441
L443:
	;
	v2226 = v2214
	goto L442
L444:
	;
	if base.Ui32(v1966) < base.Ui32(v2214) {
		goto L443
	} else {
		goto L447
	}
L445:
	;
	goto L446
L446:
	;
	if int32(0) <= v1966-v2214 {
		v2226 = v1966
		goto L442
	} else {
		goto L448
	}
L447:
	;
	v2226 = v1966
	goto L442
L448:
	;
	goto L443
L449:
	;
	v2234 = v2229
	v2235 = v2231
	goto L450
L450:
	;
	v2239 = int32(3)
	v2243 = v1996 + v2235<<(uint(int32(2))%32)
	v2244 = *(*int32)(unsafe.Add(mBase, uint32(v2243)))
	if base.B2i32(base.Ui32(v2234) < base.Ui32(v2239))|base.B2i32(base.Ui32(v2244) < base.Ui32(v2239)) == int32(0) {
		goto L454
	} else {
		goto L455
	}
L451:
	;
	v2274 = v2269
	goto L437
L452:
	;
	v2257 = *(*int32)(unsafe.Add(mBase, uint32(v2243-int32(4))))
	if base.B2i32(base.Ui32(int32(2)) < base.Ui32(v2257))&base.B2i32(base.Ui32(int32(3)) <= base.Ui32(v2254)) == int32(0) {
		goto L461
	} else {
		goto L462
	}
L453:
	;
	v2254 = v2244
	goto L452
L454:
	;
	if v2234-v2244 < int32(0) {
		goto L453
	} else {
		goto L457
	}
L455:
	;
	goto L456
L456:
	;
	if base.Ui32(v2244) <= base.Ui32(v2234) {
		v2254 = v2234
		goto L452
	} else {
		goto L458
	}
L457:
	;
	v2254 = v2234
	goto L452
L458:
	;
	goto L453
L459:
	;
	if int32(1) < v2235 {
		v2234 = v2269
		v2235 = v2235 - int32(2)
		goto L450
	} else {
		goto L466
	}
L460:
	;
	v2269 = v2257
	goto L459
L461:
	;
	if base.Ui32(v2254) < base.Ui32(v2257) {
		goto L460
	} else {
		goto L464
	}
L462:
	;
	goto L463
L463:
	;
	if int32(0) <= v2254-v2257 {
		v2269 = v2254
		goto L459
	} else {
		goto L465
	}
L464:
	;
	v2269 = v2254
	goto L459
L465:
	;
	goto L460
L466:
	;
	goto L451
L467:
	;
	v2280 = *(*int64)(unsafe.Add(mBase, _c_F_CommitTransaction[41]))
	F_SyncRepWaitForLSN(m, v2280, int32(1))
	mBase = m.M
	v2283 = m.ExcPending
	if v2283 != 0 {
		goto L20
	} else {
		goto L470
	}
L468:
	;
	goto L469
L469:
	;
	v2285 = int32(_a_F_CommitTransaction_32)
	v2286 = *(*int64)(unsafe.Add(mBase, _c_F_CommitTransaction[41]))
	*(*int64)(unsafe.Add(mBase, _c_F_CommitTransaction[53])) = v2286
	*(*int64)(unsafe.Add(mBase, _c_F_CommitTransaction[41])) = int64(0)
	v2291 = v2274
	goto L393
L470:
	;
	goto L469
L471:
	;
	F_pfree(m, v2296)
	mBase = m.M
	v2298 = m.ExcPending
	if v2298 != 0 {
		goto L20
	} else {
		goto L474
	}
L472:
	;
	goto L473
L473:
	;
	if v2000 == int32(0) {
		v2324 = v2291
		goto L374
	} else {
		goto L475
	}
L474:
	;
	goto L473
L475:
	;
	v2301 = *(*int32)(unsafe.Add(mBase, uint32(v28)+40))
	F_pfree(m, v2301)
	mBase = m.M
	v2303 = m.ExcPending
	if v2303 != 0 {
		goto L20
	} else {
		goto L476
	}
L476:
	;
	v2324 = v2291
	goto L374
L477:
	;
	F_s_lock(m, v2310, int32(_a_F_CommitTransaction_33))
	mBase = m.M
	v2316 = m.ExcPending
	if v2316 != 0 {
		goto L20
	} else {
		goto L480
	}
L478:
	;
	goto L479
L479:
	;
	v2317 = *(*int64)(unsafe.Add(mBase, uint32(v2308)+72))
	if base.Ui64(v2317) < base.Ui64(v2306) {
		goto L481
	} else {
		goto L482
	}
L480:
	;
	goto L479
L481:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v2308)+72)) = v2306
	goto L483
L482:
	;
	goto L483
L483:
	;
	v2320 = int32(0)
	atomic.StoreUint32((*uint32)(unsafe.Add(mBase, uint32(v2310))), uint32(v2320))
	v2324 = int32(0)
	goto L374
L484:
	;
	v2339 = base.B2i32(v32 == int32(5))
	v2341 = *(*int32)(unsafe.Add(mBase, _c_F_CommitTransaction[7]))
	if v2341 != 0 {
		goto L485
	} else {
		goto L486
	}
L485:
	;
	v2343 = v2341
	goto L488
L486:
	;
	goto L487
L487:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_CommitTransaction[54])) = int32(0)
	v2401 = *(*int32)(unsafe.Add(mBase, _c_F_CommitTransaction[55]))
	v2402 = int32(1)
	F_ResourceOwnerReleaseInternal(m, v2401, v2402, v2402, v2402)
	mBase = m.M
	v2406 = m.ExcPending
	if v2406 != 0 {
		goto L20
	} else {
		goto L492
	}
L488:
	;
	v2367 = *(*int32)(unsafe.Add(mBase, uint32(v2343)))
	v2368 = *(*int32)(unsafe.Add(mBase, uint32(v2343)+8))
	v2369 = *(*int32)(unsafe.Add(mBase, uint32(v2343)+4))
	m.T0[v2369].(func(*base.Module, int32, int32))(m, v2339, v2368)
	mBase = m.M
	v2371 = m.ExcPending
	if v2371 != 0 {
		goto L20
	} else {
		goto L490
	}
L489:
	;
	goto L487
L490:
	;
	if v2367 != 0 {
		v2343 = v2367
		goto L488
	} else {
		goto L491
	}
L491:
	;
	goto L489
L492:
	;
	F_AtEOXact_Aio(m)
	mBase = m.M
	v2408 = m.ExcPending
	if v2408 != 0 {
		goto L20
	} else {
		goto L493
	}
L493:
	;
	F_AtEOXact_RelationCache(m, int32(1))
	mBase = m.M
	v2411 = m.ExcPending
	if v2411 != 0 {
		goto L20
	} else {
		goto L494
	}
L494:
	;
	F_AtEOXact_TypeCache(m)
	mBase = m.M
	v2413 = m.ExcPending
	if v2413 != 0 {
		goto L20
	} else {
		goto L495
	}
L495:
	;
	F_AtEOXact_Inval(m, int32(1))
	mBase = m.M
	v2416 = m.ExcPending
	if v2416 != 0 {
		goto L20
	} else {
		goto L496
	}
L496:
	;
	v2418 = *(*int32)(unsafe.Add(mBase, _c_F_CommitTransaction[56]))
	v2419 = int32(_a_F_CommitTransaction_17)
	v2420 = *(*int32)(unsafe.Add(mBase, _c_F_CommitTransaction[23]))
	v2421 = int32(2)
	v2424 = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v2418+v2420<<(uint(v2421)%32)))) = v2424
	v2427 = *(*int32)(unsafe.Add(mBase, _c_F_CommitTransaction[57]))
	v2429 = *(*int32)(unsafe.Add(mBase, _c_F_CommitTransaction[23]))
	*(*int32)(unsafe.Add(mBase, uint32(v2427+v2429<<(uint(v2421)%32)))) = v2424
	v2436 = int32(_a_F_CommitTransaction_34)
	*(*int32)(unsafe.Add(mBase, _c_F_CommitTransaction[58])) = v2436
	*(*int32)(unsafe.Add(mBase, _c_F_CommitTransaction[59])) = v2436
	*(*int32)(unsafe.Add(mBase, _c_F_CommitTransaction[60])) = v2424
	*(*int32)(unsafe.Add(mBase, _c_F_CommitTransaction[61])) = v2424
	goto L497
L497:
	;
	v2448 = *(*int32)(unsafe.Add(mBase, _c_F_CommitTransaction[55]))
	v2450 = int32(1)
	F_ResourceOwnerReleaseInternal(m, v2448, int32(2), v2450, v2450)
	mBase = m.M
	v2453 = m.ExcPending
	if v2453 != 0 {
		goto L20
	} else {
		goto L498
	}
L498:
	;
	v2455 = *(*int32)(unsafe.Add(mBase, _c_F_CommitTransaction[55]))
	v2457 = int32(1)
	F_ResourceOwnerReleaseInternal(m, v2455, int32(3), v2457, v2457)
	mBase = m.M
	v2460 = m.ExcPending
	if v2460 != 0 {
		goto L20
	} else {
		goto L499
	}
L499:
	;
	F_smgrDoPendingDeletes(m, int32(1))
	mBase = m.M
	v2463 = m.ExcPending
	if v2463 != 0 {
		goto L20
	} else {
		goto L500
	}
L500:
	;
	v2464 = int32(0)
	v2466 = m.G0
	v2468 = v2466 - int32(80)
	m.G0 = v2468
	v2471 = *(*int32)(unsafe.Add(mBase, _c_F_CommitTransaction[8]))
	v2473 = *(*int32)(unsafe.Add(mBase, _c_F_CommitTransaction[9]))
	if v2471|v2473 != 0 {
		goto L501
	} else {
		goto L502
	}
L501:
	;
	v2476 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_CommitTransaction[10])))
	if v2476 != int32(1) {
		goto L504
	} else {
		goto L505
	}
L502:
	;
	goto L503
L503:
	;
	m.G0 = v2468 + int32(80)
	v3274 = int32(1)
	F_AtEOXact_GUC(m, v3274, v3274)
	mBase = m.M
	v3277 = m.ExcPending
	if v3277 != 0 {
		goto L20
	} else {
		goto L622
	}
L504:
	;
	F_ApplyPendingListenActions(m, int32(1))
	mBase = m.M
	v2496 = m.ExcPending
	if v2496 != 0 {
		goto L20
	} else {
		goto L510
	}
L505:
	;
	v2481 = F_errstart(m, int32(14), int32(0))
	mBase = m.M
	v2482 = m.ExcPending
	if v2482 != 0 {
		goto L20
	} else {
		goto L506
	}
L506:
	;
	if v2481 == int32(0) {
		goto L504
	} else {
		goto L507
	}
L507:
	;
	F_errmsg_internal(m, int32(_a_F_CommitTransaction_35), int32(0))
	mBase = m.M
	v2488 = m.ExcPending
	if v2488 != 0 {
		goto L20
	} else {
		goto L508
	}
L508:
	;
	F_errfinish(m, int32(_a_F_CommitTransaction_8), int32(1388), int32(_a_F_CommitTransaction_35))
	mBase = m.M
	v2493 = m.ExcPending
	if v2493 != 0 {
		goto L20
	} else {
		goto L509
	}
L509:
	;
	goto L504
L510:
	;
	v2498 = *(*int32)(unsafe.Add(mBase, _c_F_CommitTransaction[8]))
	if v2498 == int32(0) {
		goto L511
	} else {
		goto L512
	}
L511:
	;
	v2573 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_CommitTransaction[20])))
	if v2573 == int32(0) {
		goto L521
	} else {
		goto L522
	}
L512:
	;
	m.Env.Pgmem_listen(m, int32(_a_F_CommitTransaction_36), int32(2))
	mBase = m.M
	v2505 = *(*int32)(unsafe.Add(mBase, _c_F_CommitTransaction[17]))
	if v2505 == int32(0) {
		goto L511
	} else {
		goto L513
	}
L513:
	;
	v2509 = v2468 + int32(12)
	F_hash_seq_init(m, v2509, v2505)
	mBase = m.M
	v2511 = m.ExcPending
	if v2511 != 0 {
		goto L20
	} else {
		goto L514
	}
L514:
	;
	v2512 = F_hash_seq_search(m, v2509)
	mBase = m.M
	v2513 = m.ExcPending
	if v2513 != 0 {
		goto L20
	} else {
		goto L515
	}
L515:
	;
	if v2512 == int32(0) {
		goto L511
	} else {
		goto L516
	}
L516:
	;
	v2517 = v2512
	goto L517
L517:
	;
	m.Env.Pgmem_listen(m, v2517, int32(1))
	mBase = m.M
	v2545 = F_hash_seq_search(m, v2468+int32(12))
	mBase = m.M
	v2546 = m.ExcPending
	if v2546 != 0 {
		goto L20
	} else {
		goto L519
	}
L518:
	;
	goto L511
L519:
	;
	if v2545 != 0 {
		v2517 = v2545
		goto L517
	} else {
		goto L520
	}
L520:
	;
	goto L518
L521:
	;
	v2653 = *(*int32)(unsafe.Add(mBase, _c_F_CommitTransaction[9]))
	if v2653 == int32(0) {
		goto L532
	} else {
		goto L533
	}
L522:
	;
	v2577 = *(*int32)(unsafe.Add(mBase, _c_F_CommitTransaction[17]))
	if v2577 != 0 {
		goto L523
	} else {
		goto L524
	}
L523:
	;
	v2579 = *(*int32)(unsafe.Add(mBase, uint32(v2577)))
	v2580 = *(*int64)(unsafe.Add(mBase, uint32(v2579)+8))
	v2581 = *(*int64)(unsafe.Add(mBase, uint32(v2579)+808))
	if v2581 != int64(0) {
		goto L527
	} else {
		goto L528
	}
L524:
	;
	goto L525
L525:
	;
	F_asyncQueueUnregister(m)
	mBase = m.M
	v2650 = m.ExcPending
	if v2650 != 0 {
		goto L20
	} else {
		goto L531
	}
L526:
	;
	if v2646 != int64(0) {
		goto L521
	} else {
		goto L530
	}
L527:
	;
	v2584 = *(*int64)(unsafe.Add(mBase, uint32(v2579)+752))
	v2585 = *(*int64)(unsafe.Add(mBase, uint32(v2579)+728))
	v2586 = *(*int64)(unsafe.Add(mBase, uint32(v2579)+704))
	v2587 = *(*int64)(unsafe.Add(mBase, uint32(v2579)+680))
	v2588 = *(*int64)(unsafe.Add(mBase, uint32(v2579)+656))
	v2589 = *(*int64)(unsafe.Add(mBase, uint32(v2579)+632))
	v2590 = *(*int64)(unsafe.Add(mBase, uint32(v2579)+608))
	v2591 = *(*int64)(unsafe.Add(mBase, uint32(v2579)+584))
	v2592 = *(*int64)(unsafe.Add(mBase, uint32(v2579)+560))
	v2593 = *(*int64)(unsafe.Add(mBase, uint32(v2579)+536))
	v2594 = *(*int64)(unsafe.Add(mBase, uint32(v2579)+512))
	v2595 = *(*int64)(unsafe.Add(mBase, uint32(v2579)+488))
	v2596 = *(*int64)(unsafe.Add(mBase, uint32(v2579)+464))
	v2597 = *(*int64)(unsafe.Add(mBase, uint32(v2579)+440))
	v2598 = *(*int64)(unsafe.Add(mBase, uint32(v2579)+416))
	v2599 = *(*int64)(unsafe.Add(mBase, uint32(v2579)+392))
	v2600 = *(*int64)(unsafe.Add(mBase, uint32(v2579)+368))
	v2601 = *(*int64)(unsafe.Add(mBase, uint32(v2579)+344))
	v2602 = *(*int64)(unsafe.Add(mBase, uint32(v2579)+320))
	v2603 = *(*int64)(unsafe.Add(mBase, uint32(v2579)+296))
	v2604 = *(*int64)(unsafe.Add(mBase, uint32(v2579)+272))
	v2605 = *(*int64)(unsafe.Add(mBase, uint32(v2579)+248))
	v2606 = *(*int64)(unsafe.Add(mBase, uint32(v2579)+224))
	v2607 = *(*int64)(unsafe.Add(mBase, uint32(v2579)+200))
	v2608 = *(*int64)(unsafe.Add(mBase, uint32(v2579)+176))
	v2609 = *(*int64)(unsafe.Add(mBase, uint32(v2579)+152))
	v2610 = *(*int64)(unsafe.Add(mBase, uint32(v2579)+128))
	v2611 = *(*int64)(unsafe.Add(mBase, uint32(v2579)+104))
	v2612 = *(*int64)(unsafe.Add(mBase, uint32(v2579)+80))
	v2613 = *(*int64)(unsafe.Add(mBase, uint32(v2579)+56))
	v2614 = *(*int64)(unsafe.Add(mBase, uint32(v2579)+32))
	v2646 = v2584 + (v2585 + (v2586 + (v2587 + (v2588 + (v2589 + (v2590 + (v2591 + (v2592 + (v2593 + (v2594 + (v2595 + (v2596 + (v2597 + (v2598 + (v2599 + (v2600 + (v2601 + (v2602 + (v2603 + (v2604 + (v2605 + (v2606 + (v2607 + (v2608 + (v2609 + (v2610 + (v2611 + (v2612 + (v2613 + (v2614 + v2580))))))))))))))))))))))))))))))
	goto L529
L528:
	;
	v2646 = v2580
	goto L529
L529:
	;
	goto L526
L530:
	;
	goto L525
L531:
	;
	goto L521
L532:
	;
	v3231 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_CommitTransaction[35])))
	if v3231 != 0 {
		goto L618
	} else {
		goto L619
	}
L533:
	;
	v2657 = *(*int32)(unsafe.Add(mBase, _c_F_CommitTransaction[13]))
	v2661 = F_LWLockAcquire(m, v2657+int32(3456), int32(0))
	mBase = m.M
	v2662 = m.ExcPending
	if v2662 != 0 {
		goto L20
	} else {
		goto L534
	}
L534:
	;
	v2664 = *(*int32)(unsafe.Add(mBase, _c_F_CommitTransaction[9]))
	v2665 = *(*int32)(unsafe.Add(mBase, uint32(v2664)+12))
	if v2665 == int32(0) {
		v2990 = v2464
		goto L535
	} else {
		goto L536
	}
L535:
	;
	v3013 = *(*int32)(unsafe.Add(mBase, _c_F_CommitTransaction[11]))
	v3014 = *(*int32)(unsafe.Add(mBase, uint32(v3013)+40))
	if v3014 != int32(-1) {
		goto L590
	} else {
		goto L591
	}
L536:
	;
	v2668 = *(*int32)(unsafe.Add(mBase, uint32(v2665)+4))
	if v2668 <= int32(0) {
		v2990 = v2464
		goto L535
	} else {
		goto L537
	}
L537:
	;
	v2672 = v2468 + int32(16)
	v2677 = v2464
	v2682 = v2464
	goto L538
L538:
	;
	v2698 = *(*int32)(unsafe.Add(mBase, uint32(v2665)+12))
	v2702 = *(*int32)(unsafe.Add(mBase, uint32(v2698+v2682<<(uint(int32(2))%32))))
	v2704 = *(*int32)(unsafe.Add(mBase, _c_F_CommitTransaction[24]))
	v2705 = int64(0)
	*(*int64)(unsafe.Add(mBase, uint32(v2672)+56)) = v2705
	*(*int64)(unsafe.Add(mBase, uint32(v2672)+48)) = v2705
	*(*int64)(unsafe.Add(mBase, uint32(v2672)+40)) = v2705
	*(*int64)(unsafe.Add(mBase, uint32(v2672)+32)) = v2705
	*(*int64)(unsafe.Add(mBase, uint32(v2672)+24)) = v2705
	*(*int64)(unsafe.Add(mBase, uint32(v2672)+16)) = v2705
	*(*int64)(unsafe.Add(mBase, uint32(v2672)+8)) = v2705
	*(*int64)(unsafe.Add(mBase, uint32(v2672))) = v2705
	*(*int32)(unsafe.Add(mBase, uint32(v2468)+12)) = v2704
	goto L543
L539:
	;
	v2990 = v2961
	goto L535
L540:
	;
	v2842 = *(*int32)(unsafe.Add(mBase, _c_F_CommitTransaction[12]))
	v2846 = F_dshash_find(m, v2842, v2468+int32(12), int32(0))
	mBase = m.M
	v2847 = m.ExcPending
	if v2847 != 0 {
		goto L20
	} else {
		goto L571
	}
L541:
	;
	v2838 = F_strlen(m, v2827)
	mBase = m.M
	goto L540
L543:
	;
	goto L544
L544:
	;
	v2728 = int32(63)
	if (v2672^v2702)&int32(3) != 0 {
		goto L548
	} else {
		goto L549
	}
L545:
	;
	v2831 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v2828))) = uint8(v2831)
	goto L541
L546:
	;
	v2812 = v2807
	v2813 = v2808
	v2814 = v2809
	goto L567
L547:
	;
	if v2802 == int32(0) {
		v2827 = v2800
		v2828 = v2801
		goto L545
	} else {
		goto L566
	}
L548:
	;
	v2800 = v2702
	v2801 = v2672
	v2802 = v2728
	goto L547
L549:
	;
	goto L550
L550:
	;
	v2732 = int32(0)
	if base.B2i32(v2702&int32(3) == v2732)|int32(0) == v2732 {
		goto L552
	} else {
		goto L553
	}
L551:
	;
	if v2768 == int32(0) {
		v2827 = v2765
		v2828 = v2766
		goto L545
	} else {
		goto L560
	}
L552:
	;
	v2744 = v2702
	v2745 = v2672
	v2746 = v2728
	goto L555
L553:
	;
	goto L554
L554:
	;
	v2765 = v2702
	v2766 = v2672
	v2767 = v2728
	v2768 = int32(1)
	goto L551
L555:
	;
	v2748 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v2744))))
	*(*uint8)(unsafe.Add(mBase, uint32(v2745))) = uint8(v2748)
	if v2748 == int32(0) {
		v2807 = v2744
		v2808 = v2745
		v2809 = v2746
		goto L546
	} else {
		goto L557
	}
L556:
	;
	v2765 = v2759
	v2766 = v2753
	v2767 = v2755
	v2768 = v2757
	goto L551
L557:
	;
	v2752 = int32(1)
	v2753 = v2745 + v2752
	v2755 = v2746 - v2752
	v2756 = int32(0)
	v2757 = base.B2i32(v2755 != v2756)
	v2759 = v2744 + v2752
	if v2759&int32(3) == v2756 {
		v2765 = v2759
		v2766 = v2753
		v2767 = v2755
		v2768 = v2757
		goto L551
	} else {
		goto L558
	}
L558:
	;
	if v2755 != 0 {
		v2744 = v2759
		v2745 = v2753
		v2746 = v2755
		goto L555
	} else {
		goto L559
	}
L559:
	;
	goto L556
L560:
	;
	v2771 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v2765))))
	if base.B2i32(v2771 == int32(0))|base.B2i32(base.Ui32(v2767) < base.Ui32(int32(4))) != 0 {
		v2800 = v2765
		v2801 = v2766
		v2802 = v2767
		goto L547
	} else {
		goto L561
	}
L561:
	;
	v2778 = v2765
	v2779 = v2766
	v2780 = v2767
	goto L562
L562:
	;
	v2783 = *(*int32)(unsafe.Add(mBase, uint32(v2778)))
	v2786 = int32(-2139062144)
	if (int32(16843008)-v2783|v2783)&v2786 != v2786 {
		v2807 = v2778
		v2808 = v2779
		v2809 = v2780
		goto L546
	} else {
		goto L564
	}
L563:
	;
	v2800 = v2794
	v2801 = v2792
	v2802 = v2796
	goto L547
L564:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v2779))) = v2783
	v2791 = int32(4)
	v2792 = v2779 + v2791
	v2794 = v2778 + v2791
	v2796 = v2780 - v2791
	if base.Ui32(int32(3)) < base.Ui32(v2796) {
		v2778 = v2794
		v2779 = v2792
		v2780 = v2796
		goto L562
	} else {
		goto L565
	}
L565:
	;
	goto L563
L566:
	;
	v2807 = v2800
	v2808 = v2801
	v2809 = v2802
	goto L546
L567:
	;
	v2816 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v2812))))
	*(*uint8)(unsafe.Add(mBase, uint32(v2813))) = uint8(v2816)
	if v2816 == int32(0) {
		v2827 = v2812
		v2828 = v2813
		goto L545
	} else {
		goto L569
	}
L568:
	;
	v2827 = v2823
	v2828 = v2821
	goto L545
L569:
	;
	v2820 = int32(1)
	v2821 = v2813 + v2820
	v2823 = v2812 + v2820
	v2825 = v2814 - v2820
	if v2825 != 0 {
		v2812 = v2823
		v2813 = v2821
		v2814 = v2825
		goto L567
	} else {
		goto L570
	}
L570:
	;
	goto L568
L571:
	;
	if v2846 != 0 {
		goto L572
	} else {
		goto L573
	}
L572:
	;
	v2849 = *(*int32)(unsafe.Add(mBase, _c_F_CommitTransaction[16]))
	v2850 = *(*int32)(unsafe.Add(mBase, uint32(v2846)+68))
	v2851 = F_dsa_get_address(m, v2849, v2850)
	mBase = m.M
	v2852 = m.ExcPending
	if v2852 != 0 {
		goto L20
	} else {
		goto L575
	}
L573:
	;
	v2961 = v2677
	goto L574
L574:
	;
	v2983 = v2682 + int32(1)
	v2984 = *(*int32)(unsafe.Add(mBase, uint32(v2665)+4))
	if v2983 < v2984 {
		v2677 = v2961
		v2682 = v2983
		goto L538
	} else {
		goto L589
	}
L575:
	;
	v2853 = *(*int32)(unsafe.Add(mBase, uint32(v2846)+72))
	if int32(0) < v2853 {
		goto L576
	} else {
		goto L577
	}
L576:
	;
	v2858 = *(*int32)(unsafe.Add(mBase, _c_F_CommitTransaction[11]))
	v2860 = *(*int32)(unsafe.Add(mBase, _c_F_CommitTransaction[27]))
	v2862 = *(*int32)(unsafe.Add(mBase, _c_F_CommitTransaction[25]))
	v2863 = v2858
	v2864 = int32(0)
	v2867 = v2677
	v2868 = v2853
	goto L579
L577:
	;
	v2932 = v2677
	goto L578
L578:
	;
	v2954 = *(*int32)(unsafe.Add(mBase, _c_F_CommitTransaction[12]))
	F_dshash_release_lock(m, v2954, v2846)
	mBase = m.M
	v2956 = m.ExcPending
	if v2956 != 0 {
		goto L20
	} else {
		goto L588
	}
L579:
	;
	v2891 = *(*int32)(unsafe.Add(mBase, uint32(v2851+v2864<<(uint(int32(3))%32))))
	v2894 = v2863 + v2891*int32(40)
	v2896 = v2894 + int32(96)
	v2897 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v2896))))
	if v2897 != 0 {
		v2920 = v2863
		v2921 = v2867
		v2922 = v2868
		goto L581
	} else {
		goto L582
	}
L580:
	;
	v2932 = v2921
	goto L578
L581:
	;
	v2926 = v2864 + int32(1)
	if v2926 < v2922 {
		v2863 = v2920
		v2864 = v2926
		v2867 = v2921
		v2868 = v2922
		goto L579
	} else {
		goto L587
	}
L582:
	;
	v2899 = v2894 - int32(-64)
	v2900 = *(*int32)(unsafe.Add(mBase, uint32(v2899)))
	v2901 = *(*int64)(unsafe.Add(mBase, uint32(v2899)+16))
	v2902 = *(*int64)(unsafe.Add(mBase, uint32(v2863)))
	if v2901 == v2902 {
		goto L583
	} else {
		goto L584
	}
L583:
	;
	v2904 = *(*int32)(unsafe.Add(mBase, uint32(v2899)+24))
	v2905 = *(*int32)(unsafe.Add(mBase, uint32(v2863)+8))
	if v2904 == v2905 {
		v2920 = v2863
		v2921 = v2867
		v2922 = v2868
		goto L581
	} else {
		goto L586
	}
L584:
	;
	goto L585
L585:
	;
	v2907 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(v2896))) = uint8(v2907)
	v2910 = v2867 << (uint(int32(2)) % 32)
	*(*int32)(unsafe.Add(mBase, uint32(v2862+v2910))) = v2900
	*(*int32)(unsafe.Add(mBase, uint32(v2910+v2860))) = v2891
	v2917 = *(*int32)(unsafe.Add(mBase, uint32(v2846)+72))
	v2919 = *(*int32)(unsafe.Add(mBase, _c_F_CommitTransaction[11]))
	v2920 = v2919
	v2921 = v2867 + v2907
	v2922 = v2917
	goto L581
L586:
	;
	goto L585
L587:
	;
	goto L580
L588:
	;
	v2961 = v2932
	goto L574
L589:
	;
	goto L539
L590:
	;
	v3018 = *(*int32)(unsafe.Add(mBase, _c_F_CommitTransaction[27]))
	v3020 = *(*int32)(unsafe.Add(mBase, _c_F_CommitTransaction[25]))
	v3022 = *(*int32)(unsafe.Add(mBase, _c_F_CommitTransaction[31]))
	v3024 = *(*int64)(unsafe.Add(mBase, _c_F_CommitTransaction[29]))
	v3026 = *(*int32)(unsafe.Add(mBase, _c_F_CommitTransaction[30]))
	v3028 = *(*int64)(unsafe.Add(mBase, _c_F_CommitTransaction[28]))
	v3029 = v3013
	v3033 = v2990
	v3035 = v3014
	goto L593
L591:
	;
	v3112 = v2990
	goto L592
L592:
	;
	v3134 = *(*int32)(unsafe.Add(mBase, _c_F_CommitTransaction[13]))
	F_LWLockRelease(m, v3134+int32(3456))
	mBase = m.M
	v3138 = m.ExcPending
	if v3138 != 0 {
		goto L20
	} else {
		goto L603
	}
L593:
	;
	v3055 = v3035 * int32(40)
	v3056 = v3029 + v3055
	v3058 = v3056 + int32(96)
	v3059 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v3058))))
	if v3059 != 0 {
		v3099 = v3029
		v3100 = v3033
		goto L595
	} else {
		goto L596
	}
L594:
	;
	v3112 = v3100
	goto L592
L595:
	;
	v3105 = *(*int32)(unsafe.Add(mBase, uint32(v3099+v3055)+72))
	if v3105 != int32(-1) {
		v3029 = v3099
		v3033 = v3100
		v3035 = v3105
		goto L593
	} else {
		goto L602
	}
L596:
	;
	v3061 = v3056 - int32(-64)
	v3062 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v3061)+33)))
	if v3062 != 0 {
		v3099 = v3029
		v3100 = v3033
		goto L595
	} else {
		goto L597
	}
L597:
	;
	v3063 = *(*int32)(unsafe.Add(mBase, uint32(v3061)))
	v3064 = *(*int64)(unsafe.Add(mBase, uint32(v3061)+16))
	if v3064 < v3028 {
		goto L598
	} else {
		goto L599
	}
L598:
	;
	v3083 = *(*int64)(unsafe.Add(mBase, uint32(v3029)))
	if v3083-v3064 < int64(4) {
		v3099 = v3029
		v3100 = v3033
		goto L595
	} else {
		goto L601
	}
L599:
	;
	v3066 = *(*int32)(unsafe.Add(mBase, uint32(v3061)+24))
	if base.B2i32(v3066 < v3026)&base.B2i32(v3064 == v3028)|(base.B2i32(v3022 <= v3066)|base.B2i32(v3064 != v3024))&base.B2i32(v3024 <= v3064) != 0 {
		goto L598
	} else {
		goto L600
	}
L600:
	;
	v3077 = *(*int64)(unsafe.Add(mBase, _c_F_CommitTransaction[31]))
	*(*int64)(unsafe.Add(mBase, uint32(v3056)+88)) = v3077
	v3080 = *(*int64)(unsafe.Add(mBase, _c_F_CommitTransaction[29]))
	*(*int64)(unsafe.Add(mBase, uint32(v3056)+80)) = v3080
	v3099 = v3029
	v3100 = v3033
	goto L595
L601:
	;
	v3087 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(v3058))) = uint8(v3087)
	v3090 = v3033 << (uint(int32(2)) % 32)
	*(*int32)(unsafe.Add(mBase, uint32(v3020+v3090))) = v3063
	*(*int32)(unsafe.Add(mBase, uint32(v3090+v3018))) = v3035
	v3098 = *(*int32)(unsafe.Add(mBase, _c_F_CommitTransaction[11]))
	v3099 = v3098
	v3100 = v3033 + v3087
	goto L595
L602:
	;
	goto L594
L603:
	;
	if v3112 <= int32(0) {
		goto L532
	} else {
		goto L604
	}
L604:
	;
	v3142 = int32(0)
	goto L605
L605:
	;
	v3167 = v3142 << (uint(int32(2)) % 32)
	v3169 = *(*int32)(unsafe.Add(mBase, _c_F_CommitTransaction[25]))
	v3171 = *(*int32)(unsafe.Add(mBase, uint32(v3167+v3169)))
	v3173 = *(*int32)(unsafe.Add(mBase, _c_F_CommitTransaction[22]))
	if v3171 == v3173 {
		goto L608
	} else {
		goto L609
	}
L606:
	;
	goto L532
L607:
	;
	v3203 = v3142 + int32(1)
	if v3203 != v3112 {
		v3142 = v3203
		goto L605
	} else {
		goto L617
	}
L608:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_CommitTransaction[62])) = int32(1)
	goto L607
L609:
	;
	goto L610
L610:
	;
	v3180 = *(*int32)(unsafe.Add(mBase, _c_F_CommitTransaction[27]))
	v3182 = *(*int32)(unsafe.Add(mBase, uint32(v3180+v3167)))
	v3183 = F_SendProcSignal(m, v3171, int32(1), v3182)
	mBase = m.M
	v3184 = m.ExcPending
	if v3184 != 0 {
		goto L20
	} else {
		goto L611
	}
L611:
	;
	if int32(0) <= v3183 {
		goto L607
	} else {
		goto L612
	}
L612:
	;
	v3189 = F_errstart(m, int32(12), int32(0))
	mBase = m.M
	v3190 = m.ExcPending
	if v3190 != 0 {
		goto L20
	} else {
		goto L613
	}
L613:
	;
	if v3189 == int32(0) {
		goto L607
	} else {
		goto L614
	}
L614:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v2468))) = v3171
	F_errmsg_internal(m, int32(_a_F_CommitTransaction_37), v2468)
	mBase = m.M
	v3196 = m.ExcPending
	if v3196 != 0 {
		goto L20
	} else {
		goto L615
	}
L615:
	;
	F_errfinish(m, int32(_a_F_CommitTransaction_8), int32(2424), int32(_a_F_CommitTransaction_38))
	mBase = m.M
	v3201 = m.ExcPending
	if v3201 != 0 {
		goto L20
	} else {
		goto L616
	}
L616:
	;
	goto L607
L617:
	;
	goto L606
L618:
	;
	v3233 = int32(0)
	*(*uint8)(unsafe.Add(mBase, _c_F_CommitTransaction[35])) = uint8(v3233)
	F_asyncQueueAdvanceTail(m)
	mBase = m.M
	v3236 = m.ExcPending
	if v3236 != 0 {
		goto L20
	} else {
		goto L621
	}
L619:
	;
	goto L620
L620:
	;
	v3238 = int32(0)
	*(*int32)(unsafe.Add(mBase, _c_F_CommitTransaction[9])) = v3238
	*(*int32)(unsafe.Add(mBase, _c_F_CommitTransaction[8])) = v3238
	*(*int32)(unsafe.Add(mBase, _c_F_CommitTransaction[18])) = v3238
	goto L503
L621:
	;
	goto L620
L622:
	;
	F_AtEOXact_SPI(m, int32(1))
	mBase = m.M
	v3280 = m.ExcPending
	if v3280 != 0 {
		goto L20
	} else {
		goto L623
	}
L623:
	;
	v3282 = int32(0)
	*(*int32)(unsafe.Add(mBase, _c_F_CommitTransaction[63])) = v3282
	*(*int32)(unsafe.Add(mBase, _c_F_CommitTransaction[64])) = v3282
	goto L624
L624:
	;
	F_AtEOXact_on_commit_actions(m, int32(1))
	mBase = m.M
	v3289 = m.ExcPending
	if v3289 != 0 {
		goto L20
	} else {
		goto L625
	}
L625:
	;
	F_AtEOXact_Namespace(m, int32(1), v2339)
	mBase = m.M
	v3292 = m.ExcPending
	if v3292 != 0 {
		goto L20
	} else {
		goto L626
	}
L626:
	;
	F_smgrdestroyall(m)
	mBase = m.M
	v3294 = m.ExcPending
	if v3294 != 0 {
		goto L20
	} else {
		goto L627
	}
L627:
	;
	F_AtEOXact_Files(m, int32(1))
	mBase = m.M
	v3297 = m.ExcPending
	if v3297 != 0 {
		goto L20
	} else {
		goto L628
	}
L628:
	;
	v3299 = int32(0)
	*(*int32)(unsafe.Add(mBase, _c_F_CommitTransaction[65])) = v3299
	*(*int32)(unsafe.Add(mBase, _c_F_CommitTransaction[66])) = v3299
	*(*int32)(unsafe.Add(mBase, _c_F_CommitTransaction[67])) = v3299
	*(*int32)(unsafe.Add(mBase, _c_F_CommitTransaction[68])) = v3299
	goto L629
L629:
	;
	F_AtEOXact_HashTables(m, int32(1))
	mBase = m.M
	v3312 = m.ExcPending
	if v3312 != 0 {
		goto L20
	} else {
		goto L630
	}
L630:
	;
	F_AtEOXact_RI(m)
	mBase = m.M
	v3314 = m.ExcPending
	if v3314 != 0 {
		goto L20
	} else {
		goto L631
	}
L631:
	;
	F_AtEOXact_PgStat(m, int32(1), v2339)
	mBase = m.M
	v3317 = m.ExcPending
	if v3317 != 0 {
		goto L20
	} else {
		goto L632
	}
L632:
	;
	F_AtEOXact_Snapshot(m, int32(1), int32(0))
	mBase = m.M
	v3321 = m.ExcPending
	if v3321 != 0 {
		goto L20
	} else {
		goto L633
	}
L633:
	;
	goto L636
L634:
	;
	F_AtEOXact_LogicalRepWorkers(m, int32(1))
	mBase = m.M
	v3344 = m.ExcPending
	if v3344 != 0 {
		goto L20
	} else {
		goto L639
	}
L635:
	;
	v3340 = int32(0)
	*(*uint8)(unsafe.Add(mBase, _c_F_CommitTransaction[69])) = uint8(v3340)
	goto L634
L636:
	;
	v3326 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_CommitTransaction[69])))
	if v3326&int32(1) == int32(0) {
		goto L635
	} else {
		goto L637
	}
L637:
	;
	v3332 = *(*int32)(unsafe.Add(mBase, _c_F_CommitTransaction[70]))
	v3333 = *(*int32)(unsafe.Add(mBase, uint32(v3332)))
	if v3333 == int32(0) {
		goto L635
	} else {
		goto L638
	}
L638:
	;
	v3337 = F_pgmem_kill(m, v3333, int32(10))
	mBase = m.M
	goto L635
L639:
	;
	F_AtEOXact_LogicalCtl(m)
	mBase = m.M
	v3346 = m.ExcPending
	if v3346 != 0 {
		goto L20
	} else {
		goto L640
	}
L640:
	;
	v3350 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_CommitTransaction[71])))
	if v3350 != int32(1) {
		goto L642
	} else {
		goto L643
	}
L641:
	;
	v3388 = *(*int32)(unsafe.Add(mBase, _c_F_CommitTransaction[55]))
	F_ResourceOwnerDelete(m, v3388)
	mBase = m.M
	v3390 = m.ExcPending
	if v3390 != 0 {
		goto L20
	} else {
		goto L645
	}
L642:
	;
	goto L641
L643:
	;
	v3354 = *(*int32)(unsafe.Add(mBase, _c_F_CommitTransaction[72]))
	if v3354 == int32(0) {
		goto L642
	} else {
		goto L644
	}
L644:
	;
	v3357 = int32(_a_F_CommitTransaction_29)
	v3359 = *(*int32)(unsafe.Add(mBase, _c_F_CommitTransaction[45]))
	v3360 = int32(1)
	*(*int32)(unsafe.Add(mBase, _c_F_CommitTransaction[45])) = v3359 + v3360
	v3363 = *(*int32)(unsafe.Add(mBase, uint32(v3354)))
	*(*int32)(unsafe.Add(mBase, uint32(v3354))) = v3363 + v3360
	v3367 = int32(0)
	v3369 = int32(_a_F_CommitTransaction_39)
	v3370 = base.AtomicRmwOr32(m, v3367, v3369, v3367)
	*(*int64)(unsafe.Add(mBase, uint32(v3354)+24)) = int64(0)
	v3375 = base.AtomicRmwOr32(m, v3367, v3369, v3367)
	v3376 = *(*int32)(unsafe.Add(mBase, uint32(v3354)))
	*(*int32)(unsafe.Add(mBase, uint32(v3354))) = v3376 + v3360
	v3382 = *(*int32)(unsafe.Add(mBase, _c_F_CommitTransaction[45]))
	*(*int32)(unsafe.Add(mBase, _c_F_CommitTransaction[45])) = v3382 - v3360
	goto L642
L645:
	;
	v3391 = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v31)+40)) = v3391
	*(*int32)(unsafe.Add(mBase, _c_F_CommitTransaction[55])) = v3391
	*(*int32)(unsafe.Add(mBase, _c_F_CommitTransaction[73])) = v3391
	v3401 = *(*int32)(unsafe.Add(mBase, _c_F_CommitTransaction[0]))
	v3402 = *(*int32)(unsafe.Add(mBase, uint32(v3401)+44))
	*(*int32)(unsafe.Add(mBase, _c_F_CommitTransaction[14])) = v3402
	v3405 = *(*int32)(unsafe.Add(mBase, _c_F_CommitTransaction[74]))
	F_MemoryContextReset(m, v3405)
	mBase = m.M
	v3407 = m.ExcPending
	if v3407 != 0 {
		goto L20
	} else {
		goto L646
	}
L646:
	;
	v3409 = int32(0)
	*(*int32)(unsafe.Add(mBase, _c_F_CommitTransaction[19])) = v3409
	*(*int32)(unsafe.Add(mBase, uint32(v3401)+36)) = v3409
	*(*int32)(unsafe.Add(mBase, uint32(v31)+56)) = v3409
	v3415 = int64(0)
	*(*int64)(unsafe.Add(mBase, uint32(v31)+48)) = v3415
	*(*int64)(unsafe.Add(mBase, uint32(v31)+28)) = v3415
	*(*int32)(unsafe.Add(mBase, uint32(v31)+8)) = v3409
	*(*int64)(unsafe.Add(mBase, uint32(v31))) = v3415
	*(*int64)(unsafe.Add(mBase, _c_F_CommitTransaction[39])) = v3415
	*(*int32)(unsafe.Add(mBase, _c_F_CommitTransaction[75])) = v3409
	*(*int32)(unsafe.Add(mBase, uint32(v31)+20)) = v3409
	v3431 = int32(_a_F_CommitTransaction_28)
	v3433 = *(*int32)(unsafe.Add(mBase, _c_F_CommitTransaction[36]))
	*(*int32)(unsafe.Add(mBase, _c_F_CommitTransaction[36])) = v3433 - int32(1)
	m.G0 = v28 + int32(48)
	return
L647:
	;
	F_errmsg_internal(m, int32(_a_F_CommitTransaction_40), int32(0))
	mBase = m.M
	v3447 = m.ExcPending
	if v3447 != 0 {
		goto L20
	} else {
		goto L648
	}
L648:
	;
	F_errfinish(m, int32(_a_F_CommitTransaction_3), int32(1395), int32(_a_F_CommitTransaction_41))
	mBase = m.M
	v3452 = m.ExcPending
	if v3452 != 0 {
		goto L20
	} else {
		goto L649
	}
L649:
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
	var v41 int32
	_ = v41
	var v42 int64
	_ = v42
	var v43 int64
	_ = v43
	var v44 int32
	_ = v44
	var v51 int32
	_ = v51
	var v54 int32
	_ = v54
	var v67 int32
	_ = v67
	var v71 int32
	_ = v71
	var v75 int32
	_ = v75
	var v81 int32
	_ = v81
	var v84 int32
	_ = v84
	var v88 int32
	_ = v88
	var v90 int32
	_ = v90
	var v91 int32
	_ = v91
	var v96 int32
	_ = v96
	var v97 int32
	_ = v97
	var v100 int32
	_ = v100
	var v105 int32
	_ = v105
	var v107 int32
	_ = v107
	var v111 int32
	_ = v111
	var v116 int32
	_ = v116
	var v148 int32
	_ = v148
	var v150 int32
	_ = v150
	var v151 int32
	_ = v151
	var v153 int32
	_ = v153
	var v155 int32
	_ = v155
	var v183 int32
	_ = v183
	var v185 int32
	_ = v185
	var v186 int32
	_ = v186
	var v188 int32
	_ = v188
	var v219 int32
	_ = v219
	var v221 int32
	_ = v221
	var v225 int32
	_ = v225
	var v228 int32
	_ = v228
	var v230 int32
	_ = v230
	var v232 int32
	_ = v232
	var v238 int32
	_ = v238
	var v239 int32
	_ = v239
	var v241 int32
	_ = v241
	var v248 int32
	_ = v248
	var v253 int32
	_ = v253
	var v257 int32
	_ = v257
	var v258 int32
	_ = v258
	var v259 int32
	_ = v259
	var v262 int64
	_ = v262
	var v263 int64
	_ = v263
	var v273 int32
	_ = v273
	var v275 int32
	_ = v275
	var v277 int32
	_ = v277
	var v279 int32
	_ = v279
	var v281 int32
	_ = v281
	var v283 int32
	_ = v283
	var v287 int32
	_ = v287
	var v291 int32
	_ = v291
	var v297 int32
	_ = v297
	var v299 int32
	_ = v299
	var v302 int32
	_ = v302
	var v306 int32
	_ = v306
	var v307 int32
	_ = v307
	var v309 int32
	_ = v309
	var v310 int32
	_ = v310
	var v316 int32
	_ = v316
	var v347 int32
	_ = v347
	var v349 int32
	_ = v349
	var v352 int32
	_ = v352
	var v355 int32
	_ = v355
	var v358 int32
	_ = v358
	var v359 int32
	_ = v359
	var v362 int32
	_ = v362
	var v363 int32
	_ = v363
	var v366 int32
	_ = v366
	var v373 int32
	_ = v373
	var v374 int32
	_ = v374
	var v377 int32
	_ = v377
	var v382 int32
	_ = v382
	var v385 int32
	_ = v385
	var v391 int32
	_ = v391
	var v396 int32
	_ = v396
	var v426 int32
	_ = v426
	var v429 int32
	_ = v429
	var v432 int32
	_ = v432
	var v435 int32
	_ = v435
	var v436 int32
	_ = v436
	var v445 int32
	_ = v445
	var v449 int32
	_ = v449
	var v456 int32
	_ = v456
	var v459 int32
	_ = v459
	var v463 int32
	_ = v463
	var v468 int32
	_ = v468
	var v472 int32
	_ = v472
	var v475 int32
	_ = v475
	var v479 int32
	_ = v479
	var v483 int32
	_ = v483
	var v488 int32
	_ = v488
	var v492 int32
	_ = v492
	var v495 int32
	_ = v495
	var v499 int32
	_ = v499
	var v501 int32
	_ = v501
	var v507 int32
	_ = v507
	var v512 int32
	_ = v512
	var v514 int32
	_ = v514
	var v516 int32
	_ = v516
	var v518 int32
	_ = v518
	var v521 int32
	_ = v521
	var v522 int32
	_ = v522
	var v523 int32
	_ = v523
	var v524 int64
	_ = v524
	var v531 int32
	_ = v531
	var v532 int32
	_ = v532
	var v537 int32
	_ = v537
	var v540 int32
	_ = v540
	var v541 int32
	_ = v541
	var v543 int32
	_ = v543
	var v559 int32
	_ = v559
	var v561 int64
	_ = v561
	var v563 int32
	_ = v563
	var v568 int32
	_ = v568
	var v569 int32
	_ = v569
	var v570 int32
	_ = v570
	var v572 int32
	_ = v572
	var v574 int32
	_ = v574
	var v579 int32
	_ = v579
	var v580 int32
	_ = v580
	var v585 int32
	_ = v585
	var v586 int32
	_ = v586
	var v591 int32
	_ = v591
	var v592 int32
	_ = v592
	var v597 int32
	_ = v597
	var v598 int32
	_ = v598
	var v602 int32
	_ = v602
	var v603 int32
	_ = v603
	var v606 int32
	_ = v606
	var v607 int32
	_ = v607
	var v608 int64
	_ = v608
	var v613 int32
	_ = v613
	var v616 int32
	_ = v616
	var v620 int32
	_ = v620
	var v621 int32
	_ = v621
	var v623 int32
	_ = v623
	var v624 int32
	_ = v624
	var v625 int32
	_ = v625
	var v626 int32
	_ = v626
	var v633 int32
	_ = v633
	var v635 int32
	_ = v635
	var v637 int32
	_ = v637
	var v642 int32
	_ = v642
	var v643 int32
	_ = v643
	var v645 int32
	_ = v645
	var v648 int32
	_ = v648
	var v650 int32
	_ = v650
	var v651 int32
	_ = v651
	var v652 int32
	_ = v652
	var v653 int32
	_ = v653
	var v657 int32
	_ = v657
	var v659 int32
	_ = v659
	var v661 int32
	_ = v661
	var v665 int32
	_ = v665
	var v667 int32
	_ = v667
	var v669 int32
	_ = v669
	var v671 int32
	_ = v671
	var v673 int32
	_ = v673
	var v677 int32
	_ = v677
	var v679 int32
	_ = v679
	var v681 int32
	_ = v681
	var v682 int32
	_ = v682
	var v683 int32
	_ = v683
	var v684 int32
	_ = v684
	var v691 int32
	_ = v691
	var v694 int32
	_ = v694
	var v696 int32
	_ = v696
	var v698 int32
	_ = v698
	var v702 int32
	_ = v702
	var v703 int32
	_ = v703
	var v705 int32
	_ = v705
	var v707 int32
	_ = v707
	var v709 int32
	_ = v709
	var v711 int32
	_ = v711
	var v713 int32
	_ = v713
	var v714 int32
	_ = v714
	var v715 int32
	_ = v715
	var v716 int32
	_ = v716
	var v717 int32
	_ = v717
	var v720 int32
	_ = v720
	var v721 int32
	_ = v721
	var v724 int32
	_ = v724
	var v727 int32
	_ = v727
	var v729 int32
	_ = v729
	var v732 int32
	_ = v732
	var v734 int32
	_ = v734
	var v738 int32
	_ = v738
	var v740 int32
	_ = v740
	var v742 int32
	_ = v742
	var v743 int32
	_ = v743
	var v744 int32
	_ = v744
	var v745 int32
	_ = v745
	var v752 int32
	_ = v752
	var v755 int32
	_ = v755
	var v757 int32
	_ = v757
	var v759 int32
	_ = v759
	var v763 int32
	_ = v763
	var v764 int32
	_ = v764
	var v766 int32
	_ = v766
	var v768 int32
	_ = v768
	var v770 int32
	_ = v770
	var v772 int32
	_ = v772
	var v773 int32
	_ = v773
	var v774 int32
	_ = v774
	var v775 int32
	_ = v775
	var v778 int32
	_ = v778
	var v782 int32
	_ = v782
	var v784 int32
	_ = v784
	var v786 int32
	_ = v786
	var v787 int32
	_ = v787
	var v790 int32
	_ = v790
	var v791 int32
	_ = v791
	var v792 int32
	_ = v792
	var v795 int32
	_ = v795
	var v796 int32
	_ = v796
	var v797 int32
	_ = v797
	var v800 int32
	_ = v800
	var v805 int32
	_ = v805
	var v807 int32
	_ = v807
	var v817 int32
	_ = v817
	var v819 int32
	_ = v819
	var v821 int32
	_ = v821
	var v824 int32
	_ = v824
	var v826 int32
	_ = v826
	var v830 int32
	_ = v830
	var v833 int32
	_ = v833
	var v834 int32
	_ = v834
	var v836 int32
	_ = v836
	var v837 int32
	_ = v837
	var v838 int32
	_ = v838
	var v839 int32
	_ = v839
	var v846 int32
	_ = v846
	var v849 int32
	_ = v849
	var v851 int32
	_ = v851
	var v853 int32
	_ = v853
	var v857 int32
	_ = v857
	var v858 int32
	_ = v858
	var v860 int32
	_ = v860
	var v863 int32
	_ = v863
	var v865 int32
	_ = v865
	var v866 int32
	_ = v866
	var v867 int32
	_ = v867
	var v868 int32
	_ = v868
	var v869 int32
	_ = v869
	var v870 int32
	_ = v870
	var v873 int32
	_ = v873
	var v882 int32
	_ = v882
	var v884 int32
	_ = v884
	var v892 int32
	_ = v892
	var v895 int32
	_ = v895
	var v897 int32
	_ = v897
	var v901 int32
	_ = v901
	var v903 int32
	_ = v903
	var v906 int32
	_ = v906
	var v907 int32
	_ = v907
	var v909 int32
	_ = v909
	var v910 int32
	_ = v910
	var v911 int32
	_ = v911
	var v912 int32
	_ = v912
	var v919 int32
	_ = v919
	var v922 int32
	_ = v922
	var v924 int32
	_ = v924
	var v926 int32
	_ = v926
	var v930 int32
	_ = v930
	var v931 int32
	_ = v931
	var v933 int32
	_ = v933
	var v936 int32
	_ = v936
	var v937 int32
	_ = v937
	var v938 int32
	_ = v938
	var v939 int32
	_ = v939
	var v940 int32
	_ = v940
	var v943 int32
	_ = v943
	var v949 int32
	_ = v949
	var v951 int32
	_ = v951
	var v954 int32
	_ = v954
	var v956 int32
	_ = v956
	var v963 int32
	_ = v963
	var v966 int32
	_ = v966
	var v968 int32
	_ = v968
	var v970 int32
	_ = v970
	var v973 int32
	_ = v973
	var v974 int32
	_ = v974
	var v976 int32
	_ = v976
	var v977 int32
	_ = v977
	var v978 int32
	_ = v978
	var v979 int32
	_ = v979
	var v986 int32
	_ = v986
	var v989 int32
	_ = v989
	var v991 int32
	_ = v991
	var v993 int32
	_ = v993
	var v997 int32
	_ = v997
	var v998 int32
	_ = v998
	var v1000 int32
	_ = v1000
	var v1003 int32
	_ = v1003
	var v1004 int32
	_ = v1004
	var v1005 int32
	_ = v1005
	var v1006 int32
	_ = v1006
	var v1007 int32
	_ = v1007
	var v1010 int32
	_ = v1010
	var v1016 int32
	_ = v1016
	var v1018 int32
	_ = v1018
	var v1021 int32
	_ = v1021
	var v1023 int32
	_ = v1023
	var v1029 int32
	_ = v1029
	var v1032 int32
	_ = v1032
	var v1034 int32
	_ = v1034
	var v1036 int32
	_ = v1036
	var v1039 int32
	_ = v1039
	var v1040 int32
	_ = v1040
	var v1042 int32
	_ = v1042
	var v1043 int32
	_ = v1043
	var v1044 int32
	_ = v1044
	var v1045 int32
	_ = v1045
	var v1052 int32
	_ = v1052
	var v1055 int32
	_ = v1055
	var v1057 int32
	_ = v1057
	var v1059 int32
	_ = v1059
	var v1063 int32
	_ = v1063
	var v1064 int32
	_ = v1064
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
	var v1076 int32
	_ = v1076
	var v1082 int32
	_ = v1082
	var v1084 int32
	_ = v1084
	var v1087 int32
	_ = v1087
	var v1089 int32
	_ = v1089
	var v1095 int32
	_ = v1095
	var v1098 int32
	_ = v1098
	var v1100 int32
	_ = v1100
	var v1102 int32
	_ = v1102
	var v1105 int32
	_ = v1105
	var v1106 int32
	_ = v1106
	var v1108 int32
	_ = v1108
	var v1109 int32
	_ = v1109
	var v1110 int32
	_ = v1110
	var v1111 int32
	_ = v1111
	var v1118 int32
	_ = v1118
	var v1121 int32
	_ = v1121
	var v1123 int32
	_ = v1123
	var v1125 int32
	_ = v1125
	var v1129 int32
	_ = v1129
	var v1130 int32
	_ = v1130
	var v1132 int32
	_ = v1132
	var v1135 int32
	_ = v1135
	var v1136 int32
	_ = v1136
	var v1137 int32
	_ = v1137
	var v1138 int32
	_ = v1138
	var v1139 int32
	_ = v1139
	var v1142 int32
	_ = v1142
	var v1148 int32
	_ = v1148
	var v1150 int32
	_ = v1150
	var v1153 int32
	_ = v1153
	var v1155 int32
	_ = v1155
	var v1165 int32
	_ = v1165
	var v1167 int32
	_ = v1167
	var v1172 int32
	_ = v1172
	var v1175 int32
	_ = v1175
	var v1179 int32
	_ = v1179
	var v1184 int32
	_ = v1184
	var v1185 int32
	_ = v1185
	var v1187 int32
	_ = v1187
	var v1192 int32
	_ = v1192
	var v1199 int32
	_ = v1199
	var v1200 int32
	_ = v1200
	var v1202 int32
	_ = v1202
	var v1204 int32
	_ = v1204
	var v1206 int32
	_ = v1206
	var v1207 int32
	_ = v1207
	var v1208 int32
	_ = v1208
	var v1209 int32
	_ = v1209
	var v1238 int32
	_ = v1238
	var v1241 int64
	_ = v1241
	var v1244 int32
	_ = v1244
	var v1248 int32
	_ = v1248
	var v1249 int32
	_ = v1249
	var v1250 int32
	_ = v1250
	var v1253 int32
	_ = v1253
	var v1255 int32
	_ = v1255
	var v1257 int32
	_ = v1257
	var v1265 int32
	_ = v1265
	var v1266 int32
	_ = v1266
	var v1268 int32
	_ = v1268
	var v1272 int32
	_ = v1272
	var v1275 int32
	_ = v1275
	var v1306 int32
	_ = v1306
	var v1307 int32
	_ = v1307
	var v1310 int32
	_ = v1310
	var v1312 int32
	_ = v1312
	var v1316 int32
	_ = v1316
	var v1317 int32
	_ = v1317
	var v1319 int32
	_ = v1319
	var v1354 int32
	_ = v1354
	var v1357 int32
	_ = v1357
	var v1391 int32
	_ = v1391
	var v1392 int32
	_ = v1392
	var v1423 int32
	_ = v1423
	var v1425 int32
	_ = v1425
	var v1427 int32
	_ = v1427
	var v1429 int32
	_ = v1429
	var v1430 int32
	_ = v1430
	var v1431 int32
	_ = v1431
	var v1438 int32
	_ = v1438
	var v1461 int32
	_ = v1461
	var v1464 int64
	_ = v1464
	var v1467 int32
	_ = v1467
	var v1469 int32
	_ = v1469
	var v1472 int32
	_ = v1472
	var v1473 int32
	_ = v1473
	var v1474 int32
	_ = v1474
	var v1478 int32
	_ = v1478
	var v1482 int32
	_ = v1482
	var v1485 int32
	_ = v1485
	var v1492 int32
	_ = v1492
	var v1493 int32
	_ = v1493
	var v1501 int32
	_ = v1501
	var v1514 int32
	_ = v1514
	var v1516 int32
	_ = v1516
	var v1519 int32
	_ = v1519
	var v1522 int32
	_ = v1522
	var v1526 int32
	_ = v1526
	var v1528 int32
	_ = v1528
	var v1530 int32
	_ = v1530
	var v1532 int32
	_ = v1532
	var v1544 int32
	_ = v1544
	var v1546 int32
	_ = v1546
	var v1548 int32
	_ = v1548
	var v1552 int32
	_ = v1552
	var v1559 int32
	_ = v1559
	var v1560 int32
	_ = v1560
	var v1581 int32
	_ = v1581
	var v1588 int32
	_ = v1588
	var v1589 int32
	_ = v1589
	var v1596 int32
	_ = v1596
	var v1613 int32
	_ = v1613
	var v1614 int32
	_ = v1614
	var v1616 int32
	_ = v1616
	var v1619 int32
	_ = v1619
	var v1620 int32
	_ = v1620
	var v1623 int32
	_ = v1623
	var v1632 int32
	_ = v1632
	var v1633 int32
	_ = v1633
	var v1662 int32
	_ = v1662
	var v1664 int32
	_ = v1664
	var v1665 int32
	_ = v1665
	var v1667 int32
	_ = v1667
	var v1668 int32
	_ = v1668
	var v1670 int32
	_ = v1670
	var v1674 int32
	_ = v1674
	var v1675 int32
	_ = v1675
	var v1676 int32
	_ = v1676
	var v1680 int32
	_ = v1680
	var v1682 int32
	_ = v1682
	var v1684 int32
	_ = v1684
	var v1685 int32
	_ = v1685
	var v1688 int32
	_ = v1688
	var v1695 int32
	_ = v1695
	var v1724 int64
	_ = v1724
	var v1730 int32
	_ = v1730
	var v1732 int32
	_ = v1732
	var v1736 int64
	_ = v1736
	var v1738 int64
	_ = v1738
	var v1747 int64
	_ = v1747
	var v1752 int32
	_ = v1752
	var v1754 int32
	_ = v1754
	var v1758 int64
	_ = v1758
	var v1760 int64
	_ = v1760
	var v1769 int64
	_ = v1769
	var v1775 int32
	_ = v1775
	var v1778 int32
	_ = v1778
	var v1782 int32
	_ = v1782
	var v1787 int32
	_ = v1787
	var v1788 int64
	_ = v1788
	var v1789 int64
	_ = v1789
	var v1791 int32
	_ = v1791
	var v1797 int64
	_ = v1797
	var v1802 int32
	_ = v1802
	var v1803 int32
	_ = v1803
	var v1806 int32
	_ = v1806
	var v1807 int32
	_ = v1807
	var v1808 int32
	_ = v1808
	var v1809 int32
	_ = v1809
	var v1812 int32
	_ = v1812
	var v1813 int32
	_ = v1813
	var v1814 int32
	_ = v1814
	var v1819 int32
	_ = v1819
	var v1821 int32
	_ = v1821
	var v1822 int32
	_ = v1822
	var v1827 int32
	_ = v1827
	var v1828 int32
	_ = v1828
	var v1831 int32
	_ = v1831
	var v1832 int32
	_ = v1832
	var v1834 int32
	_ = v1834
	var v1839 int32
	_ = v1839
	var v1843 int32
	_ = v1843
	var v1844 int32
	_ = v1844
	var v1847 int32
	_ = v1847
	var v1848 int64
	_ = v1848
	var v1854 int32
	_ = v1854
	var v1856 int32
	_ = v1856
	var v1860 int32
	_ = v1860
	var v1868 int32
	_ = v1868
	var v1870 int32
	_ = v1870
	var v1871 int32
	_ = v1871
	var v1873 int32
	_ = v1873
	var v1874 int32
	_ = v1874
	var v1875 int32
	_ = v1875
	var v1877 int32
	_ = v1877
	var v1878 int32
	_ = v1878
	var v1883 int32
	_ = v1883
	var v1886 int32
	_ = v1886
	var v1889 int32
	_ = v1889
	var v1893 int32
	_ = v1893
	var v1895 int32
	_ = v1895
	var v1896 int32
	_ = v1896
	var v1900 int32
	_ = v1900
	var v1901 int32
	_ = v1901
	var v1910 int32
	_ = v1910
	var v1941 int32
	_ = v1941
	var v1943 int64
	_ = v1943
	var v1945 int64
	_ = v1945
	var v1947 int32
	_ = v1947
	var v1954 int32
	_ = v1954
	var v1986 int32
	_ = v1986
	var v1987 int32
	_ = v1987
	var v2021 int32
	_ = v2021
	var v2023 int32
	_ = v2023
	var v2027 int32
	_ = v2027
	var v2031 int32
	_ = v2031
	var v2034 int32
	_ = v2034
	var v2038 int32
	_ = v2038
	var v2043 int32
	_ = v2043
	var v2048 int32
	_ = v2048
	var v2052 int32
	_ = v2052
	var v2056 int32
	_ = v2056
	var v2061 int32
	_ = v2061
	var v2065 int32
	_ = v2065
	var v2069 int32
	_ = v2069
	var v2074 int32
	_ = v2074
	var v2078 int32
	_ = v2078
	var v2081 int32
	_ = v2081
	var v2085 int32
	_ = v2085
	var v2090 int32
	_ = v2090
	var v2091 int32
	_ = v2091
	var v2093 int32
	_ = v2093
	var v2096 int32
	_ = v2096
	var v2099 int32
	_ = v2099
	var v2101 int32
	_ = v2101
	var v2108 int32
	_ = v2108
	var v2110 int32
	_ = v2110
	var v2114 int32
	_ = v2114
	var v2115 int32
	_ = v2115
	var v2116 int32
	_ = v2116
	var v2120 int32
	_ = v2120
	var v2123 int32
	_ = v2123
	var v2132 int32
	_ = v2132
	var v2157 int32
	_ = v2157
	var v2158 int64
	_ = v2158
	var v2160 int64
	_ = v2160
	var v2167 int32
	_ = v2167
	var v2168 int32
	_ = v2168
	var v2200 int32
	_ = v2200
	var v2204 int32
	_ = v2204
	var v2238 int32
	_ = v2238
	var v2239 int32
	_ = v2239
	var v2241 int32
	_ = v2241
	var v2243 int32
	_ = v2243
	var v2244 int32
	_ = v2244
	var v2273 int32
	_ = v2273
	var v2274 int64
	_ = v2274
	var v2276 int64
	_ = v2276
	var v2278 int64
	_ = v2278
	var v2280 int64
	_ = v2280
	var v2282 int64
	_ = v2282
	var v2284 int64
	_ = v2284
	var v2286 int32
	_ = v2286
	var v2288 int32
	_ = v2288
	var v2290 int32
	_ = v2290
	var v2297 int32
	_ = v2297
	var v2298 int32
	_ = v2298
	var v2360 int32
	_ = v2360
	var v2362 int32
	_ = v2362
	var v2365 int32
	_ = v2365
	var v2367 int32
	_ = v2367
	var v2371 int32
	_ = v2371
	var v2378 int32
	_ = v2378
	var v2383 int32
	_ = v2383
	var v2385 int32
	_ = v2385
	var v2387 int32
	_ = v2387
	var v2389 int32
	_ = v2389
	var v2395 int32
	_ = v2395
	var v2398 int32
	_ = v2398
	var v2402 int32
	_ = v2402
	var v2407 int32
	_ = v2407
	var v2409 int32
	_ = v2409
	var v2413 int32
	_ = v2413
	var v2414 int32
	_ = v2414
	var v2416 int32
	_ = v2416
	var v2417 int32
	_ = v2417
	var v2418 int32
	_ = v2418
	var v2419 int32
	_ = v2419
	var v2426 int32
	_ = v2426
	var v2428 int32
	_ = v2428
	var v2430 int32
	_ = v2430
	var v2435 int32
	_ = v2435
	var v2436 int32
	_ = v2436
	var v2438 int32
	_ = v2438
	var v2441 int32
	_ = v2441
	var v2442 int32
	_ = v2442
	var v2443 int32
	_ = v2443
	var v2444 int32
	_ = v2444
	var v2445 int32
	_ = v2445
	var v2446 int32
	_ = v2446
	var v2447 int32
	_ = v2447
	var v2453 int32
	_ = v2453
	var v2454 int32
	_ = v2454
	var v2461 int32
	_ = v2461
	var v2463 int32
	_ = v2463
	var v2468 int32
	_ = v2468
	var v2469 int32
	_ = v2469
	var v2471 int32
	_ = v2471
	var v2474 int32
	_ = v2474
	var v2476 int32
	_ = v2476
	var v2482 int64
	_ = v2482
	var v2485 int64
	_ = v2485
	var v2491 int32
	_ = v2491
	var v2493 int32
	_ = v2493
	var v2494 int32
	_ = v2494
	var v2496 int32
	_ = v2496
	var v2497 int32
	_ = v2497
	var v2501 int32
	_ = v2501
	var v2502 int32
	_ = v2502
	var v2507 int32
	_ = v2507
	var v2509 int32
	_ = v2509
	var v2510 int32
	_ = v2510
	var v2539 int32
	_ = v2539
	var v2540 int32
	_ = v2540
	var v2542 int32
	_ = v2542
	var v2543 int32
	_ = v2543
	var v2574 int32
	_ = v2574
	var v2576 int32
	_ = v2576
	var v2577 int32
	_ = v2577
	var v2581 int64
	_ = v2581
	var v2582 int32
	_ = v2582
	var v2589 int64
	_ = v2589
	var v2591 int32
	_ = v2591
	var v2592 int64
	_ = v2592
	var v2593 int64
	_ = v2593
	var v2595 int32
	_ = v2595
	var v2597 int64
	_ = v2597
	var v2600 int32
	_ = v2600
	var v2604 int32
	_ = v2604
	var v2605 int32
	_ = v2605
	var v2606 int32
	_ = v2606
	var v2609 int32
	_ = v2609
	var v2613 int32
	_ = v2613
	var v2615 int32
	_ = v2615
	var v2616 int32
	_ = v2616
	var v2617 int32
	_ = v2617
	var v2622 int32
	_ = v2622
	var v2624 int32
	_ = v2624
	var v2625 int32
	_ = v2625
	var v2631 int32
	_ = v2631
	var v2633 int32
	_ = v2633
	var v2637 int64
	_ = v2637
	var v2640 int32
	_ = v2640
	var v2642 int32
	_ = v2642
	var v2653 int32
	_ = v2653
	var v2656 int32
	_ = v2656
	var v2660 int32
	_ = v2660
	var v2665 int32
	_ = v2665
	var v2669 int32
	_ = v2669
	var v2671 int32
	_ = v2671
	var v2674 int32
	_ = v2674
	var v2675 int32
	_ = v2675
	var v2676 int32
	_ = v2676
	var v2678 int32
	_ = v2678
	var v2683 int32
	_ = v2683
	var v2685 int32
	_ = v2685
	var v2687 int32
	_ = v2687
	var v2688 int32
	_ = v2688
	var v2689 int32
	_ = v2689
	var v2696 int32
	_ = v2696
	var v2719 int32
	_ = v2719
	var v2722 int32
	_ = v2722
	var v2725 int32
	_ = v2725
	var v2728 int32
	_ = v2728
	var v2730 int32
	_ = v2730
	var v2733 int32
	_ = v2733
	var v2734 int32
	_ = v2734
	var v2735 int32
	_ = v2735
	var v2739 int32
	_ = v2739
	var v2743 int32
	_ = v2743
	var v2746 int32
	_ = v2746
	var v2747 int32
	_ = v2747
	var v2753 int32
	_ = v2753
	var v2758 int32
	_ = v2758
	var v2775 int32
	_ = v2775
	var v2777 int32
	_ = v2777
	var v2780 int32
	_ = v2780
	var v2781 int32
	_ = v2781
	var v2785 int32
	_ = v2785
	var v2791 int32
	_ = v2791
	var v2795 int32
	_ = v2795
	var v2799 int32
	_ = v2799
	var v2805 int32
	_ = v2805
	var v2807 int32
	_ = v2807
	var v2809 int32
	_ = v2809
	var v2813 int32
	_ = v2813
	var v2814 int32
	_ = v2814
	var v2820 int32
	_ = v2820
	var v2842 int32
	_ = v2842
	var v2843 int32
	_ = v2843
	var v2849 int32
	_ = v2849
	var v2850 int32
	_ = v2850
	var v2874 int32
	_ = v2874
	var v2875 int32
	_ = v2875
	var v2877 int32
	_ = v2877
	var v2880 int32
	_ = v2880
	var v2881 int32
	_ = v2881
	var v2884 int32
	_ = v2884
	var v2886 int32
	_ = v2886
	var v2893 int32
	_ = v2893
	var v2921 int64
	_ = v2921
	var v2924 int32
	_ = v2924
	var v2926 int32
	_ = v2926
	var v2960 int32
	_ = v2960
	var v2992 int32
	_ = v2992
	var v2993 int32
	_ = v2993
	var v3026 int32
	_ = v3026
	var v3028 int32
	_ = v3028
	var v3037 int32
	_ = v3037
	var v3038 int32
	_ = v3038
	var v3042 int32
	_ = v3042
	var v3060 int32
	_ = v3060
	var v3061 int32
	_ = v3061
	var v3062 int32
	_ = v3062
	var v3066 int32
	_ = v3066
	var v3072 int32
	_ = v3072
	var v3074 int32
	_ = v3074
	var v3075 int32
	_ = v3075
	var v3076 int32
	_ = v3076
	var v3077 int32
	_ = v3077
	var v3083 int32
	_ = v3083
	var v3084 int32
	_ = v3084
	var v3113 int32
	_ = v3113
	var v3115 int32
	_ = v3115
	var v3116 int32
	_ = v3116
	var v3117 int32
	_ = v3117
	var v3122 int32
	_ = v3122
	var v3127 int32
	_ = v3127
	var v3129 int32
	_ = v3129
	var v3131 int32
	_ = v3131
	var v3139 int32
	_ = v3139
	var v3141 int32
	_ = v3141
	var v3142 int32
	_ = v3142
	var v3144 int32
	_ = v3144
	var v3146 int32
	_ = v3146
	var v3149 int32
	_ = v3149
	var v3150 int32
	_ = v3150
	var v3152 int32
	_ = v3152
	var v3153 int32
	_ = v3153
	var v3154 int32
	_ = v3154
	var v3155 int32
	_ = v3155
	var v3157 int32
	_ = v3157
	var v3159 int32
	_ = v3159
	var v3160 int32
	_ = v3160
	var v3166 int32
	_ = v3166
	var v3170 int32
	_ = v3170
	var v3177 int32
	_ = v3177
	var v3205 int32
	_ = v3205
	var v3208 int32
	_ = v3208
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
	var v3218 int32
	_ = v3218
	var v3220 int32
	_ = v3220
	var v3222 int32
	_ = v3222
	var v3223 int32
	_ = v3223
	var v3229 int32
	_ = v3229
	var v3232 int32
	_ = v3232
	var v3237 int32
	_ = v3237
	var v3238 int32
	_ = v3238
	var v3239 int32
	_ = v3239
	var v3243 int32
	_ = v3243
	var v3269 int32
	_ = v3269
	var v3273 int32
	_ = v3273
	var v3274 int32
	_ = v3274
	var v3276 int32
	_ = v3276
	var v3278 int32
	_ = v3278
	var v3307 int32
	_ = v3307
	var v3313 int32
	_ = v3313
	var v3314 int32
	_ = v3314
	var v3315 int32
	_ = v3315
	var v3346 int32
	_ = v3346
	var v3353 int32
	_ = v3353
	var v3354 int32
	_ = v3354
	var v3358 int32
	_ = v3358
	var v3363 int32
	_ = v3363
	var v3394 int32
	_ = v3394
	var v3398 int32
	_ = v3398
	var v3399 int32
	_ = v3399
	var v3405 int32
	_ = v3405
	var v3410 int32
	_ = v3410
	var v3413 int32
	_ = v3413
	var v3419 int32
	_ = v3419
	var v3483 int32
	_ = v3483
	var v3485 int32
	_ = v3485
	var v3487 int32
	_ = v3487
	var v3495 int32
	_ = v3495
	var v3496 int32
	_ = v3496
	var v3518 int32
	_ = v3518
	var v3521 int32
	_ = v3521
	var v3523 int32
	_ = v3523
	var v3533 int32
	_ = v3533
	var v3536 int32
	_ = v3536
	var v3540 int32
	_ = v3540
	var v3545 int32
	_ = v3545
	var v3549 int32
	_ = v3549
	var v3553 int32
	_ = v3553
	var v3558 int32
	_ = v3558
	var v3562 int32
	_ = v3562
	var v3566 int32
	_ = v3566
	var v3571 int32
	_ = v3571
	var v3573 int32
	_ = v3573
	var v3575 int32
	_ = v3575
	var v3579 int32
	_ = v3579
	var v3580 int32
	_ = v3580
	var v3582 int32
	_ = v3582
	var v3583 int32
	_ = v3583
	var v3584 int32
	_ = v3584
	var v3588 int32
	_ = v3588
	var v3595 int32
	_ = v3595
	var v3596 int64
	_ = v3596
	var v3600 int32
	_ = v3600
	var v3603 int32
	_ = v3603
	var v3607 int32
	_ = v3607
	var v3608 int32
	_ = v3608
	var v3609 int32
	_ = v3609
	var v3610 int32
	_ = v3610
	var v3612 int32
	_ = v3612
	var v3615 int32
	_ = v3615
	var v3616 int32
	_ = v3616
	var v3624 int32
	_ = v3624
	var v3628 int32
	_ = v3628
	var v3630 int32
	_ = v3630
	var v3632 int32
	_ = v3632
	var v3660 int32
	_ = v3660
	var v3662 int32
	_ = v3662
	var v3663 int32
	_ = v3663
	var v3665 int32
	_ = v3665
	var v3696 int32
	_ = v3696
	var v3697 int32
	_ = v3697
	var v3701 int32
	_ = v3701
	var v3703 int32
	_ = v3703
	var v3706 int32
	_ = v3706
	var v3708 int32
	_ = v3708
	var v3710 int32
	_ = v3710
	var v3711 int32
	_ = v3711
	var v3713 int32
	_ = v3713
	var v3741 int32
	_ = v3741
	var v3744 int32
	_ = v3744
	var v3807 int32
	_ = v3807
	var v3812 int32
	_ = v3812
	var v3813 int32
	_ = v3813
	var v3814 int32
	_ = v3814
	var v3817 int32
	_ = v3817
	var v3846 int32
	_ = v3846
	var v3851 int32
	_ = v3851
	var v3853 int32
	_ = v3853
	var v3884 int32
	_ = v3884
	var v3885 int32
	_ = v3885
	var v3888 int32
	_ = v3888
	var v3917 int32
	_ = v3917
	var v3922 int32
	_ = v3922
	var v3924 int32
	_ = v3924
	var v3988 int32
	_ = v3988
	var v3990 int32
	_ = v3990
	var v4019 int32
	_ = v4019
	var v4022 int32
	_ = v4022
	var v4053 int32
	_ = v4053
	var v4055 int32
	_ = v4055
	var v4059 int32
	_ = v4059
	var v4061 int32
	_ = v4061
	var v4062 int32
	_ = v4062
	var v4064 int32
	_ = v4064
	var v4068 int32
	_ = v4068
	var v4069 int32
	_ = v4069
	var v4071 int32
	_ = v4071
	var v4072 int32
	_ = v4072
	var v4079 int32
	_ = v4079
	var v4086 int32
	_ = v4086
	var v4090 int32
	_ = v4090
	var v4092 int32
	_ = v4092
	var v4094 int32
	_ = v4094
	var v4096 int32
	_ = v4096
	var v4100 int32
	_ = v4100
	var v4103 int32
	_ = v4103
	var v4115 int32
	_ = v4115
	var v4119 int32
	_ = v4119
	var v4121 int32
	_ = v4121
	var v4123 int32
	_ = v4123
	var v4132 int32
	_ = v4132
	var v4134 int32
	_ = v4134
	var v4137 int32
	_ = v4137
	var v4139 int32
	_ = v4139
	var v4141 int32
	_ = v4141
	var v4144 int32
	_ = v4144
	var v4146 int32
	_ = v4146
	var v4150 int32
	_ = v4150
	var v4151 int32
	_ = v4151
	var v4153 int32
	_ = v4153
	var v4157 int32
	_ = v4157
	var v4161 int32
	_ = v4161
	var v4165 int32
	_ = v4165
	var v4168 int32
	_ = v4168
	var v4171 int32
	_ = v4171
	var v4173 int32
	_ = v4173
	var v4180 int32
	_ = v4180
	var v4184 int32
	_ = v4184
	var v4186 int32
	_ = v4186
	var v4189 int32
	_ = v4189
	var v4191 int32
	_ = v4191
	var v4204 int32
	_ = v4204
	var v4206 int32
	_ = v4206
	var v4207 int32
	_ = v4207
	var v4210 int32
	_ = v4210
	var v4229 int32
	_ = v4229
	var v4233 int32
	_ = v4233
	var v4235 int32
	_ = v4235
	var v4239 int32
	_ = v4239
	var v4243 int32
	_ = v4243
	var v4246 int32
	_ = v4246
	var v4248 int32
	_ = v4248
	var v4249 int32
	_ = v4249
	var v4252 int32
	_ = v4252
	var v4256 int32
	_ = v4256
	var v4258 int32
	_ = v4258
	var v4259 int32
	_ = v4259
	var v4264 int32
	_ = v4264
	var v4265 int32
	_ = v4265
	var v4271 int32
	_ = v4271
	var v4280 int32
	_ = v4280
	var v4282 int32
	_ = v4282
	var v4283 int32
	_ = v4283
	var v4293 int32
	_ = v4293
	var v4294 int32
	_ = v4294
	var v4297 int32
	_ = v4297
	var v4299 int32
	_ = v4299
	var v4301 int32
	_ = v4301
	var v4307 int64
	_ = v4307
	var v4323 int32
	_ = v4323
	var v4325 int32
	_ = v4325
	var v4335 int32
	_ = v4335
	var v4338 int32
	_ = v4338
	var v4342 int32
	_ = v4342
	var v4347 int32
	_ = v4347
	var v4351 int32
	_ = v4351
	var v4354 int32
	_ = v4354
	var v4358 int32
	_ = v4358
	var v4363 int32
	_ = v4363
	v30 = m.G0
	v32 = v30 - int32(16)
	m.G0 = v32
	v35 = *(*int32)(unsafe.Add(mBase, _c_F_PrepareTransaction[0]))
	v36 = *(*int64)(unsafe.Add(mBase, uint32(v35)))
	if base.I32_wrap_i64(v36) == int32(0) {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	F_AssignTransactionId(m, v35)
	mBase = m.M
	v41 = m.ExcPending
	if v41 != 0 {
		goto L4
	} else {
		goto L5
	}
L2:
	;
	v43 = v36
	goto L3
L3:
	;
	v44 = int32(10)
	goto L8
L4:
	;
	return
L5:
	;
	v42 = *(*int64)(unsafe.Add(mBase, uint32(v35)))
	v43 = v42
	goto L3
L6:
	;
	if v84 != 0 {
		goto L19
	} else {
		goto L20
	}
L7:
	;
	goto L6
L8:
	;
	v51 = *(*int32)(unsafe.Add(mBase, _c_F_PrepareTransaction[1]))
	v54 = *(*int32)(unsafe.Add(mBase, uint32(v51<<(uint(int32(2))%32))+uint32(_c_F_PrepareTransaction[2])))
	goto L11
L9:
	;
	v67 = int32(0)
	goto L16
L11:
	;
	goto L12
L12:
	;
	if int32(0)|base.B2i32(v54 == int32(15)) != 0 {
		goto L9
	} else {
		goto L14
	}
L14:
	;
	if v54 <= v44 {
		v84 = int32(1)
		goto L7
	} else {
		goto L15
	}
L15:
	;
	goto L9
L16:
	;
	v71 = *(*int32)(unsafe.Add(mBase, _c_F_PrepareTransaction[3]))
	if v71 != int32(2) {
		v84 = v67
		goto L7
	} else {
		goto L17
	}
L17:
	;
	v75 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_PrepareTransaction[4])))
	if v75&int32(1) != 0 {
		v84 = v67
		goto L7
	} else {
		goto L18
	}
L18:
	;
	v81 = *(*int32)(unsafe.Add(mBase, _c_F_PrepareTransaction[5]))
	v84 = int32(0) | base.B2i32(v81 <= v44)
	goto L7
L19:
	;
	v88 = *(*int32)(unsafe.Add(mBase, _c_F_PrepareTransaction[0]))
	F_ShowTransactionStateRec(m, int32(_a_F_PrepareTransaction_0), v88)
	mBase = m.M
	v90 = m.ExcPending
	if v90 != 0 {
		goto L4
	} else {
		goto L22
	}
L20:
	;
	goto L21
L21:
	;
	v91 = *(*int32)(unsafe.Add(mBase, uint32(v35)+20))
	if v91 == int32(2) {
		goto L23
	} else {
		goto L24
	}
L22:
	;
	goto L21
L23:
	;
	goto L32
L24:
	;
	v96 = F_errstart(m, int32(19), int32(0))
	mBase = m.M
	v97 = m.ExcPending
	if v97 != 0 {
		goto L4
	} else {
		goto L25
	}
L25:
	;
	if v96 == int32(0) {
		goto L23
	} else {
		goto L26
	}
L26:
	;
	v100 = *(*int32)(unsafe.Add(mBase, uint32(v35)+20))
	if base.Ui32(v100) <= base.Ui32(int32(5)) {
		goto L27
	} else {
		goto L28
	}
L27:
	;
	v105 = *(*int32)(unsafe.Add(mBase, uint32(v100<<(uint(int32(2))%32))+uint32(_c_F_PrepareTransaction[6])))
	v107 = v105
	goto L29
L28:
	;
	v107 = int32(_a_F_PrepareTransaction_1)
	goto L29
L29:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v32))) = v107
	F_errmsg_internal(m, int32(_a_F_PrepareTransaction_2), v32)
	mBase = m.M
	v111 = m.ExcPending
	if v111 != 0 {
		goto L4
	} else {
		goto L30
	}
L30:
	;
	F_errfinish(m, int32(_a_F_PrepareTransaction_3), int32(2575), int32(_a_F_PrepareTransaction_0))
	mBase = m.M
	v116 = m.ExcPending
	if v116 != 0 {
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
	v148 = m.ExcPending
	if v148 != 0 {
		goto L4
	} else {
		goto L34
	}
L33:
	;
	v153 = *(*int32)(unsafe.Add(mBase, _c_F_PrepareTransaction[7]))
	if v153 != 0 {
		goto L37
	} else {
		goto L38
	}
L34:
	;
	v150 = F_PreCommit_Portals(m, int32(1))
	mBase = m.M
	v151 = m.ExcPending
	if v151 != 0 {
		goto L4
	} else {
		goto L35
	}
L35:
	;
	if v150 != 0 {
		goto L32
	} else {
		goto L36
	}
L36:
	;
	goto L33
L37:
	;
	v155 = v153
	goto L40
L38:
	;
	goto L39
L39:
	;
	F_AfterTriggerEndXact(m)
	mBase = m.M
	v219 = m.ExcPending
	if v219 != 0 {
		goto L4
	} else {
		goto L44
	}
L40:
	;
	v183 = *(*int32)(unsafe.Add(mBase, uint32(v155)))
	v185 = *(*int32)(unsafe.Add(mBase, uint32(v155)+8))
	v186 = *(*int32)(unsafe.Add(mBase, uint32(v155)+4))
	m.T0[v186].(func(*base.Module, int32, int32))(m, int32(7), v185)
	mBase = m.M
	v188 = m.ExcPending
	if v188 != 0 {
		goto L4
	} else {
		goto L42
	}
L41:
	;
	goto L39
L42:
	;
	if v183 != 0 {
		v155 = v183
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
	v221 = m.ExcPending
	if v221 != 0 {
		goto L4
	} else {
		goto L45
	}
L45:
	;
	F_smgrDoPendingSyncs(m, int32(1), int32(0))
	mBase = m.M
	v225 = m.ExcPending
	if v225 != 0 {
		goto L4
	} else {
		goto L46
	}
L46:
	;
	F_AtEOXact_LargeObject(m, int32(1))
	mBase = m.M
	v228 = m.ExcPending
	if v228 != 0 {
		goto L4
	} else {
		goto L47
	}
L47:
	;
	F_PreCommit_CheckForSerializationFailure(m)
	mBase = m.M
	v230 = m.ExcPending
	if v230 != 0 {
		goto L4
	} else {
		goto L48
	}
L48:
	;
	v232 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_PrepareTransaction[8])))
	if v232&int32(1) == int32(0) {
		goto L50
	} else {
		goto L51
	}
L49:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v4351 = m.ExcPending
	if v4351 != 0 {
		goto L4
	} else {
		goto L654
	}
L50:
	;
	v238 = *(*int32)(unsafe.Add(mBase, _c_F_PrepareTransaction[9]))
	if v238 != 0 {
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
	v4335 = m.ExcPending
	if v4335 != 0 {
		goto L4
	} else {
		goto L650
	}
L53:
	;
	v239 = int32(_a_F_PrepareTransaction_4)
	v241 = *(*int32)(unsafe.Add(mBase, _c_F_PrepareTransaction[10]))
	*(*int32)(unsafe.Add(mBase, _c_F_PrepareTransaction[10])) = v241 + int32(1)
	*(*int32)(unsafe.Add(mBase, uint32(v35)+20)) = int32(5)
	v248 = *(*int32)(unsafe.Add(mBase, _c_F_PrepareTransaction[11]))
	if int32(0) < v248 {
		goto L54
	} else {
		goto L55
	}
L54:
	;
	F_disable_timeout(m, int32(8))
	mBase = m.M
	v253 = m.ExcPending
	if v253 != 0 {
		goto L4
	} else {
		goto L57
	}
L55:
	;
	goto L56
L56:
	;
	v257 = m.G0
	v258 = int32(16)
	v259 = v257 - v258
	m.G0 = v259
	F_gettimeofday(m, v259)
	mBase = m.M
	v262 = *(*int64)(unsafe.Add(mBase, uint32(v259)))
	v263 = int64(*(*int32)(unsafe.Add(mBase, uint32(v259)+8)))
	m.G0 = v259 + v258
	goto L58
L57:
	;
	goto L56
L58:
	;
	v273 = *(*int32)(unsafe.Add(mBase, _c_F_PrepareTransaction[12]))
	v275 = *(*int32)(unsafe.Add(mBase, _c_F_PrepareTransaction[13]))
	v277 = *(*int32)(unsafe.Add(mBase, _c_F_PrepareTransaction[14]))
	v279 = m.G0
	v281 = v279 - int32(48)
	m.G0 = v281
	v283 = F_strlen(m, v273)
	mBase = m.M
	if base.Ui32(v283) < base.Ui32(int32(200)) {
		goto L62
	} else {
		goto L63
	}
L59:
	;
	v514 = int32(0)
	*(*int32)(unsafe.Add(mBase, _c_F_PrepareTransaction[12])) = v514
	v516 = m.G0
	v518 = v516 - int32(96)
	m.G0 = v518
	v521 = *(*int32)(unsafe.Add(mBase, _c_F_PrepareTransaction[15]))
	v522 = *(*int32)(unsafe.Add(mBase, uint32(v521)))
	v523 = *(*int32)(unsafe.Add(mBase, uint32(v426)+4))
	v524 = *(*int64)(unsafe.Add(mBase, uint32(v426)+32))
	*(*int32)(unsafe.Add(mBase, uint32(v518)+8)) = v514
	*(*int32)(unsafe.Add(mBase, uint32(v518)+4)) = v514
	v531 = F_palloc0(m, int32(12))
	mBase = m.M
	v532 = m.ExcPending
	if v532 != 0 {
		goto L4
	} else {
		goto L106
	}
L60:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v492 = m.ExcPending
	if v492 != 0 {
		goto L4
	} else {
		goto L101
	}
L61:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v472 = m.ExcPending
	if v472 != 0 {
		goto L4
	} else {
		goto L96
	}
L62:
	;
	v287 = *(*int32)(unsafe.Add(mBase, _c_F_PrepareTransaction[16]))
	if v287 == int32(0) {
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
	v456 = m.ExcPending
	if v456 != 0 {
		goto L4
	} else {
		goto L92
	}
L65:
	;
	v291 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_PrepareTransaction[17])))
	if v291 == int32(0) {
		goto L66
	} else {
		goto L67
	}
L66:
	;
	F_before_shmem_exit(m, int32(413), int64(0))
	mBase = m.M
	v297 = m.ExcPending
	if v297 != 0 {
		goto L4
	} else {
		goto L69
	}
L67:
	;
	goto L68
L68:
	;
	v302 = *(*int32)(unsafe.Add(mBase, _c_F_PrepareTransaction[18]))
	v306 = F_LWLockAcquire(m, v302+int32(2304), int32(0))
	mBase = m.M
	v307 = m.ExcPending
	if v307 != 0 {
		goto L4
	} else {
		goto L70
	}
L69:
	;
	v299 = int32(1)
	*(*uint8)(unsafe.Add(mBase, _c_F_PrepareTransaction[17])) = uint8(v299)
	goto L68
L70:
	;
	v309 = *(*int32)(unsafe.Add(mBase, _c_F_PrepareTransaction[19]))
	v310 = *(*int32)(unsafe.Add(mBase, uint32(v309)+4))
	if v310 <= int32(0) {
		goto L71
	} else {
		goto L72
	}
L71:
	;
	v426 = *(*int32)(unsafe.Add(mBase, uint32(v309)))
	if v426 == int32(0) {
		goto L60
	} else {
		goto L90
	}
L72:
	;
	v316 = int32(0)
	goto L73
L73:
	;
	v347 = *(*int32)(unsafe.Add(mBase, uint32(v309+int32(8)+v316<<(uint(int32(2))%32))))
	v349 = v347 + int32(51)
	v352 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v349))))
	v355 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v273))))
	if base.B2i32(v352 == int32(0))|base.B2i32(v352 != v355) != 0 {
		v373 = v352
		v374 = v355
		goto L76
	} else {
		goto L77
	}
L74:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v382 = m.ExcPending
	if v382 != 0 {
		goto L4
	} else {
		goto L86
	}
L75:
	;
	if v373-v374 != 0 {
		goto L82
	} else {
		goto L83
	}
L76:
	;
	goto L75
L77:
	;
	v358 = v349
	v359 = v273
	goto L78
L78:
	;
	v362 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v359)+1)))
	v363 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v358)+1)))
	if v363 == int32(0) {
		v373 = v363
		v374 = v362
		goto L76
	} else {
		goto L80
	}
L79:
	;
	v373 = v363
	v374 = v362
	goto L76
L80:
	;
	v366 = int32(1)
	if v363 == v362 {
		v358 = v358 + v366
		v359 = v359 + v366
		goto L78
	} else {
		goto L81
	}
L81:
	;
	goto L79
L82:
	;
	v377 = v316 + int32(1)
	if v310 != v377 {
		v316 = v377
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
	v385 = m.ExcPending
	if v385 != 0 {
		goto L4
	} else {
		goto L87
	}
L87:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v281)+16)) = v273
	F_errmsg(m, int32(_a_F_PrepareTransaction_6), v281+int32(16))
	mBase = m.M
	v391 = m.ExcPending
	if v391 != 0 {
		goto L4
	} else {
		goto L88
	}
L88:
	;
	F_errfinish(m, int32(_a_F_PrepareTransaction_7), int32(402), int32(_a_F_PrepareTransaction_8))
	mBase = m.M
	v396 = m.ExcPending
	if v396 != 0 {
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
	v429 = *(*int32)(unsafe.Add(mBase, uint32(v426)))
	*(*int32)(unsafe.Add(mBase, uint32(v309))) = v429
	F_MarkAsPreparingGuts(m, v426, v43, v273, v263+v262*int64(1000000)-int64(946684800000000), v275, v277)
	mBase = m.M
	v432 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v426)+49)) = uint8(v432)
	v435 = *(*int32)(unsafe.Add(mBase, _c_F_PrepareTransaction[19]))
	v436 = *(*int32)(unsafe.Add(mBase, uint32(v435)+4))
	*(*int32)(unsafe.Add(mBase, uint32(v435)+4)) = v436 + int32(1)
	*(*int32)(unsafe.Add(mBase, uint32(v435+v436<<(uint(int32(2))%32))+8)) = v426
	v445 = *(*int32)(unsafe.Add(mBase, _c_F_PrepareTransaction[18]))
	F_LWLockRelease(m, v445+int32(2304))
	mBase = m.M
	v449 = m.ExcPending
	if v449 != 0 {
		goto L4
	} else {
		goto L91
	}
L91:
	;
	m.G0 = v281 + int32(48)
	goto L59
L92:
	;
	F_errcode(m, int32(50856066))
	mBase = m.M
	v459 = m.ExcPending
	if v459 != 0 {
		goto L4
	} else {
		goto L93
	}
L93:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v281))) = v273
	F_errmsg(m, int32(_a_F_PrepareTransaction_9), v281)
	mBase = m.M
	v463 = m.ExcPending
	if v463 != 0 {
		goto L4
	} else {
		goto L94
	}
L94:
	;
	F_errfinish(m, int32(_a_F_PrepareTransaction_7), int32(375), int32(_a_F_PrepareTransaction_8))
	mBase = m.M
	v468 = m.ExcPending
	if v468 != 0 {
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
	v475 = m.ExcPending
	if v475 != 0 {
		goto L4
	} else {
		goto L97
	}
L97:
	;
	F_errmsg(m, int32(_a_F_PrepareTransaction_10), int32(0))
	mBase = m.M
	v479 = m.ExcPending
	if v479 != 0 {
		goto L4
	} else {
		goto L98
	}
L98:
	;
	F_errhint(m, int32(_a_F_PrepareTransaction_11), int32(0))
	mBase = m.M
	v483 = m.ExcPending
	if v483 != 0 {
		goto L4
	} else {
		goto L99
	}
L99:
	;
	F_errfinish(m, int32(_a_F_PrepareTransaction_7), int32(382), int32(_a_F_PrepareTransaction_8))
	mBase = m.M
	v488 = m.ExcPending
	if v488 != 0 {
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
	v495 = m.ExcPending
	if v495 != 0 {
		goto L4
	} else {
		goto L102
	}
L102:
	;
	F_errmsg(m, int32(_a_F_PrepareTransaction_13), int32(0))
	mBase = m.M
	v499 = m.ExcPending
	if v499 != 0 {
		goto L4
	} else {
		goto L103
	}
L103:
	;
	v501 = *(*int32)(unsafe.Add(mBase, _c_F_PrepareTransaction[16]))
	*(*int32)(unsafe.Add(mBase, uint32(v281)+32)) = v501
	F_errhint(m, int32(_a_F_PrepareTransaction_14), v281+int32(32))
	mBase = m.M
	v507 = m.ExcPending
	if v507 != 0 {
		goto L4
	} else {
		goto L104
	}
L104:
	;
	F_errfinish(m, int32(_a_F_PrepareTransaction_7), int32(412), int32(_a_F_PrepareTransaction_8))
	mBase = m.M
	v512 = m.ExcPending
	if v512 != 0 {
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
	*(*int32)(unsafe.Add(mBase, _c_F_PrepareTransaction[20])) = v531
	*(*int64)(unsafe.Add(mBase, uint32(v531)+4)) = int64(0)
	v537 = int32(512)
	*(*int32)(unsafe.Add(mBase, _c_F_PrepareTransaction[21])) = v537
	v540 = F_palloc(m, v537)
	mBase = m.M
	v541 = m.ExcPending
	if v541 != 0 {
		goto L4
	} else {
		goto L107
	}
L107:
	;
	v543 = *(*int32)(unsafe.Add(mBase, _c_F_PrepareTransaction[20]))
	*(*int32)(unsafe.Add(mBase, uint32(v543))) = v540
	*(*int32)(unsafe.Add(mBase, _c_F_PrepareTransaction[22])) = int32(1)
	*(*int32)(unsafe.Add(mBase, _c_F_PrepareTransaction[23])) = v543
	*(*int32)(unsafe.Add(mBase, _c_F_PrepareTransaction[24])) = int32(0)
	*(*uint32)(unsafe.Add(mBase, uint32(v518)+32)) = uint32(v524)
	*(*int64)(unsafe.Add(mBase, uint32(v518)+24)) = int64(1475953972)
	v559 = *(*int32)(unsafe.Add(mBase, uint32(v522+v523*int32(768))+20))
	*(*int32)(unsafe.Add(mBase, uint32(v518)+36)) = v559
	v561 = *(*int64)(unsafe.Add(mBase, uint32(v426)+8))
	*(*int64)(unsafe.Add(mBase, uint32(v518)+40)) = v561
	v563 = *(*int32)(unsafe.Add(mBase, uint32(v426)+40))
	*(*int32)(unsafe.Add(mBase, uint32(v518)+48)) = v563
	v568 = *(*int32)(unsafe.Add(mBase, _c_F_PrepareTransaction[0]))
	v569 = *(*int32)(unsafe.Add(mBase, uint32(v568)+52))
	if v569 != 0 {
		goto L109
	} else {
		goto L110
	}
L108:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v518)+52)) = v574
	v579 = F_smgrGetPendingDeletes(m, int32(1), v518+int32(16))
	mBase = m.M
	v580 = m.ExcPending
	if v580 != 0 {
		goto L4
	} else {
		goto L112
	}
L109:
	;
	v570 = *(*int32)(unsafe.Add(mBase, uint32(v568)+48))
	v572 = v570
	goto L111
L110:
	;
	v572 = int32(0)
	goto L111
L111:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v518+int32(20)))) = v572
	v574 = *(*int32)(unsafe.Add(mBase, uint32(v568)+52))
	goto L108
L112:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v518)+56)) = v579
	v585 = F_smgrGetPendingDeletes(m, int32(0), v518+int32(12))
	mBase = m.M
	v586 = m.ExcPending
	if v586 != 0 {
		goto L4
	} else {
		goto L113
	}
L113:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v518)+60)) = v585
	v591 = F_pgstat_get_transactional_drops(m, int32(1), v518+int32(4))
	mBase = m.M
	v592 = m.ExcPending
	if v592 != 0 {
		goto L4
	} else {
		goto L114
	}
L114:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v518)+64)) = v591
	v597 = F_pgstat_get_transactional_drops(m, int32(0), v518+int32(8))
	mBase = m.M
	v598 = m.ExcPending
	if v598 != 0 {
		goto L4
	} else {
		goto L115
	}
L115:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v518)+68)) = v597
	v602 = F_xactGetCommittedInvalidationMessages(m, v518, v518+int32(76))
	mBase = m.M
	v603 = m.ExcPending
	if v603 != 0 {
		goto L4
	} else {
		goto L116
	}
L116:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v518)+72)) = v602
	v606 = v426 + int32(51)
	v607 = F_strlen(m, v606)
	mBase = m.M
	v608 = int64(0)
	*(*int64)(unsafe.Add(mBase, uint32(v518)+80)) = v608
	*(*int64)(unsafe.Add(mBase, uint32(v518)+88)) = v608
	v613 = v607 + int32(1)
	*(*uint16)(unsafe.Add(mBase, uint32(v518)+78)) = uint16(v613)
	v616 = *(*int32)(unsafe.Add(mBase, _c_F_PrepareTransaction[21]))
	if base.Ui32(int32(72)) <= base.Ui32(v616) {
		goto L118
	} else {
		goto L119
	}
L117:
	;
	v653 = *(*int32)(unsafe.Add(mBase, uint32(v651)+4))
	v657 = int32(72)
	base.MemoryCopy(m, v650+v653, v518+int32(24), v657)
	v659 = *(*int32)(unsafe.Add(mBase, uint32(v651)+4))
	v661 = v659 + v657
	*(*int32)(unsafe.Add(mBase, uint32(v651)+4)) = v661
	v665 = v652 - v657
	*(*int32)(unsafe.Add(mBase, _c_F_PrepareTransaction[21])) = v665
	v667 = int32(_a_F_PrepareTransaction_15)
	v669 = *(*int32)(unsafe.Add(mBase, _c_F_PrepareTransaction[24]))
	v671 = v669 + v657
	*(*int32)(unsafe.Add(mBase, _c_F_PrepareTransaction[24])) = v671
	v673 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v518)+78)))
	v677 = (v673 + int32(7)) & int32(_a_F_PrepareTransaction_16)
	if base.Ui32(v677) <= base.Ui32(v665) {
		goto L124
	} else {
		goto L125
	}
L118:
	;
	v620 = *(*int32)(unsafe.Add(mBase, _c_F_PrepareTransaction[23]))
	v621 = *(*int32)(unsafe.Add(mBase, uint32(v620)))
	v650 = v621
	v651 = v620
	v652 = v616
	goto L117
L119:
	;
	goto L120
L120:
	;
	v623 = F_palloc0(m, int32(12))
	mBase = m.M
	v624 = m.ExcPending
	if v624 != 0 {
		goto L4
	} else {
		goto L121
	}
L121:
	;
	v625 = int32(_a_F_PrepareTransaction_17)
	v626 = *(*int32)(unsafe.Add(mBase, _c_F_PrepareTransaction[23]))
	*(*int32)(unsafe.Add(mBase, uint32(v626)+8)) = v623
	*(*int32)(unsafe.Add(mBase, _c_F_PrepareTransaction[23])) = v623
	*(*int64)(unsafe.Add(mBase, uint32(v623)+4)) = int64(0)
	v633 = int32(512)
	*(*int32)(unsafe.Add(mBase, _c_F_PrepareTransaction[21])) = v633
	v635 = int32(_a_F_PrepareTransaction_18)
	v637 = *(*int32)(unsafe.Add(mBase, _c_F_PrepareTransaction[22]))
	*(*int32)(unsafe.Add(mBase, _c_F_PrepareTransaction[22])) = v637 + int32(1)
	v642 = F_palloc(m, v633)
	mBase = m.M
	v643 = m.ExcPending
	if v643 != 0 {
		goto L4
	} else {
		goto L122
	}
L122:
	;
	v645 = *(*int32)(unsafe.Add(mBase, _c_F_PrepareTransaction[23]))
	*(*int32)(unsafe.Add(mBase, uint32(v645))) = v642
	v648 = *(*int32)(unsafe.Add(mBase, _c_F_PrepareTransaction[21]))
	v650 = v642
	v651 = v645
	v652 = v648
	goto L117
L123:
	;
	if v673 != 0 {
		goto L132
	} else {
		goto L133
	}
L124:
	;
	v679 = *(*int32)(unsafe.Add(mBase, uint32(v651)))
	v713 = v665
	v714 = v651
	v715 = v671
	v716 = v661
	v717 = v679
	goto L123
L125:
	;
	goto L126
L126:
	;
	v681 = F_palloc0(m, int32(12))
	mBase = m.M
	v682 = m.ExcPending
	if v682 != 0 {
		goto L4
	} else {
		goto L127
	}
L127:
	;
	v683 = int32(_a_F_PrepareTransaction_17)
	v684 = *(*int32)(unsafe.Add(mBase, _c_F_PrepareTransaction[23]))
	*(*int32)(unsafe.Add(mBase, uint32(v684)+8)) = v681
	*(*int32)(unsafe.Add(mBase, _c_F_PrepareTransaction[23])) = v681
	*(*int64)(unsafe.Add(mBase, uint32(v681)+4)) = int64(0)
	v691 = int32(512)
	if base.Ui32(v677) <= base.Ui32(v691) {
		goto L128
	} else {
		goto L129
	}
L128:
	;
	v694 = v691
	goto L130
L129:
	;
	v694 = v677
	goto L130
L130:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_PrepareTransaction[21])) = v694
	v696 = int32(_a_F_PrepareTransaction_18)
	v698 = *(*int32)(unsafe.Add(mBase, _c_F_PrepareTransaction[22]))
	*(*int32)(unsafe.Add(mBase, _c_F_PrepareTransaction[22])) = v698 + int32(1)
	v702 = F_palloc(m, v694)
	mBase = m.M
	v703 = m.ExcPending
	if v703 != 0 {
		goto L4
	} else {
		goto L131
	}
L131:
	;
	v705 = *(*int32)(unsafe.Add(mBase, _c_F_PrepareTransaction[23]))
	*(*int32)(unsafe.Add(mBase, uint32(v705))) = v702
	v707 = *(*int32)(unsafe.Add(mBase, uint32(v705)+4))
	v709 = *(*int32)(unsafe.Add(mBase, _c_F_PrepareTransaction[24]))
	v711 = *(*int32)(unsafe.Add(mBase, _c_F_PrepareTransaction[21]))
	v713 = v711
	v714 = v705
	v715 = v709
	v716 = v707
	v717 = v702
	goto L123
L132:
	;
	base.MemoryCopy(m, v716+v717, v606, v673)
	goto L134
L133:
	;
	goto L134
L134:
	;
	v720 = *(*int32)(unsafe.Add(mBase, uint32(v714)+4))
	v721 = v720 + v677
	*(*int32)(unsafe.Add(mBase, uint32(v714)+4)) = v721
	v724 = v677 + v715
	*(*int32)(unsafe.Add(mBase, _c_F_PrepareTransaction[24])) = v724
	v727 = v713 - v677
	*(*int32)(unsafe.Add(mBase, _c_F_PrepareTransaction[21])) = v727
	v729 = *(*int32)(unsafe.Add(mBase, uint32(v518)+52))
	if v729 <= int32(0) {
		v817 = v727
		v819 = v724
		goto L135
	} else {
		goto L136
	}
L135:
	;
	v821 = *(*int32)(unsafe.Add(mBase, uint32(v518)+56))
	if int32(0) < v821 {
		goto L157
	} else {
		goto L158
	}
L136:
	;
	v732 = *(*int32)(unsafe.Add(mBase, uint32(v518)+20))
	v734 = v729 << (uint(int32(2)) % 32)
	v738 = (v734 + int32(7)) & int32(-8)
	if base.Ui32(v738) <= base.Ui32(v727) {
		goto L138
	} else {
		goto L139
	}
L137:
	;
	if v734 != 0 {
		goto L146
	} else {
		goto L147
	}
L138:
	;
	v740 = *(*int32)(unsafe.Add(mBase, uint32(v714)))
	v772 = v714
	v773 = v727
	v774 = v721
	v775 = v740
	goto L137
L139:
	;
	goto L140
L140:
	;
	v742 = F_palloc0(m, int32(12))
	mBase = m.M
	v743 = m.ExcPending
	if v743 != 0 {
		goto L4
	} else {
		goto L141
	}
L141:
	;
	v744 = int32(_a_F_PrepareTransaction_17)
	v745 = *(*int32)(unsafe.Add(mBase, _c_F_PrepareTransaction[23]))
	*(*int32)(unsafe.Add(mBase, uint32(v745)+8)) = v742
	*(*int32)(unsafe.Add(mBase, _c_F_PrepareTransaction[23])) = v742
	*(*int64)(unsafe.Add(mBase, uint32(v742)+4)) = int64(0)
	v752 = int32(512)
	if base.Ui32(v738) <= base.Ui32(v752) {
		goto L142
	} else {
		goto L143
	}
L142:
	;
	v755 = v752
	goto L144
L143:
	;
	v755 = v738
	goto L144
L144:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_PrepareTransaction[21])) = v755
	v757 = int32(_a_F_PrepareTransaction_18)
	v759 = *(*int32)(unsafe.Add(mBase, _c_F_PrepareTransaction[22]))
	*(*int32)(unsafe.Add(mBase, _c_F_PrepareTransaction[22])) = v759 + int32(1)
	v763 = F_palloc(m, v755)
	mBase = m.M
	v764 = m.ExcPending
	if v764 != 0 {
		goto L4
	} else {
		goto L145
	}
L145:
	;
	v766 = *(*int32)(unsafe.Add(mBase, _c_F_PrepareTransaction[23]))
	*(*int32)(unsafe.Add(mBase, uint32(v766))) = v763
	v768 = *(*int32)(unsafe.Add(mBase, uint32(v766)+4))
	v770 = *(*int32)(unsafe.Add(mBase, _c_F_PrepareTransaction[21]))
	v772 = v766
	v773 = v770
	v774 = v768
	v775 = v763
	goto L137
L146:
	;
	base.MemoryCopy(m, v774+v775, v732, v734)
	goto L148
L147:
	;
	goto L148
L148:
	;
	v778 = *(*int32)(unsafe.Add(mBase, uint32(v772)+4))
	*(*int32)(unsafe.Add(mBase, uint32(v772)+4)) = v778 + v738
	v782 = v773 - v738
	*(*int32)(unsafe.Add(mBase, _c_F_PrepareTransaction[21])) = v782
	v784 = int32(_a_F_PrepareTransaction_15)
	v786 = *(*int32)(unsafe.Add(mBase, _c_F_PrepareTransaction[24]))
	v787 = v786 + v738
	*(*int32)(unsafe.Add(mBase, _c_F_PrepareTransaction[24])) = v787
	v790 = *(*int32)(unsafe.Add(mBase, _c_F_PrepareTransaction[15]))
	v791 = *(*int32)(unsafe.Add(mBase, uint32(v790)))
	v792 = *(*int32)(unsafe.Add(mBase, uint32(v426)+4))
	v795 = v791 + v792*int32(768)
	v796 = *(*int32)(unsafe.Add(mBase, uint32(v518)+20))
	v797 = *(*int32)(unsafe.Add(mBase, uint32(v518)+52))
	if int32(65) <= v797 {
		goto L150
	} else {
		goto L151
	}
L149:
	;
	v807 = v805 << (uint(int32(2)) % 32)
	if v807 != 0 {
		goto L154
	} else {
		goto L155
	}
L150:
	;
	v800 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(v795)+57)) = uint8(v800)
	v805 = int32(64)
	goto L149
L151:
	;
	goto L152
L152:
	;
	if v797 <= int32(0) {
		v817 = v782
		v819 = v787
		goto L135
	} else {
		goto L153
	}
L153:
	;
	v805 = v797
	goto L149
L154:
	;
	base.MemoryCopy(m, v795+int32(60), v796, v807)
	goto L156
L155:
	;
	goto L156
L156:
	;
	*(*uint8)(unsafe.Add(mBase, uint32(v795)+56)) = uint8(v805)
	v817 = v782
	v819 = v787
	goto L135
L157:
	;
	v824 = *(*int32)(unsafe.Add(mBase, uint32(v518)+16))
	v826 = v821 * int32(12)
	v830 = (v826 + int32(7)) & int32(-8)
	if base.Ui32(v830) <= base.Ui32(v817) {
		goto L161
	} else {
		goto L162
	}
L158:
	;
	goto L159
L159:
	;
	v892 = *(*int32)(unsafe.Add(mBase, uint32(v518)+60))
	if int32(0) < v892 {
		goto L173
	} else {
		goto L174
	}
L160:
	;
	if v826 != 0 {
		goto L169
	} else {
		goto L170
	}
L161:
	;
	v833 = *(*int32)(unsafe.Add(mBase, _c_F_PrepareTransaction[23]))
	v834 = *(*int32)(unsafe.Add(mBase, uint32(v833)))
	v866 = v833
	v867 = v817
	v868 = v834
	v869 = v819
	goto L160
L162:
	;
	goto L163
L163:
	;
	v836 = F_palloc0(m, int32(12))
	mBase = m.M
	v837 = m.ExcPending
	if v837 != 0 {
		goto L4
	} else {
		goto L164
	}
L164:
	;
	v838 = int32(_a_F_PrepareTransaction_17)
	v839 = *(*int32)(unsafe.Add(mBase, _c_F_PrepareTransaction[23]))
	*(*int32)(unsafe.Add(mBase, uint32(v839)+8)) = v836
	*(*int32)(unsafe.Add(mBase, _c_F_PrepareTransaction[23])) = v836
	*(*int64)(unsafe.Add(mBase, uint32(v836)+4)) = int64(0)
	v846 = int32(512)
	if base.Ui32(v830) <= base.Ui32(v846) {
		goto L165
	} else {
		goto L166
	}
L165:
	;
	v849 = v846
	goto L167
L166:
	;
	v849 = v830
	goto L167
L167:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_PrepareTransaction[21])) = v849
	v851 = int32(_a_F_PrepareTransaction_18)
	v853 = *(*int32)(unsafe.Add(mBase, _c_F_PrepareTransaction[22]))
	*(*int32)(unsafe.Add(mBase, _c_F_PrepareTransaction[22])) = v853 + int32(1)
	v857 = F_palloc(m, v849)
	mBase = m.M
	v858 = m.ExcPending
	if v858 != 0 {
		goto L4
	} else {
		goto L168
	}
L168:
	;
	v860 = *(*int32)(unsafe.Add(mBase, _c_F_PrepareTransaction[23]))
	*(*int32)(unsafe.Add(mBase, uint32(v860))) = v857
	v863 = *(*int32)(unsafe.Add(mBase, _c_F_PrepareTransaction[24]))
	v865 = *(*int32)(unsafe.Add(mBase, _c_F_PrepareTransaction[21]))
	v866 = v860
	v867 = v865
	v868 = v857
	v869 = v863
	goto L160
L169:
	;
	v870 = *(*int32)(unsafe.Add(mBase, uint32(v866)+4))
	base.MemoryCopy(m, v868+v870, v824, v826)
	goto L171
L170:
	;
	goto L171
L171:
	;
	v873 = *(*int32)(unsafe.Add(mBase, uint32(v866)+4))
	*(*int32)(unsafe.Add(mBase, uint32(v866)+4)) = v873 + v830
	*(*int32)(unsafe.Add(mBase, _c_F_PrepareTransaction[24])) = v830 + v869
	*(*int32)(unsafe.Add(mBase, _c_F_PrepareTransaction[21])) = v867 - v830
	v882 = *(*int32)(unsafe.Add(mBase, uint32(v518)+16))
	F_pfree(m, v882)
	mBase = m.M
	v884 = m.ExcPending
	if v884 != 0 {
		goto L4
	} else {
		goto L172
	}
L172:
	;
	goto L159
L173:
	;
	v895 = *(*int32)(unsafe.Add(mBase, uint32(v518)+12))
	v897 = v892 * int32(12)
	v901 = (v897 + int32(7)) & int32(-8)
	v903 = *(*int32)(unsafe.Add(mBase, _c_F_PrepareTransaction[21]))
	if base.Ui32(v901) <= base.Ui32(v903) {
		goto L177
	} else {
		goto L178
	}
L174:
	;
	goto L175
L175:
	;
	v963 = *(*int32)(unsafe.Add(mBase, uint32(v518)+64))
	if int32(0) < v963 {
		goto L189
	} else {
		goto L190
	}
L176:
	;
	if v897 != 0 {
		goto L185
	} else {
		goto L186
	}
L177:
	;
	v906 = *(*int32)(unsafe.Add(mBase, _c_F_PrepareTransaction[23]))
	v907 = *(*int32)(unsafe.Add(mBase, uint32(v906)))
	v937 = v906
	v938 = v903
	v939 = v907
	goto L176
L178:
	;
	goto L179
L179:
	;
	v909 = F_palloc0(m, int32(12))
	mBase = m.M
	v910 = m.ExcPending
	if v910 != 0 {
		goto L4
	} else {
		goto L180
	}
L180:
	;
	v911 = int32(_a_F_PrepareTransaction_17)
	v912 = *(*int32)(unsafe.Add(mBase, _c_F_PrepareTransaction[23]))
	*(*int32)(unsafe.Add(mBase, uint32(v912)+8)) = v909
	*(*int32)(unsafe.Add(mBase, _c_F_PrepareTransaction[23])) = v909
	*(*int64)(unsafe.Add(mBase, uint32(v909)+4)) = int64(0)
	v919 = int32(512)
	if base.Ui32(v901) <= base.Ui32(v919) {
		goto L181
	} else {
		goto L182
	}
L181:
	;
	v922 = v919
	goto L183
L182:
	;
	v922 = v901
	goto L183
L183:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_PrepareTransaction[21])) = v922
	v924 = int32(_a_F_PrepareTransaction_18)
	v926 = *(*int32)(unsafe.Add(mBase, _c_F_PrepareTransaction[22]))
	*(*int32)(unsafe.Add(mBase, _c_F_PrepareTransaction[22])) = v926 + int32(1)
	v930 = F_palloc(m, v922)
	mBase = m.M
	v931 = m.ExcPending
	if v931 != 0 {
		goto L4
	} else {
		goto L184
	}
L184:
	;
	v933 = *(*int32)(unsafe.Add(mBase, _c_F_PrepareTransaction[23]))
	*(*int32)(unsafe.Add(mBase, uint32(v933))) = v930
	v936 = *(*int32)(unsafe.Add(mBase, _c_F_PrepareTransaction[21]))
	v937 = v933
	v938 = v936
	v939 = v930
	goto L176
L185:
	;
	v940 = *(*int32)(unsafe.Add(mBase, uint32(v937)+4))
	base.MemoryCopy(m, v939+v940, v895, v897)
	goto L187
L186:
	;
	goto L187
L187:
	;
	v943 = *(*int32)(unsafe.Add(mBase, uint32(v937)+4))
	*(*int32)(unsafe.Add(mBase, uint32(v937)+4)) = v943 + v901
	*(*int32)(unsafe.Add(mBase, _c_F_PrepareTransaction[21])) = v938 - v901
	v949 = int32(_a_F_PrepareTransaction_15)
	v951 = *(*int32)(unsafe.Add(mBase, _c_F_PrepareTransaction[24]))
	*(*int32)(unsafe.Add(mBase, _c_F_PrepareTransaction[24])) = v951 + v901
	v954 = *(*int32)(unsafe.Add(mBase, uint32(v518)+12))
	F_pfree(m, v954)
	mBase = m.M
	v956 = m.ExcPending
	if v956 != 0 {
		goto L4
	} else {
		goto L188
	}
L188:
	;
	goto L175
L189:
	;
	v966 = *(*int32)(unsafe.Add(mBase, uint32(v518)+4))
	v968 = v963 << (uint(int32(4)) % 32)
	v970 = *(*int32)(unsafe.Add(mBase, _c_F_PrepareTransaction[21]))
	if base.Ui32(v968) <= base.Ui32(v970) {
		goto L193
	} else {
		goto L194
	}
L190:
	;
	goto L191
L191:
	;
	v1029 = *(*int32)(unsafe.Add(mBase, uint32(v518)+68))
	if int32(0) < v1029 {
		goto L205
	} else {
		goto L206
	}
L192:
	;
	if v968 != 0 {
		goto L201
	} else {
		goto L202
	}
L193:
	;
	v973 = *(*int32)(unsafe.Add(mBase, _c_F_PrepareTransaction[23]))
	v974 = *(*int32)(unsafe.Add(mBase, uint32(v973)))
	v1004 = v973
	v1005 = v974
	v1006 = v970
	goto L192
L194:
	;
	goto L195
L195:
	;
	v976 = F_palloc0(m, int32(12))
	mBase = m.M
	v977 = m.ExcPending
	if v977 != 0 {
		goto L4
	} else {
		goto L196
	}
L196:
	;
	v978 = int32(_a_F_PrepareTransaction_17)
	v979 = *(*int32)(unsafe.Add(mBase, _c_F_PrepareTransaction[23]))
	*(*int32)(unsafe.Add(mBase, uint32(v979)+8)) = v976
	*(*int32)(unsafe.Add(mBase, _c_F_PrepareTransaction[23])) = v976
	*(*int64)(unsafe.Add(mBase, uint32(v976)+4)) = int64(0)
	v986 = int32(512)
	if base.Ui32(v968) <= base.Ui32(v986) {
		goto L197
	} else {
		goto L198
	}
L197:
	;
	v989 = v986
	goto L199
L198:
	;
	v989 = v968
	goto L199
L199:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_PrepareTransaction[21])) = v989
	v991 = int32(_a_F_PrepareTransaction_18)
	v993 = *(*int32)(unsafe.Add(mBase, _c_F_PrepareTransaction[22]))
	*(*int32)(unsafe.Add(mBase, _c_F_PrepareTransaction[22])) = v993 + int32(1)
	v997 = F_palloc(m, v989)
	mBase = m.M
	v998 = m.ExcPending
	if v998 != 0 {
		goto L4
	} else {
		goto L200
	}
L200:
	;
	v1000 = *(*int32)(unsafe.Add(mBase, _c_F_PrepareTransaction[23]))
	*(*int32)(unsafe.Add(mBase, uint32(v1000))) = v997
	v1003 = *(*int32)(unsafe.Add(mBase, _c_F_PrepareTransaction[21]))
	v1004 = v1000
	v1005 = v997
	v1006 = v1003
	goto L192
L201:
	;
	v1007 = *(*int32)(unsafe.Add(mBase, uint32(v1004)+4))
	base.MemoryCopy(m, v1005+v1007, v966, v968)
	goto L203
L202:
	;
	goto L203
L203:
	;
	v1010 = *(*int32)(unsafe.Add(mBase, uint32(v1004)+4))
	*(*int32)(unsafe.Add(mBase, uint32(v1004)+4)) = v1010 + v968
	*(*int32)(unsafe.Add(mBase, _c_F_PrepareTransaction[21])) = v1006 - v968
	v1016 = int32(_a_F_PrepareTransaction_15)
	v1018 = *(*int32)(unsafe.Add(mBase, _c_F_PrepareTransaction[24]))
	*(*int32)(unsafe.Add(mBase, _c_F_PrepareTransaction[24])) = v1018 + v968
	v1021 = *(*int32)(unsafe.Add(mBase, uint32(v518)+4))
	F_pfree(m, v1021)
	mBase = m.M
	v1023 = m.ExcPending
	if v1023 != 0 {
		goto L4
	} else {
		goto L204
	}
L204:
	;
	goto L191
L205:
	;
	v1032 = *(*int32)(unsafe.Add(mBase, uint32(v518)+8))
	v1034 = v1029 << (uint(int32(4)) % 32)
	v1036 = *(*int32)(unsafe.Add(mBase, _c_F_PrepareTransaction[21]))
	if base.Ui32(v1034) <= base.Ui32(v1036) {
		goto L209
	} else {
		goto L210
	}
L206:
	;
	goto L207
L207:
	;
	v1095 = *(*int32)(unsafe.Add(mBase, uint32(v518)+72))
	if int32(0) < v1095 {
		goto L221
	} else {
		goto L222
	}
L208:
	;
	if v1034 != 0 {
		goto L217
	} else {
		goto L218
	}
L209:
	;
	v1039 = *(*int32)(unsafe.Add(mBase, _c_F_PrepareTransaction[23]))
	v1040 = *(*int32)(unsafe.Add(mBase, uint32(v1039)))
	v1070 = v1039
	v1071 = v1040
	v1072 = v1036
	goto L208
L210:
	;
	goto L211
L211:
	;
	v1042 = F_palloc0(m, int32(12))
	mBase = m.M
	v1043 = m.ExcPending
	if v1043 != 0 {
		goto L4
	} else {
		goto L212
	}
L212:
	;
	v1044 = int32(_a_F_PrepareTransaction_17)
	v1045 = *(*int32)(unsafe.Add(mBase, _c_F_PrepareTransaction[23]))
	*(*int32)(unsafe.Add(mBase, uint32(v1045)+8)) = v1042
	*(*int32)(unsafe.Add(mBase, _c_F_PrepareTransaction[23])) = v1042
	*(*int64)(unsafe.Add(mBase, uint32(v1042)+4)) = int64(0)
	v1052 = int32(512)
	if base.Ui32(v1034) <= base.Ui32(v1052) {
		goto L213
	} else {
		goto L214
	}
L213:
	;
	v1055 = v1052
	goto L215
L214:
	;
	v1055 = v1034
	goto L215
L215:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_PrepareTransaction[21])) = v1055
	v1057 = int32(_a_F_PrepareTransaction_18)
	v1059 = *(*int32)(unsafe.Add(mBase, _c_F_PrepareTransaction[22]))
	*(*int32)(unsafe.Add(mBase, _c_F_PrepareTransaction[22])) = v1059 + int32(1)
	v1063 = F_palloc(m, v1055)
	mBase = m.M
	v1064 = m.ExcPending
	if v1064 != 0 {
		goto L4
	} else {
		goto L216
	}
L216:
	;
	v1066 = *(*int32)(unsafe.Add(mBase, _c_F_PrepareTransaction[23]))
	*(*int32)(unsafe.Add(mBase, uint32(v1066))) = v1063
	v1069 = *(*int32)(unsafe.Add(mBase, _c_F_PrepareTransaction[21]))
	v1070 = v1066
	v1071 = v1063
	v1072 = v1069
	goto L208
L217:
	;
	v1073 = *(*int32)(unsafe.Add(mBase, uint32(v1070)+4))
	base.MemoryCopy(m, v1071+v1073, v1032, v1034)
	goto L219
L218:
	;
	goto L219
L219:
	;
	v1076 = *(*int32)(unsafe.Add(mBase, uint32(v1070)+4))
	*(*int32)(unsafe.Add(mBase, uint32(v1070)+4)) = v1076 + v1034
	*(*int32)(unsafe.Add(mBase, _c_F_PrepareTransaction[21])) = v1072 - v1034
	v1082 = int32(_a_F_PrepareTransaction_15)
	v1084 = *(*int32)(unsafe.Add(mBase, _c_F_PrepareTransaction[24]))
	*(*int32)(unsafe.Add(mBase, _c_F_PrepareTransaction[24])) = v1084 + v1034
	v1087 = *(*int32)(unsafe.Add(mBase, uint32(v518)+8))
	F_pfree(m, v1087)
	mBase = m.M
	v1089 = m.ExcPending
	if v1089 != 0 {
		goto L4
	} else {
		goto L220
	}
L220:
	;
	goto L207
L221:
	;
	v1098 = *(*int32)(unsafe.Add(mBase, uint32(v518)))
	v1100 = v1095 << (uint(int32(4)) % 32)
	v1102 = *(*int32)(unsafe.Add(mBase, _c_F_PrepareTransaction[21]))
	if base.Ui32(v1100) <= base.Ui32(v1102) {
		goto L225
	} else {
		goto L226
	}
L222:
	;
	goto L223
L223:
	;
	m.G0 = v518 + int32(96)
	v1165 = *(*int32)(unsafe.Add(mBase, _c_F_PrepareTransaction[25]))
	v1167 = *(*int32)(unsafe.Add(mBase, _c_F_PrepareTransaction[26]))
	if v1165|v1167 != 0 {
		goto L237
	} else {
		goto L238
	}
L224:
	;
	if v1100 != 0 {
		goto L233
	} else {
		goto L234
	}
L225:
	;
	v1105 = *(*int32)(unsafe.Add(mBase, _c_F_PrepareTransaction[23]))
	v1106 = *(*int32)(unsafe.Add(mBase, uint32(v1105)))
	v1136 = v1105
	v1137 = v1106
	v1138 = v1102
	goto L224
L226:
	;
	goto L227
L227:
	;
	v1108 = F_palloc0(m, int32(12))
	mBase = m.M
	v1109 = m.ExcPending
	if v1109 != 0 {
		goto L4
	} else {
		goto L228
	}
L228:
	;
	v1110 = int32(_a_F_PrepareTransaction_17)
	v1111 = *(*int32)(unsafe.Add(mBase, _c_F_PrepareTransaction[23]))
	*(*int32)(unsafe.Add(mBase, uint32(v1111)+8)) = v1108
	*(*int32)(unsafe.Add(mBase, _c_F_PrepareTransaction[23])) = v1108
	*(*int64)(unsafe.Add(mBase, uint32(v1108)+4)) = int64(0)
	v1118 = int32(512)
	if base.Ui32(v1100) <= base.Ui32(v1118) {
		goto L229
	} else {
		goto L230
	}
L229:
	;
	v1121 = v1118
	goto L231
L230:
	;
	v1121 = v1100
	goto L231
L231:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_PrepareTransaction[21])) = v1121
	v1123 = int32(_a_F_PrepareTransaction_18)
	v1125 = *(*int32)(unsafe.Add(mBase, _c_F_PrepareTransaction[22]))
	*(*int32)(unsafe.Add(mBase, _c_F_PrepareTransaction[22])) = v1125 + int32(1)
	v1129 = F_palloc(m, v1121)
	mBase = m.M
	v1130 = m.ExcPending
	if v1130 != 0 {
		goto L4
	} else {
		goto L232
	}
L232:
	;
	v1132 = *(*int32)(unsafe.Add(mBase, _c_F_PrepareTransaction[23]))
	*(*int32)(unsafe.Add(mBase, uint32(v1132))) = v1129
	v1135 = *(*int32)(unsafe.Add(mBase, _c_F_PrepareTransaction[21]))
	v1136 = v1132
	v1137 = v1129
	v1138 = v1135
	goto L224
L233:
	;
	v1139 = *(*int32)(unsafe.Add(mBase, uint32(v1136)+4))
	base.MemoryCopy(m, v1137+v1139, v1098, v1100)
	goto L235
L234:
	;
	goto L235
L235:
	;
	v1142 = *(*int32)(unsafe.Add(mBase, uint32(v1136)+4))
	*(*int32)(unsafe.Add(mBase, uint32(v1136)+4)) = v1142 + v1100
	*(*int32)(unsafe.Add(mBase, _c_F_PrepareTransaction[21])) = v1138 - v1100
	v1148 = int32(_a_F_PrepareTransaction_15)
	v1150 = *(*int32)(unsafe.Add(mBase, _c_F_PrepareTransaction[24]))
	*(*int32)(unsafe.Add(mBase, _c_F_PrepareTransaction[24])) = v1150 + v1100
	v1153 = *(*int32)(unsafe.Add(mBase, uint32(v518)))
	F_pfree(m, v1153)
	mBase = m.M
	v1155 = m.ExcPending
	if v1155 != 0 {
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
	v1172 = m.ExcPending
	if v1172 != 0 {
		goto L4
	} else {
		goto L240
	}
L238:
	;
	goto L239
L239:
	;
	v1185 = m.G0
	v1187 = v1185 - int32(80)
	m.G0 = v1187
	*(*int64)(unsafe.Add(mBase, uint32(v1187)+40)) = int64(85899345936)
	v1192 = *(*int32)(unsafe.Add(mBase, _c_F_PrepareTransaction[27]))
	*(*int32)(unsafe.Add(mBase, uint32(v1187)+68)) = v1192
	v1199 = F_hash_create(m, int32(_a_F_PrepareTransaction_19), int64(256), v1187+int32(32), int32(1064))
	mBase = m.M
	v1200 = m.ExcPending
	if v1200 != 0 {
		goto L4
	} else {
		goto L244
	}
L240:
	;
	F_errcode(m, int32(1088))
	mBase = m.M
	v1175 = m.ExcPending
	if v1175 != 0 {
		goto L4
	} else {
		goto L241
	}
L241:
	;
	F_errmsg(m, int32(_a_F_PrepareTransaction_20), int32(0))
	mBase = m.M
	v1179 = m.ExcPending
	if v1179 != 0 {
		goto L4
	} else {
		goto L242
	}
L242:
	;
	F_errfinish(m, int32(_a_F_PrepareTransaction_21), int32(1166), int32(_a_F_PrepareTransaction_22))
	mBase = m.M
	v1184 = m.ExcPending
	if v1184 != 0 {
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
	v1202 = v1187 + int32(8)
	v1204 = *(*int32)(unsafe.Add(mBase, _c_F_PrepareTransaction[28]))
	F_hash_seq_init(m, v1202, v1204)
	mBase = m.M
	v1206 = m.ExcPending
	if v1206 != 0 {
		goto L4
	} else {
		goto L245
	}
L245:
	;
	v1207 = F_hash_seq_search(m, v1202)
	mBase = m.M
	v1208 = m.ExcPending
	if v1208 != 0 {
		goto L4
	} else {
		goto L248
	}
L246:
	;
	v2091 = m.G0
	v2093 = v2091 - int32(32)
	m.G0 = v2093
	v2096 = *(*int32)(unsafe.Add(mBase, _c_F_PrepareTransaction[29]))
	if v2096 != 0 {
		goto L374
	} else {
		goto L375
	}
L247:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v2078 = m.ExcPending
	if v2078 != 0 {
		goto L4
	} else {
		goto L370
	}
L248:
	;
	if v1207 != 0 {
		goto L249
	} else {
		goto L250
	}
L249:
	;
	v1209 = v1207
	goto L252
L250:
	;
	goto L251
L251:
	;
	F_hash_destroy(m, v1199)
	mBase = m.M
	v1423 = m.ExcPending
	if v1423 != 0 {
		goto L4
	} else {
		goto L286
	}
L252:
	;
	v1238 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1209)+14)))
	if v1238 == int32(6) {
		goto L254
	} else {
		goto L255
	}
L253:
	;
	goto L251
L254:
	;
	v1391 = F_hash_seq_search(m, v1187+int32(8))
	mBase = m.M
	v1392 = m.ExcPending
	if v1392 != 0 {
		goto L4
	} else {
		goto L284
	}
L255:
	;
	v1241 = *(*int64)(unsafe.Add(mBase, uint32(v1209)+32))
	if v1241 <= int64(0) {
		goto L254
	} else {
		goto L256
	}
L256:
	;
	v1244 = *(*int32)(unsafe.Add(mBase, uint32(v1209)+48))
	v1248 = F_hash_search(m, v1199, v1209, int32(1), v1187+int32(7))
	mBase = m.M
	v1249 = m.ExcPending
	if v1249 != 0 {
		goto L4
	} else {
		goto L257
	}
L257:
	;
	v1250 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1187)+7)))
	if v1250 == int32(0) {
		goto L258
	} else {
		goto L259
	}
L258:
	;
	v1253 = int32(0)
	*(*uint16)(unsafe.Add(mBase, uint32(v1248)+16)) = uint16(v1253)
	goto L260
L259:
	;
	goto L260
L260:
	;
	v1255 = *(*int32)(unsafe.Add(mBase, uint32(v1209)+40))
	v1257 = v1255 - int32(1)
	if v1257 < int32(0) {
		goto L261
	} else {
		goto L262
	}
L261:
	;
	v1354 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1248)+16)))
	if v1354 != int32(1) {
		goto L254
	} else {
		goto L282
	}
L262:
	;
	if v1255&int32(1) != 0 {
		goto L263
	} else {
		goto L264
	}
L263:
	;
	v1265 = *(*int32)(unsafe.Add(mBase, uint32(v1244+v1257<<(uint(int32(4))%32))))
	if v1265 != 0 {
		goto L267
	} else {
		goto L268
	}
L264:
	;
	v1272 = v1257
	goto L265
L265:
	;
	if v1257 == int32(0) {
		goto L261
	} else {
		goto L270
	}
L266:
	;
	v1272 = v1255 - int32(2)
	goto L265
L267:
	;
	v1266 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(v1248)+17)) = uint8(v1266)
	goto L266
L268:
	;
	goto L269
L269:
	;
	v1268 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(v1248)+16)) = uint8(v1268)
	goto L266
L270:
	;
	v1275 = v1272
	goto L271
L271:
	;
	v1306 = v1244 + v1275<<(uint(int32(4))%32)
	v1307 = *(*int32)(unsafe.Add(mBase, uint32(v1306)))
	if v1307 == int32(0) {
		goto L274
	} else {
		goto L275
	}
L272:
	;
	goto L261
L273:
	;
	v1316 = *(*int32)(unsafe.Add(mBase, uint32(v1306-int32(16))))
	if v1316 != 0 {
		goto L278
	} else {
		goto L279
	}
L274:
	;
	v1310 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(v1248)+16)) = uint8(v1310)
	goto L273
L275:
	;
	goto L276
L276:
	;
	v1312 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(v1248)+17)) = uint8(v1312)
	goto L273
L277:
	;
	if int32(1) < v1275 {
		v1275 = v1275 - int32(2)
		goto L271
	} else {
		goto L281
	}
L278:
	;
	v1317 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(v1248)+17)) = uint8(v1317)
	goto L277
L279:
	;
	goto L280
L280:
	;
	v1319 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(v1248)+16)) = uint8(v1319)
	goto L277
L281:
	;
	goto L272
L282:
	;
	v1357 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1248)+17)))
	if v1357 == int32(1) {
		goto L247
	} else {
		goto L283
	}
L283:
	;
	goto L254
L284:
	;
	if v1391 != 0 {
		v1209 = v1391
		goto L252
	} else {
		goto L285
	}
L285:
	;
	goto L253
L286:
	;
	v1425 = v1187 + int32(32)
	v1427 = *(*int32)(unsafe.Add(mBase, _c_F_PrepareTransaction[28]))
	F_hash_seq_init(m, v1425, v1427)
	mBase = m.M
	v1429 = m.ExcPending
	if v1429 != 0 {
		goto L4
	} else {
		goto L287
	}
L287:
	;
	v1430 = F_hash_seq_search(m, v1425)
	mBase = m.M
	v1431 = m.ExcPending
	if v1431 != 0 {
		goto L4
	} else {
		goto L291
	}
L288:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v2065 = m.ExcPending
	if v2065 != 0 {
		goto L4
	} else {
		goto L367
	}
L289:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v2052 = m.ExcPending
	if v2052 != 0 {
		goto L4
	} else {
		goto L364
	}
L290:
	;
	F_LWLockRelease(m, v1695)
	mBase = m.M
	v2021 = m.ExcPending
	if v2021 != 0 {
		goto L4
	} else {
		goto L357
	}
L291:
	;
	if v1430 != 0 {
		goto L292
	} else {
		goto L293
	}
L292:
	;
	v1438 = v1430
	goto L295
L293:
	;
	goto L294
L294:
	;
	m.G0 = v1187 + int32(80)
	goto L246
L295:
	;
	v1461 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1438)+14)))
	if v1461 == int32(6) {
		goto L297
	} else {
		goto L298
	}
L296:
	;
	goto L294
L297:
	;
	v1986 = F_hash_seq_search(m, v1187+int32(32))
	mBase = m.M
	v1987 = m.ExcPending
	if v1987 != 0 {
		goto L4
	} else {
		goto L355
	}
L298:
	;
	v1464 = *(*int64)(unsafe.Add(mBase, uint32(v1438)+32))
	if v1464 <= int64(0) {
		goto L297
	} else {
		goto L299
	}
L299:
	;
	v1467 = *(*int32)(unsafe.Add(mBase, uint32(v1438)+40))
	v1469 = v1467 - int32(1)
	if v1469 < int32(0) {
		goto L297
	} else {
		goto L300
	}
L300:
	;
	v1472 = *(*int32)(unsafe.Add(mBase, uint32(v1438)+48))
	v1473 = int32(3)
	v1474 = v1467 & v1473
	if base.Ui32(v1469) < base.Ui32(v1473) {
		goto L303
	} else {
		goto L304
	}
L301:
	;
	if v1633&int32(1) == int32(0) {
		goto L297
	} else {
		goto L313
	}
L302:
	;
	v1581 = v1552
	v1588 = v1559
	v1589 = v1560
	v1596 = int32(0)
	goto L310
L303:
	;
	v1478 = int32(0)
	v1552 = v1469
	v1559 = v1478
	v1560 = v1478
	goto L302
L304:
	;
	goto L305
L305:
	;
	v1482 = int32(0)
	v1485 = v1469
	v1492 = v1482
	v1493 = v1482
	v1501 = v1482
	goto L306
L306:
	;
	v1514 = int32(4)
	v1516 = v1472 + v1485<<(uint(v1514)%32)
	v1519 = *(*int32)(unsafe.Add(mBase, uint32(v1516-int32(48))))
	v1522 = *(*int32)(unsafe.Add(mBase, uint32(v1516-int32(32))))
	v1526 = *(*int32)(unsafe.Add(mBase, uint32(v1516-int32(16))))
	v1528 = *(*int32)(unsafe.Add(mBase, uint32(v1516)))
	v1530 = int32(0)
	v1532 = base.B2i32(v1519|v1522|v1526|v1528 != v1530) | v1493
	v1544 = base.B2i32(v1528 == v1530) | (base.B2i32(v1526 == v1530) | (base.B2i32(v1519 == v1530) | base.B2i32(v1522 == v1530))) | v1492
	v1546 = v1485 - v1514
	v1548 = v1501 + v1514
	if v1548 != v1467&int32(-4) {
		v1485 = v1546
		v1492 = v1544
		v1493 = v1532
		v1501 = v1548
		goto L306
	} else {
		goto L308
	}
L307:
	;
	if v1474 == int32(0) {
		v1632 = v1544
		v1633 = v1532
		goto L301
	} else {
		goto L309
	}
L308:
	;
	goto L307
L309:
	;
	v1552 = v1546
	v1559 = v1544
	v1560 = v1532
	goto L302
L310:
	;
	v1613 = *(*int32)(unsafe.Add(mBase, uint32(v1472+v1581<<(uint(int32(4))%32))))
	v1614 = int32(0)
	v1616 = base.B2i32(v1613 != v1614) | v1589
	v1619 = base.B2i32(v1613 == v1614) | v1588
	v1620 = int32(1)
	v1623 = v1596 + v1620
	if v1623 != v1474 {
		v1581 = v1581 - v1620
		v1588 = v1619
		v1589 = v1616
		v1596 = v1623
		goto L310
	} else {
		goto L312
	}
L311:
	;
	v1632 = v1619
	v1633 = v1616
	goto L301
L312:
	;
	goto L311
L313:
	;
	if v1632&int32(1) == int32(0) {
		goto L318
	} else {
		goto L319
	}
L314:
	;
	v1941 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v1438)+52)) = uint8(v1941)
	v1943 = *(*int64)(unsafe.Add(mBase, uint32(v1438)+8))
	*(*int64)(unsafe.Add(mBase, uint32(v1187)+16)) = v1943
	v1945 = *(*int64)(unsafe.Add(mBase, uint32(v1438)))
	*(*int64)(unsafe.Add(mBase, uint32(v1187)+8)) = v1945
	v1947 = *(*int32)(unsafe.Add(mBase, uint32(v1438)+16))
	*(*int32)(unsafe.Add(mBase, uint32(v1187)+24)) = v1947
	F_RegisterTwoPhaseRecord(m, int32(1), v1187+int32(8), int32(20))
	mBase = m.M
	v1954 = m.ExcPending
	if v1954 != 0 {
		goto L4
	} else {
		goto L354
	}
L315:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v1438)+28)) = v1901
	v1910 = *(*int32)(unsafe.Add(mBase, uint32(v1901)))
	*(*int32)(unsafe.Add(mBase, uint32(v1438)+24)) = v1910
	goto L314
L316:
	;
	F_LWLockRelease(m, v1684+int32(548))
	mBase = m.M
	v1868 = m.ExcPending
	if v1868 != 0 {
		goto L4
	} else {
		goto L347
	}
L317:
	;
	v1791 = *(*int32)(unsafe.Add(mBase, uint32(v1438)+16))
	v1797 = int64(1) << (uint(base.I64_extend_i32_u(v1791+base.I32_wrap_i64(v1788)-int32(1))) % 64)
	if v1797&v1789 == int64(0) {
		goto L316
	} else {
		goto L338
	}
L318:
	;
	v1662 = *(*int32)(unsafe.Add(mBase, uint32(v1438)+28))
	if v1662 != 0 {
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
	v1775 = m.ExcPending
	if v1775 != 0 {
		goto L4
	} else {
		goto L334
	}
L321:
	;
	v1664 = *(*int32)(unsafe.Add(mBase, _c_F_PrepareTransaction[18]))
	v1665 = *(*int32)(unsafe.Add(mBase, uint32(v1438)+20))
	v1667 = *(*int32)(unsafe.Add(mBase, _c_F_PrepareTransaction[30]))
	v1668 = *(*int32)(unsafe.Add(mBase, uint32(v1438)+4))
	v1670 = *(*int32)(unsafe.Add(mBase, _c_F_PrepareTransaction[31]))
	v1674 = F_LWLockAcquire(m, v1670+int32(548), int32(0))
	mBase = m.M
	v1675 = m.ExcPending
	if v1675 != 0 {
		goto L4
	} else {
		goto L322
	}
L322:
	;
	v1676 = int32(268435455)
	v1680 = (v1667 + v1676) & (v1668 * int32(_a_F_PrepareTransaction_23))
	v1682 = v1680 & v1676
	v1684 = *(*int32)(unsafe.Add(mBase, _c_F_PrepareTransaction[31]))
	v1685 = *(*int32)(unsafe.Add(mBase, uint32(v1684)+568))
	v1688 = v1685 + v1680<<(uint(int32(6))%32)
	v1695 = v1664 + v1665&int32(15)<<(uint(int32(7))%32) + int32(_a_F_PrepareTransaction_24)
	v1724 = int64(0)
	goto L323
L323:
	;
	v1730 = *(*int32)(unsafe.Add(mBase, uint32(v1688+base.I32_wrap_i64(v1724)<<(uint(int32(2))%32))))
	if v1730 == v1668 {
		goto L325
	} else {
		goto L326
	}
L324:
	;
	goto L316
L325:
	;
	v1732 = *(*int32)(unsafe.Add(mBase, uint32(v1684)+564))
	v1736 = *(*int64)(unsafe.Add(mBase, uint32(v1732+v1682<<(uint(int32(3))%32))))
	v1738 = v1724 * int64(3)
	if int64(base.Ui64(v1736)>>(uint(v1738)%64))&int64(7) != int64(0) {
		v1788 = v1738
		v1789 = v1736
		goto L317
	} else {
		goto L328
	}
L326:
	;
	goto L327
L327:
	;
	v1747 = v1724 | int64(1)
	v1752 = *(*int32)(unsafe.Add(mBase, uint32(v1688+base.I32_wrap_i64(v1747)<<(uint(int32(2))%32))))
	if v1752 == v1668 {
		goto L329
	} else {
		goto L330
	}
L328:
	;
	goto L327
L329:
	;
	v1754 = *(*int32)(unsafe.Add(mBase, uint32(v1684)+564))
	v1758 = *(*int64)(unsafe.Add(mBase, uint32(v1754+v1682<<(uint(int32(3))%32))))
	v1760 = v1747 * int64(3)
	if int64(base.Ui64(v1758)>>(uint(v1760)%64))&int64(7) != int64(0) {
		v1788 = v1760
		v1789 = v1758
		goto L317
	} else {
		goto L332
	}
L330:
	;
	goto L331
L331:
	;
	v1769 = v1724 + int64(2)
	if v1769 != int64(16) {
		v1724 = v1769
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
	v1778 = m.ExcPending
	if v1778 != 0 {
		goto L4
	} else {
		goto L335
	}
L335:
	;
	F_errmsg(m, int32(_a_F_PrepareTransaction_25), int32(0))
	mBase = m.M
	v1782 = m.ExcPending
	if v1782 != 0 {
		goto L4
	} else {
		goto L336
	}
L336:
	;
	F_errfinish(m, int32(_a_F_PrepareTransaction_26), int32(3564), int32(_a_F_PrepareTransaction_27))
	mBase = m.M
	v1787 = m.ExcPending
	if v1787 != 0 {
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
	v1802 = F_LWLockAcquire(m, v1695, int32(0))
	mBase = m.M
	v1803 = m.ExcPending
	if v1803 != 0 {
		goto L4
	} else {
		goto L339
	}
L339:
	;
	v1806 = *(*int32)(unsafe.Add(mBase, _c_F_PrepareTransaction[31]))
	v1807 = *(*int32)(unsafe.Add(mBase, uint32(v1438)+20))
	v1808 = F_SetupLockInTable(m, int32(_a_F_PrepareTransaction_28), v1806, v1438, v1807, v1791)
	mBase = m.M
	v1809 = m.ExcPending
	if v1809 != 0 {
		goto L4
	} else {
		goto L340
	}
L340:
	;
	if v1808 == int32(0) {
		goto L290
	} else {
		goto L341
	}
L341:
	;
	v1812 = *(*int32)(unsafe.Add(mBase, uint32(v1808)))
	v1813 = *(*int32)(unsafe.Add(mBase, uint32(v1812)+128))
	v1814 = int32(1)
	*(*int32)(unsafe.Add(mBase, uint32(v1812)+128)) = v1813 + v1814
	v1819 = v1812 + v1791<<(uint(int32(2))%32)
	v1821 = v1819 + int32(88)
	v1822 = *(*int32)(unsafe.Add(mBase, uint32(v1821)))
	*(*int32)(unsafe.Add(mBase, uint32(v1821))) = v1822 + v1814
	v1827 = v1814 << (uint(v1791) % 32)
	v1828 = *(*int32)(unsafe.Add(mBase, uint32(v1812)+16))
	*(*int32)(unsafe.Add(mBase, uint32(v1812)+16)) = v1827 | v1828
	v1831 = *(*int32)(unsafe.Add(mBase, uint32(v1821)))
	v1832 = *(*int32)(unsafe.Add(mBase, uint32(v1819)+44))
	if v1831 == v1832 {
		goto L342
	} else {
		goto L343
	}
L342:
	;
	v1834 = *(*int32)(unsafe.Add(mBase, uint32(v1812)+20))
	*(*int32)(unsafe.Add(mBase, uint32(v1812)+20)) = v1834 & (v1827 ^ int32(-1))
	goto L344
L343:
	;
	goto L344
L344:
	;
	v1839 = *(*int32)(unsafe.Add(mBase, uint32(v1808)+12))
	*(*int32)(unsafe.Add(mBase, uint32(v1808)+12)) = v1839 | v1827
	v1843 = *(*int32)(unsafe.Add(mBase, _c_F_PrepareTransaction[31]))
	v1844 = *(*int32)(unsafe.Add(mBase, uint32(v1843)+564))
	v1847 = v1844 + v1682<<(uint(int32(3))%32)
	v1848 = *(*int64)(unsafe.Add(mBase, uint32(v1847)))
	*(*int64)(unsafe.Add(mBase, uint32(v1847))) = v1848 & (v1797 ^ int64(-1))
	F_LWLockRelease(m, v1695)
	mBase = m.M
	v1854 = m.ExcPending
	if v1854 != 0 {
		goto L4
	} else {
		goto L345
	}
L345:
	;
	v1856 = *(*int32)(unsafe.Add(mBase, _c_F_PrepareTransaction[31]))
	F_LWLockRelease(m, v1856+int32(548))
	mBase = m.M
	v1860 = m.ExcPending
	if v1860 != 0 {
		goto L4
	} else {
		goto L346
	}
L346:
	;
	v1901 = v1808
	goto L315
L347:
	;
	v1870 = F_LWLockAcquire(m, v1695, int32(1))
	mBase = m.M
	v1871 = m.ExcPending
	if v1871 != 0 {
		goto L4
	} else {
		goto L348
	}
L348:
	;
	v1873 = *(*int32)(unsafe.Add(mBase, _c_F_PrepareTransaction[32]))
	v1874 = *(*int32)(unsafe.Add(mBase, uint32(v1438)+20))
	v1875 = int32(0)
	v1877 = F_hash_search_with_hash_value(m, v1873, v1438, v1874, v1875, v1875)
	mBase = m.M
	v1878 = m.ExcPending
	if v1878 != 0 {
		goto L4
	} else {
		goto L349
	}
L349:
	;
	if v1877 == int32(0) {
		goto L289
	} else {
		goto L350
	}
L350:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v1187)+8)) = v1877
	v1883 = *(*int32)(unsafe.Add(mBase, _c_F_PrepareTransaction[31]))
	*(*int32)(unsafe.Add(mBase, uint32(v1187)+12)) = v1883
	v1886 = *(*int32)(unsafe.Add(mBase, _c_F_PrepareTransaction[33]))
	v1889 = *(*int32)(unsafe.Add(mBase, uint32(v1438)+20))
	v1893 = int32(0)
	v1895 = F_hash_search_with_hash_value(m, v1886, v1187+int32(8), v1889^v1883<<(uint(int32(4))%32), v1893, v1893)
	mBase = m.M
	v1896 = m.ExcPending
	if v1896 != 0 {
		goto L4
	} else {
		goto L351
	}
L351:
	;
	if v1895 == int32(0) {
		goto L288
	} else {
		goto L352
	}
L352:
	;
	F_LWLockRelease(m, v1695)
	mBase = m.M
	v1900 = m.ExcPending
	if v1900 != 0 {
		goto L4
	} else {
		goto L353
	}
L353:
	;
	v1901 = v1895
	goto L315
L354:
	;
	goto L297
L355:
	;
	if v1986 != 0 {
		v1438 = v1986
		goto L295
	} else {
		goto L356
	}
L356:
	;
	goto L296
L357:
	;
	v2023 = *(*int32)(unsafe.Add(mBase, _c_F_PrepareTransaction[31]))
	F_LWLockRelease(m, v2023+int32(548))
	mBase = m.M
	v2027 = m.ExcPending
	if v2027 != 0 {
		goto L4
	} else {
		goto L358
	}
L358:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v2031 = m.ExcPending
	if v2031 != 0 {
		goto L4
	} else {
		goto L359
	}
L359:
	;
	F_errcode(m, int32(_a_F_PrepareTransaction_12))
	mBase = m.M
	v2034 = m.ExcPending
	if v2034 != 0 {
		goto L4
	} else {
		goto L360
	}
L360:
	;
	F_errmsg(m, int32(_a_F_PrepareTransaction_29), int32(0))
	mBase = m.M
	v2038 = m.ExcPending
	if v2038 != 0 {
		goto L4
	} else {
		goto L361
	}
L361:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v1187))) = int32(_a_F_PrepareTransaction_30)
	F_errhint(m, int32(_a_F_PrepareTransaction_31), v1187)
	mBase = m.M
	v2043 = m.ExcPending
	if v2043 != 0 {
		goto L4
	} else {
		goto L362
	}
L362:
	;
	F_errfinish(m, int32(_a_F_PrepareTransaction_26), int32(3041), int32(_a_F_PrepareTransaction_32))
	mBase = m.M
	v2048 = m.ExcPending
	if v2048 != 0 {
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
	v2056 = m.ExcPending
	if v2056 != 0 {
		goto L4
	} else {
		goto L365
	}
L365:
	;
	F_errfinish(m, int32(_a_F_PrepareTransaction_26), int32(3069), int32(_a_F_PrepareTransaction_32))
	mBase = m.M
	v2061 = m.ExcPending
	if v2061 != 0 {
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
	v2069 = m.ExcPending
	if v2069 != 0 {
		goto L4
	} else {
		goto L368
	}
L368:
	;
	F_errfinish(m, int32(_a_F_PrepareTransaction_26), int32(3082), int32(_a_F_PrepareTransaction_32))
	mBase = m.M
	v2074 = m.ExcPending
	if v2074 != 0 {
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
	v2081 = m.ExcPending
	if v2081 != 0 {
		goto L4
	} else {
		goto L371
	}
L371:
	;
	F_errmsg(m, int32(_a_F_PrepareTransaction_25), int32(0))
	mBase = m.M
	v2085 = m.ExcPending
	if v2085 != 0 {
		goto L4
	} else {
		goto L372
	}
L372:
	;
	F_errfinish(m, int32(_a_F_PrepareTransaction_26), int32(3496), int32(_a_F_PrepareTransaction_35))
	mBase = m.M
	v2090 = m.ExcPending
	if v2090 != 0 {
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
	*(*int32)(unsafe.Add(mBase, uint32(v2093)+8)) = int32(0)
	v2099 = *(*int32)(unsafe.Add(mBase, uint32(v2096)+104))
	*(*int32)(unsafe.Add(mBase, uint32(v2093)+12)) = v2099
	v2101 = *(*int32)(unsafe.Add(mBase, uint32(v2096)+108))
	*(*int32)(unsafe.Add(mBase, uint32(v2093)+16)) = v2101
	F_RegisterTwoPhaseRecord(m, int32(4), v2093+int32(8), int32(24))
	mBase = m.M
	v2108 = m.ExcPending
	if v2108 != 0 {
		goto L4
	} else {
		goto L377
	}
L375:
	;
	goto L376
L376:
	;
	m.G0 = v2093 + int32(32)
	v2238 = *(*int32)(unsafe.Add(mBase, _c_F_PrepareTransaction[34]))
	if v2238 != 0 {
		goto L387
	} else {
		goto L388
	}
L377:
	;
	v2110 = *(*int32)(unsafe.Add(mBase, _c_F_PrepareTransaction[18]))
	v2114 = F_LWLockAcquire(m, v2110+int32(3840), int32(1))
	mBase = m.M
	v2115 = m.ExcPending
	if v2115 != 0 {
		goto L4
	} else {
		goto L378
	}
L378:
	;
	v2116 = *(*int32)(unsafe.Add(mBase, uint32(v2096)+52))
	if v2116 == int32(0) {
		goto L379
	} else {
		goto L380
	}
L379:
	;
	v2200 = *(*int32)(unsafe.Add(mBase, _c_F_PrepareTransaction[18]))
	F_LWLockRelease(m, v2200+int32(3840))
	mBase = m.M
	v2204 = m.ExcPending
	if v2204 != 0 {
		goto L4
	} else {
		goto L386
	}
L380:
	;
	v2120 = v2096 + int32(48)
	if v2116 == v2120 {
		goto L379
	} else {
		goto L381
	}
L381:
	;
	v2123 = v2093 + int32(12)
	v2132 = v2116
	goto L382
L382:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v2093)+8)) = int32(1)
	v2157 = *(*int32)(unsafe.Add(mBase, uint32(v2132-int32(16))))
	v2158 = *(*int64)(unsafe.Add(mBase, uint32(v2157)+8))
	*(*int64)(unsafe.Add(mBase, uint32(v2123)+8)) = v2158
	v2160 = *(*int64)(unsafe.Add(mBase, uint32(v2157)))
	*(*int64)(unsafe.Add(mBase, uint32(v2123))) = v2160
	F_RegisterTwoPhaseRecord(m, int32(4), v2093+int32(8), int32(24))
	mBase = m.M
	v2167 = m.ExcPending
	if v2167 != 0 {
		goto L4
	} else {
		goto L384
	}
L383:
	;
	goto L379
L384:
	;
	v2168 = *(*int32)(unsafe.Add(mBase, uint32(v2132)+4))
	if v2168 != v2120 {
		v2132 = v2168
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
	v2239 = m.G0
	v2241 = v2239 + int32(-64)
	m.G0 = v2241
	v2243 = *(*int32)(unsafe.Add(mBase, uint32(v2238)+20))
	if v2243 != 0 {
		goto L390
	} else {
		goto L391
	}
L388:
	;
	goto L389
L389:
	;
	v2360 = m.G0
	v2362 = v2360 - int32(16)
	m.G0 = v2362
	v2365 = *(*int32)(unsafe.Add(mBase, _c_F_PrepareTransaction[35]))
	v2367 = *(*int32)(unsafe.Add(mBase, _c_F_PrepareTransaction[36]))
	v2371 = *(*int32)(unsafe.Add(mBase, uint32(v2365+v2367<<(uint(int32(2))%32))))
	*(*int32)(unsafe.Add(mBase, uint32(v2362)+12)) = v2371
	if v2371 != 0 {
		goto L397
	} else {
		goto L398
	}
L390:
	;
	v2244 = v2243
	goto L393
L391:
	;
	goto L392
L392:
	;
	m.G0 = v2241 - int32(-64)
	goto L389
L393:
	;
	v2273 = *(*int32)(unsafe.Add(mBase, uint32(v2244)+64))
	v2274 = *(*int64)(unsafe.Add(mBase, uint32(v2244)))
	*(*int64)(unsafe.Add(mBase, uint32(v2241)+8)) = v2274
	v2276 = *(*int64)(unsafe.Add(mBase, uint32(v2244)+8))
	*(*int64)(unsafe.Add(mBase, uint32(v2241)+16)) = v2276
	v2278 = *(*int64)(unsafe.Add(mBase, uint32(v2244)+16))
	*(*int64)(unsafe.Add(mBase, uint32(v2241)+24)) = v2278
	v2280 = *(*int64)(unsafe.Add(mBase, uint32(v2244)+32))
	*(*int64)(unsafe.Add(mBase, uint32(v2241)+32)) = v2280
	v2282 = *(*int64)(unsafe.Add(mBase, uint32(v2244)+40))
	*(*int64)(unsafe.Add(mBase, uint32(v2241)+40)) = v2282
	v2284 = *(*int64)(unsafe.Add(mBase, uint32(v2244)+48))
	*(*int64)(unsafe.Add(mBase, uint32(v2241)+48)) = v2284
	v2286 = *(*int32)(unsafe.Add(mBase, uint32(v2273)))
	*(*int32)(unsafe.Add(mBase, uint32(v2241)+56)) = v2286
	v2288 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v2273)+4)))
	*(*uint8)(unsafe.Add(mBase, uint32(v2241)+60)) = uint8(v2288)
	v2290 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v2244)+24)))
	*(*uint8)(unsafe.Add(mBase, uint32(v2241)+61)) = uint8(v2290)
	F_RegisterTwoPhaseRecord(m, int32(2), v2239+int32(-56), int32(56))
	mBase = m.M
	v2297 = m.ExcPending
	if v2297 != 0 {
		goto L4
	} else {
		goto L395
	}
L394:
	;
	goto L392
L395:
	;
	v2298 = *(*int32)(unsafe.Add(mBase, uint32(v2244)+68))
	if v2298 != 0 {
		v2244 = v2298
		goto L393
	} else {
		goto L396
	}
L396:
	;
	goto L394
L397:
	;
	F_RegisterTwoPhaseRecord(m, int32(3), v2362+int32(12), int32(4))
	mBase = m.M
	v2378 = m.ExcPending
	if v2378 != 0 {
		goto L4
	} else {
		goto L400
	}
L398:
	;
	goto L399
L399:
	;
	m.G0 = v2362 + int32(16)
	v2383 = *(*int32)(unsafe.Add(mBase, _c_F_PrepareTransaction[37]))
	if v2383 != 0 {
		goto L402
	} else {
		goto L403
	}
L400:
	;
	goto L399
L401:
	;
	v2409 = *(*int32)(unsafe.Add(mBase, _c_F_PrepareTransaction[21]))
	if base.Ui32(int32(8)) <= base.Ui32(v2409) {
		goto L412
	} else {
		goto L413
	}
L402:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v2395 = m.ExcPending
	if v2395 != 0 {
		goto L4
	} else {
		goto L407
	}
L403:
	;
	v2385 = *(*int32)(unsafe.Add(mBase, _c_F_PrepareTransaction[38]))
	if v2385 != 0 {
		goto L402
	} else {
		goto L404
	}
L404:
	;
	v2387 = *(*int32)(unsafe.Add(mBase, _c_F_PrepareTransaction[39]))
	if v2387 != 0 {
		goto L402
	} else {
		goto L405
	}
L405:
	;
	v2389 = *(*int32)(unsafe.Add(mBase, _c_F_PrepareTransaction[40]))
	if v2389 == int32(0) {
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
	v2398 = m.ExcPending
	if v2398 != 0 {
		goto L4
	} else {
		goto L408
	}
L408:
	;
	F_errmsg(m, int32(_a_F_PrepareTransaction_36), int32(0))
	mBase = m.M
	v2402 = m.ExcPending
	if v2402 != 0 {
		goto L4
	} else {
		goto L409
	}
L409:
	;
	F_errfinish(m, int32(_a_F_PrepareTransaction_37), int32(597), int32(_a_F_PrepareTransaction_38))
	mBase = m.M
	v2407 = m.ExcPending
	if v2407 != 0 {
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
	v2445 = *(*int32)(unsafe.Add(mBase, uint32(v2442)+4))
	v2446 = v2443 + v2445
	v2447 = int32(0)
	*(*uint16)(unsafe.Add(mBase, uint32(v2446)+6)) = uint16(v2447)
	*(*uint8)(unsafe.Add(mBase, uint32(v2446)+4)) = uint8(v2447)
	*(*int32)(unsafe.Add(mBase, uint32(v2446))) = v2447
	v2453 = *(*int32)(unsafe.Add(mBase, uint32(v2442)+4))
	v2454 = int32(8)
	*(*int32)(unsafe.Add(mBase, uint32(v2442)+4)) = v2453 + v2454
	*(*int32)(unsafe.Add(mBase, _c_F_PrepareTransaction[21])) = v2444 - v2454
	v2461 = int32(_a_F_PrepareTransaction_15)
	v2463 = *(*int32)(unsafe.Add(mBase, _c_F_PrepareTransaction[24]))
	*(*int32)(unsafe.Add(mBase, _c_F_PrepareTransaction[24])) = v2463 + v2454
	v2468 = *(*int32)(unsafe.Add(mBase, _c_F_PrepareTransaction[20]))
	v2469 = *(*int32)(unsafe.Add(mBase, uint32(v2468)))
	v2471 = v2463 + int32(12)
	*(*int32)(unsafe.Add(mBase, uint32(v2469)+4)) = v2471
	v2474 = int32(*(*uint16)(unsafe.Add(mBase, _c_F_PrepareTransaction[41])))
	v2476 = v2474 - int32(1)
	if base.Ui32(v2476&int32(_a_F_PrepareTransaction_39)) <= base.Ui32(int32(_a_F_PrepareTransaction_40)) {
		goto L417
	} else {
		goto L418
	}
L412:
	;
	v2413 = *(*int32)(unsafe.Add(mBase, _c_F_PrepareTransaction[23]))
	v2414 = *(*int32)(unsafe.Add(mBase, uint32(v2413)))
	v2442 = v2413
	v2443 = v2414
	v2444 = v2409
	goto L411
L413:
	;
	goto L414
L414:
	;
	v2416 = F_palloc0(m, int32(12))
	mBase = m.M
	v2417 = m.ExcPending
	if v2417 != 0 {
		goto L4
	} else {
		goto L415
	}
L415:
	;
	v2418 = int32(_a_F_PrepareTransaction_17)
	v2419 = *(*int32)(unsafe.Add(mBase, _c_F_PrepareTransaction[23]))
	*(*int32)(unsafe.Add(mBase, uint32(v2419)+8)) = v2416
	*(*int32)(unsafe.Add(mBase, _c_F_PrepareTransaction[23])) = v2416
	*(*int64)(unsafe.Add(mBase, uint32(v2416)+4)) = int64(0)
	v2426 = int32(512)
	*(*int32)(unsafe.Add(mBase, _c_F_PrepareTransaction[21])) = v2426
	v2428 = int32(_a_F_PrepareTransaction_18)
	v2430 = *(*int32)(unsafe.Add(mBase, _c_F_PrepareTransaction[22]))
	*(*int32)(unsafe.Add(mBase, _c_F_PrepareTransaction[22])) = v2430 + int32(1)
	v2435 = F_palloc(m, v2426)
	mBase = m.M
	v2436 = m.ExcPending
	if v2436 != 0 {
		goto L4
	} else {
		goto L416
	}
L416:
	;
	v2438 = *(*int32)(unsafe.Add(mBase, _c_F_PrepareTransaction[23]))
	*(*int32)(unsafe.Add(mBase, uint32(v2438))) = v2435
	v2441 = *(*int32)(unsafe.Add(mBase, _c_F_PrepareTransaction[21]))
	v2442 = v2438
	v2443 = v2435
	v2444 = v2441
	goto L411
L417:
	;
	v2482 = *(*int64)(unsafe.Add(mBase, _c_F_PrepareTransaction[42]))
	*(*int64)(unsafe.Add(mBase, uint32(v2469)+56)) = v2482
	v2485 = *(*int64)(unsafe.Add(mBase, _c_F_PrepareTransaction[43]))
	*(*int64)(unsafe.Add(mBase, uint32(v2469)+64)) = v2485
	goto L419
L418:
	;
	goto L419
L419:
	;
	if base.Ui32(v2471) < base.Ui32(int32(1073741824)) {
		goto L421
	} else {
		goto L422
	}
L420:
	;
	*(*int64)(unsafe.Add(mBase, _c_F_PrepareTransaction[44])) = int64(0)
	v2669 = m.G0
	v2671 = v2669 - int32(32)
	m.G0 = v2671
	v2674 = F_TwoPhaseGetDummyProc(m, v43, int32(0))
	mBase = m.M
	v2675 = m.ExcPending
	if v2675 != 0 {
		goto L4
	} else {
		goto L448
	}
L421:
	;
	v2491 = *(*int32)(unsafe.Add(mBase, _c_F_PrepareTransaction[22]))
	F_XLogEnsureRecordSpace(m, int32(0), v2491)
	mBase = m.M
	v2493 = m.ExcPending
	if v2493 != 0 {
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
	v2653 = m.ExcPending
	if v2653 != 0 {
		goto L4
	} else {
		goto L444
	}
L424:
	;
	v2494 = int32(_a_F_PrepareTransaction_41)
	v2496 = *(*int32)(unsafe.Add(mBase, _c_F_PrepareTransaction[45]))
	v2497 = int32(1)
	*(*int32)(unsafe.Add(mBase, _c_F_PrepareTransaction[45])) = v2496 + v2497
	v2501 = *(*int32)(unsafe.Add(mBase, _c_F_PrepareTransaction[31]))
	v2502 = *(*int32)(unsafe.Add(mBase, uint32(v2501)+336))
	*(*int32)(unsafe.Add(mBase, uint32(v2501)+336)) = v2502 | v2497
	F_XLogBeginInsert(m)
	mBase = m.M
	v2507 = m.ExcPending
	if v2507 != 0 {
		goto L4
	} else {
		goto L425
	}
L425:
	;
	v2509 = *(*int32)(unsafe.Add(mBase, _c_F_PrepareTransaction[20]))
	if v2509 != 0 {
		goto L426
	} else {
		goto L427
	}
L426:
	;
	v2510 = v2509
	goto L429
L427:
	;
	goto L428
L428:
	;
	v2574 = int32(_a_F_PrepareTransaction_42)
	v2576 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_PrepareTransaction[46])))
	v2577 = v2576 | int32(1)
	*(*uint8)(unsafe.Add(mBase, _c_F_PrepareTransaction[46])) = uint8(v2577)
	goto L433
L429:
	;
	v2539 = *(*int32)(unsafe.Add(mBase, uint32(v2510)))
	v2540 = *(*int32)(unsafe.Add(mBase, uint32(v2510)+4))
	F_XLogRegisterData(m, v2539, v2540)
	mBase = m.M
	v2542 = m.ExcPending
	if v2542 != 0 {
		goto L4
	} else {
		goto L431
	}
L430:
	;
	goto L428
L431:
	;
	v2543 = *(*int32)(unsafe.Add(mBase, uint32(v2510)+8))
	if v2543 != 0 {
		v2510 = v2543
		goto L429
	} else {
		goto L432
	}
L432:
	;
	goto L430
L433:
	;
	v2581 = F_XLogInsert(m, int32(1), int32(16))
	mBase = m.M
	v2582 = m.ExcPending
	if v2582 != 0 {
		goto L4
	} else {
		goto L434
	}
L434:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v426)+24)) = v2581
	if base.Ui32(v2476&int32(_a_F_PrepareTransaction_39)) <= base.Ui32(int32(_a_F_PrepareTransaction_40)) {
		goto L435
	} else {
		goto L436
	}
L435:
	;
	v2589 = *(*int64)(unsafe.Add(mBase, _c_F_PrepareTransaction[42]))
	F_replorigin_session_advance(m, v2589, v2581)
	mBase = m.M
	v2591 = m.ExcPending
	if v2591 != 0 {
		goto L4
	} else {
		goto L438
	}
L436:
	;
	v2593 = v2581
	goto L437
L437:
	;
	F_XLogFlush(m, v2593)
	mBase = m.M
	v2595 = m.ExcPending
	if v2595 != 0 {
		goto L4
	} else {
		goto L439
	}
L438:
	;
	v2592 = *(*int64)(unsafe.Add(mBase, uint32(v426)+24))
	v2593 = v2592
	goto L437
L439:
	;
	v2597 = *(*int64)(unsafe.Add(mBase, _c_F_PrepareTransaction[47]))
	*(*int64)(unsafe.Add(mBase, uint32(v426)+16)) = v2597
	v2600 = *(*int32)(unsafe.Add(mBase, _c_F_PrepareTransaction[18]))
	v2604 = F_LWLockAcquire(m, v2600+int32(2304), int32(0))
	mBase = m.M
	v2605 = m.ExcPending
	if v2605 != 0 {
		goto L4
	} else {
		goto L440
	}
L440:
	;
	v2606 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(v426)+48)) = uint8(v2606)
	v2609 = *(*int32)(unsafe.Add(mBase, _c_F_PrepareTransaction[18]))
	F_LWLockRelease(m, v2609+int32(2304))
	mBase = m.M
	v2613 = m.ExcPending
	if v2613 != 0 {
		goto L4
	} else {
		goto L441
	}
L441:
	;
	v2615 = *(*int32)(unsafe.Add(mBase, _c_F_PrepareTransaction[15]))
	v2616 = *(*int32)(unsafe.Add(mBase, uint32(v2615)))
	v2617 = *(*int32)(unsafe.Add(mBase, uint32(v426)+4))
	F_ProcArrayAdd(m, v2616+v2617*int32(768))
	mBase = m.M
	v2622 = m.ExcPending
	if v2622 != 0 {
		goto L4
	} else {
		goto L442
	}
L442:
	;
	v2624 = *(*int32)(unsafe.Add(mBase, _c_F_PrepareTransaction[31]))
	v2625 = *(*int32)(unsafe.Add(mBase, uint32(v2624)+336))
	*(*int32)(unsafe.Add(mBase, uint32(v2624)+336)) = v2625 & int32(-2)
	*(*int32)(unsafe.Add(mBase, _c_F_PrepareTransaction[48])) = v426
	v2631 = int32(_a_F_PrepareTransaction_41)
	v2633 = *(*int32)(unsafe.Add(mBase, _c_F_PrepareTransaction[45]))
	*(*int32)(unsafe.Add(mBase, _c_F_PrepareTransaction[45])) = v2633 - int32(1)
	v2637 = *(*int64)(unsafe.Add(mBase, uint32(v426)+24))
	F_SyncRepWaitForLSN(m, v2637, int32(0))
	mBase = m.M
	v2640 = m.ExcPending
	if v2640 != 0 {
		goto L4
	} else {
		goto L443
	}
L443:
	;
	v2642 = int32(0)
	*(*int32)(unsafe.Add(mBase, _c_F_PrepareTransaction[23])) = v2642
	*(*int32)(unsafe.Add(mBase, _c_F_PrepareTransaction[20])) = v2642
	*(*int32)(unsafe.Add(mBase, _c_F_PrepareTransaction[22])) = v2642
	goto L420
L444:
	;
	F_errcode(m, int32(261))
	mBase = m.M
	v2656 = m.ExcPending
	if v2656 != 0 {
		goto L4
	} else {
		goto L445
	}
L445:
	;
	F_errmsg(m, int32(_a_F_PrepareTransaction_43), int32(0))
	mBase = m.M
	v2660 = m.ExcPending
	if v2660 != 0 {
		goto L4
	} else {
		goto L446
	}
L446:
	;
	F_errfinish(m, int32(_a_F_PrepareTransaction_7), int32(1183), int32(_a_F_PrepareTransaction_44))
	mBase = m.M
	v2665 = m.ExcPending
	if v2665 != 0 {
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
	v2676 = int32(_a_F_PrepareTransaction_41)
	v2678 = *(*int32)(unsafe.Add(mBase, _c_F_PrepareTransaction[45]))
	*(*int32)(unsafe.Add(mBase, _c_F_PrepareTransaction[45])) = v2678 + int32(1)
	v2683 = v2671 + int32(12)
	v2685 = *(*int32)(unsafe.Add(mBase, _c_F_PrepareTransaction[28]))
	F_hash_seq_init(m, v2683, v2685)
	mBase = m.M
	v2687 = m.ExcPending
	if v2687 != 0 {
		goto L4
	} else {
		goto L449
	}
L449:
	;
	v2688 = F_hash_seq_search(m, v2683)
	mBase = m.M
	v2689 = m.ExcPending
	if v2689 != 0 {
		goto L4
	} else {
		goto L454
	}
L450:
	;
	v3573 = *(*int32)(unsafe.Add(mBase, _c_F_PrepareTransaction[31]))
	v3575 = *(*int32)(unsafe.Add(mBase, _c_F_PrepareTransaction[18]))
	v3579 = F_LWLockAcquire(m, v3575+int32(512), int32(0))
	mBase = m.M
	v3580 = m.ExcPending
	if v3580 != 0 {
		goto L4
	} else {
		goto L560
	}
L451:
	;
	F_errstart_cold(m, int32(24), int32(0))
	mBase = m.M
	v3562 = m.ExcPending
	if v3562 != 0 {
		goto L4
	} else {
		goto L557
	}
L452:
	;
	F_errstart_cold(m, int32(24), int32(0))
	mBase = m.M
	v3549 = m.ExcPending
	if v3549 != 0 {
		goto L4
	} else {
		goto L554
	}
L453:
	;
	F_errstart_cold(m, int32(24), int32(0))
	mBase = m.M
	v3533 = m.ExcPending
	if v3533 != 0 {
		goto L4
	} else {
		goto L550
	}
L454:
	;
	if v2688 != 0 {
		goto L455
	} else {
		goto L456
	}
L455:
	;
	v2696 = v2688
	goto L458
L456:
	;
	goto L457
L457:
	;
	v3026 = *(*int32)(unsafe.Add(mBase, _c_F_PrepareTransaction[18]))
	v3028 = *(*int32)(unsafe.Add(mBase, _c_F_PrepareTransaction[31]))
	v3037 = v3028
	v3038 = v3026
	v3042 = int32(0)
	goto L484
L458:
	;
	v2719 = *(*int32)(unsafe.Add(mBase, uint32(v2696)+28))
	if v2719 == int32(0) {
		goto L461
	} else {
		goto L462
	}
L459:
	;
	goto L457
L460:
	;
	v2992 = F_hash_seq_search(m, v2671+int32(12))
	mBase = m.M
	v2993 = m.ExcPending
	if v2993 != 0 {
		goto L4
	} else {
		goto L482
	}
L461:
	;
	F_RemoveLocalLock(m, v2696)
	mBase = m.M
	v2960 = m.ExcPending
	if v2960 != 0 {
		goto L4
	} else {
		goto L481
	}
L462:
	;
	v2722 = *(*int32)(unsafe.Add(mBase, uint32(v2696)+24))
	if v2722 == int32(0) {
		goto L461
	} else {
		goto L463
	}
L463:
	;
	v2725 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v2696)+14)))
	if v2725 == int32(6) {
		goto L460
	} else {
		goto L464
	}
L464:
	;
	v2728 = *(*int32)(unsafe.Add(mBase, uint32(v2696)+40))
	v2730 = v2728 - int32(1)
	if v2730 < int32(0) {
		goto L460
	} else {
		goto L465
	}
L465:
	;
	v2733 = *(*int32)(unsafe.Add(mBase, uint32(v2696)+48))
	v2734 = int32(3)
	v2735 = v2728 & v2734
	if base.Ui32(v2730) < base.Ui32(v2734) {
		goto L468
	} else {
		goto L469
	}
L466:
	;
	if v2886&int32(1) == int32(0) {
		goto L460
	} else {
		goto L478
	}
L467:
	;
	v2842 = v2813
	v2843 = v2814
	v2849 = v2820
	v2850 = int32(0)
	goto L475
L468:
	;
	v2739 = int32(0)
	v2813 = v2739
	v2814 = v2730
	v2820 = v2739
	goto L467
L469:
	;
	goto L470
L470:
	;
	v2743 = int32(0)
	v2746 = v2743
	v2747 = v2730
	v2753 = v2743
	v2758 = v2743
	goto L471
L471:
	;
	v2775 = int32(4)
	v2777 = v2733 + v2747<<(uint(v2775)%32)
	v2780 = *(*int32)(unsafe.Add(mBase, uint32(v2777-int32(48))))
	v2781 = int32(0)
	v2785 = *(*int32)(unsafe.Add(mBase, uint32(v2777-int32(32))))
	v2791 = *(*int32)(unsafe.Add(mBase, uint32(v2777-int32(16))))
	v2795 = *(*int32)(unsafe.Add(mBase, uint32(v2777)))
	v2799 = base.B2i32(v2780 == v2781) | base.B2i32(v2785 == v2781) | base.B2i32(v2791 == v2781) | base.B2i32(v2795 == v2781) | v2753
	v2805 = base.B2i32(v2785|v2780|v2791|v2795 != v2781) | v2746
	v2807 = v2747 - v2775
	v2809 = v2758 + v2775
	if v2809 != v2728&int32(-4) {
		v2746 = v2805
		v2747 = v2807
		v2753 = v2799
		v2758 = v2809
		goto L471
	} else {
		goto L473
	}
L472:
	;
	if v2735 == int32(0) {
		v2886 = v2805
		v2893 = v2799
		goto L466
	} else {
		goto L474
	}
L473:
	;
	goto L472
L474:
	;
	v2813 = v2805
	v2814 = v2807
	v2820 = v2799
	goto L467
L475:
	;
	v2874 = *(*int32)(unsafe.Add(mBase, uint32(v2733+v2843<<(uint(int32(4))%32))))
	v2875 = int32(0)
	v2877 = base.B2i32(v2874 == v2875) | v2849
	v2880 = base.B2i32(v2874 != v2875) | v2842
	v2881 = int32(1)
	v2884 = v2850 + v2881
	if v2884 != v2735 {
		v2842 = v2880
		v2843 = v2843 - v2881
		v2849 = v2877
		v2850 = v2884
		goto L475
	} else {
		goto L477
	}
L476:
	;
	v2886 = v2880
	v2893 = v2877
	goto L466
L477:
	;
	goto L476
L478:
	;
	if v2893&int32(1) != 0 {
		goto L453
	} else {
		goto L479
	}
L479:
	;
	v2921 = *(*int64)(unsafe.Add(mBase, uint32(v2696)+32))
	if v2921 <= int64(0) {
		goto L461
	} else {
		goto L480
	}
L480:
	;
	v2924 = *(*int32)(unsafe.Add(mBase, uint32(v2719)+16))
	v2926 = *(*int32)(unsafe.Add(mBase, uint32(v2696)+16))
	*(*int32)(unsafe.Add(mBase, uint32(v2719)+16)) = v2924 | int32(1)<<(uint(v2926)%32)
	goto L461
L481:
	;
	goto L460
L482:
	;
	if v2992 != 0 {
		v2696 = v2992
		goto L458
	} else {
		goto L483
	}
L483:
	;
	goto L459
L484:
	;
	v3060 = v3042 << (uint(int32(3)) % 32)
	v3061 = v3037 + v3060
	v3062 = *(*int32)(unsafe.Add(mBase, uint32(v3061)+424))
	if v3062 == int32(0) {
		v3495 = v3037
		v3496 = v3038
		goto L486
	} else {
		goto L487
	}
L485:
	;
	v3521 = int32(_a_F_PrepareTransaction_41)
	v3523 = *(*int32)(unsafe.Add(mBase, _c_F_PrepareTransaction[45]))
	*(*int32)(unsafe.Add(mBase, _c_F_PrepareTransaction[45])) = v3523 - int32(1)
	m.G0 = v2671 + int32(32)
	goto L450
L486:
	;
	v3518 = v3042 + int32(1)
	if v3518 != int32(16) {
		v3037 = v3495
		v3038 = v3496
		v3042 = v3518
		goto L484
	} else {
		goto L549
	}
L487:
	;
	v3066 = v3061 + int32(420)
	if v3062 == v3066 {
		v3495 = v3037
		v3496 = v3038
		goto L486
	} else {
		goto L488
	}
L488:
	;
	v3072 = v3038 + v3042<<(uint(int32(7))%32) + int32(_a_F_PrepareTransaction_24)
	v3074 = F_LWLockAcquire(m, v3072, int32(0))
	mBase = m.M
	v3075 = m.ExcPending
	if v3075 != 0 {
		goto L4
	} else {
		goto L489
	}
L489:
	;
	v3076 = *(*int32)(unsafe.Add(mBase, uint32(v3066)+4))
	v3077 = int32(0)
	if base.B2i32(v3076 == v3077)|base.B2i32(v3076 == v3066) == v3077 {
		goto L490
	} else {
		goto L491
	}
L490:
	;
	v3083 = v3060 + (v2674 + int32(420))
	v3084 = v3076
	goto L493
L491:
	;
	goto L492
L492:
	;
	F_LWLockRelease(m, v3072)
	mBase = m.M
	v3483 = m.ExcPending
	if v3483 != 0 {
		goto L4
	} else {
		goto L548
	}
L493:
	;
	v3113 = *(*int32)(unsafe.Add(mBase, uint32(v3084)+4))
	v3115 = v3084 - int32(28)
	v3116 = *(*int32)(unsafe.Add(mBase, uint32(v3115)))
	v3117 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v3116)+14)))
	if v3117 == int32(6) {
		goto L495
	} else {
		goto L496
	}
L494:
	;
	goto L492
L495:
	;
	if v3113 != v3066 {
		v3084 = v3113
		goto L493
	} else {
		goto L547
	}
L496:
	;
	v3122 = *(*int32)(unsafe.Add(mBase, uint32(v3084-int32(12))))
	if v3122 == int32(0) {
		goto L495
	} else {
		goto L497
	}
L497:
	;
	v3127 = *(*int32)(unsafe.Add(mBase, uint32(v3084-int32(16))))
	if v3122 != v3127 {
		goto L452
	} else {
		goto L498
	}
L498:
	;
	v3129 = *(*int32)(unsafe.Add(mBase, uint32(v3084)))
	*(*int32)(unsafe.Add(mBase, uint32(v3129)+4)) = v3113
	v3131 = *(*int32)(unsafe.Add(mBase, uint32(v3084)))
	*(*int32)(unsafe.Add(mBase, uint32(v3113))) = v3131
	*(*int32)(unsafe.Add(mBase, uint32(v2671)+8)) = v2674
	*(*int32)(unsafe.Add(mBase, uint32(v2671)+4)) = v3116
	*(*int32)(unsafe.Add(mBase, uint32(v3084-int32(20)))) = v2674
	v3139 = *(*int32)(unsafe.Add(mBase, _c_F_PrepareTransaction[33]))
	v3141 = v2671 + int32(4)
	v3142 = m.G0
	v3144 = v3142 - int32(32)
	m.G0 = v3144
	v3146 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v3139)+37)))
	if v3146 != int32(1) {
		goto L502
	} else {
		goto L503
	}
L499:
	;
	if v3346 == int32(0) {
		goto L451
	} else {
		goto L543
	}
L500:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v3398 = m.ExcPending
	if v3398 != 0 {
		goto L4
	} else {
		goto L540
	}
L501:
	;
	F_hash_corrupted(m, v3139)
	mBase = m.M
	v3394 = m.ExcPending
	if v3394 != 0 {
		goto L4
	} else {
		goto L539
	}
L502:
	;
	v3149 = *(*int32)(unsafe.Add(mBase, uint32(v3139)))
	v3150 = *(*int32)(unsafe.Add(mBase, uint32(v3149)+788))
	v3152 = v3115 - int32(4)
	v3153 = *(*int32)(unsafe.Add(mBase, uint32(v3152)))
	v3154 = v3150 & v3153
	v3155 = *(*int32)(unsafe.Add(mBase, uint32(v3149)+784))
	if base.Ui32(v3155) < base.Ui32(v3154) {
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
	v3353 = m.ExcPending
	if v3353 != 0 {
		goto L4
	} else {
		goto L536
	}
L505:
	;
	v3157 = *(*int32)(unsafe.Add(mBase, uint32(v3149)+792))
	v3159 = v3157 & v3154
	goto L507
L506:
	;
	v3159 = v3154
	goto L507
L507:
	;
	v3160 = *(*int32)(unsafe.Add(mBase, uint32(v3139)+4))
	v3166 = *(*int32)(unsafe.Add(mBase, uint32(v3160+int32(base.Ui32(v3159)>>(uint(int32(6))%32))&int32(67108860))))
	if v3166 == int32(0) {
		goto L501
	} else {
		goto L508
	}
L508:
	;
	v3170 = v3115 - int32(8)
	v3177 = v3166 + v3159&int32(255)<<(uint(int32(2))%32)
	goto L509
L509:
	;
	v3205 = *(*int32)(unsafe.Add(mBase, uint32(v3177)))
	if v3205 != v3170 {
		goto L511
	} else {
		goto L512
	}
L510:
	;
	if v3205 == int32(0) {
		goto L500
	} else {
		goto L515
	}
L511:
	;
	v3208 = v3205
	goto L513
L512:
	;
	v3208 = int32(0)
	goto L513
L513:
	;
	if v3208 != 0 {
		v3177 = v3205
		goto L509
	} else {
		goto L514
	}
L514:
	;
	goto L510
L515:
	;
	v3211 = *(*int32)(unsafe.Add(mBase, uint32(v3139)+40))
	v3212 = *(*int32)(unsafe.Add(mBase, uint32(v3139)+8))
	v3213 = m.T0[v3212].(func(*base.Module, int32, int32) int32)(m, v3141, v3211)
	mBase = m.M
	v3214 = m.ExcPending
	if v3214 != 0 {
		goto L4
	} else {
		goto L516
	}
L516:
	;
	v3215 = *(*int32)(unsafe.Add(mBase, uint32(v3139)))
	v3216 = *(*int32)(unsafe.Add(mBase, uint32(v3215)+788))
	v3217 = v3213 & v3216
	v3218 = *(*int32)(unsafe.Add(mBase, uint32(v3215)+784))
	if base.Ui32(v3218) < base.Ui32(v3217) {
		goto L517
	} else {
		goto L518
	}
L517:
	;
	v3220 = *(*int32)(unsafe.Add(mBase, uint32(v3215)+792))
	v3222 = v3220 & v3217
	goto L519
L518:
	;
	v3222 = v3217
	goto L519
L519:
	;
	v3223 = *(*int32)(unsafe.Add(mBase, uint32(v3139)+4))
	v3229 = *(*int32)(unsafe.Add(mBase, uint32(v3223+int32(base.Ui32(v3222)>>(uint(int32(6))%32))&int32(67108860))))
	if v3229 == int32(0) {
		goto L501
	} else {
		goto L520
	}
L520:
	;
	v3232 = *(*int32)(unsafe.Add(mBase, uint32(v3139)+40))
	v3237 = v3229 + v3222&int32(255)<<(uint(int32(2))%32)
	v3238 = *(*int32)(unsafe.Add(mBase, uint32(v3237)))
	if v3238 != 0 {
		goto L522
	} else {
		goto L523
	}
L521:
	;
	m.G0 = v3144 + int32(32)
	goto L499
L522:
	;
	v3239 = *(*int32)(unsafe.Add(mBase, uint32(v3139)+12))
	v3243 = v3238
	goto L525
L523:
	;
	v3278 = v3237
	goto L524
L524:
	;
	if v3222 != v3159 {
		goto L532
	} else {
		goto L533
	}
L525:
	;
	v3269 = *(*int32)(unsafe.Add(mBase, uint32(v3243)+4))
	if v3269 != v3213 {
		goto L527
	} else {
		goto L528
	}
L526:
	;
	v3278 = v3243
	goto L524
L527:
	;
	v3276 = *(*int32)(unsafe.Add(mBase, uint32(v3243)))
	if v3276 != 0 {
		v3243 = v3276
		goto L525
	} else {
		goto L531
	}
L528:
	;
	v3273 = m.T0[v3239].(func(*base.Module, int32, int32, int32) int32)(m, v3243+int32(8), v3141, v3232)
	mBase = m.M
	v3274 = m.ExcPending
	if v3274 != 0 {
		goto L4
	} else {
		goto L529
	}
L529:
	;
	if v3273 != 0 {
		goto L527
	} else {
		goto L530
	}
L530:
	;
	v3346 = int32(0)
	goto L521
L531:
	;
	goto L526
L532:
	;
	v3307 = *(*int32)(unsafe.Add(mBase, uint32(v3170)))
	*(*int32)(unsafe.Add(mBase, uint32(v3177))) = v3307
	*(*int32)(unsafe.Add(mBase, uint32(v3278))) = v3170
	*(*int32)(unsafe.Add(mBase, uint32(v3170))) = int32(0)
	goto L534
L533:
	;
	goto L534
L534:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v3152))) = v3213
	v3313 = *(*int32)(unsafe.Add(mBase, uint32(v3139)+16))
	v3314 = m.T0[v3313].(func(*base.Module, int32, int32, int32) int32)(m, v3115, v3141, v3232)
	mBase = m.M
	v3315 = m.ExcPending
	if v3315 != 0 {
		goto L4
	} else {
		goto L535
	}
L535:
	;
	v3346 = int32(1)
	goto L521
L536:
	;
	v3354 = *(*int32)(unsafe.Add(mBase, uint32(v3139)+32))
	*(*int32)(unsafe.Add(mBase, uint32(v3144))) = v3354
	F_errmsg_internal(m, int32(_a_F_PrepareTransaction_45), v3144)
	mBase = m.M
	v3358 = m.ExcPending
	if v3358 != 0 {
		goto L4
	} else {
		goto L537
	}
L537:
	;
	F_errfinish(m, int32(_a_F_PrepareTransaction_46), int32(1100), int32(_a_F_PrepareTransaction_47))
	mBase = m.M
	v3363 = m.ExcPending
	if v3363 != 0 {
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
	v3399 = *(*int32)(unsafe.Add(mBase, uint32(v3139)+32))
	*(*int32)(unsafe.Add(mBase, uint32(v3144)+16)) = v3399
	F_errmsg_internal(m, int32(_a_F_PrepareTransaction_48), v3144+int32(16))
	mBase = m.M
	v3405 = m.ExcPending
	if v3405 != 0 {
		goto L4
	} else {
		goto L541
	}
L541:
	;
	F_errfinish(m, int32(_a_F_PrepareTransaction_46), int32(1121), int32(_a_F_PrepareTransaction_47))
	mBase = m.M
	v3410 = m.ExcPending
	if v3410 != 0 {
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
	v3413 = *(*int32)(unsafe.Add(mBase, uint32(v3083)+4))
	if v3413 == int32(0) {
		goto L544
	} else {
		goto L545
	}
L544:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v3083)+4)) = v3083
	*(*int32)(unsafe.Add(mBase, uint32(v3083))) = v3083
	goto L546
L545:
	;
	goto L546
L546:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v3084)+4)) = v3083
	v3419 = *(*int32)(unsafe.Add(mBase, uint32(v3083)))
	*(*int32)(unsafe.Add(mBase, uint32(v3084))) = v3419
	*(*int32)(unsafe.Add(mBase, uint32(v3419)+4)) = v3084
	*(*int32)(unsafe.Add(mBase, uint32(v3083))) = v3084
	goto L495
L547:
	;
	goto L494
L548:
	;
	v3485 = *(*int32)(unsafe.Add(mBase, _c_F_PrepareTransaction[18]))
	v3487 = *(*int32)(unsafe.Add(mBase, _c_F_PrepareTransaction[31]))
	v3495 = v3487
	v3496 = v3485
	goto L486
L549:
	;
	goto L485
L550:
	;
	F_errcode(m, int32(1088))
	mBase = m.M
	v3536 = m.ExcPending
	if v3536 != 0 {
		goto L4
	} else {
		goto L551
	}
L551:
	;
	F_errmsg(m, int32(_a_F_PrepareTransaction_25), int32(0))
	mBase = m.M
	v3540 = m.ExcPending
	if v3540 != 0 {
		goto L4
	} else {
		goto L552
	}
L552:
	;
	F_errfinish(m, int32(_a_F_PrepareTransaction_26), int32(3680), int32(_a_F_PrepareTransaction_49))
	mBase = m.M
	v3545 = m.ExcPending
	if v3545 != 0 {
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
	v3553 = m.ExcPending
	if v3553 != 0 {
		goto L4
	} else {
		goto L555
	}
L555:
	;
	F_errfinish(m, int32(_a_F_PrepareTransaction_26), int32(3739), int32(_a_F_PrepareTransaction_49))
	mBase = m.M
	v3558 = m.ExcPending
	if v3558 != 0 {
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
	v3566 = m.ExcPending
	if v3566 != 0 {
		goto L4
	} else {
		goto L558
	}
L558:
	;
	F_errfinish(m, int32(_a_F_PrepareTransaction_26), int32(3777), int32(_a_F_PrepareTransaction_49))
	mBase = m.M
	v3571 = m.ExcPending
	if v3571 != 0 {
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
	v3582 = *(*int32)(unsafe.Add(mBase, _c_F_PrepareTransaction[15]))
	v3583 = *(*int32)(unsafe.Add(mBase, uint32(v3582)+4))
	v3584 = *(*int32)(unsafe.Add(mBase, uint32(v3573)+32))
	v3588 = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v3583+v3584<<(uint(int32(2))%32)))) = v3588
	*(*int32)(unsafe.Add(mBase, uint32(v3573)+52)) = v3588
	*(*int64)(unsafe.Add(mBase, uint32(v3573)+44)) = int64(0)
	v3595 = *(*int32)(unsafe.Add(mBase, _c_F_PrepareTransaction[49]))
	v3596 = *(*int64)(unsafe.Add(mBase, uint32(v3595)+56))
	*(*int64)(unsafe.Add(mBase, uint32(v3595)+56)) = v3596 + int64(1)
	v3600 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v3573)+56)))
	if v3600 == v3588 {
		goto L562
	} else {
		goto L563
	}
L561:
	;
	v3624 = *(*int32)(unsafe.Add(mBase, _c_F_PrepareTransaction[18]))
	F_LWLockRelease(m, v3624+int32(512))
	mBase = m.M
	v3628 = m.ExcPending
	if v3628 != 0 {
		goto L4
	} else {
		goto L566
	}
L562:
	;
	v3603 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v3573)+57)))
	if v3603 != int32(1) {
		goto L561
	} else {
		goto L565
	}
L563:
	;
	goto L564
L564:
	;
	v3607 = v3584 << (uint(int32(1)) % 32)
	v3608 = int32(_a_F_PrepareTransaction_52)
	v3609 = *(*int32)(unsafe.Add(mBase, _c_F_PrepareTransaction[15]))
	v3610 = *(*int32)(unsafe.Add(mBase, uint32(v3609)+8))
	v3612 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v3607+v3610))) = uint8(v3612)
	v3615 = *(*int32)(unsafe.Add(mBase, _c_F_PrepareTransaction[15]))
	v3616 = *(*int32)(unsafe.Add(mBase, uint32(v3615)+8))
	*(*uint8)(unsafe.Add(mBase, uint32(v3616+v3607)+1)) = uint8(v3612)
	*(*uint16)(unsafe.Add(mBase, uint32(v3573)+56)) = uint16(v3612)
	goto L561
L565:
	;
	goto L564
L566:
	;
	v3630 = *(*int32)(unsafe.Add(mBase, _c_F_PrepareTransaction[7]))
	if v3630 != 0 {
		goto L567
	} else {
		goto L568
	}
L567:
	;
	v3632 = v3630
	goto L570
L568:
	;
	goto L569
L569:
	;
	v3696 = *(*int32)(unsafe.Add(mBase, _c_F_PrepareTransaction[50]))
	v3697 = int32(1)
	F_ResourceOwnerReleaseInternal(m, v3696, v3697, v3697, v3697)
	mBase = m.M
	v3701 = m.ExcPending
	if v3701 != 0 {
		goto L4
	} else {
		goto L574
	}
L570:
	;
	v3660 = *(*int32)(unsafe.Add(mBase, uint32(v3632)))
	v3662 = *(*int32)(unsafe.Add(mBase, uint32(v3632)+8))
	v3663 = *(*int32)(unsafe.Add(mBase, uint32(v3632)+4))
	m.T0[v3663].(func(*base.Module, int32, int32))(m, int32(4), v3662)
	mBase = m.M
	v3665 = m.ExcPending
	if v3665 != 0 {
		goto L4
	} else {
		goto L572
	}
L571:
	;
	goto L569
L572:
	;
	if v3660 != 0 {
		v3632 = v3660
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
	v3703 = m.ExcPending
	if v3703 != 0 {
		goto L4
	} else {
		goto L575
	}
L575:
	;
	F_AtEOXact_RelationCache(m, int32(1))
	mBase = m.M
	v3706 = m.ExcPending
	if v3706 != 0 {
		goto L4
	} else {
		goto L576
	}
L576:
	;
	F_AtEOXact_TypeCache(m)
	mBase = m.M
	v3708 = m.ExcPending
	if v3708 != 0 {
		goto L4
	} else {
		goto L577
	}
L577:
	;
	v3710 = *(*int32)(unsafe.Add(mBase, _c_F_PrepareTransaction[34]))
	if v3710 != 0 {
		goto L578
	} else {
		goto L579
	}
L578:
	;
	v3711 = *(*int32)(unsafe.Add(mBase, uint32(v3710)+20))
	if v3711 != 0 {
		goto L581
	} else {
		goto L582
	}
L579:
	;
	goto L580
L580:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_PrepareTransaction[34])) = int32(0)
	F_pgstat_clear_snapshot(m)
	mBase = m.M
	v3807 = m.ExcPending
	if v3807 != 0 {
		goto L4
	} else {
		goto L587
	}
L581:
	;
	v3713 = v3711
	goto L584
L582:
	;
	goto L583
L583:
	;
	goto L580
L584:
	;
	v3741 = *(*int32)(unsafe.Add(mBase, uint32(v3713)+64))
	*(*int32)(unsafe.Add(mBase, uint32(v3741)+8)) = int32(0)
	v3744 = *(*int32)(unsafe.Add(mBase, uint32(v3713)+68))
	if v3744 != 0 {
		v3713 = v3744
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
	*(*int32)(unsafe.Add(mBase, _c_F_PrepareTransaction[51])) = int32(0)
	v3812 = *(*int32)(unsafe.Add(mBase, _c_F_PrepareTransaction[52]))
	if v3812 != 0 {
		goto L588
	} else {
		goto L589
	}
L588:
	;
	v3813 = *(*int32)(unsafe.Add(mBase, uint32(v3812)+20))
	v3814 = *(*int32)(unsafe.Add(mBase, uint32(v3812)+28))
	if v3813 < v3814 {
		goto L591
	} else {
		goto L592
	}
L589:
	;
	goto L590
L590:
	;
	v3988 = *(*int32)(unsafe.Add(mBase, _c_F_PrepareTransaction[53]))
	if v3988 != 0 {
		goto L605
	} else {
		goto L606
	}
L591:
	;
	v3817 = v3813
	goto L594
L592:
	;
	goto L593
L593:
	;
	v3884 = *(*int32)(unsafe.Add(mBase, uint32(v3812)+24))
	v3885 = *(*int32)(unsafe.Add(mBase, uint32(v3812)+32))
	if v3884 < v3885 {
		goto L598
	} else {
		goto L599
	}
L594:
	;
	v3846 = *(*int32)(unsafe.Add(mBase, _c_F_PrepareTransaction[54]))
	F_LocalExecuteInvalidationMessage(m, v3846+v3817<<(uint(int32(4))%32))
	mBase = m.M
	v3851 = m.ExcPending
	if v3851 != 0 {
		goto L4
	} else {
		goto L596
	}
L595:
	;
	goto L593
L596:
	;
	v3853 = v3817 + int32(1)
	if v3853 != v3814 {
		v3817 = v3853
		goto L594
	} else {
		goto L597
	}
L597:
	;
	goto L595
L598:
	;
	v3888 = v3884
	goto L601
L599:
	;
	goto L600
L600:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_PrepareTransaction[52])) = int32(0)
	goto L590
L601:
	;
	v3917 = *(*int32)(unsafe.Add(mBase, _c_F_PrepareTransaction[55]))
	F_LocalExecuteInvalidationMessage(m, v3917+v3888<<(uint(int32(4))%32))
	mBase = m.M
	v3922 = m.ExcPending
	if v3922 != 0 {
		goto L4
	} else {
		goto L603
	}
L602:
	;
	goto L600
L603:
	;
	v3924 = v3888 + int32(1)
	if v3924 != v3885 {
		v3888 = v3924
		goto L601
	} else {
		goto L604
	}
L604:
	;
	goto L602
L605:
	;
	v3990 = v3988
	goto L608
L606:
	;
	goto L607
L607:
	;
	v4053 = *(*int32)(unsafe.Add(mBase, _c_F_PrepareTransaction[35]))
	v4055 = *(*int32)(unsafe.Add(mBase, _c_F_PrepareTransaction[36]))
	v4059 = *(*int32)(unsafe.Add(mBase, uint32(v4053+v4055<<(uint(int32(2))%32))))
	if v4059 != 0 {
		goto L612
	} else {
		goto L613
	}
L608:
	;
	v4019 = *(*int32)(unsafe.Add(mBase, uint32(v3990)+24))
	*(*int32)(unsafe.Add(mBase, _c_F_PrepareTransaction[53])) = v4019
	F_pfree(m, v3990)
	mBase = m.M
	v4022 = m.ExcPending
	if v4022 != 0 {
		goto L4
	} else {
		goto L610
	}
L609:
	;
	goto L607
L610:
	;
	if v4019 != 0 {
		v3990 = v4019
		goto L608
	} else {
		goto L611
	}
L611:
	;
	goto L609
L612:
	;
	v4061 = F_TwoPhaseGetDummyProcNumber(m, v43, int32(0))
	mBase = m.M
	v4062 = m.ExcPending
	if v4062 != 0 {
		goto L4
	} else {
		goto L615
	}
L613:
	;
	v4094 = v4055
	goto L614
L614:
	;
	v4096 = *(*int32)(unsafe.Add(mBase, _c_F_PrepareTransaction[56]))
	v4100 = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v4096+v4094<<(uint(int32(2))%32)))) = v4100
	v4103 = int32(_a_F_PrepareTransaction_53)
	*(*int32)(unsafe.Add(mBase, _c_F_PrepareTransaction[57])) = v4103
	*(*int32)(unsafe.Add(mBase, _c_F_PrepareTransaction[58])) = v4103
	*(*int32)(unsafe.Add(mBase, _c_F_PrepareTransaction[59])) = v4100
	*(*int32)(unsafe.Add(mBase, _c_F_PrepareTransaction[60])) = v4100
	v4115 = *(*int32)(unsafe.Add(mBase, _c_F_PrepareTransaction[29]))
	if v4115 != 0 {
		goto L618
	} else {
		goto L619
	}
L615:
	;
	v4064 = *(*int32)(unsafe.Add(mBase, _c_F_PrepareTransaction[18]))
	v4068 = F_LWLockAcquire(m, v4064+int32(1664), int32(0))
	mBase = m.M
	v4069 = m.ExcPending
	if v4069 != 0 {
		goto L4
	} else {
		goto L616
	}
L616:
	;
	v4071 = *(*int32)(unsafe.Add(mBase, _c_F_PrepareTransaction[35]))
	v4072 = int32(2)
	*(*int32)(unsafe.Add(mBase, uint32(v4071+v4061<<(uint(v4072)%32)-int32(152)))) = v4059
	v4079 = *(*int32)(unsafe.Add(mBase, _c_F_PrepareTransaction[36]))
	*(*int32)(unsafe.Add(mBase, uint32(v4071+v4079<<(uint(v4072)%32)))) = int32(0)
	v4086 = *(*int32)(unsafe.Add(mBase, _c_F_PrepareTransaction[18]))
	F_LWLockRelease(m, v4086+int32(1664))
	mBase = m.M
	v4090 = m.ExcPending
	if v4090 != 0 {
		goto L4
	} else {
		goto L617
	}
L617:
	;
	v4092 = *(*int32)(unsafe.Add(mBase, _c_F_PrepareTransaction[36]))
	v4094 = v4092
	goto L614
L618:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v4115)+112)) = int64(-4294967296)
	v4119 = *(*int32)(unsafe.Add(mBase, _c_F_PrepareTransaction[61]))
	F_hash_destroy(m, v4119)
	mBase = m.M
	v4121 = m.ExcPending
	if v4121 != 0 {
		goto L4
	} else {
		goto L621
	}
L619:
	;
	goto L620
L620:
	;
	v4132 = *(*int32)(unsafe.Add(mBase, _c_F_PrepareTransaction[50]))
	v4134 = int32(1)
	F_ResourceOwnerReleaseInternal(m, v4132, int32(2), v4134, v4134)
	mBase = m.M
	v4137 = m.ExcPending
	if v4137 != 0 {
		goto L4
	} else {
		goto L622
	}
L621:
	;
	v4123 = int32(0)
	*(*int32)(unsafe.Add(mBase, _c_F_PrepareTransaction[29])) = v4123
	*(*int32)(unsafe.Add(mBase, _c_F_PrepareTransaction[61])) = v4123
	*(*uint8)(unsafe.Add(mBase, _c_F_PrepareTransaction[62])) = uint8(v4123)
	goto L620
L622:
	;
	v4139 = *(*int32)(unsafe.Add(mBase, _c_F_PrepareTransaction[50]))
	v4141 = int32(1)
	F_ResourceOwnerReleaseInternal(m, v4139, int32(3), v4141, v4141)
	mBase = m.M
	v4144 = m.ExcPending
	if v4144 != 0 {
		goto L4
	} else {
		goto L623
	}
L623:
	;
	v4146 = *(*int32)(unsafe.Add(mBase, _c_F_PrepareTransaction[18]))
	v4150 = F_LWLockAcquire(m, v4146+int32(2304), int32(0))
	mBase = m.M
	v4151 = m.ExcPending
	if v4151 != 0 {
		goto L4
	} else {
		goto L624
	}
L624:
	;
	v4153 = *(*int32)(unsafe.Add(mBase, _c_F_PrepareTransaction[48]))
	*(*int32)(unsafe.Add(mBase, uint32(v4153)+44)) = int32(-1)
	v4157 = *(*int32)(unsafe.Add(mBase, _c_F_PrepareTransaction[18]))
	F_LWLockRelease(m, v4157+int32(2304))
	mBase = m.M
	v4161 = m.ExcPending
	if v4161 != 0 {
		goto L4
	} else {
		goto L625
	}
L625:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_PrepareTransaction[48])) = int32(0)
	v4165 = int32(1)
	F_AtEOXact_GUC(m, v4165, v4165)
	mBase = m.M
	v4168 = m.ExcPending
	if v4168 != 0 {
		goto L4
	} else {
		goto L626
	}
L626:
	;
	F_AtEOXact_SPI(m, int32(1))
	mBase = m.M
	v4171 = m.ExcPending
	if v4171 != 0 {
		goto L4
	} else {
		goto L627
	}
L627:
	;
	v4173 = int32(0)
	*(*int32)(unsafe.Add(mBase, _c_F_PrepareTransaction[63])) = v4173
	*(*int32)(unsafe.Add(mBase, _c_F_PrepareTransaction[64])) = v4173
	goto L628
L628:
	;
	F_AtEOXact_on_commit_actions(m, int32(1))
	mBase = m.M
	v4180 = m.ExcPending
	if v4180 != 0 {
		goto L4
	} else {
		goto L629
	}
L629:
	;
	F_AtEOXact_Namespace(m, int32(1), int32(0))
	mBase = m.M
	v4184 = m.ExcPending
	if v4184 != 0 {
		goto L4
	} else {
		goto L630
	}
L630:
	;
	F_smgrdestroyall(m)
	mBase = m.M
	v4186 = m.ExcPending
	if v4186 != 0 {
		goto L4
	} else {
		goto L631
	}
L631:
	;
	F_AtEOXact_Files(m, int32(1))
	mBase = m.M
	v4189 = m.ExcPending
	if v4189 != 0 {
		goto L4
	} else {
		goto L632
	}
L632:
	;
	v4191 = int32(0)
	*(*int32)(unsafe.Add(mBase, _c_F_PrepareTransaction[65])) = v4191
	*(*int32)(unsafe.Add(mBase, _c_F_PrepareTransaction[66])) = v4191
	*(*int32)(unsafe.Add(mBase, _c_F_PrepareTransaction[67])) = v4191
	*(*int32)(unsafe.Add(mBase, _c_F_PrepareTransaction[68])) = v4191
	goto L633
L633:
	;
	F_AtEOXact_HashTables(m, int32(1))
	mBase = m.M
	v4204 = m.ExcPending
	if v4204 != 0 {
		goto L4
	} else {
		goto L634
	}
L634:
	;
	F_AtEOXact_RI(m)
	mBase = m.M
	v4206 = m.ExcPending
	if v4206 != 0 {
		goto L4
	} else {
		goto L635
	}
L635:
	;
	v4207 = int32(1)
	F_AtEOXact_Snapshot(m, v4207, v4207)
	mBase = m.M
	v4210 = m.ExcPending
	if v4210 != 0 {
		goto L4
	} else {
		goto L636
	}
L636:
	;
	goto L638
L637:
	;
	F_AtEOXact_LogicalRepWorkers(m, int32(0))
	mBase = m.M
	v4233 = m.ExcPending
	if v4233 != 0 {
		goto L4
	} else {
		goto L642
	}
L638:
	;
	v4229 = int32(0)
	*(*uint8)(unsafe.Add(mBase, _c_F_PrepareTransaction[69])) = uint8(v4229)
	goto L637
L642:
	;
	F_AtEOXact_LogicalCtl(m)
	mBase = m.M
	v4235 = m.ExcPending
	if v4235 != 0 {
		goto L4
	} else {
		goto L643
	}
L643:
	;
	v4239 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_PrepareTransaction[70])))
	if v4239 != int32(1) {
		goto L645
	} else {
		goto L646
	}
L644:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_PrepareTransaction[71])) = int32(0)
	v4280 = *(*int32)(unsafe.Add(mBase, _c_F_PrepareTransaction[50]))
	F_ResourceOwnerDelete(m, v4280)
	mBase = m.M
	v4282 = m.ExcPending
	if v4282 != 0 {
		goto L4
	} else {
		goto L648
	}
L645:
	;
	goto L644
L646:
	;
	v4243 = *(*int32)(unsafe.Add(mBase, _c_F_PrepareTransaction[72]))
	if v4243 == int32(0) {
		goto L645
	} else {
		goto L647
	}
L647:
	;
	v4246 = int32(_a_F_PrepareTransaction_41)
	v4248 = *(*int32)(unsafe.Add(mBase, _c_F_PrepareTransaction[45]))
	v4249 = int32(1)
	*(*int32)(unsafe.Add(mBase, _c_F_PrepareTransaction[45])) = v4248 + v4249
	v4252 = *(*int32)(unsafe.Add(mBase, uint32(v4243)))
	*(*int32)(unsafe.Add(mBase, uint32(v4243))) = v4252 + v4249
	v4256 = int32(0)
	v4258 = int32(_a_F_PrepareTransaction_54)
	v4259 = base.AtomicRmwOr32(m, v4256, v4258, v4256)
	*(*int64)(unsafe.Add(mBase, uint32(v4243)+24)) = int64(0)
	v4264 = base.AtomicRmwOr32(m, v4256, v4258, v4256)
	v4265 = *(*int32)(unsafe.Add(mBase, uint32(v4243)))
	*(*int32)(unsafe.Add(mBase, uint32(v4243))) = v4265 + v4249
	v4271 = *(*int32)(unsafe.Add(mBase, _c_F_PrepareTransaction[45]))
	*(*int32)(unsafe.Add(mBase, _c_F_PrepareTransaction[45])) = v4271 - v4249
	goto L645
L648:
	;
	v4283 = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v35)+40)) = v4283
	*(*int32)(unsafe.Add(mBase, _c_F_PrepareTransaction[50])) = v4283
	*(*int32)(unsafe.Add(mBase, _c_F_PrepareTransaction[73])) = v4283
	v4293 = *(*int32)(unsafe.Add(mBase, _c_F_PrepareTransaction[0]))
	v4294 = *(*int32)(unsafe.Add(mBase, uint32(v4293)+44))
	*(*int32)(unsafe.Add(mBase, _c_F_PrepareTransaction[27])) = v4294
	v4297 = *(*int32)(unsafe.Add(mBase, _c_F_PrepareTransaction[74]))
	F_MemoryContextReset(m, v4297)
	mBase = m.M
	v4299 = m.ExcPending
	if v4299 != 0 {
		goto L4
	} else {
		goto L649
	}
L649:
	;
	v4301 = int32(0)
	*(*int32)(unsafe.Add(mBase, _c_F_PrepareTransaction[75])) = v4301
	*(*int32)(unsafe.Add(mBase, uint32(v4293)+36)) = v4301
	*(*int32)(unsafe.Add(mBase, uint32(v35)+56)) = v4301
	v4307 = int64(0)
	*(*int64)(unsafe.Add(mBase, uint32(v35)+48)) = v4307
	*(*int64)(unsafe.Add(mBase, uint32(v35)+28)) = v4307
	*(*int32)(unsafe.Add(mBase, uint32(v35)+8)) = v4301
	*(*int64)(unsafe.Add(mBase, uint32(v35))) = v4307
	*(*int64)(unsafe.Add(mBase, _c_F_PrepareTransaction[76])) = v4307
	*(*int32)(unsafe.Add(mBase, _c_F_PrepareTransaction[77])) = v4301
	*(*int32)(unsafe.Add(mBase, uint32(v35)+20)) = v4301
	v4323 = int32(_a_F_PrepareTransaction_4)
	v4325 = *(*int32)(unsafe.Add(mBase, _c_F_PrepareTransaction[10]))
	*(*int32)(unsafe.Add(mBase, _c_F_PrepareTransaction[10])) = v4325 - int32(1)
	m.G0 = v32 + int32(16)
	return
L650:
	;
	F_errcode(m, int32(1088))
	mBase = m.M
	v4338 = m.ExcPending
	if v4338 != 0 {
		goto L4
	} else {
		goto L651
	}
L651:
	;
	F_errmsg(m, int32(_a_F_PrepareTransaction_55), int32(0))
	mBase = m.M
	v4342 = m.ExcPending
	if v4342 != 0 {
		goto L4
	} else {
		goto L652
	}
L652:
	;
	F_errfinish(m, int32(_a_F_PrepareTransaction_3), int32(2660), int32(_a_F_PrepareTransaction_0))
	mBase = m.M
	v4347 = m.ExcPending
	if v4347 != 0 {
		goto L4
	} else {
		goto L653
	}
L653:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L654:
	;
	F_errcode(m, int32(1088))
	mBase = m.M
	v4354 = m.ExcPending
	if v4354 != 0 {
		goto L4
	} else {
		goto L655
	}
L655:
	;
	F_errmsg(m, int32(_a_F_PrepareTransaction_56), int32(0))
	mBase = m.M
	v4358 = m.ExcPending
	if v4358 != 0 {
		goto L4
	} else {
		goto L656
	}
L656:
	;
	F_errfinish(m, int32(_a_F_PrepareTransaction_3), int32(2670), int32(_a_F_PrepareTransaction_0))
	mBase = m.M
	v4363 = m.ExcPending
	if v4363 != 0 {
		goto L4
	} else {
		goto L657
	}
L657:
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
	var v292 int32
	_ = v292
	var v293 int32
	_ = v293
	var v294 int32
	_ = v294
	var v295 int32
	_ = v295
	var v298 int32
	_ = v298
	var v299 int32
	_ = v299
	var v301 int32
	_ = v301
	var v303 int32
	_ = v303
	var v306 int32
	_ = v306
	var v307 int32
	_ = v307
	var v315 int32
	_ = v315
	var v320 int32
	_ = v320
	var v328 int32
	_ = v328
	var v334 int32
	_ = v334
	var v345 int32
	_ = v345
	var v346 int32
	_ = v346
	var v347 int32
	_ = v347
	var v365 int32
	_ = v365
	var v368 int32
	_ = v368
	var v369 int32
	_ = v369
	var v376 int32
	_ = v376
	var v378 int32
	_ = v378
	var v381 int32
	_ = v381
	var v382 int32
	_ = v382
	var v383 int32
	_ = v383
	var v384 int32
	_ = v384
	var v387 int32
	_ = v387
	var v388 int32
	_ = v388
	var v390 int32
	_ = v390
	var v392 int32
	_ = v392
	var v395 int32
	_ = v395
	var v396 int32
	_ = v396
	var v402 int32
	_ = v402
	var v407 int32
	_ = v407
	var v410 int32
	_ = v410
	var v413 int32
	_ = v413
	var v414 int64
	_ = v414
	var v415 int32
	_ = v415
	var v429 int64
	_ = v429
	var v434 int32
	_ = v434
	var v438 int32
	_ = v438
	var v460 int32
	_ = v460
	var v462 int32
	_ = v462
	var v465 int32
	_ = v465
	var v467 int32
	_ = v467
	var v477 int32
	_ = v477
	var v490 int32
	_ = v490
	var v494 int32
	_ = v494
	var v499 int32
	_ = v499
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
	F_errstart_cold(m, int32(24), int32(0))
	mBase = m.M
	v490 = m.ExcPending
	if v490 != 0 {
		goto L7
	} else {
		goto L110
	}
L2:
	;
	m.G0 = v18 + int32(16)
	return v477
L3:
	;
	if l0 != 0 {
		v477 = v2
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
	v477 = v2
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
	v460 = *(*int32)(unsafe.Add(mBase, uint32(v18)+12))
	if v460 != 0 {
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
	v226 = *(*int32)(unsafe.Add(mBase, uint32(v225)+32))
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
	v334 = v225
	goto L70
L70:
	;
	v345 = v334 + int32(60)
	v346 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v334)+56)))
	v347 = v346
	goto L87
L71:
	;
	v250 = *(*int32)(unsafe.Add(mBase, _c_F_RecordTransactionAbort[9]))
	v252 = v250 + int32(60)
	v256 = *(*int32)(unsafe.Add(mBase, uint32(v45+v234<<(uint(int32(2))%32))))
	v257 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v250)+56)))
	v258 = v257
	goto L75
L72:
	;
	v328 = *(*int32)(unsafe.Add(mBase, _c_F_RecordTransactionAbort[9]))
	v334 = v328
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
	v303 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v250)+57)))
	if v303 != 0 {
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
	v289 = int32(0)
	v292 = base.AtomicRmwOr32(m, v289, int32(_a_F_RecordTransactionAbort_3), v289)
	v293 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v229))))
	v294 = int32(1)
	v295 = v293 - v294
	*(*uint8)(unsafe.Add(mBase, uint32(v229))) = uint8(v295)
	v298 = *(*int32)(unsafe.Add(mBase, _c_F_RecordTransactionAbort[9]))
	v299 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v298)+56)))
	v301 = v299 - v294
	*(*uint8)(unsafe.Add(mBase, uint32(v298)+56)) = uint8(v301)
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
	v306 = F_errstart(m, int32(19), int32(0))
	mBase = m.M
	v307 = m.ExcPending
	if v307 != 0 {
		goto L7
	} else {
		goto L80
	}
L80:
	;
	if v306 == int32(0) {
		goto L73
	} else {
		goto L81
	}
L81:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v212)+16)) = v256
	F_errmsg_internal(m, int32(_a_F_RecordTransactionAbort_4), v212+int32(16))
	mBase = m.M
	v315 = m.ExcPending
	if v315 != 0 {
		goto L7
	} else {
		goto L82
	}
L82:
	;
	F_errfinish(m, int32(_a_F_RecordTransactionAbort_5), int32(4058), int32(_a_F_RecordTransactionAbort_6))
	mBase = m.M
	v320 = m.ExcPending
	if v320 != 0 {
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
	v410 = int32(3)
	v413 = *(*int32)(unsafe.Add(mBase, _c_F_RecordTransactionAbort[10]))
	v414 = *(*int64)(unsafe.Add(mBase, uint32(v413)+48))
	v415 = base.I32_wrap_i64(v414)
	if base.B2i32(base.Ui32(v205) < base.Ui32(v410))|base.B2i32(base.Ui32(v415) < base.Ui32(v410)) == int32(0) {
		goto L98
	} else {
		goto L99
	}
L86:
	;
	v392 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v334)+57)))
	if v392 != 0 {
		goto L85
	} else {
		goto L91
	}
L87:
	;
	if v347 == int32(0) {
		goto L86
	} else {
		goto L89
	}
L88:
	;
	v376 = *(*int32)(unsafe.Add(mBase, uint32(v345+v346<<(uint(int32(2))%32)-int32(4))))
	*(*int32)(unsafe.Add(mBase, uint32(v368))) = v376
	v378 = int32(0)
	v381 = base.AtomicRmwOr32(m, v378, int32(_a_F_RecordTransactionAbort_3), v378)
	v382 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v229))))
	v383 = int32(1)
	v384 = v382 - v383
	*(*uint8)(unsafe.Add(mBase, uint32(v229))) = uint8(v384)
	v387 = *(*int32)(unsafe.Add(mBase, _c_F_RecordTransactionAbort[9]))
	v388 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v387)+56)))
	v390 = v388 - v383
	*(*uint8)(unsafe.Add(mBase, uint32(v387)+56)) = uint8(v390)
	goto L85
L89:
	;
	v365 = v347 - int32(1)
	v368 = v345 + v365<<(uint(int32(2))%32)
	v369 = *(*int32)(unsafe.Add(mBase, uint32(v368)))
	if v369 != v22 {
		v347 = v365
		goto L87
	} else {
		goto L90
	}
L90:
	;
	goto L88
L91:
	;
	v395 = F_errstart(m, int32(19), int32(0))
	mBase = m.M
	v396 = m.ExcPending
	if v396 != 0 {
		goto L7
	} else {
		goto L92
	}
L92:
	;
	if v395 == int32(0) {
		goto L85
	} else {
		goto L93
	}
L93:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v212))) = v22
	F_errmsg_internal(m, int32(_a_F_RecordTransactionAbort_4), v212)
	mBase = m.M
	v402 = m.ExcPending
	if v402 != 0 {
		goto L7
	} else {
		goto L94
	}
L94:
	;
	F_errfinish(m, int32(_a_F_RecordTransactionAbort_5), int32(4074), int32(_a_F_RecordTransactionAbort_6))
	mBase = m.M
	v407 = m.ExcPending
	if v407 != 0 {
		goto L7
	} else {
		goto L95
	}
L95:
	;
	goto L85
L96:
	;
	v429 = *(*int64)(unsafe.Add(mBase, uint32(v413)+56))
	*(*int64)(unsafe.Add(mBase, uint32(v413)+56)) = v429 + int64(1)
	v434 = *(*int32)(unsafe.Add(mBase, _c_F_RecordTransactionAbort[7]))
	F_LWLockRelease(m, v434+int32(512))
	mBase = m.M
	v438 = m.ExcPending
	if v438 != 0 {
		goto L7
	} else {
		goto L103
	}
L97:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v413)+48)) = v414 + base.I64_extend_i32_s(v205-v415)
	goto L96
L98:
	;
	if v415-v205 < int32(0) {
		goto L97
	} else {
		goto L101
	}
L99:
	;
	goto L100
L100:
	;
	if base.Ui32(v205) <= base.Ui32(v415) {
		goto L96
	} else {
		goto L102
	}
L101:
	;
	goto L96
L102:
	;
	goto L97
L103:
	;
	m.G0 = v212 + int32(32)
	goto L63
L104:
	;
	F_pfree(m, v460)
	mBase = m.M
	v462 = m.ExcPending
	if v462 != 0 {
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
		v477 = v205
		goto L2
	} else {
		goto L108
	}
L107:
	;
	goto L106
L108:
	;
	v465 = *(*int32)(unsafe.Add(mBase, uint32(v18)+8))
	F_pfree(m, v465)
	mBase = m.M
	v467 = m.ExcPending
	if v467 != 0 {
		goto L7
	} else {
		goto L109
	}
L109:
	;
	v477 = v205
	goto L2
L110:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v18))) = v22
	F_errmsg_internal(m, int32(_a_F_RecordTransactionAbort_7), v18)
	mBase = m.M
	v494 = m.ExcPending
	if v494 != 0 {
		goto L7
	} else {
		goto L111
	}
L111:
	;
	F_errfinish(m, int32(_a_F_RecordTransactionAbort_8), int32(1836), int32(_a_F_RecordTransactionAbort_9))
	mBase = m.M
	v499 = m.ExcPending
	if v499 != 0 {
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
						F_errfinish(m, int32(_a_F_StartTransactionCommand_3), int32(3167), int32(_a_F_StartTransactionCommand_4))
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
func F_TransactionIdIsInProgress(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v2 int32
	_ = v2
	var v14 int32
	_ = v14
	var v26 int32
	_ = v26
	var v28 int32
	_ = v28
	var v31 int32
	_ = v31
	var v43 int32
	_ = v43
	var v47 int32
	_ = v47
	var v51 int32
	_ = v51
	var v56 int32
	_ = v56
	var v62 int32
	_ = v62
	var v65 int32
	_ = v65
	var v68 int32
	_ = v68
	var v70 int32
	_ = v70
	var v72 int32
	_ = v72
	var v75 int32
	_ = v75
	var v78 int32
	_ = v78
	var v81 int32
	_ = v81
	var v86 int32
	_ = v86
	var v87 int32
	_ = v87
	var v88 int32
	_ = v88
	var v92 int32
	_ = v92
	var v101 int32
	_ = v101
	var v102 int32
	_ = v102
	var v105 int32
	_ = v105
	var v116 int32
	_ = v116
	var v121 int32
	_ = v121
	var v123 int32
	_ = v123
	var v126 int32
	_ = v126
	var v131 int32
	_ = v131
	var v132 int32
	_ = v132
	var v133 int32
	_ = v133
	var v137 int32
	_ = v137
	var v138 int32
	_ = v138
	var v141 int32
	_ = v141
	var v142 int32
	_ = v142
	var v145 int32
	_ = v145
	var v153 int32
	_ = v153
	var v163 int32
	_ = v163
	var v167 int32
	_ = v167
	var v173 int32
	_ = v173
	var v178 int32
	_ = v178
	var v179 int32
	_ = v179
	var v181 int32
	_ = v181
	var v183 int32
	_ = v183
	var v185 int32
	_ = v185
	var v187 int32
	_ = v187
	var v191 int32
	_ = v191
	var v192 int32
	_ = v192
	var v195 int32
	_ = v195
	var v202 int32
	_ = v202
	var v203 int32
	_ = v203
	var v205 int32
	_ = v205
	var v207 int32
	_ = v207
	var v211 int32
	_ = v211
	var v214 int32
	_ = v214
	var v215 int32
	_ = v215
	var v218 int32
	_ = v218
	var v219 int32
	_ = v219
	var v232 int32
	_ = v232
	var v235 int32
	_ = v235
	var v239 int32
	_ = v239
	var v244 int32
	_ = v244
	var v245 int32
	_ = v245
	var v249 int32
	_ = v249
	var v250 int32
	_ = v250
	var v258 int32
	_ = v258
	var v261 int32
	_ = v261
	var v270 int32
	_ = v270
	var v272 int32
	_ = v272
	var v274 int32
	_ = v274
	var v289 int32
	_ = v289
	var v290 int32
	_ = v290
	var v291 int32
	_ = v291
	var v294 int32
	_ = v294
	var v296 int32
	_ = v296
	var v298 int32
	_ = v298
	var v305 int32
	_ = v305
	var v319 int32
	_ = v319
	var v323 int32
	_ = v323
	var v325 int32
	_ = v325
	var v329 int32
	_ = v329
	var v341 int32
	_ = v341
	var v349 int32
	_ = v349
	var v356 int32
	_ = v356
	var v365 int32
	_ = v365
	var v370 int32
	_ = v370
	var v371 int32
	_ = v371
	var v373 int32
	_ = v373
	var v375 int32
	_ = v375
	var v379 int32
	_ = v379
	var v380 int32
	_ = v380
	var v381 int32
	_ = v381
	var v382 int32
	_ = v382
	var v385 int32
	_ = v385
	var v387 int32
	_ = v387
	var v390 int32
	_ = v390
	var v394 int32
	_ = v394
	var v395 int32
	_ = v395
	var v405 int32
	_ = v405
	var v406 int32
	_ = v406
	var v407 int32
	_ = v407
	var v411 int32
	_ = v411
	var v424 int32
	_ = v424
	var v425 int32
	_ = v425
	var v428 int32
	_ = v428
	var v433 int32
	_ = v433
	var v435 int32
	_ = v435
	var v450 int32
	_ = v450
	var v453 int32
	_ = v453
	var v454 int32
	_ = v454
	var v465 int32
	_ = v465
	var v466 int32
	_ = v466
	var v468 int32
	_ = v468
	var v469 int32
	_ = v469
	var v470 int32
	_ = v470
	var v474 int32
	_ = v474
	var v477 int32
	_ = v477
	var v481 int32
	_ = v481
	var v482 int32
	_ = v482
	var v485 int32
	_ = v485
	var v493 int32
	_ = v493
	var v497 int32
	_ = v497
	var v501 int32
	_ = v501
	var v518 int32
	_ = v518
	var v521 int32
	_ = v521
	var v522 int32
	_ = v522
	var v524 int32
	_ = v524
	var v531 int32
	_ = v531
	var v543 int32
	_ = v543
	var v551 int32
	_ = v551
	var v555 int32
	_ = v555
	var v556 int32
	_ = v556
	var v557 int32
	_ = v557
	var v558 int32
	_ = v558
	var v559 int32
	_ = v559
	var v563 int32
	_ = v563
	var v565 int32
	_ = v565
	var v579 int32
	_ = v579
	var v580 int32
	_ = v580
	var v584 int32
	_ = v584
	var v644 int32
	_ = v644
	var v648 int32
	_ = v648
	v2 = int32(0)
	v14 = *(*int32)(unsafe.Add(mBase, _c_F_TransactionIdIsInProgress[0]))
	if base.B2i32(base.Ui32(int32(2)) < base.Ui32(v14))&base.B2i32(base.Ui32(int32(3)) <= base.Ui32(l0)) == v2 {
		goto L4
	} else {
		goto L5
	}
L1:
	;
	v644 = *(*int32)(unsafe.Add(mBase, _c_F_TransactionIdIsInProgress[1]))
	F_LWLockRelease(m, v644+int32(512))
	mBase = m.M
	v648 = m.ExcPending
	if v648 != 0 {
		goto L66
	} else {
		goto L164
	}
L2:
	;
	return int32(0)
L3:
	;
	if v26 != 0 {
		goto L2
	} else {
		goto L7
	}
L4:
	;
	v26 = base.B2i32(base.Ui32(l0) < base.Ui32(v14))
	goto L3
L5:
	;
	goto L6
L6:
	;
	v26 = int32(base.Ui32(l0-v14) >> (uint(int32(31)) % 32))
	goto L3
L7:
	;
	v28 = *(*int32)(unsafe.Add(mBase, _c_F_TransactionIdIsInProgress[2]))
	if v28 == l0 {
		goto L2
	} else {
		goto L8
	}
L8:
	;
	v31 = *(*int32)(unsafe.Add(mBase, _c_F_TransactionIdIsInProgress[3]))
	if base.Ui32(l0) < base.Ui32(int32(3)) {
		goto L10
	} else {
		goto L11
	}
L9:
	;
	if v163 != 0 {
		goto L49
	} else {
		goto L50
	}
L10:
	;
	v163 = int32(0)
	goto L9
L11:
	;
	goto L12
L12:
	;
	v43 = *(*int32)(unsafe.Add(mBase, _c_F_TransactionIdIsInProgress[4]))
	if v43 == l0 {
		goto L13
	} else {
		goto L14
	}
L13:
	;
	v163 = int32(1)
	goto L9
L14:
	;
	goto L15
L15:
	;
	v47 = *(*int32)(unsafe.Add(mBase, _c_F_TransactionIdIsInProgress[5]))
	if v47 <= int32(0) {
		goto L17
	} else {
		goto L18
	}
L16:
	;
	v163 = v153
	goto L9
L17:
	;
	v51 = *(*int32)(unsafe.Add(mBase, _c_F_TransactionIdIsInProgress[6]))
	if v51 == int32(0) {
		v153 = int32(0)
		goto L16
	} else {
		goto L20
	}
L18:
	;
	goto L19
L19:
	;
	v121 = *(*int32)(unsafe.Add(mBase, _c_F_TransactionIdIsInProgress[7]))
	v123 = int32(0)
	v126 = v47 - int32(1)
	goto L39
L20:
	;
	v56 = v51
	goto L21
L21:
	;
	v62 = *(*int32)(unsafe.Add(mBase, uint32(v56)+20))
	if v62 == int32(4) {
		goto L23
	} else {
		goto L24
	}
L22:
	;
	v153 = int32(0)
	goto L16
L23:
	;
	v116 = *(*int32)(unsafe.Add(mBase, uint32(v56)+80))
	if v116 != 0 {
		v56 = v116
		goto L21
	} else {
		goto L38
	}
L24:
	;
	v65 = *(*int32)(unsafe.Add(mBase, uint32(v56)))
	if v65 == int32(0) {
		goto L23
	} else {
		goto L25
	}
L25:
	;
	v68 = int32(1)
	if l0 == v65 {
		v153 = v68
		goto L16
	} else {
		goto L26
	}
L26:
	;
	v70 = *(*int32)(unsafe.Add(mBase, uint32(v56)+52))
	v72 = v70 - int32(1)
	if v72 < int32(0) {
		goto L23
	} else {
		goto L27
	}
L27:
	;
	v75 = *(*int32)(unsafe.Add(mBase, uint32(v56)+48))
	v78 = int32(0)
	v81 = v72
	goto L28
L28:
	;
	v86 = int32(2)
	v87 = base.I32_div_s(v81-v78, v86)
	v88 = v87 + v78
	v92 = *(*int32)(unsafe.Add(mBase, uint32(v75+v88<<(uint(v86)%32))))
	if v92 == l0 {
		v153 = v68
		goto L16
	} else {
		goto L30
	}
L29:
	;
	goto L23
L30:
	;
	v101 = base.B2i32(v92-l0 < int32(0)) | base.B2i32(base.Ui32(v92) < base.Ui32(int32(3)))
	if v101 != 0 {
		goto L31
	} else {
		goto L32
	}
L31:
	;
	v102 = v88 + int32(1)
	goto L33
L32:
	;
	v102 = v78
	goto L33
L33:
	;
	if v101 != 0 {
		goto L34
	} else {
		goto L35
	}
L34:
	;
	v105 = v81
	goto L36
L35:
	;
	v105 = v88 - int32(1)
	goto L36
L36:
	;
	if v102 <= v105 {
		v78 = v102
		v81 = v105
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
	v131 = int32(2)
	v132 = base.I32_div_s(v126-v123, v131)
	v133 = v132 + v123
	v137 = *(*int32)(unsafe.Add(mBase, uint32(v121+v133<<(uint(v131)%32))))
	v138 = base.B2i32(v137 == l0)
	if v137 == l0 {
		v153 = v138
		goto L16
	} else {
		goto L41
	}
L40:
	;
	v153 = v138
	goto L16
L41:
	;
	v141 = base.B2i32(base.Ui32(v137) < base.Ui32(l0))
	if base.Ui32(v137) < base.Ui32(l0) {
		goto L42
	} else {
		goto L43
	}
L42:
	;
	v142 = v133 + int32(1)
	goto L44
L43:
	;
	v142 = v123
	goto L44
L44:
	;
	if base.Ui32(v137) < base.Ui32(l0) {
		goto L45
	} else {
		goto L46
	}
L45:
	;
	v145 = v126
	goto L47
L46:
	;
	v145 = v133 - int32(1)
	goto L47
L47:
	;
	if v142 <= v145 {
		v123 = v142
		v126 = v145
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
	v167 = *(*int32)(unsafe.Add(mBase, _c_F_TransactionIdIsInProgress[8]))
	if v167 == int32(0) {
		goto L54
	} else {
		goto L55
	}
L52:
	;
	v245 = *(*int32)(unsafe.Add(mBase, uint32(v31)))
	if int32(0) < v245 {
		goto L78
	} else {
		goto L79
	}
L53:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v232 = m.ExcPending
	if v232 != 0 {
		goto L66
	} else {
		goto L74
	}
L54:
	;
	v173 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_TransactionIdIsInProgress[9])))
	if v173 == int32(1) {
		goto L59
	} else {
		goto L60
	}
L55:
	;
	goto L56
L56:
	;
	v202 = *(*int32)(unsafe.Add(mBase, _c_F_TransactionIdIsInProgress[10]))
	v203 = *(*int32)(unsafe.Add(mBase, uint32(v202)+4))
	*(*int32)(unsafe.Add(mBase, _c_F_TransactionIdIsInProgress[11])) = v203
	v205 = *(*int32)(unsafe.Add(mBase, uint32(v202)+8))
	v207 = *(*int32)(unsafe.Add(mBase, _c_F_TransactionIdIsInProgress[1]))
	v211 = F_LWLockAcquire(m, v207+int32(512), int32(1))
	mBase = m.M
	v214 = m.ExcPending
	if v214 != 0 {
		goto L66
	} else {
		goto L67
	}
L57:
	;
	v195 = F_emscripten_builtin_malloc(m, v192<<(uint(int32(2))%32))
	mBase = m.M
	*(*int32)(unsafe.Add(mBase, _c_F_TransactionIdIsInProgress[8])) = v195
	if v195 == int32(0) {
		goto L53
	} else {
		goto L65
	}
L58:
	;
	if v183 != 0 {
		goto L62
	} else {
		goto L63
	}
L59:
	;
	v178 = *(*int32)(unsafe.Add(mBase, _c_F_TransactionIdIsInProgress[12]))
	v179 = *(*int32)(unsafe.Add(mBase, uint32(v178)+308))
	v181 = base.B2i32(v179 != int32(2))
	*(*uint8)(unsafe.Add(mBase, _c_F_TransactionIdIsInProgress[9])) = uint8(v181)
	v183 = v181
	goto L61
L60:
	;
	v183 = int32(0)
	goto L61
L61:
	;
	goto L58
L62:
	;
	v185 = *(*int32)(unsafe.Add(mBase, _c_F_TransactionIdIsInProgress[13]))
	v187 = *(*int32)(unsafe.Add(mBase, _c_F_TransactionIdIsInProgress[14]))
	v192 = (v185 + v187) * int32(65)
	goto L57
L63:
	;
	goto L64
L64:
	;
	v191 = *(*int32)(unsafe.Add(mBase, uint32(v31)+4))
	v192 = v191
	goto L57
L65:
	;
	goto L56
L66:
	;
	return int32(0)
L67:
	;
	v215 = int32(3)
	v218 = *(*int32)(unsafe.Add(mBase, _c_F_TransactionIdIsInProgress[15]))
	v219 = *(*int32)(unsafe.Add(mBase, uint32(v218)+48))
	if base.B2i32(base.Ui32(l0) < base.Ui32(v215))|base.B2i32(base.Ui32(v219) < base.Ui32(v215)) == int32(0) {
		goto L69
	} else {
		goto L70
	}
L68:
	;
	goto L1
L69:
	;
	if v219-l0 < int32(0) {
		goto L68
	} else {
		goto L72
	}
L70:
	;
	goto L71
L71:
	;
	if base.Ui32(l0) <= base.Ui32(v219) {
		goto L52
	} else {
		goto L73
	}
L72:
	;
	goto L52
L73:
	;
	goto L68
L74:
	;
	F_errcode(m, int32(_a_F_TransactionIdIsInProgress_0))
	mBase = m.M
	v235 = m.ExcPending
	if v235 != 0 {
		goto L66
	} else {
		goto L75
	}
L75:
	;
	F_errmsg(m, int32(_a_F_TransactionIdIsInProgress_1), int32(0))
	mBase = m.M
	v239 = m.ExcPending
	if v239 != 0 {
		goto L66
	} else {
		goto L76
	}
L76:
	;
	F_errfinish(m, int32(_a_F_TransactionIdIsInProgress_2), int32(1456), int32(_a_F_TransactionIdIsInProgress_3))
	mBase = m.M
	v244 = m.ExcPending
	if v244 != 0 {
		goto L66
	} else {
		goto L77
	}
L77:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L78:
	;
	v249 = *(*int32)(unsafe.Add(mBase, _c_F_TransactionIdIsInProgress[16]))
	v250 = *(*int32)(unsafe.Add(mBase, uint32(v249)+32))
	v258 = int32(0)
	v261 = v2
	goto L81
L79:
	;
	v356 = v2
	goto L80
L80:
	;
	v365 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_TransactionIdIsInProgress[9])))
	if v365 == int32(1) {
		goto L104
	} else {
		goto L105
	}
L81:
	;
	if v258 == v250 {
		v341 = v261
		goto L83
	} else {
		goto L84
	}
L82:
	;
	v356 = v341
	goto L80
L83:
	;
	v349 = v258 + int32(1)
	if v349 != v245 {
		v258 = v349
		v261 = v341
		goto L81
	} else {
		goto L101
	}
L84:
	;
	v270 = v258 << (uint(int32(2)) % 32)
	v272 = *(*int32)(unsafe.Add(mBase, _c_F_TransactionIdIsInProgress[11]))
	v274 = *(*int32)(unsafe.Add(mBase, uint32(v270+v272)))
	if v274 == int32(0) {
		v341 = v261
		goto L83
	} else {
		goto L85
	}
L85:
	;
	if l0 == v274 {
		goto L86
	} else {
		goto L87
	}
L86:
	;
	goto L1
L87:
	;
	goto L88
L88:
	;
	if base.B2i32(base.Ui32(l0) < base.Ui32(int32(3)))|base.B2i32(base.Ui32(v274) < base.Ui32(int32(3))) == int32(0) {
		goto L90
	} else {
		goto L91
	}
L89:
	;
	v289 = v205 + v258<<(uint(int32(1))%32)
	v290 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v289))))
	v291 = int32(0)
	v294 = base.AtomicRmwOr32(m, v291, int32(_a_F_TransactionIdIsInProgress_4), v291)
	v296 = *(*int32)(unsafe.Add(mBase, _c_F_TransactionIdIsInProgress[17]))
	v298 = *(*int32)(unsafe.Add(mBase, uint32(v31+int32(36)+v270)))
	v305 = v290
	goto L96
L90:
	;
	if int32(0) <= l0-v274 {
		goto L89
	} else {
		goto L93
	}
L91:
	;
	goto L92
L92:
	;
	if base.Ui32(l0) < base.Ui32(v274) {
		v341 = v261
		goto L83
	} else {
		goto L94
	}
L93:
	;
	v341 = v261
	goto L83
L94:
	;
	goto L89
L95:
	;
	v325 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v289)+1)))
	if v325 != int32(1) {
		v341 = v261
		goto L83
	} else {
		goto L100
	}
L96:
	;
	if v305 <= int32(0) {
		goto L95
	} else {
		goto L98
	}
L97:
	;
	goto L1
L98:
	;
	v319 = v305 - int32(1)
	v323 = *(*int32)(unsafe.Add(mBase, uint32(v296+v298*int32(768)+int32(60)+v319<<(uint(int32(2))%32))))
	if v323 != l0 {
		v305 = v319
		goto L96
	} else {
		goto L99
	}
L99:
	;
	goto L97
L100:
	;
	v329 = *(*int32)(unsafe.Add(mBase, _c_F_TransactionIdIsInProgress[8]))
	*(*int32)(unsafe.Add(mBase, uint32(v329+v261<<(uint(int32(2))%32)))) = v274
	v341 = v261 + int32(1)
	goto L83
L101:
	;
	goto L82
L102:
	;
	v551 = *(*int32)(unsafe.Add(mBase, _c_F_TransactionIdIsInProgress[1]))
	F_LWLockRelease(m, v551+int32(512))
	mBase = m.M
	v555 = m.ExcPending
	if v555 != 0 {
		goto L66
	} else {
		goto L148
	}
L103:
	;
	if v375 == int32(0) {
		v543 = v356
		goto L102
	} else {
		goto L107
	}
L104:
	;
	v370 = *(*int32)(unsafe.Add(mBase, _c_F_TransactionIdIsInProgress[12]))
	v371 = *(*int32)(unsafe.Add(mBase, uint32(v370)+308))
	v373 = base.B2i32(v371 != int32(2))
	*(*uint8)(unsafe.Add(mBase, _c_F_TransactionIdIsInProgress[9])) = uint8(v373)
	v375 = v373
	goto L106
L105:
	;
	v375 = int32(0)
	goto L106
L106:
	;
	goto L103
L107:
	;
	v379 = *(*int32)(unsafe.Add(mBase, _c_F_TransactionIdIsInProgress[3]))
	v380 = *(*int32)(unsafe.Add(mBase, uint32(v379)+16))
	v381 = *(*int32)(unsafe.Add(mBase, uint32(v379)+20))
	v382 = int32(0)
	v385 = base.AtomicRmwOr32(m, v382, int32(_a_F_TransactionIdIsInProgress_4), v382)
	v387 = v381 - int32(1)
	if v387 < v380 {
		goto L108
	} else {
		goto L109
	}
L108:
	;
	v450 = int32(3)
	v453 = *(*int32)(unsafe.Add(mBase, _c_F_TransactionIdIsInProgress[3]))
	v454 = *(*int32)(unsafe.Add(mBase, uint32(v453)+24))
	if base.B2i32(base.Ui32(l0) < base.Ui32(v450))|base.B2i32(base.Ui32(v454) < base.Ui32(v450)) == int32(0) {
		goto L128
	} else {
		goto L129
	}
L109:
	;
	v390 = *(*int32)(unsafe.Add(mBase, _c_F_TransactionIdIsInProgress[18]))
	v394 = v380
	v395 = v387
	goto L110
L110:
	;
	v405 = v394 + v395
	v406 = int32(2)
	v407 = base.I32_div_s(v405, v406)
	v411 = *(*int32)(unsafe.Add(mBase, uint32(v390+v407<<(uint(v406)%32))))
	if v411 != l0 {
		goto L112
	} else {
		goto L113
	}
L111:
	;
	if v405 < int32(-1) {
		goto L108
	} else {
		goto L125
	}
L112:
	;
	if base.B2i32(base.B2i32(base.Ui32(l0) < base.Ui32(int32(3))) == int32(0))&base.B2i32(base.Ui32(int32(2)) < base.Ui32(v411)) != 0 {
		goto L115
	} else {
		goto L116
	}
L113:
	;
	goto L114
L114:
	;
	goto L111
L115:
	;
	v424 = int32(base.Ui32(l0-v411) >> (uint(int32(31)) % 32))
	goto L117
L116:
	;
	v424 = base.B2i32(base.Ui32(l0) < base.Ui32(v411))
	goto L117
L117:
	;
	if v424 != 0 {
		goto L118
	} else {
		goto L119
	}
L118:
	;
	v425 = v394
	goto L120
L119:
	;
	v425 = v407 + int32(1)
	goto L120
L120:
	;
	if v424 != 0 {
		goto L121
	} else {
		goto L122
	}
L121:
	;
	v428 = v407 - int32(1)
	goto L123
L122:
	;
	v428 = v395
	goto L123
L123:
	;
	if v425 <= v428 {
		v394 = v425
		v395 = v428
		goto L110
	} else {
		goto L124
	}
L124:
	;
	goto L108
L125:
	;
	v433 = *(*int32)(unsafe.Add(mBase, _c_F_TransactionIdIsInProgress[19]))
	v435 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v433+v407))))
	if v435 != int32(1) {
		goto L108
	} else {
		goto L126
	}
L126:
	;
	goto L1
L127:
	;
	v465 = *(*int32)(unsafe.Add(mBase, _c_F_TransactionIdIsInProgress[8]))
	v466 = int32(0)
	v468 = *(*int32)(unsafe.Add(mBase, _c_F_TransactionIdIsInProgress[3]))
	v469 = *(*int32)(unsafe.Add(mBase, uint32(v468)+20))
	v470 = *(*int32)(unsafe.Add(mBase, uint32(v468)+16))
	v474 = base.AtomicRmwOr32(m, v466, int32(_a_F_TransactionIdIsInProgress_4), v466)
	if v469 <= v470 {
		v531 = v466
		goto L133
	} else {
		goto L134
	}
L128:
	;
	if l0-v454 <= int32(0) {
		goto L127
	} else {
		goto L131
	}
L129:
	;
	goto L130
L130:
	;
	if base.Ui32(v454) < base.Ui32(l0) {
		v543 = v356
		goto L102
	} else {
		goto L132
	}
L131:
	;
	v543 = v356
	goto L102
L132:
	;
	goto L127
L133:
	;
	v543 = v531
	goto L102
L134:
	;
	v477 = *(*int32)(unsafe.Add(mBase, _c_F_TransactionIdIsInProgress[19]))
	v481 = v470
	v482 = v477
	v485 = v466
	goto L135
L135:
	;
	v493 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v481+v482))))
	if v493 == int32(1) {
		goto L137
	} else {
		goto L138
	}
L136:
	;
	v531 = v522
	goto L133
L137:
	;
	v497 = *(*int32)(unsafe.Add(mBase, _c_F_TransactionIdIsInProgress[18]))
	v501 = *(*int32)(unsafe.Add(mBase, uint32(v497+v481<<(uint(int32(2))%32))))
	if l0 == int32(0) {
		goto L140
	} else {
		goto L141
	}
L138:
	;
	v521 = v482
	v522 = v485
	goto L139
L139:
	;
	v524 = v481 + int32(1)
	if v524 != v469 {
		v481 = v524
		v482 = v521
		v485 = v522
		goto L135
	} else {
		goto L147
	}
L140:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v465+v485<<(uint(int32(2))%32)))) = v501
	v518 = *(*int32)(unsafe.Add(mBase, _c_F_TransactionIdIsInProgress[19]))
	v521 = v518
	v522 = v485 + int32(1)
	goto L139
L141:
	;
	if base.B2i32(base.Ui32(l0) < base.Ui32(int32(3)))|base.B2i32(base.Ui32(v501) < base.Ui32(int32(3))) == int32(0) {
		goto L142
	} else {
		goto L143
	}
L142:
	;
	if v501-l0 < int32(0) {
		goto L140
	} else {
		goto L145
	}
L143:
	;
	goto L144
L144:
	;
	if base.Ui32(l0) <= base.Ui32(v501) {
		v531 = v485
		goto L133
	} else {
		goto L146
	}
L145:
	;
	v531 = v485
	goto L133
L146:
	;
	goto L140
L147:
	;
	goto L136
L148:
	;
	if v543 != 0 {
		goto L149
	} else {
		goto L150
	}
L149:
	;
	v556 = F_TransactionIdDidAbort(m, l0)
	mBase = m.M
	v557 = m.ExcPending
	if v557 != 0 {
		goto L66
	} else {
		goto L153
	}
L150:
	;
	goto L151
L151:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_TransactionIdIsInProgress[2])) = l0
	goto L2
L152:
	;
	goto L151
L153:
	;
	if v556 != 0 {
		goto L152
	} else {
		goto L154
	}
L154:
	;
	v558 = F_SubTransGetTopmostTransaction(m, l0)
	mBase = m.M
	v559 = m.ExcPending
	if v559 != 0 {
		goto L66
	} else {
		goto L155
	}
L155:
	;
	if v558 == l0 {
		goto L152
	} else {
		goto L156
	}
L156:
	;
	v563 = *(*int32)(unsafe.Add(mBase, _c_F_TransactionIdIsInProgress[8]))
	v565 = int32(0)
	goto L157
L157:
	;
	v579 = *(*int32)(unsafe.Add(mBase, uint32(v563+v565<<(uint(int32(2))%32))))
	v580 = base.B2i32(v558 == v579)
	if v580 == int32(0) {
		goto L159
	} else {
		goto L160
	}
L158:
	;
	if v580 == int32(0) {
		goto L152
	} else {
		goto L163
	}
L159:
	;
	v584 = v565 + int32(1)
	if v584 != v543 {
		v565 = v584
		goto L157
	} else {
		goto L162
	}
L160:
	;
	goto L161
L161:
	;
	goto L158
L162:
	;
	goto L161
L163:
	;
	return int32(1)
L164:
	;
	return int32(1)
}
func F_TransactionIdSetTreeStatus(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int64) {
	mBase := m.M
	_ = mBase
	var v6 int32
	_ = v6
	var v18 int32
	_ = v18
	var v19 int64
	_ = v19
	var v30 int32
	_ = v30
	var v41 int32
	_ = v41
	var v46 int32
	_ = v46
	var v56 int32
	_ = v56
	var v82 int32
	_ = v82
	var v83 int32
	_ = v83
	var v86 int32
	_ = v86
	var v87 int32
	_ = v87
	var v90 int32
	_ = v90
	var v94 int32
	_ = v94
	var v95 int32
	_ = v95
	var v97 int32
	_ = v97
	var v100 int32
	_ = v100
	var v102 int32
	_ = v102
	var v110 int32
	_ = v110
	var v111 int32
	_ = v111
	var v112 int32
	_ = v112
	var v115 int32
	_ = v115
	var v116 int32
	_ = v116
	var v118 int32
	_ = v118
	var v119 int32
	_ = v119
	var v121 int32
	_ = v121
	var v123 int32
	_ = v123
	var v126 int32
	_ = v126
	var v127 int32
	_ = v127
	var v128 int32
	_ = v128
	var v133 int32
	_ = v133
	var v134 int32
	_ = v134
	var v135 int32
	_ = v135
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
	var v151 int32
	_ = v151
	var v164 int32
	_ = v164
	var v166 int32
	_ = v166
	var v167 int32
	_ = v167
	var v169 int32
	_ = v169
	var v171 int32
	_ = v171
	var v176 int32
	_ = v176
	var v178 int32
	_ = v178
	var v182 int32
	_ = v182
	var v198 int32
	_ = v198
	var v199 int32
	_ = v199
	var v203 int64
	_ = v203
	var v204 int64
	_ = v204
	var v208 int32
	_ = v208
	var v210 int32
	_ = v210
	var v214 int32
	_ = v214
	var v217 int32
	_ = v217
	var v233 int32
	_ = v233
	var v235 int32
	_ = v235
	var v238 int32
	_ = v238
	var v240 int32
	_ = v240
	var v241 int32
	_ = v241
	var v248 int32
	_ = v248
	var v261 int32
	_ = v261
	var v263 int32
	_ = v263
	var v264 int32
	_ = v264
	var v268 int32
	_ = v268
	var v272 int32
	_ = v272
	var v274 int32
	_ = v274
	var v278 int32
	_ = v278
	var v279 int32
	_ = v279
	var v280 int64
	_ = v280
	var v282 int64
	_ = v282
	var v283 int64
	_ = v283
	var v287 int32
	_ = v287
	var v289 int32
	_ = v289
	var v290 int32
	_ = v290
	var v291 int32
	_ = v291
	var v293 int32
	_ = v293
	var v297 int32
	_ = v297
	var v299 int32
	_ = v299
	var v300 int64
	_ = v300
	var v313 int32
	_ = v313
	var v314 int32
	_ = v314
	var v317 int32
	_ = v317
	var v318 int64
	_ = v318
	var v321 int32
	_ = v321
	var v322 int32
	_ = v322
	var v324 int64
	_ = v324
	var v325 int64
	_ = v325
	var v329 int32
	_ = v329
	var v332 int32
	_ = v332
	var v334 int32
	_ = v334
	var v335 int32
	_ = v335
	var v336 int64
	_ = v336
	var v337 int64
	_ = v337
	var v338 int32
	_ = v338
	var v340 int64
	_ = v340
	var v341 int64
	_ = v341
	var v342 int32
	_ = v342
	var v343 int32
	_ = v343
	var v346 int32
	_ = v346
	var v347 int64
	_ = v347
	var v349 int32
	_ = v349
	var v350 int32
	_ = v350
	var v354 int32
	_ = v354
	var v370 int32
	_ = v370
	var v373 int32
	_ = v373
	var v390 int32
	_ = v390
	var v391 int32
	_ = v391
	var v394 int32
	_ = v394
	var v395 int32
	_ = v395
	var v398 int32
	_ = v398
	var v401 int32
	_ = v401
	var v405 int32
	_ = v405
	var v407 int32
	_ = v407
	var v409 int32
	_ = v409
	var v414 int32
	_ = v414
	var v433 int32
	_ = v433
	var v434 int32
	_ = v434
	var v452 int32
	_ = v452
	var v454 int32
	_ = v454
	var v473 int32
	_ = v473
	var v474 int32
	_ = v474
	var v481 int32
	_ = v481
	var v482 int32
	_ = v482
	var v492 int32
	_ = v492
	var v500 int64
	_ = v500
	var v503 int32
	_ = v503
	var v505 int32
	_ = v505
	var v506 int32
	_ = v506
	var v509 int32
	_ = v509
	var v513 int32
	_ = v513
	var v527 int32
	_ = v527
	var v530 int64
	_ = v530
	var v532 int32
	_ = v532
	var v535 int32
	_ = v535
	var v539 int32
	_ = v539
	var v540 int32
	_ = v540
	var v541 int64
	_ = v541
	var v543 int32
	_ = v543
	var v544 int32
	_ = v544
	var v547 int32
	_ = v547
	var v548 int32
	_ = v548
	var v551 int32
	_ = v551
	var v553 int32
	_ = v553
	var v554 int32
	_ = v554
	var v561 int32
	_ = v561
	var v563 int32
	_ = v563
	var v582 int32
	_ = v582
	var v583 int32
	_ = v583
	var v585 int32
	_ = v585
	var v586 int32
	_ = v586
	var v589 int32
	_ = v589
	var v591 int32
	_ = v591
	var v592 int32
	_ = v592
	var v594 int32
	_ = v594
	var v596 int32
	_ = v596
	var v601 int32
	_ = v601
	var v602 int32
	_ = v602
	var v607 int32
	_ = v607
	var v621 int64
	_ = v621
	var v624 int32
	_ = v624
	var v626 int32
	_ = v626
	var v627 int32
	_ = v627
	var v630 int32
	_ = v630
	var v634 int32
	_ = v634
	var v648 int32
	_ = v648
	var v651 int64
	_ = v651
	var v653 int32
	_ = v653
	var v656 int32
	_ = v656
	var v658 int32
	_ = v658
	var v661 int32
	_ = v661
	var v662 int64
	_ = v662
	var v664 int32
	_ = v664
	var v665 int32
	_ = v665
	var v668 int32
	_ = v668
	var v669 int32
	_ = v669
	var v672 int32
	_ = v672
	var v674 int32
	_ = v674
	var v675 int32
	_ = v675
	var v681 int32
	_ = v681
	var v683 int32
	_ = v683
	v6 = int32(0)
	v18 = int32(base.Ui32(l0) >> (uint(int32(15)) % 32))
	v19 = base.I64_extend_i32_u(v18)
	if l1 <= v6 {
		v56 = v6
		goto L3
	} else {
		goto L4
	}
L1:
	;
	v473 = l1 - v56
	v474 = int32(0)
	if base.B2i32(l3 != int32(1))|base.B2i32(v473 <= v474) == v474 {
		goto L91
	} else {
		goto L92
	}
L2:
	;
	v82 = *(*int32)(unsafe.Add(mBase, _c_F_TransactionIdSetTreeStatus[0]))
	v83 = *(*int32)(unsafe.Add(mBase, uint32(v82)+28))
	v86 = int32(*(*uint16)(unsafe.Add(mBase, _c_F_TransactionIdSetTreeStatus[1])))
	v87 = base.I32_rem_u_s(base.I32_wrap_i64(v19), v86)
	v90 = v83 + v87<<(uint(int32(7))%32)
	if int32(5) < l1 {
		goto L12
	} else {
		goto L13
	}
L3:
	;
	if l1 != v56 {
		goto L1
	} else {
		goto L9
	}
L4:
	;
	v30 = v6
	goto L5
L5:
	;
	v41 = *(*int32)(unsafe.Add(mBase, uint32(l2+v30<<(uint(int32(2))%32))))
	if int32(base.Ui32(v41)>>(uint(int32(15))%32)) != v18 {
		v56 = v30
		goto L3
	} else {
		goto L7
	}
L6:
	;
	goto L2
L7:
	;
	v46 = v30 + int32(1)
	if v46 != l1 {
		v30 = v46
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
	F_TransactionIdSetPageStatusInternal(m, l0, l1, l2, l3, l4, v19)
	mBase = m.M
	v452 = m.ExcPending
	if v452 != 0 {
		goto L38
	} else {
		goto L89
	}
L12:
	;
	v433 = F_LWLockAcquire(m, v90, int32(0))
	mBase = m.M
	v434 = m.ExcPending
	if v434 != 0 {
		goto L38
	} else {
		goto L88
	}
L13:
	;
	v94 = *(*int32)(unsafe.Add(mBase, _c_F_TransactionIdSetTreeStatus[2]))
	v95 = *(*int32)(unsafe.Add(mBase, uint32(v94)+48))
	if l0 != v95 {
		goto L12
	} else {
		goto L14
	}
L14:
	;
	v97 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v94)+56)))
	if l1 != v97 {
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
	v100 = v94 + int32(60)
	v102 = l1 << (uint(int32(2)) % 32)
	if base.Ui32(int32(4)) <= base.Ui32(v102) {
		goto L22
	} else {
		goto L23
	}
L17:
	;
	goto L18
L18:
	;
	v166 = F_LWLockConditionalAcquire(m, v90, int32(0))
	mBase = m.M
	v167 = m.ExcPending
	if v167 != 0 {
		goto L38
	} else {
		goto L39
	}
L19:
	;
	if v164 != 0 {
		goto L12
	} else {
		goto L37
	}
L20:
	;
	v164 = int32(0)
	goto L19
L21:
	;
	v138 = v133
	v139 = v134
	v140 = v135
	goto L31
L22:
	;
	if (l2|v100)&int32(3) != 0 {
		v133 = l2
		v134 = v100
		v135 = v102
		goto L21
	} else {
		goto L25
	}
L23:
	;
	v126 = l2
	v127 = v100
	v128 = v102
	goto L24
L24:
	;
	if v128 == int32(0) {
		goto L20
	} else {
		goto L30
	}
L25:
	;
	v110 = l2
	v111 = v100
	v112 = v102
	goto L26
L26:
	;
	v115 = *(*int32)(unsafe.Add(mBase, uint32(v110)))
	v116 = *(*int32)(unsafe.Add(mBase, uint32(v111)))
	if v115 != v116 {
		v133 = v110
		v134 = v111
		v135 = v112
		goto L21
	} else {
		goto L28
	}
L27:
	;
	v126 = v121
	v127 = v119
	v128 = v123
	goto L24
L28:
	;
	v118 = int32(4)
	v119 = v111 + v118
	v121 = v110 + v118
	v123 = v112 - v118
	if base.Ui32(int32(3)) < base.Ui32(v123) {
		v110 = v121
		v111 = v119
		v112 = v123
		goto L26
	} else {
		goto L29
	}
L29:
	;
	goto L27
L30:
	;
	v133 = v126
	v134 = v127
	v135 = v128
	goto L21
L31:
	;
	v143 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v138))))
	v144 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v139))))
	if v143 == v144 {
		goto L33
	} else {
		goto L34
	}
L32:
	;
	v164 = v143 - v144
	goto L19
L33:
	;
	v146 = int32(1)
	v151 = v140 - v146
	if v151 != 0 {
		v138 = v138 + v146
		v139 = v139 + v146
		v140 = v151
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
	if v166 != 0 {
		goto L11
	} else {
		goto L40
	}
L40:
	;
	v169 = *(*int32)(unsafe.Add(mBase, _c_F_TransactionIdSetTreeStatus[3]))
	v171 = *(*int32)(unsafe.Add(mBase, _c_F_TransactionIdSetTreeStatus[2]))
	*(*int64)(unsafe.Add(mBase, uint32(v171)+640)) = l4
	*(*int64)(unsafe.Add(mBase, uint32(v171)+632)) = v19
	*(*int32)(unsafe.Add(mBase, uint32(v171)+628)) = l3
	*(*int32)(unsafe.Add(mBase, uint32(v171)+624)) = l0
	v176 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(v171)+616)) = uint8(v176)
	v178 = *(*int32)(unsafe.Add(mBase, uint32(v169)+60))
	v182 = v178
	goto L42
L41:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v171)+620)) = int32(-1)
	v414 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v171)+616)) = uint8(v414)
	goto L12
L42:
	;
	if v182 != int32(-1) {
		goto L44
	} else {
		goto L45
	}
L43:
	;
	v278 = *(*int32)(unsafe.Add(mBase, _c_F_TransactionIdSetTreeStatus[0]))
	v279 = *(*int32)(unsafe.Add(mBase, uint32(v278)+28))
	v280 = *(*int64)(unsafe.Add(mBase, uint32(v171)+632))
	v282 = int64(*(*uint16)(unsafe.Add(mBase, _c_F_TransactionIdSetTreeStatus[1])))
	v283 = base.I64_rem_s(v280, v282)
	v287 = v279 + base.I32_wrap_i64(v283)<<(uint(int32(7))%32)
	v289 = F_LWLockAcquire(m, v287, int32(0))
	mBase = m.M
	v290 = m.ExcPending
	if v290 != 0 {
		goto L38
	} else {
		goto L59
	}
L44:
	;
	v198 = *(*int32)(unsafe.Add(mBase, _c_F_TransactionIdSetTreeStatus[3]))
	v199 = *(*int32)(unsafe.Add(mBase, uint32(v198)))
	v203 = *(*int64)(unsafe.Add(mBase, uint32(v199+v182*int32(768))+632))
	v204 = *(*int64)(unsafe.Add(mBase, uint32(v171)+632))
	if v203 != v204 {
		goto L41
	} else {
		goto L47
	}
L45:
	;
	goto L46
L46:
	;
	v268 = int32(-1)
	*(*int32)(unsafe.Add(mBase, uint32(v171)+620)) = v268
	v272 = *(*int32)(unsafe.Add(mBase, _c_F_TransactionIdSetTreeStatus[4]))
	v274 = base.AtomicRmwCmpxchg32(m, v169, int32(60), v268, v272)
	if v274 != v268 {
		v182 = v274
		goto L42
	} else {
		goto L58
	}
L47:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v171)+620)) = v182
	v208 = *(*int32)(unsafe.Add(mBase, _c_F_TransactionIdSetTreeStatus[4]))
	v210 = base.AtomicRmwCmpxchg32(m, v169, int32(60), v182, v208)
	if v182 != v210 {
		v182 = v210
		goto L42
	} else {
		goto L48
	}
L48:
	;
	v214 = *(*int32)(unsafe.Add(mBase, _c_F_TransactionIdSetTreeStatus[5]))
	*(*int32)(unsafe.Add(mBase, uint32(v214))) = int32(134217789)
	v217 = int32(0)
	goto L49
L49:
	;
	v233 = *(*int32)(unsafe.Add(mBase, uint32(v171)+332))
	F_PGSemaphoreLock(m, v233)
	mBase = m.M
	v235 = m.ExcPending
	if v235 != 0 {
		goto L38
	} else {
		goto L51
	}
L50:
	;
	v240 = *(*int32)(unsafe.Add(mBase, _c_F_TransactionIdSetTreeStatus[5]))
	v241 = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v240))) = v241
	if v217 <= v241 {
		goto L10
	} else {
		goto L53
	}
L51:
	;
	v238 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v171)+616)))
	if v238 != 0 {
		v217 = v217 + int32(1)
		goto L49
	} else {
		goto L52
	}
L52:
	;
	goto L50
L53:
	;
	v248 = v217
	goto L54
L54:
	;
	v261 = *(*int32)(unsafe.Add(mBase, uint32(v171)+332))
	F_PGSemaphoreUnlock(m, v261)
	mBase = m.M
	v263 = m.ExcPending
	if v263 != 0 {
		goto L38
	} else {
		goto L56
	}
L55:
	;
	goto L10
L56:
	;
	v264 = int32(1)
	if base.Ui32(v264) < base.Ui32(v248) {
		v248 = v248 - v264
		goto L54
	} else {
		goto L57
	}
L57:
	;
	goto L55
L58:
	;
	goto L43
L59:
	;
	v291 = int32(-1)
	v293 = base.AtomicRmwXchg32(m, v169, int32(60), v291)
	if v293 != v291 {
		goto L60
	} else {
		goto L61
	}
L60:
	;
	v297 = v287
	v299 = v293
	v300 = v280
	goto L63
L61:
	;
	v354 = v287
	goto L62
L62:
	;
	if v354 != 0 {
		goto L76
	} else {
		goto L77
	}
L63:
	;
	v313 = *(*int32)(unsafe.Add(mBase, _c_F_TransactionIdSetTreeStatus[3]))
	v314 = *(*int32)(unsafe.Add(mBase, uint32(v313)))
	v317 = v314 + v299*int32(768)
	v318 = *(*int64)(unsafe.Add(mBase, uint32(v317)+632))
	if v300 == v318 {
		goto L66
	} else {
		goto L67
	}
L64:
	;
	v354 = v338
	goto L62
L65:
	;
	v342 = *(*int32)(unsafe.Add(mBase, uint32(v317)+624))
	v343 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v317)+56)))
	v346 = *(*int32)(unsafe.Add(mBase, uint32(v317)+628))
	v347 = *(*int64)(unsafe.Add(mBase, uint32(v317)+640))
	F_TransactionIdSetPageStatusInternal(m, v342, v343, v317+int32(60), v346, v347, v341)
	mBase = m.M
	v349 = m.ExcPending
	if v349 != 0 {
		goto L38
	} else {
		goto L74
	}
L66:
	;
	v338 = v297
	v340 = v300
	v341 = v300
	goto L65
L67:
	;
	goto L68
L68:
	;
	v321 = *(*int32)(unsafe.Add(mBase, _c_F_TransactionIdSetTreeStatus[0]))
	v322 = *(*int32)(unsafe.Add(mBase, uint32(v321)+28))
	v324 = int64(*(*uint16)(unsafe.Add(mBase, _c_F_TransactionIdSetTreeStatus[1])))
	v325 = base.I64_rem_s(v318, v324)
	v329 = v322 + base.I32_wrap_i64(v325)<<(uint(int32(7))%32)
	if v329 == v297 {
		goto L69
	} else {
		goto L70
	}
L69:
	;
	v337 = v318
	goto L71
L70:
	;
	F_LWLockRelease(m, v297)
	mBase = m.M
	v332 = m.ExcPending
	if v332 != 0 {
		goto L38
	} else {
		goto L72
	}
L71:
	;
	v338 = v329
	v340 = v318
	v341 = v337
	goto L65
L72:
	;
	v334 = F_LWLockAcquire(m, v329, int32(0))
	mBase = m.M
	v335 = m.ExcPending
	if v335 != 0 {
		goto L38
	} else {
		goto L73
	}
L73:
	;
	v336 = *(*int64)(unsafe.Add(mBase, uint32(v317)+632))
	v337 = v336
	goto L71
L74:
	;
	v350 = *(*int32)(unsafe.Add(mBase, uint32(v317)+620))
	if v350 != int32(-1) {
		v297 = v338
		v299 = v350
		v300 = v340
		goto L63
	} else {
		goto L75
	}
L75:
	;
	goto L64
L76:
	;
	F_LWLockRelease(m, v354)
	mBase = m.M
	v370 = m.ExcPending
	if v370 != 0 {
		goto L38
	} else {
		goto L79
	}
L77:
	;
	goto L78
L78:
	;
	if v293 == int32(-1) {
		goto L10
	} else {
		goto L80
	}
L79:
	;
	goto L78
L80:
	;
	v373 = v293
	goto L81
L81:
	;
	v390 = *(*int32)(unsafe.Add(mBase, _c_F_TransactionIdSetTreeStatus[3]))
	v391 = *(*int32)(unsafe.Add(mBase, uint32(v390)))
	v394 = v391 + v373*int32(768)
	v395 = *(*int32)(unsafe.Add(mBase, uint32(v394)+620))
	*(*int32)(unsafe.Add(mBase, uint32(v394)+620)) = int32(-1)
	v398 = int32(0)
	v401 = base.AtomicRmwOr32(m, v398, int32(_a_F_TransactionIdSetTreeStatus_0), v398)
	*(*uint8)(unsafe.Add(mBase, uint32(v394)+616)) = uint8(v398)
	v405 = *(*int32)(unsafe.Add(mBase, _c_F_TransactionIdSetTreeStatus[2]))
	if v405 != v394 {
		goto L83
	} else {
		goto L84
	}
L82:
	;
	goto L10
L83:
	;
	v407 = *(*int32)(unsafe.Add(mBase, uint32(v394)+332))
	F_PGSemaphoreUnlock(m, v407)
	mBase = m.M
	v409 = m.ExcPending
	if v409 != 0 {
		goto L38
	} else {
		goto L86
	}
L84:
	;
	goto L85
L85:
	;
	if v395 != int32(-1) {
		v373 = v395
		goto L81
	} else {
		goto L87
	}
L86:
	;
	goto L85
L87:
	;
	goto L82
L88:
	;
	goto L11
L89:
	;
	F_LWLockRelease(m, v90)
	mBase = m.M
	v454 = m.ExcPending
	if v454 != 0 {
		goto L38
	} else {
		goto L90
	}
L90:
	;
	goto L10
L91:
	;
	v481 = l2 + v56<<(uint(int32(2))%32)
	v482 = *(*int32)(unsafe.Add(mBase, uint32(v481)))
	v492 = v6
	v500 = base.I64_extend_i32_u(int32(base.Ui32(v482) >> (uint(int32(15)) % 32)))
	goto L94
L92:
	;
	goto L93
L93:
	;
	v582 = *(*int32)(unsafe.Add(mBase, _c_F_TransactionIdSetTreeStatus[0]))
	v583 = *(*int32)(unsafe.Add(mBase, uint32(v582)+28))
	v585 = int32(*(*uint16)(unsafe.Add(mBase, _c_F_TransactionIdSetTreeStatus[1])))
	v586 = base.I32_rem_u_s(v18, v585)
	v589 = v583 + v586<<(uint(int32(7))%32)
	v591 = F_LWLockAcquire(m, v589, int32(0))
	mBase = m.M
	v592 = m.ExcPending
	if v592 != 0 {
		goto L38
	} else {
		goto L110
	}
L94:
	;
	v503 = v492 + int32(1)
	if v503 < v473 {
		goto L96
	} else {
		goto L97
	}
L95:
	;
	goto L93
L96:
	;
	v505 = v473
	goto L98
L97:
	;
	v505 = v503
	goto L98
L98:
	;
	v506 = v505 - v492
	v509 = v492
	v513 = int32(0)
	goto L100
L99:
	;
	v543 = *(*int32)(unsafe.Add(mBase, _c_F_TransactionIdSetTreeStatus[0]))
	v544 = *(*int32)(unsafe.Add(mBase, uint32(v543)+28))
	v547 = int32(*(*uint16)(unsafe.Add(mBase, _c_F_TransactionIdSetTreeStatus[1])))
	v548 = base.I32_rem_u_s(base.I32_wrap_i64(v500), v547)
	v551 = v544 + v548<<(uint(int32(7))%32)
	v553 = F_LWLockAcquire(m, v551, int32(0))
	mBase = m.M
	v554 = m.ExcPending
	if v554 != 0 {
		goto L38
	} else {
		goto L106
	}
L100:
	;
	v527 = *(*int32)(unsafe.Add(mBase, uint32(v481+v509<<(uint(int32(2))%32))))
	v530 = base.I64_extend_i32_u(int32(base.Ui32(v527) >> (uint(int32(15)) % 32)))
	if v530 != v500 {
		goto L102
	} else {
		goto L103
	}
L101:
	;
	v539 = v505
	v540 = v506
	v541 = v500
	goto L99
L102:
	;
	v539 = v509
	v540 = v513
	v541 = v530
	goto L99
L103:
	;
	goto L104
L104:
	;
	v532 = int32(1)
	v535 = v513 + v532
	if v535 != v506 {
		v509 = v509 + v532
		v513 = v535
		goto L100
	} else {
		goto L105
	}
L105:
	;
	goto L101
L106:
	;
	F_TransactionIdSetPageStatusInternal(m, int32(0), v540, v481+v492<<(uint(int32(2))%32), int32(3), l4, v500)
	mBase = m.M
	v561 = m.ExcPending
	if v561 != 0 {
		goto L38
	} else {
		goto L107
	}
L107:
	;
	F_LWLockRelease(m, v551)
	mBase = m.M
	v563 = m.ExcPending
	if v563 != 0 {
		goto L38
	} else {
		goto L108
	}
L108:
	;
	if v539 < v473 {
		v492 = v539
		v500 = v541
		goto L94
	} else {
		goto L109
	}
L109:
	;
	goto L95
L110:
	;
	F_TransactionIdSetPageStatusInternal(m, l0, v56, l2, l3, l4, v19)
	mBase = m.M
	v594 = m.ExcPending
	if v594 != 0 {
		goto L38
	} else {
		goto L111
	}
L111:
	;
	F_LWLockRelease(m, v589)
	mBase = m.M
	v596 = m.ExcPending
	if v596 != 0 {
		goto L38
	} else {
		goto L112
	}
L112:
	;
	if int32(0) < v473 {
		goto L113
	} else {
		goto L114
	}
L113:
	;
	v601 = l2 + v56<<(uint(int32(2))%32)
	v602 = *(*int32)(unsafe.Add(mBase, uint32(v601)))
	v607 = int32(0)
	v621 = base.I64_extend_i32_u(int32(base.Ui32(v602) >> (uint(int32(15)) % 32)))
	goto L116
L114:
	;
	goto L115
L115:
	;
	return
L116:
	;
	v624 = v607 + int32(1)
	if v624 < v473 {
		goto L118
	} else {
		goto L119
	}
L117:
	;
	goto L115
L118:
	;
	v626 = v473
	goto L120
L119:
	;
	v626 = v624
	goto L120
L120:
	;
	v627 = v626 - v607
	v630 = v607
	v634 = int32(0)
	goto L122
L121:
	;
	v664 = *(*int32)(unsafe.Add(mBase, _c_F_TransactionIdSetTreeStatus[0]))
	v665 = *(*int32)(unsafe.Add(mBase, uint32(v664)+28))
	v668 = int32(*(*uint16)(unsafe.Add(mBase, _c_F_TransactionIdSetTreeStatus[1])))
	v669 = base.I32_rem_u_s(base.I32_wrap_i64(v621), v668)
	v672 = v665 + v669<<(uint(int32(7))%32)
	v674 = F_LWLockAcquire(m, v672, int32(0))
	mBase = m.M
	v675 = m.ExcPending
	if v675 != 0 {
		goto L38
	} else {
		goto L128
	}
L122:
	;
	v648 = *(*int32)(unsafe.Add(mBase, uint32(v601+v630<<(uint(int32(2))%32))))
	v651 = base.I64_extend_i32_u(int32(base.Ui32(v648) >> (uint(int32(15)) % 32)))
	if v651 != v621 {
		goto L124
	} else {
		goto L125
	}
L123:
	;
	v658 = v626
	v661 = v627
	v662 = v621
	goto L121
L124:
	;
	v658 = v630
	v661 = v634
	v662 = v651
	goto L121
L125:
	;
	goto L126
L126:
	;
	v653 = int32(1)
	v656 = v634 + v653
	if v656 != v627 {
		v630 = v630 + v653
		v634 = v656
		goto L122
	} else {
		goto L127
	}
L127:
	;
	goto L123
L128:
	;
	F_TransactionIdSetPageStatusInternal(m, int32(0), v661, v601+v607<<(uint(int32(2))%32), l3, l4, v621)
	mBase = m.M
	v681 = m.ExcPending
	if v681 != 0 {
		goto L38
	} else {
		goto L129
	}
L129:
	;
	F_LWLockRelease(m, v672)
	mBase = m.M
	v683 = m.ExcPending
	if v683 != 0 {
		goto L38
	} else {
		goto L130
	}
L130:
	;
	if v658 < v473 {
		v607 = v658
		v621 = v662
		goto L116
	} else {
		goto L131
	}
L131:
	;
	goto L117
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
