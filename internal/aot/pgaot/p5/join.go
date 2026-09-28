package p5

import (
	base "github.com/shibukawa/pgmem/internal/aot/pgaot/base"
	"unsafe"
)

func F_ExecMergeJoin(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v19 int32
	_ = v19
	var v21 int32
	_ = v21
	var v24 int32
	_ = v24
	var v28 int32
	_ = v28
	var v29 int32
	_ = v29
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
	var v38 int32
	_ = v38
	var v40 int32
	_ = v40
	var v59 int32
	_ = v59
	var v62 int32
	_ = v62
	var v64 int32
	_ = v64
	var v65 int32
	_ = v65
	var v66 int32
	_ = v66
	var v67 int32
	_ = v67
	var v69 int32
	_ = v69
	var v70 int32
	_ = v70
	var v77 int32
	_ = v77
	var v78 int32
	_ = v78
	var v86 int32
	_ = v86
	var v90 int32
	_ = v90
	var v92 int32
	_ = v92
	var v93 int32
	_ = v93
	var v94 int32
	_ = v94
	var v95 int32
	_ = v95
	var v97 int32
	_ = v97
	var v98 int32
	_ = v98
	var v103 int32
	_ = v103
	var v107 int32
	_ = v107
	var v112 int32
	_ = v112
	var v113 int32
	_ = v113
	var v119 int32
	_ = v119
	var v125 int32
	_ = v125
	var v127 int32
	_ = v127
	var v129 int32
	_ = v129
	var v130 int32
	_ = v130
	var v132 int32
	_ = v132
	var v136 int32
	_ = v136
	var v137 int64
	_ = v137
	var v138 int32
	_ = v138
	var v144 int32
	_ = v144
	var v146 int32
	_ = v146
	var v151 int32
	_ = v151
	var v158 int32
	_ = v158
	var v159 int32
	_ = v159
	var v161 int32
	_ = v161
	var v165 int32
	_ = v165
	var v166 int64
	_ = v166
	var v167 int32
	_ = v167
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
	var v179 int32
	_ = v179
	var v180 int32
	_ = v180
	var v181 int32
	_ = v181
	var v183 int32
	_ = v183
	var v188 int32
	_ = v188
	var v189 int64
	_ = v189
	var v190 int32
	_ = v190
	var v193 int32
	_ = v193
	var v195 int32
	_ = v195
	var v197 int32
	_ = v197
	var v198 int32
	_ = v198
	var v200 int32
	_ = v200
	var v203 float64
	_ = v203
	var v207 int32
	_ = v207
	var v210 float64
	_ = v210
	var v218 int32
	_ = v218
	var v219 int32
	_ = v219
	var v221 int32
	_ = v221
	var v222 int32
	_ = v222
	var v224 int32
	_ = v224
	var v226 int32
	_ = v226
	var v227 int32
	_ = v227
	var v228 int32
	_ = v228
	var v229 int32
	_ = v229
	var v230 int32
	_ = v230
	var v233 int32
	_ = v233
	var v234 int32
	_ = v234
	var v237 int32
	_ = v237
	var v238 int32
	_ = v238
	var v240 int32
	_ = v240
	var v241 int32
	_ = v241
	var v242 int32
	_ = v242
	var v243 int32
	_ = v243
	var v245 int32
	_ = v245
	var v248 int32
	_ = v248
	var v252 int32
	_ = v252
	var v257 int32
	_ = v257
	var v259 int32
	_ = v259
	var v269 int32
	_ = v269
	var v272 int32
	_ = v272
	var v273 int32
	_ = v273
	var v274 int32
	_ = v274
	var v276 int32
	_ = v276
	var v278 int32
	_ = v278
	var v300 int32
	_ = v300
	var v304 int32
	_ = v304
	var v306 int32
	_ = v306
	var v309 int32
	_ = v309
	var v310 int32
	_ = v310
	var v311 int32
	_ = v311
	var v315 int32
	_ = v315
	var v316 int32
	_ = v316
	var v317 int32
	_ = v317
	var v322 int32
	_ = v322
	var v336 int32
	_ = v336
	var v359 int64
	_ = v359
	var v360 int64
	_ = v360
	var v363 int32
	_ = v363
	var v364 int32
	_ = v364
	var v365 int32
	_ = v365
	var v366 int32
	_ = v366
	var v373 int32
	_ = v373
	var v375 int32
	_ = v375
	var v376 int32
	_ = v376
	var v396 int32
	_ = v396
	var v448 int32
	_ = v448
	var v449 int32
	_ = v449
	var v453 int32
	_ = v453
	var v458 int32
	_ = v458
	var v459 int32
	_ = v459
	var v462 int32
	_ = v462
	var v464 int32
	_ = v464
	var v465 int32
	_ = v465
	var v467 int32
	_ = v467
	var v469 int32
	_ = v469
	var v470 int32
	_ = v470
	var v471 int32
	_ = v471
	var v472 int32
	_ = v472
	var v473 int32
	_ = v473
	var v474 int32
	_ = v474
	var v479 int32
	_ = v479
	var v484 int32
	_ = v484
	var v487 int32
	_ = v487
	var v489 int32
	_ = v489
	var v490 int32
	_ = v490
	var v492 int32
	_ = v492
	var v496 int32
	_ = v496
	var v497 int32
	_ = v497
	var v499 int32
	_ = v499
	var v500 int32
	_ = v500
	var v501 int32
	_ = v501
	var v502 int32
	_ = v502
	var v503 int32
	_ = v503
	var v504 int32
	_ = v504
	var v509 int32
	_ = v509
	var v518 int32
	_ = v518
	var v519 int32
	_ = v519
	var v521 int32
	_ = v521
	var v522 int32
	_ = v522
	var v524 int32
	_ = v524
	var v528 int32
	_ = v528
	var v529 int32
	_ = v529
	var v531 int32
	_ = v531
	var v532 int32
	_ = v532
	var v533 int32
	_ = v533
	var v534 int32
	_ = v534
	var v535 int32
	_ = v535
	var v538 int32
	_ = v538
	var v539 int32
	_ = v539
	var v546 int32
	_ = v546
	var v549 int32
	_ = v549
	var v552 int32
	_ = v552
	var v559 int32
	_ = v559
	var v560 int32
	_ = v560
	var v562 int32
	_ = v562
	var v563 int32
	_ = v563
	var v565 int32
	_ = v565
	var v567 int32
	_ = v567
	var v568 int32
	_ = v568
	var v569 int32
	_ = v569
	var v570 int32
	_ = v570
	var v571 int32
	_ = v571
	var v574 int32
	_ = v574
	var v575 int32
	_ = v575
	var v582 int32
	_ = v582
	var v587 int32
	_ = v587
	var v590 int32
	_ = v590
	var v595 int32
	_ = v595
	var v596 int32
	_ = v596
	var v598 int32
	_ = v598
	var v599 int32
	_ = v599
	var v600 int32
	_ = v600
	var v601 int32
	_ = v601
	var v603 int32
	_ = v603
	var v606 int32
	_ = v606
	var v610 int32
	_ = v610
	var v616 int32
	_ = v616
	var v617 int32
	_ = v617
	var v627 int32
	_ = v627
	var v630 int32
	_ = v630
	var v631 int32
	_ = v631
	var v632 int32
	_ = v632
	var v634 int32
	_ = v634
	var v636 int32
	_ = v636
	var v658 int32
	_ = v658
	var v662 int32
	_ = v662
	var v664 int32
	_ = v664
	var v667 int32
	_ = v667
	var v668 int32
	_ = v668
	var v669 int32
	_ = v669
	var v673 int32
	_ = v673
	var v674 int32
	_ = v674
	var v675 int32
	_ = v675
	var v680 int32
	_ = v680
	var v694 int32
	_ = v694
	var v717 int64
	_ = v717
	var v718 int64
	_ = v718
	var v721 int32
	_ = v721
	var v722 int32
	_ = v722
	var v723 int32
	_ = v723
	var v724 int32
	_ = v724
	var v731 int32
	_ = v731
	var v733 int32
	_ = v733
	var v734 int32
	_ = v734
	var v754 int32
	_ = v754
	var v757 int32
	_ = v757
	var v761 int32
	_ = v761
	var v762 int32
	_ = v762
	var v763 int32
	_ = v763
	var v764 int32
	_ = v764
	var v765 int32
	_ = v765
	var v767 int32
	_ = v767
	var v816 int32
	_ = v816
	var v817 int32
	_ = v817
	var v818 int32
	_ = v818
	var v819 int32
	_ = v819
	var v820 int32
	_ = v820
	var v822 int32
	_ = v822
	var v823 int32
	_ = v823
	var v824 int32
	_ = v824
	var v825 int32
	_ = v825
	var v827 int32
	_ = v827
	var v830 int32
	_ = v830
	var v834 int32
	_ = v834
	var v840 int32
	_ = v840
	var v841 int32
	_ = v841
	var v851 int32
	_ = v851
	var v854 int32
	_ = v854
	var v855 int32
	_ = v855
	var v856 int32
	_ = v856
	var v858 int32
	_ = v858
	var v860 int32
	_ = v860
	var v882 int32
	_ = v882
	var v886 int32
	_ = v886
	var v888 int32
	_ = v888
	var v891 int32
	_ = v891
	var v892 int32
	_ = v892
	var v893 int32
	_ = v893
	var v897 int32
	_ = v897
	var v898 int32
	_ = v898
	var v899 int32
	_ = v899
	var v904 int32
	_ = v904
	var v918 int32
	_ = v918
	var v939 int64
	_ = v939
	var v940 int64
	_ = v940
	var v943 int32
	_ = v943
	var v944 int32
	_ = v944
	var v945 int32
	_ = v945
	var v946 int32
	_ = v946
	var v953 int32
	_ = v953
	var v955 int32
	_ = v955
	var v956 int32
	_ = v956
	var v976 int32
	_ = v976
	var v979 int32
	_ = v979
	var v983 int32
	_ = v983
	var v1029 int32
	_ = v1029
	var v1030 int32
	_ = v1030
	var v1031 int32
	_ = v1031
	var v1064 int32
	_ = v1064
	var v1068 int32
	_ = v1068
	var v1073 int32
	_ = v1073
	var v1076 int32
	_ = v1076
	var v1077 int32
	_ = v1077
	var v1079 int32
	_ = v1079
	var v1080 int32
	_ = v1080
	var v1082 int32
	_ = v1082
	var v1084 int32
	_ = v1084
	var v1085 int32
	_ = v1085
	var v1086 int32
	_ = v1086
	var v1087 int32
	_ = v1087
	var v1088 int32
	_ = v1088
	var v1091 int32
	_ = v1091
	var v1092 int32
	_ = v1092
	var v1099 int32
	_ = v1099
	var v1104 int32
	_ = v1104
	var v1107 int32
	_ = v1107
	var v1113 int32
	_ = v1113
	var v1159 int32
	_ = v1159
	var v1163 int32
	_ = v1163
	var v1168 int32
	_ = v1168
	v19 = m.G0
	v21 = v19 - int32(16)
	m.G0 = v21
	v24 = *(*int32)(unsafe.Add(mBase, _c_F_ExecMergeJoin[0]))
	if v24 != 0 {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	F_ProcessInterrupts(m)
	mBase = m.M
	v28 = m.ExcPending
	if v28 != 0 {
		goto L4
	} else {
		goto L5
	}
L2:
	;
	goto L3
L3:
	;
	v29 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+132)))
	v30 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+131)))
	v31 = *(*int32)(unsafe.Add(mBase, uint32(l0)+32))
	v32 = *(*int32)(unsafe.Add(mBase, uint32(l0)+112))
	v33 = *(*int32)(unsafe.Add(mBase, uint32(l0)+36))
	v34 = *(*int32)(unsafe.Add(mBase, uint32(l0)+40))
	v35 = *(*int32)(unsafe.Add(mBase, uint32(l0)+64))
	v36 = *(*int32)(unsafe.Add(mBase, uint32(v35)+20))
	F_MemoryContextReset(m, v36)
	mBase = m.M
	v38 = m.ExcPending
	if v38 != 0 {
		goto L4
	} else {
		goto L6
	}
L4:
	;
	return int32(0)
L5:
	;
	goto L3
L6:
	;
	v40 = v30 & int32(1)
	goto L7
L7:
	;
	v59 = *(*int32)(unsafe.Add(mBase, uint32(l0)+124))
	switch v59 - int32(1) {
	case 0:
		goto L23
	case 1:
		goto L22
	case 2:
		goto L21
	case 3:
		goto L12
	case 4:
		goto L13
	case 5:
		goto L20
	case 6:
		goto L14
	case 7:
		goto L15
	case 8:
		goto L16
	case 9:
		goto L17
	case 10:
		goto L18
	default:
		goto L19
	}
L8:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v1159 = m.ExcPending
	if v1159 != 0 {
		goto L4
	} else {
		goto L302
	}
L9:
	;
	goto L8
L10:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+124)) = int32(4)
	*(*int32)(unsafe.Add(mBase, uint32(l0)+140)) = int32(0)
	goto L7
L11:
	;
	m.G0 = v21 + int32(16)
	return v1113
L12:
	;
	if v40 == int32(0) {
		goto L285
	} else {
		goto L286
	}
L13:
	;
	v816 = *(*int32)(unsafe.Add(mBase, uint32(l0)+144))
	v817 = F_MJEvalInnerValues(m, l0, v816)
	mBase = m.M
	v818 = m.ExcPending
	if v818 != 0 {
		goto L4
	} else {
		goto L235
	}
L14:
	;
	v595 = *(*int32)(unsafe.Add(mBase, uint32(l0)+64))
	v596 = *(*int32)(unsafe.Add(mBase, uint32(v595)+20))
	F_MemoryContextReset(m, v596)
	mBase = m.M
	v598 = m.ExcPending
	if v598 != 0 {
		goto L4
	} else {
		goto L196
	}
L15:
	;
	if v40 == int32(0) {
		goto L179
	} else {
		goto L180
	}
L16:
	;
	if v29&int32(1) == int32(0) {
		goto L158
	} else {
		goto L159
	}
L17:
	;
	v484 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+134)))
	if v484 == int32(0) {
		goto L142
	} else {
		goto L143
	}
L18:
	;
	v459 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+133)))
	if v459 == int32(0) {
		goto L130
	} else {
		goto L131
	}
L19:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v448 = m.ExcPending
	if v448 != 0 {
		goto L4
	} else {
		goto L127
	}
L20:
	;
	if v29&int32(1) == int32(0) {
		goto L81
	} else {
		goto L82
	}
L21:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+124)) = int32(6)
	v125 = *(*int32)(unsafe.Add(mBase, uint32(l0)+136))
	*(*int32)(unsafe.Add(mBase, uint32(v35)+12)) = v125
	v127 = *(*int32)(unsafe.Add(mBase, uint32(l0)+140))
	*(*int32)(unsafe.Add(mBase, uint32(v35)+8)) = v127
	if v32 != 0 {
		goto L59
	} else {
		goto L60
	}
L22:
	;
	v90 = *(*int32)(unsafe.Add(mBase, uint32(v34)+52))
	if v90 != 0 {
		goto L39
	} else {
		goto L40
	}
L23:
	;
	v62 = *(*int32)(unsafe.Add(mBase, uint32(v33)+52))
	if v62 != 0 {
		goto L24
	} else {
		goto L25
	}
L24:
	;
	F_ExecReScan(m, v33)
	mBase = m.M
	v64 = m.ExcPending
	if v64 != 0 {
		goto L4
	} else {
		goto L27
	}
L25:
	;
	goto L26
L26:
	;
	v65 = *(*int32)(unsafe.Add(mBase, uint32(v33)+12))
	v66 = m.T0[v65].(func(*base.Module, int32) int32)(m, v33)
	mBase = m.M
	v67 = m.ExcPending
	if v67 != 0 {
		goto L4
	} else {
		goto L28
	}
L27:
	;
	goto L26
L28:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+136)) = v66
	v69 = F_MJEvalOuterValues(m, l0)
	mBase = m.M
	v70 = m.ExcPending
	if v70 != 0 {
		goto L4
	} else {
		goto L32
	}
L29:
	;
	if v29&int32(1) == int32(0) {
		goto L36
	} else {
		goto L37
	}
L30:
	;
	if v40 == int32(0) {
		goto L7
	} else {
		goto L33
	}
L31:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+124)) = int32(2)
	goto L7
L32:
	;
	switch v69 - int32(1) {
	case 0:
		goto L30
	case 1:
		goto L29
	default:
		goto L31
	}
L33:
	;
	v77 = F_MJFillOuter(m, l0)
	mBase = m.M
	v78 = m.ExcPending
	if v78 != 0 {
		goto L4
	} else {
		goto L34
	}
L34:
	;
	if v77 == int32(0) {
		goto L7
	} else {
		goto L35
	}
L35:
	;
	v1113 = v77
	goto L11
L36:
	;
	v1113 = int32(0)
	goto L11
L37:
	;
	goto L38
L38:
	;
	v86 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(l0)+134)) = uint8(v86)
	*(*int32)(unsafe.Add(mBase, uint32(l0)+124)) = int32(10)
	goto L7
L39:
	;
	F_ExecReScan(m, v34)
	mBase = m.M
	v92 = m.ExcPending
	if v92 != 0 {
		goto L4
	} else {
		goto L42
	}
L40:
	;
	goto L41
L41:
	;
	v93 = *(*int32)(unsafe.Add(mBase, uint32(v34)+12))
	v94 = m.T0[v93].(func(*base.Module, int32) int32)(m, v34)
	mBase = m.M
	v95 = m.ExcPending
	if v95 != 0 {
		goto L4
	} else {
		goto L43
	}
L42:
	;
	goto L41
L43:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+140)) = v94
	v97 = F_MJEvalInnerValues(m, l0, v94)
	mBase = m.M
	v98 = m.ExcPending
	if v98 != 0 {
		goto L4
	} else {
		goto L47
	}
L44:
	;
	if v40 == int32(0) {
		goto L55
	} else {
		goto L56
	}
L45:
	;
	v103 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+129)))
	if v103 == int32(1) {
		goto L48
	} else {
		goto L49
	}
L46:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+124)) = int32(7)
	goto L7
L47:
	;
	switch v97 - int32(1) {
	case 0:
		goto L45
	case 1:
		goto L44
	default:
		goto L46
	}
L48:
	;
	F_ExecMarkPos(m, v34)
	mBase = m.M
	v107 = m.ExcPending
	if v107 != 0 {
		goto L4
	} else {
		goto L51
	}
L49:
	;
	goto L50
L50:
	;
	if v29&int32(1) == int32(0) {
		goto L7
	} else {
		goto L52
	}
L51:
	;
	goto L50
L52:
	;
	v112 = F_MJFillInner(m, l0)
	mBase = m.M
	v113 = m.ExcPending
	if v113 != 0 {
		goto L4
	} else {
		goto L53
	}
L53:
	;
	if v112 == int32(0) {
		goto L7
	} else {
		goto L54
	}
L54:
	;
	v1113 = v112
	goto L11
L55:
	;
	v1113 = int32(0)
	goto L11
L56:
	;
	goto L57
L57:
	;
	v119 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(l0)+133)) = uint8(v119)
	*(*int32)(unsafe.Add(mBase, uint32(l0)+124)) = int32(11)
	goto L7
L58:
	;
	v207 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	if v207 == int32(0) {
		goto L7
	} else {
		goto L80
	}
L59:
	;
	v129 = int32(_a_F_ExecMergeJoin_0)
	v130 = *(*int32)(unsafe.Add(mBase, _c_F_ExecMergeJoin[1]))
	v132 = *(*int32)(unsafe.Add(mBase, uint32(v35)+20))
	*(*int32)(unsafe.Add(mBase, _c_F_ExecMergeJoin[1])) = v132
	v136 = *(*int32)(unsafe.Add(mBase, uint32(v32)+24))
	v137 = m.T0[v136].(func(*base.Module, int32, int32, int32) int64)(m, v32, v35, v21+int32(14))
	mBase = m.M
	v138 = m.ExcPending
	if v138 != 0 {
		goto L4
	} else {
		goto L62
	}
L60:
	;
	goto L61
L61:
	;
	v144 = int32(257)
	*(*uint16)(unsafe.Add(mBase, uint32(l0)+133)) = uint16(v144)
	v146 = *(*int32)(unsafe.Add(mBase, uint32(l0)+104))
	if v146 == int32(5) {
		goto L64
	} else {
		goto L65
	}
L62:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_ExecMergeJoin[1])) = v130
	if v137 == int64(0) {
		goto L58
	} else {
		goto L63
	}
L63:
	;
	goto L61
L64:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+124)) = int32(4)
	goto L7
L65:
	;
	goto L66
L66:
	;
	v151 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+108)))
	if v151 == int32(1) {
		goto L67
	} else {
		goto L68
	}
L67:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+124)) = int32(4)
	goto L69
L68:
	;
	goto L69
L69:
	;
	if v146 == int32(7) {
		goto L7
	} else {
		goto L70
	}
L70:
	;
	if v31 != 0 {
		goto L72
	} else {
		goto L73
	}
L71:
	;
	v200 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	if v200 == int32(0) {
		goto L7
	} else {
		goto L79
	}
L72:
	;
	v158 = int32(_a_F_ExecMergeJoin_0)
	v159 = *(*int32)(unsafe.Add(mBase, _c_F_ExecMergeJoin[1]))
	v161 = *(*int32)(unsafe.Add(mBase, uint32(v35)+20))
	*(*int32)(unsafe.Add(mBase, _c_F_ExecMergeJoin[1])) = v161
	v165 = *(*int32)(unsafe.Add(mBase, uint32(v31)+24))
	v166 = m.T0[v165].(func(*base.Module, int32, int32, int32) int64)(m, v31, v35, v21+int32(15))
	mBase = m.M
	v167 = m.ExcPending
	if v167 != 0 {
		goto L4
	} else {
		goto L75
	}
L73:
	;
	goto L74
L74:
	;
	v173 = *(*int32)(unsafe.Add(mBase, uint32(l0)+68))
	v174 = *(*int32)(unsafe.Add(mBase, uint32(v173)+80))
	v175 = *(*int32)(unsafe.Add(mBase, uint32(v173)+24))
	v176 = *(*int32)(unsafe.Add(mBase, uint32(v175)+8))
	v177 = *(*int32)(unsafe.Add(mBase, uint32(v176)+12))
	m.T0[v177].(func(*base.Module, int32))(m, v175)
	mBase = m.M
	v179 = m.ExcPending
	if v179 != 0 {
		goto L4
	} else {
		goto L77
	}
L75:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_ExecMergeJoin[1])) = v159
	if v166 == int64(0) {
		goto L71
	} else {
		goto L76
	}
L76:
	;
	goto L74
L77:
	;
	v180 = int32(_a_F_ExecMergeJoin_0)
	v181 = *(*int32)(unsafe.Add(mBase, _c_F_ExecMergeJoin[1]))
	v183 = *(*int32)(unsafe.Add(mBase, uint32(v174)+20))
	*(*int32)(unsafe.Add(mBase, _c_F_ExecMergeJoin[1])) = v183
	v188 = *(*int32)(unsafe.Add(mBase, uint32(v173)+32))
	v189 = m.T0[v188].(func(*base.Module, int32, int32, int32) int64)(m, v173+int32(8), v174, int32(0))
	mBase = m.M
	v190 = m.ExcPending
	if v190 != 0 {
		goto L4
	} else {
		goto L78
	}
L78:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_ExecMergeJoin[1])) = v181
	v193 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v175)+4)))
	v195 = v193 & int32(_a_F_ExecMergeJoin_1)
	*(*uint16)(unsafe.Add(mBase, uint32(v175)+4)) = uint16(v195)
	v197 = *(*int32)(unsafe.Add(mBase, uint32(v175)+12))
	v198 = *(*int32)(unsafe.Add(mBase, uint32(v197)))
	*(*uint16)(unsafe.Add(mBase, uint32(v175)+6)) = uint16(v198)
	v1113 = v175
	goto L11
L79:
	;
	v203 = *(*float64)(unsafe.Add(mBase, uint32(v200)+432))
	*(*float64)(unsafe.Add(mBase, uint32(v200)+432)) = base.F64_add(v203, float64(1))
	goto L7
L80:
	;
	v210 = *(*float64)(unsafe.Add(mBase, uint32(v207)+424))
	*(*float64)(unsafe.Add(mBase, uint32(v207)+424)) = base.F64_add(v210, float64(1))
	goto L7
L81:
	;
	v224 = *(*int32)(unsafe.Add(mBase, uint32(v34)+52))
	if v224 != 0 {
		goto L86
	} else {
		goto L87
	}
L82:
	;
	v218 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+134)))
	if v218 != 0 {
		goto L81
	} else {
		goto L83
	}
L83:
	;
	v219 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(l0)+134)) = uint8(v219)
	v221 = F_MJFillInner(m, l0)
	mBase = m.M
	v222 = m.ExcPending
	if v222 != 0 {
		goto L4
	} else {
		goto L84
	}
L84:
	;
	if v221 != 0 {
		v1113 = v221
		goto L11
	} else {
		goto L85
	}
L85:
	;
	goto L81
L86:
	;
	F_ExecReScan(m, v34)
	mBase = m.M
	v226 = m.ExcPending
	if v226 != 0 {
		goto L4
	} else {
		goto L89
	}
L87:
	;
	goto L88
L88:
	;
	v227 = *(*int32)(unsafe.Add(mBase, uint32(v34)+12))
	v228 = m.T0[v227].(func(*base.Module, int32) int32)(m, v34)
	mBase = m.M
	v229 = m.ExcPending
	if v229 != 0 {
		goto L4
	} else {
		goto L90
	}
L89:
	;
	goto L88
L90:
	;
	v230 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(l0)+134)) = uint8(v230)
	*(*int32)(unsafe.Add(mBase, uint32(l0)+140)) = v228
	v233 = F_MJEvalInnerValues(m, l0, v228)
	mBase = m.M
	v234 = m.ExcPending
	if v234 != 0 {
		goto L4
	} else {
		goto L93
	}
L91:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+124)) = int32(4)
	goto L7
L92:
	;
	v237 = *(*int32)(unsafe.Add(mBase, uint32(l0)+64))
	v238 = *(*int32)(unsafe.Add(mBase, uint32(v237)+20))
	F_MemoryContextReset(m, v238)
	mBase = m.M
	v240 = m.ExcPending
	if v240 != 0 {
		goto L4
	} else {
		goto L94
	}
L93:
	;
	switch v233 - int32(1) {
	case 0:
		goto L91
	case 1:
		goto L10
	default:
		goto L92
	}
L94:
	;
	v241 = int32(0)
	v242 = int32(_a_F_ExecMergeJoin_0)
	v243 = *(*int32)(unsafe.Add(mBase, _c_F_ExecMergeJoin[1]))
	v245 = *(*int32)(unsafe.Add(mBase, uint32(v237)+20))
	*(*int32)(unsafe.Add(mBase, _c_F_ExecMergeJoin[1])) = v245
	v248 = *(*int32)(unsafe.Add(mBase, uint32(l0)+116))
	if v241 < v248 {
		goto L97
	} else {
		goto L98
	}
L95:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_ExecMergeJoin[1])) = v243
	if int32(0) <= v373 {
		goto L9
	} else {
		goto L126
	}
L96:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_ExecMergeJoin[1])) = v243
	goto L9
L97:
	;
	v252 = v241
	v257 = v248
	v259 = v241
	goto L100
L98:
	;
	goto L99
L99:
	;
	v396 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+130)))
	if v396 != 0 {
		goto L96
	} else {
		goto L125
	}
L100:
	;
	v269 = *(*int32)(unsafe.Add(mBase, uint32(l0)+120))
	v272 = v269 + v252<<(uint(int32(6))%32)
	v273 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v272)+25)))
	v274 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v272)+24)))
	if v274 != 0 {
		goto L104
	} else {
		goto L105
	}
L101:
	;
	if v322 != 0 {
		goto L96
	} else {
		goto L124
	}
L102:
	;
	v359 = *(*int64)(unsafe.Add(mBase, uint32(v316)+8))
	v360 = *(*int64)(unsafe.Add(mBase, uint32(v316)+16))
	v363 = *(*int32)(unsafe.Add(mBase, uint32(v316)+44))
	v364 = m.T0[v363].(func(*base.Module, int64, int64, int32) int32)(m, v359, v360, v316+int32(28))
	mBase = m.M
	v365 = m.ExcPending
	if v365 != 0 {
		goto L4
	} else {
		goto L117
	}
L103:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_ExecMergeJoin[1])) = v243
	*(*int32)(unsafe.Add(mBase, uint32(l0)+124)) = int32(4)
	goto L7
L104:
	;
	v276 = v252
	v278 = v273
	goto L107
L105:
	;
	v315 = v252
	v316 = v272
	v317 = v273
	v322 = v259
	goto L106
L106:
	;
	if v317&int32(1) == int32(0) {
		goto L102
	} else {
		goto L115
	}
L107:
	;
	if v278&int32(1) == int32(0) {
		goto L109
	} else {
		goto L110
	}
L108:
	;
	v315 = v304
	v316 = v309
	v317 = v310
	v322 = v306
	goto L106
L109:
	;
	v300 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v269+v276<<(uint(int32(6))%32))+37)))
	if v300 == int32(0) {
		goto L96
	} else {
		goto L112
	}
L110:
	;
	goto L111
L111:
	;
	v304 = v276 + int32(1)
	if v257 <= v304 {
		goto L96
	} else {
		goto L113
	}
L112:
	;
	goto L103
L113:
	;
	v306 = int32(1)
	v309 = v269 + v304<<(uint(int32(6))%32)
	v310 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v309)+25)))
	v311 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v309)+24)))
	if v311 == v306 {
		v276 = v304
		v278 = v310
		goto L107
	} else {
		goto L114
	}
L114:
	;
	goto L108
L115:
	;
	v336 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v316)+37)))
	if v336 != 0 {
		goto L96
	} else {
		goto L116
	}
L116:
	;
	goto L103
L117:
	;
	v366 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v316)+36)))
	if v366 == int32(1) {
		goto L118
	} else {
		goto L119
	}
L118:
	;
	if v364 < int32(0) {
		goto L96
	} else {
		goto L121
	}
L119:
	;
	v373 = v364
	goto L120
L120:
	;
	if v373 != 0 {
		goto L95
	} else {
		goto L122
	}
L121:
	;
	v373 = int32(0) - v364
	goto L120
L122:
	;
	v375 = v315 + int32(1)
	v376 = *(*int32)(unsafe.Add(mBase, uint32(l0)+116))
	if v375 < v376 {
		v252 = v375
		v257 = v376
		v259 = v322
		goto L100
	} else {
		goto L123
	}
L123:
	;
	goto L101
L124:
	;
	goto L99
L125:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_ExecMergeJoin[1])) = v243
	*(*int32)(unsafe.Add(mBase, uint32(l0)+124)) = int32(3)
	goto L7
L126:
	;
	goto L91
L127:
	;
	v449 = *(*int32)(unsafe.Add(mBase, uint32(l0)+124))
	*(*int32)(unsafe.Add(mBase, uint32(v21))) = v449
	F_errmsg_internal(m, int32(_a_F_ExecMergeJoin_2), v21)
	mBase = m.M
	v453 = m.ExcPending
	if v453 != 0 {
		goto L4
	} else {
		goto L128
	}
L128:
	;
	F_errfinish(m, int32(_a_F_ExecMergeJoin_3), int32(1431), int32(_a_F_ExecMergeJoin_4))
	mBase = m.M
	v458 = m.ExcPending
	if v458 != 0 {
		goto L4
	} else {
		goto L129
	}
L129:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L130:
	;
	v462 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(l0)+133)) = uint8(v462)
	v464 = F_MJFillOuter(m, l0)
	mBase = m.M
	v465 = m.ExcPending
	if v465 != 0 {
		goto L4
	} else {
		goto L133
	}
L131:
	;
	goto L132
L132:
	;
	v467 = *(*int32)(unsafe.Add(mBase, uint32(v33)+52))
	if v467 != 0 {
		goto L135
	} else {
		goto L136
	}
L133:
	;
	if v464 != 0 {
		v1113 = v464
		goto L11
	} else {
		goto L134
	}
L134:
	;
	goto L132
L135:
	;
	F_ExecReScan(m, v33)
	mBase = m.M
	v469 = m.ExcPending
	if v469 != 0 {
		goto L4
	} else {
		goto L138
	}
L136:
	;
	goto L137
L137:
	;
	v470 = int32(0)
	v471 = *(*int32)(unsafe.Add(mBase, uint32(v33)+12))
	v472 = m.T0[v471].(func(*base.Module, int32) int32)(m, v33)
	mBase = m.M
	v473 = m.ExcPending
	if v473 != 0 {
		goto L4
	} else {
		goto L139
	}
L138:
	;
	goto L137
L139:
	;
	v474 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(l0)+133)) = uint8(v474)
	*(*int32)(unsafe.Add(mBase, uint32(l0)+136)) = v472
	if v472 == v474 {
		v1113 = v470
		goto L11
	} else {
		goto L140
	}
L140:
	;
	v479 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v472)+4)))
	if v479&int32(2) == int32(0) {
		goto L7
	} else {
		goto L141
	}
L141:
	;
	v1113 = v470
	goto L11
L142:
	;
	v487 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(l0)+134)) = uint8(v487)
	v489 = F_MJFillInner(m, l0)
	mBase = m.M
	v490 = m.ExcPending
	if v490 != 0 {
		goto L4
	} else {
		goto L145
	}
L143:
	;
	goto L144
L144:
	;
	v492 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+129)))
	if v492 == int32(1) {
		goto L147
	} else {
		goto L148
	}
L145:
	;
	if v489 != 0 {
		v1113 = v489
		goto L11
	} else {
		goto L146
	}
L146:
	;
	goto L144
L147:
	;
	F_ExecMarkPos(m, v34)
	mBase = m.M
	v496 = m.ExcPending
	if v496 != 0 {
		goto L4
	} else {
		goto L150
	}
L148:
	;
	goto L149
L149:
	;
	v497 = *(*int32)(unsafe.Add(mBase, uint32(v34)+52))
	if v497 != 0 {
		goto L151
	} else {
		goto L152
	}
L150:
	;
	goto L149
L151:
	;
	F_ExecReScan(m, v34)
	mBase = m.M
	v499 = m.ExcPending
	if v499 != 0 {
		goto L4
	} else {
		goto L154
	}
L152:
	;
	goto L153
L153:
	;
	v500 = int32(0)
	v501 = *(*int32)(unsafe.Add(mBase, uint32(v34)+12))
	v502 = m.T0[v501].(func(*base.Module, int32) int32)(m, v34)
	mBase = m.M
	v503 = m.ExcPending
	if v503 != 0 {
		goto L4
	} else {
		goto L155
	}
L154:
	;
	goto L153
L155:
	;
	v504 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(l0)+134)) = uint8(v504)
	*(*int32)(unsafe.Add(mBase, uint32(l0)+140)) = v502
	if v502 == v504 {
		v1113 = v500
		goto L11
	} else {
		goto L156
	}
L156:
	;
	v509 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v502)+4)))
	if v509&int32(2) == int32(0) {
		goto L7
	} else {
		goto L157
	}
L157:
	;
	v1113 = v500
	goto L11
L158:
	;
	v524 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+129)))
	if v524 == int32(1) {
		goto L163
	} else {
		goto L164
	}
L159:
	;
	v518 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+134)))
	if v518 != 0 {
		goto L158
	} else {
		goto L160
	}
L160:
	;
	v519 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(l0)+134)) = uint8(v519)
	v521 = F_MJFillInner(m, l0)
	mBase = m.M
	v522 = m.ExcPending
	if v522 != 0 {
		goto L4
	} else {
		goto L161
	}
L161:
	;
	if v521 != 0 {
		v1113 = v521
		goto L11
	} else {
		goto L162
	}
L162:
	;
	goto L158
L163:
	;
	F_ExecMarkPos(m, v34)
	mBase = m.M
	v528 = m.ExcPending
	if v528 != 0 {
		goto L4
	} else {
		goto L166
	}
L164:
	;
	goto L165
L165:
	;
	v529 = *(*int32)(unsafe.Add(mBase, uint32(v34)+52))
	if v529 != 0 {
		goto L167
	} else {
		goto L168
	}
L166:
	;
	goto L165
L167:
	;
	F_ExecReScan(m, v34)
	mBase = m.M
	v531 = m.ExcPending
	if v531 != 0 {
		goto L4
	} else {
		goto L170
	}
L168:
	;
	goto L169
L169:
	;
	v532 = *(*int32)(unsafe.Add(mBase, uint32(v34)+12))
	v533 = m.T0[v532].(func(*base.Module, int32) int32)(m, v34)
	mBase = m.M
	v534 = m.ExcPending
	if v534 != 0 {
		goto L4
	} else {
		goto L171
	}
L170:
	;
	goto L169
L171:
	;
	v535 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(l0)+134)) = uint8(v535)
	*(*int32)(unsafe.Add(mBase, uint32(l0)+140)) = v533
	v538 = F_MJEvalInnerValues(m, l0, v533)
	mBase = m.M
	v539 = m.ExcPending
	if v539 != 0 {
		goto L4
	} else {
		goto L175
	}
L172:
	;
	v546 = int32(0)
	if v40 == v546 {
		v1113 = v546
		goto L11
	} else {
		goto L176
	}
L173:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+124)) = int32(9)
	goto L7
L174:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+124)) = int32(7)
	goto L7
L175:
	;
	switch v538 - int32(1) {
	case 0:
		goto L173
	case 1:
		goto L172
	default:
		goto L174
	}
L176:
	;
	v549 = *(*int32)(unsafe.Add(mBase, uint32(l0)+136))
	if v549 == int32(0) {
		v1113 = v546
		goto L11
	} else {
		goto L177
	}
L177:
	;
	v552 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v549)+4)))
	if v552&int32(2) != 0 {
		v1113 = v546
		goto L11
	} else {
		goto L178
	}
L178:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+124)) = int32(11)
	goto L7
L179:
	;
	v565 = *(*int32)(unsafe.Add(mBase, uint32(v33)+52))
	if v565 != 0 {
		goto L184
	} else {
		goto L185
	}
L180:
	;
	v559 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+133)))
	if v559 != 0 {
		goto L179
	} else {
		goto L181
	}
L181:
	;
	v560 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(l0)+133)) = uint8(v560)
	v562 = F_MJFillOuter(m, l0)
	mBase = m.M
	v563 = m.ExcPending
	if v563 != 0 {
		goto L4
	} else {
		goto L182
	}
L182:
	;
	if v562 != 0 {
		v1113 = v562
		goto L11
	} else {
		goto L183
	}
L183:
	;
	goto L179
L184:
	;
	F_ExecReScan(m, v33)
	mBase = m.M
	v567 = m.ExcPending
	if v567 != 0 {
		goto L4
	} else {
		goto L187
	}
L185:
	;
	goto L186
L186:
	;
	v568 = *(*int32)(unsafe.Add(mBase, uint32(v33)+12))
	v569 = m.T0[v568].(func(*base.Module, int32) int32)(m, v33)
	mBase = m.M
	v570 = m.ExcPending
	if v570 != 0 {
		goto L4
	} else {
		goto L188
	}
L187:
	;
	goto L186
L188:
	;
	v571 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(l0)+133)) = uint8(v571)
	*(*int32)(unsafe.Add(mBase, uint32(l0)+136)) = v569
	v574 = F_MJEvalOuterValues(m, l0)
	mBase = m.M
	v575 = m.ExcPending
	if v575 != 0 {
		goto L4
	} else {
		goto L192
	}
L189:
	;
	v582 = int32(0)
	if v29&int32(1) == v582 {
		v1113 = v582
		goto L11
	} else {
		goto L193
	}
L190:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+124)) = int32(8)
	goto L7
L191:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+124)) = int32(7)
	goto L7
L192:
	;
	switch v574 - int32(1) {
	case 0:
		goto L190
	case 1:
		goto L189
	default:
		goto L191
	}
L193:
	;
	v587 = *(*int32)(unsafe.Add(mBase, uint32(l0)+140))
	if v587 == int32(0) {
		v1113 = v582
		goto L11
	} else {
		goto L194
	}
L194:
	;
	v590 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v587)+4)))
	if v590&int32(2) != 0 {
		v1113 = v582
		goto L11
	} else {
		goto L195
	}
L195:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+124)) = int32(10)
	goto L7
L196:
	;
	v599 = int32(0)
	v600 = int32(_a_F_ExecMergeJoin_0)
	v601 = *(*int32)(unsafe.Add(mBase, _c_F_ExecMergeJoin[1]))
	v603 = *(*int32)(unsafe.Add(mBase, uint32(v595)+20))
	*(*int32)(unsafe.Add(mBase, _c_F_ExecMergeJoin[1])) = v603
	v606 = *(*int32)(unsafe.Add(mBase, uint32(l0)+116))
	if v599 < v606 {
		goto L200
	} else {
		goto L201
	}
L197:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+124)) = int32(9)
	goto L7
L198:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_ExecMergeJoin[1])) = v601
	if int32(0) <= v731 {
		goto L197
	} else {
		goto L234
	}
L199:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_ExecMergeJoin[1])) = v601
	goto L197
L200:
	;
	v610 = v599
	v616 = v606
	v617 = v599
	goto L203
L201:
	;
	goto L202
L202:
	;
	v754 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+130)))
	if v754 != 0 {
		goto L199
	} else {
		goto L228
	}
L203:
	;
	v627 = *(*int32)(unsafe.Add(mBase, uint32(l0)+120))
	v630 = v627 + v610<<(uint(int32(6))%32)
	v631 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v630)+25)))
	v632 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v630)+24)))
	if v632 != 0 {
		goto L207
	} else {
		goto L208
	}
L204:
	;
	if v680 != 0 {
		goto L199
	} else {
		goto L227
	}
L205:
	;
	v717 = *(*int64)(unsafe.Add(mBase, uint32(v674)+8))
	v718 = *(*int64)(unsafe.Add(mBase, uint32(v674)+16))
	v721 = *(*int32)(unsafe.Add(mBase, uint32(v674)+44))
	v722 = m.T0[v721].(func(*base.Module, int64, int64, int32) int32)(m, v717, v718, v674+int32(28))
	mBase = m.M
	v723 = m.ExcPending
	if v723 != 0 {
		goto L4
	} else {
		goto L220
	}
L206:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_ExecMergeJoin[1])) = v601
	*(*int32)(unsafe.Add(mBase, uint32(l0)+124)) = int32(8)
	goto L7
L207:
	;
	v634 = v610
	v636 = v631
	goto L210
L208:
	;
	v673 = v610
	v674 = v630
	v675 = v631
	v680 = v617
	goto L209
L209:
	;
	if v675&int32(1) == int32(0) {
		goto L205
	} else {
		goto L218
	}
L210:
	;
	if v636&int32(1) == int32(0) {
		goto L212
	} else {
		goto L213
	}
L211:
	;
	v673 = v662
	v674 = v667
	v675 = v668
	v680 = v664
	goto L209
L212:
	;
	v658 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v627+v634<<(uint(int32(6))%32))+37)))
	if v658 == int32(0) {
		goto L199
	} else {
		goto L215
	}
L213:
	;
	goto L214
L214:
	;
	v662 = v634 + int32(1)
	if v616 <= v662 {
		goto L199
	} else {
		goto L216
	}
L215:
	;
	goto L206
L216:
	;
	v664 = int32(1)
	v667 = v627 + v662<<(uint(int32(6))%32)
	v668 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v667)+25)))
	v669 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v667)+24)))
	if v669 == v664 {
		v634 = v662
		v636 = v668
		goto L210
	} else {
		goto L217
	}
L217:
	;
	goto L211
L218:
	;
	v694 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v674)+37)))
	if v694 != 0 {
		goto L199
	} else {
		goto L219
	}
L219:
	;
	goto L206
L220:
	;
	v724 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v674)+36)))
	if v724 == int32(1) {
		goto L221
	} else {
		goto L222
	}
L221:
	;
	if v722 < int32(0) {
		goto L199
	} else {
		goto L224
	}
L222:
	;
	v731 = v722
	goto L223
L223:
	;
	if v731 != 0 {
		goto L198
	} else {
		goto L225
	}
L224:
	;
	v731 = int32(0) - v722
	goto L223
L225:
	;
	v733 = v673 + int32(1)
	v734 = *(*int32)(unsafe.Add(mBase, uint32(l0)+116))
	if v733 < v734 {
		v610 = v733
		v616 = v734
		v617 = v680
		goto L203
	} else {
		goto L226
	}
L226:
	;
	goto L204
L227:
	;
	goto L202
L228:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_ExecMergeJoin[1])) = v601
	v757 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+128)))
	if v757 == int32(0) {
		goto L229
	} else {
		goto L230
	}
L229:
	;
	F_ExecMarkPos(m, v34)
	mBase = m.M
	v761 = m.ExcPending
	if v761 != 0 {
		goto L4
	} else {
		goto L232
	}
L230:
	;
	goto L231
L231:
	;
	v762 = *(*int32)(unsafe.Add(mBase, uint32(l0)+144))
	v763 = *(*int32)(unsafe.Add(mBase, uint32(l0)+140))
	v764 = *(*int32)(unsafe.Add(mBase, uint32(v762)+8))
	v765 = *(*int32)(unsafe.Add(mBase, uint32(v764)+32))
	m.T0[v765].(func(*base.Module, int32, int32))(m, v762, v763)
	mBase = m.M
	v767 = m.ExcPending
	if v767 != 0 {
		goto L4
	} else {
		goto L233
	}
L232:
	;
	goto L231
L233:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+124)) = int32(3)
	goto L7
L234:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+124)) = int32(8)
	goto L7
L235:
	;
	v819 = *(*int32)(unsafe.Add(mBase, uint32(l0)+64))
	v820 = *(*int32)(unsafe.Add(mBase, uint32(v819)+20))
	F_MemoryContextReset(m, v820)
	mBase = m.M
	v822 = m.ExcPending
	if v822 != 0 {
		goto L4
	} else {
		goto L236
	}
L236:
	;
	v823 = int32(0)
	v824 = int32(_a_F_ExecMergeJoin_0)
	v825 = *(*int32)(unsafe.Add(mBase, _c_F_ExecMergeJoin[1]))
	v827 = *(*int32)(unsafe.Add(mBase, uint32(v819)+20))
	*(*int32)(unsafe.Add(mBase, _c_F_ExecMergeJoin[1])) = v827
	v830 = *(*int32)(unsafe.Add(mBase, uint32(l0)+116))
	if v823 < v830 {
		goto L241
	} else {
		goto L242
	}
L237:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v1064 = m.ExcPending
	if v1064 != 0 {
		goto L4
	} else {
		goto L282
	}
L238:
	;
	v1029 = *(*int32)(unsafe.Add(mBase, uint32(l0)+140))
	v1030 = F_MJEvalInnerValues(m, l0, v1029)
	mBase = m.M
	v1031 = m.ExcPending
	if v1031 != 0 {
		goto L4
	} else {
		goto L278
	}
L239:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_ExecMergeJoin[1])) = v825
	if v953 <= int32(0) {
		goto L237
	} else {
		goto L274
	}
L240:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_ExecMergeJoin[1])) = v825
	goto L238
L241:
	;
	v834 = v823
	v840 = v830
	v841 = v823
	goto L244
L242:
	;
	goto L243
L243:
	;
	v976 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+130)))
	if v976 != 0 {
		goto L240
	} else {
		goto L269
	}
L244:
	;
	v851 = *(*int32)(unsafe.Add(mBase, uint32(l0)+120))
	v854 = v851 + v834<<(uint(int32(6))%32)
	v855 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v854)+25)))
	v856 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v854)+24)))
	if v856 != 0 {
		goto L248
	} else {
		goto L249
	}
L245:
	;
	if v904 != 0 {
		goto L240
	} else {
		goto L268
	}
L246:
	;
	v939 = *(*int64)(unsafe.Add(mBase, uint32(v898)+8))
	v940 = *(*int64)(unsafe.Add(mBase, uint32(v898)+16))
	v943 = *(*int32)(unsafe.Add(mBase, uint32(v898)+44))
	v944 = m.T0[v943].(func(*base.Module, int64, int64, int32) int32)(m, v939, v940, v898+int32(28))
	mBase = m.M
	v945 = m.ExcPending
	if v945 != 0 {
		goto L4
	} else {
		goto L261
	}
L247:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_ExecMergeJoin[1])) = v825
	goto L237
L248:
	;
	v858 = v834
	v860 = v855
	goto L251
L249:
	;
	v897 = v834
	v898 = v854
	v899 = v855
	v904 = v841
	goto L250
L250:
	;
	if v899&int32(1) == int32(0) {
		goto L246
	} else {
		goto L259
	}
L251:
	;
	if v860&int32(1) == int32(0) {
		goto L253
	} else {
		goto L254
	}
L252:
	;
	v897 = v886
	v898 = v891
	v899 = v892
	v904 = v888
	goto L250
L253:
	;
	v882 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v851+v858<<(uint(int32(6))%32))+37)))
	if v882 == int32(0) {
		goto L240
	} else {
		goto L256
	}
L254:
	;
	goto L255
L255:
	;
	v886 = v858 + int32(1)
	if v840 <= v886 {
		goto L240
	} else {
		goto L257
	}
L256:
	;
	goto L247
L257:
	;
	v888 = int32(1)
	v891 = v851 + v886<<(uint(int32(6))%32)
	v892 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v891)+25)))
	v893 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v891)+24)))
	if v893 == v888 {
		v858 = v886
		v860 = v892
		goto L251
	} else {
		goto L258
	}
L258:
	;
	goto L252
L259:
	;
	v918 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v898)+37)))
	if v918 != 0 {
		goto L240
	} else {
		goto L260
	}
L260:
	;
	goto L247
L261:
	;
	v946 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v898)+36)))
	if v946 == int32(1) {
		goto L262
	} else {
		goto L263
	}
L262:
	;
	if v944 < int32(0) {
		goto L240
	} else {
		goto L265
	}
L263:
	;
	v953 = v944
	goto L264
L264:
	;
	if v953 != 0 {
		goto L239
	} else {
		goto L266
	}
L265:
	;
	v953 = int32(0) - v944
	goto L264
L266:
	;
	v955 = v897 + int32(1)
	v956 = *(*int32)(unsafe.Add(mBase, uint32(l0)+116))
	if v955 < v956 {
		v834 = v955
		v840 = v956
		v841 = v904
		goto L244
	} else {
		goto L267
	}
L267:
	;
	goto L245
L268:
	;
	goto L243
L269:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_ExecMergeJoin[1])) = v825
	v979 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+128)))
	if v979 == int32(0) {
		goto L270
	} else {
		goto L271
	}
L270:
	;
	F_ExecRestrPos(m, v34)
	mBase = m.M
	v983 = m.ExcPending
	if v983 != 0 {
		goto L4
	} else {
		goto L273
	}
L271:
	;
	goto L272
L272:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+124)) = int32(3)
	goto L7
L273:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+140)) = v816
	goto L272
L274:
	;
	goto L238
L275:
	;
	if v40 == int32(0) {
		goto L279
	} else {
		goto L280
	}
L276:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+124)) = int32(9)
	goto L7
L277:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+124)) = int32(7)
	goto L7
L278:
	;
	switch v1030 - int32(1) {
	case 0:
		goto L276
	case 1:
		goto L275
	default:
		goto L277
	}
L279:
	;
	v1113 = int32(0)
	goto L11
L280:
	;
	goto L281
L281:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+124)) = int32(11)
	goto L7
L282:
	;
	F_errmsg_internal(m, int32(_a_F_ExecMergeJoin_5), int32(0))
	mBase = m.M
	v1068 = m.ExcPending
	if v1068 != 0 {
		goto L4
	} else {
		goto L283
	}
L283:
	;
	F_errfinish(m, int32(_a_F_ExecMergeJoin_3), int32(1147), int32(_a_F_ExecMergeJoin_4))
	mBase = m.M
	v1073 = m.ExcPending
	if v1073 != 0 {
		goto L4
	} else {
		goto L284
	}
L284:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L285:
	;
	v1082 = *(*int32)(unsafe.Add(mBase, uint32(v33)+52))
	if v1082 != 0 {
		goto L290
	} else {
		goto L291
	}
L286:
	;
	v1076 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+133)))
	if v1076 != 0 {
		goto L285
	} else {
		goto L287
	}
L287:
	;
	v1077 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(l0)+133)) = uint8(v1077)
	v1079 = F_MJFillOuter(m, l0)
	mBase = m.M
	v1080 = m.ExcPending
	if v1080 != 0 {
		goto L4
	} else {
		goto L288
	}
L288:
	;
	if v1079 != 0 {
		v1113 = v1079
		goto L11
	} else {
		goto L289
	}
L289:
	;
	goto L285
L290:
	;
	F_ExecReScan(m, v33)
	mBase = m.M
	v1084 = m.ExcPending
	if v1084 != 0 {
		goto L4
	} else {
		goto L293
	}
L291:
	;
	goto L292
L292:
	;
	v1085 = *(*int32)(unsafe.Add(mBase, uint32(v33)+12))
	v1086 = m.T0[v1085].(func(*base.Module, int32) int32)(m, v33)
	mBase = m.M
	v1087 = m.ExcPending
	if v1087 != 0 {
		goto L4
	} else {
		goto L294
	}
L293:
	;
	goto L292
L294:
	;
	v1088 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(l0)+133)) = uint8(v1088)
	*(*int32)(unsafe.Add(mBase, uint32(l0)+136)) = v1086
	v1091 = F_MJEvalOuterValues(m, l0)
	mBase = m.M
	v1092 = m.ExcPending
	if v1092 != 0 {
		goto L4
	} else {
		goto L298
	}
L295:
	;
	v1099 = int32(0)
	if v29&int32(1) == v1099 {
		v1113 = v1099
		goto L11
	} else {
		goto L299
	}
L296:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+124)) = int32(4)
	goto L7
L297:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+124)) = int32(5)
	goto L7
L298:
	;
	switch v1091 - int32(1) {
	case 0:
		goto L296
	case 1:
		goto L295
	default:
		goto L297
	}
L299:
	;
	v1104 = *(*int32)(unsafe.Add(mBase, uint32(l0)+140))
	if v1104 == int32(0) {
		v1113 = v1099
		goto L11
	} else {
		goto L300
	}
L300:
	;
	v1107 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1104)+4)))
	if v1107&int32(2) != 0 {
		v1113 = v1099
		goto L11
	} else {
		goto L301
	}
L301:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+124)) = int32(10)
	goto L7
L302:
	;
	F_errmsg_internal(m, int32(_a_F_ExecMergeJoin_5), int32(0))
	mBase = m.M
	v1163 = m.ExcPending
	if v1163 != 0 {
		goto L4
	} else {
		goto L303
	}
L303:
	;
	F_errfinish(m, int32(_a_F_ExecMergeJoin_3), int32(904), int32(_a_F_ExecMergeJoin_4))
	mBase = m.M
	v1168 = m.ExcPending
	if v1168 != 0 {
		goto L4
	} else {
		goto L304
	}
L304:
	;
	base.Wasm_trap_unreachable()
	for {
	}
}
func F_create_join_clause(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32, l5 int32) int32 {
	mBase := m.M
	_ = mBase
	var v7 int32
	_ = v7
	var v14 int32
	_ = v14
	var v17 int32
	_ = v17
	var v20 int32
	_ = v20
	var v28 int32
	_ = v28
	var v37 int32
	_ = v37
	var v38 int32
	_ = v38
	var v40 int32
	_ = v40
	var v42 int32
	_ = v42
	var v45 int32
	_ = v45
	var v47 int32
	_ = v47
	var v50 int32
	_ = v50
	var v65 int32
	_ = v65
	var v68 int32
	_ = v68
	var v69 int32
	_ = v69
	var v70 int32
	_ = v70
	var v71 int32
	_ = v71
	var v73 int32
	_ = v73
	var v75 int32
	_ = v75
	var v78 int32
	_ = v78
	var v81 int32
	_ = v81
	var v82 int32
	_ = v82
	var v83 int32
	_ = v83
	var v84 int32
	_ = v84
	var v85 int32
	_ = v85
	var v86 int32
	_ = v86
	var v88 int32
	_ = v88
	var v89 int32
	_ = v89
	var v90 int32
	_ = v90
	var v91 int32
	_ = v91
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
	var v104 int32
	_ = v104
	var v105 int32
	_ = v105
	var v107 int32
	_ = v107
	var v110 int32
	_ = v110
	var v111 int32
	_ = v111
	var v112 int32
	_ = v112
	var v113 int32
	_ = v113
	var v116 int32
	_ = v116
	var v117 int32
	_ = v117
	var v120 int32
	_ = v120
	var v123 int32
	_ = v123
	var v133 int32
	_ = v133
	var v141 int32
	_ = v141
	var v142 int32
	_ = v142
	var v144 int32
	_ = v144
	var v146 int32
	_ = v146
	var v149 int32
	_ = v149
	var v151 int32
	_ = v151
	var v154 int32
	_ = v154
	var v169 int32
	_ = v169
	var v170 int32
	_ = v170
	var v180 int32
	_ = v180
	var v186 int32
	_ = v186
	var v206 int32
	_ = v206
	var v207 int32
	_ = v207
	var v208 int32
	_ = v208
	var v210 int32
	_ = v210
	var v212 int32
	_ = v212
	var v221 int32
	_ = v221
	v7 = int32(0)
	v14 = *(*int32)(unsafe.Add(mBase, uint32(l1)+24))
	if v14 == v7 {
		goto L2
	} else {
		goto L3
	}
L1:
	;
	return v221
L2:
	;
	v65 = F_ec_search_derived_clause_for_ems(m, l0, l1, l3, l4, l5)
	mBase = m.M
	v68 = m.ExcPending
	if v68 != 0 {
		goto L16
	} else {
		goto L17
	}
L3:
	;
	v17 = *(*int32)(unsafe.Add(mBase, uint32(v14)+4))
	if v17 <= int32(0) {
		goto L2
	} else {
		goto L4
	}
L4:
	;
	v20 = *(*int32)(unsafe.Add(mBase, uint32(v14)+12))
	v28 = v7
	goto L5
L5:
	;
	v37 = *(*int32)(unsafe.Add(mBase, uint32(v20+v28<<(uint(int32(2))%32))))
	v38 = *(*int32)(unsafe.Add(mBase, uint32(v37)+108))
	if v38 != l3 {
		goto L7
	} else {
		goto L8
	}
L6:
	;
	goto L2
L7:
	;
	if l4 != v38 {
		goto L11
	} else {
		goto L12
	}
L8:
	;
	v40 = *(*int32)(unsafe.Add(mBase, uint32(v37)+112))
	if v40 != l4 {
		goto L7
	} else {
		goto L9
	}
L9:
	;
	v42 = *(*int32)(unsafe.Add(mBase, uint32(v37)+60))
	if v42 == l5 {
		v221 = v37
		goto L1
	} else {
		goto L10
	}
L10:
	;
	goto L7
L11:
	;
	v50 = v28 + int32(1)
	if v17 != v50 {
		v28 = v50
		goto L5
	} else {
		goto L15
	}
L12:
	;
	v45 = *(*int32)(unsafe.Add(mBase, uint32(v37)+112))
	if v45 != l3 {
		goto L11
	} else {
		goto L13
	}
L13:
	;
	v47 = *(*int32)(unsafe.Add(mBase, uint32(v37)+60))
	if v47 == l5 {
		v221 = v37
		goto L1
	} else {
		goto L14
	}
L14:
	;
	goto L11
L15:
	;
	goto L6
L16:
	;
	return int32(0)
L17:
	;
	if v65 != 0 {
		v221 = v65
		goto L1
	} else {
		goto L18
	}
L18:
	;
	v69 = int32(0)
	v70 = int32(_a_F_create_join_clause_0)
	v71 = *(*int32)(unsafe.Add(mBase, _c_F_create_join_clause[0]))
	v73 = *(*int32)(unsafe.Add(mBase, uint32(l0)+300))
	*(*int32)(unsafe.Add(mBase, _c_F_create_join_clause[0])) = v73
	v75 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l3)+13)))
	if v75 == v69 {
		goto L20
	} else {
		goto L21
	}
L19:
	;
	v89 = *(*int32)(unsafe.Add(mBase, uint32(l1)+8))
	v90 = *(*int32)(unsafe.Add(mBase, uint32(l3)+4))
	v91 = *(*int32)(unsafe.Add(mBase, uint32(l4)+4))
	v92 = *(*int32)(unsafe.Add(mBase, uint32(l3)+8))
	v93 = *(*int32)(unsafe.Add(mBase, uint32(l4)+8))
	v94 = F_bms_union(m, v92, v93)
	mBase = m.M
	v95 = m.ExcPending
	if v95 != 0 {
		goto L16
	} else {
		goto L31
	}
L20:
	;
	v78 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l4)+13)))
	if v78 != int32(1) {
		v88 = v69
		goto L19
	} else {
		goto L23
	}
L21:
	;
	goto L22
L22:
	;
	v81 = *(*int32)(unsafe.Add(mBase, uint32(l3)+24))
	if v81 != 0 {
		goto L24
	} else {
		goto L25
	}
L23:
	;
	goto L22
L24:
	;
	v82 = v81
	goto L26
L25:
	;
	v82 = l3
	goto L26
L26:
	;
	v83 = *(*int32)(unsafe.Add(mBase, uint32(l4)+24))
	if v83 != 0 {
		goto L27
	} else {
		goto L28
	}
L27:
	;
	v84 = v83
	goto L29
L28:
	;
	v84 = l4
	goto L29
L29:
	;
	v85 = F_create_join_clause(m, l0, l1, l2, v82, v84, l5)
	mBase = m.M
	v86 = m.ExcPending
	if v86 != 0 {
		goto L16
	} else {
		goto L30
	}
L30:
	;
	v88 = v85
	goto L19
L31:
	;
	v96 = *(*int32)(unsafe.Add(mBase, uint32(l1)+48))
	v97 = F_build_implied_join_equality(m, l0, l2, v89, v90, v91, v94, v96)
	mBase = m.M
	v98 = m.ExcPending
	if v98 != 0 {
		goto L16
	} else {
		goto L32
	}
L32:
	;
	v99 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l3)+13)))
	if v99 == int32(1) {
		goto L33
	} else {
		goto L34
	}
L33:
	;
	v102 = *(*int32)(unsafe.Add(mBase, uint32(v97)+28))
	v103 = *(*int32)(unsafe.Add(mBase, uint32(l3)+8))
	v104 = F_bms_add_members(m, v102, v103)
	mBase = m.M
	v105 = m.ExcPending
	if v105 != 0 {
		goto L16
	} else {
		goto L36
	}
L34:
	;
	goto L35
L35:
	;
	v107 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l4)+13)))
	if v107 == int32(1) {
		goto L37
	} else {
		goto L38
	}
L36:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v97)+28)) = v104
	goto L35
L37:
	;
	v110 = *(*int32)(unsafe.Add(mBase, uint32(v97)+28))
	v111 = *(*int32)(unsafe.Add(mBase, uint32(l4)+8))
	v112 = F_bms_add_members(m, v110, v111)
	mBase = m.M
	v113 = m.ExcPending
	if v113 != 0 {
		goto L16
	} else {
		goto L40
	}
L38:
	;
	goto L39
L39:
	;
	if v88 != 0 {
		v180 = v88
		goto L42
	} else {
		goto L43
	}
L40:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v97)+28)) = v112
	goto L39
L41:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v97)+112)) = l4
	*(*int32)(unsafe.Add(mBase, uint32(v97)+108)) = l3
	*(*int32)(unsafe.Add(mBase, uint32(v97)+104)) = l1
	*(*int32)(unsafe.Add(mBase, uint32(v97)+100)) = l1
	*(*int32)(unsafe.Add(mBase, uint32(v97)+60)) = l5
	v206 = *(*int32)(unsafe.Add(mBase, uint32(l1)+28))
	v207 = F_lappend(m, v206, v97)
	mBase = m.M
	v208 = m.ExcPending
	if v208 != 0 {
		goto L16
	} else {
		goto L63
	}
L42:
	;
	v186 = *(*int32)(unsafe.Add(mBase, uint32(v180)+56))
	*(*int32)(unsafe.Add(mBase, uint32(v97)+56)) = v186
	goto L41
L43:
	;
	if l5 != 0 {
		goto L44
	} else {
		goto L45
	}
L44:
	;
	v116 = int32(0)
	goto L46
L45:
	;
	v116 = l1
	goto L46
L46:
	;
	v117 = *(*int32)(unsafe.Add(mBase, uint32(l1)+24))
	if v117 == int32(0) {
		goto L47
	} else {
		goto L48
	}
L47:
	;
	v169 = F_ec_search_derived_clause_for_ems(m, l0, l1, l3, l4, v116)
	mBase = m.M
	v170 = m.ExcPending
	if v170 != 0 {
		goto L16
	} else {
		goto L61
	}
L48:
	;
	v120 = *(*int32)(unsafe.Add(mBase, uint32(v117)+4))
	if v120 <= int32(0) {
		goto L47
	} else {
		goto L49
	}
L49:
	;
	v123 = *(*int32)(unsafe.Add(mBase, uint32(v117)+12))
	v133 = int32(0)
	goto L50
L50:
	;
	v141 = *(*int32)(unsafe.Add(mBase, uint32(v123+v133<<(uint(int32(2))%32))))
	v142 = *(*int32)(unsafe.Add(mBase, uint32(v141)+108))
	if v142 != l3 {
		goto L52
	} else {
		goto L53
	}
L51:
	;
	goto L47
L52:
	;
	if l4 != v142 {
		goto L56
	} else {
		goto L57
	}
L53:
	;
	v144 = *(*int32)(unsafe.Add(mBase, uint32(v141)+112))
	if v144 != l4 {
		goto L52
	} else {
		goto L54
	}
L54:
	;
	v146 = *(*int32)(unsafe.Add(mBase, uint32(v141)+60))
	if v146 == v116 {
		v180 = v141
		goto L42
	} else {
		goto L55
	}
L55:
	;
	goto L52
L56:
	;
	v154 = v133 + int32(1)
	if v120 != v154 {
		v133 = v154
		goto L50
	} else {
		goto L60
	}
L57:
	;
	v149 = *(*int32)(unsafe.Add(mBase, uint32(v141)+112))
	if v149 != l3 {
		goto L56
	} else {
		goto L58
	}
L58:
	;
	v151 = *(*int32)(unsafe.Add(mBase, uint32(v141)+60))
	if v151 == v116 {
		v180 = v141
		goto L42
	} else {
		goto L59
	}
L59:
	;
	goto L56
L60:
	;
	goto L51
L61:
	;
	if v169 == int32(0) {
		goto L41
	} else {
		goto L62
	}
L62:
	;
	v180 = v169
	goto L42
L63:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l1)+28)) = v207
	v210 = *(*int32)(unsafe.Add(mBase, uint32(l1)+32))
	if v210 != 0 {
		goto L64
	} else {
		goto L65
	}
L64:
	;
	F_ec_add_clause_to_derives_hash(m, l1, v97)
	mBase = m.M
	v212 = m.ExcPending
	if v212 != 0 {
		goto L16
	} else {
		goto L67
	}
L65:
	;
	goto L66
L66:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_create_join_clause[0])) = v71
	v221 = v97
	goto L1
L67:
	;
	goto L66
}
func F_get_join_variables(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32, l5 int32) {
	mBase := m.M
	_ = mBase
	var v10 int32
	_ = v10
	var v13 int32
	_ = v13
	var v14 int32
	_ = v14
	var v15 int32
	_ = v15
	var v18 int32
	_ = v18
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
	var v36 int32
	_ = v36
	var v37 int32
	_ = v37
	var v39 int32
	_ = v39
	var v42 int32
	_ = v42
	var v43 int32
	_ = v43
	var v48 int32
	_ = v48
	var v55 int32
	_ = v55
	var v57 int32
	_ = v57
	var v59 int32
	_ = v59
	var v62 int32
	_ = v62
	var v64 int32
	_ = v64
	var v66 int32
	_ = v66
	var v73 int32
	_ = v73
	var v80 int32
	_ = v80
	var v83 int32
	_ = v83
	var v85 int32
	_ = v85
	var v88 int32
	_ = v88
	var v89 int32
	_ = v89
	var v90 int32
	_ = v90
	var v99 int32
	_ = v99
	var v100 int32
	_ = v100
	var v102 int32
	_ = v102
	var v105 int32
	_ = v105
	var v106 int32
	_ = v106
	var v111 int32
	_ = v111
	var v118 int32
	_ = v118
	var v120 int32
	_ = v120
	var v122 int32
	_ = v122
	var v125 int32
	_ = v125
	var v127 int32
	_ = v127
	var v129 int32
	_ = v129
	var v136 int32
	_ = v136
	var v143 int32
	_ = v143
	var v146 int32
	_ = v146
	var v148 int32
	_ = v148
	var v153 int32
	_ = v153
	var v157 int32
	_ = v157
	var v162 int32
	_ = v162
	if l1 == int32(0) {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v153 = m.ExcPending
	if v153 != 0 {
		goto L4
	} else {
		goto L41
	}
L2:
	;
	v10 = *(*int32)(unsafe.Add(mBase, uint32(l1)+4))
	if v10 != int32(2) {
		goto L1
	} else {
		goto L3
	}
L3:
	;
	v13 = *(*int32)(unsafe.Add(mBase, uint32(l1)+12))
	v14 = *(*int32)(unsafe.Add(mBase, uint32(v13)+4))
	v15 = *(*int32)(unsafe.Add(mBase, uint32(v13)))
	F_examine_variable(m, l0, v15, int32(0), l3)
	mBase = m.M
	v18 = m.ExcPending
	if v18 != 0 {
		goto L4
	} else {
		goto L5
	}
L4:
	;
	return
L5:
	;
	F_examine_variable(m, l0, v14, int32(0), l4)
	mBase = m.M
	v21 = m.ExcPending
	if v21 != 0 {
		goto L4
	} else {
		goto L6
	}
L6:
	;
	v22 = *(*int32)(unsafe.Add(mBase, uint32(l3)+4))
	if v22 == int32(0) {
		goto L7
	} else {
		goto L8
	}
L7:
	;
	v85 = *(*int32)(unsafe.Add(mBase, uint32(l4)+4))
	if v85 == int32(0) {
		goto L24
	} else {
		goto L25
	}
L8:
	;
	v25 = *(*int32)(unsafe.Add(mBase, uint32(v22)+8))
	v26 = *(*int32)(unsafe.Add(mBase, uint32(l2)+16))
	v27 = int32(0)
	if v25 == v27 {
		goto L10
	} else {
		goto L11
	}
L9:
	;
	if v80 == int32(0) {
		goto L7
	} else {
		goto L23
	}
L10:
	;
	v80 = int32(1)
	goto L9
L11:
	;
	goto L12
L12:
	;
	if v26 == int32(0) {
		v73 = v27
		goto L13
	} else {
		goto L14
	}
L13:
	;
	v80 = v73
	goto L9
L14:
	;
	v36 = *(*int32)(unsafe.Add(mBase, uint32(v25)+4))
	v37 = *(*int32)(unsafe.Add(mBase, uint32(v26)+4))
	if v37 < v36 {
		v73 = v27
		goto L13
	} else {
		goto L15
	}
L15:
	;
	v39 = int32(1)
	if v36 <= v39 {
		goto L16
	} else {
		goto L17
	}
L16:
	;
	v42 = v39
	goto L18
L17:
	;
	v42 = v36
	goto L18
L18:
	;
	v43 = int32(8)
	v48 = int32(0)
	goto L19
L19:
	;
	v55 = v48 << (uint(int32(2)) % 32)
	v57 = *(*int32)(unsafe.Add(mBase, uint32(v25+v43+v55)))
	v59 = *(*int32)(unsafe.Add(mBase, uint32(v26+v43+v55)))
	v62 = v57 & (v59 ^ int32(-1))
	v64 = base.B2i32(v62 == int32(0))
	if v62 != 0 {
		v73 = v64
		goto L13
	} else {
		goto L21
	}
L20:
	;
	v73 = v64
	goto L13
L21:
	;
	v66 = v48 + int32(1)
	if v66 != v42 {
		v48 = v66
		goto L19
	} else {
		goto L22
	}
L22:
	;
	goto L20
L23:
	;
	v83 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(l5))) = uint8(v83)
	return
L24:
	;
	v148 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(l5))) = uint8(v148)
	return
L25:
	;
	v88 = *(*int32)(unsafe.Add(mBase, uint32(v85)+8))
	v89 = *(*int32)(unsafe.Add(mBase, uint32(l2)+12))
	v90 = int32(0)
	if v88 == v90 {
		goto L27
	} else {
		goto L28
	}
L26:
	;
	if v143 == int32(0) {
		goto L24
	} else {
		goto L40
	}
L27:
	;
	v143 = int32(1)
	goto L26
L28:
	;
	goto L29
L29:
	;
	if v89 == int32(0) {
		v136 = v90
		goto L30
	} else {
		goto L31
	}
L30:
	;
	v143 = v136
	goto L26
L31:
	;
	v99 = *(*int32)(unsafe.Add(mBase, uint32(v88)+4))
	v100 = *(*int32)(unsafe.Add(mBase, uint32(v89)+4))
	if v100 < v99 {
		v136 = v90
		goto L30
	} else {
		goto L32
	}
L32:
	;
	v102 = int32(1)
	if v99 <= v102 {
		goto L33
	} else {
		goto L34
	}
L33:
	;
	v105 = v102
	goto L35
L34:
	;
	v105 = v99
	goto L35
L35:
	;
	v106 = int32(8)
	v111 = int32(0)
	goto L36
L36:
	;
	v118 = v111 << (uint(int32(2)) % 32)
	v120 = *(*int32)(unsafe.Add(mBase, uint32(v88+v106+v118)))
	v122 = *(*int32)(unsafe.Add(mBase, uint32(v89+v106+v118)))
	v125 = v120 & (v122 ^ int32(-1))
	v127 = base.B2i32(v125 == int32(0))
	if v125 != 0 {
		v136 = v127
		goto L30
	} else {
		goto L38
	}
L37:
	;
	v136 = v127
	goto L30
L38:
	;
	v129 = v111 + int32(1)
	if v129 != v105 {
		v111 = v129
		goto L36
	} else {
		goto L39
	}
L39:
	;
	goto L37
L40:
	;
	v146 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(l5))) = uint8(v146)
	return
L41:
	;
	F_errmsg_internal(m, int32(_a_F_get_join_variables_0), int32(0))
	mBase = m.M
	v157 = m.ExcPending
	if v157 != 0 {
		goto L4
	} else {
		goto L42
	}
L42:
	;
	F_errfinish(m, int32(_a_F_get_join_variables_1), int32(_a_F_get_join_variables_2), int32(_a_F_get_join_variables_3))
	mBase = m.M
	v162 = m.ExcPending
	if v162 != 0 {
		goto L4
	} else {
		goto L43
	}
L43:
	;
	base.Wasm_trap_unreachable()
	for {
	}
}
func F_join_is_legal(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32, l5 int32) int32 {
	mBase := m.M
	_ = mBase
	var v7 int32
	_ = v7
	var v19 int32
	_ = v19
	var v22 int32
	_ = v22
	var v31 int32
	_ = v31
	var v33 int32
	_ = v33
	var v34 int32
	_ = v34
	var v35 int32
	_ = v35
	var v37 int32
	_ = v37
	var v39 int32
	_ = v39
	var v43 int32
	_ = v43
	var v44 int32
	_ = v44
	var v45 int32
	_ = v45
	var v55 int32
	_ = v55
	var v56 int32
	_ = v56
	var v58 int32
	_ = v58
	var v61 int32
	_ = v61
	var v62 int32
	_ = v62
	var v67 int32
	_ = v67
	var v74 int32
	_ = v74
	var v76 int32
	_ = v76
	var v78 int32
	_ = v78
	var v79 int32
	_ = v79
	var v81 int32
	_ = v81
	var v83 int32
	_ = v83
	var v90 int32
	_ = v90
	var v93 int32
	_ = v93
	var v94 int32
	_ = v94
	var v103 int32
	_ = v103
	var v104 int32
	_ = v104
	var v106 int32
	_ = v106
	var v109 int32
	_ = v109
	var v110 int32
	_ = v110
	var v115 int32
	_ = v115
	var v122 int32
	_ = v122
	var v124 int32
	_ = v124
	var v126 int32
	_ = v126
	var v129 int32
	_ = v129
	var v131 int32
	_ = v131
	var v133 int32
	_ = v133
	var v140 int32
	_ = v140
	var v147 int32
	_ = v147
	var v148 int32
	_ = v148
	var v149 int32
	_ = v149
	var v150 int32
	_ = v150
	var v159 int32
	_ = v159
	var v160 int32
	_ = v160
	var v162 int32
	_ = v162
	var v165 int32
	_ = v165
	var v166 int32
	_ = v166
	var v171 int32
	_ = v171
	var v178 int32
	_ = v178
	var v180 int32
	_ = v180
	var v182 int32
	_ = v182
	var v185 int32
	_ = v185
	var v187 int32
	_ = v187
	var v189 int32
	_ = v189
	var v196 int32
	_ = v196
	var v203 int32
	_ = v203
	var v204 int32
	_ = v204
	var v205 int32
	_ = v205
	var v206 int32
	_ = v206
	var v215 int32
	_ = v215
	var v216 int32
	_ = v216
	var v218 int32
	_ = v218
	var v221 int32
	_ = v221
	var v222 int32
	_ = v222
	var v227 int32
	_ = v227
	var v234 int32
	_ = v234
	var v236 int32
	_ = v236
	var v238 int32
	_ = v238
	var v241 int32
	_ = v241
	var v243 int32
	_ = v243
	var v245 int32
	_ = v245
	var v252 int32
	_ = v252
	var v259 int32
	_ = v259
	var v260 int32
	_ = v260
	var v261 int32
	_ = v261
	var v262 int32
	_ = v262
	var v271 int32
	_ = v271
	var v272 int32
	_ = v272
	var v274 int32
	_ = v274
	var v277 int32
	_ = v277
	var v278 int32
	_ = v278
	var v283 int32
	_ = v283
	var v290 int32
	_ = v290
	var v292 int32
	_ = v292
	var v294 int32
	_ = v294
	var v297 int32
	_ = v297
	var v299 int32
	_ = v299
	var v301 int32
	_ = v301
	var v308 int32
	_ = v308
	var v315 int32
	_ = v315
	var v316 int32
	_ = v316
	var v317 int32
	_ = v317
	var v318 int32
	_ = v318
	var v327 int32
	_ = v327
	var v328 int32
	_ = v328
	var v330 int32
	_ = v330
	var v333 int32
	_ = v333
	var v334 int32
	_ = v334
	var v339 int32
	_ = v339
	var v346 int32
	_ = v346
	var v348 int32
	_ = v348
	var v350 int32
	_ = v350
	var v353 int32
	_ = v353
	var v355 int32
	_ = v355
	var v357 int32
	_ = v357
	var v364 int32
	_ = v364
	var v371 int32
	_ = v371
	var v372 int32
	_ = v372
	var v375 int32
	_ = v375
	var v376 int32
	_ = v376
	var v377 int32
	_ = v377
	var v386 int32
	_ = v386
	var v387 int32
	_ = v387
	var v389 int32
	_ = v389
	var v392 int32
	_ = v392
	var v393 int32
	_ = v393
	var v398 int32
	_ = v398
	var v405 int32
	_ = v405
	var v407 int32
	_ = v407
	var v409 int32
	_ = v409
	var v412 int32
	_ = v412
	var v414 int32
	_ = v414
	var v416 int32
	_ = v416
	var v423 int32
	_ = v423
	var v430 int32
	_ = v430
	var v431 int32
	_ = v431
	var v432 int32
	_ = v432
	var v433 int32
	_ = v433
	var v447 int32
	_ = v447
	var v448 int32
	_ = v448
	var v450 int32
	_ = v450
	var v453 int32
	_ = v453
	var v454 int32
	_ = v454
	var v459 int32
	_ = v459
	var v467 int32
	_ = v467
	var v469 int32
	_ = v469
	var v471 int32
	_ = v471
	var v472 int32
	_ = v472
	var v475 int32
	_ = v475
	var v479 int32
	_ = v479
	var v486 int32
	_ = v486
	var v487 int32
	_ = v487
	var v488 int32
	_ = v488
	var v497 int32
	_ = v497
	var v498 int32
	_ = v498
	var v500 int32
	_ = v500
	var v503 int32
	_ = v503
	var v504 int32
	_ = v504
	var v509 int32
	_ = v509
	var v516 int32
	_ = v516
	var v518 int32
	_ = v518
	var v520 int32
	_ = v520
	var v523 int32
	_ = v523
	var v525 int32
	_ = v525
	var v527 int32
	_ = v527
	var v534 int32
	_ = v534
	var v541 int32
	_ = v541
	var v544 int32
	_ = v544
	var v545 int32
	_ = v545
	var v546 int32
	_ = v546
	var v560 int32
	_ = v560
	var v561 int32
	_ = v561
	var v563 int32
	_ = v563
	var v566 int32
	_ = v566
	var v567 int32
	_ = v567
	var v572 int32
	_ = v572
	var v580 int32
	_ = v580
	var v582 int32
	_ = v582
	var v584 int32
	_ = v584
	var v585 int32
	_ = v585
	var v588 int32
	_ = v588
	var v592 int32
	_ = v592
	var v599 int32
	_ = v599
	var v600 int32
	_ = v600
	var v601 int32
	_ = v601
	var v610 int32
	_ = v610
	var v611 int32
	_ = v611
	var v613 int32
	_ = v613
	var v616 int32
	_ = v616
	var v617 int32
	_ = v617
	var v622 int32
	_ = v622
	var v629 int32
	_ = v629
	var v631 int32
	_ = v631
	var v633 int32
	_ = v633
	var v636 int32
	_ = v636
	var v638 int32
	_ = v638
	var v640 int32
	_ = v640
	var v647 int32
	_ = v647
	var v654 int32
	_ = v654
	var v657 int32
	_ = v657
	var v658 int32
	_ = v658
	var v659 int32
	_ = v659
	var v668 int32
	_ = v668
	var v669 int32
	_ = v669
	var v671 int32
	_ = v671
	var v674 int32
	_ = v674
	var v675 int32
	_ = v675
	var v680 int32
	_ = v680
	var v687 int32
	_ = v687
	var v689 int32
	_ = v689
	var v691 int32
	_ = v691
	var v694 int32
	_ = v694
	var v696 int32
	_ = v696
	var v698 int32
	_ = v698
	var v705 int32
	_ = v705
	var v712 int32
	_ = v712
	var v715 int32
	_ = v715
	var v718 int32
	_ = v718
	var v719 int32
	_ = v719
	var v720 int32
	_ = v720
	var v729 int32
	_ = v729
	var v730 int32
	_ = v730
	var v732 int32
	_ = v732
	var v735 int32
	_ = v735
	var v736 int32
	_ = v736
	var v741 int32
	_ = v741
	var v748 int32
	_ = v748
	var v750 int32
	_ = v750
	var v752 int32
	_ = v752
	var v755 int32
	_ = v755
	var v757 int32
	_ = v757
	var v759 int32
	_ = v759
	var v766 int32
	_ = v766
	var v773 int32
	_ = v773
	var v776 int32
	_ = v776
	var v777 int32
	_ = v777
	var v778 int32
	_ = v778
	var v787 int32
	_ = v787
	var v788 int32
	_ = v788
	var v790 int32
	_ = v790
	var v793 int32
	_ = v793
	var v794 int32
	_ = v794
	var v799 int32
	_ = v799
	var v806 int32
	_ = v806
	var v808 int32
	_ = v808
	var v810 int32
	_ = v810
	var v813 int32
	_ = v813
	var v815 int32
	_ = v815
	var v817 int32
	_ = v817
	var v824 int32
	_ = v824
	var v831 int32
	_ = v831
	var v839 int32
	_ = v839
	var v842 int32
	_ = v842
	var v843 int32
	_ = v843
	var v844 int32
	_ = v844
	var v858 int32
	_ = v858
	var v859 int32
	_ = v859
	var v861 int32
	_ = v861
	var v864 int32
	_ = v864
	var v865 int32
	_ = v865
	var v870 int32
	_ = v870
	var v878 int32
	_ = v878
	var v880 int32
	_ = v880
	var v882 int32
	_ = v882
	var v883 int32
	_ = v883
	var v886 int32
	_ = v886
	var v890 int32
	_ = v890
	var v897 int32
	_ = v897
	var v900 int32
	_ = v900
	var v907 int32
	_ = v907
	var v910 int32
	_ = v910
	var v911 int32
	_ = v911
	var v912 int32
	_ = v912
	var v926 int32
	_ = v926
	var v927 int32
	_ = v927
	var v929 int32
	_ = v929
	var v932 int32
	_ = v932
	var v933 int32
	_ = v933
	var v938 int32
	_ = v938
	var v946 int32
	_ = v946
	var v948 int32
	_ = v948
	var v950 int32
	_ = v950
	var v951 int32
	_ = v951
	var v954 int32
	_ = v954
	var v958 int32
	_ = v958
	var v965 int32
	_ = v965
	var v966 int32
	_ = v966
	var v971 int32
	_ = v971
	var v973 int32
	_ = v973
	var v974 int32
	_ = v974
	var v975 int32
	_ = v975
	var v985 int32
	_ = v985
	var v986 int32
	_ = v986
	var v988 int32
	_ = v988
	var v991 int32
	_ = v991
	var v992 int32
	_ = v992
	var v997 int32
	_ = v997
	var v1004 int32
	_ = v1004
	var v1006 int32
	_ = v1006
	var v1008 int32
	_ = v1008
	var v1009 int32
	_ = v1009
	var v1011 int32
	_ = v1011
	var v1013 int32
	_ = v1013
	var v1020 int32
	_ = v1020
	var v1021 int32
	_ = v1021
	var v1022 int32
	_ = v1022
	var v1023 int32
	_ = v1023
	var v1033 int32
	_ = v1033
	var v1034 int32
	_ = v1034
	var v1036 int32
	_ = v1036
	var v1039 int32
	_ = v1039
	var v1040 int32
	_ = v1040
	var v1045 int32
	_ = v1045
	var v1052 int32
	_ = v1052
	var v1054 int32
	_ = v1054
	var v1056 int32
	_ = v1056
	var v1057 int32
	_ = v1057
	var v1059 int32
	_ = v1059
	var v1061 int32
	_ = v1061
	var v1068 int32
	_ = v1068
	var v1069 int32
	_ = v1069
	var v1072 int32
	_ = v1072
	var v1073 int32
	_ = v1073
	var v1083 int32
	_ = v1083
	var v1084 int32
	_ = v1084
	var v1086 int32
	_ = v1086
	var v1089 int32
	_ = v1089
	var v1090 int32
	_ = v1090
	var v1095 int32
	_ = v1095
	var v1102 int32
	_ = v1102
	var v1104 int32
	_ = v1104
	var v1106 int32
	_ = v1106
	var v1107 int32
	_ = v1107
	var v1109 int32
	_ = v1109
	var v1111 int32
	_ = v1111
	var v1118 int32
	_ = v1118
	var v1120 int32
	_ = v1120
	var v1121 int32
	_ = v1121
	var v1122 int32
	_ = v1122
	var v1123 int32
	_ = v1123
	var v1124 int32
	_ = v1124
	var v1126 int32
	_ = v1126
	var v1127 int32
	_ = v1127
	var v1135 int32
	_ = v1135
	var v1138 int32
	_ = v1138
	var v1139 int32
	_ = v1139
	var v1141 int32
	_ = v1141
	var v1147 int32
	_ = v1147
	var v1150 int32
	_ = v1150
	var v1159 int32
	_ = v1159
	var v1163 int32
	_ = v1163
	var v1165 int32
	_ = v1165
	var v1167 int32
	_ = v1167
	var v1170 int32
	_ = v1170
	var v1171 int32
	_ = v1171
	var v1172 int32
	_ = v1172
	var v1182 int32
	_ = v1182
	var v1183 int32
	_ = v1183
	var v1185 int32
	_ = v1185
	var v1188 int32
	_ = v1188
	var v1189 int32
	_ = v1189
	var v1194 int32
	_ = v1194
	var v1201 int32
	_ = v1201
	var v1203 int32
	_ = v1203
	var v1205 int32
	_ = v1205
	var v1206 int32
	_ = v1206
	var v1208 int32
	_ = v1208
	var v1210 int32
	_ = v1210
	var v1217 int32
	_ = v1217
	var v1218 int32
	_ = v1218
	var v1219 int32
	_ = v1219
	var v1220 int32
	_ = v1220
	var v1231 int32
	_ = v1231
	var v1232 int32
	_ = v1232
	var v1234 int32
	_ = v1234
	var v1237 int32
	_ = v1237
	var v1238 int32
	_ = v1238
	var v1243 int32
	_ = v1243
	var v1250 int32
	_ = v1250
	var v1252 int32
	_ = v1252
	var v1254 int32
	_ = v1254
	var v1255 int32
	_ = v1255
	var v1257 int32
	_ = v1257
	var v1259 int32
	_ = v1259
	var v1266 int32
	_ = v1266
	var v1267 int32
	_ = v1267
	var v1269 int32
	_ = v1269
	var v1272 int32
	_ = v1272
	var v1273 int32
	_ = v1273
	var v1274 int32
	_ = v1274
	var v1284 int32
	_ = v1284
	var v1285 int32
	_ = v1285
	var v1287 int32
	_ = v1287
	var v1290 int32
	_ = v1290
	var v1291 int32
	_ = v1291
	var v1296 int32
	_ = v1296
	var v1303 int32
	_ = v1303
	var v1305 int32
	_ = v1305
	var v1307 int32
	_ = v1307
	var v1308 int32
	_ = v1308
	var v1310 int32
	_ = v1310
	var v1312 int32
	_ = v1312
	var v1319 int32
	_ = v1319
	var v1327 int32
	_ = v1327
	var v1330 int32
	_ = v1330
	var v1331 int32
	_ = v1331
	var v1332 int32
	_ = v1332
	var v1342 int32
	_ = v1342
	var v1343 int32
	_ = v1343
	var v1345 int32
	_ = v1345
	var v1348 int32
	_ = v1348
	var v1349 int32
	_ = v1349
	var v1354 int32
	_ = v1354
	var v1361 int32
	_ = v1361
	var v1363 int32
	_ = v1363
	var v1365 int32
	_ = v1365
	var v1366 int32
	_ = v1366
	var v1368 int32
	_ = v1368
	var v1370 int32
	_ = v1370
	var v1377 int32
	_ = v1377
	var v1380 int32
	_ = v1380
	var v1381 int32
	_ = v1381
	var v1382 int32
	_ = v1382
	var v1383 int32
	_ = v1383
	var v1384 int32
	_ = v1384
	var v1385 int32
	_ = v1385
	var v1388 int32
	_ = v1388
	var v1389 int32
	_ = v1389
	var v1393 int32
	_ = v1393
	var v1404 int32
	_ = v1404
	var v1405 int32
	_ = v1405
	var v1407 int32
	_ = v1407
	var v1411 int32
	_ = v1411
	var v1413 int32
	_ = v1413
	var v1418 int32
	_ = v1418
	var v1424 int32
	_ = v1424
	var v1425 int32
	_ = v1425
	var v1428 int32
	_ = v1428
	var v1429 int32
	_ = v1429
	var v1432 int32
	_ = v1432
	var v1433 int32
	_ = v1433
	var v1443 int32
	_ = v1443
	var v1444 int32
	_ = v1444
	var v1446 int32
	_ = v1446
	var v1449 int32
	_ = v1449
	var v1450 int32
	_ = v1450
	var v1455 int32
	_ = v1455
	var v1462 int32
	_ = v1462
	var v1464 int32
	_ = v1464
	var v1466 int32
	_ = v1466
	var v1467 int32
	_ = v1467
	var v1469 int32
	_ = v1469
	var v1471 int32
	_ = v1471
	var v1478 int32
	_ = v1478
	var v1481 int32
	_ = v1481
	var v1482 int32
	_ = v1482
	var v1491 int32
	_ = v1491
	var v1492 int32
	_ = v1492
	var v1494 int32
	_ = v1494
	var v1497 int32
	_ = v1497
	var v1498 int32
	_ = v1498
	var v1503 int32
	_ = v1503
	var v1510 int32
	_ = v1510
	var v1512 int32
	_ = v1512
	var v1514 int32
	_ = v1514
	var v1517 int32
	_ = v1517
	var v1519 int32
	_ = v1519
	var v1521 int32
	_ = v1521
	var v1528 int32
	_ = v1528
	var v1535 int32
	_ = v1535
	var v1537 int32
	_ = v1537
	var v1538 int32
	_ = v1538
	var v1539 int32
	_ = v1539
	var v1540 int32
	_ = v1540
	var v1541 int32
	_ = v1541
	var v1543 int32
	_ = v1543
	var v1544 int32
	_ = v1544
	var v1547 int32
	_ = v1547
	var v1549 int32
	_ = v1549
	var v1565 int32
	_ = v1565
	var v1576 int32
	_ = v1576
	var v1586 int32
	_ = v1586
	var v1587 int32
	_ = v1587
	var v1589 int32
	_ = v1589
	var v1592 int32
	_ = v1592
	var v1593 int32
	_ = v1593
	var v1598 int32
	_ = v1598
	var v1605 int32
	_ = v1605
	var v1607 int32
	_ = v1607
	var v1609 int32
	_ = v1609
	var v1610 int32
	_ = v1610
	var v1612 int32
	_ = v1612
	var v1614 int32
	_ = v1614
	var v1621 int32
	_ = v1621
	var v1622 int32
	_ = v1622
	var v1623 int32
	_ = v1623
	var v1633 int32
	_ = v1633
	var v1634 int32
	_ = v1634
	var v1636 int32
	_ = v1636
	var v1639 int32
	_ = v1639
	var v1640 int32
	_ = v1640
	var v1645 int32
	_ = v1645
	var v1652 int32
	_ = v1652
	var v1654 int32
	_ = v1654
	var v1656 int32
	_ = v1656
	var v1657 int32
	_ = v1657
	var v1659 int32
	_ = v1659
	var v1661 int32
	_ = v1661
	var v1668 int32
	_ = v1668
	var v1671 int32
	_ = v1671
	var v1674 int32
	_ = v1674
	var v1686 int32
	_ = v1686
	var v1692 int32
	_ = v1692
	var v1696 int32
	_ = v1696
	var v1697 int32
	_ = v1697
	var v1698 int32
	_ = v1698
	var v1699 int32
	_ = v1699
	var v1700 int32
	_ = v1700
	var v1701 int32
	_ = v1701
	var v1711 int32
	_ = v1711
	var v1712 int32
	_ = v1712
	var v1714 int32
	_ = v1714
	var v1717 int32
	_ = v1717
	var v1718 int32
	_ = v1718
	var v1723 int32
	_ = v1723
	var v1730 int32
	_ = v1730
	var v1732 int32
	_ = v1732
	var v1734 int32
	_ = v1734
	var v1735 int32
	_ = v1735
	var v1737 int32
	_ = v1737
	var v1739 int32
	_ = v1739
	var v1746 int32
	_ = v1746
	var v1747 int32
	_ = v1747
	var v1748 int32
	_ = v1748
	var v1758 int32
	_ = v1758
	var v1759 int32
	_ = v1759
	var v1761 int32
	_ = v1761
	var v1764 int32
	_ = v1764
	var v1765 int32
	_ = v1765
	var v1770 int32
	_ = v1770
	var v1777 int32
	_ = v1777
	var v1779 int32
	_ = v1779
	var v1781 int32
	_ = v1781
	var v1782 int32
	_ = v1782
	var v1784 int32
	_ = v1784
	var v1786 int32
	_ = v1786
	var v1793 int32
	_ = v1793
	var v1795 int32
	_ = v1795
	var v1796 int32
	_ = v1796
	var v1828 int32
	_ = v1828
	v7 = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(l4))) = v7
	*(*uint8)(unsafe.Add(mBase, uint32(l5))) = uint8(v7)
	v19 = *(*int32)(unsafe.Add(mBase, uint32(l0)+120))
	if v19 == v7 {
		v1159 = v7
		v1163 = v7
		v1165 = v7
		goto L2
	} else {
		goto L3
	}
L1:
	;
	return v1828
L2:
	;
	v1167 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+333)))
	if v1167 != int32(1) {
		goto L317
	} else {
		goto L318
	}
L3:
	;
	v22 = *(*int32)(unsafe.Add(mBase, uint32(v19)+4))
	if int32(0) < v22 {
		goto L4
	} else {
		goto L5
	}
L4:
	;
	v31 = v7
	v33 = v7
	v34 = v7
	v35 = v7
	v37 = v7
	goto L7
L5:
	;
	v1135 = v7
	v1138 = v7
	v1139 = v7
	v1141 = v7
	goto L6
L6:
	;
	if v1138 == int32(0) {
		v1159 = v1135
		v1163 = v1139
		v1165 = v1141
		goto L2
	} else {
		goto L313
	}
L7:
	;
	v39 = *(*int32)(unsafe.Add(mBase, uint32(v19)+12))
	v43 = *(*int32)(unsafe.Add(mBase, uint32(v39+v33<<(uint(int32(2))%32))))
	v44 = *(*int32)(unsafe.Add(mBase, uint32(v43)+8))
	v45 = int32(0)
	if base.B2i32(v44 == v45)|base.B2i32(l3 == v45) != 0 {
		v90 = v45
		goto L12
	} else {
		goto L13
	}
L8:
	;
	v1135 = v1121
	v1138 = v1122
	v1139 = v1123
	v1141 = v1124
	goto L6
L9:
	;
	v1126 = v33 + int32(1)
	v1127 = *(*int32)(unsafe.Add(mBase, uint32(v19)+4))
	if v1126 < v1127 {
		v31 = v1121
		v33 = v1126
		v34 = v1122
		v35 = v1123
		v37 = v1124
		goto L7
	} else {
		goto L312
	}
L10:
	;
	v1121 = v31
	v1122 = v1120
	v1123 = v35
	v1124 = v37
	goto L9
L11:
	;
	if v90 == int32(0) {
		v1120 = v34
		goto L10
	} else {
		goto L24
	}
L12:
	;
	goto L11
L13:
	;
	v55 = *(*int32)(unsafe.Add(mBase, uint32(v44)+4))
	v56 = *(*int32)(unsafe.Add(mBase, uint32(l3)+4))
	if v55 < v56 {
		goto L14
	} else {
		goto L15
	}
L14:
	;
	v58 = v55
	goto L16
L15:
	;
	v58 = v56
	goto L16
L16:
	;
	if v58 <= int32(1) {
		goto L17
	} else {
		goto L18
	}
L17:
	;
	v61 = int32(1)
	goto L19
L18:
	;
	v61 = v58
	goto L19
L19:
	;
	v62 = int32(8)
	v67 = int32(0)
	goto L20
L20:
	;
	v74 = v67 << (uint(int32(2)) % 32)
	v76 = *(*int32)(unsafe.Add(mBase, uint32(l3+v62+v74)))
	v78 = *(*int32)(unsafe.Add(mBase, uint32(v44+v62+v74)))
	v79 = v76 & v78
	v81 = base.B2i32(v79 != int32(0))
	if v79 != 0 {
		v90 = v81
		goto L12
	} else {
		goto L22
	}
L21:
	;
	v90 = v81
	goto L12
L22:
	;
	v83 = v67 + int32(1)
	if v83 != v61 {
		v67 = v83
		goto L20
	} else {
		goto L23
	}
L23:
	;
	goto L21
L24:
	;
	v93 = *(*int32)(unsafe.Add(mBase, uint32(v43)+8))
	v94 = int32(0)
	if l3 == v94 {
		goto L26
	} else {
		goto L27
	}
L25:
	;
	if v147 != 0 {
		v1120 = v34
		goto L10
	} else {
		goto L39
	}
L26:
	;
	v147 = int32(1)
	goto L25
L27:
	;
	goto L28
L28:
	;
	if v93 == int32(0) {
		v140 = v94
		goto L29
	} else {
		goto L30
	}
L29:
	;
	v147 = v140
	goto L25
L30:
	;
	v103 = *(*int32)(unsafe.Add(mBase, uint32(l3)+4))
	v104 = *(*int32)(unsafe.Add(mBase, uint32(v93)+4))
	if v104 < v103 {
		v140 = v94
		goto L29
	} else {
		goto L31
	}
L31:
	;
	v106 = int32(1)
	if v103 <= v106 {
		goto L32
	} else {
		goto L33
	}
L32:
	;
	v109 = v106
	goto L34
L33:
	;
	v109 = v103
	goto L34
L34:
	;
	v110 = int32(8)
	v115 = int32(0)
	goto L35
L35:
	;
	v122 = v115 << (uint(int32(2)) % 32)
	v124 = *(*int32)(unsafe.Add(mBase, uint32(l3+v110+v122)))
	v126 = *(*int32)(unsafe.Add(mBase, uint32(v93+v110+v122)))
	v129 = v124 & (v126 ^ int32(-1))
	v131 = base.B2i32(v129 == int32(0))
	if v129 != 0 {
		v140 = v131
		goto L29
	} else {
		goto L37
	}
L36:
	;
	v140 = v131
	goto L29
L37:
	;
	v133 = v115 + int32(1)
	if v133 != v109 {
		v115 = v133
		goto L35
	} else {
		goto L38
	}
L38:
	;
	goto L36
L39:
	;
	v148 = *(*int32)(unsafe.Add(mBase, uint32(v43)+4))
	v149 = *(*int32)(unsafe.Add(mBase, uint32(l1)+8))
	v150 = int32(0)
	if v148 == v150 {
		goto L41
	} else {
		goto L42
	}
L40:
	;
	if v203 != 0 {
		goto L54
	} else {
		goto L55
	}
L41:
	;
	v203 = int32(1)
	goto L40
L42:
	;
	goto L43
L43:
	;
	if v149 == int32(0) {
		v196 = v150
		goto L44
	} else {
		goto L45
	}
L44:
	;
	v203 = v196
	goto L40
L45:
	;
	v159 = *(*int32)(unsafe.Add(mBase, uint32(v148)+4))
	v160 = *(*int32)(unsafe.Add(mBase, uint32(v149)+4))
	if v160 < v159 {
		v196 = v150
		goto L44
	} else {
		goto L46
	}
L46:
	;
	v162 = int32(1)
	if v159 <= v162 {
		goto L47
	} else {
		goto L48
	}
L47:
	;
	v165 = v162
	goto L49
L48:
	;
	v165 = v159
	goto L49
L49:
	;
	v166 = int32(8)
	v171 = int32(0)
	goto L50
L50:
	;
	v178 = v171 << (uint(int32(2)) % 32)
	v180 = *(*int32)(unsafe.Add(mBase, uint32(v148+v166+v178)))
	v182 = *(*int32)(unsafe.Add(mBase, uint32(v149+v166+v178)))
	v185 = v180 & (v182 ^ int32(-1))
	v187 = base.B2i32(v185 == int32(0))
	if v185 != 0 {
		v196 = v187
		goto L44
	} else {
		goto L52
	}
L51:
	;
	v196 = v187
	goto L44
L52:
	;
	v189 = v171 + int32(1)
	if v189 != v165 {
		v171 = v189
		goto L50
	} else {
		goto L53
	}
L53:
	;
	goto L51
L54:
	;
	v204 = *(*int32)(unsafe.Add(mBase, uint32(v43)+8))
	v205 = *(*int32)(unsafe.Add(mBase, uint32(l1)+8))
	v206 = int32(0)
	if v204 == v206 {
		goto L58
	} else {
		goto L59
	}
L55:
	;
	goto L56
L56:
	;
	v260 = *(*int32)(unsafe.Add(mBase, uint32(v43)+4))
	v261 = *(*int32)(unsafe.Add(mBase, uint32(l2)+8))
	v262 = int32(0)
	if v260 == v262 {
		goto L73
	} else {
		goto L74
	}
L57:
	;
	if v259 != 0 {
		v1120 = v34
		goto L10
	} else {
		goto L71
	}
L58:
	;
	v259 = int32(1)
	goto L57
L59:
	;
	goto L60
L60:
	;
	if v205 == int32(0) {
		v252 = v206
		goto L61
	} else {
		goto L62
	}
L61:
	;
	v259 = v252
	goto L57
L62:
	;
	v215 = *(*int32)(unsafe.Add(mBase, uint32(v204)+4))
	v216 = *(*int32)(unsafe.Add(mBase, uint32(v205)+4))
	if v216 < v215 {
		v252 = v206
		goto L61
	} else {
		goto L63
	}
L63:
	;
	v218 = int32(1)
	if v215 <= v218 {
		goto L64
	} else {
		goto L65
	}
L64:
	;
	v221 = v218
	goto L66
L65:
	;
	v221 = v215
	goto L66
L66:
	;
	v222 = int32(8)
	v227 = int32(0)
	goto L67
L67:
	;
	v234 = v227 << (uint(int32(2)) % 32)
	v236 = *(*int32)(unsafe.Add(mBase, uint32(v204+v222+v234)))
	v238 = *(*int32)(unsafe.Add(mBase, uint32(v205+v222+v234)))
	v241 = v236 & (v238 ^ int32(-1))
	v243 = base.B2i32(v241 == int32(0))
	if v241 != 0 {
		v252 = v243
		goto L61
	} else {
		goto L69
	}
L68:
	;
	v252 = v243
	goto L61
L69:
	;
	v245 = v227 + int32(1)
	if v245 != v221 {
		v227 = v245
		goto L67
	} else {
		goto L70
	}
L70:
	;
	goto L68
L71:
	;
	goto L56
L72:
	;
	if v315 != 0 {
		goto L86
	} else {
		goto L87
	}
L73:
	;
	v315 = int32(1)
	goto L72
L74:
	;
	goto L75
L75:
	;
	if v261 == int32(0) {
		v308 = v262
		goto L76
	} else {
		goto L77
	}
L76:
	;
	v315 = v308
	goto L72
L77:
	;
	v271 = *(*int32)(unsafe.Add(mBase, uint32(v260)+4))
	v272 = *(*int32)(unsafe.Add(mBase, uint32(v261)+4))
	if v272 < v271 {
		v308 = v262
		goto L76
	} else {
		goto L78
	}
L78:
	;
	v274 = int32(1)
	if v271 <= v274 {
		goto L79
	} else {
		goto L80
	}
L79:
	;
	v277 = v274
	goto L81
L80:
	;
	v277 = v271
	goto L81
L81:
	;
	v278 = int32(8)
	v283 = int32(0)
	goto L82
L82:
	;
	v290 = v283 << (uint(int32(2)) % 32)
	v292 = *(*int32)(unsafe.Add(mBase, uint32(v260+v278+v290)))
	v294 = *(*int32)(unsafe.Add(mBase, uint32(v261+v278+v290)))
	v297 = v292 & (v294 ^ int32(-1))
	v299 = base.B2i32(v297 == int32(0))
	if v297 != 0 {
		v308 = v299
		goto L76
	} else {
		goto L84
	}
L83:
	;
	v308 = v299
	goto L76
L84:
	;
	v301 = v283 + int32(1)
	if v301 != v277 {
		v283 = v301
		goto L82
	} else {
		goto L85
	}
L85:
	;
	goto L83
L86:
	;
	v316 = *(*int32)(unsafe.Add(mBase, uint32(v43)+8))
	v317 = *(*int32)(unsafe.Add(mBase, uint32(l2)+8))
	v318 = int32(0)
	if v316 == v318 {
		goto L90
	} else {
		goto L91
	}
L87:
	;
	goto L88
L88:
	;
	v372 = *(*int32)(unsafe.Add(mBase, uint32(v43)+20))
	if v372 != int32(4) {
		goto L104
	} else {
		goto L105
	}
L89:
	;
	if v371 != 0 {
		v1120 = v34
		goto L10
	} else {
		goto L103
	}
L90:
	;
	v371 = int32(1)
	goto L89
L91:
	;
	goto L92
L92:
	;
	if v317 == int32(0) {
		v364 = v318
		goto L93
	} else {
		goto L94
	}
L93:
	;
	v371 = v364
	goto L89
L94:
	;
	v327 = *(*int32)(unsafe.Add(mBase, uint32(v316)+4))
	v328 = *(*int32)(unsafe.Add(mBase, uint32(v317)+4))
	if v328 < v327 {
		v364 = v318
		goto L93
	} else {
		goto L95
	}
L95:
	;
	v330 = int32(1)
	if v327 <= v330 {
		goto L96
	} else {
		goto L97
	}
L96:
	;
	v333 = v330
	goto L98
L97:
	;
	v333 = v327
	goto L98
L98:
	;
	v334 = int32(8)
	v339 = int32(0)
	goto L99
L99:
	;
	v346 = v339 << (uint(int32(2)) % 32)
	v348 = *(*int32)(unsafe.Add(mBase, uint32(v316+v334+v346)))
	v350 = *(*int32)(unsafe.Add(mBase, uint32(v317+v334+v346)))
	v353 = v348 & (v350 ^ int32(-1))
	v355 = base.B2i32(v353 == int32(0))
	if v353 != 0 {
		v364 = v355
		goto L93
	} else {
		goto L101
	}
L100:
	;
	v364 = v355
	goto L93
L101:
	;
	v357 = v339 + int32(1)
	if v357 != v333 {
		v339 = v357
		goto L99
	} else {
		goto L102
	}
L102:
	;
	goto L100
L103:
	;
	goto L88
L104:
	;
	v599 = *(*int32)(unsafe.Add(mBase, uint32(v43)+4))
	v600 = *(*int32)(unsafe.Add(mBase, uint32(l1)+8))
	v601 = int32(0)
	if v599 == v601 {
		goto L164
	} else {
		goto L165
	}
L105:
	;
	v375 = *(*int32)(unsafe.Add(mBase, uint32(v43)+16))
	v376 = *(*int32)(unsafe.Add(mBase, uint32(l1)+8))
	v377 = int32(0)
	if v375 == v377 {
		goto L107
	} else {
		goto L108
	}
L106:
	;
	if v430 != 0 {
		goto L120
	} else {
		goto L121
	}
L107:
	;
	v430 = int32(1)
	goto L106
L108:
	;
	goto L109
L109:
	;
	if v376 == int32(0) {
		v423 = v377
		goto L110
	} else {
		goto L111
	}
L110:
	;
	v430 = v423
	goto L106
L111:
	;
	v386 = *(*int32)(unsafe.Add(mBase, uint32(v375)+4))
	v387 = *(*int32)(unsafe.Add(mBase, uint32(v376)+4))
	if v387 < v386 {
		v423 = v377
		goto L110
	} else {
		goto L112
	}
L112:
	;
	v389 = int32(1)
	if v386 <= v389 {
		goto L113
	} else {
		goto L114
	}
L113:
	;
	v392 = v389
	goto L115
L114:
	;
	v392 = v386
	goto L115
L115:
	;
	v393 = int32(8)
	v398 = int32(0)
	goto L116
L116:
	;
	v405 = v398 << (uint(int32(2)) % 32)
	v407 = *(*int32)(unsafe.Add(mBase, uint32(v375+v393+v405)))
	v409 = *(*int32)(unsafe.Add(mBase, uint32(v376+v393+v405)))
	v412 = v407 & (v409 ^ int32(-1))
	v414 = base.B2i32(v412 == int32(0))
	if v412 != 0 {
		v423 = v414
		goto L110
	} else {
		goto L118
	}
L117:
	;
	v423 = v414
	goto L110
L118:
	;
	v416 = v398 + int32(1)
	if v416 != v392 {
		v398 = v416
		goto L116
	} else {
		goto L119
	}
L119:
	;
	goto L117
L120:
	;
	v431 = *(*int32)(unsafe.Add(mBase, uint32(v43)+16))
	v432 = *(*int32)(unsafe.Add(mBase, uint32(l1)+8))
	v433 = int32(0)
	if base.B2i32(v431 == v433)|base.B2i32(v432 == v433) != 0 {
		v479 = base.B2i32(v431|v432 == v433)
		goto L124
	} else {
		goto L125
	}
L121:
	;
	goto L122
L122:
	;
	v486 = *(*int32)(unsafe.Add(mBase, uint32(v43)+16))
	v487 = *(*int32)(unsafe.Add(mBase, uint32(l2)+8))
	v488 = int32(0)
	if v486 == v488 {
		goto L136
	} else {
		goto L137
	}
L123:
	;
	if v479 == int32(0) {
		v1120 = v34
		goto L10
	} else {
		goto L134
	}
L124:
	;
	goto L123
L125:
	;
	v447 = *(*int32)(unsafe.Add(mBase, uint32(v431)+4))
	v448 = *(*int32)(unsafe.Add(mBase, uint32(v432)+4))
	if v447 != v448 {
		v479 = int32(0)
		goto L124
	} else {
		goto L126
	}
L126:
	;
	v450 = int32(1)
	if v447 <= v450 {
		goto L127
	} else {
		goto L128
	}
L127:
	;
	v453 = v450
	goto L129
L128:
	;
	v453 = v447
	goto L129
L129:
	;
	v454 = int32(8)
	v459 = int32(0)
	goto L130
L130:
	;
	v467 = v459 << (uint(int32(2)) % 32)
	v469 = *(*int32)(unsafe.Add(mBase, uint32(v431+v454+v467)))
	v471 = *(*int32)(unsafe.Add(mBase, uint32(v432+v454+v467)))
	v472 = base.B2i32(v469 == v471)
	if v469 != v471 {
		v479 = v472
		goto L124
	} else {
		goto L132
	}
L131:
	;
	v479 = v472
	goto L124
L132:
	;
	v475 = v459 + int32(1)
	if v475 != v453 {
		v459 = v475
		goto L130
	} else {
		goto L133
	}
L133:
	;
	goto L131
L134:
	;
	goto L122
L135:
	;
	if v541 == int32(0) {
		goto L104
	} else {
		goto L149
	}
L136:
	;
	v541 = int32(1)
	goto L135
L137:
	;
	goto L138
L138:
	;
	if v487 == int32(0) {
		v534 = v488
		goto L139
	} else {
		goto L140
	}
L139:
	;
	v541 = v534
	goto L135
L140:
	;
	v497 = *(*int32)(unsafe.Add(mBase, uint32(v486)+4))
	v498 = *(*int32)(unsafe.Add(mBase, uint32(v487)+4))
	if v498 < v497 {
		v534 = v488
		goto L139
	} else {
		goto L141
	}
L141:
	;
	v500 = int32(1)
	if v497 <= v500 {
		goto L142
	} else {
		goto L143
	}
L142:
	;
	v503 = v500
	goto L144
L143:
	;
	v503 = v497
	goto L144
L144:
	;
	v504 = int32(8)
	v509 = int32(0)
	goto L145
L145:
	;
	v516 = v509 << (uint(int32(2)) % 32)
	v518 = *(*int32)(unsafe.Add(mBase, uint32(v486+v504+v516)))
	v520 = *(*int32)(unsafe.Add(mBase, uint32(v487+v504+v516)))
	v523 = v518 & (v520 ^ int32(-1))
	v525 = base.B2i32(v523 == int32(0))
	if v523 != 0 {
		v534 = v525
		goto L139
	} else {
		goto L147
	}
L146:
	;
	v534 = v525
	goto L139
L147:
	;
	v527 = v509 + int32(1)
	if v527 != v503 {
		v509 = v527
		goto L145
	} else {
		goto L148
	}
L148:
	;
	goto L146
L149:
	;
	v544 = *(*int32)(unsafe.Add(mBase, uint32(v43)+16))
	v545 = *(*int32)(unsafe.Add(mBase, uint32(l2)+8))
	v546 = int32(0)
	if base.B2i32(v544 == v546)|base.B2i32(v545 == v546) != 0 {
		v592 = base.B2i32(v544|v545 == v546)
		goto L151
	} else {
		goto L152
	}
L150:
	;
	if v592 == int32(0) {
		v1120 = v34
		goto L10
	} else {
		goto L161
	}
L151:
	;
	goto L150
L152:
	;
	v560 = *(*int32)(unsafe.Add(mBase, uint32(v544)+4))
	v561 = *(*int32)(unsafe.Add(mBase, uint32(v545)+4))
	if v560 != v561 {
		v592 = int32(0)
		goto L151
	} else {
		goto L153
	}
L153:
	;
	v563 = int32(1)
	if v560 <= v563 {
		goto L154
	} else {
		goto L155
	}
L154:
	;
	v566 = v563
	goto L156
L155:
	;
	v566 = v560
	goto L156
L156:
	;
	v567 = int32(8)
	v572 = int32(0)
	goto L157
L157:
	;
	v580 = v572 << (uint(int32(2)) % 32)
	v582 = *(*int32)(unsafe.Add(mBase, uint32(v544+v567+v580)))
	v584 = *(*int32)(unsafe.Add(mBase, uint32(v545+v567+v580)))
	v585 = base.B2i32(v582 == v584)
	if v582 != v584 {
		v592 = v585
		goto L151
	} else {
		goto L159
	}
L158:
	;
	v592 = v585
	goto L151
L159:
	;
	v588 = v572 + int32(1)
	if v588 != v566 {
		v572 = v588
		goto L157
	} else {
		goto L160
	}
L160:
	;
	goto L158
L161:
	;
	goto L104
L162:
	;
	v718 = *(*int32)(unsafe.Add(mBase, uint32(v43)+4))
	v719 = *(*int32)(unsafe.Add(mBase, uint32(l2)+8))
	v720 = int32(0)
	if v718 == v720 {
		goto L196
	} else {
		goto L197
	}
L163:
	;
	if v654 == int32(0) {
		goto L162
	} else {
		goto L177
	}
L164:
	;
	v654 = int32(1)
	goto L163
L165:
	;
	goto L166
L166:
	;
	if v600 == int32(0) {
		v647 = v601
		goto L167
	} else {
		goto L168
	}
L167:
	;
	v654 = v647
	goto L163
L168:
	;
	v610 = *(*int32)(unsafe.Add(mBase, uint32(v599)+4))
	v611 = *(*int32)(unsafe.Add(mBase, uint32(v600)+4))
	if v611 < v610 {
		v647 = v601
		goto L167
	} else {
		goto L169
	}
L169:
	;
	v613 = int32(1)
	if v610 <= v613 {
		goto L170
	} else {
		goto L171
	}
L170:
	;
	v616 = v613
	goto L172
L171:
	;
	v616 = v610
	goto L172
L172:
	;
	v617 = int32(8)
	v622 = int32(0)
	goto L173
L173:
	;
	v629 = v622 << (uint(int32(2)) % 32)
	v631 = *(*int32)(unsafe.Add(mBase, uint32(v599+v617+v629)))
	v633 = *(*int32)(unsafe.Add(mBase, uint32(v600+v617+v629)))
	v636 = v631 & (v633 ^ int32(-1))
	v638 = base.B2i32(v636 == int32(0))
	if v636 != 0 {
		v647 = v638
		goto L167
	} else {
		goto L175
	}
L174:
	;
	v647 = v638
	goto L167
L175:
	;
	v640 = v622 + int32(1)
	if v640 != v616 {
		v622 = v640
		goto L173
	} else {
		goto L176
	}
L176:
	;
	goto L174
L177:
	;
	v657 = *(*int32)(unsafe.Add(mBase, uint32(v43)+8))
	v658 = *(*int32)(unsafe.Add(mBase, uint32(l2)+8))
	v659 = int32(0)
	if v657 == v659 {
		goto L179
	} else {
		goto L180
	}
L178:
	;
	if v712 == int32(0) {
		goto L162
	} else {
		goto L192
	}
L179:
	;
	v712 = int32(1)
	goto L178
L180:
	;
	goto L181
L181:
	;
	if v658 == int32(0) {
		v705 = v659
		goto L182
	} else {
		goto L183
	}
L182:
	;
	v712 = v705
	goto L178
L183:
	;
	v668 = *(*int32)(unsafe.Add(mBase, uint32(v657)+4))
	v669 = *(*int32)(unsafe.Add(mBase, uint32(v658)+4))
	if v669 < v668 {
		v705 = v659
		goto L182
	} else {
		goto L184
	}
L184:
	;
	v671 = int32(1)
	if v668 <= v671 {
		goto L185
	} else {
		goto L186
	}
L185:
	;
	v674 = v671
	goto L187
L186:
	;
	v674 = v668
	goto L187
L187:
	;
	v675 = int32(8)
	v680 = int32(0)
	goto L188
L188:
	;
	v687 = v680 << (uint(int32(2)) % 32)
	v689 = *(*int32)(unsafe.Add(mBase, uint32(v657+v675+v687)))
	v691 = *(*int32)(unsafe.Add(mBase, uint32(v658+v675+v687)))
	v694 = v689 & (v691 ^ int32(-1))
	v696 = base.B2i32(v694 == int32(0))
	if v694 != 0 {
		v705 = v696
		goto L182
	} else {
		goto L190
	}
L189:
	;
	v705 = v696
	goto L182
L190:
	;
	v698 = v680 + int32(1)
	if v698 != v674 {
		v680 = v698
		goto L188
	} else {
		goto L191
	}
L191:
	;
	goto L189
L192:
	;
	v715 = int32(0)
	if v31 == v715 {
		v1121 = v43
		v1122 = v34
		v1123 = v715
		v1124 = v37
		goto L9
	} else {
		goto L193
	}
L193:
	;
	v1828 = v7
	goto L1
L194:
	;
	v839 = *(*int32)(unsafe.Add(mBase, uint32(v43)+20))
	if v839 != int32(4) {
		goto L228
	} else {
		goto L229
	}
L195:
	;
	if v773 == int32(0) {
		goto L194
	} else {
		goto L209
	}
L196:
	;
	v773 = int32(1)
	goto L195
L197:
	;
	goto L198
L198:
	;
	if v719 == int32(0) {
		v766 = v720
		goto L199
	} else {
		goto L200
	}
L199:
	;
	v773 = v766
	goto L195
L200:
	;
	v729 = *(*int32)(unsafe.Add(mBase, uint32(v718)+4))
	v730 = *(*int32)(unsafe.Add(mBase, uint32(v719)+4))
	if v730 < v729 {
		v766 = v720
		goto L199
	} else {
		goto L201
	}
L201:
	;
	v732 = int32(1)
	if v729 <= v732 {
		goto L202
	} else {
		goto L203
	}
L202:
	;
	v735 = v732
	goto L204
L203:
	;
	v735 = v729
	goto L204
L204:
	;
	v736 = int32(8)
	v741 = int32(0)
	goto L205
L205:
	;
	v748 = v741 << (uint(int32(2)) % 32)
	v750 = *(*int32)(unsafe.Add(mBase, uint32(v718+v736+v748)))
	v752 = *(*int32)(unsafe.Add(mBase, uint32(v719+v736+v748)))
	v755 = v750 & (v752 ^ int32(-1))
	v757 = base.B2i32(v755 == int32(0))
	if v755 != 0 {
		v766 = v757
		goto L199
	} else {
		goto L207
	}
L206:
	;
	v766 = v757
	goto L199
L207:
	;
	v759 = v741 + int32(1)
	if v759 != v735 {
		v741 = v759
		goto L205
	} else {
		goto L208
	}
L208:
	;
	goto L206
L209:
	;
	v776 = *(*int32)(unsafe.Add(mBase, uint32(v43)+8))
	v777 = *(*int32)(unsafe.Add(mBase, uint32(l1)+8))
	v778 = int32(0)
	if v776 == v778 {
		goto L211
	} else {
		goto L212
	}
L210:
	;
	if v831 == int32(0) {
		goto L194
	} else {
		goto L224
	}
L211:
	;
	v831 = int32(1)
	goto L210
L212:
	;
	goto L213
L213:
	;
	if v777 == int32(0) {
		v824 = v778
		goto L214
	} else {
		goto L215
	}
L214:
	;
	v831 = v824
	goto L210
L215:
	;
	v787 = *(*int32)(unsafe.Add(mBase, uint32(v776)+4))
	v788 = *(*int32)(unsafe.Add(mBase, uint32(v777)+4))
	if v788 < v787 {
		v824 = v778
		goto L214
	} else {
		goto L216
	}
L216:
	;
	v790 = int32(1)
	if v787 <= v790 {
		goto L217
	} else {
		goto L218
	}
L217:
	;
	v793 = v790
	goto L219
L218:
	;
	v793 = v787
	goto L219
L219:
	;
	v794 = int32(8)
	v799 = int32(0)
	goto L220
L220:
	;
	v806 = v799 << (uint(int32(2)) % 32)
	v808 = *(*int32)(unsafe.Add(mBase, uint32(v776+v794+v806)))
	v810 = *(*int32)(unsafe.Add(mBase, uint32(v777+v794+v806)))
	v813 = v808 & (v810 ^ int32(-1))
	v815 = base.B2i32(v813 == int32(0))
	if v813 != 0 {
		v824 = v815
		goto L214
	} else {
		goto L222
	}
L221:
	;
	v824 = v815
	goto L214
L222:
	;
	v817 = v799 + int32(1)
	if v817 != v793 {
		v799 = v817
		goto L220
	} else {
		goto L223
	}
L223:
	;
	goto L221
L224:
	;
	if v31 == int32(0) {
		goto L225
	} else {
		goto L226
	}
L225:
	;
	v1121 = v43
	v1122 = v34
	v1123 = int32(1)
	v1124 = v37
	goto L9
L226:
	;
	goto L227
L227:
	;
	return int32(0)
L228:
	;
	v973 = *(*int32)(unsafe.Add(mBase, uint32(l1)+8))
	v974 = *(*int32)(unsafe.Add(mBase, uint32(v43)+8))
	v975 = int32(0)
	if base.B2i32(v973 == v975)|base.B2i32(v974 == v975) != 0 {
		v1020 = v975
		goto L268
	} else {
		goto L269
	}
L229:
	;
	v842 = *(*int32)(unsafe.Add(mBase, uint32(v43)+16))
	v843 = *(*int32)(unsafe.Add(mBase, uint32(l2)+8))
	v844 = int32(0)
	if base.B2i32(v842 == v844)|base.B2i32(v843 == v844) != 0 {
		v890 = base.B2i32(v842|v843 == v844)
		goto L232
	} else {
		goto L233
	}
L230:
	;
	v907 = *(*int32)(unsafe.Add(mBase, uint32(v43)+20))
	if v907 != int32(4) {
		goto L228
	} else {
		goto L249
	}
L231:
	;
	if v890 == int32(0) {
		goto L230
	} else {
		goto L242
	}
L232:
	;
	goto L231
L233:
	;
	v858 = *(*int32)(unsafe.Add(mBase, uint32(v842)+4))
	v859 = *(*int32)(unsafe.Add(mBase, uint32(v843)+4))
	if v858 != v859 {
		v890 = int32(0)
		goto L232
	} else {
		goto L234
	}
L234:
	;
	v861 = int32(1)
	if v858 <= v861 {
		goto L235
	} else {
		goto L236
	}
L235:
	;
	v864 = v861
	goto L237
L236:
	;
	v864 = v858
	goto L237
L237:
	;
	v865 = int32(8)
	v870 = int32(0)
	goto L238
L238:
	;
	v878 = v870 << (uint(int32(2)) % 32)
	v880 = *(*int32)(unsafe.Add(mBase, uint32(v842+v865+v878)))
	v882 = *(*int32)(unsafe.Add(mBase, uint32(v843+v865+v878)))
	v883 = base.B2i32(v880 == v882)
	if v880 != v882 {
		v890 = v883
		goto L232
	} else {
		goto L240
	}
L239:
	;
	v890 = v883
	goto L232
L240:
	;
	v886 = v870 + int32(1)
	if v886 != v864 {
		v870 = v886
		goto L238
	} else {
		goto L241
	}
L241:
	;
	goto L239
L242:
	;
	v897 = F_create_unique_paths(m, l0, l2, v43)
	mBase = m.M
	v900 = m.ExcPending
	if v900 != 0 {
		goto L243
	} else {
		goto L244
	}
L243:
	;
	return int32(0)
L244:
	;
	if v897 == int32(0) {
		goto L230
	} else {
		goto L245
	}
L245:
	;
	if v31 != 0 {
		goto L246
	} else {
		goto L247
	}
L246:
	;
	return int32(0)
L247:
	;
	goto L248
L248:
	;
	v1121 = v43
	v1122 = v34
	v1123 = int32(0)
	v1124 = int32(1)
	goto L9
L249:
	;
	v910 = *(*int32)(unsafe.Add(mBase, uint32(v43)+16))
	v911 = *(*int32)(unsafe.Add(mBase, uint32(l1)+8))
	v912 = int32(0)
	if base.B2i32(v910 == v912)|base.B2i32(v911 == v912) != 0 {
		v958 = base.B2i32(v910|v911 == v912)
		goto L251
	} else {
		goto L252
	}
L250:
	;
	if v958 == int32(0) {
		goto L228
	} else {
		goto L261
	}
L251:
	;
	goto L250
L252:
	;
	v926 = *(*int32)(unsafe.Add(mBase, uint32(v910)+4))
	v927 = *(*int32)(unsafe.Add(mBase, uint32(v911)+4))
	if v926 != v927 {
		v958 = int32(0)
		goto L251
	} else {
		goto L253
	}
L253:
	;
	v929 = int32(1)
	if v926 <= v929 {
		goto L254
	} else {
		goto L255
	}
L254:
	;
	v932 = v929
	goto L256
L255:
	;
	v932 = v926
	goto L256
L256:
	;
	v933 = int32(8)
	v938 = int32(0)
	goto L257
L257:
	;
	v946 = v938 << (uint(int32(2)) % 32)
	v948 = *(*int32)(unsafe.Add(mBase, uint32(v910+v933+v946)))
	v950 = *(*int32)(unsafe.Add(mBase, uint32(v911+v933+v946)))
	v951 = base.B2i32(v948 == v950)
	if v948 != v950 {
		v958 = v951
		goto L251
	} else {
		goto L259
	}
L258:
	;
	v958 = v951
	goto L251
L259:
	;
	v954 = v938 + int32(1)
	if v954 != v932 {
		v938 = v954
		goto L257
	} else {
		goto L260
	}
L260:
	;
	goto L258
L261:
	;
	v965 = F_create_unique_paths(m, l0, l1, v43)
	mBase = m.M
	v966 = m.ExcPending
	if v966 != 0 {
		goto L243
	} else {
		goto L262
	}
L262:
	;
	if v965 == int32(0) {
		goto L228
	} else {
		goto L263
	}
L263:
	;
	if v31 != 0 {
		goto L264
	} else {
		goto L265
	}
L264:
	;
	return int32(0)
L265:
	;
	goto L266
L266:
	;
	v971 = int32(1)
	v1121 = v43
	v1122 = v34
	v1123 = v971
	v1124 = v971
	goto L9
L267:
	;
	if v1020 != 0 {
		goto L280
	} else {
		goto L281
	}
L268:
	;
	goto L267
L269:
	;
	v985 = *(*int32)(unsafe.Add(mBase, uint32(v973)+4))
	v986 = *(*int32)(unsafe.Add(mBase, uint32(v974)+4))
	if v985 < v986 {
		goto L270
	} else {
		goto L271
	}
L270:
	;
	v988 = v985
	goto L272
L271:
	;
	v988 = v986
	goto L272
L272:
	;
	if v988 <= int32(1) {
		goto L273
	} else {
		goto L274
	}
L273:
	;
	v991 = int32(1)
	goto L275
L274:
	;
	v991 = v988
	goto L275
L275:
	;
	v992 = int32(8)
	v997 = int32(0)
	goto L276
L276:
	;
	v1004 = v997 << (uint(int32(2)) % 32)
	v1006 = *(*int32)(unsafe.Add(mBase, uint32(v974+v992+v1004)))
	v1008 = *(*int32)(unsafe.Add(mBase, uint32(v973+v992+v1004)))
	v1009 = v1006 & v1008
	v1011 = base.B2i32(v1009 != int32(0))
	if v1009 != 0 {
		v1020 = v1011
		goto L268
	} else {
		goto L278
	}
L277:
	;
	v1020 = v1011
	goto L268
L278:
	;
	v1013 = v997 + int32(1)
	if v1013 != v991 {
		v997 = v1013
		goto L276
	} else {
		goto L279
	}
L279:
	;
	goto L277
L280:
	;
	v1021 = *(*int32)(unsafe.Add(mBase, uint32(l2)+8))
	v1022 = *(*int32)(unsafe.Add(mBase, uint32(v43)+8))
	v1023 = int32(0)
	if base.B2i32(v1021 == v1023)|base.B2i32(v1022 == v1023) != 0 {
		v1068 = v1023
		goto L284
	} else {
		goto L285
	}
L281:
	;
	goto L282
L282:
	;
	v1069 = *(*int32)(unsafe.Add(mBase, uint32(v43)+20))
	if v1069 != int32(1) {
		v1828 = v7
		goto L1
	} else {
		goto L297
	}
L283:
	;
	if v1068 != 0 {
		v1120 = v34
		goto L10
	} else {
		goto L296
	}
L284:
	;
	goto L283
L285:
	;
	v1033 = *(*int32)(unsafe.Add(mBase, uint32(v1021)+4))
	v1034 = *(*int32)(unsafe.Add(mBase, uint32(v1022)+4))
	if v1033 < v1034 {
		goto L286
	} else {
		goto L287
	}
L286:
	;
	v1036 = v1033
	goto L288
L287:
	;
	v1036 = v1034
	goto L288
L288:
	;
	if v1036 <= int32(1) {
		goto L289
	} else {
		goto L290
	}
L289:
	;
	v1039 = int32(1)
	goto L291
L290:
	;
	v1039 = v1036
	goto L291
L291:
	;
	v1040 = int32(8)
	v1045 = int32(0)
	goto L292
L292:
	;
	v1052 = v1045 << (uint(int32(2)) % 32)
	v1054 = *(*int32)(unsafe.Add(mBase, uint32(v1022+v1040+v1052)))
	v1056 = *(*int32)(unsafe.Add(mBase, uint32(v1021+v1040+v1052)))
	v1057 = v1054 & v1056
	v1059 = base.B2i32(v1057 != int32(0))
	if v1057 != 0 {
		v1068 = v1059
		goto L284
	} else {
		goto L294
	}
L293:
	;
	v1068 = v1059
	goto L284
L294:
	;
	v1061 = v1045 + int32(1)
	if v1061 != v1039 {
		v1045 = v1061
		goto L292
	} else {
		goto L295
	}
L295:
	;
	goto L293
L296:
	;
	goto L282
L297:
	;
	v1072 = *(*int32)(unsafe.Add(mBase, uint32(v43)+4))
	v1073 = int32(0)
	if base.B2i32(l3 == v1073)|base.B2i32(v1072 == v1073) != 0 {
		v1118 = v1073
		goto L299
	} else {
		goto L300
	}
L298:
	;
	if v1118 != 0 {
		v1828 = v7
		goto L1
	} else {
		goto L311
	}
L299:
	;
	goto L298
L300:
	;
	v1083 = *(*int32)(unsafe.Add(mBase, uint32(l3)+4))
	v1084 = *(*int32)(unsafe.Add(mBase, uint32(v1072)+4))
	if v1083 < v1084 {
		goto L301
	} else {
		goto L302
	}
L301:
	;
	v1086 = v1083
	goto L303
L302:
	;
	v1086 = v1084
	goto L303
L303:
	;
	if v1086 <= int32(1) {
		goto L304
	} else {
		goto L305
	}
L304:
	;
	v1089 = int32(1)
	goto L306
L305:
	;
	v1089 = v1086
	goto L306
L306:
	;
	v1090 = int32(8)
	v1095 = int32(0)
	goto L307
L307:
	;
	v1102 = v1095 << (uint(int32(2)) % 32)
	v1104 = *(*int32)(unsafe.Add(mBase, uint32(v1072+v1090+v1102)))
	v1106 = *(*int32)(unsafe.Add(mBase, uint32(l3+v1090+v1102)))
	v1107 = v1104 & v1106
	v1109 = base.B2i32(v1107 != int32(0))
	if v1107 != 0 {
		v1118 = v1109
		goto L299
	} else {
		goto L309
	}
L308:
	;
	v1118 = v1109
	goto L299
L309:
	;
	v1111 = v1095 + int32(1)
	if v1111 != v1089 {
		v1095 = v1111
		goto L307
	} else {
		goto L310
	}
L310:
	;
	goto L308
L311:
	;
	v1120 = int32(1)
	goto L10
L312:
	;
	goto L8
L313:
	;
	if v1135 == int32(0) {
		v1828 = v7
		goto L1
	} else {
		goto L314
	}
L314:
	;
	v1147 = *(*int32)(unsafe.Add(mBase, uint32(v1135)+20))
	if v1147 != int32(1) {
		v1828 = v7
		goto L1
	} else {
		goto L315
	}
L315:
	;
	v1150 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1135)+44)))
	if v1150 != int32(1) {
		v1828 = v7
		goto L1
	} else {
		goto L316
	}
L316:
	;
	v1159 = v1135
	v1163 = v1139
	v1165 = v1141
	goto L2
L317:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l4))) = v1159
	*(*uint8)(unsafe.Add(mBase, uint32(l5))) = uint8(v1163)
	v1828 = int32(1)
	goto L1
L318:
	;
	v1170 = *(*int32)(unsafe.Add(mBase, uint32(l1)+8))
	v1171 = *(*int32)(unsafe.Add(mBase, uint32(l2)+72))
	v1172 = int32(0)
	if base.B2i32(v1170 == v1172)|base.B2i32(v1171 == v1172) != 0 {
		v1217 = v1172
		goto L320
	} else {
		goto L321
	}
L319:
	;
	v1218 = int32(0)
	v1219 = *(*int32)(unsafe.Add(mBase, uint32(l2)+8))
	v1220 = *(*int32)(unsafe.Add(mBase, uint32(l1)+72))
	if base.B2i32(v1219 == v1218)|base.B2i32(v1220 == v1218) != 0 {
		v1266 = v1218
		goto L333
	} else {
		goto L334
	}
L320:
	;
	goto L319
L321:
	;
	v1182 = *(*int32)(unsafe.Add(mBase, uint32(v1170)+4))
	v1183 = *(*int32)(unsafe.Add(mBase, uint32(v1171)+4))
	if v1182 < v1183 {
		goto L322
	} else {
		goto L323
	}
L322:
	;
	v1185 = v1182
	goto L324
L323:
	;
	v1185 = v1183
	goto L324
L324:
	;
	if v1185 <= int32(1) {
		goto L325
	} else {
		goto L326
	}
L325:
	;
	v1188 = int32(1)
	goto L327
L326:
	;
	v1188 = v1185
	goto L327
L327:
	;
	v1189 = int32(8)
	v1194 = int32(0)
	goto L328
L328:
	;
	v1201 = v1194 << (uint(int32(2)) % 32)
	v1203 = *(*int32)(unsafe.Add(mBase, uint32(v1171+v1189+v1201)))
	v1205 = *(*int32)(unsafe.Add(mBase, uint32(v1170+v1189+v1201)))
	v1206 = v1203 & v1205
	v1208 = base.B2i32(v1206 != int32(0))
	if v1206 != 0 {
		v1217 = v1208
		goto L320
	} else {
		goto L330
	}
L329:
	;
	v1217 = v1208
	goto L320
L330:
	;
	v1210 = v1194 + int32(1)
	if v1210 != v1188 {
		v1194 = v1210
		goto L328
	} else {
		goto L331
	}
L331:
	;
	goto L329
L332:
	;
	if v1266 != 0 {
		goto L345
	} else {
		goto L346
	}
L333:
	;
	goto L332
L334:
	;
	v1231 = *(*int32)(unsafe.Add(mBase, uint32(v1219)+4))
	v1232 = *(*int32)(unsafe.Add(mBase, uint32(v1220)+4))
	if v1231 < v1232 {
		goto L335
	} else {
		goto L336
	}
L335:
	;
	v1234 = v1231
	goto L337
L336:
	;
	v1234 = v1232
	goto L337
L337:
	;
	if v1234 <= int32(1) {
		goto L338
	} else {
		goto L339
	}
L338:
	;
	v1237 = int32(1)
	goto L340
L339:
	;
	v1237 = v1234
	goto L340
L340:
	;
	v1238 = int32(8)
	v1243 = int32(0)
	goto L341
L341:
	;
	v1250 = v1243 << (uint(int32(2)) % 32)
	v1252 = *(*int32)(unsafe.Add(mBase, uint32(v1220+v1238+v1250)))
	v1254 = *(*int32)(unsafe.Add(mBase, uint32(v1219+v1238+v1250)))
	v1255 = v1252 & v1254
	v1257 = base.B2i32(v1255 != int32(0))
	if v1255 != 0 {
		v1266 = v1257
		goto L333
	} else {
		goto L343
	}
L342:
	;
	v1266 = v1257
	goto L333
L343:
	;
	v1259 = v1243 + int32(1)
	if v1259 != v1237 {
		v1243 = v1259
		goto L341
	} else {
		goto L344
	}
L344:
	;
	goto L342
L345:
	;
	v1267 = v1217
	goto L347
L346:
	;
	v1267 = v1218
	goto L347
L347:
	;
	if v1267 != 0 {
		v1828 = v7
		goto L1
	} else {
		goto L348
	}
L348:
	;
	if v1217 != 0 {
		goto L350
	} else {
		goto L351
	}
L349:
	;
	v1380 = *(*int32)(unsafe.Add(mBase, uint32(l1)+72))
	v1381 = *(*int32)(unsafe.Add(mBase, uint32(l2)+72))
	v1382 = F_bms_union(m, v1380, v1381)
	mBase = m.M
	v1383 = m.ExcPending
	if v1383 != 0 {
		goto L243
	} else {
		goto L392
	}
L350:
	;
	if v1159 != 0 {
		goto L353
	} else {
		goto L354
	}
L351:
	;
	goto L352
L352:
	;
	if v1266 == int32(0) {
		goto L349
	} else {
		goto L372
	}
L353:
	;
	if v1163|v1165 != 0 {
		v1828 = v7
		goto L1
	} else {
		goto L356
	}
L354:
	;
	goto L355
L355:
	;
	v1272 = *(*int32)(unsafe.Add(mBase, uint32(l1)+8))
	v1273 = *(*int32)(unsafe.Add(mBase, uint32(l2)+68))
	v1274 = int32(0)
	if base.B2i32(v1272 == v1274)|base.B2i32(v1273 == v1274) != 0 {
		v1319 = v1274
		goto L359
	} else {
		goto L360
	}
L356:
	;
	v1269 = *(*int32)(unsafe.Add(mBase, uint32(v1159)+20))
	if v1269 == int32(2) {
		v1828 = v7
		goto L1
	} else {
		goto L357
	}
L357:
	;
	goto L355
L358:
	;
	if v1319 != 0 {
		goto L349
	} else {
		goto L371
	}
L359:
	;
	goto L358
L360:
	;
	v1284 = *(*int32)(unsafe.Add(mBase, uint32(v1272)+4))
	v1285 = *(*int32)(unsafe.Add(mBase, uint32(v1273)+4))
	if v1284 < v1285 {
		goto L361
	} else {
		goto L362
	}
L361:
	;
	v1287 = v1284
	goto L363
L362:
	;
	v1287 = v1285
	goto L363
L363:
	;
	if v1287 <= int32(1) {
		goto L364
	} else {
		goto L365
	}
L364:
	;
	v1290 = int32(1)
	goto L366
L365:
	;
	v1290 = v1287
	goto L366
L366:
	;
	v1291 = int32(8)
	v1296 = int32(0)
	goto L367
L367:
	;
	v1303 = v1296 << (uint(int32(2)) % 32)
	v1305 = *(*int32)(unsafe.Add(mBase, uint32(v1273+v1291+v1303)))
	v1307 = *(*int32)(unsafe.Add(mBase, uint32(v1272+v1291+v1303)))
	v1308 = v1305 & v1307
	v1310 = base.B2i32(v1308 != int32(0))
	if v1308 != 0 {
		v1319 = v1310
		goto L359
	} else {
		goto L369
	}
L368:
	;
	v1319 = v1310
	goto L359
L369:
	;
	v1312 = v1296 + int32(1)
	if v1312 != v1290 {
		v1296 = v1312
		goto L367
	} else {
		goto L370
	}
L370:
	;
	goto L368
L371:
	;
	v1828 = v7
	goto L1
L372:
	;
	if v1159 != 0 {
		goto L373
	} else {
		goto L374
	}
L373:
	;
	if (v1163^int32(-1)|v1165)&int32(1) != 0 {
		v1828 = v7
		goto L1
	} else {
		goto L376
	}
L374:
	;
	goto L375
L375:
	;
	v1330 = *(*int32)(unsafe.Add(mBase, uint32(l2)+8))
	v1331 = *(*int32)(unsafe.Add(mBase, uint32(l1)+68))
	v1332 = int32(0)
	if base.B2i32(v1330 == v1332)|base.B2i32(v1331 == v1332) != 0 {
		v1377 = v1332
		goto L379
	} else {
		goto L380
	}
L376:
	;
	v1327 = *(*int32)(unsafe.Add(mBase, uint32(v1159)+20))
	if v1327 == int32(2) {
		v1828 = v7
		goto L1
	} else {
		goto L377
	}
L377:
	;
	goto L375
L378:
	;
	if v1377 == int32(0) {
		v1828 = v7
		goto L1
	} else {
		goto L391
	}
L379:
	;
	goto L378
L380:
	;
	v1342 = *(*int32)(unsafe.Add(mBase, uint32(v1330)+4))
	v1343 = *(*int32)(unsafe.Add(mBase, uint32(v1331)+4))
	if v1342 < v1343 {
		goto L381
	} else {
		goto L382
	}
L381:
	;
	v1345 = v1342
	goto L383
L382:
	;
	v1345 = v1343
	goto L383
L383:
	;
	if v1345 <= int32(1) {
		goto L384
	} else {
		goto L385
	}
L384:
	;
	v1348 = int32(1)
	goto L386
L385:
	;
	v1348 = v1345
	goto L386
L386:
	;
	v1349 = int32(8)
	v1354 = int32(0)
	goto L387
L387:
	;
	v1361 = v1354 << (uint(int32(2)) % 32)
	v1363 = *(*int32)(unsafe.Add(mBase, uint32(v1331+v1349+v1361)))
	v1365 = *(*int32)(unsafe.Add(mBase, uint32(v1330+v1349+v1361)))
	v1366 = v1363 & v1365
	v1368 = base.B2i32(v1366 != int32(0))
	if v1366 != 0 {
		v1377 = v1368
		goto L379
	} else {
		goto L389
	}
L388:
	;
	v1377 = v1368
	goto L379
L389:
	;
	v1370 = v1354 + int32(1)
	if v1370 != v1348 {
		v1354 = v1370
		goto L387
	} else {
		goto L390
	}
L390:
	;
	goto L388
L391:
	;
	goto L349
L392:
	;
	v1384 = F_bms_del_members(m, v1382, l3)
	mBase = m.M
	v1385 = m.ExcPending
	if v1385 != 0 {
		goto L243
	} else {
		goto L393
	}
L393:
	;
	if v1384 == int32(0) {
		goto L317
	} else {
		goto L394
	}
L394:
	;
	v1388 = F_bms_copy(m, l3)
	mBase = m.M
	v1389 = m.ExcPending
	if v1389 != 0 {
		goto L243
	} else {
		goto L395
	}
L395:
	;
	v1393 = v1388
	goto L396
L396:
	;
	v1404 = *(*int32)(unsafe.Add(mBase, uint32(l0)+120))
	if v1404 != 0 {
		goto L398
	} else {
		goto L399
	}
L397:
	;
	v1576 = int32(0)
	if base.B2i32(v1565 == v1576)|base.B2i32(v1384 == v1576) != 0 {
		v1621 = v1576
		goto L441
	} else {
		goto L442
	}
L398:
	;
	v1405 = int32(0)
	v1407 = *(*int32)(unsafe.Add(mBase, uint32(v1404)+4))
	if v1405 < v1407 {
		goto L401
	} else {
		goto L402
	}
L399:
	;
	v1565 = v1393
	goto L400
L400:
	;
	goto L397
L401:
	;
	v1411 = v1405
	v1413 = v1393
	v1418 = v1405
	goto L404
L402:
	;
	v1547 = v1405
	v1549 = v1393
	goto L403
L403:
	;
	if v1547&int32(1) != 0 {
		v1393 = v1549
		goto L396
	} else {
		goto L439
	}
L404:
	;
	v1424 = *(*int32)(unsafe.Add(mBase, uint32(v1404)+12))
	v1425 = int32(2)
	v1428 = *(*int32)(unsafe.Add(mBase, uint32(v1424+v1418<<(uint(v1425)%32))))
	v1429 = *(*int32)(unsafe.Add(mBase, uint32(v1428)+20))
	if v1429 == v1425 {
		v1540 = v1411
		v1541 = v1413
		goto L406
	} else {
		goto L407
	}
L405:
	;
	v1547 = v1540
	v1549 = v1541
	goto L403
L406:
	;
	v1543 = v1418 + int32(1)
	v1544 = *(*int32)(unsafe.Add(mBase, uint32(v1404)+4))
	if v1543 < v1544 {
		v1411 = v1540
		v1413 = v1541
		v1418 = v1543
		goto L404
	} else {
		goto L438
	}
L407:
	;
	v1432 = *(*int32)(unsafe.Add(mBase, uint32(v1428)+4))
	v1433 = int32(0)
	if base.B2i32(v1432 == v1433)|base.B2i32(v1413 == v1433) != 0 {
		v1478 = v1433
		goto L409
	} else {
		goto L410
	}
L408:
	;
	if v1478 == int32(0) {
		v1540 = v1411
		v1541 = v1413
		goto L406
	} else {
		goto L421
	}
L409:
	;
	goto L408
L410:
	;
	v1443 = *(*int32)(unsafe.Add(mBase, uint32(v1432)+4))
	v1444 = *(*int32)(unsafe.Add(mBase, uint32(v1413)+4))
	if v1443 < v1444 {
		goto L411
	} else {
		goto L412
	}
L411:
	;
	v1446 = v1443
	goto L413
L412:
	;
	v1446 = v1444
	goto L413
L413:
	;
	if v1446 <= int32(1) {
		goto L414
	} else {
		goto L415
	}
L414:
	;
	v1449 = int32(1)
	goto L416
L415:
	;
	v1449 = v1446
	goto L416
L416:
	;
	v1450 = int32(8)
	v1455 = int32(0)
	goto L417
L417:
	;
	v1462 = v1455 << (uint(int32(2)) % 32)
	v1464 = *(*int32)(unsafe.Add(mBase, uint32(v1413+v1450+v1462)))
	v1466 = *(*int32)(unsafe.Add(mBase, uint32(v1432+v1450+v1462)))
	v1467 = v1464 & v1466
	v1469 = base.B2i32(v1467 != int32(0))
	if v1467 != 0 {
		v1478 = v1469
		goto L409
	} else {
		goto L419
	}
L418:
	;
	v1478 = v1469
	goto L409
L419:
	;
	v1471 = v1455 + int32(1)
	if v1471 != v1449 {
		v1455 = v1471
		goto L417
	} else {
		goto L420
	}
L420:
	;
	goto L418
L421:
	;
	v1481 = *(*int32)(unsafe.Add(mBase, uint32(v1428)+8))
	v1482 = int32(0)
	if v1481 == v1482 {
		goto L423
	} else {
		goto L424
	}
L422:
	;
	if v1535 != 0 {
		v1540 = v1411
		v1541 = v1413
		goto L406
	} else {
		goto L436
	}
L423:
	;
	v1535 = int32(1)
	goto L422
L424:
	;
	goto L425
L425:
	;
	if v1413 == int32(0) {
		v1528 = v1482
		goto L426
	} else {
		goto L427
	}
L426:
	;
	v1535 = v1528
	goto L422
L427:
	;
	v1491 = *(*int32)(unsafe.Add(mBase, uint32(v1481)+4))
	v1492 = *(*int32)(unsafe.Add(mBase, uint32(v1413)+4))
	if v1492 < v1491 {
		v1528 = v1482
		goto L426
	} else {
		goto L428
	}
L428:
	;
	v1494 = int32(1)
	if v1491 <= v1494 {
		goto L429
	} else {
		goto L430
	}
L429:
	;
	v1497 = v1494
	goto L431
L430:
	;
	v1497 = v1491
	goto L431
L431:
	;
	v1498 = int32(8)
	v1503 = int32(0)
	goto L432
L432:
	;
	v1510 = v1503 << (uint(int32(2)) % 32)
	v1512 = *(*int32)(unsafe.Add(mBase, uint32(v1481+v1498+v1510)))
	v1514 = *(*int32)(unsafe.Add(mBase, uint32(v1413+v1498+v1510)))
	v1517 = v1512 & (v1514 ^ int32(-1))
	v1519 = base.B2i32(v1517 == int32(0))
	if v1517 != 0 {
		v1528 = v1519
		goto L426
	} else {
		goto L434
	}
L433:
	;
	v1528 = v1519
	goto L426
L434:
	;
	v1521 = v1503 + int32(1)
	if v1521 != v1497 {
		v1503 = v1521
		goto L432
	} else {
		goto L435
	}
L435:
	;
	goto L433
L436:
	;
	v1537 = *(*int32)(unsafe.Add(mBase, uint32(v1428)+8))
	v1538 = F_bms_add_members(m, v1413, v1537)
	mBase = m.M
	v1539 = m.ExcPending
	if v1539 != 0 {
		goto L243
	} else {
		goto L437
	}
L437:
	;
	v1540 = int32(1)
	v1541 = v1538
	goto L406
L438:
	;
	goto L405
L439:
	;
	v1565 = v1549
	goto L400
L440:
	;
	if v1621 != 0 {
		v1828 = v7
		goto L1
	} else {
		goto L453
	}
L441:
	;
	goto L440
L442:
	;
	v1586 = *(*int32)(unsafe.Add(mBase, uint32(v1565)+4))
	v1587 = *(*int32)(unsafe.Add(mBase, uint32(v1384)+4))
	if v1586 < v1587 {
		goto L443
	} else {
		goto L444
	}
L443:
	;
	v1589 = v1586
	goto L445
L444:
	;
	v1589 = v1587
	goto L445
L445:
	;
	if v1589 <= int32(1) {
		goto L446
	} else {
		goto L447
	}
L446:
	;
	v1592 = int32(1)
	goto L448
L447:
	;
	v1592 = v1589
	goto L448
L448:
	;
	v1593 = int32(8)
	v1598 = int32(0)
	goto L449
L449:
	;
	v1605 = v1598 << (uint(int32(2)) % 32)
	v1607 = *(*int32)(unsafe.Add(mBase, uint32(v1384+v1593+v1605)))
	v1609 = *(*int32)(unsafe.Add(mBase, uint32(v1565+v1593+v1605)))
	v1610 = v1607 & v1609
	v1612 = base.B2i32(v1610 != int32(0))
	if v1610 != 0 {
		v1621 = v1612
		goto L441
	} else {
		goto L451
	}
L450:
	;
	v1621 = v1612
	goto L441
L451:
	;
	v1614 = v1598 + int32(1)
	if v1614 != v1592 {
		v1598 = v1614
		goto L449
	} else {
		goto L452
	}
L452:
	;
	goto L450
L453:
	;
	v1622 = *(*int32)(unsafe.Add(mBase, uint32(l0)+56))
	v1623 = int32(0)
	if base.B2i32(v1384 == v1623)|base.B2i32(v1622 == v1623) != 0 {
		v1668 = v1623
		goto L455
	} else {
		goto L456
	}
L454:
	;
	if v1668 == int32(0) {
		goto L317
	} else {
		goto L467
	}
L455:
	;
	goto L454
L456:
	;
	v1633 = *(*int32)(unsafe.Add(mBase, uint32(v1384)+4))
	v1634 = *(*int32)(unsafe.Add(mBase, uint32(v1622)+4))
	if v1633 < v1634 {
		goto L457
	} else {
		goto L458
	}
L457:
	;
	v1636 = v1633
	goto L459
L458:
	;
	v1636 = v1634
	goto L459
L459:
	;
	if v1636 <= int32(1) {
		goto L460
	} else {
		goto L461
	}
L460:
	;
	v1639 = int32(1)
	goto L462
L461:
	;
	v1639 = v1636
	goto L462
L462:
	;
	v1640 = int32(8)
	v1645 = int32(0)
	goto L463
L463:
	;
	v1652 = v1645 << (uint(int32(2)) % 32)
	v1654 = *(*int32)(unsafe.Add(mBase, uint32(v1622+v1640+v1652)))
	v1656 = *(*int32)(unsafe.Add(mBase, uint32(v1384+v1640+v1652)))
	v1657 = v1654 & v1656
	v1659 = base.B2i32(v1657 != int32(0))
	if v1657 != 0 {
		v1668 = v1659
		goto L455
	} else {
		goto L465
	}
L464:
	;
	v1668 = v1659
	goto L455
L465:
	;
	v1661 = v1645 + int32(1)
	if v1661 != v1639 {
		v1645 = v1661
		goto L463
	} else {
		goto L466
	}
L466:
	;
	goto L464
L467:
	;
	v1671 = *(*int32)(unsafe.Add(mBase, uint32(l0)+120))
	if v1671 == int32(0) {
		goto L317
	} else {
		goto L468
	}
L468:
	;
	v1674 = *(*int32)(unsafe.Add(mBase, uint32(v1671)+4))
	if v1674 <= int32(0) {
		goto L317
	} else {
		goto L469
	}
L469:
	;
	v1686 = int32(0)
	goto L470
L470:
	;
	v1692 = *(*int32)(unsafe.Add(mBase, uint32(v1671)+12))
	v1696 = *(*int32)(unsafe.Add(mBase, uint32(v1692+v1686<<(uint(int32(2))%32))))
	v1697 = *(*int32)(unsafe.Add(mBase, uint32(v1696)+24))
	v1698 = F_bms_is_member(m, v1697, v1384)
	mBase = m.M
	v1699 = m.ExcPending
	if v1699 != 0 {
		goto L243
	} else {
		goto L472
	}
L471:
	;
	goto L317
L472:
	;
	if v1698 != 0 {
		goto L473
	} else {
		goto L474
	}
L473:
	;
	v1700 = *(*int32)(unsafe.Add(mBase, uint32(v1696)+4))
	v1701 = int32(0)
	if base.B2i32(v1565 == v1701)|base.B2i32(v1700 == v1701) != 0 {
		v1746 = v1701
		goto L477
	} else {
		goto L478
	}
L474:
	;
	goto L475
L475:
	;
	v1795 = v1686 + int32(1)
	v1796 = *(*int32)(unsafe.Add(mBase, uint32(v1671)+4))
	if v1795 < v1796 {
		v1686 = v1795
		goto L470
	} else {
		goto L504
	}
L476:
	;
	if v1746 != 0 {
		v1828 = v7
		goto L1
	} else {
		goto L489
	}
L477:
	;
	goto L476
L478:
	;
	v1711 = *(*int32)(unsafe.Add(mBase, uint32(v1565)+4))
	v1712 = *(*int32)(unsafe.Add(mBase, uint32(v1700)+4))
	if v1711 < v1712 {
		goto L479
	} else {
		goto L480
	}
L479:
	;
	v1714 = v1711
	goto L481
L480:
	;
	v1714 = v1712
	goto L481
L481:
	;
	if v1714 <= int32(1) {
		goto L482
	} else {
		goto L483
	}
L482:
	;
	v1717 = int32(1)
	goto L484
L483:
	;
	v1717 = v1714
	goto L484
L484:
	;
	v1718 = int32(8)
	v1723 = int32(0)
	goto L485
L485:
	;
	v1730 = v1723 << (uint(int32(2)) % 32)
	v1732 = *(*int32)(unsafe.Add(mBase, uint32(v1700+v1718+v1730)))
	v1734 = *(*int32)(unsafe.Add(mBase, uint32(v1565+v1718+v1730)))
	v1735 = v1732 & v1734
	v1737 = base.B2i32(v1735 != int32(0))
	if v1735 != 0 {
		v1746 = v1737
		goto L477
	} else {
		goto L487
	}
L486:
	;
	v1746 = v1737
	goto L477
L487:
	;
	v1739 = v1723 + int32(1)
	if v1739 != v1717 {
		v1723 = v1739
		goto L485
	} else {
		goto L488
	}
L488:
	;
	goto L486
L489:
	;
	v1747 = *(*int32)(unsafe.Add(mBase, uint32(v1696)+8))
	v1748 = int32(0)
	if base.B2i32(v1565 == v1748)|base.B2i32(v1747 == v1748) != 0 {
		v1793 = v1748
		goto L491
	} else {
		goto L492
	}
L490:
	;
	if v1793 != 0 {
		v1828 = v7
		goto L1
	} else {
		goto L503
	}
L491:
	;
	goto L490
L492:
	;
	v1758 = *(*int32)(unsafe.Add(mBase, uint32(v1565)+4))
	v1759 = *(*int32)(unsafe.Add(mBase, uint32(v1747)+4))
	if v1758 < v1759 {
		goto L493
	} else {
		goto L494
	}
L493:
	;
	v1761 = v1758
	goto L495
L494:
	;
	v1761 = v1759
	goto L495
L495:
	;
	if v1761 <= int32(1) {
		goto L496
	} else {
		goto L497
	}
L496:
	;
	v1764 = int32(1)
	goto L498
L497:
	;
	v1764 = v1761
	goto L498
L498:
	;
	v1765 = int32(8)
	v1770 = int32(0)
	goto L499
L499:
	;
	v1777 = v1770 << (uint(int32(2)) % 32)
	v1779 = *(*int32)(unsafe.Add(mBase, uint32(v1747+v1765+v1777)))
	v1781 = *(*int32)(unsafe.Add(mBase, uint32(v1565+v1765+v1777)))
	v1782 = v1779 & v1781
	v1784 = base.B2i32(v1782 != int32(0))
	if v1782 != 0 {
		v1793 = v1784
		goto L491
	} else {
		goto L501
	}
L500:
	;
	v1793 = v1784
	goto L491
L501:
	;
	v1786 = v1770 + int32(1)
	if v1786 != v1764 {
		v1770 = v1786
		goto L499
	} else {
		goto L502
	}
L502:
	;
	goto L500
L503:
	;
	goto L475
L504:
	;
	goto L471
}
func F_remove_join_from_jointree(m *base.Module, l0 int32, l1 int32, l2 int32) int32 {
	mBase := m.M
	_ = mBase
	var v4 int32
	_ = v4
	var v8 int32
	_ = v8
	var v10 int32
	_ = v10
	var v15 int32
	_ = v15
	var v18 int32
	_ = v18
	var v21 int32
	_ = v21
	var v29 int32
	_ = v29
	var v31 int32
	_ = v31
	var v34 int32
	_ = v34
	var v35 int32
	_ = v35
	var v36 int32
	_ = v36
	var v39 int32
	_ = v39
	var v42 int32
	_ = v42
	var v43 int32
	_ = v43
	var v45 int32
	_ = v45
	var v47 int32
	_ = v47
	var v48 int32
	_ = v48
	var v49 int32
	_ = v49
	var v51 int32
	_ = v51
	var v52 int32
	_ = v52
	var v53 int32
	_ = v53
	var v55 int32
	_ = v55
	var v59 int32
	_ = v59
	var v63 int32
	_ = v63
	var v64 int32
	_ = v64
	var v68 int32
	_ = v68
	var v73 int32
	_ = v73
	var v74 int32
	_ = v74
	v4 = int32(0)
	v8 = m.G0
	v10 = v8 - int32(16)
	m.G0 = v10
	if l0 == v4 {
		goto L2
	} else {
		goto L3
	}
L1:
	;
	m.G0 = v10 + int32(16)
	return v74
L2:
	;
	v74 = int32(0)
	goto L1
L3:
	;
	goto L4
L4:
	;
	v15 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	switch v15 - int32(63) {
	case 0:
		v74 = l0
		goto L1
	case 1:
		goto L6
	case 2:
		goto L7
	default:
		goto L5
	}
L5:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v63 = m.ExcPending
	if v63 != 0 {
		goto L12
	} else {
		goto L20
	}
L6:
	;
	v45 = *(*int32)(unsafe.Add(mBase, uint32(l0)+36))
	if l1 != v45 {
		goto L15
	} else {
		goto L16
	}
L7:
	;
	v18 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	if v18 == int32(0) {
		v74 = l0
		goto L1
	} else {
		goto L8
	}
L8:
	;
	v21 = *(*int32)(unsafe.Add(mBase, uint32(v18)+4))
	if v21 <= int32(0) {
		v74 = l0
		goto L1
	} else {
		goto L9
	}
L9:
	;
	v29 = v4
	goto L10
L10:
	;
	v31 = *(*int32)(unsafe.Add(mBase, uint32(v18)+12))
	v34 = v31 + v29<<(uint(int32(2))%32)
	v35 = *(*int32)(unsafe.Add(mBase, uint32(v34)))
	v36 = F_remove_join_from_jointree(m, v35, l1, l2)
	mBase = m.M
	v39 = m.ExcPending
	if v39 != 0 {
		goto L12
	} else {
		goto L13
	}
L11:
	;
	v74 = l0
	goto L1
L12:
	;
	return int32(0)
L13:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v34))) = v36
	v42 = v29 + int32(1)
	v43 = *(*int32)(unsafe.Add(mBase, uint32(v18)+4))
	if v42 < v43 {
		v29 = v42
		goto L10
	} else {
		goto L14
	}
L14:
	;
	goto L11
L15:
	;
	v47 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	v48 = F_remove_join_from_jointree(m, v47, l1, l2)
	mBase = m.M
	v49 = m.ExcPending
	if v49 != 0 {
		goto L12
	} else {
		goto L18
	}
L16:
	;
	goto L17
L17:
	;
	v55 = *(*int32)(unsafe.Add(mBase, uint32(l2)))
	*(*int32)(unsafe.Add(mBase, uint32(l2))) = v55 + int32(1)
	v59 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	v74 = v59
	goto L1
L18:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+12)) = v48
	v51 = *(*int32)(unsafe.Add(mBase, uint32(l0)+16))
	v52 = F_remove_join_from_jointree(m, v51, l1, l2)
	mBase = m.M
	v53 = m.ExcPending
	if v53 != 0 {
		goto L12
	} else {
		goto L19
	}
L19:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+16)) = v52
	v74 = l0
	goto L1
L20:
	;
	v64 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	*(*int32)(unsafe.Add(mBase, uint32(v10))) = v64
	F_errmsg_internal(m, int32(_a_F_remove_join_from_jointree_0), v10)
	mBase = m.M
	v68 = m.ExcPending
	if v68 != 0 {
		goto L12
	} else {
		goto L21
	}
L21:
	;
	F_errfinish(m, int32(_a_F_remove_join_from_jointree_1), int32(406), int32(_a_F_remove_join_from_jointree_2))
	mBase = m.M
	v73 = m.ExcPending
	if v73 != 0 {
		goto L12
	} else {
		goto L22
	}
L22:
	;
	base.Wasm_trap_unreachable()
	for {
	}
}
