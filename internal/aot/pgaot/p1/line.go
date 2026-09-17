package p1

import (
	base "github.com/shibukawa/pgmem/internal/aot/pgaot/base"
	"unsafe"
)

func F_CopyReadLine(m *base.Module, l0 int32, l1 int32) int32 {
	mBase := m.M
	_ = mBase
	var v3 int32
	_ = v3
	var v19 int32
	_ = v19
	var v20 int32
	_ = v20
	var v27 int32
	_ = v27
	var v29 int32
	_ = v29
	var v30 int32
	_ = v30
	var v32 int32
	_ = v32
	var v33 int32
	_ = v33
	var v35 int32
	_ = v35
	var v40 int32
	_ = v40
	var v41 int32
	_ = v41
	var v42 int32
	_ = v42
	var v43 int32
	_ = v43
	var v44 int32
	_ = v44
	var v47 int32
	_ = v47
	var v50 int32
	_ = v50
	var v51 int32
	_ = v51
	var v54 int32
	_ = v54
	var v55 int32
	_ = v55
	var v58 int32
	_ = v58
	var v61 int32
	_ = v61
	var v66 int32
	_ = v66
	var v73 int32
	_ = v73
	var v75 int32
	_ = v75
	var v76 int32
	_ = v76
	var v77 int32
	_ = v77
	var v78 int32
	_ = v78
	var v79 int32
	_ = v79
	var v80 int32
	_ = v80
	var v101 int32
	_ = v101
	var v104 int32
	_ = v104
	var v105 int32
	_ = v105
	var v107 int32
	_ = v107
	var v110 int32
	_ = v110
	var v112 int32
	_ = v112
	var v113 int32
	_ = v113
	var v114 int32
	_ = v114
	var v115 int32
	_ = v115
	var v120 int32
	_ = v120
	var v121 int32
	_ = v121
	var v122 int32
	_ = v122
	var v123 int32
	_ = v123
	var v137 int32
	_ = v137
	var v139 int32
	_ = v139
	var v142 int32
	_ = v142
	var v144 int32
	_ = v144
	var v145 int32
	_ = v145
	var v146 int32
	_ = v146
	var v147 int32
	_ = v147
	var v149 int32
	_ = v149
	var v154 int32
	_ = v154
	var v156 int32
	_ = v156
	var v157 int32
	_ = v157
	var v162 int32
	_ = v162
	var v163 int32
	_ = v163
	var v172 int32
	_ = v172
	var v174 int32
	_ = v174
	var v178 int32
	_ = v178
	var v179 int32
	_ = v179
	var v182 int32
	_ = v182
	var v186 int32
	_ = v186
	var v187 int32
	_ = v187
	var v189 int32
	_ = v189
	var v192 int32
	_ = v192
	var v194 int32
	_ = v194
	var v199 int32
	_ = v199
	var v201 int32
	_ = v201
	var v206 int32
	_ = v206
	var v208 int32
	_ = v208
	var v211 int32
	_ = v211
	var v213 int32
	_ = v213
	var v216 int32
	_ = v216
	var v228 int32
	_ = v228
	var v230 int32
	_ = v230
	var v232 int32
	_ = v232
	var v235 int32
	_ = v235
	var v238 int32
	_ = v238
	var v243 int32
	_ = v243
	var v245 int32
	_ = v245
	var v247 int32
	_ = v247
	var v249 int32
	_ = v249
	var v252 int32
	_ = v252
	var v253 int32
	_ = v253
	var v255 int32
	_ = v255
	var v258 int32
	_ = v258
	var v260 int32
	_ = v260
	var v261 int32
	_ = v261
	var v262 int32
	_ = v262
	var v263 int32
	_ = v263
	var v264 int32
	_ = v264
	var v273 int32
	_ = v273
	var v274 int32
	_ = v274
	var v276 int32
	_ = v276
	var v281 int32
	_ = v281
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
	var v287 int32
	_ = v287
	var v289 int32
	_ = v289
	var v290 int32
	_ = v290
	var v292 int32
	_ = v292
	var v293 int32
	_ = v293
	var v297 int32
	_ = v297
	var v298 int32
	_ = v298
	var v301 int32
	_ = v301
	var v307 int32
	_ = v307
	var v309 int32
	_ = v309
	var v312 int32
	_ = v312
	var v313 int32
	_ = v313
	var v321 int32
	_ = v321
	var v322 int32
	_ = v322
	var v323 int32
	_ = v323
	var v325 int32
	_ = v325
	var v326 int32
	_ = v326
	var v329 int32
	_ = v329
	var v330 int32
	_ = v330
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
	var v339 int32
	_ = v339
	var v340 int32
	_ = v340
	var v347 int32
	_ = v347
	var v348 int32
	_ = v348
	var v352 int32
	_ = v352
	var v356 int32
	_ = v356
	var v361 int32
	_ = v361
	var v362 int32
	_ = v362
	var v366 int32
	_ = v366
	var v369 int32
	_ = v369
	var v370 int32
	_ = v370
	var v374 int32
	_ = v374
	var v376 int32
	_ = v376
	var v380 int32
	_ = v380
	var v382 int32
	_ = v382
	var v383 int32
	_ = v383
	var v387 int32
	_ = v387
	var v395 int32
	_ = v395
	var v397 int32
	_ = v397
	var v400 int32
	_ = v400
	var v410 int32
	_ = v410
	var v412 int32
	_ = v412
	var v413 int32
	_ = v413
	var v420 int32
	_ = v420
	var v421 int32
	_ = v421
	var v424 int32
	_ = v424
	var v427 int32
	_ = v427
	var v429 int64
	_ = v429
	var v436 int32
	_ = v436
	var v437 int32
	_ = v437
	var v442 int32
	_ = v442
	var v443 int32
	_ = v443
	var v451 int32
	_ = v451
	var v463 int32
	_ = v463
	var v466 int32
	_ = v466
	var v469 int32
	_ = v469
	var v472 int32
	_ = v472
	var v475 int32
	_ = v475
	var v478 int32
	_ = v478
	var v483 int32
	_ = v483
	var v487 int32
	_ = v487
	var v490 int32
	_ = v490
	var v493 int32
	_ = v493
	var v496 int32
	_ = v496
	var v499 int32
	_ = v499
	var v502 int32
	_ = v502
	var v507 int32
	_ = v507
	var v508 int32
	_ = v508
	var v517 int32
	_ = v517
	var v520 int32
	_ = v520
	var v523 int32
	_ = v523
	var v526 int32
	_ = v526
	var v529 int32
	_ = v529
	var v532 int32
	_ = v532
	var v537 int32
	_ = v537
	var v542 int32
	_ = v542
	var v553 int32
	_ = v553
	var v555 int32
	_ = v555
	var v558 int32
	_ = v558
	var v561 int32
	_ = v561
	var v569 int32
	_ = v569
	var v577 int32
	_ = v577
	var v579 int32
	_ = v579
	var v587 int32
	_ = v587
	var v593 int32
	_ = v593
	var v596 int32
	_ = v596
	var v600 int32
	_ = v600
	var v605 int32
	_ = v605
	var v621 int32
	_ = v621
	var v622 int32
	_ = v622
	var v632 int32
	_ = v632
	var v635 int32
	_ = v635
	var v639 int32
	_ = v639
	var v644 int32
	_ = v644
	var v646 int32
	_ = v646
	var v650 int32
	_ = v650
	var v652 int32
	_ = v652
	var v673 int32
	_ = v673
	var v674 int32
	_ = v674
	var v694 int32
	_ = v694
	var v696 int32
	_ = v696
	var v697 int32
	_ = v697
	var v700 int64
	_ = v700
	var v707 int32
	_ = v707
	var v710 int32
	_ = v710
	var v714 int32
	_ = v714
	var v719 int32
	_ = v719
	var v723 int32
	_ = v723
	var v726 int32
	_ = v726
	var v730 int32
	_ = v730
	var v735 int32
	_ = v735
	var v739 int32
	_ = v739
	var v742 int32
	_ = v742
	var v746 int32
	_ = v746
	var v751 int32
	_ = v751
	var v753 int32
	_ = v753
	var v754 int32
	_ = v754
	var v759 int32
	_ = v759
	var v764 int32
	_ = v764
	var v768 int32
	_ = v768
	var v770 int32
	_ = v770
	var v771 int32
	_ = v771
	var v774 int32
	_ = v774
	var v776 int32
	_ = v776
	var v778 int32
	_ = v778
	var v780 int32
	_ = v780
	var v782 int32
	_ = v782
	var v784 int32
	_ = v784
	var v786 int32
	_ = v786
	var v788 int32
	_ = v788
	var v790 int32
	_ = v790
	var v792 int32
	_ = v792
	var v794 int32
	_ = v794
	var v796 int32
	_ = v796
	var v803 int32
	_ = v803
	var v815 int32
	_ = v815
	v3 = int32(0)
	v19 = l0 + int32(288)
	v20 = *(*int32)(unsafe.Add(mBase, uint32(v19)))
	*(*uint8)(unsafe.Add(mBase, uint32(v20))) = uint8(v3)
	*(*int32)(unsafe.Add(mBase, uint32(v19)+12)) = v3
	*(*int32)(unsafe.Add(mBase, uint32(v19)+4)) = v3
	goto L1
L1:
	;
	v27 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(l0)+304)) = uint8(v27)
	if l1 != 0 {
		goto L2
	} else {
		goto L3
	}
L2:
	;
	v29 = *(*int32)(unsafe.Add(mBase, uint32(l0)+96))
	v30 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v29))))
	v32 = *(*int32)(unsafe.Add(mBase, uint32(l0)+92))
	v33 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v32))))
	if v30 != v33 {
		goto L5
	} else {
		goto L6
	}
L3:
	;
	v40 = v3
	v41 = v3
	goto L4
L4:
	;
	v42 = *(*int32)(unsafe.Add(mBase, uint32(l0)+316))
	v43 = *(*int32)(unsafe.Add(mBase, uint32(l0)+308))
	v44 = *(*int32)(unsafe.Add(mBase, uint32(l0)+312))
	v47 = v44
	v50 = v42
	v51 = v44
	v54 = v3
	v55 = v3
	v58 = v3
	v61 = v3
	goto L17
L5:
	;
	v35 = v30
	goto L7
L6:
	;
	v35 = int32(0)
	goto L7
L7:
	;
	v40 = base.I32_extend8_s(v33)
	v41 = base.I32_extend8_s(v35)
	goto L4
L8:
	;
	v815 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(l0)+304)) = uint8(v815)
	return v803
L9:
	;
	if v383 < v759 {
		goto L221
	} else {
		goto L222
	}
L10:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+16)) = v754
	v759 = v753
	goto L9
L11:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v739 = m.ExcPending
	if v739 != 0 {
		goto L27
	} else {
		goto L217
	}
L12:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v723 = m.ExcPending
	if v723 != 0 {
		goto L27
	} else {
		goto L213
	}
L13:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v707 = m.ExcPending
	if v707 != 0 {
		goto L27
	} else {
		goto L209
	}
L14:
	;
	v673 = int32(1)
	v674 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	if v674 != v673 {
		v803 = v673
		goto L8
	} else {
		goto L204
	}
L15:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+312)) = v652
	goto L14
L16:
	;
	if v395 <= v383 {
		goto L14
	} else {
		goto L202
	}
L17:
	;
	if v54|base.B2i32(v50 <= v47) == int32(0) {
		goto L20
	} else {
		goto L21
	}
L18:
	;
	v587 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v577+v43))))
	switch v587 - int32(10) {
	case 0, 3:
		goto L188
	default:
		goto L189
	}
L19:
	;
	v395 = v380 + int32(1)
	v397 = int32(*(*int8)(unsafe.Add(mBase, uint32(v380+v43))))
	if l1 == int32(0) {
		v436 = v58
		v437 = v61
		goto L121
	} else {
		goto L122
	}
L20:
	;
	v380 = v47
	v382 = v50
	v383 = v51
	v387 = v55
	goto L19
L21:
	;
	goto L22
L22:
	;
	v66 = *(*int32)(unsafe.Add(mBase, uint32(l0)+308))
	if v47 <= v51 {
		goto L24
	} else {
		goto L25
	}
L23:
	;
	v80 = *(*int32)(unsafe.Add(mBase, uint32(l0)+324))
	if v79 == v80 {
		goto L29
	} else {
		goto L30
	}
L24:
	;
	v77 = v51
	v78 = v50
	v79 = v66
	goto L23
L25:
	;
	goto L26
L26:
	;
	F_appendBinaryStringInfo(m, v19, v51+v66, v47-v51)
	mBase = m.M
	v73 = m.ExcPending
	if v73 != 0 {
		goto L27
	} else {
		goto L28
	}
L27:
	;
	return int32(0)
L28:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+312)) = v47
	v75 = *(*int32)(unsafe.Add(mBase, uint32(l0)+308))
	v76 = *(*int32)(unsafe.Add(mBase, uint32(l0)+316))
	v77 = v47
	v78 = v76
	v79 = v75
	goto L23
L29:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+328)) = v77
	goto L31
L30:
	;
	goto L31
L31:
	;
	goto L32
L32:
	;
	v101 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+24)))
	if v101 == int32(0) {
		goto L35
	} else {
		goto L36
	}
L34:
	;
	v321 = *(*int32)(unsafe.Add(mBase, uint32(l0)+316))
	v322 = *(*int32)(unsafe.Add(mBase, uint32(l0)+312))
	v323 = v321 - v322
	if v78-v77 < v323 {
		goto L105
	} else {
		goto L106
	}
L35:
	;
	v104 = *(*int32)(unsafe.Add(mBase, uint32(l0)+332))
	v105 = *(*int32)(unsafe.Add(mBase, uint32(l0)+316))
	if v104 == v105 {
		goto L38
	} else {
		goto L39
	}
L36:
	;
	goto L37
L37:
	;
	v252 = *(*int32)(unsafe.Add(mBase, uint32(l0)+332))
	v253 = *(*int32)(unsafe.Add(mBase, uint32(l0)+328))
	if v252 == v253 {
		goto L86
	} else {
		goto L87
	}
L38:
	;
	v107 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+336)))
	if v107 != int32(1) {
		goto L34
	} else {
		goto L41
	}
L39:
	;
	goto L40
L40:
	;
	v112 = *(*int32)(unsafe.Add(mBase, uint32(l0)+324))
	v113 = v112 + v105
	v114 = v104 - v105
	v115 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	if base.Ui32(v115) <= base.Ui32(int32(41)) {
		goto L43
	} else {
		goto L44
	}
L41:
	;
	v110 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(l0)+320)) = uint8(v110)
	goto L34
L42:
	;
	if v232 == int32(0) {
		goto L75
	} else {
		goto L76
	}
L43:
	;
	v120 = *(*int32)(unsafe.Add(mBase, uint32(v115*int32(28))+uint32(_c_F_CopyReadLine[0])))
	v121 = m.T0[v120].(func(*base.Module, int32, int32) int32)(m, v113, v114)
	mBase = m.M
	v122 = m.ExcPending
	if v122 != 0 {
		goto L27
	} else {
		goto L46
	}
L44:
	;
	goto L45
L45:
	;
	v123 = int32(0)
	if base.B2i32(v113&int32(3) == v123)|base.B2i32(v114 == v123) != 0 {
		v154 = v113
		v156 = v114
		v157 = base.B2i32(v114 != v123)
		goto L50
	} else {
		goto L51
	}
L46:
	;
	v232 = v121
	goto L42
L47:
	;
	if v228 != 0 {
		goto L72
	} else {
		goto L73
	}
L48:
	;
	v228 = int32(0)
	goto L47
L49:
	;
	v206 = v199
	v208 = v201
	goto L66
L50:
	;
	if v157 == int32(0) {
		goto L48
	} else {
		goto L57
	}
L51:
	;
	v137 = v113
	v139 = v114
	goto L52
L52:
	;
	v142 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v137))))
	if v142 == int32(0) {
		v199 = v137
		v201 = v139
		goto L49
	} else {
		goto L54
	}
L53:
	;
	v154 = v149
	v156 = v145
	v157 = v147
	goto L50
L54:
	;
	v144 = int32(1)
	v145 = v139 - v144
	v146 = int32(0)
	v147 = base.B2i32(v145 != v146)
	v149 = v137 + v144
	if v149&int32(3) == v146 {
		v154 = v149
		v156 = v145
		v157 = v147
		goto L50
	} else {
		goto L55
	}
L55:
	;
	if v145 != 0 {
		v137 = v149
		v139 = v145
		goto L52
	} else {
		goto L56
	}
L56:
	;
	goto L53
L57:
	;
	v162 = int32(0)
	v163 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v154))))
	if base.B2i32(v162 == v163)|base.B2i32(base.Ui32(v156) < base.Ui32(int32(4))) == v162 {
		goto L58
	} else {
		goto L59
	}
L58:
	;
	v172 = v154
	v174 = v156
	goto L61
L59:
	;
	v192 = v154
	v194 = v156
	goto L60
L60:
	;
	if v194 == int32(0) {
		goto L48
	} else {
		goto L65
	}
L61:
	;
	v178 = *(*int32)(unsafe.Add(mBase, uint32(v172)))
	v179 = v178 ^ int32(0)
	v182 = int32(-2139062144)
	if (int32(16843008)-v179|v179)&v182 != v182 {
		v199 = v172
		v201 = v174
		goto L49
	} else {
		goto L63
	}
L62:
	;
	v192 = v187
	v194 = v189
	goto L60
L63:
	;
	v186 = int32(4)
	v187 = v172 + v186
	v189 = v174 - v186
	if base.Ui32(int32(3)) < base.Ui32(v189) {
		v172 = v187
		v174 = v189
		goto L61
	} else {
		goto L64
	}
L64:
	;
	goto L62
L65:
	;
	v199 = v192
	v201 = v194
	goto L49
L66:
	;
	v211 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v206))))
	if int32(0) == v211 {
		goto L68
	} else {
		goto L69
	}
L67:
	;
	goto L48
L68:
	;
	v228 = v206
	goto L47
L69:
	;
	goto L70
L70:
	;
	v213 = int32(1)
	v216 = v208 - v213
	if v216 != 0 {
		v206 = v206 + v213
		v208 = v216
		goto L66
	} else {
		goto L71
	}
L71:
	;
	goto L67
L72:
	;
	v230 = v228 - v113
	goto L74
L73:
	;
	v230 = v114
	goto L74
L74:
	;
	v232 = v230
	goto L42
L75:
	;
	v235 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+336)))
	if v235 == int32(0) {
		goto L78
	} else {
		goto L79
	}
L76:
	;
	goto L77
L77:
	;
	v249 = *(*int32)(unsafe.Add(mBase, uint32(l0)+316))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+316)) = v249 + v232
	goto L34
L78:
	;
	v238 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	if base.Ui32(v238) <= base.Ui32(int32(41)) {
		goto L82
	} else {
		goto L83
	}
L79:
	;
	goto L80
L80:
	;
	v247 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(l0)+321)) = uint8(v247)
	goto L34
L81:
	;
	if v114 < v245 {
		goto L34
	} else {
		goto L85
	}
L82:
	;
	v243 = *(*int32)(unsafe.Add(mBase, uint32(v238*int32(28))+uint32(_c_F_CopyReadLine[1])))
	v245 = v243
	goto L84
L83:
	;
	v245 = int32(1)
	goto L84
L84:
	;
	goto L81
L85:
	;
	goto L80
L86:
	;
	v255 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+336)))
	if v255 != int32(1) {
		goto L34
	} else {
		goto L89
	}
L87:
	;
	goto L88
L88:
	;
	v260 = *(*int32)(unsafe.Add(mBase, uint32(l0)+316))
	v261 = *(*int32)(unsafe.Add(mBase, uint32(l0)+312))
	v262 = v260 - v261
	v263 = *(*int32)(unsafe.Add(mBase, uint32(l0)+308))
	v264 = int32(0)
	if base.B2i32(v261 <= v264)|base.B2i32(v262 <= v264) == v264 {
		goto L90
	} else {
		goto L91
	}
L89:
	;
	v258 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(l0)+320)) = uint8(v258)
	goto L34
L90:
	;
	if v262 != 0 {
		goto L93
	} else {
		goto L94
	}
L91:
	;
	v274 = v263
	goto L92
L92:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+316)) = v262
	v276 = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(l0)+312)) = v276
	*(*uint8)(unsafe.Add(mBase, uint32(v262+v274))) = uint8(v276)
	v281 = *(*int32)(unsafe.Add(mBase, uint32(l0)+316))
	v282 = *(*int32)(unsafe.Add(mBase, uint32(l0)+308))
	v283 = *(*int32)(unsafe.Add(mBase, uint32(l0)+328))
	v284 = *(*int32)(unsafe.Add(mBase, uint32(l0)+332))
	v285 = *(*int32)(unsafe.Add(mBase, uint32(l0)+324))
	v286 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
	v287 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v289 = *(*int32)(unsafe.Add(mBase, _c_F_CopyReadLine[2]))
	v290 = *(*int32)(unsafe.Add(mBase, uint32(v289)+4))
	goto L96
L93:
	;
	base.MemoryCopy(m, v263, v261+v263, v262)
	goto L95
L94:
	;
	goto L95
L95:
	;
	v273 = *(*int32)(unsafe.Add(mBase, uint32(l0)+308))
	v274 = v273
	goto L92
L96:
	;
	v292 = v284 - v283
	v293 = v282 + v281
	v297 = F_pg_do_encoding_conversion_buf(m, v286, v287, v290, v285+v283, v292, v293, int32(_a_F_CopyReadLine_0)-v281, int32(1))
	mBase = m.M
	v298 = m.ExcPending
	if v298 != 0 {
		goto L27
	} else {
		goto L97
	}
L97:
	;
	if v297 == int32(0) {
		goto L98
	} else {
		goto L99
	}
L98:
	;
	v301 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+336)))
	if base.B2i32(v301 == int32(0))&base.B2i32(v292 < int32(16)) != 0 {
		goto L34
	} else {
		goto L101
	}
L99:
	;
	goto L100
L100:
	;
	v309 = *(*int32)(unsafe.Add(mBase, uint32(l0)+328))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+328)) = v309 + v297
	v312 = *(*int32)(unsafe.Add(mBase, uint32(l0)+316))
	v313 = F_strlen(m, v293)
	mBase = m.M
	*(*int32)(unsafe.Add(mBase, uint32(l0)+316)) = v312 + v313
	goto L34
L101:
	;
	v307 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(l0)+321)) = uint8(v307)
	goto L34
L102:
	;
	F_CopyLoadRawBuf(m, l0)
	mBase = m.M
	v376 = m.ExcPending
	if v376 != 0 {
		goto L27
	} else {
		goto L120
	}
L103:
	;
	v369 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v370 = *(*int32)(unsafe.Add(mBase, uint32(l0)+324))
	F_report_invalid_encoding(m, v369, v370+v321, v329-v321)
	mBase = m.M
	v374 = m.ExcPending
	if v374 != 0 {
		goto L27
	} else {
		goto L119
	}
L104:
	;
	if v323 <= int32(0) {
		goto L14
	} else {
		goto L118
	}
L105:
	;
	v325 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+320)))
	v366 = v325
	goto L104
L106:
	;
	goto L107
L107:
	;
	v326 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+321)))
	if v326 == int32(1) {
		goto L108
	} else {
		goto L109
	}
L108:
	;
	v329 = *(*int32)(unsafe.Add(mBase, uint32(l0)+332))
	v330 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+24)))
	if v330 == int32(0) {
		goto L103
	} else {
		goto L111
	}
L109:
	;
	goto L110
L110:
	;
	v362 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+320)))
	if v362 == int32(0) {
		goto L102
	} else {
		goto L117
	}
L111:
	;
	v333 = *(*int32)(unsafe.Add(mBase, uint32(l0)+308))
	v334 = *(*int32)(unsafe.Add(mBase, uint32(l0)+328))
	v335 = *(*int32)(unsafe.Add(mBase, uint32(l0)+324))
	v336 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
	v337 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v339 = *(*int32)(unsafe.Add(mBase, _c_F_CopyReadLine[2]))
	v340 = *(*int32)(unsafe.Add(mBase, uint32(v339)+4))
	goto L112
L112:
	;
	v347 = F_pg_do_encoding_conversion_buf(m, v336, v337, v340, v335+v334, v329-v334, v333+v321, int32(_a_F_CopyReadLine_0)-v321, int32(0))
	mBase = m.M
	v348 = m.ExcPending
	if v348 != 0 {
		goto L27
	} else {
		goto L113
	}
L113:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v352 = m.ExcPending
	if v352 != 0 {
		goto L27
	} else {
		goto L114
	}
L114:
	;
	F_errmsg_internal(m, int32(_a_F_CopyReadLine_1), int32(0))
	mBase = m.M
	v356 = m.ExcPending
	if v356 != 0 {
		goto L27
	} else {
		goto L115
	}
L115:
	;
	F_errfinish(m, int32(_a_F_CopyReadLine_2), int32(579), int32(_a_F_CopyReadLine_3))
	mBase = m.M
	v361 = m.ExcPending
	if v361 != 0 {
		goto L27
	} else {
		goto L116
	}
L116:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L117:
	;
	v366 = int32(1)
	goto L104
L118:
	;
	v380 = v322
	v382 = v321
	v383 = v322
	v387 = v366
	goto L19
L119:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L120:
	;
	goto L32
L121:
	;
	switch v397&int32(255) - int32(10) {
	case 0:
		goto L130
	default:
		goto L129
	case 3:
		goto L131
	}
L122:
	;
	v400 = int32(1)
	if (base.B2i32(v397 != int32(13))|base.B2i32(v395 < v382)|v387)&v400 == int32(0) {
		v47 = v380
		v50 = v382
		v51 = v383
		v54 = v400
		v55 = v387
		goto L17
	} else {
		goto L123
	}
L123:
	;
	v410 = base.B2i32(v397 == v41)
	v412 = v61 ^ v410&v58
	v413 = v410 & v412
	if v58 != (v412|base.B2i32(v397 != v40))&int32(1) {
		v436 = int32(0)
		v437 = v413
		goto L121
	} else {
		goto L124
	}
L124:
	;
	v420 = int32(1)
	v421 = int32(0)
	v424 = *(*int32)(unsafe.Add(mBase, uint32(l0)+16))
	if v424 == v420 {
		goto L125
	} else {
		goto L126
	}
L125:
	;
	v427 = int32(10)
	goto L127
L126:
	;
	v427 = int32(13)
	goto L127
L127:
	;
	if v427 != v397 {
		v47 = v395
		v50 = v382
		v51 = v383
		v54 = v421
		v55 = v387
		v58 = v420
		v61 = v413
		goto L17
	} else {
		goto L128
	}
L128:
	;
	v429 = *(*int64)(unsafe.Add(mBase, uint32(l0)+184))
	*(*int64)(unsafe.Add(mBase, uint32(l0)+184)) = v429 + int64(1)
	v47 = v395
	v50 = v382
	v51 = v383
	v54 = v421
	v55 = v387
	v58 = v420
	v61 = v413
	goto L17
L129:
	;
	if l1|base.B2i32(v397 != int32(92)) != 0 {
		v47 = v395
		v50 = v382
		v51 = v383
		v54 = int32(0)
		v55 = v387
		v58 = v436
		v61 = v437
		goto L17
	} else {
		goto L177
	}
L130:
	;
	v508 = *(*int32)(unsafe.Add(mBase, uint32(l0)+16))
	if v508&int32(-2) != int32(2) {
		goto L163
	} else {
		goto L164
	}
L131:
	;
	v442 = *(*int32)(unsafe.Add(mBase, uint32(l0)+16))
	switch v442 {
	case 0, 3:
		goto L133
	case 1:
		goto L132
	default:
		v759 = v395
		goto L9
	}
L132:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v487 = m.ExcPending
	if v487 != 0 {
		goto L27
	} else {
		goto L152
	}
L133:
	;
	v443 = int32(1)
	if (base.B2i32(v395 < v382)|v387)&v443 == int32(0) {
		v47 = v380
		v50 = v382
		v51 = v383
		v54 = v443
		v55 = v387
		v58 = v436
		v61 = v437
		goto L17
	} else {
		goto L134
	}
L134:
	;
	v451 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v395+v43))))
	if v451 == int32(10) {
		goto L135
	} else {
		goto L136
	}
L135:
	;
	v753 = v380 + int32(2)
	v754 = int32(3)
	goto L10
L136:
	;
	goto L137
L137:
	;
	if v442 != int32(3) {
		goto L138
	} else {
		goto L139
	}
L138:
	;
	v753 = v395
	v754 = int32(2)
	goto L10
L139:
	;
	goto L140
L140:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v463 = m.ExcPending
	if v463 != 0 {
		goto L27
	} else {
		goto L141
	}
L141:
	;
	F_errcode(m, int32(67240066))
	mBase = m.M
	v466 = m.ExcPending
	if v466 != 0 {
		goto L27
	} else {
		goto L142
	}
L142:
	;
	if l1 != 0 {
		goto L143
	} else {
		goto L144
	}
L143:
	;
	v469 = int32(_a_F_CopyReadLine_4)
	goto L145
L144:
	;
	v469 = int32(_a_F_CopyReadLine_5)
	goto L145
L145:
	;
	F_errmsg(m, v469, int32(0))
	mBase = m.M
	v472 = m.ExcPending
	if v472 != 0 {
		goto L27
	} else {
		goto L146
	}
L146:
	;
	if l1 != 0 {
		goto L147
	} else {
		goto L148
	}
L147:
	;
	v475 = int32(_a_F_CopyReadLine_6)
	goto L149
L148:
	;
	v475 = int32(_a_F_CopyReadLine_7)
	goto L149
L149:
	;
	F_errhint(m, v475, int32(0))
	mBase = m.M
	v478 = m.ExcPending
	if v478 != 0 {
		goto L27
	} else {
		goto L150
	}
L150:
	;
	F_errfinish(m, int32(_a_F_CopyReadLine_2), int32(1398), int32(_a_F_CopyReadLine_8))
	mBase = m.M
	v483 = m.ExcPending
	if v483 != 0 {
		goto L27
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
	F_errcode(m, int32(67240066))
	mBase = m.M
	v490 = m.ExcPending
	if v490 != 0 {
		goto L27
	} else {
		goto L153
	}
L153:
	;
	if l1 != 0 {
		goto L154
	} else {
		goto L155
	}
L154:
	;
	v493 = int32(_a_F_CopyReadLine_4)
	goto L156
L155:
	;
	v493 = int32(_a_F_CopyReadLine_5)
	goto L156
L156:
	;
	F_errmsg(m, v493, int32(0))
	mBase = m.M
	v496 = m.ExcPending
	if v496 != 0 {
		goto L27
	} else {
		goto L157
	}
L157:
	;
	if l1 != 0 {
		goto L158
	} else {
		goto L159
	}
L158:
	;
	v499 = int32(_a_F_CopyReadLine_6)
	goto L160
L159:
	;
	v499 = int32(_a_F_CopyReadLine_7)
	goto L160
L160:
	;
	F_errhint(m, v499, int32(0))
	mBase = m.M
	v502 = m.ExcPending
	if v502 != 0 {
		goto L27
	} else {
		goto L161
	}
L161:
	;
	F_errfinish(m, int32(_a_F_CopyReadLine_2), int32(1415), int32(_a_F_CopyReadLine_8))
	mBase = m.M
	v507 = m.ExcPending
	if v507 != 0 {
		goto L27
	} else {
		goto L162
	}
L162:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L163:
	;
	v753 = v395
	v754 = int32(1)
	goto L10
L164:
	;
	goto L165
L165:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v517 = m.ExcPending
	if v517 != 0 {
		goto L27
	} else {
		goto L166
	}
L166:
	;
	F_errcode(m, int32(67240066))
	mBase = m.M
	v520 = m.ExcPending
	if v520 != 0 {
		goto L27
	} else {
		goto L167
	}
L167:
	;
	if l1 != 0 {
		goto L168
	} else {
		goto L169
	}
L168:
	;
	v523 = int32(_a_F_CopyReadLine_9)
	goto L170
L169:
	;
	v523 = int32(_a_F_CopyReadLine_10)
	goto L170
L170:
	;
	F_errmsg(m, v523, int32(0))
	mBase = m.M
	v526 = m.ExcPending
	if v526 != 0 {
		goto L27
	} else {
		goto L171
	}
L171:
	;
	if l1 != 0 {
		goto L172
	} else {
		goto L173
	}
L172:
	;
	v529 = int32(_a_F_CopyReadLine_11)
	goto L174
L173:
	;
	v529 = int32(_a_F_CopyReadLine_12)
	goto L174
L174:
	;
	F_errhint(m, v529, int32(0))
	mBase = m.M
	v532 = m.ExcPending
	if v532 != 0 {
		goto L27
	} else {
		goto L175
	}
L175:
	;
	F_errfinish(m, int32(_a_F_CopyReadLine_2), int32(1431), int32(_a_F_CopyReadLine_8))
	mBase = m.M
	v537 = m.ExcPending
	if v537 != 0 {
		goto L27
	} else {
		goto L176
	}
L176:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L177:
	;
	v542 = int32(1)
	if (base.B2i32(v395 < v382)|v387)&v542 == int32(0) {
		v47 = v380
		v50 = v382
		v51 = v383
		v54 = v542
		v55 = v387
		v58 = v436
		v61 = v437
		goto L17
	} else {
		goto L178
	}
L178:
	;
	if base.B2i32(v382 <= v395)&v387 != 0 {
		goto L16
	} else {
		goto L179
	}
L179:
	;
	v553 = v380 + int32(2)
	v555 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v395+v43))))
	if v555 != int32(46) {
		v47 = v553
		v50 = v382
		v51 = v383
		v54 = int32(0)
		v55 = v387
		v58 = v436
		v61 = v437
		goto L17
	} else {
		goto L180
	}
L180:
	;
	v558 = *(*int32)(unsafe.Add(mBase, uint32(l0)+16))
	if v558 == int32(3) {
		goto L181
	} else {
		goto L182
	}
L181:
	;
	v561 = int32(1)
	if (base.B2i32(v553 < v382)|v387)&v561 == int32(0) {
		v47 = v380
		v50 = v382
		v51 = v383
		v54 = v561
		v55 = v387
		v58 = v436
		v61 = v437
		goto L17
	} else {
		goto L184
	}
L182:
	;
	v577 = v553
	goto L183
L183:
	;
	v579 = int32(1)
	if (base.B2i32(v577 < v382)|v387)&v579 == int32(0) {
		v47 = v380
		v50 = v382
		v51 = v383
		v54 = v579
		v55 = v387
		v58 = v436
		v61 = v437
		goto L17
	} else {
		goto L187
	}
L184:
	;
	v569 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v553+v43))))
	if v569 == int32(10) {
		goto L13
	} else {
		goto L185
	}
L185:
	;
	if v569 != int32(13) {
		goto L12
	} else {
		goto L186
	}
L186:
	;
	v577 = v380 + int32(3)
	goto L183
L187:
	;
	goto L18
L188:
	;
	if (base.B2i32(v558 == int32(1))|base.B2i32(v558 == int32(3)))&base.B2i32(v587 != int32(10))|base.B2i32(v558 == int32(2))&base.B2i32(v587 != int32(13)) != 0 {
		goto L11
	} else {
		goto L194
	}
L189:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v593 = m.ExcPending
	if v593 != 0 {
		goto L27
	} else {
		goto L190
	}
L190:
	;
	F_errcode(m, int32(67240066))
	mBase = m.M
	v596 = m.ExcPending
	if v596 != 0 {
		goto L27
	} else {
		goto L191
	}
L191:
	;
	F_errmsg(m, int32(_a_F_CopyReadLine_13), int32(0))
	mBase = m.M
	v600 = m.ExcPending
	if v600 != 0 {
		goto L27
	} else {
		goto L192
	}
L192:
	;
	F_errfinish(m, int32(_a_F_CopyReadLine_2), int32(1484), int32(_a_F_CopyReadLine_8))
	mBase = m.M
	v605 = m.ExcPending
	if v605 != 0 {
		goto L27
	} else {
		goto L193
	}
L193:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L194:
	;
	v621 = *(*int32)(unsafe.Add(mBase, uint32(l0)+292))
	v622 = int32(0)
	if base.B2i32(v383 < v380)|base.B2i32(v622 < v621) == v622 {
		goto L195
	} else {
		goto L196
	}
L195:
	;
	v652 = v577 + int32(1)
	goto L15
L196:
	;
	goto L197
L197:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v632 = m.ExcPending
	if v632 != 0 {
		goto L27
	} else {
		goto L198
	}
L198:
	;
	F_errcode(m, int32(67240066))
	mBase = m.M
	v635 = m.ExcPending
	if v635 != 0 {
		goto L27
	} else {
		goto L199
	}
L199:
	;
	F_errmsg(m, int32(_a_F_CopyReadLine_13), int32(0))
	mBase = m.M
	v639 = m.ExcPending
	if v639 != 0 {
		goto L27
	} else {
		goto L200
	}
L200:
	;
	F_errfinish(m, int32(_a_F_CopyReadLine_2), int32(1500), int32(_a_F_CopyReadLine_8))
	mBase = m.M
	v644 = m.ExcPending
	if v644 != 0 {
		goto L27
	} else {
		goto L201
	}
L201:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L202:
	;
	v646 = *(*int32)(unsafe.Add(mBase, uint32(l0)+308))
	F_appendBinaryStringInfo(m, v19, v646+v383, v395-v383)
	mBase = m.M
	v650 = m.ExcPending
	if v650 != 0 {
		goto L27
	} else {
		goto L203
	}
L203:
	;
	v652 = v395
	goto L15
L204:
	;
	goto L205
L205:
	;
	v694 = *(*int32)(unsafe.Add(mBase, uint32(l0)+308))
	v696 = F_CopyGetData(m, l0, v694, int32(_a_F_CopyReadLine_14))
	mBase = m.M
	v697 = m.ExcPending
	if v697 != 0 {
		goto L27
	} else {
		goto L207
	}
L206:
	;
	v700 = int64(0)
	*(*int64)(unsafe.Add(mBase, uint32(l0)+328)) = v700
	*(*int64)(unsafe.Add(mBase, uint32(l0)+312)) = v700
	v803 = v673
	goto L8
L207:
	;
	if int32(0) < v696 {
		goto L205
	} else {
		goto L208
	}
L208:
	;
	goto L206
L209:
	;
	F_errcode(m, int32(67240066))
	mBase = m.M
	v710 = m.ExcPending
	if v710 != 0 {
		goto L27
	} else {
		goto L210
	}
L210:
	;
	F_errmsg(m, int32(_a_F_CopyReadLine_15), int32(0))
	mBase = m.M
	v714 = m.ExcPending
	if v714 != 0 {
		goto L27
	} else {
		goto L211
	}
L211:
	;
	F_errfinish(m, int32(_a_F_CopyReadLine_2), int32(1469), int32(_a_F_CopyReadLine_8))
	mBase = m.M
	v719 = m.ExcPending
	if v719 != 0 {
		goto L27
	} else {
		goto L212
	}
L212:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L213:
	;
	F_errcode(m, int32(67240066))
	mBase = m.M
	v726 = m.ExcPending
	if v726 != 0 {
		goto L27
	} else {
		goto L214
	}
L214:
	;
	F_errmsg(m, int32(_a_F_CopyReadLine_13), int32(0))
	mBase = m.M
	v730 = m.ExcPending
	if v730 != 0 {
		goto L27
	} else {
		goto L215
	}
L215:
	;
	F_errfinish(m, int32(_a_F_CopyReadLine_2), int32(1473), int32(_a_F_CopyReadLine_8))
	mBase = m.M
	v735 = m.ExcPending
	if v735 != 0 {
		goto L27
	} else {
		goto L216
	}
L216:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L217:
	;
	F_errcode(m, int32(67240066))
	mBase = m.M
	v742 = m.ExcPending
	if v742 != 0 {
		goto L27
	} else {
		goto L218
	}
L218:
	;
	F_errmsg(m, int32(_a_F_CopyReadLine_15), int32(0))
	mBase = m.M
	v746 = m.ExcPending
	if v746 != 0 {
		goto L27
	} else {
		goto L219
	}
L219:
	;
	F_errfinish(m, int32(_a_F_CopyReadLine_2), int32(1491), int32(_a_F_CopyReadLine_8))
	mBase = m.M
	v751 = m.ExcPending
	if v751 != 0 {
		goto L27
	} else {
		goto L220
	}
L220:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L221:
	;
	v764 = *(*int32)(unsafe.Add(mBase, uint32(l0)+308))
	F_appendBinaryStringInfo(m, v19, v764+v383, v759-v383)
	mBase = m.M
	v768 = m.ExcPending
	if v768 != 0 {
		goto L27
	} else {
		goto L224
	}
L222:
	;
	goto L223
L223:
	;
	v770 = int32(0)
	v771 = *(*int32)(unsafe.Add(mBase, uint32(l0)+16))
	switch v771 - int32(1) {
	case 0:
		goto L227
	case 1:
		goto L226
	case 2:
		goto L225
	default:
		v803 = v770
		goto L8
	}
L224:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+312)) = v759
	goto L223
L225:
	;
	v790 = *(*int32)(unsafe.Add(mBase, uint32(l0)+292))
	v792 = v790 - int32(2)
	*(*int32)(unsafe.Add(mBase, uint32(l0)+292)) = v792
	v794 = *(*int32)(unsafe.Add(mBase, uint32(l0)+288))
	v796 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v794+v792))) = uint8(v796)
	v803 = v770
	goto L8
L226:
	;
	v782 = *(*int32)(unsafe.Add(mBase, uint32(l0)+292))
	v784 = v782 - int32(1)
	*(*int32)(unsafe.Add(mBase, uint32(l0)+292)) = v784
	v786 = *(*int32)(unsafe.Add(mBase, uint32(l0)+288))
	v788 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v786+v784))) = uint8(v788)
	v803 = v770
	goto L8
L227:
	;
	v774 = *(*int32)(unsafe.Add(mBase, uint32(l0)+292))
	v776 = v774 - int32(1)
	*(*int32)(unsafe.Add(mBase, uint32(l0)+292)) = v776
	v778 = *(*int32)(unsafe.Add(mBase, uint32(l0)+288))
	v780 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v778+v776))) = uint8(v780)
	v803 = v770
	goto L8
}
func F_line_construct_pp(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v11 int32
	_ = v11
	var v12 int32
	_ = v12
	var v14 int32
	_ = v14
	var v17 int32
	_ = v17
	var v18 float64
	_ = v18
	var v24 float64
	_ = v24
	var v26 int64
	_ = v26
	var v27 int64
	_ = v27
	var v28 float64
	_ = v28
	var v31 int64
	_ = v31
	var v38 float64
	_ = v38
	var v45 int64
	_ = v45
	var v50 float64
	_ = v50
	var v68 int32
	_ = v68
	var v75 float64
	_ = v75
	var v79 int64
	_ = v79
	var v80 float64
	_ = v80
	var v83 int64
	_ = v83
	var v102 int32
	_ = v102
	var v105 int32
	_ = v105
	var v109 int32
	_ = v109
	var v114 int32
	_ = v114
	var v122 float64
	_ = v122
	var v123 int32
	_ = v123
	var v125 int32
	_ = v125
	v11 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
	v12 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v14 = F_palloc(m, int32(24))
	mBase = m.M
	v17 = m.ExcPending
	if v17 != 0 {
		return int32(0)
	} else {
		v18 = *(*float64)(unsafe.Add(mBase, uint32(v12)))
		if base.Ui64(base.I64_reinterpret_f64(v18)&int64(9223372036854775807)) <= base.Ui64(int64(9218868437227405312)) {
			v24 = *(*float64)(unsafe.Add(mBase, uint32(v11)))
			v26 = int64(9223372036854775807)
			v27 = base.I64_reinterpret_f64(v24) & v26
			v28 = *(*float64)(unsafe.Add(mBase, uint32(v12)+8))
			v31 = base.I64_reinterpret_f64(v28) & v26
			if base.Ui64(int64(9218868437227405313)) <= base.Ui64(v31) {
				v68 = base.B2i32(base.Ui64(v27) < base.Ui64(int64(9218868437227405313)))
				if base.B2i32(v68 == int32(0))|base.F64_ne(v18, v24) != 0 {
					v122 = F_point_sl(m, v12, v11)
					mBase = m.M
					v123 = m.ExcPending
					if v123 != 0 {
						return int32(0)
					} else {
						F_line_construct(m, v14, v12, v122)
						mBase = m.M
						v125 = m.ExcPending
						if v125 != 0 {
							return int32(0)
						} else {
							return v14
						}
					}
				} else {
					v75 = v28
					v79 = v31
					v80 = *(*float64)(unsafe.Add(mBase, uint32(v11)+8))
					v83 = base.I64_reinterpret_f64(v80) & int64(9223372036854775807)
					if base.Ui64(v79) <= base.Ui64(int64(9218868437227405312)) {
						if base.B2i32(base.Ui64(int64(9218868437227405313)) <= base.Ui64(v83))|base.F64_ne(v80, v75) != 0 {
							v122 = F_point_sl(m, v12, v11)
							mBase = m.M
							v123 = m.ExcPending
							if v123 != 0 {
								return int32(0)
							} else {
								F_line_construct(m, v14, v12, v122)
								mBase = m.M
								v125 = m.ExcPending
								if v125 != 0 {
									return int32(0)
								} else {
									return v14
								}
							}
						} else {
							F_errstart_cold(m, int32(21), int32(0))
							mBase = m.M
							v102 = m.ExcPending
							if v102 != 0 {
								return int32(0)
							} else {
								F_errcode(m, int32(50856066))
								mBase = m.M
								v105 = m.ExcPending
								if v105 != 0 {
									return int32(0)
								} else {
									F_errmsg(m, int32(_a_F_line_construct_pp_0), int32(0))
									mBase = m.M
									v109 = m.ExcPending
									if v109 != 0 {
										return int32(0)
									} else {
										F_errfinish(m, int32(_a_F_line_construct_pp_1), int32(1124), int32(_a_F_line_construct_pp_2))
										mBase = m.M
										v114 = m.ExcPending
										if v114 != 0 {
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
					} else {
						if base.Ui64(v83) < base.Ui64(int64(9218868437227405313)) {
							v122 = F_point_sl(m, v12, v11)
							mBase = m.M
							v123 = m.ExcPending
							if v123 != 0 {
								return int32(0)
							} else {
								F_line_construct(m, v14, v12, v122)
								mBase = m.M
								v125 = m.ExcPending
								if v125 != 0 {
									return int32(0)
								} else {
									return v14
								}
							}
						} else {
							F_errstart_cold(m, int32(21), int32(0))
							mBase = m.M
							v102 = m.ExcPending
							if v102 != 0 {
								return int32(0)
							} else {
								F_errcode(m, int32(50856066))
								mBase = m.M
								v105 = m.ExcPending
								if v105 != 0 {
									return int32(0)
								} else {
									F_errmsg(m, int32(_a_F_line_construct_pp_0), int32(0))
									mBase = m.M
									v109 = m.ExcPending
									if v109 != 0 {
										return int32(0)
									} else {
										F_errfinish(m, int32(_a_F_line_construct_pp_1), int32(1124), int32(_a_F_line_construct_pp_2))
										mBase = m.M
										v114 = m.ExcPending
										if v114 != 0 {
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
				}
			} else {
				if base.Ui64(int64(9218868437227405312)) < base.Ui64(v27) {
					v122 = F_point_sl(m, v12, v11)
					mBase = m.M
					v123 = m.ExcPending
					if v123 != 0 {
						return int32(0)
					} else {
						F_line_construct(m, v14, v12, v122)
						mBase = m.M
						v125 = m.ExcPending
						if v125 != 0 {
							return int32(0)
						} else {
							return v14
						}
					}
				} else {
					v38 = *(*float64)(unsafe.Add(mBase, uint32(v11)+8))
					if base.Ui64(base.I64_reinterpret_f64(v38)&int64(9223372036854775807)) <= base.Ui64(int64(9218868437227405312)) {
						if base.B2i32(base.F64_le(base.F64_abs(base.F64_sub(v18, v24)), float64(1e-06)) == int32(0))&base.F64_ne(v18, v24) != 0 {
							v122 = F_point_sl(m, v12, v11)
							mBase = m.M
							v123 = m.ExcPending
							if v123 != 0 {
								return int32(0)
							} else {
								F_line_construct(m, v14, v12, v122)
								mBase = m.M
								v125 = m.ExcPending
								if v125 != 0 {
									return int32(0)
								} else {
									return v14
								}
							}
						} else {
							if base.F64_eq(v28, v38)|base.F64_le(base.F64_abs(base.F64_sub(v28, v38)), float64(1e-06)) != 0 {
								F_errstart_cold(m, int32(21), int32(0))
								mBase = m.M
								v102 = m.ExcPending
								if v102 != 0 {
									return int32(0)
								} else {
									F_errcode(m, int32(50856066))
									mBase = m.M
									v105 = m.ExcPending
									if v105 != 0 {
										return int32(0)
									} else {
										F_errmsg(m, int32(_a_F_line_construct_pp_0), int32(0))
										mBase = m.M
										v109 = m.ExcPending
										if v109 != 0 {
											return int32(0)
										} else {
											F_errfinish(m, int32(_a_F_line_construct_pp_1), int32(1124), int32(_a_F_line_construct_pp_2))
											mBase = m.M
											v114 = m.ExcPending
											if v114 != 0 {
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
								v122 = F_point_sl(m, v12, v11)
								mBase = m.M
								v123 = m.ExcPending
								if v123 != 0 {
									return int32(0)
								} else {
									F_line_construct(m, v14, v12, v122)
									mBase = m.M
									v125 = m.ExcPending
									if v125 != 0 {
										return int32(0)
									} else {
										return v14
									}
								}
							}
						}
					} else {
						v68 = int32(1)
						if base.B2i32(v68 == int32(0))|base.F64_ne(v18, v24) != 0 {
							v122 = F_point_sl(m, v12, v11)
							mBase = m.M
							v123 = m.ExcPending
							if v123 != 0 {
								return int32(0)
							} else {
								F_line_construct(m, v14, v12, v122)
								mBase = m.M
								v125 = m.ExcPending
								if v125 != 0 {
									return int32(0)
								} else {
									return v14
								}
							}
						} else {
							v75 = v28
							v79 = v31
							v80 = *(*float64)(unsafe.Add(mBase, uint32(v11)+8))
							v83 = base.I64_reinterpret_f64(v80) & int64(9223372036854775807)
							if base.Ui64(v79) <= base.Ui64(int64(9218868437227405312)) {
								if base.B2i32(base.Ui64(int64(9218868437227405313)) <= base.Ui64(v83))|base.F64_ne(v80, v75) != 0 {
									v122 = F_point_sl(m, v12, v11)
									mBase = m.M
									v123 = m.ExcPending
									if v123 != 0 {
										return int32(0)
									} else {
										F_line_construct(m, v14, v12, v122)
										mBase = m.M
										v125 = m.ExcPending
										if v125 != 0 {
											return int32(0)
										} else {
											return v14
										}
									}
								} else {
									F_errstart_cold(m, int32(21), int32(0))
									mBase = m.M
									v102 = m.ExcPending
									if v102 != 0 {
										return int32(0)
									} else {
										F_errcode(m, int32(50856066))
										mBase = m.M
										v105 = m.ExcPending
										if v105 != 0 {
											return int32(0)
										} else {
											F_errmsg(m, int32(_a_F_line_construct_pp_0), int32(0))
											mBase = m.M
											v109 = m.ExcPending
											if v109 != 0 {
												return int32(0)
											} else {
												F_errfinish(m, int32(_a_F_line_construct_pp_1), int32(1124), int32(_a_F_line_construct_pp_2))
												mBase = m.M
												v114 = m.ExcPending
												if v114 != 0 {
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
							} else {
								if base.Ui64(v83) < base.Ui64(int64(9218868437227405313)) {
									v122 = F_point_sl(m, v12, v11)
									mBase = m.M
									v123 = m.ExcPending
									if v123 != 0 {
										return int32(0)
									} else {
										F_line_construct(m, v14, v12, v122)
										mBase = m.M
										v125 = m.ExcPending
										if v125 != 0 {
											return int32(0)
										} else {
											return v14
										}
									}
								} else {
									F_errstart_cold(m, int32(21), int32(0))
									mBase = m.M
									v102 = m.ExcPending
									if v102 != 0 {
										return int32(0)
									} else {
										F_errcode(m, int32(50856066))
										mBase = m.M
										v105 = m.ExcPending
										if v105 != 0 {
											return int32(0)
										} else {
											F_errmsg(m, int32(_a_F_line_construct_pp_0), int32(0))
											mBase = m.M
											v109 = m.ExcPending
											if v109 != 0 {
												return int32(0)
											} else {
												F_errfinish(m, int32(_a_F_line_construct_pp_1), int32(1124), int32(_a_F_line_construct_pp_2))
												mBase = m.M
												v114 = m.ExcPending
												if v114 != 0 {
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
						}
					}
				}
			}
		} else {
			v45 = *(*int64)(unsafe.Add(mBase, uint32(v11)))
			if base.Ui64(v45&int64(9223372036854775807)) < base.Ui64(int64(9218868437227405313)) {
				v122 = F_point_sl(m, v12, v11)
				mBase = m.M
				v123 = m.ExcPending
				if v123 != 0 {
					return int32(0)
				} else {
					F_line_construct(m, v14, v12, v122)
					mBase = m.M
					v125 = m.ExcPending
					if v125 != 0 {
						return int32(0)
					} else {
						return v14
					}
				}
			} else {
				v50 = *(*float64)(unsafe.Add(mBase, uint32(v12)+8))
				v75 = v50
				v79 = base.I64_reinterpret_f64(v50) & int64(9223372036854775807)
				v80 = *(*float64)(unsafe.Add(mBase, uint32(v11)+8))
				v83 = base.I64_reinterpret_f64(v80) & int64(9223372036854775807)
				if base.Ui64(v79) <= base.Ui64(int64(9218868437227405312)) {
					if base.B2i32(base.Ui64(int64(9218868437227405313)) <= base.Ui64(v83))|base.F64_ne(v80, v75) != 0 {
						v122 = F_point_sl(m, v12, v11)
						mBase = m.M
						v123 = m.ExcPending
						if v123 != 0 {
							return int32(0)
						} else {
							F_line_construct(m, v14, v12, v122)
							mBase = m.M
							v125 = m.ExcPending
							if v125 != 0 {
								return int32(0)
							} else {
								return v14
							}
						}
					} else {
						F_errstart_cold(m, int32(21), int32(0))
						mBase = m.M
						v102 = m.ExcPending
						if v102 != 0 {
							return int32(0)
						} else {
							F_errcode(m, int32(50856066))
							mBase = m.M
							v105 = m.ExcPending
							if v105 != 0 {
								return int32(0)
							} else {
								F_errmsg(m, int32(_a_F_line_construct_pp_0), int32(0))
								mBase = m.M
								v109 = m.ExcPending
								if v109 != 0 {
									return int32(0)
								} else {
									F_errfinish(m, int32(_a_F_line_construct_pp_1), int32(1124), int32(_a_F_line_construct_pp_2))
									mBase = m.M
									v114 = m.ExcPending
									if v114 != 0 {
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
				} else {
					if base.Ui64(v83) < base.Ui64(int64(9218868437227405313)) {
						v122 = F_point_sl(m, v12, v11)
						mBase = m.M
						v123 = m.ExcPending
						if v123 != 0 {
							return int32(0)
						} else {
							F_line_construct(m, v14, v12, v122)
							mBase = m.M
							v125 = m.ExcPending
							if v125 != 0 {
								return int32(0)
							} else {
								return v14
							}
						}
					} else {
						F_errstart_cold(m, int32(21), int32(0))
						mBase = m.M
						v102 = m.ExcPending
						if v102 != 0 {
							return int32(0)
						} else {
							F_errcode(m, int32(50856066))
							mBase = m.M
							v105 = m.ExcPending
							if v105 != 0 {
								return int32(0)
							} else {
								F_errmsg(m, int32(_a_F_line_construct_pp_0), int32(0))
								mBase = m.M
								v109 = m.ExcPending
								if v109 != 0 {
									return int32(0)
								} else {
									F_errfinish(m, int32(_a_F_line_construct_pp_1), int32(1124), int32(_a_F_line_construct_pp_2))
									mBase = m.M
									v114 = m.ExcPending
									if v114 != 0 {
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
			}
		}
	}
}
func F_line_horizontal(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v2 int32
	_ = v2
	var v3 float64
	_ = v3
	v2 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v3 = *(*float64)(unsafe.Add(mBase, uint32(v2)))
	return base.F64_le(base.F64_abs(v3), float64(1e-06))
}
func F_line_vertical(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v2 int32
	_ = v2
	var v3 float64
	_ = v3
	v2 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v3 = *(*float64)(unsafe.Add(mBase, uint32(v2)+8))
	return base.F64_le(base.F64_abs(v3), float64(1e-06))
}
