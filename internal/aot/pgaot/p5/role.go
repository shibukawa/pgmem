package p5

import (
	base "github.com/shibukawa/pgmem/internal/aot/pgaot/base"
	"unsafe"
)

func F_AddRoleMems(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32, l5 int32, l6 int32) {
	mBase := m.M
	_ = mBase
	var v20 int32
	_ = v20
	var v22 int32
	_ = v22
	var v25 int32
	_ = v25
	var v26 int32
	_ = v26
	var v29 int32
	_ = v29
	var v30 int32
	_ = v30
	var v31 int32
	_ = v31
	var v36 int32
	_ = v36
	var v42 int32
	_ = v42
	var v56 int32
	_ = v56
	var v60 int32
	_ = v60
	var v62 int32
	_ = v62
	var v66 int32
	_ = v66
	var v71 int32
	_ = v71
	var v74 int32
	_ = v74
	var v76 int64
	_ = v76
	var v79 int32
	_ = v79
	var v85 int64
	_ = v85
	var v87 int32
	_ = v87
	var v88 int32
	_ = v88
	var v89 int32
	_ = v89
	var v90 int32
	_ = v90
	var v92 int32
	_ = v92
	var v96 int32
	_ = v96
	var v101 int32
	_ = v101
	var v102 int32
	_ = v102
	var v108 int32
	_ = v108
	var v111 int32
	_ = v111
	var v112 int32
	_ = v112
	var v113 int32
	_ = v113
	var v120 int32
	_ = v120
	var v125 int32
	_ = v125
	var v127 int32
	_ = v127
	var v128 int32
	_ = v128
	var v129 int32
	_ = v129
	var v138 int32
	_ = v138
	var v158 int32
	_ = v158
	var v159 int32
	_ = v159
	var v161 int32
	_ = v161
	var v163 int32
	_ = v163
	var v182 int32
	_ = v182
	var v187 int32
	_ = v187
	var v197 int32
	_ = v197
	var v206 int32
	_ = v206
	var v210 int32
	_ = v210
	var v216 int32
	_ = v216
	var v219 int32
	_ = v219
	var v226 int32
	_ = v226
	var v231 int32
	_ = v231
	var v232 int32
	_ = v232
	var v235 int32
	_ = v235
	var v240 int32
	_ = v240
	var v257 int32
	_ = v257
	var v258 int32
	_ = v258
	var v259 int32
	_ = v259
	var v261 int32
	_ = v261
	var v266 int32
	_ = v266
	var v267 int32
	_ = v267
	var v268 int32
	_ = v268
	var v270 int32
	_ = v270
	var v272 int32
	_ = v272
	var v292 int32
	_ = v292
	var v293 int32
	_ = v293
	var v298 int32
	_ = v298
	var v301 int32
	_ = v301
	var v302 int32
	_ = v302
	var v303 int32
	_ = v303
	var v309 int32
	_ = v309
	var v314 int32
	_ = v314
	var v315 int32
	_ = v315
	var v344 int32
	_ = v344
	var v359 int32
	_ = v359
	var v361 int32
	_ = v361
	var v363 int32
	_ = v363
	var v364 int32
	_ = v364
	var v365 int32
	_ = v365
	var v366 int32
	_ = v366
	var v367 int32
	_ = v367
	var v369 int32
	_ = v369
	var v372 int32
	_ = v372
	var v396 int32
	_ = v396
	var v399 int32
	_ = v399
	var v406 int32
	_ = v406
	var v411 int32
	_ = v411
	var v413 int32
	_ = v413
	var v433 int64
	_ = v433
	var v440 int32
	_ = v440
	var v454 int32
	_ = v454
	var v458 int32
	_ = v458
	var v460 int32
	_ = v460
	var v464 int32
	_ = v464
	var v469 int32
	_ = v469
	var v472 int32
	_ = v472
	var v478 int32
	_ = v478
	var v479 int32
	_ = v479
	var v480 int64
	_ = v480
	var v494 int32
	_ = v494
	var v499 int64
	_ = v499
	var v507 int32
	_ = v507
	var v508 int32
	_ = v508
	var v509 int32
	_ = v509
	var v510 int32
	_ = v510
	var v511 int32
	_ = v511
	var v512 int32
	_ = v512
	var v513 int32
	_ = v513
	var v518 int32
	_ = v518
	var v519 int32
	_ = v519
	var v521 int32
	_ = v521
	var v526 int32
	_ = v526
	var v532 int32
	_ = v532
	var v533 int32
	_ = v533
	var v535 int32
	_ = v535
	var v540 int32
	_ = v540
	var v546 int32
	_ = v546
	var v547 int32
	_ = v547
	var v549 int32
	_ = v549
	var v556 int32
	_ = v556
	var v557 int32
	_ = v557
	var v558 int32
	_ = v558
	var v559 int32
	_ = v559
	var v561 int32
	_ = v561
	var v562 int32
	_ = v562
	var v570 int32
	_ = v570
	var v575 int32
	_ = v575
	var v578 int32
	_ = v578
	var v582 int32
	_ = v582
	var v583 int32
	_ = v583
	var v584 int64
	_ = v584
	var v586 int64
	_ = v586
	var v588 int32
	_ = v588
	var v591 int64
	_ = v591
	var v594 int32
	_ = v594
	var v595 int32
	_ = v595
	var v598 int32
	_ = v598
	var v599 int32
	_ = v599
	var v601 int64
	_ = v601
	var v604 int32
	_ = v604
	var v609 int32
	_ = v609
	var v610 int32
	_ = v610
	var v617 int32
	_ = v617
	var v618 int32
	_ = v618
	var v620 int32
	_ = v620
	var v623 int32
	_ = v623
	var v629 int32
	_ = v629
	var v633 int32
	_ = v633
	var v640 int32
	_ = v640
	var v644 int32
	_ = v644
	var v649 int32
	_ = v649
	var v657 int32
	_ = v657
	var v658 int32
	_ = v658
	var v662 int32
	_ = v662
	var v664 int32
	_ = v664
	var v671 int32
	_ = v671
	v20 = m.G0
	v22 = v20 - int32(176)
	m.G0 = v22
	v25 = F_check_role_grantor(m, l0, l2, l5, int32(1))
	mBase = m.M
	v26 = m.ExcPending
	if v26 != 0 {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	return
L2:
	;
	v29 = F_table_open(m, int32(1261), int32(3))
	mBase = m.M
	v30 = m.ExcPending
	if v30 != 0 {
		goto L1
	} else {
		goto L3
	}
L3:
	;
	v31 = *(*int32)(unsafe.Add(mBase, uint32(v29)+52))
	F_LockSharedObject(m, int32(1260), l2, int32(4))
	mBase = m.M
	v36 = m.ExcPending
	if v36 != 0 {
		goto L1
	} else {
		goto L4
	}
L4:
	;
	v42 = int32(0)
	goto L10
L5:
	;
	v433 = base.I64_extend_i32_u(v25)
	v440 = int32(0)
	goto L79
L6:
	;
	if int32(0) < v315 {
		goto L64
	} else {
		goto L65
	}
L7:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v298 = m.ExcPending
	if v298 != 0 {
		goto L1
	} else {
		goto L58
	}
L8:
	;
	if l4 == int32(0) {
		v315 = v161
		goto L6
	} else {
		goto L36
	}
L9:
	;
	v127 = F_palloc_mul(m, int32(4), v89)
	mBase = m.M
	v128 = m.ExcPending
	if v128 != 0 {
		goto L1
	} else {
		goto L31
	}
L10:
	;
	v56 = int32(0)
	if l3 == v56 {
		v66 = v56
		goto L12
	} else {
		goto L13
	}
L11:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v108 = m.ExcPending
	if v108 != 0 {
		goto L1
	} else {
		goto L26
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
	v60 = *(*int32)(unsafe.Add(mBase, uint32(l3)+4))
	if v60 <= v42 {
		v66 = int32(0)
		goto L12
	} else {
		goto L14
	}
L14:
	;
	v62 = *(*int32)(unsafe.Add(mBase, uint32(l3)+12))
	v66 = v62 + v42<<(uint(int32(2))%32)
	goto L12
L15:
	;
	v92 = *(*int32)(unsafe.Add(mBase, uint32(v66)))
	v96 = *(*int32)(unsafe.Add(mBase, uint32(v74+v42<<(uint(int32(2))%32))))
	if v96 == int32(_a_F_AddRoleMems_0) {
		goto L7
	} else {
		goto L23
	}
L16:
	;
	v76 = base.I64_extend_i32_u(l2)
	v79 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l6)+4)))
	if base.B2i32(v25 == int32(10))|base.B2i32(v79 != int32(1)) != 0 {
		goto L5
	} else {
		goto L20
	}
L17:
	;
	v71 = *(*int32)(unsafe.Add(mBase, uint32(l4)+4))
	if base.B2i32(v66 == int32(0))|base.B2i32(v71 <= v42) != 0 {
		goto L16
	} else {
		goto L18
	}
L18:
	;
	v74 = *(*int32)(unsafe.Add(mBase, uint32(l4)+12))
	if v74 != 0 {
		goto L15
	} else {
		goto L19
	}
L19:
	;
	goto L16
L20:
	;
	v85 = int64(0)
	v87 = F_SearchSysCacheList(m, int32(9), int32(1), v76, v85, v85)
	mBase = m.M
	v88 = m.ExcPending
	if v88 != 0 {
		goto L1
	} else {
		goto L21
	}
L21:
	;
	v89 = *(*int32)(unsafe.Add(mBase, uint32(v87)+56))
	if v89 != 0 {
		goto L9
	} else {
		goto L22
	}
L22:
	;
	v90 = int32(0)
	v161 = v90
	v163 = v90
	goto L8
L23:
	;
	v101 = F_is_member_of_role_nosuper(m, l2, v96)
	mBase = m.M
	v102 = m.ExcPending
	if v102 != 0 {
		goto L1
	} else {
		goto L24
	}
L24:
	;
	if v101 == int32(0) {
		v42 = v42 + int32(1)
		goto L10
	} else {
		goto L25
	}
L25:
	;
	goto L11
L26:
	;
	F_errcode(m, int32(16910080))
	mBase = m.M
	v111 = m.ExcPending
	if v111 != 0 {
		goto L1
	} else {
		goto L27
	}
L27:
	;
	v112 = F_get_rolespec_name(m, v92)
	mBase = m.M
	v113 = m.ExcPending
	if v113 != 0 {
		goto L1
	} else {
		goto L28
	}
L28:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v22)+84)) = v112
	*(*int32)(unsafe.Add(mBase, uint32(v22)+80)) = l1
	F_errmsg(m, int32(_a_F_AddRoleMems_1), v22+int32(80))
	mBase = m.M
	v120 = m.ExcPending
	if v120 != 0 {
		goto L1
	} else {
		goto L29
	}
L29:
	;
	F_errfinish(m, int32(_a_F_AddRoleMems_2), int32(1761), int32(_a_F_AddRoleMems_3))
	mBase = m.M
	v125 = m.ExcPending
	if v125 != 0 {
		goto L1
	} else {
		goto L30
	}
L30:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L31:
	;
	v129 = *(*int32)(unsafe.Add(mBase, uint32(v87)+56))
	if v129 <= int32(0) {
		v161 = v129
		v163 = v127
		goto L8
	} else {
		goto L32
	}
L32:
	;
	v138 = int32(0)
	goto L33
L33:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v127+v138<<(uint(int32(2))%32)))) = int32(0)
	v158 = v138 + int32(1)
	v159 = *(*int32)(unsafe.Add(mBase, uint32(v87)+56))
	if v158 < v159 {
		v138 = v158
		goto L33
	} else {
		goto L35
	}
L34:
	;
	v161 = v159
	v163 = v127
	goto L8
L35:
	;
	goto L34
L36:
	;
	v182 = *(*int32)(unsafe.Add(mBase, uint32(l4)+4))
	if v182 <= int32(0) {
		v315 = v161
		goto L6
	} else {
		goto L37
	}
L37:
	;
	v187 = v161
	v197 = int32(0)
	goto L38
L38:
	;
	v206 = *(*int32)(unsafe.Add(mBase, uint32(l4)+12))
	v210 = *(*int32)(unsafe.Add(mBase, uint32(v206+v197<<(uint(int32(2))%32))))
	if v210 == int32(10) {
		goto L40
	} else {
		goto L41
	}
L39:
	;
	v315 = v272
	goto L6
L40:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v216 = m.ExcPending
	if v216 != 0 {
		goto L1
	} else {
		goto L43
	}
L41:
	;
	goto L42
L42:
	;
	v232 = int32(0)
	if v232 < v187 {
		goto L47
	} else {
		goto L48
	}
L43:
	;
	F_errcode(m, int32(16910080))
	mBase = m.M
	v219 = m.ExcPending
	if v219 != 0 {
		goto L1
	} else {
		goto L44
	}
L44:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v22)+48)) = int32(_a_F_AddRoleMems_4)
	F_errmsg(m, int32(_a_F_AddRoleMems_5), v22+int32(48))
	mBase = m.M
	v226 = m.ExcPending
	if v226 != 0 {
		goto L1
	} else {
		goto L45
	}
L45:
	;
	F_errfinish(m, int32(_a_F_AddRoleMems_2), int32(1801), int32(_a_F_AddRoleMems_3))
	mBase = m.M
	v231 = m.ExcPending
	if v231 != 0 {
		goto L1
	} else {
		goto L46
	}
L46:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L47:
	;
	v235 = v187
	v240 = v232
	goto L50
L48:
	;
	v272 = v187
	goto L49
L49:
	;
	v292 = v197 + int32(1)
	v293 = *(*int32)(unsafe.Add(mBase, uint32(l4)+4))
	if v292 < v293 {
		v187 = v272
		v197 = v292
		goto L38
	} else {
		goto L57
	}
L50:
	;
	v257 = *(*int32)(unsafe.Add(mBase, uint32(v87-int32(-64)+v240<<(uint(int32(2))%32))))
	v258 = *(*int32)(unsafe.Add(mBase, uint32(v257)+72))
	v259 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v258)+22)))
	v261 = *(*int32)(unsafe.Add(mBase, uint32(v258+v259)+8))
	if v210 == v261 {
		goto L52
	} else {
		goto L53
	}
L51:
	;
	v272 = v268
	goto L49
L52:
	;
	F_plan_recursive_revoke(m, v87, v163, v240, int32(0), int32(1))
	mBase = m.M
	v266 = m.ExcPending
	if v266 != 0 {
		goto L1
	} else {
		goto L55
	}
L53:
	;
	v268 = v235
	goto L54
L54:
	;
	v270 = v240 + int32(1)
	if v270 < v268 {
		v235 = v268
		v240 = v270
		goto L50
	} else {
		goto L56
	}
L55:
	;
	v267 = *(*int32)(unsafe.Add(mBase, uint32(v87)+56))
	v268 = v267
	goto L54
L56:
	;
	goto L51
L57:
	;
	goto L39
L58:
	;
	F_errcode(m, int32(1088))
	mBase = m.M
	v301 = m.ExcPending
	if v301 != 0 {
		goto L1
	} else {
		goto L59
	}
L59:
	;
	v302 = F_get_rolespec_name(m, v92)
	mBase = m.M
	v303 = m.ExcPending
	if v303 != 0 {
		goto L1
	} else {
		goto L60
	}
L60:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v22)+64)) = v302
	F_errmsg(m, int32(_a_F_AddRoleMems_6), v22-int32(-64))
	mBase = m.M
	v309 = m.ExcPending
	if v309 != 0 {
		goto L1
	} else {
		goto L61
	}
L61:
	;
	F_errfinish(m, int32(_a_F_AddRoleMems_2), int32(1748), int32(_a_F_AddRoleMems_3))
	mBase = m.M
	v314 = m.ExcPending
	if v314 != 0 {
		goto L1
	} else {
		goto L62
	}
L62:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L63:
	;
	F_ReleaseCatCacheList(m, v87)
	mBase = m.M
	v413 = m.ExcPending
	if v413 != 0 {
		goto L1
	} else {
		goto L78
	}
L64:
	;
	v344 = int32(0)
	goto L67
L65:
	;
	goto L66
L66:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v396 = m.ExcPending
	if v396 != 0 {
		goto L1
	} else {
		goto L74
	}
L67:
	;
	v359 = v344 << (uint(int32(2)) % 32)
	v361 = *(*int32)(unsafe.Add(mBase, uint32(v163+v359)))
	if v361 != 0 {
		goto L69
	} else {
		goto L70
	}
L68:
	;
	goto L66
L69:
	;
	v372 = v344 + int32(1)
	if v372 != v315 {
		v344 = v372
		goto L67
	} else {
		goto L73
	}
L70:
	;
	v363 = *(*int32)(unsafe.Add(mBase, uint32(v359+(v87-int32(-64)))))
	v364 = *(*int32)(unsafe.Add(mBase, uint32(v363)+72))
	v365 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v364)+22)))
	v366 = v364 + v365
	v367 = *(*int32)(unsafe.Add(mBase, uint32(v366)+8))
	if v367 != v25 {
		goto L69
	} else {
		goto L71
	}
L71:
	;
	v369 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v366)+16)))
	if v369 != 0 {
		goto L63
	} else {
		goto L72
	}
L72:
	;
	goto L69
L73:
	;
	goto L68
L74:
	;
	F_errcode(m, int32(16910080))
	mBase = m.M
	v399 = m.ExcPending
	if v399 != 0 {
		goto L1
	} else {
		goto L75
	}
L75:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v22)+32)) = int32(_a_F_AddRoleMems_4)
	F_errmsg(m, int32(_a_F_AddRoleMems_5), v22+int32(32))
	mBase = m.M
	v406 = m.ExcPending
	if v406 != 0 {
		goto L1
	} else {
		goto L76
	}
L76:
	;
	F_errfinish(m, int32(_a_F_AddRoleMems_2), int32(1827), int32(_a_F_AddRoleMems_3))
	mBase = m.M
	v411 = m.ExcPending
	if v411 != 0 {
		goto L1
	} else {
		goto L77
	}
L77:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L78:
	;
	goto L5
L79:
	;
	v454 = int32(0)
	if l3 == v454 {
		v464 = v454
		goto L81
	} else {
		goto L82
	}
L81:
	;
	if l4 == int32(0) {
		goto L87
	} else {
		goto L88
	}
L82:
	;
	v458 = *(*int32)(unsafe.Add(mBase, uint32(l3)+4))
	if v458 <= v440 {
		v464 = int32(0)
		goto L81
	} else {
		goto L83
	}
L83:
	;
	v460 = *(*int32)(unsafe.Add(mBase, uint32(l3)+12))
	v464 = v460 + v440<<(uint(int32(2))%32)
	goto L81
L84:
	;
	F_CommandCounterIncrement(m)
	mBase = m.M
	v671 = m.ExcPending
	if v671 != 0 {
		goto L1
	} else {
		goto L133
	}
L85:
	;
	v657 = F_heap_modify_tuple(m, v507, v31, v22+int32(112), v22+int32(104), v22+int32(96))
	mBase = m.M
	v658 = m.ExcPending
	if v658 != 0 {
		goto L1
	} else {
		goto L130
	}
L86:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v640 = m.ExcPending
	if v640 != 0 {
		goto L1
	} else {
		goto L127
	}
L87:
	;
	F_relation_close(m, v29, int32(0))
	mBase = m.M
	v633 = m.ExcPending
	if v633 != 0 {
		goto L1
	} else {
		goto L126
	}
L88:
	;
	v469 = *(*int32)(unsafe.Add(mBase, uint32(l4)+4))
	if base.B2i32(v464 == int32(0))|base.B2i32(v469 <= v440) != 0 {
		goto L87
	} else {
		goto L89
	}
L89:
	;
	v472 = *(*int32)(unsafe.Add(mBase, uint32(l4)+12))
	if v472 == int32(0) {
		goto L87
	} else {
		goto L90
	}
L90:
	;
	v478 = *(*int32)(unsafe.Add(mBase, uint32(v472+v440<<(uint(int32(2))%32))))
	v479 = *(*int32)(unsafe.Add(mBase, uint32(v464)))
	v480 = int64(0)
	*(*int64)(unsafe.Add(mBase, uint32(v22)+160)) = v480
	*(*int64)(unsafe.Add(mBase, uint32(v22)+152)) = v480
	*(*int64)(unsafe.Add(mBase, uint32(v22)+144)) = v480
	*(*int64)(unsafe.Add(mBase, uint32(v22)+136)) = v480
	*(*int64)(unsafe.Add(mBase, uint32(v22)+128)) = v480
	*(*int64)(unsafe.Add(mBase, uint32(v22)+120)) = v480
	*(*int64)(unsafe.Add(mBase, uint32(v22)+112)) = v480
	v494 = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v22)+107)) = v494
	*(*int32)(unsafe.Add(mBase, uint32(v22)+104)) = v494
	*(*int64)(unsafe.Add(mBase, uint32(v22)+136)) = v433
	v499 = base.I64_extend_i32_u(v478)
	*(*int64)(unsafe.Add(mBase, uint32(v22)+128)) = v499
	*(*int64)(unsafe.Add(mBase, uint32(v22)+120)) = v76
	*(*int32)(unsafe.Add(mBase, uint32(v22)+99)) = v494
	*(*int32)(unsafe.Add(mBase, uint32(v22)+96)) = v494
	v507 = F_SearchSysCache3(m, int32(9), v76, v499, v433)
	mBase = m.M
	v508 = m.ExcPending
	if v508 != 0 {
		goto L1
	} else {
		goto L91
	}
L91:
	;
	if v507 != 0 {
		goto L92
	} else {
		goto L93
	}
L92:
	;
	v509 = *(*int32)(unsafe.Add(mBase, uint32(v507)+16))
	v510 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v509)+22)))
	v511 = v509 + v510
	v512 = int32(0)
	v513 = *(*int32)(unsafe.Add(mBase, uint32(l6)))
	if v513&int32(1) == v512 {
		v526 = v512
		goto L95
	} else {
		goto L96
	}
L93:
	;
	goto L94
L94:
	;
	v582 = F_palloc(m, int32(4))
	mBase = m.M
	v583 = m.ExcPending
	if v583 != 0 {
		goto L1
	} else {
		goto L114
	}
L95:
	;
	if v513&int32(2) == int32(0) {
		v540 = v526
		goto L98
	} else {
		goto L99
	}
L96:
	;
	v518 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l6)+4)))
	v519 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v511)+16)))
	if v518 == v519 {
		v526 = v512
		goto L95
	} else {
		goto L97
	}
L97:
	;
	v521 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(v22)+100)) = uint8(v521)
	*(*int64)(unsafe.Add(mBase, uint32(v22)+144)) = base.I64_extend_i32_u(v518)
	v526 = v521
	goto L95
L98:
	;
	if v513&int32(4) == int32(0) {
		goto L101
	} else {
		goto L102
	}
L99:
	;
	v532 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l6)+5)))
	v533 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v511)+17)))
	if v532 == v533 {
		v540 = v526
		goto L98
	} else {
		goto L100
	}
L100:
	;
	v535 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(v22)+101)) = uint8(v535)
	*(*int64)(unsafe.Add(mBase, uint32(v22)+152)) = base.I64_extend_i32_u(v532)
	v540 = v535
	goto L98
L101:
	;
	if v540 != 0 {
		goto L85
	} else {
		goto L104
	}
L102:
	;
	v546 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l6)+6)))
	v547 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v511)+18)))
	if v546 == v547 {
		goto L101
	} else {
		goto L103
	}
L103:
	;
	v549 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(v22)+102)) = uint8(v549)
	*(*int64)(unsafe.Add(mBase, uint32(v22)+160)) = base.I64_extend_i32_u(v546)
	goto L85
L104:
	;
	v556 = F_errstart(m, int32(18), int32(0))
	mBase = m.M
	v557 = m.ExcPending
	if v557 != 0 {
		goto L1
	} else {
		goto L105
	}
L105:
	;
	if v556 != 0 {
		goto L106
	} else {
		goto L107
	}
L106:
	;
	v558 = F_get_rolespec_name(m, v479)
	mBase = m.M
	v559 = m.ExcPending
	if v559 != 0 {
		goto L1
	} else {
		goto L109
	}
L107:
	;
	goto L108
L108:
	;
	F_ReleaseCatCache(m, v507)
	mBase = m.M
	v578 = m.ExcPending
	if v578 != 0 {
		goto L1
	} else {
		goto L113
	}
L109:
	;
	v561 = F_GetUserNameFromId(m, v25, int32(0))
	mBase = m.M
	v562 = m.ExcPending
	if v562 != 0 {
		goto L1
	} else {
		goto L110
	}
L110:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v22)+24)) = v561
	*(*int32)(unsafe.Add(mBase, uint32(v22)+20)) = l1
	*(*int32)(unsafe.Add(mBase, uint32(v22)+16)) = v558
	F_errmsg(m, int32(_a_F_AddRoleMems_7), v22+int32(16))
	mBase = m.M
	v570 = m.ExcPending
	if v570 != 0 {
		goto L1
	} else {
		goto L111
	}
L111:
	;
	F_errfinish(m, int32(_a_F_AddRoleMems_2), int32(1905), int32(_a_F_AddRoleMems_3))
	mBase = m.M
	v575 = m.ExcPending
	if v575 != 0 {
		goto L1
	} else {
		goto L112
	}
L112:
	;
	goto L108
L113:
	;
	v440 = v440 + int32(1)
	goto L79
L114:
	;
	v584 = int64(*(*uint8)(unsafe.Add(mBase, uint32(l6)+4)))
	*(*int64)(unsafe.Add(mBase, uint32(v22)+144)) = v584
	v586 = int64(*(*uint8)(unsafe.Add(mBase, uint32(l6)+6)))
	*(*int64)(unsafe.Add(mBase, uint32(v22)+160)) = v586
	v588 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l6))))
	if v588&int32(2) != 0 {
		goto L116
	} else {
		goto L117
	}
L115:
	;
	v609 = F_GetNewOidWithIndex(m, v29, int32(_a_F_AddRoleMems_8), int32(1))
	mBase = m.M
	v610 = m.ExcPending
	if v610 != 0 {
		goto L1
	} else {
		goto L122
	}
L116:
	;
	v591 = int64(*(*uint8)(unsafe.Add(mBase, uint32(l6)+5)))
	*(*int64)(unsafe.Add(mBase, uint32(v22)+152)) = v591
	goto L115
L117:
	;
	goto L118
L118:
	;
	v594 = F_SearchSysCache1(m, int32(11), v499)
	mBase = m.M
	v595 = m.ExcPending
	if v595 != 0 {
		goto L1
	} else {
		goto L119
	}
L119:
	;
	if v594 == int32(0) {
		goto L86
	} else {
		goto L120
	}
L120:
	;
	v598 = *(*int32)(unsafe.Add(mBase, uint32(v594)+16))
	v599 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v598)+22)))
	v601 = int64(*(*uint8)(unsafe.Add(mBase, uint32(v598+v599)+69)))
	*(*int64)(unsafe.Add(mBase, uint32(v22)+152)) = v601
	F_ReleaseCatCache(m, v594)
	mBase = m.M
	v604 = m.ExcPending
	if v604 != 0 {
		goto L1
	} else {
		goto L121
	}
L121:
	;
	goto L115
L122:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v22)+112)) = base.I64_extend_i32_u(v609)
	v617 = F_heap_form_tuple(m, v31, v22+int32(112), v22+int32(104))
	mBase = m.M
	v618 = m.ExcPending
	if v618 != 0 {
		goto L1
	} else {
		goto L123
	}
L123:
	;
	F_CatalogTupleInsert(m, v29, v617)
	mBase = m.M
	v620 = m.ExcPending
	if v620 != 0 {
		goto L1
	} else {
		goto L124
	}
L124:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v582))) = v25
	v623 = int32(0)
	F_updateAclDependencies(m, int32(1261), v609, v623, v623, v623, v623, int32(1), v582)
	mBase = m.M
	v629 = m.ExcPending
	if v629 != 0 {
		goto L1
	} else {
		goto L125
	}
L125:
	;
	goto L84
L126:
	;
	m.G0 = v22 + int32(176)
	return
L127:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v22))) = v478
	F_errmsg_internal(m, int32(_a_F_AddRoleMems_9), v22)
	mBase = m.M
	v644 = m.ExcPending
	if v644 != 0 {
		goto L1
	} else {
		goto L128
	}
L128:
	;
	F_errfinish(m, int32(_a_F_AddRoleMems_2), int32(1947), int32(_a_F_AddRoleMems_3))
	mBase = m.M
	v649 = m.ExcPending
	if v649 != 0 {
		goto L1
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
	F_CatalogTupleUpdate(m, v29, v657+int32(4), v657)
	mBase = m.M
	v662 = m.ExcPending
	if v662 != 0 {
		goto L1
	} else {
		goto L131
	}
L131:
	;
	F_ReleaseCatCache(m, v507)
	mBase = m.M
	v664 = m.ExcPending
	if v664 != 0 {
		goto L1
	} else {
		goto L132
	}
L132:
	;
	goto L84
L133:
	;
	v440 = v440 + int32(1)
	goto L79
}
func F_RoleMembershipCacheCallback(m *base.Module, l0 int64, l1 int32, l2 int32) {
	mBase := m.M
	_ = mBase
	var v4 int32
	_ = v4
	var v12 int32
	_ = v12
	v4 = int32(0)
	if base.B2i32(l2 == v4)|base.B2i32(l1 != int32(21)) == v4 {
		v12 = *(*int32)(unsafe.Add(mBase, _c_F_RoleMembershipCacheCallback[0]))
		if l2 != v12 {
		} else {
			*(*int64)(unsafe.Add(mBase, _c_F_RoleMembershipCacheCallback[1])) = int64(0)
			*(*int32)(unsafe.Add(mBase, _c_F_RoleMembershipCacheCallback[2])) = int32(0)
		}
	} else {
		*(*int64)(unsafe.Add(mBase, _c_F_RoleMembershipCacheCallback[1])) = int64(0)
		*(*int32)(unsafe.Add(mBase, _c_F_RoleMembershipCacheCallback[2])) = int32(0)
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
	var v29 int32
	_ = v29
	var v36 int32
	_ = v36
	var v37 int32
	_ = v37
	var v42 int32
	_ = v42
	var v43 int32
	_ = v43
	var v44 int32
	_ = v44
	var v47 int32
	_ = v47
	var v50 int32
	_ = v50
	v3 = int32(0)
	v6 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v7 = *(*int32)(unsafe.Add(mBase, uint32(l1)+4))
	if v6 != v7 {
		v50 = v3
		goto L1
	} else {
		goto L2
	}
L1:
	;
	return v50
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
	v42 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	v43 = *(*int32)(unsafe.Add(mBase, uint32(l1)+12))
	v44 = F_equal(m, v42, v43)
	mBase = m.M
	v47 = m.ExcPending
	if v47 != 0 {
		goto L17
	} else {
		goto L18
	}
L4:
	;
	if v9 == int32(0) {
		v50 = v3
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
		v50 = v3
		goto L1
	} else {
		goto L16
	}
L7:
	;
	v15 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v10))))
	v18 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v9))))
	if base.B2i32(v15 == int32(0))|base.B2i32(v15 != v18) != 0 {
		v36 = v15
		v37 = v18
		goto L9
	} else {
		goto L10
	}
L8:
	;
	if v36-v37 == int32(0) {
		goto L3
	} else {
		goto L15
	}
L9:
	;
	goto L8
L10:
	;
	v21 = v10
	v22 = v9
	goto L11
L11:
	;
	v25 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v22)+1)))
	v26 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v21)+1)))
	if v26 == int32(0) {
		v36 = v26
		v37 = v25
		goto L9
	} else {
		goto L13
	}
L12:
	;
	v36 = v26
	v37 = v25
	goto L9
L13:
	;
	v29 = int32(1)
	if v26 == v25 {
		v21 = v21 + v29
		v22 = v22 + v29
		goto L11
	} else {
		goto L14
	}
L14:
	;
	goto L12
L15:
	;
	v50 = v3
	goto L1
L16:
	;
	goto L3
L17:
	;
	return int32(0)
L18:
	;
	v50 = v44
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
	var v48 int32
	_ = v48
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
		v48 = v16
		goto L11
	} else {
		goto L12
	}
L11:
	;
	v54 = v48
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
		v48 = v40
		goto L11
	} else {
		goto L18
	}
L17:
	;
	v48 = v40
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
	F_errmsg(m, int32(_a_F_check_can_set_role_0), v6)
	mBase = m.M
	v68 = m.ExcPending
	if v68 != 0 {
		goto L3
	} else {
		goto L24
	}
L24:
	;
	F_errfinish(m, int32(_a_F_check_can_set_role_1), int32(_a_F_check_can_set_role_2), int32(_a_F_check_can_set_role_3))
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
