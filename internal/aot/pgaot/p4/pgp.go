package p4

import (
	base "github.com/shibukawa/pgmem/internal/aot/pgaot/base"
	"unsafe"
)

func F_pgp_armor_decode(m *base.Module, l0 int32, l1 int32, l2 int32) int32 {
	mBase := m.M
	_ = mBase
	var v12 int32
	_ = v12
	var v14 int32
	_ = v14
	var v17 int32
	_ = v17
	var v18 int32
	_ = v18
	var v30 int32
	_ = v30
	var v36 int32
	_ = v36
	var v39 int32
	_ = v39
	var v43 int32
	_ = v43
	var v50 int32
	_ = v50
	var v53 int32
	_ = v53
	var v55 int32
	_ = v55
	var v57 int32
	_ = v57
	var v62 int32
	_ = v62
	var v73 int32
	_ = v73
	var v79 int32
	_ = v79
	var v85 int32
	_ = v85
	var v90 int32
	_ = v90
	var v100 int32
	_ = v100
	var v102 int32
	_ = v102
	var v104 int32
	_ = v104
	var v111 int32
	_ = v111
	var v113 int32
	_ = v113
	var v117 int32
	_ = v117
	var v127 int32
	_ = v127
	var v131 int32
	_ = v131
	var v132 int32
	_ = v132
	var v141 int32
	_ = v141
	var v143 int32
	_ = v143
	var v148 int32
	_ = v148
	var v153 int32
	_ = v153
	var v157 int32
	_ = v157
	var v164 int32
	_ = v164
	var v167 int32
	_ = v167
	var v169 int32
	_ = v169
	var v171 int32
	_ = v171
	var v176 int32
	_ = v176
	var v187 int32
	_ = v187
	var v193 int32
	_ = v193
	var v199 int32
	_ = v199
	var v204 int32
	_ = v204
	var v214 int32
	_ = v214
	var v216 int32
	_ = v216
	var v218 int32
	_ = v218
	var v225 int32
	_ = v225
	var v227 int32
	_ = v227
	var v231 int32
	_ = v231
	var v241 int32
	_ = v241
	var v245 int32
	_ = v245
	var v247 int32
	_ = v247
	var v258 int32
	_ = v258
	var v262 int32
	_ = v262
	var v263 int32
	_ = v263
	var v266 int32
	_ = v266
	var v275 int32
	_ = v275
	var v277 int32
	_ = v277
	var v280 int32
	_ = v280
	var v282 int32
	_ = v282
	var v283 int32
	_ = v283
	var v284 int32
	_ = v284
	var v285 int32
	_ = v285
	var v287 int32
	_ = v287
	var v292 int32
	_ = v292
	var v294 int32
	_ = v294
	var v295 int32
	_ = v295
	var v299 int32
	_ = v299
	var v309 int32
	_ = v309
	var v311 int32
	_ = v311
	var v315 int32
	_ = v315
	var v316 int32
	_ = v316
	var v319 int32
	_ = v319
	var v323 int32
	_ = v323
	var v324 int32
	_ = v324
	var v326 int32
	_ = v326
	var v329 int32
	_ = v329
	var v331 int32
	_ = v331
	var v336 int32
	_ = v336
	var v338 int32
	_ = v338
	var v343 int32
	_ = v343
	var v345 int32
	_ = v345
	var v348 int32
	_ = v348
	var v350 int32
	_ = v350
	var v353 int32
	_ = v353
	var v365 int32
	_ = v365
	var v369 int32
	_ = v369
	var v371 int32
	_ = v371
	var v384 int32
	_ = v384
	var v387 int32
	_ = v387
	var v391 int32
	_ = v391
	var v400 int32
	_ = v400
	var v402 int32
	_ = v402
	var v409 int32
	_ = v409
	var v411 int32
	_ = v411
	var v422 int32
	_ = v422
	var v423 int32
	_ = v423
	var v430 int32
	_ = v430
	var v431 int32
	_ = v431
	var v434 int32
	_ = v434
	var v436 int32
	_ = v436
	var v437 int32
	_ = v437
	var v438 int32
	_ = v438
	var v442 int32
	_ = v442
	var v452 int32
	_ = v452
	var v453 int32
	_ = v453
	var v455 int32
	_ = v455
	var v479 int32
	_ = v479
	var v491 int32
	_ = v491
	var v495 int32
	_ = v495
	var v497 int32
	_ = v497
	var v501 int32
	_ = v501
	var v504 int32
	_ = v504
	var v513 int32
	_ = v513
	var v519 int32
	_ = v519
	var v520 int32
	_ = v520
	var v521 int32
	_ = v521
	var v525 int32
	_ = v525
	var v529 int32
	_ = v529
	var v530 int32
	_ = v530
	var v531 int32
	_ = v531
	var v534 int32
	_ = v534
	var v535 int32
	_ = v535
	var v546 int32
	_ = v546
	var v549 int32
	_ = v549
	var v551 int32
	_ = v551
	var v552 int32
	_ = v552
	var v553 int32
	_ = v553
	var v557 int32
	_ = v557
	var v561 int32
	_ = v561
	var v583 int32
	_ = v583
	var v587 int32
	_ = v587
	var v588 int32
	_ = v588
	var v589 int32
	_ = v589
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
	var v609 int32
	_ = v609
	var v610 int32
	_ = v610
	var v613 int32
	_ = v613
	var v615 int32
	_ = v615
	var v616 int32
	_ = v616
	var v617 int32
	_ = v617
	var v621 int32
	_ = v621
	var v631 int32
	_ = v631
	var v632 int32
	_ = v632
	var v634 int32
	_ = v634
	var v658 int32
	_ = v658
	var v670 int32
	_ = v670
	var v674 int32
	_ = v674
	var v676 int32
	_ = v676
	var v680 int32
	_ = v680
	var v683 int32
	_ = v683
	var v692 int32
	_ = v692
	var v698 int32
	_ = v698
	var v699 int32
	_ = v699
	var v700 int32
	_ = v700
	var v704 int32
	_ = v704
	var v708 int32
	_ = v708
	var v709 int32
	_ = v709
	var v710 int32
	_ = v710
	var v713 int32
	_ = v713
	var v714 int32
	_ = v714
	var v725 int32
	_ = v725
	var v728 int32
	_ = v728
	var v730 int32
	_ = v730
	var v731 int32
	_ = v731
	var v732 int32
	_ = v732
	var v736 int32
	_ = v736
	var v740 int32
	_ = v740
	var v746 int32
	_ = v746
	var v762 int32
	_ = v762
	var v773 int32
	_ = v773
	var v775 int32
	_ = v775
	var v778 int32
	_ = v778
	var v780 int32
	_ = v780
	var v786 int32
	_ = v786
	var v789 int32
	_ = v789
	var v791 int32
	_ = v791
	var v796 int32
	_ = v796
	var v798 int32
	_ = v798
	var v803 int32
	_ = v803
	var v805 int32
	_ = v805
	var v810 int32
	_ = v810
	var v812 int32
	_ = v812
	var v817 int32
	_ = v817
	var v819 int32
	_ = v819
	var v824 int32
	_ = v824
	var v826 int32
	_ = v826
	var v831 int32
	_ = v831
	var v833 int32
	_ = v833
	var v838 int32
	_ = v838
	var v840 int32
	_ = v840
	var v845 int32
	_ = v845
	var v846 int32
	_ = v846
	var v849 int32
	_ = v849
	var v864 int32
	_ = v864
	var v866 int32
	_ = v866
	var v888 int32
	_ = v888
	var v898 int32
	_ = v898
	var v904 int32
	_ = v904
	var v911 int32
	_ = v911
	v12 = m.G0
	v14 = v12 - int32(16)
	m.G0 = v14
	*(*int32)(unsafe.Add(mBase, uint32(v14)+12)) = l0
	v17 = int32(-101)
	v18 = l0 + l1
	if base.Ui32(v18) <= base.Ui32(l0) {
		v127 = v17
		goto L4
	} else {
		goto L5
	}
L1:
	;
	F_errstart_cold(m, int32(22), int32(0))
	mBase = m.M
	v898 = m.ExcPending
	if v898 != 0 {
		goto L168
	} else {
		goto L245
	}
L2:
	;
	m.G0 = v14 + int32(16)
	return v888
L3:
	;
	if v127 <= int32(0) {
		v888 = v17
		goto L2
	} else {
		goto L42
	}
L4:
	;
	goto L3
L5:
	;
	v30 = int32(10)
	goto L7
L7:
	;
	goto L8
L8:
	;
	if v18-l0 < v30 {
		v127 = v17
		goto L4
	} else {
		goto L9
	}
L9:
	;
	v36 = int32(508064)
	goto L11
L11:
	;
	goto L12
L12:
	;
	v39 = int32(*(*int8)(unsafe.Add(mBase, _consts[1529])))
	v43 = l0
	goto L13
L13:
	;
	v50 = F_memchr(m, v43, v39, v18-v43)
	mBase = m.M
	if v50 == int32(0) {
		v127 = v17
		goto L4
	} else {
		goto L15
	}
L14:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v14+int32(12)))) = v50
	if base.Ui32(v18) <= base.Ui32(v53) {
		v90 = v53
		goto L26
	} else {
		goto L27
	}
L15:
	;
	v53 = v30 + v50
	if base.Ui32(v18) < base.Ui32(v53) {
		v127 = v17
		goto L4
	} else {
		goto L16
	}
L16:
	;
	v55 = F_memcmp(m, v50, v36, v30)
	mBase = m.M
	if v55 != 0 {
		goto L17
	} else {
		goto L18
	}
L17:
	;
	v57 = v50 + int32(1)
	if base.Ui32(v57) < base.Ui32(v18) {
		v43 = v57
		goto L13
	} else {
		goto L20
	}
L18:
	;
	goto L19
L19:
	;
	if l0 == v50 {
		goto L21
	} else {
		goto L22
	}
L20:
	;
	v127 = v17
	goto L4
L21:
	;
	goto L14
L22:
	;
	v62 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v50-int32(1)))))
	if v62 == int32(10) {
		goto L21
	} else {
		goto L23
	}
L23:
	;
	if base.Ui32(v18) <= base.Ui32(v53) {
		v127 = v17
		goto L4
	} else {
		goto L24
	}
L24:
	;
	if v30 <= v18-v53 {
		v43 = v53
		goto L13
	} else {
		goto L25
	}
L25:
	;
	v127 = v17
	goto L4
L26:
	;
	if v18-v90 < int32(5) {
		v127 = v17
		goto L4
	} else {
		goto L33
	}
L27:
	;
	v73 = v53
	goto L28
L28:
	;
	v79 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v73))))
	if v79 == int32(45) {
		v90 = v73
		goto L26
	} else {
		goto L30
	}
L29:
	;
	v90 = v18
	goto L26
L30:
	;
	if base.Ui32(v79) < base.Ui32(int32(32)) {
		v127 = v17
		goto L4
	} else {
		goto L31
	}
L31:
	;
	v85 = v73 + int32(1)
	if base.Ui32(v85) < base.Ui32(v18) {
		v73 = v85
		goto L28
	} else {
		goto L32
	}
L32:
	;
	goto L29
L33:
	;
	v100 = F_memcmp(m, v90, v36, int32(5))
	mBase = m.M
	if v100 != 0 {
		v127 = v17
		goto L4
	} else {
		goto L34
	}
L34:
	;
	v102 = v90 + int32(5)
	if base.Ui32(v18) <= base.Ui32(v102) {
		v117 = v102
		goto L35
	} else {
		goto L36
	}
L35:
	;
	v127 = v117 - v50
	goto L4
L36:
	;
	v104 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v102))))
	switch v104 - int32(10) {
	case 0, 3:
		goto L37
	default:
		v127 = v17
		goto L4
	}
L37:
	;
	if v104 == int32(13) {
		goto L38
	} else {
		goto L39
	}
L38:
	;
	v111 = v90 + int32(6)
	goto L40
L39:
	;
	v111 = v102
	goto L40
L40:
	;
	if base.Ui32(v18) <= base.Ui32(v111) {
		v117 = v111
		goto L35
	} else {
		goto L41
	}
L41:
	;
	v113 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v111))))
	v117 = v111 + base.B2i32(v113 == int32(10))
	goto L35
L42:
	;
	v131 = *(*int32)(unsafe.Add(mBase, uint32(v14)+12))
	v132 = v131 + v127
	v141 = int32(-101)
	if base.Ui32(v18) <= base.Ui32(v132) {
		v241 = v141
		goto L44
	} else {
		goto L45
	}
L43:
	;
	if v241 <= int32(0) {
		v888 = v17
		goto L2
	} else {
		goto L82
	}
L44:
	;
	goto L43
L45:
	;
	v143 = int32(8)
	goto L46
L46:
	;
	goto L48
L48:
	;
	if v18-v132 < v143 {
		v241 = v141
		goto L44
	} else {
		goto L49
	}
L49:
	;
	v148 = int32(520278)
	goto L50
L50:
	;
	goto L52
L52:
	;
	v153 = int32(*(*int8)(unsafe.Add(mBase, _consts[1530])))
	v157 = v132
	goto L53
L53:
	;
	v164 = F_memchr(m, v157, v153, v18-v157)
	mBase = m.M
	if v164 == int32(0) {
		v241 = v141
		goto L44
	} else {
		goto L55
	}
L54:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v14+int32(8)))) = v164
	if base.Ui32(v18) <= base.Ui32(v167) {
		v204 = v167
		goto L66
	} else {
		goto L67
	}
L55:
	;
	v167 = v143 + v164
	if base.Ui32(v18) < base.Ui32(v167) {
		v241 = v141
		goto L44
	} else {
		goto L56
	}
L56:
	;
	v169 = F_memcmp(m, v164, v148, v143)
	mBase = m.M
	if v169 != 0 {
		goto L57
	} else {
		goto L58
	}
L57:
	;
	v171 = v164 + int32(1)
	if base.Ui32(v171) < base.Ui32(v18) {
		v157 = v171
		goto L53
	} else {
		goto L60
	}
L58:
	;
	goto L59
L59:
	;
	if v132 == v164 {
		goto L61
	} else {
		goto L62
	}
L60:
	;
	v241 = v141
	goto L44
L61:
	;
	goto L54
L62:
	;
	v176 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v164-int32(1)))))
	if v176 == int32(10) {
		goto L61
	} else {
		goto L63
	}
L63:
	;
	if base.Ui32(v18) <= base.Ui32(v167) {
		v241 = v141
		goto L44
	} else {
		goto L64
	}
L64:
	;
	if v143 <= v18-v167 {
		v157 = v167
		goto L53
	} else {
		goto L65
	}
L65:
	;
	v241 = v141
	goto L44
L66:
	;
	if v18-v204 < int32(5) {
		v241 = v141
		goto L44
	} else {
		goto L73
	}
L67:
	;
	v187 = v167
	goto L68
L68:
	;
	v193 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v187))))
	if v193 == int32(45) {
		v204 = v187
		goto L66
	} else {
		goto L70
	}
L69:
	;
	v204 = v18
	goto L66
L70:
	;
	if base.Ui32(v193) < base.Ui32(int32(32)) {
		v241 = v141
		goto L44
	} else {
		goto L71
	}
L71:
	;
	v199 = v187 + int32(1)
	if base.Ui32(v199) < base.Ui32(v18) {
		v187 = v199
		goto L68
	} else {
		goto L72
	}
L72:
	;
	goto L69
L73:
	;
	v214 = F_memcmp(m, v204, v148, int32(5))
	mBase = m.M
	if v214 != 0 {
		v241 = v141
		goto L44
	} else {
		goto L74
	}
L74:
	;
	v216 = v204 + int32(5)
	if base.Ui32(v18) <= base.Ui32(v216) {
		v231 = v216
		goto L75
	} else {
		goto L76
	}
L75:
	;
	v241 = v231 - v164
	goto L44
L76:
	;
	v218 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v216))))
	switch v218 - int32(10) {
	case 0, 3:
		goto L77
	default:
		v241 = v141
		goto L44
	}
L77:
	;
	if v218 == int32(13) {
		goto L78
	} else {
		goto L79
	}
L78:
	;
	v225 = v204 + int32(6)
	goto L80
L79:
	;
	v225 = v216
	goto L80
L80:
	;
	if base.Ui32(v18) <= base.Ui32(v225) {
		v231 = v225
		goto L75
	} else {
		goto L81
	}
L81:
	;
	v227 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v225))))
	v231 = v225 + base.B2i32(v227 == int32(10))
	goto L75
L82:
	;
	v245 = *(*int32)(unsafe.Add(mBase, uint32(v14)+8))
	if base.Ui32(v245) <= base.Ui32(v132) {
		v371 = v132
		goto L83
	} else {
		goto L84
	}
L83:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v14)+12)) = v245
	if base.Ui32(v245) < base.Ui32(v371) {
		v888 = v17
		goto L2
	} else {
		goto L116
	}
L84:
	;
	v247 = v132
	goto L85
L85:
	;
	v258 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v247))))
	switch v258 - int32(10) {
	case 0, 3:
		v371 = v247
		goto L83
	default:
		goto L87
	}
L86:
	;
	v371 = v369
	goto L83
L87:
	;
	v262 = v245 - v247
	v263 = int32(0)
	v266 = base.B2i32(v262 != v263)
	if v247&int32(3) == v263 {
		v292 = v247
		v294 = v262
		v295 = v266
		goto L91
	} else {
		goto L92
	}
L88:
	;
	if v365 == int32(0) {
		v888 = v17
		goto L2
	} else {
		goto L114
	}
L89:
	;
	v365 = int32(0)
	goto L88
L90:
	;
	v343 = v336
	v345 = v338
	goto L108
L91:
	;
	if v295 == int32(0) {
		goto L89
	} else {
		goto L99
	}
L92:
	;
	if v262 == int32(0) {
		v292 = v247
		v294 = v262
		v295 = v266
		goto L91
	} else {
		goto L93
	}
L93:
	;
	v275 = v247
	v277 = v262
	goto L94
L94:
	;
	v280 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v275))))
	if v280 == int32(10) {
		v336 = v275
		v338 = v277
		goto L90
	} else {
		goto L96
	}
L95:
	;
	v292 = v287
	v294 = v283
	v295 = v285
	goto L91
L96:
	;
	v282 = int32(1)
	v283 = v277 - v282
	v284 = int32(0)
	v285 = base.B2i32(v283 != v284)
	v287 = v275 + v282
	if v287&int32(3) == v284 {
		v292 = v287
		v294 = v283
		v295 = v285
		goto L91
	} else {
		goto L97
	}
L97:
	;
	if v283 != 0 {
		v275 = v287
		v277 = v283
		goto L94
	} else {
		goto L98
	}
L98:
	;
	goto L95
L99:
	;
	v299 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v292))))
	if v299 == int32(10) {
		v329 = v292
		v331 = v294
		goto L100
	} else {
		goto L101
	}
L100:
	;
	if v331 == int32(0) {
		goto L89
	} else {
		goto L107
	}
L101:
	;
	if base.Ui32(v294) < base.Ui32(int32(4)) {
		v329 = v292
		v331 = v294
		goto L100
	} else {
		goto L102
	}
L102:
	;
	v309 = v292
	v311 = v294
	goto L103
L103:
	;
	v315 = *(*int32)(unsafe.Add(mBase, uint32(v309)))
	v316 = v315 ^ int32(168430090)
	v319 = int32(-2139062144)
	if (int32(16843008)-v316|v316)&v319 != v319 {
		v336 = v309
		v338 = v311
		goto L90
	} else {
		goto L105
	}
L104:
	;
	v329 = v324
	v331 = v326
	goto L100
L105:
	;
	v323 = int32(4)
	v324 = v309 + v323
	v326 = v311 - v323
	if base.Ui32(int32(3)) < base.Ui32(v326) {
		v309 = v324
		v311 = v326
		goto L103
	} else {
		goto L106
	}
L106:
	;
	goto L104
L107:
	;
	v336 = v329
	v338 = v331
	goto L90
L108:
	;
	v348 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v343))))
	if int32(10) == v348 {
		goto L110
	} else {
		goto L111
	}
L109:
	;
	goto L89
L110:
	;
	v365 = v343
	goto L88
L111:
	;
	goto L112
L112:
	;
	v350 = int32(1)
	v353 = v345 - v350
	if v353 != 0 {
		v343 = v343 + v350
		v345 = v353
		goto L108
	} else {
		goto L113
	}
L113:
	;
	goto L109
L114:
	;
	v369 = v365 + int32(1)
	if base.Ui32(v369) < base.Ui32(v245) {
		v247 = v369
		goto L85
	} else {
		goto L115
	}
L115:
	;
	goto L86
L116:
	;
	v384 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v245))))
	if v384 == int32(61) {
		goto L118
	} else {
		goto L119
	}
L117:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v14)+12)) = v409
	v422 = v14 + int32(4)
	v423 = int32(0)
	goto L129
L118:
	;
	v387 = *(*int32)(unsafe.Add(mBase, uint32(v14)+12))
	v409 = v387
	v411 = v245
	goto L117
L119:
	;
	goto L120
L120:
	;
	v391 = v245
	goto L121
L121:
	;
	v400 = v391 - int32(1)
	if base.Ui32(v371) <= base.Ui32(v400) {
		goto L123
	} else {
		goto L124
	}
L122:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v14)+12)) = v400
	v888 = v17
	goto L2
L123:
	;
	v402 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v400))))
	if v402 != int32(61) {
		v391 = v400
		goto L121
	} else {
		goto L126
	}
L124:
	;
	goto L125
L125:
	;
	goto L122
L126:
	;
	v409 = v400
	v411 = v400
	goto L117
L127:
	;
	if v583 != int32(3) {
		v888 = v17
		goto L2
	} else {
		goto L167
	}
L128:
	;
	goto L127
L129:
	;
	v430 = v411 + int32(5)
	v431 = v411 + int32(1)
	v434 = v422
	v436 = v423
	v437 = v423
	v438 = v423
	goto L132
L131:
	;
	v583 = v557 - v422
	goto L128
L132:
	;
	v442 = v431
	goto L141
L133:
	;
	if v561 != 0 {
		v583 = int32(-101)
		goto L128
	} else {
		goto L166
	}
L134:
	;
	goto L133
L135:
	;
	if base.Ui32(v452) < base.Ui32(v430) {
		v431 = v452
		v434 = v549
		v436 = v551
		v437 = v552
		v438 = v553
		goto L132
	} else {
		goto L165
	}
L136:
	;
	v535 = int32(0)
	if base.B2i32(v530 == v535)&base.B2i32(base.Ui32(v531) < base.Ui32(int32(3))) == v535 {
		goto L162
	} else {
		goto L163
	}
L137:
	;
	v525 = int32(base.Ui32(v519) >> (uint(int32(8)) % 32))
	*(*uint8)(unsafe.Add(mBase, uint32(v434)+1)) = uint8(v525)
	v529 = v519
	v530 = v520
	v531 = v521
	v534 = v434 + int32(2)
	goto L136
L138:
	;
	v513 = int32(base.Ui32(v437) >> (uint(int32(10)) % 32))
	*(*uint8)(unsafe.Add(mBase, uint32(v434))) = uint8(v513)
	v519 = v437 << (uint(int32(6)) % 32)
	v520 = int32(0)
	v521 = int32(2)
	goto L137
L139:
	;
	v495 = v491 + v437<<(uint(int32(6))%32)
	v497 = v438 + int32(1)
	if v497 != int32(4) {
		goto L157
	} else {
		goto L158
	}
L140:
	;
	v491 = int32(62)
	goto L139
L141:
	;
	v452 = v442 + int32(1)
	v453 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v442))))
	v455 = v453 - int32(65)
	if base.Ui32(v455&int32(255)) <= base.Ui32(int32(25)) {
		v491 = v455
		goto L139
	} else {
		goto L143
	}
L142:
	;
	if v436 != 0 {
		goto L153
	} else {
		goto L154
	}
L143:
	;
	if base.Ui32((v453-int32(97))&int32(255)) <= base.Ui32(int32(25)) {
		goto L144
	} else {
		goto L145
	}
L144:
	;
	v491 = v453 - int32(71)
	goto L139
L145:
	;
	goto L146
L146:
	;
	if base.Ui32((v453-int32(48))&int32(255)) <= base.Ui32(int32(9)) {
		goto L147
	} else {
		goto L148
	}
L147:
	;
	v491 = (v453 + int32(4)) & int32(255)
	goto L139
L148:
	;
	goto L149
L149:
	;
	v479 = int32(-101)
	switch v453 - int32(9) {
	case 0, 1, 4, 23:
		goto L151
	default:
		v583 = v479
		goto L128
	case 34:
		goto L140
	case 38:
		v491 = int32(63)
		goto L139
	case 52:
		goto L150
	}
L150:
	;
	goto L142
L151:
	;
	if base.Ui32(v452) < base.Ui32(v430) {
		v442 = v452
		goto L141
	} else {
		goto L152
	}
L152:
	;
	v557 = v434
	v561 = v438
	goto L134
L153:
	;
	v491 = int32(0)
	goto L139
L154:
	;
	goto L155
L155:
	;
	switch v438 - int32(2) {
	case 0:
		goto L156
	case 1:
		goto L138
	default:
		v583 = v479
		goto L128
	}
L156:
	;
	v549 = v434
	v551 = int32(1)
	v552 = v437 << (uint(int32(6)) % 32)
	v553 = int32(3)
	goto L135
L157:
	;
	v549 = v434
	v551 = v436
	v552 = v495
	v553 = v497
	goto L135
L158:
	;
	goto L159
L159:
	;
	v501 = int32(base.Ui32(v495) >> (uint(int32(16)) % 32))
	*(*uint8)(unsafe.Add(mBase, uint32(v434))) = uint8(v501)
	v504 = base.B2i32(v436 == int32(0))
	if v436 == int32(0) {
		v519 = v495
		v520 = v504
		v521 = v436
		goto L137
	} else {
		goto L160
	}
L160:
	;
	if base.Ui32(int32(2)) <= base.Ui32(v436) {
		v519 = v495
		v520 = v504
		v521 = v436
		goto L137
	} else {
		goto L161
	}
L161:
	;
	v529 = v495
	v530 = int32(0)
	v531 = v436
	v534 = v434 + int32(1)
	goto L136
L162:
	;
	*(*uint8)(unsafe.Add(mBase, uint32(v534))) = uint8(v529)
	v546 = v534 + int32(1)
	goto L164
L163:
	;
	v546 = v534
	goto L164
L164:
	;
	v549 = v546
	v551 = v531
	v552 = v535
	v553 = int32(0)
	goto L135
L165:
	;
	v557 = v549
	v561 = v553
	goto L134
L166:
	;
	goto L131
L167:
	;
	v587 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v14)+6)))
	v588 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v14)+5)))
	v589 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v14)+4)))
	v593 = int32(base.Ui32(l1*int32(3)) >> (uint(int32(2)) % 32))
	F_enlargeStringInfo(m, l2, v593)
	mBase = m.M
	v597 = m.ExcPending
	if v597 != 0 {
		goto L168
	} else {
		goto L169
	}
L168:
	;
	return int32(0)
L169:
	;
	v600 = v371 ^ int32(-1) + v411
	v601 = *(*int32)(unsafe.Add(mBase, uint32(l2)))
	v602 = int32(0)
	if v600 != 0 {
		goto L172
	} else {
		goto L173
	}
L170:
	;
	if v593 < v762 {
		goto L1
	} else {
		goto L210
	}
L171:
	;
	goto L170
L172:
	;
	v609 = v371 + v600
	v610 = v371
	v613 = v601
	v615 = v602
	v616 = v602
	v617 = v602
	goto L175
L173:
	;
	v746 = v601
	goto L174
L174:
	;
	v762 = v746 - v601
	goto L171
L175:
	;
	v621 = v610
	goto L184
L176:
	;
	if v740 != 0 {
		v762 = int32(-101)
		goto L171
	} else {
		goto L209
	}
L177:
	;
	goto L176
L178:
	;
	if base.Ui32(v631) < base.Ui32(v609) {
		v610 = v631
		v613 = v728
		v615 = v730
		v616 = v731
		v617 = v732
		goto L175
	} else {
		goto L208
	}
L179:
	;
	v714 = int32(0)
	if base.B2i32(v709 == v714)&base.B2i32(base.Ui32(v710) < base.Ui32(int32(3))) == v714 {
		goto L205
	} else {
		goto L206
	}
L180:
	;
	v704 = int32(base.Ui32(v698) >> (uint(int32(8)) % 32))
	*(*uint8)(unsafe.Add(mBase, uint32(v613)+1)) = uint8(v704)
	v708 = v698
	v709 = v699
	v710 = v700
	v713 = v613 + int32(2)
	goto L179
L181:
	;
	v692 = int32(base.Ui32(v616) >> (uint(int32(10)) % 32))
	*(*uint8)(unsafe.Add(mBase, uint32(v613))) = uint8(v692)
	v698 = v616 << (uint(int32(6)) % 32)
	v699 = int32(0)
	v700 = int32(2)
	goto L180
L182:
	;
	v674 = v670 + v616<<(uint(int32(6))%32)
	v676 = v617 + int32(1)
	if v676 != int32(4) {
		goto L200
	} else {
		goto L201
	}
L183:
	;
	v670 = int32(62)
	goto L182
L184:
	;
	v631 = v621 + int32(1)
	v632 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v621))))
	v634 = v632 - int32(65)
	if base.Ui32(v634&int32(255)) <= base.Ui32(int32(25)) {
		v670 = v634
		goto L182
	} else {
		goto L186
	}
L185:
	;
	if v615 != 0 {
		goto L196
	} else {
		goto L197
	}
L186:
	;
	if base.Ui32((v632-int32(97))&int32(255)) <= base.Ui32(int32(25)) {
		goto L187
	} else {
		goto L188
	}
L187:
	;
	v670 = v632 - int32(71)
	goto L182
L188:
	;
	goto L189
L189:
	;
	if base.Ui32((v632-int32(48))&int32(255)) <= base.Ui32(int32(9)) {
		goto L190
	} else {
		goto L191
	}
L190:
	;
	v670 = (v632 + int32(4)) & int32(255)
	goto L182
L191:
	;
	goto L192
L192:
	;
	v658 = int32(-101)
	switch v632 - int32(9) {
	case 0, 1, 4, 23:
		goto L194
	default:
		v762 = v658
		goto L171
	case 34:
		goto L183
	case 38:
		v670 = int32(63)
		goto L182
	case 52:
		goto L193
	}
L193:
	;
	goto L185
L194:
	;
	if base.Ui32(v631) < base.Ui32(v609) {
		v621 = v631
		goto L184
	} else {
		goto L195
	}
L195:
	;
	v736 = v613
	v740 = v617
	goto L177
L196:
	;
	v670 = int32(0)
	goto L182
L197:
	;
	goto L198
L198:
	;
	switch v617 - int32(2) {
	case 0:
		goto L199
	case 1:
		goto L181
	default:
		v762 = v658
		goto L171
	}
L199:
	;
	v728 = v613
	v730 = int32(1)
	v731 = v616 << (uint(int32(6)) % 32)
	v732 = int32(3)
	goto L178
L200:
	;
	v728 = v613
	v730 = v615
	v731 = v674
	v732 = v676
	goto L178
L201:
	;
	goto L202
L202:
	;
	v680 = int32(base.Ui32(v674) >> (uint(int32(16)) % 32))
	*(*uint8)(unsafe.Add(mBase, uint32(v613))) = uint8(v680)
	v683 = base.B2i32(v615 == int32(0))
	if v615 == int32(0) {
		v698 = v674
		v699 = v683
		v700 = v615
		goto L180
	} else {
		goto L203
	}
L203:
	;
	if base.Ui32(int32(2)) <= base.Ui32(v615) {
		v698 = v674
		v699 = v683
		v700 = v615
		goto L180
	} else {
		goto L204
	}
L204:
	;
	v708 = v674
	v709 = int32(0)
	v710 = v615
	v713 = v613 + int32(1)
	goto L179
L205:
	;
	*(*uint8)(unsafe.Add(mBase, uint32(v713))) = uint8(v708)
	v725 = v713 + int32(1)
	goto L207
L206:
	;
	v725 = v713
	goto L207
L207:
	;
	v728 = v725
	v730 = v710
	v731 = v714
	v732 = int32(0)
	goto L178
L208:
	;
	v736 = v728
	v740 = v732
	goto L177
L209:
	;
	v746 = v736
	goto L174
L210:
	;
	if int32(0) <= v762 {
		goto L211
	} else {
		goto L212
	}
L211:
	;
	if v762 != 0 {
		goto L214
	} else {
		goto L215
	}
L212:
	;
	goto L213
L213:
	;
	v888 = v762
	goto L2
L214:
	;
	v773 = *(*int32)(unsafe.Add(mBase, uint32(l2)))
	v775 = v773
	v778 = int32(11994318)
	v780 = v762
	goto L217
L215:
	;
	v864 = int32(11994318)
	goto L216
L216:
	;
	if v864 != v588<<(uint(int32(8))%32)|v589<<(uint(int32(16))%32)|v587 {
		v888 = v17
		goto L2
	} else {
		goto L244
	}
L217:
	;
	v786 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v775))))
	v789 = v786<<(uint(int32(16))%32) ^ v778
	v791 = v789 << (uint(int32(1)) % 32)
	if v789&int32(8388608) != 0 {
		goto L219
	} else {
		goto L220
	}
L218:
	;
	v864 = v845 & int32(16777215)
	goto L216
L219:
	;
	v796 = v791 ^ int32(25578747)
	goto L221
L220:
	;
	v796 = v791
	goto L221
L221:
	;
	v798 = v796 << (uint(int32(1)) % 32)
	if v796&int32(8388608) != 0 {
		goto L222
	} else {
		goto L223
	}
L222:
	;
	v803 = v798 ^ int32(25578747)
	goto L224
L223:
	;
	v803 = v798
	goto L224
L224:
	;
	v805 = v803 << (uint(int32(1)) % 32)
	if v803&int32(8388608) != 0 {
		goto L225
	} else {
		goto L226
	}
L225:
	;
	v810 = v805 ^ int32(25578747)
	goto L227
L226:
	;
	v810 = v805
	goto L227
L227:
	;
	v812 = v810 << (uint(int32(1)) % 32)
	if v810&int32(8388608) != 0 {
		goto L228
	} else {
		goto L229
	}
L228:
	;
	v817 = v812 ^ int32(25578747)
	goto L230
L229:
	;
	v817 = v812
	goto L230
L230:
	;
	v819 = v817 << (uint(int32(1)) % 32)
	if v817&int32(8388608) != 0 {
		goto L231
	} else {
		goto L232
	}
L231:
	;
	v824 = v819 ^ int32(25578747)
	goto L233
L232:
	;
	v824 = v819
	goto L233
L233:
	;
	v826 = v824 << (uint(int32(1)) % 32)
	if v824&int32(8388608) != 0 {
		goto L234
	} else {
		goto L235
	}
L234:
	;
	v831 = v826 ^ int32(25578747)
	goto L236
L235:
	;
	v831 = v826
	goto L236
L236:
	;
	v833 = v831 << (uint(int32(1)) % 32)
	if v831&int32(8388608) != 0 {
		goto L237
	} else {
		goto L238
	}
L237:
	;
	v838 = v833 ^ int32(25578747)
	goto L239
L238:
	;
	v838 = v833
	goto L239
L239:
	;
	v840 = v838 << (uint(int32(1)) % 32)
	if v838&int32(8388608) != 0 {
		goto L240
	} else {
		goto L241
	}
L240:
	;
	v845 = v840 ^ int32(25578747)
	goto L242
L241:
	;
	v845 = v840
	goto L242
L242:
	;
	v846 = int32(1)
	v849 = v780 - v846
	if v849 != 0 {
		v775 = v775 + v846
		v778 = v845
		v780 = v849
		goto L217
	} else {
		goto L243
	}
L243:
	;
	goto L218
L244:
	;
	v866 = *(*int32)(unsafe.Add(mBase, uint32(l2)+4))
	*(*int32)(unsafe.Add(mBase, uint32(l2)+4)) = v866 + v762
	goto L213
L245:
	;
	F_errmsg_internal(m, int32(291078), int32(0))
	mBase = m.M
	v904 = m.ExcPending
	if v904 != 0 {
		goto L168
	} else {
		goto L246
	}
L246:
	;
	F_errfinish(m, int32(473595), int32(370), int32(395538))
	mBase = m.M
	v911 = m.ExcPending
	if v911 != 0 {
		goto L168
	} else {
		goto L247
	}
L247:
	;
	base.Wasm_trap_unreachable()
	for {
	}
}
func F_pgp_armor_encode(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32, l5 int32) {
	mBase := m.M
	_ = mBase
	var v13 int32
	_ = v13
	var v15 int32
	_ = v15
	var v17 int32
	_ = v17
	var v24 int32
	_ = v24
	var v26 int32
	_ = v26
	var v27 int32
	_ = v27
	var v30 int32
	_ = v30
	var v33 int32
	_ = v33
	var v35 int32
	_ = v35
	var v40 int32
	_ = v40
	var v42 int32
	_ = v42
	var v47 int32
	_ = v47
	var v49 int32
	_ = v49
	var v54 int32
	_ = v54
	var v56 int32
	_ = v56
	var v61 int32
	_ = v61
	var v63 int32
	_ = v63
	var v68 int32
	_ = v68
	var v70 int32
	_ = v70
	var v75 int32
	_ = v75
	var v77 int32
	_ = v77
	var v82 int32
	_ = v82
	var v84 int32
	_ = v84
	var v89 int32
	_ = v89
	var v90 int32
	_ = v90
	var v93 int32
	_ = v93
	var v104 int32
	_ = v104
	var v110 int32
	_ = v110
	var v120 int32
	_ = v120
	var v127 int32
	_ = v127
	var v129 int32
	_ = v129
	var v131 int32
	_ = v131
	var v136 int32
	_ = v136
	var v138 int32
	_ = v138
	var v154 int32
	_ = v154
	var v155 int32
	_ = v155
	var v157 int32
	_ = v157
	var v161 int32
	_ = v161
	var v164 int32
	_ = v164
	var v166 int32
	_ = v166
	var v167 int32
	_ = v167
	var v168 int32
	_ = v168
	var v169 int32
	_ = v169
	var v176 int32
	_ = v176
	var v179 int32
	_ = v179
	var v182 int32
	_ = v182
	var v183 int32
	_ = v183
	var v185 int32
	_ = v185
	var v188 int32
	_ = v188
	var v192 int32
	_ = v192
	var v198 int32
	_ = v198
	var v201 int32
	_ = v201
	var v208 int32
	_ = v208
	var v215 int32
	_ = v215
	var v222 int32
	_ = v222
	var v228 int32
	_ = v228
	var v229 int32
	_ = v229
	var v231 int32
	_ = v231
	var v233 int32
	_ = v233
	var v239 int32
	_ = v239
	var v240 int32
	_ = v240
	var v242 int32
	_ = v242
	var v249 int32
	_ = v249
	var v252 int32
	_ = v252
	var v259 int32
	_ = v259
	var v270 int32
	_ = v270
	var v271 int32
	_ = v271
	var v272 int32
	_ = v272
	var v283 int32
	_ = v283
	var v289 int32
	_ = v289
	var v291 int32
	_ = v291
	var v292 int32
	_ = v292
	var v294 int32
	_ = v294
	var v298 int32
	_ = v298
	var v303 int32
	_ = v303
	var v306 int32
	_ = v306
	var v313 int32
	_ = v313
	var v315 int32
	_ = v315
	var v321 int32
	_ = v321
	var v323 int32
	_ = v323
	var v329 int32
	_ = v329
	var v331 int32
	_ = v331
	var v335 int32
	_ = v335
	var v337 int32
	_ = v337
	var v341 int32
	_ = v341
	var v348 int32
	_ = v348
	var v354 int32
	_ = v354
	var v361 int32
	_ = v361
	v13 = m.G0
	v15 = v13 - int32(16)
	m.G0 = v15
	v17 = int32(11994318)
	if l1 != 0 {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	v24 = l0
	v26 = v17
	v27 = l1
	goto L4
L2:
	;
	v104 = v17
	goto L3
L3:
	;
	F_appendStringInfoString(m, l2, int32(715335))
	mBase = m.M
	v110 = m.ExcPending
	if v110 != 0 {
		goto L31
	} else {
		goto L32
	}
L4:
	;
	v30 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v24))))
	v33 = v30<<(uint(int32(16))%32) ^ v26
	v35 = v33 << (uint(int32(1)) % 32)
	if v33&int32(8388608) != 0 {
		goto L6
	} else {
		goto L7
	}
L5:
	;
	v104 = v89 & int32(16777215)
	goto L3
L6:
	;
	v40 = v35 ^ int32(25578747)
	goto L8
L7:
	;
	v40 = v35
	goto L8
L8:
	;
	v42 = v40 << (uint(int32(1)) % 32)
	if v40&int32(8388608) != 0 {
		goto L9
	} else {
		goto L10
	}
L9:
	;
	v47 = v42 ^ int32(25578747)
	goto L11
L10:
	;
	v47 = v42
	goto L11
L11:
	;
	v49 = v47 << (uint(int32(1)) % 32)
	if v47&int32(8388608) != 0 {
		goto L12
	} else {
		goto L13
	}
L12:
	;
	v54 = v49 ^ int32(25578747)
	goto L14
L13:
	;
	v54 = v49
	goto L14
L14:
	;
	v56 = v54 << (uint(int32(1)) % 32)
	if v54&int32(8388608) != 0 {
		goto L15
	} else {
		goto L16
	}
L15:
	;
	v61 = v56 ^ int32(25578747)
	goto L17
L16:
	;
	v61 = v56
	goto L17
L17:
	;
	v63 = v61 << (uint(int32(1)) % 32)
	if v61&int32(8388608) != 0 {
		goto L18
	} else {
		goto L19
	}
L18:
	;
	v68 = v63 ^ int32(25578747)
	goto L20
L19:
	;
	v68 = v63
	goto L20
L20:
	;
	v70 = v68 << (uint(int32(1)) % 32)
	if v68&int32(8388608) != 0 {
		goto L21
	} else {
		goto L22
	}
L21:
	;
	v75 = v70 ^ int32(25578747)
	goto L23
L22:
	;
	v75 = v70
	goto L23
L23:
	;
	v77 = v75 << (uint(int32(1)) % 32)
	if v75&int32(8388608) != 0 {
		goto L24
	} else {
		goto L25
	}
L24:
	;
	v82 = v77 ^ int32(25578747)
	goto L26
L25:
	;
	v82 = v77
	goto L26
L26:
	;
	v84 = v82 << (uint(int32(1)) % 32)
	if v82&int32(8388608) != 0 {
		goto L27
	} else {
		goto L28
	}
L27:
	;
	v89 = v84 ^ int32(25578747)
	goto L29
L28:
	;
	v89 = v84
	goto L29
L29:
	;
	v90 = int32(1)
	v93 = v27 - v90
	if v93 != 0 {
		v24 = v24 + v90
		v26 = v89
		v27 = v93
		goto L4
	} else {
		goto L30
	}
L30:
	;
	goto L5
L31:
	;
	return
L32:
	;
	if int32(0) < l3 {
		goto L33
	} else {
		goto L34
	}
L33:
	;
	v120 = int32(0)
	goto L36
L34:
	;
	goto L35
L35:
	;
	F_appendStringInfoChar(m, l2, int32(10))
	mBase = m.M
	v154 = m.ExcPending
	if v154 != 0 {
		goto L31
	} else {
		goto L40
	}
L36:
	;
	v127 = v120 << (uint(int32(2)) % 32)
	v129 = *(*int32)(unsafe.Add(mBase, uint32(l4+v127)))
	v131 = *(*int32)(unsafe.Add(mBase, uint32(l5+v127)))
	*(*int32)(unsafe.Add(mBase, uint32(v15)+4)) = v131
	*(*int32)(unsafe.Add(mBase, uint32(v15))) = v129
	F_appendStringInfo(m, l2, int32(707863), v15)
	mBase = m.M
	v136 = m.ExcPending
	if v136 != 0 {
		goto L31
	} else {
		goto L38
	}
L37:
	;
	goto L35
L38:
	;
	v138 = v120 + int32(1)
	if v138 != l3 {
		v120 = v138
		goto L36
	} else {
		goto L39
	}
L39:
	;
	goto L37
L40:
	;
	v155 = int32(2)
	v157 = base.I32_div_u_s(l1, int32(57))
	v161 = base.I32_div_u_s(l1+v155, int32(3))
	v164 = v157 + v161<<(uint(v155)%32)
	F_enlargeStringInfo(m, l2, v164)
	mBase = m.M
	v166 = m.ExcPending
	if v166 != 0 {
		goto L31
	} else {
		goto L41
	}
L41:
	;
	v167 = *(*int32)(unsafe.Add(mBase, uint32(l2)))
	v168 = *(*int32)(unsafe.Add(mBase, uint32(l2)+4))
	v169 = v167 + v168
	if l1 == int32(0) {
		v283 = v169
		goto L42
	} else {
		goto L43
	}
L42:
	;
	v289 = v283 - v169
	if base.Ui32(v289) <= base.Ui32(v164) {
		goto L57
	} else {
		goto L58
	}
L43:
	;
	v176 = l0
	v179 = v169 + int32(76)
	v182 = v169
	v183 = int32(0)
	v185 = v155
	goto L44
L44:
	;
	v188 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v176))))
	v192 = v188<<(uint(v185<<(uint(int32(3))%32))%32) | v183
	if int32(0) < v185 {
		goto L46
	} else {
		goto L47
	}
L45:
	;
	if v231 == int32(2) {
		v283 = v240
		goto L42
	} else {
		goto L53
	}
L46:
	;
	v228 = v182
	v229 = v192
	v231 = v185 - int32(1)
	goto L48
L47:
	;
	v198 = int32(63)
	v201 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v192&v198)+uint32(_consts[1528]))))
	*(*uint8)(unsafe.Add(mBase, uint32(v182)+3)) = uint8(v201)
	v208 = int32(*(*uint8)(unsafe.Add(mBase, uint32(int32(base.Ui32(v192)>>(uint(int32(6))%32))&v198)+uint32(_consts[1528]))))
	*(*uint8)(unsafe.Add(mBase, uint32(v182)+2)) = uint8(v208)
	v215 = int32(*(*uint8)(unsafe.Add(mBase, uint32(int32(base.Ui32(v192)>>(uint(int32(12))%32))&v198)+uint32(_consts[1528]))))
	*(*uint8)(unsafe.Add(mBase, uint32(v182)+1)) = uint8(v215)
	v222 = int32(*(*uint8)(unsafe.Add(mBase, uint32(int32(base.Ui32(v192)>>(uint(int32(18))%32))&v198)+uint32(_consts[1528]))))
	*(*uint8)(unsafe.Add(mBase, uint32(v182))) = uint8(v222)
	v228 = v182 + int32(4)
	v229 = int32(0)
	v231 = int32(2)
	goto L48
L48:
	;
	if base.Ui32(v179) <= base.Ui32(v228) {
		goto L49
	} else {
		goto L50
	}
L49:
	;
	v233 = int32(10)
	*(*uint8)(unsafe.Add(mBase, uint32(v228))) = uint8(v233)
	v239 = v228 + int32(77)
	v240 = v228 + int32(1)
	goto L51
L50:
	;
	v239 = v179
	v240 = v228
	goto L51
L51:
	;
	v242 = v176 + int32(1)
	if base.Ui32(v242) < base.Ui32(l0+l1) {
		v176 = v242
		v179 = v239
		v182 = v240
		v183 = v229
		v185 = v231
		goto L44
	} else {
		goto L52
	}
L52:
	;
	goto L45
L53:
	;
	v249 = int32(63)
	v252 = int32(*(*uint8)(unsafe.Add(mBase, uint32(int32(base.Ui32(v229)>>(uint(int32(12))%32))&v249)+uint32(_consts[1528]))))
	*(*uint8)(unsafe.Add(mBase, uint32(v240)+1)) = uint8(v252)
	v259 = int32(*(*uint8)(unsafe.Add(mBase, uint32(int32(base.Ui32(v229)>>(uint(int32(18))%32))&v249)+uint32(_consts[1528]))))
	*(*uint8)(unsafe.Add(mBase, uint32(v240))) = uint8(v259)
	if v231 == int32(0) {
		goto L54
	} else {
		goto L55
	}
L54:
	;
	v270 = int32(*(*uint8)(unsafe.Add(mBase, uint32(int32(base.Ui32(v229)>>(uint(int32(6))%32))&int32(63))+uint32(_consts[1528]))))
	v271 = v270
	goto L56
L55:
	;
	v271 = int32(61)
	goto L56
L56:
	;
	v272 = int32(61)
	*(*uint8)(unsafe.Add(mBase, uint32(v240)+3)) = uint8(v272)
	*(*uint8)(unsafe.Add(mBase, uint32(v240)+2)) = uint8(v271)
	v283 = v240 + int32(4)
	goto L42
L57:
	;
	v291 = *(*int32)(unsafe.Add(mBase, uint32(l2)+4))
	v292 = v291 + v289
	*(*int32)(unsafe.Add(mBase, uint32(l2)+4)) = v292
	v294 = *(*int32)(unsafe.Add(mBase, uint32(l2)))
	v298 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v294+v292-int32(1)))))
	if v298 != int32(10) {
		goto L60
	} else {
		goto L61
	}
L58:
	;
	goto L59
L59:
	;
	F_errstart_cold(m, int32(22), int32(0))
	mBase = m.M
	v348 = m.ExcPending
	if v348 != 0 {
		goto L31
	} else {
		goto L70
	}
L60:
	;
	F_appendStringInfoChar(m, l2, int32(10))
	mBase = m.M
	v303 = m.ExcPending
	if v303 != 0 {
		goto L31
	} else {
		goto L63
	}
L61:
	;
	goto L62
L62:
	;
	F_appendStringInfoChar(m, l2, int32(61))
	mBase = m.M
	v306 = m.ExcPending
	if v306 != 0 {
		goto L31
	} else {
		goto L64
	}
L63:
	;
	goto L62
L64:
	;
	v313 = int32(*(*int8)(unsafe.Add(mBase, uint32(int32(base.Ui32(v104)>>(uint(int32(18))%32)))+uint32(_consts[1528]))))
	F_appendStringInfoChar(m, l2, v313)
	mBase = m.M
	v315 = m.ExcPending
	if v315 != 0 {
		goto L31
	} else {
		goto L65
	}
L65:
	;
	v321 = int32(*(*int8)(unsafe.Add(mBase, uint32(int32(base.Ui32(v104)>>(uint(int32(12))%32))&int32(63))+uint32(_consts[1528]))))
	F_appendStringInfoChar(m, l2, v321)
	mBase = m.M
	v323 = m.ExcPending
	if v323 != 0 {
		goto L31
	} else {
		goto L66
	}
L66:
	;
	v329 = int32(*(*int8)(unsafe.Add(mBase, uint32(int32(base.Ui32(v104)>>(uint(int32(6))%32))&int32(63))+uint32(_consts[1528]))))
	F_appendStringInfoChar(m, l2, v329)
	mBase = m.M
	v331 = m.ExcPending
	if v331 != 0 {
		goto L31
	} else {
		goto L67
	}
L67:
	;
	v335 = int32(*(*int8)(unsafe.Add(mBase, uint32(v104&int32(63))+uint32(_consts[1528]))))
	F_appendStringInfoChar(m, l2, v335)
	mBase = m.M
	v337 = m.ExcPending
	if v337 != 0 {
		goto L31
	} else {
		goto L68
	}
L68:
	;
	F_appendStringInfoString(m, l2, int32(715364))
	mBase = m.M
	v341 = m.ExcPending
	if v341 != 0 {
		goto L31
	} else {
		goto L69
	}
L69:
	;
	m.G0 = v15 + int32(16)
	return
L70:
	;
	F_errmsg_internal(m, int32(291041), int32(0))
	mBase = m.M
	v354 = m.ExcPending
	if v354 != 0 {
		goto L31
	} else {
		goto L71
	}
L71:
	;
	F_errfinish(m, int32(473595), int32(227), int32(395435))
	mBase = m.M
	v361 = m.ExcPending
	if v361 != 0 {
		goto L31
	} else {
		goto L72
	}
L72:
	;
	base.Wasm_trap_unreachable()
	for {
	}
}
func F_pgp_cfb_create(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32, l5 int32) int32 {
	mBase := m.M
	_ = mBase
	var v8 int32
	_ = v8
	var v10 int32
	_ = v10
	var v14 int32
	_ = v14
	var v17 int32
	_ = v17
	var v20 int32
	_ = v20
	var v22 int32
	_ = v22
	var v23 int32
	_ = v23
	var v24 int32
	_ = v24
	var v27 int32
	_ = v27
	var v28 int32
	_ = v28
	var v30 int32
	_ = v30
	var v32 int32
	_ = v32
	var v33 int32
	_ = v33
	var v34 int32
	_ = v34
	var v36 int32
	_ = v36
	var v37 int32
	_ = v37
	var v38 int32
	_ = v38
	var v43 int32
	_ = v43
	var v47 int32
	_ = v47
	v8 = m.G0
	v10 = v8 - int32(16)
	m.G0 = v10
	v14 = F_pgp_load_cipher(m, l1, v10+int32(12))
	mBase = m.M
	v17 = m.ExcPending
	if v17 != 0 {
		return int32(0)
	} else {
		if v14 < int32(0) {
			v47 = v14
			m.G0 = v10 + int32(16)
			return v47
		} else {
			v20 = *(*int32)(unsafe.Add(mBase, uint32(v10)+12))
			v22 = *(*int32)(unsafe.Add(mBase, uint32(v20)+12))
			v23 = m.T0[v22].(func(*base.Module, int32, int32, int32, int32) int32)(m, v20, l2, l3, int32(0))
			mBase = m.M
			v24 = m.ExcPending
			if v24 != 0 {
				return int32(0)
			} else {
				if v23 < int32(0) {
					v27 = *(*int32)(unsafe.Add(mBase, uint32(v10)+12))
					v28 = *(*int32)(unsafe.Add(mBase, uint32(v27)+24))
					m.T0[v28].(func(*base.Module, int32))(m, v27)
					mBase = m.M
					v30 = m.ExcPending
					if v30 != 0 {
						return int32(0)
					} else {
						v47 = v23
						m.G0 = v10 + int32(16)
						return v47
					}
				} else {
					v32 = F_palloc0(m, int32(116))
					mBase = m.M
					v33 = m.ExcPending
					if v33 != 0 {
						return int32(0)
					} else {
						v34 = *(*int32)(unsafe.Add(mBase, uint32(v10)+12))
						*(*int32)(unsafe.Add(mBase, uint32(v32))) = v34
						v36 = *(*int32)(unsafe.Add(mBase, uint32(v34)))
						v37 = m.T0[v36].(func(*base.Module, int32) int32)(m, v34)
						mBase = m.M
						v38 = m.ExcPending
						if v38 != 0 {
							return int32(0)
						} else {
							*(*int32)(unsafe.Add(mBase, uint32(v32)+16)) = l4
							*(*int32)(unsafe.Add(mBase, uint32(v32)+4)) = v37
							if l5 != 0 {
								if v37 != 0 {
									v43 = F__emscripten_memcpy_bulkmem(m, v32+int32(20), l5, v37)
									mBase = m.M
								} else {
								}
							} else {
							}
							*(*int32)(unsafe.Add(mBase, uint32(l0))) = v32
							v47 = int32(0)
							m.G0 = v10 + int32(16)
							return v47
						}
					}
				}
			}
		}
	}
}
func F_pgp_create_pkt_writer(m *base.Module, l0 int32, l1 int32, l2 int32) int32 {
	mBase := m.M
	_ = mBase
	var v5 int32
	_ = v5
	var v7 int32
	_ = v7
	var v10 int32
	_ = v10
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
	v5 = m.G0
	v7 = v5 - int32(16)
	m.G0 = v7
	v10 = l1 | int32(192)
	*(*uint8)(unsafe.Add(mBase, uint32(v7)+15)) = uint8(v10)
	v15 = F_pushf_write(m, l0, v7+int32(15), int32(1))
	mBase = m.M
	v18 = m.ExcPending
	if v18 != 0 {
		return int32(0)
	} else {
		if int32(0) <= v15 {
			v23 = F_pushf_create(m, l2, int32(4336412), int32(0), l0)
			mBase = m.M
			v24 = m.ExcPending
			if v24 != 0 {
				return int32(0)
			} else {
				v25 = v23
				m.G0 = v7 + int32(16)
				return v25
			}
		} else {
			v25 = v15
			m.G0 = v7 + int32(16)
			return v25
		}
	}
}
func F_pgp_decompress_filter(m *base.Module, l0 int32, l1 int32, l2 int32) int32 {
	var v5 int32
	_ = v5
	var v8 int32
	_ = v8
	v5 = F_pullf_create(m, l0, int32(4336320), l1, l2)
	v8 = m.ExcPending
	if v8 != 0 {
		return int32(0)
	} else {
		return v5
	}
}
func F_pgp_free(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v3 int32
	_ = v3
	var v7 int32
	_ = v7
	var v10 int32
	_ = v10
	var v12 int32
	_ = v12
	v3 = *(*int32)(unsafe.Add(mBase, uint32(l0)+120))
	if v3 != 0 {
		F_pgp_key_free(m, v3)
		mBase = m.M
		v7 = m.ExcPending
		if v7 != 0 {
			return int32(0)
		} else {
			v10 = F___memset(m, l0, int32(0), int32(168))
			mBase = m.M
			F_pfree(m, l0)
			mBase = m.M
			v12 = m.ExcPending
			if v12 != 0 {
				return int32(0)
			} else {
				return int32(0)
			}
		}
	} else {
		v10 = F___memset(m, l0, int32(0), int32(168))
		mBase = m.M
		F_pfree(m, l0)
		mBase = m.M
		v12 = m.ExcPending
		if v12 != 0 {
			return int32(0)
		} else {
			return int32(0)
		}
	}
}
func F_pgp_get_cipher_block_size(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v2 int32
	_ = v2
	var v4 int32
	_ = v4
	var v17 int32
	_ = v17
	var v18 int32
	_ = v18
	var v19 int32
	_ = v19
	v2 = int32(0)
	v4 = l0 - int32(2)
	if base.Ui32(int32(8)) < base.Ui32(v4) {
		v19 = v2
	} else {
		if int32(base.Ui32(int32(487))>>(uint(v4)%32))&int32(1) == int32(0) {
			v19 = v2
		} else {
			v17 = *(*int32)(unsafe.Add(mBase, uint32(v4<<(uint(int32(2))%32))+uint32(_consts[1553])))
			v18 = *(*int32)(unsafe.Add(mBase, uint32(v17)+16))
			v19 = v18
		}
	}
	return v19
}
func F_pgp_load_cipher(m *base.Module, l0 int32, l1 int32) int32 {
	mBase := m.M
	_ = mBase
	var v4 int32
	_ = v4
	var v6 int32
	_ = v6
	var v21 int32
	_ = v21
	var v22 int32
	_ = v22
	var v23 int32
	_ = v23
	var v26 int32
	_ = v26
	var v27 int32
	_ = v27
	var v28 int32
	_ = v28
	v4 = int32(-100)
	v6 = l0 - int32(2)
	if base.Ui32(int32(8)) < base.Ui32(v6) {
		v28 = v4
		return v28
	} else {
		if int32(base.Ui32(int32(487))>>(uint(v6)%32))&int32(1) == int32(0) {
			v28 = v4
			return v28
		} else {
			v21 = *(*int32)(unsafe.Add(mBase, uint32(v6<<(uint(int32(2))%32))+uint32(_consts[1553])))
			v22 = *(*int32)(unsafe.Add(mBase, uint32(v21)+8))
			v23 = F_px_find_cipher(m, v22, l1)
			mBase = m.M
			v26 = m.ExcPending
			if v26 != 0 {
				return int32(0)
			} else {
				if v23 != 0 {
					v27 = int32(-103)
				} else {
					v27 = int32(0)
				}
				v28 = v27
				return v28
			}
		}
	}
}
func F_pgp_mpi_alloc(m *base.Module, l0 int32, l1 int32) int32 {
	mBase := m.M
	_ = mBase
	var v6 int32
	_ = v6
	var v8 int32
	_ = v8
	var v17 int32
	_ = v17
	var v22 int32
	_ = v22
	var v25 int32
	_ = v25
	var v26 int32
	_ = v26
	var v36 int32
	_ = v36
	v6 = m.G0
	v8 = v6 - int32(16)
	m.G0 = v8
	if base.Ui32(int32(65536)) <= base.Ui32(l0) {
		*(*int32)(unsafe.Add(mBase, uint32(v8))) = l0
		F_px_debug(m, int32(446018), v8)
		mBase = m.M
		v17 = m.ExcPending
		if v17 != 0 {
			return int32(0)
		} else {
			v36 = int32(-100)
			m.G0 = v8 + int32(16)
			return v36
		}
	} else {
		v22 = int32(base.Ui32(l0+int32(7)) >> (uint(int32(3)) % 32))
		v25 = F_palloc(m, v22+int32(12))
		mBase = m.M
		v26 = m.ExcPending
		if v26 != 0 {
			return int32(0)
		} else {
			*(*int32)(unsafe.Add(mBase, uint32(v25)+8)) = v22
			*(*int32)(unsafe.Add(mBase, uint32(v25)+4)) = l0
			*(*int32)(unsafe.Add(mBase, uint32(v25))) = v25 + int32(12)
			*(*int32)(unsafe.Add(mBase, uint32(l1))) = v25
			v36 = int32(0)
			m.G0 = v8 + int32(16)
			return v36
		}
	}
}
func F_pgp_mpi_cksum(m *base.Module, l0 int32, l1 int32) int32 {
	mBase := m.M
	_ = mBase
	var v3 int32
	_ = v3
	var v9 int32
	_ = v9
	var v15 int32
	_ = v15
	var v16 int32
	_ = v16
	var v20 int32
	_ = v20
	var v21 int32
	_ = v21
	var v28 int32
	_ = v28
	var v29 int32
	_ = v29
	var v34 int32
	_ = v34
	var v36 int32
	_ = v36
	var v37 int32
	_ = v37
	var v39 int32
	_ = v39
	var v41 int32
	_ = v41
	var v43 int32
	_ = v43
	var v44 int32
	_ = v44
	var v45 int32
	_ = v45
	var v46 int32
	_ = v46
	var v48 int32
	_ = v48
	var v50 int32
	_ = v50
	var v51 int32
	_ = v51
	var v60 int32
	_ = v60
	var v61 int32
	_ = v61
	var v65 int32
	_ = v65
	var v69 int32
	_ = v69
	var v70 int32
	_ = v70
	var v71 int32
	_ = v71
	var v74 int32
	_ = v74
	var v76 int32
	_ = v76
	v3 = int32(0)
	v9 = *(*int32)(unsafe.Add(mBase, uint32(l1)+4))
	v15 = v9>>(uint(int32(8))%32) + l0 + v9&int32(255)
	v16 = *(*int32)(unsafe.Add(mBase, uint32(l1)+8))
	if v16 <= v3 {
		v76 = v15
	} else {
		v20 = v16 & int32(3)
		v21 = *(*int32)(unsafe.Add(mBase, uint32(l1)))
		if base.Ui32(v16) < base.Ui32(int32(4)) {
			v50 = v15
			v51 = int32(0)
		} else {
			v28 = v15
			v29 = int32(0)
			v34 = v3
			for {
				v36 = v29 + v21
				v37 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v36))))
				v39 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v36)+1)))
				v41 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v36)+2)))
				v43 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v36)+3)))
				v44 = v28 + v37 + v39 + v41 + v43
				v45 = int32(4)
				v46 = v29 + v45
				v48 = v34 + v45
				if v48 != v16&int32(2147483644) {
					v28 = v44
					v29 = v46
					v34 = v48
					continue
				} else {
					break
				}
				break
			}
			v50 = v44
			v51 = v46
		}
		if v20 == int32(0) {
			v76 = v50
		} else {
			v60 = v50
			v61 = v51
			v65 = v3
			for {
				v69 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v61+v21))))
				v70 = v60 + v69
				v71 = int32(1)
				v74 = v65 + v71
				if v74 != v20 {
					v60 = v70
					v61 = v61 + v71
					v65 = v74
					continue
				} else {
					break
				}
				break
			}
			v76 = v70
		}
	}
	return v76 & int32(65535)
}
func F_pgp_parse_pkt_hdr(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) int32 {
	mBase := m.M
	_ = mBase
	var v8 int32
	_ = v8
	var v10 int32
	_ = v10
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
	var v31 int32
	_ = v31
	var v35 int32
	_ = v35
	var v37 int32
	_ = v37
	var v38 int32
	_ = v38
	var v42 int32
	_ = v42
	var v44 int32
	_ = v44
	var v45 int32
	_ = v45
	var v50 int32
	_ = v50
	var v54 int32
	_ = v54
	var v55 int32
	_ = v55
	var v58 int32
	_ = v58
	var v64 int32
	_ = v64
	var v65 int32
	_ = v65
	var v68 int32
	_ = v68
	var v75 int32
	_ = v75
	var v76 int32
	_ = v76
	var v79 int32
	_ = v79
	var v83 int32
	_ = v83
	var v84 int32
	_ = v84
	var v87 int32
	_ = v87
	var v91 int32
	_ = v91
	var v92 int32
	_ = v92
	var v95 int32
	_ = v95
	var v96 int32
	_ = v96
	var v104 int32
	_ = v104
	var v110 int32
	_ = v110
	var v116 int32
	_ = v116
	var v124 int32
	_ = v124
	v8 = m.G0
	v10 = v8 - int32(16)
	m.G0 = v10
	v15 = F_pullf_read(m, l0, int32(1), v10+int32(4))
	mBase = m.M
	v18 = m.ExcPending
	if v18 != 0 {
		return int32(0)
	} else {
		if v15 < int32(0) {
			v124 = v15
			m.G0 = v10 + int32(16)
			return v124
		} else {
			if v15 == int32(0) {
				v124 = int32(0)
				m.G0 = v10 + int32(16)
				return v124
			} else {
				v23 = *(*int32)(unsafe.Add(mBase, uint32(v10)+4))
				v24 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v23))))
				v25 = base.I32_extend8_s(v24)
				if int32(0) <= v25 {
					F_px_debug(m, int32(218326), int32(0))
					mBase = m.M
					v31 = m.ExcPending
					if v31 != 0 {
						return int32(0)
					} else {
						v124 = int32(-100)
						m.G0 = v10 + int32(16)
						return v124
					}
				} else {
					if v24&int32(64) != 0 {
						v35 = v25 & int32(63)
						*(*uint8)(unsafe.Add(mBase, uint32(l1))) = uint8(v35)
						v37 = F_parse_new_len(m, l0, l2)
						mBase = m.M
						v38 = m.ExcPending
						if v38 != 0 {
							return int32(0)
						} else {
							v124 = v37
							m.G0 = v10 + int32(16)
							return v124
						}
					} else {
						v42 = int32(base.Ui32(v25)>>(uint(int32(2))%32)) & int32(15)
						*(*uint8)(unsafe.Add(mBase, uint32(l1))) = uint8(v42)
						v44 = int32(3)
						v45 = v24 & v44
						if v45 == v44 {
							if l3 != 0 {
								v50 = int32(3)
							} else {
								v50 = int32(-100)
							}
							v124 = v50
							m.G0 = v10 + int32(16)
							return v124
						} else {
							v54 = F_pullf_read_fixed(m, l0, int32(1), v10+int32(15))
							mBase = m.M
							v55 = m.ExcPending
							if v55 != 0 {
								return int32(0)
							} else {
								if v54 < int32(0) {
									v124 = v54
									m.G0 = v10 + int32(16)
									return v124
								} else {
									v58 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v10)+15)))
									switch v45 - int32(1) {
									case 0:
										v64 = F_pullf_read_fixed(m, l0, int32(1), v10+int32(14))
										mBase = m.M
										v65 = m.ExcPending
										if v65 != 0 {
											return int32(0)
										} else {
											if v64 < int32(0) {
												v124 = v64
											} else {
												v68 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v10)+14)))
												v116 = v68 | v58<<(uint(int32(8))%32)
												*(*int32)(unsafe.Add(mBase, uint32(l2))) = v116
												v124 = int32(1)
											}
											m.G0 = v10 + int32(16)
											return v124
										}
									case 1:
										v75 = F_pullf_read_fixed(m, l0, int32(1), v10+int32(13))
										mBase = m.M
										v76 = m.ExcPending
										if v76 != 0 {
											return int32(0)
										} else {
											if v75 < int32(0) {
												v124 = v75
												m.G0 = v10 + int32(16)
												return v124
											} else {
												v79 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v10)+13)))
												v83 = F_pullf_read_fixed(m, l0, int32(1), v10+int32(12))
												mBase = m.M
												v84 = m.ExcPending
												if v84 != 0 {
													return int32(0)
												} else {
													if v83 < int32(0) {
														v124 = v83
														m.G0 = v10 + int32(16)
														return v124
													} else {
														v87 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v10)+12)))
														v91 = F_pullf_read_fixed(m, l0, int32(1), v10+int32(11))
														mBase = m.M
														v92 = m.ExcPending
														if v92 != 0 {
															return int32(0)
														} else {
															if v91 < int32(0) {
																v124 = v91
																m.G0 = v10 + int32(16)
																return v124
															} else {
																v95 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v10)+11)))
																v96 = int32(8)
																v104 = v95 | (v79<<(uint(v96)%32)|v58<<(uint(int32(16))%32)|v87)<<(uint(v96)%32)
																if base.Ui32(v104) < base.Ui32(int32(16777217)) {
																	v116 = v104
																	*(*int32)(unsafe.Add(mBase, uint32(l2))) = v116
																	v124 = int32(1)
																	m.G0 = v10 + int32(16)
																	return v124
																} else {
																	F_px_debug(m, int32(306784), int32(0))
																	mBase = m.M
																	v110 = m.ExcPending
																	if v110 != 0 {
																		return int32(0)
																	} else {
																		v124 = int32(-100)
																		m.G0 = v10 + int32(16)
																		return v124
																	}
																}
															}
														}
													}
												}
											}
										}
									default:
										v116 = v58
										*(*int32)(unsafe.Add(mBase, uint32(l2))) = v116
										v124 = int32(1)
										m.G0 = v10 + int32(16)
										return v124
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
func F_pgp_s2k_fill(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) int32 {
	mBase := m.M
	_ = mBase
	var v2 int32
	_ = v2
	var v3 int32
	_ = v3
	var v7 int32
	_ = v7
	var v9 int32
	_ = v9
	var v14 int32
	_ = v14
	var v15 int32
	_ = v15
	var v24 int32
	_ = v24
	var v26 int32
	_ = v26
	var v32 int32
	_ = v32
	var v38 int32
	_ = v38
	var v39 int32
	_ = v39
	var v44 int32
	_ = v44
	var v48 int32
	_ = v48
	var v53 int32
	_ = v53
	var v58 int32
	_ = v58
	var v60 int32
	_ = v60
	var v65 int32
	_ = v65
	var v70 int32
	_ = v70
	var v71 int32
	_ = v71
	var v75 int32
	_ = v75
	var v79 int32
	_ = v79
	var v81 int32
	_ = v81
	var v87 int32
	_ = v87
	var v93 int32
	_ = v93
	var v94 int32
	_ = v94
	var v99 int32
	_ = v99
	var v103 int32
	_ = v103
	var v108 int32
	_ = v108
	var v113 int32
	_ = v113
	var v115 int32
	_ = v115
	var v120 int32
	_ = v120
	var v130 int32
	_ = v130
	var v134 int32
	_ = v134
	var v136 int32
	_ = v136
	var v142 int32
	_ = v142
	var v148 int32
	_ = v148
	var v149 int32
	_ = v149
	var v154 int32
	_ = v154
	var v158 int32
	_ = v158
	var v163 int32
	_ = v163
	var v168 int32
	_ = v168
	var v170 int32
	_ = v170
	var v175 int32
	_ = v175
	var v186 int32
	_ = v186
	var v198 int32
	_ = v198
	var v202 int32
	_ = v202
	var v210 int32
	_ = v210
	var v215 int32
	_ = v215
	var v224 int32
	_ = v224
	var v229 int32
	_ = v229
	v2 = l1
	v3 = l2
	v7 = m.G0
	v9 = v7 - int32(16)
	m.G0 = v9
	*(*uint8)(unsafe.Add(mBase, uint32(l0)+1)) = uint8(v3)
	*(*uint8)(unsafe.Add(mBase, uint32(l0))) = uint8(v2)
	v14 = v2 & int32(255)
	switch v14 {
	case 0:
		v229 = v14
		goto L1
	case 1:
		goto L6
	default:
		goto L4
	case 3:
		goto L5
	}
L1:
	;
	m.G0 = v9 + int32(16)
	return v229
L2:
	;
	*(*uint8)(unsafe.Add(mBase, uint32(l0)+10)) = uint8(v224)
	v229 = int32(0)
	goto L1
L3:
	;
	v215 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v9)+15)))
	v224 = v215&int32(31) | int32(96)
	goto L2
L4:
	;
	v229 = int32(-121)
	goto L1
L5:
	;
	v71 = int32(-17)
	v75 = int32(0)
	v79 = m.G0
	v81 = v79 - int32(16)
	m.G0 = v81
	*(*int32)(unsafe.Add(mBase, uint32(v81))) = v75
	v87 = F_open(m, int32(275211), v75, v81)
	mBase = m.M
	if v87 != int32(-1) {
		goto L24
	} else {
		goto L25
	}
L6:
	;
	v15 = int32(0)
	v24 = m.G0
	v26 = v24 - int32(16)
	m.G0 = v26
	*(*int32)(unsafe.Add(mBase, uint32(v26))) = v15
	v32 = F_open(m, int32(275211), v15, v26)
	mBase = m.M
	if v32 != int32(-1) {
		goto L8
	} else {
		goto L9
	}
L7:
	;
	if v65 != 0 {
		goto L20
	} else {
		goto L21
	}
L8:
	;
	goto L12
L9:
	;
	v65 = v15
	goto L10
L10:
	;
	m.G0 = v26 + int32(16)
	goto L7
L11:
	;
	v60 = F_close(m, v32)
	mBase = m.M
	v65 = v58
	goto L10
L12:
	;
	v38 = l0 + int32(2)
	v39 = int32(8)
	goto L13
L13:
	;
	v44 = F_read(m, v32, v38, v39)
	mBase = m.M
	if v44 <= int32(0) {
		goto L15
	} else {
		goto L16
	}
L14:
	;
	v58 = int32(1)
	goto L11
L15:
	;
	v48 = *(*int32)(unsafe.Add(mBase, _consts[140]))
	if v48 == int32(27) {
		goto L13
	} else {
		goto L18
	}
L16:
	;
	goto L17
L17:
	;
	v53 = v39 - v44
	if v53 != 0 {
		v38 = v38 + v44
		v39 = v53
		goto L13
	} else {
		goto L19
	}
L18:
	;
	v58 = int32(0)
	goto L11
L19:
	;
	goto L14
L20:
	;
	v70 = v15
	goto L22
L21:
	;
	v70 = int32(-17)
	goto L22
L22:
	;
	v229 = v70
	goto L1
L23:
	;
	if v120 == int32(0) {
		v229 = v71
		goto L1
	} else {
		goto L36
	}
L24:
	;
	goto L28
L25:
	;
	v120 = v75
	goto L26
L26:
	;
	m.G0 = v81 + int32(16)
	goto L23
L27:
	;
	v115 = F_close(m, v87)
	mBase = m.M
	v120 = v113
	goto L26
L28:
	;
	v93 = l0 + int32(2)
	v94 = int32(8)
	goto L29
L29:
	;
	v99 = F_read(m, v87, v93, v94)
	mBase = m.M
	if v99 <= int32(0) {
		goto L31
	} else {
		goto L32
	}
L30:
	;
	v113 = int32(1)
	goto L27
L31:
	;
	v103 = *(*int32)(unsafe.Add(mBase, _consts[140]))
	if v103 == int32(27) {
		goto L29
	} else {
		goto L34
	}
L32:
	;
	goto L33
L33:
	;
	v108 = v94 - v99
	if v108 != 0 {
		v93 = v93 + v99
		v94 = v108
		goto L29
	} else {
		goto L35
	}
L34:
	;
	v113 = int32(0)
	goto L27
L35:
	;
	goto L30
L36:
	;
	v130 = int32(0)
	v134 = m.G0
	v136 = v134 - int32(16)
	m.G0 = v136
	*(*int32)(unsafe.Add(mBase, uint32(v136))) = v130
	v142 = F_open(m, int32(275211), v130, v136)
	mBase = m.M
	if v142 != int32(-1) {
		goto L38
	} else {
		goto L39
	}
L37:
	;
	if v175 == int32(0) {
		v229 = v71
		goto L1
	} else {
		goto L50
	}
L38:
	;
	goto L42
L39:
	;
	v175 = v130
	goto L40
L40:
	;
	m.G0 = v136 + int32(16)
	goto L37
L41:
	;
	v170 = F_close(m, v142)
	mBase = m.M
	v175 = v168
	goto L40
L42:
	;
	v148 = v9 + int32(15)
	v149 = int32(1)
	goto L43
L43:
	;
	v154 = F_read(m, v142, v148, v149)
	mBase = m.M
	if v154 <= int32(0) {
		goto L45
	} else {
		goto L46
	}
L44:
	;
	v168 = int32(1)
	goto L41
L45:
	;
	v158 = *(*int32)(unsafe.Add(mBase, _consts[140]))
	if v158 == int32(27) {
		goto L43
	} else {
		goto L48
	}
L46:
	;
	goto L47
L47:
	;
	v163 = v149 - v154
	if v163 != 0 {
		v148 = v148 + v154
		v149 = v163
		goto L43
	} else {
		goto L49
	}
L48:
	;
	v168 = int32(0)
	goto L41
L49:
	;
	goto L44
L50:
	;
	if l3 == int32(-1) {
		goto L3
	} else {
		goto L51
	}
L51:
	;
	v186 = int32(0)
	goto L52
L52:
	;
	v198 = int32(base.Ui32(v186)>>(uint(int32(4))%32)) + int32(6)
	if base.Ui32(l3) <= base.Ui32((v186&int32(14)|int32(16))<<(uint(v198)%32)) {
		goto L54
	} else {
		goto L55
	}
L53:
	;
	v224 = int32(255)
	goto L2
L54:
	;
	v224 = v186
	goto L2
L55:
	;
	goto L56
L56:
	;
	v202 = v186 | int32(1)
	if base.Ui32(l3) <= base.Ui32((v202&int32(15)|int32(16))<<(uint(v198)%32)) {
		v224 = v202
		goto L2
	} else {
		goto L57
	}
L57:
	;
	v210 = v186 + int32(2)
	if v210 != int32(256) {
		v186 = v210
		goto L52
	} else {
		goto L58
	}
L58:
	;
	goto L53
}
func F_pgp_set_cipher_algo(m *base.Module, l0 int32, l1 int32) int32 {
	mBase := m.M
	_ = mBase
	var v8 int32
	_ = v8
	var v16 int32
	_ = v16
	var v24 int32
	_ = v24
	var v32 int32
	_ = v32
	var v40 int32
	_ = v40
	var v48 int32
	_ = v48
	var v56 int32
	_ = v56
	var v64 int32
	_ = v64
	var v73 int32
	_ = v73
	var v77 int32
	_ = v77
	var v78 int32
	_ = v78
	var v79 int32
	_ = v79
	v8 = F_pg_strcasecmp(m, int32(161978), l1)
	mBase = m.M
	if v8 == int32(0) {
		v77 = int32(4336544)
		v78 = *(*int32)(unsafe.Add(mBase, uint32(v77)+4))
		v79 = v78
	} else {
		v16 = F_pg_strcasecmp(m, int32(527488), l1)
		mBase = m.M
		if v16 == int32(0) {
			v77 = int32(4336564)
			v78 = *(*int32)(unsafe.Add(mBase, uint32(v77)+4))
			v79 = v78
		} else {
			v24 = F_pg_strcasecmp(m, int32(324403), l1)
			mBase = m.M
			if v24 == int32(0) {
				v77 = int32(4336584)
				v78 = *(*int32)(unsafe.Add(mBase, uint32(v77)+4))
				v79 = v78
			} else {
				v32 = F_pg_strcasecmp(m, int32(308250), l1)
				mBase = m.M
				if v32 == int32(0) {
					v77 = int32(4336604)
					v78 = *(*int32)(unsafe.Add(mBase, uint32(v77)+4))
					v79 = v78
				} else {
					v40 = F_pg_strcasecmp(m, int32(162478), l1)
					mBase = m.M
					if v40 == int32(0) {
						v77 = int32(4336624)
						v78 = *(*int32)(unsafe.Add(mBase, uint32(v77)+4))
						v79 = v78
					} else {
						v48 = F_pg_strcasecmp(m, int32(526872), l1)
						mBase = m.M
						if v48 == int32(0) {
							v77 = int32(4336644)
							v78 = *(*int32)(unsafe.Add(mBase, uint32(v77)+4))
							v79 = v78
						} else {
							v56 = F_pg_strcasecmp(m, int32(530108), l1)
							mBase = m.M
							if v56 == int32(0) {
								v77 = int32(4336664)
								v78 = *(*int32)(unsafe.Add(mBase, uint32(v77)+4))
								v79 = v78
							} else {
								v64 = F_pg_strcasecmp(m, int32(527252), l1)
								mBase = m.M
								if v64 == int32(0) {
									v77 = int32(4336684)
									v78 = *(*int32)(unsafe.Add(mBase, uint32(v77)+4))
									v79 = v78
								} else {
									v73 = F_pg_strcasecmp(m, int32(308259), l1)
									mBase = m.M
									if v73 != 0 {
										v79 = int32(-103)
									} else {
										v77 = int32(4336704)
										v78 = *(*int32)(unsafe.Add(mBase, uint32(v77)+4))
										v79 = v78
									}
								}
							}
						}
					}
				}
			}
		}
	}
	if v79 < int32(0) {
		return v79
	} else {
		*(*int32)(unsafe.Add(mBase, uint32(l0)+60)) = v79
		return int32(0)
	}
}
func F_pgp_set_compress_level(m *base.Module, l0 int32, l1 int32) int32 {
	mBase := m.M
	_ = mBase
	var v9 int32
	_ = v9
	if base.Ui32(l1) <= base.Ui32(int32(9)) {
		*(*int32)(unsafe.Add(mBase, uint32(l0)+68)) = l1
		v9 = int32(0)
	} else {
		v9 = int32(-13)
	}
	return v9
}
func F_pgp_set_convert_crlf(m *base.Module, l0 int32, l1 int32) int32 {
	mBase := m.M
	_ = mBase
	var v3 int32
	_ = v3
	v3 = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(l0)+84)) = base.B2i32(l1 != v3)
	return v3
}
