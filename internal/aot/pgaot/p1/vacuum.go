package p1

import (
	base "github.com/shibukawa/pgmem/internal/aot/pgaot/base"
	"unsafe"
)

func F_vacuum(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32) {
	mBase := m.M
	_ = mBase
	var v6 int32
	_ = v6
	var v26 int32
	_ = v26
	var v28 int32
	_ = v28
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
	var v44 int32
	_ = v44
	var v45 int32
	_ = v45
	var v46 int32
	_ = v46
	var v49 int32
	_ = v49
	var v58 int32
	_ = v58
	var v60 int32
	_ = v60
	var v73 int32
	_ = v73
	var v84 int32
	_ = v84
	var v85 int32
	_ = v85
	var v88 int32
	_ = v88
	var v90 int32
	_ = v90
	var v95 int32
	_ = v95
	var v96 int32
	_ = v96
	var v99 int32
	_ = v99
	var v112 int32
	_ = v112
	var v124 int32
	_ = v124
	var v136 int32
	_ = v136
	var v142 int32
	_ = v142
	var v156 int32
	_ = v156
	var v157 int32
	_ = v157
	var v160 int32
	_ = v160
	var v164 int32
	_ = v164
	var v174 int32
	_ = v174
	var v177 int32
	_ = v177
	var v178 int32
	_ = v178
	var v179 int32
	_ = v179
	var v181 int32
	_ = v181
	var v185 int32
	_ = v185
	var v191 int32
	_ = v191
	var v195 int32
	_ = v195
	var v196 int32
	_ = v196
	var v197 int32
	_ = v197
	var v198 int32
	_ = v198
	var v211 int32
	_ = v211
	var v212 int32
	_ = v212
	var v215 int32
	_ = v215
	var v216 int32
	_ = v216
	var v231 int32
	_ = v231
	var v233 int32
	_ = v233
	var v234 int32
	_ = v234
	var v248 int32
	_ = v248
	var v249 int32
	_ = v249
	var v252 int32
	_ = v252
	var v266 int32
	_ = v266
	var v267 int32
	_ = v267
	var v268 int32
	_ = v268
	var v283 int32
	_ = v283
	var v297 int32
	_ = v297
	var v298 int32
	_ = v298
	var v312 int32
	_ = v312
	var v313 int32
	_ = v313
	var v314 int32
	_ = v314
	var v329 int32
	_ = v329
	var v343 int32
	_ = v343
	var v355 int32
	_ = v355
	var v356 int32
	_ = v356
	var v371 int32
	_ = v371
	var v386 int32
	_ = v386
	var v400 int32
	_ = v400
	var v401 int32
	_ = v401
	var v402 int32
	_ = v402
	var v412 int32
	_ = v412
	var v413 int32
	_ = v413
	var v414 int32
	_ = v414
	var v416 int32
	_ = v416
	var v417 int32
	_ = v417
	var v420 int32
	_ = v420
	var v421 int32
	_ = v421
	var v431 int32
	_ = v431
	var v432 int32
	_ = v432
	var v443 int32
	_ = v443
	var v444 int32
	_ = v444
	var v447 int32
	_ = v447
	var v448 int32
	_ = v448
	var v451 int32
	_ = v451
	var v452 int32
	_ = v452
	var v459 int32
	_ = v459
	var v474 int32
	_ = v474
	var v475 int32
	_ = v475
	var v476 int32
	_ = v476
	var v477 int32
	_ = v477
	var v492 int32
	_ = v492
	var v506 int32
	_ = v506
	var v518 int32
	_ = v518
	var v529 int32
	_ = v529
	var v543 int32
	_ = v543
	var v545 int32
	_ = v545
	var v546 int32
	_ = v546
	var v549 int32
	_ = v549
	var v550 int32
	_ = v550
	var v564 int32
	_ = v564
	var v569 int32
	_ = v569
	var v570 int32
	_ = v570
	var v573 int32
	_ = v573
	var v578 int32
	_ = v578
	var v582 int32
	_ = v582
	var v584 int32
	_ = v584
	var v585 int32
	_ = v585
	var v588 int32
	_ = v588
	var v599 int32
	_ = v599
	var v600 int32
	_ = v600
	var v610 int32
	_ = v610
	var v611 int32
	_ = v611
	var v614 int32
	_ = v614
	var v615 int32
	_ = v615
	var v616 int32
	_ = v616
	var v617 int32
	_ = v617
	var v620 int32
	_ = v620
	var v633 int32
	_ = v633
	var v638 int32
	_ = v638
	var v658 int32
	_ = v658
	var v670 int32
	_ = v670
	var v671 int32
	_ = v671
	var v672 int32
	_ = v672
	var v675 int32
	_ = v675
	var v684 int32
	_ = v684
	var v685 int32
	_ = v685
	var v697 int32
	_ = v697
	var v698 int32
	_ = v698
	var v702 int32
	_ = v702
	var v703 int32
	_ = v703
	var v716 int32
	_ = v716
	var v717 int32
	_ = v717
	var v727 int32
	_ = v727
	var v729 int32
	_ = v729
	var v730 int32
	_ = v730
	var v740 int32
	_ = v740
	var v741 int32
	_ = v741
	var v742 int32
	_ = v742
	var v752 int32
	_ = v752
	var v753 int32
	_ = v753
	var v758 int32
	_ = v758
	var v761 int32
	_ = v761
	var v768 int32
	_ = v768
	var v769 int32
	_ = v769
	var v770 int32
	_ = v770
	var v771 int32
	_ = v771
	var v773 int32
	_ = v773
	var v783 int32
	_ = v783
	var v784 int32
	_ = v784
	var v787 int32
	_ = v787
	var v800 int32
	_ = v800
	var v801 int32
	_ = v801
	var v804 int32
	_ = v804
	var v808 int32
	_ = v808
	var v820 int32
	_ = v820
	var v821 int32
	_ = v821
	var v824 int32
	_ = v824
	var v825 int32
	_ = v825
	var v837 int32
	_ = v837
	var v839 int32
	_ = v839
	var v840 int32
	_ = v840
	var v850 int32
	_ = v850
	var v851 int32
	_ = v851
	var v854 int32
	_ = v854
	var v855 int32
	_ = v855
	var v868 int32
	_ = v868
	var v869 int32
	_ = v869
	var v879 int32
	_ = v879
	var v880 int32
	_ = v880
	var v885 int32
	_ = v885
	var v895 int32
	_ = v895
	var v896 int32
	_ = v896
	var v897 int32
	_ = v897
	var v908 int32
	_ = v908
	var v920 int32
	_ = v920
	var v929 int32
	_ = v929
	var v930 int32
	_ = v930
	var v931 int32
	_ = v931
	var v932 int32
	_ = v932
	var v933 int32
	_ = v933
	var v934 int32
	_ = v934
	var v936 int32
	_ = v936
	var v946 int32
	_ = v946
	var v947 int32
	_ = v947
	var v951 int32
	_ = v951
	var v955 int32
	_ = v955
	var v956 int32
	_ = v956
	var v957 int32
	_ = v957
	var v962 int32
	_ = v962
	var v964 int32
	_ = v964
	var v977 int32
	_ = v977
	var v990 int32
	_ = v990
	var v1001 int32
	_ = v1001
	var v1004 int32
	_ = v1004
	var v1006 int32
	_ = v1006
	var v1008 int32
	_ = v1008
	var v1020 int32
	_ = v1020
	var v1021 int32
	_ = v1021
	var v1022 int32
	_ = v1022
	var v1023 int32
	_ = v1023
	var v1024 int32
	_ = v1024
	var v1025 int32
	_ = v1025
	var v1026 int32
	_ = v1026
	var v1027 int32
	_ = v1027
	var v1029 int32
	_ = v1029
	var v1032 int32
	_ = v1032
	var v1042 int32
	_ = v1042
	var v1049 int32
	_ = v1049
	var v1061 int32
	_ = v1061
	var v1063 int32
	_ = v1063
	var v1076 int32
	_ = v1076
	var v1077 int32
	_ = v1077
	var v1098 int32
	_ = v1098
	var v1105 int32
	_ = v1105
	var v1109 int32
	_ = v1109
	var v1110 int32
	_ = v1110
	var v1113 int32
	_ = v1113
	var v1114 int32
	_ = v1114
	var v1124 int64
	_ = v1124
	var v1126 int64
	_ = v1126
	var v1128 int64
	_ = v1128
	var v1130 int64
	_ = v1130
	var v1132 int64
	_ = v1132
	var v1134 int64
	_ = v1134
	var v1136 int64
	_ = v1136
	var v1138 int64
	_ = v1138
	var v1140 int32
	_ = v1140
	var v1141 int32
	_ = v1141
	var v1144 int32
	_ = v1144
	var v1146 int32
	_ = v1146
	var v1151 int32
	_ = v1151
	var v1164 int32
	_ = v1164
	var v1174 int32
	_ = v1174
	var v1175 int32
	_ = v1175
	var v1186 int32
	_ = v1186
	var v1187 int32
	_ = v1187
	var v1188 int32
	_ = v1188
	var v1189 int32
	_ = v1189
	var v1199 int32
	_ = v1199
	var v1201 int32
	_ = v1201
	var v1202 int32
	_ = v1202
	var v1215 int32
	_ = v1215
	var v1226 int32
	_ = v1226
	var v1237 int32
	_ = v1237
	var v1248 int32
	_ = v1248
	var v1252 int32
	_ = v1252
	var v1257 int32
	_ = v1257
	var v1258 int32
	_ = v1258
	var v1290 int32
	_ = v1290
	var v1301 int32
	_ = v1301
	var v1312 int32
	_ = v1312
	var v1313 int32
	_ = v1313
	var v1328 int32
	_ = v1328
	var v1337 int32
	_ = v1337
	var v1358 int32
	_ = v1358
	var v1384 int32
	_ = v1384
	var v1385 int64
	_ = v1385
	var v1389 int32
	_ = v1389
	var v1391 int32
	_ = v1391
	var v1392 int32
	_ = v1392
	var v1395 int32
	_ = v1395
	var v1397 int32
	_ = v1397
	var v1399 int32
	_ = v1399
	var v1400 int32
	_ = v1400
	var v1401 int32
	_ = v1401
	var v1402 int32
	_ = v1402
	var v1403 int32
	_ = v1403
	var v1404 int32
	_ = v1404
	var v1405 int32
	_ = v1405
	var v1406 int32
	_ = v1406
	var v1407 int32
	_ = v1407
	var v1408 int32
	_ = v1408
	var v1410 int32
	_ = v1410
	v6 = int32(0)
	v26 = m.G0
	v28 = v26 - int32(352)
	m.G0 = v28
	v37 = v6
	v38 = v6
	v39 = v6
	v40 = v6
	v41 = v6
	v42 = v6
	v43 = v6
	v44 = v6
	v45 = v6
	v46 = int32(-1)
	v49 = v6
	goto L1
L1:
	;
	goto L4
L2:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L3:
	;
	goto L2
L4:
	;
	if v46 == int32(1) {
		goto L8
	} else {
		goto L9
	}
L5:
	;
	goto L3
L6:
	;
	v1384 = int32(m.ExcTag)
	v1385 = int64(m.ExcVals[0])
	m.ExcPending = 0
	if v1384 == int32(0) {
		goto L193
	} else {
		goto L194
	}
L7:
	;
	if v1032 == int32(0) {
		goto L152
	} else {
		goto L153
	}
L8:
	;
	v1020 = v37
	v1021 = v38
	v1022 = v39
	v1023 = v40
	v1024 = v41
	v1025 = v42
	v1026 = v43
	v1027 = v44
	v1029 = v45
	v1032 = v49
	goto L7
L9:
	;
	goto L10
L10:
	;
	v58 = *(*int32)(unsafe.Add(mBase, uint32(l1)))
	v60 = v58 & int32(1)
	if v60 != 0 {
		goto L12
	} else {
		goto L13
	}
L11:
	;
	*(*uint8)(unsafe.Add(mBase, uint32(v28)+315)) = uint8(v96)
	v99 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_vacuum[0])))
	if v99 != 0 {
		goto L19
	} else {
		goto L20
	}
L12:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v28)+320)) = v37
	*(*int32)(unsafe.Add(mBase, uint32(v28)+316)) = v38
	*(*int32)(unsafe.Add(mBase, uint32(v28)+324)) = v45
	*(*int32)(unsafe.Add(mBase, uint32(v28)+328)) = v40
	*(*int32)(unsafe.Add(mBase, uint32(v28)+332)) = v41
	*(*int32)(unsafe.Add(mBase, uint32(v28)+336)) = v39
	*(*int32)(unsafe.Add(mBase, uint32(v28)+340)) = v42
	*(*int32)(unsafe.Add(mBase, uint32(v28)+344)) = v43
	*(*int32)(unsafe.Add(mBase, uint32(v28)+348)) = v44
	F_PreventInTransactionBlock(m, l4, int32(_a_F_vacuum_0))
	mBase = m.M
	v73 = m.ExcPending
	if v73 != 0 {
		goto L6
	} else {
		goto L15
	}
L13:
	;
	goto L14
L14:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v28)+320)) = v37
	*(*int32)(unsafe.Add(mBase, uint32(v28)+316)) = v38
	*(*int32)(unsafe.Add(mBase, uint32(v28)+324)) = v45
	*(*int32)(unsafe.Add(mBase, uint32(v28)+328)) = v40
	*(*int32)(unsafe.Add(mBase, uint32(v28)+332)) = v41
	*(*int32)(unsafe.Add(mBase, uint32(v28)+336)) = v39
	*(*int32)(unsafe.Add(mBase, uint32(v28)+340)) = v42
	*(*int32)(unsafe.Add(mBase, uint32(v28)+344)) = v43
	*(*int32)(unsafe.Add(mBase, uint32(v28)+348)) = v44
	v84 = *(*int32)(unsafe.Add(mBase, _c_F_vacuum[1]))
	v85 = *(*int32)(unsafe.Add(mBase, uint32(v84)+24))
	if base.Ui32(v85) <= base.Ui32(int32(1)) {
		goto L16
	} else {
		goto L17
	}
L15:
	;
	v96 = int32(0)
	goto L11
L16:
	;
	v88 = int32(1)
	v90 = *(*int32)(unsafe.Add(mBase, uint32(v84)+28))
	v95 = l4 ^ v88 | base.B2i32(v88 < v90)
	goto L18
L17:
	;
	v95 = int32(1)
	goto L18
L18:
	;
	v96 = v95
	goto L11
L19:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v28)+320)) = v37
	*(*int32)(unsafe.Add(mBase, uint32(v28)+316)) = v38
	*(*int32)(unsafe.Add(mBase, uint32(v28)+324)) = v45
	*(*int32)(unsafe.Add(mBase, uint32(v28)+328)) = v40
	*(*int32)(unsafe.Add(mBase, uint32(v28)+332)) = v41
	*(*int32)(unsafe.Add(mBase, uint32(v28)+336)) = v39
	*(*int32)(unsafe.Add(mBase, uint32(v28)+340)) = v42
	*(*int32)(unsafe.Add(mBase, uint32(v28)+344)) = v43
	*(*int32)(unsafe.Add(mBase, uint32(v28)+348)) = v44
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v112 = m.ExcPending
	if v112 != 0 {
		goto L6
	} else {
		goto L22
	}
L20:
	;
	goto L21
L21:
	;
	v157 = *(*int32)(unsafe.Add(mBase, uint32(l1)))
	if v157&int32(1024) != 0 {
		goto L30
	} else {
		goto L31
	}
L22:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v28)+320)) = v37
	*(*int32)(unsafe.Add(mBase, uint32(v28)+316)) = v38
	*(*int32)(unsafe.Add(mBase, uint32(v28)+324)) = v45
	*(*int32)(unsafe.Add(mBase, uint32(v28)+328)) = v40
	*(*int32)(unsafe.Add(mBase, uint32(v28)+332)) = v41
	*(*int32)(unsafe.Add(mBase, uint32(v28)+336)) = v39
	*(*int32)(unsafe.Add(mBase, uint32(v28)+340)) = v42
	*(*int32)(unsafe.Add(mBase, uint32(v28)+344)) = v43
	*(*int32)(unsafe.Add(mBase, uint32(v28)+348)) = v44
	F_errcode(m, int32(1088))
	mBase = m.M
	v124 = m.ExcPending
	if v124 != 0 {
		goto L6
	} else {
		goto L23
	}
L23:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v28)+320)) = v37
	*(*int32)(unsafe.Add(mBase, uint32(v28)+316)) = v38
	*(*int32)(unsafe.Add(mBase, uint32(v28)+324)) = v45
	*(*int32)(unsafe.Add(mBase, uint32(v28)+328)) = v40
	*(*int32)(unsafe.Add(mBase, uint32(v28)+332)) = v41
	*(*int32)(unsafe.Add(mBase, uint32(v28)+336)) = v39
	*(*int32)(unsafe.Add(mBase, uint32(v28)+340)) = v42
	*(*int32)(unsafe.Add(mBase, uint32(v28)+344)) = v43
	*(*int32)(unsafe.Add(mBase, uint32(v28)+348)) = v44
	if v60 != 0 {
		goto L24
	} else {
		goto L25
	}
L24:
	;
	v136 = int32(_a_F_vacuum_0)
	goto L26
L25:
	;
	v136 = int32(_a_F_vacuum_1)
	goto L26
L26:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v28)+64)) = v136
	F_errmsg(m, int32(_a_F_vacuum_2), v28-int32(-64))
	mBase = m.M
	v142 = m.ExcPending
	if v142 != 0 {
		goto L6
	} else {
		goto L27
	}
L27:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v28)+320)) = v37
	*(*int32)(unsafe.Add(mBase, uint32(v28)+316)) = v38
	*(*int32)(unsafe.Add(mBase, uint32(v28)+324)) = v45
	*(*int32)(unsafe.Add(mBase, uint32(v28)+328)) = v40
	*(*int32)(unsafe.Add(mBase, uint32(v28)+332)) = v41
	*(*int32)(unsafe.Add(mBase, uint32(v28)+336)) = v39
	*(*int32)(unsafe.Add(mBase, uint32(v28)+340)) = v42
	*(*int32)(unsafe.Add(mBase, uint32(v28)+344)) = v43
	*(*int32)(unsafe.Add(mBase, uint32(v28)+348)) = v44
	F_errfinish(m, int32(_a_F_vacuum_3), int32(530), int32(_a_F_vacuum_4))
	mBase = m.M
	v156 = m.ExcPending
	if v156 != 0 {
		goto L6
	} else {
		goto L28
	}
L28:
	;
	goto L3
L29:
	;
	v946 = int32(1)
	v947 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l1))))
	if v947&v946 != 0 {
		v962 = v946
		goto L131
	} else {
		goto L132
	}
L30:
	;
	v929 = v39
	v930 = v40
	v931 = v41
	v932 = v42
	v933 = v43
	v934 = v44
	v936 = l0
	goto L29
L31:
	;
	goto L32
L32:
	;
	if l0 != 0 {
		goto L33
	} else {
		goto L34
	}
L33:
	;
	v160 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	if v160 <= int32(0) {
		goto L36
	} else {
		goto L37
	}
L34:
	;
	goto L35
L35:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v28)+320)) = v37
	*(*int32)(unsafe.Add(mBase, uint32(v28)+316)) = v38
	*(*int32)(unsafe.Add(mBase, uint32(v28)+324)) = v45
	*(*int32)(unsafe.Add(mBase, uint32(v28)+328)) = v40
	*(*int32)(unsafe.Add(mBase, uint32(v28)+332)) = v41
	*(*int32)(unsafe.Add(mBase, uint32(v28)+336)) = v39
	*(*int32)(unsafe.Add(mBase, uint32(v28)+340)) = v42
	*(*int32)(unsafe.Add(mBase, uint32(v28)+344)) = v43
	*(*int32)(unsafe.Add(mBase, uint32(v28)+348)) = v44
	v716 = F_table_open(m, int32(1259), int32(1))
	mBase = m.M
	v717 = m.ExcPending
	if v717 != 0 {
		goto L6
	} else {
		goto L102
	}
L36:
	;
	v929 = v39
	v930 = v40
	v931 = v41
	v932 = v42
	v933 = v43
	v934 = v44
	v936 = int32(0)
	goto L29
L37:
	;
	goto L38
L38:
	;
	v164 = int32(0)
	v174 = v39
	v177 = v42
	v178 = v43
	v179 = v44
	v181 = v164
	v185 = v164
	goto L39
L39:
	;
	v191 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	v195 = *(*int32)(unsafe.Add(mBase, uint32(v191+v181<<(uint(int32(2))%32))))
	v196 = *(*int32)(unsafe.Add(mBase, uint32(v195)+8))
	if v196 != 0 {
		goto L42
	} else {
		goto L43
	}
L40:
	;
	v929 = v697
	v930 = v40
	v931 = v41
	v932 = v670
	v933 = v671
	v934 = v672
	v936 = v697
	goto L29
L41:
	;
	v684 = int32(_a_F_vacuum_5)
	v685 = *(*int32)(unsafe.Add(mBase, _c_F_vacuum[2]))
	*(*int32)(unsafe.Add(mBase, _c_F_vacuum[2])) = l3
	*(*int32)(unsafe.Add(mBase, uint32(v28)+316)) = v38
	*(*int32)(unsafe.Add(mBase, uint32(v28)+320)) = v37
	*(*int32)(unsafe.Add(mBase, uint32(v28)+324)) = v45
	*(*int32)(unsafe.Add(mBase, uint32(v28)+328)) = v40
	*(*int32)(unsafe.Add(mBase, uint32(v28)+332)) = v41
	*(*int32)(unsafe.Add(mBase, uint32(v28)+336)) = v174
	*(*int32)(unsafe.Add(mBase, uint32(v28)+340)) = v670
	*(*int32)(unsafe.Add(mBase, uint32(v28)+344)) = v671
	*(*int32)(unsafe.Add(mBase, uint32(v28)+348)) = v672
	v697 = F_list_concat(m, v185, v675)
	mBase = m.M
	v698 = m.ExcPending
	if v698 != 0 {
		goto L6
	} else {
		goto L100
	}
L42:
	;
	v197 = int32(_a_F_vacuum_5)
	v198 = *(*int32)(unsafe.Add(mBase, _c_F_vacuum[2]))
	*(*int32)(unsafe.Add(mBase, _c_F_vacuum[2])) = l3
	*(*int32)(unsafe.Add(mBase, uint32(v28)+316)) = v38
	*(*int32)(unsafe.Add(mBase, uint32(v28)+320)) = v37
	*(*int32)(unsafe.Add(mBase, uint32(v28)+324)) = v45
	*(*int32)(unsafe.Add(mBase, uint32(v28)+328)) = v40
	*(*int32)(unsafe.Add(mBase, uint32(v28)+332)) = v41
	*(*int32)(unsafe.Add(mBase, uint32(v28)+336)) = v174
	*(*int32)(unsafe.Add(mBase, uint32(v28)+340)) = v177
	*(*int32)(unsafe.Add(mBase, uint32(v28)+344)) = v178
	*(*int32)(unsafe.Add(mBase, uint32(v28)+348)) = v179
	v211 = F_lappend(m, int32(0), v195)
	mBase = m.M
	v212 = m.ExcPending
	if v212 != 0 {
		goto L6
	} else {
		goto L45
	}
L43:
	;
	goto L44
L44:
	;
	v215 = *(*int32)(unsafe.Add(mBase, uint32(l1)))
	v216 = *(*int32)(unsafe.Add(mBase, uint32(v195)+4))
	*(*int32)(unsafe.Add(mBase, uint32(v28)+320)) = v37
	*(*int32)(unsafe.Add(mBase, uint32(v28)+316)) = v38
	*(*int32)(unsafe.Add(mBase, uint32(v28)+324)) = v45
	*(*int32)(unsafe.Add(mBase, uint32(v28)+328)) = v40
	*(*int32)(unsafe.Add(mBase, uint32(v28)+332)) = v41
	*(*int32)(unsafe.Add(mBase, uint32(v28)+336)) = v174
	*(*int32)(unsafe.Add(mBase, uint32(v28)+340)) = v177
	*(*int32)(unsafe.Add(mBase, uint32(v28)+344)) = v178
	*(*int32)(unsafe.Add(mBase, uint32(v28)+348)) = v179
	v231 = int32(0)
	v233 = F_RangeVarGetRelidExtended(m, v216, int32(1), int32(base.Ui32(v215)>>(uint(int32(3))%32))&int32(4), v231, v231)
	mBase = m.M
	v234 = m.ExcPending
	if v234 != 0 {
		goto L6
	} else {
		goto L46
	}
L45:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_vacuum[2])) = v198
	v670 = v177
	v671 = v178
	v672 = v211
	v675 = v211
	goto L41
L46:
	;
	if v233 == int32(0) {
		goto L47
	} else {
		goto L48
	}
L47:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v28)+320)) = v37
	*(*int32)(unsafe.Add(mBase, uint32(v28)+316)) = v38
	*(*int32)(unsafe.Add(mBase, uint32(v28)+324)) = v45
	*(*int32)(unsafe.Add(mBase, uint32(v28)+328)) = v40
	*(*int32)(unsafe.Add(mBase, uint32(v28)+332)) = v41
	*(*int32)(unsafe.Add(mBase, uint32(v28)+336)) = v174
	*(*int32)(unsafe.Add(mBase, uint32(v28)+340)) = v177
	*(*int32)(unsafe.Add(mBase, uint32(v28)+344)) = v178
	*(*int32)(unsafe.Add(mBase, uint32(v28)+348)) = v179
	v248 = F_errstart(m, int32(19), int32(0))
	mBase = m.M
	v249 = m.ExcPending
	if v249 != 0 {
		goto L6
	} else {
		goto L50
	}
L48:
	;
	goto L49
L49:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v28)+320)) = v37
	*(*int32)(unsafe.Add(mBase, uint32(v28)+316)) = v38
	*(*int32)(unsafe.Add(mBase, uint32(v28)+324)) = v45
	*(*int32)(unsafe.Add(mBase, uint32(v28)+328)) = v40
	*(*int32)(unsafe.Add(mBase, uint32(v28)+332)) = v41
	*(*int32)(unsafe.Add(mBase, uint32(v28)+336)) = v174
	*(*int32)(unsafe.Add(mBase, uint32(v28)+340)) = v177
	*(*int32)(unsafe.Add(mBase, uint32(v28)+344)) = v178
	*(*int32)(unsafe.Add(mBase, uint32(v28)+348)) = v179
	v355 = F_SearchSysCache1(m, int32(57), base.I64_extend_i32_u(v233))
	mBase = m.M
	v356 = m.ExcPending
	if v356 != 0 {
		goto L6
	} else {
		goto L62
	}
L50:
	;
	if v215&int32(1) != 0 {
		goto L51
	} else {
		goto L52
	}
L51:
	;
	v252 = int32(0)
	if v248 == v252 {
		v670 = v177
		v671 = v178
		v672 = v179
		v675 = v252
		goto L41
	} else {
		goto L54
	}
L52:
	;
	goto L53
L53:
	;
	v298 = int32(0)
	if v248 == v298 {
		v670 = v177
		v671 = v178
		v672 = v179
		v675 = v298
		goto L41
	} else {
		goto L58
	}
L54:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v28)+320)) = v37
	*(*int32)(unsafe.Add(mBase, uint32(v28)+316)) = v38
	*(*int32)(unsafe.Add(mBase, uint32(v28)+324)) = v45
	*(*int32)(unsafe.Add(mBase, uint32(v28)+328)) = v40
	*(*int32)(unsafe.Add(mBase, uint32(v28)+332)) = v41
	*(*int32)(unsafe.Add(mBase, uint32(v28)+336)) = v174
	*(*int32)(unsafe.Add(mBase, uint32(v28)+340)) = v177
	*(*int32)(unsafe.Add(mBase, uint32(v28)+344)) = v178
	*(*int32)(unsafe.Add(mBase, uint32(v28)+348)) = v179
	F_errcode(m, int32(50463045))
	mBase = m.M
	v266 = m.ExcPending
	if v266 != 0 {
		goto L6
	} else {
		goto L55
	}
L55:
	;
	v267 = *(*int32)(unsafe.Add(mBase, uint32(v195)+4))
	v268 = *(*int32)(unsafe.Add(mBase, uint32(v267)+12))
	*(*int32)(unsafe.Add(mBase, uint32(v28)+320)) = v37
	*(*int32)(unsafe.Add(mBase, uint32(v28)+316)) = v38
	*(*int32)(unsafe.Add(mBase, uint32(v28)+324)) = v45
	*(*int32)(unsafe.Add(mBase, uint32(v28)+328)) = v40
	*(*int32)(unsafe.Add(mBase, uint32(v28)+332)) = v41
	*(*int32)(unsafe.Add(mBase, uint32(v28)+336)) = v174
	*(*int32)(unsafe.Add(mBase, uint32(v28)+340)) = v177
	*(*int32)(unsafe.Add(mBase, uint32(v28)+344)) = v178
	*(*int32)(unsafe.Add(mBase, uint32(v28)+348)) = v179
	*(*int32)(unsafe.Add(mBase, uint32(v28)+96)) = v268
	F_errmsg(m, int32(_a_F_vacuum_6), v28+int32(96))
	mBase = m.M
	v283 = m.ExcPending
	if v283 != 0 {
		goto L6
	} else {
		goto L56
	}
L56:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v28)+320)) = v37
	*(*int32)(unsafe.Add(mBase, uint32(v28)+316)) = v38
	*(*int32)(unsafe.Add(mBase, uint32(v28)+324)) = v45
	*(*int32)(unsafe.Add(mBase, uint32(v28)+328)) = v40
	*(*int32)(unsafe.Add(mBase, uint32(v28)+332)) = v41
	*(*int32)(unsafe.Add(mBase, uint32(v28)+336)) = v174
	*(*int32)(unsafe.Add(mBase, uint32(v28)+340)) = v177
	*(*int32)(unsafe.Add(mBase, uint32(v28)+344)) = v178
	*(*int32)(unsafe.Add(mBase, uint32(v28)+348)) = v179
	F_errfinish(m, int32(_a_F_vacuum_3), int32(937), int32(_a_F_vacuum_7))
	mBase = m.M
	v297 = m.ExcPending
	if v297 != 0 {
		goto L6
	} else {
		goto L57
	}
L57:
	;
	v670 = v177
	v671 = v178
	v672 = v179
	v675 = v252
	goto L41
L58:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v28)+320)) = v37
	*(*int32)(unsafe.Add(mBase, uint32(v28)+316)) = v38
	*(*int32)(unsafe.Add(mBase, uint32(v28)+324)) = v45
	*(*int32)(unsafe.Add(mBase, uint32(v28)+328)) = v40
	*(*int32)(unsafe.Add(mBase, uint32(v28)+332)) = v41
	*(*int32)(unsafe.Add(mBase, uint32(v28)+336)) = v174
	*(*int32)(unsafe.Add(mBase, uint32(v28)+340)) = v177
	*(*int32)(unsafe.Add(mBase, uint32(v28)+344)) = v178
	*(*int32)(unsafe.Add(mBase, uint32(v28)+348)) = v179
	F_errcode(m, int32(50463045))
	mBase = m.M
	v312 = m.ExcPending
	if v312 != 0 {
		goto L6
	} else {
		goto L59
	}
L59:
	;
	v313 = *(*int32)(unsafe.Add(mBase, uint32(v195)+4))
	v314 = *(*int32)(unsafe.Add(mBase, uint32(v313)+12))
	*(*int32)(unsafe.Add(mBase, uint32(v28)+320)) = v37
	*(*int32)(unsafe.Add(mBase, uint32(v28)+316)) = v38
	*(*int32)(unsafe.Add(mBase, uint32(v28)+324)) = v45
	*(*int32)(unsafe.Add(mBase, uint32(v28)+328)) = v40
	*(*int32)(unsafe.Add(mBase, uint32(v28)+332)) = v41
	*(*int32)(unsafe.Add(mBase, uint32(v28)+336)) = v174
	*(*int32)(unsafe.Add(mBase, uint32(v28)+340)) = v177
	*(*int32)(unsafe.Add(mBase, uint32(v28)+344)) = v178
	*(*int32)(unsafe.Add(mBase, uint32(v28)+348)) = v179
	*(*int32)(unsafe.Add(mBase, uint32(v28)+80)) = v314
	F_errmsg(m, int32(_a_F_vacuum_8), v28+int32(80))
	mBase = m.M
	v329 = m.ExcPending
	if v329 != 0 {
		goto L6
	} else {
		goto L60
	}
L60:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v28)+320)) = v37
	*(*int32)(unsafe.Add(mBase, uint32(v28)+316)) = v38
	*(*int32)(unsafe.Add(mBase, uint32(v28)+324)) = v45
	*(*int32)(unsafe.Add(mBase, uint32(v28)+328)) = v40
	*(*int32)(unsafe.Add(mBase, uint32(v28)+332)) = v41
	*(*int32)(unsafe.Add(mBase, uint32(v28)+336)) = v174
	*(*int32)(unsafe.Add(mBase, uint32(v28)+340)) = v177
	*(*int32)(unsafe.Add(mBase, uint32(v28)+344)) = v178
	*(*int32)(unsafe.Add(mBase, uint32(v28)+348)) = v179
	F_errfinish(m, int32(_a_F_vacuum_3), int32(942), int32(_a_F_vacuum_7))
	mBase = m.M
	v343 = m.ExcPending
	if v343 != 0 {
		goto L6
	} else {
		goto L61
	}
L61:
	;
	v670 = v177
	v671 = v178
	v672 = v179
	v675 = v298
	goto L41
L62:
	;
	if v355 == int32(0) {
		goto L63
	} else {
		goto L64
	}
L63:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v28)+320)) = v37
	*(*int32)(unsafe.Add(mBase, uint32(v28)+316)) = v38
	*(*int32)(unsafe.Add(mBase, uint32(v28)+324)) = v45
	*(*int32)(unsafe.Add(mBase, uint32(v28)+328)) = v40
	*(*int32)(unsafe.Add(mBase, uint32(v28)+332)) = v41
	*(*int32)(unsafe.Add(mBase, uint32(v28)+336)) = v174
	*(*int32)(unsafe.Add(mBase, uint32(v28)+340)) = v177
	*(*int32)(unsafe.Add(mBase, uint32(v28)+344)) = v178
	*(*int32)(unsafe.Add(mBase, uint32(v28)+348)) = v179
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v371 = m.ExcPending
	if v371 != 0 {
		goto L6
	} else {
		goto L66
	}
L64:
	;
	goto L65
L65:
	;
	v401 = *(*int32)(unsafe.Add(mBase, uint32(v355)+16))
	v402 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v401)+22)))
	*(*int32)(unsafe.Add(mBase, uint32(v28)+320)) = v37
	*(*int32)(unsafe.Add(mBase, uint32(v28)+316)) = v38
	*(*int32)(unsafe.Add(mBase, uint32(v28)+324)) = v45
	*(*int32)(unsafe.Add(mBase, uint32(v28)+328)) = v40
	*(*int32)(unsafe.Add(mBase, uint32(v28)+332)) = v41
	*(*int32)(unsafe.Add(mBase, uint32(v28)+336)) = v174
	*(*int32)(unsafe.Add(mBase, uint32(v28)+340)) = v177
	*(*int32)(unsafe.Add(mBase, uint32(v28)+344)) = v178
	*(*int32)(unsafe.Add(mBase, uint32(v28)+348)) = v179
	v412 = v402 + v401
	v413 = F_vacuum_is_permitted_for_relation(m, v233, v412, v215)
	mBase = m.M
	v414 = m.ExcPending
	if v414 != 0 {
		goto L6
	} else {
		goto L69
	}
L66:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v28)+320)) = v37
	*(*int32)(unsafe.Add(mBase, uint32(v28)+316)) = v38
	*(*int32)(unsafe.Add(mBase, uint32(v28)+324)) = v45
	*(*int32)(unsafe.Add(mBase, uint32(v28)+328)) = v40
	*(*int32)(unsafe.Add(mBase, uint32(v28)+332)) = v41
	*(*int32)(unsafe.Add(mBase, uint32(v28)+336)) = v174
	*(*int32)(unsafe.Add(mBase, uint32(v28)+340)) = v177
	*(*int32)(unsafe.Add(mBase, uint32(v28)+344)) = v178
	*(*int32)(unsafe.Add(mBase, uint32(v28)+348)) = v179
	*(*int32)(unsafe.Add(mBase, uint32(v28)+112)) = v233
	F_errmsg_internal(m, int32(_a_F_vacuum_9), v28+int32(112))
	mBase = m.M
	v386 = m.ExcPending
	if v386 != 0 {
		goto L6
	} else {
		goto L67
	}
L67:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v28)+320)) = v37
	*(*int32)(unsafe.Add(mBase, uint32(v28)+316)) = v38
	*(*int32)(unsafe.Add(mBase, uint32(v28)+324)) = v45
	*(*int32)(unsafe.Add(mBase, uint32(v28)+328)) = v40
	*(*int32)(unsafe.Add(mBase, uint32(v28)+332)) = v41
	*(*int32)(unsafe.Add(mBase, uint32(v28)+336)) = v174
	*(*int32)(unsafe.Add(mBase, uint32(v28)+340)) = v177
	*(*int32)(unsafe.Add(mBase, uint32(v28)+344)) = v178
	*(*int32)(unsafe.Add(mBase, uint32(v28)+348)) = v179
	F_errfinish(m, int32(_a_F_vacuum_3), int32(952), int32(_a_F_vacuum_7))
	mBase = m.M
	v400 = m.ExcPending
	if v400 != 0 {
		goto L6
	} else {
		goto L68
	}
L68:
	;
	goto L3
L69:
	;
	if v413 != 0 {
		goto L70
	} else {
		goto L71
	}
L70:
	;
	v416 = int32(_a_F_vacuum_5)
	v417 = *(*int32)(unsafe.Add(mBase, _c_F_vacuum[2]))
	*(*int32)(unsafe.Add(mBase, _c_F_vacuum[2])) = l3
	v420 = *(*int32)(unsafe.Add(mBase, uint32(v195)+12))
	v421 = *(*int32)(unsafe.Add(mBase, uint32(v195)+4))
	*(*int32)(unsafe.Add(mBase, uint32(v28)+320)) = v37
	*(*int32)(unsafe.Add(mBase, uint32(v28)+316)) = v38
	*(*int32)(unsafe.Add(mBase, uint32(v28)+324)) = v45
	*(*int32)(unsafe.Add(mBase, uint32(v28)+328)) = v40
	*(*int32)(unsafe.Add(mBase, uint32(v28)+332)) = v41
	*(*int32)(unsafe.Add(mBase, uint32(v28)+336)) = v174
	*(*int32)(unsafe.Add(mBase, uint32(v28)+340)) = v177
	*(*int32)(unsafe.Add(mBase, uint32(v28)+344)) = v178
	*(*int32)(unsafe.Add(mBase, uint32(v28)+348)) = v179
	v431 = F_makeVacuumRelation(m, v421, v233, v420)
	mBase = m.M
	v432 = m.ExcPending
	if v432 != 0 {
		goto L6
	} else {
		goto L73
	}
L71:
	;
	v447 = v178
	v448 = int32(0)
	goto L72
L72:
	;
	v451 = *(*int32)(unsafe.Add(mBase, uint32(v195)+4))
	v452 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v451)+16)))
	if v215&int32(1) == int32(0) {
		goto L76
	} else {
		goto L77
	}
L73:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v28)+320)) = v37
	*(*int32)(unsafe.Add(mBase, uint32(v28)+316)) = v38
	*(*int32)(unsafe.Add(mBase, uint32(v28)+324)) = v45
	*(*int32)(unsafe.Add(mBase, uint32(v28)+328)) = v40
	*(*int32)(unsafe.Add(mBase, uint32(v28)+332)) = v41
	*(*int32)(unsafe.Add(mBase, uint32(v28)+336)) = v174
	*(*int32)(unsafe.Add(mBase, uint32(v28)+340)) = v177
	*(*int32)(unsafe.Add(mBase, uint32(v28)+344)) = v178
	*(*int32)(unsafe.Add(mBase, uint32(v28)+348)) = v179
	v443 = F_lappend(m, int32(0), v431)
	mBase = m.M
	v444 = m.ExcPending
	if v444 != 0 {
		goto L6
	} else {
		goto L74
	}
L74:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_vacuum[2])) = v417
	v447 = v443
	v448 = v443
	goto L72
L75:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v28)+320)) = v37
	*(*int32)(unsafe.Add(mBase, uint32(v28)+316)) = v38
	*(*int32)(unsafe.Add(mBase, uint32(v28)+324)) = v45
	*(*int32)(unsafe.Add(mBase, uint32(v28)+328)) = v40
	*(*int32)(unsafe.Add(mBase, uint32(v28)+332)) = v41
	*(*int32)(unsafe.Add(mBase, uint32(v28)+336)) = v174
	*(*int32)(unsafe.Add(mBase, uint32(v28)+340)) = v633
	*(*int32)(unsafe.Add(mBase, uint32(v28)+344)) = v447
	*(*int32)(unsafe.Add(mBase, uint32(v28)+348)) = v179
	F_UnlockRelationOid(m, v233, int32(1))
	mBase = m.M
	v658 = m.ExcPending
	if v658 != 0 {
		goto L6
	} else {
		goto L99
	}
L76:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v28)+320)) = v37
	*(*int32)(unsafe.Add(mBase, uint32(v28)+316)) = v38
	*(*int32)(unsafe.Add(mBase, uint32(v28)+324)) = v45
	*(*int32)(unsafe.Add(mBase, uint32(v28)+328)) = v40
	*(*int32)(unsafe.Add(mBase, uint32(v28)+332)) = v41
	*(*int32)(unsafe.Add(mBase, uint32(v28)+336)) = v174
	*(*int32)(unsafe.Add(mBase, uint32(v28)+340)) = v177
	*(*int32)(unsafe.Add(mBase, uint32(v28)+344)) = v447
	*(*int32)(unsafe.Add(mBase, uint32(v28)+348)) = v179
	F_ReleaseCatCache(m, v355)
	mBase = m.M
	v529 = m.ExcPending
	if v529 != 0 {
		goto L6
	} else {
		goto L86
	}
L77:
	;
	v459 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v412)+119)))
	if v452&int32(1)|base.B2i32(v459 != int32(112)) != 0 {
		goto L76
	} else {
		goto L78
	}
L78:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v28)+320)) = v37
	*(*int32)(unsafe.Add(mBase, uint32(v28)+316)) = v38
	*(*int32)(unsafe.Add(mBase, uint32(v28)+324)) = v45
	*(*int32)(unsafe.Add(mBase, uint32(v28)+328)) = v40
	*(*int32)(unsafe.Add(mBase, uint32(v28)+332)) = v41
	*(*int32)(unsafe.Add(mBase, uint32(v28)+336)) = v174
	*(*int32)(unsafe.Add(mBase, uint32(v28)+340)) = v177
	*(*int32)(unsafe.Add(mBase, uint32(v28)+344)) = v447
	*(*int32)(unsafe.Add(mBase, uint32(v28)+348)) = v179
	v474 = F_errstart(m, int32(19), int32(0))
	mBase = m.M
	v475 = m.ExcPending
	if v475 != 0 {
		goto L6
	} else {
		goto L79
	}
L79:
	;
	if v474 != 0 {
		goto L80
	} else {
		goto L81
	}
L80:
	;
	v476 = *(*int32)(unsafe.Add(mBase, uint32(v195)+4))
	v477 = *(*int32)(unsafe.Add(mBase, uint32(v476)+12))
	*(*int32)(unsafe.Add(mBase, uint32(v28)+320)) = v37
	*(*int32)(unsafe.Add(mBase, uint32(v28)+316)) = v38
	*(*int32)(unsafe.Add(mBase, uint32(v28)+324)) = v45
	*(*int32)(unsafe.Add(mBase, uint32(v28)+328)) = v40
	*(*int32)(unsafe.Add(mBase, uint32(v28)+332)) = v41
	*(*int32)(unsafe.Add(mBase, uint32(v28)+336)) = v174
	*(*int32)(unsafe.Add(mBase, uint32(v28)+340)) = v177
	*(*int32)(unsafe.Add(mBase, uint32(v28)+344)) = v447
	*(*int32)(unsafe.Add(mBase, uint32(v28)+348)) = v179
	*(*int32)(unsafe.Add(mBase, uint32(v28)+128)) = v477
	F_errmsg(m, int32(_a_F_vacuum_10), v28+int32(128))
	mBase = m.M
	v492 = m.ExcPending
	if v492 != 0 {
		goto L6
	} else {
		goto L83
	}
L81:
	;
	goto L82
L82:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v28)+320)) = v37
	*(*int32)(unsafe.Add(mBase, uint32(v28)+316)) = v38
	*(*int32)(unsafe.Add(mBase, uint32(v28)+324)) = v45
	*(*int32)(unsafe.Add(mBase, uint32(v28)+328)) = v40
	*(*int32)(unsafe.Add(mBase, uint32(v28)+332)) = v41
	*(*int32)(unsafe.Add(mBase, uint32(v28)+336)) = v174
	*(*int32)(unsafe.Add(mBase, uint32(v28)+340)) = v177
	*(*int32)(unsafe.Add(mBase, uint32(v28)+344)) = v447
	*(*int32)(unsafe.Add(mBase, uint32(v28)+348)) = v179
	F_ReleaseCatCache(m, v355)
	mBase = m.M
	v518 = m.ExcPending
	if v518 != 0 {
		goto L6
	} else {
		goto L85
	}
L83:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v28)+320)) = v37
	*(*int32)(unsafe.Add(mBase, uint32(v28)+316)) = v38
	*(*int32)(unsafe.Add(mBase, uint32(v28)+324)) = v45
	*(*int32)(unsafe.Add(mBase, uint32(v28)+328)) = v40
	*(*int32)(unsafe.Add(mBase, uint32(v28)+332)) = v41
	*(*int32)(unsafe.Add(mBase, uint32(v28)+336)) = v174
	*(*int32)(unsafe.Add(mBase, uint32(v28)+340)) = v177
	*(*int32)(unsafe.Add(mBase, uint32(v28)+344)) = v447
	*(*int32)(unsafe.Add(mBase, uint32(v28)+348)) = v179
	F_errfinish(m, int32(_a_F_vacuum_3), int32(978), int32(_a_F_vacuum_7))
	mBase = m.M
	v506 = m.ExcPending
	if v506 != 0 {
		goto L6
	} else {
		goto L84
	}
L84:
	;
	goto L82
L85:
	;
	v633 = v177
	v638 = v448
	goto L75
L86:
	;
	if v452&int32(1) == int32(0) {
		v633 = v177
		v638 = v448
		goto L75
	} else {
		goto L87
	}
L87:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v28)+320)) = v37
	*(*int32)(unsafe.Add(mBase, uint32(v28)+316)) = v38
	*(*int32)(unsafe.Add(mBase, uint32(v28)+324)) = v45
	*(*int32)(unsafe.Add(mBase, uint32(v28)+328)) = v40
	*(*int32)(unsafe.Add(mBase, uint32(v28)+332)) = v41
	*(*int32)(unsafe.Add(mBase, uint32(v28)+336)) = v174
	*(*int32)(unsafe.Add(mBase, uint32(v28)+340)) = v177
	*(*int32)(unsafe.Add(mBase, uint32(v28)+344)) = v447
	*(*int32)(unsafe.Add(mBase, uint32(v28)+348)) = v179
	v543 = int32(0)
	v545 = F_find_all_inheritors(m, v233, v543, v543)
	mBase = m.M
	v546 = m.ExcPending
	if v546 != 0 {
		goto L6
	} else {
		goto L88
	}
L88:
	;
	if v545 == int32(0) {
		v633 = v177
		v638 = v448
		goto L75
	} else {
		goto L89
	}
L89:
	;
	v549 = int32(0)
	v550 = *(*int32)(unsafe.Add(mBase, uint32(v545)+4))
	if v550 <= v549 {
		v633 = v177
		v638 = v448
		goto L75
	} else {
		goto L90
	}
L90:
	;
	v564 = v177
	v569 = v448
	v570 = v550
	v573 = v549
	goto L91
L91:
	;
	v578 = *(*int32)(unsafe.Add(mBase, uint32(v545)+12))
	v582 = *(*int32)(unsafe.Add(mBase, uint32(v578+v573<<(uint(int32(2))%32))))
	if v233 != v582 {
		goto L93
	} else {
		goto L94
	}
L92:
	;
	v633 = v615
	v638 = v616
	goto L75
L93:
	;
	v584 = int32(_a_F_vacuum_5)
	v585 = *(*int32)(unsafe.Add(mBase, _c_F_vacuum[2]))
	*(*int32)(unsafe.Add(mBase, _c_F_vacuum[2])) = l3
	v588 = *(*int32)(unsafe.Add(mBase, uint32(v195)+12))
	*(*int32)(unsafe.Add(mBase, uint32(v28)+320)) = v37
	*(*int32)(unsafe.Add(mBase, uint32(v28)+316)) = v38
	*(*int32)(unsafe.Add(mBase, uint32(v28)+324)) = v45
	*(*int32)(unsafe.Add(mBase, uint32(v28)+328)) = v40
	*(*int32)(unsafe.Add(mBase, uint32(v28)+332)) = v41
	*(*int32)(unsafe.Add(mBase, uint32(v28)+336)) = v174
	*(*int32)(unsafe.Add(mBase, uint32(v28)+340)) = v564
	*(*int32)(unsafe.Add(mBase, uint32(v28)+344)) = v447
	*(*int32)(unsafe.Add(mBase, uint32(v28)+348)) = v179
	v599 = F_makeVacuumRelation(m, int32(0), v582, v588)
	mBase = m.M
	v600 = m.ExcPending
	if v600 != 0 {
		goto L6
	} else {
		goto L96
	}
L94:
	;
	v615 = v564
	v616 = v569
	v617 = v570
	goto L95
L95:
	;
	v620 = v573 + int32(1)
	if v620 < v617 {
		v564 = v615
		v569 = v616
		v570 = v617
		v573 = v620
		goto L91
	} else {
		goto L98
	}
L96:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v28)+320)) = v37
	*(*int32)(unsafe.Add(mBase, uint32(v28)+316)) = v38
	*(*int32)(unsafe.Add(mBase, uint32(v28)+324)) = v45
	*(*int32)(unsafe.Add(mBase, uint32(v28)+328)) = v40
	*(*int32)(unsafe.Add(mBase, uint32(v28)+332)) = v41
	*(*int32)(unsafe.Add(mBase, uint32(v28)+336)) = v174
	*(*int32)(unsafe.Add(mBase, uint32(v28)+340)) = v564
	*(*int32)(unsafe.Add(mBase, uint32(v28)+344)) = v447
	*(*int32)(unsafe.Add(mBase, uint32(v28)+348)) = v179
	v610 = F_lappend(m, v569, v599)
	mBase = m.M
	v611 = m.ExcPending
	if v611 != 0 {
		goto L6
	} else {
		goto L97
	}
L97:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_vacuum[2])) = v585
	v614 = *(*int32)(unsafe.Add(mBase, uint32(v545)+4))
	v615 = v610
	v616 = v610
	v617 = v614
	goto L95
L98:
	;
	goto L92
L99:
	;
	v670 = v633
	v671 = v447
	v672 = v179
	v675 = v638
	goto L41
L100:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_vacuum[2])) = v685
	v702 = v181 + int32(1)
	v703 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	if v702 < v703 {
		v174 = v697
		v177 = v670
		v178 = v671
		v179 = v672
		v181 = v702
		v185 = v697
		goto L39
	} else {
		goto L101
	}
L101:
	;
	goto L40
L102:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v28)+320)) = v37
	*(*int32)(unsafe.Add(mBase, uint32(v28)+316)) = v38
	*(*int32)(unsafe.Add(mBase, uint32(v28)+324)) = v45
	*(*int32)(unsafe.Add(mBase, uint32(v28)+328)) = v40
	*(*int32)(unsafe.Add(mBase, uint32(v28)+332)) = v41
	*(*int32)(unsafe.Add(mBase, uint32(v28)+336)) = v39
	*(*int32)(unsafe.Add(mBase, uint32(v28)+340)) = v42
	*(*int32)(unsafe.Add(mBase, uint32(v28)+344)) = v43
	*(*int32)(unsafe.Add(mBase, uint32(v28)+348)) = v44
	v727 = int32(0)
	v729 = F_table_beginscan_catalog(m, v716, v727, v727)
	mBase = m.M
	v730 = m.ExcPending
	if v730 != 0 {
		goto L6
	} else {
		goto L103
	}
L103:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v28)+320)) = v37
	*(*int32)(unsafe.Add(mBase, uint32(v28)+316)) = v38
	*(*int32)(unsafe.Add(mBase, uint32(v28)+324)) = v45
	*(*int32)(unsafe.Add(mBase, uint32(v28)+328)) = v40
	*(*int32)(unsafe.Add(mBase, uint32(v28)+332)) = v41
	*(*int32)(unsafe.Add(mBase, uint32(v28)+336)) = v39
	*(*int32)(unsafe.Add(mBase, uint32(v28)+340)) = v42
	*(*int32)(unsafe.Add(mBase, uint32(v28)+344)) = v43
	*(*int32)(unsafe.Add(mBase, uint32(v28)+348)) = v44
	v740 = F_heap_getnext(m, v729)
	mBase = m.M
	v741 = m.ExcPending
	if v741 != 0 {
		goto L6
	} else {
		goto L104
	}
L104:
	;
	v742 = int32(0)
	if v740 != 0 {
		goto L105
	} else {
		goto L106
	}
L105:
	;
	v752 = v40
	v753 = v41
	v758 = v742
	v761 = v740
	goto L108
L106:
	;
	v879 = v40
	v880 = v41
	v885 = v742
	goto L107
L107:
	;
	v895 = *(*int32)(unsafe.Add(mBase, uint32(v729)))
	v896 = *(*int32)(unsafe.Add(mBase, uint32(v895)+188))
	v897 = *(*int32)(unsafe.Add(mBase, uint32(v896)+12))
	*(*int32)(unsafe.Add(mBase, uint32(v28)+320)) = v37
	*(*int32)(unsafe.Add(mBase, uint32(v28)+316)) = v38
	*(*int32)(unsafe.Add(mBase, uint32(v28)+324)) = v45
	*(*int32)(unsafe.Add(mBase, uint32(v28)+328)) = v879
	*(*int32)(unsafe.Add(mBase, uint32(v28)+332)) = v880
	*(*int32)(unsafe.Add(mBase, uint32(v28)+336)) = v39
	*(*int32)(unsafe.Add(mBase, uint32(v28)+340)) = v42
	*(*int32)(unsafe.Add(mBase, uint32(v28)+344)) = v43
	*(*int32)(unsafe.Add(mBase, uint32(v28)+348)) = v44
	m.T0[v897].(func(*base.Module, int32))(m, v729)
	mBase = m.M
	v908 = m.ExcPending
	if v908 != 0 {
		goto L6
	} else {
		goto L129
	}
L108:
	;
	v768 = *(*int32)(unsafe.Add(mBase, uint32(v761)+16))
	v769 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v768)+22)))
	v770 = v768 + v769
	v771 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v770)+119)))
	v773 = v771 - int32(109)
	if base.B2i32(base.Ui32(int32(5)) < base.Ui32(v773))|base.B2i32(int32(1)<<(uint(v773)%32)&int32(41) == int32(0)) != 0 {
		v854 = v753
		v855 = v758
		goto L110
	} else {
		goto L111
	}
L109:
	;
	v879 = v868
	v880 = v854
	v885 = v855
	goto L107
L110:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v28)+320)) = v37
	*(*int32)(unsafe.Add(mBase, uint32(v28)+316)) = v38
	*(*int32)(unsafe.Add(mBase, uint32(v28)+324)) = v45
	*(*int32)(unsafe.Add(mBase, uint32(v28)+328)) = v752
	*(*int32)(unsafe.Add(mBase, uint32(v28)+332)) = v854
	*(*int32)(unsafe.Add(mBase, uint32(v28)+336)) = v39
	*(*int32)(unsafe.Add(mBase, uint32(v28)+340)) = v42
	*(*int32)(unsafe.Add(mBase, uint32(v28)+344)) = v43
	*(*int32)(unsafe.Add(mBase, uint32(v28)+348)) = v44
	v868 = F_heap_getnext(m, v729)
	mBase = m.M
	v869 = m.ExcPending
	if v869 != 0 {
		goto L6
	} else {
		goto L127
	}
L111:
	;
	v783 = *(*int32)(unsafe.Add(mBase, uint32(v770)))
	v784 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v770)+118)))
	if v784 == int32(116) {
		goto L112
	} else {
		goto L113
	}
L112:
	;
	v787 = *(*int32)(unsafe.Add(mBase, uint32(v770)+68))
	*(*int32)(unsafe.Add(mBase, uint32(v28)+320)) = v37
	*(*int32)(unsafe.Add(mBase, uint32(v28)+316)) = v38
	*(*int32)(unsafe.Add(mBase, uint32(v28)+324)) = v45
	*(*int32)(unsafe.Add(mBase, uint32(v28)+328)) = v752
	*(*int32)(unsafe.Add(mBase, uint32(v28)+332)) = v753
	*(*int32)(unsafe.Add(mBase, uint32(v28)+336)) = v39
	*(*int32)(unsafe.Add(mBase, uint32(v28)+340)) = v42
	*(*int32)(unsafe.Add(mBase, uint32(v28)+344)) = v43
	*(*int32)(unsafe.Add(mBase, uint32(v28)+348)) = v44
	v800 = *(*int32)(unsafe.Add(mBase, _c_F_vacuum[3]))
	if v800 != 0 {
		goto L117
	} else {
		goto L118
	}
L113:
	;
	goto L114
L114:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v28)+320)) = v37
	*(*int32)(unsafe.Add(mBase, uint32(v28)+316)) = v38
	*(*int32)(unsafe.Add(mBase, uint32(v28)+324)) = v45
	*(*int32)(unsafe.Add(mBase, uint32(v28)+328)) = v752
	*(*int32)(unsafe.Add(mBase, uint32(v28)+332)) = v753
	*(*int32)(unsafe.Add(mBase, uint32(v28)+336)) = v39
	*(*int32)(unsafe.Add(mBase, uint32(v28)+340)) = v42
	*(*int32)(unsafe.Add(mBase, uint32(v28)+344)) = v43
	*(*int32)(unsafe.Add(mBase, uint32(v28)+348)) = v44
	v820 = F_vacuum_is_permitted_for_relation(m, v783, v770, v157)
	mBase = m.M
	v821 = m.ExcPending
	if v821 != 0 {
		goto L6
	} else {
		goto L123
	}
L115:
	;
	if v808 == int32(0) {
		v854 = v753
		v855 = v758
		goto L110
	} else {
		goto L122
	}
L116:
	;
	goto L115
L117:
	;
	v801 = int32(1)
	if v787 == v800 {
		v808 = v801
		goto L116
	} else {
		goto L120
	}
L118:
	;
	goto L119
L119:
	;
	v808 = int32(0)
	goto L116
L120:
	;
	v804 = *(*int32)(unsafe.Add(mBase, _c_F_vacuum[4]))
	if v804 == v787 {
		v808 = v801
		goto L116
	} else {
		goto L121
	}
L121:
	;
	goto L119
L122:
	;
	goto L114
L123:
	;
	if v820 == int32(0) {
		v854 = v753
		v855 = v758
		goto L110
	} else {
		goto L124
	}
L124:
	;
	v824 = int32(_a_F_vacuum_5)
	v825 = *(*int32)(unsafe.Add(mBase, _c_F_vacuum[2]))
	*(*int32)(unsafe.Add(mBase, _c_F_vacuum[2])) = l3
	*(*int32)(unsafe.Add(mBase, uint32(v28)+316)) = v38
	*(*int32)(unsafe.Add(mBase, uint32(v28)+320)) = v37
	*(*int32)(unsafe.Add(mBase, uint32(v28)+324)) = v45
	*(*int32)(unsafe.Add(mBase, uint32(v28)+328)) = v752
	*(*int32)(unsafe.Add(mBase, uint32(v28)+332)) = v753
	*(*int32)(unsafe.Add(mBase, uint32(v28)+336)) = v39
	*(*int32)(unsafe.Add(mBase, uint32(v28)+340)) = v42
	*(*int32)(unsafe.Add(mBase, uint32(v28)+344)) = v43
	*(*int32)(unsafe.Add(mBase, uint32(v28)+348)) = v44
	v837 = int32(0)
	v839 = F_makeVacuumRelation(m, v837, v783, v837)
	mBase = m.M
	v840 = m.ExcPending
	if v840 != 0 {
		goto L6
	} else {
		goto L125
	}
L125:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v28)+320)) = v37
	*(*int32)(unsafe.Add(mBase, uint32(v28)+316)) = v38
	*(*int32)(unsafe.Add(mBase, uint32(v28)+324)) = v45
	*(*int32)(unsafe.Add(mBase, uint32(v28)+328)) = v752
	*(*int32)(unsafe.Add(mBase, uint32(v28)+332)) = v753
	*(*int32)(unsafe.Add(mBase, uint32(v28)+336)) = v39
	*(*int32)(unsafe.Add(mBase, uint32(v28)+340)) = v42
	*(*int32)(unsafe.Add(mBase, uint32(v28)+344)) = v43
	*(*int32)(unsafe.Add(mBase, uint32(v28)+348)) = v44
	v850 = F_lappend(m, v758, v839)
	mBase = m.M
	v851 = m.ExcPending
	if v851 != 0 {
		goto L6
	} else {
		goto L126
	}
L126:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_vacuum[2])) = v825
	v854 = v850
	v855 = v850
	goto L110
L127:
	;
	if v868 != 0 {
		v752 = v868
		v753 = v854
		v758 = v855
		v761 = v868
		goto L108
	} else {
		goto L128
	}
L128:
	;
	goto L109
L129:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v28)+320)) = v37
	*(*int32)(unsafe.Add(mBase, uint32(v28)+316)) = v38
	*(*int32)(unsafe.Add(mBase, uint32(v28)+324)) = v45
	*(*int32)(unsafe.Add(mBase, uint32(v28)+328)) = v879
	*(*int32)(unsafe.Add(mBase, uint32(v28)+332)) = v880
	*(*int32)(unsafe.Add(mBase, uint32(v28)+336)) = v39
	*(*int32)(unsafe.Add(mBase, uint32(v28)+340)) = v42
	*(*int32)(unsafe.Add(mBase, uint32(v28)+344)) = v43
	*(*int32)(unsafe.Add(mBase, uint32(v28)+348)) = v44
	F_relation_close(m, v716, int32(1))
	mBase = m.M
	v920 = m.ExcPending
	if v920 != 0 {
		goto L6
	} else {
		goto L130
	}
L130:
	;
	v929 = v39
	v930 = v879
	v931 = v880
	v932 = v42
	v933 = v43
	v934 = v44
	v936 = v885
	goto L29
L131:
	;
	*(*uint8)(unsafe.Add(mBase, uint32(v28)+314)) = uint8(v962)
	v964 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v28)+314)))
	if v964 == int32(1) {
		goto L139
	} else {
		goto L140
	}
L132:
	;
	v951 = *(*int32)(unsafe.Add(mBase, _c_F_vacuum[5]))
	if v951 == int32(4) {
		v962 = v946
		goto L131
	} else {
		goto L133
	}
L133:
	;
	v955 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v28)+315)))
	if v955 != 0 {
		v962 = int32(0)
		goto L131
	} else {
		goto L134
	}
L134:
	;
	if v936 != 0 {
		goto L135
	} else {
		goto L136
	}
L135:
	;
	v956 = int32(1)
	v957 = *(*int32)(unsafe.Add(mBase, uint32(v936)+4))
	if v956 < v957 {
		v962 = v956
		goto L131
	} else {
		goto L138
	}
L136:
	;
	goto L137
L137:
	;
	v962 = int32(0)
	goto L131
L138:
	;
	goto L137
L139:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v28)+320)) = v37
	*(*int32)(unsafe.Add(mBase, uint32(v28)+316)) = v38
	*(*int32)(unsafe.Add(mBase, uint32(v28)+324)) = v936
	*(*int32)(unsafe.Add(mBase, uint32(v28)+328)) = v930
	*(*int32)(unsafe.Add(mBase, uint32(v28)+332)) = v931
	*(*int32)(unsafe.Add(mBase, uint32(v28)+336)) = v929
	*(*int32)(unsafe.Add(mBase, uint32(v28)+340)) = v932
	*(*int32)(unsafe.Add(mBase, uint32(v28)+344)) = v933
	*(*int32)(unsafe.Add(mBase, uint32(v28)+348)) = v934
	v977 = *(*int32)(unsafe.Add(mBase, _c_F_vacuum[6]))
	goto L142
L140:
	;
	goto L141
L141:
	;
	v1004 = *(*int32)(unsafe.Add(mBase, _c_F_vacuum[7]))
	v1006 = *(*int32)(unsafe.Add(mBase, _c_F_vacuum[8]))
	goto L148
L142:
	;
	if v977 != int32(0) {
		goto L143
	} else {
		goto L144
	}
L143:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v28)+320)) = v37
	*(*int32)(unsafe.Add(mBase, uint32(v28)+316)) = v38
	*(*int32)(unsafe.Add(mBase, uint32(v28)+324)) = v936
	*(*int32)(unsafe.Add(mBase, uint32(v28)+328)) = v930
	*(*int32)(unsafe.Add(mBase, uint32(v28)+332)) = v931
	*(*int32)(unsafe.Add(mBase, uint32(v28)+336)) = v929
	*(*int32)(unsafe.Add(mBase, uint32(v28)+340)) = v932
	*(*int32)(unsafe.Add(mBase, uint32(v28)+344)) = v933
	*(*int32)(unsafe.Add(mBase, uint32(v28)+348)) = v934
	F_PopActiveSnapshot(m)
	mBase = m.M
	v990 = m.ExcPending
	if v990 != 0 {
		goto L6
	} else {
		goto L146
	}
L144:
	;
	goto L145
L145:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v28)+320)) = v37
	*(*int32)(unsafe.Add(mBase, uint32(v28)+316)) = v38
	*(*int32)(unsafe.Add(mBase, uint32(v28)+324)) = v936
	*(*int32)(unsafe.Add(mBase, uint32(v28)+328)) = v930
	*(*int32)(unsafe.Add(mBase, uint32(v28)+332)) = v931
	*(*int32)(unsafe.Add(mBase, uint32(v28)+336)) = v929
	*(*int32)(unsafe.Add(mBase, uint32(v28)+340)) = v932
	*(*int32)(unsafe.Add(mBase, uint32(v28)+344)) = v933
	*(*int32)(unsafe.Add(mBase, uint32(v28)+348)) = v934
	F_CommitTransactionCommand(m)
	mBase = m.M
	v1001 = m.ExcPending
	if v1001 != 0 {
		goto L6
	} else {
		goto L147
	}
L146:
	;
	goto L145
L147:
	;
	goto L141
L148:
	;
	v1008 = v28 + int32(144)
	*(*int32)(unsafe.Add(mBase, uint32(v1008)+4)) = int32(1)
	*(*int32)(unsafe.Add(mBase, uint32(v1008))) = v28 + int32(140)
	goto L151
L149:
	;
	v1020 = v1006
	v1021 = v1004
	v1022 = v929
	v1023 = v930
	v1024 = v931
	v1025 = v932
	v1026 = v933
	v1027 = v934
	v1029 = v936
	v1032 = int32(0)
	goto L7
L151:
	;
	goto L149
L152:
	;
	v1042 = int32(1)
	*(*uint8)(unsafe.Add(mBase, _c_F_vacuum[0])) = uint8(v1042)
	*(*int32)(unsafe.Add(mBase, _c_F_vacuum[8])) = v28 + int32(144)
	v1049 = int32(0)
	*(*uint8)(unsafe.Add(mBase, _c_F_vacuum[9])) = uint8(v1049)
	*(*int32)(unsafe.Add(mBase, uint32(v28)+316)) = v1021
	*(*int32)(unsafe.Add(mBase, uint32(v28)+320)) = v1020
	*(*int32)(unsafe.Add(mBase, uint32(v28)+324)) = v1029
	*(*int32)(unsafe.Add(mBase, uint32(v28)+328)) = v1023
	*(*int32)(unsafe.Add(mBase, uint32(v28)+332)) = v1024
	*(*int32)(unsafe.Add(mBase, uint32(v28)+336)) = v1022
	*(*int32)(unsafe.Add(mBase, uint32(v28)+340)) = v1025
	*(*int32)(unsafe.Add(mBase, uint32(v28)+344)) = v1026
	*(*int32)(unsafe.Add(mBase, uint32(v28)+348)) = v1027
	F_VacuumUpdateCosts(m)
	mBase = m.M
	v1061 = m.ExcPending
	if v1061 != 0 {
		goto L6
	} else {
		goto L155
	}
L153:
	;
	goto L154
L154:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_vacuum[7])) = v1021
	*(*int32)(unsafe.Add(mBase, _c_F_vacuum[8])) = v1020
	v1337 = int32(0)
	*(*uint8)(unsafe.Add(mBase, _c_F_vacuum[0])) = uint8(v1337)
	*(*uint8)(unsafe.Add(mBase, _c_F_vacuum[10])) = uint8(v1337)
	*(*uint8)(unsafe.Add(mBase, _c_F_vacuum[9])) = uint8(v1337)
	*(*int32)(unsafe.Add(mBase, _c_F_vacuum[11])) = v1337
	*(*int32)(unsafe.Add(mBase, uint32(v28)+316)) = v1021
	*(*int32)(unsafe.Add(mBase, uint32(v28)+320)) = v1020
	*(*int32)(unsafe.Add(mBase, uint32(v28)+324)) = v1029
	*(*int32)(unsafe.Add(mBase, uint32(v28)+328)) = v1023
	*(*int32)(unsafe.Add(mBase, uint32(v28)+332)) = v1024
	*(*int32)(unsafe.Add(mBase, uint32(v28)+336)) = v1022
	*(*int32)(unsafe.Add(mBase, uint32(v28)+340)) = v1025
	*(*int32)(unsafe.Add(mBase, uint32(v28)+344)) = v1026
	*(*int32)(unsafe.Add(mBase, uint32(v28)+348)) = v1027
	F_pg_re_throw(m)
	mBase = m.M
	v1358 = m.ExcPending
	if v1358 != 0 {
		goto L6
	} else {
		goto L192
	}
L155:
	;
	v1063 = int32(0)
	*(*int32)(unsafe.Add(mBase, _c_F_vacuum[12])) = v1063
	*(*int32)(unsafe.Add(mBase, _c_F_vacuum[11])) = v1063
	*(*int32)(unsafe.Add(mBase, _c_F_vacuum[13])) = v1063
	*(*int32)(unsafe.Add(mBase, _c_F_vacuum[14])) = v1063
	if v1029 == v1063 {
		goto L156
	} else {
		goto L157
	}
L156:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_vacuum[8])) = v1020
	*(*int32)(unsafe.Add(mBase, _c_F_vacuum[7])) = v1021
	v1290 = int32(0)
	*(*uint8)(unsafe.Add(mBase, _c_F_vacuum[10])) = uint8(v1290)
	*(*uint8)(unsafe.Add(mBase, _c_F_vacuum[0])) = uint8(v1290)
	*(*uint8)(unsafe.Add(mBase, _c_F_vacuum[9])) = uint8(v1290)
	*(*int32)(unsafe.Add(mBase, _c_F_vacuum[11])) = v1290
	v1301 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v28)+314)))
	if v1301 != 0 {
		goto L184
	} else {
		goto L185
	}
L157:
	;
	v1076 = int32(0)
	v1077 = *(*int32)(unsafe.Add(mBase, uint32(v1029)+4))
	if v1077 <= v1076 {
		goto L156
	} else {
		goto L158
	}
L158:
	;
	v1098 = v1076
	goto L159
L159:
	;
	v1105 = *(*int32)(unsafe.Add(mBase, uint32(v1029)+12))
	v1109 = *(*int32)(unsafe.Add(mBase, uint32(v1105+v1098<<(uint(int32(2))%32))))
	v1110 = *(*int32)(unsafe.Add(mBase, uint32(l1)))
	if v1110&int32(1) != 0 {
		goto L163
	} else {
		goto L164
	}
L160:
	;
	goto L156
L161:
	;
	v1257 = v1098 + int32(1)
	v1258 = *(*int32)(unsafe.Add(mBase, uint32(v1029)+4))
	if v1257 < v1258 {
		v1098 = v1257
		goto L159
	} else {
		goto L183
	}
L162:
	;
	v1252 = int32(0)
	*(*uint8)(unsafe.Add(mBase, _c_F_vacuum[9])) = uint8(v1252)
	goto L161
L163:
	;
	v1113 = *(*int32)(unsafe.Add(mBase, uint32(v1109)+4))
	v1114 = *(*int32)(unsafe.Add(mBase, uint32(v1109)+8))
	*(*int32)(unsafe.Add(mBase, uint32(v28)+320)) = v1020
	*(*int32)(unsafe.Add(mBase, uint32(v28)+316)) = v1021
	*(*int32)(unsafe.Add(mBase, uint32(v28)+324)) = v1029
	*(*int32)(unsafe.Add(mBase, uint32(v28)+328)) = v1023
	*(*int32)(unsafe.Add(mBase, uint32(v28)+332)) = v1024
	*(*int32)(unsafe.Add(mBase, uint32(v28)+336)) = v1022
	*(*int32)(unsafe.Add(mBase, uint32(v28)+340)) = v1025
	*(*int32)(unsafe.Add(mBase, uint32(v28)+344)) = v1026
	*(*int32)(unsafe.Add(mBase, uint32(v28)+348)) = v1027
	v1124 = *(*int64)(unsafe.Add(mBase, uint32(l1)+56))
	*(*int64)(unsafe.Add(mBase, uint32(v28)+56)) = v1124
	v1126 = *(*int64)(unsafe.Add(mBase, uint32(l1)+48))
	*(*int64)(unsafe.Add(mBase, uint32(v28)+48)) = v1126
	v1128 = *(*int64)(unsafe.Add(mBase, uint32(l1)+40))
	*(*int64)(unsafe.Add(mBase, uint32(v28)+40)) = v1128
	v1130 = *(*int64)(unsafe.Add(mBase, uint32(l1)+32))
	*(*int64)(unsafe.Add(mBase, uint32(v28)+32)) = v1130
	v1132 = *(*int64)(unsafe.Add(mBase, uint32(l1)+24))
	*(*int64)(unsafe.Add(mBase, uint32(v28)+24)) = v1132
	v1134 = *(*int64)(unsafe.Add(mBase, uint32(l1)+16))
	*(*int64)(unsafe.Add(mBase, uint32(v28)+16)) = v1134
	v1136 = *(*int64)(unsafe.Add(mBase, uint32(l1)+8))
	*(*int64)(unsafe.Add(mBase, uint32(v28)+8)) = v1136
	v1138 = *(*int64)(unsafe.Add(mBase, uint32(l1)))
	*(*int64)(unsafe.Add(mBase, uint32(v28))) = v1138
	v1140 = F_vacuum_rel(m, v1114, v1113, v28, l2, l4)
	mBase = m.M
	v1141 = m.ExcPending
	if v1141 != 0 {
		goto L6
	} else {
		goto L166
	}
L164:
	;
	v1146 = v1110
	goto L165
L165:
	;
	if v1146&int32(2) == int32(0) {
		goto L162
	} else {
		goto L168
	}
L166:
	;
	if v1140 == int32(0) {
		goto L161
	} else {
		goto L167
	}
L167:
	;
	v1144 = *(*int32)(unsafe.Add(mBase, uint32(l1)))
	v1146 = v1144
	goto L165
L168:
	;
	v1151 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v28)+314)))
	if v1151 == int32(1) {
		goto L169
	} else {
		goto L170
	}
L169:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v28)+320)) = v1020
	*(*int32)(unsafe.Add(mBase, uint32(v28)+316)) = v1021
	*(*int32)(unsafe.Add(mBase, uint32(v28)+324)) = v1029
	*(*int32)(unsafe.Add(mBase, uint32(v28)+328)) = v1023
	*(*int32)(unsafe.Add(mBase, uint32(v28)+332)) = v1024
	*(*int32)(unsafe.Add(mBase, uint32(v28)+336)) = v1022
	*(*int32)(unsafe.Add(mBase, uint32(v28)+340)) = v1025
	*(*int32)(unsafe.Add(mBase, uint32(v28)+344)) = v1026
	*(*int32)(unsafe.Add(mBase, uint32(v28)+348)) = v1027
	F_StartTransactionCommand(m)
	mBase = m.M
	v1164 = m.ExcPending
	if v1164 != 0 {
		goto L6
	} else {
		goto L172
	}
L170:
	;
	goto L171
L171:
	;
	v1187 = *(*int32)(unsafe.Add(mBase, uint32(v1109)+12))
	v1188 = *(*int32)(unsafe.Add(mBase, uint32(v1109)+4))
	v1189 = *(*int32)(unsafe.Add(mBase, uint32(v1109)+8))
	*(*int32)(unsafe.Add(mBase, uint32(v28)+316)) = v1021
	*(*int32)(unsafe.Add(mBase, uint32(v28)+320)) = v1020
	*(*int32)(unsafe.Add(mBase, uint32(v28)+324)) = v1029
	*(*int32)(unsafe.Add(mBase, uint32(v28)+328)) = v1023
	*(*int32)(unsafe.Add(mBase, uint32(v28)+332)) = v1024
	*(*int32)(unsafe.Add(mBase, uint32(v28)+336)) = v1022
	*(*int32)(unsafe.Add(mBase, uint32(v28)+340)) = v1025
	*(*int32)(unsafe.Add(mBase, uint32(v28)+344)) = v1026
	*(*int32)(unsafe.Add(mBase, uint32(v28)+348)) = v1027
	v1199 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v28)+315)))
	F_analyze_rel(m, v1189, v1188, l1, v1187, v1199, l2)
	mBase = m.M
	v1201 = m.ExcPending
	if v1201 != 0 {
		goto L6
	} else {
		goto L175
	}
L172:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v28)+320)) = v1020
	*(*int32)(unsafe.Add(mBase, uint32(v28)+316)) = v1021
	*(*int32)(unsafe.Add(mBase, uint32(v28)+324)) = v1029
	*(*int32)(unsafe.Add(mBase, uint32(v28)+328)) = v1023
	*(*int32)(unsafe.Add(mBase, uint32(v28)+332)) = v1024
	*(*int32)(unsafe.Add(mBase, uint32(v28)+336)) = v1022
	*(*int32)(unsafe.Add(mBase, uint32(v28)+340)) = v1025
	*(*int32)(unsafe.Add(mBase, uint32(v28)+344)) = v1026
	*(*int32)(unsafe.Add(mBase, uint32(v28)+348)) = v1027
	v1174 = F_GetTransactionSnapshot(m)
	mBase = m.M
	v1175 = m.ExcPending
	if v1175 != 0 {
		goto L6
	} else {
		goto L173
	}
L173:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v28)+320)) = v1020
	*(*int32)(unsafe.Add(mBase, uint32(v28)+316)) = v1021
	*(*int32)(unsafe.Add(mBase, uint32(v28)+324)) = v1029
	*(*int32)(unsafe.Add(mBase, uint32(v28)+328)) = v1023
	*(*int32)(unsafe.Add(mBase, uint32(v28)+332)) = v1024
	*(*int32)(unsafe.Add(mBase, uint32(v28)+336)) = v1022
	*(*int32)(unsafe.Add(mBase, uint32(v28)+340)) = v1025
	*(*int32)(unsafe.Add(mBase, uint32(v28)+344)) = v1026
	*(*int32)(unsafe.Add(mBase, uint32(v28)+348)) = v1027
	F_PushActiveSnapshot(m, v1174)
	mBase = m.M
	v1186 = m.ExcPending
	if v1186 != 0 {
		goto L6
	} else {
		goto L174
	}
L174:
	;
	goto L171
L175:
	;
	v1202 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v28)+314)))
	if v1202 == int32(1) {
		goto L176
	} else {
		goto L177
	}
L176:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v28)+320)) = v1020
	*(*int32)(unsafe.Add(mBase, uint32(v28)+316)) = v1021
	*(*int32)(unsafe.Add(mBase, uint32(v28)+324)) = v1029
	*(*int32)(unsafe.Add(mBase, uint32(v28)+328)) = v1023
	*(*int32)(unsafe.Add(mBase, uint32(v28)+332)) = v1024
	*(*int32)(unsafe.Add(mBase, uint32(v28)+336)) = v1022
	*(*int32)(unsafe.Add(mBase, uint32(v28)+340)) = v1025
	*(*int32)(unsafe.Add(mBase, uint32(v28)+344)) = v1026
	*(*int32)(unsafe.Add(mBase, uint32(v28)+348)) = v1027
	F_PopActiveSnapshot(m)
	mBase = m.M
	v1215 = m.ExcPending
	if v1215 != 0 {
		goto L6
	} else {
		goto L179
	}
L177:
	;
	goto L178
L178:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v28)+320)) = v1020
	*(*int32)(unsafe.Add(mBase, uint32(v28)+316)) = v1021
	*(*int32)(unsafe.Add(mBase, uint32(v28)+324)) = v1029
	*(*int32)(unsafe.Add(mBase, uint32(v28)+328)) = v1023
	*(*int32)(unsafe.Add(mBase, uint32(v28)+332)) = v1024
	*(*int32)(unsafe.Add(mBase, uint32(v28)+336)) = v1022
	*(*int32)(unsafe.Add(mBase, uint32(v28)+340)) = v1025
	*(*int32)(unsafe.Add(mBase, uint32(v28)+344)) = v1026
	*(*int32)(unsafe.Add(mBase, uint32(v28)+348)) = v1027
	F_CommandCounterIncrement(m)
	mBase = m.M
	v1248 = m.ExcPending
	if v1248 != 0 {
		goto L6
	} else {
		goto L182
	}
L179:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v28)+320)) = v1020
	*(*int32)(unsafe.Add(mBase, uint32(v28)+316)) = v1021
	*(*int32)(unsafe.Add(mBase, uint32(v28)+324)) = v1029
	*(*int32)(unsafe.Add(mBase, uint32(v28)+328)) = v1023
	*(*int32)(unsafe.Add(mBase, uint32(v28)+332)) = v1024
	*(*int32)(unsafe.Add(mBase, uint32(v28)+336)) = v1022
	*(*int32)(unsafe.Add(mBase, uint32(v28)+340)) = v1025
	*(*int32)(unsafe.Add(mBase, uint32(v28)+344)) = v1026
	*(*int32)(unsafe.Add(mBase, uint32(v28)+348)) = v1027
	F_CommandCounterIncrement(m)
	mBase = m.M
	v1226 = m.ExcPending
	if v1226 != 0 {
		goto L6
	} else {
		goto L180
	}
L180:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v28)+320)) = v1020
	*(*int32)(unsafe.Add(mBase, uint32(v28)+316)) = v1021
	*(*int32)(unsafe.Add(mBase, uint32(v28)+324)) = v1029
	*(*int32)(unsafe.Add(mBase, uint32(v28)+328)) = v1023
	*(*int32)(unsafe.Add(mBase, uint32(v28)+332)) = v1024
	*(*int32)(unsafe.Add(mBase, uint32(v28)+336)) = v1022
	*(*int32)(unsafe.Add(mBase, uint32(v28)+340)) = v1025
	*(*int32)(unsafe.Add(mBase, uint32(v28)+344)) = v1026
	*(*int32)(unsafe.Add(mBase, uint32(v28)+348)) = v1027
	F_CommitTransactionCommand(m)
	mBase = m.M
	v1237 = m.ExcPending
	if v1237 != 0 {
		goto L6
	} else {
		goto L181
	}
L181:
	;
	goto L162
L182:
	;
	goto L162
L183:
	;
	goto L160
L184:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v28)+320)) = v1020
	*(*int32)(unsafe.Add(mBase, uint32(v28)+316)) = v1021
	*(*int32)(unsafe.Add(mBase, uint32(v28)+324)) = v1029
	*(*int32)(unsafe.Add(mBase, uint32(v28)+328)) = v1023
	*(*int32)(unsafe.Add(mBase, uint32(v28)+332)) = v1024
	*(*int32)(unsafe.Add(mBase, uint32(v28)+336)) = v1022
	*(*int32)(unsafe.Add(mBase, uint32(v28)+340)) = v1025
	*(*int32)(unsafe.Add(mBase, uint32(v28)+344)) = v1026
	*(*int32)(unsafe.Add(mBase, uint32(v28)+348)) = v1027
	F_StartTransactionCommand(m)
	mBase = m.M
	v1312 = m.ExcPending
	if v1312 != 0 {
		goto L6
	} else {
		goto L187
	}
L185:
	;
	goto L186
L186:
	;
	v1313 = *(*int32)(unsafe.Add(mBase, uint32(l1)))
	if v1313&int32(513) == int32(1) {
		goto L188
	} else {
		goto L189
	}
L187:
	;
	goto L186
L188:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v28)+320)) = v1020
	*(*int32)(unsafe.Add(mBase, uint32(v28)+316)) = v1021
	*(*int32)(unsafe.Add(mBase, uint32(v28)+324)) = v1029
	*(*int32)(unsafe.Add(mBase, uint32(v28)+328)) = v1023
	*(*int32)(unsafe.Add(mBase, uint32(v28)+332)) = v1024
	*(*int32)(unsafe.Add(mBase, uint32(v28)+336)) = v1022
	*(*int32)(unsafe.Add(mBase, uint32(v28)+340)) = v1025
	*(*int32)(unsafe.Add(mBase, uint32(v28)+344)) = v1026
	*(*int32)(unsafe.Add(mBase, uint32(v28)+348)) = v1027
	F_vac_update_datfrozenxid(m)
	mBase = m.M
	v1328 = m.ExcPending
	if v1328 != 0 {
		goto L6
	} else {
		goto L191
	}
L189:
	;
	goto L190
L190:
	;
	m.G0 = v28 + int32(352)
	return
L191:
	;
	goto L190
L192:
	;
	goto L5
L193:
	;
	v1389 = int32(v1385)
	m.G0 = v28
	v1391 = *(*int32)(unsafe.Add(mBase, uint32(v1389)+4))
	v1392 = *(*int32)(unsafe.Add(mBase, uint32(v1389)))
	v1395 = *(*int32)(unsafe.Add(mBase, uint32(v1392)))
	if v28+int32(140) == v1395 {
		goto L196
	} else {
		goto L197
	}
L194:
	;
	m.ExcPending = 1
	goto L202
L195:
	;
	if v1399 != 0 {
		goto L199
	} else {
		goto L200
	}
L196:
	;
	v1397 = *(*int32)(unsafe.Add(mBase, uint32(v1392)+4))
	v1399 = v1397
	goto L198
L197:
	;
	v1399 = int32(0)
	goto L198
L198:
	;
	goto L195
L199:
	;
	v1400 = *(*int32)(unsafe.Add(mBase, uint32(v28)+348))
	v1401 = *(*int32)(unsafe.Add(mBase, uint32(v28)+344))
	v1402 = *(*int32)(unsafe.Add(mBase, uint32(v28)+340))
	v1403 = *(*int32)(unsafe.Add(mBase, uint32(v28)+336))
	v1404 = *(*int32)(unsafe.Add(mBase, uint32(v28)+332))
	v1405 = *(*int32)(unsafe.Add(mBase, uint32(v28)+328))
	v1406 = *(*int32)(unsafe.Add(mBase, uint32(v28)+324))
	v1407 = *(*int32)(unsafe.Add(mBase, uint32(v28)+320))
	v1408 = *(*int32)(unsafe.Add(mBase, uint32(v28)+316))
	v37 = v1407
	v38 = v1408
	v39 = v1403
	v40 = v1405
	v41 = v1404
	v42 = v1402
	v43 = v1401
	v44 = v1400
	v45 = v1406
	v46 = v1399
	v49 = v1391
	goto L1
L200:
	;
	goto L201
L201:
	;
	F___wasm_longjmp(m, v1392, v1391)
	mBase = m.M
	v1410 = m.ExcPending
	if v1410 != 0 {
		goto L202
	} else {
		goto L203
	}
L202:
	;
	return
L203:
	;
	base.Wasm_trap_unreachable()
	for {
	}
}
func F_vacuum_get_cutoffs(m *base.Module, l0 int32, l1 int32, l2 int32) int32 {
	mBase := m.M
	_ = mBase
	var v15 int32
	_ = v15
	var v16 int32
	_ = v16
	var v17 int32
	_ = v17
	var v18 int32
	_ = v18
	var v19 int32
	_ = v19
	var v20 int32
	_ = v20
	var v22 int32
	_ = v22
	var v23 int32
	_ = v23
	var v25 int32
	_ = v25
	var v28 int32
	_ = v28
	var v30 int32
	_ = v30
	var v31 int32
	_ = v31
	var v33 int64
	_ = v33
	var v34 int32
	_ = v34
	var v35 int32
	_ = v35
	var v36 int32
	_ = v36
	var v38 int32
	_ = v38
	var v39 int32
	_ = v39
	var v42 int32
	_ = v42
	var v43 int32
	_ = v43
	var v44 int32
	_ = v44
	var v47 int32
	_ = v47
	var v49 int32
	_ = v49
	var v50 int32
	_ = v50
	var v53 int32
	_ = v53
	var v60 int32
	_ = v60
	var v61 int32
	_ = v61
	var v67 int32
	_ = v67
	var v71 int32
	_ = v71
	var v76 int32
	_ = v76
	var v78 int32
	_ = v78
	var v84 int32
	_ = v84
	var v85 int32
	_ = v85
	var v91 int32
	_ = v91
	var v95 int32
	_ = v95
	var v100 int32
	_ = v100
	var v101 int32
	_ = v101
	var v104 int32
	_ = v104
	var v107 int32
	_ = v107
	var v109 int32
	_ = v109
	var v111 int32
	_ = v111
	var v113 int32
	_ = v113
	var v114 int32
	_ = v114
	var v117 int32
	_ = v117
	var v121 int32
	_ = v121
	var v124 int32
	_ = v124
	var v126 int32
	_ = v126
	var v127 int32
	_ = v127
	var v130 int32
	_ = v130
	var v133 int32
	_ = v133
	var v135 int32
	_ = v135
	var v137 int32
	_ = v137
	var v140 int32
	_ = v140
	var v144 int32
	_ = v144
	var v146 int32
	_ = v146
	var v151 int32
	_ = v151
	var v154 int32
	_ = v154
	var v155 float64
	_ = v155
	var v158 float64
	_ = v158
	var v160 float64
	_ = v160
	var v162 int32
	_ = v162
	var v165 int32
	_ = v165
	var v169 int32
	_ = v169
	var v172 int32
	_ = v172
	var v175 int32
	_ = v175
	var v176 float64
	_ = v176
	var v179 float64
	_ = v179
	var v181 float64
	_ = v181
	var v182 int32
	_ = v182
	var v185 int32
	_ = v185
	var v189 int32
	_ = v189
	v15 = *(*int32)(unsafe.Add(mBase, uint32(l1)+16))
	v16 = *(*int32)(unsafe.Add(mBase, uint32(l1)+8))
	v17 = *(*int32)(unsafe.Add(mBase, uint32(l1)+12))
	v18 = *(*int32)(unsafe.Add(mBase, uint32(l1)+4))
	v19 = *(*int32)(unsafe.Add(mBase, uint32(l0)+48))
	v20 = *(*int32)(unsafe.Add(mBase, uint32(v19)+136))
	*(*int32)(unsafe.Add(mBase, uint32(l2))) = v20
	v22 = *(*int32)(unsafe.Add(mBase, uint32(l0)+48))
	v23 = *(*int32)(unsafe.Add(mBase, uint32(v22)+140))
	*(*int32)(unsafe.Add(mBase, uint32(l2)+4)) = v23
	v25 = F_GetOldestNonRemovableTransactionId(m, l0)
	mBase = m.M
	v28 = m.ExcPending
	if v28 != 0 {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	return int32(0)
L2:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l2)+8)) = v25
	v30 = F_GetOldestMultiXactId(m)
	mBase = m.M
	v31 = m.ExcPending
	if v31 != 0 {
		goto L1
	} else {
		goto L3
	}
L3:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l2)+12)) = v30
	v33 = F_ReadNextFullTransactionId(m)
	mBase = m.M
	v34 = m.ExcPending
	if v34 != 0 {
		goto L1
	} else {
		goto L4
	}
L4:
	;
	v35 = F_ReadNextMultiXactId(m)
	mBase = m.M
	v36 = m.ExcPending
	if v36 != 0 {
		goto L1
	} else {
		goto L5
	}
L5:
	;
	v38 = F_MultiXactMemberFreezeThreshold(m)
	mBase = m.M
	v39 = m.ExcPending
	if v39 != 0 {
		goto L1
	} else {
		goto L6
	}
L6:
	;
	if v35 == v38 {
		goto L7
	} else {
		goto L8
	}
L7:
	;
	v42 = int32(-1)
	goto L9
L8:
	;
	v42 = v38 - v35
	goto L9
L9:
	;
	v43 = base.I32_wrap_i64(v33)
	v44 = *(*int32)(unsafe.Add(mBase, uint32(l2)+8))
	if base.Ui32(int32(3)) <= base.Ui32(v44) {
		goto L11
	} else {
		goto L12
	}
L10:
	;
	v78 = *(*int32)(unsafe.Add(mBase, uint32(l2)+12))
	if int32(0) <= v78+v42 {
		goto L23
	} else {
		goto L24
	}
L11:
	;
	v47 = int32(3)
	v49 = *(*int32)(unsafe.Add(mBase, _c_F_vacuum_get_cutoffs[0]))
	v50 = v43 - v49
	if base.Ui32(v50) <= base.Ui32(v47) {
		goto L14
	} else {
		goto L15
	}
L12:
	;
	goto L13
L13:
	;
	v60 = F_errstart(m, int32(19), int32(0))
	mBase = m.M
	v61 = m.ExcPending
	if v61 != 0 {
		goto L1
	} else {
		goto L18
	}
L14:
	;
	v53 = v47
	goto L16
L15:
	;
	v53 = v50
	goto L16
L16:
	;
	if int32(0) <= v44-v53 {
		goto L10
	} else {
		goto L17
	}
L17:
	;
	goto L13
L18:
	;
	if v60 == int32(0) {
		goto L10
	} else {
		goto L19
	}
L19:
	;
	F_errmsg(m, int32(_a_F_vacuum_get_cutoffs_4), int32(0))
	mBase = m.M
	v67 = m.ExcPending
	if v67 != 0 {
		goto L1
	} else {
		goto L20
	}
L20:
	;
	F_errhint(m, int32(_a_F_vacuum_get_cutoffs_5), int32(0))
	mBase = m.M
	v71 = m.ExcPending
	if v71 != 0 {
		goto L1
	} else {
		goto L21
	}
L21:
	;
	F_errfinish(m, int32(_a_F_vacuum_get_cutoffs_2), int32(1175), int32(_a_F_vacuum_get_cutoffs_3))
	mBase = m.M
	v76 = m.ExcPending
	if v76 != 0 {
		goto L1
	} else {
		goto L22
	}
L22:
	;
	goto L10
L23:
	;
	v101 = *(*int32)(unsafe.Add(mBase, uint32(l2)+8))
	v104 = *(*int32)(unsafe.Add(mBase, _c_F_vacuum_get_cutoffs[1]))
	if v18 < int32(0) {
		goto L30
	} else {
		goto L31
	}
L24:
	;
	v84 = F_errstart(m, int32(19), int32(0))
	mBase = m.M
	v85 = m.ExcPending
	if v85 != 0 {
		goto L1
	} else {
		goto L25
	}
L25:
	;
	if v84 == int32(0) {
		goto L23
	} else {
		goto L26
	}
L26:
	;
	F_errmsg(m, int32(_a_F_vacuum_get_cutoffs_0), int32(0))
	mBase = m.M
	v91 = m.ExcPending
	if v91 != 0 {
		goto L1
	} else {
		goto L27
	}
L27:
	;
	F_errhint(m, int32(_a_F_vacuum_get_cutoffs_1), int32(0))
	mBase = m.M
	v95 = m.ExcPending
	if v95 != 0 {
		goto L1
	} else {
		goto L28
	}
L28:
	;
	F_errfinish(m, int32(_a_F_vacuum_get_cutoffs_2), int32(1180), int32(_a_F_vacuum_get_cutoffs_3))
	mBase = m.M
	v100 = m.ExcPending
	if v100 != 0 {
		goto L1
	} else {
		goto L29
	}
L29:
	;
	goto L23
L30:
	;
	v107 = v104
	goto L32
L31:
	;
	v107 = v18
	goto L32
L32:
	;
	v109 = *(*int32)(unsafe.Add(mBase, _c_F_vacuum_get_cutoffs[0]))
	v111 = base.I32_div_s(v109, int32(2))
	if v107 < v111 {
		goto L33
	} else {
		goto L34
	}
L33:
	;
	v113 = v107
	goto L35
L34:
	;
	v113 = v111
	goto L35
L35:
	;
	v114 = v43 - v113
	if base.Ui32(v114) <= base.Ui32(int32(3)) {
		goto L36
	} else {
		goto L37
	}
L36:
	;
	v117 = int32(3)
	goto L38
L37:
	;
	v117 = v114
	goto L38
L38:
	;
	if v101-v117 < int32(0) {
		goto L39
	} else {
		goto L40
	}
L39:
	;
	v121 = v101
	goto L41
L40:
	;
	v121 = v117
	goto L41
L41:
	;
	if base.Ui32(v101) < base.Ui32(int32(3)) {
		goto L42
	} else {
		goto L43
	}
L42:
	;
	v124 = v101
	goto L44
L43:
	;
	v124 = v121
	goto L44
L44:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l2)+16)) = v124
	v126 = int32(1)
	v127 = *(*int32)(unsafe.Add(mBase, uint32(l2)+12))
	v130 = *(*int32)(unsafe.Add(mBase, _c_F_vacuum_get_cutoffs[2]))
	if v17 < int32(0) {
		goto L45
	} else {
		goto L46
	}
L45:
	;
	v133 = v130
	goto L47
L46:
	;
	v133 = v17
	goto L47
L47:
	;
	v135 = base.I32_div_s(v38, int32(2))
	if v133 < v135 {
		goto L48
	} else {
		goto L49
	}
L48:
	;
	v137 = v133
	goto L50
L49:
	;
	v137 = v135
	goto L50
L50:
	;
	if v35 == v137 {
		goto L51
	} else {
		goto L52
	}
L51:
	;
	v140 = v126
	goto L53
L52:
	;
	v140 = v35 - v137
	goto L53
L53:
	;
	if v127-v140 < int32(0) {
		goto L54
	} else {
		goto L55
	}
L54:
	;
	v144 = v127
	goto L56
L55:
	;
	v144 = v140
	goto L56
L56:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l2)+20)) = v144
	v146 = *(*int32)(unsafe.Add(mBase, uint32(l2)))
	if base.Ui32(v146) < base.Ui32(int32(3)) {
		v189 = v126
		goto L57
	} else {
		goto L58
	}
L57:
	;
	return v189
L58:
	;
	v151 = *(*int32)(unsafe.Add(mBase, _c_F_vacuum_get_cutoffs[3]))
	if v16 < int32(0) {
		goto L59
	} else {
		goto L60
	}
L59:
	;
	v154 = v151
	goto L61
L60:
	;
	v154 = v16
	goto L61
L61:
	;
	v155 = base.F64_convert_i32_s(v154)
	v158 = base.F64_mul(base.F64_convert_i32_s(v109), float64(0.95))
	if base.F64_lt(v155, v158) != 0 {
		goto L62
	} else {
		goto L63
	}
L62:
	;
	v160 = v155
	goto L64
L63:
	;
	v160 = v158
	goto L64
L64:
	;
	v162 = v43 - base.I32_trunc_sat_f64_s(v160)
	if base.Ui32(v162) <= base.Ui32(int32(3)) {
		goto L65
	} else {
		goto L66
	}
L65:
	;
	v165 = int32(3)
	goto L67
L66:
	;
	v165 = v162
	goto L67
L67:
	;
	if v146-v165 <= int32(0) {
		v189 = v126
		goto L57
	} else {
		goto L68
	}
L68:
	;
	v169 = *(*int32)(unsafe.Add(mBase, uint32(l2)+4))
	v172 = *(*int32)(unsafe.Add(mBase, _c_F_vacuum_get_cutoffs[4]))
	if v15 < int32(0) {
		goto L69
	} else {
		goto L70
	}
L69:
	;
	v175 = v172
	goto L71
L70:
	;
	v175 = v15
	goto L71
L71:
	;
	v176 = base.F64_convert_i32_s(v175)
	v179 = base.F64_mul(base.F64_convert_i32_s(v38), float64(0.95))
	if base.F64_lt(v176, v179) != 0 {
		goto L72
	} else {
		goto L73
	}
L72:
	;
	v181 = v176
	goto L74
L73:
	;
	v181 = v179
	goto L74
L74:
	;
	v182 = base.I32_trunc_sat_f64_s(v181)
	if v182 == v35 {
		goto L75
	} else {
		goto L76
	}
L75:
	;
	v185 = int32(-1)
	goto L77
L76:
	;
	v185 = v182 - v35
	goto L77
L77:
	;
	v189 = base.B2i32(v169+v185 <= int32(0))
	goto L57
}
func F_vacuum_rel(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32) int32 {
	mBase := m.M
	_ = mBase
	var v14 int32
	_ = v14
	var v16 int32
	_ = v16
	var v18 int64
	_ = v18
	var v20 int64
	_ = v20
	var v22 int64
	_ = v22
	var v24 int64
	_ = v24
	var v26 int64
	_ = v26
	var v28 int64
	_ = v28
	var v30 int64
	_ = v30
	var v32 int64
	_ = v32
	var v37 int32
	_ = v37
	var v38 int32
	_ = v38
	var v40 int32
	_ = v40
	var v44 int32
	_ = v44
	var v48 int32
	_ = v48
	var v49 int32
	_ = v49
	var v51 int32
	_ = v51
	var v52 int32
	_ = v52
	var v54 int32
	_ = v54
	var v56 int32
	_ = v56
	var v60 int32
	_ = v60
	var v62 int32
	_ = v62
	var v64 int32
	_ = v64
	var v65 int32
	_ = v65
	var v66 int32
	_ = v66
	var v70 int32
	_ = v70
	var v74 int32
	_ = v74
	var v78 int32
	_ = v78
	var v79 int32
	_ = v79
	var v81 int32
	_ = v81
	var v83 int32
	_ = v83
	var v85 int32
	_ = v85
	var v86 int32
	_ = v86
	var v93 int32
	_ = v93
	var v94 int32
	_ = v94
	var v95 int32
	_ = v95
	var v96 int32
	_ = v96
	var v97 int32
	_ = v97
	var v98 int32
	_ = v98
	var v99 int32
	_ = v99
	var v102 int32
	_ = v102
	var v103 int32
	_ = v103
	var v106 int32
	_ = v106
	var v107 int32
	_ = v107
	var v109 int32
	_ = v109
	var v117 int32
	_ = v117
	var v120 int32
	_ = v120
	var v123 int32
	_ = v123
	var v124 int32
	_ = v124
	var v125 int32
	_ = v125
	var v131 int32
	_ = v131
	var v136 int32
	_ = v136
	var v138 int32
	_ = v138
	var v140 int32
	_ = v140
	var v142 int32
	_ = v142
	var v143 int32
	_ = v143
	var v146 int32
	_ = v146
	var v152 int32
	_ = v152
	var v154 int32
	_ = v154
	var v156 int32
	_ = v156
	var v157 int64
	_ = v157
	var v162 int32
	_ = v162
	var v163 int32
	_ = v163
	var v164 int32
	_ = v164
	var v170 int32
	_ = v170
	var v172 int32
	_ = v172
	var v181 int32
	_ = v181
	var v182 float64
	_ = v182
	var v188 int32
	_ = v188
	var v189 int32
	_ = v189
	var v191 int32
	_ = v191
	var v195 int32
	_ = v195
	var v203 int32
	_ = v203
	var v204 int32
	_ = v204
	var v205 int32
	_ = v205
	var v208 int32
	_ = v208
	var v213 int32
	_ = v213
	var v220 int32
	_ = v220
	var v221 int32
	_ = v221
	var v222 int32
	_ = v222
	var v228 int32
	_ = v228
	var v231 int32
	_ = v231
	var v233 int32
	_ = v233
	var v234 int32
	_ = v234
	var v235 int32
	_ = v235
	var v243 int32
	_ = v243
	var v245 int32
	_ = v245
	var v247 int32
	_ = v247
	var v250 int32
	_ = v250
	var v258 int32
	_ = v258
	var v264 int32
	_ = v264
	var v265 int32
	_ = v265
	var v266 int32
	_ = v266
	var v268 int32
	_ = v268
	var v269 int32
	_ = v269
	var v272 int32
	_ = v272
	var v273 int32
	_ = v273
	var v274 int32
	_ = v274
	var v281 int32
	_ = v281
	var v283 int32
	_ = v283
	var v285 int32
	_ = v285
	var v287 int32
	_ = v287
	var v291 int64
	_ = v291
	var v293 int64
	_ = v293
	var v295 int64
	_ = v295
	var v297 int64
	_ = v297
	var v299 int64
	_ = v299
	var v301 int64
	_ = v301
	var v303 int64
	_ = v303
	var v305 int64
	_ = v305
	var v310 int32
	_ = v310
	var v311 int32
	_ = v311
	var v315 int32
	_ = v315
	var v325 int32
	_ = v325
	var v330 int32
	_ = v330
	var v332 int32
	_ = v332
	var v334 int32
	_ = v334
	v14 = m.G0
	v16 = v14 - int32(160)
	m.G0 = v16
	v18 = *(*int64)(unsafe.Add(mBase, uint32(l2)+56))
	*(*int64)(unsafe.Add(mBase, uint32(v16)+136)) = v18
	v20 = *(*int64)(unsafe.Add(mBase, uint32(l2)+48))
	*(*int64)(unsafe.Add(mBase, uint32(v16)+128)) = v20
	v22 = *(*int64)(unsafe.Add(mBase, uint32(l2)+40))
	*(*int64)(unsafe.Add(mBase, uint32(v16)+120)) = v22
	v24 = *(*int64)(unsafe.Add(mBase, uint32(l2)+32))
	*(*int64)(unsafe.Add(mBase, uint32(v16)+112)) = v24
	v26 = *(*int64)(unsafe.Add(mBase, uint32(l2)+24))
	*(*int64)(unsafe.Add(mBase, uint32(v16)+104)) = v26
	v28 = *(*int64)(unsafe.Add(mBase, uint32(l2)+16))
	*(*int64)(unsafe.Add(mBase, uint32(v16)+96)) = v28
	v30 = *(*int64)(unsafe.Add(mBase, uint32(l2)+8))
	*(*int64)(unsafe.Add(mBase, uint32(v16)+88)) = v30
	v32 = *(*int64)(unsafe.Add(mBase, uint32(l2)))
	*(*int64)(unsafe.Add(mBase, uint32(v16)+80)) = v32
	F_StartTransactionCommand(m)
	mBase = m.M
	v37 = m.ExcPending
	if v37 != 0 {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	return int32(0)
L2:
	;
	v38 = *(*int32)(unsafe.Add(mBase, uint32(l2)))
	v40 = v38 & int32(16)
	if v40 == int32(0) {
		goto L3
	} else {
		goto L4
	}
L3:
	;
	v44 = *(*int32)(unsafe.Add(mBase, _c_F_vacuum_rel[0]))
	v48 = F_LWLockAcquire(m, v44+int32(512), int32(0))
	mBase = m.M
	v49 = m.ExcPending
	if v49 != 0 {
		goto L1
	} else {
		goto L6
	}
L4:
	;
	goto L5
L5:
	;
	v78 = F_GetTransactionSnapshot(m)
	mBase = m.M
	v79 = m.ExcPending
	if v79 != 0 {
		goto L1
	} else {
		goto L11
	}
L6:
	;
	v51 = *(*int32)(unsafe.Add(mBase, _c_F_vacuum_rel[1]))
	v52 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v51)+36)))
	v54 = v52 | int32(2)
	*(*uint8)(unsafe.Add(mBase, uint32(v51)+36)) = uint8(v54)
	v56 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l2)+20)))
	if v56 == int32(1) {
		goto L7
	} else {
		goto L8
	}
L7:
	;
	v60 = v52 | int32(10)
	*(*uint8)(unsafe.Add(mBase, uint32(v51)+36)) = uint8(v60)
	v62 = v60
	goto L9
L8:
	;
	v62 = v54
	goto L9
L9:
	;
	v64 = *(*int32)(unsafe.Add(mBase, _c_F_vacuum_rel[2]))
	v65 = *(*int32)(unsafe.Add(mBase, uint32(v64)+12))
	v66 = *(*int32)(unsafe.Add(mBase, uint32(v51)+32))
	*(*uint8)(unsafe.Add(mBase, uint32(v65+v66))) = uint8(v62)
	v70 = *(*int32)(unsafe.Add(mBase, _c_F_vacuum_rel[0]))
	F_LWLockRelease(m, v70+int32(512))
	mBase = m.M
	v74 = m.ExcPending
	if v74 != 0 {
		goto L1
	} else {
		goto L10
	}
L10:
	;
	goto L5
L11:
	;
	F_PushActiveSnapshot(m, v78)
	mBase = m.M
	v81 = m.ExcPending
	if v81 != 0 {
		goto L1
	} else {
		goto L12
	}
L12:
	;
	v83 = *(*int32)(unsafe.Add(mBase, _c_F_vacuum_rel[3]))
	if v83 != 0 {
		goto L13
	} else {
		goto L14
	}
L13:
	;
	F_ProcessInterrupts(m)
	mBase = m.M
	v85 = m.ExcPending
	if v85 != 0 {
		goto L1
	} else {
		goto L16
	}
L14:
	;
	goto L15
L15:
	;
	v86 = *(*int32)(unsafe.Add(mBase, uint32(l2)+24))
	if v40 != 0 {
		goto L18
	} else {
		goto L19
	}
L16:
	;
	goto L15
L17:
	;
	m.G0 = v16 + int32(160)
	return v334
L18:
	;
	v93 = int32(8)
	goto L20
L19:
	;
	v93 = int32(4)
	goto L20
L20:
	;
	v94 = F_vacuum_open_relation(m, l0, l1, v38, int32(base.Ui32(v86^int32(-1))>>(uint(int32(31))%32)), v93)
	mBase = m.M
	v95 = m.ExcPending
	if v95 != 0 {
		goto L1
	} else {
		goto L21
	}
L21:
	;
	if v94 != 0 {
		goto L22
	} else {
		goto L23
	}
L22:
	;
	v96 = *(*int32)(unsafe.Add(mBase, uint32(l2)+40))
	if v96 != 0 {
		goto L26
	} else {
		goto L27
	}
L23:
	;
	goto L24
L24:
	;
	F_PopActiveSnapshot(m)
	mBase = m.M
	v330 = m.ExcPending
	if v330 != 0 {
		goto L1
	} else {
		goto L109
	}
L25:
	;
	F_relation_close(m, v94, v93)
	mBase = m.M
	v325 = m.ExcPending
	if v325 != 0 {
		goto L1
	} else {
		goto L108
	}
L26:
	;
	v98 = v96
	goto L28
L27:
	;
	v97 = *(*int32)(unsafe.Add(mBase, uint32(v94)+56))
	v98 = v97
	goto L28
L28:
	;
	v99 = *(*int32)(unsafe.Add(mBase, uint32(v94)+48))
	v102 = F_vacuum_is_permitted_for_relation(m, v98, v99, v38&int32(-3))
	mBase = m.M
	v103 = m.ExcPending
	if v103 != 0 {
		goto L1
	} else {
		goto L29
	}
L29:
	;
	if v102 == int32(0) {
		goto L25
	} else {
		goto L30
	}
L30:
	;
	v106 = *(*int32)(unsafe.Add(mBase, uint32(v94)+48))
	v107 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v106)+119)))
	v109 = v107 - int32(109)
	if int32(1)<<(uint(v109)%32)&int32(169) != 0 {
		goto L31
	} else {
		goto L32
	}
L31:
	;
	v117 = base.B2i32(base.Ui32(v109) <= base.Ui32(int32(7)))
	goto L33
L32:
	;
	v117 = int32(0)
	goto L33
L33:
	;
	if v117 == int32(0) {
		goto L34
	} else {
		goto L35
	}
L34:
	;
	v120 = int32(0)
	v123 = F_errstart(m, int32(19), v120)
	mBase = m.M
	v124 = m.ExcPending
	if v124 != 0 {
		goto L1
	} else {
		goto L37
	}
L35:
	;
	goto L36
L36:
	;
	v143 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v106)+118)))
	if v143 == int32(116) {
		goto L46
	} else {
		goto L47
	}
L37:
	;
	if v123 != 0 {
		goto L38
	} else {
		goto L39
	}
L38:
	;
	v125 = *(*int32)(unsafe.Add(mBase, uint32(v94)+48))
	*(*int32)(unsafe.Add(mBase, uint32(v16))) = v125 + int32(4)
	F_errmsg(m, int32(_a_F_vacuum_rel_0), v16)
	mBase = m.M
	v131 = m.ExcPending
	if v131 != 0 {
		goto L1
	} else {
		goto L41
	}
L39:
	;
	goto L40
L40:
	;
	F_relation_close(m, v94, v93)
	mBase = m.M
	v138 = m.ExcPending
	if v138 != 0 {
		goto L1
	} else {
		goto L43
	}
L41:
	;
	F_errfinish(m, int32(_a_F_vacuum_rel_1), int32(2136), int32(_a_F_vacuum_rel_2))
	mBase = m.M
	v136 = m.ExcPending
	if v136 != 0 {
		goto L1
	} else {
		goto L42
	}
L42:
	;
	goto L40
L43:
	;
	F_PopActiveSnapshot(m)
	mBase = m.M
	v140 = m.ExcPending
	if v140 != 0 {
		goto L1
	} else {
		goto L44
	}
L44:
	;
	F_CommitTransactionCommand(m)
	mBase = m.M
	v142 = m.ExcPending
	if v142 != 0 {
		goto L1
	} else {
		goto L45
	}
L45:
	;
	v334 = v120
	goto L17
L46:
	;
	v146 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v94)+24)))
	if v146 == int32(0) {
		goto L25
	} else {
		goto L49
	}
L47:
	;
	goto L48
L48:
	;
	if v107 == int32(112) {
		goto L51
	} else {
		goto L52
	}
L49:
	;
	goto L48
L50:
	;
	v334 = int32(1)
	goto L17
L51:
	;
	F_relation_close(m, v94, v93)
	mBase = m.M
	v152 = m.ExcPending
	if v152 != 0 {
		goto L1
	} else {
		goto L54
	}
L52:
	;
	goto L53
L53:
	;
	v157 = *(*int64)(unsafe.Add(mBase, uint32(v94)+60))
	*(*int64)(unsafe.Add(mBase, uint32(v16)+152)) = v157
	F_LockRelationIdForSession(m, v16+int32(152), v93)
	mBase = m.M
	v162 = m.ExcPending
	if v162 != 0 {
		goto L1
	} else {
		goto L57
	}
L54:
	;
	F_PopActiveSnapshot(m)
	mBase = m.M
	v154 = m.ExcPending
	if v154 != 0 {
		goto L1
	} else {
		goto L55
	}
L55:
	;
	F_CommitTransactionCommand(m)
	mBase = m.M
	v156 = m.ExcPending
	if v156 != 0 {
		goto L1
	} else {
		goto L56
	}
L56:
	;
	goto L50
L57:
	;
	v163 = *(*int32)(unsafe.Add(mBase, uint32(v94)+180))
	v164 = *(*int32)(unsafe.Add(mBase, uint32(l2)+32))
	if v164 == int32(0) {
		goto L60
	} else {
		goto L61
	}
L58:
	;
	v191 = *(*int32)(unsafe.Add(mBase, uint32(l2)+36))
	if v191 == int32(0) {
		goto L69
	} else {
		goto L70
	}
L59:
	;
	v181 = int32(0)
	v182 = *(*float64)(unsafe.Add(mBase, uint32(v163)+128))
	if base.F64_ge(v182, float64(0)) == v181 {
		v188 = v163
		v189 = v181
		goto L58
	} else {
		goto L68
	}
L60:
	;
	if v163 == int32(0) {
		goto L63
	} else {
		goto L64
	}
L61:
	;
	goto L62
L62:
	;
	if v163 != 0 {
		goto L59
	} else {
		goto L67
	}
L63:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l2)+32)) = int32(1)
	goto L62
L64:
	;
	v170 = *(*int32)(unsafe.Add(mBase, uint32(v163)+120))
	switch v170 {
	case 0:
		goto L63
	default:
		goto L66
	case 2:
		v172 = int32(3)
		goto L65
	}
L65:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l2)+32)) = v172
	goto L59
L66:
	;
	v172 = int32(2)
	goto L65
L67:
	;
	v188 = int32(0)
	v189 = int32(1)
	goto L58
L68:
	;
	*(*float64)(unsafe.Add(mBase, uint32(l2)+48)) = v182
	v188 = v163
	v189 = v181
	goto L58
L69:
	;
	if v189 != 0 {
		goto L73
	} else {
		goto L74
	}
L70:
	;
	goto L71
L71:
	;
	v208 = int32(0)
	v213 = int32(80)
	if base.B2i32(v38&int32(128) == v208)|base.B2i32(v38&v213 == v213) == v208 {
		goto L79
	} else {
		goto L80
	}
L72:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l2)+36)) = v205
	goto L71
L73:
	;
	v203 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_vacuum_rel[4])))
	if v203 != 0 {
		goto L76
	} else {
		goto L77
	}
L74:
	;
	v195 = *(*int32)(unsafe.Add(mBase, uint32(v188)+124))
	switch v195 + int32(1) {
	case 0:
		goto L73
	default:
		goto L75
	case 2:
		v205 = int32(3)
		goto L72
	}
L75:
	;
	v205 = int32(2)
	goto L72
L76:
	;
	v204 = int32(3)
	goto L78
L77:
	;
	v204 = int32(2)
	goto L78
L78:
	;
	v205 = v204
	goto L72
L79:
	;
	v220 = *(*int32)(unsafe.Add(mBase, uint32(v94)+48))
	v221 = *(*int32)(unsafe.Add(mBase, uint32(v220)+112))
	v222 = v221
	goto L81
L80:
	;
	v222 = v208
	goto L81
L81:
	;
	v228 = *(*int32)(unsafe.Add(mBase, _c_F_vacuum_rel[5]))
	*(*int32)(unsafe.Add(mBase, uint32(v16+int32(148)))) = v228
	v231 = *(*int32)(unsafe.Add(mBase, _c_F_vacuum_rel[6]))
	*(*int32)(unsafe.Add(mBase, uint32(v16+int32(144)))) = v231
	goto L82
L82:
	;
	v233 = *(*int32)(unsafe.Add(mBase, uint32(v94)+48))
	v234 = *(*int32)(unsafe.Add(mBase, uint32(v233)+80))
	v235 = *(*int32)(unsafe.Add(mBase, uint32(v16)+144))
	*(*int32)(unsafe.Add(mBase, _c_F_vacuum_rel[6])) = v235 | int32(2)
	*(*int32)(unsafe.Add(mBase, _c_F_vacuum_rel[5])) = v234
	goto L83
L83:
	;
	v243 = int32(_a_F_vacuum_rel_3)
	v245 = *(*int32)(unsafe.Add(mBase, _c_F_vacuum_rel[7]))
	v247 = v245 + int32(1)
	*(*int32)(unsafe.Add(mBase, _c_F_vacuum_rel[7])) = v247
	goto L84
L84:
	;
	F_RestrictSearchPath(m)
	mBase = m.M
	v250 = m.ExcPending
	if v250 != 0 {
		goto L1
	} else {
		goto L85
	}
L85:
	;
	if v38&int32(64) != 0 {
		goto L87
	} else {
		goto L88
	}
L86:
	;
	F_AtEOXact_GUC(m, int32(0), v247)
	mBase = m.M
	v272 = m.ExcPending
	if v272 != 0 {
		goto L1
	} else {
		goto L95
	}
L87:
	;
	if v40 != 0 {
		goto L90
	} else {
		goto L91
	}
L88:
	;
	goto L89
L89:
	;
	v269 = v94
	goto L86
L90:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v16)+76)) = int32(base.Ui32(v38)>>(uint(int32(2))%32)) & int32(1)
	v258 = int32(0)
	F_cluster_rel(m, int32(3), v94, v258, v16+int32(76))
	mBase = m.M
	v264 = m.ExcPending
	if v264 != 0 {
		goto L1
	} else {
		goto L93
	}
L91:
	;
	goto L92
L92:
	;
	v265 = *(*int32)(unsafe.Add(mBase, uint32(v94)+188))
	v266 = *(*int32)(unsafe.Add(mBase, uint32(v265)+128))
	m.T0[v266].(func(*base.Module, int32, int32, int32))(m, v94, l2, l3)
	mBase = m.M
	v268 = m.ExcPending
	if v268 != 0 {
		goto L1
	} else {
		goto L94
	}
L93:
	;
	v269 = v258
	goto L86
L94:
	;
	goto L89
L95:
	;
	v273 = *(*int32)(unsafe.Add(mBase, uint32(v16)+148))
	v274 = *(*int32)(unsafe.Add(mBase, uint32(v16)+144))
	*(*int32)(unsafe.Add(mBase, _c_F_vacuum_rel[6])) = v274
	*(*int32)(unsafe.Add(mBase, _c_F_vacuum_rel[5])) = v273
	goto L96
L96:
	;
	if v269 != 0 {
		goto L97
	} else {
		goto L98
	}
L97:
	;
	F_relation_close(m, v269, int32(0))
	mBase = m.M
	v281 = m.ExcPending
	if v281 != 0 {
		goto L1
	} else {
		goto L100
	}
L98:
	;
	goto L99
L99:
	;
	F_PopActiveSnapshot(m)
	mBase = m.M
	v283 = m.ExcPending
	if v283 != 0 {
		goto L1
	} else {
		goto L101
	}
L100:
	;
	goto L99
L101:
	;
	F_CommitTransactionCommand(m)
	mBase = m.M
	v285 = m.ExcPending
	if v285 != 0 {
		goto L1
	} else {
		goto L102
	}
L102:
	;
	if v222 != 0 {
		goto L103
	} else {
		goto L104
	}
L103:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v16)+120)) = l0
	v287 = *(*int32)(unsafe.Add(mBase, uint32(v16)+80))
	*(*int32)(unsafe.Add(mBase, uint32(v16)+80)) = v287 | int32(64)
	v291 = *(*int64)(unsafe.Add(mBase, uint32(v16)+88))
	*(*int64)(unsafe.Add(mBase, uint32(v16)+16)) = v291
	v293 = *(*int64)(unsafe.Add(mBase, uint32(v16)+96))
	*(*int64)(unsafe.Add(mBase, uint32(v16)+24)) = v293
	v295 = *(*int64)(unsafe.Add(mBase, uint32(v16)+104))
	*(*int64)(unsafe.Add(mBase, uint32(v16)+32)) = v295
	v297 = *(*int64)(unsafe.Add(mBase, uint32(v16)+112))
	*(*int64)(unsafe.Add(mBase, uint32(v16)+40)) = v297
	v299 = *(*int64)(unsafe.Add(mBase, uint32(v16)+128))
	*(*int64)(unsafe.Add(mBase, uint32(v16)+56)) = v299
	v301 = *(*int64)(unsafe.Add(mBase, uint32(v16)+136))
	*(*int64)(unsafe.Add(mBase, uint32(v16)+64)) = v301
	v303 = *(*int64)(unsafe.Add(mBase, uint32(v16)+120))
	*(*int64)(unsafe.Add(mBase, uint32(v16)+48)) = v303
	v305 = *(*int64)(unsafe.Add(mBase, uint32(v16)+80))
	*(*int64)(unsafe.Add(mBase, uint32(v16)+8)) = v305
	v310 = F_vacuum_rel(m, v222, int32(0), v16+int32(8), l3, l4)
	mBase = m.M
	v311 = m.ExcPending
	if v311 != 0 {
		goto L1
	} else {
		goto L106
	}
L104:
	;
	goto L105
L105:
	;
	F_UnlockRelationIdForSession(m, v16+int32(152), v93)
	mBase = m.M
	v315 = m.ExcPending
	if v315 != 0 {
		goto L1
	} else {
		goto L107
	}
L106:
	;
	goto L105
L107:
	;
	goto L50
L108:
	;
	goto L24
L109:
	;
	F_CommitTransactionCommand(m)
	mBase = m.M
	v332 = m.ExcPending
	if v332 != 0 {
		goto L1
	} else {
		goto L110
	}
L110:
	;
	v334 = int32(0)
	goto L17
}
