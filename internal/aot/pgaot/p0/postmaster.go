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
	var v154 int32
	_ = v154
	var v156 int32
	_ = v156
	var v157 int32
	_ = v157
	var v174 int32
	_ = v174
	var v175 int32
	_ = v175
	var v184 int32
	_ = v184
	var v193 int32
	_ = v193
	var v194 int32
	_ = v194
	var v199 int32
	_ = v199
	var v201 int32
	_ = v201
	var v205 int32
	_ = v205
	var v215 int32
	_ = v215
	var v223 int32
	_ = v223
	var v224 int32
	_ = v224
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
	var v270 int32
	_ = v270
	var v271 int32
	_ = v271
	var v273 int32
	_ = v273
	var v277 int32
	_ = v277
	var v284 int32
	_ = v284
	var v285 int32
	_ = v285
	var v289 int32
	_ = v289
	var v294 int32
	_ = v294
	var v300 int32
	_ = v300
	var v305 int32
	_ = v305
	var v310 int32
	_ = v310
	var v312 int32
	_ = v312
	var v313 int32
	_ = v313
	var v314 int32
	_ = v314
	var v319 int32
	_ = v319
	var v320 int32
	_ = v320
	var v325 int32
	_ = v325
	var v327 int32
	_ = v327
	var v329 int32
	_ = v329
	var v335 int32
	_ = v335
	var v345 int32
	_ = v345
	var v352 int32
	_ = v352
	var v353 int32
	_ = v353
	var v357 int32
	_ = v357
	var v362 int32
	_ = v362
	var v363 int32
	_ = v363
	var v367 int32
	_ = v367
	var v370 int32
	_ = v370
	var v373 int32
	_ = v373
	var v374 int32
	_ = v374
	var v378 int32
	_ = v378
	var v383 int32
	_ = v383
	var v389 int32
	_ = v389
	var v394 int32
	_ = v394
	var v400 int32
	_ = v400
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
	var v450 int32
	_ = v450
	var v458 int32
	_ = v458
	var v460 int32
	_ = v460
	var v467 int32
	_ = v467
	var v468 int32
	_ = v468
	var v480 int32
	_ = v480
	var v484 int32
	_ = v484
	var v487 int32
	_ = v487
	var v488 int32
	_ = v488
	var v492 int32
	_ = v492
	var v497 int32
	_ = v497
	var v503 int32
	_ = v503
	var v508 int32
	_ = v508
	var v513 int32
	_ = v513
	var v518 int32
	_ = v518
	var v520 int32
	_ = v520
	var v521 int32
	_ = v521
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
	var v665 int32
	_ = v665
	var v672 int32
	_ = v672
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
	var v753 int32
	_ = v753
	var v755 int32
	_ = v755
	var v757 int32
	_ = v757
	var v760 int32
	_ = v760
	var v761 int32
	_ = v761
	var v765 int32
	_ = v765
	var v770 int32
	_ = v770
	var v774 int32
	_ = v774
	var v779 int32
	_ = v779
	var v784 int32
	_ = v784
	var v786 int32
	_ = v786
	var v787 int32
	_ = v787
	var v797 int32
	_ = v797
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
	v480 = *(*int32)(unsafe.Add(mBase, _c_F_PostmasterStateMachine[0]))
	if v480 != int32(9) {
		v521 = v480
		goto L128
	} else {
		goto L129
	}
L2:
	;
	v410 = *(*int32)(unsafe.Add(mBase, _c_F_PostmasterStateMachine[0]))
	if v410 != int32(8) {
		goto L1
	} else {
		goto L110
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
	v270 = F_CountChildren(m, v76)
	mBase = m.M
	v271 = m.ExcPending
	if v271 != 0 {
		goto L7
	} else {
		goto L69
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
	v174 = *(*int32)(unsafe.Add(mBase, _c_F_PostmasterStateMachine[7]))
	v175 = int32(0)
	if base.B2i32(v174 == v175)|base.B2i32(v174 == int32(_a_F_PostmasterStateMachine_7)) == v175 {
		goto L48
	} else {
		goto L49
	}
L31:
	;
	v102 = *(*int32)(unsafe.Add(mBase, uint32(v96)+4))
	v105 = *(*int32)(unsafe.Add(mBase, uint32(v96-int32(8))))
	v108 = v97 + v105*int32(1480)
	v109 = *(*int32)(unsafe.Add(mBase, uint32(v108)+20))
	if v109 != int32(-1) {
		v157 = v97
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
		v97 = v157
		goto L31
	} else {
		goto L47
	}
L34:
	;
	v114 = *(*int32)(unsafe.Add(mBase, uint32(v96-int32(24))))
	if v114 == int32(0) {
		v157 = v97
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
	v153 = F_kill(m, v114, int32(10))
	mBase = m.M
	v154 = m.ExcPending
	if v154 != 0 {
		goto L7
	} else {
		goto L46
	}
L46:
	;
	v156 = *(*int32)(unsafe.Add(mBase, _c_F_PostmasterStateMachine[6]))
	v157 = v156
	goto L33
L47:
	;
	goto L32
L48:
	;
	v184 = v174
	goto L51
L49:
	;
	goto L50
L50:
	;
	v237 = F_errstart(m, int32(14), int32(0))
	mBase = m.M
	v238 = m.ExcPending
	if v238 != 0 {
		goto L7
	} else {
		goto L63
	}
L51:
	;
	if v76&int32(64) != 0 {
		goto L53
	} else {
		goto L54
	}
L52:
	;
	goto L50
L53:
	;
	v215 = *(*int32)(unsafe.Add(mBase, uint32(v184-int32(12))))
	if int32(base.Ui32(v76)>>(uint(v215)%32))&int32(1) != 0 {
		goto L58
	} else {
		goto L59
	}
L54:
	;
	v193 = v184 - int32(12)
	v194 = *(*int32)(unsafe.Add(mBase, uint32(v193)))
	if v194 != int32(1) {
		goto L53
	} else {
		goto L55
	}
L55:
	;
	v199 = *(*int32)(unsafe.Add(mBase, uint32(v184-int32(16))))
	v201 = *(*int32)(unsafe.Add(mBase, _c_F_PostmasterStateMachine[8]))
	v205 = *(*int32)(unsafe.Add(mBase, uint32(v201+v199<<(uint(int32(2))%32))+44))
	goto L56
L56:
	;
	if base.B2i32(v205 == int32(3)) == int32(0) {
		goto L53
	} else {
		goto L57
	}
L57:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v193))) = int32(6)
	goto L53
L58:
	;
	F_signal_child(m, v184-int32(20), int32(15))
	mBase = m.M
	v223 = m.ExcPending
	if v223 != 0 {
		goto L7
	} else {
		goto L61
	}
L59:
	;
	goto L60
L60:
	;
	v224 = *(*int32)(unsafe.Add(mBase, uint32(v184)+4))
	if v224 != int32(_a_F_PostmasterStateMachine_7) {
		v184 = v224
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
	*(*int32)(unsafe.Add(mBase, uint32(v11)+100)) = int32(_a_F_PostmasterStateMachine_11)
	v242 = *(*int32)(unsafe.Add(mBase, _c_F_PostmasterStateMachine[0]))
	v247 = *(*int32)(unsafe.Add(mBase, uint32(v242<<(uint(int32(2))%32))+uint32(_c_F_PostmasterStateMachine[3])))
	*(*int32)(unsafe.Add(mBase, uint32(v11)+96)) = v247
	F_errmsg_internal(m, int32(_a_F_PostmasterStateMachine_2), v11+int32(96))
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
	*(*int32)(unsafe.Add(mBase, _c_F_PostmasterStateMachine[0])) = int32(6)
	goto L27
L67:
	;
	F_errfinish(m, int32(_a_F_PostmasterStateMachine_3), int32(3272), int32(_a_F_PostmasterStateMachine_4))
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
	if v270 != 0 {
		goto L2
	} else {
		goto L70
	}
L70:
	;
	v273 = *(*int32)(unsafe.Add(mBase, _c_F_PostmasterStateMachine[1]))
	if v273 <= int32(2) {
		goto L72
	} else {
		goto L73
	}
L71:
	;
	v357 = *(*int32)(unsafe.Add(mBase, _c_F_PostmasterStateMachine[9]))
	if v357 == int32(0) {
		goto L97
	} else {
		goto L98
	}
L72:
	;
	v277 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_PostmasterStateMachine[4])))
	if v277&int32(1) == int32(0) {
		goto L71
	} else {
		goto L75
	}
L73:
	;
	goto L74
L74:
	;
	v284 = F_errstart(m, int32(14), int32(0))
	mBase = m.M
	v285 = m.ExcPending
	if v285 != 0 {
		goto L7
	} else {
		goto L76
	}
L75:
	;
	goto L74
L76:
	;
	if v284 != 0 {
		goto L77
	} else {
		goto L78
	}
L77:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v11)+68)) = int32(_a_F_PostmasterStateMachine_12)
	v289 = *(*int32)(unsafe.Add(mBase, _c_F_PostmasterStateMachine[0]))
	v294 = *(*int32)(unsafe.Add(mBase, uint32(v289<<(uint(int32(2))%32))+uint32(_c_F_PostmasterStateMachine[3])))
	*(*int32)(unsafe.Add(mBase, uint32(v11)+64)) = v294
	F_errmsg_internal(m, int32(_a_F_PostmasterStateMachine_2), v11-int32(-64))
	mBase = m.M
	v300 = m.ExcPending
	if v300 != 0 {
		goto L7
	} else {
		goto L80
	}
L78:
	;
	goto L79
L79:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_PostmasterStateMachine[0])) = int32(11)
	v310 = *(*int32)(unsafe.Add(mBase, _c_F_PostmasterStateMachine[10]))
	if v310 != 0 {
		goto L82
	} else {
		goto L83
	}
L80:
	;
	F_errfinish(m, int32(_a_F_PostmasterStateMachine_3), int32(3272), int32(_a_F_PostmasterStateMachine_4))
	mBase = m.M
	v305 = m.ExcPending
	if v305 != 0 {
		goto L7
	} else {
		goto L81
	}
L81:
	;
	goto L79
L82:
	;
	F_FreeWaitEventSet(m, v310)
	mBase = m.M
	v312 = m.ExcPending
	if v312 != 0 {
		goto L7
	} else {
		goto L85
	}
L83:
	;
	goto L84
L84:
	;
	v313 = int32(_a_F_PostmasterStateMachine_13)
	v314 = int32(0)
	*(*int32)(unsafe.Add(mBase, _c_F_PostmasterStateMachine[10])) = v314
	v319 = F_CreateWaitEventSet(m, v314, int32(1))
	mBase = m.M
	v320 = m.ExcPending
	if v320 != 0 {
		goto L7
	} else {
		goto L86
	}
L85:
	;
	goto L84
L86:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_PostmasterStateMachine[10])) = v319
	v325 = *(*int32)(unsafe.Add(mBase, _c_F_PostmasterStateMachine[11]))
	F_AddWaitEventToSet(m, v319, int32(1), int32(-1), v325)
	mBase = m.M
	v327 = m.ExcPending
	if v327 != 0 {
		goto L7
	} else {
		goto L87
	}
L87:
	;
	v329 = *(*int32)(unsafe.Add(mBase, _c_F_PostmasterStateMachine[7]))
	if base.B2i32(v329 == int32(0))|base.B2i32(v329 == int32(_a_F_PostmasterStateMachine_7)) != 0 {
		goto L2
	} else {
		goto L88
	}
L88:
	;
	v335 = v329
	goto L89
L89:
	;
	v345 = *(*int32)(unsafe.Add(mBase, uint32(v335-int32(12))))
	if v345 == int32(2) {
		goto L91
	} else {
		goto L92
	}
L90:
	;
	goto L2
L91:
	;
	F_signal_child(m, v335-int32(20), int32(3))
	mBase = m.M
	v352 = m.ExcPending
	if v352 != 0 {
		goto L7
	} else {
		goto L94
	}
L92:
	;
	goto L93
L93:
	;
	v353 = *(*int32)(unsafe.Add(mBase, uint32(v335)+4))
	if v353 != int32(_a_F_PostmasterStateMachine_7) {
		v335 = v353
		goto L89
	} else {
		goto L95
	}
L94:
	;
	goto L93
L95:
	;
	goto L90
L96:
	;
	F_HandleFatalError(m, int32(0))
	mBase = m.M
	v400 = m.ExcPending
	if v400 != 0 {
		goto L7
	} else {
		goto L109
	}
L97:
	;
	v362 = F_StartChildProcess(m, int32(11))
	mBase = m.M
	v363 = m.ExcPending
	if v363 != 0 {
		goto L7
	} else {
		goto L100
	}
L98:
	;
	v367 = v357
	goto L99
L99:
	;
	F_signal_child(m, v367, int32(2))
	mBase = m.M
	v370 = m.ExcPending
	if v370 != 0 {
		goto L7
	} else {
		goto L102
	}
L100:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_PostmasterStateMachine[9])) = v362
	if v362 == int32(0) {
		goto L96
	} else {
		goto L101
	}
L101:
	;
	v367 = v362
	goto L99
L102:
	;
	v373 = F_errstart(m, int32(14), int32(0))
	mBase = m.M
	v374 = m.ExcPending
	if v374 != 0 {
		goto L7
	} else {
		goto L103
	}
L103:
	;
	if v373 != 0 {
		goto L104
	} else {
		goto L105
	}
L104:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v11)+84)) = int32(_a_F_PostmasterStateMachine_14)
	v378 = *(*int32)(unsafe.Add(mBase, _c_F_PostmasterStateMachine[0]))
	v383 = *(*int32)(unsafe.Add(mBase, uint32(v378<<(uint(int32(2))%32))+uint32(_c_F_PostmasterStateMachine[3])))
	*(*int32)(unsafe.Add(mBase, uint32(v11)+80)) = v383
	F_errmsg_internal(m, int32(_a_F_PostmasterStateMachine_2), v11+int32(80))
	mBase = m.M
	v389 = m.ExcPending
	if v389 != 0 {
		goto L7
	} else {
		goto L107
	}
L105:
	;
	goto L106
L106:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_PostmasterStateMachine[0])) = int32(7)
	goto L1
L107:
	;
	F_errfinish(m, int32(_a_F_PostmasterStateMachine_3), int32(3272), int32(_a_F_PostmasterStateMachine_4))
	mBase = m.M
	v394 = m.ExcPending
	if v394 != 0 {
		goto L7
	} else {
		goto L108
	}
L108:
	;
	goto L106
L109:
	;
	goto L2
L110:
	;
	v414 = F_CountChildren(m, int32(_a_F_PostmasterStateMachine_15))
	mBase = m.M
	v415 = m.ExcPending
	if v415 != 0 {
		goto L7
	} else {
		goto L111
	}
L111:
	;
	if v414 != 0 {
		goto L1
	} else {
		goto L112
	}
L112:
	;
	v418 = F_errstart(m, int32(14), int32(0))
	mBase = m.M
	v419 = m.ExcPending
	if v419 != 0 {
		goto L7
	} else {
		goto L113
	}
L113:
	;
	if v418 != 0 {
		goto L114
	} else {
		goto L115
	}
L114:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v11)+52)) = int32(_a_F_PostmasterStateMachine_16)
	v423 = *(*int32)(unsafe.Add(mBase, _c_F_PostmasterStateMachine[0]))
	v428 = *(*int32)(unsafe.Add(mBase, uint32(v423<<(uint(int32(2))%32))+uint32(_c_F_PostmasterStateMachine[3])))
	*(*int32)(unsafe.Add(mBase, uint32(v11)+48)) = v428
	F_errmsg_internal(m, int32(_a_F_PostmasterStateMachine_2), v11+int32(48))
	mBase = m.M
	v434 = m.ExcPending
	if v434 != 0 {
		goto L7
	} else {
		goto L117
	}
L115:
	;
	goto L116
L116:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_PostmasterStateMachine[0])) = int32(9)
	v444 = *(*int32)(unsafe.Add(mBase, _c_F_PostmasterStateMachine[7]))
	if base.B2i32(v444 == int32(0))|base.B2i32(v444 == int32(_a_F_PostmasterStateMachine_7)) != 0 {
		goto L1
	} else {
		goto L119
	}
L117:
	;
	F_errfinish(m, int32(_a_F_PostmasterStateMachine_3), int32(3272), int32(_a_F_PostmasterStateMachine_4))
	mBase = m.M
	v439 = m.ExcPending
	if v439 != 0 {
		goto L7
	} else {
		goto L118
	}
L118:
	;
	goto L116
L119:
	;
	v450 = v444
	goto L120
L120:
	;
	v458 = int32(12)
	v460 = *(*int32)(unsafe.Add(mBase, uint32(v450-v458)))
	if v460 == v458 {
		goto L122
	} else {
		goto L123
	}
L121:
	;
	goto L1
L122:
	;
	F_signal_child(m, v450-int32(20), int32(12))
	mBase = m.M
	v467 = m.ExcPending
	if v467 != 0 {
		goto L7
	} else {
		goto L125
	}
L123:
	;
	goto L124
L124:
	;
	v468 = *(*int32)(unsafe.Add(mBase, uint32(v450)+4))
	if v468 != int32(_a_F_PostmasterStateMachine_7) {
		v450 = v468
		goto L120
	} else {
		goto L126
	}
L125:
	;
	goto L124
L126:
	;
	goto L121
L127:
	;
	m.G0 = v11 + int32(128)
	return
L128:
	;
	if v521 == int32(11) {
		goto L142
	} else {
		goto L143
	}
L129:
	;
	v484 = *(*int32)(unsafe.Add(mBase, _c_F_PostmasterStateMachine[12]))
	if v484 != 0 {
		v521 = v480
		goto L128
	} else {
		goto L130
	}
L130:
	;
	v487 = F_errstart(m, int32(14), int32(0))
	mBase = m.M
	v488 = m.ExcPending
	if v488 != 0 {
		goto L7
	} else {
		goto L131
	}
L131:
	;
	if v487 != 0 {
		goto L132
	} else {
		goto L133
	}
L132:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v11)+36)) = int32(_a_F_PostmasterStateMachine_17)
	v492 = *(*int32)(unsafe.Add(mBase, _c_F_PostmasterStateMachine[0]))
	v497 = *(*int32)(unsafe.Add(mBase, uint32(v492<<(uint(int32(2))%32))+uint32(_c_F_PostmasterStateMachine[3])))
	*(*int32)(unsafe.Add(mBase, uint32(v11)+32)) = v497
	F_errmsg_internal(m, int32(_a_F_PostmasterStateMachine_2), v11+int32(32))
	mBase = m.M
	v503 = m.ExcPending
	if v503 != 0 {
		goto L7
	} else {
		goto L135
	}
L133:
	;
	goto L134
L134:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_PostmasterStateMachine[0])) = int32(10)
	v513 = *(*int32)(unsafe.Add(mBase, _c_F_PostmasterStateMachine[9]))
	if v513 == int32(0) {
		goto L127
	} else {
		goto L137
	}
L135:
	;
	F_errfinish(m, int32(_a_F_PostmasterStateMachine_3), int32(3272), int32(_a_F_PostmasterStateMachine_4))
	mBase = m.M
	v508 = m.ExcPending
	if v508 != 0 {
		goto L7
	} else {
		goto L136
	}
L136:
	;
	goto L134
L137:
	;
	F_signal_child(m, v513, int32(12))
	mBase = m.M
	v518 = m.ExcPending
	if v518 != 0 {
		goto L7
	} else {
		goto L138
	}
L138:
	;
	v520 = *(*int32)(unsafe.Add(mBase, _c_F_PostmasterStateMachine[0]))
	v521 = v520
	goto L128
L139:
	;
	v597 = *(*int32)(unsafe.Add(mBase, _c_F_PostmasterStateMachine[13]))
	if v597 == int32(3) {
		goto L167
	} else {
		goto L168
	}
L140:
	;
	v575 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_PostmasterStateMachine[4])))
	if v575 != 0 {
		goto L156
	} else {
		goto L157
	}
L141:
	;
	v544 = F_errstart(m, int32(14), int32(0))
	mBase = m.M
	v545 = m.ExcPending
	if v545 != 0 {
		goto L7
	} else {
		goto L149
	}
L142:
	;
	v526 = F_CountChildren(m, int32(_a_F_PostmasterStateMachine_18))
	mBase = m.M
	v527 = m.ExcPending
	if v527 != 0 {
		goto L7
	} else {
		goto L145
	}
L143:
	;
	v532 = v521
	goto L144
L144:
	;
	v536 = *(*int32)(unsafe.Add(mBase, _c_F_PostmasterStateMachine[1]))
	if base.B2i32(v532 == int32(12))&base.B2i32(int32(0) < v536) != 0 {
		goto L140
	} else {
		goto L147
	}
L145:
	;
	if v526 == int32(0) {
		goto L141
	} else {
		goto L146
	}
L146:
	;
	v531 = *(*int32)(unsafe.Add(mBase, _c_F_PostmasterStateMachine[0]))
	v532 = v531
	goto L144
L147:
	;
	if v532 != int32(12) {
		goto L127
	} else {
		goto L148
	}
L148:
	;
	goto L139
L149:
	;
	if v544 != 0 {
		goto L150
	} else {
		goto L151
	}
L150:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v11)+20)) = int32(_a_F_PostmasterStateMachine_19)
	v549 = *(*int32)(unsafe.Add(mBase, _c_F_PostmasterStateMachine[0]))
	v554 = *(*int32)(unsafe.Add(mBase, uint32(v549<<(uint(int32(2))%32))+uint32(_c_F_PostmasterStateMachine[3])))
	*(*int32)(unsafe.Add(mBase, uint32(v11)+16)) = v554
	F_errmsg_internal(m, int32(_a_F_PostmasterStateMachine_2), v11+int32(16))
	mBase = m.M
	v560 = m.ExcPending
	if v560 != 0 {
		goto L7
	} else {
		goto L153
	}
L151:
	;
	goto L152
L152:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_PostmasterStateMachine[0])) = int32(12)
	v570 = *(*int32)(unsafe.Add(mBase, _c_F_PostmasterStateMachine[1]))
	if v570 <= int32(0) {
		goto L139
	} else {
		goto L155
	}
L153:
	;
	F_errfinish(m, int32(_a_F_PostmasterStateMachine_3), int32(3272), int32(_a_F_PostmasterStateMachine_4))
	mBase = m.M
	v565 = m.ExcPending
	if v565 != 0 {
		goto L7
	} else {
		goto L154
	}
L154:
	;
	goto L152
L155:
	;
	goto L140
L156:
	;
	v578 = F_errstart(m, int32(15), int32(0))
	mBase = m.M
	v579 = m.ExcPending
	if v579 != 0 {
		goto L7
	} else {
		goto L159
	}
L157:
	;
	goto L158
L158:
	;
	F_ExitPostmaster(m, int32(0))
	mBase = m.M
	v594 = m.ExcPending
	if v594 != 0 {
		goto L7
	} else {
		goto L166
	}
L159:
	;
	if v578 != 0 {
		goto L160
	} else {
		goto L161
	}
L160:
	;
	F_errmsg(m, int32(_a_F_PostmasterStateMachine_20), int32(0))
	mBase = m.M
	v583 = m.ExcPending
	if v583 != 0 {
		goto L7
	} else {
		goto L163
	}
L161:
	;
	goto L162
L162:
	;
	F_ExitPostmaster(m, int32(1))
	mBase = m.M
	v591 = m.ExcPending
	if v591 != 0 {
		goto L7
	} else {
		goto L165
	}
L163:
	;
	F_errfinish(m, int32(_a_F_PostmasterStateMachine_3), int32(3161), int32(_a_F_PostmasterStateMachine_21))
	mBase = m.M
	v588 = m.ExcPending
	if v588 != 0 {
		goto L7
	} else {
		goto L164
	}
L164:
	;
	goto L162
L165:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L166:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L167:
	;
	v602 = F_errstart(m, int32(15), int32(0))
	mBase = m.M
	v603 = m.ExcPending
	if v603 != 0 {
		goto L7
	} else {
		goto L170
	}
L168:
	;
	goto L169
L169:
	;
	v617 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_PostmasterStateMachine[14])))
	if v617 == int32(0) {
		goto L177
	} else {
		goto L178
	}
L170:
	;
	if v602 != 0 {
		goto L171
	} else {
		goto L172
	}
L171:
	;
	F_errmsg(m, int32(_a_F_PostmasterStateMachine_22), int32(0))
	mBase = m.M
	v607 = m.ExcPending
	if v607 != 0 {
		goto L7
	} else {
		goto L174
	}
L172:
	;
	goto L173
L173:
	;
	F_ExitPostmaster(m, int32(1))
	mBase = m.M
	v615 = m.ExcPending
	if v615 != 0 {
		goto L7
	} else {
		goto L176
	}
L174:
	;
	F_errfinish(m, int32(_a_F_PostmasterStateMachine_3), int32(3187), int32(_a_F_PostmasterStateMachine_21))
	mBase = m.M
	v612 = m.ExcPending
	if v612 != 0 {
		goto L7
	} else {
		goto L175
	}
L175:
	;
	goto L173
L176:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L177:
	;
	v622 = F_errstart(m, int32(15), int32(0))
	mBase = m.M
	v623 = m.ExcPending
	if v623 != 0 {
		goto L7
	} else {
		goto L180
	}
L178:
	;
	goto L179
L179:
	;
	v637 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_PostmasterStateMachine[4])))
	if v637 == int32(0) {
		goto L127
	} else {
		goto L187
	}
L180:
	;
	if v622 != 0 {
		goto L181
	} else {
		goto L182
	}
L181:
	;
	F_errmsg(m, int32(_a_F_PostmasterStateMachine_23), int32(0))
	mBase = m.M
	v627 = m.ExcPending
	if v627 != 0 {
		goto L7
	} else {
		goto L184
	}
L182:
	;
	goto L183
L183:
	;
	F_ExitPostmaster(m, int32(1))
	mBase = m.M
	v635 = m.ExcPending
	if v635 != 0 {
		goto L7
	} else {
		goto L186
	}
L184:
	;
	F_errfinish(m, int32(_a_F_PostmasterStateMachine_3), int32(3193), int32(_a_F_PostmasterStateMachine_21))
	mBase = m.M
	v632 = m.ExcPending
	if v632 != 0 {
		goto L7
	} else {
		goto L185
	}
L185:
	;
	goto L183
L186:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L187:
	;
	v642 = F_errstart(m, int32(15), int32(0))
	mBase = m.M
	v643 = m.ExcPending
	if v643 != 0 {
		goto L7
	} else {
		goto L188
	}
L188:
	;
	if v642 != 0 {
		goto L189
	} else {
		goto L190
	}
L189:
	;
	F_errmsg(m, int32(_a_F_PostmasterStateMachine_24), int32(0))
	mBase = m.M
	v647 = m.ExcPending
	if v647 != 0 {
		goto L7
	} else {
		goto L192
	}
L190:
	;
	goto L191
L191:
	;
	v654 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_PostmasterStateMachine[15])))
	if v654 == int32(1) {
		goto L194
	} else {
		goto L195
	}
L192:
	;
	F_errfinish(m, int32(_a_F_PostmasterStateMachine_3), int32(3205), int32(_a_F_PostmasterStateMachine_21))
	mBase = m.M
	v652 = m.ExcPending
	if v652 != 0 {
		goto L7
	} else {
		goto L193
	}
L193:
	;
	goto L191
L194:
	;
	F_RemovePgTempFiles(m)
	mBase = m.M
	v658 = m.ExcPending
	if v658 != 0 {
		goto L7
	} else {
		goto L197
	}
L195:
	;
	goto L196
L196:
	;
	v659 = m.G0
	v661 = v659 - int32(16)
	m.G0 = v661
	v664 = *(*int32)(unsafe.Add(mBase, _c_F_PostmasterStateMachine[5]))
	v665 = int32(0)
	if base.B2i32(v664 == v665)|base.B2i32(v664 == int32(_a_F_PostmasterStateMachine_6)) == v665 {
		goto L198
	} else {
		goto L199
	}
L197:
	;
	goto L196
L198:
	;
	v672 = v664
	goto L201
L199:
	;
	goto L200
L200:
	;
	m.G0 = v661 + int32(16)
	F_shmem_exit(m, int32(1))
	mBase = m.M
	v753 = m.ExcPending
	if v753 != 0 {
		goto L7
	} else {
		goto L218
	}
L201:
	;
	v680 = *(*int32)(unsafe.Add(mBase, uint32(v672)+4))
	v683 = *(*int32)(unsafe.Add(mBase, uint32(v672-int32(1280))))
	if v683 == int32(-1) {
		goto L204
	} else {
		goto L205
	}
L202:
	;
	goto L200
L203:
	;
	if v680 != int32(_a_F_PostmasterStateMachine_6) {
		v672 = v680
		goto L201
	} else {
		goto L217
	}
L204:
	;
	v687 = *(*int32)(unsafe.Add(mBase, _c_F_PostmasterStateMachine[6]))
	v690 = *(*int32)(unsafe.Add(mBase, uint32(v672-int32(8))))
	v694 = int32(16)
	v698 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v672-int32(1288)))))
	if v698&v694 != 0 {
		goto L207
	} else {
		goto L208
	}
L205:
	;
	goto L206
L206:
	;
	v730 = v672 - int32(24)
	v731 = int64(0)
	*(*int64)(unsafe.Add(mBase, uint32(v730)+8)) = v731
	*(*int64)(unsafe.Add(mBase, uint32(v730))) = v731
	goto L203
L207:
	;
	v701 = *(*int32)(unsafe.Add(mBase, uint32(v687)+8))
	*(*int32)(unsafe.Add(mBase, uint32(v687)+8)) = v701 + int32(1)
	goto L209
L208:
	;
	goto L209
L209:
	;
	v706 = v672 - int32(1480)
	v707 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v687+v690*int32(1480)+v694))) = uint8(v707)
	v711 = F_errstart(m, int32(14), v707)
	mBase = m.M
	v712 = m.ExcPending
	if v712 != 0 {
		goto L7
	} else {
		goto L210
	}
L210:
	;
	if v711 != 0 {
		goto L211
	} else {
		goto L212
	}
L211:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v661))) = v706
	F_errmsg_internal(m, int32(_a_F_PostmasterStateMachine_8), v661)
	mBase = m.M
	v716 = m.ExcPending
	if v716 != 0 {
		goto L7
	} else {
		goto L214
	}
L212:
	;
	goto L213
L213:
	;
	v722 = *(*int32)(unsafe.Add(mBase, uint32(v672)))
	v723 = *(*int32)(unsafe.Add(mBase, uint32(v672)+4))
	*(*int32)(unsafe.Add(mBase, uint32(v722)+4)) = v723
	v725 = *(*int32)(unsafe.Add(mBase, uint32(v672)))
	*(*int32)(unsafe.Add(mBase, uint32(v723))) = v725
	F_pfree(m, v706)
	mBase = m.M
	v728 = m.ExcPending
	if v728 != 0 {
		goto L7
	} else {
		goto L216
	}
L214:
	;
	F_errfinish(m, int32(_a_F_PostmasterStateMachine_9), int32(449), int32(_a_F_PostmasterStateMachine_10))
	mBase = m.M
	v721 = m.ExcPending
	if v721 != 0 {
		goto L7
	} else {
		goto L215
	}
L215:
	;
	goto L213
L216:
	;
	goto L203
L217:
	;
	goto L202
L218:
	;
	F_LocalProcessControlFile(m)
	mBase = m.M
	v755 = m.ExcPending
	if v755 != 0 {
		goto L7
	} else {
		goto L219
	}
L219:
	;
	F_CreateSharedMemoryAndSemaphores(m)
	mBase = m.M
	v757 = m.ExcPending
	if v757 != 0 {
		goto L7
	} else {
		goto L220
	}
L220:
	;
	v760 = F_errstart(m, int32(14), int32(0))
	mBase = m.M
	v761 = m.ExcPending
	if v761 != 0 {
		goto L7
	} else {
		goto L221
	}
L221:
	;
	if v760 != 0 {
		goto L222
	} else {
		goto L223
	}
L222:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v11)+4)) = int32(_a_F_PostmasterStateMachine_25)
	v765 = *(*int32)(unsafe.Add(mBase, _c_F_PostmasterStateMachine[0]))
	v770 = *(*int32)(unsafe.Add(mBase, uint32(v765<<(uint(int32(2))%32))+uint32(_c_F_PostmasterStateMachine[3])))
	*(*int32)(unsafe.Add(mBase, uint32(v11))) = v770
	F_errmsg_internal(m, int32(_a_F_PostmasterStateMachine_2), v11)
	mBase = m.M
	v774 = m.ExcPending
	if v774 != 0 {
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
	F_maybe_adjust_io_workers(m)
	mBase = m.M
	v784 = m.ExcPending
	if v784 != 0 {
		goto L7
	} else {
		goto L227
	}
L225:
	;
	F_errfinish(m, int32(_a_F_PostmasterStateMachine_3), int32(3272), int32(_a_F_PostmasterStateMachine_4))
	mBase = m.M
	v779 = m.ExcPending
	if v779 != 0 {
		goto L7
	} else {
		goto L226
	}
L226:
	;
	goto L224
L227:
	;
	v786 = F_StartChildProcess(m, int32(13))
	mBase = m.M
	v787 = m.ExcPending
	if v787 != 0 {
		goto L7
	} else {
		goto L228
	}
L228:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_PostmasterStateMachine[13])) = int32(1)
	*(*int32)(unsafe.Add(mBase, _c_F_PostmasterStateMachine[16])) = v786
	*(*int64)(unsafe.Add(mBase, _c_F_PostmasterStateMachine[17])) = int64(0)
	F_ConfigurePostmasterWaitSet(m)
	mBase = m.M
	v797 = m.ExcPending
	if v797 != 0 {
		goto L7
	} else {
		goto L229
	}
L229:
	;
	goto L127
}
