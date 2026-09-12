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
	var v47 int32
	_ = v47
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
	var v149 int32
	_ = v149
	var v164 int64
	_ = v164
	var v166 int64
	_ = v166
	var v171 int32
	_ = v171
	var v175 int32
	_ = v175
	var v178 int32
	_ = v178
	var v180 int32
	_ = v180
	var v181 int32
	_ = v181
	var v184 int32
	_ = v184
	var v192 int32
	_ = v192
	var v198 int32
	_ = v198
	var v223 int32
	_ = v223
	var v224 int32
	_ = v224
	var v234 int32
	_ = v234
	var v237 int32
	_ = v237
	var v252 int32
	_ = v252
	var v253 int32
	_ = v253
	var v254 int32
	_ = v254
	var v255 int32
	_ = v255
	var v257 int32
	_ = v257
	var v258 int32
	_ = v258
	var v259 int32
	_ = v259
	var v260 int32
	_ = v260
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
	var v276 int32
	_ = v276
	var v281 int32
	_ = v281
	var v302 int32
	_ = v302
	var v304 int32
	_ = v304
	var v321 int32
	_ = v321
	var v322 int32
	_ = v322
	var v323 int32
	_ = v323
	var v325 int32
	_ = v325
	var v326 int32
	_ = v326
	var v329 int32
	_ = v329
	var v350 int32
	_ = v350
	var v352 int32
	_ = v352
	var v353 int32
	_ = v353
	var v354 int32
	_ = v354
	var v355 int32
	_ = v355
	var v357 int64
	_ = v357
	var v358 int32
	_ = v358
	var v361 int32
	_ = v361
	var v362 int32
	_ = v362
	var v363 int32
	_ = v363
	var v364 int32
	_ = v364
	var v365 int32
	_ = v365
	var v367 int32
	_ = v367
	var v369 int32
	_ = v369
	var v370 int32
	_ = v370
	var v371 int32
	_ = v371
	var v372 int32
	_ = v372
	var v374 int32
	_ = v374
	var v380 int32
	_ = v380
	var v385 int32
	_ = v385
	var v401 int32
	_ = v401
	var v407 int64
	_ = v407
	var v411 int32
	_ = v411
	var v412 int32
	_ = v412
	var v413 int32
	_ = v413
	var v418 int32
	_ = v418
	var v419 int32
	_ = v419
	var v420 int32
	_ = v420
	var v421 int32
	_ = v421
	var v423 int32
	_ = v423
	var v425 int32
	_ = v425
	var v426 int32
	_ = v426
	var v429 int32
	_ = v429
	var v432 int32
	_ = v432
	var v438 int64
	_ = v438
	var v443 int32
	_ = v443
	var v445 int32
	_ = v445
	var v447 int32
	_ = v447
	var v453 int32
	_ = v453
	var v454 int32
	_ = v454
	var v455 int32
	_ = v455
	var v457 int32
	_ = v457
	var v459 int32
	_ = v459
	var v480 int64
	_ = v480
	var v482 int64
	_ = v482
	var v487 int32
	_ = v487
	var v491 int32
	_ = v491
	var v494 int32
	_ = v494
	var v496 int32
	_ = v496
	var v497 int32
	_ = v497
	var v500 int32
	_ = v500
	var v508 int32
	_ = v508
	var v514 int32
	_ = v514
	var v542 int32
	_ = v542
	var v543 int32
	_ = v543
	var v566 int32
	_ = v566
	var v572 int32
	_ = v572
	var v588 int32
	_ = v588
	var v591 int32
	_ = v591
	var v592 int32
	_ = v592
	var v593 int32
	_ = v593
	var v595 int32
	_ = v595
	var v596 int32
	_ = v596
	var v598 int32
	_ = v598
	var v599 int32
	_ = v599
	var v601 int32
	_ = v601
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
	var v612 int32
	_ = v612
	var v614 int32
	_ = v614
	var v619 int32
	_ = v619
	var v638 int32
	_ = v638
	var v640 int32
	_ = v640
	var v642 int32
	_ = v642
	var v646 int32
	_ = v646
	var v647 int32
	_ = v647
	var v648 int32
	_ = v648
	var v651 int32
	_ = v651
	var v654 int32
	_ = v654
	var v656 int32
	_ = v656
	var v660 int32
	_ = v660
	var v661 int32
	_ = v661
	var v662 int32
	_ = v662
	var v663 int32
	_ = v663
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
	v47 = v4
	goto L8
L6:
	;
	goto L7
L7:
	;
	*(*int64)(unsafe.Add(mBase, uint32(l0)+4)) = int64(0)
	v566 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	if v566 == int32(0) {
		goto L1
	} else {
		goto L83
	}
L8:
	;
	v51 = *(*int32)(unsafe.Add(mBase, uint32(v24)+12))
	v55 = *(*int32)(unsafe.Add(mBase, uint32(v51+v47<<(uint(int32(2))%32))))
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
	v542 = v47 + int32(1)
	v543 = *(*int32)(unsafe.Add(mBase, uint32(v24)+4))
	if v542 < v543 {
		v47 = v542
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
	v352 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v58)+304)))
	v353 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
	v354 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v355 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v58)+304)) = uint8(v355)
	v357 = *(*int64)(unsafe.Add(mBase, uint32(v58)+184))
	v358 = *(*int32)(unsafe.Add(mBase, uint32(v57)+152))
	if v358 == v355 {
		goto L53
	} else {
		goto L54
	}
L14:
	;
	v350 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v58)+200)) = uint8(v350)
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
	v223 = v56 & int32(3)
	v224 = int32(0)
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
		v149 = v101
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
	v166 = v164 + base.I64_extend_i32_s(v149)
	*(*int64)(unsafe.Add(mBase, uint32(l2))) = v166
	v171 = *(*int32)(unsafe.Add(mBase, _consts[454]))
	if v171 == int32(0) {
		goto L34
	} else {
		goto L35
	}
L27:
	;
	v107 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v104)+9)))
	if v107 != int32(1) {
		v149 = v101
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
	v149 = v143
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
	v175 = int32(*(*uint8)(unsafe.Add(mBase, _consts[455])))
	if v175 != int32(1) {
		goto L34
	} else {
		goto L36
	}
L36:
	;
	v178 = int32(4438516)
	v180 = *(*int32)(unsafe.Add(mBase, _consts[14]))
	v181 = int32(1)
	*(*int32)(unsafe.Add(mBase, _consts[14])) = v180 + v181
	v184 = *(*int32)(unsafe.Add(mBase, uint32(v171)))
	*(*int32)(unsafe.Add(mBase, uint32(v171))) = v184 + v181
	*(*int64)(unsafe.Add(mBase, uint32(v171+int32(16))+232)) = v166
	v192 = *(*int32)(unsafe.Add(mBase, uint32(v171)))
	*(*int32)(unsafe.Add(mBase, uint32(v171))) = v192 + v181
	v198 = *(*int32)(unsafe.Add(mBase, _consts[14]))
	*(*int32)(unsafe.Add(mBase, _consts[14])) = v198 - v181
	goto L34
L37:
	;
	goto L17
L38:
	;
	v234 = v224
	v237 = int32(0)
	goto L41
L39:
	;
	v281 = v224
	goto L40
L40:
	;
	if v223 == int32(0) {
		goto L14
	} else {
		goto L48
	}
L41:
	;
	v252 = v55 + v234<<(uint(int32(2))%32)
	v253 = *(*int32)(unsafe.Add(mBase, uint32(v252)))
	v254 = *(*int32)(unsafe.Add(mBase, uint32(v253)+8))
	v255 = *(*int32)(unsafe.Add(mBase, uint32(v254)+12))
	m.T0[v255].(func(*base.Module, int32))(m, v253)
	mBase = m.M
	v257 = m.ExcPending
	if v257 != 0 {
		goto L21
	} else {
		goto L43
	}
L42:
	;
	v281 = v274
	goto L40
L43:
	;
	v258 = *(*int32)(unsafe.Add(mBase, uint32(v252)+4))
	v259 = *(*int32)(unsafe.Add(mBase, uint32(v258)+8))
	v260 = *(*int32)(unsafe.Add(mBase, uint32(v259)+12))
	m.T0[v260].(func(*base.Module, int32))(m, v258)
	mBase = m.M
	v262 = m.ExcPending
	if v262 != 0 {
		goto L21
	} else {
		goto L44
	}
L44:
	;
	v263 = *(*int32)(unsafe.Add(mBase, uint32(v252)+8))
	v264 = *(*int32)(unsafe.Add(mBase, uint32(v263)+8))
	v265 = *(*int32)(unsafe.Add(mBase, uint32(v264)+12))
	m.T0[v265].(func(*base.Module, int32))(m, v263)
	mBase = m.M
	v267 = m.ExcPending
	if v267 != 0 {
		goto L21
	} else {
		goto L45
	}
L45:
	;
	v268 = *(*int32)(unsafe.Add(mBase, uint32(v252)+12))
	v269 = *(*int32)(unsafe.Add(mBase, uint32(v268)+8))
	v270 = *(*int32)(unsafe.Add(mBase, uint32(v269)+12))
	m.T0[v270].(func(*base.Module, int32))(m, v268)
	mBase = m.M
	v272 = m.ExcPending
	if v272 != 0 {
		goto L21
	} else {
		goto L46
	}
L46:
	;
	v273 = int32(4)
	v274 = v234 + v273
	v276 = v237 + v273
	if v276 != v56&int32(-4) {
		v234 = v274
		v237 = v276
		goto L41
	} else {
		goto L47
	}
L47:
	;
	goto L42
L48:
	;
	v302 = v281
	v304 = v224
	goto L49
L49:
	;
	v321 = *(*int32)(unsafe.Add(mBase, uint32(v55+v302<<(uint(int32(2))%32))))
	v322 = *(*int32)(unsafe.Add(mBase, uint32(v321)+8))
	v323 = *(*int32)(unsafe.Add(mBase, uint32(v322)+12))
	m.T0[v323].(func(*base.Module, int32))(m, v321)
	mBase = m.M
	v325 = m.ExcPending
	if v325 != 0 {
		goto L21
	} else {
		goto L51
	}
L50:
	;
	goto L14
L51:
	;
	v326 = int32(1)
	v329 = v304 + v326
	if v329 != v223 {
		v302 = v302 + v326
		v304 = v329
		goto L49
	} else {
		goto L52
	}
L52:
	;
	goto L50
L53:
	;
	v361 = F_MakePerTupleExprContext(m, v57)
	mBase = m.M
	v362 = m.ExcPending
	if v362 != 0 {
		goto L21
	} else {
		goto L56
	}
L54:
	;
	v363 = v358
	goto L55
L55:
	;
	v364 = int32(4443856)
	v365 = *(*int32)(unsafe.Add(mBase, _consts[0]))
	v367 = *(*int32)(unsafe.Add(mBase, uint32(v363)+20))
	*(*int32)(unsafe.Add(mBase, _consts[0])) = v367
	v369 = *(*int32)(unsafe.Add(mBase, uint32(v59)+8))
	v370 = *(*int32)(unsafe.Add(mBase, uint32(v55)+4004))
	v371 = *(*int32)(unsafe.Add(mBase, uint32(v369)+188))
	v372 = *(*int32)(unsafe.Add(mBase, uint32(v371)+92))
	m.T0[v372].(func(*base.Module, int32, int32, int32, int32, int32, int32))(m, v369, v55, v56, v354, v353, v370)
	mBase = m.M
	v374 = m.ExcPending
	if v374 != 0 {
		goto L21
	} else {
		goto L57
	}
L56:
	;
	v363 = v361
	goto L55
L57:
	;
	*(*int32)(unsafe.Add(mBase, _consts[0])) = v365
	if int32(0) < v56 {
		goto L58
	} else {
		goto L59
	}
L58:
	;
	v380 = v55 + int32(4016)
	v385 = int32(0)
	goto L61
L59:
	;
	goto L60
L60:
	;
	v480 = *(*int64)(unsafe.Add(mBase, uint32(l2)))
	v482 = v480 + base.I64_extend_i32_s(v56)
	*(*int64)(unsafe.Add(mBase, uint32(l2))) = v482
	v487 = *(*int32)(unsafe.Add(mBase, _consts[454]))
	if v487 == int32(0) {
		goto L79
	} else {
		goto L80
	}
L61:
	;
	v401 = *(*int32)(unsafe.Add(mBase, uint32(v59)+12))
	if int32(0) < v401 {
		goto L64
	} else {
		goto L65
	}
L62:
	;
	goto L60
L63:
	;
	v453 = *(*int32)(unsafe.Add(mBase, uint32(v55+v385<<(uint(int32(2))%32))))
	v454 = *(*int32)(unsafe.Add(mBase, uint32(v453)+8))
	v455 = *(*int32)(unsafe.Add(mBase, uint32(v454)+12))
	m.T0[v455].(func(*base.Module, int32))(m, v453)
	mBase = m.M
	v457 = m.ExcPending
	if v457 != 0 {
		goto L21
	} else {
		goto L76
	}
L64:
	;
	v407 = *(*int64)(unsafe.Add(mBase, uint32(v380+v385<<(uint(int32(3))%32))))
	*(*int64)(unsafe.Add(mBase, uint32(v58)+184)) = v407
	v411 = v55 + v385<<(uint(int32(2))%32)
	v412 = *(*int32)(unsafe.Add(mBase, uint32(v411)))
	v413 = int32(0)
	v418 = F_ExecInsertIndexTuples(m, v59, v412, v57, v413, v413, v413, v413, v413)
	mBase = m.M
	v419 = m.ExcPending
	if v419 != 0 {
		goto L21
	} else {
		goto L67
	}
L65:
	;
	goto L66
L66:
	;
	v426 = *(*int32)(unsafe.Add(mBase, uint32(v59)+52))
	if v426 == int32(0) {
		goto L63
	} else {
		goto L70
	}
L67:
	;
	v420 = *(*int32)(unsafe.Add(mBase, uint32(v411)))
	v421 = *(*int32)(unsafe.Add(mBase, uint32(v58)+260))
	F_ExecARInsertTriggers(m, v57, v59, v420, v418, v421)
	mBase = m.M
	v423 = m.ExcPending
	if v423 != 0 {
		goto L21
	} else {
		goto L68
	}
L68:
	;
	F_list_free(m, v418)
	mBase = m.M
	v425 = m.ExcPending
	if v425 != 0 {
		goto L21
	} else {
		goto L69
	}
L69:
	;
	goto L63
L70:
	;
	v429 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v426)+9)))
	if v429 == int32(0) {
		goto L71
	} else {
		goto L72
	}
L71:
	;
	v432 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v426)+25)))
	if v432 != int32(1) {
		goto L63
	} else {
		goto L74
	}
L72:
	;
	goto L73
L73:
	;
	v438 = *(*int64)(unsafe.Add(mBase, uint32(v380+v385<<(uint(int32(3))%32))))
	*(*int64)(unsafe.Add(mBase, uint32(v58)+184)) = v438
	v443 = *(*int32)(unsafe.Add(mBase, uint32(v55+v385<<(uint(int32(2))%32))))
	v445 = *(*int32)(unsafe.Add(mBase, uint32(v58)+260))
	F_ExecARInsertTriggers(m, v57, v59, v443, int32(0), v445)
	mBase = m.M
	v447 = m.ExcPending
	if v447 != 0 {
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
	v459 = v385 + int32(1)
	if v459 != v56 {
		v385 = v459
		goto L61
	} else {
		goto L77
	}
L77:
	;
	goto L62
L78:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v58)+184)) = v357
	*(*uint8)(unsafe.Add(mBase, uint32(v58)+304)) = uint8(v352)
	goto L10
L79:
	;
	goto L78
L80:
	;
	v491 = int32(*(*uint8)(unsafe.Add(mBase, _consts[455])))
	if v491 != int32(1) {
		goto L79
	} else {
		goto L81
	}
L81:
	;
	v494 = int32(4438516)
	v496 = *(*int32)(unsafe.Add(mBase, _consts[14]))
	v497 = int32(1)
	*(*int32)(unsafe.Add(mBase, _consts[14])) = v496 + v497
	v500 = *(*int32)(unsafe.Add(mBase, uint32(v487)))
	*(*int32)(unsafe.Add(mBase, uint32(v487))) = v500 + v497
	*(*int64)(unsafe.Add(mBase, uint32(v487+int32(16))+232)) = v482
	v508 = *(*int32)(unsafe.Add(mBase, uint32(v487)))
	*(*int32)(unsafe.Add(mBase, uint32(v487))) = v508 + v497
	v514 = *(*int32)(unsafe.Add(mBase, _consts[14]))
	*(*int32)(unsafe.Add(mBase, _consts[14])) = v514 - v497
	goto L79
L82:
	;
	goto L9
L83:
	;
	v572 = v566
	goto L84
L84:
	;
	v588 = *(*int32)(unsafe.Add(mBase, uint32(v572)+4))
	if v588 < int32(33) {
		goto L1
	} else {
		goto L86
	}
L85:
	;
	goto L1
L86:
	;
	v591 = *(*int32)(unsafe.Add(mBase, uint32(v572)+12))
	v592 = *(*int32)(unsafe.Add(mBase, uint32(v591)))
	v593 = *(*int32)(unsafe.Add(mBase, uint32(v592)+4000))
	if l1 == v593 {
		goto L87
	} else {
		goto L88
	}
L87:
	;
	v595 = F_list_delete_first(m, v572)
	mBase = m.M
	v596 = m.ExcPending
	if v596 != 0 {
		goto L21
	} else {
		goto L90
	}
L88:
	;
	v605 = v592
	v606 = v593
	goto L89
L89:
	;
	v607 = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v606)+208)) = v607
	v609 = *(*int32)(unsafe.Add(mBase, uint32(v606)+84))
	if v609 == v607 {
		goto L92
	} else {
		goto L93
	}
L90:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0))) = v595
	v598 = F_lappend(m, v595, v592)
	mBase = m.M
	v599 = m.ExcPending
	if v599 != 0 {
		goto L21
	} else {
		goto L91
	}
L91:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0))) = v598
	v601 = *(*int32)(unsafe.Add(mBase, uint32(v598)+12))
	v602 = *(*int32)(unsafe.Add(mBase, uint32(v601)))
	v603 = *(*int32)(unsafe.Add(mBase, uint32(v602)+4000))
	v605 = v602
	v606 = v603
	goto L89
L92:
	;
	v612 = *(*int32)(unsafe.Add(mBase, uint32(v605)+4004))
	F_FreeBulkInsertState(m, v612)
	mBase = m.M
	v614 = m.ExcPending
	if v614 != 0 {
		goto L21
	} else {
		goto L95
	}
L93:
	;
	goto L94
L94:
	;
	v619 = int32(0)
	goto L96
L95:
	;
	goto L94
L96:
	;
	v638 = *(*int32)(unsafe.Add(mBase, uint32(v605+v619<<(uint(int32(2))%32))))
	if v638 != 0 {
		goto L98
	} else {
		goto L99
	}
L97:
	;
	v646 = *(*int32)(unsafe.Add(mBase, uint32(v606)+84))
	if v646 != 0 {
		goto L103
	} else {
		goto L104
	}
L98:
	;
	F_ExecDropSingleTupleTableSlot(m, v638)
	mBase = m.M
	v640 = m.ExcPending
	if v640 != 0 {
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
	v642 = v619 + int32(1)
	if v642 != int32(1000) {
		v619 = v642
		goto L96
	} else {
		goto L102
	}
L102:
	;
	goto L100
L103:
	;
	F_pfree(m, v605)
	mBase = m.M
	v660 = m.ExcPending
	if v660 != 0 {
		goto L21
	} else {
		goto L108
	}
L104:
	;
	v647 = *(*int32)(unsafe.Add(mBase, uint32(v606)+8))
	v648 = *(*int32)(unsafe.Add(mBase, uint32(v647)+188))
	if v648 == int32(0) {
		goto L103
	} else {
		goto L105
	}
L105:
	;
	v651 = *(*int32)(unsafe.Add(mBase, uint32(v648)+108))
	if v651 == int32(0) {
		goto L103
	} else {
		goto L106
	}
L106:
	;
	v654 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
	m.T0[v651].(func(*base.Module, int32, int32))(m, v647, v654)
	mBase = m.M
	v656 = m.ExcPending
	if v656 != 0 {
		goto L21
	} else {
		goto L107
	}
L107:
	;
	goto L103
L108:
	;
	v661 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v662 = F_list_delete_first(m, v661)
	mBase = m.M
	v663 = m.ExcPending
	if v663 != 0 {
		goto L21
	} else {
		goto L109
	}
L109:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0))) = v662
	if v662 != 0 {
		v572 = v662
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
	var v61 int32
	_ = v61
	var v62 int32
	_ = v62
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
	var v141 int32
	_ = v141
	var v144 int32
	_ = v144
	var v154 int32
	_ = v154
	var v155 int32
	_ = v155
	var v162 int32
	_ = v162
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
	var v190 int32
	_ = v190
	var v191 int32
	_ = v191
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
	var v379 int32
	_ = v379
	var v380 int32
	_ = v380
	var v382 int32
	_ = v382
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
	var v406 int32
	_ = v406
	var v409 int32
	_ = v409
	var v410 int32
	_ = v410
	var v414 int64
	_ = v414
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
	var v445 int32
	_ = v445
	var v446 int32
	_ = v446
	var v447 int32
	_ = v447
	var v448 int64
	_ = v448
	var v449 int32
	_ = v449
	var v450 int32
	_ = v450
	var v453 int32
	_ = v453
	var v460 int32
	_ = v460
	var v461 int32
	_ = v461
	var v463 int32
	_ = v463
	var v466 int32
	_ = v466
	var v476 int32
	_ = v476
	var v477 int32
	_ = v477
	var v478 int32
	_ = v478
	var v479 int32
	_ = v479
	var v482 int32
	_ = v482
	var v487 int32
	_ = v487
	var v490 int32
	_ = v490
	var v494 int32
	_ = v494
	var v499 int32
	_ = v499
	var v503 int32
	_ = v503
	var v506 int32
	_ = v506
	var v512 int32
	_ = v512
	var v517 int32
	_ = v517
	var v521 int32
	_ = v521
	var v524 int32
	_ = v524
	var v530 int32
	_ = v530
	var v535 int32
	_ = v535
	var v541 int32
	_ = v541
	var v543 int32
	_ = v543
	var v553 int32
	_ = v553
	var v555 int32
	_ = v555
	var v561 int32
	_ = v561
	var v569 int32
	_ = v569
	var v578 int32
	_ = v578
	v4 = int32(0)
	v17 = m.G0
	v19 = v17 - int32(48)
	m.G0 = v19
	if l0 == v4 {
		v561 = int32(-1)
		v569 = v4
		goto L2
	} else {
		goto L3
	}
L1:
	;
	m.G0 = v19 + int32(48)
	return v578
L2:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l1))) = v569
	v578 = v561
	goto L1
L3:
	;
	v25 = *(*int32)(unsafe.Add(mBase, _consts[89]))
	if v25 == int32(0) {
		goto L4
	} else {
		goto L5
	}
L4:
	;
	v111 = *(*int32)(unsafe.Add(mBase, _consts[99]))
	v113 = *(*int32)(unsafe.Add(mBase, _consts[100]))
	v117 = *(*int32)(unsafe.Add(mBase, uint32(v111+v113<<(uint(int32(2))%32))))
	if v117 == int32(0) {
		goto L26
	} else {
		goto L27
	}
L5:
	;
	if v25 == int32(4054512) {
		goto L4
	} else {
		goto L6
	}
L6:
	;
	v34 = v25
	goto L7
L7:
	;
	v48 = *(*int32)(unsafe.Add(mBase, uint32(v34-int32(8))))
	if l0 == v48 {
		goto L9
	} else {
		goto L10
	}
L8:
	;
	goto L4
L9:
	;
	v51 = v34 - int32(4)
	v52 = *(*int32)(unsafe.Add(mBase, uint32(v51)))
	v54 = v52 << (uint(int32(3)) % 32)
	v55 = F_palloc(m, v54)
	mBase = m.M
	v58 = m.ExcPending
	if v58 != 0 {
		goto L12
	} else {
		goto L13
	}
L10:
	;
	goto L11
L11:
	;
	v91 = *(*int32)(unsafe.Add(mBase, uint32(v34)+4))
	if v91 != int32(4054512) {
		v34 = v91
		goto L7
	} else {
		goto L25
	}
L12:
	;
	return int32(0)
L13:
	;
	if v54 != 0 {
		goto L15
	} else {
		goto L16
	}
L14:
	;
	v64 = *(*int32)(unsafe.Add(mBase, _consts[89]))
	if v34 != v64 {
		goto L18
	} else {
		goto L19
	}
L15:
	;
	v61 = F__emscripten_memcpy_bulkmem(m, v55, v34+int32(8), v54)
	mBase = m.M
	v62 = v61
	goto L17
L16:
	;
	v62 = v55
	goto L17
L17:
	;
	goto L14
L18:
	;
	v66 = *(*int32)(unsafe.Add(mBase, uint32(v34)))
	v67 = *(*int32)(unsafe.Add(mBase, uint32(v34)+4))
	*(*int32)(unsafe.Add(mBase, uint32(v66)+4)) = v67
	v69 = *(*int32)(unsafe.Add(mBase, uint32(v34)))
	*(*int32)(unsafe.Add(mBase, uint32(v67))) = v69
	v72 = *(*int32)(unsafe.Add(mBase, _consts[89]))
	if v72 == int32(0) {
		goto L21
	} else {
		goto L22
	}
L19:
	;
	goto L20
L20:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l1))) = v62
	v88 = *(*int32)(unsafe.Add(mBase, uint32(v51)))
	if int32(0) <= v88 {
		v578 = v88
		goto L1
	} else {
		goto L24
	}
L21:
	;
	v75 = int32(4054512)
	*(*int32)(unsafe.Add(mBase, _consts[90])) = v75
	v79 = v75
	goto L23
L22:
	;
	v79 = v72
	goto L23
L23:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v34))) = int32(4054512)
	*(*int32)(unsafe.Add(mBase, uint32(v34)+4)) = v79
	*(*int32)(unsafe.Add(mBase, uint32(v79))) = v34
	*(*int32)(unsafe.Add(mBase, _consts[89])) = v34
	goto L20
L24:
	;
	goto L4
L25:
	;
	goto L8
L26:
	;
	v121 = *(*int32)(unsafe.Add(mBase, _consts[44]))
	v125 = F_LWLockAcquire(m, v121+int32(1664), int32(0))
	mBase = m.M
	v126 = m.ExcPending
	if v126 != 0 {
		goto L12
	} else {
		goto L29
	}
L27:
	;
	goto L28
L28:
	;
	if l2 != 0 {
		goto L62
	} else {
		goto L63
	}
L29:
	;
	v127 = int32(1)
	v129 = *(*int32)(unsafe.Add(mBase, _consts[91]))
	v130 = *(*int32)(unsafe.Add(mBase, uint32(v129)))
	if base.Ui32(v130) <= base.Ui32(v127) {
		goto L30
	} else {
		goto L31
	}
L30:
	;
	v133 = v127
	goto L32
L31:
	;
	v133 = v130
	goto L32
L32:
	;
	v135 = *(*int32)(unsafe.Add(mBase, _consts[101]))
	v137 = *(*int32)(unsafe.Add(mBase, _consts[102]))
	v138 = v135 + v137
	if v138 <= int32(0) {
		v218 = v133
		goto L33
	} else {
		goto L34
	}
L33:
	;
	v231 = *(*int32)(unsafe.Add(mBase, _consts[99]))
	v233 = *(*int32)(unsafe.Add(mBase, _consts[100]))
	*(*int32)(unsafe.Add(mBase, uint32(v231+v233<<(uint(int32(2))%32)))) = v218
	v239 = *(*int32)(unsafe.Add(mBase, _consts[44]))
	F_LWLockRelease(m, v239+int32(1664))
	mBase = m.M
	v243 = m.ExcPending
	if v243 != 0 {
		goto L12
	} else {
		goto L61
	}
L34:
	;
	v141 = int32(1)
	v144 = *(*int32)(unsafe.Add(mBase, _consts[103]))
	if v138 == v141 {
		goto L36
	} else {
		goto L37
	}
L35:
	;
	if v138&v141 == int32(0) {
		v218 = v191
		goto L33
	} else {
		goto L54
	}
L36:
	;
	v190 = int32(0)
	v191 = v133
	goto L35
L37:
	;
	goto L38
L38:
	;
	v154 = int32(0)
	v155 = v133
	v162 = v4
	goto L39
L39:
	;
	v169 = v144 + v154<<(uint(int32(2))%32)
	v170 = *(*int32)(unsafe.Add(mBase, uint32(v169)+4))
	v171 = *(*int32)(unsafe.Add(mBase, uint32(v169)))
	if v171-v155 < int32(0) {
		goto L41
	} else {
		goto L42
	}
L40:
	;
	v190 = v183
	v191 = v181
	goto L35
L41:
	;
	v175 = v171
	goto L43
L42:
	;
	v175 = v155
	goto L43
L43:
	;
	if v171 != 0 {
		goto L44
	} else {
		goto L45
	}
L44:
	;
	v176 = v175
	goto L46
L45:
	;
	v176 = v155
	goto L46
L46:
	;
	if v170-v176 < int32(0) {
		goto L47
	} else {
		goto L48
	}
L47:
	;
	v180 = v170
	goto L49
L48:
	;
	v180 = v176
	goto L49
L49:
	;
	if v170 != 0 {
		goto L50
	} else {
		goto L51
	}
L50:
	;
	v181 = v180
	goto L52
L51:
	;
	v181 = v176
	goto L52
L52:
	;
	v182 = int32(2)
	v183 = v154 + v182
	v185 = v162 + v182
	if v185 != v138&int32(2147483646) {
		v154 = v183
		v155 = v181
		v162 = v185
		goto L39
	} else {
		goto L53
	}
L53:
	;
	goto L40
L54:
	;
	v208 = *(*int32)(unsafe.Add(mBase, uint32(v144+v190<<(uint(int32(2))%32))))
	if v208-v191 < int32(0) {
		goto L55
	} else {
		goto L56
	}
L55:
	;
	v212 = v208
	goto L57
L56:
	;
	v212 = v191
	goto L57
L57:
	;
	if v208 != 0 {
		goto L58
	} else {
		goto L59
	}
L58:
	;
	v213 = v212
	goto L60
L59:
	;
	v213 = v191
	goto L60
L60:
	;
	v218 = v213
	goto L33
L61:
	;
	goto L28
L62:
	;
	v262 = *(*int32)(unsafe.Add(mBase, _consts[99]))
	v264 = *(*int32)(unsafe.Add(mBase, _consts[100]))
	v268 = *(*int32)(unsafe.Add(mBase, uint32(v262+v264<<(uint(int32(2))%32))))
	if l0-v268 < int32(0) {
		v561 = int32(-1)
		v569 = v4
		goto L2
	} else {
		goto L65
	}
L63:
	;
	goto L64
L64:
	;
	v274 = *(*int32)(unsafe.Add(mBase, _consts[44]))
	v278 = F_LWLockAcquire(m, v274+int32(1664), int32(1))
	mBase = m.M
	v279 = m.ExcPending
	if v279 != 0 {
		goto L12
	} else {
		goto L66
	}
L65:
	;
	goto L64
L66:
	;
	v281 = *(*int32)(unsafe.Add(mBase, _consts[91]))
	v282 = *(*int32)(unsafe.Add(mBase, uint32(v281)+4))
	v283 = *(*int32)(unsafe.Add(mBase, uint32(v281)))
	v284 = *(*int32)(unsafe.Add(mBase, uint32(v281)+12))
	v286 = *(*int32)(unsafe.Add(mBase, _consts[44]))
	F_LWLockRelease(m, v286+int32(1664))
	mBase = m.M
	v290 = m.ExcPending
	if v290 != 0 {
		goto L12
	} else {
		goto L67
	}
L67:
	;
	if int32(0) <= l0-v284 {
		goto L71
	} else {
		goto L72
	}
L68:
	;
	F_LWLockRelease(m, v543)
	mBase = m.M
	v553 = m.ExcPending
	if v553 != 0 {
		goto L12
	} else {
		goto L131
	}
L69:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v521 = m.ExcPending
	if v521 != 0 {
		goto L12
	} else {
		goto L127
	}
L70:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v503 = m.ExcPending
	if v503 != 0 {
		goto L12
	} else {
		goto L123
	}
L71:
	;
	if int32(0) <= l0-v283 {
		goto L70
	} else {
		goto L74
	}
L72:
	;
	goto L73
L73:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v487 = m.ExcPending
	if v487 != 0 {
		goto L12
	} else {
		goto L119
	}
L74:
	;
	v297 = int32(0)
	v299 = *(*int32)(unsafe.Add(mBase, _consts[93]))
	v300 = *(*int32)(unsafe.Add(mBase, uint32(v299)+28))
	v302 = int32(base.Ui32(l0) >> (uint(int32(11)) % 32))
	v304 = int32(*(*uint16)(unsafe.Add(mBase, _consts[94])))
	v305 = base.I32_rem_u_s(v302, v304)
	v308 = v300 + v305<<(uint(int32(7))%32)
	v310 = F_LWLockAcquire(m, v308, v297)
	mBase = m.M
	v311 = m.ExcPending
	if v311 != 0 {
		goto L12
	} else {
		goto L75
	}
L75:
	;
	v315 = F_SimpleLruReadPage(m, int32(4338564), base.I64_extend_i32_u(v302), int32(1), l0)
	mBase = m.M
	v316 = m.ExcPending
	if v316 != 0 {
		goto L12
	} else {
		goto L76
	}
L76:
	;
	v318 = *(*int32)(unsafe.Add(mBase, _consts[93]))
	v319 = *(*int32)(unsafe.Add(mBase, uint32(v318)+4))
	v320 = int32(2)
	v323 = *(*int32)(unsafe.Add(mBase, uint32(v319+v315<<(uint(v320)%32))))
	v329 = *(*int32)(unsafe.Add(mBase, uint32(v323+l0&int32(2047)<<(uint(v320)%32))))
	v331 = l0 + int32(1)
	if v331 == v283 {
		goto L78
	} else {
		goto L79
	}
L77:
	;
	F_LWLockRelease(m, v379)
	mBase = m.M
	v385 = m.ExcPending
	if v385 != 0 {
		goto L12
	} else {
		goto L96
	}
L78:
	;
	v379 = v308
	v380 = v282
	v382 = v315
	goto L77
L79:
	;
	goto L80
L80:
	;
	v333 = int32(1)
	if base.Ui32(v331) <= base.Ui32(v333) {
		goto L81
	} else {
		goto L82
	}
L81:
	;
	v336 = v333
	goto L83
L82:
	;
	v336 = v331
	goto L83
L83:
	;
	v340 = int32(base.Ui32(v336) >> (uint(int32(11)) % 32))
	if v302 == v340 {
		goto L85
	} else {
		goto L86
	}
L84:
	;
	v375 = *(*int32)(unsafe.Add(mBase, uint32(v369+v336&int32(2047)<<(uint(int32(2))%32))))
	if v375 == int32(0) {
		goto L69
	} else {
		goto L95
	}
L85:
	;
	v368 = v308
	v369 = v323
	v370 = v315
	goto L84
L86:
	;
	goto L87
L87:
	;
	v343 = *(*int32)(unsafe.Add(mBase, uint32(v318)+28))
	v345 = int32(*(*uint16)(unsafe.Add(mBase, _consts[94])))
	v346 = base.I32_rem_u_s(v340, v345)
	v349 = v343 + v346<<(uint(int32(7))%32)
	if v308 == v349 {
		goto L89
	} else {
		goto L90
	}
L88:
	;
	v359 = F_SimpleLruReadPage(m, int32(4338564), base.I64_extend_i32_u(v340), int32(1), v336)
	mBase = m.M
	v360 = m.ExcPending
	if v360 != 0 {
		goto L12
	} else {
		goto L94
	}
L89:
	;
	v356 = v308
	goto L88
L90:
	;
	goto L91
L91:
	;
	F_LWLockRelease(m, v308)
	mBase = m.M
	v352 = m.ExcPending
	if v352 != 0 {
		goto L12
	} else {
		goto L92
	}
L92:
	;
	v354 = F_LWLockAcquire(m, v349, int32(0))
	mBase = m.M
	v355 = m.ExcPending
	if v355 != 0 {
		goto L12
	} else {
		goto L93
	}
L93:
	;
	v356 = v349
	goto L88
L94:
	;
	v362 = *(*int32)(unsafe.Add(mBase, _consts[93]))
	v363 = *(*int32)(unsafe.Add(mBase, uint32(v362)+4))
	v367 = *(*int32)(unsafe.Add(mBase, uint32(v363+v359<<(uint(int32(2))%32))))
	v368 = v356
	v369 = v367
	v370 = v359
	goto L84
L95:
	;
	v379 = v368
	v380 = v375
	v382 = v370
	goto L77
L96:
	;
	v386 = v380 - v329
	v389 = F_palloc(m, v386<<(uint(int32(3))%32))
	mBase = m.M
	v390 = m.ExcPending
	if v390 != 0 {
		goto L12
	} else {
		goto L97
	}
L97:
	;
	if v386 <= int32(0) {
		goto L98
	} else {
		goto L99
	}
L98:
	;
	v541 = v297
	v543 = int32(0)
	goto L68
L99:
	;
	goto L100
L100:
	;
	v395 = *(*int32)(unsafe.Add(mBase, _consts[95]))
	v397 = int32(0)
	v402 = v395
	v403 = v329
	v404 = v297
	v406 = v397
	v409 = v382
	v410 = v397
	v414 = int64(-1)
	goto L101
L101:
	;
	v418 = base.I32_rem_u_s(int32(base.Ui32(v403)>>(uint(int32(2))%32)), int32(409))
	v420 = v418 * int32(20)
	v422 = base.I32_div_u_s(v403, int32(1636))
	v423 = base.I64_extend_i32_u(v422)
	if v423 != v414 {
		goto L103
	} else {
		goto L104
	}
L102:
	;
	v541 = v478
	v543 = v446
	goto L68
L103:
	;
	v425 = *(*int32)(unsafe.Add(mBase, uint32(v402)+28))
	v427 = int32(*(*uint16)(unsafe.Add(mBase, _consts[96])))
	v428 = base.I32_rem_u_s(v422, v427)
	v431 = v425 + v428<<(uint(int32(7))%32)
	if v406 != v431 {
		goto L106
	} else {
		goto L107
	}
L104:
	;
	v445 = v402
	v446 = v406
	v447 = v409
	v448 = v414
	goto L105
L105:
	;
	v449 = *(*int32)(unsafe.Add(mBase, uint32(v445)+4))
	v450 = int32(2)
	v453 = *(*int32)(unsafe.Add(mBase, uint32(v449+v447<<(uint(v450)%32))))
	v460 = *(*int32)(unsafe.Add(mBase, uint32(v420+(v453+v403<<(uint(v450)%32)&int32(12)))+4))
	if v460 != 0 {
		goto L115
	} else {
		goto L116
	}
L106:
	;
	if v406 != 0 {
		goto L109
	} else {
		goto L110
	}
L107:
	;
	v438 = v406
	goto L108
L108:
	;
	v441 = F_SimpleLruReadPage(m, int32(4338644), v423, int32(1), l0)
	mBase = m.M
	v442 = m.ExcPending
	if v442 != 0 {
		goto L12
	} else {
		goto L114
	}
L109:
	;
	F_LWLockRelease(m, v406)
	mBase = m.M
	v434 = m.ExcPending
	if v434 != 0 {
		goto L12
	} else {
		goto L112
	}
L110:
	;
	goto L111
L111:
	;
	v436 = F_LWLockAcquire(m, v431, int32(0))
	mBase = m.M
	v437 = m.ExcPending
	if v437 != 0 {
		goto L12
	} else {
		goto L113
	}
L112:
	;
	goto L111
L113:
	;
	v438 = v431
	goto L108
L114:
	;
	v444 = *(*int32)(unsafe.Add(mBase, _consts[95]))
	v445 = v444
	v446 = v438
	v447 = v441
	v448 = v423
	goto L105
L115:
	;
	v461 = int32(3)
	v463 = v389 + v404<<(uint(v461)%32)
	*(*int32)(unsafe.Add(mBase, uint32(v463))) = v460
	v466 = *(*int32)(unsafe.Add(mBase, uint32(v453+v420)))
	*(*int32)(unsafe.Add(mBase, uint32(v463)+4)) = int32(base.Ui32(v466)>>(uint(v403<<(uint(v461)%32))%32)) & int32(255)
	v476 = *(*int32)(unsafe.Add(mBase, _consts[95]))
	v477 = v476
	v478 = v404 + int32(1)
	goto L117
L116:
	;
	v477 = v445
	v478 = v404
	goto L117
L117:
	;
	v479 = int32(1)
	v482 = v410 + v479
	if v386 != v482 {
		v402 = v477
		v403 = v403 + v479
		v404 = v478
		v406 = v446
		v409 = v447
		v410 = v482
		v414 = v448
		goto L101
	} else {
		goto L118
	}
L118:
	;
	goto L102
L119:
	;
	F_errcode(m, int32(2600))
	mBase = m.M
	v490 = m.ExcPending
	if v490 != 0 {
		goto L12
	} else {
		goto L120
	}
L120:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v19))) = l0
	F_errmsg(m, int32(404431), v19)
	mBase = m.M
	v494 = m.ExcPending
	if v494 != 0 {
		goto L12
	} else {
		goto L121
	}
L121:
	;
	F_errfinish(m, int32(472009), int32(1462), int32(127255))
	mBase = m.M
	v499 = m.ExcPending
	if v499 != 0 {
		goto L12
	} else {
		goto L122
	}
L122:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L123:
	;
	F_errcode(m, int32(2600))
	mBase = m.M
	v506 = m.ExcPending
	if v506 != 0 {
		goto L12
	} else {
		goto L124
	}
L124:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v19)+32)) = l0
	F_errmsg(m, int32(404490), v19+int32(32))
	mBase = m.M
	v512 = m.ExcPending
	if v512 != 0 {
		goto L12
	} else {
		goto L125
	}
L125:
	;
	F_errfinish(m, int32(472009), int32(1468), int32(127255))
	mBase = m.M
	v517 = m.ExcPending
	if v517 != 0 {
		goto L12
	} else {
		goto L126
	}
L126:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L127:
	;
	F_errcode(m, int32(16779816))
	mBase = m.M
	v524 = m.ExcPending
	if v524 != 0 {
		goto L12
	} else {
		goto L128
	}
L128:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v19)+16)) = l0
	F_errmsg(m, int32(98725), v19+int32(16))
	mBase = m.M
	v530 = m.ExcPending
	if v530 != 0 {
		goto L12
	} else {
		goto L129
	}
L129:
	;
	F_errfinish(m, int32(472009), int32(1561), int32(127255))
	mBase = m.M
	v535 = m.ExcPending
	if v535 != 0 {
		goto L12
	} else {
		goto L130
	}
L130:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L131:
	;
	F_mXactCachePut(m, l0, v541, v389)
	mBase = m.M
	v555 = m.ExcPending
	if v555 != 0 {
		goto L12
	} else {
		goto L132
	}
L132:
	;
	v561 = v541
	v569 = v389
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
	var v29 int32
	_ = v29
	var v32 int32
	_ = v32
	var v43 int32
	_ = v43
	var v46 int32
	_ = v46
	var v54 int32
	_ = v54
	var v55 int32
	_ = v55
	var v56 int32
	_ = v56
	var v59 int32
	_ = v59
	var v60 int32
	_ = v60
	var v62 int32
	_ = v62
	var v63 int32
	_ = v63
	var v65 int32
	_ = v65
	var v67 int32
	_ = v67
	var v70 int32
	_ = v70
	var v71 int32
	_ = v71
	var v72 int32
	_ = v72
	var v77 int32
	_ = v77
	var v78 int32
	_ = v78
	var v79 int32
	_ = v79
	var v82 int32
	_ = v82
	var v83 int32
	_ = v83
	var v84 int32
	_ = v84
	var v87 int32
	_ = v87
	var v88 int32
	_ = v88
	var v90 int32
	_ = v90
	var v95 int32
	_ = v95
	var v108 int32
	_ = v108
	var v112 int32
	_ = v112
	var v113 int32
	_ = v113
	var v115 int32
	_ = v115
	var v118 int32
	_ = v118
	var v121 int32
	_ = v121
	var v125 int32
	_ = v125
	var v133 int32
	_ = v133
	var v134 int32
	_ = v134
	var v153 int32
	_ = v153
	var v155 int32
	_ = v155
	var v167 int32
	_ = v167
	var v175 int32
	_ = v175
	var v190 int32
	_ = v190
	var v195 int32
	_ = v195
	var v196 int32
	_ = v196
	var v198 int32
	_ = v198
	var v200 int32
	_ = v200
	var v202 int32
	_ = v202
	var v206 int32
	_ = v206
	var v207 int32
	_ = v207
	var v209 int32
	_ = v209
	var v210 int32
	_ = v210
	var v213 int32
	_ = v213
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
	var v249 int32
	_ = v249
	var v255 int32
	_ = v255
	var v259 int32
	_ = v259
	var v264 int32
	_ = v264
	var v269 int32
	_ = v269
	var v273 int32
	_ = v273
	var v274 int32
	_ = v274
	var v277 int32
	_ = v277
	var v278 int32
	_ = v278
	var v282 int32
	_ = v282
	var v289 int32
	_ = v289
	var v294 int32
	_ = v294
	var v301 int32
	_ = v301
	var v304 int32
	_ = v304
	var v308 int32
	_ = v308
	var v312 int32
	_ = v312
	var v316 int32
	_ = v316
	var v320 int32
	_ = v320
	var v321 int32
	_ = v321
	var v322 int32
	_ = v322
	var v324 int32
	_ = v324
	var v325 int32
	_ = v325
	var v328 int32
	_ = v328
	var v331 int32
	_ = v331
	var v333 int32
	_ = v333
	var v336 int32
	_ = v336
	var v342 int32
	_ = v342
	var v343 int32
	_ = v343
	var v345 int32
	_ = v345
	var v347 int32
	_ = v347
	var v348 int32
	_ = v348
	var v351 int32
	_ = v351
	var v353 int32
	_ = v353
	var v354 int32
	_ = v354
	var v356 int64
	_ = v356
	var v357 int32
	_ = v357
	var v358 int32
	_ = v358
	var v361 int32
	_ = v361
	var v366 int32
	_ = v366
	var v369 int64
	_ = v369
	var v370 int32
	_ = v370
	var v372 int32
	_ = v372
	var v374 int32
	_ = v374
	var v376 int32
	_ = v376
	var v378 int32
	_ = v378
	var v381 int32
	_ = v381
	var v382 int32
	_ = v382
	var v385 int32
	_ = v385
	var v387 int32
	_ = v387
	var v388 int32
	_ = v388
	var v394 int32
	_ = v394
	var v398 int32
	_ = v398
	var v399 int32
	_ = v399
	var v401 int32
	_ = v401
	var v406 int32
	_ = v406
	var v407 int32
	_ = v407
	var v409 int32
	_ = v409
	var v418 int32
	_ = v418
	var v420 int32
	_ = v420
	var v421 int32
	_ = v421
	var v427 int32
	_ = v427
	var v432 int32
	_ = v432
	var v434 int32
	_ = v434
	var v435 int32
	_ = v435
	var v443 int32
	_ = v443
	var v444 int32
	_ = v444
	var v449 int32
	_ = v449
	var v451 int32
	_ = v451
	var v452 int32
	_ = v452
	var v453 int32
	_ = v453
	var v456 int32
	_ = v456
	var v463 int32
	_ = v463
	var v467 int32
	_ = v467
	var v472 int32
	_ = v472
	var v477 int32
	_ = v477
	var v480 int32
	_ = v480
	var v485 int32
	_ = v485
	var v486 int32
	_ = v486
	var v497 int32
	_ = v497
	var v504 int32
	_ = v504
	var v505 int32
	_ = v505
	var v507 int32
	_ = v507
	var v509 int32
	_ = v509
	var v510 int32
	_ = v510
	var v513 int32
	_ = v513
	var v515 int32
	_ = v515
	var v516 int32
	_ = v516
	var v518 int64
	_ = v518
	var v519 int32
	_ = v519
	var v520 int32
	_ = v520
	var v523 int32
	_ = v523
	var v528 int32
	_ = v528
	var v531 int64
	_ = v531
	var v532 int32
	_ = v532
	var v534 int32
	_ = v534
	var v542 int32
	_ = v542
	var v544 int32
	_ = v544
	var v546 int32
	_ = v546
	var v548 int32
	_ = v548
	var v562 int32
	_ = v562
	var v564 int32
	_ = v564
	var v565 int32
	_ = v565
	var v569 int32
	_ = v569
	var v570 int32
	_ = v570
	var v574 int32
	_ = v574
	var v578 int32
	_ = v578
	var v582 int32
	_ = v582
	var v587 int32
	_ = v587
	var v592 int32
	_ = v592
	var v596 int32
	_ = v596
	var v599 int64
	_ = v599
	var v600 int32
	_ = v600
	var v602 int32
	_ = v602
	var v603 int32
	_ = v603
	var v605 int32
	_ = v605
	var v610 int32
	_ = v610
	var v618 int32
	_ = v618
	var v629 int32
	_ = v629
	var v630 int32
	_ = v630
	var v632 int32
	_ = v632
	var v635 int32
	_ = v635
	var v637 int32
	_ = v637
	var v641 int32
	_ = v641
	var v642 int32
	_ = v642
	var v645 int32
	_ = v645
	var v653 int32
	_ = v653
	var v662 int32
	_ = v662
	var v669 int32
	_ = v669
	var v679 int32
	_ = v679
	var v680 int32
	_ = v680
	var v683 int32
	_ = v683
	var v689 int32
	_ = v689
	var v695 int32
	_ = v695
	var v697 int32
	_ = v697
	var v714 int32
	_ = v714
	var v717 int32
	_ = v717
	var v718 int32
	_ = v718
	var v719 int32
	_ = v719
	var v720 int32
	_ = v720
	var v722 int32
	_ = v722
	var v724 int32
	_ = v724
	var v726 int32
	_ = v726
	var v733 int32
	_ = v733
	var v739 int32
	_ = v739
	var v744 int32
	_ = v744
	var v748 int32
	_ = v748
	var v754 int32
	_ = v754
	var v759 int32
	_ = v759
	var v763 int32
	_ = v763
	var v768 int32
	_ = v768
	var v772 int32
	_ = v772
	var v776 int32
	_ = v776
	var v781 int32
	_ = v781
	var v787 int32
	_ = v787
	var v791 int32
	_ = v791
	var v796 int32
	_ = v796
	var v799 int32
	_ = v799
	var v803 int32
	_ = v803
	var v806 int32
	_ = v806
	var v810 int32
	_ = v810
	var v812 int32
	_ = v812
	var v813 int32
	_ = v813
	var v817 int32
	_ = v817
	var v824 int32
	_ = v824
	var v826 int32
	_ = v826
	var v827 int32
	_ = v827
	var v833 int32
	_ = v833
	var v838 int32
	_ = v838
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
	v23 = *(*int32)(unsafe.Add(mBase, _consts[89]))
	if v23 == int32(0) {
		goto L8
	} else {
		goto L9
	}
L3:
	;
	F_SendPostmasterSignal(m, int32(4))
	mBase = m.M
	v799 = m.ExcPending
	if v799 != 0 {
		goto L1
	} else {
		goto L202
	}
L4:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v14)+112)) = v238
	F_errmsg(m, int32(673378), v14+int32(112))
	mBase = m.M
	v787 = m.ExcPending
	if v787 != 0 {
		goto L1
	} else {
		goto L199
	}
L5:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v772 = m.ExcPending
	if v772 != 0 {
		goto L1
	} else {
		goto L196
	}
L6:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v629 = m.ExcPending
	if v629 != 0 {
		goto L1
	} else {
		goto L165
	}
L7:
	;
	m.G0 = v14 + int32(144)
	return v618
L8:
	;
	if int32(0) < l0 {
		goto L42
	} else {
		goto L43
	}
L9:
	;
	if v23 == int32(4054512) {
		goto L8
	} else {
		goto L10
	}
L10:
	;
	v29 = l0 << (uint(int32(3)) % 32)
	v32 = v23
	goto L11
L11:
	;
	v43 = *(*int32)(unsafe.Add(mBase, uint32(v32-int32(4))))
	if v43 != l0 {
		goto L13
	} else {
		goto L14
	}
L12:
	;
	goto L8
L13:
	;
	v134 = *(*int32)(unsafe.Add(mBase, uint32(v32)+4))
	if v134 != int32(4054512) {
		v32 = v134
		goto L11
	} else {
		goto L41
	}
L14:
	;
	v46 = v32 + int32(8)
	if base.Ui32(int32(4)) <= base.Ui32(v29) {
		goto L18
	} else {
		goto L19
	}
L15:
	;
	if v108 != 0 {
		goto L13
	} else {
		goto L33
	}
L16:
	;
	v108 = int32(0)
	goto L15
L17:
	;
	v82 = v77
	v83 = v78
	v84 = v79
	goto L27
L18:
	;
	if (l1|v46)&int32(3) != 0 {
		v77 = l1
		v78 = v46
		v79 = v29
		goto L17
	} else {
		goto L21
	}
L19:
	;
	v70 = l1
	v71 = v46
	v72 = v29
	goto L20
L20:
	;
	if v72 == int32(0) {
		goto L16
	} else {
		goto L26
	}
L21:
	;
	v54 = l1
	v55 = v46
	v56 = v29
	goto L22
L22:
	;
	v59 = *(*int32)(unsafe.Add(mBase, uint32(v54)))
	v60 = *(*int32)(unsafe.Add(mBase, uint32(v55)))
	if v59 != v60 {
		v77 = v54
		v78 = v55
		v79 = v56
		goto L17
	} else {
		goto L24
	}
L23:
	;
	v70 = v65
	v71 = v63
	v72 = v67
	goto L20
L24:
	;
	v62 = int32(4)
	v63 = v55 + v62
	v65 = v54 + v62
	v67 = v56 - v62
	if base.Ui32(int32(3)) < base.Ui32(v67) {
		v54 = v65
		v55 = v63
		v56 = v67
		goto L22
	} else {
		goto L25
	}
L25:
	;
	goto L23
L26:
	;
	v77 = v70
	v78 = v71
	v79 = v72
	goto L17
L27:
	;
	v87 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v82))))
	v88 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v83))))
	if v87 == v88 {
		goto L29
	} else {
		goto L30
	}
L28:
	;
	v108 = v87 - v88
	goto L15
L29:
	;
	v90 = int32(1)
	v95 = v84 - v90
	if v95 != 0 {
		v82 = v82 + v90
		v83 = v83 + v90
		v84 = v95
		goto L27
	} else {
		goto L32
	}
L30:
	;
	goto L31
L31:
	;
	goto L28
L32:
	;
	goto L16
L33:
	;
	if v32 != v23 {
		goto L34
	} else {
		goto L35
	}
L34:
	;
	v112 = *(*int32)(unsafe.Add(mBase, uint32(v32)))
	v113 = *(*int32)(unsafe.Add(mBase, uint32(v32)+4))
	*(*int32)(unsafe.Add(mBase, uint32(v112)+4)) = v113
	v115 = *(*int32)(unsafe.Add(mBase, uint32(v32)))
	*(*int32)(unsafe.Add(mBase, uint32(v113))) = v115
	v118 = *(*int32)(unsafe.Add(mBase, _consts[89]))
	if v118 == int32(0) {
		goto L37
	} else {
		goto L38
	}
L35:
	;
	goto L36
L36:
	;
	v133 = *(*int32)(unsafe.Add(mBase, uint32(v32-int32(8))))
	if v133 != 0 {
		v618 = v133
		goto L7
	} else {
		goto L40
	}
L37:
	;
	v121 = int32(4054512)
	*(*int32)(unsafe.Add(mBase, _consts[90])) = v121
	v125 = v121
	goto L39
L38:
	;
	v125 = v118
	goto L39
L39:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v32))) = int32(4054512)
	*(*int32)(unsafe.Add(mBase, uint32(v32)+4)) = v125
	*(*int32)(unsafe.Add(mBase, uint32(v125))) = v32
	*(*int32)(unsafe.Add(mBase, _consts[89])) = v32
	goto L36
L40:
	;
	goto L8
L41:
	;
	goto L12
L42:
	;
	v153 = int32(0)
	v155 = int32(0)
	goto L45
L43:
	;
	goto L44
L44:
	;
	v190 = int32(*(*uint8)(unsafe.Add(mBase, _consts[1])))
	if v190 == int32(1) {
		goto L50
	} else {
		goto L51
	}
L45:
	;
	v167 = *(*int32)(unsafe.Add(mBase, uint32(l1+v153<<(uint(int32(3))%32))+4))
	if v155&int32(1)&base.B2i32(base.Ui32(int32(4)) <= base.Ui32(v167)) != 0 {
		goto L6
	} else {
		goto L47
	}
L46:
	;
	goto L44
L47:
	;
	v175 = v153 + int32(1)
	if v175 != l0 {
		v153 = v175
		v155 = base.B2i32(base.Ui32(int32(3)) < base.Ui32(v167)) | v155
		goto L45
	} else {
		goto L48
	}
L48:
	;
	goto L46
L49:
	;
	if v200 != 0 {
		goto L5
	} else {
		goto L53
	}
L50:
	;
	v195 = *(*int32)(unsafe.Add(mBase, _consts[2]))
	v196 = *(*int32)(unsafe.Add(mBase, uint32(v195)+316))
	v198 = base.B2i32(v196 != int32(2))
	*(*uint8)(unsafe.Add(mBase, _consts[1])) = uint8(v198)
	v200 = v198
	goto L52
L51:
	;
	v200 = int32(0)
	goto L52
L52:
	;
	goto L49
L53:
	;
	v202 = *(*int32)(unsafe.Add(mBase, _consts[44]))
	v206 = F_LWLockAcquire(m, v202+int32(1664), int32(0))
	mBase = m.M
	v207 = m.ExcPending
	if v207 != 0 {
		goto L1
	} else {
		goto L54
	}
L54:
	;
	v209 = *(*int32)(unsafe.Add(mBase, _consts[91]))
	v210 = *(*int32)(unsafe.Add(mBase, uint32(v209)))
	if v210 == int32(0) {
		goto L55
	} else {
		goto L56
	}
L55:
	;
	v213 = int32(1)
	*(*int32)(unsafe.Add(mBase, uint32(v209))) = v213
	v216 = v213
	goto L57
L56:
	;
	v216 = v210
	goto L57
L57:
	;
	v217 = *(*int32)(unsafe.Add(mBase, uint32(v209)+28))
	if int32(0) <= v216-v217 {
		goto L58
	} else {
		goto L59
	}
L58:
	;
	v221 = *(*int32)(unsafe.Add(mBase, uint32(v209)+16))
	v222 = *(*int32)(unsafe.Add(mBase, uint32(v209)+40))
	v223 = *(*int32)(unsafe.Add(mBase, uint32(v209)+36))
	v224 = *(*int32)(unsafe.Add(mBase, uint32(v209)+32))
	v226 = *(*int32)(unsafe.Add(mBase, _consts[44]))
	F_LWLockRelease(m, v226+int32(1664))
	mBase = m.M
	v230 = m.ExcPending
	if v230 != 0 {
		goto L1
	} else {
		goto L61
	}
L59:
	;
	v331 = v209
	v333 = v216
	goto L60
L60:
	;
	v336 = v333 + int32(1)
	if v336&int32(2047) == int32(0) {
		goto L95
	} else {
		goto L96
	}
L61:
	;
	v232 = int32(*(*uint8)(unsafe.Add(mBase, _consts[92])))
	if v232 != int32(1) {
		goto L62
	} else {
		goto L63
	}
L62:
	;
	if v216-v224 < int32(0) {
		goto L77
	} else {
		goto L78
	}
L63:
	;
	if int32(0) <= v216-v223 {
		goto L64
	} else {
		goto L65
	}
L64:
	;
	v238 = F_get_database_name(m, v221)
	mBase = m.M
	v239 = m.ExcPending
	if v239 != 0 {
		goto L1
	} else {
		goto L67
	}
L65:
	;
	goto L66
L66:
	;
	if v216&int32(65535) != 0 {
		goto L62
	} else {
		goto L75
	}
L67:
	;
	F_SendPostmasterSignal(m, int32(4))
	mBase = m.M
	v242 = m.ExcPending
	if v242 != 0 {
		goto L1
	} else {
		goto L68
	}
L68:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v246 = m.ExcPending
	if v246 != 0 {
		goto L1
	} else {
		goto L69
	}
L69:
	;
	F_errcode(m, int32(261))
	mBase = m.M
	v249 = m.ExcPending
	if v249 != 0 {
		goto L1
	} else {
		goto L70
	}
L70:
	;
	if v238 != 0 {
		goto L4
	} else {
		goto L71
	}
L71:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v14)+96)) = v221
	F_errmsg(m, int32(53920), v14+int32(96))
	mBase = m.M
	v255 = m.ExcPending
	if v255 != 0 {
		goto L1
	} else {
		goto L72
	}
L72:
	;
	F_errhint(m, int32(546475), int32(0))
	mBase = m.M
	v259 = m.ExcPending
	if v259 != 0 {
		goto L1
	} else {
		goto L73
	}
L73:
	;
	F_errfinish(m, int32(472009), int32(1189), int32(445134))
	mBase = m.M
	v264 = m.ExcPending
	if v264 != 0 {
		goto L1
	} else {
		goto L74
	}
L74:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L75:
	;
	F_SendPostmasterSignal(m, int32(4))
	mBase = m.M
	v269 = m.ExcPending
	if v269 != 0 {
		goto L1
	} else {
		goto L76
	}
L76:
	;
	goto L62
L77:
	;
	v316 = *(*int32)(unsafe.Add(mBase, _consts[44]))
	v320 = F_LWLockAcquire(m, v316+int32(1664), int32(0))
	mBase = m.M
	v321 = m.ExcPending
	if v321 != 0 {
		goto L1
	} else {
		goto L91
	}
L78:
	;
	v273 = F_get_database_name(m, v221)
	mBase = m.M
	v274 = m.ExcPending
	if v274 != 0 {
		goto L1
	} else {
		goto L79
	}
L79:
	;
	v277 = F_errstart(m, int32(19), int32(0))
	mBase = m.M
	v278 = m.ExcPending
	if v278 != 0 {
		goto L1
	} else {
		goto L80
	}
L80:
	;
	if v273 != 0 {
		goto L82
	} else {
		goto L83
	}
L81:
	;
	F_errhint(m, int32(546475), int32(0))
	mBase = m.M
	v308 = m.ExcPending
	if v308 != 0 {
		goto L1
	} else {
		goto L89
	}
L82:
	;
	if v277 == int32(0) {
		goto L77
	} else {
		goto L85
	}
L83:
	;
	goto L84
L84:
	;
	if v277 == int32(0) {
		goto L77
	} else {
		goto L87
	}
L85:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v14)+80)) = v273
	v282 = v222 - v216
	*(*int32)(unsafe.Add(mBase, uint32(v14)+84)) = v282
	F_errmsg_plural(m, int32(430071), int32(430402), v282, v14+int32(80))
	mBase = m.M
	v289 = m.ExcPending
	if v289 != 0 {
		goto L1
	} else {
		goto L86
	}
L86:
	;
	v304 = int32(1213)
	goto L81
L87:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v14)+64)) = v221
	v294 = v222 - v216
	*(*int32)(unsafe.Add(mBase, uint32(v14)+68)) = v294
	F_errmsg_plural(m, int32(429998), int32(430327), v294, v14-int32(-64))
	mBase = m.M
	v301 = m.ExcPending
	if v301 != 0 {
		goto L1
	} else {
		goto L88
	}
L88:
	;
	v304 = int32(1222)
	goto L81
L89:
	;
	F_errfinish(m, int32(472009), v304, int32(445134))
	mBase = m.M
	v312 = m.ExcPending
	if v312 != 0 {
		goto L1
	} else {
		goto L90
	}
L90:
	;
	goto L77
L91:
	;
	v322 = int32(1)
	v324 = *(*int32)(unsafe.Add(mBase, _consts[91]))
	v325 = *(*int32)(unsafe.Add(mBase, uint32(v324)))
	if base.Ui32(v325) <= base.Ui32(v322) {
		goto L92
	} else {
		goto L93
	}
L92:
	;
	v328 = v322
	goto L94
L93:
	;
	v328 = v325
	goto L94
L94:
	;
	v331 = v324
	v333 = v328
	goto L60
L95:
	;
	v342 = *(*int32)(unsafe.Add(mBase, _consts[93]))
	v343 = *(*int32)(unsafe.Add(mBase, uint32(v342)+28))
	v345 = int32(base.Ui32(v336) >> (uint(int32(11)) % 32))
	v347 = int32(*(*uint16)(unsafe.Add(mBase, _consts[94])))
	v348 = base.I32_rem_u_s(v345, v347)
	v351 = v343 + v348<<(uint(int32(7))%32)
	v353 = F_LWLockAcquire(m, v351, int32(0))
	mBase = m.M
	v354 = m.ExcPending
	if v354 != 0 {
		goto L1
	} else {
		goto L98
	}
L96:
	;
	v376 = v331
	goto L97
L97:
	;
	v378 = *(*int32)(unsafe.Add(mBase, uint32(v376)+4))
	v381 = l0 + base.B2i32(v378 == int32(0))
	v382 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v376)+24)))
	if v382 == int32(1) {
		goto L107
	} else {
		goto L108
	}
L98:
	;
	v356 = base.I64_extend_i32_u(v345)
	v357 = F_SimpleLruZeroPage(m, int32(4338564), v356)
	mBase = m.M
	v358 = m.ExcPending
	if v358 != 0 {
		goto L1
	} else {
		goto L99
	}
L99:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v14)+128)) = v356
	F_XLogBeginInsert(m)
	mBase = m.M
	v361 = m.ExcPending
	if v361 != 0 {
		goto L1
	} else {
		goto L100
	}
L100:
	;
	F_XLogRegisterData(m, v14+int32(128), int32(8))
	mBase = m.M
	v366 = m.ExcPending
	if v366 != 0 {
		goto L1
	} else {
		goto L101
	}
L101:
	;
	v369 = F_XLogInsert(m, int32(6), int32(0))
	mBase = m.M
	v370 = m.ExcPending
	if v370 != 0 {
		goto L1
	} else {
		goto L102
	}
L102:
	;
	F_LWLockRelease(m, v351)
	mBase = m.M
	v372 = m.ExcPending
	if v372 != 0 {
		goto L1
	} else {
		goto L103
	}
L103:
	;
	v374 = *(*int32)(unsafe.Add(mBase, _consts[91]))
	v376 = v374
	goto L97
L104:
	;
	v477 = int32(1)
	if base.Ui32(v378) <= base.Ui32(v477) {
		goto L137
	} else {
		goto L138
	}
L105:
	;
	v432 = v378 + v381 + int32(1047040)
	v434 = v432 + base.B2i32(base.Ui32(v432) < base.Ui32(v378))
	v435 = *(*int32)(unsafe.Add(mBase, uint32(v427)+44))
	if base.Ui32(v378) < base.Ui32(v435) {
		goto L124
	} else {
		goto L125
	}
L106:
	;
	F_SendPostmasterSignal(m, int32(4))
	mBase = m.M
	v418 = m.ExcPending
	if v418 != 0 {
		goto L1
	} else {
		goto L121
	}
L107:
	;
	v385 = v378 + v381
	v387 = v385 + base.B2i32(base.Ui32(v385) < base.Ui32(v381))
	v388 = *(*int32)(unsafe.Add(mBase, uint32(v376)+44))
	if base.Ui32(v378) < base.Ui32(v388) {
		goto L111
	} else {
		goto L112
	}
L108:
	;
	goto L109
L109:
	;
	v406 = int32(1636)
	v407 = base.I32_div_u_s(v378+v381, v406)
	v409 = base.I32_div_u_s(v378, v406)
	if base.Ui32(v407^v409) < base.Ui32(int32(32)) {
		goto L104
	} else {
		goto L120
	}
L110:
	;
	v394 = *(*int32)(unsafe.Add(mBase, uint32(v376)+20))
	if int32(0) <= v378-v394 {
		v427 = v376
		goto L105
	} else {
		goto L118
	}
L111:
	;
	if base.Ui32(v387) < base.Ui32(v378) {
		goto L3
	} else {
		goto L114
	}
L112:
	;
	goto L113
L113:
	;
	if base.Ui32(v378) <= base.Ui32(v387) {
		goto L110
	} else {
		goto L116
	}
L114:
	;
	if base.Ui32(v387) < base.Ui32(v388) {
		goto L110
	} else {
		goto L115
	}
L115:
	;
	goto L3
L116:
	;
	if base.Ui32(v388) <= base.Ui32(v387) {
		goto L3
	} else {
		goto L117
	}
L117:
	;
	goto L110
L118:
	;
	v398 = int32(1636)
	v399 = base.I32_div_u_s(v385, v398)
	v401 = base.I32_div_u_s(v378, v398)
	if base.Ui32(int32(32)) <= base.Ui32(v399^v401) {
		goto L106
	} else {
		goto L119
	}
L119:
	;
	v427 = v376
	goto L105
L120:
	;
	goto L106
L121:
	;
	v420 = *(*int32)(unsafe.Add(mBase, _consts[91]))
	v421 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v420)+24)))
	if v421&int32(1) == int32(0) {
		goto L104
	} else {
		goto L122
	}
L122:
	;
	v427 = v420
	goto L105
L123:
	;
	v443 = F_errstart(m, int32(19), int32(0))
	mBase = m.M
	v444 = m.ExcPending
	if v444 != 0 {
		goto L1
	} else {
		goto L131
	}
L124:
	;
	if base.Ui32(v434) < base.Ui32(v378) {
		goto L123
	} else {
		goto L127
	}
L125:
	;
	goto L126
L126:
	;
	if base.Ui32(v378) <= base.Ui32(v434) {
		goto L104
	} else {
		goto L129
	}
L127:
	;
	if base.Ui32(v435) <= base.Ui32(v434) {
		goto L123
	} else {
		goto L128
	}
L128:
	;
	goto L104
L129:
	;
	if base.Ui32(v434) < base.Ui32(v435) {
		goto L104
	} else {
		goto L130
	}
L130:
	;
	goto L123
L131:
	;
	if v443 == int32(0) {
		goto L104
	} else {
		goto L132
	}
L132:
	;
	F_errcode(m, int32(261))
	mBase = m.M
	v449 = m.ExcPending
	if v449 != 0 {
		goto L1
	} else {
		goto L133
	}
L133:
	;
	v451 = *(*int32)(unsafe.Add(mBase, _consts[91]))
	v452 = *(*int32)(unsafe.Add(mBase, uint32(v451)+44))
	v453 = *(*int32)(unsafe.Add(mBase, uint32(v451)+16))
	*(*int32)(unsafe.Add(mBase, uint32(v14)+48)) = v453
	v456 = v452 - v378 + v381
	*(*int32)(unsafe.Add(mBase, uint32(v14)+52)) = v456
	F_errmsg_plural(m, int32(429920), int32(430247), v456, v14+int32(48))
	mBase = m.M
	v463 = m.ExcPending
	if v463 != 0 {
		goto L1
	} else {
		goto L134
	}
L134:
	;
	F_errhint(m, int32(556720), int32(0))
	mBase = m.M
	v467 = m.ExcPending
	if v467 != 0 {
		goto L1
	} else {
		goto L135
	}
L135:
	;
	F_errfinish(m, int32(472009), int32(1322), int32(445134))
	mBase = m.M
	v472 = m.ExcPending
	if v472 != 0 {
		goto L1
	} else {
		goto L136
	}
L136:
	;
	goto L104
L137:
	;
	v480 = v477
	goto L139
L138:
	;
	v480 = v378
	goto L139
L139:
	;
	if int32(0) < v381 {
		goto L140
	} else {
		goto L141
	}
L140:
	;
	v485 = v378
	v486 = v381
	goto L143
L141:
	;
	goto L142
L142:
	;
	v562 = int32(4438516)
	v564 = *(*int32)(unsafe.Add(mBase, _consts[14]))
	v565 = int32(1)
	*(*int32)(unsafe.Add(mBase, _consts[14])) = v564 + v565
	v569 = *(*int32)(unsafe.Add(mBase, _consts[91]))
	v570 = *(*int32)(unsafe.Add(mBase, uint32(v569)))
	*(*int32)(unsafe.Add(mBase, uint32(v569))) = v570 + v565
	v574 = *(*int32)(unsafe.Add(mBase, uint32(v569)+4))
	*(*int32)(unsafe.Add(mBase, uint32(v569)+4)) = v574 + v381
	v578 = *(*int32)(unsafe.Add(mBase, _consts[44]))
	F_LWLockRelease(m, v578+int32(1664))
	mBase = m.M
	v582 = m.ExcPending
	if v582 != 0 {
		goto L1
	} else {
		goto L158
	}
L143:
	;
	v497 = base.I32_rem_u_s(int32(base.Ui32(v485)>>(uint(int32(2))%32)), int32(409))
	if v497|v485&int32(3) == int32(0) {
		goto L145
	} else {
		goto L146
	}
L144:
	;
	goto L142
L145:
	;
	v504 = *(*int32)(unsafe.Add(mBase, _consts[95]))
	v505 = *(*int32)(unsafe.Add(mBase, uint32(v504)+28))
	v507 = base.I32_div_u_s(v485, int32(1636))
	v509 = int32(*(*uint16)(unsafe.Add(mBase, _consts[96])))
	v510 = base.I32_rem_u_s(v507, v509)
	v513 = v505 + v510<<(uint(int32(7))%32)
	v515 = F_LWLockAcquire(m, v513, int32(0))
	mBase = m.M
	v516 = m.ExcPending
	if v516 != 0 {
		goto L1
	} else {
		goto L148
	}
L146:
	;
	goto L147
L147:
	;
	if base.Ui32(int32(-1036)) <= base.Ui32(v485) {
		goto L154
	} else {
		goto L155
	}
L148:
	;
	v518 = base.I64_extend_i32_u(v507)
	v519 = F_SimpleLruZeroPage(m, int32(4338644), v518)
	mBase = m.M
	v520 = m.ExcPending
	if v520 != 0 {
		goto L1
	} else {
		goto L149
	}
L149:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v14)+128)) = v518
	F_XLogBeginInsert(m)
	mBase = m.M
	v523 = m.ExcPending
	if v523 != 0 {
		goto L1
	} else {
		goto L150
	}
L150:
	;
	F_XLogRegisterData(m, v14+int32(128), int32(8))
	mBase = m.M
	v528 = m.ExcPending
	if v528 != 0 {
		goto L1
	} else {
		goto L151
	}
L151:
	;
	v531 = F_XLogInsert(m, int32(6), int32(16))
	mBase = m.M
	v532 = m.ExcPending
	if v532 != 0 {
		goto L1
	} else {
		goto L152
	}
L152:
	;
	F_LWLockRelease(m, v513)
	mBase = m.M
	v534 = m.ExcPending
	if v534 != 0 {
		goto L1
	} else {
		goto L153
	}
L153:
	;
	goto L147
L154:
	;
	v546 = int32(0) - v485
	goto L156
L155:
	;
	v542 = int32(1636)
	v544 = base.I32_rem_u_s(v485, v542)
	v546 = v542 - v544
	goto L156
L156:
	;
	v548 = v486 - v546
	if int32(0) < v548 {
		v485 = v485 + v546
		v486 = v548
		goto L143
	} else {
		goto L157
	}
L157:
	;
	goto L144
L158:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v14)+136)) = l0
	*(*int32)(unsafe.Add(mBase, uint32(v14)+132)) = v480
	*(*int32)(unsafe.Add(mBase, uint32(v14)+128)) = v333
	F_XLogBeginInsert(m)
	mBase = m.M
	v587 = m.ExcPending
	if v587 != 0 {
		goto L1
	} else {
		goto L159
	}
L159:
	;
	F_XLogRegisterData(m, v14+int32(128), int32(12))
	mBase = m.M
	v592 = m.ExcPending
	if v592 != 0 {
		goto L1
	} else {
		goto L160
	}
L160:
	;
	F_XLogRegisterData(m, l1, l0<<(uint(int32(3))%32))
	mBase = m.M
	v596 = m.ExcPending
	if v596 != 0 {
		goto L1
	} else {
		goto L161
	}
L161:
	;
	v599 = F_XLogInsert(m, int32(6), int32(32))
	mBase = m.M
	v600 = m.ExcPending
	if v600 != 0 {
		goto L1
	} else {
		goto L162
	}
L162:
	;
	F_RecordNewMultiXact(m, v333, v480, l0, l1)
	mBase = m.M
	v602 = m.ExcPending
	if v602 != 0 {
		goto L1
	} else {
		goto L163
	}
L163:
	;
	v603 = int32(4438516)
	v605 = *(*int32)(unsafe.Add(mBase, _consts[14]))
	*(*int32)(unsafe.Add(mBase, _consts[14])) = v605 - int32(1)
	F_mXactCachePut(m, v333, l0, l1)
	mBase = m.M
	v610 = m.ExcPending
	if v610 != 0 {
		goto L1
	} else {
		goto L164
	}
L164:
	;
	v618 = v333
	goto L7
L165:
	;
	v630 = m.G0
	v632 = v630 - int32(80)
	m.G0 = v632
	v635 = *(*int32)(unsafe.Add(mBase, _consts[97]))
	if v635 != 0 {
		goto L167
	} else {
		goto L168
	}
L166:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v14))) = v726
	F_errmsg_internal(m, int32(192120), v14)
	mBase = m.M
	v763 = m.ExcPending
	if v763 != 0 {
		goto L1
	} else {
		goto L194
	}
L167:
	;
	F_pfree(m, v635)
	mBase = m.M
	v637 = m.ExcPending
	if v637 != 0 {
		goto L1
	} else {
		goto L170
	}
L168:
	;
	goto L169
L169:
	;
	F_initStringInfo(m, v632-int32(-64))
	mBase = m.M
	v641 = m.ExcPending
	if v641 != 0 {
		goto L1
	} else {
		goto L171
	}
L170:
	;
	goto L169
L171:
	;
	v642 = *(*int32)(unsafe.Add(mBase, uint32(l1)+4))
	if base.Ui32(v642) < base.Ui32(int32(6)) {
		goto L173
	} else {
		goto L174
	}
L172:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v748 = m.ExcPending
	if v748 != 0 {
		goto L1
	} else {
		goto L191
	}
L173:
	;
	v645 = *(*int32)(unsafe.Add(mBase, uint32(l1)))
	*(*int32)(unsafe.Add(mBase, uint32(v632)+40)) = v645
	*(*int32)(unsafe.Add(mBase, uint32(v632)+32)) = int32(0)
	v653 = *(*int32)(unsafe.Add(mBase, uint32(v642<<(uint(int32(2))%32))+uint32(_consts[98])))
	*(*int32)(unsafe.Add(mBase, uint32(v632)+44)) = v653
	*(*int32)(unsafe.Add(mBase, uint32(v632)+36)) = l0
	F_appendStringInfo(m, v632-int32(-64), int32(635096), v632+int32(32))
	mBase = m.M
	v662 = m.ExcPending
	if v662 != 0 {
		goto L1
	} else {
		goto L176
	}
L174:
	;
	goto L175
L175:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v733 = m.ExcPending
	if v733 != 0 {
		goto L1
	} else {
		goto L188
	}
L176:
	;
	if int32(2) <= l0 {
		goto L177
	} else {
		goto L178
	}
L177:
	;
	v669 = int32(1)
	goto L180
L178:
	;
	goto L179
L179:
	;
	F_appendStringInfoChar(m, v632-int32(-64), int32(93))
	mBase = m.M
	v714 = m.ExcPending
	if v714 != 0 {
		goto L1
	} else {
		goto L185
	}
L180:
	;
	v679 = l1 + v669<<(uint(int32(3))%32)
	v680 = *(*int32)(unsafe.Add(mBase, uint32(v679)+4))
	if base.Ui32(int32(6)) <= base.Ui32(v680) {
		goto L172
	} else {
		goto L182
	}
L181:
	;
	goto L179
L182:
	;
	v683 = *(*int32)(unsafe.Add(mBase, uint32(v679)))
	*(*int32)(unsafe.Add(mBase, uint32(v632))) = v683
	v689 = *(*int32)(unsafe.Add(mBase, uint32(v680<<(uint(int32(2))%32))+uint32(_consts[98])))
	*(*int32)(unsafe.Add(mBase, uint32(v632)+4)) = v689
	F_appendStringInfo(m, v632-int32(-64), int32(635181), v632)
	mBase = m.M
	v695 = m.ExcPending
	if v695 != 0 {
		goto L1
	} else {
		goto L183
	}
L183:
	;
	v697 = v669 + int32(1)
	if v697 != l0 {
		v669 = v697
		goto L180
	} else {
		goto L184
	}
L184:
	;
	goto L181
L185:
	;
	v717 = *(*int32)(unsafe.Add(mBase, _consts[12]))
	v718 = *(*int32)(unsafe.Add(mBase, uint32(v632)+64))
	v719 = F_MemoryContextStrdup(m, v717, v718)
	mBase = m.M
	v720 = m.ExcPending
	if v720 != 0 {
		goto L1
	} else {
		goto L186
	}
L186:
	;
	*(*int32)(unsafe.Add(mBase, _consts[97])) = v719
	v722 = *(*int32)(unsafe.Add(mBase, uint32(v632)+64))
	F_pfree(m, v722)
	mBase = m.M
	v724 = m.ExcPending
	if v724 != 0 {
		goto L1
	} else {
		goto L187
	}
L187:
	;
	v726 = *(*int32)(unsafe.Add(mBase, _consts[97]))
	m.G0 = v632 + int32(80)
	goto L166
L188:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v632)+48)) = v642
	F_errmsg_internal(m, int32(448718), v632+int32(48))
	mBase = m.M
	v739 = m.ExcPending
	if v739 != 0 {
		goto L1
	} else {
		goto L189
	}
L189:
	;
	F_errfinish(m, int32(472009), int32(1834), int32(314735))
	mBase = m.M
	v744 = m.ExcPending
	if v744 != 0 {
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
	*(*int32)(unsafe.Add(mBase, uint32(v632)+16)) = v680
	F_errmsg_internal(m, int32(448718), v632+int32(16))
	mBase = m.M
	v754 = m.ExcPending
	if v754 != 0 {
		goto L1
	} else {
		goto L192
	}
L192:
	;
	F_errfinish(m, int32(472009), int32(1834), int32(314735))
	mBase = m.M
	v759 = m.ExcPending
	if v759 != 0 {
		goto L1
	} else {
		goto L193
	}
L193:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L194:
	;
	F_errfinish(m, int32(472009), int32(849), int32(127226))
	mBase = m.M
	v768 = m.ExcPending
	if v768 != 0 {
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
	F_errmsg_internal(m, int32(13552), int32(0))
	mBase = m.M
	v776 = m.ExcPending
	if v776 != 0 {
		goto L1
	} else {
		goto L197
	}
L197:
	;
	F_errfinish(m, int32(472009), int32(1123), int32(445134))
	mBase = m.M
	v781 = m.ExcPending
	if v781 != 0 {
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
	F_errhint(m, int32(546475), int32(0))
	mBase = m.M
	v791 = m.ExcPending
	if v791 != 0 {
		goto L1
	} else {
		goto L200
	}
L200:
	;
	F_errfinish(m, int32(472009), int32(1182), int32(445134))
	mBase = m.M
	v796 = m.ExcPending
	if v796 != 0 {
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
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v803 = m.ExcPending
	if v803 != 0 {
		goto L1
	} else {
		goto L203
	}
L203:
	;
	F_errcode(m, int32(261))
	mBase = m.M
	v806 = m.ExcPending
	if v806 != 0 {
		goto L1
	} else {
		goto L204
	}
L204:
	;
	F_errmsg(m, int32(442299), int32(0))
	mBase = m.M
	v810 = m.ExcPending
	if v810 != 0 {
		goto L1
	} else {
		goto L205
	}
L205:
	;
	v812 = *(*int32)(unsafe.Add(mBase, _consts[91]))
	v813 = *(*int32)(unsafe.Add(mBase, uint32(v812)+44))
	*(*int32)(unsafe.Add(mBase, uint32(v14)+32)) = v381
	v817 = v813 + (v378 ^ int32(-1))
	*(*int32)(unsafe.Add(mBase, uint32(v14)+36)) = v817
	F_errdetail_plural(m, int32(571807), int32(550561), v817, v14+int32(32))
	mBase = m.M
	v824 = m.ExcPending
	if v824 != 0 {
		goto L1
	} else {
		goto L206
	}
L206:
	;
	v826 = *(*int32)(unsafe.Add(mBase, _consts[91]))
	v827 = *(*int32)(unsafe.Add(mBase, uint32(v826)+16))
	*(*int32)(unsafe.Add(mBase, uint32(v14)+16)) = v827
	F_errhint(m, int32(556568), v14+int32(16))
	mBase = m.M
	v833 = m.ExcPending
	if v833 != 0 {
		goto L1
	} else {
		goto L207
	}
L207:
	;
	F_errfinish(m, int32(472009), int32(1287), int32(445134))
	mBase = m.M
	v838 = m.ExcPending
	if v838 != 0 {
		goto L1
	} else {
		goto L208
	}
L208:
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
	v3 = *(*int32)(unsafe.Add(mBase, _consts[44]))
	v7 = F_LWLockAcquire(m, v3+int32(1664), int32(1))
	mBase = m.M
	v10 = m.ExcPending
	if v10 != 0 {
		return int32(0)
	} else {
		v12 = *(*int32)(unsafe.Add(mBase, _consts[91]))
		v13 = *(*int32)(unsafe.Add(mBase, uint32(v12)))
		v15 = *(*int32)(unsafe.Add(mBase, _consts[44]))
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
