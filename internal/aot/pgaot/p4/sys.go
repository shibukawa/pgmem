package p4

import (
	base "github.com/shibukawa/pgmem/internal/aot/pgaot/base"
	"unsafe"
)

func F_ReleaseSysCache(m *base.Module, l0 int32) {
	var v3 int32
	_ = v3
	F_ReleaseCatCache(m, l0)
	v3 = m.ExcPending
	if v3 != 0 {
		return
	} else {
		return
	}
}
func F_SearchSysCacheCopyAttName(m *base.Module, l0 int32, l1 int32) int32 {
	var v4 int32
	_ = v4
	var v7 int32
	_ = v7
	v4 = Fn13837(m, l0, l1, int32(_a_F_SearchSysCacheCopyAttName_0))
	v7 = m.ExcPending
	if v7 != 0 {
		return int32(0)
	} else {
		return v4
	}
}
func F_SysCacheGetAttrNotNull(m *base.Module, l0 int32, l1 int32, l2 int32) int32 {
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
	var v16 int32
	_ = v16
	var v22 int32
	_ = v22
	var v27 int32
	_ = v27
	var v28 int32
	_ = v28
	var v29 int32
	_ = v29
	var v34 int32
	_ = v34
	var v35 int32
	_ = v35
	var v36 int32
	_ = v36
	var v49 int32
	_ = v49
	var v54 int32
	_ = v54
	v6 = m.G0
	v8 = v6 - int32(16)
	m.G0 = v8
	v12 = F_SysCacheGetAttr(m, l0, l1, l2, v8+int32(15))
	mBase = m.M
	v15 = m.ExcPending
	if v15 != 0 {
		return int32(0)
	} else {
		v16 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v8)+15)))
		if v16 == int32(1) {
			F_errstart_cold(m, int32(21), int32(0))
			mBase = m.M
			v22 = m.ExcPending
			if v22 != 0 {
				return int32(0)
			} else {
				v27 = *(*int32)(unsafe.Add(mBase, uint32(l0<<(uint(int32(5))%32))+uint32(_c_F_SysCacheGetAttrNotNull[0])))
				v28 = F_get_rel_name(m, v27)
				mBase = m.M
				v29 = m.ExcPending
				if v29 != 0 {
					return int32(0)
				} else {
					v34 = *(*int32)(unsafe.Add(mBase, uint32(l0<<(uint(int32(2))%32))+uint32(_c_F_SysCacheGetAttrNotNull[1])))
					v35 = *(*int32)(unsafe.Add(mBase, uint32(v34)+8))
					v36 = *(*int32)(unsafe.Add(mBase, uint32(v35)))
					*(*int32)(unsafe.Add(mBase, uint32(v8))) = v28
					*(*int32)(unsafe.Add(mBase, uint32(v8)+4)) = v35 + v36<<(uint(int32(4))%32) + l2*int32(100) - int32(76)
					F_errmsg_internal(m, int32(_a_F_SysCacheGetAttrNotNull_0), v8)
					mBase = m.M
					v49 = m.ExcPending
					if v49 != 0 {
						return int32(0)
					} else {
						F_errfinish(m, int32(_a_F_SysCacheGetAttrNotNull_1), int32(644), int32(_a_F_SysCacheGetAttrNotNull_2))
						mBase = m.M
						v54 = m.ExcPending
						if v54 != 0 {
							return int32(0)
						} else {
							base.Wasm_trap_unreachable()
							for {
							}
						}
					}
				}
			}
		} else {
			m.G0 = v8 + int32(16)
			return v12
		}
	}
}
func F_SysLoggerMain(m *base.Module, l0 int32, l1 int32) {
	mBase := m.M
	_ = mBase
	var v24 int32
	_ = v24
	var v26 int32
	_ = v26
	var v29 int32
	_ = v29
	var v31 int32
	_ = v31
	var v39 int64
	_ = v39
	var v44 int32
	_ = v44
	var v45 int32
	_ = v45
	var v47 int32
	_ = v47
	var v53 int32
	_ = v53
	var v56 int32
	_ = v56
	var v58 int32
	_ = v58
	var v60 int32
	_ = v60
	var v64 int32
	_ = v64
	var v66 int32
	_ = v66
	var v67 int32
	_ = v67
	var v70 int32
	_ = v70
	var v73 int32
	_ = v73
	var v78 int32
	_ = v78
	var v80 int32
	_ = v80
	var v82 int32
	_ = v82
	var v92 int32
	_ = v92
	var v104 int32
	_ = v104
	var v110 int32
	_ = v110
	var v112 int32
	_ = v112
	var v114 int64
	_ = v114
	var v116 int64
	_ = v116
	var v124 int32
	_ = v124
	var v126 int32
	_ = v126
	var v128 int32
	_ = v128
	var v138 int32
	_ = v138
	var v150 int32
	_ = v150
	var v157 int32
	_ = v157
	var v158 int32
	_ = v158
	var v160 int64
	_ = v160
	var v162 int64
	_ = v162
	var v170 int32
	_ = v170
	var v172 int32
	_ = v172
	var v174 int32
	_ = v174
	var v184 int32
	_ = v184
	var v196 int32
	_ = v196
	var v203 int32
	_ = v203
	var v204 int32
	_ = v204
	var v206 int64
	_ = v206
	var v208 int64
	_ = v208
	var v216 int32
	_ = v216
	var v218 int32
	_ = v218
	var v220 int32
	_ = v220
	var v230 int32
	_ = v230
	var v242 int32
	_ = v242
	var v249 int32
	_ = v249
	var v250 int32
	_ = v250
	var v252 int64
	_ = v252
	var v254 int64
	_ = v254
	var v262 int32
	_ = v262
	var v264 int32
	_ = v264
	var v266 int32
	_ = v266
	var v276 int32
	_ = v276
	var v288 int32
	_ = v288
	var v295 int32
	_ = v295
	var v296 int32
	_ = v296
	var v298 int64
	_ = v298
	var v300 int64
	_ = v300
	var v308 int32
	_ = v308
	var v310 int32
	_ = v310
	var v312 int32
	_ = v312
	var v322 int32
	_ = v322
	var v334 int32
	_ = v334
	var v341 int32
	_ = v341
	var v342 int32
	_ = v342
	var v344 int64
	_ = v344
	var v346 int64
	_ = v346
	var v354 int32
	_ = v354
	var v356 int32
	_ = v356
	var v358 int32
	_ = v358
	var v368 int32
	_ = v368
	var v380 int32
	_ = v380
	var v387 int32
	_ = v387
	var v388 int32
	_ = v388
	var v390 int64
	_ = v390
	var v392 int64
	_ = v392
	var v400 int32
	_ = v400
	var v402 int32
	_ = v402
	var v404 int32
	_ = v404
	var v414 int32
	_ = v414
	var v426 int32
	_ = v426
	var v433 int32
	_ = v433
	var v434 int32
	_ = v434
	var v436 int64
	_ = v436
	var v438 int64
	_ = v438
	var v446 int32
	_ = v446
	var v448 int32
	_ = v448
	var v450 int32
	_ = v450
	var v460 int32
	_ = v460
	var v472 int32
	_ = v472
	var v479 int32
	_ = v479
	var v480 int32
	_ = v480
	var v482 int64
	_ = v482
	var v484 int64
	_ = v484
	var v494 int32
	_ = v494
	var v496 int64
	_ = v496
	var v499 int32
	_ = v499
	var v500 int32
	_ = v500
	var v502 int32
	_ = v502
	var v506 int32
	_ = v506
	var v507 int32
	_ = v507
	var v508 int32
	_ = v508
	var v513 int32
	_ = v513
	var v517 int32
	_ = v517
	var v518 int32
	_ = v518
	var v519 int32
	_ = v519
	var v520 int32
	_ = v520
	var v521 int32
	_ = v521
	var v525 int32
	_ = v525
	var v528 int64
	_ = v528
	var v530 int32
	_ = v530
	var v531 int32
	_ = v531
	var v534 int32
	_ = v534
	var v537 int64
	_ = v537
	var v539 int32
	_ = v539
	var v540 int32
	_ = v540
	var v543 int32
	_ = v543
	var v544 int32
	_ = v544
	var v545 int32
	_ = v545
	var v547 int32
	_ = v547
	var v548 int32
	_ = v548
	var v549 int32
	_ = v549
	var v551 int32
	_ = v551
	var v554 int64
	_ = v554
	var v559 int32
	_ = v559
	var v560 int32
	_ = v560
	var v561 int32
	_ = v561
	var v563 int64
	_ = v563
	var v566 int64
	_ = v566
	var v568 int64
	_ = v568
	var v570 int64
	_ = v570
	var v577 int32
	_ = v577
	var v579 int32
	_ = v579
	var v583 int32
	_ = v583
	var v584 int32
	_ = v584
	var v588 int32
	_ = v588
	var v590 int32
	_ = v590
	var v593 int32
	_ = v593
	var v596 int32
	_ = v596
	var v598 int32
	_ = v598
	var v608 int32
	_ = v608
	var v609 int32
	_ = v609
	var v610 int32
	_ = v610
	var v620 int64
	_ = v620
	var v622 int32
	_ = v622
	var v626 int32
	_ = v626
	var v632 int32
	_ = v632
	var v634 int32
	_ = v634
	var v637 int32
	_ = v637
	var v640 int32
	_ = v640
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
	var v658 int32
	_ = v658
	var v659 int32
	_ = v659
	var v662 int32
	_ = v662
	var v664 int32
	_ = v664
	var v665 int32
	_ = v665
	var v666 int32
	_ = v666
	var v671 int32
	_ = v671
	var v673 int32
	_ = v673
	var v674 int32
	_ = v674
	var v675 int32
	_ = v675
	var v677 int32
	_ = v677
	var v680 int32
	_ = v680
	var v683 int32
	_ = v683
	var v686 int32
	_ = v686
	var v687 int32
	_ = v687
	var v690 int32
	_ = v690
	var v691 int32
	_ = v691
	var v694 int32
	_ = v694
	var v701 int32
	_ = v701
	var v702 int32
	_ = v702
	var v705 int32
	_ = v705
	var v707 int32
	_ = v707
	var v708 int32
	_ = v708
	var v709 int32
	_ = v709
	var v713 int32
	_ = v713
	var v715 int32
	_ = v715
	var v718 int32
	_ = v718
	var v721 int32
	_ = v721
	var v732 int32
	_ = v732
	var v735 int32
	_ = v735
	var v745 int32
	_ = v745
	var v749 int64
	_ = v749
	var v754 int32
	_ = v754
	var v755 int32
	_ = v755
	var v756 int32
	_ = v756
	var v758 int64
	_ = v758
	var v761 int64
	_ = v761
	var v763 int64
	_ = v763
	var v765 int64
	_ = v765
	var v772 int32
	_ = v772
	var v776 int32
	_ = v776
	var v781 int32
	_ = v781
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
	var v794 int32
	_ = v794
	var v798 int32
	_ = v798
	var v800 int32
	_ = v800
	var v803 int64
	_ = v803
	var v805 int64
	_ = v805
	var v807 int32
	_ = v807
	var v811 int32
	_ = v811
	var v812 int64
	_ = v812
	var v814 int32
	_ = v814
	var v818 int32
	_ = v818
	var v820 int32
	_ = v820
	var v826 int32
	_ = v826
	var v827 int64
	_ = v827
	var v829 int64
	_ = v829
	var v834 int32
	_ = v834
	var v837 int32
	_ = v837
	var v839 int32
	_ = v839
	var v842 int64
	_ = v842
	var v844 int64
	_ = v844
	var v853 int32
	_ = v853
	var v855 int32
	_ = v855
	var v858 int64
	_ = v858
	var v860 int64
	_ = v860
	var v870 int32
	_ = v870
	var v872 int32
	_ = v872
	var v879 int32
	_ = v879
	var v880 int32
	_ = v880
	var v882 int64
	_ = v882
	var v883 int64
	_ = v883
	var v884 int64
	_ = v884
	var v888 int32
	_ = v888
	var v889 int32
	_ = v889
	var v895 int32
	_ = v895
	var v896 int32
	_ = v896
	var v902 int32
	_ = v902
	var v903 int32
	_ = v903
	var v907 int32
	_ = v907
	var v909 int32
	_ = v909
	var v912 int64
	_ = v912
	var v917 int32
	_ = v917
	var v918 int32
	_ = v918
	var v919 int32
	_ = v919
	var v921 int64
	_ = v921
	var v924 int64
	_ = v924
	var v926 int64
	_ = v926
	var v928 int64
	_ = v928
	var v935 int32
	_ = v935
	var v937 int32
	_ = v937
	var v941 int32
	_ = v941
	var v944 int64
	_ = v944
	var v946 int64
	_ = v946
	var v947 int64
	_ = v947
	var v950 int64
	_ = v950
	var v957 int32
	_ = v957
	var v958 int32
	_ = v958
	var v964 int32
	_ = v964
	var v965 int32
	_ = v965
	var v968 int32
	_ = v968
	var v972 int32
	_ = v972
	var v978 int32
	_ = v978
	var v982 int32
	_ = v982
	var v987 int32
	_ = v987
	var v988 int32
	_ = v988
	var v992 int32
	_ = v992
	var v996 int32
	_ = v996
	var v1001 int32
	_ = v1001
	var v1002 int32
	_ = v1002
	var v1008 int32
	_ = v1008
	var v1011 int32
	_ = v1011
	var v1016 int32
	_ = v1016
	var v1031 int32
	_ = v1031
	var v1032 int32
	_ = v1032
	var v1033 int32
	_ = v1033
	var v1040 int32
	_ = v1040
	var v1043 int32
	_ = v1043
	var v1046 int32
	_ = v1046
	var v1050 int32
	_ = v1050
	var v1061 int32
	_ = v1061
	var v1062 int32
	_ = v1062
	var v1063 int32
	_ = v1063
	var v1065 int32
	_ = v1065
	var v1067 int32
	_ = v1067
	var v1070 int32
	_ = v1070
	var v1071 int32
	_ = v1071
	var v1072 int32
	_ = v1072
	var v1075 int32
	_ = v1075
	var v1082 int32
	_ = v1082
	var v1083 int32
	_ = v1083
	var v1103 int32
	_ = v1103
	var v1104 int32
	_ = v1104
	var v1107 int32
	_ = v1107
	var v1109 int32
	_ = v1109
	var v1116 int32
	_ = v1116
	var v1136 int32
	_ = v1136
	var v1140 int32
	_ = v1140
	var v1167 int32
	_ = v1167
	var v1171 int32
	_ = v1171
	var v1172 int32
	_ = v1172
	var v1173 int32
	_ = v1173
	var v1174 int32
	_ = v1174
	var v1176 int32
	_ = v1176
	var v1179 int32
	_ = v1179
	var v1181 int32
	_ = v1181
	var v1185 int32
	_ = v1185
	var v1191 int32
	_ = v1191
	var v1192 int32
	_ = v1192
	var v1193 int32
	_ = v1193
	var v1197 int32
	_ = v1197
	var v1200 int32
	_ = v1200
	var v1202 int32
	_ = v1202
	var v1203 int32
	_ = v1203
	var v1208 int32
	_ = v1208
	var v1209 int32
	_ = v1209
	var v1212 int32
	_ = v1212
	var v1213 int32
	_ = v1213
	var v1218 int32
	_ = v1218
	var v1221 int32
	_ = v1221
	var v1223 int32
	_ = v1223
	var v1227 int32
	_ = v1227
	var v1230 int32
	_ = v1230
	var v1232 int32
	_ = v1232
	var v1233 int32
	_ = v1233
	var v1238 int32
	_ = v1238
	var v1240 int32
	_ = v1240
	var v1244 int32
	_ = v1244
	var v1245 int32
	_ = v1245
	var v1253 int32
	_ = v1253
	var v1275 int32
	_ = v1275
	var v1279 int32
	_ = v1279
	var v1281 int32
	_ = v1281
	var v1284 int32
	_ = v1284
	var v1285 int32
	_ = v1285
	var v1286 int32
	_ = v1286
	var v1290 int32
	_ = v1290
	var v1296 int32
	_ = v1296
	var v1314 int32
	_ = v1314
	var v1317 int32
	_ = v1317
	var v1323 int32
	_ = v1323
	var v1338 int32
	_ = v1338
	var v1339 int32
	_ = v1339
	var v1342 int32
	_ = v1342
	var v1345 int32
	_ = v1345
	var v1365 int32
	_ = v1365
	var v1368 int32
	_ = v1368
	var v1377 int32
	_ = v1377
	var v1385 int32
	_ = v1385
	var v1404 int32
	_ = v1404
	var v1407 int32
	_ = v1407
	var v1408 int32
	_ = v1408
	var v1411 int32
	_ = v1411
	var v1414 int32
	_ = v1414
	var v1434 int32
	_ = v1434
	var v1438 int32
	_ = v1438
	var v1439 int32
	_ = v1439
	var v1440 int32
	_ = v1440
	var v1442 int32
	_ = v1442
	var v1444 int32
	_ = v1444
	var v1445 int32
	_ = v1445
	var v1446 int32
	_ = v1446
	var v1451 int32
	_ = v1451
	var v1454 int32
	_ = v1454
	var v1456 int32
	_ = v1456
	var v1457 int32
	_ = v1457
	var v1458 int32
	_ = v1458
	var v1460 int32
	_ = v1460
	var v1486 int32
	_ = v1486
	var v1489 int32
	_ = v1489
	var v1496 int32
	_ = v1496
	var v1497 int32
	_ = v1497
	var v1498 int32
	_ = v1498
	var v1503 int32
	_ = v1503
	var v1504 int32
	_ = v1504
	var v1528 int32
	_ = v1528
	var v1533 int32
	_ = v1533
	var v1534 int32
	_ = v1534
	var v1538 int32
	_ = v1538
	var v1543 int32
	_ = v1543
	var v1546 int32
	_ = v1546
	v24 = m.G0
	v26 = v24 - int32(_a_F_SysLoggerMain_0)
	m.G0 = v26
	v29 = *(*int32)(unsafe.Add(mBase, _c_F_SysLoggerMain[0]))
	if v29 != 0 {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	F_MemoryContextDelete(m, v29)
	mBase = m.M
	v31 = m.ExcPending
	if v31 != 0 {
		goto L4
	} else {
		goto L5
	}
L2:
	;
	goto L3
L3:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_SysLoggerMain[1])) = int32(17)
	v39 = *(*int64)(unsafe.Add(mBase, _c_F_SysLoggerMain[2]))
	goto L7
L4:
	;
	return
L5:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_SysLoggerMain[0])) = int32(0)
	goto L3
L6:
	;
	v47 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_SysLoggerMain[3])))
	if v47 != int32(1) {
		goto L10
	} else {
		goto L11
	}
L7:
	;
	v44 = *(*int32)(unsafe.Add(mBase, _c_F_SysLoggerMain[1]))
	v45 = F_GetBackendTypeDesc(m, v44)
	mBase = m.M
	goto L9
L9:
	;
	goto L6
L10:
	;
	v70 = *(*int32)(unsafe.Add(mBase, _c_F_SysLoggerMain[4]))
	if int32(0) <= v70 {
		goto L13
	} else {
		goto L14
	}
L11:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v26)+16)) = int32(0)
	v53 = int32(1)
	v56 = F_open(m, int32(_a_F_SysLoggerMain_1), v53, v26+int32(16))
	mBase = m.M
	v58 = F_close(m, v53)
	mBase = m.M
	v60 = F_close(m, int32(2))
	mBase = m.M
	if v56 == int32(-1) {
		goto L10
	} else {
		goto L12
	}
L12:
	;
	v64 = F_dup2(m, v56, int32(1))
	mBase = m.M
	v66 = F_dup2(m, v56, int32(2))
	mBase = m.M
	v67 = F_close(m, v56)
	mBase = m.M
	goto L10
L13:
	;
	v73 = F_close(m, v70)
	mBase = m.M
	goto L15
L14:
	;
	goto L15
L15:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_SysLoggerMain[4])) = int32(-1)
	v78 = int32(914)
	v80 = m.G0
	v82 = v80 - int32(32)
	m.G0 = v82
	switch int32(916) {
	case 0, 2:
		v92 = v78
		goto L17
	default:
		goto L18
	}
L16:
	;
	v124 = int32(-2)
	v126 = m.G0
	v128 = v126 - int32(32)
	m.G0 = v128
	switch int32(0) {
	case 0, 2:
		v138 = v124
		goto L30
	default:
		goto L31
	}
L17:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v82)+12)) = v92
	F_sigemptyset(m, v82+int32(16))
	mBase = m.M
	goto L20
L18:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_SysLoggerMain[5])) = v78
	v92 = int32(_a_F_SysLoggerMain_2)
	goto L17
L20:
	;
	goto L21
L21:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v82)+24)) = int32(268435456)
	v104 = v82 + int32(12)
	goto L24
L22:
	;
	m.G0 = v82 + int32(32)
	goto L16
L24:
	;
	goto L25
L25:
	;
	if v104 != 0 {
		goto L26
	} else {
		goto L27
	}
L26:
	;
	v110 = int32(20)
	v112 = *(*int32)(unsafe.Add(mBase, uint32(v104)+16))
	*(*int32)(unsafe.Add(mBase, _c_F_SysLoggerMain[6])) = v112
	v114 = *(*int64)(unsafe.Add(mBase, uint32(v104)+8))
	*(*int64)(unsafe.Add(mBase, _c_F_SysLoggerMain[7])) = v114
	v116 = *(*int64)(unsafe.Add(mBase, uint32(v104)))
	*(*int64)(unsafe.Add(mBase, _c_F_SysLoggerMain[8])) = v116
	goto L28
L27:
	;
	goto L28
L28:
	;
	goto L22
L29:
	;
	v170 = int32(-2)
	v172 = m.G0
	v174 = v172 - int32(32)
	m.G0 = v174
	switch int32(0) {
	case 0, 2:
		v184 = v170
		goto L43
	default:
		goto L44
	}
L30:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v128)+12)) = v138
	F_sigemptyset(m, v128+int32(16))
	mBase = m.M
	goto L33
L31:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_SysLoggerMain[9])) = v124
	v138 = int32(_a_F_SysLoggerMain_2)
	goto L30
L33:
	;
	goto L34
L34:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v128)+24)) = int32(268435456)
	v150 = v128 + int32(12)
	goto L37
L35:
	;
	m.G0 = v128 + int32(32)
	goto L29
L37:
	;
	goto L38
L38:
	;
	if v150 != 0 {
		goto L39
	} else {
		goto L40
	}
L39:
	;
	v157 = int32(40)
	v158 = *(*int32)(unsafe.Add(mBase, uint32(v150)+16))
	*(*int32)(unsafe.Add(mBase, _c_F_SysLoggerMain[10])) = v158
	v160 = *(*int64)(unsafe.Add(mBase, uint32(v150)+8))
	*(*int64)(unsafe.Add(mBase, _c_F_SysLoggerMain[11])) = v160
	v162 = *(*int64)(unsafe.Add(mBase, uint32(v150)))
	*(*int64)(unsafe.Add(mBase, _c_F_SysLoggerMain[12])) = v162
	goto L41
L40:
	;
	goto L41
L41:
	;
	goto L35
L42:
	;
	v216 = int32(-2)
	v218 = m.G0
	v220 = v218 - int32(32)
	m.G0 = v220
	switch int32(0) {
	case 0, 2:
		v230 = v216
		goto L56
	default:
		goto L57
	}
L43:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v174)+12)) = v184
	F_sigemptyset(m, v174+int32(16))
	mBase = m.M
	goto L46
L44:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_SysLoggerMain[13])) = v170
	v184 = int32(_a_F_SysLoggerMain_2)
	goto L43
L46:
	;
	goto L47
L47:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v174)+24)) = int32(268435456)
	v196 = v174 + int32(12)
	goto L50
L48:
	;
	m.G0 = v174 + int32(32)
	goto L42
L50:
	;
	goto L51
L51:
	;
	if v196 != 0 {
		goto L52
	} else {
		goto L53
	}
L52:
	;
	v203 = int32(300)
	v204 = *(*int32)(unsafe.Add(mBase, uint32(v196)+16))
	*(*int32)(unsafe.Add(mBase, _c_F_SysLoggerMain[14])) = v204
	v206 = *(*int64)(unsafe.Add(mBase, uint32(v196)+8))
	*(*int64)(unsafe.Add(mBase, _c_F_SysLoggerMain[15])) = v206
	v208 = *(*int64)(unsafe.Add(mBase, uint32(v196)))
	*(*int64)(unsafe.Add(mBase, _c_F_SysLoggerMain[16])) = v208
	goto L54
L53:
	;
	goto L54
L54:
	;
	goto L48
L55:
	;
	v262 = int32(-2)
	v264 = m.G0
	v266 = v264 - int32(32)
	m.G0 = v266
	switch int32(0) {
	case 0, 2:
		v276 = v262
		goto L69
	default:
		goto L70
	}
L56:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v220)+12)) = v230
	F_sigemptyset(m, v220+int32(16))
	mBase = m.M
	goto L59
L57:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_SysLoggerMain[17])) = v216
	v230 = int32(_a_F_SysLoggerMain_2)
	goto L56
L59:
	;
	goto L60
L60:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v220)+24)) = int32(268435456)
	v242 = v220 + int32(12)
	goto L63
L61:
	;
	m.G0 = v220 + int32(32)
	goto L55
L63:
	;
	goto L64
L64:
	;
	if v242 != 0 {
		goto L65
	} else {
		goto L66
	}
L65:
	;
	v249 = int32(60)
	v250 = *(*int32)(unsafe.Add(mBase, uint32(v242)+16))
	*(*int32)(unsafe.Add(mBase, _c_F_SysLoggerMain[18])) = v250
	v252 = *(*int64)(unsafe.Add(mBase, uint32(v242)+8))
	*(*int64)(unsafe.Add(mBase, _c_F_SysLoggerMain[19])) = v252
	v254 = *(*int64)(unsafe.Add(mBase, uint32(v242)))
	*(*int64)(unsafe.Add(mBase, _c_F_SysLoggerMain[20])) = v254
	goto L67
L66:
	;
	goto L67
L67:
	;
	goto L61
L68:
	;
	v308 = int32(-2)
	v310 = m.G0
	v312 = v310 - int32(32)
	m.G0 = v312
	switch int32(0) {
	case 0, 2:
		v322 = v308
		goto L82
	default:
		goto L83
	}
L69:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v266)+12)) = v276
	F_sigemptyset(m, v266+int32(16))
	mBase = m.M
	goto L72
L70:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_SysLoggerMain[21])) = v262
	v276 = int32(_a_F_SysLoggerMain_2)
	goto L69
L72:
	;
	goto L73
L73:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v266)+24)) = int32(268435456)
	v288 = v266 + int32(12)
	goto L76
L74:
	;
	m.G0 = v266 + int32(32)
	goto L68
L76:
	;
	goto L77
L77:
	;
	if v288 != 0 {
		goto L78
	} else {
		goto L79
	}
L78:
	;
	v295 = int32(280)
	v296 = *(*int32)(unsafe.Add(mBase, uint32(v288)+16))
	*(*int32)(unsafe.Add(mBase, _c_F_SysLoggerMain[22])) = v296
	v298 = *(*int64)(unsafe.Add(mBase, uint32(v288)+8))
	*(*int64)(unsafe.Add(mBase, _c_F_SysLoggerMain[23])) = v298
	v300 = *(*int64)(unsafe.Add(mBase, uint32(v288)))
	*(*int64)(unsafe.Add(mBase, _c_F_SysLoggerMain[24])) = v300
	goto L80
L79:
	;
	goto L80
L80:
	;
	goto L74
L81:
	;
	v354 = int32(964)
	v356 = m.G0
	v358 = v356 - int32(32)
	m.G0 = v358
	switch int32(966) {
	case 0, 2:
		v368 = v354
		goto L95
	default:
		goto L96
	}
L82:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v312)+12)) = v322
	F_sigemptyset(m, v312+int32(16))
	mBase = m.M
	goto L85
L83:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_SysLoggerMain[25])) = v308
	v322 = int32(_a_F_SysLoggerMain_2)
	goto L82
L85:
	;
	goto L86
L86:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v312)+24)) = int32(268435456)
	v334 = v312 + int32(12)
	goto L89
L87:
	;
	m.G0 = v312 + int32(32)
	goto L81
L89:
	;
	goto L90
L90:
	;
	if v334 != 0 {
		goto L91
	} else {
		goto L92
	}
L91:
	;
	v341 = int32(260)
	v342 = *(*int32)(unsafe.Add(mBase, uint32(v334)+16))
	*(*int32)(unsafe.Add(mBase, _c_F_SysLoggerMain[26])) = v342
	v344 = *(*int64)(unsafe.Add(mBase, uint32(v334)+8))
	*(*int64)(unsafe.Add(mBase, _c_F_SysLoggerMain[27])) = v344
	v346 = *(*int64)(unsafe.Add(mBase, uint32(v334)))
	*(*int64)(unsafe.Add(mBase, _c_F_SysLoggerMain[28])) = v346
	goto L93
L92:
	;
	goto L93
L93:
	;
	goto L87
L94:
	;
	v400 = int32(-2)
	v402 = m.G0
	v404 = v402 - int32(32)
	m.G0 = v404
	switch int32(0) {
	case 0, 2:
		v414 = v400
		goto L108
	default:
		goto L109
	}
L95:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v358)+12)) = v368
	F_sigemptyset(m, v358+int32(16))
	mBase = m.M
	goto L98
L96:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_SysLoggerMain[29])) = v354
	v368 = int32(_a_F_SysLoggerMain_2)
	goto L95
L98:
	;
	goto L99
L99:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v358)+24)) = int32(268435456)
	v380 = v358 + int32(12)
	goto L102
L100:
	;
	m.G0 = v358 + int32(32)
	goto L94
L102:
	;
	goto L103
L103:
	;
	if v380 != 0 {
		goto L104
	} else {
		goto L105
	}
L104:
	;
	v387 = int32(200)
	v388 = *(*int32)(unsafe.Add(mBase, uint32(v380)+16))
	*(*int32)(unsafe.Add(mBase, _c_F_SysLoggerMain[30])) = v388
	v390 = *(*int64)(unsafe.Add(mBase, uint32(v380)+8))
	*(*int64)(unsafe.Add(mBase, _c_F_SysLoggerMain[31])) = v390
	v392 = *(*int64)(unsafe.Add(mBase, uint32(v380)))
	*(*int64)(unsafe.Add(mBase, _c_F_SysLoggerMain[32])) = v392
	goto L106
L105:
	;
	goto L106
L106:
	;
	goto L100
L107:
	;
	v446 = int32(0)
	v448 = m.G0
	v450 = v448 - int32(32)
	m.G0 = v450
	switch int32(2) {
	case 0, 2:
		v460 = v446
		goto L121
	default:
		goto L122
	}
L108:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v404)+12)) = v414
	F_sigemptyset(m, v404+int32(16))
	mBase = m.M
	goto L111
L109:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_SysLoggerMain[33])) = v400
	v414 = int32(_a_F_SysLoggerMain_2)
	goto L108
L111:
	;
	goto L112
L112:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v404)+24)) = int32(268435456)
	v426 = v404 + int32(12)
	goto L115
L113:
	;
	m.G0 = v404 + int32(32)
	goto L107
L115:
	;
	goto L116
L116:
	;
	if v426 != 0 {
		goto L117
	} else {
		goto L118
	}
L117:
	;
	v433 = int32(240)
	v434 = *(*int32)(unsafe.Add(mBase, uint32(v426)+16))
	*(*int32)(unsafe.Add(mBase, _c_F_SysLoggerMain[34])) = v434
	v436 = *(*int64)(unsafe.Add(mBase, uint32(v426)+8))
	*(*int64)(unsafe.Add(mBase, _c_F_SysLoggerMain[35])) = v436
	v438 = *(*int64)(unsafe.Add(mBase, uint32(v426)))
	*(*int64)(unsafe.Add(mBase, _c_F_SysLoggerMain[36])) = v438
	goto L119
L118:
	;
	goto L119
L119:
	;
	goto L113
L120:
	;
	F_sigprocmask(m, int32(_a_F_SysLoggerMain_3), int32(0))
	mBase = m.M
	v494 = m.ExcPending
	if v494 != 0 {
		goto L4
	} else {
		goto L133
	}
L121:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v450)+12)) = v460
	F_sigemptyset(m, v450+int32(16))
	mBase = m.M
	goto L123
L122:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_SysLoggerMain[37])) = v446
	v460 = int32(_a_F_SysLoggerMain_2)
	goto L121
L123:
	;
	goto L125
L125:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v450)+24)) = int32(268435457)
	v472 = v450 + int32(12)
	goto L128
L126:
	;
	m.G0 = v450 + int32(32)
	goto L120
L128:
	;
	goto L129
L129:
	;
	if v472 != 0 {
		goto L130
	} else {
		goto L131
	}
L130:
	;
	v479 = int32(340)
	v480 = *(*int32)(unsafe.Add(mBase, uint32(v472)+16))
	*(*int32)(unsafe.Add(mBase, _c_F_SysLoggerMain[38])) = v480
	v482 = *(*int64)(unsafe.Add(mBase, uint32(v472)+8))
	*(*int64)(unsafe.Add(mBase, _c_F_SysLoggerMain[39])) = v482
	v484 = *(*int64)(unsafe.Add(mBase, uint32(v472)))
	*(*int64)(unsafe.Add(mBase, _c_F_SysLoggerMain[40])) = v484
	goto L132
L131:
	;
	goto L132
L132:
	;
	goto L126
L133:
	;
	v496 = *(*int64)(unsafe.Add(mBase, _c_F_SysLoggerMain[41]))
	*(*int64)(unsafe.Add(mBase, uint32(v26)+48)) = v496
	v499 = F_palloc(m, int32(1024))
	mBase = m.M
	v500 = m.ExcPending
	if v500 != 0 {
		goto L4
	} else {
		goto L134
	}
L134:
	;
	v502 = *(*int32)(unsafe.Add(mBase, _c_F_SysLoggerMain[42]))
	*(*int32)(unsafe.Add(mBase, uint32(v26))) = v502
	v506 = F_pg_snprintf(m, v499, int32(1024), int32(_a_F_SysLoggerMain_4), v26)
	mBase = m.M
	v507 = m.ExcPending
	if v507 != 0 {
		goto L4
	} else {
		goto L135
	}
L135:
	;
	v508 = F_strlen(m, v499)
	mBase = m.M
	v513 = *(*int32)(unsafe.Add(mBase, _c_F_SysLoggerMain[43]))
	v517 = *(*int32)(unsafe.Add(mBase, _c_F_SysLoggerMain[44]))
	v518 = F_pg_localtime(m, v26+int32(48), v517)
	mBase = m.M
	v519 = m.ExcPending
	if v519 != 0 {
		goto L4
	} else {
		goto L136
	}
L136:
	;
	v520 = F_pg_strftime(m, v499+v508, int32(1024)-v508, v513, v518)
	mBase = m.M
	v521 = m.ExcPending
	if v521 != 0 {
		goto L4
	} else {
		goto L137
	}
L137:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_SysLoggerMain[45])) = v499
	v525 = *(*int32)(unsafe.Add(mBase, _c_F_SysLoggerMain[46]))
	if v525 != 0 {
		goto L138
	} else {
		goto L139
	}
L138:
	;
	v528 = *(*int64)(unsafe.Add(mBase, _c_F_SysLoggerMain[41]))
	v530 = F_logfile_getname(m, v528, int32(_a_F_SysLoggerMain_5))
	mBase = m.M
	v531 = m.ExcPending
	if v531 != 0 {
		goto L4
	} else {
		goto L141
	}
L139:
	;
	goto L140
L140:
	;
	v534 = *(*int32)(unsafe.Add(mBase, _c_F_SysLoggerMain[47]))
	if v534 != 0 {
		goto L142
	} else {
		goto L143
	}
L141:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_SysLoggerMain[48])) = v530
	goto L140
L142:
	;
	v537 = *(*int64)(unsafe.Add(mBase, _c_F_SysLoggerMain[41]))
	v539 = F_logfile_getname(m, v537, int32(_a_F_SysLoggerMain_6))
	mBase = m.M
	v540 = m.ExcPending
	if v540 != 0 {
		goto L4
	} else {
		goto L145
	}
L143:
	;
	goto L144
L144:
	;
	v543 = *(*int32)(unsafe.Add(mBase, _c_F_SysLoggerMain[42]))
	v544 = F_pstrdup(m, v543)
	mBase = m.M
	v545 = m.ExcPending
	if v545 != 0 {
		goto L4
	} else {
		goto L146
	}
L145:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_SysLoggerMain[49])) = v539
	goto L144
L146:
	;
	v547 = *(*int32)(unsafe.Add(mBase, _c_F_SysLoggerMain[43]))
	v548 = F_pstrdup(m, v547)
	mBase = m.M
	v549 = m.ExcPending
	if v549 != 0 {
		goto L4
	} else {
		goto L147
	}
L147:
	;
	v551 = *(*int32)(unsafe.Add(mBase, _c_F_SysLoggerMain[50]))
	if int32(0) < v551 {
		goto L148
	} else {
		goto L149
	}
L148:
	;
	v554 = F_time(m)
	mBase = m.M
	*(*int64)(unsafe.Add(mBase, uint32(v26)+48)) = v554
	v559 = *(*int32)(unsafe.Add(mBase, _c_F_SysLoggerMain[44]))
	v560 = F_pg_localtime(m, v26+int32(48), v559)
	mBase = m.M
	v561 = m.ExcPending
	if v561 != 0 {
		goto L4
	} else {
		goto L151
	}
L149:
	;
	goto L150
L150:
	;
	F_update_metainfo_datafile(m)
	mBase = m.M
	v577 = m.ExcPending
	if v577 != 0 {
		goto L4
	} else {
		goto L152
	}
L151:
	;
	v563 = *(*int64)(unsafe.Add(mBase, uint32(v26)+48))
	v566 = base.I64_extend_i32_s(v551 * int32(60))
	v568 = int64(*(*int32)(unsafe.Add(mBase, uint32(v560)+36)))
	v570 = base.I64_rem_s(v563+v568, v566)
	*(*int64)(unsafe.Add(mBase, _c_F_SysLoggerMain[51])) = v563 + v566 - v570
	goto L150
L152:
	;
	v579 = int32(0)
	*(*int32)(unsafe.Add(mBase, _c_F_SysLoggerMain[52])) = v579
	v583 = F_CreateWaitEventSet(m, v579, int32(2))
	mBase = m.M
	v584 = m.ExcPending
	if v584 != 0 {
		goto L4
	} else {
		goto L153
	}
L153:
	;
	v588 = *(*int32)(unsafe.Add(mBase, _c_F_SysLoggerMain[53]))
	F_AddWaitEventToSet(m, v583, int32(1), int32(-1), v588)
	mBase = m.M
	v590 = m.ExcPending
	if v590 != 0 {
		goto L4
	} else {
		goto L154
	}
L154:
	;
	v593 = *(*int32)(unsafe.Add(mBase, _c_F_SysLoggerMain[54]))
	F_AddWaitEventToSet(m, v583, int32(2), v593, int32(0))
	mBase = m.M
	v596 = m.ExcPending
	if v596 != 0 {
		goto L4
	} else {
		goto L155
	}
L155:
	;
	v598 = int32(0)
	v608 = v551
	v609 = v544
	v610 = v548
	v620 = v39
	goto L156
L156:
	;
	v622 = *(*int32)(unsafe.Add(mBase, _c_F_SysLoggerMain[53]))
	*(*int32)(unsafe.Add(mBase, uint32(v622))) = int32(0)
	goto L158
L157:
	;
	v1533 = F_errstart(m, int32(14), int32(0))
	mBase = m.M
	v1534 = m.ExcPending
	if v1534 != 0 {
		goto L4
	} else {
		goto L387
	}
L158:
	;
	v626 = *(*int32)(unsafe.Add(mBase, _c_F_SysLoggerMain[55]))
	if v626 != 0 {
		goto L159
	} else {
		goto L160
	}
L159:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_SysLoggerMain[55])) = int32(0)
	F_ProcessConfigFile(m, int32(2))
	mBase = m.M
	v632 = m.ExcPending
	if v632 != 0 {
		goto L4
	} else {
		goto L162
	}
L160:
	;
	v787 = v608
	v788 = v609
	v789 = v610
	goto L161
L161:
	;
	v792 = int32(0)
	v794 = *(*int32)(unsafe.Add(mBase, _c_F_SysLoggerMain[50]))
	if v794 <= v792 {
		goto L206
	} else {
		goto L207
	}
L162:
	;
	v634 = *(*int32)(unsafe.Add(mBase, _c_F_SysLoggerMain[42]))
	v637 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v634))))
	v640 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v609))))
	if base.B2i32(v637 == int32(0))|base.B2i32(v637 != v640) != 0 {
		v658 = v637
		v659 = v640
		goto L164
	} else {
		goto L165
	}
L163:
	;
	if v658-v659 != 0 {
		goto L170
	} else {
		goto L171
	}
L164:
	;
	goto L163
L165:
	;
	v643 = v634
	v644 = v609
	goto L166
L166:
	;
	v647 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v644)+1)))
	v648 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v643)+1)))
	if v648 == int32(0) {
		v658 = v648
		v659 = v647
		goto L164
	} else {
		goto L168
	}
L167:
	;
	v658 = v648
	v659 = v647
	goto L164
L168:
	;
	v651 = int32(1)
	if v648 == v647 {
		v643 = v643 + v651
		v644 = v644 + v651
		goto L166
	} else {
		goto L169
	}
L169:
	;
	goto L167
L170:
	;
	F_pfree(m, v609)
	mBase = m.M
	v662 = m.ExcPending
	if v662 != 0 {
		goto L4
	} else {
		goto L173
	}
L171:
	;
	v675 = v609
	goto L172
L172:
	;
	v677 = *(*int32)(unsafe.Add(mBase, _c_F_SysLoggerMain[43]))
	v680 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v677))))
	v683 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v610))))
	if base.B2i32(v680 == int32(0))|base.B2i32(v680 != v683) != 0 {
		v701 = v680
		v702 = v683
		goto L177
	} else {
		goto L178
	}
L173:
	;
	v664 = *(*int32)(unsafe.Add(mBase, _c_F_SysLoggerMain[42]))
	v665 = F_pstrdup(m, v664)
	mBase = m.M
	v666 = m.ExcPending
	if v666 != 0 {
		goto L4
	} else {
		goto L174
	}
L174:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_SysLoggerMain[56])) = int32(1)
	v671 = *(*int32)(unsafe.Add(mBase, _c_F_SysLoggerMain[42]))
	v673 = *(*int32)(unsafe.Add(mBase, _c_F_SysLoggerMain[57]))
	v674 = F_mkdir(m, v671, v673)
	mBase = m.M
	goto L175
L175:
	;
	v675 = v665
	goto L172
L176:
	;
	if v701-v702 != 0 {
		goto L183
	} else {
		goto L184
	}
L177:
	;
	goto L176
L178:
	;
	v686 = v677
	v687 = v610
	goto L179
L179:
	;
	v690 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v687)+1)))
	v691 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v686)+1)))
	if v691 == int32(0) {
		v701 = v691
		v702 = v690
		goto L177
	} else {
		goto L181
	}
L180:
	;
	v701 = v691
	v702 = v690
	goto L177
L181:
	;
	v694 = int32(1)
	if v691 == v690 {
		v686 = v686 + v694
		v687 = v687 + v694
		goto L179
	} else {
		goto L182
	}
L182:
	;
	goto L180
L183:
	;
	F_pfree(m, v610)
	mBase = m.M
	v705 = m.ExcPending
	if v705 != 0 {
		goto L4
	} else {
		goto L186
	}
L184:
	;
	v713 = v610
	goto L185
L185:
	;
	v715 = *(*int32)(unsafe.Add(mBase, _c_F_SysLoggerMain[58]))
	v718 = int32(0)
	v721 = *(*int32)(unsafe.Add(mBase, _c_F_SysLoggerMain[46]))
	if base.B2i32(v715&int32(8) == v718)^base.B2i32(v721 != v718) == v718 {
		goto L188
	} else {
		goto L189
	}
L186:
	;
	v707 = *(*int32)(unsafe.Add(mBase, _c_F_SysLoggerMain[43]))
	v708 = F_pstrdup(m, v707)
	mBase = m.M
	v709 = m.ExcPending
	if v709 != 0 {
		goto L4
	} else {
		goto L187
	}
L187:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_SysLoggerMain[56])) = int32(1)
	v713 = v708
	goto L185
L188:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_SysLoggerMain[56])) = int32(1)
	goto L190
L189:
	;
	goto L190
L190:
	;
	v732 = int32(0)
	v735 = *(*int32)(unsafe.Add(mBase, _c_F_SysLoggerMain[47]))
	if base.B2i32(v715&int32(16) == v732)^base.B2i32(v735 != v732) == v732 {
		goto L191
	} else {
		goto L192
	}
L191:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_SysLoggerMain[56])) = int32(1)
	goto L193
L192:
	;
	goto L193
L193:
	;
	v745 = *(*int32)(unsafe.Add(mBase, _c_F_SysLoggerMain[50]))
	if v745 != v608 {
		goto L194
	} else {
		goto L195
	}
L194:
	;
	if int32(0) < v745 {
		goto L197
	} else {
		goto L198
	}
L195:
	;
	v772 = v608
	goto L196
L196:
	;
	v776 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_SysLoggerMain[59])))
	if v776 != 0 {
		goto L201
	} else {
		goto L202
	}
L197:
	;
	v749 = F_time(m)
	mBase = m.M
	*(*int64)(unsafe.Add(mBase, uint32(v26)+32)) = v749
	v754 = *(*int32)(unsafe.Add(mBase, _c_F_SysLoggerMain[44]))
	v755 = F_pg_localtime(m, v26+int32(32), v754)
	mBase = m.M
	v756 = m.ExcPending
	if v756 != 0 {
		goto L4
	} else {
		goto L200
	}
L198:
	;
	goto L199
L199:
	;
	v772 = v745
	goto L196
L200:
	;
	v758 = *(*int64)(unsafe.Add(mBase, uint32(v26)+32))
	v761 = base.I64_extend_i32_s(v745 * int32(60))
	v763 = int64(*(*int32)(unsafe.Add(mBase, uint32(v755)+36)))
	v765 = base.I64_rem_s(v758+v763, v761)
	*(*int64)(unsafe.Add(mBase, _c_F_SysLoggerMain[51])) = v758 + v761 - v765
	goto L199
L201:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_SysLoggerMain[56])) = int32(1)
	v781 = int32(0)
	*(*uint8)(unsafe.Add(mBase, _c_F_SysLoggerMain[59])) = uint8(v781)
	goto L203
L202:
	;
	goto L203
L203:
	;
	F_update_metainfo_datafile(m)
	mBase = m.M
	v784 = m.ExcPending
	if v784 != 0 {
		goto L4
	} else {
		goto L204
	}
L204:
	;
	v787 = v772
	v788 = v675
	v789 = v713
	goto L161
L205:
	;
	v814 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_SysLoggerMain[59])))
	v818 = *(*int32)(unsafe.Add(mBase, _c_F_SysLoggerMain[56]))
	v820 = *(*int32)(unsafe.Add(mBase, _c_F_SysLoggerMain[60]))
	if v814&int32(1)|(v818|base.B2i32(v820 <= int32(0))) != 0 {
		v870 = v792
		goto L211
	} else {
		goto L212
	}
L206:
	;
	v811 = int32(0)
	v812 = v620
	goto L205
L207:
	;
	goto L208
L208:
	;
	v798 = int32(0)
	v800 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_SysLoggerMain[59])))
	if v800&int32(1) != 0 {
		v811 = v798
		v812 = v620
		goto L205
	} else {
		goto L209
	}
L209:
	;
	v803 = F_time(m)
	mBase = m.M
	v805 = *(*int64)(unsafe.Add(mBase, _c_F_SysLoggerMain[51]))
	if v803 < v805 {
		v811 = v798
		v812 = v803
		goto L205
	} else {
		goto L210
	}
L210:
	;
	v807 = int32(1)
	*(*int32)(unsafe.Add(mBase, _c_F_SysLoggerMain[56])) = v807
	v811 = v807
	v812 = v803
	goto L205
L211:
	;
	v872 = *(*int32)(unsafe.Add(mBase, _c_F_SysLoggerMain[56]))
	if v872 == int32(0) {
		goto L221
	} else {
		goto L222
	}
L212:
	;
	v826 = *(*int32)(unsafe.Add(mBase, _c_F_SysLoggerMain[61]))
	v827 = F___ftello_unlocked(m, v826)
	mBase = m.M
	v829 = int64(*(*int32)(unsafe.Add(mBase, _c_F_SysLoggerMain[60])))
	if v829<<(uint(int64(10))%64) <= v827 {
		goto L213
	} else {
		goto L214
	}
L213:
	;
	v834 = int32(1)
	*(*int32)(unsafe.Add(mBase, _c_F_SysLoggerMain[56])) = v834
	v837 = v834
	goto L215
L214:
	;
	v837 = v792
	goto L215
L215:
	;
	v839 = *(*int32)(unsafe.Add(mBase, _c_F_SysLoggerMain[46]))
	if v839 == int32(0) {
		v853 = v837
		goto L216
	} else {
		goto L217
	}
L216:
	;
	v855 = *(*int32)(unsafe.Add(mBase, _c_F_SysLoggerMain[47]))
	if v855 == int32(0) {
		v870 = v853
		goto L211
	} else {
		goto L219
	}
L217:
	;
	v842 = F___ftello_unlocked(m, v839)
	mBase = m.M
	v844 = int64(*(*int32)(unsafe.Add(mBase, _c_F_SysLoggerMain[60])))
	if v842 < v844<<(uint(int64(10))%64) {
		v853 = v837
		goto L216
	} else {
		goto L218
	}
L218:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_SysLoggerMain[56])) = int32(1)
	v853 = v837 | int32(8)
	goto L216
L219:
	;
	v858 = F___ftello_unlocked(m, v855)
	mBase = m.M
	v860 = int64(*(*int32)(unsafe.Add(mBase, _c_F_SysLoggerMain[60])))
	if v858 < v860<<(uint(int64(10))%64) {
		v870 = v853
		goto L211
	} else {
		goto L220
	}
L220:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_SysLoggerMain[56])) = int32(1)
	v870 = v853 | int32(16)
	goto L211
L221:
	;
	v935 = int32(-1)
	v937 = *(*int32)(unsafe.Add(mBase, _c_F_SysLoggerMain[50]))
	if v937 <= int32(0) {
		v958 = v935
		goto L242
	} else {
		goto L243
	}
L222:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_SysLoggerMain[56])) = int32(0)
	if v870 != 0 {
		goto L223
	} else {
		goto L224
	}
L223:
	;
	v879 = v870
	goto L225
L224:
	;
	v879 = int32(25)
	goto L225
L225:
	;
	if v811 != 0 {
		goto L226
	} else {
		goto L227
	}
L226:
	;
	v880 = v870
	goto L228
L227:
	;
	v880 = v879
	goto L228
L228:
	;
	if v811 != 0 {
		goto L230
	} else {
		goto L231
	}
L229:
	;
	v888 = F_logfile_rotate_dest(m, v811, v880, v884, int32(1), int32(_a_F_SysLoggerMain_7), int32(_a_F_SysLoggerMain_8))
	mBase = m.M
	v889 = m.ExcPending
	if v889 != 0 {
		goto L4
	} else {
		goto L233
	}
L230:
	;
	v882 = *(*int64)(unsafe.Add(mBase, _c_F_SysLoggerMain[51]))
	v884 = v882
	goto L229
L231:
	;
	goto L232
L232:
	;
	v883 = F_time(m)
	mBase = m.M
	v884 = v883
	goto L229
L233:
	;
	if v888 == int32(0) {
		goto L221
	} else {
		goto L234
	}
L234:
	;
	v895 = F_logfile_rotate_dest(m, v811, v880, v884, int32(8), int32(_a_F_SysLoggerMain_9), int32(_a_F_SysLoggerMain_10))
	mBase = m.M
	v896 = m.ExcPending
	if v896 != 0 {
		goto L4
	} else {
		goto L235
	}
L235:
	;
	if v895 == int32(0) {
		goto L221
	} else {
		goto L236
	}
L236:
	;
	v902 = F_logfile_rotate_dest(m, v811, v880, v884, int32(16), int32(_a_F_SysLoggerMain_11), int32(_a_F_SysLoggerMain_12))
	mBase = m.M
	v903 = m.ExcPending
	if v903 != 0 {
		goto L4
	} else {
		goto L237
	}
L237:
	;
	if v902 == int32(0) {
		goto L221
	} else {
		goto L238
	}
L238:
	;
	F_update_metainfo_datafile(m)
	mBase = m.M
	v907 = m.ExcPending
	if v907 != 0 {
		goto L4
	} else {
		goto L239
	}
L239:
	;
	v909 = *(*int32)(unsafe.Add(mBase, _c_F_SysLoggerMain[50]))
	if v909 <= int32(0) {
		goto L221
	} else {
		goto L240
	}
L240:
	;
	v912 = F_time(m)
	mBase = m.M
	*(*int64)(unsafe.Add(mBase, uint32(v26)+32)) = v912
	v917 = *(*int32)(unsafe.Add(mBase, _c_F_SysLoggerMain[44]))
	v918 = F_pg_localtime(m, v26+int32(32), v917)
	mBase = m.M
	v919 = m.ExcPending
	if v919 != 0 {
		goto L4
	} else {
		goto L241
	}
L241:
	;
	v921 = *(*int64)(unsafe.Add(mBase, uint32(v26)+32))
	v924 = base.I64_extend_i32_s(v909 * int32(60))
	v926 = int64(*(*int32)(unsafe.Add(mBase, uint32(v918)+36)))
	v928 = base.I64_rem_s(v921+v926, v924)
	*(*int64)(unsafe.Add(mBase, _c_F_SysLoggerMain[51])) = v921 + v924 - v928
	goto L221
L242:
	;
	v964 = F_WaitEventSetWait(m, v583, v958, v26+int32(32), int32(1), int32(83886093))
	mBase = m.M
	v965 = m.ExcPending
	if v965 != 0 {
		goto L4
	} else {
		goto L252
	}
L243:
	;
	v941 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_SysLoggerMain[59])))
	if v941&int32(1) != 0 {
		v958 = v935
		goto L242
	} else {
		goto L244
	}
L244:
	;
	v944 = int64(2147483)
	v946 = *(*int64)(unsafe.Add(mBase, _c_F_SysLoggerMain[51]))
	v947 = v946 - v812
	if v944 <= v947 {
		goto L245
	} else {
		goto L246
	}
L245:
	;
	v950 = v944
	goto L247
L246:
	;
	v950 = v947
	goto L247
L247:
	;
	if int64(0) < v947 {
		goto L248
	} else {
		goto L249
	}
L248:
	;
	v957 = base.I32_wrap_i64(v950) * int32(1000)
	goto L250
L249:
	;
	v957 = int32(0)
	goto L250
L250:
	;
	v958 = v957
	goto L242
L251:
	;
	v1528 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_SysLoggerMain[62])))
	if v1528 == int32(0) {
		v598 = v1504
		v608 = v787
		v609 = v788
		v610 = v789
		v620 = v812
		goto L156
	} else {
		goto L386
	}
L252:
	;
	if v964 != int32(1) {
		goto L253
	} else {
		goto L254
	}
L253:
	;
	v1504 = v598
	goto L251
L254:
	;
	goto L255
L255:
	;
	v968 = *(*int32)(unsafe.Add(mBase, uint32(v26)+36))
	if v968 != int32(2) {
		goto L256
	} else {
		goto L257
	}
L256:
	;
	v1504 = v598
	goto L251
L257:
	;
	goto L258
L258:
	;
	v972 = *(*int32)(unsafe.Add(mBase, _c_F_SysLoggerMain[54]))
	v978 = F_read(m, v972, v26+int32(48)+v598, int32(_a_F_SysLoggerMain_13)-v598)
	mBase = m.M
	if v978 < int32(0) {
		goto L259
	} else {
		goto L260
	}
L259:
	;
	v982 = *(*int32)(unsafe.Add(mBase, _c_F_SysLoggerMain[63]))
	if v982 == int32(27) {
		goto L262
	} else {
		goto L263
	}
L260:
	;
	goto L261
L261:
	;
	if v978 != 0 {
		goto L272
	} else {
		goto L273
	}
L262:
	;
	v1504 = v598
	goto L251
L263:
	;
	goto L264
L264:
	;
	v987 = F_errstart(m, int32(15), int32(0))
	mBase = m.M
	v988 = m.ExcPending
	if v988 != 0 {
		goto L4
	} else {
		goto L265
	}
L265:
	;
	if v987 == int32(0) {
		goto L266
	} else {
		goto L267
	}
L266:
	;
	v1504 = v598
	goto L251
L267:
	;
	goto L268
L268:
	;
	F_errcode_for_socket_access(m)
	mBase = m.M
	v992 = m.ExcPending
	if v992 != 0 {
		goto L4
	} else {
		goto L269
	}
L269:
	;
	F_errmsg(m, int32(_a_F_SysLoggerMain_14), int32(0))
	mBase = m.M
	v996 = m.ExcPending
	if v996 != 0 {
		goto L4
	} else {
		goto L270
	}
L270:
	;
	F_errfinish(m, int32(_a_F_SysLoggerMain_15), int32(527), int32(_a_F_SysLoggerMain_16))
	mBase = m.M
	v1001 = m.ExcPending
	if v1001 != 0 {
		goto L4
	} else {
		goto L271
	}
L271:
	;
	v1504 = v598
	goto L251
L272:
	;
	v1002 = v978 + v598
	if v1002 < int32(10) {
		v598 = v1002
		v608 = v787
		v609 = v788
		v610 = v789
		v620 = v812
		goto L156
	} else {
		goto L275
	}
L273:
	;
	goto L274
L274:
	;
	v1377 = int32(1)
	*(*uint8)(unsafe.Add(mBase, _c_F_SysLoggerMain[62])) = uint8(v1377)
	v1385 = int32(0)
	goto L364
L275:
	;
	v1008 = v1002
	v1011 = v26 + int32(48)
	v1016 = int32(1)
	goto L276
L276:
	;
	v1031 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1011))))
	if v1031 != 0 {
		goto L281
	} else {
		goto L282
	}
L277:
	;
	v1365 = int32(0)
	v1368 = v26 + int32(48)
	if base.B2i32(v1342 == v1365)|(base.B2i32(v1345 == v1368)|base.B2i32(v1342 <= v1365)) != 0 {
		v598 = v1342
		v608 = v787
		v609 = v788
		v610 = v789
		v620 = v812
		goto L156
	} else {
		goto L363
	}
L278:
	;
	goto L277
L279:
	;
	v1338 = v1317 + v1011
	v1339 = v1008 - v1317
	if int32(9) < v1339 {
		v1008 = v1339
		v1011 = v1338
		v1016 = v1323
		goto L276
	} else {
		goto L362
	}
L280:
	;
	F_write_stderr(m, int32(_a_F_SysLoggerMain_17), int32(0))
	mBase = m.M
	v1314 = m.ExcPending
	if v1314 != 0 {
		goto L4
	} else {
		goto L361
	}
L281:
	;
	v1253 = int32(1)
	goto L355
L282:
	;
	v1032 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1011)+1)))
	if v1032 != 0 {
		goto L281
	} else {
		goto L283
	}
L283:
	;
	v1033 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v1011)+2)))
	if base.Ui32(int32(4086)) < base.Ui32((v1033-int32(1))&int32(_a_F_SysLoggerMain_18)) {
		goto L281
	} else {
		goto L284
	}
L284:
	;
	v1040 = *(*int32)(unsafe.Add(mBase, uint32(v1011)+4))
	if v1040 == int32(0) {
		goto L281
	} else {
		goto L285
	}
L285:
	;
	v1043 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1011)+8)))
	v1046 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1043&int32(112))+uint32(_c_F_SysLoggerMain[64]))))
	if v1046 != int32(1) {
		goto L281
	} else {
		goto L286
	}
L286:
	;
	v1050 = v1033 + int32(9)
	if base.Ui32(v1008) < base.Ui32(v1050) {
		v1342 = v1008
		v1345 = v1011
		goto L278
	} else {
		goto L287
	}
L287:
	;
	if v1043&int32(16) != 0 {
		v1062 = int32(1)
		goto L288
	} else {
		goto L289
	}
L288:
	;
	v1063 = int32(0)
	v1065 = base.I32_rem_s(v1040, int32(256))
	v1067 = v1065 << (uint(int32(2)) % 32)
	v1070 = *(*int32)(unsafe.Add(mBase, uint32(v1067)+uint32(_c_F_SysLoggerMain[65])))
	if v1070 != 0 {
		goto L295
	} else {
		goto L296
	}
L289:
	;
	if v1043&int32(32) != 0 {
		v1062 = int32(8)
		goto L288
	} else {
		goto L290
	}
L290:
	;
	if v1043&int32(64) != 0 {
		goto L291
	} else {
		goto L292
	}
L291:
	;
	v1061 = int32(16)
	goto L293
L292:
	;
	v1061 = v1016
	goto L293
L293:
	;
	v1062 = v1061
	goto L288
L294:
	;
	if v1043&int32(1) == int32(0) {
		goto L306
	} else {
		goto L307
	}
L295:
	;
	v1071 = int32(0)
	v1072 = *(*int32)(unsafe.Add(mBase, uint32(v1070)+4))
	if v1072 <= v1071 {
		v1136 = v1071
		v1140 = v1063
		goto L294
	} else {
		goto L298
	}
L296:
	;
	v1116 = v1063
	goto L297
L297:
	;
	v1136 = int32(0)
	v1140 = v1116
	goto L294
L298:
	;
	v1075 = *(*int32)(unsafe.Add(mBase, uint32(v1070)+12))
	v1082 = v1063
	v1083 = int32(0)
	goto L299
L299:
	;
	v1103 = *(*int32)(unsafe.Add(mBase, uint32(v1075+v1083<<(uint(int32(2))%32))))
	v1104 = *(*int32)(unsafe.Add(mBase, uint32(v1103)))
	if v1104 == v1040 {
		v1136 = v1103
		v1140 = v1082
		goto L294
	} else {
		goto L301
	}
L300:
	;
	v1116 = v1107
	goto L297
L301:
	;
	if v1082|v1104 != 0 {
		goto L302
	} else {
		goto L303
	}
L302:
	;
	v1107 = v1082
	goto L304
L303:
	;
	v1107 = v1103
	goto L304
L304:
	;
	v1109 = v1083 + int32(1)
	if v1072 != v1109 {
		v1082 = v1107
		v1083 = v1109
		goto L299
	} else {
		goto L305
	}
L305:
	;
	goto L300
L306:
	;
	if v1136 != 0 {
		goto L309
	} else {
		goto L310
	}
L307:
	;
	goto L308
L308:
	;
	if v1136 != 0 {
		goto L320
	} else {
		goto L321
	}
L309:
	;
	F_appendBinaryStringInfo(m, v1136+int32(4), v1011+int32(9), v1033)
	mBase = m.M
	v1167 = m.ExcPending
	if v1167 != 0 {
		goto L4
	} else {
		goto L312
	}
L310:
	;
	goto L311
L311:
	;
	if v1140 == int32(0) {
		goto L313
	} else {
		goto L314
	}
L312:
	;
	v1317 = v1050
	v1323 = v1062
	goto L279
L313:
	;
	v1171 = F_palloc(m, int32(20))
	mBase = m.M
	v1172 = m.ExcPending
	if v1172 != 0 {
		goto L4
	} else {
		goto L316
	}
L314:
	;
	v1176 = v1140
	goto L315
L315:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v1176))) = v1040
	v1179 = v1176 + int32(4)
	F_initStringInfo(m, v1179)
	mBase = m.M
	v1181 = m.ExcPending
	if v1181 != 0 {
		goto L4
	} else {
		goto L318
	}
L316:
	;
	v1173 = F_lappend(m, v1070, v1171)
	mBase = m.M
	v1174 = m.ExcPending
	if v1174 != 0 {
		goto L4
	} else {
		goto L317
	}
L317:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v1067)+uint32(_c_F_SysLoggerMain[65]))) = v1173
	v1176 = v1171
	goto L315
L318:
	;
	F_appendBinaryStringInfo(m, v1179, v1011+int32(9), v1033)
	mBase = m.M
	v1185 = m.ExcPending
	if v1185 != 0 {
		goto L4
	} else {
		goto L319
	}
L319:
	;
	v1317 = v1050
	v1323 = v1062
	goto L279
L320:
	;
	F_appendBinaryStringInfo(m, v1136+int32(4), v1011+int32(9), v1033)
	mBase = m.M
	v1191 = m.ExcPending
	if v1191 != 0 {
		goto L4
	} else {
		goto L323
	}
L321:
	;
	goto L322
L322:
	;
	if v1062&int32(8) != 0 {
		goto L342
	} else {
		goto L343
	}
L323:
	;
	v1192 = *(*int32)(unsafe.Add(mBase, uint32(v1136)+8))
	v1193 = *(*int32)(unsafe.Add(mBase, uint32(v1136)+4))
	if v1062&int32(8) != 0 {
		goto L325
	} else {
		goto L326
	}
L324:
	;
	v1212 = F_fwrite(m, v1193, int32(1), v1192, v1209)
	mBase = m.M
	v1213 = m.ExcPending
	if v1213 != 0 {
		goto L4
	} else {
		goto L335
	}
L325:
	;
	v1197 = *(*int32)(unsafe.Add(mBase, _c_F_SysLoggerMain[46]))
	if v1197 != 0 {
		v1209 = v1197
		goto L324
	} else {
		goto L328
	}
L326:
	;
	goto L327
L327:
	;
	v1200 = *(*int32)(unsafe.Add(mBase, _c_F_SysLoggerMain[47]))
	v1202 = *(*int32)(unsafe.Add(mBase, _c_F_SysLoggerMain[61]))
	if v1200 != 0 {
		goto L329
	} else {
		goto L330
	}
L328:
	;
	goto L327
L329:
	;
	v1203 = v1200
	goto L331
L330:
	;
	v1203 = v1202
	goto L331
L331:
	;
	if int32(base.Ui32(v1062&int32(16))>>(uint(int32(4))%32)) != 0 {
		goto L332
	} else {
		goto L333
	}
L332:
	;
	v1208 = v1203
	goto L334
L333:
	;
	v1208 = v1202
	goto L334
L334:
	;
	v1209 = v1208
	goto L324
L335:
	;
	if v1212 != v1192 {
		goto L336
	} else {
		goto L337
	}
L336:
	;
	F_write_stderr(m, int32(_a_F_SysLoggerMain_17), int32(0))
	mBase = m.M
	v1218 = m.ExcPending
	if v1218 != 0 {
		goto L4
	} else {
		goto L339
	}
L337:
	;
	goto L338
L338:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v1136))) = int32(0)
	v1221 = *(*int32)(unsafe.Add(mBase, uint32(v1136)+4))
	F_pfree(m, v1221)
	mBase = m.M
	v1223 = m.ExcPending
	if v1223 != 0 {
		goto L4
	} else {
		goto L340
	}
L339:
	;
	goto L338
L340:
	;
	v1317 = v1050
	v1323 = v1062
	goto L279
L341:
	;
	v1244 = F_fwrite(m, v1011+int32(9), int32(1), v1033, v1240)
	mBase = m.M
	v1245 = m.ExcPending
	if v1245 != 0 {
		goto L4
	} else {
		goto L352
	}
L342:
	;
	v1227 = *(*int32)(unsafe.Add(mBase, _c_F_SysLoggerMain[46]))
	if v1227 != 0 {
		v1240 = v1227
		goto L341
	} else {
		goto L345
	}
L343:
	;
	goto L344
L344:
	;
	v1230 = *(*int32)(unsafe.Add(mBase, _c_F_SysLoggerMain[47]))
	v1232 = *(*int32)(unsafe.Add(mBase, _c_F_SysLoggerMain[61]))
	if v1230 != 0 {
		goto L346
	} else {
		goto L347
	}
L345:
	;
	goto L344
L346:
	;
	v1233 = v1230
	goto L348
L347:
	;
	v1233 = v1232
	goto L348
L348:
	;
	if int32(base.Ui32(v1062&int32(16))>>(uint(int32(4))%32)) != 0 {
		goto L349
	} else {
		goto L350
	}
L349:
	;
	v1238 = v1233
	goto L351
L350:
	;
	v1238 = v1232
	goto L351
L351:
	;
	v1240 = v1238
	goto L341
L352:
	;
	if v1244 == v1033 {
		v1317 = v1050
		v1323 = v1062
		goto L279
	} else {
		goto L353
	}
L353:
	;
	v1290 = v1050
	v1296 = v1062
	goto L280
L354:
	;
	v1284 = *(*int32)(unsafe.Add(mBase, _c_F_SysLoggerMain[61]))
	v1285 = F_fwrite(m, v1011, int32(1), v1281, v1284)
	mBase = m.M
	v1286 = m.ExcPending
	if v1286 != 0 {
		goto L4
	} else {
		goto L359
	}
L355:
	;
	v1275 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1253+v1011))))
	if v1275 == int32(0) {
		v1281 = v1253
		goto L354
	} else {
		goto L357
	}
L356:
	;
	v1281 = v1008
	goto L354
L357:
	;
	v1279 = v1253 + int32(1)
	if v1279 != v1008 {
		v1253 = v1279
		goto L355
	} else {
		goto L358
	}
L358:
	;
	goto L356
L359:
	;
	if v1285 == v1281 {
		v1317 = v1281
		v1323 = v1016
		goto L279
	} else {
		goto L360
	}
L360:
	;
	v1290 = v1281
	v1296 = v1016
	goto L280
L361:
	;
	v1317 = v1290
	v1323 = v1296
	goto L279
L362:
	;
	v1342 = v1339
	v1345 = v1338
	goto L278
L363:
	;
	base.MemoryCopy(m, v1368, v1345, v1342)
	v598 = v1342
	v608 = v787
	v609 = v788
	v610 = v789
	v620 = v812
	goto L156
L364:
	;
	v1404 = *(*int32)(unsafe.Add(mBase, uint32(v1385<<(uint(int32(2))%32))+uint32(_c_F_SysLoggerMain[65])))
	if v1404 == int32(0) {
		goto L366
	} else {
		goto L367
	}
L365:
	;
	v1489 = int32(0)
	if v598 <= v1489 {
		v1504 = v1489
		goto L251
	} else {
		goto L382
	}
L366:
	;
	v1486 = v1385 + int32(1)
	if v1486 != int32(256) {
		v1385 = v1486
		goto L364
	} else {
		goto L381
	}
L367:
	;
	v1407 = int32(0)
	v1408 = *(*int32)(unsafe.Add(mBase, uint32(v1404)+4))
	if v1408 <= v1407 {
		goto L366
	} else {
		goto L368
	}
L368:
	;
	v1411 = v1408
	v1414 = v1407
	goto L369
L369:
	;
	v1434 = *(*int32)(unsafe.Add(mBase, uint32(v1404)+12))
	v1438 = *(*int32)(unsafe.Add(mBase, uint32(v1434+v1414<<(uint(int32(2))%32))))
	v1439 = *(*int32)(unsafe.Add(mBase, uint32(v1438)))
	if v1439 != 0 {
		goto L371
	} else {
		goto L372
	}
L370:
	;
	goto L366
L371:
	;
	v1440 = *(*int32)(unsafe.Add(mBase, uint32(v1438)+4))
	v1442 = *(*int32)(unsafe.Add(mBase, uint32(v1438)+8))
	v1444 = *(*int32)(unsafe.Add(mBase, _c_F_SysLoggerMain[61]))
	v1445 = F_fwrite(m, v1440, int32(1), v1442, v1444)
	mBase = m.M
	v1446 = m.ExcPending
	if v1446 != 0 {
		goto L4
	} else {
		goto L374
	}
L372:
	;
	v1458 = v1411
	goto L373
L373:
	;
	v1460 = v1414 + int32(1)
	if v1460 < v1458 {
		v1411 = v1458
		v1414 = v1460
		goto L369
	} else {
		goto L380
	}
L374:
	;
	if v1445 != v1442 {
		goto L375
	} else {
		goto L376
	}
L375:
	;
	F_write_stderr(m, int32(_a_F_SysLoggerMain_17), int32(0))
	mBase = m.M
	v1451 = m.ExcPending
	if v1451 != 0 {
		goto L4
	} else {
		goto L378
	}
L376:
	;
	goto L377
L377:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v1438))) = int32(0)
	v1454 = *(*int32)(unsafe.Add(mBase, uint32(v1438)+4))
	F_pfree(m, v1454)
	mBase = m.M
	v1456 = m.ExcPending
	if v1456 != 0 {
		goto L4
	} else {
		goto L379
	}
L378:
	;
	goto L377
L379:
	;
	v1457 = *(*int32)(unsafe.Add(mBase, uint32(v1404)+4))
	v1458 = v1457
	goto L373
L380:
	;
	goto L370
L381:
	;
	goto L365
L382:
	;
	v1496 = *(*int32)(unsafe.Add(mBase, _c_F_SysLoggerMain[61]))
	v1497 = F_fwrite(m, v26+int32(48), int32(1), v598, v1496)
	mBase = m.M
	v1498 = m.ExcPending
	if v1498 != 0 {
		goto L4
	} else {
		goto L383
	}
L383:
	;
	if v1497 == v598 {
		v1504 = v1489
		goto L251
	} else {
		goto L384
	}
L384:
	;
	F_write_stderr(m, int32(_a_F_SysLoggerMain_17), int32(0))
	mBase = m.M
	v1503 = m.ExcPending
	if v1503 != 0 {
		goto L4
	} else {
		goto L385
	}
L385:
	;
	v1504 = v1489
	goto L251
L386:
	;
	goto L157
L387:
	;
	if v1533 != 0 {
		goto L388
	} else {
		goto L389
	}
L388:
	;
	F_errmsg_internal(m, int32(_a_F_SysLoggerMain_19), int32(0))
	mBase = m.M
	v1538 = m.ExcPending
	if v1538 != 0 {
		goto L4
	} else {
		goto L391
	}
L389:
	;
	goto L390
L390:
	;
	F_proc_exit(m, int32(0))
	mBase = m.M
	v1546 = m.ExcPending
	if v1546 != 0 {
		goto L4
	} else {
		goto L393
	}
L391:
	;
	F_errfinish(m, int32(_a_F_SysLoggerMain_15), int32(575), int32(_a_F_SysLoggerMain_16))
	mBase = m.M
	v1543 = m.ExcPending
	if v1543 != 0 {
		goto L4
	} else {
		goto L392
	}
L392:
	;
	goto L390
L393:
	;
	base.Wasm_trap_unreachable()
	for {
	}
}
