package p3

import (
	base "github.com/shibukawa/pgmem/internal/aot/pgaot/base"
	"unsafe"
)

func F_cfunc_match(m *base.Module, l0 int32, l1 int32, l2 int32) int32 {
	mBase := m.M
	_ = mBase
	var v5 int32
	_ = v5
	var v13 int32
	_ = v13
	var v14 int32
	_ = v14
	var v15 int32
	_ = v15
	var v18 int32
	_ = v18
	var v19 int32
	_ = v19
	var v21 int32
	_ = v21
	var v22 int32
	_ = v22
	var v24 int32
	_ = v24
	var v26 int32
	_ = v26
	var v36 int32
	_ = v36
	var v37 int32
	_ = v37
	var v38 int32
	_ = v38
	var v41 int32
	_ = v41
	var v42 int32
	_ = v42
	var v43 int32
	_ = v43
	var v46 int32
	_ = v46
	var v47 int32
	_ = v47
	var v49 int32
	_ = v49
	var v54 int32
	_ = v54
	var v67 int32
	_ = v67
	var v70 int32
	_ = v70
	var v71 int32
	_ = v71
	var v74 int32
	_ = v74
	var v75 int32
	_ = v75
	var v77 int32
	_ = v77
	var v79 int32
	_ = v79
	var v87 int32
	_ = v87
	var v88 int32
	_ = v88
	var v89 int32
	_ = v89
	var v92 int32
	_ = v92
	var v93 int32
	_ = v93
	var v95 int32
	_ = v95
	var v96 int32
	_ = v96
	var v98 int32
	_ = v98
	var v100 int32
	_ = v100
	var v103 int32
	_ = v103
	var v104 int32
	_ = v104
	var v105 int32
	_ = v105
	var v110 int32
	_ = v110
	var v111 int32
	_ = v111
	var v112 int32
	_ = v112
	var v115 int32
	_ = v115
	var v116 int32
	_ = v116
	var v117 int32
	_ = v117
	var v120 int32
	_ = v120
	var v121 int32
	_ = v121
	var v123 int32
	_ = v123
	var v128 int32
	_ = v128
	var v141 int32
	_ = v141
	var v142 int32
	_ = v142
	var v143 int32
	_ = v143
	var v146 int32
	_ = v146
	var v150 int32
	_ = v150
	var v151 int32
	_ = v151
	var v153 int32
	_ = v153
	var v154 int32
	_ = v154
	var v160 int32
	_ = v160
	var v162 int32
	_ = v162
	var v168 int32
	_ = v168
	var v175 int32
	_ = v175
	var v176 int32
	_ = v176
	var v177 int32
	_ = v177
	var v179 int32
	_ = v179
	var v182 int32
	_ = v182
	var v183 int32
	_ = v183
	var v184 int32
	_ = v184
	var v186 int32
	_ = v186
	var v187 int32
	_ = v187
	var v189 int32
	_ = v189
	var v190 int32
	_ = v190
	var v192 int32
	_ = v192
	var v193 int32
	_ = v193
	var v195 int32
	_ = v195
	var v197 int32
	_ = v197
	var v202 int32
	_ = v202
	var v217 int32
	_ = v217
	v5 = int32(24)
	goto L6
L1:
	;
	return int32(0)
L2:
	;
	return v217
L3:
	;
	if v67 == int32(0) {
		goto L21
	} else {
		goto L22
	}
L4:
	;
	v67 = int32(0)
	goto L3
L5:
	;
	v41 = v36
	v42 = v37
	v43 = v38
	goto L15
L6:
	;
	if (l0|l1)&int32(3) != 0 {
		v36 = l0
		v37 = l1
		v38 = v5
		goto L5
	} else {
		goto L9
	}
L8:
	;
	if v26 == int32(0) {
		goto L4
	} else {
		goto L14
	}
L9:
	;
	v13 = l0
	v14 = l1
	v15 = v5
	goto L10
L10:
	;
	v18 = *(*int32)(unsafe.Add(mBase, uint32(v13)))
	v19 = *(*int32)(unsafe.Add(mBase, uint32(v14)))
	if v18 != v19 {
		v36 = v13
		v37 = v14
		v38 = v15
		goto L5
	} else {
		goto L12
	}
L11:
	;
	goto L8
L12:
	;
	v21 = int32(4)
	v22 = v14 + v21
	v24 = v13 + v21
	v26 = v15 - v21
	if base.Ui32(int32(3)) < base.Ui32(v26) {
		v13 = v24
		v14 = v22
		v15 = v26
		goto L10
	} else {
		goto L13
	}
L13:
	;
	goto L11
L14:
	;
	v36 = v24
	v37 = v22
	v38 = v26
	goto L5
L15:
	;
	v46 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v41))))
	v47 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v42))))
	if v46 == v47 {
		goto L17
	} else {
		goto L18
	}
L16:
	;
	v67 = v46 - v47
	goto L3
L17:
	;
	v49 = int32(1)
	v54 = v43 - v49
	if v54 != 0 {
		v41 = v41 + v49
		v42 = v42 + v49
		v43 = v54
		goto L15
	} else {
		goto L20
	}
L18:
	;
	goto L19
L19:
	;
	goto L16
L20:
	;
	goto L4
L21:
	;
	v70 = int32(1)
	v71 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	if int32(0) < v71 {
		goto L24
	} else {
		goto L25
	}
L22:
	;
	goto L23
L23:
	;
	v217 = int32(1)
	goto L2
L24:
	;
	v74 = int32(28)
	v75 = l0 + v74
	v77 = l1 + v74
	v79 = v71 << (uint(int32(2)) % 32)
	if base.Ui32(int32(4)) <= base.Ui32(v79) {
		goto L30
	} else {
		goto L31
	}
L25:
	;
	goto L26
L26:
	;
	v142 = *(*int32)(unsafe.Add(mBase, uint32(l1)+24))
	v143 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
	if v143 != 0 {
		goto L46
	} else {
		goto L47
	}
L27:
	;
	if v141 != 0 {
		v217 = v70
		goto L2
	} else {
		goto L45
	}
L28:
	;
	v141 = int32(0)
	goto L27
L29:
	;
	v115 = v110
	v116 = v111
	v117 = v112
	goto L39
L30:
	;
	if (v75|v77)&int32(3) != 0 {
		v110 = v75
		v111 = v77
		v112 = v79
		goto L29
	} else {
		goto L33
	}
L31:
	;
	v103 = v75
	v104 = v77
	v105 = v79
	goto L32
L32:
	;
	if v105 == int32(0) {
		goto L28
	} else {
		goto L38
	}
L33:
	;
	v87 = v75
	v88 = v77
	v89 = v79
	goto L34
L34:
	;
	v92 = *(*int32)(unsafe.Add(mBase, uint32(v87)))
	v93 = *(*int32)(unsafe.Add(mBase, uint32(v88)))
	if v92 != v93 {
		v110 = v87
		v111 = v88
		v112 = v89
		goto L29
	} else {
		goto L36
	}
L35:
	;
	v103 = v98
	v104 = v96
	v105 = v100
	goto L32
L36:
	;
	v95 = int32(4)
	v96 = v88 + v95
	v98 = v87 + v95
	v100 = v89 - v95
	if base.Ui32(int32(3)) < base.Ui32(v100) {
		v87 = v98
		v88 = v96
		v89 = v100
		goto L34
	} else {
		goto L37
	}
L37:
	;
	goto L35
L38:
	;
	v110 = v103
	v111 = v104
	v112 = v105
	goto L29
L39:
	;
	v120 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v115))))
	v121 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v116))))
	if v120 == v121 {
		goto L41
	} else {
		goto L42
	}
L40:
	;
	v141 = v120 - v121
	goto L27
L41:
	;
	v123 = int32(1)
	v128 = v117 - v123
	if v128 != 0 {
		v115 = v115 + v123
		v116 = v116 + v123
		v117 = v128
		goto L39
	} else {
		goto L44
	}
L42:
	;
	goto L43
L43:
	;
	goto L40
L44:
	;
	goto L28
L45:
	;
	goto L26
L46:
	;
	if v142 == int32(0) {
		v217 = v70
		goto L2
	} else {
		goto L49
	}
L47:
	;
	goto L48
L48:
	;
	if v142 == int32(0) {
		goto L1
	} else {
		goto L65
	}
L49:
	;
	v146 = int32(0)
	v150 = *(*int32)(unsafe.Add(mBase, uint32(v143)))
	v151 = *(*int32)(unsafe.Add(mBase, uint32(v142)))
	if v150 != v151 {
		v202 = v146
		goto L51
	} else {
		goto L52
	}
L50:
	;
	if v202 == int32(0) {
		v217 = v70
		goto L2
	} else {
		goto L64
	}
L51:
	;
	goto L50
L52:
	;
	v153 = *(*int32)(unsafe.Add(mBase, uint32(v143)+4))
	v154 = *(*int32)(unsafe.Add(mBase, uint32(v142)+4))
	if v153 != v154 {
		v202 = v146
		goto L51
	} else {
		goto L53
	}
L53:
	;
	if v150 <= int32(0) {
		v202 = int32(1)
		goto L51
	} else {
		goto L54
	}
L54:
	;
	v160 = v150 << (uint(int32(4)) % 32)
	v162 = int32(20)
	v168 = int32(0)
	goto L55
L55:
	;
	v175 = v168 * int32(100)
	v176 = v143 + v160 + v162 + v175
	v177 = int32(4)
	v179 = v175 + (v142 + v160 + v162)
	v182 = F_strcmp(m, v176+v177, v179+v177)
	mBase = m.M
	if v182 != 0 {
		goto L57
	} else {
		goto L58
	}
L56:
	;
	v202 = int32(0)
	goto L51
L57:
	;
	goto L56
L58:
	;
	v183 = *(*int32)(unsafe.Add(mBase, uint32(v176)+68))
	v184 = *(*int32)(unsafe.Add(mBase, uint32(v179)+68))
	if v183 != v184 {
		goto L57
	} else {
		goto L59
	}
L59:
	;
	v186 = *(*int32)(unsafe.Add(mBase, uint32(v176)+76))
	v187 = *(*int32)(unsafe.Add(mBase, uint32(v179)+76))
	if v186 != v187 {
		goto L57
	} else {
		goto L60
	}
L60:
	;
	v189 = *(*int32)(unsafe.Add(mBase, uint32(v176)+96))
	v190 = *(*int32)(unsafe.Add(mBase, uint32(v179)+96))
	if v189 != v190 {
		goto L57
	} else {
		goto L61
	}
L61:
	;
	v192 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v176)+91)))
	v193 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v179)+91)))
	if v192 != v193 {
		goto L57
	} else {
		goto L62
	}
L62:
	;
	v195 = int32(1)
	v197 = v168 + v195
	if v150 != v197 {
		v168 = v197
		goto L55
	} else {
		goto L63
	}
L63:
	;
	v202 = v195
	goto L51
L64:
	;
	goto L1
L65:
	;
	goto L23
}
func F_cfunc_resolve_polymorphic_argtypes(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32, l5 int32) {
	mBase := m.M
	_ = mBase
	var v7 int32
	_ = v7
	var v39 int32
	_ = v39
	var v41 int32
	_ = v41
	var v45 int32
	_ = v45
	var v47 int32
	_ = v47
	var v56 int32
	_ = v56
	var v59 int32
	_ = v59
	var v90 int32
	_ = v90
	var v93 int32
	_ = v93
	var v94 int32
	_ = v94
	var v112 int32
	_ = v112
	var v115 int32
	_ = v115
	var v117 int32
	_ = v117
	var v118 int32
	_ = v118
	var v136 int32
	_ = v136
	var v139 int32
	_ = v139
	var v140 int32
	_ = v140
	var v142 int32
	_ = v142
	var v148 int32
	_ = v148
	var v184 int32
	_ = v184
	var v187 int32
	_ = v187
	var v188 int32
	_ = v188
	var v206 int32
	_ = v206
	var v208 int32
	_ = v208
	var v210 int32
	_ = v210
	var v212 int64
	_ = v212
	var v234 int32
	_ = v234
	var v235 int32
	_ = v235
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
	var v247 int32
	_ = v247
	var v248 int32
	_ = v248
	var v249 int32
	_ = v249
	var v251 int32
	_ = v251
	var v252 int32
	_ = v252
	var v253 int32
	_ = v253
	var v254 int32
	_ = v254
	var v255 int32
	_ = v255
	var v256 int32
	_ = v256
	var v257 int32
	_ = v257
	var v258 int32
	_ = v258
	var v260 int32
	_ = v260
	var v263 int32
	_ = v263
	var v265 int32
	_ = v265
	var v268 int32
	_ = v268
	var v269 int32
	_ = v269
	var v278 int32
	_ = v278
	var v288 int32
	_ = v288
	var v292 int32
	_ = v292
	var v293 int32
	_ = v293
	var v295 int32
	_ = v295
	var v299 int32
	_ = v299
	var v300 int32
	_ = v300
	var v302 int32
	_ = v302
	var v303 int32
	_ = v303
	var v305 int32
	_ = v305
	var v309 int32
	_ = v309
	var v310 int32
	_ = v310
	var v312 int32
	_ = v312
	var v316 int32
	_ = v316
	var v317 int32
	_ = v317
	var v319 int32
	_ = v319
	var v323 int32
	_ = v323
	var v324 int32
	_ = v324
	var v326 int32
	_ = v326
	var v330 int32
	_ = v330
	var v331 int32
	_ = v331
	var v333 int32
	_ = v333
	var v337 int32
	_ = v337
	var v338 int32
	_ = v338
	var v340 int32
	_ = v340
	var v349 int32
	_ = v349
	var v350 int32
	_ = v350
	var v351 int32
	_ = v351
	var v352 int32
	_ = v352
	var v353 int32
	_ = v353
	var v354 int32
	_ = v354
	var v355 int32
	_ = v355
	var v356 int32
	_ = v356
	var v367 int32
	_ = v367
	var v368 int32
	_ = v368
	var v369 int32
	_ = v369
	var v370 int32
	_ = v370
	var v371 int32
	_ = v371
	var v372 int32
	_ = v372
	var v373 int32
	_ = v373
	var v374 int32
	_ = v374
	var v375 int32
	_ = v375
	var v376 int32
	_ = v376
	var v377 int32
	_ = v377
	var v378 int32
	_ = v378
	var v379 int32
	_ = v379
	var v380 int32
	_ = v380
	var v381 int32
	_ = v381
	var v382 int32
	_ = v382
	var v383 int32
	_ = v383
	var v384 int32
	_ = v384
	var v385 int32
	_ = v385
	var v386 int32
	_ = v386
	var v387 int32
	_ = v387
	var v388 int32
	_ = v388
	var v389 int32
	_ = v389
	var v390 int32
	_ = v390
	var v391 int32
	_ = v391
	var v396 int32
	_ = v396
	var v398 int32
	_ = v398
	var v408 int32
	_ = v408
	var v419 int32
	_ = v419
	var v420 int32
	_ = v420
	var v421 int32
	_ = v421
	var v428 int32
	_ = v428
	var v429 int32
	_ = v429
	var v436 int32
	_ = v436
	var v437 int32
	_ = v437
	var v444 int32
	_ = v444
	var v449 int32
	_ = v449
	var v450 int32
	_ = v450
	var v451 int32
	_ = v451
	var v456 int32
	_ = v456
	var v457 int32
	_ = v457
	var v462 int32
	_ = v462
	var v463 int32
	_ = v463
	var v468 int32
	_ = v468
	var v469 int32
	_ = v469
	var v471 int32
	_ = v471
	var v472 int32
	_ = v472
	var v473 int32
	_ = v473
	var v474 int32
	_ = v474
	var v475 int32
	_ = v475
	var v476 int32
	_ = v476
	var v477 int32
	_ = v477
	var v478 int32
	_ = v478
	var v479 int32
	_ = v479
	var v496 int32
	_ = v496
	var v504 int32
	_ = v504
	var v525 int32
	_ = v525
	var v526 int32
	_ = v526
	var v541 int32
	_ = v541
	var v545 int32
	_ = v545
	var v546 int32
	_ = v546
	var v561 int32
	_ = v561
	var v564 int32
	_ = v564
	var v565 int32
	_ = v565
	var v567 int32
	_ = v567
	var v580 int32
	_ = v580
	var v611 int32
	_ = v611
	var v612 int32
	_ = v612
	var v627 int32
	_ = v627
	var v639 int32
	_ = v639
	var v674 int32
	_ = v674
	var v680 int32
	_ = v680
	var v681 int32
	_ = v681
	var v717 int32
	_ = v717
	var v722 int32
	_ = v722
	var v723 int32
	_ = v723
	var v729 int32
	_ = v729
	var v730 int32
	_ = v730
	var v737 int32
	_ = v737
	var v741 int32
	_ = v741
	var v787 int32
	_ = v787
	var v790 int32
	_ = v790
	var v794 int32
	_ = v794
	var v799 int32
	_ = v799
	v7 = int32(0)
	v39 = m.G0
	v41 = v39 - int32(16)
	m.G0 = v41
	if l4 != 0 {
		goto L3
	} else {
		goto L4
	}
L1:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v787 = m.ExcPending
	if v787 != 0 {
		goto L83
	} else {
		goto L232
	}
L2:
	;
	m.G0 = v41 + int32(16)
	return
L3:
	;
	if l0 <= int32(0) {
		goto L2
	} else {
		goto L6
	}
L4:
	;
	goto L5
L5:
	;
	v208 = m.G0
	v210 = v208 - int32(32)
	m.G0 = v210
	v212 = int64(0)
	*(*int64)(unsafe.Add(mBase, uint32(v210)+24)) = v212
	*(*int64)(unsafe.Add(mBase, uint32(v210)+8)) = v212
	*(*int64)(unsafe.Add(mBase, uint32(v210)+16)) = v212
	*(*int64)(unsafe.Add(mBase, uint32(v210))) = v212
	if l0 <= int32(0) {
		v639 = int32(1)
		goto L52
	} else {
		goto L53
	}
L6:
	;
	v45 = int32(1)
	v47 = int32(0)
	if l0 != v45 {
		goto L7
	} else {
		goto L8
	}
L7:
	;
	v56 = v47
	v59 = v7
	goto L10
L8:
	;
	v148 = v47
	goto L9
L9:
	;
	if l0&v45 == int32(0) {
		goto L2
	} else {
		goto L39
	}
L10:
	;
	v90 = int32(23)
	v93 = l1 + v56<<(uint(int32(2))%32)
	v94 = *(*int32)(unsafe.Add(mBase, uint32(v93)))
	if v94 <= int32(3830) {
		goto L16
	} else {
		goto L17
	}
L11:
	;
	v148 = v140
	goto L9
L12:
	;
	v115 = int32(23)
	v117 = v93 + int32(4)
	v118 = *(*int32)(unsafe.Add(mBase, uint32(v117)))
	if v118 <= int32(3830) {
		goto L28
	} else {
		goto L29
	}
L13:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v93))) = v112
	goto L12
L14:
	;
	v112 = int32(3904)
	goto L13
L15:
	;
	v112 = int32(1007)
	goto L13
L16:
	;
	switch v94 - int32(2277) {
	case 0:
		goto L15
	case 1, 2, 3, 4, 5:
		goto L12
	case 6:
		v112 = v90
		goto L13
	default:
		goto L19
	}
L17:
	;
	goto L18
L18:
	;
	switch v94 - int32(_a_F_cfunc_resolve_polymorphic_argtypes_0) {
	case 0, 2:
		v112 = v90
		goto L13
	case 1:
		goto L15
	case 3:
		goto L14
	default:
		goto L22
	}
L19:
	;
	if v94 == int32(2776) {
		v112 = v90
		goto L13
	} else {
		goto L20
	}
L20:
	;
	if v94 == int32(3500) {
		v112 = v90
		goto L13
	} else {
		goto L21
	}
L21:
	;
	goto L12
L22:
	;
	if v94 == int32(3831) {
		goto L14
	} else {
		goto L23
	}
L23:
	;
	if v94 != int32(_a_F_cfunc_resolve_polymorphic_argtypes_1) {
		goto L12
	} else {
		goto L24
	}
L24:
	;
	v112 = int32(_a_F_cfunc_resolve_polymorphic_argtypes_2)
	goto L13
L25:
	;
	v139 = int32(2)
	v140 = v56 + v139
	v142 = v59 + v139
	if v142 != l0&int32(2147483646) {
		v56 = v140
		v59 = v142
		goto L10
	} else {
		goto L38
	}
L26:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v117))) = v136
	goto L25
L27:
	;
	v136 = int32(1007)
	goto L26
L28:
	;
	switch v118 - int32(2277) {
	case 0:
		goto L27
	case 1, 2, 3, 4, 5:
		goto L25
	case 6:
		v136 = v115
		goto L26
	default:
		goto L31
	}
L29:
	;
	goto L30
L30:
	;
	switch v118 - int32(_a_F_cfunc_resolve_polymorphic_argtypes_0) {
	case 0, 2:
		v136 = v115
		goto L26
	case 1:
		goto L27
	case 3:
		goto L34
	default:
		goto L35
	}
L31:
	;
	if v118 == int32(2776) {
		v136 = v115
		goto L26
	} else {
		goto L32
	}
L32:
	;
	if v118 == int32(3500) {
		v136 = v115
		goto L26
	} else {
		goto L33
	}
L33:
	;
	goto L25
L34:
	;
	v136 = int32(3904)
	goto L26
L35:
	;
	if v118 == int32(3831) {
		goto L34
	} else {
		goto L36
	}
L36:
	;
	if v118 != int32(_a_F_cfunc_resolve_polymorphic_argtypes_1) {
		goto L25
	} else {
		goto L37
	}
L37:
	;
	v136 = int32(_a_F_cfunc_resolve_polymorphic_argtypes_2)
	goto L26
L38:
	;
	goto L11
L39:
	;
	v184 = int32(23)
	v187 = l1 + v148<<(uint(int32(2))%32)
	v188 = *(*int32)(unsafe.Add(mBase, uint32(v187)))
	if v188 <= int32(3830) {
		goto L42
	} else {
		goto L43
	}
L40:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v187))) = v206
	goto L2
L41:
	;
	v206 = int32(1007)
	goto L40
L42:
	;
	switch v188 - int32(2277) {
	case 0:
		goto L41
	case 1, 2, 3, 4, 5:
		goto L2
	case 6:
		v206 = v184
		goto L40
	default:
		goto L45
	}
L43:
	;
	goto L44
L44:
	;
	switch v188 - int32(_a_F_cfunc_resolve_polymorphic_argtypes_0) {
	case 0, 2:
		v206 = v184
		goto L40
	case 1:
		goto L41
	case 3:
		goto L48
	default:
		goto L49
	}
L45:
	;
	if v188 == int32(2776) {
		v206 = v184
		goto L40
	} else {
		goto L46
	}
L46:
	;
	if v188 == int32(3500) {
		v206 = v184
		goto L40
	} else {
		goto L47
	}
L47:
	;
	goto L2
L48:
	;
	v206 = int32(3904)
	goto L40
L49:
	;
	if v188 == int32(3831) {
		goto L48
	} else {
		goto L50
	}
L50:
	;
	if v188 != int32(_a_F_cfunc_resolve_polymorphic_argtypes_1) {
		goto L2
	} else {
		goto L51
	}
L51:
	;
	v206 = int32(_a_F_cfunc_resolve_polymorphic_argtypes_2)
	goto L40
L52:
	;
	m.G0 = v210 + int32(32)
	if v639 == int32(0) {
		goto L1
	} else {
		goto L220
	}
L53:
	;
	v234 = int32(0)
	v235 = v7
	v240 = v7
	v241 = v7
	v242 = v7
	v245 = v7
	v246 = v7
	v247 = v7
	v248 = v7
	v249 = v7
	v251 = v7
	v252 = v7
	v253 = v7
	v254 = v7
	v255 = v7
	v256 = v7
	v257 = v7
	v258 = v7
	v260 = v7
	goto L54
L54:
	;
	if l2 != 0 {
		goto L56
	} else {
		goto L57
	}
L55:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v210)+8)) = v390
	*(*int32)(unsafe.Add(mBase, uint32(v210)+12)) = v391
	*(*int32)(unsafe.Add(mBase, uint32(v210)+4)) = v385
	*(*int32)(unsafe.Add(mBase, uint32(v210)+28)) = v389
	*(*int32)(unsafe.Add(mBase, uint32(v210)+24)) = v388
	*(*int32)(unsafe.Add(mBase, uint32(v210)+20)) = v384
	*(*int32)(unsafe.Add(mBase, uint32(v210))) = v386
	*(*int32)(unsafe.Add(mBase, uint32(v210)+16)) = v387
	v408 = int32(1)
	if v375&v408 == int32(0) {
		v639 = v408
		goto L52
	} else {
		goto L130
	}
L56:
	;
	v263 = int32(*(*int8)(unsafe.Add(mBase, uint32(l2+v235))))
	v265 = v263
	goto L58
L57:
	;
	v265 = int32(105)
	goto L58
L58:
	;
	v268 = l1 + v235<<(uint(int32(2))%32)
	v269 = *(*int32)(unsafe.Add(mBase, uint32(v268)))
	if v269 <= int32(3830) {
		goto L70
	} else {
		goto L71
	}
L59:
	;
	switch v265 - int32(111) {
	case 0, 5:
		v396 = v248
		goto L127
	default:
		goto L128
	}
L60:
	;
	v375 = v260
	v376 = v247
	v377 = v234
	v378 = v249
	v379 = v245
	v380 = v241
	v381 = v240
	v382 = v246
	v383 = v242
	v384 = v367
	v385 = v368
	v386 = v369
	v387 = v370
	v388 = v371
	v389 = v372
	v390 = v373
	v391 = v374
	goto L59
L61:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v268))) = v340
	v367 = v349
	v368 = v350
	v369 = v351
	v370 = v352
	v371 = v353
	v372 = v354
	v373 = v355
	v374 = v356
	goto L60
L62:
	;
	v333 = int32(1)
	switch v265 - int32(111) {
	case 0, 5:
		v375 = v333
		v376 = v247
		v377 = v234
		v378 = v249
		v379 = v245
		v380 = v241
		v381 = v333
		v382 = v246
		v383 = v242
		v384 = v251
		v385 = v252
		v386 = v253
		v387 = v254
		v388 = v255
		v389 = v256
		v390 = v257
		v391 = v258
		goto L59
	default:
		goto L121
	}
L63:
	;
	v326 = int32(1)
	switch v265 - int32(111) {
	case 0, 5:
		v375 = v326
		v376 = v247
		v377 = v234
		v378 = v249
		v379 = v245
		v380 = v326
		v381 = v240
		v382 = v246
		v383 = v242
		v384 = v251
		v385 = v252
		v386 = v253
		v387 = v254
		v388 = v255
		v389 = v256
		v390 = v257
		v391 = v258
		goto L59
	default:
		goto L115
	}
L64:
	;
	v319 = int32(1)
	switch v265 - int32(111) {
	case 0, 5:
		v375 = v319
		v376 = v247
		v377 = v234
		v378 = v249
		v379 = v319
		v380 = v241
		v381 = v240
		v382 = v246
		v383 = v242
		v384 = v251
		v385 = v252
		v386 = v253
		v387 = v254
		v388 = v255
		v389 = v256
		v390 = v257
		v391 = v258
		goto L59
	default:
		goto L109
	}
L65:
	;
	v312 = int32(1)
	switch v265 - int32(111) {
	case 0, 5:
		v375 = v312
		v376 = v247
		v377 = v234
		v378 = v312
		v379 = v245
		v380 = v241
		v381 = v240
		v382 = v246
		v383 = v242
		v384 = v251
		v385 = v252
		v386 = v253
		v387 = v254
		v388 = v255
		v389 = v256
		v390 = v257
		v391 = v258
		goto L59
	default:
		goto L103
	}
L66:
	;
	v305 = int32(1)
	switch v265 - int32(111) {
	case 0, 5:
		v375 = v305
		v376 = v247
		v377 = v234
		v378 = v249
		v379 = v245
		v380 = v241
		v381 = v240
		v382 = v246
		v383 = v305
		v384 = v251
		v385 = v252
		v386 = v253
		v387 = v254
		v388 = v255
		v389 = v256
		v390 = v257
		v391 = v258
		goto L59
	default:
		goto L97
	}
L67:
	;
	if v255 != 0 {
		goto L92
	} else {
		goto L93
	}
L68:
	;
	v295 = int32(1)
	switch v265 - int32(111) {
	case 0, 5:
		v375 = v295
		v376 = v247
		v377 = v234
		v378 = v249
		v379 = v245
		v380 = v241
		v381 = v240
		v382 = v295
		v383 = v242
		v384 = v251
		v385 = v252
		v386 = v253
		v387 = v254
		v388 = v255
		v389 = v256
		v390 = v257
		v391 = v258
		goto L59
	default:
		goto L86
	}
L69:
	;
	if v254 != 0 {
		goto L80
	} else {
		goto L81
	}
L70:
	;
	switch v269 - int32(2277) {
	case 0:
		goto L68
	case 1, 2, 3, 4, 5:
		v375 = v260
		v376 = v247
		v377 = v234
		v378 = v249
		v379 = v245
		v380 = v241
		v381 = v240
		v382 = v246
		v383 = v242
		v384 = v251
		v385 = v252
		v386 = v253
		v387 = v254
		v388 = v255
		v389 = v256
		v390 = v257
		v391 = v258
		goto L59
	case 6:
		goto L73
	default:
		goto L74
	}
L71:
	;
	goto L72
L72:
	;
	switch v269 - int32(_a_F_cfunc_resolve_polymorphic_argtypes_0) {
	case 0, 2:
		goto L65
	case 1:
		goto L64
	case 3:
		goto L63
	default:
		goto L77
	}
L73:
	;
	v278 = int32(1)
	switch v265 - int32(111) {
	case 0, 5:
		v375 = v278
		v376 = v278
		v377 = v234
		v378 = v249
		v379 = v245
		v380 = v241
		v381 = v240
		v382 = v246
		v383 = v242
		v384 = v251
		v385 = v252
		v386 = v253
		v387 = v254
		v388 = v255
		v389 = v256
		v390 = v257
		v391 = v258
		goto L59
	default:
		goto L69
	}
L74:
	;
	if v269 == int32(2776) {
		goto L73
	} else {
		goto L75
	}
L75:
	;
	if v269 != int32(3500) {
		v367 = v251
		v368 = v252
		v369 = v253
		v370 = v254
		v371 = v255
		v372 = v256
		v373 = v257
		v374 = v258
		goto L60
	} else {
		goto L76
	}
L76:
	;
	goto L73
L77:
	;
	switch v269 - int32(_a_F_cfunc_resolve_polymorphic_argtypes_1) {
	case 0:
		goto L66
	case 1:
		goto L62
	default:
		goto L78
	}
L78:
	;
	if v269 != int32(3831) {
		v367 = v251
		v368 = v252
		v369 = v253
		v370 = v254
		v371 = v255
		v372 = v256
		v373 = v257
		v374 = v258
		goto L60
	} else {
		goto L79
	}
L79:
	;
	v288 = int32(1)
	switch v265 - int32(111) {
	case 0, 5:
		v375 = v288
		v376 = v247
		v377 = v288
		v378 = v249
		v379 = v245
		v380 = v241
		v381 = v240
		v382 = v246
		v383 = v242
		v384 = v251
		v385 = v252
		v386 = v253
		v387 = v254
		v388 = v255
		v389 = v256
		v390 = v257
		v391 = v258
		goto L59
	default:
		goto L67
	}
L80:
	;
	v340 = v254
	v349 = v251
	v350 = v252
	v351 = v253
	v352 = v254
	v353 = v255
	v354 = v256
	v355 = v257
	v356 = v258
	goto L61
L81:
	;
	goto L82
L82:
	;
	v292 = F_get_call_expr_argtype(m, l3, v248)
	mBase = m.M
	v293 = m.ExcPending
	if v293 != 0 {
		goto L83
	} else {
		goto L84
	}
L83:
	;
	return
L84:
	;
	if v292 != 0 {
		v340 = v292
		v349 = v251
		v350 = v252
		v351 = v253
		v352 = v292
		v353 = v255
		v354 = v256
		v355 = v257
		v356 = v258
		goto L61
	} else {
		goto L85
	}
L85:
	;
	v639 = int32(0)
	goto L52
L86:
	;
	if v251 != 0 {
		goto L87
	} else {
		goto L88
	}
L87:
	;
	v340 = v251
	v349 = v251
	v350 = v252
	v351 = v253
	v352 = v254
	v353 = v255
	v354 = v256
	v355 = v257
	v356 = v258
	goto L61
L88:
	;
	goto L89
L89:
	;
	v299 = F_get_call_expr_argtype(m, l3, v248)
	mBase = m.M
	v300 = m.ExcPending
	if v300 != 0 {
		goto L83
	} else {
		goto L90
	}
L90:
	;
	if v299 != 0 {
		v340 = v299
		v349 = v299
		v350 = v252
		v351 = v253
		v352 = v254
		v353 = v255
		v354 = v256
		v355 = v257
		v356 = v258
		goto L61
	} else {
		goto L91
	}
L91:
	;
	v639 = int32(0)
	goto L52
L92:
	;
	v340 = v255
	v349 = v251
	v350 = v252
	v351 = v253
	v352 = v254
	v353 = v255
	v354 = v256
	v355 = v257
	v356 = v258
	goto L61
L93:
	;
	goto L94
L94:
	;
	v302 = F_get_call_expr_argtype(m, l3, v248)
	mBase = m.M
	v303 = m.ExcPending
	if v303 != 0 {
		goto L83
	} else {
		goto L95
	}
L95:
	;
	if v302 != 0 {
		v340 = v302
		v349 = v251
		v350 = v252
		v351 = v253
		v352 = v254
		v353 = v302
		v354 = v256
		v355 = v257
		v356 = v258
		goto L61
	} else {
		goto L96
	}
L96:
	;
	v639 = int32(0)
	goto L52
L97:
	;
	if v256 != 0 {
		goto L98
	} else {
		goto L99
	}
L98:
	;
	v340 = v256
	v349 = v251
	v350 = v252
	v351 = v253
	v352 = v254
	v353 = v255
	v354 = v256
	v355 = v257
	v356 = v258
	goto L61
L99:
	;
	goto L100
L100:
	;
	v309 = F_get_call_expr_argtype(m, l3, v248)
	mBase = m.M
	v310 = m.ExcPending
	if v310 != 0 {
		goto L83
	} else {
		goto L101
	}
L101:
	;
	if v309 != 0 {
		v340 = v309
		v349 = v251
		v350 = v252
		v351 = v253
		v352 = v254
		v353 = v255
		v354 = v309
		v355 = v257
		v356 = v258
		goto L61
	} else {
		goto L102
	}
L102:
	;
	v639 = int32(0)
	goto L52
L103:
	;
	if v253 != 0 {
		goto L104
	} else {
		goto L105
	}
L104:
	;
	v340 = v253
	v349 = v251
	v350 = v252
	v351 = v253
	v352 = v254
	v353 = v255
	v354 = v256
	v355 = v257
	v356 = v258
	goto L61
L105:
	;
	goto L106
L106:
	;
	v316 = F_get_call_expr_argtype(m, l3, v248)
	mBase = m.M
	v317 = m.ExcPending
	if v317 != 0 {
		goto L83
	} else {
		goto L107
	}
L107:
	;
	if v316 != 0 {
		v340 = v316
		v349 = v251
		v350 = v252
		v351 = v316
		v352 = v254
		v353 = v255
		v354 = v256
		v355 = v257
		v356 = v258
		goto L61
	} else {
		goto L108
	}
L108:
	;
	v639 = int32(0)
	goto L52
L109:
	;
	if v252 != 0 {
		goto L110
	} else {
		goto L111
	}
L110:
	;
	v340 = v252
	v349 = v251
	v350 = v252
	v351 = v253
	v352 = v254
	v353 = v255
	v354 = v256
	v355 = v257
	v356 = v258
	goto L61
L111:
	;
	goto L112
L112:
	;
	v323 = F_get_call_expr_argtype(m, l3, v248)
	mBase = m.M
	v324 = m.ExcPending
	if v324 != 0 {
		goto L83
	} else {
		goto L113
	}
L113:
	;
	if v323 != 0 {
		v340 = v323
		v349 = v251
		v350 = v323
		v351 = v253
		v352 = v254
		v353 = v255
		v354 = v256
		v355 = v257
		v356 = v258
		goto L61
	} else {
		goto L114
	}
L114:
	;
	v639 = int32(0)
	goto L52
L115:
	;
	if v257 != 0 {
		goto L116
	} else {
		goto L117
	}
L116:
	;
	v340 = v257
	v349 = v251
	v350 = v252
	v351 = v253
	v352 = v254
	v353 = v255
	v354 = v256
	v355 = v257
	v356 = v258
	goto L61
L117:
	;
	goto L118
L118:
	;
	v330 = F_get_call_expr_argtype(m, l3, v248)
	mBase = m.M
	v331 = m.ExcPending
	if v331 != 0 {
		goto L83
	} else {
		goto L119
	}
L119:
	;
	if v330 != 0 {
		v340 = v330
		v349 = v251
		v350 = v252
		v351 = v253
		v352 = v254
		v353 = v255
		v354 = v256
		v355 = v330
		v356 = v258
		goto L61
	} else {
		goto L120
	}
L120:
	;
	v639 = int32(0)
	goto L52
L121:
	;
	if v258 != 0 {
		goto L122
	} else {
		goto L123
	}
L122:
	;
	v340 = v258
	v349 = v251
	v350 = v252
	v351 = v253
	v352 = v254
	v353 = v255
	v354 = v256
	v355 = v257
	v356 = v258
	goto L61
L123:
	;
	goto L124
L124:
	;
	v337 = F_get_call_expr_argtype(m, l3, v248)
	mBase = m.M
	v338 = m.ExcPending
	if v338 != 0 {
		goto L83
	} else {
		goto L125
	}
L125:
	;
	if v337 != 0 {
		v340 = v337
		v349 = v251
		v350 = v252
		v351 = v253
		v352 = v254
		v353 = v255
		v354 = v256
		v355 = v257
		v356 = v337
		goto L61
	} else {
		goto L126
	}
L126:
	;
	v639 = int32(0)
	goto L52
L127:
	;
	v398 = v235 + int32(1)
	if v398 != l0 {
		v234 = v377
		v235 = v398
		v240 = v381
		v241 = v380
		v242 = v383
		v245 = v379
		v246 = v382
		v247 = v376
		v248 = v396
		v249 = v378
		v251 = v384
		v252 = v385
		v253 = v386
		v254 = v387
		v255 = v388
		v256 = v389
		v257 = v390
		v258 = v391
		v260 = v375
		goto L54
	} else {
		goto L129
	}
L128:
	;
	v396 = v248 + int32(1)
	goto L127
L129:
	;
	goto L55
L130:
	;
	if base.B2i32(v387 == int32(0))&v376 != 0 {
		goto L131
	} else {
		goto L132
	}
L131:
	;
	F_resolve_anyelement_from_others(m, v210+int32(16))
	mBase = m.M
	v419 = m.ExcPending
	if v419 != 0 {
		goto L83
	} else {
		goto L134
	}
L132:
	;
	v421 = v384
	goto L133
L133:
	;
	if base.B2i32(v421 == int32(0))&v382 != 0 {
		goto L135
	} else {
		goto L136
	}
L134:
	;
	v420 = *(*int32)(unsafe.Add(mBase, uint32(v210)+20))
	v421 = v420
	goto L133
L135:
	;
	F_resolve_anyarray_from_others(m, v210+int32(16))
	mBase = m.M
	v428 = m.ExcPending
	if v428 != 0 {
		goto L83
	} else {
		goto L138
	}
L136:
	;
	goto L137
L137:
	;
	v429 = *(*int32)(unsafe.Add(mBase, uint32(v210)+24))
	if base.B2i32(v429 == int32(0))&v377 != 0 {
		goto L139
	} else {
		goto L140
	}
L138:
	;
	goto L137
L139:
	;
	F_resolve_anyrange_from_others(m, v210+int32(16))
	mBase = m.M
	v436 = m.ExcPending
	if v436 != 0 {
		goto L83
	} else {
		goto L142
	}
L140:
	;
	goto L141
L141:
	;
	v437 = *(*int32)(unsafe.Add(mBase, uint32(v210)+28))
	if base.B2i32(v437 == int32(0))&v383 != 0 {
		goto L143
	} else {
		goto L144
	}
L142:
	;
	goto L141
L143:
	;
	F_resolve_anymultirange_from_others(m, v210+int32(16))
	mBase = m.M
	v444 = m.ExcPending
	if v444 != 0 {
		goto L83
	} else {
		goto L146
	}
L144:
	;
	goto L145
L145:
	;
	if base.B2i32(v386 == int32(0))&v378 != 0 {
		goto L147
	} else {
		goto L148
	}
L146:
	;
	goto L145
L147:
	;
	F_resolve_anyelement_from_others(m, v210)
	mBase = m.M
	v449 = m.ExcPending
	if v449 != 0 {
		goto L83
	} else {
		goto L150
	}
L148:
	;
	v451 = v385
	goto L149
L149:
	;
	if base.B2i32(v451 == int32(0))&v379 != 0 {
		goto L151
	} else {
		goto L152
	}
L150:
	;
	v450 = *(*int32)(unsafe.Add(mBase, uint32(v210)+4))
	v451 = v450
	goto L149
L151:
	;
	F_resolve_anyarray_from_others(m, v210)
	mBase = m.M
	v456 = m.ExcPending
	if v456 != 0 {
		goto L83
	} else {
		goto L154
	}
L152:
	;
	goto L153
L153:
	;
	v457 = *(*int32)(unsafe.Add(mBase, uint32(v210)+8))
	if base.B2i32(v457 == int32(0))&v380 != 0 {
		goto L155
	} else {
		goto L156
	}
L154:
	;
	goto L153
L155:
	;
	F_resolve_anyrange_from_others(m, v210)
	mBase = m.M
	v462 = m.ExcPending
	if v462 != 0 {
		goto L83
	} else {
		goto L158
	}
L156:
	;
	goto L157
L157:
	;
	v463 = *(*int32)(unsafe.Add(mBase, uint32(v210)+12))
	if base.B2i32(v463 == int32(0))&v381 != 0 {
		goto L159
	} else {
		goto L160
	}
L158:
	;
	goto L157
L159:
	;
	F_resolve_anymultirange_from_others(m, v210)
	mBase = m.M
	v468 = m.ExcPending
	if v468 != 0 {
		goto L83
	} else {
		goto L162
	}
L160:
	;
	goto L161
L161:
	;
	v469 = int32(1)
	v471 = int32(0)
	v472 = *(*int32)(unsafe.Add(mBase, uint32(v210)+16))
	v473 = *(*int32)(unsafe.Add(mBase, uint32(v210)+20))
	v474 = *(*int32)(unsafe.Add(mBase, uint32(v210)+24))
	v475 = *(*int32)(unsafe.Add(mBase, uint32(v210)+28))
	v476 = *(*int32)(unsafe.Add(mBase, uint32(v210)))
	v477 = *(*int32)(unsafe.Add(mBase, uint32(v210)+4))
	v478 = *(*int32)(unsafe.Add(mBase, uint32(v210)+8))
	v479 = *(*int32)(unsafe.Add(mBase, uint32(v210)+12))
	if l0 != v469 {
		goto L163
	} else {
		goto L164
	}
L162:
	;
	goto L161
L163:
	;
	v496 = v471
	v504 = int32(0)
	goto L166
L164:
	;
	v580 = v471
	goto L165
L165:
	;
	if l0&v469 == int32(0) {
		v639 = v408
		goto L52
	} else {
		goto L203
	}
L166:
	;
	v525 = l1 + v496<<(uint(int32(2))%32)
	v526 = *(*int32)(unsafe.Add(mBase, uint32(v525)))
	if v526 <= int32(3830) {
		goto L176
	} else {
		goto L177
	}
L167:
	;
	v580 = v565
	goto L165
L168:
	;
	v545 = v525 + int32(4)
	v546 = *(*int32)(unsafe.Add(mBase, uint32(v545)))
	if v546 <= int32(3830) {
		goto L188
	} else {
		goto L189
	}
L169:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v525))) = v541
	goto L168
L170:
	;
	v541 = v479
	goto L169
L171:
	;
	v541 = v478
	goto L169
L172:
	;
	v541 = v477
	goto L169
L173:
	;
	v541 = v476
	goto L169
L174:
	;
	v541 = v475
	goto L169
L175:
	;
	v541 = v473
	goto L169
L176:
	;
	switch v526 - int32(2277) {
	case 0:
		goto L175
	case 1, 2, 3, 4, 5:
		goto L168
	case 6:
		v541 = v472
		goto L169
	default:
		goto L179
	}
L177:
	;
	goto L178
L178:
	;
	switch v526 - int32(_a_F_cfunc_resolve_polymorphic_argtypes_0) {
	case 0, 2:
		goto L173
	case 1:
		goto L172
	case 3:
		goto L171
	default:
		goto L182
	}
L179:
	;
	if v526 == int32(2776) {
		v541 = v472
		goto L169
	} else {
		goto L180
	}
L180:
	;
	if v526 == int32(3500) {
		v541 = v472
		goto L169
	} else {
		goto L181
	}
L181:
	;
	goto L168
L182:
	;
	switch v526 - int32(_a_F_cfunc_resolve_polymorphic_argtypes_1) {
	case 0:
		goto L174
	case 1:
		goto L170
	default:
		goto L183
	}
L183:
	;
	if v526 == int32(3831) {
		v541 = v474
		goto L169
	} else {
		goto L184
	}
L184:
	;
	goto L168
L185:
	;
	v564 = int32(2)
	v565 = v496 + v564
	v567 = v504 + v564
	if v567 != l0&int32(2147483646) {
		v496 = v565
		v504 = v567
		goto L166
	} else {
		goto L202
	}
L186:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v545))) = v561
	goto L185
L187:
	;
	v561 = v473
	goto L186
L188:
	;
	switch v546 - int32(2277) {
	case 0:
		goto L187
	case 1, 2, 3, 4, 5:
		goto L185
	case 6:
		v561 = v472
		goto L186
	default:
		goto L191
	}
L189:
	;
	goto L190
L190:
	;
	switch v546 - int32(_a_F_cfunc_resolve_polymorphic_argtypes_0) {
	case 0, 2:
		goto L195
	case 1:
		goto L196
	case 3:
		goto L197
	default:
		goto L198
	}
L191:
	;
	if v546 == int32(2776) {
		v561 = v472
		goto L186
	} else {
		goto L192
	}
L192:
	;
	if v546 == int32(3500) {
		v561 = v472
		goto L186
	} else {
		goto L193
	}
L193:
	;
	goto L185
L194:
	;
	v561 = v475
	goto L186
L195:
	;
	v561 = v476
	goto L186
L196:
	;
	v561 = v477
	goto L186
L197:
	;
	v561 = v478
	goto L186
L198:
	;
	switch v546 - int32(_a_F_cfunc_resolve_polymorphic_argtypes_1) {
	case 0:
		goto L194
	case 1:
		goto L199
	default:
		goto L200
	}
L199:
	;
	v561 = v479
	goto L186
L200:
	;
	if v546 == int32(3831) {
		v561 = v474
		goto L186
	} else {
		goto L201
	}
L201:
	;
	goto L185
L202:
	;
	goto L167
L203:
	;
	v611 = l1 + v580<<(uint(int32(2))%32)
	v612 = *(*int32)(unsafe.Add(mBase, uint32(v611)))
	if v612 <= int32(3830) {
		goto L206
	} else {
		goto L207
	}
L204:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v611))) = v627
	v639 = v408
	goto L52
L205:
	;
	v627 = v473
	goto L204
L206:
	;
	switch v612 - int32(2277) {
	case 0:
		goto L205
	case 1, 2, 3, 4, 5:
		v639 = v408
		goto L52
	case 6:
		v627 = v472
		goto L204
	default:
		goto L209
	}
L207:
	;
	goto L208
L208:
	;
	switch v612 - int32(_a_F_cfunc_resolve_polymorphic_argtypes_0) {
	case 0, 2:
		goto L213
	case 1:
		goto L214
	case 3:
		goto L215
	default:
		goto L216
	}
L209:
	;
	if v612 == int32(2776) {
		v627 = v472
		goto L204
	} else {
		goto L210
	}
L210:
	;
	if v612 == int32(3500) {
		v627 = v472
		goto L204
	} else {
		goto L211
	}
L211:
	;
	v639 = v408
	goto L52
L212:
	;
	v627 = v475
	goto L204
L213:
	;
	v627 = v476
	goto L204
L214:
	;
	v627 = v477
	goto L204
L215:
	;
	v627 = v478
	goto L204
L216:
	;
	switch v612 - int32(_a_F_cfunc_resolve_polymorphic_argtypes_1) {
	case 0:
		goto L212
	case 1:
		goto L217
	default:
		goto L218
	}
L217:
	;
	v627 = v479
	goto L204
L218:
	;
	if v612 == int32(3831) {
		v627 = v474
		goto L204
	} else {
		goto L219
	}
L219:
	;
	v639 = v408
	goto L52
L220:
	;
	if l0 <= int32(0) {
		goto L2
	} else {
		goto L221
	}
L221:
	;
	v674 = int32(0)
	v680 = v674
	v681 = v674
	goto L222
L222:
	;
	if l2 == int32(0) {
		goto L225
	} else {
		goto L226
	}
L223:
	;
	goto L2
L224:
	;
	v741 = v680 + int32(1)
	if v741 != l0 {
		v680 = v741
		v681 = v737
		goto L222
	} else {
		goto L231
	}
L225:
	;
	v722 = l1 + v680<<(uint(int32(2))%32)
	v723 = *(*int32)(unsafe.Add(mBase, uint32(v722)))
	if base.B2i32(v723 != int32(2287))&base.B2i32(v723 != int32(2249)) != 0 {
		goto L227
	} else {
		goto L228
	}
L226:
	;
	v717 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l2+v680))))
	switch v717 - int32(111) {
	case 0, 5:
		v737 = v681
		goto L224
	default:
		goto L225
	}
L227:
	;
	v737 = v681 + int32(1)
	goto L224
L228:
	;
	v729 = F_get_call_expr_argtype(m, l3, v681)
	mBase = m.M
	v730 = m.ExcPending
	if v730 != 0 {
		goto L83
	} else {
		goto L229
	}
L229:
	;
	if v729 == int32(0) {
		goto L227
	} else {
		goto L230
	}
L230:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v722))) = v729
	goto L227
L231:
	;
	goto L223
L232:
	;
	F_errcode(m, int32(1088))
	mBase = m.M
	v790 = m.ExcPending
	if v790 != 0 {
		goto L83
	} else {
		goto L233
	}
L233:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v41))) = l5
	F_errmsg(m, int32(_a_F_cfunc_resolve_polymorphic_argtypes_3), v41)
	mBase = m.M
	v794 = m.ExcPending
	if v794 != 0 {
		goto L83
	} else {
		goto L234
	}
L234:
	;
	F_errfinish(m, int32(_a_F_cfunc_resolve_polymorphic_argtypes_4), int32(366), int32(_a_F_cfunc_resolve_polymorphic_argtypes_5))
	mBase = m.M
	v799 = m.ExcPending
	if v799 != 0 {
		goto L83
	} else {
		goto L235
	}
L235:
	;
	base.Wasm_trap_unreachable()
	for {
	}
}
