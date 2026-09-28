package p5

import (
	base "github.com/shibukawa/pgmem/internal/aot/pgaot/base"
	"unsafe"
)

func F_SetStartupBufferPinWaitBufId(m *base.Module, l0 int32) {
	mBase := m.M
	_ = mBase
	var v3 int32
	_ = v3
	v3 = *(*int32)(unsafe.Add(mBase, _c_F_SetStartupBufferPinWaitBufId[0]))
	*(*int32)(unsafe.Add(mBase, uint32(v3)+76)) = l0
	return
}
func F_StartupDecodingContext(m *base.Module, l0 int32, l1 int64, l2 int32, l3 int32, l4 int32, l5 int32, l6 int32, l7 int32, l8 int32, l9 int32, l10 int32) int32 {
	mBase := m.M
	_ = mBase
	var v4 int32
	_ = v4
	var v5 int32
	_ = v5
	var v6 int32
	_ = v6
	var v22 int32
	_ = v22
	var v24 int32
	_ = v24
	var v27 int32
	_ = v27
	var v29 int32
	_ = v29
	var v34 int32
	_ = v34
	var v37 int32
	_ = v37
	var v38 int32
	_ = v38
	var v39 int32
	_ = v39
	var v43 int32
	_ = v43
	var v44 int32
	_ = v44
	var v49 int32
	_ = v49
	var v50 int32
	_ = v50
	var v53 int32
	_ = v53
	var v56 int32
	_ = v56
	var v59 int32
	_ = v59
	var v60 int32
	_ = v60
	var v63 int32
	_ = v63
	var v64 int32
	_ = v64
	var v67 int32
	_ = v67
	var v74 int32
	_ = v74
	var v75 int32
	_ = v75
	var v80 int32
	_ = v80
	var v84 int32
	_ = v84
	var v87 int32
	_ = v87
	var v94 int32
	_ = v94
	var v98 int32
	_ = v98
	var v99 int32
	_ = v99
	var v104 int32
	_ = v104
	var v106 int32
	_ = v106
	var v109 int32
	_ = v109
	var v112 int32
	_ = v112
	var v115 int32
	_ = v115
	var v116 int32
	_ = v116
	var v119 int32
	_ = v119
	var v120 int32
	_ = v120
	var v125 int32
	_ = v125
	var v126 int32
	_ = v126
	var v129 int32
	_ = v129
	var v136 int32
	_ = v136
	var v141 int32
	_ = v141
	var v142 int32
	_ = v142
	var v144 int32
	_ = v144
	var v147 int32
	_ = v147
	var v150 int32
	_ = v150
	var v153 int32
	_ = v153
	var v156 int32
	_ = v156
	var v157 int32
	_ = v157
	var v164 int32
	_ = v164
	var v182 int32
	_ = v182
	var v185 int32
	_ = v185
	var v188 int32
	_ = v188
	var v191 int32
	_ = v191
	var v192 int32
	_ = v192
	var v195 int32
	_ = v195
	var v196 int32
	_ = v196
	var v199 int32
	_ = v199
	var v206 int32
	_ = v206
	var v207 int32
	_ = v207
	var v212 int32
	_ = v212
	var v215 int32
	_ = v215
	var v217 int32
	_ = v217
	var v240 int32
	_ = v240
	var v242 int32
	_ = v242
	var v243 int32
	_ = v243
	var v249 int32
	_ = v249
	var v250 int32
	_ = v250
	var v253 int32
	_ = v253
	var v256 int32
	_ = v256
	var v281 int32
	_ = v281
	var v282 int32
	_ = v282
	var v288 int32
	_ = v288
	var v292 int32
	_ = v292
	var v293 int32
	_ = v293
	var v295 int32
	_ = v295
	var v296 int32
	_ = v296
	var v298 int32
	_ = v298
	var v301 int32
	_ = v301
	var v302 int32
	_ = v302
	var v303 int32
	_ = v303
	var v307 int32
	_ = v307
	var v311 int32
	_ = v311
	var v316 int32
	_ = v316
	var v317 int32
	_ = v317
	var v318 int32
	_ = v318
	var v322 int32
	_ = v322
	var v324 int32
	_ = v324
	var v327 int32
	_ = v327
	var v332 int32
	_ = v332
	var v333 int32
	_ = v333
	var v335 int32
	_ = v335
	var v336 int32
	_ = v336
	var v337 int64
	_ = v337
	var v353 int32
	_ = v353
	var v354 int32
	_ = v354
	var v359 int32
	_ = v359
	var v360 int32
	_ = v360
	var v363 int32
	_ = v363
	var v366 int32
	_ = v366
	var v367 int32
	_ = v367
	var v371 int32
	_ = v371
	var v376 int32
	_ = v376
	var v377 int32
	_ = v377
	var v378 int32
	_ = v378
	var v380 int64
	_ = v380
	var v387 int32
	_ = v387
	var v388 int32
	_ = v388
	var v394 int32
	_ = v394
	var v400 int32
	_ = v400
	var v404 int32
	_ = v404
	var v408 int32
	_ = v408
	var v412 int32
	_ = v412
	var v416 int32
	_ = v416
	var v421 int64
	_ = v421
	var v423 int32
	_ = v423
	var v428 int32
	_ = v428
	var v429 int32
	_ = v429
	var v430 int32
	_ = v430
	var v431 int32
	_ = v431
	var v435 int32
	_ = v435
	var v436 int32
	_ = v436
	var v445 int32
	_ = v445
	var v446 int32
	_ = v446
	var v449 int32
	_ = v449
	var v460 int32
	_ = v460
	var v462 int32
	_ = v462
	var v465 int32
	_ = v465
	var v468 int32
	_ = v468
	var v471 int32
	_ = v471
	var v474 int32
	_ = v474
	var v479 int32
	_ = v479
	var v481 int32
	_ = v481
	var v483 int32
	_ = v483
	var v485 int32
	_ = v485
	var v487 int32
	_ = v487
	var v489 int32
	_ = v489
	var v490 int32
	_ = v490
	var v493 int32
	_ = v493
	var v495 int32
	_ = v495
	var v498 int32
	_ = v498
	var v501 int32
	_ = v501
	var v504 int32
	_ = v504
	var v507 int32
	_ = v507
	var v510 int32
	_ = v510
	var v513 int32
	_ = v513
	var v516 int32
	_ = v516
	var v519 int32
	_ = v519
	var v520 int32
	_ = v520
	var v521 int32
	_ = v521
	var v522 int32
	_ = v522
	var v523 int32
	_ = v523
	var v524 int32
	_ = v524
	var v527 int32
	_ = v527
	var v529 int32
	_ = v529
	var v532 int32
	_ = v532
	var v535 int32
	_ = v535
	var v538 int32
	_ = v538
	var v541 int32
	_ = v541
	var v544 int32
	_ = v544
	var v545 int32
	_ = v545
	var v561 int32
	_ = v561
	var v565 int32
	_ = v565
	var v570 int32
	_ = v570
	var v574 int32
	_ = v574
	var v578 int32
	_ = v578
	var v583 int32
	_ = v583
	var v587 int32
	_ = v587
	var v591 int32
	_ = v591
	var v596 int32
	_ = v596
	var v600 int32
	_ = v600
	var v604 int32
	_ = v604
	var v609 int32
	_ = v609
	var v613 int32
	_ = v613
	var v616 int32
	_ = v616
	var v620 int32
	_ = v620
	var v623 int32
	_ = v623
	var v624 int32
	_ = v624
	var v629 int32
	_ = v629
	var v636 int32
	_ = v636
	var v652 int32
	_ = v652
	var v654 int32
	_ = v654
	var v679 int32
	_ = v679
	var v682 int32
	_ = v682
	var v688 int32
	_ = v688
	var v692 int32
	_ = v692
	var v698 int32
	_ = v698
	var v705 int32
	_ = v705
	var v710 int32
	_ = v710
	v4 = l3
	v5 = l4
	v6 = l5
	v22 = m.G0
	v24 = v22 - int32(96)
	m.G0 = v24
	v27 = *(*int32)(unsafe.Add(mBase, _c_F_StartupDecodingContext[0]))
	v29 = *(*int32)(unsafe.Add(mBase, _c_F_StartupDecodingContext[1]))
	v34 = F_AllocSetContextCreateInternal(m, v29, int32(_a_F_StartupDecodingContext_0), int32(0), int32(_a_F_StartupDecodingContext_1), int32(_a_F_StartupDecodingContext_2))
	mBase = m.M
	v37 = m.ExcPending
	if v37 != 0 {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	return int32(0)
L2:
	;
	v38 = int32(_a_F_StartupDecodingContext_3)
	v39 = *(*int32)(unsafe.Add(mBase, _c_F_StartupDecodingContext[1]))
	*(*int32)(unsafe.Add(mBase, _c_F_StartupDecodingContext[1])) = v34
	v43 = F_palloc0(m, int32(168))
	mBase = m.M
	v44 = m.ExcPending
	if v44 != 0 {
		goto L1
	} else {
		goto L3
	}
L3:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v43))) = v34
	if v5 == int32(0) {
		goto L11
	} else {
		goto L12
	}
L4:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v679 = m.ExcPending
	if v679 != 0 {
		goto L1
	} else {
		goto L128
	}
L5:
	;
	F_list_free(m, v636)
	mBase = m.M
	v652 = m.ExcPending
	if v652 != 0 {
		goto L1
	} else {
		goto L126
	}
L6:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v613 = m.ExcPending
	if v613 != 0 {
		goto L1
	} else {
		goto L121
	}
L7:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v600 = m.ExcPending
	if v600 != 0 {
		goto L1
	} else {
		goto L118
	}
L8:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v587 = m.ExcPending
	if v587 != 0 {
		goto L1
	} else {
		goto L115
	}
L9:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v574 = m.ExcPending
	if v574 != 0 {
		goto L1
	} else {
		goto L112
	}
L10:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v561 = m.ExcPending
	if v561 != 0 {
		goto L1
	} else {
		goto L109
	}
L11:
	;
	v49 = v27 + int32(137)
	v50 = int32(_a_F_StartupDecodingContext_4)
	v53 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v49))))
	v56 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_StartupDecodingContext[2])))
	if base.B2i32(v53 == int32(0))|base.B2i32(v53 != v56) != 0 {
		v74 = v53
		v75 = v56
		goto L16
	} else {
		goto L17
	}
L12:
	;
	goto L13
L13:
	;
	v281 = *(*int32)(unsafe.Add(mBase, _c_F_StartupDecodingContext[3]))
	v282 = *(*int32)(unsafe.Add(mBase, uint32(v281)+24))
	goto L76
L14:
	;
	v240 = int32(0)
	v242 = F_load_external_function(m, v49, int32(_a_F_StartupDecodingContext_5), v240, v240)
	mBase = m.M
	v243 = m.ExcPending
	if v243 != 0 {
		goto L1
	} else {
		goto L70
	}
L15:
	;
	if v74-v75 == int32(0) {
		goto L22
	} else {
		goto L23
	}
L16:
	;
	goto L15
L17:
	;
	v59 = v49
	v60 = v50
	goto L18
L18:
	;
	v63 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v60)+1)))
	v64 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v59)+1)))
	if v64 == int32(0) {
		v74 = v64
		v75 = v63
		goto L16
	} else {
		goto L20
	}
L19:
	;
	v74 = v64
	v75 = v63
	goto L16
L20:
	;
	v67 = int32(1)
	if v64 == v63 {
		v59 = v59 + v67
		v60 = v60 + v67
		goto L18
	} else {
		goto L21
	}
L21:
	;
	goto L19
L22:
	;
	if l6 != 0 {
		goto L25
	} else {
		goto L26
	}
L23:
	;
	goto L24
L24:
	;
	v106 = *(*int32)(unsafe.Add(mBase, _c_F_StartupDecodingContext[4]))
	if v106 == int32(0) {
		goto L4
	} else {
		goto L34
	}
L25:
	;
	v80 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_StartupDecodingContext[5])))
	if v80 != 0 {
		goto L14
	} else {
		goto L28
	}
L26:
	;
	goto L27
L27:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v84 = m.ExcPending
	if v84 != 0 {
		goto L1
	} else {
		goto L29
	}
L28:
	;
	goto L27
L29:
	;
	F_errcode(m, int32(1088))
	mBase = m.M
	v87 = m.ExcPending
	if v87 != 0 {
		goto L1
	} else {
		goto L30
	}
L30:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v24)+16)) = int32(_a_F_StartupDecodingContext_4)
	F_errmsg(m, int32(_a_F_StartupDecodingContext_6), v24+int32(16))
	mBase = m.M
	v94 = m.ExcPending
	if v94 != 0 {
		goto L1
	} else {
		goto L31
	}
L31:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v24))) = int32(_a_F_StartupDecodingContext_7)
	v98 = F_errdetail(m, int32(_a_F_StartupDecodingContext_8), v24)
	mBase = m.M
	v99 = m.ExcPending
	if v99 != 0 {
		goto L1
	} else {
		goto L32
	}
L32:
	;
	F_errfinish(m, int32(_a_F_StartupDecodingContext_9), int32(200), int32(_a_F_StartupDecodingContext_10))
	mBase = m.M
	v104 = m.ExcPending
	if v104 != 0 {
		goto L1
	} else {
		goto L33
	}
L33:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L34:
	;
	v109 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v106))))
	if v109 == int32(0) {
		goto L4
	} else {
		goto L35
	}
L35:
	;
	v112 = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v24)+92)) = v112
	v115 = F_pstrdup(m, v106)
	mBase = m.M
	v116 = m.ExcPending
	if v116 != 0 {
		goto L1
	} else {
		goto L36
	}
L36:
	;
	v119 = F_SplitGUCList(m, v115, v24+int32(92))
	mBase = m.M
	v120 = m.ExcPending
	if v120 != 0 {
		goto L1
	} else {
		goto L37
	}
L37:
	;
	if v119 == int32(0) {
		goto L38
	} else {
		goto L39
	}
L38:
	;
	v125 = F_errstart(m, int32(15), int32(0))
	mBase = m.M
	v126 = m.ExcPending
	if v126 != 0 {
		goto L1
	} else {
		goto L41
	}
L39:
	;
	goto L40
L40:
	;
	v147 = *(*int32)(unsafe.Add(mBase, uint32(v24)+92))
	if v147 == int32(0) {
		v636 = v112
		goto L5
	} else {
		goto L49
	}
L41:
	;
	if v125 != 0 {
		goto L42
	} else {
		goto L43
	}
L42:
	;
	F_errcode(m, int32(16801924))
	mBase = m.M
	v129 = m.ExcPending
	if v129 != 0 {
		goto L1
	} else {
		goto L45
	}
L43:
	;
	goto L44
L44:
	;
	v142 = *(*int32)(unsafe.Add(mBase, uint32(v24)+92))
	F_list_free(m, v142)
	mBase = m.M
	v144 = m.ExcPending
	if v144 != 0 {
		goto L1
	} else {
		goto L48
	}
L45:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v24)+80)) = int32(_a_F_StartupDecodingContext_11)
	F_errmsg(m, int32(_a_F_StartupDecodingContext_12), v24+int32(80))
	mBase = m.M
	v136 = m.ExcPending
	if v136 != 0 {
		goto L1
	} else {
		goto L46
	}
L46:
	;
	F_errfinish(m, int32(_a_F_StartupDecodingContext_9), int32(220), int32(_a_F_StartupDecodingContext_10))
	mBase = m.M
	v141 = m.ExcPending
	if v141 != 0 {
		goto L1
	} else {
		goto L47
	}
L47:
	;
	goto L44
L48:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v24)+92)) = int32(0)
	v636 = v112
	goto L5
L49:
	;
	v150 = *(*int32)(unsafe.Add(mBase, uint32(v147)+4))
	if v150 <= int32(0) {
		goto L50
	} else {
		goto L51
	}
L50:
	;
	v636 = v147
	goto L5
L51:
	;
	goto L52
L52:
	;
	v153 = int32(0)
	if v153 < v150 {
		goto L53
	} else {
		goto L54
	}
L53:
	;
	v156 = v150
	goto L55
L54:
	;
	v156 = v153
	goto L55
L55:
	;
	v157 = *(*int32)(unsafe.Add(mBase, uint32(v147)+12))
	v164 = v112
	goto L57
L56:
	;
	F_list_free(m, v147)
	mBase = m.M
	v215 = m.ExcPending
	if v215 != 0 {
		goto L1
	} else {
		goto L68
	}
L57:
	;
	v182 = *(*int32)(unsafe.Add(mBase, uint32(v157+v164<<(uint(int32(2))%32))))
	v185 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v182))))
	v188 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v49))))
	if base.B2i32(v185 == int32(0))|base.B2i32(v185 != v188) != 0 {
		v206 = v185
		v207 = v188
		goto L60
	} else {
		goto L61
	}
L58:
	;
	v636 = v147
	goto L5
L59:
	;
	if v206-v207 == int32(0) {
		goto L56
	} else {
		goto L66
	}
L60:
	;
	goto L59
L61:
	;
	v191 = v182
	v192 = v49
	goto L62
L62:
	;
	v195 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v192)+1)))
	v196 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v191)+1)))
	if v196 == int32(0) {
		v206 = v196
		v207 = v195
		goto L60
	} else {
		goto L64
	}
L63:
	;
	v206 = v196
	v207 = v195
	goto L60
L64:
	;
	v199 = int32(1)
	if v196 == v195 {
		v191 = v191 + v199
		v192 = v192 + v199
		goto L62
	} else {
		goto L65
	}
L65:
	;
	goto L63
L66:
	;
	v212 = v164 + int32(1)
	if v212 != v156 {
		v164 = v212
		goto L57
	} else {
		goto L67
	}
L67:
	;
	goto L58
L68:
	;
	F_pfree(m, v115)
	mBase = m.M
	v217 = m.ExcPending
	if v217 != 0 {
		goto L1
	} else {
		goto L69
	}
L69:
	;
	goto L14
L70:
	;
	if v242 == int32(0) {
		goto L10
	} else {
		goto L71
	}
L71:
	;
	m.T0[v242].(func(*base.Module, int32))(m, v43+int32(24))
	mBase = m.M
	v249 = m.ExcPending
	if v249 != 0 {
		goto L1
	} else {
		goto L72
	}
L72:
	;
	v250 = *(*int32)(unsafe.Add(mBase, uint32(v43)+28))
	if v250 == int32(0) {
		goto L9
	} else {
		goto L73
	}
L73:
	;
	v253 = *(*int32)(unsafe.Add(mBase, uint32(v43)+32))
	if v253 == int32(0) {
		goto L8
	} else {
		goto L74
	}
L74:
	;
	v256 = *(*int32)(unsafe.Add(mBase, uint32(v43)+40))
	if v256 == int32(0) {
		goto L7
	} else {
		goto L75
	}
L75:
	;
	goto L13
L76:
	;
	if base.B2i32(v282 != int32(0)) == int32(0) {
		goto L77
	} else {
		goto L78
	}
L77:
	;
	v288 = *(*int32)(unsafe.Add(mBase, _c_F_StartupDecodingContext[6]))
	v292 = F_LWLockAcquire(m, v288+int32(512), int32(0))
	mBase = m.M
	v293 = m.ExcPending
	if v293 != 0 {
		goto L1
	} else {
		goto L80
	}
L78:
	;
	goto L79
L79:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v43)+4)) = v27
	v316 = *(*int32)(unsafe.Add(mBase, _c_F_StartupDecodingContext[7]))
	v317 = F_XLogReaderAllocate(m, v316, l7, v43)
	mBase = m.M
	v318 = m.ExcPending
	if v318 != 0 {
		goto L1
	} else {
		goto L82
	}
L80:
	;
	v295 = *(*int32)(unsafe.Add(mBase, _c_F_StartupDecodingContext[8]))
	v296 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v295)+36)))
	v298 = v296 | int32(16)
	*(*uint8)(unsafe.Add(mBase, uint32(v295)+36)) = uint8(v298)
	v301 = *(*int32)(unsafe.Add(mBase, _c_F_StartupDecodingContext[9]))
	v302 = *(*int32)(unsafe.Add(mBase, uint32(v301)+12))
	v303 = *(*int32)(unsafe.Add(mBase, uint32(v295)+32))
	*(*uint8)(unsafe.Add(mBase, uint32(v302+v303))) = uint8(v298)
	v307 = *(*int32)(unsafe.Add(mBase, _c_F_StartupDecodingContext[6]))
	F_LWLockRelease(m, v307+int32(512))
	mBase = m.M
	v311 = m.ExcPending
	if v311 != 0 {
		goto L1
	} else {
		goto L81
	}
L81:
	;
	goto L79
L82:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v43)+8)) = v317
	if v317 == int32(0) {
		goto L6
	} else {
		goto L83
	}
L83:
	;
	v322 = m.G0
	v324 = v322 - int32(48)
	m.G0 = v324
	v327 = *(*int32)(unsafe.Add(mBase, _c_F_StartupDecodingContext[1]))
	v332 = F_AllocSetContextCreateInternal(m, v327, int32(_a_F_StartupDecodingContext_13), int32(0), int32(_a_F_StartupDecodingContext_1), int32(_a_F_StartupDecodingContext_2))
	mBase = m.M
	v333 = m.ExcPending
	if v333 != 0 {
		goto L1
	} else {
		goto L84
	}
L84:
	;
	v335 = F_MemoryContextAlloc(m, v332, int32(232))
	mBase = m.M
	v336 = m.ExcPending
	if v336 != 0 {
		goto L1
	} else {
		goto L85
	}
L85:
	;
	v337 = int64(0)
	*(*int64)(unsafe.Add(mBase, uint32(v324)+40)) = v337
	*(*int64)(unsafe.Add(mBase, uint32(v324)+32)) = v337
	*(*int64)(unsafe.Add(mBase, uint32(v324)+24)) = v337
	*(*int64)(unsafe.Add(mBase, uint32(v324)+16)) = v337
	*(*int64)(unsafe.Add(mBase, uint32(v324)+8)) = v337
	*(*int64)(unsafe.Add(mBase, uint32(v324))) = v337
	*(*int32)(unsafe.Add(mBase, uint32(v335)+120)) = v332
	v353 = F_SlabContextCreate(m, v332, int32(_a_F_StartupDecodingContext_14), int32(_a_F_StartupDecodingContext_1), int32(64))
	mBase = m.M
	v354 = m.ExcPending
	if v354 != 0 {
		goto L1
	} else {
		goto L86
	}
L86:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v335)+124)) = v353
	v359 = F_SlabContextCreate(m, v332, int32(_a_F_StartupDecodingContext_15), int32(_a_F_StartupDecodingContext_1), int32(232))
	mBase = m.M
	v360 = m.ExcPending
	if v360 != 0 {
		goto L1
	} else {
		goto L87
	}
L87:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v335)+128)) = v359
	v363 = int32(_a_F_StartupDecodingContext_1)
	v366 = F_GenerationContextCreate(m, v332, int32(_a_F_StartupDecodingContext_16), v363, v363, v363)
	mBase = m.M
	v367 = m.ExcPending
	if v367 != 0 {
		goto L1
	} else {
		goto L88
	}
L88:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v335)+132)) = v366
	*(*int64)(unsafe.Add(mBase, uint32(v324)+8)) = int64(34359738372)
	v371 = *(*int32)(unsafe.Add(mBase, uint32(v335)+120))
	*(*int32)(unsafe.Add(mBase, uint32(v324)+36)) = v371
	v376 = F_hash_create(m, int32(_a_F_StartupDecodingContext_17), int64(1000), v324, int32(1064))
	mBase = m.M
	v377 = m.ExcPending
	if v377 != 0 {
		goto L1
	} else {
		goto L89
	}
L89:
	;
	v378 = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v335)+152)) = v378
	v380 = int64(0)
	*(*int64)(unsafe.Add(mBase, uint32(v335)+144)) = v380
	*(*int64)(unsafe.Add(mBase, uint32(v335)+32)) = v380
	*(*int32)(unsafe.Add(mBase, uint32(v335))) = v376
	v387 = F_pairingheap_allocate(m, int32(1087), v378)
	mBase = m.M
	v388 = m.ExcPending
	if v388 != 0 {
		goto L1
	} else {
		goto L90
	}
L90:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v335)+136)) = int64(0)
	*(*int32)(unsafe.Add(mBase, uint32(v335)+156)) = v387
	v394 = int32(0)
	base.MemoryFill(m, v335+int32(160), v394, int32(72))
	*(*int32)(unsafe.Add(mBase, uint32(v335)+28)) = v394
	v400 = v335 + int32(20)
	*(*int32)(unsafe.Add(mBase, uint32(v335)+24)) = v400
	*(*int32)(unsafe.Add(mBase, uint32(v335)+20)) = v400
	v404 = v335 + int32(12)
	*(*int32)(unsafe.Add(mBase, uint32(v335)+16)) = v404
	*(*int32)(unsafe.Add(mBase, uint32(v335)+12)) = v404
	v408 = v335 + int32(4)
	*(*int32)(unsafe.Add(mBase, uint32(v335)+8)) = v408
	*(*int32)(unsafe.Add(mBase, uint32(v335)+4)) = v408
	v412 = *(*int32)(unsafe.Add(mBase, _c_F_StartupDecodingContext[0]))
	F_ReorderBufferCleanupSerializedTXNs(m, v412+int32(24))
	mBase = m.M
	v416 = m.ExcPending
	if v416 != 0 {
		goto L1
	} else {
		goto L91
	}
L91:
	;
	m.G0 = v324 + int32(48)
	*(*int32)(unsafe.Add(mBase, uint32(v43)+12)) = v335
	v421 = *(*int64)(unsafe.Add(mBase, uint32(v27)+128))
	v423 = *(*int32)(unsafe.Add(mBase, _c_F_StartupDecodingContext[1]))
	v428 = F_AllocSetContextCreateInternal(m, v423, int32(_a_F_StartupDecodingContext_18), int32(0), int32(_a_F_StartupDecodingContext_1), int32(_a_F_StartupDecodingContext_2))
	mBase = m.M
	v429 = m.ExcPending
	if v429 != 0 {
		goto L1
	} else {
		goto L92
	}
L92:
	;
	v430 = int32(_a_F_StartupDecodingContext_3)
	v431 = *(*int32)(unsafe.Add(mBase, _c_F_StartupDecodingContext[1]))
	*(*int32)(unsafe.Add(mBase, _c_F_StartupDecodingContext[1])) = v428
	v435 = F_palloc0(m, int32(88))
	mBase = m.M
	v436 = m.ExcPending
	if v436 != 0 {
		goto L1
	} else {
		goto L93
	}
L93:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v435)+64)) = int64(549755813888)
	*(*int32)(unsafe.Add(mBase, uint32(v435)+56)) = v335
	*(*int32)(unsafe.Add(mBase, uint32(v435)+4)) = v428
	*(*int32)(unsafe.Add(mBase, uint32(v435))) = int32(-1)
	v445 = F_palloc0_mul(m, int32(4), int32(128))
	mBase = m.M
	v446 = m.ExcPending
	if v446 != 0 {
		goto L1
	} else {
		goto L94
	}
L94:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v435)+80)) = int64(0)
	v449 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(v435)+72)) = uint8(v449)
	*(*int32)(unsafe.Add(mBase, uint32(v435)+76)) = v445
	*(*int32)(unsafe.Add(mBase, uint32(v435)+32)) = l2
	*(*uint8)(unsafe.Add(mBase, uint32(v435)+37)) = uint8(v6)
	*(*int64)(unsafe.Add(mBase, uint32(v435)+16)) = l1
	*(*uint8)(unsafe.Add(mBase, uint32(v435)+36)) = uint8(v4)
	*(*int64)(unsafe.Add(mBase, uint32(v435)+24)) = v421
	*(*int32)(unsafe.Add(mBase, _c_F_StartupDecodingContext[1])) = v431
	*(*int32)(unsafe.Add(mBase, uint32(v43)+16)) = v435
	v460 = *(*int32)(unsafe.Add(mBase, uint32(v43)+12))
	*(*int32)(unsafe.Add(mBase, uint32(v460)+112)) = v43
	v462 = *(*int32)(unsafe.Add(mBase, uint32(v43)+12))
	*(*int32)(unsafe.Add(mBase, uint32(v462)+40)) = int32(1059)
	v465 = *(*int32)(unsafe.Add(mBase, uint32(v43)+12))
	*(*int32)(unsafe.Add(mBase, uint32(v465)+44)) = int32(1060)
	v468 = *(*int32)(unsafe.Add(mBase, uint32(v43)+12))
	*(*int32)(unsafe.Add(mBase, uint32(v468)+48)) = int32(1061)
	v471 = *(*int32)(unsafe.Add(mBase, uint32(v43)+12))
	*(*int32)(unsafe.Add(mBase, uint32(v471)+52)) = int32(1062)
	v474 = *(*int32)(unsafe.Add(mBase, uint32(v43)+12))
	*(*int32)(unsafe.Add(mBase, uint32(v474)+56)) = int32(1063)
	v479 = *(*int32)(unsafe.Add(mBase, uint32(v43)+76))
	if v479 != 0 {
		v493 = v449
		goto L95
	} else {
		goto L96
	}
L95:
	;
	*(*uint8)(unsafe.Add(mBase, uint32(v43)+144)) = uint8(v493)
	v495 = *(*int32)(unsafe.Add(mBase, uint32(v43)+12))
	*(*int32)(unsafe.Add(mBase, uint32(v495)+76)) = int32(1064)
	v498 = *(*int32)(unsafe.Add(mBase, uint32(v43)+12))
	*(*int32)(unsafe.Add(mBase, uint32(v498)+80)) = int32(1065)
	v501 = *(*int32)(unsafe.Add(mBase, uint32(v43)+12))
	*(*int32)(unsafe.Add(mBase, uint32(v501)+84)) = int32(1066)
	v504 = *(*int32)(unsafe.Add(mBase, uint32(v43)+12))
	*(*int32)(unsafe.Add(mBase, uint32(v504)+88)) = int32(1067)
	v507 = *(*int32)(unsafe.Add(mBase, uint32(v43)+12))
	*(*int32)(unsafe.Add(mBase, uint32(v507)+92)) = int32(1068)
	v510 = *(*int32)(unsafe.Add(mBase, uint32(v43)+12))
	*(*int32)(unsafe.Add(mBase, uint32(v510)+96)) = int32(1069)
	v513 = *(*int32)(unsafe.Add(mBase, uint32(v43)+12))
	*(*int32)(unsafe.Add(mBase, uint32(v513)+100)) = int32(1070)
	v516 = *(*int32)(unsafe.Add(mBase, uint32(v43)+12))
	*(*int32)(unsafe.Add(mBase, uint32(v516)+104)) = int32(1071)
	v519 = *(*int32)(unsafe.Add(mBase, uint32(v43)+60))
	if v519 != 0 {
		v527 = v449
		goto L102
	} else {
		goto L103
	}
L96:
	;
	v481 = *(*int32)(unsafe.Add(mBase, uint32(v43)+80))
	if v481 != 0 {
		v493 = int32(1)
		goto L95
	} else {
		goto L97
	}
L97:
	;
	v483 = *(*int32)(unsafe.Add(mBase, uint32(v43)+84))
	if v483 != 0 {
		v493 = int32(1)
		goto L95
	} else {
		goto L98
	}
L98:
	;
	v485 = *(*int32)(unsafe.Add(mBase, uint32(v43)+92))
	if v485 != 0 {
		v493 = int32(1)
		goto L95
	} else {
		goto L99
	}
L99:
	;
	v487 = *(*int32)(unsafe.Add(mBase, uint32(v43)+96))
	if v487 != 0 {
		v493 = int32(1)
		goto L95
	} else {
		goto L100
	}
L100:
	;
	v489 = *(*int32)(unsafe.Add(mBase, uint32(v43)+100))
	if v489 != 0 {
		v493 = int32(1)
		goto L95
	} else {
		goto L101
	}
L101:
	;
	v490 = *(*int32)(unsafe.Add(mBase, uint32(v43)+104))
	v493 = base.B2i32(v490 != int32(0))
	goto L95
L102:
	;
	*(*uint8)(unsafe.Add(mBase, uint32(v43)+145)) = uint8(v527)
	v529 = *(*int32)(unsafe.Add(mBase, uint32(v43)+12))
	*(*int32)(unsafe.Add(mBase, uint32(v529)+60)) = int32(1072)
	v532 = *(*int32)(unsafe.Add(mBase, uint32(v43)+12))
	*(*int32)(unsafe.Add(mBase, uint32(v532)+64)) = int32(1073)
	v535 = *(*int32)(unsafe.Add(mBase, uint32(v43)+12))
	*(*int32)(unsafe.Add(mBase, uint32(v535)+68)) = int32(1074)
	v538 = *(*int32)(unsafe.Add(mBase, uint32(v43)+12))
	*(*int32)(unsafe.Add(mBase, uint32(v538)+72)) = int32(1075)
	v541 = *(*int32)(unsafe.Add(mBase, uint32(v43)+12))
	*(*int32)(unsafe.Add(mBase, uint32(v541)+108)) = int32(1076)
	v544 = F_makeStringInfo(m)
	mBase = m.M
	v545 = m.ExcPending
	if v545 != 0 {
		goto L1
	} else {
		goto L108
	}
L103:
	;
	v520 = *(*int32)(unsafe.Add(mBase, uint32(v43)+64))
	if v520 != 0 {
		v527 = v449
		goto L102
	} else {
		goto L104
	}
L104:
	;
	v521 = *(*int32)(unsafe.Add(mBase, uint32(v43)+68))
	if v521 != 0 {
		v527 = v449
		goto L102
	} else {
		goto L105
	}
L105:
	;
	v522 = *(*int32)(unsafe.Add(mBase, uint32(v43)+72))
	if v522 != 0 {
		v527 = v449
		goto L102
	} else {
		goto L106
	}
L106:
	;
	v523 = *(*int32)(unsafe.Add(mBase, uint32(v43)+88))
	if v523 != 0 {
		v527 = v449
		goto L102
	} else {
		goto L107
	}
L107:
	;
	v524 = *(*int32)(unsafe.Add(mBase, uint32(v43)+56))
	v527 = base.B2i32(v524 != int32(0))
	goto L102
L108:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v43)+128)) = l10
	*(*int32)(unsafe.Add(mBase, uint32(v43)+124)) = l9
	*(*int32)(unsafe.Add(mBase, uint32(v43)+120)) = l8
	*(*int32)(unsafe.Add(mBase, uint32(v43)+132)) = v544
	*(*int32)(unsafe.Add(mBase, uint32(v43)+116)) = l0
	*(*uint8)(unsafe.Add(mBase, uint32(v43)+20)) = uint8(v5)
	*(*int32)(unsafe.Add(mBase, _c_F_StartupDecodingContext[1])) = v39
	m.G0 = v24 + int32(96)
	return v43
L109:
	;
	F_errmsg_internal(m, int32(_a_F_StartupDecodingContext_19), int32(0))
	mBase = m.M
	v565 = m.ExcPending
	if v565 != 0 {
		goto L1
	} else {
		goto L110
	}
L110:
	;
	F_errfinish(m, int32(_a_F_StartupDecodingContext_9), int32(823), int32(_a_F_StartupDecodingContext_20))
	mBase = m.M
	v570 = m.ExcPending
	if v570 != 0 {
		goto L1
	} else {
		goto L111
	}
L111:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L112:
	;
	F_errmsg_internal(m, int32(_a_F_StartupDecodingContext_21), int32(0))
	mBase = m.M
	v578 = m.ExcPending
	if v578 != 0 {
		goto L1
	} else {
		goto L113
	}
L113:
	;
	F_errfinish(m, int32(_a_F_StartupDecodingContext_9), int32(829), int32(_a_F_StartupDecodingContext_20))
	mBase = m.M
	v583 = m.ExcPending
	if v583 != 0 {
		goto L1
	} else {
		goto L114
	}
L114:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L115:
	;
	F_errmsg_internal(m, int32(_a_F_StartupDecodingContext_22), int32(0))
	mBase = m.M
	v591 = m.ExcPending
	if v591 != 0 {
		goto L1
	} else {
		goto L116
	}
L116:
	;
	F_errfinish(m, int32(_a_F_StartupDecodingContext_9), int32(831), int32(_a_F_StartupDecodingContext_20))
	mBase = m.M
	v596 = m.ExcPending
	if v596 != 0 {
		goto L1
	} else {
		goto L117
	}
L117:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L118:
	;
	F_errmsg_internal(m, int32(_a_F_StartupDecodingContext_23), int32(0))
	mBase = m.M
	v604 = m.ExcPending
	if v604 != 0 {
		goto L1
	} else {
		goto L119
	}
L119:
	;
	F_errfinish(m, int32(_a_F_StartupDecodingContext_9), int32(833), int32(_a_F_StartupDecodingContext_20))
	mBase = m.M
	v609 = m.ExcPending
	if v609 != 0 {
		goto L1
	} else {
		goto L120
	}
L120:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L121:
	;
	F_errcode(m, int32(_a_F_StartupDecodingContext_24))
	mBase = m.M
	v616 = m.ExcPending
	if v616 != 0 {
		goto L1
	} else {
		goto L122
	}
L122:
	;
	F_errmsg(m, int32(_a_F_StartupDecodingContext_25), int32(0))
	mBase = m.M
	v620 = m.ExcPending
	if v620 != 0 {
		goto L1
	} else {
		goto L123
	}
L123:
	;
	v623 = F_errdetail(m, int32(_a_F_StartupDecodingContext_26), int32(0))
	mBase = m.M
	v624 = m.ExcPending
	if v624 != 0 {
		goto L1
	} else {
		goto L124
	}
L124:
	;
	F_errfinish(m, int32(_a_F_StartupDecodingContext_9), int32(288), int32(_a_F_StartupDecodingContext_10))
	mBase = m.M
	v629 = m.ExcPending
	if v629 != 0 {
		goto L1
	} else {
		goto L125
	}
L125:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L126:
	;
	F_pfree(m, v115)
	mBase = m.M
	v654 = m.ExcPending
	if v654 != 0 {
		goto L1
	} else {
		goto L127
	}
L127:
	;
	goto L4
L128:
	;
	F_errcode(m, int32(16797828))
	mBase = m.M
	v682 = m.ExcPending
	if v682 != 0 {
		goto L1
	} else {
		goto L129
	}
L129:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v24)+64)) = v49
	F_errmsg(m, int32(_a_F_StartupDecodingContext_27), v24-int32(-64))
	mBase = m.M
	v688 = m.ExcPending
	if v688 != 0 {
		goto L1
	} else {
		goto L130
	}
L130:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v24)+48)) = int32(_a_F_StartupDecodingContext_11)
	v692 = *(*int32)(unsafe.Add(mBase, _c_F_StartupDecodingContext[4]))
	*(*int32)(unsafe.Add(mBase, uint32(v24)+52)) = v692
	F_errdetail_log(m, int32(_a_F_StartupDecodingContext_28), v24+int32(48))
	mBase = m.M
	v698 = m.ExcPending
	if v698 != 0 {
		goto L1
	} else {
		goto L131
	}
L131:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v24)+32)) = int32(_a_F_StartupDecodingContext_11)
	F_errhint(m, int32(_a_F_StartupDecodingContext_29), v24+int32(32))
	mBase = m.M
	v705 = m.ExcPending
	if v705 != 0 {
		goto L1
	} else {
		goto L132
	}
L132:
	;
	F_errfinish(m, int32(_a_F_StartupDecodingContext_9), int32(255), int32(_a_F_StartupDecodingContext_10))
	mBase = m.M
	v710 = m.ExcPending
	if v710 != 0 {
		goto L1
	} else {
		goto L133
	}
L133:
	;
	base.Wasm_trap_unreachable()
	for {
	}
}
func F_StartupProcExit(m *base.Module, l0 int32, l1 int64) {
	mBase := m.M
	_ = mBase
	var v4 int32
	_ = v4
	var v6 int32
	_ = v6
	v4 = *(*int32)(unsafe.Add(mBase, _c_F_StartupProcExit[0]))
	if v4 != 0 {
		F_ShutdownRecoveryTransactionEnvironment(m)
		mBase = m.M
		v6 = m.ExcPending
		if v6 != 0 {
			return
		} else {
			return
		}
	} else {
		return
	}
}
func F_StartupProcShutdownHandler(m *base.Module, l0 int32, l1 int32) {
	mBase := m.M
	_ = mBase
	var v4 int32
	_ = v4
	var v7 int32
	_ = v7
	var v12 int32
	_ = v12
	v4 = *(*int32)(unsafe.Add(mBase, _c_F_StartupProcShutdownHandler[0]))
	if v4 != 0 {
		F_proc_exit(m, int32(1))
		mBase = m.M
		v7 = m.ExcPending
		if v7 != 0 {
			return
		} else {
			base.Wasm_trap_unreachable()
			for {
			}
		}
	} else {
		*(*int32)(unsafe.Add(mBase, _c_F_StartupProcShutdownHandler[1])) = int32(1)
		v12 = *(*int32)(unsafe.Add(mBase, _c_F_StartupProcShutdownHandler[2]))
		F_SetLatch(m, v12+int32(4))
		mBase = m.M
		return
	}
}
func F_StartupProcSigHupHandler(m *base.Module, l0 int32, l1 int32) {
	mBase := m.M
	_ = mBase
	var v7 int32
	_ = v7
	*(*int32)(unsafe.Add(mBase, _c_F_StartupProcSigHupHandler[0])) = int32(1)
	v7 = *(*int32)(unsafe.Add(mBase, _c_F_StartupProcSigHupHandler[1]))
	F_SetLatch(m, v7+int32(4))
	mBase = m.M
	return
}
func F_StartupProcessMain(m *base.Module, l0 int32, l1 int32) {
	mBase := m.M
	_ = mBase
	var v4 int32
	_ = v4
	var v8 int32
	_ = v8
	var v13 int32
	_ = v13
	var v15 int32
	_ = v15
	var v18 int32
	_ = v18
	var v47 int32
	_ = v47
	var v53 int32
	_ = v53
	var v55 int32
	_ = v55
	var v57 int32
	_ = v57
	var v89 int32
	_ = v89
	var v97 int32
	_ = v97
	var v99 int32
	_ = v99
	var v102 int32
	_ = v102
	var v131 int32
	_ = v131
	var v135 int32
	_ = v135
	var v145 int32
	_ = v145
	var v147 int32
	_ = v147
	var v148 int32
	_ = v148
	var v149 int32
	_ = v149
	var v152 int64
	_ = v152
	var v163 int32
	_ = v163
	var v165 int32
	_ = v165
	var v180 int32
	_ = v180
	var v182 int32
	_ = v182
	var v199 int32
	_ = v199
	var v201 int32
	_ = v201
	var v202 int32
	_ = v202
	var v205 int64
	_ = v205
	var v218 int32
	_ = v218
	var v225 int32
	_ = v225
	var v227 int32
	_ = v227
	var v229 int32
	_ = v229
	var v261 int32
	_ = v261
	var v269 int32
	_ = v269
	var v271 int32
	_ = v271
	var v274 int32
	_ = v274
	var v303 int32
	_ = v303
	var v311 int32
	_ = v311
	var v313 int32
	_ = v313
	var v316 int32
	_ = v316
	var v345 int32
	_ = v345
	var v353 int32
	_ = v353
	var v355 int32
	_ = v355
	var v357 int32
	_ = v357
	var v380 int32
	_ = v380
	var v387 int32
	_ = v387
	var v394 int32
	_ = v394
	var v398 int32
	_ = v398
	var v402 int32
	_ = v402
	var v406 int32
	_ = v406
	var v408 int32
	_ = v408
	var v411 int32
	_ = v411
	F_AuxiliaryProcessMainCommon(m)
	mBase = m.M
	v4 = m.ExcPending
	if v4 != 0 {
		return
	} else {
		F_on_shmem_exit(m, int32(1019), int64(0))
		mBase = m.M
		v8 = m.ExcPending
		if v8 != 0 {
			return
		} else {
			v13 = m.G0
			v15 = v13 - int32(32)
			m.G0 = v15
			v18 = int32(1022)
			switch v18 {
			case 0, 2:
			default:
				*(*int32)(unsafe.Add(mBase, _c_F_StartupProcessMain[0])) = int32(1020)
			}
			F_sigemptyset(m, v15+int32(16))
			mBase = m.M
			*(*int32)(unsafe.Add(mBase, uint32(v15)+24)) = int32(268435456)
			switch v18 {
			case 0:
				*(*int32)(unsafe.Add(mBase, uint32(v15)+12)) = int32(-2)
			default:
				*(*int32)(unsafe.Add(mBase, uint32(v15)+24)) = int32(268435460)
				*(*int32)(unsafe.Add(mBase, uint32(v15)+12)) = int32(_a_F_StartupProcessMain_0)
			case 2:
				*(*int32)(unsafe.Add(mBase, uint32(v15)+12)) = int32(0)
			}
			v47 = F___sigaction(m, int32(1), v15+int32(12), int32(0))
			mBase = m.M
			m.G0 = v15 + int32(32)
			v53 = int32(0)
			v55 = m.G0
			v57 = v55 - int32(32)
			m.G0 = v57
			switch v53 {
			case 0, 2:
			default:
				*(*int32)(unsafe.Add(mBase, _c_F_StartupProcessMain[1])) = int32(-2)
			}
			F_sigemptyset(m, v57+int32(16))
			mBase = m.M
			*(*int32)(unsafe.Add(mBase, uint32(v57)+24)) = int32(268435456)
			switch v53 {
			case 0:
				*(*int32)(unsafe.Add(mBase, uint32(v57)+12)) = int32(-2)
			default:
				*(*int32)(unsafe.Add(mBase, uint32(v57)+24)) = int32(268435460)
				*(*int32)(unsafe.Add(mBase, uint32(v57)+12)) = int32(_a_F_StartupProcessMain_0)
			case 2:
				*(*int32)(unsafe.Add(mBase, uint32(v57)+12)) = int32(0)
			}
			v89 = F___sigaction(m, int32(2), v57+int32(12), int32(0))
			mBase = m.M
			m.G0 = v57 + int32(32)
			v97 = m.G0
			v99 = v97 - int32(32)
			m.G0 = v99
			v102 = int32(1023)
			switch v102 {
			case 0, 2:
			default:
				*(*int32)(unsafe.Add(mBase, _c_F_StartupProcessMain[2])) = int32(1021)
			}
			F_sigemptyset(m, v99+int32(16))
			mBase = m.M
			*(*int32)(unsafe.Add(mBase, uint32(v99)+24)) = int32(268435456)
			switch v102 {
			case 0:
				*(*int32)(unsafe.Add(mBase, uint32(v99)+12)) = int32(-2)
			default:
				*(*int32)(unsafe.Add(mBase, uint32(v99)+24)) = int32(268435460)
				*(*int32)(unsafe.Add(mBase, uint32(v99)+12)) = int32(_a_F_StartupProcessMain_0)
			case 2:
				*(*int32)(unsafe.Add(mBase, uint32(v99)+12)) = int32(0)
			}
			v131 = F___sigaction(m, int32(15), v99+int32(12), int32(0))
			mBase = m.M
			m.G0 = v99 + int32(32)
			v135 = int32(0)
			*(*int32)(unsafe.Add(mBase, _c_F_StartupProcessMain[3])) = v135
			*(*int32)(unsafe.Add(mBase, _c_F_StartupProcessMain[4])) = v135
			v145 = v135
			for {
				v147 = int32(40)
				v148 = v145 * v147
				v149 = int32(0)
				*(*uint8)(unsafe.Add(mBase, uint32(v148)+uint32(_c_F_StartupProcessMain[5]))) = uint8(v149)
				*(*int32)(unsafe.Add(mBase, uint32(v148)+uint32(_c_F_StartupProcessMain[6]))) = v145
				v152 = int64(0)
				*(*int64)(unsafe.Add(mBase, uint32(v148)+uint32(_c_F_StartupProcessMain[7]))) = v152
				*(*int32)(unsafe.Add(mBase, uint32(v148)+uint32(_c_F_StartupProcessMain[8]))) = v149
				*(*int64)(unsafe.Add(mBase, uint32(v148)+uint32(_c_F_StartupProcessMain[9]))) = v152
				*(*int32)(unsafe.Add(mBase, uint32(v148)+uint32(_c_F_StartupProcessMain[10]))) = v149
				*(*uint8)(unsafe.Add(mBase, uint32(v148)+uint32(_c_F_StartupProcessMain[11]))) = uint8(v149)
				v163 = v145 | int32(1)
				v165 = v163 * v147
				*(*uint8)(unsafe.Add(mBase, uint32(v165)+uint32(_c_F_StartupProcessMain[5]))) = uint8(v149)
				*(*int32)(unsafe.Add(mBase, uint32(v165)+uint32(_c_F_StartupProcessMain[6]))) = v163
				*(*int64)(unsafe.Add(mBase, uint32(v165)+uint32(_c_F_StartupProcessMain[7]))) = v152
				*(*int32)(unsafe.Add(mBase, uint32(v165)+uint32(_c_F_StartupProcessMain[8]))) = v149
				*(*int64)(unsafe.Add(mBase, uint32(v165)+uint32(_c_F_StartupProcessMain[9]))) = v152
				*(*int32)(unsafe.Add(mBase, uint32(v165)+uint32(_c_F_StartupProcessMain[10]))) = v149
				*(*uint8)(unsafe.Add(mBase, uint32(v165)+uint32(_c_F_StartupProcessMain[11]))) = uint8(v149)
				v180 = v145 | int32(2)
				v182 = v180 * v147
				*(*uint8)(unsafe.Add(mBase, uint32(v182)+uint32(_c_F_StartupProcessMain[5]))) = uint8(v149)
				*(*int32)(unsafe.Add(mBase, uint32(v182)+uint32(_c_F_StartupProcessMain[6]))) = v180
				*(*int64)(unsafe.Add(mBase, uint32(v182)+uint32(_c_F_StartupProcessMain[7]))) = v152
				*(*int32)(unsafe.Add(mBase, uint32(v182)+uint32(_c_F_StartupProcessMain[8]))) = v149
				*(*int64)(unsafe.Add(mBase, uint32(v182)+uint32(_c_F_StartupProcessMain[9]))) = v152
				*(*int32)(unsafe.Add(mBase, uint32(v182)+uint32(_c_F_StartupProcessMain[10]))) = v149
				*(*uint8)(unsafe.Add(mBase, uint32(v182)+uint32(_c_F_StartupProcessMain[11]))) = uint8(v149)
				if v145 != int32(20) {
					v199 = v145 | int32(3)
					v201 = v199 * int32(40)
					v202 = int32(0)
					*(*uint8)(unsafe.Add(mBase, uint32(v201)+uint32(_c_F_StartupProcessMain[5]))) = uint8(v202)
					*(*int32)(unsafe.Add(mBase, uint32(v201)+uint32(_c_F_StartupProcessMain[6]))) = v199
					v205 = int64(0)
					*(*int64)(unsafe.Add(mBase, uint32(v201)+uint32(_c_F_StartupProcessMain[7]))) = v205
					*(*int32)(unsafe.Add(mBase, uint32(v201)+uint32(_c_F_StartupProcessMain[8]))) = v202
					*(*int64)(unsafe.Add(mBase, uint32(v201)+uint32(_c_F_StartupProcessMain[9]))) = v205
					*(*int32)(unsafe.Add(mBase, uint32(v201)+uint32(_c_F_StartupProcessMain[10]))) = v202
					*(*uint8)(unsafe.Add(mBase, uint32(v201)+uint32(_c_F_StartupProcessMain[11]))) = uint8(v202)
					v145 = v145 + int32(4)
					continue
				} else {
					break
				}
				break
			}
			v218 = int32(1)
			*(*uint8)(unsafe.Add(mBase, _c_F_StartupProcessMain[12])) = uint8(v218)
			F_pqsignal_be(m, int32(14), int32(1992))
			mBase = m.M
			v225 = int32(0)
			v227 = m.G0
			v229 = v227 - int32(32)
			m.G0 = v229
			switch v225 {
			case 0, 2:
			default:
				*(*int32)(unsafe.Add(mBase, _c_F_StartupProcessMain[13])) = int32(-2)
			}
			F_sigemptyset(m, v229+int32(16))
			mBase = m.M
			*(*int32)(unsafe.Add(mBase, uint32(v229)+24)) = int32(268435456)
			switch v225 {
			case 0:
				*(*int32)(unsafe.Add(mBase, uint32(v229)+12)) = int32(-2)
			default:
				*(*int32)(unsafe.Add(mBase, uint32(v229)+24)) = int32(268435460)
				*(*int32)(unsafe.Add(mBase, uint32(v229)+12)) = int32(_a_F_StartupProcessMain_0)
			case 2:
				*(*int32)(unsafe.Add(mBase, uint32(v229)+12)) = int32(0)
			}
			v261 = F___sigaction(m, int32(13), v229+int32(12), int32(0))
			mBase = m.M
			m.G0 = v229 + int32(32)
			v269 = m.G0
			v271 = v269 - int32(32)
			m.G0 = v271
			v274 = int32(970)
			switch v274 {
			case 0, 2:
			default:
				*(*int32)(unsafe.Add(mBase, _c_F_StartupProcessMain[14])) = int32(968)
			}
			F_sigemptyset(m, v271+int32(16))
			mBase = m.M
			*(*int32)(unsafe.Add(mBase, uint32(v271)+24)) = int32(268435456)
			switch v274 {
			case 0:
				*(*int32)(unsafe.Add(mBase, uint32(v271)+12)) = int32(-2)
			default:
				*(*int32)(unsafe.Add(mBase, uint32(v271)+24)) = int32(268435460)
				*(*int32)(unsafe.Add(mBase, uint32(v271)+12)) = int32(_a_F_StartupProcessMain_0)
			case 2:
				*(*int32)(unsafe.Add(mBase, uint32(v271)+12)) = int32(0)
			}
			v303 = F___sigaction(m, int32(10), v271+int32(12), int32(0))
			mBase = m.M
			m.G0 = v271 + int32(32)
			v311 = m.G0
			v313 = v311 - int32(32)
			m.G0 = v313
			v316 = int32(1024)
			switch v316 {
			case 0, 2:
			default:
				*(*int32)(unsafe.Add(mBase, _c_F_StartupProcessMain[15])) = int32(1022)
			}
			F_sigemptyset(m, v313+int32(16))
			mBase = m.M
			*(*int32)(unsafe.Add(mBase, uint32(v313)+24)) = int32(268435456)
			switch v316 {
			case 0:
				*(*int32)(unsafe.Add(mBase, uint32(v313)+12)) = int32(-2)
			default:
				*(*int32)(unsafe.Add(mBase, uint32(v313)+24)) = int32(268435460)
				*(*int32)(unsafe.Add(mBase, uint32(v313)+12)) = int32(_a_F_StartupProcessMain_0)
			case 2:
				*(*int32)(unsafe.Add(mBase, uint32(v313)+12)) = int32(0)
			}
			v345 = F___sigaction(m, int32(12), v313+int32(12), int32(0))
			mBase = m.M
			m.G0 = v313 + int32(32)
			v353 = m.G0
			v355 = v353 - int32(32)
			m.G0 = v355
			v357 = int32(2)
			switch v357 {
			case 0, 2:
			default:
				*(*int32)(unsafe.Add(mBase, _c_F_StartupProcessMain[16])) = int32(0)
			}
			F_sigemptyset(m, v355+int32(16))
			mBase = m.M
			*(*int32)(unsafe.Add(mBase, uint32(v355)+24)) = int32(268435456)
			switch v357 {
			case 0:
				*(*int32)(unsafe.Add(mBase, uint32(v355)+12)) = int32(-2)
				v380 = int32(268435457)
			default:
				*(*int32)(unsafe.Add(mBase, uint32(v355)+24)) = int32(268435460)
				*(*int32)(unsafe.Add(mBase, uint32(v355)+12)) = int32(_a_F_StartupProcessMain_0)
				v380 = int32(268435461)
			case 2:
				*(*int32)(unsafe.Add(mBase, uint32(v355)+12)) = int32(0)
				v380 = int32(268435457)
			}
			*(*int32)(unsafe.Add(mBase, uint32(v355)+24)) = v380
			v387 = F___sigaction(m, int32(17), v355+int32(12), int32(0))
			mBase = m.M
			m.G0 = v355 + int32(32)
			F_RegisterTimeout(m, int32(4), int32(1023))
			mBase = m.M
			v394 = m.ExcPending
			if v394 != 0 {
				return
			} else {
				F_RegisterTimeout(m, int32(5), int32(1024))
				mBase = m.M
				v398 = m.ExcPending
				if v398 != 0 {
					return
				} else {
					F_RegisterTimeout(m, int32(6), int32(1025))
					mBase = m.M
					v402 = m.ExcPending
					if v402 != 0 {
						return
					} else {
						F_pgmem_sigprocmask(m, int32(_a_F_StartupProcessMain_1), int32(0))
						mBase = m.M
						v406 = m.ExcPending
						if v406 != 0 {
							return
						} else {
							F_StartupXLOG(m)
							mBase = m.M
							v408 = m.ExcPending
							if v408 != 0 {
								return
							} else {
								F_proc_exit(m, int32(0))
								mBase = m.M
								v411 = m.ExcPending
								if v411 != 0 {
									return
								} else {
									base.Wasm_trap_unreachable()
									for {
									}
								}
							}
						}
					}
				}
			}
		}
	}
}
