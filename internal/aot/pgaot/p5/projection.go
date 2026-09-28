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
	var v67 int32
	_ = v67
	var v79 int32
	_ = v79
	var v80 int32
	_ = v80
	var v87 int32
	_ = v87
	var v89 int32
	_ = v89
	var v100 int32
	_ = v100
	var v140 int32
	_ = v140
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
	var v211 int32
	_ = v211
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
	var v240 int32
	_ = v240
	var v255 int32
	_ = v255
	var v258 int32
	_ = v258
	var v261 int32
	_ = v261
	var v266 int32
	_ = v266
	var v275 int32
	_ = v275
	var v288 int32
	_ = v288
	var v292 int32
	_ = v292
	var v295 int32
	_ = v295
	var v296 int32
	_ = v296
	var v319 int32
	_ = v319
	var v322 int32
	_ = v322
	var v323 int32
	_ = v323
	var v328 int32
	_ = v328
	var v330 int32
	_ = v330
	var v332 int32
	_ = v332
	var v333 int32
	_ = v333
	var v339 int32
	_ = v339
	var v344 int32
	_ = v344
	var v349 int32
	_ = v349
	var v353 int32
	_ = v353
	var v357 int32
	_ = v357
	var v359 int32
	_ = v359
	var v363 int32
	_ = v363
	var v366 int32
	_ = v366
	var v369 int32
	_ = v369
	var v372 int32
	_ = v372
	var v373 int32
	_ = v373
	var v376 int32
	_ = v376
	var v381 int32
	_ = v381
	var v384 int32
	_ = v384
	var v391 int32
	_ = v391
	var v392 int32
	_ = v392
	var v395 int32
	_ = v395
	var v396 int32
	_ = v396
	var v397 int32
	_ = v397
	var v398 int32
	_ = v398
	var v400 int32
	_ = v400
	var v401 int32
	_ = v401
	var v404 int32
	_ = v404
	var v405 int32
	_ = v405
	var v407 int32
	_ = v407
	var v408 int32
	_ = v408
	var v411 int32
	_ = v411
	var v415 int32
	_ = v415
	var v416 int32
	_ = v416
	var v417 int32
	_ = v417
	var v419 int32
	_ = v419
	var v424 int32
	_ = v424
	var v427 int32
	_ = v427
	var v428 int32
	_ = v428
	var v429 int32
	_ = v429
	var v432 int32
	_ = v432
	var v436 int32
	_ = v436
	var v437 int32
	_ = v437
	var v438 int32
	_ = v438
	var v440 int32
	_ = v440
	var v444 int32
	_ = v444
	var v447 int32
	_ = v447
	var v448 int32
	_ = v448
	var v449 int32
	_ = v449
	var v451 int32
	_ = v451
	var v453 int32
	_ = v453
	var v455 int32
	_ = v455
	var v456 int32
	_ = v456
	var v457 int32
	_ = v457
	var v459 int32
	_ = v459
	var v464 int32
	_ = v464
	var v469 int32
	_ = v469
	var v474 int32
	_ = v474
	var v478 int32
	_ = v478
	var v481 int32
	_ = v481
	var v484 int32
	_ = v484
	var v488 int32
	_ = v488
	var v489 int32
	_ = v489
	var v490 int32
	_ = v490
	var v492 int32
	_ = v492
	var v496 int32
	_ = v496
	var v499 int32
	_ = v499
	var v500 int32
	_ = v500
	var v501 int32
	_ = v501
	var v503 int32
	_ = v503
	var v504 int32
	_ = v504
	var v505 int32
	_ = v505
	var v510 int32
	_ = v510
	var v515 int32
	_ = v515
	var v521 int64
	_ = v521
	var v523 int64
	_ = v523
	var v525 int32
	_ = v525
	var v528 int32
	_ = v528
	var v532 int32
	_ = v532
	var v533 int32
	_ = v533
	var v534 int32
	_ = v534
	var v536 int32
	_ = v536
	var v540 int32
	_ = v540
	var v543 int32
	_ = v543
	var v544 int32
	_ = v544
	var v545 int32
	_ = v545
	var v547 int32
	_ = v547
	var v548 int32
	_ = v548
	var v549 int32
	_ = v549
	var v555 int32
	_ = v555
	var v558 int32
	_ = v558
	var v565 int32
	_ = v565
	var v566 int32
	_ = v566
	var v567 int32
	_ = v567
	var v570 int32
	_ = v570
	var v574 int32
	_ = v574
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
	var v587 int32
	_ = v587
	var v589 int32
	_ = v589
	var v590 int32
	_ = v590
	var v596 int32
	_ = v596
	var v607 int32
	_ = v607
	var v608 int32
	_ = v608
	var v609 int32
	_ = v609
	var v610 int32
	_ = v610
	var v612 int32
	_ = v612
	var v614 int64
	_ = v614
	var v616 int64
	_ = v616
	var v618 int32
	_ = v618
	var v619 int32
	_ = v619
	var v620 int32
	_ = v620
	var v622 int32
	_ = v622
	var v623 int32
	_ = v623
	var v625 int32
	_ = v625
	var v626 int32
	_ = v626
	var v632 int32
	_ = v632
	var v635 int32
	_ = v635
	var v639 int32
	_ = v639
	var v642 int32
	_ = v642
	var v643 int32
	_ = v643
	var v648 int32
	_ = v648
	var v652 int32
	_ = v652
	var v655 int32
	_ = v655
	var v659 int32
	_ = v659
	var v662 int32
	_ = v662
	var v663 int32
	_ = v663
	var v668 int32
	_ = v668
	var v672 int32
	_ = v672
	var v675 int32
	_ = v675
	var v679 int32
	_ = v679
	var v680 int32
	_ = v680
	var v681 int32
	_ = v681
	var v682 int32
	_ = v682
	var v683 int32
	_ = v683
	var v684 int32
	_ = v684
	var v685 int32
	_ = v685
	var v686 int32
	_ = v686
	var v687 int32
	_ = v687
	var v694 int32
	_ = v694
	var v695 int32
	_ = v695
	var v700 int32
	_ = v700
	var v701 int32
	_ = v701
	var v704 int32
	_ = v704
	var v705 int32
	_ = v705
	var v706 int32
	_ = v706
	var v707 int32
	_ = v707
	var v708 int32
	_ = v708
	var v713 int32
	_ = v713
	var v714 int64
	_ = v714
	var v731 int32
	_ = v731
	var v732 int32
	_ = v732
	var v733 int32
	_ = v733
	var v740 int32
	_ = v740
	var v745 int32
	_ = v745
	var v749 int32
	_ = v749
	var v752 int32
	_ = v752
	var v756 int32
	_ = v756
	var v757 int32
	_ = v757
	var v758 int32
	_ = v758
	var v760 int32
	_ = v760
	var v764 int32
	_ = v764
	var v767 int32
	_ = v767
	var v768 int32
	_ = v768
	var v769 int32
	_ = v769
	var v771 int32
	_ = v771
	var v772 int32
	_ = v772
	var v773 int32
	_ = v773
	var v778 int32
	_ = v778
	var v792 int64
	_ = v792
	var v794 int64
	_ = v794
	var v796 int32
	_ = v796
	var v797 int32
	_ = v797
	var v801 int32
	_ = v801
	v8 = int32(0)
	v19 = m.G0
	v21 = v19 + int32(-64)
	m.G0 = v21
	v24 = F_palloc0(m, int32(88))
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
	*(*int32)(unsafe.Add(mBase, uint32(v24))) = int32(390)
	v30 = int64(0)
	*(*int64)(unsafe.Add(mBase, uint32(v21)+56)) = v30
	*(*int64)(unsafe.Add(mBase, uint32(v21)+48)) = v30
	*(*int64)(unsafe.Add(mBase, uint32(v21)+39)) = v30
	*(*int64)(unsafe.Add(mBase, uint32(v21)+32)) = v30
	v38 = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v24)+56)) = v38
	*(*int32)(unsafe.Add(mBase, uint32(v24)+52)) = l6
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
	*(*int32)(unsafe.Add(mBase, uint32(v24)+36)) = v42
	*(*int32)(unsafe.Add(mBase, uint32(v24)+8)) = int32(386)
	*(*int32)(unsafe.Add(mBase, uint32(v24)+80)) = l4
	*(*int32)(unsafe.Add(mBase, uint32(v24)+24)) = l5
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
	if v140 == v149 {
		goto L7
	} else {
		goto L31
	}
L11:
	;
	v255 = int32(0)
	v258 = v8
	v261 = int32(1)
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
		v140 = v8
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
	v100 = v8
	goto L17
L17:
	;
	if l2 != 0 {
		v140 = v100
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
	v67 = v8
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
	v100 = v87
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
	v87 = v67
	goto L25
L25:
	;
	v89 = v64 + int32(1)
	if v89 != v54 {
		v63 = v80
		v64 = v89
		v67 = v87
		goto L21
	} else {
		goto L27
	}
L26:
	;
	v87 = v67 + int32(1)
	goto L25
L27:
	;
	goto L22
L28:
	;
	if v100 == int32(0) {
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
	v211 = v8
	goto L41
L39:
	;
	v240 = v8
	goto L40
L40:
	;
	v255 = v140
	v258 = v240
	v261 = v8
	goto L6
L41:
	;
	v217 = *(*int32)(unsafe.Add(mBase, uint32(l2)+12))
	v221 = int32(*(*int16)(unsafe.Add(mBase, uint32(v217+v205<<(uint(int32(2))%32)))))
	v222 = F_bms_add_member(m, v211, v221)
	mBase = m.M
	v223 = m.ExcPending
	if v223 != 0 {
		goto L1
	} else {
		goto L43
	}
L42:
	;
	v240 = v222
	goto L40
L43:
	;
	v225 = v205 + int32(1)
	v226 = *(*int32)(unsafe.Add(mBase, uint32(l2)+4))
	if v225 < v226 {
		v205 = v225
		v211 = v222
		goto L41
	} else {
		goto L44
	}
L44:
	;
	goto L42
L45:
	;
	v319 = v24 + int32(8)
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
	v292 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l3+v288<<(uint(int32(3))%32))+34)))
	if v292&int32(4) != 0 {
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
	v295 = F_bms_is_member(m, v275, v258)
	mBase = m.M
	v296 = m.ExcPending
	if v296 != 0 {
		goto L1
	} else {
		goto L51
	}
L51:
	;
	if v295 != 0 {
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
	F_ExecPushExprSetupSteps(m, v319, v19+int32(-16))
	mBase = m.M
	v328 = m.ExcPending
	if v328 != 0 {
		goto L1
	} else {
		goto L59
	}
L55:
	;
	v322 = F_expr_setup_walker(m, l0, v19+int32(-16))
	mBase = m.M
	v323 = m.ExcPending
	if v323 != 0 {
		goto L1
	} else {
		goto L58
	}
L56:
	;
	goto L57
L57:
	;
	*(*uint16)(unsafe.Add(mBase, uint32(v21)+50)) = uint16(v255)
	goto L54
L58:
	;
	goto L54
L59:
	;
	v330 = v24 + int32(13)
	v332 = v24 + int32(16)
	v333 = int32(0)
	v339 = v333
	v344 = v333
	v349 = v8
	goto L60
L60:
	;
	v353 = int32(0)
	if l0 == v353 {
		v363 = v353
		goto L62
	} else {
		goto L63
	}
L61:
	;
	v749 = *(*int32)(unsafe.Add(mBase, uint32(v24)+48))
	if v749 == int32(0) {
		goto L165
	} else {
		goto L166
	}
L62:
	;
	if v261 != 0 {
		goto L73
	} else {
		goto L74
	}
L63:
	;
	v357 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	if v357 <= v339 {
		v363 = int32(0)
		goto L62
	} else {
		goto L64
	}
L64:
	;
	v359 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	v363 = v359 + v339<<(uint(int32(2))%32)
	goto L62
L65:
	;
	goto L61
L66:
	;
	v707 = *(*int32)(unsafe.Add(mBase, uint32(v24)+44))
	v708 = int32(1)
	*(*int32)(unsafe.Add(mBase, uint32(v24)+44)) = v707 + v708
	v713 = v704 + v707*int32(40)
	v714 = int64(0)
	*(*int64)(unsafe.Add(mBase, uint32(v713)+24)) = v714
	*(*int64)(unsafe.Add(mBase, uint32(v713)+16)) = base.I64_extend_i32_u(v404) | base.I64_extend_i32_u(v705)<<(uint(int64(32))%64)
	*(*int32)(unsafe.Add(mBase, uint32(v713)+12)) = int32(0)
	*(*int64)(unsafe.Add(mBase, uint32(v713)+4)) = v714
	*(*int32)(unsafe.Add(mBase, uint32(v713))) = v706
	*(*int64)(unsafe.Add(mBase, uint32(v713)+32)) = v714
	v339 = v339 + v708
	v344 = v404
	v349 = v705
	goto L60
L67:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v24)+28)) = v701
	v704 = v701
	v705 = v349
	v706 = int32(23)
	goto L66
L68:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v672 = m.ExcPending
	if v672 != 0 {
		goto L1
	} else {
		goto L155
	}
L69:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v652 = m.ExcPending
	if v652 != 0 {
		goto L1
	} else {
		goto L150
	}
L70:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v632 = m.ExcPending
	if v632 != 0 {
		goto L1
	} else {
		goto L145
	}
L71:
	;
	v453 = int32(0)
	v455 = v453
	v456 = v453
	v457 = v372
	v459 = int32(1)
	v464 = v344
	v469 = v349
	goto L105
L72:
	;
	v381 = int32(*(*int16)(unsafe.Add(mBase, uint32(v369+v339<<(uint(int32(2))%32)))))
	if v381 <= int32(0) {
		goto L70
	} else {
		goto L78
	}
L73:
	;
	v372 = int32(0)
	v373 = *(*int32)(unsafe.Add(mBase, uint32(l3)))
	if v372 < v373 {
		goto L71
	} else {
		goto L77
	}
L74:
	;
	v366 = *(*int32)(unsafe.Add(mBase, uint32(l2+int32(4))))
	if base.B2i32(v363 == int32(0))|base.B2i32(v366 <= v339) != 0 {
		goto L73
	} else {
		goto L75
	}
L75:
	;
	v369 = *(*int32)(unsafe.Add(mBase, uint32(l2)+12))
	if v369 != 0 {
		goto L72
	} else {
		goto L76
	}
L76:
	;
	goto L73
L77:
	;
	v376 = int32(0)
	v731 = v376
	v732 = v376
	v733 = v372
	v740 = v344
	v745 = v349
	goto L65
L78:
	;
	v384 = *(*int32)(unsafe.Add(mBase, uint32(l3)))
	if v384 < v381 {
		goto L70
	} else {
		goto L79
	}
L79:
	;
	v391 = l3 + v384<<(uint(int32(3))%32) + v381*int32(100)
	v392 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v391)+19)))
	if v392 == int32(1) {
		goto L69
	} else {
		goto L80
	}
L80:
	;
	v395 = *(*int32)(unsafe.Add(mBase, uint32(v363)))
	v396 = *(*int32)(unsafe.Add(mBase, uint32(v395)+4))
	v397 = F_exprType(m, v396)
	mBase = m.M
	v398 = m.ExcPending
	if v398 != 0 {
		goto L1
	} else {
		goto L81
	}
L81:
	;
	v400 = v391 - int32(72)
	v401 = *(*int32)(unsafe.Add(mBase, uint32(v400)+68))
	if v397 != v401 {
		goto L68
	} else {
		goto L82
	}
L82:
	;
	v404 = v381 - int32(1)
	if l1 != 0 {
		goto L83
	} else {
		goto L84
	}
L83:
	;
	v405 = *(*int32)(unsafe.Add(mBase, uint32(v395)+4))
	F_ExecInitExprRec(m, v405, v319, v332, v330)
	mBase = m.M
	v407 = m.ExcPending
	if v407 != 0 {
		goto L1
	} else {
		goto L86
	}
L84:
	;
	goto L85
L85:
	;
	v429 = *(*int32)(unsafe.Add(mBase, uint32(v24)+48))
	if v429 == int32(0) {
		goto L97
	} else {
		goto L98
	}
L86:
	;
	v408 = *(*int32)(unsafe.Add(mBase, uint32(v24)+48))
	if v408 == int32(0) {
		goto L87
	} else {
		goto L88
	}
L87:
	;
	v411 = int32(16)
	*(*int32)(unsafe.Add(mBase, uint32(v24)+48)) = v411
	v415 = F_palloc_mul(m, int32(40), v411)
	mBase = m.M
	v416 = m.ExcPending
	if v416 != 0 {
		goto L1
	} else {
		goto L90
	}
L88:
	;
	goto L89
L89:
	;
	v417 = *(*int32)(unsafe.Add(mBase, uint32(v24)+44))
	if v417 != v408 {
		goto L91
	} else {
		goto L92
	}
L90:
	;
	v701 = v415
	goto L67
L91:
	;
	v419 = *(*int32)(unsafe.Add(mBase, uint32(v24)+28))
	v704 = v419
	v705 = v349
	v706 = int32(23)
	goto L66
L92:
	;
	goto L93
L93:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v24)+48)) = v408 << (uint(int32(1)) % 32)
	v424 = *(*int32)(unsafe.Add(mBase, uint32(v24)+28))
	v427 = F_repalloc(m, v424, v408*int32(80))
	mBase = m.M
	v428 = m.ExcPending
	if v428 != 0 {
		goto L1
	} else {
		goto L94
	}
L94:
	;
	v701 = v427
	goto L67
L95:
	;
	v704 = v451
	v705 = v339
	v706 = int32(19)
	goto L66
L96:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v24)+28)) = v449
	v451 = v449
	goto L95
L97:
	;
	v432 = int32(16)
	*(*int32)(unsafe.Add(mBase, uint32(v24)+48)) = v432
	v436 = F_palloc_mul(m, int32(40), v432)
	mBase = m.M
	v437 = m.ExcPending
	if v437 != 0 {
		goto L1
	} else {
		goto L100
	}
L98:
	;
	goto L99
L99:
	;
	v438 = *(*int32)(unsafe.Add(mBase, uint32(v24)+44))
	if v438 != v429 {
		goto L101
	} else {
		goto L102
	}
L100:
	;
	v449 = v436
	goto L96
L101:
	;
	v440 = *(*int32)(unsafe.Add(mBase, uint32(v24)+28))
	v451 = v440
	goto L95
L102:
	;
	goto L103
L103:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v24)+48)) = v429 << (uint(int32(1)) % 32)
	v444 = *(*int32)(unsafe.Add(mBase, uint32(v24)+28))
	v447 = F_repalloc(m, v444, v429*int32(80))
	mBase = m.M
	v448 = m.ExcPending
	if v448 != 0 {
		goto L1
	} else {
		goto L104
	}
L104:
	;
	v449 = v447
	goto L96
L105:
	;
	v474 = v459 - int32(1)
	v478 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l3+v474<<(uint(int32(3))%32))+34)))
	if v478&int32(4) != 0 {
		goto L109
	} else {
		goto L110
	}
L106:
	;
	v731 = v618
	v732 = v619
	v733 = v620
	v740 = v622
	v745 = v623
	goto L65
L107:
	;
	v625 = v459 + int32(1)
	v626 = *(*int32)(unsafe.Add(mBase, uint32(l3)))
	if v625 <= v626 {
		v455 = v618
		v456 = v619
		v457 = v620
		v459 = v625
		v464 = v622
		v469 = v623
		goto L105
	} else {
		goto L144
	}
L108:
	;
	*(*uint8)(unsafe.Add(mBase, uint32(v610)+24)) = uint8(v609)
	v614 = *(*int64)(unsafe.Add(mBase, uint32(v21)+32))
	*(*int64)(unsafe.Add(mBase, uint32(v610)+25)) = v614
	v616 = *(*int64)(unsafe.Add(mBase, uint32(v21)+39))
	*(*int64)(unsafe.Add(mBase, uint32(v610)+32)) = v616
	v618 = v607
	v619 = v608
	v620 = v609
	v622 = v474
	v623 = v612
	goto L107
L109:
	;
	v481 = *(*int32)(unsafe.Add(mBase, uint32(v24)+48))
	if v481 == int32(0) {
		goto L114
	} else {
		goto L115
	}
L110:
	;
	goto L111
L111:
	;
	v565 = F_bms_is_member(m, v459, v258)
	mBase = m.M
	v566 = m.ExcPending
	if v566 != 0 {
		goto L1
	} else {
		goto L132
	}
L112:
	;
	v504 = *(*int32)(unsafe.Add(mBase, uint32(v24)+44))
	v505 = int32(1)
	*(*int32)(unsafe.Add(mBase, uint32(v24)+44)) = v504 + v505
	v510 = v503 + v504*int32(40)
	*(*uint8)(unsafe.Add(mBase, uint32(v510)+24)) = uint8(v505)
	*(*int64)(unsafe.Add(mBase, uint32(v510)+16)) = int64(0)
	v515 = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v510)+12)) = v515
	*(*int32)(unsafe.Add(mBase, uint32(v510)+8)) = v330
	*(*int32)(unsafe.Add(mBase, uint32(v510)+4)) = v332
	*(*int32)(unsafe.Add(mBase, uint32(v510))) = int32(25)
	v521 = *(*int64)(unsafe.Add(mBase, uint32(v21)+39))
	*(*int64)(unsafe.Add(mBase, uint32(v510)+32)) = v521
	v523 = *(*int64)(unsafe.Add(mBase, uint32(v21)+32))
	*(*int64)(unsafe.Add(mBase, uint32(v510)+25)) = v523
	v525 = *(*int32)(unsafe.Add(mBase, uint32(v24)+48))
	if v525 == v515 {
		goto L124
	} else {
		goto L125
	}
L113:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v24)+28)) = v501
	v503 = v501
	goto L112
L114:
	;
	v484 = int32(16)
	*(*int32)(unsafe.Add(mBase, uint32(v24)+48)) = v484
	v488 = F_palloc_mul(m, int32(40), v484)
	mBase = m.M
	v489 = m.ExcPending
	if v489 != 0 {
		goto L1
	} else {
		goto L117
	}
L115:
	;
	goto L116
L116:
	;
	v490 = *(*int32)(unsafe.Add(mBase, uint32(v24)+44))
	if v490 != v481 {
		goto L118
	} else {
		goto L119
	}
L117:
	;
	v501 = v488
	goto L113
L118:
	;
	v492 = *(*int32)(unsafe.Add(mBase, uint32(v24)+28))
	v503 = v492
	goto L112
L119:
	;
	goto L120
L120:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v24)+48)) = v481 << (uint(int32(1)) % 32)
	v496 = *(*int32)(unsafe.Add(mBase, uint32(v24)+28))
	v499 = F_repalloc(m, v496, v481*int32(80))
	mBase = m.M
	v500 = m.ExcPending
	if v500 != 0 {
		goto L1
	} else {
		goto L121
	}
L121:
	;
	v501 = v499
	goto L113
L122:
	;
	v548 = int32(1)
	v549 = *(*int32)(unsafe.Add(mBase, uint32(v24)+44))
	*(*int32)(unsafe.Add(mBase, uint32(v24)+44)) = v549 + v548
	v555 = v547 + v549*int32(40)
	*(*int64)(unsafe.Add(mBase, uint32(v555)+16)) = base.I64_extend_i32_u(v474)
	v558 = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v555)+12)) = v558
	*(*int32)(unsafe.Add(mBase, uint32(v555)+8)) = v330
	*(*int32)(unsafe.Add(mBase, uint32(v555)+4)) = v332
	*(*int32)(unsafe.Add(mBase, uint32(v555))) = int32(23)
	v607 = v332
	v608 = v330
	v609 = v548
	v610 = v555
	v612 = v558
	goto L108
L123:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v24)+28)) = v545
	v547 = v545
	goto L122
L124:
	;
	v528 = int32(16)
	*(*int32)(unsafe.Add(mBase, uint32(v24)+48)) = v528
	v532 = F_palloc_mul(m, int32(40), v528)
	mBase = m.M
	v533 = m.ExcPending
	if v533 != 0 {
		goto L1
	} else {
		goto L127
	}
L125:
	;
	goto L126
L126:
	;
	v534 = *(*int32)(unsafe.Add(mBase, uint32(v24)+44))
	if v534 != v525 {
		goto L128
	} else {
		goto L129
	}
L127:
	;
	v545 = v532
	goto L123
L128:
	;
	v536 = *(*int32)(unsafe.Add(mBase, uint32(v24)+28))
	v547 = v536
	goto L122
L129:
	;
	goto L130
L130:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v24)+48)) = v525 << (uint(int32(1)) % 32)
	v540 = *(*int32)(unsafe.Add(mBase, uint32(v24)+28))
	v543 = F_repalloc(m, v540, v525*int32(80))
	mBase = m.M
	v544 = m.ExcPending
	if v544 != 0 {
		goto L1
	} else {
		goto L131
	}
L131:
	;
	v545 = v543
	goto L123
L132:
	;
	if v565 != 0 {
		v618 = v455
		v619 = v456
		v620 = v457
		v622 = v464
		v623 = v469
		goto L107
	} else {
		goto L133
	}
L133:
	;
	v567 = *(*int32)(unsafe.Add(mBase, uint32(v24)+48))
	if v567 == int32(0) {
		goto L136
	} else {
		goto L137
	}
L134:
	;
	v590 = *(*int32)(unsafe.Add(mBase, uint32(v24)+44))
	*(*int32)(unsafe.Add(mBase, uint32(v24)+44)) = v590 + int32(1)
	v596 = v589 + v590*int32(40)
	*(*int64)(unsafe.Add(mBase, uint32(v596)+16)) = base.I64_extend_i32_u(v474) * int64(4294967297)
	*(*int32)(unsafe.Add(mBase, uint32(v596)+12)) = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v596)+8)) = v456
	*(*int32)(unsafe.Add(mBase, uint32(v596)+4)) = v455
	*(*int32)(unsafe.Add(mBase, uint32(v596))) = int32(20)
	v607 = v455
	v608 = v456
	v609 = v457
	v610 = v596
	v612 = v474
	goto L108
L135:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v24)+28)) = v587
	v589 = v587
	goto L134
L136:
	;
	v570 = int32(16)
	*(*int32)(unsafe.Add(mBase, uint32(v24)+48)) = v570
	v574 = F_palloc_mul(m, int32(40), v570)
	mBase = m.M
	v575 = m.ExcPending
	if v575 != 0 {
		goto L1
	} else {
		goto L139
	}
L137:
	;
	goto L138
L138:
	;
	v576 = *(*int32)(unsafe.Add(mBase, uint32(v24)+44))
	if v576 != v567 {
		goto L140
	} else {
		goto L141
	}
L139:
	;
	v587 = v574
	goto L135
L140:
	;
	v578 = *(*int32)(unsafe.Add(mBase, uint32(v24)+28))
	v589 = v578
	goto L134
L141:
	;
	goto L142
L142:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v24)+48)) = v567 << (uint(int32(1)) % 32)
	v582 = *(*int32)(unsafe.Add(mBase, uint32(v24)+28))
	v585 = F_repalloc(m, v582, v567*int32(80))
	mBase = m.M
	v586 = m.ExcPending
	if v586 != 0 {
		goto L1
	} else {
		goto L143
	}
L143:
	;
	v587 = v585
	goto L135
L144:
	;
	goto L106
L145:
	;
	F_errcode(m, int32(67141764))
	mBase = m.M
	v635 = m.ExcPending
	if v635 != 0 {
		goto L1
	} else {
		goto L146
	}
L146:
	;
	F_errmsg(m, int32(_a_F_ExecBuildUpdateProjection_4), int32(0))
	mBase = m.M
	v639 = m.ExcPending
	if v639 != 0 {
		goto L1
	} else {
		goto L147
	}
L147:
	;
	v642 = F_errdetail(m, int32(_a_F_ExecBuildUpdateProjection_5), int32(0))
	mBase = m.M
	v643 = m.ExcPending
	if v643 != 0 {
		goto L1
	} else {
		goto L148
	}
L148:
	;
	F_errfinish(m, int32(_a_F_ExecBuildUpdateProjection_1), int32(668), int32(_a_F_ExecBuildUpdateProjection_2))
	mBase = m.M
	v648 = m.ExcPending
	if v648 != 0 {
		goto L1
	} else {
		goto L149
	}
L149:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L150:
	;
	F_errcode(m, int32(67141764))
	mBase = m.M
	v655 = m.ExcPending
	if v655 != 0 {
		goto L1
	} else {
		goto L151
	}
L151:
	;
	F_errmsg(m, int32(_a_F_ExecBuildUpdateProjection_4), int32(0))
	mBase = m.M
	v659 = m.ExcPending
	if v659 != 0 {
		goto L1
	} else {
		goto L152
	}
L152:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v21))) = v381
	v662 = F_errdetail(m, int32(_a_F_ExecBuildUpdateProjection_6), v21)
	mBase = m.M
	v663 = m.ExcPending
	if v663 != 0 {
		goto L1
	} else {
		goto L153
	}
L153:
	;
	F_errfinish(m, int32(_a_F_ExecBuildUpdateProjection_1), int32(676), int32(_a_F_ExecBuildUpdateProjection_2))
	mBase = m.M
	v668 = m.ExcPending
	if v668 != 0 {
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
	v675 = m.ExcPending
	if v675 != 0 {
		goto L1
	} else {
		goto L156
	}
L156:
	;
	F_errmsg(m, int32(_a_F_ExecBuildUpdateProjection_4), int32(0))
	mBase = m.M
	v679 = m.ExcPending
	if v679 != 0 {
		goto L1
	} else {
		goto L157
	}
L157:
	;
	v680 = *(*int32)(unsafe.Add(mBase, uint32(v400)+68))
	v681 = F_format_type_be(m, v680)
	mBase = m.M
	v682 = m.ExcPending
	if v682 != 0 {
		goto L1
	} else {
		goto L158
	}
L158:
	;
	v683 = *(*int32)(unsafe.Add(mBase, uint32(v395)+4))
	v684 = F_exprType(m, v683)
	mBase = m.M
	v685 = m.ExcPending
	if v685 != 0 {
		goto L1
	} else {
		goto L159
	}
L159:
	;
	v686 = F_format_type_be(m, v684)
	mBase = m.M
	v687 = m.ExcPending
	if v687 != 0 {
		goto L1
	} else {
		goto L160
	}
L160:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v21)+24)) = v686
	*(*int32)(unsafe.Add(mBase, uint32(v21)+20)) = v381
	*(*int32)(unsafe.Add(mBase, uint32(v21)+16)) = v681
	v694 = F_errdetail(m, int32(_a_F_ExecBuildUpdateProjection_7), v19+int32(-48))
	mBase = m.M
	v695 = m.ExcPending
	if v695 != 0 {
		goto L1
	} else {
		goto L161
	}
L161:
	;
	F_errfinish(m, int32(_a_F_ExecBuildUpdateProjection_1), int32(684), int32(_a_F_ExecBuildUpdateProjection_2))
	mBase = m.M
	v700 = m.ExcPending
	if v700 != 0 {
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
	v772 = *(*int32)(unsafe.Add(mBase, uint32(v24)+44))
	v773 = int32(1)
	*(*int32)(unsafe.Add(mBase, uint32(v24)+44)) = v772 + v773
	v778 = v771 + v772*int32(40)
	*(*uint8)(unsafe.Add(mBase, uint32(v778)+24)) = uint8(v733)
	*(*int64)(unsafe.Add(mBase, uint32(v778)+16)) = base.I64_extend_i32_u(v740) | base.I64_extend_i32_u(v745)<<(uint(int64(32))%64)
	*(*int32)(unsafe.Add(mBase, uint32(v778)+12)) = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v778)+8)) = v732
	*(*int32)(unsafe.Add(mBase, uint32(v778)+4)) = v731
	*(*int32)(unsafe.Add(mBase, uint32(v778))) = v773
	v792 = *(*int64)(unsafe.Add(mBase, uint32(v21)+32))
	*(*int64)(unsafe.Add(mBase, uint32(v778)+25)) = v792
	v794 = *(*int64)(unsafe.Add(mBase, uint32(v21)+39))
	*(*int64)(unsafe.Add(mBase, uint32(v778)+32)) = v794
	v796 = F_jit_compile_expr(m, v319)
	mBase = m.M
	v797 = m.ExcPending
	if v797 != 0 {
		goto L1
	} else {
		goto L173
	}
L164:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v24)+28)) = v769
	v771 = v769
	goto L163
L165:
	;
	v752 = int32(16)
	*(*int32)(unsafe.Add(mBase, uint32(v24)+48)) = v752
	v756 = F_palloc_mul(m, int32(40), v752)
	mBase = m.M
	v757 = m.ExcPending
	if v757 != 0 {
		goto L1
	} else {
		goto L168
	}
L166:
	;
	goto L167
L167:
	;
	v758 = *(*int32)(unsafe.Add(mBase, uint32(v24)+44))
	if v758 != v749 {
		goto L169
	} else {
		goto L170
	}
L168:
	;
	v769 = v756
	goto L164
L169:
	;
	v760 = *(*int32)(unsafe.Add(mBase, uint32(v24)+28))
	v771 = v760
	goto L163
L170:
	;
	goto L171
L171:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v24)+48)) = v749 << (uint(int32(1)) % 32)
	v764 = *(*int32)(unsafe.Add(mBase, uint32(v24)+28))
	v767 = F_repalloc(m, v764, v749*int32(80))
	mBase = m.M
	v768 = m.ExcPending
	if v768 != 0 {
		goto L1
	} else {
		goto L172
	}
L172:
	;
	v769 = v767
	goto L164
L173:
	;
	if v796 == int32(0) {
		goto L174
	} else {
		goto L175
	}
L174:
	;
	F_ExecReadyInterpretedExpr(m, v319)
	mBase = m.M
	v801 = m.ExcPending
	if v801 != 0 {
		goto L1
	} else {
		goto L177
	}
L175:
	;
	goto L176
L176:
	;
	m.G0 = v21 - int32(-64)
	return v24
L177:
	;
	goto L176
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
	*(*int64)(unsafe.Add(mBase, uint32(v11))) = int64(1443109011760)
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
	switch v3 - int32(336) {
	case 0, 1, 3, 4, 28, 29, 30, 31, 35, 38, 39, 40, 41:
		v20 = v2
		return v20
	case 2:
		v12 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
		if v12 != int32(293) {
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
