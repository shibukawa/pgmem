package p0

import (
	base "github.com/shibukawa/pgmem/internal/aot/pgaot/base"
	"unsafe"
)

func F_FlushBuffer(m *base.Module, l0 int32, l1 int32, l2 int32) {
	mBase := m.M
	_ = mBase
	var v4 int32
	_ = v4
	var v10 int32
	_ = v10
	var v12 int32
	_ = v12
	var v16 int32
	_ = v16
	var v17 int32
	_ = v17
	var v21 int32
	_ = v21
	var v22 int32
	_ = v22
	var v30 int32
	_ = v30
	var v32 int32
	_ = v32
	var v34 int32
	_ = v34
	var v36 int64
	_ = v36
	var v40 int32
	_ = v40
	var v41 int32
	_ = v41
	var v42 int32
	_ = v42
	var v53 int32
	_ = v53
	var v54 int32
	_ = v54
	var v71 int32
	_ = v71
	var v72 int32
	_ = v72
	var v73 int32
	_ = v73
	var v82 int32
	_ = v82
	var v90 int32
	_ = v90
	var v91 int32
	_ = v91
	var v93 int32
	_ = v93
	var v98 int32
	_ = v98
	var v101 int32
	_ = v101
	var v108 int32
	_ = v108
	var v111 int32
	_ = v111
	var v112 int32
	_ = v112
	var v115 int32
	_ = v115
	var v116 int64
	_ = v116
	var v117 int64
	_ = v117
	var v119 int32
	_ = v119
	var v127 int32
	_ = v127
	var v128 int32
	_ = v128
	var v130 int32
	_ = v130
	var v131 int32
	_ = v131
	var v134 int32
	_ = v134
	var v135 int32
	_ = v135
	var v139 int32
	_ = v139
	var v140 int32
	_ = v140
	var v146 int32
	_ = v146
	var v151 int32
	_ = v151
	var v155 int32
	_ = v155
	var v156 int32
	_ = v156
	var v158 int32
	_ = v158
	var v161 int32
	_ = v161
	var v194 int32
	_ = v194
	var v197 int32
	_ = v197
	var v198 int32
	_ = v198
	var v199 int32
	_ = v199
	var v206 int32
	_ = v206
	var v240 int32
	_ = v240
	var v247 int32
	_ = v247
	var v277 int32
	_ = v277
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
	var v291 int32
	_ = v291
	var v292 int32
	_ = v292
	var v294 int32
	_ = v294
	var v295 int32
	_ = v295
	var v296 int32
	_ = v296
	var v304 int32
	_ = v304
	var v308 int32
	_ = v308
	var v317 int32
	_ = v317
	var v349 int32
	_ = v349
	var v350 int32
	_ = v350
	var v351 int32
	_ = v351
	var v353 int32
	_ = v353
	var v357 int32
	_ = v357
	var v364 int32
	_ = v364
	var v371 int32
	_ = v371
	var v379 int32
	_ = v379
	var v388 int32
	_ = v388
	var v420 int32
	_ = v420
	var v421 int32
	_ = v421
	var v422 int32
	_ = v422
	var v424 int32
	_ = v424
	var v428 int32
	_ = v428
	var v435 int32
	_ = v435
	var v442 int32
	_ = v442
	var v450 int32
	_ = v450
	var v453 int32
	_ = v453
	var v454 int32
	_ = v454
	var v455 int32
	_ = v455
	var v456 int32
	_ = v456
	var v457 int32
	_ = v457
	var v458 int32
	_ = v458
	var v459 int32
	_ = v459
	var v460 int32
	_ = v460
	var v461 int32
	_ = v461
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
	var v468 int32
	_ = v468
	var v469 int32
	_ = v469
	var v470 int32
	_ = v470
	var v471 int32
	_ = v471
	var v472 int32
	_ = v472
	var v473 int32
	_ = v473
	var v474 int32
	_ = v474
	var v475 int32
	_ = v475
	var v476 int32
	_ = v476
	var v477 int32
	_ = v477
	var v478 int32
	_ = v478
	var v479 int32
	_ = v479
	var v480 int32
	_ = v480
	var v481 int32
	_ = v481
	var v482 int32
	_ = v482
	var v483 int32
	_ = v483
	var v484 int32
	_ = v484
	var v518 int32
	_ = v518
	var v519 int32
	_ = v519
	var v523 int32
	_ = v523
	var v525 int32
	_ = v525
	var v528 int32
	_ = v528
	var v530 int32
	_ = v530
	var v533 int32
	_ = v533
	var v535 int32
	_ = v535
	var v539 int64
	_ = v539
	var v540 int64
	_ = v540
	var v544 int64
	_ = v544
	var v548 int32
	_ = v548
	var v549 int32
	_ = v549
	var v555 int32
	_ = v555
	var v556 int32
	_ = v556
	var v558 int32
	_ = v558
	var v559 int64
	_ = v559
	var v563 int32
	_ = v563
	var v565 int32
	_ = v565
	var v571 int64
	_ = v571
	var v572 int64
	_ = v572
	var v576 int64
	_ = v576
	var v583 int32
	_ = v583
	var v585 int64
	_ = v585
	var v587 int64
	_ = v587
	var v590 int32
	_ = v590
	var v592 int64
	_ = v592
	var v595 int32
	_ = v595
	var v597 int64
	_ = v597
	var v620 int32
	_ = v620
	var v622 int32
	_ = v622
	var v627 int64
	_ = v627
	var v631 int32
	_ = v631
	var v643 int64
	_ = v643
	var v647 int32
	_ = v647
	var v659 int32
	_ = v659
	var v664 int64
	_ = v664
	var v668 int64
	_ = v668
	var v673 int32
	_ = v673
	var v681 int32
	_ = v681
	var v683 int64
	_ = v683
	var v687 int32
	_ = v687
	var v688 int32
	_ = v688
	var v692 int32
	_ = v692
	var v694 int32
	_ = v694
	v4 = int32(0)
	v10 = m.G0
	v12 = v10 + int32(-64)
	m.G0 = v12
	v16 = F_StartBufferIO(m, l0, v4, v4)
	mBase = m.M
	v17 = m.ExcPending
	if v17 != 0 {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	return
L2:
	;
	if v16 != 0 {
		goto L3
	} else {
		goto L4
	}
L3:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v12)+36)) = l0
	*(*int32)(unsafe.Add(mBase, uint32(v12)+32)) = int32(1075)
	v21 = int32(_a_F_FlushBuffer_0)
	v22 = *(*int32)(unsafe.Add(mBase, _c_F_FlushBuffer[0]))
	*(*int32)(unsafe.Add(mBase, _c_F_FlushBuffer[0])) = v10 + int32(-36)
	*(*int32)(unsafe.Add(mBase, uint32(v12)+28)) = v22
	if l1 == int32(0) {
		goto L6
	} else {
		goto L7
	}
L4:
	;
	goto L5
L5:
	;
	m.G0 = v12 - int32(-64)
	return
L6:
	;
	v30 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	*(*int32)(unsafe.Add(mBase, uint32(v12)+16)) = v30
	v32 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	*(*int32)(unsafe.Add(mBase, uint32(v12)+20)) = v32
	v34 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	*(*int32)(unsafe.Add(mBase, uint32(v12)+24)) = v34
	v36 = *(*int64)(unsafe.Add(mBase, uint32(v12)+16))
	*(*int64)(unsafe.Add(mBase, uint32(v12))) = v36
	*(*int32)(unsafe.Add(mBase, uint32(v12)+8)) = v34
	v40 = F_smgropen(m, v12, int32(-1))
	mBase = m.M
	v41 = m.ExcPending
	if v41 != 0 {
		goto L1
	} else {
		goto L9
	}
L7:
	;
	v42 = l1
	goto L8
L8:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v12)+60)) = int32(_a_F_FlushBuffer_1)
	*(*int32)(unsafe.Add(mBase, uint32(v12)+56)) = int32(_a_F_FlushBuffer_2)
	*(*int32)(unsafe.Add(mBase, uint32(v12)+52)) = int32(_a_F_FlushBuffer_3)
	*(*int32)(unsafe.Add(mBase, uint32(v12)+48)) = int32(0)
	*(*int64)(unsafe.Add(mBase, uint32(v12)+40)) = int64(0)
	v53 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
	v54 = int32(_a_F_FlushBuffer_4)
	*(*int32)(unsafe.Add(mBase, uint32(l0)+24)) = v53 | v54
	if v53&v54 != 0 {
		goto L10
	} else {
		goto L11
	}
L9:
	;
	v42 = v40
	goto L8
L10:
	;
	goto L13
L11:
	;
	v82 = v53
	goto L12
L12:
	;
	v90 = int32(_a_F_FlushBuffer_5)
	v91 = *(*int32)(unsafe.Add(mBase, _c_F_FlushBuffer[1]))
	v93 = *(*int32)(unsafe.Add(mBase, uint32(v10+int32(-24))+8))
	if v93 == int32(0) {
		goto L20
	} else {
		goto L21
	}
L13:
	;
	F_perform_spin_delay(m, v10+int32(-24))
	mBase = m.M
	v71 = m.ExcPending
	if v71 != 0 {
		goto L1
	} else {
		goto L15
	}
L14:
	;
	v82 = v72
	goto L12
L15:
	;
	v72 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
	v73 = int32(_a_F_FlushBuffer_4)
	*(*int32)(unsafe.Add(mBase, uint32(l0)+24)) = v72 | v73
	if v72&v73 != 0 {
		goto L13
	} else {
		goto L16
	}
L16:
	;
	goto L14
L17:
	;
	v111 = *(*int32)(unsafe.Add(mBase, _c_F_FlushBuffer[2]))
	v112 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v115 = v111 + v112<<(uint(int32(13))%32)
	v116 = int64(*(*uint32)(unsafe.Add(mBase, uint32(v115)+4)))
	v117 = int64(*(*uint32)(unsafe.Add(mBase, uint32(v115))))
	v119 = v82 & int32(-272629761)
	*(*int32)(unsafe.Add(mBase, uint32(l0)+24)) = v119
	if base.Ui32(int32(2143289344)) <= base.Ui32(v119) {
		goto L28
	} else {
		goto L29
	}
L18:
	;
	goto L17
L19:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_FlushBuffer[1])) = v108
	goto L18
L20:
	;
	if int32(999) < v91 {
		goto L18
	} else {
		goto L23
	}
L21:
	;
	goto L22
L22:
	;
	if v91 < int32(11) {
		goto L18
	} else {
		goto L27
	}
L23:
	;
	v98 = int32(900)
	if v98 <= v91 {
		goto L24
	} else {
		goto L25
	}
L24:
	;
	v101 = v98
	goto L26
L25:
	;
	v101 = v91
	goto L26
L26:
	;
	v108 = v101 + int32(100)
	goto L19
L27:
	;
	v108 = v91 - int32(1)
	goto L19
L28:
	;
	F_XLogFlush(m, v117<<(uint(int64(32))%64)|v116)
	mBase = m.M
	v127 = m.ExcPending
	if v127 != 0 {
		goto L1
	} else {
		goto L31
	}
L29:
	;
	goto L30
L30:
	;
	v128 = *(*int32)(unsafe.Add(mBase, uint32(l0)+16))
	v130 = *(*int32)(unsafe.Add(mBase, _c_F_FlushBuffer[2]))
	v131 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v134 = v130 + v131<<(uint(int32(13))%32)
	v135 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v134)+14)))
	if v135 == int32(0) {
		v528 = v134
		goto L32
	} else {
		goto L33
	}
L31:
	;
	goto L30
L32:
	;
	v530 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_FlushBuffer[3])))
	v533 = m.G0
	v535 = v533 - int32(16)
	m.G0 = v535
	if v530 != 0 {
		goto L54
	} else {
		goto L55
	}
L33:
	;
	v139 = *(*int32)(unsafe.Add(mBase, _c_F_FlushBuffer[4]))
	v140 = *(*int32)(unsafe.Add(mBase, uint32(v139)+252))
	goto L34
L34:
	;
	if base.B2i32(v140 != int32(0)) == int32(0) {
		v528 = v134
		goto L32
	} else {
		goto L35
	}
L35:
	;
	v146 = *(*int32)(unsafe.Add(mBase, _c_F_FlushBuffer[5]))
	if v146 == int32(0) {
		goto L36
	} else {
		goto L37
	}
L36:
	;
	v151 = *(*int32)(unsafe.Add(mBase, _c_F_FlushBuffer[6]))
	v155 = F_MemoryContextAllocAligned(m, v151, int32(_a_F_FlushBuffer_6), int32(_a_F_FlushBuffer_7), int32(0))
	mBase = m.M
	v156 = m.ExcPending
	if v156 != 0 {
		goto L1
	} else {
		goto L39
	}
L37:
	;
	v158 = v146
	goto L38
L38:
	;
	base.MemoryCopy(m, v158, v134, int32(_a_F_FlushBuffer_6))
	v161 = int32(0)
	v194 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v158)+8)))
	*(*uint16)(unsafe.Add(mBase, uint32(v158)+8)) = uint16(v161)
	v197 = m.G0
	v198 = int32(128)
	v199 = v197 - v198
	base.MemoryCopy(m, v199, int32(_a_F_FlushBuffer_8), v198)
	v206 = v161
	goto L41
L39:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_FlushBuffer[5])) = v155
	v158 = v155
	goto L38
L40:
	;
	v525 = *(*int32)(unsafe.Add(mBase, _c_F_FlushBuffer[5]))
	*(*uint16)(unsafe.Add(mBase, uint32(v525)+8)) = uint16(v523)
	v528 = v525
	goto L32
L41:
	;
	v240 = v158 + v206<<(uint(int32(7))%32)
	v247 = int32(0)
	goto L43
L42:
	;
	v317 = int32(0)
	goto L47
L43:
	;
	v277 = int32(2)
	v278 = v247 << (uint(v277) % 32)
	v279 = v199 + v278
	v281 = *(*int32)(unsafe.Add(mBase, uint32(v240+v278)))
	v282 = *(*int32)(unsafe.Add(mBase, uint32(v279)))
	v283 = v281 ^ v282
	v284 = int32(16777619)
	v286 = int32(17)
	*(*int32)(unsafe.Add(mBase, uint32(v279))) = v283*v284 ^ int32(base.Ui32(v283)>>(uint(v286)%32))
	v291 = v278 | int32(4)
	v292 = v199 + v291
	v294 = *(*int32)(unsafe.Add(mBase, uint32(v240+v291)))
	v295 = *(*int32)(unsafe.Add(mBase, uint32(v292)))
	v296 = v294 ^ v295
	*(*int32)(unsafe.Add(mBase, uint32(v292))) = v296*v284 ^ int32(base.Ui32(v296)>>(uint(v286)%32))
	v304 = v247 + v277
	if v304 != int32(32) {
		v247 = v304
		goto L43
	} else {
		goto L45
	}
L44:
	;
	v308 = v206 + int32(1)
	if v308 != int32(64) {
		v206 = v308
		goto L41
	} else {
		goto L46
	}
L45:
	;
	goto L44
L46:
	;
	goto L42
L47:
	;
	v349 = v199 + v317<<(uint(int32(2))%32)
	v350 = *(*int32)(unsafe.Add(mBase, uint32(v349)))
	v351 = int32(16777619)
	v353 = int32(17)
	*(*int32)(unsafe.Add(mBase, uint32(v349))) = v350*v351 ^ int32(base.Ui32(v350)>>(uint(v353)%32))
	v357 = *(*int32)(unsafe.Add(mBase, uint32(v349)+4))
	*(*int32)(unsafe.Add(mBase, uint32(v349)+4)) = v357*v351 ^ int32(base.Ui32(v357)>>(uint(v353)%32))
	v364 = *(*int32)(unsafe.Add(mBase, uint32(v349)+8))
	*(*int32)(unsafe.Add(mBase, uint32(v349)+8)) = v364*v351 ^ int32(base.Ui32(v364)>>(uint(v353)%32))
	v371 = *(*int32)(unsafe.Add(mBase, uint32(v349)+12))
	*(*int32)(unsafe.Add(mBase, uint32(v349)+12)) = v371*v351 ^ int32(base.Ui32(v371)>>(uint(v353)%32))
	v379 = v317 + int32(4)
	if v379 != int32(32) {
		v317 = v379
		goto L47
	} else {
		goto L49
	}
L48:
	;
	v388 = int32(0)
	goto L50
L49:
	;
	goto L48
L50:
	;
	v420 = v199 + v388<<(uint(int32(2))%32)
	v421 = *(*int32)(unsafe.Add(mBase, uint32(v420)))
	v422 = int32(16777619)
	v424 = int32(17)
	*(*int32)(unsafe.Add(mBase, uint32(v420))) = v421*v422 ^ int32(base.Ui32(v421)>>(uint(v424)%32))
	v428 = *(*int32)(unsafe.Add(mBase, uint32(v420)+4))
	*(*int32)(unsafe.Add(mBase, uint32(v420)+4)) = v428*v422 ^ int32(base.Ui32(v428)>>(uint(v424)%32))
	v435 = *(*int32)(unsafe.Add(mBase, uint32(v420)+8))
	*(*int32)(unsafe.Add(mBase, uint32(v420)+8)) = v435*v422 ^ int32(base.Ui32(v435)>>(uint(v424)%32))
	v442 = *(*int32)(unsafe.Add(mBase, uint32(v420)+12))
	*(*int32)(unsafe.Add(mBase, uint32(v420)+12)) = v442*v422 ^ int32(base.Ui32(v442)>>(uint(v424)%32))
	v450 = v388 + int32(4)
	if v450 != int32(32) {
		v388 = v450
		goto L50
	} else {
		goto L52
	}
L51:
	;
	v453 = *(*int32)(unsafe.Add(mBase, uint32(v199)+124))
	v454 = *(*int32)(unsafe.Add(mBase, uint32(v199)+120))
	v455 = *(*int32)(unsafe.Add(mBase, uint32(v199)+116))
	v456 = *(*int32)(unsafe.Add(mBase, uint32(v199)+112))
	v457 = *(*int32)(unsafe.Add(mBase, uint32(v199)+108))
	v458 = *(*int32)(unsafe.Add(mBase, uint32(v199)+104))
	v459 = *(*int32)(unsafe.Add(mBase, uint32(v199)+100))
	v460 = *(*int32)(unsafe.Add(mBase, uint32(v199)+96))
	v461 = *(*int32)(unsafe.Add(mBase, uint32(v199)+92))
	v462 = *(*int32)(unsafe.Add(mBase, uint32(v199)+88))
	v463 = *(*int32)(unsafe.Add(mBase, uint32(v199)+84))
	v464 = *(*int32)(unsafe.Add(mBase, uint32(v199)+80))
	v465 = *(*int32)(unsafe.Add(mBase, uint32(v199)+76))
	v466 = *(*int32)(unsafe.Add(mBase, uint32(v199)+72))
	v467 = *(*int32)(unsafe.Add(mBase, uint32(v199)+68))
	v468 = *(*int32)(unsafe.Add(mBase, uint32(v199)+64))
	v469 = *(*int32)(unsafe.Add(mBase, uint32(v199)+60))
	v470 = *(*int32)(unsafe.Add(mBase, uint32(v199)+56))
	v471 = *(*int32)(unsafe.Add(mBase, uint32(v199)+52))
	v472 = *(*int32)(unsafe.Add(mBase, uint32(v199)+48))
	v473 = *(*int32)(unsafe.Add(mBase, uint32(v199)+44))
	v474 = *(*int32)(unsafe.Add(mBase, uint32(v199)+40))
	v475 = *(*int32)(unsafe.Add(mBase, uint32(v199)+36))
	v476 = *(*int32)(unsafe.Add(mBase, uint32(v199)+32))
	v477 = *(*int32)(unsafe.Add(mBase, uint32(v199)+28))
	v478 = *(*int32)(unsafe.Add(mBase, uint32(v199)+24))
	v479 = *(*int32)(unsafe.Add(mBase, uint32(v199)+20))
	v480 = *(*int32)(unsafe.Add(mBase, uint32(v199)+16))
	v481 = *(*int32)(unsafe.Add(mBase, uint32(v199)+12))
	v482 = *(*int32)(unsafe.Add(mBase, uint32(v199)+8))
	v483 = *(*int32)(unsafe.Add(mBase, uint32(v199)+4))
	v484 = *(*int32)(unsafe.Add(mBase, uint32(v199)))
	*(*uint16)(unsafe.Add(mBase, uint32(v158)+8)) = uint16(v194)
	v518 = int32(_a_F_FlushBuffer_9)
	v519 = base.I32_rem_u_s(v453^(v454^(v455^(v456^(v457^(v458^(v459^(v460^(v461^(v462^(v463^(v464^(v465^(v466^(v467^(v468^(v469^(v470^(v471^(v472^(v473^(v474^(v475^(v476^(v477^(v478^(v479^(v480^(v481^(v482^(v483^(v128^v484))))))))))))))))))))))))))))))), v518)
	v523 = (v519 + int32(1)) & v518
	goto L40
L52:
	;
	goto L51
L53:
	;
	v548 = *(*int32)(unsafe.Add(mBase, uint32(l0)+16))
	v549 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	*(*int32)(unsafe.Add(mBase, uint32(v12)+40)) = v528
	F_smgrwritev(m, v42, v549, v548, v10+int32(-24), int32(0))
	mBase = m.M
	v555 = m.ExcPending
	if v555 != 0 {
		goto L1
	} else {
		goto L57
	}
L54:
	;
	F___clock_gettime(m, int32(1), v535)
	mBase = m.M
	v539 = int64(*(*int32)(unsafe.Add(mBase, uint32(v535)+8)))
	v540 = *(*int64)(unsafe.Add(mBase, uint32(v535)))
	v544 = v539 + v540*int64(1000000000)
	goto L56
L55:
	;
	v544 = int64(0)
	goto L56
L56:
	;
	m.G0 = v535 + int32(16)
	goto L53
L57:
	;
	v556 = int32(0)
	v558 = int32(1)
	v559 = int64(8192)
	v563 = m.G0
	v565 = v563 - int32(16)
	m.G0 = v565
	if v544 != int64(0) {
		goto L59
	} else {
		goto L60
	}
L58:
	;
	v681 = int32(_a_F_FlushBuffer_10)
	v683 = *(*int64)(unsafe.Add(mBase, _c_F_FlushBuffer[7]))
	*(*int64)(unsafe.Add(mBase, _c_F_FlushBuffer[7])) = v683 + int64(1)
	v687 = int32(1)
	v688 = int32(0)
	F_TerminateBufferIO(m, l0, v687, v688, v687, v688)
	mBase = m.M
	v692 = m.ExcPending
	if v692 != 0 {
		goto L1
	} else {
		goto L75
	}
L59:
	;
	F___clock_gettime(m, int32(1), v565)
	mBase = m.M
	v571 = int64(*(*int32)(unsafe.Add(mBase, uint32(v565)+8)))
	v572 = *(*int64)(unsafe.Add(mBase, uint32(v565)))
	v576 = v571 + (v572*int64(1000000000) - v544)
	goto L63
L60:
	;
	goto L61
L61:
	;
	v659 = l2 << (uint(int32(6)) % 32)
	v664 = *(*int64)(unsafe.Add(mBase, uint32(v659)+uint32(_c_F_FlushBuffer[8])))
	*(*int64)(unsafe.Add(mBase, uint32(v659)+uint32(_c_F_FlushBuffer[8]))) = v664 + base.I64_extend_i32_u(v558)
	v668 = *(*int64)(unsafe.Add(mBase, uint32(v659)+uint32(_c_F_FlushBuffer[9])))
	*(*int64)(unsafe.Add(mBase, uint32(v659)+uint32(_c_F_FlushBuffer[9]))) = v668 + v559
	F_pgstat_count_backend_io_op(m, v556, l2, int32(7), v558, v559)
	mBase = m.M
	v673 = int32(1)
	*(*uint8)(unsafe.Add(mBase, _c_F_FlushBuffer[10])) = uint8(v673)
	*(*uint8)(unsafe.Add(mBase, _c_F_FlushBuffer[11])) = uint8(v673)
	m.G0 = v565 + int32(16)
	goto L58
L62:
	;
	v620 = int32(0)
	v622 = l2 << (uint(int32(6)) % 32)
	v627 = *(*int64)(unsafe.Add(mBase, uint32(v622)+uint32(_c_F_FlushBuffer[12])))
	*(*int64)(unsafe.Add(mBase, uint32(v622)+uint32(_c_F_FlushBuffer[12]))) = v627 + v576
	v631 = *(*int32)(unsafe.Add(mBase, _c_F_FlushBuffer[13]))
	if base.B2i32(base.Ui32(int32(16)) < base.Ui32(v631))|base.B2i32(int32(1)<<(uint(v631)%32)&int32(_a_F_FlushBuffer_11) == v620) == v620 {
		goto L72
	} else {
		goto L73
	}
L63:
	;
	goto L64
L64:
	;
	v583 = int32(_a_F_FlushBuffer_12)
	v585 = *(*int64)(unsafe.Add(mBase, _c_F_FlushBuffer[14]))
	v587 = base.I64_div_s(v576, int64(1000))
	*(*int64)(unsafe.Add(mBase, _c_F_FlushBuffer[14])) = v585 + v587
	switch v556 {
	case 0:
		goto L68
	case 1:
		goto L67
	default:
		goto L62
	}
L67:
	;
	v595 = int32(_a_F_FlushBuffer_13)
	v597 = *(*int64)(unsafe.Add(mBase, _c_F_FlushBuffer[15]))
	*(*int64)(unsafe.Add(mBase, _c_F_FlushBuffer[15])) = v597 + v576
	goto L62
L68:
	;
	v590 = int32(_a_F_FlushBuffer_14)
	v592 = *(*int64)(unsafe.Add(mBase, _c_F_FlushBuffer[16]))
	*(*int64)(unsafe.Add(mBase, _c_F_FlushBuffer[16])) = v592 + v576
	goto L62
L72:
	;
	v643 = *(*int64)(unsafe.Add(mBase, uint32(v622)+uint32(_c_F_FlushBuffer[17])))
	*(*int64)(unsafe.Add(mBase, uint32(v622)+uint32(_c_F_FlushBuffer[17]))) = v643 + v576
	v647 = int32(1)
	*(*uint8)(unsafe.Add(mBase, _c_F_FlushBuffer[10])) = uint8(v647)
	*(*uint8)(unsafe.Add(mBase, _c_F_FlushBuffer[18])) = uint8(v647)
	goto L74
L73:
	;
	goto L74
L74:
	;
	goto L61
L75:
	;
	v694 = *(*int32)(unsafe.Add(mBase, uint32(v12)+28))
	*(*int32)(unsafe.Add(mBase, _c_F_FlushBuffer[0])) = v694
	goto L5
}
func F_UnpinBufferNoOwner(m *base.Module, l0 int32) {
	mBase := m.M
	_ = mBase
	var v7 int32
	_ = v7
	var v9 int32
	_ = v9
	var v11 int32
	_ = v11
	var v13 int32
	_ = v13
	var v15 int32
	_ = v15
	var v17 int32
	_ = v17
	var v19 int32
	_ = v19
	var v21 int32
	_ = v21
	var v23 int32
	_ = v23
	var v25 int32
	_ = v25
	var v27 int32
	_ = v27
	var v29 int32
	_ = v29
	var v31 int32
	_ = v31
	var v33 int32
	_ = v33
	var v35 int32
	_ = v35
	var v37 int32
	_ = v37
	var v39 int32
	_ = v39
	var v41 int32
	_ = v41
	var v43 int32
	_ = v43
	var v45 int32
	_ = v45
	var v47 int32
	_ = v47
	var v49 int32
	_ = v49
	var v53 int32
	_ = v53
	var v56 int32
	_ = v56
	var v58 int32
	_ = v58
	var v59 int32
	_ = v59
	var v60 int32
	_ = v60
	var v61 int32
	_ = v61
	var v63 int32
	_ = v63
	var v65 int32
	_ = v65
	var v67 int32
	_ = v67
	var v84 int32
	_ = v84
	var v96 int32
	_ = v96
	var v97 int32
	_ = v97
	var v101 int32
	_ = v101
	var v109 int32
	_ = v109
	var v110 int32
	_ = v110
	var v112 int32
	_ = v112
	var v117 int32
	_ = v117
	var v120 int32
	_ = v120
	var v127 int32
	_ = v127
	var v130 int32
	_ = v130
	var v136 int32
	_ = v136
	var v137 int32
	_ = v137
	var v138 int32
	_ = v138
	var v139 int32
	_ = v139
	var v157 int32
	_ = v157
	var v158 int32
	_ = v158
	var v172 int32
	_ = v172
	var v173 int32
	_ = v173
	var v174 int32
	_ = v174
	var v180 int32
	_ = v180
	var v188 int32
	_ = v188
	var v189 int32
	_ = v189
	var v191 int32
	_ = v191
	var v196 int32
	_ = v196
	var v199 int32
	_ = v199
	var v206 int32
	_ = v206
	var v212 int32
	_ = v212
	var v217 int32
	_ = v217
	var v238 int32
	_ = v238
	var v241 int32
	_ = v241
	var v247 int32
	_ = v247
	var v248 int32
	_ = v248
	var v249 int32
	_ = v249
	var v251 int32
	_ = v251
	v7 = m.G0
	v9 = v7 - int32(32)
	m.G0 = v9
	v11 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v13 = v11 + int32(1)
	*(*int32)(unsafe.Add(mBase, uint32(v9)+4)) = v13
	v15 = int32(_a_F_UnpinBufferNoOwner_0)
	v17 = *(*int32)(unsafe.Add(mBase, _c_F_UnpinBufferNoOwner[0]))
	if v13 == v17 {
		v60 = v15
		goto L1
	} else {
		goto L2
	}
L1:
	;
	v61 = *(*int32)(unsafe.Add(mBase, uint32(v60)+4))
	v63 = v61 - int32(1)
	*(*int32)(unsafe.Add(mBase, uint32(v60)+4)) = v63
	if v63 != 0 {
		goto L13
	} else {
		goto L14
	}
L2:
	;
	v19 = int32(_a_F_UnpinBufferNoOwner_1)
	v21 = *(*int32)(unsafe.Add(mBase, _c_F_UnpinBufferNoOwner[1]))
	if v13 == v21 {
		v60 = v19
		goto L1
	} else {
		goto L3
	}
L3:
	;
	v23 = int32(_a_F_UnpinBufferNoOwner_2)
	v25 = *(*int32)(unsafe.Add(mBase, _c_F_UnpinBufferNoOwner[2]))
	if v13 == v25 {
		v60 = v23
		goto L1
	} else {
		goto L4
	}
L4:
	;
	v27 = int32(_a_F_UnpinBufferNoOwner_3)
	v29 = *(*int32)(unsafe.Add(mBase, _c_F_UnpinBufferNoOwner[3]))
	if v13 == v29 {
		v60 = v27
		goto L1
	} else {
		goto L5
	}
L5:
	;
	v31 = int32(_a_F_UnpinBufferNoOwner_4)
	v33 = *(*int32)(unsafe.Add(mBase, _c_F_UnpinBufferNoOwner[4]))
	if v13 == v33 {
		v60 = v31
		goto L1
	} else {
		goto L6
	}
L6:
	;
	v35 = int32(_a_F_UnpinBufferNoOwner_5)
	v37 = *(*int32)(unsafe.Add(mBase, _c_F_UnpinBufferNoOwner[5]))
	if v13 == v37 {
		v60 = v35
		goto L1
	} else {
		goto L7
	}
L7:
	;
	v39 = int32(_a_F_UnpinBufferNoOwner_6)
	v41 = *(*int32)(unsafe.Add(mBase, _c_F_UnpinBufferNoOwner[6]))
	if v13 == v41 {
		v60 = v39
		goto L1
	} else {
		goto L8
	}
L8:
	;
	v43 = int32(_a_F_UnpinBufferNoOwner_7)
	v45 = *(*int32)(unsafe.Add(mBase, _c_F_UnpinBufferNoOwner[7]))
	if v13 == v45 {
		v60 = v43
		goto L1
	} else {
		goto L9
	}
L9:
	;
	v47 = int32(0)
	v49 = *(*int32)(unsafe.Add(mBase, _c_F_UnpinBufferNoOwner[8]))
	if v49 == v47 {
		v60 = v47
		goto L1
	} else {
		goto L10
	}
L10:
	;
	v53 = *(*int32)(unsafe.Add(mBase, _c_F_UnpinBufferNoOwner[9]))
	v56 = int32(0)
	v58 = F_hash_search(m, v53, v9+int32(4), v56, v56)
	mBase = m.M
	v59 = m.ExcPending
	if v59 != 0 {
		goto L11
	} else {
		goto L12
	}
L11:
	;
	return
L12:
	;
	v60 = v58
	goto L1
L13:
	;
	m.G0 = v9 + int32(32)
	return
L14:
	;
	v65 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
	v67 = v65
	goto L15
L15:
	;
	if v67&int32(_a_F_UnpinBufferNoOwner_8) != 0 {
		goto L17
	} else {
		goto L18
	}
L16:
	;
	if v136&int32(536870912) == int32(0) {
		goto L44
	} else {
		goto L45
	}
L17:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v9)+24)) = int32(_a_F_UnpinBufferNoOwner_9)
	*(*int32)(unsafe.Add(mBase, uint32(v9)+20)) = int32(_a_F_UnpinBufferNoOwner_10)
	*(*int32)(unsafe.Add(mBase, uint32(v9)+16)) = int32(_a_F_UnpinBufferNoOwner_11)
	*(*int32)(unsafe.Add(mBase, uint32(v9)+12)) = int32(0)
	*(*int64)(unsafe.Add(mBase, uint32(v9)+4)) = int64(0)
	v84 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
	if v84&int32(_a_F_UnpinBufferNoOwner_8) != 0 {
		goto L20
	} else {
		goto L21
	}
L18:
	;
	v130 = v67
	goto L19
L19:
	;
	v136 = v130 - int32(1)
	v137 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
	v138 = base.B2i32(v130 == v137)
	if v130 == v137 {
		goto L38
	} else {
		goto L39
	}
L20:
	;
	goto L23
L21:
	;
	v101 = v84
	goto L22
L22:
	;
	v109 = int32(_a_F_UnpinBufferNoOwner_12)
	v110 = *(*int32)(unsafe.Add(mBase, _c_F_UnpinBufferNoOwner[10]))
	v112 = *(*int32)(unsafe.Add(mBase, uint32(v9+int32(4))+8))
	if v112 == int32(0) {
		goto L30
	} else {
		goto L31
	}
L23:
	;
	F_perform_spin_delay(m, v9+int32(4))
	mBase = m.M
	v96 = m.ExcPending
	if v96 != 0 {
		goto L11
	} else {
		goto L25
	}
L24:
	;
	v101 = v97
	goto L22
L25:
	;
	v97 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
	if v97&int32(_a_F_UnpinBufferNoOwner_8) != 0 {
		goto L23
	} else {
		goto L26
	}
L26:
	;
	goto L24
L27:
	;
	v130 = v101
	goto L19
L28:
	;
	goto L27
L29:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_UnpinBufferNoOwner[10])) = v127
	goto L28
L30:
	;
	if int32(999) < v110 {
		goto L28
	} else {
		goto L33
	}
L31:
	;
	goto L32
L32:
	;
	if v110 < int32(11) {
		goto L28
	} else {
		goto L37
	}
L33:
	;
	v117 = int32(900)
	if v117 <= v110 {
		goto L34
	} else {
		goto L35
	}
L34:
	;
	v120 = v117
	goto L36
L35:
	;
	v120 = v110
	goto L36
L36:
	;
	v127 = v120 + int32(100)
	goto L29
L37:
	;
	v127 = v110 - int32(1)
	goto L29
L38:
	;
	v139 = v136
	goto L40
L39:
	;
	v139 = v137
	goto L40
L40:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+24)) = v139
	if v138 == int32(0) {
		goto L41
	} else {
		goto L42
	}
L41:
	;
	v67 = v137
	goto L15
L42:
	;
	goto L43
L43:
	;
	goto L16
L44:
	;
	if base.B2i32(base.Ui32(v60) < base.Ui32(int32(_a_F_UnpinBufferNoOwner_0)))|base.B2i32(base.Ui32(int32(_a_F_UnpinBufferNoOwner_13)) <= base.Ui32(v60)) == int32(0) {
		goto L68
	} else {
		goto L69
	}
L45:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v9)+24)) = int32(_a_F_UnpinBufferNoOwner_14)
	*(*int32)(unsafe.Add(mBase, uint32(v9)+20)) = int32(_a_F_UnpinBufferNoOwner_15)
	*(*int32)(unsafe.Add(mBase, uint32(v9)+16)) = int32(_a_F_UnpinBufferNoOwner_11)
	*(*int32)(unsafe.Add(mBase, uint32(v9)+12)) = int32(0)
	*(*int64)(unsafe.Add(mBase, uint32(v9)+4)) = int64(0)
	v157 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
	v158 = int32(_a_F_UnpinBufferNoOwner_8)
	*(*int32)(unsafe.Add(mBase, uint32(l0)+24)) = v157 | v158
	if v157&v158 != 0 {
		goto L46
	} else {
		goto L47
	}
L46:
	;
	goto L49
L47:
	;
	v180 = v157
	goto L48
L48:
	;
	v188 = int32(_a_F_UnpinBufferNoOwner_12)
	v189 = *(*int32)(unsafe.Add(mBase, _c_F_UnpinBufferNoOwner[10]))
	v191 = *(*int32)(unsafe.Add(mBase, uint32(v9+int32(4))+8))
	if v191 == int32(0) {
		goto L56
	} else {
		goto L57
	}
L49:
	;
	F_perform_spin_delay(m, v9+int32(4))
	mBase = m.M
	v172 = m.ExcPending
	if v172 != 0 {
		goto L11
	} else {
		goto L51
	}
L50:
	;
	v180 = v173
	goto L48
L51:
	;
	v173 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
	v174 = int32(_a_F_UnpinBufferNoOwner_8)
	*(*int32)(unsafe.Add(mBase, uint32(l0)+24)) = v173 | v174
	if v173&v174 != 0 {
		goto L49
	} else {
		goto L52
	}
L52:
	;
	goto L50
L53:
	;
	if v180&int32(537133055) == int32(536870913) {
		goto L64
	} else {
		goto L65
	}
L54:
	;
	goto L53
L55:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_UnpinBufferNoOwner[10])) = v206
	goto L54
L56:
	;
	if int32(999) < v189 {
		goto L54
	} else {
		goto L59
	}
L57:
	;
	goto L58
L58:
	;
	if v189 < int32(11) {
		goto L54
	} else {
		goto L63
	}
L59:
	;
	v196 = int32(900)
	if v196 <= v189 {
		goto L60
	} else {
		goto L61
	}
L60:
	;
	v199 = v196
	goto L62
L61:
	;
	v199 = v189
	goto L62
L62:
	;
	v206 = v199 + int32(100)
	goto L55
L63:
	;
	v206 = v189 - int32(1)
	goto L55
L64:
	;
	v212 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+24)) = v180 & int32(-541327359)
	F_ProcSendSignal(m, v212)
	mBase = m.M
	v217 = m.ExcPending
	if v217 != 0 {
		goto L11
	} else {
		goto L67
	}
L65:
	;
	goto L66
L66:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+24)) = v180 & int32(-4194305)
	goto L44
L67:
	;
	goto L44
L68:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v60))) = int32(0)
	*(*int32)(unsafe.Add(mBase, _c_F_UnpinBufferNoOwner[11])) = v60
	goto L13
L69:
	;
	goto L70
L70:
	;
	v238 = *(*int32)(unsafe.Add(mBase, uint32(v60)))
	*(*int32)(unsafe.Add(mBase, uint32(v9)+4)) = v238
	v241 = *(*int32)(unsafe.Add(mBase, _c_F_UnpinBufferNoOwner[9]))
	v247 = F_hash_search(m, v241, v9+int32(4), int32(2), v9+int32(31))
	mBase = m.M
	v248 = m.ExcPending
	if v248 != 0 {
		goto L11
	} else {
		goto L71
	}
L71:
	;
	v249 = int32(_a_F_UnpinBufferNoOwner_16)
	v251 = *(*int32)(unsafe.Add(mBase, _c_F_UnpinBufferNoOwner[8]))
	*(*int32)(unsafe.Add(mBase, _c_F_UnpinBufferNoOwner[8])) = v251 - int32(1)
	goto L13
}
func F_buffer_readv_report(m *base.Module, l0 int32, l1 int32, l2 int32) {
	mBase := m.M
	_ = mBase
	var v12 int32
	_ = v12
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
	var v21 int32
	_ = v21
	var v22 int32
	_ = v22
	var v24 int32
	_ = v24
	var v26 int32
	_ = v26
	var v29 int32
	_ = v29
	var v32 int32
	_ = v32
	var v35 int32
	_ = v35
	var v36 int32
	_ = v36
	var v38 int32
	_ = v38
	var v40 int32
	_ = v40
	var v42 int32
	_ = v42
	var v44 int32
	_ = v44
	var v45 int32
	_ = v45
	var v46 int32
	_ = v46
	var v50 int32
	_ = v50
	var v51 int32
	_ = v51
	var v56 int32
	_ = v56
	var v60 int32
	_ = v60
	var v61 int32
	_ = v61
	var v64 int32
	_ = v64
	var v70 int32
	_ = v70
	var v79 int32
	_ = v79
	var v82 int32
	_ = v82
	var v87 int32
	_ = v87
	var v109 int32
	_ = v109
	var v110 int32
	_ = v110
	var v111 int32
	_ = v111
	var v112 int32
	_ = v112
	var v113 int32
	_ = v113
	var v115 int32
	_ = v115
	var v116 int32
	_ = v116
	var v121 int32
	_ = v121
	var v132 int32
	_ = v132
	var v143 int32
	_ = v143
	var v144 int32
	_ = v144
	var v152 int32
	_ = v152
	var v159 int32
	_ = v159
	var v161 int32
	_ = v161
	var v168 int32
	_ = v168
	v12 = m.G0
	v14 = v12 - int32(208)
	m.G0 = v14
	v16 = *(*int32)(unsafe.Add(mBase, uint32(l1)+12))
	v17 = *(*int32)(unsafe.Add(mBase, uint32(l1)+16))
	v19 = v14 + int32(136)
	v20 = *(*int32)(unsafe.Add(mBase, uint32(l1)+4))
	v21 = *(*int32)(unsafe.Add(mBase, uint32(l1)))
	v22 = *(*int32)(unsafe.Add(mBase, uint32(l1)+8))
	v24 = *(*int32)(unsafe.Add(mBase, _c_F_buffer_readv_report[0]))
	v26 = int32(*(*uint16)(unsafe.Add(mBase, uint32(l1)+20)))
	if v26&int32(256) != 0 {
		v29 = v24
	} else {
		v29 = int32(-1)
	}
	F_GetRelationPath(m, v19, v20, v21, v22, v29, base.I32_extend8_s(v26))
	mBase = m.M
	v32 = m.ExcPending
	if v32 != 0 {
		return
	} else {
		v35 = v16 + v17 - int32(1)
		v36 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
		v38 = int32(base.Ui32(v36) >> (uint(int32(25)) % 32))
		v40 = int32(base.Ui32(v36) >> (uint(int32(18)) % 32))
		v42 = int32(base.Ui32(v36) >> (uint(int32(11)) % 32))
		v44 = v42 & int32(127)
		v45 = int32(1536)
		v46 = v36 & v45
		if v46 == v45 {
			v50 = F_errstart(m, l2, int32(0))
			mBase = m.M
			v51 = m.ExcPending
			if v51 != 0 {
				return
			} else {
				if v50 == int32(0) {
					m.G0 = v14 + int32(208)
					return
				} else {
					F_errcode(m, int32(16779816))
					mBase = m.M
					v56 = m.ExcPending
					if v56 != 0 {
						return
					} else {
						*(*int32)(unsafe.Add(mBase, uint32(v14)+48)) = v19
						*(*int32)(unsafe.Add(mBase, uint32(v14)+44)) = v35
						*(*int32)(unsafe.Add(mBase, uint32(v14)+40)) = v16
						v60 = int32(127)
						v61 = v40 & v60
						*(*int32)(unsafe.Add(mBase, uint32(v14)+36)) = v61
						v64 = v42 & v60
						*(*int32)(unsafe.Add(mBase, uint32(v14)+32)) = v64
						F_errmsg(m, int32(_a_F_buffer_readv_report_0), v14+int32(32))
						mBase = m.M
						v70 = m.ExcPending
						if v70 != 0 {
							return
						} else {
							if base.Ui32(int32(2)) <= base.Ui32(v44) {
								*(*int32)(unsafe.Add(mBase, uint32(v14)+16)) = v16 + v38
								F_errdetail(m, int32(_a_F_buffer_readv_report_1), v14+int32(16))
								mBase = m.M
								v79 = m.ExcPending
								if v79 != 0 {
									return
								} else {
									v82 = v61 + v64 - int32(1)
									*(*int32)(unsafe.Add(mBase, uint32(v14))) = v82
									F_errhint_plural(m, int32(_a_F_buffer_readv_report_2), int32(_a_F_buffer_readv_report_3), v82, v14)
									mBase = m.M
									v87 = m.ExcPending
									if v87 != 0 {
										return
									} else {
										v161 = int32(_a_F_buffer_readv_report_4)
										F_errfinish(m, int32(_a_F_buffer_readv_report_5), v161, int32(_a_F_buffer_readv_report_6))
										mBase = m.M
										v168 = m.ExcPending
										if v168 != 0 {
											return
										} else {
											m.G0 = v14 + int32(208)
											return
										}
									}
								}
							} else {
								v82 = v61 + v64 - int32(1)
								*(*int32)(unsafe.Add(mBase, uint32(v14))) = v82
								F_errhint_plural(m, int32(_a_F_buffer_readv_report_2), int32(_a_F_buffer_readv_report_3), v82, v14)
								mBase = m.M
								v87 = m.ExcPending
								if v87 != 0 {
									return
								} else {
									v161 = int32(_a_F_buffer_readv_report_4)
									F_errfinish(m, int32(_a_F_buffer_readv_report_5), v161, int32(_a_F_buffer_readv_report_6))
									mBase = m.M
									v168 = m.ExcPending
									if v168 != 0 {
										return
									} else {
										m.G0 = v14 + int32(208)
										return
									}
								}
							}
						}
					}
				}
			}
		} else {
			if v36&int32(448) == int32(256) {
				v109 = v44
				v110 = int32(_a_F_buffer_readv_report_7)
				v111 = int32(_a_F_buffer_readv_report_8)
				v112 = int32(_a_F_buffer_readv_report_9)
				v113 = int32(_a_F_buffer_readv_report_10)
			} else {
				if v46 == int32(512) {
					v109 = v44
					v110 = int32(_a_F_buffer_readv_report_11)
					v111 = int32(_a_F_buffer_readv_report_12)
					v112 = int32(_a_F_buffer_readv_report_1)
					v113 = int32(_a_F_buffer_readv_report_13)
				} else {
					v109 = v40 & int32(127)
					v110 = int32(_a_F_buffer_readv_report_14)
					v111 = int32(_a_F_buffer_readv_report_15)
					v112 = int32(_a_F_buffer_readv_report_16)
					v113 = int32(_a_F_buffer_readv_report_17)
				}
			}
			v115 = F_errstart(m, l2, int32(0))
			mBase = m.M
			v116 = m.ExcPending
			if v116 != 0 {
				return
			} else {
				if v115 == int32(0) {
					m.G0 = v14 + int32(208)
					return
				} else {
					F_errcode(m, int32(16779816))
					mBase = m.M
					v121 = m.ExcPending
					if v121 != 0 {
						return
					} else {
						if v109 == int32(1) {
							*(*int32)(unsafe.Add(mBase, uint32(v14)+64)) = v16 + v38
							*(*int32)(unsafe.Add(mBase, uint32(v14)+68)) = v14 + int32(136)
							F_errmsg_internal(m, v110, v14-int32(-64))
							mBase = m.M
							v132 = m.ExcPending
							if v132 != 0 {
								return
							} else {
								v161 = int32(_a_F_buffer_readv_report_18)
								F_errfinish(m, int32(_a_F_buffer_readv_report_5), v161, int32(_a_F_buffer_readv_report_6))
								mBase = m.M
								v168 = m.ExcPending
								if v168 != 0 {
									return
								} else {
									m.G0 = v14 + int32(208)
									return
								}
							}
						} else {
							*(*int32)(unsafe.Add(mBase, uint32(v14)+120)) = v35
							*(*int32)(unsafe.Add(mBase, uint32(v14)+116)) = v16
							*(*int32)(unsafe.Add(mBase, uint32(v14)+112)) = v109
							*(*int32)(unsafe.Add(mBase, uint32(v14)+124)) = v14 + int32(136)
							F_errmsg_internal(m, v113, v14+int32(112))
							mBase = m.M
							v143 = m.ExcPending
							if v143 != 0 {
								return
							} else {
								v144 = int32(_a_F_buffer_readv_report_18)
								if v109 == int32(0) {
									v161 = v144
									F_errfinish(m, int32(_a_F_buffer_readv_report_5), v161, int32(_a_F_buffer_readv_report_6))
									mBase = m.M
									v168 = m.ExcPending
									if v168 != 0 {
										return
									} else {
										m.G0 = v14 + int32(208)
										return
									}
								} else {
									*(*int32)(unsafe.Add(mBase, uint32(v14)+96)) = v16 + v38
									F_errdetail_internal(m, v112, v14+int32(96))
									mBase = m.M
									v152 = m.ExcPending
									if v152 != 0 {
										return
									} else {
										*(*int32)(unsafe.Add(mBase, uint32(v14)+80)) = v109 - int32(1)
										F_errhint_internal(m, v111, v14+int32(80))
										mBase = m.M
										v159 = m.ExcPending
										if v159 != 0 {
											return
										} else {
											v161 = v144
											F_errfinish(m, int32(_a_F_buffer_readv_report_5), v161, int32(_a_F_buffer_readv_report_6))
											mBase = m.M
											v168 = m.ExcPending
											if v168 != 0 {
												return
											} else {
												m.G0 = v14 + int32(208)
												return
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
func F_show_buffer_usage(m *base.Module, l0 int32, l1 int32) {
	mBase := m.M
	_ = mBase
	var v11 int32
	_ = v11
	var v13 int32
	_ = v13
	var v15 int64
	_ = v15
	var v16 int32
	_ = v16
	var v19 int32
	_ = v19
	var v24 int64
	_ = v24
	var v28 int64
	_ = v28
	var v31 int64
	_ = v31
	var v34 int32
	_ = v34
	var v35 int64
	_ = v35
	var v38 int64
	_ = v38
	var v41 int64
	_ = v41
	var v44 int64
	_ = v44
	var v47 int32
	_ = v47
	var v48 int32
	_ = v48
	var v50 int64
	_ = v50
	var v53 int64
	_ = v53
	var v56 int32
	_ = v56
	var v57 int64
	_ = v57
	var v60 int64
	_ = v60
	var v63 int32
	_ = v63
	var v64 int32
	_ = v64
	var v66 int64
	_ = v66
	var v69 int64
	_ = v69
	var v72 int32
	_ = v72
	var v73 int64
	_ = v73
	var v76 int64
	_ = v76
	var v79 int32
	_ = v79
	var v85 int32
	_ = v85
	var v86 int32
	_ = v86
	var v89 int32
	_ = v89
	var v92 int32
	_ = v92
	var v95 int32
	_ = v95
	var v96 int64
	_ = v96
	var v99 int32
	_ = v99
	var v105 int32
	_ = v105
	var v106 int64
	_ = v106
	var v109 int32
	_ = v109
	var v115 int32
	_ = v115
	var v116 int64
	_ = v116
	var v119 int32
	_ = v119
	var v125 int32
	_ = v125
	var v126 int64
	_ = v126
	var v129 int32
	_ = v129
	var v135 int32
	_ = v135
	var v139 int32
	_ = v139
	var v142 int32
	_ = v142
	var v144 int32
	_ = v144
	var v147 int32
	_ = v147
	var v148 int64
	_ = v148
	var v151 int32
	_ = v151
	var v157 int32
	_ = v157
	var v158 int64
	_ = v158
	var v161 int32
	_ = v161
	var v167 int32
	_ = v167
	var v168 int64
	_ = v168
	var v171 int32
	_ = v171
	var v177 int32
	_ = v177
	var v178 int64
	_ = v178
	var v181 int32
	_ = v181
	var v187 int32
	_ = v187
	var v190 int32
	_ = v190
	var v193 int32
	_ = v193
	var v197 int32
	_ = v197
	var v200 int32
	_ = v200
	var v201 int64
	_ = v201
	var v204 int32
	_ = v204
	var v210 int32
	_ = v210
	var v211 int64
	_ = v211
	var v214 int32
	_ = v214
	var v220 int32
	_ = v220
	var v222 int32
	_ = v222
	var v225 int32
	_ = v225
	var v232 int32
	_ = v232
	var v233 int32
	_ = v233
	var v236 int32
	_ = v236
	var v239 int32
	_ = v239
	var v242 int32
	_ = v242
	var v243 int64
	_ = v243
	var v246 int32
	_ = v246
	var v255 int32
	_ = v255
	var v256 int64
	_ = v256
	var v259 int32
	_ = v259
	var v268 int32
	_ = v268
	var v272 int32
	_ = v272
	var v275 int32
	_ = v275
	var v277 int32
	_ = v277
	var v280 int32
	_ = v280
	var v281 int64
	_ = v281
	var v284 int32
	_ = v284
	var v293 int32
	_ = v293
	var v294 int64
	_ = v294
	var v297 int32
	_ = v297
	var v306 int32
	_ = v306
	var v309 int32
	_ = v309
	var v312 int32
	_ = v312
	var v316 int32
	_ = v316
	var v319 int32
	_ = v319
	var v320 int64
	_ = v320
	var v323 int32
	_ = v323
	var v332 int32
	_ = v332
	var v333 int64
	_ = v333
	var v336 int32
	_ = v336
	var v343 int32
	_ = v343
	var v345 int32
	_ = v345
	var v348 int32
	_ = v348
	var v352 int32
	_ = v352
	var v355 int64
	_ = v355
	var v357 int32
	_ = v357
	var v360 int64
	_ = v360
	var v362 int32
	_ = v362
	var v365 int64
	_ = v365
	var v367 int32
	_ = v367
	var v370 int64
	_ = v370
	var v372 int32
	_ = v372
	var v375 int64
	_ = v375
	var v377 int32
	_ = v377
	var v380 int64
	_ = v380
	var v382 int32
	_ = v382
	var v385 int64
	_ = v385
	var v387 int32
	_ = v387
	var v390 int64
	_ = v390
	var v392 int32
	_ = v392
	var v395 int64
	_ = v395
	var v397 int32
	_ = v397
	var v399 int32
	_ = v399
	var v404 int64
	_ = v404
	var v410 int32
	_ = v410
	var v413 int64
	_ = v413
	var v419 int32
	_ = v419
	var v422 int64
	_ = v422
	var v428 int32
	_ = v428
	var v431 int64
	_ = v431
	var v437 int32
	_ = v437
	var v440 int64
	_ = v440
	var v446 int32
	_ = v446
	var v449 int64
	_ = v449
	var v455 int32
	_ = v455
	v11 = m.G0
	v13 = v11 - int32(256)
	m.G0 = v13
	v15 = *(*int64)(unsafe.Add(mBase, uint32(l1)))
	v16 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	if v16 == int32(0) {
		goto L2
	} else {
		goto L3
	}
L1:
	;
	m.G0 = v13 + int32(256)
	return
L2:
	;
	v19 = int32(1)
	if int64(0) < v15 {
		v34 = v19
		goto L5
	} else {
		goto L6
	}
L3:
	;
	goto L4
L4:
	;
	F_ExplainPropertyInteger(m, int32(_a_F_show_buffer_usage_11), int32(0), v15, l0)
	mBase = m.M
	v352 = m.ExcPending
	if v352 != 0 {
		goto L28
	} else {
		goto L126
	}
L5:
	;
	v35 = *(*int64)(unsafe.Add(mBase, uint32(l1)+32))
	if int64(0) < v35 {
		v47 = v19
		goto L9
	} else {
		goto L10
	}
L6:
	;
	v24 = *(*int64)(unsafe.Add(mBase, uint32(l1)+8))
	if int64(0) < v24 {
		v34 = int32(1)
		goto L5
	} else {
		goto L7
	}
L7:
	;
	v28 = *(*int64)(unsafe.Add(mBase, uint32(l1)+16))
	if int64(0) < v28 {
		v34 = int32(1)
		goto L5
	} else {
		goto L8
	}
L8:
	;
	v31 = *(*int64)(unsafe.Add(mBase, uint32(l1)+24))
	v34 = base.B2i32(int64(0) < v31)
	goto L5
L9:
	;
	v48 = int32(1)
	v50 = *(*int64)(unsafe.Add(mBase, uint32(l1)+64))
	if v50 <= int64(0) {
		goto L13
	} else {
		goto L14
	}
L10:
	;
	v38 = *(*int64)(unsafe.Add(mBase, uint32(l1)+40))
	if int64(0) < v38 {
		v47 = v19
		goto L9
	} else {
		goto L11
	}
L11:
	;
	v41 = *(*int64)(unsafe.Add(mBase, uint32(l1)+48))
	if int64(0) < v41 {
		v47 = v19
		goto L9
	} else {
		goto L12
	}
L12:
	;
	v44 = *(*int64)(unsafe.Add(mBase, uint32(l1)+56))
	v47 = base.B2i32(int64(0) < v44)
	goto L9
L13:
	;
	v53 = *(*int64)(unsafe.Add(mBase, uint32(l1)+72))
	v56 = base.B2i32(int64(0) < v53)
	goto L15
L14:
	;
	v56 = v48
	goto L15
L15:
	;
	v57 = *(*int64)(unsafe.Add(mBase, uint32(l1)+80))
	if v57 == int64(0) {
		goto L16
	} else {
		goto L17
	}
L16:
	;
	v60 = *(*int64)(unsafe.Add(mBase, uint32(l1)+88))
	v63 = base.B2i32(v60 != int64(0))
	goto L18
L17:
	;
	v63 = v48
	goto L18
L18:
	;
	v64 = int32(1)
	v66 = *(*int64)(unsafe.Add(mBase, uint32(l1)+96))
	if v66 == int64(0) {
		goto L19
	} else {
		goto L20
	}
L19:
	;
	v69 = *(*int64)(unsafe.Add(mBase, uint32(l1)+104))
	v72 = base.B2i32(v69 != int64(0))
	goto L21
L20:
	;
	v72 = v64
	goto L21
L21:
	;
	v73 = *(*int64)(unsafe.Add(mBase, uint32(l1)+112))
	if v73 == int64(0) {
		goto L22
	} else {
		goto L23
	}
L22:
	;
	v76 = *(*int64)(unsafe.Add(mBase, uint32(l1)+120))
	v79 = base.B2i32(v76 != int64(0))
	goto L24
L23:
	;
	v79 = v64
	goto L24
L24:
	;
	if (v47|v34|v56)&int32(1) != 0 {
		goto L25
	} else {
		goto L26
	}
L25:
	;
	F_ExplainIndentText(m, l0)
	mBase = m.M
	v85 = m.ExcPending
	if v85 != 0 {
		goto L28
	} else {
		goto L29
	}
L26:
	;
	goto L27
L27:
	;
	if v79|(v72|v63) == int32(0) {
		goto L1
	} else {
		goto L85
	}
L28:
	;
	return
L29:
	;
	v86 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	F_appendStringInfoString(m, v86, int32(_a_F_show_buffer_usage_0))
	mBase = m.M
	v89 = m.ExcPending
	if v89 != 0 {
		goto L28
	} else {
		goto L30
	}
L30:
	;
	if v34 == int32(0) {
		goto L31
	} else {
		goto L32
	}
L31:
	;
	if v47 != 0 {
		goto L54
	} else {
		goto L55
	}
L32:
	;
	v92 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	F_appendStringInfoString(m, v92, int32(_a_F_show_buffer_usage_10))
	mBase = m.M
	v95 = m.ExcPending
	if v95 != 0 {
		goto L28
	} else {
		goto L33
	}
L33:
	;
	v96 = *(*int64)(unsafe.Add(mBase, uint32(l1)))
	if int64(0) < v96 {
		goto L34
	} else {
		goto L35
	}
L34:
	;
	v99 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	*(*int64)(unsafe.Add(mBase, uint32(v13)+240)) = v96
	F_appendStringInfo(m, v99, int32(_a_F_show_buffer_usage_2), v13+int32(240))
	mBase = m.M
	v105 = m.ExcPending
	if v105 != 0 {
		goto L28
	} else {
		goto L37
	}
L35:
	;
	goto L36
L36:
	;
	v106 = *(*int64)(unsafe.Add(mBase, uint32(l1)+8))
	if int64(0) < v106 {
		goto L38
	} else {
		goto L39
	}
L37:
	;
	goto L36
L38:
	;
	v109 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	*(*int64)(unsafe.Add(mBase, uint32(v13)+224)) = v106
	F_appendStringInfo(m, v109, int32(_a_F_show_buffer_usage_3), v13+int32(224))
	mBase = m.M
	v115 = m.ExcPending
	if v115 != 0 {
		goto L28
	} else {
		goto L41
	}
L39:
	;
	goto L40
L40:
	;
	v116 = *(*int64)(unsafe.Add(mBase, uint32(l1)+16))
	if int64(0) < v116 {
		goto L42
	} else {
		goto L43
	}
L41:
	;
	goto L40
L42:
	;
	v119 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	*(*int64)(unsafe.Add(mBase, uint32(v13)+208)) = v116
	F_appendStringInfo(m, v119, int32(_a_F_show_buffer_usage_4), v13+int32(208))
	mBase = m.M
	v125 = m.ExcPending
	if v125 != 0 {
		goto L28
	} else {
		goto L45
	}
L43:
	;
	goto L44
L44:
	;
	v126 = *(*int64)(unsafe.Add(mBase, uint32(l1)+24))
	if int64(0) < v126 {
		goto L46
	} else {
		goto L47
	}
L45:
	;
	goto L44
L46:
	;
	v129 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	*(*int64)(unsafe.Add(mBase, uint32(v13)+192)) = v126
	F_appendStringInfo(m, v129, int32(_a_F_show_buffer_usage_5), v13+int32(192))
	mBase = m.M
	v135 = m.ExcPending
	if v135 != 0 {
		goto L28
	} else {
		goto L49
	}
L47:
	;
	goto L48
L48:
	;
	if v56|v47 == int32(0) {
		goto L31
	} else {
		goto L50
	}
L49:
	;
	goto L48
L50:
	;
	v139 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	F_appendStringInfoChar(m, v139, int32(44))
	mBase = m.M
	v142 = m.ExcPending
	if v142 != 0 {
		goto L28
	} else {
		goto L51
	}
L51:
	;
	goto L31
L52:
	;
	v222 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	F_appendStringInfoChar(m, v222, int32(10))
	mBase = m.M
	v225 = m.ExcPending
	if v225 != 0 {
		goto L28
	} else {
		goto L84
	}
L53:
	;
	v197 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	F_appendStringInfoString(m, v197, int32(_a_F_show_buffer_usage_9))
	mBase = m.M
	v200 = m.ExcPending
	if v200 != 0 {
		goto L28
	} else {
		goto L77
	}
L54:
	;
	v144 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	F_appendStringInfoString(m, v144, int32(_a_F_show_buffer_usage_1))
	mBase = m.M
	v147 = m.ExcPending
	if v147 != 0 {
		goto L28
	} else {
		goto L57
	}
L55:
	;
	goto L56
L56:
	;
	if v56 == int32(0) {
		goto L52
	} else {
		goto L76
	}
L57:
	;
	v148 = *(*int64)(unsafe.Add(mBase, uint32(l1)+32))
	if int64(0) < v148 {
		goto L58
	} else {
		goto L59
	}
L58:
	;
	v151 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	*(*int64)(unsafe.Add(mBase, uint32(v13)+176)) = v148
	F_appendStringInfo(m, v151, int32(_a_F_show_buffer_usage_2), v13+int32(176))
	mBase = m.M
	v157 = m.ExcPending
	if v157 != 0 {
		goto L28
	} else {
		goto L61
	}
L59:
	;
	goto L60
L60:
	;
	v158 = *(*int64)(unsafe.Add(mBase, uint32(l1)+40))
	if int64(0) < v158 {
		goto L62
	} else {
		goto L63
	}
L61:
	;
	goto L60
L62:
	;
	v161 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	*(*int64)(unsafe.Add(mBase, uint32(v13)+160)) = v158
	F_appendStringInfo(m, v161, int32(_a_F_show_buffer_usage_3), v13+int32(160))
	mBase = m.M
	v167 = m.ExcPending
	if v167 != 0 {
		goto L28
	} else {
		goto L65
	}
L63:
	;
	goto L64
L64:
	;
	v168 = *(*int64)(unsafe.Add(mBase, uint32(l1)+48))
	if int64(0) < v168 {
		goto L66
	} else {
		goto L67
	}
L65:
	;
	goto L64
L66:
	;
	v171 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	*(*int64)(unsafe.Add(mBase, uint32(v13)+144)) = v168
	F_appendStringInfo(m, v171, int32(_a_F_show_buffer_usage_4), v13+int32(144))
	mBase = m.M
	v177 = m.ExcPending
	if v177 != 0 {
		goto L28
	} else {
		goto L69
	}
L67:
	;
	goto L68
L68:
	;
	v178 = *(*int64)(unsafe.Add(mBase, uint32(l1)+56))
	if int64(0) < v178 {
		goto L70
	} else {
		goto L71
	}
L69:
	;
	goto L68
L70:
	;
	v181 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	*(*int64)(unsafe.Add(mBase, uint32(v13)+128)) = v178
	F_appendStringInfo(m, v181, int32(_a_F_show_buffer_usage_5), v13+int32(128))
	mBase = m.M
	v187 = m.ExcPending
	if v187 != 0 {
		goto L28
	} else {
		goto L73
	}
L71:
	;
	goto L72
L72:
	;
	if v56 == int32(0) {
		goto L52
	} else {
		goto L74
	}
L73:
	;
	goto L72
L74:
	;
	v190 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	F_appendStringInfoChar(m, v190, int32(44))
	mBase = m.M
	v193 = m.ExcPending
	if v193 != 0 {
		goto L28
	} else {
		goto L75
	}
L75:
	;
	goto L53
L76:
	;
	goto L53
L77:
	;
	v201 = *(*int64)(unsafe.Add(mBase, uint32(l1)+64))
	if int64(0) < v201 {
		goto L78
	} else {
		goto L79
	}
L78:
	;
	v204 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	*(*int64)(unsafe.Add(mBase, uint32(v13)+112)) = v201
	F_appendStringInfo(m, v204, int32(_a_F_show_buffer_usage_3), v13+int32(112))
	mBase = m.M
	v210 = m.ExcPending
	if v210 != 0 {
		goto L28
	} else {
		goto L81
	}
L79:
	;
	goto L80
L80:
	;
	v211 = *(*int64)(unsafe.Add(mBase, uint32(l1)+72))
	if v211 <= int64(0) {
		goto L52
	} else {
		goto L82
	}
L81:
	;
	goto L80
L82:
	;
	v214 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	*(*int64)(unsafe.Add(mBase, uint32(v13)+96)) = v211
	F_appendStringInfo(m, v214, int32(_a_F_show_buffer_usage_5), v13+int32(96))
	mBase = m.M
	v220 = m.ExcPending
	if v220 != 0 {
		goto L28
	} else {
		goto L83
	}
L83:
	;
	goto L52
L84:
	;
	goto L27
L85:
	;
	F_ExplainIndentText(m, l0)
	mBase = m.M
	v232 = m.ExcPending
	if v232 != 0 {
		goto L28
	} else {
		goto L86
	}
L86:
	;
	v233 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	F_appendStringInfoString(m, v233, int32(_a_F_show_buffer_usage_6))
	mBase = m.M
	v236 = m.ExcPending
	if v236 != 0 {
		goto L28
	} else {
		goto L87
	}
L87:
	;
	if v63 == int32(0) {
		goto L88
	} else {
		goto L89
	}
L88:
	;
	if v72 != 0 {
		goto L103
	} else {
		goto L104
	}
L89:
	;
	v239 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	F_appendStringInfoString(m, v239, int32(_a_F_show_buffer_usage_10))
	mBase = m.M
	v242 = m.ExcPending
	if v242 != 0 {
		goto L28
	} else {
		goto L90
	}
L90:
	;
	v243 = *(*int64)(unsafe.Add(mBase, uint32(l1)+80))
	if v243 != int64(0) {
		goto L91
	} else {
		goto L92
	}
L91:
	;
	v246 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	*(*float64)(unsafe.Add(mBase, uint32(v13)+80)) = base.F64_div(base.F64_convert_i64_s(v243), float64(1e+06))
	F_appendStringInfo(m, v246, int32(_a_F_show_buffer_usage_7), v13+int32(80))
	mBase = m.M
	v255 = m.ExcPending
	if v255 != 0 {
		goto L28
	} else {
		goto L94
	}
L92:
	;
	goto L93
L93:
	;
	v256 = *(*int64)(unsafe.Add(mBase, uint32(l1)+88))
	if v256 != int64(0) {
		goto L95
	} else {
		goto L96
	}
L94:
	;
	goto L93
L95:
	;
	v259 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	*(*float64)(unsafe.Add(mBase, uint32(v13)+64)) = base.F64_div(base.F64_convert_i64_s(v256), float64(1e+06))
	F_appendStringInfo(m, v259, int32(_a_F_show_buffer_usage_8), v13-int32(-64))
	mBase = m.M
	v268 = m.ExcPending
	if v268 != 0 {
		goto L28
	} else {
		goto L98
	}
L96:
	;
	goto L97
L97:
	;
	if v79|v72 == int32(0) {
		goto L88
	} else {
		goto L99
	}
L98:
	;
	goto L97
L99:
	;
	v272 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	F_appendStringInfoChar(m, v272, int32(44))
	mBase = m.M
	v275 = m.ExcPending
	if v275 != 0 {
		goto L28
	} else {
		goto L100
	}
L100:
	;
	goto L88
L101:
	;
	v345 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	F_appendStringInfoChar(m, v345, int32(10))
	mBase = m.M
	v348 = m.ExcPending
	if v348 != 0 {
		goto L28
	} else {
		goto L125
	}
L102:
	;
	v316 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	F_appendStringInfoString(m, v316, int32(_a_F_show_buffer_usage_9))
	mBase = m.M
	v319 = m.ExcPending
	if v319 != 0 {
		goto L28
	} else {
		goto L118
	}
L103:
	;
	v277 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	F_appendStringInfoString(m, v277, int32(_a_F_show_buffer_usage_1))
	mBase = m.M
	v280 = m.ExcPending
	if v280 != 0 {
		goto L28
	} else {
		goto L106
	}
L104:
	;
	goto L105
L105:
	;
	if v79 == int32(0) {
		goto L101
	} else {
		goto L117
	}
L106:
	;
	v281 = *(*int64)(unsafe.Add(mBase, uint32(l1)+96))
	if v281 != int64(0) {
		goto L107
	} else {
		goto L108
	}
L107:
	;
	v284 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	*(*float64)(unsafe.Add(mBase, uint32(v13)+48)) = base.F64_div(base.F64_convert_i64_s(v281), float64(1e+06))
	F_appendStringInfo(m, v284, int32(_a_F_show_buffer_usage_7), v13+int32(48))
	mBase = m.M
	v293 = m.ExcPending
	if v293 != 0 {
		goto L28
	} else {
		goto L110
	}
L108:
	;
	goto L109
L109:
	;
	v294 = *(*int64)(unsafe.Add(mBase, uint32(l1)+104))
	if v294 != int64(0) {
		goto L111
	} else {
		goto L112
	}
L110:
	;
	goto L109
L111:
	;
	v297 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	*(*float64)(unsafe.Add(mBase, uint32(v13)+32)) = base.F64_div(base.F64_convert_i64_s(v294), float64(1e+06))
	F_appendStringInfo(m, v297, int32(_a_F_show_buffer_usage_8), v13+int32(32))
	mBase = m.M
	v306 = m.ExcPending
	if v306 != 0 {
		goto L28
	} else {
		goto L114
	}
L112:
	;
	goto L113
L113:
	;
	if v79 == int32(0) {
		goto L101
	} else {
		goto L115
	}
L114:
	;
	goto L113
L115:
	;
	v309 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	F_appendStringInfoChar(m, v309, int32(44))
	mBase = m.M
	v312 = m.ExcPending
	if v312 != 0 {
		goto L28
	} else {
		goto L116
	}
L116:
	;
	goto L102
L117:
	;
	goto L102
L118:
	;
	v320 = *(*int64)(unsafe.Add(mBase, uint32(l1)+112))
	if v320 != int64(0) {
		goto L119
	} else {
		goto L120
	}
L119:
	;
	v323 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	*(*float64)(unsafe.Add(mBase, uint32(v13)+16)) = base.F64_div(base.F64_convert_i64_s(v320), float64(1e+06))
	F_appendStringInfo(m, v323, int32(_a_F_show_buffer_usage_7), v13+int32(16))
	mBase = m.M
	v332 = m.ExcPending
	if v332 != 0 {
		goto L28
	} else {
		goto L122
	}
L120:
	;
	goto L121
L121:
	;
	v333 = *(*int64)(unsafe.Add(mBase, uint32(l1)+120))
	if v333 == int64(0) {
		goto L101
	} else {
		goto L123
	}
L122:
	;
	goto L121
L123:
	;
	v336 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	*(*float64)(unsafe.Add(mBase, uint32(v13))) = base.F64_div(base.F64_convert_i64_s(v333), float64(1e+06))
	F_appendStringInfo(m, v336, int32(_a_F_show_buffer_usage_8), v13)
	mBase = m.M
	v343 = m.ExcPending
	if v343 != 0 {
		goto L28
	} else {
		goto L124
	}
L124:
	;
	goto L101
L125:
	;
	goto L1
L126:
	;
	v355 = *(*int64)(unsafe.Add(mBase, uint32(l1)+8))
	F_ExplainPropertyInteger(m, int32(_a_F_show_buffer_usage_12), int32(0), v355, l0)
	mBase = m.M
	v357 = m.ExcPending
	if v357 != 0 {
		goto L28
	} else {
		goto L127
	}
L127:
	;
	v360 = *(*int64)(unsafe.Add(mBase, uint32(l1)+16))
	F_ExplainPropertyInteger(m, int32(_a_F_show_buffer_usage_13), int32(0), v360, l0)
	mBase = m.M
	v362 = m.ExcPending
	if v362 != 0 {
		goto L28
	} else {
		goto L128
	}
L128:
	;
	v365 = *(*int64)(unsafe.Add(mBase, uint32(l1)+24))
	F_ExplainPropertyInteger(m, int32(_a_F_show_buffer_usage_14), int32(0), v365, l0)
	mBase = m.M
	v367 = m.ExcPending
	if v367 != 0 {
		goto L28
	} else {
		goto L129
	}
L129:
	;
	v370 = *(*int64)(unsafe.Add(mBase, uint32(l1)+32))
	F_ExplainPropertyInteger(m, int32(_a_F_show_buffer_usage_15), int32(0), v370, l0)
	mBase = m.M
	v372 = m.ExcPending
	if v372 != 0 {
		goto L28
	} else {
		goto L130
	}
L130:
	;
	v375 = *(*int64)(unsafe.Add(mBase, uint32(l1)+40))
	F_ExplainPropertyInteger(m, int32(_a_F_show_buffer_usage_16), int32(0), v375, l0)
	mBase = m.M
	v377 = m.ExcPending
	if v377 != 0 {
		goto L28
	} else {
		goto L131
	}
L131:
	;
	v380 = *(*int64)(unsafe.Add(mBase, uint32(l1)+48))
	F_ExplainPropertyInteger(m, int32(_a_F_show_buffer_usage_17), int32(0), v380, l0)
	mBase = m.M
	v382 = m.ExcPending
	if v382 != 0 {
		goto L28
	} else {
		goto L132
	}
L132:
	;
	v385 = *(*int64)(unsafe.Add(mBase, uint32(l1)+56))
	F_ExplainPropertyInteger(m, int32(_a_F_show_buffer_usage_18), int32(0), v385, l0)
	mBase = m.M
	v387 = m.ExcPending
	if v387 != 0 {
		goto L28
	} else {
		goto L133
	}
L133:
	;
	v390 = *(*int64)(unsafe.Add(mBase, uint32(l1)+64))
	F_ExplainPropertyInteger(m, int32(_a_F_show_buffer_usage_19), int32(0), v390, l0)
	mBase = m.M
	v392 = m.ExcPending
	if v392 != 0 {
		goto L28
	} else {
		goto L134
	}
L134:
	;
	v395 = *(*int64)(unsafe.Add(mBase, uint32(l1)+72))
	F_ExplainPropertyInteger(m, int32(_a_F_show_buffer_usage_20), int32(0), v395, l0)
	mBase = m.M
	v397 = m.ExcPending
	if v397 != 0 {
		goto L28
	} else {
		goto L135
	}
L135:
	;
	v399 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_show_buffer_usage[0])))
	if v399 != int32(1) {
		goto L1
	} else {
		goto L136
	}
L136:
	;
	v404 = *(*int64)(unsafe.Add(mBase, uint32(l1)+80))
	F_ExplainPropertyFloat(m, int32(_a_F_show_buffer_usage_21), int32(_a_F_show_buffer_usage_22), base.F64_div(base.F64_convert_i64_s(v404), float64(1e+06)), int32(3), l0)
	mBase = m.M
	v410 = m.ExcPending
	if v410 != 0 {
		goto L28
	} else {
		goto L137
	}
L137:
	;
	v413 = *(*int64)(unsafe.Add(mBase, uint32(l1)+88))
	F_ExplainPropertyFloat(m, int32(_a_F_show_buffer_usage_23), int32(_a_F_show_buffer_usage_22), base.F64_div(base.F64_convert_i64_s(v413), float64(1e+06)), int32(3), l0)
	mBase = m.M
	v419 = m.ExcPending
	if v419 != 0 {
		goto L28
	} else {
		goto L138
	}
L138:
	;
	v422 = *(*int64)(unsafe.Add(mBase, uint32(l1)+96))
	F_ExplainPropertyFloat(m, int32(_a_F_show_buffer_usage_24), int32(_a_F_show_buffer_usage_22), base.F64_div(base.F64_convert_i64_s(v422), float64(1e+06)), int32(3), l0)
	mBase = m.M
	v428 = m.ExcPending
	if v428 != 0 {
		goto L28
	} else {
		goto L139
	}
L139:
	;
	v431 = *(*int64)(unsafe.Add(mBase, uint32(l1)+104))
	F_ExplainPropertyFloat(m, int32(_a_F_show_buffer_usage_25), int32(_a_F_show_buffer_usage_22), base.F64_div(base.F64_convert_i64_s(v431), float64(1e+06)), int32(3), l0)
	mBase = m.M
	v437 = m.ExcPending
	if v437 != 0 {
		goto L28
	} else {
		goto L140
	}
L140:
	;
	v440 = *(*int64)(unsafe.Add(mBase, uint32(l1)+112))
	F_ExplainPropertyFloat(m, int32(_a_F_show_buffer_usage_26), int32(_a_F_show_buffer_usage_22), base.F64_div(base.F64_convert_i64_s(v440), float64(1e+06)), int32(3), l0)
	mBase = m.M
	v446 = m.ExcPending
	if v446 != 0 {
		goto L28
	} else {
		goto L141
	}
L141:
	;
	v449 = *(*int64)(unsafe.Add(mBase, uint32(l1)+120))
	F_ExplainPropertyFloat(m, int32(_a_F_show_buffer_usage_27), int32(_a_F_show_buffer_usage_22), base.F64_div(base.F64_convert_i64_s(v449), float64(1e+06)), int32(3), l0)
	mBase = m.M
	v455 = m.ExcPending
	if v455 != 0 {
		goto L28
	} else {
		goto L142
	}
L142:
	;
	goto L1
}
