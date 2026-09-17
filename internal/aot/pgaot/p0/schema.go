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
	var v153 int32
	_ = v153
	var v155 int32
	_ = v155
	var v156 int32
	_ = v156
	var v157 int32
	_ = v157
	var v159 int32
	_ = v159
	var v170 int32
	_ = v170
	var v181 int32
	_ = v181
	var v191 int32
	_ = v191
	var v199 int32
	_ = v199
	var v202 int32
	_ = v202
	var v208 int32
	_ = v208
	var v212 int32
	_ = v212
	var v217 int64
	_ = v217
	var v220 int32
	_ = v220
	var v222 int64
	_ = v222
	var v229 int32
	_ = v229
	var v230 int32
	_ = v230
	var v231 int32
	_ = v231
	var v233 int32
	_ = v233
	var v237 int32
	_ = v237
	var v251 int32
	_ = v251
	var v252 int32
	_ = v252
	var v253 int32
	_ = v253
	var v254 int32
	_ = v254
	var v255 int32
	_ = v255
	var v256 int32
	_ = v256
	var v257 int32
	_ = v257
	var v259 int32
	_ = v259
	var v263 int32
	_ = v263
	var v264 int32
	_ = v264
	var v274 int32
	_ = v274
	var v275 int32
	_ = v275
	var v279 int32
	_ = v279
	var v284 int32
	_ = v284
	var v285 int32
	_ = v285
	var v286 int32
	_ = v286
	var v290 int32
	_ = v290
	var v291 int32
	_ = v291
	var v294 int32
	_ = v294
	var v297 int32
	_ = v297
	var v300 int32
	_ = v300
	var v301 int32
	_ = v301
	var v304 int32
	_ = v304
	var v305 int32
	_ = v305
	var v308 int32
	_ = v308
	var v315 int32
	_ = v315
	var v316 int32
	_ = v316
	var v318 int32
	_ = v318
	var v319 int32
	_ = v319
	var v320 int32
	_ = v320
	var v321 int32
	_ = v321
	var v325 int32
	_ = v325
	var v326 int32
	_ = v326
	var v329 int32
	_ = v329
	var v332 int32
	_ = v332
	var v335 int32
	_ = v335
	var v336 int32
	_ = v336
	var v339 int32
	_ = v339
	var v340 int32
	_ = v340
	var v343 int32
	_ = v343
	var v350 int32
	_ = v350
	var v351 int32
	_ = v351
	var v353 int32
	_ = v353
	var v354 int32
	_ = v354
	var v355 int32
	_ = v355
	var v356 int32
	_ = v356
	var v360 int32
	_ = v360
	var v361 int32
	_ = v361
	var v364 int32
	_ = v364
	var v367 int32
	_ = v367
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
	var v402 int32
	_ = v402
	var v405 int32
	_ = v405
	var v406 int32
	_ = v406
	var v409 int32
	_ = v409
	var v410 int32
	_ = v410
	var v413 int32
	_ = v413
	var v420 int32
	_ = v420
	var v421 int32
	_ = v421
	var v423 int32
	_ = v423
	var v424 int32
	_ = v424
	var v425 int32
	_ = v425
	var v426 int32
	_ = v426
	var v430 int32
	_ = v430
	var v431 int32
	_ = v431
	var v434 int32
	_ = v434
	var v437 int32
	_ = v437
	var v440 int32
	_ = v440
	var v441 int32
	_ = v441
	var v444 int32
	_ = v444
	var v445 int32
	_ = v445
	var v448 int32
	_ = v448
	var v455 int32
	_ = v455
	var v456 int32
	_ = v456
	var v458 int32
	_ = v458
	var v459 int32
	_ = v459
	var v460 int32
	_ = v460
	var v461 int32
	_ = v461
	var v464 int32
	_ = v464
	var v465 int32
	_ = v465
	var v466 int32
	_ = v466
	var v467 int32
	_ = v467
	var v468 int32
	_ = v468
	var v469 int32
	_ = v469
	var v471 int32
	_ = v471
	var v472 int32
	_ = v472
	var v485 int32
	_ = v485
	var v486 int32
	_ = v486
	var v487 int32
	_ = v487
	var v488 int32
	_ = v488
	var v489 int32
	_ = v489
	var v491 int32
	_ = v491
	var v494 int32
	_ = v494
	var v495 int32
	_ = v495
	var v496 int32
	_ = v496
	var v497 int32
	_ = v497
	var v498 int32
	_ = v498
	var v499 int32
	_ = v499
	var v500 int32
	_ = v500
	var v501 int32
	_ = v501
	var v502 int32
	_ = v502
	var v503 int32
	_ = v503
	var v504 int32
	_ = v504
	var v505 int32
	_ = v505
	var v512 int32
	_ = v512
	var v515 int32
	_ = v515
	var v516 int32
	_ = v516
	var v523 int32
	_ = v523
	var v527 int32
	_ = v527
	var v530 int32
	_ = v530
	var v531 int32
	_ = v531
	var v538 int32
	_ = v538
	var v542 int32
	_ = v542
	var v545 int32
	_ = v545
	var v546 int32
	_ = v546
	var v553 int32
	_ = v553
	var v557 int32
	_ = v557
	var v560 int32
	_ = v560
	var v561 int32
	_ = v561
	var v568 int32
	_ = v568
	var v572 int32
	_ = v572
	var v575 int32
	_ = v575
	var v576 int32
	_ = v576
	var v583 int32
	_ = v583
	var v586 int32
	_ = v586
	var v599 int32
	_ = v599
	var v608 int32
	_ = v608
	var v612 int32
	_ = v612
	var v614 int32
	_ = v614
	var v615 int32
	_ = v615
	var v619 int32
	_ = v619
	var v628 int32
	_ = v628
	var v631 int32
	_ = v631
	var v633 int32
	_ = v633
	var v635 int32
	_ = v635
	var v636 int32
	_ = v636
	var v659 int32
	_ = v659
	var v660 int32
	_ = v660
	var v661 int32
	_ = v661
	var v691 int32
	_ = v691
	var v695 int32
	_ = v695
	var v700 int32
	_ = v700
	var v704 int32
	_ = v704
	var v707 int32
	_ = v707
	var v713 int32
	_ = v713
	var v717 int32
	_ = v717
	var v722 int32
	_ = v722
	var v730 int32
	_ = v730
	v5 = int32(0)
	v20 = m.G0
	v22 = v20 - int32(128)
	m.G0 = v22
	v24 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v26 = *(*int32)(unsafe.Add(mBase, _c_F_CreateSchemaCommand[0]))
	v32 = *(*int32)(unsafe.Add(mBase, _c_F_CreateSchemaCommand[1]))
	*(*int32)(unsafe.Add(mBase, uint32(v22+int32(124)))) = v32
	v35 = *(*int32)(unsafe.Add(mBase, _c_F_CreateSchemaCommand[2]))
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
		goto L11
	} else {
		goto L12
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
	F_errfinish(m, int32(_a_F_CreateSchemaCommand_0), int32(_a_F_CreateSchemaCommand_1), int32(_a_F_CreateSchemaCommand_2))
	mBase = m.M
	v730 = m.ExcPending
	if v730 != 0 {
		goto L6
	} else {
		goto L196
	}
L9:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v704 = m.ExcPending
	if v704 != 0 {
		goto L6
	} else {
		goto L191
	}
L10:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v691 = m.ExcPending
	if v691 != 0 {
		goto L6
	} else {
		goto L188
	}
L11:
	;
	v46 = F_SearchSysCache1(m, int32(11), v42)
	mBase = m.M
	v47 = m.ExcPending
	if v47 != 0 {
		goto L6
	} else {
		goto L14
	}
L12:
	;
	v59 = v24
	goto L13
L13:
	;
	v63 = *(*int32)(unsafe.Add(mBase, _c_F_CreateSchemaCommand[3]))
	v64 = *(*int32)(unsafe.Add(mBase, uint32(v22)+124))
	v66 = F_object_aclcheck(m, int32(1262), v63, v64, int64(512))
	mBase = m.M
	v67 = m.ExcPending
	if v67 != 0 {
		goto L6
	} else {
		goto L18
	}
L14:
	;
	if v46 == int32(0) {
		goto L10
	} else {
		goto L15
	}
L15:
	;
	v50 = *(*int32)(unsafe.Add(mBase, uint32(v46)+16))
	v51 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v50)+22)))
	v55 = F_pstrdup(m, v50+v51+int32(4))
	mBase = m.M
	v56 = m.ExcPending
	if v56 != 0 {
		goto L6
	} else {
		goto L16
	}
L16:
	;
	F_ReleaseCatCache(m, v46)
	mBase = m.M
	v58 = m.ExcPending
	if v58 != 0 {
		goto L6
	} else {
		goto L17
	}
L17:
	;
	v59 = v55
	goto L13
L18:
	;
	if v66 != 0 {
		goto L19
	} else {
		goto L20
	}
L19:
	;
	v70 = *(*int32)(unsafe.Add(mBase, _c_F_CreateSchemaCommand[3]))
	v71 = F_get_database_name(m, v70)
	mBase = m.M
	v72 = m.ExcPending
	if v72 != 0 {
		goto L6
	} else {
		goto L22
	}
L20:
	;
	goto L21
L21:
	;
	v75 = *(*int32)(unsafe.Add(mBase, uint32(v22)+124))
	F_check_can_set_role(m, v75, v42)
	mBase = m.M
	v77 = m.ExcPending
	if v77 != 0 {
		goto L6
	} else {
		goto L24
	}
L22:
	;
	F_aclcheck_error(m, v66, int32(9), v71)
	mBase = m.M
	v74 = m.ExcPending
	if v74 != 0 {
		goto L6
	} else {
		goto L23
	}
L23:
	;
	goto L21
L24:
	;
	v79 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_CreateSchemaCommand[4])))
	if v79 == int32(0) {
		goto L25
	} else {
		goto L26
	}
L25:
	;
	v82 = int32(0)
	v83 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v59))))
	if v83 != int32(112) {
		v92 = v82
		goto L29
	} else {
		goto L30
	}
L26:
	;
	goto L27
L27:
	;
	v93 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+16)))
	if v93 != int32(1) {
		goto L34
	} else {
		goto L35
	}
L28:
	;
	if v92 != 0 {
		goto L9
	} else {
		goto L32
	}
L29:
	;
	goto L28
L30:
	;
	v86 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v59)+1)))
	if v86 != int32(103) {
		v92 = v82
		goto L29
	} else {
		goto L31
	}
L31:
	;
	v89 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v59)+2)))
	v92 = base.B2i32(v89 == int32(95))
	goto L29
L32:
	;
	goto L27
L33:
	;
	m.G0 = v22 + int32(128)
	return
L34:
	;
	v131 = *(*int32)(unsafe.Add(mBase, uint32(v22)+124))
	if v42 != v131 {
		goto L44
	} else {
		goto L45
	}
L35:
	;
	v97 = F_get_namespace_oid(m, v59, int32(1))
	mBase = m.M
	v98 = m.ExcPending
	if v98 != 0 {
		goto L6
	} else {
		goto L36
	}
L36:
	;
	if v97 == int32(0) {
		goto L34
	} else {
		goto L37
	}
L37:
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
		goto L38
	}
L38:
	;
	v112 = F_errstart(m, int32(18), int32(0))
	mBase = m.M
	v113 = m.ExcPending
	if v113 != 0 {
		goto L6
	} else {
		goto L39
	}
L39:
	;
	if v112 == int32(0) {
		goto L33
	} else {
		goto L40
	}
L40:
	;
	F_errcode(m, int32(100794500))
	mBase = m.M
	v118 = m.ExcPending
	if v118 != 0 {
		goto L6
	} else {
		goto L41
	}
L41:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v22)+64)) = v59
	F_errmsg(m, int32(_a_F_CreateSchemaCommand_3), v22-int32(-64))
	mBase = m.M
	v124 = m.ExcPending
	if v124 != 0 {
		goto L6
	} else {
		goto L42
	}
L42:
	;
	F_errfinish(m, int32(_a_F_CreateSchemaCommand_4), int32(135), int32(_a_F_CreateSchemaCommand_5))
	mBase = m.M
	v129 = m.ExcPending
	if v129 != 0 {
		goto L6
	} else {
		goto L43
	}
L43:
	;
	goto L33
L44:
	;
	v133 = *(*int32)(unsafe.Add(mBase, uint32(v22)+120))
	*(*int32)(unsafe.Add(mBase, _c_F_CreateSchemaCommand[2])) = v133 | int32(1)
	*(*int32)(unsafe.Add(mBase, _c_F_CreateSchemaCommand[1])) = v42
	goto L47
L45:
	;
	goto L46
L46:
	;
	v141 = F_NamespaceCreate(m, v59, v42, int32(0))
	mBase = m.M
	v142 = m.ExcPending
	if v142 != 0 {
		goto L6
	} else {
		goto L48
	}
L47:
	;
	goto L46
L48:
	;
	F_CommandCounterIncrement(m)
	mBase = m.M
	v144 = m.ExcPending
	if v144 != 0 {
		goto L6
	} else {
		goto L49
	}
L49:
	;
	v146 = int32(_a_F_CreateSchemaCommand_6)
	v148 = *(*int32)(unsafe.Add(mBase, _c_F_CreateSchemaCommand[5]))
	v150 = v148 + int32(1)
	*(*int32)(unsafe.Add(mBase, _c_F_CreateSchemaCommand[5])) = v150
	goto L50
L50:
	;
	v153 = v22 + int32(92)
	F_initStringInfo(m, v153)
	mBase = m.M
	v155 = m.ExcPending
	if v155 != 0 {
		goto L6
	} else {
		goto L51
	}
L51:
	;
	v156 = F_quote_identifier(m, v59)
	mBase = m.M
	v157 = m.ExcPending
	if v157 != 0 {
		goto L6
	} else {
		goto L52
	}
L52:
	;
	F_appendStringInfoString(m, v153, v156)
	mBase = m.M
	v159 = m.ExcPending
	if v159 != 0 {
		goto L6
	} else {
		goto L53
	}
L53:
	;
	v170 = v26
	goto L54
L54:
	;
	v181 = int32(*(*int8)(unsafe.Add(mBase, uint32(v170))))
	goto L56
L55:
	;
	v191 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v170))))
	if v191 != 0 {
		goto L58
	} else {
		goto L59
	}
L56:
	;
	if base.B2i32(v181 == int32(32))|base.B2i32(base.Ui32((v181-int32(9))&int32(255)) < base.Ui32(int32(5))) != 0 {
		v170 = v170 + int32(1)
		goto L54
	} else {
		goto L57
	}
L57:
	;
	goto L55
L58:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v22)+48)) = v170
	F_appendStringInfo(m, v22+int32(92), int32(_a_F_CreateSchemaCommand_7), v22+int32(48))
	mBase = m.M
	v199 = m.ExcPending
	if v199 != 0 {
		goto L6
	} else {
		goto L61
	}
L59:
	;
	goto L60
L60:
	;
	v202 = *(*int32)(unsafe.Add(mBase, uint32(v22)+92))
	F_set_config_option(m, int32(_a_F_CreateSchemaCommand_8), v202, int32(6), int32(13), int32(2), int32(1))
	mBase = m.M
	v208 = m.ExcPending
	if v208 != 0 {
		goto L6
	} else {
		goto L62
	}
L61:
	;
	goto L60
L62:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v22)+112)) = v141
	*(*int32)(unsafe.Add(mBase, uint32(v22)+108)) = int32(2615)
	v212 = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v22)+116)) = v212
	*(*int32)(unsafe.Add(mBase, uint32(v22)+40)) = v212
	v217 = *(*int64)(unsafe.Add(mBase, _c_F_CreateSchemaCommand[6]))
	*(*int64)(unsafe.Add(mBase, uint32(v22)+16)) = v217
	v220 = *(*int32)(unsafe.Add(mBase, _c_F_CreateSchemaCommand[7]))
	*(*int32)(unsafe.Add(mBase, uint32(v22)+24)) = v220
	v222 = *(*int64)(unsafe.Add(mBase, uint32(v22)+108))
	*(*int64)(unsafe.Add(mBase, uint32(v22)+32)) = v222
	F_EventTriggerCollectSimpleCommand(m, v22+int32(32), v22+int32(16), l0)
	mBase = m.M
	v229 = m.ExcPending
	if v229 != 0 {
		goto L6
	} else {
		goto L63
	}
L63:
	;
	v230 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	v231 = m.G0
	v233 = v231 - int32(96)
	m.G0 = v233
	if v230 == int32(0) {
		v485 = v5
		v486 = v5
		v487 = v5
		v488 = v5
		v489 = v5
		v491 = v5
		goto L71
	} else {
		goto L72
	}
L64:
	;
	F_AtEOXact_GUC(m, int32(1), v150)
	mBase = m.M
	v659 = m.ExcPending
	if v659 != 0 {
		goto L6
	} else {
		goto L186
	}
L65:
	;
	if v504 == int32(0) {
		goto L64
	} else {
		goto L178
	}
L66:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v572 = m.ExcPending
	if v572 != 0 {
		goto L6
	} else {
		goto L175
	}
L67:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v557 = m.ExcPending
	if v557 != 0 {
		goto L6
	} else {
		goto L172
	}
L68:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v542 = m.ExcPending
	if v542 != 0 {
		goto L6
	} else {
		goto L169
	}
L69:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v527 = m.ExcPending
	if v527 != 0 {
		goto L6
	} else {
		goto L166
	}
L70:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v512 = m.ExcPending
	if v512 != 0 {
		goto L6
	} else {
		goto L163
	}
L71:
	;
	v494 = F_list_concat(m, int32(0), v485)
	mBase = m.M
	v495 = m.ExcPending
	if v495 != 0 {
		goto L6
	} else {
		goto L157
	}
L72:
	;
	v237 = *(*int32)(unsafe.Add(mBase, uint32(v230)+4))
	if v237 <= int32(0) {
		v485 = v5
		v486 = v5
		v487 = v5
		v488 = v5
		v489 = v5
		v491 = v5
		goto L71
	} else {
		goto L73
	}
L73:
	;
	v251 = v5
	v252 = v5
	v253 = v5
	v254 = v5
	v255 = v5
	v256 = v5
	v257 = v5
	goto L74
L74:
	;
	v259 = *(*int32)(unsafe.Add(mBase, uint32(v230)+12))
	v263 = *(*int32)(unsafe.Add(mBase, uint32(v259+v256<<(uint(int32(2))%32))))
	v264 = *(*int32)(unsafe.Add(mBase, uint32(v263)))
	switch v264 - int32(152) {
	case 0:
		goto L77
	case 1, 2, 3, 4, 5, 6, 7, 9, 10, 11, 12, 13, 14, 15, 16, 17, 18, 19, 20, 21, 22, 23, 24, 25, 26, 27, 28, 30, 31, 32, 33, 34, 35, 36:
		goto L83
	case 8:
		goto L81
	case 29:
		goto L78
	case 37:
		goto L82
	default:
		goto L84
	}
L75:
	;
	v485 = v464
	v486 = v465
	v487 = v466
	v488 = v467
	v489 = v468
	v491 = v469
	goto L71
L76:
	;
	v471 = v256 + int32(1)
	v472 = *(*int32)(unsafe.Add(mBase, uint32(v230)+4))
	if v471 < v472 {
		v251 = v464
		v252 = v465
		v253 = v466
		v254 = v467
		v255 = v468
		v256 = v471
		v257 = v469
		goto L74
	} else {
		goto L156
	}
L77:
	;
	v460 = F_lappend(m, v257, v263)
	mBase = m.M
	v461 = m.ExcPending
	if v461 != 0 {
		goto L6
	} else {
		goto L155
	}
L78:
	;
	v425 = *(*int32)(unsafe.Add(mBase, uint32(v263)+12))
	v426 = *(*int32)(unsafe.Add(mBase, uint32(v425)+8))
	if v426 == int32(0) {
		goto L142
	} else {
		goto L143
	}
L79:
	;
	v390 = *(*int32)(unsafe.Add(mBase, uint32(v263)+8))
	v391 = *(*int32)(unsafe.Add(mBase, uint32(v390)+8))
	if v391 == int32(0) {
		goto L129
	} else {
		goto L130
	}
L80:
	;
	v355 = *(*int32)(unsafe.Add(mBase, uint32(v263)+4))
	v356 = *(*int32)(unsafe.Add(mBase, uint32(v355)+8))
	if v356 == int32(0) {
		goto L116
	} else {
		goto L117
	}
L81:
	;
	v320 = *(*int32)(unsafe.Add(mBase, uint32(v263)+4))
	v321 = *(*int32)(unsafe.Add(mBase, uint32(v320)+8))
	if v321 == int32(0) {
		goto L103
	} else {
		goto L104
	}
L82:
	;
	v285 = *(*int32)(unsafe.Add(mBase, uint32(v263)+4))
	v286 = *(*int32)(unsafe.Add(mBase, uint32(v285)+8))
	if v286 == int32(0) {
		goto L90
	} else {
		goto L91
	}
L83:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v274 = m.ExcPending
	if v274 != 0 {
		goto L6
	} else {
		goto L87
	}
L84:
	;
	if v264 == int32(204) {
		goto L79
	} else {
		goto L85
	}
L85:
	;
	if v264 == int32(230) {
		goto L80
	} else {
		goto L86
	}
L86:
	;
	goto L83
L87:
	;
	v275 = *(*int32)(unsafe.Add(mBase, uint32(v263)))
	*(*int32)(unsafe.Add(mBase, uint32(v233))) = v275
	F_errmsg_internal(m, int32(_a_F_CreateSchemaCommand_9), v233)
	mBase = m.M
	v279 = m.ExcPending
	if v279 != 0 {
		goto L6
	} else {
		goto L88
	}
L88:
	;
	F_errfinish(m, int32(_a_F_CreateSchemaCommand_0), int32(_a_F_CreateSchemaCommand_10), int32(_a_F_CreateSchemaCommand_11))
	mBase = m.M
	v284 = m.ExcPending
	if v284 != 0 {
		goto L6
	} else {
		goto L89
	}
L89:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L90:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v285)+8)) = v59
	v290 = F_lappend(m, v251, v263)
	mBase = m.M
	v291 = m.ExcPending
	if v291 != 0 {
		goto L6
	} else {
		goto L93
	}
L91:
	;
	goto L92
L92:
	;
	v294 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v59))))
	v297 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v286))))
	if base.B2i32(v294 == int32(0))|base.B2i32(v294 != v297) != 0 {
		v315 = v294
		v316 = v297
		goto L95
	} else {
		goto L96
	}
L93:
	;
	v464 = v290
	v465 = v252
	v466 = v253
	v467 = v254
	v468 = v255
	v469 = v257
	goto L76
L94:
	;
	if v315-v316 != 0 {
		goto L70
	} else {
		goto L101
	}
L95:
	;
	goto L94
L96:
	;
	v300 = v59
	v301 = v286
	goto L97
L97:
	;
	v304 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v301)+1)))
	v305 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v300)+1)))
	if v305 == int32(0) {
		v315 = v305
		v316 = v304
		goto L95
	} else {
		goto L99
	}
L98:
	;
	v315 = v305
	v316 = v304
	goto L95
L99:
	;
	v308 = int32(1)
	if v305 == v304 {
		v300 = v300 + v308
		v301 = v301 + v308
		goto L97
	} else {
		goto L100
	}
L100:
	;
	goto L98
L101:
	;
	v318 = F_lappend(m, v251, v263)
	mBase = m.M
	v319 = m.ExcPending
	if v319 != 0 {
		goto L6
	} else {
		goto L102
	}
L102:
	;
	v464 = v318
	v465 = v252
	v466 = v253
	v467 = v254
	v468 = v255
	v469 = v257
	goto L76
L103:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v320)+8)) = v59
	v325 = F_lappend(m, v252, v263)
	mBase = m.M
	v326 = m.ExcPending
	if v326 != 0 {
		goto L6
	} else {
		goto L106
	}
L104:
	;
	goto L105
L105:
	;
	v329 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v59))))
	v332 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v321))))
	if base.B2i32(v329 == int32(0))|base.B2i32(v329 != v332) != 0 {
		v350 = v329
		v351 = v332
		goto L108
	} else {
		goto L109
	}
L106:
	;
	v464 = v251
	v465 = v325
	v466 = v253
	v467 = v254
	v468 = v255
	v469 = v257
	goto L76
L107:
	;
	if v350-v351 != 0 {
		goto L69
	} else {
		goto L114
	}
L108:
	;
	goto L107
L109:
	;
	v335 = v59
	v336 = v321
	goto L110
L110:
	;
	v339 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v336)+1)))
	v340 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v335)+1)))
	if v340 == int32(0) {
		v350 = v340
		v351 = v339
		goto L108
	} else {
		goto L112
	}
L111:
	;
	v350 = v340
	v351 = v339
	goto L108
L112:
	;
	v343 = int32(1)
	if v340 == v339 {
		v335 = v335 + v343
		v336 = v336 + v343
		goto L110
	} else {
		goto L113
	}
L113:
	;
	goto L111
L114:
	;
	v353 = F_lappend(m, v252, v263)
	mBase = m.M
	v354 = m.ExcPending
	if v354 != 0 {
		goto L6
	} else {
		goto L115
	}
L115:
	;
	v464 = v251
	v465 = v353
	v466 = v253
	v467 = v254
	v468 = v255
	v469 = v257
	goto L76
L116:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v355)+8)) = v59
	v360 = F_lappend(m, v253, v263)
	mBase = m.M
	v361 = m.ExcPending
	if v361 != 0 {
		goto L6
	} else {
		goto L119
	}
L117:
	;
	goto L118
L118:
	;
	v364 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v59))))
	v367 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v356))))
	if base.B2i32(v364 == int32(0))|base.B2i32(v364 != v367) != 0 {
		v385 = v364
		v386 = v367
		goto L121
	} else {
		goto L122
	}
L119:
	;
	v464 = v251
	v465 = v252
	v466 = v360
	v467 = v254
	v468 = v255
	v469 = v257
	goto L76
L120:
	;
	if v385-v386 != 0 {
		goto L68
	} else {
		goto L127
	}
L121:
	;
	goto L120
L122:
	;
	v370 = v59
	v371 = v356
	goto L123
L123:
	;
	v374 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v371)+1)))
	v375 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v370)+1)))
	if v375 == int32(0) {
		v385 = v375
		v386 = v374
		goto L121
	} else {
		goto L125
	}
L124:
	;
	v385 = v375
	v386 = v374
	goto L121
L125:
	;
	v378 = int32(1)
	if v375 == v374 {
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
	v388 = F_lappend(m, v253, v263)
	mBase = m.M
	v389 = m.ExcPending
	if v389 != 0 {
		goto L6
	} else {
		goto L128
	}
L128:
	;
	v464 = v251
	v465 = v252
	v466 = v388
	v467 = v254
	v468 = v255
	v469 = v257
	goto L76
L129:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v390)+8)) = v59
	v395 = F_lappend(m, v254, v263)
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
	v399 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v59))))
	v402 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v391))))
	if base.B2i32(v399 == int32(0))|base.B2i32(v399 != v402) != 0 {
		v420 = v399
		v421 = v402
		goto L134
	} else {
		goto L135
	}
L132:
	;
	v464 = v251
	v465 = v252
	v466 = v253
	v467 = v395
	v468 = v255
	v469 = v257
	goto L76
L133:
	;
	if v420-v421 != 0 {
		goto L67
	} else {
		goto L140
	}
L134:
	;
	goto L133
L135:
	;
	v405 = v59
	v406 = v391
	goto L136
L136:
	;
	v409 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v406)+1)))
	v410 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v405)+1)))
	if v410 == int32(0) {
		v420 = v410
		v421 = v409
		goto L134
	} else {
		goto L138
	}
L137:
	;
	v420 = v410
	v421 = v409
	goto L134
L138:
	;
	v413 = int32(1)
	if v410 == v409 {
		v405 = v405 + v413
		v406 = v406 + v413
		goto L136
	} else {
		goto L139
	}
L139:
	;
	goto L137
L140:
	;
	v423 = F_lappend(m, v254, v263)
	mBase = m.M
	v424 = m.ExcPending
	if v424 != 0 {
		goto L6
	} else {
		goto L141
	}
L141:
	;
	v464 = v251
	v465 = v252
	v466 = v253
	v467 = v423
	v468 = v255
	v469 = v257
	goto L76
L142:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v425)+8)) = v59
	v430 = F_lappend(m, v255, v263)
	mBase = m.M
	v431 = m.ExcPending
	if v431 != 0 {
		goto L6
	} else {
		goto L145
	}
L143:
	;
	goto L144
L144:
	;
	v434 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v59))))
	v437 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v426))))
	if base.B2i32(v434 == int32(0))|base.B2i32(v434 != v437) != 0 {
		v455 = v434
		v456 = v437
		goto L147
	} else {
		goto L148
	}
L145:
	;
	v464 = v251
	v465 = v252
	v466 = v253
	v467 = v254
	v468 = v430
	v469 = v257
	goto L76
L146:
	;
	if v455-v456 != 0 {
		goto L66
	} else {
		goto L153
	}
L147:
	;
	goto L146
L148:
	;
	v440 = v59
	v441 = v426
	goto L149
L149:
	;
	v444 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v441)+1)))
	v445 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v440)+1)))
	if v445 == int32(0) {
		v455 = v445
		v456 = v444
		goto L147
	} else {
		goto L151
	}
L150:
	;
	v455 = v445
	v456 = v444
	goto L147
L151:
	;
	v448 = int32(1)
	if v445 == v444 {
		v440 = v440 + v448
		v441 = v441 + v448
		goto L149
	} else {
		goto L152
	}
L152:
	;
	goto L150
L153:
	;
	v458 = F_lappend(m, v255, v263)
	mBase = m.M
	v459 = m.ExcPending
	if v459 != 0 {
		goto L6
	} else {
		goto L154
	}
L154:
	;
	v464 = v251
	v465 = v252
	v466 = v253
	v467 = v254
	v468 = v458
	v469 = v257
	goto L76
L155:
	;
	v464 = v251
	v465 = v252
	v466 = v253
	v467 = v254
	v468 = v255
	v469 = v460
	goto L76
L156:
	;
	goto L75
L157:
	;
	v496 = F_list_concat(m, v494, v486)
	mBase = m.M
	v497 = m.ExcPending
	if v497 != 0 {
		goto L6
	} else {
		goto L158
	}
L158:
	;
	v498 = F_list_concat(m, v496, v487)
	mBase = m.M
	v499 = m.ExcPending
	if v499 != 0 {
		goto L6
	} else {
		goto L159
	}
L159:
	;
	v500 = F_list_concat(m, v498, v488)
	mBase = m.M
	v501 = m.ExcPending
	if v501 != 0 {
		goto L6
	} else {
		goto L160
	}
L160:
	;
	v502 = F_list_concat(m, v500, v489)
	mBase = m.M
	v503 = m.ExcPending
	if v503 != 0 {
		goto L6
	} else {
		goto L161
	}
L161:
	;
	v504 = F_list_concat(m, v502, v491)
	mBase = m.M
	v505 = m.ExcPending
	if v505 != 0 {
		goto L6
	} else {
		goto L162
	}
L162:
	;
	m.G0 = v233 + int32(96)
	goto L65
L163:
	;
	F_errcode(m, int32(84279428))
	mBase = m.M
	v515 = m.ExcPending
	if v515 != 0 {
		goto L6
	} else {
		goto L164
	}
L164:
	;
	v516 = *(*int32)(unsafe.Add(mBase, uint32(v285)+8))
	*(*int32)(unsafe.Add(mBase, uint32(v233)+20)) = v59
	*(*int32)(unsafe.Add(mBase, uint32(v233)+16)) = v516
	F_errmsg(m, int32(_a_F_CreateSchemaCommand_12), v233+int32(16))
	mBase = m.M
	v523 = m.ExcPending
	if v523 != 0 {
		goto L6
	} else {
		goto L165
	}
L165:
	;
	goto L8
L166:
	;
	F_errcode(m, int32(84279428))
	mBase = m.M
	v530 = m.ExcPending
	if v530 != 0 {
		goto L6
	} else {
		goto L167
	}
L167:
	;
	v531 = *(*int32)(unsafe.Add(mBase, uint32(v320)+8))
	*(*int32)(unsafe.Add(mBase, uint32(v233)+36)) = v59
	*(*int32)(unsafe.Add(mBase, uint32(v233)+32)) = v531
	F_errmsg(m, int32(_a_F_CreateSchemaCommand_12), v233+int32(32))
	mBase = m.M
	v538 = m.ExcPending
	if v538 != 0 {
		goto L6
	} else {
		goto L168
	}
L168:
	;
	goto L8
L169:
	;
	F_errcode(m, int32(84279428))
	mBase = m.M
	v545 = m.ExcPending
	if v545 != 0 {
		goto L6
	} else {
		goto L170
	}
L170:
	;
	v546 = *(*int32)(unsafe.Add(mBase, uint32(v355)+8))
	*(*int32)(unsafe.Add(mBase, uint32(v233)+52)) = v59
	*(*int32)(unsafe.Add(mBase, uint32(v233)+48)) = v546
	F_errmsg(m, int32(_a_F_CreateSchemaCommand_12), v233+int32(48))
	mBase = m.M
	v553 = m.ExcPending
	if v553 != 0 {
		goto L6
	} else {
		goto L171
	}
L171:
	;
	goto L8
L172:
	;
	F_errcode(m, int32(84279428))
	mBase = m.M
	v560 = m.ExcPending
	if v560 != 0 {
		goto L6
	} else {
		goto L173
	}
L173:
	;
	v561 = *(*int32)(unsafe.Add(mBase, uint32(v390)+8))
	*(*int32)(unsafe.Add(mBase, uint32(v233)+68)) = v59
	*(*int32)(unsafe.Add(mBase, uint32(v233)+64)) = v561
	F_errmsg(m, int32(_a_F_CreateSchemaCommand_12), v233-int32(-64))
	mBase = m.M
	v568 = m.ExcPending
	if v568 != 0 {
		goto L6
	} else {
		goto L174
	}
L174:
	;
	goto L8
L175:
	;
	F_errcode(m, int32(84279428))
	mBase = m.M
	v575 = m.ExcPending
	if v575 != 0 {
		goto L6
	} else {
		goto L176
	}
L176:
	;
	v576 = *(*int32)(unsafe.Add(mBase, uint32(v425)+8))
	*(*int32)(unsafe.Add(mBase, uint32(v233)+84)) = v59
	*(*int32)(unsafe.Add(mBase, uint32(v233)+80)) = v576
	F_errmsg(m, int32(_a_F_CreateSchemaCommand_12), v233+int32(80))
	mBase = m.M
	v583 = m.ExcPending
	if v583 != 0 {
		goto L6
	} else {
		goto L177
	}
L177:
	;
	goto L8
L178:
	;
	v586 = *(*int32)(unsafe.Add(mBase, uint32(v504)+4))
	if v586 <= int32(0) {
		goto L64
	} else {
		goto L179
	}
L179:
	;
	v599 = int32(0)
	goto L180
L180:
	;
	v608 = *(*int32)(unsafe.Add(mBase, uint32(v504)+12))
	v612 = *(*int32)(unsafe.Add(mBase, uint32(v608+v599<<(uint(int32(2))%32))))
	v614 = F_palloc0(m, int32(104))
	mBase = m.M
	v615 = m.ExcPending
	if v615 != 0 {
		goto L6
	} else {
		goto L182
	}
L181:
	;
	goto L64
L182:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v614)+96)) = l3
	*(*int32)(unsafe.Add(mBase, uint32(v614)+92)) = l2
	*(*int32)(unsafe.Add(mBase, uint32(v614)+88)) = v612
	v619 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v614)+26)) = uint8(v619)
	*(*int64)(unsafe.Add(mBase, uint32(v614))) = int64(25769804106)
	v628 = *(*int32)(unsafe.Add(mBase, _c_F_CreateSchemaCommand[8]))
	F_ProcessUtility(m, v614, l1, v619, int32(3), v619, v619, v628, v619)
	mBase = m.M
	v631 = m.ExcPending
	if v631 != 0 {
		goto L6
	} else {
		goto L183
	}
L183:
	;
	F_CommandCounterIncrement(m)
	mBase = m.M
	v633 = m.ExcPending
	if v633 != 0 {
		goto L6
	} else {
		goto L184
	}
L184:
	;
	v635 = v599 + int32(1)
	v636 = *(*int32)(unsafe.Add(mBase, uint32(v504)+4))
	if v635 < v636 {
		v599 = v635
		goto L180
	} else {
		goto L185
	}
L185:
	;
	goto L181
L186:
	;
	v660 = *(*int32)(unsafe.Add(mBase, uint32(v22)+124))
	v661 = *(*int32)(unsafe.Add(mBase, uint32(v22)+120))
	*(*int32)(unsafe.Add(mBase, _c_F_CreateSchemaCommand[2])) = v661
	*(*int32)(unsafe.Add(mBase, _c_F_CreateSchemaCommand[1])) = v660
	goto L187
L187:
	;
	goto L33
L188:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v22))) = v42
	F_errmsg_internal(m, int32(_a_F_CreateSchemaCommand_13), v22)
	mBase = m.M
	v695 = m.ExcPending
	if v695 != 0 {
		goto L6
	} else {
		goto L189
	}
L189:
	;
	F_errfinish(m, int32(_a_F_CreateSchemaCommand_4), int32(85), int32(_a_F_CreateSchemaCommand_5))
	mBase = m.M
	v700 = m.ExcPending
	if v700 != 0 {
		goto L6
	} else {
		goto L190
	}
L190:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L191:
	;
	F_errcode(m, int32(151818372))
	mBase = m.M
	v707 = m.ExcPending
	if v707 != 0 {
		goto L6
	} else {
		goto L192
	}
L192:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v22)+80)) = v59
	F_errmsg(m, int32(_a_F_CreateSchemaCommand_14), v22+int32(80))
	mBase = m.M
	v713 = m.ExcPending
	if v713 != 0 {
		goto L6
	} else {
		goto L193
	}
L193:
	;
	F_errdetail(m, int32(_a_F_CreateSchemaCommand_15), int32(0))
	mBase = m.M
	v717 = m.ExcPending
	if v717 != 0 {
		goto L6
	} else {
		goto L194
	}
L194:
	;
	F_errfinish(m, int32(_a_F_CreateSchemaCommand_4), int32(110), int32(_a_F_CreateSchemaCommand_5))
	mBase = m.M
	v722 = m.ExcPending
	if v722 != 0 {
		goto L6
	} else {
		goto L195
	}
L195:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L196:
	;
	base.Wasm_trap_unreachable()
	for {
	}
}
func F_GetSchemaPublications(m *base.Module, l0 int32) int32 {
	var v3 int32
	_ = v3
	var v6 int32
	_ = v6
	v3 = Fn13829(m, l0, int32(50))
	v6 = m.ExcPending
	if v6 != 0 {
		return int32(0)
	} else {
		return v3
	}
}
func F_has_schema_privilege_id(m *base.Module, l0 int32) int32 {
	var v4 int32
	_ = v4
	var v7 int32
	_ = v7
	v4 = Fn13913(m, l0, int32(_a_F_has_schema_privilege_id_0), int32(2615))
	v7 = m.ExcPending
	if v7 != 0 {
		return int32(0)
	} else {
		return v4
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
			v13 = *(*int32)(unsafe.Add(mBase, _c_F_has_schema_privilege_name[0]))
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
					v21 = F_convert_any_priv_string(m, v10, int32(_a_F_has_schema_privilege_name_0))
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
	var v4 int32
	_ = v4
	var v7 int32
	_ = v7
	v4 = Fn13915(m, l0, int32(_a_F_has_schema_privilege_name_id_0), int32(2615))
	v7 = m.ExcPending
	if v7 != 0 {
		return int32(0)
	} else {
		return v4
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
						v22 = F_convert_any_priv_string(m, v11, int32(_a_F_has_schema_privilege_name_name_0))
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
	var v33 int32
	_ = v33
	var v35 int32
	_ = v35
	var v39 int32
	_ = v39
	var v40 int32
	_ = v40
	var v41 int32
	_ = v41
	var v42 int32
	_ = v42
	var v45 int32
	_ = v45
	var v50 int32
	_ = v50
	var v54 int32
	_ = v54
	var v55 int32
	_ = v55
	var v59 int32
	_ = v59
	var v61 int32
	_ = v61
	var v62 int32
	_ = v62
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
	var v70 int32
	_ = v70
	var v72 int32
	_ = v72
	var v73 int32
	_ = v73
	var v80 int32
	_ = v80
	var v81 int32
	_ = v81
	var v82 int32
	_ = v82
	var v84 int32
	_ = v84
	var v85 int32
	_ = v85
	var v87 int32
	_ = v87
	var v90 int32
	_ = v90
	var v91 int32
	_ = v91
	var v92 int32
	_ = v92
	var v93 int32
	_ = v93
	var v94 int32
	_ = v94
	var v96 int32
	_ = v96
	var v97 int32
	_ = v97
	var v98 int32
	_ = v98
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
	F_appendStringInfoString(m, v11, int32(_a_F_schema_to_xmlschema_internal_0))
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
	F_appendStringInfo(m, v11, int32(_a_F_schema_to_xmlschema_internal_1), v9+int32(16))
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
	F_appendStringInfoString(m, v11, int32(_a_F_schema_to_xmlschema_internal_2))
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
	v33 = v9 + int32(32)
	F_initStringInfo(m, v33)
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
	F_appendStringInfo(m, v33, int32(_a_F_schema_to_xmlschema_internal_3), v9)
	mBase = m.M
	v39 = m.ExcPending
	if v39 != 0 {
		goto L1
	} else {
		goto L12
	}
L12:
	;
	v40 = *(*int32)(unsafe.Add(mBase, uint32(v9)+32))
	v41 = F_query_to_oid_list(m, v40)
	mBase = m.M
	v42 = m.ExcPending
	if v42 != 0 {
		goto L1
	} else {
		goto L14
	}
L13:
	;
	v81 = F_map_sql_typecoll_to_xmlschema_types(m, v80)
	mBase = m.M
	v82 = m.ExcPending
	if v82 != 0 {
		goto L1
	} else {
		goto L24
	}
L14:
	;
	if v41 == int32(0) {
		v80 = v3
		goto L13
	} else {
		goto L15
	}
L15:
	;
	v45 = *(*int32)(unsafe.Add(mBase, uint32(v41)+4))
	if v45 <= int32(0) {
		v80 = v3
		goto L13
	} else {
		goto L16
	}
L16:
	;
	v50 = int32(0)
	v54 = v3
	goto L17
L17:
	;
	v55 = *(*int32)(unsafe.Add(mBase, uint32(v41)+12))
	v59 = *(*int32)(unsafe.Add(mBase, uint32(v55+v50<<(uint(int32(2))%32))))
	v61 = F_table_open(m, v59, int32(1))
	mBase = m.M
	v62 = m.ExcPending
	if v62 != 0 {
		goto L1
	} else {
		goto L19
	}
L18:
	;
	v80 = v66
	goto L13
L19:
	;
	v63 = *(*int32)(unsafe.Add(mBase, uint32(v61)+52))
	v64 = F_CreateTupleDescCopy(m, v63)
	mBase = m.M
	v65 = m.ExcPending
	if v65 != 0 {
		goto L1
	} else {
		goto L20
	}
L20:
	;
	v66 = F_lappend(m, v54, v64)
	mBase = m.M
	v67 = m.ExcPending
	if v67 != 0 {
		goto L1
	} else {
		goto L21
	}
L21:
	;
	F_relation_close(m, v61, int32(0))
	mBase = m.M
	v70 = m.ExcPending
	if v70 != 0 {
		goto L1
	} else {
		goto L22
	}
L22:
	;
	v72 = v50 + int32(1)
	v73 = *(*int32)(unsafe.Add(mBase, uint32(v41)+4))
	if v72 < v73 {
		v50 = v72
		v54 = v66
		goto L17
	} else {
		goto L23
	}
L23:
	;
	goto L18
L24:
	;
	F_appendStringInfoString(m, v11, v81)
	mBase = m.M
	v84 = m.ExcPending
	if v84 != 0 {
		goto L1
	} else {
		goto L25
	}
L25:
	;
	v85 = m.G0
	v87 = v85 - int32(16)
	m.G0 = v87
	v90 = *(*int32)(unsafe.Add(mBase, _c_F_schema_to_xmlschema_internal[0]))
	v91 = F_get_database_name(m, v90)
	mBase = m.M
	v92 = m.ExcPending
	if v92 != 0 {
		goto L1
	} else {
		goto L26
	}
L26:
	;
	v93 = F_get_namespace_name(m, v14)
	mBase = m.M
	v94 = m.ExcPending
	if v94 != 0 {
		goto L1
	} else {
		goto L27
	}
L27:
	;
	F_initStringInfo(m, v87)
	mBase = m.M
	v96 = m.ExcPending
	if v96 != 0 {
		goto L1
	} else {
		goto L28
	}
L28:
	;
	v97 = F_map_sql_identifier_to_xml_name(m)
	mBase = m.M
	v98 = m.ExcPending
	if v98 != 0 {
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
