package p0

import (
	base "github.com/shibukawa/pgmem/internal/aot/pgaot/base"
	"unsafe"
)

func F_internal_load_library(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v13 int32
	_ = v13
	var v15 int32
	_ = v15
	var v20 int32
	_ = v20
	var v30 int32
	_ = v30
	var v32 int32
	_ = v32
	var v35 int32
	_ = v35
	var v38 int32
	_ = v38
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
	var v56 int32
	_ = v56
	var v57 int32
	_ = v57
	var v63 int32
	_ = v63
	var v67 int32
	_ = v67
	var v68 int32
	_ = v68
	var v69 int64
	_ = v69
	var v72 int32
	_ = v72
	var v82 int64
	_ = v82
	var v84 int32
	_ = v84
	var v86 int32
	_ = v86
	var v99 int32
	_ = v99
	var v102 int32
	_ = v102
	var v114 int32
	_ = v114
	var v120 int32
	_ = v120
	var v124 int32
	_ = v124
	var v126 int32
	_ = v126
	var v127 int32
	_ = v127
	var v131 int32
	_ = v131
	var v132 int32
	_ = v132
	var v134 int32
	_ = v134
	var v138 int32
	_ = v138
	var v140 int32
	_ = v140
	var v142 int32
	_ = v142
	var v145 int32
	_ = v145
	var v150 int32
	_ = v150
	var v151 int32
	_ = v151
	var v152 int32
	_ = v152
	var v154 int32
	_ = v154
	var v155 int32
	_ = v155
	var v157 int32
	_ = v157
	var v159 int32
	_ = v159
	var v162 int32
	_ = v162
	var v167 int32
	_ = v167
	var v168 int32
	_ = v168
	var v169 int32
	_ = v169
	var v176 int32
	_ = v176
	var v178 int32
	_ = v178
	var v179 int32
	_ = v179
	var v181 int32
	_ = v181
	var v189 int32
	_ = v189
	var v191 int64
	_ = v191
	var v192 int32
	_ = v192
	var v195 int32
	_ = v195
	var v197 int32
	_ = v197
	var v203 int64
	_ = v203
	var v207 int64
	_ = v207
	var v211 int64
	_ = v211
	var v215 int64
	_ = v215
	var v220 int32
	_ = v220
	var v227 int32
	_ = v227
	var v229 int32
	_ = v229
	var v233 int32
	_ = v233
	var v234 int32
	_ = v234
	var v235 int32
	_ = v235
	var v239 int32
	_ = v239
	var v242 int32
	_ = v242
	var v246 int32
	_ = v246
	var v253 int32
	_ = v253
	var v255 int32
	_ = v255
	var v259 int32
	_ = v259
	var v260 int32
	_ = v260
	var v261 int32
	_ = v261
	var v265 int32
	_ = v265
	var v267 int32
	_ = v267
	var v268 int32
	_ = v268
	var v269 int32
	_ = v269
	var v271 int32
	_ = v271
	var v279 int32
	_ = v279
	var v287 int32
	_ = v287
	var v288 int32
	_ = v288
	var v289 int32
	_ = v289
	var v296 int32
	_ = v296
	var v297 int32
	_ = v297
	var v298 int32
	_ = v298
	var v299 int32
	_ = v299
	var v300 int32
	_ = v300
	var v302 int32
	_ = v302
	var v308 int32
	_ = v308
	var v311 int32
	_ = v311
	var v312 int32
	_ = v312
	var v313 int32
	_ = v313
	var v318 int32
	_ = v318
	var v320 int32
	_ = v320
	var v323 int32
	_ = v323
	var v327 int32
	_ = v327
	var v328 int32
	_ = v328
	var v335 int32
	_ = v335
	var v339 int32
	_ = v339
	var v357 int32
	_ = v357
	var v360 int32
	_ = v360
	var v374 int32
	_ = v374
	var v389 int32
	_ = v389
	var v397 int32
	_ = v397
	var v398 int32
	_ = v398
	var v401 int32
	_ = v401
	var v402 int32
	_ = v402
	var v403 int32
	_ = v403
	var v406 int64
	_ = v406
	var v409 int64
	_ = v409
	var v413 int64
	_ = v413
	var v416 int64
	_ = v416
	var v421 int64
	_ = v421
	var v422 int64
	_ = v422
	var v423 int64
	_ = v423
	var v430 int32
	_ = v430
	var v432 int32
	_ = v432
	var v433 int32
	_ = v433
	var v435 int32
	_ = v435
	var v437 int32
	_ = v437
	var v443 int32
	_ = v443
	var v449 int32
	_ = v449
	var v459 int32
	_ = v459
	var v467 int32
	_ = v467
	var v469 int32
	_ = v469
	var v473 int32
	_ = v473
	var v478 int32
	_ = v478
	var v482 int32
	_ = v482
	var v485 int32
	_ = v485
	var v489 int32
	_ = v489
	var v494 int32
	_ = v494
	var v496 int32
	_ = v496
	var v498 int32
	_ = v498
	var v502 int32
	_ = v502
	var v507 int32
	_ = v507
	var v509 int32
	_ = v509
	var v516 int32
	_ = v516
	var v521 int32
	_ = v521
	var v522 int64
	_ = v522
	var v524 int64
	_ = v524
	var v526 int64
	_ = v526
	var v528 int64
	_ = v528
	var v530 int64
	_ = v530
	var v532 int64
	_ = v532
	var v534 int64
	_ = v534
	var v536 int64
	_ = v536
	var v540 int32
	_ = v540
	var v542 int32
	_ = v542
	var v547 int32
	_ = v547
	var v548 int32
	_ = v548
	var v554 int32
	_ = v554
	var v562 int32
	_ = v562
	var v563 int32
	_ = v563
	var v564 int32
	_ = v564
	var v565 int32
	_ = v565
	var v577 int32
	_ = v577
	var v578 int32
	_ = v578
	var v583 int32
	_ = v583
	var v589 int32
	_ = v589
	var v599 int32
	_ = v599
	var v604 int32
	_ = v604
	var v606 int32
	_ = v606
	var v607 int32
	_ = v607
	var v610 int32
	_ = v610
	var v613 int32
	_ = v613
	var v616 int32
	_ = v616
	var v617 int32
	_ = v617
	var v620 int32
	_ = v620
	var v621 int32
	_ = v621
	var v624 int32
	_ = v624
	var v631 int32
	_ = v631
	var v632 int32
	_ = v632
	var v637 int32
	_ = v637
	var v639 int32
	_ = v639
	var v640 int32
	_ = v640
	var v643 int32
	_ = v643
	var v646 int32
	_ = v646
	var v647 int32
	_ = v647
	var v648 int32
	_ = v648
	var v660 int32
	_ = v660
	var v661 int32
	_ = v661
	var v664 int32
	_ = v664
	var v669 int32
	_ = v669
	var v670 int32
	_ = v670
	var v671 int32
	_ = v671
	var v683 int32
	_ = v683
	var v684 int32
	_ = v684
	var v687 int32
	_ = v687
	var v692 int32
	_ = v692
	var v693 int32
	_ = v693
	var v694 int32
	_ = v694
	var v706 int32
	_ = v706
	var v707 int32
	_ = v707
	var v708 int32
	_ = v708
	var v713 int32
	_ = v713
	var v716 int32
	_ = v716
	var v717 int32
	_ = v717
	var v719 int32
	_ = v719
	var v731 int32
	_ = v731
	var v732 int32
	_ = v732
	var v739 int32
	_ = v739
	var v743 int32
	_ = v743
	var v749 int32
	_ = v749
	var v750 int32
	_ = v750
	var v754 int32
	_ = v754
	var v759 int32
	_ = v759
	var v763 int32
	_ = v763
	var v769 int32
	_ = v769
	var v777 int32
	_ = v777
	var v782 int32
	_ = v782
	var v788 int32
	_ = v788
	var v794 int32
	_ = v794
	var v798 int32
	_ = v798
	var v803 int32
	_ = v803
	v13 = m.G0
	v15 = v13 - int32(208)
	m.G0 = v15
	v20 = int32(_a_F_internal_load_library_0)
	goto L7
L1:
	;
	F_emscripten_builtin_free(m, v102)
	mBase = m.M
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v788 = m.ExcPending
	if v788 != 0 {
		goto L104
	} else {
		goto L207
	}
L2:
	;
	v522 = *(*int64)(unsafe.Add(mBase, uint32(v401)+56))
	*(*int64)(unsafe.Add(mBase, uint32(v15)+104)) = v522
	v524 = *(*int64)(unsafe.Add(mBase, uint32(v401)+48))
	*(*int64)(unsafe.Add(mBase, uint32(v15)+96)) = v524
	v526 = *(*int64)(unsafe.Add(mBase, uint32(v401)+40))
	*(*int64)(unsafe.Add(mBase, uint32(v15)+88)) = v526
	v528 = *(*int64)(unsafe.Add(mBase, uint32(v401)+32))
	*(*int64)(unsafe.Add(mBase, uint32(v15)+80)) = v528
	v530 = *(*int64)(unsafe.Add(mBase, uint32(v401)+24))
	*(*int64)(unsafe.Add(mBase, uint32(v15)+72)) = v530
	v532 = *(*int64)(unsafe.Add(mBase, uint32(v401)+16))
	*(*int64)(unsafe.Add(mBase, uint32(v15)+64)) = v532
	v534 = *(*int64)(unsafe.Add(mBase, uint32(v401)+8))
	*(*int64)(unsafe.Add(mBase, uint32(v15)+56)) = v534
	v536 = *(*int64)(unsafe.Add(mBase, uint32(v401)))
	*(*int64)(unsafe.Add(mBase, uint32(v15)+48)) = v536
	F_emscripten_builtin_free(m, v102)
	mBase = m.M
	v540 = m.G0
	v542 = v540 - int32(224)
	m.G0 = v542
	v547 = v15 + int32(48) | int32(4)
	v548 = *(*int32)(unsafe.Add(mBase, uint32(v547)))
	if v548 != int32(1800) {
		goto L136
	} else {
		goto L137
	}
L3:
	;
	v496 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_internal_load_library[0])))
	if v496 != 0 {
		goto L129
	} else {
		goto L130
	}
L4:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v482 = m.ExcPending
	if v482 != 0 {
		goto L104
	} else {
		goto L125
	}
L5:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v467 = m.ExcPending
	if v467 != 0 {
		goto L104
	} else {
		goto L121
	}
L6:
	;
	v459 = *(*int32)(unsafe.Add(mBase, uint32(v449)+16))
	m.G0 = v15 + int32(208)
	return v459
L7:
	;
	v30 = *(*int32)(unsafe.Add(mBase, uint32(v20)))
	if v30 != 0 {
		goto L9
	} else {
		goto L10
	}
L8:
	;
	v63 = F___fstatat(m, int32(-100), l0, v15+int32(112), int32(0))
	mBase = m.M
	goto L20
L9:
	;
	v32 = v30 + int32(24)
	v35 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0))))
	v38 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v32))))
	if base.B2i32(v35 == int32(0))|base.B2i32(v35 != v38) != 0 {
		v56 = v35
		v57 = v38
		goto L13
	} else {
		goto L14
	}
L10:
	;
	goto L11
L11:
	;
	goto L8
L12:
	;
	if v56-v57 != 0 {
		v20 = v30
		goto L7
	} else {
		goto L19
	}
L13:
	;
	goto L12
L14:
	;
	v41 = l0
	v42 = v32
	goto L15
L15:
	;
	v45 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v42)+1)))
	v46 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v41)+1)))
	if v46 == int32(0) {
		v56 = v46
		v57 = v45
		goto L13
	} else {
		goto L17
	}
L16:
	;
	v56 = v46
	v57 = v45
	goto L13
L17:
	;
	v49 = int32(1)
	if v46 == v45 {
		v41 = v41 + v49
		v42 = v42 + v49
		goto L15
	} else {
		goto L18
	}
L18:
	;
	goto L16
L19:
	;
	v449 = v30
	goto L6
L20:
	;
	if v63 == int32(-1) {
		goto L5
	} else {
		goto L21
	}
L21:
	;
	v67 = *(*int32)(unsafe.Add(mBase, _c_F_internal_load_library[1]))
	if v67 != 0 {
		goto L22
	} else {
		goto L23
	}
L22:
	;
	v68 = *(*int32)(unsafe.Add(mBase, uint32(v15)+112))
	v69 = *(*int64)(unsafe.Add(mBase, uint32(v15)+200))
	v72 = v67
	goto L25
L23:
	;
	goto L24
L24:
	;
	v99 = F_strlen(m, l0)
	mBase = m.M
	v102 = F_emscripten_builtin_malloc(m, v99+int32(25))
	mBase = m.M
	if v102 == int32(0) {
		goto L4
	} else {
		goto L32
	}
L25:
	;
	v82 = *(*int64)(unsafe.Add(mBase, uint32(v72)+8))
	if v82 == v69 {
		goto L27
	} else {
		goto L28
	}
L26:
	;
	goto L24
L27:
	;
	v84 = *(*int32)(unsafe.Add(mBase, uint32(v72)+4))
	if v68 == v84 {
		v449 = v72
		goto L6
	} else {
		goto L30
	}
L28:
	;
	goto L29
L29:
	;
	v86 = *(*int32)(unsafe.Add(mBase, uint32(v72)))
	if v86 != 0 {
		v72 = v86
		goto L25
	} else {
		goto L31
	}
L30:
	;
	goto L29
L31:
	;
	goto L26
L32:
	;
	if v102&int32(3) == int32(0) {
		goto L34
	} else {
		goto L35
	}
L33:
	;
	v114 = v102 + int32(24)
	if (l0^v114)&int32(3) != 0 {
		goto L40
	} else {
		goto L41
	}
L34:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v102)+16)) = int64(0)
	goto L33
L35:
	;
	goto L36
L36:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v102)+16)) = int64(0)
	goto L33
L37:
	;
	v189 = *(*int32)(unsafe.Add(mBase, uint32(v15)+112))
	*(*int32)(unsafe.Add(mBase, uint32(v102)+4)) = v189
	v191 = *(*int64)(unsafe.Add(mBase, uint32(v15)+200))
	v192 = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v102))) = v192
	*(*int64)(unsafe.Add(mBase, uint32(v102)+8)) = v191
	v195 = m.G0
	v197 = v195 - int32(16)
	m.G0 = v197
	if v114 == v192 {
		goto L60
	} else {
		goto L61
	}
L38:
	;
	goto L37
L39:
	;
	*(*uint8)(unsafe.Add(mBase, uint32(v169))) = uint8(v168)
	if v168&int32(255) == int32(0) {
		goto L38
	} else {
		goto L54
	}
L40:
	;
	v120 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0))))
	v167 = l0
	v168 = v120
	v169 = v114
	goto L39
L41:
	;
	goto L42
L42:
	;
	if l0&int32(3) != 0 {
		goto L43
	} else {
		goto L44
	}
L43:
	;
	v124 = l0
	v126 = v114
	goto L46
L44:
	;
	v138 = l0
	v140 = v114
	goto L45
L45:
	;
	v142 = *(*int32)(unsafe.Add(mBase, uint32(v138)))
	v145 = int32(-2139062144)
	if (int32(16843008)-v142|v142)&v145 != v145 {
		v167 = v138
		v168 = v142
		v169 = v140
		goto L39
	} else {
		goto L50
	}
L46:
	;
	v127 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v124))))
	*(*uint8)(unsafe.Add(mBase, uint32(v126))) = uint8(v127)
	if v127 == int32(0) {
		goto L38
	} else {
		goto L48
	}
L47:
	;
	v138 = v134
	v140 = v132
	goto L45
L48:
	;
	v131 = int32(1)
	v132 = v126 + v131
	v134 = v124 + v131
	if v134&int32(3) != 0 {
		v124 = v134
		v126 = v132
		goto L46
	} else {
		goto L49
	}
L49:
	;
	goto L47
L50:
	;
	v150 = v138
	v151 = v142
	v152 = v140
	goto L51
L51:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v152))) = v151
	v154 = int32(4)
	v155 = v152 + v154
	v157 = v150 + v154
	v159 = *(*int32)(unsafe.Add(mBase, uint32(v150)+4))
	v162 = int32(-2139062144)
	if (int32(16843008)-v159|v159)&v162 == v162 {
		v150 = v157
		v151 = v159
		v152 = v155
		goto L51
	} else {
		goto L53
	}
L52:
	;
	v167 = v157
	v168 = v159
	v169 = v155
	goto L39
L53:
	;
	goto L52
L54:
	;
	v176 = v167
	v178 = v169
	goto L55
L55:
	;
	v179 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v176)+1)))
	*(*uint8)(unsafe.Add(mBase, uint32(v178)+1)) = uint8(v179)
	v181 = int32(1)
	if v179 != 0 {
		v176 = v176 + v181
		v178 = v178 + v181
		goto L55
	} else {
		goto L57
	}
L56:
	;
	goto L38
L57:
	;
	goto L56
L58:
	;
	m.G0 = v197 + int32(16)
	*(*int32)(unsafe.Add(mBase, uint32(v102)+16)) = v389
	if v389 == int32(0) {
		goto L3
	} else {
		goto L106
	}
L59:
	;
	v374 = int32(1)
	*(*uint8)(unsafe.Add(mBase, _c_F_internal_load_library[0])) = uint8(v374)
	v389 = int32(0)
	goto L58
L60:
	;
	v203 = *(*int64)(unsafe.Add(mBase, _c_F_internal_load_library[2]))
	*(*int64)(unsafe.Add(mBase, _c_F_internal_load_library[3])) = v203
	v207 = *(*int64)(unsafe.Add(mBase, _c_F_internal_load_library[4]))
	*(*int64)(unsafe.Add(mBase, _c_F_internal_load_library[5])) = v207
	v211 = *(*int64)(unsafe.Add(mBase, _c_F_internal_load_library[6]))
	*(*int64)(unsafe.Add(mBase, _c_F_internal_load_library[7])) = v211
	v215 = *(*int64)(unsafe.Add(mBase, _c_F_internal_load_library[8]))
	*(*int64)(unsafe.Add(mBase, _c_F_internal_load_library[9])) = v215
	goto L59
L61:
	;
	goto L62
L62:
	;
	v220 = F_strlen(m, v114)
	mBase = m.M
	v227 = v220 + int32(1)
	goto L66
L63:
	;
	v269 = int32(0)
	v271 = *(*int32)(unsafe.Add(mBase, _c_F_internal_load_library[10]))
	if v269 < v271 {
		goto L82
	} else {
		goto L83
	}
L64:
	;
	if v239 != 0 {
		goto L70
	} else {
		goto L71
	}
L65:
	;
	goto L64
L66:
	;
	v229 = int32(0)
	if v227 == v229 {
		v239 = v229
		goto L65
	} else {
		goto L68
	}
L67:
	;
	v239 = v234
	goto L65
L68:
	;
	v233 = v227 - int32(1)
	v234 = v114 + v233
	v235 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v234))))
	if v235 != int32(47) {
		v227 = v233
		goto L66
	} else {
		goto L69
	}
L69:
	;
	goto L67
L70:
	;
	v242 = v239 + int32(1)
	goto L72
L71:
	;
	v242 = v114
	goto L72
L72:
	;
	v246 = F_strlen(m, v242)
	mBase = m.M
	v253 = v246 + int32(1)
	goto L75
L73:
	;
	if v265 != 0 {
		goto L79
	} else {
		goto L80
	}
L74:
	;
	goto L73
L75:
	;
	v255 = int32(0)
	if v253 == v255 {
		v265 = v255
		goto L74
	} else {
		goto L77
	}
L76:
	;
	v265 = v260
	goto L74
L77:
	;
	v259 = v253 - int32(1)
	v260 = v242 + v259
	v261 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v260))))
	if v261 != int32(46) {
		v253 = v259
		goto L75
	} else {
		goto L78
	}
L78:
	;
	goto L76
L79:
	;
	v268 = v265 - v242
	goto L63
L80:
	;
	goto L81
L81:
	;
	v267 = F_strlen(m, v242)
	mBase = m.M
	v268 = v267
	goto L63
L82:
	;
	v279 = v269
	goto L85
L83:
	;
	goto L84
L84:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v197))) = v114
	v357 = F_snprintf(m, int32(_a_F_internal_load_library_1), int32(512), int32(_a_F_internal_load_library_2), v197)
	mBase = m.M
	v360 = m.ExcPending
	if v360 != 0 {
		goto L104
	} else {
		goto L105
	}
L85:
	;
	v287 = v279 * int32(12)
	v288 = *(*int32)(unsafe.Add(mBase, uint32(v287)+uint32(_c_F_internal_load_library[11])))
	v289 = F_strlen(m, v288)
	mBase = m.M
	if v289 != v268 {
		goto L87
	} else {
		goto L88
	}
L86:
	;
	goto L84
L87:
	;
	v339 = v279 + int32(1)
	if v339 != v271 {
		v279 = v339
		goto L85
	} else {
		goto L103
	}
L88:
	;
	if v268 == int32(0) {
		goto L90
	} else {
		goto L91
	}
L89:
	;
	if v335 != 0 {
		goto L87
	} else {
		goto L102
	}
L90:
	;
	v335 = int32(0)
	goto L89
L91:
	;
	goto L92
L92:
	;
	v296 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v288))))
	if v296 != 0 {
		goto L93
	} else {
		goto L94
	}
L93:
	;
	v297 = v288
	v298 = v242
	v299 = v268
	v300 = v296
	goto L97
L94:
	;
	v323 = v242
	v327 = int32(0)
	goto L95
L95:
	;
	v328 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v323))))
	v335 = v327 - v328
	goto L89
L96:
	;
	v323 = v318
	v327 = v320
	goto L95
L97:
	;
	v302 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v298))))
	if base.B2i32(v300 != v302)|base.B2i32(v302 == int32(0)) != 0 {
		v318 = v298
		v320 = v300
		goto L96
	} else {
		goto L99
	}
L98:
	;
	v318 = v312
	v320 = int32(0)
	goto L96
L99:
	;
	v308 = v299 - int32(1)
	if v308 == int32(0) {
		v318 = v298
		v320 = v300
		goto L96
	} else {
		goto L100
	}
L100:
	;
	v311 = int32(1)
	v312 = v298 + v311
	v313 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v297)+1)))
	if v313 != 0 {
		v297 = v297 + v311
		v298 = v312
		v299 = v308
		v300 = v313
		goto L97
	} else {
		goto L101
	}
L101:
	;
	goto L98
L102:
	;
	v389 = v287 + int32(_a_F_internal_load_library_3)
	goto L58
L103:
	;
	goto L86
L104:
	;
	return int32(0)
L105:
	;
	goto L59
L106:
	;
	v397 = F_pgmem_dlsym(m, v389, int32(_a_F_internal_load_library_4))
	mBase = m.M
	v398 = m.ExcPending
	if v398 != 0 {
		goto L104
	} else {
		goto L107
	}
L107:
	;
	if v397 == int32(0) {
		goto L1
	} else {
		goto L108
	}
L108:
	;
	v401 = m.T0[v397].(func(*base.Module) int32)(m)
	mBase = m.M
	v402 = m.ExcPending
	if v402 != 0 {
		goto L104
	} else {
		goto L109
	}
L109:
	;
	v403 = *(*int32)(unsafe.Add(mBase, uint32(v401)))
	if v403 != int32(64) {
		goto L2
	} else {
		goto L110
	}
L110:
	;
	v406 = *(*int64)(unsafe.Add(mBase, uint32(v401)+4))
	v409 = *(*int64)(unsafe.Add(mBase, uint32(v401)+12))
	v413 = *(*int64)(unsafe.Add(mBase, uint32(v401)+20))
	v416 = *(*int64)(unsafe.Add(mBase, uint32(v401)+28))
	v421 = int64(*(*uint32)(unsafe.Add(mBase, uint32(v401)+52)))
	v422 = *(*int64)(unsafe.Add(mBase, uint32(v401)+36))
	v423 = *(*int64)(unsafe.Add(mBase, uint32(v401)+44))
	if v406^int64(429496731400)|(v409^int64(274877906976))|(v413^int64(8391172920109432832)|(v416^int64(83912175219303)))|(v421|(v422|v423)) != int64(0) {
		goto L2
	} else {
		goto L111
	}
L111:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v102)+20)) = v401
	v430 = *(*int32)(unsafe.Add(mBase, uint32(v102)+16))
	v432 = F_pgmem_dlsym(m, v430, int32(_a_F_internal_load_library_5))
	mBase = m.M
	v433 = m.ExcPending
	if v433 != 0 {
		goto L104
	} else {
		goto L112
	}
L112:
	;
	if v432 != 0 {
		goto L113
	} else {
		goto L114
	}
L113:
	;
	m.T0[v432].(func(*base.Module))(m)
	mBase = m.M
	v435 = m.ExcPending
	if v435 != 0 {
		goto L104
	} else {
		goto L116
	}
L114:
	;
	goto L115
L115:
	;
	v437 = *(*int32)(unsafe.Add(mBase, _c_F_internal_load_library[1]))
	if v437 == int32(0) {
		goto L118
	} else {
		goto L119
	}
L116:
	;
	goto L115
L117:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_internal_load_library[12])) = v102
	v449 = v102
	goto L6
L118:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_internal_load_library[1])) = v102
	goto L117
L119:
	;
	goto L120
L120:
	;
	v443 = *(*int32)(unsafe.Add(mBase, _c_F_internal_load_library[12]))
	*(*int32)(unsafe.Add(mBase, uint32(v443))) = v102
	goto L117
L121:
	;
	F_errcode_for_file_access(m)
	mBase = m.M
	v469 = m.ExcPending
	if v469 != 0 {
		goto L104
	} else {
		goto L122
	}
L122:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v15))) = l0
	F_errmsg(m, int32(_a_F_internal_load_library_6), v15)
	mBase = m.M
	v473 = m.ExcPending
	if v473 != 0 {
		goto L104
	} else {
		goto L123
	}
L123:
	;
	F_errfinish(m, int32(_a_F_internal_load_library_7), int32(215), int32(_a_F_internal_load_library_8))
	mBase = m.M
	v478 = m.ExcPending
	if v478 != 0 {
		goto L104
	} else {
		goto L124
	}
L124:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L125:
	;
	F_errcode(m, int32(_a_F_internal_load_library_9))
	mBase = m.M
	v485 = m.ExcPending
	if v485 != 0 {
		goto L104
	} else {
		goto L126
	}
L126:
	;
	F_errmsg(m, int32(_a_F_internal_load_library_10), int32(0))
	mBase = m.M
	v489 = m.ExcPending
	if v489 != 0 {
		goto L104
	} else {
		goto L127
	}
L127:
	;
	F_errfinish(m, int32(_a_F_internal_load_library_7), int32(234), int32(_a_F_internal_load_library_8))
	mBase = m.M
	v494 = m.ExcPending
	if v494 != 0 {
		goto L104
	} else {
		goto L128
	}
L128:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L129:
	;
	v498 = int32(0)
	*(*uint8)(unsafe.Add(mBase, _c_F_internal_load_library[0])) = uint8(v498)
	v502 = int32(_a_F_internal_load_library_1)
	goto L131
L130:
	;
	v502 = int32(0)
	goto L131
L131:
	;
	F_emscripten_builtin_free(m, v102)
	mBase = m.M
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v507 = m.ExcPending
	if v507 != 0 {
		goto L104
	} else {
		goto L132
	}
L132:
	;
	F_errcode_for_file_access(m)
	mBase = m.M
	v509 = m.ExcPending
	if v509 != 0 {
		goto L104
	} else {
		goto L133
	}
L133:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v15)+20)) = v502
	*(*int32)(unsafe.Add(mBase, uint32(v15)+16)) = l0
	F_errmsg(m, int32(_a_F_internal_load_library_11), v15+int32(16))
	mBase = m.M
	v516 = m.ExcPending
	if v516 != 0 {
		goto L104
	} else {
		goto L134
	}
L134:
	;
	F_errfinish(m, int32(_a_F_internal_load_library_7), int32(253), int32(_a_F_internal_load_library_8))
	mBase = m.M
	v521 = m.ExcPending
	if v521 != 0 {
		goto L104
	} else {
		goto L135
	}
L135:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L136:
	;
	if int32(1000) <= v548 {
		goto L140
	} else {
		goto L141
	}
L137:
	;
	goto L138
L138:
	;
	v606 = v547 + int32(20)
	v607 = int32(_a_F_internal_load_library_12)
	v610 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v606))))
	v613 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_internal_load_library[13])))
	if base.B2i32(v610 == int32(0))|base.B2i32(v610 != v613) != 0 {
		v631 = v610
		v632 = v613
		goto L150
	} else {
		goto L151
	}
L139:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v583 = m.ExcPending
	if v583 != 0 {
		goto L104
	} else {
		goto L145
	}
L140:
	;
	v554 = base.I32_div_u_s(v548, int32(100))
	*(*int32)(unsafe.Add(mBase, uint32(v542)+160)) = v554
	v562 = F_pg_snprintf(m, v542+int32(192), int32(32), int32(_a_F_internal_load_library_13), v542+int32(160))
	mBase = m.M
	v563 = m.ExcPending
	if v563 != 0 {
		goto L104
	} else {
		goto L143
	}
L141:
	;
	goto L142
L142:
	;
	v564 = int32(100)
	v565 = base.I32_div_s(v548, v564)
	*(*int32)(unsafe.Add(mBase, uint32(v542)+176)) = v565
	*(*int32)(unsafe.Add(mBase, uint32(v542)+180)) = v548 - v565*v564
	v577 = F_pg_snprintf(m, v542+int32(192), int32(32), int32(_a_F_internal_load_library_14), v542+int32(176))
	mBase = m.M
	v578 = m.ExcPending
	if v578 != 0 {
		goto L104
	} else {
		goto L144
	}
L143:
	;
	goto L139
L144:
	;
	goto L139
L145:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v542)+144)) = l0
	F_errmsg(m, int32(_a_F_internal_load_library_15), v542+int32(144))
	mBase = m.M
	v589 = m.ExcPending
	if v589 != 0 {
		goto L104
	} else {
		goto L146
	}
L146:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v542)+128)) = int32(18)
	*(*int32)(unsafe.Add(mBase, uint32(v542)+132)) = v542 + int32(192)
	F_errdetail(m, int32(_a_F_internal_load_library_16), v542+int32(128))
	mBase = m.M
	v599 = m.ExcPending
	if v599 != 0 {
		goto L104
	} else {
		goto L147
	}
L147:
	;
	F_errfinish(m, int32(_a_F_internal_load_library_7), int32(340), int32(_a_F_internal_load_library_17))
	mBase = m.M
	v604 = m.ExcPending
	if v604 != 0 {
		goto L104
	} else {
		goto L148
	}
L148:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L149:
	;
	if v631-v632 == int32(0) {
		goto L156
	} else {
		goto L157
	}
L150:
	;
	goto L149
L151:
	;
	v616 = v606
	v617 = v607
	goto L152
L152:
	;
	v620 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v617)+1)))
	v621 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v616)+1)))
	if v621 == int32(0) {
		v631 = v621
		v632 = v620
		goto L150
	} else {
		goto L154
	}
L153:
	;
	v631 = v621
	v632 = v620
	goto L150
L154:
	;
	v624 = int32(1)
	if v621 == v620 {
		v616 = v616 + v624
		v617 = v617 + v624
		goto L152
	} else {
		goto L155
	}
L155:
	;
	goto L153
L156:
	;
	v637 = v542 + int32(192)
	F_initStringInfo(m, v637)
	mBase = m.M
	v639 = m.ExcPending
	if v639 != 0 {
		goto L104
	} else {
		goto L159
	}
L157:
	;
	goto L158
L158:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v763 = m.ExcPending
	if v763 != 0 {
		goto L104
	} else {
		goto L203
	}
L159:
	;
	v640 = *(*int32)(unsafe.Add(mBase, uint32(v547)+4))
	if v640 != int32(100) {
		goto L160
	} else {
		goto L161
	}
L160:
	;
	v643 = *(*int32)(unsafe.Add(mBase, uint32(v542)+196))
	if v643 != 0 {
		goto L163
	} else {
		goto L164
	}
L161:
	;
	goto L162
L162:
	;
	v661 = *(*int32)(unsafe.Add(mBase, uint32(v547)+8))
	if v661 != int32(32) {
		goto L168
	} else {
		goto L169
	}
L163:
	;
	F_appendStringInfoChar(m, v637, int32(10))
	mBase = m.M
	v646 = m.ExcPending
	if v646 != 0 {
		goto L104
	} else {
		goto L166
	}
L164:
	;
	v648 = v640
	goto L165
L165:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v542)+88)) = v648
	*(*int32)(unsafe.Add(mBase, uint32(v542)+84)) = int32(100)
	*(*int32)(unsafe.Add(mBase, uint32(v542)+80)) = int32(_a_F_internal_load_library_18)
	F_appendStringInfo(m, v542+int32(192), int32(_a_F_internal_load_library_19), v542+int32(80))
	mBase = m.M
	v660 = m.ExcPending
	if v660 != 0 {
		goto L104
	} else {
		goto L167
	}
L166:
	;
	v647 = *(*int32)(unsafe.Add(mBase, uint32(v547)+4))
	v648 = v647
	goto L165
L167:
	;
	goto L162
L168:
	;
	v664 = *(*int32)(unsafe.Add(mBase, uint32(v542)+196))
	if v664 != 0 {
		goto L171
	} else {
		goto L172
	}
L169:
	;
	goto L170
L170:
	;
	v684 = *(*int32)(unsafe.Add(mBase, uint32(v547)+12))
	if v684 != int32(64) {
		goto L176
	} else {
		goto L177
	}
L171:
	;
	F_appendStringInfoChar(m, v542+int32(192), int32(10))
	mBase = m.M
	v669 = m.ExcPending
	if v669 != 0 {
		goto L104
	} else {
		goto L174
	}
L172:
	;
	v671 = v661
	goto L173
L173:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v542)+72)) = v671
	*(*int32)(unsafe.Add(mBase, uint32(v542)+68)) = int32(32)
	*(*int32)(unsafe.Add(mBase, uint32(v542)+64)) = int32(_a_F_internal_load_library_20)
	F_appendStringInfo(m, v542+int32(192), int32(_a_F_internal_load_library_19), v542-int32(-64))
	mBase = m.M
	v683 = m.ExcPending
	if v683 != 0 {
		goto L104
	} else {
		goto L175
	}
L174:
	;
	v670 = *(*int32)(unsafe.Add(mBase, uint32(v547)+8))
	v671 = v670
	goto L173
L175:
	;
	goto L170
L176:
	;
	v687 = *(*int32)(unsafe.Add(mBase, uint32(v542)+196))
	if v687 != 0 {
		goto L179
	} else {
		goto L180
	}
L177:
	;
	goto L178
L178:
	;
	v707 = *(*int32)(unsafe.Add(mBase, uint32(v547)+16))
	if v707 != 0 {
		goto L184
	} else {
		goto L185
	}
L179:
	;
	F_appendStringInfoChar(m, v542+int32(192), int32(10))
	mBase = m.M
	v692 = m.ExcPending
	if v692 != 0 {
		goto L104
	} else {
		goto L182
	}
L180:
	;
	v694 = v684
	goto L181
L181:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v542)+56)) = v694
	*(*int32)(unsafe.Add(mBase, uint32(v542)+52)) = int32(64)
	*(*int32)(unsafe.Add(mBase, uint32(v542)+48)) = int32(_a_F_internal_load_library_21)
	F_appendStringInfo(m, v542+int32(192), int32(_a_F_internal_load_library_19), v542+int32(48))
	mBase = m.M
	v706 = m.ExcPending
	if v706 != 0 {
		goto L104
	} else {
		goto L183
	}
L182:
	;
	v693 = *(*int32)(unsafe.Add(mBase, uint32(v547)+12))
	v694 = v693
	goto L181
L183:
	;
	goto L178
L184:
	;
	v708 = *(*int32)(unsafe.Add(mBase, uint32(v542)+196))
	if v708 != 0 {
		goto L187
	} else {
		goto L188
	}
L185:
	;
	goto L186
L186:
	;
	v732 = *(*int32)(unsafe.Add(mBase, uint32(v542)+196))
	if v732 == int32(0) {
		goto L195
	} else {
		goto L196
	}
L187:
	;
	F_appendStringInfoChar(m, v542+int32(192), int32(10))
	mBase = m.M
	v713 = m.ExcPending
	if v713 != 0 {
		goto L104
	} else {
		goto L190
	}
L188:
	;
	v719 = int32(_a_F_internal_load_library_22)
	goto L189
L189:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v542)+40)) = v719
	*(*int32)(unsafe.Add(mBase, uint32(v542)+36)) = int32(_a_F_internal_load_library_23)
	*(*int32)(unsafe.Add(mBase, uint32(v542)+32)) = int32(_a_F_internal_load_library_24)
	F_appendStringInfo(m, v542+int32(192), int32(_a_F_internal_load_library_25), v542+int32(32))
	mBase = m.M
	v731 = m.ExcPending
	if v731 != 0 {
		goto L104
	} else {
		goto L194
	}
L190:
	;
	v716 = *(*int32)(unsafe.Add(mBase, uint32(v547)+16))
	if v716 != 0 {
		goto L191
	} else {
		goto L192
	}
L191:
	;
	v717 = int32(_a_F_internal_load_library_22)
	goto L193
L192:
	;
	v717 = int32(_a_F_internal_load_library_23)
	goto L193
L193:
	;
	v719 = v717
	goto L189
L194:
	;
	goto L186
L195:
	;
	F_appendStringInfoString(m, v542+int32(192), int32(_a_F_internal_load_library_26))
	mBase = m.M
	v739 = m.ExcPending
	if v739 != 0 {
		goto L104
	} else {
		goto L198
	}
L196:
	;
	goto L197
L197:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v743 = m.ExcPending
	if v743 != 0 {
		goto L104
	} else {
		goto L199
	}
L198:
	;
	goto L197
L199:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v542)+16)) = l0
	F_errmsg(m, int32(_a_F_internal_load_library_27), v542+int32(16))
	mBase = m.M
	v749 = m.ExcPending
	if v749 != 0 {
		goto L104
	} else {
		goto L200
	}
L200:
	;
	v750 = *(*int32)(unsafe.Add(mBase, uint32(v542)+192))
	*(*int32)(unsafe.Add(mBase, uint32(v542))) = v750
	F_errdetail_internal(m, int32(_a_F_internal_load_library_28), v542)
	mBase = m.M
	v754 = m.ExcPending
	if v754 != 0 {
		goto L104
	} else {
		goto L201
	}
L201:
	;
	F_errfinish(m, int32(_a_F_internal_load_library_7), int32(414), int32(_a_F_internal_load_library_17))
	mBase = m.M
	v759 = m.ExcPending
	if v759 != 0 {
		goto L104
	} else {
		goto L202
	}
L202:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L203:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v542)+112)) = l0
	F_errmsg(m, int32(_a_F_internal_load_library_29), v542+int32(112))
	mBase = m.M
	v769 = m.ExcPending
	if v769 != 0 {
		goto L104
	} else {
		goto L204
	}
L204:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v542)+100)) = v606
	*(*int32)(unsafe.Add(mBase, uint32(v542)+96)) = int32(_a_F_internal_load_library_12)
	F_errdetail(m, int32(_a_F_internal_load_library_30), v542+int32(96))
	mBase = m.M
	v777 = m.ExcPending
	if v777 != 0 {
		goto L104
	} else {
		goto L205
	}
L205:
	;
	F_errfinish(m, int32(_a_F_internal_load_library_7), int32(355), int32(_a_F_internal_load_library_17))
	mBase = m.M
	v782 = m.ExcPending
	if v782 != 0 {
		goto L104
	} else {
		goto L206
	}
L206:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L207:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v15)+32)) = l0
	F_errmsg(m, int32(_a_F_internal_load_library_31), v15+int32(32))
	mBase = m.M
	v794 = m.ExcPending
	if v794 != 0 {
		goto L104
	} else {
		goto L208
	}
L208:
	;
	F_errhint(m, int32(_a_F_internal_load_library_32), int32(0))
	mBase = m.M
	v798 = m.ExcPending
	if v798 != 0 {
		goto L104
	} else {
		goto L209
	}
L209:
	;
	F_errfinish(m, int32(_a_F_internal_load_library_7), int32(291), int32(_a_F_internal_load_library_8))
	mBase = m.M
	v803 = m.ExcPending
	if v803 != 0 {
		goto L104
	} else {
		goto L210
	}
L210:
	;
	base.Wasm_trap_unreachable()
	for {
	}
}
