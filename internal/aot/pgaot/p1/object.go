package p1

import (
	base "github.com/shibukawa/pgmem/internal/aot/pgaot/base"
	"unsafe"
)

func F_alloc_object(m *base.Module, l0 int32, l1 int32) int32 {
	mBase := m.M
	_ = mBase
	var v18 int32
	_ = v18
	var v20 int32
	_ = v20
	var v22 int32
	_ = v22
	var v25 int32
	_ = v25
	var v27 int32
	_ = v27
	var v29 int32
	_ = v29
	var v32 int32
	_ = v32
	var v33 int32
	_ = v33
	var v40 int32
	_ = v40
	var v44 int32
	_ = v44
	var v48 int32
	_ = v48
	var v49 int32
	_ = v49
	var v51 int32
	_ = v51
	var v52 int32
	_ = v52
	var v56 int32
	_ = v56
	var v60 int32
	_ = v60
	var v74 int32
	_ = v74
	var v75 int32
	_ = v75
	var v76 int32
	_ = v76
	var v81 int32
	_ = v81
	var v82 int32
	_ = v82
	var v84 int32
	_ = v84
	var v85 int32
	_ = v85
	var v89 int32
	_ = v89
	var v91 int32
	_ = v91
	var v94 int32
	_ = v94
	var v95 int32
	_ = v95
	var v96 int32
	_ = v96
	var v97 int32
	_ = v97
	var v98 int32
	_ = v98
	var v99 int32
	_ = v99
	var v102 int32
	_ = v102
	var v103 int32
	_ = v103
	var v107 int32
	_ = v107
	var v108 int32
	_ = v108
	var v109 int32
	_ = v109
	var v110 int32
	_ = v110
	var v111 int32
	_ = v111
	var v116 int32
	_ = v116
	var v117 int32
	_ = v117
	var v119 int32
	_ = v119
	var v120 int32
	_ = v120
	var v124 int32
	_ = v124
	var v128 int32
	_ = v128
	var v131 int32
	_ = v131
	var v132 int32
	_ = v132
	var v133 int32
	_ = v133
	var v134 int32
	_ = v134
	var v135 int32
	_ = v135
	var v136 int32
	_ = v136
	var v143 int32
	_ = v143
	var v146 int32
	_ = v146
	var v148 int32
	_ = v148
	var v154 int32
	_ = v154
	var v155 int32
	_ = v155
	var v156 int32
	_ = v156
	var v157 int32
	_ = v157
	var v162 int32
	_ = v162
	var v163 int32
	_ = v163
	var v165 int32
	_ = v165
	var v166 int32
	_ = v166
	var v170 int32
	_ = v170
	var v174 int32
	_ = v174
	var v177 int32
	_ = v177
	var v178 int32
	_ = v178
	var v179 int32
	_ = v179
	var v180 int32
	_ = v180
	var v181 int32
	_ = v181
	var v182 int32
	_ = v182
	var v184 int32
	_ = v184
	var v192 int32
	_ = v192
	var v200 int32
	_ = v200
	var v201 int32
	_ = v201
	var v206 int32
	_ = v206
	var v207 int32
	_ = v207
	var v208 int32
	_ = v208
	var v209 int32
	_ = v209
	var v214 int32
	_ = v214
	var v215 int32
	_ = v215
	var v217 int32
	_ = v217
	var v218 int32
	_ = v218
	var v222 int32
	_ = v222
	var v226 int32
	_ = v226
	var v229 int32
	_ = v229
	var v230 int32
	_ = v230
	var v231 int32
	_ = v231
	var v232 int32
	_ = v232
	var v233 int32
	_ = v233
	var v234 int32
	_ = v234
	var v247 int32
	_ = v247
	var v267 int32
	_ = v267
	var v268 int32
	_ = v268
	var v269 int32
	_ = v269
	var v274 int32
	_ = v274
	var v275 int32
	_ = v275
	var v281 int32
	_ = v281
	var v282 int32
	_ = v282
	var v286 int32
	_ = v286
	var v287 int32
	_ = v287
	var v288 int32
	_ = v288
	var v292 int32
	_ = v292
	var v293 int32
	_ = v293
	var v294 int32
	_ = v294
	var v295 int32
	_ = v295
	var v296 int32
	_ = v296
	var v297 int32
	_ = v297
	var v298 int32
	_ = v298
	var v302 int32
	_ = v302
	var v307 int32
	_ = v307
	var v308 int32
	_ = v308
	var v311 int32
	_ = v311
	var v312 int32
	_ = v312
	var v315 int32
	_ = v315
	var v319 int32
	_ = v319
	var v320 int32
	_ = v320
	var v327 int32
	_ = v327
	var v330 int32
	_ = v330
	var v331 int32
	_ = v331
	var v333 int32
	_ = v333
	var v335 int32
	_ = v335
	var v336 int32
	_ = v336
	var v343 int32
	_ = v343
	var v348 int32
	_ = v348
	var v358 int32
	_ = v358
	var v359 int32
	_ = v359
	var v360 int32
	_ = v360
	var v361 int32
	_ = v361
	var v367 int32
	_ = v367
	var v368 int32
	_ = v368
	var v374 int32
	_ = v374
	var v375 int32
	_ = v375
	var v381 int32
	_ = v381
	var v382 int32
	_ = v382
	var v388 int32
	_ = v388
	var v389 int32
	_ = v389
	var v391 int32
	_ = v391
	var v395 int32
	_ = v395
	var v414 int32
	_ = v414
	var v423 int32
	_ = v423
	var v429 int32
	_ = v429
	var v430 int32
	_ = v430
	var v431 int32
	_ = v431
	var v438 int32
	_ = v438
	var v441 int32
	_ = v441
	var v460 int32
	_ = v460
	var v463 int32
	_ = v463
	var v478 int32
	_ = v478
	var v479 int32
	_ = v479
	var v480 int32
	_ = v480
	var v485 int32
	_ = v485
	var v486 int32
	_ = v486
	var v488 int32
	_ = v488
	var v489 int32
	_ = v489
	var v493 int32
	_ = v493
	var v497 int32
	_ = v497
	var v499 int32
	_ = v499
	var v502 int32
	_ = v502
	var v503 int32
	_ = v503
	var v504 int32
	_ = v504
	var v505 int32
	_ = v505
	var v506 int32
	_ = v506
	var v507 int32
	_ = v507
	var v508 int32
	_ = v508
	var v509 int32
	_ = v509
	var v514 int32
	_ = v514
	var v515 int32
	_ = v515
	var v519 int32
	_ = v519
	var v520 int32
	_ = v520
	var v521 int32
	_ = v521
	var v522 int32
	_ = v522
	var v527 int32
	_ = v527
	var v528 int32
	_ = v528
	var v530 int32
	_ = v530
	var v531 int32
	_ = v531
	var v535 int32
	_ = v535
	var v537 int32
	_ = v537
	var v540 int32
	_ = v540
	var v541 int32
	_ = v541
	var v542 int32
	_ = v542
	var v543 int32
	_ = v543
	var v544 int32
	_ = v544
	var v545 int32
	_ = v545
	var v549 int32
	_ = v549
	var v551 int32
	_ = v551
	var v553 int32
	_ = v553
	var v557 int32
	_ = v557
	var v561 int32
	_ = v561
	var v563 int32
	_ = v563
	var v569 int32
	_ = v569
	var v570 int32
	_ = v570
	var v574 int32
	_ = v574
	var v588 int32
	_ = v588
	var v595 int32
	_ = v595
	var v603 int32
	_ = v603
	var v607 int32
	_ = v607
	var v612 int32
	_ = v612
	v18 = m.G0
	v20 = v18 - int32(16)
	m.G0 = v20
	v22 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v25 = v22 + l1<<(uint(int32(5))%32)
	v27 = v25 + int32(224)
	v29 = F_LWLockAcquire(m, v27, int32(0))
	mBase = m.M
	v32 = m.ExcPending
	if v32 != 0 {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	return int32(0)
L2:
	;
	v33 = *(*int32)(unsafe.Add(mBase, uint32(v25)+244))
	if v33 == int32(0) {
		goto L5
	} else {
		goto L6
	}
L3:
	;
	F_errstart_cold(m, int32(22), int32(0))
	mBase = m.M
	v603 = m.ExcPending
	if v603 != 0 {
		goto L1
	} else {
		goto L141
	}
L4:
	;
	v588 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	F_LWLockRelease(m, v588+l1<<(uint(int32(5))%32)+int32(224))
	mBase = m.M
	v595 = m.ExcPending
	if v595 != 0 {
		goto L1
	} else {
		goto L140
	}
L5:
	;
	v40 = int32(*(*uint16)(unsafe.Add(mBase, uint32(l1<<(uint(int32(1))%32))+uint32(_consts[1229]))))
	if l1 == int32(0) {
		goto L9
	} else {
		goto L10
	}
L6:
	;
	v463 = v33
	goto L7
L7:
	;
	v478 = *(*int32)(unsafe.Add(mBase, uint32(l0)+652))
	v479 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v480 = *(*int32)(unsafe.Add(mBase, uint32(v479)+1468))
	if v478 != v480 {
		goto L114
	} else {
		goto L115
	}
L8:
	;
	v51 = v25 + int32(248)
	v52 = *(*int32)(unsafe.Add(mBase, uint32(v51)))
	if v52 != 0 {
		goto L13
	} else {
		goto L14
	}
L9:
	;
	v44 = base.I32_div_u_s(int32(4096), v40)
	v49 = v44 - int32(1)
	goto L8
L10:
	;
	goto L11
L11:
	;
	v48 = base.I32_div_u_s(int32(65536), v40)
	v49 = v48
	goto L8
L12:
	;
	v460 = *(*int32)(unsafe.Add(mBase, uint32(v27)+20))
	v463 = v460
	goto L7
L13:
	;
	v56 = l0 + int32(12)
	v60 = v52
	goto L16
L14:
	;
	goto L15
L15:
	;
	v267 = F_transfer_first_span(m, l0, v27, int32(2), int32(1))
	mBase = m.M
	v268 = m.ExcPending
	if v268 != 0 {
		goto L1
	} else {
		goto L76
	}
L16:
	;
	v74 = *(*int32)(unsafe.Add(mBase, uint32(l0)+652))
	v75 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v76 = *(*int32)(unsafe.Add(mBase, uint32(v75)+1468))
	if v74 != v76 {
		goto L18
	} else {
		goto L19
	}
L17:
	;
	v247 = *(*int32)(unsafe.Add(mBase, uint32(v27)+20))
	if v247 != 0 {
		goto L12
	} else {
		goto L75
	}
L18:
	;
	v81 = F_LWLockAcquire(m, v75+int32(1476), int32(0))
	mBase = m.M
	v82 = m.ExcPending
	if v82 != 0 {
		goto L1
	} else {
		goto L21
	}
L19:
	;
	goto L20
L20:
	;
	v91 = int32(base.Ui32(v60) >> (uint(int32(27)) % 32))
	v94 = v56 + v91*int32(20)
	v95 = *(*int32)(unsafe.Add(mBase, uint32(v94)))
	if v95 != 0 {
		goto L24
	} else {
		goto L25
	}
L21:
	;
	F_check_for_freed_segments_locked(m, l0)
	mBase = m.M
	v84 = m.ExcPending
	if v84 != 0 {
		goto L1
	} else {
		goto L22
	}
L22:
	;
	v85 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	F_LWLockRelease(m, v85+int32(1476))
	mBase = m.M
	v89 = m.ExcPending
	if v89 != 0 {
		goto L1
	} else {
		goto L23
	}
L23:
	;
	goto L20
L24:
	;
	v99 = v95
	goto L26
L25:
	;
	v96 = F_get_segment_by_index(m, l0, v91)
	mBase = m.M
	v97 = m.ExcPending
	if v97 != 0 {
		goto L1
	} else {
		goto L27
	}
L26:
	;
	v102 = v99 + v60&int32(134217727)
	v103 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v102)+24)))
	v107 = base.I32_div_u_s((v49-v103)*int32(3), v49)
	v108 = *(*int32)(unsafe.Add(mBase, uint32(v102)+8))
	if v108 != 0 {
		goto L28
	} else {
		goto L29
	}
L27:
	;
	v98 = *(*int32)(unsafe.Add(mBase, uint32(v94)))
	v99 = v98
	goto L26
L28:
	;
	v109 = *(*int32)(unsafe.Add(mBase, uint32(l0)+652))
	v110 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v111 = *(*int32)(unsafe.Add(mBase, uint32(v110)+1468))
	if v109 != v111 {
		goto L31
	} else {
		goto L32
	}
L29:
	;
	v143 = int32(0)
	goto L30
L30:
	;
	if v107 <= int32(1) {
		goto L41
	} else {
		goto L42
	}
L31:
	;
	v116 = F_LWLockAcquire(m, v110+int32(1476), int32(0))
	mBase = m.M
	v117 = m.ExcPending
	if v117 != 0 {
		goto L1
	} else {
		goto L34
	}
L32:
	;
	goto L33
L33:
	;
	v128 = int32(base.Ui32(v108) >> (uint(int32(27)) % 32))
	v131 = v56 + v128*int32(20)
	v132 = *(*int32)(unsafe.Add(mBase, uint32(v131)))
	if v132 != 0 {
		goto L37
	} else {
		goto L38
	}
L34:
	;
	F_check_for_freed_segments_locked(m, l0)
	mBase = m.M
	v119 = m.ExcPending
	if v119 != 0 {
		goto L1
	} else {
		goto L35
	}
L35:
	;
	v120 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	F_LWLockRelease(m, v120+int32(1476))
	mBase = m.M
	v124 = m.ExcPending
	if v124 != 0 {
		goto L1
	} else {
		goto L36
	}
L36:
	;
	goto L33
L37:
	;
	v136 = v132
	goto L39
L38:
	;
	v133 = F_get_segment_by_index(m, l0, v128)
	mBase = m.M
	v134 = m.ExcPending
	if v134 != 0 {
		goto L1
	} else {
		goto L40
	}
L39:
	;
	v143 = v136 + v108&int32(134217727)
	goto L30
L40:
	;
	v135 = *(*int32)(unsafe.Add(mBase, uint32(v131)))
	v136 = v135
	goto L39
L41:
	;
	v146 = *(*int32)(unsafe.Add(mBase, uint32(v51)))
	if v60 == v146 {
		goto L46
	} else {
		goto L47
	}
L42:
	;
	goto L43
L43:
	;
	if v108 != 0 {
		v60 = v108
		goto L16
	} else {
		goto L74
	}
L44:
	;
	v200 = v25 + int32(240) + v107<<(uint(int32(2))%32)
	v201 = *(*int32)(unsafe.Add(mBase, uint32(v200)))
	*(*int32)(unsafe.Add(mBase, uint32(v102)+8)) = v201
	*(*int32)(unsafe.Add(mBase, uint32(v200))) = v60
	*(*int32)(unsafe.Add(mBase, uint32(v102)+4)) = int32(0)
	v206 = *(*int32)(unsafe.Add(mBase, uint32(v102)+8))
	if v206 != 0 {
		goto L61
	} else {
		goto L62
	}
L45:
	;
	v192 = *(*int32)(unsafe.Add(mBase, uint32(v102)+4))
	*(*int32)(unsafe.Add(mBase, uint32(v143)+4)) = v192
	goto L44
L46:
	;
	v148 = *(*int32)(unsafe.Add(mBase, uint32(v102)+8))
	*(*int32)(unsafe.Add(mBase, uint32(v51))) = v148
	if v143 == int32(0) {
		goto L44
	} else {
		goto L49
	}
L47:
	;
	goto L48
L48:
	;
	v154 = *(*int32)(unsafe.Add(mBase, uint32(v102)+4))
	v155 = *(*int32)(unsafe.Add(mBase, uint32(l0)+652))
	v156 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v157 = *(*int32)(unsafe.Add(mBase, uint32(v156)+1468))
	if v155 != v157 {
		goto L50
	} else {
		goto L51
	}
L49:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v143)+4)) = int32(0)
	goto L45
L50:
	;
	v162 = F_LWLockAcquire(m, v156+int32(1476), int32(0))
	mBase = m.M
	v163 = m.ExcPending
	if v163 != 0 {
		goto L1
	} else {
		goto L53
	}
L51:
	;
	goto L52
L52:
	;
	v174 = int32(base.Ui32(v154) >> (uint(int32(27)) % 32))
	v177 = v56 + v174*int32(20)
	v178 = *(*int32)(unsafe.Add(mBase, uint32(v177)))
	if v178 != 0 {
		goto L56
	} else {
		goto L57
	}
L53:
	;
	F_check_for_freed_segments_locked(m, l0)
	mBase = m.M
	v165 = m.ExcPending
	if v165 != 0 {
		goto L1
	} else {
		goto L54
	}
L54:
	;
	v166 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	F_LWLockRelease(m, v166+int32(1476))
	mBase = m.M
	v170 = m.ExcPending
	if v170 != 0 {
		goto L1
	} else {
		goto L55
	}
L55:
	;
	goto L52
L56:
	;
	v182 = v178
	goto L58
L57:
	;
	v179 = F_get_segment_by_index(m, l0, v174)
	mBase = m.M
	v180 = m.ExcPending
	if v180 != 0 {
		goto L1
	} else {
		goto L59
	}
L58:
	;
	v184 = *(*int32)(unsafe.Add(mBase, uint32(v102)+8))
	*(*int32)(unsafe.Add(mBase, uint32(v182+v154&int32(134217727))+8)) = v184
	if v143 == int32(0) {
		goto L44
	} else {
		goto L60
	}
L59:
	;
	v181 = *(*int32)(unsafe.Add(mBase, uint32(v177)))
	v182 = v181
	goto L58
L60:
	;
	goto L45
L61:
	;
	v207 = *(*int32)(unsafe.Add(mBase, uint32(l0)+652))
	v208 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v209 = *(*int32)(unsafe.Add(mBase, uint32(v208)+1468))
	if v207 != v209 {
		goto L64
	} else {
		goto L65
	}
L62:
	;
	goto L63
L63:
	;
	*(*uint16)(unsafe.Add(mBase, uint32(v102)+30)) = uint16(v107)
	goto L43
L64:
	;
	v214 = F_LWLockAcquire(m, v208+int32(1476), int32(0))
	mBase = m.M
	v215 = m.ExcPending
	if v215 != 0 {
		goto L1
	} else {
		goto L67
	}
L65:
	;
	goto L66
L66:
	;
	v226 = int32(base.Ui32(v206) >> (uint(int32(27)) % 32))
	v229 = v56 + v226*int32(20)
	v230 = *(*int32)(unsafe.Add(mBase, uint32(v229)))
	if v230 != 0 {
		goto L70
	} else {
		goto L71
	}
L67:
	;
	F_check_for_freed_segments_locked(m, l0)
	mBase = m.M
	v217 = m.ExcPending
	if v217 != 0 {
		goto L1
	} else {
		goto L68
	}
L68:
	;
	v218 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	F_LWLockRelease(m, v218+int32(1476))
	mBase = m.M
	v222 = m.ExcPending
	if v222 != 0 {
		goto L1
	} else {
		goto L69
	}
L69:
	;
	goto L66
L70:
	;
	v234 = v230
	goto L72
L71:
	;
	v231 = F_get_segment_by_index(m, l0, v226)
	mBase = m.M
	v232 = m.ExcPending
	if v232 != 0 {
		goto L1
	} else {
		goto L73
	}
L72:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v234+v206&int32(134217727))+4)) = v60
	goto L63
L73:
	;
	v233 = *(*int32)(unsafe.Add(mBase, uint32(v229)))
	v234 = v233
	goto L72
L74:
	;
	goto L17
L75:
	;
	goto L15
L76:
	;
	if v267 != 0 {
		goto L12
	} else {
		goto L77
	}
L77:
	;
	v269 = *(*int32)(unsafe.Add(mBase, uint32(v27)+20))
	if v269 == int32(0) {
		goto L78
	} else {
		goto L79
	}
L78:
	;
	v274 = F_transfer_first_span(m, l0, v27, int32(0), int32(1))
	mBase = m.M
	v275 = m.ExcPending
	if v275 != 0 {
		goto L1
	} else {
		goto L81
	}
L79:
	;
	goto L80
L80:
	;
	if l1 == int32(0) {
		goto L86
	} else {
		goto L87
	}
L81:
	;
	if v274 != 0 {
		goto L12
	} else {
		goto L82
	}
L82:
	;
	goto L80
L83:
	;
	v308 = *(*int32)(unsafe.Add(mBase, uint32(v307)+12))
	v311 = F_FreePageManagerGet(m, v308, v287, v20+int32(12))
	mBase = m.M
	v312 = m.ExcPending
	if v312 != 0 {
		goto L1
	} else {
		goto L97
	}
L84:
	;
	v574 = int32(0)
	goto L4
L85:
	;
	v288 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v292 = F_LWLockAcquire(m, v288+int32(1476), int32(0))
	mBase = m.M
	v293 = m.ExcPending
	if v293 != 0 {
		goto L1
	} else {
		goto L91
	}
L86:
	;
	v286 = int32(0)
	v287 = int32(1)
	goto L85
L87:
	;
	goto L88
L88:
	;
	v281 = F_alloc_object(m, l0, int32(0))
	mBase = m.M
	v282 = m.ExcPending
	if v282 != 0 {
		goto L1
	} else {
		goto L89
	}
L89:
	;
	if v281 == int32(0) {
		goto L84
	} else {
		goto L90
	}
L90:
	;
	v286 = v281
	v287 = int32(16)
	goto L85
L91:
	;
	v294 = F_get_best_segment(m, l0, v287)
	mBase = m.M
	v295 = m.ExcPending
	if v295 != 0 {
		goto L1
	} else {
		goto L92
	}
L92:
	;
	if v294 != 0 {
		v307 = v294
		goto L83
	} else {
		goto L93
	}
L93:
	;
	v296 = F_make_new_segment(m, l0, v287)
	mBase = m.M
	v297 = m.ExcPending
	if v297 != 0 {
		goto L1
	} else {
		goto L94
	}
L94:
	;
	if v296 != 0 {
		v307 = v296
		goto L83
	} else {
		goto L95
	}
L95:
	;
	v298 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	F_LWLockRelease(m, v298+int32(1476))
	mBase = m.M
	v302 = m.ExcPending
	if v302 != 0 {
		goto L1
	} else {
		goto L96
	}
L96:
	;
	goto L84
L97:
	;
	if v311 == int32(0) {
		goto L3
	} else {
		goto L98
	}
L98:
	;
	v315 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	F_LWLockRelease(m, v315+int32(1476))
	mBase = m.M
	v319 = m.ExcPending
	if v319 != 0 {
		goto L1
	} else {
		goto L99
	}
L99:
	;
	v320 = *(*int32)(unsafe.Add(mBase, uint32(v20)+12))
	v327 = base.I32_div_s(v307-l0-int32(8), int32(20))
	v330 = v320<<(uint(int32(12))%32) | v327<<(uint(int32(27))%32)
	if l1 != 0 {
		goto L100
	} else {
		goto L101
	}
L100:
	;
	v331 = v286
	goto L102
L101:
	;
	v331 = v330
	goto L102
L102:
	;
	F_init_span(m, l0, v331, v27, v330, v287, l1)
	mBase = m.M
	v333 = m.ExcPending
	if v333 != 0 {
		goto L1
	} else {
		goto L103
	}
L103:
	;
	v335 = v287 & int32(1)
	v336 = int32(0)
	if l1 != 0 {
		goto L104
	} else {
		goto L105
	}
L104:
	;
	v343 = v336
	v348 = int32(0)
	goto L107
L105:
	;
	v395 = v336
	goto L106
L106:
	;
	if v335 == int32(0) {
		goto L12
	} else {
		goto L110
	}
L107:
	;
	v358 = int32(2)
	v359 = v343 << (uint(v358) % 32)
	v360 = *(*int32)(unsafe.Add(mBase, uint32(v307)+16))
	v361 = *(*int32)(unsafe.Add(mBase, uint32(v20)+12))
	*(*int32)(unsafe.Add(mBase, uint32(v359+(v360+v361<<(uint(v358)%32))))) = v331
	v367 = *(*int32)(unsafe.Add(mBase, uint32(v307)+16))
	v368 = *(*int32)(unsafe.Add(mBase, uint32(v20)+12))
	*(*int32)(unsafe.Add(mBase, uint32(v367+v368<<(uint(v358)%32)+v359)+4)) = v331
	v374 = *(*int32)(unsafe.Add(mBase, uint32(v307)+16))
	v375 = *(*int32)(unsafe.Add(mBase, uint32(v20)+12))
	*(*int32)(unsafe.Add(mBase, uint32(v374+v375<<(uint(v358)%32)+v359)+8)) = v331
	v381 = *(*int32)(unsafe.Add(mBase, uint32(v307)+16))
	v382 = *(*int32)(unsafe.Add(mBase, uint32(v20)+12))
	*(*int32)(unsafe.Add(mBase, uint32(v381+v382<<(uint(v358)%32)+v359)+12)) = v331
	v388 = int32(4)
	v389 = v343 + v388
	v391 = v348 + v388
	if v391 != v287&int32(16) {
		v343 = v389
		v348 = v391
		goto L107
	} else {
		goto L109
	}
L108:
	;
	v395 = v389
	goto L106
L109:
	;
	goto L108
L110:
	;
	v414 = v395
	v423 = v336
	goto L111
L111:
	;
	v429 = *(*int32)(unsafe.Add(mBase, uint32(v307)+16))
	v430 = *(*int32)(unsafe.Add(mBase, uint32(v20)+12))
	v431 = int32(2)
	*(*int32)(unsafe.Add(mBase, uint32(v429+v430<<(uint(v431)%32)+v414<<(uint(v431)%32)))) = v331
	v438 = int32(1)
	v441 = v423 + v438
	if v441 != v335 {
		v414 = v414 + v438
		v423 = v441
		goto L111
	} else {
		goto L113
	}
L112:
	;
	goto L12
L113:
	;
	goto L112
L114:
	;
	v485 = F_LWLockAcquire(m, v479+int32(1476), int32(0))
	mBase = m.M
	v486 = m.ExcPending
	if v486 != 0 {
		goto L1
	} else {
		goto L117
	}
L115:
	;
	goto L116
L116:
	;
	v497 = l0 + int32(12)
	v499 = int32(base.Ui32(v463) >> (uint(int32(27)) % 32))
	v502 = v497 + v499*int32(20)
	v503 = *(*int32)(unsafe.Add(mBase, uint32(v502)))
	if v503 != 0 {
		goto L120
	} else {
		goto L121
	}
L117:
	;
	F_check_for_freed_segments_locked(m, l0)
	mBase = m.M
	v488 = m.ExcPending
	if v488 != 0 {
		goto L1
	} else {
		goto L118
	}
L118:
	;
	v489 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	F_LWLockRelease(m, v489+int32(1476))
	mBase = m.M
	v493 = m.ExcPending
	if v493 != 0 {
		goto L1
	} else {
		goto L119
	}
L119:
	;
	goto L116
L120:
	;
	v507 = v503
	goto L122
L121:
	;
	v504 = F_get_segment_by_index(m, l0, v499)
	mBase = m.M
	v505 = m.ExcPending
	if v505 != 0 {
		goto L1
	} else {
		goto L123
	}
L122:
	;
	v508 = v507 + v463&int32(134217727)
	v509 = *(*int32)(unsafe.Add(mBase, uint32(v508)+12))
	v514 = int32(*(*uint16)(unsafe.Add(mBase, uint32(l1<<(uint(int32(1))%32))+uint32(_consts[1229]))))
	v515 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v508)+26)))
	if v515 != int32(65535) {
		goto L125
	} else {
		goto L126
	}
L123:
	;
	v506 = *(*int32)(unsafe.Add(mBase, uint32(v502)))
	v507 = v506
	goto L122
L124:
	;
	v561 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v508)+24)))
	v563 = v561 - int32(1)
	*(*uint16)(unsafe.Add(mBase, uint32(v508)+24)) = uint16(v563)
	if v563&int32(65535) != 0 {
		v574 = v557
		goto L4
	} else {
		goto L138
	}
L125:
	;
	v519 = v514*v515 + v509
	v520 = *(*int32)(unsafe.Add(mBase, uint32(l0)+652))
	v521 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v522 = *(*int32)(unsafe.Add(mBase, uint32(v521)+1468))
	if v520 != v522 {
		goto L128
	} else {
		goto L129
	}
L126:
	;
	goto L127
L127:
	;
	v551 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v508)+22)))
	v553 = v551 + int32(1)
	*(*uint16)(unsafe.Add(mBase, uint32(v508)+22)) = uint16(v553)
	v557 = v551*v514 + v509
	goto L124
L128:
	;
	v527 = F_LWLockAcquire(m, v521+int32(1476), int32(0))
	mBase = m.M
	v528 = m.ExcPending
	if v528 != 0 {
		goto L1
	} else {
		goto L131
	}
L129:
	;
	goto L130
L130:
	;
	v537 = int32(base.Ui32(v519) >> (uint(int32(27)) % 32))
	v540 = v497 + v537*int32(20)
	v541 = *(*int32)(unsafe.Add(mBase, uint32(v540)))
	if v541 != 0 {
		goto L134
	} else {
		goto L135
	}
L131:
	;
	F_check_for_freed_segments_locked(m, l0)
	mBase = m.M
	v530 = m.ExcPending
	if v530 != 0 {
		goto L1
	} else {
		goto L132
	}
L132:
	;
	v531 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	F_LWLockRelease(m, v531+int32(1476))
	mBase = m.M
	v535 = m.ExcPending
	if v535 != 0 {
		goto L1
	} else {
		goto L133
	}
L133:
	;
	goto L130
L134:
	;
	v545 = v541
	goto L136
L135:
	;
	v542 = F_get_segment_by_index(m, l0, v537)
	mBase = m.M
	v543 = m.ExcPending
	if v543 != 0 {
		goto L1
	} else {
		goto L137
	}
L136:
	;
	v549 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v545+v519&int32(134217727)))))
	*(*uint16)(unsafe.Add(mBase, uint32(v508)+26)) = uint16(v549)
	v557 = v519
	goto L124
L137:
	;
	v544 = *(*int32)(unsafe.Add(mBase, uint32(v540)))
	v545 = v544
	goto L136
L138:
	;
	v569 = F_transfer_first_span(m, l0, v27, int32(1), int32(3))
	mBase = m.M
	v570 = m.ExcPending
	if v570 != 0 {
		goto L1
	} else {
		goto L139
	}
L139:
	;
	v574 = v557
	goto L4
L140:
	;
	m.G0 = v20 + int32(16)
	return v574
L141:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v20))) = v287
	F_errmsg_internal(m, int32(316304), v20)
	mBase = m.M
	v607 = m.ExcPending
	if v607 != 0 {
		goto L1
	} else {
		goto L142
	}
L142:
	;
	F_errfinish(m, int32(500062), int32(1719), int32(316279))
	mBase = m.M
	v612 = m.ExcPending
	if v612 != 0 {
		goto L1
	} else {
		goto L143
	}
L143:
	;
	base.Wasm_trap_unreachable()
	for {
	}
}
func F_free_object_addresses(m *base.Module, l0 int32) {
	mBase := m.M
	_ = mBase
	var v3 int32
	_ = v3
	var v5 int32
	_ = v5
	var v6 int32
	_ = v6
	var v8 int32
	_ = v8
	var v10 int32
	_ = v10
	v3 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	F_pfree(m, v3)
	mBase = m.M
	v5 = m.ExcPending
	if v5 != 0 {
		return
	} else {
		v6 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
		if v6 != 0 {
			F_pfree(m, v6)
			mBase = m.M
			v8 = m.ExcPending
			if v8 != 0 {
				return
			} else {
				F_pfree(m, l0)
				mBase = m.M
				v10 = m.ExcPending
				if v10 != 0 {
					return
				} else {
					return
				}
			}
		} else {
			F_pfree(m, l0)
			mBase = m.M
			v10 = m.ExcPending
			if v10 != 0 {
				return
			} else {
				return
			}
		}
	}
}
func F_get_object_address(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32, l5 int32) {
	mBase := m.M
	_ = mBase
	var v7 int32
	_ = v7
	var v22 int32
	_ = v22
	var v24 int32
	_ = v24
	var v31 int64
	_ = v31
	var v39 int32
	_ = v39
	var v46 int32
	_ = v46
	var v49 int32
	_ = v49
	var v50 int32
	_ = v50
	var v51 int64
	_ = v51
	var v54 int32
	_ = v54
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
	var v69 int32
	_ = v69
	var v72 int32
	_ = v72
	var v73 int32
	_ = v73
	var v81 int32
	_ = v81
	var v86 int32
	_ = v86
	var v87 int32
	_ = v87
	var v88 int32
	_ = v88
	var v94 int32
	_ = v94
	var v97 int32
	_ = v97
	var v98 int32
	_ = v98
	var v106 int32
	_ = v106
	var v111 int32
	_ = v111
	var v112 int32
	_ = v112
	var v113 int32
	_ = v113
	var v119 int32
	_ = v119
	var v122 int32
	_ = v122
	var v123 int32
	_ = v123
	var v131 int32
	_ = v131
	var v136 int32
	_ = v136
	var v137 int32
	_ = v137
	var v138 int32
	_ = v138
	var v144 int32
	_ = v144
	var v147 int32
	_ = v147
	var v148 int32
	_ = v148
	var v156 int32
	_ = v156
	var v161 int32
	_ = v161
	var v162 int32
	_ = v162
	var v163 int32
	_ = v163
	var v169 int32
	_ = v169
	var v172 int32
	_ = v172
	var v173 int32
	_ = v173
	var v181 int32
	_ = v181
	var v186 int32
	_ = v186
	var v187 int32
	_ = v187
	var v188 int32
	_ = v188
	var v194 int32
	_ = v194
	var v197 int32
	_ = v197
	var v198 int32
	_ = v198
	var v206 int32
	_ = v206
	var v211 int32
	_ = v211
	var v215 int32
	_ = v215
	var v221 int32
	_ = v221
	var v226 int32
	_ = v226
	var v227 int32
	_ = v227
	var v228 int32
	_ = v228
	var v233 int32
	_ = v233
	var v236 int32
	_ = v236
	var v242 int32
	_ = v242
	var v243 int32
	_ = v243
	var v246 int32
	_ = v246
	var v247 int32
	_ = v247
	var v248 int32
	_ = v248
	var v249 int32
	_ = v249
	var v250 int32
	_ = v250
	var v251 int32
	_ = v251
	var v252 int32
	_ = v252
	var v253 int32
	_ = v253
	var v254 int32
	_ = v254
	var v260 int32
	_ = v260
	var v261 int32
	_ = v261
	var v264 int32
	_ = v264
	var v265 int32
	_ = v265
	var v266 int32
	_ = v266
	var v270 int32
	_ = v270
	var v273 int32
	_ = v273
	var v279 int32
	_ = v279
	var v280 int32
	_ = v280
	var v283 int32
	_ = v283
	var v284 int32
	_ = v284
	var v285 int32
	_ = v285
	var v286 int32
	_ = v286
	var v287 int32
	_ = v287
	var v288 int32
	_ = v288
	var v289 int32
	_ = v289
	var v290 int32
	_ = v290
	var v291 int32
	_ = v291
	var v292 int32
	_ = v292
	var v295 int32
	_ = v295
	var v298 int32
	_ = v298
	var v299 int32
	_ = v299
	var v304 int32
	_ = v304
	var v305 int32
	_ = v305
	var v307 int32
	_ = v307
	var v308 int32
	_ = v308
	var v312 int32
	_ = v312
	var v315 int32
	_ = v315
	var v318 int32
	_ = v318
	var v324 int32
	_ = v324
	var v325 int32
	_ = v325
	var v328 int32
	_ = v328
	var v329 int32
	_ = v329
	var v330 int32
	_ = v330
	var v331 int32
	_ = v331
	var v333 int32
	_ = v333
	var v334 int32
	_ = v334
	var v335 int32
	_ = v335
	var v337 int32
	_ = v337
	var v346 int32
	_ = v346
	var v348 int32
	_ = v348
	var v350 int32
	_ = v350
	var v354 int32
	_ = v354
	var v355 int32
	_ = v355
	var v362 int32
	_ = v362
	var v369 int32
	_ = v369
	var v376 int32
	_ = v376
	var v377 int32
	_ = v377
	var v378 int32
	_ = v378
	var v379 int32
	_ = v379
	var v385 int32
	_ = v385
	var v388 int32
	_ = v388
	var v389 int32
	_ = v389
	var v390 int32
	_ = v390
	var v395 int32
	_ = v395
	var v400 int32
	_ = v400
	var v401 int32
	_ = v401
	var v402 int32
	_ = v402
	var v404 int32
	_ = v404
	var v405 int32
	_ = v405
	var v407 int32
	_ = v407
	var v410 int32
	_ = v410
	var v419 int32
	_ = v419
	var v420 int32
	_ = v420
	var v425 int32
	_ = v425
	var v427 int32
	_ = v427
	var v429 int32
	_ = v429
	var v433 int32
	_ = v433
	var v434 int32
	_ = v434
	var v437 int32
	_ = v437
	var v441 int32
	_ = v441
	var v448 int32
	_ = v448
	var v455 int32
	_ = v455
	var v456 int32
	_ = v456
	var v457 int32
	_ = v457
	var v458 int32
	_ = v458
	var v464 int32
	_ = v464
	var v467 int32
	_ = v467
	var v468 int32
	_ = v468
	var v469 int32
	_ = v469
	var v474 int32
	_ = v474
	var v479 int32
	_ = v479
	var v480 int32
	_ = v480
	var v481 int32
	_ = v481
	var v483 int32
	_ = v483
	var v484 int32
	_ = v484
	var v486 int32
	_ = v486
	var v489 int32
	_ = v489
	var v496 int32
	_ = v496
	var v502 int32
	_ = v502
	var v507 int32
	_ = v507
	var v510 int32
	_ = v510
	var v512 int32
	_ = v512
	var v515 int32
	_ = v515
	var v516 int32
	_ = v516
	var v522 int32
	_ = v522
	var v525 int32
	_ = v525
	var v526 int32
	_ = v526
	var v527 int32
	_ = v527
	var v532 int32
	_ = v532
	var v537 int32
	_ = v537
	var v538 int32
	_ = v538
	var v539 int32
	_ = v539
	var v541 int32
	_ = v541
	var v543 int32
	_ = v543
	var v545 int32
	_ = v545
	var v550 int32
	_ = v550
	var v556 int32
	_ = v556
	var v559 int32
	_ = v559
	var v562 int32
	_ = v562
	var v568 int32
	_ = v568
	var v570 int32
	_ = v570
	var v571 int32
	_ = v571
	var v572 int32
	_ = v572
	var v580 int32
	_ = v580
	var v582 int32
	_ = v582
	var v583 int32
	_ = v583
	var v585 int32
	_ = v585
	var v586 int32
	_ = v586
	var v588 int32
	_ = v588
	var v589 int32
	_ = v589
	var v591 int32
	_ = v591
	var v592 int32
	_ = v592
	var v594 int32
	_ = v594
	var v595 int32
	_ = v595
	var v597 int32
	_ = v597
	var v598 int32
	_ = v598
	var v600 int32
	_ = v600
	var v601 int32
	_ = v601
	var v603 int32
	_ = v603
	var v604 int32
	_ = v604
	var v606 int32
	_ = v606
	var v608 int32
	_ = v608
	var v611 int32
	_ = v611
	var v614 int32
	_ = v614
	var v615 int32
	_ = v615
	var v619 int32
	_ = v619
	var v622 int32
	_ = v622
	var v626 int32
	_ = v626
	var v631 int32
	_ = v631
	var v636 int32
	_ = v636
	var v637 int32
	_ = v637
	var v639 int32
	_ = v639
	var v640 int32
	_ = v640
	var v642 int32
	_ = v642
	var v643 int32
	_ = v643
	var v647 int32
	_ = v647
	var v653 int32
	_ = v653
	var v658 int32
	_ = v658
	var v661 int32
	_ = v661
	var v662 int32
	_ = v662
	var v663 int32
	_ = v663
	var v666 int32
	_ = v666
	var v667 int32
	_ = v667
	var v668 int32
	_ = v668
	var v671 int32
	_ = v671
	var v672 int32
	_ = v672
	var v673 int32
	_ = v673
	var v676 int32
	_ = v676
	var v677 int32
	_ = v677
	var v678 int32
	_ = v678
	var v681 int32
	_ = v681
	var v683 int32
	_ = v683
	var v690 int32
	_ = v690
	var v691 int32
	_ = v691
	var v693 int32
	_ = v693
	var v694 int32
	_ = v694
	var v695 int32
	_ = v695
	var v697 int32
	_ = v697
	var v698 int32
	_ = v698
	var v700 int32
	_ = v700
	var v701 int32
	_ = v701
	var v703 int32
	_ = v703
	var v704 int32
	_ = v704
	var v706 int32
	_ = v706
	var v709 int32
	_ = v709
	var v710 int32
	_ = v710
	var v713 int32
	_ = v713
	var v720 int32
	_ = v720
	var v735 int32
	_ = v735
	var v739 int32
	_ = v739
	var v741 int32
	_ = v741
	var v744 int32
	_ = v744
	var v746 int32
	_ = v746
	var v747 int32
	_ = v747
	var v750 int32
	_ = v750
	var v751 int32
	_ = v751
	var v796 int32
	_ = v796
	var v800 int32
	_ = v800
	var v803 int32
	_ = v803
	var v804 int32
	_ = v804
	var v805 int32
	_ = v805
	var v809 int32
	_ = v809
	var v814 int32
	_ = v814
	var v825 int32
	_ = v825
	var v839 int32
	_ = v839
	var v840 int32
	_ = v840
	var v841 int32
	_ = v841
	var v842 int32
	_ = v842
	var v843 int32
	_ = v843
	var v845 int32
	_ = v845
	var v846 int32
	_ = v846
	var v850 int32
	_ = v850
	var v851 int32
	_ = v851
	var v855 int32
	_ = v855
	var v861 int32
	_ = v861
	var v866 int32
	_ = v866
	var v868 int32
	_ = v868
	var v869 int32
	_ = v869
	var v870 int32
	_ = v870
	var v871 int32
	_ = v871
	var v872 int32
	_ = v872
	var v873 int32
	_ = v873
	var v879 int32
	_ = v879
	var v880 int32
	_ = v880
	var v884 int32
	_ = v884
	var v889 int32
	_ = v889
	var v890 int32
	_ = v890
	var v891 int32
	_ = v891
	var v892 int32
	_ = v892
	var v898 int32
	_ = v898
	var v899 int32
	_ = v899
	var v900 int32
	_ = v900
	var v901 int32
	_ = v901
	var v902 int32
	_ = v902
	var v903 int32
	_ = v903
	var v905 int32
	_ = v905
	var v908 int32
	_ = v908
	var v909 int32
	_ = v909
	var v910 int32
	_ = v910
	var v912 int32
	_ = v912
	var v914 int32
	_ = v914
	var v915 int32
	_ = v915
	var v919 int32
	_ = v919
	var v922 int32
	_ = v922
	var v928 int32
	_ = v928
	var v929 int32
	_ = v929
	var v932 int32
	_ = v932
	var v933 int32
	_ = v933
	var v934 int32
	_ = v934
	var v935 int32
	_ = v935
	var v936 int32
	_ = v936
	var v937 int32
	_ = v937
	var v938 int32
	_ = v938
	var v940 int32
	_ = v940
	var v941 int32
	_ = v941
	var v943 int32
	_ = v943
	var v944 int32
	_ = v944
	var v945 int32
	_ = v945
	var v950 int32
	_ = v950
	var v951 int32
	_ = v951
	var v954 int32
	_ = v954
	var v960 int32
	_ = v960
	var v961 int32
	_ = v961
	var v963 int32
	_ = v963
	var v964 int32
	_ = v964
	var v966 int32
	_ = v966
	var v972 int32
	_ = v972
	var v973 int32
	_ = v973
	var v975 int32
	_ = v975
	var v976 int32
	_ = v976
	var v977 int32
	_ = v977
	var v980 int32
	_ = v980
	var v982 int32
	_ = v982
	var v984 int32
	_ = v984
	var v987 int32
	_ = v987
	var v988 int32
	_ = v988
	var v991 int32
	_ = v991
	var v992 int32
	_ = v992
	var v996 int32
	_ = v996
	var v997 int32
	_ = v997
	var v1001 int32
	_ = v1001
	var v1006 int32
	_ = v1006
	var v1007 int32
	_ = v1007
	var v1008 int32
	_ = v1008
	var v1015 int32
	_ = v1015
	var v1016 int32
	_ = v1016
	var v1020 int32
	_ = v1020
	var v1023 int32
	_ = v1023
	var v1029 int32
	_ = v1029
	var v1034 int32
	_ = v1034
	var v1035 int32
	_ = v1035
	var v1036 int32
	_ = v1036
	var v1038 int32
	_ = v1038
	var v1039 int32
	_ = v1039
	var v1040 int32
	_ = v1040
	var v1041 int32
	_ = v1041
	var v1042 int32
	_ = v1042
	var v1045 int32
	_ = v1045
	var v1047 int32
	_ = v1047
	var v1050 int32
	_ = v1050
	var v1052 int32
	_ = v1052
	var v1053 int32
	_ = v1053
	var v1057 int32
	_ = v1057
	var v1060 int32
	_ = v1060
	var v1061 int32
	_ = v1061
	var v1062 int32
	_ = v1062
	var v1063 int32
	_ = v1063
	var v1064 int32
	_ = v1064
	var v1069 int32
	_ = v1069
	var v1074 int32
	_ = v1074
	var v1078 int32
	_ = v1078
	var v1079 int32
	_ = v1079
	var v1080 int32
	_ = v1080
	var v1082 int32
	_ = v1082
	var v1083 int32
	_ = v1083
	var v1084 int32
	_ = v1084
	var v1085 int32
	_ = v1085
	var v1086 int32
	_ = v1086
	var v1089 int32
	_ = v1089
	var v1090 int32
	_ = v1090
	var v1091 int32
	_ = v1091
	var v1094 int32
	_ = v1094
	var v1095 int32
	_ = v1095
	var v1096 int32
	_ = v1096
	var v1099 int32
	_ = v1099
	var v1100 int32
	_ = v1100
	var v1101 int32
	_ = v1101
	var v1104 int32
	_ = v1104
	var v1105 int32
	_ = v1105
	var v1106 int32
	_ = v1106
	var v1109 int32
	_ = v1109
	var v1110 int32
	_ = v1110
	var v1111 int32
	_ = v1111
	var v1112 int32
	_ = v1112
	var v1113 int32
	_ = v1113
	var v1114 int32
	_ = v1114
	var v1115 int32
	_ = v1115
	var v1116 int32
	_ = v1116
	var v1119 int32
	_ = v1119
	var v1120 int32
	_ = v1120
	var v1124 int32
	_ = v1124
	var v1125 int32
	_ = v1125
	var v1128 int32
	_ = v1128
	var v1129 int32
	_ = v1129
	var v1132 int32
	_ = v1132
	var v1139 int32
	_ = v1139
	var v1140 int32
	_ = v1140
	var v1146 int32
	_ = v1146
	var v1147 int32
	_ = v1147
	var v1154 int32
	_ = v1154
	var v1157 int32
	_ = v1157
	var v1164 int32
	_ = v1164
	var v1169 int32
	_ = v1169
	var v1170 int32
	_ = v1170
	var v1171 int32
	_ = v1171
	var v1173 int32
	_ = v1173
	var v1175 int32
	_ = v1175
	var v1177 int32
	_ = v1177
	var v1179 int32
	_ = v1179
	var v1180 int32
	_ = v1180
	var v1187 int32
	_ = v1187
	var v1190 int32
	_ = v1190
	var v1196 int32
	_ = v1196
	var v1201 int32
	_ = v1201
	var v1203 int32
	_ = v1203
	var v1204 int32
	_ = v1204
	var v1205 int32
	_ = v1205
	var v1212 int32
	_ = v1212
	var v1215 int32
	_ = v1215
	var v1222 int32
	_ = v1222
	var v1227 int32
	_ = v1227
	var v1228 int32
	_ = v1228
	var v1229 int32
	_ = v1229
	var v1231 int32
	_ = v1231
	var v1233 int32
	_ = v1233
	var v1234 int32
	_ = v1234
	var v1238 int32
	_ = v1238
	var v1242 int32
	_ = v1242
	var v1245 int32
	_ = v1245
	var v1246 int32
	_ = v1246
	var v1247 int32
	_ = v1247
	var v1248 int32
	_ = v1248
	var v1250 int32
	_ = v1250
	var v1251 int32
	_ = v1251
	var v1252 int32
	_ = v1252
	var v1253 int32
	_ = v1253
	var v1256 int32
	_ = v1256
	var v1257 int32
	_ = v1257
	var v1261 int32
	_ = v1261
	var v1262 int32
	_ = v1262
	var v1264 int32
	_ = v1264
	var v1265 int32
	_ = v1265
	var v1268 int32
	_ = v1268
	var v1273 int32
	_ = v1273
	var v1276 int32
	_ = v1276
	var v1277 int32
	_ = v1277
	var v1278 int32
	_ = v1278
	var v1279 int32
	_ = v1279
	var v1281 int32
	_ = v1281
	var v1282 int32
	_ = v1282
	var v1283 int32
	_ = v1283
	var v1284 int32
	_ = v1284
	var v1285 int32
	_ = v1285
	var v1286 int32
	_ = v1286
	var v1287 int32
	_ = v1287
	var v1289 int32
	_ = v1289
	var v1290 int32
	_ = v1290
	var v1291 int32
	_ = v1291
	var v1293 int32
	_ = v1293
	var v1294 int32
	_ = v1294
	var v1300 int32
	_ = v1300
	var v1304 int32
	_ = v1304
	var v1306 int32
	_ = v1306
	var v1308 int32
	_ = v1308
	var v1313 int32
	_ = v1313
	var v1316 int32
	_ = v1316
	var v1317 int32
	_ = v1317
	var v1319 int32
	_ = v1319
	var v1322 int32
	_ = v1322
	var v1323 int32
	_ = v1323
	var v1324 int32
	_ = v1324
	var v1325 int32
	_ = v1325
	var v1327 int32
	_ = v1327
	var v1328 int32
	_ = v1328
	var v1329 int32
	_ = v1329
	var v1331 int32
	_ = v1331
	var v1341 int32
	_ = v1341
	var v1344 int32
	_ = v1344
	var v1350 int32
	_ = v1350
	var v1361 int32
	_ = v1361
	var v1366 int32
	_ = v1366
	var v1368 int32
	_ = v1368
	var v1370 int32
	_ = v1370
	var v1371 int32
	_ = v1371
	var v1374 int32
	_ = v1374
	var v1375 int32
	_ = v1375
	var v1377 int32
	_ = v1377
	var v1379 int32
	_ = v1379
	var v1384 int32
	_ = v1384
	var v1385 int32
	_ = v1385
	var v1388 int32
	_ = v1388
	var v1390 int32
	_ = v1390
	var v1391 int32
	_ = v1391
	var v1394 int32
	_ = v1394
	var v1395 int32
	_ = v1395
	var v1397 int32
	_ = v1397
	var v1399 int32
	_ = v1399
	var v1407 int32
	_ = v1407
	var v1410 int32
	_ = v1410
	var v1417 int32
	_ = v1417
	var v1422 int32
	_ = v1422
	var v1423 int32
	_ = v1423
	var v1426 int32
	_ = v1426
	var v1427 int32
	_ = v1427
	var v1429 int32
	_ = v1429
	var v1430 int32
	_ = v1430
	var v1431 int32
	_ = v1431
	var v1435 int32
	_ = v1435
	var v1439 int32
	_ = v1439
	var v1444 int32
	_ = v1444
	var v1448 int32
	_ = v1448
	var v1449 int32
	_ = v1449
	var v1451 int32
	_ = v1451
	var v1452 int32
	_ = v1452
	var v1453 int32
	_ = v1453
	var v1454 int32
	_ = v1454
	var v1455 int32
	_ = v1455
	var v1456 int32
	_ = v1456
	var v1459 int32
	_ = v1459
	var v1460 int32
	_ = v1460
	var v1468 int32
	_ = v1468
	var v1482 int32
	_ = v1482
	var v1493 int32
	_ = v1493
	var v1507 int32
	_ = v1507
	var v1513 int32
	_ = v1513
	var v1516 int32
	_ = v1516
	var v1520 int32
	_ = v1520
	var v1525 int32
	_ = v1525
	var v1529 int32
	_ = v1529
	var v1532 int32
	_ = v1532
	var v1533 int32
	_ = v1533
	var v1534 int32
	_ = v1534
	var v1541 int32
	_ = v1541
	var v1546 int32
	_ = v1546
	var v1551 int32
	_ = v1551
	var v1554 int32
	_ = v1554
	var v1558 int32
	_ = v1558
	var v1563 int32
	_ = v1563
	var v1567 int32
	_ = v1567
	var v1570 int32
	_ = v1570
	var v1571 int32
	_ = v1571
	var v1572 int32
	_ = v1572
	var v1579 int32
	_ = v1579
	var v1584 int32
	_ = v1584
	var v1588 int32
	_ = v1588
	var v1591 int32
	_ = v1591
	var v1595 int32
	_ = v1595
	var v1600 int32
	_ = v1600
	var v1604 int32
	_ = v1604
	var v1607 int32
	_ = v1607
	var v1614 int32
	_ = v1614
	var v1619 int32
	_ = v1619
	var v1623 int32
	_ = v1623
	var v1626 int32
	_ = v1626
	var v1627 int32
	_ = v1627
	var v1636 int32
	_ = v1636
	var v1641 int32
	_ = v1641
	var v1649 int32
	_ = v1649
	var v1654 int32
	_ = v1654
	var v1655 int32
	_ = v1655
	var v1659 int32
	_ = v1659
	var v1663 int32
	_ = v1663
	var v1666 int32
	_ = v1666
	var v1669 int32
	_ = v1669
	var v1670 int32
	_ = v1670
	var v1671 int32
	_ = v1671
	var v1672 int32
	_ = v1672
	var v1673 int32
	_ = v1673
	var v1678 int32
	_ = v1678
	var v1679 int32
	_ = v1679
	var v1685 int32
	_ = v1685
	var v1688 int32
	_ = v1688
	var v1689 int32
	_ = v1689
	var v1690 int32
	_ = v1690
	var v1691 int32
	_ = v1691
	var v1692 int32
	_ = v1692
	var v1696 int32
	_ = v1696
	var v1697 int32
	_ = v1697
	var v1706 int32
	_ = v1706
	var v1711 int32
	_ = v1711
	var v1713 int32
	_ = v1713
	var v1714 int32
	_ = v1714
	var v1718 int32
	_ = v1718
	var v1719 int32
	_ = v1719
	var v1720 int32
	_ = v1720
	var v1721 int32
	_ = v1721
	var v1723 int32
	_ = v1723
	var v1725 int32
	_ = v1725
	var v1731 int32
	_ = v1731
	var v1734 int32
	_ = v1734
	var v1735 int32
	_ = v1735
	var v1736 int32
	_ = v1736
	var v1737 int32
	_ = v1737
	var v1738 int32
	_ = v1738
	var v1742 int32
	_ = v1742
	var v1743 int32
	_ = v1743
	var v1752 int32
	_ = v1752
	var v1757 int32
	_ = v1757
	var v1758 int32
	_ = v1758
	var v1759 int32
	_ = v1759
	var v1765 int32
	_ = v1765
	var v1768 int32
	_ = v1768
	var v1769 int32
	_ = v1769
	var v1770 int32
	_ = v1770
	var v1779 int32
	_ = v1779
	var v1781 int32
	_ = v1781
	var v1785 int32
	_ = v1785
	var v1794 int32
	_ = v1794
	var v1795 int32
	_ = v1795
	var v1796 int32
	_ = v1796
	var v1815 int32
	_ = v1815
	var v1821 int32
	_ = v1821
	var v1833 int32
	_ = v1833
	var v1845 int32
	_ = v1845
	var v1861 int32
	_ = v1861
	var v1889 int32
	_ = v1889
	var v1891 int32
	_ = v1891
	var v1893 int32
	_ = v1893
	var v1898 int32
	_ = v1898
	var v1910 int32
	_ = v1910
	var v1922 int32
	_ = v1922
	var v1938 int32
	_ = v1938
	var v1966 int32
	_ = v1966
	var v1968 int32
	_ = v1968
	var v1970 int32
	_ = v1970
	var v1972 int64
	_ = v1972
	var v1974 int32
	_ = v1974
	v7 = int32(0)
	v22 = m.G0
	v24 = v22 - int32(448)
	m.G0 = v24
	*(*int64)(unsafe.Add(mBase, uint32(l0))) = int64(0)
	*(*int32)(unsafe.Add(mBase, uint32(l0)+8)) = v7
	v31 = *(*int64)(unsafe.Add(mBase, _consts[315]))
	v39 = v7
	v46 = v7
	v49 = v7
	v50 = v7
	v51 = v31
	goto L2
L1:
	;
	m.G0 = v24 + int32(448)
	return
L2:
	;
	switch l1 {
	case 0, 9, 14, 15, 16, 17, 21, 27, 30, 33, 36, 38, 42:
		goto L42
	case 1, 19, 29, 34:
		goto L41
	case 2, 3:
		goto L36
	case 4, 6:
		goto L45
	case 5:
		goto L34
	case 7:
		goto L39
	case 8:
		goto L38
	case 10:
		goto L44
	case 11:
		goto L25
	case 12, 49:
		goto L23
	case 13:
		goto L21
	case 18, 20, 23, 37, 41, 51:
		goto L46
	case 22:
		goto L35
	case 24, 26:
		goto L37
	case 25:
		goto L40
	case 28, 35, 40, 44:
		goto L43
	case 31:
		goto L27
	case 32:
		goto L26
	case 39:
		goto L24
	case 43:
		goto L33
	case 45:
		goto L29
	case 46:
		goto L31
	case 47:
		goto L32
	case 48:
		goto L30
	case 50:
		goto L28
	default:
		v1431 = v39
		goto L22
	}
L3:
	;
	if l3 == int32(0) {
		goto L1
	} else {
		goto L590
	}
L4:
	;
	if v1794 == int32(0) {
		goto L1
	} else {
		goto L511
	}
L5:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v1779
	v1785 = int32(1259)
	*(*int32)(unsafe.Add(mBase, uint32(l0))) = v1785
	v1794 = v1779
	v1795 = v1785
	v1796 = v1781
	goto L4
L6:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v1768
	*(*int32)(unsafe.Add(mBase, uint32(l0))) = v1769
	v1794 = v1768
	v1795 = v1769
	v1796 = v1770
	goto L4
L7:
	;
	v1765 = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(l0)+8)) = v1765
	v1768 = v1758
	v1769 = v1759
	v1770 = v1765
	goto L6
L8:
	;
	v1673 = base.I32_extend16_s(v928)
	if l1 == int32(2) {
		goto L483
	} else {
		goto L484
	}
L9:
	;
	v1666 = int32(0)
	v1669 = v1666
	v1670 = v929
	v1671 = v1666
	v1672 = v1666
	goto L8
L10:
	;
	v1659 = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(l0)+8)) = v1659
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v1655
	v1663 = int32(826)
	*(*int32)(unsafe.Add(mBase, uint32(l0))) = v1663
	v1794 = v1655
	v1795 = v1663
	v1796 = v1659
	goto L4
L11:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v24)+408)) = v1368
	*(*int32)(unsafe.Add(mBase, uint32(v24)+404)) = v1324
	*(*int32)(unsafe.Add(mBase, uint32(v24)+400)) = v1325
	F_errmsg(m, int32(69851), v24+int32(400))
	mBase = m.M
	v1649 = m.ExcPending
	if v1649 != 0 {
		goto L47
	} else {
		goto L479
	}
L12:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v1623 = m.ExcPending
	if v1623 != 0 {
		goto L47
	} else {
		goto L475
	}
L13:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v1604 = m.ExcPending
	if v1604 != 0 {
		goto L47
	} else {
		goto L471
	}
L14:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v1588 = m.ExcPending
	if v1588 != 0 {
		goto L47
	} else {
		goto L467
	}
L15:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v1567 = m.ExcPending
	if v1567 != 0 {
		goto L47
	} else {
		goto L462
	}
L16:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v1551 = m.ExcPending
	if v1551 != 0 {
		goto L47
	} else {
		goto L458
	}
L17:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v1529 = m.ExcPending
	if v1529 != 0 {
		goto L47
	} else {
		goto L453
	}
L18:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v1513 = m.ExcPending
	if v1513 != 0 {
		goto L47
	} else {
		goto L449
	}
L19:
	;
	v1507 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v1794 = v1507
	v1795 = v1493
	v1796 = int32(0)
	goto L4
L20:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+8)) = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v1482
	v1493 = v1468
	goto L19
L21:
	;
	v1448 = *(*int32)(unsafe.Add(mBase, uint32(l2)+12))
	v1449 = *(*int32)(unsafe.Add(mBase, uint32(v1448)))
	F_get_object_address_type(m, v24+int32(436), int32(12), v1449, l5)
	mBase = m.M
	v1451 = m.ExcPending
	if v1451 != 0 {
		goto L47
	} else {
		goto L447
	}
L22:
	;
	if v1431 != 0 {
		v1493 = v1431
		goto L19
	} else {
		goto L443
	}
L23:
	;
	F_get_object_address_type(m, l0, l1, l2, l5)
	mBase = m.M
	v1429 = m.ExcPending
	if v1429 != 0 {
		goto L47
	} else {
		goto L442
	}
L24:
	;
	v1423 = int32(3381)
	*(*int32)(unsafe.Add(mBase, uint32(l0))) = v1423
	v1426 = F_get_statistics_object_oid(m, l2, l5)
	mBase = m.M
	v1427 = m.ExcPending
	if v1427 != 0 {
		goto L47
	} else {
		goto L441
	}
L25:
	;
	v1316 = *(*int32)(unsafe.Add(mBase, uint32(l2)+12))
	v1317 = *(*int32)(unsafe.Add(mBase, uint32(v1316)+4))
	v1319 = *(*int32)(unsafe.Add(mBase, uint32(l2)+4))
	if int32(3) <= v1319 {
		goto L407
	} else {
		goto L408
	}
L26:
	;
	v1276 = *(*int32)(unsafe.Add(mBase, uint32(l2)+12))
	v1277 = *(*int32)(unsafe.Add(mBase, uint32(v1276)))
	v1278 = F_makeRangeVarFromNameList(m, v1277)
	mBase = m.M
	v1279 = m.ExcPending
	if v1279 != 0 {
		goto L47
	} else {
		goto L394
	}
L27:
	;
	v1245 = *(*int32)(unsafe.Add(mBase, uint32(l2)+12))
	v1246 = *(*int32)(unsafe.Add(mBase, uint32(v1245)+4))
	v1247 = *(*int32)(unsafe.Add(mBase, uint32(v1246)+4))
	v1248 = int32(0)
	v1250 = *(*int32)(unsafe.Add(mBase, uint32(v1245)))
	v1251 = *(*int32)(unsafe.Add(mBase, uint32(v1250)+4))
	v1252 = F_get_namespace_oid(m, v1251, l5)
	mBase = m.M
	v1253 = m.ExcPending
	if v1253 != 0 {
		goto L47
	} else {
		goto L386
	}
L28:
	;
	v1111 = *(*int32)(unsafe.Add(mBase, uint32(l2)+12))
	v1112 = *(*int32)(unsafe.Add(mBase, uint32(v1111)+4))
	v1113 = *(*int32)(unsafe.Add(mBase, uint32(v1112)+4))
	v1114 = *(*int32)(unsafe.Add(mBase, uint32(v1111)))
	v1115 = *(*int32)(unsafe.Add(mBase, uint32(v1114)+4))
	v1116 = int32(491360)
	v1119 = int32(*(*uint8)(unsafe.Add(mBase, _consts[324])))
	v1120 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1115))))
	if v1120 == int32(0) {
		v1139 = v1119
		v1140 = v1120
		goto L340
	} else {
		goto L341
	}
L29:
	;
	v1106 = int32(3602)
	*(*int32)(unsafe.Add(mBase, uint32(l0))) = v1106
	v1109 = F_get_ts_config_oid(m, l2, l5)
	mBase = m.M
	v1110 = m.ExcPending
	if v1110 != 0 {
		goto L47
	} else {
		goto L336
	}
L30:
	;
	v1101 = int32(3764)
	*(*int32)(unsafe.Add(mBase, uint32(l0))) = v1101
	v1104 = F_get_ts_template_oid(m, l2, l5)
	mBase = m.M
	v1105 = m.ExcPending
	if v1105 != 0 {
		goto L47
	} else {
		goto L335
	}
L31:
	;
	v1096 = int32(3600)
	*(*int32)(unsafe.Add(mBase, uint32(l0))) = v1096
	v1099 = F_get_ts_dict_oid(m, l2, l5)
	mBase = m.M
	v1100 = m.ExcPending
	if v1100 != 0 {
		goto L47
	} else {
		goto L334
	}
L32:
	;
	v1091 = int32(3601)
	*(*int32)(unsafe.Add(mBase, uint32(l0))) = v1091
	v1094 = F_get_ts_parser_oid(m, l2, l5)
	mBase = m.M
	v1095 = m.ExcPending
	if v1095 != 0 {
		goto L47
	} else {
		goto L333
	}
L33:
	;
	v1078 = *(*int32)(unsafe.Add(mBase, uint32(l2)+12))
	v1079 = *(*int32)(unsafe.Add(mBase, uint32(v1078)+4))
	v1080 = *(*int32)(unsafe.Add(mBase, uint32(v1079)+4))
	v1082 = *(*int32)(unsafe.Add(mBase, uint32(v1078)))
	v1083 = F_LookupTypeNameOid(m, v1082, l5)
	mBase = m.M
	v1084 = m.ExcPending
	if v1084 != 0 {
		goto L47
	} else {
		goto L330
	}
L34:
	;
	v1035 = *(*int32)(unsafe.Add(mBase, uint32(l2)+12))
	v1036 = *(*int32)(unsafe.Add(mBase, uint32(v1035)+4))
	v1038 = *(*int32)(unsafe.Add(mBase, uint32(v1035)))
	v1039 = F_LookupTypeNameOid(m, v1038, l5)
	mBase = m.M
	v1040 = m.ExcPending
	if v1040 != 0 {
		goto L47
	} else {
		goto L318
	}
L35:
	;
	v977 = int32(2613)
	*(*int32)(unsafe.Add(mBase, uint32(l0))) = v977
	v980 = m.G0
	v982 = v980 - int32(16)
	m.G0 = v982
	v984 = *(*int32)(unsafe.Add(mBase, uint32(l2)))
	switch v984 - int32(465) {
	case 0:
		goto L304
	case 1:
		goto L306
	default:
		goto L305
	}
L36:
	;
	v870 = *(*int32)(unsafe.Add(mBase, uint32(l2)+12))
	v871 = *(*int32)(unsafe.Add(mBase, uint32(v870)))
	v872 = *(*int32)(unsafe.Add(mBase, uint32(v871)+12))
	v873 = *(*int32)(unsafe.Add(mBase, uint32(v871)+4))
	v879 = *(*int32)(unsafe.Add(mBase, uint32(v872+v873<<(uint(int32(2))%32)-int32(4))))
	v880 = *(*int32)(unsafe.Add(mBase, uint32(v879)+4))
	v884 = v880
	goto L279
L37:
	;
	v839 = *(*int32)(unsafe.Add(mBase, uint32(l2)+12))
	v840 = *(*int32)(unsafe.Add(mBase, uint32(v839)))
	v841 = *(*int32)(unsafe.Add(mBase, uint32(v840)+4))
	v842 = F_get_index_am_oid(m, v841)
	mBase = m.M
	v843 = m.ExcPending
	if v843 != 0 {
		goto L47
	} else {
		goto L268
	}
L38:
	;
	v678 = int32(2607)
	*(*int32)(unsafe.Add(mBase, uint32(l0))) = v678
	v681 = m.G0
	v683 = v681 - int32(16)
	m.G0 = v683
	F_DeconstructQualifiedName(m, l2, v683+int32(12), v683+int32(8))
	mBase = m.M
	v690 = m.ExcPending
	if v690 != 0 {
		goto L47
	} else {
		goto L237
	}
L39:
	;
	v673 = int32(3456)
	*(*int32)(unsafe.Add(mBase, uint32(l0))) = v673
	v676 = F_get_collation_oid(m, l2, l5)
	mBase = m.M
	v677 = m.ExcPending
	if v677 != 0 {
		goto L47
	} else {
		goto L236
	}
L40:
	;
	v668 = int32(2617)
	*(*int32)(unsafe.Add(mBase, uint32(l0))) = v668
	v671 = F_LookupOperWithArgs(m, l2, l5)
	mBase = m.M
	v672 = m.ExcPending
	if v672 != 0 {
		goto L47
	} else {
		goto L235
	}
L41:
	;
	v663 = int32(1255)
	*(*int32)(unsafe.Add(mBase, uint32(l0))) = v663
	v666 = F_LookupFuncWithArgs(m, l1, l2, l5)
	mBase = m.M
	v667 = m.ExcPending
	if v667 != 0 {
		goto L47
	} else {
		goto L234
	}
L42:
	;
	v580 = *(*int32)(unsafe.Add(mBase, uint32(l2)+4))
	switch l1 {
	case 0:
		goto L197
	default:
		goto L198
	case 9:
		goto L210
	case 14:
		goto L202
	case 15:
		goto L209
	case 16:
		goto L204
	case 17:
		goto L203
	case 21:
		goto L205
	case 27:
		goto L201
	case 30:
		goto L200
	case 33:
		goto L207
	case 36:
		goto L206
	case 38:
		goto L199
	case 42:
		goto L208
	}
L43:
	;
	v315 = *(*int32)(unsafe.Add(mBase, uint32(l2)+4))
	if v315 <= int32(1) {
		goto L14
	} else {
		goto L119
	}
L44:
	;
	if l2 == int32(0) {
		goto L16
	} else {
		goto L105
	}
L45:
	;
	if l2 == int32(0) {
		goto L18
	} else {
		goto L94
	}
L46:
	;
	v54 = F_makeRangeVarFromNameList(m, l2)
	mBase = m.M
	v55 = m.ExcPending
	if v55 != 0 {
		goto L47
	} else {
		goto L48
	}
L47:
	;
	return
L48:
	;
	v56 = F_relation_openrv_extended(m, v54, l4, l5)
	mBase = m.M
	v57 = m.ExcPending
	if v57 != 0 {
		goto L47
	} else {
		goto L49
	}
L49:
	;
	if v56 != 0 {
		goto L50
	} else {
		goto L51
	}
L50:
	;
	switch l1 - int32(18) {
	case 0:
		goto L55
	default:
		goto L54
	case 2:
		goto L60
	case 5:
		goto L56
	case 19:
		goto L59
	case 23:
		goto L58
	case 33:
		goto L57
	}
L51:
	;
	v228 = int32(0)
	goto L52
L52:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+8)) = int32(0)
	v1779 = v228
	v1781 = v56
	goto L5
L53:
	;
	v227 = *(*int32)(unsafe.Add(mBase, uint32(v56)+56))
	v228 = v227
	goto L52
L54:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v215 = m.ExcPending
	if v215 != 0 {
		goto L47
	} else {
		goto L91
	}
L55:
	;
	v187 = *(*int32)(unsafe.Add(mBase, uint32(v56)+48))
	v188 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v187)+119)))
	if v188 == int32(102) {
		goto L53
	} else {
		goto L86
	}
L56:
	;
	v162 = *(*int32)(unsafe.Add(mBase, uint32(v56)+48))
	v163 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v162)+119)))
	if v163 == int32(109) {
		goto L53
	} else {
		goto L81
	}
L57:
	;
	v137 = *(*int32)(unsafe.Add(mBase, uint32(v56)+48))
	v138 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v137)+119)))
	if v138 == int32(118) {
		goto L53
	} else {
		goto L76
	}
L58:
	;
	v112 = *(*int32)(unsafe.Add(mBase, uint32(v56)+48))
	v113 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v112)+119)))
	switch v113 - int32(112) {
	case 0, 2:
		goto L53
	default:
		goto L71
	}
L59:
	;
	v87 = *(*int32)(unsafe.Add(mBase, uint32(v56)+48))
	v88 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v87)+119)))
	if v88 == int32(83) {
		goto L53
	} else {
		goto L66
	}
L60:
	;
	v60 = *(*int32)(unsafe.Add(mBase, uint32(v56)+48))
	v61 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v60)+119)))
	if v61|int32(32) == int32(105) {
		goto L53
	} else {
		goto L61
	}
L61:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v69 = m.ExcPending
	if v69 != 0 {
		goto L47
	} else {
		goto L62
	}
L62:
	;
	F_errcode(m, int32(151027844))
	mBase = m.M
	v72 = m.ExcPending
	if v72 != 0 {
		goto L47
	} else {
		goto L63
	}
L63:
	;
	v73 = *(*int32)(unsafe.Add(mBase, uint32(v56)+48))
	*(*int32)(unsafe.Add(mBase, uint32(v24)+32)) = v73 + int32(4)
	F_errmsg(m, int32(28458), v24+int32(32))
	mBase = m.M
	v81 = m.ExcPending
	if v81 != 0 {
		goto L47
	} else {
		goto L64
	}
L64:
	;
	F_errfinish(m, int32(493656), int32(1362), int32(380897))
	mBase = m.M
	v86 = m.ExcPending
	if v86 != 0 {
		goto L47
	} else {
		goto L65
	}
L65:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L66:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v94 = m.ExcPending
	if v94 != 0 {
		goto L47
	} else {
		goto L67
	}
L67:
	;
	F_errcode(m, int32(151027844))
	mBase = m.M
	v97 = m.ExcPending
	if v97 != 0 {
		goto L47
	} else {
		goto L68
	}
L68:
	;
	v98 = *(*int32)(unsafe.Add(mBase, uint32(v56)+48))
	*(*int32)(unsafe.Add(mBase, uint32(v24)+48)) = v98 + int32(4)
	F_errmsg(m, int32(416122), v24+int32(48))
	mBase = m.M
	v106 = m.ExcPending
	if v106 != 0 {
		goto L47
	} else {
		goto L69
	}
L69:
	;
	F_errfinish(m, int32(493656), int32(1369), int32(380897))
	mBase = m.M
	v111 = m.ExcPending
	if v111 != 0 {
		goto L47
	} else {
		goto L70
	}
L70:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L71:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v119 = m.ExcPending
	if v119 != 0 {
		goto L47
	} else {
		goto L72
	}
L72:
	;
	F_errcode(m, int32(151027844))
	mBase = m.M
	v122 = m.ExcPending
	if v122 != 0 {
		goto L47
	} else {
		goto L73
	}
L73:
	;
	v123 = *(*int32)(unsafe.Add(mBase, uint32(v56)+48))
	*(*int32)(unsafe.Add(mBase, uint32(v24)+64)) = v123 + int32(4)
	F_errmsg(m, int32(395086), v24-int32(-64))
	mBase = m.M
	v131 = m.ExcPending
	if v131 != 0 {
		goto L47
	} else {
		goto L74
	}
L74:
	;
	F_errfinish(m, int32(493656), int32(1377), int32(380897))
	mBase = m.M
	v136 = m.ExcPending
	if v136 != 0 {
		goto L47
	} else {
		goto L75
	}
L75:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L76:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v144 = m.ExcPending
	if v144 != 0 {
		goto L47
	} else {
		goto L77
	}
L77:
	;
	F_errcode(m, int32(151027844))
	mBase = m.M
	v147 = m.ExcPending
	if v147 != 0 {
		goto L47
	} else {
		goto L78
	}
L78:
	;
	v148 = *(*int32)(unsafe.Add(mBase, uint32(v56)+48))
	*(*int32)(unsafe.Add(mBase, uint32(v24)+80)) = v148 + int32(4)
	F_errmsg(m, int32(32477), v24+int32(80))
	mBase = m.M
	v156 = m.ExcPending
	if v156 != 0 {
		goto L47
	} else {
		goto L79
	}
L79:
	;
	F_errfinish(m, int32(493656), int32(1384), int32(380897))
	mBase = m.M
	v161 = m.ExcPending
	if v161 != 0 {
		goto L47
	} else {
		goto L80
	}
L80:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L81:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v169 = m.ExcPending
	if v169 != 0 {
		goto L47
	} else {
		goto L82
	}
L82:
	;
	F_errcode(m, int32(151027844))
	mBase = m.M
	v172 = m.ExcPending
	if v172 != 0 {
		goto L47
	} else {
		goto L83
	}
L83:
	;
	v173 = *(*int32)(unsafe.Add(mBase, uint32(v56)+48))
	*(*int32)(unsafe.Add(mBase, uint32(v24)+96)) = v173 + int32(4)
	F_errmsg(m, int32(32422), v24+int32(96))
	mBase = m.M
	v181 = m.ExcPending
	if v181 != 0 {
		goto L47
	} else {
		goto L84
	}
L84:
	;
	F_errfinish(m, int32(493656), int32(1391), int32(380897))
	mBase = m.M
	v186 = m.ExcPending
	if v186 != 0 {
		goto L47
	} else {
		goto L85
	}
L85:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L86:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v194 = m.ExcPending
	if v194 != 0 {
		goto L47
	} else {
		goto L87
	}
L87:
	;
	F_errcode(m, int32(151027844))
	mBase = m.M
	v197 = m.ExcPending
	if v197 != 0 {
		goto L47
	} else {
		goto L88
	}
L88:
	;
	v198 = *(*int32)(unsafe.Add(mBase, uint32(v56)+48))
	*(*int32)(unsafe.Add(mBase, uint32(v24)+112)) = v198 + int32(4)
	F_errmsg(m, int32(393113), v24+int32(112))
	mBase = m.M
	v206 = m.ExcPending
	if v206 != 0 {
		goto L47
	} else {
		goto L89
	}
L89:
	;
	F_errfinish(m, int32(493656), int32(1398), int32(380897))
	mBase = m.M
	v211 = m.ExcPending
	if v211 != 0 {
		goto L47
	} else {
		goto L90
	}
L90:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L91:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v24)+16)) = l1
	F_errmsg_internal(m, int32(484755), v24+int32(16))
	mBase = m.M
	v221 = m.ExcPending
	if v221 != 0 {
		goto L47
	} else {
		goto L92
	}
L92:
	;
	F_errfinish(m, int32(493656), int32(1401), int32(380897))
	mBase = m.M
	v226 = m.ExcPending
	if v226 != 0 {
		goto L47
	} else {
		goto L93
	}
L93:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L94:
	;
	v233 = *(*int32)(unsafe.Add(mBase, uint32(l2)+4))
	if v233 <= int32(1) {
		goto L18
	} else {
		goto L95
	}
L95:
	;
	v236 = *(*int32)(unsafe.Add(mBase, uint32(l2)+12))
	v242 = *(*int32)(unsafe.Add(mBase, uint32(v236+v233<<(uint(int32(2))%32)-int32(4))))
	v243 = *(*int32)(unsafe.Add(mBase, uint32(v242)+4))
	v246 = F_list_copy_head(m, l2, v233-int32(1))
	mBase = m.M
	v247 = m.ExcPending
	if v247 != 0 {
		goto L47
	} else {
		goto L96
	}
L96:
	;
	v248 = F_makeRangeVarFromNameList(m, v246)
	mBase = m.M
	v249 = m.ExcPending
	if v249 != 0 {
		goto L47
	} else {
		goto L97
	}
L97:
	;
	v250 = F_relation_openrv(m, v248, l4)
	mBase = m.M
	v251 = m.ExcPending
	if v251 != 0 {
		goto L47
	} else {
		goto L98
	}
L98:
	;
	v252 = *(*int32)(unsafe.Add(mBase, uint32(v250)+56))
	v253 = F_get_attnum(m, v252, v243)
	mBase = m.M
	v254 = m.ExcPending
	if v254 != 0 {
		goto L47
	} else {
		goto L99
	}
L99:
	;
	if v253 == int32(0) {
		goto L100
	} else {
		goto L101
	}
L100:
	;
	if l5 == int32(0) {
		goto L17
	} else {
		goto L103
	}
L101:
	;
	v264 = v252
	v265 = v253
	v266 = v250
	goto L102
L102:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+8)) = v265
	v1779 = v264
	v1781 = v266
	goto L5
L103:
	;
	F_relation_close(m, v250, l4)
	mBase = m.M
	v260 = m.ExcPending
	if v260 != 0 {
		goto L47
	} else {
		goto L104
	}
L104:
	;
	v261 = int32(0)
	v264 = v261
	v265 = v261
	v266 = v261
	goto L102
L105:
	;
	v270 = *(*int32)(unsafe.Add(mBase, uint32(l2)+4))
	if v270 <= int32(1) {
		goto L16
	} else {
		goto L106
	}
L106:
	;
	v273 = *(*int32)(unsafe.Add(mBase, uint32(l2)+12))
	v279 = *(*int32)(unsafe.Add(mBase, uint32(v273+v270<<(uint(int32(2))%32)-int32(4))))
	v280 = *(*int32)(unsafe.Add(mBase, uint32(v279)+4))
	v283 = F_list_copy_head(m, l2, v270-int32(1))
	mBase = m.M
	v284 = m.ExcPending
	if v284 != 0 {
		goto L47
	} else {
		goto L107
	}
L107:
	;
	v285 = F_makeRangeVarFromNameList(m, v283)
	mBase = m.M
	v286 = m.ExcPending
	if v286 != 0 {
		goto L47
	} else {
		goto L108
	}
L108:
	;
	v287 = F_relation_openrv(m, v285, l4)
	mBase = m.M
	v288 = m.ExcPending
	if v288 != 0 {
		goto L47
	} else {
		goto L109
	}
L109:
	;
	v289 = *(*int32)(unsafe.Add(mBase, uint32(v287)+52))
	v290 = *(*int32)(unsafe.Add(mBase, uint32(v287)+56))
	v291 = F_get_attnum(m, v290, v280)
	mBase = m.M
	v292 = m.ExcPending
	if v292 != 0 {
		goto L47
	} else {
		goto L112
	}
L110:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+8)) = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v307
	v312 = int32(2604)
	*(*int32)(unsafe.Add(mBase, uint32(l0))) = v312
	v1794 = v307
	v1795 = v312
	v1796 = v308
	goto L4
L111:
	;
	if l5 == int32(0) {
		goto L15
	} else {
		goto L117
	}
L112:
	;
	if v291 == int32(0) {
		goto L111
	} else {
		goto L113
	}
L113:
	;
	v295 = *(*int32)(unsafe.Add(mBase, uint32(v289)+16))
	if v295 == int32(0) {
		goto L111
	} else {
		goto L114
	}
L114:
	;
	v298 = F_GetAttrDefaultOid(m, v290, v291)
	mBase = m.M
	v299 = m.ExcPending
	if v299 != 0 {
		goto L47
	} else {
		goto L115
	}
L115:
	;
	if v298 != 0 {
		v307 = v298
		v308 = v287
		goto L110
	} else {
		goto L116
	}
L116:
	;
	goto L111
L117:
	;
	F_relation_close(m, v287, l4)
	mBase = m.M
	v304 = m.ExcPending
	if v304 != 0 {
		goto L47
	} else {
		goto L118
	}
L118:
	;
	v305 = int32(0)
	v307 = v305
	v308 = v305
	goto L110
L119:
	;
	v318 = *(*int32)(unsafe.Add(mBase, uint32(l2)+12))
	v324 = *(*int32)(unsafe.Add(mBase, uint32(v318+v315<<(uint(int32(2))%32)-int32(4))))
	v325 = *(*int32)(unsafe.Add(mBase, uint32(v324)+4))
	v328 = F_list_copy_head(m, l2, v315-int32(1))
	mBase = m.M
	v329 = m.ExcPending
	if v329 != 0 {
		goto L47
	} else {
		goto L120
	}
L120:
	;
	v330 = F_makeRangeVarFromNameList(m, v328)
	mBase = m.M
	v331 = m.ExcPending
	if v331 != 0 {
		goto L47
	} else {
		goto L121
	}
L121:
	;
	v333 = F_table_openrv_extended(m, v330, int32(1), l5)
	mBase = m.M
	v334 = m.ExcPending
	if v334 != 0 {
		goto L47
	} else {
		goto L122
	}
L122:
	;
	if v333 != 0 {
		goto L123
	} else {
		goto L124
	}
L123:
	;
	v335 = *(*int32)(unsafe.Add(mBase, uint32(v333)+56))
	v337 = v335
	goto L125
L124:
	;
	v337 = int32(0)
	goto L125
L125:
	;
	switch l1 - int32(28) {
	case 0:
		goto L131
	default:
		goto L130
	case 7:
		goto L134
	case 12:
		goto L132
	case 16:
		goto L133
	}
L126:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+8)) = int32(0)
	v1768 = v570
	v1769 = v571
	v1770 = v572
	goto L6
L127:
	;
	v570 = v568
	v571 = v562
	v572 = int32(0)
	goto L126
L128:
	;
	if v556 != 0 {
		v570 = v556
		v571 = v550
		v572 = v333
		goto L126
	} else {
		goto L195
	}
L129:
	;
	v510 = m.G0
	v512 = v510 - int32(16)
	m.G0 = v512
	v515 = F_SearchSysCache2(m, int32(60), v337, v325)
	mBase = m.M
	v516 = m.ExcPending
	if v516 != 0 {
		goto L47
	} else {
		goto L184
	}
L130:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v496 = m.ExcPending
	if v496 != 0 {
		goto L47
	} else {
		goto L180
	}
L131:
	;
	if v333 == int32(0) {
		goto L160
	} else {
		goto L161
	}
L132:
	;
	if v333 == int32(0) {
		goto L156
	} else {
		goto L157
	}
L133:
	;
	if v333 == int32(0) {
		goto L136
	} else {
		goto L137
	}
L134:
	;
	if v333 != 0 {
		goto L129
	} else {
		goto L135
	}
L135:
	;
	v562 = int32(2618)
	v568 = int32(0)
	goto L127
L136:
	;
	v562 = int32(2620)
	v568 = int32(0)
	goto L127
L137:
	;
	goto L138
L138:
	;
	v346 = int32(2620)
	v348 = m.G0
	v350 = v348 - int32(112)
	m.G0 = v350
	v354 = F_table_open(m, v346, int32(1))
	mBase = m.M
	v355 = m.ExcPending
	if v355 != 0 {
		goto L47
	} else {
		goto L139
	}
L139:
	;
	F_ScanKeyInit(m, v350+int32(16), int32(2), int32(3), int32(184), v337)
	mBase = m.M
	v362 = m.ExcPending
	if v362 != 0 {
		goto L47
	} else {
		goto L140
	}
L140:
	;
	F_ScanKeyInit(m, v350-int32(-64), int32(4), int32(3), int32(62), v325)
	mBase = m.M
	v369 = m.ExcPending
	if v369 != 0 {
		goto L47
	} else {
		goto L141
	}
L141:
	;
	v376 = F_systable_beginscan(m, v354, int32(2701), int32(1), int32(0), int32(2), v350+int32(16))
	mBase = m.M
	v377 = m.ExcPending
	if v377 != 0 {
		goto L47
	} else {
		goto L143
	}
L142:
	;
	F_systable_endscan(m, v376)
	mBase = m.M
	v407 = m.ExcPending
	if v407 != 0 {
		goto L47
	} else {
		goto L154
	}
L143:
	;
	v378 = F_systable_getnext(m, v376)
	mBase = m.M
	v379 = m.ExcPending
	if v379 != 0 {
		goto L47
	} else {
		goto L144
	}
L144:
	;
	if v378 == int32(0) {
		goto L145
	} else {
		goto L146
	}
L145:
	;
	if l5 != 0 {
		v405 = int32(0)
		goto L142
	} else {
		goto L148
	}
L146:
	;
	goto L147
L147:
	;
	v401 = *(*int32)(unsafe.Add(mBase, uint32(v378)+16))
	v402 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v401)+22)))
	v404 = *(*int32)(unsafe.Add(mBase, uint32(v401+v402)))
	v405 = v404
	goto L142
L148:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v385 = m.ExcPending
	if v385 != 0 {
		goto L47
	} else {
		goto L149
	}
L149:
	;
	F_errcode(m, int32(67137668))
	mBase = m.M
	v388 = m.ExcPending
	if v388 != 0 {
		goto L47
	} else {
		goto L150
	}
L150:
	;
	v389 = F_get_rel_name(m, v337)
	mBase = m.M
	v390 = m.ExcPending
	if v390 != 0 {
		goto L47
	} else {
		goto L151
	}
L151:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v350)+4)) = v389
	*(*int32)(unsafe.Add(mBase, uint32(v350))) = v325
	F_errmsg(m, int32(72417), v350)
	mBase = m.M
	v395 = m.ExcPending
	if v395 != 0 {
		goto L47
	} else {
		goto L152
	}
L152:
	;
	F_errfinish(m, int32(495446), int32(1404), int32(434188))
	mBase = m.M
	v400 = m.ExcPending
	if v400 != 0 {
		goto L47
	} else {
		goto L153
	}
L153:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L154:
	;
	F_sequence_close(m, v354, int32(1))
	mBase = m.M
	v410 = m.ExcPending
	if v410 != 0 {
		goto L47
	} else {
		goto L155
	}
L155:
	;
	m.G0 = v350 + int32(112)
	v550 = v346
	v556 = v405
	goto L128
L156:
	;
	v562 = int32(2606)
	v568 = int32(0)
	goto L127
L157:
	;
	goto L158
L158:
	;
	v419 = F_get_relation_constraint_oid(m, v337, v325, l5)
	mBase = m.M
	v420 = m.ExcPending
	if v420 != 0 {
		goto L47
	} else {
		goto L159
	}
L159:
	;
	v550 = int32(2606)
	v556 = v419
	goto L128
L160:
	;
	v562 = int32(3256)
	v568 = int32(0)
	goto L127
L161:
	;
	goto L162
L162:
	;
	v425 = int32(3256)
	v427 = m.G0
	v429 = v427 - int32(112)
	m.G0 = v429
	v433 = F_table_open(m, v425, int32(1))
	mBase = m.M
	v434 = m.ExcPending
	if v434 != 0 {
		goto L47
	} else {
		goto L163
	}
L163:
	;
	v437 = int32(3)
	F_ScanKeyInit(m, v429+int32(16), v437, v437, int32(184), v337)
	mBase = m.M
	v441 = m.ExcPending
	if v441 != 0 {
		goto L47
	} else {
		goto L164
	}
L164:
	;
	F_ScanKeyInit(m, v429-int32(-64), int32(2), int32(3), int32(62), v325)
	mBase = m.M
	v448 = m.ExcPending
	if v448 != 0 {
		goto L47
	} else {
		goto L165
	}
L165:
	;
	v455 = F_systable_beginscan(m, v433, int32(3258), int32(1), int32(0), int32(2), v429+int32(16))
	mBase = m.M
	v456 = m.ExcPending
	if v456 != 0 {
		goto L47
	} else {
		goto L167
	}
L166:
	;
	F_systable_endscan(m, v455)
	mBase = m.M
	v486 = m.ExcPending
	if v486 != 0 {
		goto L47
	} else {
		goto L178
	}
L167:
	;
	v457 = F_systable_getnext(m, v455)
	mBase = m.M
	v458 = m.ExcPending
	if v458 != 0 {
		goto L47
	} else {
		goto L168
	}
L168:
	;
	if v457 == int32(0) {
		goto L169
	} else {
		goto L170
	}
L169:
	;
	if l5 != 0 {
		v484 = int32(0)
		goto L166
	} else {
		goto L172
	}
L170:
	;
	goto L171
L171:
	;
	v480 = *(*int32)(unsafe.Add(mBase, uint32(v457)+16))
	v481 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v480)+22)))
	v483 = *(*int32)(unsafe.Add(mBase, uint32(v480+v481)))
	v484 = v483
	goto L166
L172:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v464 = m.ExcPending
	if v464 != 0 {
		goto L47
	} else {
		goto L173
	}
L173:
	;
	F_errcode(m, int32(67137668))
	mBase = m.M
	v467 = m.ExcPending
	if v467 != 0 {
		goto L47
	} else {
		goto L174
	}
L174:
	;
	v468 = F_get_rel_name(m, v337)
	mBase = m.M
	v469 = m.ExcPending
	if v469 != 0 {
		goto L47
	} else {
		goto L175
	}
L175:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v429)+4)) = v468
	*(*int32)(unsafe.Add(mBase, uint32(v429))) = v325
	F_errmsg(m, int32(72288), v429)
	mBase = m.M
	v474 = m.ExcPending
	if v474 != 0 {
		goto L47
	} else {
		goto L176
	}
L176:
	;
	F_errfinish(m, int32(492363), int32(1238), int32(433604))
	mBase = m.M
	v479 = m.ExcPending
	if v479 != 0 {
		goto L47
	} else {
		goto L177
	}
L177:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L178:
	;
	F_sequence_close(m, v433, int32(1))
	mBase = m.M
	v489 = m.ExcPending
	if v489 != 0 {
		goto L47
	} else {
		goto L179
	}
L179:
	;
	m.G0 = v429 + int32(112)
	v550 = v425
	v556 = v484
	goto L128
L180:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v24)+160)) = l1
	F_errmsg_internal(m, int32(484755), v24+int32(160))
	mBase = m.M
	v502 = m.ExcPending
	if v502 != 0 {
		goto L47
	} else {
		goto L181
	}
L181:
	;
	F_errfinish(m, int32(493656), int32(1477), int32(110302))
	mBase = m.M
	v507 = m.ExcPending
	if v507 != 0 {
		goto L47
	} else {
		goto L182
	}
L182:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L183:
	;
	m.G0 = v512 + int32(16)
	v550 = int32(2618)
	v556 = v545
	goto L128
L184:
	;
	if v515 == int32(0) {
		goto L185
	} else {
		goto L186
	}
L185:
	;
	if l5 != 0 {
		v545 = int32(0)
		goto L183
	} else {
		goto L188
	}
L186:
	;
	goto L187
L187:
	;
	v538 = *(*int32)(unsafe.Add(mBase, uint32(v515)+16))
	v539 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v538)+22)))
	v541 = *(*int32)(unsafe.Add(mBase, uint32(v538+v539)))
	F_ReleaseCatCache(m, v515)
	mBase = m.M
	v543 = m.ExcPending
	if v543 != 0 {
		goto L47
	} else {
		goto L194
	}
L188:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v522 = m.ExcPending
	if v522 != 0 {
		goto L47
	} else {
		goto L189
	}
L189:
	;
	F_errcode(m, int32(67137668))
	mBase = m.M
	v525 = m.ExcPending
	if v525 != 0 {
		goto L47
	} else {
		goto L190
	}
L190:
	;
	v526 = F_get_rel_name(m, v337)
	mBase = m.M
	v527 = m.ExcPending
	if v527 != 0 {
		goto L47
	} else {
		goto L191
	}
L191:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v512)+4)) = v526
	*(*int32)(unsafe.Add(mBase, uint32(v512))) = v325
	F_errmsg(m, int32(71357), v512)
	mBase = m.M
	v532 = m.ExcPending
	if v532 != 0 {
		goto L47
	} else {
		goto L192
	}
L192:
	;
	F_errfinish(m, int32(492844), int32(109), int32(434559))
	mBase = m.M
	v537 = m.ExcPending
	if v537 != 0 {
		goto L47
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
	v545 = v541
	goto L183
L195:
	;
	F_sequence_close(m, v333, int32(1))
	mBase = m.M
	v559 = m.ExcPending
	if v559 != 0 {
		goto L47
	} else {
		goto L196
	}
L196:
	;
	v562 = v550
	v568 = int32(0)
	goto L127
L197:
	;
	v661 = F_get_am_type_oid(m, v580, int32(0), l5)
	mBase = m.M
	v662 = m.ExcPending
	if v662 != 0 {
		goto L47
	} else {
		goto L233
	}
L198:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v647 = m.ExcPending
	if v647 != 0 {
		goto L47
	} else {
		goto L230
	}
L199:
	;
	v642 = F_get_subscription_oid(m, v580, l5)
	mBase = m.M
	v643 = m.ExcPending
	if v643 != 0 {
		goto L47
	} else {
		goto L229
	}
L200:
	;
	v639 = F_get_publication_oid(m, v580, l5)
	mBase = m.M
	v640 = m.ExcPending
	if v640 != 0 {
		goto L47
	} else {
		goto L228
	}
L201:
	;
	v636 = F_ParameterAclLookup(m, v580, l5)
	mBase = m.M
	v637 = m.ExcPending
	if v637 != 0 {
		goto L47
	} else {
		goto L227
	}
L202:
	;
	v606 = m.G0
	v608 = v606 - int32(16)
	m.G0 = v608
	v611 = int32(0)
	v614 = F_GetSysCacheOid(m, int32(25), v580, v611, v611, v611)
	mBase = m.M
	v615 = m.ExcPending
	if v615 != 0 {
		goto L47
	} else {
		goto L219
	}
L203:
	;
	v603 = F_get_foreign_server_oid(m, v580, l5)
	mBase = m.M
	v604 = m.ExcPending
	if v604 != 0 {
		goto L47
	} else {
		goto L218
	}
L204:
	;
	v600 = F_get_foreign_data_wrapper_oid(m, v580, l5)
	mBase = m.M
	v601 = m.ExcPending
	if v601 != 0 {
		goto L47
	} else {
		goto L217
	}
L205:
	;
	v597 = F_get_language_oid(m, v580, l5)
	mBase = m.M
	v598 = m.ExcPending
	if v598 != 0 {
		goto L47
	} else {
		goto L216
	}
L206:
	;
	v594 = F_get_namespace_oid(m, v580, l5)
	mBase = m.M
	v595 = m.ExcPending
	if v595 != 0 {
		goto L47
	} else {
		goto L215
	}
L207:
	;
	v591 = F_get_role_oid(m, v580, l5)
	mBase = m.M
	v592 = m.ExcPending
	if v592 != 0 {
		goto L47
	} else {
		goto L214
	}
L208:
	;
	v588 = F_get_tablespace_oid(m, v580, l5)
	mBase = m.M
	v589 = m.ExcPending
	if v589 != 0 {
		goto L47
	} else {
		goto L213
	}
L209:
	;
	v585 = F_get_extension_oid(m, v580, l5)
	mBase = m.M
	v586 = m.ExcPending
	if v586 != 0 {
		goto L47
	} else {
		goto L212
	}
L210:
	;
	v582 = F_get_database_oid(m, v580, l5)
	mBase = m.M
	v583 = m.ExcPending
	if v583 != 0 {
		goto L47
	} else {
		goto L211
	}
L211:
	;
	v1758 = v582
	v1759 = int32(1262)
	goto L7
L212:
	;
	v1758 = v585
	v1759 = int32(3079)
	goto L7
L213:
	;
	v1758 = v588
	v1759 = int32(1213)
	goto L7
L214:
	;
	v1758 = v591
	v1759 = int32(1260)
	goto L7
L215:
	;
	v1758 = v594
	v1759 = int32(2615)
	goto L7
L216:
	;
	v1758 = v597
	v1759 = int32(2612)
	goto L7
L217:
	;
	v1758 = v600
	v1759 = int32(2328)
	goto L7
L218:
	;
	v1758 = v603
	v1759 = int32(1417)
	goto L7
L219:
	;
	if l5 != 0 {
		goto L220
	} else {
		goto L221
	}
L220:
	;
	m.G0 = v608 + int32(16)
	v1758 = v614
	v1759 = int32(3466)
	goto L7
L221:
	;
	if v614 != 0 {
		goto L220
	} else {
		goto L222
	}
L222:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v619 = m.ExcPending
	if v619 != 0 {
		goto L47
	} else {
		goto L223
	}
L223:
	;
	F_errcode(m, int32(67137668))
	mBase = m.M
	v622 = m.ExcPending
	if v622 != 0 {
		goto L47
	} else {
		goto L224
	}
L224:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v608))) = v580
	F_errmsg(m, int32(71120), v608)
	mBase = m.M
	v626 = m.ExcPending
	if v626 != 0 {
		goto L47
	} else {
		goto L225
	}
L225:
	;
	F_errfinish(m, int32(495440), int32(588), int32(434166))
	mBase = m.M
	v631 = m.ExcPending
	if v631 != 0 {
		goto L47
	} else {
		goto L226
	}
L226:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L227:
	;
	v1758 = v636
	v1759 = int32(6243)
	goto L7
L228:
	;
	v1758 = v639
	v1759 = int32(6104)
	goto L7
L229:
	;
	v1758 = v642
	v1759 = int32(6100)
	goto L7
L230:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v24)+176)) = l1
	F_errmsg_internal(m, int32(484755), v24+int32(176))
	mBase = m.M
	v653 = m.ExcPending
	if v653 != 0 {
		goto L47
	} else {
		goto L231
	}
L231:
	;
	F_errfinish(m, int32(493656), int32(1324), int32(456570))
	mBase = m.M
	v658 = m.ExcPending
	if v658 != 0 {
		goto L47
	} else {
		goto L232
	}
L232:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L233:
	;
	v1758 = v661
	v1759 = int32(2601)
	goto L7
L234:
	;
	v1468 = v663
	v1482 = v666
	goto L20
L235:
	;
	v1468 = v668
	v1482 = v671
	goto L20
L236:
	;
	v1468 = v673
	v1482 = v676
	goto L20
L237:
	;
	v691 = *(*int32)(unsafe.Add(mBase, uint32(v683)+12))
	if v691 != 0 {
		goto L241
	} else {
		goto L242
	}
L238:
	;
	m.G0 = v683 + int32(16)
	v1468 = v678
	v1482 = v825
	goto L20
L239:
	;
	if l5 != 0 {
		v825 = v796
		goto L238
	} else {
		goto L261
	}
L240:
	;
	v796 = int32(0)
	goto L239
L241:
	;
	v693 = F_LookupExplicitNamespace(m, v691, l5)
	mBase = m.M
	v694 = m.ExcPending
	if v694 != 0 {
		goto L47
	} else {
		goto L244
	}
L242:
	;
	goto L243
L243:
	;
	F_recomputeNamespacePath(m)
	mBase = m.M
	v703 = m.ExcPending
	if v703 != 0 {
		goto L47
	} else {
		goto L250
	}
L244:
	;
	if v693 != 0 {
		goto L245
	} else {
		goto L246
	}
L245:
	;
	v695 = int32(0)
	goto L247
L246:
	;
	v695 = l5
	goto L247
L247:
	;
	if v695 != 0 {
		goto L240
	} else {
		goto L248
	}
L248:
	;
	v697 = *(*int32)(unsafe.Add(mBase, uint32(v683)+8))
	v698 = int32(0)
	v700 = F_GetSysCacheOid(m, int32(18), v697, v693, v698, v698)
	mBase = m.M
	v701 = m.ExcPending
	if v701 != 0 {
		goto L47
	} else {
		goto L249
	}
L249:
	;
	v796 = v700
	goto L239
L250:
	;
	v704 = int32(0)
	v706 = *(*int32)(unsafe.Add(mBase, _consts[317]))
	if v706 == v704 {
		v796 = v704
		goto L239
	} else {
		goto L251
	}
L251:
	;
	v709 = int32(0)
	v710 = *(*int32)(unsafe.Add(mBase, uint32(v706)+4))
	if v710 <= v709 {
		goto L240
	} else {
		goto L252
	}
L252:
	;
	v713 = *(*int32)(unsafe.Add(mBase, uint32(v683)+8))
	v720 = v709
	goto L253
L253:
	;
	v735 = *(*int32)(unsafe.Add(mBase, uint32(v706)+12))
	v739 = *(*int32)(unsafe.Add(mBase, uint32(v735+v720<<(uint(int32(2))%32))))
	v741 = *(*int32)(unsafe.Add(mBase, _consts[316]))
	if v739 != v741 {
		goto L255
	} else {
		goto L256
	}
L254:
	;
	goto L240
L255:
	;
	v744 = int32(0)
	v746 = F_GetSysCacheOid(m, int32(18), v713, v739, v744, v744)
	mBase = m.M
	v747 = m.ExcPending
	if v747 != 0 {
		goto L47
	} else {
		goto L258
	}
L256:
	;
	goto L257
L257:
	;
	v750 = v720 + int32(1)
	v751 = *(*int32)(unsafe.Add(mBase, uint32(v706)+4))
	if v750 < v751 {
		v720 = v750
		goto L253
	} else {
		goto L260
	}
L258:
	;
	if v746 != 0 {
		v825 = v746
		goto L238
	} else {
		goto L259
	}
L259:
	;
	goto L257
L260:
	;
	goto L254
L261:
	;
	if v796 != 0 {
		v825 = v796
		goto L238
	} else {
		goto L262
	}
L262:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v800 = m.ExcPending
	if v800 != 0 {
		goto L47
	} else {
		goto L263
	}
L263:
	;
	F_errcode(m, int32(67137668))
	mBase = m.M
	v803 = m.ExcPending
	if v803 != 0 {
		goto L47
	} else {
		goto L264
	}
L264:
	;
	v804 = F_NameListToString(m, l2)
	mBase = m.M
	v805 = m.ExcPending
	if v805 != 0 {
		goto L47
	} else {
		goto L265
	}
L265:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v683))) = v804
	F_errmsg(m, int32(71766), v683)
	mBase = m.M
	v809 = m.ExcPending
	if v809 != 0 {
		goto L47
	} else {
		goto L266
	}
L266:
	;
	F_errfinish(m, int32(499480), int32(4075), int32(434292))
	mBase = m.M
	v814 = m.ExcPending
	if v814 != 0 {
		goto L47
	} else {
		goto L267
	}
L267:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L268:
	;
	v845 = F_list_copy_tail(m, l2, int32(1))
	mBase = m.M
	v846 = m.ExcPending
	if v846 != 0 {
		goto L47
	} else {
		goto L269
	}
L269:
	;
	switch l1 - int32(24) {
	case 0:
		goto L270
	default:
		goto L271
	case 2:
		goto L272
	}
L270:
	;
	v868 = F_get_opclass_oid(m, v842, v845, l5)
	mBase = m.M
	v869 = m.ExcPending
	if v869 != 0 {
		goto L47
	} else {
		goto L277
	}
L271:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v855 = m.ExcPending
	if v855 != 0 {
		goto L47
	} else {
		goto L274
	}
L272:
	;
	v850 = F_get_opfamily_oid(m, v842, v845, l5)
	mBase = m.M
	v851 = m.ExcPending
	if v851 != 0 {
		goto L47
	} else {
		goto L273
	}
L273:
	;
	v1758 = v850
	v1759 = int32(2753)
	goto L7
L274:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v24)+192)) = l1
	F_errmsg_internal(m, int32(484755), v24+int32(192))
	mBase = m.M
	v861 = m.ExcPending
	if v861 != 0 {
		goto L47
	} else {
		goto L275
	}
L275:
	;
	F_errfinish(m, int32(493656), int32(1669), int32(339765))
	mBase = m.M
	v866 = m.ExcPending
	if v866 != 0 {
		goto L47
	} else {
		goto L276
	}
L276:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L277:
	;
	v1758 = v868
	v1759 = int32(2616)
	goto L7
L278:
	;
	v929 = int32(0)
	v932 = F_list_copy_head(m, v871, v873-int32(1))
	mBase = m.M
	v933 = m.ExcPending
	if v933 != 0 {
		goto L47
	} else {
		goto L294
	}
L279:
	;
	v889 = v884 + int32(1)
	v890 = int32(*(*int8)(unsafe.Add(mBase, uint32(v884))))
	v891 = F___isspace(m, v890)
	mBase = m.M
	if v891 != 0 {
		v884 = v889
		goto L279
	} else {
		goto L281
	}
L280:
	;
	v892 = int32(1)
	switch v890&int32(255) - int32(43) {
	case 0:
		v898 = v892
		goto L283
	default:
		v900 = v890
		v901 = v884
		v902 = v892
		goto L282
	case 2:
		goto L284
	}
L281:
	;
	goto L280
L282:
	;
	v903 = int32(0)
	v905 = v900 - int32(48)
	if base.Ui32(v905) <= base.Ui32(int32(9)) {
		goto L285
	} else {
		goto L286
	}
L283:
	;
	v899 = int32(*(*int8)(unsafe.Add(mBase, uint32(v889))))
	v900 = v899
	v901 = v889
	v902 = v898
	goto L282
L284:
	;
	v898 = int32(0)
	goto L283
L285:
	;
	v908 = v903
	v909 = v905
	v910 = v901
	goto L288
L286:
	;
	v922 = v903
	goto L287
L287:
	;
	if v902 != 0 {
		goto L291
	} else {
		goto L292
	}
L288:
	;
	v912 = int32(10)
	v914 = v908*v912 - v909
	v915 = int32(*(*int8)(unsafe.Add(mBase, uint32(v910)+1)))
	v919 = v915 - int32(48)
	if base.Ui32(v919) < base.Ui32(v912) {
		v908 = v914
		v909 = v919
		v910 = v910 + int32(1)
		goto L288
	} else {
		goto L290
	}
L289:
	;
	v922 = v914
	goto L287
L290:
	;
	goto L289
L291:
	;
	v928 = int32(0) - v922
	goto L293
L292:
	;
	v928 = v922
	goto L293
L293:
	;
	goto L278
L294:
	;
	v934 = *(*int32)(unsafe.Add(mBase, uint32(v932)+12))
	v935 = *(*int32)(unsafe.Add(mBase, uint32(v934)))
	v936 = *(*int32)(unsafe.Add(mBase, uint32(v935)+4))
	v937 = F_get_index_am_oid(m, v936)
	mBase = m.M
	v938 = m.ExcPending
	if v938 != 0 {
		goto L47
	} else {
		goto L295
	}
L295:
	;
	v940 = F_list_copy_tail(m, v932, int32(1))
	mBase = m.M
	v941 = m.ExcPending
	if v941 != 0 {
		goto L47
	} else {
		goto L296
	}
L296:
	;
	v943 = F_get_opfamily_oid(m, v937, v940, int32(0))
	mBase = m.M
	v944 = m.ExcPending
	if v944 != 0 {
		goto L47
	} else {
		goto L297
	}
L297:
	;
	v945 = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v24)+444)) = v945
	*(*int32)(unsafe.Add(mBase, uint32(v24)+440)) = v943
	*(*int32)(unsafe.Add(mBase, uint32(v24)+436)) = int32(2753)
	v950 = *(*int32)(unsafe.Add(mBase, uint32(l2)+12))
	v951 = *(*int32)(unsafe.Add(mBase, uint32(v950)+4))
	if v951 == v945 {
		goto L9
	} else {
		goto L298
	}
L298:
	;
	v954 = *(*int32)(unsafe.Add(mBase, uint32(v951)+4))
	if v954 <= int32(0) {
		goto L9
	} else {
		goto L299
	}
L299:
	;
	v960 = *(*int32)(unsafe.Add(mBase, uint32(v951)+12))
	v961 = *(*int32)(unsafe.Add(mBase, uint32(v960)))
	F_get_object_address_type(m, v24+int32(424), int32(49), v961, l5)
	mBase = m.M
	v963 = m.ExcPending
	if v963 != 0 {
		goto L47
	} else {
		goto L300
	}
L300:
	;
	v964 = *(*int32)(unsafe.Add(mBase, uint32(v24)+428))
	v966 = *(*int32)(unsafe.Add(mBase, uint32(v951)+4))
	if v966 < int32(2) {
		v1669 = v964
		v1670 = v929
		v1671 = v961
		v1672 = int32(0)
		goto L8
	} else {
		goto L301
	}
L301:
	;
	v972 = *(*int32)(unsafe.Add(mBase, uint32(v951)+12))
	v973 = *(*int32)(unsafe.Add(mBase, uint32(v972)+4))
	F_get_object_address_type(m, v24+int32(424), int32(49), v973, l5)
	mBase = m.M
	v975 = m.ExcPending
	if v975 != 0 {
		goto L47
	} else {
		goto L302
	}
L302:
	;
	v976 = *(*int32)(unsafe.Add(mBase, uint32(v24)+428))
	v1669 = v964
	v1670 = v973
	v1671 = v961
	v1672 = v976
	goto L8
L303:
	;
	m.G0 = v982 + int32(16)
	*(*int32)(unsafe.Add(mBase, uint32(l0)+8)) = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v1008
	v1015 = F_LargeObjectExists(m, v1008)
	mBase = m.M
	v1016 = m.ExcPending
	if v1016 != 0 {
		goto L47
	} else {
		goto L311
	}
L304:
	;
	v1007 = *(*int32)(unsafe.Add(mBase, uint32(l2)+4))
	v1008 = v1007
	goto L303
L305:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v996 = m.ExcPending
	if v996 != 0 {
		goto L47
	} else {
		goto L308
	}
L306:
	;
	v987 = *(*int32)(unsafe.Add(mBase, uint32(l2)+4))
	v988 = int32(0)
	v991 = F_uint32in_subr(m, v987, v988, int32(435118), v988)
	mBase = m.M
	v992 = m.ExcPending
	if v992 != 0 {
		goto L47
	} else {
		goto L307
	}
L307:
	;
	v1008 = v991
	goto L303
L308:
	;
	v997 = *(*int32)(unsafe.Add(mBase, uint32(l2)))
	*(*int32)(unsafe.Add(mBase, uint32(v982))) = v997
	F_errmsg_internal(m, int32(485936), v982)
	mBase = m.M
	v1001 = m.ExcPending
	if v1001 != 0 {
		goto L47
	} else {
		goto L309
	}
L309:
	;
	F_errfinish(m, int32(499674), int32(280), int32(360834))
	mBase = m.M
	v1006 = m.ExcPending
	if v1006 != 0 {
		goto L47
	} else {
		goto L310
	}
L310:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L311:
	;
	if l5 != 0 {
		v1493 = v977
		goto L19
	} else {
		goto L312
	}
L312:
	;
	if v1015 != 0 {
		v1493 = v977
		goto L19
	} else {
		goto L313
	}
L313:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v1020 = m.ExcPending
	if v1020 != 0 {
		goto L47
	} else {
		goto L314
	}
L314:
	;
	F_errcode(m, int32(67137668))
	mBase = m.M
	v1023 = m.ExcPending
	if v1023 != 0 {
		goto L47
	} else {
		goto L315
	}
L315:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v24)+240)) = v1008
	F_errmsg(m, int32(68740), v24+int32(240))
	mBase = m.M
	v1029 = m.ExcPending
	if v1029 != 0 {
		goto L47
	} else {
		goto L316
	}
L316:
	;
	F_errfinish(m, int32(493656), int32(1056), int32(128300))
	mBase = m.M
	v1034 = m.ExcPending
	if v1034 != 0 {
		goto L47
	} else {
		goto L317
	}
L317:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L318:
	;
	v1041 = F_LookupTypeNameOid(m, v1036, l5)
	mBase = m.M
	v1042 = m.ExcPending
	if v1042 != 0 {
		goto L47
	} else {
		goto L319
	}
L319:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0))) = int32(2605)
	v1045 = m.G0
	v1047 = v1045 - int32(16)
	m.G0 = v1047
	v1050 = int32(0)
	v1052 = F_GetSysCacheOid(m, int32(12), v1039, v1041, v1050, v1050)
	mBase = m.M
	v1053 = m.ExcPending
	if v1053 != 0 {
		goto L47
	} else {
		goto L320
	}
L320:
	;
	if l5 != 0 {
		goto L321
	} else {
		goto L322
	}
L321:
	;
	m.G0 = v1047 + int32(16)
	v1468 = int32(2605)
	v1482 = v1052
	goto L20
L322:
	;
	if v1052 != 0 {
		goto L321
	} else {
		goto L323
	}
L323:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v1057 = m.ExcPending
	if v1057 != 0 {
		goto L47
	} else {
		goto L324
	}
L324:
	;
	F_errcode(m, int32(67137668))
	mBase = m.M
	v1060 = m.ExcPending
	if v1060 != 0 {
		goto L47
	} else {
		goto L325
	}
L325:
	;
	v1061 = F_format_type_be(m, v1039)
	mBase = m.M
	v1062 = m.ExcPending
	if v1062 != 0 {
		goto L47
	} else {
		goto L326
	}
L326:
	;
	v1063 = F_format_type_be(m, v1041)
	mBase = m.M
	v1064 = m.ExcPending
	if v1064 != 0 {
		goto L47
	} else {
		goto L327
	}
L327:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v1047)+4)) = v1063
	*(*int32)(unsafe.Add(mBase, uint32(v1047))) = v1061
	F_errmsg(m, int32(70098), v1047)
	mBase = m.M
	v1069 = m.ExcPending
	if v1069 != 0 {
		goto L47
	} else {
		goto L328
	}
L328:
	;
	F_errfinish(m, int32(499145), int32(1111), int32(433678))
	mBase = m.M
	v1074 = m.ExcPending
	if v1074 != 0 {
		goto L47
	} else {
		goto L329
	}
L329:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L330:
	;
	v1085 = F_get_language_oid(m, v1080, l5)
	mBase = m.M
	v1086 = m.ExcPending
	if v1086 != 0 {
		goto L47
	} else {
		goto L331
	}
L331:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0))) = int32(3576)
	v1089 = F_get_transform_oid(m, v1083, v1085, l5)
	mBase = m.M
	v1090 = m.ExcPending
	if v1090 != 0 {
		goto L47
	} else {
		goto L332
	}
L332:
	;
	v1468 = int32(3576)
	v1482 = v1089
	goto L20
L333:
	;
	v1468 = v1091
	v1482 = v1094
	goto L20
L334:
	;
	v1468 = v1096
	v1482 = v1099
	goto L20
L335:
	;
	v1468 = v1101
	v1482 = v1104
	goto L20
L336:
	;
	v1468 = v1106
	v1482 = v1109
	goto L20
L337:
	;
	v1238 = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(l0)+8)) = v1238
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v1234
	v1242 = int32(1418)
	*(*int32)(unsafe.Add(mBase, uint32(l0))) = v1242
	v1794 = v1234
	v1795 = v1242
	v1796 = v1238
	goto L4
L338:
	;
	v1179 = F_GetForeignServerByName(m, v1113, int32(1))
	mBase = m.M
	v1180 = m.ExcPending
	if v1180 != 0 {
		goto L47
	} else {
		goto L362
	}
L339:
	;
	if v1140-v1139 == int32(0) {
		goto L347
	} else {
		goto L348
	}
L340:
	;
	goto L339
L341:
	;
	if v1119 != v1120 {
		v1139 = v1119
		v1140 = v1120
		goto L340
	} else {
		goto L342
	}
L342:
	;
	v1124 = v1115
	v1125 = v1116
	goto L343
L343:
	;
	v1128 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1125)+1)))
	v1129 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1124)+1)))
	if v1129 == int32(0) {
		v1139 = v1128
		v1140 = v1129
		goto L340
	} else {
		goto L345
	}
L344:
	;
	v1139 = v1128
	v1140 = v1129
	goto L340
L345:
	;
	v1132 = int32(1)
	if v1128 == v1129 {
		v1124 = v1124 + v1132
		v1125 = v1125 + v1132
		goto L343
	} else {
		goto L346
	}
L346:
	;
	goto L344
L347:
	;
	v1177 = int32(0)
	goto L338
L348:
	;
	goto L349
L349:
	;
	v1146 = F_SearchSysCache1(m, int32(10), v1115)
	mBase = m.M
	v1147 = m.ExcPending
	if v1147 != 0 {
		goto L47
	} else {
		goto L350
	}
L350:
	;
	if v1146 == int32(0) {
		goto L351
	} else {
		goto L352
	}
L351:
	;
	if l5 != 0 {
		goto L354
	} else {
		goto L355
	}
L352:
	;
	goto L353
L353:
	;
	v1170 = *(*int32)(unsafe.Add(mBase, uint32(v1146)+16))
	v1171 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1170)+22)))
	v1173 = *(*int32)(unsafe.Add(mBase, uint32(v1170+v1171)))
	F_ReleaseCatCache(m, v1146)
	mBase = m.M
	v1175 = m.ExcPending
	if v1175 != 0 {
		goto L47
	} else {
		goto L361
	}
L354:
	;
	v1234 = int32(0)
	goto L337
L355:
	;
	goto L356
L356:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v1154 = m.ExcPending
	if v1154 != 0 {
		goto L47
	} else {
		goto L357
	}
L357:
	;
	F_errcode(m, int32(67137668))
	mBase = m.M
	v1157 = m.ExcPending
	if v1157 != 0 {
		goto L47
	} else {
		goto L358
	}
L358:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v24)+292)) = v1113
	*(*int32)(unsafe.Add(mBase, uint32(v24)+288)) = v1115
	F_errmsg(m, int32(70861), v24+int32(288))
	mBase = m.M
	v1164 = m.ExcPending
	if v1164 != 0 {
		goto L47
	} else {
		goto L359
	}
L359:
	;
	F_errfinish(m, int32(493656), int32(1825), int32(333966))
	mBase = m.M
	v1169 = m.ExcPending
	if v1169 != 0 {
		goto L47
	} else {
		goto L360
	}
L360:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L361:
	;
	v1177 = v1173
	goto L338
L362:
	;
	if v1179 == int32(0) {
		goto L363
	} else {
		goto L364
	}
L363:
	;
	if l5 != 0 {
		goto L366
	} else {
		goto L367
	}
L364:
	;
	goto L365
L365:
	;
	v1203 = *(*int32)(unsafe.Add(mBase, uint32(v1179)))
	v1204 = F_SearchSysCache2(m, int32(84), v1177, v1203)
	mBase = m.M
	v1205 = m.ExcPending
	if v1205 != 0 {
		goto L47
	} else {
		goto L373
	}
L366:
	;
	v1234 = int32(0)
	goto L337
L367:
	;
	goto L368
L368:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v1187 = m.ExcPending
	if v1187 != 0 {
		goto L47
	} else {
		goto L369
	}
L369:
	;
	F_errcode(m, int32(67137668))
	mBase = m.M
	v1190 = m.ExcPending
	if v1190 != 0 {
		goto L47
	} else {
		goto L370
	}
L370:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v24)+256)) = v1113
	F_errmsg(m, int32(70891), v24+int32(256))
	mBase = m.M
	v1196 = m.ExcPending
	if v1196 != 0 {
		goto L47
	} else {
		goto L371
	}
L371:
	;
	F_errfinish(m, int32(493656), int32(1839), int32(333966))
	mBase = m.M
	v1201 = m.ExcPending
	if v1201 != 0 {
		goto L47
	} else {
		goto L372
	}
L372:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L373:
	;
	if v1204 == int32(0) {
		goto L374
	} else {
		goto L375
	}
L374:
	;
	if l5 != 0 {
		goto L377
	} else {
		goto L378
	}
L375:
	;
	goto L376
L376:
	;
	v1228 = *(*int32)(unsafe.Add(mBase, uint32(v1204)+16))
	v1229 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1228)+22)))
	v1231 = *(*int32)(unsafe.Add(mBase, uint32(v1228+v1229)))
	F_ReleaseCatCache(m, v1204)
	mBase = m.M
	v1233 = m.ExcPending
	if v1233 != 0 {
		goto L47
	} else {
		goto L384
	}
L377:
	;
	v1234 = int32(0)
	goto L337
L378:
	;
	goto L379
L379:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v1212 = m.ExcPending
	if v1212 != 0 {
		goto L47
	} else {
		goto L380
	}
L380:
	;
	F_errcode(m, int32(67137668))
	mBase = m.M
	v1215 = m.ExcPending
	if v1215 != 0 {
		goto L47
	} else {
		goto L381
	}
L381:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v24)+276)) = v1113
	*(*int32)(unsafe.Add(mBase, uint32(v24)+272)) = v1115
	F_errmsg(m, int32(70861), v24+int32(272))
	mBase = m.M
	v1222 = m.ExcPending
	if v1222 != 0 {
		goto L47
	} else {
		goto L382
	}
L382:
	;
	F_errfinish(m, int32(493656), int32(1851), int32(333966))
	mBase = m.M
	v1227 = m.ExcPending
	if v1227 != 0 {
		goto L47
	} else {
		goto L383
	}
L383:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L384:
	;
	v1234 = v1231
	goto L337
L385:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+8)) = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v1268
	v1273 = int32(6237)
	*(*int32)(unsafe.Add(mBase, uint32(l0))) = v1273
	v1794 = v1268
	v1795 = v1273
	v1796 = v1248
	goto L4
L386:
	;
	if v1252 == int32(0) {
		v1268 = v1248
		goto L385
	} else {
		goto L387
	}
L387:
	;
	v1256 = F_GetPublicationByName(m, v1247, l5)
	mBase = m.M
	v1257 = m.ExcPending
	if v1257 != 0 {
		goto L47
	} else {
		goto L388
	}
L388:
	;
	if v1256 == int32(0) {
		v1268 = v1248
		goto L385
	} else {
		goto L389
	}
L389:
	;
	v1261 = *(*int32)(unsafe.Add(mBase, uint32(v1256)))
	v1262 = int32(0)
	v1264 = F_GetSysCacheOid(m, int32(50), v1252, v1261, v1262, v1262)
	mBase = m.M
	v1265 = m.ExcPending
	if v1265 != 0 {
		goto L47
	} else {
		goto L390
	}
L390:
	;
	if l5 != 0 {
		v1268 = v1264
		goto L385
	} else {
		goto L391
	}
L391:
	;
	if v1264 == int32(0) {
		goto L13
	} else {
		goto L392
	}
L392:
	;
	v1268 = v1264
	goto L385
L393:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+8)) = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v1306
	v1313 = int32(6106)
	*(*int32)(unsafe.Add(mBase, uint32(l0))) = v1313
	v1794 = v1306
	v1795 = v1313
	v1796 = v1308
	goto L4
L394:
	;
	v1281 = F_relation_openrv_extended(m, v1278, int32(1), l5)
	mBase = m.M
	v1282 = m.ExcPending
	if v1282 != 0 {
		goto L47
	} else {
		goto L395
	}
L395:
	;
	if v1281 != 0 {
		goto L396
	} else {
		goto L397
	}
L396:
	;
	v1283 = *(*int32)(unsafe.Add(mBase, uint32(l2)+12))
	v1284 = *(*int32)(unsafe.Add(mBase, uint32(v1283)+4))
	v1285 = *(*int32)(unsafe.Add(mBase, uint32(v1284)+4))
	v1286 = F_GetPublicationByName(m, v1285, l5)
	mBase = m.M
	v1287 = m.ExcPending
	if v1287 != 0 {
		goto L47
	} else {
		goto L399
	}
L397:
	;
	goto L398
L398:
	;
	v1304 = int32(0)
	v1306 = v1304
	v1308 = v1304
	goto L393
L399:
	;
	if v1286 != 0 {
		goto L400
	} else {
		goto L401
	}
L400:
	;
	v1289 = *(*int32)(unsafe.Add(mBase, uint32(v1281)+56))
	v1290 = *(*int32)(unsafe.Add(mBase, uint32(v1286)))
	v1291 = int32(0)
	v1293 = F_GetSysCacheOid(m, int32(53), v1289, v1290, v1291, v1291)
	mBase = m.M
	v1294 = m.ExcPending
	if v1294 != 0 {
		goto L47
	} else {
		goto L403
	}
L401:
	;
	goto L402
L402:
	;
	F_relation_close(m, v1281, int32(1))
	mBase = m.M
	v1300 = m.ExcPending
	if v1300 != 0 {
		goto L47
	} else {
		goto L406
	}
L403:
	;
	if v1293 != 0 {
		v1306 = v1293
		v1308 = v1281
		goto L393
	} else {
		goto L404
	}
L404:
	;
	if l5 == int32(0) {
		goto L12
	} else {
		goto L405
	}
L405:
	;
	goto L402
L406:
	;
	goto L398
L407:
	;
	v1322 = *(*int32)(unsafe.Add(mBase, uint32(v1316)+8))
	v1323 = *(*int32)(unsafe.Add(mBase, uint32(v1322)+4))
	v1324 = v1323
	goto L409
L408:
	;
	v1324 = int32(0)
	goto L409
L409:
	;
	v1325 = *(*int32)(unsafe.Add(mBase, uint32(v1317)+4))
	v1327 = *(*int32)(unsafe.Add(mBase, uint32(v1316)))
	v1328 = *(*int32)(unsafe.Add(mBase, uint32(v1327)+4))
	v1329 = int32(*(*int8)(unsafe.Add(mBase, uint32(v1328))))
	v1331 = v1329 & int32(255)
	switch v1331 - int32(76) {
	case 0:
		goto L413
	default:
		goto L412
	case 7:
		goto L411
	case 8:
		goto L415
	case 26:
		goto L416
	case 34:
		goto L414
	case 38:
		v1368 = int32(166656)
		goto L410
	}
L410:
	;
	v1370 = F_SearchSysCache1(m, int32(10), v1325)
	mBase = m.M
	v1371 = m.ExcPending
	if v1371 != 0 {
		goto L47
	} else {
		goto L423
	}
L411:
	;
	v1368 = int32(171416)
	goto L410
L412:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v1341 = m.ExcPending
	if v1341 != 0 {
		goto L47
	} else {
		goto L417
	}
L413:
	;
	v1368 = int32(125488)
	goto L410
L414:
	;
	v1368 = int32(174307)
	goto L410
L415:
	;
	v1368 = int32(162593)
	goto L410
L416:
	;
	v1368 = int32(141385)
	goto L410
L417:
	;
	F_errcode(m, int32(50856066))
	mBase = m.M
	v1344 = m.ExcPending
	if v1344 != 0 {
		goto L47
	} else {
		goto L418
	}
L418:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v24)+368)) = v1329
	F_errmsg(m, int32(729563), v24+int32(368))
	mBase = m.M
	v1350 = m.ExcPending
	if v1350 != 0 {
		goto L47
	} else {
		goto L419
	}
L419:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v24)+352)) = int64(326417514606)
	*(*int64)(unsafe.Add(mBase, uint32(v24)+344)) = int64(360777252966)
	*(*int64)(unsafe.Add(mBase, uint32(v24)+336)) = int64(356482285682)
	F_errhint(m, int32(669156), v24+int32(336))
	mBase = m.M
	v1361 = m.ExcPending
	if v1361 != 0 {
		goto L47
	} else {
		goto L420
	}
L420:
	;
	F_errfinish(m, int32(493656), int32(2021), int32(308377))
	mBase = m.M
	v1366 = m.ExcPending
	if v1366 != 0 {
		goto L47
	} else {
		goto L421
	}
L421:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L422:
	;
	if l5 != 0 {
		v1655 = int32(0)
		goto L10
	} else {
		goto L435
	}
L423:
	;
	if v1370 == int32(0) {
		goto L422
	} else {
		goto L424
	}
L424:
	;
	v1374 = *(*int32)(unsafe.Add(mBase, uint32(v1370)+16))
	v1375 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1374)+22)))
	v1377 = *(*int32)(unsafe.Add(mBase, uint32(v1374+v1375)))
	F_ReleaseCatCache(m, v1370)
	mBase = m.M
	v1379 = m.ExcPending
	if v1379 != 0 {
		goto L47
	} else {
		goto L425
	}
L425:
	;
	if v1324 == int32(0) {
		goto L427
	} else {
		goto L428
	}
L426:
	;
	v1390 = F_SearchSysCache3(m, int32(22), v1377, v1388, v1331)
	mBase = m.M
	v1391 = m.ExcPending
	if v1391 != 0 {
		goto L47
	} else {
		goto L432
	}
L427:
	;
	v1388 = int32(0)
	goto L426
L428:
	;
	goto L429
L429:
	;
	v1384 = F_get_namespace_oid(m, v1324, int32(1))
	mBase = m.M
	v1385 = m.ExcPending
	if v1385 != 0 {
		goto L47
	} else {
		goto L430
	}
L430:
	;
	if v1384 == int32(0) {
		goto L422
	} else {
		goto L431
	}
L431:
	;
	v1388 = v1384
	goto L426
L432:
	;
	if v1390 == int32(0) {
		goto L422
	} else {
		goto L433
	}
L433:
	;
	v1394 = *(*int32)(unsafe.Add(mBase, uint32(v1390)+16))
	v1395 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1394)+22)))
	v1397 = *(*int32)(unsafe.Add(mBase, uint32(v1394+v1395)))
	F_ReleaseCatCache(m, v1390)
	mBase = m.M
	v1399 = m.ExcPending
	if v1399 != 0 {
		goto L47
	} else {
		goto L434
	}
L434:
	;
	v1655 = v1397
	goto L10
L435:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v1407 = m.ExcPending
	if v1407 != 0 {
		goto L47
	} else {
		goto L436
	}
L436:
	;
	F_errcode(m, int32(67137668))
	mBase = m.M
	v1410 = m.ExcPending
	if v1410 != 0 {
		goto L47
	} else {
		goto L437
	}
L437:
	;
	if v1324 != 0 {
		goto L11
	} else {
		goto L438
	}
L438:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v24)+388)) = v1368
	*(*int32)(unsafe.Add(mBase, uint32(v24)+384)) = v1325
	F_errmsg(m, int32(69804), v24+int32(384))
	mBase = m.M
	v1417 = m.ExcPending
	if v1417 != 0 {
		goto L47
	} else {
		goto L439
	}
L439:
	;
	F_errfinish(m, int32(493656), int32(2073), int32(308377))
	mBase = m.M
	v1422 = m.ExcPending
	if v1422 != 0 {
		goto L47
	} else {
		goto L440
	}
L440:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L441:
	;
	v1468 = v1423
	v1482 = v1426
	goto L20
L442:
	;
	v1430 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v1431 = v1430
	goto L22
L443:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v1435 = m.ExcPending
	if v1435 != 0 {
		goto L47
	} else {
		goto L444
	}
L444:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v24))) = l1
	F_errmsg_internal(m, int32(484755), v24)
	mBase = m.M
	v1439 = m.ExcPending
	if v1439 != 0 {
		goto L47
	} else {
		goto L445
	}
L445:
	;
	F_errfinish(m, int32(493656), int32(1134), int32(128300))
	mBase = m.M
	v1444 = m.ExcPending
	if v1444 != 0 {
		goto L47
	} else {
		goto L446
	}
L446:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L447:
	;
	v1452 = *(*int32)(unsafe.Add(mBase, uint32(v24)+440))
	v1453 = *(*int32)(unsafe.Add(mBase, uint32(l2)+12))
	v1454 = *(*int32)(unsafe.Add(mBase, uint32(v1453)+4))
	v1455 = *(*int32)(unsafe.Add(mBase, uint32(v1454)+4))
	v1456 = int32(2606)
	*(*int32)(unsafe.Add(mBase, uint32(l0))) = v1456
	v1459 = F_get_domain_constraint_oid(m, v1452, v1455, l5)
	mBase = m.M
	v1460 = m.ExcPending
	if v1460 != 0 {
		goto L47
	} else {
		goto L448
	}
L448:
	;
	v1468 = v1456
	v1482 = v1459
	goto L20
L449:
	;
	F_errcode(m, int32(16801924))
	mBase = m.M
	v1516 = m.ExcPending
	if v1516 != 0 {
		goto L47
	} else {
		goto L450
	}
L450:
	;
	F_errmsg(m, int32(456662), int32(0))
	mBase = m.M
	v1520 = m.ExcPending
	if v1520 != 0 {
		goto L47
	} else {
		goto L451
	}
L451:
	;
	F_errfinish(m, int32(493656), int32(1514), int32(348765))
	mBase = m.M
	v1525 = m.ExcPending
	if v1525 != 0 {
		goto L47
	} else {
		goto L452
	}
L452:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L453:
	;
	F_errcode(m, int32(50360452))
	mBase = m.M
	v1532 = m.ExcPending
	if v1532 != 0 {
		goto L47
	} else {
		goto L454
	}
L454:
	;
	v1533 = F_NameListToString(m, v246)
	mBase = m.M
	v1534 = m.ExcPending
	if v1534 != 0 {
		goto L47
	} else {
		goto L455
	}
L455:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v24)+132)) = v1533
	*(*int32)(unsafe.Add(mBase, uint32(v24)+128)) = v243
	F_errmsg(m, int32(71602), v24+int32(128))
	mBase = m.M
	v1541 = m.ExcPending
	if v1541 != 0 {
		goto L47
	} else {
		goto L456
	}
L456:
	;
	F_errfinish(m, int32(493656), int32(1529), int32(348765))
	mBase = m.M
	v1546 = m.ExcPending
	if v1546 != 0 {
		goto L47
	} else {
		goto L457
	}
L457:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L458:
	;
	F_errcode(m, int32(16801924))
	mBase = m.M
	v1554 = m.ExcPending
	if v1554 != 0 {
		goto L47
	} else {
		goto L459
	}
L459:
	;
	F_errmsg(m, int32(456662), int32(0))
	mBase = m.M
	v1558 = m.ExcPending
	if v1558 != 0 {
		goto L47
	} else {
		goto L460
	}
L460:
	;
	F_errfinish(m, int32(493656), int32(1567), int32(339431))
	mBase = m.M
	v1563 = m.ExcPending
	if v1563 != 0 {
		goto L47
	} else {
		goto L461
	}
L461:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L462:
	;
	F_errcode(m, int32(50360452))
	mBase = m.M
	v1570 = m.ExcPending
	if v1570 != 0 {
		goto L47
	} else {
		goto L463
	}
L463:
	;
	v1571 = F_NameListToString(m, v283)
	mBase = m.M
	v1572 = m.ExcPending
	if v1572 != 0 {
		goto L47
	} else {
		goto L464
	}
L464:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v24)+148)) = v1571
	*(*int32)(unsafe.Add(mBase, uint32(v24)+144)) = v280
	F_errmsg(m, int32(71584), v24+int32(144))
	mBase = m.M
	v1579 = m.ExcPending
	if v1579 != 0 {
		goto L47
	} else {
		goto L465
	}
L465:
	;
	F_errfinish(m, int32(493656), int32(1587), int32(339431))
	mBase = m.M
	v1584 = m.ExcPending
	if v1584 != 0 {
		goto L47
	} else {
		goto L466
	}
L466:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L467:
	;
	F_errcode(m, int32(16801924))
	mBase = m.M
	v1591 = m.ExcPending
	if v1591 != 0 {
		goto L47
	} else {
		goto L468
	}
L468:
	;
	F_errmsg(m, int32(381277), int32(0))
	mBase = m.M
	v1595 = m.ExcPending
	if v1595 != 0 {
		goto L47
	} else {
		goto L469
	}
L469:
	;
	F_errfinish(m, int32(493656), int32(1438), int32(110302))
	mBase = m.M
	v1600 = m.ExcPending
	if v1600 != 0 {
		goto L47
	} else {
		goto L470
	}
L470:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L471:
	;
	F_errcode(m, int32(67137668))
	mBase = m.M
	v1607 = m.ExcPending
	if v1607 != 0 {
		goto L47
	} else {
		goto L472
	}
L472:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v24)+308)) = v1247
	*(*int32)(unsafe.Add(mBase, uint32(v24)+304)) = v1251
	F_errmsg(m, int32(71707), v24+int32(304))
	mBase = m.M
	v1614 = m.ExcPending
	if v1614 != 0 {
		goto L47
	} else {
		goto L473
	}
L473:
	;
	F_errfinish(m, int32(493656), int32(1954), int32(506096))
	mBase = m.M
	v1619 = m.ExcPending
	if v1619 != 0 {
		goto L47
	} else {
		goto L474
	}
L474:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L475:
	;
	F_errcode(m, int32(67137668))
	mBase = m.M
	v1626 = m.ExcPending
	if v1626 != 0 {
		goto L47
	} else {
		goto L476
	}
L476:
	;
	v1627 = *(*int32)(unsafe.Add(mBase, uint32(v1281)+48))
	*(*int32)(unsafe.Add(mBase, uint32(v24)+324)) = v1285
	*(*int32)(unsafe.Add(mBase, uint32(v24)+320)) = v1627 + int32(4)
	F_errmsg(m, int32(71646), v24+int32(320))
	mBase = m.M
	v1636 = m.ExcPending
	if v1636 != 0 {
		goto L47
	} else {
		goto L477
	}
L477:
	;
	F_errfinish(m, int32(493656), int32(1907), int32(307496))
	mBase = m.M
	v1641 = m.ExcPending
	if v1641 != 0 {
		goto L47
	} else {
		goto L478
	}
L478:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L479:
	;
	F_errfinish(m, int32(493656), int32(2068), int32(308377))
	mBase = m.M
	v1654 = m.ExcPending
	if v1654 != 0 {
		goto L47
	} else {
		goto L480
	}
L480:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L481:
	;
	if l5 != 0 {
		goto L501
	} else {
		goto L502
	}
L482:
	;
	v1720 = *(*int32)(unsafe.Add(mBase, uint32(v1718)+16))
	v1721 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1720)+22)))
	v1723 = *(*int32)(unsafe.Add(mBase, uint32(v1720+v1721)))
	F_ReleaseCatCache(m, v1718)
	mBase = m.M
	v1725 = m.ExcPending
	if v1725 != 0 {
		goto L47
	} else {
		goto L500
	}
L483:
	;
	v1678 = F_SearchSysCache4(m, int32(4), v943, v1669, v1672, v1673)
	mBase = m.M
	v1679 = m.ExcPending
	if v1679 != 0 {
		goto L47
	} else {
		goto L486
	}
L484:
	;
	goto L485
L485:
	;
	v1713 = F_SearchSysCache4(m, int32(5), v943, v1669, v1672, v1673)
	mBase = m.M
	v1714 = m.ExcPending
	if v1714 != 0 {
		goto L47
	} else {
		goto L498
	}
L486:
	;
	if v1678 != 0 {
		v1718 = v1678
		v1719 = int32(2602)
		goto L482
	} else {
		goto L487
	}
L487:
	;
	if l5 != 0 {
		goto L488
	} else {
		goto L489
	}
L488:
	;
	v1758 = int32(0)
	v1759 = int32(2602)
	goto L7
L489:
	;
	goto L490
L490:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v1685 = m.ExcPending
	if v1685 != 0 {
		goto L47
	} else {
		goto L491
	}
L491:
	;
	F_errcode(m, int32(67137668))
	mBase = m.M
	v1688 = m.ExcPending
	if v1688 != 0 {
		goto L47
	} else {
		goto L492
	}
L492:
	;
	v1689 = F_TypeNameToString(m, v1671)
	mBase = m.M
	v1690 = m.ExcPending
	if v1690 != 0 {
		goto L47
	} else {
		goto L493
	}
L493:
	;
	v1691 = F_TypeNameToString(m, v1670)
	mBase = m.M
	v1692 = m.ExcPending
	if v1692 != 0 {
		goto L47
	} else {
		goto L494
	}
L494:
	;
	v1696 = F_getObjectDescription(m, v24+int32(436), int32(0))
	mBase = m.M
	v1697 = m.ExcPending
	if v1697 != 0 {
		goto L47
	} else {
		goto L495
	}
L495:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v24)+220)) = v1696
	*(*int32)(unsafe.Add(mBase, uint32(v24)+216)) = v1691
	*(*int32)(unsafe.Add(mBase, uint32(v24)+212)) = v1689
	*(*int32)(unsafe.Add(mBase, uint32(v24)+208)) = v928
	F_errmsg(m, int32(69958), v24+int32(208))
	mBase = m.M
	v1706 = m.ExcPending
	if v1706 != 0 {
		goto L47
	} else {
		goto L496
	}
L496:
	;
	F_errfinish(m, int32(493656), int32(1746), int32(228541))
	mBase = m.M
	v1711 = m.ExcPending
	if v1711 != 0 {
		goto L47
	} else {
		goto L497
	}
L497:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L498:
	;
	if v1713 == int32(0) {
		goto L481
	} else {
		goto L499
	}
L499:
	;
	v1718 = v1713
	v1719 = int32(2603)
	goto L482
L500:
	;
	v1758 = v1723
	v1759 = v1719
	goto L7
L501:
	;
	v1758 = int32(0)
	v1759 = int32(2603)
	goto L7
L502:
	;
	goto L503
L503:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v1731 = m.ExcPending
	if v1731 != 0 {
		goto L47
	} else {
		goto L504
	}
L504:
	;
	F_errcode(m, int32(67137668))
	mBase = m.M
	v1734 = m.ExcPending
	if v1734 != 0 {
		goto L47
	} else {
		goto L505
	}
L505:
	;
	v1735 = F_TypeNameToString(m, v1671)
	mBase = m.M
	v1736 = m.ExcPending
	if v1736 != 0 {
		goto L47
	} else {
		goto L506
	}
L506:
	;
	v1737 = F_TypeNameToString(m, v1670)
	mBase = m.M
	v1738 = m.ExcPending
	if v1738 != 0 {
		goto L47
	} else {
		goto L507
	}
L507:
	;
	v1742 = F_getObjectDescription(m, v24+int32(436), int32(0))
	mBase = m.M
	v1743 = m.ExcPending
	if v1743 != 0 {
		goto L47
	} else {
		goto L508
	}
L508:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v24)+236)) = v1742
	*(*int32)(unsafe.Add(mBase, uint32(v24)+232)) = v1737
	*(*int32)(unsafe.Add(mBase, uint32(v24)+228)) = v1735
	*(*int32)(unsafe.Add(mBase, uint32(v24)+224)) = v928
	F_errmsg(m, int32(70000), v24+int32(224))
	mBase = m.M
	v1752 = m.ExcPending
	if v1752 != 0 {
		goto L47
	} else {
		goto L509
	}
L509:
	;
	F_errfinish(m, int32(493656), int32(1777), int32(228541))
	mBase = m.M
	v1757 = m.ExcPending
	if v1757 != 0 {
		goto L47
	} else {
		goto L510
	}
L510:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L511:
	;
	if v46 == int32(0) {
		goto L513
	} else {
		goto L514
	}
L512:
	;
	goto L3
L513:
	;
	if v1795 == int32(1259) {
		goto L553
	} else {
		goto L554
	}
L514:
	;
	if v1795 != v46 {
		goto L515
	} else {
		goto L516
	}
L515:
	;
	if v46 == int32(1259) {
		goto L513
	} else {
		goto L519
	}
L516:
	;
	if v1794 != v49 {
		goto L515
	} else {
		goto L517
	}
L517:
	;
	v1815 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	if v50 == v1815 {
		goto L512
	} else {
		goto L518
	}
L518:
	;
	goto L515
L519:
	;
	v1821 = int32(1)
	if v46 <= int32(3591) {
		goto L524
	} else {
		goto L525
	}
L520:
	;
	if v1889 != 0 {
		goto L548
	} else {
		goto L549
	}
L521:
	;
	goto L520
L522:
	;
	v1889 = int32(0)
	goto L521
L523:
	;
	if base.Ui32(v46-int32(2964)) < base.Ui32(int32(4)) {
		v1889 = v1821
		goto L521
	} else {
		goto L546
	}
L524:
	;
	if v46 <= int32(2670) {
		goto L527
	} else {
		goto L528
	}
L525:
	;
	goto L526
L526:
	;
	if v46 <= int32(5999) {
		goto L535
	} else {
		goto L536
	}
L527:
	;
	switch v46 - int32(1213) {
	case 0, 1, 19, 20, 47, 48, 49:
		v1889 = v1821
		goto L521
	case 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15, 16, 17, 18, 21, 22, 23, 24, 25, 26, 27, 28, 29, 30, 31, 32, 33, 34, 35, 36, 37, 38, 39, 40, 41, 42, 43, 44, 45, 46:
		goto L522
	default:
		goto L530
	}
L528:
	;
	goto L529
L529:
	;
	v1833 = v46 - int32(2671)
	if base.Ui32(int32(27)) < base.Ui32(v1833) {
		goto L523
	} else {
		goto L532
	}
L530:
	;
	if base.Ui32(int32(2)) <= base.Ui32(v46-int32(2396)) {
		goto L522
	} else {
		goto L531
	}
L531:
	;
	v1889 = v1821
	goto L521
L532:
	;
	if int32(1)<<(uint(v1833)%32)&int32(226492515) == int32(0) {
		goto L523
	} else {
		goto L533
	}
L533:
	;
	v1889 = v1821
	goto L521
L534:
	;
	if base.Ui32(v46-int32(3592)) < base.Ui32(int32(2)) {
		v1889 = v1821
		goto L521
	} else {
		goto L544
	}
L535:
	;
	v1845 = v46 - int32(4177)
	if base.Ui32(int32(9)) < base.Ui32(v1845) {
		goto L534
	} else {
		goto L538
	}
L536:
	;
	goto L537
L537:
	;
	switch v46 - int32(6243) {
	case 0, 1, 2, 3, 4, 59, 60:
		v1889 = v1821
		goto L521
	case 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15, 16, 17, 18, 19, 20, 21, 22, 23, 24, 25, 26, 27, 28, 29, 30, 31, 32, 33, 34, 35, 36, 37, 38, 39, 40, 41, 42, 43, 44, 45, 46, 47, 48, 49, 50, 51, 52, 53, 54, 55, 56, 57, 58:
		goto L522
	default:
		goto L540
	}
L538:
	;
	if int32(1)<<(uint(v1845)%32)&int32(963) == int32(0) {
		goto L534
	} else {
		goto L539
	}
L539:
	;
	v1889 = v1821
	goto L521
L540:
	;
	if base.Ui32(v46-int32(6000)) < base.Ui32(int32(3)) {
		v1889 = v1821
		goto L521
	} else {
		goto L541
	}
L541:
	;
	v1861 = v46 - int32(6100)
	if base.Ui32(int32(15)) < base.Ui32(v1861) {
		goto L522
	} else {
		goto L542
	}
L542:
	;
	if int32(1)<<(uint(v1861)%32)&int32(49153) != 0 {
		v1889 = v1821
		goto L521
	} else {
		goto L543
	}
L543:
	;
	goto L522
L544:
	;
	if base.Ui32(int32(2)) <= base.Ui32(v46-int32(4060)) {
		goto L522
	} else {
		goto L545
	}
L545:
	;
	v1889 = v1821
	goto L521
L546:
	;
	if base.Ui32(v46-int32(2846)) < base.Ui32(int32(2)) {
		v1889 = v1821
		goto L521
	} else {
		goto L547
	}
L547:
	;
	goto L522
L548:
	;
	F_UnlockSharedObject(m, v46, v49, l4)
	mBase = m.M
	v1891 = m.ExcPending
	if v1891 != 0 {
		goto L47
	} else {
		goto L551
	}
L549:
	;
	goto L550
L550:
	;
	F_UnlockDatabaseObject(m, v46, v49, l4)
	mBase = m.M
	v1893 = m.ExcPending
	if v1893 != 0 {
		goto L47
	} else {
		goto L552
	}
L551:
	;
	goto L513
L552:
	;
	goto L513
L553:
	;
	if v1796 != 0 {
		goto L512
	} else {
		goto L588
	}
L554:
	;
	v1898 = int32(1)
	if v1795 <= int32(3591) {
		goto L559
	} else {
		goto L560
	}
L555:
	;
	if v1966 != 0 {
		goto L583
	} else {
		goto L584
	}
L556:
	;
	goto L555
L557:
	;
	v1966 = int32(0)
	goto L556
L558:
	;
	if base.Ui32(v1795-int32(2964)) < base.Ui32(int32(4)) {
		v1966 = v1898
		goto L556
	} else {
		goto L581
	}
L559:
	;
	if v1795 <= int32(2670) {
		goto L562
	} else {
		goto L563
	}
L560:
	;
	goto L561
L561:
	;
	if v1795 <= int32(5999) {
		goto L570
	} else {
		goto L571
	}
L562:
	;
	switch v1795 - int32(1213) {
	case 0, 1, 19, 20, 47, 48, 49:
		v1966 = v1898
		goto L556
	case 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15, 16, 17, 18, 21, 22, 23, 24, 25, 26, 27, 28, 29, 30, 31, 32, 33, 34, 35, 36, 37, 38, 39, 40, 41, 42, 43, 44, 45, 46:
		goto L557
	default:
		goto L565
	}
L563:
	;
	goto L564
L564:
	;
	v1910 = v1795 - int32(2671)
	if base.Ui32(int32(27)) < base.Ui32(v1910) {
		goto L558
	} else {
		goto L567
	}
L565:
	;
	if base.Ui32(int32(2)) <= base.Ui32(v1795-int32(2396)) {
		goto L557
	} else {
		goto L566
	}
L566:
	;
	v1966 = v1898
	goto L556
L567:
	;
	if int32(1)<<(uint(v1910)%32)&int32(226492515) == int32(0) {
		goto L558
	} else {
		goto L568
	}
L568:
	;
	v1966 = v1898
	goto L556
L569:
	;
	if base.Ui32(v1795-int32(3592)) < base.Ui32(int32(2)) {
		v1966 = v1898
		goto L556
	} else {
		goto L579
	}
L570:
	;
	v1922 = v1795 - int32(4177)
	if base.Ui32(int32(9)) < base.Ui32(v1922) {
		goto L569
	} else {
		goto L573
	}
L571:
	;
	goto L572
L572:
	;
	switch v1795 - int32(6243) {
	case 0, 1, 2, 3, 4, 59, 60:
		v1966 = v1898
		goto L556
	case 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15, 16, 17, 18, 19, 20, 21, 22, 23, 24, 25, 26, 27, 28, 29, 30, 31, 32, 33, 34, 35, 36, 37, 38, 39, 40, 41, 42, 43, 44, 45, 46, 47, 48, 49, 50, 51, 52, 53, 54, 55, 56, 57, 58:
		goto L557
	default:
		goto L575
	}
L573:
	;
	if int32(1)<<(uint(v1922)%32)&int32(963) == int32(0) {
		goto L569
	} else {
		goto L574
	}
L574:
	;
	v1966 = v1898
	goto L556
L575:
	;
	if base.Ui32(v1795-int32(6000)) < base.Ui32(int32(3)) {
		v1966 = v1898
		goto L556
	} else {
		goto L576
	}
L576:
	;
	v1938 = v1795 - int32(6100)
	if base.Ui32(int32(15)) < base.Ui32(v1938) {
		goto L557
	} else {
		goto L577
	}
L577:
	;
	if int32(1)<<(uint(v1938)%32)&int32(49153) != 0 {
		v1966 = v1898
		goto L556
	} else {
		goto L578
	}
L578:
	;
	goto L557
L579:
	;
	if base.Ui32(int32(2)) <= base.Ui32(v1795-int32(4060)) {
		goto L557
	} else {
		goto L580
	}
L580:
	;
	v1966 = v1898
	goto L556
L581:
	;
	if base.Ui32(v1795-int32(2846)) < base.Ui32(int32(2)) {
		v1966 = v1898
		goto L556
	} else {
		goto L582
	}
L582:
	;
	goto L557
L583:
	;
	F_LockSharedObject(m, v1795, v1794, l4)
	mBase = m.M
	v1968 = m.ExcPending
	if v1968 != 0 {
		goto L47
	} else {
		goto L586
	}
L584:
	;
	goto L585
L585:
	;
	F_LockDatabaseObject(m, v1795, v1794, l4)
	mBase = m.M
	v1970 = m.ExcPending
	if v1970 != 0 {
		goto L47
	} else {
		goto L587
	}
L586:
	;
	goto L553
L587:
	;
	goto L553
L588:
	;
	v1972 = *(*int64)(unsafe.Add(mBase, _consts[315]))
	if v51 == v1972 {
		goto L512
	} else {
		goto L589
	}
L589:
	;
	v1974 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	v39 = v1795
	v46 = v1795
	v49 = v1794
	v50 = v1974
	v51 = v1972
	goto L2
L590:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l3))) = v1796
	goto L1
}
func F_get_object_attnum_oid(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v5 int32
	_ = v5
	var v7 int32
	_ = v7
	var v10 int32
	_ = v10
	var v11 int32
	_ = v11
	var v16 int32
	_ = v16
	var v19 int32
	_ = v19
	var v22 int32
	_ = v22
	var v29 int32
	_ = v29
	var v32 int32
	_ = v32
	var v39 int32
	_ = v39
	var v42 int32
	_ = v42
	var v49 int32
	_ = v49
	var v52 int32
	_ = v52
	var v61 int32
	_ = v61
	var v64 int32
	_ = v64
	var v67 int32
	_ = v67
	var v77 int32
	_ = v77
	var v81 int32
	_ = v81
	var v86 int32
	_ = v86
	v5 = m.G0
	v7 = v5 - int32(16)
	m.G0 = v7
	v10 = *(*int32)(unsafe.Add(mBase, _consts[325]))
	if v10 != 0 {
		goto L3
	} else {
		goto L4
	}
L1:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v77 = m.ExcPending
	if v77 != 0 {
		goto L23
	} else {
		goto L24
	}
L2:
	;
	v67 = int32(*(*int16)(unsafe.Add(mBase, uint32(v64)+20)))
	m.G0 = v7 + int32(16)
	return v67
L3:
	;
	v11 = *(*int32)(unsafe.Add(mBase, uint32(v10)+4))
	if v11 == l0 {
		v64 = v10
		goto L2
	} else {
		goto L6
	}
L4:
	;
	goto L5
L5:
	;
	v16 = int32(0)
	goto L8
L6:
	;
	goto L5
L7:
	;
	*(*int32)(unsafe.Add(mBase, _consts[325])) = v61
	v64 = v61
	goto L2
L8:
	;
	v19 = v16 * int32(40)
	v22 = *(*int32)(unsafe.Add(mBase, uint32(v19)+uint32(_consts[326])))
	if l0 != v22 {
		goto L10
	} else {
		goto L11
	}
L9:
	;
	v61 = v19 + int32(766464)
	goto L7
L10:
	;
	if v16 == int32(36) {
		goto L1
	} else {
		goto L13
	}
L11:
	;
	goto L12
L12:
	;
	goto L9
L13:
	;
	v29 = (v16 | int32(1)) * int32(40)
	v32 = *(*int32)(unsafe.Add(mBase, uint32(v29)+uint32(_consts[326])))
	if l0 == v32 {
		goto L14
	} else {
		goto L15
	}
L14:
	;
	v61 = v29 + int32(766464)
	goto L7
L15:
	;
	goto L16
L16:
	;
	v39 = (v16 | int32(2)) * int32(40)
	v42 = *(*int32)(unsafe.Add(mBase, uint32(v39)+uint32(_consts[326])))
	if l0 == v42 {
		goto L17
	} else {
		goto L18
	}
L17:
	;
	v61 = v39 + int32(766464)
	goto L7
L18:
	;
	goto L19
L19:
	;
	v49 = (v16 | int32(3)) * int32(40)
	v52 = *(*int32)(unsafe.Add(mBase, uint32(v49)+uint32(_consts[326])))
	if l0 == v52 {
		goto L20
	} else {
		goto L21
	}
L20:
	;
	v61 = v49 + int32(766464)
	goto L7
L21:
	;
	v16 = v16 + int32(4)
	goto L8
L23:
	;
	return int32(0)
L24:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v7))) = l0
	F_errmsg_internal(m, int32(59363), v7)
	mBase = m.M
	v81 = m.ExcPending
	if v81 != 0 {
		goto L23
	} else {
		goto L25
	}
L25:
	;
	F_errfinish(m, int32(493656), int32(2777), int32(503926))
	mBase = m.M
	v86 = m.ExcPending
	if v86 != 0 {
		goto L23
	} else {
		goto L26
	}
L26:
	;
	base.Wasm_trap_unreachable()
	for {
	}
}
func F_get_object_catcache_name(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v5 int32
	_ = v5
	var v7 int32
	_ = v7
	var v10 int32
	_ = v10
	var v11 int32
	_ = v11
	var v16 int32
	_ = v16
	var v19 int32
	_ = v19
	var v22 int32
	_ = v22
	var v29 int32
	_ = v29
	var v32 int32
	_ = v32
	var v39 int32
	_ = v39
	var v42 int32
	_ = v42
	var v49 int32
	_ = v49
	var v52 int32
	_ = v52
	var v61 int32
	_ = v61
	var v64 int32
	_ = v64
	var v67 int32
	_ = v67
	var v77 int32
	_ = v77
	var v81 int32
	_ = v81
	var v86 int32
	_ = v86
	v5 = m.G0
	v7 = v5 - int32(16)
	m.G0 = v7
	v10 = *(*int32)(unsafe.Add(mBase, _consts[325]))
	if v10 != 0 {
		goto L3
	} else {
		goto L4
	}
L1:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v77 = m.ExcPending
	if v77 != 0 {
		goto L23
	} else {
		goto L24
	}
L2:
	;
	v67 = *(*int32)(unsafe.Add(mBase, uint32(v64)+16))
	m.G0 = v7 + int32(16)
	return v67
L3:
	;
	v11 = *(*int32)(unsafe.Add(mBase, uint32(v10)+4))
	if v11 == l0 {
		v64 = v10
		goto L2
	} else {
		goto L6
	}
L4:
	;
	goto L5
L5:
	;
	v16 = int32(0)
	goto L8
L6:
	;
	goto L5
L7:
	;
	*(*int32)(unsafe.Add(mBase, _consts[325])) = v61
	v64 = v61
	goto L2
L8:
	;
	v19 = v16 * int32(40)
	v22 = *(*int32)(unsafe.Add(mBase, uint32(v19)+uint32(_consts[326])))
	if l0 != v22 {
		goto L10
	} else {
		goto L11
	}
L9:
	;
	v61 = v19 + int32(766464)
	goto L7
L10:
	;
	if v16 == int32(36) {
		goto L1
	} else {
		goto L13
	}
L11:
	;
	goto L12
L12:
	;
	goto L9
L13:
	;
	v29 = (v16 | int32(1)) * int32(40)
	v32 = *(*int32)(unsafe.Add(mBase, uint32(v29)+uint32(_consts[326])))
	if l0 == v32 {
		goto L14
	} else {
		goto L15
	}
L14:
	;
	v61 = v29 + int32(766464)
	goto L7
L15:
	;
	goto L16
L16:
	;
	v39 = (v16 | int32(2)) * int32(40)
	v42 = *(*int32)(unsafe.Add(mBase, uint32(v39)+uint32(_consts[326])))
	if l0 == v42 {
		goto L17
	} else {
		goto L18
	}
L17:
	;
	v61 = v39 + int32(766464)
	goto L7
L18:
	;
	goto L19
L19:
	;
	v49 = (v16 | int32(3)) * int32(40)
	v52 = *(*int32)(unsafe.Add(mBase, uint32(v49)+uint32(_consts[326])))
	if l0 == v52 {
		goto L20
	} else {
		goto L21
	}
L20:
	;
	v61 = v49 + int32(766464)
	goto L7
L21:
	;
	v16 = v16 + int32(4)
	goto L8
L23:
	;
	return int32(0)
L24:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v7))) = l0
	F_errmsg_internal(m, int32(59363), v7)
	mBase = m.M
	v81 = m.ExcPending
	if v81 != 0 {
		goto L23
	} else {
		goto L25
	}
L25:
	;
	F_errfinish(m, int32(493656), int32(2777), int32(503926))
	mBase = m.M
	v86 = m.ExcPending
	if v86 != 0 {
		goto L23
	} else {
		goto L26
	}
L26:
	;
	base.Wasm_trap_unreachable()
	for {
	}
}
func F_get_object_field_end(m *base.Module, l0 int32, l1 int32, l2 int32) int32 {
	mBase := m.M
	_ = mBase
	var v10 int32
	_ = v10
	var v11 int32
	_ = v11
	var v12 int32
	_ = v12
	var v14 int32
	_ = v14
	var v15 int32
	_ = v15
	var v16 int32
	_ = v16
	var v18 int32
	_ = v18
	var v21 int32
	_ = v21
	var v27 int32
	_ = v27
	var v32 int32
	_ = v32
	var v33 int32
	_ = v33
	var v37 int32
	_ = v37
	var v38 int32
	_ = v38
	var v41 int32
	_ = v41
	var v42 int32
	_ = v42
	var v45 int32
	_ = v45
	var v52 int32
	_ = v52
	var v53 int32
	_ = v53
	var v57 int32
	_ = v57
	var v61 int32
	_ = v61
	var v65 int32
	_ = v65
	var v66 int32
	_ = v66
	var v68 int32
	_ = v68
	var v71 int32
	_ = v71
	var v72 int32
	_ = v72
	v10 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v11 = *(*int32)(unsafe.Add(mBase, uint32(v10)+32))
	v12 = *(*int32)(unsafe.Add(mBase, uint32(l0)+16))
	if v12 < v11 {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	return int32(0)
L2:
	;
	v14 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
	v15 = int32(1)
	v16 = v11 - v15
	v18 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v14+v16))))
	if v18 != v15 {
		goto L1
	} else {
		goto L3
	}
L3:
	;
	v21 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	if v21 == int32(0) {
		goto L1
	} else {
		goto L4
	}
L4:
	;
	v27 = *(*int32)(unsafe.Add(mBase, uint32(v21+v16<<(uint(int32(2))%32))))
	if v27 == int32(0) {
		goto L1
	} else {
		goto L5
	}
L5:
	;
	v32 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v27))))
	v33 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l1))))
	if v33 == int32(0) {
		v52 = v32
		v53 = v33
		goto L7
	} else {
		goto L8
	}
L6:
	;
	if v53-v52 != 0 {
		goto L1
	} else {
		goto L14
	}
L7:
	;
	goto L6
L8:
	;
	if v32 != v33 {
		v52 = v32
		v53 = v33
		goto L7
	} else {
		goto L9
	}
L9:
	;
	v37 = l1
	v38 = v27
	goto L10
L10:
	;
	v41 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v38)+1)))
	v42 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v37)+1)))
	if v42 == int32(0) {
		v52 = v41
		v53 = v42
		goto L7
	} else {
		goto L12
	}
L11:
	;
	v52 = v41
	v53 = v42
	goto L7
L12:
	;
	v45 = int32(1)
	if v41 == v42 {
		v37 = v37 + v45
		v38 = v38 + v45
		goto L10
	} else {
		goto L13
	}
L13:
	;
	goto L11
L14:
	;
	if v11 < v12 {
		goto L15
	} else {
		goto L16
	}
L15:
	;
	v57 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v11+v14))) = uint8(v57)
	return v57
L16:
	;
	goto L17
L17:
	;
	v61 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	if v61 == int32(0) {
		goto L1
	} else {
		goto L18
	}
L18:
	;
	if l2 != 0 {
		goto L20
	} else {
		goto L21
	}
L19:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+8)) = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v72
	goto L1
L20:
	;
	v65 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+12)))
	if v65 != 0 {
		v72 = int32(0)
		goto L19
	} else {
		goto L23
	}
L21:
	;
	goto L22
L22:
	;
	v66 = *(*int32)(unsafe.Add(mBase, uint32(v10)+20))
	v68 = F_cstring_to_text_with_len(m, v61, v66-v61)
	mBase = m.M
	v71 = m.ExcPending
	if v71 != 0 {
		goto L24
	} else {
		goto L25
	}
L23:
	;
	goto L22
L24:
	;
	return int32(0)
L25:
	;
	v72 = v68
	goto L19
}
func F_get_object_start(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v3 int32
	_ = v3
	var v4 int32
	_ = v4
	var v5 int32
	_ = v5
	var v6 int32
	_ = v6
	v3 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v4 = *(*int32)(unsafe.Add(mBase, uint32(v3)+32))
	if v4 != 0 {
	} else {
		v5 = *(*int32)(unsafe.Add(mBase, uint32(l0)+16))
		if v5 != 0 {
		} else {
			v6 = *(*int32)(unsafe.Add(mBase, uint32(v3)+12))
			*(*int32)(unsafe.Add(mBase, uint32(l0)+8)) = v6
		}
	}
	return int32(0)
}
func F_get_object_type(m *base.Module, l0 int32, l1 int32) int32 {
	mBase := m.M
	_ = mBase
	var v6 int32
	_ = v6
	var v8 int32
	_ = v8
	var v11 int32
	_ = v11
	var v12 int32
	_ = v12
	var v17 int32
	_ = v17
	var v21 int32
	_ = v21
	var v24 int32
	_ = v24
	var v31 int32
	_ = v31
	var v34 int32
	_ = v34
	var v41 int32
	_ = v41
	var v44 int32
	_ = v44
	var v51 int32
	_ = v51
	var v54 int32
	_ = v54
	var v63 int32
	_ = v63
	var v68 int32
	_ = v68
	var v70 int32
	_ = v70
	var v74 int32
	_ = v74
	var v77 int32
	_ = v77
	var v87 int32
	_ = v87
	var v95 int32
	_ = v95
	var v99 int32
	_ = v99
	var v104 int32
	_ = v104
	v6 = m.G0
	v8 = v6 - int32(16)
	m.G0 = v8
	v11 = *(*int32)(unsafe.Add(mBase, _consts[325]))
	if v11 != 0 {
		goto L3
	} else {
		goto L4
	}
L1:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v95 = m.ExcPending
	if v95 != 0 {
		goto L30
	} else {
		goto L32
	}
L2:
	;
	v70 = *(*int32)(unsafe.Add(mBase, uint32(v68)+32))
	if v70 != int32(41) {
		v87 = v70
		goto L23
	} else {
		goto L24
	}
L3:
	;
	v12 = *(*int32)(unsafe.Add(mBase, uint32(v11)+4))
	if v12 == l0 {
		v68 = v11
		goto L2
	} else {
		goto L6
	}
L4:
	;
	goto L5
L5:
	;
	v17 = int32(0)
	goto L8
L6:
	;
	goto L5
L7:
	;
	*(*int32)(unsafe.Add(mBase, _consts[325])) = v63
	v68 = v63
	goto L2
L8:
	;
	v21 = v17 * int32(40)
	v24 = *(*int32)(unsafe.Add(mBase, uint32(v21)+uint32(_consts[326])))
	if l0 != v24 {
		goto L10
	} else {
		goto L11
	}
L9:
	;
	v63 = v21 + int32(766464)
	goto L7
L10:
	;
	if v17 == int32(36) {
		goto L1
	} else {
		goto L13
	}
L11:
	;
	goto L12
L12:
	;
	goto L9
L13:
	;
	v31 = (v17 | int32(1)) * int32(40)
	v34 = *(*int32)(unsafe.Add(mBase, uint32(v31)+uint32(_consts[326])))
	if l0 == v34 {
		goto L14
	} else {
		goto L15
	}
L14:
	;
	v63 = v31 + int32(766464)
	goto L7
L15:
	;
	goto L16
L16:
	;
	v41 = (v17 | int32(2)) * int32(40)
	v44 = *(*int32)(unsafe.Add(mBase, uint32(v41)+uint32(_consts[326])))
	if l0 == v44 {
		goto L17
	} else {
		goto L18
	}
L17:
	;
	v63 = v41 + int32(766464)
	goto L7
L18:
	;
	goto L19
L19:
	;
	v51 = (v17 | int32(3)) * int32(40)
	v54 = *(*int32)(unsafe.Add(mBase, uint32(v51)+uint32(_consts[326])))
	if l0 == v54 {
		goto L20
	} else {
		goto L21
	}
L20:
	;
	v63 = v51 + int32(766464)
	goto L7
L21:
	;
	v17 = v17 + int32(4)
	goto L8
L23:
	;
	m.G0 = v8 + int32(16)
	return v87
L24:
	;
	v74 = F_get_rel_relkind(m, l1)
	mBase = m.M
	v77 = m.ExcPending
	if v77 != 0 {
		goto L30
	} else {
		goto L31
	}
L25:
	;
	v87 = int32(18)
	goto L23
L26:
	;
	v87 = int32(23)
	goto L23
L27:
	;
	v87 = int32(51)
	goto L23
L28:
	;
	v87 = int32(37)
	goto L23
L29:
	;
	v87 = int32(20)
	goto L23
L30:
	;
	return int32(0)
L31:
	;
	switch v74&int32(255) - int32(73) {
	case 0, 32:
		goto L29
	default:
		v87 = int32(41)
		goto L23
	case 10:
		goto L28
	case 29:
		goto L25
	case 36:
		goto L26
	case 45:
		goto L27
	}
L32:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v8))) = l0
	F_errmsg_internal(m, int32(59363), v8)
	mBase = m.M
	v99 = m.ExcPending
	if v99 != 0 {
		goto L30
	} else {
		goto L33
	}
L33:
	;
	F_errfinish(m, int32(493656), int32(2777), int32(503926))
	mBase = m.M
	v104 = m.ExcPending
	if v104 != 0 {
		goto L30
	} else {
		goto L34
	}
L34:
	;
	base.Wasm_trap_unreachable()
	for {
	}
}
func F_object_ownercheck(m *base.Module, l0 int32, l1 int32, l2 int32) int32 {
	mBase := m.M
	_ = mBase
	var v8 int32
	_ = v8
	var v10 int32
	_ = v10
	var v13 int32
	_ = v13
	var v16 int32
	_ = v16
	var v22 int32
	_ = v22
	var v23 int32
	_ = v23
	var v24 int32
	_ = v24
	var v27 int32
	_ = v27
	var v28 int32
	_ = v28
	var v31 int32
	_ = v31
	var v32 int32
	_ = v32
	var v33 int32
	_ = v33
	var v34 int32
	_ = v34
	var v36 int32
	_ = v36
	var v38 int32
	_ = v38
	var v39 int32
	_ = v39
	var v42 int32
	_ = v42
	var v43 int32
	_ = v43
	var v47 int32
	_ = v47
	var v48 int32
	_ = v48
	var v49 int32
	_ = v49
	var v50 int32
	_ = v50
	var v55 int32
	_ = v55
	var v56 int32
	_ = v56
	var v57 int32
	_ = v57
	var v58 int32
	_ = v58
	var v61 int32
	_ = v61
	var v62 int32
	_ = v62
	var v63 int32
	_ = v63
	var v66 int32
	_ = v66
	var v67 int32
	_ = v67
	var v69 int32
	_ = v69
	var v72 int32
	_ = v72
	var v74 int32
	_ = v74
	var v77 int32
	_ = v77
	var v78 int32
	_ = v78
	var v80 int32
	_ = v80
	var v90 int32
	_ = v90
	var v91 int32
	_ = v91
	var v92 int32
	_ = v92
	var v99 int32
	_ = v99
	var v104 int32
	_ = v104
	var v108 int32
	_ = v108
	var v109 int32
	_ = v109
	var v110 int32
	_ = v110
	var v115 int32
	_ = v115
	var v120 int32
	_ = v120
	v8 = m.G0
	v10 = v8 - int32(80)
	m.G0 = v10
	v13 = F_superuser_arg(m, l2)
	mBase = m.M
	v16 = m.ExcPending
	if v16 != 0 {
		return int32(0)
	} else {
		if v13 == int32(0) {
			if l0 == int32(2613) {
				v22 = int32(2995)
			} else {
				v22 = l0
			}
			v23 = F_get_object_catcache_oid(m, v22)
			mBase = m.M
			v24 = m.ExcPending
			if v24 != 0 {
				return int32(0)
			} else {
				if v23 != int32(-1) {
					v27 = F_SearchSysCache1(m, v23, l1)
					mBase = m.M
					v28 = m.ExcPending
					if v28 != 0 {
						return int32(0)
					} else {
						if v27 == int32(0) {
							F_errstart_cold(m, int32(21), int32(0))
							mBase = m.M
							v90 = m.ExcPending
							if v90 != 0 {
								return int32(0)
							} else {
								v91 = F_get_object_class_descr(m, v22)
								mBase = m.M
								v92 = m.ExcPending
								if v92 != 0 {
									return int32(0)
								} else {
									*(*int32)(unsafe.Add(mBase, uint32(v10)+20)) = l1
									*(*int32)(unsafe.Add(mBase, uint32(v10)+16)) = v91
									F_errmsg_internal(m, int32(42594), v10+int32(16))
									mBase = m.M
									v99 = m.ExcPending
									if v99 != 0 {
										return int32(0)
									} else {
										F_errfinish(m, int32(497591), int32(4113), int32(317696))
										mBase = m.M
										v104 = m.ExcPending
										if v104 != 0 {
											return int32(0)
										} else {
											base.Wasm_trap_unreachable()
											for {
											}
										}
									}
								}
							}
						} else {
							v31 = F_get_object_attnum_owner(m, v22)
							mBase = m.M
							v32 = m.ExcPending
							if v32 != 0 {
								return int32(0)
							} else {
								v33 = F_SysCacheGetAttrNotNull(m, v23, v27, v31)
								mBase = m.M
								v34 = m.ExcPending
								if v34 != 0 {
									return int32(0)
								} else {
									F_ReleaseCatCache(m, v27)
									mBase = m.M
									v36 = m.ExcPending
									if v36 != 0 {
										return int32(0)
									} else {
										v74 = v33
										v77 = F_has_privs_of_role(m, l2, v74)
										mBase = m.M
										v78 = m.ExcPending
										if v78 != 0 {
											return int32(0)
										} else {
											v80 = v77
											m.G0 = v10 + int32(80)
											return v80
										}
									}
								}
							}
						}
					}
				} else {
					v38 = F_table_open(m, v22, int32(1))
					mBase = m.M
					v39 = m.ExcPending
					if v39 != 0 {
						return int32(0)
					} else {
						v42 = F_get_object_attnum_oid(m, v22)
						mBase = m.M
						v43 = m.ExcPending
						if v43 != 0 {
							return int32(0)
						} else {
							F_ScanKeyInit(m, v10+int32(32), v42, int32(3), int32(184), l1)
							mBase = m.M
							v47 = m.ExcPending
							if v47 != 0 {
								return int32(0)
							} else {
								v48 = F_get_object_oid_index(m, v22)
								mBase = m.M
								v49 = m.ExcPending
								if v49 != 0 {
									return int32(0)
								} else {
									v50 = int32(1)
									v55 = F_systable_beginscan(m, v38, v48, v50, int32(0), v50, v10+int32(32))
									mBase = m.M
									v56 = m.ExcPending
									if v56 != 0 {
										return int32(0)
									} else {
										v57 = F_systable_getnext(m, v55)
										mBase = m.M
										v58 = m.ExcPending
										if v58 != 0 {
											return int32(0)
										} else {
											if v57 == int32(0) {
												F_errstart_cold(m, int32(21), int32(0))
												mBase = m.M
												v108 = m.ExcPending
												if v108 != 0 {
													return int32(0)
												} else {
													v109 = F_get_object_class_descr(m, v22)
													mBase = m.M
													v110 = m.ExcPending
													if v110 != 0 {
														return int32(0)
													} else {
														*(*int32)(unsafe.Add(mBase, uint32(v10)+4)) = l1
														*(*int32)(unsafe.Add(mBase, uint32(v10))) = v109
														F_errmsg_internal(m, int32(42563), v10)
														mBase = m.M
														v115 = m.ExcPending
														if v115 != 0 {
															return int32(0)
														} else {
															F_errfinish(m, int32(497591), int32(4143), int32(317696))
															mBase = m.M
															v120 = m.ExcPending
															if v120 != 0 {
																return int32(0)
															} else {
																base.Wasm_trap_unreachable()
																for {
																}
															}
														}
													}
												}
											} else {
												v61 = F_get_object_attnum_owner(m, v22)
												mBase = m.M
												v62 = m.ExcPending
												if v62 != 0 {
													return int32(0)
												} else {
													v63 = *(*int32)(unsafe.Add(mBase, uint32(v38)+52))
													v66 = F_heap_getattr_2(m, v57, v61, v63, v10+int32(31))
													mBase = m.M
													v67 = m.ExcPending
													if v67 != 0 {
														return int32(0)
													} else {
														F_systable_endscan(m, v55)
														mBase = m.M
														v69 = m.ExcPending
														if v69 != 0 {
															return int32(0)
														} else {
															F_sequence_close(m, v38, int32(1))
															mBase = m.M
															v72 = m.ExcPending
															if v72 != 0 {
																return int32(0)
															} else {
																v74 = v66
																v77 = F_has_privs_of_role(m, l2, v74)
																mBase = m.M
																v78 = m.ExcPending
																if v78 != 0 {
																	return int32(0)
																} else {
																	v80 = v77
																	m.G0 = v10 + int32(80)
																	return v80
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
						}
					}
				}
			}
		} else {
			v80 = int32(1)
			m.G0 = v10 + int32(80)
			return v80
		}
	}
}
func F_parse_object(m *base.Module, l0 int32, l1 int32) int32 {
	mBase := m.M
	_ = mBase
	var v5 int32
	_ = v5
	var v6 int32
	_ = v6
	var v10 int32
	_ = v10
	var v11 int32
	_ = v11
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
	var v21 int32
	_ = v21
	var v24 int32
	_ = v24
	var v25 int32
	_ = v25
	var v30 int32
	_ = v30
	var v35 int32
	_ = v35
	var v38 int32
	_ = v38
	var v39 int32
	_ = v39
	var v42 int32
	_ = v42
	var v44 int32
	_ = v44
	var v45 int32
	_ = v45
	var v46 int32
	_ = v46
	var v47 int32
	_ = v47
	var v50 int32
	_ = v50
	var v53 int32
	_ = v53
	var v54 int32
	_ = v54
	var v57 int32
	_ = v57
	var v63 int32
	_ = v63
	var v64 int32
	_ = v64
	var v65 int32
	_ = v65
	var v69 int32
	_ = v69
	var v70 int32
	_ = v70
	var v71 int32
	_ = v71
	var v76 int32
	_ = v76
	v5 = *(*int32)(unsafe.Add(mBase, uint32(l1)+8))
	v6 = *(*int32)(unsafe.Add(mBase, uint32(l1)+4))
	F_check_stack_depth(m)
	mBase = m.M
	v10 = m.ExcPending
	if v10 != 0 {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	return int32(0)
L2:
	;
	if v6 != 0 {
		goto L4
	} else {
		goto L5
	}
L3:
	;
	return v76
L4:
	;
	v11 = *(*int32)(unsafe.Add(mBase, uint32(l1)))
	v12 = m.T0[v6].(func(*base.Module, int32) int32)(m, v11)
	mBase = m.M
	v13 = m.ExcPending
	if v13 != 0 {
		goto L1
	} else {
		goto L7
	}
L5:
	;
	goto L6
L6:
	;
	v15 = *(*int32)(unsafe.Add(mBase, uint32(l0)+32))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+32)) = v15 + int32(1)
	v19 = F_json_lex(m, l0)
	mBase = m.M
	v20 = m.ExcPending
	if v20 != 0 {
		goto L1
	} else {
		goto L9
	}
L7:
	;
	if v12 != 0 {
		v76 = v12
		goto L3
	} else {
		goto L8
	}
L8:
	;
	goto L6
L9:
	;
	if v19 != 0 {
		v76 = v19
		goto L3
	} else {
		goto L10
	}
L10:
	;
	v21 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
	switch v21 - int32(1) {
	case 0:
		goto L13
	default:
		goto L12
	case 3:
		goto L11
	}
L11:
	;
	v63 = F_json_lex(m, l0)
	mBase = m.M
	v64 = m.ExcPending
	if v64 != 0 {
		goto L1
	} else {
		goto L38
	}
L12:
	;
	v50 = int32(11)
	v53 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	if v53 != 0 {
		goto L32
	} else {
		goto L33
	}
L13:
	;
	v24 = F_parse_object_field(m, l0, l1)
	mBase = m.M
	v25 = m.ExcPending
	if v25 != 0 {
		goto L1
	} else {
		goto L14
	}
L14:
	;
	if v24 != 0 {
		v76 = v24
		goto L3
	} else {
		goto L15
	}
L15:
	;
	goto L16
L16:
	;
	v30 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
	if v30 != int32(7) {
		goto L18
	} else {
		goto L19
	}
L17:
	;
	v76 = v46
	goto L3
L18:
	;
	if v30 == int32(4) {
		goto L11
	} else {
		goto L21
	}
L19:
	;
	goto L20
L20:
	;
	v44 = F_json_lex(m, l0)
	mBase = m.M
	v45 = m.ExcPending
	if v45 != 0 {
		goto L1
	} else {
		goto L28
	}
L21:
	;
	v35 = int32(11)
	v38 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	if v38 != 0 {
		goto L22
	} else {
		goto L23
	}
L22:
	;
	v39 = int32(13)
	goto L24
L23:
	;
	v39 = v35
	goto L24
L24:
	;
	if v30 == int32(12) {
		goto L25
	} else {
		goto L26
	}
L25:
	;
	v42 = v35
	goto L27
L26:
	;
	v42 = v39
	goto L27
L27:
	;
	return v42
L28:
	;
	if v44 != 0 {
		v76 = v44
		goto L3
	} else {
		goto L29
	}
L29:
	;
	v46 = F_parse_object_field(m, l0, l1)
	mBase = m.M
	v47 = m.ExcPending
	if v47 != 0 {
		goto L1
	} else {
		goto L30
	}
L30:
	;
	if v46 == int32(0) {
		goto L16
	} else {
		goto L31
	}
L31:
	;
	goto L17
L32:
	;
	v54 = int32(12)
	goto L34
L33:
	;
	v54 = v50
	goto L34
L34:
	;
	if v21 == int32(12) {
		goto L35
	} else {
		goto L36
	}
L35:
	;
	v57 = v50
	goto L37
L36:
	;
	v57 = v54
	goto L37
L37:
	;
	return v57
L38:
	;
	if v63 != 0 {
		v76 = v63
		goto L3
	} else {
		goto L39
	}
L39:
	;
	v65 = *(*int32)(unsafe.Add(mBase, uint32(l0)+32))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+32)) = v65 - int32(1)
	if v5 != 0 {
		goto L40
	} else {
		goto L41
	}
L40:
	;
	v69 = *(*int32)(unsafe.Add(mBase, uint32(l1)))
	v70 = m.T0[v5].(func(*base.Module, int32) int32)(m, v69)
	mBase = m.M
	v71 = m.ExcPending
	if v71 != 0 {
		goto L1
	} else {
		goto L43
	}
L41:
	;
	goto L42
L42:
	;
	v76 = int32(0)
	goto L3
L43:
	;
	if v70 != 0 {
		v76 = v70
		goto L3
	} else {
		goto L44
	}
L44:
	;
	goto L42
}
