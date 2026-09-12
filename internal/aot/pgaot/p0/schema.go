package p0

import (
	base "github.com/shibukawa/pgmem/internal/aot/pgaot/base"
	"unsafe"
)

func F_CreateSchemaCommand(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) {
	mBase := m.M
	_ = mBase
	var v5 int32
	_ = v5
	var v20 int32
	_ = v20
	var v22 int32
	_ = v22
	var v24 int32
	_ = v24
	var v26 int32
	_ = v26
	var v32 int32
	_ = v32
	var v35 int32
	_ = v35
	var v37 int32
	_ = v37
	var v39 int32
	_ = v39
	var v40 int32
	_ = v40
	var v41 int32
	_ = v41
	var v42 int32
	_ = v42
	var v46 int32
	_ = v46
	var v47 int32
	_ = v47
	var v50 int32
	_ = v50
	var v51 int32
	_ = v51
	var v55 int32
	_ = v55
	var v56 int32
	_ = v56
	var v58 int32
	_ = v58
	var v59 int32
	_ = v59
	var v63 int32
	_ = v63
	var v64 int32
	_ = v64
	var v66 int32
	_ = v66
	var v67 int32
	_ = v67
	var v70 int32
	_ = v70
	var v71 int32
	_ = v71
	var v72 int32
	_ = v72
	var v74 int32
	_ = v74
	var v75 int32
	_ = v75
	var v77 int32
	_ = v77
	var v79 int32
	_ = v79
	var v82 int32
	_ = v82
	var v83 int32
	_ = v83
	var v86 int32
	_ = v86
	var v89 int32
	_ = v89
	var v92 int32
	_ = v92
	var v93 int32
	_ = v93
	var v97 int32
	_ = v97
	var v98 int32
	_ = v98
	var v109 int32
	_ = v109
	var v112 int32
	_ = v112
	var v113 int32
	_ = v113
	var v118 int32
	_ = v118
	var v124 int32
	_ = v124
	var v129 int32
	_ = v129
	var v131 int32
	_ = v131
	var v133 int32
	_ = v133
	var v141 int32
	_ = v141
	var v142 int32
	_ = v142
	var v144 int32
	_ = v144
	var v146 int32
	_ = v146
	var v148 int32
	_ = v148
	var v150 int32
	_ = v150
	var v155 int32
	_ = v155
	var v158 int32
	_ = v158
	var v159 int32
	_ = v159
	var v161 int32
	_ = v161
	var v172 int32
	_ = v172
	var v183 int32
	_ = v183
	var v193 int32
	_ = v193
	var v201 int32
	_ = v201
	var v204 int32
	_ = v204
	var v210 int32
	_ = v210
	var v211 int32
	_ = v211
	var v216 int32
	_ = v216
	var v222 int64
	_ = v222
	var v224 int64
	_ = v224
	var v231 int32
	_ = v231
	var v232 int32
	_ = v232
	var v233 int32
	_ = v233
	var v234 int32
	_ = v234
	var v236 int32
	_ = v236
	var v240 int32
	_ = v240
	var v250 int32
	_ = v250
	var v254 int32
	_ = v254
	var v255 int32
	_ = v255
	var v256 int32
	_ = v256
	var v257 int32
	_ = v257
	var v258 int32
	_ = v258
	var v261 int32
	_ = v261
	var v262 int32
	_ = v262
	var v266 int32
	_ = v266
	var v267 int32
	_ = v267
	var v277 int32
	_ = v277
	var v278 int32
	_ = v278
	var v282 int32
	_ = v282
	var v287 int32
	_ = v287
	var v288 int32
	_ = v288
	var v289 int32
	_ = v289
	var v293 int32
	_ = v293
	var v294 int32
	_ = v294
	var v297 int32
	_ = v297
	var v298 int32
	_ = v298
	var v302 int32
	_ = v302
	var v303 int32
	_ = v303
	var v306 int32
	_ = v306
	var v307 int32
	_ = v307
	var v310 int32
	_ = v310
	var v317 int32
	_ = v317
	var v318 int32
	_ = v318
	var v320 int32
	_ = v320
	var v321 int32
	_ = v321
	var v322 int32
	_ = v322
	var v323 int32
	_ = v323
	var v327 int32
	_ = v327
	var v328 int32
	_ = v328
	var v331 int32
	_ = v331
	var v332 int32
	_ = v332
	var v336 int32
	_ = v336
	var v337 int32
	_ = v337
	var v340 int32
	_ = v340
	var v341 int32
	_ = v341
	var v344 int32
	_ = v344
	var v351 int32
	_ = v351
	var v352 int32
	_ = v352
	var v354 int32
	_ = v354
	var v355 int32
	_ = v355
	var v356 int32
	_ = v356
	var v357 int32
	_ = v357
	var v361 int32
	_ = v361
	var v362 int32
	_ = v362
	var v365 int32
	_ = v365
	var v366 int32
	_ = v366
	var v370 int32
	_ = v370
	var v371 int32
	_ = v371
	var v374 int32
	_ = v374
	var v375 int32
	_ = v375
	var v378 int32
	_ = v378
	var v385 int32
	_ = v385
	var v386 int32
	_ = v386
	var v388 int32
	_ = v388
	var v389 int32
	_ = v389
	var v390 int32
	_ = v390
	var v391 int32
	_ = v391
	var v395 int32
	_ = v395
	var v396 int32
	_ = v396
	var v399 int32
	_ = v399
	var v400 int32
	_ = v400
	var v404 int32
	_ = v404
	var v405 int32
	_ = v405
	var v408 int32
	_ = v408
	var v409 int32
	_ = v409
	var v412 int32
	_ = v412
	var v419 int32
	_ = v419
	var v420 int32
	_ = v420
	var v422 int32
	_ = v422
	var v423 int32
	_ = v423
	var v424 int32
	_ = v424
	var v425 int32
	_ = v425
	var v429 int32
	_ = v429
	var v430 int32
	_ = v430
	var v433 int32
	_ = v433
	var v434 int32
	_ = v434
	var v438 int32
	_ = v438
	var v439 int32
	_ = v439
	var v442 int32
	_ = v442
	var v443 int32
	_ = v443
	var v446 int32
	_ = v446
	var v453 int32
	_ = v453
	var v454 int32
	_ = v454
	var v456 int32
	_ = v456
	var v457 int32
	_ = v457
	var v458 int32
	_ = v458
	var v459 int32
	_ = v459
	var v462 int32
	_ = v462
	var v463 int32
	_ = v463
	var v464 int32
	_ = v464
	var v465 int32
	_ = v465
	var v466 int32
	_ = v466
	var v467 int32
	_ = v467
	var v469 int32
	_ = v469
	var v470 int32
	_ = v470
	var v475 int32
	_ = v475
	var v478 int32
	_ = v478
	var v479 int32
	_ = v479
	var v486 int32
	_ = v486
	var v491 int32
	_ = v491
	var v495 int32
	_ = v495
	var v498 int32
	_ = v498
	var v499 int32
	_ = v499
	var v506 int32
	_ = v506
	var v511 int32
	_ = v511
	var v515 int32
	_ = v515
	var v518 int32
	_ = v518
	var v519 int32
	_ = v519
	var v526 int32
	_ = v526
	var v531 int32
	_ = v531
	var v535 int32
	_ = v535
	var v538 int32
	_ = v538
	var v539 int32
	_ = v539
	var v546 int32
	_ = v546
	var v551 int32
	_ = v551
	var v555 int32
	_ = v555
	var v558 int32
	_ = v558
	var v559 int32
	_ = v559
	var v566 int32
	_ = v566
	var v571 int32
	_ = v571
	var v583 int32
	_ = v583
	var v584 int32
	_ = v584
	var v585 int32
	_ = v585
	var v586 int32
	_ = v586
	var v587 int32
	_ = v587
	var v590 int32
	_ = v590
	var v592 int32
	_ = v592
	var v593 int32
	_ = v593
	var v594 int32
	_ = v594
	var v595 int32
	_ = v595
	var v596 int32
	_ = v596
	var v597 int32
	_ = v597
	var v598 int32
	_ = v598
	var v599 int32
	_ = v599
	var v600 int32
	_ = v600
	var v601 int32
	_ = v601
	var v602 int32
	_ = v602
	var v603 int32
	_ = v603
	var v609 int32
	_ = v609
	var v622 int32
	_ = v622
	var v631 int32
	_ = v631
	var v635 int32
	_ = v635
	var v637 int32
	_ = v637
	var v638 int32
	_ = v638
	var v642 int32
	_ = v642
	var v651 int32
	_ = v651
	var v654 int32
	_ = v654
	var v656 int32
	_ = v656
	var v658 int32
	_ = v658
	var v659 int32
	_ = v659
	var v682 int32
	_ = v682
	var v683 int32
	_ = v683
	var v684 int32
	_ = v684
	var v714 int32
	_ = v714
	var v718 int32
	_ = v718
	var v723 int32
	_ = v723
	var v727 int32
	_ = v727
	var v730 int32
	_ = v730
	var v736 int32
	_ = v736
	var v740 int32
	_ = v740
	var v745 int32
	_ = v745
	v5 = int32(0)
	v20 = m.G0
	v22 = v20 - int32(128)
	m.G0 = v22
	v24 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v26 = *(*int32)(unsafe.Add(mBase, _consts[210]))
	v32 = *(*int32)(unsafe.Add(mBase, _consts[4]))
	*(*int32)(unsafe.Add(mBase, uint32(v22+int32(124)))) = v32
	v35 = *(*int32)(unsafe.Add(mBase, _consts[122]))
	*(*int32)(unsafe.Add(mBase, uint32(v22+int32(120)))) = v35
	goto L1
L1:
	;
	v37 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	if v37 != 0 {
		goto L3
	} else {
		goto L4
	}
L2:
	;
	if v24 == int32(0) {
		goto L10
	} else {
		goto L11
	}
L3:
	;
	v39 = F_get_rolespec_oid(m, v37, int32(0))
	mBase = m.M
	v40 = m.ExcPending
	if v40 != 0 {
		goto L6
	} else {
		goto L7
	}
L4:
	;
	goto L5
L5:
	;
	v41 = *(*int32)(unsafe.Add(mBase, uint32(v22)+124))
	v42 = v41
	goto L2
L6:
	;
	return
L7:
	;
	v42 = v39
	goto L2
L8:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v727 = m.ExcPending
	if v727 != 0 {
		goto L6
	} else {
		goto L199
	}
L9:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v714 = m.ExcPending
	if v714 != 0 {
		goto L6
	} else {
		goto L196
	}
L10:
	;
	v46 = F_SearchSysCache1(m, int32(11), v42)
	mBase = m.M
	v47 = m.ExcPending
	if v47 != 0 {
		goto L6
	} else {
		goto L13
	}
L11:
	;
	v59 = v24
	goto L12
L12:
	;
	v63 = *(*int32)(unsafe.Add(mBase, _consts[226]))
	v64 = *(*int32)(unsafe.Add(mBase, uint32(v22)+124))
	v66 = F_object_aclcheck(m, int32(1262), v63, v64, int64(512))
	mBase = m.M
	v67 = m.ExcPending
	if v67 != 0 {
		goto L6
	} else {
		goto L17
	}
L13:
	;
	if v46 == int32(0) {
		goto L9
	} else {
		goto L14
	}
L14:
	;
	v50 = *(*int32)(unsafe.Add(mBase, uint32(v46)+16))
	v51 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v50)+22)))
	v55 = F_pstrdup(m, v50+v51+int32(4))
	mBase = m.M
	v56 = m.ExcPending
	if v56 != 0 {
		goto L6
	} else {
		goto L15
	}
L15:
	;
	F_ReleaseCatCache(m, v46)
	mBase = m.M
	v58 = m.ExcPending
	if v58 != 0 {
		goto L6
	} else {
		goto L16
	}
L16:
	;
	v59 = v55
	goto L12
L17:
	;
	if v66 != 0 {
		goto L18
	} else {
		goto L19
	}
L18:
	;
	v70 = *(*int32)(unsafe.Add(mBase, _consts[226]))
	v71 = F_get_database_name(m, v70)
	mBase = m.M
	v72 = m.ExcPending
	if v72 != 0 {
		goto L6
	} else {
		goto L21
	}
L19:
	;
	goto L20
L20:
	;
	v75 = *(*int32)(unsafe.Add(mBase, uint32(v22)+124))
	F_check_can_set_role(m, v75, v42)
	mBase = m.M
	v77 = m.ExcPending
	if v77 != 0 {
		goto L6
	} else {
		goto L23
	}
L21:
	;
	F_aclcheck_error(m, v66, int32(9), v71)
	mBase = m.M
	v74 = m.ExcPending
	if v74 != 0 {
		goto L6
	} else {
		goto L22
	}
L22:
	;
	goto L20
L23:
	;
	v79 = int32(*(*uint8)(unsafe.Add(mBase, _consts[273])))
	if v79 == int32(0) {
		goto L24
	} else {
		goto L25
	}
L24:
	;
	v82 = int32(0)
	v83 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v59))))
	if v83 != int32(112) {
		v92 = v82
		goto L28
	} else {
		goto L29
	}
L25:
	;
	goto L26
L26:
	;
	v93 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+16)))
	if v93 != int32(1) {
		goto L33
	} else {
		goto L34
	}
L27:
	;
	if v92 != 0 {
		goto L8
	} else {
		goto L31
	}
L28:
	;
	goto L27
L29:
	;
	v86 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v59)+1)))
	if v86 != int32(103) {
		v92 = v82
		goto L28
	} else {
		goto L30
	}
L30:
	;
	v89 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v59)+2)))
	v92 = base.B2i32(v89 == int32(95))
	goto L28
L31:
	;
	goto L26
L32:
	;
	m.G0 = v22 + int32(128)
	return
L33:
	;
	v131 = *(*int32)(unsafe.Add(mBase, uint32(v22)+124))
	if v42 != v131 {
		goto L43
	} else {
		goto L44
	}
L34:
	;
	v97 = F_get_namespace_oid(m, v59, int32(1))
	mBase = m.M
	v98 = m.ExcPending
	if v98 != 0 {
		goto L6
	} else {
		goto L35
	}
L35:
	;
	if v97 == int32(0) {
		goto L33
	} else {
		goto L36
	}
L36:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v22)+116)) = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v22)+112)) = v97
	*(*int32)(unsafe.Add(mBase, uint32(v22)+108)) = int32(2615)
	F_checkMembershipInCurrentExtension(m, v22+int32(108))
	mBase = m.M
	v109 = m.ExcPending
	if v109 != 0 {
		goto L6
	} else {
		goto L37
	}
L37:
	;
	v112 = F_errstart(m, int32(18), int32(0))
	mBase = m.M
	v113 = m.ExcPending
	if v113 != 0 {
		goto L6
	} else {
		goto L38
	}
L38:
	;
	if v112 == int32(0) {
		goto L32
	} else {
		goto L39
	}
L39:
	;
	F_errcode(m, int32(100794500))
	mBase = m.M
	v118 = m.ExcPending
	if v118 != 0 {
		goto L6
	} else {
		goto L40
	}
L40:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v22)+64)) = v59
	F_errmsg(m, int32(321267), v22-int32(-64))
	mBase = m.M
	v124 = m.ExcPending
	if v124 != 0 {
		goto L6
	} else {
		goto L41
	}
L41:
	;
	F_errfinish(m, int32(477739), int32(135), int32(414260))
	mBase = m.M
	v129 = m.ExcPending
	if v129 != 0 {
		goto L6
	} else {
		goto L42
	}
L42:
	;
	goto L32
L43:
	;
	v133 = *(*int32)(unsafe.Add(mBase, uint32(v22)+120))
	*(*int32)(unsafe.Add(mBase, _consts[122])) = v133 | int32(1)
	*(*int32)(unsafe.Add(mBase, _consts[4])) = v42
	goto L46
L44:
	;
	goto L45
L45:
	;
	v141 = F_NamespaceCreate(m, v59, v42, int32(0))
	mBase = m.M
	v142 = m.ExcPending
	if v142 != 0 {
		goto L6
	} else {
		goto L47
	}
L46:
	;
	goto L45
L47:
	;
	F_CommandCounterIncrement(m)
	mBase = m.M
	v144 = m.ExcPending
	if v144 != 0 {
		goto L6
	} else {
		goto L48
	}
L48:
	;
	v146 = int32(4453256)
	v148 = *(*int32)(unsafe.Add(mBase, _consts[159]))
	v150 = v148 + int32(1)
	*(*int32)(unsafe.Add(mBase, _consts[159])) = v150
	goto L49
L49:
	;
	F_initStringInfo(m, v22+int32(92))
	mBase = m.M
	v155 = m.ExcPending
	if v155 != 0 {
		goto L6
	} else {
		goto L50
	}
L50:
	;
	v158 = F_quote_identifier(m, v59)
	mBase = m.M
	v159 = m.ExcPending
	if v159 != 0 {
		goto L6
	} else {
		goto L51
	}
L51:
	;
	F_appendStringInfoString(m, v22+int32(92), v158)
	mBase = m.M
	v161 = m.ExcPending
	if v161 != 0 {
		goto L6
	} else {
		goto L52
	}
L52:
	;
	v172 = v26
	goto L53
L53:
	;
	v183 = int32(*(*int8)(unsafe.Add(mBase, uint32(v172))))
	goto L55
L54:
	;
	v193 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v172))))
	if v193 != 0 {
		goto L57
	} else {
		goto L58
	}
L55:
	;
	if base.B2i32(v183 == int32(32))|base.B2i32(base.Ui32((v183-int32(9))&int32(255)) < base.Ui32(int32(5))) != 0 {
		v172 = v172 + int32(1)
		goto L53
	} else {
		goto L56
	}
L56:
	;
	goto L54
L57:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v22)+48)) = v172
	F_appendStringInfo(m, v22+int32(92), int32(198468), v22+int32(48))
	mBase = m.M
	v201 = m.ExcPending
	if v201 != 0 {
		goto L6
	} else {
		goto L60
	}
L58:
	;
	goto L59
L59:
	;
	v204 = *(*int32)(unsafe.Add(mBase, uint32(v22)+92))
	F_set_config_option(m, int32(310148), v204, int32(6), int32(13), int32(2), int32(1))
	mBase = m.M
	v210 = m.ExcPending
	if v210 != 0 {
		goto L6
	} else {
		goto L61
	}
L60:
	;
	goto L59
L61:
	;
	v211 = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v22)+116)) = v211
	*(*int32)(unsafe.Add(mBase, uint32(v22)+40)) = v211
	v216 = *(*int32)(unsafe.Add(mBase, _consts[225]))
	*(*int32)(unsafe.Add(mBase, uint32(v22)+24)) = v216
	*(*int32)(unsafe.Add(mBase, uint32(v22)+112)) = v141
	*(*int32)(unsafe.Add(mBase, uint32(v22)+108)) = int32(2615)
	v222 = *(*int64)(unsafe.Add(mBase, _consts[224]))
	*(*int64)(unsafe.Add(mBase, uint32(v22)+16)) = v222
	v224 = *(*int64)(unsafe.Add(mBase, uint32(v22)+108))
	*(*int64)(unsafe.Add(mBase, uint32(v22)+32)) = v224
	F_EventTriggerCollectSimpleCommand(m, v22+int32(32), v22+int32(16), l0)
	mBase = m.M
	v231 = m.ExcPending
	if v231 != 0 {
		goto L6
	} else {
		goto L62
	}
L62:
	;
	v232 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	v233 = int32(0)
	v234 = m.G0
	v236 = v234 - int32(96)
	m.G0 = v236
	if v232 == v233 {
		v583 = v5
		v584 = v5
		v585 = v5
		v586 = v5
		v587 = v5
		v590 = v5
		goto L63
	} else {
		goto L64
	}
L63:
	;
	v592 = F_list_concat(m, int32(0), v583)
	mBase = m.M
	v593 = m.ExcPending
	if v593 != 0 {
		goto L6
	} else {
		goto L179
	}
L64:
	;
	v240 = *(*int32)(unsafe.Add(mBase, uint32(v232)+4))
	if v240 <= int32(0) {
		v583 = v5
		v584 = v5
		v585 = v5
		v586 = v5
		v587 = v5
		v590 = v5
		goto L63
	} else {
		goto L65
	}
L65:
	;
	v250 = v233
	v254 = v5
	v255 = v5
	v256 = v5
	v257 = v5
	v258 = v5
	v261 = v5
	goto L70
L66:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v555 = m.ExcPending
	if v555 != 0 {
		goto L6
	} else {
		goto L175
	}
L67:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v535 = m.ExcPending
	if v535 != 0 {
		goto L6
	} else {
		goto L171
	}
L68:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v515 = m.ExcPending
	if v515 != 0 {
		goto L6
	} else {
		goto L167
	}
L69:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v495 = m.ExcPending
	if v495 != 0 {
		goto L6
	} else {
		goto L163
	}
L70:
	;
	v262 = *(*int32)(unsafe.Add(mBase, uint32(v232)+12))
	v266 = *(*int32)(unsafe.Add(mBase, uint32(v262+v250<<(uint(int32(2))%32))))
	v267 = *(*int32)(unsafe.Add(mBase, uint32(v266)))
	switch v267 - int32(152) {
	case 0:
		goto L74
	case 1, 2, 3, 4, 5, 6, 7, 9, 10, 11, 12, 13, 14, 15, 16, 17, 18, 19, 20, 21, 22, 23, 24, 25, 26, 27, 28, 30, 31, 32, 33, 34, 35, 36:
		goto L80
	case 8:
		goto L78
	case 29:
		goto L75
	case 37:
		goto L79
	default:
		goto L81
	}
L71:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v475 = m.ExcPending
	if v475 != 0 {
		goto L6
	} else {
		goto L159
	}
L72:
	;
	goto L71
L73:
	;
	v469 = v250 + int32(1)
	v470 = *(*int32)(unsafe.Add(mBase, uint32(v232)+4))
	if v469 < v470 {
		v250 = v469
		v254 = v462
		v255 = v463
		v256 = v464
		v257 = v465
		v258 = v466
		v261 = v467
		goto L70
	} else {
		goto L158
	}
L74:
	;
	v458 = F_lappend(m, v261, v266)
	mBase = m.M
	v459 = m.ExcPending
	if v459 != 0 {
		goto L6
	} else {
		goto L157
	}
L75:
	;
	v424 = *(*int32)(unsafe.Add(mBase, uint32(v266)+12))
	v425 = *(*int32)(unsafe.Add(mBase, uint32(v424)+8))
	if v425 == int32(0) {
		goto L143
	} else {
		goto L144
	}
L76:
	;
	v390 = *(*int32)(unsafe.Add(mBase, uint32(v266)+8))
	v391 = *(*int32)(unsafe.Add(mBase, uint32(v390)+8))
	if v391 == int32(0) {
		goto L129
	} else {
		goto L130
	}
L77:
	;
	v356 = *(*int32)(unsafe.Add(mBase, uint32(v266)+4))
	v357 = *(*int32)(unsafe.Add(mBase, uint32(v356)+8))
	if v357 == int32(0) {
		goto L115
	} else {
		goto L116
	}
L78:
	;
	v322 = *(*int32)(unsafe.Add(mBase, uint32(v266)+4))
	v323 = *(*int32)(unsafe.Add(mBase, uint32(v322)+8))
	if v323 == int32(0) {
		goto L101
	} else {
		goto L102
	}
L79:
	;
	v288 = *(*int32)(unsafe.Add(mBase, uint32(v266)+4))
	v289 = *(*int32)(unsafe.Add(mBase, uint32(v288)+8))
	if v289 == int32(0) {
		goto L87
	} else {
		goto L88
	}
L80:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v277 = m.ExcPending
	if v277 != 0 {
		goto L6
	} else {
		goto L84
	}
L81:
	;
	if v267 == int32(204) {
		goto L76
	} else {
		goto L82
	}
L82:
	;
	if v267 == int32(230) {
		goto L77
	} else {
		goto L83
	}
L83:
	;
	goto L80
L84:
	;
	v278 = *(*int32)(unsafe.Add(mBase, uint32(v266)))
	*(*int32)(unsafe.Add(mBase, uint32(v236))) = v278
	F_errmsg_internal(m, int32(469568), v236)
	mBase = m.M
	v282 = m.ExcPending
	if v282 != 0 {
		goto L6
	} else {
		goto L85
	}
L85:
	;
	F_errfinish(m, int32(482739), int32(4185), int32(117409))
	mBase = m.M
	v287 = m.ExcPending
	if v287 != 0 {
		goto L6
	} else {
		goto L86
	}
L86:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L87:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v288)+8)) = v59
	v293 = F_lappend(m, v254, v266)
	mBase = m.M
	v294 = m.ExcPending
	if v294 != 0 {
		goto L6
	} else {
		goto L90
	}
L88:
	;
	goto L89
L89:
	;
	v297 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v289))))
	v298 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v59))))
	if v298 == int32(0) {
		v317 = v297
		v318 = v298
		goto L92
	} else {
		goto L93
	}
L90:
	;
	v462 = v293
	v463 = v255
	v464 = v256
	v465 = v257
	v466 = v258
	v467 = v261
	goto L73
L91:
	;
	if v318-v317 != 0 {
		goto L72
	} else {
		goto L99
	}
L92:
	;
	goto L91
L93:
	;
	if v297 != v298 {
		v317 = v297
		v318 = v298
		goto L92
	} else {
		goto L94
	}
L94:
	;
	v302 = v59
	v303 = v289
	goto L95
L95:
	;
	v306 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v303)+1)))
	v307 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v302)+1)))
	if v307 == int32(0) {
		v317 = v306
		v318 = v307
		goto L92
	} else {
		goto L97
	}
L96:
	;
	v317 = v306
	v318 = v307
	goto L92
L97:
	;
	v310 = int32(1)
	if v306 == v307 {
		v302 = v302 + v310
		v303 = v303 + v310
		goto L95
	} else {
		goto L98
	}
L98:
	;
	goto L96
L99:
	;
	v320 = F_lappend(m, v254, v266)
	mBase = m.M
	v321 = m.ExcPending
	if v321 != 0 {
		goto L6
	} else {
		goto L100
	}
L100:
	;
	v462 = v320
	v463 = v255
	v464 = v256
	v465 = v257
	v466 = v258
	v467 = v261
	goto L73
L101:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v322)+8)) = v59
	v327 = F_lappend(m, v255, v266)
	mBase = m.M
	v328 = m.ExcPending
	if v328 != 0 {
		goto L6
	} else {
		goto L104
	}
L102:
	;
	goto L103
L103:
	;
	v331 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v323))))
	v332 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v59))))
	if v332 == int32(0) {
		v351 = v331
		v352 = v332
		goto L106
	} else {
		goto L107
	}
L104:
	;
	v462 = v254
	v463 = v327
	v464 = v256
	v465 = v257
	v466 = v258
	v467 = v261
	goto L73
L105:
	;
	if v352-v351 != 0 {
		goto L69
	} else {
		goto L113
	}
L106:
	;
	goto L105
L107:
	;
	if v331 != v332 {
		v351 = v331
		v352 = v332
		goto L106
	} else {
		goto L108
	}
L108:
	;
	v336 = v59
	v337 = v323
	goto L109
L109:
	;
	v340 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v337)+1)))
	v341 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v336)+1)))
	if v341 == int32(0) {
		v351 = v340
		v352 = v341
		goto L106
	} else {
		goto L111
	}
L110:
	;
	v351 = v340
	v352 = v341
	goto L106
L111:
	;
	v344 = int32(1)
	if v340 == v341 {
		v336 = v336 + v344
		v337 = v337 + v344
		goto L109
	} else {
		goto L112
	}
L112:
	;
	goto L110
L113:
	;
	v354 = F_lappend(m, v255, v266)
	mBase = m.M
	v355 = m.ExcPending
	if v355 != 0 {
		goto L6
	} else {
		goto L114
	}
L114:
	;
	v462 = v254
	v463 = v354
	v464 = v256
	v465 = v257
	v466 = v258
	v467 = v261
	goto L73
L115:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v356)+8)) = v59
	v361 = F_lappend(m, v256, v266)
	mBase = m.M
	v362 = m.ExcPending
	if v362 != 0 {
		goto L6
	} else {
		goto L118
	}
L116:
	;
	goto L117
L117:
	;
	v365 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v357))))
	v366 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v59))))
	if v366 == int32(0) {
		v385 = v365
		v386 = v366
		goto L120
	} else {
		goto L121
	}
L118:
	;
	v462 = v254
	v463 = v255
	v464 = v361
	v465 = v257
	v466 = v258
	v467 = v261
	goto L73
L119:
	;
	if v386-v385 != 0 {
		goto L68
	} else {
		goto L127
	}
L120:
	;
	goto L119
L121:
	;
	if v365 != v366 {
		v385 = v365
		v386 = v366
		goto L120
	} else {
		goto L122
	}
L122:
	;
	v370 = v59
	v371 = v357
	goto L123
L123:
	;
	v374 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v371)+1)))
	v375 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v370)+1)))
	if v375 == int32(0) {
		v385 = v374
		v386 = v375
		goto L120
	} else {
		goto L125
	}
L124:
	;
	v385 = v374
	v386 = v375
	goto L120
L125:
	;
	v378 = int32(1)
	if v374 == v375 {
		v370 = v370 + v378
		v371 = v371 + v378
		goto L123
	} else {
		goto L126
	}
L126:
	;
	goto L124
L127:
	;
	v388 = F_lappend(m, v256, v266)
	mBase = m.M
	v389 = m.ExcPending
	if v389 != 0 {
		goto L6
	} else {
		goto L128
	}
L128:
	;
	v462 = v254
	v463 = v255
	v464 = v388
	v465 = v257
	v466 = v258
	v467 = v261
	goto L73
L129:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v390)+8)) = v59
	v395 = F_lappend(m, v257, v266)
	mBase = m.M
	v396 = m.ExcPending
	if v396 != 0 {
		goto L6
	} else {
		goto L132
	}
L130:
	;
	goto L131
L131:
	;
	v399 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v391))))
	v400 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v59))))
	if v400 == int32(0) {
		v419 = v399
		v420 = v400
		goto L134
	} else {
		goto L135
	}
L132:
	;
	v462 = v254
	v463 = v255
	v464 = v256
	v465 = v395
	v466 = v258
	v467 = v261
	goto L73
L133:
	;
	if v420-v419 != 0 {
		goto L67
	} else {
		goto L141
	}
L134:
	;
	goto L133
L135:
	;
	if v399 != v400 {
		v419 = v399
		v420 = v400
		goto L134
	} else {
		goto L136
	}
L136:
	;
	v404 = v59
	v405 = v391
	goto L137
L137:
	;
	v408 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v405)+1)))
	v409 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v404)+1)))
	if v409 == int32(0) {
		v419 = v408
		v420 = v409
		goto L134
	} else {
		goto L139
	}
L138:
	;
	v419 = v408
	v420 = v409
	goto L134
L139:
	;
	v412 = int32(1)
	if v408 == v409 {
		v404 = v404 + v412
		v405 = v405 + v412
		goto L137
	} else {
		goto L140
	}
L140:
	;
	goto L138
L141:
	;
	v422 = F_lappend(m, v257, v266)
	mBase = m.M
	v423 = m.ExcPending
	if v423 != 0 {
		goto L6
	} else {
		goto L142
	}
L142:
	;
	v462 = v254
	v463 = v255
	v464 = v256
	v465 = v422
	v466 = v258
	v467 = v261
	goto L73
L143:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v424)+8)) = v59
	v429 = F_lappend(m, v258, v266)
	mBase = m.M
	v430 = m.ExcPending
	if v430 != 0 {
		goto L6
	} else {
		goto L146
	}
L144:
	;
	goto L145
L145:
	;
	v433 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v425))))
	v434 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v59))))
	if v434 == int32(0) {
		v453 = v433
		v454 = v434
		goto L148
	} else {
		goto L149
	}
L146:
	;
	v462 = v254
	v463 = v255
	v464 = v256
	v465 = v257
	v466 = v429
	v467 = v261
	goto L73
L147:
	;
	if v454-v453 != 0 {
		goto L66
	} else {
		goto L155
	}
L148:
	;
	goto L147
L149:
	;
	if v433 != v434 {
		v453 = v433
		v454 = v434
		goto L148
	} else {
		goto L150
	}
L150:
	;
	v438 = v59
	v439 = v425
	goto L151
L151:
	;
	v442 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v439)+1)))
	v443 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v438)+1)))
	if v443 == int32(0) {
		v453 = v442
		v454 = v443
		goto L148
	} else {
		goto L153
	}
L152:
	;
	v453 = v442
	v454 = v443
	goto L148
L153:
	;
	v446 = int32(1)
	if v442 == v443 {
		v438 = v438 + v446
		v439 = v439 + v446
		goto L151
	} else {
		goto L154
	}
L154:
	;
	goto L152
L155:
	;
	v456 = F_lappend(m, v258, v266)
	mBase = m.M
	v457 = m.ExcPending
	if v457 != 0 {
		goto L6
	} else {
		goto L156
	}
L156:
	;
	v462 = v254
	v463 = v255
	v464 = v256
	v465 = v257
	v466 = v456
	v467 = v261
	goto L73
L157:
	;
	v462 = v254
	v463 = v255
	v464 = v256
	v465 = v257
	v466 = v258
	v467 = v458
	goto L73
L158:
	;
	v583 = v462
	v584 = v463
	v585 = v464
	v586 = v465
	v587 = v466
	v590 = v467
	goto L63
L159:
	;
	F_errcode(m, int32(84279428))
	mBase = m.M
	v478 = m.ExcPending
	if v478 != 0 {
		goto L6
	} else {
		goto L160
	}
L160:
	;
	v479 = *(*int32)(unsafe.Add(mBase, uint32(v288)+8))
	*(*int32)(unsafe.Add(mBase, uint32(v236)+20)) = v59
	*(*int32)(unsafe.Add(mBase, uint32(v236)+16)) = v479
	F_errmsg(m, int32(640959), v236+int32(16))
	mBase = m.M
	v486 = m.ExcPending
	if v486 != 0 {
		goto L6
	} else {
		goto L161
	}
L161:
	;
	F_errfinish(m, int32(482739), int32(4214), int32(368465))
	mBase = m.M
	v491 = m.ExcPending
	if v491 != 0 {
		goto L6
	} else {
		goto L162
	}
L162:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L163:
	;
	F_errcode(m, int32(84279428))
	mBase = m.M
	v498 = m.ExcPending
	if v498 != 0 {
		goto L6
	} else {
		goto L164
	}
L164:
	;
	v499 = *(*int32)(unsafe.Add(mBase, uint32(v322)+8))
	*(*int32)(unsafe.Add(mBase, uint32(v236)+36)) = v59
	*(*int32)(unsafe.Add(mBase, uint32(v236)+32)) = v499
	F_errmsg(m, int32(640959), v236+int32(32))
	mBase = m.M
	v506 = m.ExcPending
	if v506 != 0 {
		goto L6
	} else {
		goto L165
	}
L165:
	;
	F_errfinish(m, int32(482739), int32(4214), int32(368465))
	mBase = m.M
	v511 = m.ExcPending
	if v511 != 0 {
		goto L6
	} else {
		goto L166
	}
L166:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L167:
	;
	F_errcode(m, int32(84279428))
	mBase = m.M
	v518 = m.ExcPending
	if v518 != 0 {
		goto L6
	} else {
		goto L168
	}
L168:
	;
	v519 = *(*int32)(unsafe.Add(mBase, uint32(v356)+8))
	*(*int32)(unsafe.Add(mBase, uint32(v236)+52)) = v59
	*(*int32)(unsafe.Add(mBase, uint32(v236)+48)) = v519
	F_errmsg(m, int32(640959), v236+int32(48))
	mBase = m.M
	v526 = m.ExcPending
	if v526 != 0 {
		goto L6
	} else {
		goto L169
	}
L169:
	;
	F_errfinish(m, int32(482739), int32(4214), int32(368465))
	mBase = m.M
	v531 = m.ExcPending
	if v531 != 0 {
		goto L6
	} else {
		goto L170
	}
L170:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L171:
	;
	F_errcode(m, int32(84279428))
	mBase = m.M
	v538 = m.ExcPending
	if v538 != 0 {
		goto L6
	} else {
		goto L172
	}
L172:
	;
	v539 = *(*int32)(unsafe.Add(mBase, uint32(v390)+8))
	*(*int32)(unsafe.Add(mBase, uint32(v236)+68)) = v59
	*(*int32)(unsafe.Add(mBase, uint32(v236)+64)) = v539
	F_errmsg(m, int32(640959), v236-int32(-64))
	mBase = m.M
	v546 = m.ExcPending
	if v546 != 0 {
		goto L6
	} else {
		goto L173
	}
L173:
	;
	F_errfinish(m, int32(482739), int32(4214), int32(368465))
	mBase = m.M
	v551 = m.ExcPending
	if v551 != 0 {
		goto L6
	} else {
		goto L174
	}
L174:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L175:
	;
	F_errcode(m, int32(84279428))
	mBase = m.M
	v558 = m.ExcPending
	if v558 != 0 {
		goto L6
	} else {
		goto L176
	}
L176:
	;
	v559 = *(*int32)(unsafe.Add(mBase, uint32(v424)+8))
	*(*int32)(unsafe.Add(mBase, uint32(v236)+84)) = v59
	*(*int32)(unsafe.Add(mBase, uint32(v236)+80)) = v559
	F_errmsg(m, int32(640959), v236+int32(80))
	mBase = m.M
	v566 = m.ExcPending
	if v566 != 0 {
		goto L6
	} else {
		goto L177
	}
L177:
	;
	F_errfinish(m, int32(482739), int32(4214), int32(368465))
	mBase = m.M
	v571 = m.ExcPending
	if v571 != 0 {
		goto L6
	} else {
		goto L178
	}
L178:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L179:
	;
	v594 = F_list_concat(m, v592, v584)
	mBase = m.M
	v595 = m.ExcPending
	if v595 != 0 {
		goto L6
	} else {
		goto L180
	}
L180:
	;
	v596 = F_list_concat(m, v594, v585)
	mBase = m.M
	v597 = m.ExcPending
	if v597 != 0 {
		goto L6
	} else {
		goto L181
	}
L181:
	;
	v598 = F_list_concat(m, v596, v586)
	mBase = m.M
	v599 = m.ExcPending
	if v599 != 0 {
		goto L6
	} else {
		goto L182
	}
L182:
	;
	v600 = F_list_concat(m, v598, v587)
	mBase = m.M
	v601 = m.ExcPending
	if v601 != 0 {
		goto L6
	} else {
		goto L183
	}
L183:
	;
	v602 = F_list_concat(m, v600, v590)
	mBase = m.M
	v603 = m.ExcPending
	if v603 != 0 {
		goto L6
	} else {
		goto L184
	}
L184:
	;
	m.G0 = v236 + int32(96)
	if v602 == int32(0) {
		goto L185
	} else {
		goto L186
	}
L185:
	;
	F_AtEOXact_GUC(m, int32(1), v150)
	mBase = m.M
	v682 = m.ExcPending
	if v682 != 0 {
		goto L6
	} else {
		goto L194
	}
L186:
	;
	v609 = *(*int32)(unsafe.Add(mBase, uint32(v602)+4))
	if v609 <= int32(0) {
		goto L185
	} else {
		goto L187
	}
L187:
	;
	v622 = int32(0)
	goto L188
L188:
	;
	v631 = *(*int32)(unsafe.Add(mBase, uint32(v602)+12))
	v635 = *(*int32)(unsafe.Add(mBase, uint32(v631+v622<<(uint(int32(2))%32))))
	v637 = F_palloc0(m, int32(104))
	mBase = m.M
	v638 = m.ExcPending
	if v638 != 0 {
		goto L6
	} else {
		goto L190
	}
L189:
	;
	goto L185
L190:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v637)+96)) = l3
	*(*int32)(unsafe.Add(mBase, uint32(v637)+92)) = l2
	*(*int32)(unsafe.Add(mBase, uint32(v637)+88)) = v635
	v642 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v637)+26)) = uint8(v642)
	*(*int64)(unsafe.Add(mBase, uint32(v637))) = int64(25769804106)
	v651 = *(*int32)(unsafe.Add(mBase, _consts[274]))
	F_ProcessUtility(m, v637, l1, v642, int32(3), v642, v642, v651, v642)
	mBase = m.M
	v654 = m.ExcPending
	if v654 != 0 {
		goto L6
	} else {
		goto L191
	}
L191:
	;
	F_CommandCounterIncrement(m)
	mBase = m.M
	v656 = m.ExcPending
	if v656 != 0 {
		goto L6
	} else {
		goto L192
	}
L192:
	;
	v658 = v622 + int32(1)
	v659 = *(*int32)(unsafe.Add(mBase, uint32(v602)+4))
	if v658 < v659 {
		v622 = v658
		goto L188
	} else {
		goto L193
	}
L193:
	;
	goto L189
L194:
	;
	v683 = *(*int32)(unsafe.Add(mBase, uint32(v22)+124))
	v684 = *(*int32)(unsafe.Add(mBase, uint32(v22)+120))
	*(*int32)(unsafe.Add(mBase, _consts[122])) = v684
	*(*int32)(unsafe.Add(mBase, _consts[4])) = v683
	goto L195
L195:
	;
	goto L32
L196:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v22))) = v42
	F_errmsg_internal(m, int32(50220), v22)
	mBase = m.M
	v718 = m.ExcPending
	if v718 != 0 {
		goto L6
	} else {
		goto L197
	}
L197:
	;
	F_errfinish(m, int32(477739), int32(85), int32(414260))
	mBase = m.M
	v723 = m.ExcPending
	if v723 != 0 {
		goto L6
	} else {
		goto L198
	}
L198:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L199:
	;
	F_errcode(m, int32(151818372))
	mBase = m.M
	v730 = m.ExcPending
	if v730 != 0 {
		goto L6
	} else {
		goto L200
	}
L200:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v22)+80)) = v59
	F_errmsg(m, int32(680352), v22+int32(80))
	mBase = m.M
	v736 = m.ExcPending
	if v736 != 0 {
		goto L6
	} else {
		goto L201
	}
L201:
	;
	F_errdetail(m, int32(568895), int32(0))
	mBase = m.M
	v740 = m.ExcPending
	if v740 != 0 {
		goto L6
	} else {
		goto L202
	}
L202:
	;
	F_errfinish(m, int32(477739), int32(110), int32(414260))
	mBase = m.M
	v745 = m.ExcPending
	if v745 != 0 {
		goto L6
	} else {
		goto L203
	}
L203:
	;
	base.Wasm_trap_unreachable()
	for {
	}
}
func F_GetSchemaPublications(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v2 int32
	_ = v2
	var v9 int32
	_ = v9
	var v12 int32
	_ = v12
	var v13 int32
	_ = v13
	var v19 int32
	_ = v19
	var v20 int32
	_ = v20
	var v26 int32
	_ = v26
	var v27 int32
	_ = v27
	var v28 int32
	_ = v28
	var v30 int32
	_ = v30
	var v31 int32
	_ = v31
	var v32 int32
	_ = v32
	var v34 int32
	_ = v34
	var v35 int32
	_ = v35
	var v38 int32
	_ = v38
	var v42 int32
	_ = v42
	v2 = int32(0)
	v9 = F_SearchSysCacheList(m, int32(50), int32(1), l0, v2, v2)
	mBase = m.M
	v12 = m.ExcPending
	if v12 != 0 {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	return int32(0)
L2:
	;
	v13 = *(*int32)(unsafe.Add(mBase, uint32(v9)+40))
	if int32(0) < v13 {
		goto L3
	} else {
		goto L4
	}
L3:
	;
	v19 = int32(0)
	v20 = v2
	goto L6
L4:
	;
	v38 = v2
	goto L5
L5:
	;
	F_ReleaseCatCacheList(m, v9)
	mBase = m.M
	v42 = m.ExcPending
	if v42 != 0 {
		goto L1
	} else {
		goto L10
	}
L6:
	;
	v26 = *(*int32)(unsafe.Add(mBase, uint32(v9+int32(48)+v19<<(uint(int32(2))%32))))
	v27 = *(*int32)(unsafe.Add(mBase, uint32(v26)+56))
	v28 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v27)+22)))
	v30 = *(*int32)(unsafe.Add(mBase, uint32(v27+v28)+4))
	v31 = F_lappend_oid(m, v20, v30)
	mBase = m.M
	v32 = m.ExcPending
	if v32 != 0 {
		goto L1
	} else {
		goto L8
	}
L7:
	;
	v38 = v31
	goto L5
L8:
	;
	v34 = v19 + int32(1)
	v35 = *(*int32)(unsafe.Add(mBase, uint32(v9)+40))
	if v34 < v35 {
		v19 = v34
		v20 = v31
		goto L6
	} else {
		goto L9
	}
L9:
	;
	goto L7
L10:
	;
	return v38
}
func F_has_schema_privilege_id(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v6 int32
	_ = v6
	var v8 int32
	_ = v8
	var v10 int32
	_ = v10
	var v11 int32
	_ = v11
	var v12 int32
	_ = v12
	var v15 int32
	_ = v15
	var v16 int32
	_ = v16
	var v20 int32
	_ = v20
	var v22 int64
	_ = v22
	var v23 int32
	_ = v23
	var v26 int32
	_ = v26
	var v27 int32
	_ = v27
	var v28 int32
	_ = v28
	var v31 int32
	_ = v31
	var v35 int32
	_ = v35
	v6 = m.G0
	v8 = v6 - int32(16)
	m.G0 = v8
	v10 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v11 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
	v12 = F_pg_detoast_datum_packed(m, v11)
	mBase = m.M
	v15 = m.ExcPending
	if v15 != 0 {
		return int32(0)
	} else {
		v16 = int32(0)
		*(*uint8)(unsafe.Add(mBase, uint32(v8)+15)) = uint8(v16)
		v20 = *(*int32)(unsafe.Add(mBase, _consts[4]))
		v22 = F_convert_any_priv_string(m, v12, int32(1624432))
		mBase = m.M
		v23 = m.ExcPending
		if v23 != 0 {
			return int32(0)
		} else {
			v26 = F_object_aclcheck_ext(m, int32(2615), v10, v20, v22, v8+int32(15))
			mBase = m.M
			v27 = m.ExcPending
			if v27 != 0 {
				return int32(0)
			} else {
				v28 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v8)+15)))
				if v28 == int32(1) {
					v31 = int32(1)
					*(*uint8)(unsafe.Add(mBase, uint32(l0)+16)) = uint8(v31)
					v35 = int32(0)
				} else {
					v35 = base.B2i32(v26 == int32(0))
				}
				m.G0 = v8 + int32(16)
				return v35
			}
		}
	}
}
func F_has_schema_privilege_name(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v4 int32
	_ = v4
	var v5 int32
	_ = v5
	var v8 int32
	_ = v8
	var v9 int32
	_ = v9
	var v10 int32
	_ = v10
	var v11 int32
	_ = v11
	var v13 int32
	_ = v13
	var v15 int32
	_ = v15
	var v16 int32
	_ = v16
	var v18 int32
	_ = v18
	var v19 int32
	_ = v19
	var v21 int64
	_ = v21
	var v22 int32
	_ = v22
	var v23 int32
	_ = v23
	var v24 int32
	_ = v24
	v4 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v5 = F_pg_detoast_datum_packed(m, v4)
	mBase = m.M
	v8 = m.ExcPending
	if v8 != 0 {
		return int32(0)
	} else {
		v9 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
		v10 = F_pg_detoast_datum_packed(m, v9)
		mBase = m.M
		v11 = m.ExcPending
		if v11 != 0 {
			return int32(0)
		} else {
			v13 = *(*int32)(unsafe.Add(mBase, _consts[4]))
			v15 = F_text_to_cstring(m, v5)
			mBase = m.M
			v16 = m.ExcPending
			if v16 != 0 {
				return int32(0)
			} else {
				v18 = F_get_namespace_oid(m, v15, int32(0))
				mBase = m.M
				v19 = m.ExcPending
				if v19 != 0 {
					return int32(0)
				} else {
					v21 = F_convert_any_priv_string(m, v10, int32(1624432))
					mBase = m.M
					v22 = m.ExcPending
					if v22 != 0 {
						return int32(0)
					} else {
						v23 = F_object_aclcheck(m, int32(2615), v18, v13, v21)
						mBase = m.M
						v24 = m.ExcPending
						if v24 != 0 {
							return int32(0)
						} else {
							return base.B2i32(v23 == int32(0))
						}
					}
				}
			}
		}
	}
}
func F_has_schema_privilege_name_id(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v7 int32
	_ = v7
	var v9 int32
	_ = v9
	var v11 int32
	_ = v11
	var v12 int32
	_ = v12
	var v13 int32
	_ = v13
	var v14 int32
	_ = v14
	var v17 int32
	_ = v17
	var v18 int32
	_ = v18
	var v21 int32
	_ = v21
	var v22 int32
	_ = v22
	var v24 int64
	_ = v24
	var v25 int32
	_ = v25
	var v28 int32
	_ = v28
	var v29 int32
	_ = v29
	var v30 int32
	_ = v30
	var v33 int32
	_ = v33
	var v37 int32
	_ = v37
	v7 = m.G0
	v9 = v7 - int32(16)
	m.G0 = v9
	v11 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
	v12 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v13 = *(*int32)(unsafe.Add(mBase, uint32(l0)+36))
	v14 = F_pg_detoast_datum_packed(m, v13)
	mBase = m.M
	v17 = m.ExcPending
	if v17 != 0 {
		return int32(0)
	} else {
		v18 = int32(0)
		*(*uint8)(unsafe.Add(mBase, uint32(v9)+15)) = uint8(v18)
		v21 = F_get_role_oid_or_public(m, v12)
		mBase = m.M
		v22 = m.ExcPending
		if v22 != 0 {
			return int32(0)
		} else {
			v24 = F_convert_any_priv_string(m, v14, int32(1624432))
			mBase = m.M
			v25 = m.ExcPending
			if v25 != 0 {
				return int32(0)
			} else {
				v28 = F_object_aclcheck_ext(m, int32(2615), v11, v21, v24, v9+int32(15))
				mBase = m.M
				v29 = m.ExcPending
				if v29 != 0 {
					return int32(0)
				} else {
					v30 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v9)+15)))
					if v30 == int32(1) {
						v33 = int32(1)
						*(*uint8)(unsafe.Add(mBase, uint32(l0)+16)) = uint8(v33)
						v37 = int32(0)
					} else {
						v37 = base.B2i32(v28 == int32(0))
					}
					m.G0 = v9 + int32(16)
					return v37
				}
			}
		}
	}
}
func F_has_schema_privilege_name_name(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v4 int32
	_ = v4
	var v5 int32
	_ = v5
	var v6 int32
	_ = v6
	var v9 int32
	_ = v9
	var v10 int32
	_ = v10
	var v11 int32
	_ = v11
	var v12 int32
	_ = v12
	var v13 int32
	_ = v13
	var v14 int32
	_ = v14
	var v16 int32
	_ = v16
	var v17 int32
	_ = v17
	var v19 int32
	_ = v19
	var v20 int32
	_ = v20
	var v22 int64
	_ = v22
	var v23 int32
	_ = v23
	var v24 int32
	_ = v24
	var v25 int32
	_ = v25
	v4 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v5 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
	v6 = F_pg_detoast_datum_packed(m, v5)
	mBase = m.M
	v9 = m.ExcPending
	if v9 != 0 {
		return int32(0)
	} else {
		v10 = *(*int32)(unsafe.Add(mBase, uint32(l0)+36))
		v11 = F_pg_detoast_datum_packed(m, v10)
		mBase = m.M
		v12 = m.ExcPending
		if v12 != 0 {
			return int32(0)
		} else {
			v13 = F_get_role_oid_or_public(m, v4)
			mBase = m.M
			v14 = m.ExcPending
			if v14 != 0 {
				return int32(0)
			} else {
				v16 = F_text_to_cstring(m, v6)
				mBase = m.M
				v17 = m.ExcPending
				if v17 != 0 {
					return int32(0)
				} else {
					v19 = F_get_namespace_oid(m, v16, int32(0))
					mBase = m.M
					v20 = m.ExcPending
					if v20 != 0 {
						return int32(0)
					} else {
						v22 = F_convert_any_priv_string(m, v11, int32(1624432))
						mBase = m.M
						v23 = m.ExcPending
						if v23 != 0 {
							return int32(0)
						} else {
							v24 = F_object_aclcheck(m, int32(2615), v19, v13, v22)
							mBase = m.M
							v25 = m.ExcPending
							if v25 != 0 {
								return int32(0)
							} else {
								return base.B2i32(v24 == int32(0))
							}
						}
					}
				}
			}
		}
	}
}
func F_schema_to_xml_and_xmlschema(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v3 int32
	_ = v3
	var v4 int32
	_ = v4
	var v5 int32
	_ = v5
	var v8 int32
	_ = v8
	var v9 int32
	_ = v9
	var v10 int32
	_ = v10
	var v12 int32
	_ = v12
	var v13 int32
	_ = v13
	var v15 int32
	_ = v15
	v3 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v4 = *(*int32)(unsafe.Add(mBase, uint32(l0)+44))
	v5 = F_pg_detoast_datum_packed(m, v4)
	mBase = m.M
	v8 = m.ExcPending
	if v8 != 0 {
		return int32(0)
	} else {
		v9 = F_text_to_cstring(m, v5)
		mBase = m.M
		v10 = m.ExcPending
		if v10 != 0 {
			return int32(0)
		} else {
			v12 = F_LookupExplicitNamespace(m, v3, int32(0))
			mBase = m.M
			v13 = m.ExcPending
			if v13 != 0 {
				return int32(0)
			} else {
				F_schema_to_xmlschema_internal(m, v3, v9)
				mBase = m.M
				v15 = m.ExcPending
				if v15 != 0 {
					return int32(0)
				} else {
					base.Wasm_trap_unreachable()
					for {
					}
				}
			}
		}
	}
}
func F_schema_to_xmlschema_internal(m *base.Module, l0 int32, l1 int32) {
	mBase := m.M
	_ = mBase
	var v3 int32
	_ = v3
	var v7 int32
	_ = v7
	var v9 int32
	_ = v9
	var v11 int32
	_ = v11
	var v12 int32
	_ = v12
	var v14 int32
	_ = v14
	var v15 int32
	_ = v15
	var v18 int32
	_ = v18
	var v19 int32
	_ = v19
	var v25 int32
	_ = v25
	var v28 int32
	_ = v28
	var v31 int32
	_ = v31
	var v35 int32
	_ = v35
	var v41 int32
	_ = v41
	var v42 int32
	_ = v42
	var v43 int32
	_ = v43
	var v44 int32
	_ = v44
	var v47 int32
	_ = v47
	var v48 int32
	_ = v48
	var v53 int32
	_ = v53
	var v56 int32
	_ = v56
	var v57 int32
	_ = v57
	var v61 int32
	_ = v61
	var v63 int32
	_ = v63
	var v64 int32
	_ = v64
	var v65 int32
	_ = v65
	var v66 int32
	_ = v66
	var v67 int32
	_ = v67
	var v68 int32
	_ = v68
	var v69 int32
	_ = v69
	var v72 int32
	_ = v72
	var v74 int32
	_ = v74
	var v75 int32
	_ = v75
	var v82 int32
	_ = v82
	var v83 int32
	_ = v83
	var v84 int32
	_ = v84
	var v86 int32
	_ = v86
	var v87 int32
	_ = v87
	var v89 int32
	_ = v89
	var v92 int32
	_ = v92
	var v93 int32
	_ = v93
	var v94 int32
	_ = v94
	var v95 int32
	_ = v95
	var v96 int32
	_ = v96
	var v98 int32
	_ = v98
	var v99 int32
	_ = v99
	var v100 int32
	_ = v100
	v3 = int32(0)
	v7 = m.G0
	v9 = v7 - int32(48)
	m.G0 = v9
	v11 = F_makeStringInfo(m)
	mBase = m.M
	v12 = m.ExcPending
	if v12 != 0 {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	return
L2:
	;
	v14 = F_LookupExplicitNamespace(m, l0, int32(0))
	mBase = m.M
	v15 = m.ExcPending
	if v15 != 0 {
		goto L1
	} else {
		goto L3
	}
L3:
	;
	F_appendStringInfoString(m, v11, int32(695008))
	mBase = m.M
	v18 = m.ExcPending
	if v18 != 0 {
		goto L1
	} else {
		goto L4
	}
L4:
	;
	v19 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l1))))
	if v19 != 0 {
		goto L5
	} else {
		goto L6
	}
L5:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v9)+16)) = l1
	F_appendStringInfo(m, v11, int32(694596), v9+int32(16))
	mBase = m.M
	v25 = m.ExcPending
	if v25 != 0 {
		goto L1
	} else {
		goto L8
	}
L6:
	;
	goto L7
L7:
	;
	F_appendStringInfoString(m, v11, int32(722246))
	mBase = m.M
	v28 = m.ExcPending
	if v28 != 0 {
		goto L1
	} else {
		goto L9
	}
L8:
	;
	goto L7
L9:
	;
	F_SPI_connect_ext(m, int32(0))
	mBase = m.M
	v31 = m.ExcPending
	if v31 != 0 {
		goto L1
	} else {
		goto L10
	}
L10:
	;
	F_initStringInfo(m, v9+int32(32))
	mBase = m.M
	v35 = m.ExcPending
	if v35 != 0 {
		goto L1
	} else {
		goto L11
	}
L11:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v9))) = v14
	F_appendStringInfo(m, v9+int32(32), int32(528856), v9)
	mBase = m.M
	v41 = m.ExcPending
	if v41 != 0 {
		goto L1
	} else {
		goto L12
	}
L12:
	;
	v42 = *(*int32)(unsafe.Add(mBase, uint32(v9)+32))
	v43 = F_query_to_oid_list(m, v42)
	mBase = m.M
	v44 = m.ExcPending
	if v44 != 0 {
		goto L1
	} else {
		goto L14
	}
L13:
	;
	v83 = F_map_sql_typecoll_to_xmlschema_types(m, v82)
	mBase = m.M
	v84 = m.ExcPending
	if v84 != 0 {
		goto L1
	} else {
		goto L24
	}
L14:
	;
	if v43 == int32(0) {
		v82 = v3
		goto L13
	} else {
		goto L15
	}
L15:
	;
	v47 = int32(0)
	v48 = *(*int32)(unsafe.Add(mBase, uint32(v43)+4))
	if v48 <= v47 {
		v82 = v3
		goto L13
	} else {
		goto L16
	}
L16:
	;
	v53 = v47
	v56 = v3
	goto L17
L17:
	;
	v57 = *(*int32)(unsafe.Add(mBase, uint32(v43)+12))
	v61 = *(*int32)(unsafe.Add(mBase, uint32(v57+v53<<(uint(int32(2))%32))))
	v63 = F_table_open(m, v61, int32(1))
	mBase = m.M
	v64 = m.ExcPending
	if v64 != 0 {
		goto L1
	} else {
		goto L19
	}
L18:
	;
	v82 = v68
	goto L13
L19:
	;
	v65 = *(*int32)(unsafe.Add(mBase, uint32(v63)+52))
	v66 = F_CreateTupleDescCopy(m, v65)
	mBase = m.M
	v67 = m.ExcPending
	if v67 != 0 {
		goto L1
	} else {
		goto L20
	}
L20:
	;
	v68 = F_lappend(m, v56, v66)
	mBase = m.M
	v69 = m.ExcPending
	if v69 != 0 {
		goto L1
	} else {
		goto L21
	}
L21:
	;
	F_sequence_close(m, v63, int32(0))
	mBase = m.M
	v72 = m.ExcPending
	if v72 != 0 {
		goto L1
	} else {
		goto L22
	}
L22:
	;
	v74 = v53 + int32(1)
	v75 = *(*int32)(unsafe.Add(mBase, uint32(v43)+4))
	if v74 < v75 {
		v53 = v74
		v56 = v68
		goto L17
	} else {
		goto L23
	}
L23:
	;
	goto L18
L24:
	;
	F_appendStringInfoString(m, v11, v83)
	mBase = m.M
	v86 = m.ExcPending
	if v86 != 0 {
		goto L1
	} else {
		goto L25
	}
L25:
	;
	v87 = m.G0
	v89 = v87 - int32(16)
	m.G0 = v89
	v92 = *(*int32)(unsafe.Add(mBase, _consts[226]))
	v93 = F_get_database_name(m, v92)
	mBase = m.M
	v94 = m.ExcPending
	if v94 != 0 {
		goto L1
	} else {
		goto L26
	}
L26:
	;
	v95 = F_get_namespace_name(m, v14)
	mBase = m.M
	v96 = m.ExcPending
	if v96 != 0 {
		goto L1
	} else {
		goto L27
	}
L27:
	;
	F_initStringInfo(m, v89)
	mBase = m.M
	v98 = m.ExcPending
	if v98 != 0 {
		goto L1
	} else {
		goto L28
	}
L28:
	;
	v99 = F_map_sql_identifier_to_xml_name(m)
	mBase = m.M
	v100 = m.ExcPending
	if v100 != 0 {
		goto L1
	} else {
		goto L29
	}
L29:
	;
	base.Wasm_trap_unreachable()
	for {
	}
}
