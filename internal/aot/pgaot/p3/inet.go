package p3

import (
	base "github.com/shibukawa/pgmem/internal/aot/pgaot/base"
	"unsafe"
)

func F_inet_cidr_pton_ipv6(m *base.Module, l0 int32, l1 int32, l2 int32) int32 {
	mBase := m.M
	_ = mBase
	var v4 int32
	_ = v4
	var v17 int32
	_ = v17
	var v18 int32
	_ = v18
	var v19 int32
	_ = v19
	var v23 int64
	_ = v23
	var v27 int32
	_ = v27
	var v32 int32
	_ = v32
	var v38 int32
	_ = v38
	var v39 int32
	_ = v39
	var v41 int32
	_ = v41
	var v42 int32
	_ = v42
	var v44 int32
	_ = v44
	var v45 int32
	_ = v45
	var v46 int32
	_ = v46
	var v47 int32
	_ = v47
	var v48 int32
	_ = v48
	var v51 int32
	_ = v51
	var v53 int32
	_ = v53
	var v56 int32
	_ = v56
	var v59 int32
	_ = v59
	var v60 int32
	_ = v60
	var v62 int32
	_ = v62
	var v63 int32
	_ = v63
	var v102 int32
	_ = v102
	var v103 int32
	_ = v103
	var v112 int32
	_ = v112
	var v114 int32
	_ = v114
	var v118 int32
	_ = v118
	var v119 int32
	_ = v119
	var v122 int32
	_ = v122
	var v126 int32
	_ = v126
	var v127 int32
	_ = v127
	var v129 int32
	_ = v129
	var v132 int32
	_ = v132
	var v134 int32
	_ = v134
	var v139 int32
	_ = v139
	var v141 int32
	_ = v141
	var v146 int32
	_ = v146
	var v148 int32
	_ = v148
	var v151 int32
	_ = v151
	var v153 int32
	_ = v153
	var v156 int32
	_ = v156
	var v168 int32
	_ = v168
	var v171 int32
	_ = v171
	var v173 int32
	_ = v173
	var v212 int32
	_ = v212
	var v213 int32
	_ = v213
	var v222 int32
	_ = v222
	var v224 int32
	_ = v224
	var v228 int32
	_ = v228
	var v229 int32
	_ = v229
	var v232 int32
	_ = v232
	var v236 int32
	_ = v236
	var v237 int32
	_ = v237
	var v239 int32
	_ = v239
	var v242 int32
	_ = v242
	var v244 int32
	_ = v244
	var v249 int32
	_ = v249
	var v251 int32
	_ = v251
	var v256 int32
	_ = v256
	var v258 int32
	_ = v258
	var v261 int32
	_ = v261
	var v263 int32
	_ = v263
	var v266 int32
	_ = v266
	var v278 int32
	_ = v278
	var v281 int32
	_ = v281
	var v282 int32
	_ = v282
	var v283 int32
	_ = v283
	var v285 int32
	_ = v285
	var v293 int32
	_ = v293
	var v301 int32
	_ = v301
	var v304 int32
	_ = v304
	var v310 int32
	_ = v310
	var v316 int32
	_ = v316
	var v320 int32
	_ = v320
	var v325 int32
	_ = v325
	var v326 int32
	_ = v326
	var v327 int32
	_ = v327
	var v328 int32
	_ = v328
	var v330 int32
	_ = v330
	var v331 int32
	_ = v331
	var v333 int32
	_ = v333
	var v338 int32
	_ = v338
	var v343 int32
	_ = v343
	var v345 int32
	_ = v345
	var v355 int32
	_ = v355
	var v357 int32
	_ = v357
	var v359 int32
	_ = v359
	var v363 int32
	_ = v363
	var v373 int32
	_ = v373
	var v377 int32
	_ = v377
	var v379 int32
	_ = v379
	var v394 int32
	_ = v394
	var v396 int32
	_ = v396
	var v399 int32
	_ = v399
	var v401 int32
	_ = v401
	var v402 int32
	_ = v402
	var v403 int32
	_ = v403
	var v406 int32
	_ = v406
	var v419 int32
	_ = v419
	var v420 int32
	_ = v420
	var v429 int32
	_ = v429
	var v431 int32
	_ = v431
	var v435 int32
	_ = v435
	var v436 int32
	_ = v436
	var v439 int32
	_ = v439
	var v443 int32
	_ = v443
	var v444 int32
	_ = v444
	var v446 int32
	_ = v446
	var v449 int32
	_ = v449
	var v451 int32
	_ = v451
	var v456 int32
	_ = v456
	var v458 int32
	_ = v458
	var v463 int32
	_ = v463
	var v465 int32
	_ = v465
	var v468 int32
	_ = v468
	var v470 int32
	_ = v470
	var v473 int32
	_ = v473
	var v485 int32
	_ = v485
	var v487 int32
	_ = v487
	var v494 int32
	_ = v494
	var v511 int32
	_ = v511
	var v517 int32
	_ = v517
	var v518 int32
	_ = v518
	var v527 int32
	_ = v527
	var v530 int32
	_ = v530
	var v531 int32
	_ = v531
	var v570 int32
	_ = v570
	var v571 int32
	_ = v571
	var v580 int32
	_ = v580
	var v582 int32
	_ = v582
	var v586 int32
	_ = v586
	var v587 int32
	_ = v587
	var v590 int32
	_ = v590
	var v594 int32
	_ = v594
	var v595 int32
	_ = v595
	var v597 int32
	_ = v597
	var v600 int32
	_ = v600
	var v602 int32
	_ = v602
	var v607 int32
	_ = v607
	var v609 int32
	_ = v609
	var v614 int32
	_ = v614
	var v616 int32
	_ = v616
	var v619 int32
	_ = v619
	var v621 int32
	_ = v621
	var v624 int32
	_ = v624
	var v636 int32
	_ = v636
	var v637 int32
	_ = v637
	var v640 int32
	_ = v640
	var v642 int32
	_ = v642
	var v650 int32
	_ = v650
	var v661 int32
	_ = v661
	var v664 int32
	_ = v664
	var v668 int32
	_ = v668
	var v677 int32
	_ = v677
	var v678 int32
	_ = v678
	var v679 int32
	_ = v679
	var v718 int32
	_ = v718
	var v719 int32
	_ = v719
	var v728 int32
	_ = v728
	var v730 int32
	_ = v730
	var v734 int32
	_ = v734
	var v735 int32
	_ = v735
	var v738 int32
	_ = v738
	var v742 int32
	_ = v742
	var v743 int32
	_ = v743
	var v745 int32
	_ = v745
	var v748 int32
	_ = v748
	var v750 int32
	_ = v750
	var v755 int32
	_ = v755
	var v757 int32
	_ = v757
	var v762 int32
	_ = v762
	var v764 int32
	_ = v764
	var v767 int32
	_ = v767
	var v769 int32
	_ = v769
	var v772 int32
	_ = v772
	var v784 int32
	_ = v784
	var v785 int32
	_ = v785
	var v788 int32
	_ = v788
	var v790 int32
	_ = v790
	var v798 int32
	_ = v798
	var v807 int32
	_ = v807
	var v808 int32
	_ = v808
	var v810 int32
	_ = v810
	var v813 int32
	_ = v813
	var v814 int32
	_ = v814
	var v824 int32
	_ = v824
	var v830 int32
	_ = v830
	var v842 int32
	_ = v842
	var v844 int32
	_ = v844
	var v847 int32
	_ = v847
	var v856 int32
	_ = v856
	var v862 int32
	_ = v862
	var v867 int32
	_ = v867
	var v873 int32
	_ = v873
	var v876 int32
	_ = v876
	var v877 int32
	_ = v877
	var v878 int32
	_ = v878
	var v892 int32
	_ = v892
	var v893 int32
	_ = v893
	var v908 int32
	_ = v908
	var v909 int32
	_ = v909
	var v911 int32
	_ = v911
	var v914 int32
	_ = v914
	var v916 int32
	_ = v916
	var v917 int32
	_ = v917
	var v921 int32
	_ = v921
	var v922 int32
	_ = v922
	var v924 int32
	_ = v924
	var v930 int32
	_ = v930
	var v946 int32
	_ = v946
	var v947 int32
	_ = v947
	var v949 int32
	_ = v949
	var v1005 int64
	_ = v1005
	var v1007 int64
	_ = v1007
	var v1009 int32
	_ = v1009
	v4 = int32(0)
	v17 = m.G0
	v18 = int32(16)
	v19 = v17 - v18
	m.G0 = v19
	if base.Ui32(v18) <= base.Ui32(l2) {
		goto L7
	} else {
		goto L8
	}
L1:
	;
	m.G0 = v19 + int32(16)
	return v1009
L2:
	;
	v1005 = *(*int64)(unsafe.Add(mBase, uint32(v19)+8))
	*(*int64)(unsafe.Add(mBase, uint32(l1)+8)) = v1005
	v1007 = *(*int64)(unsafe.Add(mBase, uint32(v19)))
	*(*int64)(unsafe.Add(mBase, uint32(l1))) = v1007
	v1009 = v873
	goto L1
L3:
	;
	v1009 = int32(-1)
	goto L1
L4:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_inet_cidr_pton_ipv6[0])) = int32(44)
	goto L3
L5:
	;
	if v867 != 0 {
		goto L205
	} else {
		goto L206
	}
L6:
	;
	if v844 == int32(-1) {
		goto L202
	} else {
		goto L203
	}
L7:
	;
	v23 = int64(0)
	*(*int64)(unsafe.Add(mBase, uint32(v19)+8)) = v23
	*(*int64)(unsafe.Add(mBase, uint32(v19))) = v23
	v27 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0))))
	if v27 == int32(0) {
		goto L4
	} else {
		goto L10
	}
L8:
	;
	goto L9
L9:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_inet_cidr_pton_ipv6[0])) = int32(35)
	goto L3
L10:
	;
	if v27 == int32(58) {
		goto L11
	} else {
		goto L12
	}
L11:
	;
	v32 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+1)))
	if v32 != int32(58) {
		goto L4
	} else {
		goto L14
	}
L12:
	;
	v38 = l0
	v39 = v27
	goto L13
L13:
	;
	v41 = v19 + int32(16)
	v42 = v38
	v44 = v39
	v45 = v4
	v46 = v4
	v47 = v4
	v48 = v38
	v51 = v19
	v53 = v4
	v56 = v4
	goto L19
L14:
	;
	v38 = l0 + int32(1)
	v39 = int32(58)
	goto L13
L15:
	;
	if v511 != 0 {
		v842 = v338
		v844 = v518
		v847 = v45
		goto L6
	} else {
		goto L201
	}
L16:
	;
	if v814 == int32(0) {
		v842 = v808
		v844 = v810
		v847 = v813
		goto L6
	} else {
		goto L199
	}
L17:
	;
	v661 = v59
	v664 = int32(0)
	v668 = v4
	goto L163
L18:
	;
	if int32(12) < v47 {
		goto L4
	} else {
		goto L87
	}
L19:
	;
	v59 = v42 + int32(1)
	v60 = int32(_a_F_inet_cidr_pton_ipv6_0)
	v62 = base.I32_extend8_s(v44)
	v63 = int32(17)
	goto L26
L20:
	;
	v807 = v325
	v808 = v326
	v810 = int32(-1)
	v813 = v331
	v814 = v328
	goto L16
L21:
	;
	v333 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v59))))
	if v333 != 0 {
		v42 = v59
		v44 = v333
		v45 = v331
		v46 = v325
		v47 = v326
		v48 = v327
		v51 = v326 + v19
		v53 = v328
		v56 = v330
		goto L19
	} else {
		goto L86
	}
L22:
	;
	v293 = v44 & int32(255)
	if v293 != int32(58) {
		goto L78
	} else {
		goto L79
	}
L23:
	;
	if v168 == int32(0) {
		goto L48
	} else {
		goto L49
	}
L24:
	;
	v168 = int32(0)
	goto L23
L25:
	;
	v146 = v139
	v148 = v141
	goto L42
L26:
	;
	goto L33
L33:
	;
	v102 = v62 & int32(255)
	v103 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_inet_cidr_pton_ipv6[1])))
	if base.B2i32(v102 == v103)|int32(0) == int32(0) {
		goto L34
	} else {
		goto L35
	}
L34:
	;
	v112 = v60
	v114 = v63
	goto L37
L35:
	;
	v132 = v60
	v134 = v63
	goto L36
L36:
	;
	if v134 == int32(0) {
		goto L24
	} else {
		goto L41
	}
L37:
	;
	v118 = *(*int32)(unsafe.Add(mBase, uint32(v112)))
	v119 = v118 ^ v102*int32(16843009)
	v122 = int32(-2139062144)
	if (int32(16843008)-v119|v119)&v122 != v122 {
		v139 = v112
		v141 = v114
		goto L25
	} else {
		goto L39
	}
L38:
	;
	v132 = v127
	v134 = v129
	goto L36
L39:
	;
	v126 = int32(4)
	v127 = v112 + v126
	v129 = v114 - v126
	if base.Ui32(int32(3)) < base.Ui32(v129) {
		v112 = v127
		v114 = v129
		goto L37
	} else {
		goto L40
	}
L40:
	;
	goto L38
L41:
	;
	v139 = v132
	v141 = v134
	goto L25
L42:
	;
	v151 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v146))))
	if v62&int32(255) == v151 {
		goto L44
	} else {
		goto L45
	}
L43:
	;
	goto L24
L44:
	;
	v168 = v146
	goto L23
L45:
	;
	goto L46
L46:
	;
	v153 = int32(1)
	v156 = v148 - v153
	if v156 != 0 {
		v146 = v146 + v153
		v148 = v156
		goto L42
	} else {
		goto L47
	}
L47:
	;
	goto L43
L48:
	;
	v171 = int32(_a_F_inet_cidr_pton_ipv6_1)
	v173 = int32(17)
	goto L54
L49:
	;
	v281 = v60
	v282 = v168
	goto L50
L50:
	;
	v283 = int32(1)
	v285 = v56 + v283
	if int32(4) < v285 {
		goto L4
	} else {
		goto L77
	}
L51:
	;
	if v278 == int32(0) {
		goto L22
	} else {
		goto L76
	}
L52:
	;
	v278 = int32(0)
	goto L51
L53:
	;
	v256 = v249
	v258 = v251
	goto L70
L54:
	;
	goto L61
L61:
	;
	v212 = v62 & int32(255)
	v213 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_inet_cidr_pton_ipv6[2])))
	if base.B2i32(v212 == v213)|int32(0) == int32(0) {
		goto L62
	} else {
		goto L63
	}
L62:
	;
	v222 = v171
	v224 = v173
	goto L65
L63:
	;
	v242 = v171
	v244 = v173
	goto L64
L64:
	;
	if v244 == int32(0) {
		goto L52
	} else {
		goto L69
	}
L65:
	;
	v228 = *(*int32)(unsafe.Add(mBase, uint32(v222)))
	v229 = v228 ^ v212*int32(16843009)
	v232 = int32(-2139062144)
	if (int32(16843008)-v229|v229)&v232 != v232 {
		v249 = v222
		v251 = v224
		goto L53
	} else {
		goto L67
	}
L66:
	;
	v242 = v237
	v244 = v239
	goto L64
L67:
	;
	v236 = int32(4)
	v237 = v222 + v236
	v239 = v224 - v236
	if base.Ui32(int32(3)) < base.Ui32(v239) {
		v222 = v237
		v224 = v239
		goto L65
	} else {
		goto L68
	}
L68:
	;
	goto L66
L69:
	;
	v249 = v242
	v251 = v244
	goto L53
L70:
	;
	v261 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v256))))
	if v62&int32(255) == v261 {
		goto L72
	} else {
		goto L73
	}
L71:
	;
	goto L52
L72:
	;
	v278 = v256
	goto L51
L73:
	;
	goto L74
L74:
	;
	v263 = int32(1)
	v266 = v258 - v263
	if v266 != 0 {
		v256 = v256 + v263
		v258 = v266
		goto L70
	} else {
		goto L75
	}
L75:
	;
	goto L71
L76:
	;
	v281 = v171
	v282 = v278
	goto L50
L77:
	;
	v325 = v282 - v281 | v46<<(uint(int32(4))%32)
	v326 = v47
	v327 = v48
	v328 = v283
	v330 = v285
	v331 = v45
	goto L21
L78:
	;
	switch v293 - int32(46) {
	case 0:
		goto L18
	case 1:
		goto L17
	default:
		goto L4
	}
L79:
	;
	goto L80
L80:
	;
	if v53 == int32(0) {
		goto L81
	} else {
		goto L82
	}
L81:
	;
	v301 = int32(0)
	if v45 == v301 {
		v325 = v46
		v326 = v47
		v327 = v59
		v328 = v301
		v330 = v56
		v331 = v51
		goto L21
	} else {
		goto L84
	}
L82:
	;
	goto L83
L83:
	;
	v304 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v59))))
	if base.B2i32(v304 == int32(0))|base.B2i32(base.Ui32(int32(14)) < base.Ui32(v47)) != 0 {
		goto L4
	} else {
		goto L85
	}
L84:
	;
	goto L4
L85:
	;
	v310 = int32(8)
	v316 = v46<<(uint(v310)%32) | int32(base.Ui32(v46&int32(_a_F_inet_cidr_pton_ipv6_2))>>(uint(v310)%32))
	*(*uint16)(unsafe.Add(mBase, uint32(v51))) = uint16(v316)
	v320 = int32(0)
	v325 = v320
	v326 = v47 + int32(2)
	v327 = v59
	v328 = v320
	v330 = v320
	v331 = v45
	goto L21
L86:
	;
	goto L20
L87:
	;
	v338 = v47 + int32(4)
	v343 = v51
	v345 = v48
	goto L89
L88:
	;
	if base.B2i32(v357 == int32(0))|base.B2i32(int32(3) < v343-v51) != 0 {
		goto L4
	} else {
		goto L162
	}
L89:
	;
	v355 = int32(0)
	v357 = v355
	v359 = v355
	v363 = v345
	goto L91
L90:
	;
	v511 = int32(0)
	v517 = v377
	v518 = v4
	goto L129
L91:
	;
	v373 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v363))))
	if v373 == int32(0) {
		goto L88
	} else {
		goto L93
	}
L92:
	;
	if base.B2i32(v373&int32(254) != int32(46))|base.B2i32(int32(3) < v343-v51) != 0 {
		goto L4
	} else {
		goto L127
	}
L93:
	;
	v377 = v363 + int32(1)
	v379 = base.I32_extend8_s(v373)
	goto L98
L94:
	;
	if v485 != 0 {
		goto L119
	} else {
		goto L120
	}
L95:
	;
	v485 = int32(0)
	goto L94
L96:
	;
	v463 = v456
	v465 = v458
	goto L113
L97:
	;
	if base.B2i32(v402 != v403) == int32(0) {
		goto L95
	} else {
		goto L104
	}
L98:
	;
	v394 = int32(_a_F_inet_cidr_pton_ipv6_3)
	v396 = int32(11)
	goto L99
L99:
	;
	v399 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v394))))
	if v399 == v379&int32(255) {
		v456 = v394
		v458 = v396
		goto L96
	} else {
		goto L101
	}
L100:
	;
	goto L97
L101:
	;
	v401 = int32(1)
	v402 = v396 - v401
	v403 = int32(0)
	v406 = v394 + v401
	if v406&int32(3) == v403 {
		goto L97
	} else {
		goto L102
	}
L102:
	;
	if v402 != 0 {
		v394 = v406
		v396 = v402
		goto L99
	} else {
		goto L103
	}
L103:
	;
	goto L100
L104:
	;
	v419 = v379 & int32(255)
	v420 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v406))))
	if base.B2i32(v419 == v420)|base.B2i32(base.Ui32(v402) < base.Ui32(int32(4))) == int32(0) {
		goto L105
	} else {
		goto L106
	}
L105:
	;
	v429 = v406
	v431 = v402
	goto L108
L106:
	;
	v449 = v406
	v451 = v402
	goto L107
L107:
	;
	if v451 == int32(0) {
		goto L95
	} else {
		goto L112
	}
L108:
	;
	v435 = *(*int32)(unsafe.Add(mBase, uint32(v429)))
	v436 = v435 ^ v419*int32(16843009)
	v439 = int32(-2139062144)
	if (int32(16843008)-v436|v436)&v439 != v439 {
		v456 = v429
		v458 = v431
		goto L96
	} else {
		goto L110
	}
L109:
	;
	v449 = v444
	v451 = v446
	goto L107
L110:
	;
	v443 = int32(4)
	v444 = v429 + v443
	v446 = v431 - v443
	if base.Ui32(int32(3)) < base.Ui32(v446) {
		v429 = v444
		v431 = v446
		goto L108
	} else {
		goto L111
	}
L111:
	;
	goto L109
L112:
	;
	v456 = v449
	v458 = v451
	goto L96
L113:
	;
	v468 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v463))))
	if v379&int32(255) == v468 {
		goto L115
	} else {
		goto L116
	}
L114:
	;
	goto L95
L115:
	;
	v485 = v463
	goto L94
L116:
	;
	goto L117
L117:
	;
	v470 = int32(1)
	v473 = v465 - v470
	if v473 != 0 {
		v463 = v463 + v470
		v465 = v473
		goto L113
	} else {
		goto L118
	}
L118:
	;
	goto L114
L119:
	;
	if v359 != 0 {
		goto L122
	} else {
		goto L123
	}
L120:
	;
	goto L121
L121:
	;
	goto L92
L122:
	;
	v487 = int32(0)
	goto L124
L123:
	;
	v487 = v357
	goto L124
L124:
	;
	if v487 != 0 {
		goto L4
	} else {
		goto L125
	}
L125:
	;
	v494 = v485 - int32(_a_F_inet_cidr_pton_ipv6_3) + v359*int32(10)
	if base.Ui32(v494) < base.Ui32(int32(256)) {
		v357 = v357 + int32(1)
		v359 = v494
		v363 = v377
		goto L91
	} else {
		goto L126
	}
L126:
	;
	goto L4
L127:
	;
	*(*uint8)(unsafe.Add(mBase, uint32(v343))) = uint8(v359)
	if v373 != int32(47) {
		v343 = v343 + int32(1)
		v345 = v377
		goto L89
	} else {
		goto L128
	}
L128:
	;
	goto L90
L129:
	;
	v527 = int32(*(*int8)(unsafe.Add(mBase, uint32(v517))))
	if v527 == int32(0) {
		goto L15
	} else {
		goto L131
	}
L130:
	;
	goto L4
L131:
	;
	v530 = int32(_a_F_inet_cidr_pton_ipv6_4)
	v531 = int32(11)
	goto L135
L132:
	;
	v637 = int32(0)
	if v518 != 0 {
		goto L157
	} else {
		goto L158
	}
L133:
	;
	v636 = int32(0)
	goto L132
L134:
	;
	v614 = v607
	v616 = v609
	goto L151
L135:
	;
	goto L142
L142:
	;
	v570 = v527 & int32(255)
	v571 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_inet_cidr_pton_ipv6[3])))
	if base.B2i32(v570 == v571)|int32(0) == int32(0) {
		goto L143
	} else {
		goto L144
	}
L143:
	;
	v580 = v530
	v582 = v531
	goto L146
L144:
	;
	v600 = v530
	v602 = v531
	goto L145
L145:
	;
	if v602 == int32(0) {
		goto L133
	} else {
		goto L150
	}
L146:
	;
	v586 = *(*int32)(unsafe.Add(mBase, uint32(v580)))
	v587 = v586 ^ v570*int32(16843009)
	v590 = int32(-2139062144)
	if (int32(16843008)-v587|v587)&v590 != v590 {
		v607 = v580
		v609 = v582
		goto L134
	} else {
		goto L148
	}
L147:
	;
	v600 = v595
	v602 = v597
	goto L145
L148:
	;
	v594 = int32(4)
	v595 = v580 + v594
	v597 = v582 - v594
	if base.Ui32(int32(3)) < base.Ui32(v597) {
		v580 = v595
		v582 = v597
		goto L146
	} else {
		goto L149
	}
L149:
	;
	goto L147
L150:
	;
	v607 = v600
	v609 = v602
	goto L134
L151:
	;
	v619 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v614))))
	if v527&int32(255) == v619 {
		goto L153
	} else {
		goto L154
	}
L152:
	;
	goto L133
L153:
	;
	v636 = v614
	goto L132
L154:
	;
	goto L155
L155:
	;
	v621 = int32(1)
	v624 = v616 - v621
	if v624 != 0 {
		v614 = v614 + v621
		v616 = v624
		goto L151
	} else {
		goto L156
	}
L156:
	;
	goto L152
L157:
	;
	v640 = v637
	goto L159
L158:
	;
	v640 = v511
	goto L159
L159:
	;
	if base.B2i32(v636 == v637)|v640 != 0 {
		goto L4
	} else {
		goto L160
	}
L160:
	;
	v642 = int32(1)
	v650 = v636 - int32(_a_F_inet_cidr_pton_ipv6_4) + v518*int32(10)
	if v650 < int32(129) {
		v511 = v511 + v642
		v517 = v517 + v642
		v518 = v650
		goto L129
	} else {
		goto L161
	}
L161:
	;
	goto L130
L162:
	;
	*(*uint8)(unsafe.Add(mBase, uint32(v343))) = uint8(v359)
	v862 = v338
	v867 = v45
	v873 = int32(128)
	goto L5
L163:
	;
	v677 = int32(*(*int8)(unsafe.Add(mBase, uint32(v661))))
	if v677 != 0 {
		goto L165
	} else {
		goto L166
	}
L164:
	;
	if v664 == int32(0) {
		goto L4
	} else {
		goto L198
	}
L165:
	;
	v678 = int32(_a_F_inet_cidr_pton_ipv6_4)
	v679 = int32(11)
	goto L171
L166:
	;
	goto L167
L167:
	;
	goto L164
L168:
	;
	v785 = int32(0)
	if v668 != 0 {
		goto L193
	} else {
		goto L194
	}
L169:
	;
	v784 = int32(0)
	goto L168
L170:
	;
	v762 = v755
	v764 = v757
	goto L187
L171:
	;
	goto L178
L178:
	;
	v718 = v677 & int32(255)
	v719 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_inet_cidr_pton_ipv6[3])))
	if base.B2i32(v718 == v719)|int32(0) == int32(0) {
		goto L179
	} else {
		goto L180
	}
L179:
	;
	v728 = v678
	v730 = v679
	goto L182
L180:
	;
	v748 = v678
	v750 = v679
	goto L181
L181:
	;
	if v750 == int32(0) {
		goto L169
	} else {
		goto L186
	}
L182:
	;
	v734 = *(*int32)(unsafe.Add(mBase, uint32(v728)))
	v735 = v734 ^ v718*int32(16843009)
	v738 = int32(-2139062144)
	if (int32(16843008)-v735|v735)&v738 != v738 {
		v755 = v728
		v757 = v730
		goto L170
	} else {
		goto L184
	}
L183:
	;
	v748 = v743
	v750 = v745
	goto L181
L184:
	;
	v742 = int32(4)
	v743 = v728 + v742
	v745 = v730 - v742
	if base.Ui32(int32(3)) < base.Ui32(v745) {
		v728 = v743
		v730 = v745
		goto L182
	} else {
		goto L185
	}
L185:
	;
	goto L183
L186:
	;
	v755 = v748
	v757 = v750
	goto L170
L187:
	;
	v767 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v762))))
	if v677&int32(255) == v767 {
		goto L189
	} else {
		goto L190
	}
L188:
	;
	goto L169
L189:
	;
	v784 = v762
	goto L168
L190:
	;
	goto L191
L191:
	;
	v769 = int32(1)
	v772 = v764 - v769
	if v772 != 0 {
		v762 = v762 + v769
		v764 = v772
		goto L187
	} else {
		goto L192
	}
L192:
	;
	goto L188
L193:
	;
	v788 = v785
	goto L195
L194:
	;
	v788 = v664
	goto L195
L195:
	;
	if base.B2i32(v784 == v785)|v788 != 0 {
		goto L4
	} else {
		goto L196
	}
L196:
	;
	v790 = int32(1)
	v798 = v784 - int32(_a_F_inet_cidr_pton_ipv6_4) + v668*int32(10)
	if v798 < int32(129) {
		v661 = v661 + v790
		v664 = v664 + v790
		v668 = v798
		goto L163
	} else {
		goto L197
	}
L197:
	;
	goto L4
L198:
	;
	v807 = v46
	v808 = v47
	v810 = v668
	v813 = v45
	v814 = v53
	goto L16
L199:
	;
	if int32(14) < v808 {
		goto L4
	} else {
		goto L200
	}
L200:
	;
	v824 = int32(8)
	v830 = v807<<(uint(v824)%32) | int32(base.Ui32(v807&int32(_a_F_inet_cidr_pton_ipv6_2))>>(uint(v824)%32))
	*(*uint16)(unsafe.Add(mBase, uint32(v808+v19))) = uint16(v830)
	v842 = v808 + int32(2)
	v844 = v810
	v847 = v813
	goto L6
L201:
	;
	goto L4
L202:
	;
	v856 = int32(128)
	goto L204
L203:
	;
	v856 = v844
	goto L204
L204:
	;
	v862 = v842
	v867 = v847
	v873 = v856
	goto L5
L205:
	;
	if v862 == int32(16) {
		goto L4
	} else {
		goto L208
	}
L206:
	;
	goto L207
L207:
	;
	if v862 == int32(16) {
		goto L2
	} else {
		goto L217
	}
L208:
	;
	v876 = int32(1)
	v877 = v862 + v19
	v878 = v877 - v867
	if v878 <= int32(0) {
		goto L2
	} else {
		goto L209
	}
L209:
	;
	if v867+int32(1) != v877 {
		goto L210
	} else {
		goto L211
	}
L210:
	;
	v892 = v876
	v893 = int32(0)
	goto L213
L211:
	;
	v930 = v876
	goto L212
L212:
	;
	v946 = v867 + (v878 - v930)
	v947 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v946))))
	*(*uint8)(unsafe.Add(mBase, uint32(v41-v930))) = uint8(v947)
	v949 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v946))) = uint8(v949)
	goto L2
L213:
	;
	v908 = v867 + (v878 - v892)
	v909 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v908))))
	*(*uint8)(unsafe.Add(mBase, uint32(v41-v892))) = uint8(v909)
	v911 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v908))) = uint8(v911)
	v914 = v892 ^ int32(-1)
	v916 = v878 + v867 + v914
	v917 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v916))))
	*(*uint8)(unsafe.Add(mBase, uint32(v41+v914))) = uint8(v917)
	*(*uint8)(unsafe.Add(mBase, uint32(v916))) = uint8(v911)
	v921 = int32(2)
	v922 = v892 + v921
	v924 = v893 + v921
	if v924 != v878&int32(2147483646) {
		v892 = v922
		v893 = v924
		goto L213
	} else {
		goto L215
	}
L214:
	;
	if v878&int32(1) == int32(0) {
		goto L2
	} else {
		goto L216
	}
L215:
	;
	goto L214
L216:
	;
	v930 = v922
	goto L212
L217:
	;
	goto L4
}
func F_inet_gist_compress(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v9 int32
	_ = v9
	var v10 int32
	_ = v10
	var v15 int32
	_ = v15
	var v18 int32
	_ = v18
	var v19 int32
	_ = v19
	var v20 int32
	_ = v20
	var v21 int32
	_ = v21
	var v23 int32
	_ = v23
	var v24 int32
	_ = v24
	var v25 int32
	_ = v25
	var v27 int32
	_ = v27
	var v30 int32
	_ = v30
	var v32 int32
	_ = v32
	var v34 int32
	_ = v34
	var v35 int32
	_ = v35
	var v37 int32
	_ = v37
	var v38 int32
	_ = v38
	var v41 int32
	_ = v41
	var v42 int32
	_ = v42
	var v46 int32
	_ = v46
	var v47 int32
	_ = v47
	var v52 int32
	_ = v52
	var v55 int32
	_ = v55
	var v58 int32
	_ = v58
	var v64 int32
	_ = v64
	var v67 int32
	_ = v67
	var v73 int32
	_ = v73
	var v75 int32
	_ = v75
	var v77 int32
	_ = v77
	var v78 int32
	_ = v78
	v9 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v10 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v9)+14)))
	if v10 != int32(1) {
		return v9
	} else {
		v15 = F_palloc(m, int32(16))
		mBase = m.M
		v18 = m.ExcPending
		if v18 != 0 {
			return int32(0)
		} else {
			v19 = *(*int32)(unsafe.Add(mBase, uint32(v9)))
			if v19 != 0 {
				v20 = F_pg_detoast_datum_packed(m, v19)
				mBase = m.M
				v21 = m.ExcPending
				if v21 != 0 {
					return int32(0)
				} else {
					v23 = F_palloc0(m, int32(20))
					mBase = m.M
					v24 = m.ExcPending
					if v24 != 0 {
						return int32(0)
					} else {
						v25 = int32(1)
						v27 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v20))))
						if v27&v25 != 0 {
							v30 = v25
						} else {
							v30 = int32(4)
						}
						v32 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v20+v30))))
						*(*uint8)(unsafe.Add(mBase, uint32(v23)+1)) = uint8(v32)
						v34 = int32(1)
						v35 = v20 + v34
						v37 = v20 + int32(4)
						v38 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v20))))
						if v38&v34 != 0 {
							v41 = v35
						} else {
							v41 = v37
						}
						v42 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v41)+1)))
						v46 = base.B2i32(v32 == int32(3))
						if v32 == int32(3) {
							v47 = int32(-128)
						} else {
							v47 = int32(32)
						}
						*(*uint8)(unsafe.Add(mBase, uint32(v23)+3)) = uint8(v47)
						*(*uint8)(unsafe.Add(mBase, uint32(v23)+2)) = uint8(v42)
						if v32 == int32(3) {
							v52 = int32(16)
						} else {
							v52 = int32(4)
						}
						if v52 != 0 {
							v55 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v20))))
							if v55&int32(1) != 0 {
								v58 = v35
							} else {
								v58 = v37
							}
							base.MemoryCopy(m, v23+int32(4), v58+int32(2), v52)
						} else {
						}
						if v32 == int32(3) {
							v64 = int32(41)
						} else {
							v64 = int32(17)
						}
						*(*uint8)(unsafe.Add(mBase, uint32(v23))) = uint8(v64)
						v67 = v23
						*(*int32)(unsafe.Add(mBase, uint32(v15))) = v67
						v73 = *(*int32)(unsafe.Add(mBase, uint32(v9)+4))
						*(*int32)(unsafe.Add(mBase, uint32(v15)+4)) = v73
						v75 = *(*int32)(unsafe.Add(mBase, uint32(v9)+8))
						*(*int32)(unsafe.Add(mBase, uint32(v15)+8)) = v75
						v77 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v9)+12)))
						v78 = int32(0)
						*(*uint8)(unsafe.Add(mBase, uint32(v15)+14)) = uint8(v78)
						*(*uint16)(unsafe.Add(mBase, uint32(v15)+12)) = uint16(v77)
						return v15
					}
				}
			} else {
				v67 = int32(0)
				*(*int32)(unsafe.Add(mBase, uint32(v15))) = v67
				v73 = *(*int32)(unsafe.Add(mBase, uint32(v9)+4))
				*(*int32)(unsafe.Add(mBase, uint32(v15)+4)) = v73
				v75 = *(*int32)(unsafe.Add(mBase, uint32(v9)+8))
				*(*int32)(unsafe.Add(mBase, uint32(v15)+8)) = v75
				v77 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v9)+12)))
				v78 = int32(0)
				*(*uint8)(unsafe.Add(mBase, uint32(v15)+14)) = uint8(v78)
				*(*uint16)(unsafe.Add(mBase, uint32(v15)+12)) = uint16(v77)
				return v15
			}
		}
	}
}
func F_inet_merge(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v2 int32
	_ = v2
	var v10 int32
	_ = v10
	var v11 int32
	_ = v11
	var v14 int32
	_ = v14
	var v15 int32
	_ = v15
	var v16 int32
	_ = v16
	var v17 int32
	_ = v17
	var v18 int32
	_ = v18
	var v20 int32
	_ = v20
	var v23 int32
	_ = v23
	var v24 int32
	_ = v24
	var v25 int32
	_ = v25
	var v26 int32
	_ = v26
	var v28 int32
	_ = v28
	var v31 int32
	_ = v31
	var v32 int32
	_ = v32
	var v33 int32
	_ = v33
	var v35 int32
	_ = v35
	var v36 int32
	_ = v36
	var v38 int32
	_ = v38
	var v40 int32
	_ = v40
	var v41 int32
	_ = v41
	var v43 int32
	_ = v43
	var v47 int32
	_ = v47
	var v50 int32
	_ = v50
	var v58 int32
	_ = v58
	var v60 int32
	_ = v60
	var v64 int32
	_ = v64
	var v68 int32
	_ = v68
	var v76 int32
	_ = v76
	var v80 int32
	_ = v80
	var v82 int32
	_ = v82
	var v84 int32
	_ = v84
	var v85 int32
	_ = v85
	var v90 int32
	_ = v90
	var v91 int32
	_ = v91
	var v94 int32
	_ = v94
	var v107 int32
	_ = v107
	var v109 int32
	_ = v109
	var v117 int32
	_ = v117
	var v118 int32
	_ = v118
	var v119 int32
	_ = v119
	var v121 int32
	_ = v121
	var v123 int32
	_ = v123
	var v124 int32
	_ = v124
	var v126 int32
	_ = v126
	var v128 int32
	_ = v128
	var v131 int32
	_ = v131
	var v133 int32
	_ = v133
	var v136 int32
	_ = v136
	var v138 int32
	_ = v138
	var v139 int32
	_ = v139
	var v142 int32
	_ = v142
	var v149 int32
	_ = v149
	var v150 int32
	_ = v150
	var v153 int32
	_ = v153
	var v156 int32
	_ = v156
	var v160 int32
	_ = v160
	var v163 int32
	_ = v163
	var v168 int32
	_ = v168
	var v173 int32
	_ = v173
	var v176 int32
	_ = v176
	var v177 int32
	_ = v177
	var v178 int32
	_ = v178
	var v181 int32
	_ = v181
	var v188 int32
	_ = v188
	var v190 int32
	_ = v190
	var v193 int32
	_ = v193
	var v195 int32
	_ = v195
	var v198 int32
	_ = v198
	var v204 int32
	_ = v204
	var v207 int32
	_ = v207
	var v211 int32
	_ = v211
	var v216 int32
	_ = v216
	v2 = int32(0)
	v10 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v11 = F_pg_detoast_datum_packed(m, v10)
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
	v15 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
	v16 = F_pg_detoast_datum_packed(m, v15)
	mBase = m.M
	v17 = m.ExcPending
	if v17 != 0 {
		goto L1
	} else {
		goto L3
	}
L3:
	;
	v18 = int32(1)
	v20 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v11))))
	if v20&v18 != 0 {
		goto L4
	} else {
		goto L5
	}
L4:
	;
	v23 = v18
	goto L6
L5:
	;
	v23 = int32(4)
	goto L6
L6:
	;
	v24 = v11 + v23
	v25 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v24))))
	v26 = int32(1)
	v28 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v16))))
	if v28&v26 != 0 {
		goto L7
	} else {
		goto L8
	}
L7:
	;
	v31 = v26
	goto L9
L8:
	;
	v31 = int32(4)
	goto L9
L9:
	;
	v32 = v16 + v31
	v33 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v32))))
	if v25 == v33 {
		goto L10
	} else {
		goto L11
	}
L10:
	;
	v35 = int32(2)
	v36 = v24 + v35
	v38 = v32 + v35
	v40 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v24)+1)))
	v41 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v32)+1)))
	if base.Ui32(v40) < base.Ui32(v41) {
		goto L15
	} else {
		goto L16
	}
L11:
	;
	goto L12
L12:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v204 = m.ExcPending
	if v204 != 0 {
		goto L1
	} else {
		goto L62
	}
L13:
	;
	v117 = F_palloc0(m, int32(22))
	mBase = m.M
	v118 = m.ExcPending
	if v118 != 0 {
		goto L1
	} else {
		goto L31
	}
L14:
	;
	v94 = v84
	goto L28
L15:
	;
	v43 = v40
	goto L17
L16:
	;
	v43 = v41
	goto L17
L17:
	;
	if base.Ui32(int32(8)) <= base.Ui32(v43) {
		goto L18
	} else {
		goto L19
	}
L18:
	;
	v47 = int32(base.Ui32(v43) >> (uint(int32(3)) % 32))
	v50 = v2
	goto L21
L19:
	;
	v68 = v2
	goto L20
L20:
	;
	v76 = v43 & int32(7)
	if v76 == int32(0) {
		v107 = int32(0)
		v109 = v68
		goto L13
	} else {
		goto L27
	}
L21:
	;
	v58 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v50+v36))))
	v60 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v50+v38))))
	if v58 != v60 {
		goto L23
	} else {
		goto L24
	}
L22:
	;
	v68 = v47
	goto L20
L23:
	;
	v84 = int32(7)
	v85 = v50
	v90 = v58
	v91 = v60
	goto L14
L24:
	;
	goto L25
L25:
	;
	v64 = v50 + int32(1)
	if v64 != v47 {
		v50 = v64
		goto L21
	} else {
		goto L26
	}
L26:
	;
	goto L22
L27:
	;
	v80 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v68+v38))))
	v82 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v68+v36))))
	v84 = v76
	v85 = v68
	v90 = v82
	v91 = v80
	goto L14
L28:
	;
	if int32(base.Ui32(v90^v91)>>(uint(int32(8)-v94)%32)) != 0 {
		v94 = v94 - int32(1)
		goto L28
	} else {
		goto L30
	}
L29:
	;
	v107 = v94
	v109 = v85
	goto L13
L30:
	;
	goto L29
L31:
	;
	v119 = int32(1)
	v121 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v117))))
	v123 = v121 & v119
	if v123 != 0 {
		goto L32
	} else {
		goto L33
	}
L32:
	;
	v124 = v119
	goto L34
L33:
	;
	v124 = int32(4)
	goto L34
L34:
	;
	v126 = int32(1)
	v128 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v11))))
	if v128&v126 != 0 {
		goto L35
	} else {
		goto L36
	}
L35:
	;
	v131 = v126
	goto L37
L36:
	;
	v131 = int32(4)
	goto L37
L37:
	;
	v133 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v11+v131))))
	*(*uint8)(unsafe.Add(mBase, uint32(v117+v124))) = uint8(v133)
	v136 = v117 + int32(1)
	v138 = v117 + int32(4)
	if v123 != 0 {
		goto L38
	} else {
		goto L39
	}
L38:
	;
	v139 = v136
	goto L40
L39:
	;
	v139 = v138
	goto L40
L40:
	;
	v142 = v107 + v109<<(uint(int32(3))%32)
	*(*uint8)(unsafe.Add(mBase, uint32(v139)+1)) = uint8(v142)
	if v142 <= int32(0) {
		goto L41
	} else {
		goto L42
	}
L41:
	;
	v188 = int32(1)
	v190 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v117))))
	if v190&v188 != 0 {
		goto L56
	} else {
		goto L57
	}
L42:
	;
	v149 = base.I32_div_s(v142+int32(7), int32(8))
	if v149 != 0 {
		goto L43
	} else {
		goto L44
	}
L43:
	;
	v150 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v117))))
	if v150&int32(1) != 0 {
		goto L46
	} else {
		goto L47
	}
L44:
	;
	goto L45
L45:
	;
	v168 = v107 & int32(7)
	if v168 == int32(0) {
		goto L41
	} else {
		goto L52
	}
L46:
	;
	v153 = v136
	goto L48
L47:
	;
	v153 = v138
	goto L48
L48:
	;
	v156 = int32(1)
	v160 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v11))))
	if v160&v156 != 0 {
		goto L49
	} else {
		goto L50
	}
L49:
	;
	v163 = v11 + v156
	goto L51
L50:
	;
	v163 = v11 + int32(4)
	goto L51
L51:
	;
	base.MemoryCopy(m, v153+int32(2), v163+int32(2), v149)
	goto L45
L52:
	;
	v173 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v117))))
	if v173&int32(1) != 0 {
		goto L53
	} else {
		goto L54
	}
L53:
	;
	v176 = v136
	goto L55
L54:
	;
	v176 = v138
	goto L55
L55:
	;
	v177 = int32(base.Ui32(v142)>>(uint(int32(3))%32)) + v176
	v178 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v177)+2)))
	v181 = v178 & (int32(-256) >> (uint(v168) % 32))
	*(*uint8)(unsafe.Add(mBase, uint32(v177)+2)) = uint8(v181)
	goto L41
L56:
	;
	v193 = v188
	goto L58
L57:
	;
	v193 = int32(4)
	goto L58
L58:
	;
	v195 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v117+v193))))
	if v195 == int32(2) {
		goto L59
	} else {
		goto L60
	}
L59:
	;
	v198 = int32(40)
	goto L61
L60:
	;
	v198 = int32(88)
	goto L61
L61:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v117))) = v198
	return v117
L62:
	;
	F_errcode(m, int32(50856066))
	mBase = m.M
	v207 = m.ExcPending
	if v207 != 0 {
		goto L1
	} else {
		goto L63
	}
L63:
	;
	F_errmsg(m, int32(_a_F_inet_merge_0), int32(0))
	mBase = m.M
	v211 = m.ExcPending
	if v211 != 0 {
		goto L1
	} else {
		goto L64
	}
L64:
	;
	F_errfinish(m, int32(_a_F_inet_merge_1), int32(1450), int32(_a_F_inet_merge_2))
	mBase = m.M
	v216 = m.ExcPending
	if v216 != 0 {
		goto L1
	} else {
		goto L65
	}
L65:
	;
	base.Wasm_trap_unreachable()
	for {
	}
}
func F_inet_spg_inner_consistent(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v7 int32
	_ = v7
	var v8 int32
	_ = v8
	var v9 int32
	_ = v9
	var v12 int32
	_ = v12
	var v13 int32
	_ = v13
	var v17 int32
	_ = v17
	var v18 int32
	_ = v18
	var v23 int32
	_ = v23
	var v26 int32
	_ = v26
	var v27 int32
	_ = v27
	var v28 int32
	_ = v28
	var v29 int32
	_ = v29
	var v32 int32
	_ = v32
	var v35 int32
	_ = v35
	var v39 int32
	_ = v39
	var v42 int32
	_ = v42
	var v44 int32
	_ = v44
	var v47 int32
	_ = v47
	var v50 int32
	_ = v50
	var v52 int32
	_ = v52
	var v55 int32
	_ = v55
	var v57 int32
	_ = v57
	var v60 int32
	_ = v60
	var v61 int32
	_ = v61
	var v63 int32
	_ = v63
	var v66 int32
	_ = v66
	var v68 int32
	_ = v68
	var v75 int32
	_ = v75
	var v77 int32
	_ = v77
	var v78 int32
	_ = v78
	var v81 int32
	_ = v81
	var v82 int32
	_ = v82
	var v83 int32
	_ = v83
	var v84 int32
	_ = v84
	var v85 int32
	_ = v85
	var v86 int32
	_ = v86
	var v88 int32
	_ = v88
	var v89 int32
	_ = v89
	var v91 int32
	_ = v91
	var v98 int32
	_ = v98
	var v102 int32
	_ = v102
	var v107 int32
	_ = v107
	var v110 int32
	_ = v110
	var v111 int32
	_ = v111
	var v113 int32
	_ = v113
	var v117 int32
	_ = v117
	var v119 int32
	_ = v119
	var v126 int32
	_ = v126
	var v127 int32
	_ = v127
	var v132 int32
	_ = v132
	var v136 int32
	_ = v136
	var v137 int32
	_ = v137
	var v139 int32
	_ = v139
	v7 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
	v8 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v9 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v8)+34)))
	if v9 == int32(0) {
		goto L5
	} else {
		goto L6
	}
L1:
	;
	return int32(0)
L2:
	;
	v107 = *(*int32)(unsafe.Add(mBase, uint32(v8)+40))
	v110 = F_palloc(m, v107<<(uint(int32(2))%32))
	mBase = m.M
	v111 = m.ExcPending
	if v111 != 0 {
		goto L11
	} else {
		goto L40
	}
L3:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v7))) = int32(0)
	v102 = v98
	goto L2
L4:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v7))) = int32(0)
	if v91 != 0 {
		v102 = v91
		goto L2
	} else {
		goto L39
	}
L5:
	;
	v12 = int32(3)
	v13 = *(*int32)(unsafe.Add(mBase, uint32(v8)+8))
	if v13 <= int32(0) {
		v98 = v12
		goto L3
	} else {
		goto L8
	}
L6:
	;
	goto L7
L7:
	;
	v81 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v8)+33)))
	if v81 != 0 {
		v98 = int32(-1)
		goto L3
	} else {
		goto L36
	}
L8:
	;
	v17 = int32(0)
	v18 = v12
	goto L9
L9:
	;
	v23 = *(*int32)(unsafe.Add(mBase, uint32(v8)))
	v26 = v23 + v17*int32(48)
	v27 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v26)+6)))
	v28 = *(*int32)(unsafe.Add(mBase, uint32(v26)+44))
	v29 = F_pg_detoast_datum_packed(m, v28)
	mBase = m.M
	v32 = m.ExcPending
	if v32 != 0 {
		goto L11
	} else {
		goto L12
	}
L10:
	;
	v91 = v75
	goto L4
L11:
	;
	return int32(0)
L12:
	;
	switch v27 - int32(19) {
	case 0:
		v75 = v18
		goto L13
	case 1, 2:
		goto L16
	case 3, 4:
		goto L15
	default:
		goto L14
	}
L13:
	;
	v77 = v17 + int32(1)
	v78 = *(*int32)(unsafe.Add(mBase, uint32(v8)+8))
	if v77 < v78 {
		v17 = v77
		v18 = v75
		goto L9
	} else {
		goto L35
	}
L14:
	;
	v61 = int32(1)
	v63 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v29))))
	if v63&v61 != 0 {
		goto L29
	} else {
		goto L30
	}
L15:
	;
	v50 = int32(1)
	v52 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v29))))
	if v52&v50 != 0 {
		goto L23
	} else {
		goto L24
	}
L16:
	;
	v35 = int32(1)
	v39 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v29))))
	if v39&v35 != 0 {
		goto L17
	} else {
		goto L18
	}
L17:
	;
	v42 = v35
	goto L19
L18:
	;
	v42 = int32(4)
	goto L19
L19:
	;
	v44 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v29+v42))))
	if v44 == int32(2) {
		goto L20
	} else {
		goto L21
	}
L20:
	;
	v47 = v18 & v35
	goto L22
L21:
	;
	v47 = v18
	goto L22
L22:
	;
	v75 = v47
	goto L13
L23:
	;
	v55 = v50
	goto L25
L24:
	;
	v55 = int32(4)
	goto L25
L25:
	;
	v57 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v29+v55))))
	if v57 == int32(3) {
		goto L26
	} else {
		goto L27
	}
L26:
	;
	v60 = v18 & int32(2)
	goto L28
L27:
	;
	v60 = v18
	goto L28
L28:
	;
	v75 = v60
	goto L13
L29:
	;
	v66 = v61
	goto L31
L30:
	;
	v66 = int32(4)
	goto L31
L31:
	;
	v68 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v29+v66))))
	if v68 == int32(2) {
		goto L32
	} else {
		goto L33
	}
L32:
	;
	v75 = v18 & int32(1)
	goto L13
L33:
	;
	goto L34
L34:
	;
	v75 = v18 & int32(2)
	goto L13
L35:
	;
	goto L10
L36:
	;
	v82 = *(*int32)(unsafe.Add(mBase, uint32(v8)+36))
	v83 = F_pg_detoast_datum_packed(m, v82)
	mBase = m.M
	v84 = m.ExcPending
	if v84 != 0 {
		goto L11
	} else {
		goto L37
	}
L37:
	;
	v85 = *(*int32)(unsafe.Add(mBase, uint32(v8)+8))
	v86 = *(*int32)(unsafe.Add(mBase, uint32(v8)))
	v88 = F_inet_spg_consistent_bitmap(m, v83, v85, v86, int32(0))
	mBase = m.M
	v89 = m.ExcPending
	if v89 != 0 {
		goto L11
	} else {
		goto L38
	}
L38:
	;
	v91 = v88
	goto L4
L39:
	;
	goto L1
L40:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v7)+4)) = v110
	v113 = *(*int32)(unsafe.Add(mBase, uint32(v8)+40))
	if v113 <= int32(0) {
		goto L1
	} else {
		goto L41
	}
L41:
	;
	v117 = int32(0)
	v119 = v113
	goto L42
L42:
	;
	if int32(base.Ui32(v102)>>(uint(v117)%32))&int32(1) != 0 {
		goto L44
	} else {
		goto L45
	}
L43:
	;
	goto L1
L44:
	;
	v126 = *(*int32)(unsafe.Add(mBase, uint32(v7)+4))
	v127 = *(*int32)(unsafe.Add(mBase, uint32(v7)))
	*(*int32)(unsafe.Add(mBase, uint32(v126+v127<<(uint(int32(2))%32)))) = v117
	v132 = *(*int32)(unsafe.Add(mBase, uint32(v7)))
	*(*int32)(unsafe.Add(mBase, uint32(v7))) = v132 + int32(1)
	v136 = *(*int32)(unsafe.Add(mBase, uint32(v8)+40))
	v137 = v136
	goto L46
L45:
	;
	v137 = v119
	goto L46
L46:
	;
	v139 = v117 + int32(1)
	if v139 < v137 {
		v117 = v139
		v119 = v137
		goto L42
	} else {
		goto L47
	}
L47:
	;
	goto L43
}
