package p3

import (
	base "github.com/shibukawa/pgmem/internal/aot/pgaot/base"
	"unsafe"
)

func F_CopyMultiInsertInfoFlush(m *base.Module, l0 int32, l1 int32, l2 int32) {
	mBase := m.M
	_ = mBase
	var v4 int32
	_ = v4
	var v20 int32
	_ = v20
	var v22 int32
	_ = v22
	var v24 int32
	_ = v24
	var v29 int32
	_ = v29
	var v48 int32
	_ = v48
	var v51 int32
	_ = v51
	var v55 int32
	_ = v55
	var v56 int32
	_ = v56
	var v57 int32
	_ = v57
	var v58 int32
	_ = v58
	var v59 int32
	_ = v59
	var v60 int32
	_ = v60
	var v61 int32
	_ = v61
	var v62 int32
	_ = v62
	var v70 int32
	_ = v70
	var v86 int32
	_ = v86
	var v88 int32
	_ = v88
	var v90 int32
	_ = v90
	var v97 int32
	_ = v97
	var v98 int32
	_ = v98
	var v99 int32
	_ = v99
	var v100 int32
	_ = v100
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
	var v116 int32
	_ = v116
	var v135 int32
	_ = v135
	var v138 int32
	_ = v138
	var v140 int32
	_ = v140
	var v142 int32
	_ = v142
	var v143 int32
	_ = v143
	var v155 int32
	_ = v155
	var v164 int64
	_ = v164
	var v166 int64
	_ = v166
	var v171 int32
	_ = v171
	var v175 int32
	_ = v175
	var v180 int32
	_ = v180
	var v182 int32
	_ = v182
	var v183 int32
	_ = v183
	var v186 int32
	_ = v186
	var v190 int32
	_ = v190
	var v192 int32
	_ = v192
	var v193 int32
	_ = v193
	var v201 int32
	_ = v201
	var v202 int32
	_ = v202
	var v208 int32
	_ = v208
	var v233 int32
	_ = v233
	var v234 int32
	_ = v234
	var v244 int32
	_ = v244
	var v247 int32
	_ = v247
	var v262 int32
	_ = v262
	var v263 int32
	_ = v263
	var v264 int32
	_ = v264
	var v265 int32
	_ = v265
	var v267 int32
	_ = v267
	var v268 int32
	_ = v268
	var v269 int32
	_ = v269
	var v270 int32
	_ = v270
	var v272 int32
	_ = v272
	var v273 int32
	_ = v273
	var v274 int32
	_ = v274
	var v275 int32
	_ = v275
	var v277 int32
	_ = v277
	var v278 int32
	_ = v278
	var v279 int32
	_ = v279
	var v280 int32
	_ = v280
	var v282 int32
	_ = v282
	var v283 int32
	_ = v283
	var v284 int32
	_ = v284
	var v286 int32
	_ = v286
	var v293 int32
	_ = v293
	var v312 int32
	_ = v312
	var v314 int32
	_ = v314
	var v331 int32
	_ = v331
	var v332 int32
	_ = v332
	var v333 int32
	_ = v333
	var v335 int32
	_ = v335
	var v336 int32
	_ = v336
	var v339 int32
	_ = v339
	var v360 int32
	_ = v360
	var v362 int32
	_ = v362
	var v363 int32
	_ = v363
	var v364 int32
	_ = v364
	var v365 int32
	_ = v365
	var v367 int64
	_ = v367
	var v368 int32
	_ = v368
	var v371 int32
	_ = v371
	var v372 int32
	_ = v372
	var v373 int32
	_ = v373
	var v374 int32
	_ = v374
	var v375 int32
	_ = v375
	var v377 int32
	_ = v377
	var v379 int32
	_ = v379
	var v380 int32
	_ = v380
	var v381 int32
	_ = v381
	var v382 int32
	_ = v382
	var v384 int32
	_ = v384
	var v390 int32
	_ = v390
	var v395 int32
	_ = v395
	var v411 int32
	_ = v411
	var v417 int64
	_ = v417
	var v419 int32
	_ = v419
	var v422 int32
	_ = v422
	var v423 int32
	_ = v423
	var v426 int32
	_ = v426
	var v427 int32
	_ = v427
	var v428 int32
	_ = v428
	var v429 int32
	_ = v429
	var v431 int32
	_ = v431
	var v433 int32
	_ = v433
	var v434 int32
	_ = v434
	var v437 int32
	_ = v437
	var v440 int32
	_ = v440
	var v446 int64
	_ = v446
	var v451 int32
	_ = v451
	var v453 int32
	_ = v453
	var v455 int32
	_ = v455
	var v461 int32
	_ = v461
	var v462 int32
	_ = v462
	var v463 int32
	_ = v463
	var v465 int32
	_ = v465
	var v467 int32
	_ = v467
	var v488 int64
	_ = v488
	var v490 int64
	_ = v490
	var v495 int32
	_ = v495
	var v499 int32
	_ = v499
	var v504 int32
	_ = v504
	var v506 int32
	_ = v506
	var v507 int32
	_ = v507
	var v510 int32
	_ = v510
	var v514 int32
	_ = v514
	var v516 int32
	_ = v516
	var v517 int32
	_ = v517
	var v525 int32
	_ = v525
	var v526 int32
	_ = v526
	var v532 int32
	_ = v532
	var v560 int32
	_ = v560
	var v561 int32
	_ = v561
	var v584 int32
	_ = v584
	var v590 int32
	_ = v590
	var v606 int32
	_ = v606
	var v609 int32
	_ = v609
	var v610 int32
	_ = v610
	var v611 int32
	_ = v611
	var v613 int32
	_ = v613
	var v614 int32
	_ = v614
	var v616 int32
	_ = v616
	var v617 int32
	_ = v617
	var v619 int32
	_ = v619
	var v620 int32
	_ = v620
	var v621 int32
	_ = v621
	var v623 int32
	_ = v623
	var v624 int32
	_ = v624
	var v625 int32
	_ = v625
	var v627 int32
	_ = v627
	var v630 int32
	_ = v630
	var v632 int32
	_ = v632
	var v637 int32
	_ = v637
	var v656 int32
	_ = v656
	var v658 int32
	_ = v658
	var v660 int32
	_ = v660
	var v664 int32
	_ = v664
	var v665 int32
	_ = v665
	var v666 int32
	_ = v666
	var v669 int32
	_ = v669
	var v672 int32
	_ = v672
	var v674 int32
	_ = v674
	var v678 int32
	_ = v678
	var v679 int32
	_ = v679
	var v680 int32
	_ = v680
	var v681 int32
	_ = v681
	v4 = int32(0)
	v20 = m.G0
	v22 = v20 - int32(16)
	m.G0 = v22
	v24 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	if v24 == v4 {
		goto L2
	} else {
		goto L3
	}
L1:
	;
	m.G0 = v22 + int32(16)
	return
L2:
	;
	*(*int64)(unsafe.Add(mBase, uint32(l0)+4)) = int64(0)
	goto L1
L3:
	;
	goto L4
L4:
	;
	v29 = *(*int32)(unsafe.Add(mBase, uint32(v24)+4))
	if int32(0) < v29 {
		goto L5
	} else {
		goto L6
	}
L5:
	;
	v48 = v4
	goto L8
L6:
	;
	goto L7
L7:
	;
	*(*int64)(unsafe.Add(mBase, uint32(l0)+4)) = int64(0)
	v584 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	if v584 == int32(0) {
		goto L1
	} else {
		goto L83
	}
L8:
	;
	v51 = *(*int32)(unsafe.Add(mBase, uint32(v24)+12))
	v55 = *(*int32)(unsafe.Add(mBase, uint32(v51+v48<<(uint(int32(2))%32))))
	v56 = *(*int32)(unsafe.Add(mBase, uint32(v55)+4008))
	v57 = *(*int32)(unsafe.Add(mBase, uint32(l0)+16))
	v58 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	v59 = *(*int32)(unsafe.Add(mBase, uint32(v55)+4000))
	v60 = *(*int32)(unsafe.Add(mBase, uint32(v59)+84))
	if v60 != 0 {
		goto L11
	} else {
		goto L12
	}
L9:
	;
	goto L7
L10:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v55)+4008)) = int32(0)
	v560 = v48 + int32(1)
	v561 = *(*int32)(unsafe.Add(mBase, uint32(v24)+4))
	if v560 < v561 {
		v48 = v560
		goto L8
	} else {
		goto L82
	}
L11:
	;
	v61 = *(*int32)(unsafe.Add(mBase, uint32(v59)+104))
	v62 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(v58)+208)) = uint8(v62)
	if v56 <= int32(0) {
		goto L14
	} else {
		goto L15
	}
L12:
	;
	goto L13
L13:
	;
	v362 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v58)+320)))
	v363 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
	v364 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v365 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v58)+320)) = uint8(v365)
	v367 = *(*int64)(unsafe.Add(mBase, uint32(v58)+192))
	v368 = *(*int32)(unsafe.Add(mBase, uint32(v57)+152))
	if v368 == v365 {
		goto L53
	} else {
		goto L54
	}
L14:
	;
	v360 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v58)+208)) = uint8(v360)
	goto L10
L15:
	;
	v70 = int32(0)
	goto L16
L16:
	;
	v86 = v56 - v70
	if v61 < v86 {
		goto L18
	} else {
		goto L19
	}
L17:
	;
	v233 = v56 & int32(3)
	v234 = int32(0)
	if base.Ui32(int32(4)) <= base.Ui32(v56) {
		goto L38
	} else {
		goto L39
	}
L18:
	;
	v88 = v61
	goto L20
L19:
	;
	v88 = v86
	goto L20
L20:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v22)+12)) = v88
	v90 = v70 + v88
	v97 = *(*int32)(unsafe.Add(mBase, uint32(v59)+84))
	v98 = *(*int32)(unsafe.Add(mBase, uint32(v97)+56))
	v99 = m.T0[v98].(func(*base.Module, int32, int32, int32, int32, int32) int32)(m, v57, v59, v55+v70<<(uint(int32(2))%32), int32(0), v22+int32(12))
	mBase = m.M
	v100 = m.ExcPending
	if v100 != 0 {
		goto L21
	} else {
		goto L22
	}
L21:
	;
	return
L22:
	;
	v101 = *(*int32)(unsafe.Add(mBase, uint32(v22)+12))
	if int32(0) < v101 {
		goto L23
	} else {
		goto L24
	}
L23:
	;
	v104 = *(*int32)(unsafe.Add(mBase, uint32(v59)+52))
	if v104 == int32(0) {
		v155 = v101
		goto L26
	} else {
		goto L27
	}
L24:
	;
	goto L25
L25:
	;
	if v90 < v56 {
		v70 = v90
		goto L16
	} else {
		goto L37
	}
L26:
	;
	v164 = *(*int64)(unsafe.Add(mBase, uint32(l2)))
	v166 = v164 + base.I64_extend_i32_s(v155)
	*(*int64)(unsafe.Add(mBase, uint32(l2))) = v166
	v171 = *(*int32)(unsafe.Add(mBase, _c_F_CopyMultiInsertInfoFlush[0]))
	if v171 == int32(0) {
		goto L34
	} else {
		goto L35
	}
L27:
	;
	v107 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v104)+9)))
	if v107 != int32(1) {
		v155 = v101
		goto L26
	} else {
		goto L28
	}
L28:
	;
	v110 = *(*int32)(unsafe.Add(mBase, uint32(v59)+8))
	v111 = *(*int32)(unsafe.Add(mBase, uint32(v110)+56))
	v116 = int32(0)
	goto L29
L29:
	;
	v135 = *(*int32)(unsafe.Add(mBase, uint32(v99+v116<<(uint(int32(2))%32))))
	*(*int32)(unsafe.Add(mBase, uint32(v135)+40)) = v111
	v138 = *(*int32)(unsafe.Add(mBase, uint32(v58)+276))
	F_ExecARInsertTriggers(m, v57, v59, v135, int32(0), v138)
	mBase = m.M
	v140 = m.ExcPending
	if v140 != 0 {
		goto L21
	} else {
		goto L31
	}
L30:
	;
	v155 = v143
	goto L26
L31:
	;
	v142 = v116 + int32(1)
	v143 = *(*int32)(unsafe.Add(mBase, uint32(v22)+12))
	if v142 < v143 {
		v116 = v142
		goto L29
	} else {
		goto L32
	}
L32:
	;
	goto L30
L33:
	;
	goto L25
L34:
	;
	goto L33
L35:
	;
	v175 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_CopyMultiInsertInfoFlush[1])))
	if v175&int32(1) == int32(0) {
		goto L34
	} else {
		goto L36
	}
L36:
	;
	v180 = int32(_a_F_CopyMultiInsertInfoFlush_0)
	v182 = *(*int32)(unsafe.Add(mBase, _c_F_CopyMultiInsertInfoFlush[2]))
	v183 = int32(1)
	*(*int32)(unsafe.Add(mBase, _c_F_CopyMultiInsertInfoFlush[2])) = v182 + v183
	v186 = *(*int32)(unsafe.Add(mBase, uint32(v171)))
	*(*int32)(unsafe.Add(mBase, uint32(v171))) = v186 + v183
	v190 = int32(0)
	v192 = int32(_a_F_CopyMultiInsertInfoFlush_1)
	v193 = base.AtomicRmwOr32(m, v190, v192, v190)
	*(*int64)(unsafe.Add(mBase, uint32(v171+int32(16))+232)) = v166
	v201 = base.AtomicRmwOr32(m, v190, v192, v190)
	v202 = *(*int32)(unsafe.Add(mBase, uint32(v171)))
	*(*int32)(unsafe.Add(mBase, uint32(v171))) = v202 + v183
	v208 = *(*int32)(unsafe.Add(mBase, _c_F_CopyMultiInsertInfoFlush[2]))
	*(*int32)(unsafe.Add(mBase, _c_F_CopyMultiInsertInfoFlush[2])) = v208 - v183
	goto L34
L37:
	;
	goto L17
L38:
	;
	v244 = v234
	v247 = int32(0)
	goto L41
L39:
	;
	v293 = v234
	goto L40
L40:
	;
	v312 = v293
	v314 = v234
	goto L49
L41:
	;
	v262 = v55 + v244<<(uint(int32(2))%32)
	v263 = *(*int32)(unsafe.Add(mBase, uint32(v262)))
	v264 = *(*int32)(unsafe.Add(mBase, uint32(v263)+8))
	v265 = *(*int32)(unsafe.Add(mBase, uint32(v264)+12))
	m.T0[v265].(func(*base.Module, int32))(m, v263)
	mBase = m.M
	v267 = m.ExcPending
	if v267 != 0 {
		goto L21
	} else {
		goto L43
	}
L42:
	;
	if v233 == int32(0) {
		goto L14
	} else {
		goto L48
	}
L43:
	;
	v268 = *(*int32)(unsafe.Add(mBase, uint32(v262)+4))
	v269 = *(*int32)(unsafe.Add(mBase, uint32(v268)+8))
	v270 = *(*int32)(unsafe.Add(mBase, uint32(v269)+12))
	m.T0[v270].(func(*base.Module, int32))(m, v268)
	mBase = m.M
	v272 = m.ExcPending
	if v272 != 0 {
		goto L21
	} else {
		goto L44
	}
L44:
	;
	v273 = *(*int32)(unsafe.Add(mBase, uint32(v262)+8))
	v274 = *(*int32)(unsafe.Add(mBase, uint32(v273)+8))
	v275 = *(*int32)(unsafe.Add(mBase, uint32(v274)+12))
	m.T0[v275].(func(*base.Module, int32))(m, v273)
	mBase = m.M
	v277 = m.ExcPending
	if v277 != 0 {
		goto L21
	} else {
		goto L45
	}
L45:
	;
	v278 = *(*int32)(unsafe.Add(mBase, uint32(v262)+12))
	v279 = *(*int32)(unsafe.Add(mBase, uint32(v278)+8))
	v280 = *(*int32)(unsafe.Add(mBase, uint32(v279)+12))
	m.T0[v280].(func(*base.Module, int32))(m, v278)
	mBase = m.M
	v282 = m.ExcPending
	if v282 != 0 {
		goto L21
	} else {
		goto L46
	}
L46:
	;
	v283 = int32(4)
	v284 = v244 + v283
	v286 = v247 + v283
	if v286 != v56&int32(-4) {
		v244 = v284
		v247 = v286
		goto L41
	} else {
		goto L47
	}
L47:
	;
	goto L42
L48:
	;
	v293 = v284
	goto L40
L49:
	;
	v331 = *(*int32)(unsafe.Add(mBase, uint32(v55+v312<<(uint(int32(2))%32))))
	v332 = *(*int32)(unsafe.Add(mBase, uint32(v331)+8))
	v333 = *(*int32)(unsafe.Add(mBase, uint32(v332)+12))
	m.T0[v333].(func(*base.Module, int32))(m, v331)
	mBase = m.M
	v335 = m.ExcPending
	if v335 != 0 {
		goto L21
	} else {
		goto L51
	}
L50:
	;
	goto L14
L51:
	;
	v336 = int32(1)
	v339 = v314 + v336
	if v339 != v233 {
		v312 = v312 + v336
		v314 = v339
		goto L49
	} else {
		goto L52
	}
L52:
	;
	goto L50
L53:
	;
	v371 = F_MakePerTupleExprContext(m, v57)
	mBase = m.M
	v372 = m.ExcPending
	if v372 != 0 {
		goto L21
	} else {
		goto L56
	}
L54:
	;
	v373 = v368
	goto L55
L55:
	;
	v374 = int32(_a_F_CopyMultiInsertInfoFlush_2)
	v375 = *(*int32)(unsafe.Add(mBase, _c_F_CopyMultiInsertInfoFlush[3]))
	v377 = *(*int32)(unsafe.Add(mBase, uint32(v373)+20))
	*(*int32)(unsafe.Add(mBase, _c_F_CopyMultiInsertInfoFlush[3])) = v377
	v379 = *(*int32)(unsafe.Add(mBase, uint32(v59)+8))
	v380 = *(*int32)(unsafe.Add(mBase, uint32(v55)+4004))
	v381 = *(*int32)(unsafe.Add(mBase, uint32(v379)+188))
	v382 = *(*int32)(unsafe.Add(mBase, uint32(v381)+92))
	m.T0[v382].(func(*base.Module, int32, int32, int32, int32, int32, int32))(m, v379, v55, v56, v364, v363, v380)
	mBase = m.M
	v384 = m.ExcPending
	if v384 != 0 {
		goto L21
	} else {
		goto L57
	}
L56:
	;
	v373 = v371
	goto L55
L57:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_CopyMultiInsertInfoFlush[3])) = v375
	if int32(0) < v56 {
		goto L58
	} else {
		goto L59
	}
L58:
	;
	v390 = v55 + int32(4016)
	v395 = int32(0)
	goto L61
L59:
	;
	goto L60
L60:
	;
	v488 = *(*int64)(unsafe.Add(mBase, uint32(l2)))
	v490 = v488 + base.I64_extend_i32_s(v56)
	*(*int64)(unsafe.Add(mBase, uint32(l2))) = v490
	v495 = *(*int32)(unsafe.Add(mBase, _c_F_CopyMultiInsertInfoFlush[0]))
	if v495 == int32(0) {
		goto L79
	} else {
		goto L80
	}
L61:
	;
	v411 = *(*int32)(unsafe.Add(mBase, uint32(v59)+12))
	if int32(0) < v411 {
		goto L64
	} else {
		goto L65
	}
L62:
	;
	goto L60
L63:
	;
	v461 = *(*int32)(unsafe.Add(mBase, uint32(v55+v395<<(uint(int32(2))%32))))
	v462 = *(*int32)(unsafe.Add(mBase, uint32(v461)+8))
	v463 = *(*int32)(unsafe.Add(mBase, uint32(v462)+12))
	m.T0[v463].(func(*base.Module, int32))(m, v461)
	mBase = m.M
	v465 = m.ExcPending
	if v465 != 0 {
		goto L21
	} else {
		goto L76
	}
L64:
	;
	v417 = *(*int64)(unsafe.Add(mBase, uint32(v390+v395<<(uint(int32(3))%32))))
	*(*int64)(unsafe.Add(mBase, uint32(v58)+192)) = v417
	v419 = int32(0)
	v422 = v55 + v395<<(uint(int32(2))%32)
	v423 = *(*int32)(unsafe.Add(mBase, uint32(v422)))
	v426 = F_ExecInsertIndexTuples(m, v59, v57, v419, v423, v419, v419)
	mBase = m.M
	v427 = m.ExcPending
	if v427 != 0 {
		goto L21
	} else {
		goto L67
	}
L65:
	;
	goto L66
L66:
	;
	v434 = *(*int32)(unsafe.Add(mBase, uint32(v59)+52))
	if v434 == int32(0) {
		goto L63
	} else {
		goto L70
	}
L67:
	;
	v428 = *(*int32)(unsafe.Add(mBase, uint32(v422)))
	v429 = *(*int32)(unsafe.Add(mBase, uint32(v58)+276))
	F_ExecARInsertTriggers(m, v57, v59, v428, v426, v429)
	mBase = m.M
	v431 = m.ExcPending
	if v431 != 0 {
		goto L21
	} else {
		goto L68
	}
L68:
	;
	F_list_free(m, v426)
	mBase = m.M
	v433 = m.ExcPending
	if v433 != 0 {
		goto L21
	} else {
		goto L69
	}
L69:
	;
	goto L63
L70:
	;
	v437 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v434)+9)))
	if v437 == int32(0) {
		goto L71
	} else {
		goto L72
	}
L71:
	;
	v440 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v434)+25)))
	if v440 != int32(1) {
		goto L63
	} else {
		goto L74
	}
L72:
	;
	goto L73
L73:
	;
	v446 = *(*int64)(unsafe.Add(mBase, uint32(v390+v395<<(uint(int32(3))%32))))
	*(*int64)(unsafe.Add(mBase, uint32(v58)+192)) = v446
	v451 = *(*int32)(unsafe.Add(mBase, uint32(v55+v395<<(uint(int32(2))%32))))
	v453 = *(*int32)(unsafe.Add(mBase, uint32(v58)+276))
	F_ExecARInsertTriggers(m, v57, v59, v451, int32(0), v453)
	mBase = m.M
	v455 = m.ExcPending
	if v455 != 0 {
		goto L21
	} else {
		goto L75
	}
L74:
	;
	goto L73
L75:
	;
	goto L63
L76:
	;
	v467 = v395 + int32(1)
	if v467 != v56 {
		v395 = v467
		goto L61
	} else {
		goto L77
	}
L77:
	;
	goto L62
L78:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v58)+192)) = v367
	*(*uint8)(unsafe.Add(mBase, uint32(v58)+320)) = uint8(v362)
	goto L10
L79:
	;
	goto L78
L80:
	;
	v499 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_CopyMultiInsertInfoFlush[1])))
	if v499&int32(1) == int32(0) {
		goto L79
	} else {
		goto L81
	}
L81:
	;
	v504 = int32(_a_F_CopyMultiInsertInfoFlush_0)
	v506 = *(*int32)(unsafe.Add(mBase, _c_F_CopyMultiInsertInfoFlush[2]))
	v507 = int32(1)
	*(*int32)(unsafe.Add(mBase, _c_F_CopyMultiInsertInfoFlush[2])) = v506 + v507
	v510 = *(*int32)(unsafe.Add(mBase, uint32(v495)))
	*(*int32)(unsafe.Add(mBase, uint32(v495))) = v510 + v507
	v514 = int32(0)
	v516 = int32(_a_F_CopyMultiInsertInfoFlush_1)
	v517 = base.AtomicRmwOr32(m, v514, v516, v514)
	*(*int64)(unsafe.Add(mBase, uint32(v495+int32(16))+232)) = v490
	v525 = base.AtomicRmwOr32(m, v514, v516, v514)
	v526 = *(*int32)(unsafe.Add(mBase, uint32(v495)))
	*(*int32)(unsafe.Add(mBase, uint32(v495))) = v526 + v507
	v532 = *(*int32)(unsafe.Add(mBase, _c_F_CopyMultiInsertInfoFlush[2]))
	*(*int32)(unsafe.Add(mBase, _c_F_CopyMultiInsertInfoFlush[2])) = v532 - v507
	goto L79
L82:
	;
	goto L9
L83:
	;
	v590 = v584
	goto L84
L84:
	;
	v606 = *(*int32)(unsafe.Add(mBase, uint32(v590)+4))
	if v606 < int32(33) {
		goto L1
	} else {
		goto L86
	}
L85:
	;
	goto L1
L86:
	;
	v609 = *(*int32)(unsafe.Add(mBase, uint32(v590)+12))
	v610 = *(*int32)(unsafe.Add(mBase, uint32(v609)))
	v611 = *(*int32)(unsafe.Add(mBase, uint32(v610)+4000))
	if l1 == v611 {
		goto L87
	} else {
		goto L88
	}
L87:
	;
	v613 = F_list_delete_first(m, v590)
	mBase = m.M
	v614 = m.ExcPending
	if v614 != 0 {
		goto L21
	} else {
		goto L90
	}
L88:
	;
	v623 = v610
	v624 = v611
	goto L89
L89:
	;
	v625 = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v624)+208)) = v625
	v627 = *(*int32)(unsafe.Add(mBase, uint32(v624)+84))
	if v627 == v625 {
		goto L92
	} else {
		goto L93
	}
L90:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0))) = v613
	v616 = F_lappend(m, v613, v610)
	mBase = m.M
	v617 = m.ExcPending
	if v617 != 0 {
		goto L21
	} else {
		goto L91
	}
L91:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0))) = v616
	v619 = *(*int32)(unsafe.Add(mBase, uint32(v616)+12))
	v620 = *(*int32)(unsafe.Add(mBase, uint32(v619)))
	v621 = *(*int32)(unsafe.Add(mBase, uint32(v620)+4000))
	v623 = v620
	v624 = v621
	goto L89
L92:
	;
	v630 = *(*int32)(unsafe.Add(mBase, uint32(v623)+4004))
	F_FreeBulkInsertState(m, v630)
	mBase = m.M
	v632 = m.ExcPending
	if v632 != 0 {
		goto L21
	} else {
		goto L95
	}
L93:
	;
	goto L94
L94:
	;
	v637 = int32(0)
	goto L96
L95:
	;
	goto L94
L96:
	;
	v656 = *(*int32)(unsafe.Add(mBase, uint32(v623+v637<<(uint(int32(2))%32))))
	if v656 != 0 {
		goto L98
	} else {
		goto L99
	}
L97:
	;
	v664 = *(*int32)(unsafe.Add(mBase, uint32(v624)+84))
	if v664 != 0 {
		goto L103
	} else {
		goto L104
	}
L98:
	;
	F_ExecDropSingleTupleTableSlot(m, v656)
	mBase = m.M
	v658 = m.ExcPending
	if v658 != 0 {
		goto L21
	} else {
		goto L101
	}
L99:
	;
	goto L100
L100:
	;
	goto L97
L101:
	;
	v660 = v637 + int32(1)
	if v660 != int32(1000) {
		v637 = v660
		goto L96
	} else {
		goto L102
	}
L102:
	;
	goto L100
L103:
	;
	F_pfree(m, v623)
	mBase = m.M
	v678 = m.ExcPending
	if v678 != 0 {
		goto L21
	} else {
		goto L108
	}
L104:
	;
	v665 = *(*int32)(unsafe.Add(mBase, uint32(v624)+8))
	v666 = *(*int32)(unsafe.Add(mBase, uint32(v665)+188))
	if v666 == int32(0) {
		goto L103
	} else {
		goto L105
	}
L105:
	;
	v669 = *(*int32)(unsafe.Add(mBase, uint32(v666)+108))
	if v669 == int32(0) {
		goto L103
	} else {
		goto L106
	}
L106:
	;
	v672 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
	m.T0[v669].(func(*base.Module, int32, int32))(m, v665, v672)
	mBase = m.M
	v674 = m.ExcPending
	if v674 != 0 {
		goto L21
	} else {
		goto L107
	}
L107:
	;
	goto L103
L108:
	;
	v679 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v680 = F_list_delete_first(m, v679)
	mBase = m.M
	v681 = m.ExcPending
	if v681 != 0 {
		goto L21
	} else {
		goto L109
	}
L109:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0))) = v680
	if v680 != 0 {
		v590 = v680
		goto L84
	} else {
		goto L110
	}
L110:
	;
	goto L85
}
func F_GetMultiXactIdMembers(m *base.Module, l0 int32, l1 int32, l2 int32) int32 {
	mBase := m.M
	_ = mBase
	var v4 int32
	_ = v4
	var v16 int32
	_ = v16
	var v18 int32
	_ = v18
	var v25 int32
	_ = v25
	var v34 int32
	_ = v34
	var v48 int32
	_ = v48
	var v51 int32
	_ = v51
	var v52 int32
	_ = v52
	var v54 int32
	_ = v54
	var v55 int32
	_ = v55
	var v58 int32
	_ = v58
	var v63 int32
	_ = v63
	var v65 int32
	_ = v65
	var v66 int32
	_ = v66
	var v68 int32
	_ = v68
	var v71 int32
	_ = v71
	var v74 int32
	_ = v74
	var v78 int32
	_ = v78
	var v87 int32
	_ = v87
	var v90 int32
	_ = v90
	var v109 int32
	_ = v109
	var v111 int32
	_ = v111
	var v115 int32
	_ = v115
	var v119 int32
	_ = v119
	var v123 int32
	_ = v123
	var v124 int32
	_ = v124
	var v126 int32
	_ = v126
	var v127 int32
	_ = v127
	var v129 int32
	_ = v129
	var v131 int32
	_ = v131
	var v132 int32
	_ = v132
	var v136 int32
	_ = v136
	var v148 int32
	_ = v148
	var v150 int32
	_ = v150
	var v154 int32
	_ = v154
	var v162 int32
	_ = v162
	var v163 int32
	_ = v163
	var v164 int32
	_ = v164
	var v168 int32
	_ = v168
	var v169 int32
	_ = v169
	var v173 int32
	_ = v173
	var v174 int32
	_ = v174
	var v175 int32
	_ = v175
	var v176 int32
	_ = v176
	var v178 int32
	_ = v178
	var v185 int32
	_ = v185
	var v187 int32
	_ = v187
	var v200 int32
	_ = v200
	var v204 int32
	_ = v204
	var v205 int32
	_ = v205
	var v209 int32
	_ = v209
	var v222 int32
	_ = v222
	var v224 int32
	_ = v224
	var v230 int32
	_ = v230
	var v234 int32
	_ = v234
	var v250 int32
	_ = v250
	var v253 int32
	_ = v253
	var v255 int32
	_ = v255
	var v259 int32
	_ = v259
	var v266 int32
	_ = v266
	var v270 int32
	_ = v270
	var v271 int32
	_ = v271
	var v273 int32
	_ = v273
	var v274 int32
	_ = v274
	var v275 int32
	_ = v275
	var v277 int32
	_ = v277
	var v281 int32
	_ = v281
	var v289 int32
	_ = v289
	var v290 int32
	_ = v290
	var v292 int32
	_ = v292
	var v294 int32
	_ = v294
	var v295 int32
	_ = v295
	var v298 int32
	_ = v298
	var v300 int32
	_ = v300
	var v301 int32
	_ = v301
	var v307 int32
	_ = v307
	var v308 int32
	_ = v308
	var v310 int32
	_ = v310
	var v311 int32
	_ = v311
	var v315 int32
	_ = v315
	var v321 int64
	_ = v321
	var v324 int32
	_ = v324
	var v325 int32
	_ = v325
	var v326 int32
	_ = v326
	var v328 int32
	_ = v328
	var v333 int32
	_ = v333
	var v336 int32
	_ = v336
	var v338 int32
	_ = v338
	var v339 int32
	_ = v339
	var v342 int32
	_ = v342
	var v345 int32
	_ = v345
	var v347 int32
	_ = v347
	var v348 int32
	_ = v348
	var v349 int32
	_ = v349
	var v354 int32
	_ = v354
	var v355 int32
	_ = v355
	var v357 int32
	_ = v357
	var v358 int32
	_ = v358
	var v362 int32
	_ = v362
	var v363 int32
	_ = v363
	var v364 int32
	_ = v364
	var v365 int32
	_ = v365
	var v370 int64
	_ = v370
	var v372 int32
	_ = v372
	var v377 int64
	_ = v377
	var v381 int32
	_ = v381
	var v384 int32
	_ = v384
	var v385 int32
	_ = v385
	var v386 int32
	_ = v386
	var v388 int32
	_ = v388
	var v391 int32
	_ = v391
	var v397 int32
	_ = v397
	var v401 int64
	_ = v401
	var v402 int64
	_ = v402
	var v403 int32
	_ = v403
	var v409 int64
	_ = v409
	var v412 int32
	_ = v412
	var v417 int32
	_ = v417
	var v418 int32
	_ = v418
	var v420 int64
	_ = v420
	var v421 int64
	_ = v421
	var v425 int32
	_ = v425
	var v428 int32
	_ = v428
	var v430 int32
	_ = v430
	var v431 int32
	_ = v431
	var v432 int32
	_ = v432
	var v437 int32
	_ = v437
	var v438 int32
	_ = v438
	var v439 int32
	_ = v439
	var v441 int32
	_ = v441
	var v442 int64
	_ = v442
	var v443 int32
	_ = v443
	var v445 int32
	_ = v445
	var v447 int32
	_ = v447
	var v448 int32
	_ = v448
	var v452 int32
	_ = v452
	var v456 int64
	_ = v456
	var v460 int32
	_ = v460
	var v462 int32
	_ = v462
	var v464 int32
	_ = v464
	var v474 int32
	_ = v474
	var v477 int32
	_ = v477
	var v478 int32
	_ = v478
	var v480 int32
	_ = v480
	var v486 int32
	_ = v486
	var v487 int32
	_ = v487
	var v502 int32
	_ = v502
	var v519 int32
	_ = v519
	var v522 int32
	_ = v522
	var v526 int32
	_ = v526
	var v531 int32
	_ = v531
	var v535 int32
	_ = v535
	var v538 int32
	_ = v538
	var v544 int32
	_ = v544
	var v549 int32
	_ = v549
	var v553 int32
	_ = v553
	var v556 int32
	_ = v556
	var v557 int32
	_ = v557
	var v563 int32
	_ = v563
	var v568 int32
	_ = v568
	var v572 int32
	_ = v572
	var v575 int32
	_ = v575
	var v576 int32
	_ = v576
	var v582 int32
	_ = v582
	var v587 int32
	_ = v587
	var v591 int32
	_ = v591
	var v594 int32
	_ = v594
	var v596 int32
	_ = v596
	var v602 int32
	_ = v602
	var v607 int32
	_ = v607
	var v611 int32
	_ = v611
	var v614 int32
	_ = v614
	var v617 int32
	_ = v617
	var v623 int32
	_ = v623
	var v628 int32
	_ = v628
	var v632 int32
	_ = v632
	var v635 int32
	_ = v635
	var v637 int32
	_ = v637
	var v643 int32
	_ = v643
	var v648 int32
	_ = v648
	v4 = int32(0)
	v16 = m.G0
	v18 = v16 - int32(144)
	m.G0 = v18
	*(*int32)(unsafe.Add(mBase, uint32(v18)+140)) = l0
	if l0 == v4 {
		v486 = int32(-1)
		v487 = v4
		goto L9
	} else {
		goto L10
	}
L1:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v632 = m.ExcPending
	if v632 != 0 {
		goto L18
	} else {
		goto L136
	}
L2:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v611 = m.ExcPending
	if v611 != 0 {
		goto L18
	} else {
		goto L132
	}
L3:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v591 = m.ExcPending
	if v591 != 0 {
		goto L18
	} else {
		goto L128
	}
L4:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v572 = m.ExcPending
	if v572 != 0 {
		goto L18
	} else {
		goto L124
	}
L5:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v553 = m.ExcPending
	if v553 != 0 {
		goto L18
	} else {
		goto L120
	}
L6:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v535 = m.ExcPending
	if v535 != 0 {
		goto L18
	} else {
		goto L116
	}
L7:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v519 = m.ExcPending
	if v519 != 0 {
		goto L18
	} else {
		goto L112
	}
L8:
	;
	m.G0 = v18 + int32(144)
	return v502
L9:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l1))) = v487
	v502 = v486
	goto L8
L10:
	;
	v25 = *(*int32)(unsafe.Add(mBase, _c_F_GetMultiXactIdMembers[0]))
	if base.B2i32(v25 == int32(0))|base.B2i32(v25 == int32(_a_F_GetMultiXactIdMembers_0)) != 0 {
		goto L11
	} else {
		goto L12
	}
L11:
	;
	v109 = *(*int32)(unsafe.Add(mBase, _c_F_GetMultiXactIdMembers[1]))
	v111 = *(*int32)(unsafe.Add(mBase, _c_F_GetMultiXactIdMembers[2]))
	v115 = *(*int32)(unsafe.Add(mBase, uint32(v109+v111<<(uint(int32(2))%32))))
	if v115 == int32(0) {
		goto L31
	} else {
		goto L32
	}
L12:
	;
	v34 = v25
	goto L13
L13:
	;
	v48 = *(*int32)(unsafe.Add(mBase, uint32(v34-int32(8))))
	if l0 == v48 {
		goto L15
	} else {
		goto L16
	}
L14:
	;
	goto L11
L15:
	;
	v51 = v34 - int32(4)
	v52 = *(*int32)(unsafe.Add(mBase, uint32(v51)))
	v54 = v52 << (uint(int32(3)) % 32)
	v55 = F_palloc(m, v54)
	mBase = m.M
	v58 = m.ExcPending
	if v58 != 0 {
		goto L18
	} else {
		goto L19
	}
L16:
	;
	goto L17
L17:
	;
	v90 = *(*int32)(unsafe.Add(mBase, uint32(v34)+4))
	if v90 != int32(_a_F_GetMultiXactIdMembers_0) {
		v34 = v90
		goto L13
	} else {
		goto L30
	}
L18:
	;
	return int32(0)
L19:
	;
	if v54 != 0 {
		goto L20
	} else {
		goto L21
	}
L20:
	;
	base.MemoryCopy(m, v55, v34+int32(8), v54)
	goto L22
L21:
	;
	goto L22
L22:
	;
	v63 = *(*int32)(unsafe.Add(mBase, _c_F_GetMultiXactIdMembers[0]))
	if v34 != v63 {
		goto L23
	} else {
		goto L24
	}
L23:
	;
	v65 = *(*int32)(unsafe.Add(mBase, uint32(v34)))
	v66 = *(*int32)(unsafe.Add(mBase, uint32(v34)+4))
	*(*int32)(unsafe.Add(mBase, uint32(v65)+4)) = v66
	v68 = *(*int32)(unsafe.Add(mBase, uint32(v34)))
	*(*int32)(unsafe.Add(mBase, uint32(v66))) = v68
	v71 = *(*int32)(unsafe.Add(mBase, _c_F_GetMultiXactIdMembers[0]))
	if v71 == int32(0) {
		goto L26
	} else {
		goto L27
	}
L24:
	;
	goto L25
L25:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l1))) = v55
	v87 = *(*int32)(unsafe.Add(mBase, uint32(v51)))
	if int32(0) <= v87 {
		v502 = v87
		goto L8
	} else {
		goto L29
	}
L26:
	;
	v74 = int32(_a_F_GetMultiXactIdMembers_0)
	*(*int32)(unsafe.Add(mBase, _c_F_GetMultiXactIdMembers[3])) = v74
	v78 = v74
	goto L28
L27:
	;
	v78 = v71
	goto L28
L28:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v34))) = int32(_a_F_GetMultiXactIdMembers_0)
	*(*int32)(unsafe.Add(mBase, uint32(v34)+4)) = v78
	*(*int32)(unsafe.Add(mBase, uint32(v78))) = v34
	*(*int32)(unsafe.Add(mBase, _c_F_GetMultiXactIdMembers[0])) = v34
	goto L25
L29:
	;
	goto L11
L30:
	;
	goto L14
L31:
	;
	v119 = *(*int32)(unsafe.Add(mBase, _c_F_GetMultiXactIdMembers[4]))
	v123 = F_LWLockAcquire(m, v119+int32(1664), int32(0))
	mBase = m.M
	v124 = m.ExcPending
	if v124 != 0 {
		goto L18
	} else {
		goto L34
	}
L32:
	;
	goto L33
L33:
	;
	if l2 != 0 {
		goto L64
	} else {
		goto L65
	}
L34:
	;
	v126 = *(*int32)(unsafe.Add(mBase, _c_F_GetMultiXactIdMembers[5]))
	v127 = *(*int32)(unsafe.Add(mBase, uint32(v126)))
	v129 = *(*int32)(unsafe.Add(mBase, _c_F_GetMultiXactIdMembers[6]))
	v131 = *(*int32)(unsafe.Add(mBase, _c_F_GetMultiXactIdMembers[7]))
	v132 = v129 + v131
	if v132 <= int32(0) {
		v209 = v127
		goto L35
	} else {
		goto L36
	}
L35:
	;
	v222 = *(*int32)(unsafe.Add(mBase, _c_F_GetMultiXactIdMembers[1]))
	v224 = *(*int32)(unsafe.Add(mBase, _c_F_GetMultiXactIdMembers[2]))
	*(*int32)(unsafe.Add(mBase, uint32(v222+v224<<(uint(int32(2))%32)))) = v209
	v230 = *(*int32)(unsafe.Add(mBase, _c_F_GetMultiXactIdMembers[4]))
	F_LWLockRelease(m, v230+int32(1664))
	mBase = m.M
	v234 = m.ExcPending
	if v234 != 0 {
		goto L18
	} else {
		goto L63
	}
L36:
	;
	v136 = *(*int32)(unsafe.Add(mBase, _c_F_GetMultiXactIdMembers[8]))
	if v132 == int32(1) {
		goto L38
	} else {
		goto L39
	}
L37:
	;
	v200 = *(*int32)(unsafe.Add(mBase, uint32(v136+v187<<(uint(int32(2))%32))))
	if v200-v185 < int32(0) {
		goto L57
	} else {
		goto L58
	}
L38:
	;
	v185 = v127
	v187 = int32(0)
	goto L37
L39:
	;
	goto L40
L40:
	;
	v148 = v127
	v150 = int32(0)
	v154 = v4
	goto L41
L41:
	;
	v162 = v136 + v150<<(uint(int32(2))%32)
	v163 = *(*int32)(unsafe.Add(mBase, uint32(v162)+4))
	v164 = *(*int32)(unsafe.Add(mBase, uint32(v162)))
	if v164-v148 < int32(0) {
		goto L43
	} else {
		goto L44
	}
L42:
	;
	if v132&int32(1) == int32(0) {
		v209 = v174
		goto L35
	} else {
		goto L56
	}
L43:
	;
	v168 = v164
	goto L45
L44:
	;
	v168 = v148
	goto L45
L45:
	;
	if v164 != 0 {
		goto L46
	} else {
		goto L47
	}
L46:
	;
	v169 = v168
	goto L48
L47:
	;
	v169 = v148
	goto L48
L48:
	;
	if v163-v169 < int32(0) {
		goto L49
	} else {
		goto L50
	}
L49:
	;
	v173 = v163
	goto L51
L50:
	;
	v173 = v169
	goto L51
L51:
	;
	if v163 != 0 {
		goto L52
	} else {
		goto L53
	}
L52:
	;
	v174 = v173
	goto L54
L53:
	;
	v174 = v169
	goto L54
L54:
	;
	v175 = int32(2)
	v176 = v150 + v175
	v178 = v154 + v175
	if v178 != v132&int32(2147483646) {
		v148 = v174
		v150 = v176
		v154 = v178
		goto L41
	} else {
		goto L55
	}
L55:
	;
	goto L42
L56:
	;
	v185 = v174
	v187 = v176
	goto L37
L57:
	;
	v204 = v200
	goto L59
L58:
	;
	v204 = v185
	goto L59
L59:
	;
	if v200 != 0 {
		goto L60
	} else {
		goto L61
	}
L60:
	;
	v205 = v204
	goto L62
L61:
	;
	v205 = v185
	goto L62
L62:
	;
	v209 = v205
	goto L35
L63:
	;
	goto L33
L64:
	;
	v250 = int32(0)
	v253 = *(*int32)(unsafe.Add(mBase, _c_F_GetMultiXactIdMembers[1]))
	v255 = *(*int32)(unsafe.Add(mBase, _c_F_GetMultiXactIdMembers[2]))
	v259 = *(*int32)(unsafe.Add(mBase, uint32(v253+v255<<(uint(int32(2))%32))))
	if l0-v259 < v250 {
		v486 = int32(-1)
		v487 = v250
		goto L9
	} else {
		goto L67
	}
L65:
	;
	goto L66
L66:
	;
	v266 = *(*int32)(unsafe.Add(mBase, _c_F_GetMultiXactIdMembers[4]))
	v270 = F_LWLockAcquire(m, v266+int32(1664), int32(1))
	mBase = m.M
	v271 = m.ExcPending
	if v271 != 0 {
		goto L18
	} else {
		goto L68
	}
L67:
	;
	goto L66
L68:
	;
	v273 = *(*int32)(unsafe.Add(mBase, _c_F_GetMultiXactIdMembers[5]))
	v274 = *(*int32)(unsafe.Add(mBase, uint32(v273)))
	v275 = *(*int32)(unsafe.Add(mBase, uint32(v273)+20))
	v277 = *(*int32)(unsafe.Add(mBase, _c_F_GetMultiXactIdMembers[4]))
	F_LWLockRelease(m, v277+int32(1664))
	mBase = m.M
	v281 = m.ExcPending
	if v281 != 0 {
		goto L18
	} else {
		goto L69
	}
L69:
	;
	if l0-v275 < int32(0) {
		goto L7
	} else {
		goto L70
	}
L70:
	;
	if int32(0) <= l0-v274 {
		goto L6
	} else {
		goto L71
	}
L71:
	;
	v289 = *(*int32)(unsafe.Add(mBase, _c_F_GetMultiXactIdMembers[9]))
	v290 = *(*int32)(unsafe.Add(mBase, uint32(v289)+28))
	v292 = int32(base.Ui32(l0) >> (uint(int32(10)) % 32))
	v294 = int32(*(*uint16)(unsafe.Add(mBase, _c_F_GetMultiXactIdMembers[10])))
	v295 = base.I32_rem_u_s(v292, v294)
	v298 = v290 + v295<<(uint(int32(7))%32)
	v300 = F_LWLockAcquire(m, v298, int32(0))
	mBase = m.M
	v301 = m.ExcPending
	if v301 != 0 {
		goto L18
	} else {
		goto L72
	}
L72:
	;
	v307 = F_SimpleLruReadPage(m, int32(_a_F_GetMultiXactIdMembers_1), base.I64_extend_i32_u(v292), int32(1), v18+int32(140))
	mBase = m.M
	v308 = m.ExcPending
	if v308 != 0 {
		goto L18
	} else {
		goto L73
	}
L73:
	;
	v310 = *(*int32)(unsafe.Add(mBase, _c_F_GetMultiXactIdMembers[9]))
	v311 = *(*int32)(unsafe.Add(mBase, uint32(v310)+4))
	v315 = *(*int32)(unsafe.Add(mBase, uint32(v311+v307<<(uint(int32(2))%32))))
	v321 = *(*int64)(unsafe.Add(mBase, uint32(v315+l0&int32(1023)<<(uint(int32(3))%32))))
	if v321 == int64(0) {
		goto L5
	} else {
		goto L74
	}
L74:
	;
	v324 = *(*int32)(unsafe.Add(mBase, uint32(v18)+140))
	v325 = int32(1)
	v326 = v324 + v325
	if v326 != 0 {
		goto L75
	} else {
		goto L76
	}
L75:
	;
	v328 = v326
	goto L77
L76:
	;
	v328 = v325
	goto L77
L77:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v18)+120)) = v328
	v333 = int32(base.Ui32(v328) >> (uint(int32(10)) % 32))
	if v292 == v333 {
		goto L79
	} else {
		goto L80
	}
L78:
	;
	v370 = *(*int64)(unsafe.Add(mBase, uint32(v364+v328&int32(1023)<<(uint(int32(3))%32))))
	F_LWLockRelease(m, v363)
	mBase = m.M
	v372 = m.ExcPending
	if v372 != 0 {
		goto L18
	} else {
		goto L89
	}
L79:
	;
	v363 = v298
	v364 = v315
	v365 = v307
	goto L78
L80:
	;
	goto L81
L81:
	;
	v336 = *(*int32)(unsafe.Add(mBase, uint32(v310)+28))
	v338 = int32(*(*uint16)(unsafe.Add(mBase, _c_F_GetMultiXactIdMembers[10])))
	v339 = base.I32_rem_u_s(v333, v338)
	v342 = v336 + v339<<(uint(int32(7))%32)
	if v298 == v342 {
		goto L83
	} else {
		goto L84
	}
L82:
	;
	v354 = F_SimpleLruReadPage(m, int32(_a_F_GetMultiXactIdMembers_1), base.I64_extend_i32_u(v333), int32(1), v18+int32(120))
	mBase = m.M
	v355 = m.ExcPending
	if v355 != 0 {
		goto L18
	} else {
		goto L88
	}
L83:
	;
	v349 = v298
	goto L82
L84:
	;
	goto L85
L85:
	;
	F_LWLockRelease(m, v298)
	mBase = m.M
	v345 = m.ExcPending
	if v345 != 0 {
		goto L18
	} else {
		goto L86
	}
L86:
	;
	v347 = F_LWLockAcquire(m, v342, int32(0))
	mBase = m.M
	v348 = m.ExcPending
	if v348 != 0 {
		goto L18
	} else {
		goto L87
	}
L87:
	;
	v349 = v342
	goto L82
L88:
	;
	v357 = *(*int32)(unsafe.Add(mBase, _c_F_GetMultiXactIdMembers[9]))
	v358 = *(*int32)(unsafe.Add(mBase, uint32(v357)+4))
	v362 = *(*int32)(unsafe.Add(mBase, uint32(v358+v354<<(uint(int32(2))%32))))
	v363 = v349
	v364 = v362
	v365 = v354
	goto L78
L89:
	;
	if v370 == int64(0) {
		goto L4
	} else {
		goto L90
	}
L90:
	;
	if v370 == v321 {
		goto L3
	} else {
		goto L91
	}
L91:
	;
	if base.Ui64(v370) < base.Ui64(v321) {
		goto L2
	} else {
		goto L92
	}
L92:
	;
	v377 = v370 - v321
	if base.Ui64(int64(2147483648)) <= base.Ui64(v377) {
		goto L1
	} else {
		goto L93
	}
L93:
	;
	v381 = base.I32_wrap_i64(v377)
	v384 = F_palloc(m, v381<<(uint(int32(3))%32))
	mBase = m.M
	v385 = m.ExcPending
	if v385 != 0 {
		goto L18
	} else {
		goto L94
	}
L94:
	;
	v386 = int32(0)
	v388 = v386
	v391 = v386
	v397 = v365
	v401 = v321
	v402 = int64(-1)
	goto L95
L95:
	;
	v403 = base.I32_wrap_i64(v401)
	v409 = base.I64_div_u_s(v401, int64(1636))
	if v402 != v409 {
		goto L97
	} else {
		goto L98
	}
L96:
	;
	F_LWLockRelease(m, v439)
	mBase = m.M
	v477 = m.ExcPending
	if v477 != 0 {
		goto L18
	} else {
		goto L110
	}
L97:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v18)+128)) = v401
	v412 = *(*int32)(unsafe.Add(mBase, uint32(v18)+140))
	*(*int32)(unsafe.Add(mBase, uint32(v18)+120)) = v412
	*(*int32)(unsafe.Add(mBase, uint32(v18)+124)) = int32(0)
	v417 = *(*int32)(unsafe.Add(mBase, _c_F_GetMultiXactIdMembers[11]))
	v418 = *(*int32)(unsafe.Add(mBase, uint32(v417)+28))
	v420 = int64(*(*uint16)(unsafe.Add(mBase, _c_F_GetMultiXactIdMembers[12])))
	v421 = base.I64_rem_u_s(v409, v420)
	v425 = v418 + base.I32_wrap_i64(v421)<<(uint(int32(7))%32)
	if v388 != v425 {
		goto L100
	} else {
		goto L101
	}
L98:
	;
	v439 = v388
	v441 = v397
	v442 = v402
	goto L99
L99:
	;
	v443 = int32(3)
	v445 = v384 + v391<<(uint(v443)%32)
	v447 = *(*int32)(unsafe.Add(mBase, _c_F_GetMultiXactIdMembers[11]))
	v448 = *(*int32)(unsafe.Add(mBase, uint32(v447)+4))
	v452 = *(*int32)(unsafe.Add(mBase, uint32(v448+v441<<(uint(int32(2))%32))))
	v456 = base.I64_rem_u_s(int64(base.Ui64(v401)>>(uint(int64(2))%64)), int64(409))
	v460 = v452 + base.I32_wrap_i64(v456)*int32(20)
	v462 = *(*int32)(unsafe.Add(mBase, uint32(v403<<(uint(int32(2))%32)&int32(12)+v460)+4))
	*(*int32)(unsafe.Add(mBase, uint32(v445))) = v462
	v464 = *(*int32)(unsafe.Add(mBase, uint32(v460)))
	*(*int32)(unsafe.Add(mBase, uint32(v445)+4)) = int32(base.Ui32(v464)>>(uint(v403<<(uint(v443)%32))%32)) & int32(255)
	v474 = v391 + int32(1)
	if v474 != v381 {
		v388 = v439
		v391 = v474
		v397 = v441
		v401 = v401 + int64(1)
		v402 = v442
		goto L95
	} else {
		goto L109
	}
L100:
	;
	if v388 != 0 {
		goto L103
	} else {
		goto L104
	}
L101:
	;
	v432 = v388
	goto L102
L102:
	;
	v437 = F_SimpleLruReadPage(m, int32(_a_F_GetMultiXactIdMembers_2), v409, int32(1), v18+int32(120))
	mBase = m.M
	v438 = m.ExcPending
	if v438 != 0 {
		goto L18
	} else {
		goto L108
	}
L103:
	;
	F_LWLockRelease(m, v388)
	mBase = m.M
	v428 = m.ExcPending
	if v428 != 0 {
		goto L18
	} else {
		goto L106
	}
L104:
	;
	goto L105
L105:
	;
	v430 = F_LWLockAcquire(m, v425, int32(0))
	mBase = m.M
	v431 = m.ExcPending
	if v431 != 0 {
		goto L18
	} else {
		goto L107
	}
L106:
	;
	goto L105
L107:
	;
	v432 = v425
	goto L102
L108:
	;
	v439 = v432
	v441 = v437
	v442 = v409
	goto L99
L109:
	;
	goto L96
L110:
	;
	v478 = *(*int32)(unsafe.Add(mBase, uint32(v18)+140))
	F_mXactCachePut(m, v478, v381, v384)
	mBase = m.M
	v480 = m.ExcPending
	if v480 != 0 {
		goto L18
	} else {
		goto L111
	}
L111:
	;
	v486 = v381
	v487 = v384
	goto L9
L112:
	;
	F_errcode(m, int32(2600))
	mBase = m.M
	v522 = m.ExcPending
	if v522 != 0 {
		goto L18
	} else {
		goto L113
	}
L113:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v18))) = l0
	F_errmsg(m, int32(_a_F_GetMultiXactIdMembers_3), v18)
	mBase = m.M
	v526 = m.ExcPending
	if v526 != 0 {
		goto L18
	} else {
		goto L114
	}
L114:
	;
	F_errfinish(m, int32(_a_F_GetMultiXactIdMembers_4), int32(1246), int32(_a_F_GetMultiXactIdMembers_5))
	mBase = m.M
	v531 = m.ExcPending
	if v531 != 0 {
		goto L18
	} else {
		goto L115
	}
L115:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L116:
	;
	F_errcode(m, int32(2600))
	mBase = m.M
	v538 = m.ExcPending
	if v538 != 0 {
		goto L18
	} else {
		goto L117
	}
L117:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v18)+112)) = l0
	F_errmsg(m, int32(_a_F_GetMultiXactIdMembers_6), v18+int32(112))
	mBase = m.M
	v544 = m.ExcPending
	if v544 != 0 {
		goto L18
	} else {
		goto L118
	}
L118:
	;
	F_errfinish(m, int32(_a_F_GetMultiXactIdMembers_4), int32(1252), int32(_a_F_GetMultiXactIdMembers_5))
	mBase = m.M
	v549 = m.ExcPending
	if v549 != 0 {
		goto L18
	} else {
		goto L119
	}
L119:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L120:
	;
	F_errcode(m, int32(16779816))
	mBase = m.M
	v556 = m.ExcPending
	if v556 != 0 {
		goto L18
	} else {
		goto L121
	}
L121:
	;
	v557 = *(*int32)(unsafe.Add(mBase, uint32(v18)+140))
	*(*int32)(unsafe.Add(mBase, uint32(v18)+16)) = v557
	F_errmsg(m, int32(_a_F_GetMultiXactIdMembers_7), v18+int32(16))
	mBase = m.M
	v563 = m.ExcPending
	if v563 != 0 {
		goto L18
	} else {
		goto L122
	}
L122:
	;
	F_errfinish(m, int32(_a_F_GetMultiXactIdMembers_4), int32(1276), int32(_a_F_GetMultiXactIdMembers_5))
	mBase = m.M
	v568 = m.ExcPending
	if v568 != 0 {
		goto L18
	} else {
		goto L123
	}
L123:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L124:
	;
	F_errcode(m, int32(16779816))
	mBase = m.M
	v575 = m.ExcPending
	if v575 != 0 {
		goto L18
	} else {
		goto L125
	}
L125:
	;
	v576 = *(*int32)(unsafe.Add(mBase, uint32(v18)+140))
	*(*int32)(unsafe.Add(mBase, uint32(v18)+32)) = v576
	F_errmsg(m, int32(_a_F_GetMultiXactIdMembers_8), v18+int32(32))
	mBase = m.M
	v582 = m.ExcPending
	if v582 != 0 {
		goto L18
	} else {
		goto L126
	}
L126:
	;
	F_errfinish(m, int32(_a_F_GetMultiXactIdMembers_4), int32(1321), int32(_a_F_GetMultiXactIdMembers_5))
	mBase = m.M
	v587 = m.ExcPending
	if v587 != 0 {
		goto L18
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
	F_errcode(m, int32(16779816))
	mBase = m.M
	v594 = m.ExcPending
	if v594 != 0 {
		goto L18
	} else {
		goto L129
	}
L129:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v18)+56)) = v321
	v596 = *(*int32)(unsafe.Add(mBase, uint32(v18)+140))
	*(*int32)(unsafe.Add(mBase, uint32(v18)+48)) = v596
	F_errmsg(m, int32(_a_F_GetMultiXactIdMembers_9), v18+int32(48))
	mBase = m.M
	v602 = m.ExcPending
	if v602 != 0 {
		goto L18
	} else {
		goto L130
	}
L130:
	;
	F_errfinish(m, int32(_a_F_GetMultiXactIdMembers_4), int32(1326), int32(_a_F_GetMultiXactIdMembers_5))
	mBase = m.M
	v607 = m.ExcPending
	if v607 != 0 {
		goto L18
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
	F_errcode(m, int32(16779816))
	mBase = m.M
	v614 = m.ExcPending
	if v614 != 0 {
		goto L18
	} else {
		goto L133
	}
L133:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v18)+80)) = v370
	*(*int64)(unsafe.Add(mBase, uint32(v18)+72)) = v321
	v617 = *(*int32)(unsafe.Add(mBase, uint32(v18)+140))
	*(*int32)(unsafe.Add(mBase, uint32(v18)+64)) = v617
	F_errmsg(m, int32(_a_F_GetMultiXactIdMembers_10), v18-int32(-64))
	mBase = m.M
	v623 = m.ExcPending
	if v623 != 0 {
		goto L18
	} else {
		goto L134
	}
L134:
	;
	F_errfinish(m, int32(_a_F_GetMultiXactIdMembers_4), int32(1331), int32(_a_F_GetMultiXactIdMembers_5))
	mBase = m.M
	v628 = m.ExcPending
	if v628 != 0 {
		goto L18
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
	F_errcode(m, int32(16779816))
	mBase = m.M
	v635 = m.ExcPending
	if v635 != 0 {
		goto L18
	} else {
		goto L137
	}
L137:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v18)+104)) = v377
	v637 = *(*int32)(unsafe.Add(mBase, uint32(v18)+140))
	*(*int32)(unsafe.Add(mBase, uint32(v18)+96)) = v637
	F_errmsg(m, int32(_a_F_GetMultiXactIdMembers_11), v18+int32(96))
	mBase = m.M
	v643 = m.ExcPending
	if v643 != 0 {
		goto L18
	} else {
		goto L138
	}
L138:
	;
	F_errfinish(m, int32(_a_F_GetMultiXactIdMembers_4), int32(1336), int32(_a_F_GetMultiXactIdMembers_5))
	mBase = m.M
	v648 = m.ExcPending
	if v648 != 0 {
		goto L18
	} else {
		goto L139
	}
L139:
	;
	base.Wasm_trap_unreachable()
	for {
	}
}
func F_MultiXactIdCreateFromMembers(m *base.Module, l0 int32, l1 int32) int32 {
	mBase := m.M
	_ = mBase
	var v13 int32
	_ = v13
	var v15 int32
	_ = v15
	var v22 int32
	_ = v22
	var v24 int32
	_ = v24
	var v31 int32
	_ = v31
	var v35 int32
	_ = v35
	var v46 int32
	_ = v46
	var v49 int32
	_ = v49
	var v57 int32
	_ = v57
	var v58 int32
	_ = v58
	var v59 int32
	_ = v59
	var v62 int32
	_ = v62
	var v63 int32
	_ = v63
	var v65 int32
	_ = v65
	var v66 int32
	_ = v66
	var v68 int32
	_ = v68
	var v70 int32
	_ = v70
	var v73 int32
	_ = v73
	var v74 int32
	_ = v74
	var v75 int32
	_ = v75
	var v80 int32
	_ = v80
	var v81 int32
	_ = v81
	var v82 int32
	_ = v82
	var v85 int32
	_ = v85
	var v86 int32
	_ = v86
	var v87 int32
	_ = v87
	var v90 int32
	_ = v90
	var v91 int32
	_ = v91
	var v93 int32
	_ = v93
	var v98 int32
	_ = v98
	var v111 int32
	_ = v111
	var v113 int32
	_ = v113
	var v114 int32
	_ = v114
	var v116 int32
	_ = v116
	var v119 int32
	_ = v119
	var v122 int32
	_ = v122
	var v126 int32
	_ = v126
	var v136 int32
	_ = v136
	var v137 int32
	_ = v137
	var v157 int32
	_ = v157
	var v158 int32
	_ = v158
	var v172 int32
	_ = v172
	var v180 int32
	_ = v180
	var v196 int32
	_ = v196
	var v201 int32
	_ = v201
	var v202 int32
	_ = v202
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
	var v215 int32
	_ = v215
	var v216 int32
	_ = v216
	var v217 int32
	_ = v217
	var v221 int32
	_ = v221
	var v222 int32
	_ = v222
	var v223 int32
	_ = v223
	var v224 int32
	_ = v224
	var v226 int32
	_ = v226
	var v230 int32
	_ = v230
	var v232 int32
	_ = v232
	var v238 int32
	_ = v238
	var v239 int32
	_ = v239
	var v242 int32
	_ = v242
	var v246 int32
	_ = v246
	var v253 int32
	_ = v253
	var v255 int32
	_ = v255
	var v259 int32
	_ = v259
	var v262 int32
	_ = v262
	var v268 int32
	_ = v268
	var v272 int32
	_ = v272
	var v277 int32
	_ = v277
	var v283 int32
	_ = v283
	var v286 int32
	_ = v286
	var v290 int32
	_ = v290
	var v297 int32
	_ = v297
	var v299 int32
	_ = v299
	var v303 int32
	_ = v303
	var v304 int32
	_ = v304
	var v307 int32
	_ = v307
	var v308 int32
	_ = v308
	var v312 int32
	_ = v312
	var v319 int32
	_ = v319
	var v324 int32
	_ = v324
	var v331 int32
	_ = v331
	var v333 int32
	_ = v333
	var v334 int32
	_ = v334
	var v344 int32
	_ = v344
	var v345 int32
	_ = v345
	var v349 int32
	_ = v349
	var v353 int32
	_ = v353
	var v358 int32
	_ = v358
	var v362 int32
	_ = v362
	var v363 int32
	_ = v363
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
	var v373 int32
	_ = v373
	var v375 int32
	_ = v375
	var v381 int32
	_ = v381
	var v385 int32
	_ = v385
	var v386 int32
	_ = v386
	var v388 int32
	_ = v388
	var v390 int32
	_ = v390
	var v391 int32
	_ = v391
	var v394 int32
	_ = v394
	var v396 int32
	_ = v396
	var v397 int32
	_ = v397
	var v399 int64
	_ = v399
	var v400 int32
	_ = v400
	var v401 int32
	_ = v401
	var v405 int32
	_ = v405
	var v407 int32
	_ = v407
	var v409 int32
	_ = v409
	var v410 int32
	_ = v410
	var v413 int64
	_ = v413
	var v414 int64
	_ = v414
	var v422 int32
	_ = v422
	var v428 int64
	_ = v428
	var v435 int64
	_ = v435
	var v442 int32
	_ = v442
	var v443 int32
	_ = v443
	var v445 int64
	_ = v445
	var v447 int64
	_ = v447
	var v448 int64
	_ = v448
	var v452 int32
	_ = v452
	var v454 int32
	_ = v454
	var v455 int32
	_ = v455
	var v457 int32
	_ = v457
	var v458 int32
	_ = v458
	var v462 int32
	_ = v462
	var v464 int32
	_ = v464
	var v469 int64
	_ = v469
	var v471 int32
	_ = v471
	var v474 int32
	_ = v474
	var v478 int32
	_ = v478
	var v481 int32
	_ = v481
	var v491 int32
	_ = v491
	var v493 int32
	_ = v493
	var v499 int32
	_ = v499
	var v500 int64
	_ = v500
	var v504 int32
	_ = v504
	var v508 int32
	_ = v508
	var v513 int32
	_ = v513
	var v518 int32
	_ = v518
	var v522 int32
	_ = v522
	var v525 int64
	_ = v525
	var v526 int32
	_ = v526
	var v528 int32
	_ = v528
	var v529 int32
	_ = v529
	var v531 int32
	_ = v531
	var v536 int32
	_ = v536
	var v541 int32
	_ = v541
	var v556 int32
	_ = v556
	var v557 int32
	_ = v557
	var v559 int32
	_ = v559
	var v562 int32
	_ = v562
	var v564 int32
	_ = v564
	var v566 int32
	_ = v566
	var v568 int32
	_ = v568
	var v569 int32
	_ = v569
	var v572 int32
	_ = v572
	var v578 int32
	_ = v578
	var v585 int32
	_ = v585
	var v592 int32
	_ = v592
	var v603 int32
	_ = v603
	var v604 int32
	_ = v604
	var v607 int32
	_ = v607
	var v611 int32
	_ = v611
	var v617 int32
	_ = v617
	var v619 int32
	_ = v619
	var v637 int32
	_ = v637
	var v640 int32
	_ = v640
	var v641 int32
	_ = v641
	var v642 int32
	_ = v642
	var v643 int32
	_ = v643
	var v645 int32
	_ = v645
	var v647 int32
	_ = v647
	var v649 int32
	_ = v649
	var v656 int32
	_ = v656
	var v662 int32
	_ = v662
	var v667 int32
	_ = v667
	var v671 int32
	_ = v671
	var v677 int32
	_ = v677
	var v682 int32
	_ = v682
	var v686 int32
	_ = v686
	var v691 int32
	_ = v691
	var v695 int32
	_ = v695
	var v699 int32
	_ = v699
	var v704 int32
	_ = v704
	var v710 int32
	_ = v710
	var v714 int32
	_ = v714
	var v719 int32
	_ = v719
	var v723 int32
	_ = v723
	var v726 int32
	_ = v726
	var v730 int32
	_ = v730
	var v735 int32
	_ = v735
	v13 = m.G0
	v15 = v13 - int32(112)
	m.G0 = v15
	F_pg_qsort(m, l1, l0, int32(8), int32(309))
	mBase = m.M
	v22 = m.ExcPending
	if v22 != 0 {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	return int32(0)
L2:
	;
	v24 = *(*int32)(unsafe.Add(mBase, _c_F_MultiXactIdCreateFromMembers[0]))
	if base.B2i32(v24 == int32(0))|base.B2i32(v24 == int32(_a_F_MultiXactIdCreateFromMembers_0)) != 0 {
		goto L8
	} else {
		goto L9
	}
L3:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v723 = m.ExcPending
	if v723 != 0 {
		goto L1
	} else {
		goto L169
	}
L4:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v15)+80)) = v238
	F_errmsg(m, int32(_a_F_MultiXactIdCreateFromMembers_1), v15+int32(80))
	mBase = m.M
	v710 = m.ExcPending
	if v710 != 0 {
		goto L1
	} else {
		goto L166
	}
L5:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v695 = m.ExcPending
	if v695 != 0 {
		goto L1
	} else {
		goto L163
	}
L6:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v556 = m.ExcPending
	if v556 != 0 {
		goto L1
	} else {
		goto L132
	}
L7:
	;
	m.G0 = v15 + int32(112)
	return v541
L8:
	;
	if int32(0) < l0 {
		goto L41
	} else {
		goto L42
	}
L9:
	;
	v31 = l0 << (uint(int32(3)) % 32)
	v35 = v24
	goto L10
L10:
	;
	v46 = *(*int32)(unsafe.Add(mBase, uint32(v35-int32(4))))
	if v46 != l0 {
		goto L12
	} else {
		goto L13
	}
L11:
	;
	goto L8
L12:
	;
	v137 = *(*int32)(unsafe.Add(mBase, uint32(v35)+4))
	if v137 != int32(_a_F_MultiXactIdCreateFromMembers_0) {
		v35 = v137
		goto L10
	} else {
		goto L40
	}
L13:
	;
	v49 = v35 + int32(8)
	if base.Ui32(int32(4)) <= base.Ui32(v31) {
		goto L17
	} else {
		goto L18
	}
L14:
	;
	if v111 != 0 {
		goto L12
	} else {
		goto L32
	}
L15:
	;
	v111 = int32(0)
	goto L14
L16:
	;
	v85 = v80
	v86 = v81
	v87 = v82
	goto L26
L17:
	;
	if (l1|v49)&int32(3) != 0 {
		v80 = l1
		v81 = v49
		v82 = v31
		goto L16
	} else {
		goto L20
	}
L18:
	;
	v73 = l1
	v74 = v49
	v75 = v31
	goto L19
L19:
	;
	if v75 == int32(0) {
		goto L15
	} else {
		goto L25
	}
L20:
	;
	v57 = l1
	v58 = v49
	v59 = v31
	goto L21
L21:
	;
	v62 = *(*int32)(unsafe.Add(mBase, uint32(v57)))
	v63 = *(*int32)(unsafe.Add(mBase, uint32(v58)))
	if v62 != v63 {
		v80 = v57
		v81 = v58
		v82 = v59
		goto L16
	} else {
		goto L23
	}
L22:
	;
	v73 = v68
	v74 = v66
	v75 = v70
	goto L19
L23:
	;
	v65 = int32(4)
	v66 = v58 + v65
	v68 = v57 + v65
	v70 = v59 - v65
	if base.Ui32(int32(3)) < base.Ui32(v70) {
		v57 = v68
		v58 = v66
		v59 = v70
		goto L21
	} else {
		goto L24
	}
L24:
	;
	goto L22
L25:
	;
	v80 = v73
	v81 = v74
	v82 = v75
	goto L16
L26:
	;
	v90 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v85))))
	v91 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v86))))
	if v90 == v91 {
		goto L28
	} else {
		goto L29
	}
L27:
	;
	v111 = v90 - v91
	goto L14
L28:
	;
	v93 = int32(1)
	v98 = v87 - v93
	if v98 != 0 {
		v85 = v85 + v93
		v86 = v86 + v93
		v87 = v98
		goto L26
	} else {
		goto L31
	}
L29:
	;
	goto L30
L30:
	;
	goto L27
L31:
	;
	goto L15
L32:
	;
	if v35 != v24 {
		goto L33
	} else {
		goto L34
	}
L33:
	;
	v113 = *(*int32)(unsafe.Add(mBase, uint32(v35)))
	v114 = *(*int32)(unsafe.Add(mBase, uint32(v35)+4))
	*(*int32)(unsafe.Add(mBase, uint32(v113)+4)) = v114
	v116 = *(*int32)(unsafe.Add(mBase, uint32(v35)))
	*(*int32)(unsafe.Add(mBase, uint32(v114))) = v116
	v119 = *(*int32)(unsafe.Add(mBase, _c_F_MultiXactIdCreateFromMembers[0]))
	if v119 == int32(0) {
		goto L36
	} else {
		goto L37
	}
L34:
	;
	goto L35
L35:
	;
	v136 = *(*int32)(unsafe.Add(mBase, uint32(v35-int32(8))))
	if v136 != 0 {
		v541 = v136
		goto L7
	} else {
		goto L39
	}
L36:
	;
	v122 = int32(_a_F_MultiXactIdCreateFromMembers_0)
	*(*int32)(unsafe.Add(mBase, _c_F_MultiXactIdCreateFromMembers[1])) = v122
	v126 = v122
	goto L38
L37:
	;
	v126 = v119
	goto L38
L38:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v35))) = int32(_a_F_MultiXactIdCreateFromMembers_0)
	*(*int32)(unsafe.Add(mBase, uint32(v35)+4)) = v126
	*(*int32)(unsafe.Add(mBase, uint32(v126))) = v35
	*(*int32)(unsafe.Add(mBase, _c_F_MultiXactIdCreateFromMembers[0])) = v35
	goto L35
L39:
	;
	goto L8
L40:
	;
	goto L11
L41:
	;
	v157 = int32(0)
	v158 = int32(0)
	goto L44
L42:
	;
	goto L43
L43:
	;
	v196 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_MultiXactIdCreateFromMembers[2])))
	if v196 == int32(1) {
		goto L49
	} else {
		goto L50
	}
L44:
	;
	v172 = *(*int32)(unsafe.Add(mBase, uint32(l1+v157<<(uint(int32(3))%32))+4))
	if v158&int32(1)&base.B2i32(base.Ui32(int32(4)) <= base.Ui32(v172)) != 0 {
		goto L6
	} else {
		goto L46
	}
L45:
	;
	goto L43
L46:
	;
	v180 = v157 + int32(1)
	if v180 != l0 {
		v157 = v180
		v158 = base.B2i32(base.Ui32(int32(3)) < base.Ui32(v172)) | v158
		goto L44
	} else {
		goto L47
	}
L47:
	;
	goto L45
L48:
	;
	if v206 != 0 {
		goto L5
	} else {
		goto L52
	}
L49:
	;
	v201 = *(*int32)(unsafe.Add(mBase, _c_F_MultiXactIdCreateFromMembers[3]))
	v202 = *(*int32)(unsafe.Add(mBase, uint32(v201)+308))
	v204 = base.B2i32(v202 != int32(2))
	*(*uint8)(unsafe.Add(mBase, _c_F_MultiXactIdCreateFromMembers[2])) = uint8(v204)
	v206 = v204
	goto L51
L50:
	;
	v206 = int32(0)
	goto L51
L51:
	;
	goto L48
L52:
	;
	v208 = *(*int32)(unsafe.Add(mBase, _c_F_MultiXactIdCreateFromMembers[4]))
	v212 = F_LWLockAcquire(m, v208+int32(1664), int32(0))
	mBase = m.M
	v213 = m.ExcPending
	if v213 != 0 {
		goto L1
	} else {
		goto L53
	}
L53:
	;
	v215 = *(*int32)(unsafe.Add(mBase, _c_F_MultiXactIdCreateFromMembers[5]))
	v216 = *(*int32)(unsafe.Add(mBase, uint32(v215)))
	v217 = *(*int32)(unsafe.Add(mBase, uint32(v215)+40))
	if int32(0) <= v216-v217 {
		goto L54
	} else {
		goto L55
	}
L54:
	;
	v221 = *(*int32)(unsafe.Add(mBase, uint32(v215)+24))
	v222 = *(*int32)(unsafe.Add(mBase, uint32(v215)+52))
	v223 = *(*int32)(unsafe.Add(mBase, uint32(v215)+48))
	v224 = *(*int32)(unsafe.Add(mBase, uint32(v215)+44))
	v226 = *(*int32)(unsafe.Add(mBase, _c_F_MultiXactIdCreateFromMembers[4]))
	F_LWLockRelease(m, v226+int32(1664))
	mBase = m.M
	v230 = m.ExcPending
	if v230 != 0 {
		goto L1
	} else {
		goto L57
	}
L55:
	;
	v367 = v215
	v369 = v216
	goto L56
L56:
	;
	v372 = int32(1)
	v373 = v369 + v372
	if v373 != 0 {
		goto L98
	} else {
		goto L99
	}
L57:
	;
	v232 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_MultiXactIdCreateFromMembers[6])))
	if v232 != int32(1) {
		goto L58
	} else {
		goto L59
	}
L58:
	;
	if v216-v224 < int32(0) {
		goto L82
	} else {
		goto L83
	}
L59:
	;
	if int32(0) <= v216-v223 {
		goto L60
	} else {
		goto L61
	}
L60:
	;
	v238 = F_get_database_name(m, v221)
	mBase = m.M
	v239 = m.ExcPending
	if v239 != 0 {
		goto L1
	} else {
		goto L63
	}
L61:
	;
	goto L62
L62:
	;
	if v216&int32(_a_F_MultiXactIdCreateFromMembers_2) != 0 {
		goto L74
	} else {
		goto L75
	}
L63:
	;
	v242 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_MultiXactIdCreateFromMembers[6])))
	if v242 == int32(1) {
		goto L65
	} else {
		goto L66
	}
L64:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v259 = m.ExcPending
	if v259 != 0 {
		goto L1
	} else {
		goto L68
	}
L65:
	;
	v246 = *(*int32)(unsafe.Add(mBase, _c_F_MultiXactIdCreateFromMembers[7]))
	*(*int32)(unsafe.Add(mBase, uint32(v246+int32(16)))) = int32(1)
	v253 = *(*int32)(unsafe.Add(mBase, _c_F_MultiXactIdCreateFromMembers[8]))
	v255 = F_pgmem_kill(m, v253, int32(10))
	mBase = m.M
	goto L67
L66:
	;
	goto L67
L67:
	;
	goto L64
L68:
	;
	F_errcode(m, int32(261))
	mBase = m.M
	v262 = m.ExcPending
	if v262 != 0 {
		goto L1
	} else {
		goto L69
	}
L69:
	;
	if v238 != 0 {
		goto L4
	} else {
		goto L70
	}
L70:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v15)+64)) = v221
	F_errmsg(m, int32(_a_F_MultiXactIdCreateFromMembers_3), v15-int32(-64))
	mBase = m.M
	v268 = m.ExcPending
	if v268 != 0 {
		goto L1
	} else {
		goto L71
	}
L71:
	;
	F_errhint(m, int32(_a_F_MultiXactIdCreateFromMembers_4), int32(0))
	mBase = m.M
	v272 = m.ExcPending
	if v272 != 0 {
		goto L1
	} else {
		goto L72
	}
L72:
	;
	F_errfinish(m, int32(_a_F_MultiXactIdCreateFromMembers_5), int32(1050), int32(_a_F_MultiXactIdCreateFromMembers_6))
	mBase = m.M
	v277 = m.ExcPending
	if v277 != 0 {
		goto L1
	} else {
		goto L73
	}
L73:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L74:
	;
	v283 = base.B2i32(v216 != int32(1))
	goto L76
L75:
	;
	v283 = int32(0)
	goto L76
L76:
	;
	if v283 != 0 {
		goto L58
	} else {
		goto L77
	}
L77:
	;
	v286 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_MultiXactIdCreateFromMembers[6])))
	if v286 == int32(1) {
		goto L79
	} else {
		goto L80
	}
L78:
	;
	goto L58
L79:
	;
	v290 = *(*int32)(unsafe.Add(mBase, _c_F_MultiXactIdCreateFromMembers[7]))
	*(*int32)(unsafe.Add(mBase, uint32(v290+int32(16)))) = int32(1)
	v297 = *(*int32)(unsafe.Add(mBase, _c_F_MultiXactIdCreateFromMembers[8]))
	v299 = F_pgmem_kill(m, v297, int32(10))
	mBase = m.M
	goto L81
L80:
	;
	goto L81
L81:
	;
	goto L78
L82:
	;
	v358 = *(*int32)(unsafe.Add(mBase, _c_F_MultiXactIdCreateFromMembers[4]))
	v362 = F_LWLockAcquire(m, v358+int32(1664), int32(0))
	mBase = m.M
	v363 = m.ExcPending
	if v363 != 0 {
		goto L1
	} else {
		goto L97
	}
L83:
	;
	v303 = F_get_database_name(m, v221)
	mBase = m.M
	v304 = m.ExcPending
	if v304 != 0 {
		goto L1
	} else {
		goto L84
	}
L84:
	;
	v307 = F_errstart(m, int32(19), int32(0))
	mBase = m.M
	v308 = m.ExcPending
	if v308 != 0 {
		goto L1
	} else {
		goto L85
	}
L85:
	;
	if v303 != 0 {
		goto L87
	} else {
		goto L88
	}
L86:
	;
	*(*float64)(unsafe.Add(mBase, uint32(v15)+16)) = base.F64_mul(base.F64_div(base.F64_convert_i32_u(v333), float64(2.147483647e+09)), float64(100))
	v344 = F_errdetail(m, int32(_a_F_MultiXactIdCreateFromMembers_7), v15+int32(16))
	mBase = m.M
	v345 = m.ExcPending
	if v345 != 0 {
		goto L1
	} else {
		goto L94
	}
L87:
	;
	if v307 == int32(0) {
		goto L82
	} else {
		goto L90
	}
L88:
	;
	goto L89
L89:
	;
	if v307 == int32(0) {
		goto L82
	} else {
		goto L92
	}
L90:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v15)+48)) = v303
	v312 = v222 - v216
	*(*int32)(unsafe.Add(mBase, uint32(v15)+52)) = v312
	F_errmsg_plural(m, int32(_a_F_MultiXactIdCreateFromMembers_8), int32(_a_F_MultiXactIdCreateFromMembers_9), v312, v15+int32(48))
	mBase = m.M
	v319 = m.ExcPending
	if v319 != 0 {
		goto L1
	} else {
		goto L91
	}
L91:
	;
	v333 = v312
	v334 = int32(1076)
	goto L86
L92:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v15)+32)) = v221
	v324 = v222 - v216
	*(*int32)(unsafe.Add(mBase, uint32(v15)+36)) = v324
	F_errmsg_plural(m, int32(_a_F_MultiXactIdCreateFromMembers_10), int32(_a_F_MultiXactIdCreateFromMembers_11), v324, v15+int32(32))
	mBase = m.M
	v331 = m.ExcPending
	if v331 != 0 {
		goto L1
	} else {
		goto L93
	}
L93:
	;
	v333 = v324
	v334 = int32(1087)
	goto L86
L94:
	;
	F_errhint(m, int32(_a_F_MultiXactIdCreateFromMembers_4), int32(0))
	mBase = m.M
	v349 = m.ExcPending
	if v349 != 0 {
		goto L1
	} else {
		goto L95
	}
L95:
	;
	F_errfinish(m, int32(_a_F_MultiXactIdCreateFromMembers_5), v334, int32(_a_F_MultiXactIdCreateFromMembers_6))
	mBase = m.M
	v353 = m.ExcPending
	if v353 != 0 {
		goto L1
	} else {
		goto L96
	}
L96:
	;
	goto L82
L97:
	;
	v365 = *(*int32)(unsafe.Add(mBase, _c_F_MultiXactIdCreateFromMembers[5]))
	v366 = *(*int32)(unsafe.Add(mBase, uint32(v365)))
	v367 = v365
	v369 = v366
	goto L56
L98:
	;
	v375 = v373
	goto L100
L99:
	;
	v375 = v372
	goto L100
L100:
	;
	if v375&int32(1023) != 0 {
		goto L101
	} else {
		goto L102
	}
L101:
	;
	v381 = base.B2i32(v375 != int32(1))
	goto L103
L102:
	;
	v381 = int32(0)
	goto L103
L103:
	;
	if v381 == int32(0) {
		goto L104
	} else {
		goto L105
	}
L104:
	;
	v385 = *(*int32)(unsafe.Add(mBase, _c_F_MultiXactIdCreateFromMembers[9]))
	v386 = *(*int32)(unsafe.Add(mBase, uint32(v385)+28))
	v388 = int32(base.Ui32(v375) >> (uint(int32(10)) % 32))
	v390 = int32(*(*uint16)(unsafe.Add(mBase, _c_F_MultiXactIdCreateFromMembers[10])))
	v391 = base.I32_rem_u_s(v388, v390)
	v394 = v386 + v391<<(uint(int32(7))%32)
	v396 = F_LWLockAcquire(m, v394, int32(0))
	mBase = m.M
	v397 = m.ExcPending
	if v397 != 0 {
		goto L1
	} else {
		goto L107
	}
L105:
	;
	v410 = v367
	goto L106
L106:
	;
	v413 = base.I64_extend_i32_s(l0)
	v414 = *(*int64)(unsafe.Add(mBase, uint32(v410)+8))
	if base.Ui64(v414^int64(-1)) < base.Ui64(v413) {
		goto L3
	} else {
		goto L111
	}
L107:
	;
	v399 = base.I64_extend_i32_u(v388)
	v400 = F_SimpleLruZeroPage(m, int32(_a_F_MultiXactIdCreateFromMembers_12), v399)
	mBase = m.M
	v401 = m.ExcPending
	if v401 != 0 {
		goto L1
	} else {
		goto L108
	}
L108:
	;
	F_XLogSimpleInsertInt64(m, int32(6), int32(0), v399)
	mBase = m.M
	v405 = m.ExcPending
	if v405 != 0 {
		goto L1
	} else {
		goto L109
	}
L109:
	;
	F_LWLockRelease(m, v394)
	mBase = m.M
	v407 = m.ExcPending
	if v407 != 0 {
		goto L1
	} else {
		goto L110
	}
L110:
	;
	v409 = *(*int32)(unsafe.Add(mBase, _c_F_MultiXactIdCreateFromMembers[5]))
	v410 = v409
	goto L106
L111:
	;
	if int32(0) < l0 {
		goto L112
	} else {
		goto L113
	}
L112:
	;
	v422 = l0
	v428 = v414
	goto L115
L113:
	;
	v481 = v410
	goto L114
L114:
	;
	v491 = int32(_a_F_MultiXactIdCreateFromMembers_13)
	v493 = *(*int32)(unsafe.Add(mBase, _c_F_MultiXactIdCreateFromMembers[11]))
	*(*int32)(unsafe.Add(mBase, _c_F_MultiXactIdCreateFromMembers[11])) = v493 + int32(1)
	*(*int32)(unsafe.Add(mBase, uint32(v481))) = v375
	v499 = *(*int32)(unsafe.Add(mBase, _c_F_MultiXactIdCreateFromMembers[5]))
	v500 = *(*int64)(unsafe.Add(mBase, uint32(v499)+8))
	*(*int64)(unsafe.Add(mBase, uint32(v499)+8)) = v500 + v413
	v504 = *(*int32)(unsafe.Add(mBase, _c_F_MultiXactIdCreateFromMembers[4]))
	F_LWLockRelease(m, v504+int32(1664))
	mBase = m.M
	v508 = m.ExcPending
	if v508 != 0 {
		goto L1
	} else {
		goto L125
	}
L115:
	;
	v435 = base.I64_rem_u_s(int64(base.Ui64(v428)>>(uint(int64(2))%64)), int64(409))
	if v435|v428&int64(3) == int64(0) {
		goto L117
	} else {
		goto L118
	}
L116:
	;
	v478 = *(*int32)(unsafe.Add(mBase, _c_F_MultiXactIdCreateFromMembers[5]))
	v481 = v478
	goto L114
L117:
	;
	v442 = *(*int32)(unsafe.Add(mBase, _c_F_MultiXactIdCreateFromMembers[12]))
	v443 = *(*int32)(unsafe.Add(mBase, uint32(v442)+28))
	v445 = base.I64_div_u_s(v428, int64(1636))
	v447 = int64(*(*uint16)(unsafe.Add(mBase, _c_F_MultiXactIdCreateFromMembers[13])))
	v448 = base.I64_rem_u_s(v445, v447)
	v452 = v443 + base.I32_wrap_i64(v448)<<(uint(int32(7))%32)
	v454 = F_LWLockAcquire(m, v452, int32(0))
	mBase = m.M
	v455 = m.ExcPending
	if v455 != 0 {
		goto L1
	} else {
		goto L120
	}
L118:
	;
	goto L119
L119:
	;
	v469 = base.I64_rem_u_s(v428, int64(1636))
	v471 = int32(1636) - base.I32_wrap_i64(v469)
	v474 = v422 - v471
	if int32(0) < v474 {
		v422 = v474
		v428 = v428 + base.I64_extend_i32_u(v471)
		goto L115
	} else {
		goto L124
	}
L120:
	;
	v457 = F_SimpleLruZeroPage(m, int32(_a_F_MultiXactIdCreateFromMembers_14), v445)
	mBase = m.M
	v458 = m.ExcPending
	if v458 != 0 {
		goto L1
	} else {
		goto L121
	}
L121:
	;
	F_XLogSimpleInsertInt64(m, int32(6), int32(16), v445)
	mBase = m.M
	v462 = m.ExcPending
	if v462 != 0 {
		goto L1
	} else {
		goto L122
	}
L122:
	;
	F_LWLockRelease(m, v452)
	mBase = m.M
	v464 = m.ExcPending
	if v464 != 0 {
		goto L1
	} else {
		goto L123
	}
L123:
	;
	goto L119
L124:
	;
	goto L116
L125:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v15)+104)) = l0
	*(*int64)(unsafe.Add(mBase, uint32(v15)+96)) = v414
	*(*int32)(unsafe.Add(mBase, uint32(v15)+88)) = v369
	F_XLogBeginInsert(m)
	mBase = m.M
	v513 = m.ExcPending
	if v513 != 0 {
		goto L1
	} else {
		goto L126
	}
L126:
	;
	F_XLogRegisterData(m, v15+int32(88), int32(20))
	mBase = m.M
	v518 = m.ExcPending
	if v518 != 0 {
		goto L1
	} else {
		goto L127
	}
L127:
	;
	F_XLogRegisterData(m, l1, l0<<(uint(int32(3))%32))
	mBase = m.M
	v522 = m.ExcPending
	if v522 != 0 {
		goto L1
	} else {
		goto L128
	}
L128:
	;
	v525 = F_XLogInsert(m, int32(6), int32(32))
	mBase = m.M
	v526 = m.ExcPending
	if v526 != 0 {
		goto L1
	} else {
		goto L129
	}
L129:
	;
	F_RecordNewMultiXact(m, v369, v414, l0, l1)
	mBase = m.M
	v528 = m.ExcPending
	if v528 != 0 {
		goto L1
	} else {
		goto L130
	}
L130:
	;
	v529 = int32(_a_F_MultiXactIdCreateFromMembers_13)
	v531 = *(*int32)(unsafe.Add(mBase, _c_F_MultiXactIdCreateFromMembers[11]))
	*(*int32)(unsafe.Add(mBase, _c_F_MultiXactIdCreateFromMembers[11])) = v531 - int32(1)
	F_mXactCachePut(m, v369, l0, l1)
	mBase = m.M
	v536 = m.ExcPending
	if v536 != 0 {
		goto L1
	} else {
		goto L131
	}
L131:
	;
	v541 = v369
	goto L7
L132:
	;
	v557 = m.G0
	v559 = v557 - int32(80)
	m.G0 = v559
	v562 = *(*int32)(unsafe.Add(mBase, _c_F_MultiXactIdCreateFromMembers[14]))
	if v562 != 0 {
		goto L134
	} else {
		goto L135
	}
L133:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v15))) = v649
	F_errmsg_internal(m, int32(_a_F_MultiXactIdCreateFromMembers_15), v15)
	mBase = m.M
	v686 = m.ExcPending
	if v686 != 0 {
		goto L1
	} else {
		goto L161
	}
L134:
	;
	F_pfree(m, v562)
	mBase = m.M
	v564 = m.ExcPending
	if v564 != 0 {
		goto L1
	} else {
		goto L137
	}
L135:
	;
	goto L136
L136:
	;
	v566 = v559 - int32(-64)
	F_initStringInfo(m, v566)
	mBase = m.M
	v568 = m.ExcPending
	if v568 != 0 {
		goto L1
	} else {
		goto L138
	}
L137:
	;
	goto L136
L138:
	;
	v569 = *(*int32)(unsafe.Add(mBase, uint32(l1)+4))
	if base.Ui32(v569) < base.Ui32(int32(6)) {
		goto L140
	} else {
		goto L141
	}
L139:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v671 = m.ExcPending
	if v671 != 0 {
		goto L1
	} else {
		goto L158
	}
L140:
	;
	v572 = *(*int32)(unsafe.Add(mBase, uint32(l1)))
	*(*int32)(unsafe.Add(mBase, uint32(v559)+40)) = v572
	*(*int32)(unsafe.Add(mBase, uint32(v559)+32)) = int32(0)
	v578 = *(*int32)(unsafe.Add(mBase, uint32(v569<<(uint(int32(2))%32))+uint32(_c_F_MultiXactIdCreateFromMembers[15])))
	*(*int32)(unsafe.Add(mBase, uint32(v559)+44)) = v578
	*(*int32)(unsafe.Add(mBase, uint32(v559)+36)) = l0
	F_appendStringInfo(m, v566, int32(_a_F_MultiXactIdCreateFromMembers_16), v559+int32(32))
	mBase = m.M
	v585 = m.ExcPending
	if v585 != 0 {
		goto L1
	} else {
		goto L143
	}
L141:
	;
	goto L142
L142:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v656 = m.ExcPending
	if v656 != 0 {
		goto L1
	} else {
		goto L155
	}
L143:
	;
	if int32(2) <= l0 {
		goto L144
	} else {
		goto L145
	}
L144:
	;
	v592 = int32(1)
	goto L147
L145:
	;
	goto L146
L146:
	;
	F_appendStringInfoChar(m, v559-int32(-64), int32(93))
	mBase = m.M
	v637 = m.ExcPending
	if v637 != 0 {
		goto L1
	} else {
		goto L152
	}
L147:
	;
	v603 = l1 + v592<<(uint(int32(3))%32)
	v604 = *(*int32)(unsafe.Add(mBase, uint32(v603)+4))
	if base.Ui32(int32(6)) <= base.Ui32(v604) {
		goto L139
	} else {
		goto L149
	}
L148:
	;
	goto L146
L149:
	;
	v607 = *(*int32)(unsafe.Add(mBase, uint32(v603)))
	*(*int32)(unsafe.Add(mBase, uint32(v559))) = v607
	v611 = *(*int32)(unsafe.Add(mBase, uint32(v604<<(uint(int32(2))%32))+uint32(_c_F_MultiXactIdCreateFromMembers[15])))
	*(*int32)(unsafe.Add(mBase, uint32(v559)+4)) = v611
	F_appendStringInfo(m, v559-int32(-64), int32(_a_F_MultiXactIdCreateFromMembers_17), v559)
	mBase = m.M
	v617 = m.ExcPending
	if v617 != 0 {
		goto L1
	} else {
		goto L150
	}
L150:
	;
	v619 = v592 + int32(1)
	if v619 != l0 {
		v592 = v619
		goto L147
	} else {
		goto L151
	}
L151:
	;
	goto L148
L152:
	;
	v640 = *(*int32)(unsafe.Add(mBase, _c_F_MultiXactIdCreateFromMembers[16]))
	v641 = *(*int32)(unsafe.Add(mBase, uint32(v559)+64))
	v642 = F_MemoryContextStrdup(m, v640, v641)
	mBase = m.M
	v643 = m.ExcPending
	if v643 != 0 {
		goto L1
	} else {
		goto L153
	}
L153:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_MultiXactIdCreateFromMembers[14])) = v642
	v645 = *(*int32)(unsafe.Add(mBase, uint32(v559)+64))
	F_pfree(m, v645)
	mBase = m.M
	v647 = m.ExcPending
	if v647 != 0 {
		goto L1
	} else {
		goto L154
	}
L154:
	;
	v649 = *(*int32)(unsafe.Add(mBase, _c_F_MultiXactIdCreateFromMembers[14]))
	m.G0 = v559 + int32(80)
	goto L133
L155:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v559)+48)) = v569
	F_errmsg_internal(m, int32(_a_F_MultiXactIdCreateFromMembers_18), v559+int32(48))
	mBase = m.M
	v662 = m.ExcPending
	if v662 != 0 {
		goto L1
	} else {
		goto L156
	}
L156:
	;
	F_errfinish(m, int32(_a_F_MultiXactIdCreateFromMembers_5), int32(1591), int32(_a_F_MultiXactIdCreateFromMembers_19))
	mBase = m.M
	v667 = m.ExcPending
	if v667 != 0 {
		goto L1
	} else {
		goto L157
	}
L157:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L158:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v559)+16)) = v604
	F_errmsg_internal(m, int32(_a_F_MultiXactIdCreateFromMembers_18), v559+int32(16))
	mBase = m.M
	v677 = m.ExcPending
	if v677 != 0 {
		goto L1
	} else {
		goto L159
	}
L159:
	;
	F_errfinish(m, int32(_a_F_MultiXactIdCreateFromMembers_5), int32(1591), int32(_a_F_MultiXactIdCreateFromMembers_19))
	mBase = m.M
	v682 = m.ExcPending
	if v682 != 0 {
		goto L1
	} else {
		goto L160
	}
L160:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L161:
	;
	F_errfinish(m, int32(_a_F_MultiXactIdCreateFromMembers_5), int32(752), int32(_a_F_MultiXactIdCreateFromMembers_20))
	mBase = m.M
	v691 = m.ExcPending
	if v691 != 0 {
		goto L1
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
	F_errmsg_internal(m, int32(_a_F_MultiXactIdCreateFromMembers_21), int32(0))
	mBase = m.M
	v699 = m.ExcPending
	if v699 != 0 {
		goto L1
	} else {
		goto L164
	}
L164:
	;
	F_errfinish(m, int32(_a_F_MultiXactIdCreateFromMembers_5), int32(988), int32(_a_F_MultiXactIdCreateFromMembers_6))
	mBase = m.M
	v704 = m.ExcPending
	if v704 != 0 {
		goto L1
	} else {
		goto L165
	}
L165:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L166:
	;
	F_errhint(m, int32(_a_F_MultiXactIdCreateFromMembers_4), int32(0))
	mBase = m.M
	v714 = m.ExcPending
	if v714 != 0 {
		goto L1
	} else {
		goto L167
	}
L167:
	;
	F_errfinish(m, int32(_a_F_MultiXactIdCreateFromMembers_5), int32(1043), int32(_a_F_MultiXactIdCreateFromMembers_6))
	mBase = m.M
	v719 = m.ExcPending
	if v719 != 0 {
		goto L1
	} else {
		goto L168
	}
L168:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L169:
	;
	F_errcode(m, int32(261))
	mBase = m.M
	v726 = m.ExcPending
	if v726 != 0 {
		goto L1
	} else {
		goto L170
	}
L170:
	;
	F_errmsg(m, int32(_a_F_MultiXactIdCreateFromMembers_22), int32(0))
	mBase = m.M
	v730 = m.ExcPending
	if v730 != 0 {
		goto L1
	} else {
		goto L171
	}
L171:
	;
	F_errfinish(m, int32(_a_F_MultiXactIdCreateFromMembers_5), int32(1116), int32(_a_F_MultiXactIdCreateFromMembers_6))
	mBase = m.M
	v735 = m.ExcPending
	if v735 != 0 {
		goto L1
	} else {
		goto L172
	}
L172:
	;
	base.Wasm_trap_unreachable()
	for {
	}
}
func F_ReadNextMultiXactId(m *base.Module) int32 {
	mBase := m.M
	_ = mBase
	var v3 int32
	_ = v3
	var v7 int32
	_ = v7
	var v10 int32
	_ = v10
	var v12 int32
	_ = v12
	var v13 int32
	_ = v13
	var v15 int32
	_ = v15
	var v19 int32
	_ = v19
	v3 = *(*int32)(unsafe.Add(mBase, _c_F_ReadNextMultiXactId[0]))
	v7 = F_LWLockAcquire(m, v3+int32(1664), int32(1))
	mBase = m.M
	v10 = m.ExcPending
	if v10 != 0 {
		return int32(0)
	} else {
		v12 = *(*int32)(unsafe.Add(mBase, _c_F_ReadNextMultiXactId[1]))
		v13 = *(*int32)(unsafe.Add(mBase, uint32(v12)))
		v15 = *(*int32)(unsafe.Add(mBase, _c_F_ReadNextMultiXactId[0]))
		F_LWLockRelease(m, v15+int32(1664))
		mBase = m.M
		v19 = m.ExcPending
		if v19 != 0 {
			return int32(0)
		} else {
			return v13
		}
	}
}
