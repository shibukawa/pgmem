package p5

import (
	base "github.com/shibukawa/pgmem/internal/aot/pgaot/base"
	"unsafe"
)

func F_ExecBuildUpdateProjection(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32, l5 int32, l6 int32) int32 {
	mBase := m.M
	_ = mBase
	var v8 int32
	_ = v8
	var v19 int32
	_ = v19
	var v21 int32
	_ = v21
	var v24 int32
	_ = v24
	var v27 int32
	_ = v27
	var v30 int64
	_ = v30
	var v38 int32
	_ = v38
	var v42 int32
	_ = v42
	var v48 int32
	_ = v48
	var v51 int32
	_ = v51
	var v54 int32
	_ = v54
	var v55 int32
	_ = v55
	var v56 int32
	_ = v56
	var v63 int32
	_ = v63
	var v64 int32
	_ = v64
	var v66 int32
	_ = v66
	var v79 int32
	_ = v79
	var v80 int32
	_ = v80
	var v87 int32
	_ = v87
	var v89 int32
	_ = v89
	var v99 int32
	_ = v99
	var v139 int32
	_ = v139
	var v149 int32
	_ = v149
	var v172 int32
	_ = v172
	var v176 int32
	_ = v176
	var v181 int32
	_ = v181
	var v185 int32
	_ = v185
	var v189 int32
	_ = v189
	var v194 int32
	_ = v194
	var v195 int32
	_ = v195
	var v205 int32
	_ = v205
	var v212 int32
	_ = v212
	var v217 int32
	_ = v217
	var v221 int32
	_ = v221
	var v222 int32
	_ = v222
	var v223 int32
	_ = v223
	var v225 int32
	_ = v225
	var v226 int32
	_ = v226
	var v241 int32
	_ = v241
	var v254 int32
	_ = v254
	var v257 int32
	_ = v257
	var v259 int32
	_ = v259
	var v266 int32
	_ = v266
	var v275 int32
	_ = v275
	var v288 int32
	_ = v288
	var v292 int32
	_ = v292
	var v293 int32
	_ = v293
	var v294 int32
	_ = v294
	var v317 int32
	_ = v317
	var v320 int32
	_ = v320
	var v321 int32
	_ = v321
	var v326 int32
	_ = v326
	var v328 int32
	_ = v328
	var v330 int32
	_ = v330
	var v331 int32
	_ = v331
	var v340 int32
	_ = v340
	var v342 int32
	_ = v342
	var v344 int32
	_ = v344
	var v349 int32
	_ = v349
	var v352 int32
	_ = v352
	var v356 int32
	_ = v356
	var v358 int32
	_ = v358
	var v362 int32
	_ = v362
	var v365 int32
	_ = v365
	var v368 int32
	_ = v368
	var v371 int32
	_ = v371
	var v372 int32
	_ = v372
	var v379 int32
	_ = v379
	var v382 int32
	_ = v382
	var v389 int32
	_ = v389
	var v390 int32
	_ = v390
	var v393 int32
	_ = v393
	var v394 int32
	_ = v394
	var v395 int32
	_ = v395
	var v396 int32
	_ = v396
	var v398 int32
	_ = v398
	var v399 int32
	_ = v399
	var v402 int32
	_ = v402
	var v403 int32
	_ = v403
	var v405 int32
	_ = v405
	var v406 int32
	_ = v406
	var v412 int32
	_ = v412
	var v413 int32
	_ = v413
	var v414 int32
	_ = v414
	var v416 int32
	_ = v416
	var v420 int32
	_ = v420
	var v423 int32
	_ = v423
	var v424 int32
	_ = v424
	var v425 int32
	_ = v425
	var v427 int32
	_ = v427
	var v428 int32
	_ = v428
	var v434 int32
	_ = v434
	var v446 int32
	_ = v446
	var v452 int32
	_ = v452
	var v453 int32
	_ = v453
	var v454 int32
	_ = v454
	var v456 int32
	_ = v456
	var v460 int32
	_ = v460
	var v463 int32
	_ = v463
	var v464 int32
	_ = v464
	var v465 int32
	_ = v465
	var v467 int32
	_ = v467
	var v470 int32
	_ = v470
	var v476 int32
	_ = v476
	var v483 int32
	_ = v483
	var v485 int32
	_ = v485
	var v486 int32
	_ = v486
	var v487 int64
	_ = v487
	var v495 int32
	_ = v495
	var v499 int32
	_ = v499
	var v502 int32
	_ = v502
	var v504 int32
	_ = v504
	var v505 int32
	_ = v505
	var v509 int32
	_ = v509
	var v512 int32
	_ = v512
	var v513 int32
	_ = v513
	var v517 int32
	_ = v517
	var v520 int32
	_ = v520
	var v526 int32
	_ = v526
	var v527 int32
	_ = v527
	var v528 int32
	_ = v528
	var v530 int32
	_ = v530
	var v534 int32
	_ = v534
	var v537 int32
	_ = v537
	var v538 int32
	_ = v538
	var v539 int32
	_ = v539
	var v541 int32
	_ = v541
	var v542 int32
	_ = v542
	var v543 int32
	_ = v543
	var v548 int32
	_ = v548
	var v552 int32
	_ = v552
	var v560 int64
	_ = v560
	var v562 int64
	_ = v562
	var v564 int32
	_ = v564
	var v570 int32
	_ = v570
	var v571 int32
	_ = v571
	var v573 int32
	_ = v573
	var v575 int32
	_ = v575
	var v576 int32
	_ = v576
	var v578 int32
	_ = v578
	var v582 int32
	_ = v582
	var v585 int32
	_ = v585
	var v586 int32
	_ = v586
	var v588 int32
	_ = v588
	var v589 int32
	_ = v589
	var v592 int32
	_ = v592
	var v598 int32
	_ = v598
	var v599 int32
	_ = v599
	var v600 int32
	_ = v600
	var v602 int32
	_ = v602
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
	var v615 int32
	_ = v615
	var v616 int32
	_ = v616
	var v619 int32
	_ = v619
	var v620 int32
	_ = v620
	var v621 int32
	_ = v621
	var v622 int32
	_ = v622
	var v623 int32
	_ = v623
	var v624 int32
	_ = v624
	var v625 int32
	_ = v625
	var v626 int32
	_ = v626
	var v632 int32
	_ = v632
	var v640 int64
	_ = v640
	var v642 int64
	_ = v642
	var v645 int32
	_ = v645
	var v647 int32
	_ = v647
	var v649 int32
	_ = v649
	var v650 int32
	_ = v650
	var v651 int32
	_ = v651
	var v652 int32
	_ = v652
	var v654 int32
	_ = v654
	var v655 int32
	_ = v655
	var v660 int32
	_ = v660
	var v663 int32
	_ = v663
	var v667 int32
	_ = v667
	var v668 int32
	_ = v668
	var v669 int32
	_ = v669
	var v670 int32
	_ = v670
	var v671 int32
	_ = v671
	var v672 int32
	_ = v672
	var v673 int32
	_ = v673
	var v674 int32
	_ = v674
	var v675 int32
	_ = v675
	var v683 int32
	_ = v683
	var v688 int32
	_ = v688
	var v692 int32
	_ = v692
	var v695 int32
	_ = v695
	var v699 int32
	_ = v699
	var v703 int32
	_ = v703
	var v708 int32
	_ = v708
	var v713 int32
	_ = v713
	var v716 int32
	_ = v716
	var v720 int32
	_ = v720
	var v724 int32
	_ = v724
	var v729 int32
	_ = v729
	var v731 int32
	_ = v731
	var v734 int32
	_ = v734
	var v740 int32
	_ = v740
	var v741 int32
	_ = v741
	var v745 int32
	_ = v745
	var v748 int32
	_ = v748
	var v754 int32
	_ = v754
	var v755 int32
	_ = v755
	var v756 int32
	_ = v756
	var v758 int32
	_ = v758
	var v762 int32
	_ = v762
	var v765 int32
	_ = v765
	var v766 int32
	_ = v766
	var v767 int32
	_ = v767
	var v769 int32
	_ = v769
	var v770 int32
	_ = v770
	var v771 int32
	_ = v771
	var v776 int32
	_ = v776
	var v790 int64
	_ = v790
	var v792 int64
	_ = v792
	var v794 int32
	_ = v794
	var v795 int32
	_ = v795
	var v799 int32
	_ = v799
	v8 = int32(0)
	v19 = m.G0
	v21 = v19 + int32(-64)
	m.G0 = v21
	v24 = F_palloc0(m, int32(76))
	mBase = m.M
	v27 = m.ExcPending
	if v27 != 0 {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	return int32(0)
L2:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v24))) = int32(384)
	v30 = int64(0)
	*(*int64)(unsafe.Add(mBase, uint32(v21)+56)) = v30
	*(*int64)(unsafe.Add(mBase, uint32(v21)+48)) = v30
	*(*int64)(unsafe.Add(mBase, uint32(v21)+40)) = v30
	*(*int64)(unsafe.Add(mBase, uint32(v21)+32)) = v30
	v38 = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v24)+48)) = v38
	*(*int32)(unsafe.Add(mBase, uint32(v24)+44)) = l6
	if l1 != 0 {
		goto L3
	} else {
		goto L4
	}
L3:
	;
	v42 = l0
	goto L5
L4:
	;
	v42 = v38
	goto L5
L5:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v24)+28)) = v42
	*(*int32)(unsafe.Add(mBase, uint32(v24)+4)) = int32(380)
	*(*int32)(unsafe.Add(mBase, uint32(v24)+72)) = l4
	*(*int32)(unsafe.Add(mBase, uint32(v24)+16)) = l5
	if l0 != 0 {
		goto L12
	} else {
		goto L13
	}
L6:
	;
	v266 = *(*int32)(unsafe.Add(mBase, uint32(l3)))
	if v266 <= int32(0) {
		goto L45
	} else {
		goto L46
	}
L7:
	;
	v195 = *(*int32)(unsafe.Add(mBase, uint32(l2)+4))
	if int32(0) < v195 {
		goto L38
	} else {
		goto L39
	}
L8:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v185 = m.ExcPending
	if v185 != 0 {
		goto L1
	} else {
		goto L35
	}
L9:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v172 = m.ExcPending
	if v172 != 0 {
		goto L1
	} else {
		goto L32
	}
L10:
	;
	v149 = *(*int32)(unsafe.Add(mBase, uint32(l2)+4))
	if v139 == v149 {
		goto L7
	} else {
		goto L31
	}
L11:
	;
	v254 = int32(0)
	v257 = int32(1)
	v259 = v8
	goto L6
L12:
	;
	v48 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	if int32(0) < v48 {
		goto L15
	} else {
		goto L16
	}
L13:
	;
	goto L14
L14:
	;
	if l2 != 0 {
		v139 = v8
		goto L10
	} else {
		goto L30
	}
L15:
	;
	v51 = int32(0)
	if v51 < v48 {
		goto L18
	} else {
		goto L19
	}
L16:
	;
	v99 = v8
	goto L17
L17:
	;
	if l2 != 0 {
		v139 = v99
		goto L10
	} else {
		goto L28
	}
L18:
	;
	v54 = v48
	goto L20
L19:
	;
	v54 = v51
	goto L20
L20:
	;
	v55 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	v56 = int32(0)
	v63 = v56
	v64 = v56
	v66 = v8
	goto L21
L21:
	;
	v79 = *(*int32)(unsafe.Add(mBase, uint32(v55+v64<<(uint(int32(2))%32))))
	v80 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v79)+26)))
	if v80 == int32(0) {
		goto L23
	} else {
		goto L24
	}
L22:
	;
	v99 = v87
	goto L17
L23:
	;
	if v63&int32(1) != 0 {
		goto L8
	} else {
		goto L26
	}
L24:
	;
	v87 = v66
	goto L25
L25:
	;
	v89 = v64 + int32(1)
	if v89 != v54 {
		v63 = v80
		v64 = v89
		v66 = v87
		goto L21
	} else {
		goto L27
	}
L26:
	;
	v87 = v66 + int32(1)
	goto L25
L27:
	;
	goto L22
L28:
	;
	if v99 == int32(0) {
		goto L11
	} else {
		goto L29
	}
L29:
	;
	goto L9
L30:
	;
	goto L11
L31:
	;
	goto L9
L32:
	;
	F_errmsg_internal(m, int32(_a_F_ExecBuildUpdateProjection_0), int32(0))
	mBase = m.M
	v176 = m.ExcPending
	if v176 != 0 {
		goto L1
	} else {
		goto L33
	}
L33:
	;
	F_errfinish(m, int32(_a_F_ExecBuildUpdateProjection_1), int32(601), int32(_a_F_ExecBuildUpdateProjection_2))
	mBase = m.M
	v181 = m.ExcPending
	if v181 != 0 {
		goto L1
	} else {
		goto L34
	}
L34:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L35:
	;
	F_errmsg_internal(m, int32(_a_F_ExecBuildUpdateProjection_3), int32(0))
	mBase = m.M
	v189 = m.ExcPending
	if v189 != 0 {
		goto L1
	} else {
		goto L36
	}
L36:
	;
	F_errfinish(m, int32(_a_F_ExecBuildUpdateProjection_1), int32(594), int32(_a_F_ExecBuildUpdateProjection_2))
	mBase = m.M
	v194 = m.ExcPending
	if v194 != 0 {
		goto L1
	} else {
		goto L37
	}
L37:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L38:
	;
	v205 = int32(0)
	v212 = v8
	goto L41
L39:
	;
	v241 = v8
	goto L40
L40:
	;
	v254 = v139
	v257 = v8
	v259 = v241
	goto L6
L41:
	;
	v217 = *(*int32)(unsafe.Add(mBase, uint32(l2)+12))
	v221 = int32(*(*int16)(unsafe.Add(mBase, uint32(v217+v205<<(uint(int32(2))%32)))))
	v222 = F_bms_add_member(m, v212, v221)
	mBase = m.M
	v223 = m.ExcPending
	if v223 != 0 {
		goto L1
	} else {
		goto L43
	}
L42:
	;
	v241 = v222
	goto L40
L43:
	;
	v225 = v205 + int32(1)
	v226 = *(*int32)(unsafe.Add(mBase, uint32(l2)+4))
	if v225 < v226 {
		v205 = v225
		v212 = v222
		goto L41
	} else {
		goto L44
	}
L44:
	;
	goto L42
L45:
	;
	v317 = v24 + int32(4)
	if l1 != 0 {
		goto L55
	} else {
		goto L56
	}
L46:
	;
	v275 = v266
	goto L47
L47:
	;
	v288 = v275 - int32(1)
	v292 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l3+v288<<(uint(int32(4))%32))+29)))
	if v292 != 0 {
		goto L49
	} else {
		goto L50
	}
L48:
	;
	goto L45
L49:
	;
	if base.Ui32(int32(1)) < base.Ui32(v275) {
		v275 = v288
		goto L47
	} else {
		goto L53
	}
L50:
	;
	v293 = F_bms_is_member(m, v275, v259)
	mBase = m.M
	v294 = m.ExcPending
	if v294 != 0 {
		goto L1
	} else {
		goto L51
	}
L51:
	;
	if v293 != 0 {
		goto L49
	} else {
		goto L52
	}
L52:
	;
	*(*uint16)(unsafe.Add(mBase, uint32(v21)+52)) = uint16(v275)
	goto L45
L53:
	;
	goto L48
L54:
	;
	F_ExecPushExprSetupSteps(m, v317, v19+int32(-16))
	mBase = m.M
	v326 = m.ExcPending
	if v326 != 0 {
		goto L1
	} else {
		goto L59
	}
L55:
	;
	v320 = F_expr_setup_walker(m, l0, v19+int32(-16))
	mBase = m.M
	v321 = m.ExcPending
	if v321 != 0 {
		goto L1
	} else {
		goto L58
	}
L56:
	;
	goto L57
L57:
	;
	*(*uint16)(unsafe.Add(mBase, uint32(v21)+50)) = uint16(v254)
	goto L54
L58:
	;
	goto L54
L59:
	;
	v328 = v24 + int32(9)
	v330 = v24 + int32(12)
	v331 = int32(0)
	v340 = v331
	v342 = v331
	v344 = v331
	v349 = v8
	goto L64
L60:
	;
	v748 = *(*int32)(unsafe.Add(mBase, uint32(v24)+40))
	if v748 == int32(0) {
		goto L167
	} else {
		goto L168
	}
L61:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v713 = m.ExcPending
	if v713 != 0 {
		goto L1
	} else {
		goto L160
	}
L62:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v692 = m.ExcPending
	if v692 != 0 {
		goto L1
	} else {
		goto L155
	}
L63:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v660 = m.ExcPending
	if v660 != 0 {
		goto L1
	} else {
		goto L147
	}
L64:
	;
	v352 = int32(0)
	if l0 == v352 {
		v362 = v352
		goto L66
	} else {
		goto L67
	}
L65:
	;
	v495 = v371
	v499 = int32(1)
	v502 = v342
	v504 = v344
	v505 = int32(0)
	v509 = v349
	goto L106
L66:
	;
	if v257 != 0 {
		goto L71
	} else {
		goto L72
	}
L67:
	;
	v356 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	if v356 <= v340 {
		v362 = int32(0)
		goto L66
	} else {
		goto L68
	}
L68:
	;
	v358 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	v362 = v358 + v340<<(uint(int32(2))%32)
	goto L66
L69:
	;
	goto L65
L70:
	;
	v379 = int32(*(*int16)(unsafe.Add(mBase, uint32(v368+v340<<(uint(int32(2))%32)))))
	if v379 <= int32(0) {
		goto L61
	} else {
		goto L76
	}
L71:
	;
	v371 = int32(0)
	v372 = *(*int32)(unsafe.Add(mBase, uint32(l3)))
	if v371 < v372 {
		goto L69
	} else {
		goto L75
	}
L72:
	;
	v365 = *(*int32)(unsafe.Add(mBase, uint32(l2+int32(4))))
	if base.B2i32(v362 == int32(0))|base.B2i32(v365 <= v340) != 0 {
		goto L71
	} else {
		goto L73
	}
L73:
	;
	v368 = *(*int32)(unsafe.Add(mBase, uint32(l2)+12))
	if v368 != 0 {
		goto L70
	} else {
		goto L74
	}
L74:
	;
	goto L71
L75:
	;
	v731 = v371
	v734 = v342
	v740 = v344
	v741 = int32(0)
	v745 = v349
	goto L60
L76:
	;
	v382 = *(*int32)(unsafe.Add(mBase, uint32(l3)))
	if v382 < v379 {
		goto L61
	} else {
		goto L77
	}
L77:
	;
	v389 = l3 + v382<<(uint(int32(4))%32) + v379*int32(100)
	v390 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v389)+11)))
	if v390 == int32(1) {
		goto L62
	} else {
		goto L78
	}
L78:
	;
	v393 = *(*int32)(unsafe.Add(mBase, uint32(v362)))
	v394 = *(*int32)(unsafe.Add(mBase, uint32(v393)+4))
	v395 = F_exprType(m, v394)
	mBase = m.M
	v396 = m.ExcPending
	if v396 != 0 {
		goto L1
	} else {
		goto L79
	}
L79:
	;
	v398 = v389 - int32(80)
	v399 = *(*int32)(unsafe.Add(mBase, uint32(v398)+68))
	if v395 != v399 {
		goto L63
	} else {
		goto L80
	}
L80:
	;
	v402 = v379 - int32(1)
	if l1 != 0 {
		goto L82
	} else {
		goto L83
	}
L81:
	;
	v487 = int64(0)
	*(*int64)(unsafe.Add(mBase, uint32(v483)+32)) = v487
	*(*int64)(unsafe.Add(mBase, uint32(v483)+24)) = v487
	v340 = v340 + int32(1)
	v342 = v402
	v344 = v485
	v349 = v486
	goto L64
L82:
	;
	v403 = *(*int32)(unsafe.Add(mBase, uint32(v393)+4))
	F_ExecInitExprRec(m, v403, v317, v330, v328)
	mBase = m.M
	v405 = m.ExcPending
	if v405 != 0 {
		goto L1
	} else {
		goto L85
	}
L83:
	;
	goto L84
L84:
	;
	v446 = *(*int32)(unsafe.Add(mBase, uint32(v24)+40))
	if v446 == int32(0) {
		goto L98
	} else {
		goto L99
	}
L85:
	;
	v406 = *(*int32)(unsafe.Add(mBase, uint32(v24)+40))
	if v406 == int32(0) {
		goto L88
	} else {
		goto L89
	}
L86:
	;
	v428 = *(*int32)(unsafe.Add(mBase, uint32(v24)+36))
	*(*int32)(unsafe.Add(mBase, uint32(v24)+36)) = v428 + int32(1)
	v434 = v427 + v428*int32(40)
	*(*int32)(unsafe.Add(mBase, uint32(v434)+20)) = v344&int32(255) | v349<<(uint(int32(8))%32)
	*(*int32)(unsafe.Add(mBase, uint32(v434)+16)) = v402
	*(*int64)(unsafe.Add(mBase, uint32(v434)+8)) = int64(0)
	*(*int64)(unsafe.Add(mBase, uint32(v434))) = int64(23)
	v483 = v434
	v485 = v344
	v486 = v349
	goto L81
L87:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v24)+20)) = v425
	v427 = v425
	goto L86
L88:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v24)+40)) = int32(16)
	v412 = F_palloc(m, int32(640))
	mBase = m.M
	v413 = m.ExcPending
	if v413 != 0 {
		goto L1
	} else {
		goto L91
	}
L89:
	;
	goto L90
L90:
	;
	v414 = *(*int32)(unsafe.Add(mBase, uint32(v24)+36))
	if v414 != v406 {
		goto L92
	} else {
		goto L93
	}
L91:
	;
	v425 = v412
	goto L87
L92:
	;
	v416 = *(*int32)(unsafe.Add(mBase, uint32(v24)+20))
	v427 = v416
	goto L86
L93:
	;
	goto L94
L94:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v24)+40)) = v406 << (uint(int32(1)) % 32)
	v420 = *(*int32)(unsafe.Add(mBase, uint32(v24)+20))
	v423 = F_repalloc(m, v420, v406*int32(80))
	mBase = m.M
	v424 = m.ExcPending
	if v424 != 0 {
		goto L1
	} else {
		goto L95
	}
L95:
	;
	v425 = v423
	goto L87
L96:
	;
	v470 = *(*int32)(unsafe.Add(mBase, uint32(v24)+36))
	*(*int32)(unsafe.Add(mBase, uint32(v24)+36)) = v470 + int32(1)
	v476 = v467 + v470*int32(40)
	*(*int32)(unsafe.Add(mBase, uint32(v476)+20)) = v340
	*(*int32)(unsafe.Add(mBase, uint32(v476)+16)) = v402
	*(*int64)(unsafe.Add(mBase, uint32(v476)+8)) = int64(0)
	*(*int64)(unsafe.Add(mBase, uint32(v476))) = int64(19)
	v483 = v476
	v485 = v340
	v486 = int32(base.Ui32(v340) >> (uint(int32(8)) % 32))
	goto L81
L97:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v24)+20)) = v465
	v467 = v465
	goto L96
L98:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v24)+40)) = int32(16)
	v452 = F_palloc(m, int32(640))
	mBase = m.M
	v453 = m.ExcPending
	if v453 != 0 {
		goto L1
	} else {
		goto L101
	}
L99:
	;
	goto L100
L100:
	;
	v454 = *(*int32)(unsafe.Add(mBase, uint32(v24)+36))
	if v454 != v446 {
		goto L102
	} else {
		goto L103
	}
L101:
	;
	v465 = v452
	goto L97
L102:
	;
	v456 = *(*int32)(unsafe.Add(mBase, uint32(v24)+20))
	v467 = v456
	goto L96
L103:
	;
	goto L104
L104:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v24)+40)) = v446 << (uint(int32(1)) % 32)
	v460 = *(*int32)(unsafe.Add(mBase, uint32(v24)+20))
	v463 = F_repalloc(m, v460, v446*int32(80))
	mBase = m.M
	v464 = m.ExcPending
	if v464 != 0 {
		goto L1
	} else {
		goto L105
	}
L105:
	;
	v465 = v463
	goto L97
L106:
	;
	v512 = int32(1)
	v513 = v499 - v512
	v517 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l3+v513<<(uint(int32(4))%32))+29)))
	if v517 == v512 {
		goto L111
	} else {
		goto L112
	}
L107:
	;
	v731 = v645
	v734 = v647
	v740 = v650
	v741 = v651
	v745 = v652
	goto L60
L108:
	;
	v654 = v499 + int32(1)
	v655 = *(*int32)(unsafe.Add(mBase, uint32(l3)))
	if v654 <= v655 {
		v495 = v645
		v499 = v654
		v502 = v649
		v504 = v650
		v505 = v651
		v509 = v652
		goto L106
	} else {
		goto L146
	}
L109:
	;
	v626 = *(*int32)(unsafe.Add(mBase, uint32(v24)+36))
	*(*int32)(unsafe.Add(mBase, uint32(v24)+36)) = v626 + int32(1)
	v632 = v620 + v626*int32(40)
	*(*int32)(unsafe.Add(mBase, uint32(v632)+20)) = v621
	*(*int32)(unsafe.Add(mBase, uint32(v632)+16)) = v513
	*(*int32)(unsafe.Add(mBase, uint32(v632)+12)) = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v632)+8)) = v619
	*(*int32)(unsafe.Add(mBase, uint32(v632)+4)) = v623
	*(*int32)(unsafe.Add(mBase, uint32(v632))) = v625
	v640 = *(*int64)(unsafe.Add(mBase, uint32(v21)+32))
	*(*int64)(unsafe.Add(mBase, uint32(v632)+24)) = v640
	v642 = *(*int64)(unsafe.Add(mBase, uint32(v21)+40))
	*(*int64)(unsafe.Add(mBase, uint32(v632)+32)) = v642
	v645 = v619
	v647 = v513
	v649 = v513
	v650 = v622
	v651 = v623
	v652 = v624
	goto L108
L110:
	;
	v619 = v328
	v620 = v615
	v621 = v552
	v622 = v616
	v623 = v330
	v624 = v509
	v625 = int32(23)
	goto L109
L111:
	;
	v520 = *(*int32)(unsafe.Add(mBase, uint32(v24)+40))
	if v520 == int32(0) {
		goto L116
	} else {
		goto L117
	}
L112:
	;
	goto L113
L113:
	;
	v588 = F_bms_is_member(m, v499, v259)
	mBase = m.M
	v589 = m.ExcPending
	if v589 != 0 {
		goto L1
	} else {
		goto L132
	}
L114:
	;
	v542 = *(*int32)(unsafe.Add(mBase, uint32(v24)+36))
	v543 = int32(1)
	*(*int32)(unsafe.Add(mBase, uint32(v24)+36)) = v542 + v543
	v548 = v541 + v542*int32(40)
	v552 = v509<<(uint(int32(8))%32) | v543
	*(*int32)(unsafe.Add(mBase, uint32(v548)+20)) = v552
	*(*int64)(unsafe.Add(mBase, uint32(v548)+12)) = int64(0)
	*(*int32)(unsafe.Add(mBase, uint32(v548)+8)) = v328
	*(*int32)(unsafe.Add(mBase, uint32(v548)+4)) = v330
	*(*int32)(unsafe.Add(mBase, uint32(v548))) = int32(25)
	v560 = *(*int64)(unsafe.Add(mBase, uint32(v21)+40))
	*(*int64)(unsafe.Add(mBase, uint32(v548)+32)) = v560
	v562 = *(*int64)(unsafe.Add(mBase, uint32(v21)+32))
	*(*int64)(unsafe.Add(mBase, uint32(v548)+24)) = v562
	v564 = *(*int32)(unsafe.Add(mBase, uint32(v24)+40))
	if v564 == int32(0) {
		goto L126
	} else {
		goto L127
	}
L115:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v24)+20)) = v539
	v541 = v539
	goto L114
L116:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v24)+40)) = int32(16)
	v526 = F_palloc(m, int32(640))
	mBase = m.M
	v527 = m.ExcPending
	if v527 != 0 {
		goto L1
	} else {
		goto L119
	}
L117:
	;
	goto L118
L118:
	;
	v528 = *(*int32)(unsafe.Add(mBase, uint32(v24)+36))
	if v528 != v520 {
		goto L120
	} else {
		goto L121
	}
L119:
	;
	v539 = v526
	goto L115
L120:
	;
	v530 = *(*int32)(unsafe.Add(mBase, uint32(v24)+20))
	v541 = v530
	goto L114
L121:
	;
	goto L122
L122:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v24)+40)) = v520 << (uint(int32(1)) % 32)
	v534 = *(*int32)(unsafe.Add(mBase, uint32(v24)+20))
	v537 = F_repalloc(m, v534, v520*int32(80))
	mBase = m.M
	v538 = m.ExcPending
	if v538 != 0 {
		goto L1
	} else {
		goto L123
	}
L123:
	;
	v539 = v537
	goto L115
L124:
	;
	v578 = int32(1)
	*(*int32)(unsafe.Add(mBase, uint32(v24)+40)) = v564 << (uint(v578) % 32)
	v582 = *(*int32)(unsafe.Add(mBase, uint32(v24)+20))
	v585 = F_repalloc(m, v582, v564*int32(80))
	mBase = m.M
	v586 = m.ExcPending
	if v586 != 0 {
		goto L1
	} else {
		goto L131
	}
L125:
	;
	v615 = v576
	v616 = int32(1)
	goto L110
L126:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v24)+40)) = int32(16)
	v570 = F_palloc(m, int32(640))
	mBase = m.M
	v571 = m.ExcPending
	if v571 != 0 {
		goto L1
	} else {
		goto L129
	}
L127:
	;
	goto L128
L128:
	;
	v573 = *(*int32)(unsafe.Add(mBase, uint32(v24)+36))
	if v564 == v573 {
		goto L124
	} else {
		goto L130
	}
L129:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v24)+20)) = v570
	v576 = v570
	goto L125
L130:
	;
	v575 = *(*int32)(unsafe.Add(mBase, uint32(v24)+20))
	v576 = v575
	goto L125
L131:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v24)+20)) = v585
	v615 = v585
	v616 = v578
	goto L110
L132:
	;
	if v588 != 0 {
		goto L133
	} else {
		goto L134
	}
L133:
	;
	v645 = v495
	v647 = v502
	v649 = v502
	v650 = v504
	v651 = v505
	v652 = v509
	goto L108
L134:
	;
	goto L135
L135:
	;
	v592 = *(*int32)(unsafe.Add(mBase, uint32(v24)+40))
	if v592 == int32(0) {
		goto L138
	} else {
		goto L139
	}
L136:
	;
	v619 = v495
	v620 = v613
	v621 = v513
	v622 = v513
	v623 = v505
	v624 = int32(base.Ui32(v513) >> (uint(int32(8)) % 32))
	v625 = int32(20)
	goto L109
L137:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v24)+20)) = v611
	v613 = v611
	goto L136
L138:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v24)+40)) = int32(16)
	v598 = F_palloc(m, int32(640))
	mBase = m.M
	v599 = m.ExcPending
	if v599 != 0 {
		goto L1
	} else {
		goto L141
	}
L139:
	;
	goto L140
L140:
	;
	v600 = *(*int32)(unsafe.Add(mBase, uint32(v24)+36))
	if v600 != v592 {
		goto L142
	} else {
		goto L143
	}
L141:
	;
	v611 = v598
	goto L137
L142:
	;
	v602 = *(*int32)(unsafe.Add(mBase, uint32(v24)+20))
	v613 = v602
	goto L136
L143:
	;
	goto L144
L144:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v24)+40)) = v592 << (uint(int32(1)) % 32)
	v606 = *(*int32)(unsafe.Add(mBase, uint32(v24)+20))
	v609 = F_repalloc(m, v606, v592*int32(80))
	mBase = m.M
	v610 = m.ExcPending
	if v610 != 0 {
		goto L1
	} else {
		goto L145
	}
L145:
	;
	v611 = v609
	goto L137
L146:
	;
	goto L107
L147:
	;
	F_errcode(m, int32(67141764))
	mBase = m.M
	v663 = m.ExcPending
	if v663 != 0 {
		goto L1
	} else {
		goto L148
	}
L148:
	;
	F_errmsg(m, int32(_a_F_ExecBuildUpdateProjection_4), int32(0))
	mBase = m.M
	v667 = m.ExcPending
	if v667 != 0 {
		goto L1
	} else {
		goto L149
	}
L149:
	;
	v668 = *(*int32)(unsafe.Add(mBase, uint32(v398)+68))
	v669 = F_format_type_be(m, v668)
	mBase = m.M
	v670 = m.ExcPending
	if v670 != 0 {
		goto L1
	} else {
		goto L150
	}
L150:
	;
	v671 = *(*int32)(unsafe.Add(mBase, uint32(v393)+4))
	v672 = F_exprType(m, v671)
	mBase = m.M
	v673 = m.ExcPending
	if v673 != 0 {
		goto L1
	} else {
		goto L151
	}
L151:
	;
	v674 = F_format_type_be(m, v672)
	mBase = m.M
	v675 = m.ExcPending
	if v675 != 0 {
		goto L1
	} else {
		goto L152
	}
L152:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v21)+24)) = v674
	*(*int32)(unsafe.Add(mBase, uint32(v21)+20)) = v379
	*(*int32)(unsafe.Add(mBase, uint32(v21)+16)) = v669
	F_errdetail(m, int32(_a_F_ExecBuildUpdateProjection_5), v19+int32(-48))
	mBase = m.M
	v683 = m.ExcPending
	if v683 != 0 {
		goto L1
	} else {
		goto L153
	}
L153:
	;
	F_errfinish(m, int32(_a_F_ExecBuildUpdateProjection_1), int32(684), int32(_a_F_ExecBuildUpdateProjection_2))
	mBase = m.M
	v688 = m.ExcPending
	if v688 != 0 {
		goto L1
	} else {
		goto L154
	}
L154:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L155:
	;
	F_errcode(m, int32(67141764))
	mBase = m.M
	v695 = m.ExcPending
	if v695 != 0 {
		goto L1
	} else {
		goto L156
	}
L156:
	;
	F_errmsg(m, int32(_a_F_ExecBuildUpdateProjection_4), int32(0))
	mBase = m.M
	v699 = m.ExcPending
	if v699 != 0 {
		goto L1
	} else {
		goto L157
	}
L157:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v21))) = v379
	F_errdetail(m, int32(_a_F_ExecBuildUpdateProjection_6), v21)
	mBase = m.M
	v703 = m.ExcPending
	if v703 != 0 {
		goto L1
	} else {
		goto L158
	}
L158:
	;
	F_errfinish(m, int32(_a_F_ExecBuildUpdateProjection_1), int32(676), int32(_a_F_ExecBuildUpdateProjection_2))
	mBase = m.M
	v708 = m.ExcPending
	if v708 != 0 {
		goto L1
	} else {
		goto L159
	}
L159:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L160:
	;
	F_errcode(m, int32(67141764))
	mBase = m.M
	v716 = m.ExcPending
	if v716 != 0 {
		goto L1
	} else {
		goto L161
	}
L161:
	;
	F_errmsg(m, int32(_a_F_ExecBuildUpdateProjection_4), int32(0))
	mBase = m.M
	v720 = m.ExcPending
	if v720 != 0 {
		goto L1
	} else {
		goto L162
	}
L162:
	;
	F_errdetail(m, int32(_a_F_ExecBuildUpdateProjection_7), int32(0))
	mBase = m.M
	v724 = m.ExcPending
	if v724 != 0 {
		goto L1
	} else {
		goto L163
	}
L163:
	;
	F_errfinish(m, int32(_a_F_ExecBuildUpdateProjection_1), int32(668), int32(_a_F_ExecBuildUpdateProjection_2))
	mBase = m.M
	v729 = m.ExcPending
	if v729 != 0 {
		goto L1
	} else {
		goto L164
	}
L164:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L165:
	;
	v770 = *(*int32)(unsafe.Add(mBase, uint32(v24)+36))
	v771 = int32(1)
	*(*int32)(unsafe.Add(mBase, uint32(v24)+36)) = v770 + v771
	v776 = v769 + v770*int32(40)
	*(*int32)(unsafe.Add(mBase, uint32(v776)+20)) = v740&int32(255) | v745<<(uint(int32(8))%32)
	*(*int32)(unsafe.Add(mBase, uint32(v776)+16)) = v734
	*(*int32)(unsafe.Add(mBase, uint32(v776)+12)) = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v776)+8)) = v731
	*(*int32)(unsafe.Add(mBase, uint32(v776)+4)) = v741
	*(*int32)(unsafe.Add(mBase, uint32(v776))) = v771
	v790 = *(*int64)(unsafe.Add(mBase, uint32(v21)+32))
	*(*int64)(unsafe.Add(mBase, uint32(v776)+24)) = v790
	v792 = *(*int64)(unsafe.Add(mBase, uint32(v21)+40))
	*(*int64)(unsafe.Add(mBase, uint32(v776)+32)) = v792
	v794 = F_jit_compile_expr(m, v317)
	mBase = m.M
	v795 = m.ExcPending
	if v795 != 0 {
		goto L1
	} else {
		goto L175
	}
L166:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v24)+20)) = v767
	v769 = v767
	goto L165
L167:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v24)+40)) = int32(16)
	v754 = F_palloc(m, int32(640))
	mBase = m.M
	v755 = m.ExcPending
	if v755 != 0 {
		goto L1
	} else {
		goto L170
	}
L168:
	;
	goto L169
L169:
	;
	v756 = *(*int32)(unsafe.Add(mBase, uint32(v24)+36))
	if v756 != v748 {
		goto L171
	} else {
		goto L172
	}
L170:
	;
	v767 = v754
	goto L166
L171:
	;
	v758 = *(*int32)(unsafe.Add(mBase, uint32(v24)+20))
	v769 = v758
	goto L165
L172:
	;
	goto L173
L173:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v24)+40)) = v748 << (uint(int32(1)) % 32)
	v762 = *(*int32)(unsafe.Add(mBase, uint32(v24)+20))
	v765 = F_repalloc(m, v762, v748*int32(80))
	mBase = m.M
	v766 = m.ExcPending
	if v766 != 0 {
		goto L1
	} else {
		goto L174
	}
L174:
	;
	v767 = v765
	goto L166
L175:
	;
	if v794 == int32(0) {
		goto L176
	} else {
		goto L177
	}
L176:
	;
	F_ExecReadyInterpretedExpr(m, v317)
	mBase = m.M
	v799 = m.ExcPending
	if v799 != 0 {
		goto L1
	} else {
		goto L179
	}
L177:
	;
	goto L178
L178:
	;
	m.G0 = v21 - int32(-64)
	return v24
L179:
	;
	goto L178
}
func F_create_set_projection_path(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) int32 {
	mBase := m.M
	_ = mBase
	var v5 int32
	_ = v5
	var v11 int32
	_ = v11
	var v14 int32
	_ = v14
	var v15 int32
	_ = v15
	var v23 int32
	_ = v23
	var v26 int32
	_ = v26
	var v29 int32
	_ = v29
	var v30 int32
	_ = v30
	var v31 int32
	_ = v31
	var v32 int32
	_ = v32
	var v34 int32
	_ = v34
	var v36 int32
	_ = v36
	var v39 int32
	_ = v39
	var v43 float64
	_ = v43
	var v44 int32
	_ = v44
	var v49 int32
	_ = v49
	var v54 float64
	_ = v54
	var v57 int32
	_ = v57
	var v61 int32
	_ = v61
	var v62 float64
	_ = v62
	var v63 int32
	_ = v63
	var v65 float64
	_ = v65
	var v67 int32
	_ = v67
	var v68 int32
	_ = v68
	var v76 float64
	_ = v76
	var v79 int32
	_ = v79
	var v81 float64
	_ = v81
	var v82 float64
	_ = v82
	var v84 float64
	_ = v84
	var v85 float64
	_ = v85
	var v89 float64
	_ = v89
	var v90 float64
	_ = v90
	var v92 float64
	_ = v92
	var v94 float64
	_ = v94
	var v95 float64
	_ = v95
	v5 = int32(0)
	v11 = F_palloc0(m, int32(80))
	mBase = m.M
	v14 = m.ExcPending
	if v14 != 0 {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	return int32(0)
L2:
	;
	v15 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v11)+20)) = uint8(v15)
	*(*int32)(unsafe.Add(mBase, uint32(v11)+16)) = v15
	*(*int32)(unsafe.Add(mBase, uint32(v11)+12)) = l3
	*(*int32)(unsafe.Add(mBase, uint32(v11)+8)) = l1
	*(*int64)(unsafe.Add(mBase, uint32(v11))) = int64(1425929142574)
	v23 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l1)+26)))
	if v23 != int32(1) {
		v32 = v5
		goto L3
	} else {
		goto L4
	}
L3:
	;
	*(*uint8)(unsafe.Add(mBase, uint32(v11)+21)) = uint8(v32)
	v34 = *(*int32)(unsafe.Add(mBase, uint32(l2)+24))
	*(*int32)(unsafe.Add(mBase, uint32(v11)+24)) = v34
	v36 = *(*int32)(unsafe.Add(mBase, uint32(l2)+64))
	*(*int32)(unsafe.Add(mBase, uint32(v11)+72)) = l2
	*(*int32)(unsafe.Add(mBase, uint32(v11)+64)) = v36
	v39 = *(*int32)(unsafe.Add(mBase, uint32(l3)+4))
	if v39 == int32(0) {
		goto L8
	} else {
		goto L9
	}
L4:
	;
	v26 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l2)+21)))
	if v26 != int32(1) {
		v32 = v5
		goto L3
	} else {
		goto L5
	}
L5:
	;
	v29 = *(*int32)(unsafe.Add(mBase, uint32(l3)+4))
	v30 = F_is_parallel_safe(m, l0, v29)
	mBase = m.M
	v31 = m.ExcPending
	if v31 != 0 {
		goto L1
	} else {
		goto L6
	}
L6:
	;
	v32 = v30
	goto L3
L7:
	;
	v79 = *(*int32)(unsafe.Add(mBase, uint32(l2)+40))
	*(*int32)(unsafe.Add(mBase, uint32(v11)+40)) = v79
	v81 = *(*float64)(unsafe.Add(mBase, uint32(l2)+32))
	v82 = base.F64_mul(v76, v81)
	*(*float64)(unsafe.Add(mBase, uint32(v11)+32)) = v82
	v84 = *(*float64)(unsafe.Add(mBase, uint32(l2)+48))
	v85 = *(*float64)(unsafe.Add(mBase, uint32(l3)+16))
	*(*float64)(unsafe.Add(mBase, uint32(v11)+48)) = base.F64_add(v84, v85)
	v89 = *(*float64)(unsafe.Add(mBase, _c_F_create_set_projection_path[0]))
	v90 = *(*float64)(unsafe.Add(mBase, uint32(l3)+24))
	v92 = *(*float64)(unsafe.Add(mBase, uint32(l2)+32))
	v94 = *(*float64)(unsafe.Add(mBase, uint32(l2)+56))
	v95 = *(*float64)(unsafe.Add(mBase, uint32(l3)+16))
	*(*float64)(unsafe.Add(mBase, uint32(v11)+56)) = base.F64_add(base.F64_add(base.F64_mul(base.F64_add(v89, v90), v92), base.F64_add(v94, v95)), base.F64_mul(base.F64_mul(v89, base.F64_sub(v82, v92)), float64(0.5)))
	return v11
L8:
	;
	v76 = float64(1)
	goto L7
L9:
	;
	goto L10
L10:
	;
	v43 = float64(1)
	v44 = *(*int32)(unsafe.Add(mBase, uint32(v39)+4))
	if v44 <= int32(0) {
		v76 = v43
		goto L7
	} else {
		goto L11
	}
L11:
	;
	v49 = int32(0)
	v54 = v43
	goto L12
L12:
	;
	v57 = *(*int32)(unsafe.Add(mBase, uint32(v39)+12))
	v61 = *(*int32)(unsafe.Add(mBase, uint32(v57+v49<<(uint(int32(2))%32))))
	v62 = F_expression_returns_set_rows(m, l0, v61)
	mBase = m.M
	v63 = m.ExcPending
	if v63 != 0 {
		goto L1
	} else {
		goto L14
	}
L13:
	;
	v76 = v65
	goto L7
L14:
	;
	if base.F64_lt(v54, v62) != 0 {
		goto L15
	} else {
		goto L16
	}
L15:
	;
	v65 = v62
	goto L17
L16:
	;
	v65 = v54
	goto L17
L17:
	;
	v67 = v49 + int32(1)
	v68 = *(*int32)(unsafe.Add(mBase, uint32(v39)+4))
	if v67 < v68 {
		v49 = v67
		v54 = v65
		goto L12
	} else {
		goto L18
	}
L18:
	;
	goto L13
}
func F_is_projection_capable_path(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v2 int32
	_ = v2
	var v3 int32
	_ = v3
	var v6 int32
	_ = v6
	var v12 int32
	_ = v12
	var v15 int32
	_ = v15
	var v20 int32
	_ = v20
	v2 = int32(0)
	v3 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	switch v3 - int32(332) {
	case 0, 1, 3, 4, 28, 29, 30, 31, 35, 38, 39, 40, 41:
		v20 = v2
		return v20
	case 2:
		v12 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
		if v12 != int32(290) {
			v20 = v2
			return v20
		} else {
			v15 = *(*int32)(unsafe.Add(mBase, uint32(l0)+72))
			return base.B2i32(v15 == int32(0))
		}
	default:
		v20 = int32(1)
		return v20
	case 23:
		v6 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+72)))
		return int32(base.Ui32(v6&int32(4)) >> (uint(int32(2)) % 32))
	}
}
