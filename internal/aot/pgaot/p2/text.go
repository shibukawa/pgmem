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
	var v34 int32
	_ = v34
	var v37 int32
	_ = v37
	var v44 int32
	_ = v44
	var v45 int32
	_ = v45
	var v51 int32
	_ = v51
	var v57 int32
	_ = v57
	var v63 int32
	_ = v63
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
	var v74 int32
	_ = v74
	var v77 int32
	_ = v77
	var v78 int32
	_ = v78
	var v79 int32
	_ = v79
	var v80 int32
	_ = v80
	var v81 int32
	_ = v81
	var v83 int32
	_ = v83
	var v84 int32
	_ = v84
	var v86 int32
	_ = v86
	var v87 int32
	_ = v87
	var v90 int32
	_ = v90
	var v101 int32
	_ = v101
	var v102 int32
	_ = v102
	var v106 int32
	_ = v106
	var v112 int32
	_ = v112
	var v117 int32
	_ = v117
	var v125 int32
	_ = v125
	var v134 int32
	_ = v134
	var v142 int32
	_ = v142
	var v143 int32
	_ = v143
	var v156 int32
	_ = v156
	var v158 int32
	_ = v158
	var v161 int32
	_ = v161
	var v163 int32
	_ = v163
	var v164 int32
	_ = v164
	var v165 int32
	_ = v165
	var v166 int32
	_ = v166
	var v168 int32
	_ = v168
	var v173 int32
	_ = v173
	var v175 int32
	_ = v175
	var v176 int32
	_ = v176
	var v182 int32
	_ = v182
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
	var v283 int32
	_ = v283
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
	var v305 int32
	_ = v305
	var v312 int32
	_ = v312
	var v315 int32
	_ = v315
	var v316 int32
	_ = v316
	var v332 int32
	_ = v332
	var v335 int32
	_ = v335
	var v337 int32
	_ = v337
	var v340 int32
	_ = v340
	var v341 int32
	_ = v341
	var v345 int32
	_ = v345
	var v347 int32
	_ = v347
	var v351 int32
	_ = v351
	var v354 int32
	_ = v354
	var v358 int32
	_ = v358
	var v363 int32
	_ = v363
	var v365 int32
	_ = v365
	var v367 int32
	_ = v367
	var v372 int32
	_ = v372
	var v373 int32
	_ = v373
	var v374 int32
	_ = v374
	var v380 int32
	_ = v380
	var v381 int32
	_ = v381
	var v386 int32
	_ = v386
	var v392 int32
	_ = v392
	var v396 int32
	_ = v396
	var v412 int32
	_ = v412
	var v413 int32
	_ = v413
	var v414 int32
	_ = v414
	var v415 int32
	_ = v415
	var v427 int32
	_ = v427
	var v444 int32
	_ = v444
	var v446 int32
	_ = v446
	var v448 int32
	_ = v448
	var v451 int32
	_ = v451
	var v470 int32
	_ = v470
	var v472 int32
	_ = v472
	var v473 int32
	_ = v473
	var v476 int32
	_ = v476
	var v487 int32
	_ = v487
	var v490 int32
	_ = v490
	var v496 int32
	_ = v496
	var v502 int32
	_ = v502
	var v507 int32
	_ = v507
	var v517 int32
	_ = v517
	var v532 int32
	_ = v532
	var v533 int32
	_ = v533
	var v546 int32
	_ = v546
	var v548 int32
	_ = v548
	var v551 int32
	_ = v551
	var v553 int32
	_ = v553
	var v554 int32
	_ = v554
	var v555 int32
	_ = v555
	var v556 int32
	_ = v556
	var v558 int32
	_ = v558
	var v563 int32
	_ = v563
	var v565 int32
	_ = v565
	var v566 int32
	_ = v566
	var v572 int32
	_ = v572
	var v581 int32
	_ = v581
	var v583 int32
	_ = v583
	var v587 int32
	_ = v587
	var v588 int32
	_ = v588
	var v591 int32
	_ = v591
	var v595 int32
	_ = v595
	var v596 int32
	_ = v596
	var v598 int32
	_ = v598
	var v601 int32
	_ = v601
	var v603 int32
	_ = v603
	var v608 int32
	_ = v608
	var v610 int32
	_ = v610
	var v615 int32
	_ = v615
	var v617 int32
	_ = v617
	var v620 int32
	_ = v620
	var v622 int32
	_ = v622
	var v625 int32
	_ = v625
	var v637 int32
	_ = v637
	var v638 int32
	_ = v638
	var v644 int32
	_ = v644
	var v647 int32
	_ = v647
	var v653 int32
	_ = v653
	var v654 int32
	_ = v654
	var v665 int32
	_ = v665
	var v680 int32
	_ = v680
	var v687 int32
	_ = v687
	var v688 int32
	_ = v688
	var v689 int32
	_ = v689
	var v691 int32
	_ = v691
	var v692 int32
	_ = v692
	var v695 int32
	_ = v695
	var v698 int32
	_ = v698
	var v700 int32
	_ = v700
	var v701 int32
	_ = v701
	var v706 int32
	_ = v706
	var v714 int32
	_ = v714
	var v718 int32
	_ = v718
	var v734 int32
	_ = v734
	var v735 int32
	_ = v735
	var v736 int32
	_ = v736
	var v737 int32
	_ = v737
	var v744 int32
	_ = v744
	var v772 int32
	_ = v772
	var v788 int32
	_ = v788
	var v791 int32
	_ = v791
	var v793 int32
	_ = v793
	var v794 int32
	_ = v794
	var v799 int32
	_ = v799
	var v807 int32
	_ = v807
	var v811 int32
	_ = v811
	var v827 int32
	_ = v827
	var v828 int32
	_ = v828
	var v829 int32
	_ = v829
	var v830 int32
	_ = v830
	var v841 int32
	_ = v841
	var v881 int32
	_ = v881
	var v883 int32
	_ = v883
	var v893 int32
	_ = v893
	var v910 int32
	_ = v910
	var v916 int32
	_ = v916
	var v919 int32
	_ = v919
	var v926 int32
	_ = v926
	var v927 int32
	_ = v927
	var v931 int32
	_ = v931
	var v937 int32
	_ = v937
	var v939 int32
	_ = v939
	var v963 int32
	_ = v963
	var v964 int32
	_ = v964
	var v966 int32
	_ = v966
	var v967 int32
	_ = v967
	var v972 int32
	_ = v972
	var v980 int32
	_ = v980
	var v984 int32
	_ = v984
	var v1000 int32
	_ = v1000
	var v1001 int32
	_ = v1001
	var v1002 int32
	_ = v1002
	var v1003 int32
	_ = v1003
	var v1010 int32
	_ = v1010
	var v1054 int32
	_ = v1054
	var v1055 int32
	_ = v1055
	var v1056 int32
	_ = v1056
	var v1060 int32
	_ = v1060
	var v1063 int32
	_ = v1063
	var v1066 int32
	_ = v1066
	var v1082 int32
	_ = v1082
	var v1084 int32
	_ = v1084
	var v1090 int32
	_ = v1090
	var v1093 int32
	_ = v1093
	var v1112 int32
	_ = v1112
	var v1116 int32
	_ = v1116
	var v1118 int32
	_ = v1118
	var v1121 int32
	_ = v1121
	var v1128 int32
	_ = v1128
	var v1129 int32
	_ = v1129
	var v1133 int32
	_ = v1133
	var v1137 int32
	_ = v1137
	var v1141 int32
	_ = v1141
	var v1143 int32
	_ = v1143
	var v1144 int32
	_ = v1144
	var v1146 int32
	_ = v1146
	var v1147 int32
	_ = v1147
	var v1148 int32
	_ = v1148
	var v1155 int32
	_ = v1155
	var v1157 int32
	_ = v1157
	var v1159 int32
	_ = v1159
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
	v63 = m.ExcPending
	if v63 != 0 {
		goto L12
	} else {
		goto L13
	}
L2:
	;
	v34 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+1)))
	if v34 == int32(18) {
		goto L5
	} else {
		goto L6
	}
L3:
	;
	goto L4
L4:
	;
	v45 = int32(1)
	if v28&v45 != 0 {
		v57 = int32(base.Ui32(v28)>>(uint(v45)%32)) - v45
		goto L1
	} else {
		goto L11
	}
L5:
	;
	v37 = int32(16)
	goto L7
L6:
	;
	v37 = int32(0)
	goto L7
L7:
	;
	if base.Ui32((v34-int32(1))&int32(255)) < base.Ui32(int32(3)) {
		goto L8
	} else {
		goto L9
	}
L8:
	;
	v44 = int32(4)
	goto L10
L9:
	;
	v44 = v37
	goto L10
L10:
	;
	v57 = v44
	goto L1
L11:
	;
	v51 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v57 = int32(base.Ui32(v51)>>(uint(int32(2))%32)) - int32(4)
	goto L1
L12:
	;
	return int32(0)
L13:
	;
	v68 = F_palloc(m, v57<<(uint(int32(2))%32)+int32(4))
	mBase = m.M
	v69 = m.ExcPending
	if v69 != 0 {
		goto L12
	} else {
		goto L14
	}
L14:
	;
	v70 = int32(1)
	v71 = l0 + v70
	v73 = l0 + int32(4)
	v74 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0))))
	if v74&v70 != 0 {
		goto L15
	} else {
		goto L16
	}
L15:
	;
	v77 = v71
	goto L17
L16:
	;
	v77 = v73
	goto L17
L17:
	;
	v78 = F_pg_mb2wchar_with_len(m, v77, v68, v57)
	mBase = m.M
	v79 = m.ExcPending
	if v79 != 0 {
		goto L12
	} else {
		goto L18
	}
L18:
	;
	v80 = int32(1)
	v81 = l2 + v80
	v83 = l2 + int32(4)
	v84 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l2))))
	v86 = v84 & v80
	if v86 != 0 {
		goto L19
	} else {
		goto L20
	}
L19:
	;
	v87 = v81
	goto L21
L20:
	;
	v87 = v83
	goto L21
L21:
	;
	if v84 == int32(1) {
		goto L26
	} else {
		goto L27
	}
L22:
	;
	v293 = base.B2i32(base.Ui32(v283) < base.Ui32(int32(2)))
	if base.Ui32(v283) < base.Ui32(int32(2)) {
		goto L72
	} else {
		goto L73
	}
L23:
	;
	v125 = v87
	v134 = v8
	goto L37
L24:
	;
	v117 = v87 + int32(4)
	goto L23
L25:
	;
	if v112 != 0 {
		goto L34
	} else {
		goto L35
	}
L26:
	;
	v90 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v81))))
	if base.Ui32((v90-int32(1))&int32(255)) < base.Ui32(int32(3)) {
		goto L24
	} else {
		goto L29
	}
L27:
	;
	goto L28
L28:
	;
	v102 = int32(1)
	if v86 != 0 {
		v112 = int32(base.Ui32(v84)>>(uint(v102)%32)) - v102
		goto L25
	} else {
		goto L33
	}
L29:
	;
	if v90 == int32(18) {
		goto L30
	} else {
		goto L31
	}
L30:
	;
	v101 = int32(16)
	goto L32
L31:
	;
	v101 = int32(0)
	goto L32
L32:
	;
	v112 = v101
	goto L25
L33:
	;
	v106 = *(*int32)(unsafe.Add(mBase, uint32(l2)))
	v112 = int32(base.Ui32(v106)>>(uint(int32(2))%32)) - int32(4)
	goto L25
L34:
	;
	v117 = v87 + v112
	goto L23
L35:
	;
	goto L36
L36:
	;
	v283 = v8
	goto L22
L37:
	;
	v142 = v117 - v125
	v143 = int32(0)
	if base.B2i32(v125&int32(3) == v143)|base.B2i32(v142 == v143) != 0 {
		v173 = v125
		v175 = v142
		v176 = base.B2i32(v142 != v143)
		goto L42
	} else {
		goto L43
	}
L38:
	;
	v283 = v265
	goto L22
L39:
	;
	if v247 == int32(0) {
		v283 = v134
		goto L22
	} else {
		goto L64
	}
L40:
	;
	v247 = int32(0)
	goto L39
L41:
	;
	v225 = v218
	v227 = v220
	goto L58
L42:
	;
	if v176 == int32(0) {
		goto L40
	} else {
		goto L49
	}
L43:
	;
	v156 = v125
	v158 = v142
	goto L44
L44:
	;
	v161 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v156))))
	if v161 == int32(92) {
		v218 = v156
		v220 = v158
		goto L41
	} else {
		goto L46
	}
L45:
	;
	v173 = v168
	v175 = v164
	v176 = v166
	goto L42
L46:
	;
	v163 = int32(1)
	v164 = v158 - v163
	v165 = int32(0)
	v166 = base.B2i32(v164 != v165)
	v168 = v156 + v163
	if v168&int32(3) == v165 {
		v173 = v168
		v175 = v164
		v176 = v166
		goto L42
	} else {
		goto L47
	}
L47:
	;
	if v164 != 0 {
		v156 = v168
		v158 = v164
		goto L44
	} else {
		goto L48
	}
L48:
	;
	goto L45
L49:
	;
	v182 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v173))))
	if base.B2i32(int32(92) == v182)|base.B2i32(base.Ui32(v175) < base.Ui32(int32(4))) == int32(0) {
		goto L50
	} else {
		goto L51
	}
L50:
	;
	v191 = v173
	v193 = v175
	goto L53
L51:
	;
	v211 = v173
	v213 = v175
	goto L52
L52:
	;
	if v213 == int32(0) {
		goto L40
	} else {
		goto L57
	}
L53:
	;
	v197 = *(*int32)(unsafe.Add(mBase, uint32(v191)))
	v198 = v197 ^ int32(1549556828)
	v201 = int32(-2139062144)
	if (int32(16843008)-v198|v198)&v201 != v201 {
		v218 = v191
		v220 = v193
		goto L41
	} else {
		goto L55
	}
L54:
	;
	v211 = v206
	v213 = v208
	goto L52
L55:
	;
	v205 = int32(4)
	v206 = v191 + v205
	v208 = v193 - v205
	if base.Ui32(int32(3)) < base.Ui32(v208) {
		v191 = v206
		v193 = v208
		goto L53
	} else {
		goto L56
	}
L56:
	;
	goto L54
L57:
	;
	v218 = v211
	v220 = v213
	goto L41
L58:
	;
	v230 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v225))))
	if int32(92) == v230 {
		goto L60
	} else {
		goto L61
	}
L59:
	;
	goto L40
L60:
	;
	v247 = v225
	goto L39
L61:
	;
	goto L62
L62:
	;
	v232 = int32(1)
	v235 = v227 - v232
	if v235 != 0 {
		v225 = v225 + v232
		v227 = v235
		goto L58
	} else {
		goto L63
	}
L63:
	;
	goto L59
L64:
	;
	v251 = v247 + int32(1)
	if base.Ui32(v251) < base.Ui32(v117) {
		goto L65
	} else {
		goto L66
	}
L65:
	;
	v253 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v251))))
	if base.Ui32((v253-int32(49))&int32(255)) < base.Ui32(int32(9)) {
		goto L68
	} else {
		goto L69
	}
L66:
	;
	v264 = v251
	v265 = v134
	goto L67
L67:
	;
	if base.Ui32(v264) < base.Ui32(v117) {
		v125 = v264
		v134 = v265
		goto L37
	} else {
		goto L71
	}
L68:
	;
	v283 = int32(2)
	goto L22
L69:
	;
	goto L70
L70:
	;
	v264 = v247 + int32(2)
	v265 = int32(1)
	goto L67
L71:
	;
	goto L38
L72:
	;
	v294 = l3 | int32(16)
	goto L74
L73:
	;
	v294 = l3
	goto L74
L74:
	;
	v295 = F_RE_compile_and_cache(m, l1, v294, l4)
	mBase = m.M
	v296 = m.ExcPending
	if v296 != 0 {
		goto L12
	} else {
		goto L75
	}
L75:
	;
	v297 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0))))
	if v297&int32(1) != 0 {
		goto L76
	} else {
		goto L77
	}
L76:
	;
	v300 = v71
	goto L78
L77:
	;
	v300 = v73
	goto L78
L78:
	;
	if base.Ui32(v78) < base.Ui32(l5) {
		v1090 = v300
		v1093 = int32(0)
		goto L79
	} else {
		goto L80
	}
L79:
	;
	if base.Ui32(v1093) < base.Ui32(v78) {
		goto L242
	} else {
		goto L243
	}
L80:
	;
	if base.Ui32(v283) < base.Ui32(int32(2)) {
		goto L81
	} else {
		goto L82
	}
L81:
	;
	v305 = int32(1)
	goto L83
L82:
	;
	v305 = int32(10)
	goto L83
L83:
	;
	v312 = int32(0)
	v315 = v300
	v316 = l5
	v332 = v8
	goto L84
L84:
	;
	v335 = *(*int32)(unsafe.Add(mBase, _c_F_replace_text_regexp[0]))
	if v335 != 0 {
		goto L86
	} else {
		goto L87
	}
L85:
	;
	v1090 = v1063
	v1093 = v1060
	goto L79
L86:
	;
	F_ProcessInterrupts(m)
	mBase = m.M
	v337 = m.ExcPending
	if v337 != 0 {
		goto L12
	} else {
		goto L89
	}
L87:
	;
	goto L88
L88:
	;
	v340 = F_pg_regexec(m, v295, v68, v78, v316, v305, v26+int32(128))
	mBase = m.M
	v341 = m.ExcPending
	if v341 != 0 {
		goto L12
	} else {
		goto L90
	}
L89:
	;
	goto L88
L90:
	;
	if v340 != 0 {
		goto L91
	} else {
		goto L92
	}
L91:
	;
	if v340 == int32(1) {
		v1090 = v315
		v1093 = v312
		goto L79
	} else {
		goto L94
	}
L92:
	;
	goto L93
L93:
	;
	v365 = v332 + int32(1)
	v367 = int32(0)
	if base.B2i32(v365 == l6)|base.B2i32(l6 <= v367) == v367 {
		goto L101
	} else {
		goto L102
	}
L94:
	;
	v345 = v26 + int32(16)
	F_pg_regerror(m, v340, v345)
	mBase = m.M
	v347 = m.ExcPending
	if v347 != 0 {
		goto L12
	} else {
		goto L95
	}
L95:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v351 = m.ExcPending
	if v351 != 0 {
		goto L12
	} else {
		goto L96
	}
L96:
	;
	F_errcode(m, int32(302252162))
	mBase = m.M
	v354 = m.ExcPending
	if v354 != 0 {
		goto L12
	} else {
		goto L97
	}
L97:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v26))) = v345
	F_errmsg(m, int32(_a_F_replace_text_regexp_0), v26)
	mBase = m.M
	v358 = m.ExcPending
	if v358 != 0 {
		goto L12
	} else {
		goto L98
	}
L98:
	;
	F_errfinish(m, int32(_a_F_replace_text_regexp_1), int32(_a_F_replace_text_regexp_2), int32(_a_F_replace_text_regexp_3))
	mBase = m.M
	v363 = m.ExcPending
	if v363 != 0 {
		goto L12
	} else {
		goto L99
	}
L99:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L100:
	;
	v1082 = *(*int32)(unsafe.Add(mBase, uint32(v26)+128))
	v1084 = v1066 + base.B2i32(v1082 == v1066)
	if base.Ui32(v1084) <= base.Ui32(v78) {
		v312 = v1060
		v315 = v1063
		v316 = v1084
		v332 = v365
		goto L84
	} else {
		goto L241
	}
L101:
	;
	v372 = *(*int32)(unsafe.Add(mBase, uint32(v26)+132))
	v1060 = v312
	v1063 = v315
	v1066 = v372
	goto L100
L102:
	;
	goto L103
L103:
	;
	v373 = *(*int32)(unsafe.Add(mBase, uint32(v26)+128))
	v374 = v373 - v312
	if int32(0) < v374 {
		goto L104
	} else {
		goto L105
	}
L104:
	;
	v380 = *(*int32)(unsafe.Add(mBase, _c_F_replace_text_regexp[1]))
	v381 = *(*int32)(unsafe.Add(mBase, uint32(v380)+4))
	v386 = *(*int32)(unsafe.Add(mBase, uint32(v381*int32(28))+uint32(_c_F_replace_text_regexp[2])))
	goto L107
L105:
	;
	v448 = v312
	v451 = v315
	goto L106
L106:
	;
	v470 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l2))))
	v472 = v470 & int32(1)
	if v283 != 0 {
		goto L117
	} else {
		goto L118
	}
L107:
	;
	if v386 != int32(1) {
		goto L108
	} else {
		goto L109
	}
L108:
	;
	v392 = v315
	v396 = v374
	goto L111
L109:
	;
	v427 = v374
	goto L110
L110:
	;
	F_appendBinaryStringInfo(m, v26+int32(208), v315, v427)
	mBase = m.M
	v444 = m.ExcPending
	if v444 != 0 {
		goto L12
	} else {
		goto L115
	}
L111:
	;
	v412 = F_pg_mblen_unbounded(m, v392)
	mBase = m.M
	v413 = m.ExcPending
	if v413 != 0 {
		goto L12
	} else {
		goto L113
	}
L112:
	;
	v427 = v414 - v315
	goto L110
L113:
	;
	v414 = v412 + v392
	v415 = int32(1)
	if base.Ui32(v415) < base.Ui32(v396) {
		v392 = v414
		v396 = v396 - v415
		goto L111
	} else {
		goto L114
	}
L114:
	;
	goto L112
L115:
	;
	v446 = *(*int32)(unsafe.Add(mBase, uint32(v26)+128))
	v448 = v446
	v451 = v315 + v427
	goto L106
L116:
	;
	v963 = *(*int32)(unsafe.Add(mBase, uint32(v26)+132))
	v964 = v963 - v448
	v966 = *(*int32)(unsafe.Add(mBase, _c_F_replace_text_regexp[1]))
	v967 = *(*int32)(unsafe.Add(mBase, uint32(v966)+4))
	v972 = *(*int32)(unsafe.Add(mBase, uint32(v967*int32(28))+uint32(_c_F_replace_text_regexp[2])))
	goto L229
L117:
	;
	if v472 != 0 {
		goto L120
	} else {
		goto L121
	}
L118:
	;
	goto L119
L119:
	;
	if v472 != 0 {
		goto L214
	} else {
		goto L215
	}
L120:
	;
	v473 = v81
	goto L122
L121:
	;
	v473 = v83
	goto L122
L122:
	;
	if v470 == int32(1) {
		goto L127
	} else {
		goto L128
	}
L123:
	;
	v517 = v473
	goto L136
L124:
	;
	if v502 == int32(0) {
		goto L116
	} else {
		goto L135
	}
L125:
	;
	v496 = *(*int32)(unsafe.Add(mBase, uint32(l2)))
	v502 = int32(base.Ui32(v496)>>(uint(int32(2))%32)) - int32(4)
	goto L124
L126:
	;
	v507 = v473 + int32(4)
	goto L123
L127:
	;
	v476 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v81))))
	if base.Ui32((v476-int32(1))&int32(255)) < base.Ui32(int32(3)) {
		goto L126
	} else {
		goto L130
	}
L128:
	;
	goto L129
L129:
	;
	if v472 == int32(0) {
		goto L125
	} else {
		goto L134
	}
L130:
	;
	if v476 == int32(18) {
		goto L131
	} else {
		goto L132
	}
L131:
	;
	v487 = int32(16)
	goto L133
L132:
	;
	v487 = int32(0)
	goto L133
L133:
	;
	v502 = v487
	goto L124
L134:
	;
	v490 = int32(1)
	v502 = int32(base.Ui32(v470)>>(uint(v490)%32)) - v490
	goto L124
L135:
	;
	v507 = v502 + v473
	goto L123
L136:
	;
	v532 = v507 - v517
	v533 = int32(0)
	if base.B2i32(v517&int32(3) == v533)|base.B2i32(v532 == v533) != 0 {
		v563 = v517
		v565 = v532
		v566 = base.B2i32(v532 != v533)
		goto L141
	} else {
		goto L142
	}
L137:
	;
	goto L116
L138:
	;
	if v637 != 0 {
		goto L163
	} else {
		goto L164
	}
L139:
	;
	v637 = int32(0)
	goto L138
L140:
	;
	v615 = v608
	v617 = v610
	goto L157
L141:
	;
	if v566 == int32(0) {
		goto L139
	} else {
		goto L148
	}
L142:
	;
	v546 = v517
	v548 = v532
	goto L143
L143:
	;
	v551 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v546))))
	if v551 == int32(92) {
		v608 = v546
		v610 = v548
		goto L140
	} else {
		goto L145
	}
L144:
	;
	v563 = v558
	v565 = v554
	v566 = v556
	goto L141
L145:
	;
	v553 = int32(1)
	v554 = v548 - v553
	v555 = int32(0)
	v556 = base.B2i32(v554 != v555)
	v558 = v546 + v553
	if v558&int32(3) == v555 {
		v563 = v558
		v565 = v554
		v566 = v556
		goto L141
	} else {
		goto L146
	}
L146:
	;
	if v554 != 0 {
		v546 = v558
		v548 = v554
		goto L143
	} else {
		goto L147
	}
L147:
	;
	goto L144
L148:
	;
	v572 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v563))))
	if base.B2i32(int32(92) == v572)|base.B2i32(base.Ui32(v565) < base.Ui32(int32(4))) == int32(0) {
		goto L149
	} else {
		goto L150
	}
L149:
	;
	v581 = v563
	v583 = v565
	goto L152
L150:
	;
	v601 = v563
	v603 = v565
	goto L151
L151:
	;
	if v603 == int32(0) {
		goto L139
	} else {
		goto L156
	}
L152:
	;
	v587 = *(*int32)(unsafe.Add(mBase, uint32(v581)))
	v588 = v587 ^ int32(1549556828)
	v591 = int32(-2139062144)
	if (int32(16843008)-v588|v588)&v591 != v591 {
		v608 = v581
		v610 = v583
		goto L140
	} else {
		goto L154
	}
L153:
	;
	v601 = v596
	v603 = v598
	goto L151
L154:
	;
	v595 = int32(4)
	v596 = v581 + v595
	v598 = v583 - v595
	if base.Ui32(int32(3)) < base.Ui32(v598) {
		v581 = v596
		v583 = v598
		goto L152
	} else {
		goto L155
	}
L155:
	;
	goto L153
L156:
	;
	v608 = v601
	v610 = v603
	goto L140
L157:
	;
	v620 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v615))))
	if int32(92) == v620 {
		goto L159
	} else {
		goto L160
	}
L158:
	;
	goto L139
L159:
	;
	v637 = v615
	goto L138
L160:
	;
	goto L161
L161:
	;
	v622 = int32(1)
	v625 = v617 - v622
	if v625 != 0 {
		v615 = v615 + v622
		v617 = v625
		goto L157
	} else {
		goto L162
	}
L162:
	;
	goto L158
L163:
	;
	v638 = v637
	goto L165
L164:
	;
	v638 = v507
	goto L165
L165:
	;
	if base.Ui32(v517) < base.Ui32(v638) {
		goto L166
	} else {
		goto L167
	}
L166:
	;
	F_appendBinaryStringInfo(m, v26+int32(208), v517, v638-v517)
	mBase = m.M
	v644 = m.ExcPending
	if v644 != 0 {
		goto L12
	} else {
		goto L169
	}
L167:
	;
	goto L168
L168:
	;
	if base.Ui32(v507) <= base.Ui32(v638) {
		goto L116
	} else {
		goto L170
	}
L169:
	;
	goto L168
L170:
	;
	v647 = v638 + int32(1)
	if base.Ui32(v507) <= base.Ui32(v647) {
		goto L171
	} else {
		goto L172
	}
L171:
	;
	F_appendStringInfoChar(m, v26+int32(208), int32(92))
	mBase = m.M
	v653 = m.ExcPending
	if v653 != 0 {
		goto L12
	} else {
		goto L174
	}
L172:
	;
	goto L173
L173:
	;
	v654 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v647))))
	if base.Ui32((v654-int32(49))&int32(255)) <= base.Ui32(int32(8)) {
		goto L177
	} else {
		goto L178
	}
L174:
	;
	goto L116
L175:
	;
	if base.Ui32(v893) < base.Ui32(v507) {
		v517 = v893
		goto L136
	} else {
		goto L213
	}
L176:
	;
	v691 = v638 + int32(2)
	v692 = *(*int32)(unsafe.Add(mBase, uint32(v688)))
	if v692 < int32(0) {
		v893 = v691
		goto L175
	} else {
		goto L188
	}
L177:
	;
	v665 = v26 + int32(128) + v654<<(uint(int32(3))%32)
	v688 = v665 - int32(384)
	v689 = v665 - int32(380)
	goto L176
L178:
	;
	goto L179
L179:
	;
	if v654 == int32(38) {
		goto L180
	} else {
		goto L181
	}
L180:
	;
	v688 = v26 + int32(128)
	v689 = v26 + int32(128) | int32(4)
	goto L176
L181:
	;
	goto L182
L182:
	;
	if v654 == int32(92) {
		goto L183
	} else {
		goto L184
	}
L183:
	;
	F_appendStringInfoChar(m, v26+int32(208), int32(92))
	mBase = m.M
	v680 = m.ExcPending
	if v680 != 0 {
		goto L12
	} else {
		goto L186
	}
L184:
	;
	goto L185
L185:
	;
	F_appendStringInfoChar(m, v26+int32(208), int32(92))
	mBase = m.M
	v687 = m.ExcPending
	if v687 != 0 {
		goto L12
	} else {
		goto L187
	}
L186:
	;
	v893 = v638 + int32(2)
	goto L175
L187:
	;
	v893 = v647
	goto L175
L188:
	;
	v695 = *(*int32)(unsafe.Add(mBase, uint32(v689)))
	if v695 < int32(0) {
		v893 = v691
		goto L175
	} else {
		goto L189
	}
L189:
	;
	v698 = v692 - v448
	v700 = *(*int32)(unsafe.Add(mBase, _c_F_replace_text_regexp[1]))
	v701 = *(*int32)(unsafe.Add(mBase, uint32(v700)+4))
	v706 = *(*int32)(unsafe.Add(mBase, uint32(v701*int32(28))+uint32(_c_F_replace_text_regexp[2])))
	goto L190
L190:
	;
	if v706 != int32(1) {
		goto L191
	} else {
		goto L192
	}
L191:
	;
	if int32(0) < v698 {
		goto L194
	} else {
		goto L195
	}
L192:
	;
	v772 = v698
	goto L193
L193:
	;
	v788 = v695 - v692
	v791 = v451 + v772
	v793 = *(*int32)(unsafe.Add(mBase, _c_F_replace_text_regexp[1]))
	v794 = *(*int32)(unsafe.Add(mBase, uint32(v793)+4))
	v799 = *(*int32)(unsafe.Add(mBase, uint32(v794*int32(28))+uint32(_c_F_replace_text_regexp[2])))
	goto L201
L194:
	;
	v714 = v451
	v718 = v698
	goto L197
L195:
	;
	v744 = v451
	goto L196
L196:
	;
	v772 = v744 - v451
	goto L193
L197:
	;
	v734 = F_pg_mblen_unbounded(m, v714)
	mBase = m.M
	v735 = m.ExcPending
	if v735 != 0 {
		goto L12
	} else {
		goto L199
	}
L198:
	;
	v744 = v736
	goto L196
L199:
	;
	v736 = v734 + v714
	v737 = int32(1)
	if base.Ui32(v737) < base.Ui32(v718) {
		v714 = v736
		v718 = v718 - v737
		goto L197
	} else {
		goto L200
	}
L200:
	;
	goto L198
L201:
	;
	if v799 != int32(1) {
		goto L202
	} else {
		goto L203
	}
L202:
	;
	if int32(0) < v788 {
		goto L205
	} else {
		goto L206
	}
L203:
	;
	v881 = v788
	goto L204
L204:
	;
	F_appendBinaryStringInfo(m, v26+int32(208), v791, v881)
	mBase = m.M
	v883 = m.ExcPending
	if v883 != 0 {
		goto L12
	} else {
		goto L212
	}
L205:
	;
	v807 = v788
	v811 = v791
	goto L208
L206:
	;
	v841 = v791
	goto L207
L207:
	;
	v881 = v841 - v791
	goto L204
L208:
	;
	v827 = F_pg_mblen_unbounded(m, v811)
	mBase = m.M
	v828 = m.ExcPending
	if v828 != 0 {
		goto L12
	} else {
		goto L210
	}
L209:
	;
	v841 = v829
	goto L207
L210:
	;
	v829 = v827 + v811
	v830 = int32(1)
	if base.Ui32(v830) < base.Ui32(v807) {
		v807 = v807 - v830
		v811 = v829
		goto L208
	} else {
		goto L211
	}
L211:
	;
	goto L209
L212:
	;
	v893 = v691
	goto L175
L213:
	;
	goto L137
L214:
	;
	v910 = v81
	goto L216
L215:
	;
	v910 = v83
	goto L216
L216:
	;
	if v470 == int32(1) {
		goto L218
	} else {
		goto L219
	}
L217:
	;
	F_appendBinaryStringInfo(m, v26+int32(208), v910, v937)
	mBase = m.M
	v939 = m.ExcPending
	if v939 != 0 {
		goto L12
	} else {
		goto L228
	}
L218:
	;
	v916 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v81))))
	if v916 == int32(18) {
		goto L221
	} else {
		goto L222
	}
L219:
	;
	goto L220
L220:
	;
	v927 = int32(1)
	if v472 != 0 {
		v937 = int32(base.Ui32(v470)>>(uint(v927)%32)) - v927
		goto L217
	} else {
		goto L227
	}
L221:
	;
	v919 = int32(16)
	goto L223
L222:
	;
	v919 = int32(0)
	goto L223
L223:
	;
	if base.Ui32((v916-int32(1))&int32(255)) < base.Ui32(int32(3)) {
		goto L224
	} else {
		goto L225
	}
L224:
	;
	v926 = int32(4)
	goto L226
L225:
	;
	v926 = v919
	goto L226
L226:
	;
	v937 = v926
	goto L217
L227:
	;
	v931 = *(*int32)(unsafe.Add(mBase, uint32(l2)))
	v937 = int32(base.Ui32(v931)>>(uint(int32(2))%32)) - int32(4)
	goto L217
L228:
	;
	goto L116
L229:
	;
	if v972 != int32(1) {
		goto L230
	} else {
		goto L231
	}
L230:
	;
	if int32(0) < v964 {
		goto L233
	} else {
		goto L234
	}
L231:
	;
	v1054 = v964
	goto L232
L232:
	;
	v1055 = v1054 + v451
	v1056 = *(*int32)(unsafe.Add(mBase, uint32(v26)+132))
	if int32(0) < l6 {
		v1090 = v1055
		v1093 = v1056
		goto L79
	} else {
		goto L240
	}
L233:
	;
	v980 = v451
	v984 = v964
	goto L236
L234:
	;
	v1010 = v451
	goto L235
L235:
	;
	v1054 = v1010 - v451
	goto L232
L236:
	;
	v1000 = F_pg_mblen_unbounded(m, v980)
	mBase = m.M
	v1001 = m.ExcPending
	if v1001 != 0 {
		goto L12
	} else {
		goto L238
	}
L237:
	;
	v1010 = v1002
	goto L235
L238:
	;
	v1002 = v1000 + v980
	v1003 = int32(1)
	if base.Ui32(v1003) < base.Ui32(v984) {
		v980 = v1002
		v984 = v984 - v1003
		goto L236
	} else {
		goto L239
	}
L239:
	;
	goto L237
L240:
	;
	v1060 = v1056
	v1063 = v1055
	v1066 = v1056
	goto L100
L241:
	;
	goto L85
L242:
	;
	v1112 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0))))
	if v1112 == int32(1) {
		goto L246
	} else {
		goto L247
	}
L243:
	;
	goto L244
L244:
	;
	v1143 = *(*int32)(unsafe.Add(mBase, uint32(v26)+208))
	v1144 = *(*int32)(unsafe.Add(mBase, uint32(v26)+212))
	v1146 = v1144 + int32(4)
	v1147 = F_palloc(m, v1146)
	mBase = m.M
	v1148 = m.ExcPending
	if v1148 != 0 {
		goto L12
	} else {
		goto L257
	}
L245:
	;
	F_appendBinaryStringInfo(m, v26+int32(208), v1090, v1137+l0-v1090)
	mBase = m.M
	v1141 = m.ExcPending
	if v1141 != 0 {
		goto L12
	} else {
		goto L256
	}
L246:
	;
	v1116 = int32(18)
	v1118 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v71))))
	if v1118 == v1116 {
		goto L249
	} else {
		goto L250
	}
L247:
	;
	goto L248
L248:
	;
	v1129 = int32(1)
	if v1112&v1129 != 0 {
		v1137 = int32(base.Ui32(v1112) >> (uint(v1129) % 32))
		goto L245
	} else {
		goto L255
	}
L249:
	;
	v1121 = v1116
	goto L251
L250:
	;
	v1121 = int32(2)
	goto L251
L251:
	;
	if base.Ui32((v1118-int32(1))&int32(255)) < base.Ui32(int32(3)) {
		goto L252
	} else {
		goto L253
	}
L252:
	;
	v1128 = int32(6)
	goto L254
L253:
	;
	v1128 = v1121
	goto L254
L254:
	;
	v1137 = v1128
	goto L245
L255:
	;
	v1133 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v1137 = int32(base.Ui32(v1133) >> (uint(int32(2)) % 32))
	goto L245
L256:
	;
	goto L244
L257:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v1147))) = v1146 << (uint(int32(2)) % 32)
	if v1144 != 0 {
		goto L258
	} else {
		goto L259
	}
L258:
	;
	base.MemoryCopy(m, v1147+int32(4), v1143, v1144)
	goto L260
L259:
	;
	goto L260
L260:
	;
	v1155 = *(*int32)(unsafe.Add(mBase, uint32(v26)+208))
	F_pfree(m, v1155)
	mBase = m.M
	v1157 = m.ExcPending
	if v1157 != 0 {
		goto L12
	} else {
		goto L261
	}
L261:
	;
	F_pfree(m, v68)
	mBase = m.M
	v1159 = m.ExcPending
	if v1159 != 0 {
		goto L12
	} else {
		goto L262
	}
L262:
	;
	m.G0 = v26 + int32(224)
	return v1147
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
	var v46 int32
	_ = v46
	var v49 int32
	_ = v49
	var v56 int32
	_ = v56
	var v57 int32
	_ = v57
	var v63 int32
	_ = v63
	var v69 int32
	_ = v69
	var v70 int32
	_ = v70
	var v76 int32
	_ = v76
	var v79 int32
	_ = v79
	var v86 int32
	_ = v86
	var v87 int32
	_ = v87
	var v93 int32
	_ = v93
	var v99 int32
	_ = v99
	var v102 int32
	_ = v102
	var v106 int32
	_ = v106
	var v107 int32
	_ = v107
	var v110 int32
	_ = v110
	var v111 int32
	_ = v111
	var v116 int32
	_ = v116
	var v122 int32
	_ = v122
	var v126 int32
	_ = v126
	var v127 int32
	_ = v127
	var v128 int32
	_ = v128
	var v131 int32
	_ = v131
	var v134 int32
	_ = v134
	var v138 int32
	_ = v138
	var v149 int32
	_ = v149
	var v151 int32
	_ = v151
	var v154 int32
	_ = v154
	var v155 int32
	_ = v155
	var v158 int32
	_ = v158
	var v162 int32
	_ = v162
	var v164 int32
	_ = v164
	var v167 int32
	_ = v167
	var v174 int32
	_ = v174
	var v175 int32
	_ = v175
	var v179 int32
	_ = v179
	var v183 int32
	_ = v183
	var v186 int32
	_ = v186
	var v187 int32
	_ = v187
	var v188 int32
	_ = v188
	var v189 int32
	_ = v189
	var v191 int32
	_ = v191
	var v192 int32
	_ = v192
	var v193 int32
	_ = v193
	var v201 int32
	_ = v201
	var v202 int32
	_ = v202
	var v206 int32
	_ = v206
	var v207 int32
	_ = v207
	var v210 int32
	_ = v210
	var v216 int32
	_ = v216
	var v217 int32
	_ = v217
	var v220 int32
	_ = v220
	var v221 int32
	_ = v221
	var v222 int32
	_ = v222
	var v225 int32
	_ = v225
	var v226 int32
	_ = v226
	var v228 int32
	_ = v228
	var v231 int32
	_ = v231
	var v242 int32
	_ = v242
	var v247 int32
	_ = v247
	var v256 int32
	_ = v256
	var v262 int32
	_ = v262
	var v265 int32
	_ = v265
	var v269 int32
	_ = v269
	var v270 int32
	_ = v270
	var v272 int32
	_ = v272
	var v274 int32
	_ = v274
	var v275 int32
	_ = v275
	var v276 int32
	_ = v276
	var v281 int32
	_ = v281
	var v289 int32
	_ = v289
	var v290 int32
	_ = v290
	var v292 int32
	_ = v292
	var v294 int32
	_ = v294
	var v296 int32
	_ = v296
	var v297 int32
	_ = v297
	var v298 int32
	_ = v298
	var v306 int32
	_ = v306
	var v307 int32
	_ = v307
	var v311 int32
	_ = v311
	var v312 int32
	_ = v312
	var v315 int32
	_ = v315
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
	var v330 int32
	_ = v330
	var v332 int32
	_ = v332
	var v335 int32
	_ = v335
	var v338 int32
	_ = v338
	var v339 int32
	_ = v339
	var v340 int32
	_ = v340
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
	v335 = *(*int32)(unsafe.Add(mBase, uint32(l1)))
	v338 = *(*int32)(unsafe.Add(mBase, _c_F_split_text[0]))
	v339 = F_accumArrayResult(m, v335, v21, v110, int32(25), v338)
	mBase = m.M
	v340 = m.ExcPending
	if v340 != 0 {
		goto L3
	} else {
		goto L130
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
	v228 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v21))))
	if v228 == int32(1) {
		goto L95
	} else {
		goto L96
	}
L17:
	;
	v70 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v31))))
	if v70 == int32(1) {
		goto L29
	} else {
		goto L30
	}
L18:
	;
	v46 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v21)+1)))
	if v46 == int32(18) {
		goto L21
	} else {
		goto L22
	}
L19:
	;
	goto L20
L20:
	;
	v57 = int32(1)
	if v40&v57 != 0 {
		v69 = int32(base.Ui32(v40)>>(uint(v57)%32)) - v57
		goto L17
	} else {
		goto L27
	}
L21:
	;
	v49 = int32(16)
	goto L23
L22:
	;
	v49 = int32(0)
	goto L23
L23:
	;
	if base.Ui32((v46-int32(1))&int32(255)) < base.Ui32(int32(3)) {
		goto L24
	} else {
		goto L25
	}
L24:
	;
	v56 = int32(4)
	goto L26
L25:
	;
	v56 = v49
	goto L26
L26:
	;
	v69 = v56
	goto L17
L27:
	;
	v63 = *(*int32)(unsafe.Add(mBase, uint32(v21)))
	v69 = int32(base.Ui32(v63)>>(uint(int32(2))%32)) - int32(4)
	goto L17
L28:
	;
	if v69 <= int32(0) {
		goto L1
	} else {
		goto L39
	}
L29:
	;
	v76 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v31)+1)))
	if v76 == int32(18) {
		goto L32
	} else {
		goto L33
	}
L30:
	;
	goto L31
L31:
	;
	v87 = int32(1)
	if v70&v87 != 0 {
		v99 = int32(base.Ui32(v70)>>(uint(v87)%32)) - v87
		goto L28
	} else {
		goto L38
	}
L32:
	;
	v79 = int32(16)
	goto L34
L33:
	;
	v79 = int32(0)
	goto L34
L34:
	;
	if base.Ui32((v76-int32(1))&int32(255)) < base.Ui32(int32(3)) {
		goto L35
	} else {
		goto L36
	}
L35:
	;
	v86 = int32(4)
	goto L37
L36:
	;
	v86 = v79
	goto L37
L37:
	;
	v99 = v86
	goto L28
L38:
	;
	v93 = *(*int32)(unsafe.Add(mBase, uint32(v31)))
	v99 = int32(base.Ui32(v93)>>(uint(int32(2))%32)) - int32(4)
	goto L28
L39:
	;
	v102 = int32(0)
	if v99 <= v102 {
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
	v126 = m.ExcPending
	if v126 != 0 {
		goto L3
	} else {
		goto L49
	}
L43:
	;
	v106 = F_DirectFunctionCall2Coll(m, int32(1544), v19, v21, v39)
	mBase = m.M
	v107 = m.ExcPending
	if v107 != 0 {
		goto L3
	} else {
		goto L46
	}
L44:
	;
	v110 = v102
	goto L45
L45:
	;
	v111 = *(*int32)(unsafe.Add(mBase, uint32(l1)+4))
	if v111 == int32(0) {
		goto L13
	} else {
		goto L47
	}
L46:
	;
	v110 = base.B2i32(v106 != int32(0))
	goto L45
L47:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v16)+8)) = v21
	*(*uint8)(unsafe.Add(mBase, uint32(v16)+1084)) = uint8(v110)
	v116 = *(*int32)(unsafe.Add(mBase, uint32(l1)+8))
	F_tuplestore_putvalues(m, v111, v116, v16+int32(8), v16+int32(1084))
	mBase = m.M
	v122 = m.ExcPending
	if v122 != 0 {
		goto L3
	} else {
		goto L48
	}
L48:
	;
	goto L1
L49:
	;
	v127 = int32(1)
	v128 = v21 + v127
	v131 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v21))))
	if v131&v127 != 0 {
		goto L50
	} else {
		goto L51
	}
L50:
	;
	v134 = v128
	goto L52
L51:
	;
	v134 = v21 + int32(4)
	goto L52
L52:
	;
	v138 = v134
	goto L53
L53:
	;
	v149 = *(*int32)(unsafe.Add(mBase, _c_F_split_text[1]))
	if v149 != 0 {
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
	v151 = m.ExcPending
	if v151 != 0 {
		goto L3
	} else {
		goto L58
	}
L56:
	;
	goto L57
L57:
	;
	v154 = F_text_position_next(m, v16+int32(8))
	mBase = m.M
	v155 = m.ExcPending
	if v155 != 0 {
		goto L3
	} else {
		goto L60
	}
L58:
	;
	goto L57
L59:
	;
	v189 = v187 - v138
	v191 = v189 + int32(4)
	v192 = F_palloc(m, v191)
	mBase = m.M
	v193 = m.ExcPending
	if v193 != 0 {
		goto L3
	} else {
		goto L75
	}
L60:
	;
	if v154 == int32(0) {
		goto L61
	} else {
		goto L62
	}
L61:
	;
	v158 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v21))))
	if v158 == int32(1) {
		goto L65
	} else {
		goto L66
	}
L62:
	;
	goto L63
L63:
	;
	v186 = *(*int32)(unsafe.Add(mBase, uint32(v16)+1060))
	v187 = v186
	v188 = v186
	goto L59
L64:
	;
	v187 = v183 + v21
	v188 = int32(0)
	goto L59
L65:
	;
	v162 = int32(18)
	v164 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v128))))
	if v164 == v162 {
		goto L68
	} else {
		goto L69
	}
L66:
	;
	goto L67
L67:
	;
	v175 = int32(1)
	if v158&v175 != 0 {
		v183 = int32(base.Ui32(v158) >> (uint(v175) % 32))
		goto L64
	} else {
		goto L74
	}
L68:
	;
	v167 = v162
	goto L70
L69:
	;
	v167 = int32(2)
	goto L70
L70:
	;
	if base.Ui32((v164-int32(1))&int32(255)) < base.Ui32(int32(3)) {
		goto L71
	} else {
		goto L72
	}
L71:
	;
	v174 = int32(6)
	goto L73
L72:
	;
	v174 = v167
	goto L73
L73:
	;
	v183 = v174
	goto L64
L74:
	;
	v179 = *(*int32)(unsafe.Add(mBase, uint32(v21)))
	v183 = int32(base.Ui32(v179) >> (uint(int32(2)) % 32))
	goto L64
L75:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v192))) = v191 << (uint(int32(2)) % 32)
	if v189 != 0 {
		goto L76
	} else {
		goto L77
	}
L76:
	;
	base.MemoryCopy(m, v192+int32(4), v138, v189)
	goto L78
L77:
	;
	goto L78
L78:
	;
	if v39 != 0 {
		goto L79
	} else {
		goto L80
	}
L79:
	;
	v201 = F_DirectFunctionCall2Coll(m, int32(1544), v19, v192, v39)
	mBase = m.M
	v202 = m.ExcPending
	if v202 != 0 {
		goto L3
	} else {
		goto L82
	}
L80:
	;
	v206 = int32(0)
	goto L81
L81:
	;
	v207 = *(*int32)(unsafe.Add(mBase, uint32(l1)+4))
	if v207 != 0 {
		goto L84
	} else {
		goto L85
	}
L82:
	;
	v206 = base.B2i32(v201 != int32(0))
	goto L81
L83:
	;
	F_pfree(m, v192)
	mBase = m.M
	v225 = m.ExcPending
	if v225 != 0 {
		goto L3
	} else {
		goto L89
	}
L84:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v16)+1084)) = v192
	*(*uint8)(unsafe.Add(mBase, uint32(v16)+1083)) = uint8(v206)
	v210 = *(*int32)(unsafe.Add(mBase, uint32(l1)+8))
	F_tuplestore_putvalues(m, v207, v210, v16+int32(1084), v16+int32(1083))
	mBase = m.M
	v216 = m.ExcPending
	if v216 != 0 {
		goto L3
	} else {
		goto L87
	}
L85:
	;
	goto L86
L86:
	;
	v217 = *(*int32)(unsafe.Add(mBase, uint32(l1)))
	v220 = *(*int32)(unsafe.Add(mBase, _c_F_split_text[0]))
	v221 = F_accumArrayResult(m, v217, v192, v206, int32(25), v220)
	mBase = m.M
	v222 = m.ExcPending
	if v222 != 0 {
		goto L3
	} else {
		goto L88
	}
L87:
	;
	goto L83
L88:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l1))) = v221
	goto L83
L89:
	;
	v226 = *(*int32)(unsafe.Add(mBase, uint32(v16)+1064))
	if v154 != 0 {
		v138 = v188 + v226
		goto L53
	} else {
		goto L90
	}
L90:
	;
	goto L54
L91:
	;
	v276 = v272
	v281 = v274
	goto L107
L92:
	;
	if v262 <= int32(0) {
		goto L1
	} else {
		goto L103
	}
L93:
	;
	v256 = *(*int32)(unsafe.Add(mBase, uint32(v21)))
	v262 = int32(base.Ui32(v256)>>(uint(int32(2))%32)) - int32(4)
	goto L92
L94:
	;
	v272 = v21 + int32(1)
	v274 = int32(4)
	v275 = v21 + int32(5)
	goto L91
L95:
	;
	v231 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v21)+1)))
	if base.Ui32((v231-int32(1))&int32(255)) < base.Ui32(int32(3)) {
		goto L94
	} else {
		goto L98
	}
L96:
	;
	goto L97
L97:
	;
	if v228&int32(1) == int32(0) {
		goto L93
	} else {
		goto L102
	}
L98:
	;
	if v231 == int32(18) {
		goto L99
	} else {
		goto L100
	}
L99:
	;
	v242 = int32(16)
	goto L101
L100:
	;
	v242 = int32(0)
	goto L101
L101:
	;
	v262 = v242
	goto L92
L102:
	;
	v247 = int32(1)
	v262 = int32(base.Ui32(v228)>>(uint(v247)%32)) - v247
	goto L92
L103:
	;
	v265 = int32(1)
	if v228&v265 != 0 {
		goto L104
	} else {
		goto L105
	}
L104:
	;
	v269 = v265
	goto L106
L105:
	;
	v269 = int32(4)
	goto L106
L106:
	;
	v270 = v21 + v269
	v272 = v270
	v274 = v262
	v275 = v270 + v262
	goto L91
L107:
	;
	v289 = F_pg_mblen_range(m, v276, v275)
	mBase = m.M
	v290 = m.ExcPending
	if v290 != 0 {
		goto L3
	} else {
		goto L109
	}
L108:
	;
	goto L1
L109:
	;
	v292 = *(*int32)(unsafe.Add(mBase, _c_F_split_text[1]))
	if v292 != 0 {
		goto L110
	} else {
		goto L111
	}
L110:
	;
	F_ProcessInterrupts(m)
	mBase = m.M
	v294 = m.ExcPending
	if v294 != 0 {
		goto L3
	} else {
		goto L113
	}
L111:
	;
	goto L112
L112:
	;
	v296 = v289 + int32(4)
	v297 = F_palloc(m, v296)
	mBase = m.M
	v298 = m.ExcPending
	if v298 != 0 {
		goto L3
	} else {
		goto L114
	}
L113:
	;
	goto L112
L114:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v297))) = v296 << (uint(int32(2)) % 32)
	if v289 != 0 {
		goto L115
	} else {
		goto L116
	}
L115:
	;
	base.MemoryCopy(m, v297+int32(4), v276, v289)
	goto L117
L116:
	;
	goto L117
L117:
	;
	if v39 != 0 {
		goto L118
	} else {
		goto L119
	}
L118:
	;
	v306 = F_DirectFunctionCall2Coll(m, int32(1544), v19, v297, v39)
	mBase = m.M
	v307 = m.ExcPending
	if v307 != 0 {
		goto L3
	} else {
		goto L121
	}
L119:
	;
	v311 = int32(0)
	goto L120
L120:
	;
	v312 = *(*int32)(unsafe.Add(mBase, uint32(l1)+4))
	if v312 != 0 {
		goto L123
	} else {
		goto L124
	}
L121:
	;
	v311 = base.B2i32(v306 != int32(0))
	goto L120
L122:
	;
	F_pfree(m, v297)
	mBase = m.M
	v330 = m.ExcPending
	if v330 != 0 {
		goto L3
	} else {
		goto L128
	}
L123:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v16)+8)) = v297
	*(*uint8)(unsafe.Add(mBase, uint32(v16)+1084)) = uint8(v311)
	v315 = *(*int32)(unsafe.Add(mBase, uint32(l1)+8))
	F_tuplestore_putvalues(m, v312, v315, v16+int32(8), v16+int32(1084))
	mBase = m.M
	v321 = m.ExcPending
	if v321 != 0 {
		goto L3
	} else {
		goto L126
	}
L124:
	;
	goto L125
L125:
	;
	v322 = *(*int32)(unsafe.Add(mBase, uint32(l1)))
	v325 = *(*int32)(unsafe.Add(mBase, _c_F_split_text[0]))
	v326 = F_accumArrayResult(m, v322, v297, v311, int32(25), v325)
	mBase = m.M
	v327 = m.ExcPending
	if v327 != 0 {
		goto L3
	} else {
		goto L127
	}
L126:
	;
	goto L122
L127:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l1))) = v326
	goto L122
L128:
	;
	v332 = v281 - v289
	if int32(0) < v332 {
		v276 = v276 + v289
		v281 = v332
		goto L107
	} else {
		goto L129
	}
L129:
	;
	goto L108
L130:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l1))) = v339
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
	var v92 int32
	_ = v92
	var v96 int32
	_ = v96
	var v100 int32
	_ = v100
	var v109 int32
	_ = v109
	var v116 int32
	_ = v116
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
						v96 = int32(1)
						if v20 != 0 {
							v109 = int32(base.Ui32(v18)>>(uint(v96)%32)) - v96
						} else {
							v100 = *(*int32)(unsafe.Add(mBase, uint32(v10)))
							v109 = int32(base.Ui32(v100)>>(uint(int32(2))%32)) - int32(4)
						}
						if v109 != 0 {
							v116 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v21))))
							return base.I32_extend8_s(v116)
						} else {
							return int32(0)
						}
					} else {
						v81 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v15))))
						if base.Ui32((v81-int32(1))&int32(255)) < base.Ui32(int32(3)) {
							v116 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v21))))
							return base.I32_extend8_s(v116)
						} else {
							if v81 == int32(18) {
								v92 = int32(16)
							} else {
								v92 = int32(0)
							}
							v109 = v92
							if v109 != 0 {
								v116 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v21))))
								return base.I32_extend8_s(v116)
							} else {
								return int32(0)
							}
						}
					}
				} else {
					v51 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v21)+1)))
					if v51&int32(248) != int32(48) {
						if v18 != int32(1) {
							v96 = int32(1)
							if v20 != 0 {
								v109 = int32(base.Ui32(v18)>>(uint(v96)%32)) - v96
							} else {
								v100 = *(*int32)(unsafe.Add(mBase, uint32(v10)))
								v109 = int32(base.Ui32(v100)>>(uint(int32(2))%32)) - int32(4)
							}
							if v109 != 0 {
								v116 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v21))))
								return base.I32_extend8_s(v116)
							} else {
								return int32(0)
							}
						} else {
							v81 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v15))))
							if base.Ui32((v81-int32(1))&int32(255)) < base.Ui32(int32(3)) {
								v116 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v21))))
								return base.I32_extend8_s(v116)
							} else {
								if v81 == int32(18) {
									v92 = int32(16)
								} else {
									v92 = int32(0)
								}
								v109 = v92
								if v109 != 0 {
									v116 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v21))))
									return base.I32_extend8_s(v116)
								} else {
									return int32(0)
								}
							}
						}
					} else {
						v56 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v21)+2)))
						if v56&int32(248) != int32(48) {
							if v18 != int32(1) {
								v96 = int32(1)
								if v20 != 0 {
									v109 = int32(base.Ui32(v18)>>(uint(v96)%32)) - v96
								} else {
									v100 = *(*int32)(unsafe.Add(mBase, uint32(v10)))
									v109 = int32(base.Ui32(v100)>>(uint(int32(2))%32)) - int32(4)
								}
								if v109 != 0 {
									v116 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v21))))
									return base.I32_extend8_s(v116)
								} else {
									return int32(0)
								}
							} else {
								v81 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v15))))
								if base.Ui32((v81-int32(1))&int32(255)) < base.Ui32(int32(3)) {
									v116 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v21))))
									return base.I32_extend8_s(v116)
								} else {
									if v81 == int32(18) {
										v92 = int32(16)
									} else {
										v92 = int32(0)
									}
									v109 = v92
									if v109 != 0 {
										v116 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v21))))
										return base.I32_extend8_s(v116)
									} else {
										return int32(0)
									}
								}
							}
						} else {
							v61 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v21)+3)))
							if v61&int32(248) != int32(48) {
								if v18 != int32(1) {
									v96 = int32(1)
									if v20 != 0 {
										v109 = int32(base.Ui32(v18)>>(uint(v96)%32)) - v96
									} else {
										v100 = *(*int32)(unsafe.Add(mBase, uint32(v10)))
										v109 = int32(base.Ui32(v100)>>(uint(int32(2))%32)) - int32(4)
									}
									if v109 != 0 {
										v116 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v21))))
										return base.I32_extend8_s(v116)
									} else {
										return int32(0)
									}
								} else {
									v81 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v15))))
									if base.Ui32((v81-int32(1))&int32(255)) < base.Ui32(int32(3)) {
										v116 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v21))))
										return base.I32_extend8_s(v116)
									} else {
										if v81 == int32(18) {
											v92 = int32(16)
										} else {
											v92 = int32(0)
										}
										v109 = v92
										if v109 != 0 {
											v116 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v21))))
											return base.I32_extend8_s(v116)
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
					v116 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v21))))
					return base.I32_extend8_s(v116)
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
				v96 = int32(1)
				if v20 != 0 {
					v109 = int32(base.Ui32(v18)>>(uint(v96)%32)) - v96
				} else {
					v100 = *(*int32)(unsafe.Add(mBase, uint32(v10)))
					v109 = int32(base.Ui32(v100)>>(uint(int32(2))%32)) - int32(4)
				}
				if v109 != 0 {
					v116 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v21))))
					return base.I32_extend8_s(v116)
				} else {
					return int32(0)
				}
			} else {
				v48 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v21))))
				if v48 != int32(92) {
					if v18 != int32(1) {
						v96 = int32(1)
						if v20 != 0 {
							v109 = int32(base.Ui32(v18)>>(uint(v96)%32)) - v96
						} else {
							v100 = *(*int32)(unsafe.Add(mBase, uint32(v10)))
							v109 = int32(base.Ui32(v100)>>(uint(int32(2))%32)) - int32(4)
						}
						if v109 != 0 {
							v116 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v21))))
							return base.I32_extend8_s(v116)
						} else {
							return int32(0)
						}
					} else {
						v81 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v15))))
						if base.Ui32((v81-int32(1))&int32(255)) < base.Ui32(int32(3)) {
							v116 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v21))))
							return base.I32_extend8_s(v116)
						} else {
							if v81 == int32(18) {
								v92 = int32(16)
							} else {
								v92 = int32(0)
							}
							v109 = v92
							if v109 != 0 {
								v116 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v21))))
								return base.I32_extend8_s(v116)
							} else {
								return int32(0)
							}
						}
					}
				} else {
					v51 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v21)+1)))
					if v51&int32(248) != int32(48) {
						if v18 != int32(1) {
							v96 = int32(1)
							if v20 != 0 {
								v109 = int32(base.Ui32(v18)>>(uint(v96)%32)) - v96
							} else {
								v100 = *(*int32)(unsafe.Add(mBase, uint32(v10)))
								v109 = int32(base.Ui32(v100)>>(uint(int32(2))%32)) - int32(4)
							}
							if v109 != 0 {
								v116 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v21))))
								return base.I32_extend8_s(v116)
							} else {
								return int32(0)
							}
						} else {
							v81 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v15))))
							if base.Ui32((v81-int32(1))&int32(255)) < base.Ui32(int32(3)) {
								v116 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v21))))
								return base.I32_extend8_s(v116)
							} else {
								if v81 == int32(18) {
									v92 = int32(16)
								} else {
									v92 = int32(0)
								}
								v109 = v92
								if v109 != 0 {
									v116 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v21))))
									return base.I32_extend8_s(v116)
								} else {
									return int32(0)
								}
							}
						}
					} else {
						v56 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v21)+2)))
						if v56&int32(248) != int32(48) {
							if v18 != int32(1) {
								v96 = int32(1)
								if v20 != 0 {
									v109 = int32(base.Ui32(v18)>>(uint(v96)%32)) - v96
								} else {
									v100 = *(*int32)(unsafe.Add(mBase, uint32(v10)))
									v109 = int32(base.Ui32(v100)>>(uint(int32(2))%32)) - int32(4)
								}
								if v109 != 0 {
									v116 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v21))))
									return base.I32_extend8_s(v116)
								} else {
									return int32(0)
								}
							} else {
								v81 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v15))))
								if base.Ui32((v81-int32(1))&int32(255)) < base.Ui32(int32(3)) {
									v116 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v21))))
									return base.I32_extend8_s(v116)
								} else {
									if v81 == int32(18) {
										v92 = int32(16)
									} else {
										v92 = int32(0)
									}
									v109 = v92
									if v109 != 0 {
										v116 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v21))))
										return base.I32_extend8_s(v116)
									} else {
										return int32(0)
									}
								}
							}
						} else {
							v61 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v21)+3)))
							if v61&int32(248) != int32(48) {
								if v18 != int32(1) {
									v96 = int32(1)
									if v20 != 0 {
										v109 = int32(base.Ui32(v18)>>(uint(v96)%32)) - v96
									} else {
										v100 = *(*int32)(unsafe.Add(mBase, uint32(v10)))
										v109 = int32(base.Ui32(v100)>>(uint(int32(2))%32)) - int32(4)
									}
									if v109 != 0 {
										v116 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v21))))
										return base.I32_extend8_s(v116)
									} else {
										return int32(0)
									}
								} else {
									v81 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v15))))
									if base.Ui32((v81-int32(1))&int32(255)) < base.Ui32(int32(3)) {
										v116 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v21))))
										return base.I32_extend8_s(v116)
									} else {
										if v81 == int32(18) {
											v92 = int32(16)
										} else {
											v92 = int32(0)
										}
										v109 = v92
										if v109 != 0 {
											v116 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v21))))
											return base.I32_extend8_s(v116)
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
	var v43 int32
	_ = v43
	var v46 int32
	_ = v46
	var v48 int32
	_ = v48
	var v60 int32
	_ = v60
	var v63 int32
	_ = v63
	var v67 int32
	_ = v67
	var v71 int32
	_ = v71
	var v76 int32
	_ = v76
	var v80 int32
	_ = v80
	var v82 int32
	_ = v82
	var v93 int32
	_ = v93
	var v96 int32
	_ = v96
	var v100 int32
	_ = v100
	var v105 int32
	_ = v105
	v4 = int32(0)
	v10 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v11 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v10))))
	v17 = base.B2i32(base.Ui32((v11-int32(48))&int32(255)) < base.Ui32(int32(10)))
	if v17 == v4 {
		v80 = v10
		v82 = v4
		goto L2
	} else {
		goto L3
	}
L1:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v93 = m.ExcPending
	if v93 != 0 {
		goto L12
	} else {
		goto L18
	}
L2:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0))) = v80
	*(*int32)(unsafe.Add(mBase, uint32(l2))) = v82
	return v17
L3:
	;
	v23 = v10
	v24 = v11
	v25 = v4
	goto L4
L4:
	;
	v31 = base.I64_extend_i32_s(v25) * int64(10)
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
	v60 = m.ExcPending
	if v60 != 0 {
		goto L12
	} else {
		goto L13
	}
L6:
	;
	v43 = (v24-int32(48))&int32(255) + v35
	if v43 < v35 {
		goto L1
	} else {
		goto L7
	}
L7:
	;
	v46 = v23 + int32(1)
	if base.Ui32(v46) < base.Ui32(l1) {
		goto L8
	} else {
		goto L9
	}
L8:
	;
	v48 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v46))))
	if base.Ui32(int32(9)) < base.Ui32((v48-int32(48))&int32(255)) {
		v80 = v46
		v82 = v43
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
	v23 = v46
	v24 = v48
	v25 = v43
	goto L4
L12:
	;
	return int32(0)
L13:
	;
	F_errcode(m, int32(50856066))
	mBase = m.M
	v63 = m.ExcPending
	if v63 != 0 {
		goto L12
	} else {
		goto L14
	}
L14:
	;
	F_errmsg(m, int32(_a_F_text_format_parse_digits_0), int32(0))
	mBase = m.M
	v67 = m.ExcPending
	if v67 != 0 {
		goto L12
	} else {
		goto L15
	}
L15:
	;
	F_errhint(m, int32(_a_F_text_format_parse_digits_1), int32(0))
	mBase = m.M
	v71 = m.ExcPending
	if v71 != 0 {
		goto L12
	} else {
		goto L16
	}
L16:
	;
	F_errfinish(m, int32(_a_F_text_format_parse_digits_2), int32(_a_F_text_format_parse_digits_3), int32(_a_F_text_format_parse_digits_4))
	mBase = m.M
	v76 = m.ExcPending
	if v76 != 0 {
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
	v96 = m.ExcPending
	if v96 != 0 {
		goto L12
	} else {
		goto L19
	}
L19:
	;
	F_errmsg(m, int32(_a_F_text_format_parse_digits_5), int32(0))
	mBase = m.M
	v100 = m.ExcPending
	if v100 != 0 {
		goto L12
	} else {
		goto L20
	}
L20:
	;
	F_errfinish(m, int32(_a_F_text_format_parse_digits_2), int32(_a_F_text_format_parse_digits_6), int32(_a_F_text_format_parse_digits_4))
	mBase = m.M
	v105 = m.ExcPending
	if v105 != 0 {
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
					F_errmsg(m, int32(_a_F_text_overlay_0), int32(0))
					v55 = m.ExcPending
					if v55 != 0 {
						return int32(0)
					} else {
						F_errfinish(m, int32(_a_F_text_overlay_1), int32(1186), int32(_a_F_text_overlay_2))
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
				F_errmsg(m, int32(_a_F_text_overlay_3), int32(0))
				v39 = m.ExcPending
				if v39 != 0 {
					return int32(0)
				} else {
					F_errfinish(m, int32(_a_F_text_overlay_1), int32(1182), int32(_a_F_text_overlay_2))
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
	var v5 int32
	_ = v5
	var v6 int32
	_ = v6
	var v9 int32
	_ = v9
	var v10 int32
	_ = v10
	var v11 int32
	_ = v11
	var v12 int32
	_ = v12
	var v17 int32
	_ = v17
	var v23 int32
	_ = v23
	var v26 int32
	_ = v26
	var v33 int32
	_ = v33
	var v34 int32
	_ = v34
	var v40 int32
	_ = v40
	var v46 int32
	_ = v46
	var v47 int32
	_ = v47
	var v53 int32
	_ = v53
	var v56 int32
	_ = v56
	var v63 int32
	_ = v63
	var v64 int32
	_ = v64
	var v70 int32
	_ = v70
	var v76 int32
	_ = v76
	var v77 int32
	_ = v77
	var v81 int32
	_ = v81
	var v83 int32
	_ = v83
	var v87 int32
	_ = v87
	var v89 int32
	_ = v89
	var v90 int32
	_ = v90
	var v91 int32
	_ = v91
	var v94 int32
	_ = v94
	var v95 int32
	_ = v95
	var v98 int32
	_ = v98
	var v99 int32
	_ = v99
	var v102 int32
	_ = v102
	v5 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v6 = F_pg_detoast_datum_packed(m, v5)
	mBase = m.M
	v9 = m.ExcPending
	if v9 != 0 {
		return int32(0)
	} else {
		v10 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
		v11 = F_pg_detoast_datum_packed(m, v10)
		mBase = m.M
		v12 = m.ExcPending
		if v12 != 0 {
			return int32(0)
		} else {
			v17 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v6))))
			if v17 == int32(1) {
				v23 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v6)+1)))
				if v23 == int32(18) {
					v26 = int32(16)
				} else {
					v26 = int32(0)
				}
				if base.Ui32((v23-int32(1))&int32(255)) < base.Ui32(int32(3)) {
					v33 = int32(4)
				} else {
					v33 = v26
				}
				v46 = v33
			} else {
				v34 = int32(1)
				if v17&v34 != 0 {
					v46 = int32(base.Ui32(v17)>>(uint(v34)%32)) - v34
				} else {
					v40 = *(*int32)(unsafe.Add(mBase, uint32(v6)))
					v46 = int32(base.Ui32(v40)>>(uint(int32(2))%32)) - int32(4)
				}
			}
			v47 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v11))))
			if v47 == int32(1) {
				v53 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v11)+1)))
				if v53 == int32(18) {
					v56 = int32(16)
				} else {
					v56 = int32(0)
				}
				if base.Ui32((v53-int32(1))&int32(255)) < base.Ui32(int32(3)) {
					v63 = int32(4)
				} else {
					v63 = v56
				}
				v76 = v63
			} else {
				v64 = int32(1)
				if v47&v64 != 0 {
					v76 = int32(base.Ui32(v47)>>(uint(v64)%32)) - v64
				} else {
					v70 = *(*int32)(unsafe.Add(mBase, uint32(v11)))
					v76 = int32(base.Ui32(v70)>>(uint(int32(2))%32)) - int32(4)
				}
			}
			v77 = int32(1)
			if v17&v77 != 0 {
				v81 = v77
			} else {
				v81 = int32(4)
			}
			v83 = int32(1)
			if v47&v83 != 0 {
				v87 = v83
			} else {
				v87 = int32(4)
			}
			v89 = base.B2i32(v46 < v76)
			if v46 < v76 {
				v90 = v46
			} else {
				v90 = v76
			}
			v91 = F_memcmp(m, v6+v81, v11+v87, v90)
			mBase = m.M
			if v91 != 0 {
				v94 = v91
			} else {
				if v46 < v76 {
					v94 = int32(-1)
				} else {
					v94 = base.B2i32(v76 < v46)
				}
			}
			v95 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
			if v95 != v6 {
				F_pfree(m, v6)
				mBase = m.M
				v98 = m.ExcPending
				if v98 != 0 {
					return int32(0)
				} else {
					v99 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
					if v99 != v11 {
						F_pfree(m, v11)
						mBase = m.M
						v102 = m.ExcPending
						if v102 != 0 {
							return int32(0)
						} else {
							return base.B2i32(v94 <= int32(0))
						}
					} else {
						return base.B2i32(v94 <= int32(0))
					}
				}
			} else {
				v99 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
				if v99 != v11 {
					F_pfree(m, v11)
					mBase = m.M
					v102 = m.ExcPending
					if v102 != 0 {
						return int32(0)
					} else {
						return base.B2i32(v94 <= int32(0))
					}
				} else {
					return base.B2i32(v94 <= int32(0))
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
	var v24 int32
	_ = v24
	var v27 int32
	_ = v27
	var v34 int32
	_ = v34
	var v35 int32
	_ = v35
	var v39 int32
	_ = v39
	var v45 int32
	_ = v45
	var v46 int32
	_ = v46
	var v47 int32
	_ = v47
	var v50 int32
	_ = v50
	var v51 int32
	_ = v51
	var v53 int32
	_ = v53
	var v55 int32
	_ = v55
	var v56 int32
	_ = v56
	var v57 int32
	_ = v57
	var v59 int32
	_ = v59
	var v60 int32
	_ = v60
	var v61 int32
	_ = v61
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
			v24 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v15))))
			if v24 == int32(18) {
				v27 = int32(16)
			} else {
				v27 = int32(0)
			}
			if base.Ui32((v24-int32(1))&int32(255)) < base.Ui32(int32(3)) {
				v34 = int32(4)
			} else {
				v34 = v27
			}
			v45 = v34
		} else {
			v35 = int32(1)
			if v18 != 0 {
				v45 = int32(base.Ui32(v16)>>(uint(v35)%32)) - v35
			} else {
				v39 = *(*int32)(unsafe.Add(mBase, uint32(v8)))
				v45 = int32(base.Ui32(v39)>>(uint(int32(2))%32)) - int32(4)
			}
		}
		if v18 != 0 {
			v46 = v15
		} else {
			v46 = v8 + int32(4)
		}
		v47 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
		if int32(0) <= v47 {
			v50 = F_pg_mbstrlen_with_len(m, v46, v45)
			mBase = m.M
			v51 = m.ExcPending
			if v51 != 0 {
				return int32(0)
			} else {
				v53 = v50
				v55 = F_pg_mbcharcliplen(m, v46, v45, v53-v47)
				mBase = m.M
				v56 = m.ExcPending
				if v56 != 0 {
					return int32(0)
				} else {
					v57 = v45 - v55
					v59 = v57 + int32(4)
					v60 = F_palloc(m, v59)
					mBase = m.M
					v61 = m.ExcPending
					if v61 != 0 {
						return int32(0)
					} else {
						*(*int32)(unsafe.Add(mBase, uint32(v60))) = v59 << (uint(int32(2)) % 32)
						if v57 != 0 {
							base.MemoryCopy(m, v60+int32(4), v46+v55, v57)
						} else {
						}
						return v60
					}
				}
			}
		} else {
			v53 = int32(0)
			v55 = F_pg_mbcharcliplen(m, v46, v45, v53-v47)
			mBase = m.M
			v56 = m.ExcPending
			if v56 != 0 {
				return int32(0)
			} else {
				v57 = v45 - v55
				v59 = v57 + int32(4)
				v60 = F_palloc(m, v59)
				mBase = m.M
				v61 = m.ExcPending
				if v61 != 0 {
					return int32(0)
				} else {
					*(*int32)(unsafe.Add(mBase, uint32(v60))) = v59 << (uint(int32(2)) % 32)
					if v57 != 0 {
						base.MemoryCopy(m, v60+int32(4), v46+v55, v57)
					} else {
					}
					return v60
				}
			}
		}
	}
}
