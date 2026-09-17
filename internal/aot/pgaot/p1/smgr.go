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
	var v84 int32
	_ = v84
	var v88 int32
	_ = v88
	var v89 int32
	_ = v89
	var v93 int32
	_ = v93
	var v97 int32
	_ = v97
	var v99 int64
	_ = v99
	var v104 int32
	_ = v104
	var v105 int32
	_ = v105
	var v106 int32
	_ = v106
	var v108 int32
	_ = v108
	var v111 int32
	_ = v111
	var v112 int32
	_ = v112
	var v114 int32
	_ = v114
	var v115 int32
	_ = v115
	var v117 int32
	_ = v117
	var v118 int64
	_ = v118
	var v120 int32
	_ = v120
	var v121 int32
	_ = v121
	var v123 int32
	_ = v123
	var v124 int32
	_ = v124
	var v127 int32
	_ = v127
	var v128 int64
	_ = v128
	var v131 int32
	_ = v131
	var v132 int32
	_ = v132
	var v137 int32
	_ = v137
	var v138 int32
	_ = v138
	var v141 int32
	_ = v141
	var v142 int64
	_ = v142
	var v144 int32
	_ = v144
	var v145 int32
	_ = v145
	var v147 int32
	_ = v147
	var v148 int32
	_ = v148
	var v151 int32
	_ = v151
	var v152 int64
	_ = v152
	var v153 int32
	_ = v153
	var v155 int64
	_ = v155
	var v163 int32
	_ = v163
	var v165 int64
	_ = v165
	var v169 int32
	_ = v169
	var v170 int32
	_ = v170
	var v171 int32
	_ = v171
	var v174 int32
	_ = v174
	var v176 int32
	_ = v176
	var v180 int32
	_ = v180
	var v182 int64
	_ = v182
	var v186 int32
	_ = v186
	var v187 int32
	_ = v187
	var v191 int32
	_ = v191
	var v193 int32
	_ = v193
	var v197 int32
	_ = v197
	var v199 int64
	_ = v199
	var v203 int32
	_ = v203
	var v204 int32
	_ = v204
	var v208 int32
	_ = v208
	var v210 int32
	_ = v210
	var v214 int32
	_ = v214
	var v216 int64
	_ = v216
	var v218 int32
	_ = v218
	var v219 int32
	_ = v219
	var v223 int32
	_ = v223
	var v225 int32
	_ = v225
	var v235 int32
	_ = v235
	var v236 int32
	_ = v236
	var v240 int32
	_ = v240
	var v241 int32
	_ = v241
	var v244 int32
	_ = v244
	var v245 int32
	_ = v245
	var v256 int32
	_ = v256
	var v257 int32
	_ = v257
	var v260 int32
	_ = v260
	var v264 int32
	_ = v264
	var v265 int32
	_ = v265
	var v267 int32
	_ = v267
	var v270 int32
	_ = v270
	var v272 int32
	_ = v272
	var v274 int32
	_ = v274
	var v276 int32
	_ = v276
	var v280 int32
	_ = v280
	var v281 int32
	_ = v281
	var v291 int32
	_ = v291
	var v292 int32
	_ = v292
	var v303 int32
	_ = v303
	var v305 int32
	_ = v305
	var v306 int32
	_ = v306
	var v308 int32
	_ = v308
	var v309 int32
	_ = v309
	var v310 int32
	_ = v310
	var v312 int64
	_ = v312
	var v314 int32
	_ = v314
	var v317 int32
	_ = v317
	var v320 int32
	_ = v320
	var v323 int32
	_ = v323
	var v324 int32
	_ = v324
	var v325 int32
	_ = v325
	var v327 int64
	_ = v327
	var v329 int32
	_ = v329
	var v332 int32
	_ = v332
	var v334 int32
	_ = v334
	var v340 int32
	_ = v340
	var v354 int32
	_ = v354
	var v357 int32
	_ = v357
	var v358 int32
	_ = v358
	var v359 int32
	_ = v359
	var v361 int64
	_ = v361
	var v363 int32
	_ = v363
	var v385 int32
	_ = v385
	var v401 int32
	_ = v401
	var v403 int32
	_ = v403
	var v415 int32
	_ = v415
	var v422 int32
	_ = v422
	var v425 int32
	_ = v425
	var v430 int32
	_ = v430
	var v434 int32
	_ = v434
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
	var v472 int32
	_ = v472
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
	var v534 int32
	_ = v534
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
	var v580 int32
	_ = v580
	var v581 int32
	_ = v581
	var v583 int32
	_ = v583
	var v584 int32
	_ = v584
	var v589 int32
	_ = v589
	var v590 int32
	_ = v590
	var v591 int32
	_ = v591
	var v598 int32
	_ = v598
	var v601 int32
	_ = v601
	var v604 int32
	_ = v604
	var v606 int32
	_ = v606
	var v608 int32
	_ = v608
	var v609 int32
	_ = v609
	var v610 int32
	_ = v610
	var v613 int32
	_ = v613
	var v615 int32
	_ = v615
	var v617 int32
	_ = v617
	var v618 int32
	_ = v618
	var v623 int32
	_ = v623
	var v625 int32
	_ = v625
	var v644 int32
	_ = v644
	var v646 int32
	_ = v646
	var v663 int32
	_ = v663
	var v681 int32
	_ = v681
	var v683 int32
	_ = v683
	var v690 int32
	_ = v690
	var v705 int32
	_ = v705
	var v706 int32
	_ = v706
	var v708 int32
	_ = v708
	var v709 int32
	_ = v709
	var v710 int32
	_ = v710
	var v713 int32
	_ = v713
	var v714 int32
	_ = v714
	var v716 int32
	_ = v716
	var v717 int32
	_ = v717
	var v718 int32
	_ = v718
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
	var v729 int32
	_ = v729
	var v730 int32
	_ = v730
	var v732 int32
	_ = v732
	var v733 int32
	_ = v733
	var v734 int32
	_ = v734
	var v737 int32
	_ = v737
	var v739 int32
	_ = v739
	var v755 int32
	_ = v755
	var v757 int32
	_ = v757
	var v776 int32
	_ = v776
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
	v84 = v79
	v88 = v3
	v89 = v3
	v93 = v3
	goto L24
L24:
	;
	v97 = *(*int32)(unsafe.Add(mBase, uint32(v84)+8))
	*(*int32)(unsafe.Add(mBase, uint32(v17)+72)) = v97
	v99 = *(*int64)(unsafe.Add(mBase, uint32(v84)))
	*(*int64)(unsafe.Add(mBase, uint32(v17)+64)) = v99
	v104 = F_smgropen(m, v17-int32(-64), int32(-1))
	mBase = m.M
	v105 = m.ExcPending
	if v105 != 0 {
		goto L18
	} else {
		goto L26
	}
L25:
	;
	v267 = int32(0)
	*(*int32)(unsafe.Add(mBase, _c_F_smgrDoPendingSyncs[0])) = v267
	v270 = base.B2i32(v256 <= v267)
	if v256 <= v267 {
		goto L1
	} else {
		goto L84
	}
L26:
	;
	v106 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v84)+12)))
	if v106 != 0 {
		goto L28
	} else {
		goto L29
	}
L27:
	;
	v264 = F_hash_seq_search(m, v17+int32(76))
	mBase = m.M
	v265 = m.ExcPending
	if v265 != 0 {
		goto L18
	} else {
		goto L82
	}
L28:
	;
	if v89 == int32(0) {
		goto L76
	} else {
		goto L77
	}
L29:
	;
	v108 = int32(-1)
	v111 = F_smgrexists(m, v104, int32(0))
	mBase = m.M
	v112 = m.ExcPending
	if v112 != 0 {
		goto L18
	} else {
		goto L30
	}
L30:
	;
	if v111 != 0 {
		goto L31
	} else {
		goto L32
	}
L31:
	;
	v114 = F_smgrnblocks(m, v104, int32(0))
	mBase = m.M
	v115 = m.ExcPending
	if v115 != 0 {
		goto L18
	} else {
		goto L34
	}
L32:
	;
	v117 = v108
	v118 = int64(0)
	goto L33
L33:
	;
	v120 = F_smgrexists(m, v104, int32(1))
	mBase = m.M
	v121 = m.ExcPending
	if v121 != 0 {
		goto L18
	} else {
		goto L35
	}
L34:
	;
	v117 = v114
	v118 = base.I64_extend_i32_u(v114)
	goto L33
L35:
	;
	if v120 != 0 {
		goto L36
	} else {
		goto L37
	}
L36:
	;
	v123 = F_smgrnblocks(m, v104, int32(1))
	mBase = m.M
	v124 = m.ExcPending
	if v124 != 0 {
		goto L18
	} else {
		goto L39
	}
L37:
	;
	v127 = v108
	v128 = v118
	goto L38
L38:
	;
	v131 = F_smgrexists(m, v104, int32(2))
	mBase = m.M
	v132 = m.ExcPending
	if v132 != 0 {
		goto L18
	} else {
		goto L41
	}
L39:
	;
	v127 = v123
	v128 = v118 + base.I64_extend_i32_u(v123)
	goto L38
L40:
	;
	v144 = F_smgrexists(m, v104, int32(3))
	mBase = m.M
	v145 = m.ExcPending
	if v145 != 0 {
		goto L18
	} else {
		goto L46
	}
L41:
	;
	if v131 == int32(0) {
		goto L42
	} else {
		goto L43
	}
L42:
	;
	v141 = int32(-1)
	v142 = v128
	goto L40
L43:
	;
	goto L44
L44:
	;
	v137 = F_smgrnblocks(m, v104, int32(2))
	mBase = m.M
	v138 = m.ExcPending
	if v138 != 0 {
		goto L18
	} else {
		goto L45
	}
L45:
	;
	v141 = v137
	v142 = v128 + base.I64_extend_i32_u(v137)
	goto L40
L46:
	;
	if v144 != 0 {
		goto L47
	} else {
		goto L48
	}
L47:
	;
	v147 = F_smgrnblocks(m, v104, int32(3))
	mBase = m.M
	v148 = m.ExcPending
	if v148 != 0 {
		goto L18
	} else {
		goto L50
	}
L48:
	;
	v151 = int32(-1)
	v152 = v142
	goto L49
L49:
	;
	v153 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v84)+12)))
	if v153 != 0 {
		goto L28
	} else {
		goto L51
	}
L50:
	;
	v151 = v147
	v152 = v142 + base.I64_extend_i32_u(v147)
	goto L49
L51:
	;
	v155 = int64(*(*int32)(unsafe.Add(mBase, _c_F_smgrDoPendingSyncs[2])))
	if base.Ui64(int64(base.Ui64(v155)>>(uint(int64(3))%64))&int64(2251799813685247)) <= base.Ui64(v152) {
		goto L28
	} else {
		goto L52
	}
L52:
	;
	if v117 != int32(-1) {
		goto L53
	} else {
		goto L54
	}
L53:
	;
	v163 = *(*int32)(unsafe.Add(mBase, uint32(v104)+8))
	*(*int32)(unsafe.Add(mBase, uint32(v17)+56)) = v163
	v165 = *(*int64)(unsafe.Add(mBase, uint32(v104)))
	*(*int64)(unsafe.Add(mBase, uint32(v17)+48)) = v165
	v169 = F_CreateFakeRelcacheEntry(m, v17+int32(48))
	mBase = m.M
	v170 = m.ExcPending
	if v170 != 0 {
		goto L18
	} else {
		goto L56
	}
L54:
	;
	goto L55
L55:
	;
	if v127 != int32(-1) {
		goto L59
	} else {
		goto L60
	}
L56:
	;
	v171 = int32(0)
	F_log_newpage_range(m, v169, v171, v117, v171)
	mBase = m.M
	v174 = m.ExcPending
	if v174 != 0 {
		goto L18
	} else {
		goto L57
	}
L57:
	;
	F_pfree(m, v169)
	mBase = m.M
	v176 = m.ExcPending
	if v176 != 0 {
		goto L18
	} else {
		goto L58
	}
L58:
	;
	goto L55
L59:
	;
	v180 = *(*int32)(unsafe.Add(mBase, uint32(v104)+8))
	*(*int32)(unsafe.Add(mBase, uint32(v17)+40)) = v180
	v182 = *(*int64)(unsafe.Add(mBase, uint32(v104)))
	*(*int64)(unsafe.Add(mBase, uint32(v17)+32)) = v182
	v186 = F_CreateFakeRelcacheEntry(m, v17+int32(32))
	mBase = m.M
	v187 = m.ExcPending
	if v187 != 0 {
		goto L18
	} else {
		goto L62
	}
L60:
	;
	goto L61
L61:
	;
	if v141 != int32(-1) {
		goto L65
	} else {
		goto L66
	}
L62:
	;
	F_log_newpage_range(m, v186, int32(1), v127, int32(0))
	mBase = m.M
	v191 = m.ExcPending
	if v191 != 0 {
		goto L18
	} else {
		goto L63
	}
L63:
	;
	F_pfree(m, v186)
	mBase = m.M
	v193 = m.ExcPending
	if v193 != 0 {
		goto L18
	} else {
		goto L64
	}
L64:
	;
	goto L61
L65:
	;
	v197 = *(*int32)(unsafe.Add(mBase, uint32(v104)+8))
	*(*int32)(unsafe.Add(mBase, uint32(v17)+24)) = v197
	v199 = *(*int64)(unsafe.Add(mBase, uint32(v104)))
	*(*int64)(unsafe.Add(mBase, uint32(v17)+16)) = v199
	v203 = F_CreateFakeRelcacheEntry(m, v17+int32(16))
	mBase = m.M
	v204 = m.ExcPending
	if v204 != 0 {
		goto L18
	} else {
		goto L68
	}
L66:
	;
	goto L67
L67:
	;
	if v151 == int32(-1) {
		v256 = v88
		v257 = v89
		v260 = v93
		goto L27
	} else {
		goto L71
	}
L68:
	;
	F_log_newpage_range(m, v203, int32(2), v141, int32(0))
	mBase = m.M
	v208 = m.ExcPending
	if v208 != 0 {
		goto L18
	} else {
		goto L69
	}
L69:
	;
	F_pfree(m, v203)
	mBase = m.M
	v210 = m.ExcPending
	if v210 != 0 {
		goto L18
	} else {
		goto L70
	}
L70:
	;
	goto L67
L71:
	;
	v214 = *(*int32)(unsafe.Add(mBase, uint32(v104)+8))
	*(*int32)(unsafe.Add(mBase, uint32(v17)+8)) = v214
	v216 = *(*int64)(unsafe.Add(mBase, uint32(v104)))
	*(*int64)(unsafe.Add(mBase, uint32(v17))) = v216
	v218 = F_CreateFakeRelcacheEntry(m, v17)
	mBase = m.M
	v219 = m.ExcPending
	if v219 != 0 {
		goto L18
	} else {
		goto L72
	}
L72:
	;
	F_log_newpage_range(m, v218, int32(3), v151, int32(0))
	mBase = m.M
	v223 = m.ExcPending
	if v223 != 0 {
		goto L18
	} else {
		goto L73
	}
L73:
	;
	F_pfree(m, v218)
	mBase = m.M
	v225 = m.ExcPending
	if v225 != 0 {
		goto L18
	} else {
		goto L74
	}
L74:
	;
	v256 = v88
	v257 = v89
	v260 = v93
	goto L27
L75:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v245+v88<<(uint(int32(2))%32)))) = v104
	v256 = v88 + int32(1)
	v257 = v244
	v260 = v245
	goto L27
L76:
	;
	v235 = F_palloc(m, int32(32))
	mBase = m.M
	v236 = m.ExcPending
	if v236 != 0 {
		goto L18
	} else {
		goto L79
	}
L77:
	;
	goto L78
L78:
	;
	if v88 < v89 {
		v244 = v89
		v245 = v93
		goto L75
	} else {
		goto L80
	}
L79:
	;
	v244 = int32(8)
	v245 = v235
	goto L75
L80:
	;
	v240 = F_repalloc(m, v93, v89<<(uint(int32(3))%32))
	mBase = m.M
	v241 = m.ExcPending
	if v241 != 0 {
		goto L18
	} else {
		goto L81
	}
L81:
	;
	v244 = v89 << (uint(int32(1)) % 32)
	v245 = v240
	goto L75
L82:
	;
	if v264 != 0 {
		v84 = v264
		v88 = v256
		v89 = v257
		v93 = v260
		goto L24
	} else {
		goto L83
	}
L83:
	;
	goto L25
L84:
	;
	if v256 != 0 {
		goto L85
	} else {
		goto L86
	}
L85:
	;
	v272 = int32(0)
	v274 = m.G0
	v276 = v274 - int32(32)
	m.G0 = v276
	if v256 != 0 {
		goto L88
	} else {
		goto L89
	}
L86:
	;
	goto L87
L87:
	;
	F_pfree(m, v260)
	mBase = m.M
	v776 = m.ExcPending
	if v776 != 0 {
		goto L18
	} else {
		goto L182
	}
L88:
	;
	v280 = F_palloc(m, v256<<(uint(int32(4))%32))
	mBase = m.M
	v281 = m.ExcPending
	if v281 != 0 {
		goto L18
	} else {
		goto L91
	}
L89:
	;
	goto L90
L90:
	;
	m.G0 = v276 + int32(32)
	v681 = int32(_a_F_smgrDoPendingSyncs_0)
	v683 = *(*int32)(unsafe.Add(mBase, _c_F_smgrDoPendingSyncs[3]))
	*(*int32)(unsafe.Add(mBase, _c_F_smgrDoPendingSyncs[3])) = v683 + int32(1)
	if int32(0) < v256 {
		goto L156
	} else {
		goto L157
	}
L91:
	;
	if v256 <= v267 {
		v401 = int32(0)
		goto L92
	} else {
		goto L93
	}
L92:
	;
	v403 = *(*int32)(unsafe.Add(mBase, _c_F_smgrDoPendingSyncs[4]))
	if int32(0) < v403 {
		goto L104
	} else {
		goto L105
	}
L93:
	;
	if v256 != int32(1) {
		goto L95
	} else {
		goto L96
	}
L94:
	;
	if base.Ui32(v256) < base.Ui32(int32(21)) {
		v401 = int32(0)
		goto L92
	} else {
		goto L102
	}
L95:
	;
	v291 = v272
	v292 = v272
	goto L98
L96:
	;
	v340 = v272
	goto L97
L97:
	;
	v354 = v280 + v340<<(uint(int32(4))%32)
	v357 = v260 + v340<<(uint(int32(2))%32)
	v358 = *(*int32)(unsafe.Add(mBase, uint32(v357)))
	v359 = *(*int32)(unsafe.Add(mBase, uint32(v358)+8))
	*(*int32)(unsafe.Add(mBase, uint32(v354)+8)) = v359
	v361 = *(*int64)(unsafe.Add(mBase, uint32(v358)))
	*(*int64)(unsafe.Add(mBase, uint32(v354))) = v361
	v363 = *(*int32)(unsafe.Add(mBase, uint32(v357)))
	*(*int32)(unsafe.Add(mBase, uint32(v354)+12)) = v363
	goto L94
L98:
	;
	v303 = int32(4)
	v305 = v280 + v291<<(uint(v303)%32)
	v306 = int32(2)
	v308 = v260 + v291<<(uint(v306)%32)
	v309 = *(*int32)(unsafe.Add(mBase, uint32(v308)))
	v310 = *(*int32)(unsafe.Add(mBase, uint32(v309)+8))
	*(*int32)(unsafe.Add(mBase, uint32(v305)+8)) = v310
	v312 = *(*int64)(unsafe.Add(mBase, uint32(v309)))
	*(*int64)(unsafe.Add(mBase, uint32(v305))) = v312
	v314 = *(*int32)(unsafe.Add(mBase, uint32(v308)))
	*(*int32)(unsafe.Add(mBase, uint32(v305)+12)) = v314
	v317 = v291 | int32(1)
	v320 = v280 + v317<<(uint(v303)%32)
	v323 = v260 + v317<<(uint(v306)%32)
	v324 = *(*int32)(unsafe.Add(mBase, uint32(v323)))
	v325 = *(*int32)(unsafe.Add(mBase, uint32(v324)+8))
	*(*int32)(unsafe.Add(mBase, uint32(v320)+8)) = v325
	v327 = *(*int64)(unsafe.Add(mBase, uint32(v324)))
	*(*int64)(unsafe.Add(mBase, uint32(v320))) = v327
	v329 = *(*int32)(unsafe.Add(mBase, uint32(v323)))
	*(*int32)(unsafe.Add(mBase, uint32(v320)+12)) = v329
	v332 = v291 + v306
	v334 = v292 + v306
	if v334 != v256&int32(2147483646) {
		v291 = v332
		v292 = v334
		goto L98
	} else {
		goto L100
	}
L99:
	;
	if v256&int32(1) == int32(0) {
		goto L94
	} else {
		goto L101
	}
L100:
	;
	goto L99
L101:
	;
	v340 = v332
	goto L97
L102:
	;
	F_pg_qsort(m, v280, v256, int32(16), int32(1078))
	mBase = m.M
	v385 = m.ExcPending
	if v385 != 0 {
		goto L18
	} else {
		goto L103
	}
L103:
	;
	v401 = int32(1)
	goto L92
L104:
	;
	v415 = int32(0)
	goto L107
L105:
	;
	goto L106
L106:
	;
	F_pfree(m, v280)
	mBase = m.M
	v663 = m.ExcPending
	if v663 != 0 {
		goto L18
	} else {
		goto L155
	}
L107:
	;
	v422 = *(*int32)(unsafe.Add(mBase, _c_F_smgrDoPendingSyncs[5]))
	v425 = v422 + v415<<(uint(int32(6))%32)
	if v401 == int32(0) {
		goto L111
	} else {
		goto L112
	}
L108:
	;
	goto L106
L109:
	;
	v644 = v415 + int32(1)
	v646 = *(*int32)(unsafe.Add(mBase, _c_F_smgrDoPendingSyncs[4]))
	if v644 < v646 {
		v415 = v644
		goto L107
	} else {
		goto L154
	}
L110:
	;
	F_ReservePrivateRefCountEntry(m)
	mBase = m.M
	v487 = m.ExcPending
	if v487 != 0 {
		goto L18
	} else {
		goto L124
	}
L111:
	;
	if v256 <= int32(0) {
		goto L109
	} else {
		goto L114
	}
L112:
	;
	goto L113
L113:
	;
	v460 = *(*int64)(unsafe.Add(mBase, uint32(v425)))
	v461 = *(*int32)(unsafe.Add(mBase, uint32(v425)+8))
	*(*int32)(unsafe.Add(mBase, uint32(v276)+16)) = v461
	*(*int64)(unsafe.Add(mBase, uint32(v276)+8)) = v460
	v468 = F_bsearch(m, v276+int32(8), v280, v256, int32(16), int32(1078))
	mBase = m.M
	v469 = m.ExcPending
	if v469 != 0 {
		goto L18
	} else {
		goto L122
	}
L114:
	;
	v430 = *(*int32)(unsafe.Add(mBase, uint32(v425)))
	v434 = int32(0)
	goto L115
L115:
	;
	v448 = v280 + v434<<(uint(int32(4))%32)
	v449 = *(*int32)(unsafe.Add(mBase, uint32(v448)))
	if v430 != v449 {
		goto L117
	} else {
		goto L118
	}
L116:
	;
	goto L109
L117:
	;
	v458 = v434 + int32(1)
	if v458 != v256 {
		v434 = v458
		goto L115
	} else {
		goto L121
	}
L118:
	;
	v451 = *(*int32)(unsafe.Add(mBase, uint32(v425)+4))
	v452 = *(*int32)(unsafe.Add(mBase, uint32(v448)+4))
	if v451 != v452 {
		goto L117
	} else {
		goto L119
	}
L119:
	;
	v454 = *(*int32)(unsafe.Add(mBase, uint32(v425)+8))
	v455 = *(*int32)(unsafe.Add(mBase, uint32(v448)+8))
	if v454 == v455 {
		v472 = v448
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
	if v468 == int32(0) {
		goto L109
	} else {
		goto L123
	}
L123:
	;
	v472 = v468
	goto L110
L124:
	;
	v489 = *(*int32)(unsafe.Add(mBase, _c_F_smgrDoPendingSyncs[6]))
	F_ResourceOwnerEnlarge(m, v489)
	mBase = m.M
	v491 = m.ExcPending
	if v491 != 0 {
		goto L18
	} else {
		goto L125
	}
L125:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v276)+28)) = int32(_a_F_smgrDoPendingSyncs_1)
	*(*int32)(unsafe.Add(mBase, uint32(v276)+24)) = int32(_a_F_smgrDoPendingSyncs_2)
	*(*int32)(unsafe.Add(mBase, uint32(v276)+20)) = int32(_a_F_smgrDoPendingSyncs_3)
	*(*int32)(unsafe.Add(mBase, uint32(v276)+16)) = int32(0)
	*(*int64)(unsafe.Add(mBase, uint32(v276)+8)) = int64(0)
	v502 = *(*int32)(unsafe.Add(mBase, uint32(v425)+24))
	v503 = int32(_a_F_smgrDoPendingSyncs_4)
	*(*int32)(unsafe.Add(mBase, uint32(v425)+24)) = v502 | v503
	if v502&v503 != 0 {
		goto L126
	} else {
		goto L127
	}
L126:
	;
	goto L129
L127:
	;
	v534 = v502
	goto L128
L128:
	;
	v549 = int32(_a_F_smgrDoPendingSyncs_5)
	v550 = *(*int32)(unsafe.Add(mBase, _c_F_smgrDoPendingSyncs[7]))
	v552 = *(*int32)(unsafe.Add(mBase, uint32(v276+int32(8))+8))
	if v552 == int32(0) {
		goto L136
	} else {
		goto L137
	}
L129:
	;
	F_perform_spin_delay(m, v276+int32(8))
	mBase = m.M
	v525 = m.ExcPending
	if v525 != 0 {
		goto L18
	} else {
		goto L131
	}
L130:
	;
	v534 = v526
	goto L128
L131:
	;
	v526 = *(*int32)(unsafe.Add(mBase, uint32(v425)+24))
	v527 = int32(_a_F_smgrDoPendingSyncs_4)
	*(*int32)(unsafe.Add(mBase, uint32(v425)+24)) = v526 | v527
	if v526&v527 != 0 {
		goto L129
	} else {
		goto L132
	}
L132:
	;
	goto L130
L133:
	;
	v569 = *(*int32)(unsafe.Add(mBase, uint32(v425)))
	v570 = *(*int32)(unsafe.Add(mBase, uint32(v472)))
	if v569 != v570 {
		goto L144
	} else {
		goto L145
	}
L134:
	;
	goto L133
L135:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_smgrDoPendingSyncs[7])) = v567
	goto L134
L136:
	;
	if int32(999) < v550 {
		goto L134
	} else {
		goto L139
	}
L137:
	;
	goto L138
L138:
	;
	if v550 < int32(11) {
		goto L134
	} else {
		goto L143
	}
L139:
	;
	v557 = int32(900)
	if v557 <= v550 {
		goto L140
	} else {
		goto L141
	}
L140:
	;
	v560 = v557
	goto L142
L141:
	;
	v560 = v550
	goto L142
L142:
	;
	v567 = v560 + int32(100)
	goto L135
L143:
	;
	v567 = v550 - int32(1)
	goto L135
L144:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v425)+24)) = v534 & int32(-4194305)
	goto L109
L145:
	;
	v572 = *(*int32)(unsafe.Add(mBase, uint32(v425)+4))
	v573 = *(*int32)(unsafe.Add(mBase, uint32(v472)+4))
	v575 = int32(25165824)
	if base.B2i32(v572 != v573)|base.B2i32(v534&v575 != v575) != 0 {
		goto L144
	} else {
		goto L146
	}
L146:
	;
	v580 = *(*int32)(unsafe.Add(mBase, uint32(v425)+8))
	v581 = *(*int32)(unsafe.Add(mBase, uint32(v472)+8))
	if v580 != v581 {
		goto L144
	} else {
		goto L147
	}
L147:
	;
	v583 = *(*int32)(unsafe.Add(mBase, uint32(v425)+24))
	v584 = int32(1)
	*(*int32)(unsafe.Add(mBase, uint32(v425)+24)) = (v583 + v584) & int32(-4194305)
	v589 = *(*int32)(unsafe.Add(mBase, uint32(v425)+20))
	v590 = int32(_a_F_smgrDoPendingSyncs_6)
	v591 = *(*int32)(unsafe.Add(mBase, _c_F_smgrDoPendingSyncs[8]))
	*(*int32)(unsafe.Add(mBase, _c_F_smgrDoPendingSyncs[8])) = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v591)+4)) = v584
	v598 = v589 + v584
	*(*int32)(unsafe.Add(mBase, uint32(v591))) = v598
	v601 = *(*int32)(unsafe.Add(mBase, _c_F_smgrDoPendingSyncs[6]))
	F_ResourceOwnerRemember(m, v601, v598, int32(_a_F_smgrDoPendingSyncs_7))
	mBase = m.M
	v604 = m.ExcPending
	if v604 != 0 {
		goto L18
	} else {
		goto L148
	}
L148:
	;
	v606 = v425 + int32(48)
	v608 = F_LWLockAcquire(m, v606, int32(1))
	mBase = m.M
	v609 = m.ExcPending
	if v609 != 0 {
		goto L18
	} else {
		goto L149
	}
L149:
	;
	v610 = *(*int32)(unsafe.Add(mBase, uint32(v472)+12))
	F_FlushBuffer(m, v425, v610, int32(3))
	mBase = m.M
	v613 = m.ExcPending
	if v613 != 0 {
		goto L18
	} else {
		goto L150
	}
L150:
	;
	F_LWLockRelease(m, v606)
	mBase = m.M
	v615 = m.ExcPending
	if v615 != 0 {
		goto L18
	} else {
		goto L151
	}
L151:
	;
	v617 = *(*int32)(unsafe.Add(mBase, _c_F_smgrDoPendingSyncs[6]))
	v618 = *(*int32)(unsafe.Add(mBase, uint32(v425)+20))
	F_ResourceOwnerForget(m, v617, v618+int32(1), int32(_a_F_smgrDoPendingSyncs_7))
	mBase = m.M
	v623 = m.ExcPending
	if v623 != 0 {
		goto L18
	} else {
		goto L152
	}
L152:
	;
	F_UnpinBufferNoOwner(m, v425)
	mBase = m.M
	v625 = m.ExcPending
	if v625 != 0 {
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
	v690 = int32(0)
	goto L159
L157:
	;
	goto L158
L158:
	;
	v755 = int32(_a_F_smgrDoPendingSyncs_0)
	v757 = *(*int32)(unsafe.Add(mBase, _c_F_smgrDoPendingSyncs[3]))
	*(*int32)(unsafe.Add(mBase, _c_F_smgrDoPendingSyncs[3])) = v757 - int32(1)
	goto L87
L159:
	;
	v705 = v260 + v690<<(uint(int32(2))%32)
	v706 = *(*int32)(unsafe.Add(mBase, uint32(v705)))
	v708 = F_mdexists(m, v706, int32(0))
	mBase = m.M
	v709 = m.ExcPending
	if v709 != 0 {
		goto L18
	} else {
		goto L161
	}
L160:
	;
	goto L158
L161:
	;
	if v708 != 0 {
		goto L162
	} else {
		goto L163
	}
L162:
	;
	v710 = *(*int32)(unsafe.Add(mBase, uint32(v705)))
	F_mdimmedsync(m, v710, int32(0))
	mBase = m.M
	v713 = m.ExcPending
	if v713 != 0 {
		goto L18
	} else {
		goto L165
	}
L163:
	;
	goto L164
L164:
	;
	v714 = *(*int32)(unsafe.Add(mBase, uint32(v705)))
	v716 = F_mdexists(m, v714, int32(1))
	mBase = m.M
	v717 = m.ExcPending
	if v717 != 0 {
		goto L18
	} else {
		goto L166
	}
L165:
	;
	goto L164
L166:
	;
	if v716 != 0 {
		goto L167
	} else {
		goto L168
	}
L167:
	;
	v718 = *(*int32)(unsafe.Add(mBase, uint32(v705)))
	F_mdimmedsync(m, v718, int32(1))
	mBase = m.M
	v721 = m.ExcPending
	if v721 != 0 {
		goto L18
	} else {
		goto L170
	}
L168:
	;
	goto L169
L169:
	;
	v722 = *(*int32)(unsafe.Add(mBase, uint32(v705)))
	v724 = F_mdexists(m, v722, int32(2))
	mBase = m.M
	v725 = m.ExcPending
	if v725 != 0 {
		goto L18
	} else {
		goto L171
	}
L170:
	;
	goto L169
L171:
	;
	if v724 != 0 {
		goto L172
	} else {
		goto L173
	}
L172:
	;
	v726 = *(*int32)(unsafe.Add(mBase, uint32(v705)))
	F_mdimmedsync(m, v726, int32(2))
	mBase = m.M
	v729 = m.ExcPending
	if v729 != 0 {
		goto L18
	} else {
		goto L175
	}
L173:
	;
	goto L174
L174:
	;
	v730 = *(*int32)(unsafe.Add(mBase, uint32(v705)))
	v732 = F_mdexists(m, v730, int32(3))
	mBase = m.M
	v733 = m.ExcPending
	if v733 != 0 {
		goto L18
	} else {
		goto L176
	}
L175:
	;
	goto L174
L176:
	;
	if v732 != 0 {
		goto L177
	} else {
		goto L178
	}
L177:
	;
	v734 = *(*int32)(unsafe.Add(mBase, uint32(v705)))
	F_mdimmedsync(m, v734, int32(3))
	mBase = m.M
	v737 = m.ExcPending
	if v737 != 0 {
		goto L18
	} else {
		goto L180
	}
L178:
	;
	goto L179
L179:
	;
	v739 = v690 + int32(1)
	if v739 != v256 {
		v690 = v739
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
