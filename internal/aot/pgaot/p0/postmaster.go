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
	var v128 int32
	_ = v128
	var v131 int32
	_ = v131
	var v138 int32
	_ = v138
	var v139 int32
	_ = v139
	var v143 int32
	_ = v143
	var v148 int32
	_ = v148
	var v149 int32
	_ = v149
	var v150 int32
	_ = v150
	var v152 int32
	_ = v152
	var v155 int32
	_ = v155
	var v157 int32
	_ = v157
	var v159 int32
	_ = v159
	var v160 int32
	_ = v160
	var v177 int32
	_ = v177
	var v178 int32
	_ = v178
	var v187 int32
	_ = v187
	var v196 int32
	_ = v196
	var v197 int32
	_ = v197
	var v202 int32
	_ = v202
	var v204 int32
	_ = v204
	var v208 int32
	_ = v208
	var v218 int32
	_ = v218
	var v226 int32
	_ = v226
	var v227 int32
	_ = v227
	var v240 int32
	_ = v240
	var v241 int32
	_ = v241
	var v245 int32
	_ = v245
	var v250 int32
	_ = v250
	var v256 int32
	_ = v256
	var v261 int32
	_ = v261
	var v273 int32
	_ = v273
	var v274 int32
	_ = v274
	var v276 int32
	_ = v276
	var v280 int32
	_ = v280
	var v287 int32
	_ = v287
	var v288 int32
	_ = v288
	var v292 int32
	_ = v292
	var v297 int32
	_ = v297
	var v303 int32
	_ = v303
	var v308 int32
	_ = v308
	var v313 int32
	_ = v313
	var v315 int32
	_ = v315
	var v316 int32
	_ = v316
	var v317 int32
	_ = v317
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
	var v338 int32
	_ = v338
	var v348 int32
	_ = v348
	var v355 int32
	_ = v355
	var v356 int32
	_ = v356
	var v360 int32
	_ = v360
	var v365 int32
	_ = v365
	var v366 int32
	_ = v366
	var v370 int32
	_ = v370
	var v373 int32
	_ = v373
	var v376 int32
	_ = v376
	var v377 int32
	_ = v377
	var v381 int32
	_ = v381
	var v386 int32
	_ = v386
	var v392 int32
	_ = v392
	var v397 int32
	_ = v397
	var v403 int32
	_ = v403
	var v413 int32
	_ = v413
	var v417 int32
	_ = v417
	var v418 int32
	_ = v418
	var v421 int32
	_ = v421
	var v422 int32
	_ = v422
	var v426 int32
	_ = v426
	var v431 int32
	_ = v431
	var v437 int32
	_ = v437
	var v442 int32
	_ = v442
	var v447 int32
	_ = v447
	var v453 int32
	_ = v453
	var v461 int32
	_ = v461
	var v463 int32
	_ = v463
	var v470 int32
	_ = v470
	var v471 int32
	_ = v471
	var v483 int32
	_ = v483
	var v487 int32
	_ = v487
	var v490 int32
	_ = v490
	var v491 int32
	_ = v491
	var v495 int32
	_ = v495
	var v500 int32
	_ = v500
	var v506 int32
	_ = v506
	var v511 int32
	_ = v511
	var v516 int32
	_ = v516
	var v521 int32
	_ = v521
	var v523 int32
	_ = v523
	var v524 int32
	_ = v524
	var v529 int32
	_ = v529
	var v530 int32
	_ = v530
	var v534 int32
	_ = v534
	var v535 int32
	_ = v535
	var v539 int32
	_ = v539
	var v547 int32
	_ = v547
	var v548 int32
	_ = v548
	var v552 int32
	_ = v552
	var v557 int32
	_ = v557
	var v563 int32
	_ = v563
	var v568 int32
	_ = v568
	var v573 int32
	_ = v573
	var v578 int32
	_ = v578
	var v581 int32
	_ = v581
	var v582 int32
	_ = v582
	var v586 int32
	_ = v586
	var v591 int32
	_ = v591
	var v594 int32
	_ = v594
	var v597 int32
	_ = v597
	var v600 int32
	_ = v600
	var v605 int32
	_ = v605
	var v606 int32
	_ = v606
	var v610 int32
	_ = v610
	var v615 int32
	_ = v615
	var v618 int32
	_ = v618
	var v620 int32
	_ = v620
	var v625 int32
	_ = v625
	var v626 int32
	_ = v626
	var v630 int32
	_ = v630
	var v635 int32
	_ = v635
	var v638 int32
	_ = v638
	var v640 int32
	_ = v640
	var v645 int32
	_ = v645
	var v646 int32
	_ = v646
	var v650 int32
	_ = v650
	var v655 int32
	_ = v655
	var v657 int32
	_ = v657
	var v661 int32
	_ = v661
	var v662 int32
	_ = v662
	var v664 int32
	_ = v664
	var v667 int32
	_ = v667
	var v668 int32
	_ = v668
	var v675 int32
	_ = v675
	var v683 int32
	_ = v683
	var v686 int32
	_ = v686
	var v690 int32
	_ = v690
	var v693 int32
	_ = v693
	var v697 int32
	_ = v697
	var v701 int32
	_ = v701
	var v704 int32
	_ = v704
	var v709 int32
	_ = v709
	var v710 int32
	_ = v710
	var v713 int32
	_ = v713
	var v718 int32
	_ = v718
	var v719 int32
	_ = v719
	var v723 int32
	_ = v723
	var v728 int32
	_ = v728
	var v729 int32
	_ = v729
	var v730 int32
	_ = v730
	var v732 int32
	_ = v732
	var v735 int32
	_ = v735
	var v738 int32
	_ = v738
	var v765 int32
	_ = v765
	var v767 int32
	_ = v767
	var v769 int32
	_ = v769
	var v775 int32
	_ = v775
	var v777 int32
	_ = v777
	var v780 int32
	_ = v780
	var v781 int32
	_ = v781
	var v785 int32
	_ = v785
	var v790 int32
	_ = v790
	var v794 int32
	_ = v794
	var v799 int32
	_ = v799
	var v804 int32
	_ = v804
	var v806 int32
	_ = v806
	var v807 int32
	_ = v807
	var v817 int32
	_ = v817
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
	v483 = *(*int32)(unsafe.Add(mBase, _c_F_PostmasterStateMachine[0]))
	if v483 != int32(9) {
		v524 = v483
		goto L127
	} else {
		goto L128
	}
L2:
	;
	v413 = *(*int32)(unsafe.Add(mBase, _c_F_PostmasterStateMachine[0]))
	if v413 != int32(8) {
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
	F_errfinish(m, int32(_a_F_PostmasterStateMachine_3), int32(3320), int32(_a_F_PostmasterStateMachine_4))
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
	v273 = F_CountChildren(m, v76)
	mBase = m.M
	v274 = m.ExcPending
	if v274 != 0 {
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
	v177 = *(*int32)(unsafe.Add(mBase, _c_F_PostmasterStateMachine[7]))
	v178 = int32(0)
	if base.B2i32(v177 == v178)|base.B2i32(v177 == int32(_a_F_PostmasterStateMachine_7)) == v178 {
		goto L47
	} else {
		goto L48
	}
L31:
	;
	v102 = *(*int32)(unsafe.Add(mBase, uint32(v96)+4))
	v105 = *(*int32)(unsafe.Add(mBase, uint32(v96-int32(8))))
	v108 = v97 + v105*int32(1488)
	v109 = *(*int32)(unsafe.Add(mBase, uint32(v108)+20))
	if v109 != int32(-1) {
		v160 = v97
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
		v97 = v160
		goto L31
	} else {
		goto L46
	}
L34:
	;
	v114 = *(*int32)(unsafe.Add(mBase, uint32(v96-int32(32))))
	if v114 == int32(0) {
		v160 = v97
		goto L33
	} else {
		goto L35
	}
L35:
	;
	v119 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v96-int32(1304)))))
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
	v127 = v96 - int32(1496)
	v128 = int32(0)
	v131 = base.AtomicRmwOr32(m, v128, int32(_a_F_PostmasterStateMachine_8), v128)
	*(*uint8)(unsafe.Add(mBase, uint32(v108+int32(16)))) = uint8(v128)
	v138 = F_errstart(m, int32(14), v128)
	mBase = m.M
	v139 = m.ExcPending
	if v139 != 0 {
		goto L7
	} else {
		goto L39
	}
L39:
	;
	if v138 != 0 {
		goto L40
	} else {
		goto L41
	}
L40:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v81))) = v127
	F_errmsg_internal(m, int32(_a_F_PostmasterStateMachine_9), v81)
	mBase = m.M
	v143 = m.ExcPending
	if v143 != 0 {
		goto L7
	} else {
		goto L43
	}
L41:
	;
	goto L42
L42:
	;
	v149 = *(*int32)(unsafe.Add(mBase, uint32(v96)))
	v150 = *(*int32)(unsafe.Add(mBase, uint32(v96)+4))
	*(*int32)(unsafe.Add(mBase, uint32(v149)+4)) = v150
	v152 = *(*int32)(unsafe.Add(mBase, uint32(v96)))
	*(*int32)(unsafe.Add(mBase, uint32(v150))) = v152
	F_pfree(m, v127)
	mBase = m.M
	v155 = m.ExcPending
	if v155 != 0 {
		goto L7
	} else {
		goto L45
	}
L43:
	;
	F_errfinish(m, int32(_a_F_PostmasterStateMachine_10), int32(466), int32(_a_F_PostmasterStateMachine_11))
	mBase = m.M
	v148 = m.ExcPending
	if v148 != 0 {
		goto L7
	} else {
		goto L44
	}
L44:
	;
	goto L42
L45:
	;
	v157 = F_pgmem_kill(m, v114, int32(10))
	mBase = m.M
	v159 = *(*int32)(unsafe.Add(mBase, _c_F_PostmasterStateMachine[6]))
	v160 = v159
	goto L33
L46:
	;
	goto L32
L47:
	;
	v187 = v177
	goto L50
L48:
	;
	goto L49
L49:
	;
	v240 = F_errstart(m, int32(14), int32(0))
	mBase = m.M
	v241 = m.ExcPending
	if v241 != 0 {
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
	v218 = *(*int32)(unsafe.Add(mBase, uint32(v187-int32(12))))
	if int32(base.Ui32(v76)>>(uint(v218)%32))&int32(1) != 0 {
		goto L57
	} else {
		goto L58
	}
L53:
	;
	v196 = v187 - int32(12)
	v197 = *(*int32)(unsafe.Add(mBase, uint32(v196)))
	if v197 != int32(1) {
		goto L52
	} else {
		goto L54
	}
L54:
	;
	v202 = *(*int32)(unsafe.Add(mBase, uint32(v187-int32(16))))
	v204 = *(*int32)(unsafe.Add(mBase, _c_F_PostmasterStateMachine[8]))
	v208 = *(*int32)(unsafe.Add(mBase, uint32(v204+v202<<(uint(int32(2))%32))+48))
	goto L55
L55:
	;
	if base.B2i32(v208 == int32(3)) == int32(0) {
		goto L52
	} else {
		goto L56
	}
L56:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v196))) = int32(6)
	goto L52
L57:
	;
	F_signal_child(m, v187-int32(20), int32(15))
	mBase = m.M
	v226 = m.ExcPending
	if v226 != 0 {
		goto L7
	} else {
		goto L60
	}
L58:
	;
	goto L59
L59:
	;
	v227 = *(*int32)(unsafe.Add(mBase, uint32(v187)+4))
	if v227 != int32(_a_F_PostmasterStateMachine_7) {
		v187 = v227
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
	if v240 != 0 {
		goto L63
	} else {
		goto L64
	}
L63:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v11)+100)) = int32(_a_F_PostmasterStateMachine_12)
	v245 = *(*int32)(unsafe.Add(mBase, _c_F_PostmasterStateMachine[0]))
	v250 = *(*int32)(unsafe.Add(mBase, uint32(v245<<(uint(int32(2))%32))+uint32(_c_F_PostmasterStateMachine[3])))
	*(*int32)(unsafe.Add(mBase, uint32(v11)+96)) = v250
	F_errmsg_internal(m, int32(_a_F_PostmasterStateMachine_2), v11+int32(96))
	mBase = m.M
	v256 = m.ExcPending
	if v256 != 0 {
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
	F_errfinish(m, int32(_a_F_PostmasterStateMachine_3), int32(3320), int32(_a_F_PostmasterStateMachine_4))
	mBase = m.M
	v261 = m.ExcPending
	if v261 != 0 {
		goto L7
	} else {
		goto L67
	}
L67:
	;
	goto L65
L68:
	;
	if v273 != 0 {
		goto L2
	} else {
		goto L69
	}
L69:
	;
	v276 = *(*int32)(unsafe.Add(mBase, _c_F_PostmasterStateMachine[1]))
	if v276 <= int32(2) {
		goto L71
	} else {
		goto L72
	}
L70:
	;
	v360 = *(*int32)(unsafe.Add(mBase, _c_F_PostmasterStateMachine[9]))
	if v360 == int32(0) {
		goto L96
	} else {
		goto L97
	}
L71:
	;
	v280 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_PostmasterStateMachine[4])))
	if v280&int32(1) == int32(0) {
		goto L70
	} else {
		goto L74
	}
L72:
	;
	goto L73
L73:
	;
	v287 = F_errstart(m, int32(14), int32(0))
	mBase = m.M
	v288 = m.ExcPending
	if v288 != 0 {
		goto L7
	} else {
		goto L75
	}
L74:
	;
	goto L73
L75:
	;
	if v287 != 0 {
		goto L76
	} else {
		goto L77
	}
L76:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v11)+68)) = int32(_a_F_PostmasterStateMachine_13)
	v292 = *(*int32)(unsafe.Add(mBase, _c_F_PostmasterStateMachine[0]))
	v297 = *(*int32)(unsafe.Add(mBase, uint32(v292<<(uint(int32(2))%32))+uint32(_c_F_PostmasterStateMachine[3])))
	*(*int32)(unsafe.Add(mBase, uint32(v11)+64)) = v297
	F_errmsg_internal(m, int32(_a_F_PostmasterStateMachine_2), v11-int32(-64))
	mBase = m.M
	v303 = m.ExcPending
	if v303 != 0 {
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
	v313 = *(*int32)(unsafe.Add(mBase, _c_F_PostmasterStateMachine[10]))
	if v313 != 0 {
		goto L81
	} else {
		goto L82
	}
L79:
	;
	F_errfinish(m, int32(_a_F_PostmasterStateMachine_3), int32(3320), int32(_a_F_PostmasterStateMachine_4))
	mBase = m.M
	v308 = m.ExcPending
	if v308 != 0 {
		goto L7
	} else {
		goto L80
	}
L80:
	;
	goto L78
L81:
	;
	F_FreeWaitEventSet(m, v313)
	mBase = m.M
	v315 = m.ExcPending
	if v315 != 0 {
		goto L7
	} else {
		goto L84
	}
L82:
	;
	goto L83
L83:
	;
	v316 = int32(_a_F_PostmasterStateMachine_14)
	v317 = int32(0)
	*(*int32)(unsafe.Add(mBase, _c_F_PostmasterStateMachine[10])) = v317
	v322 = F_CreateWaitEventSet(m, v317, int32(1))
	mBase = m.M
	v323 = m.ExcPending
	if v323 != 0 {
		goto L7
	} else {
		goto L85
	}
L84:
	;
	goto L83
L85:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_PostmasterStateMachine[10])) = v322
	v328 = *(*int32)(unsafe.Add(mBase, _c_F_PostmasterStateMachine[11]))
	F_AddWaitEventToSet(m, v322, int32(1), int32(-1), v328)
	mBase = m.M
	v330 = m.ExcPending
	if v330 != 0 {
		goto L7
	} else {
		goto L86
	}
L86:
	;
	v332 = *(*int32)(unsafe.Add(mBase, _c_F_PostmasterStateMachine[7]))
	if base.B2i32(v332 == int32(0))|base.B2i32(v332 == int32(_a_F_PostmasterStateMachine_7)) != 0 {
		goto L2
	} else {
		goto L87
	}
L87:
	;
	v338 = v332
	goto L88
L88:
	;
	v348 = *(*int32)(unsafe.Add(mBase, uint32(v338-int32(12))))
	if v348 == int32(2) {
		goto L90
	} else {
		goto L91
	}
L89:
	;
	goto L2
L90:
	;
	F_signal_child(m, v338-int32(20), int32(3))
	mBase = m.M
	v355 = m.ExcPending
	if v355 != 0 {
		goto L7
	} else {
		goto L93
	}
L91:
	;
	goto L92
L92:
	;
	v356 = *(*int32)(unsafe.Add(mBase, uint32(v338)+4))
	if v356 != int32(_a_F_PostmasterStateMachine_7) {
		v338 = v356
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
	v403 = m.ExcPending
	if v403 != 0 {
		goto L7
	} else {
		goto L108
	}
L96:
	;
	v365 = F_StartChildProcess(m, int32(11))
	mBase = m.M
	v366 = m.ExcPending
	if v366 != 0 {
		goto L7
	} else {
		goto L99
	}
L97:
	;
	v370 = v360
	goto L98
L98:
	;
	F_signal_child(m, v370, int32(2))
	mBase = m.M
	v373 = m.ExcPending
	if v373 != 0 {
		goto L7
	} else {
		goto L101
	}
L99:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_PostmasterStateMachine[9])) = v365
	if v365 == int32(0) {
		goto L95
	} else {
		goto L100
	}
L100:
	;
	v370 = v365
	goto L98
L101:
	;
	v376 = F_errstart(m, int32(14), int32(0))
	mBase = m.M
	v377 = m.ExcPending
	if v377 != 0 {
		goto L7
	} else {
		goto L102
	}
L102:
	;
	if v376 != 0 {
		goto L103
	} else {
		goto L104
	}
L103:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v11)+84)) = int32(_a_F_PostmasterStateMachine_15)
	v381 = *(*int32)(unsafe.Add(mBase, _c_F_PostmasterStateMachine[0]))
	v386 = *(*int32)(unsafe.Add(mBase, uint32(v381<<(uint(int32(2))%32))+uint32(_c_F_PostmasterStateMachine[3])))
	*(*int32)(unsafe.Add(mBase, uint32(v11)+80)) = v386
	F_errmsg_internal(m, int32(_a_F_PostmasterStateMachine_2), v11+int32(80))
	mBase = m.M
	v392 = m.ExcPending
	if v392 != 0 {
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
	F_errfinish(m, int32(_a_F_PostmasterStateMachine_3), int32(3320), int32(_a_F_PostmasterStateMachine_4))
	mBase = m.M
	v397 = m.ExcPending
	if v397 != 0 {
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
	v417 = F_CountChildren(m, int32(_a_F_PostmasterStateMachine_16))
	mBase = m.M
	v418 = m.ExcPending
	if v418 != 0 {
		goto L7
	} else {
		goto L110
	}
L110:
	;
	if v417 != 0 {
		goto L1
	} else {
		goto L111
	}
L111:
	;
	v421 = F_errstart(m, int32(14), int32(0))
	mBase = m.M
	v422 = m.ExcPending
	if v422 != 0 {
		goto L7
	} else {
		goto L112
	}
L112:
	;
	if v421 != 0 {
		goto L113
	} else {
		goto L114
	}
L113:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v11)+52)) = int32(_a_F_PostmasterStateMachine_17)
	v426 = *(*int32)(unsafe.Add(mBase, _c_F_PostmasterStateMachine[0]))
	v431 = *(*int32)(unsafe.Add(mBase, uint32(v426<<(uint(int32(2))%32))+uint32(_c_F_PostmasterStateMachine[3])))
	*(*int32)(unsafe.Add(mBase, uint32(v11)+48)) = v431
	F_errmsg_internal(m, int32(_a_F_PostmasterStateMachine_2), v11+int32(48))
	mBase = m.M
	v437 = m.ExcPending
	if v437 != 0 {
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
	v447 = *(*int32)(unsafe.Add(mBase, _c_F_PostmasterStateMachine[7]))
	if base.B2i32(v447 == int32(0))|base.B2i32(v447 == int32(_a_F_PostmasterStateMachine_7)) != 0 {
		goto L1
	} else {
		goto L118
	}
L116:
	;
	F_errfinish(m, int32(_a_F_PostmasterStateMachine_3), int32(3320), int32(_a_F_PostmasterStateMachine_4))
	mBase = m.M
	v442 = m.ExcPending
	if v442 != 0 {
		goto L7
	} else {
		goto L117
	}
L117:
	;
	goto L115
L118:
	;
	v453 = v447
	goto L119
L119:
	;
	v461 = int32(12)
	v463 = *(*int32)(unsafe.Add(mBase, uint32(v453-v461)))
	if v463 == v461 {
		goto L121
	} else {
		goto L122
	}
L120:
	;
	goto L1
L121:
	;
	F_signal_child(m, v453-int32(20), int32(12))
	mBase = m.M
	v470 = m.ExcPending
	if v470 != 0 {
		goto L7
	} else {
		goto L124
	}
L122:
	;
	goto L123
L123:
	;
	v471 = *(*int32)(unsafe.Add(mBase, uint32(v453)+4))
	if v471 != int32(_a_F_PostmasterStateMachine_7) {
		v453 = v471
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
	if v524 == int32(11) {
		goto L141
	} else {
		goto L142
	}
L128:
	;
	v487 = *(*int32)(unsafe.Add(mBase, _c_F_PostmasterStateMachine[12]))
	if v487 != 0 {
		v524 = v483
		goto L127
	} else {
		goto L129
	}
L129:
	;
	v490 = F_errstart(m, int32(14), int32(0))
	mBase = m.M
	v491 = m.ExcPending
	if v491 != 0 {
		goto L7
	} else {
		goto L130
	}
L130:
	;
	if v490 != 0 {
		goto L131
	} else {
		goto L132
	}
L131:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v11)+36)) = int32(_a_F_PostmasterStateMachine_18)
	v495 = *(*int32)(unsafe.Add(mBase, _c_F_PostmasterStateMachine[0]))
	v500 = *(*int32)(unsafe.Add(mBase, uint32(v495<<(uint(int32(2))%32))+uint32(_c_F_PostmasterStateMachine[3])))
	*(*int32)(unsafe.Add(mBase, uint32(v11)+32)) = v500
	F_errmsg_internal(m, int32(_a_F_PostmasterStateMachine_2), v11+int32(32))
	mBase = m.M
	v506 = m.ExcPending
	if v506 != 0 {
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
	v516 = *(*int32)(unsafe.Add(mBase, _c_F_PostmasterStateMachine[9]))
	if v516 == int32(0) {
		goto L126
	} else {
		goto L136
	}
L134:
	;
	F_errfinish(m, int32(_a_F_PostmasterStateMachine_3), int32(3320), int32(_a_F_PostmasterStateMachine_4))
	mBase = m.M
	v511 = m.ExcPending
	if v511 != 0 {
		goto L7
	} else {
		goto L135
	}
L135:
	;
	goto L133
L136:
	;
	F_signal_child(m, v516, int32(12))
	mBase = m.M
	v521 = m.ExcPending
	if v521 != 0 {
		goto L7
	} else {
		goto L137
	}
L137:
	;
	v523 = *(*int32)(unsafe.Add(mBase, _c_F_PostmasterStateMachine[0]))
	v524 = v523
	goto L127
L138:
	;
	v600 = *(*int32)(unsafe.Add(mBase, _c_F_PostmasterStateMachine[13]))
	if v600 == int32(3) {
		goto L166
	} else {
		goto L167
	}
L139:
	;
	v578 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_PostmasterStateMachine[4])))
	if v578 != 0 {
		goto L155
	} else {
		goto L156
	}
L140:
	;
	v547 = F_errstart(m, int32(14), int32(0))
	mBase = m.M
	v548 = m.ExcPending
	if v548 != 0 {
		goto L7
	} else {
		goto L148
	}
L141:
	;
	v529 = F_CountChildren(m, int32(_a_F_PostmasterStateMachine_19))
	mBase = m.M
	v530 = m.ExcPending
	if v530 != 0 {
		goto L7
	} else {
		goto L144
	}
L142:
	;
	v535 = v524
	goto L143
L143:
	;
	v539 = *(*int32)(unsafe.Add(mBase, _c_F_PostmasterStateMachine[1]))
	if base.B2i32(v535 == int32(12))&base.B2i32(int32(0) < v539) != 0 {
		goto L139
	} else {
		goto L146
	}
L144:
	;
	if v529 == int32(0) {
		goto L140
	} else {
		goto L145
	}
L145:
	;
	v534 = *(*int32)(unsafe.Add(mBase, _c_F_PostmasterStateMachine[0]))
	v535 = v534
	goto L143
L146:
	;
	if v535 != int32(12) {
		goto L126
	} else {
		goto L147
	}
L147:
	;
	goto L138
L148:
	;
	if v547 != 0 {
		goto L149
	} else {
		goto L150
	}
L149:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v11)+20)) = int32(_a_F_PostmasterStateMachine_20)
	v552 = *(*int32)(unsafe.Add(mBase, _c_F_PostmasterStateMachine[0]))
	v557 = *(*int32)(unsafe.Add(mBase, uint32(v552<<(uint(int32(2))%32))+uint32(_c_F_PostmasterStateMachine[3])))
	*(*int32)(unsafe.Add(mBase, uint32(v11)+16)) = v557
	F_errmsg_internal(m, int32(_a_F_PostmasterStateMachine_2), v11+int32(16))
	mBase = m.M
	v563 = m.ExcPending
	if v563 != 0 {
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
	v573 = *(*int32)(unsafe.Add(mBase, _c_F_PostmasterStateMachine[1]))
	if v573 <= int32(0) {
		goto L138
	} else {
		goto L154
	}
L152:
	;
	F_errfinish(m, int32(_a_F_PostmasterStateMachine_3), int32(3320), int32(_a_F_PostmasterStateMachine_4))
	mBase = m.M
	v568 = m.ExcPending
	if v568 != 0 {
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
	v581 = F_errstart(m, int32(15), int32(0))
	mBase = m.M
	v582 = m.ExcPending
	if v582 != 0 {
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
	v597 = m.ExcPending
	if v597 != 0 {
		goto L7
	} else {
		goto L165
	}
L158:
	;
	if v581 != 0 {
		goto L159
	} else {
		goto L160
	}
L159:
	;
	F_errmsg(m, int32(_a_F_PostmasterStateMachine_21), int32(0))
	mBase = m.M
	v586 = m.ExcPending
	if v586 != 0 {
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
	v594 = m.ExcPending
	if v594 != 0 {
		goto L7
	} else {
		goto L164
	}
L162:
	;
	F_errfinish(m, int32(_a_F_PostmasterStateMachine_3), int32(3202), int32(_a_F_PostmasterStateMachine_22))
	mBase = m.M
	v591 = m.ExcPending
	if v591 != 0 {
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
	v605 = F_errstart(m, int32(15), int32(0))
	mBase = m.M
	v606 = m.ExcPending
	if v606 != 0 {
		goto L7
	} else {
		goto L169
	}
L167:
	;
	goto L168
L168:
	;
	v620 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_PostmasterStateMachine[14])))
	if v620 == int32(0) {
		goto L176
	} else {
		goto L177
	}
L169:
	;
	if v605 != 0 {
		goto L170
	} else {
		goto L171
	}
L170:
	;
	F_errmsg(m, int32(_a_F_PostmasterStateMachine_23), int32(0))
	mBase = m.M
	v610 = m.ExcPending
	if v610 != 0 {
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
	v618 = m.ExcPending
	if v618 != 0 {
		goto L7
	} else {
		goto L175
	}
L173:
	;
	F_errfinish(m, int32(_a_F_PostmasterStateMachine_3), int32(3228), int32(_a_F_PostmasterStateMachine_22))
	mBase = m.M
	v615 = m.ExcPending
	if v615 != 0 {
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
	v625 = F_errstart(m, int32(15), int32(0))
	mBase = m.M
	v626 = m.ExcPending
	if v626 != 0 {
		goto L7
	} else {
		goto L179
	}
L177:
	;
	goto L178
L178:
	;
	v640 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_PostmasterStateMachine[4])))
	if v640 == int32(0) {
		goto L126
	} else {
		goto L186
	}
L179:
	;
	if v625 != 0 {
		goto L180
	} else {
		goto L181
	}
L180:
	;
	F_errmsg(m, int32(_a_F_PostmasterStateMachine_24), int32(0))
	mBase = m.M
	v630 = m.ExcPending
	if v630 != 0 {
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
	v638 = m.ExcPending
	if v638 != 0 {
		goto L7
	} else {
		goto L185
	}
L183:
	;
	F_errfinish(m, int32(_a_F_PostmasterStateMachine_3), int32(3234), int32(_a_F_PostmasterStateMachine_22))
	mBase = m.M
	v635 = m.ExcPending
	if v635 != 0 {
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
	v645 = F_errstart(m, int32(15), int32(0))
	mBase = m.M
	v646 = m.ExcPending
	if v646 != 0 {
		goto L7
	} else {
		goto L187
	}
L187:
	;
	if v645 != 0 {
		goto L188
	} else {
		goto L189
	}
L188:
	;
	F_errmsg(m, int32(_a_F_PostmasterStateMachine_25), int32(0))
	mBase = m.M
	v650 = m.ExcPending
	if v650 != 0 {
		goto L7
	} else {
		goto L191
	}
L189:
	;
	goto L190
L190:
	;
	v657 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_PostmasterStateMachine[15])))
	if v657 == int32(1) {
		goto L193
	} else {
		goto L194
	}
L191:
	;
	F_errfinish(m, int32(_a_F_PostmasterStateMachine_3), int32(3246), int32(_a_F_PostmasterStateMachine_22))
	mBase = m.M
	v655 = m.ExcPending
	if v655 != 0 {
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
	v661 = m.ExcPending
	if v661 != 0 {
		goto L7
	} else {
		goto L196
	}
L194:
	;
	goto L195
L195:
	;
	v662 = m.G0
	v664 = v662 - int32(16)
	m.G0 = v664
	v667 = *(*int32)(unsafe.Add(mBase, _c_F_PostmasterStateMachine[5]))
	v668 = int32(0)
	if base.B2i32(v667 == v668)|base.B2i32(v667 == int32(_a_F_PostmasterStateMachine_6)) == v668 {
		goto L197
	} else {
		goto L198
	}
L196:
	;
	goto L195
L197:
	;
	v675 = v667
	goto L200
L198:
	;
	goto L199
L199:
	;
	m.G0 = v664 + int32(16)
	F_shmem_exit(m, int32(1))
	mBase = m.M
	v765 = m.ExcPending
	if v765 != 0 {
		goto L7
	} else {
		goto L217
	}
L200:
	;
	v683 = *(*int32)(unsafe.Add(mBase, uint32(v675)+4))
	v686 = *(*int32)(unsafe.Add(mBase, uint32(v675-int32(1296))))
	if v686 == int32(-1) {
		goto L203
	} else {
		goto L204
	}
L201:
	;
	goto L199
L202:
	;
	if v683 != int32(_a_F_PostmasterStateMachine_6) {
		v675 = v683
		goto L200
	} else {
		goto L216
	}
L203:
	;
	v690 = *(*int32)(unsafe.Add(mBase, _c_F_PostmasterStateMachine[6]))
	v693 = *(*int32)(unsafe.Add(mBase, uint32(v675-int32(8))))
	v697 = int32(16)
	v701 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v675-int32(1304)))))
	if v701&v697 != 0 {
		goto L206
	} else {
		goto L207
	}
L204:
	;
	goto L205
L205:
	;
	v738 = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v675-int32(24)))) = v738
	*(*int64)(unsafe.Add(mBase, uint32(v675-int32(16)))) = int64(0)
	*(*int32)(unsafe.Add(mBase, uint32(v675-int32(32)))) = v738
	goto L202
L206:
	;
	v704 = *(*int32)(unsafe.Add(mBase, uint32(v690)+8))
	*(*int32)(unsafe.Add(mBase, uint32(v690)+8)) = v704 + int32(1)
	goto L208
L207:
	;
	goto L208
L208:
	;
	v709 = v675 - int32(1496)
	v710 = int32(0)
	v713 = base.AtomicRmwOr32(m, v710, int32(_a_F_PostmasterStateMachine_8), v710)
	*(*uint8)(unsafe.Add(mBase, uint32(v690+v693*int32(1488)+v697))) = uint8(v710)
	v718 = F_errstart(m, int32(14), v710)
	mBase = m.M
	v719 = m.ExcPending
	if v719 != 0 {
		goto L7
	} else {
		goto L209
	}
L209:
	;
	if v718 != 0 {
		goto L210
	} else {
		goto L211
	}
L210:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v664))) = v709
	F_errmsg_internal(m, int32(_a_F_PostmasterStateMachine_9), v664)
	mBase = m.M
	v723 = m.ExcPending
	if v723 != 0 {
		goto L7
	} else {
		goto L213
	}
L211:
	;
	goto L212
L212:
	;
	v729 = *(*int32)(unsafe.Add(mBase, uint32(v675)))
	v730 = *(*int32)(unsafe.Add(mBase, uint32(v675)+4))
	*(*int32)(unsafe.Add(mBase, uint32(v729)+4)) = v730
	v732 = *(*int32)(unsafe.Add(mBase, uint32(v675)))
	*(*int32)(unsafe.Add(mBase, uint32(v730))) = v732
	F_pfree(m, v709)
	mBase = m.M
	v735 = m.ExcPending
	if v735 != 0 {
		goto L7
	} else {
		goto L215
	}
L213:
	;
	F_errfinish(m, int32(_a_F_PostmasterStateMachine_10), int32(466), int32(_a_F_PostmasterStateMachine_11))
	mBase = m.M
	v728 = m.ExcPending
	if v728 != 0 {
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
	v767 = m.ExcPending
	if v767 != 0 {
		goto L7
	} else {
		goto L218
	}
L218:
	;
	v769 = int32(0)
	*(*int32)(unsafe.Add(mBase, _c_F_PostmasterStateMachine[16])) = v769
	*(*int32)(unsafe.Add(mBase, _c_F_PostmasterStateMachine[17])) = v769
	F_ShmemCallRequestCallbacks(m)
	mBase = m.M
	v775 = m.ExcPending
	if v775 != 0 {
		goto L7
	} else {
		goto L219
	}
L219:
	;
	F_CreateSharedMemoryAndSemaphores(m)
	mBase = m.M
	v777 = m.ExcPending
	if v777 != 0 {
		goto L7
	} else {
		goto L220
	}
L220:
	;
	v780 = F_errstart(m, int32(14), int32(0))
	mBase = m.M
	v781 = m.ExcPending
	if v781 != 0 {
		goto L7
	} else {
		goto L221
	}
L221:
	;
	if v780 != 0 {
		goto L222
	} else {
		goto L223
	}
L222:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v11)+4)) = int32(_a_F_PostmasterStateMachine_26)
	v785 = *(*int32)(unsafe.Add(mBase, _c_F_PostmasterStateMachine[0]))
	v790 = *(*int32)(unsafe.Add(mBase, uint32(v785<<(uint(int32(2))%32))+uint32(_c_F_PostmasterStateMachine[3])))
	*(*int32)(unsafe.Add(mBase, uint32(v11))) = v790
	F_errmsg_internal(m, int32(_a_F_PostmasterStateMachine_2), v11)
	mBase = m.M
	v794 = m.ExcPending
	if v794 != 0 {
		goto L7
	} else {
		goto L225
	}
L223:
	;
	goto L224
L224:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_PostmasterStateMachine[0])) = int32(1)
	F_maybe_start_io_workers(m)
	mBase = m.M
	v804 = m.ExcPending
	if v804 != 0 {
		goto L7
	} else {
		goto L227
	}
L225:
	;
	F_errfinish(m, int32(_a_F_PostmasterStateMachine_3), int32(3320), int32(_a_F_PostmasterStateMachine_4))
	mBase = m.M
	v799 = m.ExcPending
	if v799 != 0 {
		goto L7
	} else {
		goto L226
	}
L226:
	;
	goto L224
L227:
	;
	v806 = F_StartChildProcess(m, int32(13))
	mBase = m.M
	v807 = m.ExcPending
	if v807 != 0 {
		goto L7
	} else {
		goto L228
	}
L228:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_PostmasterStateMachine[13])) = int32(1)
	*(*int32)(unsafe.Add(mBase, _c_F_PostmasterStateMachine[18])) = v806
	*(*int64)(unsafe.Add(mBase, _c_F_PostmasterStateMachine[19])) = int64(0)
	F_ConfigurePostmasterWaitSet(m)
	mBase = m.M
	v817 = m.ExcPending
	if v817 != 0 {
		goto L7
	} else {
		goto L229
	}
L229:
	;
	goto L126
}
