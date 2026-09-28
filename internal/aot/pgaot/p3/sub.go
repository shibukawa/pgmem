package p3

import (
	base "github.com/shibukawa/pgmem/internal/aot/pgaot/base"
	"unsafe"
)

func F_CommitSubTransaction(m *base.Module) {
	mBase := m.M
	_ = mBase
	var v10 int32
	_ = v10
	var v12 int32
	_ = v12
	var v15 int32
	_ = v15
	var v16 int32
	_ = v16
	var v23 int32
	_ = v23
	var v26 int32
	_ = v26
	var v39 int32
	_ = v39
	var v43 int32
	_ = v43
	var v47 int32
	_ = v47
	var v53 int32
	_ = v53
	var v56 int32
	_ = v56
	var v60 int32
	_ = v60
	var v62 int32
	_ = v62
	var v63 int32
	_ = v63
	var v68 int32
	_ = v68
	var v69 int32
	_ = v69
	var v72 int32
	_ = v72
	var v77 int32
	_ = v77
	var v79 int32
	_ = v79
	var v85 int32
	_ = v85
	var v90 int32
	_ = v90
	var v92 int32
	_ = v92
	var v95 int32
	_ = v95
	var v96 int32
	_ = v96
	var v97 int32
	_ = v97
	var v98 int32
	_ = v98
	var v107 int32
	_ = v107
	var v109 int32
	_ = v109
	var v110 int32
	_ = v110
	var v112 int32
	_ = v112
	var v113 int32
	_ = v113
	var v123 int32
	_ = v123
	var v125 int32
	_ = v125
	var v126 int32
	_ = v126
	var v129 int32
	_ = v129
	var v130 int32
	_ = v130
	var v131 int32
	_ = v131
	var v137 int32
	_ = v137
	var v142 int32
	_ = v142
	var v148 int32
	_ = v148
	var v149 int32
	_ = v149
	var v151 int32
	_ = v151
	var v152 int32
	_ = v152
	var v153 int32
	_ = v153
	var v154 int32
	_ = v154
	var v155 int32
	_ = v155
	var v158 int32
	_ = v158
	var v160 int32
	_ = v160
	var v162 int32
	_ = v162
	var v165 int32
	_ = v165
	var v167 int32
	_ = v167
	var v171 int32
	_ = v171
	var v174 int32
	_ = v174
	var v175 int32
	_ = v175
	var v178 int32
	_ = v178
	var v179 int32
	_ = v179
	var v180 int32
	_ = v180
	var v181 int32
	_ = v181
	var v183 int32
	_ = v183
	var v185 int32
	_ = v185
	var v186 int32
	_ = v186
	var v187 int32
	_ = v187
	var v188 int32
	_ = v188
	var v189 int32
	_ = v189
	var v193 int64
	_ = v193
	var v195 int32
	_ = v195
	var v199 int32
	_ = v199
	var v202 int32
	_ = v202
	var v203 int32
	_ = v203
	var v204 int32
	_ = v204
	var v210 int32
	_ = v210
	var v214 int32
	_ = v214
	var v216 int32
	_ = v216
	var v218 int32
	_ = v218
	var v229 int32
	_ = v229
	var v230 int32
	_ = v230
	var v231 int32
	_ = v231
	var v232 int32
	_ = v232
	var v233 int32
	_ = v233
	var v234 int32
	_ = v234
	var v235 int32
	_ = v235
	var v237 int32
	_ = v237
	var v240 int32
	_ = v240
	var v242 int32
	_ = v242
	var v244 int32
	_ = v244
	var v245 int32
	_ = v245
	var v246 int32
	_ = v246
	var v247 int32
	_ = v247
	var v256 int32
	_ = v256
	var v257 int32
	_ = v257
	var v261 int32
	_ = v261
	var v266 int32
	_ = v266
	var v269 int32
	_ = v269
	var v271 int32
	_ = v271
	var v276 int32
	_ = v276
	var v279 int32
	_ = v279
	var v281 int32
	_ = v281
	var v288 int32
	_ = v288
	var v291 int32
	_ = v291
	var v296 int32
	_ = v296
	var v301 int32
	_ = v301
	var v302 int32
	_ = v302
	var v316 int32
	_ = v316
	var v317 int32
	_ = v317
	var v318 int32
	_ = v318
	var v320 int32
	_ = v320
	var v322 int32
	_ = v322
	var v323 int32
	_ = v323
	var v325 int32
	_ = v325
	var v328 int32
	_ = v328
	var v330 int32
	_ = v330
	var v331 int32
	_ = v331
	var v340 int32
	_ = v340
	var v341 int32
	_ = v341
	var v342 int32
	_ = v342
	var v343 int32
	_ = v343
	var v345 int32
	_ = v345
	var v348 int32
	_ = v348
	var v352 int32
	_ = v352
	var v355 int32
	_ = v355
	var v357 int32
	_ = v357
	var v358 int32
	_ = v358
	var v365 int32
	_ = v365
	var v368 int32
	_ = v368
	var v371 int32
	_ = v371
	var v375 int32
	_ = v375
	var v383 int32
	_ = v383
	var v387 int32
	_ = v387
	var v388 int32
	_ = v388
	var v389 int32
	_ = v389
	var v393 int32
	_ = v393
	var v395 int32
	_ = v395
	var v396 int32
	_ = v396
	var v408 int32
	_ = v408
	var v419 int32
	_ = v419
	var v420 int32
	_ = v420
	var v421 int32
	_ = v421
	var v422 int32
	_ = v422
	var v423 int32
	_ = v423
	var v432 int32
	_ = v432
	var v434 int32
	_ = v434
	var v435 int32
	_ = v435
	var v437 int32
	_ = v437
	var v447 int32
	_ = v447
	var v448 int32
	_ = v448
	var v452 int32
	_ = v452
	var v454 int32
	_ = v454
	var v455 int32
	_ = v455
	var v456 int32
	_ = v456
	var v458 int32
	_ = v458
	var v460 int32
	_ = v460
	var v463 int32
	_ = v463
	var v465 int32
	_ = v465
	var v466 int32
	_ = v466
	var v468 int32
	_ = v468
	var v471 int32
	_ = v471
	var v480 int32
	_ = v480
	var v483 int32
	_ = v483
	var v494 int32
	_ = v494
	var v496 int32
	_ = v496
	var v497 int32
	_ = v497
	var v499 int32
	_ = v499
	var v508 int32
	_ = v508
	var v509 int32
	_ = v509
	var v513 int32
	_ = v513
	var v515 int32
	_ = v515
	var v520 int32
	_ = v520
	var v521 int32
	_ = v521
	var v526 int32
	_ = v526
	var v528 int32
	_ = v528
	var v530 int32
	_ = v530
	var v532 int32
	_ = v532
	var v534 int32
	_ = v534
	var v536 int32
	_ = v536
	var v537 int32
	_ = v537
	var v538 int32
	_ = v538
	var v540 int32
	_ = v540
	var v542 int32
	_ = v542
	var v543 int32
	_ = v543
	var v544 int32
	_ = v544
	var v546 int32
	_ = v546
	var v570 int32
	_ = v570
	var v571 int32
	_ = v571
	var v572 int32
	_ = v572
	var v574 int32
	_ = v574
	var v576 int32
	_ = v576
	var v578 int32
	_ = v578
	var v580 int32
	_ = v580
	var v582 int32
	_ = v582
	var v583 int32
	_ = v583
	var v585 int32
	_ = v585
	var v590 int32
	_ = v590
	var v599 int32
	_ = v599
	var v602 int32
	_ = v602
	var v613 int32
	_ = v613
	var v616 int32
	_ = v616
	var v617 int32
	_ = v617
	var v621 int32
	_ = v621
	var v623 int32
	_ = v623
	var v628 int32
	_ = v628
	var v629 int32
	_ = v629
	var v630 int32
	_ = v630
	var v634 int32
	_ = v634
	var v635 int32
	_ = v635
	var v637 int32
	_ = v637
	var v638 int32
	_ = v638
	var v639 int32
	_ = v639
	var v640 int32
	_ = v640
	var v641 int32
	_ = v641
	var v642 int32
	_ = v642
	var v644 int32
	_ = v644
	var v650 int32
	_ = v650
	var v657 int32
	_ = v657
	var v660 int32
	_ = v660
	var v665 int32
	_ = v665
	var v670 int32
	_ = v670
	v10 = m.G0
	v12 = v10 - int32(48)
	m.G0 = v12
	v15 = *(*int32)(unsafe.Add(mBase, _c_F_CommitSubTransaction[0]))
	v16 = int32(10)
	goto L3
L1:
	;
	if v56 != 0 {
		goto L14
	} else {
		goto L15
	}
L2:
	;
	goto L1
L3:
	;
	v23 = *(*int32)(unsafe.Add(mBase, _c_F_CommitSubTransaction[1]))
	v26 = *(*int32)(unsafe.Add(mBase, uint32(v23<<(uint(int32(2))%32))+uint32(_c_F_CommitSubTransaction[2])))
	goto L6
L4:
	;
	v39 = int32(0)
	goto L11
L6:
	;
	goto L7
L7:
	;
	if int32(0)|base.B2i32(v26 == int32(15)) != 0 {
		goto L4
	} else {
		goto L9
	}
L9:
	;
	if v26 <= v16 {
		v56 = int32(1)
		goto L2
	} else {
		goto L10
	}
L10:
	;
	goto L4
L11:
	;
	v43 = *(*int32)(unsafe.Add(mBase, _c_F_CommitSubTransaction[3]))
	if v43 != int32(2) {
		v56 = v39
		goto L2
	} else {
		goto L12
	}
L12:
	;
	v47 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_CommitSubTransaction[4])))
	if v47&int32(1) != 0 {
		v56 = v39
		goto L2
	} else {
		goto L13
	}
L13:
	;
	v53 = *(*int32)(unsafe.Add(mBase, _c_F_CommitSubTransaction[5]))
	v56 = int32(0) | base.B2i32(v53 <= v16)
	goto L2
L14:
	;
	v60 = *(*int32)(unsafe.Add(mBase, _c_F_CommitSubTransaction[0]))
	F_ShowTransactionStateRec(m, int32(_a_F_CommitSubTransaction_0), v60)
	mBase = m.M
	v62 = m.ExcPending
	if v62 != 0 {
		goto L17
	} else {
		goto L18
	}
L15:
	;
	goto L16
L16:
	;
	v63 = *(*int32)(unsafe.Add(mBase, uint32(v15)+20))
	if v63 == int32(2) {
		goto L19
	} else {
		goto L20
	}
L17:
	;
	return
L18:
	;
	goto L16
L19:
	;
	v92 = *(*int32)(unsafe.Add(mBase, uint32(v15)+8))
	v95 = *(*int32)(unsafe.Add(mBase, _c_F_CommitSubTransaction[6]))
	if v95 != 0 {
		goto L28
	} else {
		goto L29
	}
L20:
	;
	v68 = F_errstart(m, int32(19), int32(0))
	mBase = m.M
	v69 = m.ExcPending
	if v69 != 0 {
		goto L17
	} else {
		goto L21
	}
L21:
	;
	if v68 == int32(0) {
		goto L19
	} else {
		goto L22
	}
L22:
	;
	v72 = *(*int32)(unsafe.Add(mBase, uint32(v15)+20))
	if base.Ui32(v72) <= base.Ui32(int32(5)) {
		goto L23
	} else {
		goto L24
	}
L23:
	;
	v77 = *(*int32)(unsafe.Add(mBase, uint32(v72<<(uint(int32(2))%32))+uint32(_c_F_CommitSubTransaction[7])))
	v79 = v77
	goto L25
L24:
	;
	v79 = int32(_a_F_CommitSubTransaction_1)
	goto L25
L25:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v12)+32)) = v79
	F_errmsg_internal(m, int32(_a_F_CommitSubTransaction_2), v12+int32(32))
	mBase = m.M
	v85 = m.ExcPending
	if v85 != 0 {
		goto L17
	} else {
		goto L26
	}
L26:
	;
	F_errfinish(m, int32(_a_F_CommitSubTransaction_3), int32(_a_F_CommitSubTransaction_4), int32(_a_F_CommitSubTransaction_0))
	mBase = m.M
	v90 = m.ExcPending
	if v90 != 0 {
		goto L17
	} else {
		goto L27
	}
L27:
	;
	goto L19
L28:
	;
	v96 = *(*int32)(unsafe.Add(mBase, uint32(v15)+80))
	v97 = *(*int32)(unsafe.Add(mBase, uint32(v96)+8))
	v98 = v95
	goto L31
L29:
	;
	v123 = v92
	goto L30
L30:
	;
	F_AtEOSubXact_Parallel(m, int32(1), v123)
	mBase = m.M
	v125 = m.ExcPending
	if v125 != 0 {
		goto L17
	} else {
		goto L35
	}
L31:
	;
	v107 = *(*int32)(unsafe.Add(mBase, uint32(v98)))
	v109 = *(*int32)(unsafe.Add(mBase, uint32(v98)+8))
	v110 = *(*int32)(unsafe.Add(mBase, uint32(v98)+4))
	m.T0[v110].(func(*base.Module, int32, int32, int32, int32))(m, int32(3), v92, v97, v109)
	mBase = m.M
	v112 = m.ExcPending
	if v112 != 0 {
		goto L17
	} else {
		goto L33
	}
L32:
	;
	v113 = *(*int32)(unsafe.Add(mBase, uint32(v15)+8))
	v123 = v113
	goto L30
L33:
	;
	if v107 != 0 {
		v98 = v107
		goto L31
	} else {
		goto L34
	}
L34:
	;
	goto L32
L35:
	;
	v126 = *(*int32)(unsafe.Add(mBase, uint32(v15)+72))
	if v126 != 0 {
		goto L36
	} else {
		goto L37
	}
L36:
	;
	v129 = F_errstart(m, int32(19), int32(0))
	mBase = m.M
	v130 = m.ExcPending
	if v130 != 0 {
		goto L17
	} else {
		goto L39
	}
L37:
	;
	goto L38
L38:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v15)+20)) = int32(3)
	F_CommandCounterIncrement(m)
	mBase = m.M
	v148 = m.ExcPending
	if v148 != 0 {
		goto L17
	} else {
		goto L45
	}
L39:
	;
	if v129 != 0 {
		goto L40
	} else {
		goto L41
	}
L40:
	;
	v131 = *(*int32)(unsafe.Add(mBase, uint32(v15)+72))
	*(*int32)(unsafe.Add(mBase, uint32(v12)+16)) = v131
	F_errmsg_internal(m, int32(_a_F_CommitSubTransaction_5), v12+int32(16))
	mBase = m.M
	v137 = m.ExcPending
	if v137 != 0 {
		goto L17
	} else {
		goto L43
	}
L41:
	;
	goto L42
L42:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v15)+72)) = int32(0)
	goto L38
L43:
	;
	F_errfinish(m, int32(_a_F_CommitSubTransaction_3), int32(_a_F_CommitSubTransaction_6), int32(_a_F_CommitSubTransaction_0))
	mBase = m.M
	v142 = m.ExcPending
	if v142 != 0 {
		goto L17
	} else {
		goto L44
	}
L44:
	;
	goto L42
L45:
	;
	v149 = *(*int32)(unsafe.Add(mBase, uint32(v15)))
	if v149 != 0 {
		goto L47
	} else {
		goto L48
	}
L46:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v657 = m.ExcPending
	if v657 != 0 {
		goto L17
	} else {
		goto L187
	}
L47:
	;
	v151 = *(*int32)(unsafe.Add(mBase, _c_F_CommitSubTransaction[0]))
	v152 = *(*int32)(unsafe.Add(mBase, uint32(v151)+80))
	v153 = *(*int32)(unsafe.Add(mBase, uint32(v152)+56))
	v154 = *(*int32)(unsafe.Add(mBase, uint32(v152)+52))
	v155 = *(*int32)(unsafe.Add(mBase, uint32(v151)+52))
	v158 = v154 + v155 + int32(1)
	if v153 < v158 {
		goto L50
	} else {
		goto L51
	}
L48:
	;
	goto L49
L49:
	;
	F_AfterTriggerEndSubXact(m, int32(1))
	mBase = m.M
	v229 = m.ExcPending
	if v229 != 0 {
		goto L17
	} else {
		goto L70
	}
L50:
	;
	v160 = int32(268435455)
	v162 = v158 << (uint(int32(1)) % 32)
	if v160 <= v162 {
		goto L53
	} else {
		goto L54
	}
L51:
	;
	v187 = v154
	v188 = v152
	goto L52
L52:
	;
	v189 = *(*int32)(unsafe.Add(mBase, uint32(v188)+48))
	v193 = *(*int64)(unsafe.Add(mBase, uint32(v151)))
	*(*uint32)(unsafe.Add(mBase, uint32(v189+v187<<(uint(int32(2))%32)))) = uint32(v193)
	v195 = *(*int32)(unsafe.Add(mBase, uint32(v151)+52))
	if v195 <= int32(0) {
		goto L63
	} else {
		goto L64
	}
L53:
	;
	v165 = v160
	goto L55
L54:
	;
	v165 = v162
	goto L55
L55:
	;
	if v165 < v158 {
		goto L46
	} else {
		goto L56
	}
L56:
	;
	v167 = *(*int32)(unsafe.Add(mBase, uint32(v152)+48))
	if v167 == int32(0) {
		goto L58
	} else {
		goto L59
	}
L57:
	;
	v181 = *(*int32)(unsafe.Add(mBase, uint32(v151)+80))
	*(*int32)(unsafe.Add(mBase, uint32(v181)+48)) = v180
	v183 = *(*int32)(unsafe.Add(mBase, uint32(v151)+80))
	*(*int32)(unsafe.Add(mBase, uint32(v183)+56)) = v165
	v185 = *(*int32)(unsafe.Add(mBase, uint32(v151)+80))
	v186 = *(*int32)(unsafe.Add(mBase, uint32(v185)+52))
	v187 = v186
	v188 = v185
	goto L52
L58:
	;
	v171 = *(*int32)(unsafe.Add(mBase, _c_F_CommitSubTransaction[8]))
	v174 = F_MemoryContextAlloc(m, v171, v165<<(uint(int32(2))%32))
	mBase = m.M
	v175 = m.ExcPending
	if v175 != 0 {
		goto L17
	} else {
		goto L61
	}
L59:
	;
	goto L60
L60:
	;
	v178 = F_repalloc(m, v167, v165<<(uint(int32(2))%32))
	mBase = m.M
	v179 = m.ExcPending
	if v179 != 0 {
		goto L17
	} else {
		goto L62
	}
L61:
	;
	v180 = v174
	goto L57
L62:
	;
	v180 = v178
	goto L57
L63:
	;
	v214 = *(*int32)(unsafe.Add(mBase, uint32(v151)+80))
	*(*int32)(unsafe.Add(mBase, uint32(v214)+52)) = v158
	v216 = *(*int32)(unsafe.Add(mBase, uint32(v151)+48))
	if v216 != 0 {
		goto L66
	} else {
		goto L67
	}
L64:
	;
	v199 = v195 << (uint(int32(2)) % 32)
	if v199 == int32(0) {
		goto L63
	} else {
		goto L65
	}
L65:
	;
	v202 = *(*int32)(unsafe.Add(mBase, uint32(v151)+80))
	v203 = *(*int32)(unsafe.Add(mBase, uint32(v202)+48))
	v204 = *(*int32)(unsafe.Add(mBase, uint32(v202)+52))
	v210 = *(*int32)(unsafe.Add(mBase, uint32(v151)+48))
	base.MemoryCopy(m, v203+v204<<(uint(int32(2))%32)+int32(4), v210, v199)
	goto L63
L66:
	;
	F_pfree(m, v216)
	mBase = m.M
	v218 = m.ExcPending
	if v218 != 0 {
		goto L17
	} else {
		goto L69
	}
L67:
	;
	goto L68
L68:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v151)+56)) = int32(0)
	*(*int64)(unsafe.Add(mBase, uint32(v151)+48)) = int64(0)
	goto L49
L69:
	;
	goto L68
L70:
	;
	v230 = *(*int32)(unsafe.Add(mBase, uint32(v15)+8))
	v231 = *(*int32)(unsafe.Add(mBase, uint32(v15)+80))
	v232 = *(*int32)(unsafe.Add(mBase, uint32(v231)+8))
	v233 = *(*int32)(unsafe.Add(mBase, uint32(v231)+28))
	v234 = *(*int32)(unsafe.Add(mBase, uint32(v231)+40))
	v235 = m.G0
	v237 = v235 - int32(32)
	m.G0 = v237
	v240 = v237 + int32(12)
	v242 = *(*int32)(unsafe.Add(mBase, _c_F_CommitSubTransaction[9]))
	F_hash_seq_init(m, v240, v242)
	mBase = m.M
	v244 = m.ExcPending
	if v244 != 0 {
		goto L17
	} else {
		goto L71
	}
L71:
	;
	v245 = F_hash_seq_search(m, v240)
	mBase = m.M
	v246 = m.ExcPending
	if v246 != 0 {
		goto L17
	} else {
		goto L72
	}
L72:
	;
	if v245 != 0 {
		goto L73
	} else {
		goto L74
	}
L73:
	;
	v247 = v245
	goto L76
L74:
	;
	goto L75
L75:
	;
	m.G0 = v237 + int32(32)
	v316 = *(*int32)(unsafe.Add(mBase, uint32(v15)+8))
	v317 = *(*int32)(unsafe.Add(mBase, uint32(v15)+80))
	v318 = *(*int32)(unsafe.Add(mBase, uint32(v317)+8))
	F_AtEOSubXact_LargeObject(m, int32(1), v316, v318)
	mBase = m.M
	v320 = m.ExcPending
	if v320 != 0 {
		goto L17
	} else {
		goto L99
	}
L76:
	;
	v256 = *(*int32)(unsafe.Add(mBase, uint32(v247)+64))
	v257 = *(*int32)(unsafe.Add(mBase, uint32(v256)+20))
	if v257 != v230 {
		goto L78
	} else {
		goto L79
	}
L77:
	;
	goto L75
L78:
	;
	v296 = *(*int32)(unsafe.Add(mBase, uint32(v256)+24))
	if v230 == v296 {
		goto L94
	} else {
		goto L95
	}
L79:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v256)+28)) = v233
	*(*int32)(unsafe.Add(mBase, uint32(v256)+20)) = v232
	v261 = *(*int32)(unsafe.Add(mBase, uint32(v256)+12))
	if v261 == int32(0) {
		goto L78
	} else {
		goto L80
	}
L80:
	;
	v266 = *(*int32)(unsafe.Add(mBase, uint32(v261)))
	if v266 == int32(0) {
		goto L82
	} else {
		goto L83
	}
L81:
	;
	goto L78
L82:
	;
	if v234 != 0 {
		goto L91
	} else {
		goto L92
	}
L83:
	;
	v269 = *(*int32)(unsafe.Add(mBase, uint32(v266)+4))
	if v269 == v261 {
		goto L84
	} else {
		goto L85
	}
L84:
	;
	v271 = *(*int32)(unsafe.Add(mBase, uint32(v261)+8))
	*(*int32)(unsafe.Add(mBase, uint32(v266)+4)) = v271
	goto L82
L85:
	;
	goto L86
L86:
	;
	v276 = v269
	goto L87
L87:
	;
	if v276 == int32(0) {
		goto L82
	} else {
		goto L89
	}
L88:
	;
	v281 = *(*int32)(unsafe.Add(mBase, uint32(v261)+8))
	*(*int32)(unsafe.Add(mBase, uint32(v276)+8)) = v281
	goto L82
L89:
	;
	v279 = *(*int32)(unsafe.Add(mBase, uint32(v276)+8))
	if v261 != v279 {
		v276 = v279
		goto L87
	} else {
		goto L90
	}
L90:
	;
	goto L88
L91:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v261))) = v234
	v288 = *(*int32)(unsafe.Add(mBase, uint32(v234)+4))
	*(*int32)(unsafe.Add(mBase, uint32(v261)+8)) = v288
	*(*int32)(unsafe.Add(mBase, uint32(v234)+4)) = v261
	goto L81
L92:
	;
	goto L93
L93:
	;
	v291 = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v261)+8)) = v291
	*(*int32)(unsafe.Add(mBase, uint32(v261))) = v291
	goto L81
L94:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v256)+24)) = v232
	goto L96
L95:
	;
	goto L96
L96:
	;
	v301 = F_hash_seq_search(m, v237+int32(12))
	mBase = m.M
	v302 = m.ExcPending
	if v302 != 0 {
		goto L17
	} else {
		goto L97
	}
L97:
	;
	if v301 != 0 {
		v247 = v301
		goto L76
	} else {
		goto L98
	}
L98:
	;
	goto L77
L99:
	;
	v322 = *(*int32)(unsafe.Add(mBase, _c_F_CommitSubTransaction[0]))
	v323 = *(*int32)(unsafe.Add(mBase, uint32(v322)+28))
	goto L100
L100:
	;
	v325 = *(*int32)(unsafe.Add(mBase, _c_F_CommitSubTransaction[10]))
	if v325 == int32(0) {
		goto L101
	} else {
		goto L102
	}
L101:
	;
	v352 = *(*int32)(unsafe.Add(mBase, _c_F_CommitSubTransaction[11]))
	if v352 == int32(0) {
		goto L111
	} else {
		goto L112
	}
L102:
	;
	v328 = *(*int32)(unsafe.Add(mBase, uint32(v325)))
	if v328 < v323 {
		goto L101
	} else {
		goto L103
	}
L103:
	;
	v330 = *(*int32)(unsafe.Add(mBase, uint32(v325)+8))
	if v330 != 0 {
		goto L105
	} else {
		goto L106
	}
L104:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_CommitSubTransaction[10])) = v330
	v340 = *(*int32)(unsafe.Add(mBase, uint32(v330)+4))
	v341 = *(*int32)(unsafe.Add(mBase, uint32(v325)+4))
	v342 = F_list_concat(m, v340, v341)
	mBase = m.M
	v343 = m.ExcPending
	if v343 != 0 {
		goto L17
	} else {
		goto L109
	}
L105:
	;
	v331 = *(*int32)(unsafe.Add(mBase, uint32(v330)))
	if v323-int32(1) <= v331 {
		goto L104
	} else {
		goto L108
	}
L106:
	;
	goto L107
L107:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v325))) = v328 - int32(1)
	goto L101
L108:
	;
	goto L107
L109:
	;
	v345 = *(*int32)(unsafe.Add(mBase, _c_F_CommitSubTransaction[10]))
	*(*int32)(unsafe.Add(mBase, uint32(v345)+4)) = v342
	F_pfree(m, v325)
	mBase = m.M
	v348 = m.ExcPending
	if v348 != 0 {
		goto L17
	} else {
		goto L110
	}
L110:
	;
	goto L101
L111:
	;
	v419 = *(*int32)(unsafe.Add(mBase, _c_F_CommitSubTransaction[6]))
	if v419 != 0 {
		goto L131
	} else {
		goto L132
	}
L112:
	;
	v355 = *(*int32)(unsafe.Add(mBase, uint32(v352)))
	if v355 < v323 {
		goto L111
	} else {
		goto L113
	}
L113:
	;
	v357 = *(*int32)(unsafe.Add(mBase, uint32(v352)+20))
	if v357 != 0 {
		goto L115
	} else {
		goto L116
	}
L114:
	;
	v365 = int32(0)
	*(*int32)(unsafe.Add(mBase, _c_F_CommitSubTransaction[11])) = v357
	v368 = *(*int32)(unsafe.Add(mBase, uint32(v352)+4))
	if v368 == v365 {
		goto L119
	} else {
		goto L120
	}
L115:
	;
	v358 = *(*int32)(unsafe.Add(mBase, uint32(v357)))
	if v323-int32(1) <= v358 {
		goto L114
	} else {
		goto L118
	}
L116:
	;
	goto L117
L117:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v352))) = v355 - int32(1)
	goto L111
L118:
	;
	goto L117
L119:
	;
	F_pfree(m, v352)
	mBase = m.M
	v408 = m.ExcPending
	if v408 != 0 {
		goto L17
	} else {
		goto L130
	}
L120:
	;
	v371 = *(*int32)(unsafe.Add(mBase, uint32(v368)+4))
	if v371 <= int32(0) {
		goto L119
	} else {
		goto L121
	}
L121:
	;
	v375 = v365
	goto L122
L122:
	;
	v383 = *(*int32)(unsafe.Add(mBase, uint32(v368)+12))
	v387 = *(*int32)(unsafe.Add(mBase, uint32(v383+v375<<(uint(int32(2))%32))))
	v388 = F_AsyncExistsPendingNotify(m, v387)
	mBase = m.M
	v389 = m.ExcPending
	if v389 != 0 {
		goto L17
	} else {
		goto L124
	}
L123:
	;
	goto L119
L124:
	;
	if v388 == int32(0) {
		goto L125
	} else {
		goto L126
	}
L125:
	;
	F_AddEventToPendingNotifies(m, v387)
	mBase = m.M
	v393 = m.ExcPending
	if v393 != 0 {
		goto L17
	} else {
		goto L128
	}
L126:
	;
	goto L127
L127:
	;
	v395 = v375 + int32(1)
	v396 = *(*int32)(unsafe.Add(mBase, uint32(v368)+4))
	if v395 < v396 {
		v375 = v395
		goto L122
	} else {
		goto L129
	}
L128:
	;
	goto L127
L129:
	;
	goto L123
L130:
	;
	goto L111
L131:
	;
	v420 = *(*int32)(unsafe.Add(mBase, uint32(v15)+8))
	v421 = *(*int32)(unsafe.Add(mBase, uint32(v15)+80))
	v422 = *(*int32)(unsafe.Add(mBase, uint32(v421)+8))
	v423 = v419
	goto L134
L132:
	;
	goto L133
L133:
	;
	v447 = *(*int32)(unsafe.Add(mBase, uint32(v15)+40))
	v448 = int32(1)
	F_ResourceOwnerReleaseInternal(m, v447, v448, v448, int32(0))
	mBase = m.M
	v452 = m.ExcPending
	if v452 != 0 {
		goto L17
	} else {
		goto L138
	}
L134:
	;
	v432 = *(*int32)(unsafe.Add(mBase, uint32(v423)))
	v434 = *(*int32)(unsafe.Add(mBase, uint32(v423)+8))
	v435 = *(*int32)(unsafe.Add(mBase, uint32(v423)+4))
	m.T0[v435].(func(*base.Module, int32, int32, int32, int32))(m, int32(1), v420, v422, v434)
	mBase = m.M
	v437 = m.ExcPending
	if v437 != 0 {
		goto L17
	} else {
		goto L136
	}
L135:
	;
	goto L133
L136:
	;
	if v432 != 0 {
		v423 = v432
		goto L134
	} else {
		goto L137
	}
L137:
	;
	goto L135
L138:
	;
	v454 = *(*int32)(unsafe.Add(mBase, uint32(v15)+8))
	v455 = *(*int32)(unsafe.Add(mBase, uint32(v15)+80))
	v456 = *(*int32)(unsafe.Add(mBase, uint32(v455)+8))
	F_AtEOSubXact_RelationCache(m, int32(1), v454, v456)
	mBase = m.M
	v458 = m.ExcPending
	if v458 != 0 {
		goto L17
	} else {
		goto L139
	}
L139:
	;
	F_AtEOXact_TypeCache(m)
	mBase = m.M
	v460 = m.ExcPending
	if v460 != 0 {
		goto L17
	} else {
		goto L140
	}
L140:
	;
	F_AtEOSubXact_Inval(m, int32(1))
	mBase = m.M
	v463 = m.ExcPending
	if v463 != 0 {
		goto L17
	} else {
		goto L141
	}
L141:
	;
	v465 = *(*int32)(unsafe.Add(mBase, _c_F_CommitSubTransaction[0]))
	v466 = *(*int32)(unsafe.Add(mBase, uint32(v465)+28))
	goto L142
L142:
	;
	v468 = *(*int32)(unsafe.Add(mBase, _c_F_CommitSubTransaction[12]))
	if v468 != 0 {
		goto L143
	} else {
		goto L144
	}
L143:
	;
	v471 = v468
	goto L146
L144:
	;
	goto L145
L145:
	;
	v494 = *(*int32)(unsafe.Add(mBase, uint32(v15)+40))
	*(*int32)(unsafe.Add(mBase, _c_F_CommitSubTransaction[13])) = v494
	v496 = *(*int32)(unsafe.Add(mBase, uint32(v15)))
	if v496 != 0 {
		goto L152
	} else {
		goto L153
	}
L146:
	;
	v480 = *(*int32)(unsafe.Add(mBase, uint32(v471)+20))
	if v466 <= v480 {
		goto L148
	} else {
		goto L149
	}
L147:
	;
	goto L145
L148:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v471)+20)) = v466 - int32(1)
	goto L150
L149:
	;
	goto L150
L150:
	;
	v483 = *(*int32)(unsafe.Add(mBase, uint32(v471)+24))
	if v483 != 0 {
		v471 = v483
		goto L146
	} else {
		goto L151
	}
L151:
	;
	goto L147
L152:
	;
	v497 = m.G0
	v499 = v497 - int32(16)
	m.G0 = v499
	*(*int32)(unsafe.Add(mBase, uint32(v499)+12)) = int32(17104896)
	*(*int64)(unsafe.Add(mBase, uint32(v499)+4)) = int64(0)
	*(*int32)(unsafe.Add(mBase, uint32(v499))) = v496
	v508 = F_LockRelease(m, v499, int32(7), int32(0))
	mBase = m.M
	v509 = m.ExcPending
	if v509 != 0 {
		goto L17
	} else {
		goto L155
	}
L153:
	;
	v515 = v494
	goto L154
L154:
	;
	F_ResourceOwnerReleaseInternal(m, v515, int32(2), int32(1), int32(0))
	mBase = m.M
	v520 = m.ExcPending
	if v520 != 0 {
		goto L17
	} else {
		goto L156
	}
L155:
	;
	m.G0 = v499 + int32(16)
	v513 = *(*int32)(unsafe.Add(mBase, uint32(v15)+40))
	v515 = v513
	goto L154
L156:
	;
	v521 = *(*int32)(unsafe.Add(mBase, uint32(v15)+40))
	F_ResourceOwnerReleaseInternal(m, v521, int32(3), int32(1), int32(0))
	mBase = m.M
	v526 = m.ExcPending
	if v526 != 0 {
		goto L17
	} else {
		goto L157
	}
L157:
	;
	v528 = *(*int32)(unsafe.Add(mBase, uint32(v15)+32))
	F_AtEOXact_GUC(m, int32(1), v528)
	mBase = m.M
	v530 = m.ExcPending
	if v530 != 0 {
		goto L17
	} else {
		goto L158
	}
L158:
	;
	v532 = *(*int32)(unsafe.Add(mBase, uint32(v15)+8))
	F_AtEOSubXact_SPI(m, int32(1), v532)
	mBase = m.M
	v534 = m.ExcPending
	if v534 != 0 {
		goto L17
	} else {
		goto L159
	}
L159:
	;
	v536 = *(*int32)(unsafe.Add(mBase, uint32(v15)+8))
	v537 = *(*int32)(unsafe.Add(mBase, uint32(v15)+80))
	v538 = *(*int32)(unsafe.Add(mBase, uint32(v537)+8))
	F_AtEOSubXact_on_commit_actions(m, int32(1), v536, v538)
	mBase = m.M
	v540 = m.ExcPending
	if v540 != 0 {
		goto L17
	} else {
		goto L160
	}
L160:
	;
	v542 = *(*int32)(unsafe.Add(mBase, uint32(v15)+8))
	v543 = *(*int32)(unsafe.Add(mBase, uint32(v15)+80))
	v544 = *(*int32)(unsafe.Add(mBase, uint32(v543)+8))
	v546 = *(*int32)(unsafe.Add(mBase, _c_F_CommitSubTransaction[14]))
	if v542 == v546 {
		goto L162
	} else {
		goto L163
	}
L161:
	;
	v570 = *(*int32)(unsafe.Add(mBase, uint32(v15)+8))
	v571 = *(*int32)(unsafe.Add(mBase, uint32(v15)+80))
	v572 = *(*int32)(unsafe.Add(mBase, uint32(v571)+8))
	F_AtEOSubXact_Files(m, int32(1), v570, v572)
	mBase = m.M
	v574 = m.ExcPending
	if v574 != 0 {
		goto L17
	} else {
		goto L168
	}
L162:
	;
	goto L165
L163:
	;
	goto L164
L164:
	;
	goto L161
L165:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_CommitSubTransaction[14])) = v544
	goto L161
L168:
	;
	v576 = *(*int32)(unsafe.Add(mBase, uint32(v15)+28))
	F_AtEOSubXact_HashTables(m, int32(1), v576)
	mBase = m.M
	v578 = m.ExcPending
	if v578 != 0 {
		goto L17
	} else {
		goto L169
	}
L169:
	;
	v580 = *(*int32)(unsafe.Add(mBase, uint32(v15)+28))
	F_AtEOSubXact_PgStat(m, int32(1), v580)
	mBase = m.M
	v582 = m.ExcPending
	if v582 != 0 {
		goto L17
	} else {
		goto L170
	}
L170:
	;
	v583 = *(*int32)(unsafe.Add(mBase, uint32(v15)+28))
	v585 = *(*int32)(unsafe.Add(mBase, _c_F_CommitSubTransaction[15]))
	if v585 == int32(0) {
		goto L171
	} else {
		goto L172
	}
L171:
	;
	v613 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v15)+68)))
	*(*uint8)(unsafe.Add(mBase, _c_F_CommitSubTransaction[16])) = uint8(v613)
	v616 = *(*int32)(unsafe.Add(mBase, uint32(v15)+80))
	v617 = *(*int32)(unsafe.Add(mBase, uint32(v616)+40))
	*(*int32)(unsafe.Add(mBase, _c_F_CommitSubTransaction[17])) = v617
	*(*int32)(unsafe.Add(mBase, _c_F_CommitSubTransaction[13])) = v617
	v621 = *(*int32)(unsafe.Add(mBase, uint32(v15)+40))
	F_ResourceOwnerDelete(m, v621)
	mBase = m.M
	v623 = m.ExcPending
	if v623 != 0 {
		goto L17
	} else {
		goto L177
	}
L172:
	;
	v590 = v585
	goto L173
L173:
	;
	v599 = *(*int32)(unsafe.Add(mBase, uint32(v590)+4))
	if v599 < v583 {
		goto L171
	} else {
		goto L175
	}
L174:
	;
	goto L171
L175:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v590)+4)) = v583 - int32(1)
	v602 = *(*int32)(unsafe.Add(mBase, uint32(v590)+8))
	if v602 != 0 {
		v590 = v602
		goto L173
	} else {
		goto L176
	}
L176:
	;
	goto L174
L177:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v15)+40)) = int32(0)
	v628 = *(*int32)(unsafe.Add(mBase, _c_F_CommitSubTransaction[0]))
	v629 = *(*int32)(unsafe.Add(mBase, uint32(v628)+80))
	v630 = *(*int32)(unsafe.Add(mBase, uint32(v629)+36))
	*(*int32)(unsafe.Add(mBase, _c_F_CommitSubTransaction[18])) = v630
	*(*int32)(unsafe.Add(mBase, _c_F_CommitSubTransaction[19])) = v630
	v634 = *(*int32)(unsafe.Add(mBase, uint32(v628)+36))
	v635 = *(*int32)(unsafe.Add(mBase, uint32(v634)+20))
	if v635 != 0 {
		goto L178
	} else {
		goto L179
	}
L178:
	;
	v641 = int32(0)
	goto L180
L179:
	;
	v637 = *(*int32)(unsafe.Add(mBase, uint32(v634)+12))
	v638 = *(*int32)(unsafe.Add(mBase, uint32(v637)+28))
	v639 = m.T0[v638].(func(*base.Module, int32) int32)(m, v634)
	mBase = m.M
	v640 = m.ExcPending
	if v640 != 0 {
		goto L17
	} else {
		goto L181
	}
L180:
	;
	if v641 != 0 {
		goto L182
	} else {
		goto L183
	}
L181:
	;
	v641 = v639
	goto L180
L182:
	;
	v642 = *(*int32)(unsafe.Add(mBase, uint32(v628)+36))
	F_MemoryContextDelete(m, v642)
	mBase = m.M
	v644 = m.ExcPending
	if v644 != 0 {
		goto L17
	} else {
		goto L185
	}
L183:
	;
	goto L184
L184:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v15)+20)) = int32(0)
	F_PopTransaction(m)
	mBase = m.M
	v650 = m.ExcPending
	if v650 != 0 {
		goto L17
	} else {
		goto L186
	}
L185:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v628)+36)) = int32(0)
	goto L184
L186:
	;
	m.G0 = v12 + int32(48)
	return
L187:
	;
	F_errcode(m, int32(261))
	mBase = m.M
	v660 = m.ExcPending
	if v660 != 0 {
		goto L17
	} else {
		goto L188
	}
L188:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v12))) = int32(268435455)
	F_errmsg(m, int32(_a_F_CommitSubTransaction_7), v12)
	mBase = m.M
	v665 = m.ExcPending
	if v665 != 0 {
		goto L17
	} else {
		goto L189
	}
L189:
	;
	F_errfinish(m, int32(_a_F_CommitSubTransaction_3), int32(1738), int32(_a_F_CommitSubTransaction_8))
	mBase = m.M
	v670 = m.ExcPending
	if v670 != 0 {
		goto L17
	} else {
		goto L190
	}
L190:
	;
	base.Wasm_trap_unreachable()
	for {
	}
}
func F_parse_sub_analyze(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) int32 {
	mBase := m.M
	_ = mBase
	var v4 int32
	_ = v4
	var v5 int32
	_ = v5
	var v8 int32
	_ = v8
	var v9 int32
	_ = v9
	var v13 int32
	_ = v13
	var v14 int32
	_ = v14
	var v16 int32
	_ = v16
	v4 = l3
	v5 = F_make_parsestate(m, l1)
	mBase = m.M
	v8 = m.ExcPending
	if v8 != 0 {
		return int32(0)
	} else {
		v9 = int32(1)
		*(*uint8)(unsafe.Add(mBase, uint32(v5)+81)) = uint8(v9)
		*(*uint8)(unsafe.Add(mBase, uint32(v5)+80)) = uint8(v4)
		*(*int32)(unsafe.Add(mBase, uint32(v5)+44)) = l2
		v13 = F_transformStmt(m, v5, l0)
		mBase = m.M
		v14 = m.ExcPending
		if v14 != 0 {
			return int32(0)
		} else {
			F_free_parsestate(m, v5)
			mBase = m.M
			v16 = m.ExcPending
			if v16 != 0 {
				return int32(0)
			} else {
				return v13
			}
		}
	}
}
func F_sub_abs(m *base.Module, l0 int32, l1 int32, l2 int32) {
	mBase := m.M
	_ = mBase
	var v17 int32
	_ = v17
	var v18 int32
	_ = v18
	var v19 int32
	_ = v19
	var v20 int32
	_ = v20
	var v22 int32
	_ = v22
	var v23 int32
	_ = v23
	var v24 int32
	_ = v24
	var v26 int32
	_ = v26
	var v27 int32
	_ = v27
	var v28 int32
	_ = v28
	var v31 int32
	_ = v31
	var v33 int32
	_ = v33
	var v34 int32
	_ = v34
	var v35 int32
	_ = v35
	var v36 int32
	_ = v36
	var v39 int32
	_ = v39
	var v41 int32
	_ = v41
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
	var v53 int32
	_ = v53
	var v54 int32
	_ = v54
	var v59 int32
	_ = v59
	var v60 int32
	_ = v60
	var v70 int32
	_ = v70
	var v71 int32
	_ = v71
	var v80 int32
	_ = v80
	var v82 int32
	_ = v82
	var v83 int32
	_ = v83
	var v86 int32
	_ = v86
	var v87 int32
	_ = v87
	var v96 int32
	_ = v96
	var v98 int32
	_ = v98
	var v106 int32
	_ = v106
	var v112 int32
	_ = v112
	var v114 int32
	_ = v114
	var v118 int32
	_ = v118
	var v122 int32
	_ = v122
	var v126 int32
	_ = v126
	var v128 int32
	_ = v128
	var v130 int32
	_ = v130
	var v141 int32
	_ = v141
	var v145 int32
	_ = v145
	var v163 int32
	_ = v163
	var v164 int32
	_ = v164
	var v168 int32
	_ = v168
	var v169 int32
	_ = v169
	var v178 int32
	_ = v178
	var v197 int32
	_ = v197
	var v199 int32
	_ = v199
	v17 = *(*int32)(unsafe.Add(mBase, uint32(l1)+12))
	v18 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	v19 = *(*int32)(unsafe.Add(mBase, uint32(l1)+20))
	v20 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v22 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v23 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v24 = int32(-1)
	v26 = v22 + (v23 ^ v24)
	v27 = *(*int32)(unsafe.Add(mBase, uint32(l1)))
	v28 = *(*int32)(unsafe.Add(mBase, uint32(l1)+4))
	v31 = v27 + (v28 ^ v24)
	if v31 < v26 {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	v33 = v26
	goto L3
L2:
	;
	v33 = v31
	goto L3
L3:
	;
	v34 = int32(1)
	v35 = v33 + v34
	v36 = v35 + v23
	if v36 <= v34 {
		goto L4
	} else {
		goto L5
	}
L4:
	;
	v39 = int32(1)
	goto L6
L5:
	;
	v39 = v36
	goto L6
L6:
	;
	v41 = v39 << (uint(int32(1)) % 32)
	v44 = F_palloc(m, v41+int32(2))
	mBase = m.M
	v45 = m.ExcPending
	if v45 != 0 {
		goto L7
	} else {
		goto L8
	}
L7:
	;
	return
L8:
	;
	v46 = int32(0)
	*(*uint16)(unsafe.Add(mBase, uint32(v44))) = uint16(v46)
	v48 = *(*int32)(unsafe.Add(mBase, uint32(l1)+4))
	v50 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v53 = v39
	v54 = v46
	v59 = v35 + v48
	v60 = v35 + v50
	goto L9
L9:
	;
	v70 = v60 - int32(1)
	v71 = int32(0)
	if base.B2i32(v70 < v71)|base.B2i32(v22 <= v70) == v71 {
		goto L11
	} else {
		goto L12
	}
L10:
	;
	v112 = *(*int32)(unsafe.Add(mBase, uint32(l2)+16))
	if v112 != 0 {
		goto L21
	} else {
		goto L22
	}
L11:
	;
	v80 = int32(*(*int16)(unsafe.Add(mBase, uint32(v20+v70<<(uint(int32(1))%32)))))
	v82 = v54 + v80
	goto L13
L12:
	;
	v82 = v54
	goto L13
L13:
	;
	v83 = int32(1)
	v86 = v59 - v83
	v87 = int32(0)
	if base.B2i32(v86 < v87)|base.B2i32(v27 <= v86) == v87 {
		goto L14
	} else {
		goto L15
	}
L14:
	;
	v96 = int32(*(*int16)(unsafe.Add(mBase, uint32(v19+v86<<(uint(int32(1))%32)))))
	v98 = v82 - v96
	goto L16
L15:
	;
	v98 = v82
	goto L16
L16:
	;
	if v98 < int32(0) {
		goto L17
	} else {
		goto L18
	}
L17:
	;
	v106 = v98 + int32(_a_F_sub_abs_0)
	goto L19
L18:
	;
	v106 = v98
	goto L19
L19:
	;
	*(*uint16)(unsafe.Add(mBase, uint32(v44+v53<<(uint(int32(1))%32)))) = uint16(v106)
	if base.Ui32(int32(1)) < base.Ui32(v53) {
		v53 = v53 - v83
		v54 = v98 >> (uint(int32(31)) % 32)
		v59 = v86
		v60 = v70
		goto L9
	} else {
		goto L20
	}
L20:
	;
	goto L10
L21:
	;
	F_pfree(m, v112)
	mBase = m.M
	v114 = m.ExcPending
	if v114 != 0 {
		goto L7
	} else {
		goto L24
	}
L22:
	;
	goto L23
L23:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l2)+16)) = v44
	*(*int32)(unsafe.Add(mBase, uint32(l2))) = v39
	if v17 < v18 {
		goto L25
	} else {
		goto L26
	}
L24:
	;
	goto L23
L25:
	;
	v118 = v18
	goto L27
L26:
	;
	v118 = v17
	goto L27
L27:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l2)+12)) = v118
	*(*int32)(unsafe.Add(mBase, uint32(l2)+4)) = v23
	v122 = v44 + int32(2)
	*(*int32)(unsafe.Add(mBase, uint32(l2)+20)) = v122
	v126 = v122
	v128 = v39
	v130 = v23
	goto L30
L28:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l2))) = v199
	*(*int32)(unsafe.Add(mBase, uint32(l2)+20)) = v197
	return
L29:
	;
	*(*int64)(unsafe.Add(mBase, uint32(l2)+4)) = int64(0)
	v197 = v178
	v199 = int32(0)
	goto L28
L30:
	;
	v141 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v126))))
	if v141 != 0 {
		goto L32
	} else {
		goto L33
	}
L31:
	;
	v178 = v122 + v41
	goto L29
L32:
	;
	v145 = v128
	goto L35
L33:
	;
	goto L34
L34:
	;
	v168 = int32(1)
	v169 = v130 - v168
	*(*int32)(unsafe.Add(mBase, uint32(l2)+4)) = v169
	if v168 < v128 {
		v126 = v126 + int32(2)
		v128 = v128 - v168
		v130 = v169
		goto L30
	} else {
		goto L39
	}
L35:
	;
	v163 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v126+v145<<(uint(int32(1))%32)-int32(2)))))
	if v163 != 0 {
		v197 = v126
		v199 = v145
		goto L28
	} else {
		goto L37
	}
L37:
	;
	v164 = int32(1)
	if v164 < v145 {
		v145 = v145 - v164
		goto L35
	} else {
		goto L38
	}
L38:
	;
	v178 = v126
	goto L29
L39:
	;
	goto L31
}
