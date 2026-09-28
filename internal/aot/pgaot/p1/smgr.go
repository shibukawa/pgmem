package p1

import (
	base "github.com/shibukawa/pgmem/internal/aot/pgaot/base"
	"unsafe"
)

func F_smgrDoPendingSyncs(m *base.Module, l0 int32, l1 int32) {
	mBase := m.M
	_ = mBase
	var v4 int32
	_ = v4
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
	var v95 int32
	_ = v95
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
	var v118 int64
	_ = v118
	var v119 int32
	_ = v119
	var v121 int32
	_ = v121
	var v122 int32
	_ = v122
	var v124 int32
	_ = v124
	var v125 int32
	_ = v125
	var v128 int64
	_ = v128
	var v129 int32
	_ = v129
	var v132 int32
	_ = v132
	var v133 int32
	_ = v133
	var v138 int32
	_ = v138
	var v139 int32
	_ = v139
	var v142 int64
	_ = v142
	var v143 int32
	_ = v143
	var v145 int32
	_ = v145
	var v146 int32
	_ = v146
	var v148 int32
	_ = v148
	var v149 int32
	_ = v149
	var v152 int64
	_ = v152
	var v153 int32
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
	var v234 int32
	_ = v234
	var v237 int32
	_ = v237
	var v238 int32
	_ = v238
	var v242 int32
	_ = v242
	var v243 int32
	_ = v243
	var v244 int32
	_ = v244
	var v245 int32
	_ = v245
	var v246 int32
	_ = v246
	var v257 int32
	_ = v257
	var v258 int32
	_ = v258
	var v262 int32
	_ = v262
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
	var v274 int32
	_ = v274
	var v276 int32
	_ = v276
	var v279 int32
	_ = v279
	var v280 int32
	_ = v280
	var v291 int32
	_ = v291
	var v301 int32
	_ = v301
	var v302 int32
	_ = v302
	var v304 int32
	_ = v304
	var v305 int32
	_ = v305
	var v307 int32
	_ = v307
	var v308 int32
	_ = v308
	var v309 int32
	_ = v309
	var v311 int64
	_ = v311
	var v313 int32
	_ = v313
	var v316 int32
	_ = v316
	var v319 int32
	_ = v319
	var v322 int32
	_ = v322
	var v323 int32
	_ = v323
	var v324 int32
	_ = v324
	var v326 int64
	_ = v326
	var v328 int32
	_ = v328
	var v331 int32
	_ = v331
	var v333 int32
	_ = v333
	var v340 int32
	_ = v340
	var v353 int32
	_ = v353
	var v356 int32
	_ = v356
	var v357 int32
	_ = v357
	var v358 int32
	_ = v358
	var v360 int64
	_ = v360
	var v362 int32
	_ = v362
	var v384 int32
	_ = v384
	var v400 int32
	_ = v400
	var v402 int32
	_ = v402
	var v416 int32
	_ = v416
	var v421 int32
	_ = v421
	var v424 int32
	_ = v424
	var v429 int32
	_ = v429
	var v434 int32
	_ = v434
	var v447 int32
	_ = v447
	var v448 int32
	_ = v448
	var v450 int32
	_ = v450
	var v451 int32
	_ = v451
	var v453 int32
	_ = v453
	var v454 int32
	_ = v454
	var v457 int32
	_ = v457
	var v459 int64
	_ = v459
	var v460 int32
	_ = v460
	var v467 int32
	_ = v467
	var v468 int32
	_ = v468
	var v477 int32
	_ = v477
	var v486 int32
	_ = v486
	var v488 int32
	_ = v488
	var v490 int32
	_ = v490
	var v491 int64
	_ = v491
	var v493 int64
	_ = v493
	var v500 int64
	_ = v500
	var v520 int64
	_ = v520
	var v543 int32
	_ = v543
	var v544 int64
	_ = v544
	var v547 int64
	_ = v547
	var v569 int32
	_ = v569
	var v570 int32
	_ = v570
	var v572 int32
	_ = v572
	var v577 int32
	_ = v577
	var v580 int32
	_ = v580
	var v587 int32
	_ = v587
	var v589 int64
	_ = v589
	var v591 int64
	_ = v591
	var v598 int64
	_ = v598
	var v610 int32
	_ = v610
	var v611 int32
	_ = v611
	var v613 int32
	_ = v613
	var v614 int32
	_ = v614
	var v616 int64
	_ = v616
	var v621 int32
	_ = v621
	var v622 int32
	_ = v622
	var v624 int64
	_ = v624
	var v626 int32
	_ = v626
	var v627 int64
	_ = v627
	var v633 int64
	_ = v633
	var v637 int64
	_ = v637
	var v654 int64
	_ = v654
	var v670 int32
	_ = v670
	var v671 int32
	_ = v671
	var v676 int32
	_ = v676
	var v677 int32
	_ = v677
	var v678 int32
	_ = v678
	var v681 int32
	_ = v681
	var v699 int32
	_ = v699
	var v703 int32
	_ = v703
	var v704 int32
	_ = v704
	var v705 int32
	_ = v705
	var v707 int32
	_ = v707
	var v710 int32
	_ = v710
	var v713 int32
	_ = v713
	var v715 int32
	_ = v715
	var v717 int32
	_ = v717
	var v718 int32
	_ = v718
	var v724 int32
	_ = v724
	var v726 int32
	_ = v726
	var v729 int64
	_ = v729
	var v745 int32
	_ = v745
	var v747 int32
	_ = v747
	var v764 int32
	_ = v764
	var v782 int32
	_ = v782
	var v784 int32
	_ = v784
	var v791 int32
	_ = v791
	var v806 int32
	_ = v806
	var v807 int32
	_ = v807
	var v809 int32
	_ = v809
	var v810 int32
	_ = v810
	var v811 int32
	_ = v811
	var v814 int32
	_ = v814
	var v815 int32
	_ = v815
	var v817 int32
	_ = v817
	var v818 int32
	_ = v818
	var v819 int32
	_ = v819
	var v822 int32
	_ = v822
	var v823 int32
	_ = v823
	var v825 int32
	_ = v825
	var v826 int32
	_ = v826
	var v827 int32
	_ = v827
	var v830 int32
	_ = v830
	var v831 int32
	_ = v831
	var v833 int32
	_ = v833
	var v834 int32
	_ = v834
	var v835 int32
	_ = v835
	var v838 int32
	_ = v838
	var v840 int32
	_ = v840
	var v856 int32
	_ = v856
	var v858 int32
	_ = v858
	var v867 int32
	_ = v867
	var v877 int32
	_ = v877
	var v889 int32
	_ = v889
	v4 = int32(0)
	v15 = m.G0
	v17 = v15 - int32(96)
	m.G0 = v17
	v20 = *(*int32)(unsafe.Add(mBase, _c_F_smgrDoPendingSyncs[0]))
	if v20 == v4 {
		v889 = v17
		goto L1
	} else {
		goto L2
	}
L1:
	;
	m.G0 = v889 + int32(96)
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
	v889 = v17
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
	v889 = v17
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
	v889 = v17
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
	v90 = v4
	v95 = v4
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
		v889 = v17
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
	v118 = int64(0)
	v119 = v109
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
	v118 = base.I64_extend_i32_u(v115)
	v119 = v115
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
	v128 = v118
	v129 = v109
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
	v128 = v118 + base.I64_extend_i32_u(v124)
	v129 = v124
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
	v142 = v128
	v143 = int32(-1)
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
	v142 = v128 + base.I64_extend_i32_u(v138)
	v143 = v138
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
	v152 = v142
	v153 = int32(-1)
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
	v152 = v142 + base.I64_extend_i32_u(v148)
	v153 = v148
	goto L49
L51:
	;
	v156 = int64(*(*int32)(unsafe.Add(mBase, _c_F_smgrDoPendingSyncs[2])))
	if base.Ui64(int64(base.Ui64(v156)>>(uint(int64(3))%64))&int64(2251799813685247)) <= base.Ui64(v152) {
		goto L28
	} else {
		goto L52
	}
L52:
	;
	if v119 != int32(-1) {
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
	if v129 != int32(-1) {
		goto L59
	} else {
		goto L60
	}
L56:
	;
	v172 = int32(0)
	F_log_newpage_range(m, v170, v172, v119, v172)
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
	if v143 != int32(-1) {
		goto L65
	} else {
		goto L66
	}
L62:
	;
	F_log_newpage_range(m, v187, int32(1), v129, int32(0))
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
	if v153 == int32(-1) {
		v257 = v88
		v258 = v90
		v262 = v95
		goto L27
	} else {
		goto L71
	}
L68:
	;
	F_log_newpage_range(m, v204, int32(2), v143, int32(0))
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
	F_log_newpage_range(m, v219, int32(3), v153, int32(0))
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
	v262 = v95
	goto L27
L75:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v246+v88<<(uint(int32(2))%32)))) = v105
	v257 = v88 + int32(1)
	v258 = v245
	v262 = v246
	goto L27
L76:
	;
	v234 = int32(8)
	v237 = F_palloc_mul(m, int32(4), v234)
	mBase = m.M
	v238 = m.ExcPending
	if v238 != 0 {
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
		v246 = v95
		goto L75
	} else {
		goto L80
	}
L79:
	;
	v245 = v234
	v246 = v237
	goto L75
L80:
	;
	v242 = v90 << (uint(int32(1)) % 32)
	v243 = F_repalloc_mul(m, v95, int32(4), v242)
	mBase = m.M
	v244 = m.ExcPending
	if v244 != 0 {
		goto L18
	} else {
		goto L81
	}
L81:
	;
	v245 = v242
	v246 = v243
	goto L75
L82:
	;
	if v265 != 0 {
		v85 = v265
		v88 = v257
		v90 = v258
		v95 = v262
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
	v274 = m.G0
	v276 = v274 - int32(32)
	m.G0 = v276
	if v257 != 0 {
		goto L88
	} else {
		goto L89
	}
L86:
	;
	v867 = v17
	goto L87
L87:
	;
	F_pfree(m, v262)
	mBase = m.M
	v877 = m.ExcPending
	if v877 != 0 {
		goto L18
	} else {
		goto L194
	}
L88:
	;
	v279 = F_palloc_mul(m, int32(16), v257)
	mBase = m.M
	v280 = m.ExcPending
	if v280 != 0 {
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
	v782 = int32(_a_F_smgrDoPendingSyncs_0)
	v784 = *(*int32)(unsafe.Add(mBase, _c_F_smgrDoPendingSyncs[3]))
	*(*int32)(unsafe.Add(mBase, _c_F_smgrDoPendingSyncs[3])) = v784 + int32(1)
	if int32(0) < v257 {
		goto L168
	} else {
		goto L169
	}
L91:
	;
	if v257 <= v268 {
		v400 = int32(0)
		goto L92
	} else {
		goto L93
	}
L92:
	;
	v402 = *(*int32)(unsafe.Add(mBase, _c_F_smgrDoPendingSyncs[4]))
	if int32(0) < v402 {
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
		v400 = int32(0)
		goto L92
	} else {
		goto L102
	}
L95:
	;
	v291 = v273
	v301 = v4
	goto L98
L96:
	;
	v340 = v273
	goto L97
L97:
	;
	v353 = v279 + v340<<(uint(int32(4))%32)
	v356 = v262 + v340<<(uint(int32(2))%32)
	v357 = *(*int32)(unsafe.Add(mBase, uint32(v356)))
	v358 = *(*int32)(unsafe.Add(mBase, uint32(v357)+8))
	*(*int32)(unsafe.Add(mBase, uint32(v353)+8)) = v358
	v360 = *(*int64)(unsafe.Add(mBase, uint32(v357)))
	*(*int64)(unsafe.Add(mBase, uint32(v353))) = v360
	v362 = *(*int32)(unsafe.Add(mBase, uint32(v356)))
	*(*int32)(unsafe.Add(mBase, uint32(v353)+12)) = v362
	goto L94
L98:
	;
	v302 = int32(4)
	v304 = v279 + v291<<(uint(v302)%32)
	v305 = int32(2)
	v307 = v262 + v291<<(uint(v305)%32)
	v308 = *(*int32)(unsafe.Add(mBase, uint32(v307)))
	v309 = *(*int32)(unsafe.Add(mBase, uint32(v308)+8))
	*(*int32)(unsafe.Add(mBase, uint32(v304)+8)) = v309
	v311 = *(*int64)(unsafe.Add(mBase, uint32(v308)))
	*(*int64)(unsafe.Add(mBase, uint32(v304))) = v311
	v313 = *(*int32)(unsafe.Add(mBase, uint32(v307)))
	*(*int32)(unsafe.Add(mBase, uint32(v304)+12)) = v313
	v316 = v291 | int32(1)
	v319 = v279 + v316<<(uint(v302)%32)
	v322 = v262 + v316<<(uint(v305)%32)
	v323 = *(*int32)(unsafe.Add(mBase, uint32(v322)))
	v324 = *(*int32)(unsafe.Add(mBase, uint32(v323)+8))
	*(*int32)(unsafe.Add(mBase, uint32(v319)+8)) = v324
	v326 = *(*int64)(unsafe.Add(mBase, uint32(v323)))
	*(*int64)(unsafe.Add(mBase, uint32(v319))) = v326
	v328 = *(*int32)(unsafe.Add(mBase, uint32(v322)))
	*(*int32)(unsafe.Add(mBase, uint32(v319)+12)) = v328
	v331 = v291 + v305
	v333 = v301 + v305
	if v333 != v257&int32(2147483646) {
		v291 = v331
		v301 = v333
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
	v340 = v331
	goto L97
L102:
	;
	F_pg_qsort(m, v279, v257, int32(16), int32(1165))
	mBase = m.M
	v384 = m.ExcPending
	if v384 != 0 {
		goto L18
	} else {
		goto L103
	}
L103:
	;
	v400 = int32(1)
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
	F_pfree(m, v279)
	mBase = m.M
	v764 = m.ExcPending
	if v764 != 0 {
		goto L18
	} else {
		goto L167
	}
L107:
	;
	v421 = *(*int32)(unsafe.Add(mBase, _c_F_smgrDoPendingSyncs[5]))
	v424 = v421 + v416*int32(56)
	if v400 == int32(0) {
		goto L111
	} else {
		goto L112
	}
L108:
	;
	goto L106
L109:
	;
	v745 = v416 + int32(1)
	v747 = *(*int32)(unsafe.Add(mBase, _c_F_smgrDoPendingSyncs[4]))
	if v745 < v747 {
		v416 = v745
		goto L107
	} else {
		goto L166
	}
L110:
	;
	F_ReservePrivateRefCountEntry(m)
	mBase = m.M
	v486 = m.ExcPending
	if v486 != 0 {
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
	v459 = *(*int64)(unsafe.Add(mBase, uint32(v424)))
	v460 = *(*int32)(unsafe.Add(mBase, uint32(v424)+8))
	*(*int32)(unsafe.Add(mBase, uint32(v276)+16)) = v460
	*(*int64)(unsafe.Add(mBase, uint32(v276)+8)) = v459
	v467 = F_bsearch(m, v276+int32(8), v279, v257, int32(16), int32(1165))
	mBase = m.M
	v468 = m.ExcPending
	if v468 != 0 {
		goto L18
	} else {
		goto L122
	}
L114:
	;
	v429 = *(*int32)(unsafe.Add(mBase, uint32(v424)))
	v434 = int32(0)
	goto L115
L115:
	;
	v447 = v279 + v434<<(uint(int32(4))%32)
	v448 = *(*int32)(unsafe.Add(mBase, uint32(v447)))
	if v429 != v448 {
		goto L117
	} else {
		goto L118
	}
L116:
	;
	goto L109
L117:
	;
	v457 = v434 + int32(1)
	if v457 != v257 {
		v434 = v457
		goto L115
	} else {
		goto L121
	}
L118:
	;
	v450 = *(*int32)(unsafe.Add(mBase, uint32(v424)+4))
	v451 = *(*int32)(unsafe.Add(mBase, uint32(v447)+4))
	if v450 != v451 {
		goto L117
	} else {
		goto L119
	}
L119:
	;
	v453 = *(*int32)(unsafe.Add(mBase, uint32(v424)+8))
	v454 = *(*int32)(unsafe.Add(mBase, uint32(v447)+8))
	if v453 == v454 {
		v477 = v447
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
	if v467 == int32(0) {
		goto L109
	} else {
		goto L123
	}
L123:
	;
	v477 = v467
	goto L110
L124:
	;
	v488 = *(*int32)(unsafe.Add(mBase, _c_F_smgrDoPendingSyncs[6]))
	F_ResourceOwnerEnlarge(m, v488)
	mBase = m.M
	v490 = m.ExcPending
	if v490 != 0 {
		goto L18
	} else {
		goto L125
	}
L125:
	;
	v491 = int64(4194304)
	v493 = base.AtomicRmwOr64(m, v424, int32(24), v491)
	if v493&v491 != int64(0) {
		goto L126
	} else {
		goto L127
	}
L126:
	;
	v500 = v493
	goto L129
L127:
	;
	v598 = v493
	goto L128
L128:
	;
	v610 = *(*int32)(unsafe.Add(mBase, uint32(v424)))
	v611 = *(*int32)(unsafe.Add(mBase, uint32(v477)))
	if v610 != v611 {
		goto L150
	} else {
		goto L151
	}
L129:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v276)+28)) = int32(_a_F_smgrDoPendingSyncs_1)
	*(*int32)(unsafe.Add(mBase, uint32(v276)+24)) = int32(_a_F_smgrDoPendingSyncs_2)
	*(*int32)(unsafe.Add(mBase, uint32(v276)+20)) = int32(_a_F_smgrDoPendingSyncs_3)
	*(*int32)(unsafe.Add(mBase, uint32(v276)+16)) = int32(0)
	v520 = int64(0)
	*(*int64)(unsafe.Add(mBase, uint32(v276)+8)) = v520
	if v500&int64(4194304) != v520 {
		goto L131
	} else {
		goto L132
	}
L130:
	;
	v598 = v591
	goto L128
L131:
	;
	goto L134
L132:
	;
	goto L133
L133:
	;
	v569 = int32(_a_F_smgrDoPendingSyncs_4)
	v570 = *(*int32)(unsafe.Add(mBase, _c_F_smgrDoPendingSyncs[7]))
	v572 = *(*int32)(unsafe.Add(mBase, uint32(v276+int32(8))+8))
	if v572 == int32(0) {
		goto L141
	} else {
		goto L142
	}
L134:
	;
	F_perform_spin_delay(m, v276+int32(8))
	mBase = m.M
	v543 = m.ExcPending
	if v543 != 0 {
		goto L18
	} else {
		goto L136
	}
L135:
	;
	goto L133
L136:
	;
	v544 = int64(0)
	v547 = base.AtomicRmwCmpxchg64(m, v424, int32(24), v544, v544)
	if v547&int64(4194304) != v544 {
		goto L134
	} else {
		goto L137
	}
L137:
	;
	goto L135
L138:
	;
	v589 = int64(4194304)
	v591 = base.AtomicRmwOr64(m, v424, int32(24), v589)
	if v591&v589 != int64(0) {
		v500 = v591
		goto L129
	} else {
		goto L149
	}
L139:
	;
	goto L138
L140:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_smgrDoPendingSyncs[7])) = v587
	goto L139
L141:
	;
	if int32(999) < v570 {
		goto L139
	} else {
		goto L144
	}
L142:
	;
	goto L143
L143:
	;
	if v570 < int32(11) {
		goto L139
	} else {
		goto L148
	}
L144:
	;
	v577 = int32(900)
	if v577 <= v570 {
		goto L145
	} else {
		goto L146
	}
L145:
	;
	v580 = v577
	goto L147
L146:
	;
	v580 = v570
	goto L147
L147:
	;
	v587 = v580 + int32(100)
	goto L140
L148:
	;
	v587 = v570 - int32(1)
	goto L140
L149:
	;
	goto L130
L150:
	;
	v729 = base.AtomicRmwSub64(m, v424, int32(24), int64(4194304))
	goto L109
L151:
	;
	v613 = *(*int32)(unsafe.Add(mBase, uint32(v424)+4))
	v614 = *(*int32)(unsafe.Add(mBase, uint32(v477)+4))
	v616 = int64(25165824)
	if base.B2i32(v613 != v614)|base.B2i32(v598&v616 != v616) != 0 {
		goto L150
	} else {
		goto L152
	}
L152:
	;
	v621 = *(*int32)(unsafe.Add(mBase, uint32(v424)+8))
	v622 = *(*int32)(unsafe.Add(mBase, uint32(v477)+8))
	if v621 != v622 {
		goto L150
	} else {
		goto L153
	}
L153:
	;
	v624 = int64(0)
	v626 = int32(24)
	v627 = base.AtomicRmwCmpxchg64(m, v424, v626, v624, v624)
	v633 = base.AtomicRmwCmpxchg64(m, v424, v626, v627, v627&int64(-4194305)+int64(1))
	if v627 != v633 {
		goto L154
	} else {
		goto L155
	}
L154:
	;
	v637 = v633
	goto L157
L155:
	;
	goto L156
L156:
	;
	v670 = int32(_a_F_smgrDoPendingSyncs_5)
	v671 = *(*int32)(unsafe.Add(mBase, _c_F_smgrDoPendingSyncs[8]))
	v676 = *(*int32)(unsafe.Add(mBase, uint32(v424)+20))
	v677 = int32(1)
	v678 = v676 + v677
	*(*int32)(unsafe.Add(mBase, uint32(v671<<(uint(int32(2))%32))+uint32(_c_F_smgrDoPendingSyncs[9]))) = v678
	v681 = v671 << (uint(int32(4)) % 32)
	*(*int32)(unsafe.Add(mBase, uint32(v681)+uint32(_c_F_smgrDoPendingSyncs[10]))) = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v681)+uint32(_c_F_smgrDoPendingSyncs[11]))) = v678
	*(*int32)(unsafe.Add(mBase, _c_F_smgrDoPendingSyncs[8])) = int32(-1)
	*(*int32)(unsafe.Add(mBase, _c_F_smgrDoPendingSyncs[12])) = v671
	*(*int32)(unsafe.Add(mBase, uint32(v681)+uint32(_c_F_smgrDoPendingSyncs[13]))) = v677
	v699 = *(*int32)(unsafe.Add(mBase, _c_F_smgrDoPendingSyncs[6]))
	F_ResourceOwnerRemember(m, v699, base.I64_extend_i32_s(v678), int32(_a_F_smgrDoPendingSyncs_6))
	mBase = m.M
	v703 = m.ExcPending
	if v703 != 0 {
		goto L18
	} else {
		goto L160
	}
L157:
	;
	v654 = base.AtomicRmwCmpxchg64(m, v424, int32(24), v637, v637&int64(-4194305)+int64(1))
	if v637 != v654 {
		v637 = v654
		goto L157
	} else {
		goto L159
	}
L158:
	;
	goto L156
L159:
	;
	goto L158
L160:
	;
	v704 = *(*int32)(unsafe.Add(mBase, uint32(v477)+12))
	v705 = *(*int32)(unsafe.Add(mBase, uint32(v424)+20))
	v707 = v705 + int32(1)
	F_BufferLockAcquire(m, v707, v424, int32(2))
	mBase = m.M
	v710 = m.ExcPending
	if v710 != 0 {
		goto L18
	} else {
		goto L161
	}
L161:
	;
	F_FlushBuffer(m, v424, v704, int32(3))
	mBase = m.M
	v713 = m.ExcPending
	if v713 != 0 {
		goto L18
	} else {
		goto L162
	}
L162:
	;
	F_BufferLockUnlock(m, v707, v424)
	mBase = m.M
	v715 = m.ExcPending
	if v715 != 0 {
		goto L18
	} else {
		goto L163
	}
L163:
	;
	v717 = *(*int32)(unsafe.Add(mBase, _c_F_smgrDoPendingSyncs[6]))
	v718 = *(*int32)(unsafe.Add(mBase, uint32(v424)+20))
	F_ResourceOwnerForget(m, v717, base.I64_extend_i32_s(v718+int32(1)), int32(_a_F_smgrDoPendingSyncs_6))
	mBase = m.M
	v724 = m.ExcPending
	if v724 != 0 {
		goto L18
	} else {
		goto L164
	}
L164:
	;
	F_UnpinBufferNoOwner(m, v424)
	mBase = m.M
	v726 = m.ExcPending
	if v726 != 0 {
		goto L18
	} else {
		goto L165
	}
L165:
	;
	goto L109
L166:
	;
	goto L108
L167:
	;
	goto L90
L168:
	;
	v791 = int32(0)
	goto L171
L169:
	;
	goto L170
L170:
	;
	v856 = int32(_a_F_smgrDoPendingSyncs_0)
	v858 = *(*int32)(unsafe.Add(mBase, _c_F_smgrDoPendingSyncs[3]))
	*(*int32)(unsafe.Add(mBase, _c_F_smgrDoPendingSyncs[3])) = v858 - int32(1)
	v867 = v17
	goto L87
L171:
	;
	v806 = v262 + v791<<(uint(int32(2))%32)
	v807 = *(*int32)(unsafe.Add(mBase, uint32(v806)))
	v809 = F_mdexists(m, v807, int32(0))
	mBase = m.M
	v810 = m.ExcPending
	if v810 != 0 {
		goto L18
	} else {
		goto L173
	}
L172:
	;
	goto L170
L173:
	;
	if v809 != 0 {
		goto L174
	} else {
		goto L175
	}
L174:
	;
	v811 = *(*int32)(unsafe.Add(mBase, uint32(v806)))
	F_mdimmedsync(m, v811, int32(0))
	mBase = m.M
	v814 = m.ExcPending
	if v814 != 0 {
		goto L18
	} else {
		goto L177
	}
L175:
	;
	goto L176
L176:
	;
	v815 = *(*int32)(unsafe.Add(mBase, uint32(v806)))
	v817 = F_mdexists(m, v815, int32(1))
	mBase = m.M
	v818 = m.ExcPending
	if v818 != 0 {
		goto L18
	} else {
		goto L178
	}
L177:
	;
	goto L176
L178:
	;
	if v817 != 0 {
		goto L179
	} else {
		goto L180
	}
L179:
	;
	v819 = *(*int32)(unsafe.Add(mBase, uint32(v806)))
	F_mdimmedsync(m, v819, int32(1))
	mBase = m.M
	v822 = m.ExcPending
	if v822 != 0 {
		goto L18
	} else {
		goto L182
	}
L180:
	;
	goto L181
L181:
	;
	v823 = *(*int32)(unsafe.Add(mBase, uint32(v806)))
	v825 = F_mdexists(m, v823, int32(2))
	mBase = m.M
	v826 = m.ExcPending
	if v826 != 0 {
		goto L18
	} else {
		goto L183
	}
L182:
	;
	goto L181
L183:
	;
	if v825 != 0 {
		goto L184
	} else {
		goto L185
	}
L184:
	;
	v827 = *(*int32)(unsafe.Add(mBase, uint32(v806)))
	F_mdimmedsync(m, v827, int32(2))
	mBase = m.M
	v830 = m.ExcPending
	if v830 != 0 {
		goto L18
	} else {
		goto L187
	}
L185:
	;
	goto L186
L186:
	;
	v831 = *(*int32)(unsafe.Add(mBase, uint32(v806)))
	v833 = F_mdexists(m, v831, int32(3))
	mBase = m.M
	v834 = m.ExcPending
	if v834 != 0 {
		goto L18
	} else {
		goto L188
	}
L187:
	;
	goto L186
L188:
	;
	if v833 != 0 {
		goto L189
	} else {
		goto L190
	}
L189:
	;
	v835 = *(*int32)(unsafe.Add(mBase, uint32(v806)))
	F_mdimmedsync(m, v835, int32(3))
	mBase = m.M
	v838 = m.ExcPending
	if v838 != 0 {
		goto L18
	} else {
		goto L192
	}
L190:
	;
	goto L191
L191:
	;
	v840 = v791 + int32(1)
	if v840 != v257 {
		v791 = v840
		goto L171
	} else {
		goto L193
	}
L192:
	;
	goto L191
L193:
	;
	goto L172
L194:
	;
	v889 = v867
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
