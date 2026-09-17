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
	var v31 int32
	_ = v31
	var v33 int32
	_ = v33
	var v35 int32
	_ = v35
	var v37 int32
	_ = v37
	var v41 int32
	_ = v41
	var v44 int32
	_ = v44
	var v46 int32
	_ = v46
	var v48 int32
	_ = v48
	var v50 int32
	_ = v50
	var v55 int32
	_ = v55
	var v56 int64
	_ = v56
	var v60 int32
	_ = v60
	var v61 int32
	_ = v61
	var v62 int32
	_ = v62
	var v85 int32
	_ = v85
	var v88 int32
	_ = v88
	var v91 int32
	_ = v91
	var v114 int32
	_ = v114
	var v116 int32
	_ = v116
	var v118 int32
	_ = v118
	var v119 int32
	_ = v119
	var v120 int32
	_ = v120
	var v122 int32
	_ = v122
	var v125 int32
	_ = v125
	var v126 int32
	_ = v126
	var v128 int32
	_ = v128
	var v130 int32
	_ = v130
	var v132 int32
	_ = v132
	var v133 int32
	_ = v133
	var v137 int32
	_ = v137
	var v142 int32
	_ = v142
	var v144 int32
	_ = v144
	var v145 int32
	_ = v145
	var v148 int32
	_ = v148
	var v149 int32
	_ = v149
	var v150 int32
	_ = v150
	var v155 int32
	_ = v155
	var v157 int32
	_ = v157
	var v161 int32
	_ = v161
	var v166 int32
	_ = v166
	var v168 int32
	_ = v168
	var v180 int32
	_ = v180
	var v192 int32
	_ = v192
	var v197 int32
	_ = v197
	var v203 int32
	_ = v203
	var v204 int64
	_ = v204
	var v211 int32
	_ = v211
	var v212 int64
	_ = v212
	var v218 int32
	_ = v218
	var v220 int32
	_ = v220
	var v234 int32
	_ = v234
	var v237 int32
	_ = v237
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
	var v273 int32
	_ = v273
	var v278 int32
	_ = v278
	var v280 int32
	_ = v280
	var v281 int32
	_ = v281
	var v284 int32
	_ = v284
	var v285 int32
	_ = v285
	var v286 int32
	_ = v286
	var v291 int32
	_ = v291
	var v293 int32
	_ = v293
	var v297 int32
	_ = v297
	var v302 int32
	_ = v302
	var v304 int32
	_ = v304
	var v305 int32
	_ = v305
	var v321 int32
	_ = v321
	var v326 int32
	_ = v326
	var v332 int32
	_ = v332
	var v334 int32
	_ = v334
	var v341 int64
	_ = v341
	var v344 int32
	_ = v344
	var v345 int32
	_ = v345
	var v349 int32
	_ = v349
	var v350 int64
	_ = v350
	var v352 int64
	_ = v352
	var v354 int64
	_ = v354
	var v388 int32
	_ = v388
	var v389 int32
	_ = v389
	var v391 int32
	_ = v391
	var v393 int32
	_ = v393
	var v395 int32
	_ = v395
	var v396 int32
	_ = v396
	var v398 int32
	_ = v398
	var v403 int32
	_ = v403
	var v405 int32
	_ = v405
	var v406 int32
	_ = v406
	var v409 int32
	_ = v409
	var v410 int32
	_ = v410
	var v411 int32
	_ = v411
	var v416 int32
	_ = v416
	var v417 int32
	_ = v417
	var v418 int32
	_ = v418
	var v422 int32
	_ = v422
	var v427 int32
	_ = v427
	var v428 int32
	_ = v428
	var v429 int32
	_ = v429
	var v446 int32
	_ = v446
	var v451 int32
	_ = v451
	var v459 int32
	_ = v459
	var v463 int32
	_ = v463
	var v464 int32
	_ = v464
	var v465 int32
	_ = v465
	var v467 int32
	_ = v467
	var v469 int32
	_ = v469
	var v475 int32
	_ = v475
	var v477 int32
	_ = v477
	var v479 int32
	_ = v479
	var v480 int32
	_ = v480
	var v482 int32
	_ = v482
	var v486 int64
	_ = v486
	var v492 int32
	_ = v492
	var v493 int32
	_ = v493
	var v501 int32
	_ = v501
	var v506 int32
	_ = v506
	var v508 int32
	_ = v508
	var v509 int32
	_ = v509
	var v517 int32
	_ = v517
	var v521 int32
	_ = v521
	var v524 int32
	_ = v524
	var v525 int32
	_ = v525
	var v528 int32
	_ = v528
	var v530 int32
	_ = v530
	v2 = int32(0)
	v31 = m.G0
	v33 = v31 - int32(48)
	m.G0 = v33
	v35 = int32(_a_F_pgaio_io_reclaim_0)
	v37 = *(*int32)(unsafe.Add(mBase, _c_F_pgaio_io_reclaim[0]))
	*(*int32)(unsafe.Add(mBase, _c_F_pgaio_io_reclaim[0])) = v37 + int32(1)
	v41 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0))))
	if v41 != int32(6) {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	v388 = F_errstart(m, int32(11), int32(0))
	mBase = m.M
	v389 = m.ExcPending
	if v389 != 0 {
		goto L11
	} else {
		goto L60
	}
L2:
	;
	v44 = m.G0
	v46 = v44 - int32(112)
	m.G0 = v46
	v48 = int32(_a_F_pgaio_io_reclaim_1)
	v50 = *(*int32)(unsafe.Add(mBase, _c_F_pgaio_io_reclaim[1]))
	*(*int32)(unsafe.Add(mBase, _c_F_pgaio_io_reclaim[1])) = v50 + int32(1)
	v55 = v33 + int32(40)
	v56 = *(*int64)(unsafe.Add(mBase, uint32(l0)+68))
	*(*int64)(unsafe.Add(mBase, uint32(v55))) = v56
	v60 = base.I32_wrap_i64(int64(base.Ui64(v56) >> (uint(int64(32)) % 64)))
	v61 = base.I32_wrap_i64(v56)
	v62 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+4)))
	if v62 != 0 {
		goto L3
	} else {
		goto L4
	}
L3:
	;
	v85 = v62
	v88 = v61
	v91 = v60
	goto L6
L4:
	;
	v234 = v61
	v237 = v60
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
	v114 = v85 - int32(1)
	v116 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0+int32(5)+v114))))
	v118 = v116 << (uint(int32(3)) % 32)
	v119 = *(*int32)(unsafe.Add(mBase, uint32(v118)+uint32(_c_F_pgaio_io_reclaim[2])))
	v120 = *(*int32)(unsafe.Add(mBase, uint32(v119)+8))
	if v120 != 0 {
		goto L8
	} else {
		goto L9
	}
L7:
	;
	v234 = v218
	v237 = v220
	goto L5
L8:
	;
	v122 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v114+(l0+int32(9))))))
	v125 = F_errstart(m, int32(11), int32(0))
	mBase = m.M
	v126 = m.ExcPending
	if v126 != 0 {
		goto L11
	} else {
		goto L12
	}
L9:
	;
	v218 = v88
	v220 = v91
	goto L10
L10:
	;
	if base.Ui32(int32(1)) < base.Ui32(v85) {
		v85 = v114
		v88 = v218
		v91 = v220
		goto L6
	} else {
		goto L35
	}
L11:
	;
	return
L12:
	;
	if v125 != 0 {
		goto L13
	} else {
		goto L14
	}
L13:
	;
	F_errhidestmt(m)
	mBase = m.M
	v128 = m.ExcPending
	if v128 != 0 {
		goto L11
	} else {
		goto L16
	}
L14:
	;
	goto L15
L15:
	;
	v203 = *(*int32)(unsafe.Add(mBase, uint32(v119)+8))
	v204 = *(*int64)(unsafe.Add(mBase, uint32(v55)))
	*(*int64)(unsafe.Add(mBase, uint32(v46)+40)) = v204
	m.T0[v203].(func(*base.Module, int32, int32, int32, int32))(m, v46+int32(104), l0, v46+int32(40), v122)
	mBase = m.M
	v211 = m.ExcPending
	if v211 != 0 {
		goto L11
	} else {
		goto L34
	}
L16:
	;
	F_errhidecontext(m)
	mBase = m.M
	v130 = m.ExcPending
	if v130 != 0 {
		goto L11
	} else {
		goto L17
	}
L17:
	;
	v132 = *(*int32)(unsafe.Add(mBase, _c_F_pgaio_io_reclaim[3]))
	v133 = *(*int32)(unsafe.Add(mBase, uint32(v132)+24))
	goto L18
L18:
	;
	v137 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+2)))
	if base.Ui32(v137) <= base.Ui32(int32(2)) {
		goto L20
	} else {
		goto L21
	}
L19:
	;
	v145 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+1)))
	v148 = *(*int32)(unsafe.Add(mBase, uint32(v145<<(uint(int32(2))%32))+uint32(_c_F_pgaio_io_reclaim[4])))
	v149 = *(*int32)(unsafe.Add(mBase, uint32(v148)+8))
	goto L23
L20:
	;
	v142 = *(*int32)(unsafe.Add(mBase, uint32(v137<<(uint(int32(2))%32))+uint32(_c_F_pgaio_io_reclaim[5])))
	v144 = v142
	goto L22
L21:
	;
	v144 = int32(0)
	goto L22
L22:
	;
	goto L19
L23:
	;
	v150 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0))))
	if base.Ui32(v150) <= base.Ui32(int32(7)) {
		goto L25
	} else {
		goto L26
	}
L24:
	;
	v161 = int32(base.Ui32(v88)>>(uint(int32(6))%32)) & int32(7)
	if base.Ui32(v161) <= base.Ui32(int32(4)) {
		goto L29
	} else {
		goto L30
	}
L25:
	;
	v155 = *(*int32)(unsafe.Add(mBase, uint32(v150<<(uint(int32(2))%32))+uint32(_c_F_pgaio_io_reclaim[6])))
	v157 = v155
	goto L27
L26:
	;
	v157 = int32(0)
	goto L27
L27:
	;
	goto L24
L28:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v46+int32(92)))) = v91
	*(*int32)(unsafe.Add(mBase, uint32(v46+int32(88)))) = int32(base.Ui32(v88) >> (uint(int32(9)) % 32))
	*(*int32)(unsafe.Add(mBase, uint32(v46+int32(84)))) = v88 & int32(63)
	*(*int32)(unsafe.Add(mBase, uint32(v46+int32(80)))) = v168
	*(*int32)(unsafe.Add(mBase, uint32(v46+int32(76)))) = v122
	v180 = *(*int32)(unsafe.Add(mBase, uint32(v118)+uint32(_c_F_pgaio_io_reclaim[7])))
	*(*int32)(unsafe.Add(mBase, uint32(v46+int32(72)))) = v180
	*(*int32)(unsafe.Add(mBase, uint32(v46+int32(68)))) = v116
	*(*int32)(unsafe.Add(mBase, uint32(v46-int32(-64)))) = v85
	*(*int32)(unsafe.Add(mBase, uint32(v46)+60)) = v157
	*(*int32)(unsafe.Add(mBase, uint32(v46)+56)) = v149
	*(*int32)(unsafe.Add(mBase, uint32(v46)+52)) = v144
	*(*int32)(unsafe.Add(mBase, uint32(v46)+48)) = (l0 - v133) >> (uint(int32(7)) % 32)
	F_errmsg_internal(m, int32(_a_F_pgaio_io_reclaim_2), v46+int32(48))
	mBase = m.M
	v192 = m.ExcPending
	if v192 != 0 {
		goto L11
	} else {
		goto L32
	}
L29:
	;
	v166 = *(*int32)(unsafe.Add(mBase, uint32(v161<<(uint(int32(2))%32))+uint32(_c_F_pgaio_io_reclaim[8])))
	v168 = v166
	goto L31
L30:
	;
	v168 = int32(0)
	goto L31
L31:
	;
	goto L28
L32:
	;
	F_errfinish(m, int32(_a_F_pgaio_io_reclaim_3), int32(311), int32(_a_F_pgaio_io_reclaim_4))
	mBase = m.M
	v197 = m.ExcPending
	if v197 != 0 {
		goto L11
	} else {
		goto L33
	}
L33:
	;
	goto L15
L34:
	;
	v212 = *(*int64)(unsafe.Add(mBase, uint32(v46)+104))
	*(*int64)(unsafe.Add(mBase, uint32(v55))) = v212
	v218 = base.I32_wrap_i64(v212)
	v220 = base.I32_wrap_i64(int64(base.Ui64(v212) >> (uint(int64(32)) % 64)))
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
	v332 = int32(_a_F_pgaio_io_reclaim_1)
	v334 = *(*int32)(unsafe.Add(mBase, _c_F_pgaio_io_reclaim[1]))
	*(*int32)(unsafe.Add(mBase, _c_F_pgaio_io_reclaim[1])) = v334 - int32(1)
	m.G0 = v46 + int32(112)
	v341 = *(*int64)(unsafe.Add(mBase, uint32(v33)+40))
	F_pgaio_io_update_state(m, l0, int32(7))
	mBase = m.M
	v344 = m.ExcPending
	if v344 != 0 {
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
	v268 = *(*int32)(unsafe.Add(mBase, _c_F_pgaio_io_reclaim[3]))
	v269 = *(*int32)(unsafe.Add(mBase, uint32(v268)+24))
	goto L42
L42:
	;
	v273 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+2)))
	if base.Ui32(v273) <= base.Ui32(int32(2)) {
		goto L44
	} else {
		goto L45
	}
L43:
	;
	v281 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+1)))
	v284 = *(*int32)(unsafe.Add(mBase, uint32(v281<<(uint(int32(2))%32))+uint32(_c_F_pgaio_io_reclaim[4])))
	v285 = *(*int32)(unsafe.Add(mBase, uint32(v284)+8))
	goto L47
L44:
	;
	v278 = *(*int32)(unsafe.Add(mBase, uint32(v273<<(uint(int32(2))%32))+uint32(_c_F_pgaio_io_reclaim[5])))
	v280 = v278
	goto L46
L45:
	;
	v280 = int32(0)
	goto L46
L46:
	;
	goto L43
L47:
	;
	v286 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0))))
	if base.Ui32(v286) <= base.Ui32(int32(7)) {
		goto L49
	} else {
		goto L50
	}
L48:
	;
	v297 = int32(base.Ui32(v234)>>(uint(int32(6))%32)) & int32(7)
	if base.Ui32(v297) <= base.Ui32(int32(4)) {
		goto L53
	} else {
		goto L54
	}
L49:
	;
	v291 = *(*int32)(unsafe.Add(mBase, uint32(v286<<(uint(int32(2))%32))+uint32(_c_F_pgaio_io_reclaim[6])))
	v293 = v291
	goto L51
L50:
	;
	v293 = int32(0)
	goto L51
L51:
	;
	goto L48
L52:
	;
	v305 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	*(*int32)(unsafe.Add(mBase, uint32(v46)+32)) = v305
	*(*int32)(unsafe.Add(mBase, uint32(v46)+28)) = v237
	*(*int32)(unsafe.Add(mBase, uint32(v46)+24)) = int32(base.Ui32(v234) >> (uint(int32(9)) % 32))
	*(*int32)(unsafe.Add(mBase, uint32(v46)+20)) = v234 & int32(63)
	*(*int32)(unsafe.Add(mBase, uint32(v46)+16)) = v304
	*(*int32)(unsafe.Add(mBase, uint32(v46)+12)) = v293
	*(*int32)(unsafe.Add(mBase, uint32(v46)+8)) = v285
	*(*int32)(unsafe.Add(mBase, uint32(v46)+4)) = v280
	*(*int32)(unsafe.Add(mBase, uint32(v46))) = (l0 - v269) >> (uint(int32(7)) % 32)
	F_errmsg_internal(m, int32(_a_F_pgaio_io_reclaim_5), v46)
	mBase = m.M
	v321 = m.ExcPending
	if v321 != 0 {
		goto L11
	} else {
		goto L56
	}
L53:
	;
	v302 = *(*int32)(unsafe.Add(mBase, uint32(v297<<(uint(int32(2))%32))+uint32(_c_F_pgaio_io_reclaim[8])))
	v304 = v302
	goto L55
L54:
	;
	v304 = int32(0)
	goto L55
L55:
	;
	goto L52
L56:
	;
	F_errfinish(m, int32(_a_F_pgaio_io_reclaim_3), int32(328), int32(_a_F_pgaio_io_reclaim_4))
	mBase = m.M
	v326 = m.ExcPending
	if v326 != 0 {
		goto L11
	} else {
		goto L57
	}
L57:
	;
	goto L39
L58:
	;
	v345 = *(*int32)(unsafe.Add(mBase, uint32(l0)+80))
	if v345 == int32(0) {
		goto L1
	} else {
		goto L59
	}
L59:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v345))) = v341
	v349 = *(*int32)(unsafe.Add(mBase, uint32(l0)+80))
	v350 = *(*int64)(unsafe.Add(mBase, uint32(l0)+120))
	*(*int64)(unsafe.Add(mBase, uint32(v349)+24)) = v350
	v352 = *(*int64)(unsafe.Add(mBase, uint32(l0)+112))
	*(*int64)(unsafe.Add(mBase, uint32(v349)+16)) = v352
	v354 = *(*int64)(unsafe.Add(mBase, uint32(l0)+104))
	*(*int64)(unsafe.Add(mBase, uint32(v349)+8)) = v354
	goto L1
L60:
	;
	if v388 != 0 {
		goto L61
	} else {
		goto L62
	}
L61:
	;
	F_errhidestmt(m)
	mBase = m.M
	v391 = m.ExcPending
	if v391 != 0 {
		goto L11
	} else {
		goto L64
	}
L62:
	;
	goto L63
L63:
	;
	v459 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0))))
	if v459 != int32(1) {
		goto L79
	} else {
		goto L80
	}
L64:
	;
	F_errhidecontext(m)
	mBase = m.M
	v393 = m.ExcPending
	if v393 != 0 {
		goto L11
	} else {
		goto L65
	}
L65:
	;
	v395 = *(*int32)(unsafe.Add(mBase, _c_F_pgaio_io_reclaim[3]))
	v396 = *(*int32)(unsafe.Add(mBase, uint32(v395)+24))
	v398 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+2)))
	if base.Ui32(v398) <= base.Ui32(int32(2)) {
		goto L67
	} else {
		goto L68
	}
L66:
	;
	v406 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+1)))
	v409 = *(*int32)(unsafe.Add(mBase, uint32(v406<<(uint(int32(2))%32))+uint32(_c_F_pgaio_io_reclaim[4])))
	v410 = *(*int32)(unsafe.Add(mBase, uint32(v409)+8))
	goto L70
L67:
	;
	v403 = *(*int32)(unsafe.Add(mBase, uint32(v398<<(uint(int32(2))%32))+uint32(_c_F_pgaio_io_reclaim[5])))
	v405 = v403
	goto L69
L68:
	;
	v405 = int32(0)
	goto L69
L69:
	;
	goto L66
L70:
	;
	v411 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0))))
	if base.Ui32(v411) <= base.Ui32(int32(7)) {
		goto L71
	} else {
		goto L72
	}
L71:
	;
	v416 = *(*int32)(unsafe.Add(mBase, uint32(v411<<(uint(int32(2))%32))+uint32(_c_F_pgaio_io_reclaim[6])))
	v417 = v416
	goto L73
L72:
	;
	v417 = v2
	goto L73
L73:
	;
	v418 = *(*int32)(unsafe.Add(mBase, uint32(l0)+68))
	v422 = int32(base.Ui32(v418)>>(uint(int32(6))%32)) & int32(7)
	if base.Ui32(v422) <= base.Ui32(int32(4)) {
		goto L74
	} else {
		goto L75
	}
L74:
	;
	v427 = *(*int32)(unsafe.Add(mBase, uint32(v422<<(uint(int32(2))%32))+uint32(_c_F_pgaio_io_reclaim[8])))
	v428 = v427
	goto L76
L75:
	;
	v428 = v2
	goto L76
L76:
	;
	v429 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	*(*int32)(unsafe.Add(mBase, uint32(v33)+28)) = v429
	*(*int32)(unsafe.Add(mBase, uint32(v33)+24)) = int32(base.Ui32(v418) >> (uint(int32(9)) % 32))
	*(*int32)(unsafe.Add(mBase, uint32(v33)+20)) = v418 & int32(63)
	*(*int32)(unsafe.Add(mBase, uint32(v33)+16)) = v428
	*(*int32)(unsafe.Add(mBase, uint32(v33)+12)) = v417
	*(*int32)(unsafe.Add(mBase, uint32(v33)+8)) = v410
	*(*int32)(unsafe.Add(mBase, uint32(v33)+4)) = v405
	*(*int32)(unsafe.Add(mBase, uint32(v33))) = (l0 - v396) >> (uint(int32(7)) % 32)
	F_errmsg_internal(m, int32(_a_F_pgaio_io_reclaim_6), v33)
	mBase = m.M
	v446 = m.ExcPending
	if v446 != 0 {
		goto L11
	} else {
		goto L77
	}
L77:
	;
	F_errfinish(m, int32(_a_F_pgaio_io_reclaim_7), int32(708), int32(_a_F_pgaio_io_reclaim_8))
	mBase = m.M
	v451 = m.ExcPending
	if v451 != 0 {
		goto L11
	} else {
		goto L78
	}
L78:
	;
	goto L63
L79:
	;
	v463 = *(*int32)(unsafe.Add(mBase, _c_F_pgaio_io_reclaim[9]))
	v464 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
	v465 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
	*(*int32)(unsafe.Add(mBase, uint32(v464)+4)) = v465
	v467 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
	*(*int32)(unsafe.Add(mBase, uint32(v465))) = v467
	v469 = *(*int32)(unsafe.Add(mBase, uint32(v463)+160))
	*(*int32)(unsafe.Add(mBase, uint32(v463)+160)) = v469 - int32(1)
	goto L81
L80:
	;
	goto L81
L81:
	;
	v475 = *(*int32)(unsafe.Add(mBase, uint32(l0)+32))
	if v475 != 0 {
		goto L82
	} else {
		goto L83
	}
L82:
	;
	v477 = l0 + int32(36)
	v479 = *(*int32)(unsafe.Add(mBase, uint32(v477)))
	v480 = *(*int32)(unsafe.Add(mBase, uint32(v477)+4))
	*(*int32)(unsafe.Add(mBase, uint32(v479)+4)) = v480
	v482 = *(*int32)(unsafe.Add(mBase, uint32(v477)))
	*(*int32)(unsafe.Add(mBase, uint32(v480))) = v482
	goto L85
L83:
	;
	goto L84
L84:
	;
	v486 = *(*int64)(unsafe.Add(mBase, uint32(l0)+48))
	*(*int64)(unsafe.Add(mBase, uint32(l0)+48)) = v486 + int64(1)
	F_pgaio_io_update_state(m, l0, int32(0))
	mBase = m.M
	v492 = m.ExcPending
	if v492 != 0 {
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
	v493 = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(l0)+80)) = v493
	*(*uint8)(unsafe.Add(mBase, uint32(l0)+13)) = uint8(v493)
	*(*int32)(unsafe.Add(mBase, uint32(l0)+20)) = v493
	*(*int32)(unsafe.Add(mBase, uint32(l0)+1)) = v493
	v501 = *(*int32)(unsafe.Add(mBase, uint32(l0)+68))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+68)) = v501 & int32(-449)
	v506 = *(*int32)(unsafe.Add(mBase, _c_F_pgaio_io_reclaim[9]))
	v508 = v506 + int32(4)
	v509 = *(*int32)(unsafe.Add(mBase, uint32(v506)+8))
	if v509 == v493 {
		goto L87
	} else {
		goto L88
	}
L87:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v506)+12)) = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v506)+4)) = v506 + int32(4)
	v517 = v508
	goto L89
L88:
	;
	v517 = v509
	goto L89
L89:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+24)) = v508
	*(*int32)(unsafe.Add(mBase, uint32(l0)+28)) = v517
	v521 = l0 + int32(24)
	*(*int32)(unsafe.Add(mBase, uint32(v517))) = v521
	*(*int32)(unsafe.Add(mBase, uint32(v506)+8)) = v521
	v524 = *(*int32)(unsafe.Add(mBase, uint32(v506)+12))
	v525 = int32(1)
	*(*int32)(unsafe.Add(mBase, uint32(v506)+12)) = v524 + v525
	v528 = int32(_a_F_pgaio_io_reclaim_0)
	v530 = *(*int32)(unsafe.Add(mBase, _c_F_pgaio_io_reclaim[0]))
	*(*int32)(unsafe.Add(mBase, _c_F_pgaio_io_reclaim[0])) = v530 - v525
	m.G0 = v33 + int32(48)
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
	var v17 int32
	_ = v17
	var v18 int32
	_ = v18
	var v28 int32
	_ = v28
	var v30 int32
	_ = v30
	var v33 int32
	_ = v33
	var v41 int32
	_ = v41
	var v47 int32
	_ = v47
	var v52 int32
	_ = v52
	var v54 int32
	_ = v54
	var v55 int32
	_ = v55
	var v58 int32
	_ = v58
	var v62 int32
	_ = v62
	var v65 int32
	_ = v65
	var v66 int32
	_ = v66
	var v71 int64
	_ = v71
	var v83 int32
	_ = v83
	var v84 int32
	_ = v84
	var v85 int64
	_ = v85
	var v99 int32
	_ = v99
	var v101 int32
	_ = v101
	var v109 int32
	_ = v109
	var v110 int64
	_ = v110
	var v125 int32
	_ = v125
	var v127 int32
	_ = v127
	var v128 int32
	_ = v128
	var v129 int32
	_ = v129
	var v134 int32
	_ = v134
	var v136 int32
	_ = v136
	var v144 int32
	_ = v144
	var v149 int32
	_ = v149
	v8 = m.G0
	v10 = v8 - int32(32)
	m.G0 = v10
	v12 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0))))
	v14 = *(*int32)(unsafe.Add(mBase, _c_F_pgaio_io_wait[0]))
	v15 = *(*int32)(unsafe.Add(mBase, uint32(l0)+16))
	v16 = *(*int64)(unsafe.Add(mBase, uint32(l0)+48))
	v17 = base.B2i32(v16 != l1)
	if v16 != l1 {
		goto L2
	} else {
		goto L3
	}
L1:
	;
	F_errstart_cold(m, int32(23), int32(0))
	mBase = m.M
	v125 = m.ExcPending
	if v125 != 0 {
		goto L13
	} else {
		goto L32
	}
L2:
	;
	m.G0 = v10 + int32(32)
	return
L3:
	;
	v18 = base.B2i32(v14 != v15)
	if base.B2i32(v18 == int32(0))&base.B2i32(base.Ui32((v12-int32(8))&int32(255)) < base.Ui32(int32(252))) != 0 {
		goto L1
	} else {
		goto L4
	}
L4:
	;
	v28 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0))))
	if v16 != l1 {
		goto L2
	} else {
		goto L5
	}
L5:
	;
	v30 = l0 + int32(56)
	v33 = v28
	goto L6
L6:
	;
	switch v33 {
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
	v109 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0))))
	v110 = *(*int64)(unsafe.Add(mBase, uint32(l0)+48))
	if v110 == l1 {
		v33 = v109
		goto L6
	} else {
		goto L31
	}
L9:
	;
	if v14 != v15 {
		goto L2
	} else {
		goto L29
	}
L10:
	;
	F_ConditionVariablePrepareToSleep(m, v30)
	mBase = m.M
	v65 = m.ExcPending
	if v65 != 0 {
		goto L13
	} else {
		goto L20
	}
L11:
	;
	v54 = *(*int32)(unsafe.Add(mBase, _c_F_pgaio_io_wait[1]))
	v55 = *(*int32)(unsafe.Add(mBase, uint32(v54)+24))
	if v55 == int32(0) {
		goto L10
	} else {
		goto L17
	}
L12:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v41 = m.ExcPending
	if v41 != 0 {
		goto L13
	} else {
		goto L14
	}
L13:
	;
	return
L14:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v10)+16)) = v33
	F_errmsg_internal(m, int32(_a_F_pgaio_io_wait_0), v10+int32(16))
	mBase = m.M
	v47 = m.ExcPending
	if v47 != 0 {
		goto L13
	} else {
		goto L15
	}
L15:
	;
	F_errfinish(m, int32(_a_F_pgaio_io_wait_1), int32(610), int32(_a_F_pgaio_io_wait_2))
	mBase = m.M
	v52 = m.ExcPending
	if v52 != 0 {
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
	v58 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+3)))
	if v58&int32(1) != 0 {
		goto L10
	} else {
		goto L18
	}
L18:
	;
	m.T0[v55].(func(*base.Module, int32, int64))(m, l0, l1)
	mBase = m.M
	v62 = m.ExcPending
	if v62 != 0 {
		goto L13
	} else {
		goto L19
	}
L19:
	;
	goto L8
L20:
	;
	v66 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0))))
	v71 = *(*int64)(unsafe.Add(mBase, uint32(l0)+48))
	if base.B2i32(v66&int32(254) == int32(6))|base.B2i32(v71 != l1) != 0 {
		goto L21
	} else {
		goto L22
	}
L21:
	;
	F_ConditionVariableCancelSleep(m)
	mBase = m.M
	v99 = m.ExcPending
	if v99 != 0 {
		goto L13
	} else {
		goto L28
	}
L22:
	;
	goto L23
L23:
	;
	F_ConditionVariableSleep(m, v30, int32(167772160))
	mBase = m.M
	v83 = m.ExcPending
	if v83 != 0 {
		goto L13
	} else {
		goto L25
	}
L24:
	;
	goto L21
L25:
	;
	v84 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0))))
	v85 = *(*int64)(unsafe.Add(mBase, uint32(l0)+48))
	if v85 != l1 {
		goto L21
	} else {
		goto L26
	}
L26:
	;
	if v84&int32(254) != int32(6) {
		goto L23
	} else {
		goto L27
	}
L27:
	;
	goto L24
L28:
	;
	goto L8
L29:
	;
	F_pgaio_io_reclaim(m, l0)
	mBase = m.M
	v101 = m.ExcPending
	if v101 != 0 {
		goto L13
	} else {
		goto L30
	}
L30:
	;
	goto L2
L31:
	;
	goto L7
L32:
	;
	v127 = *(*int32)(unsafe.Add(mBase, _c_F_pgaio_io_wait[2]))
	v128 = *(*int32)(unsafe.Add(mBase, uint32(v127)+24))
	v129 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0))))
	if base.Ui32(v129) <= base.Ui32(int32(7)) {
		goto L34
	} else {
		goto L35
	}
L33:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v10)+4)) = v136
	*(*int32)(unsafe.Add(mBase, uint32(v10))) = (l0 - v128) >> (uint(int32(7)) % 32)
	F_errmsg_internal(m, int32(_a_F_pgaio_io_wait_3), v10)
	mBase = m.M
	v144 = m.ExcPending
	if v144 != 0 {
		goto L13
	} else {
		goto L37
	}
L34:
	;
	v134 = *(*int32)(unsafe.Add(mBase, uint32(v129<<(uint(int32(2))%32))+uint32(_c_F_pgaio_io_wait[3])))
	v136 = v134
	goto L36
L35:
	;
	v136 = int32(0)
	goto L36
L36:
	;
	goto L33
L37:
	;
	F_errfinish(m, int32(_a_F_pgaio_io_wait_1), int32(597), int32(_a_F_pgaio_io_wait_2))
	mBase = m.M
	v149 = m.ExcPending
	if v149 != 0 {
		goto L13
	} else {
		goto L38
	}
L38:
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
	var v17 int32
	_ = v17
	var v18 int32
	_ = v18
	var v24 int32
	_ = v24
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
	v17 = *(*int32)(unsafe.Add(mBase, uint32(v16)+uint32(_c_F_pgaio_result_report[0])))
	v18 = *(*int32)(unsafe.Add(mBase, uint32(v17)+12))
	if v18 == int32(0) {
		F_errstart_cold(m, int32(21), int32(0))
		mBase = m.M
		v24 = m.ExcPending
		if v24 != 0 {
			return
		} else {
			v27 = *(*int32)(unsafe.Add(mBase, uint32(v16)+uint32(_c_F_pgaio_result_report[1])))
			*(*int32)(unsafe.Add(mBase, uint32(v10)+4)) = v27
			*(*int32)(unsafe.Add(mBase, uint32(v10))) = v14
			F_errmsg_internal(m, int32(_a_F_pgaio_result_report_0), v10)
			mBase = m.M
			v32 = m.ExcPending
			if v32 != 0 {
				return
			} else {
				F_errfinish(m, int32(_a_F_pgaio_result_report_1), int32(183), int32(_a_F_pgaio_result_report_2))
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
		m.T0[v18].(func(*base.Module, int32, int32, int32))(m, v10+int32(8), l1, l2)
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
	var v6 int32
	_ = v6
	var v8 int32
	_ = v8
	if base.Ui32(l0) <= base.Ui32(int32(4)) {
		v6 = *(*int32)(unsafe.Add(mBase, uint32(l0<<(uint(int32(2))%32))+uint32(_c_F_pgaio_result_status_string[0])))
		v8 = v6
	} else {
		v8 = int32(0)
	}
	return v8
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
	v3 = *(*int32)(unsafe.Add(mBase, _c_F_pgaio_wref_wait[0]))
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
