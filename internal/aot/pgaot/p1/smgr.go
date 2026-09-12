package p1

import (
	base "github.com/shibukawa/pgmem/internal/aot/pgaot/base"
	"unsafe"
)

func F_smgrDoPendingSyncs(m *base.Module, l0 int32, l1 int32) {
	mBase := m.M
	_ = mBase
	var v3 int32
	_ = v3
	var v15 int32
	_ = v15
	var v17 int32
	_ = v17
	var v20 int32
	_ = v20
	var v33 int32
	_ = v33
	var v34 int32
	_ = v34
	var v48 int32
	_ = v48
	var v52 int32
	_ = v52
	var v55 int32
	_ = v55
	var v56 int32
	_ = v56
	var v57 int32
	_ = v57
	var v59 int32
	_ = v59
	var v74 int32
	_ = v74
	var v76 int32
	_ = v76
	var v79 int32
	_ = v79
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
	var v91 int32
	_ = v91
	var v92 int32
	_ = v92
	var v94 int32
	_ = v94
	var v96 int32
	_ = v96
	var v98 int64
	_ = v98
	var v103 int32
	_ = v103
	var v104 int32
	_ = v104
	var v105 int32
	_ = v105
	var v109 int32
	_ = v109
	var v112 int32
	_ = v112
	var v113 int32
	_ = v113
	var v115 int32
	_ = v115
	var v116 int32
	_ = v116
	var v118 int32
	_ = v118
	var v119 int64
	_ = v119
	var v121 int32
	_ = v121
	var v122 int32
	_ = v122
	var v124 int32
	_ = v124
	var v125 int32
	_ = v125
	var v128 int32
	_ = v128
	var v129 int64
	_ = v129
	var v132 int32
	_ = v132
	var v133 int32
	_ = v133
	var v138 int32
	_ = v138
	var v139 int32
	_ = v139
	var v142 int32
	_ = v142
	var v143 int64
	_ = v143
	var v145 int32
	_ = v145
	var v146 int32
	_ = v146
	var v148 int32
	_ = v148
	var v149 int32
	_ = v149
	var v152 int32
	_ = v152
	var v153 int64
	_ = v153
	var v154 int32
	_ = v154
	var v160 int32
	_ = v160
	var v161 int32
	_ = v161
	var v162 int32
	_ = v162
	var v163 int32
	_ = v163
	var v164 int64
	_ = v164
	var v166 int64
	_ = v166
	var v174 int32
	_ = v174
	var v176 int64
	_ = v176
	var v180 int32
	_ = v180
	var v181 int32
	_ = v181
	var v182 int32
	_ = v182
	var v185 int32
	_ = v185
	var v187 int32
	_ = v187
	var v191 int32
	_ = v191
	var v193 int64
	_ = v193
	var v197 int32
	_ = v197
	var v198 int32
	_ = v198
	var v202 int32
	_ = v202
	var v204 int32
	_ = v204
	var v208 int32
	_ = v208
	var v210 int64
	_ = v210
	var v214 int32
	_ = v214
	var v215 int32
	_ = v215
	var v219 int32
	_ = v219
	var v221 int32
	_ = v221
	var v226 int32
	_ = v226
	var v228 int64
	_ = v228
	var v230 int32
	_ = v230
	var v231 int32
	_ = v231
	var v235 int32
	_ = v235
	var v237 int32
	_ = v237
	var v238 int32
	_ = v238
	var v239 int32
	_ = v239
	var v240 int32
	_ = v240
	var v241 int32
	_ = v241
	var v247 int32
	_ = v247
	var v248 int32
	_ = v248
	var v252 int32
	_ = v252
	var v253 int32
	_ = v253
	var v256 int32
	_ = v256
	var v257 int32
	_ = v257
	var v266 int32
	_ = v266
	var v267 int32
	_ = v267
	var v268 int32
	_ = v268
	var v269 int32
	_ = v269
	var v270 int32
	_ = v270
	var v271 int32
	_ = v271
	var v272 int32
	_ = v272
	var v276 int32
	_ = v276
	var v277 int32
	_ = v277
	var v279 int32
	_ = v279
	var v285 int32
	_ = v285
	var v287 int32
	_ = v287
	var v291 int32
	_ = v291
	var v292 int32
	_ = v292
	var v293 int32
	_ = v293
	var v304 int32
	_ = v304
	var v308 int32
	_ = v308
	var v314 int32
	_ = v314
	var v316 int32
	_ = v316
	var v317 int32
	_ = v317
	var v319 int32
	_ = v319
	var v320 int32
	_ = v320
	var v321 int64
	_ = v321
	var v323 int32
	_ = v323
	var v325 int32
	_ = v325
	var v328 int32
	_ = v328
	var v331 int32
	_ = v331
	var v334 int32
	_ = v334
	var v335 int32
	_ = v335
	var v336 int64
	_ = v336
	var v338 int32
	_ = v338
	var v340 int32
	_ = v340
	var v343 int32
	_ = v343
	var v345 int32
	_ = v345
	var v351 int32
	_ = v351
	var v365 int32
	_ = v365
	var v368 int32
	_ = v368
	var v369 int32
	_ = v369
	var v370 int64
	_ = v370
	var v372 int32
	_ = v372
	var v374 int32
	_ = v374
	var v385 int32
	_ = v385
	var v401 int32
	_ = v401
	var v403 int32
	_ = v403
	var v414 int32
	_ = v414
	var v422 int32
	_ = v422
	var v425 int32
	_ = v425
	var v430 int32
	_ = v430
	var v436 int32
	_ = v436
	var v448 int32
	_ = v448
	var v449 int32
	_ = v449
	var v451 int32
	_ = v451
	var v452 int32
	_ = v452
	var v454 int32
	_ = v454
	var v455 int32
	_ = v455
	var v458 int32
	_ = v458
	var v460 int64
	_ = v460
	var v461 int32
	_ = v461
	var v468 int32
	_ = v468
	var v469 int32
	_ = v469
	var v470 int32
	_ = v470
	var v487 int32
	_ = v487
	var v489 int32
	_ = v489
	var v491 int32
	_ = v491
	var v502 int32
	_ = v502
	var v503 int32
	_ = v503
	var v525 int32
	_ = v525
	var v526 int32
	_ = v526
	var v527 int32
	_ = v527
	var v536 int32
	_ = v536
	var v549 int32
	_ = v549
	var v550 int32
	_ = v550
	var v552 int32
	_ = v552
	var v557 int32
	_ = v557
	var v560 int32
	_ = v560
	var v567 int32
	_ = v567
	var v569 int32
	_ = v569
	var v570 int32
	_ = v570
	var v572 int32
	_ = v572
	var v573 int32
	_ = v573
	var v575 int32
	_ = v575
	var v579 int32
	_ = v579
	var v580 int32
	_ = v580
	var v582 int32
	_ = v582
	var v583 int32
	_ = v583
	var v588 int32
	_ = v588
	var v589 int32
	_ = v589
	var v590 int32
	_ = v590
	var v597 int32
	_ = v597
	var v600 int32
	_ = v600
	var v603 int32
	_ = v603
	var v605 int32
	_ = v605
	var v607 int32
	_ = v607
	var v608 int32
	_ = v608
	var v609 int32
	_ = v609
	var v612 int32
	_ = v612
	var v614 int32
	_ = v614
	var v616 int32
	_ = v616
	var v617 int32
	_ = v617
	var v622 int32
	_ = v622
	var v624 int32
	_ = v624
	var v643 int32
	_ = v643
	var v645 int32
	_ = v645
	var v662 int32
	_ = v662
	var v680 int32
	_ = v680
	var v682 int32
	_ = v682
	var v698 int32
	_ = v698
	var v704 int32
	_ = v704
	var v705 int32
	_ = v705
	var v706 int32
	_ = v706
	var v708 int32
	_ = v708
	var v714 int32
	_ = v714
	var v715 int32
	_ = v715
	var v716 int32
	_ = v716
	var v717 int32
	_ = v717
	var v719 int32
	_ = v719
	var v721 int32
	_ = v721
	var v722 int32
	_ = v722
	var v724 int32
	_ = v724
	var v725 int32
	_ = v725
	var v726 int32
	_ = v726
	var v728 int32
	_ = v728
	var v730 int32
	_ = v730
	var v731 int32
	_ = v731
	var v733 int32
	_ = v733
	var v734 int32
	_ = v734
	var v735 int32
	_ = v735
	var v737 int32
	_ = v737
	var v739 int32
	_ = v739
	var v740 int32
	_ = v740
	var v742 int32
	_ = v742
	var v743 int32
	_ = v743
	var v744 int32
	_ = v744
	var v746 int32
	_ = v746
	var v748 int32
	_ = v748
	var v750 int32
	_ = v750
	var v766 int32
	_ = v766
	var v768 int32
	_ = v768
	var v787 int32
	_ = v787
	v3 = int32(0)
	v15 = m.G0
	v17 = v15 - int32(96)
	m.G0 = v17
	v20 = *(*int32)(unsafe.Add(mBase, _consts[330]))
	if v20 == v3 {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	m.G0 = v17 + int32(96)
	return
L2:
	;
	if l0 == int32(0) {
		goto L3
	} else {
		goto L4
	}
L3:
	;
	*(*int32)(unsafe.Add(mBase, _consts[330])) = int32(0)
	goto L1
L4:
	;
	goto L5
L5:
	;
	if l1 == int32(0) {
		goto L6
	} else {
		goto L7
	}
L6:
	;
	v33 = *(*int32)(unsafe.Add(mBase, _consts[212]))
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
	*(*int32)(unsafe.Add(mBase, _consts[330])) = int32(0)
	goto L1
L9:
	;
	v34 = v33
	goto L12
L10:
	;
	v74 = v20
	goto L11
L11:
	;
	F_hash_seq_init(m, v17+int32(76), v74)
	mBase = m.M
	v76 = m.ExcPending
	if v76 != 0 {
		goto L17
	} else {
		goto L20
	}
L12:
	;
	v48 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v34)+16)))
	if v48 == int32(1) {
		goto L14
	} else {
		goto L15
	}
L13:
	;
	v59 = *(*int32)(unsafe.Add(mBase, _consts[330]))
	v74 = v59
	goto L11
L14:
	;
	v52 = *(*int32)(unsafe.Add(mBase, _consts[330]))
	v55 = F_hash_search(m, v52, v34, int32(2), int32(0))
	mBase = m.M
	v56 = m.ExcPending
	if v56 != 0 {
		goto L17
	} else {
		goto L18
	}
L15:
	;
	goto L16
L16:
	;
	v57 = *(*int32)(unsafe.Add(mBase, uint32(v34)+24))
	if v57 != 0 {
		v34 = v57
		goto L12
	} else {
		goto L19
	}
L17:
	;
	return
L18:
	;
	goto L16
L19:
	;
	goto L13
L20:
	;
	v79 = F_hash_seq_search(m, v17+int32(76))
	mBase = m.M
	v80 = m.ExcPending
	if v80 != 0 {
		goto L17
	} else {
		goto L21
	}
L21:
	;
	if v79 != 0 {
		goto L22
	} else {
		goto L23
	}
L22:
	;
	v83 = v79
	v85 = int32(0)
	v88 = v3
	v89 = v3
	v90 = v3
	v91 = v3
	v92 = v3
	v94 = v3
	goto L25
L23:
	;
	goto L24
L24:
	;
	*(*int32)(unsafe.Add(mBase, _consts[330])) = int32(0)
	goto L1
L25:
	;
	v96 = *(*int32)(unsafe.Add(mBase, uint32(v83)+8))
	*(*int32)(unsafe.Add(mBase, uint32(v17)+72)) = v96
	v98 = *(*int64)(unsafe.Add(mBase, uint32(v83)))
	*(*int64)(unsafe.Add(mBase, uint32(v17)+64)) = v98
	v103 = F_smgropen(m, v17-int32(-64), int32(-1))
	mBase = m.M
	v104 = m.ExcPending
	if v104 != 0 {
		goto L17
	} else {
		goto L27
	}
L26:
	;
	v279 = int32(0)
	*(*int32)(unsafe.Add(mBase, _consts[330])) = v279
	if v267 <= v279 {
		goto L1
	} else {
		goto L91
	}
L27:
	;
	v105 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v83)+12)))
	if v105 == int32(0) {
		goto L31
	} else {
		goto L32
	}
L28:
	;
	v276 = F_hash_seq_search(m, v17+int32(76))
	mBase = m.M
	v277 = m.ExcPending
	if v277 != 0 {
		goto L17
	} else {
		goto L89
	}
L29:
	;
	if v85 == int32(0) {
		goto L83
	} else {
		goto L84
	}
L30:
	;
	v166 = int64(*(*int32)(unsafe.Add(mBase, _consts[331])))
	if base.Ui64(int64(base.Ui64(v166)>>(uint(int64(3))%64))&int64(2251799813685247)) <= base.Ui64(v164) {
		v238 = v160
		v239 = v161
		v240 = v162
		v241 = v163
		goto L29
	} else {
		goto L57
	}
L31:
	;
	v109 = int32(-1)
	v112 = F_smgrexists(m, v103, int32(0))
	mBase = m.M
	v113 = m.ExcPending
	if v113 != 0 {
		goto L17
	} else {
		goto L34
	}
L32:
	;
	goto L33
L33:
	;
	if v105 != 0 {
		v238 = v89
		v239 = v90
		v240 = v91
		v241 = v92
		goto L29
	} else {
		goto L56
	}
L34:
	;
	if v112 != 0 {
		goto L35
	} else {
		goto L36
	}
L35:
	;
	v115 = F_smgrnblocks(m, v103, int32(0))
	mBase = m.M
	v116 = m.ExcPending
	if v116 != 0 {
		goto L17
	} else {
		goto L38
	}
L36:
	;
	v118 = v109
	v119 = int64(0)
	goto L37
L37:
	;
	v121 = F_smgrexists(m, v103, int32(1))
	mBase = m.M
	v122 = m.ExcPending
	if v122 != 0 {
		goto L17
	} else {
		goto L39
	}
L38:
	;
	v118 = v115
	v119 = base.I64_extend_i32_u(v115)
	goto L37
L39:
	;
	if v121 != 0 {
		goto L40
	} else {
		goto L41
	}
L40:
	;
	v124 = F_smgrnblocks(m, v103, int32(1))
	mBase = m.M
	v125 = m.ExcPending
	if v125 != 0 {
		goto L17
	} else {
		goto L43
	}
L41:
	;
	v128 = v109
	v129 = v119
	goto L42
L42:
	;
	v132 = F_smgrexists(m, v103, int32(2))
	mBase = m.M
	v133 = m.ExcPending
	if v133 != 0 {
		goto L17
	} else {
		goto L45
	}
L43:
	;
	v128 = v124
	v129 = v119 + base.I64_extend_i32_u(v124)
	goto L42
L44:
	;
	v145 = F_smgrexists(m, v103, int32(3))
	mBase = m.M
	v146 = m.ExcPending
	if v146 != 0 {
		goto L17
	} else {
		goto L50
	}
L45:
	;
	if v132 == int32(0) {
		goto L46
	} else {
		goto L47
	}
L46:
	;
	v142 = int32(-1)
	v143 = v129
	goto L44
L47:
	;
	goto L48
L48:
	;
	v138 = F_smgrnblocks(m, v103, int32(2))
	mBase = m.M
	v139 = m.ExcPending
	if v139 != 0 {
		goto L17
	} else {
		goto L49
	}
L49:
	;
	v142 = v138
	v143 = v129 + base.I64_extend_i32_u(v138)
	goto L44
L50:
	;
	if v145 != 0 {
		goto L51
	} else {
		goto L52
	}
L51:
	;
	v148 = F_smgrnblocks(m, v103, int32(3))
	mBase = m.M
	v149 = m.ExcPending
	if v149 != 0 {
		goto L17
	} else {
		goto L54
	}
L52:
	;
	v152 = int32(-1)
	v153 = v143
	goto L53
L53:
	;
	v154 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v83)+12)))
	if v154&int32(1) == int32(0) {
		v160 = v128
		v161 = v118
		v162 = v142
		v163 = v152
		v164 = v153
		goto L30
	} else {
		goto L55
	}
L54:
	;
	v152 = v148
	v153 = v143 + base.I64_extend_i32_u(v148)
	goto L53
L55:
	;
	v238 = v128
	v239 = v118
	v240 = v142
	v241 = v152
	goto L29
L56:
	;
	v160 = v89
	v161 = v90
	v162 = v91
	v163 = v92
	v164 = int64(0)
	goto L30
L57:
	;
	if v161 != int32(-1) {
		goto L58
	} else {
		goto L59
	}
L58:
	;
	v174 = *(*int32)(unsafe.Add(mBase, uint32(v103)+8))
	*(*int32)(unsafe.Add(mBase, uint32(v17)+56)) = v174
	v176 = *(*int64)(unsafe.Add(mBase, uint32(v103)))
	*(*int64)(unsafe.Add(mBase, uint32(v17)+48)) = v176
	v180 = F_CreateFakeRelcacheEntry(m, v17+int32(48))
	mBase = m.M
	v181 = m.ExcPending
	if v181 != 0 {
		goto L17
	} else {
		goto L61
	}
L59:
	;
	goto L60
L60:
	;
	if v160 != int32(-1) {
		goto L64
	} else {
		goto L65
	}
L61:
	;
	v182 = int32(0)
	F_log_newpage_range(m, v180, v182, v161, v182)
	mBase = m.M
	v185 = m.ExcPending
	if v185 != 0 {
		goto L17
	} else {
		goto L62
	}
L62:
	;
	F_pfree(m, v180)
	mBase = m.M
	v187 = m.ExcPending
	if v187 != 0 {
		goto L17
	} else {
		goto L63
	}
L63:
	;
	goto L60
L64:
	;
	v191 = *(*int32)(unsafe.Add(mBase, uint32(v103)+8))
	*(*int32)(unsafe.Add(mBase, uint32(v17)+40)) = v191
	v193 = *(*int64)(unsafe.Add(mBase, uint32(v103)))
	*(*int64)(unsafe.Add(mBase, uint32(v17)+32)) = v193
	v197 = F_CreateFakeRelcacheEntry(m, v17+int32(32))
	mBase = m.M
	v198 = m.ExcPending
	if v198 != 0 {
		goto L17
	} else {
		goto L67
	}
L65:
	;
	goto L66
L66:
	;
	if v162 != int32(-1) {
		goto L70
	} else {
		goto L71
	}
L67:
	;
	F_log_newpage_range(m, v197, int32(1), v160, int32(0))
	mBase = m.M
	v202 = m.ExcPending
	if v202 != 0 {
		goto L17
	} else {
		goto L68
	}
L68:
	;
	F_pfree(m, v197)
	mBase = m.M
	v204 = m.ExcPending
	if v204 != 0 {
		goto L17
	} else {
		goto L69
	}
L69:
	;
	goto L66
L70:
	;
	v208 = *(*int32)(unsafe.Add(mBase, uint32(v103)+8))
	*(*int32)(unsafe.Add(mBase, uint32(v17)+24)) = v208
	v210 = *(*int64)(unsafe.Add(mBase, uint32(v103)))
	*(*int64)(unsafe.Add(mBase, uint32(v17)+16)) = v210
	v214 = F_CreateFakeRelcacheEntry(m, v17+int32(16))
	mBase = m.M
	v215 = m.ExcPending
	if v215 != 0 {
		goto L17
	} else {
		goto L73
	}
L71:
	;
	goto L72
L72:
	;
	if v163 == int32(-1) {
		goto L76
	} else {
		goto L77
	}
L73:
	;
	F_log_newpage_range(m, v214, int32(2), v162, int32(0))
	mBase = m.M
	v219 = m.ExcPending
	if v219 != 0 {
		goto L17
	} else {
		goto L74
	}
L74:
	;
	F_pfree(m, v214)
	mBase = m.M
	v221 = m.ExcPending
	if v221 != 0 {
		goto L17
	} else {
		goto L75
	}
L75:
	;
	goto L72
L76:
	;
	v266 = v85
	v267 = v88
	v268 = v160
	v269 = v161
	v270 = v162
	v271 = int32(-1)
	v272 = v94
	goto L28
L77:
	;
	goto L78
L78:
	;
	v226 = *(*int32)(unsafe.Add(mBase, uint32(v103)+8))
	*(*int32)(unsafe.Add(mBase, uint32(v17)+8)) = v226
	v228 = *(*int64)(unsafe.Add(mBase, uint32(v103)))
	*(*int64)(unsafe.Add(mBase, uint32(v17))) = v228
	v230 = F_CreateFakeRelcacheEntry(m, v17)
	mBase = m.M
	v231 = m.ExcPending
	if v231 != 0 {
		goto L17
	} else {
		goto L79
	}
L79:
	;
	F_log_newpage_range(m, v230, int32(3), v163, int32(0))
	mBase = m.M
	v235 = m.ExcPending
	if v235 != 0 {
		goto L17
	} else {
		goto L80
	}
L80:
	;
	F_pfree(m, v230)
	mBase = m.M
	v237 = m.ExcPending
	if v237 != 0 {
		goto L17
	} else {
		goto L81
	}
L81:
	;
	v266 = v85
	v267 = v88
	v268 = v160
	v269 = v161
	v270 = v162
	v271 = v163
	v272 = v94
	goto L28
L82:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v257+v88<<(uint(int32(2))%32)))) = v103
	v266 = v256
	v267 = v88 + int32(1)
	v268 = v238
	v269 = v239
	v270 = v240
	v271 = v241
	v272 = v257
	goto L28
L83:
	;
	v247 = F_palloc(m, int32(32))
	mBase = m.M
	v248 = m.ExcPending
	if v248 != 0 {
		goto L17
	} else {
		goto L86
	}
L84:
	;
	goto L85
L85:
	;
	if v88 < v85 {
		v256 = v85
		v257 = v94
		goto L82
	} else {
		goto L87
	}
L86:
	;
	v256 = int32(8)
	v257 = v247
	goto L82
L87:
	;
	v252 = F_repalloc(m, v94, v85<<(uint(int32(3))%32))
	mBase = m.M
	v253 = m.ExcPending
	if v253 != 0 {
		goto L17
	} else {
		goto L88
	}
L88:
	;
	v256 = v85 << (uint(int32(1)) % 32)
	v257 = v252
	goto L82
L89:
	;
	if v276 != 0 {
		v83 = v276
		v85 = v266
		v88 = v267
		v89 = v268
		v90 = v269
		v91 = v270
		v92 = v271
		v94 = v272
		goto L25
	} else {
		goto L90
	}
L90:
	;
	goto L26
L91:
	;
	if v267 != 0 {
		goto L92
	} else {
		goto L93
	}
L92:
	;
	v285 = m.G0
	v287 = v285 - int32(32)
	m.G0 = v287
	if v267 != 0 {
		goto L95
	} else {
		goto L96
	}
L93:
	;
	goto L94
L94:
	;
	F_pfree(m, v272)
	mBase = m.M
	v787 = m.ExcPending
	if v787 != 0 {
		goto L17
	} else {
		goto L191
	}
L95:
	;
	v291 = F_palloc(m, v267<<(uint(int32(4))%32))
	mBase = m.M
	v292 = m.ExcPending
	if v292 != 0 {
		goto L17
	} else {
		goto L98
	}
L96:
	;
	goto L97
L97:
	;
	m.G0 = v287 + int32(32)
	v680 = int32(4437228)
	v682 = *(*int32)(unsafe.Add(mBase, _consts[163]))
	*(*int32)(unsafe.Add(mBase, _consts[163])) = v682 + int32(1)
	if int32(0) < v267 {
		goto L165
	} else {
		goto L166
	}
L98:
	;
	v293 = int32(0)
	if v267 <= v293 {
		v401 = v293
		goto L99
	} else {
		goto L100
	}
L99:
	;
	v403 = *(*int32)(unsafe.Add(mBase, _consts[136]))
	if int32(0) < v403 {
		goto L112
	} else {
		goto L113
	}
L100:
	;
	if v267 != int32(1) {
		goto L101
	} else {
		goto L102
	}
L101:
	;
	v304 = v3
	v308 = int32(0)
	goto L104
L102:
	;
	v351 = v3
	goto L103
L103:
	;
	if v267&int32(1) != 0 {
		goto L107
	} else {
		goto L108
	}
L104:
	;
	v314 = int32(4)
	v316 = v291 + v304<<(uint(v314)%32)
	v317 = int32(2)
	v319 = v272 + v304<<(uint(v317)%32)
	v320 = *(*int32)(unsafe.Add(mBase, uint32(v319)))
	v321 = *(*int64)(unsafe.Add(mBase, uint32(v320)))
	*(*int64)(unsafe.Add(mBase, uint32(v316))) = v321
	v323 = *(*int32)(unsafe.Add(mBase, uint32(v320)+8))
	*(*int32)(unsafe.Add(mBase, uint32(v316)+8)) = v323
	v325 = *(*int32)(unsafe.Add(mBase, uint32(v319)))
	*(*int32)(unsafe.Add(mBase, uint32(v316)+12)) = v325
	v328 = v304 | int32(1)
	v331 = v291 + v328<<(uint(v314)%32)
	v334 = v272 + v328<<(uint(v317)%32)
	v335 = *(*int32)(unsafe.Add(mBase, uint32(v334)))
	v336 = *(*int64)(unsafe.Add(mBase, uint32(v335)))
	*(*int64)(unsafe.Add(mBase, uint32(v331))) = v336
	v338 = *(*int32)(unsafe.Add(mBase, uint32(v335)+8))
	*(*int32)(unsafe.Add(mBase, uint32(v331)+8)) = v338
	v340 = *(*int32)(unsafe.Add(mBase, uint32(v334)))
	*(*int32)(unsafe.Add(mBase, uint32(v331)+12)) = v340
	v343 = v304 + v317
	v345 = v308 + v317
	if v345 != v267&int32(2147483646) {
		v304 = v343
		v308 = v345
		goto L104
	} else {
		goto L106
	}
L105:
	;
	v351 = v343
	goto L103
L106:
	;
	goto L105
L107:
	;
	v365 = v291 + v351<<(uint(int32(4))%32)
	v368 = v272 + v351<<(uint(int32(2))%32)
	v369 = *(*int32)(unsafe.Add(mBase, uint32(v368)))
	v370 = *(*int64)(unsafe.Add(mBase, uint32(v369)))
	*(*int64)(unsafe.Add(mBase, uint32(v365))) = v370
	v372 = *(*int32)(unsafe.Add(mBase, uint32(v369)+8))
	*(*int32)(unsafe.Add(mBase, uint32(v365)+8)) = v372
	v374 = *(*int32)(unsafe.Add(mBase, uint32(v368)))
	*(*int32)(unsafe.Add(mBase, uint32(v365)+12)) = v374
	goto L109
L108:
	;
	goto L109
L109:
	;
	if v267 < int32(21) {
		v401 = int32(0)
		goto L99
	} else {
		goto L110
	}
L110:
	;
	F_pg_qsort(m, v291, v267, int32(16), int32(1077))
	mBase = m.M
	v385 = m.ExcPending
	if v385 != 0 {
		goto L17
	} else {
		goto L111
	}
L111:
	;
	v401 = int32(1)
	goto L99
L112:
	;
	v414 = int32(0)
	goto L115
L113:
	;
	goto L114
L114:
	;
	F_pfree(m, v291)
	mBase = m.M
	v662 = m.ExcPending
	if v662 != 0 {
		goto L17
	} else {
		goto L164
	}
L115:
	;
	v422 = *(*int32)(unsafe.Add(mBase, _consts[9]))
	v425 = v422 + v414<<(uint(int32(6))%32)
	if v401 == int32(0) {
		goto L119
	} else {
		goto L120
	}
L116:
	;
	goto L114
L117:
	;
	v643 = v414 + int32(1)
	v645 = *(*int32)(unsafe.Add(mBase, _consts[136]))
	if v643 < v645 {
		v414 = v643
		goto L115
	} else {
		goto L163
	}
L118:
	;
	if v470 == int32(0) {
		goto L117
	} else {
		goto L131
	}
L119:
	;
	if v267 <= int32(0) {
		goto L117
	} else {
		goto L122
	}
L120:
	;
	goto L121
L121:
	;
	v460 = *(*int64)(unsafe.Add(mBase, uint32(v425)))
	v461 = *(*int32)(unsafe.Add(mBase, uint32(v425)+8))
	*(*int32)(unsafe.Add(mBase, uint32(v287)+16)) = v461
	*(*int64)(unsafe.Add(mBase, uint32(v287)+8)) = v460
	v468 = F_bsearch(m, v287+int32(8), v291, v267, int32(16), int32(1077))
	mBase = m.M
	v469 = m.ExcPending
	if v469 != 0 {
		goto L17
	} else {
		goto L130
	}
L122:
	;
	v430 = *(*int32)(unsafe.Add(mBase, uint32(v425)))
	v436 = int32(0)
	goto L123
L123:
	;
	v448 = v291 + v436<<(uint(int32(4))%32)
	v449 = *(*int32)(unsafe.Add(mBase, uint32(v448)))
	if v430 != v449 {
		goto L125
	} else {
		goto L126
	}
L124:
	;
	goto L117
L125:
	;
	v458 = v436 + int32(1)
	if v458 != v267 {
		v436 = v458
		goto L123
	} else {
		goto L129
	}
L126:
	;
	v451 = *(*int32)(unsafe.Add(mBase, uint32(v425)+4))
	v452 = *(*int32)(unsafe.Add(mBase, uint32(v448)+4))
	if v451 != v452 {
		goto L125
	} else {
		goto L127
	}
L127:
	;
	v454 = *(*int32)(unsafe.Add(mBase, uint32(v425)+8))
	v455 = *(*int32)(unsafe.Add(mBase, uint32(v448)+8))
	if v454 == v455 {
		v470 = v448
		goto L118
	} else {
		goto L128
	}
L128:
	;
	goto L125
L129:
	;
	goto L124
L130:
	;
	v470 = v468
	goto L118
L131:
	;
	F_ReservePrivateRefCountEntry(m)
	mBase = m.M
	v487 = m.ExcPending
	if v487 != 0 {
		goto L17
	} else {
		goto L132
	}
L132:
	;
	v489 = *(*int32)(unsafe.Add(mBase, _consts[175]))
	F_ResourceOwnerEnlarge(m, v489)
	mBase = m.M
	v491 = m.ExcPending
	if v491 != 0 {
		goto L17
	} else {
		goto L133
	}
L133:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v287)+28)) = int32(217995)
	*(*int32)(unsafe.Add(mBase, uint32(v287)+24)) = int32(6259)
	*(*int32)(unsafe.Add(mBase, uint32(v287)+20)) = int32(472001)
	*(*int32)(unsafe.Add(mBase, uint32(v287)+16)) = int32(0)
	*(*int64)(unsafe.Add(mBase, uint32(v287)+8)) = int64(0)
	v502 = *(*int32)(unsafe.Add(mBase, uint32(v425)+24))
	v503 = int32(4194304)
	*(*int32)(unsafe.Add(mBase, uint32(v425)+24)) = v502 | v503
	if v502&v503 != 0 {
		goto L134
	} else {
		goto L135
	}
L134:
	;
	goto L137
L135:
	;
	v536 = v502
	goto L136
L136:
	;
	v549 = int32(4062540)
	v550 = *(*int32)(unsafe.Add(mBase, _consts[276]))
	v552 = *(*int32)(unsafe.Add(mBase, uint32(v287+int32(8))+8))
	if v552 == int32(0) {
		goto L144
	} else {
		goto L145
	}
L137:
	;
	F_perform_spin_delay(m, v287+int32(8))
	mBase = m.M
	v525 = m.ExcPending
	if v525 != 0 {
		goto L17
	} else {
		goto L139
	}
L138:
	;
	v536 = v526
	goto L136
L139:
	;
	v526 = *(*int32)(unsafe.Add(mBase, uint32(v425)+24))
	v527 = int32(4194304)
	*(*int32)(unsafe.Add(mBase, uint32(v425)+24)) = v526 | v527
	if v526&v527 != 0 {
		goto L137
	} else {
		goto L140
	}
L140:
	;
	goto L138
L141:
	;
	v569 = *(*int32)(unsafe.Add(mBase, uint32(v425)))
	v570 = *(*int32)(unsafe.Add(mBase, uint32(v470)))
	if v569 != v570 {
		goto L152
	} else {
		goto L153
	}
L142:
	;
	goto L141
L143:
	;
	*(*int32)(unsafe.Add(mBase, _consts[276])) = v567
	goto L142
L144:
	;
	if int32(999) < v550 {
		goto L142
	} else {
		goto L147
	}
L145:
	;
	goto L146
L146:
	;
	if v550 < int32(11) {
		goto L142
	} else {
		goto L151
	}
L147:
	;
	v557 = int32(900)
	if v557 <= v550 {
		goto L148
	} else {
		goto L149
	}
L148:
	;
	v560 = v557
	goto L150
L149:
	;
	v560 = v550
	goto L150
L150:
	;
	v567 = v560 + int32(100)
	goto L143
L151:
	;
	v567 = v550 - int32(1)
	goto L143
L152:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v425)+24)) = v536 & int32(-4194305)
	goto L117
L153:
	;
	v572 = *(*int32)(unsafe.Add(mBase, uint32(v425)+4))
	v573 = *(*int32)(unsafe.Add(mBase, uint32(v470)+4))
	if v572 != v573 {
		goto L152
	} else {
		goto L154
	}
L154:
	;
	v575 = int32(25165824)
	if v536&v575 != v575 {
		goto L152
	} else {
		goto L155
	}
L155:
	;
	v579 = *(*int32)(unsafe.Add(mBase, uint32(v425)+8))
	v580 = *(*int32)(unsafe.Add(mBase, uint32(v470)+8))
	if v579 != v580 {
		goto L152
	} else {
		goto L156
	}
L156:
	;
	v582 = *(*int32)(unsafe.Add(mBase, uint32(v425)+24))
	v583 = int32(1)
	*(*int32)(unsafe.Add(mBase, uint32(v425)+24)) = (v582 + v583) & int32(-4194305)
	v588 = *(*int32)(unsafe.Add(mBase, uint32(v425)+20))
	v589 = int32(4358488)
	v590 = *(*int32)(unsafe.Add(mBase, _consts[278]))
	*(*int32)(unsafe.Add(mBase, _consts[278])) = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v590)+4)) = v583
	v597 = v588 + v583
	*(*int32)(unsafe.Add(mBase, uint32(v590))) = v597
	v600 = *(*int32)(unsafe.Add(mBase, _consts[175]))
	F_ResourceOwnerRemember(m, v600, v597, int32(1586128))
	mBase = m.M
	v603 = m.ExcPending
	if v603 != 0 {
		goto L17
	} else {
		goto L157
	}
L157:
	;
	v605 = v425 + int32(48)
	v607 = F_LWLockAcquire(m, v605, int32(1))
	mBase = m.M
	v608 = m.ExcPending
	if v608 != 0 {
		goto L17
	} else {
		goto L158
	}
L158:
	;
	v609 = *(*int32)(unsafe.Add(mBase, uint32(v470)+12))
	F_FlushBuffer(m, v425, v609, int32(3))
	mBase = m.M
	v612 = m.ExcPending
	if v612 != 0 {
		goto L17
	} else {
		goto L159
	}
L159:
	;
	F_LWLockRelease(m, v605)
	mBase = m.M
	v614 = m.ExcPending
	if v614 != 0 {
		goto L17
	} else {
		goto L160
	}
L160:
	;
	v616 = *(*int32)(unsafe.Add(mBase, _consts[175]))
	v617 = *(*int32)(unsafe.Add(mBase, uint32(v425)+20))
	F_ResourceOwnerForget(m, v616, v617+int32(1), int32(1586128))
	mBase = m.M
	v622 = m.ExcPending
	if v622 != 0 {
		goto L17
	} else {
		goto L161
	}
L161:
	;
	F_UnpinBufferNoOwner(m, v425)
	mBase = m.M
	v624 = m.ExcPending
	if v624 != 0 {
		goto L17
	} else {
		goto L162
	}
L162:
	;
	goto L117
L163:
	;
	goto L116
L164:
	;
	goto L97
L165:
	;
	v698 = int32(0)
	goto L168
L166:
	;
	goto L167
L167:
	;
	v766 = int32(4437228)
	v768 = *(*int32)(unsafe.Add(mBase, _consts[163]))
	*(*int32)(unsafe.Add(mBase, _consts[163])) = v768 - int32(1)
	goto L94
L168:
	;
	v704 = v272 + v698<<(uint(int32(2))%32)
	v705 = *(*int32)(unsafe.Add(mBase, uint32(v704)))
	v706 = *(*int32)(unsafe.Add(mBase, uint32(v705)+36))
	v708 = v706 * int32(80)
	v714 = *(*int32)(unsafe.Add(mBase, uint32(v708)+uint32(_consts[332])))
	v715 = m.T0[v714].(func(*base.Module, int32, int32) int32)(m, v705, int32(0))
	mBase = m.M
	v716 = m.ExcPending
	if v716 != 0 {
		goto L17
	} else {
		goto L170
	}
L169:
	;
	goto L167
L170:
	;
	if v715 != 0 {
		goto L171
	} else {
		goto L172
	}
L171:
	;
	v717 = *(*int32)(unsafe.Add(mBase, uint32(v704)))
	v719 = *(*int32)(unsafe.Add(mBase, uint32(v708)+uint32(_consts[333])))
	m.T0[v719].(func(*base.Module, int32, int32))(m, v717, int32(0))
	mBase = m.M
	v721 = m.ExcPending
	if v721 != 0 {
		goto L17
	} else {
		goto L174
	}
L172:
	;
	goto L173
L173:
	;
	v722 = *(*int32)(unsafe.Add(mBase, uint32(v704)))
	v724 = m.T0[v714].(func(*base.Module, int32, int32) int32)(m, v722, int32(1))
	mBase = m.M
	v725 = m.ExcPending
	if v725 != 0 {
		goto L17
	} else {
		goto L175
	}
L174:
	;
	goto L173
L175:
	;
	if v724 != 0 {
		goto L176
	} else {
		goto L177
	}
L176:
	;
	v726 = *(*int32)(unsafe.Add(mBase, uint32(v704)))
	v728 = *(*int32)(unsafe.Add(mBase, uint32(v708)+uint32(_consts[333])))
	m.T0[v728].(func(*base.Module, int32, int32))(m, v726, int32(1))
	mBase = m.M
	v730 = m.ExcPending
	if v730 != 0 {
		goto L17
	} else {
		goto L179
	}
L177:
	;
	goto L178
L178:
	;
	v731 = *(*int32)(unsafe.Add(mBase, uint32(v704)))
	v733 = m.T0[v714].(func(*base.Module, int32, int32) int32)(m, v731, int32(2))
	mBase = m.M
	v734 = m.ExcPending
	if v734 != 0 {
		goto L17
	} else {
		goto L180
	}
L179:
	;
	goto L178
L180:
	;
	if v733 != 0 {
		goto L181
	} else {
		goto L182
	}
L181:
	;
	v735 = *(*int32)(unsafe.Add(mBase, uint32(v704)))
	v737 = *(*int32)(unsafe.Add(mBase, uint32(v708)+uint32(_consts[333])))
	m.T0[v737].(func(*base.Module, int32, int32))(m, v735, int32(2))
	mBase = m.M
	v739 = m.ExcPending
	if v739 != 0 {
		goto L17
	} else {
		goto L184
	}
L182:
	;
	goto L183
L183:
	;
	v740 = *(*int32)(unsafe.Add(mBase, uint32(v704)))
	v742 = m.T0[v714].(func(*base.Module, int32, int32) int32)(m, v740, int32(3))
	mBase = m.M
	v743 = m.ExcPending
	if v743 != 0 {
		goto L17
	} else {
		goto L185
	}
L184:
	;
	goto L183
L185:
	;
	if v742 != 0 {
		goto L186
	} else {
		goto L187
	}
L186:
	;
	v744 = *(*int32)(unsafe.Add(mBase, uint32(v704)))
	v746 = *(*int32)(unsafe.Add(mBase, uint32(v708)+uint32(_consts[333])))
	m.T0[v746].(func(*base.Module, int32, int32))(m, v744, int32(3))
	mBase = m.M
	v748 = m.ExcPending
	if v748 != 0 {
		goto L17
	} else {
		goto L189
	}
L187:
	;
	goto L188
L188:
	;
	v750 = v698 + int32(1)
	if v750 != v267 {
		v698 = v750
		goto L168
	} else {
		goto L190
	}
L189:
	;
	goto L188
L190:
	;
	goto L169
L191:
	;
	goto L1
}
func F_smgr_bulk_get_buf(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v2 int32
	_ = v2
	var v6 int32
	_ = v6
	var v9 int32
	_ = v9
	v2 = *(*int32)(unsafe.Add(mBase, uint32(l0)+416))
	v6 = F_MemoryContextAllocAligned(m, v2, int32(8192), int32(4096), int32(0))
	mBase = m.M
	v9 = m.ExcPending
	if v9 != 0 {
		return int32(0)
	} else {
		return v6
	}
}
func F_smgr_bulk_start_rel(m *base.Module, l0 int32, l1 int32) int32 {
	mBase := m.M
	_ = mBase
	var v6 int32
	_ = v6
	var v8 int32
	_ = v8
	var v10 int32
	_ = v10
	var v13 int32
	_ = v13
	var v14 int32
	_ = v14
	var v16 int64
	_ = v16
	var v18 int32
	_ = v18
	var v21 int32
	_ = v21
	var v24 int32
	_ = v24
	var v25 int32
	_ = v25
	var v26 int32
	_ = v26
	var v28 int32
	_ = v28
	var v30 int32
	_ = v30
	var v32 int32
	_ = v32
	var v36 int32
	_ = v36
	var v37 int32
	_ = v37
	var v38 int32
	_ = v38
	var v39 int32
	_ = v39
	var v42 int32
	_ = v42
	var v44 int32
	_ = v44
	var v47 int32
	_ = v47
	var v48 int32
	_ = v48
	var v54 int32
	_ = v54
	var v56 int32
	_ = v56
	var v57 int32
	_ = v57
	var v63 int32
	_ = v63
	var v64 int32
	_ = v64
	var v66 int64
	_ = v66
	var v67 int32
	_ = v67
	var v70 int32
	_ = v70
	v6 = m.G0
	v8 = v6 - int32(16)
	m.G0 = v8
	v10 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	if v10 == int32(0) {
		v13 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
		v14 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
		*(*int32)(unsafe.Add(mBase, uint32(v8)+8)) = v14
		v16 = *(*int64)(unsafe.Add(mBase, uint32(l0)))
		*(*int64)(unsafe.Add(mBase, uint32(v8))) = v16
		v18 = F_smgropen(m, v8, v13)
		mBase = m.M
		v21 = m.ExcPending
		if v21 != 0 {
			return int32(0)
		} else {
			*(*int32)(unsafe.Add(mBase, uint32(l0)+12)) = v18
			v24 = *(*int32)(unsafe.Add(mBase, uint32(v18)+72))
			if v24 != 0 {
				v32 = v24
			} else {
				v25 = *(*int32)(unsafe.Add(mBase, uint32(v18)+76))
				v26 = *(*int32)(unsafe.Add(mBase, uint32(v18)+80))
				*(*int32)(unsafe.Add(mBase, uint32(v25)+4)) = v26
				v28 = *(*int32)(unsafe.Add(mBase, uint32(v18)+76))
				*(*int32)(unsafe.Add(mBase, uint32(v26))) = v28
				v30 = *(*int32)(unsafe.Add(mBase, uint32(v18)+72))
				v32 = v30
			}
			*(*int32)(unsafe.Add(mBase, uint32(v18)+72)) = v32 + int32(1)
			v36 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
			v37 = v36
			v38 = *(*int32)(unsafe.Add(mBase, uint32(l0)+48))
			v39 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v38)+118)))
			if v39 != int32(112) {
				v54 = base.B2i32(l1 == int32(3))
			} else {
				v42 = int32(1)
				v44 = *(*int32)(unsafe.Add(mBase, _consts[2]))
				if int32(0) < v44 {
					v54 = v42
				} else {
					v47 = *(*int32)(unsafe.Add(mBase, uint32(l0)+32))
					if v47 != 0 {
						v54 = base.B2i32(l1 == int32(3))
					} else {
						v48 = *(*int32)(unsafe.Add(mBase, uint32(l0)+40))
						if v48 == int32(0) {
							v54 = v42
						} else {
							v54 = base.B2i32(l1 == int32(3))
						}
					}
				}
			}
			v56 = F_palloc(m, int32(424))
			mBase = m.M
			v57 = m.ExcPending
			if v57 != 0 {
				return int32(0)
			} else {
				*(*int32)(unsafe.Add(mBase, uint32(v56)+12)) = int32(0)
				*(*uint8)(unsafe.Add(mBase, uint32(v56)+8)) = uint8(v54)
				*(*int32)(unsafe.Add(mBase, uint32(v56)+4)) = l1
				*(*int32)(unsafe.Add(mBase, uint32(v56))) = v37
				v63 = F_smgrnblocks(m, v37, l1)
				mBase = m.M
				v64 = m.ExcPending
				if v64 != 0 {
					return int32(0)
				} else {
					*(*int32)(unsafe.Add(mBase, uint32(v56)+400)) = v63
					v66 = F_GetRedoRecPtr(m)
					mBase = m.M
					v67 = m.ExcPending
					if v67 != 0 {
						return int32(0)
					} else {
						*(*int64)(unsafe.Add(mBase, uint32(v56)+408)) = v66
						v70 = *(*int32)(unsafe.Add(mBase, _consts[3]))
						*(*int32)(unsafe.Add(mBase, uint32(v56)+416)) = v70
						m.G0 = v8 + int32(16)
						return v56
					}
				}
			}
		}
	} else {
		v37 = v10
		v38 = *(*int32)(unsafe.Add(mBase, uint32(l0)+48))
		v39 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v38)+118)))
		if v39 != int32(112) {
			v54 = base.B2i32(l1 == int32(3))
		} else {
			v42 = int32(1)
			v44 = *(*int32)(unsafe.Add(mBase, _consts[2]))
			if int32(0) < v44 {
				v54 = v42
			} else {
				v47 = *(*int32)(unsafe.Add(mBase, uint32(l0)+32))
				if v47 != 0 {
					v54 = base.B2i32(l1 == int32(3))
				} else {
					v48 = *(*int32)(unsafe.Add(mBase, uint32(l0)+40))
					if v48 == int32(0) {
						v54 = v42
					} else {
						v54 = base.B2i32(l1 == int32(3))
					}
				}
			}
		}
		v56 = F_palloc(m, int32(424))
		mBase = m.M
		v57 = m.ExcPending
		if v57 != 0 {
			return int32(0)
		} else {
			*(*int32)(unsafe.Add(mBase, uint32(v56)+12)) = int32(0)
			*(*uint8)(unsafe.Add(mBase, uint32(v56)+8)) = uint8(v54)
			*(*int32)(unsafe.Add(mBase, uint32(v56)+4)) = l1
			*(*int32)(unsafe.Add(mBase, uint32(v56))) = v37
			v63 = F_smgrnblocks(m, v37, l1)
			mBase = m.M
			v64 = m.ExcPending
			if v64 != 0 {
				return int32(0)
			} else {
				*(*int32)(unsafe.Add(mBase, uint32(v56)+400)) = v63
				v66 = F_GetRedoRecPtr(m)
				mBase = m.M
				v67 = m.ExcPending
				if v67 != 0 {
					return int32(0)
				} else {
					*(*int64)(unsafe.Add(mBase, uint32(v56)+408)) = v66
					v70 = *(*int32)(unsafe.Add(mBase, _consts[3]))
					*(*int32)(unsafe.Add(mBase, uint32(v56)+416)) = v70
					m.G0 = v8 + int32(16)
					return v56
				}
			}
		}
	}
}
