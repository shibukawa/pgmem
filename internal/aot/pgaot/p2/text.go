package p2

import (
	base "github.com/shibukawa/pgmem/internal/aot/pgaot/base"
	"unsafe"
)

func F_replace_text_regexp(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32, l5 int32, l6 int32) int32 {
	mBase := m.M
	_ = mBase
	var v8 int32
	_ = v8
	var v21 int32
	_ = v21
	var v23 int32
	_ = v23
	var v25 int32
	_ = v25
	var v31 int32
	_ = v31
	var v34 int32
	_ = v34
	var v41 int32
	_ = v41
	var v42 int32
	_ = v42
	var v48 int32
	_ = v48
	var v54 int32
	_ = v54
	var v60 int32
	_ = v60
	var v65 int32
	_ = v65
	var v66 int32
	_ = v66
	var v67 int32
	_ = v67
	var v69 int32
	_ = v69
	var v72 int32
	_ = v72
	var v74 int32
	_ = v74
	var v75 int32
	_ = v75
	var v76 int32
	_ = v76
	var v78 int32
	_ = v78
	var v80 int32
	_ = v80
	var v81 int32
	_ = v81
	var v82 int32
	_ = v82
	var v85 int32
	_ = v85
	var v96 int32
	_ = v96
	var v97 int32
	_ = v97
	var v101 int32
	_ = v101
	var v107 int32
	_ = v107
	var v112 int32
	_ = v112
	var v120 int32
	_ = v120
	var v128 int32
	_ = v128
	var v134 int32
	_ = v134
	var v135 int32
	_ = v135
	var v148 int32
	_ = v148
	var v150 int32
	_ = v150
	var v153 int32
	_ = v153
	var v155 int32
	_ = v155
	var v156 int32
	_ = v156
	var v157 int32
	_ = v157
	var v158 int32
	_ = v158
	var v160 int32
	_ = v160
	var v165 int32
	_ = v165
	var v167 int32
	_ = v167
	var v168 int32
	_ = v168
	var v174 int32
	_ = v174
	var v183 int32
	_ = v183
	var v185 int32
	_ = v185
	var v189 int32
	_ = v189
	var v190 int32
	_ = v190
	var v193 int32
	_ = v193
	var v197 int32
	_ = v197
	var v198 int32
	_ = v198
	var v200 int32
	_ = v200
	var v203 int32
	_ = v203
	var v205 int32
	_ = v205
	var v210 int32
	_ = v210
	var v212 int32
	_ = v212
	var v217 int32
	_ = v217
	var v219 int32
	_ = v219
	var v222 int32
	_ = v222
	var v224 int32
	_ = v224
	var v227 int32
	_ = v227
	var v239 int32
	_ = v239
	var v243 int32
	_ = v243
	var v245 int32
	_ = v245
	var v256 int32
	_ = v256
	var v257 int32
	_ = v257
	var v274 int32
	_ = v274
	var v282 int32
	_ = v282
	var v283 int32
	_ = v283
	var v284 int32
	_ = v284
	var v285 int32
	_ = v285
	var v286 int32
	_ = v286
	var v288 int32
	_ = v288
	var v291 int32
	_ = v291
	var v292 int32
	_ = v292
	var v297 int32
	_ = v297
	var v304 int32
	_ = v304
	var v307 int32
	_ = v307
	var v308 int32
	_ = v308
	var v321 int32
	_ = v321
	var v324 int32
	_ = v324
	var v326 int32
	_ = v326
	var v329 int32
	_ = v329
	var v330 int32
	_ = v330
	var v334 int32
	_ = v334
	var v336 int32
	_ = v336
	var v340 int32
	_ = v340
	var v343 int32
	_ = v343
	var v347 int32
	_ = v347
	var v352 int32
	_ = v352
	var v354 int32
	_ = v354
	var v356 int32
	_ = v356
	var v361 int32
	_ = v361
	var v362 int32
	_ = v362
	var v363 int32
	_ = v363
	var v369 int32
	_ = v369
	var v370 int32
	_ = v370
	var v375 int32
	_ = v375
	var v381 int32
	_ = v381
	var v385 int32
	_ = v385
	var v398 int32
	_ = v398
	var v399 int32
	_ = v399
	var v400 int32
	_ = v400
	var v401 int32
	_ = v401
	var v413 int32
	_ = v413
	var v427 int32
	_ = v427
	var v429 int32
	_ = v429
	var v431 int32
	_ = v431
	var v434 int32
	_ = v434
	var v450 int32
	_ = v450
	var v451 int32
	_ = v451
	var v454 int32
	_ = v454
	var v455 int32
	_ = v455
	var v456 int32
	_ = v456
	var v459 int32
	_ = v459
	var v470 int32
	_ = v470
	var v473 int32
	_ = v473
	var v479 int32
	_ = v479
	var v485 int32
	_ = v485
	var v490 int32
	_ = v490
	var v500 int32
	_ = v500
	var v512 int32
	_ = v512
	var v513 int32
	_ = v513
	var v526 int32
	_ = v526
	var v528 int32
	_ = v528
	var v531 int32
	_ = v531
	var v533 int32
	_ = v533
	var v534 int32
	_ = v534
	var v535 int32
	_ = v535
	var v536 int32
	_ = v536
	var v538 int32
	_ = v538
	var v543 int32
	_ = v543
	var v545 int32
	_ = v545
	var v546 int32
	_ = v546
	var v552 int32
	_ = v552
	var v561 int32
	_ = v561
	var v563 int32
	_ = v563
	var v567 int32
	_ = v567
	var v568 int32
	_ = v568
	var v571 int32
	_ = v571
	var v575 int32
	_ = v575
	var v576 int32
	_ = v576
	var v578 int32
	_ = v578
	var v581 int32
	_ = v581
	var v583 int32
	_ = v583
	var v588 int32
	_ = v588
	var v590 int32
	_ = v590
	var v595 int32
	_ = v595
	var v597 int32
	_ = v597
	var v600 int32
	_ = v600
	var v602 int32
	_ = v602
	var v605 int32
	_ = v605
	var v617 int32
	_ = v617
	var v618 int32
	_ = v618
	var v624 int32
	_ = v624
	var v627 int32
	_ = v627
	var v633 int32
	_ = v633
	var v634 int32
	_ = v634
	var v645 int32
	_ = v645
	var v660 int32
	_ = v660
	var v667 int32
	_ = v667
	var v668 int32
	_ = v668
	var v669 int32
	_ = v669
	var v671 int32
	_ = v671
	var v672 int32
	_ = v672
	var v675 int32
	_ = v675
	var v678 int32
	_ = v678
	var v680 int32
	_ = v680
	var v681 int32
	_ = v681
	var v686 int32
	_ = v686
	var v694 int32
	_ = v694
	var v698 int32
	_ = v698
	var v711 int32
	_ = v711
	var v712 int32
	_ = v712
	var v713 int32
	_ = v713
	var v714 int32
	_ = v714
	var v721 int32
	_ = v721
	var v746 int32
	_ = v746
	var v759 int32
	_ = v759
	var v762 int32
	_ = v762
	var v764 int32
	_ = v764
	var v765 int32
	_ = v765
	var v770 int32
	_ = v770
	var v778 int32
	_ = v778
	var v782 int32
	_ = v782
	var v795 int32
	_ = v795
	var v796 int32
	_ = v796
	var v797 int32
	_ = v797
	var v798 int32
	_ = v798
	var v809 int32
	_ = v809
	var v843 int32
	_ = v843
	var v845 int32
	_ = v845
	var v855 int32
	_ = v855
	var v869 int32
	_ = v869
	var v873 int32
	_ = v873
	var v880 int32
	_ = v880
	var v883 int32
	_ = v883
	var v890 int32
	_ = v890
	var v891 int32
	_ = v891
	var v897 int32
	_ = v897
	var v903 int32
	_ = v903
	var v905 int32
	_ = v905
	var v926 int32
	_ = v926
	var v927 int32
	_ = v927
	var v929 int32
	_ = v929
	var v930 int32
	_ = v930
	var v935 int32
	_ = v935
	var v943 int32
	_ = v943
	var v947 int32
	_ = v947
	var v960 int32
	_ = v960
	var v961 int32
	_ = v961
	var v962 int32
	_ = v962
	var v963 int32
	_ = v963
	var v970 int32
	_ = v970
	var v1008 int32
	_ = v1008
	var v1009 int32
	_ = v1009
	var v1010 int32
	_ = v1010
	var v1014 int32
	_ = v1014
	var v1017 int32
	_ = v1017
	var v1020 int32
	_ = v1020
	var v1033 int32
	_ = v1033
	var v1035 int32
	_ = v1035
	var v1041 int32
	_ = v1041
	var v1044 int32
	_ = v1044
	var v1060 int32
	_ = v1060
	var v1064 int32
	_ = v1064
	var v1066 int32
	_ = v1066
	var v1069 int32
	_ = v1069
	var v1076 int32
	_ = v1076
	var v1077 int32
	_ = v1077
	var v1081 int32
	_ = v1081
	var v1085 int32
	_ = v1085
	var v1089 int32
	_ = v1089
	var v1091 int32
	_ = v1091
	var v1092 int32
	_ = v1092
	var v1094 int32
	_ = v1094
	var v1095 int32
	_ = v1095
	var v1096 int32
	_ = v1096
	var v1103 int32
	_ = v1103
	var v1105 int32
	_ = v1105
	var v1107 int32
	_ = v1107
	v8 = int32(0)
	v21 = m.G0
	v23 = v21 - int32(224)
	m.G0 = v23
	v25 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0))))
	if v25 == int32(1) {
		goto L2
	} else {
		goto L3
	}
L1:
	;
	F_initStringInfo(m, v23+int32(208))
	mBase = m.M
	v60 = m.ExcPending
	if v60 != 0 {
		goto L12
	} else {
		goto L13
	}
L2:
	;
	v31 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+1)))
	if v31 == int32(18) {
		goto L5
	} else {
		goto L6
	}
L3:
	;
	goto L4
L4:
	;
	v42 = int32(1)
	if v25&v42 != 0 {
		v54 = int32(base.Ui32(v25)>>(uint(v42)%32)) - v42
		goto L1
	} else {
		goto L11
	}
L5:
	;
	v34 = int32(16)
	goto L7
L6:
	;
	v34 = int32(0)
	goto L7
L7:
	;
	if base.Ui32((v31-int32(1))&int32(255)) < base.Ui32(int32(3)) {
		goto L8
	} else {
		goto L9
	}
L8:
	;
	v41 = int32(4)
	goto L10
L9:
	;
	v41 = v34
	goto L10
L10:
	;
	v54 = v41
	goto L1
L11:
	;
	v48 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v54 = int32(base.Ui32(v48)>>(uint(int32(2))%32)) - int32(4)
	goto L1
L12:
	;
	return int32(0)
L13:
	;
	v65 = F_palloc(m, v54<<(uint(int32(2))%32)+int32(4))
	mBase = m.M
	v66 = m.ExcPending
	if v66 != 0 {
		goto L12
	} else {
		goto L14
	}
L14:
	;
	v67 = int32(1)
	v69 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0))))
	if v69&v67 != 0 {
		goto L15
	} else {
		goto L16
	}
L15:
	;
	v72 = v67
	goto L17
L16:
	;
	v72 = int32(4)
	goto L17
L17:
	;
	v74 = F_pg_mb2wchar_with_len(m, l0+v72, v65, v54)
	mBase = m.M
	v75 = m.ExcPending
	if v75 != 0 {
		goto L12
	} else {
		goto L18
	}
L18:
	;
	v76 = int32(1)
	v78 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l2))))
	v80 = v78 & v76
	if v80 != 0 {
		goto L19
	} else {
		goto L20
	}
L19:
	;
	v81 = v76
	goto L21
L20:
	;
	v81 = int32(4)
	goto L21
L21:
	;
	v82 = l2 + v81
	if v78 == int32(1) {
		goto L26
	} else {
		goto L27
	}
L22:
	;
	v282 = base.B2i32(base.Ui32(v274) < base.Ui32(int32(2)))
	if base.Ui32(v274) < base.Ui32(int32(2)) {
		goto L72
	} else {
		goto L73
	}
L23:
	;
	v120 = v82
	v128 = v8
	goto L37
L24:
	;
	v112 = v82 + int32(4)
	goto L23
L25:
	;
	if v107 != 0 {
		goto L34
	} else {
		goto L35
	}
L26:
	;
	v85 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l2)+1)))
	if base.Ui32((v85-int32(1))&int32(255)) < base.Ui32(int32(3)) {
		goto L24
	} else {
		goto L29
	}
L27:
	;
	goto L28
L28:
	;
	v97 = int32(1)
	if v80 != 0 {
		v107 = int32(base.Ui32(v78)>>(uint(v97)%32)) - v97
		goto L25
	} else {
		goto L33
	}
L29:
	;
	if v85 == int32(18) {
		goto L30
	} else {
		goto L31
	}
L30:
	;
	v96 = int32(16)
	goto L32
L31:
	;
	v96 = int32(0)
	goto L32
L32:
	;
	v107 = v96
	goto L25
L33:
	;
	v101 = *(*int32)(unsafe.Add(mBase, uint32(l2)))
	v107 = int32(base.Ui32(v101)>>(uint(int32(2))%32)) - int32(4)
	goto L25
L34:
	;
	v112 = v82 + v107
	goto L23
L35:
	;
	goto L36
L36:
	;
	v274 = v8
	goto L22
L37:
	;
	v134 = v112 - v120
	v135 = int32(0)
	if base.B2i32(v120&int32(3) == v135)|base.B2i32(v134 == v135) != 0 {
		v165 = v120
		v167 = v134
		v168 = base.B2i32(v134 != v135)
		goto L42
	} else {
		goto L43
	}
L38:
	;
	v274 = v257
	goto L22
L39:
	;
	if v239 == int32(0) {
		v274 = v128
		goto L22
	} else {
		goto L64
	}
L40:
	;
	v239 = int32(0)
	goto L39
L41:
	;
	v217 = v210
	v219 = v212
	goto L58
L42:
	;
	if v168 == int32(0) {
		goto L40
	} else {
		goto L49
	}
L43:
	;
	v148 = v120
	v150 = v134
	goto L44
L44:
	;
	v153 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v148))))
	if v153 == int32(92) {
		v210 = v148
		v212 = v150
		goto L41
	} else {
		goto L46
	}
L45:
	;
	v165 = v160
	v167 = v156
	v168 = v158
	goto L42
L46:
	;
	v155 = int32(1)
	v156 = v150 - v155
	v157 = int32(0)
	v158 = base.B2i32(v156 != v157)
	v160 = v148 + v155
	if v160&int32(3) == v157 {
		v165 = v160
		v167 = v156
		v168 = v158
		goto L42
	} else {
		goto L47
	}
L47:
	;
	if v156 != 0 {
		v148 = v160
		v150 = v156
		goto L44
	} else {
		goto L48
	}
L48:
	;
	goto L45
L49:
	;
	v174 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v165))))
	if base.B2i32(int32(92) == v174)|base.B2i32(base.Ui32(v167) < base.Ui32(int32(4))) == int32(0) {
		goto L50
	} else {
		goto L51
	}
L50:
	;
	v183 = v165
	v185 = v167
	goto L53
L51:
	;
	v203 = v165
	v205 = v167
	goto L52
L52:
	;
	if v205 == int32(0) {
		goto L40
	} else {
		goto L57
	}
L53:
	;
	v189 = *(*int32)(unsafe.Add(mBase, uint32(v183)))
	v190 = v189 ^ int32(1549556828)
	v193 = int32(-2139062144)
	if (int32(16843008)-v190|v190)&v193 != v193 {
		v210 = v183
		v212 = v185
		goto L41
	} else {
		goto L55
	}
L54:
	;
	v203 = v198
	v205 = v200
	goto L52
L55:
	;
	v197 = int32(4)
	v198 = v183 + v197
	v200 = v185 - v197
	if base.Ui32(int32(3)) < base.Ui32(v200) {
		v183 = v198
		v185 = v200
		goto L53
	} else {
		goto L56
	}
L56:
	;
	goto L54
L57:
	;
	v210 = v203
	v212 = v205
	goto L41
L58:
	;
	v222 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v217))))
	if int32(92) == v222 {
		goto L60
	} else {
		goto L61
	}
L59:
	;
	goto L40
L60:
	;
	v239 = v217
	goto L39
L61:
	;
	goto L62
L62:
	;
	v224 = int32(1)
	v227 = v219 - v224
	if v227 != 0 {
		v217 = v217 + v224
		v219 = v227
		goto L58
	} else {
		goto L63
	}
L63:
	;
	goto L59
L64:
	;
	v243 = v239 + int32(1)
	if base.Ui32(v243) < base.Ui32(v112) {
		goto L65
	} else {
		goto L66
	}
L65:
	;
	v245 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v243))))
	if base.Ui32((v245-int32(49))&int32(255)) < base.Ui32(int32(9)) {
		goto L68
	} else {
		goto L69
	}
L66:
	;
	v256 = v243
	v257 = v128
	goto L67
L67:
	;
	if base.Ui32(v256) < base.Ui32(v112) {
		v120 = v256
		v128 = v257
		goto L37
	} else {
		goto L71
	}
L68:
	;
	v274 = int32(2)
	goto L22
L69:
	;
	goto L70
L70:
	;
	v256 = v239 + int32(2)
	v257 = int32(1)
	goto L67
L71:
	;
	goto L38
L72:
	;
	v283 = l3 | int32(16)
	goto L74
L73:
	;
	v283 = l3
	goto L74
L74:
	;
	v284 = F_RE_compile_and_cache(m, l1, v283, l4)
	mBase = m.M
	v285 = m.ExcPending
	if v285 != 0 {
		goto L12
	} else {
		goto L75
	}
L75:
	;
	v286 = int32(1)
	v288 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0))))
	if v288&v286 != 0 {
		goto L76
	} else {
		goto L77
	}
L76:
	;
	v291 = v286
	goto L78
L77:
	;
	v291 = int32(4)
	goto L78
L78:
	;
	v292 = l0 + v291
	if base.Ui32(v74) < base.Ui32(l5) {
		v1041 = v292
		v1044 = int32(0)
		goto L79
	} else {
		goto L80
	}
L79:
	;
	if base.Ui32(v1044) < base.Ui32(v74) {
		goto L242
	} else {
		goto L243
	}
L80:
	;
	if base.Ui32(v274) < base.Ui32(int32(2)) {
		goto L81
	} else {
		goto L82
	}
L81:
	;
	v297 = int32(1)
	goto L83
L82:
	;
	v297 = int32(10)
	goto L83
L83:
	;
	v304 = int32(0)
	v307 = v292
	v308 = l5
	v321 = v8
	goto L84
L84:
	;
	v324 = *(*int32)(unsafe.Add(mBase, _c_F_replace_text_regexp[0]))
	if v324 != 0 {
		goto L86
	} else {
		goto L87
	}
L85:
	;
	v1041 = v1017
	v1044 = v1014
	goto L79
L86:
	;
	F_ProcessInterrupts(m)
	mBase = m.M
	v326 = m.ExcPending
	if v326 != 0 {
		goto L12
	} else {
		goto L89
	}
L87:
	;
	goto L88
L88:
	;
	v329 = F_pg_regexec(m, v284, v65, v74, v308, v297, v23+int32(128))
	mBase = m.M
	v330 = m.ExcPending
	if v330 != 0 {
		goto L12
	} else {
		goto L90
	}
L89:
	;
	goto L88
L90:
	;
	if v329 != 0 {
		goto L91
	} else {
		goto L92
	}
L91:
	;
	if v329 == int32(1) {
		v1041 = v307
		v1044 = v304
		goto L79
	} else {
		goto L94
	}
L92:
	;
	goto L93
L93:
	;
	v354 = v321 + int32(1)
	v356 = int32(0)
	if base.B2i32(v354 == l6)|base.B2i32(l6 <= v356) == v356 {
		goto L101
	} else {
		goto L102
	}
L94:
	;
	v334 = v23 + int32(16)
	F_pg_regerror(m, v329, v334)
	mBase = m.M
	v336 = m.ExcPending
	if v336 != 0 {
		goto L12
	} else {
		goto L95
	}
L95:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v340 = m.ExcPending
	if v340 != 0 {
		goto L12
	} else {
		goto L96
	}
L96:
	;
	F_errcode(m, int32(302252162))
	mBase = m.M
	v343 = m.ExcPending
	if v343 != 0 {
		goto L12
	} else {
		goto L97
	}
L97:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v23))) = v334
	F_errmsg(m, int32(_a_F_replace_text_regexp_0), v23)
	mBase = m.M
	v347 = m.ExcPending
	if v347 != 0 {
		goto L12
	} else {
		goto L98
	}
L98:
	;
	F_errfinish(m, int32(_a_F_replace_text_regexp_1), int32(3414), int32(_a_F_replace_text_regexp_2))
	mBase = m.M
	v352 = m.ExcPending
	if v352 != 0 {
		goto L12
	} else {
		goto L99
	}
L99:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L100:
	;
	v1033 = *(*int32)(unsafe.Add(mBase, uint32(v23)+128))
	v1035 = v1020 + base.B2i32(v1033 == v1020)
	if base.Ui32(v1035) <= base.Ui32(v74) {
		v304 = v1014
		v307 = v1017
		v308 = v1035
		v321 = v354
		goto L84
	} else {
		goto L241
	}
L101:
	;
	v361 = *(*int32)(unsafe.Add(mBase, uint32(v23)+132))
	v1014 = v304
	v1017 = v307
	v1020 = v361
	goto L100
L102:
	;
	goto L103
L103:
	;
	v362 = *(*int32)(unsafe.Add(mBase, uint32(v23)+128))
	v363 = v362 - v304
	if int32(0) < v363 {
		goto L104
	} else {
		goto L105
	}
L104:
	;
	v369 = *(*int32)(unsafe.Add(mBase, _c_F_replace_text_regexp[1]))
	v370 = *(*int32)(unsafe.Add(mBase, uint32(v369)+4))
	v375 = *(*int32)(unsafe.Add(mBase, uint32(v370*int32(28))+uint32(_c_F_replace_text_regexp[2])))
	goto L107
L105:
	;
	v431 = v304
	v434 = v307
	goto L106
L106:
	;
	v450 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l2))))
	if v274 != 0 {
		goto L117
	} else {
		goto L118
	}
L107:
	;
	if v375 != int32(1) {
		goto L108
	} else {
		goto L109
	}
L108:
	;
	v381 = v307
	v385 = v363
	goto L111
L109:
	;
	v413 = v363
	goto L110
L110:
	;
	F_appendBinaryStringInfo(m, v23+int32(208), v307, v413)
	mBase = m.M
	v427 = m.ExcPending
	if v427 != 0 {
		goto L12
	} else {
		goto L115
	}
L111:
	;
	v398 = F_pg_mblen_unbounded(m, v381)
	mBase = m.M
	v399 = m.ExcPending
	if v399 != 0 {
		goto L12
	} else {
		goto L113
	}
L112:
	;
	v413 = v400 - v307
	goto L110
L113:
	;
	v400 = v398 + v381
	v401 = int32(1)
	if base.Ui32(v401) < base.Ui32(v385) {
		v381 = v400
		v385 = v385 - v401
		goto L111
	} else {
		goto L114
	}
L114:
	;
	goto L112
L115:
	;
	v429 = *(*int32)(unsafe.Add(mBase, uint32(v23)+128))
	v431 = v429
	v434 = v307 + v413
	goto L106
L116:
	;
	v926 = *(*int32)(unsafe.Add(mBase, uint32(v23)+132))
	v927 = v926 - v431
	v929 = *(*int32)(unsafe.Add(mBase, _c_F_replace_text_regexp[1]))
	v930 = *(*int32)(unsafe.Add(mBase, uint32(v929)+4))
	v935 = *(*int32)(unsafe.Add(mBase, uint32(v930*int32(28))+uint32(_c_F_replace_text_regexp[2])))
	goto L229
L117:
	;
	v451 = int32(1)
	v454 = v450 & v451
	if v454 != 0 {
		goto L120
	} else {
		goto L121
	}
L118:
	;
	goto L119
L119:
	;
	v869 = int32(1)
	if v450&v869 != 0 {
		goto L214
	} else {
		goto L215
	}
L120:
	;
	v455 = v451
	goto L122
L121:
	;
	v455 = int32(4)
	goto L122
L122:
	;
	v456 = l2 + v455
	if v450 == int32(1) {
		goto L127
	} else {
		goto L128
	}
L123:
	;
	v500 = v456
	goto L136
L124:
	;
	if v485 == int32(0) {
		goto L116
	} else {
		goto L135
	}
L125:
	;
	v479 = *(*int32)(unsafe.Add(mBase, uint32(l2)))
	v485 = int32(base.Ui32(v479)>>(uint(int32(2))%32)) - int32(4)
	goto L124
L126:
	;
	v490 = v456 + int32(4)
	goto L123
L127:
	;
	v459 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l2)+1)))
	if base.Ui32((v459-int32(1))&int32(255)) < base.Ui32(int32(3)) {
		goto L126
	} else {
		goto L130
	}
L128:
	;
	goto L129
L129:
	;
	if v454 == int32(0) {
		goto L125
	} else {
		goto L134
	}
L130:
	;
	if v459 == int32(18) {
		goto L131
	} else {
		goto L132
	}
L131:
	;
	v470 = int32(16)
	goto L133
L132:
	;
	v470 = int32(0)
	goto L133
L133:
	;
	v485 = v470
	goto L124
L134:
	;
	v473 = int32(1)
	v485 = int32(base.Ui32(v450)>>(uint(v473)%32)) - v473
	goto L124
L135:
	;
	v490 = v485 + v456
	goto L123
L136:
	;
	v512 = v490 - v500
	v513 = int32(0)
	if base.B2i32(v500&int32(3) == v513)|base.B2i32(v512 == v513) != 0 {
		v543 = v500
		v545 = v512
		v546 = base.B2i32(v512 != v513)
		goto L141
	} else {
		goto L142
	}
L137:
	;
	goto L116
L138:
	;
	if v617 != 0 {
		goto L163
	} else {
		goto L164
	}
L139:
	;
	v617 = int32(0)
	goto L138
L140:
	;
	v595 = v588
	v597 = v590
	goto L157
L141:
	;
	if v546 == int32(0) {
		goto L139
	} else {
		goto L148
	}
L142:
	;
	v526 = v500
	v528 = v512
	goto L143
L143:
	;
	v531 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v526))))
	if v531 == int32(92) {
		v588 = v526
		v590 = v528
		goto L140
	} else {
		goto L145
	}
L144:
	;
	v543 = v538
	v545 = v534
	v546 = v536
	goto L141
L145:
	;
	v533 = int32(1)
	v534 = v528 - v533
	v535 = int32(0)
	v536 = base.B2i32(v534 != v535)
	v538 = v526 + v533
	if v538&int32(3) == v535 {
		v543 = v538
		v545 = v534
		v546 = v536
		goto L141
	} else {
		goto L146
	}
L146:
	;
	if v534 != 0 {
		v526 = v538
		v528 = v534
		goto L143
	} else {
		goto L147
	}
L147:
	;
	goto L144
L148:
	;
	v552 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v543))))
	if base.B2i32(int32(92) == v552)|base.B2i32(base.Ui32(v545) < base.Ui32(int32(4))) == int32(0) {
		goto L149
	} else {
		goto L150
	}
L149:
	;
	v561 = v543
	v563 = v545
	goto L152
L150:
	;
	v581 = v543
	v583 = v545
	goto L151
L151:
	;
	if v583 == int32(0) {
		goto L139
	} else {
		goto L156
	}
L152:
	;
	v567 = *(*int32)(unsafe.Add(mBase, uint32(v561)))
	v568 = v567 ^ int32(1549556828)
	v571 = int32(-2139062144)
	if (int32(16843008)-v568|v568)&v571 != v571 {
		v588 = v561
		v590 = v563
		goto L140
	} else {
		goto L154
	}
L153:
	;
	v581 = v576
	v583 = v578
	goto L151
L154:
	;
	v575 = int32(4)
	v576 = v561 + v575
	v578 = v563 - v575
	if base.Ui32(int32(3)) < base.Ui32(v578) {
		v561 = v576
		v563 = v578
		goto L152
	} else {
		goto L155
	}
L155:
	;
	goto L153
L156:
	;
	v588 = v581
	v590 = v583
	goto L140
L157:
	;
	v600 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v595))))
	if int32(92) == v600 {
		goto L159
	} else {
		goto L160
	}
L158:
	;
	goto L139
L159:
	;
	v617 = v595
	goto L138
L160:
	;
	goto L161
L161:
	;
	v602 = int32(1)
	v605 = v597 - v602
	if v605 != 0 {
		v595 = v595 + v602
		v597 = v605
		goto L157
	} else {
		goto L162
	}
L162:
	;
	goto L158
L163:
	;
	v618 = v617
	goto L165
L164:
	;
	v618 = v490
	goto L165
L165:
	;
	if base.Ui32(v500) < base.Ui32(v618) {
		goto L166
	} else {
		goto L167
	}
L166:
	;
	F_appendBinaryStringInfo(m, v23+int32(208), v500, v618-v500)
	mBase = m.M
	v624 = m.ExcPending
	if v624 != 0 {
		goto L12
	} else {
		goto L169
	}
L167:
	;
	goto L168
L168:
	;
	if base.Ui32(v490) <= base.Ui32(v618) {
		goto L116
	} else {
		goto L170
	}
L169:
	;
	goto L168
L170:
	;
	v627 = v618 + int32(1)
	if base.Ui32(v490) <= base.Ui32(v627) {
		goto L171
	} else {
		goto L172
	}
L171:
	;
	F_appendStringInfoChar(m, v23+int32(208), int32(92))
	mBase = m.M
	v633 = m.ExcPending
	if v633 != 0 {
		goto L12
	} else {
		goto L174
	}
L172:
	;
	goto L173
L173:
	;
	v634 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v627))))
	if base.Ui32((v634-int32(49))&int32(255)) <= base.Ui32(int32(8)) {
		goto L177
	} else {
		goto L178
	}
L174:
	;
	goto L116
L175:
	;
	if base.Ui32(v855) < base.Ui32(v490) {
		v500 = v855
		goto L136
	} else {
		goto L213
	}
L176:
	;
	v671 = v618 + int32(2)
	v672 = *(*int32)(unsafe.Add(mBase, uint32(v668)))
	if v672 < int32(0) {
		v855 = v671
		goto L175
	} else {
		goto L188
	}
L177:
	;
	v645 = v23 + int32(128) + v634<<(uint(int32(3))%32)
	v668 = v645 - int32(384)
	v669 = v645 - int32(380)
	goto L176
L178:
	;
	goto L179
L179:
	;
	if v634 == int32(38) {
		goto L180
	} else {
		goto L181
	}
L180:
	;
	v668 = v23 + int32(128)
	v669 = v23 + int32(128) | int32(4)
	goto L176
L181:
	;
	goto L182
L182:
	;
	if v634 == int32(92) {
		goto L183
	} else {
		goto L184
	}
L183:
	;
	F_appendStringInfoChar(m, v23+int32(208), int32(92))
	mBase = m.M
	v660 = m.ExcPending
	if v660 != 0 {
		goto L12
	} else {
		goto L186
	}
L184:
	;
	goto L185
L185:
	;
	F_appendStringInfoChar(m, v23+int32(208), int32(92))
	mBase = m.M
	v667 = m.ExcPending
	if v667 != 0 {
		goto L12
	} else {
		goto L187
	}
L186:
	;
	v855 = v618 + int32(2)
	goto L175
L187:
	;
	v855 = v627
	goto L175
L188:
	;
	v675 = *(*int32)(unsafe.Add(mBase, uint32(v669)))
	if v675 < int32(0) {
		v855 = v671
		goto L175
	} else {
		goto L189
	}
L189:
	;
	v678 = v672 - v431
	v680 = *(*int32)(unsafe.Add(mBase, _c_F_replace_text_regexp[1]))
	v681 = *(*int32)(unsafe.Add(mBase, uint32(v680)+4))
	v686 = *(*int32)(unsafe.Add(mBase, uint32(v681*int32(28))+uint32(_c_F_replace_text_regexp[2])))
	goto L190
L190:
	;
	if v686 != int32(1) {
		goto L191
	} else {
		goto L192
	}
L191:
	;
	if int32(0) < v678 {
		goto L194
	} else {
		goto L195
	}
L192:
	;
	v746 = v678
	goto L193
L193:
	;
	v759 = v675 - v672
	v762 = v434 + v746
	v764 = *(*int32)(unsafe.Add(mBase, _c_F_replace_text_regexp[1]))
	v765 = *(*int32)(unsafe.Add(mBase, uint32(v764)+4))
	v770 = *(*int32)(unsafe.Add(mBase, uint32(v765*int32(28))+uint32(_c_F_replace_text_regexp[2])))
	goto L201
L194:
	;
	v694 = v434
	v698 = v678
	goto L197
L195:
	;
	v721 = v434
	goto L196
L196:
	;
	v746 = v721 - v434
	goto L193
L197:
	;
	v711 = F_pg_mblen_unbounded(m, v694)
	mBase = m.M
	v712 = m.ExcPending
	if v712 != 0 {
		goto L12
	} else {
		goto L199
	}
L198:
	;
	v721 = v713
	goto L196
L199:
	;
	v713 = v711 + v694
	v714 = int32(1)
	if base.Ui32(v714) < base.Ui32(v698) {
		v694 = v713
		v698 = v698 - v714
		goto L197
	} else {
		goto L200
	}
L200:
	;
	goto L198
L201:
	;
	if v770 != int32(1) {
		goto L202
	} else {
		goto L203
	}
L202:
	;
	if int32(0) < v759 {
		goto L205
	} else {
		goto L206
	}
L203:
	;
	v843 = v759
	goto L204
L204:
	;
	F_appendBinaryStringInfo(m, v23+int32(208), v762, v843)
	mBase = m.M
	v845 = m.ExcPending
	if v845 != 0 {
		goto L12
	} else {
		goto L212
	}
L205:
	;
	v778 = v759
	v782 = v762
	goto L208
L206:
	;
	v809 = v762
	goto L207
L207:
	;
	v843 = v809 - v762
	goto L204
L208:
	;
	v795 = F_pg_mblen_unbounded(m, v782)
	mBase = m.M
	v796 = m.ExcPending
	if v796 != 0 {
		goto L12
	} else {
		goto L210
	}
L209:
	;
	v809 = v797
	goto L207
L210:
	;
	v797 = v795 + v782
	v798 = int32(1)
	if base.Ui32(v798) < base.Ui32(v778) {
		v778 = v778 - v798
		v782 = v797
		goto L208
	} else {
		goto L211
	}
L211:
	;
	goto L209
L212:
	;
	v855 = v671
	goto L175
L213:
	;
	goto L137
L214:
	;
	v873 = v869
	goto L216
L215:
	;
	v873 = int32(4)
	goto L216
L216:
	;
	if v450 == int32(1) {
		goto L218
	} else {
		goto L219
	}
L217:
	;
	F_appendBinaryStringInfo(m, v23+int32(208), l2+v873, v903)
	mBase = m.M
	v905 = m.ExcPending
	if v905 != 0 {
		goto L12
	} else {
		goto L228
	}
L218:
	;
	v880 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l2)+1)))
	if v880 == int32(18) {
		goto L221
	} else {
		goto L222
	}
L219:
	;
	goto L220
L220:
	;
	v891 = int32(1)
	if v450&v891 != 0 {
		v903 = int32(base.Ui32(v450)>>(uint(v891)%32)) - v891
		goto L217
	} else {
		goto L227
	}
L221:
	;
	v883 = int32(16)
	goto L223
L222:
	;
	v883 = int32(0)
	goto L223
L223:
	;
	if base.Ui32((v880-int32(1))&int32(255)) < base.Ui32(int32(3)) {
		goto L224
	} else {
		goto L225
	}
L224:
	;
	v890 = int32(4)
	goto L226
L225:
	;
	v890 = v883
	goto L226
L226:
	;
	v903 = v890
	goto L217
L227:
	;
	v897 = *(*int32)(unsafe.Add(mBase, uint32(l2)))
	v903 = int32(base.Ui32(v897)>>(uint(int32(2))%32)) - int32(4)
	goto L217
L228:
	;
	goto L116
L229:
	;
	if v935 != int32(1) {
		goto L230
	} else {
		goto L231
	}
L230:
	;
	if int32(0) < v927 {
		goto L233
	} else {
		goto L234
	}
L231:
	;
	v1008 = v927
	goto L232
L232:
	;
	v1009 = v1008 + v434
	v1010 = *(*int32)(unsafe.Add(mBase, uint32(v23)+132))
	if int32(0) < l6 {
		v1041 = v1009
		v1044 = v1010
		goto L79
	} else {
		goto L240
	}
L233:
	;
	v943 = v434
	v947 = v927
	goto L236
L234:
	;
	v970 = v434
	goto L235
L235:
	;
	v1008 = v970 - v434
	goto L232
L236:
	;
	v960 = F_pg_mblen_unbounded(m, v943)
	mBase = m.M
	v961 = m.ExcPending
	if v961 != 0 {
		goto L12
	} else {
		goto L238
	}
L237:
	;
	v970 = v962
	goto L235
L238:
	;
	v962 = v960 + v943
	v963 = int32(1)
	if base.Ui32(v963) < base.Ui32(v947) {
		v943 = v962
		v947 = v947 - v963
		goto L236
	} else {
		goto L239
	}
L239:
	;
	goto L237
L240:
	;
	v1014 = v1010
	v1017 = v1009
	v1020 = v1010
	goto L100
L241:
	;
	goto L85
L242:
	;
	v1060 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0))))
	if v1060 == int32(1) {
		goto L246
	} else {
		goto L247
	}
L243:
	;
	goto L244
L244:
	;
	v1091 = *(*int32)(unsafe.Add(mBase, uint32(v23)+208))
	v1092 = *(*int32)(unsafe.Add(mBase, uint32(v23)+212))
	v1094 = v1092 + int32(4)
	v1095 = F_palloc(m, v1094)
	mBase = m.M
	v1096 = m.ExcPending
	if v1096 != 0 {
		goto L12
	} else {
		goto L257
	}
L245:
	;
	F_appendBinaryStringInfo(m, v23+int32(208), v1041, v1085+l0-v1041)
	mBase = m.M
	v1089 = m.ExcPending
	if v1089 != 0 {
		goto L12
	} else {
		goto L256
	}
L246:
	;
	v1064 = int32(18)
	v1066 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+1)))
	if v1066 == v1064 {
		goto L249
	} else {
		goto L250
	}
L247:
	;
	goto L248
L248:
	;
	v1077 = int32(1)
	if v1060&v1077 != 0 {
		v1085 = int32(base.Ui32(v1060) >> (uint(v1077) % 32))
		goto L245
	} else {
		goto L255
	}
L249:
	;
	v1069 = v1064
	goto L251
L250:
	;
	v1069 = int32(2)
	goto L251
L251:
	;
	if base.Ui32((v1066-int32(1))&int32(255)) < base.Ui32(int32(3)) {
		goto L252
	} else {
		goto L253
	}
L252:
	;
	v1076 = int32(6)
	goto L254
L253:
	;
	v1076 = v1069
	goto L254
L254:
	;
	v1085 = v1076
	goto L245
L255:
	;
	v1081 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v1085 = int32(base.Ui32(v1081) >> (uint(int32(2)) % 32))
	goto L245
L256:
	;
	goto L244
L257:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v1095))) = v1094 << (uint(int32(2)) % 32)
	if v1092 != 0 {
		goto L258
	} else {
		goto L259
	}
L258:
	;
	base.MemoryCopy(m, v1095+int32(4), v1091, v1092)
	goto L260
L259:
	;
	goto L260
L260:
	;
	v1103 = *(*int32)(unsafe.Add(mBase, uint32(v23)+208))
	F_pfree(m, v1103)
	mBase = m.M
	v1105 = m.ExcPending
	if v1105 != 0 {
		goto L12
	} else {
		goto L261
	}
L261:
	;
	F_pfree(m, v65)
	mBase = m.M
	v1107 = m.ExcPending
	if v1107 != 0 {
		goto L12
	} else {
		goto L262
	}
L262:
	;
	m.G0 = v23 + int32(224)
	return v1095
}
func F_split_text(m *base.Module, l0 int32, l1 int32) int32 {
	mBase := m.M
	_ = mBase
	var v3 int32
	_ = v3
	var v14 int32
	_ = v14
	var v16 int32
	_ = v16
	var v18 int32
	_ = v18
	var v19 int32
	_ = v19
	var v20 int32
	_ = v20
	var v21 int32
	_ = v21
	var v24 int32
	_ = v24
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
	var v32 int32
	_ = v32
	var v35 int32
	_ = v35
	var v36 int32
	_ = v36
	var v37 int32
	_ = v37
	var v38 int32
	_ = v38
	var v39 int32
	_ = v39
	var v40 int32
	_ = v40
	var v46 int32
	_ = v46
	var v49 int32
	_ = v49
	var v56 int32
	_ = v56
	var v57 int32
	_ = v57
	var v63 int32
	_ = v63
	var v69 int32
	_ = v69
	var v70 int32
	_ = v70
	var v76 int32
	_ = v76
	var v79 int32
	_ = v79
	var v86 int32
	_ = v86
	var v87 int32
	_ = v87
	var v93 int32
	_ = v93
	var v99 int32
	_ = v99
	var v102 int32
	_ = v102
	var v108 int64
	_ = v108
	var v109 int32
	_ = v109
	var v112 int32
	_ = v112
	var v113 int32
	_ = v113
	var v119 int32
	_ = v119
	var v123 int32
	_ = v123
	var v125 int32
	_ = v125
	var v126 int32
	_ = v126
	var v128 int32
	_ = v128
	var v131 int32
	_ = v131
	var v137 int32
	_ = v137
	var v148 int32
	_ = v148
	var v150 int32
	_ = v150
	var v151 int32
	_ = v151
	var v152 int32
	_ = v152
	var v155 int32
	_ = v155
	var v159 int32
	_ = v159
	var v161 int32
	_ = v161
	var v164 int32
	_ = v164
	var v171 int32
	_ = v171
	var v172 int32
	_ = v172
	var v176 int32
	_ = v176
	var v180 int32
	_ = v180
	var v183 int32
	_ = v183
	var v184 int32
	_ = v184
	var v185 int32
	_ = v185
	var v186 int32
	_ = v186
	var v188 int32
	_ = v188
	var v189 int32
	_ = v189
	var v190 int32
	_ = v190
	var v199 int64
	_ = v199
	var v200 int32
	_ = v200
	var v204 int32
	_ = v204
	var v205 int32
	_ = v205
	var v209 int32
	_ = v209
	var v215 int32
	_ = v215
	var v216 int32
	_ = v216
	var v220 int32
	_ = v220
	var v221 int32
	_ = v221
	var v222 int32
	_ = v222
	var v225 int32
	_ = v225
	var v226 int32
	_ = v226
	var v228 int32
	_ = v228
	var v231 int32
	_ = v231
	var v242 int32
	_ = v242
	var v247 int32
	_ = v247
	var v256 int32
	_ = v256
	var v262 int32
	_ = v262
	var v265 int32
	_ = v265
	var v269 int32
	_ = v269
	var v270 int32
	_ = v270
	var v272 int32
	_ = v272
	var v274 int32
	_ = v274
	var v275 int32
	_ = v275
	var v277 int32
	_ = v277
	var v282 int32
	_ = v282
	var v290 int32
	_ = v290
	var v291 int32
	_ = v291
	var v293 int32
	_ = v293
	var v295 int32
	_ = v295
	var v297 int32
	_ = v297
	var v298 int32
	_ = v298
	var v299 int32
	_ = v299
	var v308 int64
	_ = v308
	var v309 int32
	_ = v309
	var v313 int32
	_ = v313
	var v314 int32
	_ = v314
	var v318 int32
	_ = v318
	var v322 int32
	_ = v322
	var v323 int32
	_ = v323
	var v327 int32
	_ = v327
	var v328 int32
	_ = v328
	var v329 int32
	_ = v329
	var v332 int32
	_ = v332
	var v334 int32
	_ = v334
	var v337 int32
	_ = v337
	var v341 int32
	_ = v341
	var v342 int32
	_ = v342
	var v343 int32
	_ = v343
	v3 = int32(0)
	v14 = m.G0
	v16 = v14 - int32(1088)
	m.G0 = v16
	v18 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+32)))
	if v18 != 0 {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	m.G0 = v16 + int32(1088)
	return v18 ^ int32(1)
L2:
	;
	v19 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	v20 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
	v21 = F_pg_detoast_datum_packed(m, v20)
	mBase = m.M
	v24 = m.ExcPending
	if v24 != 0 {
		goto L3
	} else {
		goto L4
	}
L3:
	;
	return int32(0)
L4:
	;
	v25 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+48)))
	if v25 == int32(0) {
		goto L5
	} else {
		goto L6
	}
L5:
	;
	v28 = *(*int32)(unsafe.Add(mBase, uint32(l0)+40))
	v29 = F_pg_detoast_datum_packed(m, v28)
	mBase = m.M
	v30 = m.ExcPending
	if v30 != 0 {
		goto L3
	} else {
		goto L8
	}
L6:
	;
	v31 = v3
	goto L7
L7:
	;
	v32 = int32(*(*int16)(unsafe.Add(mBase, uint32(l0)+18)))
	if v32 < int32(3) {
		v39 = v3
		goto L9
	} else {
		goto L10
	}
L8:
	;
	v31 = v29
	goto L7
L9:
	;
	if v31 != 0 {
		goto L14
	} else {
		goto L15
	}
L10:
	;
	v35 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+64)))
	if v35 != 0 {
		v39 = v3
		goto L9
	} else {
		goto L11
	}
L11:
	;
	v36 = *(*int32)(unsafe.Add(mBase, uint32(l0)+56))
	v37 = F_pg_detoast_datum_packed(m, v36)
	mBase = m.M
	v38 = m.ExcPending
	if v38 != 0 {
		goto L3
	} else {
		goto L12
	}
L12:
	;
	v39 = v37
	goto L9
L13:
	;
	v337 = *(*int32)(unsafe.Add(mBase, uint32(l1)))
	v341 = *(*int32)(unsafe.Add(mBase, _c_F_split_text[0]))
	v342 = F_accumArrayResult(m, v337, base.I64_extend_i32_u(v21), v112, int32(25), v341)
	mBase = m.M
	v343 = m.ExcPending
	if v343 != 0 {
		goto L3
	} else {
		goto L130
	}
L14:
	;
	v40 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v21))))
	if v40 == int32(1) {
		goto L18
	} else {
		goto L19
	}
L15:
	;
	goto L16
L16:
	;
	v228 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v21))))
	if v228 == int32(1) {
		goto L95
	} else {
		goto L96
	}
L17:
	;
	v70 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v31))))
	if v70 == int32(1) {
		goto L29
	} else {
		goto L30
	}
L18:
	;
	v46 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v21)+1)))
	if v46 == int32(18) {
		goto L21
	} else {
		goto L22
	}
L19:
	;
	goto L20
L20:
	;
	v57 = int32(1)
	if v40&v57 != 0 {
		v69 = int32(base.Ui32(v40)>>(uint(v57)%32)) - v57
		goto L17
	} else {
		goto L27
	}
L21:
	;
	v49 = int32(16)
	goto L23
L22:
	;
	v49 = int32(0)
	goto L23
L23:
	;
	if base.Ui32((v46-int32(1))&int32(255)) < base.Ui32(int32(3)) {
		goto L24
	} else {
		goto L25
	}
L24:
	;
	v56 = int32(4)
	goto L26
L25:
	;
	v56 = v49
	goto L26
L26:
	;
	v69 = v56
	goto L17
L27:
	;
	v63 = *(*int32)(unsafe.Add(mBase, uint32(v21)))
	v69 = int32(base.Ui32(v63)>>(uint(int32(2))%32)) - int32(4)
	goto L17
L28:
	;
	if v69 <= int32(0) {
		goto L1
	} else {
		goto L39
	}
L29:
	;
	v76 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v31)+1)))
	if v76 == int32(18) {
		goto L32
	} else {
		goto L33
	}
L30:
	;
	goto L31
L31:
	;
	v87 = int32(1)
	if v70&v87 != 0 {
		v99 = int32(base.Ui32(v70)>>(uint(v87)%32)) - v87
		goto L28
	} else {
		goto L38
	}
L32:
	;
	v79 = int32(16)
	goto L34
L33:
	;
	v79 = int32(0)
	goto L34
L34:
	;
	if base.Ui32((v76-int32(1))&int32(255)) < base.Ui32(int32(3)) {
		goto L35
	} else {
		goto L36
	}
L35:
	;
	v86 = int32(4)
	goto L37
L36:
	;
	v86 = v79
	goto L37
L37:
	;
	v99 = v86
	goto L28
L38:
	;
	v93 = *(*int32)(unsafe.Add(mBase, uint32(v31)))
	v99 = int32(base.Ui32(v93)>>(uint(int32(2))%32)) - int32(4)
	goto L28
L39:
	;
	v102 = int32(0)
	if v99 <= v102 {
		goto L40
	} else {
		goto L41
	}
L40:
	;
	if v39 != 0 {
		goto L43
	} else {
		goto L44
	}
L41:
	;
	goto L42
L42:
	;
	F_text_position_setup(m, v21, v31, v19, v16)
	mBase = m.M
	v125 = m.ExcPending
	if v125 != 0 {
		goto L3
	} else {
		goto L49
	}
L43:
	;
	v108 = F_DirectFunctionCall2Coll(m, int32(1757), v19, base.I64_extend_i32_u(v21), base.I64_extend_i32_u(v39))
	mBase = m.M
	v109 = m.ExcPending
	if v109 != 0 {
		goto L3
	} else {
		goto L46
	}
L44:
	;
	v112 = v102
	goto L45
L45:
	;
	v113 = *(*int32)(unsafe.Add(mBase, uint32(l1)+4))
	if v113 == int32(0) {
		goto L13
	} else {
		goto L47
	}
L46:
	;
	v112 = base.B2i32(v108 != int64(0))
	goto L45
L47:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v16))) = base.I64_extend_i32_u(v21)
	*(*uint8)(unsafe.Add(mBase, uint32(v16)+1080)) = uint8(v112)
	v119 = *(*int32)(unsafe.Add(mBase, uint32(l1)+8))
	F_tuplestore_putvalues(m, v113, v119, v16, v16+int32(1080))
	mBase = m.M
	v123 = m.ExcPending
	if v123 != 0 {
		goto L3
	} else {
		goto L48
	}
L48:
	;
	goto L1
L49:
	;
	v126 = int32(1)
	v128 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v21))))
	if v128&v126 != 0 {
		goto L50
	} else {
		goto L51
	}
L50:
	;
	v131 = v126
	goto L52
L51:
	;
	v131 = int32(4)
	goto L52
L52:
	;
	v137 = v21 + v131
	goto L53
L53:
	;
	v148 = *(*int32)(unsafe.Add(mBase, _c_F_split_text[1]))
	if v148 != 0 {
		goto L55
	} else {
		goto L56
	}
L54:
	;
	goto L1
L55:
	;
	F_ProcessInterrupts(m)
	mBase = m.M
	v150 = m.ExcPending
	if v150 != 0 {
		goto L3
	} else {
		goto L58
	}
L56:
	;
	goto L57
L57:
	;
	v151 = F_text_position_next(m, v16)
	mBase = m.M
	v152 = m.ExcPending
	if v152 != 0 {
		goto L3
	} else {
		goto L60
	}
L58:
	;
	goto L57
L59:
	;
	v186 = v184 - v137
	v188 = v186 + int32(4)
	v189 = F_palloc(m, v188)
	mBase = m.M
	v190 = m.ExcPending
	if v190 != 0 {
		goto L3
	} else {
		goto L75
	}
L60:
	;
	if v151 == int32(0) {
		goto L61
	} else {
		goto L62
	}
L61:
	;
	v155 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v21))))
	if v155 == int32(1) {
		goto L65
	} else {
		goto L66
	}
L62:
	;
	goto L63
L63:
	;
	v183 = *(*int32)(unsafe.Add(mBase, uint32(v16)+1052))
	v184 = v183
	v185 = v183
	goto L59
L64:
	;
	v184 = v180 + v21
	v185 = int32(0)
	goto L59
L65:
	;
	v159 = int32(18)
	v161 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v21)+1)))
	if v161 == v159 {
		goto L68
	} else {
		goto L69
	}
L66:
	;
	goto L67
L67:
	;
	v172 = int32(1)
	if v155&v172 != 0 {
		v180 = int32(base.Ui32(v155) >> (uint(v172) % 32))
		goto L64
	} else {
		goto L74
	}
L68:
	;
	v164 = v159
	goto L70
L69:
	;
	v164 = int32(2)
	goto L70
L70:
	;
	if base.Ui32((v161-int32(1))&int32(255)) < base.Ui32(int32(3)) {
		goto L71
	} else {
		goto L72
	}
L71:
	;
	v171 = int32(6)
	goto L73
L72:
	;
	v171 = v164
	goto L73
L73:
	;
	v180 = v171
	goto L64
L74:
	;
	v176 = *(*int32)(unsafe.Add(mBase, uint32(v21)))
	v180 = int32(base.Ui32(v176) >> (uint(int32(2)) % 32))
	goto L64
L75:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v189))) = v188 << (uint(int32(2)) % 32)
	if v186 != 0 {
		goto L76
	} else {
		goto L77
	}
L76:
	;
	base.MemoryCopy(m, v189+int32(4), v137, v186)
	goto L78
L77:
	;
	goto L78
L78:
	;
	if v39 != 0 {
		goto L79
	} else {
		goto L80
	}
L79:
	;
	v199 = F_DirectFunctionCall2Coll(m, int32(1757), v19, base.I64_extend_i32_u(v189), base.I64_extend_i32_u(v39))
	mBase = m.M
	v200 = m.ExcPending
	if v200 != 0 {
		goto L3
	} else {
		goto L82
	}
L80:
	;
	v204 = int32(0)
	goto L81
L81:
	;
	v205 = *(*int32)(unsafe.Add(mBase, uint32(l1)+4))
	if v205 != 0 {
		goto L84
	} else {
		goto L85
	}
L82:
	;
	v204 = base.B2i32(v199 != int64(0))
	goto L81
L83:
	;
	F_pfree(m, v189)
	mBase = m.M
	v225 = m.ExcPending
	if v225 != 0 {
		goto L3
	} else {
		goto L89
	}
L84:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v16)+1080)) = base.I64_extend_i32_u(v189)
	*(*uint8)(unsafe.Add(mBase, uint32(v16)+1079)) = uint8(v204)
	v209 = *(*int32)(unsafe.Add(mBase, uint32(l1)+8))
	F_tuplestore_putvalues(m, v205, v209, v16+int32(1080), v16+int32(1079))
	mBase = m.M
	v215 = m.ExcPending
	if v215 != 0 {
		goto L3
	} else {
		goto L87
	}
L85:
	;
	goto L86
L86:
	;
	v216 = *(*int32)(unsafe.Add(mBase, uint32(l1)))
	v220 = *(*int32)(unsafe.Add(mBase, _c_F_split_text[0]))
	v221 = F_accumArrayResult(m, v216, base.I64_extend_i32_u(v189), v204, int32(25), v220)
	mBase = m.M
	v222 = m.ExcPending
	if v222 != 0 {
		goto L3
	} else {
		goto L88
	}
L87:
	;
	goto L83
L88:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l1))) = v221
	goto L83
L89:
	;
	v226 = *(*int32)(unsafe.Add(mBase, uint32(v16)+1056))
	if v151 != 0 {
		v137 = v185 + v226
		goto L53
	} else {
		goto L90
	}
L90:
	;
	goto L54
L91:
	;
	v277 = v272
	v282 = v274
	goto L107
L92:
	;
	if v262 <= int32(0) {
		goto L1
	} else {
		goto L103
	}
L93:
	;
	v256 = *(*int32)(unsafe.Add(mBase, uint32(v21)))
	v262 = int32(base.Ui32(v256)>>(uint(int32(2))%32)) - int32(4)
	goto L92
L94:
	;
	v272 = v21 + int32(1)
	v274 = int32(4)
	v275 = v21 + int32(5)
	goto L91
L95:
	;
	v231 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v21)+1)))
	if base.Ui32((v231-int32(1))&int32(255)) < base.Ui32(int32(3)) {
		goto L94
	} else {
		goto L98
	}
L96:
	;
	goto L97
L97:
	;
	if v228&int32(1) == int32(0) {
		goto L93
	} else {
		goto L102
	}
L98:
	;
	if v231 == int32(18) {
		goto L99
	} else {
		goto L100
	}
L99:
	;
	v242 = int32(16)
	goto L101
L100:
	;
	v242 = int32(0)
	goto L101
L101:
	;
	v262 = v242
	goto L92
L102:
	;
	v247 = int32(1)
	v262 = int32(base.Ui32(v228)>>(uint(v247)%32)) - v247
	goto L92
L103:
	;
	v265 = int32(1)
	if v228&v265 != 0 {
		goto L104
	} else {
		goto L105
	}
L104:
	;
	v269 = v265
	goto L106
L105:
	;
	v269 = int32(4)
	goto L106
L106:
	;
	v270 = v21 + v269
	v272 = v270
	v274 = v262
	v275 = v270 + v262
	goto L91
L107:
	;
	v290 = F_pg_mblen_range(m, v277, v275)
	mBase = m.M
	v291 = m.ExcPending
	if v291 != 0 {
		goto L3
	} else {
		goto L109
	}
L108:
	;
	goto L1
L109:
	;
	v293 = *(*int32)(unsafe.Add(mBase, _c_F_split_text[1]))
	if v293 != 0 {
		goto L110
	} else {
		goto L111
	}
L110:
	;
	F_ProcessInterrupts(m)
	mBase = m.M
	v295 = m.ExcPending
	if v295 != 0 {
		goto L3
	} else {
		goto L113
	}
L111:
	;
	goto L112
L112:
	;
	v297 = v290 + int32(4)
	v298 = F_palloc(m, v297)
	mBase = m.M
	v299 = m.ExcPending
	if v299 != 0 {
		goto L3
	} else {
		goto L114
	}
L113:
	;
	goto L112
L114:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v298))) = v297 << (uint(int32(2)) % 32)
	if v290 != 0 {
		goto L115
	} else {
		goto L116
	}
L115:
	;
	base.MemoryCopy(m, v298+int32(4), v277, v290)
	goto L117
L116:
	;
	goto L117
L117:
	;
	if v39 != 0 {
		goto L118
	} else {
		goto L119
	}
L118:
	;
	v308 = F_DirectFunctionCall2Coll(m, int32(1757), v19, base.I64_extend_i32_u(v298), base.I64_extend_i32_u(v39))
	mBase = m.M
	v309 = m.ExcPending
	if v309 != 0 {
		goto L3
	} else {
		goto L121
	}
L119:
	;
	v313 = int32(0)
	goto L120
L120:
	;
	v314 = *(*int32)(unsafe.Add(mBase, uint32(l1)+4))
	if v314 != 0 {
		goto L123
	} else {
		goto L124
	}
L121:
	;
	v313 = base.B2i32(v308 != int64(0))
	goto L120
L122:
	;
	F_pfree(m, v298)
	mBase = m.M
	v332 = m.ExcPending
	if v332 != 0 {
		goto L3
	} else {
		goto L128
	}
L123:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v16))) = base.I64_extend_i32_u(v298)
	*(*uint8)(unsafe.Add(mBase, uint32(v16)+1080)) = uint8(v313)
	v318 = *(*int32)(unsafe.Add(mBase, uint32(l1)+8))
	F_tuplestore_putvalues(m, v314, v318, v16, v16+int32(1080))
	mBase = m.M
	v322 = m.ExcPending
	if v322 != 0 {
		goto L3
	} else {
		goto L126
	}
L124:
	;
	goto L125
L125:
	;
	v323 = *(*int32)(unsafe.Add(mBase, uint32(l1)))
	v327 = *(*int32)(unsafe.Add(mBase, _c_F_split_text[0]))
	v328 = F_accumArrayResult(m, v323, base.I64_extend_i32_u(v298), v313, int32(25), v327)
	mBase = m.M
	v329 = m.ExcPending
	if v329 != 0 {
		goto L3
	} else {
		goto L127
	}
L126:
	;
	goto L122
L127:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l1))) = v328
	goto L122
L128:
	;
	v334 = v282 - v290
	if int32(0) < v334 {
		v277 = v277 + v290
		v282 = v334
		goto L107
	} else {
		goto L129
	}
L129:
	;
	goto L108
L130:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l1))) = v342
	goto L1
}
func F_text_char(m *base.Module, l0 int32) int64 {
	mBase := m.M
	_ = mBase
	var v7 int32
	_ = v7
	var v8 int32
	_ = v8
	var v11 int32
	_ = v11
	var v12 int32
	_ = v12
	var v14 int32
	_ = v14
	var v16 int32
	_ = v16
	var v17 int32
	_ = v17
	var v18 int32
	_ = v18
	var v21 int32
	_ = v21
	var v31 int32
	_ = v31
	var v35 int32
	_ = v35
	var v40 int32
	_ = v40
	var v44 int32
	_ = v44
	var v47 int32
	_ = v47
	var v52 int32
	_ = v52
	var v57 int32
	_ = v57
	var v77 int32
	_ = v77
	var v82 int32
	_ = v82
	var v86 int32
	_ = v86
	var v97 int32
	_ = v97
	var v98 int32
	_ = v98
	var v107 int32
	_ = v107
	var v113 int32
	_ = v113
	var v118 int32
	_ = v118
	v7 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
	v8 = F_pg_detoast_datum_packed(m, v7)
	mBase = m.M
	v11 = m.ExcPending
	if v11 != 0 {
		return int64(0)
	} else {
		v12 = int32(1)
		v14 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v8))))
		v16 = v14 & v12
		if v16 != 0 {
			v17 = v12
		} else {
			v17 = int32(4)
		}
		v18 = v8 + v17
		if v14 == int32(1) {
			v21 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v8)+1)))
			if base.Ui32((v21-int32(1))&int32(255)) < base.Ui32(int32(3)) {
				v44 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v18))))
				if v44 != int32(92) {
					if v14 == int32(1) {
						v86 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v8)+1)))
						if base.Ui32((v86-int32(1))&int32(255)) < base.Ui32(int32(3)) {
							v113 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v18))))
							v118 = v113
						} else {
							if v86 == int32(18) {
								v97 = int32(16)
							} else {
								v97 = int32(0)
							}
							v107 = v97
							if v107 != 0 {
								v113 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v18))))
								v118 = v113
							} else {
								v118 = int32(0)
							}
						}
					} else {
						v77 = v14 & int32(1)
						if v77 == int32(0) {
							v98 = *(*int32)(unsafe.Add(mBase, uint32(v8)))
							v107 = int32(base.Ui32(v98)>>(uint(int32(2))%32)) - int32(4)
						} else {
							v82 = int32(1)
							v107 = int32(base.Ui32(v14)>>(uint(v82)%32)) - v82
						}
						if v107 != 0 {
							v113 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v18))))
							v118 = v113
						} else {
							v118 = int32(0)
						}
					}
				} else {
					v47 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v18)+1)))
					if v47&int32(248) != int32(48) {
						if v14 == int32(1) {
							v86 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v8)+1)))
							if base.Ui32((v86-int32(1))&int32(255)) < base.Ui32(int32(3)) {
								v113 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v18))))
								v118 = v113
							} else {
								if v86 == int32(18) {
									v97 = int32(16)
								} else {
									v97 = int32(0)
								}
								v107 = v97
								if v107 != 0 {
									v113 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v18))))
									v118 = v113
								} else {
									v118 = int32(0)
								}
							}
						} else {
							v77 = v14 & int32(1)
							if v77 == int32(0) {
								v98 = *(*int32)(unsafe.Add(mBase, uint32(v8)))
								v107 = int32(base.Ui32(v98)>>(uint(int32(2))%32)) - int32(4)
							} else {
								v82 = int32(1)
								v107 = int32(base.Ui32(v14)>>(uint(v82)%32)) - v82
							}
							if v107 != 0 {
								v113 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v18))))
								v118 = v113
							} else {
								v118 = int32(0)
							}
						}
					} else {
						v52 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v18)+2)))
						if v52&int32(248) != int32(48) {
							if v14 == int32(1) {
								v86 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v8)+1)))
								if base.Ui32((v86-int32(1))&int32(255)) < base.Ui32(int32(3)) {
									v113 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v18))))
									v118 = v113
								} else {
									if v86 == int32(18) {
										v97 = int32(16)
									} else {
										v97 = int32(0)
									}
									v107 = v97
									if v107 != 0 {
										v113 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v18))))
										v118 = v113
									} else {
										v118 = int32(0)
									}
								}
							} else {
								v77 = v14 & int32(1)
								if v77 == int32(0) {
									v98 = *(*int32)(unsafe.Add(mBase, uint32(v8)))
									v107 = int32(base.Ui32(v98)>>(uint(int32(2))%32)) - int32(4)
								} else {
									v82 = int32(1)
									v107 = int32(base.Ui32(v14)>>(uint(v82)%32)) - v82
								}
								if v107 != 0 {
									v113 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v18))))
									v118 = v113
								} else {
									v118 = int32(0)
								}
							}
						} else {
							v57 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v18)+3)))
							if v57&int32(248) != int32(48) {
								if v14 == int32(1) {
									v86 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v8)+1)))
									if base.Ui32((v86-int32(1))&int32(255)) < base.Ui32(int32(3)) {
										v113 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v18))))
										v118 = v113
									} else {
										if v86 == int32(18) {
											v97 = int32(16)
										} else {
											v97 = int32(0)
										}
										v107 = v97
										if v107 != 0 {
											v113 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v18))))
											v118 = v113
										} else {
											v118 = int32(0)
										}
									}
								} else {
									v77 = v14 & int32(1)
									if v77 == int32(0) {
										v98 = *(*int32)(unsafe.Add(mBase, uint32(v8)))
										v107 = int32(base.Ui32(v98)>>(uint(int32(2))%32)) - int32(4)
									} else {
										v82 = int32(1)
										v107 = int32(base.Ui32(v14)>>(uint(v82)%32)) - v82
									}
									if v107 != 0 {
										v113 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v18))))
										v118 = v113
									} else {
										v118 = int32(0)
									}
								}
							} else {
								v118 = v47<<(uint(int32(6))%32) + v52<<(uint(int32(3))%32) + v57 + int32(80)
							}
						}
					}
				}
			} else {
				if v21 == int32(18) {
					v113 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v18))))
					v118 = v113
				} else {
					v118 = int32(0)
				}
			}
		} else {
			if v16 != 0 {
				v31 = int32(1)
				v40 = int32(base.Ui32(v14)>>(uint(v31)%32)) - v31
			} else {
				v35 = *(*int32)(unsafe.Add(mBase, uint32(v8)))
				v40 = int32(base.Ui32(v35)>>(uint(int32(2))%32)) - int32(4)
			}
			if v40 != int32(4) {
				v77 = v16
				if v77 == int32(0) {
					v98 = *(*int32)(unsafe.Add(mBase, uint32(v8)))
					v107 = int32(base.Ui32(v98)>>(uint(int32(2))%32)) - int32(4)
				} else {
					v82 = int32(1)
					v107 = int32(base.Ui32(v14)>>(uint(v82)%32)) - v82
				}
				if v107 != 0 {
					v113 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v18))))
					v118 = v113
				} else {
					v118 = int32(0)
				}
			} else {
				v44 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v18))))
				if v44 != int32(92) {
					if v14 == int32(1) {
						v86 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v8)+1)))
						if base.Ui32((v86-int32(1))&int32(255)) < base.Ui32(int32(3)) {
							v113 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v18))))
							v118 = v113
						} else {
							if v86 == int32(18) {
								v97 = int32(16)
							} else {
								v97 = int32(0)
							}
							v107 = v97
							if v107 != 0 {
								v113 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v18))))
								v118 = v113
							} else {
								v118 = int32(0)
							}
						}
					} else {
						v77 = v14 & int32(1)
						if v77 == int32(0) {
							v98 = *(*int32)(unsafe.Add(mBase, uint32(v8)))
							v107 = int32(base.Ui32(v98)>>(uint(int32(2))%32)) - int32(4)
						} else {
							v82 = int32(1)
							v107 = int32(base.Ui32(v14)>>(uint(v82)%32)) - v82
						}
						if v107 != 0 {
							v113 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v18))))
							v118 = v113
						} else {
							v118 = int32(0)
						}
					}
				} else {
					v47 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v18)+1)))
					if v47&int32(248) != int32(48) {
						if v14 == int32(1) {
							v86 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v8)+1)))
							if base.Ui32((v86-int32(1))&int32(255)) < base.Ui32(int32(3)) {
								v113 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v18))))
								v118 = v113
							} else {
								if v86 == int32(18) {
									v97 = int32(16)
								} else {
									v97 = int32(0)
								}
								v107 = v97
								if v107 != 0 {
									v113 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v18))))
									v118 = v113
								} else {
									v118 = int32(0)
								}
							}
						} else {
							v77 = v14 & int32(1)
							if v77 == int32(0) {
								v98 = *(*int32)(unsafe.Add(mBase, uint32(v8)))
								v107 = int32(base.Ui32(v98)>>(uint(int32(2))%32)) - int32(4)
							} else {
								v82 = int32(1)
								v107 = int32(base.Ui32(v14)>>(uint(v82)%32)) - v82
							}
							if v107 != 0 {
								v113 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v18))))
								v118 = v113
							} else {
								v118 = int32(0)
							}
						}
					} else {
						v52 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v18)+2)))
						if v52&int32(248) != int32(48) {
							if v14 == int32(1) {
								v86 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v8)+1)))
								if base.Ui32((v86-int32(1))&int32(255)) < base.Ui32(int32(3)) {
									v113 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v18))))
									v118 = v113
								} else {
									if v86 == int32(18) {
										v97 = int32(16)
									} else {
										v97 = int32(0)
									}
									v107 = v97
									if v107 != 0 {
										v113 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v18))))
										v118 = v113
									} else {
										v118 = int32(0)
									}
								}
							} else {
								v77 = v14 & int32(1)
								if v77 == int32(0) {
									v98 = *(*int32)(unsafe.Add(mBase, uint32(v8)))
									v107 = int32(base.Ui32(v98)>>(uint(int32(2))%32)) - int32(4)
								} else {
									v82 = int32(1)
									v107 = int32(base.Ui32(v14)>>(uint(v82)%32)) - v82
								}
								if v107 != 0 {
									v113 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v18))))
									v118 = v113
								} else {
									v118 = int32(0)
								}
							}
						} else {
							v57 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v18)+3)))
							if v57&int32(248) != int32(48) {
								if v14 == int32(1) {
									v86 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v8)+1)))
									if base.Ui32((v86-int32(1))&int32(255)) < base.Ui32(int32(3)) {
										v113 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v18))))
										v118 = v113
									} else {
										if v86 == int32(18) {
											v97 = int32(16)
										} else {
											v97 = int32(0)
										}
										v107 = v97
										if v107 != 0 {
											v113 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v18))))
											v118 = v113
										} else {
											v118 = int32(0)
										}
									}
								} else {
									v77 = v14 & int32(1)
									if v77 == int32(0) {
										v98 = *(*int32)(unsafe.Add(mBase, uint32(v8)))
										v107 = int32(base.Ui32(v98)>>(uint(int32(2))%32)) - int32(4)
									} else {
										v82 = int32(1)
										v107 = int32(base.Ui32(v14)>>(uint(v82)%32)) - v82
									}
									if v107 != 0 {
										v113 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v18))))
										v118 = v113
									} else {
										v118 = int32(0)
									}
								}
							} else {
								v118 = v47<<(uint(int32(6))%32) + v52<<(uint(int32(3))%32) + v57 + int32(80)
							}
						}
					}
				}
			}
		}
		return base.I64_extend8_s(base.I64_extend_i32_u(v118))
	}
}
func F_text_format_nv(m *base.Module, l0 int32) int64 {
	var v2 int64
	_ = v2
	var v5 int32
	_ = v5
	v2 = F_text_format(m, l0)
	v5 = m.ExcPending
	if v5 != 0 {
		return int64(0)
	} else {
		return v2
	}
}
func F_text_format_parse_digits(m *base.Module, l0 int32, l1 int32, l2 int32) int32 {
	mBase := m.M
	_ = mBase
	var v4 int32
	_ = v4
	var v10 int32
	_ = v10
	var v11 int32
	_ = v11
	var v17 int32
	_ = v17
	var v23 int32
	_ = v23
	var v24 int32
	_ = v24
	var v25 int32
	_ = v25
	var v31 int64
	_ = v31
	var v35 int32
	_ = v35
	var v43 int32
	_ = v43
	var v46 int32
	_ = v46
	var v48 int32
	_ = v48
	var v60 int32
	_ = v60
	var v63 int32
	_ = v63
	var v67 int32
	_ = v67
	var v71 int32
	_ = v71
	var v76 int32
	_ = v76
	var v80 int32
	_ = v80
	var v82 int32
	_ = v82
	var v93 int32
	_ = v93
	var v96 int32
	_ = v96
	var v100 int32
	_ = v100
	var v105 int32
	_ = v105
	v4 = int32(0)
	v10 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v11 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v10))))
	v17 = base.B2i32(base.Ui32((v11-int32(48))&int32(255)) < base.Ui32(int32(10)))
	if v17 == v4 {
		v80 = v10
		v82 = v4
		goto L2
	} else {
		goto L3
	}
L1:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v93 = m.ExcPending
	if v93 != 0 {
		goto L12
	} else {
		goto L18
	}
L2:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0))) = v80
	*(*int32)(unsafe.Add(mBase, uint32(l2))) = v82
	return v17
L3:
	;
	v23 = v10
	v24 = v11
	v25 = v4
	goto L4
L4:
	;
	v31 = base.I64_extend_i32_s(v25) * int64(10)
	v35 = base.I32_wrap_i64(v31)
	if base.I32_wrap_i64(int64(base.Ui64(v31)>>(uint(int64(32))%64))) != v35>>(uint(int32(31))%32) {
		goto L1
	} else {
		goto L6
	}
L5:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v60 = m.ExcPending
	if v60 != 0 {
		goto L12
	} else {
		goto L13
	}
L6:
	;
	v43 = (v24-int32(48))&int32(255) + v35
	if v43 < v35 {
		goto L1
	} else {
		goto L7
	}
L7:
	;
	v46 = v23 + int32(1)
	if base.Ui32(v46) < base.Ui32(l1) {
		goto L8
	} else {
		goto L9
	}
L8:
	;
	v48 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v46))))
	if base.Ui32(int32(9)) < base.Ui32((v48-int32(48))&int32(255)) {
		v80 = v46
		v82 = v43
		goto L2
	} else {
		goto L11
	}
L9:
	;
	goto L10
L10:
	;
	goto L5
L11:
	;
	v23 = v46
	v24 = v48
	v25 = v43
	goto L4
L12:
	;
	return int32(0)
L13:
	;
	F_errcode(m, int32(50856066))
	mBase = m.M
	v63 = m.ExcPending
	if v63 != 0 {
		goto L12
	} else {
		goto L14
	}
L14:
	;
	F_errmsg(m, int32(_a_F_text_format_parse_digits_0), int32(0))
	mBase = m.M
	v67 = m.ExcPending
	if v67 != 0 {
		goto L12
	} else {
		goto L15
	}
L15:
	;
	F_errhint(m, int32(_a_F_text_format_parse_digits_1), int32(0))
	mBase = m.M
	v71 = m.ExcPending
	if v71 != 0 {
		goto L12
	} else {
		goto L16
	}
L16:
	;
	F_errfinish(m, int32(_a_F_text_format_parse_digits_2), int32(_a_F_text_format_parse_digits_3), int32(_a_F_text_format_parse_digits_4))
	mBase = m.M
	v76 = m.ExcPending
	if v76 != 0 {
		goto L12
	} else {
		goto L17
	}
L17:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L18:
	;
	F_errcode(m, int32(50331778))
	mBase = m.M
	v96 = m.ExcPending
	if v96 != 0 {
		goto L12
	} else {
		goto L19
	}
L19:
	;
	F_errmsg(m, int32(_a_F_text_format_parse_digits_5), int32(0))
	mBase = m.M
	v100 = m.ExcPending
	if v100 != 0 {
		goto L12
	} else {
		goto L20
	}
L20:
	;
	F_errfinish(m, int32(_a_F_text_format_parse_digits_2), int32(_a_F_text_format_parse_digits_6), int32(_a_F_text_format_parse_digits_4))
	mBase = m.M
	v105 = m.ExcPending
	if v105 != 0 {
		goto L12
	} else {
		goto L21
	}
L21:
	;
	base.Wasm_trap_unreachable()
	for {
	}
}
func F_text_length(m *base.Module, l0 int64) int32 {
	mBase := m.M
	_ = mBase
	var v6 int32
	_ = v6
	var v7 int32
	_ = v7
	var v12 int32
	_ = v12
	var v15 int32
	_ = v15
	var v18 int32
	_ = v18
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
	var v31 int32
	_ = v31
	var v37 int32
	_ = v37
	var v40 int32
	_ = v40
	var v47 int32
	_ = v47
	var v48 int32
	_ = v48
	var v49 int32
	_ = v49
	var v53 int32
	_ = v53
	var v57 int32
	_ = v57
	var v58 int32
	_ = v58
	var v60 int32
	_ = v60
	var v65 int32
	_ = v65
	var v66 int32
	_ = v66
	v6 = *(*int32)(unsafe.Add(mBase, _c_F_text_length[0]))
	v7 = *(*int32)(unsafe.Add(mBase, uint32(v6)+4))
	v12 = *(*int32)(unsafe.Add(mBase, uint32(v7*int32(28))+uint32(_c_F_text_length[1])))
	if v12 == int32(1) {
		v15 = F_toast_raw_datum_size(m, l0)
		mBase = m.M
		v18 = m.ExcPending
		if v18 != 0 {
			return int32(0)
		} else {
			return v15 - int32(4)
		}
	} else {
		v23 = F_pg_detoast_datum_packed(m, base.I32_wrap_i64(l0))
		mBase = m.M
		v24 = m.ExcPending
		if v24 != 0 {
			return int32(0)
		} else {
			v25 = int32(1)
			v27 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v23))))
			if v27&v25 != 0 {
				v30 = v25
			} else {
				v30 = int32(4)
			}
			v31 = v23 + v30
			if v27 == int32(1) {
				v37 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v23)+1)))
				if v37 == int32(18) {
					v40 = int32(16)
				} else {
					v40 = int32(0)
				}
				if base.Ui32((v37-int32(1))&int32(255)) < base.Ui32(int32(3)) {
					v47 = int32(4)
				} else {
					v47 = v40
				}
				v48 = F_pg_mbstrlen_with_len(m, v31, v47)
				mBase = m.M
				v49 = m.ExcPending
				if v49 != 0 {
					return int32(0)
				} else {
					return v48
				}
			} else {
				if v27&int32(1) != 0 {
					v53 = int32(1)
					v57 = F_pg_mbstrlen_with_len(m, v31, int32(base.Ui32(v27)>>(uint(v53)%32))-v53)
					mBase = m.M
					v58 = m.ExcPending
					if v58 != 0 {
						return int32(0)
					} else {
						return v57
					}
				} else {
					v60 = *(*int32)(unsafe.Add(mBase, uint32(v23)))
					v65 = F_pg_mbstrlen_with_len(m, v31, int32(base.Ui32(v60)>>(uint(int32(2))%32))-int32(4))
					mBase = m.M
					v66 = m.ExcPending
					if v66 != 0 {
						return int32(0)
					} else {
						return v65
					}
				}
			}
		}
	}
}
func F_text_overlay(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) int32 {
	var v10 int32
	_ = v10
	var v13 int64
	_ = v13
	var v14 int32
	_ = v14
	var v18 int32
	_ = v18
	var v21 int32
	_ = v21
	var v24 int32
	_ = v24
	var v25 int32
	_ = v25
	var v26 int32
	_ = v26
	var v27 int32
	_ = v27
	var v28 int32
	_ = v28
	var v29 int32
	_ = v29
	var v34 int32
	_ = v34
	var v37 int32
	_ = v37
	var v41 int32
	_ = v41
	var v46 int32
	_ = v46
	var v50 int32
	_ = v50
	var v53 int32
	_ = v53
	var v57 int32
	_ = v57
	var v62 int32
	_ = v62
	if int32(0) < l2 {
		v10 = l2 + l3
		if base.B2i32(l3 < int32(0)) != base.B2i32(v10 < l2) {
			F_errstart_cold(m, int32(21), int32(0))
			v50 = m.ExcPending
			if v50 != 0 {
				return int32(0)
			} else {
				F_errcode(m, int32(50331778))
				v53 = m.ExcPending
				if v53 != 0 {
					return int32(0)
				} else {
					F_errmsg(m, int32(_a_F_text_overlay_0), int32(0))
					v57 = m.ExcPending
					if v57 != 0 {
						return int32(0)
					} else {
						F_errfinish(m, int32(_a_F_text_overlay_1), int32(875), int32(_a_F_text_overlay_2))
						v62 = m.ExcPending
						if v62 != 0 {
							return int32(0)
						} else {
							base.Wasm_trap_unreachable()
							for {
							}
						}
					}
				}
			}
		} else {
			v13 = base.I64_extend_i32_u(l0)
			v14 = int32(1)
			v18 = F_text_substring(m, v13, v14, l2-v14, int32(0))
			v21 = m.ExcPending
			if v21 != 0 {
				return int32(0)
			} else {
				v24 = F_text_substring(m, v13, v10, int32(-1), int32(1))
				v25 = m.ExcPending
				if v25 != 0 {
					return int32(0)
				} else {
					v26 = F_bytea_catenate(m, v18, l1)
					v27 = m.ExcPending
					if v27 != 0 {
						return int32(0)
					} else {
						v28 = F_bytea_catenate(m, v26, v24)
						v29 = m.ExcPending
						if v29 != 0 {
							return int32(0)
						} else {
							return v28
						}
					}
				}
			}
		}
	} else {
		F_errstart_cold(m, int32(21), int32(0))
		v34 = m.ExcPending
		if v34 != 0 {
			return int32(0)
		} else {
			F_errcode(m, int32(17039490))
			v37 = m.ExcPending
			if v37 != 0 {
				return int32(0)
			} else {
				F_errmsg(m, int32(_a_F_text_overlay_3), int32(0))
				v41 = m.ExcPending
				if v41 != 0 {
					return int32(0)
				} else {
					F_errfinish(m, int32(_a_F_text_overlay_1), int32(871), int32(_a_F_text_overlay_2))
					v46 = m.ExcPending
					if v46 != 0 {
						return int32(0)
					} else {
						base.Wasm_trap_unreachable()
						for {
						}
					}
				}
			}
		}
	}
}
func F_text_pattern_le(m *base.Module, l0 int32) int64 {
	mBase := m.M
	_ = mBase
	var v5 int32
	_ = v5
	var v6 int32
	_ = v6
	var v9 int32
	_ = v9
	var v10 int32
	_ = v10
	var v11 int32
	_ = v11
	var v12 int32
	_ = v12
	var v17 int32
	_ = v17
	var v23 int32
	_ = v23
	var v26 int32
	_ = v26
	var v33 int32
	_ = v33
	var v34 int32
	_ = v34
	var v40 int32
	_ = v40
	var v46 int32
	_ = v46
	var v47 int32
	_ = v47
	var v53 int32
	_ = v53
	var v56 int32
	_ = v56
	var v63 int32
	_ = v63
	var v64 int32
	_ = v64
	var v70 int32
	_ = v70
	var v76 int32
	_ = v76
	var v77 int32
	_ = v77
	var v81 int32
	_ = v81
	var v83 int32
	_ = v83
	var v87 int32
	_ = v87
	var v89 int32
	_ = v89
	var v90 int32
	_ = v90
	var v91 int32
	_ = v91
	var v94 int32
	_ = v94
	var v95 int32
	_ = v95
	var v98 int32
	_ = v98
	var v99 int32
	_ = v99
	var v102 int32
	_ = v102
	v5 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
	v6 = F_pg_detoast_datum_packed(m, v5)
	mBase = m.M
	v9 = m.ExcPending
	if v9 != 0 {
		return int64(0)
	} else {
		v10 = *(*int32)(unsafe.Add(mBase, uint32(l0)+40))
		v11 = F_pg_detoast_datum_packed(m, v10)
		mBase = m.M
		v12 = m.ExcPending
		if v12 != 0 {
			return int64(0)
		} else {
			v17 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v6))))
			if v17 == int32(1) {
				v23 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v6)+1)))
				if v23 == int32(18) {
					v26 = int32(16)
				} else {
					v26 = int32(0)
				}
				if base.Ui32((v23-int32(1))&int32(255)) < base.Ui32(int32(3)) {
					v33 = int32(4)
				} else {
					v33 = v26
				}
				v46 = v33
			} else {
				v34 = int32(1)
				if v17&v34 != 0 {
					v46 = int32(base.Ui32(v17)>>(uint(v34)%32)) - v34
				} else {
					v40 = *(*int32)(unsafe.Add(mBase, uint32(v6)))
					v46 = int32(base.Ui32(v40)>>(uint(int32(2))%32)) - int32(4)
				}
			}
			v47 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v11))))
			if v47 == int32(1) {
				v53 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v11)+1)))
				if v53 == int32(18) {
					v56 = int32(16)
				} else {
					v56 = int32(0)
				}
				if base.Ui32((v53-int32(1))&int32(255)) < base.Ui32(int32(3)) {
					v63 = int32(4)
				} else {
					v63 = v56
				}
				v76 = v63
			} else {
				v64 = int32(1)
				if v47&v64 != 0 {
					v76 = int32(base.Ui32(v47)>>(uint(v64)%32)) - v64
				} else {
					v70 = *(*int32)(unsafe.Add(mBase, uint32(v11)))
					v76 = int32(base.Ui32(v70)>>(uint(int32(2))%32)) - int32(4)
				}
			}
			v77 = int32(1)
			if v17&v77 != 0 {
				v81 = v77
			} else {
				v81 = int32(4)
			}
			v83 = int32(1)
			if v47&v83 != 0 {
				v87 = v83
			} else {
				v87 = int32(4)
			}
			v89 = base.B2i32(v46 < v76)
			if v46 < v76 {
				v90 = v46
			} else {
				v90 = v76
			}
			v91 = F_memcmp(m, v6+v81, v11+v87, v90)
			mBase = m.M
			if v91 != 0 {
				v94 = v91
			} else {
				if v46 < v76 {
					v94 = int32(-1)
				} else {
					v94 = base.B2i32(v76 < v46)
				}
			}
			v95 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
			if v95 != v6 {
				F_pfree(m, v6)
				mBase = m.M
				v98 = m.ExcPending
				if v98 != 0 {
					return int64(0)
				} else {
					v99 = *(*int32)(unsafe.Add(mBase, uint32(l0)+40))
					if v99 != v11 {
						F_pfree(m, v11)
						mBase = m.M
						v102 = m.ExcPending
						if v102 != 0 {
							return int64(0)
						} else {
							return base.I64_extend_i32_u(base.B2i32(v94 <= int32(0)))
						}
					} else {
						return base.I64_extend_i32_u(base.B2i32(v94 <= int32(0)))
					}
				}
			} else {
				v99 = *(*int32)(unsafe.Add(mBase, uint32(l0)+40))
				if v99 != v11 {
					F_pfree(m, v11)
					mBase = m.M
					v102 = m.ExcPending
					if v102 != 0 {
						return int64(0)
					} else {
						return base.I64_extend_i32_u(base.B2i32(v94 <= int32(0)))
					}
				} else {
					return base.I64_extend_i32_u(base.B2i32(v94 <= int32(0)))
				}
			}
		}
	}
}
func F_text_right(m *base.Module, l0 int32) int64 {
	mBase := m.M
	_ = mBase
	var v8 int32
	_ = v8
	var v9 int32
	_ = v9
	var v12 int32
	_ = v12
	var v13 int32
	_ = v13
	var v15 int32
	_ = v15
	var v16 int32
	_ = v16
	var v22 int32
	_ = v22
	var v25 int32
	_ = v25
	var v32 int32
	_ = v32
	var v33 int32
	_ = v33
	var v37 int32
	_ = v37
	var v43 int32
	_ = v43
	var v44 int32
	_ = v44
	var v45 int32
	_ = v45
	var v53 int32
	_ = v53
	var v54 int32
	_ = v54
	var v55 int32
	_ = v55
	var v57 int32
	_ = v57
	var v58 int32
	_ = v58
	var v59 int32
	_ = v59
	var v60 int32
	_ = v60
	var v62 int32
	_ = v62
	var v63 int32
	_ = v63
	var v64 int32
	_ = v64
	v8 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
	v9 = F_pg_detoast_datum_packed(m, v8)
	mBase = m.M
	v12 = m.ExcPending
	if v12 != 0 {
		return int64(0)
	} else {
		v13 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v9))))
		v15 = v13 & int32(1)
		if v15 != 0 {
			v16 = int32(1)
		} else {
			v16 = int32(4)
		}
		if v13 == int32(1) {
			v22 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v9)+1)))
			if v22 == int32(18) {
				v25 = int32(16)
			} else {
				v25 = int32(0)
			}
			if base.Ui32((v22-int32(1))&int32(255)) < base.Ui32(int32(3)) {
				v32 = int32(4)
			} else {
				v32 = v25
			}
			v43 = v32
		} else {
			v33 = int32(1)
			if v15 != 0 {
				v43 = int32(base.Ui32(v13)>>(uint(v33)%32)) - v33
			} else {
				v37 = *(*int32)(unsafe.Add(mBase, uint32(v9)))
				v43 = int32(base.Ui32(v37)>>(uint(int32(2))%32)) - int32(4)
			}
		}
		v44 = v9 + v16
		v45 = *(*int32)(unsafe.Add(mBase, uint32(l0)+40))
		if v45 < int32(0) {
			if v45 == int32(-2147483648) {
				v53 = int32(2147483647)
			} else {
				v53 = int32(0) - v45
			}
			v57 = v53
			v58 = F_pg_mbcharcliplen(m, v44, v43, v57)
			mBase = m.M
			v59 = m.ExcPending
			if v59 != 0 {
				return int64(0)
			} else {
				v60 = v43 - v58
				v62 = v60 + int32(4)
				v63 = F_palloc(m, v62)
				mBase = m.M
				v64 = m.ExcPending
				if v64 != 0 {
					return int64(0)
				} else {
					*(*int32)(unsafe.Add(mBase, uint32(v63))) = v62 << (uint(int32(2)) % 32)
					if v60 != 0 {
						base.MemoryCopy(m, v63+int32(4), v44+v58, v60)
					} else {
					}
					return base.I64_extend_i32_u(v63)
				}
			}
		} else {
			v54 = F_pg_mbstrlen_with_len(m, v44, v43)
			mBase = m.M
			v55 = m.ExcPending
			if v55 != 0 {
				return int64(0)
			} else {
				v57 = v54 - v45
				v58 = F_pg_mbcharcliplen(m, v44, v43, v57)
				mBase = m.M
				v59 = m.ExcPending
				if v59 != 0 {
					return int64(0)
				} else {
					v60 = v43 - v58
					v62 = v60 + int32(4)
					v63 = F_palloc(m, v62)
					mBase = m.M
					v64 = m.ExcPending
					if v64 != 0 {
						return int64(0)
					} else {
						*(*int32)(unsafe.Add(mBase, uint32(v63))) = v62 << (uint(int32(2)) % 32)
						if v60 != 0 {
							base.MemoryCopy(m, v63+int32(4), v44+v58, v60)
						} else {
						}
						return base.I64_extend_i32_u(v63)
					}
				}
			}
		}
	}
}
