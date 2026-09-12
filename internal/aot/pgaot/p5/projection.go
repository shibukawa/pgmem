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
	var v21 int32
	_ = v21
	var v23 int32
	_ = v23
	var v26 int32
	_ = v26
	var v29 int32
	_ = v29
	var v32 int64
	_ = v32
	var v40 int32
	_ = v40
	var v44 int32
	_ = v44
	var v50 int32
	_ = v50
	var v53 int32
	_ = v53
	var v56 int32
	_ = v56
	var v57 int32
	_ = v57
	var v58 int32
	_ = v58
	var v65 int32
	_ = v65
	var v66 int32
	_ = v66
	var v68 int32
	_ = v68
	var v83 int32
	_ = v83
	var v84 int32
	_ = v84
	var v91 int32
	_ = v91
	var v93 int32
	_ = v93
	var v103 int32
	_ = v103
	var v137 int32
	_ = v137
	var v150 int32
	_ = v150
	var v162 int32
	_ = v162
	var v187 int32
	_ = v187
	var v191 int32
	_ = v191
	var v196 int32
	_ = v196
	var v200 int32
	_ = v200
	var v204 int32
	_ = v204
	var v209 int32
	_ = v209
	var v210 int32
	_ = v210
	var v211 int32
	_ = v211
	var v221 int32
	_ = v221
	var v226 int32
	_ = v226
	var v235 int32
	_ = v235
	var v239 int32
	_ = v239
	var v240 int32
	_ = v240
	var v241 int32
	_ = v241
	var v243 int32
	_ = v243
	var v244 int32
	_ = v244
	var v257 int32
	_ = v257
	var v276 int32
	_ = v276
	var v279 int32
	_ = v279
	var v282 int32
	_ = v282
	var v288 int32
	_ = v288
	var v289 int32
	_ = v289
	var v300 int32
	_ = v300
	var v315 int32
	_ = v315
	var v319 int32
	_ = v319
	var v320 int32
	_ = v320
	var v321 int32
	_ = v321
	var v346 int32
	_ = v346
	var v349 int32
	_ = v349
	var v350 int32
	_ = v350
	var v355 int32
	_ = v355
	var v357 int32
	_ = v357
	var v359 int32
	_ = v359
	var v362 int32
	_ = v362
	var v369 int32
	_ = v369
	var v371 int32
	_ = v371
	var v375 int32
	_ = v375
	var v378 int32
	_ = v378
	var v385 int32
	_ = v385
	var v389 int32
	_ = v389
	var v391 int32
	_ = v391
	var v395 int32
	_ = v395
	var v396 int32
	_ = v396
	var v400 int32
	_ = v400
	var v403 int32
	_ = v403
	var v405 int32
	_ = v405
	var v408 int32
	_ = v408
	var v410 int32
	_ = v410
	var v413 int32
	_ = v413
	var v418 int32
	_ = v418
	var v419 int32
	_ = v419
	var v422 int32
	_ = v422
	var v423 int32
	_ = v423
	var v426 int32
	_ = v426
	var v427 int32
	_ = v427
	var v428 int32
	_ = v428
	var v429 int32
	_ = v429
	var v430 int32
	_ = v430
	var v432 int32
	_ = v432
	var v434 int32
	_ = v434
	var v435 int32
	_ = v435
	var v441 int32
	_ = v441
	var v442 int32
	_ = v442
	var v443 int32
	_ = v443
	var v445 int32
	_ = v445
	var v449 int32
	_ = v449
	var v452 int32
	_ = v452
	var v453 int32
	_ = v453
	var v454 int32
	_ = v454
	var v456 int32
	_ = v456
	var v457 int32
	_ = v457
	var v463 int32
	_ = v463
	var v475 int32
	_ = v475
	var v481 int32
	_ = v481
	var v482 int32
	_ = v482
	var v483 int32
	_ = v483
	var v485 int32
	_ = v485
	var v489 int32
	_ = v489
	var v492 int32
	_ = v492
	var v493 int32
	_ = v493
	var v494 int32
	_ = v494
	var v496 int32
	_ = v496
	var v499 int32
	_ = v499
	var v505 int32
	_ = v505
	var v512 int32
	_ = v512
	var v514 int32
	_ = v514
	var v515 int32
	_ = v515
	var v516 int64
	_ = v516
	var v525 int32
	_ = v525
	var v528 int32
	_ = v528
	var v532 int32
	_ = v532
	var v533 int32
	_ = v533
	var v537 int32
	_ = v537
	var v539 int32
	_ = v539
	var v540 int32
	_ = v540
	var v547 int32
	_ = v547
	var v548 int32
	_ = v548
	var v552 int32
	_ = v552
	var v555 int32
	_ = v555
	var v561 int32
	_ = v561
	var v562 int32
	_ = v562
	var v563 int32
	_ = v563
	var v565 int32
	_ = v565
	var v569 int32
	_ = v569
	var v572 int32
	_ = v572
	var v573 int32
	_ = v573
	var v574 int32
	_ = v574
	var v576 int32
	_ = v576
	var v577 int32
	_ = v577
	var v578 int32
	_ = v578
	var v583 int32
	_ = v583
	var v587 int32
	_ = v587
	var v595 int64
	_ = v595
	var v597 int64
	_ = v597
	var v599 int32
	_ = v599
	var v605 int32
	_ = v605
	var v606 int32
	_ = v606
	var v608 int32
	_ = v608
	var v610 int32
	_ = v610
	var v611 int32
	_ = v611
	var v613 int32
	_ = v613
	var v617 int32
	_ = v617
	var v620 int32
	_ = v620
	var v621 int32
	_ = v621
	var v623 int32
	_ = v623
	var v624 int32
	_ = v624
	var v627 int32
	_ = v627
	var v633 int32
	_ = v633
	var v634 int32
	_ = v634
	var v635 int32
	_ = v635
	var v637 int32
	_ = v637
	var v641 int32
	_ = v641
	var v644 int32
	_ = v644
	var v645 int32
	_ = v645
	var v646 int32
	_ = v646
	var v648 int32
	_ = v648
	var v650 int32
	_ = v650
	var v651 int32
	_ = v651
	var v653 int32
	_ = v653
	var v654 int32
	_ = v654
	var v655 int32
	_ = v655
	var v656 int32
	_ = v656
	var v657 int32
	_ = v657
	var v658 int32
	_ = v658
	var v659 int32
	_ = v659
	var v660 int32
	_ = v660
	var v666 int32
	_ = v666
	var v674 int64
	_ = v674
	var v676 int64
	_ = v676
	var v679 int32
	_ = v679
	var v681 int32
	_ = v681
	var v682 int32
	_ = v682
	var v684 int32
	_ = v684
	var v685 int32
	_ = v685
	var v686 int32
	_ = v686
	var v688 int32
	_ = v688
	var v689 int32
	_ = v689
	var v694 int32
	_ = v694
	var v697 int32
	_ = v697
	var v701 int32
	_ = v701
	var v702 int32
	_ = v702
	var v703 int32
	_ = v703
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
	var v709 int32
	_ = v709
	var v717 int32
	_ = v717
	var v722 int32
	_ = v722
	var v726 int32
	_ = v726
	var v729 int32
	_ = v729
	var v733 int32
	_ = v733
	var v737 int32
	_ = v737
	var v742 int32
	_ = v742
	var v747 int32
	_ = v747
	var v750 int32
	_ = v750
	var v754 int32
	_ = v754
	var v758 int32
	_ = v758
	var v763 int32
	_ = v763
	var v765 int32
	_ = v765
	var v768 int32
	_ = v768
	var v774 int32
	_ = v774
	var v776 int32
	_ = v776
	var v777 int32
	_ = v777
	var v784 int32
	_ = v784
	var v790 int32
	_ = v790
	var v791 int32
	_ = v791
	var v792 int32
	_ = v792
	var v794 int32
	_ = v794
	var v798 int32
	_ = v798
	var v801 int32
	_ = v801
	var v802 int32
	_ = v802
	var v803 int32
	_ = v803
	var v805 int32
	_ = v805
	var v806 int32
	_ = v806
	var v807 int32
	_ = v807
	var v812 int32
	_ = v812
	var v826 int64
	_ = v826
	var v828 int64
	_ = v828
	var v830 int32
	_ = v830
	var v831 int32
	_ = v831
	var v835 int32
	_ = v835
	v8 = int32(0)
	v21 = m.G0
	v23 = v21 + int32(-64)
	m.G0 = v23
	v26 = F_palloc0(m, int32(76))
	mBase = m.M
	v29 = m.ExcPending
	if v29 != 0 {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	return int32(0)
L2:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v26))) = int32(384)
	v32 = int64(0)
	*(*int64)(unsafe.Add(mBase, uint32(v23)+56)) = v32
	*(*int64)(unsafe.Add(mBase, uint32(v23)+48)) = v32
	*(*int64)(unsafe.Add(mBase, uint32(v23)+40)) = v32
	*(*int64)(unsafe.Add(mBase, uint32(v23)+32)) = v32
	v40 = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v26)+48)) = v40
	*(*int32)(unsafe.Add(mBase, uint32(v26)+44)) = l6
	if l1 != 0 {
		goto L3
	} else {
		goto L4
	}
L3:
	;
	v44 = l0
	goto L5
L4:
	;
	v44 = v40
	goto L5
L5:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v26)+28)) = v44
	*(*int32)(unsafe.Add(mBase, uint32(v26)+4)) = int32(380)
	*(*int32)(unsafe.Add(mBase, uint32(v26)+72)) = l4
	*(*int32)(unsafe.Add(mBase, uint32(v26)+16)) = l5
	if l0 != 0 {
		goto L12
	} else {
		goto L13
	}
L6:
	;
	v289 = *(*int32)(unsafe.Add(mBase, uint32(l3)))
	if v289 <= int32(0) {
		goto L45
	} else {
		goto L46
	}
L7:
	;
	v210 = int32(0)
	v211 = *(*int32)(unsafe.Add(mBase, uint32(l2)+4))
	if v210 < v211 {
		goto L38
	} else {
		goto L39
	}
L8:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v200 = m.ExcPending
	if v200 != 0 {
		goto L1
	} else {
		goto L35
	}
L9:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v187 = m.ExcPending
	if v187 != 0 {
		goto L1
	} else {
		goto L32
	}
L10:
	;
	v162 = *(*int32)(unsafe.Add(mBase, uint32(l2)+4))
	if v150 == v162 {
		goto L7
	} else {
		goto L31
	}
L11:
	;
	v137 = int32(0)
	v276 = v137
	v279 = v137
	v282 = int32(1)
	v288 = l2 + int32(4)
	goto L6
L12:
	;
	v50 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	if int32(0) < v50 {
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
		v150 = v8
		goto L10
	} else {
		goto L30
	}
L15:
	;
	v53 = int32(0)
	if v53 < v50 {
		goto L18
	} else {
		goto L19
	}
L16:
	;
	v103 = v8
	goto L17
L17:
	;
	if l2 != 0 {
		v150 = v103
		goto L10
	} else {
		goto L28
	}
L18:
	;
	v56 = v50
	goto L20
L19:
	;
	v56 = v53
	goto L20
L20:
	;
	v57 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	v58 = int32(0)
	v65 = v58
	v66 = v58
	v68 = v8
	goto L21
L21:
	;
	v83 = *(*int32)(unsafe.Add(mBase, uint32(v57+v66<<(uint(int32(2))%32))))
	v84 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v83)+26)))
	if v84 == int32(0) {
		goto L23
	} else {
		goto L24
	}
L22:
	;
	v103 = v91
	goto L17
L23:
	;
	if v65&int32(1) != 0 {
		goto L8
	} else {
		goto L26
	}
L24:
	;
	v91 = v68
	goto L25
L25:
	;
	v93 = v66 + int32(1)
	if v93 != v56 {
		v65 = v84
		v66 = v93
		v68 = v91
		goto L21
	} else {
		goto L27
	}
L26:
	;
	v91 = v68 + int32(1)
	goto L25
L27:
	;
	goto L22
L28:
	;
	if v103 == int32(0) {
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
	F_errmsg_internal(m, int32(73160), int32(0))
	mBase = m.M
	v191 = m.ExcPending
	if v191 != 0 {
		goto L1
	} else {
		goto L33
	}
L33:
	;
	F_errfinish(m, int32(486589), int32(601), int32(251454))
	mBase = m.M
	v196 = m.ExcPending
	if v196 != 0 {
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
	F_errmsg_internal(m, int32(223567), int32(0))
	mBase = m.M
	v204 = m.ExcPending
	if v204 != 0 {
		goto L1
	} else {
		goto L36
	}
L36:
	;
	F_errfinish(m, int32(486589), int32(594), int32(251454))
	mBase = m.M
	v209 = m.ExcPending
	if v209 != 0 {
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
	v221 = int32(0)
	v226 = v210
	goto L41
L39:
	;
	v257 = v210
	goto L40
L40:
	;
	v276 = v150
	v279 = v257
	v282 = v8
	v288 = l2 + int32(4)
	goto L6
L41:
	;
	v235 = *(*int32)(unsafe.Add(mBase, uint32(l2)+12))
	v239 = int32(*(*int16)(unsafe.Add(mBase, uint32(v235+v221<<(uint(int32(2))%32)))))
	v240 = F_bms_add_member(m, v226, v239)
	mBase = m.M
	v241 = m.ExcPending
	if v241 != 0 {
		goto L1
	} else {
		goto L43
	}
L42:
	;
	v257 = v240
	goto L40
L43:
	;
	v243 = v221 + int32(1)
	v244 = *(*int32)(unsafe.Add(mBase, uint32(l2)+4))
	if v243 < v244 {
		v221 = v243
		v226 = v240
		goto L41
	} else {
		goto L44
	}
L44:
	;
	goto L42
L45:
	;
	v346 = v26 + int32(4)
	if l1 != 0 {
		goto L55
	} else {
		goto L56
	}
L46:
	;
	v300 = v289
	goto L47
L47:
	;
	v315 = v300 - int32(1)
	v319 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l3+int32(29)+v315<<(uint(int32(4))%32)))))
	if v319 != 0 {
		goto L49
	} else {
		goto L50
	}
L48:
	;
	goto L45
L49:
	;
	if base.Ui32(int32(1)) < base.Ui32(v300) {
		v300 = v315
		goto L47
	} else {
		goto L53
	}
L50:
	;
	v320 = F_bms_is_member(m, v300, v279)
	mBase = m.M
	v321 = m.ExcPending
	if v321 != 0 {
		goto L1
	} else {
		goto L51
	}
L51:
	;
	if v320 != 0 {
		goto L49
	} else {
		goto L52
	}
L52:
	;
	*(*uint16)(unsafe.Add(mBase, uint32(v23)+52)) = uint16(v300)
	goto L45
L53:
	;
	goto L48
L54:
	;
	F_ExecPushExprSetupSteps(m, v346, v21+int32(-16))
	mBase = m.M
	v355 = m.ExcPending
	if v355 != 0 {
		goto L1
	} else {
		goto L59
	}
L55:
	;
	v349 = F_expr_setup_walker(m, l0, v21+int32(-16))
	mBase = m.M
	v350 = m.ExcPending
	if v350 != 0 {
		goto L1
	} else {
		goto L58
	}
L56:
	;
	goto L57
L57:
	;
	*(*uint16)(unsafe.Add(mBase, uint32(v23)+50)) = uint16(v276)
	goto L54
L58:
	;
	goto L54
L59:
	;
	v357 = v26 + int32(9)
	v359 = v26 + int32(12)
	v362 = int32(0)
	v369 = v362
	v371 = v362
	v375 = v362
	v378 = v8
	goto L64
L60:
	;
	v784 = *(*int32)(unsafe.Add(mBase, uint32(v26)+40))
	if v784 == int32(0) {
		goto L168
	} else {
		goto L169
	}
L61:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v747 = m.ExcPending
	if v747 != 0 {
		goto L1
	} else {
		goto L161
	}
L62:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v726 = m.ExcPending
	if v726 != 0 {
		goto L1
	} else {
		goto L156
	}
L63:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v694 = m.ExcPending
	if v694 != 0 {
		goto L1
	} else {
		goto L148
	}
L64:
	;
	v385 = int32(0)
	if l0 == v385 {
		v395 = v385
		goto L66
	} else {
		goto L67
	}
L65:
	;
	v525 = int32(0)
	v528 = v525
	v532 = int32(1)
	v533 = v369
	v537 = v375
	v539 = v525
	v540 = v378
	goto L107
L66:
	;
	if v282 != 0 {
		goto L71
	} else {
		goto L72
	}
L67:
	;
	v389 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	if v389 <= v371 {
		v395 = int32(0)
		goto L66
	} else {
		goto L68
	}
L68:
	;
	v391 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	v395 = v391 + v371<<(uint(int32(2))%32)
	goto L66
L69:
	;
	goto L65
L70:
	;
	v410 = int32(*(*int16)(unsafe.Add(mBase, uint32(v403))))
	if v410 <= int32(0) {
		goto L61
	} else {
		goto L77
	}
L71:
	;
	v405 = *(*int32)(unsafe.Add(mBase, uint32(l3)))
	if int32(0) < v405 {
		goto L69
	} else {
		goto L76
	}
L72:
	;
	v396 = *(*int32)(unsafe.Add(mBase, uint32(v288)))
	if v396 <= v371 {
		goto L71
	} else {
		goto L73
	}
L73:
	;
	if v395 == int32(0) {
		goto L71
	} else {
		goto L74
	}
L74:
	;
	v400 = *(*int32)(unsafe.Add(mBase, uint32(l2)+12))
	v403 = v400 + v371<<(uint(int32(2))%32)
	if v403 != 0 {
		goto L70
	} else {
		goto L75
	}
L75:
	;
	goto L71
L76:
	;
	v408 = int32(0)
	v765 = v408
	v768 = v369
	v774 = v375
	v776 = v408
	v777 = v378
	goto L60
L77:
	;
	v413 = *(*int32)(unsafe.Add(mBase, uint32(l3)))
	if v413 < v410 {
		goto L61
	} else {
		goto L78
	}
L78:
	;
	v418 = int32(1)
	v419 = v410 - v418
	v422 = l3 + int32(20) + v413<<(uint(int32(4))%32) + v419*int32(100)
	v423 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v422)+91)))
	if v423 == v418 {
		goto L62
	} else {
		goto L79
	}
L79:
	;
	v426 = *(*int32)(unsafe.Add(mBase, uint32(v395)))
	v427 = *(*int32)(unsafe.Add(mBase, uint32(v426)+4))
	v428 = F_exprType(m, v427)
	mBase = m.M
	v429 = m.ExcPending
	if v429 != 0 {
		goto L1
	} else {
		goto L80
	}
L80:
	;
	v430 = *(*int32)(unsafe.Add(mBase, uint32(v422)+68))
	if v428 != v430 {
		goto L63
	} else {
		goto L81
	}
L81:
	;
	if l1 != 0 {
		goto L83
	} else {
		goto L84
	}
L82:
	;
	v516 = int64(0)
	*(*int64)(unsafe.Add(mBase, uint32(v512)+24)) = v516
	*(*int64)(unsafe.Add(mBase, uint32(v512)+32)) = v516
	v369 = v419
	v371 = v371 + int32(1)
	v375 = v514
	v378 = v515
	goto L64
L83:
	;
	v432 = *(*int32)(unsafe.Add(mBase, uint32(v426)+4))
	F_ExecInitExprRec(m, v432, v346, v359, v357)
	mBase = m.M
	v434 = m.ExcPending
	if v434 != 0 {
		goto L1
	} else {
		goto L86
	}
L84:
	;
	goto L85
L85:
	;
	v475 = *(*int32)(unsafe.Add(mBase, uint32(v26)+40))
	if v475 == int32(0) {
		goto L99
	} else {
		goto L100
	}
L86:
	;
	v435 = *(*int32)(unsafe.Add(mBase, uint32(v26)+40))
	if v435 == int32(0) {
		goto L89
	} else {
		goto L90
	}
L87:
	;
	v457 = *(*int32)(unsafe.Add(mBase, uint32(v26)+36))
	*(*int32)(unsafe.Add(mBase, uint32(v26)+36)) = v457 + int32(1)
	v463 = v456 + v457*int32(40)
	*(*int32)(unsafe.Add(mBase, uint32(v463)+20)) = v375&int32(255) | v378<<(uint(int32(8))%32)
	*(*int32)(unsafe.Add(mBase, uint32(v463)+16)) = v419
	*(*int64)(unsafe.Add(mBase, uint32(v463)+8)) = int64(0)
	*(*int64)(unsafe.Add(mBase, uint32(v463))) = int64(23)
	v512 = v463
	v514 = v375
	v515 = v378
	goto L82
L88:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v26)+20)) = v454
	v456 = v454
	goto L87
L89:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v26)+40)) = int32(16)
	v441 = F_palloc(m, int32(640))
	mBase = m.M
	v442 = m.ExcPending
	if v442 != 0 {
		goto L1
	} else {
		goto L92
	}
L90:
	;
	goto L91
L91:
	;
	v443 = *(*int32)(unsafe.Add(mBase, uint32(v26)+36))
	if v443 != v435 {
		goto L93
	} else {
		goto L94
	}
L92:
	;
	v454 = v441
	goto L88
L93:
	;
	v445 = *(*int32)(unsafe.Add(mBase, uint32(v26)+20))
	v456 = v445
	goto L87
L94:
	;
	goto L95
L95:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v26)+40)) = v435 << (uint(int32(1)) % 32)
	v449 = *(*int32)(unsafe.Add(mBase, uint32(v26)+20))
	v452 = F_repalloc(m, v449, v435*int32(80))
	mBase = m.M
	v453 = m.ExcPending
	if v453 != 0 {
		goto L1
	} else {
		goto L96
	}
L96:
	;
	v454 = v452
	goto L88
L97:
	;
	v499 = *(*int32)(unsafe.Add(mBase, uint32(v26)+36))
	*(*int32)(unsafe.Add(mBase, uint32(v26)+36)) = v499 + int32(1)
	v505 = v496 + v499*int32(40)
	*(*int32)(unsafe.Add(mBase, uint32(v505)+20)) = v371
	*(*int32)(unsafe.Add(mBase, uint32(v505)+16)) = v419
	*(*int64)(unsafe.Add(mBase, uint32(v505)+8)) = int64(0)
	*(*int64)(unsafe.Add(mBase, uint32(v505))) = int64(19)
	v512 = v505
	v514 = v371
	v515 = int32(base.Ui32(v371) >> (uint(int32(8)) % 32))
	goto L82
L98:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v26)+20)) = v494
	v496 = v494
	goto L97
L99:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v26)+40)) = int32(16)
	v481 = F_palloc(m, int32(640))
	mBase = m.M
	v482 = m.ExcPending
	if v482 != 0 {
		goto L1
	} else {
		goto L102
	}
L100:
	;
	goto L101
L101:
	;
	v483 = *(*int32)(unsafe.Add(mBase, uint32(v26)+36))
	if v483 != v475 {
		goto L103
	} else {
		goto L104
	}
L102:
	;
	v494 = v481
	goto L98
L103:
	;
	v485 = *(*int32)(unsafe.Add(mBase, uint32(v26)+20))
	v496 = v485
	goto L97
L104:
	;
	goto L105
L105:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v26)+40)) = v475 << (uint(int32(1)) % 32)
	v489 = *(*int32)(unsafe.Add(mBase, uint32(v26)+20))
	v492 = F_repalloc(m, v489, v475*int32(80))
	mBase = m.M
	v493 = m.ExcPending
	if v493 != 0 {
		goto L1
	} else {
		goto L106
	}
L106:
	;
	v494 = v492
	goto L98
L107:
	;
	v547 = int32(1)
	v548 = v532 - v547
	v552 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l3+int32(29)+v548<<(uint(int32(4))%32)))))
	if v552 == v547 {
		goto L112
	} else {
		goto L113
	}
L108:
	;
	v765 = v679
	v768 = v681
	v774 = v684
	v776 = v685
	v777 = v686
	goto L60
L109:
	;
	v688 = v532 + int32(1)
	v689 = *(*int32)(unsafe.Add(mBase, uint32(l3)))
	if v688 <= v689 {
		v528 = v679
		v532 = v688
		v533 = v682
		v537 = v684
		v539 = v685
		v540 = v686
		goto L107
	} else {
		goto L147
	}
L110:
	;
	v660 = *(*int32)(unsafe.Add(mBase, uint32(v26)+36))
	*(*int32)(unsafe.Add(mBase, uint32(v26)+36)) = v660 + int32(1)
	v666 = v655 + v660*int32(40)
	*(*int32)(unsafe.Add(mBase, uint32(v666)+20)) = v654
	*(*int32)(unsafe.Add(mBase, uint32(v666)+16)) = v548
	*(*int32)(unsafe.Add(mBase, uint32(v666)+12)) = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v666)+8)) = v653
	*(*int32)(unsafe.Add(mBase, uint32(v666)+4)) = v657
	*(*int32)(unsafe.Add(mBase, uint32(v666))) = v659
	v674 = *(*int64)(unsafe.Add(mBase, uint32(v23)+32))
	*(*int64)(unsafe.Add(mBase, uint32(v666)+24)) = v674
	v676 = *(*int64)(unsafe.Add(mBase, uint32(v23)+40))
	*(*int64)(unsafe.Add(mBase, uint32(v666)+32)) = v676
	v679 = v653
	v681 = v548
	v682 = v548
	v684 = v656
	v685 = v657
	v686 = v658
	goto L109
L111:
	;
	v653 = v357
	v654 = v587
	v655 = v650
	v656 = v651
	v657 = v359
	v658 = v540
	v659 = int32(23)
	goto L110
L112:
	;
	v555 = *(*int32)(unsafe.Add(mBase, uint32(v26)+40))
	if v555 == int32(0) {
		goto L117
	} else {
		goto L118
	}
L113:
	;
	goto L114
L114:
	;
	v623 = F_bms_is_member(m, v532, v279)
	mBase = m.M
	v624 = m.ExcPending
	if v624 != 0 {
		goto L1
	} else {
		goto L133
	}
L115:
	;
	v577 = *(*int32)(unsafe.Add(mBase, uint32(v26)+36))
	v578 = int32(1)
	*(*int32)(unsafe.Add(mBase, uint32(v26)+36)) = v577 + v578
	v583 = v576 + v577*int32(40)
	v587 = v540<<(uint(int32(8))%32) | v578
	*(*int32)(unsafe.Add(mBase, uint32(v583)+20)) = v587
	*(*int64)(unsafe.Add(mBase, uint32(v583)+12)) = int64(0)
	*(*int32)(unsafe.Add(mBase, uint32(v583)+8)) = v357
	*(*int32)(unsafe.Add(mBase, uint32(v583)+4)) = v359
	*(*int32)(unsafe.Add(mBase, uint32(v583))) = int32(25)
	v595 = *(*int64)(unsafe.Add(mBase, uint32(v23)+40))
	*(*int64)(unsafe.Add(mBase, uint32(v583)+32)) = v595
	v597 = *(*int64)(unsafe.Add(mBase, uint32(v23)+32))
	*(*int64)(unsafe.Add(mBase, uint32(v583)+24)) = v597
	v599 = *(*int32)(unsafe.Add(mBase, uint32(v26)+40))
	if v599 == int32(0) {
		goto L127
	} else {
		goto L128
	}
L116:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v26)+20)) = v574
	v576 = v574
	goto L115
L117:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v26)+40)) = int32(16)
	v561 = F_palloc(m, int32(640))
	mBase = m.M
	v562 = m.ExcPending
	if v562 != 0 {
		goto L1
	} else {
		goto L120
	}
L118:
	;
	goto L119
L119:
	;
	v563 = *(*int32)(unsafe.Add(mBase, uint32(v26)+36))
	if v563 != v555 {
		goto L121
	} else {
		goto L122
	}
L120:
	;
	v574 = v561
	goto L116
L121:
	;
	v565 = *(*int32)(unsafe.Add(mBase, uint32(v26)+20))
	v576 = v565
	goto L115
L122:
	;
	goto L123
L123:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v26)+40)) = v555 << (uint(int32(1)) % 32)
	v569 = *(*int32)(unsafe.Add(mBase, uint32(v26)+20))
	v572 = F_repalloc(m, v569, v555*int32(80))
	mBase = m.M
	v573 = m.ExcPending
	if v573 != 0 {
		goto L1
	} else {
		goto L124
	}
L124:
	;
	v574 = v572
	goto L116
L125:
	;
	v613 = int32(1)
	*(*int32)(unsafe.Add(mBase, uint32(v26)+40)) = v599 << (uint(v613) % 32)
	v617 = *(*int32)(unsafe.Add(mBase, uint32(v26)+20))
	v620 = F_repalloc(m, v617, v599*int32(80))
	mBase = m.M
	v621 = m.ExcPending
	if v621 != 0 {
		goto L1
	} else {
		goto L132
	}
L126:
	;
	v650 = v611
	v651 = int32(1)
	goto L111
L127:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v26)+40)) = int32(16)
	v605 = F_palloc(m, int32(640))
	mBase = m.M
	v606 = m.ExcPending
	if v606 != 0 {
		goto L1
	} else {
		goto L130
	}
L128:
	;
	goto L129
L129:
	;
	v608 = *(*int32)(unsafe.Add(mBase, uint32(v26)+36))
	if v599 == v608 {
		goto L125
	} else {
		goto L131
	}
L130:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v26)+20)) = v605
	v611 = v605
	goto L126
L131:
	;
	v610 = *(*int32)(unsafe.Add(mBase, uint32(v26)+20))
	v611 = v610
	goto L126
L132:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v26)+20)) = v620
	v650 = v620
	v651 = v613
	goto L111
L133:
	;
	if v623 != 0 {
		goto L134
	} else {
		goto L135
	}
L134:
	;
	v679 = v528
	v681 = v533
	v682 = v533
	v684 = v537
	v685 = v539
	v686 = v540
	goto L109
L135:
	;
	goto L136
L136:
	;
	v627 = *(*int32)(unsafe.Add(mBase, uint32(v26)+40))
	if v627 == int32(0) {
		goto L139
	} else {
		goto L140
	}
L137:
	;
	v653 = v528
	v654 = v548
	v655 = v648
	v656 = v548
	v657 = v539
	v658 = int32(base.Ui32(v548) >> (uint(int32(8)) % 32))
	v659 = int32(20)
	goto L110
L138:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v26)+20)) = v646
	v648 = v646
	goto L137
L139:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v26)+40)) = int32(16)
	v633 = F_palloc(m, int32(640))
	mBase = m.M
	v634 = m.ExcPending
	if v634 != 0 {
		goto L1
	} else {
		goto L142
	}
L140:
	;
	goto L141
L141:
	;
	v635 = *(*int32)(unsafe.Add(mBase, uint32(v26)+36))
	if v635 != v627 {
		goto L143
	} else {
		goto L144
	}
L142:
	;
	v646 = v633
	goto L138
L143:
	;
	v637 = *(*int32)(unsafe.Add(mBase, uint32(v26)+20))
	v648 = v637
	goto L137
L144:
	;
	goto L145
L145:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v26)+40)) = v627 << (uint(int32(1)) % 32)
	v641 = *(*int32)(unsafe.Add(mBase, uint32(v26)+20))
	v644 = F_repalloc(m, v641, v627*int32(80))
	mBase = m.M
	v645 = m.ExcPending
	if v645 != 0 {
		goto L1
	} else {
		goto L146
	}
L146:
	;
	v646 = v644
	goto L138
L147:
	;
	goto L108
L148:
	;
	F_errcode(m, int32(67141764))
	mBase = m.M
	v697 = m.ExcPending
	if v697 != 0 {
		goto L1
	} else {
		goto L149
	}
L149:
	;
	F_errmsg(m, int32(320123), int32(0))
	mBase = m.M
	v701 = m.ExcPending
	if v701 != 0 {
		goto L1
	} else {
		goto L150
	}
L150:
	;
	v702 = *(*int32)(unsafe.Add(mBase, uint32(v422)+68))
	v703 = F_format_type_be(m, v702)
	mBase = m.M
	v704 = m.ExcPending
	if v704 != 0 {
		goto L1
	} else {
		goto L151
	}
L151:
	;
	v705 = *(*int32)(unsafe.Add(mBase, uint32(v426)+4))
	v706 = F_exprType(m, v705)
	mBase = m.M
	v707 = m.ExcPending
	if v707 != 0 {
		goto L1
	} else {
		goto L152
	}
L152:
	;
	v708 = F_format_type_be(m, v706)
	mBase = m.M
	v709 = m.ExcPending
	if v709 != 0 {
		goto L1
	} else {
		goto L153
	}
L153:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v23)+24)) = v708
	*(*int32)(unsafe.Add(mBase, uint32(v23)+20)) = v410
	*(*int32)(unsafe.Add(mBase, uint32(v23)+16)) = v703
	F_errdetail(m, int32(579224), v21+int32(-48))
	mBase = m.M
	v717 = m.ExcPending
	if v717 != 0 {
		goto L1
	} else {
		goto L154
	}
L154:
	;
	F_errfinish(m, int32(486589), int32(684), int32(251454))
	mBase = m.M
	v722 = m.ExcPending
	if v722 != 0 {
		goto L1
	} else {
		goto L155
	}
L155:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L156:
	;
	F_errcode(m, int32(67141764))
	mBase = m.M
	v729 = m.ExcPending
	if v729 != 0 {
		goto L1
	} else {
		goto L157
	}
L157:
	;
	F_errmsg(m, int32(320123), int32(0))
	mBase = m.M
	v733 = m.ExcPending
	if v733 != 0 {
		goto L1
	} else {
		goto L158
	}
L158:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v23))) = v410
	F_errdetail(m, int32(628542), v23)
	mBase = m.M
	v737 = m.ExcPending
	if v737 != 0 {
		goto L1
	} else {
		goto L159
	}
L159:
	;
	F_errfinish(m, int32(486589), int32(676), int32(251454))
	mBase = m.M
	v742 = m.ExcPending
	if v742 != 0 {
		goto L1
	} else {
		goto L160
	}
L160:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L161:
	;
	F_errcode(m, int32(67141764))
	mBase = m.M
	v750 = m.ExcPending
	if v750 != 0 {
		goto L1
	} else {
		goto L162
	}
L162:
	;
	F_errmsg(m, int32(320123), int32(0))
	mBase = m.M
	v754 = m.ExcPending
	if v754 != 0 {
		goto L1
	} else {
		goto L163
	}
L163:
	;
	F_errdetail(m, int32(567595), int32(0))
	mBase = m.M
	v758 = m.ExcPending
	if v758 != 0 {
		goto L1
	} else {
		goto L164
	}
L164:
	;
	F_errfinish(m, int32(486589), int32(668), int32(251454))
	mBase = m.M
	v763 = m.ExcPending
	if v763 != 0 {
		goto L1
	} else {
		goto L165
	}
L165:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L166:
	;
	v806 = *(*int32)(unsafe.Add(mBase, uint32(v26)+36))
	v807 = int32(1)
	*(*int32)(unsafe.Add(mBase, uint32(v26)+36)) = v806 + v807
	v812 = v805 + v806*int32(40)
	*(*int32)(unsafe.Add(mBase, uint32(v812)+20)) = v774&int32(255) | v777<<(uint(int32(8))%32)
	*(*int32)(unsafe.Add(mBase, uint32(v812)+16)) = v768
	*(*int32)(unsafe.Add(mBase, uint32(v812)+12)) = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v812)+8)) = v765
	*(*int32)(unsafe.Add(mBase, uint32(v812)+4)) = v776
	*(*int32)(unsafe.Add(mBase, uint32(v812))) = v807
	v826 = *(*int64)(unsafe.Add(mBase, uint32(v23)+32))
	*(*int64)(unsafe.Add(mBase, uint32(v812)+24)) = v826
	v828 = *(*int64)(unsafe.Add(mBase, uint32(v23)+40))
	*(*int64)(unsafe.Add(mBase, uint32(v812)+32)) = v828
	v830 = F_jit_compile_expr(m, v346)
	mBase = m.M
	v831 = m.ExcPending
	if v831 != 0 {
		goto L1
	} else {
		goto L176
	}
L167:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v26)+20)) = v803
	v805 = v803
	goto L166
L168:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v26)+40)) = int32(16)
	v790 = F_palloc(m, int32(640))
	mBase = m.M
	v791 = m.ExcPending
	if v791 != 0 {
		goto L1
	} else {
		goto L171
	}
L169:
	;
	goto L170
L170:
	;
	v792 = *(*int32)(unsafe.Add(mBase, uint32(v26)+36))
	if v792 != v784 {
		goto L172
	} else {
		goto L173
	}
L171:
	;
	v803 = v790
	goto L167
L172:
	;
	v794 = *(*int32)(unsafe.Add(mBase, uint32(v26)+20))
	v805 = v794
	goto L166
L173:
	;
	goto L174
L174:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v26)+40)) = v784 << (uint(int32(1)) % 32)
	v798 = *(*int32)(unsafe.Add(mBase, uint32(v26)+20))
	v801 = F_repalloc(m, v798, v784*int32(80))
	mBase = m.M
	v802 = m.ExcPending
	if v802 != 0 {
		goto L1
	} else {
		goto L175
	}
L175:
	;
	v803 = v801
	goto L167
L176:
	;
	if v830 == int32(0) {
		goto L177
	} else {
		goto L178
	}
L177:
	;
	F_ExecReadyInterpretedExpr(m, v346)
	mBase = m.M
	v835 = m.ExcPending
	if v835 != 0 {
		goto L1
	} else {
		goto L180
	}
L178:
	;
	goto L179
L179:
	;
	m.G0 = v23 - int32(-64)
	return v26
L180:
	;
	goto L179
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
	var v43 int32
	_ = v43
	var v44 float64
	_ = v44
	var v45 int32
	_ = v45
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
	v89 = *(*float64)(unsafe.Add(mBase, _consts[388]))
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
	v43 = int32(0)
	v44 = float64(1)
	v45 = *(*int32)(unsafe.Add(mBase, uint32(v39)+4))
	if v45 <= v43 {
		v76 = v44
		goto L7
	} else {
		goto L11
	}
L11:
	;
	v49 = v43
	v54 = v44
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
