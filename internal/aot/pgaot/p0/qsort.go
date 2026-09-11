package p0

import (
	base "github.com/shibukawa/pgmem/internal/aot/pgaot/base"
	"unsafe"
)

func F_qsort_partition_list_value_cmp(m *base.Module, l0 int32, l1 int32, l2 int32) int32 {
	mBase := m.M
	_ = mBase
	var v4 int32
	_ = v4
	var v5 int32
	_ = v5
	var v6 int32
	_ = v6
	var v7 int32
	_ = v7
	var v8 int32
	_ = v8
	var v9 int32
	_ = v9
	var v12 int32
	_ = v12
	v4 = *(*int32)(unsafe.Add(mBase, uint32(l2)+24))
	v5 = *(*int32)(unsafe.Add(mBase, uint32(l2)+28))
	v6 = *(*int32)(unsafe.Add(mBase, uint32(v5)))
	v7 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v8 = *(*int32)(unsafe.Add(mBase, uint32(l1)+4))
	v9 = F_FunctionCall2Coll(m, v4, v6, v7, v8)
	mBase = m.M
	v12 = m.ExcPending
	if v12 != 0 {
		return int32(0)
	} else {
		return v9
	}
}
func F_qsort_tuple_unsigned(m *base.Module, l0 int32, l1 int32, l2 int32) {
	mBase := m.M
	_ = mBase
	var v18 int32
	_ = v18
	var v20 int32
	_ = v20
	var v22 int32
	_ = v22
	var v23 int32
	_ = v23
	var v40 int32
	_ = v40
	var v42 int32
	_ = v42
	var v59 int32
	_ = v59
	var v61 int32
	_ = v61
	var v64 int32
	_ = v64
	var v82 int32
	_ = v82
	var v91 int32
	_ = v91
	var v104 int32
	_ = v104
	var v105 int32
	_ = v105
	var v106 int32
	_ = v106
	var v109 int32
	_ = v109
	var v114 int32
	_ = v114
	var v119 int32
	_ = v119
	var v122 int32
	_ = v122
	var v123 int32
	_ = v123
	var v125 int32
	_ = v125
	var v126 int32
	_ = v126
	var v127 int32
	_ = v127
	var v132 int32
	_ = v132
	var v135 int32
	_ = v135
	var v136 int32
	_ = v136
	var v137 int32
	_ = v137
	var v138 int32
	_ = v138
	var v139 int32
	_ = v139
	var v145 int32
	_ = v145
	var v146 int32
	_ = v146
	var v148 int32
	_ = v148
	var v149 int64
	_ = v149
	var v151 int64
	_ = v151
	var v154 int32
	_ = v154
	var v155 int64
	_ = v155
	var v157 int64
	_ = v157
	var v159 int64
	_ = v159
	var v161 int64
	_ = v161
	var v182 int32
	_ = v182
	var v189 int32
	_ = v189
	var v203 int32
	_ = v203
	var v205 int32
	_ = v205
	var v206 int32
	_ = v206
	var v207 int32
	_ = v207
	var v210 int32
	_ = v210
	var v215 int32
	_ = v215
	var v220 int32
	_ = v220
	var v225 int32
	_ = v225
	var v226 int32
	_ = v226
	var v228 int32
	_ = v228
	var v229 int32
	_ = v229
	var v230 int32
	_ = v230
	var v235 int32
	_ = v235
	var v238 int32
	_ = v238
	var v241 int32
	_ = v241
	var v242 int32
	_ = v242
	var v243 int32
	_ = v243
	var v245 int32
	_ = v245
	var v251 int32
	_ = v251
	var v259 int32
	_ = v259
	var v263 int32
	_ = v263
	var v267 int32
	_ = v267
	var v269 int32
	_ = v269
	var v272 int32
	_ = v272
	var v274 int32
	_ = v274
	var v275 int32
	_ = v275
	var v278 int32
	_ = v278
	var v279 int32
	_ = v279
	var v282 int32
	_ = v282
	var v283 int32
	_ = v283
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
	var v296 int32
	_ = v296
	var v297 int32
	_ = v297
	var v299 int32
	_ = v299
	var v300 int64
	_ = v300
	var v302 int64
	_ = v302
	var v305 int32
	_ = v305
	var v306 int64
	_ = v306
	var v308 int64
	_ = v308
	var v310 int64
	_ = v310
	var v312 int64
	_ = v312
	var v315 int32
	_ = v315
	var v320 int32
	_ = v320
	var v321 int32
	_ = v321
	var v325 int32
	_ = v325
	var v329 int32
	_ = v329
	var v339 int32
	_ = v339
	var v343 int32
	_ = v343
	var v351 int32
	_ = v351
	var v352 int32
	_ = v352
	var v353 int32
	_ = v353
	var v358 int32
	_ = v358
	var v361 int32
	_ = v361
	var v364 int32
	_ = v364
	var v365 int32
	_ = v365
	var v367 int32
	_ = v367
	var v368 int32
	_ = v368
	var v369 int32
	_ = v369
	var v374 int32
	_ = v374
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
	var v387 int32
	_ = v387
	var v388 int32
	_ = v388
	var v389 int64
	_ = v389
	var v391 int64
	_ = v391
	var v394 int32
	_ = v394
	var v395 int64
	_ = v395
	var v397 int64
	_ = v397
	var v399 int64
	_ = v399
	var v401 int64
	_ = v401
	var v407 int32
	_ = v407
	var v410 int32
	_ = v410
	var v412 int32
	_ = v412
	var v414 int32
	_ = v414
	var v421 int32
	_ = v421
	var v425 int32
	_ = v425
	var v438 int32
	_ = v438
	var v447 int32
	_ = v447
	var v451 int32
	_ = v451
	var v452 int32
	_ = v452
	var v453 int32
	_ = v453
	var v458 int32
	_ = v458
	var v461 int32
	_ = v461
	var v462 int32
	_ = v462
	var v463 int32
	_ = v463
	var v465 int32
	_ = v465
	var v466 int32
	_ = v466
	var v467 int32
	_ = v467
	var v472 int32
	_ = v472
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
	var v485 int32
	_ = v485
	var v486 int32
	_ = v486
	var v487 int64
	_ = v487
	var v489 int64
	_ = v489
	var v492 int32
	_ = v492
	var v493 int64
	_ = v493
	var v495 int64
	_ = v495
	var v497 int64
	_ = v497
	var v499 int64
	_ = v499
	var v506 int32
	_ = v506
	var v508 int32
	_ = v508
	var v510 int32
	_ = v510
	var v512 int32
	_ = v512
	var v518 int32
	_ = v518
	var v527 int32
	_ = v527
	var v532 int32
	_ = v532
	var v533 int32
	_ = v533
	var v536 int32
	_ = v536
	var v538 int32
	_ = v538
	var v546 int32
	_ = v546
	var v561 int32
	_ = v561
	var v562 int32
	_ = v562
	var v563 int32
	_ = v563
	var v564 int32
	_ = v564
	var v565 int64
	_ = v565
	var v567 int64
	_ = v567
	var v569 int32
	_ = v569
	var v571 int32
	_ = v571
	var v572 int64
	_ = v572
	var v574 int64
	_ = v574
	var v576 int64
	_ = v576
	var v578 int64
	_ = v578
	var v581 int32
	_ = v581
	var v601 int32
	_ = v601
	var v602 int32
	_ = v602
	var v607 int32
	_ = v607
	var v609 int32
	_ = v609
	var v618 int32
	_ = v618
	var v632 int32
	_ = v632
	var v633 int32
	_ = v633
	var v634 int32
	_ = v634
	var v635 int32
	_ = v635
	var v636 int64
	_ = v636
	var v638 int64
	_ = v638
	var v640 int32
	_ = v640
	var v642 int32
	_ = v642
	var v643 int64
	_ = v643
	var v645 int64
	_ = v645
	var v647 int64
	_ = v647
	var v649 int64
	_ = v649
	var v652 int32
	_ = v652
	var v675 int32
	_ = v675
	var v687 int32
	_ = v687
	var v692 int32
	_ = v692
	var v693 int32
	_ = v693
	var v694 int64
	_ = v694
	var v696 int64
	_ = v696
	var v699 int32
	_ = v699
	var v700 int64
	_ = v700
	var v702 int64
	_ = v702
	var v704 int64
	_ = v704
	var v706 int64
	_ = v706
	var v708 int32
	_ = v708
	v18 = m.G0
	v20 = v18 - int32(16)
	m.G0 = v20
	v22 = l0
	v23 = l1
	goto L1
L1:
	;
	v40 = v22 + int32(16)
	v42 = v23
	goto L3
L2:
	;
	m.G0 = v20 + int32(16)
	return
L3:
	;
	v59 = *(*int32)(unsafe.Add(mBase, _consts[1]))
	if v59 != 0 {
		goto L5
	} else {
		goto L6
	}
L4:
	;
	goto L2
L5:
	;
	F_ProcessInterrupts(m)
	mBase = m.M
	v61 = m.ExcPending
	if v61 != 0 {
		goto L8
	} else {
		goto L9
	}
L6:
	;
	goto L7
L7:
	;
	v64 = v22 + v42<<(uint(int32(4))%32)
	if base.Ui32(v42) <= base.Ui32(int32(6)) {
		goto L11
	} else {
		goto L12
	}
L8:
	;
	return
L9:
	;
	goto L7
L10:
	;
	goto L4
L11:
	;
	if base.Ui32(v64) <= base.Ui32(v40) {
		goto L10
	} else {
		goto L14
	}
L12:
	;
	goto L13
L13:
	;
	if base.Ui32(v64) <= base.Ui32(v40) {
		goto L10
	} else {
		goto L43
	}
L14:
	;
	v82 = v40
	goto L15
L15:
	;
	if base.Ui32(v82) <= base.Ui32(v22) {
		goto L17
	} else {
		goto L18
	}
L16:
	;
	goto L10
L17:
	;
	v182 = v82 + int32(16)
	if base.Ui32(v182) < base.Ui32(v64) {
		v82 = v182
		goto L15
	} else {
		goto L42
	}
L18:
	;
	v91 = v82
	goto L19
L19:
	;
	v104 = v91 - int32(16)
	v105 = *(*int32)(unsafe.Add(mBase, uint32(l2)+44))
	v106 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v91)+8)))
	v109 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v91-int32(8)))))
	if v109 == int32(1) {
		goto L24
	} else {
		goto L25
	}
L20:
	;
	goto L17
L21:
	;
	v145 = int32(8)
	v146 = v20 + v145
	v148 = v91 + v145
	v149 = *(*int64)(unsafe.Add(mBase, uint32(v148)))
	*(*int64)(unsafe.Add(mBase, uint32(v146))) = v149
	v151 = *(*int64)(unsafe.Add(mBase, uint32(v91)))
	*(*int64)(unsafe.Add(mBase, uint32(v20))) = v151
	v154 = v104 + v145
	v155 = *(*int64)(unsafe.Add(mBase, uint32(v154)))
	*(*int64)(unsafe.Add(mBase, uint32(v148))) = v155
	v157 = *(*int64)(unsafe.Add(mBase, uint32(v104)))
	*(*int64)(unsafe.Add(mBase, uint32(v91))) = v157
	v159 = *(*int64)(unsafe.Add(mBase, uint32(v146)))
	*(*int64)(unsafe.Add(mBase, uint32(v154))) = v159
	v161 = *(*int64)(unsafe.Add(mBase, uint32(v20)))
	*(*int64)(unsafe.Add(mBase, uint32(v104))) = v161
	if base.Ui32(v22) < base.Ui32(v104) {
		v91 = v104
		goto L19
	} else {
		goto L41
	}
L22:
	;
	if v139 <= int32(0) {
		goto L17
	} else {
		goto L40
	}
L23:
	;
	v135 = *(*int32)(unsafe.Add(mBase, uint32(l2)+48))
	if v135 != 0 {
		goto L17
	} else {
		goto L38
	}
L24:
	;
	if v106&int32(1) != 0 {
		goto L23
	} else {
		goto L27
	}
L25:
	;
	goto L26
L26:
	;
	if v106&int32(1) != 0 {
		goto L29
	} else {
		goto L30
	}
L27:
	;
	v114 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v105)+9)))
	if v114 == int32(0) {
		goto L21
	} else {
		goto L28
	}
L28:
	;
	goto L17
L29:
	;
	v119 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v105)+9)))
	if v119 != 0 {
		goto L21
	} else {
		goto L32
	}
L30:
	;
	goto L31
L31:
	;
	v122 = *(*int32)(unsafe.Add(mBase, uint32(v91-int32(12))))
	v123 = *(*int32)(unsafe.Add(mBase, uint32(v91)+4))
	v125 = base.B2i32(base.Ui32(v122) < base.Ui32(v123))
	v126 = base.B2i32(base.Ui32(v123) < base.Ui32(v122)) - v125
	v127 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v105)+8)))
	if v127 == int32(1) {
		goto L33
	} else {
		goto L34
	}
L32:
	;
	goto L17
L33:
	;
	if base.Ui32(v122) < base.Ui32(v123) {
		goto L21
	} else {
		goto L36
	}
L34:
	;
	v132 = v126
	goto L35
L35:
	;
	if v132 != 0 {
		v139 = v132
		goto L22
	} else {
		goto L37
	}
L36:
	;
	v132 = int32(0) - v126
	goto L35
L37:
	;
	goto L23
L38:
	;
	v136 = *(*int32)(unsafe.Add(mBase, uint32(l2)+4))
	v137 = m.T0[v136].(func(*base.Module, int32, int32, int32) int32)(m, v104, v91, l2)
	mBase = m.M
	v138 = m.ExcPending
	if v138 != 0 {
		goto L8
	} else {
		goto L39
	}
L39:
	;
	v139 = v137
	goto L22
L40:
	;
	goto L21
L41:
	;
	goto L20
L42:
	;
	goto L16
L43:
	;
	v189 = v40
	goto L44
L44:
	;
	v203 = *(*int32)(unsafe.Add(mBase, _consts[1]))
	if v203 != 0 {
		goto L46
	} else {
		goto L47
	}
L45:
	;
	v259 = v22 + v42<<(uint(int32(3))%32)&int32(-16)
	if v42 != int32(7) {
		goto L72
	} else {
		goto L73
	}
L46:
	;
	F_ProcessInterrupts(m)
	mBase = m.M
	v205 = m.ExcPending
	if v205 != 0 {
		goto L8
	} else {
		goto L49
	}
L47:
	;
	goto L48
L48:
	;
	v206 = *(*int32)(unsafe.Add(mBase, uint32(l2)+44))
	v207 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v189)+8)))
	v210 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v189-int32(8)))))
	if v210 == int32(1) {
		goto L54
	} else {
		goto L55
	}
L49:
	;
	goto L48
L50:
	;
	goto L45
L51:
	;
	v251 = v189 + int32(16)
	if base.Ui32(v251) < base.Ui32(v64) {
		v189 = v251
		goto L44
	} else {
		goto L71
	}
L52:
	;
	if int32(0) < v245 {
		goto L50
	} else {
		goto L70
	}
L53:
	;
	v238 = *(*int32)(unsafe.Add(mBase, uint32(l2)+48))
	if v238 != 0 {
		goto L51
	} else {
		goto L68
	}
L54:
	;
	if v207&int32(1) != 0 {
		goto L53
	} else {
		goto L57
	}
L55:
	;
	goto L56
L56:
	;
	if v207&int32(1) != 0 {
		goto L59
	} else {
		goto L60
	}
L57:
	;
	v215 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v206)+9)))
	if v215 == int32(0) {
		goto L50
	} else {
		goto L58
	}
L58:
	;
	goto L51
L59:
	;
	v220 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v206)+9)))
	if v220 == int32(0) {
		goto L51
	} else {
		goto L62
	}
L60:
	;
	goto L61
L61:
	;
	v225 = *(*int32)(unsafe.Add(mBase, uint32(v189-int32(12))))
	v226 = *(*int32)(unsafe.Add(mBase, uint32(v189)+4))
	v228 = base.B2i32(base.Ui32(v225) < base.Ui32(v226))
	v229 = base.B2i32(base.Ui32(v226) < base.Ui32(v225)) - v228
	v230 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v206)+8)))
	if v230 == int32(1) {
		goto L63
	} else {
		goto L64
	}
L62:
	;
	goto L50
L63:
	;
	if base.Ui32(v225) < base.Ui32(v226) {
		goto L50
	} else {
		goto L66
	}
L64:
	;
	v235 = v229
	goto L65
L65:
	;
	if v235 != 0 {
		v245 = v235
		goto L52
	} else {
		goto L67
	}
L66:
	;
	v235 = int32(0) - v229
	goto L65
L67:
	;
	goto L53
L68:
	;
	v241 = *(*int32)(unsafe.Add(mBase, uint32(l2)+4))
	v242 = m.T0[v241].(func(*base.Module, int32, int32, int32) int32)(m, v189-int32(16), v189, l2)
	mBase = m.M
	v243 = m.ExcPending
	if v243 != 0 {
		goto L8
	} else {
		goto L69
	}
L69:
	;
	v245 = v242
	goto L52
L70:
	;
	goto L51
L71:
	;
	goto L10
L72:
	;
	v263 = v64 - int32(16)
	if base.Ui32(v42) < base.Ui32(int32(41)) {
		goto L76
	} else {
		goto L77
	}
L73:
	;
	v292 = v259
	goto L74
L74:
	;
	v296 = int32(8)
	v297 = v20 + v296
	v299 = v22 + v296
	v300 = *(*int64)(unsafe.Add(mBase, uint32(v299)))
	*(*int64)(unsafe.Add(mBase, uint32(v297))) = v300
	v302 = *(*int64)(unsafe.Add(mBase, uint32(v22)))
	*(*int64)(unsafe.Add(mBase, uint32(v20))) = v302
	v305 = v292 + v296
	v306 = *(*int64)(unsafe.Add(mBase, uint32(v305)))
	*(*int64)(unsafe.Add(mBase, uint32(v299))) = v306
	v308 = *(*int64)(unsafe.Add(mBase, uint32(v292)))
	*(*int64)(unsafe.Add(mBase, uint32(v22))) = v308
	v310 = *(*int64)(unsafe.Add(mBase, uint32(v297)))
	*(*int64)(unsafe.Add(mBase, uint32(v305))) = v310
	v312 = *(*int64)(unsafe.Add(mBase, uint32(v20)))
	*(*int64)(unsafe.Add(mBase, uint32(v292))) = v312
	v315 = v64 - int32(16)
	v320 = v315
	v321 = v40
	v325 = v40
	v329 = v315
	goto L83
L75:
	;
	v289 = F_qsort_tuple_unsigned_med3(m, v287, v285, v286, l2)
	mBase = m.M
	v290 = m.ExcPending
	if v290 != 0 {
		goto L8
	} else {
		goto L82
	}
L76:
	;
	v285 = v259
	v286 = v263
	v287 = v22
	goto L75
L77:
	;
	goto L78
L78:
	;
	v267 = int32(base.Ui32(v42) >> (uint(int32(3)) % 32))
	v269 = v267 << (uint(int32(4)) % 32)
	v272 = v267 << (uint(int32(5)) % 32)
	v274 = F_qsort_tuple_unsigned_med3(m, v22, v22+v269, v22+v272, l2)
	mBase = m.M
	v275 = m.ExcPending
	if v275 != 0 {
		goto L8
	} else {
		goto L79
	}
L79:
	;
	v278 = F_qsort_tuple_unsigned_med3(m, v259-v269, v259, v269+v259, l2)
	mBase = m.M
	v279 = m.ExcPending
	if v279 != 0 {
		goto L8
	} else {
		goto L80
	}
L80:
	;
	v282 = F_qsort_tuple_unsigned_med3(m, v263-v272, v263-v269, v263, l2)
	mBase = m.M
	v283 = m.ExcPending
	if v283 != 0 {
		goto L8
	} else {
		goto L81
	}
L81:
	;
	v285 = v278
	v286 = v282
	v287 = v274
	goto L75
L82:
	;
	v292 = v289
	goto L74
L83:
	;
	if base.Ui32(v320) < base.Ui32(v321) {
		v421 = v321
		v425 = v325
		goto L85
	} else {
		goto L86
	}
L85:
	;
	if base.Ui32(v421) <= base.Ui32(v320) {
		goto L117
	} else {
		goto L118
	}
L86:
	;
	v339 = v321
	v343 = v325
	goto L87
L87:
	;
	v351 = *(*int32)(unsafe.Add(mBase, uint32(l2)+44))
	v352 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v22)+8)))
	v353 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v339)+8)))
	if v353 == int32(1) {
		goto L93
	} else {
		goto L94
	}
L88:
	;
	v421 = v414
	v425 = v407
	goto L85
L89:
	;
	v410 = *(*int32)(unsafe.Add(mBase, _consts[1]))
	if v410 != 0 {
		goto L111
	} else {
		goto L112
	}
L90:
	;
	v387 = int32(8)
	v388 = v343 + v387
	v389 = *(*int64)(unsafe.Add(mBase, uint32(v388)))
	*(*int64)(unsafe.Add(mBase, uint32(v297))) = v389
	v391 = *(*int64)(unsafe.Add(mBase, uint32(v343)))
	*(*int64)(unsafe.Add(mBase, uint32(v20))) = v391
	v394 = v339 + v387
	v395 = *(*int64)(unsafe.Add(mBase, uint32(v394)))
	*(*int64)(unsafe.Add(mBase, uint32(v388))) = v395
	v397 = *(*int64)(unsafe.Add(mBase, uint32(v339)))
	*(*int64)(unsafe.Add(mBase, uint32(v343))) = v397
	v399 = *(*int64)(unsafe.Add(mBase, uint32(v297)))
	*(*int64)(unsafe.Add(mBase, uint32(v394))) = v399
	v401 = *(*int64)(unsafe.Add(mBase, uint32(v20)))
	*(*int64)(unsafe.Add(mBase, uint32(v339))) = v401
	v407 = v343 + int32(16)
	goto L89
L91:
	;
	if int32(0) < v381 {
		v421 = v339
		v425 = v343
		goto L85
	} else {
		goto L109
	}
L92:
	;
	v377 = *(*int32)(unsafe.Add(mBase, uint32(l2)+48))
	if v377 != 0 {
		goto L90
	} else {
		goto L107
	}
L93:
	;
	if v352&int32(1) != 0 {
		goto L92
	} else {
		goto L96
	}
L94:
	;
	goto L95
L95:
	;
	if v352&int32(1) != 0 {
		goto L98
	} else {
		goto L99
	}
L96:
	;
	v358 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v351)+9)))
	if v358 != 0 {
		v407 = v343
		goto L89
	} else {
		goto L97
	}
L97:
	;
	v421 = v339
	v425 = v343
	goto L85
L98:
	;
	v361 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v351)+9)))
	if v361 == int32(0) {
		v407 = v343
		goto L89
	} else {
		goto L101
	}
L99:
	;
	goto L100
L100:
	;
	v364 = *(*int32)(unsafe.Add(mBase, uint32(v339)+4))
	v365 = *(*int32)(unsafe.Add(mBase, uint32(v22)+4))
	v367 = base.B2i32(base.Ui32(v364) < base.Ui32(v365))
	v368 = base.B2i32(base.Ui32(v365) < base.Ui32(v364)) - v367
	v369 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v351)+8)))
	if v369 == int32(1) {
		goto L102
	} else {
		goto L103
	}
L101:
	;
	v421 = v339
	v425 = v343
	goto L85
L102:
	;
	if base.Ui32(v364) < base.Ui32(v365) {
		v421 = v339
		v425 = v343
		goto L85
	} else {
		goto L105
	}
L103:
	;
	v374 = v368
	goto L104
L104:
	;
	if v374 != 0 {
		v381 = v374
		goto L91
	} else {
		goto L106
	}
L105:
	;
	v374 = int32(0) - v368
	goto L104
L106:
	;
	goto L92
L107:
	;
	v378 = *(*int32)(unsafe.Add(mBase, uint32(l2)+4))
	v379 = m.T0[v378].(func(*base.Module, int32, int32, int32) int32)(m, v339, v22, l2)
	mBase = m.M
	v380 = m.ExcPending
	if v380 != 0 {
		goto L8
	} else {
		goto L108
	}
L108:
	;
	v381 = v379
	goto L91
L109:
	;
	if v381 != 0 {
		v407 = v343
		goto L89
	} else {
		goto L110
	}
L110:
	;
	goto L90
L111:
	;
	F_ProcessInterrupts(m)
	mBase = m.M
	v412 = m.ExcPending
	if v412 != 0 {
		goto L8
	} else {
		goto L114
	}
L112:
	;
	goto L113
L113:
	;
	v414 = v339 + int32(16)
	if base.Ui32(v414) <= base.Ui32(v320) {
		v339 = v414
		v343 = v407
		goto L87
	} else {
		goto L115
	}
L114:
	;
	goto L113
L115:
	;
	goto L88
L116:
	;
	v692 = int32(8)
	v693 = v421 + v692
	v694 = *(*int64)(unsafe.Add(mBase, uint32(v693)))
	*(*int64)(unsafe.Add(mBase, uint32(v297))) = v694
	v696 = *(*int64)(unsafe.Add(mBase, uint32(v421)))
	*(*int64)(unsafe.Add(mBase, uint32(v20))) = v696
	v699 = v438 + v692
	v700 = *(*int64)(unsafe.Add(mBase, uint32(v699)))
	*(*int64)(unsafe.Add(mBase, uint32(v693))) = v700
	v702 = *(*int64)(unsafe.Add(mBase, uint32(v438)))
	*(*int64)(unsafe.Add(mBase, uint32(v421))) = v702
	v704 = *(*int64)(unsafe.Add(mBase, uint32(v297)))
	*(*int64)(unsafe.Add(mBase, uint32(v699))) = v704
	v706 = *(*int64)(unsafe.Add(mBase, uint32(v20)))
	*(*int64)(unsafe.Add(mBase, uint32(v438))) = v706
	v708 = int32(16)
	v320 = v438 - v708
	v321 = v421 + v708
	v325 = v425
	v329 = v447
	goto L83
L117:
	;
	v438 = v320
	v447 = v329
	goto L120
L118:
	;
	v518 = v320
	v527 = v329
	goto L119
L119:
	;
	v532 = int32(4)
	v533 = (v425 - v22) >> (uint(v532) % 32)
	v536 = (v421 - v425) >> (uint(v532) % 32)
	if v533 < v536 {
		goto L149
	} else {
		goto L150
	}
L120:
	;
	v451 = *(*int32)(unsafe.Add(mBase, uint32(l2)+44))
	v452 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v22)+8)))
	v453 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v438)+8)))
	if v453 == int32(1) {
		goto L126
	} else {
		goto L127
	}
L121:
	;
	v518 = v512
	v527 = v506
	goto L119
L122:
	;
	v508 = *(*int32)(unsafe.Add(mBase, _consts[1]))
	if v508 != 0 {
		goto L144
	} else {
		goto L145
	}
L123:
	;
	v485 = int32(8)
	v486 = v438 + v485
	v487 = *(*int64)(unsafe.Add(mBase, uint32(v486)))
	*(*int64)(unsafe.Add(mBase, uint32(v297))) = v487
	v489 = *(*int64)(unsafe.Add(mBase, uint32(v438)))
	*(*int64)(unsafe.Add(mBase, uint32(v20))) = v489
	v492 = v447 + v485
	v493 = *(*int64)(unsafe.Add(mBase, uint32(v492)))
	*(*int64)(unsafe.Add(mBase, uint32(v486))) = v493
	v495 = *(*int64)(unsafe.Add(mBase, uint32(v447)))
	*(*int64)(unsafe.Add(mBase, uint32(v438))) = v495
	v497 = *(*int64)(unsafe.Add(mBase, uint32(v297)))
	*(*int64)(unsafe.Add(mBase, uint32(v492))) = v497
	v499 = *(*int64)(unsafe.Add(mBase, uint32(v20)))
	*(*int64)(unsafe.Add(mBase, uint32(v447))) = v499
	v506 = v447 - int32(16)
	goto L122
L124:
	;
	if v479 < int32(0) {
		goto L116
	} else {
		goto L142
	}
L125:
	;
	v475 = *(*int32)(unsafe.Add(mBase, uint32(l2)+48))
	if v475 != 0 {
		goto L123
	} else {
		goto L140
	}
L126:
	;
	if v452&int32(1) != 0 {
		goto L125
	} else {
		goto L129
	}
L127:
	;
	goto L128
L128:
	;
	if v452&int32(1) != 0 {
		goto L131
	} else {
		goto L132
	}
L129:
	;
	v458 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v451)+9)))
	if v458 != 0 {
		goto L116
	} else {
		goto L130
	}
L130:
	;
	v506 = v447
	goto L122
L131:
	;
	v461 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v451)+9)))
	if v461 != 0 {
		v506 = v447
		goto L122
	} else {
		goto L134
	}
L132:
	;
	goto L133
L133:
	;
	v462 = *(*int32)(unsafe.Add(mBase, uint32(v438)+4))
	v463 = *(*int32)(unsafe.Add(mBase, uint32(v22)+4))
	v465 = base.B2i32(base.Ui32(v462) < base.Ui32(v463))
	v466 = base.B2i32(base.Ui32(v463) < base.Ui32(v462)) - v465
	v467 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v451)+8)))
	if v467 == int32(1) {
		goto L135
	} else {
		goto L136
	}
L134:
	;
	goto L116
L135:
	;
	if base.Ui32(v462) < base.Ui32(v463) {
		v506 = v447
		goto L122
	} else {
		goto L138
	}
L136:
	;
	v472 = v466
	goto L137
L137:
	;
	if v472 != 0 {
		v479 = v472
		goto L124
	} else {
		goto L139
	}
L138:
	;
	v472 = int32(0) - v466
	goto L137
L139:
	;
	goto L125
L140:
	;
	v476 = *(*int32)(unsafe.Add(mBase, uint32(l2)+4))
	v477 = m.T0[v476].(func(*base.Module, int32, int32, int32) int32)(m, v438, v22, l2)
	mBase = m.M
	v478 = m.ExcPending
	if v478 != 0 {
		goto L8
	} else {
		goto L141
	}
L141:
	;
	v479 = v477
	goto L124
L142:
	;
	if v479 != 0 {
		v506 = v447
		goto L122
	} else {
		goto L143
	}
L143:
	;
	goto L123
L144:
	;
	F_ProcessInterrupts(m)
	mBase = m.M
	v510 = m.ExcPending
	if v510 != 0 {
		goto L8
	} else {
		goto L147
	}
L145:
	;
	goto L146
L146:
	;
	v512 = v438 - int32(16)
	if base.Ui32(v421) <= base.Ui32(v512) {
		v438 = v512
		v447 = v506
		goto L120
	} else {
		goto L148
	}
L147:
	;
	goto L146
L148:
	;
	goto L121
L149:
	;
	v538 = v533
	goto L151
L150:
	;
	v538 = v536
	goto L151
L151:
	;
	if v538 != 0 {
		goto L152
	} else {
		goto L153
	}
L152:
	;
	v546 = int32(0)
	goto L155
L153:
	;
	goto L154
L154:
	;
	v601 = int32(4)
	v602 = (v527 - v518) >> (uint(v601) % 32)
	v607 = (v64-v527)>>(uint(v601)%32) - int32(1)
	if v602 < v607 {
		goto L158
	} else {
		goto L159
	}
L155:
	;
	v561 = v546 << (uint(int32(4)) % 32)
	v562 = v22 + v561
	v563 = int32(8)
	v564 = v562 + v563
	v565 = *(*int64)(unsafe.Add(mBase, uint32(v564)))
	*(*int64)(unsafe.Add(mBase, uint32(v297))) = v565
	v567 = *(*int64)(unsafe.Add(mBase, uint32(v562)))
	*(*int64)(unsafe.Add(mBase, uint32(v20))) = v567
	v569 = v561 + (v421 - v538<<(uint(int32(4))%32))
	v571 = v569 + v563
	v572 = *(*int64)(unsafe.Add(mBase, uint32(v571)))
	*(*int64)(unsafe.Add(mBase, uint32(v564))) = v572
	v574 = *(*int64)(unsafe.Add(mBase, uint32(v569)))
	*(*int64)(unsafe.Add(mBase, uint32(v562))) = v574
	v576 = *(*int64)(unsafe.Add(mBase, uint32(v297)))
	*(*int64)(unsafe.Add(mBase, uint32(v571))) = v576
	v578 = *(*int64)(unsafe.Add(mBase, uint32(v20)))
	*(*int64)(unsafe.Add(mBase, uint32(v569))) = v578
	v581 = v546 + int32(1)
	if v581 != v538 {
		v546 = v581
		goto L155
	} else {
		goto L157
	}
L156:
	;
	goto L154
L157:
	;
	goto L156
L158:
	;
	v609 = v602
	goto L160
L159:
	;
	v609 = v607
	goto L160
L160:
	;
	if v609 != 0 {
		goto L161
	} else {
		goto L162
	}
L161:
	;
	v618 = int32(0)
	goto L164
L162:
	;
	goto L163
L163:
	;
	if base.Ui32(v536) <= base.Ui32(v602) {
		goto L167
	} else {
		goto L168
	}
L164:
	;
	v632 = v618 << (uint(int32(4)) % 32)
	v633 = v421 + v632
	v634 = int32(8)
	v635 = v633 + v634
	v636 = *(*int64)(unsafe.Add(mBase, uint32(v635)))
	*(*int64)(unsafe.Add(mBase, uint32(v297))) = v636
	v638 = *(*int64)(unsafe.Add(mBase, uint32(v633)))
	*(*int64)(unsafe.Add(mBase, uint32(v20))) = v638
	v640 = v632 + (v64 - v609<<(uint(int32(4))%32))
	v642 = v640 + v634
	v643 = *(*int64)(unsafe.Add(mBase, uint32(v642)))
	*(*int64)(unsafe.Add(mBase, uint32(v635))) = v643
	v645 = *(*int64)(unsafe.Add(mBase, uint32(v640)))
	*(*int64)(unsafe.Add(mBase, uint32(v633))) = v645
	v647 = *(*int64)(unsafe.Add(mBase, uint32(v297)))
	*(*int64)(unsafe.Add(mBase, uint32(v642))) = v647
	v649 = *(*int64)(unsafe.Add(mBase, uint32(v20)))
	*(*int64)(unsafe.Add(mBase, uint32(v640))) = v649
	v652 = v618 + int32(1)
	if v652 != v609 {
		v618 = v652
		goto L164
	} else {
		goto L166
	}
L165:
	;
	goto L163
L166:
	;
	goto L165
L167:
	;
	if base.Ui32(int32(2)) <= base.Ui32(v536) {
		goto L170
	} else {
		goto L171
	}
L168:
	;
	goto L169
L169:
	;
	if base.Ui32(int32(2)) <= base.Ui32(v602) {
		goto L175
	} else {
		goto L176
	}
L170:
	;
	F_qsort_tuple_unsigned(m, v22, v536, l2)
	mBase = m.M
	v675 = m.ExcPending
	if v675 != 0 {
		goto L8
	} else {
		goto L173
	}
L171:
	;
	goto L172
L172:
	;
	if base.Ui32(v602) < base.Ui32(int32(2)) {
		goto L10
	} else {
		goto L174
	}
L173:
	;
	goto L172
L174:
	;
	v22 = v64 - v602<<(uint(int32(4))%32)
	v23 = v602
	goto L1
L175:
	;
	F_qsort_tuple_unsigned(m, v64-v602<<(uint(int32(4))%32), v602, l2)
	mBase = m.M
	v687 = m.ExcPending
	if v687 != 0 {
		goto L8
	} else {
		goto L178
	}
L176:
	;
	goto L177
L177:
	;
	if base.Ui32(int32(1)) < base.Ui32(v536) {
		v42 = v536
		goto L3
	} else {
		goto L179
	}
L178:
	;
	goto L177
L179:
	;
	goto L10
}
func F_qsort_tuple_unsigned_med3(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) int32 {
	mBase := m.M
	_ = mBase
	var v10 int32
	_ = v10
	var v11 int32
	_ = v11
	var v12 int32
	_ = v12
	var v13 int32
	_ = v13
	var v18 int32
	_ = v18
	var v21 int32
	_ = v21
	var v22 int32
	_ = v22
	var v24 int32
	_ = v24
	var v25 int32
	_ = v25
	var v26 int32
	_ = v26
	var v31 int32
	_ = v31
	var v34 int32
	_ = v34
	var v35 int32
	_ = v35
	var v36 int32
	_ = v36
	var v39 int32
	_ = v39
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
	var v45 int32
	_ = v45
	var v47 int32
	_ = v47
	var v50 int32
	_ = v50
	var v52 int32
	_ = v52
	var v54 int32
	_ = v54
	var v55 int32
	_ = v55
	var v56 int32
	_ = v56
	var v61 int32
	_ = v61
	var v66 int32
	_ = v66
	var v68 int32
	_ = v68
	var v69 int32
	_ = v69
	var v70 int32
	_ = v70
	var v75 int32
	_ = v75
	var v78 int32
	_ = v78
	var v79 int32
	_ = v79
	var v80 int32
	_ = v80
	var v81 int32
	_ = v81
	var v82 int32
	_ = v82
	var v86 int32
	_ = v86
	var v87 int32
	_ = v87
	var v88 int32
	_ = v88
	var v90 int32
	_ = v90
	var v91 int32
	_ = v91
	var v92 int32
	_ = v92
	var v94 int32
	_ = v94
	var v99 int32
	_ = v99
	var v102 int32
	_ = v102
	var v105 int32
	_ = v105
	var v107 int32
	_ = v107
	var v108 int32
	_ = v108
	var v109 int32
	_ = v109
	var v114 int32
	_ = v114
	var v117 int32
	_ = v117
	var v118 int32
	_ = v118
	var v119 int32
	_ = v119
	var v120 int32
	_ = v120
	var v122 int32
	_ = v122
	var v127 int32
	_ = v127
	var v129 int32
	_ = v129
	var v131 int32
	_ = v131
	var v132 int32
	_ = v132
	var v133 int32
	_ = v133
	var v138 int32
	_ = v138
	var v141 int32
	_ = v141
	var v145 int32
	_ = v145
	var v146 int32
	_ = v146
	var v147 int32
	_ = v147
	var v152 int32
	_ = v152
	var v155 int32
	_ = v155
	var v156 int32
	_ = v156
	var v157 int32
	_ = v157
	var v158 int32
	_ = v158
	var v159 int32
	_ = v159
	var v163 int32
	_ = v163
	var v164 int32
	_ = v164
	var v165 int32
	_ = v165
	var v167 int32
	_ = v167
	var v168 int32
	_ = v168
	var v169 int32
	_ = v169
	var v171 int32
	_ = v171
	var v176 int32
	_ = v176
	var v179 int32
	_ = v179
	var v182 int32
	_ = v182
	var v184 int32
	_ = v184
	var v185 int32
	_ = v185
	var v186 int32
	_ = v186
	var v192 int32
	_ = v192
	var v195 int32
	_ = v195
	var v197 int32
	_ = v197
	var v198 int32
	_ = v198
	var v199 int32
	_ = v199
	var v201 int32
	_ = v201
	var v210 int32
	_ = v210
	v10 = *(*int32)(unsafe.Add(mBase, uint32(l3)+44))
	v11 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l1)+8)))
	v12 = *(*int32)(unsafe.Add(mBase, uint32(l1)+4))
	v13 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+8)))
	if v13 == int32(1) {
		goto L7
	} else {
		goto L8
	}
L1:
	;
	return v210
L2:
	;
	v210 = l0
	goto L1
L3:
	;
	v132 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l2)+8)))
	v133 = *(*int32)(unsafe.Add(mBase, uint32(l2)+4))
	if v127&int32(1) != 0 {
		goto L68
	} else {
		goto L69
	}
L4:
	;
	v55 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l2)+8)))
	v56 = *(*int32)(unsafe.Add(mBase, uint32(l2)+4))
	if v50&int32(1) != 0 {
		goto L28
	} else {
		goto L29
	}
L5:
	;
	if int32(0) <= v44 {
		v127 = v43
		v129 = v45
		v131 = v47
		goto L3
	} else {
		goto L24
	}
L6:
	;
	v34 = *(*int32)(unsafe.Add(mBase, uint32(l3)+48))
	if v34 != 0 {
		v127 = v11
		v129 = v10
		v131 = v12
		goto L3
	} else {
		goto L21
	}
L7:
	;
	if v11&int32(1) != 0 {
		goto L6
	} else {
		goto L10
	}
L8:
	;
	goto L9
L9:
	;
	if v11&int32(1) != 0 {
		goto L12
	} else {
		goto L13
	}
L10:
	;
	v18 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v10)+9)))
	if v18 != 0 {
		v50 = v11
		v52 = v10
		v54 = v12
		goto L4
	} else {
		goto L11
	}
L11:
	;
	v127 = v11
	v129 = v10
	v131 = v12
	goto L3
L12:
	;
	v21 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v10)+9)))
	if v21 != 0 {
		v127 = v11
		v129 = v10
		v131 = v12
		goto L3
	} else {
		goto L15
	}
L13:
	;
	goto L14
L14:
	;
	v22 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v24 = base.B2i32(base.Ui32(v22) < base.Ui32(v12))
	v25 = base.B2i32(base.Ui32(v12) < base.Ui32(v22)) - v24
	v26 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v10)+8)))
	if v26 == int32(1) {
		goto L16
	} else {
		goto L17
	}
L15:
	;
	v50 = v11
	v52 = v10
	v54 = v12
	goto L4
L16:
	;
	if base.Ui32(v22) < base.Ui32(v12) {
		v127 = v11
		v129 = v10
		v131 = v12
		goto L3
	} else {
		goto L19
	}
L17:
	;
	v31 = v25
	goto L18
L18:
	;
	if v31 != 0 {
		v43 = v11
		v44 = v31
		v45 = v10
		v47 = v12
		goto L5
	} else {
		goto L20
	}
L19:
	;
	v31 = int32(0) - v25
	goto L18
L20:
	;
	goto L6
L21:
	;
	v35 = *(*int32)(unsafe.Add(mBase, uint32(l3)+4))
	v36 = m.T0[v35].(func(*base.Module, int32, int32, int32) int32)(m, l0, l1, l3)
	mBase = m.M
	v39 = m.ExcPending
	if v39 != 0 {
		goto L22
	} else {
		goto L23
	}
L22:
	;
	return int32(0)
L23:
	;
	v40 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l1)+8)))
	v41 = *(*int32)(unsafe.Add(mBase, uint32(l3)+44))
	v42 = *(*int32)(unsafe.Add(mBase, uint32(l1)+4))
	v43 = v40
	v44 = v36
	v45 = v41
	v47 = v42
	goto L5
L24:
	;
	v50 = v43
	v52 = v45
	v54 = v47
	goto L4
L25:
	;
	v94 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+8)))
	if v94 == int32(1) {
		goto L48
	} else {
		goto L49
	}
L26:
	;
	if v82 < int32(0) {
		v210 = l1
		goto L1
	} else {
		goto L44
	}
L27:
	;
	v78 = *(*int32)(unsafe.Add(mBase, uint32(l3)+48))
	if v78 != 0 {
		v90 = v55
		v91 = v52
		v92 = v56
		goto L25
	} else {
		goto L42
	}
L28:
	;
	if v55&int32(1) != 0 {
		goto L27
	} else {
		goto L31
	}
L29:
	;
	goto L30
L30:
	;
	if v55&int32(1) != 0 {
		goto L33
	} else {
		goto L34
	}
L31:
	;
	v61 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v52)+9)))
	if v61 == int32(0) {
		v90 = v55
		v91 = v52
		v92 = v56
		goto L25
	} else {
		goto L32
	}
L32:
	;
	v210 = l1
	goto L1
L33:
	;
	v66 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v52)+9)))
	if v66 != 0 {
		v90 = v55
		v91 = v52
		v92 = v56
		goto L25
	} else {
		goto L36
	}
L34:
	;
	goto L35
L35:
	;
	v68 = base.B2i32(base.Ui32(v54) < base.Ui32(v56))
	v69 = base.B2i32(base.Ui32(v56) < base.Ui32(v54)) - v68
	v70 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v52)+8)))
	if v70 == int32(1) {
		goto L37
	} else {
		goto L38
	}
L36:
	;
	v210 = l1
	goto L1
L37:
	;
	if base.Ui32(v54) < base.Ui32(v56) {
		v90 = v55
		v91 = v52
		v92 = v56
		goto L25
	} else {
		goto L40
	}
L38:
	;
	v75 = v69
	goto L39
L39:
	;
	if v75 != 0 {
		v82 = v75
		goto L26
	} else {
		goto L41
	}
L40:
	;
	v75 = int32(0) - v69
	goto L39
L41:
	;
	goto L27
L42:
	;
	v79 = *(*int32)(unsafe.Add(mBase, uint32(l3)+4))
	v80 = m.T0[v79].(func(*base.Module, int32, int32, int32) int32)(m, l1, l2, l3)
	mBase = m.M
	v81 = m.ExcPending
	if v81 != 0 {
		goto L22
	} else {
		goto L43
	}
L43:
	;
	v82 = v80
	goto L26
L44:
	;
	v86 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l2)+8)))
	v87 = *(*int32)(unsafe.Add(mBase, uint32(l3)+44))
	v88 = *(*int32)(unsafe.Add(mBase, uint32(l2)+4))
	v90 = v86
	v91 = v87
	v92 = v88
	goto L25
L45:
	;
	return l2
L46:
	;
	if int32(0) <= v122 {
		v210 = l0
		goto L1
	} else {
		goto L64
	}
L47:
	;
	v117 = *(*int32)(unsafe.Add(mBase, uint32(l3)+48))
	if v117 != 0 {
		goto L2
	} else {
		goto L62
	}
L48:
	;
	if v90&int32(1) != 0 {
		goto L47
	} else {
		goto L51
	}
L49:
	;
	goto L50
L50:
	;
	if v90&int32(1) != 0 {
		goto L53
	} else {
		goto L54
	}
L51:
	;
	v99 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v91)+9)))
	if v99 != 0 {
		goto L45
	} else {
		goto L52
	}
L52:
	;
	v210 = l0
	goto L1
L53:
	;
	v102 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v91)+9)))
	if v102 == int32(0) {
		goto L45
	} else {
		goto L56
	}
L54:
	;
	goto L55
L55:
	;
	v105 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v107 = base.B2i32(base.Ui32(v105) < base.Ui32(v92))
	v108 = base.B2i32(base.Ui32(v92) < base.Ui32(v105)) - v107
	v109 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v91)+8)))
	if v109 == int32(1) {
		goto L57
	} else {
		goto L58
	}
L56:
	;
	v210 = l0
	goto L1
L57:
	;
	if base.Ui32(v105) < base.Ui32(v92) {
		goto L2
	} else {
		goto L60
	}
L58:
	;
	v114 = v108
	goto L59
L59:
	;
	if v114 != 0 {
		v122 = v114
		goto L46
	} else {
		goto L61
	}
L60:
	;
	v114 = int32(0) - v108
	goto L59
L61:
	;
	goto L47
L62:
	;
	v118 = *(*int32)(unsafe.Add(mBase, uint32(l3)+4))
	v119 = m.T0[v118].(func(*base.Module, int32, int32, int32) int32)(m, l0, l2, l3)
	mBase = m.M
	v120 = m.ExcPending
	if v120 != 0 {
		goto L22
	} else {
		goto L63
	}
L63:
	;
	v122 = v119
	goto L46
L64:
	;
	goto L45
L65:
	;
	v171 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+8)))
	if v171 == int32(1) {
		goto L87
	} else {
		goto L88
	}
L66:
	;
	if int32(0) < v159 {
		v210 = l1
		goto L1
	} else {
		goto L84
	}
L67:
	;
	v155 = *(*int32)(unsafe.Add(mBase, uint32(l3)+48))
	if v155 != 0 {
		v167 = v132
		v168 = v129
		v169 = v133
		goto L65
	} else {
		goto L82
	}
L68:
	;
	if v132&int32(1) != 0 {
		goto L67
	} else {
		goto L71
	}
L69:
	;
	goto L70
L70:
	;
	if v132&int32(1) != 0 {
		goto L73
	} else {
		goto L74
	}
L71:
	;
	v138 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v129)+9)))
	if v138 != 0 {
		v167 = v132
		v168 = v129
		v169 = v133
		goto L65
	} else {
		goto L72
	}
L72:
	;
	v210 = l1
	goto L1
L73:
	;
	v141 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v129)+9)))
	if v141 == int32(0) {
		v167 = v132
		v168 = v129
		v169 = v133
		goto L65
	} else {
		goto L76
	}
L74:
	;
	goto L75
L75:
	;
	v145 = base.B2i32(base.Ui32(v131) < base.Ui32(v133))
	v146 = base.B2i32(base.Ui32(v133) < base.Ui32(v131)) - v145
	v147 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v129)+8)))
	if v147 == int32(1) {
		goto L77
	} else {
		goto L78
	}
L76:
	;
	v210 = l1
	goto L1
L77:
	;
	if base.Ui32(v131) < base.Ui32(v133) {
		v210 = l1
		goto L1
	} else {
		goto L80
	}
L78:
	;
	v152 = v146
	goto L79
L79:
	;
	if v152 != 0 {
		v159 = v152
		goto L66
	} else {
		goto L81
	}
L80:
	;
	v152 = int32(0) - v146
	goto L79
L81:
	;
	goto L67
L82:
	;
	v156 = *(*int32)(unsafe.Add(mBase, uint32(l3)+4))
	v157 = m.T0[v156].(func(*base.Module, int32, int32, int32) int32)(m, l1, l2, l3)
	mBase = m.M
	v158 = m.ExcPending
	if v158 != 0 {
		goto L22
	} else {
		goto L83
	}
L83:
	;
	v159 = v157
	goto L66
L84:
	;
	v163 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l2)+8)))
	v164 = *(*int32)(unsafe.Add(mBase, uint32(l3)+44))
	v165 = *(*int32)(unsafe.Add(mBase, uint32(l2)+4))
	v167 = v163
	v168 = v164
	v169 = v165
	goto L65
L85:
	;
	if int32(0) <= v201 {
		v210 = l2
		goto L1
	} else {
		goto L107
	}
L86:
	;
	v195 = *(*int32)(unsafe.Add(mBase, uint32(l3)+48))
	if v195 != 0 {
		goto L103
	} else {
		goto L104
	}
L87:
	;
	if v167&int32(1) != 0 {
		goto L86
	} else {
		goto L90
	}
L88:
	;
	goto L89
L89:
	;
	if v167&int32(1) != 0 {
		goto L92
	} else {
		goto L93
	}
L90:
	;
	v176 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v168)+9)))
	if v176 != 0 {
		goto L2
	} else {
		goto L91
	}
L91:
	;
	v210 = l2
	goto L1
L92:
	;
	v179 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v168)+9)))
	if v179 == int32(0) {
		goto L2
	} else {
		goto L95
	}
L93:
	;
	goto L94
L94:
	;
	v182 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v184 = base.B2i32(base.Ui32(v182) < base.Ui32(v169))
	v185 = base.B2i32(base.Ui32(v169) < base.Ui32(v182)) - v184
	v186 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v168)+8)))
	if v186 == int32(1) {
		goto L96
	} else {
		goto L97
	}
L95:
	;
	v210 = l2
	goto L1
L96:
	;
	if base.Ui32(v182) < base.Ui32(v169) {
		goto L99
	} else {
		goto L100
	}
L97:
	;
	v192 = v185
	goto L98
L98:
	;
	if v192 != 0 {
		v201 = v192
		goto L85
	} else {
		goto L102
	}
L99:
	;
	return l2
L100:
	;
	goto L101
L101:
	;
	v192 = int32(0) - v185
	goto L98
L102:
	;
	goto L86
L103:
	;
	return l2
L104:
	;
	goto L105
L105:
	;
	v197 = *(*int32)(unsafe.Add(mBase, uint32(l3)+4))
	v198 = m.T0[v197].(func(*base.Module, int32, int32, int32) int32)(m, l0, l2, l3)
	mBase = m.M
	v199 = m.ExcPending
	if v199 != 0 {
		goto L22
	} else {
		goto L106
	}
L106:
	;
	v201 = v198
	goto L85
L107:
	;
	goto L2
}
