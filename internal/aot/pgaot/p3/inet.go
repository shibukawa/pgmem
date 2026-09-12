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
	var v16 int32
	_ = v16
	var v17 int32
	_ = v17
	var v18 int32
	_ = v18
	var v22 int64
	_ = v22
	var v26 int32
	_ = v26
	var v31 int32
	_ = v31
	var v33 int32
	_ = v33
	var v34 int32
	_ = v34
	var v37 int32
	_ = v37
	var v38 int32
	_ = v38
	var v40 int32
	_ = v40
	var v41 int32
	_ = v41
	var v43 int32
	_ = v43
	var v44 int32
	_ = v44
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
	var v57 int32
	_ = v57
	var v58 int32
	_ = v58
	var v60 int32
	_ = v60
	var v61 int32
	_ = v61
	var v98 int32
	_ = v98
	var v108 int32
	_ = v108
	var v110 int32
	_ = v110
	var v114 int32
	_ = v114
	var v115 int32
	_ = v115
	var v118 int32
	_ = v118
	var v122 int32
	_ = v122
	var v123 int32
	_ = v123
	var v125 int32
	_ = v125
	var v128 int32
	_ = v128
	var v130 int32
	_ = v130
	var v135 int32
	_ = v135
	var v137 int32
	_ = v137
	var v142 int32
	_ = v142
	var v144 int32
	_ = v144
	var v147 int32
	_ = v147
	var v149 int32
	_ = v149
	var v152 int32
	_ = v152
	var v164 int32
	_ = v164
	var v167 int32
	_ = v167
	var v169 int32
	_ = v169
	var v206 int32
	_ = v206
	var v216 int32
	_ = v216
	var v218 int32
	_ = v218
	var v222 int32
	_ = v222
	var v223 int32
	_ = v223
	var v226 int32
	_ = v226
	var v230 int32
	_ = v230
	var v231 int32
	_ = v231
	var v233 int32
	_ = v233
	var v236 int32
	_ = v236
	var v238 int32
	_ = v238
	var v243 int32
	_ = v243
	var v245 int32
	_ = v245
	var v250 int32
	_ = v250
	var v252 int32
	_ = v252
	var v255 int32
	_ = v255
	var v257 int32
	_ = v257
	var v260 int32
	_ = v260
	var v272 int32
	_ = v272
	var v275 int32
	_ = v275
	var v276 int32
	_ = v276
	var v277 int32
	_ = v277
	var v279 int32
	_ = v279
	var v287 int32
	_ = v287
	var v295 int32
	_ = v295
	var v298 int32
	_ = v298
	var v302 int32
	_ = v302
	var v304 int32
	_ = v304
	var v310 int32
	_ = v310
	var v312 int32
	_ = v312
	var v318 int32
	_ = v318
	var v319 int32
	_ = v319
	var v320 int32
	_ = v320
	var v321 int32
	_ = v321
	var v322 int32
	_ = v322
	var v323 int32
	_ = v323
	var v324 int32
	_ = v324
	var v327 int32
	_ = v327
	var v336 int32
	_ = v336
	var v339 int32
	_ = v339
	var v344 int32
	_ = v344
	var v346 int32
	_ = v346
	var v348 int32
	_ = v348
	var v356 int32
	_ = v356
	var v361 int32
	_ = v361
	var v365 int32
	_ = v365
	var v367 int32
	_ = v367
	var v381 int32
	_ = v381
	var v383 int32
	_ = v383
	var v386 int32
	_ = v386
	var v388 int32
	_ = v388
	var v389 int32
	_ = v389
	var v390 int32
	_ = v390
	var v393 int32
	_ = v393
	var v405 int32
	_ = v405
	var v415 int32
	_ = v415
	var v417 int32
	_ = v417
	var v421 int32
	_ = v421
	var v422 int32
	_ = v422
	var v425 int32
	_ = v425
	var v429 int32
	_ = v429
	var v430 int32
	_ = v430
	var v432 int32
	_ = v432
	var v435 int32
	_ = v435
	var v437 int32
	_ = v437
	var v442 int32
	_ = v442
	var v444 int32
	_ = v444
	var v449 int32
	_ = v449
	var v451 int32
	_ = v451
	var v454 int32
	_ = v454
	var v456 int32
	_ = v456
	var v459 int32
	_ = v459
	var v471 int32
	_ = v471
	var v473 int32
	_ = v473
	var v480 int32
	_ = v480
	var v495 int32
	_ = v495
	var v497 int32
	_ = v497
	var v499 int32
	_ = v499
	var v507 int32
	_ = v507
	var v512 int32
	_ = v512
	var v515 int32
	_ = v515
	var v516 int32
	_ = v516
	var v553 int32
	_ = v553
	var v563 int32
	_ = v563
	var v565 int32
	_ = v565
	var v569 int32
	_ = v569
	var v570 int32
	_ = v570
	var v573 int32
	_ = v573
	var v577 int32
	_ = v577
	var v578 int32
	_ = v578
	var v580 int32
	_ = v580
	var v583 int32
	_ = v583
	var v585 int32
	_ = v585
	var v590 int32
	_ = v590
	var v592 int32
	_ = v592
	var v597 int32
	_ = v597
	var v599 int32
	_ = v599
	var v602 int32
	_ = v602
	var v604 int32
	_ = v604
	var v607 int32
	_ = v607
	var v619 int32
	_ = v619
	var v623 int32
	_ = v623
	var v624 int32
	_ = v624
	var v632 int32
	_ = v632
	var v642 int32
	_ = v642
	var v644 int32
	_ = v644
	var v656 int32
	_ = v656
	var v657 int32
	_ = v657
	var v658 int32
	_ = v658
	var v659 int32
	_ = v659
	var v696 int32
	_ = v696
	var v706 int32
	_ = v706
	var v708 int32
	_ = v708
	var v712 int32
	_ = v712
	var v713 int32
	_ = v713
	var v716 int32
	_ = v716
	var v720 int32
	_ = v720
	var v721 int32
	_ = v721
	var v723 int32
	_ = v723
	var v726 int32
	_ = v726
	var v728 int32
	_ = v728
	var v733 int32
	_ = v733
	var v735 int32
	_ = v735
	var v740 int32
	_ = v740
	var v742 int32
	_ = v742
	var v745 int32
	_ = v745
	var v747 int32
	_ = v747
	var v750 int32
	_ = v750
	var v762 int32
	_ = v762
	var v766 int32
	_ = v766
	var v767 int32
	_ = v767
	var v775 int32
	_ = v775
	var v782 int32
	_ = v782
	var v785 int32
	_ = v785
	var v786 int32
	_ = v786
	var v788 int32
	_ = v788
	var v792 int32
	_ = v792
	var v798 int32
	_ = v798
	var v800 int32
	_ = v800
	var v806 int32
	_ = v806
	var v813 int32
	_ = v813
	var v814 int32
	_ = v814
	var v819 int32
	_ = v819
	var v829 int32
	_ = v829
	var v833 int32
	_ = v833
	var v838 int32
	_ = v838
	var v845 int32
	_ = v845
	var v847 int32
	_ = v847
	var v848 int32
	_ = v848
	var v851 int32
	_ = v851
	var v862 int32
	_ = v862
	var v863 int32
	_ = v863
	var v877 int32
	_ = v877
	var v878 int32
	_ = v878
	var v880 int32
	_ = v880
	var v883 int32
	_ = v883
	var v885 int32
	_ = v885
	var v886 int32
	_ = v886
	var v890 int32
	_ = v890
	var v891 int32
	_ = v891
	var v893 int32
	_ = v893
	var v897 int32
	_ = v897
	var v914 int32
	_ = v914
	var v915 int32
	_ = v915
	var v917 int32
	_ = v917
	var v969 int64
	_ = v969
	var v971 int64
	_ = v971
	var v973 int32
	_ = v973
	v4 = int32(0)
	v16 = m.G0
	v17 = int32(16)
	v18 = v16 - v17
	m.G0 = v18
	if base.Ui32(v17) <= base.Ui32(l2) {
		goto L7
	} else {
		goto L8
	}
L1:
	;
	m.G0 = v18 + int32(16)
	return v973
L2:
	;
	v969 = *(*int64)(unsafe.Add(mBase, uint32(v18)))
	*(*int64)(unsafe.Add(mBase, uint32(l1))) = v969
	v971 = *(*int64)(unsafe.Add(mBase, uint32(v18)+8))
	*(*int64)(unsafe.Add(mBase, uint32(l1)+8)) = v971
	v973 = v845
	goto L1
L3:
	;
	v973 = int32(-1)
	goto L1
L4:
	;
	*(*int32)(unsafe.Add(mBase, _consts[43])) = int32(44)
	goto L3
L5:
	;
	if v838 != 0 {
		goto L217
	} else {
		goto L218
	}
L6:
	;
	if v813 == int32(-1) {
		goto L214
	} else {
		goto L215
	}
L7:
	;
	v22 = int64(0)
	*(*int64)(unsafe.Add(mBase, uint32(v18)+8)) = v22
	*(*int64)(unsafe.Add(mBase, uint32(v18))) = v22
	v26 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0))))
	if v26 == int32(0) {
		goto L4
	} else {
		goto L10
	}
L8:
	;
	goto L9
L9:
	;
	*(*int32)(unsafe.Add(mBase, _consts[43])) = int32(35)
	goto L3
L10:
	;
	if v26 == int32(58) {
		goto L11
	} else {
		goto L12
	}
L11:
	;
	v31 = int32(58)
	v33 = l0 + int32(1)
	v34 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v33))))
	if v34 != v31 {
		goto L4
	} else {
		goto L14
	}
L12:
	;
	v37 = l0
	v38 = v26
	goto L13
L13:
	;
	v40 = v18 + int32(16)
	v41 = v37
	v43 = v38
	v44 = v4
	v46 = v4
	v47 = v18
	v48 = v4
	v51 = v37
	v53 = v4
	goto L19
L14:
	;
	v37 = v33
	v38 = v31
	goto L13
L15:
	;
	if v497 != 0 {
		v813 = v499
		v814 = v327
		v819 = v44
		goto L6
	} else {
		goto L213
	}
L16:
	;
	if v792 == int32(0) {
		goto L209
	} else {
		goto L210
	}
L17:
	;
	v642 = v57
	v644 = int32(0)
	v656 = v4
	goto L171
L18:
	;
	v327 = v47 + int32(4)
	if base.Ui32(v40) < base.Ui32(v327) {
		goto L4
	} else {
		goto L90
	}
L19:
	;
	v57 = v41 + int32(1)
	v58 = int32(1694848)
	v60 = base.I32_extend8_s(v43)
	v61 = int32(17)
	goto L26
L20:
	;
	v782 = int32(-1)
	v785 = v318
	v786 = v319
	v788 = v323
	v792 = v322
	goto L16
L21:
	;
	v324 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v57))))
	if v324 != 0 {
		v41 = v57
		v43 = v324
		v44 = v323
		v46 = v318
		v47 = v319
		v48 = v320
		v51 = v321
		v53 = v322
		goto L19
	} else {
		goto L89
	}
L22:
	;
	v287 = v43 & int32(255)
	if v287 != int32(58) {
		goto L80
	} else {
		goto L81
	}
L23:
	;
	if v164 == int32(0) {
		goto L49
	} else {
		goto L50
	}
L24:
	;
	v164 = int32(0)
	goto L23
L25:
	;
	v142 = v135
	v144 = v137
	goto L43
L26:
	;
	goto L34
L34:
	;
	v98 = int32(*(*uint8)(unsafe.Add(mBase, _consts[976])))
	if v98 == v60&int32(255) {
		v128 = v58
		v130 = v61
		goto L35
	} else {
		goto L36
	}
L35:
	;
	if v130 == int32(0) {
		goto L24
	} else {
		goto L42
	}
L36:
	;
	goto L37
L37:
	;
	v108 = v58
	v110 = v61
	goto L38
L38:
	;
	v114 = *(*int32)(unsafe.Add(mBase, uint32(v108)))
	v115 = v114 ^ v60&int32(255)*int32(16843009)
	v118 = int32(-2139062144)
	if (int32(16843008)-v115|v115)&v118 != v118 {
		v135 = v108
		v137 = v110
		goto L25
	} else {
		goto L40
	}
L39:
	;
	v128 = v123
	v130 = v125
	goto L35
L40:
	;
	v122 = int32(4)
	v123 = v108 + v122
	v125 = v110 - v122
	if base.Ui32(int32(3)) < base.Ui32(v125) {
		v108 = v123
		v110 = v125
		goto L38
	} else {
		goto L41
	}
L41:
	;
	goto L39
L42:
	;
	v135 = v128
	v137 = v130
	goto L25
L43:
	;
	v147 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v142))))
	if v60&int32(255) == v147 {
		goto L45
	} else {
		goto L46
	}
L44:
	;
	goto L24
L45:
	;
	v164 = v142
	goto L23
L46:
	;
	goto L47
L47:
	;
	v149 = int32(1)
	v152 = v144 - v149
	if v152 != 0 {
		v142 = v142 + v149
		v144 = v152
		goto L43
	} else {
		goto L48
	}
L48:
	;
	goto L44
L49:
	;
	v167 = int32(1694880)
	v169 = int32(17)
	goto L55
L50:
	;
	v275 = v58
	v276 = v164
	goto L51
L51:
	;
	v277 = int32(1)
	v279 = v48 + v277
	if int32(4) < v279 {
		goto L4
	} else {
		goto L79
	}
L52:
	;
	if v272 == int32(0) {
		goto L22
	} else {
		goto L78
	}
L53:
	;
	v272 = int32(0)
	goto L52
L54:
	;
	v250 = v243
	v252 = v245
	goto L72
L55:
	;
	goto L63
L63:
	;
	v206 = int32(*(*uint8)(unsafe.Add(mBase, _consts[977])))
	if v206 == v60&int32(255) {
		v236 = v167
		v238 = v169
		goto L64
	} else {
		goto L65
	}
L64:
	;
	if v238 == int32(0) {
		goto L53
	} else {
		goto L71
	}
L65:
	;
	goto L66
L66:
	;
	v216 = v167
	v218 = v169
	goto L67
L67:
	;
	v222 = *(*int32)(unsafe.Add(mBase, uint32(v216)))
	v223 = v222 ^ v60&int32(255)*int32(16843009)
	v226 = int32(-2139062144)
	if (int32(16843008)-v223|v223)&v226 != v226 {
		v243 = v216
		v245 = v218
		goto L54
	} else {
		goto L69
	}
L68:
	;
	v236 = v231
	v238 = v233
	goto L64
L69:
	;
	v230 = int32(4)
	v231 = v216 + v230
	v233 = v218 - v230
	if base.Ui32(int32(3)) < base.Ui32(v233) {
		v216 = v231
		v218 = v233
		goto L67
	} else {
		goto L70
	}
L70:
	;
	goto L68
L71:
	;
	v243 = v236
	v245 = v238
	goto L54
L72:
	;
	v255 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v250))))
	if v60&int32(255) == v255 {
		goto L74
	} else {
		goto L75
	}
L73:
	;
	goto L53
L74:
	;
	v272 = v250
	goto L52
L75:
	;
	goto L76
L76:
	;
	v257 = int32(1)
	v260 = v252 - v257
	if v260 != 0 {
		v250 = v250 + v257
		v252 = v260
		goto L72
	} else {
		goto L77
	}
L77:
	;
	goto L73
L78:
	;
	v275 = v167
	v276 = v272
	goto L51
L79:
	;
	v318 = v276 - v275 | v46<<(uint(int32(4))%32)
	v319 = v47
	v320 = v279
	v321 = v51
	v322 = v277
	v323 = v44
	goto L21
L80:
	;
	switch v287 - int32(46) {
	case 0:
		goto L18
	case 1:
		goto L17
	default:
		goto L4
	}
L81:
	;
	goto L82
L82:
	;
	if v53 == int32(0) {
		goto L83
	} else {
		goto L84
	}
L83:
	;
	v295 = int32(0)
	if v44 == v295 {
		v318 = v46
		v319 = v47
		v320 = v48
		v321 = v57
		v322 = v295
		v323 = v47
		goto L21
	} else {
		goto L86
	}
L84:
	;
	goto L85
L85:
	;
	v298 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v57))))
	if v298 == int32(0) {
		goto L4
	} else {
		goto L87
	}
L86:
	;
	goto L4
L87:
	;
	v302 = v47 + int32(2)
	if base.Ui32(v40) < base.Ui32(v302) {
		goto L4
	} else {
		goto L88
	}
L88:
	;
	v304 = int32(8)
	v310 = v46<<(uint(v304)%32) | int32(base.Ui32(v46&int32(65280))>>(uint(v304)%32))
	*(*uint16)(unsafe.Add(mBase, uint32(v47))) = uint16(v310)
	v312 = int32(0)
	v318 = v312
	v319 = v302
	v320 = v312
	v321 = v57
	v322 = v312
	v323 = v44
	goto L21
L89:
	;
	goto L20
L90:
	;
	v336 = v47
	v339 = v51
	goto L92
L91:
	;
	if v346 == int32(0) {
		goto L4
	} else {
		goto L169
	}
L92:
	;
	v344 = int32(0)
	v346 = v344
	v348 = v344
	v356 = v339
	goto L94
L93:
	;
	v495 = int32(0)
	v497 = v495
	v499 = v495
	v507 = v365
	goto L134
L94:
	;
	v361 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v356))))
	if v361 == int32(0) {
		goto L91
	} else {
		goto L96
	}
L95:
	;
	if v361&int32(254) != int32(46) {
		goto L4
	} else {
		goto L131
	}
L96:
	;
	v365 = v356 + int32(1)
	v367 = base.I32_extend8_s(v361)
	goto L101
L97:
	;
	if v471 != 0 {
		goto L123
	} else {
		goto L124
	}
L98:
	;
	v471 = int32(0)
	goto L97
L99:
	;
	v449 = v442
	v451 = v444
	goto L117
L100:
	;
	if base.B2i32(v389 != v390) == int32(0) {
		goto L98
	} else {
		goto L108
	}
L101:
	;
	goto L102
L102:
	;
	v381 = int32(1694897)
	v383 = int32(11)
	goto L103
L103:
	;
	v386 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v381))))
	if v386 == v367&int32(255) {
		v442 = v381
		v444 = v383
		goto L99
	} else {
		goto L105
	}
L104:
	;
	goto L100
L105:
	;
	v388 = int32(1)
	v389 = v383 - v388
	v390 = int32(0)
	v393 = v381 + v388
	if v393&int32(3) == v390 {
		goto L100
	} else {
		goto L106
	}
L106:
	;
	if v389 != 0 {
		v381 = v393
		v383 = v389
		goto L103
	} else {
		goto L107
	}
L107:
	;
	goto L104
L108:
	;
	v405 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v393))))
	if v405 == v367&int32(255) {
		v435 = v393
		v437 = v389
		goto L109
	} else {
		goto L110
	}
L109:
	;
	if v437 == int32(0) {
		goto L98
	} else {
		goto L116
	}
L110:
	;
	if base.Ui32(v389) < base.Ui32(int32(4)) {
		v435 = v393
		v437 = v389
		goto L109
	} else {
		goto L111
	}
L111:
	;
	v415 = v393
	v417 = v389
	goto L112
L112:
	;
	v421 = *(*int32)(unsafe.Add(mBase, uint32(v415)))
	v422 = v421 ^ v367&int32(255)*int32(16843009)
	v425 = int32(-2139062144)
	if (int32(16843008)-v422|v422)&v425 != v425 {
		v442 = v415
		v444 = v417
		goto L99
	} else {
		goto L114
	}
L113:
	;
	v435 = v430
	v437 = v432
	goto L109
L114:
	;
	v429 = int32(4)
	v430 = v415 + v429
	v432 = v417 - v429
	if base.Ui32(int32(3)) < base.Ui32(v432) {
		v415 = v430
		v417 = v432
		goto L112
	} else {
		goto L115
	}
L115:
	;
	goto L113
L116:
	;
	v442 = v435
	v444 = v437
	goto L99
L117:
	;
	v454 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v449))))
	if v367&int32(255) == v454 {
		goto L119
	} else {
		goto L120
	}
L118:
	;
	goto L98
L119:
	;
	v471 = v449
	goto L97
L120:
	;
	goto L121
L121:
	;
	v456 = int32(1)
	v459 = v451 - v456
	if v459 != 0 {
		v449 = v449 + v456
		v451 = v459
		goto L117
	} else {
		goto L122
	}
L122:
	;
	goto L118
L123:
	;
	if v348 != 0 {
		goto L126
	} else {
		goto L127
	}
L124:
	;
	goto L125
L125:
	;
	goto L95
L126:
	;
	v473 = int32(0)
	goto L128
L127:
	;
	v473 = v346
	goto L128
L128:
	;
	if v473 != 0 {
		goto L4
	} else {
		goto L129
	}
L129:
	;
	v480 = v471 - int32(1694897) + v348*int32(10)
	if base.Ui32(v480) < base.Ui32(int32(256)) {
		v346 = v346 + int32(1)
		v348 = v480
		v356 = v365
		goto L94
	} else {
		goto L130
	}
L130:
	;
	goto L4
L131:
	;
	if int32(3) < v336-v47 {
		goto L4
	} else {
		goto L132
	}
L132:
	;
	*(*uint8)(unsafe.Add(mBase, uint32(v336))) = uint8(v348)
	if v361 != int32(47) {
		v336 = v336 + int32(1)
		v339 = v365
		goto L92
	} else {
		goto L133
	}
L133:
	;
	goto L93
L134:
	;
	v512 = int32(*(*int8)(unsafe.Add(mBase, uint32(v507))))
	if v512 == int32(0) {
		goto L15
	} else {
		goto L136
	}
L135:
	;
	goto L4
L136:
	;
	v515 = int32(1694908)
	v516 = int32(11)
	goto L140
L137:
	;
	if v619 == int32(0) {
		goto L4
	} else {
		goto L163
	}
L138:
	;
	v619 = int32(0)
	goto L137
L139:
	;
	v597 = v590
	v599 = v592
	goto L157
L140:
	;
	goto L148
L148:
	;
	v553 = int32(*(*uint8)(unsafe.Add(mBase, _consts[978])))
	if v553 == v512&int32(255) {
		v583 = v515
		v585 = v516
		goto L149
	} else {
		goto L150
	}
L149:
	;
	if v585 == int32(0) {
		goto L138
	} else {
		goto L156
	}
L150:
	;
	goto L151
L151:
	;
	v563 = v515
	v565 = v516
	goto L152
L152:
	;
	v569 = *(*int32)(unsafe.Add(mBase, uint32(v563)))
	v570 = v569 ^ v512&int32(255)*int32(16843009)
	v573 = int32(-2139062144)
	if (int32(16843008)-v570|v570)&v573 != v573 {
		v590 = v563
		v592 = v565
		goto L139
	} else {
		goto L154
	}
L153:
	;
	v583 = v578
	v585 = v580
	goto L149
L154:
	;
	v577 = int32(4)
	v578 = v563 + v577
	v580 = v565 - v577
	if base.Ui32(int32(3)) < base.Ui32(v580) {
		v563 = v578
		v565 = v580
		goto L152
	} else {
		goto L155
	}
L155:
	;
	goto L153
L156:
	;
	v590 = v583
	v592 = v585
	goto L139
L157:
	;
	v602 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v597))))
	if v512&int32(255) == v602 {
		goto L159
	} else {
		goto L160
	}
L158:
	;
	goto L138
L159:
	;
	v619 = v597
	goto L137
L160:
	;
	goto L161
L161:
	;
	v604 = int32(1)
	v607 = v599 - v604
	if v607 != 0 {
		v597 = v597 + v604
		v599 = v607
		goto L157
	} else {
		goto L162
	}
L162:
	;
	goto L158
L163:
	;
	if v499 != 0 {
		goto L164
	} else {
		goto L165
	}
L164:
	;
	v623 = int32(0)
	goto L166
L165:
	;
	v623 = v497
	goto L166
L166:
	;
	if v623 != 0 {
		goto L4
	} else {
		goto L167
	}
L167:
	;
	v624 = int32(1)
	v632 = v619 - int32(1694908) + v499*int32(10)
	if v632 < int32(129) {
		v497 = v497 + v624
		v499 = v632
		v507 = v507 + v624
		goto L134
	} else {
		goto L168
	}
L168:
	;
	goto L135
L169:
	;
	if int32(3) < v336-v47 {
		goto L4
	} else {
		goto L170
	}
L170:
	;
	*(*uint8)(unsafe.Add(mBase, uint32(v336))) = uint8(v348)
	v833 = v327
	v838 = v44
	v845 = int32(128)
	goto L5
L171:
	;
	v657 = int32(*(*int8)(unsafe.Add(mBase, uint32(v642))))
	if v657 != 0 {
		goto L173
	} else {
		goto L174
	}
L172:
	;
	if v656 == int32(0) {
		goto L4
	} else {
		goto L208
	}
L173:
	;
	v658 = int32(1694908)
	v659 = int32(11)
	goto L179
L174:
	;
	goto L175
L175:
	;
	goto L172
L176:
	;
	if v762 == int32(0) {
		goto L4
	} else {
		goto L202
	}
L177:
	;
	v762 = int32(0)
	goto L176
L178:
	;
	v740 = v733
	v742 = v735
	goto L196
L179:
	;
	goto L187
L187:
	;
	v696 = int32(*(*uint8)(unsafe.Add(mBase, _consts[978])))
	if v696 == v657&int32(255) {
		v726 = v658
		v728 = v659
		goto L188
	} else {
		goto L189
	}
L188:
	;
	if v728 == int32(0) {
		goto L177
	} else {
		goto L195
	}
L189:
	;
	goto L190
L190:
	;
	v706 = v658
	v708 = v659
	goto L191
L191:
	;
	v712 = *(*int32)(unsafe.Add(mBase, uint32(v706)))
	v713 = v712 ^ v657&int32(255)*int32(16843009)
	v716 = int32(-2139062144)
	if (int32(16843008)-v713|v713)&v716 != v716 {
		v733 = v706
		v735 = v708
		goto L178
	} else {
		goto L193
	}
L192:
	;
	v726 = v721
	v728 = v723
	goto L188
L193:
	;
	v720 = int32(4)
	v721 = v706 + v720
	v723 = v708 - v720
	if base.Ui32(int32(3)) < base.Ui32(v723) {
		v706 = v721
		v708 = v723
		goto L191
	} else {
		goto L194
	}
L194:
	;
	goto L192
L195:
	;
	v733 = v726
	v735 = v728
	goto L178
L196:
	;
	v745 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v740))))
	if v657&int32(255) == v745 {
		goto L198
	} else {
		goto L199
	}
L197:
	;
	goto L177
L198:
	;
	v762 = v740
	goto L176
L199:
	;
	goto L200
L200:
	;
	v747 = int32(1)
	v750 = v742 - v747
	if v750 != 0 {
		v740 = v740 + v747
		v742 = v750
		goto L196
	} else {
		goto L201
	}
L201:
	;
	goto L197
L202:
	;
	if v644 != 0 {
		goto L203
	} else {
		goto L204
	}
L203:
	;
	v766 = int32(0)
	goto L205
L204:
	;
	v766 = v656
	goto L205
L205:
	;
	if v766 != 0 {
		goto L4
	} else {
		goto L206
	}
L206:
	;
	v767 = int32(1)
	v775 = v762 - int32(1694908) + v644*int32(10)
	if v775 < int32(129) {
		v642 = v642 + v767
		v644 = v775
		v656 = v656 + v767
		goto L171
	} else {
		goto L207
	}
L207:
	;
	goto L4
L208:
	;
	v782 = v644
	v785 = v46
	v786 = v47
	v788 = v44
	v792 = v53
	goto L16
L209:
	;
	v813 = v782
	v814 = v786
	v819 = v788
	goto L6
L210:
	;
	goto L211
L211:
	;
	v798 = v786 + int32(2)
	if base.Ui32(v40) < base.Ui32(v798) {
		goto L4
	} else {
		goto L212
	}
L212:
	;
	v800 = int32(8)
	v806 = v785<<(uint(v800)%32) | int32(base.Ui32(v785&int32(65280))>>(uint(v800)%32))
	*(*uint16)(unsafe.Add(mBase, uint32(v786))) = uint16(v806)
	v813 = v782
	v814 = v798
	v819 = v788
	goto L6
L213:
	;
	goto L4
L214:
	;
	v829 = int32(128)
	goto L216
L215:
	;
	v829 = v813
	goto L216
L216:
	;
	v833 = v814
	v838 = v819
	v845 = v829
	goto L5
L217:
	;
	if v833 == v40 {
		goto L4
	} else {
		goto L220
	}
L218:
	;
	goto L219
L219:
	;
	if v833 == v40 {
		goto L2
	} else {
		goto L229
	}
L220:
	;
	v847 = int32(1)
	v848 = v833 - v838
	if v848 <= int32(0) {
		goto L2
	} else {
		goto L221
	}
L221:
	;
	v851 = int32(1)
	if v838 != v833-v851 {
		goto L222
	} else {
		goto L223
	}
L222:
	;
	v862 = v847
	v863 = int32(0)
	goto L225
L223:
	;
	v897 = v847
	goto L224
L224:
	;
	if v848&v851 == int32(0) {
		goto L2
	} else {
		goto L228
	}
L225:
	;
	v877 = v838 + (v848 - v862)
	v878 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v877))))
	*(*uint8)(unsafe.Add(mBase, uint32(v40-v862))) = uint8(v878)
	v880 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v877))) = uint8(v880)
	v883 = v862 ^ int32(-1)
	v885 = v883 + (v848 + v838)
	v886 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v885))))
	*(*uint8)(unsafe.Add(mBase, uint32(v40+v883))) = uint8(v886)
	*(*uint8)(unsafe.Add(mBase, uint32(v885))) = uint8(v880)
	v890 = int32(2)
	v891 = v862 + v890
	v893 = v863 + v890
	if v893 != v848&int32(2147483646) {
		v862 = v891
		v863 = v893
		goto L225
	} else {
		goto L227
	}
L226:
	;
	v897 = v891
	goto L224
L227:
	;
	goto L226
L228:
	;
	v914 = v838 + (v848 - v897)
	v915 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v914))))
	*(*uint8)(unsafe.Add(mBase, uint32(v40-v897))) = uint8(v915)
	v917 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v914))) = uint8(v917)
	goto L2
L229:
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
	var v60 int32
	_ = v60
	var v61 int32
	_ = v61
	var v65 int32
	_ = v65
	var v68 int32
	_ = v68
	var v74 int32
	_ = v74
	var v76 int32
	_ = v76
	var v78 int32
	_ = v78
	var v79 int32
	_ = v79
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
						v52 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v20))))
						if v52&int32(1) != 0 {
							v55 = v35
						} else {
							v55 = v37
						}
						if v32 == int32(3) {
							v60 = int32(16)
						} else {
							v60 = int32(4)
						}
						if v60 != 0 {
							v61 = F__emscripten_memcpy_bulkmem(m, v23+int32(4), v55+int32(2), v60)
							mBase = m.M
						} else {
						}
						if v32 == int32(3) {
							v65 = int32(41)
						} else {
							v65 = int32(17)
						}
						*(*uint8)(unsafe.Add(mBase, uint32(v23))) = uint8(v65)
						v68 = v23
						*(*int32)(unsafe.Add(mBase, uint32(v15))) = v68
						v74 = *(*int32)(unsafe.Add(mBase, uint32(v9)+4))
						*(*int32)(unsafe.Add(mBase, uint32(v15)+4)) = v74
						v76 = *(*int32)(unsafe.Add(mBase, uint32(v9)+8))
						*(*int32)(unsafe.Add(mBase, uint32(v15)+8)) = v76
						v78 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v9)+12)))
						v79 = int32(0)
						*(*uint8)(unsafe.Add(mBase, uint32(v15)+14)) = uint8(v79)
						*(*uint16)(unsafe.Add(mBase, uint32(v15)+12)) = uint16(v78)
						return v15
					}
				}
			} else {
				v68 = int32(0)
				*(*int32)(unsafe.Add(mBase, uint32(v15))) = v68
				v74 = *(*int32)(unsafe.Add(mBase, uint32(v9)+4))
				*(*int32)(unsafe.Add(mBase, uint32(v15)+4)) = v74
				v76 = *(*int32)(unsafe.Add(mBase, uint32(v9)+8))
				*(*int32)(unsafe.Add(mBase, uint32(v15)+8)) = v76
				v78 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v9)+12)))
				v79 = int32(0)
				*(*uint8)(unsafe.Add(mBase, uint32(v15)+14)) = uint8(v79)
				*(*uint16)(unsafe.Add(mBase, uint32(v15)+12)) = uint16(v78)
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
	var v86 int32
	_ = v86
	var v91 int32
	_ = v91
	var v96 int32
	_ = v96
	var v109 int32
	_ = v109
	var v111 int32
	_ = v111
	var v119 int32
	_ = v119
	var v120 int32
	_ = v120
	var v121 int32
	_ = v121
	var v123 int32
	_ = v123
	var v125 int32
	_ = v125
	var v126 int32
	_ = v126
	var v128 int32
	_ = v128
	var v130 int32
	_ = v130
	var v133 int32
	_ = v133
	var v135 int32
	_ = v135
	var v138 int32
	_ = v138
	var v140 int32
	_ = v140
	var v141 int32
	_ = v141
	var v144 int32
	_ = v144
	var v148 int32
	_ = v148
	var v151 int32
	_ = v151
	var v154 int32
	_ = v154
	var v158 int32
	_ = v158
	var v161 int32
	_ = v161
	var v167 int32
	_ = v167
	var v168 int32
	_ = v168
	var v171 int32
	_ = v171
	var v176 int32
	_ = v176
	var v179 int32
	_ = v179
	var v182 int32
	_ = v182
	var v183 int32
	_ = v183
	var v186 int32
	_ = v186
	var v192 int32
	_ = v192
	var v194 int32
	_ = v194
	var v197 int32
	_ = v197
	var v199 int32
	_ = v199
	var v202 int32
	_ = v202
	var v208 int32
	_ = v208
	var v211 int32
	_ = v211
	var v215 int32
	_ = v215
	var v220 int32
	_ = v220
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
	v208 = m.ExcPending
	if v208 != 0 {
		goto L1
	} else {
		goto L63
	}
L13:
	;
	v119 = F_palloc0(m, int32(22))
	mBase = m.M
	v120 = m.ExcPending
	if v120 != 0 {
		goto L1
	} else {
		goto L31
	}
L14:
	;
	v96 = v84
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
		v109 = int32(0)
		v111 = v68
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
	v86 = v58
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
	v86 = v82
	v91 = v80
	goto L14
L28:
	;
	if int32(base.Ui32((v86^v91)&int32(255))>>(uint(int32(8)-v96)%32)) != 0 {
		v96 = v96 - int32(1)
		goto L28
	} else {
		goto L30
	}
L29:
	;
	v109 = v96
	v111 = v85
	goto L13
L30:
	;
	goto L29
L31:
	;
	v121 = int32(1)
	v123 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v119))))
	v125 = v123 & v121
	if v125 != 0 {
		goto L32
	} else {
		goto L33
	}
L32:
	;
	v126 = v121
	goto L34
L33:
	;
	v126 = int32(4)
	goto L34
L34:
	;
	v128 = int32(1)
	v130 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v11))))
	if v130&v128 != 0 {
		goto L35
	} else {
		goto L36
	}
L35:
	;
	v133 = v128
	goto L37
L36:
	;
	v133 = int32(4)
	goto L37
L37:
	;
	v135 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v11+v133))))
	*(*uint8)(unsafe.Add(mBase, uint32(v119+v126))) = uint8(v135)
	v138 = v119 + int32(1)
	v140 = v119 + int32(4)
	if v125 != 0 {
		goto L38
	} else {
		goto L39
	}
L38:
	;
	v141 = v138
	goto L40
L39:
	;
	v141 = v140
	goto L40
L40:
	;
	v144 = v109 + v111<<(uint(int32(3))%32)
	*(*uint8)(unsafe.Add(mBase, uint32(v141)+1)) = uint8(v144)
	if v144 <= int32(0) {
		goto L41
	} else {
		goto L42
	}
L41:
	;
	v192 = int32(1)
	v194 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v119))))
	if v194&v192 != 0 {
		goto L57
	} else {
		goto L58
	}
L42:
	;
	v148 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v119))))
	if v148&int32(1) != 0 {
		goto L43
	} else {
		goto L44
	}
L43:
	;
	v151 = v138
	goto L45
L44:
	;
	v151 = v140
	goto L45
L45:
	;
	v154 = int32(1)
	v158 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v11))))
	if v158&v154 != 0 {
		goto L46
	} else {
		goto L47
	}
L46:
	;
	v161 = v11 + v154
	goto L48
L47:
	;
	v161 = v11 + int32(4)
	goto L48
L48:
	;
	v167 = base.I32_div_s(v144+int32(7), int32(8))
	if v167 != 0 {
		goto L50
	} else {
		goto L51
	}
L49:
	;
	v171 = v109 & int32(7)
	if v171 == int32(0) {
		goto L41
	} else {
		goto L53
	}
L50:
	;
	v168 = F__emscripten_memcpy_bulkmem(m, v151+int32(2), v161+int32(2), v167)
	mBase = m.M
	goto L52
L51:
	;
	goto L52
L52:
	;
	goto L49
L53:
	;
	v176 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v119))))
	if v176&int32(1) != 0 {
		goto L54
	} else {
		goto L55
	}
L54:
	;
	v179 = v138
	goto L56
L55:
	;
	v179 = v140
	goto L56
L56:
	;
	v182 = int32(base.Ui32(v144)>>(uint(int32(3))%32)) + v179 + int32(2)
	v183 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v182))))
	v186 = v183 & (int32(-256) >> (uint(v171) % 32))
	*(*uint8)(unsafe.Add(mBase, uint32(v182))) = uint8(v186)
	goto L41
L57:
	;
	v197 = v192
	goto L59
L58:
	;
	v197 = int32(4)
	goto L59
L59:
	;
	v199 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v119+v197))))
	if v199 == int32(2) {
		goto L60
	} else {
		goto L61
	}
L60:
	;
	v202 = int32(40)
	goto L62
L61:
	;
	v202 = int32(88)
	goto L62
L62:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v119))) = v202
	return v119
L63:
	;
	F_errcode(m, int32(50856066))
	mBase = m.M
	v211 = m.ExcPending
	if v211 != 0 {
		goto L1
	} else {
		goto L64
	}
L64:
	;
	F_errmsg(m, int32(177051), int32(0))
	mBase = m.M
	v215 = m.ExcPending
	if v215 != 0 {
		goto L1
	} else {
		goto L65
	}
L65:
	;
	F_errfinish(m, int32(517688), int32(1450), int32(415736))
	mBase = m.M
	v220 = m.ExcPending
	if v220 != 0 {
		goto L1
	} else {
		goto L66
	}
L66:
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
