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
	var v170 int32
	_ = v170
	var v180 int32
	_ = v180
	var v183 int32
	_ = v183
	var v197 int32
	_ = v197
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
	var v266 int32
	_ = v266
	var v270 int32
	_ = v270
	var v272 int32
	_ = v272
	var v273 int32
	_ = v273
	var v287 int32
	_ = v287
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
	var v372 int32
	_ = v372
	var v375 int32
	_ = v375
	var v376 int32
	_ = v376
	var v378 int32
	_ = v378
	var v379 int32
	_ = v379
	var v385 int32
	_ = v385
	var v388 int32
	_ = v388
	var v391 int32
	_ = v391
	var v393 int32
	_ = v393
	var v394 int32
	_ = v394
	var v398 int32
	_ = v398
	var v399 int32
	_ = v399
	var v402 int32
	_ = v402
	var v403 int32
	_ = v403
	var v404 int32
	_ = v404
	var v405 int32
	_ = v405
	var v409 int32
	_ = v409
	var v411 int32
	_ = v411
	var v412 int32
	_ = v412
	var v415 int32
	_ = v415
	var v420 int32
	_ = v420
	var v422 int32
	_ = v422
	var v426 int32
	_ = v426
	var v428 int32
	_ = v428
	var v431 int32
	_ = v431
	var v442 int32
	_ = v442
	var v446 int32
	_ = v446
	var v449 int32
	_ = v449
	var v457 int32
	_ = v457
	var v461 int32
	_ = v461
	var v462 int32
	_ = v462
	var v463 int32
	_ = v463
	var v464 int32
	_ = v464
	var v465 int32
	_ = v465
	var v469 int32
	_ = v469
	var v470 int32
	_ = v470
	var v472 int32
	_ = v472
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
	var v481 int32
	_ = v481
	var v482 int32
	_ = v482
	var v491 int32
	_ = v491
	var v495 int32
	_ = v495
	var v507 int32
	_ = v507
	var v508 int32
	_ = v508
	var v509 int32
	_ = v509
	var v512 int32
	_ = v512
	var v513 int32
	_ = v513
	var v515 int64
	_ = v515
	var v517 int32
	_ = v517
	var v518 int32
	_ = v518
	var v522 int32
	_ = v522
	var v525 int32
	_ = v525
	var v527 int32
	_ = v527
	var v534 int32
	_ = v534
	var v537 int32
	_ = v537
	var v541 int32
	_ = v541
	var v546 int32
	_ = v546
	var v571 int32
	_ = v571
	var v574 int32
	_ = v574
	var v575 int32
	_ = v575
	var v582 int32
	_ = v582
	var v587 int32
	_ = v587
	var v591 int32
	_ = v591
	var v594 int32
	_ = v594
	var v595 int32
	_ = v595
	var v596 int32
	_ = v596
	var v601 int32
	_ = v601
	var v606 int32
	_ = v606
	var v610 int32
	_ = v610
	var v613 int32
	_ = v613
	var v617 int32
	_ = v617
	var v622 int32
	_ = v622
	var v626 int32
	_ = v626
	var v629 int32
	_ = v629
	var v633 int32
	_ = v633
	var v638 int32
	_ = v638
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
	v626 = m.ExcPending
	if v626 != 0 {
		goto L1
	} else {
		goto L144
	}
L6:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v610 = m.ExcPending
	if v610 != 0 {
		goto L1
	} else {
		goto L140
	}
L7:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v591 = m.ExcPending
	if v591 != 0 {
		goto L1
	} else {
		goto L136
	}
L8:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v571 = m.ExcPending
	if v571 != 0 {
		goto L1
	} else {
		goto L132
	}
L9:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v534 = m.ExcPending
	if v534 != 0 {
		goto L1
	} else {
		goto L128
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
		v170 = v59
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
	v170 = v157
	goto L20
L39:
	;
	v298 = *(*int32)(unsafe.Add(mBase, uint32(v287)))
	v302 = v298
	v317 = v295
	goto L19
L40:
	;
	v287 = v166
	v295 = v8
	goto L39
L41:
	;
	goto L42
L42:
	;
	if v170 == int32(0) {
		goto L8
	} else {
		goto L43
	}
L43:
	;
	v183 = *(*int32)(unsafe.Add(mBase, uint32(v170)+4))
	if v183 <= int32(0) {
		goto L8
	} else {
		goto L44
	}
L44:
	;
	v197 = v8
	v201 = int32(0)
	v205 = v8
	goto L45
L45:
	;
	v208 = *(*int32)(unsafe.Add(mBase, uint32(v170)+12))
	v212 = *(*int32)(unsafe.Add(mBase, uint32(v208+v201<<(uint(int32(2))%32))))
	v213 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v212)+8)))
	if v213 != int32(1) {
		v266 = v197
		v270 = v205
		goto L47
	} else {
		goto L48
	}
L46:
	;
	if v266 == int32(0) {
		goto L8
	} else {
		goto L69
	}
L47:
	;
	v272 = v201 + int32(1)
	v273 = *(*int32)(unsafe.Add(mBase, uint32(v170)+4))
	if v272 < v273 {
		v197 = v266
		v201 = v272
		v205 = v270
		goto L45
	} else {
		goto L68
	}
L48:
	;
	v216 = int32(1)
	v218 = F_find_update_path(m, v170, v212, v166, v216, v216)
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
		v266 = v197
		v270 = v205
		goto L47
	} else {
		goto L50
	}
L50:
	;
	if v197 == int32(0) {
		goto L51
	} else {
		goto L52
	}
L51:
	;
	v266 = v212
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
		v266 = v197
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
	v234 = *(*int32)(unsafe.Add(mBase, uint32(v197)))
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
		v266 = v197
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
	v287 = v266
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
	v415 = *(*int32)(unsafe.Add(mBase, _c_F_CreateExtensionInternal[1]))
	goto L102
L73:
	;
	v398 = F_fetch_search_path(m, int32(0))
	mBase = m.M
	v399 = m.ExcPending
	if v399 != 0 {
		goto L1
	} else {
		goto L97
	}
L74:
	;
	if v341 != 0 {
		v411 = l2
		v412 = v341
		goto L72
	} else {
		goto L96
	}
L75:
	;
	v375 = F_get_namespace_oid(m, v372, int32(1))
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
		v372 = v339
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
		v372 = v343
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
	v372 = v343
	goto L75
L91:
	;
	if v375 != 0 {
		v411 = v372
		v412 = v375
		goto L72
	} else {
		goto L92
	}
L92:
	;
	v378 = F_palloc0(m, int32(20))
	mBase = m.M
	v379 = m.ExcPending
	if v379 != 0 {
		goto L1
	} else {
		goto L93
	}
L93:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v378)+8)) = int64(0)
	*(*int32)(unsafe.Add(mBase, uint32(v378)+4)) = v372
	*(*int32)(unsafe.Add(mBase, uint32(v378))) = int32(145)
	v385 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v378)+16)) = uint8(v385)
	v388 = int32(-1)
	F_CreateSchemaCommand(m, v378, int32(_a_F_CreateExtensionInternal_0), v388, v388)
	mBase = m.M
	v391 = m.ExcPending
	if v391 != 0 {
		goto L1
	} else {
		goto L94
	}
L94:
	;
	v393 = F_get_namespace_oid(m, v372, int32(0))
	mBase = m.M
	v394 = m.ExcPending
	if v394 != 0 {
		goto L1
	} else {
		goto L95
	}
L95:
	;
	v411 = v372
	v412 = v393
	goto L72
L96:
	;
	goto L73
L97:
	;
	if v398 == int32(0) {
		goto L6
	} else {
		goto L98
	}
L98:
	;
	v402 = *(*int32)(unsafe.Add(mBase, uint32(v398)+12))
	v403 = *(*int32)(unsafe.Add(mBase, uint32(v402)))
	v404 = F_get_namespace_name(m, v403)
	mBase = m.M
	v405 = m.ExcPending
	if v405 != 0 {
		goto L1
	} else {
		goto L99
	}
L99:
	;
	if v404 == int32(0) {
		goto L5
	} else {
		goto L100
	}
L100:
	;
	F_list_free(m, v398)
	mBase = m.M
	v409 = m.ExcPending
	if v409 != 0 {
		goto L1
	} else {
		goto L101
	}
L101:
	;
	v411 = v404
	v412 = v403
	goto L72
L102:
	;
	if base.B2i32(v415 != int32(0))&base.B2i32(v412 == v415) != 0 {
		goto L103
	} else {
		goto L104
	}
L103:
	;
	v420 = int32(_a_F_CreateExtensionInternal_1)
	v422 = *(*int32)(unsafe.Add(mBase, _c_F_CreateExtensionInternal[2]))
	*(*int32)(unsafe.Add(mBase, _c_F_CreateExtensionInternal[2])) = v422 | int32(1)
	goto L105
L104:
	;
	goto L105
L105:
	;
	v426 = int32(0)
	v428 = *(*int32)(unsafe.Add(mBase, uint32(v321)+40))
	if v428 == v426 {
		v491 = v426
		v495 = v426
		goto L106
	} else {
		goto L107
	}
L106:
	;
	v507 = *(*int32)(unsafe.Add(mBase, uint32(v321)))
	v508 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v321)+32)))
	v509 = int32(0)
	F_InsertExtensionTuple(m, v24+int32(36), v507, v27, v412, v508, v302, v509, v509, v491)
	mBase = m.M
	v512 = m.ExcPending
	if v512 != 0 {
		goto L1
	} else {
		goto L121
	}
L107:
	;
	v431 = *(*int32)(unsafe.Add(mBase, uint32(v428)+4))
	if v431 <= int32(0) {
		v491 = v426
		v495 = v426
		goto L106
	} else {
		goto L108
	}
L108:
	;
	v442 = v426
	v446 = v426
	v449 = int32(0)
	goto L109
L109:
	;
	v457 = *(*int32)(unsafe.Add(mBase, uint32(v428)+12))
	v461 = *(*int32)(unsafe.Add(mBase, uint32(v457+v449<<(uint(int32(2))%32))))
	v462 = F_get_required_extension(m, v461, l1, l2, l4, l5, l6)
	mBase = m.M
	v463 = m.ExcPending
	if v463 != 0 {
		goto L1
	} else {
		goto L112
	}
L110:
	;
	v491 = v476
	v495 = v478
	goto L106
L111:
	;
	v476 = F_lappend_oid(m, v442, v462)
	mBase = m.M
	v477 = m.ExcPending
	if v477 != 0 {
		goto L1
	} else {
		goto L118
	}
L112:
	;
	v464 = F_SearchSysCache1(m, int32(28), v462)
	mBase = m.M
	v465 = m.ExcPending
	if v465 != 0 {
		goto L1
	} else {
		goto L113
	}
L113:
	;
	if v464 == int32(0) {
		goto L114
	} else {
		goto L115
	}
L114:
	;
	v475 = int32(0)
	goto L111
L115:
	;
	goto L116
L116:
	;
	v469 = *(*int32)(unsafe.Add(mBase, uint32(v464)+16))
	v470 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v469)+22)))
	v472 = *(*int32)(unsafe.Add(mBase, uint32(v469+v470)+72))
	F_ReleaseCatCache(m, v464)
	mBase = m.M
	v474 = m.ExcPending
	if v474 != 0 {
		goto L1
	} else {
		goto L117
	}
L117:
	;
	v475 = v472
	goto L111
L118:
	;
	v478 = F_lappend_oid(m, v446, v475)
	mBase = m.M
	v479 = m.ExcPending
	if v479 != 0 {
		goto L1
	} else {
		goto L119
	}
L119:
	;
	v481 = v449 + int32(1)
	v482 = *(*int32)(unsafe.Add(mBase, uint32(v428)+4))
	if v481 < v482 {
		v442 = v476
		v446 = v478
		v449 = v481
		goto L109
	} else {
		goto L120
	}
L120:
	;
	goto L110
L121:
	;
	v513 = *(*int32)(unsafe.Add(mBase, uint32(v24)+44))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+8)) = v513
	v515 = *(*int64)(unsafe.Add(mBase, uint32(v24)+36))
	*(*int64)(unsafe.Add(mBase, uint32(l0))) = v515
	v517 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v518 = *(*int32)(unsafe.Add(mBase, uint32(v321)+24))
	if v518 != 0 {
		goto L122
	} else {
		goto L123
	}
L122:
	;
	F_CreateComments(m, v517, int32(3079), int32(0), v518)
	mBase = m.M
	v522 = m.ExcPending
	if v522 != 0 {
		goto L1
	} else {
		goto L125
	}
L123:
	;
	goto L124
L124:
	;
	F_execute_extension_script(m, v517, v321, int32(0), v302, v495, v411)
	mBase = m.M
	v525 = m.ExcPending
	if v525 != 0 {
		goto L1
	} else {
		goto L126
	}
L125:
	;
	goto L124
L126:
	;
	F_ApplyExtensionUpdates(m, v517, v29, v302, v317, l2, l4, l6)
	mBase = m.M
	v527 = m.ExcPending
	if v527 != 0 {
		goto L1
	} else {
		goto L127
	}
L127:
	;
	m.G0 = v24 + int32(144)
	return
L128:
	;
	F_errcode(m, int32(50856066))
	mBase = m.M
	v537 = m.ExcPending
	if v537 != 0 {
		goto L1
	} else {
		goto L129
	}
L129:
	;
	F_errmsg(m, int32(_a_F_CreateExtensionInternal_2), int32(0))
	mBase = m.M
	v541 = m.ExcPending
	if v541 != 0 {
		goto L1
	} else {
		goto L130
	}
L130:
	;
	F_errfinish(m, int32(_a_F_CreateExtensionInternal_3), int32(1822), int32(_a_F_CreateExtensionInternal_4))
	mBase = m.M
	v546 = m.ExcPending
	if v546 != 0 {
		goto L1
	} else {
		goto L131
	}
L131:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L132:
	;
	F_errcode(m, int32(50856066))
	mBase = m.M
	v574 = m.ExcPending
	if v574 != 0 {
		goto L1
	} else {
		goto L133
	}
L133:
	;
	v575 = *(*int32)(unsafe.Add(mBase, uint32(v29)))
	*(*int32)(unsafe.Add(mBase, uint32(v24)+20)) = v48
	*(*int32)(unsafe.Add(mBase, uint32(v24)+16)) = v575
	F_errmsg(m, int32(_a_F_CreateExtensionInternal_5), v24+int32(16))
	mBase = m.M
	v582 = m.ExcPending
	if v582 != 0 {
		goto L1
	} else {
		goto L134
	}
L134:
	;
	F_errfinish(m, int32(_a_F_CreateExtensionInternal_3), int32(1860), int32(_a_F_CreateExtensionInternal_4))
	mBase = m.M
	v587 = m.ExcPending
	if v587 != 0 {
		goto L1
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
	F_errcode(m, int32(1088))
	mBase = m.M
	v594 = m.ExcPending
	if v594 != 0 {
		goto L1
	} else {
		goto L137
	}
L137:
	;
	v595 = *(*int32)(unsafe.Add(mBase, uint32(v321)))
	v596 = *(*int32)(unsafe.Add(mBase, uint32(v321)+28))
	*(*int32)(unsafe.Add(mBase, uint32(v24)+4)) = v596
	*(*int32)(unsafe.Add(mBase, uint32(v24))) = v595
	F_errmsg(m, int32(_a_F_CreateExtensionInternal_6), v24)
	mBase = m.M
	v601 = m.ExcPending
	if v601 != 0 {
		goto L1
	} else {
		goto L138
	}
L138:
	;
	F_errfinish(m, int32(_a_F_CreateExtensionInternal_3), int32(1895), int32(_a_F_CreateExtensionInternal_4))
	mBase = m.M
	v606 = m.ExcPending
	if v606 != 0 {
		goto L1
	} else {
		goto L139
	}
L139:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L140:
	;
	F_errcode(m, int32(1411))
	mBase = m.M
	v613 = m.ExcPending
	if v613 != 0 {
		goto L1
	} else {
		goto L141
	}
L141:
	;
	F_errmsg(m, int32(_a_F_CreateExtensionInternal_7), int32(0))
	mBase = m.M
	v617 = m.ExcPending
	if v617 != 0 {
		goto L1
	} else {
		goto L142
	}
L142:
	;
	F_errfinish(m, int32(_a_F_CreateExtensionInternal_3), int32(1933), int32(_a_F_CreateExtensionInternal_4))
	mBase = m.M
	v622 = m.ExcPending
	if v622 != 0 {
		goto L1
	} else {
		goto L143
	}
L143:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L144:
	;
	F_errcode(m, int32(1411))
	mBase = m.M
	v629 = m.ExcPending
	if v629 != 0 {
		goto L1
	} else {
		goto L145
	}
L145:
	;
	F_errmsg(m, int32(_a_F_CreateExtensionInternal_7), int32(0))
	mBase = m.M
	v633 = m.ExcPending
	if v633 != 0 {
		goto L1
	} else {
		goto L146
	}
L146:
	;
	F_errfinish(m, int32(_a_F_CreateExtensionInternal_3), int32(1939), int32(_a_F_CreateExtensionInternal_4))
	mBase = m.M
	v638 = m.ExcPending
	if v638 != 0 {
		goto L1
	} else {
		goto L147
	}
L147:
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
	v3 = Fn13897(m, l0, int32(28))
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
