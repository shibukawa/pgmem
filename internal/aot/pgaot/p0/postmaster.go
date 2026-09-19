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
	v4 = *(*int32)(unsafe.Add(mBase, _c_F_CheckPostmasterSignal[0]))
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
	var v9 int32
	_ = v9
	var v11 int32
	_ = v11
	var v14 int32
	_ = v14
	var v20 int32
	_ = v20
	var v26 int32
	_ = v26
	var v27 int32
	_ = v27
	var v32 int32
	_ = v32
	var v33 int32
	_ = v33
	var v37 int32
	_ = v37
	var v42 int32
	_ = v42
	var v48 int32
	_ = v48
	var v53 int32
	_ = v53
	var v54 int32
	_ = v54
	var v59 int32
	_ = v59
	var v60 int32
	_ = v60
	var v65 int32
	_ = v65
	var v66 int32
	_ = v66
	var v70 int32
	_ = v70
	var v73 int32
	_ = v73
	var v75 int32
	_ = v75
	var v76 int32
	_ = v76
	var v79 int32
	_ = v79
	var v81 int32
	_ = v81
	var v84 int32
	_ = v84
	var v85 int32
	_ = v85
	var v93 int32
	_ = v93
	var v96 int32
	_ = v96
	var v97 int32
	_ = v97
	var v102 int32
	_ = v102
	var v105 int32
	_ = v105
	var v108 int32
	_ = v108
	var v109 int32
	_ = v109
	var v114 int32
	_ = v114
	var v119 int32
	_ = v119
	var v122 int32
	_ = v122
	var v127 int32
	_ = v127
	var v130 int32
	_ = v130
	var v134 int32
	_ = v134
	var v135 int32
	_ = v135
	var v139 int32
	_ = v139
	var v144 int32
	_ = v144
	var v145 int32
	_ = v145
	var v146 int32
	_ = v146
	var v148 int32
	_ = v148
	var v151 int32
	_ = v151
	var v153 int32
	_ = v153
	var v155 int32
	_ = v155
	var v156 int32
	_ = v156
	var v173 int32
	_ = v173
	var v174 int32
	_ = v174
	var v183 int32
	_ = v183
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
	var v236 int32
	_ = v236
	var v237 int32
	_ = v237
	var v241 int32
	_ = v241
	var v246 int32
	_ = v246
	var v252 int32
	_ = v252
	var v257 int32
	_ = v257
	var v269 int32
	_ = v269
	var v270 int32
	_ = v270
	var v272 int32
	_ = v272
	var v276 int32
	_ = v276
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
	var v334 int32
	_ = v334
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
	var v409 int32
	_ = v409
	var v413 int32
	_ = v413
	var v414 int32
	_ = v414
	var v417 int32
	_ = v417
	var v418 int32
	_ = v418
	var v422 int32
	_ = v422
	var v427 int32
	_ = v427
	var v433 int32
	_ = v433
	var v438 int32
	_ = v438
	var v443 int32
	_ = v443
	var v449 int32
	_ = v449
	var v457 int32
	_ = v457
	var v459 int32
	_ = v459
	var v466 int32
	_ = v466
	var v467 int32
	_ = v467
	var v479 int32
	_ = v479
	var v483 int32
	_ = v483
	var v486 int32
	_ = v486
	var v487 int32
	_ = v487
	var v491 int32
	_ = v491
	var v496 int32
	_ = v496
	var v502 int32
	_ = v502
	var v507 int32
	_ = v507
	var v512 int32
	_ = v512
	var v517 int32
	_ = v517
	var v519 int32
	_ = v519
	var v520 int32
	_ = v520
	var v525 int32
	_ = v525
	var v526 int32
	_ = v526
	var v530 int32
	_ = v530
	var v531 int32
	_ = v531
	var v535 int32
	_ = v535
	var v543 int32
	_ = v543
	var v544 int32
	_ = v544
	var v548 int32
	_ = v548
	var v553 int32
	_ = v553
	var v559 int32
	_ = v559
	var v564 int32
	_ = v564
	var v569 int32
	_ = v569
	var v574 int32
	_ = v574
	var v577 int32
	_ = v577
	var v578 int32
	_ = v578
	var v582 int32
	_ = v582
	var v587 int32
	_ = v587
	var v590 int32
	_ = v590
	var v593 int32
	_ = v593
	var v596 int32
	_ = v596
	var v601 int32
	_ = v601
	var v602 int32
	_ = v602
	var v606 int32
	_ = v606
	var v611 int32
	_ = v611
	var v614 int32
	_ = v614
	var v616 int32
	_ = v616
	var v621 int32
	_ = v621
	var v622 int32
	_ = v622
	var v626 int32
	_ = v626
	var v631 int32
	_ = v631
	var v634 int32
	_ = v634
	var v636 int32
	_ = v636
	var v641 int32
	_ = v641
	var v642 int32
	_ = v642
	var v646 int32
	_ = v646
	var v651 int32
	_ = v651
	var v653 int32
	_ = v653
	var v657 int32
	_ = v657
	var v658 int32
	_ = v658
	var v660 int32
	_ = v660
	var v663 int32
	_ = v663
	var v664 int32
	_ = v664
	var v671 int32
	_ = v671
	var v679 int32
	_ = v679
	var v682 int32
	_ = v682
	var v686 int32
	_ = v686
	var v689 int32
	_ = v689
	var v693 int32
	_ = v693
	var v697 int32
	_ = v697
	var v700 int32
	_ = v700
	var v705 int32
	_ = v705
	var v706 int32
	_ = v706
	var v710 int32
	_ = v710
	var v711 int32
	_ = v711
	var v715 int32
	_ = v715
	var v720 int32
	_ = v720
	var v721 int32
	_ = v721
	var v722 int32
	_ = v722
	var v724 int32
	_ = v724
	var v727 int32
	_ = v727
	var v729 int32
	_ = v729
	var v730 int64
	_ = v730
	var v752 int32
	_ = v752
	var v754 int32
	_ = v754
	var v756 int32
	_ = v756
	var v759 int32
	_ = v759
	var v760 int32
	_ = v760
	var v764 int32
	_ = v764
	var v769 int32
	_ = v769
	var v773 int32
	_ = v773
	var v778 int32
	_ = v778
	var v783 int32
	_ = v783
	var v785 int32
	_ = v785
	var v786 int32
	_ = v786
	var v796 int32
	_ = v796
	v9 = m.G0
	v11 = v9 - int32(128)
	m.G0 = v11
	v14 = *(*int32)(unsafe.Add(mBase, _c_F_PostmasterStateMachine[0]))
	if base.Ui32(int32(1)) < base.Ui32(v14-int32(3)) {
		v60 = v14
		goto L4
	} else {
		goto L5
	}
L1:
	;
	v479 = *(*int32)(unsafe.Add(mBase, _c_F_PostmasterStateMachine[0]))
	if v479 != int32(9) {
		v520 = v479
		goto L127
	} else {
		goto L128
	}
L2:
	;
	v409 = *(*int32)(unsafe.Add(mBase, _c_F_PostmasterStateMachine[0]))
	if v409 != int32(8) {
		goto L1
	} else {
		goto L109
	}
L3:
	;
	v66 = int32(_a_F_PostmasterStateMachine_0)
	v70 = *(*int32)(unsafe.Add(mBase, _c_F_PostmasterStateMachine[1]))
	if int32(2) < v70 {
		goto L19
	} else {
		goto L20
	}
L4:
	;
	if base.Ui32(int32(1)) < base.Ui32(v60-int32(5)) {
		goto L2
	} else {
		goto L18
	}
L5:
	;
	v20 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_PostmasterStateMachine[2])))
	if v20&int32(1) == int32(0) {
		v60 = v14
		goto L4
	} else {
		goto L6
	}
L6:
	;
	v26 = F_CountChildren(m, int32(2))
	mBase = m.M
	v27 = m.ExcPending
	if v27 != 0 {
		goto L7
	} else {
		goto L8
	}
L7:
	;
	return
L8:
	;
	if v26 == int32(0) {
		goto L9
	} else {
		goto L10
	}
L9:
	;
	v32 = F_errstart(m, int32(14), int32(0))
	mBase = m.M
	v33 = m.ExcPending
	if v33 != 0 {
		goto L7
	} else {
		goto L12
	}
L10:
	;
	goto L11
L11:
	;
	v59 = *(*int32)(unsafe.Add(mBase, _c_F_PostmasterStateMachine[0]))
	v60 = v59
	goto L4
L12:
	;
	if v32 != 0 {
		goto L13
	} else {
		goto L14
	}
L13:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v11)+116)) = int32(_a_F_PostmasterStateMachine_1)
	v37 = *(*int32)(unsafe.Add(mBase, _c_F_PostmasterStateMachine[0]))
	v42 = *(*int32)(unsafe.Add(mBase, uint32(v37<<(uint(int32(2))%32))+uint32(_c_F_PostmasterStateMachine[3])))
	*(*int32)(unsafe.Add(mBase, uint32(v11)+112)) = v42
	F_errmsg_internal(m, int32(_a_F_PostmasterStateMachine_2), v11+int32(112))
	mBase = m.M
	v48 = m.ExcPending
	if v48 != 0 {
		goto L7
	} else {
		goto L16
	}
L14:
	;
	goto L15
L15:
	;
	v54 = int32(5)
	*(*int32)(unsafe.Add(mBase, _c_F_PostmasterStateMachine[0])) = v54
	v65 = v54
	goto L3
L16:
	;
	F_errfinish(m, int32(_a_F_PostmasterStateMachine_3), int32(3272), int32(_a_F_PostmasterStateMachine_4))
	mBase = m.M
	v53 = m.ExcPending
	if v53 != 0 {
		goto L7
	} else {
		goto L17
	}
L17:
	;
	goto L15
L18:
	;
	v65 = v60
	goto L3
L19:
	;
	v73 = v66
	goto L21
L20:
	;
	v73 = int32(_a_F_PostmasterStateMachine_5)
	goto L21
L21:
	;
	v75 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_PostmasterStateMachine[4])))
	if v75 != 0 {
		goto L22
	} else {
		goto L23
	}
L22:
	;
	v76 = v66
	goto L24
L23:
	;
	v76 = v73
	goto L24
L24:
	;
	if v65 == int32(5) {
		goto L25
	} else {
		goto L26
	}
L25:
	;
	v79 = m.G0
	v81 = v79 - int32(16)
	m.G0 = v81
	v84 = *(*int32)(unsafe.Add(mBase, _c_F_PostmasterStateMachine[5]))
	v85 = int32(0)
	if base.B2i32(v84 == v85)|base.B2i32(v84 == int32(_a_F_PostmasterStateMachine_6)) == v85 {
		goto L28
	} else {
		goto L29
	}
L26:
	;
	goto L27
L27:
	;
	v269 = F_CountChildren(m, v76)
	mBase = m.M
	v270 = m.ExcPending
	if v270 != 0 {
		goto L7
	} else {
		goto L68
	}
L28:
	;
	v93 = *(*int32)(unsafe.Add(mBase, _c_F_PostmasterStateMachine[6]))
	v96 = v84
	v97 = v93
	goto L31
L29:
	;
	goto L30
L30:
	;
	m.G0 = v81 + int32(16)
	v173 = *(*int32)(unsafe.Add(mBase, _c_F_PostmasterStateMachine[7]))
	v174 = int32(0)
	if base.B2i32(v173 == v174)|base.B2i32(v173 == int32(_a_F_PostmasterStateMachine_7)) == v174 {
		goto L47
	} else {
		goto L48
	}
L31:
	;
	v102 = *(*int32)(unsafe.Add(mBase, uint32(v96)+4))
	v105 = *(*int32)(unsafe.Add(mBase, uint32(v96-int32(8))))
	v108 = v97 + v105*int32(1480)
	v109 = *(*int32)(unsafe.Add(mBase, uint32(v108)+20))
	if v109 != int32(-1) {
		v156 = v97
		goto L33
	} else {
		goto L34
	}
L32:
	;
	goto L30
L33:
	;
	if v102 != int32(_a_F_PostmasterStateMachine_6) {
		v96 = v102
		v97 = v156
		goto L31
	} else {
		goto L46
	}
L34:
	;
	v114 = *(*int32)(unsafe.Add(mBase, uint32(v96-int32(24))))
	if v114 == int32(0) {
		v156 = v97
		goto L33
	} else {
		goto L35
	}
L35:
	;
	v119 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v96-int32(1288)))))
	if v119&int32(16) != 0 {
		goto L36
	} else {
		goto L37
	}
L36:
	;
	v122 = *(*int32)(unsafe.Add(mBase, uint32(v97)+8))
	*(*int32)(unsafe.Add(mBase, uint32(v97)+8)) = v122 + int32(1)
	goto L38
L37:
	;
	goto L38
L38:
	;
	v127 = v96 - int32(1480)
	v130 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v108+int32(16)))) = uint8(v130)
	v134 = F_errstart(m, int32(14), v130)
	mBase = m.M
	v135 = m.ExcPending
	if v135 != 0 {
		goto L7
	} else {
		goto L39
	}
L39:
	;
	if v134 != 0 {
		goto L40
	} else {
		goto L41
	}
L40:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v81))) = v127
	F_errmsg_internal(m, int32(_a_F_PostmasterStateMachine_8), v81)
	mBase = m.M
	v139 = m.ExcPending
	if v139 != 0 {
		goto L7
	} else {
		goto L43
	}
L41:
	;
	goto L42
L42:
	;
	v145 = *(*int32)(unsafe.Add(mBase, uint32(v96)))
	v146 = *(*int32)(unsafe.Add(mBase, uint32(v96)+4))
	*(*int32)(unsafe.Add(mBase, uint32(v145)+4)) = v146
	v148 = *(*int32)(unsafe.Add(mBase, uint32(v96)))
	*(*int32)(unsafe.Add(mBase, uint32(v146))) = v148
	F_pfree(m, v127)
	mBase = m.M
	v151 = m.ExcPending
	if v151 != 0 {
		goto L7
	} else {
		goto L45
	}
L43:
	;
	F_errfinish(m, int32(_a_F_PostmasterStateMachine_9), int32(449), int32(_a_F_PostmasterStateMachine_10))
	mBase = m.M
	v144 = m.ExcPending
	if v144 != 0 {
		goto L7
	} else {
		goto L44
	}
L44:
	;
	goto L42
L45:
	;
	v153 = F_pgmem_kill(m, v114, int32(10))
	mBase = m.M
	v155 = *(*int32)(unsafe.Add(mBase, _c_F_PostmasterStateMachine[6]))
	v156 = v155
	goto L33
L46:
	;
	goto L32
L47:
	;
	v183 = v173
	goto L50
L48:
	;
	goto L49
L49:
	;
	v236 = F_errstart(m, int32(14), int32(0))
	mBase = m.M
	v237 = m.ExcPending
	if v237 != 0 {
		goto L7
	} else {
		goto L62
	}
L50:
	;
	if v76&int32(64) != 0 {
		goto L52
	} else {
		goto L53
	}
L51:
	;
	goto L49
L52:
	;
	v214 = *(*int32)(unsafe.Add(mBase, uint32(v183-int32(12))))
	if int32(base.Ui32(v76)>>(uint(v214)%32))&int32(1) != 0 {
		goto L57
	} else {
		goto L58
	}
L53:
	;
	v192 = v183 - int32(12)
	v193 = *(*int32)(unsafe.Add(mBase, uint32(v192)))
	if v193 != int32(1) {
		goto L52
	} else {
		goto L54
	}
L54:
	;
	v198 = *(*int32)(unsafe.Add(mBase, uint32(v183-int32(16))))
	v200 = *(*int32)(unsafe.Add(mBase, _c_F_PostmasterStateMachine[8]))
	v204 = *(*int32)(unsafe.Add(mBase, uint32(v200+v198<<(uint(int32(2))%32))+44))
	goto L55
L55:
	;
	if base.B2i32(v204 == int32(3)) == int32(0) {
		goto L52
	} else {
		goto L56
	}
L56:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v192))) = int32(6)
	goto L52
L57:
	;
	F_signal_child(m, v183-int32(20), int32(15))
	mBase = m.M
	v222 = m.ExcPending
	if v222 != 0 {
		goto L7
	} else {
		goto L60
	}
L58:
	;
	goto L59
L59:
	;
	v223 = *(*int32)(unsafe.Add(mBase, uint32(v183)+4))
	if v223 != int32(_a_F_PostmasterStateMachine_7) {
		v183 = v223
		goto L50
	} else {
		goto L61
	}
L60:
	;
	goto L59
L61:
	;
	goto L51
L62:
	;
	if v236 != 0 {
		goto L63
	} else {
		goto L64
	}
L63:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v11)+100)) = int32(_a_F_PostmasterStateMachine_11)
	v241 = *(*int32)(unsafe.Add(mBase, _c_F_PostmasterStateMachine[0]))
	v246 = *(*int32)(unsafe.Add(mBase, uint32(v241<<(uint(int32(2))%32))+uint32(_c_F_PostmasterStateMachine[3])))
	*(*int32)(unsafe.Add(mBase, uint32(v11)+96)) = v246
	F_errmsg_internal(m, int32(_a_F_PostmasterStateMachine_2), v11+int32(96))
	mBase = m.M
	v252 = m.ExcPending
	if v252 != 0 {
		goto L7
	} else {
		goto L66
	}
L64:
	;
	goto L65
L65:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_PostmasterStateMachine[0])) = int32(6)
	goto L27
L66:
	;
	F_errfinish(m, int32(_a_F_PostmasterStateMachine_3), int32(3272), int32(_a_F_PostmasterStateMachine_4))
	mBase = m.M
	v257 = m.ExcPending
	if v257 != 0 {
		goto L7
	} else {
		goto L67
	}
L67:
	;
	goto L65
L68:
	;
	if v269 != 0 {
		goto L2
	} else {
		goto L69
	}
L69:
	;
	v272 = *(*int32)(unsafe.Add(mBase, _c_F_PostmasterStateMachine[1]))
	if v272 <= int32(2) {
		goto L71
	} else {
		goto L72
	}
L70:
	;
	v356 = *(*int32)(unsafe.Add(mBase, _c_F_PostmasterStateMachine[9]))
	if v356 == int32(0) {
		goto L96
	} else {
		goto L97
	}
L71:
	;
	v276 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_PostmasterStateMachine[4])))
	if v276&int32(1) == int32(0) {
		goto L70
	} else {
		goto L74
	}
L72:
	;
	goto L73
L73:
	;
	v283 = F_errstart(m, int32(14), int32(0))
	mBase = m.M
	v284 = m.ExcPending
	if v284 != 0 {
		goto L7
	} else {
		goto L75
	}
L74:
	;
	goto L73
L75:
	;
	if v283 != 0 {
		goto L76
	} else {
		goto L77
	}
L76:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v11)+68)) = int32(_a_F_PostmasterStateMachine_12)
	v288 = *(*int32)(unsafe.Add(mBase, _c_F_PostmasterStateMachine[0]))
	v293 = *(*int32)(unsafe.Add(mBase, uint32(v288<<(uint(int32(2))%32))+uint32(_c_F_PostmasterStateMachine[3])))
	*(*int32)(unsafe.Add(mBase, uint32(v11)+64)) = v293
	F_errmsg_internal(m, int32(_a_F_PostmasterStateMachine_2), v11-int32(-64))
	mBase = m.M
	v299 = m.ExcPending
	if v299 != 0 {
		goto L7
	} else {
		goto L79
	}
L77:
	;
	goto L78
L78:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_PostmasterStateMachine[0])) = int32(11)
	v309 = *(*int32)(unsafe.Add(mBase, _c_F_PostmasterStateMachine[10]))
	if v309 != 0 {
		goto L81
	} else {
		goto L82
	}
L79:
	;
	F_errfinish(m, int32(_a_F_PostmasterStateMachine_3), int32(3272), int32(_a_F_PostmasterStateMachine_4))
	mBase = m.M
	v304 = m.ExcPending
	if v304 != 0 {
		goto L7
	} else {
		goto L80
	}
L80:
	;
	goto L78
L81:
	;
	F_FreeWaitEventSet(m, v309)
	mBase = m.M
	v311 = m.ExcPending
	if v311 != 0 {
		goto L7
	} else {
		goto L84
	}
L82:
	;
	goto L83
L83:
	;
	v312 = int32(_a_F_PostmasterStateMachine_13)
	v313 = int32(0)
	*(*int32)(unsafe.Add(mBase, _c_F_PostmasterStateMachine[10])) = v313
	v318 = F_CreateWaitEventSet(m, v313, int32(1))
	mBase = m.M
	v319 = m.ExcPending
	if v319 != 0 {
		goto L7
	} else {
		goto L85
	}
L84:
	;
	goto L83
L85:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_PostmasterStateMachine[10])) = v318
	v324 = *(*int32)(unsafe.Add(mBase, _c_F_PostmasterStateMachine[11]))
	F_AddWaitEventToSet(m, v318, int32(1), int32(-1), v324)
	mBase = m.M
	v326 = m.ExcPending
	if v326 != 0 {
		goto L7
	} else {
		goto L86
	}
L86:
	;
	v328 = *(*int32)(unsafe.Add(mBase, _c_F_PostmasterStateMachine[7]))
	if base.B2i32(v328 == int32(0))|base.B2i32(v328 == int32(_a_F_PostmasterStateMachine_7)) != 0 {
		goto L2
	} else {
		goto L87
	}
L87:
	;
	v334 = v328
	goto L88
L88:
	;
	v344 = *(*int32)(unsafe.Add(mBase, uint32(v334-int32(12))))
	if v344 == int32(2) {
		goto L90
	} else {
		goto L91
	}
L89:
	;
	goto L2
L90:
	;
	F_signal_child(m, v334-int32(20), int32(3))
	mBase = m.M
	v351 = m.ExcPending
	if v351 != 0 {
		goto L7
	} else {
		goto L93
	}
L91:
	;
	goto L92
L92:
	;
	v352 = *(*int32)(unsafe.Add(mBase, uint32(v334)+4))
	if v352 != int32(_a_F_PostmasterStateMachine_7) {
		v334 = v352
		goto L88
	} else {
		goto L94
	}
L93:
	;
	goto L92
L94:
	;
	goto L89
L95:
	;
	F_HandleFatalError(m, int32(0))
	mBase = m.M
	v399 = m.ExcPending
	if v399 != 0 {
		goto L7
	} else {
		goto L108
	}
L96:
	;
	v361 = F_StartChildProcess(m, int32(11))
	mBase = m.M
	v362 = m.ExcPending
	if v362 != 0 {
		goto L7
	} else {
		goto L99
	}
L97:
	;
	v366 = v356
	goto L98
L98:
	;
	F_signal_child(m, v366, int32(2))
	mBase = m.M
	v369 = m.ExcPending
	if v369 != 0 {
		goto L7
	} else {
		goto L101
	}
L99:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_PostmasterStateMachine[9])) = v361
	if v361 == int32(0) {
		goto L95
	} else {
		goto L100
	}
L100:
	;
	v366 = v361
	goto L98
L101:
	;
	v372 = F_errstart(m, int32(14), int32(0))
	mBase = m.M
	v373 = m.ExcPending
	if v373 != 0 {
		goto L7
	} else {
		goto L102
	}
L102:
	;
	if v372 != 0 {
		goto L103
	} else {
		goto L104
	}
L103:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v11)+84)) = int32(_a_F_PostmasterStateMachine_14)
	v377 = *(*int32)(unsafe.Add(mBase, _c_F_PostmasterStateMachine[0]))
	v382 = *(*int32)(unsafe.Add(mBase, uint32(v377<<(uint(int32(2))%32))+uint32(_c_F_PostmasterStateMachine[3])))
	*(*int32)(unsafe.Add(mBase, uint32(v11)+80)) = v382
	F_errmsg_internal(m, int32(_a_F_PostmasterStateMachine_2), v11+int32(80))
	mBase = m.M
	v388 = m.ExcPending
	if v388 != 0 {
		goto L7
	} else {
		goto L106
	}
L104:
	;
	goto L105
L105:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_PostmasterStateMachine[0])) = int32(7)
	goto L1
L106:
	;
	F_errfinish(m, int32(_a_F_PostmasterStateMachine_3), int32(3272), int32(_a_F_PostmasterStateMachine_4))
	mBase = m.M
	v393 = m.ExcPending
	if v393 != 0 {
		goto L7
	} else {
		goto L107
	}
L107:
	;
	goto L105
L108:
	;
	goto L2
L109:
	;
	v413 = F_CountChildren(m, int32(_a_F_PostmasterStateMachine_15))
	mBase = m.M
	v414 = m.ExcPending
	if v414 != 0 {
		goto L7
	} else {
		goto L110
	}
L110:
	;
	if v413 != 0 {
		goto L1
	} else {
		goto L111
	}
L111:
	;
	v417 = F_errstart(m, int32(14), int32(0))
	mBase = m.M
	v418 = m.ExcPending
	if v418 != 0 {
		goto L7
	} else {
		goto L112
	}
L112:
	;
	if v417 != 0 {
		goto L113
	} else {
		goto L114
	}
L113:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v11)+52)) = int32(_a_F_PostmasterStateMachine_16)
	v422 = *(*int32)(unsafe.Add(mBase, _c_F_PostmasterStateMachine[0]))
	v427 = *(*int32)(unsafe.Add(mBase, uint32(v422<<(uint(int32(2))%32))+uint32(_c_F_PostmasterStateMachine[3])))
	*(*int32)(unsafe.Add(mBase, uint32(v11)+48)) = v427
	F_errmsg_internal(m, int32(_a_F_PostmasterStateMachine_2), v11+int32(48))
	mBase = m.M
	v433 = m.ExcPending
	if v433 != 0 {
		goto L7
	} else {
		goto L116
	}
L114:
	;
	goto L115
L115:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_PostmasterStateMachine[0])) = int32(9)
	v443 = *(*int32)(unsafe.Add(mBase, _c_F_PostmasterStateMachine[7]))
	if base.B2i32(v443 == int32(0))|base.B2i32(v443 == int32(_a_F_PostmasterStateMachine_7)) != 0 {
		goto L1
	} else {
		goto L118
	}
L116:
	;
	F_errfinish(m, int32(_a_F_PostmasterStateMachine_3), int32(3272), int32(_a_F_PostmasterStateMachine_4))
	mBase = m.M
	v438 = m.ExcPending
	if v438 != 0 {
		goto L7
	} else {
		goto L117
	}
L117:
	;
	goto L115
L118:
	;
	v449 = v443
	goto L119
L119:
	;
	v457 = int32(12)
	v459 = *(*int32)(unsafe.Add(mBase, uint32(v449-v457)))
	if v459 == v457 {
		goto L121
	} else {
		goto L122
	}
L120:
	;
	goto L1
L121:
	;
	F_signal_child(m, v449-int32(20), int32(12))
	mBase = m.M
	v466 = m.ExcPending
	if v466 != 0 {
		goto L7
	} else {
		goto L124
	}
L122:
	;
	goto L123
L123:
	;
	v467 = *(*int32)(unsafe.Add(mBase, uint32(v449)+4))
	if v467 != int32(_a_F_PostmasterStateMachine_7) {
		v449 = v467
		goto L119
	} else {
		goto L125
	}
L124:
	;
	goto L123
L125:
	;
	goto L120
L126:
	;
	m.G0 = v11 + int32(128)
	return
L127:
	;
	if v520 == int32(11) {
		goto L141
	} else {
		goto L142
	}
L128:
	;
	v483 = *(*int32)(unsafe.Add(mBase, _c_F_PostmasterStateMachine[12]))
	if v483 != 0 {
		v520 = v479
		goto L127
	} else {
		goto L129
	}
L129:
	;
	v486 = F_errstart(m, int32(14), int32(0))
	mBase = m.M
	v487 = m.ExcPending
	if v487 != 0 {
		goto L7
	} else {
		goto L130
	}
L130:
	;
	if v486 != 0 {
		goto L131
	} else {
		goto L132
	}
L131:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v11)+36)) = int32(_a_F_PostmasterStateMachine_17)
	v491 = *(*int32)(unsafe.Add(mBase, _c_F_PostmasterStateMachine[0]))
	v496 = *(*int32)(unsafe.Add(mBase, uint32(v491<<(uint(int32(2))%32))+uint32(_c_F_PostmasterStateMachine[3])))
	*(*int32)(unsafe.Add(mBase, uint32(v11)+32)) = v496
	F_errmsg_internal(m, int32(_a_F_PostmasterStateMachine_2), v11+int32(32))
	mBase = m.M
	v502 = m.ExcPending
	if v502 != 0 {
		goto L7
	} else {
		goto L134
	}
L132:
	;
	goto L133
L133:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_PostmasterStateMachine[0])) = int32(10)
	v512 = *(*int32)(unsafe.Add(mBase, _c_F_PostmasterStateMachine[9]))
	if v512 == int32(0) {
		goto L126
	} else {
		goto L136
	}
L134:
	;
	F_errfinish(m, int32(_a_F_PostmasterStateMachine_3), int32(3272), int32(_a_F_PostmasterStateMachine_4))
	mBase = m.M
	v507 = m.ExcPending
	if v507 != 0 {
		goto L7
	} else {
		goto L135
	}
L135:
	;
	goto L133
L136:
	;
	F_signal_child(m, v512, int32(12))
	mBase = m.M
	v517 = m.ExcPending
	if v517 != 0 {
		goto L7
	} else {
		goto L137
	}
L137:
	;
	v519 = *(*int32)(unsafe.Add(mBase, _c_F_PostmasterStateMachine[0]))
	v520 = v519
	goto L127
L138:
	;
	v596 = *(*int32)(unsafe.Add(mBase, _c_F_PostmasterStateMachine[13]))
	if v596 == int32(3) {
		goto L166
	} else {
		goto L167
	}
L139:
	;
	v574 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_PostmasterStateMachine[4])))
	if v574 != 0 {
		goto L155
	} else {
		goto L156
	}
L140:
	;
	v543 = F_errstart(m, int32(14), int32(0))
	mBase = m.M
	v544 = m.ExcPending
	if v544 != 0 {
		goto L7
	} else {
		goto L148
	}
L141:
	;
	v525 = F_CountChildren(m, int32(_a_F_PostmasterStateMachine_18))
	mBase = m.M
	v526 = m.ExcPending
	if v526 != 0 {
		goto L7
	} else {
		goto L144
	}
L142:
	;
	v531 = v520
	goto L143
L143:
	;
	v535 = *(*int32)(unsafe.Add(mBase, _c_F_PostmasterStateMachine[1]))
	if base.B2i32(v531 == int32(12))&base.B2i32(int32(0) < v535) != 0 {
		goto L139
	} else {
		goto L146
	}
L144:
	;
	if v525 == int32(0) {
		goto L140
	} else {
		goto L145
	}
L145:
	;
	v530 = *(*int32)(unsafe.Add(mBase, _c_F_PostmasterStateMachine[0]))
	v531 = v530
	goto L143
L146:
	;
	if v531 != int32(12) {
		goto L126
	} else {
		goto L147
	}
L147:
	;
	goto L138
L148:
	;
	if v543 != 0 {
		goto L149
	} else {
		goto L150
	}
L149:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v11)+20)) = int32(_a_F_PostmasterStateMachine_19)
	v548 = *(*int32)(unsafe.Add(mBase, _c_F_PostmasterStateMachine[0]))
	v553 = *(*int32)(unsafe.Add(mBase, uint32(v548<<(uint(int32(2))%32))+uint32(_c_F_PostmasterStateMachine[3])))
	*(*int32)(unsafe.Add(mBase, uint32(v11)+16)) = v553
	F_errmsg_internal(m, int32(_a_F_PostmasterStateMachine_2), v11+int32(16))
	mBase = m.M
	v559 = m.ExcPending
	if v559 != 0 {
		goto L7
	} else {
		goto L152
	}
L150:
	;
	goto L151
L151:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_PostmasterStateMachine[0])) = int32(12)
	v569 = *(*int32)(unsafe.Add(mBase, _c_F_PostmasterStateMachine[1]))
	if v569 <= int32(0) {
		goto L138
	} else {
		goto L154
	}
L152:
	;
	F_errfinish(m, int32(_a_F_PostmasterStateMachine_3), int32(3272), int32(_a_F_PostmasterStateMachine_4))
	mBase = m.M
	v564 = m.ExcPending
	if v564 != 0 {
		goto L7
	} else {
		goto L153
	}
L153:
	;
	goto L151
L154:
	;
	goto L139
L155:
	;
	v577 = F_errstart(m, int32(15), int32(0))
	mBase = m.M
	v578 = m.ExcPending
	if v578 != 0 {
		goto L7
	} else {
		goto L158
	}
L156:
	;
	goto L157
L157:
	;
	F_ExitPostmaster(m, int32(0))
	mBase = m.M
	v593 = m.ExcPending
	if v593 != 0 {
		goto L7
	} else {
		goto L165
	}
L158:
	;
	if v577 != 0 {
		goto L159
	} else {
		goto L160
	}
L159:
	;
	F_errmsg(m, int32(_a_F_PostmasterStateMachine_20), int32(0))
	mBase = m.M
	v582 = m.ExcPending
	if v582 != 0 {
		goto L7
	} else {
		goto L162
	}
L160:
	;
	goto L161
L161:
	;
	F_ExitPostmaster(m, int32(1))
	mBase = m.M
	v590 = m.ExcPending
	if v590 != 0 {
		goto L7
	} else {
		goto L164
	}
L162:
	;
	F_errfinish(m, int32(_a_F_PostmasterStateMachine_3), int32(3161), int32(_a_F_PostmasterStateMachine_21))
	mBase = m.M
	v587 = m.ExcPending
	if v587 != 0 {
		goto L7
	} else {
		goto L163
	}
L163:
	;
	goto L161
L164:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L165:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L166:
	;
	v601 = F_errstart(m, int32(15), int32(0))
	mBase = m.M
	v602 = m.ExcPending
	if v602 != 0 {
		goto L7
	} else {
		goto L169
	}
L167:
	;
	goto L168
L168:
	;
	v616 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_PostmasterStateMachine[14])))
	if v616 == int32(0) {
		goto L176
	} else {
		goto L177
	}
L169:
	;
	if v601 != 0 {
		goto L170
	} else {
		goto L171
	}
L170:
	;
	F_errmsg(m, int32(_a_F_PostmasterStateMachine_22), int32(0))
	mBase = m.M
	v606 = m.ExcPending
	if v606 != 0 {
		goto L7
	} else {
		goto L173
	}
L171:
	;
	goto L172
L172:
	;
	F_ExitPostmaster(m, int32(1))
	mBase = m.M
	v614 = m.ExcPending
	if v614 != 0 {
		goto L7
	} else {
		goto L175
	}
L173:
	;
	F_errfinish(m, int32(_a_F_PostmasterStateMachine_3), int32(3187), int32(_a_F_PostmasterStateMachine_21))
	mBase = m.M
	v611 = m.ExcPending
	if v611 != 0 {
		goto L7
	} else {
		goto L174
	}
L174:
	;
	goto L172
L175:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L176:
	;
	v621 = F_errstart(m, int32(15), int32(0))
	mBase = m.M
	v622 = m.ExcPending
	if v622 != 0 {
		goto L7
	} else {
		goto L179
	}
L177:
	;
	goto L178
L178:
	;
	v636 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_PostmasterStateMachine[4])))
	if v636 == int32(0) {
		goto L126
	} else {
		goto L186
	}
L179:
	;
	if v621 != 0 {
		goto L180
	} else {
		goto L181
	}
L180:
	;
	F_errmsg(m, int32(_a_F_PostmasterStateMachine_23), int32(0))
	mBase = m.M
	v626 = m.ExcPending
	if v626 != 0 {
		goto L7
	} else {
		goto L183
	}
L181:
	;
	goto L182
L182:
	;
	F_ExitPostmaster(m, int32(1))
	mBase = m.M
	v634 = m.ExcPending
	if v634 != 0 {
		goto L7
	} else {
		goto L185
	}
L183:
	;
	F_errfinish(m, int32(_a_F_PostmasterStateMachine_3), int32(3193), int32(_a_F_PostmasterStateMachine_21))
	mBase = m.M
	v631 = m.ExcPending
	if v631 != 0 {
		goto L7
	} else {
		goto L184
	}
L184:
	;
	goto L182
L185:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L186:
	;
	v641 = F_errstart(m, int32(15), int32(0))
	mBase = m.M
	v642 = m.ExcPending
	if v642 != 0 {
		goto L7
	} else {
		goto L187
	}
L187:
	;
	if v641 != 0 {
		goto L188
	} else {
		goto L189
	}
L188:
	;
	F_errmsg(m, int32(_a_F_PostmasterStateMachine_24), int32(0))
	mBase = m.M
	v646 = m.ExcPending
	if v646 != 0 {
		goto L7
	} else {
		goto L191
	}
L189:
	;
	goto L190
L190:
	;
	v653 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_PostmasterStateMachine[15])))
	if v653 == int32(1) {
		goto L193
	} else {
		goto L194
	}
L191:
	;
	F_errfinish(m, int32(_a_F_PostmasterStateMachine_3), int32(3205), int32(_a_F_PostmasterStateMachine_21))
	mBase = m.M
	v651 = m.ExcPending
	if v651 != 0 {
		goto L7
	} else {
		goto L192
	}
L192:
	;
	goto L190
L193:
	;
	F_RemovePgTempFiles(m)
	mBase = m.M
	v657 = m.ExcPending
	if v657 != 0 {
		goto L7
	} else {
		goto L196
	}
L194:
	;
	goto L195
L195:
	;
	v658 = m.G0
	v660 = v658 - int32(16)
	m.G0 = v660
	v663 = *(*int32)(unsafe.Add(mBase, _c_F_PostmasterStateMachine[5]))
	v664 = int32(0)
	if base.B2i32(v663 == v664)|base.B2i32(v663 == int32(_a_F_PostmasterStateMachine_6)) == v664 {
		goto L197
	} else {
		goto L198
	}
L196:
	;
	goto L195
L197:
	;
	v671 = v663
	goto L200
L198:
	;
	goto L199
L199:
	;
	m.G0 = v660 + int32(16)
	F_shmem_exit(m, int32(1))
	mBase = m.M
	v752 = m.ExcPending
	if v752 != 0 {
		goto L7
	} else {
		goto L217
	}
L200:
	;
	v679 = *(*int32)(unsafe.Add(mBase, uint32(v671)+4))
	v682 = *(*int32)(unsafe.Add(mBase, uint32(v671-int32(1280))))
	if v682 == int32(-1) {
		goto L203
	} else {
		goto L204
	}
L201:
	;
	goto L199
L202:
	;
	if v679 != int32(_a_F_PostmasterStateMachine_6) {
		v671 = v679
		goto L200
	} else {
		goto L216
	}
L203:
	;
	v686 = *(*int32)(unsafe.Add(mBase, _c_F_PostmasterStateMachine[6]))
	v689 = *(*int32)(unsafe.Add(mBase, uint32(v671-int32(8))))
	v693 = int32(16)
	v697 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v671-int32(1288)))))
	if v697&v693 != 0 {
		goto L206
	} else {
		goto L207
	}
L204:
	;
	goto L205
L205:
	;
	v729 = v671 - int32(24)
	v730 = int64(0)
	*(*int64)(unsafe.Add(mBase, uint32(v729)+8)) = v730
	*(*int64)(unsafe.Add(mBase, uint32(v729))) = v730
	goto L202
L206:
	;
	v700 = *(*int32)(unsafe.Add(mBase, uint32(v686)+8))
	*(*int32)(unsafe.Add(mBase, uint32(v686)+8)) = v700 + int32(1)
	goto L208
L207:
	;
	goto L208
L208:
	;
	v705 = v671 - int32(1480)
	v706 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v686+v689*int32(1480)+v693))) = uint8(v706)
	v710 = F_errstart(m, int32(14), v706)
	mBase = m.M
	v711 = m.ExcPending
	if v711 != 0 {
		goto L7
	} else {
		goto L209
	}
L209:
	;
	if v710 != 0 {
		goto L210
	} else {
		goto L211
	}
L210:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v660))) = v705
	F_errmsg_internal(m, int32(_a_F_PostmasterStateMachine_8), v660)
	mBase = m.M
	v715 = m.ExcPending
	if v715 != 0 {
		goto L7
	} else {
		goto L213
	}
L211:
	;
	goto L212
L212:
	;
	v721 = *(*int32)(unsafe.Add(mBase, uint32(v671)))
	v722 = *(*int32)(unsafe.Add(mBase, uint32(v671)+4))
	*(*int32)(unsafe.Add(mBase, uint32(v721)+4)) = v722
	v724 = *(*int32)(unsafe.Add(mBase, uint32(v671)))
	*(*int32)(unsafe.Add(mBase, uint32(v722))) = v724
	F_pfree(m, v705)
	mBase = m.M
	v727 = m.ExcPending
	if v727 != 0 {
		goto L7
	} else {
		goto L215
	}
L213:
	;
	F_errfinish(m, int32(_a_F_PostmasterStateMachine_9), int32(449), int32(_a_F_PostmasterStateMachine_10))
	mBase = m.M
	v720 = m.ExcPending
	if v720 != 0 {
		goto L7
	} else {
		goto L214
	}
L214:
	;
	goto L212
L215:
	;
	goto L202
L216:
	;
	goto L201
L217:
	;
	F_LocalProcessControlFile(m)
	mBase = m.M
	v754 = m.ExcPending
	if v754 != 0 {
		goto L7
	} else {
		goto L218
	}
L218:
	;
	F_CreateSharedMemoryAndSemaphores(m)
	mBase = m.M
	v756 = m.ExcPending
	if v756 != 0 {
		goto L7
	} else {
		goto L219
	}
L219:
	;
	v759 = F_errstart(m, int32(14), int32(0))
	mBase = m.M
	v760 = m.ExcPending
	if v760 != 0 {
		goto L7
	} else {
		goto L220
	}
L220:
	;
	if v759 != 0 {
		goto L221
	} else {
		goto L222
	}
L221:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v11)+4)) = int32(_a_F_PostmasterStateMachine_25)
	v764 = *(*int32)(unsafe.Add(mBase, _c_F_PostmasterStateMachine[0]))
	v769 = *(*int32)(unsafe.Add(mBase, uint32(v764<<(uint(int32(2))%32))+uint32(_c_F_PostmasterStateMachine[3])))
	*(*int32)(unsafe.Add(mBase, uint32(v11))) = v769
	F_errmsg_internal(m, int32(_a_F_PostmasterStateMachine_2), v11)
	mBase = m.M
	v773 = m.ExcPending
	if v773 != 0 {
		goto L7
	} else {
		goto L224
	}
L222:
	;
	goto L223
L223:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_PostmasterStateMachine[0])) = int32(1)
	F_maybe_adjust_io_workers(m)
	mBase = m.M
	v783 = m.ExcPending
	if v783 != 0 {
		goto L7
	} else {
		goto L226
	}
L224:
	;
	F_errfinish(m, int32(_a_F_PostmasterStateMachine_3), int32(3272), int32(_a_F_PostmasterStateMachine_4))
	mBase = m.M
	v778 = m.ExcPending
	if v778 != 0 {
		goto L7
	} else {
		goto L225
	}
L225:
	;
	goto L223
L226:
	;
	v785 = F_StartChildProcess(m, int32(13))
	mBase = m.M
	v786 = m.ExcPending
	if v786 != 0 {
		goto L7
	} else {
		goto L227
	}
L227:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_PostmasterStateMachine[13])) = int32(1)
	*(*int32)(unsafe.Add(mBase, _c_F_PostmasterStateMachine[16])) = v785
	*(*int64)(unsafe.Add(mBase, _c_F_PostmasterStateMachine[17])) = int64(0)
	F_ConfigurePostmasterWaitSet(m)
	mBase = m.M
	v796 = m.ExcPending
	if v796 != 0 {
		goto L7
	} else {
		goto L228
	}
L228:
	;
	goto L126
}
