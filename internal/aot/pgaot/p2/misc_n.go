package p2

import (
	base "github.com/shibukawa/pgmem/internal/aot/pgaot/base"
	"unsafe"
)

func F_NIAddAffix(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32, l5 int32, l6 int32) {
	mBase := m.M
	_ = mBase
	var v18 int32
	_ = v18
	var v20 int32
	_ = v20
	var v22 int32
	_ = v22
	var v23 int32
	_ = v23
	var v25 int32
	_ = v25
	var v29 int32
	_ = v29
	var v32 int32
	_ = v32
	var v33 int32
	_ = v33
	var v37 int32
	_ = v37
	var v38 int32
	_ = v38
	var v39 int32
	_ = v39
	var v41 int32
	_ = v41
	var v42 int32
	_ = v42
	var v45 int32
	_ = v45
	var v46 int32
	_ = v46
	var v49 int32
	_ = v49
	var v50 int32
	_ = v50
	var v56 int32
	_ = v56
	var v58 int32
	_ = v58
	var v61 int32
	_ = v61
	var v72 int32
	_ = v72
	var v73 int32
	_ = v73
	var v75 int32
	_ = v75
	var v84 int32
	_ = v84
	var v85 int32
	_ = v85
	var v87 int32
	_ = v87
	var v94 int32
	_ = v94
	var v95 int32
	_ = v95
	var v102 int32
	_ = v102
	var v103 int32
	_ = v103
	var v105 int32
	_ = v105
	var v112 int32
	_ = v112
	var v117 int32
	_ = v117
	var v122 int32
	_ = v122
	var v124 int32
	_ = v124
	var v125 int32
	_ = v125
	var v126 int32
	_ = v126
	var v127 int32
	_ = v127
	var v128 int32
	_ = v128
	var v148 int32
	_ = v148
	var v152 int32
	_ = v152
	var v154 int32
	_ = v154
	var v159 int32
	_ = v159
	var v163 int32
	_ = v163
	var v164 int32
	_ = v164
	var v165 int32
	_ = v165
	var v167 int32
	_ = v167
	var v176 int32
	_ = v176
	var v181 int32
	_ = v181
	var v185 int32
	_ = v185
	var v190 int32
	_ = v190
	var v192 int32
	_ = v192
	var v196 int32
	_ = v196
	var v202 int32
	_ = v202
	var v205 int32
	_ = v205
	var v211 int32
	_ = v211
	var v215 int32
	_ = v215
	var v217 int32
	_ = v217
	var v225 int32
	_ = v225
	var v229 int32
	_ = v229
	var v231 int32
	_ = v231
	var v237 int32
	_ = v237
	var v241 int32
	_ = v241
	var v242 int32
	_ = v242
	var v246 int32
	_ = v246
	var v253 int32
	_ = v253
	var v254 int32
	_ = v254
	var v255 int32
	_ = v255
	var v256 int32
	_ = v256
	var v259 int32
	_ = v259
	var v267 int32
	_ = v267
	var v268 int32
	_ = v268
	var v269 int32
	_ = v269
	var v271 int32
	_ = v271
	var v281 int32
	_ = v281
	var v284 int32
	_ = v284
	var v285 int32
	_ = v285
	var v288 int32
	_ = v288
	var v289 int32
	_ = v289
	var v298 int32
	_ = v298
	var v304 int32
	_ = v304
	var v309 int32
	_ = v309
	var v314 int32
	_ = v314
	var v321 int32
	_ = v321
	var v322 int32
	_ = v322
	var v325 int32
	_ = v325
	var v326 int32
	_ = v326
	var v327 int32
	_ = v327
	var v329 int32
	_ = v329
	var v330 int32
	_ = v330
	var v342 int32
	_ = v342
	var v348 int32
	_ = v348
	var v353 int32
	_ = v353
	var v358 int32
	_ = v358
	var v359 int32
	_ = v359
	var v360 int32
	_ = v360
	var v368 int32
	_ = v368
	var v369 int32
	_ = v369
	var v370 int32
	_ = v370
	var v372 int32
	_ = v372
	var v383 int32
	_ = v383
	var v389 int32
	_ = v389
	var v395 int32
	_ = v395
	var v400 int32
	_ = v400
	var v404 int32
	_ = v404
	var v411 int32
	_ = v411
	var v416 int32
	_ = v416
	var v417 int32
	_ = v417
	var v419 int32
	_ = v419
	var v421 int32
	_ = v421
	var v422 int32
	_ = v422
	var v423 int32
	_ = v423
	var v424 int32
	_ = v424
	var v444 int32
	_ = v444
	var v445 int32
	_ = v445
	var v449 int32
	_ = v449
	var v454 int32
	_ = v454
	var v469 int32
	_ = v469
	var v471 int32
	_ = v471
	var v495 int32
	_ = v495
	var v499 int32
	_ = v499
	var v504 int32
	_ = v504
	var v506 int32
	_ = v506
	var v514 int32
	_ = v514
	var v519 int32
	_ = v519
	var v523 int32
	_ = v523
	var v528 int32
	_ = v528
	var v530 int32
	_ = v530
	var v534 int32
	_ = v534
	var v540 int32
	_ = v540
	var v543 int32
	_ = v543
	var v549 int32
	_ = v549
	var v553 int32
	_ = v553
	var v555 int32
	_ = v555
	var v563 int32
	_ = v563
	var v566 int32
	_ = v566
	var v567 int32
	_ = v567
	var v571 int32
	_ = v571
	var v574 int32
	_ = v574
	var v575 int32
	_ = v575
	var v576 int32
	_ = v576
	var v584 int32
	_ = v584
	var v589 int32
	_ = v589
	var v593 int32
	_ = v593
	var v598 int32
	_ = v598
	var v600 int32
	_ = v600
	var v604 int32
	_ = v604
	var v610 int32
	_ = v610
	var v613 int32
	_ = v613
	var v619 int32
	_ = v619
	var v623 int32
	_ = v623
	var v625 int32
	_ = v625
	var v633 int32
	_ = v633
	var v638 int32
	_ = v638
	var v639 int32
	_ = v639
	var v640 int32
	_ = v640
	var v641 int32
	_ = v641
	var v643 int32
	_ = v643
	var v644 int32
	_ = v644
	var v648 int32
	_ = v648
	var v649 int32
	_ = v649
	var v667 int32
	_ = v667
	var v671 int32
	_ = v671
	var v674 int32
	_ = v674
	var v675 int32
	_ = v675
	var v679 int32
	_ = v679
	var v682 int32
	_ = v682
	var v691 int32
	_ = v691
	var v696 int32
	_ = v696
	var v700 int32
	_ = v700
	var v705 int32
	_ = v705
	var v707 int32
	_ = v707
	var v711 int32
	_ = v711
	var v717 int32
	_ = v717
	var v720 int32
	_ = v720
	var v726 int32
	_ = v726
	var v730 int32
	_ = v730
	var v732 int32
	_ = v732
	var v740 int32
	_ = v740
	var v742 int32
	_ = v742
	var v745 int32
	_ = v745
	var v746 int32
	_ = v746
	var v750 int32
	_ = v750
	var v751 int32
	_ = v751
	var v753 int32
	_ = v753
	var v754 int32
	_ = v754
	var v756 int32
	_ = v756
	var v757 int32
	_ = v757
	var v758 int32
	_ = v758
	var v759 int32
	_ = v759
	var v765 int32
	_ = v765
	var v772 int32
	_ = v772
	var v776 int32
	_ = v776
	var v778 int32
	_ = v778
	var v779 int32
	_ = v779
	var v783 int32
	_ = v783
	var v784 int32
	_ = v784
	var v786 int32
	_ = v786
	var v790 int32
	_ = v790
	var v792 int32
	_ = v792
	var v794 int32
	_ = v794
	var v797 int32
	_ = v797
	var v802 int32
	_ = v802
	var v803 int32
	_ = v803
	var v804 int32
	_ = v804
	var v806 int32
	_ = v806
	var v807 int32
	_ = v807
	var v808 int32
	_ = v808
	var v810 int32
	_ = v810
	var v814 int32
	_ = v814
	var v819 int32
	_ = v819
	var v820 int32
	_ = v820
	var v821 int32
	_ = v821
	var v828 int32
	_ = v828
	var v830 int32
	_ = v830
	var v831 int32
	_ = v831
	var v833 int32
	_ = v833
	var v842 int32
	_ = v842
	var v845 int32
	_ = v845
	var v851 int32
	_ = v851
	var v861 int32
	_ = v861
	var v866 int32
	_ = v866
	var v870 int32
	_ = v870
	var v875 int32
	_ = v875
	var v877 int32
	_ = v877
	var v881 int32
	_ = v881
	var v887 int32
	_ = v887
	var v890 int32
	_ = v890
	var v896 int32
	_ = v896
	var v900 int32
	_ = v900
	var v902 int32
	_ = v902
	var v910 int32
	_ = v910
	var v912 int32
	_ = v912
	var v915 int32
	_ = v915
	var v916 int32
	_ = v916
	var v920 int32
	_ = v920
	var v921 int32
	_ = v921
	var v923 int32
	_ = v923
	var v924 int32
	_ = v924
	var v926 int32
	_ = v926
	var v927 int32
	_ = v927
	var v928 int32
	_ = v928
	var v929 int32
	_ = v929
	var v934 int32
	_ = v934
	var v942 int32
	_ = v942
	var v946 int32
	_ = v946
	var v948 int32
	_ = v948
	var v949 int32
	_ = v949
	var v953 int32
	_ = v953
	var v954 int32
	_ = v954
	var v956 int32
	_ = v956
	var v960 int32
	_ = v960
	var v962 int32
	_ = v962
	var v964 int32
	_ = v964
	var v967 int32
	_ = v967
	var v972 int32
	_ = v972
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
	var v980 int32
	_ = v980
	var v984 int32
	_ = v984
	var v989 int32
	_ = v989
	var v990 int32
	_ = v990
	var v991 int32
	_ = v991
	var v998 int32
	_ = v998
	var v1000 int32
	_ = v1000
	var v1001 int32
	_ = v1001
	var v1003 int32
	_ = v1003
	var v1011 int32
	_ = v1011
	var v1012 int32
	_ = v1012
	var v1013 int32
	_ = v1013
	var v1025 int32
	_ = v1025
	var v1030 int32
	_ = v1030
	var v1034 int32
	_ = v1034
	var v1039 int32
	_ = v1039
	var v1041 int32
	_ = v1041
	var v1045 int32
	_ = v1045
	var v1051 int32
	_ = v1051
	var v1054 int32
	_ = v1054
	var v1060 int32
	_ = v1060
	var v1064 int32
	_ = v1064
	var v1066 int32
	_ = v1066
	var v1074 int32
	_ = v1074
	var v1076 int32
	_ = v1076
	var v1088 int32
	_ = v1088
	var v1093 int32
	_ = v1093
	var v1097 int32
	_ = v1097
	var v1102 int32
	_ = v1102
	var v1104 int32
	_ = v1104
	var v1108 int32
	_ = v1108
	var v1114 int32
	_ = v1114
	var v1117 int32
	_ = v1117
	var v1123 int32
	_ = v1123
	var v1127 int32
	_ = v1127
	var v1129 int32
	_ = v1129
	var v1137 int32
	_ = v1137
	var v1139 int32
	_ = v1139
	var v1142 int32
	_ = v1142
	var v1143 int32
	_ = v1143
	var v1147 int32
	_ = v1147
	var v1148 int32
	_ = v1148
	var v1150 int32
	_ = v1150
	var v1151 int32
	_ = v1151
	var v1153 int32
	_ = v1153
	var v1154 int32
	_ = v1154
	var v1155 int32
	_ = v1155
	var v1156 int32
	_ = v1156
	var v1162 int32
	_ = v1162
	var v1169 int32
	_ = v1169
	var v1173 int32
	_ = v1173
	var v1175 int32
	_ = v1175
	var v1176 int32
	_ = v1176
	var v1180 int32
	_ = v1180
	var v1181 int32
	_ = v1181
	var v1183 int32
	_ = v1183
	var v1187 int32
	_ = v1187
	var v1189 int32
	_ = v1189
	var v1191 int32
	_ = v1191
	var v1194 int32
	_ = v1194
	var v1199 int32
	_ = v1199
	var v1200 int32
	_ = v1200
	var v1201 int32
	_ = v1201
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
	var v1216 int32
	_ = v1216
	var v1217 int32
	_ = v1217
	var v1218 int32
	_ = v1218
	var v1225 int32
	_ = v1225
	var v1227 int32
	_ = v1227
	var v1228 int32
	_ = v1228
	var v1230 int32
	_ = v1230
	var v1242 int32
	_ = v1242
	var v1244 int32
	_ = v1244
	var v1254 int32
	_ = v1254
	var v1255 int32
	_ = v1255
	var v1259 int32
	_ = v1259
	var v1262 int32
	_ = v1262
	var v1268 int32
	_ = v1268
	var v1273 int32
	_ = v1273
	v18 = m.G0
	v20 = v18 - int32(144)
	m.G0 = v20
	v22 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v23 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	if v23 < v22 {
		goto L2
	} else {
		goto L3
	}
L1:
	;
	v42 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v45 = v41 + v42*int32(24)
	v46 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l3))))
	if v46 != 0 {
		goto L15
	} else {
		goto L16
	}
L2:
	;
	v25 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	v41 = v25
	goto L1
L3:
	;
	goto L4
L4:
	;
	if v22 != 0 {
		goto L6
	} else {
		goto L7
	}
L5:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+8)) = v39
	v41 = v39
	goto L1
L6:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0))) = v22 << (uint(int32(1)) % 32)
	v29 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	v32 = F_repalloc(m, v29, v22*int32(48))
	mBase = m.M
	v33 = m.ExcPending
	if v33 != 0 {
		goto L9
	} else {
		goto L10
	}
L7:
	;
	goto L8
L8:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0))) = int32(16)
	v37 = F_palloc(m, int32(384))
	mBase = m.M
	v38 = m.ExcPending
	if v38 != 0 {
		goto L9
	} else {
		goto L11
	}
L9:
	;
	return
L10:
	;
	v39 = v32
	goto L5
L11:
	;
	v39 = v37
	goto L5
L12:
	;
	v1254 = F_pg_regerror(m, v648, v20+int32(32))
	mBase = m.M
	v1255 = m.ExcPending
	if v1255 != 0 {
		goto L9
	} else {
		goto L376
	}
L13:
	;
	v667 = *(*int32)(unsafe.Add(mBase, uint32(v45)+4))
	v671 = l2 << (uint(int32(1)) % 32)
	v674 = v667&int32(-255) | v671&int32(254)
	v675 = int32(28)
	if v671&v675 != 0 {
		goto L201
	} else {
		goto L202
	}
L14:
	;
	v56 = m.G0
	v58 = v56 - int32(16)
	m.G0 = v58
	v61 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l3))))
	if v61 == int32(0) {
		v148 = int32(1)
		goto L20
	} else {
		goto L21
	}
L15:
	;
	if v46 != int32(46) {
		goto L14
	} else {
		goto L18
	}
L16:
	;
	goto L17
L17:
	;
	v50 = *(*int32)(unsafe.Add(mBase, uint32(v45)+4))
	*(*int32)(unsafe.Add(mBase, uint32(v45)+4)) = v50&int32(-769) | int32(256)
	goto L13
L18:
	;
	v49 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l3)+1)))
	if v49 != 0 {
		goto L14
	} else {
		goto L19
	}
L19:
	;
	goto L17
L20:
	;
	m.G0 = v58 + int32(16)
	v152 = *(*int32)(unsafe.Add(mBase, uint32(v45)+4))
	v154 = v152 & int32(-769)
	if v148 != 0 {
		goto L52
	} else {
		goto L53
	}
L21:
	;
	v72 = l3
	v73 = int32(4)
	v75 = v61
	goto L22
L22:
	;
	switch v73 - int32(1) {
	case 0:
		goto L27
	default:
		goto L26
	case 3:
		goto L28
	}
L23:
	;
	v148 = base.B2i32(v124 == int32(4))
	goto L20
L24:
	;
	v125 = F_pg_mblen_cstr(m, v72)
	mBase = m.M
	v126 = m.ExcPending
	if v126 != 0 {
		goto L9
	} else {
		goto L50
	}
L25:
	;
	v124 = int32(1)
	goto L24
L26:
	;
	if v73&int32(-2) == int32(2) {
		goto L41
	} else {
		goto L42
	}
L27:
	;
	if v75 == int32(94) {
		goto L34
	} else {
		goto L35
	}
L28:
	;
	v84 = F_t_isalpha_cstr(m, v72)
	mBase = m.M
	v85 = m.ExcPending
	if v85 != 0 {
		goto L9
	} else {
		goto L29
	}
L29:
	;
	if v84 != 0 {
		goto L30
	} else {
		goto L31
	}
L30:
	;
	v124 = int32(4)
	goto L24
L31:
	;
	goto L32
L32:
	;
	v87 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v72))))
	if v87 == int32(91) {
		goto L25
	} else {
		goto L33
	}
L33:
	;
	v148 = int32(0)
	goto L20
L34:
	;
	v124 = int32(3)
	goto L24
L35:
	;
	goto L36
L36:
	;
	v94 = F_t_isalpha_cstr(m, v72)
	mBase = m.M
	v95 = m.ExcPending
	if v95 != 0 {
		goto L9
	} else {
		goto L37
	}
L37:
	;
	if v94 != 0 {
		goto L38
	} else {
		goto L39
	}
L38:
	;
	v124 = int32(2)
	goto L24
L39:
	;
	goto L40
L40:
	;
	v148 = int32(0)
	goto L20
L41:
	;
	v102 = F_t_isalpha_cstr(m, v72)
	mBase = m.M
	v103 = m.ExcPending
	if v103 != 0 {
		goto L9
	} else {
		goto L44
	}
L42:
	;
	goto L43
L43:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v112 = m.ExcPending
	if v112 != 0 {
		goto L9
	} else {
		goto L47
	}
L44:
	;
	if v102 != 0 {
		v124 = v73
		goto L24
	} else {
		goto L45
	}
L45:
	;
	v105 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v72))))
	if v105 != int32(93) {
		v148 = int32(0)
		goto L20
	} else {
		goto L46
	}
L46:
	;
	v124 = int32(4)
	goto L24
L47:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v58))) = int32(3)
	F_errmsg_internal(m, int32(_a_F_NIAddAffix_0), v58)
	mBase = m.M
	v117 = m.ExcPending
	if v117 != 0 {
		goto L9
	} else {
		goto L48
	}
L48:
	;
	F_errfinish(m, int32(_a_F_NIAddAffix_1), int32(66), int32(_a_F_NIAddAffix_2))
	mBase = m.M
	v122 = m.ExcPending
	if v122 != 0 {
		goto L9
	} else {
		goto L49
	}
L49:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L50:
	;
	v127 = v125 + v72
	v128 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v127))))
	if v128 != 0 {
		v72 = v127
		v73 = v124
		v75 = v128
		goto L22
	} else {
		goto L51
	}
L51:
	;
	goto L23
L52:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v45)+4)) = v154 | int32(512)
	v159 = v45 + int32(16)
	v163 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l3))))
	if v163 != 0 {
		goto L55
	} else {
		goto L56
	}
L53:
	;
	goto L54
L54:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v45)+4)) = v154
	v506 = *(*int32)(unsafe.Add(mBase, uint32(l0)+64))
	if l3&int32(3) == int32(0) {
		v530 = l3
		goto L159
	} else {
		goto L160
	}
L55:
	;
	v164 = l3
	goto L57
L56:
	;
	v164 = int32(_a_F_NIAddAffix_3)
	goto L57
L57:
	;
	v165 = m.G0
	v167 = v165 - int32(80)
	m.G0 = v167
	if v164&int32(3) == int32(0) {
		v192 = v164
		goto L60
	} else {
		goto L61
	}
L58:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v159))) = int64(0)
	*(*int32)(unsafe.Add(mBase, uint32(v159)+4)) = base.B2i32(l6 != int32(0))
	v229 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v164))))
	if v229 != 0 {
		goto L77
	} else {
		goto L78
	}
L59:
	;
	v225 = v217 - v164
	goto L58
L60:
	;
	v196 = v192
	goto L69
L61:
	;
	v176 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v164))))
	if v176 == int32(0) {
		goto L62
	} else {
		goto L63
	}
L62:
	;
	v225 = int32(0)
	goto L58
L63:
	;
	goto L64
L64:
	;
	v181 = v164
	goto L65
L65:
	;
	v185 = v181 + int32(1)
	if v185&int32(3) == int32(0) {
		v192 = v185
		goto L60
	} else {
		goto L67
	}
L66:
	;
	v217 = v185
	goto L59
L67:
	;
	v190 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v185))))
	if v190 != 0 {
		v181 = v185
		goto L65
	} else {
		goto L68
	}
L68:
	;
	goto L66
L69:
	;
	v202 = *(*int32)(unsafe.Add(mBase, uint32(v196)))
	v205 = int32(-2139062144)
	if (int32(16843008)-v202|v202)&v205 == v205 {
		v196 = v196 + int32(4)
		goto L69
	} else {
		goto L71
	}
L70:
	;
	v211 = v196
	goto L72
L71:
	;
	goto L70
L72:
	;
	v215 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v211))))
	if v215 != 0 {
		v211 = v211 + int32(1)
		goto L72
	} else {
		goto L74
	}
L73:
	;
	v217 = v211
	goto L59
L74:
	;
	goto L73
L75:
	;
	goto L13
L76:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v495 = m.ExcPending
	if v495 != 0 {
		goto L9
	} else {
		goto L154
	}
L77:
	;
	v231 = v225 + int32(9)
	v237 = int32(0)
	v241 = v229
	v242 = v164
	v246 = int32(4)
	goto L80
L78:
	;
	goto L79
L79:
	;
	v444 = *(*int32)(unsafe.Add(mBase, uint32(v159)))
	if v444 != 0 {
		goto L148
	} else {
		goto L149
	}
L80:
	;
	switch v246 - int32(1) {
	case 0:
		goto L84
	default:
		goto L83
	case 3:
		goto L85
	}
L81:
	;
	if v419 != int32(4) {
		goto L76
	} else {
		goto L147
	}
L82:
	;
	v421 = F_pg_mblen_cstr(m, v242)
	mBase = m.M
	v422 = m.ExcPending
	if v422 != 0 {
		goto L9
	} else {
		goto L145
	}
L83:
	;
	if v246&int32(-2) == int32(2) {
		goto L126
	} else {
		goto L127
	}
L84:
	;
	if v241&int32(255) == int32(94) {
		goto L111
	} else {
		goto L112
	}
L85:
	;
	v253 = F_t_isalpha_cstr(m, v242)
	mBase = m.M
	v254 = m.ExcPending
	if v254 != 0 {
		goto L9
	} else {
		goto L86
	}
L86:
	;
	if v253 != 0 {
		goto L87
	} else {
		goto L88
	}
L87:
	;
	v255 = F_palloc0(m, v231)
	mBase = m.M
	v256 = m.ExcPending
	if v256 != 0 {
		goto L9
	} else {
		goto L90
	}
L88:
	;
	goto L89
L89:
	;
	v281 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v242))))
	if v281 == int32(91) {
		goto L100
	} else {
		goto L101
	}
L90:
	;
	if v237 != 0 {
		goto L92
	} else {
		goto L93
	}
L91:
	;
	v259 = *(*int32)(unsafe.Add(mBase, uint32(v255)))
	*(*int32)(unsafe.Add(mBase, uint32(v255))) = v259&int32(-4) | int32(1)
	v267 = F_pg_mblen_cstr(m, v242)
	mBase = m.M
	v268 = m.ExcPending
	if v268 != 0 {
		goto L9
	} else {
		goto L95
	}
L92:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v237)+4)) = v255
	goto L91
L93:
	;
	goto L94
L94:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v159))) = v255
	goto L91
L95:
	;
	if v267 != 0 {
		goto L97
	} else {
		goto L98
	}
L96:
	;
	v271 = *(*int32)(unsafe.Add(mBase, uint32(v255)))
	*(*int32)(unsafe.Add(mBase, uint32(v255))) = v271&int32(-262141) | v267<<(uint(int32(2))%32)&int32(_a_F_NIAddAffix_4)
	v417 = v255
	v419 = int32(4)
	goto L82
L97:
	;
	v269 = F__emscripten_memcpy_bulkmem(m, v255+int32(8), v242, v267)
	mBase = m.M
	goto L99
L98:
	;
	goto L99
L99:
	;
	goto L96
L100:
	;
	v284 = F_palloc0(m, v231)
	mBase = m.M
	v285 = m.ExcPending
	if v285 != 0 {
		goto L9
	} else {
		goto L103
	}
L101:
	;
	goto L102
L102:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v298 = m.ExcPending
	if v298 != 0 {
		goto L9
	} else {
		goto L108
	}
L103:
	;
	if v237 != 0 {
		goto L105
	} else {
		goto L106
	}
L104:
	;
	v288 = int32(1)
	v289 = *(*int32)(unsafe.Add(mBase, uint32(v284)))
	*(*int32)(unsafe.Add(mBase, uint32(v284))) = v289&int32(-4) | v288
	v417 = v284
	v419 = v288
	goto L82
L105:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v237)+4)) = v284
	goto L104
L106:
	;
	goto L107
L107:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v159))) = v284
	goto L104
L108:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v167)+48)) = v164
	F_errmsg_internal(m, int32(_a_F_NIAddAffix_5), v167+int32(48))
	mBase = m.M
	v304 = m.ExcPending
	if v304 != 0 {
		goto L9
	} else {
		goto L109
	}
L109:
	;
	F_errfinish(m, int32(_a_F_NIAddAffix_1), int32(118), int32(_a_F_NIAddAffix_6))
	mBase = m.M
	v309 = m.ExcPending
	if v309 != 0 {
		goto L9
	} else {
		goto L110
	}
L110:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L111:
	;
	v314 = *(*int32)(unsafe.Add(mBase, uint32(v237)))
	*(*int32)(unsafe.Add(mBase, uint32(v237))) = v314&int32(-4) | int32(2)
	v417 = v237
	v419 = int32(3)
	goto L82
L112:
	;
	goto L113
L113:
	;
	v321 = F_t_isalpha_cstr(m, v242)
	mBase = m.M
	v322 = m.ExcPending
	if v322 != 0 {
		goto L9
	} else {
		goto L114
	}
L114:
	;
	if v321 != 0 {
		goto L115
	} else {
		goto L116
	}
L115:
	;
	v325 = F_pg_mblen_cstr(m, v242)
	mBase = m.M
	v326 = m.ExcPending
	if v326 != 0 {
		goto L9
	} else {
		goto L118
	}
L116:
	;
	goto L117
L117:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v342 = m.ExcPending
	if v342 != 0 {
		goto L9
	} else {
		goto L123
	}
L118:
	;
	if v325 != 0 {
		goto L120
	} else {
		goto L121
	}
L119:
	;
	v329 = int32(2)
	v330 = *(*int32)(unsafe.Add(mBase, uint32(v237)))
	*(*int32)(unsafe.Add(mBase, uint32(v237))) = v330&int32(-262141) | v325<<(uint(v329)%32)&int32(_a_F_NIAddAffix_4)
	v417 = v237
	v419 = v329
	goto L82
L120:
	;
	v327 = F__emscripten_memcpy_bulkmem(m, v237+int32(8), v242, v325)
	mBase = m.M
	goto L122
L121:
	;
	goto L122
L122:
	;
	goto L119
L123:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v167)+64)) = v164
	F_errmsg_internal(m, int32(_a_F_NIAddAffix_5), v167-int32(-64))
	mBase = m.M
	v348 = m.ExcPending
	if v348 != 0 {
		goto L9
	} else {
		goto L124
	}
L124:
	;
	F_errfinish(m, int32(_a_F_NIAddAffix_1), int32(133), int32(_a_F_NIAddAffix_6))
	mBase = m.M
	v353 = m.ExcPending
	if v353 != 0 {
		goto L9
	} else {
		goto L125
	}
L125:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L126:
	;
	v358 = F_t_isalpha_cstr(m, v242)
	mBase = m.M
	v359 = m.ExcPending
	if v359 != 0 {
		goto L9
	} else {
		goto L129
	}
L127:
	;
	goto L128
L128:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v404 = m.ExcPending
	if v404 != 0 {
		goto L9
	} else {
		goto L142
	}
L129:
	;
	if v358 != 0 {
		goto L130
	} else {
		goto L131
	}
L130:
	;
	v360 = *(*int32)(unsafe.Add(mBase, uint32(v237)))
	v368 = F_pg_mblen_cstr(m, v242)
	mBase = m.M
	v369 = m.ExcPending
	if v369 != 0 {
		goto L9
	} else {
		goto L133
	}
L131:
	;
	goto L132
L132:
	;
	v383 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v242))))
	if v383 == int32(93) {
		v417 = v237
		v419 = int32(4)
		goto L82
	} else {
		goto L138
	}
L133:
	;
	if v368 != 0 {
		goto L135
	} else {
		goto L136
	}
L134:
	;
	v372 = *(*int32)(unsafe.Add(mBase, uint32(v237)))
	*(*int32)(unsafe.Add(mBase, uint32(v237))) = (v372+v368<<(uint(int32(2))%32))&int32(_a_F_NIAddAffix_4) | v372&int32(-262141)
	v417 = v237
	v419 = v246
	goto L82
L135:
	;
	v370 = F__emscripten_memcpy_bulkmem(m, v237+int32(base.Ui32(v360)>>(uint(int32(2))%32))&int32(_a_F_NIAddAffix_7)+int32(8), v242, v368)
	mBase = m.M
	goto L137
L136:
	;
	goto L137
L137:
	;
	goto L134
L138:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v389 = m.ExcPending
	if v389 != 0 {
		goto L9
	} else {
		goto L139
	}
L139:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v167)+16)) = v164
	F_errmsg_internal(m, int32(_a_F_NIAddAffix_5), v167+int32(16))
	mBase = m.M
	v395 = m.ExcPending
	if v395 != 0 {
		goto L9
	} else {
		goto L140
	}
L140:
	;
	F_errfinish(m, int32(_a_F_NIAddAffix_1), int32(142), int32(_a_F_NIAddAffix_6))
	mBase = m.M
	v400 = m.ExcPending
	if v400 != 0 {
		goto L9
	} else {
		goto L141
	}
L141:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L142:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v167)+32)) = int32(3)
	F_errmsg_internal(m, int32(_a_F_NIAddAffix_8), v167+int32(32))
	mBase = m.M
	v411 = m.ExcPending
	if v411 != 0 {
		goto L9
	} else {
		goto L143
	}
L143:
	;
	F_errfinish(m, int32(_a_F_NIAddAffix_1), int32(145), int32(_a_F_NIAddAffix_6))
	mBase = m.M
	v416 = m.ExcPending
	if v416 != 0 {
		goto L9
	} else {
		goto L144
	}
L144:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L145:
	;
	v423 = v421 + v242
	v424 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v423))))
	if v424 != 0 {
		v237 = v417
		v241 = v424
		v242 = v423
		v246 = v419
		goto L80
	} else {
		goto L146
	}
L146:
	;
	goto L81
L147:
	;
	goto L79
L148:
	;
	v445 = *(*int32)(unsafe.Add(mBase, uint32(v159)+4))
	v449 = v444
	v454 = v445
	goto L151
L149:
	;
	goto L150
L150:
	;
	m.G0 = v167 + int32(80)
	goto L75
L151:
	;
	v469 = (v454+int32(2))&int32(_a_F_NIAddAffix_9) | v454&int32(-131071)
	*(*int32)(unsafe.Add(mBase, uint32(v159)+4)) = v469
	v471 = *(*int32)(unsafe.Add(mBase, uint32(v449)+4))
	if v471 != 0 {
		v449 = v471
		v454 = v469
		goto L151
	} else {
		goto L153
	}
L152:
	;
	goto L150
L153:
	;
	goto L152
L154:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v167))) = v164
	F_errmsg_internal(m, int32(_a_F_NIAddAffix_5), v167)
	mBase = m.M
	v499 = m.ExcPending
	if v499 != 0 {
		goto L9
	} else {
		goto L155
	}
L155:
	;
	F_errfinish(m, int32(_a_F_NIAddAffix_1), int32(150), int32(_a_F_NIAddAffix_6))
	mBase = m.M
	v504 = m.ExcPending
	if v504 != 0 {
		goto L9
	} else {
		goto L156
	}
L156:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L157:
	;
	v566 = F_MemoryContextAlloc(m, v506, v563+int32(3))
	mBase = m.M
	v567 = m.ExcPending
	if v567 != 0 {
		goto L9
	} else {
		goto L174
	}
L158:
	;
	v563 = v555 - l3
	goto L157
L159:
	;
	v534 = v530
	goto L168
L160:
	;
	v514 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l3))))
	if v514 == int32(0) {
		goto L161
	} else {
		goto L162
	}
L161:
	;
	v563 = int32(0)
	goto L157
L162:
	;
	goto L163
L163:
	;
	v519 = l3
	goto L164
L164:
	;
	v523 = v519 + int32(1)
	if v523&int32(3) == int32(0) {
		v530 = v523
		goto L159
	} else {
		goto L166
	}
L165:
	;
	v555 = v523
	goto L158
L166:
	;
	v528 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v523))))
	if v528 != 0 {
		v519 = v523
		goto L164
	} else {
		goto L167
	}
L167:
	;
	goto L165
L168:
	;
	v540 = *(*int32)(unsafe.Add(mBase, uint32(v534)))
	v543 = int32(-2139062144)
	if (int32(16843008)-v540|v540)&v543 == v543 {
		v534 = v534 + int32(4)
		goto L168
	} else {
		goto L170
	}
L169:
	;
	v549 = v534
	goto L171
L170:
	;
	goto L169
L171:
	;
	v553 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v549))))
	if v553 != 0 {
		v549 = v549 + int32(1)
		goto L171
	} else {
		goto L173
	}
L172:
	;
	v555 = v549
	goto L158
L173:
	;
	goto L172
L174:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v20)+16)) = l3
	if l6 != 0 {
		goto L175
	} else {
		goto L176
	}
L175:
	;
	v571 = int32(_a_F_NIAddAffix_10)
	goto L177
L176:
	;
	v571 = int32(_a_F_NIAddAffix_11)
	goto L177
L177:
	;
	v574 = F_pg_sprintf(m, v566, v571, v20+int32(16))
	mBase = m.M
	v575 = m.ExcPending
	if v575 != 0 {
		goto L9
	} else {
		goto L178
	}
L178:
	;
	v576 = *(*int32)(unsafe.Add(mBase, uint32(l0)+64))
	if v566&int32(3) == int32(0) {
		v600 = v566
		goto L181
	} else {
		goto L182
	}
L179:
	;
	v638 = F_MemoryContextAlloc(m, v576, v633<<(uint(int32(2))%32)+int32(4))
	mBase = m.M
	v639 = m.ExcPending
	if v639 != 0 {
		goto L9
	} else {
		goto L196
	}
L180:
	;
	v633 = v625 - v566
	goto L179
L181:
	;
	v604 = v600
	goto L190
L182:
	;
	v584 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v566))))
	if v584 == int32(0) {
		goto L183
	} else {
		goto L184
	}
L183:
	;
	v633 = int32(0)
	goto L179
L184:
	;
	goto L185
L185:
	;
	v589 = v566
	goto L186
L186:
	;
	v593 = v589 + int32(1)
	if v593&int32(3) == int32(0) {
		v600 = v593
		goto L181
	} else {
		goto L188
	}
L187:
	;
	v625 = v593
	goto L180
L188:
	;
	v598 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v593))))
	if v598 != 0 {
		v589 = v593
		goto L186
	} else {
		goto L189
	}
L189:
	;
	goto L187
L190:
	;
	v610 = *(*int32)(unsafe.Add(mBase, uint32(v604)))
	v613 = int32(-2139062144)
	if (int32(16843008)-v610|v610)&v613 == v613 {
		v604 = v604 + int32(4)
		goto L190
	} else {
		goto L192
	}
L191:
	;
	v619 = v604
	goto L193
L192:
	;
	goto L191
L193:
	;
	v623 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v619))))
	if v623 != 0 {
		v619 = v619 + int32(1)
		goto L193
	} else {
		goto L195
	}
L194:
	;
	v625 = v619
	goto L180
L195:
	;
	goto L194
L196:
	;
	v640 = F_pg_mb2wchar_with_len(m, v566, v638, v633)
	mBase = m.M
	v641 = m.ExcPending
	if v641 != 0 {
		goto L9
	} else {
		goto L197
	}
L197:
	;
	v643 = F_palloc(m, int32(32))
	mBase = m.M
	v644 = m.ExcPending
	if v644 != 0 {
		goto L9
	} else {
		goto L198
	}
L198:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v45)+16)) = v643
	v648 = F_pg_regcomp(m, v643, v638, v640, int32(19), int32(100))
	mBase = m.M
	v649 = m.ExcPending
	if v649 != 0 {
		goto L9
	} else {
		goto L199
	}
L199:
	;
	if v648 != 0 {
		goto L12
	} else {
		goto L200
	}
L200:
	;
	goto L13
L201:
	;
	v679 = v674
	goto L203
L202:
	;
	v679 = v674 | v675
	goto L203
L203:
	;
	if v671&int32(34) != 0 {
		goto L204
	} else {
		goto L205
	}
L204:
	;
	v682 = v679
	goto L206
L205:
	;
	v682 = v674
	goto L206
L206:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v45)+4)) = v682
	if l1&int32(3) == int32(0) {
		v707 = l1
		goto L210
	} else {
		goto L211
	}
L207:
	;
	if (l1^v765)&int32(3) != 0 {
		goto L237
	} else {
		goto L238
	}
L208:
	;
	v742 = v740 + int32(1)
	if base.Ui32(int32(1025)) <= base.Ui32(v742) {
		goto L225
	} else {
		goto L226
	}
L209:
	;
	v740 = v732 - l1
	goto L208
L210:
	;
	v711 = v707
	goto L219
L211:
	;
	v691 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l1))))
	if v691 == int32(0) {
		goto L212
	} else {
		goto L213
	}
L212:
	;
	v740 = int32(0)
	goto L208
L213:
	;
	goto L214
L214:
	;
	v696 = l1
	goto L215
L215:
	;
	v700 = v696 + int32(1)
	if v700&int32(3) == int32(0) {
		v707 = v700
		goto L210
	} else {
		goto L217
	}
L216:
	;
	v732 = v700
	goto L209
L217:
	;
	v705 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v700))))
	if v705 != 0 {
		v696 = v700
		goto L215
	} else {
		goto L218
	}
L218:
	;
	goto L216
L219:
	;
	v717 = *(*int32)(unsafe.Add(mBase, uint32(v711)))
	v720 = int32(-2139062144)
	if (int32(16843008)-v717|v717)&v720 == v720 {
		v711 = v711 + int32(4)
		goto L219
	} else {
		goto L221
	}
L220:
	;
	v726 = v711
	goto L222
L221:
	;
	goto L220
L222:
	;
	v730 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v726))))
	if v730 != 0 {
		v726 = v726 + int32(1)
		goto L222
	} else {
		goto L224
	}
L223:
	;
	v732 = v726
	goto L209
L224:
	;
	goto L223
L225:
	;
	v745 = F_palloc0(m, v742)
	mBase = m.M
	v746 = m.ExcPending
	if v746 != 0 {
		goto L9
	} else {
		goto L228
	}
L226:
	;
	goto L227
L227:
	;
	v750 = (v740 + int32(8)) & int32(4088)
	v751 = *(*int32)(unsafe.Add(mBase, uint32(l0)+84))
	if base.Ui32(v750) <= base.Ui32(v751) {
		goto L230
	} else {
		goto L231
	}
L228:
	;
	v765 = v745
	goto L207
L229:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+84)) = v758 - v750
	*(*int32)(unsafe.Add(mBase, uint32(l0)+80)) = v750 + v759
	v765 = v759
	goto L207
L230:
	;
	v753 = *(*int32)(unsafe.Add(mBase, uint32(l0)+80))
	v758 = v751
	v759 = v753
	goto L229
L231:
	;
	goto L232
L232:
	;
	v754 = int32(_a_F_NIAddAffix_12)
	v756 = F_palloc0(m, v754)
	mBase = m.M
	v757 = m.ExcPending
	if v757 != 0 {
		goto L9
	} else {
		goto L233
	}
L233:
	;
	v758 = v754
	v759 = v756
	goto L229
L234:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v45))) = v765
	v842 = *(*int32)(unsafe.Add(mBase, uint32(v45)+4))
	v845 = v842&int32(-2) | l6
	*(*int32)(unsafe.Add(mBase, uint32(v45)+4)) = v845
	if l4 == int32(0) {
		goto L256
	} else {
		goto L257
	}
L235:
	;
	goto L234
L236:
	;
	*(*uint8)(unsafe.Add(mBase, uint32(v821))) = uint8(v820)
	if v820&int32(255) == int32(0) {
		goto L235
	} else {
		goto L251
	}
L237:
	;
	v772 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l1))))
	v819 = l1
	v820 = v772
	v821 = v765
	goto L236
L238:
	;
	goto L239
L239:
	;
	if l1&int32(3) != 0 {
		goto L240
	} else {
		goto L241
	}
L240:
	;
	v776 = l1
	v778 = v765
	goto L243
L241:
	;
	v790 = l1
	v792 = v765
	goto L242
L242:
	;
	v794 = *(*int32)(unsafe.Add(mBase, uint32(v790)))
	v797 = int32(-2139062144)
	if (int32(16843008)-v794|v794)&v797 != v797 {
		v819 = v790
		v820 = v794
		v821 = v792
		goto L236
	} else {
		goto L247
	}
L243:
	;
	v779 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v776))))
	*(*uint8)(unsafe.Add(mBase, uint32(v778))) = uint8(v779)
	if v779 == int32(0) {
		goto L235
	} else {
		goto L245
	}
L244:
	;
	v790 = v786
	v792 = v784
	goto L242
L245:
	;
	v783 = int32(1)
	v784 = v778 + v783
	v786 = v776 + v783
	if v786&int32(3) != 0 {
		v776 = v786
		v778 = v784
		goto L243
	} else {
		goto L246
	}
L246:
	;
	goto L244
L247:
	;
	v802 = v790
	v803 = v794
	v804 = v792
	goto L248
L248:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v804))) = v803
	v806 = int32(4)
	v807 = v804 + v806
	v808 = *(*int32)(unsafe.Add(mBase, uint32(v802)+4))
	v810 = v802 + v806
	v814 = int32(-2139062144)
	if (v808|(int32(16843008)-v808))&v814 == v814 {
		v802 = v810
		v803 = v808
		v804 = v807
		goto L248
	} else {
		goto L250
	}
L249:
	;
	v819 = v810
	v820 = v808
	v821 = v807
	goto L236
L250:
	;
	goto L249
L251:
	;
	v828 = v819
	v830 = v821
	goto L252
L252:
	;
	v831 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v828)+1)))
	*(*uint8)(unsafe.Add(mBase, uint32(v830)+1)) = uint8(v831)
	v833 = int32(1)
	if v831 != 0 {
		v828 = v828 + v833
		v830 = v830 + v833
		goto L252
	} else {
		goto L254
	}
L253:
	;
	goto L235
L254:
	;
	goto L253
L255:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v45)+8)) = v1013
	if l5&int32(3) == int32(0) {
		v1041 = l5
		goto L310
	} else {
		goto L311
	}
L256:
	;
	v1012 = v845
	v1013 = int32(_a_F_NIAddAffix_3)
	goto L255
L257:
	;
	goto L258
L258:
	;
	v851 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l4))))
	if v851 == int32(0) {
		v1012 = v845
		v1013 = int32(_a_F_NIAddAffix_3)
		goto L255
	} else {
		goto L259
	}
L259:
	;
	if l4&int32(3) == int32(0) {
		v877 = l4
		goto L263
	} else {
		goto L264
	}
L260:
	;
	if (l4^v934)&int32(3) != 0 {
		goto L290
	} else {
		goto L291
	}
L261:
	;
	v912 = v910 + int32(1)
	if base.Ui32(int32(1025)) <= base.Ui32(v912) {
		goto L278
	} else {
		goto L279
	}
L262:
	;
	v910 = v902 - l4
	goto L261
L263:
	;
	v881 = v877
	goto L272
L264:
	;
	v861 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l4))))
	if v861 == int32(0) {
		goto L265
	} else {
		goto L266
	}
L265:
	;
	v910 = int32(0)
	goto L261
L266:
	;
	goto L267
L267:
	;
	v866 = l4
	goto L268
L268:
	;
	v870 = v866 + int32(1)
	if v870&int32(3) == int32(0) {
		v877 = v870
		goto L263
	} else {
		goto L270
	}
L269:
	;
	v902 = v870
	goto L262
L270:
	;
	v875 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v870))))
	if v875 != 0 {
		v866 = v870
		goto L268
	} else {
		goto L271
	}
L271:
	;
	goto L269
L272:
	;
	v887 = *(*int32)(unsafe.Add(mBase, uint32(v881)))
	v890 = int32(-2139062144)
	if (int32(16843008)-v887|v887)&v890 == v890 {
		v881 = v881 + int32(4)
		goto L272
	} else {
		goto L274
	}
L273:
	;
	v896 = v881
	goto L275
L274:
	;
	goto L273
L275:
	;
	v900 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v896))))
	if v900 != 0 {
		v896 = v896 + int32(1)
		goto L275
	} else {
		goto L277
	}
L276:
	;
	v902 = v896
	goto L262
L277:
	;
	goto L276
L278:
	;
	v915 = F_palloc0(m, v912)
	mBase = m.M
	v916 = m.ExcPending
	if v916 != 0 {
		goto L9
	} else {
		goto L281
	}
L279:
	;
	goto L280
L280:
	;
	v920 = (v910 + int32(8)) & int32(4088)
	v921 = *(*int32)(unsafe.Add(mBase, uint32(l0)+84))
	if base.Ui32(v920) <= base.Ui32(v921) {
		goto L283
	} else {
		goto L284
	}
L281:
	;
	v934 = v915
	goto L260
L282:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+84)) = v928 - v920
	*(*int32)(unsafe.Add(mBase, uint32(l0)+80)) = v929 + v920
	v934 = v929
	goto L260
L283:
	;
	v923 = *(*int32)(unsafe.Add(mBase, uint32(l0)+80))
	v928 = v921
	v929 = v923
	goto L282
L284:
	;
	goto L285
L285:
	;
	v924 = int32(_a_F_NIAddAffix_12)
	v926 = F_palloc0(m, v924)
	mBase = m.M
	v927 = m.ExcPending
	if v927 != 0 {
		goto L9
	} else {
		goto L286
	}
L286:
	;
	v928 = v924
	v929 = v926
	goto L282
L287:
	;
	v1011 = *(*int32)(unsafe.Add(mBase, uint32(v45)+4))
	v1012 = v1011
	v1013 = v934
	goto L255
L288:
	;
	goto L287
L289:
	;
	*(*uint8)(unsafe.Add(mBase, uint32(v991))) = uint8(v990)
	if v990&int32(255) == int32(0) {
		goto L288
	} else {
		goto L304
	}
L290:
	;
	v942 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l4))))
	v989 = l4
	v990 = v942
	v991 = v934
	goto L289
L291:
	;
	goto L292
L292:
	;
	if l4&int32(3) != 0 {
		goto L293
	} else {
		goto L294
	}
L293:
	;
	v946 = l4
	v948 = v934
	goto L296
L294:
	;
	v960 = l4
	v962 = v934
	goto L295
L295:
	;
	v964 = *(*int32)(unsafe.Add(mBase, uint32(v960)))
	v967 = int32(-2139062144)
	if (int32(16843008)-v964|v964)&v967 != v967 {
		v989 = v960
		v990 = v964
		v991 = v962
		goto L289
	} else {
		goto L300
	}
L296:
	;
	v949 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v946))))
	*(*uint8)(unsafe.Add(mBase, uint32(v948))) = uint8(v949)
	if v949 == int32(0) {
		goto L288
	} else {
		goto L298
	}
L297:
	;
	v960 = v956
	v962 = v954
	goto L295
L298:
	;
	v953 = int32(1)
	v954 = v948 + v953
	v956 = v946 + v953
	if v956&int32(3) != 0 {
		v946 = v956
		v948 = v954
		goto L296
	} else {
		goto L299
	}
L299:
	;
	goto L297
L300:
	;
	v972 = v960
	v973 = v964
	v974 = v962
	goto L301
L301:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v974))) = v973
	v976 = int32(4)
	v977 = v974 + v976
	v978 = *(*int32)(unsafe.Add(mBase, uint32(v972)+4))
	v980 = v972 + v976
	v984 = int32(-2139062144)
	if (v978|(int32(16843008)-v978))&v984 == v984 {
		v972 = v980
		v973 = v978
		v974 = v977
		goto L301
	} else {
		goto L303
	}
L302:
	;
	v989 = v980
	v990 = v978
	v991 = v977
	goto L289
L303:
	;
	goto L302
L304:
	;
	v998 = v989
	v1000 = v991
	goto L305
L305:
	;
	v1001 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v998)+1)))
	*(*uint8)(unsafe.Add(mBase, uint32(v1000)+1)) = uint8(v1001)
	v1003 = int32(1)
	if v1001 != 0 {
		v998 = v998 + v1003
		v1000 = v1000 + v1003
		goto L305
	} else {
		goto L307
	}
L306:
	;
	goto L288
L307:
	;
	goto L306
L308:
	;
	v1076 = v1074 & int32(_a_F_NIAddAffix_13)
	*(*int32)(unsafe.Add(mBase, uint32(v45)+4)) = v1012&int32(-16776193) | v1076<<(uint(int32(10))%32)
	if v1076 != 0 {
		goto L325
	} else {
		goto L326
	}
L309:
	;
	v1074 = v1066 - l5
	goto L308
L310:
	;
	v1045 = v1041
	goto L319
L311:
	;
	v1025 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l5))))
	if v1025 == int32(0) {
		goto L312
	} else {
		goto L313
	}
L312:
	;
	v1074 = int32(0)
	goto L308
L313:
	;
	goto L314
L314:
	;
	v1030 = l5
	goto L315
L315:
	;
	v1034 = v1030 + int32(1)
	if v1034&int32(3) == int32(0) {
		v1041 = v1034
		goto L310
	} else {
		goto L317
	}
L316:
	;
	v1066 = v1034
	goto L309
L317:
	;
	v1039 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1034))))
	if v1039 != 0 {
		v1030 = v1034
		goto L315
	} else {
		goto L318
	}
L318:
	;
	goto L316
L319:
	;
	v1051 = *(*int32)(unsafe.Add(mBase, uint32(v1045)))
	v1054 = int32(-2139062144)
	if (int32(16843008)-v1051|v1051)&v1054 == v1054 {
		v1045 = v1045 + int32(4)
		goto L319
	} else {
		goto L321
	}
L320:
	;
	v1060 = v1045
	goto L322
L321:
	;
	goto L320
L322:
	;
	v1064 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1060))))
	if v1064 != 0 {
		v1060 = v1060 + int32(1)
		goto L322
	} else {
		goto L324
	}
L323:
	;
	v1066 = v1060
	goto L309
L324:
	;
	goto L323
L325:
	;
	if l5&int32(3) == int32(0) {
		v1104 = l5
		goto L331
	} else {
		goto L332
	}
L326:
	;
	v1242 = int32(_a_F_NIAddAffix_3)
	goto L327
L327:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v45)+12)) = v1242
	v1244 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v1244 + int32(1)
	m.G0 = v20 + int32(144)
	return
L328:
	;
	if (l5^v1162)&int32(3) != 0 {
		goto L358
	} else {
		goto L359
	}
L329:
	;
	v1139 = v1137 + int32(1)
	if base.Ui32(int32(1025)) <= base.Ui32(v1139) {
		goto L346
	} else {
		goto L347
	}
L330:
	;
	v1137 = v1129 - l5
	goto L329
L331:
	;
	v1108 = v1104
	goto L340
L332:
	;
	v1088 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l5))))
	if v1088 == int32(0) {
		goto L333
	} else {
		goto L334
	}
L333:
	;
	v1137 = int32(0)
	goto L329
L334:
	;
	goto L335
L335:
	;
	v1093 = l5
	goto L336
L336:
	;
	v1097 = v1093 + int32(1)
	if v1097&int32(3) == int32(0) {
		v1104 = v1097
		goto L331
	} else {
		goto L338
	}
L337:
	;
	v1129 = v1097
	goto L330
L338:
	;
	v1102 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1097))))
	if v1102 != 0 {
		v1093 = v1097
		goto L336
	} else {
		goto L339
	}
L339:
	;
	goto L337
L340:
	;
	v1114 = *(*int32)(unsafe.Add(mBase, uint32(v1108)))
	v1117 = int32(-2139062144)
	if (int32(16843008)-v1114|v1114)&v1117 == v1117 {
		v1108 = v1108 + int32(4)
		goto L340
	} else {
		goto L342
	}
L341:
	;
	v1123 = v1108
	goto L343
L342:
	;
	goto L341
L343:
	;
	v1127 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1123))))
	if v1127 != 0 {
		v1123 = v1123 + int32(1)
		goto L343
	} else {
		goto L345
	}
L344:
	;
	v1129 = v1123
	goto L330
L345:
	;
	goto L344
L346:
	;
	v1142 = F_palloc0(m, v1139)
	mBase = m.M
	v1143 = m.ExcPending
	if v1143 != 0 {
		goto L9
	} else {
		goto L349
	}
L347:
	;
	goto L348
L348:
	;
	v1147 = (v1137 + int32(8)) & int32(4088)
	v1148 = *(*int32)(unsafe.Add(mBase, uint32(l0)+84))
	if base.Ui32(v1147) <= base.Ui32(v1148) {
		goto L351
	} else {
		goto L352
	}
L349:
	;
	v1162 = v1142
	goto L328
L350:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+84)) = v1155 - v1147
	*(*int32)(unsafe.Add(mBase, uint32(l0)+80)) = v1156 + v1147
	v1162 = v1156
	goto L328
L351:
	;
	v1150 = *(*int32)(unsafe.Add(mBase, uint32(l0)+80))
	v1155 = v1148
	v1156 = v1150
	goto L350
L352:
	;
	goto L353
L353:
	;
	v1151 = int32(_a_F_NIAddAffix_12)
	v1153 = F_palloc0(m, v1151)
	mBase = m.M
	v1154 = m.ExcPending
	if v1154 != 0 {
		goto L9
	} else {
		goto L354
	}
L354:
	;
	v1155 = v1151
	v1156 = v1153
	goto L350
L355:
	;
	v1242 = v1162
	goto L327
L356:
	;
	goto L355
L357:
	;
	*(*uint8)(unsafe.Add(mBase, uint32(v1218))) = uint8(v1217)
	if v1217&int32(255) == int32(0) {
		goto L356
	} else {
		goto L372
	}
L358:
	;
	v1169 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l5))))
	v1216 = l5
	v1217 = v1169
	v1218 = v1162
	goto L357
L359:
	;
	goto L360
L360:
	;
	if l5&int32(3) != 0 {
		goto L361
	} else {
		goto L362
	}
L361:
	;
	v1173 = l5
	v1175 = v1162
	goto L364
L362:
	;
	v1187 = l5
	v1189 = v1162
	goto L363
L363:
	;
	v1191 = *(*int32)(unsafe.Add(mBase, uint32(v1187)))
	v1194 = int32(-2139062144)
	if (int32(16843008)-v1191|v1191)&v1194 != v1194 {
		v1216 = v1187
		v1217 = v1191
		v1218 = v1189
		goto L357
	} else {
		goto L368
	}
L364:
	;
	v1176 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1173))))
	*(*uint8)(unsafe.Add(mBase, uint32(v1175))) = uint8(v1176)
	if v1176 == int32(0) {
		goto L356
	} else {
		goto L366
	}
L365:
	;
	v1187 = v1183
	v1189 = v1181
	goto L363
L366:
	;
	v1180 = int32(1)
	v1181 = v1175 + v1180
	v1183 = v1173 + v1180
	if v1183&int32(3) != 0 {
		v1173 = v1183
		v1175 = v1181
		goto L364
	} else {
		goto L367
	}
L367:
	;
	goto L365
L368:
	;
	v1199 = v1187
	v1200 = v1191
	v1201 = v1189
	goto L369
L369:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v1201))) = v1200
	v1203 = int32(4)
	v1204 = v1201 + v1203
	v1205 = *(*int32)(unsafe.Add(mBase, uint32(v1199)+4))
	v1207 = v1199 + v1203
	v1211 = int32(-2139062144)
	if (v1205|(int32(16843008)-v1205))&v1211 == v1211 {
		v1199 = v1207
		v1200 = v1205
		v1201 = v1204
		goto L369
	} else {
		goto L371
	}
L370:
	;
	v1216 = v1207
	v1217 = v1205
	v1218 = v1204
	goto L357
L371:
	;
	goto L370
L372:
	;
	v1225 = v1216
	v1227 = v1218
	goto L373
L373:
	;
	v1228 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1225)+1)))
	*(*uint8)(unsafe.Add(mBase, uint32(v1227)+1)) = uint8(v1228)
	v1230 = int32(1)
	if v1228 != 0 {
		v1225 = v1225 + v1230
		v1227 = v1227 + v1230
		goto L373
	} else {
		goto L375
	}
L374:
	;
	goto L356
L375:
	;
	goto L374
L376:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v1259 = m.ExcPending
	if v1259 != 0 {
		goto L9
	} else {
		goto L377
	}
L377:
	;
	F_errcode(m, int32(302252162))
	mBase = m.M
	v1262 = m.ExcPending
	if v1262 != 0 {
		goto L9
	} else {
		goto L378
	}
L378:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v20))) = v20 + int32(32)
	F_errmsg(m, int32(_a_F_NIAddAffix_14), v20)
	mBase = m.M
	v1268 = m.ExcPending
	if v1268 != 0 {
		goto L9
	} else {
		goto L379
	}
L379:
	;
	F_errfinish(m, int32(_a_F_NIAddAffix_15), int32(752), int32(_a_F_NIAddAffix_16))
	mBase = m.M
	v1273 = m.ExcPending
	if v1273 != 0 {
		goto L9
	} else {
		goto L380
	}
L380:
	;
	base.Wasm_trap_unreachable()
	for {
	}
}
func F_NullCommand(m *base.Module, l0 int32) {
	var v2 int32
	_ = v2
	var v8 int32
	_ = v8
	v2 = int32(2)
	if base.Ui32(l0-v2) <= base.Ui32(v2) {
		F_pq_putemptymessage(m, int32(73))
		v8 = m.ExcPending
		if v8 != 0 {
			return
		} else {
			return
		}
	} else {
		return
	}
}
func F___nl_langinfo_l(m *base.Module, l0 int32, l1 int32) int32 {
	mBase := m.M
	_ = mBase
	var v9 int32
	_ = v9
	var v10 int32
	_ = v10
	var v13 int32
	_ = v13
	var v14 int32
	_ = v14
	var v15 int32
	_ = v15
	var v23 int32
	_ = v23
	var v27 int32
	_ = v27
	var v29 int32
	_ = v29
	var v41 int32
	_ = v41
	var v45 int32
	_ = v45
	var v47 int32
	_ = v47
	var v49 int32
	_ = v49
	var v51 int32
	_ = v51
	var v53 int32
	_ = v53
	var v54 int32
	_ = v54
	if l0 == int32(14) {
		v9 = *(*int32)(unsafe.Add(mBase, uint32(l1)))
		if v9 != 0 {
			v10 = int32(_a_F___nl_langinfo_l_0)
		} else {
			v10 = int32(_a_F___nl_langinfo_l_1)
		}
		return v10
	} else {
		v13 = l0 >> (uint(int32(16)) % 32)
		v14 = int32(_a_F___nl_langinfo_l_2)
		v15 = l0 & v14
		if v15 != v14 {
			v29 = int32(_a_F___nl_langinfo_l_3)
			switch v13 - int32(1) {
			case 0:
				if base.Ui32(int32(1)) < base.Ui32(v15) {
					v54 = v29
					return v54
				} else {
					v41 = int32(_a_F___nl_langinfo_l_4)
					if v15 == int32(0) {
						return v41
					} else {
						v45 = v41
						v47 = v15
						for {
							v49 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v45))))
							v51 = v45 + int32(1)
							if v49 != 0 {
								v45 = v51
								continue
							} else {
							}
							v53 = v47 - int32(1)
							if v53 != 0 {
								v45 = v51
								v47 = v53
								continue
							} else {
								break
							}
							break
						}
						v54 = v51
						return v54
					}
				}
			case 1:
				if base.Ui32(int32(49)) < base.Ui32(v15) {
					v54 = v29
					return v54
				} else {
					v41 = int32(_a_F___nl_langinfo_l_5)
					if v15 == int32(0) {
						return v41
					} else {
						v45 = v41
						v47 = v15
						for {
							v49 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v45))))
							v51 = v45 + int32(1)
							if v49 != 0 {
								v45 = v51
								continue
							} else {
							}
							v53 = v47 - int32(1)
							if v53 != 0 {
								v45 = v51
								v47 = v53
								continue
							} else {
								break
							}
							break
						}
						v54 = v51
						return v54
					}
				}
			default:
				v54 = v29
				return v54
			case 4:
				if base.Ui32(int32(3)) < base.Ui32(v15) {
					v54 = v29
					return v54
				} else {
					v41 = int32(_a_F___nl_langinfo_l_6)
					if v15 == int32(0) {
						return v41
					} else {
						v45 = v41
						v47 = v15
						for {
							v49 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v45))))
							v51 = v45 + int32(1)
							if v49 != 0 {
								v45 = v51
								continue
							} else {
							}
							v53 = v47 - int32(1)
							if v53 != 0 {
								v45 = v51
								v47 = v53
								continue
							} else {
								break
							}
							break
						}
						v54 = v51
						return v54
					}
				}
			}
		} else {
			if int32(5) < v13 {
				v29 = int32(_a_F___nl_langinfo_l_3)
				switch v13 - int32(1) {
				case 0:
					if base.Ui32(int32(1)) < base.Ui32(v15) {
						v54 = v29
						return v54
					} else {
						v41 = int32(_a_F___nl_langinfo_l_4)
						if v15 == int32(0) {
							return v41
						} else {
							v45 = v41
							v47 = v15
							for {
								v49 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v45))))
								v51 = v45 + int32(1)
								if v49 != 0 {
									v45 = v51
									continue
								} else {
								}
								v53 = v47 - int32(1)
								if v53 != 0 {
									v45 = v51
									v47 = v53
									continue
								} else {
									break
								}
								break
							}
							v54 = v51
							return v54
						}
					}
				case 1:
					if base.Ui32(int32(49)) < base.Ui32(v15) {
						v54 = v29
						return v54
					} else {
						v41 = int32(_a_F___nl_langinfo_l_5)
						if v15 == int32(0) {
							return v41
						} else {
							v45 = v41
							v47 = v15
							for {
								v49 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v45))))
								v51 = v45 + int32(1)
								if v49 != 0 {
									v45 = v51
									continue
								} else {
								}
								v53 = v47 - int32(1)
								if v53 != 0 {
									v45 = v51
									v47 = v53
									continue
								} else {
									break
								}
								break
							}
							v54 = v51
							return v54
						}
					}
				default:
					v54 = v29
					return v54
				case 4:
					if base.Ui32(int32(3)) < base.Ui32(v15) {
						v54 = v29
						return v54
					} else {
						v41 = int32(_a_F___nl_langinfo_l_6)
						if v15 == int32(0) {
							return v41
						} else {
							v45 = v41
							v47 = v15
							for {
								v49 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v45))))
								v51 = v45 + int32(1)
								if v49 != 0 {
									v45 = v51
									continue
								} else {
								}
								v53 = v47 - int32(1)
								if v53 != 0 {
									v45 = v51
									v47 = v53
									continue
								} else {
									break
								}
								break
							}
							v54 = v51
							return v54
						}
					}
				}
			} else {
				v23 = *(*int32)(unsafe.Add(mBase, uint32(l1+v13<<(uint(int32(2))%32))))
				if v23 != 0 {
					v27 = v23 + int32(8)
				} else {
					v27 = int32(_a_F___nl_langinfo_l_7)
				}
				return v27
			}
		}
	}
}
func F_namefastcmp_locale(m *base.Module, l0 int32, l1 int32, l2 int32) int32 {
	mBase := m.M
	_ = mBase
	var v11 int32
	_ = v11
	var v16 int32
	_ = v16
	var v20 int32
	_ = v20
	var v25 int32
	_ = v25
	var v27 int32
	_ = v27
	var v31 int32
	_ = v31
	var v37 int32
	_ = v37
	var v40 int32
	_ = v40
	var v46 int32
	_ = v46
	var v50 int32
	_ = v50
	var v52 int32
	_ = v52
	var v60 int32
	_ = v60
	var v68 int32
	_ = v68
	var v73 int32
	_ = v73
	var v77 int32
	_ = v77
	var v82 int32
	_ = v82
	var v84 int32
	_ = v84
	var v88 int32
	_ = v88
	var v94 int32
	_ = v94
	var v97 int32
	_ = v97
	var v103 int32
	_ = v103
	var v107 int32
	_ = v107
	var v109 int32
	_ = v109
	var v117 int32
	_ = v117
	var v118 int32
	_ = v118
	var v121 int32
	_ = v121
	if l0&int32(3) == int32(0) {
		v27 = l0
		goto L3
	} else {
		goto L4
	}
L1:
	;
	if l1&int32(3) == int32(0) {
		v84 = l1
		goto L20
	} else {
		goto L21
	}
L2:
	;
	v60 = v52 - l0
	goto L1
L3:
	;
	v31 = v27
	goto L12
L4:
	;
	v11 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0))))
	if v11 == int32(0) {
		goto L5
	} else {
		goto L6
	}
L5:
	;
	v60 = int32(0)
	goto L1
L6:
	;
	goto L7
L7:
	;
	v16 = l0
	goto L8
L8:
	;
	v20 = v16 + int32(1)
	if v20&int32(3) == int32(0) {
		v27 = v20
		goto L3
	} else {
		goto L10
	}
L9:
	;
	v52 = v20
	goto L2
L10:
	;
	v25 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v20))))
	if v25 != 0 {
		v16 = v20
		goto L8
	} else {
		goto L11
	}
L11:
	;
	goto L9
L12:
	;
	v37 = *(*int32)(unsafe.Add(mBase, uint32(v31)))
	v40 = int32(-2139062144)
	if (int32(16843008)-v37|v37)&v40 == v40 {
		v31 = v31 + int32(4)
		goto L12
	} else {
		goto L14
	}
L13:
	;
	v46 = v31
	goto L15
L14:
	;
	goto L13
L15:
	;
	v50 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v46))))
	if v50 != 0 {
		v46 = v46 + int32(1)
		goto L15
	} else {
		goto L17
	}
L16:
	;
	v52 = v46
	goto L2
L17:
	;
	goto L16
L18:
	;
	v118 = F_varstrfastcmp_locale(m, l0, v60, l1, v117, l2)
	mBase = m.M
	v121 = m.ExcPending
	if v121 != 0 {
		goto L35
	} else {
		goto L36
	}
L19:
	;
	v117 = v109 - l1
	goto L18
L20:
	;
	v88 = v84
	goto L29
L21:
	;
	v68 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l1))))
	if v68 == int32(0) {
		goto L22
	} else {
		goto L23
	}
L22:
	;
	v117 = int32(0)
	goto L18
L23:
	;
	goto L24
L24:
	;
	v73 = l1
	goto L25
L25:
	;
	v77 = v73 + int32(1)
	if v77&int32(3) == int32(0) {
		v84 = v77
		goto L20
	} else {
		goto L27
	}
L26:
	;
	v109 = v77
	goto L19
L27:
	;
	v82 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v77))))
	if v82 != 0 {
		v73 = v77
		goto L25
	} else {
		goto L28
	}
L28:
	;
	goto L26
L29:
	;
	v94 = *(*int32)(unsafe.Add(mBase, uint32(v88)))
	v97 = int32(-2139062144)
	if (int32(16843008)-v94|v94)&v97 == v97 {
		v88 = v88 + int32(4)
		goto L29
	} else {
		goto L31
	}
L30:
	;
	v103 = v88
	goto L32
L31:
	;
	goto L30
L32:
	;
	v107 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v103))))
	if v107 != 0 {
		v103 = v103 + int32(1)
		goto L32
	} else {
		goto L34
	}
L33:
	;
	v109 = v103
	goto L19
L34:
	;
	goto L33
L35:
	;
	return int32(0)
L36:
	;
	return v118
}
func F_namelike(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v6 int32
	_ = v6
	var v7 int32
	_ = v7
	var v8 int32
	_ = v8
	var v11 int32
	_ = v11
	var v13 int32
	_ = v13
	var v21 int32
	_ = v21
	var v26 int32
	_ = v26
	var v30 int32
	_ = v30
	var v35 int32
	_ = v35
	var v37 int32
	_ = v37
	var v41 int32
	_ = v41
	var v47 int32
	_ = v47
	var v50 int32
	_ = v50
	var v56 int32
	_ = v56
	var v60 int32
	_ = v60
	var v62 int32
	_ = v62
	var v70 int32
	_ = v70
	var v73 int32
	_ = v73
	var v75 int32
	_ = v75
	var v76 int32
	_ = v76
	var v79 int32
	_ = v79
	var v81 int32
	_ = v81
	var v90 int32
	_ = v90
	var v93 int32
	_ = v93
	var v94 int32
	_ = v94
	var v98 int32
	_ = v98
	var v104 int32
	_ = v104
	var v105 int32
	_ = v105
	var v106 int32
	_ = v106
	var v107 int32
	_ = v107
	v6 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v7 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
	v8 = F_pg_detoast_datum_packed(m, v7)
	mBase = m.M
	v11 = m.ExcPending
	if v11 != 0 {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	return int32(0)
L2:
	;
	v13 = v8 + int32(1)
	if v6&int32(3) == int32(0) {
		v37 = v6
		goto L5
	} else {
		goto L6
	}
L3:
	;
	v73 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v8))))
	v75 = v73 & int32(1)
	if v75 != 0 {
		goto L20
	} else {
		goto L21
	}
L4:
	;
	v70 = v62 - v6
	goto L3
L5:
	;
	v41 = v37
	goto L14
L6:
	;
	v21 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v6))))
	if v21 == int32(0) {
		goto L7
	} else {
		goto L8
	}
L7:
	;
	v70 = int32(0)
	goto L3
L8:
	;
	goto L9
L9:
	;
	v26 = v6
	goto L10
L10:
	;
	v30 = v26 + int32(1)
	if v30&int32(3) == int32(0) {
		v37 = v30
		goto L5
	} else {
		goto L12
	}
L11:
	;
	v62 = v30
	goto L4
L12:
	;
	v35 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v30))))
	if v35 != 0 {
		v26 = v30
		goto L10
	} else {
		goto L13
	}
L13:
	;
	goto L11
L14:
	;
	v47 = *(*int32)(unsafe.Add(mBase, uint32(v41)))
	v50 = int32(-2139062144)
	if (int32(16843008)-v47|v47)&v50 == v50 {
		v41 = v41 + int32(4)
		goto L14
	} else {
		goto L16
	}
L15:
	;
	v56 = v41
	goto L17
L16:
	;
	goto L15
L17:
	;
	v60 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v56))))
	if v60 != 0 {
		v56 = v56 + int32(1)
		goto L17
	} else {
		goto L19
	}
L18:
	;
	v62 = v56
	goto L4
L19:
	;
	goto L18
L20:
	;
	v76 = v13
	goto L22
L21:
	;
	v76 = v8 + int32(4)
	goto L22
L22:
	;
	if v73 == int32(1) {
		goto L24
	} else {
		goto L25
	}
L23:
	;
	v105 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	v106 = F_GenericMatchText(m, v6, v70, v76, v104, v105)
	mBase = m.M
	v107 = m.ExcPending
	if v107 != 0 {
		goto L1
	} else {
		goto L34
	}
L24:
	;
	v79 = int32(4)
	v81 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v13))))
	if v81&int32(254) == int32(2) {
		goto L27
	} else {
		goto L28
	}
L25:
	;
	goto L26
L26:
	;
	v94 = int32(1)
	if v75 != 0 {
		v104 = int32(base.Ui32(v73)>>(uint(v94)%32)) - v94
		goto L23
	} else {
		goto L33
	}
L27:
	;
	v90 = v79
	goto L29
L28:
	;
	v90 = base.B2i32(v81 == int32(18)) << (uint(v79) % 32)
	goto L29
L29:
	;
	if v81 == int32(1) {
		goto L30
	} else {
		goto L31
	}
L30:
	;
	v93 = v79
	goto L32
L31:
	;
	v93 = v90
	goto L32
L32:
	;
	v104 = v93
	goto L23
L33:
	;
	v98 = *(*int32)(unsafe.Add(mBase, uint32(v8)))
	v104 = int32(base.Ui32(v98)>>(uint(int32(2))%32)) - int32(4)
	goto L23
L34:
	;
	return base.B2i32(v106 == int32(1))
}
func F_nameregexeq(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v5 int32
	_ = v5
	var v6 int32
	_ = v6
	var v7 int32
	_ = v7
	var v10 int32
	_ = v10
	var v18 int32
	_ = v18
	var v23 int32
	_ = v23
	var v27 int32
	_ = v27
	var v32 int32
	_ = v32
	var v34 int32
	_ = v34
	var v38 int32
	_ = v38
	var v44 int32
	_ = v44
	var v47 int32
	_ = v47
	var v53 int32
	_ = v53
	var v57 int32
	_ = v57
	var v59 int32
	_ = v59
	var v67 int32
	_ = v67
	var v69 int32
	_ = v69
	var v70 int32
	_ = v70
	var v71 int32
	_ = v71
	var v76 int32
	_ = v76
	var v77 int32
	_ = v77
	var v78 int32
	_ = v78
	var v79 int32
	_ = v79
	var v80 int32
	_ = v80
	var v83 int32
	_ = v83
	var v84 int32
	_ = v84
	var v86 int32
	_ = v86
	v5 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v6 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
	v7 = F_pg_detoast_datum_packed(m, v6)
	mBase = m.M
	v10 = m.ExcPending
	if v10 != 0 {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	return int32(0)
L2:
	;
	if v5&int32(3) == int32(0) {
		v34 = v5
		goto L5
	} else {
		goto L6
	}
L3:
	;
	v69 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	v70 = F_RE_compile_and_cache(m, v7, int32(19), v69)
	mBase = m.M
	v71 = m.ExcPending
	if v71 != 0 {
		goto L1
	} else {
		goto L20
	}
L4:
	;
	v67 = v59 - v5
	goto L3
L5:
	;
	v38 = v34
	goto L14
L6:
	;
	v18 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v5))))
	if v18 == int32(0) {
		goto L7
	} else {
		goto L8
	}
L7:
	;
	v67 = int32(0)
	goto L3
L8:
	;
	goto L9
L9:
	;
	v23 = v5
	goto L10
L10:
	;
	v27 = v23 + int32(1)
	if v27&int32(3) == int32(0) {
		v34 = v27
		goto L5
	} else {
		goto L12
	}
L11:
	;
	v59 = v27
	goto L4
L12:
	;
	v32 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v27))))
	if v32 != 0 {
		v23 = v27
		goto L10
	} else {
		goto L13
	}
L13:
	;
	goto L11
L14:
	;
	v44 = *(*int32)(unsafe.Add(mBase, uint32(v38)))
	v47 = int32(-2139062144)
	if (int32(16843008)-v44|v44)&v47 == v47 {
		v38 = v38 + int32(4)
		goto L14
	} else {
		goto L16
	}
L15:
	;
	v53 = v38
	goto L17
L16:
	;
	goto L15
L17:
	;
	v57 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v53))))
	if v57 != 0 {
		v53 = v53 + int32(1)
		goto L17
	} else {
		goto L19
	}
L18:
	;
	v59 = v53
	goto L4
L19:
	;
	goto L18
L20:
	;
	v76 = F_palloc(m, v67<<(uint(int32(2))%32)+int32(4))
	mBase = m.M
	v77 = m.ExcPending
	if v77 != 0 {
		goto L1
	} else {
		goto L21
	}
L21:
	;
	v78 = F_pg_mb2wchar_with_len(m, v5, v76, v67)
	mBase = m.M
	v79 = m.ExcPending
	if v79 != 0 {
		goto L1
	} else {
		goto L22
	}
L22:
	;
	v80 = int32(0)
	v83 = F_RE_wchar_execute(m, v76, v78, v80, v80, v80)
	mBase = m.M
	v84 = m.ExcPending
	if v84 != 0 {
		goto L1
	} else {
		goto L23
	}
L23:
	;
	F_pfree(m, v76)
	mBase = m.M
	v86 = m.ExcPending
	if v86 != 0 {
		goto L1
	} else {
		goto L24
	}
L24:
	;
	return v83
}
func F_namestrcpy(m *base.Module, l0 int32, l1 int32) {
	mBase := m.M
	_ = mBase
	var v3 int32
	_ = v3
	var v18 int32
	_ = v18
	var v19 int32
	_ = v19
	var v20 int32
	_ = v20
	var v22 int32
	_ = v22
	var v26 int32
	_ = v26
	var v27 int32
	_ = v27
	var v29 int32
	_ = v29
	var v30 int32
	_ = v30
	var v31 int32
	_ = v31
	var v33 int32
	_ = v33
	var v39 int32
	_ = v39
	var v40 int32
	_ = v40
	var v41 int32
	_ = v41
	var v42 int32
	_ = v42
	var v45 int32
	_ = v45
	var v51 int32
	_ = v51
	var v52 int32
	_ = v52
	var v53 int32
	_ = v53
	var v56 int32
	_ = v56
	var v59 int32
	_ = v59
	var v64 int32
	_ = v64
	var v65 int32
	_ = v65
	var v67 int32
	_ = v67
	var v69 int32
	_ = v69
	var v73 int32
	_ = v73
	var v74 int32
	_ = v74
	var v75 int32
	_ = v75
	var v80 int32
	_ = v80
	var v81 int32
	_ = v81
	var v82 int32
	_ = v82
	var v85 int32
	_ = v85
	var v86 int32
	_ = v86
	var v87 int32
	_ = v87
	var v89 int32
	_ = v89
	var v93 int32
	_ = v93
	var v94 int32
	_ = v94
	var v98 int32
	_ = v98
	var v102 int32
	_ = v102
	var v107 int32
	_ = v107
	var v108 int32
	_ = v108
	var v111 int32
	_ = v111
	var v112 int32
	_ = v112
	v3 = int32(64)
	if (l1^l0)&int32(3) != 0 {
		v73 = l1
		v74 = v3
		v75 = l0
		goto L5
	} else {
		goto L6
	}
L1:
	;
	v112 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(l0)+63)) = uint8(v112)
	return
L2:
	;
	v111 = F___memset(m, v108, int32(0), v107)
	mBase = m.M
	goto L1
L3:
	;
	v107 = int32(0)
	v108 = v102
	goto L2
L4:
	;
	v85 = v80
	v86 = v81
	v87 = v82
	goto L23
L5:
	;
	if v74 == int32(0) {
		v102 = v75
		goto L3
	} else {
		goto L22
	}
L6:
	;
	if l1&int32(3) == int32(0) {
		v39 = l1
		v40 = v3
		v41 = l0
		v42 = int32(1)
		goto L7
	} else {
		goto L8
	}
L7:
	;
	if v42 == int32(0) {
		v102 = v41
		goto L3
	} else {
		goto L15
	}
L8:
	;
	goto L9
L9:
	;
	v18 = l1
	v19 = v3
	v20 = l0
	goto L10
L10:
	;
	v22 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v18))))
	*(*uint8)(unsafe.Add(mBase, uint32(v20))) = uint8(v22)
	if v22 == int32(0) {
		v107 = v19
		v108 = v20
		goto L2
	} else {
		goto L12
	}
L11:
	;
	v39 = v33
	v40 = v29
	v41 = v27
	v42 = v31
	goto L7
L12:
	;
	v26 = int32(1)
	v27 = v20 + v26
	v29 = v19 - v26
	v30 = int32(0)
	v31 = base.B2i32(v29 != v30)
	v33 = v18 + v26
	if v33&int32(3) == v30 {
		v39 = v33
		v40 = v29
		v41 = v27
		v42 = v31
		goto L7
	} else {
		goto L13
	}
L13:
	;
	if v29 != 0 {
		v18 = v33
		v19 = v29
		v20 = v27
		goto L10
	} else {
		goto L14
	}
L14:
	;
	goto L11
L15:
	;
	v45 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v39))))
	if v45 == int32(0) {
		v107 = v40
		v108 = v41
		goto L2
	} else {
		goto L16
	}
L16:
	;
	if base.Ui32(v40) < base.Ui32(int32(4)) {
		v73 = v39
		v74 = v40
		v75 = v41
		goto L5
	} else {
		goto L17
	}
L17:
	;
	v51 = v39
	v52 = v40
	v53 = v41
	goto L18
L18:
	;
	v56 = *(*int32)(unsafe.Add(mBase, uint32(v51)))
	v59 = int32(-2139062144)
	if (int32(16843008)-v56|v56)&v59 != v59 {
		v80 = v51
		v81 = v52
		v82 = v53
		goto L4
	} else {
		goto L20
	}
L19:
	;
	v73 = v67
	v74 = v69
	v75 = v65
	goto L5
L20:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v53))) = v56
	v64 = int32(4)
	v65 = v53 + v64
	v67 = v51 + v64
	v69 = v52 - v64
	if base.Ui32(int32(3)) < base.Ui32(v69) {
		v51 = v67
		v52 = v69
		v53 = v65
		goto L18
	} else {
		goto L21
	}
L21:
	;
	goto L19
L22:
	;
	v80 = v73
	v81 = v74
	v82 = v75
	goto L4
L23:
	;
	v89 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v85))))
	*(*uint8)(unsafe.Add(mBase, uint32(v87))) = uint8(v89)
	if v89 == int32(0) {
		v107 = v86
		v108 = v87
		goto L2
	} else {
		goto L25
	}
L24:
	;
	v102 = v94
	goto L3
L25:
	;
	v93 = int32(1)
	v94 = v87 + v93
	v98 = v86 - v93
	if v98 != 0 {
		v85 = v85 + v93
		v86 = v98
		v87 = v94
		goto L23
	} else {
		goto L26
	}
L26:
	;
	goto L24
}
func F_negate_clause(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v2 int32
	_ = v2
	var v5 int32
	_ = v5
	var v7 int32
	_ = v7
	var v9 int32
	_ = v9
	var v12 int32
	_ = v12
	var v17 int32
	_ = v17
	var v20 int32
	_ = v20
	var v21 int32
	_ = v21
	var v22 int32
	_ = v22
	var v25 int32
	_ = v25
	var v26 int32
	_ = v26
	var v27 int32
	_ = v27
	var v28 int32
	_ = v28
	var v29 int32
	_ = v29
	var v33 int32
	_ = v33
	var v34 int32
	_ = v34
	var v40 int32
	_ = v40
	var v42 int32
	_ = v42
	var v44 int32
	_ = v44
	var v46 int32
	_ = v46
	var v48 int32
	_ = v48
	var v50 int32
	_ = v50
	var v52 int32
	_ = v52
	var v53 int32
	_ = v53
	var v54 int32
	_ = v54
	var v58 int32
	_ = v58
	var v59 int32
	_ = v59
	var v67 int32
	_ = v67
	var v69 int32
	_ = v69
	var v71 int32
	_ = v71
	var v73 int32
	_ = v73
	var v75 int32
	_ = v75
	var v77 int32
	_ = v77
	var v78 int32
	_ = v78
	var v79 int32
	_ = v79
	var v80 int32
	_ = v80
	var v84 int32
	_ = v84
	var v85 int32
	_ = v85
	var v88 int32
	_ = v88
	var v92 int32
	_ = v92
	var v93 int32
	_ = v93
	var v94 int32
	_ = v94
	var v95 int32
	_ = v95
	var v96 int32
	_ = v96
	var v98 int32
	_ = v98
	var v99 int32
	_ = v99
	var v101 int32
	_ = v101
	var v105 int32
	_ = v105
	var v106 int32
	_ = v106
	var v107 int32
	_ = v107
	var v111 int32
	_ = v111
	var v113 int32
	_ = v113
	var v115 int32
	_ = v115
	var v119 int32
	_ = v119
	var v120 int32
	_ = v120
	var v121 int32
	_ = v121
	var v122 int32
	_ = v122
	var v123 int32
	_ = v123
	var v125 int32
	_ = v125
	var v126 int32
	_ = v126
	var v130 int32
	_ = v130
	var v132 int32
	_ = v132
	var v133 int32
	_ = v133
	var v134 int32
	_ = v134
	var v135 int32
	_ = v135
	var v136 int32
	_ = v136
	var v140 int32
	_ = v140
	var v141 int32
	_ = v141
	var v145 int32
	_ = v145
	var v150 int32
	_ = v150
	var v151 int32
	_ = v151
	var v153 int32
	_ = v153
	var v154 int32
	_ = v154
	var v157 int32
	_ = v157
	var v159 int32
	_ = v159
	var v163 int32
	_ = v163
	var v165 int32
	_ = v165
	var v168 int32
	_ = v168
	var v169 int32
	_ = v169
	var v172 int32
	_ = v172
	var v174 int32
	_ = v174
	var v181 int32
	_ = v181
	var v183 int32
	_ = v183
	var v186 int32
	_ = v186
	var v187 int32
	_ = v187
	var v188 int32
	_ = v188
	var v191 int32
	_ = v191
	var v192 int32
	_ = v192
	var v200 int32
	_ = v200
	var v201 int32
	_ = v201
	var v211 int32
	_ = v211
	var v215 int32
	_ = v215
	var v220 int32
	_ = v220
	var v224 int32
	_ = v224
	var v225 int32
	_ = v225
	var v231 int32
	_ = v231
	var v236 int32
	_ = v236
	var v238 int32
	_ = v238
	var v241 int32
	_ = v241
	var v242 int32
	_ = v242
	var v244 int32
	_ = v244
	v2 = int32(0)
	v5 = m.G0
	v7 = v5 - int32(32)
	m.G0 = v7
	if l0 != 0 {
		goto L4
	} else {
		goto L5
	}
L1:
	;
	m.G0 = v7 + int32(32)
	return v244
L2:
	;
	v241 = F_make_orclause(m, v238)
	mBase = m.M
	v242 = m.ExcPending
	if v242 != 0 {
		goto L17
	} else {
		goto L68
	}
L3:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v224 = m.ExcPending
	if v224 != 0 {
		goto L17
	} else {
		goto L65
	}
L4:
	;
	v9 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	switch v9 - int32(7) {
	case 0:
		goto L13
	default:
		goto L7
	case 10:
		goto L12
	case 13:
		goto L11
	case 14:
		goto L10
	case 45:
		goto L9
	case 46:
		goto L8
	}
L5:
	;
	goto L6
L6:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v211 = m.ExcPending
	if v211 != 0 {
		goto L17
	} else {
		goto L62
	}
L7:
	;
	v186 = m.G0
	v187 = int32(16)
	v188 = v186 - v187
	m.G0 = v188
	v191 = F_palloc0(m, v187)
	mBase = m.M
	v192 = m.ExcPending
	if v192 != 0 {
		goto L17
	} else {
		goto L60
	}
L8:
	;
	v168 = F_palloc0(m, int32(16))
	mBase = m.M
	v169 = m.ExcPending
	if v169 != 0 {
		goto L17
	} else {
		goto L58
	}
L9:
	;
	v151 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+12)))
	if v151 != 0 {
		goto L7
	} else {
		goto L56
	}
L10:
	;
	v77 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	switch v77 {
	case 0:
		goto L29
	case 1:
		goto L28
	case 2:
		goto L27
	default:
		goto L26
	}
L11:
	;
	v52 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v53 = F_get_negator(m, v52)
	mBase = m.M
	v54 = m.ExcPending
	if v54 != 0 {
		goto L17
	} else {
		goto L23
	}
L12:
	;
	v27 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v28 = F_get_negator(m, v27)
	mBase = m.M
	v29 = m.ExcPending
	if v29 != 0 {
		goto L17
	} else {
		goto L20
	}
L13:
	;
	v12 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+24)))
	if v12 == int32(1) {
		goto L14
	} else {
		goto L15
	}
L14:
	;
	v17 = F_makeBoolConst(m, int32(0), int32(1))
	mBase = m.M
	v20 = m.ExcPending
	if v20 != 0 {
		goto L17
	} else {
		goto L18
	}
L15:
	;
	goto L16
L16:
	;
	v21 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v22 = int32(0)
	v25 = F_makeBoolConst(m, base.B2i32(v21 == v22), v22)
	mBase = m.M
	v26 = m.ExcPending
	if v26 != 0 {
		goto L17
	} else {
		goto L19
	}
L17:
	;
	return int32(0)
L18:
	;
	v244 = v17
	goto L1
L19:
	;
	v244 = v25
	goto L1
L20:
	;
	if v28 == int32(0) {
		goto L7
	} else {
		goto L21
	}
L21:
	;
	v33 = F_palloc0(m, int32(36))
	mBase = m.M
	v34 = m.ExcPending
	if v34 != 0 {
		goto L17
	} else {
		goto L22
	}
L22:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v33)+8)) = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v33)+4)) = v28
	*(*int32)(unsafe.Add(mBase, uint32(v33))) = int32(17)
	v40 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	*(*int32)(unsafe.Add(mBase, uint32(v33)+12)) = v40
	v42 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+16)))
	*(*uint8)(unsafe.Add(mBase, uint32(v33)+16)) = uint8(v42)
	v44 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	*(*int32)(unsafe.Add(mBase, uint32(v33)+20)) = v44
	v46 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
	*(*int32)(unsafe.Add(mBase, uint32(v33)+24)) = v46
	v48 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
	*(*int32)(unsafe.Add(mBase, uint32(v33)+28)) = v48
	v50 = *(*int32)(unsafe.Add(mBase, uint32(l0)+32))
	*(*int32)(unsafe.Add(mBase, uint32(v33)+32)) = v50
	v244 = v33
	goto L1
L23:
	;
	if v53 == int32(0) {
		goto L7
	} else {
		goto L24
	}
L24:
	;
	v58 = F_palloc0(m, int32(36))
	mBase = m.M
	v59 = m.ExcPending
	if v59 != 0 {
		goto L17
	} else {
		goto L25
	}
L25:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v58)+16)) = int32(0)
	*(*int64)(unsafe.Add(mBase, uint32(v58)+8)) = int64(0)
	*(*int32)(unsafe.Add(mBase, uint32(v58)+4)) = v53
	*(*int32)(unsafe.Add(mBase, uint32(v58))) = int32(20)
	v67 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+20)))
	v69 = v67 ^ int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(v58)+20)) = uint8(v69)
	v71 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
	*(*int32)(unsafe.Add(mBase, uint32(v58)+24)) = v71
	v73 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
	*(*int32)(unsafe.Add(mBase, uint32(v58)+28)) = v73
	v75 = *(*int32)(unsafe.Add(mBase, uint32(l0)+32))
	*(*int32)(unsafe.Add(mBase, uint32(v58)+32)) = v75
	v244 = v58
	goto L1
L26:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v140 = m.ExcPending
	if v140 != 0 {
		goto L17
	} else {
		goto L53
	}
L27:
	;
	v134 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	v135 = *(*int32)(unsafe.Add(mBase, uint32(v134)+12))
	v136 = *(*int32)(unsafe.Add(mBase, uint32(v135)))
	v244 = v136
	goto L1
L28:
	;
	v101 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	if v101 == int32(0) {
		goto L40
	} else {
		goto L41
	}
L29:
	;
	v78 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	if v78 != 0 {
		goto L31
	} else {
		goto L32
	}
L30:
	;
	v84 = v79
	v85 = v2
	goto L35
L31:
	;
	v79 = int32(0)
	v80 = *(*int32)(unsafe.Add(mBase, uint32(v78)+4))
	if v79 < v80 {
		goto L30
	} else {
		goto L34
	}
L32:
	;
	goto L33
L33:
	;
	v238 = v2
	goto L2
L34:
	;
	goto L33
L35:
	;
	v88 = *(*int32)(unsafe.Add(mBase, uint32(v78)+12))
	v92 = *(*int32)(unsafe.Add(mBase, uint32(v88+v84<<(uint(int32(2))%32))))
	v93 = F_negate_clause(m, v92)
	mBase = m.M
	v94 = m.ExcPending
	if v94 != 0 {
		goto L17
	} else {
		goto L37
	}
L36:
	;
	v238 = v95
	goto L2
L37:
	;
	v95 = F_lappend(m, v85, v93)
	mBase = m.M
	v96 = m.ExcPending
	if v96 != 0 {
		goto L17
	} else {
		goto L38
	}
L38:
	;
	v98 = v84 + int32(1)
	v99 = *(*int32)(unsafe.Add(mBase, uint32(v78)+4))
	if v98 < v99 {
		v84 = v98
		v85 = v95
		goto L35
	} else {
		goto L39
	}
L39:
	;
	goto L36
L40:
	;
	v105 = F_make_andclause(m, int32(0))
	mBase = m.M
	v106 = m.ExcPending
	if v106 != 0 {
		goto L17
	} else {
		goto L43
	}
L41:
	;
	goto L42
L42:
	;
	v107 = *(*int32)(unsafe.Add(mBase, uint32(v101)+4))
	if int32(0) < v107 {
		goto L44
	} else {
		goto L45
	}
L43:
	;
	v244 = v105
	goto L1
L44:
	;
	v111 = int32(0)
	v113 = v2
	goto L47
L45:
	;
	v130 = v2
	goto L46
L46:
	;
	v132 = F_make_andclause(m, v130)
	mBase = m.M
	v133 = m.ExcPending
	if v133 != 0 {
		goto L17
	} else {
		goto L52
	}
L47:
	;
	v115 = *(*int32)(unsafe.Add(mBase, uint32(v101)+12))
	v119 = *(*int32)(unsafe.Add(mBase, uint32(v115+v111<<(uint(int32(2))%32))))
	v120 = F_negate_clause(m, v119)
	mBase = m.M
	v121 = m.ExcPending
	if v121 != 0 {
		goto L17
	} else {
		goto L49
	}
L48:
	;
	v130 = v122
	goto L46
L49:
	;
	v122 = F_lappend(m, v113, v120)
	mBase = m.M
	v123 = m.ExcPending
	if v123 != 0 {
		goto L17
	} else {
		goto L50
	}
L50:
	;
	v125 = v111 + int32(1)
	v126 = *(*int32)(unsafe.Add(mBase, uint32(v101)+4))
	if v125 < v126 {
		v111 = v125
		v113 = v122
		goto L47
	} else {
		goto L51
	}
L51:
	;
	goto L48
L52:
	;
	v244 = v132
	goto L1
L53:
	;
	v141 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	*(*int32)(unsafe.Add(mBase, uint32(v7))) = v141
	F_errmsg_internal(m, int32(_a_F_negate_clause_0), v7)
	mBase = m.M
	v145 = m.ExcPending
	if v145 != 0 {
		goto L17
	} else {
		goto L54
	}
L54:
	;
	F_errfinish(m, int32(_a_F_negate_clause_1), int32(195), int32(_a_F_negate_clause_2))
	mBase = m.M
	v150 = m.ExcPending
	if v150 != 0 {
		goto L17
	} else {
		goto L55
	}
L55:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L56:
	;
	v153 = F_palloc0(m, int32(20))
	mBase = m.M
	v154 = m.ExcPending
	if v154 != 0 {
		goto L17
	} else {
		goto L57
	}
L57:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v153))) = int32(52)
	v157 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	*(*int32)(unsafe.Add(mBase, uint32(v153)+4)) = v157
	v159 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	*(*int32)(unsafe.Add(mBase, uint32(v153)+8)) = base.B2i32(v159 == int32(0))
	v163 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+12)))
	*(*uint8)(unsafe.Add(mBase, uint32(v153)+12)) = uint8(v163)
	v165 = *(*int32)(unsafe.Add(mBase, uint32(l0)+16))
	*(*int32)(unsafe.Add(mBase, uint32(v153)+16)) = v165
	v244 = v153
	goto L1
L58:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v168))) = int32(53)
	v172 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	*(*int32)(unsafe.Add(mBase, uint32(v168)+4)) = v172
	v174 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	if base.Ui32(int32(6)) <= base.Ui32(v174) {
		goto L3
	} else {
		goto L59
	}
L59:
	;
	v181 = *(*int32)(unsafe.Add(mBase, uint32(v174<<(uint(int32(2))%32))+uint32(_c_F_negate_clause[0])))
	*(*int32)(unsafe.Add(mBase, uint32(v168)+8)) = v181
	v183 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	*(*int32)(unsafe.Add(mBase, uint32(v168)+12)) = v183
	v244 = v168
	goto L1
L60:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v191))) = int64(8589934613)
	*(*int32)(unsafe.Add(mBase, uint32(v188)+8)) = l0
	*(*int32)(unsafe.Add(mBase, uint32(v188)+12)) = l0
	v200 = F_list_make1_impl(m, int32(1), v188+int32(8))
	mBase = m.M
	v201 = m.ExcPending
	if v201 != 0 {
		goto L17
	} else {
		goto L61
	}
L61:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v191)+12)) = int32(-1)
	*(*int32)(unsafe.Add(mBase, uint32(v191)+8)) = v200
	m.G0 = v188 + int32(16)
	v244 = v191
	goto L1
L62:
	;
	F_errmsg_internal(m, int32(_a_F_negate_clause_3), int32(0))
	mBase = m.M
	v215 = m.ExcPending
	if v215 != 0 {
		goto L17
	} else {
		goto L63
	}
L63:
	;
	F_errfinish(m, int32(_a_F_negate_clause_1), int32(76), int32(_a_F_negate_clause_2))
	mBase = m.M
	v220 = m.ExcPending
	if v220 != 0 {
		goto L17
	} else {
		goto L64
	}
L64:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L65:
	;
	v225 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	*(*int32)(unsafe.Add(mBase, uint32(v7)+16)) = v225
	F_errmsg_internal(m, int32(_a_F_negate_clause_4), v7+int32(16))
	mBase = m.M
	v231 = m.ExcPending
	if v231 != 0 {
		goto L17
	} else {
		goto L66
	}
L66:
	;
	F_errfinish(m, int32(_a_F_negate_clause_1), int32(250), int32(_a_F_negate_clause_2))
	mBase = m.M
	v236 = m.ExcPending
	if v236 != 0 {
		goto L17
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
	v244 = v241
	goto L1
}
func F_neqjoinsel(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v9 int32
	_ = v9
	var v11 int32
	_ = v11
	var v13 int32
	_ = v13
	var v14 int32
	_ = v14
	var v15 int32
	_ = v15
	var v16 int32
	_ = v16
	var v30 int32
	_ = v30
	var v31 int32
	_ = v31
	var v32 int32
	_ = v32
	var v33 int32
	_ = v33
	var v34 int32
	_ = v34
	var v35 int32
	_ = v35
	var v36 int32
	_ = v36
	var v38 float32
	_ = v38
	var v42 float64
	_ = v42
	var v43 int32
	_ = v43
	var v45 int32
	_ = v45
	var v46 int32
	_ = v46
	var v47 int32
	_ = v47
	var v50 int32
	_ = v50
	var v52 int32
	_ = v52
	var v53 int32
	_ = v53
	var v54 int32
	_ = v54
	var v55 int32
	_ = v55
	var v56 int32
	_ = v56
	var v62 int32
	_ = v62
	var v63 int32
	_ = v63
	var v64 float64
	_ = v64
	var v69 float64
	_ = v69
	var v72 int32
	_ = v72
	var v73 int32
	_ = v73
	v9 = m.G0
	v11 = v9 - int32(80)
	m.G0 = v11
	v13 = *(*int32)(unsafe.Add(mBase, uint32(l0)+52))
	v14 = *(*int32)(unsafe.Add(mBase, uint32(l0)+36))
	v15 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v16 = *(*int32)(unsafe.Add(mBase, uint32(l0)+44))
	if v16&int32(_a_F_neqjoinsel_0) == int32(4) {
		F_get_join_variables(m, v15, v14, v13, v11+int32(48), v11+int32(16), v11+int32(15))
		mBase = m.M
		v30 = m.ExcPending
		if v30 != 0 {
			return int32(0)
		} else {
			v31 = *(*int32)(unsafe.Add(mBase, uint32(v11)+24))
			v32 = *(*int32)(unsafe.Add(mBase, uint32(v11)+56))
			v33 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v11)+15)))
			if v33 != 0 {
				v34 = v31
			} else {
				v34 = v32
			}
			if v34 != 0 {
				v35 = *(*int32)(unsafe.Add(mBase, uint32(v34)+16))
				v36 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v35)+22)))
				v38 = *(*float32)(unsafe.Add(mBase, uint32(v35+v36)+8))
				v42 = base.F64_promote_f32(v38)
			} else {
				v42 = float64(0)
			}
			if v32 != 0 {
				v43 = *(*int32)(unsafe.Add(mBase, uint32(v11)+60))
				m.T0[v43].(func(*base.Module, int32))(m, v32)
				mBase = m.M
				v45 = m.ExcPending
				if v45 != 0 {
					return int32(0)
				} else {
					v46 = *(*int32)(unsafe.Add(mBase, uint32(v11)+24))
					v47 = v46
					if v47 == int32(0) {
						v69 = v42
						v72 = F_Float8GetDatum(m, base.F64_sub(float64(1), v69))
						mBase = m.M
						v73 = m.ExcPending
						if v73 != 0 {
							return int32(0)
						} else {
							m.G0 = v11 + int32(80)
							return v72
						}
					} else {
						v50 = *(*int32)(unsafe.Add(mBase, uint32(v11)+28))
						m.T0[v50].(func(*base.Module, int32))(m, v47)
						mBase = m.M
						v52 = m.ExcPending
						if v52 != 0 {
							return int32(0)
						} else {
							v69 = v42
							v72 = F_Float8GetDatum(m, base.F64_sub(float64(1), v69))
							mBase = m.M
							v73 = m.ExcPending
							if v73 != 0 {
								return int32(0)
							} else {
								m.G0 = v11 + int32(80)
								return v72
							}
						}
					}
				}
			} else {
				v47 = v31
				if v47 == int32(0) {
					v69 = v42
					v72 = F_Float8GetDatum(m, base.F64_sub(float64(1), v69))
					mBase = m.M
					v73 = m.ExcPending
					if v73 != 0 {
						return int32(0)
					} else {
						m.G0 = v11 + int32(80)
						return v72
					}
				} else {
					v50 = *(*int32)(unsafe.Add(mBase, uint32(v11)+28))
					m.T0[v50].(func(*base.Module, int32))(m, v47)
					mBase = m.M
					v52 = m.ExcPending
					if v52 != 0 {
						return int32(0)
					} else {
						v69 = v42
						v72 = F_Float8GetDatum(m, base.F64_sub(float64(1), v69))
						mBase = m.M
						v73 = m.ExcPending
						if v73 != 0 {
							return int32(0)
						} else {
							m.G0 = v11 + int32(80)
							return v72
						}
					}
				}
			}
		}
	} else {
		v53 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
		v54 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
		v55 = F_get_negator(m, v54)
		mBase = m.M
		v56 = m.ExcPending
		if v56 != 0 {
			return int32(0)
		} else {
			if v55 == int32(0) {
				v69 = float64(0.005)
				v72 = F_Float8GetDatum(m, base.F64_sub(float64(1), v69))
				mBase = m.M
				v73 = m.ExcPending
				if v73 != 0 {
					return int32(0)
				} else {
					m.G0 = v11 + int32(80)
					return v72
				}
			} else {
				v62 = F_DirectFunctionCall5Coll(m, int32(1508), v53, v15, v55, v14, base.I32_extend16_s(v16), v13)
				mBase = m.M
				v63 = m.ExcPending
				if v63 != 0 {
					return int32(0)
				} else {
					v64 = *(*float64)(unsafe.Add(mBase, uint32(v62)))
					v69 = v64
					v72 = F_Float8GetDatum(m, base.F64_sub(float64(1), v69))
					mBase = m.M
					v73 = m.ExcPending
					if v73 != 0 {
						return int32(0)
					} else {
						m.G0 = v11 + int32(80)
						return v72
					}
				}
			}
		}
	}
}
func F_neqsel(m *base.Module, l0 int32) int32 {
	var v3 float64
	_ = v3
	var v6 int32
	_ = v6
	var v7 int32
	_ = v7
	var v8 int32
	_ = v8
	v3 = F_eqsel_internal(m, l0, int32(1))
	v6 = m.ExcPending
	if v6 != 0 {
		return int32(0)
	} else {
		v7 = F_Float8GetDatum(m, v3)
		v8 = m.ExcPending
		if v8 != 0 {
			return int32(0)
		} else {
			return v7
		}
	}
}
func F_networksel(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v7 float64
	_ = v7
	var v11 int32
	_ = v11
	var v13 int32
	_ = v13
	var v15 int32
	_ = v15
	var v16 int32
	_ = v16
	var v17 int32
	_ = v17
	var v19 int32
	_ = v19
	var v31 int32
	_ = v31
	var v35 int32
	_ = v35
	var v40 int32
	_ = v40
	var v43 int32
	_ = v43
	var v50 int32
	_ = v50
	var v51 int32
	_ = v51
	var v58 float64
	_ = v58
	var v59 int32
	_ = v59
	var v60 int32
	_ = v60
	var v63 int32
	_ = v63
	var v64 int32
	_ = v64
	var v66 int32
	_ = v66
	var v71 float64
	_ = v71
	var v72 int32
	_ = v72
	var v75 int32
	_ = v75
	var v78 int32
	_ = v78
	var v80 int32
	_ = v80
	var v81 int32
	_ = v81
	var v88 float64
	_ = v88
	var v89 int32
	_ = v89
	var v90 int32
	_ = v90
	var v91 int32
	_ = v91
	var v93 float32
	_ = v93
	var v94 int32
	_ = v94
	var v95 int32
	_ = v95
	var v99 int32
	_ = v99
	var v106 int32
	_ = v106
	var v109 float64
	_ = v109
	var v110 int32
	_ = v110
	var v113 int32
	_ = v113
	var v117 int32
	_ = v117
	var v118 int32
	_ = v118
	var v119 int32
	_ = v119
	var v120 int32
	_ = v120
	var v123 int32
	_ = v123
	var v124 int32
	_ = v124
	var v125 float64
	_ = v125
	var v126 int32
	_ = v126
	var v130 int32
	_ = v130
	var v135 float64
	_ = v135
	var v136 float64
	_ = v136
	var v137 float64
	_ = v137
	var v140 float64
	_ = v140
	var v143 float64
	_ = v143
	var v151 float64
	_ = v151
	var v152 int32
	_ = v152
	var v155 int32
	_ = v155
	var v157 int32
	_ = v157
	var v161 float64
	_ = v161
	var v165 int32
	_ = v165
	var v166 int32
	_ = v166
	v7 = float64(0)
	v11 = m.G0
	v13 = v11 - int32(128)
	m.G0 = v13
	v15 = *(*int32)(unsafe.Add(mBase, uint32(l0)+44))
	v16 = *(*int32)(unsafe.Add(mBase, uint32(l0)+36))
	v17 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v19 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
	switch v19 - int32(931) {
	case 0:
		v43 = int32(2)
		v50 = F_get_restriction_variable(m, v17, v16, v15, v13+int32(96), v13+int32(92), v13+int32(91))
		mBase = m.M
		v51 = m.ExcPending
		if v51 != 0 {
			return int32(0)
		} else {
			if v50 == int32(0) {
				if v19 == int32(3552) {
					v58 = float64(0.01)
				} else {
					v58 = float64(0.005)
				}
				v161 = v58
				v165 = F_Float8GetDatum(m, v161)
				mBase = m.M
				v166 = m.ExcPending
				if v166 != 0 {
					return int32(0)
				} else {
					m.G0 = v13 + int32(128)
					return v165
				}
			} else {
				v59 = *(*int32)(unsafe.Add(mBase, uint32(v13)+92))
				v60 = *(*int32)(unsafe.Add(mBase, uint32(v59)))
				if v60 != int32(7) {
					v63 = *(*int32)(unsafe.Add(mBase, uint32(v13)+104))
					if v63 != 0 {
						v64 = *(*int32)(unsafe.Add(mBase, uint32(v13)+108))
						m.T0[v64].(func(*base.Module, int32))(m, v63)
						mBase = m.M
						v66 = m.ExcPending
						if v66 != 0 {
							return int32(0)
						} else {
							if v19 == int32(3552) {
								v71 = float64(0.01)
							} else {
								v71 = float64(0.005)
							}
							v161 = v71
							v165 = F_Float8GetDatum(m, v161)
							mBase = m.M
							v166 = m.ExcPending
							if v166 != 0 {
								return int32(0)
							} else {
								m.G0 = v13 + int32(128)
								return v165
							}
						}
					} else {
						if v19 == int32(3552) {
							v71 = float64(0.01)
						} else {
							v71 = float64(0.005)
						}
						v161 = v71
						v165 = F_Float8GetDatum(m, v161)
						mBase = m.M
						v166 = m.ExcPending
						if v166 != 0 {
							return int32(0)
						} else {
							m.G0 = v13 + int32(128)
							return v165
						}
					}
				} else {
					v72 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v59)+24)))
					if v72 == int32(1) {
						v75 = *(*int32)(unsafe.Add(mBase, uint32(v13)+104))
						if v75 == int32(0) {
							v161 = v7
							v165 = F_Float8GetDatum(m, v161)
							mBase = m.M
							v166 = m.ExcPending
							if v166 != 0 {
								return int32(0)
							} else {
								m.G0 = v13 + int32(128)
								return v165
							}
						} else {
							v78 = *(*int32)(unsafe.Add(mBase, uint32(v13)+108))
							m.T0[v78].(func(*base.Module, int32))(m, v75)
							mBase = m.M
							v80 = m.ExcPending
							if v80 != 0 {
								return int32(0)
							} else {
								v161 = v7
								v165 = F_Float8GetDatum(m, v161)
								mBase = m.M
								v166 = m.ExcPending
								if v166 != 0 {
									return int32(0)
								} else {
									m.G0 = v13 + int32(128)
									return v165
								}
							}
						}
					} else {
						v81 = *(*int32)(unsafe.Add(mBase, uint32(v13)+104))
						if v81 == int32(0) {
							if v19 == int32(3552) {
								v88 = float64(0.01)
							} else {
								v88 = float64(0.005)
							}
							v161 = v88
							v165 = F_Float8GetDatum(m, v161)
							mBase = m.M
							v166 = m.ExcPending
							if v166 != 0 {
								return int32(0)
							} else {
								m.G0 = v13 + int32(128)
								return v165
							}
						} else {
							v89 = *(*int32)(unsafe.Add(mBase, uint32(v59)+20))
							v90 = *(*int32)(unsafe.Add(mBase, uint32(v81)+16))
							v91 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v90)+22)))
							v93 = *(*float32)(unsafe.Add(mBase, uint32(v90+v91)+8))
							v94 = F_get_opcode(m, v19)
							mBase = m.M
							v95 = m.ExcPending
							if v95 != 0 {
								return int32(0)
							} else {
								F_fmgr_info(m, v94, v13+int32(12))
								mBase = m.M
								v99 = m.ExcPending
								if v99 != 0 {
									return int32(0)
								} else {
									v106 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v13)+91)))
									v109 = F_mcv_selectivity(m, v13+int32(96), v13+int32(12), int32(0), v89, v106, v13+int32(40))
									mBase = m.M
									v110 = m.ExcPending
									if v110 != 0 {
										return int32(0)
									} else {
										v113 = *(*int32)(unsafe.Add(mBase, uint32(v13)+104))
										v117 = F_get_attstatsslot(m, v13+int32(52), v113, int32(2), int32(0), int32(1))
										mBase = m.M
										v118 = m.ExcPending
										if v118 != 0 {
											return int32(0)
										} else {
											if v117 != 0 {
												v119 = *(*int32)(unsafe.Add(mBase, uint32(v13)+64))
												v120 = *(*int32)(unsafe.Add(mBase, uint32(v13)+68))
												v123 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v13)+91)))
												if v123 != 0 {
													v124 = v43
												} else {
													v124 = int32(0) - v43
												}
												v125 = F_inet_hist_value_sel(m, v119, v120, v89, v124)
												mBase = m.M
												v126 = m.ExcPending
												if v126 != 0 {
													return int32(0)
												} else {
													F_free_attstatsslot(m, v13+int32(52))
													mBase = m.M
													v130 = m.ExcPending
													if v130 != 0 {
														return int32(0)
													} else {
														v136 = v125
														v137 = float64(0)
														v140 = *(*float64)(unsafe.Add(mBase, uint32(v13)+40))
														v143 = base.F64_add(base.F64_mul(base.F64_sub(base.F64_sub(float64(1), base.F64_promote_f32(v93)), v140), v136), v109)
														if base.F64_lt(v143, v137) != 0 {
															v151 = v137
														} else {
															if base.F64_gt(v143, float64(1)) == int32(0) {
																v151 = v143
															} else {
																v151 = float64(1)
															}
														}
														v152 = *(*int32)(unsafe.Add(mBase, uint32(v13)+104))
														if v152 == int32(0) {
															v161 = v151
															v165 = F_Float8GetDatum(m, v161)
															mBase = m.M
															v166 = m.ExcPending
															if v166 != 0 {
																return int32(0)
															} else {
																m.G0 = v13 + int32(128)
																return v165
															}
														} else {
															v155 = *(*int32)(unsafe.Add(mBase, uint32(v13)+108))
															m.T0[v155].(func(*base.Module, int32))(m, v152)
															mBase = m.M
															v157 = m.ExcPending
															if v157 != 0 {
																return int32(0)
															} else {
																v161 = v151
																v165 = F_Float8GetDatum(m, v161)
																mBase = m.M
																v166 = m.ExcPending
																if v166 != 0 {
																	return int32(0)
																} else {
																	m.G0 = v13 + int32(128)
																	return v165
																}
															}
														}
													}
												}
											} else {
												if v19 == int32(3552) {
													v135 = float64(0.01)
												} else {
													v135 = float64(0.005)
												}
												v136 = v135
												v137 = float64(0)
												v140 = *(*float64)(unsafe.Add(mBase, uint32(v13)+40))
												v143 = base.F64_add(base.F64_mul(base.F64_sub(base.F64_sub(float64(1), base.F64_promote_f32(v93)), v140), v136), v109)
												if base.F64_lt(v143, v137) != 0 {
													v151 = v137
												} else {
													if base.F64_gt(v143, float64(1)) == int32(0) {
														v151 = v143
													} else {
														v151 = float64(1)
													}
												}
												v152 = *(*int32)(unsafe.Add(mBase, uint32(v13)+104))
												if v152 == int32(0) {
													v161 = v151
													v165 = F_Float8GetDatum(m, v161)
													mBase = m.M
													v166 = m.ExcPending
													if v166 != 0 {
														return int32(0)
													} else {
														m.G0 = v13 + int32(128)
														return v165
													}
												} else {
													v155 = *(*int32)(unsafe.Add(mBase, uint32(v13)+108))
													m.T0[v155].(func(*base.Module, int32))(m, v152)
													mBase = m.M
													v157 = m.ExcPending
													if v157 != 0 {
														return int32(0)
													} else {
														v161 = v151
														v165 = F_Float8GetDatum(m, v161)
														mBase = m.M
														v166 = m.ExcPending
														if v166 != 0 {
															return int32(0)
														} else {
															m.G0 = v13 + int32(128)
															return v165
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
			}
		}
	case 1:
		v43 = int32(1)
		v50 = F_get_restriction_variable(m, v17, v16, v15, v13+int32(96), v13+int32(92), v13+int32(91))
		mBase = m.M
		v51 = m.ExcPending
		if v51 != 0 {
			return int32(0)
		} else {
			if v50 == int32(0) {
				if v19 == int32(3552) {
					v58 = float64(0.01)
				} else {
					v58 = float64(0.005)
				}
				v161 = v58
				v165 = F_Float8GetDatum(m, v161)
				mBase = m.M
				v166 = m.ExcPending
				if v166 != 0 {
					return int32(0)
				} else {
					m.G0 = v13 + int32(128)
					return v165
				}
			} else {
				v59 = *(*int32)(unsafe.Add(mBase, uint32(v13)+92))
				v60 = *(*int32)(unsafe.Add(mBase, uint32(v59)))
				if v60 != int32(7) {
					v63 = *(*int32)(unsafe.Add(mBase, uint32(v13)+104))
					if v63 != 0 {
						v64 = *(*int32)(unsafe.Add(mBase, uint32(v13)+108))
						m.T0[v64].(func(*base.Module, int32))(m, v63)
						mBase = m.M
						v66 = m.ExcPending
						if v66 != 0 {
							return int32(0)
						} else {
							if v19 == int32(3552) {
								v71 = float64(0.01)
							} else {
								v71 = float64(0.005)
							}
							v161 = v71
							v165 = F_Float8GetDatum(m, v161)
							mBase = m.M
							v166 = m.ExcPending
							if v166 != 0 {
								return int32(0)
							} else {
								m.G0 = v13 + int32(128)
								return v165
							}
						}
					} else {
						if v19 == int32(3552) {
							v71 = float64(0.01)
						} else {
							v71 = float64(0.005)
						}
						v161 = v71
						v165 = F_Float8GetDatum(m, v161)
						mBase = m.M
						v166 = m.ExcPending
						if v166 != 0 {
							return int32(0)
						} else {
							m.G0 = v13 + int32(128)
							return v165
						}
					}
				} else {
					v72 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v59)+24)))
					if v72 == int32(1) {
						v75 = *(*int32)(unsafe.Add(mBase, uint32(v13)+104))
						if v75 == int32(0) {
							v161 = v7
							v165 = F_Float8GetDatum(m, v161)
							mBase = m.M
							v166 = m.ExcPending
							if v166 != 0 {
								return int32(0)
							} else {
								m.G0 = v13 + int32(128)
								return v165
							}
						} else {
							v78 = *(*int32)(unsafe.Add(mBase, uint32(v13)+108))
							m.T0[v78].(func(*base.Module, int32))(m, v75)
							mBase = m.M
							v80 = m.ExcPending
							if v80 != 0 {
								return int32(0)
							} else {
								v161 = v7
								v165 = F_Float8GetDatum(m, v161)
								mBase = m.M
								v166 = m.ExcPending
								if v166 != 0 {
									return int32(0)
								} else {
									m.G0 = v13 + int32(128)
									return v165
								}
							}
						}
					} else {
						v81 = *(*int32)(unsafe.Add(mBase, uint32(v13)+104))
						if v81 == int32(0) {
							if v19 == int32(3552) {
								v88 = float64(0.01)
							} else {
								v88 = float64(0.005)
							}
							v161 = v88
							v165 = F_Float8GetDatum(m, v161)
							mBase = m.M
							v166 = m.ExcPending
							if v166 != 0 {
								return int32(0)
							} else {
								m.G0 = v13 + int32(128)
								return v165
							}
						} else {
							v89 = *(*int32)(unsafe.Add(mBase, uint32(v59)+20))
							v90 = *(*int32)(unsafe.Add(mBase, uint32(v81)+16))
							v91 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v90)+22)))
							v93 = *(*float32)(unsafe.Add(mBase, uint32(v90+v91)+8))
							v94 = F_get_opcode(m, v19)
							mBase = m.M
							v95 = m.ExcPending
							if v95 != 0 {
								return int32(0)
							} else {
								F_fmgr_info(m, v94, v13+int32(12))
								mBase = m.M
								v99 = m.ExcPending
								if v99 != 0 {
									return int32(0)
								} else {
									v106 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v13)+91)))
									v109 = F_mcv_selectivity(m, v13+int32(96), v13+int32(12), int32(0), v89, v106, v13+int32(40))
									mBase = m.M
									v110 = m.ExcPending
									if v110 != 0 {
										return int32(0)
									} else {
										v113 = *(*int32)(unsafe.Add(mBase, uint32(v13)+104))
										v117 = F_get_attstatsslot(m, v13+int32(52), v113, int32(2), int32(0), int32(1))
										mBase = m.M
										v118 = m.ExcPending
										if v118 != 0 {
											return int32(0)
										} else {
											if v117 != 0 {
												v119 = *(*int32)(unsafe.Add(mBase, uint32(v13)+64))
												v120 = *(*int32)(unsafe.Add(mBase, uint32(v13)+68))
												v123 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v13)+91)))
												if v123 != 0 {
													v124 = v43
												} else {
													v124 = int32(0) - v43
												}
												v125 = F_inet_hist_value_sel(m, v119, v120, v89, v124)
												mBase = m.M
												v126 = m.ExcPending
												if v126 != 0 {
													return int32(0)
												} else {
													F_free_attstatsslot(m, v13+int32(52))
													mBase = m.M
													v130 = m.ExcPending
													if v130 != 0 {
														return int32(0)
													} else {
														v136 = v125
														v137 = float64(0)
														v140 = *(*float64)(unsafe.Add(mBase, uint32(v13)+40))
														v143 = base.F64_add(base.F64_mul(base.F64_sub(base.F64_sub(float64(1), base.F64_promote_f32(v93)), v140), v136), v109)
														if base.F64_lt(v143, v137) != 0 {
															v151 = v137
														} else {
															if base.F64_gt(v143, float64(1)) == int32(0) {
																v151 = v143
															} else {
																v151 = float64(1)
															}
														}
														v152 = *(*int32)(unsafe.Add(mBase, uint32(v13)+104))
														if v152 == int32(0) {
															v161 = v151
															v165 = F_Float8GetDatum(m, v161)
															mBase = m.M
															v166 = m.ExcPending
															if v166 != 0 {
																return int32(0)
															} else {
																m.G0 = v13 + int32(128)
																return v165
															}
														} else {
															v155 = *(*int32)(unsafe.Add(mBase, uint32(v13)+108))
															m.T0[v155].(func(*base.Module, int32))(m, v152)
															mBase = m.M
															v157 = m.ExcPending
															if v157 != 0 {
																return int32(0)
															} else {
																v161 = v151
																v165 = F_Float8GetDatum(m, v161)
																mBase = m.M
																v166 = m.ExcPending
																if v166 != 0 {
																	return int32(0)
																} else {
																	m.G0 = v13 + int32(128)
																	return v165
																}
															}
														}
													}
												}
											} else {
												if v19 == int32(3552) {
													v135 = float64(0.01)
												} else {
													v135 = float64(0.005)
												}
												v136 = v135
												v137 = float64(0)
												v140 = *(*float64)(unsafe.Add(mBase, uint32(v13)+40))
												v143 = base.F64_add(base.F64_mul(base.F64_sub(base.F64_sub(float64(1), base.F64_promote_f32(v93)), v140), v136), v109)
												if base.F64_lt(v143, v137) != 0 {
													v151 = v137
												} else {
													if base.F64_gt(v143, float64(1)) == int32(0) {
														v151 = v143
													} else {
														v151 = float64(1)
													}
												}
												v152 = *(*int32)(unsafe.Add(mBase, uint32(v13)+104))
												if v152 == int32(0) {
													v161 = v151
													v165 = F_Float8GetDatum(m, v161)
													mBase = m.M
													v166 = m.ExcPending
													if v166 != 0 {
														return int32(0)
													} else {
														m.G0 = v13 + int32(128)
														return v165
													}
												} else {
													v155 = *(*int32)(unsafe.Add(mBase, uint32(v13)+108))
													m.T0[v155].(func(*base.Module, int32))(m, v152)
													mBase = m.M
													v157 = m.ExcPending
													if v157 != 0 {
														return int32(0)
													} else {
														v161 = v151
														v165 = F_Float8GetDatum(m, v161)
														mBase = m.M
														v166 = m.ExcPending
														if v166 != 0 {
															return int32(0)
														} else {
															m.G0 = v13 + int32(128)
															return v165
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
			}
		}
	case 2:
		v43 = int32(-2)
		v50 = F_get_restriction_variable(m, v17, v16, v15, v13+int32(96), v13+int32(92), v13+int32(91))
		mBase = m.M
		v51 = m.ExcPending
		if v51 != 0 {
			return int32(0)
		} else {
			if v50 == int32(0) {
				if v19 == int32(3552) {
					v58 = float64(0.01)
				} else {
					v58 = float64(0.005)
				}
				v161 = v58
				v165 = F_Float8GetDatum(m, v161)
				mBase = m.M
				v166 = m.ExcPending
				if v166 != 0 {
					return int32(0)
				} else {
					m.G0 = v13 + int32(128)
					return v165
				}
			} else {
				v59 = *(*int32)(unsafe.Add(mBase, uint32(v13)+92))
				v60 = *(*int32)(unsafe.Add(mBase, uint32(v59)))
				if v60 != int32(7) {
					v63 = *(*int32)(unsafe.Add(mBase, uint32(v13)+104))
					if v63 != 0 {
						v64 = *(*int32)(unsafe.Add(mBase, uint32(v13)+108))
						m.T0[v64].(func(*base.Module, int32))(m, v63)
						mBase = m.M
						v66 = m.ExcPending
						if v66 != 0 {
							return int32(0)
						} else {
							if v19 == int32(3552) {
								v71 = float64(0.01)
							} else {
								v71 = float64(0.005)
							}
							v161 = v71
							v165 = F_Float8GetDatum(m, v161)
							mBase = m.M
							v166 = m.ExcPending
							if v166 != 0 {
								return int32(0)
							} else {
								m.G0 = v13 + int32(128)
								return v165
							}
						}
					} else {
						if v19 == int32(3552) {
							v71 = float64(0.01)
						} else {
							v71 = float64(0.005)
						}
						v161 = v71
						v165 = F_Float8GetDatum(m, v161)
						mBase = m.M
						v166 = m.ExcPending
						if v166 != 0 {
							return int32(0)
						} else {
							m.G0 = v13 + int32(128)
							return v165
						}
					}
				} else {
					v72 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v59)+24)))
					if v72 == int32(1) {
						v75 = *(*int32)(unsafe.Add(mBase, uint32(v13)+104))
						if v75 == int32(0) {
							v161 = v7
							v165 = F_Float8GetDatum(m, v161)
							mBase = m.M
							v166 = m.ExcPending
							if v166 != 0 {
								return int32(0)
							} else {
								m.G0 = v13 + int32(128)
								return v165
							}
						} else {
							v78 = *(*int32)(unsafe.Add(mBase, uint32(v13)+108))
							m.T0[v78].(func(*base.Module, int32))(m, v75)
							mBase = m.M
							v80 = m.ExcPending
							if v80 != 0 {
								return int32(0)
							} else {
								v161 = v7
								v165 = F_Float8GetDatum(m, v161)
								mBase = m.M
								v166 = m.ExcPending
								if v166 != 0 {
									return int32(0)
								} else {
									m.G0 = v13 + int32(128)
									return v165
								}
							}
						}
					} else {
						v81 = *(*int32)(unsafe.Add(mBase, uint32(v13)+104))
						if v81 == int32(0) {
							if v19 == int32(3552) {
								v88 = float64(0.01)
							} else {
								v88 = float64(0.005)
							}
							v161 = v88
							v165 = F_Float8GetDatum(m, v161)
							mBase = m.M
							v166 = m.ExcPending
							if v166 != 0 {
								return int32(0)
							} else {
								m.G0 = v13 + int32(128)
								return v165
							}
						} else {
							v89 = *(*int32)(unsafe.Add(mBase, uint32(v59)+20))
							v90 = *(*int32)(unsafe.Add(mBase, uint32(v81)+16))
							v91 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v90)+22)))
							v93 = *(*float32)(unsafe.Add(mBase, uint32(v90+v91)+8))
							v94 = F_get_opcode(m, v19)
							mBase = m.M
							v95 = m.ExcPending
							if v95 != 0 {
								return int32(0)
							} else {
								F_fmgr_info(m, v94, v13+int32(12))
								mBase = m.M
								v99 = m.ExcPending
								if v99 != 0 {
									return int32(0)
								} else {
									v106 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v13)+91)))
									v109 = F_mcv_selectivity(m, v13+int32(96), v13+int32(12), int32(0), v89, v106, v13+int32(40))
									mBase = m.M
									v110 = m.ExcPending
									if v110 != 0 {
										return int32(0)
									} else {
										v113 = *(*int32)(unsafe.Add(mBase, uint32(v13)+104))
										v117 = F_get_attstatsslot(m, v13+int32(52), v113, int32(2), int32(0), int32(1))
										mBase = m.M
										v118 = m.ExcPending
										if v118 != 0 {
											return int32(0)
										} else {
											if v117 != 0 {
												v119 = *(*int32)(unsafe.Add(mBase, uint32(v13)+64))
												v120 = *(*int32)(unsafe.Add(mBase, uint32(v13)+68))
												v123 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v13)+91)))
												if v123 != 0 {
													v124 = v43
												} else {
													v124 = int32(0) - v43
												}
												v125 = F_inet_hist_value_sel(m, v119, v120, v89, v124)
												mBase = m.M
												v126 = m.ExcPending
												if v126 != 0 {
													return int32(0)
												} else {
													F_free_attstatsslot(m, v13+int32(52))
													mBase = m.M
													v130 = m.ExcPending
													if v130 != 0 {
														return int32(0)
													} else {
														v136 = v125
														v137 = float64(0)
														v140 = *(*float64)(unsafe.Add(mBase, uint32(v13)+40))
														v143 = base.F64_add(base.F64_mul(base.F64_sub(base.F64_sub(float64(1), base.F64_promote_f32(v93)), v140), v136), v109)
														if base.F64_lt(v143, v137) != 0 {
															v151 = v137
														} else {
															if base.F64_gt(v143, float64(1)) == int32(0) {
																v151 = v143
															} else {
																v151 = float64(1)
															}
														}
														v152 = *(*int32)(unsafe.Add(mBase, uint32(v13)+104))
														if v152 == int32(0) {
															v161 = v151
															v165 = F_Float8GetDatum(m, v161)
															mBase = m.M
															v166 = m.ExcPending
															if v166 != 0 {
																return int32(0)
															} else {
																m.G0 = v13 + int32(128)
																return v165
															}
														} else {
															v155 = *(*int32)(unsafe.Add(mBase, uint32(v13)+108))
															m.T0[v155].(func(*base.Module, int32))(m, v152)
															mBase = m.M
															v157 = m.ExcPending
															if v157 != 0 {
																return int32(0)
															} else {
																v161 = v151
																v165 = F_Float8GetDatum(m, v161)
																mBase = m.M
																v166 = m.ExcPending
																if v166 != 0 {
																	return int32(0)
																} else {
																	m.G0 = v13 + int32(128)
																	return v165
																}
															}
														}
													}
												}
											} else {
												if v19 == int32(3552) {
													v135 = float64(0.01)
												} else {
													v135 = float64(0.005)
												}
												v136 = v135
												v137 = float64(0)
												v140 = *(*float64)(unsafe.Add(mBase, uint32(v13)+40))
												v143 = base.F64_add(base.F64_mul(base.F64_sub(base.F64_sub(float64(1), base.F64_promote_f32(v93)), v140), v136), v109)
												if base.F64_lt(v143, v137) != 0 {
													v151 = v137
												} else {
													if base.F64_gt(v143, float64(1)) == int32(0) {
														v151 = v143
													} else {
														v151 = float64(1)
													}
												}
												v152 = *(*int32)(unsafe.Add(mBase, uint32(v13)+104))
												if v152 == int32(0) {
													v161 = v151
													v165 = F_Float8GetDatum(m, v161)
													mBase = m.M
													v166 = m.ExcPending
													if v166 != 0 {
														return int32(0)
													} else {
														m.G0 = v13 + int32(128)
														return v165
													}
												} else {
													v155 = *(*int32)(unsafe.Add(mBase, uint32(v13)+108))
													m.T0[v155].(func(*base.Module, int32))(m, v152)
													mBase = m.M
													v157 = m.ExcPending
													if v157 != 0 {
														return int32(0)
													} else {
														v161 = v151
														v165 = F_Float8GetDatum(m, v161)
														mBase = m.M
														v166 = m.ExcPending
														if v166 != 0 {
															return int32(0)
														} else {
															m.G0 = v13 + int32(128)
															return v165
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
			}
		}
	case 3:
		v43 = int32(-1)
		v50 = F_get_restriction_variable(m, v17, v16, v15, v13+int32(96), v13+int32(92), v13+int32(91))
		mBase = m.M
		v51 = m.ExcPending
		if v51 != 0 {
			return int32(0)
		} else {
			if v50 == int32(0) {
				if v19 == int32(3552) {
					v58 = float64(0.01)
				} else {
					v58 = float64(0.005)
				}
				v161 = v58
				v165 = F_Float8GetDatum(m, v161)
				mBase = m.M
				v166 = m.ExcPending
				if v166 != 0 {
					return int32(0)
				} else {
					m.G0 = v13 + int32(128)
					return v165
				}
			} else {
				v59 = *(*int32)(unsafe.Add(mBase, uint32(v13)+92))
				v60 = *(*int32)(unsafe.Add(mBase, uint32(v59)))
				if v60 != int32(7) {
					v63 = *(*int32)(unsafe.Add(mBase, uint32(v13)+104))
					if v63 != 0 {
						v64 = *(*int32)(unsafe.Add(mBase, uint32(v13)+108))
						m.T0[v64].(func(*base.Module, int32))(m, v63)
						mBase = m.M
						v66 = m.ExcPending
						if v66 != 0 {
							return int32(0)
						} else {
							if v19 == int32(3552) {
								v71 = float64(0.01)
							} else {
								v71 = float64(0.005)
							}
							v161 = v71
							v165 = F_Float8GetDatum(m, v161)
							mBase = m.M
							v166 = m.ExcPending
							if v166 != 0 {
								return int32(0)
							} else {
								m.G0 = v13 + int32(128)
								return v165
							}
						}
					} else {
						if v19 == int32(3552) {
							v71 = float64(0.01)
						} else {
							v71 = float64(0.005)
						}
						v161 = v71
						v165 = F_Float8GetDatum(m, v161)
						mBase = m.M
						v166 = m.ExcPending
						if v166 != 0 {
							return int32(0)
						} else {
							m.G0 = v13 + int32(128)
							return v165
						}
					}
				} else {
					v72 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v59)+24)))
					if v72 == int32(1) {
						v75 = *(*int32)(unsafe.Add(mBase, uint32(v13)+104))
						if v75 == int32(0) {
							v161 = v7
							v165 = F_Float8GetDatum(m, v161)
							mBase = m.M
							v166 = m.ExcPending
							if v166 != 0 {
								return int32(0)
							} else {
								m.G0 = v13 + int32(128)
								return v165
							}
						} else {
							v78 = *(*int32)(unsafe.Add(mBase, uint32(v13)+108))
							m.T0[v78].(func(*base.Module, int32))(m, v75)
							mBase = m.M
							v80 = m.ExcPending
							if v80 != 0 {
								return int32(0)
							} else {
								v161 = v7
								v165 = F_Float8GetDatum(m, v161)
								mBase = m.M
								v166 = m.ExcPending
								if v166 != 0 {
									return int32(0)
								} else {
									m.G0 = v13 + int32(128)
									return v165
								}
							}
						}
					} else {
						v81 = *(*int32)(unsafe.Add(mBase, uint32(v13)+104))
						if v81 == int32(0) {
							if v19 == int32(3552) {
								v88 = float64(0.01)
							} else {
								v88 = float64(0.005)
							}
							v161 = v88
							v165 = F_Float8GetDatum(m, v161)
							mBase = m.M
							v166 = m.ExcPending
							if v166 != 0 {
								return int32(0)
							} else {
								m.G0 = v13 + int32(128)
								return v165
							}
						} else {
							v89 = *(*int32)(unsafe.Add(mBase, uint32(v59)+20))
							v90 = *(*int32)(unsafe.Add(mBase, uint32(v81)+16))
							v91 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v90)+22)))
							v93 = *(*float32)(unsafe.Add(mBase, uint32(v90+v91)+8))
							v94 = F_get_opcode(m, v19)
							mBase = m.M
							v95 = m.ExcPending
							if v95 != 0 {
								return int32(0)
							} else {
								F_fmgr_info(m, v94, v13+int32(12))
								mBase = m.M
								v99 = m.ExcPending
								if v99 != 0 {
									return int32(0)
								} else {
									v106 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v13)+91)))
									v109 = F_mcv_selectivity(m, v13+int32(96), v13+int32(12), int32(0), v89, v106, v13+int32(40))
									mBase = m.M
									v110 = m.ExcPending
									if v110 != 0 {
										return int32(0)
									} else {
										v113 = *(*int32)(unsafe.Add(mBase, uint32(v13)+104))
										v117 = F_get_attstatsslot(m, v13+int32(52), v113, int32(2), int32(0), int32(1))
										mBase = m.M
										v118 = m.ExcPending
										if v118 != 0 {
											return int32(0)
										} else {
											if v117 != 0 {
												v119 = *(*int32)(unsafe.Add(mBase, uint32(v13)+64))
												v120 = *(*int32)(unsafe.Add(mBase, uint32(v13)+68))
												v123 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v13)+91)))
												if v123 != 0 {
													v124 = v43
												} else {
													v124 = int32(0) - v43
												}
												v125 = F_inet_hist_value_sel(m, v119, v120, v89, v124)
												mBase = m.M
												v126 = m.ExcPending
												if v126 != 0 {
													return int32(0)
												} else {
													F_free_attstatsslot(m, v13+int32(52))
													mBase = m.M
													v130 = m.ExcPending
													if v130 != 0 {
														return int32(0)
													} else {
														v136 = v125
														v137 = float64(0)
														v140 = *(*float64)(unsafe.Add(mBase, uint32(v13)+40))
														v143 = base.F64_add(base.F64_mul(base.F64_sub(base.F64_sub(float64(1), base.F64_promote_f32(v93)), v140), v136), v109)
														if base.F64_lt(v143, v137) != 0 {
															v151 = v137
														} else {
															if base.F64_gt(v143, float64(1)) == int32(0) {
																v151 = v143
															} else {
																v151 = float64(1)
															}
														}
														v152 = *(*int32)(unsafe.Add(mBase, uint32(v13)+104))
														if v152 == int32(0) {
															v161 = v151
															v165 = F_Float8GetDatum(m, v161)
															mBase = m.M
															v166 = m.ExcPending
															if v166 != 0 {
																return int32(0)
															} else {
																m.G0 = v13 + int32(128)
																return v165
															}
														} else {
															v155 = *(*int32)(unsafe.Add(mBase, uint32(v13)+108))
															m.T0[v155].(func(*base.Module, int32))(m, v152)
															mBase = m.M
															v157 = m.ExcPending
															if v157 != 0 {
																return int32(0)
															} else {
																v161 = v151
																v165 = F_Float8GetDatum(m, v161)
																mBase = m.M
																v166 = m.ExcPending
																if v166 != 0 {
																	return int32(0)
																} else {
																	m.G0 = v13 + int32(128)
																	return v165
																}
															}
														}
													}
												}
											} else {
												if v19 == int32(3552) {
													v135 = float64(0.01)
												} else {
													v135 = float64(0.005)
												}
												v136 = v135
												v137 = float64(0)
												v140 = *(*float64)(unsafe.Add(mBase, uint32(v13)+40))
												v143 = base.F64_add(base.F64_mul(base.F64_sub(base.F64_sub(float64(1), base.F64_promote_f32(v93)), v140), v136), v109)
												if base.F64_lt(v143, v137) != 0 {
													v151 = v137
												} else {
													if base.F64_gt(v143, float64(1)) == int32(0) {
														v151 = v143
													} else {
														v151 = float64(1)
													}
												}
												v152 = *(*int32)(unsafe.Add(mBase, uint32(v13)+104))
												if v152 == int32(0) {
													v161 = v151
													v165 = F_Float8GetDatum(m, v161)
													mBase = m.M
													v166 = m.ExcPending
													if v166 != 0 {
														return int32(0)
													} else {
														m.G0 = v13 + int32(128)
														return v165
													}
												} else {
													v155 = *(*int32)(unsafe.Add(mBase, uint32(v13)+108))
													m.T0[v155].(func(*base.Module, int32))(m, v152)
													mBase = m.M
													v157 = m.ExcPending
													if v157 != 0 {
														return int32(0)
													} else {
														v161 = v151
														v165 = F_Float8GetDatum(m, v161)
														mBase = m.M
														v166 = m.ExcPending
														if v166 != 0 {
															return int32(0)
														} else {
															m.G0 = v13 + int32(128)
															return v165
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
			}
		}
	default:
		if v19 == int32(3552) {
			v43 = int32(0)
			v50 = F_get_restriction_variable(m, v17, v16, v15, v13+int32(96), v13+int32(92), v13+int32(91))
			mBase = m.M
			v51 = m.ExcPending
			if v51 != 0 {
				return int32(0)
			} else {
				if v50 == int32(0) {
					if v19 == int32(3552) {
						v58 = float64(0.01)
					} else {
						v58 = float64(0.005)
					}
					v161 = v58
					v165 = F_Float8GetDatum(m, v161)
					mBase = m.M
					v166 = m.ExcPending
					if v166 != 0 {
						return int32(0)
					} else {
						m.G0 = v13 + int32(128)
						return v165
					}
				} else {
					v59 = *(*int32)(unsafe.Add(mBase, uint32(v13)+92))
					v60 = *(*int32)(unsafe.Add(mBase, uint32(v59)))
					if v60 != int32(7) {
						v63 = *(*int32)(unsafe.Add(mBase, uint32(v13)+104))
						if v63 != 0 {
							v64 = *(*int32)(unsafe.Add(mBase, uint32(v13)+108))
							m.T0[v64].(func(*base.Module, int32))(m, v63)
							mBase = m.M
							v66 = m.ExcPending
							if v66 != 0 {
								return int32(0)
							} else {
								if v19 == int32(3552) {
									v71 = float64(0.01)
								} else {
									v71 = float64(0.005)
								}
								v161 = v71
								v165 = F_Float8GetDatum(m, v161)
								mBase = m.M
								v166 = m.ExcPending
								if v166 != 0 {
									return int32(0)
								} else {
									m.G0 = v13 + int32(128)
									return v165
								}
							}
						} else {
							if v19 == int32(3552) {
								v71 = float64(0.01)
							} else {
								v71 = float64(0.005)
							}
							v161 = v71
							v165 = F_Float8GetDatum(m, v161)
							mBase = m.M
							v166 = m.ExcPending
							if v166 != 0 {
								return int32(0)
							} else {
								m.G0 = v13 + int32(128)
								return v165
							}
						}
					} else {
						v72 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v59)+24)))
						if v72 == int32(1) {
							v75 = *(*int32)(unsafe.Add(mBase, uint32(v13)+104))
							if v75 == int32(0) {
								v161 = v7
								v165 = F_Float8GetDatum(m, v161)
								mBase = m.M
								v166 = m.ExcPending
								if v166 != 0 {
									return int32(0)
								} else {
									m.G0 = v13 + int32(128)
									return v165
								}
							} else {
								v78 = *(*int32)(unsafe.Add(mBase, uint32(v13)+108))
								m.T0[v78].(func(*base.Module, int32))(m, v75)
								mBase = m.M
								v80 = m.ExcPending
								if v80 != 0 {
									return int32(0)
								} else {
									v161 = v7
									v165 = F_Float8GetDatum(m, v161)
									mBase = m.M
									v166 = m.ExcPending
									if v166 != 0 {
										return int32(0)
									} else {
										m.G0 = v13 + int32(128)
										return v165
									}
								}
							}
						} else {
							v81 = *(*int32)(unsafe.Add(mBase, uint32(v13)+104))
							if v81 == int32(0) {
								if v19 == int32(3552) {
									v88 = float64(0.01)
								} else {
									v88 = float64(0.005)
								}
								v161 = v88
								v165 = F_Float8GetDatum(m, v161)
								mBase = m.M
								v166 = m.ExcPending
								if v166 != 0 {
									return int32(0)
								} else {
									m.G0 = v13 + int32(128)
									return v165
								}
							} else {
								v89 = *(*int32)(unsafe.Add(mBase, uint32(v59)+20))
								v90 = *(*int32)(unsafe.Add(mBase, uint32(v81)+16))
								v91 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v90)+22)))
								v93 = *(*float32)(unsafe.Add(mBase, uint32(v90+v91)+8))
								v94 = F_get_opcode(m, v19)
								mBase = m.M
								v95 = m.ExcPending
								if v95 != 0 {
									return int32(0)
								} else {
									F_fmgr_info(m, v94, v13+int32(12))
									mBase = m.M
									v99 = m.ExcPending
									if v99 != 0 {
										return int32(0)
									} else {
										v106 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v13)+91)))
										v109 = F_mcv_selectivity(m, v13+int32(96), v13+int32(12), int32(0), v89, v106, v13+int32(40))
										mBase = m.M
										v110 = m.ExcPending
										if v110 != 0 {
											return int32(0)
										} else {
											v113 = *(*int32)(unsafe.Add(mBase, uint32(v13)+104))
											v117 = F_get_attstatsslot(m, v13+int32(52), v113, int32(2), int32(0), int32(1))
											mBase = m.M
											v118 = m.ExcPending
											if v118 != 0 {
												return int32(0)
											} else {
												if v117 != 0 {
													v119 = *(*int32)(unsafe.Add(mBase, uint32(v13)+64))
													v120 = *(*int32)(unsafe.Add(mBase, uint32(v13)+68))
													v123 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v13)+91)))
													if v123 != 0 {
														v124 = v43
													} else {
														v124 = int32(0) - v43
													}
													v125 = F_inet_hist_value_sel(m, v119, v120, v89, v124)
													mBase = m.M
													v126 = m.ExcPending
													if v126 != 0 {
														return int32(0)
													} else {
														F_free_attstatsslot(m, v13+int32(52))
														mBase = m.M
														v130 = m.ExcPending
														if v130 != 0 {
															return int32(0)
														} else {
															v136 = v125
															v137 = float64(0)
															v140 = *(*float64)(unsafe.Add(mBase, uint32(v13)+40))
															v143 = base.F64_add(base.F64_mul(base.F64_sub(base.F64_sub(float64(1), base.F64_promote_f32(v93)), v140), v136), v109)
															if base.F64_lt(v143, v137) != 0 {
																v151 = v137
															} else {
																if base.F64_gt(v143, float64(1)) == int32(0) {
																	v151 = v143
																} else {
																	v151 = float64(1)
																}
															}
															v152 = *(*int32)(unsafe.Add(mBase, uint32(v13)+104))
															if v152 == int32(0) {
																v161 = v151
																v165 = F_Float8GetDatum(m, v161)
																mBase = m.M
																v166 = m.ExcPending
																if v166 != 0 {
																	return int32(0)
																} else {
																	m.G0 = v13 + int32(128)
																	return v165
																}
															} else {
																v155 = *(*int32)(unsafe.Add(mBase, uint32(v13)+108))
																m.T0[v155].(func(*base.Module, int32))(m, v152)
																mBase = m.M
																v157 = m.ExcPending
																if v157 != 0 {
																	return int32(0)
																} else {
																	v161 = v151
																	v165 = F_Float8GetDatum(m, v161)
																	mBase = m.M
																	v166 = m.ExcPending
																	if v166 != 0 {
																		return int32(0)
																	} else {
																		m.G0 = v13 + int32(128)
																		return v165
																	}
																}
															}
														}
													}
												} else {
													if v19 == int32(3552) {
														v135 = float64(0.01)
													} else {
														v135 = float64(0.005)
													}
													v136 = v135
													v137 = float64(0)
													v140 = *(*float64)(unsafe.Add(mBase, uint32(v13)+40))
													v143 = base.F64_add(base.F64_mul(base.F64_sub(base.F64_sub(float64(1), base.F64_promote_f32(v93)), v140), v136), v109)
													if base.F64_lt(v143, v137) != 0 {
														v151 = v137
													} else {
														if base.F64_gt(v143, float64(1)) == int32(0) {
															v151 = v143
														} else {
															v151 = float64(1)
														}
													}
													v152 = *(*int32)(unsafe.Add(mBase, uint32(v13)+104))
													if v152 == int32(0) {
														v161 = v151
														v165 = F_Float8GetDatum(m, v161)
														mBase = m.M
														v166 = m.ExcPending
														if v166 != 0 {
															return int32(0)
														} else {
															m.G0 = v13 + int32(128)
															return v165
														}
													} else {
														v155 = *(*int32)(unsafe.Add(mBase, uint32(v13)+108))
														m.T0[v155].(func(*base.Module, int32))(m, v152)
														mBase = m.M
														v157 = m.ExcPending
														if v157 != 0 {
															return int32(0)
														} else {
															v161 = v151
															v165 = F_Float8GetDatum(m, v161)
															mBase = m.M
															v166 = m.ExcPending
															if v166 != 0 {
																return int32(0)
															} else {
																m.G0 = v13 + int32(128)
																return v165
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
				}
			}
		} else {
			F_errstart_cold(m, int32(21), int32(0))
			mBase = m.M
			v31 = m.ExcPending
			if v31 != 0 {
				return int32(0)
			} else {
				*(*int32)(unsafe.Add(mBase, uint32(v13))) = v19
				F_errmsg_internal(m, int32(_a_F_networksel_0), v13)
				mBase = m.M
				v35 = m.ExcPending
				if v35 != 0 {
					return int32(0)
				} else {
					F_errfinish(m, int32(_a_F_networksel_1), int32(870), int32(_a_F_networksel_2))
					mBase = m.M
					v40 = m.ExcPending
					if v40 != 0 {
						return int32(0)
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
func F_newarc(m *base.Module, l0 int32, l1 int32, l2 int32) {
	mBase := m.M
	_ = mBase
	var v6 int32
	_ = v6
	var v8 int32
	_ = v8
	var v9 int32
	_ = v9
	var v10 int32
	_ = v10
	var v12 int32
	_ = v12
	var v18 int32
	_ = v18
	var v19 int32
	_ = v19
	var v21 int32
	_ = v21
	var v22 int32
	_ = v22
	var v25 int32
	_ = v25
	var v26 int32
	_ = v26
	var v32 int32
	_ = v32
	var v33 int32
	_ = v33
	var v35 int32
	_ = v35
	var v36 int32
	_ = v36
	var v39 int32
	_ = v39
	var v47 int32
	_ = v47
	v6 = *(*int32)(unsafe.Add(mBase, _c_F_newarc[0]))
	if v6 != 0 {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	F_ProcessInterrupts(m)
	mBase = m.M
	v8 = m.ExcPending
	if v8 != 0 {
		goto L4
	} else {
		goto L5
	}
L2:
	;
	goto L3
L3:
	;
	v9 = *(*int32)(unsafe.Add(mBase, uint32(l1)+12))
	v10 = *(*int32)(unsafe.Add(mBase, uint32(l2)+8))
	if v9 <= v10 {
		goto L8
	} else {
		goto L9
	}
L4:
	;
	return
L5:
	;
	goto L3
L6:
	;
	return
L7:
	;
	F_createarc(m, l0, int32(110), int32(0), l1, l2)
	mBase = m.M
	v47 = m.ExcPending
	if v47 != 0 {
		goto L4
	} else {
		goto L27
	}
L8:
	;
	v12 = *(*int32)(unsafe.Add(mBase, uint32(l1)+20))
	if v12 == int32(0) {
		goto L7
	} else {
		goto L11
	}
L9:
	;
	goto L10
L10:
	;
	v26 = *(*int32)(unsafe.Add(mBase, uint32(l2)+16))
	if v26 == int32(0) {
		goto L7
	} else {
		goto L19
	}
L11:
	;
	v18 = v12
	goto L12
L12:
	;
	v19 = *(*int32)(unsafe.Add(mBase, uint32(v18)+12))
	if v19 != l2 {
		goto L14
	} else {
		goto L15
	}
L13:
	;
	goto L7
L14:
	;
	v25 = *(*int32)(unsafe.Add(mBase, uint32(v18)+16))
	if v25 != 0 {
		v18 = v25
		goto L12
	} else {
		goto L18
	}
L15:
	;
	v21 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v18)+4)))
	if v21 != 0 {
		goto L14
	} else {
		goto L16
	}
L16:
	;
	v22 = *(*int32)(unsafe.Add(mBase, uint32(v18)))
	if v22 == int32(110) {
		goto L6
	} else {
		goto L17
	}
L17:
	;
	goto L14
L18:
	;
	goto L13
L19:
	;
	v32 = v26
	goto L20
L20:
	;
	v33 = *(*int32)(unsafe.Add(mBase, uint32(v32)+8))
	if v33 != l1 {
		goto L22
	} else {
		goto L23
	}
L21:
	;
	goto L7
L22:
	;
	v39 = *(*int32)(unsafe.Add(mBase, uint32(v32)+24))
	if v39 != 0 {
		v32 = v39
		goto L20
	} else {
		goto L26
	}
L23:
	;
	v35 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v32)+4)))
	if v35 != 0 {
		goto L22
	} else {
		goto L24
	}
L24:
	;
	v36 = *(*int32)(unsafe.Add(mBase, uint32(v32)))
	if v36 == int32(110) {
		goto L6
	} else {
		goto L25
	}
L25:
	;
	goto L22
L26:
	;
	goto L21
L27:
	;
	goto L6
}
func F_nocache_index_getattr(m *base.Module, l0 int32, l1 int32, l2 int32) int32 {
	mBase := m.M
	_ = mBase
	var v4 int32
	_ = v4
	var v15 int32
	_ = v15
	var v17 int32
	_ = v17
	var v21 int32
	_ = v21
	var v23 int32
	_ = v23
	var v24 int32
	_ = v24
	var v26 int32
	_ = v26
	var v28 int32
	_ = v28
	var v30 int32
	_ = v30
	var v32 int32
	_ = v32
	var v33 int32
	_ = v33
	var v40 int32
	_ = v40
	var v44 int32
	_ = v44
	var v58 int32
	_ = v58
	var v62 int32
	_ = v62
	var v89 int32
	_ = v89
	var v93 int32
	_ = v93
	var v95 int32
	_ = v95
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
	var v111 int32
	_ = v111
	var v112 int32
	_ = v112
	var v118 int32
	_ = v118
	var v124 int32
	_ = v124
	var v129 int32
	_ = v129
	var v138 int32
	_ = v138
	var v154 int32
	_ = v154
	var v158 int32
	_ = v158
	var v176 int32
	_ = v176
	var v177 int32
	_ = v177
	var v181 int32
	_ = v181
	var v197 int32
	_ = v197
	var v201 int32
	_ = v201
	var v204 int32
	_ = v204
	var v222 int32
	_ = v222
	var v223 int32
	_ = v223
	var v224 int32
	_ = v224
	var v227 int32
	_ = v227
	var v232 int32
	_ = v232
	var v242 int32
	_ = v242
	var v243 int32
	_ = v243
	var v246 int32
	_ = v246
	var v248 int32
	_ = v248
	var v252 int32
	_ = v252
	var v256 int32
	_ = v256
	var v272 int32
	_ = v272
	var v283 int32
	_ = v283
	var v284 int32
	_ = v284
	var v287 int32
	_ = v287
	var v290 int32
	_ = v290
	var v294 int32
	_ = v294
	var v296 int32
	_ = v296
	var v297 int32
	_ = v297
	var v298 int32
	_ = v298
	var v313 int32
	_ = v313
	var v324 int32
	_ = v324
	var v327 int32
	_ = v327
	var v328 int32
	_ = v328
	var v331 int32
	_ = v331
	var v334 int32
	_ = v334
	var v335 int32
	_ = v335
	var v342 int32
	_ = v342
	var v348 int32
	_ = v348
	var v349 int32
	_ = v349
	var v357 int32
	_ = v357
	var v359 int32
	_ = v359
	var v360 int32
	_ = v360
	var v368 int32
	_ = v368
	var v369 int32
	_ = v369
	var v371 int32
	_ = v371
	var v374 int32
	_ = v374
	var v377 int32
	_ = v377
	var v380 int32
	_ = v380
	var v382 int32
	_ = v382
	var v384 int32
	_ = v384
	var v387 int32
	_ = v387
	var v392 int32
	_ = v392
	var v395 int32
	_ = v395
	var v400 int32
	_ = v400
	var v410 int32
	_ = v410
	var v415 int32
	_ = v415
	var v419 int32
	_ = v419
	var v424 int32
	_ = v424
	var v426 int32
	_ = v426
	var v430 int32
	_ = v430
	var v436 int32
	_ = v436
	var v439 int32
	_ = v439
	var v445 int32
	_ = v445
	var v449 int32
	_ = v449
	var v451 int32
	_ = v451
	var v459 int32
	_ = v459
	var v462 int32
	_ = v462
	var v468 int32
	_ = v468
	var v475 int32
	_ = v475
	var v482 int32
	_ = v482
	var v485 int32
	_ = v485
	var v488 int32
	_ = v488
	var v489 int32
	_ = v489
	var v492 int32
	_ = v492
	var v497 int32
	_ = v497
	var v498 int32
	_ = v498
	var v502 int32
	_ = v502
	var v506 int32
	_ = v506
	var v511 int32
	_ = v511
	var v513 int32
	_ = v513
	var v526 int32
	_ = v526
	var v528 int32
	_ = v528
	v4 = int32(0)
	v15 = m.G0
	v17 = v15 - int32(32)
	m.G0 = v17
	v21 = int32(*(*int16)(unsafe.Add(mBase, uint32(l0)+6)))
	v23 = base.B2i32(v4 <= v21)
	if v4 <= v21 {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	v24 = int32(8)
	goto L3
L2:
	;
	v24 = int32(16)
	goto L3
L3:
	;
	v26 = l1 - int32(1)
	if v4 <= v21 {
		v89 = v4
		goto L8
	} else {
		goto L9
	}
L4:
	;
	m.G0 = v17 + int32(32)
	return v528
L5:
	;
	v526 = *(*int32)(unsafe.Add(mBase, uint32(v513)))
	v528 = v526
	goto L4
L6:
	;
	v485 = v475 + v482
	v488 = l2 + v26<<(uint(int32(4))%32)
	v489 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v488)+26)))
	if v489 != int32(1) {
		v528 = v485
		goto L4
	} else {
		goto L103
	}
L7:
	;
	v287 = l0 + v24
	v290 = int32(0)
	v294 = v290
	v296 = int32(1)
	v297 = v290
	v298 = v21
	goto L47
L8:
	;
	v93 = l0 + v24
	v95 = l2 + int32(20)
	v98 = v95 + v26<<(uint(int32(4))%32)
	v99 = *(*int32)(unsafe.Add(mBase, uint32(v98)))
	if int32(0) <= v99 {
		goto L17
	} else {
		goto L18
	}
L9:
	;
	v28 = l0 + int32(8)
	v30 = v26 >> (uint(int32(3)) % 32)
	v32 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v28+v30))))
	v33 = int32(-1)
	if v32|v33<<(uint(v26&int32(7))%32) != v33 {
		goto L10
	} else {
		goto L11
	}
L10:
	;
	v283 = v28
	v284 = l0 + v24
	goto L7
L11:
	;
	v40 = int32(0)
	if v30 <= v40 {
		v89 = v28
		goto L8
	} else {
		goto L12
	}
L12:
	;
	v44 = v40
	goto L13
L13:
	;
	v58 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v44+v28))))
	if v58 != int32(255) {
		goto L10
	} else {
		goto L15
	}
L14:
	;
	v89 = v28
	goto L8
L15:
	;
	v62 = v44 + int32(1)
	if v30 != v62 {
		v44 = v62
		goto L13
	} else {
		goto L16
	}
L16:
	;
	goto L14
L17:
	;
	v102 = v99 + v93
	v103 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v98)+6)))
	if v103 != int32(1) {
		v528 = v102
		goto L4
	} else {
		goto L20
	}
L18:
	;
	goto L19
L19:
	;
	if v21&int32(_a_F_nocache_index_getattr_0) == int32(0) {
		goto L28
	} else {
		goto L29
	}
L20:
	;
	v106 = int32(*(*int16)(unsafe.Add(mBase, uint32(v98)+4)))
	switch v106&int32(_a_F_nocache_index_getattr_1) - int32(1) {
	case 0:
		goto L23
	case 1:
		goto L22
	default:
		goto L21
	case 3:
		v513 = v102
		goto L5
	}
L21:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v118 = m.ExcPending
	if v118 != 0 {
		goto L24
	} else {
		goto L25
	}
L22:
	;
	v112 = int32(*(*int16)(unsafe.Add(mBase, uint32(v102))))
	v528 = v112
	goto L4
L23:
	;
	v111 = int32(*(*int8)(unsafe.Add(mBase, uint32(v102))))
	v528 = v111
	goto L4
L24:
	;
	return int32(0)
L25:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v17)+16)) = v106
	F_errmsg_internal(m, int32(_a_F_nocache_index_getattr_2), v17+int32(16))
	mBase = m.M
	v124 = m.ExcPending
	if v124 != 0 {
		goto L24
	} else {
		goto L26
	}
L26:
	;
	F_errfinish(m, int32(_a_F_nocache_index_getattr_3), int32(70), int32(_a_F_nocache_index_getattr_4))
	mBase = m.M
	v129 = m.ExcPending
	if v129 != 0 {
		goto L24
	} else {
		goto L27
	}
L27:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L28:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l2)+20)) = int32(0)
	v176 = int32(1)
	v177 = *(*int32)(unsafe.Add(mBase, uint32(l2)))
	if v177 < int32(2) {
		v204 = v176
		goto L36
	} else {
		goto L37
	}
L29:
	;
	if v26 < int32(0) {
		goto L28
	} else {
		goto L30
	}
L30:
	;
	v138 = int32(0)
	goto L31
L31:
	;
	v154 = int32(*(*int16)(unsafe.Add(mBase, uint32(v95+v138<<(uint(int32(4))%32))+4)))
	if v154 <= int32(0) {
		v283 = v89
		v284 = v93
		goto L7
	} else {
		goto L33
	}
L32:
	;
	goto L28
L33:
	;
	v158 = v138 + int32(1)
	if v158 <= v26 {
		v138 = v158
		goto L31
	} else {
		goto L34
	}
L34:
	;
	goto L32
L35:
	;
	v272 = *(*int32)(unsafe.Add(mBase, uint32(v98)))
	v475 = v272
	v482 = v93
	goto L6
L36:
	;
	if v177 <= v204 {
		goto L35
	} else {
		goto L42
	}
L37:
	;
	v181 = v176
	goto L38
L38:
	;
	v197 = *(*int32)(unsafe.Add(mBase, uint32(v95+v181<<(uint(int32(4))%32))))
	if v197 <= int32(0) {
		v204 = v181
		goto L36
	} else {
		goto L40
	}
L39:
	;
	goto L35
L40:
	;
	v201 = v181 + int32(1)
	if v201 != v177 {
		v181 = v201
		goto L38
	} else {
		goto L41
	}
L41:
	;
	goto L39
L42:
	;
	v222 = v204<<(uint(int32(4))%32) + v95 - int32(16)
	v223 = *(*int32)(unsafe.Add(mBase, uint32(v222)))
	v224 = int32(*(*int16)(unsafe.Add(mBase, uint32(v222)+4)))
	v227 = v204
	v232 = v223 + v224
	goto L43
L43:
	;
	v242 = v95 + v227<<(uint(int32(4))%32)
	v243 = int32(*(*int16)(unsafe.Add(mBase, uint32(v242)+4)))
	if v243 <= int32(0) {
		goto L35
	} else {
		goto L45
	}
L44:
	;
	goto L35
L45:
	;
	v246 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v242)+12)))
	v248 = int32(1)
	v252 = (v232 + v246 - v248) & (int32(0) - v246)
	*(*int32)(unsafe.Add(mBase, uint32(v242))) = v252
	v256 = v227 + v248
	if v256 != v177 {
		v227 = v256
		v232 = v243 + v252
		goto L43
	} else {
		goto L46
	}
L46:
	;
	goto L44
L47:
	;
	if int32(0) <= base.I32_extend16_s(v298) {
		goto L49
	} else {
		goto L50
	}
L49:
	;
	v324 = l2 + int32(20) + v294<<(uint(int32(4))%32)
	if v296&int32(1) != 0 {
		goto L55
	} else {
		goto L56
	}
L50:
	;
	v313 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v283+v294>>(uint(int32(3))%32)))))
	if int32(base.Ui32(v313)>>(uint(v294&int32(7))%32))&int32(1) != 0 {
		goto L49
	} else {
		goto L51
	}
L51:
	;
	v294 = v294 + int32(1)
	v296 = int32(0)
	goto L47
L52:
	;
	if v294 == v26 {
		v475 = v368
		v482 = v284
		goto L6
	} else {
		goto L65
	}
L53:
	;
	v357 = int32(0)
	v359 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v297+v287))))
	if v359 != 0 {
		v368 = v297
		v369 = v357
		goto L52
	} else {
		goto L64
	}
L54:
	;
	v342 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v324)+12)))
	v348 = (v297 + v342 - int32(1)) & (int32(0) - v342)
	v349 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v324)+4)))
	if v349 != int32(_a_F_nocache_index_getattr_1) {
		goto L60
	} else {
		goto L61
	}
L55:
	;
	v327 = int32(1)
	v328 = *(*int32)(unsafe.Add(mBase, uint32(v324)))
	if v328 < int32(0) {
		goto L54
	} else {
		goto L58
	}
L56:
	;
	goto L57
L57:
	;
	v331 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v324)+4)))
	if v331 == int32(_a_F_nocache_index_getattr_1) {
		goto L53
	} else {
		goto L59
	}
L58:
	;
	v368 = v328
	v369 = v327
	goto L52
L59:
	;
	v334 = int32(0)
	v335 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v324)+12)))
	v368 = (v297 + v335 - int32(1)) & (v334 - v335)
	v369 = v334
	goto L52
L60:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v324))) = v348
	v368 = v348
	v369 = v327
	goto L52
L61:
	;
	goto L62
L62:
	;
	if v348 != v297 {
		goto L53
	} else {
		goto L63
	}
L63:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v324))) = v297
	v368 = v297
	v369 = v327
	goto L52
L64:
	;
	v360 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v324)+12)))
	v368 = (v297 + v360 - int32(1)) & (int32(0) - v360)
	v369 = v357
	goto L52
L65:
	;
	v371 = int32(*(*int16)(unsafe.Add(mBase, uint32(v324)+4)))
	if int32(0) < v371 {
		v462 = v371
		goto L66
	} else {
		goto L67
	}
L66:
	;
	v468 = int32(*(*uint16)(unsafe.Add(mBase, uint32(l0)+6)))
	v294 = v294 + int32(1)
	v296 = v369 & base.B2i32(int32(0) < v371)
	v297 = v462 + v368
	v298 = v468
	goto L47
L67:
	;
	v374 = v368 + v287
	if v371 == int32(-1) {
		goto L68
	} else {
		goto L69
	}
L68:
	;
	v377 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v374))))
	if v377 == int32(1) {
		goto L71
	} else {
		goto L72
	}
L69:
	;
	goto L70
L70:
	;
	if v374&int32(3) == int32(0) {
		v426 = v374
		goto L88
	} else {
		goto L89
	}
L71:
	;
	v380 = int32(6)
	v382 = int32(18)
	v384 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v374)+1)))
	if v384 == v382 {
		goto L74
	} else {
		goto L75
	}
L72:
	;
	goto L73
L73:
	;
	if v377&int32(1) != 0 {
		goto L83
	} else {
		goto L84
	}
L74:
	;
	v387 = v382
	goto L76
L75:
	;
	v387 = int32(2)
	goto L76
L76:
	;
	if v384&int32(254) == int32(2) {
		goto L77
	} else {
		goto L78
	}
L77:
	;
	v392 = v380
	goto L79
L78:
	;
	v392 = v387
	goto L79
L79:
	;
	if v384 == int32(1) {
		goto L80
	} else {
		goto L81
	}
L80:
	;
	v395 = v380
	goto L82
L81:
	;
	v395 = v392
	goto L82
L82:
	;
	v462 = v395
	goto L66
L83:
	;
	v462 = int32(base.Ui32(v377) >> (uint(int32(1)) % 32))
	goto L66
L84:
	;
	goto L85
L85:
	;
	v400 = *(*int32)(unsafe.Add(mBase, uint32(v374)))
	v462 = int32(base.Ui32(v400) >> (uint(int32(2)) % 32))
	goto L66
L86:
	;
	v462 = v459 + int32(1)
	goto L66
L87:
	;
	v459 = v451 - v374
	goto L86
L88:
	;
	v430 = v426
	goto L97
L89:
	;
	v410 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v374))))
	if v410 == int32(0) {
		goto L90
	} else {
		goto L91
	}
L90:
	;
	v459 = int32(0)
	goto L86
L91:
	;
	goto L92
L92:
	;
	v415 = v374
	goto L93
L93:
	;
	v419 = v415 + int32(1)
	if v419&int32(3) == int32(0) {
		v426 = v419
		goto L88
	} else {
		goto L95
	}
L94:
	;
	v451 = v419
	goto L87
L95:
	;
	v424 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v419))))
	if v424 != 0 {
		v415 = v419
		goto L93
	} else {
		goto L96
	}
L96:
	;
	goto L94
L97:
	;
	v436 = *(*int32)(unsafe.Add(mBase, uint32(v430)))
	v439 = int32(-2139062144)
	if (int32(16843008)-v436|v436)&v439 == v439 {
		v430 = v430 + int32(4)
		goto L97
	} else {
		goto L99
	}
L98:
	;
	v445 = v430
	goto L100
L99:
	;
	goto L98
L100:
	;
	v449 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v445))))
	if v449 != 0 {
		v445 = v445 + int32(1)
		goto L100
	} else {
		goto L102
	}
L101:
	;
	v451 = v445
	goto L87
L102:
	;
	goto L101
L103:
	;
	v492 = int32(*(*int16)(unsafe.Add(mBase, uint32(v488)+24)))
	switch v492&int32(_a_F_nocache_index_getattr_1) - int32(1) {
	case 0:
		goto L106
	case 1:
		goto L105
	default:
		goto L104
	case 3:
		v513 = v485
		goto L5
	}
L104:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v502 = m.ExcPending
	if v502 != 0 {
		goto L24
	} else {
		goto L107
	}
L105:
	;
	v498 = int32(*(*int16)(unsafe.Add(mBase, uint32(v485))))
	v528 = v498
	goto L4
L106:
	;
	v497 = int32(*(*int8)(unsafe.Add(mBase, uint32(v485))))
	v528 = v497
	goto L4
L107:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v17))) = v492
	F_errmsg_internal(m, int32(_a_F_nocache_index_getattr_2), v17)
	mBase = m.M
	v506 = m.ExcPending
	if v506 != 0 {
		goto L24
	} else {
		goto L108
	}
L108:
	;
	F_errfinish(m, int32(_a_F_nocache_index_getattr_3), int32(70), int32(_a_F_nocache_index_getattr_4))
	mBase = m.M
	v511 = m.ExcPending
	if v511 != 0 {
		goto L24
	} else {
		goto L109
	}
L109:
	;
	base.Wasm_trap_unreachable()
	for {
	}
}
func F_nodeToString(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v2 int32
	_ = v2
	var v4 int32
	_ = v4
	var v6 int32
	_ = v6
	var v8 int32
	_ = v8
	var v9 int32
	_ = v9
	var v16 int32
	_ = v16
	var v18 int32
	_ = v18
	var v21 int32
	_ = v21
	var v23 int32
	_ = v23
	v2 = int32(0)
	v4 = m.G0
	v6 = v4 - int32(16)
	m.G0 = v6
	v8 = int32(_a_F_nodeToString_0)
	v9 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_nodeToString[0])))
	*(*uint8)(unsafe.Add(mBase, _c_F_nodeToString[0])) = uint8(v2)
	F_initStringInfo(m, v6)
	mBase = m.M
	v16 = m.ExcPending
	if v16 != 0 {
		return int32(0)
	} else {
		F_outNode(m, v6, l0)
		mBase = m.M
		v18 = m.ExcPending
		if v18 != 0 {
			return int32(0)
		} else {
			v21 = v9 & int32(1)
			*(*uint8)(unsafe.Add(mBase, _c_F_nodeToString[0])) = uint8(v21)
			v23 = *(*int32)(unsafe.Add(mBase, uint32(v6)))
			m.G0 = v6 + int32(16)
			return v23
		}
	}
}
func F_notification_hash(m *base.Module, l0 int32, l1 int32) int32 {
	mBase := m.M
	_ = mBase
	var v3 int32
	_ = v3
	var v5 int32
	_ = v5
	var v6 int32
	_ = v6
	var v7 int32
	_ = v7
	var v10 int32
	_ = v10
	var v16 int32
	_ = v16
	var v23 int32
	_ = v23
	var v24 int32
	_ = v24
	var v25 int32
	_ = v25
	var v26 int32
	_ = v26
	var v27 int32
	_ = v27
	var v29 int32
	_ = v29
	var v30 int32
	_ = v30
	var v31 int32
	_ = v31
	var v33 int32
	_ = v33
	var v34 int32
	_ = v34
	var v36 int32
	_ = v36
	var v38 int32
	_ = v38
	var v42 int32
	_ = v42
	var v43 int32
	_ = v43
	var v44 int32
	_ = v44
	var v45 int32
	_ = v45
	var v49 int32
	_ = v49
	var v53 int32
	_ = v53
	var v57 int32
	_ = v57
	var v58 int32
	_ = v58
	var v59 int32
	_ = v59
	var v60 int32
	_ = v60
	var v64 int32
	_ = v64
	var v65 int32
	_ = v65
	var v66 int32
	_ = v66
	var v68 int32
	_ = v68
	var v71 int32
	_ = v71
	var v72 int32
	_ = v72
	var v73 int32
	_ = v73
	var v74 int32
	_ = v74
	var v75 int32
	_ = v75
	var v79 int32
	_ = v79
	var v83 int32
	_ = v83
	var v84 int32
	_ = v84
	var v88 int32
	_ = v88
	var v89 int32
	_ = v89
	var v93 int32
	_ = v93
	var v94 int32
	_ = v94
	var v96 int32
	_ = v96
	var v98 int32
	_ = v98
	var v102 int32
	_ = v102
	var v103 int32
	_ = v103
	var v107 int32
	_ = v107
	var v108 int32
	_ = v108
	var v110 int32
	_ = v110
	var v111 int32
	_ = v111
	var v113 int32
	_ = v113
	var v117 int32
	_ = v117
	var v118 int32
	_ = v118
	var v122 int32
	_ = v122
	var v123 int32
	_ = v123
	var v125 int32
	_ = v125
	var v126 int32
	_ = v126
	var v127 int32
	_ = v127
	var v128 int32
	_ = v128
	var v129 int32
	_ = v129
	var v131 int32
	_ = v131
	var v132 int32
	_ = v132
	var v133 int32
	_ = v133
	var v135 int32
	_ = v135
	var v136 int32
	_ = v136
	var v138 int32
	_ = v138
	var v140 int32
	_ = v140
	var v144 int32
	_ = v144
	var v145 int32
	_ = v145
	var v146 int32
	_ = v146
	var v147 int32
	_ = v147
	var v151 int32
	_ = v151
	var v155 int32
	_ = v155
	var v159 int32
	_ = v159
	var v160 int32
	_ = v160
	var v161 int32
	_ = v161
	var v162 int32
	_ = v162
	var v166 int32
	_ = v166
	var v167 int32
	_ = v167
	var v168 int32
	_ = v168
	var v170 int32
	_ = v170
	var v173 int32
	_ = v173
	var v174 int32
	_ = v174
	var v175 int32
	_ = v175
	var v176 int32
	_ = v176
	var v177 int32
	_ = v177
	var v181 int32
	_ = v181
	var v185 int32
	_ = v185
	var v186 int32
	_ = v186
	var v190 int32
	_ = v190
	var v191 int32
	_ = v191
	var v195 int32
	_ = v195
	var v196 int32
	_ = v196
	var v200 int32
	_ = v200
	var v201 int32
	_ = v201
	var v202 int32
	_ = v202
	var v206 int32
	_ = v206
	var v207 int32
	_ = v207
	var v208 int32
	_ = v208
	var v212 int32
	_ = v212
	var v213 int32
	_ = v213
	var v214 int32
	_ = v214
	var v216 int32
	_ = v216
	var v217 int32
	_ = v217
	var v218 int32
	_ = v218
	var v222 int32
	_ = v222
	var v223 int32
	_ = v223
	var v224 int32
	_ = v224
	var v225 int32
	_ = v225
	var v229 int32
	_ = v229
	var v230 int32
	_ = v230
	var v231 int32
	_ = v231
	var v232 int32
	_ = v232
	var v236 int32
	_ = v236
	var v237 int32
	_ = v237
	var v238 int32
	_ = v238
	var v239 int32
	_ = v239
	var v243 int32
	_ = v243
	var v244 int32
	_ = v244
	var v245 int32
	_ = v245
	var v248 int32
	_ = v248
	var v250 int32
	_ = v250
	var v254 int32
	_ = v254
	var v258 int32
	_ = v258
	var v262 int32
	_ = v262
	var v266 int32
	_ = v266
	var v270 int32
	_ = v270
	v3 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v5 = v3 + int32(4)
	v6 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v3))))
	v7 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v3)+2)))
	v10 = v6 + v7 + int32(1)
	v16 = v10 - int32(1636608432)
	if v5&int32(3) != 0 {
		if base.Ui32(int32(11)) < base.Ui32(v10) {
			v125 = v5
			v126 = v10
			v127 = v16
			v128 = v16
			v129 = v16
			for {
				v131 = *(*int32)(unsafe.Add(mBase, uint32(v125)+4))
				v132 = v131 + v128
				v133 = *(*int32)(unsafe.Add(mBase, uint32(v125)))
				v135 = *(*int32)(unsafe.Add(mBase, uint32(v125)+8))
				v136 = v135 + v129
				v138 = int32(4)
				v140 = v133 + v127 - v136 ^ base.I32_rotl(v136, v138)
				v144 = v132 - v140 ^ base.I32_rotl(v140, int32(6))
				v145 = v136 + v132
				v146 = v140 + v145
				v147 = v144 + v146
				v151 = v145 - v144 ^ base.I32_rotl(v144, int32(8))
				v155 = v146 - v151 ^ base.I32_rotl(v151, int32(16))
				v159 = v147 - v155 ^ base.I32_rotl(v155, int32(19))
				v160 = v151 + v147
				v161 = v155 + v160
				v162 = v159 + v161
				v166 = v160 - v159 ^ base.I32_rotl(v159, v138)
				v167 = int32(12)
				v168 = v125 + v167
				v170 = v126 - v167
				if base.Ui32(int32(11)) < base.Ui32(v170) {
					v125 = v168
					v126 = v170
					v127 = v161
					v128 = v162
					v129 = v166
					continue
				} else {
					break
				}
				break
			}
			v173 = v168
			v174 = v170
			v175 = v161
			v176 = v162
			v177 = v166
		} else {
			v173 = v5
			v174 = v10
			v175 = v16
			v176 = v16
			v177 = v16
		}
		switch v174 - int32(1) {
		case 0:
			v236 = v175
			v237 = v176
			v238 = v177
			v239 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v173))))
			v243 = v236 + v239
			v244 = v237
			v245 = v238
		case 1:
			v229 = v175
			v230 = v176
			v231 = v177
			v232 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v173)+1)))
			v236 = v232<<(uint(int32(8))%32) + v229
			v237 = v230
			v238 = v231
			v239 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v173))))
			v243 = v236 + v239
			v244 = v237
			v245 = v238
		case 2:
			v222 = v175
			v223 = v176
			v224 = v177
			v225 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v173)+2)))
			v229 = v225<<(uint(int32(16))%32) + v222
			v230 = v223
			v231 = v224
			v232 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v173)+1)))
			v236 = v232<<(uint(int32(8))%32) + v229
			v237 = v230
			v238 = v231
			v239 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v173))))
			v243 = v236 + v239
			v244 = v237
			v245 = v238
		case 3:
			v216 = v176
			v217 = v177
			v218 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v173)+3)))
			v222 = v218<<(uint(int32(24))%32) + v175
			v223 = v216
			v224 = v217
			v225 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v173)+2)))
			v229 = v225<<(uint(int32(16))%32) + v222
			v230 = v223
			v231 = v224
			v232 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v173)+1)))
			v236 = v232<<(uint(int32(8))%32) + v229
			v237 = v230
			v238 = v231
			v239 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v173))))
			v243 = v236 + v239
			v244 = v237
			v245 = v238
		case 4:
			v212 = v176
			v213 = v177
			v214 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v173)+4)))
			v216 = v212 + v214
			v217 = v213
			v218 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v173)+3)))
			v222 = v218<<(uint(int32(24))%32) + v175
			v223 = v216
			v224 = v217
			v225 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v173)+2)))
			v229 = v225<<(uint(int32(16))%32) + v222
			v230 = v223
			v231 = v224
			v232 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v173)+1)))
			v236 = v232<<(uint(int32(8))%32) + v229
			v237 = v230
			v238 = v231
			v239 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v173))))
			v243 = v236 + v239
			v244 = v237
			v245 = v238
		case 5:
			v206 = v176
			v207 = v177
			v208 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v173)+5)))
			v212 = v208<<(uint(int32(8))%32) + v206
			v213 = v207
			v214 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v173)+4)))
			v216 = v212 + v214
			v217 = v213
			v218 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v173)+3)))
			v222 = v218<<(uint(int32(24))%32) + v175
			v223 = v216
			v224 = v217
			v225 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v173)+2)))
			v229 = v225<<(uint(int32(16))%32) + v222
			v230 = v223
			v231 = v224
			v232 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v173)+1)))
			v236 = v232<<(uint(int32(8))%32) + v229
			v237 = v230
			v238 = v231
			v239 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v173))))
			v243 = v236 + v239
			v244 = v237
			v245 = v238
		case 6:
			v200 = v176
			v201 = v177
			v202 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v173)+6)))
			v206 = v202<<(uint(int32(16))%32) + v200
			v207 = v201
			v208 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v173)+5)))
			v212 = v208<<(uint(int32(8))%32) + v206
			v213 = v207
			v214 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v173)+4)))
			v216 = v212 + v214
			v217 = v213
			v218 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v173)+3)))
			v222 = v218<<(uint(int32(24))%32) + v175
			v223 = v216
			v224 = v217
			v225 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v173)+2)))
			v229 = v225<<(uint(int32(16))%32) + v222
			v230 = v223
			v231 = v224
			v232 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v173)+1)))
			v236 = v232<<(uint(int32(8))%32) + v229
			v237 = v230
			v238 = v231
			v239 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v173))))
			v243 = v236 + v239
			v244 = v237
			v245 = v238
		case 7:
			v195 = v177
			v196 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v173)+7)))
			v200 = v196<<(uint(int32(24))%32) + v176
			v201 = v195
			v202 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v173)+6)))
			v206 = v202<<(uint(int32(16))%32) + v200
			v207 = v201
			v208 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v173)+5)))
			v212 = v208<<(uint(int32(8))%32) + v206
			v213 = v207
			v214 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v173)+4)))
			v216 = v212 + v214
			v217 = v213
			v218 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v173)+3)))
			v222 = v218<<(uint(int32(24))%32) + v175
			v223 = v216
			v224 = v217
			v225 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v173)+2)))
			v229 = v225<<(uint(int32(16))%32) + v222
			v230 = v223
			v231 = v224
			v232 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v173)+1)))
			v236 = v232<<(uint(int32(8))%32) + v229
			v237 = v230
			v238 = v231
			v239 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v173))))
			v243 = v236 + v239
			v244 = v237
			v245 = v238
		case 8:
			v190 = v177
			v191 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v173)+8)))
			v195 = v191<<(uint(int32(8))%32) + v190
			v196 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v173)+7)))
			v200 = v196<<(uint(int32(24))%32) + v176
			v201 = v195
			v202 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v173)+6)))
			v206 = v202<<(uint(int32(16))%32) + v200
			v207 = v201
			v208 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v173)+5)))
			v212 = v208<<(uint(int32(8))%32) + v206
			v213 = v207
			v214 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v173)+4)))
			v216 = v212 + v214
			v217 = v213
			v218 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v173)+3)))
			v222 = v218<<(uint(int32(24))%32) + v175
			v223 = v216
			v224 = v217
			v225 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v173)+2)))
			v229 = v225<<(uint(int32(16))%32) + v222
			v230 = v223
			v231 = v224
			v232 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v173)+1)))
			v236 = v232<<(uint(int32(8))%32) + v229
			v237 = v230
			v238 = v231
			v239 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v173))))
			v243 = v236 + v239
			v244 = v237
			v245 = v238
		case 9:
			v185 = v177
			v186 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v173)+9)))
			v190 = v186<<(uint(int32(16))%32) + v185
			v191 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v173)+8)))
			v195 = v191<<(uint(int32(8))%32) + v190
			v196 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v173)+7)))
			v200 = v196<<(uint(int32(24))%32) + v176
			v201 = v195
			v202 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v173)+6)))
			v206 = v202<<(uint(int32(16))%32) + v200
			v207 = v201
			v208 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v173)+5)))
			v212 = v208<<(uint(int32(8))%32) + v206
			v213 = v207
			v214 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v173)+4)))
			v216 = v212 + v214
			v217 = v213
			v218 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v173)+3)))
			v222 = v218<<(uint(int32(24))%32) + v175
			v223 = v216
			v224 = v217
			v225 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v173)+2)))
			v229 = v225<<(uint(int32(16))%32) + v222
			v230 = v223
			v231 = v224
			v232 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v173)+1)))
			v236 = v232<<(uint(int32(8))%32) + v229
			v237 = v230
			v238 = v231
			v239 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v173))))
			v243 = v236 + v239
			v244 = v237
			v245 = v238
		case 10:
			v181 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v173)+10)))
			v185 = v181<<(uint(int32(24))%32) + v177
			v186 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v173)+9)))
			v190 = v186<<(uint(int32(16))%32) + v185
			v191 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v173)+8)))
			v195 = v191<<(uint(int32(8))%32) + v190
			v196 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v173)+7)))
			v200 = v196<<(uint(int32(24))%32) + v176
			v201 = v195
			v202 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v173)+6)))
			v206 = v202<<(uint(int32(16))%32) + v200
			v207 = v201
			v208 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v173)+5)))
			v212 = v208<<(uint(int32(8))%32) + v206
			v213 = v207
			v214 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v173)+4)))
			v216 = v212 + v214
			v217 = v213
			v218 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v173)+3)))
			v222 = v218<<(uint(int32(24))%32) + v175
			v223 = v216
			v224 = v217
			v225 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v173)+2)))
			v229 = v225<<(uint(int32(16))%32) + v222
			v230 = v223
			v231 = v224
			v232 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v173)+1)))
			v236 = v232<<(uint(int32(8))%32) + v229
			v237 = v230
			v238 = v231
			v239 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v173))))
			v243 = v236 + v239
			v244 = v237
			v245 = v238
		default:
			v243 = v175
			v244 = v176
			v245 = v177
		}
	} else {
		if base.Ui32(v10) < base.Ui32(int32(12)) {
			v71 = v5
			v72 = v10
			v73 = v16
			v74 = v16
			v75 = v16
		} else {
			v23 = v5
			v24 = v10
			v25 = v16
			v26 = v16
			v27 = v16
			for {
				v29 = *(*int32)(unsafe.Add(mBase, uint32(v23)+4))
				v30 = v29 + v26
				v31 = *(*int32)(unsafe.Add(mBase, uint32(v23)))
				v33 = *(*int32)(unsafe.Add(mBase, uint32(v23)+8))
				v34 = v33 + v27
				v36 = int32(4)
				v38 = v31 + v25 - v34 ^ base.I32_rotl(v34, v36)
				v42 = v30 - v38 ^ base.I32_rotl(v38, int32(6))
				v43 = v34 + v30
				v44 = v38 + v43
				v45 = v42 + v44
				v49 = v43 - v42 ^ base.I32_rotl(v42, int32(8))
				v53 = v44 - v49 ^ base.I32_rotl(v49, int32(16))
				v57 = v45 - v53 ^ base.I32_rotl(v53, int32(19))
				v58 = v49 + v45
				v59 = v53 + v58
				v60 = v57 + v59
				v64 = v58 - v57 ^ base.I32_rotl(v57, v36)
				v65 = int32(12)
				v66 = v23 + v65
				v68 = v24 - v65
				if base.Ui32(int32(11)) < base.Ui32(v68) {
					v23 = v66
					v24 = v68
					v25 = v59
					v26 = v60
					v27 = v64
					continue
				} else {
					break
				}
				break
			}
			v71 = v66
			v72 = v68
			v73 = v59
			v74 = v60
			v75 = v64
		}
		switch v72 - int32(1) {
		case 0:
			v122 = v73
			v123 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v71))))
			v243 = v122 + v123
			v244 = v74
			v245 = v75
		case 1:
			v117 = v73
			v118 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v71)+1)))
			v122 = v118<<(uint(int32(8))%32) + v117
			v123 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v71))))
			v243 = v122 + v123
			v244 = v74
			v245 = v75
		case 2:
			v113 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v71)+2)))
			v117 = v113<<(uint(int32(16))%32) + v73
			v118 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v71)+1)))
			v122 = v118<<(uint(int32(8))%32) + v117
			v123 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v71))))
			v243 = v122 + v123
			v244 = v74
			v245 = v75
		case 3:
			v110 = v74
			v111 = *(*int32)(unsafe.Add(mBase, uint32(v71)))
			v243 = v111 + v73
			v244 = v110
			v245 = v75
		case 4:
			v107 = v74
			v108 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v71)+4)))
			v110 = v107 + v108
			v111 = *(*int32)(unsafe.Add(mBase, uint32(v71)))
			v243 = v111 + v73
			v244 = v110
			v245 = v75
		case 5:
			v102 = v74
			v103 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v71)+5)))
			v107 = v103<<(uint(int32(8))%32) + v102
			v108 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v71)+4)))
			v110 = v107 + v108
			v111 = *(*int32)(unsafe.Add(mBase, uint32(v71)))
			v243 = v111 + v73
			v244 = v110
			v245 = v75
		case 6:
			v98 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v71)+6)))
			v102 = v98<<(uint(int32(16))%32) + v74
			v103 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v71)+5)))
			v107 = v103<<(uint(int32(8))%32) + v102
			v108 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v71)+4)))
			v110 = v107 + v108
			v111 = *(*int32)(unsafe.Add(mBase, uint32(v71)))
			v243 = v111 + v73
			v244 = v110
			v245 = v75
		case 7:
			v93 = v75
			v94 = *(*int32)(unsafe.Add(mBase, uint32(v71)))
			v96 = *(*int32)(unsafe.Add(mBase, uint32(v71)+4))
			v243 = v94 + v73
			v244 = v96 + v74
			v245 = v93
		case 8:
			v88 = v75
			v89 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v71)+8)))
			v93 = v89<<(uint(int32(8))%32) + v88
			v94 = *(*int32)(unsafe.Add(mBase, uint32(v71)))
			v96 = *(*int32)(unsafe.Add(mBase, uint32(v71)+4))
			v243 = v94 + v73
			v244 = v96 + v74
			v245 = v93
		case 9:
			v83 = v75
			v84 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v71)+9)))
			v88 = v84<<(uint(int32(16))%32) + v83
			v89 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v71)+8)))
			v93 = v89<<(uint(int32(8))%32) + v88
			v94 = *(*int32)(unsafe.Add(mBase, uint32(v71)))
			v96 = *(*int32)(unsafe.Add(mBase, uint32(v71)+4))
			v243 = v94 + v73
			v244 = v96 + v74
			v245 = v93
		case 10:
			v79 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v71)+10)))
			v83 = v79<<(uint(int32(24))%32) + v75
			v84 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v71)+9)))
			v88 = v84<<(uint(int32(16))%32) + v83
			v89 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v71)+8)))
			v93 = v89<<(uint(int32(8))%32) + v88
			v94 = *(*int32)(unsafe.Add(mBase, uint32(v71)))
			v96 = *(*int32)(unsafe.Add(mBase, uint32(v71)+4))
			v243 = v94 + v73
			v244 = v96 + v74
			v245 = v93
		default:
			v243 = v73
			v244 = v74
			v245 = v75
		}
	}
	v248 = int32(14)
	v250 = v244 ^ v245 - base.I32_rotl(v244, v248)
	v254 = v250 ^ v243 - base.I32_rotl(v250, int32(11))
	v258 = v254 ^ v244 - base.I32_rotl(v254, int32(25))
	v262 = v258 ^ v250 - base.I32_rotl(v258, int32(16))
	v266 = v262 ^ v254 - base.I32_rotl(v262, int32(4))
	v270 = v266 ^ v258 - base.I32_rotl(v266, v248)
	return v270 ^ v262 - base.I32_rotl(v270, int32(24))
}
