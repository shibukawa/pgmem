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
	var v85 int32
	_ = v85
	var v88 int32
	_ = v88
	var v90 int32
	_ = v90
	var v94 int32
	_ = v94
	var v98 int32
	_ = v98
	var v100 int64
	_ = v100
	var v105 int32
	_ = v105
	var v106 int32
	_ = v106
	var v107 int32
	_ = v107
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
	var v156 int64
	_ = v156
	var v164 int32
	_ = v164
	var v166 int64
	_ = v166
	var v170 int32
	_ = v170
	var v171 int32
	_ = v171
	var v172 int32
	_ = v172
	var v175 int32
	_ = v175
	var v177 int32
	_ = v177
	var v181 int32
	_ = v181
	var v183 int64
	_ = v183
	var v187 int32
	_ = v187
	var v188 int32
	_ = v188
	var v192 int32
	_ = v192
	var v194 int32
	_ = v194
	var v198 int32
	_ = v198
	var v200 int64
	_ = v200
	var v204 int32
	_ = v204
	var v205 int32
	_ = v205
	var v209 int32
	_ = v209
	var v211 int32
	_ = v211
	var v215 int32
	_ = v215
	var v217 int64
	_ = v217
	var v219 int32
	_ = v219
	var v220 int32
	_ = v220
	var v224 int32
	_ = v224
	var v226 int32
	_ = v226
	var v236 int32
	_ = v236
	var v237 int32
	_ = v237
	var v241 int32
	_ = v241
	var v242 int32
	_ = v242
	var v245 int32
	_ = v245
	var v246 int32
	_ = v246
	var v257 int32
	_ = v257
	var v258 int32
	_ = v258
	var v261 int32
	_ = v261
	var v265 int32
	_ = v265
	var v266 int32
	_ = v266
	var v268 int32
	_ = v268
	var v271 int32
	_ = v271
	var v273 int32
	_ = v273
	var v275 int32
	_ = v275
	var v277 int32
	_ = v277
	var v281 int32
	_ = v281
	var v282 int32
	_ = v282
	var v292 int32
	_ = v292
	var v293 int32
	_ = v293
	var v304 int32
	_ = v304
	var v306 int32
	_ = v306
	var v307 int32
	_ = v307
	var v309 int32
	_ = v309
	var v310 int32
	_ = v310
	var v311 int32
	_ = v311
	var v313 int64
	_ = v313
	var v315 int32
	_ = v315
	var v318 int32
	_ = v318
	var v321 int32
	_ = v321
	var v324 int32
	_ = v324
	var v325 int32
	_ = v325
	var v326 int32
	_ = v326
	var v328 int64
	_ = v328
	var v330 int32
	_ = v330
	var v333 int32
	_ = v333
	var v335 int32
	_ = v335
	var v341 int32
	_ = v341
	var v355 int32
	_ = v355
	var v358 int32
	_ = v358
	var v359 int32
	_ = v359
	var v360 int32
	_ = v360
	var v362 int64
	_ = v362
	var v364 int32
	_ = v364
	var v386 int32
	_ = v386
	var v402 int32
	_ = v402
	var v404 int32
	_ = v404
	var v416 int32
	_ = v416
	var v423 int32
	_ = v423
	var v426 int32
	_ = v426
	var v431 int32
	_ = v431
	var v435 int32
	_ = v435
	var v449 int32
	_ = v449
	var v450 int32
	_ = v450
	var v452 int32
	_ = v452
	var v453 int32
	_ = v453
	var v455 int32
	_ = v455
	var v456 int32
	_ = v456
	var v459 int32
	_ = v459
	var v461 int64
	_ = v461
	var v462 int32
	_ = v462
	var v469 int32
	_ = v469
	var v470 int32
	_ = v470
	var v473 int32
	_ = v473
	var v488 int32
	_ = v488
	var v490 int32
	_ = v490
	var v492 int32
	_ = v492
	var v503 int32
	_ = v503
	var v505 int32
	_ = v505
	var v525 int32
	_ = v525
	var v526 int32
	_ = v526
	var v528 int32
	_ = v528
	var v533 int32
	_ = v533
	var v548 int32
	_ = v548
	var v549 int32
	_ = v549
	var v551 int32
	_ = v551
	var v556 int32
	_ = v556
	var v559 int32
	_ = v559
	var v566 int32
	_ = v566
	var v568 int32
	_ = v568
	var v569 int32
	_ = v569
	var v571 int32
	_ = v571
	var v572 int32
	_ = v572
	var v574 int32
	_ = v574
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
	var v689 int32
	_ = v689
	var v704 int32
	_ = v704
	var v705 int32
	_ = v705
	var v707 int32
	_ = v707
	var v708 int32
	_ = v708
	var v709 int32
	_ = v709
	var v712 int32
	_ = v712
	var v713 int32
	_ = v713
	var v715 int32
	_ = v715
	var v716 int32
	_ = v716
	var v717 int32
	_ = v717
	var v720 int32
	_ = v720
	var v721 int32
	_ = v721
	var v723 int32
	_ = v723
	var v724 int32
	_ = v724
	var v725 int32
	_ = v725
	var v728 int32
	_ = v728
	var v729 int32
	_ = v729
	var v731 int32
	_ = v731
	var v732 int32
	_ = v732
	var v733 int32
	_ = v733
	var v736 int32
	_ = v736
	var v738 int32
	_ = v738
	var v754 int32
	_ = v754
	var v756 int32
	_ = v756
	var v775 int32
	_ = v775
	v3 = int32(0)
	v15 = m.G0
	v17 = v15 - int32(96)
	m.G0 = v17
	v20 = *(*int32)(unsafe.Add(mBase, _c_F_smgrDoPendingSyncs[0]))
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
	*(*int32)(unsafe.Add(mBase, _c_F_smgrDoPendingSyncs[0])) = int32(0)
	goto L1
L4:
	;
	goto L5
L5:
	;
	if l1 == int32(0) {
		goto L7
	} else {
		goto L8
	}
L6:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_smgrDoPendingSyncs[0])) = int32(0)
	goto L1
L7:
	;
	v33 = *(*int32)(unsafe.Add(mBase, _c_F_smgrDoPendingSyncs[1]))
	if v33 != 0 {
		goto L10
	} else {
		goto L11
	}
L8:
	;
	goto L9
L9:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_smgrDoPendingSyncs[0])) = int32(0)
	goto L1
L10:
	;
	v34 = v33
	goto L13
L11:
	;
	v74 = v20
	goto L12
L12:
	;
	F_hash_seq_init(m, v17+int32(76), v74)
	mBase = m.M
	v76 = m.ExcPending
	if v76 != 0 {
		goto L18
	} else {
		goto L21
	}
L13:
	;
	v48 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v34)+16)))
	if v48 == int32(1) {
		goto L15
	} else {
		goto L16
	}
L14:
	;
	v59 = *(*int32)(unsafe.Add(mBase, _c_F_smgrDoPendingSyncs[0]))
	v74 = v59
	goto L12
L15:
	;
	v52 = *(*int32)(unsafe.Add(mBase, _c_F_smgrDoPendingSyncs[0]))
	v55 = F_hash_search(m, v52, v34, int32(2), int32(0))
	mBase = m.M
	v56 = m.ExcPending
	if v56 != 0 {
		goto L18
	} else {
		goto L19
	}
L16:
	;
	goto L17
L17:
	;
	v57 = *(*int32)(unsafe.Add(mBase, uint32(v34)+24))
	if v57 != 0 {
		v34 = v57
		goto L13
	} else {
		goto L20
	}
L18:
	;
	return
L19:
	;
	goto L17
L20:
	;
	goto L14
L21:
	;
	v79 = F_hash_seq_search(m, v17+int32(76))
	mBase = m.M
	v80 = m.ExcPending
	if v80 != 0 {
		goto L18
	} else {
		goto L22
	}
L22:
	;
	if v79 == int32(0) {
		goto L6
	} else {
		goto L23
	}
L23:
	;
	v85 = v79
	v88 = int32(0)
	v90 = v3
	v94 = v3
	goto L24
L24:
	;
	v98 = *(*int32)(unsafe.Add(mBase, uint32(v85)+8))
	*(*int32)(unsafe.Add(mBase, uint32(v17)+72)) = v98
	v100 = *(*int64)(unsafe.Add(mBase, uint32(v85)))
	*(*int64)(unsafe.Add(mBase, uint32(v17)+64)) = v100
	v105 = F_smgropen(m, v17-int32(-64), int32(-1))
	mBase = m.M
	v106 = m.ExcPending
	if v106 != 0 {
		goto L18
	} else {
		goto L26
	}
L25:
	;
	v268 = int32(0)
	*(*int32)(unsafe.Add(mBase, _c_F_smgrDoPendingSyncs[0])) = v268
	v271 = base.B2i32(v257 <= v268)
	if v257 <= v268 {
		goto L1
	} else {
		goto L84
	}
L26:
	;
	v107 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v85)+12)))
	if v107 != 0 {
		goto L28
	} else {
		goto L29
	}
L27:
	;
	v265 = F_hash_seq_search(m, v17+int32(76))
	mBase = m.M
	v266 = m.ExcPending
	if v266 != 0 {
		goto L18
	} else {
		goto L82
	}
L28:
	;
	if v90 == int32(0) {
		goto L76
	} else {
		goto L77
	}
L29:
	;
	v109 = int32(-1)
	v112 = F_smgrexists(m, v105, int32(0))
	mBase = m.M
	v113 = m.ExcPending
	if v113 != 0 {
		goto L18
	} else {
		goto L30
	}
L30:
	;
	if v112 != 0 {
		goto L31
	} else {
		goto L32
	}
L31:
	;
	v115 = F_smgrnblocks(m, v105, int32(0))
	mBase = m.M
	v116 = m.ExcPending
	if v116 != 0 {
		goto L18
	} else {
		goto L34
	}
L32:
	;
	v118 = v109
	v119 = int64(0)
	goto L33
L33:
	;
	v121 = F_smgrexists(m, v105, int32(1))
	mBase = m.M
	v122 = m.ExcPending
	if v122 != 0 {
		goto L18
	} else {
		goto L35
	}
L34:
	;
	v118 = v115
	v119 = base.I64_extend_i32_u(v115)
	goto L33
L35:
	;
	if v121 != 0 {
		goto L36
	} else {
		goto L37
	}
L36:
	;
	v124 = F_smgrnblocks(m, v105, int32(1))
	mBase = m.M
	v125 = m.ExcPending
	if v125 != 0 {
		goto L18
	} else {
		goto L39
	}
L37:
	;
	v128 = v109
	v129 = v119
	goto L38
L38:
	;
	v132 = F_smgrexists(m, v105, int32(2))
	mBase = m.M
	v133 = m.ExcPending
	if v133 != 0 {
		goto L18
	} else {
		goto L41
	}
L39:
	;
	v128 = v124
	v129 = v119 + base.I64_extend_i32_u(v124)
	goto L38
L40:
	;
	v145 = F_smgrexists(m, v105, int32(3))
	mBase = m.M
	v146 = m.ExcPending
	if v146 != 0 {
		goto L18
	} else {
		goto L46
	}
L41:
	;
	if v132 == int32(0) {
		goto L42
	} else {
		goto L43
	}
L42:
	;
	v142 = int32(-1)
	v143 = v129
	goto L40
L43:
	;
	goto L44
L44:
	;
	v138 = F_smgrnblocks(m, v105, int32(2))
	mBase = m.M
	v139 = m.ExcPending
	if v139 != 0 {
		goto L18
	} else {
		goto L45
	}
L45:
	;
	v142 = v138
	v143 = v129 + base.I64_extend_i32_u(v138)
	goto L40
L46:
	;
	if v145 != 0 {
		goto L47
	} else {
		goto L48
	}
L47:
	;
	v148 = F_smgrnblocks(m, v105, int32(3))
	mBase = m.M
	v149 = m.ExcPending
	if v149 != 0 {
		goto L18
	} else {
		goto L50
	}
L48:
	;
	v152 = int32(-1)
	v153 = v143
	goto L49
L49:
	;
	v154 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v85)+12)))
	if v154 != 0 {
		goto L28
	} else {
		goto L51
	}
L50:
	;
	v152 = v148
	v153 = v143 + base.I64_extend_i32_u(v148)
	goto L49
L51:
	;
	v156 = int64(*(*int32)(unsafe.Add(mBase, _c_F_smgrDoPendingSyncs[2])))
	if base.Ui64(int64(base.Ui64(v156)>>(uint(int64(3))%64))&int64(2251799813685247)) <= base.Ui64(v153) {
		goto L28
	} else {
		goto L52
	}
L52:
	;
	if v118 != int32(-1) {
		goto L53
	} else {
		goto L54
	}
L53:
	;
	v164 = *(*int32)(unsafe.Add(mBase, uint32(v105)+8))
	*(*int32)(unsafe.Add(mBase, uint32(v17)+56)) = v164
	v166 = *(*int64)(unsafe.Add(mBase, uint32(v105)))
	*(*int64)(unsafe.Add(mBase, uint32(v17)+48)) = v166
	v170 = F_CreateFakeRelcacheEntry(m, v17+int32(48))
	mBase = m.M
	v171 = m.ExcPending
	if v171 != 0 {
		goto L18
	} else {
		goto L56
	}
L54:
	;
	goto L55
L55:
	;
	if v128 != int32(-1) {
		goto L59
	} else {
		goto L60
	}
L56:
	;
	v172 = int32(0)
	F_log_newpage_range(m, v170, v172, v118, v172)
	mBase = m.M
	v175 = m.ExcPending
	if v175 != 0 {
		goto L18
	} else {
		goto L57
	}
L57:
	;
	F_pfree(m, v170)
	mBase = m.M
	v177 = m.ExcPending
	if v177 != 0 {
		goto L18
	} else {
		goto L58
	}
L58:
	;
	goto L55
L59:
	;
	v181 = *(*int32)(unsafe.Add(mBase, uint32(v105)+8))
	*(*int32)(unsafe.Add(mBase, uint32(v17)+40)) = v181
	v183 = *(*int64)(unsafe.Add(mBase, uint32(v105)))
	*(*int64)(unsafe.Add(mBase, uint32(v17)+32)) = v183
	v187 = F_CreateFakeRelcacheEntry(m, v17+int32(32))
	mBase = m.M
	v188 = m.ExcPending
	if v188 != 0 {
		goto L18
	} else {
		goto L62
	}
L60:
	;
	goto L61
L61:
	;
	if v142 != int32(-1) {
		goto L65
	} else {
		goto L66
	}
L62:
	;
	F_log_newpage_range(m, v187, int32(1), v128, int32(0))
	mBase = m.M
	v192 = m.ExcPending
	if v192 != 0 {
		goto L18
	} else {
		goto L63
	}
L63:
	;
	F_pfree(m, v187)
	mBase = m.M
	v194 = m.ExcPending
	if v194 != 0 {
		goto L18
	} else {
		goto L64
	}
L64:
	;
	goto L61
L65:
	;
	v198 = *(*int32)(unsafe.Add(mBase, uint32(v105)+8))
	*(*int32)(unsafe.Add(mBase, uint32(v17)+24)) = v198
	v200 = *(*int64)(unsafe.Add(mBase, uint32(v105)))
	*(*int64)(unsafe.Add(mBase, uint32(v17)+16)) = v200
	v204 = F_CreateFakeRelcacheEntry(m, v17+int32(16))
	mBase = m.M
	v205 = m.ExcPending
	if v205 != 0 {
		goto L18
	} else {
		goto L68
	}
L66:
	;
	goto L67
L67:
	;
	if v152 == int32(-1) {
		v257 = v88
		v258 = v90
		v261 = v94
		goto L27
	} else {
		goto L71
	}
L68:
	;
	F_log_newpage_range(m, v204, int32(2), v142, int32(0))
	mBase = m.M
	v209 = m.ExcPending
	if v209 != 0 {
		goto L18
	} else {
		goto L69
	}
L69:
	;
	F_pfree(m, v204)
	mBase = m.M
	v211 = m.ExcPending
	if v211 != 0 {
		goto L18
	} else {
		goto L70
	}
L70:
	;
	goto L67
L71:
	;
	v215 = *(*int32)(unsafe.Add(mBase, uint32(v105)+8))
	*(*int32)(unsafe.Add(mBase, uint32(v17)+8)) = v215
	v217 = *(*int64)(unsafe.Add(mBase, uint32(v105)))
	*(*int64)(unsafe.Add(mBase, uint32(v17))) = v217
	v219 = F_CreateFakeRelcacheEntry(m, v17)
	mBase = m.M
	v220 = m.ExcPending
	if v220 != 0 {
		goto L18
	} else {
		goto L72
	}
L72:
	;
	F_log_newpage_range(m, v219, int32(3), v152, int32(0))
	mBase = m.M
	v224 = m.ExcPending
	if v224 != 0 {
		goto L18
	} else {
		goto L73
	}
L73:
	;
	F_pfree(m, v219)
	mBase = m.M
	v226 = m.ExcPending
	if v226 != 0 {
		goto L18
	} else {
		goto L74
	}
L74:
	;
	v257 = v88
	v258 = v90
	v261 = v94
	goto L27
L75:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v246+v88<<(uint(int32(2))%32)))) = v105
	v257 = v88 + int32(1)
	v258 = v245
	v261 = v246
	goto L27
L76:
	;
	v236 = F_palloc(m, int32(32))
	mBase = m.M
	v237 = m.ExcPending
	if v237 != 0 {
		goto L18
	} else {
		goto L79
	}
L77:
	;
	goto L78
L78:
	;
	if v88 < v90 {
		v245 = v90
		v246 = v94
		goto L75
	} else {
		goto L80
	}
L79:
	;
	v245 = int32(8)
	v246 = v236
	goto L75
L80:
	;
	v241 = F_repalloc(m, v94, v90<<(uint(int32(3))%32))
	mBase = m.M
	v242 = m.ExcPending
	if v242 != 0 {
		goto L18
	} else {
		goto L81
	}
L81:
	;
	v245 = v90 << (uint(int32(1)) % 32)
	v246 = v241
	goto L75
L82:
	;
	if v265 != 0 {
		v85 = v265
		v88 = v257
		v90 = v258
		v94 = v261
		goto L24
	} else {
		goto L83
	}
L83:
	;
	goto L25
L84:
	;
	if v257 != 0 {
		goto L85
	} else {
		goto L86
	}
L85:
	;
	v273 = int32(0)
	v275 = m.G0
	v277 = v275 - int32(32)
	m.G0 = v277
	if v257 != 0 {
		goto L88
	} else {
		goto L89
	}
L86:
	;
	goto L87
L87:
	;
	F_pfree(m, v261)
	mBase = m.M
	v775 = m.ExcPending
	if v775 != 0 {
		goto L18
	} else {
		goto L182
	}
L88:
	;
	v281 = F_palloc(m, v257<<(uint(int32(4))%32))
	mBase = m.M
	v282 = m.ExcPending
	if v282 != 0 {
		goto L18
	} else {
		goto L91
	}
L89:
	;
	goto L90
L90:
	;
	m.G0 = v277 + int32(32)
	v680 = int32(_a_F_smgrDoPendingSyncs_0)
	v682 = *(*int32)(unsafe.Add(mBase, _c_F_smgrDoPendingSyncs[3]))
	*(*int32)(unsafe.Add(mBase, _c_F_smgrDoPendingSyncs[3])) = v682 + int32(1)
	if int32(0) < v257 {
		goto L156
	} else {
		goto L157
	}
L91:
	;
	if v257 <= v268 {
		v402 = int32(0)
		goto L92
	} else {
		goto L93
	}
L92:
	;
	v404 = *(*int32)(unsafe.Add(mBase, _c_F_smgrDoPendingSyncs[4]))
	if int32(0) < v404 {
		goto L104
	} else {
		goto L105
	}
L93:
	;
	if v257 != int32(1) {
		goto L95
	} else {
		goto L96
	}
L94:
	;
	if base.Ui32(v257) < base.Ui32(int32(21)) {
		v402 = int32(0)
		goto L92
	} else {
		goto L102
	}
L95:
	;
	v292 = v273
	v293 = v273
	goto L98
L96:
	;
	v341 = v273
	goto L97
L97:
	;
	v355 = v281 + v341<<(uint(int32(4))%32)
	v358 = v261 + v341<<(uint(int32(2))%32)
	v359 = *(*int32)(unsafe.Add(mBase, uint32(v358)))
	v360 = *(*int32)(unsafe.Add(mBase, uint32(v359)+8))
	*(*int32)(unsafe.Add(mBase, uint32(v355)+8)) = v360
	v362 = *(*int64)(unsafe.Add(mBase, uint32(v359)))
	*(*int64)(unsafe.Add(mBase, uint32(v355))) = v362
	v364 = *(*int32)(unsafe.Add(mBase, uint32(v358)))
	*(*int32)(unsafe.Add(mBase, uint32(v355)+12)) = v364
	goto L94
L98:
	;
	v304 = int32(4)
	v306 = v281 + v292<<(uint(v304)%32)
	v307 = int32(2)
	v309 = v261 + v292<<(uint(v307)%32)
	v310 = *(*int32)(unsafe.Add(mBase, uint32(v309)))
	v311 = *(*int32)(unsafe.Add(mBase, uint32(v310)+8))
	*(*int32)(unsafe.Add(mBase, uint32(v306)+8)) = v311
	v313 = *(*int64)(unsafe.Add(mBase, uint32(v310)))
	*(*int64)(unsafe.Add(mBase, uint32(v306))) = v313
	v315 = *(*int32)(unsafe.Add(mBase, uint32(v309)))
	*(*int32)(unsafe.Add(mBase, uint32(v306)+12)) = v315
	v318 = v292 | int32(1)
	v321 = v281 + v318<<(uint(v304)%32)
	v324 = v261 + v318<<(uint(v307)%32)
	v325 = *(*int32)(unsafe.Add(mBase, uint32(v324)))
	v326 = *(*int32)(unsafe.Add(mBase, uint32(v325)+8))
	*(*int32)(unsafe.Add(mBase, uint32(v321)+8)) = v326
	v328 = *(*int64)(unsafe.Add(mBase, uint32(v325)))
	*(*int64)(unsafe.Add(mBase, uint32(v321))) = v328
	v330 = *(*int32)(unsafe.Add(mBase, uint32(v324)))
	*(*int32)(unsafe.Add(mBase, uint32(v321)+12)) = v330
	v333 = v292 + v307
	v335 = v293 + v307
	if v335 != v257&int32(2147483646) {
		v292 = v333
		v293 = v335
		goto L98
	} else {
		goto L100
	}
L99:
	;
	if v257&int32(1) == int32(0) {
		goto L94
	} else {
		goto L101
	}
L100:
	;
	goto L99
L101:
	;
	v341 = v333
	goto L97
L102:
	;
	F_pg_qsort(m, v281, v257, int32(16), int32(1078))
	mBase = m.M
	v386 = m.ExcPending
	if v386 != 0 {
		goto L18
	} else {
		goto L103
	}
L103:
	;
	v402 = int32(1)
	goto L92
L104:
	;
	v416 = int32(0)
	goto L107
L105:
	;
	goto L106
L106:
	;
	F_pfree(m, v281)
	mBase = m.M
	v662 = m.ExcPending
	if v662 != 0 {
		goto L18
	} else {
		goto L155
	}
L107:
	;
	v423 = *(*int32)(unsafe.Add(mBase, _c_F_smgrDoPendingSyncs[5]))
	v426 = v423 + v416<<(uint(int32(6))%32)
	if v402 == int32(0) {
		goto L111
	} else {
		goto L112
	}
L108:
	;
	goto L106
L109:
	;
	v643 = v416 + int32(1)
	v645 = *(*int32)(unsafe.Add(mBase, _c_F_smgrDoPendingSyncs[4]))
	if v643 < v645 {
		v416 = v643
		goto L107
	} else {
		goto L154
	}
L110:
	;
	F_ReservePrivateRefCountEntry(m)
	mBase = m.M
	v488 = m.ExcPending
	if v488 != 0 {
		goto L18
	} else {
		goto L124
	}
L111:
	;
	if v257 <= int32(0) {
		goto L109
	} else {
		goto L114
	}
L112:
	;
	goto L113
L113:
	;
	v461 = *(*int64)(unsafe.Add(mBase, uint32(v426)))
	v462 = *(*int32)(unsafe.Add(mBase, uint32(v426)+8))
	*(*int32)(unsafe.Add(mBase, uint32(v277)+16)) = v462
	*(*int64)(unsafe.Add(mBase, uint32(v277)+8)) = v461
	v469 = F_bsearch(m, v277+int32(8), v281, v257, int32(16), int32(1078))
	mBase = m.M
	v470 = m.ExcPending
	if v470 != 0 {
		goto L18
	} else {
		goto L122
	}
L114:
	;
	v431 = *(*int32)(unsafe.Add(mBase, uint32(v426)))
	v435 = int32(0)
	goto L115
L115:
	;
	v449 = v281 + v435<<(uint(int32(4))%32)
	v450 = *(*int32)(unsafe.Add(mBase, uint32(v449)))
	if v431 != v450 {
		goto L117
	} else {
		goto L118
	}
L116:
	;
	goto L109
L117:
	;
	v459 = v435 + int32(1)
	if v459 != v257 {
		v435 = v459
		goto L115
	} else {
		goto L121
	}
L118:
	;
	v452 = *(*int32)(unsafe.Add(mBase, uint32(v426)+4))
	v453 = *(*int32)(unsafe.Add(mBase, uint32(v449)+4))
	if v452 != v453 {
		goto L117
	} else {
		goto L119
	}
L119:
	;
	v455 = *(*int32)(unsafe.Add(mBase, uint32(v426)+8))
	v456 = *(*int32)(unsafe.Add(mBase, uint32(v449)+8))
	if v455 == v456 {
		v473 = v449
		goto L110
	} else {
		goto L120
	}
L120:
	;
	goto L117
L121:
	;
	goto L116
L122:
	;
	if v469 == int32(0) {
		goto L109
	} else {
		goto L123
	}
L123:
	;
	v473 = v469
	goto L110
L124:
	;
	v490 = *(*int32)(unsafe.Add(mBase, _c_F_smgrDoPendingSyncs[6]))
	F_ResourceOwnerEnlarge(m, v490)
	mBase = m.M
	v492 = m.ExcPending
	if v492 != 0 {
		goto L18
	} else {
		goto L125
	}
L125:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v277)+28)) = int32(_a_F_smgrDoPendingSyncs_1)
	*(*int32)(unsafe.Add(mBase, uint32(v277)+24)) = int32(_a_F_smgrDoPendingSyncs_2)
	*(*int32)(unsafe.Add(mBase, uint32(v277)+20)) = int32(_a_F_smgrDoPendingSyncs_3)
	*(*int32)(unsafe.Add(mBase, uint32(v277)+16)) = int32(0)
	*(*int64)(unsafe.Add(mBase, uint32(v277)+8)) = int64(0)
	v503 = int32(_a_F_smgrDoPendingSyncs_4)
	v505 = base.AtomicRmwOr32(m, v426, int32(24), v503)
	if v505&v503 != 0 {
		goto L126
	} else {
		goto L127
	}
L126:
	;
	goto L129
L127:
	;
	v533 = v505
	goto L128
L128:
	;
	v548 = int32(_a_F_smgrDoPendingSyncs_5)
	v549 = *(*int32)(unsafe.Add(mBase, _c_F_smgrDoPendingSyncs[7]))
	v551 = *(*int32)(unsafe.Add(mBase, uint32(v277+int32(8))+8))
	if v551 == int32(0) {
		goto L136
	} else {
		goto L137
	}
L129:
	;
	F_perform_spin_delay(m, v277+int32(8))
	mBase = m.M
	v525 = m.ExcPending
	if v525 != 0 {
		goto L18
	} else {
		goto L131
	}
L130:
	;
	v533 = v528
	goto L128
L131:
	;
	v526 = int32(_a_F_smgrDoPendingSyncs_4)
	v528 = base.AtomicRmwOr32(m, v426, int32(24), v526)
	if v528&v526 != 0 {
		goto L129
	} else {
		goto L132
	}
L132:
	;
	goto L130
L133:
	;
	v568 = *(*int32)(unsafe.Add(mBase, uint32(v426)))
	v569 = *(*int32)(unsafe.Add(mBase, uint32(v473)))
	if v568 != v569 {
		goto L144
	} else {
		goto L145
	}
L134:
	;
	goto L133
L135:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_smgrDoPendingSyncs[7])) = v566
	goto L134
L136:
	;
	if int32(999) < v549 {
		goto L134
	} else {
		goto L139
	}
L137:
	;
	goto L138
L138:
	;
	if v549 < int32(11) {
		goto L134
	} else {
		goto L143
	}
L139:
	;
	v556 = int32(900)
	if v556 <= v549 {
		goto L140
	} else {
		goto L141
	}
L140:
	;
	v559 = v556
	goto L142
L141:
	;
	v559 = v549
	goto L142
L142:
	;
	v566 = v559 + int32(100)
	goto L135
L143:
	;
	v566 = v549 - int32(1)
	goto L135
L144:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v426)+24)) = v533 & int32(-4194305)
	goto L109
L145:
	;
	v571 = *(*int32)(unsafe.Add(mBase, uint32(v426)+4))
	v572 = *(*int32)(unsafe.Add(mBase, uint32(v473)+4))
	v574 = int32(25165824)
	if base.B2i32(v571 != v572)|base.B2i32(v533&v574 != v574) != 0 {
		goto L144
	} else {
		goto L146
	}
L146:
	;
	v579 = *(*int32)(unsafe.Add(mBase, uint32(v426)+8))
	v580 = *(*int32)(unsafe.Add(mBase, uint32(v473)+8))
	if v579 != v580 {
		goto L144
	} else {
		goto L147
	}
L147:
	;
	v582 = *(*int32)(unsafe.Add(mBase, uint32(v426)+24))
	v583 = int32(1)
	*(*int32)(unsafe.Add(mBase, uint32(v426)+24)) = (v582 + v583) & int32(-4194305)
	v588 = *(*int32)(unsafe.Add(mBase, uint32(v426)+20))
	v589 = int32(_a_F_smgrDoPendingSyncs_6)
	v590 = *(*int32)(unsafe.Add(mBase, _c_F_smgrDoPendingSyncs[8]))
	*(*int32)(unsafe.Add(mBase, _c_F_smgrDoPendingSyncs[8])) = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v590)+4)) = v583
	v597 = v588 + v583
	*(*int32)(unsafe.Add(mBase, uint32(v590))) = v597
	v600 = *(*int32)(unsafe.Add(mBase, _c_F_smgrDoPendingSyncs[6]))
	F_ResourceOwnerRemember(m, v600, v597, int32(_a_F_smgrDoPendingSyncs_7))
	mBase = m.M
	v603 = m.ExcPending
	if v603 != 0 {
		goto L18
	} else {
		goto L148
	}
L148:
	;
	v605 = v426 + int32(48)
	v607 = F_LWLockAcquire(m, v605, int32(1))
	mBase = m.M
	v608 = m.ExcPending
	if v608 != 0 {
		goto L18
	} else {
		goto L149
	}
L149:
	;
	v609 = *(*int32)(unsafe.Add(mBase, uint32(v473)+12))
	F_FlushBuffer(m, v426, v609, int32(3))
	mBase = m.M
	v612 = m.ExcPending
	if v612 != 0 {
		goto L18
	} else {
		goto L150
	}
L150:
	;
	F_LWLockRelease(m, v605)
	mBase = m.M
	v614 = m.ExcPending
	if v614 != 0 {
		goto L18
	} else {
		goto L151
	}
L151:
	;
	v616 = *(*int32)(unsafe.Add(mBase, _c_F_smgrDoPendingSyncs[6]))
	v617 = *(*int32)(unsafe.Add(mBase, uint32(v426)+20))
	F_ResourceOwnerForget(m, v616, v617+int32(1), int32(_a_F_smgrDoPendingSyncs_7))
	mBase = m.M
	v622 = m.ExcPending
	if v622 != 0 {
		goto L18
	} else {
		goto L152
	}
L152:
	;
	F_UnpinBufferNoOwner(m, v426)
	mBase = m.M
	v624 = m.ExcPending
	if v624 != 0 {
		goto L18
	} else {
		goto L153
	}
L153:
	;
	goto L109
L154:
	;
	goto L108
L155:
	;
	goto L90
L156:
	;
	v689 = int32(0)
	goto L159
L157:
	;
	goto L158
L158:
	;
	v754 = int32(_a_F_smgrDoPendingSyncs_0)
	v756 = *(*int32)(unsafe.Add(mBase, _c_F_smgrDoPendingSyncs[3]))
	*(*int32)(unsafe.Add(mBase, _c_F_smgrDoPendingSyncs[3])) = v756 - int32(1)
	goto L87
L159:
	;
	v704 = v261 + v689<<(uint(int32(2))%32)
	v705 = *(*int32)(unsafe.Add(mBase, uint32(v704)))
	v707 = F_mdexists(m, v705, int32(0))
	mBase = m.M
	v708 = m.ExcPending
	if v708 != 0 {
		goto L18
	} else {
		goto L161
	}
L160:
	;
	goto L158
L161:
	;
	if v707 != 0 {
		goto L162
	} else {
		goto L163
	}
L162:
	;
	v709 = *(*int32)(unsafe.Add(mBase, uint32(v704)))
	F_mdimmedsync(m, v709, int32(0))
	mBase = m.M
	v712 = m.ExcPending
	if v712 != 0 {
		goto L18
	} else {
		goto L165
	}
L163:
	;
	goto L164
L164:
	;
	v713 = *(*int32)(unsafe.Add(mBase, uint32(v704)))
	v715 = F_mdexists(m, v713, int32(1))
	mBase = m.M
	v716 = m.ExcPending
	if v716 != 0 {
		goto L18
	} else {
		goto L166
	}
L165:
	;
	goto L164
L166:
	;
	if v715 != 0 {
		goto L167
	} else {
		goto L168
	}
L167:
	;
	v717 = *(*int32)(unsafe.Add(mBase, uint32(v704)))
	F_mdimmedsync(m, v717, int32(1))
	mBase = m.M
	v720 = m.ExcPending
	if v720 != 0 {
		goto L18
	} else {
		goto L170
	}
L168:
	;
	goto L169
L169:
	;
	v721 = *(*int32)(unsafe.Add(mBase, uint32(v704)))
	v723 = F_mdexists(m, v721, int32(2))
	mBase = m.M
	v724 = m.ExcPending
	if v724 != 0 {
		goto L18
	} else {
		goto L171
	}
L170:
	;
	goto L169
L171:
	;
	if v723 != 0 {
		goto L172
	} else {
		goto L173
	}
L172:
	;
	v725 = *(*int32)(unsafe.Add(mBase, uint32(v704)))
	F_mdimmedsync(m, v725, int32(2))
	mBase = m.M
	v728 = m.ExcPending
	if v728 != 0 {
		goto L18
	} else {
		goto L175
	}
L173:
	;
	goto L174
L174:
	;
	v729 = *(*int32)(unsafe.Add(mBase, uint32(v704)))
	v731 = F_mdexists(m, v729, int32(3))
	mBase = m.M
	v732 = m.ExcPending
	if v732 != 0 {
		goto L18
	} else {
		goto L176
	}
L175:
	;
	goto L174
L176:
	;
	if v731 != 0 {
		goto L177
	} else {
		goto L178
	}
L177:
	;
	v733 = *(*int32)(unsafe.Add(mBase, uint32(v704)))
	F_mdimmedsync(m, v733, int32(3))
	mBase = m.M
	v736 = m.ExcPending
	if v736 != 0 {
		goto L18
	} else {
		goto L180
	}
L178:
	;
	goto L179
L179:
	;
	v738 = v689 + int32(1)
	if v738 != v257 {
		v689 = v738
		goto L159
	} else {
		goto L181
	}
L180:
	;
	goto L179
L181:
	;
	goto L160
L182:
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
	v6 = F_MemoryContextAllocAligned(m, v2, int32(_a_F_smgr_bulk_get_buf_0), int32(_a_F_smgr_bulk_get_buf_1), int32(0))
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
				v44 = *(*int32)(unsafe.Add(mBase, _c_F_smgr_bulk_start_rel[0]))
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
						v70 = *(*int32)(unsafe.Add(mBase, _c_F_smgr_bulk_start_rel[1]))
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
			v44 = *(*int32)(unsafe.Add(mBase, _c_F_smgr_bulk_start_rel[0]))
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
					v70 = *(*int32)(unsafe.Add(mBase, _c_F_smgr_bulk_start_rel[1]))
					*(*int32)(unsafe.Add(mBase, uint32(v56)+416)) = v70
					m.G0 = v8 + int32(16)
					return v56
				}
			}
		}
	}
}
