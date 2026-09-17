package p5

import (
	base "github.com/shibukawa/pgmem/internal/aot/pgaot/base"
	"unsafe"
)

func F_ExecMergeJoin(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v18 int32
	_ = v18
	var v20 int32
	_ = v20
	var v23 int32
	_ = v23
	var v27 int32
	_ = v27
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
	var v37 int32
	_ = v37
	var v39 int32
	_ = v39
	var v57 int32
	_ = v57
	var v60 int32
	_ = v60
	var v62 int32
	_ = v62
	var v63 int32
	_ = v63
	var v64 int32
	_ = v64
	var v65 int32
	_ = v65
	var v67 int32
	_ = v67
	var v68 int32
	_ = v68
	var v75 int32
	_ = v75
	var v76 int32
	_ = v76
	var v84 int32
	_ = v84
	var v88 int32
	_ = v88
	var v90 int32
	_ = v90
	var v91 int32
	_ = v91
	var v92 int32
	_ = v92
	var v93 int32
	_ = v93
	var v95 int32
	_ = v95
	var v96 int32
	_ = v96
	var v101 int32
	_ = v101
	var v105 int32
	_ = v105
	var v110 int32
	_ = v110
	var v111 int32
	_ = v111
	var v117 int32
	_ = v117
	var v123 int32
	_ = v123
	var v125 int32
	_ = v125
	var v127 int32
	_ = v127
	var v128 int32
	_ = v128
	var v130 int32
	_ = v130
	var v134 int32
	_ = v134
	var v135 int32
	_ = v135
	var v136 int32
	_ = v136
	var v142 int32
	_ = v142
	var v144 int32
	_ = v144
	var v149 int32
	_ = v149
	var v156 int32
	_ = v156
	var v157 int32
	_ = v157
	var v159 int32
	_ = v159
	var v163 int32
	_ = v163
	var v164 int32
	_ = v164
	var v165 int32
	_ = v165
	var v171 int32
	_ = v171
	var v172 int32
	_ = v172
	var v173 int32
	_ = v173
	var v174 int32
	_ = v174
	var v175 int32
	_ = v175
	var v177 int32
	_ = v177
	var v178 int32
	_ = v178
	var v179 int32
	_ = v179
	var v181 int32
	_ = v181
	var v186 int32
	_ = v186
	var v187 int32
	_ = v187
	var v188 int32
	_ = v188
	var v191 int32
	_ = v191
	var v193 int32
	_ = v193
	var v195 int32
	_ = v195
	var v196 int32
	_ = v196
	var v198 int32
	_ = v198
	var v201 float64
	_ = v201
	var v205 int32
	_ = v205
	var v208 float64
	_ = v208
	var v216 int32
	_ = v216
	var v217 int32
	_ = v217
	var v219 int32
	_ = v219
	var v220 int32
	_ = v220
	var v222 int32
	_ = v222
	var v224 int32
	_ = v224
	var v225 int32
	_ = v225
	var v226 int32
	_ = v226
	var v227 int32
	_ = v227
	var v228 int32
	_ = v228
	var v231 int32
	_ = v231
	var v232 int32
	_ = v232
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
	var v250 int32
	_ = v250
	var v255 int32
	_ = v255
	var v257 int32
	_ = v257
	var v266 int32
	_ = v266
	var v269 int32
	_ = v269
	var v270 int32
	_ = v270
	var v271 int32
	_ = v271
	var v273 int32
	_ = v273
	var v275 int32
	_ = v275
	var v296 int32
	_ = v296
	var v300 int32
	_ = v300
	var v302 int32
	_ = v302
	var v305 int32
	_ = v305
	var v306 int32
	_ = v306
	var v307 int32
	_ = v307
	var v311 int32
	_ = v311
	var v312 int32
	_ = v312
	var v313 int32
	_ = v313
	var v318 int32
	_ = v318
	var v331 int32
	_ = v331
	var v353 int32
	_ = v353
	var v354 int32
	_ = v354
	var v357 int32
	_ = v357
	var v358 int32
	_ = v358
	var v359 int32
	_ = v359
	var v360 int32
	_ = v360
	var v367 int32
	_ = v367
	var v369 int32
	_ = v369
	var v370 int32
	_ = v370
	var v389 int32
	_ = v389
	var v439 int32
	_ = v439
	var v440 int32
	_ = v440
	var v444 int32
	_ = v444
	var v449 int32
	_ = v449
	var v450 int32
	_ = v450
	var v453 int32
	_ = v453
	var v455 int32
	_ = v455
	var v456 int32
	_ = v456
	var v458 int32
	_ = v458
	var v460 int32
	_ = v460
	var v461 int32
	_ = v461
	var v462 int32
	_ = v462
	var v463 int32
	_ = v463
	var v464 int32
	_ = v464
	var v465 int32
	_ = v465
	var v470 int32
	_ = v470
	var v475 int32
	_ = v475
	var v478 int32
	_ = v478
	var v480 int32
	_ = v480
	var v481 int32
	_ = v481
	var v483 int32
	_ = v483
	var v487 int32
	_ = v487
	var v488 int32
	_ = v488
	var v490 int32
	_ = v490
	var v491 int32
	_ = v491
	var v492 int32
	_ = v492
	var v493 int32
	_ = v493
	var v494 int32
	_ = v494
	var v495 int32
	_ = v495
	var v500 int32
	_ = v500
	var v509 int32
	_ = v509
	var v510 int32
	_ = v510
	var v512 int32
	_ = v512
	var v513 int32
	_ = v513
	var v515 int32
	_ = v515
	var v519 int32
	_ = v519
	var v520 int32
	_ = v520
	var v522 int32
	_ = v522
	var v523 int32
	_ = v523
	var v524 int32
	_ = v524
	var v525 int32
	_ = v525
	var v526 int32
	_ = v526
	var v529 int32
	_ = v529
	var v530 int32
	_ = v530
	var v537 int32
	_ = v537
	var v540 int32
	_ = v540
	var v543 int32
	_ = v543
	var v550 int32
	_ = v550
	var v551 int32
	_ = v551
	var v553 int32
	_ = v553
	var v554 int32
	_ = v554
	var v556 int32
	_ = v556
	var v558 int32
	_ = v558
	var v559 int32
	_ = v559
	var v560 int32
	_ = v560
	var v561 int32
	_ = v561
	var v562 int32
	_ = v562
	var v565 int32
	_ = v565
	var v566 int32
	_ = v566
	var v573 int32
	_ = v573
	var v578 int32
	_ = v578
	var v581 int32
	_ = v581
	var v586 int32
	_ = v586
	var v587 int32
	_ = v587
	var v589 int32
	_ = v589
	var v590 int32
	_ = v590
	var v591 int32
	_ = v591
	var v592 int32
	_ = v592
	var v594 int32
	_ = v594
	var v597 int32
	_ = v597
	var v601 int32
	_ = v601
	var v607 int32
	_ = v607
	var v608 int32
	_ = v608
	var v617 int32
	_ = v617
	var v620 int32
	_ = v620
	var v621 int32
	_ = v621
	var v622 int32
	_ = v622
	var v624 int32
	_ = v624
	var v626 int32
	_ = v626
	var v647 int32
	_ = v647
	var v651 int32
	_ = v651
	var v653 int32
	_ = v653
	var v656 int32
	_ = v656
	var v657 int32
	_ = v657
	var v658 int32
	_ = v658
	var v662 int32
	_ = v662
	var v663 int32
	_ = v663
	var v664 int32
	_ = v664
	var v669 int32
	_ = v669
	var v682 int32
	_ = v682
	var v704 int32
	_ = v704
	var v705 int32
	_ = v705
	var v708 int32
	_ = v708
	var v709 int32
	_ = v709
	var v710 int32
	_ = v710
	var v711 int32
	_ = v711
	var v718 int32
	_ = v718
	var v720 int32
	_ = v720
	var v721 int32
	_ = v721
	var v740 int32
	_ = v740
	var v743 int32
	_ = v743
	var v747 int32
	_ = v747
	var v748 int32
	_ = v748
	var v749 int32
	_ = v749
	var v750 int32
	_ = v750
	var v751 int32
	_ = v751
	var v753 int32
	_ = v753
	var v800 int32
	_ = v800
	var v801 int32
	_ = v801
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
	var v809 int32
	_ = v809
	var v811 int32
	_ = v811
	var v814 int32
	_ = v814
	var v818 int32
	_ = v818
	var v824 int32
	_ = v824
	var v825 int32
	_ = v825
	var v834 int32
	_ = v834
	var v837 int32
	_ = v837
	var v838 int32
	_ = v838
	var v839 int32
	_ = v839
	var v841 int32
	_ = v841
	var v843 int32
	_ = v843
	var v864 int32
	_ = v864
	var v868 int32
	_ = v868
	var v870 int32
	_ = v870
	var v873 int32
	_ = v873
	var v874 int32
	_ = v874
	var v875 int32
	_ = v875
	var v879 int32
	_ = v879
	var v880 int32
	_ = v880
	var v881 int32
	_ = v881
	var v886 int32
	_ = v886
	var v899 int32
	_ = v899
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
	var v926 int32
	_ = v926
	var v933 int32
	_ = v933
	var v935 int32
	_ = v935
	var v936 int32
	_ = v936
	var v955 int32
	_ = v955
	var v958 int32
	_ = v958
	var v962 int32
	_ = v962
	var v1006 int32
	_ = v1006
	var v1007 int32
	_ = v1007
	var v1008 int32
	_ = v1008
	var v1040 int32
	_ = v1040
	var v1044 int32
	_ = v1044
	var v1049 int32
	_ = v1049
	var v1052 int32
	_ = v1052
	var v1053 int32
	_ = v1053
	var v1055 int32
	_ = v1055
	var v1056 int32
	_ = v1056
	var v1058 int32
	_ = v1058
	var v1060 int32
	_ = v1060
	var v1061 int32
	_ = v1061
	var v1062 int32
	_ = v1062
	var v1063 int32
	_ = v1063
	var v1064 int32
	_ = v1064
	var v1067 int32
	_ = v1067
	var v1068 int32
	_ = v1068
	var v1075 int32
	_ = v1075
	var v1080 int32
	_ = v1080
	var v1083 int32
	_ = v1083
	var v1089 int32
	_ = v1089
	var v1133 int32
	_ = v1133
	var v1137 int32
	_ = v1137
	var v1142 int32
	_ = v1142
	v18 = m.G0
	v20 = v18 - int32(16)
	m.G0 = v20
	v23 = *(*int32)(unsafe.Add(mBase, _c_F_ExecMergeJoin[0]))
	if v23 != 0 {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	F_ProcessInterrupts(m)
	mBase = m.M
	v27 = m.ExcPending
	if v27 != 0 {
		goto L4
	} else {
		goto L5
	}
L2:
	;
	goto L3
L3:
	;
	v28 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+132)))
	v29 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+131)))
	v30 = *(*int32)(unsafe.Add(mBase, uint32(l0)+32))
	v31 = *(*int32)(unsafe.Add(mBase, uint32(l0)+112))
	v32 = *(*int32)(unsafe.Add(mBase, uint32(l0)+36))
	v33 = *(*int32)(unsafe.Add(mBase, uint32(l0)+40))
	v34 = *(*int32)(unsafe.Add(mBase, uint32(l0)+64))
	v35 = *(*int32)(unsafe.Add(mBase, uint32(v34)+20))
	F_MemoryContextReset(m, v35)
	mBase = m.M
	v37 = m.ExcPending
	if v37 != 0 {
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
	v39 = v29 & int32(1)
	goto L7
L7:
	;
	v57 = *(*int32)(unsafe.Add(mBase, uint32(l0)+124))
	switch v57 - int32(1) {
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
	v1133 = m.ExcPending
	if v1133 != 0 {
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
	m.G0 = v20 + int32(16)
	return v1089
L12:
	;
	if v39 == int32(0) {
		goto L285
	} else {
		goto L286
	}
L13:
	;
	v800 = *(*int32)(unsafe.Add(mBase, uint32(l0)+144))
	v801 = F_MJEvalInnerValues(m, l0, v800)
	mBase = m.M
	v802 = m.ExcPending
	if v802 != 0 {
		goto L4
	} else {
		goto L235
	}
L14:
	;
	v586 = *(*int32)(unsafe.Add(mBase, uint32(l0)+64))
	v587 = *(*int32)(unsafe.Add(mBase, uint32(v586)+20))
	F_MemoryContextReset(m, v587)
	mBase = m.M
	v589 = m.ExcPending
	if v589 != 0 {
		goto L4
	} else {
		goto L196
	}
L15:
	;
	if v39 == int32(0) {
		goto L179
	} else {
		goto L180
	}
L16:
	;
	if v28&int32(1) == int32(0) {
		goto L158
	} else {
		goto L159
	}
L17:
	;
	v475 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+134)))
	if v475 == int32(0) {
		goto L142
	} else {
		goto L143
	}
L18:
	;
	v450 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+133)))
	if v450 == int32(0) {
		goto L130
	} else {
		goto L131
	}
L19:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v439 = m.ExcPending
	if v439 != 0 {
		goto L4
	} else {
		goto L127
	}
L20:
	;
	if v28&int32(1) == int32(0) {
		goto L81
	} else {
		goto L82
	}
L21:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+124)) = int32(6)
	v123 = *(*int32)(unsafe.Add(mBase, uint32(l0)+136))
	*(*int32)(unsafe.Add(mBase, uint32(v34)+12)) = v123
	v125 = *(*int32)(unsafe.Add(mBase, uint32(l0)+140))
	*(*int32)(unsafe.Add(mBase, uint32(v34)+8)) = v125
	if v31 != 0 {
		goto L59
	} else {
		goto L60
	}
L22:
	;
	v88 = *(*int32)(unsafe.Add(mBase, uint32(v33)+52))
	if v88 != 0 {
		goto L39
	} else {
		goto L40
	}
L23:
	;
	v60 = *(*int32)(unsafe.Add(mBase, uint32(v32)+52))
	if v60 != 0 {
		goto L24
	} else {
		goto L25
	}
L24:
	;
	F_ExecReScan(m, v32)
	mBase = m.M
	v62 = m.ExcPending
	if v62 != 0 {
		goto L4
	} else {
		goto L27
	}
L25:
	;
	goto L26
L26:
	;
	v63 = *(*int32)(unsafe.Add(mBase, uint32(v32)+12))
	v64 = m.T0[v63].(func(*base.Module, int32) int32)(m, v32)
	mBase = m.M
	v65 = m.ExcPending
	if v65 != 0 {
		goto L4
	} else {
		goto L28
	}
L27:
	;
	goto L26
L28:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+136)) = v64
	v67 = F_MJEvalOuterValues(m, l0)
	mBase = m.M
	v68 = m.ExcPending
	if v68 != 0 {
		goto L4
	} else {
		goto L32
	}
L29:
	;
	if v28&int32(1) == int32(0) {
		goto L36
	} else {
		goto L37
	}
L30:
	;
	if v39 == int32(0) {
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
	switch v67 - int32(1) {
	case 0:
		goto L30
	case 1:
		goto L29
	default:
		goto L31
	}
L33:
	;
	v75 = F_MJFillOuter(m, l0)
	mBase = m.M
	v76 = m.ExcPending
	if v76 != 0 {
		goto L4
	} else {
		goto L34
	}
L34:
	;
	if v75 == int32(0) {
		goto L7
	} else {
		goto L35
	}
L35:
	;
	v1089 = v75
	goto L11
L36:
	;
	v1089 = int32(0)
	goto L11
L37:
	;
	goto L38
L38:
	;
	v84 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(l0)+134)) = uint8(v84)
	*(*int32)(unsafe.Add(mBase, uint32(l0)+124)) = int32(10)
	goto L7
L39:
	;
	F_ExecReScan(m, v33)
	mBase = m.M
	v90 = m.ExcPending
	if v90 != 0 {
		goto L4
	} else {
		goto L42
	}
L40:
	;
	goto L41
L41:
	;
	v91 = *(*int32)(unsafe.Add(mBase, uint32(v33)+12))
	v92 = m.T0[v91].(func(*base.Module, int32) int32)(m, v33)
	mBase = m.M
	v93 = m.ExcPending
	if v93 != 0 {
		goto L4
	} else {
		goto L43
	}
L42:
	;
	goto L41
L43:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+140)) = v92
	v95 = F_MJEvalInnerValues(m, l0, v92)
	mBase = m.M
	v96 = m.ExcPending
	if v96 != 0 {
		goto L4
	} else {
		goto L47
	}
L44:
	;
	if v39 == int32(0) {
		goto L55
	} else {
		goto L56
	}
L45:
	;
	v101 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+129)))
	if v101 == int32(1) {
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
	switch v95 - int32(1) {
	case 0:
		goto L45
	case 1:
		goto L44
	default:
		goto L46
	}
L48:
	;
	F_ExecMarkPos(m, v33)
	mBase = m.M
	v105 = m.ExcPending
	if v105 != 0 {
		goto L4
	} else {
		goto L51
	}
L49:
	;
	goto L50
L50:
	;
	if v28&int32(1) == int32(0) {
		goto L7
	} else {
		goto L52
	}
L51:
	;
	goto L50
L52:
	;
	v110 = F_MJFillInner(m, l0)
	mBase = m.M
	v111 = m.ExcPending
	if v111 != 0 {
		goto L4
	} else {
		goto L53
	}
L53:
	;
	if v110 == int32(0) {
		goto L7
	} else {
		goto L54
	}
L54:
	;
	v1089 = v110
	goto L11
L55:
	;
	v1089 = int32(0)
	goto L11
L56:
	;
	goto L57
L57:
	;
	v117 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(l0)+133)) = uint8(v117)
	*(*int32)(unsafe.Add(mBase, uint32(l0)+124)) = int32(11)
	goto L7
L58:
	;
	v205 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	if v205 == int32(0) {
		goto L7
	} else {
		goto L80
	}
L59:
	;
	v127 = int32(_a_F_ExecMergeJoin_0)
	v128 = *(*int32)(unsafe.Add(mBase, _c_F_ExecMergeJoin[1]))
	v130 = *(*int32)(unsafe.Add(mBase, uint32(v34)+20))
	*(*int32)(unsafe.Add(mBase, _c_F_ExecMergeJoin[1])) = v130
	v134 = *(*int32)(unsafe.Add(mBase, uint32(v31)+20))
	v135 = m.T0[v134].(func(*base.Module, int32, int32, int32) int32)(m, v31, v34, v20+int32(14))
	mBase = m.M
	v136 = m.ExcPending
	if v136 != 0 {
		goto L4
	} else {
		goto L62
	}
L60:
	;
	goto L61
L61:
	;
	v142 = int32(257)
	*(*uint16)(unsafe.Add(mBase, uint32(l0)+133)) = uint16(v142)
	v144 = *(*int32)(unsafe.Add(mBase, uint32(l0)+104))
	if v144 == int32(5) {
		goto L64
	} else {
		goto L65
	}
L62:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_ExecMergeJoin[1])) = v128
	if v135 == int32(0) {
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
	v149 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+108)))
	if v149 == int32(1) {
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
	if v144 == int32(7) {
		goto L7
	} else {
		goto L70
	}
L70:
	;
	if v30 != 0 {
		goto L72
	} else {
		goto L73
	}
L71:
	;
	v198 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	if v198 == int32(0) {
		goto L7
	} else {
		goto L79
	}
L72:
	;
	v156 = int32(_a_F_ExecMergeJoin_0)
	v157 = *(*int32)(unsafe.Add(mBase, _c_F_ExecMergeJoin[1]))
	v159 = *(*int32)(unsafe.Add(mBase, uint32(v34)+20))
	*(*int32)(unsafe.Add(mBase, _c_F_ExecMergeJoin[1])) = v159
	v163 = *(*int32)(unsafe.Add(mBase, uint32(v30)+20))
	v164 = m.T0[v163].(func(*base.Module, int32, int32, int32) int32)(m, v30, v34, v20+int32(15))
	mBase = m.M
	v165 = m.ExcPending
	if v165 != 0 {
		goto L4
	} else {
		goto L75
	}
L73:
	;
	goto L74
L74:
	;
	v171 = *(*int32)(unsafe.Add(mBase, uint32(l0)+68))
	v172 = *(*int32)(unsafe.Add(mBase, uint32(v171)+72))
	v173 = *(*int32)(unsafe.Add(mBase, uint32(v171)+16))
	v174 = *(*int32)(unsafe.Add(mBase, uint32(v173)+8))
	v175 = *(*int32)(unsafe.Add(mBase, uint32(v174)+12))
	m.T0[v175].(func(*base.Module, int32))(m, v173)
	mBase = m.M
	v177 = m.ExcPending
	if v177 != 0 {
		goto L4
	} else {
		goto L77
	}
L75:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_ExecMergeJoin[1])) = v157
	if v164 == int32(0) {
		goto L71
	} else {
		goto L76
	}
L76:
	;
	goto L74
L77:
	;
	v178 = int32(_a_F_ExecMergeJoin_0)
	v179 = *(*int32)(unsafe.Add(mBase, _c_F_ExecMergeJoin[1]))
	v181 = *(*int32)(unsafe.Add(mBase, uint32(v172)+20))
	*(*int32)(unsafe.Add(mBase, _c_F_ExecMergeJoin[1])) = v181
	v186 = *(*int32)(unsafe.Add(mBase, uint32(v171)+24))
	v187 = m.T0[v186].(func(*base.Module, int32, int32, int32) int32)(m, v171+int32(4), v172, int32(0))
	mBase = m.M
	v188 = m.ExcPending
	if v188 != 0 {
		goto L4
	} else {
		goto L78
	}
L78:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_ExecMergeJoin[1])) = v179
	v191 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v173)+4)))
	v193 = v191 & int32(_a_F_ExecMergeJoin_1)
	*(*uint16)(unsafe.Add(mBase, uint32(v173)+4)) = uint16(v193)
	v195 = *(*int32)(unsafe.Add(mBase, uint32(v173)+12))
	v196 = *(*int32)(unsafe.Add(mBase, uint32(v195)))
	*(*uint16)(unsafe.Add(mBase, uint32(v173)+6)) = uint16(v196)
	v1089 = v173
	goto L11
L79:
	;
	v201 = *(*float64)(unsafe.Add(mBase, uint32(v198)+248))
	*(*float64)(unsafe.Add(mBase, uint32(v198)+248)) = base.F64_add(v201, float64(1))
	goto L7
L80:
	;
	v208 = *(*float64)(unsafe.Add(mBase, uint32(v205)+240))
	*(*float64)(unsafe.Add(mBase, uint32(v205)+240)) = base.F64_add(v208, float64(1))
	goto L7
L81:
	;
	v222 = *(*int32)(unsafe.Add(mBase, uint32(v33)+52))
	if v222 != 0 {
		goto L86
	} else {
		goto L87
	}
L82:
	;
	v216 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+134)))
	if v216 != 0 {
		goto L81
	} else {
		goto L83
	}
L83:
	;
	v217 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(l0)+134)) = uint8(v217)
	v219 = F_MJFillInner(m, l0)
	mBase = m.M
	v220 = m.ExcPending
	if v220 != 0 {
		goto L4
	} else {
		goto L84
	}
L84:
	;
	if v219 != 0 {
		v1089 = v219
		goto L11
	} else {
		goto L85
	}
L85:
	;
	goto L81
L86:
	;
	F_ExecReScan(m, v33)
	mBase = m.M
	v224 = m.ExcPending
	if v224 != 0 {
		goto L4
	} else {
		goto L89
	}
L87:
	;
	goto L88
L88:
	;
	v225 = *(*int32)(unsafe.Add(mBase, uint32(v33)+12))
	v226 = m.T0[v225].(func(*base.Module, int32) int32)(m, v33)
	mBase = m.M
	v227 = m.ExcPending
	if v227 != 0 {
		goto L4
	} else {
		goto L90
	}
L89:
	;
	goto L88
L90:
	;
	v228 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(l0)+134)) = uint8(v228)
	*(*int32)(unsafe.Add(mBase, uint32(l0)+140)) = v226
	v231 = F_MJEvalInnerValues(m, l0, v226)
	mBase = m.M
	v232 = m.ExcPending
	if v232 != 0 {
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
	v235 = *(*int32)(unsafe.Add(mBase, uint32(l0)+64))
	v236 = *(*int32)(unsafe.Add(mBase, uint32(v235)+20))
	F_MemoryContextReset(m, v236)
	mBase = m.M
	v238 = m.ExcPending
	if v238 != 0 {
		goto L4
	} else {
		goto L94
	}
L93:
	;
	switch v231 - int32(1) {
	case 0:
		goto L91
	case 1:
		goto L10
	default:
		goto L92
	}
L94:
	;
	v239 = int32(0)
	v240 = int32(_a_F_ExecMergeJoin_0)
	v241 = *(*int32)(unsafe.Add(mBase, _c_F_ExecMergeJoin[1]))
	v243 = *(*int32)(unsafe.Add(mBase, uint32(v235)+20))
	*(*int32)(unsafe.Add(mBase, _c_F_ExecMergeJoin[1])) = v243
	v246 = *(*int32)(unsafe.Add(mBase, uint32(l0)+116))
	if v239 < v246 {
		goto L97
	} else {
		goto L98
	}
L95:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_ExecMergeJoin[1])) = v241
	if int32(0) <= v367 {
		goto L9
	} else {
		goto L126
	}
L96:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_ExecMergeJoin[1])) = v241
	goto L9
L97:
	;
	v250 = v239
	v255 = v246
	v257 = v239
	goto L100
L98:
	;
	goto L99
L99:
	;
	v389 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+130)))
	if v389 != 0 {
		goto L96
	} else {
		goto L125
	}
L100:
	;
	v266 = *(*int32)(unsafe.Add(mBase, uint32(l0)+120))
	v269 = v266 + v250*int32(56)
	v270 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v269)+17)))
	v271 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v269)+16)))
	if v271 != 0 {
		goto L104
	} else {
		goto L105
	}
L101:
	;
	if v318 != 0 {
		goto L96
	} else {
		goto L124
	}
L102:
	;
	v353 = *(*int32)(unsafe.Add(mBase, uint32(v312)+8))
	v354 = *(*int32)(unsafe.Add(mBase, uint32(v312)+12))
	v357 = *(*int32)(unsafe.Add(mBase, uint32(v312)+36))
	v358 = m.T0[v357].(func(*base.Module, int32, int32, int32) int32)(m, v353, v354, v312+int32(20))
	mBase = m.M
	v359 = m.ExcPending
	if v359 != 0 {
		goto L4
	} else {
		goto L117
	}
L103:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_ExecMergeJoin[1])) = v241
	*(*int32)(unsafe.Add(mBase, uint32(l0)+124)) = int32(4)
	goto L7
L104:
	;
	v273 = v250
	v275 = v270
	goto L107
L105:
	;
	v311 = v250
	v312 = v269
	v313 = v270
	v318 = v257
	goto L106
L106:
	;
	if v313&int32(1) == int32(0) {
		goto L102
	} else {
		goto L115
	}
L107:
	;
	if v275&int32(1) == int32(0) {
		goto L109
	} else {
		goto L110
	}
L108:
	;
	v311 = v300
	v312 = v305
	v313 = v306
	v318 = v302
	goto L106
L109:
	;
	v296 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v266+v273*int32(56))+29)))
	if v296 == int32(0) {
		goto L96
	} else {
		goto L112
	}
L110:
	;
	goto L111
L111:
	;
	v300 = v273 + int32(1)
	if v255 <= v300 {
		goto L96
	} else {
		goto L113
	}
L112:
	;
	goto L103
L113:
	;
	v302 = int32(1)
	v305 = v266 + v300*int32(56)
	v306 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v305)+17)))
	v307 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v305)+16)))
	if v307 == v302 {
		v273 = v300
		v275 = v306
		goto L107
	} else {
		goto L114
	}
L114:
	;
	goto L108
L115:
	;
	v331 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v312)+29)))
	if v331 != 0 {
		goto L96
	} else {
		goto L116
	}
L116:
	;
	goto L103
L117:
	;
	v360 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v312)+28)))
	if v360 == int32(1) {
		goto L118
	} else {
		goto L119
	}
L118:
	;
	if v358 < int32(0) {
		goto L96
	} else {
		goto L121
	}
L119:
	;
	v367 = v358
	goto L120
L120:
	;
	if v367 != 0 {
		goto L95
	} else {
		goto L122
	}
L121:
	;
	v367 = int32(0) - v358
	goto L120
L122:
	;
	v369 = v311 + int32(1)
	v370 = *(*int32)(unsafe.Add(mBase, uint32(l0)+116))
	if v369 < v370 {
		v250 = v369
		v255 = v370
		v257 = v318
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
	*(*int32)(unsafe.Add(mBase, _c_F_ExecMergeJoin[1])) = v241
	*(*int32)(unsafe.Add(mBase, uint32(l0)+124)) = int32(3)
	goto L7
L126:
	;
	goto L91
L127:
	;
	v440 = *(*int32)(unsafe.Add(mBase, uint32(l0)+124))
	*(*int32)(unsafe.Add(mBase, uint32(v20))) = v440
	F_errmsg_internal(m, int32(_a_F_ExecMergeJoin_2), v20)
	mBase = m.M
	v444 = m.ExcPending
	if v444 != 0 {
		goto L4
	} else {
		goto L128
	}
L128:
	;
	F_errfinish(m, int32(_a_F_ExecMergeJoin_3), int32(1429), int32(_a_F_ExecMergeJoin_4))
	mBase = m.M
	v449 = m.ExcPending
	if v449 != 0 {
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
	v453 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(l0)+133)) = uint8(v453)
	v455 = F_MJFillOuter(m, l0)
	mBase = m.M
	v456 = m.ExcPending
	if v456 != 0 {
		goto L4
	} else {
		goto L133
	}
L131:
	;
	goto L132
L132:
	;
	v458 = *(*int32)(unsafe.Add(mBase, uint32(v32)+52))
	if v458 != 0 {
		goto L135
	} else {
		goto L136
	}
L133:
	;
	if v455 != 0 {
		v1089 = v455
		goto L11
	} else {
		goto L134
	}
L134:
	;
	goto L132
L135:
	;
	F_ExecReScan(m, v32)
	mBase = m.M
	v460 = m.ExcPending
	if v460 != 0 {
		goto L4
	} else {
		goto L138
	}
L136:
	;
	goto L137
L137:
	;
	v461 = int32(0)
	v462 = *(*int32)(unsafe.Add(mBase, uint32(v32)+12))
	v463 = m.T0[v462].(func(*base.Module, int32) int32)(m, v32)
	mBase = m.M
	v464 = m.ExcPending
	if v464 != 0 {
		goto L4
	} else {
		goto L139
	}
L138:
	;
	goto L137
L139:
	;
	v465 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(l0)+133)) = uint8(v465)
	*(*int32)(unsafe.Add(mBase, uint32(l0)+136)) = v463
	if v463 == v465 {
		v1089 = v461
		goto L11
	} else {
		goto L140
	}
L140:
	;
	v470 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v463)+4)))
	if v470&int32(2) == int32(0) {
		goto L7
	} else {
		goto L141
	}
L141:
	;
	v1089 = v461
	goto L11
L142:
	;
	v478 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(l0)+134)) = uint8(v478)
	v480 = F_MJFillInner(m, l0)
	mBase = m.M
	v481 = m.ExcPending
	if v481 != 0 {
		goto L4
	} else {
		goto L145
	}
L143:
	;
	goto L144
L144:
	;
	v483 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+129)))
	if v483 == int32(1) {
		goto L147
	} else {
		goto L148
	}
L145:
	;
	if v480 != 0 {
		v1089 = v480
		goto L11
	} else {
		goto L146
	}
L146:
	;
	goto L144
L147:
	;
	F_ExecMarkPos(m, v33)
	mBase = m.M
	v487 = m.ExcPending
	if v487 != 0 {
		goto L4
	} else {
		goto L150
	}
L148:
	;
	goto L149
L149:
	;
	v488 = *(*int32)(unsafe.Add(mBase, uint32(v33)+52))
	if v488 != 0 {
		goto L151
	} else {
		goto L152
	}
L150:
	;
	goto L149
L151:
	;
	F_ExecReScan(m, v33)
	mBase = m.M
	v490 = m.ExcPending
	if v490 != 0 {
		goto L4
	} else {
		goto L154
	}
L152:
	;
	goto L153
L153:
	;
	v491 = int32(0)
	v492 = *(*int32)(unsafe.Add(mBase, uint32(v33)+12))
	v493 = m.T0[v492].(func(*base.Module, int32) int32)(m, v33)
	mBase = m.M
	v494 = m.ExcPending
	if v494 != 0 {
		goto L4
	} else {
		goto L155
	}
L154:
	;
	goto L153
L155:
	;
	v495 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(l0)+134)) = uint8(v495)
	*(*int32)(unsafe.Add(mBase, uint32(l0)+140)) = v493
	if v493 == v495 {
		v1089 = v491
		goto L11
	} else {
		goto L156
	}
L156:
	;
	v500 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v493)+4)))
	if v500&int32(2) == int32(0) {
		goto L7
	} else {
		goto L157
	}
L157:
	;
	v1089 = v491
	goto L11
L158:
	;
	v515 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+129)))
	if v515 == int32(1) {
		goto L163
	} else {
		goto L164
	}
L159:
	;
	v509 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+134)))
	if v509 != 0 {
		goto L158
	} else {
		goto L160
	}
L160:
	;
	v510 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(l0)+134)) = uint8(v510)
	v512 = F_MJFillInner(m, l0)
	mBase = m.M
	v513 = m.ExcPending
	if v513 != 0 {
		goto L4
	} else {
		goto L161
	}
L161:
	;
	if v512 != 0 {
		v1089 = v512
		goto L11
	} else {
		goto L162
	}
L162:
	;
	goto L158
L163:
	;
	F_ExecMarkPos(m, v33)
	mBase = m.M
	v519 = m.ExcPending
	if v519 != 0 {
		goto L4
	} else {
		goto L166
	}
L164:
	;
	goto L165
L165:
	;
	v520 = *(*int32)(unsafe.Add(mBase, uint32(v33)+52))
	if v520 != 0 {
		goto L167
	} else {
		goto L168
	}
L166:
	;
	goto L165
L167:
	;
	F_ExecReScan(m, v33)
	mBase = m.M
	v522 = m.ExcPending
	if v522 != 0 {
		goto L4
	} else {
		goto L170
	}
L168:
	;
	goto L169
L169:
	;
	v523 = *(*int32)(unsafe.Add(mBase, uint32(v33)+12))
	v524 = m.T0[v523].(func(*base.Module, int32) int32)(m, v33)
	mBase = m.M
	v525 = m.ExcPending
	if v525 != 0 {
		goto L4
	} else {
		goto L171
	}
L170:
	;
	goto L169
L171:
	;
	v526 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(l0)+134)) = uint8(v526)
	*(*int32)(unsafe.Add(mBase, uint32(l0)+140)) = v524
	v529 = F_MJEvalInnerValues(m, l0, v524)
	mBase = m.M
	v530 = m.ExcPending
	if v530 != 0 {
		goto L4
	} else {
		goto L175
	}
L172:
	;
	v537 = int32(0)
	if v39 == v537 {
		v1089 = v537
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
	switch v529 - int32(1) {
	case 0:
		goto L173
	case 1:
		goto L172
	default:
		goto L174
	}
L176:
	;
	v540 = *(*int32)(unsafe.Add(mBase, uint32(l0)+136))
	if v540 == int32(0) {
		v1089 = v537
		goto L11
	} else {
		goto L177
	}
L177:
	;
	v543 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v540)+4)))
	if v543&int32(2) != 0 {
		v1089 = v537
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
	v556 = *(*int32)(unsafe.Add(mBase, uint32(v32)+52))
	if v556 != 0 {
		goto L184
	} else {
		goto L185
	}
L180:
	;
	v550 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+133)))
	if v550 != 0 {
		goto L179
	} else {
		goto L181
	}
L181:
	;
	v551 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(l0)+133)) = uint8(v551)
	v553 = F_MJFillOuter(m, l0)
	mBase = m.M
	v554 = m.ExcPending
	if v554 != 0 {
		goto L4
	} else {
		goto L182
	}
L182:
	;
	if v553 != 0 {
		v1089 = v553
		goto L11
	} else {
		goto L183
	}
L183:
	;
	goto L179
L184:
	;
	F_ExecReScan(m, v32)
	mBase = m.M
	v558 = m.ExcPending
	if v558 != 0 {
		goto L4
	} else {
		goto L187
	}
L185:
	;
	goto L186
L186:
	;
	v559 = *(*int32)(unsafe.Add(mBase, uint32(v32)+12))
	v560 = m.T0[v559].(func(*base.Module, int32) int32)(m, v32)
	mBase = m.M
	v561 = m.ExcPending
	if v561 != 0 {
		goto L4
	} else {
		goto L188
	}
L187:
	;
	goto L186
L188:
	;
	v562 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(l0)+133)) = uint8(v562)
	*(*int32)(unsafe.Add(mBase, uint32(l0)+136)) = v560
	v565 = F_MJEvalOuterValues(m, l0)
	mBase = m.M
	v566 = m.ExcPending
	if v566 != 0 {
		goto L4
	} else {
		goto L192
	}
L189:
	;
	v573 = int32(0)
	if v28&int32(1) == v573 {
		v1089 = v573
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
	switch v565 - int32(1) {
	case 0:
		goto L190
	case 1:
		goto L189
	default:
		goto L191
	}
L193:
	;
	v578 = *(*int32)(unsafe.Add(mBase, uint32(l0)+140))
	if v578 == int32(0) {
		v1089 = v573
		goto L11
	} else {
		goto L194
	}
L194:
	;
	v581 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v578)+4)))
	if v581&int32(2) != 0 {
		v1089 = v573
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
	v590 = int32(0)
	v591 = int32(_a_F_ExecMergeJoin_0)
	v592 = *(*int32)(unsafe.Add(mBase, _c_F_ExecMergeJoin[1]))
	v594 = *(*int32)(unsafe.Add(mBase, uint32(v586)+20))
	*(*int32)(unsafe.Add(mBase, _c_F_ExecMergeJoin[1])) = v594
	v597 = *(*int32)(unsafe.Add(mBase, uint32(l0)+116))
	if v590 < v597 {
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
	*(*int32)(unsafe.Add(mBase, _c_F_ExecMergeJoin[1])) = v592
	if int32(0) <= v718 {
		goto L197
	} else {
		goto L234
	}
L199:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_ExecMergeJoin[1])) = v592
	goto L197
L200:
	;
	v601 = v590
	v607 = v597
	v608 = v590
	goto L203
L201:
	;
	goto L202
L202:
	;
	v740 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+130)))
	if v740 != 0 {
		goto L199
	} else {
		goto L228
	}
L203:
	;
	v617 = *(*int32)(unsafe.Add(mBase, uint32(l0)+120))
	v620 = v617 + v601*int32(56)
	v621 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v620)+17)))
	v622 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v620)+16)))
	if v622 != 0 {
		goto L207
	} else {
		goto L208
	}
L204:
	;
	if v669 != 0 {
		goto L199
	} else {
		goto L227
	}
L205:
	;
	v704 = *(*int32)(unsafe.Add(mBase, uint32(v663)+8))
	v705 = *(*int32)(unsafe.Add(mBase, uint32(v663)+12))
	v708 = *(*int32)(unsafe.Add(mBase, uint32(v663)+36))
	v709 = m.T0[v708].(func(*base.Module, int32, int32, int32) int32)(m, v704, v705, v663+int32(20))
	mBase = m.M
	v710 = m.ExcPending
	if v710 != 0 {
		goto L4
	} else {
		goto L220
	}
L206:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_ExecMergeJoin[1])) = v592
	*(*int32)(unsafe.Add(mBase, uint32(l0)+124)) = int32(8)
	goto L7
L207:
	;
	v624 = v601
	v626 = v621
	goto L210
L208:
	;
	v662 = v601
	v663 = v620
	v664 = v621
	v669 = v608
	goto L209
L209:
	;
	if v664&int32(1) == int32(0) {
		goto L205
	} else {
		goto L218
	}
L210:
	;
	if v626&int32(1) == int32(0) {
		goto L212
	} else {
		goto L213
	}
L211:
	;
	v662 = v651
	v663 = v656
	v664 = v657
	v669 = v653
	goto L209
L212:
	;
	v647 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v617+v624*int32(56))+29)))
	if v647 == int32(0) {
		goto L199
	} else {
		goto L215
	}
L213:
	;
	goto L214
L214:
	;
	v651 = v624 + int32(1)
	if v607 <= v651 {
		goto L199
	} else {
		goto L216
	}
L215:
	;
	goto L206
L216:
	;
	v653 = int32(1)
	v656 = v617 + v651*int32(56)
	v657 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v656)+17)))
	v658 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v656)+16)))
	if v658 == v653 {
		v624 = v651
		v626 = v657
		goto L210
	} else {
		goto L217
	}
L217:
	;
	goto L211
L218:
	;
	v682 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v663)+29)))
	if v682 != 0 {
		goto L199
	} else {
		goto L219
	}
L219:
	;
	goto L206
L220:
	;
	v711 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v663)+28)))
	if v711 == int32(1) {
		goto L221
	} else {
		goto L222
	}
L221:
	;
	if v709 < int32(0) {
		goto L199
	} else {
		goto L224
	}
L222:
	;
	v718 = v709
	goto L223
L223:
	;
	if v718 != 0 {
		goto L198
	} else {
		goto L225
	}
L224:
	;
	v718 = int32(0) - v709
	goto L223
L225:
	;
	v720 = v662 + int32(1)
	v721 = *(*int32)(unsafe.Add(mBase, uint32(l0)+116))
	if v720 < v721 {
		v601 = v720
		v607 = v721
		v608 = v669
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
	*(*int32)(unsafe.Add(mBase, _c_F_ExecMergeJoin[1])) = v592
	v743 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+128)))
	if v743 == int32(0) {
		goto L229
	} else {
		goto L230
	}
L229:
	;
	F_ExecMarkPos(m, v33)
	mBase = m.M
	v747 = m.ExcPending
	if v747 != 0 {
		goto L4
	} else {
		goto L232
	}
L230:
	;
	goto L231
L231:
	;
	v748 = *(*int32)(unsafe.Add(mBase, uint32(l0)+144))
	v749 = *(*int32)(unsafe.Add(mBase, uint32(l0)+140))
	v750 = *(*int32)(unsafe.Add(mBase, uint32(v748)+8))
	v751 = *(*int32)(unsafe.Add(mBase, uint32(v750)+32))
	m.T0[v751].(func(*base.Module, int32, int32))(m, v748, v749)
	mBase = m.M
	v753 = m.ExcPending
	if v753 != 0 {
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
	v803 = *(*int32)(unsafe.Add(mBase, uint32(l0)+64))
	v804 = *(*int32)(unsafe.Add(mBase, uint32(v803)+20))
	F_MemoryContextReset(m, v804)
	mBase = m.M
	v806 = m.ExcPending
	if v806 != 0 {
		goto L4
	} else {
		goto L236
	}
L236:
	;
	v807 = int32(0)
	v808 = int32(_a_F_ExecMergeJoin_0)
	v809 = *(*int32)(unsafe.Add(mBase, _c_F_ExecMergeJoin[1]))
	v811 = *(*int32)(unsafe.Add(mBase, uint32(v803)+20))
	*(*int32)(unsafe.Add(mBase, _c_F_ExecMergeJoin[1])) = v811
	v814 = *(*int32)(unsafe.Add(mBase, uint32(l0)+116))
	if v807 < v814 {
		goto L241
	} else {
		goto L242
	}
L237:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v1040 = m.ExcPending
	if v1040 != 0 {
		goto L4
	} else {
		goto L282
	}
L238:
	;
	v1006 = *(*int32)(unsafe.Add(mBase, uint32(l0)+140))
	v1007 = F_MJEvalInnerValues(m, l0, v1006)
	mBase = m.M
	v1008 = m.ExcPending
	if v1008 != 0 {
		goto L4
	} else {
		goto L278
	}
L239:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_ExecMergeJoin[1])) = v809
	if v933 <= int32(0) {
		goto L237
	} else {
		goto L274
	}
L240:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_ExecMergeJoin[1])) = v809
	goto L238
L241:
	;
	v818 = v807
	v824 = v814
	v825 = v807
	goto L244
L242:
	;
	goto L243
L243:
	;
	v955 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+130)))
	if v955 != 0 {
		goto L240
	} else {
		goto L269
	}
L244:
	;
	v834 = *(*int32)(unsafe.Add(mBase, uint32(l0)+120))
	v837 = v834 + v818*int32(56)
	v838 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v837)+17)))
	v839 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v837)+16)))
	if v839 != 0 {
		goto L248
	} else {
		goto L249
	}
L245:
	;
	if v886 != 0 {
		goto L240
	} else {
		goto L268
	}
L246:
	;
	v919 = *(*int32)(unsafe.Add(mBase, uint32(v880)+8))
	v920 = *(*int32)(unsafe.Add(mBase, uint32(v880)+12))
	v923 = *(*int32)(unsafe.Add(mBase, uint32(v880)+36))
	v924 = m.T0[v923].(func(*base.Module, int32, int32, int32) int32)(m, v919, v920, v880+int32(20))
	mBase = m.M
	v925 = m.ExcPending
	if v925 != 0 {
		goto L4
	} else {
		goto L261
	}
L247:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_ExecMergeJoin[1])) = v809
	goto L237
L248:
	;
	v841 = v818
	v843 = v838
	goto L251
L249:
	;
	v879 = v818
	v880 = v837
	v881 = v838
	v886 = v825
	goto L250
L250:
	;
	if v881&int32(1) == int32(0) {
		goto L246
	} else {
		goto L259
	}
L251:
	;
	if v843&int32(1) == int32(0) {
		goto L253
	} else {
		goto L254
	}
L252:
	;
	v879 = v868
	v880 = v873
	v881 = v874
	v886 = v870
	goto L250
L253:
	;
	v864 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v834+v841*int32(56))+29)))
	if v864 == int32(0) {
		goto L240
	} else {
		goto L256
	}
L254:
	;
	goto L255
L255:
	;
	v868 = v841 + int32(1)
	if v824 <= v868 {
		goto L240
	} else {
		goto L257
	}
L256:
	;
	goto L247
L257:
	;
	v870 = int32(1)
	v873 = v834 + v868*int32(56)
	v874 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v873)+17)))
	v875 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v873)+16)))
	if v875 == v870 {
		v841 = v868
		v843 = v874
		goto L251
	} else {
		goto L258
	}
L258:
	;
	goto L252
L259:
	;
	v899 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v880)+29)))
	if v899 != 0 {
		goto L240
	} else {
		goto L260
	}
L260:
	;
	goto L247
L261:
	;
	v926 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v880)+28)))
	if v926 == int32(1) {
		goto L262
	} else {
		goto L263
	}
L262:
	;
	if v924 < int32(0) {
		goto L240
	} else {
		goto L265
	}
L263:
	;
	v933 = v924
	goto L264
L264:
	;
	if v933 != 0 {
		goto L239
	} else {
		goto L266
	}
L265:
	;
	v933 = int32(0) - v924
	goto L264
L266:
	;
	v935 = v879 + int32(1)
	v936 = *(*int32)(unsafe.Add(mBase, uint32(l0)+116))
	if v935 < v936 {
		v818 = v935
		v824 = v936
		v825 = v886
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
	*(*int32)(unsafe.Add(mBase, _c_F_ExecMergeJoin[1])) = v809
	v958 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+128)))
	if v958 == int32(0) {
		goto L270
	} else {
		goto L271
	}
L270:
	;
	F_ExecRestrPos(m, v33)
	mBase = m.M
	v962 = m.ExcPending
	if v962 != 0 {
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
	*(*int32)(unsafe.Add(mBase, uint32(l0)+140)) = v800
	goto L272
L274:
	;
	goto L238
L275:
	;
	if v39 == int32(0) {
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
	switch v1007 - int32(1) {
	case 0:
		goto L276
	case 1:
		goto L275
	default:
		goto L277
	}
L279:
	;
	v1089 = int32(0)
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
	v1044 = m.ExcPending
	if v1044 != 0 {
		goto L4
	} else {
		goto L283
	}
L283:
	;
	F_errfinish(m, int32(_a_F_ExecMergeJoin_3), int32(1145), int32(_a_F_ExecMergeJoin_4))
	mBase = m.M
	v1049 = m.ExcPending
	if v1049 != 0 {
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
	v1058 = *(*int32)(unsafe.Add(mBase, uint32(v32)+52))
	if v1058 != 0 {
		goto L290
	} else {
		goto L291
	}
L286:
	;
	v1052 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+133)))
	if v1052 != 0 {
		goto L285
	} else {
		goto L287
	}
L287:
	;
	v1053 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(l0)+133)) = uint8(v1053)
	v1055 = F_MJFillOuter(m, l0)
	mBase = m.M
	v1056 = m.ExcPending
	if v1056 != 0 {
		goto L4
	} else {
		goto L288
	}
L288:
	;
	if v1055 != 0 {
		v1089 = v1055
		goto L11
	} else {
		goto L289
	}
L289:
	;
	goto L285
L290:
	;
	F_ExecReScan(m, v32)
	mBase = m.M
	v1060 = m.ExcPending
	if v1060 != 0 {
		goto L4
	} else {
		goto L293
	}
L291:
	;
	goto L292
L292:
	;
	v1061 = *(*int32)(unsafe.Add(mBase, uint32(v32)+12))
	v1062 = m.T0[v1061].(func(*base.Module, int32) int32)(m, v32)
	mBase = m.M
	v1063 = m.ExcPending
	if v1063 != 0 {
		goto L4
	} else {
		goto L294
	}
L293:
	;
	goto L292
L294:
	;
	v1064 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(l0)+133)) = uint8(v1064)
	*(*int32)(unsafe.Add(mBase, uint32(l0)+136)) = v1062
	v1067 = F_MJEvalOuterValues(m, l0)
	mBase = m.M
	v1068 = m.ExcPending
	if v1068 != 0 {
		goto L4
	} else {
		goto L298
	}
L295:
	;
	v1075 = int32(0)
	if v28&int32(1) == v1075 {
		v1089 = v1075
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
	switch v1067 - int32(1) {
	case 0:
		goto L296
	case 1:
		goto L295
	default:
		goto L297
	}
L299:
	;
	v1080 = *(*int32)(unsafe.Add(mBase, uint32(l0)+140))
	if v1080 == int32(0) {
		v1089 = v1075
		goto L11
	} else {
		goto L300
	}
L300:
	;
	v1083 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1080)+4)))
	if v1083&int32(2) != 0 {
		v1089 = v1075
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
	v1137 = m.ExcPending
	if v1137 != 0 {
		goto L4
	} else {
		goto L303
	}
L303:
	;
	F_errfinish(m, int32(_a_F_ExecMergeJoin_3), int32(902), int32(_a_F_ExecMergeJoin_4))
	mBase = m.M
	v1142 = m.ExcPending
	if v1142 != 0 {
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
	var v12 int32
	_ = v12
	var v15 int32
	_ = v15
	var v18 int32
	_ = v18
	var v26 int32
	_ = v26
	var v33 int32
	_ = v33
	var v34 int32
	_ = v34
	var v36 int32
	_ = v36
	var v38 int32
	_ = v38
	var v41 int32
	_ = v41
	var v43 int32
	_ = v43
	var v46 int32
	_ = v46
	var v59 int32
	_ = v59
	var v62 int32
	_ = v62
	var v63 int32
	_ = v63
	var v64 int32
	_ = v64
	var v65 int32
	_ = v65
	var v67 int32
	_ = v67
	var v69 int32
	_ = v69
	var v72 int32
	_ = v72
	var v75 int32
	_ = v75
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
	var v87 int32
	_ = v87
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
	var v96 int32
	_ = v96
	var v97 int32
	_ = v97
	var v98 int32
	_ = v98
	var v99 int32
	_ = v99
	var v101 int32
	_ = v101
	var v104 int32
	_ = v104
	var v105 int32
	_ = v105
	var v106 int32
	_ = v106
	var v107 int32
	_ = v107
	var v109 int32
	_ = v109
	var v116 int32
	_ = v116
	var v117 int32
	_ = v117
	var v118 int32
	_ = v118
	var v120 int32
	_ = v120
	var v122 int32
	_ = v122
	var v131 int32
	_ = v131
	v7 = int32(0)
	v12 = *(*int32)(unsafe.Add(mBase, uint32(l1)+24))
	if v12 == v7 {
		goto L2
	} else {
		goto L3
	}
L1:
	;
	return v131
L2:
	;
	v59 = F_ec_search_derived_clause_for_ems(m, l0, l1, l3, l4, l5)
	mBase = m.M
	v62 = m.ExcPending
	if v62 != 0 {
		goto L16
	} else {
		goto L17
	}
L3:
	;
	v15 = *(*int32)(unsafe.Add(mBase, uint32(v12)+4))
	if v15 <= int32(0) {
		goto L2
	} else {
		goto L4
	}
L4:
	;
	v18 = *(*int32)(unsafe.Add(mBase, uint32(v12)+12))
	v26 = v7
	goto L5
L5:
	;
	v33 = *(*int32)(unsafe.Add(mBase, uint32(v18+v26<<(uint(int32(2))%32))))
	v34 = *(*int32)(unsafe.Add(mBase, uint32(v33)+108))
	if v34 != l3 {
		goto L7
	} else {
		goto L8
	}
L6:
	;
	goto L2
L7:
	;
	if l4 != v34 {
		goto L11
	} else {
		goto L12
	}
L8:
	;
	v36 = *(*int32)(unsafe.Add(mBase, uint32(v33)+112))
	if v36 != l4 {
		goto L7
	} else {
		goto L9
	}
L9:
	;
	v38 = *(*int32)(unsafe.Add(mBase, uint32(v33)+60))
	if v38 == l5 {
		v131 = v33
		goto L1
	} else {
		goto L10
	}
L10:
	;
	goto L7
L11:
	;
	v46 = v26 + int32(1)
	if v15 != v46 {
		v26 = v46
		goto L5
	} else {
		goto L15
	}
L12:
	;
	v41 = *(*int32)(unsafe.Add(mBase, uint32(v33)+112))
	if v41 != l3 {
		goto L11
	} else {
		goto L13
	}
L13:
	;
	v43 = *(*int32)(unsafe.Add(mBase, uint32(v33)+60))
	if v43 == l5 {
		v131 = v33
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
	if v59 != 0 {
		v131 = v59
		goto L1
	} else {
		goto L18
	}
L18:
	;
	v63 = int32(0)
	v64 = int32(_a_F_create_join_clause_0)
	v65 = *(*int32)(unsafe.Add(mBase, _c_F_create_join_clause[0]))
	v67 = *(*int32)(unsafe.Add(mBase, uint32(l0)+280))
	*(*int32)(unsafe.Add(mBase, _c_F_create_join_clause[0])) = v67
	v69 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l3)+13)))
	if v69 == v63 {
		goto L20
	} else {
		goto L21
	}
L19:
	;
	v83 = *(*int32)(unsafe.Add(mBase, uint32(l1)+8))
	v84 = *(*int32)(unsafe.Add(mBase, uint32(l3)+4))
	v85 = *(*int32)(unsafe.Add(mBase, uint32(l4)+4))
	v86 = *(*int32)(unsafe.Add(mBase, uint32(l3)+8))
	v87 = *(*int32)(unsafe.Add(mBase, uint32(l4)+8))
	v88 = F_bms_union(m, v86, v87)
	mBase = m.M
	v89 = m.ExcPending
	if v89 != 0 {
		goto L16
	} else {
		goto L31
	}
L20:
	;
	v72 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l4)+13)))
	if v72 != int32(1) {
		v82 = v63
		goto L19
	} else {
		goto L23
	}
L21:
	;
	goto L22
L22:
	;
	v75 = *(*int32)(unsafe.Add(mBase, uint32(l3)+24))
	if v75 != 0 {
		goto L24
	} else {
		goto L25
	}
L23:
	;
	goto L22
L24:
	;
	v76 = v75
	goto L26
L25:
	;
	v76 = l3
	goto L26
L26:
	;
	v77 = *(*int32)(unsafe.Add(mBase, uint32(l4)+24))
	if v77 != 0 {
		goto L27
	} else {
		goto L28
	}
L27:
	;
	v78 = v77
	goto L29
L28:
	;
	v78 = l4
	goto L29
L29:
	;
	v79 = F_create_join_clause(m, l0, l1, l2, v76, v78, l5)
	mBase = m.M
	v80 = m.ExcPending
	if v80 != 0 {
		goto L16
	} else {
		goto L30
	}
L30:
	;
	v82 = v79
	goto L19
L31:
	;
	v90 = *(*int32)(unsafe.Add(mBase, uint32(l1)+48))
	v91 = F_build_implied_join_equality(m, l0, l2, v83, v84, v85, v88, v90)
	mBase = m.M
	v92 = m.ExcPending
	if v92 != 0 {
		goto L16
	} else {
		goto L32
	}
L32:
	;
	v93 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l3)+13)))
	if v93 == int32(1) {
		goto L33
	} else {
		goto L34
	}
L33:
	;
	v96 = *(*int32)(unsafe.Add(mBase, uint32(v91)+28))
	v97 = *(*int32)(unsafe.Add(mBase, uint32(l3)+8))
	v98 = F_bms_add_members(m, v96, v97)
	mBase = m.M
	v99 = m.ExcPending
	if v99 != 0 {
		goto L16
	} else {
		goto L36
	}
L34:
	;
	goto L35
L35:
	;
	v101 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l4)+13)))
	if v101 == int32(1) {
		goto L37
	} else {
		goto L38
	}
L36:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v91)+28)) = v98
	goto L35
L37:
	;
	v104 = *(*int32)(unsafe.Add(mBase, uint32(v91)+28))
	v105 = *(*int32)(unsafe.Add(mBase, uint32(l4)+8))
	v106 = F_bms_add_members(m, v104, v105)
	mBase = m.M
	v107 = m.ExcPending
	if v107 != 0 {
		goto L16
	} else {
		goto L40
	}
L38:
	;
	goto L39
L39:
	;
	if v82 != 0 {
		goto L41
	} else {
		goto L42
	}
L40:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v91)+28)) = v106
	goto L39
L41:
	;
	v109 = *(*int32)(unsafe.Add(mBase, uint32(v82)+56))
	*(*int32)(unsafe.Add(mBase, uint32(v91)+56)) = v109
	goto L43
L42:
	;
	goto L43
L43:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v91)+112)) = l4
	*(*int32)(unsafe.Add(mBase, uint32(v91)+108)) = l3
	*(*int32)(unsafe.Add(mBase, uint32(v91)+104)) = l1
	*(*int32)(unsafe.Add(mBase, uint32(v91)+100)) = l1
	*(*int32)(unsafe.Add(mBase, uint32(v91)+60)) = l5
	v116 = *(*int32)(unsafe.Add(mBase, uint32(l1)+28))
	v117 = F_lappend(m, v116, v91)
	mBase = m.M
	v118 = m.ExcPending
	if v118 != 0 {
		goto L16
	} else {
		goto L44
	}
L44:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l1)+28)) = v117
	v120 = *(*int32)(unsafe.Add(mBase, uint32(l1)+32))
	if v120 != 0 {
		goto L45
	} else {
		goto L46
	}
L45:
	;
	F_ec_add_clause_to_derives_hash(m, l1, v91)
	mBase = m.M
	v122 = m.ExcPending
	if v122 != 0 {
		goto L16
	} else {
		goto L48
	}
L46:
	;
	goto L47
L47:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_create_join_clause[0])) = v65
	v131 = v91
	goto L1
L48:
	;
	goto L47
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
	var v898 int32
	_ = v898
	var v901 int32
	_ = v901
	var v908 int32
	_ = v908
	var v911 int32
	_ = v911
	var v912 int32
	_ = v912
	var v913 int32
	_ = v913
	var v927 int32
	_ = v927
	var v928 int32
	_ = v928
	var v930 int32
	_ = v930
	var v933 int32
	_ = v933
	var v934 int32
	_ = v934
	var v939 int32
	_ = v939
	var v947 int32
	_ = v947
	var v949 int32
	_ = v949
	var v951 int32
	_ = v951
	var v952 int32
	_ = v952
	var v955 int32
	_ = v955
	var v959 int32
	_ = v959
	var v966 int32
	_ = v966
	var v967 int32
	_ = v967
	var v968 int32
	_ = v968
	var v973 int32
	_ = v973
	var v975 int32
	_ = v975
	var v976 int32
	_ = v976
	var v977 int32
	_ = v977
	var v987 int32
	_ = v987
	var v988 int32
	_ = v988
	var v990 int32
	_ = v990
	var v993 int32
	_ = v993
	var v994 int32
	_ = v994
	var v999 int32
	_ = v999
	var v1006 int32
	_ = v1006
	var v1008 int32
	_ = v1008
	var v1010 int32
	_ = v1010
	var v1011 int32
	_ = v1011
	var v1013 int32
	_ = v1013
	var v1015 int32
	_ = v1015
	var v1022 int32
	_ = v1022
	var v1023 int32
	_ = v1023
	var v1024 int32
	_ = v1024
	var v1025 int32
	_ = v1025
	var v1035 int32
	_ = v1035
	var v1036 int32
	_ = v1036
	var v1038 int32
	_ = v1038
	var v1041 int32
	_ = v1041
	var v1042 int32
	_ = v1042
	var v1047 int32
	_ = v1047
	var v1054 int32
	_ = v1054
	var v1056 int32
	_ = v1056
	var v1058 int32
	_ = v1058
	var v1059 int32
	_ = v1059
	var v1061 int32
	_ = v1061
	var v1063 int32
	_ = v1063
	var v1070 int32
	_ = v1070
	var v1071 int32
	_ = v1071
	var v1074 int32
	_ = v1074
	var v1075 int32
	_ = v1075
	var v1085 int32
	_ = v1085
	var v1086 int32
	_ = v1086
	var v1088 int32
	_ = v1088
	var v1091 int32
	_ = v1091
	var v1092 int32
	_ = v1092
	var v1097 int32
	_ = v1097
	var v1104 int32
	_ = v1104
	var v1106 int32
	_ = v1106
	var v1108 int32
	_ = v1108
	var v1109 int32
	_ = v1109
	var v1111 int32
	_ = v1111
	var v1113 int32
	_ = v1113
	var v1120 int32
	_ = v1120
	var v1122 int32
	_ = v1122
	var v1123 int32
	_ = v1123
	var v1124 int32
	_ = v1124
	var v1125 int32
	_ = v1125
	var v1126 int32
	_ = v1126
	var v1128 int32
	_ = v1128
	var v1129 int32
	_ = v1129
	var v1137 int32
	_ = v1137
	var v1139 int32
	_ = v1139
	var v1141 int32
	_ = v1141
	var v1143 int32
	_ = v1143
	var v1149 int32
	_ = v1149
	var v1152 int32
	_ = v1152
	var v1161 int32
	_ = v1161
	var v1165 int32
	_ = v1165
	var v1167 int32
	_ = v1167
	var v1169 int32
	_ = v1169
	var v1172 int32
	_ = v1172
	var v1173 int32
	_ = v1173
	var v1174 int32
	_ = v1174
	var v1184 int32
	_ = v1184
	var v1185 int32
	_ = v1185
	var v1187 int32
	_ = v1187
	var v1190 int32
	_ = v1190
	var v1191 int32
	_ = v1191
	var v1196 int32
	_ = v1196
	var v1203 int32
	_ = v1203
	var v1205 int32
	_ = v1205
	var v1207 int32
	_ = v1207
	var v1208 int32
	_ = v1208
	var v1210 int32
	_ = v1210
	var v1212 int32
	_ = v1212
	var v1219 int32
	_ = v1219
	var v1220 int32
	_ = v1220
	var v1221 int32
	_ = v1221
	var v1222 int32
	_ = v1222
	var v1233 int32
	_ = v1233
	var v1234 int32
	_ = v1234
	var v1236 int32
	_ = v1236
	var v1239 int32
	_ = v1239
	var v1240 int32
	_ = v1240
	var v1245 int32
	_ = v1245
	var v1252 int32
	_ = v1252
	var v1254 int32
	_ = v1254
	var v1256 int32
	_ = v1256
	var v1257 int32
	_ = v1257
	var v1259 int32
	_ = v1259
	var v1261 int32
	_ = v1261
	var v1268 int32
	_ = v1268
	var v1269 int32
	_ = v1269
	var v1271 int32
	_ = v1271
	var v1274 int32
	_ = v1274
	var v1275 int32
	_ = v1275
	var v1276 int32
	_ = v1276
	var v1286 int32
	_ = v1286
	var v1287 int32
	_ = v1287
	var v1289 int32
	_ = v1289
	var v1292 int32
	_ = v1292
	var v1293 int32
	_ = v1293
	var v1298 int32
	_ = v1298
	var v1305 int32
	_ = v1305
	var v1307 int32
	_ = v1307
	var v1309 int32
	_ = v1309
	var v1310 int32
	_ = v1310
	var v1312 int32
	_ = v1312
	var v1314 int32
	_ = v1314
	var v1321 int32
	_ = v1321
	var v1329 int32
	_ = v1329
	var v1332 int32
	_ = v1332
	var v1333 int32
	_ = v1333
	var v1334 int32
	_ = v1334
	var v1344 int32
	_ = v1344
	var v1345 int32
	_ = v1345
	var v1347 int32
	_ = v1347
	var v1350 int32
	_ = v1350
	var v1351 int32
	_ = v1351
	var v1356 int32
	_ = v1356
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
	var v1372 int32
	_ = v1372
	var v1379 int32
	_ = v1379
	var v1382 int32
	_ = v1382
	var v1383 int32
	_ = v1383
	var v1384 int32
	_ = v1384
	var v1385 int32
	_ = v1385
	var v1386 int32
	_ = v1386
	var v1387 int32
	_ = v1387
	var v1390 int32
	_ = v1390
	var v1391 int32
	_ = v1391
	var v1395 int32
	_ = v1395
	var v1406 int32
	_ = v1406
	var v1407 int32
	_ = v1407
	var v1409 int32
	_ = v1409
	var v1413 int32
	_ = v1413
	var v1415 int32
	_ = v1415
	var v1421 int32
	_ = v1421
	var v1426 int32
	_ = v1426
	var v1427 int32
	_ = v1427
	var v1430 int32
	_ = v1430
	var v1431 int32
	_ = v1431
	var v1434 int32
	_ = v1434
	var v1435 int32
	_ = v1435
	var v1445 int32
	_ = v1445
	var v1446 int32
	_ = v1446
	var v1448 int32
	_ = v1448
	var v1451 int32
	_ = v1451
	var v1452 int32
	_ = v1452
	var v1457 int32
	_ = v1457
	var v1464 int32
	_ = v1464
	var v1466 int32
	_ = v1466
	var v1468 int32
	_ = v1468
	var v1469 int32
	_ = v1469
	var v1471 int32
	_ = v1471
	var v1473 int32
	_ = v1473
	var v1480 int32
	_ = v1480
	var v1483 int32
	_ = v1483
	var v1484 int32
	_ = v1484
	var v1493 int32
	_ = v1493
	var v1494 int32
	_ = v1494
	var v1496 int32
	_ = v1496
	var v1499 int32
	_ = v1499
	var v1500 int32
	_ = v1500
	var v1505 int32
	_ = v1505
	var v1512 int32
	_ = v1512
	var v1514 int32
	_ = v1514
	var v1516 int32
	_ = v1516
	var v1519 int32
	_ = v1519
	var v1521 int32
	_ = v1521
	var v1523 int32
	_ = v1523
	var v1530 int32
	_ = v1530
	var v1537 int32
	_ = v1537
	var v1539 int32
	_ = v1539
	var v1540 int32
	_ = v1540
	var v1541 int32
	_ = v1541
	var v1542 int32
	_ = v1542
	var v1543 int32
	_ = v1543
	var v1545 int32
	_ = v1545
	var v1546 int32
	_ = v1546
	var v1549 int32
	_ = v1549
	var v1551 int32
	_ = v1551
	var v1567 int32
	_ = v1567
	var v1578 int32
	_ = v1578
	var v1588 int32
	_ = v1588
	var v1589 int32
	_ = v1589
	var v1591 int32
	_ = v1591
	var v1594 int32
	_ = v1594
	var v1595 int32
	_ = v1595
	var v1600 int32
	_ = v1600
	var v1607 int32
	_ = v1607
	var v1609 int32
	_ = v1609
	var v1611 int32
	_ = v1611
	var v1612 int32
	_ = v1612
	var v1614 int32
	_ = v1614
	var v1616 int32
	_ = v1616
	var v1623 int32
	_ = v1623
	var v1654 int32
	_ = v1654
	v7 = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(l4))) = v7
	*(*uint8)(unsafe.Add(mBase, uint32(l5))) = uint8(v7)
	v19 = *(*int32)(unsafe.Add(mBase, uint32(l0)+112))
	if v19 == v7 {
		v1161 = v7
		v1165 = v7
		v1167 = v7
		goto L2
	} else {
		goto L3
	}
L1:
	;
	return v1654
L2:
	;
	v1169 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+317)))
	if v1169 != int32(1) {
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
	v1137 = v7
	v1139 = v7
	v1141 = v7
	v1143 = v7
	goto L6
L6:
	;
	if v1139 == int32(0) {
		v1161 = v1137
		v1165 = v1141
		v1167 = v1143
		goto L2
	} else {
		goto L313
	}
L7:
	;
	v39 = *(*int32)(unsafe.Add(mBase, uint32(v19)+12))
	v43 = *(*int32)(unsafe.Add(mBase, uint32(v39+v34<<(uint(int32(2))%32))))
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
	v1137 = v1123
	v1139 = v1124
	v1141 = v1125
	v1143 = v1126
	goto L6
L9:
	;
	v1128 = v34 + int32(1)
	v1129 = *(*int32)(unsafe.Add(mBase, uint32(v19)+4))
	if v1128 < v1129 {
		v31 = v1123
		v33 = v1124
		v34 = v1128
		v35 = v1125
		v37 = v1126
		goto L7
	} else {
		goto L312
	}
L10:
	;
	v1123 = v31
	v1124 = v1122
	v1125 = v35
	v1126 = v37
	goto L9
L11:
	;
	if v90 == int32(0) {
		v1122 = v33
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
		v1122 = v33
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
		v1122 = v33
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
		v1122 = v33
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
		v1122 = v33
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
		v1122 = v33
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
		v1123 = v43
		v1124 = v33
		v1125 = v715
		v1126 = v37
		goto L9
	} else {
		goto L193
	}
L193:
	;
	v1654 = v7
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
	v1123 = v43
	v1124 = v33
	v1125 = int32(1)
	v1126 = v37
	goto L9
L226:
	;
	goto L227
L227:
	;
	return int32(0)
L228:
	;
	v975 = *(*int32)(unsafe.Add(mBase, uint32(l1)+8))
	v976 = *(*int32)(unsafe.Add(mBase, uint32(v43)+8))
	v977 = int32(0)
	if base.B2i32(v975 == v977)|base.B2i32(v976 == v977) != 0 {
		v1022 = v977
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
	v908 = *(*int32)(unsafe.Add(mBase, uint32(v43)+20))
	if v908 != int32(4) {
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
	v897 = *(*int32)(unsafe.Add(mBase, uint32(l2)+48))
	v898 = F_create_unique_path(m, l0, l2, v897, v43)
	mBase = m.M
	v901 = m.ExcPending
	if v901 != 0 {
		goto L243
	} else {
		goto L244
	}
L243:
	;
	return int32(0)
L244:
	;
	if v898 == int32(0) {
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
	v1123 = v43
	v1124 = v33
	v1125 = int32(0)
	v1126 = int32(1)
	goto L9
L249:
	;
	v911 = *(*int32)(unsafe.Add(mBase, uint32(v43)+16))
	v912 = *(*int32)(unsafe.Add(mBase, uint32(l1)+8))
	v913 = int32(0)
	if base.B2i32(v911 == v913)|base.B2i32(v912 == v913) != 0 {
		v959 = base.B2i32(v911|v912 == v913)
		goto L251
	} else {
		goto L252
	}
L250:
	;
	if v959 == int32(0) {
		goto L228
	} else {
		goto L261
	}
L251:
	;
	goto L250
L252:
	;
	v927 = *(*int32)(unsafe.Add(mBase, uint32(v911)+4))
	v928 = *(*int32)(unsafe.Add(mBase, uint32(v912)+4))
	if v927 != v928 {
		v959 = int32(0)
		goto L251
	} else {
		goto L253
	}
L253:
	;
	v930 = int32(1)
	if v927 <= v930 {
		goto L254
	} else {
		goto L255
	}
L254:
	;
	v933 = v930
	goto L256
L255:
	;
	v933 = v927
	goto L256
L256:
	;
	v934 = int32(8)
	v939 = int32(0)
	goto L257
L257:
	;
	v947 = v939 << (uint(int32(2)) % 32)
	v949 = *(*int32)(unsafe.Add(mBase, uint32(v911+v934+v947)))
	v951 = *(*int32)(unsafe.Add(mBase, uint32(v912+v934+v947)))
	v952 = base.B2i32(v949 == v951)
	if v949 != v951 {
		v959 = v952
		goto L251
	} else {
		goto L259
	}
L258:
	;
	v959 = v952
	goto L251
L259:
	;
	v955 = v939 + int32(1)
	if v955 != v933 {
		v939 = v955
		goto L257
	} else {
		goto L260
	}
L260:
	;
	goto L258
L261:
	;
	v966 = *(*int32)(unsafe.Add(mBase, uint32(l1)+48))
	v967 = F_create_unique_path(m, l0, l1, v966, v43)
	mBase = m.M
	v968 = m.ExcPending
	if v968 != 0 {
		goto L243
	} else {
		goto L262
	}
L262:
	;
	if v967 == int32(0) {
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
	v973 = int32(1)
	v1123 = v43
	v1124 = v33
	v1125 = v973
	v1126 = v973
	goto L9
L267:
	;
	if v1022 != 0 {
		goto L280
	} else {
		goto L281
	}
L268:
	;
	goto L267
L269:
	;
	v987 = *(*int32)(unsafe.Add(mBase, uint32(v975)+4))
	v988 = *(*int32)(unsafe.Add(mBase, uint32(v976)+4))
	if v987 < v988 {
		goto L270
	} else {
		goto L271
	}
L270:
	;
	v990 = v987
	goto L272
L271:
	;
	v990 = v988
	goto L272
L272:
	;
	if v990 <= int32(1) {
		goto L273
	} else {
		goto L274
	}
L273:
	;
	v993 = int32(1)
	goto L275
L274:
	;
	v993 = v990
	goto L275
L275:
	;
	v994 = int32(8)
	v999 = int32(0)
	goto L276
L276:
	;
	v1006 = v999 << (uint(int32(2)) % 32)
	v1008 = *(*int32)(unsafe.Add(mBase, uint32(v976+v994+v1006)))
	v1010 = *(*int32)(unsafe.Add(mBase, uint32(v975+v994+v1006)))
	v1011 = v1008 & v1010
	v1013 = base.B2i32(v1011 != int32(0))
	if v1011 != 0 {
		v1022 = v1013
		goto L268
	} else {
		goto L278
	}
L277:
	;
	v1022 = v1013
	goto L268
L278:
	;
	v1015 = v999 + int32(1)
	if v1015 != v993 {
		v999 = v1015
		goto L276
	} else {
		goto L279
	}
L279:
	;
	goto L277
L280:
	;
	v1023 = *(*int32)(unsafe.Add(mBase, uint32(l2)+8))
	v1024 = *(*int32)(unsafe.Add(mBase, uint32(v43)+8))
	v1025 = int32(0)
	if base.B2i32(v1023 == v1025)|base.B2i32(v1024 == v1025) != 0 {
		v1070 = v1025
		goto L284
	} else {
		goto L285
	}
L281:
	;
	goto L282
L282:
	;
	v1071 = *(*int32)(unsafe.Add(mBase, uint32(v43)+20))
	if v1071 != int32(1) {
		v1654 = v7
		goto L1
	} else {
		goto L297
	}
L283:
	;
	if v1070 != 0 {
		v1122 = v33
		goto L10
	} else {
		goto L296
	}
L284:
	;
	goto L283
L285:
	;
	v1035 = *(*int32)(unsafe.Add(mBase, uint32(v1023)+4))
	v1036 = *(*int32)(unsafe.Add(mBase, uint32(v1024)+4))
	if v1035 < v1036 {
		goto L286
	} else {
		goto L287
	}
L286:
	;
	v1038 = v1035
	goto L288
L287:
	;
	v1038 = v1036
	goto L288
L288:
	;
	if v1038 <= int32(1) {
		goto L289
	} else {
		goto L290
	}
L289:
	;
	v1041 = int32(1)
	goto L291
L290:
	;
	v1041 = v1038
	goto L291
L291:
	;
	v1042 = int32(8)
	v1047 = int32(0)
	goto L292
L292:
	;
	v1054 = v1047 << (uint(int32(2)) % 32)
	v1056 = *(*int32)(unsafe.Add(mBase, uint32(v1024+v1042+v1054)))
	v1058 = *(*int32)(unsafe.Add(mBase, uint32(v1023+v1042+v1054)))
	v1059 = v1056 & v1058
	v1061 = base.B2i32(v1059 != int32(0))
	if v1059 != 0 {
		v1070 = v1061
		goto L284
	} else {
		goto L294
	}
L293:
	;
	v1070 = v1061
	goto L284
L294:
	;
	v1063 = v1047 + int32(1)
	if v1063 != v1041 {
		v1047 = v1063
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
	v1074 = *(*int32)(unsafe.Add(mBase, uint32(v43)+4))
	v1075 = int32(0)
	if base.B2i32(l3 == v1075)|base.B2i32(v1074 == v1075) != 0 {
		v1120 = v1075
		goto L299
	} else {
		goto L300
	}
L298:
	;
	if v1120 != 0 {
		v1654 = v7
		goto L1
	} else {
		goto L311
	}
L299:
	;
	goto L298
L300:
	;
	v1085 = *(*int32)(unsafe.Add(mBase, uint32(l3)+4))
	v1086 = *(*int32)(unsafe.Add(mBase, uint32(v1074)+4))
	if v1085 < v1086 {
		goto L301
	} else {
		goto L302
	}
L301:
	;
	v1088 = v1085
	goto L303
L302:
	;
	v1088 = v1086
	goto L303
L303:
	;
	if v1088 <= int32(1) {
		goto L304
	} else {
		goto L305
	}
L304:
	;
	v1091 = int32(1)
	goto L306
L305:
	;
	v1091 = v1088
	goto L306
L306:
	;
	v1092 = int32(8)
	v1097 = int32(0)
	goto L307
L307:
	;
	v1104 = v1097 << (uint(int32(2)) % 32)
	v1106 = *(*int32)(unsafe.Add(mBase, uint32(v1074+v1092+v1104)))
	v1108 = *(*int32)(unsafe.Add(mBase, uint32(l3+v1092+v1104)))
	v1109 = v1106 & v1108
	v1111 = base.B2i32(v1109 != int32(0))
	if v1109 != 0 {
		v1120 = v1111
		goto L299
	} else {
		goto L309
	}
L308:
	;
	v1120 = v1111
	goto L299
L309:
	;
	v1113 = v1097 + int32(1)
	if v1113 != v1091 {
		v1097 = v1113
		goto L307
	} else {
		goto L310
	}
L310:
	;
	goto L308
L311:
	;
	v1122 = int32(1)
	goto L10
L312:
	;
	goto L8
L313:
	;
	if v1137 == int32(0) {
		v1654 = v7
		goto L1
	} else {
		goto L314
	}
L314:
	;
	v1149 = *(*int32)(unsafe.Add(mBase, uint32(v1137)+20))
	if v1149 != int32(1) {
		v1654 = v7
		goto L1
	} else {
		goto L315
	}
L315:
	;
	v1152 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1137)+44)))
	if v1152 != int32(1) {
		v1654 = v7
		goto L1
	} else {
		goto L316
	}
L316:
	;
	v1161 = v1137
	v1165 = v1141
	v1167 = v1143
	goto L2
L317:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l4))) = v1161
	*(*uint8)(unsafe.Add(mBase, uint32(l5))) = uint8(v1165)
	v1654 = int32(1)
	goto L1
L318:
	;
	v1172 = *(*int32)(unsafe.Add(mBase, uint32(l1)+8))
	v1173 = *(*int32)(unsafe.Add(mBase, uint32(l2)+64))
	v1174 = int32(0)
	if base.B2i32(v1172 == v1174)|base.B2i32(v1173 == v1174) != 0 {
		v1219 = v1174
		goto L320
	} else {
		goto L321
	}
L319:
	;
	v1220 = int32(0)
	v1221 = *(*int32)(unsafe.Add(mBase, uint32(l2)+8))
	v1222 = *(*int32)(unsafe.Add(mBase, uint32(l1)+64))
	if base.B2i32(v1221 == v1220)|base.B2i32(v1222 == v1220) != 0 {
		v1268 = v1220
		goto L333
	} else {
		goto L334
	}
L320:
	;
	goto L319
L321:
	;
	v1184 = *(*int32)(unsafe.Add(mBase, uint32(v1172)+4))
	v1185 = *(*int32)(unsafe.Add(mBase, uint32(v1173)+4))
	if v1184 < v1185 {
		goto L322
	} else {
		goto L323
	}
L322:
	;
	v1187 = v1184
	goto L324
L323:
	;
	v1187 = v1185
	goto L324
L324:
	;
	if v1187 <= int32(1) {
		goto L325
	} else {
		goto L326
	}
L325:
	;
	v1190 = int32(1)
	goto L327
L326:
	;
	v1190 = v1187
	goto L327
L327:
	;
	v1191 = int32(8)
	v1196 = int32(0)
	goto L328
L328:
	;
	v1203 = v1196 << (uint(int32(2)) % 32)
	v1205 = *(*int32)(unsafe.Add(mBase, uint32(v1173+v1191+v1203)))
	v1207 = *(*int32)(unsafe.Add(mBase, uint32(v1172+v1191+v1203)))
	v1208 = v1205 & v1207
	v1210 = base.B2i32(v1208 != int32(0))
	if v1208 != 0 {
		v1219 = v1210
		goto L320
	} else {
		goto L330
	}
L329:
	;
	v1219 = v1210
	goto L320
L330:
	;
	v1212 = v1196 + int32(1)
	if v1212 != v1190 {
		v1196 = v1212
		goto L328
	} else {
		goto L331
	}
L331:
	;
	goto L329
L332:
	;
	if v1268 != 0 {
		goto L345
	} else {
		goto L346
	}
L333:
	;
	goto L332
L334:
	;
	v1233 = *(*int32)(unsafe.Add(mBase, uint32(v1221)+4))
	v1234 = *(*int32)(unsafe.Add(mBase, uint32(v1222)+4))
	if v1233 < v1234 {
		goto L335
	} else {
		goto L336
	}
L335:
	;
	v1236 = v1233
	goto L337
L336:
	;
	v1236 = v1234
	goto L337
L337:
	;
	if v1236 <= int32(1) {
		goto L338
	} else {
		goto L339
	}
L338:
	;
	v1239 = int32(1)
	goto L340
L339:
	;
	v1239 = v1236
	goto L340
L340:
	;
	v1240 = int32(8)
	v1245 = int32(0)
	goto L341
L341:
	;
	v1252 = v1245 << (uint(int32(2)) % 32)
	v1254 = *(*int32)(unsafe.Add(mBase, uint32(v1222+v1240+v1252)))
	v1256 = *(*int32)(unsafe.Add(mBase, uint32(v1221+v1240+v1252)))
	v1257 = v1254 & v1256
	v1259 = base.B2i32(v1257 != int32(0))
	if v1257 != 0 {
		v1268 = v1259
		goto L333
	} else {
		goto L343
	}
L342:
	;
	v1268 = v1259
	goto L333
L343:
	;
	v1261 = v1245 + int32(1)
	if v1261 != v1239 {
		v1245 = v1261
		goto L341
	} else {
		goto L344
	}
L344:
	;
	goto L342
L345:
	;
	v1269 = v1219
	goto L347
L346:
	;
	v1269 = v1220
	goto L347
L347:
	;
	if v1269 != 0 {
		v1654 = v7
		goto L1
	} else {
		goto L348
	}
L348:
	;
	if v1219 != 0 {
		goto L350
	} else {
		goto L351
	}
L349:
	;
	v1382 = *(*int32)(unsafe.Add(mBase, uint32(l1)+64))
	v1383 = *(*int32)(unsafe.Add(mBase, uint32(l2)+64))
	v1384 = F_bms_union(m, v1382, v1383)
	mBase = m.M
	v1385 = m.ExcPending
	if v1385 != 0 {
		goto L243
	} else {
		goto L392
	}
L350:
	;
	if v1161 != 0 {
		goto L353
	} else {
		goto L354
	}
L351:
	;
	goto L352
L352:
	;
	if v1268 == int32(0) {
		goto L349
	} else {
		goto L372
	}
L353:
	;
	if v1165|v1167 != 0 {
		v1654 = v7
		goto L1
	} else {
		goto L356
	}
L354:
	;
	goto L355
L355:
	;
	v1274 = *(*int32)(unsafe.Add(mBase, uint32(l1)+8))
	v1275 = *(*int32)(unsafe.Add(mBase, uint32(l2)+60))
	v1276 = int32(0)
	if base.B2i32(v1274 == v1276)|base.B2i32(v1275 == v1276) != 0 {
		v1321 = v1276
		goto L359
	} else {
		goto L360
	}
L356:
	;
	v1271 = *(*int32)(unsafe.Add(mBase, uint32(v1161)+20))
	if v1271 == int32(2) {
		v1654 = v7
		goto L1
	} else {
		goto L357
	}
L357:
	;
	goto L355
L358:
	;
	if v1321 != 0 {
		goto L349
	} else {
		goto L371
	}
L359:
	;
	goto L358
L360:
	;
	v1286 = *(*int32)(unsafe.Add(mBase, uint32(v1274)+4))
	v1287 = *(*int32)(unsafe.Add(mBase, uint32(v1275)+4))
	if v1286 < v1287 {
		goto L361
	} else {
		goto L362
	}
L361:
	;
	v1289 = v1286
	goto L363
L362:
	;
	v1289 = v1287
	goto L363
L363:
	;
	if v1289 <= int32(1) {
		goto L364
	} else {
		goto L365
	}
L364:
	;
	v1292 = int32(1)
	goto L366
L365:
	;
	v1292 = v1289
	goto L366
L366:
	;
	v1293 = int32(8)
	v1298 = int32(0)
	goto L367
L367:
	;
	v1305 = v1298 << (uint(int32(2)) % 32)
	v1307 = *(*int32)(unsafe.Add(mBase, uint32(v1275+v1293+v1305)))
	v1309 = *(*int32)(unsafe.Add(mBase, uint32(v1274+v1293+v1305)))
	v1310 = v1307 & v1309
	v1312 = base.B2i32(v1310 != int32(0))
	if v1310 != 0 {
		v1321 = v1312
		goto L359
	} else {
		goto L369
	}
L368:
	;
	v1321 = v1312
	goto L359
L369:
	;
	v1314 = v1298 + int32(1)
	if v1314 != v1292 {
		v1298 = v1314
		goto L367
	} else {
		goto L370
	}
L370:
	;
	goto L368
L371:
	;
	v1654 = v7
	goto L1
L372:
	;
	if v1161 != 0 {
		goto L373
	} else {
		goto L374
	}
L373:
	;
	if (v1165^int32(-1)|v1167)&int32(1) != 0 {
		v1654 = v7
		goto L1
	} else {
		goto L376
	}
L374:
	;
	goto L375
L375:
	;
	v1332 = *(*int32)(unsafe.Add(mBase, uint32(l2)+8))
	v1333 = *(*int32)(unsafe.Add(mBase, uint32(l1)+60))
	v1334 = int32(0)
	if base.B2i32(v1332 == v1334)|base.B2i32(v1333 == v1334) != 0 {
		v1379 = v1334
		goto L379
	} else {
		goto L380
	}
L376:
	;
	v1329 = *(*int32)(unsafe.Add(mBase, uint32(v1161)+20))
	if v1329 == int32(2) {
		v1654 = v7
		goto L1
	} else {
		goto L377
	}
L377:
	;
	goto L375
L378:
	;
	if v1379 == int32(0) {
		v1654 = v7
		goto L1
	} else {
		goto L391
	}
L379:
	;
	goto L378
L380:
	;
	v1344 = *(*int32)(unsafe.Add(mBase, uint32(v1332)+4))
	v1345 = *(*int32)(unsafe.Add(mBase, uint32(v1333)+4))
	if v1344 < v1345 {
		goto L381
	} else {
		goto L382
	}
L381:
	;
	v1347 = v1344
	goto L383
L382:
	;
	v1347 = v1345
	goto L383
L383:
	;
	if v1347 <= int32(1) {
		goto L384
	} else {
		goto L385
	}
L384:
	;
	v1350 = int32(1)
	goto L386
L385:
	;
	v1350 = v1347
	goto L386
L386:
	;
	v1351 = int32(8)
	v1356 = int32(0)
	goto L387
L387:
	;
	v1363 = v1356 << (uint(int32(2)) % 32)
	v1365 = *(*int32)(unsafe.Add(mBase, uint32(v1333+v1351+v1363)))
	v1367 = *(*int32)(unsafe.Add(mBase, uint32(v1332+v1351+v1363)))
	v1368 = v1365 & v1367
	v1370 = base.B2i32(v1368 != int32(0))
	if v1368 != 0 {
		v1379 = v1370
		goto L379
	} else {
		goto L389
	}
L388:
	;
	v1379 = v1370
	goto L379
L389:
	;
	v1372 = v1356 + int32(1)
	if v1372 != v1350 {
		v1356 = v1372
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
	v1386 = F_bms_del_members(m, v1384, l3)
	mBase = m.M
	v1387 = m.ExcPending
	if v1387 != 0 {
		goto L243
	} else {
		goto L393
	}
L393:
	;
	if v1386 == int32(0) {
		goto L317
	} else {
		goto L394
	}
L394:
	;
	v1390 = F_bms_copy(m, l3)
	mBase = m.M
	v1391 = m.ExcPending
	if v1391 != 0 {
		goto L243
	} else {
		goto L395
	}
L395:
	;
	v1395 = v1390
	goto L396
L396:
	;
	v1406 = *(*int32)(unsafe.Add(mBase, uint32(l0)+112))
	if v1406 != 0 {
		goto L398
	} else {
		goto L399
	}
L397:
	;
	v1578 = int32(0)
	if base.B2i32(v1567 == v1578)|base.B2i32(v1386 == v1578) != 0 {
		v1623 = v1578
		goto L441
	} else {
		goto L442
	}
L398:
	;
	v1407 = int32(0)
	v1409 = *(*int32)(unsafe.Add(mBase, uint32(v1406)+4))
	if v1407 < v1409 {
		goto L401
	} else {
		goto L402
	}
L399:
	;
	v1567 = v1395
	goto L400
L400:
	;
	goto L397
L401:
	;
	v1413 = v1407
	v1415 = v1395
	v1421 = v1407
	goto L404
L402:
	;
	v1549 = v1407
	v1551 = v1395
	goto L403
L403:
	;
	if v1549&int32(1) != 0 {
		v1395 = v1551
		goto L396
	} else {
		goto L439
	}
L404:
	;
	v1426 = *(*int32)(unsafe.Add(mBase, uint32(v1406)+12))
	v1427 = int32(2)
	v1430 = *(*int32)(unsafe.Add(mBase, uint32(v1426+v1421<<(uint(v1427)%32))))
	v1431 = *(*int32)(unsafe.Add(mBase, uint32(v1430)+20))
	if v1431 == v1427 {
		v1542 = v1413
		v1543 = v1415
		goto L406
	} else {
		goto L407
	}
L405:
	;
	v1549 = v1542
	v1551 = v1543
	goto L403
L406:
	;
	v1545 = v1421 + int32(1)
	v1546 = *(*int32)(unsafe.Add(mBase, uint32(v1406)+4))
	if v1545 < v1546 {
		v1413 = v1542
		v1415 = v1543
		v1421 = v1545
		goto L404
	} else {
		goto L438
	}
L407:
	;
	v1434 = *(*int32)(unsafe.Add(mBase, uint32(v1430)+4))
	v1435 = int32(0)
	if base.B2i32(v1434 == v1435)|base.B2i32(v1415 == v1435) != 0 {
		v1480 = v1435
		goto L409
	} else {
		goto L410
	}
L408:
	;
	if v1480 == int32(0) {
		v1542 = v1413
		v1543 = v1415
		goto L406
	} else {
		goto L421
	}
L409:
	;
	goto L408
L410:
	;
	v1445 = *(*int32)(unsafe.Add(mBase, uint32(v1434)+4))
	v1446 = *(*int32)(unsafe.Add(mBase, uint32(v1415)+4))
	if v1445 < v1446 {
		goto L411
	} else {
		goto L412
	}
L411:
	;
	v1448 = v1445
	goto L413
L412:
	;
	v1448 = v1446
	goto L413
L413:
	;
	if v1448 <= int32(1) {
		goto L414
	} else {
		goto L415
	}
L414:
	;
	v1451 = int32(1)
	goto L416
L415:
	;
	v1451 = v1448
	goto L416
L416:
	;
	v1452 = int32(8)
	v1457 = int32(0)
	goto L417
L417:
	;
	v1464 = v1457 << (uint(int32(2)) % 32)
	v1466 = *(*int32)(unsafe.Add(mBase, uint32(v1415+v1452+v1464)))
	v1468 = *(*int32)(unsafe.Add(mBase, uint32(v1434+v1452+v1464)))
	v1469 = v1466 & v1468
	v1471 = base.B2i32(v1469 != int32(0))
	if v1469 != 0 {
		v1480 = v1471
		goto L409
	} else {
		goto L419
	}
L418:
	;
	v1480 = v1471
	goto L409
L419:
	;
	v1473 = v1457 + int32(1)
	if v1473 != v1451 {
		v1457 = v1473
		goto L417
	} else {
		goto L420
	}
L420:
	;
	goto L418
L421:
	;
	v1483 = *(*int32)(unsafe.Add(mBase, uint32(v1430)+8))
	v1484 = int32(0)
	if v1483 == v1484 {
		goto L423
	} else {
		goto L424
	}
L422:
	;
	if v1537 != 0 {
		v1542 = v1413
		v1543 = v1415
		goto L406
	} else {
		goto L436
	}
L423:
	;
	v1537 = int32(1)
	goto L422
L424:
	;
	goto L425
L425:
	;
	if v1415 == int32(0) {
		v1530 = v1484
		goto L426
	} else {
		goto L427
	}
L426:
	;
	v1537 = v1530
	goto L422
L427:
	;
	v1493 = *(*int32)(unsafe.Add(mBase, uint32(v1483)+4))
	v1494 = *(*int32)(unsafe.Add(mBase, uint32(v1415)+4))
	if v1494 < v1493 {
		v1530 = v1484
		goto L426
	} else {
		goto L428
	}
L428:
	;
	v1496 = int32(1)
	if v1493 <= v1496 {
		goto L429
	} else {
		goto L430
	}
L429:
	;
	v1499 = v1496
	goto L431
L430:
	;
	v1499 = v1493
	goto L431
L431:
	;
	v1500 = int32(8)
	v1505 = int32(0)
	goto L432
L432:
	;
	v1512 = v1505 << (uint(int32(2)) % 32)
	v1514 = *(*int32)(unsafe.Add(mBase, uint32(v1483+v1500+v1512)))
	v1516 = *(*int32)(unsafe.Add(mBase, uint32(v1415+v1500+v1512)))
	v1519 = v1514 & (v1516 ^ int32(-1))
	v1521 = base.B2i32(v1519 == int32(0))
	if v1519 != 0 {
		v1530 = v1521
		goto L426
	} else {
		goto L434
	}
L433:
	;
	v1530 = v1521
	goto L426
L434:
	;
	v1523 = v1505 + int32(1)
	if v1523 != v1499 {
		v1505 = v1523
		goto L432
	} else {
		goto L435
	}
L435:
	;
	goto L433
L436:
	;
	v1539 = *(*int32)(unsafe.Add(mBase, uint32(v1430)+8))
	v1540 = F_bms_add_members(m, v1415, v1539)
	mBase = m.M
	v1541 = m.ExcPending
	if v1541 != 0 {
		goto L243
	} else {
		goto L437
	}
L437:
	;
	v1542 = int32(1)
	v1543 = v1540
	goto L406
L438:
	;
	goto L405
L439:
	;
	v1567 = v1551
	goto L400
L440:
	;
	if v1623 != 0 {
		v1654 = v7
		goto L1
	} else {
		goto L453
	}
L441:
	;
	goto L440
L442:
	;
	v1588 = *(*int32)(unsafe.Add(mBase, uint32(v1567)+4))
	v1589 = *(*int32)(unsafe.Add(mBase, uint32(v1386)+4))
	if v1588 < v1589 {
		goto L443
	} else {
		goto L444
	}
L443:
	;
	v1591 = v1588
	goto L445
L444:
	;
	v1591 = v1589
	goto L445
L445:
	;
	if v1591 <= int32(1) {
		goto L446
	} else {
		goto L447
	}
L446:
	;
	v1594 = int32(1)
	goto L448
L447:
	;
	v1594 = v1591
	goto L448
L448:
	;
	v1595 = int32(8)
	v1600 = int32(0)
	goto L449
L449:
	;
	v1607 = v1600 << (uint(int32(2)) % 32)
	v1609 = *(*int32)(unsafe.Add(mBase, uint32(v1386+v1595+v1607)))
	v1611 = *(*int32)(unsafe.Add(mBase, uint32(v1567+v1595+v1607)))
	v1612 = v1609 & v1611
	v1614 = base.B2i32(v1612 != int32(0))
	if v1612 != 0 {
		v1623 = v1614
		goto L441
	} else {
		goto L451
	}
L450:
	;
	v1623 = v1614
	goto L441
L451:
	;
	v1616 = v1600 + int32(1)
	if v1616 != v1594 {
		v1600 = v1616
		goto L449
	} else {
		goto L452
	}
L452:
	;
	goto L450
L453:
	;
	goto L317
}
