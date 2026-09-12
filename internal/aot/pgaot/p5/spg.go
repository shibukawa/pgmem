package p5

import (
	base "github.com/shibukawa/pgmem/internal/aot/pgaot/base"
	"unsafe"
)

func F_spgUpdateNodeLink(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) {
	mBase := m.M
	_ = mBase
	var v3 int32
	_ = v3
	var v4 int32
	_ = v4
	var v8 int32
	_ = v8
	var v10 int32
	_ = v10
	var v12 int32
	_ = v12
	var v16 int32
	_ = v16
	var v23 int32
	_ = v23
	var v27 int32
	_ = v27
	var v31 int32
	_ = v31
	var v36 int32
	_ = v36
	var v48 int32
	_ = v48
	var v52 int32
	_ = v52
	var v57 int32
	_ = v57
	var v60 int32
	_ = v60
	var v61 int32
	_ = v61
	v3 = l2
	v4 = l3
	v8 = m.G0
	v10 = v8 - int32(16)
	m.G0 = v10
	v12 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v16 = int32(base.Ui32(v12)>>(uint(int32(3))%32)) & int32(8191)
	if v16 != 0 {
		goto L2
	} else {
		goto L3
	}
L1:
	;
	*(*uint16)(unsafe.Add(mBase, uint32(v23)+4)) = uint16(v4)
	*(*uint16)(unsafe.Add(mBase, uint32(v23)+2)) = uint16(v3)
	v60 = int32(16)
	v61 = int32(base.Ui32(v3) >> (uint(v60) % 32))
	*(*uint16)(unsafe.Add(mBase, uint32(v23))) = uint16(v61)
	m.G0 = v10 + v60
	return
L2:
	;
	v23 = l0 + int32(base.Ui32(v12)>>(uint(int32(16))%32)) + int32(8)
	v27 = int32(0)
	goto L5
L3:
	;
	goto L4
L4:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v48 = m.ExcPending
	if v48 != 0 {
		goto L9
	} else {
		goto L10
	}
L5:
	;
	if l1 == v27 {
		goto L1
	} else {
		goto L7
	}
L6:
	;
	goto L4
L7:
	;
	v31 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v23)+6)))
	v36 = v27 + int32(1)
	if v36 != v16 {
		v23 = v23 + v31&int32(8191)
		v27 = v36
		goto L5
	} else {
		goto L8
	}
L8:
	;
	goto L6
L9:
	;
	return
L10:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v10))) = l1
	F_errmsg_internal(m, int32(402048), v10)
	mBase = m.M
	v52 = m.ExcPending
	if v52 != 0 {
		goto L9
	} else {
		goto L11
	}
L11:
	;
	F_errfinish(m, int32(515200), int32(68), int32(330515))
	mBase = m.M
	v57 = m.ExcPending
	if v57 != 0 {
		goto L9
	} else {
		goto L12
	}
L12:
	;
	base.Wasm_trap_unreachable()
	for {
	}
}
func F_spg_kd_config(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v2 int32
	_ = v2
	var v3 int32
	_ = v3
	v2 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
	v3 = int32(1)
	*(*uint16)(unsafe.Add(mBase, uint32(v2)+12)) = uint16(v3)
	*(*int64)(unsafe.Add(mBase, uint32(v2))) = int64(9783935500989)
	return int32(0)
}
func F_spg_mask(m *base.Module, l0 int32, l1 int32) {
	mBase := m.M
	_ = mBase
	var v3 int32
	_ = v3
	var v9 int32
	_ = v9
	var v11 int32
	_ = v11
	var v13 int32
	_ = v13
	var v17 int32
	_ = v17
	v3 = int32(0)
	*(*uint16)(unsafe.Add(mBase, uint32(l0)+8)) = uint16(v3)
	*(*int64)(unsafe.Add(mBase, uint32(l0))) = int64(0)
	*(*int32)(unsafe.Add(mBase, uint32(l0)+20)) = int32(0)
	v9 = int32(*(*uint16)(unsafe.Add(mBase, uint32(l0)+10)))
	v11 = v9 & int32(65528)
	*(*uint16)(unsafe.Add(mBase, uint32(l0)+10)) = uint16(v11)
	v13 = int32(*(*uint16)(unsafe.Add(mBase, uint32(l0)+12)))
	if base.Ui32(int32(24)) <= base.Ui32(v13) {
		F_mask_unused_space(m, l0)
		mBase = m.M
		v17 = m.ExcPending
		if v17 != 0 {
			return
		} else {
			return
		}
	} else {
		return
	}
}
func F_spg_redo(m *base.Module, l0 int32) {
	mBase := m.M
	_ = mBase
	var v2 int32
	_ = v2
	var v17 int32
	_ = v17
	var v19 int32
	_ = v19
	var v21 int32
	_ = v21
	var v22 int32
	_ = v22
	var v23 int32
	_ = v23
	var v24 int32
	_ = v24
	var v27 int32
	_ = v27
	var v30 int32
	_ = v30
	var v35 int64
	_ = v35
	var v36 int32
	_ = v36
	var v38 int32
	_ = v38
	var v39 int32
	_ = v39
	var v40 int32
	_ = v40
	var v44 int32
	_ = v44
	var v45 int32
	_ = v45
	var v49 int32
	_ = v49
	var v50 int32
	_ = v50
	var v54 int32
	_ = v54
	var v60 int32
	_ = v60
	var v62 int32
	_ = v62
	var v68 int32
	_ = v68
	var v72 int32
	_ = v72
	var v73 int32
	_ = v73
	var v74 int32
	_ = v74
	var v80 int32
	_ = v80
	var v81 int32
	_ = v81
	var v82 int32
	_ = v82
	var v83 int32
	_ = v83
	var v87 int32
	_ = v87
	var v93 int32
	_ = v93
	var v95 int32
	_ = v95
	var v101 int32
	_ = v101
	var v102 int32
	_ = v102
	var v103 int32
	_ = v103
	var v108 int32
	_ = v108
	var v109 int32
	_ = v109
	var v115 int32
	_ = v115
	var v118 int32
	_ = v118
	var v119 int32
	_ = v119
	var v122 int32
	_ = v122
	var v125 int32
	_ = v125
	var v128 int32
	_ = v128
	var v130 int32
	_ = v130
	var v131 int32
	_ = v131
	var v133 int32
	_ = v133
	var v134 int32
	_ = v134
	var v135 int32
	_ = v135
	var v142 int32
	_ = v142
	var v144 int32
	_ = v144
	var v149 int32
	_ = v149
	var v151 int32
	_ = v151
	var v152 int32
	_ = v152
	var v158 int32
	_ = v158
	var v159 int32
	_ = v159
	var v162 int32
	_ = v162
	var v168 int32
	_ = v168
	var v169 int32
	_ = v169
	var v173 int32
	_ = v173
	var v179 int32
	_ = v179
	var v181 int32
	_ = v181
	var v187 int32
	_ = v187
	var v188 int32
	_ = v188
	var v192 int32
	_ = v192
	var v196 int32
	_ = v196
	var v197 int32
	_ = v197
	var v198 int32
	_ = v198
	var v200 int32
	_ = v200
	var v204 int32
	_ = v204
	var v206 int32
	_ = v206
	var v208 int32
	_ = v208
	var v212 int32
	_ = v212
	var v213 int64
	_ = v213
	var v214 int32
	_ = v214
	var v216 int32
	_ = v216
	var v221 int32
	_ = v221
	var v222 int32
	_ = v222
	var v223 int32
	_ = v223
	var v229 int32
	_ = v229
	var v233 int32
	_ = v233
	var v234 int32
	_ = v234
	var v237 int32
	_ = v237
	var v238 int32
	_ = v238
	var v239 int32
	_ = v239
	var v241 int32
	_ = v241
	var v245 int32
	_ = v245
	var v246 int32
	_ = v246
	var v247 int32
	_ = v247
	var v251 int32
	_ = v251
	var v252 int32
	_ = v252
	var v256 int32
	_ = v256
	var v257 int32
	_ = v257
	var v261 int32
	_ = v261
	var v267 int32
	_ = v267
	var v269 int32
	_ = v269
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
	var v288 int32
	_ = v288
	var v292 int32
	_ = v292
	var v296 int32
	_ = v296
	var v302 int32
	_ = v302
	var v304 int32
	_ = v304
	var v310 int32
	_ = v310
	var v314 int32
	_ = v314
	var v316 int32
	_ = v316
	var v329 int32
	_ = v329
	var v331 int32
	_ = v331
	var v335 int32
	_ = v335
	var v337 int32
	_ = v337
	var v340 int32
	_ = v340
	var v345 int32
	_ = v345
	var v347 int32
	_ = v347
	var v364 int32
	_ = v364
	var v366 int32
	_ = v366
	var v370 int32
	_ = v370
	var v371 int32
	_ = v371
	var v376 int32
	_ = v376
	var v380 int32
	_ = v380
	var v386 int32
	_ = v386
	var v388 int32
	_ = v388
	var v394 int32
	_ = v394
	var v395 int32
	_ = v395
	var v397 int32
	_ = v397
	var v400 int32
	_ = v400
	var v402 int32
	_ = v402
	var v408 int32
	_ = v408
	var v410 int32
	_ = v410
	var v414 int32
	_ = v414
	var v416 int32
	_ = v416
	var v418 int32
	_ = v418
	var v420 int32
	_ = v420
	var v424 int32
	_ = v424
	var v425 int32
	_ = v425
	var v428 int32
	_ = v428
	var v432 int32
	_ = v432
	var v438 int32
	_ = v438
	var v440 int32
	_ = v440
	var v446 int32
	_ = v446
	var v447 int32
	_ = v447
	var v448 int32
	_ = v448
	var v451 int32
	_ = v451
	var v455 int32
	_ = v455
	var v456 int32
	_ = v456
	var v462 int32
	_ = v462
	var v464 int32
	_ = v464
	var v468 int32
	_ = v468
	var v470 int32
	_ = v470
	var v472 int32
	_ = v472
	var v476 int32
	_ = v476
	var v477 int64
	_ = v477
	var v478 int32
	_ = v478
	var v479 int32
	_ = v479
	var v480 int32
	_ = v480
	var v481 int32
	_ = v481
	var v487 int32
	_ = v487
	var v491 int32
	_ = v491
	var v492 int32
	_ = v492
	var v495 int32
	_ = v495
	var v496 int32
	_ = v496
	var v497 int32
	_ = v497
	var v500 int32
	_ = v500
	var v504 int32
	_ = v504
	var v505 int32
	_ = v505
	var v508 int32
	_ = v508
	var v512 int32
	_ = v512
	var v518 int32
	_ = v518
	var v520 int32
	_ = v520
	var v526 int32
	_ = v526
	var v527 int32
	_ = v527
	var v529 int32
	_ = v529
	var v530 int32
	_ = v530
	var v532 int32
	_ = v532
	var v533 int32
	_ = v533
	var v534 int32
	_ = v534
	var v539 int32
	_ = v539
	var v541 int32
	_ = v541
	var v543 int32
	_ = v543
	var v547 int32
	_ = v547
	var v548 int32
	_ = v548
	var v554 int32
	_ = v554
	var v556 int32
	_ = v556
	var v561 int32
	_ = v561
	var v562 int32
	_ = v562
	var v566 int32
	_ = v566
	var v567 int32
	_ = v567
	var v569 int32
	_ = v569
	var v573 int32
	_ = v573
	var v579 int32
	_ = v579
	var v581 int32
	_ = v581
	var v587 int32
	_ = v587
	var v591 int32
	_ = v591
	var v592 int32
	_ = v592
	var v593 int32
	_ = v593
	var v599 int32
	_ = v599
	var v600 int32
	_ = v600
	var v602 int32
	_ = v602
	var v606 int32
	_ = v606
	var v612 int32
	_ = v612
	var v614 int32
	_ = v614
	var v620 int32
	_ = v620
	var v621 int32
	_ = v621
	var v623 int32
	_ = v623
	var v624 int32
	_ = v624
	var v627 int32
	_ = v627
	var v631 int32
	_ = v631
	var v635 int32
	_ = v635
	var v636 int32
	_ = v636
	var v637 int32
	_ = v637
	var v639 int32
	_ = v639
	var v643 int32
	_ = v643
	var v645 int32
	_ = v645
	var v647 int32
	_ = v647
	var v649 int32
	_ = v649
	var v653 int32
	_ = v653
	var v654 int32
	_ = v654
	var v657 int32
	_ = v657
	var v661 int32
	_ = v661
	var v667 int32
	_ = v667
	var v669 int32
	_ = v669
	var v675 int32
	_ = v675
	var v676 int32
	_ = v676
	var v685 int32
	_ = v685
	var v691 int32
	_ = v691
	var v693 int32
	_ = v693
	var v704 int32
	_ = v704
	var v711 int32
	_ = v711
	var v713 int32
	_ = v713
	var v714 int32
	_ = v714
	var v716 int32
	_ = v716
	var v722 int32
	_ = v722
	var v724 int32
	_ = v724
	var v731 int32
	_ = v731
	var v733 int32
	_ = v733
	var v741 int32
	_ = v741
	var v742 int32
	_ = v742
	var v744 int32
	_ = v744
	var v745 int32
	_ = v745
	var v748 int32
	_ = v748
	var v750 int32
	_ = v750
	var v751 int32
	_ = v751
	var v752 int32
	_ = v752
	var v754 int32
	_ = v754
	var v758 int32
	_ = v758
	var v759 int32
	_ = v759
	var v760 int32
	_ = v760
	var v761 int32
	_ = v761
	var v763 int32
	_ = v763
	var v765 int32
	_ = v765
	var v768 int32
	_ = v768
	var v772 int32
	_ = v772
	var v776 int32
	_ = v776
	var v777 int32
	_ = v777
	var v778 int32
	_ = v778
	var v780 int32
	_ = v780
	var v784 int32
	_ = v784
	var v786 int32
	_ = v786
	var v789 int32
	_ = v789
	var v791 int32
	_ = v791
	var v792 int32
	_ = v792
	var v798 int32
	_ = v798
	var v799 int32
	_ = v799
	var v802 int32
	_ = v802
	var v806 int32
	_ = v806
	var v812 int32
	_ = v812
	var v814 int32
	_ = v814
	var v820 int32
	_ = v820
	var v821 int32
	_ = v821
	var v825 int32
	_ = v825
	var v829 int32
	_ = v829
	var v830 int32
	_ = v830
	var v831 int32
	_ = v831
	var v833 int32
	_ = v833
	var v837 int32
	_ = v837
	var v839 int32
	_ = v839
	var v841 int32
	_ = v841
	var v845 int32
	_ = v845
	var v846 int64
	_ = v846
	var v847 int32
	_ = v847
	var v849 int32
	_ = v849
	var v850 int32
	_ = v850
	var v851 int32
	_ = v851
	var v852 int32
	_ = v852
	var v853 int32
	_ = v853
	var v854 int32
	_ = v854
	var v858 int32
	_ = v858
	var v859 int32
	_ = v859
	var v861 int32
	_ = v861
	var v865 int32
	_ = v865
	var v871 int32
	_ = v871
	var v873 int32
	_ = v873
	var v879 int32
	_ = v879
	var v883 int32
	_ = v883
	var v884 int32
	_ = v884
	var v885 int32
	_ = v885
	var v891 int32
	_ = v891
	var v892 int32
	_ = v892
	var v893 int32
	_ = v893
	var v894 int32
	_ = v894
	var v898 int32
	_ = v898
	var v904 int32
	_ = v904
	var v906 int32
	_ = v906
	var v912 int32
	_ = v912
	var v913 int32
	_ = v913
	var v915 int32
	_ = v915
	var v919 int32
	_ = v919
	var v921 int32
	_ = v921
	var v923 int32
	_ = v923
	var v927 int32
	_ = v927
	var v932 int32
	_ = v932
	var v933 int32
	_ = v933
	var v936 int32
	_ = v936
	var v940 int32
	_ = v940
	var v946 int32
	_ = v946
	var v948 int32
	_ = v948
	var v954 int32
	_ = v954
	var v955 int32
	_ = v955
	var v957 int32
	_ = v957
	var v958 int32
	_ = v958
	var v960 int32
	_ = v960
	var v961 int32
	_ = v961
	var v962 int32
	_ = v962
	var v964 int32
	_ = v964
	var v967 int32
	_ = v967
	var v969 int32
	_ = v969
	var v973 int32
	_ = v973
	var v975 int32
	_ = v975
	var v977 int32
	_ = v977
	var v981 int32
	_ = v981
	var v982 int64
	_ = v982
	var v983 int32
	_ = v983
	var v985 int32
	_ = v985
	var v990 int32
	_ = v990
	var v991 int32
	_ = v991
	var v992 int32
	_ = v992
	var v998 int32
	_ = v998
	var v1002 int32
	_ = v1002
	var v1003 int32
	_ = v1003
	var v1006 int32
	_ = v1006
	var v1007 int32
	_ = v1007
	var v1008 int32
	_ = v1008
	var v1010 int32
	_ = v1010
	var v1011 int32
	_ = v1011
	var v1014 int32
	_ = v1014
	var v1015 int32
	_ = v1015
	var v1016 int32
	_ = v1016
	var v1017 int32
	_ = v1017
	var v1022 int32
	_ = v1022
	var v1026 int32
	_ = v1026
	var v1027 int32
	_ = v1027
	var v1032 int32
	_ = v1032
	var v1038 int32
	_ = v1038
	var v1040 int32
	_ = v1040
	var v1046 int32
	_ = v1046
	var v1049 int32
	_ = v1049
	var v1050 int32
	_ = v1050
	var v1054 int32
	_ = v1054
	var v1060 int32
	_ = v1060
	var v1062 int32
	_ = v1062
	var v1068 int32
	_ = v1068
	var v1072 int32
	_ = v1072
	var v1073 int32
	_ = v1073
	var v1074 int32
	_ = v1074
	var v1080 int32
	_ = v1080
	var v1081 int32
	_ = v1081
	var v1082 int32
	_ = v1082
	var v1086 int32
	_ = v1086
	var v1092 int32
	_ = v1092
	var v1094 int32
	_ = v1094
	var v1100 int32
	_ = v1100
	var v1101 int32
	_ = v1101
	var v1102 int32
	_ = v1102
	var v1109 int32
	_ = v1109
	var v1110 int32
	_ = v1110
	var v1112 int32
	_ = v1112
	var v1115 int32
	_ = v1115
	var v1120 int32
	_ = v1120
	var v1122 int32
	_ = v1122
	var v1123 int32
	_ = v1123
	var v1124 int32
	_ = v1124
	var v1127 int32
	_ = v1127
	var v1130 int32
	_ = v1130
	var v1134 int32
	_ = v1134
	var v1135 int32
	_ = v1135
	var v1140 int32
	_ = v1140
	var v1146 int32
	_ = v1146
	var v1148 int32
	_ = v1148
	var v1154 int32
	_ = v1154
	var v1157 int32
	_ = v1157
	var v1158 int32
	_ = v1158
	var v1162 int32
	_ = v1162
	var v1168 int32
	_ = v1168
	var v1170 int32
	_ = v1170
	var v1176 int32
	_ = v1176
	var v1180 int32
	_ = v1180
	var v1181 int32
	_ = v1181
	var v1182 int32
	_ = v1182
	var v1188 int32
	_ = v1188
	var v1189 int32
	_ = v1189
	var v1190 int32
	_ = v1190
	var v1194 int32
	_ = v1194
	var v1200 int32
	_ = v1200
	var v1202 int32
	_ = v1202
	var v1209 int32
	_ = v1209
	var v1210 int32
	_ = v1210
	var v1214 int32
	_ = v1214
	var v1216 int32
	_ = v1216
	var v1218 int32
	_ = v1218
	var v1229 int32
	_ = v1229
	var v1231 int32
	_ = v1231
	var v1233 int32
	_ = v1233
	var v1234 int32
	_ = v1234
	var v1238 int32
	_ = v1238
	var v1240 int32
	_ = v1240
	var v1241 int32
	_ = v1241
	var v1242 int32
	_ = v1242
	var v1245 int32
	_ = v1245
	var v1268 int32
	_ = v1268
	var v1270 int32
	_ = v1270
	var v1274 int32
	_ = v1274
	var v1276 int32
	_ = v1276
	var v1277 int32
	_ = v1277
	var v1281 int32
	_ = v1281
	var v1282 int32
	_ = v1282
	var v1286 int32
	_ = v1286
	var v1287 int32
	_ = v1287
	var v1291 int32
	_ = v1291
	var v1297 int32
	_ = v1297
	var v1299 int32
	_ = v1299
	var v1305 int32
	_ = v1305
	var v1309 int32
	_ = v1309
	var v1310 int32
	_ = v1310
	var v1311 int32
	_ = v1311
	var v1317 int32
	_ = v1317
	var v1318 int32
	_ = v1318
	var v1319 int32
	_ = v1319
	var v1320 int32
	_ = v1320
	var v1324 int32
	_ = v1324
	var v1330 int32
	_ = v1330
	var v1332 int32
	_ = v1332
	var v1338 int32
	_ = v1338
	var v1339 int32
	_ = v1339
	var v1341 int32
	_ = v1341
	var v1342 int32
	_ = v1342
	var v1345 int32
	_ = v1345
	var v1349 int32
	_ = v1349
	var v1353 int32
	_ = v1353
	var v1354 int32
	_ = v1354
	var v1355 int32
	_ = v1355
	var v1357 int32
	_ = v1357
	var v1361 int32
	_ = v1361
	var v1363 int32
	_ = v1363
	var v1365 int32
	_ = v1365
	var v1367 int32
	_ = v1367
	var v1368 int32
	_ = v1368
	var v1370 int32
	_ = v1370
	var v1371 int32
	_ = v1371
	var v1373 int32
	_ = v1373
	var v1374 int32
	_ = v1374
	var v1375 int32
	_ = v1375
	var v1378 int32
	_ = v1378
	var v1384 int32
	_ = v1384
	var v1385 int32
	_ = v1385
	var v1388 int32
	_ = v1388
	var v1392 int32
	_ = v1392
	var v1398 int32
	_ = v1398
	var v1400 int32
	_ = v1400
	var v1406 int32
	_ = v1406
	var v1407 int32
	_ = v1407
	var v1411 int32
	_ = v1411
	var v1415 int32
	_ = v1415
	var v1416 int32
	_ = v1416
	var v1417 int32
	_ = v1417
	var v1419 int32
	_ = v1419
	var v1423 int32
	_ = v1423
	var v1425 int32
	_ = v1425
	var v1427 int32
	_ = v1427
	var v1431 int32
	_ = v1431
	var v1432 int64
	_ = v1432
	var v1433 int32
	_ = v1433
	var v1434 int32
	_ = v1434
	var v1435 int32
	_ = v1435
	var v1441 int32
	_ = v1441
	var v1445 int32
	_ = v1445
	var v1446 int32
	_ = v1446
	var v1448 int32
	_ = v1448
	var v1449 int32
	_ = v1449
	var v1450 int32
	_ = v1450
	var v1451 int32
	_ = v1451
	var v1455 int32
	_ = v1455
	var v1456 int32
	_ = v1456
	var v1460 int32
	_ = v1460
	var v1461 int32
	_ = v1461
	var v1463 int32
	_ = v1463
	var v1466 int32
	_ = v1466
	var v1468 int32
	_ = v1468
	var v1469 int32
	_ = v1469
	var v1472 int32
	_ = v1472
	var v1476 int32
	_ = v1476
	var v1482 int32
	_ = v1482
	var v1484 int32
	_ = v1484
	var v1490 int32
	_ = v1490
	var v1491 int32
	_ = v1491
	var v1492 int32
	_ = v1492
	var v1497 int32
	_ = v1497
	var v1500 int32
	_ = v1500
	var v1501 int32
	_ = v1501
	var v1506 int32
	_ = v1506
	var v1508 int32
	_ = v1508
	var v1510 int32
	_ = v1510
	var v1515 int32
	_ = v1515
	var v1528 int32
	_ = v1528
	var v1529 int32
	_ = v1529
	var v1531 int32
	_ = v1531
	var v1532 int32
	_ = v1532
	var v1535 int32
	_ = v1535
	var v1536 int32
	_ = v1536
	var v1537 int32
	_ = v1537
	var v1539 int32
	_ = v1539
	var v1544 int32
	_ = v1544
	var v1545 int32
	_ = v1545
	var v1549 int32
	_ = v1549
	var v1550 int32
	_ = v1550
	var v1552 int32
	_ = v1552
	var v1570 int32
	_ = v1570
	var v1575 int32
	_ = v1575
	var v1576 int32
	_ = v1576
	var v1577 int32
	_ = v1577
	var v1589 int32
	_ = v1589
	var v1602 int32
	_ = v1602
	var v1603 int32
	_ = v1603
	var v1605 int32
	_ = v1605
	var v1611 int32
	_ = v1611
	var v1614 int32
	_ = v1614
	var v1616 int32
	_ = v1616
	var v1619 int32
	_ = v1619
	var v1622 int32
	_ = v1622
	var v1625 int32
	_ = v1625
	var v1626 int32
	_ = v1626
	var v1647 int32
	_ = v1647
	var v1649 int32
	_ = v1649
	var v1666 int32
	_ = v1666
	var v1670 int32
	_ = v1670
	var v1671 int32
	_ = v1671
	var v1672 int64
	_ = v1672
	var v1676 int32
	_ = v1676
	var v1677 int32
	_ = v1677
	var v1682 int32
	_ = v1682
	var v1686 int32
	_ = v1686
	var v1692 int32
	_ = v1692
	var v1694 int32
	_ = v1694
	var v1700 int32
	_ = v1700
	var v1701 int32
	_ = v1701
	var v1703 int32
	_ = v1703
	var v1707 int32
	_ = v1707
	var v1709 int32
	_ = v1709
	var v1712 int32
	_ = v1712
	var v1716 int32
	_ = v1716
	var v1717 int32
	_ = v1717
	var v1718 int64
	_ = v1718
	var v1720 int32
	_ = v1720
	var v1723 int32
	_ = v1723
	var v1729 int32
	_ = v1729
	var v1730 int32
	_ = v1730
	var v1731 int32
	_ = v1731
	var v1732 int32
	_ = v1732
	var v1734 int64
	_ = v1734
	var v1739 int32
	_ = v1739
	var v1745 int32
	_ = v1745
	var v1746 int32
	_ = v1746
	var v1749 int32
	_ = v1749
	var v1753 int32
	_ = v1753
	var v1759 int32
	_ = v1759
	var v1761 int32
	_ = v1761
	var v1767 int32
	_ = v1767
	var v1768 int32
	_ = v1768
	var v1769 int32
	_ = v1769
	var v1770 int32
	_ = v1770
	var v1782 int32
	_ = v1782
	var v1795 int32
	_ = v1795
	var v1798 int32
	_ = v1798
	var v1804 int32
	_ = v1804
	var v1807 int32
	_ = v1807
	var v1808 int32
	_ = v1808
	var v1812 int32
	_ = v1812
	var v1817 int32
	_ = v1817
	var v1818 int32
	_ = v1818
	var v1821 int32
	_ = v1821
	var v1836 int32
	_ = v1836
	var v1837 int32
	_ = v1837
	var v1839 int32
	_ = v1839
	var v1840 int32
	_ = v1840
	var v1841 int32
	_ = v1841
	var v1843 int32
	_ = v1843
	var v1844 int32
	_ = v1844
	var v1845 int32
	_ = v1845
	var v1855 int32
	_ = v1855
	var v1858 int32
	_ = v1858
	var v1859 int32
	_ = v1859
	var v1860 int32
	_ = v1860
	var v1862 int32
	_ = v1862
	var v1866 int32
	_ = v1866
	var v1868 int32
	_ = v1868
	var v1872 int32
	_ = v1872
	var v1883 int32
	_ = v1883
	var v1885 int32
	_ = v1885
	var v1890 int32
	_ = v1890
	var v1892 int32
	_ = v1892
	var v1895 int32
	_ = v1895
	var v1913 int32
	_ = v1913
	var v1928 int32
	_ = v1928
	var v1930 int32
	_ = v1930
	var v1935 int32
	_ = v1935
	var v1936 int32
	_ = v1936
	var v1943 int32
	_ = v1943
	var v1944 int32
	_ = v1944
	var v1951 int32
	_ = v1951
	var v1952 int32
	_ = v1952
	var v1977 int32
	_ = v1977
	var v1979 int32
	_ = v1979
	var v1994 int32
	_ = v1994
	var v1997 int32
	_ = v1997
	var v1998 int32
	_ = v1998
	var v2001 int32
	_ = v2001
	var v2003 int32
	_ = v2003
	var v2023 int32
	_ = v2023
	var v2025 int32
	_ = v2025
	var v2042 int32
	_ = v2042
	var v2046 int32
	_ = v2046
	var v2066 int32
	_ = v2066
	var v2068 int32
	_ = v2068
	var v2075 int32
	_ = v2075
	var v2081 int32
	_ = v2081
	var v2086 int32
	_ = v2086
	var v2090 int32
	_ = v2090
	var v2096 int32
	_ = v2096
	var v2101 int32
	_ = v2101
	var v2105 int32
	_ = v2105
	var v2106 int32
	_ = v2106
	var v2114 int32
	_ = v2114
	var v2119 int32
	_ = v2119
	var v2123 int32
	_ = v2123
	var v2129 int32
	_ = v2129
	var v2134 int32
	_ = v2134
	var v2138 int32
	_ = v2138
	var v2142 int32
	_ = v2142
	var v2147 int32
	_ = v2147
	v2 = int32(0)
	v17 = m.G0
	v19 = v17 - int32(192)
	m.G0 = v19
	v21 = int32(4554128)
	v22 = *(*int32)(unsafe.Add(mBase, _consts[10]))
	v23 = *(*int32)(unsafe.Add(mBase, uint32(l0)+96))
	v24 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v23)+48)))
	v27 = *(*int32)(unsafe.Add(mBase, _consts[67]))
	*(*int32)(unsafe.Add(mBase, _consts[10])) = v27
	v30 = v24 & int32(240)
	switch int32(base.Ui32(v30-int32(16)) >> (uint(int32(4)) % 32)) {
	case 0:
		goto L14
	case 1:
		goto L13
	case 2:
		goto L12
	case 3:
		goto L11
	case 4:
		goto L10
	case 5:
		goto L9
	case 6:
		goto L8
	case 7:
		goto L7
	default:
		goto L1
	}
L1:
	;
	F_errstart_cold(m, int32(23), int32(0))
	mBase = m.M
	v2138 = m.ExcPending
	if v2138 != 0 {
		goto L20
	} else {
		goto L501
	}
L2:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v2123 = m.ExcPending
	if v2123 != 0 {
		goto L20
	} else {
		goto L498
	}
L3:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v2105 = m.ExcPending
	if v2105 != 0 {
		goto L20
	} else {
		goto L495
	}
L4:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v2090 = m.ExcPending
	if v2090 != 0 {
		goto L20
	} else {
		goto L492
	}
L5:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v2075 = m.ExcPending
	if v2075 != 0 {
		goto L20
	} else {
		goto L489
	}
L6:
	;
	*(*int32)(unsafe.Add(mBase, _consts[10])) = v22
	v2066 = *(*int32)(unsafe.Add(mBase, _consts[67]))
	F_MemoryContextReset(m, v2066)
	mBase = m.M
	v2068 = m.ExcPending
	if v2068 != 0 {
		goto L20
	} else {
		goto L488
	}
L7:
	;
	v1717 = *(*int32)(unsafe.Add(mBase, uint32(v23)+64))
	v1718 = *(*int64)(unsafe.Add(mBase, uint32(l0)+40))
	v1720 = *(*int32)(unsafe.Add(mBase, _consts[33]))
	if base.Ui32(int32(2)) <= base.Ui32(v1720) {
		goto L441
	} else {
		goto L442
	}
L8:
	;
	v1671 = *(*int32)(unsafe.Add(mBase, uint32(v23)+64))
	v1672 = *(*int64)(unsafe.Add(mBase, uint32(l0)+40))
	v1676 = F_XLogReadBufferForRedo(m, l0, int32(0), v19+int32(108))
	mBase = m.M
	v1677 = m.ExcPending
	if v1677 != 0 {
		goto L20
	} else {
		goto L429
	}
L9:
	;
	v1432 = *(*int64)(unsafe.Add(mBase, uint32(l0)+40))
	v1433 = *(*int32)(unsafe.Add(mBase, uint32(v23)+64))
	v1434 = *(*int32)(unsafe.Add(mBase, uint32(v1433)+8))
	v1435 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1433)+12)))
	v1441 = F__emscripten_memset_bulkmem(m, v19+int32(108), base.I32_extend8_s(int32(0)), int32(84))
	mBase = m.M
	goto L401
L10:
	;
	v982 = *(*int64)(unsafe.Add(mBase, uint32(l0)+40))
	v983 = *(*int32)(unsafe.Add(mBase, uint32(v23)+64))
	v985 = int32(0)
	F_XLogRecGetBlockTag(m, l0, int32(2), v985, v985, v19+int32(92))
	mBase = m.M
	v990 = m.ExcPending
	if v990 != 0 {
		goto L20
	} else {
		goto L267
	}
L11:
	;
	v846 = *(*int64)(unsafe.Add(mBase, uint32(l0)+40))
	v847 = *(*int32)(unsafe.Add(mBase, uint32(v23)+64))
	v849 = v847 + int32(6)
	v850 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v847)+10)))
	v851 = v849 + v850
	v852 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v851)+4)))
	v853 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v847)+5)))
	if v853 != 0 {
		goto L226
	} else {
		goto L227
	}
L12:
	;
	v477 = *(*int64)(unsafe.Add(mBase, uint32(l0)+40))
	v478 = *(*int32)(unsafe.Add(mBase, uint32(v23)+64))
	v479 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v478)+24)))
	v480 = *(*int32)(unsafe.Add(mBase, uint32(v478)+12))
	v481 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v478)+16)))
	v487 = F__emscripten_memset_bulkmem(m, v19+int32(108), base.I32_extend8_s(int32(0)), int32(84))
	mBase = m.M
	goto L128
L13:
	;
	v213 = *(*int64)(unsafe.Add(mBase, uint32(l0)+40))
	v214 = *(*int32)(unsafe.Add(mBase, uint32(v23)+64))
	v216 = int32(0)
	F_XLogRecGetBlockTag(m, l0, int32(1), v216, v216, v19+int32(100))
	mBase = m.M
	v221 = m.ExcPending
	if v221 != 0 {
		goto L20
	} else {
		goto L64
	}
L14:
	;
	v35 = *(*int64)(unsafe.Add(mBase, uint32(l0)+40))
	v36 = *(*int32)(unsafe.Add(mBase, uint32(v23)+64))
	v38 = v36 + int32(10)
	v39 = *(*int32)(unsafe.Add(mBase, uint32(v38)))
	v40 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v36))))
	if v40 == int32(1) {
		goto L17
	} else {
		goto L18
	}
L15:
	;
	v149 = *(*int32)(unsafe.Add(mBase, uint32(v19)+108))
	if v149 != 0 {
		goto L46
	} else {
		goto L47
	}
L16:
	;
	if v83 < int32(0) {
		goto L33
	} else {
		goto L34
	}
L17:
	;
	v44 = F_XLogInitBufferForRedo(m, l0, int32(0))
	mBase = m.M
	v45 = m.ExcPending
	if v45 != 0 {
		goto L20
	} else {
		goto L21
	}
L18:
	;
	goto L19
L19:
	;
	v80 = F_XLogReadBufferForRedo(m, l0, int32(0), v19+int32(108))
	mBase = m.M
	v81 = m.ExcPending
	if v81 != 0 {
		goto L20
	} else {
		goto L30
	}
L20:
	;
	return
L21:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v19)+108)) = v44
	v49 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v36)+1)))
	if v49 != 0 {
		goto L22
	} else {
		goto L23
	}
L22:
	;
	v50 = int32(12)
	goto L24
L23:
	;
	v50 = int32(4)
	goto L24
L24:
	;
	if v44 < int32(0) {
		goto L27
	} else {
		goto L28
	}
L25:
	;
	v83 = v44
	goto L16
L26:
	;
	F_PageInit(m, v68, int32(8192), int32(8))
	mBase = m.M
	v72 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v68)+16)))
	v73 = v68 + v72
	v74 = int32(65410)
	*(*uint16)(unsafe.Add(mBase, uint32(v73)+6)) = uint16(v74)
	*(*uint16)(unsafe.Add(mBase, uint32(v73))) = uint16(v50)
	goto L25
L27:
	;
	v54 = *(*int32)(unsafe.Add(mBase, _consts[12]))
	v60 = *(*int32)(unsafe.Add(mBase, uint32(v54+(v44^int32(-1))<<(uint(int32(2))%32))))
	v68 = v60
	goto L26
L28:
	;
	goto L29
L29:
	;
	v62 = *(*int32)(unsafe.Add(mBase, _consts[13]))
	v68 = v62 + v44<<(uint(int32(13))%32) + int32(-8192)
	goto L26
L30:
	;
	if v80 != 0 {
		goto L15
	} else {
		goto L31
	}
L31:
	;
	v82 = *(*int32)(unsafe.Add(mBase, uint32(v19)+108))
	v83 = v82
	goto L16
L32:
	;
	v102 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v36)+2)))
	v103 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v36)+4)))
	if v102 != v103 {
		goto L37
	} else {
		goto L38
	}
L33:
	;
	v87 = *(*int32)(unsafe.Add(mBase, _consts[12]))
	v93 = *(*int32)(unsafe.Add(mBase, uint32(v87+(v83^int32(-1))<<(uint(int32(2))%32))))
	v101 = v93
	goto L32
L34:
	;
	goto L35
L35:
	;
	v95 = *(*int32)(unsafe.Add(mBase, _consts[13]))
	v101 = v95 + v83<<(uint(int32(13))%32) + int32(-8192)
	goto L32
L36:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v101))) = base.I64_rotr(v35, int64(32))
	v142 = *(*int32)(unsafe.Add(mBase, uint32(v19)+108))
	F_MarkBufferDirty(m, v142)
	mBase = m.M
	v144 = m.ExcPending
	if v144 != 0 {
		goto L20
	} else {
		goto L45
	}
L37:
	;
	F_addOrReplaceTuple(m, v101, v38, int32(base.Ui32(v39)>>(uint(int32(2))%32)), v102)
	mBase = m.M
	v108 = m.ExcPending
	if v108 != 0 {
		goto L20
	} else {
		goto L40
	}
L38:
	;
	goto L39
L39:
	;
	F_PageIndexTupleDelete(m, v101, v102)
	mBase = m.M
	v128 = m.ExcPending
	if v128 != 0 {
		goto L20
	} else {
		goto L42
	}
L40:
	;
	v109 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v36)+4)))
	if v109 == int32(0) {
		goto L36
	} else {
		goto L41
	}
L41:
	;
	v115 = *(*int32)(unsafe.Add(mBase, uint32(v109<<(uint(int32(2))%32)+v101)+20))
	v118 = v101 + v115&int32(32767)
	v119 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v36)+2)))
	v122 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v118)+4)))
	v125 = v119&int32(16383) | v122&int32(49152)
	*(*uint16)(unsafe.Add(mBase, uint32(v118)+4)) = uint16(v125)
	goto L36
L42:
	;
	v130 = int32(base.Ui32(v39) >> (uint(int32(2)) % 32))
	v131 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v36)+2)))
	v133 = F_PageAddItemExtended(m, v101, v38, v130, v131, int32(0))
	mBase = m.M
	v134 = m.ExcPending
	if v134 != 0 {
		goto L20
	} else {
		goto L43
	}
L43:
	;
	v135 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v36)+2)))
	if v133 != v135 {
		goto L5
	} else {
		goto L44
	}
L44:
	;
	goto L36
L45:
	;
	goto L15
L46:
	;
	F_UnlockReleaseBuffer(m, v149)
	mBase = m.M
	v151 = m.ExcPending
	if v151 != 0 {
		goto L20
	} else {
		goto L49
	}
L47:
	;
	goto L48
L48:
	;
	v152 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v36)+6)))
	if v152 == int32(0) {
		goto L6
	} else {
		goto L50
	}
L49:
	;
	goto L48
L50:
	;
	v158 = F_XLogReadBufferForRedo(m, l0, int32(1), v19+int32(108))
	mBase = m.M
	v159 = m.ExcPending
	if v159 != 0 {
		goto L20
	} else {
		goto L51
	}
L51:
	;
	if v158 == int32(0) {
		goto L52
	} else {
		goto L53
	}
L52:
	;
	v162 = int32(0)
	F_XLogRecGetBlockTag(m, l0, v162, v162, v162, v19+int32(104))
	mBase = m.M
	v168 = m.ExcPending
	if v168 != 0 {
		goto L20
	} else {
		goto L55
	}
L53:
	;
	goto L54
L54:
	;
	v208 = *(*int32)(unsafe.Add(mBase, uint32(v19)+108))
	if v208 == int32(0) {
		goto L6
	} else {
		goto L62
	}
L55:
	;
	v169 = *(*int32)(unsafe.Add(mBase, uint32(v19)+108))
	if v169 < int32(0) {
		goto L57
	} else {
		goto L58
	}
L56:
	;
	v188 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v36)+6)))
	v192 = *(*int32)(unsafe.Add(mBase, uint32(v187+v188<<(uint(int32(2))%32))+20))
	v196 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v36)+8)))
	v197 = *(*int32)(unsafe.Add(mBase, uint32(v19)+104))
	v198 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v36)+2)))
	F_spgUpdateNodeLink(m, v192&int32(32767)+v187, v196, v197, v198)
	mBase = m.M
	v200 = m.ExcPending
	if v200 != 0 {
		goto L20
	} else {
		goto L60
	}
L57:
	;
	v173 = *(*int32)(unsafe.Add(mBase, _consts[12]))
	v179 = *(*int32)(unsafe.Add(mBase, uint32(v173+(v169^int32(-1))<<(uint(int32(2))%32))))
	v187 = v179
	goto L56
L58:
	;
	goto L59
L59:
	;
	v181 = *(*int32)(unsafe.Add(mBase, _consts[13]))
	v187 = v181 + v169<<(uint(int32(13))%32) + int32(-8192)
	goto L56
L60:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v187))) = base.I64_rotr(v35, int64(32))
	v204 = *(*int32)(unsafe.Add(mBase, uint32(v19)+108))
	F_MarkBufferDirty(m, v204)
	mBase = m.M
	v206 = m.ExcPending
	if v206 != 0 {
		goto L20
	} else {
		goto L61
	}
L61:
	;
	goto L54
L62:
	;
	F_UnlockReleaseBuffer(m, v208)
	mBase = m.M
	v212 = m.ExcPending
	if v212 != 0 {
		goto L20
	} else {
		goto L63
	}
L63:
	;
	goto L6
L64:
	;
	v222 = *(*int32)(unsafe.Add(mBase, uint32(v214)+12))
	v223 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v214)+16)))
	v229 = F__emscripten_memset_bulkmem(m, v19+int32(108), base.I32_extend8_s(int32(0)), int32(84))
	mBase = m.M
	goto L65
L65:
	;
	*(*uint8)(unsafe.Add(mBase, uint32(v19)+188)) = uint8(v223)
	*(*int32)(unsafe.Add(mBase, uint32(v19)+184)) = v222
	v233 = F_palloc0(m, int32(16))
	mBase = m.M
	v234 = m.ExcPending
	if v234 != 0 {
		goto L20
	} else {
		goto L66
	}
L66:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v19)+180)) = v233
	v237 = v214 + int32(20)
	v238 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v214))))
	v239 = int32(1)
	v241 = v237 + v238<<(uint(v239)%32)
	v245 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v214)+3)))
	if v245 != 0 {
		goto L67
	} else {
		goto L68
	}
L67:
	;
	v246 = v239
	goto L69
L68:
	;
	v246 = v238 + v239
	goto L69
L69:
	;
	v247 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v214)+2)))
	if v247 == int32(1) {
		goto L72
	} else {
		goto L73
	}
L70:
	;
	v364 = *(*int32)(unsafe.Add(mBase, uint32(v19)+104))
	if v364 != 0 {
		goto L95
	} else {
		goto L96
	}
L71:
	;
	v292 = *(*int32)(unsafe.Add(mBase, uint32(v19)+104))
	if v292 < int32(0) {
		goto L87
	} else {
		goto L88
	}
L72:
	;
	v251 = F_XLogInitBufferForRedo(m, l0, int32(1))
	mBase = m.M
	v252 = m.ExcPending
	if v252 != 0 {
		goto L20
	} else {
		goto L75
	}
L73:
	;
	goto L74
L74:
	;
	v287 = F_XLogReadBufferForRedo(m, l0, int32(1), v19+int32(104))
	mBase = m.M
	v288 = m.ExcPending
	if v288 != 0 {
		goto L20
	} else {
		goto L84
	}
L75:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v19)+104)) = v251
	v256 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v214)+4)))
	if v256 != 0 {
		goto L76
	} else {
		goto L77
	}
L76:
	;
	v257 = int32(12)
	goto L78
L77:
	;
	v257 = int32(4)
	goto L78
L78:
	;
	if v251 < int32(0) {
		goto L81
	} else {
		goto L82
	}
L79:
	;
	goto L71
L80:
	;
	F_PageInit(m, v275, int32(8192), int32(8))
	mBase = m.M
	v279 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v275)+16)))
	v280 = v275 + v279
	v281 = int32(65410)
	*(*uint16)(unsafe.Add(mBase, uint32(v280)+6)) = uint16(v281)
	*(*uint16)(unsafe.Add(mBase, uint32(v280))) = uint16(v257)
	goto L79
L81:
	;
	v261 = *(*int32)(unsafe.Add(mBase, _consts[12]))
	v267 = *(*int32)(unsafe.Add(mBase, uint32(v261+(v251^int32(-1))<<(uint(int32(2))%32))))
	v275 = v267
	goto L80
L82:
	;
	goto L83
L83:
	;
	v269 = *(*int32)(unsafe.Add(mBase, _consts[13]))
	v275 = v269 + v251<<(uint(int32(13))%32) + int32(-8192)
	goto L80
L84:
	;
	if v287 != 0 {
		goto L70
	} else {
		goto L85
	}
L85:
	;
	goto L71
L86:
	;
	v314 = v246<<(uint(int32(1))%32) + v241
	v316 = int32(0)
	goto L90
L87:
	;
	v296 = *(*int32)(unsafe.Add(mBase, _consts[12]))
	v302 = *(*int32)(unsafe.Add(mBase, uint32(v296+(v292^int32(-1))<<(uint(int32(2))%32))))
	v310 = v302
	goto L86
L88:
	;
	goto L89
L89:
	;
	v304 = *(*int32)(unsafe.Add(mBase, _consts[13]))
	v310 = v304 + v292<<(uint(int32(13))%32) + int32(-8192)
	goto L86
L90:
	;
	v329 = *(*int32)(unsafe.Add(mBase, uint32(v314)))
	v331 = int32(base.Ui32(v329) >> (uint(int32(2)) % 32))
	v335 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v241+v316<<(uint(int32(1))%32)))))
	F_addOrReplaceTuple(m, v310, v314, v331, v335)
	mBase = m.M
	v337 = m.ExcPending
	if v337 != 0 {
		goto L20
	} else {
		goto L92
	}
L91:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v310))) = base.I64_rotr(v213, int64(32))
	v345 = *(*int32)(unsafe.Add(mBase, uint32(v19)+104))
	F_MarkBufferDirty(m, v345)
	mBase = m.M
	v347 = m.ExcPending
	if v347 != 0 {
		goto L20
	} else {
		goto L94
	}
L92:
	;
	v340 = v316 + int32(1)
	if v340 != v246 {
		v314 = v314 + v331
		v316 = v340
		goto L90
	} else {
		goto L93
	}
L93:
	;
	goto L91
L94:
	;
	goto L70
L95:
	;
	F_UnlockReleaseBuffer(m, v364)
	mBase = m.M
	v366 = m.ExcPending
	if v366 != 0 {
		goto L20
	} else {
		goto L98
	}
L96:
	;
	goto L97
L97:
	;
	v370 = F_XLogReadBufferForRedo(m, l0, int32(0), v19+int32(104))
	mBase = m.M
	v371 = m.ExcPending
	if v371 != 0 {
		goto L20
	} else {
		goto L99
	}
L98:
	;
	goto L97
L99:
	;
	if v370 == int32(0) {
		goto L100
	} else {
		goto L101
	}
L100:
	;
	v376 = *(*int32)(unsafe.Add(mBase, uint32(v19)+104))
	if v376 < int32(0) {
		goto L104
	} else {
		goto L105
	}
L101:
	;
	goto L102
L102:
	;
	v418 = *(*int32)(unsafe.Add(mBase, uint32(v19)+104))
	if v418 != 0 {
		goto L112
	} else {
		goto L113
	}
L103:
	;
	v395 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v214))))
	v397 = int32(1)
	if v223&v397 != 0 {
		goto L107
	} else {
		goto L108
	}
L104:
	;
	v380 = *(*int32)(unsafe.Add(mBase, _consts[12]))
	v386 = *(*int32)(unsafe.Add(mBase, uint32(v380+(v376^int32(-1))<<(uint(int32(2))%32))))
	v394 = v386
	goto L103
L105:
	;
	goto L106
L106:
	;
	v388 = *(*int32)(unsafe.Add(mBase, _consts[13]))
	v394 = v388 + v376<<(uint(int32(13))%32) + int32(-8192)
	goto L103
L107:
	;
	v400 = int32(3)
	goto L109
L108:
	;
	v400 = v397
	goto L109
L109:
	;
	v402 = *(*int32)(unsafe.Add(mBase, uint32(v19)+100))
	v408 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v241+v246<<(uint(int32(1))%32)-int32(2)))))
	F_spgPageIndexMultiDelete(m, v19+int32(108), v394, v237, v395, v400, int32(3), v402, v408)
	mBase = m.M
	v410 = m.ExcPending
	if v410 != 0 {
		goto L20
	} else {
		goto L110
	}
L110:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v394))) = base.I64_rotr(v213, int64(32))
	v414 = *(*int32)(unsafe.Add(mBase, uint32(v19)+104))
	F_MarkBufferDirty(m, v414)
	mBase = m.M
	v416 = m.ExcPending
	if v416 != 0 {
		goto L20
	} else {
		goto L111
	}
L111:
	;
	goto L102
L112:
	;
	F_UnlockReleaseBuffer(m, v418)
	mBase = m.M
	v420 = m.ExcPending
	if v420 != 0 {
		goto L20
	} else {
		goto L115
	}
L113:
	;
	goto L114
L114:
	;
	v424 = F_XLogReadBufferForRedo(m, l0, int32(2), v19+int32(104))
	mBase = m.M
	v425 = m.ExcPending
	if v425 != 0 {
		goto L20
	} else {
		goto L116
	}
L115:
	;
	goto L114
L116:
	;
	if v424 == int32(0) {
		goto L117
	} else {
		goto L118
	}
L117:
	;
	v428 = *(*int32)(unsafe.Add(mBase, uint32(v19)+104))
	if v428 < int32(0) {
		goto L121
	} else {
		goto L122
	}
L118:
	;
	goto L119
L119:
	;
	v472 = *(*int32)(unsafe.Add(mBase, uint32(v19)+104))
	if v472 == int32(0) {
		goto L6
	} else {
		goto L126
	}
L120:
	;
	v447 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v214)+6)))
	v448 = int32(2)
	v451 = *(*int32)(unsafe.Add(mBase, uint32(v446+v447<<(uint(v448)%32))+20))
	v455 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v214)+8)))
	v456 = *(*int32)(unsafe.Add(mBase, uint32(v19)+100))
	v462 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v241+v246<<(uint(int32(1))%32)-v448))))
	F_spgUpdateNodeLink(m, v451&int32(32767)+v446, v455, v456, v462)
	mBase = m.M
	v464 = m.ExcPending
	if v464 != 0 {
		goto L20
	} else {
		goto L124
	}
L121:
	;
	v432 = *(*int32)(unsafe.Add(mBase, _consts[12]))
	v438 = *(*int32)(unsafe.Add(mBase, uint32(v432+(v428^int32(-1))<<(uint(int32(2))%32))))
	v446 = v438
	goto L120
L122:
	;
	goto L123
L123:
	;
	v440 = *(*int32)(unsafe.Add(mBase, _consts[13]))
	v446 = v440 + v428<<(uint(int32(13))%32) + int32(-8192)
	goto L120
L124:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v446))) = base.I64_rotr(v213, int64(32))
	v468 = *(*int32)(unsafe.Add(mBase, uint32(v19)+104))
	F_MarkBufferDirty(m, v468)
	mBase = m.M
	v470 = m.ExcPending
	if v470 != 0 {
		goto L20
	} else {
		goto L125
	}
L125:
	;
	goto L119
L126:
	;
	F_UnlockReleaseBuffer(m, v472)
	mBase = m.M
	v476 = m.ExcPending
	if v476 != 0 {
		goto L20
	} else {
		goto L127
	}
L127:
	;
	goto L6
L128:
	;
	*(*uint8)(unsafe.Add(mBase, uint32(v19)+188)) = uint8(v481)
	*(*int32)(unsafe.Add(mBase, uint32(v19)+184)) = v480
	v491 = F_palloc0(m, int32(16))
	mBase = m.M
	v492 = m.ExcPending
	if v492 != 0 {
		goto L20
	} else {
		goto L129
	}
L129:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v19)+180)) = v491
	v495 = v478 + int32(20)
	v496 = *(*int32)(unsafe.Add(mBase, uint32(l0)+96))
	v497 = *(*int32)(unsafe.Add(mBase, uint32(v496)+72))
	if int32(0) < v497 {
		goto L131
	} else {
		goto L132
	}
L130:
	;
	v548 = int32(0)
	F_XLogRecGetBlockTag(m, l0, v548, v548, v548, v19+int32(100))
	mBase = m.M
	v554 = m.ExcPending
	if v554 != 0 {
		goto L20
	} else {
		goto L149
	}
L131:
	;
	v500 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v496)+128)))
	if v500 != 0 {
		goto L130
	} else {
		goto L134
	}
L132:
	;
	goto L133
L133:
	;
	v504 = F_XLogReadBufferForRedo(m, l0, int32(0), v19+int32(104))
	mBase = m.M
	v505 = m.ExcPending
	if v505 != 0 {
		goto L20
	} else {
		goto L135
	}
L134:
	;
	goto L133
L135:
	;
	if v504 == int32(0) {
		goto L136
	} else {
		goto L137
	}
L136:
	;
	v508 = *(*int32)(unsafe.Add(mBase, uint32(v19)+104))
	if v508 < int32(0) {
		goto L140
	} else {
		goto L141
	}
L137:
	;
	goto L138
L138:
	;
	v543 = *(*int32)(unsafe.Add(mBase, uint32(v19)+104))
	if v543 == int32(0) {
		goto L6
	} else {
		goto L147
	}
L139:
	;
	v527 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v478))))
	F_PageIndexTupleDelete(m, v526, v527)
	mBase = m.M
	v529 = m.ExcPending
	if v529 != 0 {
		goto L20
	} else {
		goto L143
	}
L140:
	;
	v512 = *(*int32)(unsafe.Add(mBase, _consts[12]))
	v518 = *(*int32)(unsafe.Add(mBase, uint32(v512+(v508^int32(-1))<<(uint(int32(2))%32))))
	v526 = v518
	goto L139
L141:
	;
	goto L142
L142:
	;
	v520 = *(*int32)(unsafe.Add(mBase, _consts[13]))
	v526 = v520 + v508<<(uint(int32(13))%32) + int32(-8192)
	goto L139
L143:
	;
	v530 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v478))))
	v532 = F_PageAddItemExtended(m, v526, v495, v479, v530, int32(0))
	mBase = m.M
	v533 = m.ExcPending
	if v533 != 0 {
		goto L20
	} else {
		goto L144
	}
L144:
	;
	v534 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v478))))
	if v532 != v534 {
		goto L4
	} else {
		goto L145
	}
L145:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v526))) = base.I64_rotr(v477, int64(32))
	v539 = *(*int32)(unsafe.Add(mBase, uint32(v19)+104))
	F_MarkBufferDirty(m, v539)
	mBase = m.M
	v541 = m.ExcPending
	if v541 != 0 {
		goto L20
	} else {
		goto L146
	}
L146:
	;
	goto L138
L147:
	;
	F_UnlockReleaseBuffer(m, v543)
	mBase = m.M
	v547 = m.ExcPending
	if v547 != 0 {
		goto L20
	} else {
		goto L148
	}
L148:
	;
	goto L6
L149:
	;
	v556 = int32(0)
	F_XLogRecGetBlockTag(m, l0, int32(1), v556, v556, v19+int32(96))
	mBase = m.M
	v561 = m.ExcPending
	if v561 != 0 {
		goto L20
	} else {
		goto L150
	}
L150:
	;
	v562 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v478)+4)))
	if v562 == int32(1) {
		goto L153
	} else {
		goto L154
	}
L151:
	;
	v647 = *(*int32)(unsafe.Add(mBase, uint32(v19)+104))
	if v647 != 0 {
		goto L174
	} else {
		goto L175
	}
L152:
	;
	v602 = *(*int32)(unsafe.Add(mBase, uint32(v19)+104))
	if v602 < int32(0) {
		goto L165
	} else {
		goto L166
	}
L153:
	;
	v566 = F_XLogInitBufferForRedo(m, l0, int32(1))
	mBase = m.M
	v567 = m.ExcPending
	if v567 != 0 {
		goto L20
	} else {
		goto L156
	}
L154:
	;
	goto L155
L155:
	;
	v599 = F_XLogReadBufferForRedo(m, l0, int32(1), v19+int32(104))
	mBase = m.M
	v600 = m.ExcPending
	if v600 != 0 {
		goto L20
	} else {
		goto L162
	}
L156:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v19)+104)) = v566
	v569 = int32(0)
	if v566 < v569 {
		goto L159
	} else {
		goto L160
	}
L157:
	;
	goto L152
L158:
	;
	F_PageInit(m, v587, int32(8192), int32(8))
	mBase = m.M
	v591 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v587)+16)))
	v592 = v587 + v591
	v593 = int32(65410)
	*(*uint16)(unsafe.Add(mBase, uint32(v592)+6)) = uint16(v593)
	*(*uint16)(unsafe.Add(mBase, uint32(v592))) = uint16(v569)
	goto L157
L159:
	;
	v573 = *(*int32)(unsafe.Add(mBase, _consts[12]))
	v579 = *(*int32)(unsafe.Add(mBase, uint32(v573+(v566^int32(-1))<<(uint(int32(2))%32))))
	v587 = v579
	goto L158
L160:
	;
	goto L161
L161:
	;
	v581 = *(*int32)(unsafe.Add(mBase, _consts[13]))
	v587 = v581 + v566<<(uint(int32(13))%32) + int32(-8192)
	goto L158
L162:
	;
	if v599 != 0 {
		goto L151
	} else {
		goto L163
	}
L163:
	;
	goto L152
L164:
	;
	v621 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v478)+2)))
	F_addOrReplaceTuple(m, v620, v495, v479, v621)
	mBase = m.M
	v623 = m.ExcPending
	if v623 != 0 {
		goto L20
	} else {
		goto L168
	}
L165:
	;
	v606 = *(*int32)(unsafe.Add(mBase, _consts[12]))
	v612 = *(*int32)(unsafe.Add(mBase, uint32(v606+(v602^int32(-1))<<(uint(int32(2))%32))))
	v620 = v612
	goto L164
L166:
	;
	goto L167
L167:
	;
	v614 = *(*int32)(unsafe.Add(mBase, _consts[13]))
	v620 = v614 + v602<<(uint(int32(13))%32) + int32(-8192)
	goto L164
L168:
	;
	v624 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v478)+5)))
	if v624 == int32(1) {
		goto L169
	} else {
		goto L170
	}
L169:
	;
	v627 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v478)+6)))
	v631 = *(*int32)(unsafe.Add(mBase, uint32(v627<<(uint(int32(2))%32)+v620)+20))
	v635 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v478)+8)))
	v636 = *(*int32)(unsafe.Add(mBase, uint32(v19)+96))
	v637 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v478)+2)))
	F_spgUpdateNodeLink(m, v620+v631&int32(32767), v635, v636, v637)
	mBase = m.M
	v639 = m.ExcPending
	if v639 != 0 {
		goto L20
	} else {
		goto L172
	}
L170:
	;
	goto L171
L171:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v620))) = base.I64_rotr(v477, int64(32))
	v643 = *(*int32)(unsafe.Add(mBase, uint32(v19)+104))
	F_MarkBufferDirty(m, v643)
	mBase = m.M
	v645 = m.ExcPending
	if v645 != 0 {
		goto L20
	} else {
		goto L173
	}
L172:
	;
	goto L171
L173:
	;
	goto L151
L174:
	;
	F_UnlockReleaseBuffer(m, v647)
	mBase = m.M
	v649 = m.ExcPending
	if v649 != 0 {
		goto L20
	} else {
		goto L177
	}
L175:
	;
	goto L176
L176:
	;
	v653 = F_XLogReadBufferForRedo(m, l0, int32(0), v19+int32(104))
	mBase = m.M
	v654 = m.ExcPending
	if v654 != 0 {
		goto L20
	} else {
		goto L178
	}
L177:
	;
	goto L176
L178:
	;
	if v653 == int32(0) {
		goto L179
	} else {
		goto L180
	}
L179:
	;
	v657 = *(*int32)(unsafe.Add(mBase, uint32(v19)+104))
	if v657 < int32(0) {
		goto L183
	} else {
		goto L184
	}
L180:
	;
	goto L181
L181:
	;
	v789 = *(*int32)(unsafe.Add(mBase, uint32(v19)+104))
	if v789 != 0 {
		goto L209
	} else {
		goto L210
	}
L182:
	;
	v676 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v19)+188)))
	if v676 == int32(1) {
		goto L187
	} else {
		goto L188
	}
L183:
	;
	v661 = *(*int32)(unsafe.Add(mBase, _consts[12]))
	v667 = *(*int32)(unsafe.Add(mBase, uint32(v661+(v657^int32(-1))<<(uint(int32(2))%32))))
	v675 = v667
	goto L182
L184:
	;
	goto L185
L185:
	;
	v669 = *(*int32)(unsafe.Add(mBase, _consts[13]))
	v675 = v669 + v657<<(uint(int32(13))%32) + int32(-8192)
	goto L182
L186:
	;
	v742 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v478))))
	F_PageIndexTupleDelete(m, v675, v742)
	mBase = m.M
	v744 = m.ExcPending
	if v744 != 0 {
		goto L20
	} else {
		goto L198
	}
L187:
	;
	v685 = *(*int32)(unsafe.Add(mBase, uint32(v19+int32(108))+72))
	*(*int32)(unsafe.Add(mBase, uint32(v685))) = int32(67)
	v691 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v685)+4)))
	v693 = v691 & int32(49152)
	*(*uint16)(unsafe.Add(mBase, uint32(v685)+4)) = uint16(v693)
	goto L192
L188:
	;
	goto L189
L189:
	;
	v711 = v19 + int32(108)
	v713 = *(*int32)(unsafe.Add(mBase, uint32(v19)+96))
	v714 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v478)+2)))
	v716 = *(*int32)(unsafe.Add(mBase, uint32(v711)+72))
	*(*int32)(unsafe.Add(mBase, uint32(v716))) = int32(65)
	v722 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v716)+4)))
	v724 = v722 & int32(49152)
	*(*uint16)(unsafe.Add(mBase, uint32(v716)+4)) = uint16(v724)
	goto L195
L190:
	;
	v741 = v685
	goto L186
L192:
	;
	goto L193
L193:
	;
	v704 = int32(0)
	*(*uint16)(unsafe.Add(mBase, uint32(v685)+10)) = uint16(v704)
	*(*int32)(unsafe.Add(mBase, uint32(v685)+6)) = int32(-1)
	*(*int32)(unsafe.Add(mBase, uint32(v685)+12)) = v704
	goto L190
L194:
	;
	v741 = v716
	goto L186
L195:
	;
	*(*uint16)(unsafe.Add(mBase, uint32(v716)+10)) = uint16(v714)
	*(*uint16)(unsafe.Add(mBase, uint32(v716)+8)) = uint16(v713)
	v731 = int32(base.Ui32(v713) >> (uint(int32(16)) % 32))
	*(*uint16)(unsafe.Add(mBase, uint32(v716)+6)) = uint16(v731)
	v733 = *(*int32)(unsafe.Add(mBase, uint32(v711)+76))
	*(*int32)(unsafe.Add(mBase, uint32(v716)+12)) = v733
	goto L194
L198:
	;
	v745 = *(*int32)(unsafe.Add(mBase, uint32(v741)))
	v748 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v478))))
	v750 = F_PageAddItemExtended(m, v675, v741, int32(base.Ui32(v745)>>(uint(int32(2))%32)), v748, int32(0))
	mBase = m.M
	v751 = m.ExcPending
	if v751 != 0 {
		goto L20
	} else {
		goto L199
	}
L199:
	;
	v752 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v478))))
	if v750 != v752 {
		goto L3
	} else {
		goto L200
	}
L200:
	;
	v754 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v675)+16)))
	v758 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v19)+188)))
	if v758 != 0 {
		goto L201
	} else {
		goto L202
	}
L201:
	;
	v759 = int32(4)
	goto L203
L202:
	;
	v759 = int32(2)
	goto L203
L203:
	;
	v760 = v675 + v754 + v759
	v761 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v760))))
	v763 = v761 + int32(1)
	*(*uint16)(unsafe.Add(mBase, uint32(v760))) = uint16(v763)
	v765 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v478)+5)))
	if v765 == int32(0) {
		goto L204
	} else {
		goto L205
	}
L204:
	;
	v768 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v478)+6)))
	v772 = *(*int32)(unsafe.Add(mBase, uint32(v768<<(uint(int32(2))%32)+v675)+20))
	v776 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v478)+8)))
	v777 = *(*int32)(unsafe.Add(mBase, uint32(v19)+96))
	v778 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v478)+2)))
	F_spgUpdateNodeLink(m, v675+v772&int32(32767), v776, v777, v778)
	mBase = m.M
	v780 = m.ExcPending
	if v780 != 0 {
		goto L20
	} else {
		goto L207
	}
L205:
	;
	goto L206
L206:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v675))) = base.I64_rotr(v477, int64(32))
	v784 = *(*int32)(unsafe.Add(mBase, uint32(v19)+104))
	F_MarkBufferDirty(m, v784)
	mBase = m.M
	v786 = m.ExcPending
	if v786 != 0 {
		goto L20
	} else {
		goto L208
	}
L207:
	;
	goto L206
L208:
	;
	goto L181
L209:
	;
	F_UnlockReleaseBuffer(m, v789)
	mBase = m.M
	v791 = m.ExcPending
	if v791 != 0 {
		goto L20
	} else {
		goto L212
	}
L210:
	;
	goto L211
L211:
	;
	v792 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v478)+5)))
	if v792 != int32(2) {
		goto L6
	} else {
		goto L213
	}
L212:
	;
	goto L211
L213:
	;
	v798 = F_XLogReadBufferForRedo(m, l0, int32(2), v19+int32(104))
	mBase = m.M
	v799 = m.ExcPending
	if v799 != 0 {
		goto L20
	} else {
		goto L214
	}
L214:
	;
	if v798 == int32(0) {
		goto L215
	} else {
		goto L216
	}
L215:
	;
	v802 = *(*int32)(unsafe.Add(mBase, uint32(v19)+104))
	if v802 < int32(0) {
		goto L219
	} else {
		goto L220
	}
L216:
	;
	goto L217
L217:
	;
	v841 = *(*int32)(unsafe.Add(mBase, uint32(v19)+104))
	if v841 == int32(0) {
		goto L6
	} else {
		goto L224
	}
L218:
	;
	v821 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v478)+6)))
	v825 = *(*int32)(unsafe.Add(mBase, uint32(v820+v821<<(uint(int32(2))%32))+20))
	v829 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v478)+8)))
	v830 = *(*int32)(unsafe.Add(mBase, uint32(v19)+96))
	v831 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v478)+2)))
	F_spgUpdateNodeLink(m, v825&int32(32767)+v820, v829, v830, v831)
	mBase = m.M
	v833 = m.ExcPending
	if v833 != 0 {
		goto L20
	} else {
		goto L222
	}
L219:
	;
	v806 = *(*int32)(unsafe.Add(mBase, _consts[12]))
	v812 = *(*int32)(unsafe.Add(mBase, uint32(v806+(v802^int32(-1))<<(uint(int32(2))%32))))
	v820 = v812
	goto L218
L220:
	;
	goto L221
L221:
	;
	v814 = *(*int32)(unsafe.Add(mBase, _consts[13]))
	v820 = v814 + v802<<(uint(int32(13))%32) + int32(-8192)
	goto L218
L222:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v820))) = base.I64_rotr(v477, int64(32))
	v837 = *(*int32)(unsafe.Add(mBase, uint32(v19)+104))
	F_MarkBufferDirty(m, v837)
	mBase = m.M
	v839 = m.ExcPending
	if v839 != 0 {
		goto L20
	} else {
		goto L223
	}
L223:
	;
	goto L217
L224:
	;
	F_UnlockReleaseBuffer(m, v841)
	mBase = m.M
	v845 = m.ExcPending
	if v845 != 0 {
		goto L20
	} else {
		goto L225
	}
L225:
	;
	goto L6
L226:
	;
	v932 = F_XLogReadBufferForRedo(m, l0, int32(0), v19+int32(108))
	mBase = m.M
	v933 = m.ExcPending
	if v933 != 0 {
		goto L20
	} else {
		goto L249
	}
L227:
	;
	v854 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v847)+4)))
	if v854 == int32(1) {
		goto L230
	} else {
		goto L231
	}
L228:
	;
	v923 = *(*int32)(unsafe.Add(mBase, uint32(v19)+108))
	if v923 == int32(0) {
		goto L226
	} else {
		goto L247
	}
L229:
	;
	if v894 < int32(0) {
		goto L242
	} else {
		goto L243
	}
L230:
	;
	v858 = F_XLogInitBufferForRedo(m, l0, int32(1))
	mBase = m.M
	v859 = m.ExcPending
	if v859 != 0 {
		goto L20
	} else {
		goto L233
	}
L231:
	;
	goto L232
L232:
	;
	v891 = F_XLogReadBufferForRedo(m, l0, int32(1), v19+int32(108))
	mBase = m.M
	v892 = m.ExcPending
	if v892 != 0 {
		goto L20
	} else {
		goto L239
	}
L233:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v19)+108)) = v858
	v861 = int32(0)
	if v858 < v861 {
		goto L236
	} else {
		goto L237
	}
L234:
	;
	v894 = v858
	goto L229
L235:
	;
	F_PageInit(m, v879, int32(8192), int32(8))
	mBase = m.M
	v883 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v879)+16)))
	v884 = v879 + v883
	v885 = int32(65410)
	*(*uint16)(unsafe.Add(mBase, uint32(v884)+6)) = uint16(v885)
	*(*uint16)(unsafe.Add(mBase, uint32(v884))) = uint16(v861)
	goto L234
L236:
	;
	v865 = *(*int32)(unsafe.Add(mBase, _consts[12]))
	v871 = *(*int32)(unsafe.Add(mBase, uint32(v865+(v858^int32(-1))<<(uint(int32(2))%32))))
	v879 = v871
	goto L235
L237:
	;
	goto L238
L238:
	;
	v873 = *(*int32)(unsafe.Add(mBase, _consts[13]))
	v879 = v873 + v858<<(uint(int32(13))%32) + int32(-8192)
	goto L235
L239:
	;
	if v891 != 0 {
		goto L228
	} else {
		goto L240
	}
L240:
	;
	v893 = *(*int32)(unsafe.Add(mBase, uint32(v19)+108))
	v894 = v893
	goto L229
L241:
	;
	v913 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v847)+2)))
	F_addOrReplaceTuple(m, v912, v851, v852, v913)
	mBase = m.M
	v915 = m.ExcPending
	if v915 != 0 {
		goto L20
	} else {
		goto L245
	}
L242:
	;
	v898 = *(*int32)(unsafe.Add(mBase, _consts[12]))
	v904 = *(*int32)(unsafe.Add(mBase, uint32(v898+(v894^int32(-1))<<(uint(int32(2))%32))))
	v912 = v904
	goto L241
L243:
	;
	goto L244
L244:
	;
	v906 = *(*int32)(unsafe.Add(mBase, _consts[13]))
	v912 = v906 + v894<<(uint(int32(13))%32) + int32(-8192)
	goto L241
L245:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v912))) = base.I64_rotr(v846, int64(32))
	v919 = *(*int32)(unsafe.Add(mBase, uint32(v19)+108))
	F_MarkBufferDirty(m, v919)
	mBase = m.M
	v921 = m.ExcPending
	if v921 != 0 {
		goto L20
	} else {
		goto L246
	}
L246:
	;
	goto L228
L247:
	;
	F_UnlockReleaseBuffer(m, v923)
	mBase = m.M
	v927 = m.ExcPending
	if v927 != 0 {
		goto L20
	} else {
		goto L248
	}
L248:
	;
	goto L226
L249:
	;
	if v932 == int32(0) {
		goto L250
	} else {
		goto L251
	}
L250:
	;
	v936 = *(*int32)(unsafe.Add(mBase, uint32(v19)+108))
	if v936 < int32(0) {
		goto L254
	} else {
		goto L255
	}
L251:
	;
	goto L252
L252:
	;
	v977 = *(*int32)(unsafe.Add(mBase, uint32(v19)+108))
	if v977 == int32(0) {
		goto L6
	} else {
		goto L265
	}
L253:
	;
	v955 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v847))))
	F_PageIndexTupleDelete(m, v954, v955)
	mBase = m.M
	v957 = m.ExcPending
	if v957 != 0 {
		goto L20
	} else {
		goto L257
	}
L254:
	;
	v940 = *(*int32)(unsafe.Add(mBase, _consts[12]))
	v946 = *(*int32)(unsafe.Add(mBase, uint32(v940+(v936^int32(-1))<<(uint(int32(2))%32))))
	v954 = v946
	goto L253
L255:
	;
	goto L256
L256:
	;
	v948 = *(*int32)(unsafe.Add(mBase, _consts[13]))
	v954 = v948 + v936<<(uint(int32(13))%32) + int32(-8192)
	goto L253
L257:
	;
	v958 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v847))))
	v960 = F_PageAddItemExtended(m, v954, v849, v850, v958, int32(0))
	mBase = m.M
	v961 = m.ExcPending
	if v961 != 0 {
		goto L20
	} else {
		goto L258
	}
L258:
	;
	v962 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v847))))
	if v960 != v962 {
		goto L2
	} else {
		goto L259
	}
L259:
	;
	v964 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v847)+5)))
	if v964 == int32(1) {
		goto L260
	} else {
		goto L261
	}
L260:
	;
	v967 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v847)+2)))
	F_addOrReplaceTuple(m, v954, v851, v852, v967)
	mBase = m.M
	v969 = m.ExcPending
	if v969 != 0 {
		goto L20
	} else {
		goto L263
	}
L261:
	;
	goto L262
L262:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v954))) = base.I64_rotr(v846, int64(32))
	v973 = *(*int32)(unsafe.Add(mBase, uint32(v19)+108))
	F_MarkBufferDirty(m, v973)
	mBase = m.M
	v975 = m.ExcPending
	if v975 != 0 {
		goto L20
	} else {
		goto L264
	}
L263:
	;
	goto L262
L264:
	;
	goto L252
L265:
	;
	F_UnlockReleaseBuffer(m, v977)
	mBase = m.M
	v981 = m.ExcPending
	if v981 != 0 {
		goto L20
	} else {
		goto L266
	}
L266:
	;
	goto L6
L267:
	;
	v991 = *(*int32)(unsafe.Add(mBase, uint32(v983)+20))
	v992 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v983)+24)))
	v998 = F__emscripten_memset_bulkmem(m, v19+int32(108), base.I32_extend8_s(int32(0)), int32(84))
	mBase = m.M
	goto L268
L268:
	;
	*(*uint8)(unsafe.Add(mBase, uint32(v19)+188)) = uint8(v992)
	*(*int32)(unsafe.Add(mBase, uint32(v19)+184)) = v991
	v1002 = F_palloc0(m, int32(16))
	mBase = m.M
	v1003 = m.ExcPending
	if v1003 != 0 {
		goto L20
	} else {
		goto L269
	}
L269:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v19)+180)) = v1002
	v1006 = v983 + int32(28)
	v1007 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v983)+2)))
	v1008 = int32(1)
	v1010 = v1006 + v1007<<(uint(v1008)%32)
	v1011 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v983)+4)))
	v1014 = v1010 + v1011<<(uint(v1008)%32)
	v1015 = v1014 + v1011
	v1016 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v1015)+4)))
	v1017 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v983))))
	if v1017 == v1008 {
		goto L271
	} else {
		goto L272
	}
L270:
	;
	v1123 = *(*int32)(unsafe.Add(mBase, uint32(l0)+96))
	v1124 = *(*int32)(unsafe.Add(mBase, uint32(v1123)+72))
	if int32(0) < v1124 {
		goto L303
	} else {
		goto L304
	}
L271:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v19)+104)) = int32(0)
	v1122 = v2
	goto L270
L272:
	;
	goto L273
L273:
	;
	v1022 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v983)+6)))
	if v1022 == int32(1) {
		goto L274
	} else {
		goto L275
	}
L274:
	;
	v1026 = F_XLogInitBufferForRedo(m, l0, int32(0))
	mBase = m.M
	v1027 = m.ExcPending
	if v1027 != 0 {
		goto L20
	} else {
		goto L277
	}
L275:
	;
	goto L276
L276:
	;
	v1080 = F_XLogReadBufferForRedo(m, l0, int32(0), v19+int32(104))
	mBase = m.M
	v1081 = m.ExcPending
	if v1081 != 0 {
		goto L20
	} else {
		goto L290
	}
L277:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v19)+104)) = v1026
	if v1026 < int32(0) {
		goto L279
	} else {
		goto L280
	}
L278:
	;
	v1049 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v983)+11)))
	if v1049 != 0 {
		goto L282
	} else {
		goto L283
	}
L279:
	;
	v1032 = *(*int32)(unsafe.Add(mBase, _consts[12]))
	v1038 = *(*int32)(unsafe.Add(mBase, uint32(v1032+(v1026^int32(-1))<<(uint(int32(2))%32))))
	v1046 = v1038
	goto L278
L280:
	;
	goto L281
L281:
	;
	v1040 = *(*int32)(unsafe.Add(mBase, _consts[13]))
	v1046 = v1040 + v1026<<(uint(int32(13))%32) + int32(-8192)
	goto L278
L282:
	;
	v1050 = int32(12)
	goto L284
L283:
	;
	v1050 = int32(4)
	goto L284
L284:
	;
	if v1026 < int32(0) {
		goto L287
	} else {
		goto L288
	}
L285:
	;
	v1122 = v1046
	goto L270
L286:
	;
	F_PageInit(m, v1068, int32(8192), int32(8))
	mBase = m.M
	v1072 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v1068)+16)))
	v1073 = v1068 + v1072
	v1074 = int32(65410)
	*(*uint16)(unsafe.Add(mBase, uint32(v1073)+6)) = uint16(v1074)
	*(*uint16)(unsafe.Add(mBase, uint32(v1073))) = uint16(v1050)
	goto L285
L287:
	;
	v1054 = *(*int32)(unsafe.Add(mBase, _consts[12]))
	v1060 = *(*int32)(unsafe.Add(mBase, uint32(v1054+(v1026^int32(-1))<<(uint(int32(2))%32))))
	v1068 = v1060
	goto L286
L288:
	;
	goto L289
L289:
	;
	v1062 = *(*int32)(unsafe.Add(mBase, _consts[13]))
	v1068 = v1062 + v1026<<(uint(int32(13))%32) + int32(-8192)
	goto L286
L290:
	;
	if v1080 != 0 {
		v1122 = v2
		goto L270
	} else {
		goto L291
	}
L291:
	;
	v1082 = *(*int32)(unsafe.Add(mBase, uint32(v19)+104))
	if v1082 < int32(0) {
		goto L293
	} else {
		goto L294
	}
L292:
	;
	v1101 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v983)+2)))
	v1102 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v19)+188)))
	if v1102 == int32(0) {
		goto L296
	} else {
		goto L297
	}
L293:
	;
	v1086 = *(*int32)(unsafe.Add(mBase, _consts[12]))
	v1092 = *(*int32)(unsafe.Add(mBase, uint32(v1086+(v1082^int32(-1))<<(uint(int32(2))%32))))
	v1100 = v1092
	goto L292
L294:
	;
	goto L295
L295:
	;
	v1094 = *(*int32)(unsafe.Add(mBase, _consts[13]))
	v1100 = v1094 + v1082<<(uint(int32(13))%32) + int32(-8192)
	goto L292
L296:
	;
	v1109 = *(*int32)(unsafe.Add(mBase, uint32(v19)+92))
	v1110 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v983)+8)))
	F_spgPageIndexMultiDelete(m, v19+int32(108), v1100, v1006, v1101, int32(1), int32(3), v1109, v1110)
	mBase = m.M
	v1112 = m.ExcPending
	if v1112 != 0 {
		goto L20
	} else {
		goto L299
	}
L297:
	;
	goto L298
L298:
	;
	v1115 = int32(3)
	F_spgPageIndexMultiDelete(m, v19+int32(108), v1100, v1006, v1101, v1115, v1115, int32(-1), int32(0))
	mBase = m.M
	v1120 = m.ExcPending
	if v1120 != 0 {
		goto L20
	} else {
		goto L300
	}
L299:
	;
	v1122 = v1100
	goto L270
L300:
	;
	v1122 = v1100
	goto L270
L301:
	;
	v1210 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v983)+4)))
	if v1210 != 0 {
		goto L328
	} else {
		goto L329
	}
L302:
	;
	v1130 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v983)+7)))
	if v1130 == int32(1) {
		goto L307
	} else {
		goto L308
	}
L303:
	;
	v1127 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1123)+128)))
	if v1127 != 0 {
		goto L302
	} else {
		goto L306
	}
L304:
	;
	goto L305
L305:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v19)+100)) = int32(0)
	v1209 = v2
	goto L301
L306:
	;
	goto L305
L307:
	;
	v1134 = F_XLogInitBufferForRedo(m, l0, int32(1))
	mBase = m.M
	v1135 = m.ExcPending
	if v1135 != 0 {
		goto L20
	} else {
		goto L310
	}
L308:
	;
	goto L309
L309:
	;
	v1188 = F_XLogReadBufferForRedo(m, l0, int32(1), v19+int32(100))
	mBase = m.M
	v1189 = m.ExcPending
	if v1189 != 0 {
		goto L20
	} else {
		goto L323
	}
L310:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v19)+100)) = v1134
	if v1134 < int32(0) {
		goto L312
	} else {
		goto L313
	}
L311:
	;
	v1157 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v983)+11)))
	if v1157 != 0 {
		goto L315
	} else {
		goto L316
	}
L312:
	;
	v1140 = *(*int32)(unsafe.Add(mBase, _consts[12]))
	v1146 = *(*int32)(unsafe.Add(mBase, uint32(v1140+(v1134^int32(-1))<<(uint(int32(2))%32))))
	v1154 = v1146
	goto L311
L313:
	;
	goto L314
L314:
	;
	v1148 = *(*int32)(unsafe.Add(mBase, _consts[13]))
	v1154 = v1148 + v1134<<(uint(int32(13))%32) + int32(-8192)
	goto L311
L315:
	;
	v1158 = int32(12)
	goto L317
L316:
	;
	v1158 = int32(4)
	goto L317
L317:
	;
	if v1134 < int32(0) {
		goto L320
	} else {
		goto L321
	}
L318:
	;
	v1209 = v1154
	goto L301
L319:
	;
	F_PageInit(m, v1176, int32(8192), int32(8))
	mBase = m.M
	v1180 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v1176)+16)))
	v1181 = v1176 + v1180
	v1182 = int32(65410)
	*(*uint16)(unsafe.Add(mBase, uint32(v1181)+6)) = uint16(v1182)
	*(*uint16)(unsafe.Add(mBase, uint32(v1181))) = uint16(v1158)
	goto L318
L320:
	;
	v1162 = *(*int32)(unsafe.Add(mBase, _consts[12]))
	v1168 = *(*int32)(unsafe.Add(mBase, uint32(v1162+(v1134^int32(-1))<<(uint(int32(2))%32))))
	v1176 = v1168
	goto L319
L321:
	;
	goto L322
L322:
	;
	v1170 = *(*int32)(unsafe.Add(mBase, _consts[13]))
	v1176 = v1170 + v1134<<(uint(int32(13))%32) + int32(-8192)
	goto L319
L323:
	;
	if v1188 != 0 {
		v1209 = v2
		goto L301
	} else {
		goto L324
	}
L324:
	;
	v1190 = *(*int32)(unsafe.Add(mBase, uint32(v19)+100))
	if v1190 < int32(0) {
		goto L325
	} else {
		goto L326
	}
L325:
	;
	v1194 = *(*int32)(unsafe.Add(mBase, _consts[12]))
	v1200 = *(*int32)(unsafe.Add(mBase, uint32(v1194+(v1190^int32(-1))<<(uint(int32(2))%32))))
	v1209 = v1200
	goto L301
L326:
	;
	goto L327
L327:
	;
	v1202 = *(*int32)(unsafe.Add(mBase, _consts[13]))
	v1209 = v1202 + v1190<<(uint(int32(13))%32) + int32(-8192)
	goto L301
L328:
	;
	v1214 = v1015 + v1016
	v1216 = int32(0)
	v1218 = v1210
	goto L331
L329:
	;
	goto L330
L330:
	;
	if v1122 != 0 {
		goto L341
	} else {
		goto L342
	}
L331:
	;
	v1229 = *(*int32)(unsafe.Add(mBase, uint32(v1214)))
	v1231 = int32(base.Ui32(v1229) >> (uint(int32(2)) % 32))
	v1233 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1216+v1014))))
	if v1233 != 0 {
		goto L333
	} else {
		goto L334
	}
L332:
	;
	goto L330
L333:
	;
	v1234 = v1209
	goto L335
L334:
	;
	v1234 = v1122
	goto L335
L335:
	;
	if v1234 != 0 {
		goto L336
	} else {
		goto L337
	}
L336:
	;
	v1238 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v1010+v1216<<(uint(int32(1))%32)))))
	F_addOrReplaceTuple(m, v1234, v1214, v1231, v1238)
	mBase = m.M
	v1240 = m.ExcPending
	if v1240 != 0 {
		goto L20
	} else {
		goto L339
	}
L337:
	;
	v1242 = v1218
	goto L338
L338:
	;
	v1245 = v1216 + int32(1)
	if base.Ui32(v1245) < base.Ui32(v1242&int32(65535)) {
		v1214 = v1214 + v1231
		v1216 = v1245
		v1218 = v1242
		goto L331
	} else {
		goto L340
	}
L339:
	;
	v1241 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v983)+4)))
	v1242 = v1241
	goto L338
L340:
	;
	goto L332
L341:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v1122))) = base.I64_rotr(v982, int64(32))
	v1268 = *(*int32)(unsafe.Add(mBase, uint32(v19)+104))
	F_MarkBufferDirty(m, v1268)
	mBase = m.M
	v1270 = m.ExcPending
	if v1270 != 0 {
		goto L20
	} else {
		goto L344
	}
L342:
	;
	goto L343
L343:
	;
	if v1209 != 0 {
		goto L345
	} else {
		goto L346
	}
L344:
	;
	goto L343
L345:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v1209))) = base.I64_rotr(v982, int64(32))
	v1274 = *(*int32)(unsafe.Add(mBase, uint32(v19)+100))
	F_MarkBufferDirty(m, v1274)
	mBase = m.M
	v1276 = m.ExcPending
	if v1276 != 0 {
		goto L20
	} else {
		goto L348
	}
L346:
	;
	goto L347
L347:
	;
	v1277 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v983)+10)))
	if v1277 == int32(1) {
		goto L351
	} else {
		goto L352
	}
L348:
	;
	goto L347
L349:
	;
	v1365 = *(*int32)(unsafe.Add(mBase, uint32(v19)+96))
	if v1365 != 0 {
		goto L375
	} else {
		goto L376
	}
L350:
	;
	if v1320 < int32(0) {
		goto L366
	} else {
		goto L367
	}
L351:
	;
	v1281 = F_XLogInitBufferForRedo(m, l0, int32(2))
	mBase = m.M
	v1282 = m.ExcPending
	if v1282 != 0 {
		goto L20
	} else {
		goto L354
	}
L352:
	;
	goto L353
L353:
	;
	v1317 = F_XLogReadBufferForRedo(m, l0, int32(2), v19+int32(96))
	mBase = m.M
	v1318 = m.ExcPending
	if v1318 != 0 {
		goto L20
	} else {
		goto L363
	}
L354:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v19)+96)) = v1281
	v1286 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v983)+11)))
	if v1286 != 0 {
		goto L355
	} else {
		goto L356
	}
L355:
	;
	v1287 = int32(8)
	goto L357
L356:
	;
	v1287 = int32(0)
	goto L357
L357:
	;
	if v1281 < int32(0) {
		goto L360
	} else {
		goto L361
	}
L358:
	;
	v1320 = v1281
	goto L350
L359:
	;
	F_PageInit(m, v1305, int32(8192), int32(8))
	mBase = m.M
	v1309 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v1305)+16)))
	v1310 = v1305 + v1309
	v1311 = int32(65410)
	*(*uint16)(unsafe.Add(mBase, uint32(v1310)+6)) = uint16(v1311)
	*(*uint16)(unsafe.Add(mBase, uint32(v1310))) = uint16(v1287)
	goto L358
L360:
	;
	v1291 = *(*int32)(unsafe.Add(mBase, _consts[12]))
	v1297 = *(*int32)(unsafe.Add(mBase, uint32(v1291+(v1281^int32(-1))<<(uint(int32(2))%32))))
	v1305 = v1297
	goto L359
L361:
	;
	goto L362
L362:
	;
	v1299 = *(*int32)(unsafe.Add(mBase, _consts[13]))
	v1305 = v1299 + v1281<<(uint(int32(13))%32) + int32(-8192)
	goto L359
L363:
	;
	if v1317 != 0 {
		goto L349
	} else {
		goto L364
	}
L364:
	;
	v1319 = *(*int32)(unsafe.Add(mBase, uint32(v19)+96))
	v1320 = v1319
	goto L350
L365:
	;
	v1339 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v983)+8)))
	F_addOrReplaceTuple(m, v1338, v1015, v1016, v1339)
	mBase = m.M
	v1341 = m.ExcPending
	if v1341 != 0 {
		goto L20
	} else {
		goto L369
	}
L366:
	;
	v1324 = *(*int32)(unsafe.Add(mBase, _consts[12]))
	v1330 = *(*int32)(unsafe.Add(mBase, uint32(v1324+(v1320^int32(-1))<<(uint(int32(2))%32))))
	v1338 = v1330
	goto L365
L367:
	;
	goto L368
L368:
	;
	v1332 = *(*int32)(unsafe.Add(mBase, _consts[13]))
	v1338 = v1332 + v1320<<(uint(int32(13))%32) + int32(-8192)
	goto L365
L369:
	;
	v1342 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v983)+12)))
	if v1342 == int32(1) {
		goto L370
	} else {
		goto L371
	}
L370:
	;
	v1345 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v983)+14)))
	v1349 = *(*int32)(unsafe.Add(mBase, uint32(v1345<<(uint(int32(2))%32)+v1338)+20))
	v1353 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v983)+16)))
	v1354 = *(*int32)(unsafe.Add(mBase, uint32(v19)+92))
	v1355 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v983)+8)))
	F_spgUpdateNodeLink(m, v1338+v1349&int32(32767), v1353, v1354, v1355)
	mBase = m.M
	v1357 = m.ExcPending
	if v1357 != 0 {
		goto L20
	} else {
		goto L373
	}
L371:
	;
	goto L372
L372:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v1338))) = base.I64_rotr(v982, int64(32))
	v1361 = *(*int32)(unsafe.Add(mBase, uint32(v19)+96))
	F_MarkBufferDirty(m, v1361)
	mBase = m.M
	v1363 = m.ExcPending
	if v1363 != 0 {
		goto L20
	} else {
		goto L374
	}
L373:
	;
	goto L372
L374:
	;
	goto L349
L375:
	;
	F_UnlockReleaseBuffer(m, v1365)
	mBase = m.M
	v1367 = m.ExcPending
	if v1367 != 0 {
		goto L20
	} else {
		goto L378
	}
L376:
	;
	goto L377
L377:
	;
	v1368 = *(*int32)(unsafe.Add(mBase, uint32(v19)+104))
	if v1368 != 0 {
		goto L379
	} else {
		goto L380
	}
L378:
	;
	goto L377
L379:
	;
	F_UnlockReleaseBuffer(m, v1368)
	mBase = m.M
	v1370 = m.ExcPending
	if v1370 != 0 {
		goto L20
	} else {
		goto L382
	}
L380:
	;
	goto L381
L381:
	;
	v1371 = *(*int32)(unsafe.Add(mBase, uint32(v19)+100))
	if v1371 != 0 {
		goto L383
	} else {
		goto L384
	}
L382:
	;
	goto L381
L383:
	;
	F_UnlockReleaseBuffer(m, v1371)
	mBase = m.M
	v1373 = m.ExcPending
	if v1373 != 0 {
		goto L20
	} else {
		goto L386
	}
L384:
	;
	goto L385
L385:
	;
	v1374 = *(*int32)(unsafe.Add(mBase, uint32(l0)+96))
	v1375 = *(*int32)(unsafe.Add(mBase, uint32(v1374)+72))
	if v1375 < int32(3) {
		goto L6
	} else {
		goto L387
	}
L386:
	;
	goto L385
L387:
	;
	v1378 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1374)+232)))
	if v1378 != int32(1) {
		goto L6
	} else {
		goto L388
	}
L388:
	;
	v1384 = F_XLogReadBufferForRedo(m, l0, int32(3), v19+int32(88))
	mBase = m.M
	v1385 = m.ExcPending
	if v1385 != 0 {
		goto L20
	} else {
		goto L389
	}
L389:
	;
	if v1384 == int32(0) {
		goto L390
	} else {
		goto L391
	}
L390:
	;
	v1388 = *(*int32)(unsafe.Add(mBase, uint32(v19)+88))
	if v1388 < int32(0) {
		goto L394
	} else {
		goto L395
	}
L391:
	;
	goto L392
L392:
	;
	v1427 = *(*int32)(unsafe.Add(mBase, uint32(v19)+88))
	if v1427 == int32(0) {
		goto L6
	} else {
		goto L399
	}
L393:
	;
	v1407 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v983)+14)))
	v1411 = *(*int32)(unsafe.Add(mBase, uint32(v1406+v1407<<(uint(int32(2))%32))+20))
	v1415 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v983)+16)))
	v1416 = *(*int32)(unsafe.Add(mBase, uint32(v19)+92))
	v1417 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v983)+8)))
	F_spgUpdateNodeLink(m, v1411&int32(32767)+v1406, v1415, v1416, v1417)
	mBase = m.M
	v1419 = m.ExcPending
	if v1419 != 0 {
		goto L20
	} else {
		goto L397
	}
L394:
	;
	v1392 = *(*int32)(unsafe.Add(mBase, _consts[12]))
	v1398 = *(*int32)(unsafe.Add(mBase, uint32(v1392+(v1388^int32(-1))<<(uint(int32(2))%32))))
	v1406 = v1398
	goto L393
L395:
	;
	goto L396
L396:
	;
	v1400 = *(*int32)(unsafe.Add(mBase, _consts[13]))
	v1406 = v1400 + v1388<<(uint(int32(13))%32) + int32(-8192)
	goto L393
L397:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v1406))) = base.I64_rotr(v982, int64(32))
	v1423 = *(*int32)(unsafe.Add(mBase, uint32(v19)+88))
	F_MarkBufferDirty(m, v1423)
	mBase = m.M
	v1425 = m.ExcPending
	if v1425 != 0 {
		goto L20
	} else {
		goto L398
	}
L398:
	;
	goto L392
L399:
	;
	F_UnlockReleaseBuffer(m, v1427)
	mBase = m.M
	v1431 = m.ExcPending
	if v1431 != 0 {
		goto L20
	} else {
		goto L400
	}
L400:
	;
	goto L6
L401:
	;
	*(*uint8)(unsafe.Add(mBase, uint32(v19)+188)) = uint8(v1435)
	*(*int32)(unsafe.Add(mBase, uint32(v19)+184)) = v1434
	v1445 = F_palloc0(m, int32(16))
	mBase = m.M
	v1446 = m.ExcPending
	if v1446 != 0 {
		goto L20
	} else {
		goto L402
	}
L402:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v19)+180)) = v1445
	v1448 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v1433))))
	v1449 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v1433)+2)))
	v1450 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v1433)+4)))
	v1451 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v1433)+6)))
	v1455 = F_XLogReadBufferForRedo(m, l0, int32(0), v19+int32(104))
	mBase = m.M
	v1456 = m.ExcPending
	if v1456 != 0 {
		goto L20
	} else {
		goto L403
	}
L403:
	;
	if v1455 == int32(0) {
		goto L404
	} else {
		goto L405
	}
L404:
	;
	v1460 = v1433 + int32(16)
	v1461 = int32(1)
	v1463 = v1460 + v1448<<(uint(v1461)%32)
	v1466 = v1463 + v1449<<(uint(v1461)%32)
	v1468 = v1450 << (uint(v1461) % 32)
	v1469 = v1466 + v1468
	v1472 = *(*int32)(unsafe.Add(mBase, uint32(v19)+104))
	if v1472 < int32(0) {
		goto L408
	} else {
		goto L409
	}
L405:
	;
	goto L406
L406:
	;
	v1666 = *(*int32)(unsafe.Add(mBase, uint32(v19)+104))
	if v1666 == int32(0) {
		goto L6
	} else {
		goto L427
	}
L407:
	;
	v1491 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v1433))))
	v1492 = int32(2)
	F_spgPageIndexMultiDelete(m, v19+int32(108), v1490, v1460, v1491, v1492, v1492, int32(-1), int32(0))
	mBase = m.M
	v1497 = m.ExcPending
	if v1497 != 0 {
		goto L20
	} else {
		goto L411
	}
L408:
	;
	v1476 = *(*int32)(unsafe.Add(mBase, _consts[12]))
	v1482 = *(*int32)(unsafe.Add(mBase, uint32(v1476+(v1472^int32(-1))<<(uint(int32(2))%32))))
	v1490 = v1482
	goto L407
L409:
	;
	goto L410
L410:
	;
	v1484 = *(*int32)(unsafe.Add(mBase, _consts[13]))
	v1490 = v1484 + v1472<<(uint(int32(13))%32) + int32(-8192)
	goto L407
L411:
	;
	v1500 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v1433)+2)))
	v1501 = int32(3)
	F_spgPageIndexMultiDelete(m, v19+int32(108), v1490, v1463, v1500, v1501, v1501, int32(-1), int32(0))
	mBase = m.M
	v1506 = m.ExcPending
	if v1506 != 0 {
		goto L20
	} else {
		goto L412
	}
L412:
	;
	v1508 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v1433)+4)))
	if v1508 != 0 {
		goto L413
	} else {
		goto L414
	}
L413:
	;
	v1510 = v1490 + int32(24)
	v1515 = int32(0)
	goto L416
L414:
	;
	v1552 = int32(0)
	goto L415
L415:
	;
	v1570 = int32(3)
	F_spgPageIndexMultiDelete(m, v19+int32(108), v1490, v1466, v1552, v1570, v1570, int32(-1), int32(0))
	mBase = m.M
	v1575 = m.ExcPending
	if v1575 != 0 {
		goto L20
	} else {
		goto L419
	}
L416:
	;
	v1528 = int32(1)
	v1529 = v1515 << (uint(v1528) % 32)
	v1531 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v1466+v1529))))
	v1532 = int32(2)
	v1535 = int32(4)
	v1536 = v1531<<(uint(v1532)%32) + v1510 - v1535
	v1537 = *(*int32)(unsafe.Add(mBase, uint32(v1536)))
	v1539 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v1529+v1469))))
	v1544 = v1539<<(uint(v1532)%32) + v1510 - v1535
	v1545 = *(*int32)(unsafe.Add(mBase, uint32(v1544)))
	*(*int32)(unsafe.Add(mBase, uint32(v1536))) = v1545
	*(*int32)(unsafe.Add(mBase, uint32(v1544))) = v1537
	v1549 = v1515 + v1528
	v1550 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v1433)+4)))
	if base.Ui32(v1549) < base.Ui32(v1550) {
		v1515 = v1549
		goto L416
	} else {
		goto L418
	}
L417:
	;
	v1552 = v1550
	goto L415
L418:
	;
	goto L417
L419:
	;
	v1576 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v1433)+6)))
	if v1576 != 0 {
		goto L420
	} else {
		goto L421
	}
L420:
	;
	v1577 = v1469 + v1468
	v1589 = int32(0)
	goto L423
L421:
	;
	goto L422
L422:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v1490))) = base.I64_rotr(v1432, int64(32))
	v1647 = *(*int32)(unsafe.Add(mBase, uint32(v19)+104))
	F_MarkBufferDirty(m, v1647)
	mBase = m.M
	v1649 = m.ExcPending
	if v1649 != 0 {
		goto L20
	} else {
		goto L426
	}
L423:
	;
	v1602 = int32(1)
	v1603 = v1589 << (uint(v1602) % 32)
	v1605 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v1577+v1603))))
	v1611 = *(*int32)(unsafe.Add(mBase, uint32(v1605<<(uint(int32(2))%32)+(v1490+int32(24))-int32(4))))
	v1614 = v1490 + int32(4) + v1611&int32(32767)
	v1616 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v1603+(v1577+v1451<<(uint(int32(1))%32))))))
	v1619 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v1614))))
	v1622 = v1616&int32(16383) | v1619&int32(49152)
	*(*uint16)(unsafe.Add(mBase, uint32(v1614))) = uint16(v1622)
	v1625 = v1589 + v1602
	v1626 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v1433)+6)))
	if base.Ui32(v1625) < base.Ui32(v1626) {
		v1589 = v1625
		goto L423
	} else {
		goto L425
	}
L424:
	;
	goto L422
L425:
	;
	goto L424
L426:
	;
	goto L406
L427:
	;
	F_UnlockReleaseBuffer(m, v1666)
	mBase = m.M
	v1670 = m.ExcPending
	if v1670 != 0 {
		goto L20
	} else {
		goto L428
	}
L428:
	;
	goto L6
L429:
	;
	if v1676 == int32(0) {
		goto L430
	} else {
		goto L431
	}
L430:
	;
	v1682 = *(*int32)(unsafe.Add(mBase, uint32(v19)+108))
	if v1682 < int32(0) {
		goto L434
	} else {
		goto L435
	}
L431:
	;
	goto L432
L432:
	;
	v1712 = *(*int32)(unsafe.Add(mBase, uint32(v19)+108))
	if v1712 == int32(0) {
		goto L6
	} else {
		goto L439
	}
L433:
	;
	v1701 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v1671))))
	F_PageIndexMultiDelete(m, v1700, v1671+int32(12), v1701)
	mBase = m.M
	v1703 = m.ExcPending
	if v1703 != 0 {
		goto L20
	} else {
		goto L437
	}
L434:
	;
	v1686 = *(*int32)(unsafe.Add(mBase, _consts[12]))
	v1692 = *(*int32)(unsafe.Add(mBase, uint32(v1686+(v1682^int32(-1))<<(uint(int32(2))%32))))
	v1700 = v1692
	goto L433
L435:
	;
	goto L436
L436:
	;
	v1694 = *(*int32)(unsafe.Add(mBase, _consts[13]))
	v1700 = v1694 + v1682<<(uint(int32(13))%32) + int32(-8192)
	goto L433
L437:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v1700))) = base.I64_rotr(v1672, int64(32))
	v1707 = *(*int32)(unsafe.Add(mBase, uint32(v19)+108))
	F_MarkBufferDirty(m, v1707)
	mBase = m.M
	v1709 = m.ExcPending
	if v1709 != 0 {
		goto L20
	} else {
		goto L438
	}
L438:
	;
	goto L432
L439:
	;
	F_UnlockReleaseBuffer(m, v1712)
	mBase = m.M
	v1716 = m.ExcPending
	if v1716 != 0 {
		goto L20
	} else {
		goto L440
	}
L440:
	;
	goto L6
L441:
	;
	v1723 = int32(0)
	F_XLogRecGetBlockTag(m, l0, v1723, v19+int32(108), v1723, v1723)
	mBase = m.M
	v1729 = m.ExcPending
	if v1729 != 0 {
		goto L20
	} else {
		goto L444
	}
L442:
	;
	goto L443
L443:
	;
	v1745 = F_XLogReadBufferForRedo(m, l0, int32(0), v19+int32(108))
	mBase = m.M
	v1746 = m.ExcPending
	if v1746 != 0 {
		goto L20
	} else {
		goto L446
	}
L444:
	;
	v1730 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1717)+8)))
	v1731 = *(*int32)(unsafe.Add(mBase, uint32(v1717)+4))
	v1732 = *(*int32)(unsafe.Add(mBase, uint32(v19)+116))
	*(*int32)(unsafe.Add(mBase, uint32(v19)+80)) = v1732
	v1734 = *(*int64)(unsafe.Add(mBase, uint32(v19)+108))
	*(*int64)(unsafe.Add(mBase, uint32(v19)+72)) = v1734
	F_ResolveRecoveryConflictWithSnapshot(m, v1731, v1730, v19+int32(72))
	mBase = m.M
	v1739 = m.ExcPending
	if v1739 != 0 {
		goto L20
	} else {
		goto L445
	}
L445:
	;
	goto L443
L446:
	;
	if v1745 == int32(0) {
		goto L447
	} else {
		goto L448
	}
L447:
	;
	v1749 = *(*int32)(unsafe.Add(mBase, uint32(v19)+108))
	if v1749 < int32(0) {
		goto L451
	} else {
		goto L452
	}
L448:
	;
	goto L449
L449:
	;
	v2042 = *(*int32)(unsafe.Add(mBase, uint32(v19)+108))
	if v2042 == int32(0) {
		goto L6
	} else {
		goto L486
	}
L450:
	;
	v1768 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v1767)+16)))
	v1769 = v1768 + v1767
	v1770 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v1717))))
	if v1770 == int32(0) {
		goto L455
	} else {
		goto L456
	}
L451:
	;
	v1753 = *(*int32)(unsafe.Add(mBase, _consts[12]))
	v1759 = *(*int32)(unsafe.Add(mBase, uint32(v1753+(v1749^int32(-1))<<(uint(int32(2))%32))))
	v1767 = v1759
	goto L450
L452:
	;
	goto L453
L453:
	;
	v1761 = *(*int32)(unsafe.Add(mBase, _consts[13]))
	v1767 = v1761 + v1749<<(uint(int32(13))%32) + int32(-8192)
	goto L450
L454:
	;
	v1836 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v1769)+2)))
	v1837 = v1836 - v1821
	*(*uint16)(unsafe.Add(mBase, uint32(v1769)+2)) = uint16(v1837)
	v1839 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v1769)+4)))
	v1840 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v1717))))
	v1841 = v1839 + v1840
	*(*uint16)(unsafe.Add(mBase, uint32(v1769)+4)) = uint16(v1841)
	v1843 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v1717)+2)))
	if v1843 != 0 {
		goto L461
	} else {
		goto L462
	}
L455:
	;
	v1821 = int32(0)
	goto L454
L456:
	;
	goto L457
L457:
	;
	v1782 = int32(0)
	goto L458
L458:
	;
	v1795 = int32(1)
	v1798 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v1717+int32(10)+v1782<<(uint(v1795)%32)))))
	v1804 = *(*int32)(unsafe.Add(mBase, uint32(v1798<<(uint(int32(2))%32)+(v1767+int32(24))-int32(4))))
	v1807 = v1767 + v1804&int32(32767)
	v1808 = int32(0)
	*(*uint16)(unsafe.Add(mBase, uint32(v1807)+10)) = uint16(v1808)
	*(*int32)(unsafe.Add(mBase, uint32(v1807)+6)) = int32(-1)
	v1812 = *(*int32)(unsafe.Add(mBase, uint32(v1807)))
	*(*int32)(unsafe.Add(mBase, uint32(v1807))) = v1812 | int32(3)
	v1817 = v1782 + v1795
	v1818 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v1717))))
	if base.Ui32(v1817) < base.Ui32(v1818) {
		v1782 = v1817
		goto L458
	} else {
		goto L460
	}
L459:
	;
	v1821 = v1818
	goto L454
L460:
	;
	goto L459
L461:
	;
	v1844 = int32(0)
	v1845 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v1767)+12)))
	if base.Ui32(int32(25)) <= base.Ui32(v1845) {
		goto L464
	} else {
		goto L465
	}
L462:
	;
	goto L463
L463:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v1767))) = base.I64_rotr(v1718, int64(32))
	v2023 = *(*int32)(unsafe.Add(mBase, uint32(v19)+108))
	F_MarkBufferDirty(m, v2023)
	mBase = m.M
	v2025 = m.ExcPending
	if v2025 != 0 {
		goto L20
	} else {
		goto L485
	}
L464:
	;
	v1855 = int32(base.Ui32(v1845+int32(262120))>>(uint(int32(2))%32)) & int32(65535)
	goto L466
L465:
	;
	v1855 = v1844
	goto L466
L466:
	;
	v1858 = F_palloc(m, v1855<<(uint(int32(1))%32))
	mBase = m.M
	v1859 = m.ExcPending
	if v1859 != 0 {
		goto L20
	} else {
		goto L467
	}
L467:
	;
	v1860 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v1717)+2)))
	if base.Ui32(v1860) <= base.Ui32(v1855) {
		goto L468
	} else {
		goto L469
	}
L468:
	;
	v1862 = v1855 - v1860
	v1866 = (v1862 + int32(1)) & int32(3)
	if v1866 != 0 {
		goto L471
	} else {
		goto L472
	}
L469:
	;
	v1979 = v1860
	goto L470
L470:
	;
	v1994 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v1769)+4)))
	v1997 = v1855 - v1979 + int32(1)
	v1998 = v1994 - v1997
	*(*uint16)(unsafe.Add(mBase, uint32(v1769)+4)) = uint16(v1998)
	F_PageIndexMultiDelete(m, v1767, v1858, v1997)
	mBase = m.M
	v2001 = m.ExcPending
	if v2001 != 0 {
		goto L20
	} else {
		goto L483
	}
L471:
	;
	v1868 = v1860
	v1872 = v1844
	goto L474
L472:
	;
	v1895 = v1860
	goto L473
L473:
	;
	if base.Ui32(int32(3)) <= base.Ui32(v1862) {
		goto L477
	} else {
		goto L478
	}
L474:
	;
	v1883 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v1717)+2)))
	v1885 = int32(1)
	*(*uint16)(unsafe.Add(mBase, uint32(v1858+(v1868-v1883)<<(uint(v1885)%32)))) = uint16(v1868)
	v1890 = v1868 + v1885
	v1892 = v1872 + v1885
	if v1892 != v1866 {
		v1868 = v1890
		v1872 = v1892
		goto L474
	} else {
		goto L476
	}
L475:
	;
	v1895 = v1890
	goto L473
L476:
	;
	goto L475
L477:
	;
	v1913 = v1895
	goto L480
L478:
	;
	goto L479
L479:
	;
	v1977 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v1717)+2)))
	v1979 = v1977
	goto L470
L480:
	;
	v1928 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v1717)+2)))
	v1930 = int32(1)
	*(*uint16)(unsafe.Add(mBase, uint32(v1858+(v1913-v1928)<<(uint(v1930)%32)))) = uint16(v1913)
	v1935 = v1913 + v1930
	v1936 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v1717)+2)))
	*(*uint16)(unsafe.Add(mBase, uint32(v1858+(v1935-v1936)<<(uint(v1930)%32)))) = uint16(v1935)
	v1943 = v1913 + int32(2)
	v1944 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v1717)+2)))
	*(*uint16)(unsafe.Add(mBase, uint32(v1858+(v1943-v1944)<<(uint(v1930)%32)))) = uint16(v1943)
	v1951 = v1913 + int32(3)
	v1952 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v1717)+2)))
	*(*uint16)(unsafe.Add(mBase, uint32(v1858+(v1951-v1952)<<(uint(v1930)%32)))) = uint16(v1951)
	if v1951 != v1855 {
		v1913 = v1913 + int32(4)
		goto L480
	} else {
		goto L482
	}
L481:
	;
	goto L479
L482:
	;
	goto L481
L483:
	;
	F_pfree(m, v1858)
	mBase = m.M
	v2003 = m.ExcPending
	if v2003 != 0 {
		goto L20
	} else {
		goto L484
	}
L484:
	;
	goto L463
L485:
	;
	goto L449
L486:
	;
	F_UnlockReleaseBuffer(m, v2042)
	mBase = m.M
	v2046 = m.ExcPending
	if v2046 != 0 {
		goto L20
	} else {
		goto L487
	}
L487:
	;
	goto L6
L488:
	;
	m.G0 = v19 + int32(192)
	return
L489:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v19)+16)) = v130
	F_errmsg_internal(m, int32(426008), v19+int32(16))
	mBase = m.M
	v2081 = m.ExcPending
	if v2081 != 0 {
		goto L20
	} else {
		goto L490
	}
L490:
	;
	F_errfinish(m, int32(521542), int32(135), int32(356144))
	mBase = m.M
	v2086 = m.ExcPending
	if v2086 != 0 {
		goto L20
	} else {
		goto L491
	}
L491:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L492:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v19)+48)) = v479
	F_errmsg_internal(m, int32(426008), v19+int32(48))
	mBase = m.M
	v2096 = m.ExcPending
	if v2096 != 0 {
		goto L20
	} else {
		goto L493
	}
L493:
	;
	F_errfinish(m, int32(521542), int32(316), int32(433197))
	mBase = m.M
	v2101 = m.ExcPending
	if v2101 != 0 {
		goto L20
	} else {
		goto L494
	}
L494:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L495:
	;
	v2106 = *(*int32)(unsafe.Add(mBase, uint32(v741)))
	*(*int32)(unsafe.Add(mBase, uint32(v19)+32)) = int32(base.Ui32(v2106) >> (uint(int32(2)) % 32))
	F_errmsg_internal(m, int32(426008), v19+int32(32))
	mBase = m.M
	v2114 = m.ExcPending
	if v2114 != 0 {
		goto L20
	} else {
		goto L496
	}
L496:
	;
	F_errfinish(m, int32(521542), int32(397), int32(433197))
	mBase = m.M
	v2119 = m.ExcPending
	if v2119 != 0 {
		goto L20
	} else {
		goto L497
	}
L497:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L498:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v19)+64)) = v850
	F_errmsg_internal(m, int32(426008), v19-int32(-64))
	mBase = m.M
	v2129 = m.ExcPending
	if v2129 != 0 {
		goto L20
	} else {
		goto L499
	}
L499:
	;
	F_errfinish(m, int32(521542), int32(514), int32(402733))
	mBase = m.M
	v2134 = m.ExcPending
	if v2134 != 0 {
		goto L20
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
	*(*int32)(unsafe.Add(mBase, uint32(v19))) = v30
	F_errmsg_internal(m, int32(57487), v19)
	mBase = m.M
	v2142 = m.ExcPending
	if v2142 != 0 {
		goto L20
	} else {
		goto L502
	}
L502:
	;
	F_errfinish(m, int32(521542), int32(968), int32(254625))
	mBase = m.M
	v2147 = m.ExcPending
	if v2147 != 0 {
		goto L20
	} else {
		goto L503
	}
L503:
	;
	base.Wasm_trap_unreachable()
	for {
	}
}
