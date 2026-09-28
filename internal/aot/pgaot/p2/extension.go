package p2

import (
	base "github.com/shibukawa/pgmem/internal/aot/pgaot/base"
	"unsafe"
)

func F_CreateExtensionInternal(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32, l5 int32, l6 int32) {
	mBase := m.M
	_ = mBase
	var v8 int32
	_ = v8
	var v22 int32
	_ = v22
	var v24 int32
	_ = v24
	var v27 int32
	_ = v27
	var v29 int32
	_ = v29
	var v30 int32
	_ = v30
	var v31 int32
	_ = v31
	var v32 int32
	_ = v32
	var v35 int32
	_ = v35
	var v37 int32
	_ = v37
	var v42 int32
	_ = v42
	var v45 int32
	_ = v45
	var v48 int32
	_ = v48
	var v50 int32
	_ = v50
	var v52 int32
	_ = v52
	var v53 int32
	_ = v53
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
	var v81 int32
	_ = v81
	var v91 int32
	_ = v91
	var v92 int32
	_ = v92
	var v95 int32
	_ = v95
	var v98 int32
	_ = v98
	var v101 int32
	_ = v101
	var v102 int32
	_ = v102
	var v105 int32
	_ = v105
	var v106 int32
	_ = v106
	var v109 int32
	_ = v109
	var v116 int32
	_ = v116
	var v117 int32
	_ = v117
	var v122 int32
	_ = v122
	var v146 int32
	_ = v146
	var v147 int32
	_ = v147
	var v148 int32
	_ = v148
	var v149 int32
	_ = v149
	var v152 int32
	_ = v152
	var v157 int32
	_ = v157
	var v158 int32
	_ = v158
	var v166 int32
	_ = v166
	var v169 int32
	_ = v169
	var v180 int32
	_ = v180
	var v183 int32
	_ = v183
	var v199 int32
	_ = v199
	var v201 int32
	_ = v201
	var v205 int32
	_ = v205
	var v208 int32
	_ = v208
	var v212 int32
	_ = v212
	var v213 int32
	_ = v213
	var v216 int32
	_ = v216
	var v218 int32
	_ = v218
	var v219 int32
	_ = v219
	var v224 int32
	_ = v224
	var v227 int32
	_ = v227
	var v230 int32
	_ = v230
	var v232 int32
	_ = v232
	var v234 int32
	_ = v234
	var v235 int32
	_ = v235
	var v238 int32
	_ = v238
	var v241 int32
	_ = v241
	var v244 int32
	_ = v244
	var v245 int32
	_ = v245
	var v248 int32
	_ = v248
	var v249 int32
	_ = v249
	var v252 int32
	_ = v252
	var v259 int32
	_ = v259
	var v260 int32
	_ = v260
	var v267 int32
	_ = v267
	var v270 int32
	_ = v270
	var v272 int32
	_ = v272
	var v273 int32
	_ = v273
	var v289 int32
	_ = v289
	var v295 int32
	_ = v295
	var v298 int32
	_ = v298
	var v302 int32
	_ = v302
	var v317 int32
	_ = v317
	var v321 int32
	_ = v321
	var v322 int32
	_ = v322
	var v323 int64
	_ = v323
	var v325 int64
	_ = v325
	var v327 int64
	_ = v327
	var v329 int64
	_ = v329
	var v331 int64
	_ = v331
	var v333 int64
	_ = v333
	var v336 int32
	_ = v336
	var v339 int32
	_ = v339
	var v341 int32
	_ = v341
	var v342 int32
	_ = v342
	var v343 int32
	_ = v343
	var v348 int32
	_ = v348
	var v351 int32
	_ = v351
	var v354 int32
	_ = v354
	var v355 int32
	_ = v355
	var v358 int32
	_ = v358
	var v359 int32
	_ = v359
	var v362 int32
	_ = v362
	var v369 int32
	_ = v369
	var v370 int32
	_ = v370
	var v373 int32
	_ = v373
	var v375 int32
	_ = v375
	var v376 int32
	_ = v376
	var v378 int32
	_ = v378
	var v379 int32
	_ = v379
	var v381 int32
	_ = v381
	var v382 int32
	_ = v382
	var v390 int32
	_ = v390
	var v392 int32
	_ = v392
	var v395 int32
	_ = v395
	var v397 int32
	_ = v397
	var v398 int32
	_ = v398
	var v402 int32
	_ = v402
	var v403 int32
	_ = v403
	var v406 int32
	_ = v406
	var v407 int32
	_ = v407
	var v408 int32
	_ = v408
	var v409 int32
	_ = v409
	var v413 int32
	_ = v413
	var v415 int32
	_ = v415
	var v416 int32
	_ = v416
	var v419 int32
	_ = v419
	var v424 int32
	_ = v424
	var v426 int32
	_ = v426
	var v430 int32
	_ = v430
	var v433 int32
	_ = v433
	var v435 int32
	_ = v435
	var v436 int32
	_ = v436
	var v440 int32
	_ = v440
	var v449 int32
	_ = v449
	var v452 int32
	_ = v452
	var v456 int32
	_ = v456
	var v464 int32
	_ = v464
	var v468 int32
	_ = v468
	var v469 int32
	_ = v469
	var v470 int32
	_ = v470
	var v472 int32
	_ = v472
	var v473 int32
	_ = v473
	var v477 int32
	_ = v477
	var v478 int32
	_ = v478
	var v480 int32
	_ = v480
	var v482 int32
	_ = v482
	var v483 int32
	_ = v483
	var v484 int32
	_ = v484
	var v485 int32
	_ = v485
	var v486 int32
	_ = v486
	var v487 int32
	_ = v487
	var v489 int32
	_ = v489
	var v490 int32
	_ = v490
	var v499 int32
	_ = v499
	var v502 int32
	_ = v502
	var v515 int32
	_ = v515
	var v516 int32
	_ = v516
	var v517 int64
	_ = v517
	var v520 int32
	_ = v520
	var v521 int32
	_ = v521
	var v523 int64
	_ = v523
	var v525 int32
	_ = v525
	var v526 int32
	_ = v526
	var v530 int32
	_ = v530
	var v533 int32
	_ = v533
	var v535 int32
	_ = v535
	var v542 int32
	_ = v542
	var v545 int32
	_ = v545
	var v549 int32
	_ = v549
	var v554 int32
	_ = v554
	var v579 int32
	_ = v579
	var v582 int32
	_ = v582
	var v583 int32
	_ = v583
	var v590 int32
	_ = v590
	var v595 int32
	_ = v595
	var v599 int32
	_ = v599
	var v602 int32
	_ = v602
	var v603 int32
	_ = v603
	var v604 int32
	_ = v604
	var v609 int32
	_ = v609
	var v614 int32
	_ = v614
	var v618 int32
	_ = v618
	var v621 int32
	_ = v621
	var v625 int32
	_ = v625
	var v630 int32
	_ = v630
	var v634 int32
	_ = v634
	var v637 int32
	_ = v637
	var v641 int32
	_ = v641
	var v646 int32
	_ = v646
	v8 = int32(0)
	v22 = m.G0
	v24 = v22 - int32(144)
	m.G0 = v24
	v27 = *(*int32)(unsafe.Add(mBase, _c_F_CreateExtensionInternal[0]))
	v29 = F_palloc0(m, int32(48))
	mBase = m.M
	v30 = m.ExcPending
	if v30 != 0 {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	return
L2:
	;
	v31 = F_pstrdup(m, l1)
	mBase = m.M
	v32 = m.ExcPending
	if v32 != 0 {
		goto L1
	} else {
		goto L3
	}
L3:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v29)+36)) = int32(-1)
	v35 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v29)+34)) = uint8(v35)
	v37 = int32(256)
	*(*uint16)(unsafe.Add(mBase, uint32(v29)+32)) = uint16(v37)
	*(*int32)(unsafe.Add(mBase, uint32(v29))) = v31
	F_parse_extension_control_file(m, v29, v35)
	mBase = m.M
	v42 = m.ExcPending
	if v42 != 0 {
		goto L1
	} else {
		goto L4
	}
L4:
	;
	if l3 == int32(0) {
		goto L10
	} else {
		goto L11
	}
L5:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v634 = m.ExcPending
	if v634 != 0 {
		goto L1
	} else {
		goto L149
	}
L6:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v618 = m.ExcPending
	if v618 != 0 {
		goto L1
	} else {
		goto L145
	}
L7:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v599 = m.ExcPending
	if v599 != 0 {
		goto L1
	} else {
		goto L141
	}
L8:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v579 = m.ExcPending
	if v579 != 0 {
		goto L1
	} else {
		goto L137
	}
L9:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v542 = m.ExcPending
	if v542 != 0 {
		goto L1
	} else {
		goto L133
	}
L10:
	;
	v45 = *(*int32)(unsafe.Add(mBase, uint32(v29)+16))
	if v45 == int32(0) {
		goto L9
	} else {
		goto L13
	}
L11:
	;
	v48 = l3
	goto L12
L12:
	;
	F_check_valid_version_name(m, v48)
	mBase = m.M
	v50 = m.ExcPending
	if v50 != 0 {
		goto L1
	} else {
		goto L14
	}
L13:
	;
	v48 = v45
	goto L12
L14:
	;
	v52 = F_get_extension_script_filename(m, v29, int32(0), v48)
	mBase = m.M
	v53 = m.ExcPending
	if v53 != 0 {
		goto L1
	} else {
		goto L15
	}
L15:
	;
	v58 = F___fstatat(m, int32(-100), v52, v24+int32(48), int32(0))
	mBase = m.M
	goto L16
L16:
	;
	if v58 != 0 {
		goto L17
	} else {
		goto L18
	}
L17:
	;
	v59 = F_get_ext_ver_list(m, v29)
	mBase = m.M
	v60 = m.ExcPending
	if v60 != 0 {
		goto L1
	} else {
		goto L22
	}
L18:
	;
	v302 = v48
	v317 = v8
	goto L19
L19:
	;
	v321 = F_palloc(m, int32(48))
	mBase = m.M
	v322 = m.ExcPending
	if v322 != 0 {
		goto L1
	} else {
		goto L70
	}
L20:
	;
	v180 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v166)+8)))
	if v180 != 0 {
		goto L40
	} else {
		goto L41
	}
L21:
	;
	v146 = F_palloc(m, int32(20))
	mBase = m.M
	v147 = m.ExcPending
	if v147 != 0 {
		goto L1
	} else {
		goto L36
	}
L22:
	;
	if v59 == int32(0) {
		goto L21
	} else {
		goto L23
	}
L23:
	;
	v63 = *(*int32)(unsafe.Add(mBase, uint32(v59)+4))
	if v63 <= int32(0) {
		goto L21
	} else {
		goto L24
	}
L24:
	;
	v66 = *(*int32)(unsafe.Add(mBase, uint32(v59)+12))
	v81 = v8
	goto L25
L25:
	;
	v91 = *(*int32)(unsafe.Add(mBase, uint32(v66+v81<<(uint(int32(2))%32))))
	v92 = *(*int32)(unsafe.Add(mBase, uint32(v91)))
	v95 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v92))))
	v98 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v48))))
	if base.B2i32(v95 == int32(0))|base.B2i32(v95 != v98) != 0 {
		v116 = v95
		v117 = v98
		goto L28
	} else {
		goto L29
	}
L26:
	;
	goto L21
L27:
	;
	if v116-v117 == int32(0) {
		v166 = v91
		v169 = v59
		goto L20
	} else {
		goto L34
	}
L28:
	;
	goto L27
L29:
	;
	v101 = v92
	v102 = v48
	goto L30
L30:
	;
	v105 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v102)+1)))
	v106 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v101)+1)))
	if v106 == int32(0) {
		v116 = v106
		v117 = v105
		goto L28
	} else {
		goto L32
	}
L31:
	;
	v116 = v106
	v117 = v105
	goto L28
L32:
	;
	v109 = int32(1)
	if v106 == v105 {
		v101 = v101 + v109
		v102 = v102 + v109
		goto L30
	} else {
		goto L33
	}
L33:
	;
	goto L31
L34:
	;
	v122 = v81 + int32(1)
	if v63 != v122 {
		v81 = v122
		goto L25
	} else {
		goto L35
	}
L35:
	;
	goto L26
L36:
	;
	v148 = F_pstrdup(m, v48)
	mBase = m.M
	v149 = m.ExcPending
	if v149 != 0 {
		goto L1
	} else {
		goto L37
	}
L37:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v146)+12)) = int64(2147483647)
	v152 = int32(0)
	*(*uint16)(unsafe.Add(mBase, uint32(v146)+8)) = uint16(v152)
	*(*int32)(unsafe.Add(mBase, uint32(v146)+4)) = v152
	*(*int32)(unsafe.Add(mBase, uint32(v146))) = v148
	v157 = F_lappend(m, v59, v146)
	mBase = m.M
	v158 = m.ExcPending
	if v158 != 0 {
		goto L1
	} else {
		goto L38
	}
L38:
	;
	v166 = v146
	v169 = v157
	goto L20
L39:
	;
	v298 = *(*int32)(unsafe.Add(mBase, uint32(v289)))
	v302 = v298
	v317 = v295
	goto L19
L40:
	;
	v289 = v166
	v295 = v8
	goto L39
L41:
	;
	goto L42
L42:
	;
	if v169 == int32(0) {
		goto L8
	} else {
		goto L43
	}
L43:
	;
	v183 = *(*int32)(unsafe.Add(mBase, uint32(v169)+4))
	if v183 <= int32(0) {
		goto L8
	} else {
		goto L44
	}
L44:
	;
	v199 = v8
	v201 = int32(0)
	v205 = v8
	goto L45
L45:
	;
	v208 = *(*int32)(unsafe.Add(mBase, uint32(v169)+12))
	v212 = *(*int32)(unsafe.Add(mBase, uint32(v208+v201<<(uint(int32(2))%32))))
	v213 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v212)+8)))
	if v213 != int32(1) {
		v267 = v199
		v270 = v205
		goto L47
	} else {
		goto L48
	}
L46:
	;
	if v267 == int32(0) {
		goto L8
	} else {
		goto L69
	}
L47:
	;
	v272 = v201 + int32(1)
	v273 = *(*int32)(unsafe.Add(mBase, uint32(v169)+4))
	if v272 < v273 {
		v199 = v267
		v201 = v272
		v205 = v270
		goto L45
	} else {
		goto L68
	}
L48:
	;
	v216 = int32(1)
	v218 = F_find_update_path(m, v169, v212, v166, v216, v216)
	mBase = m.M
	v219 = m.ExcPending
	if v219 != 0 {
		goto L1
	} else {
		goto L49
	}
L49:
	;
	if v218 == int32(0) {
		v267 = v199
		v270 = v205
		goto L47
	} else {
		goto L50
	}
L50:
	;
	if v199 == int32(0) {
		goto L51
	} else {
		goto L52
	}
L51:
	;
	v267 = v212
	v270 = v218
	goto L47
L52:
	;
	v224 = *(*int32)(unsafe.Add(mBase, uint32(v218)+4))
	if v205 == int32(0) {
		goto L54
	} else {
		goto L55
	}
L53:
	;
	if v224 != v232 {
		v267 = v199
		v270 = v205
		goto L47
	} else {
		goto L59
	}
L54:
	;
	v227 = int32(0)
	if v227 <= v224 {
		v232 = v227
		goto L53
	} else {
		goto L57
	}
L55:
	;
	goto L56
L56:
	;
	v230 = *(*int32)(unsafe.Add(mBase, uint32(v205)+4))
	if v224 < v230 {
		goto L51
	} else {
		goto L58
	}
L57:
	;
	goto L51
L58:
	;
	v232 = v230
	goto L53
L59:
	;
	v234 = *(*int32)(unsafe.Add(mBase, uint32(v199)))
	v235 = *(*int32)(unsafe.Add(mBase, uint32(v212)))
	v238 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v234))))
	v241 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v235))))
	if base.B2i32(v238 == int32(0))|base.B2i32(v238 != v241) != 0 {
		v259 = v238
		v260 = v241
		goto L61
	} else {
		goto L62
	}
L60:
	;
	if int32(0) <= v259-v260 {
		v267 = v199
		v270 = v205
		goto L47
	} else {
		goto L67
	}
L61:
	;
	goto L60
L62:
	;
	v244 = v234
	v245 = v235
	goto L63
L63:
	;
	v248 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v245)+1)))
	v249 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v244)+1)))
	if v249 == int32(0) {
		v259 = v249
		v260 = v248
		goto L61
	} else {
		goto L65
	}
L64:
	;
	v259 = v249
	v260 = v248
	goto L61
L65:
	;
	v252 = int32(1)
	if v249 == v248 {
		v244 = v244 + v252
		v245 = v245 + v252
		goto L63
	} else {
		goto L66
	}
L66:
	;
	goto L64
L67:
	;
	goto L51
L68:
	;
	goto L46
L69:
	;
	v289 = v267
	v295 = v270
	goto L39
L70:
	;
	v323 = *(*int64)(unsafe.Add(mBase, uint32(v29)+40))
	*(*int64)(unsafe.Add(mBase, uint32(v321)+40)) = v323
	v325 = *(*int64)(unsafe.Add(mBase, uint32(v29)+32))
	*(*int64)(unsafe.Add(mBase, uint32(v321)+32)) = v325
	v327 = *(*int64)(unsafe.Add(mBase, uint32(v29)+24))
	*(*int64)(unsafe.Add(mBase, uint32(v321)+24)) = v327
	v329 = *(*int64)(unsafe.Add(mBase, uint32(v29)+16))
	*(*int64)(unsafe.Add(mBase, uint32(v321)+16)) = v329
	v331 = *(*int64)(unsafe.Add(mBase, uint32(v29)+8))
	*(*int64)(unsafe.Add(mBase, uint32(v321)+8)) = v331
	v333 = *(*int64)(unsafe.Add(mBase, uint32(v29)))
	*(*int64)(unsafe.Add(mBase, uint32(v321))) = v333
	F_parse_extension_control_file(m, v321, v302)
	mBase = m.M
	v336 = m.ExcPending
	if v336 != 0 {
		goto L1
	} else {
		goto L71
	}
L71:
	;
	if l2 == int32(0) {
		goto L76
	} else {
		goto L77
	}
L72:
	;
	v419 = *(*int32)(unsafe.Add(mBase, _c_F_CreateExtensionInternal[1]))
	goto L103
L73:
	;
	v402 = F_fetch_search_path(m, int32(0))
	mBase = m.M
	v403 = m.ExcPending
	if v403 != 0 {
		goto L1
	} else {
		goto L98
	}
L74:
	;
	if v341 != 0 {
		v415 = v341
		v416 = l2
		goto L72
	} else {
		goto L97
	}
L75:
	;
	v375 = F_get_namespace_oid(m, v373, int32(1))
	mBase = m.M
	v376 = m.ExcPending
	if v376 != 0 {
		goto L1
	} else {
		goto L91
	}
L76:
	;
	v339 = *(*int32)(unsafe.Add(mBase, uint32(v321)+28))
	if v339 != 0 {
		v373 = v339
		goto L75
	} else {
		goto L79
	}
L77:
	;
	goto L78
L78:
	;
	v341 = F_get_namespace_oid(m, l2, int32(0))
	mBase = m.M
	v342 = m.ExcPending
	if v342 != 0 {
		goto L1
	} else {
		goto L80
	}
L79:
	;
	goto L73
L80:
	;
	v343 = *(*int32)(unsafe.Add(mBase, uint32(v321)+28))
	if v343 == int32(0) {
		goto L74
	} else {
		goto L81
	}
L81:
	;
	v348 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v343))))
	v351 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l2))))
	if base.B2i32(v348 == int32(0))|base.B2i32(v348 != v351) != 0 {
		v369 = v348
		v370 = v351
		goto L83
	} else {
		goto L84
	}
L82:
	;
	if l4 != 0 {
		v373 = v343
		goto L75
	} else {
		goto L89
	}
L83:
	;
	goto L82
L84:
	;
	v354 = v343
	v355 = l2
	goto L85
L85:
	;
	v358 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v355)+1)))
	v359 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v354)+1)))
	if v359 == int32(0) {
		v369 = v359
		v370 = v358
		goto L83
	} else {
		goto L87
	}
L86:
	;
	v369 = v359
	v370 = v358
	goto L83
L87:
	;
	v362 = int32(1)
	if v359 == v358 {
		v354 = v354 + v362
		v355 = v355 + v362
		goto L85
	} else {
		goto L88
	}
L88:
	;
	goto L86
L89:
	;
	if v369-v370 != 0 {
		goto L7
	} else {
		goto L90
	}
L90:
	;
	v373 = v343
	goto L75
L91:
	;
	if v375 != 0 {
		v415 = v375
		v416 = v373
		goto L72
	} else {
		goto L92
	}
L92:
	;
	v378 = F_make_parsestate(m, int32(0))
	mBase = m.M
	v379 = m.ExcPending
	if v379 != 0 {
		goto L1
	} else {
		goto L93
	}
L93:
	;
	v381 = F_palloc0(m, int32(20))
	mBase = m.M
	v382 = m.ExcPending
	if v382 != 0 {
		goto L1
	} else {
		goto L94
	}
L94:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v381))) = int32(145)
	*(*int32)(unsafe.Add(mBase, uint32(v378)+4)) = int32(_a_F_CreateExtensionInternal_0)
	*(*int64)(unsafe.Add(mBase, uint32(v381)+8)) = int64(0)
	*(*int32)(unsafe.Add(mBase, uint32(v381)+4)) = v373
	v390 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v381)+16)) = uint8(v390)
	v392 = int32(-1)
	F_CreateSchemaCommand(m, v378, v381, v392, v392)
	mBase = m.M
	v395 = m.ExcPending
	if v395 != 0 {
		goto L1
	} else {
		goto L95
	}
L95:
	;
	v397 = F_get_namespace_oid(m, v373, int32(0))
	mBase = m.M
	v398 = m.ExcPending
	if v398 != 0 {
		goto L1
	} else {
		goto L96
	}
L96:
	;
	v415 = v397
	v416 = v373
	goto L72
L97:
	;
	goto L73
L98:
	;
	if v402 == int32(0) {
		goto L6
	} else {
		goto L99
	}
L99:
	;
	v406 = *(*int32)(unsafe.Add(mBase, uint32(v402)+12))
	v407 = *(*int32)(unsafe.Add(mBase, uint32(v406)))
	v408 = F_get_namespace_name(m, v407)
	mBase = m.M
	v409 = m.ExcPending
	if v409 != 0 {
		goto L1
	} else {
		goto L100
	}
L100:
	;
	if v408 == int32(0) {
		goto L5
	} else {
		goto L101
	}
L101:
	;
	F_list_free(m, v402)
	mBase = m.M
	v413 = m.ExcPending
	if v413 != 0 {
		goto L1
	} else {
		goto L102
	}
L102:
	;
	v415 = v407
	v416 = v408
	goto L72
L103:
	;
	if base.B2i32(v419 != int32(0))&base.B2i32(v415 == v419) != 0 {
		goto L104
	} else {
		goto L105
	}
L104:
	;
	v424 = int32(_a_F_CreateExtensionInternal_1)
	v426 = *(*int32)(unsafe.Add(mBase, _c_F_CreateExtensionInternal[2]))
	*(*int32)(unsafe.Add(mBase, _c_F_CreateExtensionInternal[2])) = v426 | int32(1)
	goto L106
L105:
	;
	goto L106
L106:
	;
	v430 = *(*int32)(unsafe.Add(mBase, uint32(v321)+40))
	if v430 == int32(0) {
		goto L108
	} else {
		goto L109
	}
L107:
	;
	v515 = *(*int32)(unsafe.Add(mBase, uint32(v321)))
	v516 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v321)+32)))
	v517 = int64(0)
	F_InsertExtensionTuple(m, v24+int32(36), v515, v27, v415, v516, v302, v517, v517, v499)
	mBase = m.M
	v520 = m.ExcPending
	if v520 != 0 {
		goto L1
	} else {
		goto L126
	}
L108:
	;
	v433 = int32(0)
	v499 = v433
	v502 = v433
	goto L107
L109:
	;
	goto L110
L110:
	;
	v435 = int32(0)
	v436 = *(*int32)(unsafe.Add(mBase, uint32(v430)+4))
	if v436 <= v435 {
		goto L111
	} else {
		goto L112
	}
L111:
	;
	v499 = int32(0)
	v502 = v435
	goto L107
L112:
	;
	goto L113
L113:
	;
	v440 = int32(0)
	v449 = v440
	v452 = v435
	v456 = v440
	goto L114
L114:
	;
	v464 = *(*int32)(unsafe.Add(mBase, uint32(v430)+12))
	v468 = *(*int32)(unsafe.Add(mBase, uint32(v464+v456<<(uint(int32(2))%32))))
	v469 = F_get_required_extension(m, v468, l1, l2, l4, l5, l6)
	mBase = m.M
	v470 = m.ExcPending
	if v470 != 0 {
		goto L1
	} else {
		goto L117
	}
L115:
	;
	v499 = v484
	v502 = v486
	goto L107
L116:
	;
	v484 = F_lappend_oid(m, v449, v469)
	mBase = m.M
	v485 = m.ExcPending
	if v485 != 0 {
		goto L1
	} else {
		goto L123
	}
L117:
	;
	v472 = F_SearchSysCache1(m, int32(28), base.I64_extend_i32_u(v469))
	mBase = m.M
	v473 = m.ExcPending
	if v473 != 0 {
		goto L1
	} else {
		goto L118
	}
L118:
	;
	if v472 == int32(0) {
		goto L119
	} else {
		goto L120
	}
L119:
	;
	v483 = int32(0)
	goto L116
L120:
	;
	goto L121
L121:
	;
	v477 = *(*int32)(unsafe.Add(mBase, uint32(v472)+16))
	v478 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v477)+22)))
	v480 = *(*int32)(unsafe.Add(mBase, uint32(v477+v478)+72))
	F_ReleaseCatCache(m, v472)
	mBase = m.M
	v482 = m.ExcPending
	if v482 != 0 {
		goto L1
	} else {
		goto L122
	}
L122:
	;
	v483 = v480
	goto L116
L123:
	;
	v486 = F_lappend_oid(m, v452, v483)
	mBase = m.M
	v487 = m.ExcPending
	if v487 != 0 {
		goto L1
	} else {
		goto L124
	}
L124:
	;
	v489 = v456 + int32(1)
	v490 = *(*int32)(unsafe.Add(mBase, uint32(v430)+4))
	if v489 < v490 {
		v449 = v484
		v452 = v486
		v456 = v489
		goto L114
	} else {
		goto L125
	}
L125:
	;
	goto L115
L126:
	;
	v521 = *(*int32)(unsafe.Add(mBase, uint32(v24)+44))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+8)) = v521
	v523 = *(*int64)(unsafe.Add(mBase, uint32(v24)+36))
	*(*int64)(unsafe.Add(mBase, uint32(l0))) = v523
	v525 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v526 = *(*int32)(unsafe.Add(mBase, uint32(v321)+24))
	if v526 != 0 {
		goto L127
	} else {
		goto L128
	}
L127:
	;
	F_CreateComments(m, v525, int32(3079), int32(0), v526)
	mBase = m.M
	v530 = m.ExcPending
	if v530 != 0 {
		goto L1
	} else {
		goto L130
	}
L128:
	;
	goto L129
L129:
	;
	F_execute_extension_script(m, v525, v321, int32(0), v302, v502, v416)
	mBase = m.M
	v533 = m.ExcPending
	if v533 != 0 {
		goto L1
	} else {
		goto L131
	}
L130:
	;
	goto L129
L131:
	;
	F_ApplyExtensionUpdates(m, v525, v29, v302, v317, l2, l4, l6)
	mBase = m.M
	v535 = m.ExcPending
	if v535 != 0 {
		goto L1
	} else {
		goto L132
	}
L132:
	;
	m.G0 = v24 + int32(144)
	return
L133:
	;
	F_errcode(m, int32(50856066))
	mBase = m.M
	v545 = m.ExcPending
	if v545 != 0 {
		goto L1
	} else {
		goto L134
	}
L134:
	;
	F_errmsg(m, int32(_a_F_CreateExtensionInternal_2), int32(0))
	mBase = m.M
	v549 = m.ExcPending
	if v549 != 0 {
		goto L1
	} else {
		goto L135
	}
L135:
	;
	F_errfinish(m, int32(_a_F_CreateExtensionInternal_3), int32(1874), int32(_a_F_CreateExtensionInternal_4))
	mBase = m.M
	v554 = m.ExcPending
	if v554 != 0 {
		goto L1
	} else {
		goto L136
	}
L136:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L137:
	;
	F_errcode(m, int32(50856066))
	mBase = m.M
	v582 = m.ExcPending
	if v582 != 0 {
		goto L1
	} else {
		goto L138
	}
L138:
	;
	v583 = *(*int32)(unsafe.Add(mBase, uint32(v29)))
	*(*int32)(unsafe.Add(mBase, uint32(v24)+20)) = v48
	*(*int32)(unsafe.Add(mBase, uint32(v24)+16)) = v583
	F_errmsg(m, int32(_a_F_CreateExtensionInternal_5), v24+int32(16))
	mBase = m.M
	v590 = m.ExcPending
	if v590 != 0 {
		goto L1
	} else {
		goto L139
	}
L139:
	;
	F_errfinish(m, int32(_a_F_CreateExtensionInternal_3), int32(1912), int32(_a_F_CreateExtensionInternal_4))
	mBase = m.M
	v595 = m.ExcPending
	if v595 != 0 {
		goto L1
	} else {
		goto L140
	}
L140:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L141:
	;
	F_errcode(m, int32(1088))
	mBase = m.M
	v602 = m.ExcPending
	if v602 != 0 {
		goto L1
	} else {
		goto L142
	}
L142:
	;
	v603 = *(*int32)(unsafe.Add(mBase, uint32(v321)))
	v604 = *(*int32)(unsafe.Add(mBase, uint32(v321)+28))
	*(*int32)(unsafe.Add(mBase, uint32(v24)+4)) = v604
	*(*int32)(unsafe.Add(mBase, uint32(v24))) = v603
	F_errmsg(m, int32(_a_F_CreateExtensionInternal_6), v24)
	mBase = m.M
	v609 = m.ExcPending
	if v609 != 0 {
		goto L1
	} else {
		goto L143
	}
L143:
	;
	F_errfinish(m, int32(_a_F_CreateExtensionInternal_3), int32(1947), int32(_a_F_CreateExtensionInternal_4))
	mBase = m.M
	v614 = m.ExcPending
	if v614 != 0 {
		goto L1
	} else {
		goto L144
	}
L144:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L145:
	;
	F_errcode(m, int32(1411))
	mBase = m.M
	v621 = m.ExcPending
	if v621 != 0 {
		goto L1
	} else {
		goto L146
	}
L146:
	;
	F_errmsg(m, int32(_a_F_CreateExtensionInternal_7), int32(0))
	mBase = m.M
	v625 = m.ExcPending
	if v625 != 0 {
		goto L1
	} else {
		goto L147
	}
L147:
	;
	F_errfinish(m, int32(_a_F_CreateExtensionInternal_3), int32(1988), int32(_a_F_CreateExtensionInternal_4))
	mBase = m.M
	v630 = m.ExcPending
	if v630 != 0 {
		goto L1
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
	F_errcode(m, int32(1411))
	mBase = m.M
	v637 = m.ExcPending
	if v637 != 0 {
		goto L1
	} else {
		goto L150
	}
L150:
	;
	F_errmsg(m, int32(_a_F_CreateExtensionInternal_7), int32(0))
	mBase = m.M
	v641 = m.ExcPending
	if v641 != 0 {
		goto L1
	} else {
		goto L151
	}
L151:
	;
	F_errfinish(m, int32(_a_F_CreateExtensionInternal_3), int32(1994), int32(_a_F_CreateExtensionInternal_4))
	mBase = m.M
	v646 = m.ExcPending
	if v646 != 0 {
		goto L1
	} else {
		goto L152
	}
L152:
	;
	base.Wasm_trap_unreachable()
	for {
	}
}
func F_get_extension_name(m *base.Module, l0 int32) int32 {
	var v3 int32
	_ = v3
	var v6 int32
	_ = v6
	v3 = Fn14286(m, l0, int32(28))
	v6 = m.ExcPending
	if v6 != 0 {
		return int32(0)
	} else {
		return v3
	}
}
func F_get_extension_script_filename(m *base.Module, l0 int32, l1 int32, l2 int32) int32 {
	mBase := m.M
	_ = mBase
	var v7 int32
	_ = v7
	var v9 int32
	_ = v9
	var v11 int32
	_ = v11
	var v14 int32
	_ = v14
	var v15 int32
	_ = v15
	var v18 int32
	_ = v18
	var v19 int32
	_ = v19
	var v22 int32
	_ = v22
	var v23 int32
	_ = v23
	var v24 int32
	_ = v24
	var v30 int32
	_ = v30
	var v31 int32
	_ = v31
	var v33 int32
	_ = v33
	var v35 int32
	_ = v35
	var v36 int32
	_ = v36
	var v37 int32
	_ = v37
	var v46 int32
	_ = v46
	var v47 int32
	_ = v47
	var v53 int32
	_ = v53
	var v54 int32
	_ = v54
	var v56 int32
	_ = v56
	v7 = m.G0
	v9 = v7 - int32(48)
	m.G0 = v9
	v11 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	if v11 == int32(0) {
		v14 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
		v15 = F_pstrdup(m, v14)
		mBase = m.M
		v18 = m.ExcPending
		if v18 != 0 {
			return int32(0)
		} else {
			v33 = v15
			v35 = F_palloc(m, int32(1024))
			mBase = m.M
			v36 = m.ExcPending
			if v36 != 0 {
				return int32(0)
			} else {
				v37 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
				if l1 != 0 {
					*(*int32)(unsafe.Add(mBase, uint32(v9)+28)) = l2
					*(*int32)(unsafe.Add(mBase, uint32(v9)+24)) = l1
					*(*int32)(unsafe.Add(mBase, uint32(v9)+20)) = v37
					*(*int32)(unsafe.Add(mBase, uint32(v9)+16)) = v33
					v46 = F_pg_snprintf(m, v35, int32(1024), int32(_a_F_get_extension_script_filename_0), v9+int32(16))
					mBase = m.M
					v47 = m.ExcPending
					if v47 != 0 {
						return int32(0)
					} else {
						F_pfree(m, v33)
						mBase = m.M
						v56 = m.ExcPending
						if v56 != 0 {
							return int32(0)
						} else {
							m.G0 = v9 + int32(48)
							return v35
						}
					}
				} else {
					*(*int32)(unsafe.Add(mBase, uint32(v9)+8)) = l2
					*(*int32)(unsafe.Add(mBase, uint32(v9)+4)) = v37
					*(*int32)(unsafe.Add(mBase, uint32(v9))) = v33
					v53 = F_pg_snprintf(m, v35, int32(1024), int32(_a_F_get_extension_script_filename_1), v9)
					mBase = m.M
					v54 = m.ExcPending
					if v54 != 0 {
						return int32(0)
					} else {
						F_pfree(m, v33)
						mBase = m.M
						v56 = m.ExcPending
						if v56 != 0 {
							return int32(0)
						} else {
							m.G0 = v9 + int32(48)
							return v35
						}
					}
				}
			}
		}
	} else {
		v19 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v11))))
		if v19 == int32(47) {
			v22 = F_pstrdup(m, v11)
			mBase = m.M
			v23 = m.ExcPending
			if v23 != 0 {
				return int32(0)
			} else {
				v33 = v22
				v35 = F_palloc(m, int32(1024))
				mBase = m.M
				v36 = m.ExcPending
				if v36 != 0 {
					return int32(0)
				} else {
					v37 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
					if l1 != 0 {
						*(*int32)(unsafe.Add(mBase, uint32(v9)+28)) = l2
						*(*int32)(unsafe.Add(mBase, uint32(v9)+24)) = l1
						*(*int32)(unsafe.Add(mBase, uint32(v9)+20)) = v37
						*(*int32)(unsafe.Add(mBase, uint32(v9)+16)) = v33
						v46 = F_pg_snprintf(m, v35, int32(1024), int32(_a_F_get_extension_script_filename_0), v9+int32(16))
						mBase = m.M
						v47 = m.ExcPending
						if v47 != 0 {
							return int32(0)
						} else {
							F_pfree(m, v33)
							mBase = m.M
							v56 = m.ExcPending
							if v56 != 0 {
								return int32(0)
							} else {
								m.G0 = v9 + int32(48)
								return v35
							}
						}
					} else {
						*(*int32)(unsafe.Add(mBase, uint32(v9)+8)) = l2
						*(*int32)(unsafe.Add(mBase, uint32(v9)+4)) = v37
						*(*int32)(unsafe.Add(mBase, uint32(v9))) = v33
						v53 = F_pg_snprintf(m, v35, int32(1024), int32(_a_F_get_extension_script_filename_1), v9)
						mBase = m.M
						v54 = m.ExcPending
						if v54 != 0 {
							return int32(0)
						} else {
							F_pfree(m, v33)
							mBase = m.M
							v56 = m.ExcPending
							if v56 != 0 {
								return int32(0)
							} else {
								m.G0 = v9 + int32(48)
								return v35
							}
						}
					}
				}
			}
		} else {
			v24 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
			*(*int32)(unsafe.Add(mBase, uint32(v9)+36)) = v11
			*(*int32)(unsafe.Add(mBase, uint32(v9)+32)) = v24
			v30 = F_psprintf(m, int32(_a_F_get_extension_script_filename_2), v9+int32(32))
			mBase = m.M
			v31 = m.ExcPending
			if v31 != 0 {
				return int32(0)
			} else {
				v33 = v30
				v35 = F_palloc(m, int32(1024))
				mBase = m.M
				v36 = m.ExcPending
				if v36 != 0 {
					return int32(0)
				} else {
					v37 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
					if l1 != 0 {
						*(*int32)(unsafe.Add(mBase, uint32(v9)+28)) = l2
						*(*int32)(unsafe.Add(mBase, uint32(v9)+24)) = l1
						*(*int32)(unsafe.Add(mBase, uint32(v9)+20)) = v37
						*(*int32)(unsafe.Add(mBase, uint32(v9)+16)) = v33
						v46 = F_pg_snprintf(m, v35, int32(1024), int32(_a_F_get_extension_script_filename_0), v9+int32(16))
						mBase = m.M
						v47 = m.ExcPending
						if v47 != 0 {
							return int32(0)
						} else {
							F_pfree(m, v33)
							mBase = m.M
							v56 = m.ExcPending
							if v56 != 0 {
								return int32(0)
							} else {
								m.G0 = v9 + int32(48)
								return v35
							}
						}
					} else {
						*(*int32)(unsafe.Add(mBase, uint32(v9)+8)) = l2
						*(*int32)(unsafe.Add(mBase, uint32(v9)+4)) = v37
						*(*int32)(unsafe.Add(mBase, uint32(v9))) = v33
						v53 = F_pg_snprintf(m, v35, int32(1024), int32(_a_F_get_extension_script_filename_1), v9)
						mBase = m.M
						v54 = m.ExcPending
						if v54 != 0 {
							return int32(0)
						} else {
							F_pfree(m, v33)
							mBase = m.M
							v56 = m.ExcPending
							if v56 != 0 {
								return int32(0)
							} else {
								m.G0 = v9 + int32(48)
								return v35
							}
						}
					}
				}
			}
		}
	}
}
