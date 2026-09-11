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
	var v75 int32
	_ = v75
	var v92 int32
	_ = v92
	var v93 int32
	_ = v93
	var v96 int32
	_ = v96
	var v97 int32
	_ = v97
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
	var v169 int32
	_ = v169
	var v170 int32
	_ = v170
	var v180 int32
	_ = v180
	var v183 int32
	_ = v183
	var v194 int32
	_ = v194
	var v204 int32
	_ = v204
	var v206 int32
	_ = v206
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
	var v239 int32
	_ = v239
	var v243 int32
	_ = v243
	var v244 int32
	_ = v244
	var v247 int32
	_ = v247
	var v248 int32
	_ = v248
	var v251 int32
	_ = v251
	var v258 int32
	_ = v258
	var v259 int32
	_ = v259
	var v268 int32
	_ = v268
	var v269 int32
	_ = v269
	var v271 int32
	_ = v271
	var v272 int32
	_ = v272
	var v293 int32
	_ = v293
	var v295 int32
	_ = v295
	var v297 int32
	_ = v297
	var v301 int32
	_ = v301
	var v317 int32
	_ = v317
	var v320 int32
	_ = v320
	var v321 int32
	_ = v321
	var v322 int64
	_ = v322
	var v324 int64
	_ = v324
	var v326 int64
	_ = v326
	var v328 int64
	_ = v328
	var v330 int64
	_ = v330
	var v332 int64
	_ = v332
	var v335 int32
	_ = v335
	var v338 int32
	_ = v338
	var v340 int32
	_ = v340
	var v341 int32
	_ = v341
	var v342 int32
	_ = v342
	var v347 int32
	_ = v347
	var v348 int32
	_ = v348
	var v352 int32
	_ = v352
	var v353 int32
	_ = v353
	var v356 int32
	_ = v356
	var v357 int32
	_ = v357
	var v360 int32
	_ = v360
	var v367 int32
	_ = v367
	var v368 int32
	_ = v368
	var v371 int32
	_ = v371
	var v374 int32
	_ = v374
	var v375 int32
	_ = v375
	var v377 int32
	_ = v377
	var v378 int32
	_ = v378
	var v384 int32
	_ = v384
	var v387 int32
	_ = v387
	var v390 int32
	_ = v390
	var v392 int32
	_ = v392
	var v393 int32
	_ = v393
	var v397 int32
	_ = v397
	var v398 int32
	_ = v398
	var v401 int32
	_ = v401
	var v402 int32
	_ = v402
	var v403 int32
	_ = v403
	var v404 int32
	_ = v404
	var v408 int32
	_ = v408
	var v410 int32
	_ = v410
	var v411 int32
	_ = v411
	var v414 int32
	_ = v414
	var v419 int32
	_ = v419
	var v421 int32
	_ = v421
	var v425 int32
	_ = v425
	var v427 int32
	_ = v427
	var v430 int32
	_ = v430
	var v441 int32
	_ = v441
	var v445 int32
	_ = v445
	var v446 int32
	_ = v446
	var v456 int32
	_ = v456
	var v460 int32
	_ = v460
	var v461 int32
	_ = v461
	var v462 int32
	_ = v462
	var v463 int32
	_ = v463
	var v464 int32
	_ = v464
	var v468 int32
	_ = v468
	var v469 int32
	_ = v469
	var v471 int32
	_ = v471
	var v473 int32
	_ = v473
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
	var v480 int32
	_ = v480
	var v481 int32
	_ = v481
	var v486 int32
	_ = v486
	var v489 int32
	_ = v489
	var v493 int32
	_ = v493
	var v498 int32
	_ = v498
	var v523 int32
	_ = v523
	var v526 int32
	_ = v526
	var v527 int32
	_ = v527
	var v534 int32
	_ = v534
	var v539 int32
	_ = v539
	var v543 int32
	_ = v543
	var v546 int32
	_ = v546
	var v547 int32
	_ = v547
	var v548 int32
	_ = v548
	var v553 int32
	_ = v553
	var v558 int32
	_ = v558
	var v562 int32
	_ = v562
	var v565 int32
	_ = v565
	var v569 int32
	_ = v569
	var v574 int32
	_ = v574
	var v578 int32
	_ = v578
	var v581 int32
	_ = v581
	var v585 int32
	_ = v585
	var v590 int32
	_ = v590
	var v602 int32
	_ = v602
	var v603 int32
	_ = v603
	var v612 int32
	_ = v612
	var v613 int32
	_ = v613
	var v614 int32
	_ = v614
	var v617 int32
	_ = v617
	var v618 int32
	_ = v618
	var v619 int32
	_ = v619
	var v623 int32
	_ = v623
	var v626 int32
	_ = v626
	var v628 int32
	_ = v628
	v8 = int32(0)
	v22 = m.G0
	v24 = v22 - int32(128)
	m.G0 = v24
	v27 = *(*int32)(unsafe.Add(mBase, _consts[276]))
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
		goto L11
	} else {
		goto L12
	}
L5:
	;
	v612 = *(*int32)(unsafe.Add(mBase, uint32(v320)))
	v613 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v320)+32)))
	v614 = int32(0)
	F_InsertExtensionTuple(m, l0, v612, v27, v411, v613, v301, v614, v614, v602)
	mBase = m.M
	v617 = m.ExcPending
	if v617 != 0 {
		goto L1
	} else {
		goto L144
	}
L6:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v578 = m.ExcPending
	if v578 != 0 {
		goto L1
	} else {
		goto L140
	}
L7:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v562 = m.ExcPending
	if v562 != 0 {
		goto L1
	} else {
		goto L136
	}
L8:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v543 = m.ExcPending
	if v543 != 0 {
		goto L1
	} else {
		goto L132
	}
L9:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v523 = m.ExcPending
	if v523 != 0 {
		goto L1
	} else {
		goto L128
	}
L10:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v486 = m.ExcPending
	if v486 != 0 {
		goto L1
	} else {
		goto L124
	}
L11:
	;
	v45 = *(*int32)(unsafe.Add(mBase, uint32(v29)+16))
	if v45 == int32(0) {
		goto L10
	} else {
		goto L14
	}
L12:
	;
	v48 = l3
	goto L13
L13:
	;
	F_check_valid_version_name(m, v48)
	mBase = m.M
	v50 = m.ExcPending
	if v50 != 0 {
		goto L1
	} else {
		goto L15
	}
L14:
	;
	v48 = v45
	goto L13
L15:
	;
	v52 = F_get_extension_script_filename(m, v29, int32(0), v48)
	mBase = m.M
	v53 = m.ExcPending
	if v53 != 0 {
		goto L1
	} else {
		goto L16
	}
L16:
	;
	v58 = F___fstatat(m, int32(-100), v52, v24+int32(32), int32(0))
	mBase = m.M
	goto L17
L17:
	;
	if v58 != 0 {
		goto L18
	} else {
		goto L19
	}
L18:
	;
	v59 = F_get_ext_ver_list(m, v29)
	mBase = m.M
	v60 = m.ExcPending
	if v60 != 0 {
		goto L1
	} else {
		goto L23
	}
L19:
	;
	v301 = v48
	v317 = v8
	goto L20
L20:
	;
	v320 = F_palloc(m, int32(48))
	mBase = m.M
	v321 = m.ExcPending
	if v321 != 0 {
		goto L1
	} else {
		goto L73
	}
L21:
	;
	v180 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v170)+8)))
	if v180 != 0 {
		goto L42
	} else {
		goto L43
	}
L22:
	;
	v146 = F_palloc(m, int32(20))
	mBase = m.M
	v147 = m.ExcPending
	if v147 != 0 {
		goto L1
	} else {
		goto L38
	}
L23:
	;
	if v59 == int32(0) {
		goto L22
	} else {
		goto L24
	}
L24:
	;
	v63 = *(*int32)(unsafe.Add(mBase, uint32(v59)+4))
	if v63 <= int32(0) {
		goto L22
	} else {
		goto L25
	}
L25:
	;
	v66 = *(*int32)(unsafe.Add(mBase, uint32(v59)+12))
	v75 = int32(0)
	goto L26
L26:
	;
	v92 = *(*int32)(unsafe.Add(mBase, uint32(v66+v75<<(uint(int32(2))%32))))
	v93 = *(*int32)(unsafe.Add(mBase, uint32(v92)))
	v96 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v48))))
	v97 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v93))))
	if v97 == int32(0) {
		v116 = v96
		v117 = v97
		goto L29
	} else {
		goto L30
	}
L27:
	;
	goto L22
L28:
	;
	if v117-v116 == int32(0) {
		v169 = v59
		v170 = v92
		goto L21
	} else {
		goto L36
	}
L29:
	;
	goto L28
L30:
	;
	if v96 != v97 {
		v116 = v96
		v117 = v97
		goto L29
	} else {
		goto L31
	}
L31:
	;
	v101 = v93
	v102 = v48
	goto L32
L32:
	;
	v105 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v102)+1)))
	v106 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v101)+1)))
	if v106 == int32(0) {
		v116 = v105
		v117 = v106
		goto L29
	} else {
		goto L34
	}
L33:
	;
	v116 = v105
	v117 = v106
	goto L29
L34:
	;
	v109 = int32(1)
	if v105 == v106 {
		v101 = v101 + v109
		v102 = v102 + v109
		goto L32
	} else {
		goto L35
	}
L35:
	;
	goto L33
L36:
	;
	v122 = v75 + int32(1)
	if v63 != v122 {
		v75 = v122
		goto L26
	} else {
		goto L37
	}
L37:
	;
	goto L27
L38:
	;
	v148 = F_pstrdup(m, v48)
	mBase = m.M
	v149 = m.ExcPending
	if v149 != 0 {
		goto L1
	} else {
		goto L39
	}
L39:
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
		goto L40
	}
L40:
	;
	v169 = v157
	v170 = v146
	goto L21
L41:
	;
	v297 = *(*int32)(unsafe.Add(mBase, uint32(v293)))
	v301 = v297
	v317 = v295
	goto L20
L42:
	;
	v293 = v170
	v295 = v8
	goto L41
L43:
	;
	goto L44
L44:
	;
	if v169 == int32(0) {
		goto L9
	} else {
		goto L45
	}
L45:
	;
	v183 = *(*int32)(unsafe.Add(mBase, uint32(v169)+4))
	if v183 <= int32(0) {
		goto L9
	} else {
		goto L46
	}
L46:
	;
	v194 = int32(0)
	v204 = v8
	v206 = v8
	goto L47
L47:
	;
	v208 = *(*int32)(unsafe.Add(mBase, uint32(v169)+12))
	v212 = *(*int32)(unsafe.Add(mBase, uint32(v208+v194<<(uint(int32(2))%32))))
	v213 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v212)+8)))
	if v213 != int32(1) {
		v268 = v204
		v269 = v206
		goto L49
	} else {
		goto L50
	}
L48:
	;
	if v268 == int32(0) {
		goto L9
	} else {
		goto L72
	}
L49:
	;
	v271 = v194 + int32(1)
	v272 = *(*int32)(unsafe.Add(mBase, uint32(v169)+4))
	if v271 < v272 {
		v194 = v271
		v204 = v268
		v206 = v269
		goto L47
	} else {
		goto L71
	}
L50:
	;
	v216 = int32(1)
	v218 = F_find_update_path(m, v169, v212, v170, v216, v216)
	mBase = m.M
	v219 = m.ExcPending
	if v219 != 0 {
		goto L1
	} else {
		goto L51
	}
L51:
	;
	if v218 == int32(0) {
		v268 = v204
		v269 = v206
		goto L49
	} else {
		goto L52
	}
L52:
	;
	if v204 == int32(0) {
		goto L53
	} else {
		goto L54
	}
L53:
	;
	v268 = v212
	v269 = v218
	goto L49
L54:
	;
	v224 = *(*int32)(unsafe.Add(mBase, uint32(v218)+4))
	if v206 == int32(0) {
		goto L56
	} else {
		goto L57
	}
L55:
	;
	if v232 != v224 {
		v268 = v204
		v269 = v206
		goto L49
	} else {
		goto L61
	}
L56:
	;
	v227 = int32(0)
	if v227 <= v224 {
		v232 = v227
		goto L55
	} else {
		goto L59
	}
L57:
	;
	goto L58
L58:
	;
	v230 = *(*int32)(unsafe.Add(mBase, uint32(v206)+4))
	if v224 < v230 {
		goto L53
	} else {
		goto L60
	}
L59:
	;
	goto L53
L60:
	;
	v232 = v230
	goto L55
L61:
	;
	v234 = *(*int32)(unsafe.Add(mBase, uint32(v204)))
	v235 = *(*int32)(unsafe.Add(mBase, uint32(v212)))
	v238 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v235))))
	v239 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v234))))
	if v239 == int32(0) {
		v258 = v238
		v259 = v239
		goto L63
	} else {
		goto L64
	}
L62:
	;
	if int32(0) <= v259-v258 {
		v268 = v204
		v269 = v206
		goto L49
	} else {
		goto L70
	}
L63:
	;
	goto L62
L64:
	;
	if v238 != v239 {
		v258 = v238
		v259 = v239
		goto L63
	} else {
		goto L65
	}
L65:
	;
	v243 = v234
	v244 = v235
	goto L66
L66:
	;
	v247 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v244)+1)))
	v248 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v243)+1)))
	if v248 == int32(0) {
		v258 = v247
		v259 = v248
		goto L63
	} else {
		goto L68
	}
L67:
	;
	v258 = v247
	v259 = v248
	goto L63
L68:
	;
	v251 = int32(1)
	if v247 == v248 {
		v243 = v243 + v251
		v244 = v244 + v251
		goto L66
	} else {
		goto L69
	}
L69:
	;
	goto L67
L70:
	;
	goto L53
L71:
	;
	goto L48
L72:
	;
	v293 = v268
	v295 = v269
	goto L41
L73:
	;
	v322 = *(*int64)(unsafe.Add(mBase, uint32(v29)+40))
	*(*int64)(unsafe.Add(mBase, uint32(v320)+40)) = v322
	v324 = *(*int64)(unsafe.Add(mBase, uint32(v29)+32))
	*(*int64)(unsafe.Add(mBase, uint32(v320)+32)) = v324
	v326 = *(*int64)(unsafe.Add(mBase, uint32(v29)+24))
	*(*int64)(unsafe.Add(mBase, uint32(v320)+24)) = v326
	v328 = *(*int64)(unsafe.Add(mBase, uint32(v29)+16))
	*(*int64)(unsafe.Add(mBase, uint32(v320)+16)) = v328
	v330 = *(*int64)(unsafe.Add(mBase, uint32(v29)+8))
	*(*int64)(unsafe.Add(mBase, uint32(v320)+8)) = v330
	v332 = *(*int64)(unsafe.Add(mBase, uint32(v29)))
	*(*int64)(unsafe.Add(mBase, uint32(v320))) = v332
	F_parse_extension_control_file(m, v320, v301)
	mBase = m.M
	v335 = m.ExcPending
	if v335 != 0 {
		goto L1
	} else {
		goto L74
	}
L74:
	;
	if l2 == int32(0) {
		goto L79
	} else {
		goto L80
	}
L75:
	;
	v414 = *(*int32)(unsafe.Add(mBase, _consts[287]))
	goto L106
L76:
	;
	v397 = F_fetch_search_path(m, int32(0))
	mBase = m.M
	v398 = m.ExcPending
	if v398 != 0 {
		goto L1
	} else {
		goto L101
	}
L77:
	;
	if v340 != 0 {
		v410 = l2
		v411 = v340
		goto L75
	} else {
		goto L100
	}
L78:
	;
	v374 = F_get_namespace_oid(m, v371, int32(1))
	mBase = m.M
	v375 = m.ExcPending
	if v375 != 0 {
		goto L1
	} else {
		goto L95
	}
L79:
	;
	v338 = *(*int32)(unsafe.Add(mBase, uint32(v320)+28))
	if v338 != 0 {
		v371 = v338
		goto L78
	} else {
		goto L82
	}
L80:
	;
	goto L81
L81:
	;
	v340 = F_get_namespace_oid(m, l2, int32(0))
	mBase = m.M
	v341 = m.ExcPending
	if v341 != 0 {
		goto L1
	} else {
		goto L83
	}
L82:
	;
	goto L76
L83:
	;
	v342 = *(*int32)(unsafe.Add(mBase, uint32(v320)+28))
	if v342 == int32(0) {
		goto L77
	} else {
		goto L84
	}
L84:
	;
	v347 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l2))))
	v348 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v342))))
	if v348 == int32(0) {
		v367 = v347
		v368 = v348
		goto L86
	} else {
		goto L87
	}
L85:
	;
	if l4 != 0 {
		v371 = v342
		goto L78
	} else {
		goto L93
	}
L86:
	;
	goto L85
L87:
	;
	if v347 != v348 {
		v367 = v347
		v368 = v348
		goto L86
	} else {
		goto L88
	}
L88:
	;
	v352 = v342
	v353 = l2
	goto L89
L89:
	;
	v356 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v353)+1)))
	v357 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v352)+1)))
	if v357 == int32(0) {
		v367 = v356
		v368 = v357
		goto L86
	} else {
		goto L91
	}
L90:
	;
	v367 = v356
	v368 = v357
	goto L86
L91:
	;
	v360 = int32(1)
	if v356 == v357 {
		v352 = v352 + v360
		v353 = v353 + v360
		goto L89
	} else {
		goto L92
	}
L92:
	;
	goto L90
L93:
	;
	if v368-v367 != 0 {
		goto L8
	} else {
		goto L94
	}
L94:
	;
	v371 = v342
	goto L78
L95:
	;
	if v374 != 0 {
		v410 = v371
		v411 = v374
		goto L75
	} else {
		goto L96
	}
L96:
	;
	v377 = F_palloc0(m, int32(20))
	mBase = m.M
	v378 = m.ExcPending
	if v378 != 0 {
		goto L1
	} else {
		goto L97
	}
L97:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v377)+8)) = int64(0)
	*(*int32)(unsafe.Add(mBase, uint32(v377)+4)) = v371
	*(*int32)(unsafe.Add(mBase, uint32(v377))) = int32(145)
	v384 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v377)+16)) = uint8(v384)
	v387 = int32(-1)
	F_CreateSchemaCommand(m, v377, int32(627333), v387, v387)
	mBase = m.M
	v390 = m.ExcPending
	if v390 != 0 {
		goto L1
	} else {
		goto L98
	}
L98:
	;
	v392 = F_get_namespace_oid(m, v371, int32(0))
	mBase = m.M
	v393 = m.ExcPending
	if v393 != 0 {
		goto L1
	} else {
		goto L99
	}
L99:
	;
	v410 = v371
	v411 = v392
	goto L75
L100:
	;
	goto L76
L101:
	;
	if v397 == int32(0) {
		goto L7
	} else {
		goto L102
	}
L102:
	;
	v401 = *(*int32)(unsafe.Add(mBase, uint32(v397)+12))
	v402 = *(*int32)(unsafe.Add(mBase, uint32(v401)))
	v403 = F_get_namespace_name(m, v402)
	mBase = m.M
	v404 = m.ExcPending
	if v404 != 0 {
		goto L1
	} else {
		goto L103
	}
L103:
	;
	if v403 == int32(0) {
		goto L6
	} else {
		goto L104
	}
L104:
	;
	F_list_free(m, v397)
	mBase = m.M
	v408 = m.ExcPending
	if v408 != 0 {
		goto L1
	} else {
		goto L105
	}
L105:
	;
	v410 = v403
	v411 = v402
	goto L75
L106:
	;
	if base.B2i32(v414 != int32(0))&base.B2i32(v411 == v414) != 0 {
		goto L107
	} else {
		goto L108
	}
L107:
	;
	v419 = int32(4320404)
	v421 = *(*int32)(unsafe.Add(mBase, _consts[386]))
	*(*int32)(unsafe.Add(mBase, _consts[386])) = v421 | int32(1)
	goto L109
L108:
	;
	goto L109
L109:
	;
	v425 = int32(0)
	v427 = *(*int32)(unsafe.Add(mBase, uint32(v320)+40))
	if v427 == v425 {
		v602 = v425
		v603 = v425
		goto L5
	} else {
		goto L110
	}
L110:
	;
	v430 = *(*int32)(unsafe.Add(mBase, uint32(v427)+4))
	if v430 <= int32(0) {
		v602 = v425
		v603 = v425
		goto L5
	} else {
		goto L111
	}
L111:
	;
	v441 = int32(0)
	v445 = v425
	v446 = v425
	goto L112
L112:
	;
	v456 = *(*int32)(unsafe.Add(mBase, uint32(v427)+12))
	v460 = *(*int32)(unsafe.Add(mBase, uint32(v456+v441<<(uint(int32(2))%32))))
	v461 = F_get_required_extension(m, v460, l1, l2, l4, l5, l6)
	mBase = m.M
	v462 = m.ExcPending
	if v462 != 0 {
		goto L1
	} else {
		goto L115
	}
L113:
	;
	v602 = v475
	v603 = v477
	goto L5
L114:
	;
	v475 = F_lappend_oid(m, v445, v461)
	mBase = m.M
	v476 = m.ExcPending
	if v476 != 0 {
		goto L1
	} else {
		goto L121
	}
L115:
	;
	v463 = F_SearchSysCache1(m, int32(28), v461)
	mBase = m.M
	v464 = m.ExcPending
	if v464 != 0 {
		goto L1
	} else {
		goto L116
	}
L116:
	;
	if v463 == int32(0) {
		goto L117
	} else {
		goto L118
	}
L117:
	;
	v474 = int32(0)
	goto L114
L118:
	;
	goto L119
L119:
	;
	v468 = *(*int32)(unsafe.Add(mBase, uint32(v463)+16))
	v469 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v468)+22)))
	v471 = *(*int32)(unsafe.Add(mBase, uint32(v468+v469)+72))
	F_ReleaseCatCache(m, v463)
	mBase = m.M
	v473 = m.ExcPending
	if v473 != 0 {
		goto L1
	} else {
		goto L120
	}
L120:
	;
	v474 = v471
	goto L114
L121:
	;
	v477 = F_lappend_oid(m, v446, v474)
	mBase = m.M
	v478 = m.ExcPending
	if v478 != 0 {
		goto L1
	} else {
		goto L122
	}
L122:
	;
	v480 = v441 + int32(1)
	v481 = *(*int32)(unsafe.Add(mBase, uint32(v427)+4))
	if v480 < v481 {
		v441 = v480
		v445 = v475
		v446 = v477
		goto L112
	} else {
		goto L123
	}
L123:
	;
	goto L113
L124:
	;
	F_errcode(m, int32(50856066))
	mBase = m.M
	v489 = m.ExcPending
	if v489 != 0 {
		goto L1
	} else {
		goto L125
	}
L125:
	;
	F_errmsg(m, int32(430605), int32(0))
	mBase = m.M
	v493 = m.ExcPending
	if v493 != 0 {
		goto L1
	} else {
		goto L126
	}
L126:
	;
	F_errfinish(m, int32(465954), int32(1822), int32(293454))
	mBase = m.M
	v498 = m.ExcPending
	if v498 != 0 {
		goto L1
	} else {
		goto L127
	}
L127:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L128:
	;
	F_errcode(m, int32(50856066))
	mBase = m.M
	v526 = m.ExcPending
	if v526 != 0 {
		goto L1
	} else {
		goto L129
	}
L129:
	;
	v527 = *(*int32)(unsafe.Add(mBase, uint32(v29)))
	*(*int32)(unsafe.Add(mBase, uint32(v24)+20)) = v48
	*(*int32)(unsafe.Add(mBase, uint32(v24)+16)) = v527
	F_errmsg(m, int32(659074), v24+int32(16))
	mBase = m.M
	v534 = m.ExcPending
	if v534 != 0 {
		goto L1
	} else {
		goto L130
	}
L130:
	;
	F_errfinish(m, int32(465954), int32(1860), int32(293454))
	mBase = m.M
	v539 = m.ExcPending
	if v539 != 0 {
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
	F_errcode(m, int32(1088))
	mBase = m.M
	v546 = m.ExcPending
	if v546 != 0 {
		goto L1
	} else {
		goto L133
	}
L133:
	;
	v547 = *(*int32)(unsafe.Add(mBase, uint32(v320)))
	v548 = *(*int32)(unsafe.Add(mBase, uint32(v320)+28))
	*(*int32)(unsafe.Add(mBase, uint32(v24)+4)) = v548
	*(*int32)(unsafe.Add(mBase, uint32(v24))) = v547
	F_errmsg(m, int32(673814), v24)
	mBase = m.M
	v553 = m.ExcPending
	if v553 != 0 {
		goto L1
	} else {
		goto L134
	}
L134:
	;
	F_errfinish(m, int32(465954), int32(1895), int32(293454))
	mBase = m.M
	v558 = m.ExcPending
	if v558 != 0 {
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
	F_errcode(m, int32(1411))
	mBase = m.M
	v565 = m.ExcPending
	if v565 != 0 {
		goto L1
	} else {
		goto L137
	}
L137:
	;
	F_errmsg(m, int32(263008), int32(0))
	mBase = m.M
	v569 = m.ExcPending
	if v569 != 0 {
		goto L1
	} else {
		goto L138
	}
L138:
	;
	F_errfinish(m, int32(465954), int32(1933), int32(293454))
	mBase = m.M
	v574 = m.ExcPending
	if v574 != 0 {
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
	v581 = m.ExcPending
	if v581 != 0 {
		goto L1
	} else {
		goto L141
	}
L141:
	;
	F_errmsg(m, int32(263008), int32(0))
	mBase = m.M
	v585 = m.ExcPending
	if v585 != 0 {
		goto L1
	} else {
		goto L142
	}
L142:
	;
	F_errfinish(m, int32(465954), int32(1939), int32(293454))
	mBase = m.M
	v590 = m.ExcPending
	if v590 != 0 {
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
	v618 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v619 = *(*int32)(unsafe.Add(mBase, uint32(v320)+24))
	if v619 != 0 {
		goto L145
	} else {
		goto L146
	}
L145:
	;
	F_CreateComments(m, v618, int32(3079), int32(0), v619)
	mBase = m.M
	v623 = m.ExcPending
	if v623 != 0 {
		goto L1
	} else {
		goto L148
	}
L146:
	;
	goto L147
L147:
	;
	F_execute_extension_script(m, v618, v320, int32(0), v301, v603, v410)
	mBase = m.M
	v626 = m.ExcPending
	if v626 != 0 {
		goto L1
	} else {
		goto L149
	}
L148:
	;
	goto L147
L149:
	;
	F_ApplyExtensionUpdates(m, v618, v29, v301, v317, l2, l4, l6)
	mBase = m.M
	v628 = m.ExcPending
	if v628 != 0 {
		goto L1
	} else {
		goto L150
	}
L150:
	;
	m.G0 = v24 + int32(128)
	return
}
func F_get_extension_name(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v4 int32
	_ = v4
	var v7 int32
	_ = v7
	var v12 int32
	_ = v12
	var v13 int32
	_ = v13
	var v17 int32
	_ = v17
	var v18 int32
	_ = v18
	var v20 int32
	_ = v20
	v4 = F_SearchSysCache1(m, int32(28), l0)
	mBase = m.M
	v7 = m.ExcPending
	if v7 != 0 {
		return int32(0)
	} else {
		if v4 == int32(0) {
			return int32(0)
		} else {
			v12 = *(*int32)(unsafe.Add(mBase, uint32(v4)+16))
			v13 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v12)+22)))
			v17 = F_pstrdup(m, v12+v13+int32(4))
			mBase = m.M
			v18 = m.ExcPending
			if v18 != 0 {
				return int32(0)
			} else {
				F_ReleaseCatCache(m, v4)
				mBase = m.M
				v20 = m.ExcPending
				if v20 != 0 {
					return int32(0)
				} else {
					return v17
				}
			}
		}
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
					v46 = F_pg_snprintf(m, v35, int32(1024), int32(282596), v9+int32(16))
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
					v53 = F_pg_snprintf(m, v35, int32(1024), int32(282582), v9)
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
						v46 = F_pg_snprintf(m, v35, int32(1024), int32(282596), v9+int32(16))
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
						v53 = F_pg_snprintf(m, v35, int32(1024), int32(282582), v9)
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
			v30 = F_psprintf(m, int32(165202), v9+int32(32))
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
						v46 = F_pg_snprintf(m, v35, int32(1024), int32(282596), v9+int32(16))
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
						v53 = F_pg_snprintf(m, v35, int32(1024), int32(282582), v9)
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
