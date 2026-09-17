package p5

import (
	base "github.com/shibukawa/pgmem/internal/aot/pgaot/base"
	"unsafe"
)

func F_AddRoleMems(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32, l5 int32, l6 int32) {
	mBase := m.M
	_ = mBase
	var v8 int32
	_ = v8
	var v18 int32
	_ = v18
	var v20 int32
	_ = v20
	var v23 int32
	_ = v23
	var v24 int32
	_ = v24
	var v27 int32
	_ = v27
	var v28 int32
	_ = v28
	var v29 int32
	_ = v29
	var v34 int32
	_ = v34
	var v40 int32
	_ = v40
	var v52 int32
	_ = v52
	var v56 int32
	_ = v56
	var v58 int32
	_ = v58
	var v62 int32
	_ = v62
	var v67 int32
	_ = v67
	var v70 int32
	_ = v70
	var v74 int32
	_ = v74
	var v78 int32
	_ = v78
	var v83 int32
	_ = v83
	var v84 int32
	_ = v84
	var v85 int32
	_ = v85
	var v86 int32
	_ = v86
	var v90 int32
	_ = v90
	var v95 int32
	_ = v95
	var v96 int32
	_ = v96
	var v102 int32
	_ = v102
	var v105 int32
	_ = v105
	var v106 int32
	_ = v106
	var v107 int32
	_ = v107
	var v114 int32
	_ = v114
	var v119 int32
	_ = v119
	var v122 int32
	_ = v122
	var v123 int32
	_ = v123
	var v124 int32
	_ = v124
	var v133 int32
	_ = v133
	var v151 int32
	_ = v151
	var v152 int32
	_ = v152
	var v154 int32
	_ = v154
	var v163 int32
	_ = v163
	var v173 int32
	_ = v173
	var v178 int32
	_ = v178
	var v189 int32
	_ = v189
	var v195 int32
	_ = v195
	var v199 int32
	_ = v199
	var v205 int32
	_ = v205
	var v208 int32
	_ = v208
	var v215 int32
	_ = v215
	var v220 int32
	_ = v220
	var v221 int32
	_ = v221
	var v224 int32
	_ = v224
	var v229 int32
	_ = v229
	var v244 int32
	_ = v244
	var v245 int32
	_ = v245
	var v246 int32
	_ = v246
	var v248 int32
	_ = v248
	var v253 int32
	_ = v253
	var v254 int32
	_ = v254
	var v255 int32
	_ = v255
	var v257 int32
	_ = v257
	var v259 int32
	_ = v259
	var v277 int32
	_ = v277
	var v278 int32
	_ = v278
	var v283 int32
	_ = v283
	var v286 int32
	_ = v286
	var v287 int32
	_ = v287
	var v288 int32
	_ = v288
	var v294 int32
	_ = v294
	var v299 int32
	_ = v299
	var v300 int32
	_ = v300
	var v327 int32
	_ = v327
	var v340 int32
	_ = v340
	var v342 int32
	_ = v342
	var v344 int32
	_ = v344
	var v345 int32
	_ = v345
	var v346 int32
	_ = v346
	var v347 int32
	_ = v347
	var v348 int32
	_ = v348
	var v350 int32
	_ = v350
	var v353 int32
	_ = v353
	var v375 int32
	_ = v375
	var v378 int32
	_ = v378
	var v385 int32
	_ = v385
	var v390 int32
	_ = v390
	var v392 int32
	_ = v392
	var v416 int32
	_ = v416
	var v428 int32
	_ = v428
	var v432 int32
	_ = v432
	var v434 int32
	_ = v434
	var v438 int32
	_ = v438
	var v443 int32
	_ = v443
	var v446 int32
	_ = v446
	var v452 int32
	_ = v452
	var v453 int32
	_ = v453
	var v454 int32
	_ = v454
	var v456 int64
	_ = v456
	var v474 int32
	_ = v474
	var v475 int32
	_ = v475
	var v476 int32
	_ = v476
	var v477 int32
	_ = v477
	var v478 int32
	_ = v478
	var v479 int32
	_ = v479
	var v480 int32
	_ = v480
	var v485 int32
	_ = v485
	var v486 int32
	_ = v486
	var v488 int32
	_ = v488
	var v492 int32
	_ = v492
	var v498 int32
	_ = v498
	var v499 int32
	_ = v499
	var v501 int32
	_ = v501
	var v505 int32
	_ = v505
	var v511 int32
	_ = v511
	var v512 int32
	_ = v512
	var v514 int32
	_ = v514
	var v520 int32
	_ = v520
	var v521 int32
	_ = v521
	var v522 int32
	_ = v522
	var v523 int32
	_ = v523
	var v525 int32
	_ = v525
	var v526 int32
	_ = v526
	var v534 int32
	_ = v534
	var v539 int32
	_ = v539
	var v542 int32
	_ = v542
	var v546 int32
	_ = v546
	var v547 int32
	_ = v547
	var v548 int32
	_ = v548
	var v550 int32
	_ = v550
	var v552 int32
	_ = v552
	var v555 int32
	_ = v555
	var v558 int32
	_ = v558
	var v559 int32
	_ = v559
	var v562 int32
	_ = v562
	var v563 int32
	_ = v563
	var v565 int32
	_ = v565
	var v568 int32
	_ = v568
	var v573 int32
	_ = v573
	var v574 int32
	_ = v574
	var v580 int32
	_ = v580
	var v581 int32
	_ = v581
	var v583 int32
	_ = v583
	var v586 int32
	_ = v586
	var v592 int32
	_ = v592
	var v596 int32
	_ = v596
	var v603 int32
	_ = v603
	var v607 int32
	_ = v607
	var v612 int32
	_ = v612
	var v620 int32
	_ = v620
	var v621 int32
	_ = v621
	var v625 int32
	_ = v625
	var v627 int32
	_ = v627
	var v634 int32
	_ = v634
	v8 = int32(0)
	v18 = m.G0
	v20 = v18 - int32(144)
	m.G0 = v20
	v23 = F_check_role_grantor(m, l0, l2, l5, int32(1))
	mBase = m.M
	v24 = m.ExcPending
	if v24 != 0 {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	return
L2:
	;
	v27 = F_table_open(m, int32(1261), int32(3))
	mBase = m.M
	v28 = m.ExcPending
	if v28 != 0 {
		goto L1
	} else {
		goto L3
	}
L3:
	;
	v29 = *(*int32)(unsafe.Add(mBase, uint32(v27)+52))
	F_LockSharedObject(m, int32(1260), l2, int32(4))
	mBase = m.M
	v34 = m.ExcPending
	if v34 != 0 {
		goto L1
	} else {
		goto L4
	}
L4:
	;
	v40 = int32(0)
	goto L10
L5:
	;
	v416 = int32(0)
	goto L79
L6:
	;
	if int32(0) < v300 {
		goto L64
	} else {
		goto L65
	}
L7:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v283 = m.ExcPending
	if v283 != 0 {
		goto L1
	} else {
		goto L58
	}
L8:
	;
	if l4 == int32(0) {
		v300 = v154
		goto L6
	} else {
		goto L36
	}
L9:
	;
	v122 = F_palloc(m, v85<<(uint(int32(2))%32))
	mBase = m.M
	v123 = m.ExcPending
	if v123 != 0 {
		goto L1
	} else {
		goto L31
	}
L10:
	;
	v52 = int32(0)
	if l3 == v52 {
		v62 = v52
		goto L12
	} else {
		goto L13
	}
L11:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v102 = m.ExcPending
	if v102 != 0 {
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
	v56 = *(*int32)(unsafe.Add(mBase, uint32(l3)+4))
	if v56 <= v40 {
		v62 = int32(0)
		goto L12
	} else {
		goto L14
	}
L14:
	;
	v58 = *(*int32)(unsafe.Add(mBase, uint32(l3)+12))
	v62 = v58 + v40<<(uint(int32(2))%32)
	goto L12
L15:
	;
	v86 = *(*int32)(unsafe.Add(mBase, uint32(v62)))
	v90 = *(*int32)(unsafe.Add(mBase, uint32(v70+v40<<(uint(int32(2))%32))))
	if v90 == int32(_a_F_AddRoleMems_0) {
		goto L7
	} else {
		goto L23
	}
L16:
	;
	v74 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l6)+4)))
	if base.B2i32(v23 == int32(10))|base.B2i32(v74 != int32(1)) != 0 {
		goto L5
	} else {
		goto L20
	}
L17:
	;
	v67 = *(*int32)(unsafe.Add(mBase, uint32(l4)+4))
	if base.B2i32(v62 == int32(0))|base.B2i32(v67 <= v40) != 0 {
		goto L16
	} else {
		goto L18
	}
L18:
	;
	v70 = *(*int32)(unsafe.Add(mBase, uint32(l4)+12))
	if v70 != 0 {
		goto L15
	} else {
		goto L19
	}
L19:
	;
	goto L16
L20:
	;
	v78 = int32(0)
	v83 = F_SearchSysCacheList(m, int32(9), int32(1), l2, v78, v78)
	mBase = m.M
	v84 = m.ExcPending
	if v84 != 0 {
		goto L1
	} else {
		goto L21
	}
L21:
	;
	v85 = *(*int32)(unsafe.Add(mBase, uint32(v83)+40))
	if v85 != 0 {
		goto L9
	} else {
		goto L22
	}
L22:
	;
	v154 = v78
	v163 = v8
	goto L8
L23:
	;
	v95 = F_is_member_of_role_nosuper(m, l2, v90)
	mBase = m.M
	v96 = m.ExcPending
	if v96 != 0 {
		goto L1
	} else {
		goto L24
	}
L24:
	;
	if v95 == int32(0) {
		v40 = v40 + int32(1)
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
	v105 = m.ExcPending
	if v105 != 0 {
		goto L1
	} else {
		goto L27
	}
L27:
	;
	v106 = F_get_rolespec_name(m, v86)
	mBase = m.M
	v107 = m.ExcPending
	if v107 != 0 {
		goto L1
	} else {
		goto L28
	}
L28:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v20)+84)) = v106
	*(*int32)(unsafe.Add(mBase, uint32(v20)+80)) = l1
	F_errmsg(m, int32(_a_F_AddRoleMems_1), v20+int32(80))
	mBase = m.M
	v114 = m.ExcPending
	if v114 != 0 {
		goto L1
	} else {
		goto L29
	}
L29:
	;
	F_errfinish(m, int32(_a_F_AddRoleMems_2), int32(1748), int32(_a_F_AddRoleMems_3))
	mBase = m.M
	v119 = m.ExcPending
	if v119 != 0 {
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
	v124 = *(*int32)(unsafe.Add(mBase, uint32(v83)+40))
	if v124 <= int32(0) {
		v154 = v124
		v163 = v122
		goto L8
	} else {
		goto L32
	}
L32:
	;
	v133 = int32(0)
	goto L33
L33:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v122+v133<<(uint(int32(2))%32)))) = int32(0)
	v151 = v133 + int32(1)
	v152 = *(*int32)(unsafe.Add(mBase, uint32(v83)+40))
	if v151 < v152 {
		v133 = v151
		goto L33
	} else {
		goto L35
	}
L34:
	;
	v154 = v152
	v163 = v122
	goto L8
L35:
	;
	goto L34
L36:
	;
	v173 = *(*int32)(unsafe.Add(mBase, uint32(l4)+4))
	if v173 <= int32(0) {
		v300 = v154
		goto L6
	} else {
		goto L37
	}
L37:
	;
	v178 = v154
	v189 = v8
	goto L38
L38:
	;
	v195 = *(*int32)(unsafe.Add(mBase, uint32(l4)+12))
	v199 = *(*int32)(unsafe.Add(mBase, uint32(v195+v189<<(uint(int32(2))%32))))
	if v199 == int32(10) {
		goto L40
	} else {
		goto L41
	}
L39:
	;
	v300 = v259
	goto L6
L40:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v205 = m.ExcPending
	if v205 != 0 {
		goto L1
	} else {
		goto L43
	}
L41:
	;
	goto L42
L42:
	;
	v221 = int32(0)
	if v221 < v178 {
		goto L47
	} else {
		goto L48
	}
L43:
	;
	F_errcode(m, int32(16910080))
	mBase = m.M
	v208 = m.ExcPending
	if v208 != 0 {
		goto L1
	} else {
		goto L44
	}
L44:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v20)+48)) = int32(_a_F_AddRoleMems_4)
	F_errmsg(m, int32(_a_F_AddRoleMems_5), v20+int32(48))
	mBase = m.M
	v215 = m.ExcPending
	if v215 != 0 {
		goto L1
	} else {
		goto L45
	}
L45:
	;
	F_errfinish(m, int32(_a_F_AddRoleMems_2), int32(1788), int32(_a_F_AddRoleMems_3))
	mBase = m.M
	v220 = m.ExcPending
	if v220 != 0 {
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
	v224 = v178
	v229 = v221
	goto L50
L48:
	;
	v259 = v178
	goto L49
L49:
	;
	v277 = v189 + int32(1)
	v278 = *(*int32)(unsafe.Add(mBase, uint32(l4)+4))
	if v277 < v278 {
		v178 = v259
		v189 = v277
		goto L38
	} else {
		goto L57
	}
L50:
	;
	v244 = *(*int32)(unsafe.Add(mBase, uint32(v83+int32(48)+v229<<(uint(int32(2))%32))))
	v245 = *(*int32)(unsafe.Add(mBase, uint32(v244)+56))
	v246 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v245)+22)))
	v248 = *(*int32)(unsafe.Add(mBase, uint32(v245+v246)+8))
	if v199 == v248 {
		goto L52
	} else {
		goto L53
	}
L51:
	;
	v259 = v255
	goto L49
L52:
	;
	F_plan_recursive_revoke(m, v83, v163, v229, int32(0), int32(1))
	mBase = m.M
	v253 = m.ExcPending
	if v253 != 0 {
		goto L1
	} else {
		goto L55
	}
L53:
	;
	v255 = v224
	goto L54
L54:
	;
	v257 = v229 + int32(1)
	if v257 < v255 {
		v224 = v255
		v229 = v257
		goto L50
	} else {
		goto L56
	}
L55:
	;
	v254 = *(*int32)(unsafe.Add(mBase, uint32(v83)+40))
	v255 = v254
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
	v286 = m.ExcPending
	if v286 != 0 {
		goto L1
	} else {
		goto L59
	}
L59:
	;
	v287 = F_get_rolespec_name(m, v86)
	mBase = m.M
	v288 = m.ExcPending
	if v288 != 0 {
		goto L1
	} else {
		goto L60
	}
L60:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v20)+64)) = v287
	F_errmsg(m, int32(_a_F_AddRoleMems_6), v20-int32(-64))
	mBase = m.M
	v294 = m.ExcPending
	if v294 != 0 {
		goto L1
	} else {
		goto L61
	}
L61:
	;
	F_errfinish(m, int32(_a_F_AddRoleMems_2), int32(1735), int32(_a_F_AddRoleMems_3))
	mBase = m.M
	v299 = m.ExcPending
	if v299 != 0 {
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
	F_ReleaseCatCacheList(m, v83)
	mBase = m.M
	v392 = m.ExcPending
	if v392 != 0 {
		goto L1
	} else {
		goto L78
	}
L64:
	;
	v327 = int32(0)
	goto L67
L65:
	;
	goto L66
L66:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v375 = m.ExcPending
	if v375 != 0 {
		goto L1
	} else {
		goto L74
	}
L67:
	;
	v340 = v327 << (uint(int32(2)) % 32)
	v342 = *(*int32)(unsafe.Add(mBase, uint32(v163+v340)))
	if v342 != 0 {
		goto L69
	} else {
		goto L70
	}
L68:
	;
	goto L66
L69:
	;
	v353 = v327 + int32(1)
	if v353 != v300 {
		v327 = v353
		goto L67
	} else {
		goto L73
	}
L70:
	;
	v344 = *(*int32)(unsafe.Add(mBase, uint32(v340+(v83+int32(48)))))
	v345 = *(*int32)(unsafe.Add(mBase, uint32(v344)+56))
	v346 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v345)+22)))
	v347 = v345 + v346
	v348 = *(*int32)(unsafe.Add(mBase, uint32(v347)+8))
	if v348 != v23 {
		goto L69
	} else {
		goto L71
	}
L71:
	;
	v350 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v347)+16)))
	if v350 != 0 {
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
	v378 = m.ExcPending
	if v378 != 0 {
		goto L1
	} else {
		goto L75
	}
L75:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v20)+32)) = int32(_a_F_AddRoleMems_4)
	F_errmsg(m, int32(_a_F_AddRoleMems_5), v20+int32(32))
	mBase = m.M
	v385 = m.ExcPending
	if v385 != 0 {
		goto L1
	} else {
		goto L76
	}
L76:
	;
	F_errfinish(m, int32(_a_F_AddRoleMems_2), int32(1814), int32(_a_F_AddRoleMems_3))
	mBase = m.M
	v390 = m.ExcPending
	if v390 != 0 {
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
	v428 = int32(0)
	if l3 == v428 {
		v438 = v428
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
	v432 = *(*int32)(unsafe.Add(mBase, uint32(l3)+4))
	if v432 <= v416 {
		v438 = int32(0)
		goto L81
	} else {
		goto L83
	}
L83:
	;
	v434 = *(*int32)(unsafe.Add(mBase, uint32(l3)+12))
	v438 = v434 + v416<<(uint(int32(2))%32)
	goto L81
L84:
	;
	F_CommandCounterIncrement(m)
	mBase = m.M
	v634 = m.ExcPending
	if v634 != 0 {
		goto L1
	} else {
		goto L133
	}
L85:
	;
	v620 = F_heap_modify_tuple(m, v474, v29, v20+int32(112), v20+int32(104), v20+int32(96))
	mBase = m.M
	v621 = m.ExcPending
	if v621 != 0 {
		goto L1
	} else {
		goto L130
	}
L86:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v603 = m.ExcPending
	if v603 != 0 {
		goto L1
	} else {
		goto L127
	}
L87:
	;
	F_relation_close(m, v27, int32(0))
	mBase = m.M
	v596 = m.ExcPending
	if v596 != 0 {
		goto L1
	} else {
		goto L126
	}
L88:
	;
	v443 = *(*int32)(unsafe.Add(mBase, uint32(l4)+4))
	if base.B2i32(v438 == int32(0))|base.B2i32(v443 <= v416) != 0 {
		goto L87
	} else {
		goto L89
	}
L89:
	;
	v446 = *(*int32)(unsafe.Add(mBase, uint32(l4)+12))
	if v446 == int32(0) {
		goto L87
	} else {
		goto L90
	}
L90:
	;
	v452 = *(*int32)(unsafe.Add(mBase, uint32(v446+v416<<(uint(int32(2))%32))))
	v453 = *(*int32)(unsafe.Add(mBase, uint32(v438)))
	v454 = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v20)+136)) = v454
	v456 = int64(0)
	*(*int64)(unsafe.Add(mBase, uint32(v20)+128)) = v456
	*(*int64)(unsafe.Add(mBase, uint32(v20)+120)) = v456
	*(*int64)(unsafe.Add(mBase, uint32(v20)+112)) = v456
	*(*int32)(unsafe.Add(mBase, uint32(v20)+107)) = v454
	*(*int32)(unsafe.Add(mBase, uint32(v20)+104)) = v454
	*(*int32)(unsafe.Add(mBase, uint32(v20)+124)) = v23
	*(*int32)(unsafe.Add(mBase, uint32(v20)+120)) = v452
	*(*int32)(unsafe.Add(mBase, uint32(v20)+116)) = l2
	*(*int32)(unsafe.Add(mBase, uint32(v20)+99)) = v454
	*(*int32)(unsafe.Add(mBase, uint32(v20)+96)) = v454
	v474 = F_SearchSysCache3(m, int32(9), l2, v452, v23)
	mBase = m.M
	v475 = m.ExcPending
	if v475 != 0 {
		goto L1
	} else {
		goto L91
	}
L91:
	;
	if v474 != 0 {
		goto L92
	} else {
		goto L93
	}
L92:
	;
	v476 = *(*int32)(unsafe.Add(mBase, uint32(v474)+16))
	v477 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v476)+22)))
	v478 = v476 + v477
	v479 = int32(0)
	v480 = *(*int32)(unsafe.Add(mBase, uint32(l6)))
	if v480&int32(1) == v479 {
		v492 = v479
		goto L95
	} else {
		goto L96
	}
L93:
	;
	goto L94
L94:
	;
	v546 = F_palloc(m, int32(4))
	mBase = m.M
	v547 = m.ExcPending
	if v547 != 0 {
		goto L1
	} else {
		goto L114
	}
L95:
	;
	if v480&int32(2) == int32(0) {
		v505 = v492
		goto L98
	} else {
		goto L99
	}
L96:
	;
	v485 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l6)+4)))
	v486 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v478)+16)))
	if v485 == v486 {
		v492 = v479
		goto L95
	} else {
		goto L97
	}
L97:
	;
	v488 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(v20)+100)) = uint8(v488)
	*(*int32)(unsafe.Add(mBase, uint32(v20)+128)) = v485
	v492 = v488
	goto L95
L98:
	;
	if v480&int32(4) == int32(0) {
		goto L101
	} else {
		goto L102
	}
L99:
	;
	v498 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l6)+5)))
	v499 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v478)+17)))
	if v498 == v499 {
		v505 = v492
		goto L98
	} else {
		goto L100
	}
L100:
	;
	v501 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(v20)+101)) = uint8(v501)
	*(*int32)(unsafe.Add(mBase, uint32(v20)+132)) = v498
	v505 = v501
	goto L98
L101:
	;
	if v505 != 0 {
		goto L85
	} else {
		goto L104
	}
L102:
	;
	v511 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l6)+6)))
	v512 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v478)+18)))
	if v511 == v512 {
		goto L101
	} else {
		goto L103
	}
L103:
	;
	v514 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(v20)+102)) = uint8(v514)
	*(*int32)(unsafe.Add(mBase, uint32(v20)+136)) = v511
	goto L85
L104:
	;
	v520 = F_errstart(m, int32(18), int32(0))
	mBase = m.M
	v521 = m.ExcPending
	if v521 != 0 {
		goto L1
	} else {
		goto L105
	}
L105:
	;
	if v520 != 0 {
		goto L106
	} else {
		goto L107
	}
L106:
	;
	v522 = F_get_rolespec_name(m, v453)
	mBase = m.M
	v523 = m.ExcPending
	if v523 != 0 {
		goto L1
	} else {
		goto L109
	}
L107:
	;
	goto L108
L108:
	;
	F_ReleaseCatCache(m, v474)
	mBase = m.M
	v542 = m.ExcPending
	if v542 != 0 {
		goto L1
	} else {
		goto L113
	}
L109:
	;
	v525 = F_GetUserNameFromId(m, v23, int32(0))
	mBase = m.M
	v526 = m.ExcPending
	if v526 != 0 {
		goto L1
	} else {
		goto L110
	}
L110:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v20)+24)) = v525
	*(*int32)(unsafe.Add(mBase, uint32(v20)+20)) = l1
	*(*int32)(unsafe.Add(mBase, uint32(v20)+16)) = v522
	F_errmsg(m, int32(_a_F_AddRoleMems_7), v20+int32(16))
	mBase = m.M
	v534 = m.ExcPending
	if v534 != 0 {
		goto L1
	} else {
		goto L111
	}
L111:
	;
	F_errfinish(m, int32(_a_F_AddRoleMems_2), int32(1892), int32(_a_F_AddRoleMems_3))
	mBase = m.M
	v539 = m.ExcPending
	if v539 != 0 {
		goto L1
	} else {
		goto L112
	}
L112:
	;
	goto L108
L113:
	;
	v416 = v416 + int32(1)
	goto L79
L114:
	;
	v548 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l6)+4)))
	*(*int32)(unsafe.Add(mBase, uint32(v20)+128)) = v548
	v550 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l6)+6)))
	*(*int32)(unsafe.Add(mBase, uint32(v20)+136)) = v550
	v552 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l6))))
	if v552&int32(2) != 0 {
		goto L116
	} else {
		goto L117
	}
L115:
	;
	v573 = F_GetNewOidWithIndex(m, v27, int32(_a_F_AddRoleMems_8), int32(1))
	mBase = m.M
	v574 = m.ExcPending
	if v574 != 0 {
		goto L1
	} else {
		goto L122
	}
L116:
	;
	v555 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l6)+5)))
	*(*int32)(unsafe.Add(mBase, uint32(v20)+132)) = v555
	goto L115
L117:
	;
	goto L118
L118:
	;
	v558 = F_SearchSysCache1(m, int32(11), v452)
	mBase = m.M
	v559 = m.ExcPending
	if v559 != 0 {
		goto L1
	} else {
		goto L119
	}
L119:
	;
	if v558 == int32(0) {
		goto L86
	} else {
		goto L120
	}
L120:
	;
	v562 = *(*int32)(unsafe.Add(mBase, uint32(v558)+16))
	v563 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v562)+22)))
	v565 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v562+v563)+69)))
	*(*int32)(unsafe.Add(mBase, uint32(v20)+132)) = v565
	F_ReleaseCatCache(m, v558)
	mBase = m.M
	v568 = m.ExcPending
	if v568 != 0 {
		goto L1
	} else {
		goto L121
	}
L121:
	;
	goto L115
L122:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v20)+112)) = v573
	v580 = F_heap_form_tuple(m, v29, v20+int32(112), v20+int32(104))
	mBase = m.M
	v581 = m.ExcPending
	if v581 != 0 {
		goto L1
	} else {
		goto L123
	}
L123:
	;
	F_CatalogTupleInsert(m, v27, v580)
	mBase = m.M
	v583 = m.ExcPending
	if v583 != 0 {
		goto L1
	} else {
		goto L124
	}
L124:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v546))) = v23
	v586 = int32(0)
	F_updateAclDependencies(m, int32(1261), v573, v586, v586, v586, v586, int32(1), v546)
	mBase = m.M
	v592 = m.ExcPending
	if v592 != 0 {
		goto L1
	} else {
		goto L125
	}
L125:
	;
	goto L84
L126:
	;
	m.G0 = v20 + int32(144)
	return
L127:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v20))) = v452
	F_errmsg_internal(m, int32(_a_F_AddRoleMems_9), v20)
	mBase = m.M
	v607 = m.ExcPending
	if v607 != 0 {
		goto L1
	} else {
		goto L128
	}
L128:
	;
	F_errfinish(m, int32(_a_F_AddRoleMems_2), int32(1934), int32(_a_F_AddRoleMems_3))
	mBase = m.M
	v612 = m.ExcPending
	if v612 != 0 {
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
	F_CatalogTupleUpdate(m, v27, v620+int32(4), v620)
	mBase = m.M
	v625 = m.ExcPending
	if v625 != 0 {
		goto L1
	} else {
		goto L131
	}
L131:
	;
	F_ReleaseCatCache(m, v474)
	mBase = m.M
	v627 = m.ExcPending
	if v627 != 0 {
		goto L1
	} else {
		goto L132
	}
L132:
	;
	goto L84
L133:
	;
	v416 = v416 + int32(1)
	goto L79
}
func F_RoleMembershipCacheCallback(m *base.Module, l0 int32, l1 int32, l2 int32) {
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
