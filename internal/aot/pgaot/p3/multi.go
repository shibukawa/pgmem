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
	var v194 int32
	_ = v194
	var v200 int32
	_ = v200
	var v225 int32
	_ = v225
	var v226 int32
	_ = v226
	var v236 int32
	_ = v236
	var v239 int32
	_ = v239
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
	var v260 int32
	_ = v260
	var v261 int32
	_ = v261
	var v262 int32
	_ = v262
	var v264 int32
	_ = v264
	var v265 int32
	_ = v265
	var v266 int32
	_ = v266
	var v267 int32
	_ = v267
	var v269 int32
	_ = v269
	var v270 int32
	_ = v270
	var v271 int32
	_ = v271
	var v272 int32
	_ = v272
	var v274 int32
	_ = v274
	var v275 int32
	_ = v275
	var v276 int32
	_ = v276
	var v278 int32
	_ = v278
	var v285 int32
	_ = v285
	var v304 int32
	_ = v304
	var v306 int32
	_ = v306
	var v323 int32
	_ = v323
	var v324 int32
	_ = v324
	var v325 int32
	_ = v325
	var v327 int32
	_ = v327
	var v328 int32
	_ = v328
	var v331 int32
	_ = v331
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
	var v359 int64
	_ = v359
	var v360 int32
	_ = v360
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
	var v371 int32
	_ = v371
	var v372 int32
	_ = v372
	var v373 int32
	_ = v373
	var v374 int32
	_ = v374
	var v376 int32
	_ = v376
	var v382 int32
	_ = v382
	var v387 int32
	_ = v387
	var v403 int32
	_ = v403
	var v409 int64
	_ = v409
	var v413 int32
	_ = v413
	var v414 int32
	_ = v414
	var v415 int32
	_ = v415
	var v420 int32
	_ = v420
	var v421 int32
	_ = v421
	var v422 int32
	_ = v422
	var v423 int32
	_ = v423
	var v425 int32
	_ = v425
	var v427 int32
	_ = v427
	var v428 int32
	_ = v428
	var v431 int32
	_ = v431
	var v434 int32
	_ = v434
	var v440 int64
	_ = v440
	var v445 int32
	_ = v445
	var v447 int32
	_ = v447
	var v449 int32
	_ = v449
	var v455 int32
	_ = v455
	var v456 int32
	_ = v456
	var v457 int32
	_ = v457
	var v459 int32
	_ = v459
	var v461 int32
	_ = v461
	var v482 int64
	_ = v482
	var v484 int64
	_ = v484
	var v489 int32
	_ = v489
	var v493 int32
	_ = v493
	var v498 int32
	_ = v498
	var v500 int32
	_ = v500
	var v501 int32
	_ = v501
	var v504 int32
	_ = v504
	var v512 int32
	_ = v512
	var v518 int32
	_ = v518
	var v546 int32
	_ = v546
	var v547 int32
	_ = v547
	var v570 int32
	_ = v570
	var v576 int32
	_ = v576
	var v592 int32
	_ = v592
	var v595 int32
	_ = v595
	var v596 int32
	_ = v596
	var v597 int32
	_ = v597
	var v599 int32
	_ = v599
	var v600 int32
	_ = v600
	var v602 int32
	_ = v602
	var v603 int32
	_ = v603
	var v605 int32
	_ = v605
	var v606 int32
	_ = v606
	var v607 int32
	_ = v607
	var v609 int32
	_ = v609
	var v610 int32
	_ = v610
	var v611 int32
	_ = v611
	var v613 int32
	_ = v613
	var v616 int32
	_ = v616
	var v618 int32
	_ = v618
	var v623 int32
	_ = v623
	var v642 int32
	_ = v642
	var v644 int32
	_ = v644
	var v646 int32
	_ = v646
	var v650 int32
	_ = v650
	var v651 int32
	_ = v651
	var v652 int32
	_ = v652
	var v655 int32
	_ = v655
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
	var v667 int32
	_ = v667
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
	v570 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	if v570 == int32(0) {
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
	v546 = v48 + int32(1)
	v547 = *(*int32)(unsafe.Add(mBase, uint32(v24)+4))
	if v546 < v547 {
		v48 = v546
		goto L8
	} else {
		goto L82
	}
L11:
	;
	v61 = *(*int32)(unsafe.Add(mBase, uint32(v59)+104))
	v62 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(v58)+200)) = uint8(v62)
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
	v354 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v58)+304)))
	v355 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
	v356 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v357 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v58)+304)) = uint8(v357)
	v359 = *(*int64)(unsafe.Add(mBase, uint32(v58)+184))
	v360 = *(*int32)(unsafe.Add(mBase, uint32(v57)+152))
	if v360 == v357 {
		goto L53
	} else {
		goto L54
	}
L14:
	;
	v352 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v58)+200)) = uint8(v352)
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
	v225 = v56 & int32(3)
	v226 = int32(0)
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
	*(*int32)(unsafe.Add(mBase, uint32(v135)+36)) = v111
	v138 = *(*int32)(unsafe.Add(mBase, uint32(v58)+260))
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
	*(*int64)(unsafe.Add(mBase, uint32(v171+int32(16))+232)) = v166
	v194 = *(*int32)(unsafe.Add(mBase, uint32(v171)))
	*(*int32)(unsafe.Add(mBase, uint32(v171))) = v194 + v183
	v200 = *(*int32)(unsafe.Add(mBase, _c_F_CopyMultiInsertInfoFlush[2]))
	*(*int32)(unsafe.Add(mBase, _c_F_CopyMultiInsertInfoFlush[2])) = v200 - v183
	goto L34
L37:
	;
	goto L17
L38:
	;
	v236 = v226
	v239 = int32(0)
	goto L41
L39:
	;
	v285 = v226
	goto L40
L40:
	;
	v304 = v285
	v306 = v226
	goto L49
L41:
	;
	v254 = v55 + v236<<(uint(int32(2))%32)
	v255 = *(*int32)(unsafe.Add(mBase, uint32(v254)))
	v256 = *(*int32)(unsafe.Add(mBase, uint32(v255)+8))
	v257 = *(*int32)(unsafe.Add(mBase, uint32(v256)+12))
	m.T0[v257].(func(*base.Module, int32))(m, v255)
	mBase = m.M
	v259 = m.ExcPending
	if v259 != 0 {
		goto L21
	} else {
		goto L43
	}
L42:
	;
	if v225 == int32(0) {
		goto L14
	} else {
		goto L48
	}
L43:
	;
	v260 = *(*int32)(unsafe.Add(mBase, uint32(v254)+4))
	v261 = *(*int32)(unsafe.Add(mBase, uint32(v260)+8))
	v262 = *(*int32)(unsafe.Add(mBase, uint32(v261)+12))
	m.T0[v262].(func(*base.Module, int32))(m, v260)
	mBase = m.M
	v264 = m.ExcPending
	if v264 != 0 {
		goto L21
	} else {
		goto L44
	}
L44:
	;
	v265 = *(*int32)(unsafe.Add(mBase, uint32(v254)+8))
	v266 = *(*int32)(unsafe.Add(mBase, uint32(v265)+8))
	v267 = *(*int32)(unsafe.Add(mBase, uint32(v266)+12))
	m.T0[v267].(func(*base.Module, int32))(m, v265)
	mBase = m.M
	v269 = m.ExcPending
	if v269 != 0 {
		goto L21
	} else {
		goto L45
	}
L45:
	;
	v270 = *(*int32)(unsafe.Add(mBase, uint32(v254)+12))
	v271 = *(*int32)(unsafe.Add(mBase, uint32(v270)+8))
	v272 = *(*int32)(unsafe.Add(mBase, uint32(v271)+12))
	m.T0[v272].(func(*base.Module, int32))(m, v270)
	mBase = m.M
	v274 = m.ExcPending
	if v274 != 0 {
		goto L21
	} else {
		goto L46
	}
L46:
	;
	v275 = int32(4)
	v276 = v236 + v275
	v278 = v239 + v275
	if v278 != v56&int32(-4) {
		v236 = v276
		v239 = v278
		goto L41
	} else {
		goto L47
	}
L47:
	;
	goto L42
L48:
	;
	v285 = v276
	goto L40
L49:
	;
	v323 = *(*int32)(unsafe.Add(mBase, uint32(v55+v304<<(uint(int32(2))%32))))
	v324 = *(*int32)(unsafe.Add(mBase, uint32(v323)+8))
	v325 = *(*int32)(unsafe.Add(mBase, uint32(v324)+12))
	m.T0[v325].(func(*base.Module, int32))(m, v323)
	mBase = m.M
	v327 = m.ExcPending
	if v327 != 0 {
		goto L21
	} else {
		goto L51
	}
L50:
	;
	goto L14
L51:
	;
	v328 = int32(1)
	v331 = v306 + v328
	if v331 != v225 {
		v304 = v304 + v328
		v306 = v331
		goto L49
	} else {
		goto L52
	}
L52:
	;
	goto L50
L53:
	;
	v363 = F_MakePerTupleExprContext(m, v57)
	mBase = m.M
	v364 = m.ExcPending
	if v364 != 0 {
		goto L21
	} else {
		goto L56
	}
L54:
	;
	v365 = v360
	goto L55
L55:
	;
	v366 = int32(_a_F_CopyMultiInsertInfoFlush_1)
	v367 = *(*int32)(unsafe.Add(mBase, _c_F_CopyMultiInsertInfoFlush[3]))
	v369 = *(*int32)(unsafe.Add(mBase, uint32(v365)+20))
	*(*int32)(unsafe.Add(mBase, _c_F_CopyMultiInsertInfoFlush[3])) = v369
	v371 = *(*int32)(unsafe.Add(mBase, uint32(v59)+8))
	v372 = *(*int32)(unsafe.Add(mBase, uint32(v55)+4004))
	v373 = *(*int32)(unsafe.Add(mBase, uint32(v371)+188))
	v374 = *(*int32)(unsafe.Add(mBase, uint32(v373)+92))
	m.T0[v374].(func(*base.Module, int32, int32, int32, int32, int32, int32))(m, v371, v55, v56, v356, v355, v372)
	mBase = m.M
	v376 = m.ExcPending
	if v376 != 0 {
		goto L21
	} else {
		goto L57
	}
L56:
	;
	v365 = v363
	goto L55
L57:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_CopyMultiInsertInfoFlush[3])) = v367
	if int32(0) < v56 {
		goto L58
	} else {
		goto L59
	}
L58:
	;
	v382 = v55 + int32(4016)
	v387 = int32(0)
	goto L61
L59:
	;
	goto L60
L60:
	;
	v482 = *(*int64)(unsafe.Add(mBase, uint32(l2)))
	v484 = v482 + base.I64_extend_i32_s(v56)
	*(*int64)(unsafe.Add(mBase, uint32(l2))) = v484
	v489 = *(*int32)(unsafe.Add(mBase, _c_F_CopyMultiInsertInfoFlush[0]))
	if v489 == int32(0) {
		goto L79
	} else {
		goto L80
	}
L61:
	;
	v403 = *(*int32)(unsafe.Add(mBase, uint32(v59)+12))
	if int32(0) < v403 {
		goto L64
	} else {
		goto L65
	}
L62:
	;
	goto L60
L63:
	;
	v455 = *(*int32)(unsafe.Add(mBase, uint32(v55+v387<<(uint(int32(2))%32))))
	v456 = *(*int32)(unsafe.Add(mBase, uint32(v455)+8))
	v457 = *(*int32)(unsafe.Add(mBase, uint32(v456)+12))
	m.T0[v457].(func(*base.Module, int32))(m, v455)
	mBase = m.M
	v459 = m.ExcPending
	if v459 != 0 {
		goto L21
	} else {
		goto L76
	}
L64:
	;
	v409 = *(*int64)(unsafe.Add(mBase, uint32(v382+v387<<(uint(int32(3))%32))))
	*(*int64)(unsafe.Add(mBase, uint32(v58)+184)) = v409
	v413 = v55 + v387<<(uint(int32(2))%32)
	v414 = *(*int32)(unsafe.Add(mBase, uint32(v413)))
	v415 = int32(0)
	v420 = F_ExecInsertIndexTuples(m, v59, v414, v57, v415, v415, v415, v415, v415)
	mBase = m.M
	v421 = m.ExcPending
	if v421 != 0 {
		goto L21
	} else {
		goto L67
	}
L65:
	;
	goto L66
L66:
	;
	v428 = *(*int32)(unsafe.Add(mBase, uint32(v59)+52))
	if v428 == int32(0) {
		goto L63
	} else {
		goto L70
	}
L67:
	;
	v422 = *(*int32)(unsafe.Add(mBase, uint32(v413)))
	v423 = *(*int32)(unsafe.Add(mBase, uint32(v58)+260))
	F_ExecARInsertTriggers(m, v57, v59, v422, v420, v423)
	mBase = m.M
	v425 = m.ExcPending
	if v425 != 0 {
		goto L21
	} else {
		goto L68
	}
L68:
	;
	F_list_free(m, v420)
	mBase = m.M
	v427 = m.ExcPending
	if v427 != 0 {
		goto L21
	} else {
		goto L69
	}
L69:
	;
	goto L63
L70:
	;
	v431 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v428)+9)))
	if v431 == int32(0) {
		goto L71
	} else {
		goto L72
	}
L71:
	;
	v434 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v428)+25)))
	if v434 != int32(1) {
		goto L63
	} else {
		goto L74
	}
L72:
	;
	goto L73
L73:
	;
	v440 = *(*int64)(unsafe.Add(mBase, uint32(v382+v387<<(uint(int32(3))%32))))
	*(*int64)(unsafe.Add(mBase, uint32(v58)+184)) = v440
	v445 = *(*int32)(unsafe.Add(mBase, uint32(v55+v387<<(uint(int32(2))%32))))
	v447 = *(*int32)(unsafe.Add(mBase, uint32(v58)+260))
	F_ExecARInsertTriggers(m, v57, v59, v445, int32(0), v447)
	mBase = m.M
	v449 = m.ExcPending
	if v449 != 0 {
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
	v461 = v387 + int32(1)
	if v461 != v56 {
		v387 = v461
		goto L61
	} else {
		goto L77
	}
L77:
	;
	goto L62
L78:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v58)+184)) = v359
	*(*uint8)(unsafe.Add(mBase, uint32(v58)+304)) = uint8(v354)
	goto L10
L79:
	;
	goto L78
L80:
	;
	v493 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_CopyMultiInsertInfoFlush[1])))
	if v493&int32(1) == int32(0) {
		goto L79
	} else {
		goto L81
	}
L81:
	;
	v498 = int32(_a_F_CopyMultiInsertInfoFlush_0)
	v500 = *(*int32)(unsafe.Add(mBase, _c_F_CopyMultiInsertInfoFlush[2]))
	v501 = int32(1)
	*(*int32)(unsafe.Add(mBase, _c_F_CopyMultiInsertInfoFlush[2])) = v500 + v501
	v504 = *(*int32)(unsafe.Add(mBase, uint32(v489)))
	*(*int32)(unsafe.Add(mBase, uint32(v489))) = v504 + v501
	*(*int64)(unsafe.Add(mBase, uint32(v489+int32(16))+232)) = v484
	v512 = *(*int32)(unsafe.Add(mBase, uint32(v489)))
	*(*int32)(unsafe.Add(mBase, uint32(v489))) = v512 + v501
	v518 = *(*int32)(unsafe.Add(mBase, _c_F_CopyMultiInsertInfoFlush[2]))
	*(*int32)(unsafe.Add(mBase, _c_F_CopyMultiInsertInfoFlush[2])) = v518 - v501
	goto L79
L82:
	;
	goto L9
L83:
	;
	v576 = v570
	goto L84
L84:
	;
	v592 = *(*int32)(unsafe.Add(mBase, uint32(v576)+4))
	if v592 < int32(33) {
		goto L1
	} else {
		goto L86
	}
L85:
	;
	goto L1
L86:
	;
	v595 = *(*int32)(unsafe.Add(mBase, uint32(v576)+12))
	v596 = *(*int32)(unsafe.Add(mBase, uint32(v595)))
	v597 = *(*int32)(unsafe.Add(mBase, uint32(v596)+4000))
	if l1 == v597 {
		goto L87
	} else {
		goto L88
	}
L87:
	;
	v599 = F_list_delete_first(m, v576)
	mBase = m.M
	v600 = m.ExcPending
	if v600 != 0 {
		goto L21
	} else {
		goto L90
	}
L88:
	;
	v609 = v596
	v610 = v597
	goto L89
L89:
	;
	v611 = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v610)+208)) = v611
	v613 = *(*int32)(unsafe.Add(mBase, uint32(v610)+84))
	if v613 == v611 {
		goto L92
	} else {
		goto L93
	}
L90:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0))) = v599
	v602 = F_lappend(m, v599, v596)
	mBase = m.M
	v603 = m.ExcPending
	if v603 != 0 {
		goto L21
	} else {
		goto L91
	}
L91:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0))) = v602
	v605 = *(*int32)(unsafe.Add(mBase, uint32(v602)+12))
	v606 = *(*int32)(unsafe.Add(mBase, uint32(v605)))
	v607 = *(*int32)(unsafe.Add(mBase, uint32(v606)+4000))
	v609 = v606
	v610 = v607
	goto L89
L92:
	;
	v616 = *(*int32)(unsafe.Add(mBase, uint32(v609)+4004))
	F_FreeBulkInsertState(m, v616)
	mBase = m.M
	v618 = m.ExcPending
	if v618 != 0 {
		goto L21
	} else {
		goto L95
	}
L93:
	;
	goto L94
L94:
	;
	v623 = int32(0)
	goto L96
L95:
	;
	goto L94
L96:
	;
	v642 = *(*int32)(unsafe.Add(mBase, uint32(v609+v623<<(uint(int32(2))%32))))
	if v642 != 0 {
		goto L98
	} else {
		goto L99
	}
L97:
	;
	v650 = *(*int32)(unsafe.Add(mBase, uint32(v610)+84))
	if v650 != 0 {
		goto L103
	} else {
		goto L104
	}
L98:
	;
	F_ExecDropSingleTupleTableSlot(m, v642)
	mBase = m.M
	v644 = m.ExcPending
	if v644 != 0 {
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
	v646 = v623 + int32(1)
	if v646 != int32(1000) {
		v623 = v646
		goto L96
	} else {
		goto L102
	}
L102:
	;
	goto L100
L103:
	;
	F_pfree(m, v609)
	mBase = m.M
	v664 = m.ExcPending
	if v664 != 0 {
		goto L21
	} else {
		goto L108
	}
L104:
	;
	v651 = *(*int32)(unsafe.Add(mBase, uint32(v610)+8))
	v652 = *(*int32)(unsafe.Add(mBase, uint32(v651)+188))
	if v652 == int32(0) {
		goto L103
	} else {
		goto L105
	}
L105:
	;
	v655 = *(*int32)(unsafe.Add(mBase, uint32(v652)+108))
	if v655 == int32(0) {
		goto L103
	} else {
		goto L106
	}
L106:
	;
	v658 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
	m.T0[v655].(func(*base.Module, int32, int32))(m, v651, v658)
	mBase = m.M
	v660 = m.ExcPending
	if v660 != 0 {
		goto L21
	} else {
		goto L107
	}
L107:
	;
	goto L103
L108:
	;
	v665 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v666 = F_list_delete_first(m, v665)
	mBase = m.M
	v667 = m.ExcPending
	if v667 != 0 {
		goto L21
	} else {
		goto L109
	}
L109:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0))) = v666
	if v666 != 0 {
		v576 = v666
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
	var v17 int32
	_ = v17
	var v19 int32
	_ = v19
	var v25 int32
	_ = v25
	var v35 int32
	_ = v35
	var v49 int32
	_ = v49
	var v52 int32
	_ = v52
	var v53 int32
	_ = v53
	var v55 int32
	_ = v55
	var v56 int32
	_ = v56
	var v59 int32
	_ = v59
	var v64 int32
	_ = v64
	var v66 int32
	_ = v66
	var v67 int32
	_ = v67
	var v69 int32
	_ = v69
	var v72 int32
	_ = v72
	var v75 int32
	_ = v75
	var v79 int32
	_ = v79
	var v88 int32
	_ = v88
	var v91 int32
	_ = v91
	var v111 int32
	_ = v111
	var v113 int32
	_ = v113
	var v117 int32
	_ = v117
	var v121 int32
	_ = v121
	var v125 int32
	_ = v125
	var v126 int32
	_ = v126
	var v127 int32
	_ = v127
	var v129 int32
	_ = v129
	var v130 int32
	_ = v130
	var v133 int32
	_ = v133
	var v135 int32
	_ = v135
	var v137 int32
	_ = v137
	var v138 int32
	_ = v138
	var v142 int32
	_ = v142
	var v154 int32
	_ = v154
	var v155 int32
	_ = v155
	var v159 int32
	_ = v159
	var v169 int32
	_ = v169
	var v170 int32
	_ = v170
	var v171 int32
	_ = v171
	var v175 int32
	_ = v175
	var v176 int32
	_ = v176
	var v180 int32
	_ = v180
	var v181 int32
	_ = v181
	var v182 int32
	_ = v182
	var v183 int32
	_ = v183
	var v185 int32
	_ = v185
	var v192 int32
	_ = v192
	var v193 int32
	_ = v193
	var v208 int32
	_ = v208
	var v212 int32
	_ = v212
	var v213 int32
	_ = v213
	var v218 int32
	_ = v218
	var v231 int32
	_ = v231
	var v233 int32
	_ = v233
	var v239 int32
	_ = v239
	var v243 int32
	_ = v243
	var v262 int32
	_ = v262
	var v264 int32
	_ = v264
	var v268 int32
	_ = v268
	var v274 int32
	_ = v274
	var v278 int32
	_ = v278
	var v279 int32
	_ = v279
	var v281 int32
	_ = v281
	var v282 int32
	_ = v282
	var v283 int32
	_ = v283
	var v284 int32
	_ = v284
	var v286 int32
	_ = v286
	var v290 int32
	_ = v290
	var v297 int32
	_ = v297
	var v299 int32
	_ = v299
	var v300 int32
	_ = v300
	var v302 int32
	_ = v302
	var v304 int32
	_ = v304
	var v305 int32
	_ = v305
	var v308 int32
	_ = v308
	var v310 int32
	_ = v310
	var v311 int32
	_ = v311
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
	var v323 int32
	_ = v323
	var v329 int32
	_ = v329
	var v331 int32
	_ = v331
	var v333 int32
	_ = v333
	var v336 int32
	_ = v336
	var v340 int32
	_ = v340
	var v343 int32
	_ = v343
	var v345 int32
	_ = v345
	var v346 int32
	_ = v346
	var v349 int32
	_ = v349
	var v352 int32
	_ = v352
	var v354 int32
	_ = v354
	var v355 int32
	_ = v355
	var v356 int32
	_ = v356
	var v359 int32
	_ = v359
	var v360 int32
	_ = v360
	var v362 int32
	_ = v362
	var v363 int32
	_ = v363
	var v367 int32
	_ = v367
	var v368 int32
	_ = v368
	var v369 int32
	_ = v369
	var v370 int32
	_ = v370
	var v375 int32
	_ = v375
	var v378 int32
	_ = v378
	var v380 int32
	_ = v380
	var v381 int32
	_ = v381
	var v385 int32
	_ = v385
	var v386 int32
	_ = v386
	var v389 int32
	_ = v389
	var v390 int32
	_ = v390
	var v395 int32
	_ = v395
	var v397 int32
	_ = v397
	var v402 int32
	_ = v402
	var v403 int32
	_ = v403
	var v404 int32
	_ = v404
	var v405 int32
	_ = v405
	var v406 int32
	_ = v406
	var v407 int32
	_ = v407
	var v413 int64
	_ = v413
	var v418 int32
	_ = v418
	var v420 int32
	_ = v420
	var v422 int32
	_ = v422
	var v423 int64
	_ = v423
	var v425 int32
	_ = v425
	var v427 int32
	_ = v427
	var v428 int32
	_ = v428
	var v431 int32
	_ = v431
	var v434 int32
	_ = v434
	var v436 int32
	_ = v436
	var v437 int32
	_ = v437
	var v438 int32
	_ = v438
	var v441 int32
	_ = v441
	var v442 int32
	_ = v442
	var v444 int32
	_ = v444
	var v446 int32
	_ = v446
	var v447 int32
	_ = v447
	var v448 int32
	_ = v448
	var v449 int64
	_ = v449
	var v450 int32
	_ = v450
	var v451 int32
	_ = v451
	var v454 int32
	_ = v454
	var v461 int32
	_ = v461
	var v462 int32
	_ = v462
	var v464 int32
	_ = v464
	var v467 int32
	_ = v467
	var v477 int32
	_ = v477
	var v478 int32
	_ = v478
	var v479 int32
	_ = v479
	var v480 int32
	_ = v480
	var v483 int32
	_ = v483
	var v488 int32
	_ = v488
	var v491 int32
	_ = v491
	var v495 int32
	_ = v495
	var v500 int32
	_ = v500
	var v504 int32
	_ = v504
	var v507 int32
	_ = v507
	var v513 int32
	_ = v513
	var v518 int32
	_ = v518
	var v522 int32
	_ = v522
	var v525 int32
	_ = v525
	var v531 int32
	_ = v531
	var v536 int32
	_ = v536
	var v542 int32
	_ = v542
	var v543 int32
	_ = v543
	var v554 int32
	_ = v554
	var v556 int32
	_ = v556
	var v563 int32
	_ = v563
	var v568 int32
	_ = v568
	var v580 int32
	_ = v580
	v4 = int32(0)
	v17 = m.G0
	v19 = v17 - int32(48)
	m.G0 = v19
	if l0 == v4 {
		v563 = int32(-1)
		v568 = v4
		goto L2
	} else {
		goto L3
	}
L1:
	;
	m.G0 = v19 + int32(48)
	return v580
L2:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l1))) = v568
	v580 = v563
	goto L1
L3:
	;
	v25 = *(*int32)(unsafe.Add(mBase, _c_F_GetMultiXactIdMembers[0]))
	if base.B2i32(v25 == int32(0))|base.B2i32(v25 == int32(_a_F_GetMultiXactIdMembers_0)) != 0 {
		goto L4
	} else {
		goto L5
	}
L4:
	;
	v111 = *(*int32)(unsafe.Add(mBase, _c_F_GetMultiXactIdMembers[1]))
	v113 = *(*int32)(unsafe.Add(mBase, _c_F_GetMultiXactIdMembers[2]))
	v117 = *(*int32)(unsafe.Add(mBase, uint32(v111+v113<<(uint(int32(2))%32))))
	if v117 == int32(0) {
		goto L24
	} else {
		goto L25
	}
L5:
	;
	v35 = v25
	goto L6
L6:
	;
	v49 = *(*int32)(unsafe.Add(mBase, uint32(v35-int32(8))))
	if l0 == v49 {
		goto L8
	} else {
		goto L9
	}
L7:
	;
	goto L4
L8:
	;
	v52 = v35 - int32(4)
	v53 = *(*int32)(unsafe.Add(mBase, uint32(v52)))
	v55 = v53 << (uint(int32(3)) % 32)
	v56 = F_palloc(m, v55)
	mBase = m.M
	v59 = m.ExcPending
	if v59 != 0 {
		goto L11
	} else {
		goto L12
	}
L9:
	;
	goto L10
L10:
	;
	v91 = *(*int32)(unsafe.Add(mBase, uint32(v35)+4))
	if v91 != int32(_a_F_GetMultiXactIdMembers_0) {
		v35 = v91
		goto L6
	} else {
		goto L23
	}
L11:
	;
	return int32(0)
L12:
	;
	if v55 != 0 {
		goto L13
	} else {
		goto L14
	}
L13:
	;
	base.MemoryCopy(m, v56, v35+int32(8), v55)
	goto L15
L14:
	;
	goto L15
L15:
	;
	v64 = *(*int32)(unsafe.Add(mBase, _c_F_GetMultiXactIdMembers[0]))
	if v35 != v64 {
		goto L16
	} else {
		goto L17
	}
L16:
	;
	v66 = *(*int32)(unsafe.Add(mBase, uint32(v35)))
	v67 = *(*int32)(unsafe.Add(mBase, uint32(v35)+4))
	*(*int32)(unsafe.Add(mBase, uint32(v66)+4)) = v67
	v69 = *(*int32)(unsafe.Add(mBase, uint32(v35)))
	*(*int32)(unsafe.Add(mBase, uint32(v67))) = v69
	v72 = *(*int32)(unsafe.Add(mBase, _c_F_GetMultiXactIdMembers[0]))
	if v72 == int32(0) {
		goto L19
	} else {
		goto L20
	}
L17:
	;
	goto L18
L18:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l1))) = v56
	v88 = *(*int32)(unsafe.Add(mBase, uint32(v52)))
	if int32(0) <= v88 {
		v580 = v88
		goto L1
	} else {
		goto L22
	}
L19:
	;
	v75 = int32(_a_F_GetMultiXactIdMembers_0)
	*(*int32)(unsafe.Add(mBase, _c_F_GetMultiXactIdMembers[3])) = v75
	v79 = v75
	goto L21
L20:
	;
	v79 = v72
	goto L21
L21:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v35))) = int32(_a_F_GetMultiXactIdMembers_0)
	*(*int32)(unsafe.Add(mBase, uint32(v35)+4)) = v79
	*(*int32)(unsafe.Add(mBase, uint32(v79))) = v35
	*(*int32)(unsafe.Add(mBase, _c_F_GetMultiXactIdMembers[0])) = v35
	goto L18
L22:
	;
	goto L4
L23:
	;
	goto L7
L24:
	;
	v121 = *(*int32)(unsafe.Add(mBase, _c_F_GetMultiXactIdMembers[4]))
	v125 = F_LWLockAcquire(m, v121+int32(1664), int32(0))
	mBase = m.M
	v126 = m.ExcPending
	if v126 != 0 {
		goto L11
	} else {
		goto L27
	}
L25:
	;
	goto L26
L26:
	;
	if l2 != 0 {
		goto L60
	} else {
		goto L61
	}
L27:
	;
	v127 = int32(1)
	v129 = *(*int32)(unsafe.Add(mBase, _c_F_GetMultiXactIdMembers[5]))
	v130 = *(*int32)(unsafe.Add(mBase, uint32(v129)))
	if base.Ui32(v130) <= base.Ui32(v127) {
		goto L28
	} else {
		goto L29
	}
L28:
	;
	v133 = v127
	goto L30
L29:
	;
	v133 = v130
	goto L30
L30:
	;
	v135 = *(*int32)(unsafe.Add(mBase, _c_F_GetMultiXactIdMembers[6]))
	v137 = *(*int32)(unsafe.Add(mBase, _c_F_GetMultiXactIdMembers[7]))
	v138 = v135 + v137
	if v138 <= int32(0) {
		v218 = v133
		goto L31
	} else {
		goto L32
	}
L31:
	;
	v231 = *(*int32)(unsafe.Add(mBase, _c_F_GetMultiXactIdMembers[1]))
	v233 = *(*int32)(unsafe.Add(mBase, _c_F_GetMultiXactIdMembers[2]))
	*(*int32)(unsafe.Add(mBase, uint32(v231+v233<<(uint(int32(2))%32)))) = v218
	v239 = *(*int32)(unsafe.Add(mBase, _c_F_GetMultiXactIdMembers[4]))
	F_LWLockRelease(m, v239+int32(1664))
	mBase = m.M
	v243 = m.ExcPending
	if v243 != 0 {
		goto L11
	} else {
		goto L59
	}
L32:
	;
	v142 = *(*int32)(unsafe.Add(mBase, _c_F_GetMultiXactIdMembers[8]))
	if v138 == int32(1) {
		goto L34
	} else {
		goto L35
	}
L33:
	;
	v208 = *(*int32)(unsafe.Add(mBase, uint32(v142+v192<<(uint(int32(2))%32))))
	if v208-v193 < int32(0) {
		goto L53
	} else {
		goto L54
	}
L34:
	;
	v192 = int32(0)
	v193 = v133
	goto L33
L35:
	;
	goto L36
L36:
	;
	v154 = int32(0)
	v155 = v133
	v159 = v4
	goto L37
L37:
	;
	v169 = v142 + v154<<(uint(int32(2))%32)
	v170 = *(*int32)(unsafe.Add(mBase, uint32(v169)+4))
	v171 = *(*int32)(unsafe.Add(mBase, uint32(v169)))
	if v171-v155 < int32(0) {
		goto L39
	} else {
		goto L40
	}
L38:
	;
	if v138&int32(1) == int32(0) {
		v218 = v181
		goto L31
	} else {
		goto L52
	}
L39:
	;
	v175 = v171
	goto L41
L40:
	;
	v175 = v155
	goto L41
L41:
	;
	if v171 != 0 {
		goto L42
	} else {
		goto L43
	}
L42:
	;
	v176 = v175
	goto L44
L43:
	;
	v176 = v155
	goto L44
L44:
	;
	if v170-v176 < int32(0) {
		goto L45
	} else {
		goto L46
	}
L45:
	;
	v180 = v170
	goto L47
L46:
	;
	v180 = v176
	goto L47
L47:
	;
	if v170 != 0 {
		goto L48
	} else {
		goto L49
	}
L48:
	;
	v181 = v180
	goto L50
L49:
	;
	v181 = v176
	goto L50
L50:
	;
	v182 = int32(2)
	v183 = v154 + v182
	v185 = v159 + v182
	if v185 != v138&int32(2147483646) {
		v154 = v183
		v155 = v181
		v159 = v185
		goto L37
	} else {
		goto L51
	}
L51:
	;
	goto L38
L52:
	;
	v192 = v183
	v193 = v181
	goto L33
L53:
	;
	v212 = v208
	goto L55
L54:
	;
	v212 = v193
	goto L55
L55:
	;
	if v208 != 0 {
		goto L56
	} else {
		goto L57
	}
L56:
	;
	v213 = v212
	goto L58
L57:
	;
	v213 = v193
	goto L58
L58:
	;
	v218 = v213
	goto L31
L59:
	;
	goto L26
L60:
	;
	v262 = *(*int32)(unsafe.Add(mBase, _c_F_GetMultiXactIdMembers[1]))
	v264 = *(*int32)(unsafe.Add(mBase, _c_F_GetMultiXactIdMembers[2]))
	v268 = *(*int32)(unsafe.Add(mBase, uint32(v262+v264<<(uint(int32(2))%32))))
	if l0-v268 < int32(0) {
		v563 = int32(-1)
		v568 = v4
		goto L2
	} else {
		goto L63
	}
L61:
	;
	goto L62
L62:
	;
	v274 = *(*int32)(unsafe.Add(mBase, _c_F_GetMultiXactIdMembers[4]))
	v278 = F_LWLockAcquire(m, v274+int32(1664), int32(1))
	mBase = m.M
	v279 = m.ExcPending
	if v279 != 0 {
		goto L11
	} else {
		goto L64
	}
L63:
	;
	goto L62
L64:
	;
	v281 = *(*int32)(unsafe.Add(mBase, _c_F_GetMultiXactIdMembers[5]))
	v282 = *(*int32)(unsafe.Add(mBase, uint32(v281)+4))
	v283 = *(*int32)(unsafe.Add(mBase, uint32(v281)))
	v284 = *(*int32)(unsafe.Add(mBase, uint32(v281)+12))
	v286 = *(*int32)(unsafe.Add(mBase, _c_F_GetMultiXactIdMembers[4]))
	F_LWLockRelease(m, v286+int32(1664))
	mBase = m.M
	v290 = m.ExcPending
	if v290 != 0 {
		goto L11
	} else {
		goto L65
	}
L65:
	;
	if int32(0) <= l0-v284 {
		goto L69
	} else {
		goto L70
	}
L66:
	;
	F_LWLockRelease(m, v542)
	mBase = m.M
	v554 = m.ExcPending
	if v554 != 0 {
		goto L11
	} else {
		goto L129
	}
L67:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v522 = m.ExcPending
	if v522 != 0 {
		goto L11
	} else {
		goto L125
	}
L68:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v504 = m.ExcPending
	if v504 != 0 {
		goto L11
	} else {
		goto L121
	}
L69:
	;
	if int32(0) <= l0-v283 {
		goto L68
	} else {
		goto L72
	}
L70:
	;
	goto L71
L71:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v488 = m.ExcPending
	if v488 != 0 {
		goto L11
	} else {
		goto L117
	}
L72:
	;
	v297 = int32(0)
	v299 = *(*int32)(unsafe.Add(mBase, _c_F_GetMultiXactIdMembers[9]))
	v300 = *(*int32)(unsafe.Add(mBase, uint32(v299)+28))
	v302 = int32(base.Ui32(l0) >> (uint(int32(11)) % 32))
	v304 = int32(*(*uint16)(unsafe.Add(mBase, _c_F_GetMultiXactIdMembers[10])))
	v305 = base.I32_rem_u_s(v302, v304)
	v308 = v300 + v305<<(uint(int32(7))%32)
	v310 = F_LWLockAcquire(m, v308, v297)
	mBase = m.M
	v311 = m.ExcPending
	if v311 != 0 {
		goto L11
	} else {
		goto L73
	}
L73:
	;
	v315 = F_SimpleLruReadPage(m, int32(_a_F_GetMultiXactIdMembers_1), base.I64_extend_i32_u(v302), int32(1), l0)
	mBase = m.M
	v316 = m.ExcPending
	if v316 != 0 {
		goto L11
	} else {
		goto L74
	}
L74:
	;
	v318 = *(*int32)(unsafe.Add(mBase, _c_F_GetMultiXactIdMembers[9]))
	v319 = *(*int32)(unsafe.Add(mBase, uint32(v318)+4))
	v320 = int32(2)
	v323 = *(*int32)(unsafe.Add(mBase, uint32(v319+v315<<(uint(v320)%32))))
	v329 = *(*int32)(unsafe.Add(mBase, uint32(v323+l0&int32(2047)<<(uint(v320)%32))))
	v331 = l0 + int32(1)
	if v283 == v331 {
		goto L76
	} else {
		goto L77
	}
L75:
	;
	F_LWLockRelease(m, v378)
	mBase = m.M
	v385 = m.ExcPending
	if v385 != 0 {
		goto L11
	} else {
		goto L94
	}
L76:
	;
	v378 = v308
	v380 = v315
	v381 = v282
	goto L75
L77:
	;
	goto L78
L78:
	;
	v333 = int32(1)
	if base.Ui32(v331) <= base.Ui32(v333) {
		goto L79
	} else {
		goto L80
	}
L79:
	;
	v336 = v333
	goto L81
L80:
	;
	v336 = v331
	goto L81
L81:
	;
	v340 = int32(base.Ui32(v336) >> (uint(int32(11)) % 32))
	if v302 == v340 {
		goto L83
	} else {
		goto L84
	}
L82:
	;
	v375 = *(*int32)(unsafe.Add(mBase, uint32(v369+v336&int32(2047)<<(uint(int32(2))%32))))
	if v375 == int32(0) {
		goto L67
	} else {
		goto L93
	}
L83:
	;
	v368 = v308
	v369 = v323
	v370 = v315
	goto L82
L84:
	;
	goto L85
L85:
	;
	v343 = *(*int32)(unsafe.Add(mBase, uint32(v318)+28))
	v345 = int32(*(*uint16)(unsafe.Add(mBase, _c_F_GetMultiXactIdMembers[10])))
	v346 = base.I32_rem_u_s(v340, v345)
	v349 = v343 + v346<<(uint(int32(7))%32)
	if v308 == v349 {
		goto L87
	} else {
		goto L88
	}
L86:
	;
	v359 = F_SimpleLruReadPage(m, int32(_a_F_GetMultiXactIdMembers_1), base.I64_extend_i32_u(v340), int32(1), v336)
	mBase = m.M
	v360 = m.ExcPending
	if v360 != 0 {
		goto L11
	} else {
		goto L92
	}
L87:
	;
	v356 = v308
	goto L86
L88:
	;
	goto L89
L89:
	;
	F_LWLockRelease(m, v308)
	mBase = m.M
	v352 = m.ExcPending
	if v352 != 0 {
		goto L11
	} else {
		goto L90
	}
L90:
	;
	v354 = F_LWLockAcquire(m, v349, int32(0))
	mBase = m.M
	v355 = m.ExcPending
	if v355 != 0 {
		goto L11
	} else {
		goto L91
	}
L91:
	;
	v356 = v349
	goto L86
L92:
	;
	v362 = *(*int32)(unsafe.Add(mBase, _c_F_GetMultiXactIdMembers[9]))
	v363 = *(*int32)(unsafe.Add(mBase, uint32(v362)+4))
	v367 = *(*int32)(unsafe.Add(mBase, uint32(v363+v359<<(uint(int32(2))%32))))
	v368 = v356
	v369 = v367
	v370 = v359
	goto L82
L93:
	;
	v378 = v368
	v380 = v370
	v381 = v375
	goto L75
L94:
	;
	v386 = v381 - v329
	v389 = F_palloc(m, v386<<(uint(int32(3))%32))
	mBase = m.M
	v390 = m.ExcPending
	if v390 != 0 {
		goto L11
	} else {
		goto L95
	}
L95:
	;
	if v386 <= int32(0) {
		goto L96
	} else {
		goto L97
	}
L96:
	;
	v542 = int32(0)
	v543 = v297
	goto L66
L97:
	;
	goto L98
L98:
	;
	v395 = *(*int32)(unsafe.Add(mBase, _c_F_GetMultiXactIdMembers[11]))
	v397 = int32(0)
	v402 = v395
	v403 = v329
	v404 = v397
	v405 = v297
	v406 = v380
	v407 = v397
	v413 = int64(-1)
	goto L99
L99:
	;
	v418 = base.I32_rem_u_s(int32(base.Ui32(v403)>>(uint(int32(2))%32)), int32(409))
	v420 = v418 * int32(20)
	v422 = base.I32_div_u_s(v403, int32(1636))
	v423 = base.I64_extend_i32_u(v422)
	if v423 != v413 {
		goto L101
	} else {
		goto L102
	}
L100:
	;
	v542 = v447
	v543 = v479
	goto L66
L101:
	;
	v425 = *(*int32)(unsafe.Add(mBase, uint32(v402)+28))
	v427 = int32(*(*uint16)(unsafe.Add(mBase, _c_F_GetMultiXactIdMembers[12])))
	v428 = base.I32_rem_u_s(v422, v427)
	v431 = v425 + v428<<(uint(int32(7))%32)
	if v404 != v431 {
		goto L104
	} else {
		goto L105
	}
L102:
	;
	v446 = v402
	v447 = v404
	v448 = v406
	v449 = v413
	goto L103
L103:
	;
	v450 = *(*int32)(unsafe.Add(mBase, uint32(v446)+4))
	v451 = int32(2)
	v454 = *(*int32)(unsafe.Add(mBase, uint32(v450+v448<<(uint(v451)%32))))
	v461 = *(*int32)(unsafe.Add(mBase, uint32(v420+(v454+v403<<(uint(v451)%32)&int32(12)))+4))
	if v461 != 0 {
		goto L113
	} else {
		goto L114
	}
L104:
	;
	if v404 != 0 {
		goto L107
	} else {
		goto L108
	}
L105:
	;
	v438 = v404
	goto L106
L106:
	;
	v441 = F_SimpleLruReadPage(m, int32(_a_F_GetMultiXactIdMembers_2), v423, int32(1), l0)
	mBase = m.M
	v442 = m.ExcPending
	if v442 != 0 {
		goto L11
	} else {
		goto L112
	}
L107:
	;
	F_LWLockRelease(m, v404)
	mBase = m.M
	v434 = m.ExcPending
	if v434 != 0 {
		goto L11
	} else {
		goto L110
	}
L108:
	;
	goto L109
L109:
	;
	v436 = F_LWLockAcquire(m, v431, int32(0))
	mBase = m.M
	v437 = m.ExcPending
	if v437 != 0 {
		goto L11
	} else {
		goto L111
	}
L110:
	;
	goto L109
L111:
	;
	v438 = v431
	goto L106
L112:
	;
	v444 = *(*int32)(unsafe.Add(mBase, _c_F_GetMultiXactIdMembers[11]))
	v446 = v444
	v447 = v438
	v448 = v441
	v449 = v423
	goto L103
L113:
	;
	v462 = int32(3)
	v464 = v389 + v405<<(uint(v462)%32)
	*(*int32)(unsafe.Add(mBase, uint32(v464))) = v461
	v467 = *(*int32)(unsafe.Add(mBase, uint32(v454+v420)))
	*(*int32)(unsafe.Add(mBase, uint32(v464)+4)) = int32(base.Ui32(v467)>>(uint(v403<<(uint(v462)%32))%32)) & int32(255)
	v477 = *(*int32)(unsafe.Add(mBase, _c_F_GetMultiXactIdMembers[11]))
	v478 = v477
	v479 = v405 + int32(1)
	goto L115
L114:
	;
	v478 = v446
	v479 = v405
	goto L115
L115:
	;
	v480 = int32(1)
	v483 = v407 + v480
	if v386 != v483 {
		v402 = v478
		v403 = v403 + v480
		v404 = v447
		v405 = v479
		v406 = v448
		v407 = v483
		v413 = v449
		goto L99
	} else {
		goto L116
	}
L116:
	;
	goto L100
L117:
	;
	F_errcode(m, int32(2600))
	mBase = m.M
	v491 = m.ExcPending
	if v491 != 0 {
		goto L11
	} else {
		goto L118
	}
L118:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v19))) = l0
	F_errmsg(m, int32(_a_F_GetMultiXactIdMembers_3), v19)
	mBase = m.M
	v495 = m.ExcPending
	if v495 != 0 {
		goto L11
	} else {
		goto L119
	}
L119:
	;
	F_errfinish(m, int32(_a_F_GetMultiXactIdMembers_4), int32(1462), int32(_a_F_GetMultiXactIdMembers_5))
	mBase = m.M
	v500 = m.ExcPending
	if v500 != 0 {
		goto L11
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
	F_errcode(m, int32(2600))
	mBase = m.M
	v507 = m.ExcPending
	if v507 != 0 {
		goto L11
	} else {
		goto L122
	}
L122:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v19)+32)) = l0
	F_errmsg(m, int32(_a_F_GetMultiXactIdMembers_6), v19+int32(32))
	mBase = m.M
	v513 = m.ExcPending
	if v513 != 0 {
		goto L11
	} else {
		goto L123
	}
L123:
	;
	F_errfinish(m, int32(_a_F_GetMultiXactIdMembers_4), int32(1468), int32(_a_F_GetMultiXactIdMembers_5))
	mBase = m.M
	v518 = m.ExcPending
	if v518 != 0 {
		goto L11
	} else {
		goto L124
	}
L124:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L125:
	;
	F_errcode(m, int32(16779816))
	mBase = m.M
	v525 = m.ExcPending
	if v525 != 0 {
		goto L11
	} else {
		goto L126
	}
L126:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v19)+16)) = l0
	F_errmsg(m, int32(_a_F_GetMultiXactIdMembers_7), v19+int32(16))
	mBase = m.M
	v531 = m.ExcPending
	if v531 != 0 {
		goto L11
	} else {
		goto L127
	}
L127:
	;
	F_errfinish(m, int32(_a_F_GetMultiXactIdMembers_4), int32(1561), int32(_a_F_GetMultiXactIdMembers_5))
	mBase = m.M
	v536 = m.ExcPending
	if v536 != 0 {
		goto L11
	} else {
		goto L128
	}
L128:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L129:
	;
	F_mXactCachePut(m, l0, v543, v389)
	mBase = m.M
	v556 = m.ExcPending
	if v556 != 0 {
		goto L11
	} else {
		goto L130
	}
L130:
	;
	v563 = v543
	v568 = v389
	goto L2
}
func F_MultiXactIdCreateFromMembers(m *base.Module, l0 int32, l1 int32) int32 {
	mBase := m.M
	_ = mBase
	var v12 int32
	_ = v12
	var v14 int32
	_ = v14
	var v21 int32
	_ = v21
	var v23 int32
	_ = v23
	var v30 int32
	_ = v30
	var v33 int32
	_ = v33
	var v44 int32
	_ = v44
	var v47 int32
	_ = v47
	var v55 int32
	_ = v55
	var v56 int32
	_ = v56
	var v57 int32
	_ = v57
	var v60 int32
	_ = v60
	var v61 int32
	_ = v61
	var v63 int32
	_ = v63
	var v64 int32
	_ = v64
	var v66 int32
	_ = v66
	var v68 int32
	_ = v68
	var v71 int32
	_ = v71
	var v72 int32
	_ = v72
	var v73 int32
	_ = v73
	var v78 int32
	_ = v78
	var v79 int32
	_ = v79
	var v80 int32
	_ = v80
	var v83 int32
	_ = v83
	var v84 int32
	_ = v84
	var v85 int32
	_ = v85
	var v88 int32
	_ = v88
	var v89 int32
	_ = v89
	var v91 int32
	_ = v91
	var v96 int32
	_ = v96
	var v109 int32
	_ = v109
	var v111 int32
	_ = v111
	var v112 int32
	_ = v112
	var v114 int32
	_ = v114
	var v117 int32
	_ = v117
	var v120 int32
	_ = v120
	var v124 int32
	_ = v124
	var v135 int32
	_ = v135
	var v136 int32
	_ = v136
	var v155 int32
	_ = v155
	var v159 int32
	_ = v159
	var v169 int32
	_ = v169
	var v177 int32
	_ = v177
	var v192 int32
	_ = v192
	var v197 int32
	_ = v197
	var v198 int32
	_ = v198
	var v200 int32
	_ = v200
	var v202 int32
	_ = v202
	var v204 int32
	_ = v204
	var v208 int32
	_ = v208
	var v209 int32
	_ = v209
	var v211 int32
	_ = v211
	var v212 int32
	_ = v212
	var v215 int32
	_ = v215
	var v218 int32
	_ = v218
	var v219 int32
	_ = v219
	var v223 int32
	_ = v223
	var v224 int32
	_ = v224
	var v225 int32
	_ = v225
	var v226 int32
	_ = v226
	var v228 int32
	_ = v228
	var v232 int32
	_ = v232
	var v234 int32
	_ = v234
	var v240 int32
	_ = v240
	var v241 int32
	_ = v241
	var v244 int32
	_ = v244
	var v248 int32
	_ = v248
	var v251 int32
	_ = v251
	var v257 int32
	_ = v257
	var v261 int32
	_ = v261
	var v266 int32
	_ = v266
	var v271 int32
	_ = v271
	var v275 int32
	_ = v275
	var v276 int32
	_ = v276
	var v279 int32
	_ = v279
	var v280 int32
	_ = v280
	var v284 int32
	_ = v284
	var v291 int32
	_ = v291
	var v296 int32
	_ = v296
	var v303 int32
	_ = v303
	var v306 int32
	_ = v306
	var v310 int32
	_ = v310
	var v314 int32
	_ = v314
	var v318 int32
	_ = v318
	var v322 int32
	_ = v322
	var v323 int32
	_ = v323
	var v324 int32
	_ = v324
	var v326 int32
	_ = v326
	var v327 int32
	_ = v327
	var v330 int32
	_ = v330
	var v331 int32
	_ = v331
	var v335 int32
	_ = v335
	var v338 int32
	_ = v338
	var v344 int32
	_ = v344
	var v345 int32
	_ = v345
	var v347 int32
	_ = v347
	var v349 int32
	_ = v349
	var v350 int32
	_ = v350
	var v353 int32
	_ = v353
	var v355 int32
	_ = v355
	var v356 int32
	_ = v356
	var v358 int64
	_ = v358
	var v359 int32
	_ = v359
	var v360 int32
	_ = v360
	var v363 int32
	_ = v363
	var v368 int32
	_ = v368
	var v371 int64
	_ = v371
	var v372 int32
	_ = v372
	var v374 int32
	_ = v374
	var v376 int32
	_ = v376
	var v377 int32
	_ = v377
	var v381 int32
	_ = v381
	var v384 int32
	_ = v384
	var v385 int32
	_ = v385
	var v388 int32
	_ = v388
	var v390 int32
	_ = v390
	var v391 int32
	_ = v391
	var v397 int32
	_ = v397
	var v401 int32
	_ = v401
	var v402 int32
	_ = v402
	var v404 int32
	_ = v404
	var v409 int32
	_ = v409
	var v410 int32
	_ = v410
	var v412 int32
	_ = v412
	var v421 int32
	_ = v421
	var v423 int32
	_ = v423
	var v424 int32
	_ = v424
	var v428 int32
	_ = v428
	var v429 int32
	_ = v429
	var v431 int32
	_ = v431
	var v433 int32
	_ = v433
	var v435 int32
	_ = v435
	var v445 int32
	_ = v445
	var v446 int32
	_ = v446
	var v451 int32
	_ = v451
	var v453 int32
	_ = v453
	var v454 int32
	_ = v454
	var v455 int32
	_ = v455
	var v458 int32
	_ = v458
	var v465 int32
	_ = v465
	var v469 int32
	_ = v469
	var v474 int32
	_ = v474
	var v479 int32
	_ = v479
	var v482 int32
	_ = v482
	var v489 int32
	_ = v489
	var v491 int32
	_ = v491
	var v499 int32
	_ = v499
	var v506 int32
	_ = v506
	var v507 int32
	_ = v507
	var v509 int32
	_ = v509
	var v511 int32
	_ = v511
	var v512 int32
	_ = v512
	var v515 int32
	_ = v515
	var v517 int32
	_ = v517
	var v518 int32
	_ = v518
	var v520 int64
	_ = v520
	var v521 int32
	_ = v521
	var v522 int32
	_ = v522
	var v525 int32
	_ = v525
	var v530 int32
	_ = v530
	var v533 int64
	_ = v533
	var v534 int32
	_ = v534
	var v536 int32
	_ = v536
	var v542 int32
	_ = v542
	var v544 int32
	_ = v544
	var v548 int32
	_ = v548
	var v550 int32
	_ = v550
	var v564 int32
	_ = v564
	var v566 int32
	_ = v566
	var v567 int32
	_ = v567
	var v571 int32
	_ = v571
	var v572 int32
	_ = v572
	var v576 int32
	_ = v576
	var v580 int32
	_ = v580
	var v584 int32
	_ = v584
	var v589 int32
	_ = v589
	var v594 int32
	_ = v594
	var v598 int32
	_ = v598
	var v601 int64
	_ = v601
	var v602 int32
	_ = v602
	var v604 int32
	_ = v604
	var v605 int32
	_ = v605
	var v607 int32
	_ = v607
	var v612 int32
	_ = v612
	var v621 int32
	_ = v621
	var v631 int32
	_ = v631
	var v632 int32
	_ = v632
	var v634 int32
	_ = v634
	var v637 int32
	_ = v637
	var v639 int32
	_ = v639
	var v641 int32
	_ = v641
	var v643 int32
	_ = v643
	var v644 int32
	_ = v644
	var v647 int32
	_ = v647
	var v653 int32
	_ = v653
	var v660 int32
	_ = v660
	var v668 int32
	_ = v668
	var v677 int32
	_ = v677
	var v678 int32
	_ = v678
	var v681 int32
	_ = v681
	var v685 int32
	_ = v685
	var v691 int32
	_ = v691
	var v693 int32
	_ = v693
	var v710 int32
	_ = v710
	var v713 int32
	_ = v713
	var v714 int32
	_ = v714
	var v715 int32
	_ = v715
	var v716 int32
	_ = v716
	var v718 int32
	_ = v718
	var v720 int32
	_ = v720
	var v722 int32
	_ = v722
	var v729 int32
	_ = v729
	var v735 int32
	_ = v735
	var v740 int32
	_ = v740
	var v744 int32
	_ = v744
	var v750 int32
	_ = v750
	var v755 int32
	_ = v755
	var v759 int32
	_ = v759
	var v764 int32
	_ = v764
	var v768 int32
	_ = v768
	var v772 int32
	_ = v772
	var v777 int32
	_ = v777
	var v783 int32
	_ = v783
	var v787 int32
	_ = v787
	var v792 int32
	_ = v792
	var v795 int32
	_ = v795
	var v799 int32
	_ = v799
	var v802 int32
	_ = v802
	var v806 int32
	_ = v806
	var v808 int32
	_ = v808
	var v809 int32
	_ = v809
	var v813 int32
	_ = v813
	var v820 int32
	_ = v820
	var v822 int32
	_ = v822
	var v823 int32
	_ = v823
	var v829 int32
	_ = v829
	var v834 int32
	_ = v834
	v12 = m.G0
	v14 = v12 - int32(144)
	m.G0 = v14
	F_pg_qsort(m, l1, l0, int32(8), int32(291))
	mBase = m.M
	v21 = m.ExcPending
	if v21 != 0 {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	return int32(0)
L2:
	;
	v23 = *(*int32)(unsafe.Add(mBase, _c_F_MultiXactIdCreateFromMembers[0]))
	if base.B2i32(v23 == int32(0))|base.B2i32(v23 == int32(_a_F_MultiXactIdCreateFromMembers_0)) != 0 {
		goto L8
	} else {
		goto L9
	}
L3:
	;
	F_SendPostmasterSignal(m, int32(4))
	mBase = m.M
	v795 = m.ExcPending
	if v795 != 0 {
		goto L1
	} else {
		goto L199
	}
L4:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v14)+112)) = v240
	F_errmsg(m, int32(_a_F_MultiXactIdCreateFromMembers_1), v14+int32(112))
	mBase = m.M
	v783 = m.ExcPending
	if v783 != 0 {
		goto L1
	} else {
		goto L196
	}
L5:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v768 = m.ExcPending
	if v768 != 0 {
		goto L1
	} else {
		goto L193
	}
L6:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v631 = m.ExcPending
	if v631 != 0 {
		goto L1
	} else {
		goto L162
	}
L7:
	;
	m.G0 = v14 + int32(144)
	return v621
L8:
	;
	if int32(0) < l0 {
		goto L41
	} else {
		goto L42
	}
L9:
	;
	v30 = l0 << (uint(int32(3)) % 32)
	v33 = v23
	goto L10
L10:
	;
	v44 = *(*int32)(unsafe.Add(mBase, uint32(v33-int32(4))))
	if v44 != l0 {
		goto L12
	} else {
		goto L13
	}
L11:
	;
	goto L8
L12:
	;
	v136 = *(*int32)(unsafe.Add(mBase, uint32(v33)+4))
	if v136 != int32(_a_F_MultiXactIdCreateFromMembers_0) {
		v33 = v136
		goto L10
	} else {
		goto L40
	}
L13:
	;
	v47 = v33 + int32(8)
	if base.Ui32(int32(4)) <= base.Ui32(v30) {
		goto L17
	} else {
		goto L18
	}
L14:
	;
	if v109 != 0 {
		goto L12
	} else {
		goto L32
	}
L15:
	;
	v109 = int32(0)
	goto L14
L16:
	;
	v83 = v78
	v84 = v79
	v85 = v80
	goto L26
L17:
	;
	if (l1|v47)&int32(3) != 0 {
		v78 = l1
		v79 = v47
		v80 = v30
		goto L16
	} else {
		goto L20
	}
L18:
	;
	v71 = l1
	v72 = v47
	v73 = v30
	goto L19
L19:
	;
	if v73 == int32(0) {
		goto L15
	} else {
		goto L25
	}
L20:
	;
	v55 = l1
	v56 = v47
	v57 = v30
	goto L21
L21:
	;
	v60 = *(*int32)(unsafe.Add(mBase, uint32(v55)))
	v61 = *(*int32)(unsafe.Add(mBase, uint32(v56)))
	if v60 != v61 {
		v78 = v55
		v79 = v56
		v80 = v57
		goto L16
	} else {
		goto L23
	}
L22:
	;
	v71 = v66
	v72 = v64
	v73 = v68
	goto L19
L23:
	;
	v63 = int32(4)
	v64 = v56 + v63
	v66 = v55 + v63
	v68 = v57 - v63
	if base.Ui32(int32(3)) < base.Ui32(v68) {
		v55 = v66
		v56 = v64
		v57 = v68
		goto L21
	} else {
		goto L24
	}
L24:
	;
	goto L22
L25:
	;
	v78 = v71
	v79 = v72
	v80 = v73
	goto L16
L26:
	;
	v88 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v83))))
	v89 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v84))))
	if v88 == v89 {
		goto L28
	} else {
		goto L29
	}
L27:
	;
	v109 = v88 - v89
	goto L14
L28:
	;
	v91 = int32(1)
	v96 = v85 - v91
	if v96 != 0 {
		v83 = v83 + v91
		v84 = v84 + v91
		v85 = v96
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
	if v33 != v23 {
		goto L33
	} else {
		goto L34
	}
L33:
	;
	v111 = *(*int32)(unsafe.Add(mBase, uint32(v33)))
	v112 = *(*int32)(unsafe.Add(mBase, uint32(v33)+4))
	*(*int32)(unsafe.Add(mBase, uint32(v111)+4)) = v112
	v114 = *(*int32)(unsafe.Add(mBase, uint32(v33)))
	*(*int32)(unsafe.Add(mBase, uint32(v112))) = v114
	v117 = *(*int32)(unsafe.Add(mBase, _c_F_MultiXactIdCreateFromMembers[0]))
	if v117 == int32(0) {
		goto L36
	} else {
		goto L37
	}
L34:
	;
	goto L35
L35:
	;
	v135 = *(*int32)(unsafe.Add(mBase, uint32(v33-int32(8))))
	if v135 != 0 {
		v621 = v135
		goto L7
	} else {
		goto L39
	}
L36:
	;
	v120 = int32(_a_F_MultiXactIdCreateFromMembers_0)
	*(*int32)(unsafe.Add(mBase, _c_F_MultiXactIdCreateFromMembers[1])) = v120
	v124 = v120
	goto L38
L37:
	;
	v124 = v117
	goto L38
L38:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v33))) = int32(_a_F_MultiXactIdCreateFromMembers_0)
	*(*int32)(unsafe.Add(mBase, uint32(v33)+4)) = v124
	*(*int32)(unsafe.Add(mBase, uint32(v124))) = v33
	*(*int32)(unsafe.Add(mBase, _c_F_MultiXactIdCreateFromMembers[0])) = v33
	goto L35
L39:
	;
	goto L8
L40:
	;
	goto L11
L41:
	;
	v155 = int32(0)
	v159 = int32(0)
	goto L44
L42:
	;
	goto L43
L43:
	;
	v192 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_MultiXactIdCreateFromMembers[2])))
	if v192 == int32(1) {
		goto L49
	} else {
		goto L50
	}
L44:
	;
	v169 = *(*int32)(unsafe.Add(mBase, uint32(l1+v159<<(uint(int32(3))%32))+4))
	if v155&int32(1)&base.B2i32(base.Ui32(int32(4)) <= base.Ui32(v169)) != 0 {
		goto L6
	} else {
		goto L46
	}
L45:
	;
	goto L43
L46:
	;
	v177 = v159 + int32(1)
	if v177 != l0 {
		v155 = base.B2i32(base.Ui32(int32(3)) < base.Ui32(v169)) | v155
		v159 = v177
		goto L44
	} else {
		goto L47
	}
L47:
	;
	goto L45
L48:
	;
	if v202 != 0 {
		goto L5
	} else {
		goto L52
	}
L49:
	;
	v197 = *(*int32)(unsafe.Add(mBase, _c_F_MultiXactIdCreateFromMembers[3]))
	v198 = *(*int32)(unsafe.Add(mBase, uint32(v197)+316))
	v200 = base.B2i32(v198 != int32(2))
	*(*uint8)(unsafe.Add(mBase, _c_F_MultiXactIdCreateFromMembers[2])) = uint8(v200)
	v202 = v200
	goto L51
L50:
	;
	v202 = int32(0)
	goto L51
L51:
	;
	goto L48
L52:
	;
	v204 = *(*int32)(unsafe.Add(mBase, _c_F_MultiXactIdCreateFromMembers[4]))
	v208 = F_LWLockAcquire(m, v204+int32(1664), int32(0))
	mBase = m.M
	v209 = m.ExcPending
	if v209 != 0 {
		goto L1
	} else {
		goto L53
	}
L53:
	;
	v211 = *(*int32)(unsafe.Add(mBase, _c_F_MultiXactIdCreateFromMembers[5]))
	v212 = *(*int32)(unsafe.Add(mBase, uint32(v211)))
	if v212 == int32(0) {
		goto L54
	} else {
		goto L55
	}
L54:
	;
	v215 = int32(1)
	*(*int32)(unsafe.Add(mBase, uint32(v211))) = v215
	v218 = v215
	goto L56
L55:
	;
	v218 = v212
	goto L56
L56:
	;
	v219 = *(*int32)(unsafe.Add(mBase, uint32(v211)+28))
	if int32(0) <= v218-v219 {
		goto L57
	} else {
		goto L58
	}
L57:
	;
	v223 = *(*int32)(unsafe.Add(mBase, uint32(v211)+16))
	v224 = *(*int32)(unsafe.Add(mBase, uint32(v211)+40))
	v225 = *(*int32)(unsafe.Add(mBase, uint32(v211)+36))
	v226 = *(*int32)(unsafe.Add(mBase, uint32(v211)+32))
	v228 = *(*int32)(unsafe.Add(mBase, _c_F_MultiXactIdCreateFromMembers[4]))
	F_LWLockRelease(m, v228+int32(1664))
	mBase = m.M
	v232 = m.ExcPending
	if v232 != 0 {
		goto L1
	} else {
		goto L60
	}
L58:
	;
	v331 = v211
	v335 = v218
	goto L59
L59:
	;
	v338 = v335 + int32(1)
	if v338&int32(2047) == int32(0) {
		goto L94
	} else {
		goto L95
	}
L60:
	;
	v234 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_MultiXactIdCreateFromMembers[6])))
	if v234 != int32(1) {
		goto L61
	} else {
		goto L62
	}
L61:
	;
	if v218-v226 < int32(0) {
		goto L76
	} else {
		goto L77
	}
L62:
	;
	if int32(0) <= v218-v225 {
		goto L63
	} else {
		goto L64
	}
L63:
	;
	v240 = F_get_database_name(m, v223)
	mBase = m.M
	v241 = m.ExcPending
	if v241 != 0 {
		goto L1
	} else {
		goto L66
	}
L64:
	;
	goto L65
L65:
	;
	if v218&int32(_a_F_MultiXactIdCreateFromMembers_2) != 0 {
		goto L61
	} else {
		goto L74
	}
L66:
	;
	F_SendPostmasterSignal(m, int32(4))
	mBase = m.M
	v244 = m.ExcPending
	if v244 != 0 {
		goto L1
	} else {
		goto L67
	}
L67:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v248 = m.ExcPending
	if v248 != 0 {
		goto L1
	} else {
		goto L68
	}
L68:
	;
	F_errcode(m, int32(261))
	mBase = m.M
	v251 = m.ExcPending
	if v251 != 0 {
		goto L1
	} else {
		goto L69
	}
L69:
	;
	if v240 != 0 {
		goto L4
	} else {
		goto L70
	}
L70:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v14)+96)) = v223
	F_errmsg(m, int32(_a_F_MultiXactIdCreateFromMembers_3), v14+int32(96))
	mBase = m.M
	v257 = m.ExcPending
	if v257 != 0 {
		goto L1
	} else {
		goto L71
	}
L71:
	;
	F_errhint(m, int32(_a_F_MultiXactIdCreateFromMembers_4), int32(0))
	mBase = m.M
	v261 = m.ExcPending
	if v261 != 0 {
		goto L1
	} else {
		goto L72
	}
L72:
	;
	F_errfinish(m, int32(_a_F_MultiXactIdCreateFromMembers_5), int32(1189), int32(_a_F_MultiXactIdCreateFromMembers_6))
	mBase = m.M
	v266 = m.ExcPending
	if v266 != 0 {
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
	F_SendPostmasterSignal(m, int32(4))
	mBase = m.M
	v271 = m.ExcPending
	if v271 != 0 {
		goto L1
	} else {
		goto L75
	}
L75:
	;
	goto L61
L76:
	;
	v318 = *(*int32)(unsafe.Add(mBase, _c_F_MultiXactIdCreateFromMembers[4]))
	v322 = F_LWLockAcquire(m, v318+int32(1664), int32(0))
	mBase = m.M
	v323 = m.ExcPending
	if v323 != 0 {
		goto L1
	} else {
		goto L90
	}
L77:
	;
	v275 = F_get_database_name(m, v223)
	mBase = m.M
	v276 = m.ExcPending
	if v276 != 0 {
		goto L1
	} else {
		goto L78
	}
L78:
	;
	v279 = F_errstart(m, int32(19), int32(0))
	mBase = m.M
	v280 = m.ExcPending
	if v280 != 0 {
		goto L1
	} else {
		goto L79
	}
L79:
	;
	if v275 != 0 {
		goto L81
	} else {
		goto L82
	}
L80:
	;
	F_errhint(m, int32(_a_F_MultiXactIdCreateFromMembers_4), int32(0))
	mBase = m.M
	v310 = m.ExcPending
	if v310 != 0 {
		goto L1
	} else {
		goto L88
	}
L81:
	;
	if v279 == int32(0) {
		goto L76
	} else {
		goto L84
	}
L82:
	;
	goto L83
L83:
	;
	if v279 == int32(0) {
		goto L76
	} else {
		goto L86
	}
L84:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v14)+80)) = v275
	v284 = v224 - v218
	*(*int32)(unsafe.Add(mBase, uint32(v14)+84)) = v284
	F_errmsg_plural(m, int32(_a_F_MultiXactIdCreateFromMembers_7), int32(_a_F_MultiXactIdCreateFromMembers_8), v284, v14+int32(80))
	mBase = m.M
	v291 = m.ExcPending
	if v291 != 0 {
		goto L1
	} else {
		goto L85
	}
L85:
	;
	v306 = int32(1213)
	goto L80
L86:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v14)+64)) = v223
	v296 = v224 - v218
	*(*int32)(unsafe.Add(mBase, uint32(v14)+68)) = v296
	F_errmsg_plural(m, int32(_a_F_MultiXactIdCreateFromMembers_9), int32(_a_F_MultiXactIdCreateFromMembers_10), v296, v14-int32(-64))
	mBase = m.M
	v303 = m.ExcPending
	if v303 != 0 {
		goto L1
	} else {
		goto L87
	}
L87:
	;
	v306 = int32(1222)
	goto L80
L88:
	;
	F_errfinish(m, int32(_a_F_MultiXactIdCreateFromMembers_5), v306, int32(_a_F_MultiXactIdCreateFromMembers_6))
	mBase = m.M
	v314 = m.ExcPending
	if v314 != 0 {
		goto L1
	} else {
		goto L89
	}
L89:
	;
	goto L76
L90:
	;
	v324 = int32(1)
	v326 = *(*int32)(unsafe.Add(mBase, _c_F_MultiXactIdCreateFromMembers[5]))
	v327 = *(*int32)(unsafe.Add(mBase, uint32(v326)))
	if base.Ui32(v327) <= base.Ui32(v324) {
		goto L91
	} else {
		goto L92
	}
L91:
	;
	v330 = v324
	goto L93
L92:
	;
	v330 = v327
	goto L93
L93:
	;
	v331 = v326
	v335 = v330
	goto L59
L94:
	;
	v344 = *(*int32)(unsafe.Add(mBase, _c_F_MultiXactIdCreateFromMembers[7]))
	v345 = *(*int32)(unsafe.Add(mBase, uint32(v344)+28))
	v347 = int32(base.Ui32(v338) >> (uint(int32(11)) % 32))
	v349 = int32(*(*uint16)(unsafe.Add(mBase, _c_F_MultiXactIdCreateFromMembers[8])))
	v350 = base.I32_rem_u_s(v347, v349)
	v353 = v345 + v350<<(uint(int32(7))%32)
	v355 = F_LWLockAcquire(m, v353, int32(0))
	mBase = m.M
	v356 = m.ExcPending
	if v356 != 0 {
		goto L1
	} else {
		goto L97
	}
L95:
	;
	v377 = v331
	goto L96
L96:
	;
	v381 = *(*int32)(unsafe.Add(mBase, uint32(v377)+4))
	v384 = l0 + base.B2i32(v381 == int32(0))
	v385 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v377)+24)))
	if v385 == int32(1) {
		goto L106
	} else {
		goto L107
	}
L97:
	;
	v358 = base.I64_extend_i32_u(v347)
	v359 = F_SimpleLruZeroPage(m, int32(_a_F_MultiXactIdCreateFromMembers_11), v358)
	mBase = m.M
	v360 = m.ExcPending
	if v360 != 0 {
		goto L1
	} else {
		goto L98
	}
L98:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v14)+128)) = v358
	F_XLogBeginInsert(m)
	mBase = m.M
	v363 = m.ExcPending
	if v363 != 0 {
		goto L1
	} else {
		goto L99
	}
L99:
	;
	F_XLogRegisterData(m, v14+int32(128), int32(8))
	mBase = m.M
	v368 = m.ExcPending
	if v368 != 0 {
		goto L1
	} else {
		goto L100
	}
L100:
	;
	v371 = F_XLogInsert(m, int32(6), int32(0))
	mBase = m.M
	v372 = m.ExcPending
	if v372 != 0 {
		goto L1
	} else {
		goto L101
	}
L101:
	;
	F_LWLockRelease(m, v353)
	mBase = m.M
	v374 = m.ExcPending
	if v374 != 0 {
		goto L1
	} else {
		goto L102
	}
L102:
	;
	v376 = *(*int32)(unsafe.Add(mBase, _c_F_MultiXactIdCreateFromMembers[5]))
	v377 = v376
	goto L96
L103:
	;
	v479 = int32(1)
	if base.Ui32(v381) <= base.Ui32(v479) {
		goto L134
	} else {
		goto L135
	}
L104:
	;
	v433 = v431 + int32(_a_F_MultiXactIdCreateFromMembers_12)
	v435 = v433 + base.B2i32(base.Ui32(v433) < base.Ui32(v381))
	if base.Ui32(v381) < base.Ui32(v429) {
		goto L123
	} else {
		goto L124
	}
L105:
	;
	F_SendPostmasterSignal(m, int32(4))
	mBase = m.M
	v421 = m.ExcPending
	if v421 != 0 {
		goto L1
	} else {
		goto L120
	}
L106:
	;
	v388 = v384 + v381
	v390 = v388 + base.B2i32(base.Ui32(v388) < base.Ui32(v384))
	v391 = *(*int32)(unsafe.Add(mBase, uint32(v377)+44))
	if base.Ui32(v381) < base.Ui32(v391) {
		goto L110
	} else {
		goto L111
	}
L107:
	;
	goto L108
L108:
	;
	v409 = int32(1636)
	v410 = base.I32_div_u_s(v384+v381, v409)
	v412 = base.I32_div_u_s(v381, v409)
	if base.Ui32(v410^v412) < base.Ui32(int32(32)) {
		goto L103
	} else {
		goto L119
	}
L109:
	;
	v397 = *(*int32)(unsafe.Add(mBase, uint32(v377)+20))
	if int32(0) <= v381-v397 {
		v429 = v391
		v431 = v388
		goto L104
	} else {
		goto L117
	}
L110:
	;
	if base.Ui32(v390) < base.Ui32(v381) {
		goto L3
	} else {
		goto L113
	}
L111:
	;
	goto L112
L112:
	;
	if base.Ui32(v381) <= base.Ui32(v390) {
		goto L109
	} else {
		goto L115
	}
L113:
	;
	if base.Ui32(v390) < base.Ui32(v391) {
		goto L109
	} else {
		goto L114
	}
L114:
	;
	goto L3
L115:
	;
	if base.Ui32(v391) <= base.Ui32(v390) {
		goto L3
	} else {
		goto L116
	}
L116:
	;
	goto L109
L117:
	;
	v401 = int32(1636)
	v402 = base.I32_div_u_s(v388, v401)
	v404 = base.I32_div_u_s(v381, v401)
	if base.Ui32(int32(32)) <= base.Ui32(v402^v404) {
		goto L105
	} else {
		goto L118
	}
L118:
	;
	v429 = v391
	v431 = v388
	goto L104
L119:
	;
	goto L105
L120:
	;
	v423 = *(*int32)(unsafe.Add(mBase, _c_F_MultiXactIdCreateFromMembers[5]))
	v424 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v423)+24)))
	if v424 != int32(1) {
		goto L103
	} else {
		goto L121
	}
L121:
	;
	v428 = *(*int32)(unsafe.Add(mBase, uint32(v423)+44))
	v429 = v428
	v431 = v384 + v381
	goto L104
L122:
	;
	v445 = F_errstart(m, int32(19), int32(0))
	mBase = m.M
	v446 = m.ExcPending
	if v446 != 0 {
		goto L1
	} else {
		goto L128
	}
L123:
	;
	if base.B2i32(base.Ui32(v435) < base.Ui32(v381))|base.B2i32(base.Ui32(v429) <= base.Ui32(v435)) != 0 {
		goto L122
	} else {
		goto L126
	}
L124:
	;
	goto L125
L125:
	;
	if base.B2i32(base.Ui32(v435) < base.Ui32(v429))|base.B2i32(base.Ui32(v381) <= base.Ui32(v435)) != 0 {
		goto L103
	} else {
		goto L127
	}
L126:
	;
	goto L103
L127:
	;
	goto L122
L128:
	;
	if v445 == int32(0) {
		goto L103
	} else {
		goto L129
	}
L129:
	;
	F_errcode(m, int32(261))
	mBase = m.M
	v451 = m.ExcPending
	if v451 != 0 {
		goto L1
	} else {
		goto L130
	}
L130:
	;
	v453 = *(*int32)(unsafe.Add(mBase, _c_F_MultiXactIdCreateFromMembers[5]))
	v454 = *(*int32)(unsafe.Add(mBase, uint32(v453)+44))
	v455 = *(*int32)(unsafe.Add(mBase, uint32(v453)+16))
	*(*int32)(unsafe.Add(mBase, uint32(v14)+48)) = v455
	v458 = v454 - v381 + v384
	*(*int32)(unsafe.Add(mBase, uint32(v14)+52)) = v458
	F_errmsg_plural(m, int32(_a_F_MultiXactIdCreateFromMembers_13), int32(_a_F_MultiXactIdCreateFromMembers_14), v458, v14+int32(48))
	mBase = m.M
	v465 = m.ExcPending
	if v465 != 0 {
		goto L1
	} else {
		goto L131
	}
L131:
	;
	F_errhint(m, int32(_a_F_MultiXactIdCreateFromMembers_15), int32(0))
	mBase = m.M
	v469 = m.ExcPending
	if v469 != 0 {
		goto L1
	} else {
		goto L132
	}
L132:
	;
	F_errfinish(m, int32(_a_F_MultiXactIdCreateFromMembers_5), int32(1322), int32(_a_F_MultiXactIdCreateFromMembers_6))
	mBase = m.M
	v474 = m.ExcPending
	if v474 != 0 {
		goto L1
	} else {
		goto L133
	}
L133:
	;
	goto L103
L134:
	;
	v482 = v479
	goto L136
L135:
	;
	v482 = v381
	goto L136
L136:
	;
	if int32(0) < v384 {
		goto L137
	} else {
		goto L138
	}
L137:
	;
	v489 = v384
	v491 = v381
	goto L140
L138:
	;
	goto L139
L139:
	;
	v564 = int32(_a_F_MultiXactIdCreateFromMembers_16)
	v566 = *(*int32)(unsafe.Add(mBase, _c_F_MultiXactIdCreateFromMembers[9]))
	v567 = int32(1)
	*(*int32)(unsafe.Add(mBase, _c_F_MultiXactIdCreateFromMembers[9])) = v566 + v567
	v571 = *(*int32)(unsafe.Add(mBase, _c_F_MultiXactIdCreateFromMembers[5]))
	v572 = *(*int32)(unsafe.Add(mBase, uint32(v571)))
	*(*int32)(unsafe.Add(mBase, uint32(v571))) = v572 + v567
	v576 = *(*int32)(unsafe.Add(mBase, uint32(v571)+4))
	*(*int32)(unsafe.Add(mBase, uint32(v571)+4)) = v576 + v384
	v580 = *(*int32)(unsafe.Add(mBase, _c_F_MultiXactIdCreateFromMembers[4]))
	F_LWLockRelease(m, v580+int32(1664))
	mBase = m.M
	v584 = m.ExcPending
	if v584 != 0 {
		goto L1
	} else {
		goto L155
	}
L140:
	;
	v499 = base.I32_rem_u_s(int32(base.Ui32(v491)>>(uint(int32(2))%32)), int32(409))
	if v499|v491&int32(3) == int32(0) {
		goto L142
	} else {
		goto L143
	}
L141:
	;
	goto L139
L142:
	;
	v506 = *(*int32)(unsafe.Add(mBase, _c_F_MultiXactIdCreateFromMembers[10]))
	v507 = *(*int32)(unsafe.Add(mBase, uint32(v506)+28))
	v509 = base.I32_div_u_s(v491, int32(1636))
	v511 = int32(*(*uint16)(unsafe.Add(mBase, _c_F_MultiXactIdCreateFromMembers[11])))
	v512 = base.I32_rem_u_s(v509, v511)
	v515 = v507 + v512<<(uint(int32(7))%32)
	v517 = F_LWLockAcquire(m, v515, int32(0))
	mBase = m.M
	v518 = m.ExcPending
	if v518 != 0 {
		goto L1
	} else {
		goto L145
	}
L143:
	;
	goto L144
L144:
	;
	v542 = int32(1636)
	v544 = base.I32_rem_u_s(v491, v542)
	if base.Ui32(int32(-1036)) <= base.Ui32(v491) {
		goto L151
	} else {
		goto L152
	}
L145:
	;
	v520 = base.I64_extend_i32_u(v509)
	v521 = F_SimpleLruZeroPage(m, int32(_a_F_MultiXactIdCreateFromMembers_17), v520)
	mBase = m.M
	v522 = m.ExcPending
	if v522 != 0 {
		goto L1
	} else {
		goto L146
	}
L146:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v14)+128)) = v520
	F_XLogBeginInsert(m)
	mBase = m.M
	v525 = m.ExcPending
	if v525 != 0 {
		goto L1
	} else {
		goto L147
	}
L147:
	;
	F_XLogRegisterData(m, v14+int32(128), int32(8))
	mBase = m.M
	v530 = m.ExcPending
	if v530 != 0 {
		goto L1
	} else {
		goto L148
	}
L148:
	;
	v533 = F_XLogInsert(m, int32(6), int32(16))
	mBase = m.M
	v534 = m.ExcPending
	if v534 != 0 {
		goto L1
	} else {
		goto L149
	}
L149:
	;
	F_LWLockRelease(m, v515)
	mBase = m.M
	v536 = m.ExcPending
	if v536 != 0 {
		goto L1
	} else {
		goto L150
	}
L150:
	;
	goto L144
L151:
	;
	v548 = int32(0) - v491
	goto L153
L152:
	;
	v548 = v542 - v544
	goto L153
L153:
	;
	v550 = v489 - v548
	if int32(0) < v550 {
		v489 = v550
		v491 = v548 + v491
		goto L140
	} else {
		goto L154
	}
L154:
	;
	goto L141
L155:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v14)+136)) = l0
	*(*int32)(unsafe.Add(mBase, uint32(v14)+132)) = v482
	*(*int32)(unsafe.Add(mBase, uint32(v14)+128)) = v335
	F_XLogBeginInsert(m)
	mBase = m.M
	v589 = m.ExcPending
	if v589 != 0 {
		goto L1
	} else {
		goto L156
	}
L156:
	;
	F_XLogRegisterData(m, v14+int32(128), int32(12))
	mBase = m.M
	v594 = m.ExcPending
	if v594 != 0 {
		goto L1
	} else {
		goto L157
	}
L157:
	;
	F_XLogRegisterData(m, l1, l0<<(uint(int32(3))%32))
	mBase = m.M
	v598 = m.ExcPending
	if v598 != 0 {
		goto L1
	} else {
		goto L158
	}
L158:
	;
	v601 = F_XLogInsert(m, int32(6), int32(32))
	mBase = m.M
	v602 = m.ExcPending
	if v602 != 0 {
		goto L1
	} else {
		goto L159
	}
L159:
	;
	F_RecordNewMultiXact(m, v335, v482, l0, l1)
	mBase = m.M
	v604 = m.ExcPending
	if v604 != 0 {
		goto L1
	} else {
		goto L160
	}
L160:
	;
	v605 = int32(_a_F_MultiXactIdCreateFromMembers_16)
	v607 = *(*int32)(unsafe.Add(mBase, _c_F_MultiXactIdCreateFromMembers[9]))
	*(*int32)(unsafe.Add(mBase, _c_F_MultiXactIdCreateFromMembers[9])) = v607 - int32(1)
	F_mXactCachePut(m, v335, l0, l1)
	mBase = m.M
	v612 = m.ExcPending
	if v612 != 0 {
		goto L1
	} else {
		goto L161
	}
L161:
	;
	v621 = v335
	goto L7
L162:
	;
	v632 = m.G0
	v634 = v632 - int32(80)
	m.G0 = v634
	v637 = *(*int32)(unsafe.Add(mBase, _c_F_MultiXactIdCreateFromMembers[12]))
	if v637 != 0 {
		goto L164
	} else {
		goto L165
	}
L163:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v14))) = v722
	F_errmsg_internal(m, int32(_a_F_MultiXactIdCreateFromMembers_18), v14)
	mBase = m.M
	v759 = m.ExcPending
	if v759 != 0 {
		goto L1
	} else {
		goto L191
	}
L164:
	;
	F_pfree(m, v637)
	mBase = m.M
	v639 = m.ExcPending
	if v639 != 0 {
		goto L1
	} else {
		goto L167
	}
L165:
	;
	goto L166
L166:
	;
	v641 = v634 - int32(-64)
	F_initStringInfo(m, v641)
	mBase = m.M
	v643 = m.ExcPending
	if v643 != 0 {
		goto L1
	} else {
		goto L168
	}
L167:
	;
	goto L166
L168:
	;
	v644 = *(*int32)(unsafe.Add(mBase, uint32(l1)+4))
	if base.Ui32(v644) < base.Ui32(int32(6)) {
		goto L170
	} else {
		goto L171
	}
L169:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v744 = m.ExcPending
	if v744 != 0 {
		goto L1
	} else {
		goto L188
	}
L170:
	;
	v647 = *(*int32)(unsafe.Add(mBase, uint32(l1)))
	*(*int32)(unsafe.Add(mBase, uint32(v634)+40)) = v647
	*(*int32)(unsafe.Add(mBase, uint32(v634)+32)) = int32(0)
	v653 = *(*int32)(unsafe.Add(mBase, uint32(v644<<(uint(int32(2))%32))+uint32(_c_F_MultiXactIdCreateFromMembers[13])))
	*(*int32)(unsafe.Add(mBase, uint32(v634)+44)) = v653
	*(*int32)(unsafe.Add(mBase, uint32(v634)+36)) = l0
	F_appendStringInfo(m, v641, int32(_a_F_MultiXactIdCreateFromMembers_19), v634+int32(32))
	mBase = m.M
	v660 = m.ExcPending
	if v660 != 0 {
		goto L1
	} else {
		goto L173
	}
L171:
	;
	goto L172
L172:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v729 = m.ExcPending
	if v729 != 0 {
		goto L1
	} else {
		goto L185
	}
L173:
	;
	if int32(2) <= l0 {
		goto L174
	} else {
		goto L175
	}
L174:
	;
	v668 = int32(1)
	goto L177
L175:
	;
	goto L176
L176:
	;
	F_appendStringInfoChar(m, v634-int32(-64), int32(93))
	mBase = m.M
	v710 = m.ExcPending
	if v710 != 0 {
		goto L1
	} else {
		goto L182
	}
L177:
	;
	v677 = l1 + v668<<(uint(int32(3))%32)
	v678 = *(*int32)(unsafe.Add(mBase, uint32(v677)+4))
	if base.Ui32(int32(6)) <= base.Ui32(v678) {
		goto L169
	} else {
		goto L179
	}
L178:
	;
	goto L176
L179:
	;
	v681 = *(*int32)(unsafe.Add(mBase, uint32(v677)))
	*(*int32)(unsafe.Add(mBase, uint32(v634))) = v681
	v685 = *(*int32)(unsafe.Add(mBase, uint32(v678<<(uint(int32(2))%32))+uint32(_c_F_MultiXactIdCreateFromMembers[13])))
	*(*int32)(unsafe.Add(mBase, uint32(v634)+4)) = v685
	F_appendStringInfo(m, v634-int32(-64), int32(_a_F_MultiXactIdCreateFromMembers_20), v634)
	mBase = m.M
	v691 = m.ExcPending
	if v691 != 0 {
		goto L1
	} else {
		goto L180
	}
L180:
	;
	v693 = v668 + int32(1)
	if v693 != l0 {
		v668 = v693
		goto L177
	} else {
		goto L181
	}
L181:
	;
	goto L178
L182:
	;
	v713 = *(*int32)(unsafe.Add(mBase, _c_F_MultiXactIdCreateFromMembers[14]))
	v714 = *(*int32)(unsafe.Add(mBase, uint32(v634)+64))
	v715 = F_MemoryContextStrdup(m, v713, v714)
	mBase = m.M
	v716 = m.ExcPending
	if v716 != 0 {
		goto L1
	} else {
		goto L183
	}
L183:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_MultiXactIdCreateFromMembers[12])) = v715
	v718 = *(*int32)(unsafe.Add(mBase, uint32(v634)+64))
	F_pfree(m, v718)
	mBase = m.M
	v720 = m.ExcPending
	if v720 != 0 {
		goto L1
	} else {
		goto L184
	}
L184:
	;
	v722 = *(*int32)(unsafe.Add(mBase, _c_F_MultiXactIdCreateFromMembers[12]))
	m.G0 = v634 + int32(80)
	goto L163
L185:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v634)+48)) = v644
	F_errmsg_internal(m, int32(_a_F_MultiXactIdCreateFromMembers_21), v634+int32(48))
	mBase = m.M
	v735 = m.ExcPending
	if v735 != 0 {
		goto L1
	} else {
		goto L186
	}
L186:
	;
	F_errfinish(m, int32(_a_F_MultiXactIdCreateFromMembers_5), int32(1834), int32(_a_F_MultiXactIdCreateFromMembers_22))
	mBase = m.M
	v740 = m.ExcPending
	if v740 != 0 {
		goto L1
	} else {
		goto L187
	}
L187:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L188:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v634)+16)) = v678
	F_errmsg_internal(m, int32(_a_F_MultiXactIdCreateFromMembers_21), v634+int32(16))
	mBase = m.M
	v750 = m.ExcPending
	if v750 != 0 {
		goto L1
	} else {
		goto L189
	}
L189:
	;
	F_errfinish(m, int32(_a_F_MultiXactIdCreateFromMembers_5), int32(1834), int32(_a_F_MultiXactIdCreateFromMembers_22))
	mBase = m.M
	v755 = m.ExcPending
	if v755 != 0 {
		goto L1
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
	F_errfinish(m, int32(_a_F_MultiXactIdCreateFromMembers_5), int32(849), int32(_a_F_MultiXactIdCreateFromMembers_23))
	mBase = m.M
	v764 = m.ExcPending
	if v764 != 0 {
		goto L1
	} else {
		goto L192
	}
L192:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L193:
	;
	F_errmsg_internal(m, int32(_a_F_MultiXactIdCreateFromMembers_24), int32(0))
	mBase = m.M
	v772 = m.ExcPending
	if v772 != 0 {
		goto L1
	} else {
		goto L194
	}
L194:
	;
	F_errfinish(m, int32(_a_F_MultiXactIdCreateFromMembers_5), int32(1123), int32(_a_F_MultiXactIdCreateFromMembers_6))
	mBase = m.M
	v777 = m.ExcPending
	if v777 != 0 {
		goto L1
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
	F_errhint(m, int32(_a_F_MultiXactIdCreateFromMembers_4), int32(0))
	mBase = m.M
	v787 = m.ExcPending
	if v787 != 0 {
		goto L1
	} else {
		goto L197
	}
L197:
	;
	F_errfinish(m, int32(_a_F_MultiXactIdCreateFromMembers_5), int32(1182), int32(_a_F_MultiXactIdCreateFromMembers_6))
	mBase = m.M
	v792 = m.ExcPending
	if v792 != 0 {
		goto L1
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
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v799 = m.ExcPending
	if v799 != 0 {
		goto L1
	} else {
		goto L200
	}
L200:
	;
	F_errcode(m, int32(261))
	mBase = m.M
	v802 = m.ExcPending
	if v802 != 0 {
		goto L1
	} else {
		goto L201
	}
L201:
	;
	F_errmsg(m, int32(_a_F_MultiXactIdCreateFromMembers_25), int32(0))
	mBase = m.M
	v806 = m.ExcPending
	if v806 != 0 {
		goto L1
	} else {
		goto L202
	}
L202:
	;
	v808 = *(*int32)(unsafe.Add(mBase, _c_F_MultiXactIdCreateFromMembers[5]))
	v809 = *(*int32)(unsafe.Add(mBase, uint32(v808)+44))
	*(*int32)(unsafe.Add(mBase, uint32(v14)+32)) = v384
	v813 = v809 + (v381 ^ int32(-1))
	*(*int32)(unsafe.Add(mBase, uint32(v14)+36)) = v813
	F_errdetail_plural(m, int32(_a_F_MultiXactIdCreateFromMembers_26), int32(_a_F_MultiXactIdCreateFromMembers_27), v813, v14+int32(32))
	mBase = m.M
	v820 = m.ExcPending
	if v820 != 0 {
		goto L1
	} else {
		goto L203
	}
L203:
	;
	v822 = *(*int32)(unsafe.Add(mBase, _c_F_MultiXactIdCreateFromMembers[5]))
	v823 = *(*int32)(unsafe.Add(mBase, uint32(v822)+16))
	*(*int32)(unsafe.Add(mBase, uint32(v14)+16)) = v823
	F_errhint(m, int32(_a_F_MultiXactIdCreateFromMembers_28), v14+int32(16))
	mBase = m.M
	v829 = m.ExcPending
	if v829 != 0 {
		goto L1
	} else {
		goto L204
	}
L204:
	;
	F_errfinish(m, int32(_a_F_MultiXactIdCreateFromMembers_5), int32(1287), int32(_a_F_MultiXactIdCreateFromMembers_6))
	mBase = m.M
	v834 = m.ExcPending
	if v834 != 0 {
		goto L1
	} else {
		goto L205
	}
L205:
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
	var v20 int32
	_ = v20
	var v23 int32
	_ = v23
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
			v20 = int32(1)
			if base.Ui32(v13) <= base.Ui32(v20) {
				v23 = v20
			} else {
				v23 = v13
			}
			return v23
		}
	}
}
