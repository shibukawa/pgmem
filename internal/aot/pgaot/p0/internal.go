package p0

import (
	base "github.com/shibukawa/pgmem/internal/aot/pgaot/base"
	"unsafe"
)

func F_internal_load_library(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v13 int32
	_ = v13
	var v15 int32
	_ = v15
	var v20 int32
	_ = v20
	var v30 int32
	_ = v30
	var v32 int32
	_ = v32
	var v35 int32
	_ = v35
	var v36 int32
	_ = v36
	var v40 int32
	_ = v40
	var v41 int32
	_ = v41
	var v44 int32
	_ = v44
	var v45 int32
	_ = v45
	var v48 int32
	_ = v48
	var v55 int32
	_ = v55
	var v56 int32
	_ = v56
	var v62 int32
	_ = v62
	var v66 int32
	_ = v66
	var v67 int32
	_ = v67
	var v68 int64
	_ = v68
	var v71 int32
	_ = v71
	var v81 int64
	_ = v81
	var v83 int32
	_ = v83
	var v85 int32
	_ = v85
	var v98 int32
	_ = v98
	var v101 int32
	_ = v101
	var v109 int32
	_ = v109
	var v115 int32
	_ = v115
	var v117 int32
	_ = v117
	var v124 int32
	_ = v124
	var v130 int32
	_ = v130
	var v136 int32
	_ = v136
	var v140 int32
	_ = v140
	var v142 int32
	_ = v142
	var v143 int32
	_ = v143
	var v147 int32
	_ = v147
	var v148 int32
	_ = v148
	var v150 int32
	_ = v150
	var v154 int32
	_ = v154
	var v156 int32
	_ = v156
	var v158 int32
	_ = v158
	var v161 int32
	_ = v161
	var v166 int32
	_ = v166
	var v167 int32
	_ = v167
	var v168 int32
	_ = v168
	var v170 int32
	_ = v170
	var v171 int32
	_ = v171
	var v172 int32
	_ = v172
	var v174 int32
	_ = v174
	var v178 int32
	_ = v178
	var v183 int32
	_ = v183
	var v184 int32
	_ = v184
	var v185 int32
	_ = v185
	var v192 int32
	_ = v192
	var v194 int32
	_ = v194
	var v195 int32
	_ = v195
	var v197 int32
	_ = v197
	var v205 int32
	_ = v205
	var v207 int64
	_ = v207
	var v208 int32
	_ = v208
	var v211 int32
	_ = v211
	var v213 int32
	_ = v213
	var v219 int64
	_ = v219
	var v223 int64
	_ = v223
	var v227 int64
	_ = v227
	var v231 int64
	_ = v231
	var v236 int32
	_ = v236
	var v243 int32
	_ = v243
	var v245 int32
	_ = v245
	var v249 int32
	_ = v249
	var v250 int32
	_ = v250
	var v251 int32
	_ = v251
	var v255 int32
	_ = v255
	var v258 int32
	_ = v258
	var v262 int32
	_ = v262
	var v269 int32
	_ = v269
	var v271 int32
	_ = v271
	var v275 int32
	_ = v275
	var v276 int32
	_ = v276
	var v277 int32
	_ = v277
	var v281 int32
	_ = v281
	var v283 int32
	_ = v283
	var v284 int32
	_ = v284
	var v285 int32
	_ = v285
	var v287 int32
	_ = v287
	var v294 int32
	_ = v294
	var v303 int32
	_ = v303
	var v306 int32
	_ = v306
	var v307 int32
	_ = v307
	var v314 int32
	_ = v314
	var v315 int32
	_ = v315
	var v316 int32
	_ = v316
	var v317 int32
	_ = v317
	var v318 int32
	_ = v318
	var v320 int32
	_ = v320
	var v325 int32
	_ = v325
	var v328 int32
	_ = v328
	var v329 int32
	_ = v329
	var v330 int32
	_ = v330
	var v335 int32
	_ = v335
	var v337 int32
	_ = v337
	var v340 int32
	_ = v340
	var v344 int32
	_ = v344
	var v345 int32
	_ = v345
	var v352 int32
	_ = v352
	var v356 int32
	_ = v356
	var v374 int32
	_ = v374
	var v377 int32
	_ = v377
	var v392 int32
	_ = v392
	var v402 int32
	_ = v402
	var v413 int32
	_ = v413
	var v414 int32
	_ = v414
	var v417 int32
	_ = v417
	var v418 int32
	_ = v418
	var v419 int32
	_ = v419
	var v423 int32
	_ = v423
	var v424 int32
	_ = v424
	var v425 int32
	_ = v425
	var v433 int32
	_ = v433
	var v434 int32
	_ = v434
	var v435 int32
	_ = v435
	var v438 int32
	_ = v438
	var v439 int32
	_ = v439
	var v441 int32
	_ = v441
	var v442 int32
	_ = v442
	var v444 int32
	_ = v444
	var v446 int32
	_ = v446
	var v456 int32
	_ = v456
	var v457 int32
	_ = v457
	var v458 int32
	_ = v458
	var v461 int32
	_ = v461
	var v462 int32
	_ = v462
	var v463 int32
	_ = v463
	var v466 int32
	_ = v466
	var v467 int32
	_ = v467
	var v469 int32
	_ = v469
	var v474 int32
	_ = v474
	var v487 int32
	_ = v487
	var v489 int32
	_ = v489
	var v491 int32
	_ = v491
	var v492 int32
	_ = v492
	var v494 int32
	_ = v494
	var v496 int32
	_ = v496
	var v502 int32
	_ = v502
	var v508 int32
	_ = v508
	var v518 int32
	_ = v518
	var v526 int32
	_ = v526
	var v528 int32
	_ = v528
	var v532 int32
	_ = v532
	var v537 int32
	_ = v537
	var v541 int32
	_ = v541
	var v544 int32
	_ = v544
	var v548 int32
	_ = v548
	var v553 int32
	_ = v553
	var v555 int32
	_ = v555
	var v557 int32
	_ = v557
	var v561 int32
	_ = v561
	var v566 int32
	_ = v566
	var v568 int32
	_ = v568
	var v575 int32
	_ = v575
	var v580 int32
	_ = v580
	var v581 int64
	_ = v581
	var v583 int64
	_ = v583
	var v585 int64
	_ = v585
	var v587 int64
	_ = v587
	var v589 int64
	_ = v589
	var v593 int64
	_ = v593
	var v595 int64
	_ = v595
	var v597 int64
	_ = v597
	var v601 int32
	_ = v601
	var v603 int32
	_ = v603
	var v608 int32
	_ = v608
	var v609 int32
	_ = v609
	var v615 int32
	_ = v615
	var v623 int32
	_ = v623
	var v624 int32
	_ = v624
	var v625 int32
	_ = v625
	var v626 int32
	_ = v626
	var v638 int32
	_ = v638
	var v639 int32
	_ = v639
	var v644 int32
	_ = v644
	var v650 int32
	_ = v650
	var v660 int32
	_ = v660
	var v665 int32
	_ = v665
	var v667 int32
	_ = v667
	var v668 int32
	_ = v668
	var v671 int32
	_ = v671
	var v672 int32
	_ = v672
	var v676 int32
	_ = v676
	var v677 int32
	_ = v677
	var v680 int32
	_ = v680
	var v681 int32
	_ = v681
	var v684 int32
	_ = v684
	var v691 int32
	_ = v691
	var v692 int32
	_ = v692
	var v699 int32
	_ = v699
	var v700 int32
	_ = v700
	var v703 int32
	_ = v703
	var v708 int32
	_ = v708
	var v709 int32
	_ = v709
	var v710 int32
	_ = v710
	var v722 int32
	_ = v722
	var v723 int32
	_ = v723
	var v726 int32
	_ = v726
	var v731 int32
	_ = v731
	var v732 int32
	_ = v732
	var v733 int32
	_ = v733
	var v745 int32
	_ = v745
	var v746 int32
	_ = v746
	var v749 int32
	_ = v749
	var v754 int32
	_ = v754
	var v755 int32
	_ = v755
	var v756 int32
	_ = v756
	var v768 int32
	_ = v768
	var v769 int32
	_ = v769
	var v770 int32
	_ = v770
	var v775 int32
	_ = v775
	var v778 int32
	_ = v778
	var v779 int32
	_ = v779
	var v781 int32
	_ = v781
	var v793 int32
	_ = v793
	var v794 int32
	_ = v794
	var v801 int32
	_ = v801
	var v805 int32
	_ = v805
	var v811 int32
	_ = v811
	var v812 int32
	_ = v812
	var v816 int32
	_ = v816
	var v821 int32
	_ = v821
	var v825 int32
	_ = v825
	var v831 int32
	_ = v831
	var v839 int32
	_ = v839
	var v844 int32
	_ = v844
	var v850 int32
	_ = v850
	var v856 int32
	_ = v856
	var v860 int32
	_ = v860
	var v865 int32
	_ = v865
	v13 = m.G0
	v15 = v13 - int32(208)
	m.G0 = v15
	v20 = int32(4514032)
	goto L7
L1:
	;
	F_emscripten_builtin_free(m, v101)
	mBase = m.M
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v850 = m.ExcPending
	if v850 != 0 {
		goto L112
	} else {
		goto L234
	}
L2:
	;
	v581 = *(*int64)(unsafe.Add(mBase, uint32(v417)+56))
	*(*int64)(unsafe.Add(mBase, uint32(v15)+104)) = v581
	v583 = *(*int64)(unsafe.Add(mBase, uint32(v417)+48))
	*(*int64)(unsafe.Add(mBase, uint32(v15)+96)) = v583
	v585 = *(*int64)(unsafe.Add(mBase, uint32(v417)+40))
	*(*int64)(unsafe.Add(mBase, uint32(v15)+88)) = v585
	v587 = *(*int64)(unsafe.Add(mBase, uint32(v417)+32))
	*(*int64)(unsafe.Add(mBase, uint32(v15)+80)) = v587
	v589 = *(*int64)(unsafe.Add(mBase, uint32(v417)+24))
	*(*int64)(unsafe.Add(mBase, uint32(v15)+72)) = v589
	v593 = *(*int64)(unsafe.Add(mBase, uint32(v417)+16))
	*(*int64)(unsafe.Add(mBase, uint32(v15-int32(-64)))) = v593
	v595 = *(*int64)(unsafe.Add(mBase, uint32(v417)+8))
	*(*int64)(unsafe.Add(mBase, uint32(v15)+56)) = v595
	v597 = *(*int64)(unsafe.Add(mBase, uint32(v417)))
	*(*int64)(unsafe.Add(mBase, uint32(v15)+48)) = v597
	F_emscripten_builtin_free(m, v101)
	mBase = m.M
	v601 = m.G0
	v603 = v601 - int32(224)
	m.G0 = v603
	v608 = v15 + int32(48) | int32(4)
	v609 = *(*int32)(unsafe.Add(mBase, uint32(v608)))
	if v609 != int32(1800) {
		goto L162
	} else {
		goto L163
	}
L3:
	;
	v555 = int32(*(*uint8)(unsafe.Add(mBase, _consts[925])))
	if v555 != 0 {
		goto L155
	} else {
		goto L156
	}
L4:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v541 = m.ExcPending
	if v541 != 0 {
		goto L112
	} else {
		goto L151
	}
L5:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v526 = m.ExcPending
	if v526 != 0 {
		goto L112
	} else {
		goto L147
	}
L6:
	;
	v518 = *(*int32)(unsafe.Add(mBase, uint32(v508)+16))
	m.G0 = v15 + int32(208)
	return v518
L7:
	;
	v30 = *(*int32)(unsafe.Add(mBase, uint32(v20)))
	if v30 != 0 {
		goto L9
	} else {
		goto L10
	}
L8:
	;
	v62 = F___fstatat(m, int32(-100), l0, v15+int32(112), int32(0))
	mBase = m.M
	goto L21
L9:
	;
	v32 = v30 + int32(24)
	v35 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v32))))
	v36 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0))))
	if v36 == int32(0) {
		v55 = v35
		v56 = v36
		goto L13
	} else {
		goto L14
	}
L10:
	;
	goto L11
L11:
	;
	goto L8
L12:
	;
	if v56-v55 != 0 {
		v20 = v30
		goto L7
	} else {
		goto L20
	}
L13:
	;
	goto L12
L14:
	;
	if v35 != v36 {
		v55 = v35
		v56 = v36
		goto L13
	} else {
		goto L15
	}
L15:
	;
	v40 = l0
	v41 = v32
	goto L16
L16:
	;
	v44 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v41)+1)))
	v45 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v40)+1)))
	if v45 == int32(0) {
		v55 = v44
		v56 = v45
		goto L13
	} else {
		goto L18
	}
L17:
	;
	v55 = v44
	v56 = v45
	goto L13
L18:
	;
	v48 = int32(1)
	if v44 == v45 {
		v40 = v40 + v48
		v41 = v41 + v48
		goto L16
	} else {
		goto L19
	}
L19:
	;
	goto L17
L20:
	;
	v508 = v30
	goto L6
L21:
	;
	if v62 == int32(-1) {
		goto L5
	} else {
		goto L22
	}
L22:
	;
	v66 = *(*int32)(unsafe.Add(mBase, _consts[926]))
	if v66 != 0 {
		goto L23
	} else {
		goto L24
	}
L23:
	;
	v67 = *(*int32)(unsafe.Add(mBase, uint32(v15)+112))
	v68 = *(*int64)(unsafe.Add(mBase, uint32(v15)+200))
	v71 = v66
	goto L26
L24:
	;
	goto L25
L25:
	;
	v98 = F_strlen(m, l0)
	mBase = m.M
	v101 = F_emscripten_builtin_malloc(m, v98+int32(25))
	mBase = m.M
	if v101 == int32(0) {
		goto L4
	} else {
		goto L33
	}
L26:
	;
	v81 = *(*int64)(unsafe.Add(mBase, uint32(v71)+8))
	if v81 == v68 {
		goto L28
	} else {
		goto L29
	}
L27:
	;
	goto L25
L28:
	;
	v83 = *(*int32)(unsafe.Add(mBase, uint32(v71)+4))
	if v67 == v83 {
		v508 = v71
		goto L6
	} else {
		goto L31
	}
L29:
	;
	goto L30
L30:
	;
	v85 = *(*int32)(unsafe.Add(mBase, uint32(v71)))
	if v85 != 0 {
		v71 = v85
		goto L26
	} else {
		goto L32
	}
L31:
	;
	goto L30
L32:
	;
	goto L27
L33:
	;
	if v101&int32(3) == int32(0) {
		goto L35
	} else {
		goto L36
	}
L34:
	;
	v130 = v101 + int32(24)
	if (l0^v130)&int32(3) != 0 {
		goto L46
	} else {
		goto L47
	}
L35:
	;
	v109 = v101 + int32(24)
	if base.Ui32(v109) <= base.Ui32(v101) {
		goto L34
	} else {
		goto L38
	}
L36:
	;
	goto L37
L37:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v101)+16)) = int64(0)
	goto L34
L38:
	;
	v115 = v101 + int32(4)
	if base.Ui32(v115) < base.Ui32(v109) {
		goto L39
	} else {
		goto L40
	}
L39:
	;
	v117 = v109
	goto L41
L40:
	;
	v117 = v115
	goto L41
L41:
	;
	v124 = F__emscripten_memset_bulkmem(m, v101, base.I32_extend8_s(int32(0)), (v101^int32(-1)+v117)&int32(-4)+int32(4))
	mBase = m.M
	goto L42
L42:
	;
	goto L34
L43:
	;
	v205 = *(*int32)(unsafe.Add(mBase, uint32(v15)+112))
	*(*int32)(unsafe.Add(mBase, uint32(v101)+4)) = v205
	v207 = *(*int64)(unsafe.Add(mBase, uint32(v15)+200))
	v208 = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v101))) = v208
	*(*int64)(unsafe.Add(mBase, uint32(v101)+8)) = v207
	v211 = m.G0
	v213 = v211 - int32(16)
	m.G0 = v213
	if v130 == v208 {
		goto L66
	} else {
		goto L67
	}
L44:
	;
	goto L43
L45:
	;
	*(*uint8)(unsafe.Add(mBase, uint32(v185))) = uint8(v184)
	if v184&int32(255) == int32(0) {
		goto L44
	} else {
		goto L60
	}
L46:
	;
	v136 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0))))
	v183 = l0
	v184 = v136
	v185 = v130
	goto L45
L47:
	;
	goto L48
L48:
	;
	if l0&int32(3) != 0 {
		goto L49
	} else {
		goto L50
	}
L49:
	;
	v140 = l0
	v142 = v130
	goto L52
L50:
	;
	v154 = l0
	v156 = v130
	goto L51
L51:
	;
	v158 = *(*int32)(unsafe.Add(mBase, uint32(v154)))
	v161 = int32(-2139062144)
	if (int32(16843008)-v158|v158)&v161 != v161 {
		v183 = v154
		v184 = v158
		v185 = v156
		goto L45
	} else {
		goto L56
	}
L52:
	;
	v143 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v140))))
	*(*uint8)(unsafe.Add(mBase, uint32(v142))) = uint8(v143)
	if v143 == int32(0) {
		goto L44
	} else {
		goto L54
	}
L53:
	;
	v154 = v150
	v156 = v148
	goto L51
L54:
	;
	v147 = int32(1)
	v148 = v142 + v147
	v150 = v140 + v147
	if v150&int32(3) != 0 {
		v140 = v150
		v142 = v148
		goto L52
	} else {
		goto L55
	}
L55:
	;
	goto L53
L56:
	;
	v166 = v154
	v167 = v158
	v168 = v156
	goto L57
L57:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v168))) = v167
	v170 = int32(4)
	v171 = v168 + v170
	v172 = *(*int32)(unsafe.Add(mBase, uint32(v166)+4))
	v174 = v166 + v170
	v178 = int32(-2139062144)
	if (v172|(int32(16843008)-v172))&v178 == v178 {
		v166 = v174
		v167 = v172
		v168 = v171
		goto L57
	} else {
		goto L59
	}
L58:
	;
	v183 = v174
	v184 = v172
	v185 = v171
	goto L45
L59:
	;
	goto L58
L60:
	;
	v192 = v183
	v194 = v185
	goto L61
L61:
	;
	v195 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v192)+1)))
	*(*uint8)(unsafe.Add(mBase, uint32(v194)+1)) = uint8(v195)
	v197 = int32(1)
	if v195 != 0 {
		v192 = v192 + v197
		v194 = v194 + v197
		goto L61
	} else {
		goto L63
	}
L62:
	;
	goto L44
L63:
	;
	goto L62
L64:
	;
	m.G0 = v213 + int32(16)
	*(*int32)(unsafe.Add(mBase, uint32(v101)+16)) = v402
	if v402 == int32(0) {
		goto L3
	} else {
		goto L114
	}
L65:
	;
	v392 = int32(1)
	*(*uint8)(unsafe.Add(mBase, _consts[925])) = uint8(v392)
	v402 = int32(0)
	goto L64
L66:
	;
	v219 = *(*int64)(unsafe.Add(mBase, _consts[927]))
	*(*int64)(unsafe.Add(mBase, _consts[928])) = v219
	v223 = *(*int64)(unsafe.Add(mBase, _consts[929]))
	*(*int64)(unsafe.Add(mBase, _consts[930])) = v223
	v227 = *(*int64)(unsafe.Add(mBase, _consts[931]))
	*(*int64)(unsafe.Add(mBase, _consts[932])) = v227
	v231 = *(*int64)(unsafe.Add(mBase, _consts[933]))
	*(*int64)(unsafe.Add(mBase, _consts[934])) = v231
	goto L65
L67:
	;
	goto L68
L68:
	;
	v236 = F_strlen(m, v130)
	mBase = m.M
	v243 = v236 + int32(1)
	goto L72
L69:
	;
	v285 = int32(0)
	v287 = *(*int32)(unsafe.Add(mBase, _consts[935]))
	if v285 < v287 {
		goto L88
	} else {
		goto L89
	}
L70:
	;
	if v255 != 0 {
		goto L76
	} else {
		goto L77
	}
L71:
	;
	goto L70
L72:
	;
	v245 = int32(0)
	if v243 == v245 {
		v255 = v245
		goto L71
	} else {
		goto L74
	}
L73:
	;
	v255 = v250
	goto L71
L74:
	;
	v249 = v243 - int32(1)
	v250 = v130 + v249
	v251 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v250))))
	if v251 != int32(47) {
		v243 = v249
		goto L72
	} else {
		goto L75
	}
L75:
	;
	goto L73
L76:
	;
	v258 = v255 + int32(1)
	goto L78
L77:
	;
	v258 = v130
	goto L78
L78:
	;
	v262 = F_strlen(m, v258)
	mBase = m.M
	v269 = v262 + int32(1)
	goto L81
L79:
	;
	if v281 != 0 {
		goto L85
	} else {
		goto L86
	}
L80:
	;
	goto L79
L81:
	;
	v271 = int32(0)
	if v269 == v271 {
		v281 = v271
		goto L80
	} else {
		goto L83
	}
L82:
	;
	v281 = v276
	goto L80
L83:
	;
	v275 = v269 - int32(1)
	v276 = v258 + v275
	v277 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v276))))
	if v277 != int32(46) {
		v269 = v275
		goto L81
	} else {
		goto L84
	}
L84:
	;
	goto L82
L85:
	;
	v284 = v281 - v258
	goto L69
L86:
	;
	goto L87
L87:
	;
	v283 = F_strlen(m, v258)
	mBase = m.M
	v284 = v283
	goto L69
L88:
	;
	v294 = v285
	goto L91
L89:
	;
	goto L90
L90:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v213))) = v130
	v374 = F_snprintf(m, int32(4612944), int32(512), int32(432179), v213)
	mBase = m.M
	v377 = m.ExcPending
	if v377 != 0 {
		goto L112
	} else {
		goto L113
	}
L91:
	;
	v303 = v294 * int32(12)
	v306 = *(*int32)(unsafe.Add(mBase, uint32(v303)+uint32(_consts[936])))
	v307 = F_strlen(m, v306)
	mBase = m.M
	if v307 == v284 {
		goto L93
	} else {
		goto L94
	}
L92:
	;
	goto L90
L93:
	;
	if v284 == int32(0) {
		goto L97
	} else {
		goto L98
	}
L94:
	;
	goto L95
L95:
	;
	v356 = v294 + int32(1)
	if v356 != v287 {
		v294 = v356
		goto L91
	} else {
		goto L111
	}
L96:
	;
	if v352 == int32(0) {
		v402 = v303 + int32(2167584)
		goto L64
	} else {
		goto L110
	}
L97:
	;
	v352 = int32(0)
	goto L96
L98:
	;
	goto L99
L99:
	;
	v314 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v306))))
	if v314 != 0 {
		goto L100
	} else {
		goto L101
	}
L100:
	;
	v315 = v306
	v316 = v258
	v317 = v284
	v318 = v314
	goto L104
L101:
	;
	v340 = v258
	v344 = int32(0)
	goto L102
L102:
	;
	v345 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v340))))
	v352 = v344 - v345
	goto L96
L103:
	;
	v340 = v335
	v344 = v337
	goto L102
L104:
	;
	v320 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v316))))
	if v318 != v320 {
		v335 = v316
		v337 = v318
		goto L103
	} else {
		goto L106
	}
L105:
	;
	v335 = v329
	v337 = int32(0)
	goto L103
L106:
	;
	if v320 == int32(0) {
		v335 = v316
		v337 = v318
		goto L103
	} else {
		goto L107
	}
L107:
	;
	v325 = v317 - int32(1)
	if v325 == int32(0) {
		v335 = v316
		v337 = v318
		goto L103
	} else {
		goto L108
	}
L108:
	;
	v328 = int32(1)
	v329 = v316 + v328
	v330 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v315)+1)))
	if v330 != 0 {
		v315 = v315 + v328
		v316 = v329
		v317 = v325
		v318 = v330
		goto L104
	} else {
		goto L109
	}
L109:
	;
	goto L105
L110:
	;
	goto L95
L111:
	;
	goto L92
L112:
	;
	return int32(0)
L113:
	;
	goto L65
L114:
	;
	v413 = F_pgmem_dlsym(m, v402, int32(491106))
	mBase = m.M
	v414 = m.ExcPending
	if v414 != 0 {
		goto L112
	} else {
		goto L115
	}
L115:
	;
	if v413 == int32(0) {
		goto L1
	} else {
		goto L116
	}
L116:
	;
	v417 = m.T0[v413].(func(*base.Module) int32)(m)
	mBase = m.M
	v418 = m.ExcPending
	if v418 != 0 {
		goto L112
	} else {
		goto L117
	}
L117:
	;
	v419 = *(*int32)(unsafe.Add(mBase, uint32(v417)))
	if v419 != int32(64) {
		goto L2
	} else {
		goto L118
	}
L118:
	;
	v423 = v417 + int32(4)
	v424 = int32(1766376)
	v425 = int32(52)
	goto L122
L119:
	;
	if v487 != 0 {
		goto L2
	} else {
		goto L137
	}
L120:
	;
	v487 = int32(0)
	goto L119
L121:
	;
	v461 = v456
	v462 = v457
	v463 = v458
	goto L131
L122:
	;
	if (v423|v424)&int32(3) != 0 {
		v456 = v423
		v457 = v424
		v458 = v425
		goto L121
	} else {
		goto L125
	}
L124:
	;
	if v446 == int32(0) {
		goto L120
	} else {
		goto L130
	}
L125:
	;
	v433 = v423
	v434 = v424
	v435 = v425
	goto L126
L126:
	;
	v438 = *(*int32)(unsafe.Add(mBase, uint32(v433)))
	v439 = *(*int32)(unsafe.Add(mBase, uint32(v434)))
	if v438 != v439 {
		v456 = v433
		v457 = v434
		v458 = v435
		goto L121
	} else {
		goto L128
	}
L127:
	;
	goto L124
L128:
	;
	v441 = int32(4)
	v442 = v434 + v441
	v444 = v433 + v441
	v446 = v435 - v441
	if base.Ui32(int32(3)) < base.Ui32(v446) {
		v433 = v444
		v434 = v442
		v435 = v446
		goto L126
	} else {
		goto L129
	}
L129:
	;
	goto L127
L130:
	;
	v456 = v444
	v457 = v442
	v458 = v446
	goto L121
L131:
	;
	v466 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v461))))
	v467 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v462))))
	if v466 == v467 {
		goto L133
	} else {
		goto L134
	}
L132:
	;
	v487 = v466 - v467
	goto L119
L133:
	;
	v469 = int32(1)
	v474 = v463 - v469
	if v474 != 0 {
		v461 = v461 + v469
		v462 = v462 + v469
		v463 = v474
		goto L131
	} else {
		goto L136
	}
L134:
	;
	goto L135
L135:
	;
	goto L132
L136:
	;
	goto L120
L137:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v101)+20)) = v417
	v489 = *(*int32)(unsafe.Add(mBase, uint32(v101)+16))
	v491 = F_pgmem_dlsym(m, v489, int32(100670))
	mBase = m.M
	v492 = m.ExcPending
	if v492 != 0 {
		goto L112
	} else {
		goto L138
	}
L138:
	;
	if v491 != 0 {
		goto L139
	} else {
		goto L140
	}
L139:
	;
	m.T0[v491].(func(*base.Module))(m)
	mBase = m.M
	v494 = m.ExcPending
	if v494 != 0 {
		goto L112
	} else {
		goto L142
	}
L140:
	;
	goto L141
L141:
	;
	v496 = *(*int32)(unsafe.Add(mBase, _consts[926]))
	if v496 == int32(0) {
		goto L144
	} else {
		goto L145
	}
L142:
	;
	goto L141
L143:
	;
	*(*int32)(unsafe.Add(mBase, _consts[937])) = v101
	v508 = v101
	goto L6
L144:
	;
	*(*int32)(unsafe.Add(mBase, _consts[926])) = v101
	goto L143
L145:
	;
	goto L146
L146:
	;
	v502 = *(*int32)(unsafe.Add(mBase, _consts[937]))
	*(*int32)(unsafe.Add(mBase, uint32(v502))) = v101
	goto L143
L147:
	;
	F_errcode_for_file_access(m)
	mBase = m.M
	v528 = m.ExcPending
	if v528 != 0 {
		goto L112
	} else {
		goto L148
	}
L148:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v15))) = l0
	F_errmsg(m, int32(298600), v15)
	mBase = m.M
	v532 = m.ExcPending
	if v532 != 0 {
		goto L112
	} else {
		goto L149
	}
L149:
	;
	F_errfinish(m, int32(496310), int32(215), int32(17606))
	mBase = m.M
	v537 = m.ExcPending
	if v537 != 0 {
		goto L112
	} else {
		goto L150
	}
L150:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L151:
	;
	F_errcode(m, int32(8389))
	mBase = m.M
	v544 = m.ExcPending
	if v544 != 0 {
		goto L112
	} else {
		goto L152
	}
L152:
	;
	F_errmsg(m, int32(13904), int32(0))
	mBase = m.M
	v548 = m.ExcPending
	if v548 != 0 {
		goto L112
	} else {
		goto L153
	}
L153:
	;
	F_errfinish(m, int32(496310), int32(234), int32(17606))
	mBase = m.M
	v553 = m.ExcPending
	if v553 != 0 {
		goto L112
	} else {
		goto L154
	}
L154:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L155:
	;
	v557 = int32(0)
	*(*uint8)(unsafe.Add(mBase, _consts[925])) = uint8(v557)
	v561 = int32(4612944)
	goto L157
L156:
	;
	v561 = int32(0)
	goto L157
L157:
	;
	F_emscripten_builtin_free(m, v101)
	mBase = m.M
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v566 = m.ExcPending
	if v566 != 0 {
		goto L112
	} else {
		goto L158
	}
L158:
	;
	F_errcode_for_file_access(m)
	mBase = m.M
	v568 = m.ExcPending
	if v568 != 0 {
		goto L112
	} else {
		goto L159
	}
L159:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v15)+20)) = v561
	*(*int32)(unsafe.Add(mBase, uint32(v15)+16)) = l0
	F_errmsg(m, int32(205709), v15+int32(16))
	mBase = m.M
	v575 = m.ExcPending
	if v575 != 0 {
		goto L112
	} else {
		goto L160
	}
L160:
	;
	F_errfinish(m, int32(496310), int32(253), int32(17606))
	mBase = m.M
	v580 = m.ExcPending
	if v580 != 0 {
		goto L112
	} else {
		goto L161
	}
L161:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L162:
	;
	if int32(1000) <= v609 {
		goto L166
	} else {
		goto L167
	}
L163:
	;
	goto L164
L164:
	;
	v667 = v608 + int32(20)
	v668 = int32(1766396)
	v671 = int32(*(*uint8)(unsafe.Add(mBase, _consts[938])))
	v672 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v667))))
	if v672 == int32(0) {
		v691 = v671
		v692 = v672
		goto L176
	} else {
		goto L177
	}
L165:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v644 = m.ExcPending
	if v644 != 0 {
		goto L112
	} else {
		goto L171
	}
L166:
	;
	v615 = base.I32_div_u_s(v609, int32(100))
	*(*int32)(unsafe.Add(mBase, uint32(v603)+160)) = v615
	v623 = F_pg_snprintf(m, v603+int32(192), int32(32), int32(489370), v603+int32(160))
	mBase = m.M
	v624 = m.ExcPending
	if v624 != 0 {
		goto L112
	} else {
		goto L169
	}
L167:
	;
	goto L168
L168:
	;
	v625 = int32(100)
	v626 = base.I32_div_s(v609, v625)
	*(*int32)(unsafe.Add(mBase, uint32(v603)+176)) = v626
	*(*int32)(unsafe.Add(mBase, uint32(v603)+180)) = v609 - v626*v625
	v638 = F_pg_snprintf(m, v603+int32(192), int32(32), int32(467821), v603+int32(176))
	mBase = m.M
	v639 = m.ExcPending
	if v639 != 0 {
		goto L112
	} else {
		goto L170
	}
L169:
	;
	goto L165
L170:
	;
	goto L165
L171:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v603)+144)) = l0
	F_errmsg(m, int32(325532), v603+int32(144))
	mBase = m.M
	v650 = m.ExcPending
	if v650 != 0 {
		goto L112
	} else {
		goto L172
	}
L172:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v603)+128)) = int32(18)
	*(*int32)(unsafe.Add(mBase, uint32(v603)+132)) = v603 + int32(192)
	F_errdetail(m, int32(606052), v603+int32(128))
	mBase = m.M
	v660 = m.ExcPending
	if v660 != 0 {
		goto L112
	} else {
		goto L173
	}
L173:
	;
	F_errfinish(m, int32(496310), int32(340), int32(212399))
	mBase = m.M
	v665 = m.ExcPending
	if v665 != 0 {
		goto L112
	} else {
		goto L174
	}
L174:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L175:
	;
	if v692-v691 == int32(0) {
		goto L183
	} else {
		goto L184
	}
L176:
	;
	goto L175
L177:
	;
	if v671 != v672 {
		v691 = v671
		v692 = v672
		goto L176
	} else {
		goto L178
	}
L178:
	;
	v676 = v667
	v677 = v668
	goto L179
L179:
	;
	v680 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v677)+1)))
	v681 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v676)+1)))
	if v681 == int32(0) {
		v691 = v680
		v692 = v681
		goto L176
	} else {
		goto L181
	}
L180:
	;
	v691 = v680
	v692 = v681
	goto L176
L181:
	;
	v684 = int32(1)
	if v680 == v681 {
		v676 = v676 + v684
		v677 = v677 + v684
		goto L179
	} else {
		goto L182
	}
L182:
	;
	goto L180
L183:
	;
	F_initStringInfo(m, v603+int32(192))
	mBase = m.M
	v699 = m.ExcPending
	if v699 != 0 {
		goto L112
	} else {
		goto L186
	}
L184:
	;
	goto L185
L185:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v825 = m.ExcPending
	if v825 != 0 {
		goto L112
	} else {
		goto L230
	}
L186:
	;
	v700 = *(*int32)(unsafe.Add(mBase, uint32(v608)+4))
	if v700 != int32(100) {
		goto L187
	} else {
		goto L188
	}
L187:
	;
	v703 = *(*int32)(unsafe.Add(mBase, uint32(v603)+196))
	if v703 != 0 {
		goto L190
	} else {
		goto L191
	}
L188:
	;
	goto L189
L189:
	;
	v723 = *(*int32)(unsafe.Add(mBase, uint32(v608)+8))
	if v723 != int32(32) {
		goto L195
	} else {
		goto L196
	}
L190:
	;
	F_appendStringInfoChar(m, v603+int32(192), int32(10))
	mBase = m.M
	v708 = m.ExcPending
	if v708 != 0 {
		goto L112
	} else {
		goto L193
	}
L191:
	;
	v710 = v700
	goto L192
L192:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v603)+88)) = v710
	*(*int32)(unsafe.Add(mBase, uint32(v603)+84)) = int32(100)
	*(*int32)(unsafe.Add(mBase, uint32(v603)+80)) = int32(524413)
	F_appendStringInfo(m, v603+int32(192), int32(654697), v603+int32(80))
	mBase = m.M
	v722 = m.ExcPending
	if v722 != 0 {
		goto L112
	} else {
		goto L194
	}
L193:
	;
	v709 = *(*int32)(unsafe.Add(mBase, uint32(v608)+4))
	v710 = v709
	goto L192
L194:
	;
	goto L189
L195:
	;
	v726 = *(*int32)(unsafe.Add(mBase, uint32(v603)+196))
	if v726 != 0 {
		goto L198
	} else {
		goto L199
	}
L196:
	;
	goto L197
L197:
	;
	v746 = *(*int32)(unsafe.Add(mBase, uint32(v608)+12))
	if v746 != int32(64) {
		goto L203
	} else {
		goto L204
	}
L198:
	;
	F_appendStringInfoChar(m, v603+int32(192), int32(10))
	mBase = m.M
	v731 = m.ExcPending
	if v731 != 0 {
		goto L112
	} else {
		goto L201
	}
L199:
	;
	v733 = v723
	goto L200
L200:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v603)+72)) = v733
	*(*int32)(unsafe.Add(mBase, uint32(v603)+68)) = int32(32)
	*(*int32)(unsafe.Add(mBase, uint32(v603)+64)) = int32(523477)
	F_appendStringInfo(m, v603+int32(192), int32(654697), v603-int32(-64))
	mBase = m.M
	v745 = m.ExcPending
	if v745 != 0 {
		goto L112
	} else {
		goto L202
	}
L201:
	;
	v732 = *(*int32)(unsafe.Add(mBase, uint32(v608)+8))
	v733 = v732
	goto L200
L202:
	;
	goto L197
L203:
	;
	v749 = *(*int32)(unsafe.Add(mBase, uint32(v603)+196))
	if v749 != 0 {
		goto L206
	} else {
		goto L207
	}
L204:
	;
	goto L205
L205:
	;
	v769 = *(*int32)(unsafe.Add(mBase, uint32(v608)+16))
	if v769 != 0 {
		goto L211
	} else {
		goto L212
	}
L206:
	;
	F_appendStringInfoChar(m, v603+int32(192), int32(10))
	mBase = m.M
	v754 = m.ExcPending
	if v754 != 0 {
		goto L112
	} else {
		goto L209
	}
L207:
	;
	v756 = v746
	goto L208
L208:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v603)+56)) = v756
	*(*int32)(unsafe.Add(mBase, uint32(v603)+52)) = int32(64)
	*(*int32)(unsafe.Add(mBase, uint32(v603)+48)) = int32(531941)
	F_appendStringInfo(m, v603+int32(192), int32(654697), v603+int32(48))
	mBase = m.M
	v768 = m.ExcPending
	if v768 != 0 {
		goto L112
	} else {
		goto L210
	}
L209:
	;
	v755 = *(*int32)(unsafe.Add(mBase, uint32(v608)+12))
	v756 = v755
	goto L208
L210:
	;
	goto L205
L211:
	;
	v770 = *(*int32)(unsafe.Add(mBase, uint32(v603)+196))
	if v770 != 0 {
		goto L214
	} else {
		goto L215
	}
L212:
	;
	goto L213
L213:
	;
	v794 = *(*int32)(unsafe.Add(mBase, uint32(v603)+196))
	if v794 == int32(0) {
		goto L222
	} else {
		goto L223
	}
L214:
	;
	F_appendStringInfoChar(m, v603+int32(192), int32(10))
	mBase = m.M
	v775 = m.ExcPending
	if v775 != 0 {
		goto L112
	} else {
		goto L217
	}
L215:
	;
	v781 = int32(345210)
	goto L216
L216:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v603)+40)) = v781
	*(*int32)(unsafe.Add(mBase, uint32(v603)+36)) = int32(362226)
	*(*int32)(unsafe.Add(mBase, uint32(v603)+32)) = int32(535075)
	F_appendStringInfo(m, v603+int32(192), int32(605670), v603+int32(32))
	mBase = m.M
	v793 = m.ExcPending
	if v793 != 0 {
		goto L112
	} else {
		goto L221
	}
L217:
	;
	v778 = *(*int32)(unsafe.Add(mBase, uint32(v608)+16))
	if v778 != 0 {
		goto L218
	} else {
		goto L219
	}
L218:
	;
	v779 = int32(345210)
	goto L220
L219:
	;
	v779 = int32(362226)
	goto L220
L220:
	;
	v781 = v779
	goto L216
L221:
	;
	goto L213
L222:
	;
	F_appendStringInfoString(m, v603+int32(192), int32(645301))
	mBase = m.M
	v801 = m.ExcPending
	if v801 != 0 {
		goto L112
	} else {
		goto L225
	}
L223:
	;
	goto L224
L224:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v805 = m.ExcPending
	if v805 != 0 {
		goto L112
	} else {
		goto L226
	}
L225:
	;
	goto L224
L226:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v603)+16)) = l0
	F_errmsg(m, int32(325603), v603+int32(16))
	mBase = m.M
	v811 = m.ExcPending
	if v811 != 0 {
		goto L112
	} else {
		goto L227
	}
L227:
	;
	v812 = *(*int32)(unsafe.Add(mBase, uint32(v603)+192))
	*(*int32)(unsafe.Add(mBase, uint32(v603))) = v812
	F_errdetail_internal(m, int32(206576), v603)
	mBase = m.M
	v816 = m.ExcPending
	if v816 != 0 {
		goto L112
	} else {
		goto L228
	}
L228:
	;
	F_errfinish(m, int32(496310), int32(414), int32(212399))
	mBase = m.M
	v821 = m.ExcPending
	if v821 != 0 {
		goto L112
	} else {
		goto L229
	}
L229:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L230:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v603)+112)) = l0
	F_errmsg(m, int32(325729), v603+int32(112))
	mBase = m.M
	v831 = m.ExcPending
	if v831 != 0 {
		goto L112
	} else {
		goto L231
	}
L231:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v603)+100)) = v667
	*(*int32)(unsafe.Add(mBase, uint32(v603)+96)) = int32(1766396)
	F_errdetail(m, int32(666510), v603+int32(96))
	mBase = m.M
	v839 = m.ExcPending
	if v839 != 0 {
		goto L112
	} else {
		goto L232
	}
L232:
	;
	F_errfinish(m, int32(496310), int32(355), int32(212399))
	mBase = m.M
	v844 = m.ExcPending
	if v844 != 0 {
		goto L112
	} else {
		goto L233
	}
L233:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L234:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v15)+32)) = l0
	F_errmsg(m, int32(317553), v15+int32(32))
	mBase = m.M
	v856 = m.ExcPending
	if v856 != 0 {
		goto L112
	} else {
		goto L235
	}
L235:
	;
	F_errhint(m, int32(613787), int32(0))
	mBase = m.M
	v860 = m.ExcPending
	if v860 != 0 {
		goto L112
	} else {
		goto L236
	}
L236:
	;
	F_errfinish(m, int32(496310), int32(291), int32(17606))
	mBase = m.M
	v865 = m.ExcPending
	if v865 != 0 {
		goto L112
	} else {
		goto L237
	}
L237:
	;
	base.Wasm_trap_unreachable()
	for {
	}
}
