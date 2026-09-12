package p5

import (
	base "github.com/shibukawa/pgmem/internal/aot/pgaot/base"
	"unsafe"
)

func F_check_sql_stmt_retval(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32) int32 {
	mBase := m.M
	_ = mBase
	var v6 int32
	_ = v6
	var v22 int32
	_ = v22
	var v24 int32
	_ = v24
	var v34 int32
	_ = v34
	var v37 int32
	_ = v37
	var v40 int32
	_ = v40
	var v42 int32
	_ = v42
	var v45 int32
	_ = v45
	var v49 int32
	_ = v49
	var v50 int32
	_ = v50
	var v57 int32
	_ = v57
	var v60 int32
	_ = v60
	var v61 int32
	_ = v61
	var v69 int32
	_ = v69
	var v75 int32
	_ = v75
	var v77 int32
	_ = v77
	var v78 int32
	_ = v78
	var v80 int32
	_ = v80
	var v81 int32
	_ = v81
	var v83 int32
	_ = v83
	var v84 int32
	_ = v84
	var v85 int32
	_ = v85
	var v86 int32
	_ = v86
	var v87 int32
	_ = v87
	var v88 int32
	_ = v88
	var v89 int32
	_ = v89
	var v90 int32
	_ = v90
	var v91 int32
	_ = v91
	var v92 int32
	_ = v92
	var v93 int32
	_ = v93
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
	var v101 int32
	_ = v101
	var v108 int32
	_ = v108
	var v112 int32
	_ = v112
	var v120 int32
	_ = v120
	var v126 int32
	_ = v126
	var v133 int32
	_ = v133
	var v135 int32
	_ = v135
	var v137 int32
	_ = v137
	var v145 int32
	_ = v145
	var v151 int32
	_ = v151
	var v152 int32
	_ = v152
	var v153 int32
	_ = v153
	var v154 int32
	_ = v154
	var v155 int32
	_ = v155
	var v156 int32
	_ = v156
	var v159 int32
	_ = v159
	var v170 int32
	_ = v170
	var v178 int32
	_ = v178
	var v184 int32
	_ = v184
	var v187 int32
	_ = v187
	var v188 int32
	_ = v188
	var v196 int32
	_ = v196
	var v199 int32
	_ = v199
	var v200 int32
	_ = v200
	var v201 int32
	_ = v201
	var v211 int32
	_ = v211
	var v214 int32
	_ = v214
	var v217 int32
	_ = v217
	var v218 int32
	_ = v218
	var v222 int32
	_ = v222
	var v226 int32
	_ = v226
	var v227 int32
	_ = v227
	var v230 int32
	_ = v230
	var v231 int32
	_ = v231
	var v232 int32
	_ = v232
	var v237 int32
	_ = v237
	var v239 int32
	_ = v239
	var v240 int32
	_ = v240
	var v241 int32
	_ = v241
	var v242 int32
	_ = v242
	var v245 int32
	_ = v245
	var v246 int32
	_ = v246
	var v249 int32
	_ = v249
	var v251 int32
	_ = v251
	var v253 int32
	_ = v253
	var v256 int32
	_ = v256
	var v257 int32
	_ = v257
	var v265 int32
	_ = v265
	var v269 int32
	_ = v269
	var v270 int32
	_ = v270
	var v275 int32
	_ = v275
	var v290 int32
	_ = v290
	var v291 int32
	_ = v291
	var v294 int32
	_ = v294
	var v298 int32
	_ = v298
	var v309 int32
	_ = v309
	var v310 int32
	_ = v310
	var v316 int32
	_ = v316
	var v317 int32
	_ = v317
	var v322 int32
	_ = v322
	var v325 int32
	_ = v325
	var v326 int32
	_ = v326
	var v327 int32
	_ = v327
	var v333 int32
	_ = v333
	var v334 int32
	_ = v334
	var v335 int32
	_ = v335
	var v336 int32
	_ = v336
	var v337 int32
	_ = v337
	var v338 int32
	_ = v338
	var v344 int32
	_ = v344
	var v349 int32
	_ = v349
	var v359 int32
	_ = v359
	var v360 int32
	_ = v360
	var v366 int32
	_ = v366
	var v367 int32
	_ = v367
	var v374 int32
	_ = v374
	var v378 int32
	_ = v378
	var v379 int32
	_ = v379
	var v384 int32
	_ = v384
	var v394 int32
	_ = v394
	var v401 int32
	_ = v401
	var v405 int32
	_ = v405
	var v407 int32
	_ = v407
	var v411 int32
	_ = v411
	var v412 int32
	_ = v412
	var v415 int32
	_ = v415
	var v416 int32
	_ = v416
	var v425 int32
	_ = v425
	var v431 int32
	_ = v431
	var v437 int32
	_ = v437
	var v443 int32
	_ = v443
	var v460 int32
	_ = v460
	var v462 int32
	_ = v462
	var v468 int32
	_ = v468
	var v469 int32
	_ = v469
	var v476 int32
	_ = v476
	var v479 int32
	_ = v479
	var v484 int32
	_ = v484
	var v485 int32
	_ = v485
	var v487 int32
	_ = v487
	var v490 int32
	_ = v490
	var v492 int32
	_ = v492
	var v494 int32
	_ = v494
	var v495 int32
	_ = v495
	var v496 int32
	_ = v496
	var v497 int32
	_ = v497
	var v498 int32
	_ = v498
	var v505 int32
	_ = v505
	var v506 int32
	_ = v506
	var v510 int32
	_ = v510
	var v511 int32
	_ = v511
	var v512 int32
	_ = v512
	var v517 int32
	_ = v517
	var v518 int32
	_ = v518
	var v529 int32
	_ = v529
	var v540 int32
	_ = v540
	var v543 int32
	_ = v543
	var v544 int32
	_ = v544
	var v549 int32
	_ = v549
	var v552 int32
	_ = v552
	var v553 int32
	_ = v553
	var v554 int32
	_ = v554
	var v560 int32
	_ = v560
	var v564 int32
	_ = v564
	var v569 int32
	_ = v569
	var v573 int32
	_ = v573
	var v576 int32
	_ = v576
	var v577 int32
	_ = v577
	var v578 int32
	_ = v578
	var v584 int32
	_ = v584
	var v588 int32
	_ = v588
	var v593 int32
	_ = v593
	var v597 int32
	_ = v597
	var v600 int32
	_ = v600
	var v601 int32
	_ = v601
	var v602 int32
	_ = v602
	var v608 int32
	_ = v608
	var v609 int32
	_ = v609
	var v610 int32
	_ = v610
	var v611 int32
	_ = v611
	var v612 int32
	_ = v612
	var v613 int32
	_ = v613
	var v614 int32
	_ = v614
	var v615 int32
	_ = v615
	var v616 int32
	_ = v616
	var v624 int32
	_ = v624
	var v629 int32
	_ = v629
	var v633 int32
	_ = v633
	var v636 int32
	_ = v636
	var v637 int32
	_ = v637
	var v638 int32
	_ = v638
	var v644 int32
	_ = v644
	var v649 int32
	_ = v649
	var v658 int32
	_ = v658
	var v671 int32
	_ = v671
	var v672 int32
	_ = v672
	var v673 int32
	_ = v673
	var v675 int32
	_ = v675
	var v684 int32
	_ = v684
	var __phi684 int32
	_ = __phi684
	var v686 int32
	_ = v686
	var __phi686 int32
	_ = __phi686
	var v687 int32
	_ = v687
	var __phi687 int32
	_ = __phi687
	var v694 int32
	_ = v694
	var __phi694 int32
	_ = __phi694
	var v703 int32
	_ = v703
	var v706 int32
	_ = v706
	var v709 int32
	_ = v709
	var v714 int32
	_ = v714
	var v715 int32
	_ = v715
	var v717 int32
	_ = v717
	var v720 int32
	_ = v720
	var v722 int32
	_ = v722
	var v724 int32
	_ = v724
	var v725 int32
	_ = v725
	var v726 int32
	_ = v726
	var v727 int32
	_ = v727
	var v728 int32
	_ = v728
	var v731 int32
	_ = v731
	var v733 int32
	_ = v733
	var v742 int32
	_ = v742
	var v750 int32
	_ = v750
	var v771 int32
	_ = v771
	var v779 int32
	_ = v779
	var v783 int32
	_ = v783
	var v784 int32
	_ = v784
	var v787 int32
	_ = v787
	var v788 int32
	_ = v788
	var v791 int32
	_ = v791
	var v793 int32
	_ = v793
	var v794 int32
	_ = v794
	var v797 int32
	_ = v797
	var v806 int32
	_ = v806
	var v811 int32
	_ = v811
	var v822 int32
	_ = v822
	var v826 int32
	_ = v826
	var v827 int32
	_ = v827
	var v830 int32
	_ = v830
	var v832 int32
	_ = v832
	var v833 int32
	_ = v833
	var v834 int32
	_ = v834
	var v835 int32
	_ = v835
	var v836 int32
	_ = v836
	var v838 int32
	_ = v838
	var v840 int32
	_ = v840
	var v841 int32
	_ = v841
	var v853 int32
	_ = v853
	var v865 int32
	_ = v865
	var v866 int32
	_ = v866
	var v873 int32
	_ = v873
	var v874 int32
	_ = v874
	var v877 int32
	_ = v877
	var v879 int32
	_ = v879
	var v886 int32
	_ = v886
	var v887 int32
	_ = v887
	var v890 int32
	_ = v890
	var v891 int32
	_ = v891
	var v899 int32
	_ = v899
	var v900 int32
	_ = v900
	var v902 int32
	_ = v902
	var v903 int32
	_ = v903
	var v905 int32
	_ = v905
	var v921 int32
	_ = v921
	var v936 int32
	_ = v936
	var v939 int32
	_ = v939
	var v940 int32
	_ = v940
	var v941 int32
	_ = v941
	var v947 int32
	_ = v947
	var v951 int32
	_ = v951
	var v956 int32
	_ = v956
	var v981 int32
	_ = v981
	var v984 int32
	_ = v984
	var v985 int32
	_ = v985
	var v986 int32
	_ = v986
	var v990 int32
	_ = v990
	var v994 int32
	_ = v994
	var v999 int32
	_ = v999
	v6 = int32(0)
	v22 = m.G0
	v24 = v22 - int32(160)
	m.G0 = v24
	*(*int32)(unsafe.Add(mBase, uint32(v24)+156)) = v6
	*(*uint8)(unsafe.Add(mBase, uint32(v24)+155)) = uint8(v6)
	if l1 == int32(2278) {
		v921 = v6
		goto L3
	} else {
		goto L4
	}
L1:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v981 = m.ExcPending
	if v981 != 0 {
		goto L85
	} else {
		goto L210
	}
L2:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v936 = m.ExcPending
	if v936 != 0 {
		goto L85
	} else {
		goto L204
	}
L3:
	;
	m.G0 = v24 + int32(160)
	return v921
L4:
	;
	if l0 == int32(0) {
		goto L1
	} else {
		goto L5
	}
L5:
	;
	v34 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	if v34 <= int32(0) {
		goto L7
	} else {
		goto L8
	}
L6:
	;
	if v170 == int32(0) {
		goto L1
	} else {
		goto L54
	}
L7:
	;
	v170 = v6
	v178 = v6
	goto L6
L8:
	;
	goto L9
L9:
	;
	v37 = int32(0)
	if v37 < v34 {
		goto L10
	} else {
		goto L11
	}
L10:
	;
	v40 = v34
	goto L12
L11:
	;
	v40 = v37
	goto L12
L12:
	;
	v42 = v40 & int32(3)
	if v34 < int32(4) {
		goto L14
	} else {
		goto L15
	}
L13:
	;
	if v42 == int32(0) {
		v170 = v112
		v178 = v120
		goto L6
	} else {
		goto L44
	}
L14:
	;
	v45 = int32(0)
	v108 = v45
	v112 = v45
	v120 = v6
	goto L13
L15:
	;
	goto L16
L16:
	;
	v49 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	v50 = int32(0)
	v57 = v50
	v60 = v6
	v61 = v50
	v69 = v6
	goto L17
L17:
	;
	v75 = v49 + v57<<(uint(int32(2))%32)
	v77 = v75 + int32(12)
	v78 = *(*int32)(unsafe.Add(mBase, uint32(v77)))
	v80 = v75 + int32(8)
	v81 = *(*int32)(unsafe.Add(mBase, uint32(v80)))
	v83 = v75 + int32(4)
	v84 = *(*int32)(unsafe.Add(mBase, uint32(v83)))
	v85 = *(*int32)(unsafe.Add(mBase, uint32(v75)))
	v86 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v85)+24)))
	if v86 != 0 {
		goto L19
	} else {
		goto L20
	}
L18:
	;
	v108 = v99
	v112 = v93
	v120 = v97
	goto L13
L19:
	;
	v87 = v85
	goto L21
L20:
	;
	v87 = v61
	goto L21
L21:
	;
	v88 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v84)+24)))
	if v88 != 0 {
		goto L22
	} else {
		goto L23
	}
L22:
	;
	v89 = v84
	goto L24
L23:
	;
	v89 = v87
	goto L24
L24:
	;
	v90 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v81)+24)))
	if v90 != 0 {
		goto L25
	} else {
		goto L26
	}
L25:
	;
	v91 = v81
	goto L27
L26:
	;
	v91 = v89
	goto L27
L27:
	;
	v92 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v78)+24)))
	if v92 != 0 {
		goto L28
	} else {
		goto L29
	}
L28:
	;
	v93 = v78
	goto L30
L29:
	;
	v93 = v91
	goto L30
L30:
	;
	if v86 != 0 {
		goto L31
	} else {
		goto L32
	}
L31:
	;
	v94 = v75
	goto L33
L32:
	;
	v94 = v69
	goto L33
L33:
	;
	if v88 != 0 {
		goto L34
	} else {
		goto L35
	}
L34:
	;
	v95 = v83
	goto L36
L35:
	;
	v95 = v94
	goto L36
L36:
	;
	if v90 != 0 {
		goto L37
	} else {
		goto L38
	}
L37:
	;
	v96 = v80
	goto L39
L38:
	;
	v96 = v95
	goto L39
L39:
	;
	if v92 != 0 {
		goto L40
	} else {
		goto L41
	}
L40:
	;
	v97 = v77
	goto L42
L41:
	;
	v97 = v96
	goto L42
L42:
	;
	v98 = int32(4)
	v99 = v57 + v98
	v101 = v60 + v98
	if v101 != v40&int32(2147483644) {
		v57 = v99
		v60 = v101
		v61 = v93
		v69 = v97
		goto L17
	} else {
		goto L43
	}
L43:
	;
	goto L18
L44:
	;
	v126 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	v133 = v108
	v135 = int32(0)
	v137 = v112
	v145 = v120
	goto L45
L45:
	;
	v151 = v126 + v133<<(uint(int32(2))%32)
	v152 = *(*int32)(unsafe.Add(mBase, uint32(v151)))
	v153 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v152)+24)))
	if v153 != 0 {
		goto L47
	} else {
		goto L48
	}
L46:
	;
	v170 = v154
	v178 = v155
	goto L6
L47:
	;
	v154 = v152
	goto L49
L48:
	;
	v154 = v137
	goto L49
L49:
	;
	if v153 != 0 {
		goto L50
	} else {
		goto L51
	}
L50:
	;
	v155 = v151
	goto L52
L51:
	;
	v155 = v145
	goto L52
L52:
	;
	v156 = int32(1)
	v159 = v135 + v156
	if v159 != v42 {
		v133 = v133 + v156
		v135 = v159
		v137 = v154
		v145 = v155
		goto L45
	} else {
		goto L53
	}
L53:
	;
	goto L46
L54:
	;
	v184 = *(*int32)(unsafe.Add(mBase, uint32(v170)+4))
	if v184 == int32(1) {
		goto L56
	} else {
		goto L57
	}
L55:
	;
	v201 = int32(0)
	if v199 == v201 {
		goto L62
	} else {
		goto L63
	}
L56:
	;
	v187 = *(*int32)(unsafe.Add(mBase, uint32(v170)+76))
	v188 = *(*int32)(unsafe.Add(mBase, uint32(v170)+144))
	v199 = v187
	v200 = base.B2i32(v188 == int32(0))
	goto L55
L57:
	;
	goto L58
L58:
	;
	if base.Ui32(int32(3)) < base.Ui32(v184-int32(2)) {
		goto L1
	} else {
		goto L59
	}
L59:
	;
	v196 = *(*int32)(unsafe.Add(mBase, uint32(v170)+96))
	if v196 == int32(0) {
		goto L1
	} else {
		goto L60
	}
L60:
	;
	v199 = v196
	v200 = int32(1)
	goto L55
L61:
	;
	v291 = F_get_typtype(m, l1)
	mBase = m.M
	v294 = m.ExcPending
	if v294 != 0 {
		goto L85
	} else {
		goto L86
	}
L62:
	;
	v290 = int32(0)
	goto L61
L63:
	;
	goto L64
L64:
	;
	v211 = *(*int32)(unsafe.Add(mBase, uint32(v199)+4))
	if v211 <= int32(0) {
		v275 = v201
		goto L65
	} else {
		goto L66
	}
L65:
	;
	v290 = v275
	goto L61
L66:
	;
	v214 = int32(0)
	if v214 < v211 {
		goto L67
	} else {
		goto L68
	}
L67:
	;
	v217 = v211
	goto L69
L68:
	;
	v217 = v214
	goto L69
L69:
	;
	v218 = int32(1)
	if v211 == v218 {
		goto L71
	} else {
		goto L72
	}
L70:
	;
	if v217&v218 == int32(0) {
		v275 = v256
		goto L65
	} else {
		goto L77
	}
L71:
	;
	v222 = int32(0)
	v256 = v222
	v257 = v222
	goto L70
L72:
	;
	goto L73
L73:
	;
	v226 = *(*int32)(unsafe.Add(mBase, uint32(v199)+12))
	v227 = int32(0)
	v230 = v227
	v231 = v227
	v232 = v201
	goto L74
L74:
	;
	v237 = int32(2)
	v239 = v226 + v231<<(uint(v237)%32)
	v240 = *(*int32)(unsafe.Add(mBase, uint32(v239)))
	v241 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v240)+26)))
	v242 = int32(1)
	v245 = *(*int32)(unsafe.Add(mBase, uint32(v239)+4))
	v246 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v245)+26)))
	v249 = v230 + (v241 ^ v242) + (v246 ^ v242)
	v251 = v231 + v237
	v253 = v232 + v237
	if v253 != v217&int32(2147483646) {
		v230 = v249
		v231 = v251
		v232 = v253
		goto L74
	} else {
		goto L76
	}
L75:
	;
	v256 = v249
	v257 = v251
	goto L70
L76:
	;
	goto L75
L77:
	;
	v265 = *(*int32)(unsafe.Add(mBase, uint32(v199)+12))
	v269 = *(*int32)(unsafe.Add(mBase, uint32(v265+v257<<(uint(int32(2))%32))))
	v270 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v269)+26)))
	v275 = v256 + (v270 ^ int32(1))
	goto L65
L78:
	;
	v779 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v24)+155)))
	if v779 != int32(1) {
		v921 = v771
		goto L3
	} else {
		goto L182
	}
L79:
	;
	v671 = int32(1)
	v672 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v24)+155)))
	v673 = *(*int32)(unsafe.Add(mBase, uint32(v24)+156))
	v675 = v658 + v671
	if v675 <= v374 {
		goto L166
	} else {
		goto L167
	}
L80:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v633 = m.ExcPending
	if v633 != 0 {
		goto L85
	} else {
		goto L161
	}
L81:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v597 = m.ExcPending
	if v597 != 0 {
		goto L85
	} else {
		goto L152
	}
L82:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v573 = m.ExcPending
	if v573 != 0 {
		goto L85
	} else {
		goto L146
	}
L83:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v549 = m.ExcPending
	if v549 != 0 {
		goto L85
	} else {
		goto L140
	}
L84:
	;
	if base.B2i32(l1 != int32(2249))&base.B2i32(v291 != int32(99)) != 0 {
		goto L80
	} else {
		goto L102
	}
L85:
	;
	return int32(0)
L86:
	;
	v298 = v291&int32(255) - int32(98)
	if base.Ui32(int32(16)) < base.Ui32(v298) {
		goto L84
	} else {
		goto L87
	}
L87:
	;
	if int32(1)<<(uint(v298)%32)&int32(67597) == int32(0) {
		goto L84
	} else {
		goto L88
	}
L88:
	;
	if v290 != int32(1) {
		goto L83
	} else {
		goto L89
	}
L89:
	;
	v309 = *(*int32)(unsafe.Add(mBase, uint32(v199)+12))
	v310 = *(*int32)(unsafe.Add(mBase, uint32(v309)))
	v316 = F_coerce_fn_result_column(m, v310, l1, int32(-1), v200, v24+int32(156), v24+int32(155))
	mBase = m.M
	v317 = m.ExcPending
	if v317 != 0 {
		goto L85
	} else {
		goto L90
	}
L90:
	;
	if v316 != 0 {
		goto L91
	} else {
		goto L92
	}
L91:
	;
	v771 = int32(0)
	goto L78
L92:
	;
	goto L93
L93:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v322 = m.ExcPending
	if v322 != 0 {
		goto L85
	} else {
		goto L94
	}
L94:
	;
	F_errcode(m, int32(50724996))
	mBase = m.M
	v325 = m.ExcPending
	if v325 != 0 {
		goto L85
	} else {
		goto L95
	}
L95:
	;
	v326 = F_format_type_be(m, l1)
	mBase = m.M
	v327 = m.ExcPending
	if v327 != 0 {
		goto L85
	} else {
		goto L96
	}
L96:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v24)+112)) = v326
	F_errmsg(m, int32(193997), v24+int32(112))
	mBase = m.M
	v333 = m.ExcPending
	if v333 != 0 {
		goto L85
	} else {
		goto L97
	}
L97:
	;
	v334 = *(*int32)(unsafe.Add(mBase, uint32(v310)+4))
	v335 = F_exprType(m, v334)
	mBase = m.M
	v336 = m.ExcPending
	if v336 != 0 {
		goto L85
	} else {
		goto L98
	}
L98:
	;
	v337 = F_format_type_be(m, v335)
	mBase = m.M
	v338 = m.ExcPending
	if v338 != 0 {
		goto L85
	} else {
		goto L99
	}
L99:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v24)+96)) = v337
	F_errdetail(m, int32(632526), v24+int32(96))
	mBase = m.M
	v344 = m.ExcPending
	if v344 != 0 {
		goto L85
	} else {
		goto L100
	}
L100:
	;
	F_errfinish(m, int32(517657), int32(2279), int32(323771))
	mBase = m.M
	v349 = m.ExcPending
	if v349 != 0 {
		goto L85
	} else {
		goto L101
	}
L101:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L102:
	;
	if l3 == int32(112) {
		goto L103
	} else {
		goto L104
	}
L103:
	;
	if l2 == int32(0) {
		goto L108
	} else {
		goto L109
	}
L104:
	;
	if v290 != int32(1) {
		goto L103
	} else {
		goto L105
	}
L105:
	;
	v359 = *(*int32)(unsafe.Add(mBase, uint32(v199)+12))
	v360 = *(*int32)(unsafe.Add(mBase, uint32(v359)))
	v366 = F_coerce_fn_result_column(m, v360, l1, int32(-1), v200, v24+int32(156), v24+int32(155))
	mBase = m.M
	v367 = m.ExcPending
	if v367 != 0 {
		goto L85
	} else {
		goto L106
	}
L106:
	;
	if v366 == int32(0) {
		goto L103
	} else {
		goto L107
	}
L107:
	;
	v771 = int32(0)
	goto L78
L108:
	;
	v921 = int32(1)
	goto L3
L109:
	;
	goto L110
L110:
	;
	v374 = *(*int32)(unsafe.Add(mBase, uint32(l2)))
	if v199 == int32(0) {
		goto L111
	} else {
		goto L112
	}
L111:
	;
	v658 = int32(0)
	goto L79
L112:
	;
	goto L113
L113:
	;
	v378 = int32(0)
	v379 = *(*int32)(unsafe.Add(mBase, uint32(v199)+4))
	if v379 <= v378 {
		v658 = v378
		goto L79
	} else {
		goto L114
	}
L114:
	;
	v384 = int32(0)
	v394 = v378
	v401 = v384
	v405 = v384
	goto L115
L115:
	;
	v407 = *(*int32)(unsafe.Add(mBase, uint32(v199)+12))
	v411 = *(*int32)(unsafe.Add(mBase, uint32(v407+v401<<(uint(int32(2))%32))))
	v412 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v411)+26)))
	if v412 == int32(0) {
		goto L117
	} else {
		goto L118
	}
L116:
	;
	v658 = v529
	goto L79
L117:
	;
	v415 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v24)+155)))
	v416 = *(*int32)(unsafe.Add(mBase, uint32(v24)+156))
	v425 = v394
	v431 = v416
	v437 = v415
	goto L120
L118:
	;
	v529 = v394
	v540 = v405
	goto L119
L119:
	;
	v543 = v401 + int32(1)
	v544 = *(*int32)(unsafe.Add(mBase, uint32(v199)+4))
	if v543 < v544 {
		v394 = v529
		v401 = v543
		v405 = v540
		goto L115
	} else {
		goto L139
	}
L120:
	;
	v443 = v425
	goto L122
L121:
	;
	*(*uint8)(unsafe.Add(mBase, uint32(v24)+155)) = uint8(v506)
	*(*int32)(unsafe.Add(mBase, uint32(v24)+156)) = v505
	v510 = v405 + int32(1)
	v511 = *(*int32)(unsafe.Add(mBase, uint32(v468)+68))
	v512 = *(*int32)(unsafe.Add(mBase, uint32(v468)+76))
	v517 = F_coerce_fn_result_column(m, v411, v511, v512, v200, v24+int32(156), v24+int32(155))
	mBase = m.M
	v518 = m.ExcPending
	if v518 != 0 {
		goto L85
	} else {
		goto L137
	}
L122:
	;
	v460 = v443 + int32(1)
	if v374 < v460 {
		goto L82
	} else {
		goto L124
	}
L123:
	;
	goto L121
L124:
	;
	v462 = *(*int32)(unsafe.Add(mBase, uint32(l2)))
	v468 = l2 + int32(20) + v462<<(uint(int32(4))%32) + v443*int32(100)
	v469 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v468)+91)))
	if l4 == int32(0) {
		goto L126
	} else {
		goto L127
	}
L125:
	;
	goto L123
L126:
	;
	if v469&int32(1) != 0 {
		v443 = v460
		goto L122
	} else {
		goto L136
	}
L127:
	;
	if v469&int32(1) == int32(0) {
		goto L126
	} else {
		goto L128
	}
L128:
	;
	v476 = int32(1)
	v479 = int32(0)
	v484 = F_makeConst(m, int32(23), int32(-1), v479, int32(4), v479, v476, v476)
	mBase = m.M
	v485 = m.ExcPending
	if v485 != 0 {
		goto L85
	} else {
		goto L129
	}
L129:
	;
	if v431 != 0 {
		goto L130
	} else {
		goto L131
	}
L130:
	;
	v487 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v431)+4)))
	v490 = v487 + int32(1)
	goto L132
L131:
	;
	v490 = int32(1)
	goto L132
L132:
	;
	v492 = int32(0)
	v494 = F_makeTargetEntry(m, v484, base.I32_extend16_s(v490), v492, v492)
	mBase = m.M
	v495 = m.ExcPending
	if v495 != 0 {
		goto L85
	} else {
		goto L133
	}
L133:
	;
	v496 = F_lappend(m, v431, v494)
	mBase = m.M
	v497 = m.ExcPending
	if v497 != 0 {
		goto L85
	} else {
		goto L134
	}
L134:
	;
	v498 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v468)+91)))
	if v498&int32(1) != 0 {
		v425 = v460
		v431 = v496
		v437 = v476
		goto L120
	} else {
		goto L135
	}
L135:
	;
	v505 = v496
	v506 = v476
	goto L125
L136:
	;
	v505 = v431
	v506 = v437
	goto L125
L137:
	;
	if v517 == int32(0) {
		goto L81
	} else {
		goto L138
	}
L138:
	;
	v529 = v460
	v540 = v510
	goto L119
L139:
	;
	goto L116
L140:
	;
	F_errcode(m, int32(50724996))
	mBase = m.M
	v552 = m.ExcPending
	if v552 != 0 {
		goto L85
	} else {
		goto L141
	}
L141:
	;
	v553 = F_format_type_be(m, l1)
	mBase = m.M
	v554 = m.ExcPending
	if v554 != 0 {
		goto L85
	} else {
		goto L142
	}
L142:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v24)+128)) = v553
	F_errmsg(m, int32(193997), v24+int32(128))
	mBase = m.M
	v560 = m.ExcPending
	if v560 != 0 {
		goto L85
	} else {
		goto L143
	}
L143:
	;
	F_errdetail(m, int32(647592), int32(0))
	mBase = m.M
	v564 = m.ExcPending
	if v564 != 0 {
		goto L85
	} else {
		goto L144
	}
L144:
	;
	F_errfinish(m, int32(517657), int32(2264), int32(323771))
	mBase = m.M
	v569 = m.ExcPending
	if v569 != 0 {
		goto L85
	} else {
		goto L145
	}
L145:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L146:
	;
	F_errcode(m, int32(50724996))
	mBase = m.M
	v576 = m.ExcPending
	if v576 != 0 {
		goto L85
	} else {
		goto L147
	}
L147:
	;
	v577 = F_format_type_be(m, l1)
	mBase = m.M
	v578 = m.ExcPending
	if v578 != 0 {
		goto L85
	} else {
		goto L148
	}
L148:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v24)+32)) = v577
	F_errmsg(m, int32(193997), v24+int32(32))
	mBase = m.M
	v584 = m.ExcPending
	if v584 != 0 {
		goto L85
	} else {
		goto L149
	}
L149:
	;
	F_errdetail(m, int32(620349), int32(0))
	mBase = m.M
	v588 = m.ExcPending
	if v588 != 0 {
		goto L85
	} else {
		goto L150
	}
L150:
	;
	F_errfinish(m, int32(517657), int32(2368), int32(323771))
	mBase = m.M
	v593 = m.ExcPending
	if v593 != 0 {
		goto L85
	} else {
		goto L151
	}
L151:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L152:
	;
	F_errcode(m, int32(50724996))
	mBase = m.M
	v600 = m.ExcPending
	if v600 != 0 {
		goto L85
	} else {
		goto L153
	}
L153:
	;
	v601 = F_format_type_be(m, l1)
	mBase = m.M
	v602 = m.ExcPending
	if v602 != 0 {
		goto L85
	} else {
		goto L154
	}
L154:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v24)+64)) = v601
	F_errmsg(m, int32(193997), v24-int32(-64))
	mBase = m.M
	v608 = m.ExcPending
	if v608 != 0 {
		goto L85
	} else {
		goto L155
	}
L155:
	;
	v609 = *(*int32)(unsafe.Add(mBase, uint32(v411)+4))
	v610 = F_exprType(m, v609)
	mBase = m.M
	v611 = m.ExcPending
	if v611 != 0 {
		goto L85
	} else {
		goto L156
	}
L156:
	;
	v612 = F_format_type_be(m, v610)
	mBase = m.M
	v613 = m.ExcPending
	if v613 != 0 {
		goto L85
	} else {
		goto L157
	}
L157:
	;
	v614 = *(*int32)(unsafe.Add(mBase, uint32(v468)+68))
	v615 = F_format_type_be(m, v614)
	mBase = m.M
	v616 = m.ExcPending
	if v616 != 0 {
		goto L85
	} else {
		goto L158
	}
L158:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v24)+56)) = v510
	*(*int32)(unsafe.Add(mBase, uint32(v24)+52)) = v615
	*(*int32)(unsafe.Add(mBase, uint32(v24)+48)) = v612
	F_errdetail(m, int32(683123), v24+int32(48))
	mBase = m.M
	v624 = m.ExcPending
	if v624 != 0 {
		goto L85
	} else {
		goto L159
	}
L159:
	;
	F_errfinish(m, int32(517657), int32(2404), int32(323771))
	mBase = m.M
	v629 = m.ExcPending
	if v629 != 0 {
		goto L85
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
	F_errcode(m, int32(50724996))
	mBase = m.M
	v636 = m.ExcPending
	if v636 != 0 {
		goto L85
	} else {
		goto L162
	}
L162:
	;
	v637 = F_format_type_be(m, l1)
	mBase = m.M
	v638 = m.ExcPending
	if v638 != 0 {
		goto L85
	} else {
		goto L163
	}
L163:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v24)+80)) = v637
	F_errmsg(m, int32(150020), v24+int32(80))
	mBase = m.M
	v644 = m.ExcPending
	if v644 != 0 {
		goto L85
	} else {
		goto L164
	}
L164:
	;
	F_errfinish(m, int32(517657), int32(2444), int32(323771))
	mBase = m.M
	v649 = m.ExcPending
	if v649 != 0 {
		goto L85
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
	__phi684 = v675
	__phi686 = v673
	__phi687 = v658
	__phi694 = v672
	v684 = __phi684
	v686 = __phi686
	v687 = __phi687
	v694 = __phi694
	goto L169
L167:
	;
	v742 = v673
	v750 = v672
	goto L168
L168:
	;
	*(*uint8)(unsafe.Add(mBase, uint32(v24)+155)) = uint8(v750)
	*(*int32)(unsafe.Add(mBase, uint32(v24)+156)) = v742
	v771 = v671
	goto L78
L169:
	;
	v703 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l2+int32(29)+v687<<(uint(int32(4))%32)))))
	if v703 == int32(0) {
		goto L2
	} else {
		goto L171
	}
L170:
	;
	v742 = v728
	v750 = v731
	goto L168
L171:
	;
	if l4 != 0 {
		goto L172
	} else {
		goto L173
	}
L172:
	;
	v706 = int32(1)
	v709 = int32(0)
	v714 = F_makeConst(m, int32(23), int32(-1), v709, int32(4), v709, v706, v706)
	mBase = m.M
	v715 = m.ExcPending
	if v715 != 0 {
		goto L85
	} else {
		goto L175
	}
L173:
	;
	v728 = v686
	v731 = v694
	goto L174
L174:
	;
	v733 = v684 + int32(1)
	if v733 <= v374 {
		__phi684 = v733
		__phi686 = v728
		__phi687 = v684
		__phi694 = v731
		v684 = __phi684
		v686 = __phi686
		v687 = __phi687
		v694 = __phi694
		goto L169
	} else {
		goto L181
	}
L175:
	;
	if v686 != 0 {
		goto L176
	} else {
		goto L177
	}
L176:
	;
	v717 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v686)+4)))
	v720 = v717 + int32(1)
	goto L178
L177:
	;
	v720 = int32(1)
	goto L178
L178:
	;
	v722 = int32(0)
	v724 = F_makeTargetEntry(m, v714, base.I32_extend16_s(v720), v722, v722)
	mBase = m.M
	v725 = m.ExcPending
	if v725 != 0 {
		goto L85
	} else {
		goto L179
	}
L179:
	;
	v726 = F_lappend(m, v686, v724)
	mBase = m.M
	v727 = m.ExcPending
	if v727 != 0 {
		goto L85
	} else {
		goto L180
	}
L180:
	;
	v728 = v726
	v731 = v706
	goto L174
L181:
	;
	goto L170
L182:
	;
	v783 = F_palloc0(m, int32(168))
	mBase = m.M
	v784 = m.ExcPending
	if v784 != 0 {
		goto L85
	} else {
		goto L183
	}
L183:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v783))) = int64(4294967363)
	v787 = *(*int32)(unsafe.Add(mBase, uint32(v170)+8))
	v788 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(v783)+24)) = uint8(v788)
	*(*int32)(unsafe.Add(mBase, uint32(v783)+8)) = v787
	v791 = *(*int32)(unsafe.Add(mBase, uint32(v24)+156))
	*(*int32)(unsafe.Add(mBase, uint32(v783)+76)) = v791
	v793 = int32(0)
	v794 = *(*int32)(unsafe.Add(mBase, uint32(v170)+76))
	if v794 == v793 {
		v853 = v793
		goto L184
	} else {
		goto L185
	}
L184:
	;
	v865 = F_palloc0(m, int32(136))
	mBase = m.M
	v866 = m.ExcPending
	if v866 != 0 {
		goto L85
	} else {
		goto L198
	}
L185:
	;
	v797 = *(*int32)(unsafe.Add(mBase, uint32(v794)+4))
	if v797 <= int32(0) {
		v853 = v793
		goto L184
	} else {
		goto L186
	}
L186:
	;
	v806 = int32(0)
	v811 = v793
	goto L187
L187:
	;
	v822 = *(*int32)(unsafe.Add(mBase, uint32(v794)+12))
	v826 = *(*int32)(unsafe.Add(mBase, uint32(v822+v806<<(uint(int32(2))%32))))
	v827 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v826)+26)))
	if v827 == int32(0) {
		goto L189
	} else {
		goto L190
	}
L188:
	;
	v853 = v838
	goto L184
L189:
	;
	v830 = *(*int32)(unsafe.Add(mBase, uint32(v826)+12))
	if v830 != 0 {
		goto L192
	} else {
		goto L193
	}
L190:
	;
	v838 = v811
	goto L191
L191:
	;
	v840 = v806 + int32(1)
	v841 = *(*int32)(unsafe.Add(mBase, uint32(v794)+4))
	if v840 < v841 {
		v806 = v840
		v811 = v838
		goto L187
	} else {
		goto L197
	}
L192:
	;
	v832 = v830
	goto L194
L193:
	;
	v832 = int32(791891)
	goto L194
L194:
	;
	v833 = F_makeString(m, v832)
	mBase = m.M
	v834 = m.ExcPending
	if v834 != 0 {
		goto L85
	} else {
		goto L195
	}
L195:
	;
	v835 = F_lappend(m, v811, v833)
	mBase = m.M
	v836 = m.ExcPending
	if v836 != 0 {
		goto L85
	} else {
		goto L196
	}
L196:
	;
	v838 = v835
	goto L191
L197:
	;
	goto L188
L198:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v865)+36)) = v170
	*(*int32)(unsafe.Add(mBase, uint32(v865)+12)) = int32(1)
	*(*int32)(unsafe.Add(mBase, uint32(v865))) = int32(101)
	v873 = F_makeAlias(m, int32(700641), v853)
	mBase = m.M
	v874 = m.ExcPending
	if v874 != 0 {
		goto L85
	} else {
		goto L199
	}
L199:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v865)+8)) = v873
	*(*int32)(unsafe.Add(mBase, uint32(v865)+4)) = v873
	v877 = int32(256)
	*(*uint16)(unsafe.Add(mBase, uint32(v865)+124)) = uint16(v877)
	v879 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v865)+20)) = uint8(v879)
	*(*int32)(unsafe.Add(mBase, uint32(v24)+12)) = v865
	*(*int32)(unsafe.Add(mBase, uint32(v24)+148)) = v865
	v886 = F_list_make1_impl(m, int32(1), v24+int32(12))
	mBase = m.M
	v887 = m.ExcPending
	if v887 != 0 {
		goto L85
	} else {
		goto L200
	}
L200:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v783)+52)) = v886
	v890 = F_palloc0(m, int32(8))
	mBase = m.M
	v891 = m.ExcPending
	if v891 != 0 {
		goto L85
	} else {
		goto L201
	}
L201:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v890))) = int64(4294967359)
	*(*int32)(unsafe.Add(mBase, uint32(v24)+8)) = v890
	*(*int32)(unsafe.Add(mBase, uint32(v24)+144)) = v890
	v899 = F_list_make1_impl(m, int32(1), v24+int32(8))
	mBase = m.M
	v900 = m.ExcPending
	if v900 != 0 {
		goto L85
	} else {
		goto L202
	}
L202:
	;
	v902 = F_makeFromExpr(m, v899, int32(0))
	mBase = m.M
	v903 = m.ExcPending
	if v903 != 0 {
		goto L85
	} else {
		goto L203
	}
L203:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v783)+60)) = v902
	v905 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v170)+44)))
	*(*uint8)(unsafe.Add(mBase, uint32(v783)+44)) = uint8(v905)
	*(*int32)(unsafe.Add(mBase, uint32(v178))) = v783
	v921 = v771
	goto L3
L204:
	;
	F_errcode(m, int32(50724996))
	mBase = m.M
	v939 = m.ExcPending
	if v939 != 0 {
		goto L85
	} else {
		goto L205
	}
L205:
	;
	v940 = F_format_type_be(m, l1)
	mBase = m.M
	v941 = m.ExcPending
	if v941 != 0 {
		goto L85
	} else {
		goto L206
	}
L206:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v24)+16)) = v940
	F_errmsg(m, int32(193997), v24+int32(16))
	mBase = m.M
	v947 = m.ExcPending
	if v947 != 0 {
		goto L85
	} else {
		goto L207
	}
L207:
	;
	F_errdetail(m, int32(620465), int32(0))
	mBase = m.M
	v951 = m.ExcPending
	if v951 != 0 {
		goto L85
	} else {
		goto L208
	}
L208:
	;
	F_errfinish(m, int32(517657), int32(2415), int32(323771))
	mBase = m.M
	v956 = m.ExcPending
	if v956 != 0 {
		goto L85
	} else {
		goto L209
	}
L209:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L210:
	;
	F_errcode(m, int32(50724996))
	mBase = m.M
	v984 = m.ExcPending
	if v984 != 0 {
		goto L85
	} else {
		goto L211
	}
L211:
	;
	v985 = F_format_type_be(m, l1)
	mBase = m.M
	v986 = m.ExcPending
	if v986 != 0 {
		goto L85
	} else {
		goto L212
	}
L212:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v24))) = v985
	F_errmsg(m, int32(193997), v24)
	mBase = m.M
	v990 = m.ExcPending
	if v990 != 0 {
		goto L85
	} else {
		goto L213
	}
L213:
	;
	F_errdetail(m, int32(688784), int32(0))
	mBase = m.M
	v994 = m.ExcPending
	if v994 != 0 {
		goto L85
	} else {
		goto L214
	}
L214:
	;
	F_errfinish(m, int32(517657), int32(2226), int32(323771))
	mBase = m.M
	v999 = m.ExcPending
	if v999 != 0 {
		goto L85
	} else {
		goto L215
	}
L215:
	;
	base.Wasm_trap_unreachable()
	for {
	}
}
func F_map_sql_table_to_xmlschema(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) int32 {
	mBase := m.M
	_ = mBase
	var v6 int32
	_ = v6
	var v8 int32
	_ = v8
	var v15 int32
	_ = v15
	var v17 int32
	_ = v17
	var v18 int32
	_ = v18
	var v21 int32
	_ = v21
	var v22 int32
	_ = v22
	var v27 int32
	_ = v27
	var v28 int32
	_ = v28
	var v36 int32
	_ = v36
	var v41 int32
	_ = v41
	var v49 int32
	_ = v49
	var v50 int32
	_ = v50
	var v51 int32
	_ = v51
	var v52 int32
	_ = v52
	var v54 int32
	_ = v54
	var v63 int32
	_ = v63
	var v64 int32
	_ = v64
	var v73 int32
	_ = v73
	var v81 int32
	_ = v81
	var v83 int32
	_ = v83
	var v85 int32
	_ = v85
	var v86 int32
	_ = v86
	var v96 int32
	_ = v96
	var v109 int32
	_ = v109
	var v120 int32
	_ = v120
	var v129 int32
	_ = v129
	var v134 int32
	_ = v134
	var v135 int32
	_ = v135
	var v143 int32
	_ = v143
	var v149 int32
	_ = v149
	var v154 int32
	_ = v154
	v6 = m.G0
	v8 = v6 - int32(112)
	m.G0 = v8
	F_initStringInfo(m, v8+int32(96))
	mBase = m.M
	v15 = m.ExcPending
	if v15 != 0 {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	return int32(0)
L2:
	;
	if l1 != 0 {
		goto L4
	} else {
		goto L5
	}
L3:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v143 = m.ExcPending
	if v143 != 0 {
		goto L1
	} else {
		goto L38
	}
L4:
	;
	v17 = F_SearchSysCache1(m, int32(57), l1)
	mBase = m.M
	v18 = m.ExcPending
	if v18 != 0 {
		goto L1
	} else {
		goto L7
	}
L5:
	;
	goto L6
L6:
	;
	F_appendStringInfoString(m, v8+int32(96), int32(764322))
	mBase = m.M
	v27 = m.ExcPending
	if v27 != 0 {
		goto L1
	} else {
		goto L10
	}
L7:
	;
	if v17 == int32(0) {
		goto L3
	} else {
		goto L8
	}
L8:
	;
	v21 = F_map_sql_identifier_to_xml_name(m)
	mBase = m.M
	v22 = m.ExcPending
	if v22 != 0 {
		goto L1
	} else {
		goto L9
	}
L9:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L10:
	;
	v28 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l3))))
	if v28 != 0 {
		goto L11
	} else {
		goto L12
	}
L11:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v8)+64)) = l3
	F_appendStringInfo(m, v8+int32(96), int32(763910), v8-int32(-64))
	mBase = m.M
	v36 = m.ExcPending
	if v36 != 0 {
		goto L1
	} else {
		goto L14
	}
L12:
	;
	goto L13
L13:
	;
	F_appendStringInfoString(m, v8+int32(96), int32(791682))
	mBase = m.M
	v41 = m.ExcPending
	if v41 != 0 {
		goto L1
	} else {
		goto L15
	}
L14:
	;
	goto L13
L15:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v8)+60)) = l0
	*(*int32)(unsafe.Add(mBase, uint32(v8)+92)) = l0
	v49 = F_list_make1_impl(m, int32(1), v8+int32(60))
	mBase = m.M
	v50 = m.ExcPending
	if v50 != 0 {
		goto L1
	} else {
		goto L16
	}
L16:
	;
	v51 = F_map_sql_typecoll_to_xmlschema_types(m, v49)
	mBase = m.M
	v52 = m.ExcPending
	if v52 != 0 {
		goto L1
	} else {
		goto L17
	}
L17:
	;
	F_appendStringInfoString(m, v8+int32(96), v51)
	mBase = m.M
	v54 = m.ExcPending
	if v54 != 0 {
		goto L1
	} else {
		goto L18
	}
L18:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v8)+48)) = int32(389828)
	F_appendStringInfo(m, v8+int32(96), int32(788347), v8+int32(48))
	mBase = m.M
	v63 = m.ExcPending
	if v63 != 0 {
		goto L1
	} else {
		goto L19
	}
L19:
	;
	v64 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	if v64 <= int32(0) {
		goto L20
	} else {
		goto L21
	}
L20:
	;
	F_appendStringInfoString(m, v8+int32(96), int32(791610))
	mBase = m.M
	v96 = m.ExcPending
	if v96 != 0 {
		goto L1
	} else {
		goto L29
	}
L21:
	;
	v73 = int32(0)
	goto L22
L22:
	;
	v81 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0+v64<<(uint(int32(4))%32)+int32(20)+v73*int32(100))+91)))
	if v81 != 0 {
		goto L24
	} else {
		goto L25
	}
L23:
	;
	v85 = F_map_sql_identifier_to_xml_name(m)
	mBase = m.M
	v86 = m.ExcPending
	if v86 != 0 {
		goto L1
	} else {
		goto L28
	}
L24:
	;
	v83 = v73 + int32(1)
	if v64 != v83 {
		v73 = v83
		goto L22
	} else {
		goto L27
	}
L25:
	;
	goto L26
L26:
	;
	goto L23
L27:
	;
	goto L20
L28:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L29:
	;
	if l2 == int32(0) {
		goto L31
	} else {
		goto L32
	}
L30:
	;
	F_appendStringInfoString(m, v8+int32(96), int32(573078))
	mBase = m.M
	v134 = m.ExcPending
	if v134 != 0 {
		goto L1
	} else {
		goto L37
	}
L31:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v8)+36)) = int32(389828)
	*(*int32)(unsafe.Add(mBase, uint32(v8)+32)) = int32(390088)
	F_appendStringInfo(m, v8+int32(96), int32(791489), v8+int32(32))
	mBase = m.M
	v109 = m.ExcPending
	if v109 != 0 {
		goto L1
	} else {
		goto L34
	}
L32:
	;
	goto L33
L33:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v8)+4)) = int32(389828)
	*(*int32)(unsafe.Add(mBase, uint32(v8))) = int32(32320)
	F_appendStringInfo(m, v8+int32(96), int32(791649), v8)
	mBase = m.M
	v129 = m.ExcPending
	if v129 != 0 {
		goto L1
	} else {
		goto L36
	}
L34:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v8)+20)) = int32(390088)
	*(*int32)(unsafe.Add(mBase, uint32(v8)+16)) = int32(414263)
	F_appendStringInfo(m, v8+int32(96), int32(791649), v8+int32(16))
	mBase = m.M
	v120 = m.ExcPending
	if v120 != 0 {
		goto L1
	} else {
		goto L35
	}
L35:
	;
	goto L30
L36:
	;
	goto L30
L37:
	;
	v135 = *(*int32)(unsafe.Add(mBase, uint32(v8)+96))
	m.G0 = v8 + int32(112)
	return v135
L38:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v8)+80)) = l1
	F_errmsg_internal(m, int32(50028), v8+int32(80))
	mBase = m.M
	v149 = m.ExcPending
	if v149 != 0 {
		goto L1
	} else {
		goto L39
	}
L39:
	;
	F_errfinish(m, int32(522311), int32(3533), int32(531893))
	mBase = m.M
	v154 = m.ExcPending
	if v154 != 0 {
		goto L1
	} else {
		goto L40
	}
L40:
	;
	base.Wasm_trap_unreachable()
	for {
	}
}
func F_map_sql_typecoll_to_xmlschema_types(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v2 int32
	_ = v2
	var v10 int32
	_ = v10
	var v12 int32
	_ = v12
	var v16 int32
	_ = v16
	var v22 int32
	_ = v22
	var v26 int32
	_ = v26
	var v28 int32
	_ = v28
	var v32 int32
	_ = v32
	var v33 int32
	_ = v33
	var v41 int32
	_ = v41
	var v42 int32
	_ = v42
	var v44 int32
	_ = v44
	var v53 int32
	_ = v53
	var v54 int32
	_ = v54
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
	var v64 int32
	_ = v64
	var v66 int32
	_ = v66
	var v71 int32
	_ = v71
	var v78 int32
	_ = v78
	var v79 int32
	_ = v79
	var v84 int32
	_ = v84
	var v92 int32
	_ = v92
	var v93 int32
	_ = v93
	var v98 int32
	_ = v98
	var v100 int32
	_ = v100
	var v105 int32
	_ = v105
	var v109 int32
	_ = v109
	var v110 int32
	_ = v110
	var v111 int32
	_ = v111
	var v113 int32
	_ = v113
	var v114 int32
	_ = v114
	var v115 int32
	_ = v115
	var v117 int32
	_ = v117
	var v118 int32
	_ = v118
	var v132 int32
	_ = v132
	var v137 int32
	_ = v137
	var v145 int32
	_ = v145
	var v148 int32
	_ = v148
	var v149 int32
	_ = v149
	var v157 int32
	_ = v157
	var v161 int32
	_ = v161
	var v165 int32
	_ = v165
	var v167 int32
	_ = v167
	var v168 int32
	_ = v168
	var v172 int32
	_ = v172
	var v183 int32
	_ = v183
	var v184 int32
	_ = v184
	var v210 int32
	_ = v210
	var v218 int32
	_ = v218
	var v223 int32
	_ = v223
	var v224 int32
	_ = v224
	var v231 int32
	_ = v231
	var v232 int32
	_ = v232
	var v233 int32
	_ = v233
	var v234 int32
	_ = v234
	var v235 int32
	_ = v235
	var v243 int32
	_ = v243
	var v248 int32
	_ = v248
	var v253 int32
	_ = v253
	var v257 int32
	_ = v257
	var v258 int32
	_ = v258
	var v266 int32
	_ = v266
	var v275 int32
	_ = v275
	var v284 int32
	_ = v284
	var v295 int32
	_ = v295
	var v300 int32
	_ = v300
	var v305 int32
	_ = v305
	var v310 int32
	_ = v310
	var v315 int32
	_ = v315
	var v323 int32
	_ = v323
	var v328 int32
	_ = v328
	var v329 int32
	_ = v329
	var v333 int32
	_ = v333
	var v334 int32
	_ = v334
	var v340 int32
	_ = v340
	var v342 int32
	_ = v342
	var v343 int32
	_ = v343
	var v354 int32
	_ = v354
	v2 = int32(0)
	v10 = m.G0
	v12 = v10 - int32(176)
	m.G0 = v12
	if l0 == v2 {
		goto L3
	} else {
		goto L4
	}
L1:
	;
	v354 = *(*int32)(unsafe.Add(mBase, uint32(v12)+140))
	m.G0 = v12 + int32(176)
	return v354
L2:
	;
	F_initStringInfo(m, v12+int32(140))
	mBase = m.M
	v145 = m.ExcPending
	if v145 != 0 {
		goto L19
	} else {
		goto L36
	}
L3:
	;
	F_initStringInfo(m, v12+int32(140))
	mBase = m.M
	v132 = m.ExcPending
	if v132 != 0 {
		goto L19
	} else {
		goto L35
	}
L4:
	;
	v16 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	if v16 <= int32(0) {
		goto L6
	} else {
		goto L7
	}
L5:
	;
	if v84 == int32(0) {
		goto L3
	} else {
		goto L23
	}
L6:
	;
	v84 = v2
	goto L5
L7:
	;
	goto L8
L8:
	;
	v22 = v2
	v26 = v2
	goto L9
L9:
	;
	v28 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	v32 = *(*int32)(unsafe.Add(mBase, uint32(v28+v26<<(uint(int32(2))%32))))
	v33 = *(*int32)(unsafe.Add(mBase, uint32(v32)))
	if int32(0) < v33 {
		goto L11
	} else {
		goto L12
	}
L10:
	;
	v84 = v71
	goto L5
L11:
	;
	v41 = int32(0)
	v42 = v22
	v44 = v33
	goto L14
L12:
	;
	v71 = v22
	goto L13
L13:
	;
	v78 = v26 + int32(1)
	v79 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	if v78 < v79 {
		v22 = v71
		v26 = v78
		goto L9
	} else {
		goto L22
	}
L14:
	;
	v53 = v32 + int32(20) + v44<<(uint(int32(4))%32) + v41*int32(100)
	v54 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v53)+91)))
	if v54 == int32(0) {
		goto L16
	} else {
		goto L17
	}
L15:
	;
	v71 = v63
	goto L13
L16:
	;
	v57 = *(*int32)(unsafe.Add(mBase, uint32(v53)+68))
	v58 = F_list_append_unique_oid(m, v42, v57)
	mBase = m.M
	v61 = m.ExcPending
	if v61 != 0 {
		goto L19
	} else {
		goto L20
	}
L17:
	;
	v63 = v42
	v64 = v44
	goto L18
L18:
	;
	v66 = v41 + int32(1)
	if v66 < v64 {
		v41 = v66
		v42 = v63
		v44 = v64
		goto L14
	} else {
		goto L21
	}
L19:
	;
	return int32(0)
L20:
	;
	v62 = *(*int32)(unsafe.Add(mBase, uint32(v32)))
	v63 = v58
	v64 = v62
	goto L18
L21:
	;
	goto L15
L22:
	;
	goto L10
L23:
	;
	v92 = int32(0)
	v93 = *(*int32)(unsafe.Add(mBase, uint32(v84)+4))
	if v93 <= v92 {
		goto L24
	} else {
		goto L25
	}
L24:
	;
	v137 = v84
	goto L2
L25:
	;
	goto L26
L26:
	;
	v98 = v92
	v100 = v84
	goto L27
L27:
	;
	v105 = *(*int32)(unsafe.Add(mBase, uint32(v84)+12))
	v109 = *(*int32)(unsafe.Add(mBase, uint32(v105+v98<<(uint(int32(2))%32))))
	v110 = F_getBaseType(m, v109)
	mBase = m.M
	v111 = m.ExcPending
	if v111 != 0 {
		goto L19
	} else {
		goto L29
	}
L28:
	;
	v137 = v115
	goto L2
L29:
	;
	if v109 != v110 {
		goto L30
	} else {
		goto L31
	}
L30:
	;
	v113 = F_list_append_unique_oid(m, v100, v110)
	mBase = m.M
	v114 = m.ExcPending
	if v114 != 0 {
		goto L19
	} else {
		goto L33
	}
L31:
	;
	v115 = v100
	goto L32
L32:
	;
	v117 = v98 + int32(1)
	v118 = *(*int32)(unsafe.Add(mBase, uint32(v84)+4))
	if v117 < v118 {
		v98 = v117
		v100 = v115
		goto L27
	} else {
		goto L34
	}
L33:
	;
	v115 = v113
	goto L32
L34:
	;
	goto L28
L35:
	;
	goto L1
L36:
	;
	if v137 == int32(0) {
		goto L1
	} else {
		goto L37
	}
L37:
	;
	v148 = int32(0)
	v149 = *(*int32)(unsafe.Add(mBase, uint32(v137)+4))
	if v149 <= v148 {
		goto L1
	} else {
		goto L38
	}
L38:
	;
	v157 = v148
	goto L39
L39:
	;
	v161 = *(*int32)(unsafe.Add(mBase, uint32(v137)+12))
	v165 = *(*int32)(unsafe.Add(mBase, uint32(v161+v157<<(uint(int32(2))%32))))
	v167 = F_map_sql_type_to_xml_name(m, v165, int32(-1))
	mBase = m.M
	v168 = m.ExcPending
	if v168 != 0 {
		goto L19
	} else {
		goto L41
	}
L40:
	;
	goto L1
L41:
	;
	F_initStringInfo(m, v12+int32(160))
	mBase = m.M
	v172 = m.ExcPending
	if v172 != 0 {
		goto L19
	} else {
		goto L42
	}
L42:
	;
	if v165 == int32(142) {
		v329 = int32(788153)
		goto L43
	} else {
		goto L44
	}
L43:
	;
	F_appendStringInfoString(m, v12+int32(160), v329)
	mBase = m.M
	v333 = m.ExcPending
	if v333 != 0 {
		goto L19
	} else {
		goto L97
	}
L44:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v12)+128)) = v167
	F_appendStringInfo(m, v12+int32(160), int32(788447), v12+int32(128))
	mBase = m.M
	v183 = m.ExcPending
	if v183 != 0 {
		goto L19
	} else {
		goto L45
	}
L45:
	;
	v184 = int32(788328)
	if v165 <= int32(1041) {
		goto L57
	} else {
		goto L58
	}
L46:
	;
	F_appendStringInfoString(m, v12+int32(160), int32(787295))
	mBase = m.M
	v328 = m.ExcPending
	if v328 != 0 {
		goto L19
	} else {
		goto L96
	}
L47:
	;
	if v165 == int32(1266) {
		goto L92
	} else {
		goto L93
	}
L48:
	;
	F_appendStringInfoString(m, v12+int32(160), int32(787180))
	mBase = m.M
	v310 = m.ExcPending
	if v310 != 0 {
		goto L19
	} else {
		goto L91
	}
L49:
	;
	F_appendStringInfoString(m, v12+int32(160), int32(787238))
	mBase = m.M
	v305 = m.ExcPending
	if v305 != 0 {
		goto L19
	} else {
		goto L90
	}
L50:
	;
	F_appendStringInfoString(m, v12+int32(160), int32(787124))
	mBase = m.M
	v300 = m.ExcPending
	if v300 != 0 {
		goto L19
	} else {
		goto L89
	}
L51:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v12)+88)) = int64(-9223372036854775807 - 1)
	*(*int64)(unsafe.Add(mBase, uint32(v12)+80)) = int64(9223372036854775807)
	F_appendStringInfo(m, v12+int32(160), int32(787700), v12+int32(80))
	mBase = m.M
	v295 = m.ExcPending
	if v295 != 0 {
		goto L19
	} else {
		goto L88
	}
L52:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v12)+64)) = int64(-9223372034707292161)
	F_appendStringInfo(m, v12+int32(160), int32(787961), v12-int32(-64))
	mBase = m.M
	v284 = m.ExcPending
	if v284 != 0 {
		goto L19
	} else {
		goto L87
	}
L53:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v12)+48)) = int64(-140737488322561)
	F_appendStringInfo(m, v12+int32(160), int32(787832), v12+int32(48))
	mBase = m.M
	v275 = m.ExcPending
	if v275 != 0 {
		goto L19
	} else {
		goto L86
	}
L54:
	;
	v257 = *(*int32)(unsafe.Add(mBase, _consts[1111]))
	if v257 != 0 {
		goto L82
	} else {
		goto L83
	}
L55:
	;
	F_appendStringInfoString(m, v12+int32(160), int32(788533))
	mBase = m.M
	v248 = m.ExcPending
	if v248 != 0 {
		goto L19
	} else {
		goto L80
	}
L56:
	;
	v223 = F_get_typtype(m, v165)
	mBase = m.M
	v224 = m.ExcPending
	if v224 != 0 {
		goto L19
	} else {
		goto L75
	}
L57:
	;
	switch v165 - int32(16) {
	case 0:
		goto L48
	case 1:
		goto L54
	case 2, 3, 6, 8:
		goto L56
	case 4:
		goto L51
	case 5:
		goto L53
	case 7:
		goto L52
	case 9:
		goto L55
	default:
		goto L60
	}
L58:
	;
	goto L59
L59:
	;
	if v165 <= int32(1113) {
		goto L61
	} else {
		goto L62
	}
L60:
	;
	switch v165 - int32(700) {
	case 0:
		goto L50
	case 1:
		goto L49
	default:
		goto L56
	}
L61:
	;
	if base.Ui32(v165-int32(1042)) < base.Ui32(int32(2)) {
		goto L55
	} else {
		goto L64
	}
L62:
	;
	goto L63
L63:
	;
	if v165 <= int32(1265) {
		goto L65
	} else {
		goto L66
	}
L64:
	;
	switch v165 - int32(1082) {
	case 0:
		goto L46
	case 1:
		goto L47
	default:
		goto L56
	}
L65:
	;
	if base.B2i32(v165 != int32(1114))&base.B2i32(v165 != int32(1184)) != 0 {
		goto L56
	} else {
		goto L68
	}
L66:
	;
	goto L67
L67:
	;
	if v165 == int32(1266) {
		goto L47
	} else {
		goto L73
	}
L68:
	;
	if v165 == int32(1184) {
		goto L69
	} else {
		goto L70
	}
L69:
	;
	v210 = int32(6811)
	goto L71
L70:
	;
	v210 = int32(791891)
	goto L71
L71:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v12)+112)) = v210
	F_appendStringInfo(m, v12+int32(160), int32(787410), v12+int32(112))
	mBase = m.M
	v218 = m.ExcPending
	if v218 != 0 {
		goto L19
	} else {
		goto L72
	}
L72:
	;
	v329 = v184
	goto L43
L73:
	;
	if v165 == int32(1700) {
		v329 = v184
		goto L43
	} else {
		goto L74
	}
L74:
	;
	goto L56
L75:
	;
	if v223 != int32(100) {
		v329 = v184
		goto L43
	} else {
		goto L76
	}
L76:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v12)+156)) = int32(-1)
	v231 = F_getBaseTypeAndTypmod(m, v165, v12+int32(156))
	mBase = m.M
	v232 = m.ExcPending
	if v232 != 0 {
		goto L19
	} else {
		goto L77
	}
L77:
	;
	v233 = *(*int32)(unsafe.Add(mBase, uint32(v12)+156))
	v234 = F_map_sql_type_to_xml_name(m, v231, v233)
	mBase = m.M
	v235 = m.ExcPending
	if v235 != 0 {
		goto L19
	} else {
		goto L78
	}
L78:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v12)+16)) = v234
	F_appendStringInfo(m, v12+int32(160), int32(788415), v12+int32(16))
	mBase = m.M
	v243 = m.ExcPending
	if v243 != 0 {
		goto L19
	} else {
		goto L79
	}
L79:
	;
	v329 = v184
	goto L43
L80:
	;
	F_appendStringInfoString(m, v12+int32(160), int32(788122))
	mBase = m.M
	v253 = m.ExcPending
	if v253 != 0 {
		goto L19
	} else {
		goto L81
	}
L81:
	;
	v329 = v184
	goto L43
L82:
	;
	v258 = int32(18039)
	goto L84
L83:
	;
	v258 = int32(18049)
	goto L84
L84:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v12)+32)) = v258
	F_appendStringInfo(m, v12+int32(160), int32(788088), v12+int32(32))
	mBase = m.M
	v266 = m.ExcPending
	if v266 != 0 {
		goto L19
	} else {
		goto L85
	}
L85:
	;
	v329 = v184
	goto L43
L86:
	;
	v329 = v184
	goto L43
L87:
	;
	v329 = v184
	goto L43
L88:
	;
	v329 = v184
	goto L43
L89:
	;
	v329 = v184
	goto L43
L90:
	;
	v329 = v184
	goto L43
L91:
	;
	v329 = v184
	goto L43
L92:
	;
	v315 = int32(6811)
	goto L94
L93:
	;
	v315 = int32(791891)
	goto L94
L94:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v12)+96)) = v315
	F_appendStringInfo(m, v12+int32(160), int32(787572), v12+int32(96))
	mBase = m.M
	v323 = m.ExcPending
	if v323 != 0 {
		goto L19
	} else {
		goto L95
	}
L95:
	;
	v329 = v184
	goto L43
L96:
	;
	v329 = v184
	goto L43
L97:
	;
	v334 = *(*int32)(unsafe.Add(mBase, uint32(v12)+160))
	*(*int32)(unsafe.Add(mBase, uint32(v12))) = v334
	F_appendStringInfo(m, v12+int32(140), int32(782775), v12)
	mBase = m.M
	v340 = m.ExcPending
	if v340 != 0 {
		goto L19
	} else {
		goto L98
	}
L98:
	;
	v342 = v157 + int32(1)
	v343 = *(*int32)(unsafe.Add(mBase, uint32(v137)+4))
	if v342 < v343 {
		v157 = v342
		goto L39
	} else {
		goto L99
	}
L99:
	;
	goto L40
}
func F_read_sql_construct(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32, l5 int32, l6 int32, l7 int32, l8 int32, l9 int32, l10 int32, l11 int32) int32 {
	mBase := m.M
	_ = mBase
	var v19 int32
	_ = v19
	var v21 int32
	_ = v21
	var v28 int32
	_ = v28
	var v29 int32
	_ = v29
	var v30 int32
	_ = v30
	var v34 int32
	_ = v34
	var v50 int32
	_ = v50
	var v51 int32
	_ = v51
	var v52 int32
	_ = v52
	var v54 int32
	_ = v54
	var v55 int32
	_ = v55
	var v58 int32
	_ = v58
	var v59 int32
	_ = v59
	var v77 int32
	_ = v77
	var v80 int32
	_ = v80
	var v84 int32
	_ = v84
	var v90 int32
	_ = v90
	var v93 int32
	_ = v93
	var v99 int32
	_ = v99
	var v100 int32
	_ = v100
	var v101 int32
	_ = v101
	var v102 int32
	_ = v102
	var v107 int32
	_ = v107
	var v108 int32
	_ = v108
	var v109 int32
	_ = v109
	var v110 int32
	_ = v110
	var v120 int32
	_ = v120
	var v124 int32
	_ = v124
	var v125 int32
	_ = v125
	var v127 int32
	_ = v127
	var v128 int32
	_ = v128
	var v129 int32
	_ = v129
	var v130 int32
	_ = v130
	var v134 int32
	_ = v134
	var v137 int32
	_ = v137
	var v138 int32
	_ = v138
	var v143 int32
	_ = v143
	var v145 int32
	_ = v145
	var v149 int32
	_ = v149
	var v152 int32
	_ = v152
	var v153 int32
	_ = v153
	var v154 int32
	_ = v154
	var v155 int32
	_ = v155
	var v158 int32
	_ = v158
	var v164 int32
	_ = v164
	var v165 int32
	_ = v165
	var v174 int32
	_ = v174
	var v175 int32
	_ = v175
	var v179 int32
	_ = v179
	var v192 int32
	_ = v192
	var v196 int32
	_ = v196
	var v197 int32
	_ = v197
	var v198 int32
	_ = v198
	var v199 int32
	_ = v199
	var v204 int32
	_ = v204
	var v208 int32
	_ = v208
	v19 = m.G0
	v21 = v19 + int32(-64)
	m.G0 = v21
	F_initStringInfo(m, v19+int32(-36))
	mBase = m.M
	v28 = m.ExcPending
	if v28 != 0 {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	return int32(0)
L2:
	;
	v29 = int32(4655356)
	v30 = *(*int32)(unsafe.Add(mBase, _consts[1345]))
	*(*int32)(unsafe.Add(mBase, _consts[1345])) = int32(2)
	v34 = int32(-1)
	v50 = int32(0)
	v51 = v34
	v52 = v34
	goto L6
L3:
	;
	F_plpgsql_yyerror(m, l10, int32(0), l11, int32(283067))
	mBase = m.M
	v208 = m.ExcPending
	if v208 != 0 {
		goto L1
	} else {
		goto L60
	}
L4:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v21))) = l3
	F_errmsg(m, int32(283409), v21)
	mBase = m.M
	v196 = m.ExcPending
	if v196 != 0 {
		goto L1
	} else {
		goto L57
	}
L5:
	;
	F_plpgsql_yyerror(m, l10, int32(0), l11, int32(171330))
	mBase = m.M
	v192 = m.ExcPending
	if v192 != 0 {
		goto L1
	} else {
		goto L56
	}
L6:
	;
	v54 = F_plpgsql_yylex(m, l9, l10, l11)
	mBase = m.M
	v55 = m.ExcPending
	if v55 != 0 {
		goto L1
	} else {
		goto L8
	}
L7:
	;
	*(*int32)(unsafe.Add(mBase, _consts[1345])) = v30
	if l7 != 0 {
		goto L37
	} else {
		goto L38
	}
L8:
	;
	if v51 < int32(0) {
		goto L9
	} else {
		goto L10
	}
L9:
	;
	v58 = *(*int32)(unsafe.Add(mBase, uint32(l10)))
	v59 = v58
	goto L11
L10:
	;
	v59 = v51
	goto L11
L11:
	;
	if base.B2i32(l0 != v54)&base.B2i32(l1 != v54)&base.B2i32(l2 != v54)|v50 != 0 {
		goto L12
	} else {
		goto L13
	}
L12:
	;
	switch v54 - int32(40) {
	case 0:
		goto L16
	case 1:
		goto L17
	default:
		goto L18
	}
L13:
	;
	goto L14
L14:
	;
	goto L7
L15:
	;
	if v54 != 0 {
		goto L23
	} else {
		goto L24
	}
L16:
	;
	v80 = v50 + int32(1)
	goto L15
L17:
	;
	if int32(0) < v50 {
		goto L19
	} else {
		goto L20
	}
L18:
	;
	switch v54 - int32(91) {
	case 0:
		goto L16
	default:
		v80 = v50
		goto L15
	case 2:
		goto L17
	}
L19:
	;
	v80 = v50 - int32(1)
	goto L15
L20:
	;
	goto L21
L21:
	;
	F_plpgsql_yyerror(m, l10, int32(0), l11, int32(171330))
	mBase = m.M
	v77 = m.ExcPending
	if v77 != 0 {
		goto L1
	} else {
		goto L22
	}
L22:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L23:
	;
	v84 = base.B2i32(v54 != int32(59))
	goto L25
L24:
	;
	v84 = int32(0)
	goto L25
L25:
	;
	if v84 == int32(0) {
		goto L26
	} else {
		goto L27
	}
L26:
	;
	if v80 != 0 {
		goto L5
	} else {
		goto L29
	}
L27:
	;
	goto L28
L28:
	;
	v108 = *(*int32)(unsafe.Add(mBase, uint32(l10)))
	v109 = *(*int32)(unsafe.Add(mBase, uint32(l11)))
	v110 = *(*int32)(unsafe.Add(mBase, uint32(v109)+64))
	goto L36
L29:
	;
	F_errstart_cold(m, int32(21), int32(583335))
	mBase = m.M
	v90 = m.ExcPending
	if v90 != 0 {
		goto L1
	} else {
		goto L30
	}
L30:
	;
	F_errcode(m, int32(16801924))
	mBase = m.M
	v93 = m.ExcPending
	if v93 != 0 {
		goto L1
	} else {
		goto L31
	}
L31:
	;
	if l5 != 0 {
		goto L4
	} else {
		goto L32
	}
L32:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v21)+16)) = l3
	F_errmsg(m, int32(102323), v19+int32(-48))
	mBase = m.M
	v99 = m.ExcPending
	if v99 != 0 {
		goto L1
	} else {
		goto L33
	}
L33:
	;
	v100 = *(*int32)(unsafe.Add(mBase, uint32(l10)))
	v101 = F_plpgsql_scanner_errposition(m, v100, l11)
	mBase = m.M
	v102 = m.ExcPending
	if v102 != 0 {
		goto L1
	} else {
		goto L34
	}
L34:
	;
	F_errfinish(m, int32(27545), int32(2822), int32(115946))
	mBase = m.M
	v107 = m.ExcPending
	if v107 != 0 {
		goto L1
	} else {
		goto L35
	}
L35:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L36:
	;
	v50 = v80
	v51 = v59
	v52 = v108 + v110
	goto L6
L37:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l7))) = v59
	goto L39
L38:
	;
	goto L39
L39:
	;
	if l8 != 0 {
		goto L40
	} else {
		goto L41
	}
L40:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l8))) = v54
	goto L42
L41:
	;
	goto L42
L42:
	;
	if v52 <= v59 {
		goto L43
	} else {
		goto L44
	}
L43:
	;
	if l5 != 0 {
		goto L3
	} else {
		goto L46
	}
L44:
	;
	goto L45
L45:
	;
	F_plpgsql_append_source_text(m, v19+int32(-36), v59, v52, l11)
	mBase = m.M
	v124 = m.ExcPending
	if v124 != 0 {
		goto L1
	} else {
		goto L48
	}
L46:
	;
	F_plpgsql_yyerror(m, l10, int32(0), l11, int32(102301))
	mBase = m.M
	v120 = m.ExcPending
	if v120 != 0 {
		goto L1
	} else {
		goto L47
	}
L47:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L48:
	;
	v125 = *(*int32)(unsafe.Add(mBase, uint32(v21)+28))
	v127 = F_palloc0(m, int32(80))
	mBase = m.M
	v128 = m.ExcPending
	if v128 != 0 {
		goto L1
	} else {
		goto L49
	}
L49:
	;
	v129 = F_pstrdup(m, v125)
	mBase = m.M
	v130 = m.ExcPending
	if v130 != 0 {
		goto L1
	} else {
		goto L50
	}
L50:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v127)+4)) = l4
	*(*int32)(unsafe.Add(mBase, uint32(v127))) = v129
	v134 = *(*int32)(unsafe.Add(mBase, _consts[1364]))
	*(*int32)(unsafe.Add(mBase, uint32(v127)+8)) = v134
	v137 = *(*int32)(unsafe.Add(mBase, _consts[1346]))
	v138 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v127)+20)) = uint8(v138)
	*(*int32)(unsafe.Add(mBase, uint32(v127)+16)) = int32(-1)
	*(*int32)(unsafe.Add(mBase, uint32(v127)+12)) = v137
	v143 = *(*int32)(unsafe.Add(mBase, uint32(v21)+28))
	F_pfree(m, v143)
	mBase = m.M
	v145 = m.ExcPending
	if v145 != 0 {
		goto L1
	} else {
		goto L51
	}
L51:
	;
	if l6 == int32(0) {
		goto L52
	} else {
		goto L53
	}
L52:
	;
	m.G0 = v21 - int32(-64)
	return v127
L53:
	;
	v149 = int32(*(*uint8)(unsafe.Add(mBase, _consts[1352])))
	if v149 != int32(1) {
		goto L52
	} else {
		goto L54
	}
L54:
	;
	v152 = *(*int32)(unsafe.Add(mBase, uint32(v127)+4))
	v153 = *(*int32)(unsafe.Add(mBase, uint32(v127)))
	v154 = int32(4562080)
	v155 = *(*int32)(unsafe.Add(mBase, _consts[10]))
	v158 = *(*int32)(unsafe.Add(mBase, _consts[1347]))
	*(*int32)(unsafe.Add(mBase, _consts[10])) = v158
	*(*int32)(unsafe.Add(mBase, uint32(v21)+60)) = l11
	*(*int32)(unsafe.Add(mBase, uint32(v21)+56)) = v59
	*(*int32)(unsafe.Add(mBase, uint32(v21)+48)) = int32(7342)
	v164 = int32(4554984)
	v165 = *(*int32)(unsafe.Add(mBase, _consts[140]))
	*(*int32)(unsafe.Add(mBase, _consts[140])) = v19 + int32(-20)
	*(*int32)(unsafe.Add(mBase, uint32(v21)+44)) = v165
	*(*int32)(unsafe.Add(mBase, uint32(v21)+52)) = v19 + int32(-8)
	v174 = F_raw_parser(m, v153, v152)
	mBase = m.M
	v175 = m.ExcPending
	if v175 != 0 {
		goto L1
	} else {
		goto L55
	}
L55:
	;
	*(*int32)(unsafe.Add(mBase, _consts[10])) = v155
	v179 = *(*int32)(unsafe.Add(mBase, uint32(v21)+44))
	*(*int32)(unsafe.Add(mBase, _consts[140])) = v179
	goto L52
L56:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L57:
	;
	v197 = *(*int32)(unsafe.Add(mBase, uint32(l10)))
	v198 = F_plpgsql_scanner_errposition(m, v197, l11)
	mBase = m.M
	v199 = m.ExcPending
	if v199 != 0 {
		goto L1
	} else {
		goto L58
	}
L58:
	;
	F_errfinish(m, int32(27545), int32(2816), int32(115946))
	mBase = m.M
	v204 = m.ExcPending
	if v204 != 0 {
		goto L1
	} else {
		goto L59
	}
L59:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L60:
	;
	base.Wasm_trap_unreachable()
	for {
	}
}
func F_sql_postrewrite_callback(m *base.Module, l0 int32, l1 int32) {
	mBase := m.M
	_ = mBase
	var v3 int32
	_ = v3
	var v9 int32
	_ = v9
	var v12 int32
	_ = v12
	var v15 int32
	_ = v15
	var v22 int32
	_ = v22
	var v23 int32
	_ = v23
	var v26 int32
	_ = v26
	var v27 int32
	_ = v27
	var v30 int32
	_ = v30
	var v33 int32
	_ = v33
	var v41 int32
	_ = v41
	var v42 int32
	_ = v42
	var v43 int32
	_ = v43
	var v45 int32
	_ = v45
	var v46 int32
	_ = v46
	var v47 int32
	_ = v47
	var v52 int32
	_ = v52
	var v55 int32
	_ = v55
	var v59 int32
	_ = v59
	var v64 int32
	_ = v64
	var v68 int32
	_ = v68
	var v71 int32
	_ = v71
	var v75 int32
	_ = v75
	var v80 int32
	_ = v80
	v3 = int32(0)
	if l0 == v3 {
		goto L3
	} else {
		goto L4
	}
L1:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v68 = m.ExcPending
	if v68 != 0 {
		goto L16
	} else {
		goto L23
	}
L2:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v52 = m.ExcPending
	if v52 != 0 {
		goto L16
	} else {
		goto L19
	}
L3:
	;
	if l1 != 0 {
		goto L13
	} else {
		goto L14
	}
L4:
	;
	v9 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	if v9 <= int32(0) {
		goto L3
	} else {
		goto L5
	}
L5:
	;
	v12 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	v15 = v3
	goto L6
L6:
	;
	v22 = *(*int32)(unsafe.Add(mBase, uint32(v12+v15<<(uint(int32(2))%32))))
	v23 = *(*int32)(unsafe.Add(mBase, uint32(v22)+4))
	if v23 != int32(6) {
		goto L8
	} else {
		goto L9
	}
L7:
	;
	goto L3
L8:
	;
	v33 = v15 + int32(1)
	if v9 != v33 {
		v15 = v33
		goto L6
	} else {
		goto L12
	}
L9:
	;
	v26 = *(*int32)(unsafe.Add(mBase, uint32(v22)+28))
	v27 = *(*int32)(unsafe.Add(mBase, uint32(v26)))
	if v27 != int32(213) {
		goto L8
	} else {
		goto L10
	}
L10:
	;
	v30 = *(*int32)(unsafe.Add(mBase, uint32(v26)+12))
	if v30 != 0 {
		goto L2
	} else {
		goto L11
	}
L11:
	;
	goto L8
L12:
	;
	goto L7
L13:
	;
	v41 = *(*int32)(unsafe.Add(mBase, uint32(l1)+48))
	v42 = *(*int32)(unsafe.Add(mBase, uint32(l1)+60))
	v43 = int32(*(*int8)(unsafe.Add(mBase, uint32(l1)+58)))
	v45 = F_check_sql_stmt_retval(m, l0, v41, v42, v43, int32(0))
	mBase = m.M
	v46 = m.ExcPending
	if v46 != 0 {
		goto L16
	} else {
		goto L17
	}
L14:
	;
	goto L15
L15:
	;
	return
L16:
	;
	return
L17:
	;
	v47 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l1)+56)))
	if v45 != v47 {
		goto L1
	} else {
		goto L18
	}
L18:
	;
	goto L15
L19:
	;
	F_errcode(m, int32(1088))
	mBase = m.M
	v55 = m.ExcPending
	if v55 != 0 {
		goto L16
	} else {
		goto L20
	}
L20:
	;
	F_errmsg(m, int32(150070), int32(0))
	mBase = m.M
	v59 = m.ExcPending
	if v59 != 0 {
		goto L16
	} else {
		goto L21
	}
L21:
	;
	F_errfinish(m, int32(517657), int32(2075), int32(101990))
	mBase = m.M
	v64 = m.ExcPending
	if v64 != 0 {
		goto L16
	} else {
		goto L22
	}
L22:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L23:
	;
	F_errcode(m, int32(1088))
	mBase = m.M
	v71 = m.ExcPending
	if v71 != 0 {
		goto L16
	} else {
		goto L24
	}
L24:
	;
	F_errmsg(m, int32(386322), int32(0))
	mBase = m.M
	v75 = m.ExcPending
	if v75 != 0 {
		goto L16
	} else {
		goto L25
	}
L25:
	;
	F_errfinish(m, int32(517657), int32(1270), int32(334511))
	mBase = m.M
	v80 = m.ExcPending
	if v80 != 0 {
		goto L16
	} else {
		goto L26
	}
L26:
	;
	base.Wasm_trap_unreachable()
	for {
	}
}
func F_unpack_sql_state(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v2 int32
	_ = v2
	var v3 int32
	_ = v3
	var v5 int32
	_ = v5
	var v6 int32
	_ = v6
	var v14 int32
	_ = v14
	var v22 int32
	_ = v22
	var v30 int32
	_ = v30
	var v38 int32
	_ = v38
	var v41 int32
	_ = v41
	v2 = int32(4555816)
	v3 = int32(63)
	v5 = int32(48)
	v6 = l0&v3 + v5
	*(*uint8)(unsafe.Add(mBase, _consts[1149])) = uint8(v6)
	v14 = int32(base.Ui32(l0)>>(uint(int32(24))%32))&v3 + v5
	*(*uint8)(unsafe.Add(mBase, _consts[1150])) = uint8(v14)
	v22 = int32(base.Ui32(l0)>>(uint(int32(18))%32))&v3 + v5
	*(*uint8)(unsafe.Add(mBase, _consts[1151])) = uint8(v22)
	v30 = int32(base.Ui32(l0)>>(uint(int32(12))%32))&v3 + v5
	*(*uint8)(unsafe.Add(mBase, _consts[1152])) = uint8(v30)
	v38 = int32(base.Ui32(l0)>>(uint(int32(6))%32))&v3 + v5
	*(*uint8)(unsafe.Add(mBase, _consts[1153])) = uint8(v38)
	v41 = int32(0)
	*(*uint8)(unsafe.Add(mBase, _consts[1154])) = uint8(v41)
	return v2
}
