package p0

import (
	base "github.com/shibukawa/pgmem/internal/aot/pgaot/base"
	"unsafe"
)

func F_px_crypt_des(m *base.Module, l0 int32, l1 int32) int32 {
	mBase := m.M
	_ = mBase
	var v3 int32
	_ = v3
	var v19 int32
	_ = v19
	var v21 int32
	_ = v21
	var v24 int32
	_ = v24
	var v28 int32
	_ = v28
	var v29 int32
	_ = v29
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
	var v41 int32
	_ = v41
	var v42 int32
	_ = v42
	var v44 int32
	_ = v44
	var v48 int32
	_ = v48
	var v49 int32
	_ = v49
	var v51 int32
	_ = v51
	var v55 int32
	_ = v55
	var v56 int32
	_ = v56
	var v58 int32
	_ = v58
	var v62 int32
	_ = v62
	var v63 int32
	_ = v63
	var v65 int32
	_ = v65
	var v69 int32
	_ = v69
	var v70 int32
	_ = v70
	var v72 int32
	_ = v72
	var v76 int32
	_ = v76
	var v77 int32
	_ = v77
	var v79 int32
	_ = v79
	var v82 int32
	_ = v82
	var v83 int32
	_ = v83
	var v89 int32
	_ = v89
	var v104 int32
	_ = v104
	var v110 int32
	_ = v110
	var v112 int32
	_ = v112
	var v113 int32
	_ = v113
	var v123 int32
	_ = v123
	var v129 int32
	_ = v129
	var v135 int32
	_ = v135
	var v137 int32
	_ = v137
	var v148 int32
	_ = v148
	var v154 int32
	_ = v154
	var v160 int32
	_ = v160
	var v161 int32
	_ = v161
	var v171 int32
	_ = v171
	var v177 int32
	_ = v177
	var v183 int32
	_ = v183
	var v185 int32
	_ = v185
	var v200 int32
	_ = v200
	var v206 int32
	_ = v206
	var v208 int32
	_ = v208
	var v209 int32
	_ = v209
	var v219 int32
	_ = v219
	var v225 int32
	_ = v225
	var v231 int32
	_ = v231
	var v232 int32
	_ = v232
	var v235 int32
	_ = v235
	var v246 int32
	_ = v246
	var v252 int32
	_ = v252
	var v258 int32
	_ = v258
	var v259 int32
	_ = v259
	var v260 int32
	_ = v260
	var v270 int32
	_ = v270
	var v276 int32
	_ = v276
	var v282 int32
	_ = v282
	var v283 int32
	_ = v283
	var v284 int32
	_ = v284
	var v303 int32
	_ = v303
	var v308 int32
	_ = v308
	var v309 int32
	_ = v309
	var v319 int32
	_ = v319
	var v320 int32
	_ = v320
	var v322 int32
	_ = v322
	var v324 int32
	_ = v324
	var v335 int32
	_ = v335
	var v356 int32
	_ = v356
	var v359 int32
	_ = v359
	var v361 int32
	_ = v361
	var v362 int32
	_ = v362
	var v364 int32
	_ = v364
	var v366 int32
	_ = v366
	var v368 int32
	_ = v368
	var v376 int32
	_ = v376
	var v378 int32
	_ = v378
	var v385 int32
	_ = v385
	var v393 int32
	_ = v393
	var v395 int32
	_ = v395
	var v400 int32
	_ = v400
	var v402 int32
	_ = v402
	var v411 int32
	_ = v411
	var v413 int32
	_ = v413
	var v422 int32
	_ = v422
	var v424 int32
	_ = v424
	var v433 int32
	_ = v433
	var v435 int32
	_ = v435
	var v442 int32
	_ = v442
	var v444 int32
	_ = v444
	var v453 int32
	_ = v453
	var v455 int32
	_ = v455
	var v464 int32
	_ = v464
	var v466 int32
	_ = v466
	var v475 int32
	_ = v475
	var v479 int32
	_ = v479
	var v484 int32
	_ = v484
	var v503 int32
	_ = v503
	var v506 int32
	_ = v506
	var v509 int32
	_ = v509
	var v521 int32
	_ = v521
	var v528 int32
	_ = v528
	var v541 int32
	_ = v541
	var v542 int32
	_ = v542
	var v543 int32
	_ = v543
	var v545 int32
	_ = v545
	var v549 int32
	_ = v549
	var v550 int32
	_ = v550
	var v552 int32
	_ = v552
	var v553 int32
	_ = v553
	var v554 int32
	_ = v554
	var v556 int32
	_ = v556
	var v562 int32
	_ = v562
	var v563 int32
	_ = v563
	var v564 int32
	_ = v564
	var v565 int32
	_ = v565
	var v568 int32
	_ = v568
	var v574 int32
	_ = v574
	var v575 int32
	_ = v575
	var v576 int32
	_ = v576
	var v579 int32
	_ = v579
	var v582 int32
	_ = v582
	var v587 int32
	_ = v587
	var v588 int32
	_ = v588
	var v590 int32
	_ = v590
	var v592 int32
	_ = v592
	var v596 int32
	_ = v596
	var v597 int32
	_ = v597
	var v598 int32
	_ = v598
	var v603 int32
	_ = v603
	var v604 int32
	_ = v604
	var v605 int32
	_ = v605
	var v608 int32
	_ = v608
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
	var v619 int32
	_ = v619
	var v621 int32
	_ = v621
	var v623 int32
	_ = v623
	var v624 int32
	_ = v624
	var v627 int32
	_ = v627
	var v634 int32
	_ = v634
	var v637 int32
	_ = v637
	var v641 int32
	_ = v641
	var v643 int32
	_ = v643
	var v658 int32
	_ = v658
	var v664 int32
	_ = v664
	var v666 int32
	_ = v666
	var v682 int32
	_ = v682
	var v688 int32
	_ = v688
	var v690 int32
	_ = v690
	var v691 int32
	_ = v691
	var v692 int32
	_ = v692
	var v704 int32
	_ = v704
	var v708 int32
	_ = v708
	var v717 int32
	_ = v717
	var v719 int32
	_ = v719
	var v721 int32
	_ = v721
	var v733 int32
	_ = v733
	var v735 int32
	_ = v735
	var v737 int32
	_ = v737
	var v740 int32
	_ = v740
	var v752 int32
	_ = v752
	var v753 int32
	_ = v753
	var v754 int32
	_ = v754
	var v759 int32
	_ = v759
	var v760 int32
	_ = v760
	var v761 int32
	_ = v761
	var v766 int32
	_ = v766
	var v767 int32
	_ = v767
	var v772 int32
	_ = v772
	var v777 int32
	_ = v777
	var v798 int32
	_ = v798
	var v805 int32
	_ = v805
	var v806 int32
	_ = v806
	var v807 int32
	_ = v807
	var v810 int32
	_ = v810
	var v814 int32
	_ = v814
	var v818 int32
	_ = v818
	var v821 int32
	_ = v821
	var v828 int32
	_ = v828
	var v835 int32
	_ = v835
	var v837 int32
	_ = v837
	var v840 int32
	_ = v840
	var v845 int32
	_ = v845
	var v852 int32
	_ = v852
	var v859 int32
	_ = v859
	var v866 int32
	_ = v866
	var v873 int32
	_ = v873
	var v875 int32
	_ = v875
	var v879 int32
	_ = v879
	var v883 int32
	_ = v883
	var v890 int32
	_ = v890
	var v894 int32
	_ = v894
	var v919 int32
	_ = v919
	var v922 int32
	_ = v922
	var v928 int32
	_ = v928
	var v935 int32
	_ = v935
	var v939 int32
	_ = v939
	var v942 int32
	_ = v942
	var v948 int32
	_ = v948
	var v955 int32
	_ = v955
	v3 = int32(0)
	v19 = m.G0
	v21 = v19 - int32(16)
	m.G0 = v21
	v24 = int32(*(*uint8)(unsafe.Add(mBase, _consts[1089])))
	if v24 == v3 {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	F_des_init(m)
	mBase = m.M
	goto L3
L2:
	;
	goto L3
L3:
	;
	v28 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0))))
	v29 = int32(1)
	v30 = v28 << (uint(v29) % 32)
	*(*uint8)(unsafe.Add(mBase, uint32(v21))) = uint8(v30)
	v32 = int32(0)
	v34 = l0 + base.B2i32(v28 != v32)
	v35 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v34))))
	v37 = v35 << (uint(v29) % 32)
	*(*uint8)(unsafe.Add(mBase, uint32(v21)+1)) = uint8(v37)
	v41 = v34 + base.B2i32(v35 != v32)
	v42 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v41))))
	v44 = v42 << (uint(v29) % 32)
	*(*uint8)(unsafe.Add(mBase, uint32(v21)+2)) = uint8(v44)
	v48 = v41 + base.B2i32(v42 != v32)
	v49 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v48))))
	v51 = v49 << (uint(v29) % 32)
	*(*uint8)(unsafe.Add(mBase, uint32(v21)+3)) = uint8(v51)
	v55 = v48 + base.B2i32(v49 != v32)
	v56 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v55))))
	v58 = v56 << (uint(v29) % 32)
	*(*uint8)(unsafe.Add(mBase, uint32(v21)+4)) = uint8(v58)
	v62 = v55 + base.B2i32(v56 != v32)
	v63 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v62))))
	v65 = v63 << (uint(v29) % 32)
	*(*uint8)(unsafe.Add(mBase, uint32(v21)+5)) = uint8(v65)
	v69 = v62 + base.B2i32(v63 != v32)
	v70 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v69))))
	v72 = v70 << (uint(v29) % 32)
	*(*uint8)(unsafe.Add(mBase, uint32(v21)+6)) = uint8(v72)
	v76 = v69 + base.B2i32(v70 != v32)
	v77 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v76))))
	v79 = v77 << (uint(v29) % 32)
	*(*uint8)(unsafe.Add(mBase, uint32(v21)+7)) = uint8(v79)
	F_des_setkey(m, v21)
	mBase = m.M
	v82 = F_strlen(m, l1)
	mBase = m.M
	v83 = int32(*(*int8)(unsafe.Add(mBase, uint32(l1))))
	if v83 == int32(95) {
		goto L8
	} else {
		goto L9
	}
L4:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v939 = m.ExcPending
	if v939 != 0 {
		goto L111
	} else {
		goto L222
	}
L5:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v919 = m.ExcPending
	if v919 != 0 {
		goto L111
	} else {
		goto L218
	}
L6:
	;
	m.G0 = v21 + int32(16)
	return v894
L7:
	;
	v719 = *(*int32)(unsafe.Add(mBase, _consts[1090]))
	if v719 != v704 {
		goto L196
	} else {
		goto L197
	}
L8:
	;
	if base.Ui32(v82) < base.Ui32(int32(9)) {
		goto L5
	} else {
		goto L11
	}
L9:
	;
	goto L10
L10:
	;
	if base.Ui32(v82) <= base.Ui32(int32(1)) {
		goto L4
	} else {
		goto L172
	}
L11:
	;
	v89 = int32(*(*int8)(unsafe.Add(mBase, uint32(l1)+1)))
	if int32(122) < v89 {
		v112 = int32(0)
		goto L12
	} else {
		goto L13
	}
L12:
	;
	v113 = int32(*(*int8)(unsafe.Add(mBase, uint32(l1)+2)))
	if int32(122) < v113 {
		v135 = v3
		goto L20
	} else {
		goto L21
	}
L13:
	;
	if int32(97) <= v89 {
		v112 = v89 - int32(59)
		goto L12
	} else {
		goto L14
	}
L14:
	;
	if int32(90) < v89 {
		v112 = int32(0)
		goto L12
	} else {
		goto L15
	}
L15:
	;
	if int32(65) <= v89 {
		v112 = v89 - int32(53)
		goto L12
	} else {
		goto L16
	}
L16:
	;
	v104 = v89 - int32(46)
	if base.Ui32(v104&int32(255)) < base.Ui32(int32(12)) {
		goto L17
	} else {
		goto L18
	}
L17:
	;
	v110 = v104
	goto L19
L18:
	;
	v110 = int32(0)
	goto L19
L19:
	;
	v112 = v110
	goto L12
L20:
	;
	v137 = int32(*(*int8)(unsafe.Add(mBase, uint32(l1)+3)))
	if int32(122) < v137 {
		v160 = int32(0)
		goto L32
	} else {
		goto L33
	}
L21:
	;
	if v113 <= int32(96) {
		goto L22
	} else {
		goto L23
	}
L22:
	;
	if int32(90) < v113 {
		v135 = v3
		goto L20
	} else {
		goto L25
	}
L23:
	;
	goto L24
L24:
	;
	v135 = v113 - int32(59)
	goto L20
L25:
	;
	if v113 <= int32(64) {
		goto L26
	} else {
		goto L27
	}
L26:
	;
	v123 = v113 - int32(46)
	if base.Ui32(v123&int32(255)) < base.Ui32(int32(12)) {
		goto L29
	} else {
		goto L30
	}
L27:
	;
	goto L28
L28:
	;
	v135 = v113 - int32(53)
	goto L20
L29:
	;
	v129 = v123
	goto L31
L30:
	;
	v129 = int32(0)
	goto L31
L31:
	;
	v135 = v129
	goto L20
L32:
	;
	v161 = int32(*(*int8)(unsafe.Add(mBase, uint32(l1)+4)))
	if int32(122) < v161 {
		v183 = v3
		goto L44
	} else {
		goto L45
	}
L33:
	;
	if v137 <= int32(96) {
		goto L34
	} else {
		goto L35
	}
L34:
	;
	if int32(90) < v137 {
		v160 = int32(0)
		goto L32
	} else {
		goto L37
	}
L35:
	;
	goto L36
L36:
	;
	v160 = v137 - int32(59)
	goto L32
L37:
	;
	if v137 <= int32(64) {
		goto L38
	} else {
		goto L39
	}
L38:
	;
	v148 = v137 - int32(46)
	if base.Ui32(v148&int32(255)) < base.Ui32(int32(12)) {
		goto L41
	} else {
		goto L42
	}
L39:
	;
	goto L40
L40:
	;
	v160 = v137 - int32(53)
	goto L32
L41:
	;
	v154 = v148
	goto L43
L42:
	;
	v154 = int32(0)
	goto L43
L43:
	;
	v160 = v154
	goto L32
L44:
	;
	v185 = int32(*(*int8)(unsafe.Add(mBase, uint32(l1)+5)))
	if int32(122) < v185 {
		v208 = int32(0)
		goto L56
	} else {
		goto L57
	}
L45:
	;
	if v161 <= int32(96) {
		goto L46
	} else {
		goto L47
	}
L46:
	;
	if int32(90) < v161 {
		v183 = v3
		goto L44
	} else {
		goto L49
	}
L47:
	;
	goto L48
L48:
	;
	v183 = v161 - int32(59)
	goto L44
L49:
	;
	if v161 <= int32(64) {
		goto L50
	} else {
		goto L51
	}
L50:
	;
	v171 = v161 - int32(46)
	if base.Ui32(v171&int32(255)) < base.Ui32(int32(12)) {
		goto L53
	} else {
		goto L54
	}
L51:
	;
	goto L52
L52:
	;
	v183 = v161 - int32(53)
	goto L44
L53:
	;
	v177 = v171
	goto L55
L54:
	;
	v177 = int32(0)
	goto L55
L55:
	;
	v183 = v177
	goto L44
L56:
	;
	v209 = int32(*(*int8)(unsafe.Add(mBase, uint32(l1)+6)))
	if int32(122) < v209 {
		v231 = v3
		goto L64
	} else {
		goto L65
	}
L57:
	;
	if int32(97) <= v185 {
		v208 = v185 - int32(59)
		goto L56
	} else {
		goto L58
	}
L58:
	;
	if int32(90) < v185 {
		v208 = int32(0)
		goto L56
	} else {
		goto L59
	}
L59:
	;
	if int32(65) <= v185 {
		v208 = v185 - int32(53)
		goto L56
	} else {
		goto L60
	}
L60:
	;
	v200 = v185 - int32(46)
	if base.Ui32(v200&int32(255)) < base.Ui32(int32(12)) {
		goto L61
	} else {
		goto L62
	}
L61:
	;
	v206 = v200
	goto L63
L62:
	;
	v206 = int32(0)
	goto L63
L63:
	;
	v208 = v206
	goto L56
L64:
	;
	v232 = int32(0)
	v235 = int32(*(*int8)(unsafe.Add(mBase, uint32(l1)+7)))
	if int32(122) < v235 {
		v258 = v232
		goto L76
	} else {
		goto L77
	}
L65:
	;
	if v209 <= int32(96) {
		goto L66
	} else {
		goto L67
	}
L66:
	;
	if int32(90) < v209 {
		v231 = v3
		goto L64
	} else {
		goto L69
	}
L67:
	;
	goto L68
L68:
	;
	v231 = v209 - int32(59)
	goto L64
L69:
	;
	if v209 <= int32(64) {
		goto L70
	} else {
		goto L71
	}
L70:
	;
	v219 = v209 - int32(46)
	if base.Ui32(v219&int32(255)) < base.Ui32(int32(12)) {
		goto L73
	} else {
		goto L74
	}
L71:
	;
	goto L72
L72:
	;
	v231 = v209 - int32(53)
	goto L64
L73:
	;
	v225 = v219
	goto L75
L74:
	;
	v225 = int32(0)
	goto L75
L75:
	;
	v231 = v225
	goto L64
L76:
	;
	v259 = base.B2i32(v77 != v232) + v76
	v260 = int32(*(*int8)(unsafe.Add(mBase, uint32(l1)+8)))
	if int32(122) < v260 {
		v282 = v3
		goto L88
	} else {
		goto L89
	}
L77:
	;
	if v235 <= int32(96) {
		goto L78
	} else {
		goto L79
	}
L78:
	;
	if int32(90) < v235 {
		v258 = int32(0)
		goto L76
	} else {
		goto L81
	}
L79:
	;
	goto L80
L80:
	;
	v258 = v235 - int32(59)
	goto L76
L81:
	;
	if v235 <= int32(64) {
		goto L82
	} else {
		goto L83
	}
L82:
	;
	v246 = v235 - int32(46)
	if base.Ui32(v246&int32(255)) < base.Ui32(int32(12)) {
		goto L85
	} else {
		goto L86
	}
L83:
	;
	goto L84
L84:
	;
	v258 = v235 - int32(53)
	goto L76
L85:
	;
	v252 = v246
	goto L87
L86:
	;
	v252 = int32(0)
	goto L87
L87:
	;
	v258 = v252
	goto L76
L88:
	;
	v283 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v259))))
	if v283 != 0 {
		goto L100
	} else {
		goto L101
	}
L89:
	;
	if v260 <= int32(96) {
		goto L90
	} else {
		goto L91
	}
L90:
	;
	if int32(90) < v260 {
		v282 = v3
		goto L88
	} else {
		goto L93
	}
L91:
	;
	goto L92
L92:
	;
	v282 = v260 - int32(59)
	goto L88
L93:
	;
	if v260 <= int32(64) {
		goto L94
	} else {
		goto L95
	}
L94:
	;
	v270 = v260 - int32(46)
	if base.Ui32(v270&int32(255)) < base.Ui32(int32(12)) {
		goto L97
	} else {
		goto L98
	}
L95:
	;
	goto L96
L96:
	;
	v282 = v260 - int32(53)
	goto L88
L97:
	;
	v276 = v270
	goto L99
L98:
	;
	v276 = int32(0)
	goto L99
L99:
	;
	v282 = v276
	goto L88
L100:
	;
	v284 = v259
	goto L103
L101:
	;
	goto L102
L102:
	;
	v503 = int32(6)
	v506 = int32(12)
	v509 = int32(18)
	v521 = int32(4608528)
	goto L143
L103:
	;
	v303 = int32(*(*uint8)(unsafe.Add(mBase, _consts[1089])))
	if v303 == int32(0) {
		goto L105
	} else {
		goto L106
	}
L104:
	;
	goto L102
L105:
	;
	F_des_init(m)
	mBase = m.M
	goto L107
L106:
	;
	goto L107
L107:
	;
	v308 = *(*int32)(unsafe.Add(mBase, _consts[1090]))
	if v308 != 0 {
		goto L108
	} else {
		goto L109
	}
L108:
	;
	v309 = int32(0)
	*(*int32)(unsafe.Add(mBase, _consts[1091])) = v309
	*(*int32)(unsafe.Add(mBase, _consts[1090])) = v309
	goto L110
L109:
	;
	goto L110
L110:
	;
	v319 = *(*int32)(unsafe.Add(mBase, uint32(v21)))
	v320 = int32(24)
	v322 = int32(65280)
	v324 = int32(8)
	v335 = *(*int32)(unsafe.Add(mBase, uint32(v21)+4))
	v356 = F_do_des(m, v319<<(uint(v320)%32)|v319&v322<<(uint(v324)%32)|(int32(base.Ui32(v319)>>(uint(v324)%32))&v322|int32(base.Ui32(v319)>>(uint(v320)%32))), v335<<(uint(v320)%32)|v335&v322<<(uint(v324)%32)|(int32(base.Ui32(v335)>>(uint(v324)%32))&v322|int32(base.Ui32(v335)>>(uint(v320)%32))), v21+int32(12), v21+v324, int32(1))
	mBase = m.M
	v359 = m.ExcPending
	if v359 != 0 {
		goto L111
	} else {
		goto L112
	}
L111:
	;
	return int32(0)
L112:
	;
	if v356 != 0 {
		goto L113
	} else {
		goto L114
	}
L113:
	;
	v894 = int32(0)
	goto L6
L114:
	;
	goto L115
L115:
	;
	v361 = *(*int32)(unsafe.Add(mBase, uint32(v21)+8))
	v362 = int32(24)
	v364 = int32(65280)
	v366 = int32(8)
	v368 = v361<<(uint(v362)%32) | v361&v364<<(uint(v366)%32)
	v376 = v368 | (int32(base.Ui32(v361)>>(uint(v366)%32))&v364 | int32(base.Ui32(v361)>>(uint(v362)%32)))
	*(*int32)(unsafe.Add(mBase, uint32(v21)+4)) = v376
	v378 = *(*int32)(unsafe.Add(mBase, uint32(v21)+12))
	v385 = v378<<(uint(v362)%32) | v378&v364<<(uint(v366)%32)
	v393 = v385 | (int32(base.Ui32(v378)>>(uint(v366)%32))&v364 | int32(base.Ui32(v378)>>(uint(v362)%32)))
	*(*int32)(unsafe.Add(mBase, uint32(v21))) = v393
	v395 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v284))))
	if v395 == int32(0) {
		v479 = v284
		goto L116
	} else {
		goto L117
	}
L116:
	;
	F_des_setkey(m, v21)
	mBase = m.M
	v484 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v479))))
	if v484 != 0 {
		v284 = v479
		goto L103
	} else {
		goto L139
	}
L117:
	;
	v400 = v395<<(uint(int32(1))%32) ^ v393
	*(*uint8)(unsafe.Add(mBase, uint32(v21))) = uint8(v400)
	v402 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v284)+1)))
	if v402 == int32(0) {
		goto L118
	} else {
		goto L119
	}
L118:
	;
	v479 = v284 + int32(1)
	goto L116
L119:
	;
	goto L120
L120:
	;
	v411 = v402<<(uint(int32(1))%32) ^ int32(base.Ui32(v393)>>(uint(int32(8))%32))
	*(*uint8)(unsafe.Add(mBase, uint32(v21)+1)) = uint8(v411)
	v413 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v284)+2)))
	if v413 == int32(0) {
		goto L121
	} else {
		goto L122
	}
L121:
	;
	v479 = v284 + int32(2)
	goto L116
L122:
	;
	goto L123
L123:
	;
	v422 = v413<<(uint(int32(1))%32) ^ int32(base.Ui32(v385)>>(uint(int32(16))%32))
	*(*uint8)(unsafe.Add(mBase, uint32(v21)+2)) = uint8(v422)
	v424 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v284)+3)))
	if v424 == int32(0) {
		goto L124
	} else {
		goto L125
	}
L124:
	;
	v479 = v284 + int32(3)
	goto L116
L125:
	;
	goto L126
L126:
	;
	v433 = v378&int32(255) ^ v424<<(uint(int32(1))%32)
	*(*uint8)(unsafe.Add(mBase, uint32(v21)+3)) = uint8(v433)
	v435 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v284)+4)))
	if v435 == int32(0) {
		goto L127
	} else {
		goto L128
	}
L127:
	;
	v479 = v284 + int32(4)
	goto L116
L128:
	;
	goto L129
L129:
	;
	v442 = v435<<(uint(int32(1))%32) ^ v376
	*(*uint8)(unsafe.Add(mBase, uint32(v21)+4)) = uint8(v442)
	v444 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v284)+5)))
	if v444 == int32(0) {
		goto L130
	} else {
		goto L131
	}
L130:
	;
	v479 = v284 + int32(5)
	goto L116
L131:
	;
	goto L132
L132:
	;
	v453 = v444<<(uint(int32(1))%32) ^ int32(base.Ui32(v376)>>(uint(int32(8))%32))
	*(*uint8)(unsafe.Add(mBase, uint32(v21)+5)) = uint8(v453)
	v455 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v284)+6)))
	if v455 == int32(0) {
		goto L133
	} else {
		goto L134
	}
L133:
	;
	v479 = v284 + int32(6)
	goto L116
L134:
	;
	goto L135
L135:
	;
	v464 = v455<<(uint(int32(1))%32) ^ int32(base.Ui32(v368)>>(uint(int32(16))%32))
	*(*uint8)(unsafe.Add(mBase, uint32(v21)+6)) = uint8(v464)
	v466 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v284)+7)))
	if v466 == int32(0) {
		goto L136
	} else {
		goto L137
	}
L136:
	;
	v479 = v284 + int32(7)
	goto L116
L137:
	;
	goto L138
L138:
	;
	v475 = v361&int32(255) ^ v466<<(uint(int32(1))%32)
	*(*uint8)(unsafe.Add(mBase, uint32(v21)+7)) = uint8(v475)
	v479 = v284 + int32(8)
	goto L116
L139:
	;
	goto L104
L140:
	;
	v637 = F_strlen(m, v521)
	mBase = m.M
	v704 = v231<<(uint(v503)%32) | v208 | v258<<(uint(v506)%32) | v282<<(uint(v509)%32)
	v708 = v135<<(uint(v503)%32) | v112 | v160<<(uint(v506)%32) | v183<<(uint(v509)%32)
	v717 = v637 + v521
	goto L7
L141:
	;
	v634 = F_strlen(m, v623)
	mBase = m.M
	goto L140
L143:
	;
	goto L144
L144:
	;
	v528 = int32(9)
	if (v521^l1)&int32(3) != 0 {
		goto L148
	} else {
		goto L149
	}
L145:
	;
	v627 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v624))) = uint8(v627)
	goto L141
L146:
	;
	v608 = v603
	v609 = v604
	v610 = v605
	goto L168
L147:
	;
	if v598 == int32(0) {
		v623 = v596
		v624 = v597
		goto L145
	} else {
		goto L167
	}
L148:
	;
	v596 = l1
	v597 = v521
	v598 = v528
	goto L147
L149:
	;
	goto L150
L150:
	;
	if l1&int32(3) == int32(0) {
		goto L152
	} else {
		goto L153
	}
L151:
	;
	if v565 == int32(0) {
		v623 = v562
		v624 = v563
		goto L145
	} else {
		goto L160
	}
L152:
	;
	v562 = l1
	v563 = v521
	v564 = v528
	v565 = int32(1)
	goto L151
L153:
	;
	goto L154
L154:
	;
	v541 = l1
	v542 = v521
	v543 = v528
	goto L155
L155:
	;
	v545 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v541))))
	*(*uint8)(unsafe.Add(mBase, uint32(v542))) = uint8(v545)
	if v545 == int32(0) {
		v603 = v541
		v604 = v542
		v605 = v543
		goto L146
	} else {
		goto L157
	}
L156:
	;
	v562 = v556
	v563 = v550
	v564 = v552
	v565 = v554
	goto L151
L157:
	;
	v549 = int32(1)
	v550 = v542 + v549
	v552 = v543 - v549
	v553 = int32(0)
	v554 = base.B2i32(v552 != v553)
	v556 = v541 + v549
	if v556&int32(3) == v553 {
		v562 = v556
		v563 = v550
		v564 = v552
		v565 = v554
		goto L151
	} else {
		goto L158
	}
L158:
	;
	if v552 != 0 {
		v541 = v556
		v542 = v550
		v543 = v552
		goto L155
	} else {
		goto L159
	}
L159:
	;
	goto L156
L160:
	;
	v568 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v562))))
	if v568 == int32(0) {
		v596 = v562
		v597 = v563
		v598 = v564
		goto L147
	} else {
		goto L161
	}
L161:
	;
	if base.Ui32(v564) < base.Ui32(int32(4)) {
		v596 = v562
		v597 = v563
		v598 = v564
		goto L147
	} else {
		goto L162
	}
L162:
	;
	v574 = v562
	v575 = v563
	v576 = v564
	goto L163
L163:
	;
	v579 = *(*int32)(unsafe.Add(mBase, uint32(v574)))
	v582 = int32(-2139062144)
	if (int32(16843008)-v579|v579)&v582 != v582 {
		v603 = v574
		v604 = v575
		v605 = v576
		goto L146
	} else {
		goto L165
	}
L164:
	;
	v596 = v590
	v597 = v588
	v598 = v592
	goto L147
L165:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v575))) = v579
	v587 = int32(4)
	v588 = v575 + v587
	v590 = v574 + v587
	v592 = v576 - v587
	if base.Ui32(int32(3)) < base.Ui32(v592) {
		v574 = v590
		v575 = v588
		v576 = v592
		goto L163
	} else {
		goto L166
	}
L166:
	;
	goto L164
L167:
	;
	v603 = v596
	v604 = v597
	v605 = v598
	goto L146
L168:
	;
	v612 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v608))))
	*(*uint8)(unsafe.Add(mBase, uint32(v609))) = uint8(v612)
	if v612 == int32(0) {
		v623 = v608
		v624 = v609
		goto L145
	} else {
		goto L170
	}
L169:
	;
	v623 = v619
	v624 = v617
	goto L145
L170:
	;
	v616 = int32(1)
	v617 = v609 + v616
	v619 = v608 + v616
	v621 = v610 - v616
	if v621 != 0 {
		v608 = v619
		v609 = v617
		v610 = v621
		goto L168
	} else {
		goto L171
	}
L171:
	;
	goto L169
L172:
	;
	v641 = int32(0)
	v643 = int32(*(*int8)(unsafe.Add(mBase, uint32(l1)+1)))
	if int32(122) < v643 {
		v666 = v641
		goto L173
	} else {
		goto L174
	}
L173:
	;
	if int32(122) < v83 {
		v690 = v641
		goto L181
	} else {
		goto L182
	}
L174:
	;
	if int32(97) <= v643 {
		v666 = v643 - int32(59)
		goto L173
	} else {
		goto L175
	}
L175:
	;
	if int32(90) < v643 {
		v666 = int32(0)
		goto L173
	} else {
		goto L176
	}
L176:
	;
	if int32(65) <= v643 {
		v666 = v643 - int32(53)
		goto L173
	} else {
		goto L177
	}
L177:
	;
	v658 = v643 - int32(46)
	if base.Ui32(v658&int32(255)) < base.Ui32(int32(12)) {
		goto L178
	} else {
		goto L179
	}
L178:
	;
	v664 = v658
	goto L180
L179:
	;
	v664 = int32(0)
	goto L180
L180:
	;
	v666 = v664
	goto L173
L181:
	;
	v691 = int32(4608528)
	if v643 != 0 {
		goto L193
	} else {
		goto L194
	}
L182:
	;
	if int32(97) <= v83 {
		goto L183
	} else {
		goto L184
	}
L183:
	;
	v690 = v83 - int32(59)
	goto L181
L184:
	;
	goto L185
L185:
	;
	if int32(90) < v83 {
		v690 = v641
		goto L181
	} else {
		goto L186
	}
L186:
	;
	if int32(65) <= v83 {
		goto L187
	} else {
		goto L188
	}
L187:
	;
	v690 = v83 - int32(53)
	goto L181
L188:
	;
	goto L189
L189:
	;
	v682 = v83 - int32(46)
	if base.Ui32(v682&int32(255)) < base.Ui32(int32(12)) {
		goto L190
	} else {
		goto L191
	}
L190:
	;
	v688 = v682
	goto L192
L191:
	;
	v688 = int32(0)
	goto L192
L192:
	;
	v690 = v688
	goto L181
L193:
	;
	v692 = v643
	goto L195
L194:
	;
	v692 = v83
	goto L195
L195:
	;
	*(*uint8)(unsafe.Add(mBase, _consts[1092])) = uint8(v692)
	*(*uint8)(unsafe.Add(mBase, _consts[1093])) = uint8(v83)
	v704 = v690 + v666<<(uint(int32(6))%32)
	v708 = int32(25)
	v717 = int32(4608530)
	goto L7
L196:
	;
	v721 = int32(0)
	*(*int32)(unsafe.Add(mBase, _consts[1091])) = v721
	*(*int32)(unsafe.Add(mBase, _consts[1090])) = v704
	v733 = int32(8388608)
	v735 = int32(1)
	v737 = v721
	v740 = v721
	goto L199
L197:
	;
	goto L198
L198:
	;
	v798 = int32(0)
	v805 = F_do_des(m, v798, v798, v21+int32(12), v21+int32(8), v708)
	mBase = m.M
	v806 = m.ExcPending
	if v806 != 0 {
		goto L111
	} else {
		goto L216
	}
L199:
	;
	v752 = v735 & v704
	if v752 != 0 {
		goto L201
	} else {
		goto L202
	}
L200:
	;
	goto L198
L201:
	;
	v753 = v733 | v737
	goto L203
L202:
	;
	v753 = v737
	goto L203
L203:
	;
	v754 = int32(1)
	v759 = v735 << (uint(v754) % 32) & v704
	if v759 != 0 {
		goto L204
	} else {
		goto L205
	}
L204:
	;
	v760 = v753 | int32(base.Ui32(v733)>>(uint(v754)%32))
	goto L206
L205:
	;
	v760 = v753
	goto L206
L206:
	;
	v761 = int32(2)
	v766 = v735 << (uint(v761) % 32) & v704
	if v766 != 0 {
		goto L207
	} else {
		goto L208
	}
L207:
	;
	v767 = v760 | int32(base.Ui32(v733)>>(uint(v761)%32))
	goto L209
L208:
	;
	v767 = v760
	goto L209
L209:
	;
	if v752 != 0 {
		goto L211
	} else {
		goto L212
	}
L210:
	;
	v772 = int32(3)
	v777 = v740 + v772
	if v777 != int32(24) {
		v733 = int32(base.Ui32(v733) >> (uint(v772) % 32))
		v735 = v735 << (uint(v772) % 32)
		v737 = v767
		v740 = v777
		goto L199
	} else {
		goto L215
	}
L211:
	;
	*(*int32)(unsafe.Add(mBase, _consts[1091])) = v767
	goto L210
L212:
	;
	if v759 != 0 {
		goto L211
	} else {
		goto L213
	}
L213:
	;
	if v766 == int32(0) {
		goto L210
	} else {
		goto L214
	}
L214:
	;
	goto L211
L215:
	;
	goto L200
L216:
	;
	if v805 != 0 {
		v894 = v798
		goto L6
	} else {
		goto L217
	}
L217:
	;
	v807 = int32(0)
	v810 = *(*int32)(unsafe.Add(mBase, uint32(v21)+12))
	v814 = int32(*(*uint8)(unsafe.Add(mBase, uint32(int32(base.Ui32(v810)>>(uint(int32(26))%32)))+uint32(_consts[1094]))))
	*(*uint8)(unsafe.Add(mBase, uint32(v717))) = uint8(v814)
	v818 = int32(63)
	v821 = int32(*(*uint8)(unsafe.Add(mBase, uint32(int32(base.Ui32(v810)>>(uint(int32(8))%32))&v818)+uint32(_consts[1094]))))
	*(*uint8)(unsafe.Add(mBase, uint32(v717)+3)) = uint8(v821)
	v828 = int32(*(*uint8)(unsafe.Add(mBase, uint32(int32(base.Ui32(v810)>>(uint(int32(14))%32))&v818)+uint32(_consts[1094]))))
	*(*uint8)(unsafe.Add(mBase, uint32(v717)+2)) = uint8(v828)
	v835 = int32(*(*uint8)(unsafe.Add(mBase, uint32(int32(base.Ui32(v810)>>(uint(int32(20))%32))&v818)+uint32(_consts[1094]))))
	*(*uint8)(unsafe.Add(mBase, uint32(v717)+1)) = uint8(v835)
	v837 = *(*int32)(unsafe.Add(mBase, uint32(v21)+8))
	*(*uint8)(unsafe.Add(mBase, uint32(v717)+11)) = uint8(v807)
	v840 = int32(2)
	v845 = int32(*(*uint8)(unsafe.Add(mBase, uint32(int32(base.Ui32(v810)>>(uint(v840)%32))&v818)+uint32(_consts[1094]))))
	*(*uint8)(unsafe.Add(mBase, uint32(v717)+4)) = uint8(v845)
	v852 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v837<<(uint(v840)%32)&int32(60))+uint32(_consts[1094]))))
	*(*uint8)(unsafe.Add(mBase, uint32(v717)+10)) = uint8(v852)
	v859 = int32(*(*uint8)(unsafe.Add(mBase, uint32(int32(base.Ui32(v837)>>(uint(int32(4))%32))&v818)+uint32(_consts[1094]))))
	*(*uint8)(unsafe.Add(mBase, uint32(v717)+9)) = uint8(v859)
	v866 = int32(*(*uint8)(unsafe.Add(mBase, uint32(int32(base.Ui32(v837)>>(uint(int32(10))%32))&v818)+uint32(_consts[1094]))))
	*(*uint8)(unsafe.Add(mBase, uint32(v717)+8)) = uint8(v866)
	v873 = int32(*(*uint8)(unsafe.Add(mBase, uint32(int32(base.Ui32(v837)>>(uint(int32(22))%32))&v818)+uint32(_consts[1094]))))
	*(*uint8)(unsafe.Add(mBase, uint32(v717)+6)) = uint8(v873)
	v875 = int32(16)
	v879 = v810<<(uint(v875)%32) | int32(base.Ui32(v837)>>(uint(v875)%32))
	v883 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v879&v818)+uint32(_consts[1094]))))
	*(*uint8)(unsafe.Add(mBase, uint32(v717)+7)) = uint8(v883)
	v890 = int32(*(*uint8)(unsafe.Add(mBase, uint32(int32(base.Ui32(v879)>>(uint(int32(12))%32))&v818)+uint32(_consts[1094]))))
	*(*uint8)(unsafe.Add(mBase, uint32(v717)+5)) = uint8(v890)
	v894 = int32(4608528)
	goto L6
L218:
	;
	F_errcode(m, int32(50856066))
	mBase = m.M
	v922 = m.ExcPending
	if v922 != 0 {
		goto L111
	} else {
		goto L219
	}
L219:
	;
	F_errmsg(m, int32(98760), int32(0))
	mBase = m.M
	v928 = m.ExcPending
	if v928 != 0 {
		goto L111
	} else {
		goto L220
	}
L220:
	;
	F_errfinish(m, int32(494044), int32(697), int32(171254))
	mBase = m.M
	v935 = m.ExcPending
	if v935 != 0 {
		goto L111
	} else {
		goto L221
	}
L221:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L222:
	;
	F_errcode(m, int32(50856066))
	mBase = m.M
	v942 = m.ExcPending
	if v942 != 0 {
		goto L111
	} else {
		goto L223
	}
L223:
	;
	F_errmsg(m, int32(98760), int32(0))
	mBase = m.M
	v948 = m.ExcPending
	if v948 != 0 {
		goto L111
	} else {
		goto L224
	}
L224:
	;
	F_errfinish(m, int32(494044), int32(745), int32(171254))
	mBase = m.M
	v955 = m.ExcPending
	if v955 != 0 {
		goto L111
	} else {
		goto L225
	}
L225:
	;
	base.Wasm_trap_unreachable()
	for {
	}
}
func F_px_debug(m *base.Module, l0 int32, l1 int32) {
	mBase := m.M
	_ = mBase
	var v4 int32
	_ = v4
	var v6 int32
	_ = v6
	var v10 int32
	_ = v10
	var v12 int32
	_ = v12
	var v13 int32
	_ = v13
	var v14 int32
	_ = v14
	var v16 int32
	_ = v16
	var v18 int32
	_ = v18
	v4 = m.G0
	v6 = v4 - int32(528)
	m.G0 = v6
	*(*int32)(unsafe.Add(mBase, uint32(v6)+524)) = l1
	v10 = *(*int32)(unsafe.Add(mBase, _consts[1097]))
	if v10 != 0 {
		v12 = *(*int32)(unsafe.Add(mBase, uint32(v6)+524))
		v13 = F_pg_vsnprintf(m, v6, int32(512), l0, v12)
		mBase = m.M
		v14 = m.ExcPending
		if v14 != 0 {
			return
		} else {
			v16 = *(*int32)(unsafe.Add(mBase, _consts[1097]))
			m.T0[v16].(func(*base.Module, int32))(m, v6)
			mBase = m.M
			v18 = m.ExcPending
			if v18 != 0 {
				return
			} else {
				m.G0 = v6 + int32(528)
				return
			}
		}
	} else {
		m.G0 = v6 + int32(528)
		return
	}
}
func F_px_find_hmac(m *base.Module, l0 int32, l1 int32) int32 {
	mBase := m.M
	_ = mBase
	var v6 int32
	_ = v6
	var v8 int32
	_ = v8
	var v12 int32
	_ = v12
	var v15 int32
	_ = v15
	var v16 int32
	_ = v16
	var v17 int32
	_ = v17
	var v18 int32
	_ = v18
	var v19 int32
	_ = v19
	var v22 int32
	_ = v22
	var v23 int32
	_ = v23
	var v25 int32
	_ = v25
	var v28 int32
	_ = v28
	var v29 int32
	_ = v29
	var v30 int32
	_ = v30
	var v31 int32
	_ = v31
	var v33 int32
	_ = v33
	var v34 int32
	_ = v34
	var v36 int32
	_ = v36
	var v62 int32
	_ = v62
	v6 = m.G0
	v8 = v6 - int32(16)
	m.G0 = v8
	v12 = F_px_find_digest(m, l0, v8+int32(12))
	mBase = m.M
	v15 = m.ExcPending
	if v15 != 0 {
		return int32(0)
	} else {
		if v12 != 0 {
			v62 = v12
			m.G0 = v8 + int32(16)
			return v62
		} else {
			v16 = *(*int32)(unsafe.Add(mBase, uint32(v8)+12))
			v17 = *(*int32)(unsafe.Add(mBase, uint32(v16)+4))
			v18 = m.T0[v17].(func(*base.Module, int32) int32)(m, v16)
			mBase = m.M
			v19 = m.ExcPending
			if v19 != 0 {
				return int32(0)
			} else {
				if base.Ui32(v18) <= base.Ui32(int32(1)) {
					v22 = *(*int32)(unsafe.Add(mBase, uint32(v8)+12))
					v23 = *(*int32)(unsafe.Add(mBase, uint32(v22)+20))
					m.T0[v23].(func(*base.Module, int32))(m, v22)
					mBase = m.M
					v25 = m.ExcPending
					if v25 != 0 {
						return int32(0)
					} else {
						v62 = int32(-9)
						m.G0 = v8 + int32(16)
						return v62
					}
				} else {
					v28 = F_palloc(m, int32(40))
					mBase = m.M
					v29 = m.ExcPending
					if v29 != 0 {
						return int32(0)
					} else {
						v30 = F_palloc(m, v18)
						mBase = m.M
						v31 = m.ExcPending
						if v31 != 0 {
							return int32(0)
						} else {
							*(*int32)(unsafe.Add(mBase, uint32(v28)+32)) = v30
							v33 = F_palloc(m, v18)
							mBase = m.M
							v34 = m.ExcPending
							if v34 != 0 {
								return int32(0)
							} else {
								*(*int32)(unsafe.Add(mBase, uint32(v28)+36)) = v33
								v36 = *(*int32)(unsafe.Add(mBase, uint32(v8)+12))
								*(*int32)(unsafe.Add(mBase, uint32(v28)+24)) = int32(6894)
								*(*int32)(unsafe.Add(mBase, uint32(v28)+20)) = int32(6895)
								*(*int32)(unsafe.Add(mBase, uint32(v28)+16)) = int32(6896)
								*(*int32)(unsafe.Add(mBase, uint32(v28)+12)) = int32(6897)
								*(*int32)(unsafe.Add(mBase, uint32(v28)+8)) = int32(6898)
								*(*int32)(unsafe.Add(mBase, uint32(v28)+4)) = int32(6899)
								*(*int32)(unsafe.Add(mBase, uint32(v28))) = int32(6900)
								*(*int32)(unsafe.Add(mBase, uint32(v28)+28)) = v36
								*(*int32)(unsafe.Add(mBase, uint32(l1))) = v28
								v62 = int32(0)
								m.G0 = v8 + int32(16)
								return v62
							}
						}
					}
				}
			}
		}
	}
}
func F_px_set_debug_handler(m *base.Module, l0 int32) {
	mBase := m.M
	_ = mBase
	*(*int32)(unsafe.Add(mBase, _consts[1097])) = l0
	return
}
func F_px_strerror(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v13 int32
	_ = v13
	var v16 int32
	_ = v16
	var v19 int32
	_ = v19
	var v25 int32
	_ = v25
	var v29 int32
	_ = v29
	var v32 int32
	_ = v32
	if l0 == int32(0) {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	return int32(315301)
L2:
	;
	goto L3
L3:
	;
	v13 = int32(4395472)
	goto L5
L4:
	;
	return v32
L5:
	;
	v16 = *(*int32)(unsafe.Add(mBase, uint32(v13)+8))
	if l0 != v16 {
		goto L7
	} else {
		goto L8
	}
L6:
	;
	v29 = *(*int32)(unsafe.Add(mBase, uint32(v13)+12))
	v32 = v29
	goto L4
L7:
	;
	v19 = *(*int32)(unsafe.Add(mBase, uint32(v13)+20))
	if v19 == int32(0) {
		goto L10
	} else {
		goto L11
	}
L8:
	;
	goto L9
L9:
	;
	goto L6
L10:
	;
	return int32(413729)
L11:
	;
	goto L12
L12:
	;
	v25 = *(*int32)(unsafe.Add(mBase, uint32(v13)+16))
	if l0 != v25 {
		v13 = v13 + int32(16)
		goto L5
	} else {
		goto L13
	}
L13:
	;
	v32 = v19
	goto L4
}
