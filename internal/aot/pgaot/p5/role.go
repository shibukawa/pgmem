package p5

import (
	base "github.com/shibukawa/pgmem/internal/aot/pgaot/base"
	"unsafe"
)

func F_AddRoleMems(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32, l5 int32, l6 int32) {
	mBase := m.M
	_ = mBase
	var v19 int32
	_ = v19
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
	var v35 int32
	_ = v35
	var v41 int32
	_ = v41
	var v54 int32
	_ = v54
	var v58 int32
	_ = v58
	var v60 int32
	_ = v60
	var v64 int32
	_ = v64
	var v67 int32
	_ = v67
	var v71 int32
	_ = v71
	var v74 int32
	_ = v74
	var v76 int32
	_ = v76
	var v81 int32
	_ = v81
	var v86 int32
	_ = v86
	var v87 int32
	_ = v87
	var v88 int32
	_ = v88
	var v90 int32
	_ = v90
	var v91 int32
	_ = v91
	var v96 int32
	_ = v96
	var v97 int32
	_ = v97
	var v103 int32
	_ = v103
	var v106 int32
	_ = v106
	var v107 int32
	_ = v107
	var v108 int32
	_ = v108
	var v115 int32
	_ = v115
	var v120 int32
	_ = v120
	var v123 int32
	_ = v123
	var v124 int32
	_ = v124
	var v125 int32
	_ = v125
	var v134 int32
	_ = v134
	var v153 int32
	_ = v153
	var v154 int32
	_ = v154
	var v156 int32
	_ = v156
	var v165 int32
	_ = v165
	var v176 int32
	_ = v176
	var v181 int32
	_ = v181
	var v195 int32
	_ = v195
	var v199 int32
	_ = v199
	var v203 int32
	_ = v203
	var v206 int32
	_ = v206
	var v209 int32
	_ = v209
	var v214 int32
	_ = v214
	var v230 int32
	_ = v230
	var v231 int32
	_ = v231
	var v232 int32
	_ = v232
	var v234 int32
	_ = v234
	var v239 int32
	_ = v239
	var v240 int32
	_ = v240
	var v241 int32
	_ = v241
	var v243 int32
	_ = v243
	var v245 int32
	_ = v245
	var v264 int32
	_ = v264
	var v265 int32
	_ = v265
	var v267 int32
	_ = v267
	var v295 int32
	_ = v295
	var v309 int32
	_ = v309
	var v311 int32
	_ = v311
	var v313 int32
	_ = v313
	var v314 int32
	_ = v314
	var v315 int32
	_ = v315
	var v316 int32
	_ = v316
	var v317 int32
	_ = v317
	var v319 int32
	_ = v319
	var v322 int32
	_ = v322
	var v345 int32
	_ = v345
	var v348 int32
	_ = v348
	var v355 int32
	_ = v355
	var v360 int32
	_ = v360
	var v362 int32
	_ = v362
	var v391 int32
	_ = v391
	var v404 int32
	_ = v404
	var v408 int32
	_ = v408
	var v410 int32
	_ = v410
	var v414 int32
	_ = v414
	var v417 int32
	_ = v417
	var v421 int32
	_ = v421
	var v424 int32
	_ = v424
	var v427 int32
	_ = v427
	var v428 int32
	_ = v428
	var v429 int32
	_ = v429
	var v431 int64
	_ = v431
	var v449 int32
	_ = v449
	var v450 int32
	_ = v450
	var v451 int32
	_ = v451
	var v452 int32
	_ = v452
	var v453 int32
	_ = v453
	var v454 int32
	_ = v454
	var v455 int32
	_ = v455
	var v460 int32
	_ = v460
	var v461 int32
	_ = v461
	var v463 int32
	_ = v463
	var v467 int32
	_ = v467
	var v473 int32
	_ = v473
	var v474 int32
	_ = v474
	var v476 int32
	_ = v476
	var v480 int32
	_ = v480
	var v486 int32
	_ = v486
	var v487 int32
	_ = v487
	var v489 int32
	_ = v489
	var v495 int32
	_ = v495
	var v496 int32
	_ = v496
	var v497 int32
	_ = v497
	var v498 int32
	_ = v498
	var v500 int32
	_ = v500
	var v501 int32
	_ = v501
	var v509 int32
	_ = v509
	var v514 int32
	_ = v514
	var v517 int32
	_ = v517
	var v521 int32
	_ = v521
	var v522 int32
	_ = v522
	var v523 int32
	_ = v523
	var v525 int32
	_ = v525
	var v527 int32
	_ = v527
	var v530 int32
	_ = v530
	var v533 int32
	_ = v533
	var v534 int32
	_ = v534
	var v537 int32
	_ = v537
	var v538 int32
	_ = v538
	var v540 int32
	_ = v540
	var v543 int32
	_ = v543
	var v548 int32
	_ = v548
	var v549 int32
	_ = v549
	var v555 int32
	_ = v555
	var v556 int32
	_ = v556
	var v558 int32
	_ = v558
	var v561 int32
	_ = v561
	var v567 int32
	_ = v567
	var v571 int32
	_ = v571
	var v578 int32
	_ = v578
	var v582 int32
	_ = v582
	var v587 int32
	_ = v587
	var v595 int32
	_ = v595
	var v596 int32
	_ = v596
	var v600 int32
	_ = v600
	var v602 int32
	_ = v602
	var v609 int32
	_ = v609
	var v615 int32
	_ = v615
	var v618 int32
	_ = v618
	var v625 int32
	_ = v625
	var v630 int32
	_ = v630
	var v634 int32
	_ = v634
	var v637 int32
	_ = v637
	var v638 int32
	_ = v638
	var v639 int32
	_ = v639
	var v645 int32
	_ = v645
	var v650 int32
	_ = v650
	v19 = m.G0
	v21 = v19 - int32(144)
	m.G0 = v21
	v24 = F_check_role_grantor(m, l0, l2, l5, int32(1))
	mBase = m.M
	v25 = m.ExcPending
	if v25 != 0 {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	return
L2:
	;
	v28 = F_table_open(m, int32(1261), int32(3))
	mBase = m.M
	v29 = m.ExcPending
	if v29 != 0 {
		goto L1
	} else {
		goto L3
	}
L3:
	;
	v30 = *(*int32)(unsafe.Add(mBase, uint32(v28)+52))
	F_LockSharedObject(m, int32(1260), l2, int32(4))
	mBase = m.M
	v35 = m.ExcPending
	if v35 != 0 {
		goto L1
	} else {
		goto L4
	}
L4:
	;
	v41 = int32(0)
	goto L10
L5:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v634 = m.ExcPending
	if v634 != 0 {
		goto L1
	} else {
		goto L131
	}
L6:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v615 = m.ExcPending
	if v615 != 0 {
		goto L1
	} else {
		goto L127
	}
L7:
	;
	v391 = int32(0)
	goto L71
L8:
	;
	if l4 == int32(0) {
		v267 = v156
		goto L38
	} else {
		goto L39
	}
L9:
	;
	v123 = F_palloc(m, v88<<(uint(int32(2))%32))
	mBase = m.M
	v124 = m.ExcPending
	if v124 != 0 {
		goto L1
	} else {
		goto L33
	}
L10:
	;
	v54 = int32(0)
	if l3 == v54 {
		v64 = v54
		goto L12
	} else {
		goto L13
	}
L11:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v103 = m.ExcPending
	if v103 != 0 {
		goto L1
	} else {
		goto L28
	}
L12:
	;
	if l4 == int32(0) {
		goto L16
	} else {
		goto L17
	}
L13:
	;
	v58 = *(*int32)(unsafe.Add(mBase, uint32(l3)+4))
	if v58 <= v41 {
		v64 = int32(0)
		goto L12
	} else {
		goto L14
	}
L14:
	;
	v60 = *(*int32)(unsafe.Add(mBase, uint32(l3)+12))
	v64 = v60 + v41<<(uint(int32(2))%32)
	goto L12
L15:
	;
	v90 = *(*int32)(unsafe.Add(mBase, uint32(v64)))
	v91 = *(*int32)(unsafe.Add(mBase, uint32(v74)))
	if v91 == int32(6171) {
		goto L5
	} else {
		goto L25
	}
L16:
	;
	v76 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l6)+4)))
	if v76 != int32(1) {
		goto L7
	} else {
		goto L21
	}
L17:
	;
	v67 = *(*int32)(unsafe.Add(mBase, uint32(l4)+4))
	if v67 <= v41 {
		goto L16
	} else {
		goto L18
	}
L18:
	;
	if v64 == int32(0) {
		goto L16
	} else {
		goto L19
	}
L19:
	;
	v71 = *(*int32)(unsafe.Add(mBase, uint32(l4)+12))
	v74 = v71 + v41<<(uint(int32(2))%32)
	if v74 != 0 {
		goto L15
	} else {
		goto L20
	}
L20:
	;
	goto L16
L21:
	;
	if v24 == int32(10) {
		goto L7
	} else {
		goto L22
	}
L22:
	;
	v81 = int32(0)
	v86 = F_SearchSysCacheList(m, int32(9), int32(1), l2, v81, v81)
	mBase = m.M
	v87 = m.ExcPending
	if v87 != 0 {
		goto L1
	} else {
		goto L23
	}
L23:
	;
	v88 = *(*int32)(unsafe.Add(mBase, uint32(v86)+40))
	if v88 != 0 {
		goto L9
	} else {
		goto L24
	}
L24:
	;
	v156 = v81
	v165 = int32(0)
	goto L8
L25:
	;
	v96 = F_is_member_of_role_nosuper(m, l2, v91)
	mBase = m.M
	v97 = m.ExcPending
	if v97 != 0 {
		goto L1
	} else {
		goto L26
	}
L26:
	;
	if v96 == int32(0) {
		v41 = v41 + int32(1)
		goto L10
	} else {
		goto L27
	}
L27:
	;
	goto L11
L28:
	;
	F_errcode(m, int32(16910080))
	mBase = m.M
	v106 = m.ExcPending
	if v106 != 0 {
		goto L1
	} else {
		goto L29
	}
L29:
	;
	v107 = F_get_rolespec_name(m, v90)
	mBase = m.M
	v108 = m.ExcPending
	if v108 != 0 {
		goto L1
	} else {
		goto L30
	}
L30:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v21)+84)) = v107
	*(*int32)(unsafe.Add(mBase, uint32(v21)+80)) = l1
	F_errmsg(m, int32(717465), v21+int32(80))
	mBase = m.M
	v115 = m.ExcPending
	if v115 != 0 {
		goto L1
	} else {
		goto L31
	}
L31:
	;
	F_errfinish(m, int32(496702), int32(1748), int32(151399))
	mBase = m.M
	v120 = m.ExcPending
	if v120 != 0 {
		goto L1
	} else {
		goto L32
	}
L32:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L33:
	;
	v125 = *(*int32)(unsafe.Add(mBase, uint32(v86)+40))
	if v125 <= int32(0) {
		v156 = v125
		v165 = v123
		goto L8
	} else {
		goto L34
	}
L34:
	;
	v134 = int32(0)
	goto L35
L35:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v123+v134<<(uint(int32(2))%32)))) = int32(0)
	v153 = v134 + int32(1)
	v154 = *(*int32)(unsafe.Add(mBase, uint32(v86)+40))
	if v153 < v154 {
		v134 = v153
		goto L35
	} else {
		goto L37
	}
L36:
	;
	v156 = v154
	v165 = v123
	goto L8
L37:
	;
	goto L36
L38:
	;
	if int32(0) < v267 {
		goto L56
	} else {
		goto L57
	}
L39:
	;
	v176 = *(*int32)(unsafe.Add(mBase, uint32(l4)+4))
	if v176 <= int32(0) {
		v267 = v156
		goto L38
	} else {
		goto L40
	}
L40:
	;
	v181 = v156
	v195 = int32(0)
	goto L41
L41:
	;
	v199 = *(*int32)(unsafe.Add(mBase, uint32(l4)+12))
	v203 = *(*int32)(unsafe.Add(mBase, uint32(v199+v195<<(uint(int32(2))%32))))
	if v203 == int32(10) {
		goto L6
	} else {
		goto L43
	}
L42:
	;
	v267 = v245
	goto L38
L43:
	;
	v206 = int32(0)
	if v206 < v181 {
		goto L44
	} else {
		goto L45
	}
L44:
	;
	v209 = v181
	v214 = v206
	goto L47
L45:
	;
	v245 = v181
	goto L46
L46:
	;
	v264 = v195 + int32(1)
	v265 = *(*int32)(unsafe.Add(mBase, uint32(l4)+4))
	if v264 < v265 {
		v181 = v245
		v195 = v264
		goto L41
	} else {
		goto L54
	}
L47:
	;
	v230 = *(*int32)(unsafe.Add(mBase, uint32(v86+int32(48)+v214<<(uint(int32(2))%32))))
	v231 = *(*int32)(unsafe.Add(mBase, uint32(v230)+56))
	v232 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v231)+22)))
	v234 = *(*int32)(unsafe.Add(mBase, uint32(v231+v232)+8))
	if v203 == v234 {
		goto L49
	} else {
		goto L50
	}
L48:
	;
	v245 = v241
	goto L46
L49:
	;
	F_plan_recursive_revoke(m, v86, v165, v214, int32(0), int32(1))
	mBase = m.M
	v239 = m.ExcPending
	if v239 != 0 {
		goto L1
	} else {
		goto L52
	}
L50:
	;
	v241 = v209
	goto L51
L51:
	;
	v243 = v214 + int32(1)
	if v243 < v241 {
		v209 = v241
		v214 = v243
		goto L47
	} else {
		goto L53
	}
L52:
	;
	v240 = *(*int32)(unsafe.Add(mBase, uint32(v86)+40))
	v241 = v240
	goto L51
L53:
	;
	goto L48
L54:
	;
	goto L42
L55:
	;
	F_ReleaseCatCacheList(m, v86)
	mBase = m.M
	v362 = m.ExcPending
	if v362 != 0 {
		goto L1
	} else {
		goto L70
	}
L56:
	;
	v295 = int32(0)
	goto L59
L57:
	;
	goto L58
L58:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v345 = m.ExcPending
	if v345 != 0 {
		goto L1
	} else {
		goto L66
	}
L59:
	;
	v309 = v295 << (uint(int32(2)) % 32)
	v311 = *(*int32)(unsafe.Add(mBase, uint32(v165+v309)))
	if v311 != 0 {
		goto L61
	} else {
		goto L62
	}
L60:
	;
	goto L58
L61:
	;
	v322 = v295 + int32(1)
	if v322 != v267 {
		v295 = v322
		goto L59
	} else {
		goto L65
	}
L62:
	;
	v313 = *(*int32)(unsafe.Add(mBase, uint32(v309+(v86+int32(48)))))
	v314 = *(*int32)(unsafe.Add(mBase, uint32(v313)+56))
	v315 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v314)+22)))
	v316 = v314 + v315
	v317 = *(*int32)(unsafe.Add(mBase, uint32(v316)+8))
	if v317 != v24 {
		goto L61
	} else {
		goto L63
	}
L63:
	;
	v319 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v316)+16)))
	if v319 != 0 {
		goto L55
	} else {
		goto L64
	}
L64:
	;
	goto L61
L65:
	;
	goto L60
L66:
	;
	F_errcode(m, int32(16910080))
	mBase = m.M
	v348 = m.ExcPending
	if v348 != 0 {
		goto L1
	} else {
		goto L67
	}
L67:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v21)+32)) = int32(531854)
	F_errmsg(m, int32(208414), v21+int32(32))
	mBase = m.M
	v355 = m.ExcPending
	if v355 != 0 {
		goto L1
	} else {
		goto L68
	}
L68:
	;
	F_errfinish(m, int32(496702), int32(1814), int32(151399))
	mBase = m.M
	v360 = m.ExcPending
	if v360 != 0 {
		goto L1
	} else {
		goto L69
	}
L69:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L70:
	;
	goto L7
L71:
	;
	v404 = int32(0)
	if l3 == v404 {
		v414 = v404
		goto L73
	} else {
		goto L74
	}
L73:
	;
	if l4 == int32(0) {
		goto L79
	} else {
		goto L80
	}
L74:
	;
	v408 = *(*int32)(unsafe.Add(mBase, uint32(l3)+4))
	if v408 <= v391 {
		v414 = int32(0)
		goto L73
	} else {
		goto L75
	}
L75:
	;
	v410 = *(*int32)(unsafe.Add(mBase, uint32(l3)+12))
	v414 = v410 + v391<<(uint(int32(2))%32)
	goto L73
L76:
	;
	F_CommandCounterIncrement(m)
	mBase = m.M
	v609 = m.ExcPending
	if v609 != 0 {
		goto L1
	} else {
		goto L126
	}
L77:
	;
	v595 = F_heap_modify_tuple(m, v449, v30, v21+int32(112), v21+int32(104), v21+int32(96))
	mBase = m.M
	v596 = m.ExcPending
	if v596 != 0 {
		goto L1
	} else {
		goto L123
	}
L78:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v578 = m.ExcPending
	if v578 != 0 {
		goto L1
	} else {
		goto L120
	}
L79:
	;
	F_sequence_close(m, v28, int32(0))
	mBase = m.M
	v571 = m.ExcPending
	if v571 != 0 {
		goto L1
	} else {
		goto L119
	}
L80:
	;
	v417 = *(*int32)(unsafe.Add(mBase, uint32(l4)+4))
	if v417 <= v391 {
		goto L79
	} else {
		goto L81
	}
L81:
	;
	if v414 == int32(0) {
		goto L79
	} else {
		goto L82
	}
L82:
	;
	v421 = *(*int32)(unsafe.Add(mBase, uint32(l4)+12))
	v424 = v421 + v391<<(uint(int32(2))%32)
	if v424 == int32(0) {
		goto L79
	} else {
		goto L83
	}
L83:
	;
	v427 = *(*int32)(unsafe.Add(mBase, uint32(v414)))
	v428 = *(*int32)(unsafe.Add(mBase, uint32(v424)))
	v429 = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v21+int32(136)))) = v429
	v431 = int64(0)
	*(*int64)(unsafe.Add(mBase, uint32(v21+int32(128)))) = v431
	*(*int64)(unsafe.Add(mBase, uint32(v21)+120)) = v431
	*(*int64)(unsafe.Add(mBase, uint32(v21)+112)) = v431
	*(*int32)(unsafe.Add(mBase, uint32(v21)+107)) = v429
	*(*int32)(unsafe.Add(mBase, uint32(v21)+104)) = v429
	*(*int32)(unsafe.Add(mBase, uint32(v21)+124)) = v24
	*(*int32)(unsafe.Add(mBase, uint32(v21)+120)) = v428
	*(*int32)(unsafe.Add(mBase, uint32(v21)+116)) = l2
	*(*int32)(unsafe.Add(mBase, uint32(v21)+99)) = v429
	*(*int32)(unsafe.Add(mBase, uint32(v21)+96)) = v429
	v449 = F_SearchSysCache3(m, int32(9), l2, v428, v24)
	mBase = m.M
	v450 = m.ExcPending
	if v450 != 0 {
		goto L1
	} else {
		goto L84
	}
L84:
	;
	if v449 != 0 {
		goto L85
	} else {
		goto L86
	}
L85:
	;
	v451 = *(*int32)(unsafe.Add(mBase, uint32(v449)+16))
	v452 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v451)+22)))
	v453 = v451 + v452
	v454 = int32(0)
	v455 = *(*int32)(unsafe.Add(mBase, uint32(l6)))
	if v455&int32(1) == v454 {
		v467 = v454
		goto L88
	} else {
		goto L89
	}
L86:
	;
	goto L87
L87:
	;
	v521 = F_palloc(m, int32(4))
	mBase = m.M
	v522 = m.ExcPending
	if v522 != 0 {
		goto L1
	} else {
		goto L107
	}
L88:
	;
	if v455&int32(2) == int32(0) {
		v480 = v467
		goto L91
	} else {
		goto L92
	}
L89:
	;
	v460 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l6)+4)))
	v461 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v453)+16)))
	if v460 == v461 {
		v467 = v454
		goto L88
	} else {
		goto L90
	}
L90:
	;
	v463 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(v21)+100)) = uint8(v463)
	*(*int32)(unsafe.Add(mBase, uint32(v21)+128)) = v460
	v467 = v463
	goto L88
L91:
	;
	if v455&int32(4) == int32(0) {
		goto L94
	} else {
		goto L95
	}
L92:
	;
	v473 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l6)+5)))
	v474 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v453)+17)))
	if v473 == v474 {
		v480 = v467
		goto L91
	} else {
		goto L93
	}
L93:
	;
	v476 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(v21)+101)) = uint8(v476)
	*(*int32)(unsafe.Add(mBase, uint32(v21)+132)) = v473
	v480 = v476
	goto L91
L94:
	;
	if v480 != 0 {
		goto L77
	} else {
		goto L97
	}
L95:
	;
	v486 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l6)+6)))
	v487 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v453)+18)))
	if v486 == v487 {
		goto L94
	} else {
		goto L96
	}
L96:
	;
	v489 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(v21)+102)) = uint8(v489)
	*(*int32)(unsafe.Add(mBase, uint32(v21)+136)) = v486
	goto L77
L97:
	;
	v495 = F_errstart(m, int32(18), int32(0))
	mBase = m.M
	v496 = m.ExcPending
	if v496 != 0 {
		goto L1
	} else {
		goto L98
	}
L98:
	;
	if v495 != 0 {
		goto L99
	} else {
		goto L100
	}
L99:
	;
	v497 = F_get_rolespec_name(m, v427)
	mBase = m.M
	v498 = m.ExcPending
	if v498 != 0 {
		goto L1
	} else {
		goto L102
	}
L100:
	;
	goto L101
L101:
	;
	F_ReleaseCatCache(m, v449)
	mBase = m.M
	v517 = m.ExcPending
	if v517 != 0 {
		goto L1
	} else {
		goto L106
	}
L102:
	;
	v500 = F_GetUserNameFromId(m, v24, int32(0))
	mBase = m.M
	v501 = m.ExcPending
	if v501 != 0 {
		goto L1
	} else {
		goto L103
	}
L103:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v21)+24)) = v500
	*(*int32)(unsafe.Add(mBase, uint32(v21)+20)) = l1
	*(*int32)(unsafe.Add(mBase, uint32(v21)+16)) = v497
	F_errmsg(m, int32(717124), v21+int32(16))
	mBase = m.M
	v509 = m.ExcPending
	if v509 != 0 {
		goto L1
	} else {
		goto L104
	}
L104:
	;
	F_errfinish(m, int32(496702), int32(1892), int32(151399))
	mBase = m.M
	v514 = m.ExcPending
	if v514 != 0 {
		goto L1
	} else {
		goto L105
	}
L105:
	;
	goto L101
L106:
	;
	v391 = v391 + int32(1)
	goto L71
L107:
	;
	v523 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l6)+4)))
	*(*int32)(unsafe.Add(mBase, uint32(v21)+128)) = v523
	v525 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l6)+6)))
	*(*int32)(unsafe.Add(mBase, uint32(v21)+136)) = v525
	v527 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l6))))
	if v527&int32(2) != 0 {
		goto L109
	} else {
		goto L110
	}
L108:
	;
	v548 = F_GetNewOidWithIndex(m, v28, int32(6303), int32(1))
	mBase = m.M
	v549 = m.ExcPending
	if v549 != 0 {
		goto L1
	} else {
		goto L115
	}
L109:
	;
	v530 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l6)+5)))
	*(*int32)(unsafe.Add(mBase, uint32(v21)+132)) = v530
	goto L108
L110:
	;
	goto L111
L111:
	;
	v533 = F_SearchSysCache1(m, int32(11), v428)
	mBase = m.M
	v534 = m.ExcPending
	if v534 != 0 {
		goto L1
	} else {
		goto L112
	}
L112:
	;
	if v533 == int32(0) {
		goto L78
	} else {
		goto L113
	}
L113:
	;
	v537 = *(*int32)(unsafe.Add(mBase, uint32(v533)+16))
	v538 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v537)+22)))
	v540 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v537+v538)+69)))
	*(*int32)(unsafe.Add(mBase, uint32(v21)+132)) = v540
	F_ReleaseCatCache(m, v533)
	mBase = m.M
	v543 = m.ExcPending
	if v543 != 0 {
		goto L1
	} else {
		goto L114
	}
L114:
	;
	goto L108
L115:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v21)+112)) = v548
	v555 = F_heap_form_tuple(m, v30, v21+int32(112), v21+int32(104))
	mBase = m.M
	v556 = m.ExcPending
	if v556 != 0 {
		goto L1
	} else {
		goto L116
	}
L116:
	;
	F_CatalogTupleInsert(m, v28, v555)
	mBase = m.M
	v558 = m.ExcPending
	if v558 != 0 {
		goto L1
	} else {
		goto L117
	}
L117:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v521))) = v24
	v561 = int32(0)
	F_updateAclDependencies(m, int32(1261), v548, v561, v561, v561, v561, int32(1), v521)
	mBase = m.M
	v567 = m.ExcPending
	if v567 != 0 {
		goto L1
	} else {
		goto L118
	}
L118:
	;
	goto L76
L119:
	;
	m.G0 = v21 + int32(144)
	return
L120:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v21))) = v428
	F_errmsg_internal(m, int32(52148), v21)
	mBase = m.M
	v582 = m.ExcPending
	if v582 != 0 {
		goto L1
	} else {
		goto L121
	}
L121:
	;
	F_errfinish(m, int32(496702), int32(1934), int32(151399))
	mBase = m.M
	v587 = m.ExcPending
	if v587 != 0 {
		goto L1
	} else {
		goto L122
	}
L122:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L123:
	;
	F_CatalogTupleUpdate(m, v28, v595+int32(4), v595)
	mBase = m.M
	v600 = m.ExcPending
	if v600 != 0 {
		goto L1
	} else {
		goto L124
	}
L124:
	;
	F_ReleaseCatCache(m, v449)
	mBase = m.M
	v602 = m.ExcPending
	if v602 != 0 {
		goto L1
	} else {
		goto L125
	}
L125:
	;
	goto L76
L126:
	;
	v391 = v391 + int32(1)
	goto L71
L127:
	;
	F_errcode(m, int32(16910080))
	mBase = m.M
	v618 = m.ExcPending
	if v618 != 0 {
		goto L1
	} else {
		goto L128
	}
L128:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v21)+48)) = int32(531854)
	F_errmsg(m, int32(208414), v21+int32(48))
	mBase = m.M
	v625 = m.ExcPending
	if v625 != 0 {
		goto L1
	} else {
		goto L129
	}
L129:
	;
	F_errfinish(m, int32(496702), int32(1788), int32(151399))
	mBase = m.M
	v630 = m.ExcPending
	if v630 != 0 {
		goto L1
	} else {
		goto L130
	}
L130:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L131:
	;
	F_errcode(m, int32(1088))
	mBase = m.M
	v637 = m.ExcPending
	if v637 != 0 {
		goto L1
	} else {
		goto L132
	}
L132:
	;
	v638 = F_get_rolespec_name(m, v90)
	mBase = m.M
	v639 = m.ExcPending
	if v639 != 0 {
		goto L1
	} else {
		goto L133
	}
L133:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v21)+64)) = v638
	F_errmsg(m, int32(386615), v21-int32(-64))
	mBase = m.M
	v645 = m.ExcPending
	if v645 != 0 {
		goto L1
	} else {
		goto L134
	}
L134:
	;
	F_errfinish(m, int32(496702), int32(1735), int32(151399))
	mBase = m.M
	v650 = m.ExcPending
	if v650 != 0 {
		goto L1
	} else {
		goto L135
	}
L135:
	;
	base.Wasm_trap_unreachable()
	for {
	}
}
func F_RoleMembershipCacheCallback(m *base.Module, l0 int32, l1 int32, l2 int32) {
	mBase := m.M
	_ = mBase
	var v9 int32
	_ = v9
	if l1 != int32(21) {
		*(*int64)(unsafe.Add(mBase, _consts[1032])) = int64(0)
		*(*int32)(unsafe.Add(mBase, _consts[1035])) = int32(0)
	} else {
		if l2 == int32(0) {
			*(*int64)(unsafe.Add(mBase, _consts[1032])) = int64(0)
			*(*int32)(unsafe.Add(mBase, _consts[1035])) = int32(0)
		} else {
			v9 = *(*int32)(unsafe.Add(mBase, _consts[1036]))
			if l2 != v9 {
			} else {
				*(*int64)(unsafe.Add(mBase, _consts[1032])) = int64(0)
				*(*int32)(unsafe.Add(mBase, _consts[1035])) = int32(0)
			}
		}
	}
	return
}
func F__equalCreateRoleStmt(m *base.Module, l0 int32, l1 int32) int32 {
	mBase := m.M
	_ = mBase
	var v3 int32
	_ = v3
	var v6 int32
	_ = v6
	var v7 int32
	_ = v7
	var v9 int32
	_ = v9
	var v10 int32
	_ = v10
	var v15 int32
	_ = v15
	var v16 int32
	_ = v16
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
	var v35 int32
	_ = v35
	var v36 int32
	_ = v36
	var v41 int32
	_ = v41
	var v42 int32
	_ = v42
	var v43 int32
	_ = v43
	var v46 int32
	_ = v46
	var v48 int32
	_ = v48
	v3 = int32(0)
	v6 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v7 = *(*int32)(unsafe.Add(mBase, uint32(l1)+4))
	if v6 != v7 {
		v48 = v3
		goto L1
	} else {
		goto L2
	}
L1:
	;
	return v48
L2:
	;
	v9 = *(*int32)(unsafe.Add(mBase, uint32(l1)+8))
	v10 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	if v10 != 0 {
		goto L4
	} else {
		goto L5
	}
L3:
	;
	v41 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	v42 = *(*int32)(unsafe.Add(mBase, uint32(l1)+12))
	v43 = F_equal(m, v41, v42)
	mBase = m.M
	v46 = m.ExcPending
	if v46 != 0 {
		goto L18
	} else {
		goto L19
	}
L4:
	;
	if v9 == int32(0) {
		v48 = v3
		goto L1
	} else {
		goto L7
	}
L5:
	;
	goto L6
L6:
	;
	if v9 != v10 {
		v48 = v3
		goto L1
	} else {
		goto L17
	}
L7:
	;
	v15 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v9))))
	v16 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v10))))
	if v16 == int32(0) {
		v35 = v15
		v36 = v16
		goto L9
	} else {
		goto L10
	}
L8:
	;
	if v36-v35 == int32(0) {
		goto L3
	} else {
		goto L16
	}
L9:
	;
	goto L8
L10:
	;
	if v15 != v16 {
		v35 = v15
		v36 = v16
		goto L9
	} else {
		goto L11
	}
L11:
	;
	v20 = v10
	v21 = v9
	goto L12
L12:
	;
	v24 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v21)+1)))
	v25 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v20)+1)))
	if v25 == int32(0) {
		v35 = v24
		v36 = v25
		goto L9
	} else {
		goto L14
	}
L13:
	;
	v35 = v24
	v36 = v25
	goto L9
L14:
	;
	v28 = int32(1)
	if v24 == v25 {
		v20 = v20 + v28
		v21 = v21 + v28
		goto L12
	} else {
		goto L15
	}
L15:
	;
	goto L13
L16:
	;
	v48 = v3
	goto L1
L17:
	;
	goto L3
L18:
	;
	return int32(0)
L19:
	;
	v48 = v43
	goto L1
}
func F_check_can_set_role(m *base.Module, l0 int32, l1 int32) {
	mBase := m.M
	_ = mBase
	var v4 int32
	_ = v4
	var v6 int32
	_ = v6
	var v9 int32
	_ = v9
	var v10 int32
	_ = v10
	var v12 int32
	_ = v12
	var v14 int32
	_ = v14
	var v15 int32
	_ = v15
	var v16 int32
	_ = v16
	var v22 int32
	_ = v22
	var v25 int32
	_ = v25
	var v28 int32
	_ = v28
	var v29 int32
	_ = v29
	var v31 int32
	_ = v31
	var v39 int32
	_ = v39
	var v40 int32
	_ = v40
	var v42 int32
	_ = v42
	var v47 int32
	_ = v47
	var v54 int32
	_ = v54
	var v58 int32
	_ = v58
	var v61 int32
	_ = v61
	var v63 int32
	_ = v63
	var v64 int32
	_ = v64
	var v68 int32
	_ = v68
	var v73 int32
	_ = v73
	v4 = m.G0
	v6 = v4 - int32(16)
	m.G0 = v6
	if l0 == l1 {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	m.G0 = v6 + int32(16)
	return
L2:
	;
	v9 = F_superuser_arg(m, l0)
	mBase = m.M
	v10 = m.ExcPending
	if v10 != 0 {
		goto L3
	} else {
		goto L4
	}
L3:
	;
	return
L4:
	;
	if v9 != 0 {
		goto L1
	} else {
		goto L5
	}
L5:
	;
	v12 = int32(0)
	v14 = F_roles_is_member_of(m, l0, int32(2), v12, v12)
	mBase = m.M
	v15 = m.ExcPending
	if v15 != 0 {
		goto L3
	} else {
		goto L6
	}
L6:
	;
	v16 = int32(0)
	if v14 == v16 {
		goto L8
	} else {
		goto L9
	}
L7:
	;
	if v54 != 0 {
		goto L1
	} else {
		goto L20
	}
L8:
	;
	v54 = int32(0)
	goto L7
L9:
	;
	goto L10
L10:
	;
	v22 = *(*int32)(unsafe.Add(mBase, uint32(v14)+4))
	if v22 <= int32(0) {
		v47 = v16
		goto L11
	} else {
		goto L12
	}
L11:
	;
	v54 = v47
	goto L7
L12:
	;
	v25 = int32(0)
	if v25 < v22 {
		goto L13
	} else {
		goto L14
	}
L13:
	;
	v28 = v22
	goto L15
L14:
	;
	v28 = v25
	goto L15
L15:
	;
	v29 = *(*int32)(unsafe.Add(mBase, uint32(v14)+12))
	v31 = int32(0)
	goto L16
L16:
	;
	v39 = *(*int32)(unsafe.Add(mBase, uint32(v29+v31<<(uint(int32(2))%32))))
	v40 = base.B2i32(v39 == l1)
	if v39 == l1 {
		v47 = v40
		goto L11
	} else {
		goto L18
	}
L17:
	;
	v47 = v40
	goto L11
L18:
	;
	v42 = v31 + int32(1)
	if v42 != v28 {
		v31 = v42
		goto L16
	} else {
		goto L19
	}
L19:
	;
	goto L17
L20:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v58 = m.ExcPending
	if v58 != 0 {
		goto L3
	} else {
		goto L21
	}
L21:
	;
	F_errcode(m, int32(16797828))
	mBase = m.M
	v61 = m.ExcPending
	if v61 != 0 {
		goto L3
	} else {
		goto L22
	}
L22:
	;
	v63 = F_GetUserNameFromId(m, l1, int32(0))
	mBase = m.M
	v64 = m.ExcPending
	if v64 != 0 {
		goto L3
	} else {
		goto L23
	}
L23:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v6))) = v63
	F_errmsg(m, int32(726516), v6)
	mBase = m.M
	v68 = m.ExcPending
	if v68 != 0 {
		goto L3
	} else {
		goto L24
	}
L24:
	;
	F_errfinish(m, int32(498986), int32(5348), int32(386560))
	mBase = m.M
	v73 = m.ExcPending
	if v73 != 0 {
		goto L3
	} else {
		goto L25
	}
L25:
	;
	base.Wasm_trap_unreachable()
	for {
	}
}
