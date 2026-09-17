package p5

import (
	base "github.com/shibukawa/pgmem/internal/aot/pgaot/base"
	"unsafe"
)

func F_TidListEval(m *base.Module, l0 int32) {
	mBase := m.M
	_ = mBase
	var v2 int32
	_ = v2
	var v21 int32
	_ = v21
	var v23 int32
	_ = v23
	var v25 int32
	_ = v25
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
	var v36 int32
	_ = v36
	var v37 int32
	_ = v37
	var v38 int32
	_ = v38
	var v39 int32
	_ = v39
	var v42 int32
	_ = v42
	var v43 int32
	_ = v43
	var v44 int32
	_ = v44
	var v45 int32
	_ = v45
	var v48 int32
	_ = v48
	var v49 int32
	_ = v49
	var v50 int32
	_ = v50
	var v53 int32
	_ = v53
	var v61 int32
	_ = v61
	var v63 int32
	_ = v63
	var v64 int32
	_ = v64
	var v65 int32
	_ = v65
	var v76 int32
	_ = v76
	var v80 int32
	_ = v80
	var v81 int32
	_ = v81
	var v84 int32
	_ = v84
	var v85 int32
	_ = v85
	var v86 int32
	_ = v86
	var v87 int32
	_ = v87
	var v89 int32
	_ = v89
	var v91 int32
	_ = v91
	var v94 int32
	_ = v94
	var v97 int32
	_ = v97
	var v100 int32
	_ = v100
	var v102 int32
	_ = v102
	var v106 int32
	_ = v106
	var v107 int32
	_ = v107
	var v113 int32
	_ = v113
	var v114 int32
	_ = v114
	var v117 int32
	_ = v117
	var v123 int32
	_ = v123
	var v126 int32
	_ = v126
	var v127 int32
	_ = v127
	var v128 int32
	_ = v128
	var v129 int32
	_ = v129
	var v131 int32
	_ = v131
	var v132 int32
	_ = v132
	var v140 int32
	_ = v140
	var v145 int32
	_ = v145
	var v151 int32
	_ = v151
	var v154 int32
	_ = v154
	var v158 int32
	_ = v158
	var v163 int32
	_ = v163
	var v164 int32
	_ = v164
	var v165 int32
	_ = v165
	var v166 int32
	_ = v166
	var v169 int32
	_ = v169
	var v171 int32
	_ = v171
	var v172 int32
	_ = v172
	var v173 int32
	_ = v173
	var v174 int32
	_ = v174
	var v175 int32
	_ = v175
	var v178 int32
	_ = v178
	var v179 int32
	_ = v179
	var v182 int32
	_ = v182
	var v185 int32
	_ = v185
	var v186 int32
	_ = v186
	var v189 int32
	_ = v189
	var v192 int32
	_ = v192
	var v208 int32
	_ = v208
	var v214 int32
	_ = v214
	var v217 int32
	_ = v217
	var v220 int32
	_ = v220
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
	var v238 int32
	_ = v238
	var v240 int32
	_ = v240
	var v242 int32
	_ = v242
	var v244 int32
	_ = v244
	var v252 int32
	_ = v252
	var v255 int32
	_ = v255
	var v262 int32
	_ = v262
	var v263 int32
	_ = v263
	var v266 int32
	_ = v266
	var v270 int32
	_ = v270
	var v271 int32
	_ = v271
	var v272 int32
	_ = v272
	var v280 int32
	_ = v280
	var v281 int32
	_ = v281
	var v283 int32
	_ = v283
	var v284 int32
	_ = v284
	var v287 int32
	_ = v287
	var v289 int32
	_ = v289
	var v290 int32
	_ = v290
	var v294 int32
	_ = v294
	var v300 int32
	_ = v300
	var v303 int32
	_ = v303
	var v308 int32
	_ = v308
	var v312 int32
	_ = v312
	var v313 int32
	_ = v313
	var v321 int32
	_ = v321
	var v329 int32
	_ = v329
	var v332 int32
	_ = v332
	var v333 int32
	_ = v333
	var v336 int32
	_ = v336
	var v337 int32
	_ = v337
	var v340 int32
	_ = v340
	var v343 int32
	_ = v343
	var v346 int32
	_ = v346
	var v349 int32
	_ = v349
	var v354 int32
	_ = v354
	var v355 int32
	_ = v355
	var v357 int32
	_ = v357
	var v381 int32
	_ = v381
	var v406 int32
	_ = v406
	var v412 int32
	_ = v412
	var v417 int32
	_ = v417
	var v421 int32
	_ = v421
	var v424 int32
	_ = v424
	var v430 int32
	_ = v430
	var v435 int32
	_ = v435
	var v439 int32
	_ = v439
	var v442 int32
	_ = v442
	var v448 int32
	_ = v448
	var v453 int32
	_ = v453
	var v458 int32
	_ = v458
	var v461 int32
	_ = v461
	var v467 int32
	_ = v467
	var v472 int32
	_ = v472
	var v476 int32
	_ = v476
	var v479 int32
	_ = v479
	var v486 int32
	_ = v486
	var v491 int32
	_ = v491
	var v515 int32
	_ = v515
	var v518 int32
	_ = v518
	var v525 int32
	_ = v525
	var v530 int32
	_ = v530
	var v534 int32
	_ = v534
	var v537 int32
	_ = v537
	var v543 int32
	_ = v543
	var v548 int32
	_ = v548
	var v552 int32
	_ = v552
	var v555 int32
	_ = v555
	var v562 int32
	_ = v562
	var v567 int32
	_ = v567
	var v571 int32
	_ = v571
	var v574 int32
	_ = v574
	var v580 int32
	_ = v580
	var v585 int32
	_ = v585
	var v591 int32
	_ = v591
	var v592 int32
	_ = v592
	var v595 int32
	_ = v595
	var v596 int32
	_ = v596
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
	var v612 int32
	_ = v612
	var v616 int32
	_ = v616
	var v617 int32
	_ = v617
	var v618 int32
	_ = v618
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
	var v633 int32
	_ = v633
	var v636 int32
	_ = v636
	var v637 int32
	_ = v637
	var v640 int32
	_ = v640
	var v641 int32
	_ = v641
	var v643 int32
	_ = v643
	var v647 int32
	_ = v647
	var v648 int32
	_ = v648
	var v650 int32
	_ = v650
	var v654 int32
	_ = v654
	var v655 int32
	_ = v655
	var v656 int32
	_ = v656
	var v659 int32
	_ = v659
	var v660 int32
	_ = v660
	var v661 int32
	_ = v661
	var v670 int32
	_ = v670
	var v671 int32
	_ = v671
	var v672 int32
	_ = v672
	var v676 int32
	_ = v676
	var v677 int32
	_ = v677
	var v678 int32
	_ = v678
	var v679 int32
	_ = v679
	var v680 int32
	_ = v680
	var v681 int32
	_ = v681
	var v682 int32
	_ = v682
	var v687 int32
	_ = v687
	var v692 int32
	_ = v692
	var v705 int32
	_ = v705
	var v707 int32
	_ = v707
	var v708 int32
	_ = v708
	var v712 int32
	_ = v712
	var v713 int32
	_ = v713
	var v714 int32
	_ = v714
	var v715 int32
	_ = v715
	var v716 int32
	_ = v716
	var v717 int32
	_ = v717
	var v722 int32
	_ = v722
	var v723 int32
	_ = v723
	var v725 int32
	_ = v725
	var v731 int32
	_ = v731
	var v733 int32
	_ = v733
	var v734 int32
	_ = v734
	var v743 int32
	_ = v743
	var v756 int32
	_ = v756
	var v758 int32
	_ = v758
	var v759 int32
	_ = v759
	var v761 int32
	_ = v761
	var v767 int32
	_ = v767
	var v769 int32
	_ = v769
	var v770 int32
	_ = v770
	var v783 int32
	_ = v783
	var v784 int32
	_ = v784
	var v793 int32
	_ = v793
	var v794 int32
	_ = v794
	var v811 int32
	_ = v811
	var v816 int32
	_ = v816
	var v823 int32
	_ = v823
	var v834 int32
	_ = v834
	var v836 int32
	_ = v836
	var v837 int32
	_ = v837
	var v838 int32
	_ = v838
	var v840 int32
	_ = v840
	var v844 int32
	_ = v844
	var v845 int32
	_ = v845
	var v848 int32
	_ = v848
	var v851 int32
	_ = v851
	var v852 int32
	_ = v852
	var v855 int32
	_ = v855
	var v859 int32
	_ = v859
	var v860 int32
	_ = v860
	var v862 int32
	_ = v862
	var v865 int32
	_ = v865
	var v867 int32
	_ = v867
	var v878 int32
	_ = v878
	var v879 int32
	_ = v879
	v2 = int32(0)
	v21 = m.G0
	v23 = v21 - int32(32)
	m.G0 = v23
	v25 = *(*int32)(unsafe.Add(mBase, uint32(l0)+64))
	v26 = *(*int32)(unsafe.Add(mBase, uint32(l0)+108))
	if v26 == v2 {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	v29 = *(*int32)(unsafe.Add(mBase, uint32(l0)+104))
	v30 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	v31 = *(*int32)(unsafe.Add(mBase, uint32(v30)+8))
	v32 = int32(0)
	v36 = *(*int32)(unsafe.Add(mBase, uint32(v29)+188))
	v37 = *(*int32)(unsafe.Add(mBase, uint32(v36)+8))
	v38 = m.T0[v37].(func(*base.Module, int32, int32, int32, int32, int32, int32) int32)(m, v29, v31, v32, v32, v32, int32(8))
	mBase = m.M
	v39 = m.ExcPending
	if v39 != 0 {
		goto L4
	} else {
		goto L5
	}
L2:
	;
	v42 = v26
	goto L3
L3:
	;
	v43 = *(*int32)(unsafe.Add(mBase, uint32(l0)+116))
	if v43 != 0 {
		goto L6
	} else {
		goto L7
	}
L4:
	;
	return
L5:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+108)) = v38
	v42 = v38
	goto L3
L6:
	;
	v44 = *(*int32)(unsafe.Add(mBase, uint32(v43)+4))
	v45 = v44
	goto L8
L7:
	;
	v45 = v2
	goto L8
L8:
	;
	v48 = F_palloc(m, v45*int32(6))
	mBase = m.M
	v49 = m.ExcPending
	if v49 != 0 {
		goto L4
	} else {
		goto L9
	}
L9:
	;
	v50 = *(*int32)(unsafe.Add(mBase, uint32(l0)+116))
	if v50 == int32(0) {
		v878 = v2
		v879 = v48
		goto L10
	} else {
		goto L11
	}
L10:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+128)) = int32(-1)
	*(*int32)(unsafe.Add(mBase, uint32(l0)+124)) = v878
	*(*int32)(unsafe.Add(mBase, uint32(l0)+132)) = v879
	m.G0 = v23 + int32(32)
	return
L11:
	;
	v53 = *(*int32)(unsafe.Add(mBase, uint32(v50)+4))
	if int32(0) < v53 {
		goto L12
	} else {
		goto L13
	}
L12:
	;
	v61 = v45
	v63 = v2
	v64 = v48
	v65 = v2
	goto L15
L13:
	;
	v793 = v2
	v794 = v48
	goto L14
L14:
	;
	if v793 <= int32(1) {
		v878 = v793
		v879 = v794
		goto L10
	} else {
		goto L202
	}
L15:
	;
	v76 = *(*int32)(unsafe.Add(mBase, uint32(v50)+12))
	v80 = *(*int32)(unsafe.Add(mBase, uint32(v76+v65<<(uint(int32(2))%32))))
	v81 = *(*int32)(unsafe.Add(mBase, uint32(v80)))
	if v81 == int32(0) {
		goto L18
	} else {
		goto L19
	}
L16:
	;
	v793 = v769
	v794 = v770
	goto L14
L17:
	;
	v783 = v65 + int32(1)
	v784 = *(*int32)(unsafe.Add(mBase, uint32(v50)+4))
	if v783 < v784 {
		v61 = v767
		v63 = v769
		v64 = v770
		v65 = v783
		goto L15
	} else {
		goto L201
	}
L18:
	;
	v84 = *(*int32)(unsafe.Add(mBase, uint32(v80)+8))
	v85 = *(*int32)(unsafe.Add(mBase, uint32(l0)+104))
	v86 = *(*int32)(unsafe.Add(mBase, uint32(v85)+56))
	v87 = m.G0
	v89 = v87 - int32(192)
	m.G0 = v89
	v91 = *(*int32)(unsafe.Add(mBase, uint32(v84)+8))
	if v91 == int32(0) {
		goto L22
	} else {
		goto L23
	}
L19:
	;
	goto L20
L20:
	;
	v606 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v80)+4)))
	if v606 == int32(0) {
		goto L170
	} else {
		goto L171
	}
L21:
	;
	if v381 == int32(0) {
		v767 = v61
		v769 = v63
		v770 = v64
		goto L17
	} else {
		goto L165
	}
L22:
	;
	v94 = *(*int32)(unsafe.Add(mBase, uint32(v84)+12))
	if v94 <= int32(0) {
		goto L26
	} else {
		goto L27
	}
L23:
	;
	v169 = v91
	goto L24
L24:
	;
	v171 = v23 + int32(12)
	v172 = F_get_rel_name(m, v86)
	mBase = m.M
	v173 = m.ExcPending
	if v173 != 0 {
		goto L4
	} else {
		goto L57
	}
L25:
	;
	v164 = *(*int32)(unsafe.Add(mBase, uint32(v113)))
	v165 = F_text_to_cstring(m, v164)
	mBase = m.M
	v166 = m.ExcPending
	if v166 != 0 {
		goto L4
	} else {
		goto L48
	}
L26:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v151 = m.ExcPending
	if v151 != 0 {
		goto L4
	} else {
		goto L44
	}
L27:
	;
	v97 = *(*int32)(unsafe.Add(mBase, uint32(v25)+28))
	if v97 == int32(0) {
		goto L26
	} else {
		goto L28
	}
L28:
	;
	v100 = *(*int32)(unsafe.Add(mBase, uint32(v97)+28))
	if v100 < v94 {
		goto L26
	} else {
		goto L29
	}
L29:
	;
	v102 = *(*int32)(unsafe.Add(mBase, uint32(v97)))
	if v102 != 0 {
		goto L31
	} else {
		goto L32
	}
L30:
	;
	v114 = *(*int32)(unsafe.Add(mBase, uint32(v113)+8))
	if v114 == int32(0) {
		goto L26
	} else {
		goto L35
	}
L31:
	;
	v106 = m.T0[v102].(func(*base.Module, int32, int32, int32, int32) int32)(m, v97, v94, int32(0), v89+int32(180))
	mBase = m.M
	v107 = m.ExcPending
	if v107 != 0 {
		goto L4
	} else {
		goto L34
	}
L32:
	;
	goto L33
L33:
	;
	v113 = v97 + v94*int32(12) + int32(20)
	goto L30
L34:
	;
	v113 = v106
	goto L30
L35:
	;
	v117 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v113)+4)))
	if v117 != 0 {
		goto L26
	} else {
		goto L36
	}
L36:
	;
	if v114 == int32(1790) {
		goto L25
	} else {
		goto L37
	}
L37:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v123 = m.ExcPending
	if v123 != 0 {
		goto L4
	} else {
		goto L38
	}
L38:
	;
	F_errcode(m, int32(67141764))
	mBase = m.M
	v126 = m.ExcPending
	if v126 != 0 {
		goto L4
	} else {
		goto L39
	}
L39:
	;
	v127 = *(*int32)(unsafe.Add(mBase, uint32(v113)+8))
	v128 = F_format_type_be(m, v127)
	mBase = m.M
	v129 = m.ExcPending
	if v129 != 0 {
		goto L4
	} else {
		goto L40
	}
L40:
	;
	v131 = F_format_type_be(m, int32(1790))
	mBase = m.M
	v132 = m.ExcPending
	if v132 != 0 {
		goto L4
	} else {
		goto L41
	}
L41:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v89)+168)) = v131
	*(*int32)(unsafe.Add(mBase, uint32(v89)+164)) = v128
	*(*int32)(unsafe.Add(mBase, uint32(v89)+160)) = v94
	F_errmsg(m, int32(_a_F_TidListEval_0), v89+int32(160))
	mBase = m.M
	v140 = m.ExcPending
	if v140 != 0 {
		goto L4
	} else {
		goto L42
	}
L42:
	;
	F_errfinish(m, int32(_a_F_TidListEval_1), int32(283), int32(_a_F_TidListEval_2))
	mBase = m.M
	v145 = m.ExcPending
	if v145 != 0 {
		goto L4
	} else {
		goto L43
	}
L43:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L44:
	;
	F_errcode(m, int32(67137668))
	mBase = m.M
	v154 = m.ExcPending
	if v154 != 0 {
		goto L4
	} else {
		goto L45
	}
L45:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v89))) = v94
	F_errmsg(m, int32(_a_F_TidListEval_3), v89)
	mBase = m.M
	v158 = m.ExcPending
	if v158 != 0 {
		goto L4
	} else {
		goto L46
	}
L46:
	;
	F_errfinish(m, int32(_a_F_TidListEval_1), int32(292), int32(_a_F_TidListEval_2))
	mBase = m.M
	v163 = m.ExcPending
	if v163 != 0 {
		goto L4
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
	v169 = v165
	goto L24
L49:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v571 = m.ExcPending
	if v571 != 0 {
		goto L4
	} else {
		goto L161
	}
L50:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v552 = m.ExcPending
	if v552 != 0 {
		goto L4
	} else {
		goto L157
	}
L51:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v534 = m.ExcPending
	if v534 != 0 {
		goto L4
	} else {
		goto L153
	}
L52:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v515 = m.ExcPending
	if v515 != 0 {
		goto L4
	} else {
		goto L149
	}
L53:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v476 = m.ExcPending
	if v476 != 0 {
		goto L4
	} else {
		goto L145
	}
L54:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v458 = m.ExcPending
	if v458 != 0 {
		goto L4
	} else {
		goto L141
	}
L55:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v439 = m.ExcPending
	if v439 != 0 {
		goto L4
	} else {
		goto L137
	}
L56:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v421 = m.ExcPending
	if v421 != 0 {
		goto L4
	} else {
		goto L133
	}
L57:
	;
	if v172 != 0 {
		goto L58
	} else {
		goto L59
	}
L58:
	;
	v174 = F_GetPortalByName(m, v169)
	mBase = m.M
	v175 = m.ExcPending
	if v175 != 0 {
		goto L4
	} else {
		goto L61
	}
L59:
	;
	goto L60
L60:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v406 = m.ExcPending
	if v406 != 0 {
		goto L4
	} else {
		goto L130
	}
L61:
	;
	if v174 == int32(0) {
		goto L56
	} else {
		goto L62
	}
L62:
	;
	v178 = *(*int32)(unsafe.Add(mBase, uint32(v174)+72))
	if v178 != 0 {
		goto L55
	} else {
		goto L63
	}
L63:
	;
	v179 = *(*int32)(unsafe.Add(mBase, uint32(v174)+88))
	if v179 == int32(0) {
		goto L54
	} else {
		goto L64
	}
L64:
	;
	v182 = *(*int32)(unsafe.Add(mBase, uint32(v179)+40))
	if v182 == int32(0) {
		goto L54
	} else {
		goto L65
	}
L65:
	;
	v185 = *(*int32)(unsafe.Add(mBase, uint32(v182)+28))
	if v185 != 0 {
		goto L68
	} else {
		goto L69
	}
L66:
	;
	m.G0 = v89 + int32(192)
	goto L21
L67:
	;
	v381 = int32(1)
	goto L66
L68:
	;
	v186 = *(*int32)(unsafe.Add(mBase, uint32(v182)+20))
	if v186 == int32(0) {
		goto L52
	} else {
		goto L71
	}
L69:
	;
	goto L70
L70:
	;
	v240 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v89)+180)) = uint8(v240)
	v242 = *(*int32)(unsafe.Add(mBase, uint32(v179)+44))
	v244 = v89 + int32(180)
	if v242 == v240 {
		v321 = v240
		goto L93
	} else {
		goto L94
	}
L71:
	;
	v189 = int32(0)
	v192 = v189
	v208 = v189
	goto L72
L72:
	;
	v214 = *(*int32)(unsafe.Add(mBase, uint32(v185+v208<<(uint(int32(2))%32))))
	if v214 == int32(0) {
		goto L75
	} else {
		goto L76
	}
L73:
	;
	if v222 == int32(0) {
		goto L52
	} else {
		goto L86
	}
L74:
	;
	v224 = v208 + int32(1)
	if v224 != v186 {
		v192 = v222
		v208 = v224
		goto L72
	} else {
		goto L85
	}
L75:
	;
	v222 = v192
	goto L74
L76:
	;
	goto L77
L77:
	;
	v217 = *(*int32)(unsafe.Add(mBase, uint32(v214)+20))
	if base.Ui32(int32(3)) < base.Ui32(v217) {
		goto L78
	} else {
		goto L79
	}
L78:
	;
	v222 = v192
	goto L74
L79:
	;
	goto L80
L80:
	;
	v220 = *(*int32)(unsafe.Add(mBase, uint32(v214)+4))
	if v86 != v220 {
		goto L81
	} else {
		goto L82
	}
L81:
	;
	v222 = v192
	goto L74
L82:
	;
	goto L83
L83:
	;
	if v192 != 0 {
		goto L53
	} else {
		goto L84
	}
L84:
	;
	v222 = v214
	goto L74
L85:
	;
	goto L73
L86:
	;
	v228 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v174)+116)))
	if v228 != 0 {
		goto L51
	} else {
		goto L87
	}
L87:
	;
	v229 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v174)+117)))
	if v229 == int32(1) {
		goto L51
	} else {
		goto L88
	}
L88:
	;
	v232 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v222)+38)))
	if v232 == int32(0) {
		goto L89
	} else {
		goto L90
	}
L89:
	;
	v381 = int32(0)
	goto L66
L90:
	;
	goto L91
L91:
	;
	v236 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v222)+38)))
	*(*uint16)(unsafe.Add(mBase, uint32(v171)+4)) = uint16(v236)
	v238 = *(*int32)(unsafe.Add(mBase, uint32(v222)+34))
	*(*int32)(unsafe.Add(mBase, uint32(v171))) = v238
	goto L67
L92:
	;
	if v329 == int32(0) {
		goto L50
	} else {
		goto L120
	}
L93:
	;
	v329 = v321
	goto L92
L94:
	;
	v252 = *(*int32)(unsafe.Add(mBase, uint32(v242)))
	switch v252 - int32(394) {
	case 0, 43:
		v287 = int32(36)
		goto L98
	default:
		v321 = v240
		goto L93
	case 3:
		goto L100
	case 9, 10, 11, 12, 14, 15, 16, 24, 25:
		goto L96
	case 17:
		goto L99
	}
L95:
	;
	v312 = *(*int32)(unsafe.Add(mBase, uint32(v242)+52))
	if v312 != 0 {
		goto L117
	} else {
		goto L118
	}
L96:
	;
	v300 = *(*int32)(unsafe.Add(mBase, uint32(v242)+104))
	if v300 == int32(0) {
		v321 = v240
		goto L93
	} else {
		goto L115
	}
L97:
	;
	if v294 == int32(0) {
		v321 = v240
		goto L93
	} else {
		goto L114
	}
L98:
	;
	v289 = *(*int32)(unsafe.Add(mBase, uint32(v242+v287)))
	v290 = F_search_plan_tree(m, v289, v86, v244)
	mBase = m.M
	v294 = v290
	goto L97
L99:
	;
	v287 = int32(116)
	goto L98
L100:
	;
	v255 = *(*int32)(unsafe.Add(mBase, uint32(v242)+108))
	if v255 <= int32(0) {
		v321 = v240
		goto L93
	} else {
		goto L101
	}
L101:
	;
	v262 = int32(0)
	v263 = v240
	goto L102
L102:
	;
	v266 = *(*int32)(unsafe.Add(mBase, uint32(v242)+104))
	v270 = *(*int32)(unsafe.Add(mBase, uint32(v266+v263<<(uint(int32(2))%32))))
	v271 = F_search_plan_tree(m, v270, v86, v244)
	mBase = m.M
	v272 = int32(0)
	if base.B2i32(v271 == v272)|base.B2i32(v262 == v272) == v272 {
		goto L104
	} else {
		goto L105
	}
L103:
	;
	v294 = v281
	goto L97
L104:
	;
	v329 = int32(0)
	goto L92
L105:
	;
	goto L106
L106:
	;
	if v262 != 0 {
		goto L107
	} else {
		goto L108
	}
L107:
	;
	v280 = v262
	goto L109
L108:
	;
	v280 = v271
	goto L109
L109:
	;
	if v271 != 0 {
		goto L110
	} else {
		goto L111
	}
L110:
	;
	v281 = v280
	goto L112
L111:
	;
	v281 = v262
	goto L112
L112:
	;
	v283 = v263 + int32(1)
	v284 = *(*int32)(unsafe.Add(mBase, uint32(v242)+108))
	if v283 < v284 {
		v262 = v281
		v263 = v283
		goto L102
	} else {
		goto L113
	}
L113:
	;
	goto L103
L114:
	;
	v308 = v294
	goto L95
L115:
	;
	v303 = *(*int32)(unsafe.Add(mBase, uint32(v300)+56))
	if v303 != v86 {
		v321 = v240
		goto L93
	} else {
		goto L116
	}
L116:
	;
	v308 = v242
	goto L95
L117:
	;
	v313 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(v244))) = uint8(v313)
	goto L119
L118:
	;
	goto L119
L119:
	;
	v321 = v308
	goto L93
L120:
	;
	v332 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v174)+116)))
	if v332 != 0 {
		goto L49
	} else {
		goto L121
	}
L121:
	;
	v333 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v174)+117)))
	if v333 == int32(1) {
		goto L49
	} else {
		goto L122
	}
L122:
	;
	v336 = int32(0)
	v337 = *(*int32)(unsafe.Add(mBase, uint32(v329)+112))
	if v337 == v336 {
		v381 = v336
		goto L66
	} else {
		goto L123
	}
L123:
	;
	v340 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v337)+4)))
	if v340&int32(2) != 0 {
		v381 = v336
		goto L66
	} else {
		goto L124
	}
L124:
	;
	v343 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v89)+180)))
	if v343&int32(1) != 0 {
		v381 = v336
		goto L66
	} else {
		goto L125
	}
L125:
	;
	v346 = *(*int32)(unsafe.Add(mBase, uint32(v329)))
	if v346 == int32(406) {
		goto L127
	} else {
		goto L128
	}
L126:
	;
	v355 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v354)+4)))
	*(*uint16)(unsafe.Add(mBase, uint32(v171)+4)) = uint16(v355)
	v357 = *(*int32)(unsafe.Add(mBase, uint32(v354)))
	*(*int32)(unsafe.Add(mBase, uint32(v171))) = v357
	goto L67
L127:
	;
	v349 = *(*int32)(unsafe.Add(mBase, uint32(v329)+156))
	v354 = v349 + int32(60)
	goto L126
L128:
	;
	goto L129
L129:
	;
	v354 = v337 + int32(28)
	goto L126
L130:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v89)+16)) = v86
	F_errmsg_internal(m, int32(_a_F_TidListEval_4), v89+int32(16))
	mBase = m.M
	v412 = m.ExcPending
	if v412 != 0 {
		goto L4
	} else {
		goto L131
	}
L131:
	;
	F_errfinish(m, int32(_a_F_TidListEval_1), int32(63), int32(_a_F_TidListEval_5))
	mBase = m.M
	v417 = m.ExcPending
	if v417 != 0 {
		goto L4
	} else {
		goto L132
	}
L132:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L133:
	;
	F_errcode(m, int32(259))
	mBase = m.M
	v424 = m.ExcPending
	if v424 != 0 {
		goto L4
	} else {
		goto L134
	}
L134:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v89)+32)) = v169
	F_errmsg(m, int32(_a_F_TidListEval_6), v89+int32(32))
	mBase = m.M
	v430 = m.ExcPending
	if v430 != 0 {
		goto L4
	} else {
		goto L135
	}
L135:
	;
	F_errfinish(m, int32(_a_F_TidListEval_1), int32(70), int32(_a_F_TidListEval_5))
	mBase = m.M
	v435 = m.ExcPending
	if v435 != 0 {
		goto L4
	} else {
		goto L136
	}
L136:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L137:
	;
	F_errcode(m, int32(258))
	mBase = m.M
	v442 = m.ExcPending
	if v442 != 0 {
		goto L4
	} else {
		goto L138
	}
L138:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v89)+144)) = v169
	F_errmsg(m, int32(_a_F_TidListEval_7), v89+int32(144))
	mBase = m.M
	v448 = m.ExcPending
	if v448 != 0 {
		goto L4
	} else {
		goto L139
	}
L139:
	;
	F_errfinish(m, int32(_a_F_TidListEval_1), int32(80), int32(_a_F_TidListEval_5))
	mBase = m.M
	v453 = m.ExcPending
	if v453 != 0 {
		goto L4
	} else {
		goto L140
	}
L140:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L141:
	;
	F_errcode(m, int32(258))
	mBase = m.M
	v461 = m.ExcPending
	if v461 != 0 {
		goto L4
	} else {
		goto L142
	}
L142:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v89)+48)) = v169
	F_errmsg(m, int32(_a_F_TidListEval_8), v89+int32(48))
	mBase = m.M
	v467 = m.ExcPending
	if v467 != 0 {
		goto L4
	} else {
		goto L143
	}
L143:
	;
	F_errfinish(m, int32(_a_F_TidListEval_1), int32(86), int32(_a_F_TidListEval_5))
	mBase = m.M
	v472 = m.ExcPending
	if v472 != 0 {
		goto L4
	} else {
		goto L144
	}
L144:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L145:
	;
	F_errcode(m, int32(258))
	mBase = m.M
	v479 = m.ExcPending
	if v479 != 0 {
		goto L4
	} else {
		goto L146
	}
L146:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v89)+132)) = v172
	*(*int32)(unsafe.Add(mBase, uint32(v89)+128)) = v169
	F_errmsg(m, int32(_a_F_TidListEval_9), v89+int32(128))
	mBase = m.M
	v486 = m.ExcPending
	if v486 != 0 {
		goto L4
	} else {
		goto L147
	}
L147:
	;
	F_errfinish(m, int32(_a_F_TidListEval_1), int32(119), int32(_a_F_TidListEval_5))
	mBase = m.M
	v491 = m.ExcPending
	if v491 != 0 {
		goto L4
	} else {
		goto L148
	}
L148:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L149:
	;
	F_errcode(m, int32(258))
	mBase = m.M
	v518 = m.ExcPending
	if v518 != 0 {
		goto L4
	} else {
		goto L150
	}
L150:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v89)+100)) = v172
	*(*int32)(unsafe.Add(mBase, uint32(v89)+96)) = v169
	F_errmsg(m, int32(_a_F_TidListEval_10), v89+int32(96))
	mBase = m.M
	v525 = m.ExcPending
	if v525 != 0 {
		goto L4
	} else {
		goto L151
	}
L151:
	;
	F_errfinish(m, int32(_a_F_TidListEval_1), int32(128), int32(_a_F_TidListEval_5))
	mBase = m.M
	v530 = m.ExcPending
	if v530 != 0 {
		goto L4
	} else {
		goto L152
	}
L152:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L153:
	;
	F_errcode(m, int32(258))
	mBase = m.M
	v537 = m.ExcPending
	if v537 != 0 {
		goto L4
	} else {
		goto L154
	}
L154:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v89)+112)) = v169
	F_errmsg(m, int32(_a_F_TidListEval_11), v89+int32(112))
	mBase = m.M
	v543 = m.ExcPending
	if v543 != 0 {
		goto L4
	} else {
		goto L155
	}
L155:
	;
	F_errfinish(m, int32(_a_F_TidListEval_1), int32(138), int32(_a_F_TidListEval_5))
	mBase = m.M
	v548 = m.ExcPending
	if v548 != 0 {
		goto L4
	} else {
		goto L156
	}
L156:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L157:
	;
	F_errcode(m, int32(258))
	mBase = m.M
	v555 = m.ExcPending
	if v555 != 0 {
		goto L4
	} else {
		goto L158
	}
L158:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v89)+68)) = v172
	*(*int32)(unsafe.Add(mBase, uint32(v89)+64)) = v169
	F_errmsg(m, int32(_a_F_TidListEval_12), v89-int32(-64))
	mBase = m.M
	v562 = m.ExcPending
	if v562 != 0 {
		goto L4
	} else {
		goto L159
	}
L159:
	;
	F_errfinish(m, int32(_a_F_TidListEval_1), int32(170), int32(_a_F_TidListEval_5))
	mBase = m.M
	v567 = m.ExcPending
	if v567 != 0 {
		goto L4
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
	F_errcode(m, int32(258))
	mBase = m.M
	v574 = m.ExcPending
	if v574 != 0 {
		goto L4
	} else {
		goto L162
	}
L162:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v89)+80)) = v169
	F_errmsg(m, int32(_a_F_TidListEval_11), v89+int32(80))
	mBase = m.M
	v580 = m.ExcPending
	if v580 != 0 {
		goto L4
	} else {
		goto L163
	}
L163:
	;
	F_errfinish(m, int32(_a_F_TidListEval_1), int32(183), int32(_a_F_TidListEval_5))
	mBase = m.M
	v585 = m.ExcPending
	if v585 != 0 {
		goto L4
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
	if v61 <= v63 {
		goto L166
	} else {
		goto L167
	}
L166:
	;
	v591 = F_repalloc(m, v64, v61*int32(12))
	mBase = m.M
	v592 = m.ExcPending
	if v592 != 0 {
		goto L4
	} else {
		goto L169
	}
L167:
	;
	v595 = v61
	v596 = v64
	goto L168
L168:
	;
	v599 = v596 + v63*int32(6)
	v600 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v23)+16)))
	*(*uint16)(unsafe.Add(mBase, uint32(v599)+4)) = uint16(v600)
	v602 = *(*int32)(unsafe.Add(mBase, uint32(v23)+12))
	*(*int32)(unsafe.Add(mBase, uint32(v599))) = v602
	v767 = v595
	v769 = v63 + int32(1)
	v770 = v596
	goto L17
L169:
	;
	v595 = v61 << (uint(int32(1)) % 32)
	v596 = v591
	goto L168
L170:
	;
	v609 = int32(_a_F_TidListEval_13)
	v610 = *(*int32)(unsafe.Add(mBase, _c_F_TidListEval[0]))
	v612 = *(*int32)(unsafe.Add(mBase, uint32(v25)+20))
	*(*int32)(unsafe.Add(mBase, _c_F_TidListEval[0])) = v612
	v616 = *(*int32)(unsafe.Add(mBase, uint32(v81)+20))
	v617 = m.T0[v616].(func(*base.Module, int32, int32, int32) int32)(m, v81, v25, v23+int32(31))
	mBase = m.M
	v618 = m.ExcPending
	if v618 != 0 {
		goto L4
	} else {
		goto L173
	}
L171:
	;
	goto L172
L172:
	;
	v647 = int32(_a_F_TidListEval_13)
	v648 = *(*int32)(unsafe.Add(mBase, _c_F_TidListEval[0]))
	v650 = *(*int32)(unsafe.Add(mBase, uint32(v25)+20))
	*(*int32)(unsafe.Add(mBase, _c_F_TidListEval[0])) = v650
	v654 = *(*int32)(unsafe.Add(mBase, uint32(v81)+20))
	v655 = m.T0[v654].(func(*base.Module, int32, int32, int32) int32)(m, v81, v25, v23+int32(31))
	mBase = m.M
	v656 = m.ExcPending
	if v656 != 0 {
		goto L4
	} else {
		goto L181
	}
L173:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_TidListEval[0])) = v610
	v621 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v23)+31)))
	if v621 != 0 {
		v767 = v61
		v769 = v63
		v770 = v64
		goto L17
	} else {
		goto L174
	}
L174:
	;
	v622 = *(*int32)(unsafe.Add(mBase, uint32(v42)))
	v623 = *(*int32)(unsafe.Add(mBase, uint32(v622)+188))
	v624 = *(*int32)(unsafe.Add(mBase, uint32(v623)+64))
	v625 = m.T0[v624].(func(*base.Module, int32, int32) int32)(m, v42, v617)
	mBase = m.M
	v626 = m.ExcPending
	if v626 != 0 {
		goto L4
	} else {
		goto L175
	}
L175:
	;
	if v625 == int32(0) {
		v767 = v61
		v769 = v63
		v770 = v64
		goto L17
	} else {
		goto L176
	}
L176:
	;
	if v61 <= v63 {
		goto L177
	} else {
		goto L178
	}
L177:
	;
	v632 = F_repalloc(m, v64, v61*int32(12))
	mBase = m.M
	v633 = m.ExcPending
	if v633 != 0 {
		goto L4
	} else {
		goto L180
	}
L178:
	;
	v636 = v61
	v637 = v64
	goto L179
L179:
	;
	v640 = v637 + v63*int32(6)
	v641 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v617)+4)))
	*(*uint16)(unsafe.Add(mBase, uint32(v640)+4)) = uint16(v641)
	v643 = *(*int32)(unsafe.Add(mBase, uint32(v617)))
	*(*int32)(unsafe.Add(mBase, uint32(v640))) = v643
	v767 = v636
	v769 = v63 + int32(1)
	v770 = v637
	goto L17
L180:
	;
	v636 = v61 << (uint(int32(1)) % 32)
	v637 = v632
	goto L179
L181:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_TidListEval[0])) = v648
	v659 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v23)+31)))
	if v659 != 0 {
		v767 = v61
		v769 = v63
		v770 = v64
		goto L17
	} else {
		goto L182
	}
L182:
	;
	v660 = F_pg_detoast_datum(m, v655)
	mBase = m.M
	v661 = m.ExcPending
	if v661 != 0 {
		goto L4
	} else {
		goto L183
	}
L183:
	;
	F_deconstruct_array_builtin(m, v660, int32(27), v23+int32(12), v23+int32(24), v23+int32(20))
	mBase = m.M
	v670 = m.ExcPending
	if v670 != 0 {
		goto L4
	} else {
		goto L184
	}
L184:
	;
	v671 = *(*int32)(unsafe.Add(mBase, uint32(v23)+20))
	v672 = v671 + v63
	if v61 < v672 {
		goto L185
	} else {
		goto L186
	}
L185:
	;
	v676 = F_repalloc(m, v64, v672*int32(6))
	mBase = m.M
	v677 = m.ExcPending
	if v677 != 0 {
		goto L4
	} else {
		goto L188
	}
L186:
	;
	v679 = v671
	v680 = v61
	v681 = v64
	goto L187
L187:
	;
	v682 = int32(0)
	if v682 < v679 {
		goto L189
	} else {
		goto L190
	}
L188:
	;
	v678 = *(*int32)(unsafe.Add(mBase, uint32(v23)+20))
	v679 = v678
	v680 = v672
	v681 = v676
	goto L187
L189:
	;
	v687 = v682
	v692 = v63
	goto L192
L190:
	;
	v743 = v63
	goto L191
L191:
	;
	v756 = *(*int32)(unsafe.Add(mBase, uint32(v23)+12))
	F_pfree(m, v756)
	mBase = m.M
	v758 = m.ExcPending
	if v758 != 0 {
		goto L4
	} else {
		goto L199
	}
L192:
	;
	v705 = *(*int32)(unsafe.Add(mBase, uint32(v23)+24))
	v707 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v705+v687))))
	if v707 != 0 {
		v731 = v692
		goto L194
	} else {
		goto L195
	}
L193:
	;
	v743 = v731
	goto L191
L194:
	;
	v733 = v687 + int32(1)
	v734 = *(*int32)(unsafe.Add(mBase, uint32(v23)+20))
	if v733 < v734 {
		v687 = v733
		v692 = v731
		goto L192
	} else {
		goto L198
	}
L195:
	;
	v708 = *(*int32)(unsafe.Add(mBase, uint32(v23)+12))
	v712 = *(*int32)(unsafe.Add(mBase, uint32(v708+v687<<(uint(int32(2))%32))))
	v713 = *(*int32)(unsafe.Add(mBase, uint32(v42)))
	v714 = *(*int32)(unsafe.Add(mBase, uint32(v713)+188))
	v715 = *(*int32)(unsafe.Add(mBase, uint32(v714)+64))
	v716 = m.T0[v715].(func(*base.Module, int32, int32) int32)(m, v42, v712)
	mBase = m.M
	v717 = m.ExcPending
	if v717 != 0 {
		goto L4
	} else {
		goto L196
	}
L196:
	;
	if v716 == int32(0) {
		v731 = v692
		goto L194
	} else {
		goto L197
	}
L197:
	;
	v722 = v681 + v692*int32(6)
	v723 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v712)+4)))
	*(*uint16)(unsafe.Add(mBase, uint32(v722)+4)) = uint16(v723)
	v725 = *(*int32)(unsafe.Add(mBase, uint32(v712)))
	*(*int32)(unsafe.Add(mBase, uint32(v722))) = v725
	v731 = v692 + int32(1)
	goto L194
L198:
	;
	goto L193
L199:
	;
	v759 = *(*int32)(unsafe.Add(mBase, uint32(v23)+24))
	F_pfree(m, v759)
	mBase = m.M
	v761 = m.ExcPending
	if v761 != 0 {
		goto L4
	} else {
		goto L200
	}
L200:
	;
	v767 = v680
	v769 = v743
	v770 = v681
	goto L17
L201:
	;
	goto L16
L202:
	;
	F_pg_qsort(m, v794, v793, int32(6), int32(770))
	mBase = m.M
	v811 = m.ExcPending
	if v811 != 0 {
		goto L4
	} else {
		goto L203
	}
L203:
	;
	v816 = int32(1)
	v823 = int32(0)
	goto L204
L204:
	;
	v834 = int32(6)
	v836 = v794 + v816*v834
	v837 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v836))))
	v838 = int32(16)
	v840 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v836)+2)))
	v844 = v794 + v823*v834
	v845 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v844))))
	v848 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v844)+2)))
	if v837<<(uint(v838)%32)|v840 == v845<<(uint(v838)%32)|v848 {
		goto L207
	} else {
		goto L208
	}
L205:
	;
	v878 = v865 + int32(1)
	v879 = v794
	goto L10
L206:
	;
	v867 = v816 + int32(1)
	if v867 != v793 {
		v816 = v867
		v823 = v865
		goto L204
	} else {
		goto L212
	}
L207:
	;
	v851 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v836)+4)))
	v852 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v844)+4)))
	if v851 == v852 {
		v865 = v823
		goto L206
	} else {
		goto L210
	}
L208:
	;
	goto L209
L209:
	;
	v855 = v823 + int32(1)
	if v855 == v816 {
		v865 = v816
		goto L206
	} else {
		goto L211
	}
L210:
	;
	goto L209
L211:
	;
	v859 = v794 + v855*int32(6)
	v860 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v836)+4)))
	*(*uint16)(unsafe.Add(mBase, uint32(v859)+4)) = uint16(v860)
	v862 = *(*int32)(unsafe.Add(mBase, uint32(v836)))
	*(*int32)(unsafe.Add(mBase, uint32(v859))) = v862
	v865 = v855
	goto L206
L212:
	;
	goto L205
}
func F_TidStoreCreateLocal(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v7 int32
	_ = v7
	var v10 int32
	_ = v10
	var v13 int32
	_ = v13
	var v21 int32
	_ = v21
	var v23 int32
	_ = v23
	var v26 int32
	_ = v26
	var v27 int32
	_ = v27
	var v28 int32
	_ = v28
	var v31 int32
	_ = v31
	var v32 int32
	_ = v32
	var v34 int32
	_ = v34
	var v35 int32
	_ = v35
	var v40 int32
	_ = v40
	var v41 int32
	_ = v41
	var v46 int32
	_ = v46
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
	var v68 int32
	_ = v68
	var v70 int32
	_ = v70
	var v71 int32
	_ = v71
	var v74 int32
	_ = v74
	var v76 int32
	_ = v76
	var v79 int32
	_ = v79
	v7 = F_palloc0(m, int32(12))
	mBase = m.M
	v10 = m.ExcPending
	if v10 != 0 {
		return int32(0)
	} else {
		v13 = int32(_a_F_TidStoreCreateLocal_0)
		for {
			if base.Ui32(l0) < base.Ui32(v13<<(uint(int32(4))%32)) {
				v13 = int32(base.Ui32(v13) >> (uint(int32(1)) % 32))
				continue
			} else {
				break
			}
			break
		}
		v21 = *(*int32)(unsafe.Add(mBase, _c_F_TidStoreCreateLocal[0]))
		v23 = int32(_a_F_TidStoreCreateLocal_1)
		if base.Ui32(v13) <= base.Ui32(v23) {
			v26 = v23
		} else {
			v26 = v13
		}
		v27 = F_BumpContextCreate(m, v21, int32(_a_F_TidStoreCreateLocal_2), v26)
		mBase = m.M
		v28 = m.ExcPending
		if v28 != 0 {
			return int32(0)
		} else {
			*(*int32)(unsafe.Add(mBase, uint32(v7))) = v27
			v31 = F_palloc0(m, int32(28))
			mBase = m.M
			v32 = m.ExcPending
			if v32 != 0 {
				return int32(0)
			} else {
				v34 = F_palloc0(m, int32(32))
				mBase = m.M
				v35 = m.ExcPending
				if v35 != 0 {
					return int32(0)
				} else {
					*(*int32)(unsafe.Add(mBase, uint32(v31))) = v34
					v40 = F_SlabContextCreate(m, v27, int32(_a_F_TidStoreCreateLocal_3), int32(_a_F_TidStoreCreateLocal_1), int32(24))
					mBase = m.M
					v41 = m.ExcPending
					if v41 != 0 {
						return int32(0)
					} else {
						*(*int32)(unsafe.Add(mBase, uint32(v31)+4)) = v40
						v46 = F_SlabContextCreate(m, v27, int32(_a_F_TidStoreCreateLocal_4), int32(_a_F_TidStoreCreateLocal_1), int32(100))
						mBase = m.M
						v47 = m.ExcPending
						if v47 != 0 {
							return int32(0)
						} else {
							*(*int32)(unsafe.Add(mBase, uint32(v31)+8)) = v46
							v52 = F_SlabContextCreate(m, v27, int32(_a_F_TidStoreCreateLocal_5), int32(_a_F_TidStoreCreateLocal_1), int32(164))
							mBase = m.M
							v53 = m.ExcPending
							if v53 != 0 {
								return int32(0)
							} else {
								*(*int32)(unsafe.Add(mBase, uint32(v31)+12)) = v52
								v58 = F_SlabContextCreate(m, v27, int32(_a_F_TidStoreCreateLocal_6), int32(_a_F_TidStoreCreateLocal_7), int32(524))
								mBase = m.M
								v59 = m.ExcPending
								if v59 != 0 {
									return int32(0)
								} else {
									*(*int32)(unsafe.Add(mBase, uint32(v31)+16)) = v58
									v64 = F_SlabContextCreate(m, v27, int32(_a_F_TidStoreCreateLocal_8), int32(_a_F_TidStoreCreateLocal_9), int32(1060))
									mBase = m.M
									v65 = m.ExcPending
									if v65 != 0 {
										return int32(0)
									} else {
										*(*int32)(unsafe.Add(mBase, uint32(v31)+24)) = v27
										*(*int32)(unsafe.Add(mBase, uint32(v31)+20)) = v64
										v68 = *(*int32)(unsafe.Add(mBase, uint32(v31)+4))
										v70 = F_MemoryContextAlloc(m, v68, int32(24))
										mBase = m.M
										v71 = m.ExcPending
										if v71 != 0 {
											return int32(0)
										} else {
											*(*int64)(unsafe.Add(mBase, uint32(v70))) = int64(1024)
											v74 = *(*int32)(unsafe.Add(mBase, uint32(v31)))
											*(*int32)(unsafe.Add(mBase, uint32(v74))) = v70
											v76 = *(*int32)(unsafe.Add(mBase, uint32(v31)))
											*(*int32)(unsafe.Add(mBase, uint32(v76)+24)) = int32(0)
											v79 = *(*int32)(unsafe.Add(mBase, uint32(v31)))
											*(*int64)(unsafe.Add(mBase, uint32(v79)+8)) = int64(255)
											*(*int32)(unsafe.Add(mBase, uint32(v7)+4)) = v31
											return v7
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
