package p2

import (
	base "github.com/shibukawa/pgmem/internal/aot/pgaot/base"
	"unsafe"
)

func F_replace_text_regexp(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32, l5 int32, l6 int32) int32 {
	mBase := m.M
	_ = mBase
	var v8 int32
	_ = v8
	var v24 int32
	_ = v24
	var v26 int32
	_ = v26
	var v28 int32
	_ = v28
	var v31 int32
	_ = v31
	var v33 int32
	_ = v33
	var v42 int32
	_ = v42
	var v45 int32
	_ = v45
	var v46 int32
	_ = v46
	var v52 int32
	_ = v52
	var v58 int32
	_ = v58
	var v64 int32
	_ = v64
	var v69 int32
	_ = v69
	var v70 int32
	_ = v70
	var v71 int32
	_ = v71
	var v72 int32
	_ = v72
	var v74 int32
	_ = v74
	var v75 int32
	_ = v75
	var v78 int32
	_ = v78
	var v79 int32
	_ = v79
	var v80 int32
	_ = v80
	var v81 int32
	_ = v81
	var v82 int32
	_ = v82
	var v84 int32
	_ = v84
	var v85 int32
	_ = v85
	var v87 int32
	_ = v87
	var v88 int32
	_ = v88
	var v91 int32
	_ = v91
	var v93 int32
	_ = v93
	var v102 int32
	_ = v102
	var v105 int32
	_ = v105
	var v106 int32
	_ = v106
	var v110 int32
	_ = v110
	var v116 int32
	_ = v116
	var v118 int32
	_ = v118
	var v127 int32
	_ = v127
	var v134 int32
	_ = v134
	var v144 int32
	_ = v144
	var v145 int32
	_ = v145
	var v148 int32
	_ = v148
	var v157 int32
	_ = v157
	var v159 int32
	_ = v159
	var v162 int32
	_ = v162
	var v164 int32
	_ = v164
	var v165 int32
	_ = v165
	var v166 int32
	_ = v166
	var v167 int32
	_ = v167
	var v169 int32
	_ = v169
	var v174 int32
	_ = v174
	var v176 int32
	_ = v176
	var v177 int32
	_ = v177
	var v181 int32
	_ = v181
	var v191 int32
	_ = v191
	var v193 int32
	_ = v193
	var v197 int32
	_ = v197
	var v198 int32
	_ = v198
	var v201 int32
	_ = v201
	var v205 int32
	_ = v205
	var v206 int32
	_ = v206
	var v208 int32
	_ = v208
	var v211 int32
	_ = v211
	var v213 int32
	_ = v213
	var v218 int32
	_ = v218
	var v220 int32
	_ = v220
	var v225 int32
	_ = v225
	var v227 int32
	_ = v227
	var v230 int32
	_ = v230
	var v232 int32
	_ = v232
	var v235 int32
	_ = v235
	var v247 int32
	_ = v247
	var v251 int32
	_ = v251
	var v253 int32
	_ = v253
	var v264 int32
	_ = v264
	var v265 int32
	_ = v265
	var v281 int32
	_ = v281
	var v293 int32
	_ = v293
	var v294 int32
	_ = v294
	var v295 int32
	_ = v295
	var v296 int32
	_ = v296
	var v297 int32
	_ = v297
	var v300 int32
	_ = v300
	var v304 int32
	_ = v304
	var v309 int32
	_ = v309
	var v312 int32
	_ = v312
	var v314 int32
	_ = v314
	var v315 int32
	_ = v315
	var v316 int32
	_ = v316
	var v335 int32
	_ = v335
	var v337 int32
	_ = v337
	var v340 int32
	_ = v340
	var v341 int32
	_ = v341
	var v346 int32
	_ = v346
	var v347 int32
	_ = v347
	var v351 int32
	_ = v351
	var v354 int32
	_ = v354
	var v360 int32
	_ = v360
	var v365 int32
	_ = v365
	var v367 int32
	_ = v367
	var v371 int32
	_ = v371
	var v372 int32
	_ = v372
	var v373 int32
	_ = v373
	var v379 int32
	_ = v379
	var v380 int32
	_ = v380
	var v385 int32
	_ = v385
	var v395 int32
	_ = v395
	var v396 int32
	_ = v396
	var v411 int32
	_ = v411
	var v412 int32
	_ = v412
	var v413 int32
	_ = v413
	var v414 int32
	_ = v414
	var v426 int32
	_ = v426
	var v443 int32
	_ = v443
	var v444 int32
	_ = v444
	var v449 int32
	_ = v449
	var v450 int32
	_ = v450
	var v469 int32
	_ = v469
	var v471 int32
	_ = v471
	var v472 int32
	_ = v472
	var v475 int32
	_ = v475
	var v477 int32
	_ = v477
	var v486 int32
	_ = v486
	var v489 int32
	_ = v489
	var v490 int32
	_ = v490
	var v494 int32
	_ = v494
	var v500 int32
	_ = v500
	var v501 int32
	_ = v501
	var v508 int32
	_ = v508
	var v527 int32
	_ = v527
	var v528 int32
	_ = v528
	var v531 int32
	_ = v531
	var v540 int32
	_ = v540
	var v542 int32
	_ = v542
	var v545 int32
	_ = v545
	var v547 int32
	_ = v547
	var v548 int32
	_ = v548
	var v549 int32
	_ = v549
	var v550 int32
	_ = v550
	var v552 int32
	_ = v552
	var v557 int32
	_ = v557
	var v559 int32
	_ = v559
	var v560 int32
	_ = v560
	var v564 int32
	_ = v564
	var v574 int32
	_ = v574
	var v576 int32
	_ = v576
	var v580 int32
	_ = v580
	var v581 int32
	_ = v581
	var v584 int32
	_ = v584
	var v588 int32
	_ = v588
	var v589 int32
	_ = v589
	var v591 int32
	_ = v591
	var v594 int32
	_ = v594
	var v596 int32
	_ = v596
	var v601 int32
	_ = v601
	var v603 int32
	_ = v603
	var v608 int32
	_ = v608
	var v610 int32
	_ = v610
	var v613 int32
	_ = v613
	var v615 int32
	_ = v615
	var v618 int32
	_ = v618
	var v630 int32
	_ = v630
	var v631 int32
	_ = v631
	var v637 int32
	_ = v637
	var v640 int32
	_ = v640
	var v646 int32
	_ = v646
	var v647 int32
	_ = v647
	var v658 int32
	_ = v658
	var v673 int32
	_ = v673
	var v680 int32
	_ = v680
	var v681 int32
	_ = v681
	var v682 int32
	_ = v682
	var v684 int32
	_ = v684
	var v685 int32
	_ = v685
	var v688 int32
	_ = v688
	var v691 int32
	_ = v691
	var v693 int32
	_ = v693
	var v694 int32
	_ = v694
	var v699 int32
	_ = v699
	var v711 int32
	_ = v711
	var v712 int32
	_ = v712
	var v727 int32
	_ = v727
	var v728 int32
	_ = v728
	var v729 int32
	_ = v729
	var v730 int32
	_ = v730
	var v742 int32
	_ = v742
	var v765 int32
	_ = v765
	var v781 int32
	_ = v781
	var v784 int32
	_ = v784
	var v786 int32
	_ = v786
	var v787 int32
	_ = v787
	var v792 int32
	_ = v792
	var v804 int32
	_ = v804
	var v805 int32
	_ = v805
	var v820 int32
	_ = v820
	var v821 int32
	_ = v821
	var v822 int32
	_ = v822
	var v823 int32
	_ = v823
	var v834 int32
	_ = v834
	var v874 int32
	_ = v874
	var v876 int32
	_ = v876
	var v882 int32
	_ = v882
	var v903 int32
	_ = v903
	var v906 int32
	_ = v906
	var v908 int32
	_ = v908
	var v917 int32
	_ = v917
	var v920 int32
	_ = v920
	var v921 int32
	_ = v921
	var v925 int32
	_ = v925
	var v931 int32
	_ = v931
	var v933 int32
	_ = v933
	var v957 int32
	_ = v957
	var v958 int32
	_ = v958
	var v960 int32
	_ = v960
	var v961 int32
	_ = v961
	var v966 int32
	_ = v966
	var v978 int32
	_ = v978
	var v979 int32
	_ = v979
	var v994 int32
	_ = v994
	var v995 int32
	_ = v995
	var v996 int32
	_ = v996
	var v997 int32
	_ = v997
	var v1009 int32
	_ = v1009
	var v1048 int32
	_ = v1048
	var v1049 int32
	_ = v1049
	var v1050 int32
	_ = v1050
	var v1056 int32
	_ = v1056
	var v1057 int32
	_ = v1057
	var v1061 int32
	_ = v1061
	var v1076 int32
	_ = v1076
	var v1078 int32
	_ = v1078
	var v1083 int32
	_ = v1083
	var v1088 int32
	_ = v1088
	var v1106 int32
	_ = v1106
	var v1109 int32
	_ = v1109
	var v1111 int32
	_ = v1111
	var v1113 int32
	_ = v1113
	var v1116 int32
	_ = v1116
	var v1121 int32
	_ = v1121
	var v1124 int32
	_ = v1124
	var v1125 int32
	_ = v1125
	var v1129 int32
	_ = v1129
	var v1133 int32
	_ = v1133
	var v1137 int32
	_ = v1137
	var v1139 int32
	_ = v1139
	var v1140 int32
	_ = v1140
	var v1142 int32
	_ = v1142
	var v1143 int32
	_ = v1143
	var v1144 int32
	_ = v1144
	var v1150 int32
	_ = v1150
	var v1152 int32
	_ = v1152
	var v1154 int32
	_ = v1154
	var v1156 int32
	_ = v1156
	v8 = int32(0)
	v24 = m.G0
	v26 = v24 - int32(224)
	m.G0 = v26
	v28 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0))))
	if v28 == int32(1) {
		goto L2
	} else {
		goto L3
	}
L1:
	;
	F_initStringInfo(m, v26+int32(208))
	mBase = m.M
	v64 = m.ExcPending
	if v64 != 0 {
		goto L12
	} else {
		goto L13
	}
L2:
	;
	v31 = int32(4)
	v33 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+1)))
	if v33&int32(254) == int32(2) {
		goto L5
	} else {
		goto L6
	}
L3:
	;
	goto L4
L4:
	;
	v46 = int32(1)
	if v28&v46 != 0 {
		v58 = int32(base.Ui32(v28)>>(uint(v46)%32)) - v46
		goto L1
	} else {
		goto L11
	}
L5:
	;
	v42 = v31
	goto L7
L6:
	;
	v42 = base.B2i32(v33 == int32(18)) << (uint(v31) % 32)
	goto L7
L7:
	;
	if v33 == int32(1) {
		goto L8
	} else {
		goto L9
	}
L8:
	;
	v45 = v31
	goto L10
L9:
	;
	v45 = v42
	goto L10
L10:
	;
	v58 = v45
	goto L1
L11:
	;
	v52 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v58 = int32(base.Ui32(v52)>>(uint(int32(2))%32)) - int32(4)
	goto L1
L12:
	;
	return int32(0)
L13:
	;
	v69 = F_palloc(m, v58<<(uint(int32(2))%32)+int32(4))
	mBase = m.M
	v70 = m.ExcPending
	if v70 != 0 {
		goto L12
	} else {
		goto L14
	}
L14:
	;
	v71 = int32(1)
	v72 = l0 + v71
	v74 = l0 + int32(4)
	v75 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0))))
	if v75&v71 != 0 {
		goto L15
	} else {
		goto L16
	}
L15:
	;
	v78 = v72
	goto L17
L16:
	;
	v78 = v74
	goto L17
L17:
	;
	v79 = F_pg_mb2wchar_with_len(m, v78, v69, v58)
	mBase = m.M
	v80 = m.ExcPending
	if v80 != 0 {
		goto L12
	} else {
		goto L18
	}
L18:
	;
	v81 = int32(1)
	v82 = l2 + v81
	v84 = l2 + int32(4)
	v85 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l2))))
	v87 = v85 & v81
	if v87 != 0 {
		goto L19
	} else {
		goto L20
	}
L19:
	;
	v88 = v82
	goto L21
L20:
	;
	v88 = v84
	goto L21
L21:
	;
	if v85 == int32(1) {
		goto L23
	} else {
		goto L24
	}
L22:
	;
	v118 = v88 + v116
	if base.Ui32(v118) <= base.Ui32(v88) {
		v281 = v8
		goto L33
	} else {
		goto L34
	}
L23:
	;
	v91 = int32(4)
	v93 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v82))))
	if v93&int32(254) == int32(2) {
		goto L26
	} else {
		goto L27
	}
L24:
	;
	goto L25
L25:
	;
	v106 = int32(1)
	if v87 != 0 {
		v116 = int32(base.Ui32(v85)>>(uint(v106)%32)) - v106
		goto L22
	} else {
		goto L32
	}
L26:
	;
	v102 = v91
	goto L28
L27:
	;
	v102 = base.B2i32(v93 == int32(18)) << (uint(v91) % 32)
	goto L28
L28:
	;
	if v93 == int32(1) {
		goto L29
	} else {
		goto L30
	}
L29:
	;
	v105 = v91
	goto L31
L30:
	;
	v105 = v102
	goto L31
L31:
	;
	v116 = v105
	goto L22
L32:
	;
	v110 = *(*int32)(unsafe.Add(mBase, uint32(l2)))
	v116 = int32(base.Ui32(v110)>>(uint(int32(2))%32)) - int32(4)
	goto L22
L33:
	;
	v293 = base.B2i32(base.Ui32(v281) < base.Ui32(int32(2)))
	if base.Ui32(v281) < base.Ui32(int32(2)) {
		goto L71
	} else {
		goto L72
	}
L34:
	;
	v127 = v88
	v134 = v8
	goto L35
L35:
	;
	v144 = v118 - v127
	v145 = int32(0)
	v148 = base.B2i32(v144 != v145)
	if v127&int32(3) == v145 {
		v174 = v127
		v176 = v144
		v177 = v148
		goto L40
	} else {
		goto L41
	}
L36:
	;
	v281 = v265
	goto L33
L37:
	;
	if v247 == int32(0) {
		v281 = v134
		goto L33
	} else {
		goto L63
	}
L38:
	;
	v247 = int32(0)
	goto L37
L39:
	;
	v225 = v218
	v227 = v220
	goto L57
L40:
	;
	if v177 == int32(0) {
		goto L38
	} else {
		goto L48
	}
L41:
	;
	if v144 == int32(0) {
		v174 = v127
		v176 = v144
		v177 = v148
		goto L40
	} else {
		goto L42
	}
L42:
	;
	v157 = v127
	v159 = v144
	goto L43
L43:
	;
	v162 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v157))))
	if v162 == int32(92) {
		v218 = v157
		v220 = v159
		goto L39
	} else {
		goto L45
	}
L44:
	;
	v174 = v169
	v176 = v165
	v177 = v167
	goto L40
L45:
	;
	v164 = int32(1)
	v165 = v159 - v164
	v166 = int32(0)
	v167 = base.B2i32(v165 != v166)
	v169 = v157 + v164
	if v169&int32(3) == v166 {
		v174 = v169
		v176 = v165
		v177 = v167
		goto L40
	} else {
		goto L46
	}
L46:
	;
	if v165 != 0 {
		v157 = v169
		v159 = v165
		goto L43
	} else {
		goto L47
	}
L47:
	;
	goto L44
L48:
	;
	v181 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v174))))
	if v181 == int32(92) {
		v211 = v174
		v213 = v176
		goto L49
	} else {
		goto L50
	}
L49:
	;
	if v213 == int32(0) {
		goto L38
	} else {
		goto L56
	}
L50:
	;
	if base.Ui32(v176) < base.Ui32(int32(4)) {
		v211 = v174
		v213 = v176
		goto L49
	} else {
		goto L51
	}
L51:
	;
	v191 = v174
	v193 = v176
	goto L52
L52:
	;
	v197 = *(*int32)(unsafe.Add(mBase, uint32(v191)))
	v198 = v197 ^ int32(1549556828)
	v201 = int32(-2139062144)
	if (int32(16843008)-v198|v198)&v201 != v201 {
		v218 = v191
		v220 = v193
		goto L39
	} else {
		goto L54
	}
L53:
	;
	v211 = v206
	v213 = v208
	goto L49
L54:
	;
	v205 = int32(4)
	v206 = v191 + v205
	v208 = v193 - v205
	if base.Ui32(int32(3)) < base.Ui32(v208) {
		v191 = v206
		v193 = v208
		goto L52
	} else {
		goto L55
	}
L55:
	;
	goto L53
L56:
	;
	v218 = v211
	v220 = v213
	goto L39
L57:
	;
	v230 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v225))))
	if int32(92) == v230 {
		goto L59
	} else {
		goto L60
	}
L58:
	;
	goto L38
L59:
	;
	v247 = v225
	goto L37
L60:
	;
	goto L61
L61:
	;
	v232 = int32(1)
	v235 = v227 - v232
	if v235 != 0 {
		v225 = v225 + v232
		v227 = v235
		goto L57
	} else {
		goto L62
	}
L62:
	;
	goto L58
L63:
	;
	v251 = v247 + int32(1)
	if base.Ui32(v251) < base.Ui32(v118) {
		goto L64
	} else {
		goto L65
	}
L64:
	;
	v253 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v251))))
	if base.Ui32((v253-int32(49))&int32(255)) < base.Ui32(int32(9)) {
		goto L67
	} else {
		goto L68
	}
L65:
	;
	v264 = v251
	v265 = v134
	goto L66
L66:
	;
	if base.Ui32(v264) < base.Ui32(v118) {
		v127 = v264
		v134 = v265
		goto L35
	} else {
		goto L70
	}
L67:
	;
	v281 = int32(2)
	goto L33
L68:
	;
	goto L69
L69:
	;
	v264 = v247 + int32(2)
	v265 = int32(1)
	goto L66
L70:
	;
	goto L36
L71:
	;
	v294 = l3 | int32(16)
	goto L73
L72:
	;
	v294 = l3
	goto L73
L73:
	;
	v295 = F_RE_compile_and_cache(m, l1, v294, l4)
	mBase = m.M
	v296 = m.ExcPending
	if v296 != 0 {
		goto L12
	} else {
		goto L74
	}
L74:
	;
	v297 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0))))
	if v297&int32(1) != 0 {
		goto L75
	} else {
		goto L76
	}
L75:
	;
	v300 = v72
	goto L77
L76:
	;
	v300 = v74
	goto L77
L77:
	;
	if base.Ui32(v79) < base.Ui32(l5) {
		v1083 = v300
		v1088 = int32(0)
		goto L78
	} else {
		goto L79
	}
L78:
	;
	if base.Ui32(v1088) < base.Ui32(v79) {
		goto L241
	} else {
		goto L242
	}
L79:
	;
	if base.Ui32(v281) < base.Ui32(int32(2)) {
		goto L80
	} else {
		goto L81
	}
L80:
	;
	v304 = int32(1)
	goto L82
L81:
	;
	v304 = int32(10)
	goto L82
L82:
	;
	v309 = int32(0)
	v312 = v309
	v314 = v300
	v315 = v309
	v316 = l5
	goto L83
L83:
	;
	v335 = *(*int32)(unsafe.Add(mBase, _consts[1]))
	if v335 != 0 {
		goto L85
	} else {
		goto L86
	}
L84:
	;
	v1083 = v1056
	v1088 = v1057
	goto L78
L85:
	;
	F_ProcessInterrupts(m)
	mBase = m.M
	v337 = m.ExcPending
	if v337 != 0 {
		goto L12
	} else {
		goto L88
	}
L86:
	;
	goto L87
L87:
	;
	v340 = F_pg_regexec(m, v295, v69, v79, v316, v304, v26+int32(128))
	mBase = m.M
	v341 = m.ExcPending
	if v341 != 0 {
		goto L12
	} else {
		goto L89
	}
L88:
	;
	goto L87
L89:
	;
	if v340 != 0 {
		goto L90
	} else {
		goto L91
	}
L90:
	;
	if v340 == int32(1) {
		v1083 = v314
		v1088 = v315
		goto L78
	} else {
		goto L93
	}
L91:
	;
	goto L92
L92:
	;
	v367 = v312 + int32(1)
	if l6 <= int32(0) {
		goto L100
	} else {
		goto L101
	}
L93:
	;
	v346 = F_pg_regerror(m, v340, v26+int32(16))
	mBase = m.M
	v347 = m.ExcPending
	if v347 != 0 {
		goto L12
	} else {
		goto L94
	}
L94:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v351 = m.ExcPending
	if v351 != 0 {
		goto L12
	} else {
		goto L95
	}
L95:
	;
	F_errcode(m, int32(302252162))
	mBase = m.M
	v354 = m.ExcPending
	if v354 != 0 {
		goto L12
	} else {
		goto L96
	}
L96:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v26))) = v26 + int32(16)
	F_errmsg(m, int32(202274), v26)
	mBase = m.M
	v360 = m.ExcPending
	if v360 != 0 {
		goto L12
	} else {
		goto L97
	}
L97:
	;
	F_errfinish(m, int32(494617), int32(4531), int32(229320))
	mBase = m.M
	v365 = m.ExcPending
	if v365 != 0 {
		goto L12
	} else {
		goto L98
	}
L98:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L99:
	;
	v1076 = *(*int32)(unsafe.Add(mBase, uint32(v26)+128))
	v1078 = v1061 + base.B2i32(v1076 == v1061)
	if base.Ui32(v1078) <= base.Ui32(v79) {
		v312 = v367
		v314 = v1056
		v315 = v1057
		v316 = v1078
		goto L83
	} else {
		goto L240
	}
L100:
	;
	v372 = *(*int32)(unsafe.Add(mBase, uint32(v26)+128))
	v373 = v372 - v315
	if int32(0) < v373 {
		goto L103
	} else {
		goto L104
	}
L101:
	;
	if v367 == l6 {
		goto L100
	} else {
		goto L102
	}
L102:
	;
	v371 = *(*int32)(unsafe.Add(mBase, uint32(v26)+132))
	v1056 = v314
	v1057 = v315
	v1061 = v371
	goto L99
L103:
	;
	v379 = *(*int32)(unsafe.Add(mBase, _consts[358]))
	v380 = *(*int32)(unsafe.Add(mBase, uint32(v379)+4))
	v385 = *(*int32)(unsafe.Add(mBase, uint32(v380*int32(28))+uint32(_consts[1003])))
	goto L106
L104:
	;
	v449 = v314
	v450 = v315
	goto L105
L105:
	;
	v469 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l2))))
	v471 = v469 & int32(1)
	if v281 != 0 {
		goto L116
	} else {
		goto L117
	}
L106:
	;
	if v385 != int32(1) {
		goto L107
	} else {
		goto L108
	}
L107:
	;
	v395 = v373
	v396 = v314
	goto L110
L108:
	;
	v426 = v373
	goto L109
L109:
	;
	F_appendBinaryStringInfo(m, v26+int32(208), v314, v426)
	mBase = m.M
	v443 = m.ExcPending
	if v443 != 0 {
		goto L12
	} else {
		goto L114
	}
L110:
	;
	v411 = F_pg_mblen_unbounded(m, v396)
	mBase = m.M
	v412 = m.ExcPending
	if v412 != 0 {
		goto L12
	} else {
		goto L112
	}
L111:
	;
	v426 = v413 - v314
	goto L109
L112:
	;
	v413 = v411 + v396
	v414 = int32(1)
	if base.Ui32(v414) < base.Ui32(v395) {
		v395 = v395 - v414
		v396 = v413
		goto L110
	} else {
		goto L113
	}
L113:
	;
	goto L111
L114:
	;
	v444 = *(*int32)(unsafe.Add(mBase, uint32(v26)+128))
	v449 = v314 + v426
	v450 = v444
	goto L105
L115:
	;
	v957 = *(*int32)(unsafe.Add(mBase, uint32(v26)+132))
	v958 = v957 - v450
	v960 = *(*int32)(unsafe.Add(mBase, _consts[358]))
	v961 = *(*int32)(unsafe.Add(mBase, uint32(v960)+4))
	v966 = *(*int32)(unsafe.Add(mBase, uint32(v961*int32(28))+uint32(_consts[1003])))
	goto L228
L116:
	;
	if v471 != 0 {
		goto L119
	} else {
		goto L120
	}
L117:
	;
	goto L118
L118:
	;
	if v471 != 0 {
		goto L213
	} else {
		goto L214
	}
L119:
	;
	v472 = v82
	goto L121
L120:
	;
	v472 = v84
	goto L121
L121:
	;
	if v469 == int32(1) {
		goto L123
	} else {
		goto L124
	}
L122:
	;
	v501 = v472 + v500
	if base.Ui32(v501) <= base.Ui32(v472) {
		goto L115
	} else {
		goto L133
	}
L123:
	;
	v475 = int32(4)
	v477 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v82))))
	if v477&int32(254) == int32(2) {
		goto L126
	} else {
		goto L127
	}
L124:
	;
	goto L125
L125:
	;
	v490 = int32(1)
	if v471 != 0 {
		v500 = int32(base.Ui32(v469)>>(uint(v490)%32)) - v490
		goto L122
	} else {
		goto L132
	}
L126:
	;
	v486 = v475
	goto L128
L127:
	;
	v486 = base.B2i32(v477 == int32(18)) << (uint(v475) % 32)
	goto L128
L128:
	;
	if v477 == int32(1) {
		goto L129
	} else {
		goto L130
	}
L129:
	;
	v489 = v475
	goto L131
L130:
	;
	v489 = v486
	goto L131
L131:
	;
	v500 = v489
	goto L122
L132:
	;
	v494 = *(*int32)(unsafe.Add(mBase, uint32(l2)))
	v500 = int32(base.Ui32(v494)>>(uint(int32(2))%32)) - int32(4)
	goto L122
L133:
	;
	v508 = v472
	goto L134
L134:
	;
	v527 = v501 - v508
	v528 = int32(0)
	v531 = base.B2i32(v527 != v528)
	if v508&int32(3) == v528 {
		v557 = v508
		v559 = v527
		v560 = v531
		goto L139
	} else {
		goto L140
	}
L135:
	;
	goto L115
L136:
	;
	if v630 != 0 {
		goto L162
	} else {
		goto L163
	}
L137:
	;
	v630 = int32(0)
	goto L136
L138:
	;
	v608 = v601
	v610 = v603
	goto L156
L139:
	;
	if v560 == int32(0) {
		goto L137
	} else {
		goto L147
	}
L140:
	;
	if v527 == int32(0) {
		v557 = v508
		v559 = v527
		v560 = v531
		goto L139
	} else {
		goto L141
	}
L141:
	;
	v540 = v508
	v542 = v527
	goto L142
L142:
	;
	v545 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v540))))
	if v545 == int32(92) {
		v601 = v540
		v603 = v542
		goto L138
	} else {
		goto L144
	}
L143:
	;
	v557 = v552
	v559 = v548
	v560 = v550
	goto L139
L144:
	;
	v547 = int32(1)
	v548 = v542 - v547
	v549 = int32(0)
	v550 = base.B2i32(v548 != v549)
	v552 = v540 + v547
	if v552&int32(3) == v549 {
		v557 = v552
		v559 = v548
		v560 = v550
		goto L139
	} else {
		goto L145
	}
L145:
	;
	if v548 != 0 {
		v540 = v552
		v542 = v548
		goto L142
	} else {
		goto L146
	}
L146:
	;
	goto L143
L147:
	;
	v564 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v557))))
	if v564 == int32(92) {
		v594 = v557
		v596 = v559
		goto L148
	} else {
		goto L149
	}
L148:
	;
	if v596 == int32(0) {
		goto L137
	} else {
		goto L155
	}
L149:
	;
	if base.Ui32(v559) < base.Ui32(int32(4)) {
		v594 = v557
		v596 = v559
		goto L148
	} else {
		goto L150
	}
L150:
	;
	v574 = v557
	v576 = v559
	goto L151
L151:
	;
	v580 = *(*int32)(unsafe.Add(mBase, uint32(v574)))
	v581 = v580 ^ int32(1549556828)
	v584 = int32(-2139062144)
	if (int32(16843008)-v581|v581)&v584 != v584 {
		v601 = v574
		v603 = v576
		goto L138
	} else {
		goto L153
	}
L152:
	;
	v594 = v589
	v596 = v591
	goto L148
L153:
	;
	v588 = int32(4)
	v589 = v574 + v588
	v591 = v576 - v588
	if base.Ui32(int32(3)) < base.Ui32(v591) {
		v574 = v589
		v576 = v591
		goto L151
	} else {
		goto L154
	}
L154:
	;
	goto L152
L155:
	;
	v601 = v594
	v603 = v596
	goto L138
L156:
	;
	v613 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v608))))
	if int32(92) == v613 {
		goto L158
	} else {
		goto L159
	}
L157:
	;
	goto L137
L158:
	;
	v630 = v608
	goto L136
L159:
	;
	goto L160
L160:
	;
	v615 = int32(1)
	v618 = v610 - v615
	if v618 != 0 {
		v608 = v608 + v615
		v610 = v618
		goto L156
	} else {
		goto L161
	}
L161:
	;
	goto L157
L162:
	;
	v631 = v630
	goto L164
L163:
	;
	v631 = v501
	goto L164
L164:
	;
	if base.Ui32(v508) < base.Ui32(v631) {
		goto L165
	} else {
		goto L166
	}
L165:
	;
	F_appendBinaryStringInfo(m, v26+int32(208), v508, v631-v508)
	mBase = m.M
	v637 = m.ExcPending
	if v637 != 0 {
		goto L12
	} else {
		goto L168
	}
L166:
	;
	goto L167
L167:
	;
	if base.Ui32(v501) <= base.Ui32(v631) {
		goto L115
	} else {
		goto L169
	}
L168:
	;
	goto L167
L169:
	;
	v640 = v631 + int32(1)
	if base.Ui32(v501) <= base.Ui32(v640) {
		goto L170
	} else {
		goto L171
	}
L170:
	;
	F_appendStringInfoChar(m, v26+int32(208), int32(92))
	mBase = m.M
	v646 = m.ExcPending
	if v646 != 0 {
		goto L12
	} else {
		goto L173
	}
L171:
	;
	goto L172
L172:
	;
	v647 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v640))))
	if base.Ui32((v647-int32(49))&int32(255)) <= base.Ui32(int32(8)) {
		goto L176
	} else {
		goto L177
	}
L173:
	;
	goto L115
L174:
	;
	if base.Ui32(v882) < base.Ui32(v501) {
		v508 = v882
		goto L134
	} else {
		goto L212
	}
L175:
	;
	v684 = v631 + int32(2)
	v685 = *(*int32)(unsafe.Add(mBase, uint32(v681)))
	if v685 < int32(0) {
		v882 = v684
		goto L174
	} else {
		goto L187
	}
L176:
	;
	v658 = v26 + int32(128) + v647<<(uint(int32(3))%32)
	v681 = v658 - int32(384)
	v682 = v658 - int32(380)
	goto L175
L177:
	;
	goto L178
L178:
	;
	if v647 == int32(38) {
		goto L179
	} else {
		goto L180
	}
L179:
	;
	v681 = v26 + int32(128)
	v682 = v26 + int32(128) | int32(4)
	goto L175
L180:
	;
	goto L181
L181:
	;
	if v647 == int32(92) {
		goto L182
	} else {
		goto L183
	}
L182:
	;
	F_appendStringInfoChar(m, v26+int32(208), int32(92))
	mBase = m.M
	v673 = m.ExcPending
	if v673 != 0 {
		goto L12
	} else {
		goto L185
	}
L183:
	;
	goto L184
L184:
	;
	F_appendStringInfoChar(m, v26+int32(208), int32(92))
	mBase = m.M
	v680 = m.ExcPending
	if v680 != 0 {
		goto L12
	} else {
		goto L186
	}
L185:
	;
	v882 = v631 + int32(2)
	goto L174
L186:
	;
	v882 = v640
	goto L174
L187:
	;
	v688 = *(*int32)(unsafe.Add(mBase, uint32(v682)))
	if v688 < int32(0) {
		v882 = v684
		goto L174
	} else {
		goto L188
	}
L188:
	;
	v691 = v685 - v450
	v693 = *(*int32)(unsafe.Add(mBase, _consts[358]))
	v694 = *(*int32)(unsafe.Add(mBase, uint32(v693)+4))
	v699 = *(*int32)(unsafe.Add(mBase, uint32(v694*int32(28))+uint32(_consts[1003])))
	goto L189
L189:
	;
	if v699 != int32(1) {
		goto L190
	} else {
		goto L191
	}
L190:
	;
	if int32(0) < v691 {
		goto L193
	} else {
		goto L194
	}
L191:
	;
	v765 = v691
	goto L192
L192:
	;
	v781 = v688 - v685
	v784 = v449 + v765
	v786 = *(*int32)(unsafe.Add(mBase, _consts[358]))
	v787 = *(*int32)(unsafe.Add(mBase, uint32(v786)+4))
	v792 = *(*int32)(unsafe.Add(mBase, uint32(v787*int32(28))+uint32(_consts[1003])))
	goto L200
L193:
	;
	v711 = v691
	v712 = v449
	goto L196
L194:
	;
	v742 = v449
	goto L195
L195:
	;
	v765 = v742 - v449
	goto L192
L196:
	;
	v727 = F_pg_mblen_unbounded(m, v712)
	mBase = m.M
	v728 = m.ExcPending
	if v728 != 0 {
		goto L12
	} else {
		goto L198
	}
L197:
	;
	v742 = v729
	goto L195
L198:
	;
	v729 = v727 + v712
	v730 = int32(1)
	if base.Ui32(v730) < base.Ui32(v711) {
		v711 = v711 - v730
		v712 = v729
		goto L196
	} else {
		goto L199
	}
L199:
	;
	goto L197
L200:
	;
	if v792 != int32(1) {
		goto L201
	} else {
		goto L202
	}
L201:
	;
	if int32(0) < v781 {
		goto L204
	} else {
		goto L205
	}
L202:
	;
	v874 = v781
	goto L203
L203:
	;
	F_appendBinaryStringInfo(m, v26+int32(208), v784, v874)
	mBase = m.M
	v876 = m.ExcPending
	if v876 != 0 {
		goto L12
	} else {
		goto L211
	}
L204:
	;
	v804 = v784
	v805 = v781
	goto L207
L205:
	;
	v834 = v784
	goto L206
L206:
	;
	v874 = v834 - v784
	goto L203
L207:
	;
	v820 = F_pg_mblen_unbounded(m, v804)
	mBase = m.M
	v821 = m.ExcPending
	if v821 != 0 {
		goto L12
	} else {
		goto L209
	}
L208:
	;
	v834 = v822
	goto L206
L209:
	;
	v822 = v820 + v804
	v823 = int32(1)
	if base.Ui32(v823) < base.Ui32(v805) {
		v804 = v822
		v805 = v805 - v823
		goto L207
	} else {
		goto L210
	}
L210:
	;
	goto L208
L211:
	;
	v882 = v684
	goto L174
L212:
	;
	goto L135
L213:
	;
	v903 = v82
	goto L215
L214:
	;
	v903 = v84
	goto L215
L215:
	;
	if v469 == int32(1) {
		goto L217
	} else {
		goto L218
	}
L216:
	;
	F_appendBinaryStringInfo(m, v26+int32(208), v903, v931)
	mBase = m.M
	v933 = m.ExcPending
	if v933 != 0 {
		goto L12
	} else {
		goto L227
	}
L217:
	;
	v906 = int32(4)
	v908 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v82))))
	if v908&int32(254) == int32(2) {
		goto L220
	} else {
		goto L221
	}
L218:
	;
	goto L219
L219:
	;
	v921 = int32(1)
	if v471 != 0 {
		v931 = int32(base.Ui32(v469)>>(uint(v921)%32)) - v921
		goto L216
	} else {
		goto L226
	}
L220:
	;
	v917 = v906
	goto L222
L221:
	;
	v917 = base.B2i32(v908 == int32(18)) << (uint(v906) % 32)
	goto L222
L222:
	;
	if v908 == int32(1) {
		goto L223
	} else {
		goto L224
	}
L223:
	;
	v920 = v906
	goto L225
L224:
	;
	v920 = v917
	goto L225
L225:
	;
	v931 = v920
	goto L216
L226:
	;
	v925 = *(*int32)(unsafe.Add(mBase, uint32(l2)))
	v931 = int32(base.Ui32(v925)>>(uint(int32(2))%32)) - int32(4)
	goto L216
L227:
	;
	goto L115
L228:
	;
	if v966 != int32(1) {
		goto L229
	} else {
		goto L230
	}
L229:
	;
	if int32(0) < v958 {
		goto L232
	} else {
		goto L233
	}
L230:
	;
	v1048 = v958
	goto L231
L231:
	;
	v1049 = v1048 + v449
	v1050 = *(*int32)(unsafe.Add(mBase, uint32(v26)+132))
	if int32(0) < l6 {
		v1083 = v1049
		v1088 = v1050
		goto L78
	} else {
		goto L239
	}
L232:
	;
	v978 = v958
	v979 = v449
	goto L235
L233:
	;
	v1009 = v449
	goto L234
L234:
	;
	v1048 = v1009 - v449
	goto L231
L235:
	;
	v994 = F_pg_mblen_unbounded(m, v979)
	mBase = m.M
	v995 = m.ExcPending
	if v995 != 0 {
		goto L12
	} else {
		goto L237
	}
L236:
	;
	v1009 = v996
	goto L234
L237:
	;
	v996 = v994 + v979
	v997 = int32(1)
	if base.Ui32(v997) < base.Ui32(v978) {
		v978 = v978 - v997
		v979 = v996
		goto L235
	} else {
		goto L238
	}
L238:
	;
	goto L236
L239:
	;
	v1056 = v1049
	v1057 = v1050
	v1061 = v1050
	goto L99
L240:
	;
	goto L84
L241:
	;
	v1106 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0))))
	if v1106 == int32(1) {
		goto L245
	} else {
		goto L246
	}
L242:
	;
	goto L243
L243:
	;
	v1139 = *(*int32)(unsafe.Add(mBase, uint32(v26)+208))
	v1140 = *(*int32)(unsafe.Add(mBase, uint32(v26)+212))
	v1142 = v1140 + int32(4)
	v1143 = F_palloc(m, v1142)
	mBase = m.M
	v1144 = m.ExcPending
	if v1144 != 0 {
		goto L12
	} else {
		goto L259
	}
L244:
	;
	F_appendBinaryStringInfo(m, v26+int32(208), v1083, v1133+l0-v1083)
	mBase = m.M
	v1137 = m.ExcPending
	if v1137 != 0 {
		goto L12
	} else {
		goto L258
	}
L245:
	;
	v1109 = int32(6)
	v1111 = int32(18)
	v1113 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v72))))
	if v1113 == v1111 {
		goto L248
	} else {
		goto L249
	}
L246:
	;
	goto L247
L247:
	;
	v1125 = int32(1)
	if v1106&v1125 != 0 {
		v1133 = int32(base.Ui32(v1106) >> (uint(v1125) % 32))
		goto L244
	} else {
		goto L257
	}
L248:
	;
	v1116 = v1111
	goto L250
L249:
	;
	v1116 = int32(2)
	goto L250
L250:
	;
	if v1113&int32(254) == int32(2) {
		goto L251
	} else {
		goto L252
	}
L251:
	;
	v1121 = v1109
	goto L253
L252:
	;
	v1121 = v1116
	goto L253
L253:
	;
	if v1113 == int32(1) {
		goto L254
	} else {
		goto L255
	}
L254:
	;
	v1124 = v1109
	goto L256
L255:
	;
	v1124 = v1121
	goto L256
L256:
	;
	v1133 = v1124
	goto L244
L257:
	;
	v1129 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v1133 = int32(base.Ui32(v1129) >> (uint(int32(2)) % 32))
	goto L244
L258:
	;
	goto L243
L259:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v1143))) = v1142 << (uint(int32(2)) % 32)
	if v1140 != 0 {
		goto L261
	} else {
		goto L262
	}
L260:
	;
	v1152 = *(*int32)(unsafe.Add(mBase, uint32(v26)+208))
	F_pfree(m, v1152)
	mBase = m.M
	v1154 = m.ExcPending
	if v1154 != 0 {
		goto L12
	} else {
		goto L264
	}
L261:
	;
	v1150 = F__emscripten_memcpy_bulkmem(m, v1143+int32(4), v1139, v1140)
	mBase = m.M
	goto L263
L262:
	;
	goto L263
L263:
	;
	goto L260
L264:
	;
	F_pfree(m, v69)
	mBase = m.M
	v1156 = m.ExcPending
	if v1156 != 0 {
		goto L12
	} else {
		goto L265
	}
L265:
	;
	m.G0 = v26 + int32(224)
	return v1143
}
func F_split_text(m *base.Module, l0 int32, l1 int32) int32 {
	mBase := m.M
	_ = mBase
	var v3 int32
	_ = v3
	var v14 int32
	_ = v14
	var v16 int32
	_ = v16
	var v18 int32
	_ = v18
	var v19 int32
	_ = v19
	var v20 int32
	_ = v20
	var v21 int32
	_ = v21
	var v24 int32
	_ = v24
	var v25 int32
	_ = v25
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
	var v35 int32
	_ = v35
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
	var v43 int32
	_ = v43
	var v45 int32
	_ = v45
	var v54 int32
	_ = v54
	var v57 int32
	_ = v57
	var v58 int32
	_ = v58
	var v64 int32
	_ = v64
	var v70 int32
	_ = v70
	var v71 int32
	_ = v71
	var v74 int32
	_ = v74
	var v76 int32
	_ = v76
	var v85 int32
	_ = v85
	var v88 int32
	_ = v88
	var v89 int32
	_ = v89
	var v95 int32
	_ = v95
	var v101 int32
	_ = v101
	var v104 int32
	_ = v104
	var v108 int32
	_ = v108
	var v109 int32
	_ = v109
	var v112 int32
	_ = v112
	var v113 int32
	_ = v113
	var v118 int32
	_ = v118
	var v124 int32
	_ = v124
	var v128 int32
	_ = v128
	var v129 int32
	_ = v129
	var v130 int32
	_ = v130
	var v133 int32
	_ = v133
	var v136 int32
	_ = v136
	var v139 int32
	_ = v139
	var v151 int32
	_ = v151
	var v153 int32
	_ = v153
	var v156 int32
	_ = v156
	var v157 int32
	_ = v157
	var v160 int32
	_ = v160
	var v163 int32
	_ = v163
	var v165 int32
	_ = v165
	var v167 int32
	_ = v167
	var v170 int32
	_ = v170
	var v175 int32
	_ = v175
	var v178 int32
	_ = v178
	var v179 int32
	_ = v179
	var v183 int32
	_ = v183
	var v187 int32
	_ = v187
	var v190 int32
	_ = v190
	var v191 int32
	_ = v191
	var v192 int32
	_ = v192
	var v193 int32
	_ = v193
	var v195 int32
	_ = v195
	var v196 int32
	_ = v196
	var v197 int32
	_ = v197
	var v203 int32
	_ = v203
	var v206 int32
	_ = v206
	var v207 int32
	_ = v207
	var v211 int32
	_ = v211
	var v212 int32
	_ = v212
	var v215 int32
	_ = v215
	var v221 int32
	_ = v221
	var v222 int32
	_ = v222
	var v225 int32
	_ = v225
	var v226 int32
	_ = v226
	var v227 int32
	_ = v227
	var v230 int32
	_ = v230
	var v231 int32
	_ = v231
	var v233 int32
	_ = v233
	var v237 int32
	_ = v237
	var v252 int32
	_ = v252
	var v260 int32
	_ = v260
	var v267 int32
	_ = v267
	var v270 int32
	_ = v270
	var v274 int32
	_ = v274
	var v275 int32
	_ = v275
	var v278 int32
	_ = v278
	var v279 int32
	_ = v279
	var v280 int32
	_ = v280
	var v286 int32
	_ = v286
	var v289 int32
	_ = v289
	var v294 int32
	_ = v294
	var v295 int32
	_ = v295
	var v297 int32
	_ = v297
	var v299 int32
	_ = v299
	var v301 int32
	_ = v301
	var v302 int32
	_ = v302
	var v303 int32
	_ = v303
	var v309 int32
	_ = v309
	var v312 int32
	_ = v312
	var v313 int32
	_ = v313
	var v317 int32
	_ = v317
	var v318 int32
	_ = v318
	var v321 int32
	_ = v321
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
	var v336 int32
	_ = v336
	var v338 int32
	_ = v338
	var v341 int32
	_ = v341
	var v344 int32
	_ = v344
	var v345 int32
	_ = v345
	var v346 int32
	_ = v346
	v3 = int32(0)
	v14 = m.G0
	v16 = v14 - int32(1088)
	m.G0 = v16
	v18 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+24)))
	if v18 != 0 {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	m.G0 = v16 + int32(1088)
	return v18 ^ int32(1)
L2:
	;
	v19 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	v20 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v21 = F_pg_detoast_datum_packed(m, v20)
	mBase = m.M
	v24 = m.ExcPending
	if v24 != 0 {
		goto L3
	} else {
		goto L4
	}
L3:
	;
	return int32(0)
L4:
	;
	v25 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+32)))
	if v25 == int32(0) {
		goto L5
	} else {
		goto L6
	}
L5:
	;
	v28 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
	v29 = F_pg_detoast_datum_packed(m, v28)
	mBase = m.M
	v30 = m.ExcPending
	if v30 != 0 {
		goto L3
	} else {
		goto L8
	}
L6:
	;
	v31 = v3
	goto L7
L7:
	;
	v32 = int32(*(*int16)(unsafe.Add(mBase, uint32(l0)+18)))
	if v32 < int32(3) {
		v39 = v3
		goto L9
	} else {
		goto L10
	}
L8:
	;
	v31 = v29
	goto L7
L9:
	;
	if v31 != 0 {
		goto L14
	} else {
		goto L15
	}
L10:
	;
	v35 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+40)))
	if v35 != 0 {
		v39 = v3
		goto L9
	} else {
		goto L11
	}
L11:
	;
	v36 = *(*int32)(unsafe.Add(mBase, uint32(l0)+36))
	v37 = F_pg_detoast_datum_packed(m, v36)
	mBase = m.M
	v38 = m.ExcPending
	if v38 != 0 {
		goto L3
	} else {
		goto L12
	}
L12:
	;
	v39 = v37
	goto L9
L13:
	;
	v341 = *(*int32)(unsafe.Add(mBase, uint32(l1)))
	v344 = *(*int32)(unsafe.Add(mBase, _consts[0]))
	v345 = F_accumArrayResult(m, v341, v21, v112, int32(25), v344)
	mBase = m.M
	v346 = m.ExcPending
	if v346 != 0 {
		goto L3
	} else {
		goto L132
	}
L14:
	;
	v40 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v21))))
	if v40 == int32(1) {
		goto L18
	} else {
		goto L19
	}
L15:
	;
	goto L16
L16:
	;
	v233 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v21))))
	if v233 == int32(1) {
		goto L99
	} else {
		goto L100
	}
L17:
	;
	v71 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v31))))
	if v71 == int32(1) {
		goto L29
	} else {
		goto L30
	}
L18:
	;
	v43 = int32(4)
	v45 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v21)+1)))
	if v45&int32(254) == int32(2) {
		goto L21
	} else {
		goto L22
	}
L19:
	;
	goto L20
L20:
	;
	v58 = int32(1)
	if v40&v58 != 0 {
		v70 = int32(base.Ui32(v40)>>(uint(v58)%32)) - v58
		goto L17
	} else {
		goto L27
	}
L21:
	;
	v54 = v43
	goto L23
L22:
	;
	v54 = base.B2i32(v45 == int32(18)) << (uint(v43) % 32)
	goto L23
L23:
	;
	if v45 == int32(1) {
		goto L24
	} else {
		goto L25
	}
L24:
	;
	v57 = v43
	goto L26
L25:
	;
	v57 = v54
	goto L26
L26:
	;
	v70 = v57
	goto L17
L27:
	;
	v64 = *(*int32)(unsafe.Add(mBase, uint32(v21)))
	v70 = int32(base.Ui32(v64)>>(uint(int32(2))%32)) - int32(4)
	goto L17
L28:
	;
	if v70 <= int32(0) {
		goto L1
	} else {
		goto L39
	}
L29:
	;
	v74 = int32(4)
	v76 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v31)+1)))
	if v76&int32(254) == int32(2) {
		goto L32
	} else {
		goto L33
	}
L30:
	;
	goto L31
L31:
	;
	v89 = int32(1)
	if v71&v89 != 0 {
		v101 = int32(base.Ui32(v71)>>(uint(v89)%32)) - v89
		goto L28
	} else {
		goto L38
	}
L32:
	;
	v85 = v74
	goto L34
L33:
	;
	v85 = base.B2i32(v76 == int32(18)) << (uint(v74) % 32)
	goto L34
L34:
	;
	if v76 == int32(1) {
		goto L35
	} else {
		goto L36
	}
L35:
	;
	v88 = v74
	goto L37
L36:
	;
	v88 = v85
	goto L37
L37:
	;
	v101 = v88
	goto L28
L38:
	;
	v95 = *(*int32)(unsafe.Add(mBase, uint32(v31)))
	v101 = int32(base.Ui32(v95)>>(uint(int32(2))%32)) - int32(4)
	goto L28
L39:
	;
	v104 = int32(0)
	if v101 <= v104 {
		goto L40
	} else {
		goto L41
	}
L40:
	;
	if v39 != 0 {
		goto L43
	} else {
		goto L44
	}
L41:
	;
	goto L42
L42:
	;
	F_text_position_setup(m, v21, v31, v19, v16+int32(8))
	mBase = m.M
	v128 = m.ExcPending
	if v128 != 0 {
		goto L3
	} else {
		goto L49
	}
L43:
	;
	v108 = F_DirectFunctionCall2Coll(m, int32(1560), v19, v21, v39)
	mBase = m.M
	v109 = m.ExcPending
	if v109 != 0 {
		goto L3
	} else {
		goto L46
	}
L44:
	;
	v112 = v104
	goto L45
L45:
	;
	v113 = *(*int32)(unsafe.Add(mBase, uint32(l1)+4))
	if v113 == int32(0) {
		goto L13
	} else {
		goto L47
	}
L46:
	;
	v112 = base.B2i32(v108 != int32(0))
	goto L45
L47:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v16)+8)) = v21
	*(*uint8)(unsafe.Add(mBase, uint32(v16)+1084)) = uint8(v112)
	v118 = *(*int32)(unsafe.Add(mBase, uint32(l1)+8))
	F_tuplestore_putvalues(m, v113, v118, v16+int32(8), v16+int32(1084))
	mBase = m.M
	v124 = m.ExcPending
	if v124 != 0 {
		goto L3
	} else {
		goto L48
	}
L48:
	;
	goto L1
L49:
	;
	v129 = int32(1)
	v130 = v21 + v129
	v133 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v21))))
	if v133&v129 != 0 {
		goto L50
	} else {
		goto L51
	}
L50:
	;
	v136 = v130
	goto L52
L51:
	;
	v136 = v21 + int32(4)
	goto L52
L52:
	;
	v139 = v136
	goto L53
L53:
	;
	v151 = *(*int32)(unsafe.Add(mBase, _consts[1]))
	if v151 != 0 {
		goto L55
	} else {
		goto L56
	}
L54:
	;
	goto L1
L55:
	;
	F_ProcessInterrupts(m)
	mBase = m.M
	v153 = m.ExcPending
	if v153 != 0 {
		goto L3
	} else {
		goto L58
	}
L56:
	;
	goto L57
L57:
	;
	v156 = F_text_position_next(m, v16+int32(8))
	mBase = m.M
	v157 = m.ExcPending
	if v157 != 0 {
		goto L3
	} else {
		goto L60
	}
L58:
	;
	goto L57
L59:
	;
	v193 = v191 - v139
	v195 = v193 + int32(4)
	v196 = F_palloc(m, v195)
	mBase = m.M
	v197 = m.ExcPending
	if v197 != 0 {
		goto L3
	} else {
		goto L78
	}
L60:
	;
	if v156 == int32(0) {
		goto L61
	} else {
		goto L62
	}
L61:
	;
	v160 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v21))))
	if v160 == int32(1) {
		goto L65
	} else {
		goto L66
	}
L62:
	;
	goto L63
L63:
	;
	v190 = *(*int32)(unsafe.Add(mBase, uint32(v16)+1060))
	v191 = v190
	v192 = v190
	goto L59
L64:
	;
	v191 = v187 + v21
	v192 = int32(0)
	goto L59
L65:
	;
	v163 = int32(6)
	v165 = int32(18)
	v167 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v130))))
	if v167 == v165 {
		goto L68
	} else {
		goto L69
	}
L66:
	;
	goto L67
L67:
	;
	v179 = int32(1)
	if v160&v179 != 0 {
		v187 = int32(base.Ui32(v160) >> (uint(v179) % 32))
		goto L64
	} else {
		goto L77
	}
L68:
	;
	v170 = v165
	goto L70
L69:
	;
	v170 = int32(2)
	goto L70
L70:
	;
	if v167&int32(254) == int32(2) {
		goto L71
	} else {
		goto L72
	}
L71:
	;
	v175 = v163
	goto L73
L72:
	;
	v175 = v170
	goto L73
L73:
	;
	if v167 == int32(1) {
		goto L74
	} else {
		goto L75
	}
L74:
	;
	v178 = v163
	goto L76
L75:
	;
	v178 = v175
	goto L76
L76:
	;
	v187 = v178
	goto L64
L77:
	;
	v183 = *(*int32)(unsafe.Add(mBase, uint32(v21)))
	v187 = int32(base.Ui32(v183) >> (uint(int32(2)) % 32))
	goto L64
L78:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v196))) = v195 << (uint(int32(2)) % 32)
	if v193 != 0 {
		goto L80
	} else {
		goto L81
	}
L79:
	;
	if v39 != 0 {
		goto L83
	} else {
		goto L84
	}
L80:
	;
	v203 = F__emscripten_memcpy_bulkmem(m, v196+int32(4), v139, v193)
	mBase = m.M
	goto L82
L81:
	;
	goto L82
L82:
	;
	goto L79
L83:
	;
	v206 = F_DirectFunctionCall2Coll(m, int32(1560), v19, v196, v39)
	mBase = m.M
	v207 = m.ExcPending
	if v207 != 0 {
		goto L3
	} else {
		goto L86
	}
L84:
	;
	v211 = int32(0)
	goto L85
L85:
	;
	v212 = *(*int32)(unsafe.Add(mBase, uint32(l1)+4))
	if v212 != 0 {
		goto L88
	} else {
		goto L89
	}
L86:
	;
	v211 = base.B2i32(v206 != int32(0))
	goto L85
L87:
	;
	F_pfree(m, v196)
	mBase = m.M
	v230 = m.ExcPending
	if v230 != 0 {
		goto L3
	} else {
		goto L93
	}
L88:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v16)+1084)) = v196
	*(*uint8)(unsafe.Add(mBase, uint32(v16)+1083)) = uint8(v211)
	v215 = *(*int32)(unsafe.Add(mBase, uint32(l1)+8))
	F_tuplestore_putvalues(m, v212, v215, v16+int32(1084), v16+int32(1083))
	mBase = m.M
	v221 = m.ExcPending
	if v221 != 0 {
		goto L3
	} else {
		goto L91
	}
L89:
	;
	goto L90
L90:
	;
	v222 = *(*int32)(unsafe.Add(mBase, uint32(l1)))
	v225 = *(*int32)(unsafe.Add(mBase, _consts[0]))
	v226 = F_accumArrayResult(m, v222, v196, v211, int32(25), v225)
	mBase = m.M
	v227 = m.ExcPending
	if v227 != 0 {
		goto L3
	} else {
		goto L92
	}
L91:
	;
	goto L87
L92:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l1))) = v226
	goto L87
L93:
	;
	v231 = *(*int32)(unsafe.Add(mBase, uint32(v16)+1064))
	if v156 != 0 {
		v139 = v192 + v231
		goto L53
	} else {
		goto L94
	}
L94:
	;
	goto L54
L95:
	;
	v286 = v278
	v289 = v279
	goto L108
L96:
	;
	if v267 <= int32(0) {
		goto L1
	} else {
		goto L104
	}
L97:
	;
	v260 = *(*int32)(unsafe.Add(mBase, uint32(v21)))
	v267 = int32(base.Ui32(v260)>>(uint(int32(2))%32)) - int32(4)
	goto L96
L98:
	;
	v278 = v21 + int32(1)
	v279 = int32(4)
	v280 = v21 + int32(5)
	goto L95
L99:
	;
	v237 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v21)+1)))
	if base.Ui32((v237-int32(1))&int32(255)) < base.Ui32(int32(3)) {
		goto L98
	} else {
		goto L102
	}
L100:
	;
	goto L101
L101:
	;
	if v233&int32(1) == int32(0) {
		goto L97
	} else {
		goto L103
	}
L102:
	;
	v267 = base.B2i32(v237 == int32(18)) << (uint(int32(4)) % 32)
	goto L96
L103:
	;
	v252 = int32(1)
	v267 = int32(base.Ui32(v233)>>(uint(v252)%32)) - v252
	goto L96
L104:
	;
	v270 = int32(1)
	if v233&v270 != 0 {
		goto L105
	} else {
		goto L106
	}
L105:
	;
	v274 = v270
	goto L107
L106:
	;
	v274 = int32(4)
	goto L107
L107:
	;
	v275 = v21 + v274
	v278 = v275
	v279 = v267
	v280 = v275 + v267
	goto L95
L108:
	;
	v294 = F_pg_mblen_range(m, v286, v280)
	mBase = m.M
	v295 = m.ExcPending
	if v295 != 0 {
		goto L3
	} else {
		goto L110
	}
L109:
	;
	goto L1
L110:
	;
	v297 = *(*int32)(unsafe.Add(mBase, _consts[1]))
	if v297 != 0 {
		goto L111
	} else {
		goto L112
	}
L111:
	;
	F_ProcessInterrupts(m)
	mBase = m.M
	v299 = m.ExcPending
	if v299 != 0 {
		goto L3
	} else {
		goto L114
	}
L112:
	;
	goto L113
L113:
	;
	v301 = v294 + int32(4)
	v302 = F_palloc(m, v301)
	mBase = m.M
	v303 = m.ExcPending
	if v303 != 0 {
		goto L3
	} else {
		goto L115
	}
L114:
	;
	goto L113
L115:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v302))) = v301 << (uint(int32(2)) % 32)
	if v294 != 0 {
		goto L117
	} else {
		goto L118
	}
L116:
	;
	if v39 != 0 {
		goto L120
	} else {
		goto L121
	}
L117:
	;
	v309 = F__emscripten_memcpy_bulkmem(m, v302+int32(4), v286, v294)
	mBase = m.M
	goto L119
L118:
	;
	goto L119
L119:
	;
	goto L116
L120:
	;
	v312 = F_DirectFunctionCall2Coll(m, int32(1560), v19, v302, v39)
	mBase = m.M
	v313 = m.ExcPending
	if v313 != 0 {
		goto L3
	} else {
		goto L123
	}
L121:
	;
	v317 = int32(0)
	goto L122
L122:
	;
	v318 = *(*int32)(unsafe.Add(mBase, uint32(l1)+4))
	if v318 != 0 {
		goto L125
	} else {
		goto L126
	}
L123:
	;
	v317 = base.B2i32(v312 != int32(0))
	goto L122
L124:
	;
	F_pfree(m, v302)
	mBase = m.M
	v336 = m.ExcPending
	if v336 != 0 {
		goto L3
	} else {
		goto L130
	}
L125:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v16)+8)) = v302
	*(*uint8)(unsafe.Add(mBase, uint32(v16)+1084)) = uint8(v317)
	v321 = *(*int32)(unsafe.Add(mBase, uint32(l1)+8))
	F_tuplestore_putvalues(m, v318, v321, v16+int32(8), v16+int32(1084))
	mBase = m.M
	v327 = m.ExcPending
	if v327 != 0 {
		goto L3
	} else {
		goto L128
	}
L126:
	;
	goto L127
L127:
	;
	v328 = *(*int32)(unsafe.Add(mBase, uint32(l1)))
	v331 = *(*int32)(unsafe.Add(mBase, _consts[0]))
	v332 = F_accumArrayResult(m, v328, v302, v317, int32(25), v331)
	mBase = m.M
	v333 = m.ExcPending
	if v333 != 0 {
		goto L3
	} else {
		goto L129
	}
L128:
	;
	goto L124
L129:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l1))) = v332
	goto L124
L130:
	;
	v338 = v289 - v294
	if int32(0) < v338 {
		v286 = v294 + v286
		v289 = v338
		goto L108
	} else {
		goto L131
	}
L131:
	;
	goto L109
L132:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l1))) = v345
	goto L1
}
func F_text_char(m *base.Module, l0 int32) int32 {
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
	var v18 int32
	_ = v18
	var v20 int32
	_ = v20
	var v21 int32
	_ = v21
	var v24 int32
	_ = v24
	var v35 int32
	_ = v35
	var v39 int32
	_ = v39
	var v44 int32
	_ = v44
	var v48 int32
	_ = v48
	var v51 int32
	_ = v51
	var v56 int32
	_ = v56
	var v61 int32
	_ = v61
	var v81 int32
	_ = v81
	var v95 int32
	_ = v95
	var v99 int32
	_ = v99
	var v108 int32
	_ = v108
	var v115 int32
	_ = v115
	v9 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v10 = F_pg_detoast_datum_packed(m, v9)
	mBase = m.M
	v13 = m.ExcPending
	if v13 != 0 {
		return int32(0)
	} else {
		v14 = int32(1)
		v15 = v10 + v14
		v18 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v10))))
		v20 = v18 & v14
		if v20 != 0 {
			v21 = v15
		} else {
			v21 = v10 + int32(4)
		}
		if v18 == int32(1) {
			v24 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v15))))
			if base.Ui32((v24-int32(1))&int32(255)) < base.Ui32(int32(3)) {
				v48 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v21))))
				if v48 != int32(92) {
					if v18 != int32(1) {
						v95 = int32(1)
						if v20 != 0 {
							v108 = int32(base.Ui32(v18)>>(uint(v95)%32)) - v95
						} else {
							v99 = *(*int32)(unsafe.Add(mBase, uint32(v10)))
							v108 = int32(base.Ui32(v99)>>(uint(int32(2))%32)) - int32(4)
						}
						if v108 != 0 {
							v115 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v21))))
							return base.I32_extend8_s(v115)
						} else {
							return int32(0)
						}
					} else {
						v81 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v15))))
						if base.Ui32((v81-int32(1))&int32(255)) < base.Ui32(int32(3)) {
							v115 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v21))))
							return base.I32_extend8_s(v115)
						} else {
							v108 = base.B2i32(v81 == int32(18)) << (uint(int32(4)) % 32)
							if v108 != 0 {
								v115 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v21))))
								return base.I32_extend8_s(v115)
							} else {
								return int32(0)
							}
						}
					}
				} else {
					v51 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v21)+1)))
					if v51&int32(248) != int32(48) {
						if v18 != int32(1) {
							v95 = int32(1)
							if v20 != 0 {
								v108 = int32(base.Ui32(v18)>>(uint(v95)%32)) - v95
							} else {
								v99 = *(*int32)(unsafe.Add(mBase, uint32(v10)))
								v108 = int32(base.Ui32(v99)>>(uint(int32(2))%32)) - int32(4)
							}
							if v108 != 0 {
								v115 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v21))))
								return base.I32_extend8_s(v115)
							} else {
								return int32(0)
							}
						} else {
							v81 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v15))))
							if base.Ui32((v81-int32(1))&int32(255)) < base.Ui32(int32(3)) {
								v115 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v21))))
								return base.I32_extend8_s(v115)
							} else {
								v108 = base.B2i32(v81 == int32(18)) << (uint(int32(4)) % 32)
								if v108 != 0 {
									v115 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v21))))
									return base.I32_extend8_s(v115)
								} else {
									return int32(0)
								}
							}
						}
					} else {
						v56 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v21)+2)))
						if v56&int32(248) != int32(48) {
							if v18 != int32(1) {
								v95 = int32(1)
								if v20 != 0 {
									v108 = int32(base.Ui32(v18)>>(uint(v95)%32)) - v95
								} else {
									v99 = *(*int32)(unsafe.Add(mBase, uint32(v10)))
									v108 = int32(base.Ui32(v99)>>(uint(int32(2))%32)) - int32(4)
								}
								if v108 != 0 {
									v115 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v21))))
									return base.I32_extend8_s(v115)
								} else {
									return int32(0)
								}
							} else {
								v81 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v15))))
								if base.Ui32((v81-int32(1))&int32(255)) < base.Ui32(int32(3)) {
									v115 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v21))))
									return base.I32_extend8_s(v115)
								} else {
									v108 = base.B2i32(v81 == int32(18)) << (uint(int32(4)) % 32)
									if v108 != 0 {
										v115 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v21))))
										return base.I32_extend8_s(v115)
									} else {
										return int32(0)
									}
								}
							}
						} else {
							v61 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v21)+3)))
							if v61&int32(248) != int32(48) {
								if v18 != int32(1) {
									v95 = int32(1)
									if v20 != 0 {
										v108 = int32(base.Ui32(v18)>>(uint(v95)%32)) - v95
									} else {
										v99 = *(*int32)(unsafe.Add(mBase, uint32(v10)))
										v108 = int32(base.Ui32(v99)>>(uint(int32(2))%32)) - int32(4)
									}
									if v108 != 0 {
										v115 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v21))))
										return base.I32_extend8_s(v115)
									} else {
										return int32(0)
									}
								} else {
									v81 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v15))))
									if base.Ui32((v81-int32(1))&int32(255)) < base.Ui32(int32(3)) {
										v115 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v21))))
										return base.I32_extend8_s(v115)
									} else {
										v108 = base.B2i32(v81 == int32(18)) << (uint(int32(4)) % 32)
										if v108 != 0 {
											v115 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v21))))
											return base.I32_extend8_s(v115)
										} else {
											return int32(0)
										}
									}
								}
							} else {
								return base.I32_extend8_s(v51<<(uint(int32(6))%32) + v56<<(uint(int32(3))%32) + v61 + int32(80))
							}
						}
					}
				}
			} else {
				if v24 == int32(18) {
					v115 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v21))))
					return base.I32_extend8_s(v115)
				} else {
					return int32(0)
				}
			}
		} else {
			if v20 != 0 {
				v35 = int32(1)
				v44 = int32(base.Ui32(v18)>>(uint(v35)%32)) - v35
			} else {
				v39 = *(*int32)(unsafe.Add(mBase, uint32(v10)))
				v44 = int32(base.Ui32(v39)>>(uint(int32(2))%32)) - int32(4)
			}
			if v44 != int32(4) {
				v95 = int32(1)
				if v20 != 0 {
					v108 = int32(base.Ui32(v18)>>(uint(v95)%32)) - v95
				} else {
					v99 = *(*int32)(unsafe.Add(mBase, uint32(v10)))
					v108 = int32(base.Ui32(v99)>>(uint(int32(2))%32)) - int32(4)
				}
				if v108 != 0 {
					v115 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v21))))
					return base.I32_extend8_s(v115)
				} else {
					return int32(0)
				}
			} else {
				v48 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v21))))
				if v48 != int32(92) {
					if v18 != int32(1) {
						v95 = int32(1)
						if v20 != 0 {
							v108 = int32(base.Ui32(v18)>>(uint(v95)%32)) - v95
						} else {
							v99 = *(*int32)(unsafe.Add(mBase, uint32(v10)))
							v108 = int32(base.Ui32(v99)>>(uint(int32(2))%32)) - int32(4)
						}
						if v108 != 0 {
							v115 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v21))))
							return base.I32_extend8_s(v115)
						} else {
							return int32(0)
						}
					} else {
						v81 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v15))))
						if base.Ui32((v81-int32(1))&int32(255)) < base.Ui32(int32(3)) {
							v115 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v21))))
							return base.I32_extend8_s(v115)
						} else {
							v108 = base.B2i32(v81 == int32(18)) << (uint(int32(4)) % 32)
							if v108 != 0 {
								v115 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v21))))
								return base.I32_extend8_s(v115)
							} else {
								return int32(0)
							}
						}
					}
				} else {
					v51 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v21)+1)))
					if v51&int32(248) != int32(48) {
						if v18 != int32(1) {
							v95 = int32(1)
							if v20 != 0 {
								v108 = int32(base.Ui32(v18)>>(uint(v95)%32)) - v95
							} else {
								v99 = *(*int32)(unsafe.Add(mBase, uint32(v10)))
								v108 = int32(base.Ui32(v99)>>(uint(int32(2))%32)) - int32(4)
							}
							if v108 != 0 {
								v115 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v21))))
								return base.I32_extend8_s(v115)
							} else {
								return int32(0)
							}
						} else {
							v81 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v15))))
							if base.Ui32((v81-int32(1))&int32(255)) < base.Ui32(int32(3)) {
								v115 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v21))))
								return base.I32_extend8_s(v115)
							} else {
								v108 = base.B2i32(v81 == int32(18)) << (uint(int32(4)) % 32)
								if v108 != 0 {
									v115 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v21))))
									return base.I32_extend8_s(v115)
								} else {
									return int32(0)
								}
							}
						}
					} else {
						v56 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v21)+2)))
						if v56&int32(248) != int32(48) {
							if v18 != int32(1) {
								v95 = int32(1)
								if v20 != 0 {
									v108 = int32(base.Ui32(v18)>>(uint(v95)%32)) - v95
								} else {
									v99 = *(*int32)(unsafe.Add(mBase, uint32(v10)))
									v108 = int32(base.Ui32(v99)>>(uint(int32(2))%32)) - int32(4)
								}
								if v108 != 0 {
									v115 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v21))))
									return base.I32_extend8_s(v115)
								} else {
									return int32(0)
								}
							} else {
								v81 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v15))))
								if base.Ui32((v81-int32(1))&int32(255)) < base.Ui32(int32(3)) {
									v115 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v21))))
									return base.I32_extend8_s(v115)
								} else {
									v108 = base.B2i32(v81 == int32(18)) << (uint(int32(4)) % 32)
									if v108 != 0 {
										v115 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v21))))
										return base.I32_extend8_s(v115)
									} else {
										return int32(0)
									}
								}
							}
						} else {
							v61 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v21)+3)))
							if v61&int32(248) != int32(48) {
								if v18 != int32(1) {
									v95 = int32(1)
									if v20 != 0 {
										v108 = int32(base.Ui32(v18)>>(uint(v95)%32)) - v95
									} else {
										v99 = *(*int32)(unsafe.Add(mBase, uint32(v10)))
										v108 = int32(base.Ui32(v99)>>(uint(int32(2))%32)) - int32(4)
									}
									if v108 != 0 {
										v115 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v21))))
										return base.I32_extend8_s(v115)
									} else {
										return int32(0)
									}
								} else {
									v81 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v15))))
									if base.Ui32((v81-int32(1))&int32(255)) < base.Ui32(int32(3)) {
										v115 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v21))))
										return base.I32_extend8_s(v115)
									} else {
										v108 = base.B2i32(v81 == int32(18)) << (uint(int32(4)) % 32)
										if v108 != 0 {
											v115 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v21))))
											return base.I32_extend8_s(v115)
										} else {
											return int32(0)
										}
									}
								}
							} else {
								return base.I32_extend8_s(v51<<(uint(int32(6))%32) + v56<<(uint(int32(3))%32) + v61 + int32(80))
							}
						}
					}
				}
			}
		}
	}
}
func F_text_format_nv(m *base.Module, l0 int32) int32 {
	var v2 int32
	_ = v2
	var v5 int32
	_ = v5
	v2 = F_text_format(m, l0)
	v5 = m.ExcPending
	if v5 != 0 {
		return int32(0)
	} else {
		return v2
	}
}
func F_text_format_parse_digits(m *base.Module, l0 int32, l1 int32, l2 int32) int32 {
	mBase := m.M
	_ = mBase
	var v4 int32
	_ = v4
	var v10 int32
	_ = v10
	var v11 int32
	_ = v11
	var v17 int32
	_ = v17
	var v23 int32
	_ = v23
	var v24 int32
	_ = v24
	var v25 int32
	_ = v25
	var v31 int64
	_ = v31
	var v35 int32
	_ = v35
	var v42 int32
	_ = v42
	var v45 int32
	_ = v45
	var v49 int32
	_ = v49
	var v51 int32
	_ = v51
	var v63 int32
	_ = v63
	var v66 int32
	_ = v66
	var v70 int32
	_ = v70
	var v74 int32
	_ = v74
	var v79 int32
	_ = v79
	var v83 int32
	_ = v83
	var v84 int32
	_ = v84
	var v96 int32
	_ = v96
	var v99 int32
	_ = v99
	var v103 int32
	_ = v103
	var v108 int32
	_ = v108
	v4 = int32(0)
	v10 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v11 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v10))))
	v17 = base.B2i32(base.Ui32((v11-int32(48))&int32(255)) < base.Ui32(int32(10)))
	if v17 == v4 {
		v83 = v4
		v84 = v10
		goto L2
	} else {
		goto L3
	}
L1:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v96 = m.ExcPending
	if v96 != 0 {
		goto L12
	} else {
		goto L18
	}
L2:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0))) = v84
	*(*int32)(unsafe.Add(mBase, uint32(l2))) = v83
	return v17
L3:
	;
	v23 = v4
	v24 = v10
	v25 = v11
	goto L4
L4:
	;
	v31 = base.I64_extend_i32_s(v23) * int64(10)
	v35 = base.I32_wrap_i64(v31)
	if base.I32_wrap_i64(int64(base.Ui64(v31)>>(uint(int64(32))%64))) != v35>>(uint(int32(31))%32) {
		goto L1
	} else {
		goto L6
	}
L5:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v63 = m.ExcPending
	if v63 != 0 {
		goto L12
	} else {
		goto L13
	}
L6:
	;
	v42 = (v25 - int32(48)) & int32(255)
	v45 = v42 + v35
	if base.B2i32(base.Ui32(v42) < base.Ui32(int32(0))) != base.B2i32(v45 < v35) {
		goto L1
	} else {
		goto L7
	}
L7:
	;
	v49 = v24 + int32(1)
	if base.Ui32(v49) < base.Ui32(l1) {
		goto L8
	} else {
		goto L9
	}
L8:
	;
	v51 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v49))))
	if base.Ui32(int32(9)) < base.Ui32((v51-int32(48))&int32(255)) {
		v83 = v45
		v84 = v49
		goto L2
	} else {
		goto L11
	}
L9:
	;
	goto L10
L10:
	;
	goto L5
L11:
	;
	v23 = v45
	v24 = v49
	v25 = v51
	goto L4
L12:
	;
	return int32(0)
L13:
	;
	F_errcode(m, int32(50856066))
	mBase = m.M
	v66 = m.ExcPending
	if v66 != 0 {
		goto L12
	} else {
		goto L14
	}
L14:
	;
	F_errmsg(m, int32(220578), int32(0))
	mBase = m.M
	v70 = m.ExcPending
	if v70 != 0 {
		goto L12
	} else {
		goto L15
	}
L15:
	;
	F_errhint(m, int32(648285), int32(0))
	mBase = m.M
	v74 = m.ExcPending
	if v74 != 0 {
		goto L12
	} else {
		goto L16
	}
L16:
	;
	F_errfinish(m, int32(494617), int32(6191), int32(122376))
	mBase = m.M
	v79 = m.ExcPending
	if v79 != 0 {
		goto L12
	} else {
		goto L17
	}
L17:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L18:
	;
	F_errcode(m, int32(50331778))
	mBase = m.M
	v99 = m.ExcPending
	if v99 != 0 {
		goto L12
	} else {
		goto L19
	}
L19:
	;
	F_errmsg(m, int32(397885), int32(0))
	mBase = m.M
	v103 = m.ExcPending
	if v103 != 0 {
		goto L12
	} else {
		goto L20
	}
L20:
	;
	F_errfinish(m, int32(494617), int32(6190), int32(122376))
	mBase = m.M
	v108 = m.ExcPending
	if v108 != 0 {
		goto L12
	} else {
		goto L21
	}
L21:
	;
	base.Wasm_trap_unreachable()
	for {
	}
}
func F_text_overlay(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) int32 {
	var v9 int32
	_ = v9
	var v12 int32
	_ = v12
	var v16 int32
	_ = v16
	var v19 int32
	_ = v19
	var v22 int32
	_ = v22
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
	var v32 int32
	_ = v32
	var v35 int32
	_ = v35
	var v39 int32
	_ = v39
	var v44 int32
	_ = v44
	var v48 int32
	_ = v48
	var v51 int32
	_ = v51
	var v55 int32
	_ = v55
	var v60 int32
	_ = v60
	if int32(0) < l2 {
		v9 = l2 + l3
		if base.B2i32(l3 < int32(0)) != base.B2i32(v9 < l2) {
			F_errstart_cold(m, int32(21), int32(0))
			v48 = m.ExcPending
			if v48 != 0 {
				return int32(0)
			} else {
				F_errcode(m, int32(50331778))
				v51 = m.ExcPending
				if v51 != 0 {
					return int32(0)
				} else {
					F_errmsg(m, int32(398054), int32(0))
					v55 = m.ExcPending
					if v55 != 0 {
						return int32(0)
					} else {
						F_errfinish(m, int32(494617), int32(1186), int32(26319))
						v60 = m.ExcPending
						if v60 != 0 {
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
			v12 = int32(1)
			v16 = F_text_substring(m, l0, v12, l2-v12, int32(0))
			v19 = m.ExcPending
			if v19 != 0 {
				return int32(0)
			} else {
				v22 = F_text_substring(m, l0, v9, int32(-1), int32(1))
				v23 = m.ExcPending
				if v23 != 0 {
					return int32(0)
				} else {
					v24 = F_text_catenate(m, v16, l1)
					v25 = m.ExcPending
					if v25 != 0 {
						return int32(0)
					} else {
						v26 = F_text_catenate(m, v24, v22)
						v27 = m.ExcPending
						if v27 != 0 {
							return int32(0)
						} else {
							return v26
						}
					}
				}
			}
		}
	} else {
		F_errstart_cold(m, int32(21), int32(0))
		v32 = m.ExcPending
		if v32 != 0 {
			return int32(0)
		} else {
			F_errcode(m, int32(17039490))
			v35 = m.ExcPending
			if v35 != 0 {
				return int32(0)
			} else {
				F_errmsg(m, int32(434838), int32(0))
				v39 = m.ExcPending
				if v39 != 0 {
					return int32(0)
				} else {
					F_errfinish(m, int32(494617), int32(1182), int32(26319))
					v44 = m.ExcPending
					if v44 != 0 {
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
func F_text_pattern_le(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v6 int32
	_ = v6
	var v7 int32
	_ = v7
	var v10 int32
	_ = v10
	var v12 int32
	_ = v12
	var v13 int32
	_ = v13
	var v14 int32
	_ = v14
	var v15 int32
	_ = v15
	var v20 int32
	_ = v20
	var v23 int32
	_ = v23
	var v25 int32
	_ = v25
	var v34 int32
	_ = v34
	var v37 int32
	_ = v37
	var v38 int32
	_ = v38
	var v44 int32
	_ = v44
	var v50 int32
	_ = v50
	var v51 int32
	_ = v51
	var v54 int32
	_ = v54
	var v56 int32
	_ = v56
	var v65 int32
	_ = v65
	var v68 int32
	_ = v68
	var v69 int32
	_ = v69
	var v75 int32
	_ = v75
	var v81 int32
	_ = v81
	var v82 int32
	_ = v82
	var v86 int32
	_ = v86
	var v88 int32
	_ = v88
	var v92 int32
	_ = v92
	var v94 int32
	_ = v94
	var v95 int32
	_ = v95
	var v96 int32
	_ = v96
	var v99 int32
	_ = v99
	var v100 int32
	_ = v100
	var v103 int32
	_ = v103
	var v104 int32
	_ = v104
	var v107 int32
	_ = v107
	v6 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v7 = F_pg_detoast_datum_packed(m, v6)
	mBase = m.M
	v10 = m.ExcPending
	if v10 != 0 {
		return int32(0)
	} else {
		v12 = l0 + int32(28)
		v13 = *(*int32)(unsafe.Add(mBase, uint32(v12)))
		v14 = F_pg_detoast_datum_packed(m, v13)
		mBase = m.M
		v15 = m.ExcPending
		if v15 != 0 {
			return int32(0)
		} else {
			v20 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v7))))
			if v20 == int32(1) {
				v23 = int32(4)
				v25 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v7)+1)))
				if v25&int32(254) == int32(2) {
					v34 = v23
				} else {
					v34 = base.B2i32(v25 == int32(18)) << (uint(v23) % 32)
				}
				if v25 == int32(1) {
					v37 = v23
				} else {
					v37 = v34
				}
				v50 = v37
			} else {
				v38 = int32(1)
				if v20&v38 != 0 {
					v50 = int32(base.Ui32(v20)>>(uint(v38)%32)) - v38
				} else {
					v44 = *(*int32)(unsafe.Add(mBase, uint32(v7)))
					v50 = int32(base.Ui32(v44)>>(uint(int32(2))%32)) - int32(4)
				}
			}
			v51 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v14))))
			if v51 == int32(1) {
				v54 = int32(4)
				v56 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v14)+1)))
				if v56&int32(254) == int32(2) {
					v65 = v54
				} else {
					v65 = base.B2i32(v56 == int32(18)) << (uint(v54) % 32)
				}
				if v56 == int32(1) {
					v68 = v54
				} else {
					v68 = v65
				}
				v81 = v68
			} else {
				v69 = int32(1)
				if v51&v69 != 0 {
					v81 = int32(base.Ui32(v51)>>(uint(v69)%32)) - v69
				} else {
					v75 = *(*int32)(unsafe.Add(mBase, uint32(v14)))
					v81 = int32(base.Ui32(v75)>>(uint(int32(2))%32)) - int32(4)
				}
			}
			v82 = int32(1)
			if v20&v82 != 0 {
				v86 = v82
			} else {
				v86 = int32(4)
			}
			v88 = int32(1)
			if v51&v88 != 0 {
				v92 = v88
			} else {
				v92 = int32(4)
			}
			v94 = base.B2i32(v50 < v81)
			if v50 < v81 {
				v95 = v50
			} else {
				v95 = v81
			}
			v96 = F_memcmp(m, v7+v86, v14+v92, v95)
			mBase = m.M
			if v96 != 0 {
				v99 = v96
			} else {
				if v50 < v81 {
					v99 = int32(-1)
				} else {
					v99 = base.B2i32(v81 < v50)
				}
			}
			v100 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
			if v100 != v7 {
				F_pfree(m, v7)
				mBase = m.M
				v103 = m.ExcPending
				if v103 != 0 {
					return int32(0)
				} else {
					v104 = *(*int32)(unsafe.Add(mBase, uint32(v12)))
					if v104 != v14 {
						F_pfree(m, v14)
						mBase = m.M
						v107 = m.ExcPending
						if v107 != 0 {
							return int32(0)
						} else {
							return base.B2i32(v99 <= int32(0))
						}
					} else {
						return base.B2i32(v99 <= int32(0))
					}
				}
			} else {
				v104 = *(*int32)(unsafe.Add(mBase, uint32(v12)))
				if v104 != v14 {
					F_pfree(m, v14)
					mBase = m.M
					v107 = m.ExcPending
					if v107 != 0 {
						return int32(0)
					} else {
						return base.B2i32(v99 <= int32(0))
					}
				} else {
					return base.B2i32(v99 <= int32(0))
				}
			}
		}
	}
}
func F_text_right(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v7 int32
	_ = v7
	var v8 int32
	_ = v8
	var v11 int32
	_ = v11
	var v14 int32
	_ = v14
	var v15 int32
	_ = v15
	var v16 int32
	_ = v16
	var v18 int32
	_ = v18
	var v21 int32
	_ = v21
	var v23 int32
	_ = v23
	var v32 int32
	_ = v32
	var v35 int32
	_ = v35
	var v36 int32
	_ = v36
	var v40 int32
	_ = v40
	var v46 int32
	_ = v46
	var v47 int32
	_ = v47
	var v48 int32
	_ = v48
	var v49 int32
	_ = v49
	var v52 int32
	_ = v52
	var v53 int32
	_ = v53
	var v54 int32
	_ = v54
	var v56 int32
	_ = v56
	var v57 int32
	_ = v57
	var v58 int32
	_ = v58
	var v60 int32
	_ = v60
	var v61 int32
	_ = v61
	var v62 int32
	_ = v62
	var v69 int32
	_ = v69
	v7 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v8 = F_pg_detoast_datum_packed(m, v7)
	mBase = m.M
	v11 = m.ExcPending
	if v11 != 0 {
		return int32(0)
	} else {
		v14 = int32(1)
		v15 = v8 + v14
		v16 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v8))))
		v18 = v16 & v14
		if v16 == v14 {
			v21 = int32(4)
			v23 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v15))))
			if v23&int32(254) == int32(2) {
				v32 = v21
			} else {
				v32 = base.B2i32(v23 == int32(18)) << (uint(v21) % 32)
			}
			if v23 == int32(1) {
				v35 = v21
			} else {
				v35 = v32
			}
			v46 = v35
		} else {
			v36 = int32(1)
			if v18 != 0 {
				v46 = int32(base.Ui32(v16)>>(uint(v36)%32)) - v36
			} else {
				v40 = *(*int32)(unsafe.Add(mBase, uint32(v8)))
				v46 = int32(base.Ui32(v40)>>(uint(int32(2))%32)) - int32(4)
			}
		}
		if v18 != 0 {
			v47 = v15
		} else {
			v47 = v8 + int32(4)
		}
		v48 = int32(0)
		v49 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
		if v48 <= v49 {
			v52 = F_pg_mbstrlen_with_len(m, v47, v46)
			mBase = m.M
			v53 = m.ExcPending
			if v53 != 0 {
				return int32(0)
			} else {
				v54 = v52
				v56 = F_pg_mbcharcliplen(m, v47, v46, v54-v49)
				mBase = m.M
				v57 = m.ExcPending
				if v57 != 0 {
					return int32(0)
				} else {
					v58 = v46 - v56
					v60 = v58 + int32(4)
					v61 = F_palloc(m, v60)
					mBase = m.M
					v62 = m.ExcPending
					if v62 != 0 {
						return int32(0)
					} else {
						*(*int32)(unsafe.Add(mBase, uint32(v61))) = v60 << (uint(int32(2)) % 32)
						if v58 != 0 {
							v69 = F__emscripten_memcpy_bulkmem(m, v61+int32(4), v56+v47, v58)
							mBase = m.M
						} else {
						}
						return v61
					}
				}
			}
		} else {
			v54 = v48
			v56 = F_pg_mbcharcliplen(m, v47, v46, v54-v49)
			mBase = m.M
			v57 = m.ExcPending
			if v57 != 0 {
				return int32(0)
			} else {
				v58 = v46 - v56
				v60 = v58 + int32(4)
				v61 = F_palloc(m, v60)
				mBase = m.M
				v62 = m.ExcPending
				if v62 != 0 {
					return int32(0)
				} else {
					*(*int32)(unsafe.Add(mBase, uint32(v61))) = v60 << (uint(int32(2)) % 32)
					if v58 != 0 {
						v69 = F__emscripten_memcpy_bulkmem(m, v61+int32(4), v56+v47, v58)
						mBase = m.M
					} else {
					}
					return v61
				}
			}
		}
	}
}
