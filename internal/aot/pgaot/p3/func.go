package p3

import (
	base "github.com/shibukawa/pgmem/internal/aot/pgaot/base"
	"unsafe"
)

func F_check_func_name(m *base.Module, l0 int32, l1 int32) int32 {
	mBase := m.M
	_ = mBase
	var v3 int32
	_ = v3
	var v8 int32
	_ = v8
	var v11 int32
	_ = v11
	var v14 int32
	_ = v14
	var v15 int32
	_ = v15
	var v19 int32
	_ = v19
	var v24 int32
	_ = v24
	var v25 int32
	_ = v25
	var v29 int32
	_ = v29
	var v35 int32
	_ = v35
	v3 = int32(0)
	if l0 == v3 {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	return l0
L2:
	;
	v8 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	if v8 <= int32(0) {
		goto L1
	} else {
		goto L3
	}
L3:
	;
	v11 = int32(0)
	if v11 < v8 {
		goto L4
	} else {
		goto L5
	}
L4:
	;
	v14 = v8
	goto L6
L5:
	;
	v14 = v11
	goto L6
L6:
	;
	v15 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	v19 = v3
	goto L7
L7:
	;
	v24 = *(*int32)(unsafe.Add(mBase, uint32(v15+v19<<(uint(int32(2))%32))))
	v25 = *(*int32)(unsafe.Add(mBase, uint32(v24)))
	if v25 == int32(476) {
		goto L9
	} else {
		goto L10
	}
L8:
	;
	F_scanner_yyerror(m, int32(_a_F_check_func_name_0), l1)
	mBase = m.M
	v35 = m.ExcPending
	if v35 != 0 {
		goto L13
	} else {
		goto L14
	}
L9:
	;
	v29 = v19 + int32(1)
	if v14 != v29 {
		v19 = v29
		goto L7
	} else {
		goto L12
	}
L10:
	;
	goto L11
L11:
	;
	goto L8
L12:
	;
	goto L1
L13:
	;
	return int32(0)
L14:
	;
	base.Wasm_trap_unreachable()
	for {
	}
}
func F_func_select_candidate(m *base.Module, l0 int32, l1 int32, l2 int32) int32 {
	mBase := m.M
	_ = mBase
	var v4 int32
	_ = v4
	var v17 int32
	_ = v17
	var v19 int32
	_ = v19
	var v28 int32
	_ = v28
	var v39 int32
	_ = v39
	var v42 int32
	_ = v42
	var v44 int32
	_ = v44
	var v47 int32
	_ = v47
	var v50 int32
	_ = v50
	var v54 int32
	_ = v54
	var v55 int32
	_ = v55
	var v61 int32
	_ = v61
	var v77 int32
	_ = v77
	var v85 int32
	_ = v85
	var v89 int32
	_ = v89
	var v90 int32
	_ = v90
	var v93 int32
	_ = v93
	var v95 int32
	_ = v95
	var v99 int32
	_ = v99
	var v103 int32
	_ = v103
	var v104 int32
	_ = v104
	var v112 int32
	_ = v112
	var v117 int32
	_ = v117
	var v120 int32
	_ = v120
	var v126 int32
	_ = v126
	var v130 int32
	_ = v130
	var v134 int32
	_ = v134
	var v137 int32
	_ = v137
	var v141 int32
	_ = v141
	var v145 int32
	_ = v145
	var v149 int32
	_ = v149
	var v152 int32
	_ = v152
	var v153 int32
	_ = v153
	var v154 int32
	_ = v154
	var v156 int32
	_ = v156
	var v161 int32
	_ = v161
	var v168 int32
	_ = v168
	var v177 int32
	_ = v177
	var v181 int32
	_ = v181
	var v185 int32
	_ = v185
	var v196 int32
	_ = v196
	var v204 int32
	_ = v204
	var v215 int32
	_ = v215
	var v216 int32
	_ = v216
	var v217 int32
	_ = v217
	var v218 int32
	_ = v218
	var v219 int32
	_ = v219
	var v231 int32
	_ = v231
	var v240 int32
	_ = v240
	var v246 int32
	_ = v246
	var v267 int32
	_ = v267
	var v268 int32
	_ = v268
	var v269 int32
	_ = v269
	var v272 int32
	_ = v272
	var v290 int32
	_ = v290
	var v294 int32
	_ = v294
	var v300 int32
	_ = v300
	var v303 int32
	_ = v303
	var v304 int32
	_ = v304
	var v305 int32
	_ = v305
	var v308 int32
	_ = v308
	var v314 int32
	_ = v314
	var v321 int32
	_ = v321
	var v323 int32
	_ = v323
	var v337 int32
	_ = v337
	var v341 int32
	_ = v341
	var v345 int32
	_ = v345
	var v350 int32
	_ = v350
	var v351 int32
	_ = v351
	var v352 int32
	_ = v352
	var v357 int32
	_ = v357
	var v360 int32
	_ = v360
	var v363 int32
	_ = v363
	var v378 int32
	_ = v378
	var v389 int32
	_ = v389
	var v390 int32
	_ = v390
	var v391 int32
	_ = v391
	var v392 int32
	_ = v392
	var v393 int32
	_ = v393
	var v408 int32
	_ = v408
	var v414 int32
	_ = v414
	var v420 int32
	_ = v420
	var v424 int32
	_ = v424
	var v430 int32
	_ = v430
	var v439 int32
	_ = v439
	var v443 int32
	_ = v443
	var v446 int32
	_ = v446
	var v449 int32
	_ = v449
	var v454 int32
	_ = v454
	var v465 int32
	_ = v465
	var v467 int32
	_ = v467
	var v468 int32
	_ = v468
	var v471 int32
	_ = v471
	var v479 int32
	_ = v479
	var v485 int32
	_ = v485
	var v486 int32
	_ = v486
	var v488 int32
	_ = v488
	var v495 int32
	_ = v495
	var v508 int32
	_ = v508
	var v512 int32
	_ = v512
	var v518 int32
	_ = v518
	var v519 int32
	_ = v519
	var v528 int32
	_ = v528
	var v531 int32
	_ = v531
	var v541 int32
	_ = v541
	var v544 int32
	_ = v544
	var v546 int32
	_ = v546
	var v550 int32
	_ = v550
	var v553 int32
	_ = v553
	var v563 int32
	_ = v563
	var v568 int32
	_ = v568
	var v570 int32
	_ = v570
	var v571 int32
	_ = v571
	var v581 int32
	_ = v581
	var v583 int32
	_ = v583
	var v584 int32
	_ = v584
	var v594 int32
	_ = v594
	var v618 int32
	_ = v618
	var v627 int32
	_ = v627
	var v629 int32
	_ = v629
	var v631 int32
	_ = v631
	var v638 int32
	_ = v638
	var v641 int32
	_ = v641
	var __phi641 int32
	_ = __phi641
	var v645 int32
	_ = v645
	var __phi645 int32
	_ = __phi645
	var v646 int32
	_ = v646
	var __phi646 int32
	_ = __phi646
	var v649 int32
	_ = v649
	var __phi649 int32
	_ = __phi649
	var v662 int32
	_ = v662
	var v676 int32
	_ = v676
	var v680 int32
	_ = v680
	var v684 int32
	_ = v684
	var v690 int32
	_ = v690
	var v691 int32
	_ = v691
	var v695 int32
	_ = v695
	var v700 int32
	_ = v700
	var v703 int32
	_ = v703
	var v709 int32
	_ = v709
	var v712 int32
	_ = v712
	var v713 int32
	_ = v713
	var v714 int32
	_ = v714
	var v721 int32
	_ = v721
	var v722 int32
	_ = v722
	var v733 int32
	_ = v733
	var v743 int32
	_ = v743
	var v752 int32
	_ = v752
	var v756 int32
	_ = v756
	var v770 int32
	_ = v770
	var v776 int32
	_ = v776
	var v778 int32
	_ = v778
	var v780 int32
	_ = v780
	var v786 int32
	_ = v786
	var v787 int32
	_ = v787
	var v803 int32
	_ = v803
	var v805 int32
	_ = v805
	var v814 int32
	_ = v814
	var v823 int32
	_ = v823
	var v824 int32
	_ = v824
	var v826 int32
	_ = v826
	var v841 int32
	_ = v841
	var v847 int32
	_ = v847
	var v857 int32
	_ = v857
	var v868 int32
	_ = v868
	var v871 int32
	_ = v871
	var v897 int32
	_ = v897
	var v912 int32
	_ = v912
	var v913 int32
	_ = v913
	var v914 int32
	_ = v914
	var v916 int32
	_ = v916
	var v927 int32
	_ = v927
	var v940 int32
	_ = v940
	var v941 int32
	_ = v941
	var v942 int32
	_ = v942
	var v959 int32
	_ = v959
	var v984 int32
	_ = v984
	var v1002 int32
	_ = v1002
	var v1018 int32
	_ = v1018
	var v1021 int32
	_ = v1021
	var v1022 int32
	_ = v1022
	var v1028 int32
	_ = v1028
	var v1033 int32
	_ = v1033
	v4 = int32(0)
	v17 = m.G0
	v19 = v17 - int32(640)
	m.G0 = v19
	if l0 <= int32(100) {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	if int32(0) < l0 {
		goto L4
	} else {
		goto L5
	}
L2:
	;
	goto L3
L3:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v1018 = m.ExcPending
	if v1018 != 0 {
		goto L13
	} else {
		goto L179
	}
L4:
	;
	v28 = v4
	v39 = v4
	goto L7
L5:
	;
	v77 = v4
	goto L6
L6:
	;
	if l2 != 0 {
		goto L17
	} else {
		goto L18
	}
L7:
	;
	v42 = v28 << (uint(int32(2)) % 32)
	v44 = *(*int32)(unsafe.Add(mBase, uint32(l1+v42)))
	if v44 != int32(705) {
		goto L10
	} else {
		goto L11
	}
L8:
	;
	v77 = v54
	goto L6
L9:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v19+int32(240)+v42))) = v55
	v61 = v28 + int32(1)
	if v61 != l0 {
		v28 = v61
		v39 = v54
		goto L7
	} else {
		goto L15
	}
L10:
	;
	v47 = F_getBaseType(m, v44)
	mBase = m.M
	v50 = m.ExcPending
	if v50 != 0 {
		goto L13
	} else {
		goto L14
	}
L11:
	;
	goto L12
L12:
	;
	v54 = v39 + int32(1)
	v55 = int32(705)
	goto L9
L13:
	;
	return int32(0)
L14:
	;
	v54 = v39
	v55 = v47
	goto L9
L15:
	;
	goto L8
L16:
	;
	m.G0 = v19 + int32(640)
	return v1002
L17:
	;
	v85 = l2
	v89 = v4
	v90 = l2
	v93 = v4
	v95 = v4
	goto L20
L18:
	;
	v231 = v4
	goto L19
L19:
	;
	v240 = int32(0)
	if v240 < l0 {
		goto L45
	} else {
		goto L46
	}
L20:
	;
	v99 = int32(0)
	if l0 <= v99 {
		v196 = v99
		goto L22
	} else {
		goto L23
	}
L21:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v215))) = int32(0)
	if v218 == int32(1) {
		v1002 = v216
		goto L16
	} else {
		goto L44
	}
L22:
	;
	v204 = int32(0)
	if base.B2i32(v89 == v204)|base.B2i32(v93 < v196) == v204 {
		goto L39
	} else {
		goto L40
	}
L23:
	;
	v103 = v85 + int32(32)
	v104 = int32(0)
	if l0 != int32(1) {
		goto L24
	} else {
		goto L25
	}
L24:
	;
	v112 = v104
	v117 = v99
	v120 = v104
	goto L27
L25:
	;
	v161 = v104
	v168 = v99
	goto L26
L26:
	;
	v177 = v161 << (uint(int32(2)) % 32)
	v181 = *(*int32)(unsafe.Add(mBase, uint32(v177+(v19+int32(240)))))
	if v181 == int32(705) {
		v196 = v168
		goto L22
	} else {
		goto L37
	}
L27:
	;
	v126 = v112 << (uint(int32(2)) % 32)
	v130 = *(*int32)(unsafe.Add(mBase, uint32(v126+(v19+int32(240)))))
	if v130 != int32(705) {
		goto L29
	} else {
		goto L30
	}
L28:
	;
	if l0&int32(1) == int32(0) {
		v196 = v152
		goto L22
	} else {
		goto L36
	}
L29:
	;
	v134 = *(*int32)(unsafe.Add(mBase, uint32(v126+v103)))
	v137 = v117 + base.B2i32(v134 == v130)
	goto L31
L30:
	;
	v137 = v117
	goto L31
L31:
	;
	v141 = (v112 | int32(1)) << (uint(int32(2)) % 32)
	v145 = *(*int32)(unsafe.Add(mBase, uint32(v141+(v19+int32(240)))))
	if v145 != int32(705) {
		goto L32
	} else {
		goto L33
	}
L32:
	;
	v149 = *(*int32)(unsafe.Add(mBase, uint32(v141+v103)))
	v152 = v137 + base.B2i32(v149 == v145)
	goto L34
L33:
	;
	v152 = v137
	goto L34
L34:
	;
	v153 = int32(2)
	v154 = v112 + v153
	v156 = v120 + v153
	if v156 != l0&int32(2147483646) {
		v112 = v154
		v117 = v152
		v120 = v156
		goto L27
	} else {
		goto L35
	}
L35:
	;
	goto L28
L36:
	;
	v161 = v154
	v168 = v152
	goto L26
L37:
	;
	v185 = *(*int32)(unsafe.Add(mBase, uint32(v177+v103)))
	v196 = v168 + base.B2i32(v185 == v181)
	goto L22
L38:
	;
	v219 = *(*int32)(unsafe.Add(mBase, uint32(v85)))
	if v219 != 0 {
		v85 = v219
		v89 = v215
		v90 = v216
		v93 = v217
		v95 = v218
		goto L20
	} else {
		goto L43
	}
L39:
	;
	if v196 != v93 {
		v215 = v89
		v216 = v90
		v217 = v93
		v218 = v95
		goto L38
	} else {
		goto L42
	}
L40:
	;
	goto L41
L41:
	;
	v215 = v85
	v216 = v85
	v217 = v196
	v218 = int32(1)
	goto L38
L42:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v89))) = v85
	v215 = v85
	v216 = v90
	v217 = v93
	v218 = v95 + int32(1)
	goto L38
L43:
	;
	goto L21
L44:
	;
	v231 = v216
	goto L19
L45:
	;
	v246 = v240
	goto L48
L46:
	;
	goto L47
L47:
	;
	v290 = int32(0)
	if v231 == v290 {
		goto L53
	} else {
		goto L54
	}
L48:
	;
	v267 = *(*int32)(unsafe.Add(mBase, uint32(v19+int32(240)+v246<<(uint(int32(2))%32))))
	v268 = F_TypeCategory(m, v267)
	mBase = m.M
	v269 = m.ExcPending
	if v269 != 0 {
		goto L13
	} else {
		goto L50
	}
L49:
	;
	goto L47
L50:
	;
	*(*uint8)(unsafe.Add(mBase, uint32(v19+int32(128)+v246))) = uint8(v268)
	v272 = v246 + int32(1)
	if v272 != l0 {
		v246 = v272
		goto L48
	} else {
		goto L51
	}
L51:
	;
	goto L49
L52:
	;
	v414 = int32(0)
	if base.B2i32(v77 == v414)|base.B2i32(l0 <= v414) != 0 {
		v1002 = v414
		goto L16
	} else {
		goto L78
	}
L53:
	;
	v408 = int32(0)
	goto L52
L54:
	;
	goto L55
L55:
	;
	v294 = int32(0)
	v300 = v294
	v303 = v290
	v304 = v294
	v305 = v231
	v308 = v231
	goto L56
L56:
	;
	v314 = int32(0)
	if base.B2i32(l0 <= v294) == v314 {
		goto L58
	} else {
		goto L59
	}
L57:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v391))) = int32(0)
	if v389 == int32(1) {
		v1002 = v392
		goto L16
	} else {
		goto L77
	}
L58:
	;
	v321 = v314
	v323 = int32(0)
	goto L61
L59:
	;
	v363 = v314
	goto L60
L60:
	;
	v378 = int32(0)
	if base.B2i32(v304 == v378)|base.B2i32(v303 < v363) == v378 {
		goto L72
	} else {
		goto L73
	}
L61:
	;
	v337 = v323 << (uint(int32(2)) % 32)
	v341 = *(*int32)(unsafe.Add(mBase, uint32(v337+(v19+int32(240)))))
	if v341 == int32(705) {
		v357 = v321
		goto L63
	} else {
		goto L64
	}
L62:
	;
	v363 = v357
	goto L60
L63:
	;
	v360 = v323 + int32(1)
	if v360 != l0 {
		v321 = v357
		v323 = v360
		goto L61
	} else {
		goto L70
	}
L64:
	;
	v345 = *(*int32)(unsafe.Add(mBase, uint32(v337+(v305+int32(32)))))
	if v341 != v345 {
		goto L65
	} else {
		goto L66
	}
L65:
	;
	v350 = int32(*(*int8)(unsafe.Add(mBase, uint32(v19+int32(128)+v323))))
	v351 = F_IsPreferredType(m, v350, v345)
	mBase = m.M
	v352 = m.ExcPending
	if v352 != 0 {
		goto L13
	} else {
		goto L68
	}
L66:
	;
	goto L67
L67:
	;
	v357 = v321 + int32(1)
	goto L63
L68:
	;
	if v351 == int32(0) {
		v357 = v321
		goto L63
	} else {
		goto L69
	}
L69:
	;
	goto L67
L70:
	;
	goto L62
L71:
	;
	v393 = *(*int32)(unsafe.Add(mBase, uint32(v305)))
	if v393 != 0 {
		v300 = v389
		v303 = v390
		v304 = v391
		v305 = v393
		v308 = v392
		goto L56
	} else {
		goto L76
	}
L72:
	;
	if v363 != v303 {
		v389 = v300
		v390 = v303
		v391 = v304
		v392 = v308
		goto L71
	} else {
		goto L75
	}
L73:
	;
	goto L74
L74:
	;
	v389 = int32(1)
	v390 = v363
	v391 = v305
	v392 = v305
	goto L71
L75:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v304))) = v305
	v389 = v300 + int32(1)
	v390 = v303
	v391 = v305
	v392 = v308
	goto L71
L76:
	;
	goto L57
L77:
	;
	v408 = v392
	goto L52
L78:
	;
	v420 = int32(0)
	v424 = v420
	v430 = v420
	goto L80
L79:
	;
	v743 = int32(0)
	if base.B2i32(l0 <= v77)|base.B2i32(l0 <= v743) != 0 {
		v1002 = v743
		goto L16
	} else {
		goto L137
	}
L80:
	;
	v439 = v424 << (uint(int32(2)) % 32)
	v443 = *(*int32)(unsafe.Add(mBase, uint32(v439+(v19+int32(240)))))
	if v443 != int32(705) {
		v618 = v430
		goto L82
	} else {
		goto L83
	}
L81:
	;
	v629 = int32(0)
	v631 = v618 & int32(1)
	if v631&base.B2i32(v408 != v629) == v629 {
		goto L110
	} else {
		goto L111
	}
L82:
	;
	v627 = v424 + int32(1)
	if v627 != l0 {
		v424 = v627
		v430 = v618
		goto L80
	} else {
		goto L109
	}
L83:
	;
	v446 = int32(0)
	v449 = v19 + int32(16) + v424
	*(*uint8)(unsafe.Add(mBase, uint32(v449))) = uint8(v446)
	v454 = v19 + int32(128) + v424
	*(*uint8)(unsafe.Add(mBase, uint32(v454))) = uint8(v446)
	if v408 == v446 {
		v618 = int32(1)
		goto L82
	} else {
		goto L84
	}
L84:
	;
	v465 = v408
	v467 = v446
	v468 = v446
	v471 = v446
	goto L86
L85:
	;
	if v594&int32(255) == int32(83) {
		v618 = int32(1)
		goto L82
	} else {
		goto L108
	}
L86:
	;
	v479 = *(*int32)(unsafe.Add(mBase, uint32(v465+v439)+32))
	F_get_type_category_preferred(m, v479, v19+int32(127), v19+int32(126))
	mBase = m.M
	v485 = m.ExcPending
	if v485 != 0 {
		goto L13
	} else {
		goto L88
	}
L87:
	;
	v584 = int32(1)
	if v571&v584 == int32(0) {
		v618 = v584
		goto L82
	} else {
		goto L107
	}
L88:
	;
	v486 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v19)+127)))
	v488 = v467 & int32(255)
	if v488 != 0 {
		goto L91
	} else {
		goto L92
	}
L89:
	;
	*(*uint8)(unsafe.Add(mBase, uint32(v449))) = uint8(v581)
	v583 = *(*int32)(unsafe.Add(mBase, uint32(v568)))
	if v583 != 0 {
		v465 = v583
		v467 = v570
		v468 = v571
		v471 = v581
		goto L86
	} else {
		goto L106
	}
L90:
	;
	v563 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v19)+126)))
	v568 = v550
	v570 = v467
	v571 = v553
	v581 = v471 | v563
	goto L89
L91:
	;
	if v486 == v488 {
		v550 = v465
		v553 = v468
		goto L90
	} else {
		goto L94
	}
L92:
	;
	goto L93
L93:
	;
	*(*uint8)(unsafe.Add(mBase, uint32(v454))) = uint8(v486)
	v546 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v19)+126)))
	v568 = v465
	v570 = v486
	v571 = v468
	v581 = v546
	goto L89
L94:
	;
	if v486 != int32(83) {
		goto L95
	} else {
		goto L96
	}
L95:
	;
	v495 = v465
	goto L98
L96:
	;
	v528 = v465
	v531 = v468
	goto L97
L97:
	;
	v541 = int32(83)
	*(*uint8)(unsafe.Add(mBase, uint32(v454))) = uint8(v541)
	v544 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v19)+126)))
	v568 = v528
	v570 = v541
	v571 = v531
	v581 = v544
	goto L89
L98:
	;
	v508 = *(*int32)(unsafe.Add(mBase, uint32(v495)))
	if v508 == int32(0) {
		v594 = v467
		goto L85
	} else {
		goto L100
	}
L99:
	;
	v528 = v508
	v531 = int32(1)
	goto L97
L100:
	;
	v512 = *(*int32)(unsafe.Add(mBase, uint32(v508+v439)+32))
	F_get_type_category_preferred(m, v512, v19+int32(127), v19+int32(126))
	mBase = m.M
	v518 = m.ExcPending
	if v518 != 0 {
		goto L13
	} else {
		goto L101
	}
L101:
	;
	v519 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v19)+127)))
	if v488 == v519 {
		goto L102
	} else {
		goto L103
	}
L102:
	;
	v550 = v508
	v553 = int32(1)
	goto L90
L103:
	;
	goto L104
L104:
	;
	if v519 != int32(83) {
		v495 = v508
		goto L98
	} else {
		goto L105
	}
L105:
	;
	goto L99
L106:
	;
	goto L87
L107:
	;
	v594 = v570
	goto L85
L108:
	;
	v733 = v408
	goto L79
L109:
	;
	goto L81
L110:
	;
	if v631 != 0 {
		goto L113
	} else {
		goto L114
	}
L111:
	;
	goto L112
L112:
	;
	__phi641 = v408
	__phi645 = int32(0)
	__phi646 = v408
	__phi649 = v629
	v641 = __phi641
	v645 = __phi645
	v646 = __phi646
	v649 = __phi649
	goto L117
L113:
	;
	v638 = int32(0)
	goto L115
L114:
	;
	v638 = v408
	goto L115
L115:
	;
	v733 = v638
	goto L79
L116:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v722))) = int32(0)
	if v721 == int32(1) {
		v1002 = v646
		goto L16
	} else {
		goto L136
	}
L117:
	;
	v662 = int32(0)
	goto L120
L118:
	;
	if v645 == int32(1) {
		v1002 = v408
		goto L16
	} else {
		goto L135
	}
L119:
	;
	v714 = *(*int32)(unsafe.Add(mBase, uint32(v641)))
	if v649 != 0 {
		goto L130
	} else {
		goto L131
	}
L120:
	;
	v676 = v662 << (uint(int32(2)) % 32)
	v680 = *(*int32)(unsafe.Add(mBase, uint32(v676+(v19+int32(240)))))
	if v680 != int32(705) {
		goto L122
	} else {
		goto L123
	}
L121:
	;
	v712 = v645 + int32(1)
	v713 = *(*int32)(unsafe.Add(mBase, uint32(v641)))
	if v713 != 0 {
		__phi641 = v713
		__phi645 = v712
		__phi649 = v641
		v641 = __phi641
		v645 = __phi645
		v649 = __phi649
		goto L117
	} else {
		goto L129
	}
L122:
	;
	v709 = v662 + int32(1)
	if v709 != l0 {
		v662 = v709
		goto L120
	} else {
		goto L128
	}
L123:
	;
	v684 = *(*int32)(unsafe.Add(mBase, uint32(v676+(v641+int32(32)))))
	F_get_type_category_preferred(m, v684, v19+int32(127), v19+int32(126))
	mBase = m.M
	v690 = m.ExcPending
	if v690 != 0 {
		goto L13
	} else {
		goto L124
	}
L124:
	;
	v691 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v19)+127)))
	v695 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v19+int32(128)+v662))))
	if v691 != v695 {
		goto L119
	} else {
		goto L125
	}
L125:
	;
	v700 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v19+int32(16)+v662))))
	if v700 != int32(1) {
		goto L122
	} else {
		goto L126
	}
L126:
	;
	v703 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v19)+126)))
	if v703&int32(1) == int32(0) {
		goto L119
	} else {
		goto L127
	}
L127:
	;
	goto L122
L128:
	;
	goto L121
L129:
	;
	v721 = v712
	v722 = v641
	goto L116
L130:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v649))) = v714
	if v714 != 0 {
		__phi641 = v714
		v641 = __phi641
		goto L117
	} else {
		goto L133
	}
L131:
	;
	goto L132
L132:
	;
	if v714 != 0 {
		__phi641 = v714
		__phi646 = v714
		__phi649 = int32(0)
		v641 = __phi641
		v646 = __phi646
		v649 = __phi649
		goto L117
	} else {
		goto L134
	}
L133:
	;
	v721 = v645
	v722 = v649
	goto L116
L134:
	;
	goto L118
L135:
	;
	v733 = v408
	goto L79
L136:
	;
	v733 = v646
	goto L79
L137:
	;
	v752 = int32(705)
	v756 = v743
	goto L140
L138:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v984))) = int32(0)
	v1002 = v984
	goto L16
L139:
	;
	v1002 = int32(0)
	goto L16
L140:
	;
	v770 = *(*int32)(unsafe.Add(mBase, uint32(v19+int32(240)+v756<<(uint(int32(2))%32))))
	if v770 == int32(705) {
		v776 = v752
		goto L142
	} else {
		goto L143
	}
L141:
	;
	v780 = int32(0)
	if v776 == int32(705) {
		v1002 = v780
		goto L16
	} else {
		goto L149
	}
L142:
	;
	v778 = v756 + int32(1)
	if v778 != l0 {
		v752 = v776
		v756 = v778
		goto L140
	} else {
		goto L148
	}
L143:
	;
	if v752 == int32(705) {
		goto L144
	} else {
		goto L145
	}
L144:
	;
	v776 = v770
	goto L142
L145:
	;
	goto L146
L146:
	;
	if v770 != v752 {
		goto L139
	} else {
		goto L147
	}
L147:
	;
	v776 = v752
	goto L142
L148:
	;
	goto L141
L149:
	;
	if l0 <= int32(0) {
		goto L150
	} else {
		goto L151
	}
L150:
	;
	if v733 == int32(0) {
		v1002 = v780
		goto L16
	} else {
		goto L162
	}
L151:
	;
	v786 = l0 & int32(7)
	v787 = int32(0)
	if base.Ui32(int32(8)) <= base.Ui32(l0) {
		goto L152
	} else {
		goto L153
	}
L152:
	;
	v803 = int32(0)
	v805 = v787
	goto L155
L153:
	;
	v841 = v787
	goto L154
L154:
	;
	v847 = v787
	v857 = v841
	goto L159
L155:
	;
	v814 = v19 + int32(240) + v805<<(uint(int32(2))%32)
	*(*int32)(unsafe.Add(mBase, uint32(v814)+28)) = v776
	*(*int32)(unsafe.Add(mBase, uint32(v814)+24)) = v776
	*(*int32)(unsafe.Add(mBase, uint32(v814)+20)) = v776
	*(*int32)(unsafe.Add(mBase, uint32(v814)+16)) = v776
	*(*int32)(unsafe.Add(mBase, uint32(v814)+12)) = v776
	*(*int32)(unsafe.Add(mBase, uint32(v814)+8)) = v776
	*(*int32)(unsafe.Add(mBase, uint32(v814)+4)) = v776
	*(*int32)(unsafe.Add(mBase, uint32(v814))) = v776
	v823 = int32(8)
	v824 = v805 + v823
	v826 = v803 + v823
	if v826 != l0&int32(2147483640) {
		v803 = v826
		v805 = v824
		goto L155
	} else {
		goto L157
	}
L156:
	;
	if v786 == int32(0) {
		goto L150
	} else {
		goto L158
	}
L157:
	;
	goto L156
L158:
	;
	v841 = v824
	goto L154
L159:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v19+int32(240)+v857<<(uint(int32(2))%32)))) = v776
	v868 = int32(1)
	v871 = v847 + v868
	if v871 != v786 {
		v847 = v871
		v857 = v857 + v868
		goto L159
	} else {
		goto L161
	}
L160:
	;
	goto L150
L161:
	;
	goto L160
L162:
	;
	v897 = v733
	goto L165
L163:
	;
	if v912 != 0 {
		v984 = v959
		goto L138
	} else {
		goto L178
	}
L164:
	;
	v916 = *(*int32)(unsafe.Add(mBase, uint32(v897)))
	if v916 == int32(0) {
		goto L170
	} else {
		goto L171
	}
L165:
	;
	v912 = F_can_coerce_type(m, l0, v19+int32(240), v897+int32(32), int32(0))
	mBase = m.M
	v913 = m.ExcPending
	if v913 != 0 {
		goto L13
	} else {
		goto L167
	}
L166:
	;
	v959 = int32(0)
	goto L163
L167:
	;
	if v912 != 0 {
		goto L164
	} else {
		goto L168
	}
L168:
	;
	v914 = *(*int32)(unsafe.Add(mBase, uint32(v897)))
	if v914 != 0 {
		v897 = v914
		goto L165
	} else {
		goto L169
	}
L169:
	;
	goto L166
L170:
	;
	v984 = v897
	goto L138
L171:
	;
	goto L172
L172:
	;
	v927 = v916
	goto L173
L173:
	;
	v940 = F_can_coerce_type(m, l0, v19+int32(240), v927+int32(32), int32(0))
	mBase = m.M
	v941 = m.ExcPending
	if v941 != 0 {
		goto L13
	} else {
		goto L175
	}
L174:
	;
	v959 = v897
	goto L163
L175:
	;
	if v940 != 0 {
		v1002 = v780
		goto L16
	} else {
		goto L176
	}
L176:
	;
	v942 = *(*int32)(unsafe.Add(mBase, uint32(v927)))
	if v942 != 0 {
		v927 = v942
		goto L173
	} else {
		goto L177
	}
L177:
	;
	goto L174
L178:
	;
	goto L139
L179:
	;
	F_errcode(m, int32(50856197))
	mBase = m.M
	v1021 = m.ExcPending
	if v1021 != 0 {
		goto L13
	} else {
		goto L180
	}
L180:
	;
	v1022 = int32(100)
	*(*int32)(unsafe.Add(mBase, uint32(v19))) = v1022
	F_errmsg_plural(m, int32(_a_F_func_select_candidate_0), int32(_a_F_func_select_candidate_1), v1022, v19)
	mBase = m.M
	v1028 = m.ExcPending
	if v1028 != 0 {
		goto L13
	} else {
		goto L181
	}
L181:
	;
	F_errfinish(m, int32(_a_F_func_select_candidate_2), int32(1212), int32(_a_F_func_select_candidate_3))
	mBase = m.M
	v1033 = m.ExcPending
	if v1033 != 0 {
		goto L13
	} else {
		goto L182
	}
L182:
	;
	base.Wasm_trap_unreachable()
	for {
	}
}
func F_func_volatile(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v4 int32
	_ = v4
	var v6 int32
	_ = v6
	var v10 int32
	_ = v10
	var v13 int32
	_ = v13
	var v19 int32
	_ = v19
	var v23 int32
	_ = v23
	var v28 int32
	_ = v28
	var v29 int32
	_ = v29
	var v30 int32
	_ = v30
	var v32 int32
	_ = v32
	var v34 int32
	_ = v34
	v4 = m.G0
	v6 = v4 - int32(16)
	m.G0 = v6
	v10 = F_SearchSysCache1(m, int32(47), base.I64_extend_i32_u(l0))
	mBase = m.M
	v13 = m.ExcPending
	if v13 != 0 {
		return int32(0)
	} else {
		if v10 == int32(0) {
			F_errstart_cold(m, int32(21), int32(0))
			mBase = m.M
			v19 = m.ExcPending
			if v19 != 0 {
				return int32(0)
			} else {
				*(*int32)(unsafe.Add(mBase, uint32(v6))) = l0
				F_errmsg_internal(m, int32(_a_F_func_volatile_0), v6)
				mBase = m.M
				v23 = m.ExcPending
				if v23 != 0 {
					return int32(0)
				} else {
					F_errfinish(m, int32(_a_F_func_volatile_1), int32(2099), int32(_a_F_func_volatile_2))
					mBase = m.M
					v28 = m.ExcPending
					if v28 != 0 {
						return int32(0)
					} else {
						base.Wasm_trap_unreachable()
						for {
						}
					}
				}
			}
		} else {
			v29 = *(*int32)(unsafe.Add(mBase, uint32(v10)+16))
			v30 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v29)+22)))
			v32 = int32(*(*int8)(unsafe.Add(mBase, uint32(v29+v30)+101)))
			F_ReleaseCatCache(m, v10)
			mBase = m.M
			v34 = m.ExcPending
			if v34 != 0 {
				return int32(0)
			} else {
				m.G0 = v6 + int32(16)
				return v32
			}
		}
	}
}
func F_get_func_variadictype(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v4 int32
	_ = v4
	var v6 int32
	_ = v6
	var v10 int32
	_ = v10
	var v13 int32
	_ = v13
	var v19 int32
	_ = v19
	var v23 int32
	_ = v23
	var v28 int32
	_ = v28
	var v29 int32
	_ = v29
	var v30 int32
	_ = v30
	var v32 int32
	_ = v32
	var v34 int32
	_ = v34
	v4 = m.G0
	v6 = v4 - int32(16)
	m.G0 = v6
	v10 = F_SearchSysCache1(m, int32(47), base.I64_extend_i32_u(l0))
	mBase = m.M
	v13 = m.ExcPending
	if v13 != 0 {
		return int32(0)
	} else {
		if v10 == int32(0) {
			F_errstart_cold(m, int32(21), int32(0))
			mBase = m.M
			v19 = m.ExcPending
			if v19 != 0 {
				return int32(0)
			} else {
				*(*int32)(unsafe.Add(mBase, uint32(v6))) = l0
				F_errmsg_internal(m, int32(_a_F_get_func_variadictype_0), v6)
				mBase = m.M
				v23 = m.ExcPending
				if v23 != 0 {
					return int32(0)
				} else {
					F_errfinish(m, int32(_a_F_get_func_variadictype_1), int32(2042), int32(_a_F_get_func_variadictype_2))
					mBase = m.M
					v28 = m.ExcPending
					if v28 != 0 {
						return int32(0)
					} else {
						base.Wasm_trap_unreachable()
						for {
						}
					}
				}
			}
		} else {
			v29 = *(*int32)(unsafe.Add(mBase, uint32(v10)+16))
			v30 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v29)+22)))
			v32 = *(*int32)(unsafe.Add(mBase, uint32(v29+v30)+88))
			F_ReleaseCatCache(m, v10)
			mBase = m.M
			v34 = m.ExcPending
			if v34 != 0 {
				return int32(0)
			} else {
				m.G0 = v6 + int32(16)
				return v32
			}
		}
	}
}
func F_makeFuncExpr(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32) int32 {
	mBase := m.M
	_ = mBase
	var v8 int32
	_ = v8
	var v11 int32
	_ = v11
	var v16 int32
	_ = v16
	v8 = F_palloc0(m, int32(36))
	mBase = m.M
	v11 = m.ExcPending
	if v11 != 0 {
		return int32(0)
	} else {
		*(*int32)(unsafe.Add(mBase, uint32(v8)+32)) = int32(-1)
		*(*int32)(unsafe.Add(mBase, uint32(v8)+28)) = l2
		*(*int32)(unsafe.Add(mBase, uint32(v8)+24)) = l3
		v16 = int32(0)
		*(*int32)(unsafe.Add(mBase, uint32(v8)+20)) = v16
		*(*int32)(unsafe.Add(mBase, uint32(v8)+16)) = l4
		*(*uint16)(unsafe.Add(mBase, uint32(v8)+12)) = uint16(v16)
		*(*int32)(unsafe.Add(mBase, uint32(v8)+8)) = l1
		*(*int32)(unsafe.Add(mBase, uint32(v8)+4)) = l0
		*(*int32)(unsafe.Add(mBase, uint32(v8))) = int32(15)
		return v8
	}
}
func F_parse_func_options(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32, l5 int32, l6 int32, l7 int32) {
	mBase := m.M
	_ = mBase
	var v9 int32
	_ = v9
	var v13 int32
	_ = v13
	var v15 int32
	_ = v15
	var v31 int32
	_ = v31
	var v45 int32
	_ = v45
	var v46 int32
	_ = v46
	var v50 int32
	_ = v50
	var v51 int32
	_ = v51
	var v52 int32
	_ = v52
	var v55 int32
	_ = v55
	var v58 int32
	_ = v58
	var v61 int32
	_ = v61
	var v62 int32
	_ = v62
	var v65 int32
	_ = v65
	var v66 int32
	_ = v66
	var v69 int32
	_ = v69
	var v76 int32
	_ = v76
	var v77 int32
	_ = v77
	var v81 int32
	_ = v81
	var v84 int32
	_ = v84
	var v86 int32
	_ = v86
	var v91 int32
	_ = v91
	var v94 int32
	_ = v94
	var v95 int32
	_ = v95
	var v96 int32
	_ = v96
	var v97 int32
	_ = v97
	var v101 int32
	_ = v101
	var v104 int32
	_ = v104
	var v107 int32
	_ = v107
	var v110 int32
	_ = v110
	var v111 int32
	_ = v111
	var v114 int32
	_ = v114
	var v115 int32
	_ = v115
	var v118 int32
	_ = v118
	var v125 int32
	_ = v125
	var v126 int32
	_ = v126
	var v130 int32
	_ = v130
	var v133 int32
	_ = v133
	var v135 int32
	_ = v135
	var v146 int32
	_ = v146
	var v147 int32
	_ = v147
	var v149 int32
	_ = v149
	var v152 int32
	_ = v152
	var v155 int32
	_ = v155
	var v158 int32
	_ = v158
	var v159 int32
	_ = v159
	var v162 int32
	_ = v162
	var v163 int32
	_ = v163
	var v166 int32
	_ = v166
	var v173 int32
	_ = v173
	var v174 int32
	_ = v174
	var v176 int32
	_ = v176
	var v179 int32
	_ = v179
	var v181 int32
	_ = v181
	var v193 int32
	_ = v193
	var v194 int32
	_ = v194
	var v195 int32
	_ = v195
	var v196 int32
	_ = v196
	var v199 int32
	_ = v199
	var v204 int32
	_ = v204
	var v205 int32
	_ = v205
	var v223 int32
	_ = v223
	var v227 int32
	_ = v227
	var v230 int32
	_ = v230
	var v231 int32
	_ = v231
	var v232 int32
	_ = v232
	var v233 int32
	_ = v233
	var v239 int32
	_ = v239
	var v244 int32
	_ = v244
	var v248 int32
	_ = v248
	var v251 int32
	_ = v251
	var v252 int32
	_ = v252
	var v253 int32
	_ = v253
	var v254 int32
	_ = v254
	var v262 int32
	_ = v262
	var v267 int32
	_ = v267
	var v271 int32
	_ = v271
	var v272 int32
	_ = v272
	var v278 int32
	_ = v278
	var v283 int32
	_ = v283
	v9 = int32(0)
	v13 = m.G0
	v15 = v13 - int32(48)
	m.G0 = v15
	*(*uint8)(unsafe.Add(mBase, uint32(l2))) = uint8(v9)
	*(*uint8)(unsafe.Add(mBase, uint32(l4))) = uint8(v9)
	*(*uint8)(unsafe.Add(mBase, uint32(l6))) = uint8(v9)
	*(*int32)(unsafe.Add(mBase, uint32(l3))) = v9
	*(*int32)(unsafe.Add(mBase, uint32(l5))) = v9
	*(*int32)(unsafe.Add(mBase, uint32(l7))) = v9
	if l1 == v9 {
		goto L5
	} else {
		goto L6
	}
L1:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v271 = m.ExcPending
	if v271 != 0 {
		goto L25
	} else {
		goto L72
	}
L2:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v248 = m.ExcPending
	if v248 != 0 {
		goto L25
	} else {
		goto L67
	}
L3:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v227 = m.ExcPending
	if v227 != 0 {
		goto L25
	} else {
		goto L62
	}
L4:
	;
	F_errorConflictingDefElem(m, v50, l0)
	mBase = m.M
	v223 = m.ExcPending
	if v223 != 0 {
		goto L25
	} else {
		goto L61
	}
L5:
	;
	m.G0 = v15 + int32(48)
	return
L6:
	;
	v31 = *(*int32)(unsafe.Add(mBase, uint32(l1)+4))
	if v31 <= int32(0) {
		goto L5
	} else {
		goto L7
	}
L7:
	;
	v45 = v9
	goto L8
L8:
	;
	v46 = *(*int32)(unsafe.Add(mBase, uint32(l1)+12))
	v50 = *(*int32)(unsafe.Add(mBase, uint32(v46+v45<<(uint(int32(2))%32))))
	v51 = *(*int32)(unsafe.Add(mBase, uint32(v50)+8))
	v52 = int32(_a_F_parse_func_options_0)
	v55 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v51))))
	v58 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_parse_func_options[0])))
	if base.B2i32(v55 == int32(0))|base.B2i32(v55 != v58) != 0 {
		v76 = v55
		v77 = v58
		goto L12
	} else {
		goto L13
	}
L9:
	;
	goto L5
L10:
	;
	v204 = v45 + int32(1)
	v205 = *(*int32)(unsafe.Add(mBase, uint32(l1)+4))
	if v204 < v205 {
		v45 = v204
		goto L8
	} else {
		goto L60
	}
L11:
	;
	if v76-v77 == int32(0) {
		goto L18
	} else {
		goto L19
	}
L12:
	;
	goto L11
L13:
	;
	v61 = v51
	v62 = v52
	goto L14
L14:
	;
	v65 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v62)+1)))
	v66 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v61)+1)))
	if v66 == int32(0) {
		v76 = v66
		v77 = v65
		goto L12
	} else {
		goto L16
	}
L15:
	;
	v76 = v66
	v77 = v65
	goto L12
L16:
	;
	v69 = int32(1)
	if v66 == v65 {
		v61 = v61 + v69
		v62 = v62 + v69
		goto L14
	} else {
		goto L17
	}
L17:
	;
	goto L15
L18:
	;
	v81 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l2))))
	if v81 == int32(1) {
		goto L4
	} else {
		goto L21
	}
L19:
	;
	goto L20
L20:
	;
	v101 = int32(_a_F_parse_func_options_1)
	v104 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v51))))
	v107 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_parse_func_options[1])))
	if base.B2i32(v104 == int32(0))|base.B2i32(v104 != v107) != 0 {
		v125 = v104
		v126 = v107
		goto L30
	} else {
		goto L31
	}
L21:
	;
	v84 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(l2))) = uint8(v84)
	v86 = *(*int32)(unsafe.Add(mBase, uint32(v50)+12))
	if v86 == int32(0) {
		goto L22
	} else {
		goto L23
	}
L22:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l3))) = int32(0)
	goto L10
L23:
	;
	goto L24
L24:
	;
	v91 = int32(0)
	v94 = F_LookupFuncName(m, v86, v91, v91, v91)
	mBase = m.M
	v95 = m.ExcPending
	if v95 != 0 {
		goto L25
	} else {
		goto L26
	}
L25:
	;
	return
L26:
	;
	v96 = F_get_func_rettype(m, v94)
	mBase = m.M
	v97 = m.ExcPending
	if v97 != 0 {
		goto L25
	} else {
		goto L27
	}
L27:
	;
	if v96 != int32(3115) {
		goto L3
	} else {
		goto L28
	}
L28:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l3))) = v94
	goto L10
L29:
	;
	if v125-v126 == int32(0) {
		goto L36
	} else {
		goto L37
	}
L30:
	;
	goto L29
L31:
	;
	v110 = v51
	v111 = v101
	goto L32
L32:
	;
	v114 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v111)+1)))
	v115 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v110)+1)))
	if v115 == int32(0) {
		v125 = v115
		v126 = v114
		goto L30
	} else {
		goto L34
	}
L33:
	;
	v125 = v115
	v126 = v114
	goto L30
L34:
	;
	v118 = int32(1)
	if v115 == v114 {
		v110 = v110 + v118
		v111 = v111 + v118
		goto L32
	} else {
		goto L35
	}
L35:
	;
	goto L33
L36:
	;
	v130 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l4))))
	if v130 == int32(1) {
		goto L4
	} else {
		goto L39
	}
L37:
	;
	goto L38
L38:
	;
	v149 = int32(_a_F_parse_func_options_2)
	v152 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v51))))
	v155 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_parse_func_options[2])))
	if base.B2i32(v152 == int32(0))|base.B2i32(v152 != v155) != 0 {
		v173 = v152
		v174 = v155
		goto L45
	} else {
		goto L46
	}
L39:
	;
	v133 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(l4))) = uint8(v133)
	v135 = *(*int32)(unsafe.Add(mBase, uint32(v50)+12))
	if v135 == int32(0) {
		goto L40
	} else {
		goto L41
	}
L40:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l5))) = int32(0)
	goto L10
L41:
	;
	goto L42
L42:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v15)+36)) = int64(111669150705)
	v146 = F_LookupFuncName(m, v135, int32(2), v15+int32(36), int32(0))
	mBase = m.M
	v147 = m.ExcPending
	if v147 != 0 {
		goto L25
	} else {
		goto L43
	}
L43:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l5))) = v146
	goto L10
L44:
	;
	if v173-v174 != 0 {
		goto L1
	} else {
		goto L51
	}
L45:
	;
	goto L44
L46:
	;
	v158 = v51
	v159 = v149
	goto L47
L47:
	;
	v162 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v159)+1)))
	v163 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v158)+1)))
	if v163 == int32(0) {
		v173 = v163
		v174 = v162
		goto L45
	} else {
		goto L49
	}
L48:
	;
	v173 = v163
	v174 = v162
	goto L45
L49:
	;
	v166 = int32(1)
	if v163 == v162 {
		v158 = v158 + v166
		v159 = v159 + v166
		goto L47
	} else {
		goto L50
	}
L50:
	;
	goto L48
L51:
	;
	v176 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l6))))
	if v176 == int32(1) {
		goto L4
	} else {
		goto L52
	}
L52:
	;
	v179 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(l6))) = uint8(v179)
	v181 = *(*int32)(unsafe.Add(mBase, uint32(v50)+12))
	if v181 == int32(0) {
		goto L54
	} else {
		goto L55
	}
L53:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l7))) = v199
	goto L10
L54:
	;
	v199 = int32(0)
	goto L53
L55:
	;
	goto L56
L56:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v15)+44)) = int32(2281)
	*(*int64)(unsafe.Add(mBase, uint32(v15)+36)) = int64(111669149722)
	v193 = F_LookupFuncName(m, v181, int32(3), v15+int32(36), int32(0))
	mBase = m.M
	v194 = m.ExcPending
	if v194 != 0 {
		goto L25
	} else {
		goto L57
	}
L57:
	;
	v195 = F_get_func_rettype(m, v193)
	mBase = m.M
	v196 = m.ExcPending
	if v196 != 0 {
		goto L25
	} else {
		goto L58
	}
L58:
	;
	if v195 != int32(25) {
		goto L2
	} else {
		goto L59
	}
L59:
	;
	v199 = v193
	goto L53
L60:
	;
	goto L9
L61:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L62:
	;
	F_errcode(m, int32(151027844))
	mBase = m.M
	v230 = m.ExcPending
	if v230 != 0 {
		goto L25
	} else {
		goto L63
	}
L63:
	;
	v231 = *(*int32)(unsafe.Add(mBase, uint32(v50)+12))
	v232 = F_NameListToString(m, v231)
	mBase = m.M
	v233 = m.ExcPending
	if v233 != 0 {
		goto L25
	} else {
		goto L64
	}
L64:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v15)+4)) = int32(_a_F_parse_func_options_3)
	*(*int32)(unsafe.Add(mBase, uint32(v15))) = v232
	F_errmsg(m, int32(_a_F_parse_func_options_4), v15)
	mBase = m.M
	v239 = m.ExcPending
	if v239 != 0 {
		goto L25
	} else {
		goto L65
	}
L65:
	;
	F_errfinish(m, int32(_a_F_parse_func_options_5), int32(501), int32(_a_F_parse_func_options_6))
	mBase = m.M
	v244 = m.ExcPending
	if v244 != 0 {
		goto L25
	} else {
		goto L66
	}
L66:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L67:
	;
	F_errcode(m, int32(151027844))
	mBase = m.M
	v251 = m.ExcPending
	if v251 != 0 {
		goto L25
	} else {
		goto L68
	}
L68:
	;
	v252 = *(*int32)(unsafe.Add(mBase, uint32(v50)+12))
	v253 = F_NameListToString(m, v252)
	mBase = m.M
	v254 = m.ExcPending
	if v254 != 0 {
		goto L25
	} else {
		goto L69
	}
L69:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v15)+20)) = int32(_a_F_parse_func_options_7)
	*(*int32)(unsafe.Add(mBase, uint32(v15)+16)) = v253
	F_errmsg(m, int32(_a_F_parse_func_options_4), v15+int32(16))
	mBase = m.M
	v262 = m.ExcPending
	if v262 != 0 {
		goto L25
	} else {
		goto L70
	}
L70:
	;
	F_errfinish(m, int32(_a_F_parse_func_options_5), int32(549), int32(_a_F_parse_func_options_8))
	mBase = m.M
	v267 = m.ExcPending
	if v267 != 0 {
		goto L25
	} else {
		goto L71
	}
L71:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L72:
	;
	v272 = *(*int32)(unsafe.Add(mBase, uint32(v50)+8))
	*(*int32)(unsafe.Add(mBase, uint32(v15)+32)) = v272
	F_errmsg_internal(m, int32(_a_F_parse_func_options_9), v15+int32(32))
	mBase = m.M
	v278 = m.ExcPending
	if v278 != 0 {
		goto L25
	} else {
		goto L73
	}
L73:
	;
	F_errfinish(m, int32(_a_F_parse_func_options_5), int32(600), int32(_a_F_parse_func_options_10))
	mBase = m.M
	v283 = m.ExcPending
	if v283 != 0 {
		goto L25
	} else {
		goto L74
	}
L74:
	;
	base.Wasm_trap_unreachable()
	for {
	}
}
