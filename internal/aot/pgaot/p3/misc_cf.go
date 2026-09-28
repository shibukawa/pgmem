package p3

import (
	base "github.com/shibukawa/pgmem/internal/aot/pgaot/base"
	"unsafe"
)

func F_cfunc_match(m *base.Module, l0 int32, l1 int32, l2 int32) int32 {
	mBase := m.M
	_ = mBase
	var v4 int64
	_ = v4
	var v5 int64
	_ = v5
	var v7 int64
	_ = v7
	var v8 int64
	_ = v8
	var v11 int64
	_ = v11
	var v12 int64
	_ = v12
	var v17 int32
	_ = v17
	var v20 int32
	_ = v20
	var v21 int32
	_ = v21
	var v23 int32
	_ = v23
	var v25 int32
	_ = v25
	var v33 int32
	_ = v33
	var v34 int32
	_ = v34
	var v35 int32
	_ = v35
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
	var v46 int32
	_ = v46
	var v49 int32
	_ = v49
	var v50 int32
	_ = v50
	var v51 int32
	_ = v51
	var v56 int32
	_ = v56
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
	var v66 int32
	_ = v66
	var v67 int32
	_ = v67
	var v69 int32
	_ = v69
	var v74 int32
	_ = v74
	var v87 int32
	_ = v87
	var v88 int32
	_ = v88
	var v89 int32
	_ = v89
	var v92 int32
	_ = v92
	var v96 int32
	_ = v96
	var v97 int32
	_ = v97
	var v99 int32
	_ = v99
	var v100 int32
	_ = v100
	var v106 int32
	_ = v106
	var v108 int32
	_ = v108
	var v114 int32
	_ = v114
	var v121 int32
	_ = v121
	var v122 int32
	_ = v122
	var v123 int32
	_ = v123
	var v125 int32
	_ = v125
	var v128 int32
	_ = v128
	var v129 int32
	_ = v129
	var v130 int32
	_ = v130
	var v132 int32
	_ = v132
	var v133 int32
	_ = v133
	var v135 int32
	_ = v135
	var v136 int32
	_ = v136
	var v138 int32
	_ = v138
	var v139 int32
	_ = v139
	var v141 int32
	_ = v141
	var v143 int32
	_ = v143
	var v148 int32
	_ = v148
	v4 = *(*int64)(unsafe.Add(mBase, uint32(l0)))
	v5 = *(*int64)(unsafe.Add(mBase, uint32(l1)))
	v7 = *(*int64)(unsafe.Add(mBase, uint32(l0)+8))
	v8 = *(*int64)(unsafe.Add(mBase, uint32(l1)+8))
	v11 = *(*int64)(unsafe.Add(mBase, uint32(l0)+16))
	v12 = *(*int64)(unsafe.Add(mBase, uint32(l1)+16))
	if v4^v5|(v7^v8)|(v11^v12) == int64(0) {
		goto L2
	} else {
		goto L3
	}
L1:
	;
	return int32(0)
L2:
	;
	v17 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	if int32(0) < v17 {
		goto L6
	} else {
		goto L7
	}
L3:
	;
	goto L4
L4:
	;
	return int32(1)
L5:
	;
	goto L4
L6:
	;
	v20 = int32(28)
	v21 = l0 + v20
	v23 = l1 + v20
	v25 = v17 << (uint(int32(2)) % 32)
	if base.Ui32(int32(4)) <= base.Ui32(v25) {
		goto L12
	} else {
		goto L13
	}
L7:
	;
	goto L8
L8:
	;
	v88 = *(*int32)(unsafe.Add(mBase, uint32(l1)+24))
	v89 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
	if v89 != 0 {
		goto L28
	} else {
		goto L29
	}
L9:
	;
	if v87 != 0 {
		goto L5
	} else {
		goto L27
	}
L10:
	;
	v87 = int32(0)
	goto L9
L11:
	;
	v61 = v56
	v62 = v57
	v63 = v58
	goto L21
L12:
	;
	if (v21|v23)&int32(3) != 0 {
		v56 = v21
		v57 = v23
		v58 = v25
		goto L11
	} else {
		goto L15
	}
L13:
	;
	v49 = v21
	v50 = v23
	v51 = v25
	goto L14
L14:
	;
	if v51 == int32(0) {
		goto L10
	} else {
		goto L20
	}
L15:
	;
	v33 = v21
	v34 = v23
	v35 = v25
	goto L16
L16:
	;
	v38 = *(*int32)(unsafe.Add(mBase, uint32(v33)))
	v39 = *(*int32)(unsafe.Add(mBase, uint32(v34)))
	if v38 != v39 {
		v56 = v33
		v57 = v34
		v58 = v35
		goto L11
	} else {
		goto L18
	}
L17:
	;
	v49 = v44
	v50 = v42
	v51 = v46
	goto L14
L18:
	;
	v41 = int32(4)
	v42 = v34 + v41
	v44 = v33 + v41
	v46 = v35 - v41
	if base.Ui32(int32(3)) < base.Ui32(v46) {
		v33 = v44
		v34 = v42
		v35 = v46
		goto L16
	} else {
		goto L19
	}
L19:
	;
	goto L17
L20:
	;
	v56 = v49
	v57 = v50
	v58 = v51
	goto L11
L21:
	;
	v66 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v61))))
	v67 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v62))))
	if v66 == v67 {
		goto L23
	} else {
		goto L24
	}
L22:
	;
	v87 = v66 - v67
	goto L9
L23:
	;
	v69 = int32(1)
	v74 = v63 - v69
	if v74 != 0 {
		v61 = v61 + v69
		v62 = v62 + v69
		v63 = v74
		goto L21
	} else {
		goto L26
	}
L24:
	;
	goto L25
L25:
	;
	goto L22
L26:
	;
	goto L10
L27:
	;
	goto L8
L28:
	;
	if v88 == int32(0) {
		goto L5
	} else {
		goto L31
	}
L29:
	;
	goto L30
L30:
	;
	if v88 == int32(0) {
		goto L1
	} else {
		goto L47
	}
L31:
	;
	v92 = int32(0)
	v96 = *(*int32)(unsafe.Add(mBase, uint32(v89)))
	v97 = *(*int32)(unsafe.Add(mBase, uint32(v88)))
	if v96 != v97 {
		v148 = v92
		goto L33
	} else {
		goto L34
	}
L32:
	;
	if v148 == int32(0) {
		goto L5
	} else {
		goto L46
	}
L33:
	;
	goto L32
L34:
	;
	v99 = *(*int32)(unsafe.Add(mBase, uint32(v89)+4))
	v100 = *(*int32)(unsafe.Add(mBase, uint32(v88)+4))
	if v99 != v100 {
		v148 = v92
		goto L33
	} else {
		goto L35
	}
L35:
	;
	if v96 <= int32(0) {
		v148 = int32(1)
		goto L33
	} else {
		goto L36
	}
L36:
	;
	v106 = v96 << (uint(int32(3)) % 32)
	v108 = int32(28)
	v114 = int32(0)
	goto L37
L37:
	;
	v121 = v114 * int32(100)
	v122 = v89 + v106 + v108 + v121
	v123 = int32(4)
	v125 = v121 + (v88 + v106 + v108)
	v128 = F_strcmp(m, v122+v123, v125+v123)
	mBase = m.M
	if v128 != 0 {
		goto L39
	} else {
		goto L40
	}
L38:
	;
	v148 = int32(0)
	goto L33
L39:
	;
	goto L38
L40:
	;
	v129 = *(*int32)(unsafe.Add(mBase, uint32(v122)+68))
	v130 = *(*int32)(unsafe.Add(mBase, uint32(v125)+68))
	if v129 != v130 {
		goto L39
	} else {
		goto L41
	}
L41:
	;
	v132 = *(*int32)(unsafe.Add(mBase, uint32(v122)+76))
	v133 = *(*int32)(unsafe.Add(mBase, uint32(v125)+76))
	if v132 != v133 {
		goto L39
	} else {
		goto L42
	}
L42:
	;
	v135 = *(*int32)(unsafe.Add(mBase, uint32(v122)+96))
	v136 = *(*int32)(unsafe.Add(mBase, uint32(v125)+96))
	if v135 != v136 {
		goto L39
	} else {
		goto L43
	}
L43:
	;
	v138 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v122)+91)))
	v139 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v125)+91)))
	if v138 != v139 {
		goto L39
	} else {
		goto L44
	}
L44:
	;
	v141 = int32(1)
	v143 = v114 + v141
	if v96 != v143 {
		v114 = v143
		goto L37
	} else {
		goto L45
	}
L45:
	;
	v148 = v141
	goto L33
L46:
	;
	goto L1
L47:
	;
	goto L5
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
	var v54 int32
	_ = v54
	var v63 int32
	_ = v63
	var v90 int32
	_ = v90
	var v93 int32
	_ = v93
	var v94 int32
	_ = v94
	var v113 int32
	_ = v113
	var v116 int32
	_ = v116
	var v117 int32
	_ = v117
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
	var v207 int32
	_ = v207
	var v209 int32
	_ = v209
	var v211 int32
	_ = v211
	var v213 int64
	_ = v213
	var v235 int32
	_ = v235
	var v236 int32
	_ = v236
	var v238 int32
	_ = v238
	var v239 int32
	_ = v239
	var v241 int32
	_ = v241
	var v245 int32
	_ = v245
	var v246 int32
	_ = v246
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
	var v256 int32
	_ = v256
	var v257 int32
	_ = v257
	var v258 int32
	_ = v258
	var v259 int32
	_ = v259
	var v260 int32
	_ = v260
	var v261 int32
	_ = v261
	var v264 int32
	_ = v264
	var v266 int32
	_ = v266
	var v269 int32
	_ = v269
	var v270 int32
	_ = v270
	var v279 int32
	_ = v279
	var v289 int32
	_ = v289
	var v293 int32
	_ = v293
	var v294 int32
	_ = v294
	var v296 int32
	_ = v296
	var v300 int32
	_ = v300
	var v301 int32
	_ = v301
	var v303 int32
	_ = v303
	var v304 int32
	_ = v304
	var v306 int32
	_ = v306
	var v310 int32
	_ = v310
	var v311 int32
	_ = v311
	var v313 int32
	_ = v313
	var v317 int32
	_ = v317
	var v318 int32
	_ = v318
	var v320 int32
	_ = v320
	var v324 int32
	_ = v324
	var v325 int32
	_ = v325
	var v327 int32
	_ = v327
	var v331 int32
	_ = v331
	var v332 int32
	_ = v332
	var v334 int32
	_ = v334
	var v338 int32
	_ = v338
	var v339 int32
	_ = v339
	var v342 int32
	_ = v342
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
	var v357 int32
	_ = v357
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
	var v392 int32
	_ = v392
	var v397 int32
	_ = v397
	var v399 int32
	_ = v399
	var v409 int32
	_ = v409
	var v420 int32
	_ = v420
	var v421 int32
	_ = v421
	var v422 int32
	_ = v422
	var v429 int32
	_ = v429
	var v430 int32
	_ = v430
	var v437 int32
	_ = v437
	var v438 int32
	_ = v438
	var v445 int32
	_ = v445
	var v450 int32
	_ = v450
	var v451 int32
	_ = v451
	var v452 int32
	_ = v452
	var v457 int32
	_ = v457
	var v458 int32
	_ = v458
	var v463 int32
	_ = v463
	var v464 int32
	_ = v464
	var v469 int32
	_ = v469
	var v470 int32
	_ = v470
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
	var v498 int32
	_ = v498
	var v500 int32
	_ = v500
	var v526 int32
	_ = v526
	var v527 int32
	_ = v527
	var v543 int32
	_ = v543
	var v546 int32
	_ = v546
	var v562 int32
	_ = v562
	var v565 int32
	_ = v565
	var v566 int32
	_ = v566
	var v568 int32
	_ = v568
	var v584 int32
	_ = v584
	var v610 int32
	_ = v610
	var v611 int32
	_ = v611
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
		goto L80
	} else {
		goto L225
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
	v209 = m.G0
	v211 = v209 - int32(32)
	m.G0 = v211
	v213 = int64(0)
	*(*int64)(unsafe.Add(mBase, uint32(v211)+16)) = v213
	*(*int64)(unsafe.Add(mBase, uint32(v211)+24)) = v213
	*(*int64)(unsafe.Add(mBase, uint32(v211))) = v213
	*(*int64)(unsafe.Add(mBase, uint32(v211)+8)) = v213
	if l0 <= int32(0) {
		v639 = int32(1)
		goto L49
	} else {
		goto L50
	}
L6:
	;
	v45 = int32(0)
	if l0 != int32(1) {
		goto L7
	} else {
		goto L8
	}
L7:
	;
	v54 = v45
	v63 = v7
	goto L10
L8:
	;
	v148 = v45
	goto L9
L9:
	;
	v184 = int32(23)
	v187 = l1 + v148<<(uint(int32(2))%32)
	v188 = *(*int32)(unsafe.Add(mBase, uint32(v187)))
	if v188 <= int32(3830) {
		goto L40
	} else {
		goto L41
	}
L10:
	;
	v90 = int32(23)
	v93 = l1 + v54<<(uint(int32(2))%32)
	v94 = *(*int32)(unsafe.Add(mBase, uint32(v93)))
	if v94 <= int32(3830) {
		goto L16
	} else {
		goto L17
	}
L11:
	;
	if l0&int32(1) == int32(0) {
		goto L2
	} else {
		goto L37
	}
L12:
	;
	v116 = int32(23)
	v117 = *(*int32)(unsafe.Add(mBase, uint32(v93)+4))
	if v117 <= int32(3830) {
		goto L27
	} else {
		goto L28
	}
L13:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v93))) = v113
	goto L12
L14:
	;
	v113 = int32(3904)
	goto L13
L15:
	;
	v113 = int32(1007)
	goto L13
L16:
	;
	switch v94 - int32(2277) {
	case 0:
		goto L15
	case 1, 2, 3, 4, 5:
		goto L12
	case 6:
		v113 = v90
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
		v113 = v90
		goto L13
	case 1:
		goto L15
	case 3:
		goto L14
	default:
		goto L21
	}
L19:
	;
	if base.B2i32(v94 == int32(2776))|base.B2i32(v94 == int32(3500)) != 0 {
		v113 = v90
		goto L13
	} else {
		goto L20
	}
L20:
	;
	goto L12
L21:
	;
	if v94 == int32(3831) {
		goto L14
	} else {
		goto L22
	}
L22:
	;
	if v94 != int32(_a_F_cfunc_resolve_polymorphic_argtypes_1) {
		goto L12
	} else {
		goto L23
	}
L23:
	;
	v113 = int32(_a_F_cfunc_resolve_polymorphic_argtypes_2)
	goto L13
L24:
	;
	v139 = int32(2)
	v140 = v54 + v139
	v142 = v63 + v139
	if v142 != l0&int32(2147483646) {
		v54 = v140
		v63 = v142
		goto L10
	} else {
		goto L36
	}
L25:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v93)+4)) = v136
	goto L24
L26:
	;
	v136 = int32(1007)
	goto L25
L27:
	;
	switch v117 - int32(2277) {
	case 0:
		goto L26
	case 1, 2, 3, 4, 5:
		goto L24
	case 6:
		v136 = v116
		goto L25
	default:
		goto L30
	}
L28:
	;
	goto L29
L29:
	;
	switch v117 - int32(_a_F_cfunc_resolve_polymorphic_argtypes_0) {
	case 0, 2:
		v136 = v116
		goto L25
	case 1:
		goto L26
	case 3:
		goto L32
	default:
		goto L33
	}
L30:
	;
	if base.B2i32(v117 == int32(2776))|base.B2i32(v117 == int32(3500)) != 0 {
		v136 = v116
		goto L25
	} else {
		goto L31
	}
L31:
	;
	goto L24
L32:
	;
	v136 = int32(3904)
	goto L25
L33:
	;
	if v117 == int32(3831) {
		goto L32
	} else {
		goto L34
	}
L34:
	;
	if v117 != int32(_a_F_cfunc_resolve_polymorphic_argtypes_1) {
		goto L24
	} else {
		goto L35
	}
L35:
	;
	v136 = int32(_a_F_cfunc_resolve_polymorphic_argtypes_2)
	goto L25
L36:
	;
	goto L11
L37:
	;
	v148 = v140
	goto L9
L38:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v187))) = v207
	goto L2
L39:
	;
	v207 = int32(1007)
	goto L38
L40:
	;
	switch v188 - int32(2277) {
	case 0:
		goto L39
	case 1, 2, 3, 4, 5:
		goto L2
	case 6:
		v207 = v184
		goto L38
	default:
		goto L43
	}
L41:
	;
	goto L42
L42:
	;
	switch v188 - int32(_a_F_cfunc_resolve_polymorphic_argtypes_0) {
	case 0, 2:
		v207 = v184
		goto L38
	case 1:
		goto L39
	case 3:
		goto L45
	default:
		goto L46
	}
L43:
	;
	if base.B2i32(v188 == int32(2776))|base.B2i32(v188 == int32(3500)) != 0 {
		v207 = v184
		goto L38
	} else {
		goto L44
	}
L44:
	;
	goto L2
L45:
	;
	v207 = int32(3904)
	goto L38
L46:
	;
	if v188 == int32(3831) {
		goto L45
	} else {
		goto L47
	}
L47:
	;
	if v188 != int32(_a_F_cfunc_resolve_polymorphic_argtypes_1) {
		goto L2
	} else {
		goto L48
	}
L48:
	;
	v207 = int32(_a_F_cfunc_resolve_polymorphic_argtypes_2)
	goto L38
L49:
	;
	m.G0 = v211 + int32(32)
	if v639 == int32(0) {
		goto L1
	} else {
		goto L213
	}
L50:
	;
	v235 = int32(0)
	v236 = v7
	v238 = v7
	v239 = v7
	v241 = v7
	v245 = v7
	v246 = v7
	v248 = v7
	v249 = v7
	v251 = v7
	v252 = v7
	v253 = v7
	v254 = v7
	v256 = v7
	v257 = v7
	v258 = v7
	v259 = v7
	v260 = v7
	v261 = v7
	goto L51
L51:
	;
	if l2 != 0 {
		goto L53
	} else {
		goto L54
	}
L52:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v211)+8)) = v391
	*(*int32)(unsafe.Add(mBase, uint32(v211)+12)) = v392
	*(*int32)(unsafe.Add(mBase, uint32(v211)+4)) = v388
	*(*int32)(unsafe.Add(mBase, uint32(v211)+28)) = v390
	*(*int32)(unsafe.Add(mBase, uint32(v211)+24)) = v389
	*(*int32)(unsafe.Add(mBase, uint32(v211)+20)) = v386
	*(*int32)(unsafe.Add(mBase, uint32(v211))) = v387
	*(*int32)(unsafe.Add(mBase, uint32(v211)+16)) = v385
	v409 = int32(1)
	if v377&v409 == int32(0) {
		v639 = v409
		goto L49
	} else {
		goto L127
	}
L53:
	;
	v264 = int32(*(*int8)(unsafe.Add(mBase, uint32(l2+v256))))
	v266 = v264
	goto L55
L54:
	;
	v266 = int32(105)
	goto L55
L55:
	;
	v269 = l1 + v256<<(uint(int32(2))%32)
	v270 = *(*int32)(unsafe.Add(mBase, uint32(v269)))
	if v270 <= int32(3830) {
		goto L67
	} else {
		goto L68
	}
L56:
	;
	switch v266 - int32(111) {
	case 0, 5:
		v397 = v249
		goto L124
	default:
		goto L125
	}
L57:
	;
	v376 = v239
	v377 = v261
	v378 = v235
	v379 = v248
	v380 = v238
	v381 = v236
	v382 = v241
	v383 = v245
	v384 = v246
	v385 = v368
	v386 = v369
	v387 = v370
	v388 = v371
	v389 = v372
	v390 = v373
	v391 = v374
	v392 = v375
	goto L56
L58:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v269))) = v342
	v368 = v350
	v369 = v351
	v370 = v352
	v371 = v353
	v372 = v354
	v373 = v355
	v374 = v356
	v375 = v357
	goto L57
L59:
	;
	v334 = int32(1)
	switch v266 - int32(111) {
	case 0, 5:
		v376 = v239
		v377 = v334
		v378 = v235
		v379 = v334
		v380 = v238
		v381 = v236
		v382 = v241
		v383 = v245
		v384 = v246
		v385 = v251
		v386 = v252
		v387 = v253
		v388 = v254
		v389 = v257
		v390 = v258
		v391 = v259
		v392 = v260
		goto L56
	default:
		goto L118
	}
L60:
	;
	v327 = int32(1)
	switch v266 - int32(111) {
	case 0, 5:
		v376 = v239
		v377 = v327
		v378 = v235
		v379 = v248
		v380 = v238
		v381 = v236
		v382 = v241
		v383 = v245
		v384 = v327
		v385 = v251
		v386 = v252
		v387 = v253
		v388 = v254
		v389 = v257
		v390 = v258
		v391 = v259
		v392 = v260
		goto L56
	default:
		goto L112
	}
L61:
	;
	v320 = int32(1)
	switch v266 - int32(111) {
	case 0, 5:
		v376 = v239
		v377 = v320
		v378 = v235
		v379 = v248
		v380 = v238
		v381 = v236
		v382 = v241
		v383 = v320
		v384 = v246
		v385 = v251
		v386 = v252
		v387 = v253
		v388 = v254
		v389 = v257
		v390 = v258
		v391 = v259
		v392 = v260
		goto L56
	default:
		goto L106
	}
L62:
	;
	v313 = int32(1)
	switch v266 - int32(111) {
	case 0, 5:
		v376 = v239
		v377 = v313
		v378 = v235
		v379 = v248
		v380 = v238
		v381 = v236
		v382 = v313
		v383 = v245
		v384 = v246
		v385 = v251
		v386 = v252
		v387 = v253
		v388 = v254
		v389 = v257
		v390 = v258
		v391 = v259
		v392 = v260
		goto L56
	default:
		goto L100
	}
L63:
	;
	v306 = int32(1)
	switch v266 - int32(111) {
	case 0, 5:
		v376 = v239
		v377 = v306
		v378 = v235
		v379 = v248
		v380 = v238
		v381 = v306
		v382 = v241
		v383 = v245
		v384 = v246
		v385 = v251
		v386 = v252
		v387 = v253
		v388 = v254
		v389 = v257
		v390 = v258
		v391 = v259
		v392 = v260
		goto L56
	default:
		goto L94
	}
L64:
	;
	if v257 != 0 {
		goto L89
	} else {
		goto L90
	}
L65:
	;
	v296 = int32(1)
	switch v266 - int32(111) {
	case 0, 5:
		v376 = v239
		v377 = v296
		v378 = v235
		v379 = v248
		v380 = v296
		v381 = v236
		v382 = v241
		v383 = v245
		v384 = v246
		v385 = v251
		v386 = v252
		v387 = v253
		v388 = v254
		v389 = v257
		v390 = v258
		v391 = v259
		v392 = v260
		goto L56
	default:
		goto L83
	}
L66:
	;
	if v251 != 0 {
		goto L77
	} else {
		goto L78
	}
L67:
	;
	switch v270 - int32(2277) {
	case 0:
		goto L65
	case 1, 2, 3, 4, 5:
		v376 = v239
		v377 = v261
		v378 = v235
		v379 = v248
		v380 = v238
		v381 = v236
		v382 = v241
		v383 = v245
		v384 = v246
		v385 = v251
		v386 = v252
		v387 = v253
		v388 = v254
		v389 = v257
		v390 = v258
		v391 = v259
		v392 = v260
		goto L56
	case 6:
		goto L70
	default:
		goto L71
	}
L68:
	;
	goto L69
L69:
	;
	switch v270 - int32(_a_F_cfunc_resolve_polymorphic_argtypes_0) {
	case 0, 2:
		goto L62
	case 1:
		goto L61
	case 3:
		goto L60
	default:
		goto L74
	}
L70:
	;
	v279 = int32(1)
	switch v266 - int32(111) {
	case 0, 5:
		v376 = v279
		v377 = v279
		v378 = v235
		v379 = v248
		v380 = v238
		v381 = v236
		v382 = v241
		v383 = v245
		v384 = v246
		v385 = v251
		v386 = v252
		v387 = v253
		v388 = v254
		v389 = v257
		v390 = v258
		v391 = v259
		v392 = v260
		goto L56
	default:
		goto L66
	}
L71:
	;
	if v270 == int32(2776) {
		goto L70
	} else {
		goto L72
	}
L72:
	;
	if v270 != int32(3500) {
		v368 = v251
		v369 = v252
		v370 = v253
		v371 = v254
		v372 = v257
		v373 = v258
		v374 = v259
		v375 = v260
		goto L57
	} else {
		goto L73
	}
L73:
	;
	goto L70
L74:
	;
	switch v270 - int32(_a_F_cfunc_resolve_polymorphic_argtypes_1) {
	case 0:
		goto L63
	case 1:
		goto L59
	default:
		goto L75
	}
L75:
	;
	if v270 != int32(3831) {
		v368 = v251
		v369 = v252
		v370 = v253
		v371 = v254
		v372 = v257
		v373 = v258
		v374 = v259
		v375 = v260
		goto L57
	} else {
		goto L76
	}
L76:
	;
	v289 = int32(1)
	switch v266 - int32(111) {
	case 0, 5:
		v376 = v239
		v377 = v289
		v378 = v289
		v379 = v248
		v380 = v238
		v381 = v236
		v382 = v241
		v383 = v245
		v384 = v246
		v385 = v251
		v386 = v252
		v387 = v253
		v388 = v254
		v389 = v257
		v390 = v258
		v391 = v259
		v392 = v260
		goto L56
	default:
		goto L64
	}
L77:
	;
	v342 = v251
	v350 = v251
	v351 = v252
	v352 = v253
	v353 = v254
	v354 = v257
	v355 = v258
	v356 = v259
	v357 = v260
	goto L58
L78:
	;
	goto L79
L79:
	;
	v293 = F_get_call_expr_argtype(m, l3, v249)
	mBase = m.M
	v294 = m.ExcPending
	if v294 != 0 {
		goto L80
	} else {
		goto L81
	}
L80:
	;
	return
L81:
	;
	if v293 != 0 {
		v342 = v293
		v350 = v293
		v351 = v252
		v352 = v253
		v353 = v254
		v354 = v257
		v355 = v258
		v356 = v259
		v357 = v260
		goto L58
	} else {
		goto L82
	}
L82:
	;
	v639 = int32(0)
	goto L49
L83:
	;
	if v252 != 0 {
		goto L84
	} else {
		goto L85
	}
L84:
	;
	v342 = v252
	v350 = v251
	v351 = v252
	v352 = v253
	v353 = v254
	v354 = v257
	v355 = v258
	v356 = v259
	v357 = v260
	goto L58
L85:
	;
	goto L86
L86:
	;
	v300 = F_get_call_expr_argtype(m, l3, v249)
	mBase = m.M
	v301 = m.ExcPending
	if v301 != 0 {
		goto L80
	} else {
		goto L87
	}
L87:
	;
	if v300 != 0 {
		v342 = v300
		v350 = v251
		v351 = v300
		v352 = v253
		v353 = v254
		v354 = v257
		v355 = v258
		v356 = v259
		v357 = v260
		goto L58
	} else {
		goto L88
	}
L88:
	;
	v639 = int32(0)
	goto L49
L89:
	;
	v342 = v257
	v350 = v251
	v351 = v252
	v352 = v253
	v353 = v254
	v354 = v257
	v355 = v258
	v356 = v259
	v357 = v260
	goto L58
L90:
	;
	goto L91
L91:
	;
	v303 = F_get_call_expr_argtype(m, l3, v249)
	mBase = m.M
	v304 = m.ExcPending
	if v304 != 0 {
		goto L80
	} else {
		goto L92
	}
L92:
	;
	if v303 != 0 {
		v342 = v303
		v350 = v251
		v351 = v252
		v352 = v253
		v353 = v254
		v354 = v303
		v355 = v258
		v356 = v259
		v357 = v260
		goto L58
	} else {
		goto L93
	}
L93:
	;
	v639 = int32(0)
	goto L49
L94:
	;
	if v258 != 0 {
		goto L95
	} else {
		goto L96
	}
L95:
	;
	v342 = v258
	v350 = v251
	v351 = v252
	v352 = v253
	v353 = v254
	v354 = v257
	v355 = v258
	v356 = v259
	v357 = v260
	goto L58
L96:
	;
	goto L97
L97:
	;
	v310 = F_get_call_expr_argtype(m, l3, v249)
	mBase = m.M
	v311 = m.ExcPending
	if v311 != 0 {
		goto L80
	} else {
		goto L98
	}
L98:
	;
	if v310 != 0 {
		v342 = v310
		v350 = v251
		v351 = v252
		v352 = v253
		v353 = v254
		v354 = v257
		v355 = v310
		v356 = v259
		v357 = v260
		goto L58
	} else {
		goto L99
	}
L99:
	;
	v639 = int32(0)
	goto L49
L100:
	;
	if v253 != 0 {
		goto L101
	} else {
		goto L102
	}
L101:
	;
	v342 = v253
	v350 = v251
	v351 = v252
	v352 = v253
	v353 = v254
	v354 = v257
	v355 = v258
	v356 = v259
	v357 = v260
	goto L58
L102:
	;
	goto L103
L103:
	;
	v317 = F_get_call_expr_argtype(m, l3, v249)
	mBase = m.M
	v318 = m.ExcPending
	if v318 != 0 {
		goto L80
	} else {
		goto L104
	}
L104:
	;
	if v317 != 0 {
		v342 = v317
		v350 = v251
		v351 = v252
		v352 = v317
		v353 = v254
		v354 = v257
		v355 = v258
		v356 = v259
		v357 = v260
		goto L58
	} else {
		goto L105
	}
L105:
	;
	v639 = int32(0)
	goto L49
L106:
	;
	if v254 != 0 {
		goto L107
	} else {
		goto L108
	}
L107:
	;
	v342 = v254
	v350 = v251
	v351 = v252
	v352 = v253
	v353 = v254
	v354 = v257
	v355 = v258
	v356 = v259
	v357 = v260
	goto L58
L108:
	;
	goto L109
L109:
	;
	v324 = F_get_call_expr_argtype(m, l3, v249)
	mBase = m.M
	v325 = m.ExcPending
	if v325 != 0 {
		goto L80
	} else {
		goto L110
	}
L110:
	;
	if v324 != 0 {
		v342 = v324
		v350 = v251
		v351 = v252
		v352 = v253
		v353 = v324
		v354 = v257
		v355 = v258
		v356 = v259
		v357 = v260
		goto L58
	} else {
		goto L111
	}
L111:
	;
	v639 = int32(0)
	goto L49
L112:
	;
	if v259 != 0 {
		goto L113
	} else {
		goto L114
	}
L113:
	;
	v342 = v259
	v350 = v251
	v351 = v252
	v352 = v253
	v353 = v254
	v354 = v257
	v355 = v258
	v356 = v259
	v357 = v260
	goto L58
L114:
	;
	goto L115
L115:
	;
	v331 = F_get_call_expr_argtype(m, l3, v249)
	mBase = m.M
	v332 = m.ExcPending
	if v332 != 0 {
		goto L80
	} else {
		goto L116
	}
L116:
	;
	if v331 != 0 {
		v342 = v331
		v350 = v251
		v351 = v252
		v352 = v253
		v353 = v254
		v354 = v257
		v355 = v258
		v356 = v331
		v357 = v260
		goto L58
	} else {
		goto L117
	}
L117:
	;
	v639 = int32(0)
	goto L49
L118:
	;
	if v260 != 0 {
		goto L119
	} else {
		goto L120
	}
L119:
	;
	v342 = v260
	v350 = v251
	v351 = v252
	v352 = v253
	v353 = v254
	v354 = v257
	v355 = v258
	v356 = v259
	v357 = v260
	goto L58
L120:
	;
	goto L121
L121:
	;
	v338 = F_get_call_expr_argtype(m, l3, v249)
	mBase = m.M
	v339 = m.ExcPending
	if v339 != 0 {
		goto L80
	} else {
		goto L122
	}
L122:
	;
	if v338 != 0 {
		v342 = v338
		v350 = v251
		v351 = v252
		v352 = v253
		v353 = v254
		v354 = v257
		v355 = v258
		v356 = v259
		v357 = v338
		goto L58
	} else {
		goto L123
	}
L123:
	;
	v639 = int32(0)
	goto L49
L124:
	;
	v399 = v256 + int32(1)
	if v399 != l0 {
		v235 = v378
		v236 = v381
		v238 = v380
		v239 = v376
		v241 = v382
		v245 = v383
		v246 = v384
		v248 = v379
		v249 = v397
		v251 = v385
		v252 = v386
		v253 = v387
		v254 = v388
		v256 = v399
		v257 = v389
		v258 = v390
		v259 = v391
		v260 = v392
		v261 = v377
		goto L51
	} else {
		goto L126
	}
L125:
	;
	v397 = v249 + int32(1)
	goto L124
L126:
	;
	goto L52
L127:
	;
	if base.B2i32(v385 == int32(0))&v376 != 0 {
		goto L128
	} else {
		goto L129
	}
L128:
	;
	F_resolve_anyelement_from_others(m, v211+int32(16))
	mBase = m.M
	v420 = m.ExcPending
	if v420 != 0 {
		goto L80
	} else {
		goto L131
	}
L129:
	;
	v422 = v386
	goto L130
L130:
	;
	if base.B2i32(v422 == int32(0))&v380 != 0 {
		goto L132
	} else {
		goto L133
	}
L131:
	;
	v421 = *(*int32)(unsafe.Add(mBase, uint32(v211)+20))
	v422 = v421
	goto L130
L132:
	;
	F_resolve_anyarray_from_others(m, v211+int32(16))
	mBase = m.M
	v429 = m.ExcPending
	if v429 != 0 {
		goto L80
	} else {
		goto L135
	}
L133:
	;
	goto L134
L134:
	;
	v430 = *(*int32)(unsafe.Add(mBase, uint32(v211)+24))
	if base.B2i32(v430 == int32(0))&v378 != 0 {
		goto L136
	} else {
		goto L137
	}
L135:
	;
	goto L134
L136:
	;
	F_resolve_anyrange_from_others(m, v211+int32(16))
	mBase = m.M
	v437 = m.ExcPending
	if v437 != 0 {
		goto L80
	} else {
		goto L139
	}
L137:
	;
	goto L138
L138:
	;
	v438 = *(*int32)(unsafe.Add(mBase, uint32(v211)+28))
	if base.B2i32(v438 == int32(0))&v381 != 0 {
		goto L140
	} else {
		goto L141
	}
L139:
	;
	goto L138
L140:
	;
	F_resolve_anymultirange_from_others(m, v211+int32(16))
	mBase = m.M
	v445 = m.ExcPending
	if v445 != 0 {
		goto L80
	} else {
		goto L143
	}
L141:
	;
	goto L142
L142:
	;
	if base.B2i32(v387 == int32(0))&v382 != 0 {
		goto L144
	} else {
		goto L145
	}
L143:
	;
	goto L142
L144:
	;
	F_resolve_anyelement_from_others(m, v211)
	mBase = m.M
	v450 = m.ExcPending
	if v450 != 0 {
		goto L80
	} else {
		goto L147
	}
L145:
	;
	v452 = v388
	goto L146
L146:
	;
	if base.B2i32(v452 == int32(0))&v383 != 0 {
		goto L148
	} else {
		goto L149
	}
L147:
	;
	v451 = *(*int32)(unsafe.Add(mBase, uint32(v211)+4))
	v452 = v451
	goto L146
L148:
	;
	F_resolve_anyarray_from_others(m, v211)
	mBase = m.M
	v457 = m.ExcPending
	if v457 != 0 {
		goto L80
	} else {
		goto L151
	}
L149:
	;
	goto L150
L150:
	;
	v458 = *(*int32)(unsafe.Add(mBase, uint32(v211)+8))
	if base.B2i32(v458 == int32(0))&v384 != 0 {
		goto L152
	} else {
		goto L153
	}
L151:
	;
	goto L150
L152:
	;
	F_resolve_anyrange_from_others(m, v211)
	mBase = m.M
	v463 = m.ExcPending
	if v463 != 0 {
		goto L80
	} else {
		goto L155
	}
L153:
	;
	goto L154
L154:
	;
	v464 = *(*int32)(unsafe.Add(mBase, uint32(v211)+12))
	if base.B2i32(v464 == int32(0))&v379 != 0 {
		goto L156
	} else {
		goto L157
	}
L155:
	;
	goto L154
L156:
	;
	F_resolve_anymultirange_from_others(m, v211)
	mBase = m.M
	v469 = m.ExcPending
	if v469 != 0 {
		goto L80
	} else {
		goto L159
	}
L157:
	;
	goto L158
L158:
	;
	v470 = int32(0)
	v471 = *(*int32)(unsafe.Add(mBase, uint32(v211)+16))
	v472 = *(*int32)(unsafe.Add(mBase, uint32(v211)+20))
	v473 = *(*int32)(unsafe.Add(mBase, uint32(v211)+24))
	v474 = *(*int32)(unsafe.Add(mBase, uint32(v211)+28))
	v475 = *(*int32)(unsafe.Add(mBase, uint32(v211)))
	v476 = *(*int32)(unsafe.Add(mBase, uint32(v211)+4))
	v477 = *(*int32)(unsafe.Add(mBase, uint32(v211)+8))
	v478 = *(*int32)(unsafe.Add(mBase, uint32(v211)+12))
	if l0 == int32(1) {
		v584 = v470
		goto L160
	} else {
		goto L161
	}
L159:
	;
	goto L158
L160:
	;
	v610 = l1 + v584<<(uint(int32(2))%32)
	v611 = *(*int32)(unsafe.Add(mBase, uint32(v610)))
	if v611 <= int32(3830) {
		goto L200
	} else {
		goto L201
	}
L161:
	;
	v498 = int32(0)
	v500 = v470
	goto L162
L162:
	;
	v526 = l1 + v500<<(uint(int32(2))%32)
	v527 = *(*int32)(unsafe.Add(mBase, uint32(v526)))
	if v527 <= int32(3830) {
		goto L172
	} else {
		goto L173
	}
L163:
	;
	if l0&int32(1) != 0 {
		v584 = v566
		goto L160
	} else {
		goto L197
	}
L164:
	;
	v546 = *(*int32)(unsafe.Add(mBase, uint32(v526)+4))
	if v546 <= int32(3830) {
		goto L183
	} else {
		goto L184
	}
L165:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v526))) = v543
	goto L164
L166:
	;
	v543 = v478
	goto L165
L167:
	;
	v543 = v477
	goto L165
L168:
	;
	v543 = v476
	goto L165
L169:
	;
	v543 = v475
	goto L165
L170:
	;
	v543 = v474
	goto L165
L171:
	;
	v543 = v472
	goto L165
L172:
	;
	switch v527 - int32(2277) {
	case 0:
		goto L171
	case 1, 2, 3, 4, 5:
		goto L164
	case 6:
		v543 = v471
		goto L165
	default:
		goto L175
	}
L173:
	;
	goto L174
L174:
	;
	switch v527 - int32(_a_F_cfunc_resolve_polymorphic_argtypes_0) {
	case 0, 2:
		goto L169
	case 1:
		goto L168
	case 3:
		goto L167
	default:
		goto L177
	}
L175:
	;
	if base.B2i32(v527 == int32(2776))|base.B2i32(v527 == int32(3500)) != 0 {
		v543 = v471
		goto L165
	} else {
		goto L176
	}
L176:
	;
	goto L164
L177:
	;
	switch v527 - int32(_a_F_cfunc_resolve_polymorphic_argtypes_1) {
	case 0:
		goto L170
	case 1:
		goto L166
	default:
		goto L178
	}
L178:
	;
	if v527 == int32(3831) {
		v543 = v473
		goto L165
	} else {
		goto L179
	}
L179:
	;
	goto L164
L180:
	;
	v565 = int32(2)
	v566 = v500 + v565
	v568 = v498 + v565
	if v568 != l0&int32(2147483646) {
		v498 = v568
		v500 = v566
		goto L162
	} else {
		goto L196
	}
L181:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v526)+4)) = v562
	goto L180
L182:
	;
	v562 = v472
	goto L181
L183:
	;
	switch v546 - int32(2277) {
	case 0:
		goto L182
	case 1, 2, 3, 4, 5:
		goto L180
	case 6:
		v562 = v471
		goto L181
	default:
		goto L186
	}
L184:
	;
	goto L185
L185:
	;
	switch v546 - int32(_a_F_cfunc_resolve_polymorphic_argtypes_0) {
	case 0, 2:
		goto L189
	case 1:
		goto L190
	case 3:
		goto L191
	default:
		goto L192
	}
L186:
	;
	if base.B2i32(v546 == int32(2776))|base.B2i32(v546 == int32(3500)) != 0 {
		v562 = v471
		goto L181
	} else {
		goto L187
	}
L187:
	;
	goto L180
L188:
	;
	v562 = v474
	goto L181
L189:
	;
	v562 = v475
	goto L181
L190:
	;
	v562 = v476
	goto L181
L191:
	;
	v562 = v477
	goto L181
L192:
	;
	switch v546 - int32(_a_F_cfunc_resolve_polymorphic_argtypes_1) {
	case 0:
		goto L188
	case 1:
		goto L193
	default:
		goto L194
	}
L193:
	;
	v562 = v478
	goto L181
L194:
	;
	if v546 == int32(3831) {
		v562 = v473
		goto L181
	} else {
		goto L195
	}
L195:
	;
	goto L180
L196:
	;
	goto L163
L197:
	;
	v639 = v409
	goto L49
L198:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v610))) = v627
	v639 = v409
	goto L49
L199:
	;
	v627 = v472
	goto L198
L200:
	;
	switch v611 - int32(2277) {
	case 0:
		goto L199
	case 1, 2, 3, 4, 5:
		v639 = v409
		goto L49
	case 6:
		v627 = v471
		goto L198
	default:
		goto L203
	}
L201:
	;
	goto L202
L202:
	;
	switch v611 - int32(_a_F_cfunc_resolve_polymorphic_argtypes_0) {
	case 0, 2:
		goto L206
	case 1:
		goto L207
	case 3:
		goto L208
	default:
		goto L209
	}
L203:
	;
	if base.B2i32(v611 == int32(2776))|base.B2i32(v611 == int32(3500)) != 0 {
		v627 = v471
		goto L198
	} else {
		goto L204
	}
L204:
	;
	v639 = v409
	goto L49
L205:
	;
	v627 = v474
	goto L198
L206:
	;
	v627 = v475
	goto L198
L207:
	;
	v627 = v476
	goto L198
L208:
	;
	v627 = v477
	goto L198
L209:
	;
	switch v611 - int32(_a_F_cfunc_resolve_polymorphic_argtypes_1) {
	case 0:
		goto L205
	case 1:
		goto L210
	default:
		goto L211
	}
L210:
	;
	v627 = v478
	goto L198
L211:
	;
	if v611 == int32(3831) {
		v627 = v473
		goto L198
	} else {
		goto L212
	}
L212:
	;
	v639 = v409
	goto L49
L213:
	;
	if l0 <= int32(0) {
		goto L2
	} else {
		goto L214
	}
L214:
	;
	v674 = int32(0)
	v680 = v674
	v681 = v674
	goto L215
L215:
	;
	if l2 == int32(0) {
		goto L218
	} else {
		goto L219
	}
L216:
	;
	goto L2
L217:
	;
	v741 = v680 + int32(1)
	if v741 != l0 {
		v680 = v741
		v681 = v737
		goto L215
	} else {
		goto L224
	}
L218:
	;
	v722 = l1 + v680<<(uint(int32(2))%32)
	v723 = *(*int32)(unsafe.Add(mBase, uint32(v722)))
	if base.B2i32(v723 != int32(2287))&base.B2i32(v723 != int32(2249)) != 0 {
		goto L220
	} else {
		goto L221
	}
L219:
	;
	v717 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l2+v680))))
	switch v717 - int32(111) {
	case 0, 5:
		v737 = v681
		goto L217
	default:
		goto L218
	}
L220:
	;
	v737 = v681 + int32(1)
	goto L217
L221:
	;
	v729 = F_get_call_expr_argtype(m, l3, v681)
	mBase = m.M
	v730 = m.ExcPending
	if v730 != 0 {
		goto L80
	} else {
		goto L222
	}
L222:
	;
	if v729 == int32(0) {
		goto L220
	} else {
		goto L223
	}
L223:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v722))) = v729
	goto L220
L224:
	;
	goto L216
L225:
	;
	F_errcode(m, int32(1088))
	mBase = m.M
	v790 = m.ExcPending
	if v790 != 0 {
		goto L80
	} else {
		goto L226
	}
L226:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v41))) = l5
	F_errmsg(m, int32(_a_F_cfunc_resolve_polymorphic_argtypes_3), v41)
	mBase = m.M
	v794 = m.ExcPending
	if v794 != 0 {
		goto L80
	} else {
		goto L227
	}
L227:
	;
	F_errfinish(m, int32(_a_F_cfunc_resolve_polymorphic_argtypes_4), int32(380), int32(_a_F_cfunc_resolve_polymorphic_argtypes_5))
	mBase = m.M
	v799 = m.ExcPending
	if v799 != 0 {
		goto L80
	} else {
		goto L228
	}
L228:
	;
	base.Wasm_trap_unreachable()
	for {
	}
}
