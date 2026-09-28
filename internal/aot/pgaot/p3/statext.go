package p3

import (
	base "github.com/shibukawa/pgmem/internal/aot/pgaot/base"
	"unsafe"
)

func F_statext_is_compatible_clause_internal(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32) int32 {
	mBase := m.M
	_ = mBase
	var v6 int32
	_ = v6
	var v10 int32
	_ = v10
	var v11 int32
	_ = v11
	var v16 int32
	_ = v16
	var v22 int32
	_ = v22
	var v23 int32
	_ = v23
	var v24 int32
	_ = v24
	var v25 int32
	_ = v25
	var v28 int32
	_ = v28
	var v30 int32
	_ = v30
	var v31 int32
	_ = v31
	var v34 int32
	_ = v34
	var v35 int32
	_ = v35
	var v38 int32
	_ = v38
	var v40 int32
	_ = v40
	var v43 int32
	_ = v43
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
	var v57 int32
	_ = v57
	var v58 int32
	_ = v58
	var v59 int32
	_ = v59
	var v60 int32
	_ = v60
	var v63 int32
	_ = v63
	var v66 int32
	_ = v66
	var v67 int32
	_ = v67
	var v68 int32
	_ = v68
	var v69 int32
	_ = v69
	var v79 int32
	_ = v79
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
	var v92 int32
	_ = v92
	var v93 int32
	_ = v93
	var v94 int32
	_ = v94
	var v96 int32
	_ = v96
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
	var v108 int32
	_ = v108
	var v109 int32
	_ = v109
	var v110 int32
	_ = v110
	var v113 int32
	_ = v113
	var v114 int32
	_ = v114
	var v115 int32
	_ = v115
	var v118 int32
	_ = v118
	var v119 int32
	_ = v119
	var v120 int32
	_ = v120
	var v130 int32
	_ = v130
	var v133 int32
	_ = v133
	var v134 int32
	_ = v134
	var v135 int32
	_ = v135
	var v136 int32
	_ = v136
	var v137 int32
	_ = v137
	var v139 int32
	_ = v139
	var v140 int32
	_ = v140
	var v143 int32
	_ = v143
	var v144 int32
	_ = v144
	var v145 int32
	_ = v145
	var v147 int32
	_ = v147
	var v150 int32
	_ = v150
	var v151 int32
	_ = v151
	var v154 int32
	_ = v154
	var v155 int32
	_ = v155
	var v158 int32
	_ = v158
	var v167 int32
	_ = v167
	var v171 int32
	_ = v171
	var v172 int32
	_ = v172
	var v173 int32
	_ = v173
	var v177 int32
	_ = v177
	var v178 int32
	_ = v178
	var v180 int32
	_ = v180
	var v181 int32
	_ = v181
	var v182 int32
	_ = v182
	var v185 int32
	_ = v185
	var v186 int32
	_ = v186
	var v187 int32
	_ = v187
	var v189 int32
	_ = v189
	var v190 int32
	_ = v190
	var v191 int32
	_ = v191
	var v205 int32
	_ = v205
	v6 = int32(0)
	v10 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v11 = l0
	v16 = v10
	goto L4
L1:
	;
	return v205
L2:
	;
	v205 = int32(1)
	goto L1
L3:
	;
	v189 = *(*int32)(unsafe.Add(mBase, uint32(l3)))
	v190 = F_lappend(m, v189, v25)
	mBase = m.M
	v191 = m.ExcPending
	if v191 != 0 {
		goto L18
	} else {
		goto L70
	}
L4:
	;
	if v16 != int32(27) {
		goto L7
	} else {
		goto L8
	}
L5:
	;
	v185 = *(*int32)(unsafe.Add(mBase, uint32(l3)))
	v186 = F_lappend(m, v185, v181)
	mBase = m.M
	v187 = m.ExcPending
	if v187 != 0 {
		goto L18
	} else {
		goto L69
	}
L6:
	;
	switch v24 - int32(6) {
	case 0:
		goto L14
	default:
		goto L3
	case 11:
		goto L13
	case 14:
		goto L12
	case 15:
		goto L11
	case 46:
		goto L10
	}
L7:
	;
	v24 = v16
	v25 = v11
	goto L6
L8:
	;
	goto L9
L9:
	;
	v22 = *(*int32)(unsafe.Add(mBase, uint32(v11)+4))
	v23 = *(*int32)(unsafe.Add(mBase, uint32(v22)))
	v24 = v23
	v25 = v22
	goto L6
L10:
	;
	v180 = int32(6)
	v181 = *(*int32)(unsafe.Add(mBase, uint32(v25)+4))
	v182 = *(*int32)(unsafe.Add(mBase, uint32(v181)))
	if v182 == v180 {
		v11 = v181
		v16 = v180
		goto L4
	} else {
		goto L68
	}
L11:
	;
	v147 = *(*int32)(unsafe.Add(mBase, uint32(v25)+4))
	if base.Ui32(int32(2)) < base.Ui32(v147) {
		goto L3
	} else {
		goto L60
	}
L12:
	;
	v96 = *(*int32)(unsafe.Add(mBase, uint32(v25)+28))
	if v96 == int32(0) {
		v205 = v6
		goto L1
	} else {
		goto L42
	}
L13:
	;
	v40 = *(*int32)(unsafe.Add(mBase, uint32(v25)+28))
	if v40 == int32(0) {
		v205 = v6
		goto L1
	} else {
		goto L20
	}
L14:
	;
	v28 = *(*int32)(unsafe.Add(mBase, uint32(v25)+4))
	if v28 != l1 {
		v205 = v6
		goto L1
	} else {
		goto L15
	}
L15:
	;
	v30 = *(*int32)(unsafe.Add(mBase, uint32(v25)+28))
	if v30 != 0 {
		v205 = v6
		goto L1
	} else {
		goto L16
	}
L16:
	;
	v31 = int32(*(*int16)(unsafe.Add(mBase, uint32(v25)+8)))
	if v31 <= int32(0) {
		v205 = v6
		goto L1
	} else {
		goto L17
	}
L17:
	;
	v34 = *(*int32)(unsafe.Add(mBase, uint32(l2)))
	v35 = F_bms_add_member(m, v34, v31)
	mBase = m.M
	v38 = m.ExcPending
	if v38 != 0 {
		goto L18
	} else {
		goto L19
	}
L18:
	;
	return int32(0)
L19:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l2))) = v35
	goto L2
L20:
	;
	v43 = *(*int32)(unsafe.Add(mBase, uint32(v40)+4))
	if v43 != int32(2) {
		v205 = v6
		goto L1
	} else {
		goto L21
	}
L21:
	;
	v46 = *(*int32)(unsafe.Add(mBase, uint32(v40)+12))
	v47 = *(*int32)(unsafe.Add(mBase, uint32(v46)+4))
	v48 = *(*int32)(unsafe.Add(mBase, uint32(v46)))
	v49 = *(*int32)(unsafe.Add(mBase, uint32(v48)))
	if v49 == int32(27) {
		goto L22
	} else {
		goto L23
	}
L22:
	;
	v52 = *(*int32)(unsafe.Add(mBase, uint32(v48)+4))
	v53 = v52
	goto L24
L23:
	;
	v53 = v48
	goto L24
L24:
	;
	v54 = *(*int32)(unsafe.Add(mBase, uint32(v47)))
	if v54 == int32(27) {
		goto L26
	} else {
		goto L27
	}
L25:
	;
	v67 = *(*int32)(unsafe.Add(mBase, uint32(v25)+4))
	v68 = F_get_oprrest(m, v67)
	mBase = m.M
	v69 = m.ExcPending
	if v69 != 0 {
		goto L18
	} else {
		goto L33
	}
L26:
	;
	v57 = *(*int32)(unsafe.Add(mBase, uint32(v47)+4))
	v58 = *(*int32)(unsafe.Add(mBase, uint32(v57)))
	v59 = v57
	v60 = v58
	goto L28
L27:
	;
	v59 = v47
	v60 = v54
	goto L28
L28:
	;
	if v60 == int32(7) {
		goto L29
	} else {
		goto L30
	}
L29:
	;
	v66 = v53
	goto L25
L30:
	;
	goto L31
L31:
	;
	v63 = *(*int32)(unsafe.Add(mBase, uint32(v53)))
	if v63 != int32(7) {
		v205 = v6
		goto L1
	} else {
		goto L32
	}
L32:
	;
	v66 = v59
	goto L25
L33:
	;
	if base.B2i32(base.Ui32(int32(4)) <= base.Ui32(v68-int32(101)))&base.B2i32(base.Ui32(int32(1)) < base.Ui32(v68-int32(336))) != 0 {
		v205 = v6
		goto L1
	} else {
		goto L34
	}
L34:
	;
	v79 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l4))))
	if v79 == int32(1) {
		goto L35
	} else {
		goto L36
	}
L35:
	;
	v82 = *(*int32)(unsafe.Add(mBase, uint32(v25)+4))
	v83 = F_get_opcode(m, v82)
	mBase = m.M
	v84 = m.ExcPending
	if v84 != 0 {
		goto L18
	} else {
		goto L38
	}
L36:
	;
	goto L37
L37:
	;
	v88 = int32(6)
	v89 = *(*int32)(unsafe.Add(mBase, uint32(v66)))
	if v89 == v88 {
		v11 = v66
		v16 = v88
		goto L4
	} else {
		goto L40
	}
L38:
	;
	v85 = F_get_func_leakproof(m, v83)
	mBase = m.M
	v86 = m.ExcPending
	if v86 != 0 {
		goto L18
	} else {
		goto L39
	}
L39:
	;
	*(*uint8)(unsafe.Add(mBase, uint32(l4))) = uint8(v85)
	goto L37
L40:
	;
	v92 = *(*int32)(unsafe.Add(mBase, uint32(l3)))
	v93 = F_lappend(m, v92, v66)
	mBase = m.M
	v94 = m.ExcPending
	if v94 != 0 {
		goto L18
	} else {
		goto L41
	}
L41:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l3))) = v93
	goto L2
L42:
	;
	v99 = *(*int32)(unsafe.Add(mBase, uint32(v96)+4))
	if v99 != int32(2) {
		v205 = v6
		goto L1
	} else {
		goto L43
	}
L43:
	;
	v102 = *(*int32)(unsafe.Add(mBase, uint32(v96)+12))
	v103 = *(*int32)(unsafe.Add(mBase, uint32(v102)+4))
	v104 = *(*int32)(unsafe.Add(mBase, uint32(v102)))
	v105 = *(*int32)(unsafe.Add(mBase, uint32(v104)))
	if v105 == int32(27) {
		goto L44
	} else {
		goto L45
	}
L44:
	;
	v108 = *(*int32)(unsafe.Add(mBase, uint32(v104)+4))
	v109 = v108
	goto L46
L45:
	;
	v109 = v104
	goto L46
L46:
	;
	v110 = *(*int32)(unsafe.Add(mBase, uint32(v103)))
	if v110 == int32(27) {
		goto L47
	} else {
		goto L48
	}
L47:
	;
	v113 = *(*int32)(unsafe.Add(mBase, uint32(v103)+4))
	v114 = *(*int32)(unsafe.Add(mBase, uint32(v113)))
	v115 = v114
	goto L49
L48:
	;
	v115 = v110
	goto L49
L49:
	;
	if v115 != int32(7) {
		v205 = v6
		goto L1
	} else {
		goto L50
	}
L50:
	;
	v118 = *(*int32)(unsafe.Add(mBase, uint32(v25)+4))
	v119 = F_get_oprrest(m, v118)
	mBase = m.M
	v120 = m.ExcPending
	if v120 != 0 {
		goto L18
	} else {
		goto L51
	}
L51:
	;
	if base.B2i32(base.Ui32(int32(4)) <= base.Ui32(v119-int32(101)))&base.B2i32(base.Ui32(int32(1)) < base.Ui32(v119-int32(336))) != 0 {
		v205 = v6
		goto L1
	} else {
		goto L52
	}
L52:
	;
	v130 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l4))))
	if v130 == int32(1) {
		goto L53
	} else {
		goto L54
	}
L53:
	;
	v133 = *(*int32)(unsafe.Add(mBase, uint32(v25)+4))
	v134 = F_get_opcode(m, v133)
	mBase = m.M
	v135 = m.ExcPending
	if v135 != 0 {
		goto L18
	} else {
		goto L56
	}
L54:
	;
	goto L55
L55:
	;
	v139 = int32(6)
	v140 = *(*int32)(unsafe.Add(mBase, uint32(v109)))
	if v140 == v139 {
		v11 = v109
		v16 = v139
		goto L4
	} else {
		goto L58
	}
L56:
	;
	v136 = F_get_func_leakproof(m, v134)
	mBase = m.M
	v137 = m.ExcPending
	if v137 != 0 {
		goto L18
	} else {
		goto L57
	}
L57:
	;
	*(*uint8)(unsafe.Add(mBase, uint32(l4))) = uint8(v136)
	goto L55
L58:
	;
	v143 = *(*int32)(unsafe.Add(mBase, uint32(l3)))
	v144 = F_lappend(m, v143, v109)
	mBase = m.M
	v145 = m.ExcPending
	if v145 != 0 {
		goto L18
	} else {
		goto L59
	}
L59:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l3))) = v144
	goto L2
L60:
	;
	v150 = int32(1)
	v151 = *(*int32)(unsafe.Add(mBase, uint32(v25)+8))
	if v151 == int32(0) {
		v205 = v150
		goto L1
	} else {
		goto L61
	}
L61:
	;
	v154 = int32(0)
	v155 = *(*int32)(unsafe.Add(mBase, uint32(v151)+4))
	if v155 <= v154 {
		v205 = v150
		goto L1
	} else {
		goto L62
	}
L62:
	;
	v158 = v154
	goto L63
L63:
	;
	v167 = *(*int32)(unsafe.Add(mBase, uint32(v151)+12))
	v171 = *(*int32)(unsafe.Add(mBase, uint32(v167+v158<<(uint(int32(2))%32))))
	v172 = F_statext_is_compatible_clause_internal(m, v171, l1, l2, l3, l4)
	mBase = m.M
	v173 = m.ExcPending
	if v173 != 0 {
		goto L18
	} else {
		goto L65
	}
L64:
	;
	v205 = v172
	goto L1
L65:
	;
	if v172 == int32(0) {
		v205 = v172
		goto L1
	} else {
		goto L66
	}
L66:
	;
	v177 = v158 + int32(1)
	v178 = *(*int32)(unsafe.Add(mBase, uint32(v151)+4))
	if v177 < v178 {
		v158 = v177
		goto L63
	} else {
		goto L67
	}
L67:
	;
	goto L64
L68:
	;
	goto L5
L69:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l3))) = v186
	goto L2
L70:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l3))) = v190
	goto L2
}
func F_statext_mcv_serialize(m *base.Module, l0 int32, l1 int32) int32 {
	mBase := m.M
	_ = mBase
	var v21 int32
	_ = v21
	var v23 int32
	_ = v23
	var v26 int32
	_ = v26
	var v27 int32
	_ = v27
	var v30 int32
	_ = v30
	var v32 int32
	_ = v32
	var v33 int32
	_ = v33
	var v35 int32
	_ = v35
	var v36 int32
	_ = v36
	var v38 int32
	_ = v38
	var v39 int32
	_ = v39
	var v43 int32
	_ = v43
	var v45 int32
	_ = v45
	var v58 int32
	_ = v58
	var v72 int32
	_ = v72
	var v73 int32
	_ = v73
	var v74 int32
	_ = v74
	var v75 int32
	_ = v75
	var v76 int32
	_ = v76
	var v78 int32
	_ = v78
	var v79 int32
	_ = v79
	var v82 int32
	_ = v82
	var v83 int32
	_ = v83
	var v84 int32
	_ = v84
	var v85 int32
	_ = v85
	var v87 int32
	_ = v87
	var v88 int32
	_ = v88
	var v89 int32
	_ = v89
	var v91 int32
	_ = v91
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
	var v102 int32
	_ = v102
	var v103 int32
	_ = v103
	var v122 int32
	_ = v122
	var v123 int32
	_ = v123
	var v125 int32
	_ = v125
	var v128 int32
	_ = v128
	var v129 int32
	_ = v129
	var v130 int32
	_ = v130
	var v133 int32
	_ = v133
	var v137 int64
	_ = v137
	var v139 int32
	_ = v139
	var v143 int32
	_ = v143
	var v144 int32
	_ = v144
	var v146 int32
	_ = v146
	var v168 int32
	_ = v168
	var v169 int32
	_ = v169
	var v174 int32
	_ = v174
	var v176 int32
	_ = v176
	var v178 int32
	_ = v178
	var v179 int32
	_ = v179
	var v180 int32
	_ = v180
	var v183 int32
	_ = v183
	var v185 int32
	_ = v185
	var v186 int32
	_ = v186
	var v187 int32
	_ = v187
	var v191 int32
	_ = v191
	var v192 int32
	_ = v192
	var v194 int32
	_ = v194
	var v199 int32
	_ = v199
	var v200 int32
	_ = v200
	var v219 int32
	_ = v219
	var v220 int32
	_ = v220
	var v221 int32
	_ = v221
	var v224 int64
	_ = v224
	var v225 int64
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
	var v233 int32
	_ = v233
	var v234 int32
	_ = v234
	var v239 int64
	_ = v239
	var v243 int32
	_ = v243
	var v246 int32
	_ = v246
	var v247 int32
	_ = v247
	var v252 int32
	_ = v252
	var v270 int32
	_ = v270
	var v271 int32
	_ = v271
	var v292 int32
	_ = v292
	var v297 int32
	_ = v297
	var v316 int32
	_ = v316
	var v317 int32
	_ = v317
	var v319 int32
	_ = v319
	var v320 int32
	_ = v320
	var v321 int32
	_ = v321
	var v322 int32
	_ = v322
	var v326 int32
	_ = v326
	var v328 int32
	_ = v328
	var v329 int32
	_ = v329
	var v335 int32
	_ = v335
	var v338 int32
	_ = v338
	var v345 int32
	_ = v345
	var v346 int32
	_ = v346
	var v352 int32
	_ = v352
	var v358 int32
	_ = v358
	var v359 int32
	_ = v359
	var v364 int32
	_ = v364
	var v372 int32
	_ = v372
	var v373 int32
	_ = v373
	var v377 int32
	_ = v377
	var v384 int32
	_ = v384
	var v386 int32
	_ = v386
	var v391 int32
	_ = v391
	var v402 int32
	_ = v402
	var v406 int32
	_ = v406
	var v407 int32
	_ = v407
	var v410 int32
	_ = v410
	var v416 int32
	_ = v416
	var v419 int32
	_ = v419
	var v442 int32
	_ = v442
	var v445 int32
	_ = v445
	var v447 int32
	_ = v447
	var v450 int32
	_ = v450
	var v454 int32
	_ = v454
	var v461 int32
	_ = v461
	var v465 int32
	_ = v465
	var v467 int32
	_ = v467
	var v469 int32
	_ = v469
	var v485 int32
	_ = v485
	var v486 int32
	_ = v486
	var v487 int32
	_ = v487
	var v488 int32
	_ = v488
	var v489 int32
	_ = v489
	var v493 int32
	_ = v493
	var v494 int32
	_ = v494
	var v495 int32
	_ = v495
	var v497 int32
	_ = v497
	var v505 int32
	_ = v505
	var v509 int32
	_ = v509
	var v527 int32
	_ = v527
	var v528 int32
	_ = v528
	var v531 int32
	_ = v531
	var v548 int32
	_ = v548
	var v549 int32
	_ = v549
	var v550 int32
	_ = v550
	var v553 int32
	_ = v553
	var v556 int32
	_ = v556
	var v561 int32
	_ = v561
	var v562 int32
	_ = v562
	var v564 int32
	_ = v564
	var v575 int32
	_ = v575
	var v579 int32
	_ = v579
	var v580 int32
	_ = v580
	var v581 int32
	_ = v581
	var v585 int32
	_ = v585
	var v587 int32
	_ = v587
	var v589 int32
	_ = v589
	var v591 int32
	_ = v591
	var v594 int32
	_ = v594
	var v598 int32
	_ = v598
	var v600 int32
	_ = v600
	var v605 int32
	_ = v605
	var v606 int32
	_ = v606
	var v626 int32
	_ = v626
	var v627 int32
	_ = v627
	var v636 int32
	_ = v636
	var v638 int32
	_ = v638
	var v654 int32
	_ = v654
	var v658 int64
	_ = v658
	var v659 int32
	_ = v659
	var v662 int32
	_ = v662
	var v673 int32
	_ = v673
	var v677 int32
	_ = v677
	var v682 int32
	_ = v682
	var v687 int32
	_ = v687
	var v689 int32
	_ = v689
	var v694 int32
	_ = v694
	var v698 int32
	_ = v698
	var v699 int32
	_ = v699
	var v705 int32
	_ = v705
	var v708 int32
	_ = v708
	var v715 int32
	_ = v715
	var v716 int32
	_ = v716
	var v722 int32
	_ = v722
	var v728 int32
	_ = v728
	var v731 int32
	_ = v731
	var v732 int32
	_ = v732
	var v734 int32
	_ = v734
	var v737 int32
	_ = v737
	var v741 int32
	_ = v741
	var v742 int32
	_ = v742
	var v744 int32
	_ = v744
	var v747 int32
	_ = v747
	var v750 int32
	_ = v750
	var v754 int32
	_ = v754
	var v755 int32
	_ = v755
	var v759 int32
	_ = v759
	var v778 int32
	_ = v778
	var v782 int32
	_ = v782
	var v800 int32
	_ = v800
	var v803 int32
	_ = v803
	var v808 int32
	_ = v808
	var v810 int32
	_ = v810
	var v828 int32
	_ = v828
	var v829 int32
	_ = v829
	var v831 int32
	_ = v831
	var v832 int64
	_ = v832
	var v834 int64
	_ = v834
	var v837 int32
	_ = v837
	var v843 int32
	_ = v843
	var v847 int32
	_ = v847
	var v861 int32
	_ = v861
	var v863 int32
	_ = v863
	var v865 int32
	_ = v865
	var v871 int32
	_ = v871
	var v872 int32
	_ = v872
	var v876 int32
	_ = v876
	var v882 int32
	_ = v882
	var v883 int32
	_ = v883
	var v884 int32
	_ = v884
	var v889 int32
	_ = v889
	var v892 int32
	_ = v892
	var v894 int32
	_ = v894
	var v898 int32
	_ = v898
	var v917 int32
	_ = v917
	var v918 int32
	_ = v918
	var v941 int32
	_ = v941
	var v943 int32
	_ = v943
	v21 = m.G0
	v23 = v21 - int32(16)
	m.G0 = v23
	v26 = int32(*(*int16)(unsafe.Add(mBase, uint32(l0)+12)))
	v27 = F_palloc0_mul(m, int32(4), v26)
	mBase = m.M
	v30 = m.ExcPending
	if v30 != 0 {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	return int32(0)
L2:
	;
	v32 = F_palloc0_mul(m, int32(4), v26)
	mBase = m.M
	v33 = m.ExcPending
	if v33 != 0 {
		goto L1
	} else {
		goto L3
	}
L3:
	;
	v35 = F_palloc0_mul(m, int32(20), v26)
	mBase = m.M
	v36 = m.ExcPending
	if v36 != 0 {
		goto L1
	} else {
		goto L4
	}
L4:
	;
	v38 = F_palloc0_mul(m, int32(36), v26)
	mBase = m.M
	v39 = m.ExcPending
	if v39 != 0 {
		goto L1
	} else {
		goto L5
	}
L5:
	;
	if v26 <= int32(0) {
		goto L7
	} else {
		goto L8
	}
L6:
	;
	v575 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	v579 = v561 + v575*v564 + int32(4)
	v580 = F_palloc0(m, v579)
	mBase = m.M
	v581 = m.ExcPending
	if v581 != 0 {
		goto L1
	} else {
		goto L83
	}
L7:
	;
	v43 = v26 << (uint(int32(2)) % 32)
	v45 = v26 * int32(20)
	v556 = v43
	v561 = v43 + v45 + int32(14)
	v562 = v45
	v564 = int32(16)
	goto L6
L8:
	;
	goto L9
L9:
	;
	v58 = int32(0)
	goto L10
L10:
	;
	v72 = int32(2)
	v73 = v58 << (uint(v72) % 32)
	v74 = l1 + v73
	v75 = *(*int32)(unsafe.Add(mBase, uint32(v74)))
	v76 = *(*int32)(unsafe.Add(mBase, uint32(v75)+4))
	v78 = F_lookup_type_cache(m, v76, v72)
	mBase = m.M
	v79 = m.ExcPending
	if v79 != 0 {
		goto L1
	} else {
		goto L12
	}
L11:
	;
	v445 = v26 << (uint(int32(2)) % 32)
	v447 = v26 * int32(20)
	v450 = v445 + v447 + int32(14)
	v454 = v26*int32(3) + int32(16)
	if base.Ui32(v26) < base.Ui32(int32(4)) {
		goto L73
	} else {
		goto L74
	}
L12:
	;
	v82 = v35 + v58*int32(20)
	v83 = *(*int32)(unsafe.Add(mBase, uint32(v74)))
	v84 = *(*int32)(unsafe.Add(mBase, uint32(v83)+12))
	v85 = int32(*(*int16)(unsafe.Add(mBase, uint32(v84)+76)))
	*(*int32)(unsafe.Add(mBase, uint32(v82)+12)) = v85
	v87 = *(*int32)(unsafe.Add(mBase, uint32(v74)))
	v88 = *(*int32)(unsafe.Add(mBase, uint32(v87)+12))
	v89 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v88)+78)))
	*(*uint8)(unsafe.Add(mBase, uint32(v82)+16)) = uint8(v89)
	v91 = v73 + v27
	v93 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	v94 = F_palloc0_mul(m, int32(8), v93)
	mBase = m.M
	v95 = m.ExcPending
	if v95 != 0 {
		goto L1
	} else {
		goto L13
	}
L13:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v91))) = v94
	v97 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	if v97 != 0 {
		goto L14
	} else {
		goto L15
	}
L14:
	;
	v98 = v73 + v32
	v102 = int32(0)
	v103 = v97
	goto L17
L15:
	;
	goto L16
L16:
	;
	v168 = v73 + v32
	v169 = *(*int32)(unsafe.Add(mBase, uint32(v168)))
	if v169 == int32(0) {
		goto L23
	} else {
		goto L24
	}
L17:
	;
	v122 = l0 + int32(48) + v102*int32(24)
	v123 = *(*int32)(unsafe.Add(mBase, uint32(v122)+16))
	v125 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v123+v58))))
	if v125 == int32(0) {
		goto L19
	} else {
		goto L20
	}
L18:
	;
	goto L16
L19:
	;
	v128 = *(*int32)(unsafe.Add(mBase, uint32(v91)))
	v129 = *(*int32)(unsafe.Add(mBase, uint32(v98)))
	v130 = int32(3)
	v133 = *(*int32)(unsafe.Add(mBase, uint32(v122)+20))
	v137 = *(*int64)(unsafe.Add(mBase, uint32(v133+v58<<(uint(v130)%32))))
	*(*int64)(unsafe.Add(mBase, uint32(v128+v129<<(uint(v130)%32)))) = v137
	v139 = *(*int32)(unsafe.Add(mBase, uint32(v98)))
	*(*int32)(unsafe.Add(mBase, uint32(v98))) = v139 + int32(1)
	v143 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	v144 = v143
	goto L21
L20:
	;
	v144 = v103
	goto L21
L21:
	;
	v146 = v102 + int32(1)
	if base.Ui32(v146) < base.Ui32(v144) {
		v102 = v146
		v103 = v144
		goto L17
	} else {
		goto L22
	}
L22:
	;
	goto L18
L23:
	;
	v442 = v58 + int32(1)
	if v442 != v26 {
		v58 = v442
		goto L10
	} else {
		goto L71
	}
L24:
	;
	v174 = v38 + v58*int32(36)
	v176 = *(*int32)(unsafe.Add(mBase, _c_F_statext_mcv_serialize[0]))
	*(*int32)(unsafe.Add(mBase, uint32(v174))) = v176
	v178 = *(*int32)(unsafe.Add(mBase, uint32(v74)))
	v179 = *(*int32)(unsafe.Add(mBase, uint32(v178)+16))
	v180 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v174)+9)) = uint8(v180)
	*(*int32)(unsafe.Add(mBase, uint32(v174)+4)) = v179
	v183 = *(*int32)(unsafe.Add(mBase, uint32(v78)+56))
	F_PrepareSortSupportFromOrderingOp(m, v183, v174)
	mBase = m.M
	v185 = m.ExcPending
	if v185 != 0 {
		goto L1
	} else {
		goto L25
	}
L25:
	;
	v186 = *(*int32)(unsafe.Add(mBase, uint32(v91)))
	v187 = *(*int32)(unsafe.Add(mBase, uint32(v168)))
	F_qsort_interruptible(m, v186, v187, int32(8), int32(1145), v174)
	mBase = m.M
	v191 = m.ExcPending
	if v191 != 0 {
		goto L1
	} else {
		goto L26
	}
L26:
	;
	v192 = int32(1)
	v194 = *(*int32)(unsafe.Add(mBase, uint32(v168)))
	if int32(2) <= v194 {
		goto L27
	} else {
		goto L28
	}
L27:
	;
	v199 = v192
	v200 = v192
	goto L30
L28:
	;
	v252 = v192
	goto L29
L29:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v82))) = v252
	v270 = *(*int32)(unsafe.Add(mBase, uint32(v82)+12))
	v271 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v82)+16)))
	if v271 == int32(1) {
		goto L43
	} else {
		goto L44
	}
L30:
	;
	v219 = v199 << (uint(int32(3)) % 32)
	v220 = *(*int32)(unsafe.Add(mBase, uint32(v91)))
	v221 = v219 + v220
	v224 = *(*int64)(unsafe.Add(mBase, uint32(v221-int32(8))))
	v225 = *(*int64)(unsafe.Add(mBase, uint32(v221)))
	v226 = *(*int32)(unsafe.Add(mBase, uint32(v174)+16))
	v227 = m.T0[v226].(func(*base.Module, int64, int64, int32) int32)(m, v224, v225, v174)
	mBase = m.M
	v228 = m.ExcPending
	if v228 != 0 {
		goto L1
	} else {
		goto L32
	}
L31:
	;
	v252 = v243
	goto L29
L32:
	;
	if v227 < int32(0) {
		goto L33
	} else {
		goto L34
	}
L33:
	;
	v231 = int32(1)
	goto L35
L34:
	;
	v231 = v227
	goto L35
L35:
	;
	v232 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v174)+8)))
	if v232 != 0 {
		goto L36
	} else {
		goto L37
	}
L36:
	;
	v233 = v231
	goto L38
L37:
	;
	v233 = v227
	goto L38
L38:
	;
	if v233 != 0 {
		goto L39
	} else {
		goto L40
	}
L39:
	;
	v234 = *(*int32)(unsafe.Add(mBase, uint32(v91)))
	v239 = *(*int64)(unsafe.Add(mBase, uint32(v234+v219)))
	*(*int64)(unsafe.Add(mBase, uint32(v234+v200<<(uint(int32(3))%32)))) = v239
	v243 = v200 + int32(1)
	goto L41
L40:
	;
	v243 = v200
	goto L41
L41:
	;
	v246 = v199 + int32(1)
	v247 = *(*int32)(unsafe.Add(mBase, uint32(v168)))
	if v246 < v247 {
		v199 = v246
		v200 = v243
		goto L30
	} else {
		goto L42
	}
L42:
	;
	goto L31
L43:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v82)+4)) = v270 * v252
	*(*int32)(unsafe.Add(mBase, uint32(v82)+8)) = int32(0)
	goto L23
L44:
	;
	goto L45
L45:
	;
	if int32(0) < v270 {
		goto L46
	} else {
		goto L47
	}
L46:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v82)+4)) = v270 * v252
	*(*int32)(unsafe.Add(mBase, uint32(v82)+8)) = (v270 + int32(7)) & int32(-8) * v252
	goto L23
L47:
	;
	goto L48
L48:
	;
	switch v270 + int32(2) {
	case 0:
		goto L49
	case 1:
		goto L50
	default:
		goto L23
	}
L49:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v82)+4)) = int64(0)
	v377 = int32(0)
	if v252 <= v377 {
		goto L23
	} else {
		goto L67
	}
L50:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v82)+4)) = int64(0)
	v292 = int32(0)
	if v252 <= v292 {
		goto L23
	} else {
		goto L51
	}
L51:
	;
	v297 = v292
	goto L52
L52:
	;
	v316 = v297 << (uint(int32(3)) % 32)
	v317 = *(*int32)(unsafe.Add(mBase, uint32(v91)))
	v319 = *(*int32)(unsafe.Add(mBase, uint32(v316+v317)))
	v320 = F_pg_detoast_datum(m, v319)
	mBase = m.M
	v321 = m.ExcPending
	if v321 != 0 {
		goto L1
	} else {
		goto L54
	}
L53:
	;
	goto L23
L54:
	;
	v322 = *(*int32)(unsafe.Add(mBase, uint32(v91)))
	*(*int64)(unsafe.Add(mBase, uint32(v322+v316))) = base.I64_extend_i32_u(v320)
	v326 = *(*int32)(unsafe.Add(mBase, uint32(v91)))
	v328 = *(*int32)(unsafe.Add(mBase, uint32(v326+v316)))
	v329 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v328))))
	if v329 == int32(1) {
		goto L56
	} else {
		goto L57
	}
L55:
	;
	v359 = *(*int32)(unsafe.Add(mBase, uint32(v82)+4))
	*(*int32)(unsafe.Add(mBase, uint32(v82)+4)) = v358 + v359 + int32(4)
	v364 = *(*int32)(unsafe.Add(mBase, uint32(v82)+8))
	*(*int32)(unsafe.Add(mBase, uint32(v82)+8)) = v364 + (v358+int32(11))&int32(-8)
	v372 = v297 + int32(1)
	v373 = *(*int32)(unsafe.Add(mBase, uint32(v82)))
	if v372 < v373 {
		v297 = v372
		goto L52
	} else {
		goto L66
	}
L56:
	;
	v335 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v328)+1)))
	if v335 == int32(18) {
		goto L59
	} else {
		goto L60
	}
L57:
	;
	goto L58
L58:
	;
	v346 = int32(1)
	if v329&v346 != 0 {
		v358 = int32(base.Ui32(v329)>>(uint(v346)%32)) - v346
		goto L55
	} else {
		goto L65
	}
L59:
	;
	v338 = int32(16)
	goto L61
L60:
	;
	v338 = int32(0)
	goto L61
L61:
	;
	if base.Ui32((v335-int32(1))&int32(255)) < base.Ui32(int32(3)) {
		goto L62
	} else {
		goto L63
	}
L62:
	;
	v345 = int32(4)
	goto L64
L63:
	;
	v345 = v338
	goto L64
L64:
	;
	v358 = v345
	goto L55
L65:
	;
	v352 = *(*int32)(unsafe.Add(mBase, uint32(v328)))
	v358 = int32(base.Ui32(v352)>>(uint(int32(2))%32)) - int32(4)
	goto L55
L66:
	;
	goto L53
L67:
	;
	v384 = v377
	v386 = v377
	v391 = v377
	goto L68
L68:
	;
	v402 = *(*int32)(unsafe.Add(mBase, uint32(v91)))
	v406 = *(*int32)(unsafe.Add(mBase, uint32(v402+v384<<(uint(int32(3))%32))))
	v407 = F_strlen(m, v406)
	mBase = m.M
	v410 = v407 + v391 + int32(5)
	*(*int32)(unsafe.Add(mBase, uint32(v82)+4)) = v410
	v416 = v407&int32(-8) + v386 + int32(8)
	*(*int32)(unsafe.Add(mBase, uint32(v82)+8)) = v416
	v419 = v384 + int32(1)
	if v419 < v252 {
		v384 = v419
		v386 = v416
		v391 = v410
		goto L68
	} else {
		goto L70
	}
L69:
	;
	goto L23
L70:
	;
	goto L69
L71:
	;
	goto L11
L72:
	;
	v527 = v505
	v528 = int32(0)
	v531 = v509
	goto L80
L73:
	;
	v505 = int32(0)
	v509 = v450
	goto L72
L74:
	;
	goto L75
L75:
	;
	v461 = int32(0)
	v465 = v461
	v467 = v461
	v469 = v450
	goto L76
L76:
	;
	v485 = v35 + v465*int32(20)
	v486 = *(*int32)(unsafe.Add(mBase, uint32(v485)+64))
	v487 = *(*int32)(unsafe.Add(mBase, uint32(v485)+44))
	v488 = *(*int32)(unsafe.Add(mBase, uint32(v485)+24))
	v489 = *(*int32)(unsafe.Add(mBase, uint32(v485)+4))
	v493 = v486 + (v487 + (v488 + (v489 + v469)))
	v494 = int32(4)
	v495 = v465 + v494
	v497 = v467 + v494
	if v497 != v26&int32(_a_F_statext_mcv_serialize_0) {
		v465 = v495
		v467 = v497
		v469 = v493
		goto L76
	} else {
		goto L78
	}
L77:
	;
	if v26&int32(3) == int32(0) {
		v556 = v445
		v561 = v493
		v562 = v447
		v564 = v454
		goto L6
	} else {
		goto L79
	}
L78:
	;
	goto L77
L79:
	;
	v505 = v495
	v509 = v493
	goto L72
L80:
	;
	v548 = *(*int32)(unsafe.Add(mBase, uint32(v35+v527*int32(20))+4))
	v549 = v548 + v531
	v550 = int32(1)
	v553 = v528 + v550
	if v553 != v26&int32(3) {
		v527 = v527 + v550
		v528 = v553
		v531 = v549
		goto L80
	} else {
		goto L82
	}
L81:
	;
	v556 = v445
	v561 = v549
	v562 = v447
	v564 = v454
	goto L6
L82:
	;
	goto L81
L83:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v580))) = v579 << (uint(int32(2)) % 32)
	v585 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	*(*int32)(unsafe.Add(mBase, uint32(v580)+4)) = v585
	v587 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	*(*int32)(unsafe.Add(mBase, uint32(v580)+8)) = v587
	v589 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	*(*int32)(unsafe.Add(mBase, uint32(v580)+12)) = v589
	v591 = int32(*(*uint16)(unsafe.Add(mBase, uint32(l0)+12)))
	*(*uint16)(unsafe.Add(mBase, uint32(v580)+16)) = uint16(v591)
	v594 = v580 + int32(18)
	if v556 != 0 {
		goto L84
	} else {
		goto L85
	}
L84:
	;
	base.MemoryCopy(m, v594, l0+int32(16), v556)
	goto L86
L85:
	;
	goto L86
L86:
	;
	v598 = v556 + v594
	if v562 != 0 {
		goto L87
	} else {
		goto L88
	}
L87:
	;
	base.MemoryCopy(m, v598, v35, v562)
	goto L89
L88:
	;
	goto L89
L89:
	;
	v600 = v598 + v562
	if int32(0) < v26 {
		goto L90
	} else {
		goto L91
	}
L90:
	;
	v605 = int32(0)
	v606 = v600
	goto L93
L91:
	;
	v782 = v600
	goto L92
L92:
	;
	v800 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	if v800 != 0 {
		goto L147
	} else {
		goto L148
	}
L93:
	;
	v626 = v35 + v605*int32(20)
	v627 = *(*int32)(unsafe.Add(mBase, uint32(v626)))
	if int32(0) < v627 {
		goto L95
	} else {
		goto L96
	}
L94:
	;
	v782 = v759
	goto L92
L95:
	;
	v636 = v606
	v638 = int32(0)
	goto L98
L96:
	;
	v759 = v606
	goto L97
L97:
	;
	v778 = v605 + int32(1)
	if v778 != v26 {
		v605 = v778
		v606 = v759
		goto L93
	} else {
		goto L146
	}
L98:
	;
	v654 = *(*int32)(unsafe.Add(mBase, uint32(v27+v605<<(uint(int32(2))%32))))
	v658 = *(*int64)(unsafe.Add(mBase, uint32(v654+v638<<(uint(int32(3))%32))))
	v659 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v626)+16)))
	if v659 == int32(1) {
		goto L101
	} else {
		goto L102
	}
L99:
	;
	v759 = v750
	goto L97
L100:
	;
	v754 = v638 + int32(1)
	v755 = *(*int32)(unsafe.Add(mBase, uint32(v626)))
	if v754 < v755 {
		v636 = v750
		v638 = v754
		goto L98
	} else {
		goto L145
	}
L101:
	;
	v662 = *(*int32)(unsafe.Add(mBase, uint32(v626)+12))
	if base.I32_popcnt(v662) != int32(1) {
		goto L106
	} else {
		goto L107
	}
L102:
	;
	goto L103
L103:
	;
	v689 = *(*int32)(unsafe.Add(mBase, uint32(v626)+12))
	if int32(0) < v689 {
		goto L117
	} else {
		goto L118
	}
L104:
	;
	if v662 != 0 {
		goto L114
	} else {
		goto L115
	}
L105:
	;
	*(*uint8)(unsafe.Add(mBase, uint32(v23)+8)) = uint8(v658)
	goto L104
L106:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v673 = m.ExcPending
	if v673 != 0 {
		goto L1
	} else {
		goto L111
	}
L107:
	;
	switch base.I32_ctz(v662) {
	case 0:
		goto L105
	case 1:
		goto L110
	case 2:
		goto L109
	case 3:
		goto L108
	default:
		goto L106
	}
L108:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v23)+8)) = v658
	goto L104
L109:
	;
	*(*uint32)(unsafe.Add(mBase, uint32(v23)+8)) = uint32(v658)
	goto L104
L110:
	;
	*(*uint16)(unsafe.Add(mBase, uint32(v23)+8)) = uint16(v658)
	goto L104
L111:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v23))) = v662
	F_errmsg_internal(m, int32(_a_F_statext_mcv_serialize_1), v23)
	mBase = m.M
	v677 = m.ExcPending
	if v677 != 0 {
		goto L1
	} else {
		goto L112
	}
L112:
	;
	F_errfinish(m, int32(_a_F_statext_mcv_serialize_2), int32(474), int32(_a_F_statext_mcv_serialize_3))
	mBase = m.M
	v682 = m.ExcPending
	if v682 != 0 {
		goto L1
	} else {
		goto L113
	}
L113:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L114:
	;
	base.MemoryCopy(m, v636, v23+int32(8), v662)
	goto L116
L115:
	;
	goto L116
L116:
	;
	v687 = *(*int32)(unsafe.Add(mBase, uint32(v626)+12))
	v750 = v636 + v687
	goto L100
L117:
	;
	if v689 != 0 {
		goto L120
	} else {
		goto L121
	}
L118:
	;
	goto L119
L119:
	;
	switch v689 + int32(2) {
	case 0:
		goto L123
	case 1:
		goto L124
	default:
		v750 = v636
		goto L100
	}
L120:
	;
	base.MemoryCopy(m, v636, base.I32_wrap_i64(v658), v689)
	goto L122
L121:
	;
	goto L122
L122:
	;
	v694 = *(*int32)(unsafe.Add(mBase, uint32(v626)+12))
	v750 = v636 + v694
	goto L100
L123:
	;
	v741 = base.I32_wrap_i64(v658)
	v742 = F_strlen(m, v741)
	mBase = m.M
	v744 = v742 + int32(1)
	*(*int32)(unsafe.Add(mBase, uint32(v636))) = v744
	v747 = v636 + int32(4)
	if v744 != 0 {
		goto L142
	} else {
		goto L143
	}
L124:
	;
	v698 = base.I32_wrap_i64(v658)
	v699 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v698))))
	if v699 == int32(1) {
		goto L126
	} else {
		goto L127
	}
L125:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v636))) = v728
	v731 = v636 + int32(4)
	if v728 != 0 {
		goto L136
	} else {
		goto L137
	}
L126:
	;
	v705 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v698)+1)))
	if v705 == int32(18) {
		goto L129
	} else {
		goto L130
	}
L127:
	;
	goto L128
L128:
	;
	v716 = int32(1)
	if v699&v716 != 0 {
		v728 = int32(base.Ui32(v699)>>(uint(v716)%32)) - v716
		goto L125
	} else {
		goto L135
	}
L129:
	;
	v708 = int32(16)
	goto L131
L130:
	;
	v708 = int32(0)
	goto L131
L131:
	;
	if base.Ui32((v705-int32(1))&int32(255)) < base.Ui32(int32(3)) {
		goto L132
	} else {
		goto L133
	}
L132:
	;
	v715 = int32(4)
	goto L134
L133:
	;
	v715 = v708
	goto L134
L134:
	;
	v728 = v715
	goto L125
L135:
	;
	v722 = *(*int32)(unsafe.Add(mBase, uint32(v698)))
	v728 = int32(base.Ui32(v722)>>(uint(int32(2))%32)) - int32(4)
	goto L125
L136:
	;
	v732 = int32(1)
	v734 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v698))))
	if v734&v732 != 0 {
		goto L139
	} else {
		goto L140
	}
L137:
	;
	goto L138
L138:
	;
	v750 = v731 + v728
	goto L100
L139:
	;
	v737 = v732
	goto L141
L140:
	;
	v737 = int32(4)
	goto L141
L141:
	;
	base.MemoryCopy(m, v731, v698+v737, v728)
	goto L138
L142:
	;
	base.MemoryCopy(m, v747, v741, v744)
	goto L144
L143:
	;
	goto L144
L144:
	;
	v750 = v747 + v744
	goto L100
L145:
	;
	goto L99
L146:
	;
	goto L94
L147:
	;
	v803 = int32(0)
	v808 = v782
	v810 = v803
	goto L150
L148:
	;
	goto L149
L149:
	;
	F_pfree(m, v27)
	mBase = m.M
	v941 = m.ExcPending
	if v941 != 0 {
		goto L1
	} else {
		goto L166
	}
L150:
	;
	v828 = l0 + int32(48) + v810*int32(24)
	if v26 != 0 {
		goto L152
	} else {
		goto L153
	}
L151:
	;
	goto L149
L152:
	;
	v829 = *(*int32)(unsafe.Add(mBase, uint32(v828)+16))
	base.MemoryCopy(m, v808, v829, v26)
	goto L154
L153:
	;
	goto L154
L154:
	;
	v831 = v808 + v26
	v832 = *(*int64)(unsafe.Add(mBase, uint32(v828)))
	*(*int64)(unsafe.Add(mBase, uint32(v831))) = v832
	v834 = *(*int64)(unsafe.Add(mBase, uint32(v828)+8))
	*(*int64)(unsafe.Add(mBase, uint32(v831)+8)) = v834
	v837 = v831 + int32(16)
	if base.B2i32(v26 <= v803) == int32(0) {
		goto L155
	} else {
		goto L156
	}
L155:
	;
	v843 = v837
	v847 = int32(0)
	goto L158
L156:
	;
	v898 = v837
	goto L157
L157:
	;
	v917 = v810 + int32(1)
	v918 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	if base.Ui32(v917) < base.Ui32(v918) {
		v808 = v898
		v810 = v917
		goto L150
	} else {
		goto L165
	}
L158:
	;
	v861 = *(*int32)(unsafe.Add(mBase, uint32(v828)+16))
	v863 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v861+v847))))
	if v863 != 0 {
		goto L160
	} else {
		goto L161
	}
L159:
	;
	v898 = v892
	goto L157
L160:
	;
	v889 = int32(0)
	goto L162
L161:
	;
	v865 = *(*int32)(unsafe.Add(mBase, uint32(v828)+20))
	v871 = v27 + v847<<(uint(int32(2))%32)
	v872 = *(*int32)(unsafe.Add(mBase, uint32(v871)))
	v876 = *(*int32)(unsafe.Add(mBase, uint32(v35+v847*int32(20))))
	v882 = F_bsearch_arg(m, v865+v847<<(uint(int32(3))%32), v872, v876, int32(8), int32(1145), v38+v847*int32(36))
	mBase = m.M
	v883 = m.ExcPending
	if v883 != 0 {
		goto L1
	} else {
		goto L163
	}
L162:
	;
	*(*uint16)(unsafe.Add(mBase, uint32(v843))) = uint16(v889)
	v892 = v843 + int32(2)
	v894 = v847 + int32(1)
	if v894 != v26 {
		v843 = v892
		v847 = v894
		goto L158
	} else {
		goto L164
	}
L163:
	;
	v884 = *(*int32)(unsafe.Add(mBase, uint32(v871)))
	v889 = int32(base.Ui32(v882-v884) >> (uint(int32(3)) % 32))
	goto L162
L164:
	;
	goto L159
L165:
	;
	goto L151
L166:
	;
	F_pfree(m, v32)
	mBase = m.M
	v943 = m.ExcPending
	if v943 != 0 {
		goto L1
	} else {
		goto L167
	}
L167:
	;
	m.G0 = v23 + int32(16)
	return v580
}
