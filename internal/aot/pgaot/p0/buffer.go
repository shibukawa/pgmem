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
	var v55 int32
	_ = v55
	var v70 int32
	_ = v70
	var v71 int32
	_ = v71
	var v73 int32
	_ = v73
	var v80 int32
	_ = v80
	var v88 int32
	_ = v88
	var v89 int32
	_ = v89
	var v91 int32
	_ = v91
	var v96 int32
	_ = v96
	var v99 int32
	_ = v99
	var v106 int32
	_ = v106
	var v109 int32
	_ = v109
	var v110 int32
	_ = v110
	var v113 int32
	_ = v113
	var v114 int64
	_ = v114
	var v115 int64
	_ = v115
	var v117 int32
	_ = v117
	var v125 int32
	_ = v125
	var v126 int32
	_ = v126
	var v128 int32
	_ = v128
	var v129 int32
	_ = v129
	var v132 int32
	_ = v132
	var v133 int32
	_ = v133
	var v137 int32
	_ = v137
	var v138 int32
	_ = v138
	var v144 int32
	_ = v144
	var v149 int32
	_ = v149
	var v153 int32
	_ = v153
	var v154 int32
	_ = v154
	var v156 int32
	_ = v156
	var v159 int32
	_ = v159
	var v192 int32
	_ = v192
	var v195 int32
	_ = v195
	var v196 int32
	_ = v196
	var v197 int32
	_ = v197
	var v204 int32
	_ = v204
	var v238 int32
	_ = v238
	var v245 int32
	_ = v245
	var v275 int32
	_ = v275
	var v276 int32
	_ = v276
	var v277 int32
	_ = v277
	var v279 int32
	_ = v279
	var v280 int32
	_ = v280
	var v281 int32
	_ = v281
	var v282 int32
	_ = v282
	var v284 int32
	_ = v284
	var v289 int32
	_ = v289
	var v290 int32
	_ = v290
	var v292 int32
	_ = v292
	var v293 int32
	_ = v293
	var v294 int32
	_ = v294
	var v302 int32
	_ = v302
	var v306 int32
	_ = v306
	var v315 int32
	_ = v315
	var v347 int32
	_ = v347
	var v348 int32
	_ = v348
	var v349 int32
	_ = v349
	var v351 int32
	_ = v351
	var v355 int32
	_ = v355
	var v362 int32
	_ = v362
	var v369 int32
	_ = v369
	var v377 int32
	_ = v377
	var v386 int32
	_ = v386
	var v418 int32
	_ = v418
	var v419 int32
	_ = v419
	var v420 int32
	_ = v420
	var v422 int32
	_ = v422
	var v426 int32
	_ = v426
	var v433 int32
	_ = v433
	var v440 int32
	_ = v440
	var v448 int32
	_ = v448
	var v451 int32
	_ = v451
	var v452 int32
	_ = v452
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
	var v516 int32
	_ = v516
	var v517 int32
	_ = v517
	var v521 int32
	_ = v521
	var v523 int32
	_ = v523
	var v526 int32
	_ = v526
	var v528 int32
	_ = v528
	var v531 int32
	_ = v531
	var v533 int32
	_ = v533
	var v537 int64
	_ = v537
	var v538 int64
	_ = v538
	var v542 int64
	_ = v542
	var v546 int32
	_ = v546
	var v547 int32
	_ = v547
	var v553 int32
	_ = v553
	var v554 int32
	_ = v554
	var v556 int32
	_ = v556
	var v557 int64
	_ = v557
	var v561 int32
	_ = v561
	var v563 int32
	_ = v563
	var v569 int64
	_ = v569
	var v570 int64
	_ = v570
	var v574 int64
	_ = v574
	var v581 int32
	_ = v581
	var v583 int64
	_ = v583
	var v585 int64
	_ = v585
	var v588 int32
	_ = v588
	var v590 int64
	_ = v590
	var v593 int32
	_ = v593
	var v595 int64
	_ = v595
	var v618 int32
	_ = v618
	var v620 int32
	_ = v620
	var v625 int64
	_ = v625
	var v629 int32
	_ = v629
	var v641 int64
	_ = v641
	var v645 int32
	_ = v645
	var v657 int32
	_ = v657
	var v662 int64
	_ = v662
	var v666 int64
	_ = v666
	var v671 int32
	_ = v671
	var v679 int32
	_ = v679
	var v681 int64
	_ = v681
	var v685 int32
	_ = v685
	var v686 int32
	_ = v686
	var v690 int32
	_ = v690
	var v692 int32
	_ = v692
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
	v53 = int32(_a_F_FlushBuffer_4)
	v55 = base.AtomicRmwOr32(m, l0, int32(24), v53)
	if v55&v53 != 0 {
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
	v80 = v55
	goto L12
L12:
	;
	v88 = int32(_a_F_FlushBuffer_5)
	v89 = *(*int32)(unsafe.Add(mBase, _c_F_FlushBuffer[1]))
	v91 = *(*int32)(unsafe.Add(mBase, uint32(v10+int32(-24))+8))
	if v91 == int32(0) {
		goto L20
	} else {
		goto L21
	}
L13:
	;
	F_perform_spin_delay(m, v10+int32(-24))
	mBase = m.M
	v70 = m.ExcPending
	if v70 != 0 {
		goto L1
	} else {
		goto L15
	}
L14:
	;
	v80 = v73
	goto L12
L15:
	;
	v71 = int32(_a_F_FlushBuffer_4)
	v73 = base.AtomicRmwOr32(m, l0, int32(24), v71)
	if v73&v71 != 0 {
		goto L13
	} else {
		goto L16
	}
L16:
	;
	goto L14
L17:
	;
	v109 = *(*int32)(unsafe.Add(mBase, _c_F_FlushBuffer[2]))
	v110 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v113 = v109 + v110<<(uint(int32(13))%32)
	v114 = int64(*(*uint32)(unsafe.Add(mBase, uint32(v113)+4)))
	v115 = int64(*(*uint32)(unsafe.Add(mBase, uint32(v113))))
	v117 = v80 & int32(-272629761)
	*(*int32)(unsafe.Add(mBase, uint32(l0)+24)) = v117
	if base.Ui32(int32(2143289344)) <= base.Ui32(v117) {
		goto L28
	} else {
		goto L29
	}
L18:
	;
	goto L17
L19:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_FlushBuffer[1])) = v106
	goto L18
L20:
	;
	if int32(999) < v89 {
		goto L18
	} else {
		goto L23
	}
L21:
	;
	goto L22
L22:
	;
	if v89 < int32(11) {
		goto L18
	} else {
		goto L27
	}
L23:
	;
	v96 = int32(900)
	if v96 <= v89 {
		goto L24
	} else {
		goto L25
	}
L24:
	;
	v99 = v96
	goto L26
L25:
	;
	v99 = v89
	goto L26
L26:
	;
	v106 = v99 + int32(100)
	goto L19
L27:
	;
	v106 = v89 - int32(1)
	goto L19
L28:
	;
	F_XLogFlush(m, v115<<(uint(int64(32))%64)|v114)
	mBase = m.M
	v125 = m.ExcPending
	if v125 != 0 {
		goto L1
	} else {
		goto L31
	}
L29:
	;
	goto L30
L30:
	;
	v126 = *(*int32)(unsafe.Add(mBase, uint32(l0)+16))
	v128 = *(*int32)(unsafe.Add(mBase, _c_F_FlushBuffer[2]))
	v129 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v132 = v128 + v129<<(uint(int32(13))%32)
	v133 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v132)+14)))
	if v133 == int32(0) {
		v526 = v132
		goto L32
	} else {
		goto L33
	}
L31:
	;
	goto L30
L32:
	;
	v528 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_FlushBuffer[3])))
	v531 = m.G0
	v533 = v531 - int32(16)
	m.G0 = v533
	if v528 != 0 {
		goto L54
	} else {
		goto L55
	}
L33:
	;
	v137 = *(*int32)(unsafe.Add(mBase, _c_F_FlushBuffer[4]))
	v138 = *(*int32)(unsafe.Add(mBase, uint32(v137)+252))
	goto L34
L34:
	;
	if base.B2i32(v138 != int32(0)) == int32(0) {
		v526 = v132
		goto L32
	} else {
		goto L35
	}
L35:
	;
	v144 = *(*int32)(unsafe.Add(mBase, _c_F_FlushBuffer[5]))
	if v144 == int32(0) {
		goto L36
	} else {
		goto L37
	}
L36:
	;
	v149 = *(*int32)(unsafe.Add(mBase, _c_F_FlushBuffer[6]))
	v153 = F_MemoryContextAllocAligned(m, v149, int32(_a_F_FlushBuffer_6), int32(_a_F_FlushBuffer_7), int32(0))
	mBase = m.M
	v154 = m.ExcPending
	if v154 != 0 {
		goto L1
	} else {
		goto L39
	}
L37:
	;
	v156 = v144
	goto L38
L38:
	;
	base.MemoryCopy(m, v156, v132, int32(_a_F_FlushBuffer_6))
	v159 = int32(0)
	v192 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v156)+8)))
	*(*uint16)(unsafe.Add(mBase, uint32(v156)+8)) = uint16(v159)
	v195 = m.G0
	v196 = int32(128)
	v197 = v195 - v196
	base.MemoryCopy(m, v197, int32(_a_F_FlushBuffer_8), v196)
	v204 = v159
	goto L41
L39:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_FlushBuffer[5])) = v153
	v156 = v153
	goto L38
L40:
	;
	v523 = *(*int32)(unsafe.Add(mBase, _c_F_FlushBuffer[5]))
	*(*uint16)(unsafe.Add(mBase, uint32(v523)+8)) = uint16(v521)
	v526 = v523
	goto L32
L41:
	;
	v238 = v156 + v204<<(uint(int32(7))%32)
	v245 = int32(0)
	goto L43
L42:
	;
	v315 = int32(0)
	goto L47
L43:
	;
	v275 = int32(2)
	v276 = v245 << (uint(v275) % 32)
	v277 = v197 + v276
	v279 = *(*int32)(unsafe.Add(mBase, uint32(v238+v276)))
	v280 = *(*int32)(unsafe.Add(mBase, uint32(v277)))
	v281 = v279 ^ v280
	v282 = int32(16777619)
	v284 = int32(17)
	*(*int32)(unsafe.Add(mBase, uint32(v277))) = v281*v282 ^ int32(base.Ui32(v281)>>(uint(v284)%32))
	v289 = v276 | int32(4)
	v290 = v197 + v289
	v292 = *(*int32)(unsafe.Add(mBase, uint32(v238+v289)))
	v293 = *(*int32)(unsafe.Add(mBase, uint32(v290)))
	v294 = v292 ^ v293
	*(*int32)(unsafe.Add(mBase, uint32(v290))) = v294*v282 ^ int32(base.Ui32(v294)>>(uint(v284)%32))
	v302 = v245 + v275
	if v302 != int32(32) {
		v245 = v302
		goto L43
	} else {
		goto L45
	}
L44:
	;
	v306 = v204 + int32(1)
	if v306 != int32(64) {
		v204 = v306
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
	v347 = v197 + v315<<(uint(int32(2))%32)
	v348 = *(*int32)(unsafe.Add(mBase, uint32(v347)))
	v349 = int32(16777619)
	v351 = int32(17)
	*(*int32)(unsafe.Add(mBase, uint32(v347))) = v348*v349 ^ int32(base.Ui32(v348)>>(uint(v351)%32))
	v355 = *(*int32)(unsafe.Add(mBase, uint32(v347)+4))
	*(*int32)(unsafe.Add(mBase, uint32(v347)+4)) = v355*v349 ^ int32(base.Ui32(v355)>>(uint(v351)%32))
	v362 = *(*int32)(unsafe.Add(mBase, uint32(v347)+8))
	*(*int32)(unsafe.Add(mBase, uint32(v347)+8)) = v362*v349 ^ int32(base.Ui32(v362)>>(uint(v351)%32))
	v369 = *(*int32)(unsafe.Add(mBase, uint32(v347)+12))
	*(*int32)(unsafe.Add(mBase, uint32(v347)+12)) = v369*v349 ^ int32(base.Ui32(v369)>>(uint(v351)%32))
	v377 = v315 + int32(4)
	if v377 != int32(32) {
		v315 = v377
		goto L47
	} else {
		goto L49
	}
L48:
	;
	v386 = int32(0)
	goto L50
L49:
	;
	goto L48
L50:
	;
	v418 = v197 + v386<<(uint(int32(2))%32)
	v419 = *(*int32)(unsafe.Add(mBase, uint32(v418)))
	v420 = int32(16777619)
	v422 = int32(17)
	*(*int32)(unsafe.Add(mBase, uint32(v418))) = v419*v420 ^ int32(base.Ui32(v419)>>(uint(v422)%32))
	v426 = *(*int32)(unsafe.Add(mBase, uint32(v418)+4))
	*(*int32)(unsafe.Add(mBase, uint32(v418)+4)) = v426*v420 ^ int32(base.Ui32(v426)>>(uint(v422)%32))
	v433 = *(*int32)(unsafe.Add(mBase, uint32(v418)+8))
	*(*int32)(unsafe.Add(mBase, uint32(v418)+8)) = v433*v420 ^ int32(base.Ui32(v433)>>(uint(v422)%32))
	v440 = *(*int32)(unsafe.Add(mBase, uint32(v418)+12))
	*(*int32)(unsafe.Add(mBase, uint32(v418)+12)) = v440*v420 ^ int32(base.Ui32(v440)>>(uint(v422)%32))
	v448 = v386 + int32(4)
	if v448 != int32(32) {
		v386 = v448
		goto L50
	} else {
		goto L52
	}
L51:
	;
	v451 = *(*int32)(unsafe.Add(mBase, uint32(v197)+124))
	v452 = *(*int32)(unsafe.Add(mBase, uint32(v197)+120))
	v453 = *(*int32)(unsafe.Add(mBase, uint32(v197)+116))
	v454 = *(*int32)(unsafe.Add(mBase, uint32(v197)+112))
	v455 = *(*int32)(unsafe.Add(mBase, uint32(v197)+108))
	v456 = *(*int32)(unsafe.Add(mBase, uint32(v197)+104))
	v457 = *(*int32)(unsafe.Add(mBase, uint32(v197)+100))
	v458 = *(*int32)(unsafe.Add(mBase, uint32(v197)+96))
	v459 = *(*int32)(unsafe.Add(mBase, uint32(v197)+92))
	v460 = *(*int32)(unsafe.Add(mBase, uint32(v197)+88))
	v461 = *(*int32)(unsafe.Add(mBase, uint32(v197)+84))
	v462 = *(*int32)(unsafe.Add(mBase, uint32(v197)+80))
	v463 = *(*int32)(unsafe.Add(mBase, uint32(v197)+76))
	v464 = *(*int32)(unsafe.Add(mBase, uint32(v197)+72))
	v465 = *(*int32)(unsafe.Add(mBase, uint32(v197)+68))
	v466 = *(*int32)(unsafe.Add(mBase, uint32(v197)+64))
	v467 = *(*int32)(unsafe.Add(mBase, uint32(v197)+60))
	v468 = *(*int32)(unsafe.Add(mBase, uint32(v197)+56))
	v469 = *(*int32)(unsafe.Add(mBase, uint32(v197)+52))
	v470 = *(*int32)(unsafe.Add(mBase, uint32(v197)+48))
	v471 = *(*int32)(unsafe.Add(mBase, uint32(v197)+44))
	v472 = *(*int32)(unsafe.Add(mBase, uint32(v197)+40))
	v473 = *(*int32)(unsafe.Add(mBase, uint32(v197)+36))
	v474 = *(*int32)(unsafe.Add(mBase, uint32(v197)+32))
	v475 = *(*int32)(unsafe.Add(mBase, uint32(v197)+28))
	v476 = *(*int32)(unsafe.Add(mBase, uint32(v197)+24))
	v477 = *(*int32)(unsafe.Add(mBase, uint32(v197)+20))
	v478 = *(*int32)(unsafe.Add(mBase, uint32(v197)+16))
	v479 = *(*int32)(unsafe.Add(mBase, uint32(v197)+12))
	v480 = *(*int32)(unsafe.Add(mBase, uint32(v197)+8))
	v481 = *(*int32)(unsafe.Add(mBase, uint32(v197)+4))
	v482 = *(*int32)(unsafe.Add(mBase, uint32(v197)))
	*(*uint16)(unsafe.Add(mBase, uint32(v156)+8)) = uint16(v192)
	v516 = int32(_a_F_FlushBuffer_9)
	v517 = base.I32_rem_u_s(v451^(v452^(v453^(v454^(v455^(v456^(v457^(v458^(v459^(v460^(v461^(v462^(v463^(v464^(v465^(v466^(v467^(v468^(v469^(v470^(v471^(v472^(v473^(v474^(v475^(v476^(v477^(v478^(v479^(v480^(v481^(v126^v482))))))))))))))))))))))))))))))), v516)
	v521 = (v517 + int32(1)) & v516
	goto L40
L52:
	;
	goto L51
L53:
	;
	v546 = *(*int32)(unsafe.Add(mBase, uint32(l0)+16))
	v547 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	*(*int32)(unsafe.Add(mBase, uint32(v12)+40)) = v526
	F_smgrwritev(m, v42, v547, v546, v10+int32(-24), int32(0))
	mBase = m.M
	v553 = m.ExcPending
	if v553 != 0 {
		goto L1
	} else {
		goto L57
	}
L54:
	;
	F___clock_gettime(m, int32(1), v533)
	mBase = m.M
	v537 = int64(*(*int32)(unsafe.Add(mBase, uint32(v533)+8)))
	v538 = *(*int64)(unsafe.Add(mBase, uint32(v533)))
	v542 = v537 + v538*int64(1000000000)
	goto L56
L55:
	;
	v542 = int64(0)
	goto L56
L56:
	;
	m.G0 = v533 + int32(16)
	goto L53
L57:
	;
	v554 = int32(0)
	v556 = int32(1)
	v557 = int64(8192)
	v561 = m.G0
	v563 = v561 - int32(16)
	m.G0 = v563
	if v542 != int64(0) {
		goto L59
	} else {
		goto L60
	}
L58:
	;
	v679 = int32(_a_F_FlushBuffer_10)
	v681 = *(*int64)(unsafe.Add(mBase, _c_F_FlushBuffer[7]))
	*(*int64)(unsafe.Add(mBase, _c_F_FlushBuffer[7])) = v681 + int64(1)
	v685 = int32(1)
	v686 = int32(0)
	F_TerminateBufferIO(m, l0, v685, v686, v685, v686)
	mBase = m.M
	v690 = m.ExcPending
	if v690 != 0 {
		goto L1
	} else {
		goto L75
	}
L59:
	;
	F___clock_gettime(m, int32(1), v563)
	mBase = m.M
	v569 = int64(*(*int32)(unsafe.Add(mBase, uint32(v563)+8)))
	v570 = *(*int64)(unsafe.Add(mBase, uint32(v563)))
	v574 = v569 + (v570*int64(1000000000) - v542)
	goto L63
L60:
	;
	goto L61
L61:
	;
	v657 = l2 << (uint(int32(6)) % 32)
	v662 = *(*int64)(unsafe.Add(mBase, uint32(v657)+uint32(_c_F_FlushBuffer[8])))
	*(*int64)(unsafe.Add(mBase, uint32(v657)+uint32(_c_F_FlushBuffer[8]))) = v662 + base.I64_extend_i32_u(v556)
	v666 = *(*int64)(unsafe.Add(mBase, uint32(v657)+uint32(_c_F_FlushBuffer[9])))
	*(*int64)(unsafe.Add(mBase, uint32(v657)+uint32(_c_F_FlushBuffer[9]))) = v666 + v557
	F_pgstat_count_backend_io_op(m, v554, l2, int32(7), v556, v557)
	mBase = m.M
	v671 = int32(1)
	*(*uint8)(unsafe.Add(mBase, _c_F_FlushBuffer[10])) = uint8(v671)
	*(*uint8)(unsafe.Add(mBase, _c_F_FlushBuffer[11])) = uint8(v671)
	m.G0 = v563 + int32(16)
	goto L58
L62:
	;
	v618 = int32(0)
	v620 = l2 << (uint(int32(6)) % 32)
	v625 = *(*int64)(unsafe.Add(mBase, uint32(v620)+uint32(_c_F_FlushBuffer[12])))
	*(*int64)(unsafe.Add(mBase, uint32(v620)+uint32(_c_F_FlushBuffer[12]))) = v625 + v574
	v629 = *(*int32)(unsafe.Add(mBase, _c_F_FlushBuffer[13]))
	if base.B2i32(base.Ui32(int32(16)) < base.Ui32(v629))|base.B2i32(int32(1)<<(uint(v629)%32)&int32(_a_F_FlushBuffer_11) == v618) == v618 {
		goto L72
	} else {
		goto L73
	}
L63:
	;
	goto L64
L64:
	;
	v581 = int32(_a_F_FlushBuffer_12)
	v583 = *(*int64)(unsafe.Add(mBase, _c_F_FlushBuffer[14]))
	v585 = base.I64_div_s(v574, int64(1000))
	*(*int64)(unsafe.Add(mBase, _c_F_FlushBuffer[14])) = v583 + v585
	switch v554 {
	case 0:
		goto L68
	case 1:
		goto L67
	default:
		goto L62
	}
L67:
	;
	v593 = int32(_a_F_FlushBuffer_13)
	v595 = *(*int64)(unsafe.Add(mBase, _c_F_FlushBuffer[15]))
	*(*int64)(unsafe.Add(mBase, _c_F_FlushBuffer[15])) = v595 + v574
	goto L62
L68:
	;
	v588 = int32(_a_F_FlushBuffer_14)
	v590 = *(*int64)(unsafe.Add(mBase, _c_F_FlushBuffer[16]))
	*(*int64)(unsafe.Add(mBase, _c_F_FlushBuffer[16])) = v590 + v574
	goto L62
L72:
	;
	v641 = *(*int64)(unsafe.Add(mBase, uint32(v620)+uint32(_c_F_FlushBuffer[17])))
	*(*int64)(unsafe.Add(mBase, uint32(v620)+uint32(_c_F_FlushBuffer[17]))) = v641 + v574
	v645 = int32(1)
	*(*uint8)(unsafe.Add(mBase, _c_F_FlushBuffer[10])) = uint8(v645)
	*(*uint8)(unsafe.Add(mBase, _c_F_FlushBuffer[18])) = uint8(v645)
	goto L74
L73:
	;
	goto L74
L74:
	;
	goto L61
L75:
	;
	v692 = *(*int32)(unsafe.Add(mBase, uint32(v12)+28))
	*(*int32)(unsafe.Add(mBase, _c_F_FlushBuffer[0])) = v692
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
	var v138 int32
	_ = v138
	var v154 int32
	_ = v154
	var v156 int32
	_ = v156
	var v168 int32
	_ = v168
	var v169 int32
	_ = v169
	var v171 int32
	_ = v171
	var v175 int32
	_ = v175
	var v183 int32
	_ = v183
	var v184 int32
	_ = v184
	var v186 int32
	_ = v186
	var v191 int32
	_ = v191
	var v194 int32
	_ = v194
	var v201 int32
	_ = v201
	var v207 int32
	_ = v207
	var v212 int32
	_ = v212
	var v233 int32
	_ = v233
	var v236 int32
	_ = v236
	var v242 int32
	_ = v242
	var v243 int32
	_ = v243
	var v244 int32
	_ = v244
	var v246 int32
	_ = v246
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
		goto L41
	} else {
		goto L42
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
	v138 = base.AtomicRmwCmpxchg32(m, l0, int32(24), v130, v136)
	if v130 != v138 {
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
	v67 = v138
	goto L15
L39:
	;
	goto L40
L40:
	;
	goto L16
L41:
	;
	if base.B2i32(base.Ui32(v60) < base.Ui32(int32(_a_F_UnpinBufferNoOwner_0)))|base.B2i32(base.Ui32(int32(_a_F_UnpinBufferNoOwner_13)) <= base.Ui32(v60)) == int32(0) {
		goto L65
	} else {
		goto L66
	}
L42:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v9)+24)) = int32(_a_F_UnpinBufferNoOwner_14)
	*(*int32)(unsafe.Add(mBase, uint32(v9)+20)) = int32(_a_F_UnpinBufferNoOwner_15)
	*(*int32)(unsafe.Add(mBase, uint32(v9)+16)) = int32(_a_F_UnpinBufferNoOwner_11)
	*(*int32)(unsafe.Add(mBase, uint32(v9)+12)) = int32(0)
	*(*int64)(unsafe.Add(mBase, uint32(v9)+4)) = int64(0)
	v154 = int32(_a_F_UnpinBufferNoOwner_8)
	v156 = base.AtomicRmwOr32(m, l0, int32(24), v154)
	if v156&v154 != 0 {
		goto L43
	} else {
		goto L44
	}
L43:
	;
	goto L46
L44:
	;
	v175 = v156
	goto L45
L45:
	;
	v183 = int32(_a_F_UnpinBufferNoOwner_12)
	v184 = *(*int32)(unsafe.Add(mBase, _c_F_UnpinBufferNoOwner[10]))
	v186 = *(*int32)(unsafe.Add(mBase, uint32(v9+int32(4))+8))
	if v186 == int32(0) {
		goto L53
	} else {
		goto L54
	}
L46:
	;
	F_perform_spin_delay(m, v9+int32(4))
	mBase = m.M
	v168 = m.ExcPending
	if v168 != 0 {
		goto L11
	} else {
		goto L48
	}
L47:
	;
	v175 = v171
	goto L45
L48:
	;
	v169 = int32(_a_F_UnpinBufferNoOwner_8)
	v171 = base.AtomicRmwOr32(m, l0, int32(24), v169)
	if v171&v169 != 0 {
		goto L46
	} else {
		goto L49
	}
L49:
	;
	goto L47
L50:
	;
	if v175&int32(537133055) == int32(536870913) {
		goto L61
	} else {
		goto L62
	}
L51:
	;
	goto L50
L52:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_UnpinBufferNoOwner[10])) = v201
	goto L51
L53:
	;
	if int32(999) < v184 {
		goto L51
	} else {
		goto L56
	}
L54:
	;
	goto L55
L55:
	;
	if v184 < int32(11) {
		goto L51
	} else {
		goto L60
	}
L56:
	;
	v191 = int32(900)
	if v191 <= v184 {
		goto L57
	} else {
		goto L58
	}
L57:
	;
	v194 = v191
	goto L59
L58:
	;
	v194 = v184
	goto L59
L59:
	;
	v201 = v194 + int32(100)
	goto L52
L60:
	;
	v201 = v184 - int32(1)
	goto L52
L61:
	;
	v207 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+24)) = v175 & int32(-541327359)
	F_ProcSendSignal(m, v207)
	mBase = m.M
	v212 = m.ExcPending
	if v212 != 0 {
		goto L11
	} else {
		goto L64
	}
L62:
	;
	goto L63
L63:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+24)) = v175 & int32(-4194305)
	goto L41
L64:
	;
	goto L41
L65:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v60))) = int32(0)
	*(*int32)(unsafe.Add(mBase, _c_F_UnpinBufferNoOwner[11])) = v60
	goto L13
L66:
	;
	goto L67
L67:
	;
	v233 = *(*int32)(unsafe.Add(mBase, uint32(v60)))
	*(*int32)(unsafe.Add(mBase, uint32(v9)+4)) = v233
	v236 = *(*int32)(unsafe.Add(mBase, _c_F_UnpinBufferNoOwner[9]))
	v242 = F_hash_search(m, v236, v9+int32(4), int32(2), v9+int32(31))
	mBase = m.M
	v243 = m.ExcPending
	if v243 != 0 {
		goto L11
	} else {
		goto L68
	}
L68:
	;
	v244 = int32(_a_F_UnpinBufferNoOwner_16)
	v246 = *(*int32)(unsafe.Add(mBase, _c_F_UnpinBufferNoOwner[8]))
	*(*int32)(unsafe.Add(mBase, _c_F_UnpinBufferNoOwner[8])) = v246 - int32(1)
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
