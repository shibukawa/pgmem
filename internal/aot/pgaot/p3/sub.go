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
	var v36 int32
	_ = v36
	var v40 int32
	_ = v40
	var v44 int32
	_ = v44
	var v50 int32
	_ = v50
	var v53 int32
	_ = v53
	var v57 int32
	_ = v57
	var v59 int32
	_ = v59
	var v60 int32
	_ = v60
	var v65 int32
	_ = v65
	var v66 int32
	_ = v66
	var v69 int32
	_ = v69
	var v74 int32
	_ = v74
	var v76 int32
	_ = v76
	var v82 int32
	_ = v82
	var v87 int32
	_ = v87
	var v89 int32
	_ = v89
	var v92 int32
	_ = v92
	var v93 int32
	_ = v93
	var v94 int32
	_ = v94
	var v95 int32
	_ = v95
	var v104 int32
	_ = v104
	var v106 int32
	_ = v106
	var v107 int32
	_ = v107
	var v109 int32
	_ = v109
	var v110 int32
	_ = v110
	var v120 int32
	_ = v120
	var v122 int32
	_ = v122
	var v123 int32
	_ = v123
	var v126 int32
	_ = v126
	var v127 int32
	_ = v127
	var v128 int32
	_ = v128
	var v134 int32
	_ = v134
	var v139 int32
	_ = v139
	var v145 int32
	_ = v145
	var v146 int32
	_ = v146
	var v148 int32
	_ = v148
	var v149 int32
	_ = v149
	var v150 int32
	_ = v150
	var v151 int32
	_ = v151
	var v152 int32
	_ = v152
	var v155 int32
	_ = v155
	var v157 int32
	_ = v157
	var v159 int32
	_ = v159
	var v162 int32
	_ = v162
	var v164 int32
	_ = v164
	var v168 int32
	_ = v168
	var v171 int32
	_ = v171
	var v172 int32
	_ = v172
	var v175 int32
	_ = v175
	var v176 int32
	_ = v176
	var v177 int32
	_ = v177
	var v178 int32
	_ = v178
	var v180 int32
	_ = v180
	var v182 int32
	_ = v182
	var v183 int32
	_ = v183
	var v184 int32
	_ = v184
	var v185 int32
	_ = v185
	var v186 int32
	_ = v186
	var v190 int64
	_ = v190
	var v192 int32
	_ = v192
	var v196 int32
	_ = v196
	var v199 int32
	_ = v199
	var v200 int32
	_ = v200
	var v201 int32
	_ = v201
	var v207 int32
	_ = v207
	var v211 int32
	_ = v211
	var v213 int32
	_ = v213
	var v215 int32
	_ = v215
	var v226 int32
	_ = v226
	var v227 int32
	_ = v227
	var v228 int32
	_ = v228
	var v229 int32
	_ = v229
	var v230 int32
	_ = v230
	var v231 int32
	_ = v231
	var v232 int32
	_ = v232
	var v234 int32
	_ = v234
	var v237 int32
	_ = v237
	var v239 int32
	_ = v239
	var v241 int32
	_ = v241
	var v242 int32
	_ = v242
	var v243 int32
	_ = v243
	var v244 int32
	_ = v244
	var v253 int32
	_ = v253
	var v254 int32
	_ = v254
	var v258 int32
	_ = v258
	var v263 int32
	_ = v263
	var v266 int32
	_ = v266
	var v268 int32
	_ = v268
	var v273 int32
	_ = v273
	var v276 int32
	_ = v276
	var v278 int32
	_ = v278
	var v285 int32
	_ = v285
	var v288 int32
	_ = v288
	var v293 int32
	_ = v293
	var v298 int32
	_ = v298
	var v299 int32
	_ = v299
	var v313 int32
	_ = v313
	var v314 int32
	_ = v314
	var v315 int32
	_ = v315
	var v317 int32
	_ = v317
	var v319 int32
	_ = v319
	var v320 int32
	_ = v320
	var v322 int32
	_ = v322
	var v325 int32
	_ = v325
	var v327 int32
	_ = v327
	var v328 int32
	_ = v328
	var v337 int32
	_ = v337
	var v338 int32
	_ = v338
	var v339 int32
	_ = v339
	var v340 int32
	_ = v340
	var v342 int32
	_ = v342
	var v345 int32
	_ = v345
	var v349 int32
	_ = v349
	var v352 int32
	_ = v352
	var v354 int32
	_ = v354
	var v355 int32
	_ = v355
	var v362 int32
	_ = v362
	var v365 int32
	_ = v365
	var v368 int32
	_ = v368
	var v372 int32
	_ = v372
	var v380 int32
	_ = v380
	var v384 int32
	_ = v384
	var v385 int32
	_ = v385
	var v386 int32
	_ = v386
	var v390 int32
	_ = v390
	var v392 int32
	_ = v392
	var v393 int32
	_ = v393
	var v405 int32
	_ = v405
	var v416 int32
	_ = v416
	var v417 int32
	_ = v417
	var v418 int32
	_ = v418
	var v419 int32
	_ = v419
	var v420 int32
	_ = v420
	var v429 int32
	_ = v429
	var v431 int32
	_ = v431
	var v432 int32
	_ = v432
	var v434 int32
	_ = v434
	var v444 int32
	_ = v444
	var v445 int32
	_ = v445
	var v449 int32
	_ = v449
	var v451 int32
	_ = v451
	var v452 int32
	_ = v452
	var v453 int32
	_ = v453
	var v455 int32
	_ = v455
	var v457 int32
	_ = v457
	var v460 int32
	_ = v460
	var v462 int32
	_ = v462
	var v463 int32
	_ = v463
	var v465 int32
	_ = v465
	var v468 int32
	_ = v468
	var v477 int32
	_ = v477
	var v480 int32
	_ = v480
	var v491 int32
	_ = v491
	var v493 int32
	_ = v493
	var v494 int32
	_ = v494
	var v496 int32
	_ = v496
	var v505 int32
	_ = v505
	var v506 int32
	_ = v506
	var v510 int32
	_ = v510
	var v512 int32
	_ = v512
	var v517 int32
	_ = v517
	var v518 int32
	_ = v518
	var v523 int32
	_ = v523
	var v525 int32
	_ = v525
	var v527 int32
	_ = v527
	var v529 int32
	_ = v529
	var v531 int32
	_ = v531
	var v533 int32
	_ = v533
	var v534 int32
	_ = v534
	var v535 int32
	_ = v535
	var v537 int32
	_ = v537
	var v539 int32
	_ = v539
	var v540 int32
	_ = v540
	var v541 int32
	_ = v541
	var v543 int32
	_ = v543
	var v567 int32
	_ = v567
	var v568 int32
	_ = v568
	var v569 int32
	_ = v569
	var v571 int32
	_ = v571
	var v573 int32
	_ = v573
	var v575 int32
	_ = v575
	var v577 int32
	_ = v577
	var v579 int32
	_ = v579
	var v580 int32
	_ = v580
	var v582 int32
	_ = v582
	var v587 int32
	_ = v587
	var v596 int32
	_ = v596
	var v599 int32
	_ = v599
	var v610 int32
	_ = v610
	var v613 int32
	_ = v613
	var v614 int32
	_ = v614
	var v618 int32
	_ = v618
	var v620 int32
	_ = v620
	var v625 int32
	_ = v625
	var v626 int32
	_ = v626
	var v627 int32
	_ = v627
	var v631 int32
	_ = v631
	var v632 int32
	_ = v632
	var v634 int32
	_ = v634
	var v635 int32
	_ = v635
	var v636 int32
	_ = v636
	var v637 int32
	_ = v637
	var v638 int32
	_ = v638
	var v639 int32
	_ = v639
	var v641 int32
	_ = v641
	var v647 int32
	_ = v647
	var v654 int32
	_ = v654
	var v657 int32
	_ = v657
	var v662 int32
	_ = v662
	var v667 int32
	_ = v667
	v10 = m.G0
	v12 = v10 - int32(48)
	m.G0 = v12
	v15 = *(*int32)(unsafe.Add(mBase, _c_F_CommitSubTransaction[0]))
	v16 = int32(10)
	goto L3
L1:
	;
	if v53 != 0 {
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
	goto L6
L4:
	;
	v36 = int32(0)
	goto L11
L6:
	;
	goto L7
L7:
	;
	if int32(0)|base.B2i32(v23 == int32(15)) != 0 {
		goto L4
	} else {
		goto L9
	}
L9:
	;
	if v23 <= v16 {
		v53 = int32(1)
		goto L2
	} else {
		goto L10
	}
L10:
	;
	goto L4
L11:
	;
	v40 = *(*int32)(unsafe.Add(mBase, _c_F_CommitSubTransaction[2]))
	if v40 != int32(2) {
		v53 = v36
		goto L2
	} else {
		goto L12
	}
L12:
	;
	v44 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_CommitSubTransaction[3])))
	if v44&int32(1) != 0 {
		v53 = v36
		goto L2
	} else {
		goto L13
	}
L13:
	;
	v50 = *(*int32)(unsafe.Add(mBase, _c_F_CommitSubTransaction[4]))
	v53 = int32(0) | base.B2i32(v50 <= v16)
	goto L2
L14:
	;
	v57 = *(*int32)(unsafe.Add(mBase, _c_F_CommitSubTransaction[0]))
	F_ShowTransactionStateRec(m, int32(_a_F_CommitSubTransaction_0), v57)
	mBase = m.M
	v59 = m.ExcPending
	if v59 != 0 {
		goto L17
	} else {
		goto L18
	}
L15:
	;
	goto L16
L16:
	;
	v60 = *(*int32)(unsafe.Add(mBase, uint32(v15)+20))
	if v60 == int32(2) {
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
	v89 = *(*int32)(unsafe.Add(mBase, uint32(v15)+8))
	v92 = *(*int32)(unsafe.Add(mBase, _c_F_CommitSubTransaction[5]))
	if v92 != 0 {
		goto L28
	} else {
		goto L29
	}
L20:
	;
	v65 = F_errstart(m, int32(19), int32(0))
	mBase = m.M
	v66 = m.ExcPending
	if v66 != 0 {
		goto L17
	} else {
		goto L21
	}
L21:
	;
	if v65 == int32(0) {
		goto L19
	} else {
		goto L22
	}
L22:
	;
	v69 = *(*int32)(unsafe.Add(mBase, uint32(v15)+20))
	if base.Ui32(v69) <= base.Ui32(int32(5)) {
		goto L23
	} else {
		goto L24
	}
L23:
	;
	v74 = *(*int32)(unsafe.Add(mBase, uint32(v69<<(uint(int32(2))%32))+uint32(_c_F_CommitSubTransaction[6])))
	v76 = v74
	goto L25
L24:
	;
	v76 = int32(_a_F_CommitSubTransaction_1)
	goto L25
L25:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v12)+32)) = v76
	F_errmsg_internal(m, int32(_a_F_CommitSubTransaction_2), v12+int32(32))
	mBase = m.M
	v82 = m.ExcPending
	if v82 != 0 {
		goto L17
	} else {
		goto L26
	}
L26:
	;
	F_errfinish(m, int32(_a_F_CommitSubTransaction_3), int32(_a_F_CommitSubTransaction_4), int32(_a_F_CommitSubTransaction_0))
	mBase = m.M
	v87 = m.ExcPending
	if v87 != 0 {
		goto L17
	} else {
		goto L27
	}
L27:
	;
	goto L19
L28:
	;
	v93 = *(*int32)(unsafe.Add(mBase, uint32(v15)+80))
	v94 = *(*int32)(unsafe.Add(mBase, uint32(v93)+8))
	v95 = v92
	goto L31
L29:
	;
	v120 = v89
	goto L30
L30:
	;
	F_AtEOSubXact_Parallel(m, int32(1), v120)
	mBase = m.M
	v122 = m.ExcPending
	if v122 != 0 {
		goto L17
	} else {
		goto L35
	}
L31:
	;
	v104 = *(*int32)(unsafe.Add(mBase, uint32(v95)))
	v106 = *(*int32)(unsafe.Add(mBase, uint32(v95)+8))
	v107 = *(*int32)(unsafe.Add(mBase, uint32(v95)+4))
	m.T0[v107].(func(*base.Module, int32, int32, int32, int32))(m, int32(3), v89, v94, v106)
	mBase = m.M
	v109 = m.ExcPending
	if v109 != 0 {
		goto L17
	} else {
		goto L33
	}
L32:
	;
	v110 = *(*int32)(unsafe.Add(mBase, uint32(v15)+8))
	v120 = v110
	goto L30
L33:
	;
	if v104 != 0 {
		v95 = v104
		goto L31
	} else {
		goto L34
	}
L34:
	;
	goto L32
L35:
	;
	v123 = *(*int32)(unsafe.Add(mBase, uint32(v15)+72))
	if v123 != 0 {
		goto L36
	} else {
		goto L37
	}
L36:
	;
	v126 = F_errstart(m, int32(19), int32(0))
	mBase = m.M
	v127 = m.ExcPending
	if v127 != 0 {
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
	v145 = m.ExcPending
	if v145 != 0 {
		goto L17
	} else {
		goto L45
	}
L39:
	;
	if v126 != 0 {
		goto L40
	} else {
		goto L41
	}
L40:
	;
	v128 = *(*int32)(unsafe.Add(mBase, uint32(v15)+72))
	*(*int32)(unsafe.Add(mBase, uint32(v12)+16)) = v128
	F_errmsg_internal(m, int32(_a_F_CommitSubTransaction_5), v12+int32(16))
	mBase = m.M
	v134 = m.ExcPending
	if v134 != 0 {
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
	v139 = m.ExcPending
	if v139 != 0 {
		goto L17
	} else {
		goto L44
	}
L44:
	;
	goto L42
L45:
	;
	v146 = *(*int32)(unsafe.Add(mBase, uint32(v15)))
	if v146 != 0 {
		goto L47
	} else {
		goto L48
	}
L46:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v654 = m.ExcPending
	if v654 != 0 {
		goto L17
	} else {
		goto L187
	}
L47:
	;
	v148 = *(*int32)(unsafe.Add(mBase, _c_F_CommitSubTransaction[0]))
	v149 = *(*int32)(unsafe.Add(mBase, uint32(v148)+80))
	v150 = *(*int32)(unsafe.Add(mBase, uint32(v149)+56))
	v151 = *(*int32)(unsafe.Add(mBase, uint32(v149)+52))
	v152 = *(*int32)(unsafe.Add(mBase, uint32(v148)+52))
	v155 = v151 + v152 + int32(1)
	if v150 < v155 {
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
	v226 = m.ExcPending
	if v226 != 0 {
		goto L17
	} else {
		goto L70
	}
L50:
	;
	v157 = int32(268435455)
	v159 = v155 << (uint(int32(1)) % 32)
	if v157 <= v159 {
		goto L53
	} else {
		goto L54
	}
L51:
	;
	v184 = v151
	v185 = v149
	goto L52
L52:
	;
	v186 = *(*int32)(unsafe.Add(mBase, uint32(v185)+48))
	v190 = *(*int64)(unsafe.Add(mBase, uint32(v148)))
	*(*uint32)(unsafe.Add(mBase, uint32(v186+v184<<(uint(int32(2))%32)))) = uint32(v190)
	v192 = *(*int32)(unsafe.Add(mBase, uint32(v148)+52))
	if v192 <= int32(0) {
		goto L63
	} else {
		goto L64
	}
L53:
	;
	v162 = v157
	goto L55
L54:
	;
	v162 = v159
	goto L55
L55:
	;
	if v162 < v155 {
		goto L46
	} else {
		goto L56
	}
L56:
	;
	v164 = *(*int32)(unsafe.Add(mBase, uint32(v149)+48))
	if v164 == int32(0) {
		goto L58
	} else {
		goto L59
	}
L57:
	;
	v178 = *(*int32)(unsafe.Add(mBase, uint32(v148)+80))
	*(*int32)(unsafe.Add(mBase, uint32(v178)+48)) = v177
	v180 = *(*int32)(unsafe.Add(mBase, uint32(v148)+80))
	*(*int32)(unsafe.Add(mBase, uint32(v180)+56)) = v162
	v182 = *(*int32)(unsafe.Add(mBase, uint32(v148)+80))
	v183 = *(*int32)(unsafe.Add(mBase, uint32(v182)+52))
	v184 = v183
	v185 = v182
	goto L52
L58:
	;
	v168 = *(*int32)(unsafe.Add(mBase, _c_F_CommitSubTransaction[7]))
	v171 = F_MemoryContextAlloc(m, v168, v162<<(uint(int32(2))%32))
	mBase = m.M
	v172 = m.ExcPending
	if v172 != 0 {
		goto L17
	} else {
		goto L61
	}
L59:
	;
	goto L60
L60:
	;
	v175 = F_repalloc(m, v164, v162<<(uint(int32(2))%32))
	mBase = m.M
	v176 = m.ExcPending
	if v176 != 0 {
		goto L17
	} else {
		goto L62
	}
L61:
	;
	v177 = v171
	goto L57
L62:
	;
	v177 = v175
	goto L57
L63:
	;
	v211 = *(*int32)(unsafe.Add(mBase, uint32(v148)+80))
	*(*int32)(unsafe.Add(mBase, uint32(v211)+52)) = v155
	v213 = *(*int32)(unsafe.Add(mBase, uint32(v148)+48))
	if v213 != 0 {
		goto L66
	} else {
		goto L67
	}
L64:
	;
	v196 = v192 << (uint(int32(2)) % 32)
	if v196 == int32(0) {
		goto L63
	} else {
		goto L65
	}
L65:
	;
	v199 = *(*int32)(unsafe.Add(mBase, uint32(v148)+80))
	v200 = *(*int32)(unsafe.Add(mBase, uint32(v199)+48))
	v201 = *(*int32)(unsafe.Add(mBase, uint32(v199)+52))
	v207 = *(*int32)(unsafe.Add(mBase, uint32(v148)+48))
	base.MemoryCopy(m, v200+v201<<(uint(int32(2))%32)+int32(4), v207, v196)
	goto L63
L66:
	;
	F_pfree(m, v213)
	mBase = m.M
	v215 = m.ExcPending
	if v215 != 0 {
		goto L17
	} else {
		goto L69
	}
L67:
	;
	goto L68
L68:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v148)+56)) = int32(0)
	*(*int64)(unsafe.Add(mBase, uint32(v148)+48)) = int64(0)
	goto L49
L69:
	;
	goto L68
L70:
	;
	v227 = *(*int32)(unsafe.Add(mBase, uint32(v15)+8))
	v228 = *(*int32)(unsafe.Add(mBase, uint32(v15)+80))
	v229 = *(*int32)(unsafe.Add(mBase, uint32(v228)+8))
	v230 = *(*int32)(unsafe.Add(mBase, uint32(v228)+28))
	v231 = *(*int32)(unsafe.Add(mBase, uint32(v228)+40))
	v232 = m.G0
	v234 = v232 - int32(32)
	m.G0 = v234
	v237 = v234 + int32(12)
	v239 = *(*int32)(unsafe.Add(mBase, _c_F_CommitSubTransaction[8]))
	F_hash_seq_init(m, v237, v239)
	mBase = m.M
	v241 = m.ExcPending
	if v241 != 0 {
		goto L17
	} else {
		goto L71
	}
L71:
	;
	v242 = F_hash_seq_search(m, v237)
	mBase = m.M
	v243 = m.ExcPending
	if v243 != 0 {
		goto L17
	} else {
		goto L72
	}
L72:
	;
	if v242 != 0 {
		goto L73
	} else {
		goto L74
	}
L73:
	;
	v244 = v242
	goto L76
L74:
	;
	goto L75
L75:
	;
	m.G0 = v234 + int32(32)
	v313 = *(*int32)(unsafe.Add(mBase, uint32(v15)+8))
	v314 = *(*int32)(unsafe.Add(mBase, uint32(v15)+80))
	v315 = *(*int32)(unsafe.Add(mBase, uint32(v314)+8))
	F_AtEOSubXact_LargeObject(m, int32(1), v313, v315)
	mBase = m.M
	v317 = m.ExcPending
	if v317 != 0 {
		goto L17
	} else {
		goto L99
	}
L76:
	;
	v253 = *(*int32)(unsafe.Add(mBase, uint32(v244)+64))
	v254 = *(*int32)(unsafe.Add(mBase, uint32(v253)+20))
	if v254 != v227 {
		goto L78
	} else {
		goto L79
	}
L77:
	;
	goto L75
L78:
	;
	v293 = *(*int32)(unsafe.Add(mBase, uint32(v253)+24))
	if v227 == v293 {
		goto L94
	} else {
		goto L95
	}
L79:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v253)+28)) = v230
	*(*int32)(unsafe.Add(mBase, uint32(v253)+20)) = v229
	v258 = *(*int32)(unsafe.Add(mBase, uint32(v253)+12))
	if v258 == int32(0) {
		goto L78
	} else {
		goto L80
	}
L80:
	;
	v263 = *(*int32)(unsafe.Add(mBase, uint32(v258)))
	if v263 == int32(0) {
		goto L82
	} else {
		goto L83
	}
L81:
	;
	goto L78
L82:
	;
	if v231 != 0 {
		goto L91
	} else {
		goto L92
	}
L83:
	;
	v266 = *(*int32)(unsafe.Add(mBase, uint32(v263)+4))
	if v266 == v258 {
		goto L84
	} else {
		goto L85
	}
L84:
	;
	v268 = *(*int32)(unsafe.Add(mBase, uint32(v258)+8))
	*(*int32)(unsafe.Add(mBase, uint32(v263)+4)) = v268
	goto L82
L85:
	;
	goto L86
L86:
	;
	v273 = v266
	goto L87
L87:
	;
	if v273 == int32(0) {
		goto L82
	} else {
		goto L89
	}
L88:
	;
	v278 = *(*int32)(unsafe.Add(mBase, uint32(v258)+8))
	*(*int32)(unsafe.Add(mBase, uint32(v273)+8)) = v278
	goto L82
L89:
	;
	v276 = *(*int32)(unsafe.Add(mBase, uint32(v273)+8))
	if v258 != v276 {
		v273 = v276
		goto L87
	} else {
		goto L90
	}
L90:
	;
	goto L88
L91:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v258))) = v231
	v285 = *(*int32)(unsafe.Add(mBase, uint32(v231)+4))
	*(*int32)(unsafe.Add(mBase, uint32(v258)+8)) = v285
	*(*int32)(unsafe.Add(mBase, uint32(v231)+4)) = v258
	goto L81
L92:
	;
	goto L93
L93:
	;
	v288 = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v258)+8)) = v288
	*(*int32)(unsafe.Add(mBase, uint32(v258))) = v288
	goto L81
L94:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v253)+24)) = v229
	goto L96
L95:
	;
	goto L96
L96:
	;
	v298 = F_hash_seq_search(m, v234+int32(12))
	mBase = m.M
	v299 = m.ExcPending
	if v299 != 0 {
		goto L17
	} else {
		goto L97
	}
L97:
	;
	if v298 != 0 {
		v244 = v298
		goto L76
	} else {
		goto L98
	}
L98:
	;
	goto L77
L99:
	;
	v319 = *(*int32)(unsafe.Add(mBase, _c_F_CommitSubTransaction[0]))
	v320 = *(*int32)(unsafe.Add(mBase, uint32(v319)+28))
	goto L100
L100:
	;
	v322 = *(*int32)(unsafe.Add(mBase, _c_F_CommitSubTransaction[9]))
	if v322 == int32(0) {
		goto L101
	} else {
		goto L102
	}
L101:
	;
	v349 = *(*int32)(unsafe.Add(mBase, _c_F_CommitSubTransaction[10]))
	if v349 == int32(0) {
		goto L111
	} else {
		goto L112
	}
L102:
	;
	v325 = *(*int32)(unsafe.Add(mBase, uint32(v322)))
	if v325 < v320 {
		goto L101
	} else {
		goto L103
	}
L103:
	;
	v327 = *(*int32)(unsafe.Add(mBase, uint32(v322)+8))
	if v327 != 0 {
		goto L105
	} else {
		goto L106
	}
L104:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_CommitSubTransaction[9])) = v327
	v337 = *(*int32)(unsafe.Add(mBase, uint32(v327)+4))
	v338 = *(*int32)(unsafe.Add(mBase, uint32(v322)+4))
	v339 = F_list_concat(m, v337, v338)
	mBase = m.M
	v340 = m.ExcPending
	if v340 != 0 {
		goto L17
	} else {
		goto L109
	}
L105:
	;
	v328 = *(*int32)(unsafe.Add(mBase, uint32(v327)))
	if v320-int32(1) <= v328 {
		goto L104
	} else {
		goto L108
	}
L106:
	;
	goto L107
L107:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v322))) = v325 - int32(1)
	goto L101
L108:
	;
	goto L107
L109:
	;
	v342 = *(*int32)(unsafe.Add(mBase, _c_F_CommitSubTransaction[9]))
	*(*int32)(unsafe.Add(mBase, uint32(v342)+4)) = v339
	F_pfree(m, v322)
	mBase = m.M
	v345 = m.ExcPending
	if v345 != 0 {
		goto L17
	} else {
		goto L110
	}
L110:
	;
	goto L101
L111:
	;
	v416 = *(*int32)(unsafe.Add(mBase, _c_F_CommitSubTransaction[5]))
	if v416 != 0 {
		goto L131
	} else {
		goto L132
	}
L112:
	;
	v352 = *(*int32)(unsafe.Add(mBase, uint32(v349)))
	if v352 < v320 {
		goto L111
	} else {
		goto L113
	}
L113:
	;
	v354 = *(*int32)(unsafe.Add(mBase, uint32(v349)+12))
	if v354 != 0 {
		goto L115
	} else {
		goto L116
	}
L114:
	;
	v362 = int32(0)
	*(*int32)(unsafe.Add(mBase, _c_F_CommitSubTransaction[10])) = v354
	v365 = *(*int32)(unsafe.Add(mBase, uint32(v349)+4))
	if v365 == v362 {
		goto L119
	} else {
		goto L120
	}
L115:
	;
	v355 = *(*int32)(unsafe.Add(mBase, uint32(v354)))
	if v320-int32(1) <= v355 {
		goto L114
	} else {
		goto L118
	}
L116:
	;
	goto L117
L117:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v349))) = v352 - int32(1)
	goto L111
L118:
	;
	goto L117
L119:
	;
	F_pfree(m, v349)
	mBase = m.M
	v405 = m.ExcPending
	if v405 != 0 {
		goto L17
	} else {
		goto L130
	}
L120:
	;
	v368 = *(*int32)(unsafe.Add(mBase, uint32(v365)+4))
	if v368 <= int32(0) {
		goto L119
	} else {
		goto L121
	}
L121:
	;
	v372 = v362
	goto L122
L122:
	;
	v380 = *(*int32)(unsafe.Add(mBase, uint32(v365)+12))
	v384 = *(*int32)(unsafe.Add(mBase, uint32(v380+v372<<(uint(int32(2))%32))))
	v385 = F_AsyncExistsPendingNotify(m, v384)
	mBase = m.M
	v386 = m.ExcPending
	if v386 != 0 {
		goto L17
	} else {
		goto L124
	}
L123:
	;
	goto L119
L124:
	;
	if v385 == int32(0) {
		goto L125
	} else {
		goto L126
	}
L125:
	;
	F_AddEventToPendingNotifies(m, v384)
	mBase = m.M
	v390 = m.ExcPending
	if v390 != 0 {
		goto L17
	} else {
		goto L128
	}
L126:
	;
	goto L127
L127:
	;
	v392 = v372 + int32(1)
	v393 = *(*int32)(unsafe.Add(mBase, uint32(v365)+4))
	if v392 < v393 {
		v372 = v392
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
	v417 = *(*int32)(unsafe.Add(mBase, uint32(v15)+8))
	v418 = *(*int32)(unsafe.Add(mBase, uint32(v15)+80))
	v419 = *(*int32)(unsafe.Add(mBase, uint32(v418)+8))
	v420 = v416
	goto L134
L132:
	;
	goto L133
L133:
	;
	v444 = *(*int32)(unsafe.Add(mBase, uint32(v15)+40))
	v445 = int32(1)
	F_ResourceOwnerReleaseInternal(m, v444, v445, v445, int32(0))
	mBase = m.M
	v449 = m.ExcPending
	if v449 != 0 {
		goto L17
	} else {
		goto L138
	}
L134:
	;
	v429 = *(*int32)(unsafe.Add(mBase, uint32(v420)))
	v431 = *(*int32)(unsafe.Add(mBase, uint32(v420)+8))
	v432 = *(*int32)(unsafe.Add(mBase, uint32(v420)+4))
	m.T0[v432].(func(*base.Module, int32, int32, int32, int32))(m, int32(1), v417, v419, v431)
	mBase = m.M
	v434 = m.ExcPending
	if v434 != 0 {
		goto L17
	} else {
		goto L136
	}
L135:
	;
	goto L133
L136:
	;
	if v429 != 0 {
		v420 = v429
		goto L134
	} else {
		goto L137
	}
L137:
	;
	goto L135
L138:
	;
	v451 = *(*int32)(unsafe.Add(mBase, uint32(v15)+8))
	v452 = *(*int32)(unsafe.Add(mBase, uint32(v15)+80))
	v453 = *(*int32)(unsafe.Add(mBase, uint32(v452)+8))
	F_AtEOSubXact_RelationCache(m, int32(1), v451, v453)
	mBase = m.M
	v455 = m.ExcPending
	if v455 != 0 {
		goto L17
	} else {
		goto L139
	}
L139:
	;
	F_AtEOXact_TypeCache(m)
	mBase = m.M
	v457 = m.ExcPending
	if v457 != 0 {
		goto L17
	} else {
		goto L140
	}
L140:
	;
	F_AtEOSubXact_Inval(m, int32(1))
	mBase = m.M
	v460 = m.ExcPending
	if v460 != 0 {
		goto L17
	} else {
		goto L141
	}
L141:
	;
	v462 = *(*int32)(unsafe.Add(mBase, _c_F_CommitSubTransaction[0]))
	v463 = *(*int32)(unsafe.Add(mBase, uint32(v462)+28))
	goto L142
L142:
	;
	v465 = *(*int32)(unsafe.Add(mBase, _c_F_CommitSubTransaction[11]))
	if v465 != 0 {
		goto L143
	} else {
		goto L144
	}
L143:
	;
	v468 = v465
	goto L146
L144:
	;
	goto L145
L145:
	;
	v491 = *(*int32)(unsafe.Add(mBase, uint32(v15)+40))
	*(*int32)(unsafe.Add(mBase, _c_F_CommitSubTransaction[12])) = v491
	v493 = *(*int32)(unsafe.Add(mBase, uint32(v15)))
	if v493 != 0 {
		goto L152
	} else {
		goto L153
	}
L146:
	;
	v477 = *(*int32)(unsafe.Add(mBase, uint32(v468)+20))
	if v463 <= v477 {
		goto L148
	} else {
		goto L149
	}
L147:
	;
	goto L145
L148:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v468)+20)) = v463 - int32(1)
	goto L150
L149:
	;
	goto L150
L150:
	;
	v480 = *(*int32)(unsafe.Add(mBase, uint32(v468)+24))
	if v480 != 0 {
		v468 = v480
		goto L146
	} else {
		goto L151
	}
L151:
	;
	goto L147
L152:
	;
	v494 = m.G0
	v496 = v494 - int32(16)
	m.G0 = v496
	*(*int32)(unsafe.Add(mBase, uint32(v496)+12)) = int32(17104896)
	*(*int64)(unsafe.Add(mBase, uint32(v496)+4)) = int64(0)
	*(*int32)(unsafe.Add(mBase, uint32(v496))) = v493
	v505 = F_LockRelease(m, v496, int32(7), int32(0))
	mBase = m.M
	v506 = m.ExcPending
	if v506 != 0 {
		goto L17
	} else {
		goto L155
	}
L153:
	;
	v512 = v491
	goto L154
L154:
	;
	F_ResourceOwnerReleaseInternal(m, v512, int32(2), int32(1), int32(0))
	mBase = m.M
	v517 = m.ExcPending
	if v517 != 0 {
		goto L17
	} else {
		goto L156
	}
L155:
	;
	m.G0 = v496 + int32(16)
	v510 = *(*int32)(unsafe.Add(mBase, uint32(v15)+40))
	v512 = v510
	goto L154
L156:
	;
	v518 = *(*int32)(unsafe.Add(mBase, uint32(v15)+40))
	F_ResourceOwnerReleaseInternal(m, v518, int32(3), int32(1), int32(0))
	mBase = m.M
	v523 = m.ExcPending
	if v523 != 0 {
		goto L17
	} else {
		goto L157
	}
L157:
	;
	v525 = *(*int32)(unsafe.Add(mBase, uint32(v15)+32))
	F_AtEOXact_GUC(m, int32(1), v525)
	mBase = m.M
	v527 = m.ExcPending
	if v527 != 0 {
		goto L17
	} else {
		goto L158
	}
L158:
	;
	v529 = *(*int32)(unsafe.Add(mBase, uint32(v15)+8))
	F_AtEOSubXact_SPI(m, int32(1), v529)
	mBase = m.M
	v531 = m.ExcPending
	if v531 != 0 {
		goto L17
	} else {
		goto L159
	}
L159:
	;
	v533 = *(*int32)(unsafe.Add(mBase, uint32(v15)+8))
	v534 = *(*int32)(unsafe.Add(mBase, uint32(v15)+80))
	v535 = *(*int32)(unsafe.Add(mBase, uint32(v534)+8))
	F_AtEOSubXact_on_commit_actions(m, int32(1), v533, v535)
	mBase = m.M
	v537 = m.ExcPending
	if v537 != 0 {
		goto L17
	} else {
		goto L160
	}
L160:
	;
	v539 = *(*int32)(unsafe.Add(mBase, uint32(v15)+8))
	v540 = *(*int32)(unsafe.Add(mBase, uint32(v15)+80))
	v541 = *(*int32)(unsafe.Add(mBase, uint32(v540)+8))
	v543 = *(*int32)(unsafe.Add(mBase, _c_F_CommitSubTransaction[13]))
	if v539 == v543 {
		goto L162
	} else {
		goto L163
	}
L161:
	;
	v567 = *(*int32)(unsafe.Add(mBase, uint32(v15)+8))
	v568 = *(*int32)(unsafe.Add(mBase, uint32(v15)+80))
	v569 = *(*int32)(unsafe.Add(mBase, uint32(v568)+8))
	F_AtEOSubXact_Files(m, int32(1), v567, v569)
	mBase = m.M
	v571 = m.ExcPending
	if v571 != 0 {
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
	*(*int32)(unsafe.Add(mBase, _c_F_CommitSubTransaction[13])) = v541
	goto L161
L168:
	;
	v573 = *(*int32)(unsafe.Add(mBase, uint32(v15)+28))
	F_AtEOSubXact_HashTables(m, int32(1), v573)
	mBase = m.M
	v575 = m.ExcPending
	if v575 != 0 {
		goto L17
	} else {
		goto L169
	}
L169:
	;
	v577 = *(*int32)(unsafe.Add(mBase, uint32(v15)+28))
	F_AtEOSubXact_PgStat(m, int32(1), v577)
	mBase = m.M
	v579 = m.ExcPending
	if v579 != 0 {
		goto L17
	} else {
		goto L170
	}
L170:
	;
	v580 = *(*int32)(unsafe.Add(mBase, uint32(v15)+28))
	v582 = *(*int32)(unsafe.Add(mBase, _c_F_CommitSubTransaction[14]))
	if v582 == int32(0) {
		goto L171
	} else {
		goto L172
	}
L171:
	;
	v610 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v15)+68)))
	*(*uint8)(unsafe.Add(mBase, _c_F_CommitSubTransaction[15])) = uint8(v610)
	v613 = *(*int32)(unsafe.Add(mBase, uint32(v15)+80))
	v614 = *(*int32)(unsafe.Add(mBase, uint32(v613)+40))
	*(*int32)(unsafe.Add(mBase, _c_F_CommitSubTransaction[16])) = v614
	*(*int32)(unsafe.Add(mBase, _c_F_CommitSubTransaction[12])) = v614
	v618 = *(*int32)(unsafe.Add(mBase, uint32(v15)+40))
	F_ResourceOwnerDelete(m, v618)
	mBase = m.M
	v620 = m.ExcPending
	if v620 != 0 {
		goto L17
	} else {
		goto L177
	}
L172:
	;
	v587 = v582
	goto L173
L173:
	;
	v596 = *(*int32)(unsafe.Add(mBase, uint32(v587)+4))
	if v596 < v580 {
		goto L171
	} else {
		goto L175
	}
L174:
	;
	goto L171
L175:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v587)+4)) = v580 - int32(1)
	v599 = *(*int32)(unsafe.Add(mBase, uint32(v587)+8))
	if v599 != 0 {
		v587 = v599
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
	v625 = *(*int32)(unsafe.Add(mBase, _c_F_CommitSubTransaction[0]))
	v626 = *(*int32)(unsafe.Add(mBase, uint32(v625)+80))
	v627 = *(*int32)(unsafe.Add(mBase, uint32(v626)+36))
	*(*int32)(unsafe.Add(mBase, _c_F_CommitSubTransaction[17])) = v627
	*(*int32)(unsafe.Add(mBase, _c_F_CommitSubTransaction[18])) = v627
	v631 = *(*int32)(unsafe.Add(mBase, uint32(v625)+36))
	v632 = *(*int32)(unsafe.Add(mBase, uint32(v631)+20))
	if v632 != 0 {
		goto L178
	} else {
		goto L179
	}
L178:
	;
	v638 = int32(0)
	goto L180
L179:
	;
	v634 = *(*int32)(unsafe.Add(mBase, uint32(v631)+12))
	v635 = *(*int32)(unsafe.Add(mBase, uint32(v634)+28))
	v636 = m.T0[v635].(func(*base.Module, int32) int32)(m, v631)
	mBase = m.M
	v637 = m.ExcPending
	if v637 != 0 {
		goto L17
	} else {
		goto L181
	}
L180:
	;
	if v638 != 0 {
		goto L182
	} else {
		goto L183
	}
L181:
	;
	v638 = v636
	goto L180
L182:
	;
	v639 = *(*int32)(unsafe.Add(mBase, uint32(v625)+36))
	F_MemoryContextDelete(m, v639)
	mBase = m.M
	v641 = m.ExcPending
	if v641 != 0 {
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
	v647 = m.ExcPending
	if v647 != 0 {
		goto L17
	} else {
		goto L186
	}
L185:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v625)+36)) = int32(0)
	goto L184
L186:
	;
	m.G0 = v12 + int32(48)
	return
L187:
	;
	F_errcode(m, int32(261))
	mBase = m.M
	v657 = m.ExcPending
	if v657 != 0 {
		goto L17
	} else {
		goto L188
	}
L188:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v12))) = int32(268435455)
	F_errmsg(m, int32(_a_F_CommitSubTransaction_7), v12)
	mBase = m.M
	v662 = m.ExcPending
	if v662 != 0 {
		goto L17
	} else {
		goto L189
	}
L189:
	;
	F_errfinish(m, int32(_a_F_CommitSubTransaction_3), int32(1696), int32(_a_F_CommitSubTransaction_8))
	mBase = m.M
	v667 = m.ExcPending
	if v667 != 0 {
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
		*(*uint8)(unsafe.Add(mBase, uint32(v5)+85)) = uint8(v9)
		*(*uint8)(unsafe.Add(mBase, uint32(v5)+84)) = uint8(v4)
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
