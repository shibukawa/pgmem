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
	var v255 int32
	_ = v255
	var v257 int32
	_ = v257
	var v261 int32
	_ = v261
	var v264 int32
	_ = v264
	var v270 int32
	_ = v270
	var v274 int32
	_ = v274
	var v279 int32
	_ = v279
	var v284 int32
	_ = v284
	var v288 int32
	_ = v288
	var v295 int32
	_ = v295
	var v297 int32
	_ = v297
	var v301 int32
	_ = v301
	var v302 int32
	_ = v302
	var v305 int32
	_ = v305
	var v306 int32
	_ = v306
	var v310 int32
	_ = v310
	var v317 int32
	_ = v317
	var v322 int32
	_ = v322
	var v329 int32
	_ = v329
	var v332 int32
	_ = v332
	var v336 int32
	_ = v336
	var v340 int32
	_ = v340
	var v344 int32
	_ = v344
	var v348 int32
	_ = v348
	var v349 int32
	_ = v349
	var v350 int32
	_ = v350
	var v352 int32
	_ = v352
	var v353 int32
	_ = v353
	var v356 int32
	_ = v356
	var v357 int32
	_ = v357
	var v361 int32
	_ = v361
	var v364 int32
	_ = v364
	var v370 int32
	_ = v370
	var v371 int32
	_ = v371
	var v373 int32
	_ = v373
	var v375 int32
	_ = v375
	var v376 int32
	_ = v376
	var v379 int32
	_ = v379
	var v381 int32
	_ = v381
	var v382 int32
	_ = v382
	var v384 int64
	_ = v384
	var v385 int32
	_ = v385
	var v386 int32
	_ = v386
	var v389 int32
	_ = v389
	var v394 int32
	_ = v394
	var v397 int64
	_ = v397
	var v398 int32
	_ = v398
	var v400 int32
	_ = v400
	var v402 int32
	_ = v402
	var v403 int32
	_ = v403
	var v407 int32
	_ = v407
	var v410 int32
	_ = v410
	var v411 int32
	_ = v411
	var v414 int32
	_ = v414
	var v416 int32
	_ = v416
	var v417 int32
	_ = v417
	var v423 int32
	_ = v423
	var v427 int32
	_ = v427
	var v428 int32
	_ = v428
	var v430 int32
	_ = v430
	var v435 int32
	_ = v435
	var v436 int32
	_ = v436
	var v438 int32
	_ = v438
	var v447 int32
	_ = v447
	var v451 int32
	_ = v451
	var v458 int32
	_ = v458
	var v460 int32
	_ = v460
	var v462 int32
	_ = v462
	var v463 int32
	_ = v463
	var v467 int32
	_ = v467
	var v468 int32
	_ = v468
	var v470 int32
	_ = v470
	var v472 int32
	_ = v472
	var v474 int32
	_ = v474
	var v484 int32
	_ = v484
	var v485 int32
	_ = v485
	var v490 int32
	_ = v490
	var v492 int32
	_ = v492
	var v493 int32
	_ = v493
	var v494 int32
	_ = v494
	var v497 int32
	_ = v497
	var v504 int32
	_ = v504
	var v508 int32
	_ = v508
	var v513 int32
	_ = v513
	var v518 int32
	_ = v518
	var v521 int32
	_ = v521
	var v528 int32
	_ = v528
	var v530 int32
	_ = v530
	var v538 int32
	_ = v538
	var v545 int32
	_ = v545
	var v546 int32
	_ = v546
	var v548 int32
	_ = v548
	var v550 int32
	_ = v550
	var v551 int32
	_ = v551
	var v554 int32
	_ = v554
	var v556 int32
	_ = v556
	var v557 int32
	_ = v557
	var v559 int64
	_ = v559
	var v560 int32
	_ = v560
	var v561 int32
	_ = v561
	var v564 int32
	_ = v564
	var v569 int32
	_ = v569
	var v572 int64
	_ = v572
	var v573 int32
	_ = v573
	var v575 int32
	_ = v575
	var v581 int32
	_ = v581
	var v583 int32
	_ = v583
	var v587 int32
	_ = v587
	var v589 int32
	_ = v589
	var v603 int32
	_ = v603
	var v605 int32
	_ = v605
	var v606 int32
	_ = v606
	var v610 int32
	_ = v610
	var v611 int32
	_ = v611
	var v615 int32
	_ = v615
	var v619 int32
	_ = v619
	var v623 int32
	_ = v623
	var v628 int32
	_ = v628
	var v633 int32
	_ = v633
	var v637 int32
	_ = v637
	var v640 int64
	_ = v640
	var v641 int32
	_ = v641
	var v643 int32
	_ = v643
	var v644 int32
	_ = v644
	var v646 int32
	_ = v646
	var v651 int32
	_ = v651
	var v660 int32
	_ = v660
	var v670 int32
	_ = v670
	var v671 int32
	_ = v671
	var v673 int32
	_ = v673
	var v676 int32
	_ = v676
	var v678 int32
	_ = v678
	var v680 int32
	_ = v680
	var v682 int32
	_ = v682
	var v683 int32
	_ = v683
	var v686 int32
	_ = v686
	var v692 int32
	_ = v692
	var v699 int32
	_ = v699
	var v707 int32
	_ = v707
	var v716 int32
	_ = v716
	var v717 int32
	_ = v717
	var v720 int32
	_ = v720
	var v724 int32
	_ = v724
	var v730 int32
	_ = v730
	var v732 int32
	_ = v732
	var v749 int32
	_ = v749
	var v752 int32
	_ = v752
	var v753 int32
	_ = v753
	var v754 int32
	_ = v754
	var v755 int32
	_ = v755
	var v757 int32
	_ = v757
	var v759 int32
	_ = v759
	var v761 int32
	_ = v761
	var v768 int32
	_ = v768
	var v774 int32
	_ = v774
	var v779 int32
	_ = v779
	var v783 int32
	_ = v783
	var v789 int32
	_ = v789
	var v794 int32
	_ = v794
	var v798 int32
	_ = v798
	var v803 int32
	_ = v803
	var v807 int32
	_ = v807
	var v811 int32
	_ = v811
	var v816 int32
	_ = v816
	var v822 int32
	_ = v822
	var v826 int32
	_ = v826
	var v831 int32
	_ = v831
	var v834 int32
	_ = v834
	var v838 int32
	_ = v838
	var v845 int32
	_ = v845
	var v847 int32
	_ = v847
	var v851 int32
	_ = v851
	var v854 int32
	_ = v854
	var v858 int32
	_ = v858
	var v860 int32
	_ = v860
	var v861 int32
	_ = v861
	var v865 int32
	_ = v865
	var v872 int32
	_ = v872
	var v874 int32
	_ = v874
	var v875 int32
	_ = v875
	var v881 int32
	_ = v881
	var v886 int32
	_ = v886
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
	v834 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_MultiXactIdCreateFromMembers[1])))
	if v834 == int32(1) {
		goto L209
	} else {
		goto L210
	}
L4:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v14)+112)) = v240
	F_errmsg(m, int32(_a_F_MultiXactIdCreateFromMembers_1), v14+int32(112))
	mBase = m.M
	v822 = m.ExcPending
	if v822 != 0 {
		goto L1
	} else {
		goto L205
	}
L5:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v807 = m.ExcPending
	if v807 != 0 {
		goto L1
	} else {
		goto L202
	}
L6:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v670 = m.ExcPending
	if v670 != 0 {
		goto L1
	} else {
		goto L171
	}
L7:
	;
	m.G0 = v14 + int32(144)
	return v660
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
		v660 = v135
		goto L7
	} else {
		goto L39
	}
L36:
	;
	v120 = int32(_a_F_MultiXactIdCreateFromMembers_0)
	*(*int32)(unsafe.Add(mBase, _c_F_MultiXactIdCreateFromMembers[2])) = v120
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
	v192 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_MultiXactIdCreateFromMembers[3])))
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
	v197 = *(*int32)(unsafe.Add(mBase, _c_F_MultiXactIdCreateFromMembers[4]))
	v198 = *(*int32)(unsafe.Add(mBase, uint32(v197)+316))
	v200 = base.B2i32(v198 != int32(2))
	*(*uint8)(unsafe.Add(mBase, _c_F_MultiXactIdCreateFromMembers[3])) = uint8(v200)
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
	v204 = *(*int32)(unsafe.Add(mBase, _c_F_MultiXactIdCreateFromMembers[5]))
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
	v211 = *(*int32)(unsafe.Add(mBase, _c_F_MultiXactIdCreateFromMembers[6]))
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
	v228 = *(*int32)(unsafe.Add(mBase, _c_F_MultiXactIdCreateFromMembers[5]))
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
	v357 = v211
	v361 = v218
	goto L59
L59:
	;
	v364 = v361 + int32(1)
	if v364&int32(2047) == int32(0) {
		goto L100
	} else {
		goto L101
	}
L60:
	;
	v234 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_MultiXactIdCreateFromMembers[1])))
	if v234 != int32(1) {
		goto L61
	} else {
		goto L62
	}
L61:
	;
	if v218-v226 < int32(0) {
		goto L82
	} else {
		goto L83
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
		goto L77
	}
L66:
	;
	v244 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_MultiXactIdCreateFromMembers[1])))
	if v244 == int32(1) {
		goto L68
	} else {
		goto L69
	}
L67:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v261 = m.ExcPending
	if v261 != 0 {
		goto L1
	} else {
		goto L71
	}
L68:
	;
	v248 = *(*int32)(unsafe.Add(mBase, _c_F_MultiXactIdCreateFromMembers[7]))
	*(*int32)(unsafe.Add(mBase, uint32(v248+int32(16)))) = int32(1)
	v255 = *(*int32)(unsafe.Add(mBase, _c_F_MultiXactIdCreateFromMembers[8]))
	v257 = F_pgmem_kill(m, v255, int32(10))
	mBase = m.M
	goto L70
L69:
	;
	goto L70
L70:
	;
	goto L67
L71:
	;
	F_errcode(m, int32(261))
	mBase = m.M
	v264 = m.ExcPending
	if v264 != 0 {
		goto L1
	} else {
		goto L72
	}
L72:
	;
	if v240 != 0 {
		goto L4
	} else {
		goto L73
	}
L73:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v14)+96)) = v223
	F_errmsg(m, int32(_a_F_MultiXactIdCreateFromMembers_3), v14+int32(96))
	mBase = m.M
	v270 = m.ExcPending
	if v270 != 0 {
		goto L1
	} else {
		goto L74
	}
L74:
	;
	F_errhint(m, int32(_a_F_MultiXactIdCreateFromMembers_4), int32(0))
	mBase = m.M
	v274 = m.ExcPending
	if v274 != 0 {
		goto L1
	} else {
		goto L75
	}
L75:
	;
	F_errfinish(m, int32(_a_F_MultiXactIdCreateFromMembers_5), int32(1189), int32(_a_F_MultiXactIdCreateFromMembers_6))
	mBase = m.M
	v279 = m.ExcPending
	if v279 != 0 {
		goto L1
	} else {
		goto L76
	}
L76:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L77:
	;
	v284 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_MultiXactIdCreateFromMembers[1])))
	if v284 == int32(1) {
		goto L79
	} else {
		goto L80
	}
L78:
	;
	goto L61
L79:
	;
	v288 = *(*int32)(unsafe.Add(mBase, _c_F_MultiXactIdCreateFromMembers[7]))
	*(*int32)(unsafe.Add(mBase, uint32(v288+int32(16)))) = int32(1)
	v295 = *(*int32)(unsafe.Add(mBase, _c_F_MultiXactIdCreateFromMembers[8]))
	v297 = F_pgmem_kill(m, v295, int32(10))
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
	v344 = *(*int32)(unsafe.Add(mBase, _c_F_MultiXactIdCreateFromMembers[5]))
	v348 = F_LWLockAcquire(m, v344+int32(1664), int32(0))
	mBase = m.M
	v349 = m.ExcPending
	if v349 != 0 {
		goto L1
	} else {
		goto L96
	}
L83:
	;
	v301 = F_get_database_name(m, v223)
	mBase = m.M
	v302 = m.ExcPending
	if v302 != 0 {
		goto L1
	} else {
		goto L84
	}
L84:
	;
	v305 = F_errstart(m, int32(19), int32(0))
	mBase = m.M
	v306 = m.ExcPending
	if v306 != 0 {
		goto L1
	} else {
		goto L85
	}
L85:
	;
	if v301 != 0 {
		goto L87
	} else {
		goto L88
	}
L86:
	;
	F_errhint(m, int32(_a_F_MultiXactIdCreateFromMembers_4), int32(0))
	mBase = m.M
	v336 = m.ExcPending
	if v336 != 0 {
		goto L1
	} else {
		goto L94
	}
L87:
	;
	if v305 == int32(0) {
		goto L82
	} else {
		goto L90
	}
L88:
	;
	goto L89
L89:
	;
	if v305 == int32(0) {
		goto L82
	} else {
		goto L92
	}
L90:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v14)+80)) = v301
	v310 = v224 - v218
	*(*int32)(unsafe.Add(mBase, uint32(v14)+84)) = v310
	F_errmsg_plural(m, int32(_a_F_MultiXactIdCreateFromMembers_7), int32(_a_F_MultiXactIdCreateFromMembers_8), v310, v14+int32(80))
	mBase = m.M
	v317 = m.ExcPending
	if v317 != 0 {
		goto L1
	} else {
		goto L91
	}
L91:
	;
	v332 = int32(1213)
	goto L86
L92:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v14)+64)) = v223
	v322 = v224 - v218
	*(*int32)(unsafe.Add(mBase, uint32(v14)+68)) = v322
	F_errmsg_plural(m, int32(_a_F_MultiXactIdCreateFromMembers_9), int32(_a_F_MultiXactIdCreateFromMembers_10), v322, v14-int32(-64))
	mBase = m.M
	v329 = m.ExcPending
	if v329 != 0 {
		goto L1
	} else {
		goto L93
	}
L93:
	;
	v332 = int32(1222)
	goto L86
L94:
	;
	F_errfinish(m, int32(_a_F_MultiXactIdCreateFromMembers_5), v332, int32(_a_F_MultiXactIdCreateFromMembers_6))
	mBase = m.M
	v340 = m.ExcPending
	if v340 != 0 {
		goto L1
	} else {
		goto L95
	}
L95:
	;
	goto L82
L96:
	;
	v350 = int32(1)
	v352 = *(*int32)(unsafe.Add(mBase, _c_F_MultiXactIdCreateFromMembers[6]))
	v353 = *(*int32)(unsafe.Add(mBase, uint32(v352)))
	if base.Ui32(v353) <= base.Ui32(v350) {
		goto L97
	} else {
		goto L98
	}
L97:
	;
	v356 = v350
	goto L99
L98:
	;
	v356 = v353
	goto L99
L99:
	;
	v357 = v352
	v361 = v356
	goto L59
L100:
	;
	v370 = *(*int32)(unsafe.Add(mBase, _c_F_MultiXactIdCreateFromMembers[9]))
	v371 = *(*int32)(unsafe.Add(mBase, uint32(v370)+28))
	v373 = int32(base.Ui32(v364) >> (uint(int32(11)) % 32))
	v375 = int32(*(*uint16)(unsafe.Add(mBase, _c_F_MultiXactIdCreateFromMembers[10])))
	v376 = base.I32_rem_u_s(v373, v375)
	v379 = v371 + v376<<(uint(int32(7))%32)
	v381 = F_LWLockAcquire(m, v379, int32(0))
	mBase = m.M
	v382 = m.ExcPending
	if v382 != 0 {
		goto L1
	} else {
		goto L103
	}
L101:
	;
	v403 = v357
	goto L102
L102:
	;
	v407 = *(*int32)(unsafe.Add(mBase, uint32(v403)+4))
	v410 = l0 + base.B2i32(v407 == int32(0))
	v411 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v403)+24)))
	if v411 == int32(1) {
		goto L112
	} else {
		goto L113
	}
L103:
	;
	v384 = base.I64_extend_i32_u(v373)
	v385 = F_SimpleLruZeroPage(m, int32(_a_F_MultiXactIdCreateFromMembers_11), v384)
	mBase = m.M
	v386 = m.ExcPending
	if v386 != 0 {
		goto L1
	} else {
		goto L104
	}
L104:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v14)+128)) = v384
	F_XLogBeginInsert(m)
	mBase = m.M
	v389 = m.ExcPending
	if v389 != 0 {
		goto L1
	} else {
		goto L105
	}
L105:
	;
	F_XLogRegisterData(m, v14+int32(128), int32(8))
	mBase = m.M
	v394 = m.ExcPending
	if v394 != 0 {
		goto L1
	} else {
		goto L106
	}
L106:
	;
	v397 = F_XLogInsert(m, int32(6), int32(0))
	mBase = m.M
	v398 = m.ExcPending
	if v398 != 0 {
		goto L1
	} else {
		goto L107
	}
L107:
	;
	F_LWLockRelease(m, v379)
	mBase = m.M
	v400 = m.ExcPending
	if v400 != 0 {
		goto L1
	} else {
		goto L108
	}
L108:
	;
	v402 = *(*int32)(unsafe.Add(mBase, _c_F_MultiXactIdCreateFromMembers[6]))
	v403 = v402
	goto L102
L109:
	;
	v518 = int32(1)
	if base.Ui32(v407) <= base.Ui32(v518) {
		goto L143
	} else {
		goto L144
	}
L110:
	;
	v472 = v470 + int32(_a_F_MultiXactIdCreateFromMembers_12)
	v474 = v472 + base.B2i32(base.Ui32(v472) < base.Ui32(v407))
	if base.Ui32(v407) < base.Ui32(v468) {
		goto L132
	} else {
		goto L133
	}
L111:
	;
	v447 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_MultiXactIdCreateFromMembers[1])))
	if v447 == int32(1) {
		goto L127
	} else {
		goto L128
	}
L112:
	;
	v414 = v410 + v407
	v416 = v414 + base.B2i32(base.Ui32(v414) < base.Ui32(v410))
	v417 = *(*int32)(unsafe.Add(mBase, uint32(v403)+44))
	if base.Ui32(v407) < base.Ui32(v417) {
		goto L116
	} else {
		goto L117
	}
L113:
	;
	goto L114
L114:
	;
	v435 = int32(1636)
	v436 = base.I32_div_u_s(v410+v407, v435)
	v438 = base.I32_div_u_s(v407, v435)
	if base.Ui32(v436^v438) < base.Ui32(int32(32)) {
		goto L109
	} else {
		goto L125
	}
L115:
	;
	v423 = *(*int32)(unsafe.Add(mBase, uint32(v403)+20))
	if int32(0) <= v407-v423 {
		v468 = v417
		v470 = v414
		goto L110
	} else {
		goto L123
	}
L116:
	;
	if base.Ui32(v416) < base.Ui32(v407) {
		goto L3
	} else {
		goto L119
	}
L117:
	;
	goto L118
L118:
	;
	if base.Ui32(v407) <= base.Ui32(v416) {
		goto L115
	} else {
		goto L121
	}
L119:
	;
	if base.Ui32(v416) < base.Ui32(v417) {
		goto L115
	} else {
		goto L120
	}
L120:
	;
	goto L3
L121:
	;
	if base.Ui32(v417) <= base.Ui32(v416) {
		goto L3
	} else {
		goto L122
	}
L122:
	;
	goto L115
L123:
	;
	v427 = int32(1636)
	v428 = base.I32_div_u_s(v414, v427)
	v430 = base.I32_div_u_s(v407, v427)
	if base.Ui32(int32(32)) <= base.Ui32(v428^v430) {
		goto L111
	} else {
		goto L124
	}
L124:
	;
	v468 = v417
	v470 = v414
	goto L110
L125:
	;
	goto L111
L126:
	;
	v462 = *(*int32)(unsafe.Add(mBase, _c_F_MultiXactIdCreateFromMembers[6]))
	v463 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v462)+24)))
	if v463 != int32(1) {
		goto L109
	} else {
		goto L130
	}
L127:
	;
	v451 = *(*int32)(unsafe.Add(mBase, _c_F_MultiXactIdCreateFromMembers[7]))
	*(*int32)(unsafe.Add(mBase, uint32(v451+int32(16)))) = int32(1)
	v458 = *(*int32)(unsafe.Add(mBase, _c_F_MultiXactIdCreateFromMembers[8]))
	v460 = F_pgmem_kill(m, v458, int32(10))
	mBase = m.M
	goto L129
L128:
	;
	goto L129
L129:
	;
	goto L126
L130:
	;
	v467 = *(*int32)(unsafe.Add(mBase, uint32(v462)+44))
	v468 = v467
	v470 = v410 + v407
	goto L110
L131:
	;
	v484 = F_errstart(m, int32(19), int32(0))
	mBase = m.M
	v485 = m.ExcPending
	if v485 != 0 {
		goto L1
	} else {
		goto L137
	}
L132:
	;
	if base.B2i32(base.Ui32(v474) < base.Ui32(v407))|base.B2i32(base.Ui32(v468) <= base.Ui32(v474)) != 0 {
		goto L131
	} else {
		goto L135
	}
L133:
	;
	goto L134
L134:
	;
	if base.B2i32(base.Ui32(v474) < base.Ui32(v468))|base.B2i32(base.Ui32(v407) <= base.Ui32(v474)) != 0 {
		goto L109
	} else {
		goto L136
	}
L135:
	;
	goto L109
L136:
	;
	goto L131
L137:
	;
	if v484 == int32(0) {
		goto L109
	} else {
		goto L138
	}
L138:
	;
	F_errcode(m, int32(261))
	mBase = m.M
	v490 = m.ExcPending
	if v490 != 0 {
		goto L1
	} else {
		goto L139
	}
L139:
	;
	v492 = *(*int32)(unsafe.Add(mBase, _c_F_MultiXactIdCreateFromMembers[6]))
	v493 = *(*int32)(unsafe.Add(mBase, uint32(v492)+44))
	v494 = *(*int32)(unsafe.Add(mBase, uint32(v492)+16))
	*(*int32)(unsafe.Add(mBase, uint32(v14)+48)) = v494
	v497 = v493 - v407 + v410
	*(*int32)(unsafe.Add(mBase, uint32(v14)+52)) = v497
	F_errmsg_plural(m, int32(_a_F_MultiXactIdCreateFromMembers_13), int32(_a_F_MultiXactIdCreateFromMembers_14), v497, v14+int32(48))
	mBase = m.M
	v504 = m.ExcPending
	if v504 != 0 {
		goto L1
	} else {
		goto L140
	}
L140:
	;
	F_errhint(m, int32(_a_F_MultiXactIdCreateFromMembers_15), int32(0))
	mBase = m.M
	v508 = m.ExcPending
	if v508 != 0 {
		goto L1
	} else {
		goto L141
	}
L141:
	;
	F_errfinish(m, int32(_a_F_MultiXactIdCreateFromMembers_5), int32(1322), int32(_a_F_MultiXactIdCreateFromMembers_6))
	mBase = m.M
	v513 = m.ExcPending
	if v513 != 0 {
		goto L1
	} else {
		goto L142
	}
L142:
	;
	goto L109
L143:
	;
	v521 = v518
	goto L145
L144:
	;
	v521 = v407
	goto L145
L145:
	;
	if int32(0) < v410 {
		goto L146
	} else {
		goto L147
	}
L146:
	;
	v528 = v410
	v530 = v407
	goto L149
L147:
	;
	goto L148
L148:
	;
	v603 = int32(_a_F_MultiXactIdCreateFromMembers_16)
	v605 = *(*int32)(unsafe.Add(mBase, _c_F_MultiXactIdCreateFromMembers[11]))
	v606 = int32(1)
	*(*int32)(unsafe.Add(mBase, _c_F_MultiXactIdCreateFromMembers[11])) = v605 + v606
	v610 = *(*int32)(unsafe.Add(mBase, _c_F_MultiXactIdCreateFromMembers[6]))
	v611 = *(*int32)(unsafe.Add(mBase, uint32(v610)))
	*(*int32)(unsafe.Add(mBase, uint32(v610))) = v611 + v606
	v615 = *(*int32)(unsafe.Add(mBase, uint32(v610)+4))
	*(*int32)(unsafe.Add(mBase, uint32(v610)+4)) = v615 + v410
	v619 = *(*int32)(unsafe.Add(mBase, _c_F_MultiXactIdCreateFromMembers[5]))
	F_LWLockRelease(m, v619+int32(1664))
	mBase = m.M
	v623 = m.ExcPending
	if v623 != 0 {
		goto L1
	} else {
		goto L164
	}
L149:
	;
	v538 = base.I32_rem_u_s(int32(base.Ui32(v530)>>(uint(int32(2))%32)), int32(409))
	if v538|v530&int32(3) == int32(0) {
		goto L151
	} else {
		goto L152
	}
L150:
	;
	goto L148
L151:
	;
	v545 = *(*int32)(unsafe.Add(mBase, _c_F_MultiXactIdCreateFromMembers[12]))
	v546 = *(*int32)(unsafe.Add(mBase, uint32(v545)+28))
	v548 = base.I32_div_u_s(v530, int32(1636))
	v550 = int32(*(*uint16)(unsafe.Add(mBase, _c_F_MultiXactIdCreateFromMembers[13])))
	v551 = base.I32_rem_u_s(v548, v550)
	v554 = v546 + v551<<(uint(int32(7))%32)
	v556 = F_LWLockAcquire(m, v554, int32(0))
	mBase = m.M
	v557 = m.ExcPending
	if v557 != 0 {
		goto L1
	} else {
		goto L154
	}
L152:
	;
	goto L153
L153:
	;
	v581 = int32(1636)
	v583 = base.I32_rem_u_s(v530, v581)
	if base.Ui32(int32(-1036)) <= base.Ui32(v530) {
		goto L160
	} else {
		goto L161
	}
L154:
	;
	v559 = base.I64_extend_i32_u(v548)
	v560 = F_SimpleLruZeroPage(m, int32(_a_F_MultiXactIdCreateFromMembers_17), v559)
	mBase = m.M
	v561 = m.ExcPending
	if v561 != 0 {
		goto L1
	} else {
		goto L155
	}
L155:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v14)+128)) = v559
	F_XLogBeginInsert(m)
	mBase = m.M
	v564 = m.ExcPending
	if v564 != 0 {
		goto L1
	} else {
		goto L156
	}
L156:
	;
	F_XLogRegisterData(m, v14+int32(128), int32(8))
	mBase = m.M
	v569 = m.ExcPending
	if v569 != 0 {
		goto L1
	} else {
		goto L157
	}
L157:
	;
	v572 = F_XLogInsert(m, int32(6), int32(16))
	mBase = m.M
	v573 = m.ExcPending
	if v573 != 0 {
		goto L1
	} else {
		goto L158
	}
L158:
	;
	F_LWLockRelease(m, v554)
	mBase = m.M
	v575 = m.ExcPending
	if v575 != 0 {
		goto L1
	} else {
		goto L159
	}
L159:
	;
	goto L153
L160:
	;
	v587 = int32(0) - v530
	goto L162
L161:
	;
	v587 = v581 - v583
	goto L162
L162:
	;
	v589 = v528 - v587
	if int32(0) < v589 {
		v528 = v589
		v530 = v587 + v530
		goto L149
	} else {
		goto L163
	}
L163:
	;
	goto L150
L164:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v14)+136)) = l0
	*(*int32)(unsafe.Add(mBase, uint32(v14)+132)) = v521
	*(*int32)(unsafe.Add(mBase, uint32(v14)+128)) = v361
	F_XLogBeginInsert(m)
	mBase = m.M
	v628 = m.ExcPending
	if v628 != 0 {
		goto L1
	} else {
		goto L165
	}
L165:
	;
	F_XLogRegisterData(m, v14+int32(128), int32(12))
	mBase = m.M
	v633 = m.ExcPending
	if v633 != 0 {
		goto L1
	} else {
		goto L166
	}
L166:
	;
	F_XLogRegisterData(m, l1, l0<<(uint(int32(3))%32))
	mBase = m.M
	v637 = m.ExcPending
	if v637 != 0 {
		goto L1
	} else {
		goto L167
	}
L167:
	;
	v640 = F_XLogInsert(m, int32(6), int32(32))
	mBase = m.M
	v641 = m.ExcPending
	if v641 != 0 {
		goto L1
	} else {
		goto L168
	}
L168:
	;
	F_RecordNewMultiXact(m, v361, v521, l0, l1)
	mBase = m.M
	v643 = m.ExcPending
	if v643 != 0 {
		goto L1
	} else {
		goto L169
	}
L169:
	;
	v644 = int32(_a_F_MultiXactIdCreateFromMembers_16)
	v646 = *(*int32)(unsafe.Add(mBase, _c_F_MultiXactIdCreateFromMembers[11]))
	*(*int32)(unsafe.Add(mBase, _c_F_MultiXactIdCreateFromMembers[11])) = v646 - int32(1)
	F_mXactCachePut(m, v361, l0, l1)
	mBase = m.M
	v651 = m.ExcPending
	if v651 != 0 {
		goto L1
	} else {
		goto L170
	}
L170:
	;
	v660 = v361
	goto L7
L171:
	;
	v671 = m.G0
	v673 = v671 - int32(80)
	m.G0 = v673
	v676 = *(*int32)(unsafe.Add(mBase, _c_F_MultiXactIdCreateFromMembers[14]))
	if v676 != 0 {
		goto L173
	} else {
		goto L174
	}
L172:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v14))) = v761
	F_errmsg_internal(m, int32(_a_F_MultiXactIdCreateFromMembers_18), v14)
	mBase = m.M
	v798 = m.ExcPending
	if v798 != 0 {
		goto L1
	} else {
		goto L200
	}
L173:
	;
	F_pfree(m, v676)
	mBase = m.M
	v678 = m.ExcPending
	if v678 != 0 {
		goto L1
	} else {
		goto L176
	}
L174:
	;
	goto L175
L175:
	;
	v680 = v673 - int32(-64)
	F_initStringInfo(m, v680)
	mBase = m.M
	v682 = m.ExcPending
	if v682 != 0 {
		goto L1
	} else {
		goto L177
	}
L176:
	;
	goto L175
L177:
	;
	v683 = *(*int32)(unsafe.Add(mBase, uint32(l1)+4))
	if base.Ui32(v683) < base.Ui32(int32(6)) {
		goto L179
	} else {
		goto L180
	}
L178:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v783 = m.ExcPending
	if v783 != 0 {
		goto L1
	} else {
		goto L197
	}
L179:
	;
	v686 = *(*int32)(unsafe.Add(mBase, uint32(l1)))
	*(*int32)(unsafe.Add(mBase, uint32(v673)+40)) = v686
	*(*int32)(unsafe.Add(mBase, uint32(v673)+32)) = int32(0)
	v692 = *(*int32)(unsafe.Add(mBase, uint32(v683<<(uint(int32(2))%32))+uint32(_c_F_MultiXactIdCreateFromMembers[15])))
	*(*int32)(unsafe.Add(mBase, uint32(v673)+44)) = v692
	*(*int32)(unsafe.Add(mBase, uint32(v673)+36)) = l0
	F_appendStringInfo(m, v680, int32(_a_F_MultiXactIdCreateFromMembers_19), v673+int32(32))
	mBase = m.M
	v699 = m.ExcPending
	if v699 != 0 {
		goto L1
	} else {
		goto L182
	}
L180:
	;
	goto L181
L181:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v768 = m.ExcPending
	if v768 != 0 {
		goto L1
	} else {
		goto L194
	}
L182:
	;
	if int32(2) <= l0 {
		goto L183
	} else {
		goto L184
	}
L183:
	;
	v707 = int32(1)
	goto L186
L184:
	;
	goto L185
L185:
	;
	F_appendStringInfoChar(m, v673-int32(-64), int32(93))
	mBase = m.M
	v749 = m.ExcPending
	if v749 != 0 {
		goto L1
	} else {
		goto L191
	}
L186:
	;
	v716 = l1 + v707<<(uint(int32(3))%32)
	v717 = *(*int32)(unsafe.Add(mBase, uint32(v716)+4))
	if base.Ui32(int32(6)) <= base.Ui32(v717) {
		goto L178
	} else {
		goto L188
	}
L187:
	;
	goto L185
L188:
	;
	v720 = *(*int32)(unsafe.Add(mBase, uint32(v716)))
	*(*int32)(unsafe.Add(mBase, uint32(v673))) = v720
	v724 = *(*int32)(unsafe.Add(mBase, uint32(v717<<(uint(int32(2))%32))+uint32(_c_F_MultiXactIdCreateFromMembers[15])))
	*(*int32)(unsafe.Add(mBase, uint32(v673)+4)) = v724
	F_appendStringInfo(m, v673-int32(-64), int32(_a_F_MultiXactIdCreateFromMembers_20), v673)
	mBase = m.M
	v730 = m.ExcPending
	if v730 != 0 {
		goto L1
	} else {
		goto L189
	}
L189:
	;
	v732 = v707 + int32(1)
	if v732 != l0 {
		v707 = v732
		goto L186
	} else {
		goto L190
	}
L190:
	;
	goto L187
L191:
	;
	v752 = *(*int32)(unsafe.Add(mBase, _c_F_MultiXactIdCreateFromMembers[16]))
	v753 = *(*int32)(unsafe.Add(mBase, uint32(v673)+64))
	v754 = F_MemoryContextStrdup(m, v752, v753)
	mBase = m.M
	v755 = m.ExcPending
	if v755 != 0 {
		goto L1
	} else {
		goto L192
	}
L192:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_MultiXactIdCreateFromMembers[14])) = v754
	v757 = *(*int32)(unsafe.Add(mBase, uint32(v673)+64))
	F_pfree(m, v757)
	mBase = m.M
	v759 = m.ExcPending
	if v759 != 0 {
		goto L1
	} else {
		goto L193
	}
L193:
	;
	v761 = *(*int32)(unsafe.Add(mBase, _c_F_MultiXactIdCreateFromMembers[14]))
	m.G0 = v673 + int32(80)
	goto L172
L194:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v673)+48)) = v683
	F_errmsg_internal(m, int32(_a_F_MultiXactIdCreateFromMembers_21), v673+int32(48))
	mBase = m.M
	v774 = m.ExcPending
	if v774 != 0 {
		goto L1
	} else {
		goto L195
	}
L195:
	;
	F_errfinish(m, int32(_a_F_MultiXactIdCreateFromMembers_5), int32(1834), int32(_a_F_MultiXactIdCreateFromMembers_22))
	mBase = m.M
	v779 = m.ExcPending
	if v779 != 0 {
		goto L1
	} else {
		goto L196
	}
L196:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L197:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v673)+16)) = v717
	F_errmsg_internal(m, int32(_a_F_MultiXactIdCreateFromMembers_21), v673+int32(16))
	mBase = m.M
	v789 = m.ExcPending
	if v789 != 0 {
		goto L1
	} else {
		goto L198
	}
L198:
	;
	F_errfinish(m, int32(_a_F_MultiXactIdCreateFromMembers_5), int32(1834), int32(_a_F_MultiXactIdCreateFromMembers_22))
	mBase = m.M
	v794 = m.ExcPending
	if v794 != 0 {
		goto L1
	} else {
		goto L199
	}
L199:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L200:
	;
	F_errfinish(m, int32(_a_F_MultiXactIdCreateFromMembers_5), int32(849), int32(_a_F_MultiXactIdCreateFromMembers_23))
	mBase = m.M
	v803 = m.ExcPending
	if v803 != 0 {
		goto L1
	} else {
		goto L201
	}
L201:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L202:
	;
	F_errmsg_internal(m, int32(_a_F_MultiXactIdCreateFromMembers_24), int32(0))
	mBase = m.M
	v811 = m.ExcPending
	if v811 != 0 {
		goto L1
	} else {
		goto L203
	}
L203:
	;
	F_errfinish(m, int32(_a_F_MultiXactIdCreateFromMembers_5), int32(1123), int32(_a_F_MultiXactIdCreateFromMembers_6))
	mBase = m.M
	v816 = m.ExcPending
	if v816 != 0 {
		goto L1
	} else {
		goto L204
	}
L204:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L205:
	;
	F_errhint(m, int32(_a_F_MultiXactIdCreateFromMembers_4), int32(0))
	mBase = m.M
	v826 = m.ExcPending
	if v826 != 0 {
		goto L1
	} else {
		goto L206
	}
L206:
	;
	F_errfinish(m, int32(_a_F_MultiXactIdCreateFromMembers_5), int32(1182), int32(_a_F_MultiXactIdCreateFromMembers_6))
	mBase = m.M
	v831 = m.ExcPending
	if v831 != 0 {
		goto L1
	} else {
		goto L207
	}
L207:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L208:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v851 = m.ExcPending
	if v851 != 0 {
		goto L1
	} else {
		goto L212
	}
L209:
	;
	v838 = *(*int32)(unsafe.Add(mBase, _c_F_MultiXactIdCreateFromMembers[7]))
	*(*int32)(unsafe.Add(mBase, uint32(v838+int32(16)))) = int32(1)
	v845 = *(*int32)(unsafe.Add(mBase, _c_F_MultiXactIdCreateFromMembers[8]))
	v847 = F_pgmem_kill(m, v845, int32(10))
	mBase = m.M
	goto L211
L210:
	;
	goto L211
L211:
	;
	goto L208
L212:
	;
	F_errcode(m, int32(261))
	mBase = m.M
	v854 = m.ExcPending
	if v854 != 0 {
		goto L1
	} else {
		goto L213
	}
L213:
	;
	F_errmsg(m, int32(_a_F_MultiXactIdCreateFromMembers_25), int32(0))
	mBase = m.M
	v858 = m.ExcPending
	if v858 != 0 {
		goto L1
	} else {
		goto L214
	}
L214:
	;
	v860 = *(*int32)(unsafe.Add(mBase, _c_F_MultiXactIdCreateFromMembers[6]))
	v861 = *(*int32)(unsafe.Add(mBase, uint32(v860)+44))
	*(*int32)(unsafe.Add(mBase, uint32(v14)+32)) = v410
	v865 = v861 + (v407 ^ int32(-1))
	*(*int32)(unsafe.Add(mBase, uint32(v14)+36)) = v865
	F_errdetail_plural(m, int32(_a_F_MultiXactIdCreateFromMembers_26), int32(_a_F_MultiXactIdCreateFromMembers_27), v865, v14+int32(32))
	mBase = m.M
	v872 = m.ExcPending
	if v872 != 0 {
		goto L1
	} else {
		goto L215
	}
L215:
	;
	v874 = *(*int32)(unsafe.Add(mBase, _c_F_MultiXactIdCreateFromMembers[6]))
	v875 = *(*int32)(unsafe.Add(mBase, uint32(v874)+16))
	*(*int32)(unsafe.Add(mBase, uint32(v14)+16)) = v875
	F_errhint(m, int32(_a_F_MultiXactIdCreateFromMembers_28), v14+int32(16))
	mBase = m.M
	v881 = m.ExcPending
	if v881 != 0 {
		goto L1
	} else {
		goto L216
	}
L216:
	;
	F_errfinish(m, int32(_a_F_MultiXactIdCreateFromMembers_5), int32(1287), int32(_a_F_MultiXactIdCreateFromMembers_6))
	mBase = m.M
	v886 = m.ExcPending
	if v886 != 0 {
		goto L1
	} else {
		goto L217
	}
L217:
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
