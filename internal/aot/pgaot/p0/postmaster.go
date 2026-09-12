package p0

import (
	base "github.com/shibukawa/pgmem/internal/aot/pgaot/base"
	"unsafe"
)

func F_CheckPostmasterSignal(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v4 int32
	_ = v4
	var v7 int32
	_ = v7
	var v8 int32
	_ = v8
	v4 = *(*int32)(unsafe.Add(mBase, _consts[463]))
	v7 = v4 + l0<<(uint(int32(2))%32)
	v8 = *(*int32)(unsafe.Add(mBase, uint32(v7)))
	if v8 != 0 {
		*(*int32)(unsafe.Add(mBase, uint32(v7))) = int32(0)
	} else {
	}
	return base.B2i32(v8 != int32(0))
}
func F_PostmasterStateMachine(m *base.Module) {
	mBase := m.M
	_ = mBase
	var v10 int32
	_ = v10
	var v12 int32
	_ = v12
	var v15 int32
	_ = v15
	var v21 int32
	_ = v21
	var v25 int32
	_ = v25
	var v26 int32
	_ = v26
	var v31 int32
	_ = v31
	var v32 int32
	_ = v32
	var v36 int32
	_ = v36
	var v41 int32
	_ = v41
	var v47 int32
	_ = v47
	var v52 int32
	_ = v52
	var v53 int32
	_ = v53
	var v58 int32
	_ = v58
	var v59 int32
	_ = v59
	var v64 int32
	_ = v64
	var v65 int32
	_ = v65
	var v69 int32
	_ = v69
	var v72 int32
	_ = v72
	var v74 int32
	_ = v74
	var v75 int32
	_ = v75
	var v78 int32
	_ = v78
	var v80 int32
	_ = v80
	var v83 int32
	_ = v83
	var v89 int32
	_ = v89
	var v92 int32
	_ = v92
	var v93 int32
	_ = v93
	var v100 int32
	_ = v100
	var v101 int32
	_ = v101
	var v104 int32
	_ = v104
	var v107 int32
	_ = v107
	var v108 int32
	_ = v108
	var v113 int32
	_ = v113
	var v116 int32
	_ = v116
	var v120 int32
	_ = v120
	var v123 int32
	_ = v123
	var v128 int32
	_ = v128
	var v129 int32
	_ = v129
	var v133 int32
	_ = v133
	var v134 int32
	_ = v134
	var v138 int32
	_ = v138
	var v143 int32
	_ = v143
	var v144 int32
	_ = v144
	var v145 int32
	_ = v145
	var v147 int32
	_ = v147
	var v150 int32
	_ = v150
	var v152 int32
	_ = v152
	var v153 int32
	_ = v153
	var v155 int32
	_ = v155
	var v156 int32
	_ = v156
	var v175 int32
	_ = v175
	var v182 int32
	_ = v182
	var v192 int32
	_ = v192
	var v193 int32
	_ = v193
	var v198 int32
	_ = v198
	var v200 int32
	_ = v200
	var v204 int32
	_ = v204
	var v214 int32
	_ = v214
	var v222 int32
	_ = v222
	var v223 int32
	_ = v223
	var v237 int32
	_ = v237
	var v238 int32
	_ = v238
	var v242 int32
	_ = v242
	var v247 int32
	_ = v247
	var v253 int32
	_ = v253
	var v258 int32
	_ = v258
	var v271 int32
	_ = v271
	var v272 int32
	_ = v272
	var v274 int32
	_ = v274
	var v278 int32
	_ = v278
	var v283 int32
	_ = v283
	var v284 int32
	_ = v284
	var v288 int32
	_ = v288
	var v293 int32
	_ = v293
	var v299 int32
	_ = v299
	var v304 int32
	_ = v304
	var v309 int32
	_ = v309
	var v311 int32
	_ = v311
	var v312 int32
	_ = v312
	var v313 int32
	_ = v313
	var v318 int32
	_ = v318
	var v319 int32
	_ = v319
	var v324 int32
	_ = v324
	var v326 int32
	_ = v326
	var v328 int32
	_ = v328
	var v333 int32
	_ = v333
	var v344 int32
	_ = v344
	var v351 int32
	_ = v351
	var v352 int32
	_ = v352
	var v356 int32
	_ = v356
	var v361 int32
	_ = v361
	var v362 int32
	_ = v362
	var v366 int32
	_ = v366
	var v369 int32
	_ = v369
	var v372 int32
	_ = v372
	var v373 int32
	_ = v373
	var v377 int32
	_ = v377
	var v382 int32
	_ = v382
	var v388 int32
	_ = v388
	var v393 int32
	_ = v393
	var v399 int32
	_ = v399
	var v410 int32
	_ = v410
	var v414 int32
	_ = v414
	var v415 int32
	_ = v415
	var v418 int32
	_ = v418
	var v419 int32
	_ = v419
	var v423 int32
	_ = v423
	var v428 int32
	_ = v428
	var v434 int32
	_ = v434
	var v439 int32
	_ = v439
	var v444 int32
	_ = v444
	var v449 int32
	_ = v449
	var v458 int32
	_ = v458
	var v460 int32
	_ = v460
	var v467 int32
	_ = v467
	var v468 int32
	_ = v468
	var v481 int32
	_ = v481
	var v485 int32
	_ = v485
	var v488 int32
	_ = v488
	var v489 int32
	_ = v489
	var v493 int32
	_ = v493
	var v498 int32
	_ = v498
	var v504 int32
	_ = v504
	var v509 int32
	_ = v509
	var v514 int32
	_ = v514
	var v519 int32
	_ = v519
	var v521 int32
	_ = v521
	var v522 int32
	_ = v522
	var v526 int32
	_ = v526
	var v527 int32
	_ = v527
	var v531 int32
	_ = v531
	var v532 int32
	_ = v532
	var v536 int32
	_ = v536
	var v544 int32
	_ = v544
	var v545 int32
	_ = v545
	var v549 int32
	_ = v549
	var v554 int32
	_ = v554
	var v560 int32
	_ = v560
	var v565 int32
	_ = v565
	var v570 int32
	_ = v570
	var v575 int32
	_ = v575
	var v578 int32
	_ = v578
	var v579 int32
	_ = v579
	var v583 int32
	_ = v583
	var v588 int32
	_ = v588
	var v591 int32
	_ = v591
	var v594 int32
	_ = v594
	var v597 int32
	_ = v597
	var v602 int32
	_ = v602
	var v603 int32
	_ = v603
	var v607 int32
	_ = v607
	var v612 int32
	_ = v612
	var v615 int32
	_ = v615
	var v617 int32
	_ = v617
	var v622 int32
	_ = v622
	var v623 int32
	_ = v623
	var v627 int32
	_ = v627
	var v632 int32
	_ = v632
	var v635 int32
	_ = v635
	var v637 int32
	_ = v637
	var v642 int32
	_ = v642
	var v643 int32
	_ = v643
	var v647 int32
	_ = v647
	var v652 int32
	_ = v652
	var v654 int32
	_ = v654
	var v658 int32
	_ = v658
	var v659 int32
	_ = v659
	var v661 int32
	_ = v661
	var v664 int32
	_ = v664
	var v669 int32
	_ = v669
	var v679 int32
	_ = v679
	var v680 int32
	_ = v680
	var v683 int32
	_ = v683
	var v687 int32
	_ = v687
	var v690 int32
	_ = v690
	var v694 int32
	_ = v694
	var v698 int32
	_ = v698
	var v701 int32
	_ = v701
	var v706 int32
	_ = v706
	var v707 int32
	_ = v707
	var v711 int32
	_ = v711
	var v712 int32
	_ = v712
	var v716 int32
	_ = v716
	var v721 int32
	_ = v721
	var v722 int32
	_ = v722
	var v723 int32
	_ = v723
	var v725 int32
	_ = v725
	var v728 int32
	_ = v728
	var v730 int32
	_ = v730
	var v731 int64
	_ = v731
	var v755 int32
	_ = v755
	var v757 int32
	_ = v757
	var v759 int32
	_ = v759
	var v762 int32
	_ = v762
	var v763 int32
	_ = v763
	var v767 int32
	_ = v767
	var v772 int32
	_ = v772
	var v776 int32
	_ = v776
	var v781 int32
	_ = v781
	var v786 int32
	_ = v786
	var v788 int32
	_ = v788
	var v789 int32
	_ = v789
	var v799 int32
	_ = v799
	v10 = m.G0
	v12 = v10 - int32(128)
	m.G0 = v12
	v15 = *(*int32)(unsafe.Add(mBase, _consts[458]))
	if base.Ui32(int32(1)) < base.Ui32(v15-int32(3)) {
		v59 = v15
		goto L4
	} else {
		goto L5
	}
L1:
	;
	v481 = *(*int32)(unsafe.Add(mBase, _consts[458]))
	if v481 != int32(9) {
		v522 = v481
		goto L130
	} else {
		goto L131
	}
L2:
	;
	v410 = *(*int32)(unsafe.Add(mBase, _consts[458]))
	if v410 != int32(8) {
		goto L1
	} else {
		goto L111
	}
L3:
	;
	v65 = int32(130810)
	v69 = *(*int32)(unsafe.Add(mBase, _consts[459]))
	if int32(2) < v69 {
		goto L19
	} else {
		goto L20
	}
L4:
	;
	if base.Ui32(int32(1)) < base.Ui32(v59-int32(5)) {
		goto L2
	} else {
		goto L18
	}
L5:
	;
	v21 = int32(*(*uint8)(unsafe.Add(mBase, _consts[460])))
	if v21 == int32(0) {
		v59 = v15
		goto L4
	} else {
		goto L6
	}
L6:
	;
	v25 = F_CountChildren(m, int32(2))
	mBase = m.M
	v26 = m.ExcPending
	if v26 != 0 {
		goto L7
	} else {
		goto L8
	}
L7:
	;
	return
L8:
	;
	if v25 == int32(0) {
		goto L9
	} else {
		goto L10
	}
L9:
	;
	v31 = F_errstart(m, int32(14), int32(0))
	mBase = m.M
	v32 = m.ExcPending
	if v32 != 0 {
		goto L7
	} else {
		goto L12
	}
L10:
	;
	goto L11
L11:
	;
	v58 = *(*int32)(unsafe.Add(mBase, _consts[458]))
	v59 = v58
	goto L4
L12:
	;
	if v31 != 0 {
		goto L13
	} else {
		goto L14
	}
L13:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v12)+116)) = int32(544538)
	v36 = *(*int32)(unsafe.Add(mBase, _consts[458]))
	v41 = *(*int32)(unsafe.Add(mBase, uint32(v36<<(uint(int32(2))%32))+uint32(_consts[461])))
	*(*int32)(unsafe.Add(mBase, uint32(v12)+112)) = v41
	F_errmsg_internal(m, int32(191755), v12+int32(112))
	mBase = m.M
	v47 = m.ExcPending
	if v47 != 0 {
		goto L7
	} else {
		goto L16
	}
L14:
	;
	goto L15
L15:
	;
	v53 = int32(5)
	*(*int32)(unsafe.Add(mBase, _consts[458])) = v53
	v64 = v53
	goto L3
L16:
	;
	F_errfinish(m, int32(514840), int32(3272), int32(368732))
	mBase = m.M
	v52 = m.ExcPending
	if v52 != 0 {
		goto L7
	} else {
		goto L17
	}
L17:
	;
	goto L15
L18:
	;
	v64 = v59
	goto L3
L19:
	;
	v72 = v65
	goto L21
L20:
	;
	v72 = int32(124090)
	goto L21
L21:
	;
	v74 = int32(*(*uint8)(unsafe.Add(mBase, _consts[456])))
	if v74 != 0 {
		goto L22
	} else {
		goto L23
	}
L22:
	;
	v75 = v65
	goto L24
L23:
	;
	v75 = v72
	goto L24
L24:
	;
	if v64 == int32(5) {
		goto L25
	} else {
		goto L26
	}
L25:
	;
	v78 = m.G0
	v80 = v78 - int32(16)
	m.G0 = v80
	v83 = *(*int32)(unsafe.Add(mBase, _consts[457]))
	if v83 == int32(0) {
		goto L28
	} else {
		goto L29
	}
L26:
	;
	goto L27
L27:
	;
	v271 = F_CountChildren(m, v75)
	mBase = m.M
	v272 = m.ExcPending
	if v272 != 0 {
		goto L7
	} else {
		goto L69
	}
L28:
	;
	m.G0 = v80 + int32(16)
	v175 = *(*int32)(unsafe.Add(mBase, _consts[462]))
	if v175 == int32(0) {
		goto L48
	} else {
		goto L49
	}
L29:
	;
	if v83 == int32(4154516) {
		goto L28
	} else {
		goto L30
	}
L30:
	;
	v89 = *(*int32)(unsafe.Add(mBase, _consts[79]))
	v92 = v89
	v93 = v83
	goto L31
L31:
	;
	v100 = v93 + int32(4)
	v101 = *(*int32)(unsafe.Add(mBase, uint32(v100)))
	v104 = *(*int32)(unsafe.Add(mBase, uint32(v93-int32(8))))
	v107 = v92 + v104*int32(1480)
	v108 = *(*int32)(unsafe.Add(mBase, uint32(v107)+20))
	if v108 != int32(-1) {
		v156 = v92
		goto L33
	} else {
		goto L34
	}
L32:
	;
	goto L28
L33:
	;
	if v101 != int32(4154516) {
		v92 = v156
		v93 = v101
		goto L31
	} else {
		goto L47
	}
L34:
	;
	v113 = *(*int32)(unsafe.Add(mBase, uint32(v93-int32(24))))
	if v113 == int32(0) {
		v156 = v92
		goto L33
	} else {
		goto L35
	}
L35:
	;
	v116 = int32(16)
	v120 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v93-int32(1288)))))
	if v120&v116 != 0 {
		goto L36
	} else {
		goto L37
	}
L36:
	;
	v123 = *(*int32)(unsafe.Add(mBase, uint32(v92)+8))
	*(*int32)(unsafe.Add(mBase, uint32(v92)+8)) = v123 + int32(1)
	goto L38
L37:
	;
	goto L38
L38:
	;
	v128 = v93 - int32(1480)
	v129 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v107+v116))) = uint8(v129)
	v133 = F_errstart(m, int32(14), v129)
	mBase = m.M
	v134 = m.ExcPending
	if v134 != 0 {
		goto L7
	} else {
		goto L39
	}
L39:
	;
	if v133 != 0 {
		goto L40
	} else {
		goto L41
	}
L40:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v80))) = v128
	F_errmsg_internal(m, int32(729506), v80)
	mBase = m.M
	v138 = m.ExcPending
	if v138 != 0 {
		goto L7
	} else {
		goto L43
	}
L41:
	;
	goto L42
L42:
	;
	v144 = *(*int32)(unsafe.Add(mBase, uint32(v93)))
	v145 = *(*int32)(unsafe.Add(mBase, uint32(v100)))
	*(*int32)(unsafe.Add(mBase, uint32(v144)+4)) = v145
	v147 = *(*int32)(unsafe.Add(mBase, uint32(v93)))
	*(*int32)(unsafe.Add(mBase, uint32(v145))) = v147
	F_pfree(m, v128)
	mBase = m.M
	v150 = m.ExcPending
	if v150 != 0 {
		goto L7
	} else {
		goto L45
	}
L43:
	;
	F_errfinish(m, int32(515039), int32(449), int32(230590))
	mBase = m.M
	v143 = m.ExcPending
	if v143 != 0 {
		goto L7
	} else {
		goto L44
	}
L44:
	;
	goto L42
L45:
	;
	v152 = F_kill(m, v113, int32(10))
	mBase = m.M
	v153 = m.ExcPending
	if v153 != 0 {
		goto L7
	} else {
		goto L46
	}
L46:
	;
	v155 = *(*int32)(unsafe.Add(mBase, _consts[79]))
	v156 = v155
	goto L33
L47:
	;
	goto L32
L48:
	;
	v237 = F_errstart(m, int32(14), int32(0))
	mBase = m.M
	v238 = m.ExcPending
	if v238 != 0 {
		goto L7
	} else {
		goto L63
	}
L49:
	;
	if v175 == int32(4457200) {
		goto L48
	} else {
		goto L50
	}
L50:
	;
	v182 = v175
	goto L51
L51:
	;
	if v75&int32(64) != 0 {
		goto L53
	} else {
		goto L54
	}
L52:
	;
	goto L48
L53:
	;
	v214 = *(*int32)(unsafe.Add(mBase, uint32(v182-int32(12))))
	if int32(base.Ui32(v75)>>(uint(v214)%32))&int32(1) != 0 {
		goto L58
	} else {
		goto L59
	}
L54:
	;
	v192 = v182 - int32(12)
	v193 = *(*int32)(unsafe.Add(mBase, uint32(v192)))
	if v193 != int32(1) {
		goto L53
	} else {
		goto L55
	}
L55:
	;
	v198 = *(*int32)(unsafe.Add(mBase, uint32(v182-int32(16))))
	v200 = *(*int32)(unsafe.Add(mBase, _consts[463]))
	v204 = *(*int32)(unsafe.Add(mBase, uint32(v200+v198<<(uint(int32(2))%32))+44))
	goto L56
L56:
	;
	if base.B2i32(v204 == int32(3)) == int32(0) {
		goto L53
	} else {
		goto L57
	}
L57:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v192))) = int32(6)
	goto L53
L58:
	;
	F_signal_child(m, v182-int32(20), int32(15))
	mBase = m.M
	v222 = m.ExcPending
	if v222 != 0 {
		goto L7
	} else {
		goto L61
	}
L59:
	;
	goto L60
L60:
	;
	v223 = *(*int32)(unsafe.Add(mBase, uint32(v182)+4))
	if v223 != int32(4457200) {
		v182 = v223
		goto L51
	} else {
		goto L62
	}
L61:
	;
	goto L60
L62:
	;
	goto L52
L63:
	;
	if v237 != 0 {
		goto L64
	} else {
		goto L65
	}
L64:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v12)+100)) = int32(544521)
	v242 = *(*int32)(unsafe.Add(mBase, _consts[458]))
	v247 = *(*int32)(unsafe.Add(mBase, uint32(v242<<(uint(int32(2))%32))+uint32(_consts[461])))
	*(*int32)(unsafe.Add(mBase, uint32(v12)+96)) = v247
	F_errmsg_internal(m, int32(191755), v12+int32(96))
	mBase = m.M
	v253 = m.ExcPending
	if v253 != 0 {
		goto L7
	} else {
		goto L67
	}
L65:
	;
	goto L66
L66:
	;
	*(*int32)(unsafe.Add(mBase, _consts[458])) = int32(6)
	goto L27
L67:
	;
	F_errfinish(m, int32(514840), int32(3272), int32(368732))
	mBase = m.M
	v258 = m.ExcPending
	if v258 != 0 {
		goto L7
	} else {
		goto L68
	}
L68:
	;
	goto L66
L69:
	;
	if v271 != 0 {
		goto L2
	} else {
		goto L70
	}
L70:
	;
	v274 = *(*int32)(unsafe.Add(mBase, _consts[459]))
	if v274 <= int32(2) {
		goto L72
	} else {
		goto L73
	}
L71:
	;
	v356 = *(*int32)(unsafe.Add(mBase, _consts[464]))
	if v356 == int32(0) {
		goto L98
	} else {
		goto L99
	}
L72:
	;
	v278 = int32(*(*uint8)(unsafe.Add(mBase, _consts[456])))
	if v278 == int32(0) {
		goto L71
	} else {
		goto L75
	}
L73:
	;
	goto L74
L74:
	;
	v283 = F_errstart(m, int32(14), int32(0))
	mBase = m.M
	v284 = m.ExcPending
	if v284 != 0 {
		goto L7
	} else {
		goto L76
	}
L75:
	;
	goto L74
L76:
	;
	if v283 != 0 {
		goto L77
	} else {
		goto L78
	}
L77:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v12)+68)) = int32(563602)
	v288 = *(*int32)(unsafe.Add(mBase, _consts[458]))
	v293 = *(*int32)(unsafe.Add(mBase, uint32(v288<<(uint(int32(2))%32))+uint32(_consts[461])))
	*(*int32)(unsafe.Add(mBase, uint32(v12)+64)) = v293
	F_errmsg_internal(m, int32(191755), v12-int32(-64))
	mBase = m.M
	v299 = m.ExcPending
	if v299 != 0 {
		goto L7
	} else {
		goto L80
	}
L78:
	;
	goto L79
L79:
	;
	*(*int32)(unsafe.Add(mBase, _consts[458])) = int32(11)
	v309 = *(*int32)(unsafe.Add(mBase, _consts[465]))
	if v309 != 0 {
		goto L82
	} else {
		goto L83
	}
L80:
	;
	F_errfinish(m, int32(514840), int32(3272), int32(368732))
	mBase = m.M
	v304 = m.ExcPending
	if v304 != 0 {
		goto L7
	} else {
		goto L81
	}
L81:
	;
	goto L79
L82:
	;
	F_FreeWaitEventSet(m, v309)
	mBase = m.M
	v311 = m.ExcPending
	if v311 != 0 {
		goto L7
	} else {
		goto L85
	}
L83:
	;
	goto L84
L84:
	;
	v312 = int32(4457268)
	v313 = int32(0)
	*(*int32)(unsafe.Add(mBase, _consts[465])) = v313
	v318 = F_CreateWaitEventSet(m, v313, int32(1))
	mBase = m.M
	v319 = m.ExcPending
	if v319 != 0 {
		goto L7
	} else {
		goto L86
	}
L85:
	;
	goto L84
L86:
	;
	*(*int32)(unsafe.Add(mBase, _consts[465])) = v318
	v324 = *(*int32)(unsafe.Add(mBase, _consts[77]))
	F_AddWaitEventToSet(m, v318, int32(1), int32(-1), v324)
	mBase = m.M
	v326 = m.ExcPending
	if v326 != 0 {
		goto L7
	} else {
		goto L87
	}
L87:
	;
	v328 = *(*int32)(unsafe.Add(mBase, _consts[462]))
	if v328 == int32(0) {
		goto L2
	} else {
		goto L88
	}
L88:
	;
	if v328 == int32(4457200) {
		goto L2
	} else {
		goto L89
	}
L89:
	;
	v333 = v328
	goto L90
L90:
	;
	v344 = *(*int32)(unsafe.Add(mBase, uint32(v333-int32(12))))
	if v344 == int32(2) {
		goto L92
	} else {
		goto L93
	}
L91:
	;
	goto L2
L92:
	;
	F_signal_child(m, v333-int32(20), int32(3))
	mBase = m.M
	v351 = m.ExcPending
	if v351 != 0 {
		goto L7
	} else {
		goto L95
	}
L93:
	;
	goto L94
L94:
	;
	v352 = *(*int32)(unsafe.Add(mBase, uint32(v333)+4))
	if v352 != int32(4457200) {
		v333 = v352
		goto L90
	} else {
		goto L96
	}
L95:
	;
	goto L94
L96:
	;
	goto L91
L97:
	;
	F_HandleFatalError(m, int32(0))
	mBase = m.M
	v399 = m.ExcPending
	if v399 != 0 {
		goto L7
	} else {
		goto L110
	}
L98:
	;
	v361 = F_StartChildProcess(m, int32(11))
	mBase = m.M
	v362 = m.ExcPending
	if v362 != 0 {
		goto L7
	} else {
		goto L101
	}
L99:
	;
	v366 = v356
	goto L100
L100:
	;
	F_signal_child(m, v366, int32(2))
	mBase = m.M
	v369 = m.ExcPending
	if v369 != 0 {
		goto L7
	} else {
		goto L103
	}
L101:
	;
	*(*int32)(unsafe.Add(mBase, _consts[464])) = v361
	if v361 == int32(0) {
		goto L97
	} else {
		goto L102
	}
L102:
	;
	v366 = v361
	goto L100
L103:
	;
	v372 = F_errstart(m, int32(14), int32(0))
	mBase = m.M
	v373 = m.ExcPending
	if v373 != 0 {
		goto L7
	} else {
		goto L104
	}
L104:
	;
	if v372 != 0 {
		goto L105
	} else {
		goto L106
	}
L105:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v12)+84)) = int32(548839)
	v377 = *(*int32)(unsafe.Add(mBase, _consts[458]))
	v382 = *(*int32)(unsafe.Add(mBase, uint32(v377<<(uint(int32(2))%32))+uint32(_consts[461])))
	*(*int32)(unsafe.Add(mBase, uint32(v12)+80)) = v382
	F_errmsg_internal(m, int32(191755), v12+int32(80))
	mBase = m.M
	v388 = m.ExcPending
	if v388 != 0 {
		goto L7
	} else {
		goto L108
	}
L106:
	;
	goto L107
L107:
	;
	*(*int32)(unsafe.Add(mBase, _consts[458])) = int32(7)
	goto L1
L108:
	;
	F_errfinish(m, int32(514840), int32(3272), int32(368732))
	mBase = m.M
	v393 = m.ExcPending
	if v393 != 0 {
		goto L7
	} else {
		goto L109
	}
L109:
	;
	goto L107
L110:
	;
	goto L2
L111:
	;
	v414 = F_CountChildren(m, int32(124923))
	mBase = m.M
	v415 = m.ExcPending
	if v415 != 0 {
		goto L7
	} else {
		goto L112
	}
L112:
	;
	if v414 != 0 {
		goto L1
	} else {
		goto L113
	}
L113:
	;
	v418 = F_errstart(m, int32(14), int32(0))
	mBase = m.M
	v419 = m.ExcPending
	if v419 != 0 {
		goto L7
	} else {
		goto L114
	}
L114:
	;
	if v418 != 0 {
		goto L115
	} else {
		goto L116
	}
L115:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v12)+52)) = int32(543777)
	v423 = *(*int32)(unsafe.Add(mBase, _consts[458]))
	v428 = *(*int32)(unsafe.Add(mBase, uint32(v423<<(uint(int32(2))%32))+uint32(_consts[461])))
	*(*int32)(unsafe.Add(mBase, uint32(v12)+48)) = v428
	F_errmsg_internal(m, int32(191755), v12+int32(48))
	mBase = m.M
	v434 = m.ExcPending
	if v434 != 0 {
		goto L7
	} else {
		goto L118
	}
L116:
	;
	goto L117
L117:
	;
	*(*int32)(unsafe.Add(mBase, _consts[458])) = int32(9)
	v444 = *(*int32)(unsafe.Add(mBase, _consts[462]))
	if v444 == int32(0) {
		goto L1
	} else {
		goto L120
	}
L118:
	;
	F_errfinish(m, int32(514840), int32(3272), int32(368732))
	mBase = m.M
	v439 = m.ExcPending
	if v439 != 0 {
		goto L7
	} else {
		goto L119
	}
L119:
	;
	goto L117
L120:
	;
	if v444 == int32(4457200) {
		goto L1
	} else {
		goto L121
	}
L121:
	;
	v449 = v444
	goto L122
L122:
	;
	v458 = int32(12)
	v460 = *(*int32)(unsafe.Add(mBase, uint32(v449-v458)))
	if v460 == v458 {
		goto L124
	} else {
		goto L125
	}
L123:
	;
	goto L1
L124:
	;
	F_signal_child(m, v449-int32(20), int32(12))
	mBase = m.M
	v467 = m.ExcPending
	if v467 != 0 {
		goto L7
	} else {
		goto L127
	}
L125:
	;
	goto L126
L126:
	;
	v468 = *(*int32)(unsafe.Add(mBase, uint32(v449)+4))
	if v468 != int32(4457200) {
		v449 = v468
		goto L122
	} else {
		goto L128
	}
L127:
	;
	goto L126
L128:
	;
	goto L123
L129:
	;
	m.G0 = v12 + int32(128)
	return
L130:
	;
	if v522 == int32(11) {
		goto L144
	} else {
		goto L145
	}
L131:
	;
	v485 = *(*int32)(unsafe.Add(mBase, _consts[466]))
	if v485 != 0 {
		v522 = v481
		goto L130
	} else {
		goto L132
	}
L132:
	;
	v488 = F_errstart(m, int32(14), int32(0))
	mBase = m.M
	v489 = m.ExcPending
	if v489 != 0 {
		goto L7
	} else {
		goto L133
	}
L133:
	;
	if v488 != 0 {
		goto L134
	} else {
		goto L135
	}
L134:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v12)+36)) = int32(546222)
	v493 = *(*int32)(unsafe.Add(mBase, _consts[458]))
	v498 = *(*int32)(unsafe.Add(mBase, uint32(v493<<(uint(int32(2))%32))+uint32(_consts[461])))
	*(*int32)(unsafe.Add(mBase, uint32(v12)+32)) = v498
	F_errmsg_internal(m, int32(191755), v12+int32(32))
	mBase = m.M
	v504 = m.ExcPending
	if v504 != 0 {
		goto L7
	} else {
		goto L137
	}
L135:
	;
	goto L136
L136:
	;
	*(*int32)(unsafe.Add(mBase, _consts[458])) = int32(10)
	v514 = *(*int32)(unsafe.Add(mBase, _consts[464]))
	if v514 == int32(0) {
		goto L129
	} else {
		goto L139
	}
L137:
	;
	F_errfinish(m, int32(514840), int32(3272), int32(368732))
	mBase = m.M
	v509 = m.ExcPending
	if v509 != 0 {
		goto L7
	} else {
		goto L138
	}
L138:
	;
	goto L136
L139:
	;
	F_signal_child(m, v514, int32(12))
	mBase = m.M
	v519 = m.ExcPending
	if v519 != 0 {
		goto L7
	} else {
		goto L140
	}
L140:
	;
	v521 = *(*int32)(unsafe.Add(mBase, _consts[458]))
	v522 = v521
	goto L130
L141:
	;
	v597 = *(*int32)(unsafe.Add(mBase, _consts[467]))
	if v597 == int32(3) {
		goto L169
	} else {
		goto L170
	}
L142:
	;
	v575 = int32(*(*uint8)(unsafe.Add(mBase, _consts[456])))
	if v575 != 0 {
		goto L158
	} else {
		goto L159
	}
L143:
	;
	v544 = F_errstart(m, int32(14), int32(0))
	mBase = m.M
	v545 = m.ExcPending
	if v545 != 0 {
		goto L7
	} else {
		goto L151
	}
L144:
	;
	v526 = F_CountChildren(m, int32(131071))
	mBase = m.M
	v527 = m.ExcPending
	if v527 != 0 {
		goto L7
	} else {
		goto L147
	}
L145:
	;
	v532 = v522
	goto L146
L146:
	;
	v536 = *(*int32)(unsafe.Add(mBase, _consts[459]))
	if base.B2i32(v532 == int32(12))&base.B2i32(int32(0) < v536) != 0 {
		goto L142
	} else {
		goto L149
	}
L147:
	;
	if v526 == int32(0) {
		goto L143
	} else {
		goto L148
	}
L148:
	;
	v531 = *(*int32)(unsafe.Add(mBase, _consts[458]))
	v532 = v531
	goto L146
L149:
	;
	if v532 != int32(12) {
		goto L129
	} else {
		goto L150
	}
L150:
	;
	goto L141
L151:
	;
	if v544 != 0 {
		goto L152
	} else {
		goto L153
	}
L152:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v12)+20)) = int32(551614)
	v549 = *(*int32)(unsafe.Add(mBase, _consts[458]))
	v554 = *(*int32)(unsafe.Add(mBase, uint32(v549<<(uint(int32(2))%32))+uint32(_consts[461])))
	*(*int32)(unsafe.Add(mBase, uint32(v12)+16)) = v554
	F_errmsg_internal(m, int32(191755), v12+int32(16))
	mBase = m.M
	v560 = m.ExcPending
	if v560 != 0 {
		goto L7
	} else {
		goto L155
	}
L153:
	;
	goto L154
L154:
	;
	*(*int32)(unsafe.Add(mBase, _consts[458])) = int32(12)
	v570 = *(*int32)(unsafe.Add(mBase, _consts[459]))
	if v570 <= int32(0) {
		goto L141
	} else {
		goto L157
	}
L155:
	;
	F_errfinish(m, int32(514840), int32(3272), int32(368732))
	mBase = m.M
	v565 = m.ExcPending
	if v565 != 0 {
		goto L7
	} else {
		goto L156
	}
L156:
	;
	goto L154
L157:
	;
	goto L142
L158:
	;
	v578 = F_errstart(m, int32(15), int32(0))
	mBase = m.M
	v579 = m.ExcPending
	if v579 != 0 {
		goto L7
	} else {
		goto L161
	}
L159:
	;
	goto L160
L160:
	;
	F_ExitPostmaster(m, int32(0))
	mBase = m.M
	v594 = m.ExcPending
	if v594 != 0 {
		goto L7
	} else {
		goto L168
	}
L161:
	;
	if v578 != 0 {
		goto L162
	} else {
		goto L163
	}
L162:
	;
	F_errmsg(m, int32(253904), int32(0))
	mBase = m.M
	v583 = m.ExcPending
	if v583 != 0 {
		goto L7
	} else {
		goto L165
	}
L163:
	;
	goto L164
L164:
	;
	F_ExitPostmaster(m, int32(1))
	mBase = m.M
	v591 = m.ExcPending
	if v591 != 0 {
		goto L7
	} else {
		goto L167
	}
L165:
	;
	F_errfinish(m, int32(514840), int32(3161), int32(388769))
	mBase = m.M
	v588 = m.ExcPending
	if v588 != 0 {
		goto L7
	} else {
		goto L166
	}
L166:
	;
	goto L164
L167:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L168:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L169:
	;
	v602 = F_errstart(m, int32(15), int32(0))
	mBase = m.M
	v603 = m.ExcPending
	if v603 != 0 {
		goto L7
	} else {
		goto L172
	}
L170:
	;
	goto L171
L171:
	;
	v617 = int32(*(*uint8)(unsafe.Add(mBase, _consts[468])))
	if v617 == int32(0) {
		goto L179
	} else {
		goto L180
	}
L172:
	;
	if v602 != 0 {
		goto L173
	} else {
		goto L174
	}
L173:
	;
	F_errmsg(m, int32(378188), int32(0))
	mBase = m.M
	v607 = m.ExcPending
	if v607 != 0 {
		goto L7
	} else {
		goto L176
	}
L174:
	;
	goto L175
L175:
	;
	F_ExitPostmaster(m, int32(1))
	mBase = m.M
	v615 = m.ExcPending
	if v615 != 0 {
		goto L7
	} else {
		goto L178
	}
L176:
	;
	F_errfinish(m, int32(514840), int32(3187), int32(388769))
	mBase = m.M
	v612 = m.ExcPending
	if v612 != 0 {
		goto L7
	} else {
		goto L177
	}
L177:
	;
	goto L175
L178:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L179:
	;
	v622 = F_errstart(m, int32(15), int32(0))
	mBase = m.M
	v623 = m.ExcPending
	if v623 != 0 {
		goto L7
	} else {
		goto L182
	}
L180:
	;
	goto L181
L181:
	;
	v637 = int32(*(*uint8)(unsafe.Add(mBase, _consts[456])))
	if v637 == int32(0) {
		goto L129
	} else {
		goto L189
	}
L182:
	;
	if v622 != 0 {
		goto L183
	} else {
		goto L184
	}
L183:
	;
	F_errmsg(m, int32(352807), int32(0))
	mBase = m.M
	v627 = m.ExcPending
	if v627 != 0 {
		goto L7
	} else {
		goto L186
	}
L184:
	;
	goto L185
L185:
	;
	F_ExitPostmaster(m, int32(1))
	mBase = m.M
	v635 = m.ExcPending
	if v635 != 0 {
		goto L7
	} else {
		goto L188
	}
L186:
	;
	F_errfinish(m, int32(514840), int32(3193), int32(388769))
	mBase = m.M
	v632 = m.ExcPending
	if v632 != 0 {
		goto L7
	} else {
		goto L187
	}
L187:
	;
	goto L185
L188:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L189:
	;
	v642 = F_errstart(m, int32(15), int32(0))
	mBase = m.M
	v643 = m.ExcPending
	if v643 != 0 {
		goto L7
	} else {
		goto L190
	}
L190:
	;
	if v642 != 0 {
		goto L191
	} else {
		goto L192
	}
L191:
	;
	F_errmsg(m, int32(341969), int32(0))
	mBase = m.M
	v647 = m.ExcPending
	if v647 != 0 {
		goto L7
	} else {
		goto L194
	}
L192:
	;
	goto L193
L193:
	;
	v654 = int32(*(*uint8)(unsafe.Add(mBase, _consts[469])))
	if v654 == int32(1) {
		goto L196
	} else {
		goto L197
	}
L194:
	;
	F_errfinish(m, int32(514840), int32(3205), int32(388769))
	mBase = m.M
	v652 = m.ExcPending
	if v652 != 0 {
		goto L7
	} else {
		goto L195
	}
L195:
	;
	goto L193
L196:
	;
	F_RemovePgTempFiles(m)
	mBase = m.M
	v658 = m.ExcPending
	if v658 != 0 {
		goto L7
	} else {
		goto L199
	}
L197:
	;
	goto L198
L198:
	;
	v659 = m.G0
	v661 = v659 - int32(16)
	m.G0 = v661
	v664 = *(*int32)(unsafe.Add(mBase, _consts[457]))
	if v664 == int32(0) {
		goto L200
	} else {
		goto L201
	}
L199:
	;
	goto L198
L200:
	;
	m.G0 = v661 + int32(16)
	F_shmem_exit(m, int32(1))
	mBase = m.M
	v755 = m.ExcPending
	if v755 != 0 {
		goto L7
	} else {
		goto L220
	}
L201:
	;
	if v664 == int32(4154516) {
		goto L200
	} else {
		goto L202
	}
L202:
	;
	v669 = v664
	goto L203
L203:
	;
	v679 = v669 + int32(4)
	v680 = *(*int32)(unsafe.Add(mBase, uint32(v679)))
	v683 = *(*int32)(unsafe.Add(mBase, uint32(v669-int32(1280))))
	if v683 == int32(-1) {
		goto L206
	} else {
		goto L207
	}
L204:
	;
	goto L200
L205:
	;
	if v680 != int32(4154516) {
		v669 = v680
		goto L203
	} else {
		goto L219
	}
L206:
	;
	v687 = *(*int32)(unsafe.Add(mBase, _consts[79]))
	v690 = *(*int32)(unsafe.Add(mBase, uint32(v669-int32(8))))
	v694 = int32(16)
	v698 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v669-int32(1288)))))
	if v698&v694 != 0 {
		goto L209
	} else {
		goto L210
	}
L207:
	;
	goto L208
L208:
	;
	v730 = v669 - int32(24)
	v731 = int64(0)
	*(*int64)(unsafe.Add(mBase, uint32(v730))) = v731
	*(*int64)(unsafe.Add(mBase, uint32(v730)+8)) = v731
	goto L205
L209:
	;
	v701 = *(*int32)(unsafe.Add(mBase, uint32(v687)+8))
	*(*int32)(unsafe.Add(mBase, uint32(v687)+8)) = v701 + int32(1)
	goto L211
L210:
	;
	goto L211
L211:
	;
	v706 = v669 - int32(1480)
	v707 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v687+v690*int32(1480)+v694))) = uint8(v707)
	v711 = F_errstart(m, int32(14), v707)
	mBase = m.M
	v712 = m.ExcPending
	if v712 != 0 {
		goto L7
	} else {
		goto L212
	}
L212:
	;
	if v711 != 0 {
		goto L213
	} else {
		goto L214
	}
L213:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v661))) = v706
	F_errmsg_internal(m, int32(729506), v661)
	mBase = m.M
	v716 = m.ExcPending
	if v716 != 0 {
		goto L7
	} else {
		goto L216
	}
L214:
	;
	goto L215
L215:
	;
	v722 = *(*int32)(unsafe.Add(mBase, uint32(v669)))
	v723 = *(*int32)(unsafe.Add(mBase, uint32(v679)))
	*(*int32)(unsafe.Add(mBase, uint32(v722)+4)) = v723
	v725 = *(*int32)(unsafe.Add(mBase, uint32(v669)))
	*(*int32)(unsafe.Add(mBase, uint32(v723))) = v725
	F_pfree(m, v706)
	mBase = m.M
	v728 = m.ExcPending
	if v728 != 0 {
		goto L7
	} else {
		goto L218
	}
L216:
	;
	F_errfinish(m, int32(515039), int32(449), int32(230590))
	mBase = m.M
	v721 = m.ExcPending
	if v721 != 0 {
		goto L7
	} else {
		goto L217
	}
L217:
	;
	goto L215
L218:
	;
	goto L205
L219:
	;
	goto L204
L220:
	;
	F_LocalProcessControlFile(m)
	mBase = m.M
	v757 = m.ExcPending
	if v757 != 0 {
		goto L7
	} else {
		goto L221
	}
L221:
	;
	F_CreateSharedMemoryAndSemaphores(m)
	mBase = m.M
	v759 = m.ExcPending
	if v759 != 0 {
		goto L7
	} else {
		goto L222
	}
L222:
	;
	v762 = F_errstart(m, int32(14), int32(0))
	mBase = m.M
	v763 = m.ExcPending
	if v763 != 0 {
		goto L7
	} else {
		goto L223
	}
L223:
	;
	if v762 != 0 {
		goto L224
	} else {
		goto L225
	}
L224:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v12)+4)) = int32(547026)
	v767 = *(*int32)(unsafe.Add(mBase, _consts[458]))
	v772 = *(*int32)(unsafe.Add(mBase, uint32(v767<<(uint(int32(2))%32))+uint32(_consts[461])))
	*(*int32)(unsafe.Add(mBase, uint32(v12))) = v772
	F_errmsg_internal(m, int32(191755), v12)
	mBase = m.M
	v776 = m.ExcPending
	if v776 != 0 {
		goto L7
	} else {
		goto L227
	}
L225:
	;
	goto L226
L226:
	;
	*(*int32)(unsafe.Add(mBase, _consts[458])) = int32(1)
	F_maybe_adjust_io_workers(m)
	mBase = m.M
	v786 = m.ExcPending
	if v786 != 0 {
		goto L7
	} else {
		goto L229
	}
L227:
	;
	F_errfinish(m, int32(514840), int32(3272), int32(368732))
	mBase = m.M
	v781 = m.ExcPending
	if v781 != 0 {
		goto L7
	} else {
		goto L228
	}
L228:
	;
	goto L226
L229:
	;
	v788 = F_StartChildProcess(m, int32(13))
	mBase = m.M
	v789 = m.ExcPending
	if v789 != 0 {
		goto L7
	} else {
		goto L230
	}
L230:
	;
	*(*int32)(unsafe.Add(mBase, _consts[467])) = int32(1)
	*(*int32)(unsafe.Add(mBase, _consts[470])) = v788
	*(*int64)(unsafe.Add(mBase, _consts[471])) = int64(0)
	F_ConfigurePostmasterWaitSet(m)
	mBase = m.M
	v799 = m.ExcPending
	if v799 != 0 {
		goto L7
	} else {
		goto L231
	}
L231:
	;
	goto L129
}
