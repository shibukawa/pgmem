package p5

import (
	base "github.com/shibukawa/pgmem/internal/aot/pgaot/base"
	"unsafe"
)

func F_pgaio_io_reclaim(m *base.Module, l0 int32) {
	mBase := m.M
	_ = mBase
	var v2 int32
	_ = v2
	var v30 int32
	_ = v30
	var v32 int32
	_ = v32
	var v34 int32
	_ = v34
	var v36 int32
	_ = v36
	var v40 int32
	_ = v40
	var v43 int32
	_ = v43
	var v45 int32
	_ = v45
	var v47 int32
	_ = v47
	var v49 int32
	_ = v49
	var v54 int32
	_ = v54
	var v55 int64
	_ = v55
	var v57 int32
	_ = v57
	var v80 int32
	_ = v80
	var v108 int32
	_ = v108
	var v110 int32
	_ = v110
	var v112 int32
	_ = v112
	var v115 int32
	_ = v115
	var v116 int32
	_ = v116
	var v118 int32
	_ = v118
	var v121 int32
	_ = v121
	var v122 int32
	_ = v122
	var v124 int32
	_ = v124
	var v126 int32
	_ = v126
	var v128 int32
	_ = v128
	var v129 int32
	_ = v129
	var v134 int32
	_ = v134
	var v141 int32
	_ = v141
	var v142 int32
	_ = v142
	var v143 int32
	_ = v143
	var v148 int32
	_ = v148
	var v149 int32
	_ = v149
	var v151 int32
	_ = v151
	var v158 int32
	_ = v158
	var v159 int32
	_ = v159
	var v160 int32
	_ = v160
	var v164 int32
	_ = v164
	var v172 int32
	_ = v172
	var v173 int32
	_ = v173
	var v174 int32
	_ = v174
	var v175 int32
	_ = v175
	var v179 int32
	_ = v179
	var v197 int32
	_ = v197
	var v202 int32
	_ = v202
	var v209 int32
	_ = v209
	var v210 int64
	_ = v210
	var v217 int32
	_ = v217
	var v218 int64
	_ = v218
	var v261 int32
	_ = v261
	var v262 int32
	_ = v262
	var v264 int32
	_ = v264
	var v266 int32
	_ = v266
	var v268 int32
	_ = v268
	var v269 int32
	_ = v269
	var v274 int32
	_ = v274
	var v281 int32
	_ = v281
	var v282 int32
	_ = v282
	var v283 int32
	_ = v283
	var v288 int32
	_ = v288
	var v289 int32
	_ = v289
	var v291 int32
	_ = v291
	var v298 int32
	_ = v298
	var v299 int32
	_ = v299
	var v300 int32
	_ = v300
	var v304 int32
	_ = v304
	var v312 int32
	_ = v312
	var v313 int32
	_ = v313
	var v314 int32
	_ = v314
	var v315 int32
	_ = v315
	var v316 int32
	_ = v316
	var v332 int32
	_ = v332
	var v337 int32
	_ = v337
	var v345 int32
	_ = v345
	var v347 int32
	_ = v347
	var v354 int64
	_ = v354
	var v357 int32
	_ = v357
	var v358 int32
	_ = v358
	var v362 int32
	_ = v362
	var v363 int64
	_ = v363
	var v365 int64
	_ = v365
	var v367 int64
	_ = v367
	var v400 int32
	_ = v400
	var v401 int32
	_ = v401
	var v403 int32
	_ = v403
	var v405 int32
	_ = v405
	var v407 int32
	_ = v407
	var v408 int32
	_ = v408
	var v411 int32
	_ = v411
	var v418 int32
	_ = v418
	var v419 int32
	_ = v419
	var v420 int32
	_ = v420
	var v425 int32
	_ = v425
	var v426 int32
	_ = v426
	var v427 int32
	_ = v427
	var v434 int32
	_ = v434
	var v435 int32
	_ = v435
	var v436 int32
	_ = v436
	var v438 int32
	_ = v438
	var v442 int32
	_ = v442
	var v449 int32
	_ = v449
	var v450 int32
	_ = v450
	var v451 int32
	_ = v451
	var v466 int32
	_ = v466
	var v471 int32
	_ = v471
	var v479 int32
	_ = v479
	var v483 int32
	_ = v483
	var v484 int32
	_ = v484
	var v485 int32
	_ = v485
	var v487 int32
	_ = v487
	var v489 int32
	_ = v489
	var v495 int32
	_ = v495
	var v497 int32
	_ = v497
	var v499 int32
	_ = v499
	var v500 int32
	_ = v500
	var v502 int32
	_ = v502
	var v506 int64
	_ = v506
	var v512 int32
	_ = v512
	var v513 int32
	_ = v513
	var v521 int32
	_ = v521
	var v526 int32
	_ = v526
	var v528 int32
	_ = v528
	var v529 int32
	_ = v529
	var v537 int32
	_ = v537
	var v541 int32
	_ = v541
	var v544 int32
	_ = v544
	var v545 int32
	_ = v545
	var v548 int32
	_ = v548
	var v550 int32
	_ = v550
	v2 = int32(0)
	v30 = m.G0
	v32 = v30 - int32(48)
	m.G0 = v32
	v34 = int32(4543420)
	v36 = *(*int32)(unsafe.Add(mBase, _consts[412]))
	*(*int32)(unsafe.Add(mBase, _consts[412])) = v36 + int32(1)
	v40 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0))))
	if v40 != int32(6) {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	v400 = F_errstart(m, int32(11), int32(0))
	mBase = m.M
	v401 = m.ExcPending
	if v401 != 0 {
		goto L11
	} else {
		goto L60
	}
L2:
	;
	v43 = m.G0
	v45 = v43 - int32(112)
	m.G0 = v45
	v47 = int32(4543428)
	v49 = *(*int32)(unsafe.Add(mBase, _consts[14]))
	*(*int32)(unsafe.Add(mBase, _consts[14])) = v49 + int32(1)
	v54 = v32 + int32(40)
	v55 = *(*int64)(unsafe.Add(mBase, uint32(l0)+68))
	*(*int64)(unsafe.Add(mBase, uint32(v54))) = v55
	v57 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+4)))
	if v57 != 0 {
		goto L3
	} else {
		goto L4
	}
L3:
	;
	v80 = v57
	goto L6
L4:
	;
	goto L5
L5:
	;
	v261 = F_errstart(m, int32(12), int32(0))
	mBase = m.M
	v262 = m.ExcPending
	if v262 != 0 {
		goto L11
	} else {
		goto L36
	}
L6:
	;
	v108 = v80 - int32(1)
	v110 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0+int32(5)+v108))))
	v112 = v110 << (uint(int32(3)) % 32)
	v115 = *(*int32)(unsafe.Add(mBase, uint32(v112)+uint32(_consts[589])))
	v116 = *(*int32)(unsafe.Add(mBase, uint32(v115)+8))
	if v116 != 0 {
		goto L8
	} else {
		goto L9
	}
L7:
	;
	goto L5
L8:
	;
	v118 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v108+(l0+int32(9))))))
	v121 = F_errstart(m, int32(11), int32(0))
	mBase = m.M
	v122 = m.ExcPending
	if v122 != 0 {
		goto L11
	} else {
		goto L12
	}
L9:
	;
	goto L10
L10:
	;
	if base.Ui32(int32(1)) < base.Ui32(v80) {
		v80 = v108
		goto L6
	} else {
		goto L35
	}
L11:
	;
	return
L12:
	;
	if v121 != 0 {
		goto L13
	} else {
		goto L14
	}
L13:
	;
	F_errhidestmt(m)
	mBase = m.M
	v124 = m.ExcPending
	if v124 != 0 {
		goto L11
	} else {
		goto L16
	}
L14:
	;
	goto L15
L15:
	;
	v209 = *(*int32)(unsafe.Add(mBase, uint32(v115)+8))
	v210 = *(*int64)(unsafe.Add(mBase, uint32(v54)))
	*(*int64)(unsafe.Add(mBase, uint32(v45)+40)) = v210
	m.T0[v209].(func(*base.Module, int32, int32, int32, int32))(m, v45+int32(104), l0, v45+int32(40), v118)
	mBase = m.M
	v217 = m.ExcPending
	if v217 != 0 {
		goto L11
	} else {
		goto L34
	}
L16:
	;
	F_errhidecontext(m)
	mBase = m.M
	v126 = m.ExcPending
	if v126 != 0 {
		goto L11
	} else {
		goto L17
	}
L17:
	;
	v128 = *(*int32)(unsafe.Add(mBase, _consts[590]))
	v129 = *(*int32)(unsafe.Add(mBase, uint32(v128)+24))
	goto L18
L18:
	;
	v134 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+2)))
	if base.Ui32(v134) <= base.Ui32(int32(2)) {
		goto L20
	} else {
		goto L21
	}
L19:
	;
	v143 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+1)))
	v148 = *(*int32)(unsafe.Add(mBase, uint32(v143<<(uint(int32(2))%32))+uint32(_consts[591])))
	v149 = *(*int32)(unsafe.Add(mBase, uint32(v148)+8))
	goto L23
L20:
	;
	v141 = *(*int32)(unsafe.Add(mBase, uint32(v134<<(uint(int32(2))%32))+uint32(_consts[592])))
	v142 = v141
	goto L22
L21:
	;
	v142 = int32(0)
	goto L22
L22:
	;
	goto L19
L23:
	;
	v151 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0))))
	if base.Ui32(v151) <= base.Ui32(int32(7)) {
		goto L25
	} else {
		goto L26
	}
L24:
	;
	v160 = *(*int32)(unsafe.Add(mBase, uint32(v54)))
	v164 = int32(base.Ui32(v160)>>(uint(int32(6))%32)) & int32(7)
	if base.Ui32(v164) <= base.Ui32(int32(4)) {
		goto L29
	} else {
		goto L30
	}
L25:
	;
	v158 = *(*int32)(unsafe.Add(mBase, uint32(v151<<(uint(int32(2))%32))+uint32(_consts[593])))
	v159 = v158
	goto L27
L26:
	;
	v159 = int32(0)
	goto L27
L27:
	;
	goto L24
L28:
	;
	v174 = *(*int32)(unsafe.Add(mBase, uint32(v54)))
	v175 = *(*int32)(unsafe.Add(mBase, uint32(v54)+4))
	*(*int32)(unsafe.Add(mBase, uint32(v45+int32(92)))) = v175
	*(*int32)(unsafe.Add(mBase, uint32(v45+int32(80)))) = v173
	*(*int32)(unsafe.Add(mBase, uint32(v45+int32(76)))) = v118
	v179 = *(*int32)(unsafe.Add(mBase, uint32(v112)+uint32(_consts[594])))
	*(*int32)(unsafe.Add(mBase, uint32(v45+int32(72)))) = v179
	*(*int32)(unsafe.Add(mBase, uint32(v45+int32(68)))) = v110
	*(*int32)(unsafe.Add(mBase, uint32(v45-int32(-64)))) = v80
	*(*int32)(unsafe.Add(mBase, uint32(v45+int32(88)))) = int32(base.Ui32(v174) >> (uint(int32(9)) % 32))
	*(*int32)(unsafe.Add(mBase, uint32(v45+int32(84)))) = v174 & int32(63)
	*(*int32)(unsafe.Add(mBase, uint32(v45)+60)) = v159
	*(*int32)(unsafe.Add(mBase, uint32(v45)+56)) = v149
	*(*int32)(unsafe.Add(mBase, uint32(v45)+52)) = v142
	*(*int32)(unsafe.Add(mBase, uint32(v45)+48)) = (l0 - v129) >> (uint(int32(7)) % 32)
	F_errmsg_internal(m, int32(485919), v45+int32(48))
	mBase = m.M
	v197 = m.ExcPending
	if v197 != 0 {
		goto L11
	} else {
		goto L32
	}
L29:
	;
	v172 = *(*int32)(unsafe.Add(mBase, uint32(v164<<(uint(int32(2))%32))+uint32(_consts[595])))
	v173 = v172
	goto L31
L30:
	;
	v173 = int32(0)
	goto L31
L31:
	;
	goto L28
L32:
	;
	F_errfinish(m, int32(517755), int32(311), int32(326459))
	mBase = m.M
	v202 = m.ExcPending
	if v202 != 0 {
		goto L11
	} else {
		goto L33
	}
L33:
	;
	goto L15
L34:
	;
	v218 = *(*int64)(unsafe.Add(mBase, uint32(v45)+104))
	*(*int64)(unsafe.Add(mBase, uint32(v54))) = v218
	goto L10
L35:
	;
	goto L7
L36:
	;
	if v261 != 0 {
		goto L37
	} else {
		goto L38
	}
L37:
	;
	F_errhidestmt(m)
	mBase = m.M
	v264 = m.ExcPending
	if v264 != 0 {
		goto L11
	} else {
		goto L40
	}
L38:
	;
	goto L39
L39:
	;
	v345 = int32(4543428)
	v347 = *(*int32)(unsafe.Add(mBase, _consts[14]))
	*(*int32)(unsafe.Add(mBase, _consts[14])) = v347 - int32(1)
	m.G0 = v45 + int32(112)
	v354 = *(*int64)(unsafe.Add(mBase, uint32(v32)+40))
	F_pgaio_io_update_state(m, l0, int32(7))
	mBase = m.M
	v357 = m.ExcPending
	if v357 != 0 {
		goto L11
	} else {
		goto L58
	}
L40:
	;
	F_errhidecontext(m)
	mBase = m.M
	v266 = m.ExcPending
	if v266 != 0 {
		goto L11
	} else {
		goto L41
	}
L41:
	;
	v268 = *(*int32)(unsafe.Add(mBase, _consts[590]))
	v269 = *(*int32)(unsafe.Add(mBase, uint32(v268)+24))
	goto L42
L42:
	;
	v274 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+2)))
	if base.Ui32(v274) <= base.Ui32(int32(2)) {
		goto L44
	} else {
		goto L45
	}
L43:
	;
	v283 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+1)))
	v288 = *(*int32)(unsafe.Add(mBase, uint32(v283<<(uint(int32(2))%32))+uint32(_consts[591])))
	v289 = *(*int32)(unsafe.Add(mBase, uint32(v288)+8))
	goto L47
L44:
	;
	v281 = *(*int32)(unsafe.Add(mBase, uint32(v274<<(uint(int32(2))%32))+uint32(_consts[592])))
	v282 = v281
	goto L46
L45:
	;
	v282 = int32(0)
	goto L46
L46:
	;
	goto L43
L47:
	;
	v291 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0))))
	if base.Ui32(v291) <= base.Ui32(int32(7)) {
		goto L49
	} else {
		goto L50
	}
L48:
	;
	v300 = *(*int32)(unsafe.Add(mBase, uint32(v54)))
	v304 = int32(base.Ui32(v300)>>(uint(int32(6))%32)) & int32(7)
	if base.Ui32(v304) <= base.Ui32(int32(4)) {
		goto L53
	} else {
		goto L54
	}
L49:
	;
	v298 = *(*int32)(unsafe.Add(mBase, uint32(v291<<(uint(int32(2))%32))+uint32(_consts[593])))
	v299 = v298
	goto L51
L50:
	;
	v299 = int32(0)
	goto L51
L51:
	;
	goto L48
L52:
	;
	v314 = *(*int32)(unsafe.Add(mBase, uint32(v54)))
	v315 = *(*int32)(unsafe.Add(mBase, uint32(v54)+4))
	v316 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	*(*int32)(unsafe.Add(mBase, uint32(v45)+32)) = v316
	*(*int32)(unsafe.Add(mBase, uint32(v45)+28)) = v315
	*(*int32)(unsafe.Add(mBase, uint32(v45)+16)) = v313
	*(*int32)(unsafe.Add(mBase, uint32(v45)+24)) = int32(base.Ui32(v314) >> (uint(int32(9)) % 32))
	*(*int32)(unsafe.Add(mBase, uint32(v45)+20)) = v314 & int32(63)
	*(*int32)(unsafe.Add(mBase, uint32(v45)+12)) = v299
	*(*int32)(unsafe.Add(mBase, uint32(v45)+8)) = v289
	*(*int32)(unsafe.Add(mBase, uint32(v45)+4)) = v282
	*(*int32)(unsafe.Add(mBase, uint32(v45))) = (l0 - v269) >> (uint(int32(7)) % 32)
	F_errmsg_internal(m, int32(499273), v45)
	mBase = m.M
	v332 = m.ExcPending
	if v332 != 0 {
		goto L11
	} else {
		goto L56
	}
L53:
	;
	v312 = *(*int32)(unsafe.Add(mBase, uint32(v304<<(uint(int32(2))%32))+uint32(_consts[595])))
	v313 = v312
	goto L55
L54:
	;
	v313 = int32(0)
	goto L55
L55:
	;
	goto L52
L56:
	;
	F_errfinish(m, int32(517755), int32(328), int32(326459))
	mBase = m.M
	v337 = m.ExcPending
	if v337 != 0 {
		goto L11
	} else {
		goto L57
	}
L57:
	;
	goto L39
L58:
	;
	v358 = *(*int32)(unsafe.Add(mBase, uint32(l0)+80))
	if v358 == int32(0) {
		goto L1
	} else {
		goto L59
	}
L59:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v358))) = v354
	v362 = *(*int32)(unsafe.Add(mBase, uint32(l0)+80))
	v363 = *(*int64)(unsafe.Add(mBase, uint32(l0)+104))
	*(*int64)(unsafe.Add(mBase, uint32(v362)+8)) = v363
	v365 = *(*int64)(unsafe.Add(mBase, uint32(l0)+120))
	*(*int64)(unsafe.Add(mBase, uint32(v362)+24)) = v365
	v367 = *(*int64)(unsafe.Add(mBase, uint32(l0)+112))
	*(*int64)(unsafe.Add(mBase, uint32(v362)+16)) = v367
	goto L1
L60:
	;
	if v400 != 0 {
		goto L61
	} else {
		goto L62
	}
L61:
	;
	F_errhidestmt(m)
	mBase = m.M
	v403 = m.ExcPending
	if v403 != 0 {
		goto L11
	} else {
		goto L64
	}
L62:
	;
	goto L63
L63:
	;
	v479 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0))))
	if v479 != int32(1) {
		goto L79
	} else {
		goto L80
	}
L64:
	;
	F_errhidecontext(m)
	mBase = m.M
	v405 = m.ExcPending
	if v405 != 0 {
		goto L11
	} else {
		goto L65
	}
L65:
	;
	v407 = *(*int32)(unsafe.Add(mBase, _consts[590]))
	v408 = *(*int32)(unsafe.Add(mBase, uint32(v407)+24))
	v411 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+2)))
	if base.Ui32(v411) <= base.Ui32(int32(2)) {
		goto L67
	} else {
		goto L68
	}
L66:
	;
	v420 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+1)))
	v425 = *(*int32)(unsafe.Add(mBase, uint32(v420<<(uint(int32(2))%32))+uint32(_consts[591])))
	v426 = *(*int32)(unsafe.Add(mBase, uint32(v425)+8))
	goto L70
L67:
	;
	v418 = *(*int32)(unsafe.Add(mBase, uint32(v411<<(uint(int32(2))%32))+uint32(_consts[592])))
	v419 = v418
	goto L69
L68:
	;
	v419 = int32(0)
	goto L69
L69:
	;
	goto L66
L70:
	;
	v427 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0))))
	if base.Ui32(v427) <= base.Ui32(int32(7)) {
		goto L71
	} else {
		goto L72
	}
L71:
	;
	v434 = *(*int32)(unsafe.Add(mBase, uint32(v427<<(uint(int32(2))%32))+uint32(_consts[593])))
	v435 = v434
	goto L73
L72:
	;
	v435 = v2
	goto L73
L73:
	;
	v436 = int32(7)
	v438 = *(*int32)(unsafe.Add(mBase, uint32(l0)+68))
	v442 = int32(base.Ui32(v438)>>(uint(int32(6))%32)) & v436
	if base.Ui32(v442) <= base.Ui32(int32(4)) {
		goto L74
	} else {
		goto L75
	}
L74:
	;
	v449 = *(*int32)(unsafe.Add(mBase, uint32(v442<<(uint(int32(2))%32))+uint32(_consts[595])))
	v450 = v449
	goto L76
L75:
	;
	v450 = v2
	goto L76
L76:
	;
	v451 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	*(*int32)(unsafe.Add(mBase, uint32(v32)+28)) = v451
	*(*int32)(unsafe.Add(mBase, uint32(v32)+24)) = int32(base.Ui32(v438) >> (uint(int32(9)) % 32))
	*(*int32)(unsafe.Add(mBase, uint32(v32)+20)) = v438 & int32(63)
	*(*int32)(unsafe.Add(mBase, uint32(v32)+16)) = v450
	*(*int32)(unsafe.Add(mBase, uint32(v32)+12)) = v435
	*(*int32)(unsafe.Add(mBase, uint32(v32)+8)) = v426
	*(*int32)(unsafe.Add(mBase, uint32(v32)+4)) = v419
	*(*int32)(unsafe.Add(mBase, uint32(v32))) = (l0 - v408) >> (uint(v436) % 32)
	F_errmsg_internal(m, int32(499555), v32)
	mBase = m.M
	v466 = m.ExcPending
	if v466 != 0 {
		goto L11
	} else {
		goto L77
	}
L77:
	;
	F_errfinish(m, int32(515733), int32(708), int32(300059))
	mBase = m.M
	v471 = m.ExcPending
	if v471 != 0 {
		goto L11
	} else {
		goto L78
	}
L78:
	;
	goto L63
L79:
	;
	v483 = *(*int32)(unsafe.Add(mBase, _consts[596]))
	v484 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
	v485 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
	*(*int32)(unsafe.Add(mBase, uint32(v484)+4)) = v485
	v487 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
	*(*int32)(unsafe.Add(mBase, uint32(v485))) = v487
	v489 = *(*int32)(unsafe.Add(mBase, uint32(v483)+160))
	*(*int32)(unsafe.Add(mBase, uint32(v483)+160)) = v489 - int32(1)
	goto L81
L80:
	;
	goto L81
L81:
	;
	v495 = *(*int32)(unsafe.Add(mBase, uint32(l0)+32))
	if v495 != 0 {
		goto L82
	} else {
		goto L83
	}
L82:
	;
	v497 = l0 + int32(36)
	v499 = *(*int32)(unsafe.Add(mBase, uint32(v497)))
	v500 = *(*int32)(unsafe.Add(mBase, uint32(v497)+4))
	*(*int32)(unsafe.Add(mBase, uint32(v499)+4)) = v500
	v502 = *(*int32)(unsafe.Add(mBase, uint32(v497)))
	*(*int32)(unsafe.Add(mBase, uint32(v500))) = v502
	goto L85
L83:
	;
	goto L84
L84:
	;
	v506 = *(*int64)(unsafe.Add(mBase, uint32(l0)+48))
	*(*int64)(unsafe.Add(mBase, uint32(l0)+48)) = v506 + int64(1)
	F_pgaio_io_update_state(m, l0, int32(0))
	mBase = m.M
	v512 = m.ExcPending
	if v512 != 0 {
		goto L11
	} else {
		goto L86
	}
L85:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+32)) = int32(0)
	goto L84
L86:
	;
	v513 = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(l0)+80)) = v513
	*(*uint8)(unsafe.Add(mBase, uint32(l0)+13)) = uint8(v513)
	*(*int32)(unsafe.Add(mBase, uint32(l0)+20)) = v513
	*(*int32)(unsafe.Add(mBase, uint32(l0)+1)) = v513
	v521 = *(*int32)(unsafe.Add(mBase, uint32(l0)+68))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+68)) = v521 & int32(-449)
	v526 = *(*int32)(unsafe.Add(mBase, _consts[596]))
	v528 = v526 + int32(4)
	v529 = *(*int32)(unsafe.Add(mBase, uint32(v526)+8))
	if v529 == v513 {
		goto L87
	} else {
		goto L88
	}
L87:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v526)+12)) = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v526)+4)) = v526 + int32(4)
	v537 = v528
	goto L89
L88:
	;
	v537 = v529
	goto L89
L89:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+24)) = v528
	*(*int32)(unsafe.Add(mBase, uint32(l0)+28)) = v537
	v541 = l0 + int32(24)
	*(*int32)(unsafe.Add(mBase, uint32(v537))) = v541
	*(*int32)(unsafe.Add(mBase, uint32(v526)+8)) = v541
	v544 = *(*int32)(unsafe.Add(mBase, uint32(v526)+12))
	v545 = int32(1)
	*(*int32)(unsafe.Add(mBase, uint32(v526)+12)) = v544 + v545
	v548 = int32(4543420)
	v550 = *(*int32)(unsafe.Add(mBase, _consts[412]))
	*(*int32)(unsafe.Add(mBase, _consts[412])) = v550 - v545
	m.G0 = v32 + int32(48)
	return
}
func F_pgaio_io_wait(m *base.Module, l0 int32, l1 int64) {
	mBase := m.M
	_ = mBase
	var v8 int32
	_ = v8
	var v10 int32
	_ = v10
	var v12 int32
	_ = v12
	var v14 int32
	_ = v14
	var v15 int32
	_ = v15
	var v16 int64
	_ = v16
	var v18 int32
	_ = v18
	var v26 int32
	_ = v26
	var v27 int64
	_ = v27
	var v30 int32
	_ = v30
	var v33 int32
	_ = v33
	var v43 int32
	_ = v43
	var v51 int32
	_ = v51
	var v56 int32
	_ = v56
	var v58 int32
	_ = v58
	var v59 int32
	_ = v59
	var v62 int32
	_ = v62
	var v66 int32
	_ = v66
	var v69 int32
	_ = v69
	var v70 int32
	_ = v70
	var v71 int64
	_ = v71
	var v86 int32
	_ = v86
	var v87 int32
	_ = v87
	var v88 int64
	_ = v88
	var v102 int32
	_ = v102
	var v104 int32
	_ = v104
	var v112 int32
	_ = v112
	var v113 int64
	_ = v113
	var v128 int32
	_ = v128
	var v130 int32
	_ = v130
	var v131 int32
	_ = v131
	var v133 int32
	_ = v133
	var v140 int32
	_ = v140
	var v141 int32
	_ = v141
	var v149 int32
	_ = v149
	var v154 int32
	_ = v154
	v8 = m.G0
	v10 = v8 - int32(32)
	m.G0 = v10
	v12 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0))))
	v14 = *(*int32)(unsafe.Add(mBase, _consts[74]))
	v15 = *(*int32)(unsafe.Add(mBase, uint32(l0)+16))
	v16 = *(*int64)(unsafe.Add(mBase, uint32(l0)+48))
	if v16 != l1 {
		goto L2
	} else {
		goto L3
	}
L1:
	;
	F_errstart_cold(m, int32(23), int32(0))
	mBase = m.M
	v128 = m.ExcPending
	if v128 != 0 {
		goto L13
	} else {
		goto L33
	}
L2:
	;
	m.G0 = v10 + int32(32)
	return
L3:
	;
	v18 = base.B2i32(v14 != v15)
	if base.B2i32(v18 == int32(0))&base.B2i32(base.Ui32(v12-int32(8)) < base.Ui32(int32(-4))) != 0 {
		goto L1
	} else {
		goto L4
	}
L4:
	;
	v26 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0))))
	v27 = *(*int64)(unsafe.Add(mBase, uint32(l0)+48))
	if v27 != l1 {
		goto L2
	} else {
		goto L5
	}
L5:
	;
	v30 = l0 + int32(56)
	v33 = v26
	goto L6
L6:
	;
	switch v33 & int32(255) {
	case 0, 1:
		goto L12
	case 2, 3, 5:
		goto L10
	case 4:
		goto L11
	case 6, 7:
		goto L9
	default:
		goto L8
	}
L7:
	;
	goto L2
L8:
	;
	v112 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0))))
	v113 = *(*int64)(unsafe.Add(mBase, uint32(l0)+48))
	if v113 == l1 {
		v33 = v112
		goto L6
	} else {
		goto L32
	}
L9:
	;
	if v14 != v15 {
		goto L2
	} else {
		goto L30
	}
L10:
	;
	F_ConditionVariablePrepareToSleep(m, v30)
	mBase = m.M
	v69 = m.ExcPending
	if v69 != 0 {
		goto L13
	} else {
		goto L20
	}
L11:
	;
	v58 = *(*int32)(unsafe.Add(mBase, _consts[597]))
	v59 = *(*int32)(unsafe.Add(mBase, uint32(v58)+24))
	if v59 == int32(0) {
		goto L10
	} else {
		goto L17
	}
L12:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v43 = m.ExcPending
	if v43 != 0 {
		goto L13
	} else {
		goto L14
	}
L13:
	;
	return
L14:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v10)+16)) = v33 & int32(255)
	F_errmsg_internal(m, int32(502148), v10+int32(16))
	mBase = m.M
	v51 = m.ExcPending
	if v51 != 0 {
		goto L13
	} else {
		goto L15
	}
L15:
	;
	F_errfinish(m, int32(515733), int32(610), int32(109691))
	mBase = m.M
	v56 = m.ExcPending
	if v56 != 0 {
		goto L13
	} else {
		goto L16
	}
L16:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L17:
	;
	v62 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+3)))
	if v62&int32(1) != 0 {
		goto L10
	} else {
		goto L18
	}
L18:
	;
	m.T0[v59].(func(*base.Module, int32, int64))(m, l0, l1)
	mBase = m.M
	v66 = m.ExcPending
	if v66 != 0 {
		goto L13
	} else {
		goto L19
	}
L19:
	;
	goto L8
L20:
	;
	v70 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0))))
	v71 = *(*int64)(unsafe.Add(mBase, uint32(l0)+48))
	if v71 != l1 {
		goto L21
	} else {
		goto L22
	}
L21:
	;
	F_ConditionVariableCancelSleep(m)
	mBase = m.M
	v102 = m.ExcPending
	if v102 != 0 {
		goto L13
	} else {
		goto L29
	}
L22:
	;
	if v70&int32(254) == int32(6) {
		goto L21
	} else {
		goto L23
	}
L23:
	;
	goto L24
L24:
	;
	F_ConditionVariableSleep(m, v30, int32(167772160))
	mBase = m.M
	v86 = m.ExcPending
	if v86 != 0 {
		goto L13
	} else {
		goto L26
	}
L25:
	;
	goto L21
L26:
	;
	v87 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0))))
	v88 = *(*int64)(unsafe.Add(mBase, uint32(l0)+48))
	if v88 != l1 {
		goto L21
	} else {
		goto L27
	}
L27:
	;
	if v87&int32(254) != int32(6) {
		goto L24
	} else {
		goto L28
	}
L28:
	;
	goto L25
L29:
	;
	goto L8
L30:
	;
	F_pgaio_io_reclaim(m, l0)
	mBase = m.M
	v104 = m.ExcPending
	if v104 != 0 {
		goto L13
	} else {
		goto L31
	}
L31:
	;
	goto L2
L32:
	;
	goto L7
L33:
	;
	v130 = *(*int32)(unsafe.Add(mBase, _consts[590]))
	v131 = *(*int32)(unsafe.Add(mBase, uint32(v130)+24))
	v133 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0))))
	if base.Ui32(v133) <= base.Ui32(int32(7)) {
		goto L35
	} else {
		goto L36
	}
L34:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v10)+4)) = v141
	*(*int32)(unsafe.Add(mBase, uint32(v10))) = (l0 - v131) >> (uint(int32(7)) % 32)
	F_errmsg_internal(m, int32(211959), v10)
	mBase = m.M
	v149 = m.ExcPending
	if v149 != 0 {
		goto L13
	} else {
		goto L38
	}
L35:
	;
	v140 = *(*int32)(unsafe.Add(mBase, uint32(v133<<(uint(int32(2))%32))+uint32(_consts[593])))
	v141 = v140
	goto L37
L36:
	;
	v141 = int32(0)
	goto L37
L37:
	;
	goto L34
L38:
	;
	F_errfinish(m, int32(515733), int32(597), int32(109691))
	mBase = m.M
	v154 = m.ExcPending
	if v154 != 0 {
		goto L13
	} else {
		goto L39
	}
L39:
	;
	base.Wasm_trap_unreachable()
	for {
	}
}
func F_pgaio_result_report(m *base.Module, l0 int32, l1 int32, l2 int32) {
	mBase := m.M
	_ = mBase
	var v8 int32
	_ = v8
	var v10 int32
	_ = v10
	var v12 int32
	_ = v12
	var v14 int32
	_ = v14
	var v16 int32
	_ = v16
	var v19 int32
	_ = v19
	var v20 int32
	_ = v20
	var v26 int32
	_ = v26
	var v27 int32
	_ = v27
	var v32 int32
	_ = v32
	var v37 int32
	_ = v37
	var v38 int64
	_ = v38
	var v43 int32
	_ = v43
	v8 = m.G0
	v10 = v8 - int32(16)
	m.G0 = v10
	v12 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v14 = v12 & int32(63)
	v16 = v14 << (uint(int32(3)) % 32)
	v19 = *(*int32)(unsafe.Add(mBase, uint32(v16)+uint32(_consts[589])))
	v20 = *(*int32)(unsafe.Add(mBase, uint32(v19)+12))
	if v20 == int32(0) {
		F_errstart_cold(m, int32(21), int32(0))
		mBase = m.M
		v26 = m.ExcPending
		if v26 != 0 {
			return
		} else {
			v27 = *(*int32)(unsafe.Add(mBase, uint32(v16)+uint32(_consts[594])))
			*(*int32)(unsafe.Add(mBase, uint32(v10)+4)) = v27
			*(*int32)(unsafe.Add(mBase, uint32(v10))) = v14
			F_errmsg_internal(m, int32(331476), v10)
			mBase = m.M
			v32 = m.ExcPending
			if v32 != 0 {
				return
			} else {
				F_errfinish(m, int32(517755), int32(183), int32(86394))
				mBase = m.M
				v37 = m.ExcPending
				if v37 != 0 {
					return
				} else {
					base.Wasm_trap_unreachable()
					for {
					}
				}
			}
		}
	} else {
		v38 = *(*int64)(unsafe.Add(mBase, uint32(l0)))
		*(*int64)(unsafe.Add(mBase, uint32(v10)+8)) = v38
		m.T0[v20].(func(*base.Module, int32, int32, int32))(m, v10+int32(8), l1, l2)
		mBase = m.M
		v43 = m.ExcPending
		if v43 != 0 {
			return
		} else {
			m.G0 = v10 + int32(16)
			return
		}
	}
}
func F_pgaio_result_status_string(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v9 int32
	_ = v9
	var v10 int32
	_ = v10
	if base.Ui32(l0) <= base.Ui32(int32(4)) {
		v9 = *(*int32)(unsafe.Add(mBase, uint32(l0<<(uint(int32(2))%32))+uint32(_consts[595])))
		v10 = v9
	} else {
		v10 = int32(0)
	}
	return v10
}
func F_pgaio_wref_wait(m *base.Module, l0 int32) {
	mBase := m.M
	_ = mBase
	var v3 int32
	_ = v3
	var v4 int32
	_ = v4
	var v5 int32
	_ = v5
	var v9 int64
	_ = v9
	var v10 int64
	_ = v10
	var v15 int32
	_ = v15
	v3 = *(*int32)(unsafe.Add(mBase, _consts[590]))
	v4 = *(*int32)(unsafe.Add(mBase, uint32(v3)+24))
	v5 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v9 = int64(*(*uint32)(unsafe.Add(mBase, uint32(l0)+8)))
	v10 = int64(*(*uint32)(unsafe.Add(mBase, uint32(l0)+4)))
	F_pgaio_io_wait(m, v4+v5<<(uint(int32(7))%32), v9|v10<<(uint(int64(32))%64))
	mBase = m.M
	v15 = m.ExcPending
	if v15 != 0 {
		return
	} else {
		return
	}
}
