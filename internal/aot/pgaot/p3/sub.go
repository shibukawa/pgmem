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
	var v35 int32
	_ = v35
	var v39 int32
	_ = v39
	var v43 int32
	_ = v43
	var v47 int32
	_ = v47
	var v50 int32
	_ = v50
	var v54 int32
	_ = v54
	var v56 int32
	_ = v56
	var v57 int32
	_ = v57
	var v62 int32
	_ = v62
	var v63 int32
	_ = v63
	var v66 int32
	_ = v66
	var v73 int32
	_ = v73
	var v75 int32
	_ = v75
	var v81 int32
	_ = v81
	var v86 int32
	_ = v86
	var v88 int32
	_ = v88
	var v91 int32
	_ = v91
	var v92 int32
	_ = v92
	var v93 int32
	_ = v93
	var v94 int32
	_ = v94
	var v103 int32
	_ = v103
	var v105 int32
	_ = v105
	var v106 int32
	_ = v106
	var v108 int32
	_ = v108
	var v109 int32
	_ = v109
	var v119 int32
	_ = v119
	var v121 int32
	_ = v121
	var v122 int32
	_ = v122
	var v125 int32
	_ = v125
	var v126 int32
	_ = v126
	var v127 int32
	_ = v127
	var v133 int32
	_ = v133
	var v138 int32
	_ = v138
	var v144 int32
	_ = v144
	var v145 int32
	_ = v145
	var v147 int32
	_ = v147
	var v148 int32
	_ = v148
	var v149 int32
	_ = v149
	var v150 int32
	_ = v150
	var v151 int32
	_ = v151
	var v154 int32
	_ = v154
	var v156 int32
	_ = v156
	var v158 int32
	_ = v158
	var v161 int32
	_ = v161
	var v163 int32
	_ = v163
	var v167 int32
	_ = v167
	var v170 int32
	_ = v170
	var v171 int32
	_ = v171
	var v174 int32
	_ = v174
	var v175 int32
	_ = v175
	var v176 int32
	_ = v176
	var v177 int32
	_ = v177
	var v179 int32
	_ = v179
	var v181 int32
	_ = v181
	var v182 int32
	_ = v182
	var v183 int32
	_ = v183
	var v184 int32
	_ = v184
	var v185 int32
	_ = v185
	var v189 int64
	_ = v189
	var v191 int32
	_ = v191
	var v194 int32
	_ = v194
	var v195 int32
	_ = v195
	var v196 int32
	_ = v196
	var v197 int32
	_ = v197
	var v202 int32
	_ = v202
	var v204 int32
	_ = v204
	var v205 int32
	_ = v205
	var v208 int32
	_ = v208
	var v210 int32
	_ = v210
	var v212 int32
	_ = v212
	var v223 int32
	_ = v223
	var v224 int32
	_ = v224
	var v225 int32
	_ = v225
	var v226 int32
	_ = v226
	var v227 int32
	_ = v227
	var v228 int32
	_ = v228
	var v229 int32
	_ = v229
	var v231 int32
	_ = v231
	var v236 int32
	_ = v236
	var v238 int32
	_ = v238
	var v241 int32
	_ = v241
	var v242 int32
	_ = v242
	var v243 int32
	_ = v243
	var v252 int32
	_ = v252
	var v253 int32
	_ = v253
	var v257 int32
	_ = v257
	var v262 int32
	_ = v262
	var v265 int32
	_ = v265
	var v267 int32
	_ = v267
	var v272 int32
	_ = v272
	var v275 int32
	_ = v275
	var v277 int32
	_ = v277
	var v284 int32
	_ = v284
	var v287 int32
	_ = v287
	var v292 int32
	_ = v292
	var v297 int32
	_ = v297
	var v298 int32
	_ = v298
	var v312 int32
	_ = v312
	var v313 int32
	_ = v313
	var v314 int32
	_ = v314
	var v316 int32
	_ = v316
	var v318 int32
	_ = v318
	var v319 int32
	_ = v319
	var v321 int32
	_ = v321
	var v324 int32
	_ = v324
	var v326 int32
	_ = v326
	var v327 int32
	_ = v327
	var v336 int32
	_ = v336
	var v337 int32
	_ = v337
	var v338 int32
	_ = v338
	var v339 int32
	_ = v339
	var v341 int32
	_ = v341
	var v344 int32
	_ = v344
	var v348 int32
	_ = v348
	var v351 int32
	_ = v351
	var v353 int32
	_ = v353
	var v354 int32
	_ = v354
	var v361 int32
	_ = v361
	var v364 int32
	_ = v364
	var v367 int32
	_ = v367
	var v373 int32
	_ = v373
	var v379 int32
	_ = v379
	var v383 int32
	_ = v383
	var v384 int32
	_ = v384
	var v385 int32
	_ = v385
	var v389 int32
	_ = v389
	var v391 int32
	_ = v391
	var v392 int32
	_ = v392
	var v404 int32
	_ = v404
	var v415 int32
	_ = v415
	var v416 int32
	_ = v416
	var v417 int32
	_ = v417
	var v418 int32
	_ = v418
	var v419 int32
	_ = v419
	var v428 int32
	_ = v428
	var v430 int32
	_ = v430
	var v431 int32
	_ = v431
	var v433 int32
	_ = v433
	var v443 int32
	_ = v443
	var v444 int32
	_ = v444
	var v448 int32
	_ = v448
	var v450 int32
	_ = v450
	var v451 int32
	_ = v451
	var v452 int32
	_ = v452
	var v454 int32
	_ = v454
	var v456 int32
	_ = v456
	var v459 int32
	_ = v459
	var v461 int32
	_ = v461
	var v462 int32
	_ = v462
	var v464 int32
	_ = v464
	var v467 int32
	_ = v467
	var v476 int32
	_ = v476
	var v479 int32
	_ = v479
	var v490 int32
	_ = v490
	var v492 int32
	_ = v492
	var v493 int32
	_ = v493
	var v495 int32
	_ = v495
	var v504 int32
	_ = v504
	var v505 int32
	_ = v505
	var v509 int32
	_ = v509
	var v511 int32
	_ = v511
	var v516 int32
	_ = v516
	var v517 int32
	_ = v517
	var v522 int32
	_ = v522
	var v524 int32
	_ = v524
	var v526 int32
	_ = v526
	var v528 int32
	_ = v528
	var v530 int32
	_ = v530
	var v532 int32
	_ = v532
	var v533 int32
	_ = v533
	var v534 int32
	_ = v534
	var v536 int32
	_ = v536
	var v538 int32
	_ = v538
	var v539 int32
	_ = v539
	var v540 int32
	_ = v540
	var v542 int32
	_ = v542
	var v566 int32
	_ = v566
	var v567 int32
	_ = v567
	var v568 int32
	_ = v568
	var v570 int32
	_ = v570
	var v572 int32
	_ = v572
	var v574 int32
	_ = v574
	var v576 int32
	_ = v576
	var v578 int32
	_ = v578
	var v579 int32
	_ = v579
	var v581 int32
	_ = v581
	var v586 int32
	_ = v586
	var v595 int32
	_ = v595
	var v598 int32
	_ = v598
	var v609 int32
	_ = v609
	var v612 int32
	_ = v612
	var v613 int32
	_ = v613
	var v617 int32
	_ = v617
	var v619 int32
	_ = v619
	var v624 int32
	_ = v624
	var v625 int32
	_ = v625
	var v626 int32
	_ = v626
	var v630 int32
	_ = v630
	var v631 int32
	_ = v631
	var v633 int32
	_ = v633
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
	var v640 int32
	_ = v640
	var v646 int32
	_ = v646
	var v653 int32
	_ = v653
	var v656 int32
	_ = v656
	var v661 int32
	_ = v661
	var v666 int32
	_ = v666
	v10 = m.G0
	v12 = v10 - int32(48)
	m.G0 = v12
	v15 = *(*int32)(unsafe.Add(mBase, _consts[72]))
	v16 = int32(10)
	goto L3
L1:
	;
	if v50 != 0 {
		goto L15
	} else {
		goto L16
	}
L2:
	;
	goto L1
L3:
	;
	v23 = *(*int32)(unsafe.Add(mBase, _consts[161]))
	goto L6
L4:
	;
	v35 = int32(0)
	goto L12
L6:
	;
	goto L7
L7:
	;
	goto L9
L9:
	;
	if v23 == int32(15) {
		goto L4
	} else {
		goto L10
	}
L10:
	;
	if v23 <= v16 {
		v50 = int32(1)
		goto L2
	} else {
		goto L11
	}
L11:
	;
	goto L4
L12:
	;
	v39 = *(*int32)(unsafe.Add(mBase, _consts[162]))
	if v39 != int32(2) {
		v50 = v35
		goto L2
	} else {
		goto L13
	}
L13:
	;
	v43 = int32(*(*uint8)(unsafe.Add(mBase, _consts[163])))
	if v43 != 0 {
		v50 = v35
		goto L2
	} else {
		goto L14
	}
L14:
	;
	v47 = *(*int32)(unsafe.Add(mBase, _consts[164]))
	v50 = int32(0) | base.B2i32(v47 <= v16)
	goto L2
L15:
	;
	v54 = *(*int32)(unsafe.Add(mBase, _consts[72]))
	F_ShowTransactionStateRec(m, int32(257561), v54)
	mBase = m.M
	v56 = m.ExcPending
	if v56 != 0 {
		goto L18
	} else {
		goto L19
	}
L16:
	;
	goto L17
L17:
	;
	v57 = *(*int32)(unsafe.Add(mBase, uint32(v15)+20))
	if v57 == int32(2) {
		goto L20
	} else {
		goto L21
	}
L18:
	;
	return
L19:
	;
	goto L17
L20:
	;
	v88 = *(*int32)(unsafe.Add(mBase, uint32(v15)+8))
	v91 = *(*int32)(unsafe.Add(mBase, _consts[165]))
	if v91 != 0 {
		goto L29
	} else {
		goto L30
	}
L21:
	;
	v62 = F_errstart(m, int32(19), int32(0))
	mBase = m.M
	v63 = m.ExcPending
	if v63 != 0 {
		goto L18
	} else {
		goto L22
	}
L22:
	;
	if v62 == int32(0) {
		goto L20
	} else {
		goto L23
	}
L23:
	;
	v66 = *(*int32)(unsafe.Add(mBase, uint32(v15)+20))
	if base.Ui32(v66) <= base.Ui32(int32(5)) {
		goto L24
	} else {
		goto L25
	}
L24:
	;
	v73 = *(*int32)(unsafe.Add(mBase, uint32(v66<<(uint(int32(2))%32))+uint32(_consts[166])))
	v75 = v73
	goto L26
L25:
	;
	v75 = int32(541631)
	goto L26
L26:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v12)+32)) = v75
	F_errmsg_internal(m, int32(351416), v12+int32(32))
	mBase = m.M
	v81 = m.ExcPending
	if v81 != 0 {
		goto L18
	} else {
		goto L27
	}
L27:
	;
	F_errfinish(m, int32(491705), int32(5112), int32(257561))
	mBase = m.M
	v86 = m.ExcPending
	if v86 != 0 {
		goto L18
	} else {
		goto L28
	}
L28:
	;
	goto L20
L29:
	;
	v92 = *(*int32)(unsafe.Add(mBase, uint32(v15)+80))
	v93 = *(*int32)(unsafe.Add(mBase, uint32(v92)+8))
	v94 = v91
	goto L32
L30:
	;
	v119 = v88
	goto L31
L31:
	;
	F_AtEOSubXact_Parallel(m, int32(1), v119)
	mBase = m.M
	v121 = m.ExcPending
	if v121 != 0 {
		goto L18
	} else {
		goto L36
	}
L32:
	;
	v103 = *(*int32)(unsafe.Add(mBase, uint32(v94)))
	v105 = *(*int32)(unsafe.Add(mBase, uint32(v94)+8))
	v106 = *(*int32)(unsafe.Add(mBase, uint32(v94)+4))
	m.T0[v106].(func(*base.Module, int32, int32, int32, int32))(m, int32(3), v88, v93, v105)
	mBase = m.M
	v108 = m.ExcPending
	if v108 != 0 {
		goto L18
	} else {
		goto L34
	}
L33:
	;
	v109 = *(*int32)(unsafe.Add(mBase, uint32(v15)+8))
	v119 = v109
	goto L31
L34:
	;
	if v103 != 0 {
		v94 = v103
		goto L32
	} else {
		goto L35
	}
L35:
	;
	goto L33
L36:
	;
	v122 = *(*int32)(unsafe.Add(mBase, uint32(v15)+72))
	if v122 != 0 {
		goto L37
	} else {
		goto L38
	}
L37:
	;
	v125 = F_errstart(m, int32(19), int32(0))
	mBase = m.M
	v126 = m.ExcPending
	if v126 != 0 {
		goto L18
	} else {
		goto L40
	}
L38:
	;
	goto L39
L39:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v15)+20)) = int32(3)
	F_CommandCounterIncrement(m)
	mBase = m.M
	v144 = m.ExcPending
	if v144 != 0 {
		goto L18
	} else {
		goto L46
	}
L40:
	;
	if v125 != 0 {
		goto L41
	} else {
		goto L42
	}
L41:
	;
	v127 = *(*int32)(unsafe.Add(mBase, uint32(v15)+72))
	*(*int32)(unsafe.Add(mBase, uint32(v12)+16)) = v127
	F_errmsg_internal(m, int32(255021), v12+int32(16))
	mBase = m.M
	v133 = m.ExcPending
	if v133 != 0 {
		goto L18
	} else {
		goto L44
	}
L42:
	;
	goto L43
L43:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v15)+72)) = int32(0)
	goto L39
L44:
	;
	F_errfinish(m, int32(491705), int32(5127), int32(257561))
	mBase = m.M
	v138 = m.ExcPending
	if v138 != 0 {
		goto L18
	} else {
		goto L45
	}
L45:
	;
	goto L43
L46:
	;
	v145 = *(*int32)(unsafe.Add(mBase, uint32(v15)))
	if v145 != 0 {
		goto L48
	} else {
		goto L49
	}
L47:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v653 = m.ExcPending
	if v653 != 0 {
		goto L18
	} else {
		goto L192
	}
L48:
	;
	v147 = *(*int32)(unsafe.Add(mBase, _consts[72]))
	v148 = *(*int32)(unsafe.Add(mBase, uint32(v147)+80))
	v149 = *(*int32)(unsafe.Add(mBase, uint32(v148)+56))
	v150 = *(*int32)(unsafe.Add(mBase, uint32(v148)+52))
	v151 = *(*int32)(unsafe.Add(mBase, uint32(v147)+52))
	v154 = v150 + v151 + int32(1)
	if v149 < v154 {
		goto L51
	} else {
		goto L52
	}
L49:
	;
	goto L50
L50:
	;
	F_AfterTriggerEndSubXact(m, int32(1))
	mBase = m.M
	v223 = m.ExcPending
	if v223 != 0 {
		goto L18
	} else {
		goto L75
	}
L51:
	;
	v156 = int32(268435455)
	v158 = v154 << (uint(int32(1)) % 32)
	if v156 <= v158 {
		goto L54
	} else {
		goto L55
	}
L52:
	;
	v183 = v148
	v184 = v150
	goto L53
L53:
	;
	v185 = *(*int32)(unsafe.Add(mBase, uint32(v183)+48))
	v189 = *(*int64)(unsafe.Add(mBase, uint32(v147)))
	*(*uint32)(unsafe.Add(mBase, uint32(v185+v184<<(uint(int32(2))%32)))) = uint32(v189)
	v191 = *(*int32)(unsafe.Add(mBase, uint32(v147)+52))
	if int32(0) < v191 {
		goto L64
	} else {
		goto L65
	}
L54:
	;
	v161 = v156
	goto L56
L55:
	;
	v161 = v158
	goto L56
L56:
	;
	if v161 < v154 {
		goto L47
	} else {
		goto L57
	}
L57:
	;
	v163 = *(*int32)(unsafe.Add(mBase, uint32(v148)+48))
	if v163 == int32(0) {
		goto L59
	} else {
		goto L60
	}
L58:
	;
	v177 = *(*int32)(unsafe.Add(mBase, uint32(v147)+80))
	*(*int32)(unsafe.Add(mBase, uint32(v177)+48)) = v176
	v179 = *(*int32)(unsafe.Add(mBase, uint32(v147)+80))
	*(*int32)(unsafe.Add(mBase, uint32(v179)+56)) = v161
	v181 = *(*int32)(unsafe.Add(mBase, uint32(v147)+80))
	v182 = *(*int32)(unsafe.Add(mBase, uint32(v181)+52))
	v183 = v181
	v184 = v182
	goto L53
L59:
	;
	v167 = *(*int32)(unsafe.Add(mBase, _consts[105]))
	v170 = F_MemoryContextAlloc(m, v167, v161<<(uint(int32(2))%32))
	mBase = m.M
	v171 = m.ExcPending
	if v171 != 0 {
		goto L18
	} else {
		goto L62
	}
L60:
	;
	goto L61
L61:
	;
	v174 = F_repalloc(m, v163, v161<<(uint(int32(2))%32))
	mBase = m.M
	v175 = m.ExcPending
	if v175 != 0 {
		goto L18
	} else {
		goto L63
	}
L62:
	;
	v176 = v170
	goto L58
L63:
	;
	v176 = v174
	goto L58
L64:
	;
	v194 = *(*int32)(unsafe.Add(mBase, uint32(v147)+80))
	v195 = *(*int32)(unsafe.Add(mBase, uint32(v194)+48))
	v196 = *(*int32)(unsafe.Add(mBase, uint32(v194)+52))
	v197 = int32(2)
	v202 = *(*int32)(unsafe.Add(mBase, uint32(v147)+48))
	v204 = v191 << (uint(v197) % 32)
	if v204 != 0 {
		goto L68
	} else {
		goto L69
	}
L65:
	;
	goto L66
L66:
	;
	v208 = *(*int32)(unsafe.Add(mBase, uint32(v147)+80))
	*(*int32)(unsafe.Add(mBase, uint32(v208)+52)) = v154
	v210 = *(*int32)(unsafe.Add(mBase, uint32(v147)+48))
	if v210 != 0 {
		goto L71
	} else {
		goto L72
	}
L67:
	;
	goto L66
L68:
	;
	v205 = F__emscripten_memcpy_bulkmem(m, v195+v196<<(uint(v197)%32)+int32(4), v202, v204)
	mBase = m.M
	goto L70
L69:
	;
	goto L70
L70:
	;
	goto L67
L71:
	;
	F_pfree(m, v210)
	mBase = m.M
	v212 = m.ExcPending
	if v212 != 0 {
		goto L18
	} else {
		goto L74
	}
L72:
	;
	goto L73
L73:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v147)+56)) = int32(0)
	*(*int64)(unsafe.Add(mBase, uint32(v147)+48)) = int64(0)
	goto L50
L74:
	;
	goto L73
L75:
	;
	v224 = *(*int32)(unsafe.Add(mBase, uint32(v15)+8))
	v225 = *(*int32)(unsafe.Add(mBase, uint32(v15)+80))
	v226 = *(*int32)(unsafe.Add(mBase, uint32(v225)+8))
	v227 = *(*int32)(unsafe.Add(mBase, uint32(v225)+28))
	v228 = *(*int32)(unsafe.Add(mBase, uint32(v225)+40))
	v229 = m.G0
	v231 = v229 - int32(32)
	m.G0 = v231
	v236 = *(*int32)(unsafe.Add(mBase, _consts[167]))
	F_hash_seq_init(m, v231+int32(12), v236)
	mBase = m.M
	v238 = m.ExcPending
	if v238 != 0 {
		goto L18
	} else {
		goto L76
	}
L76:
	;
	v241 = F_hash_seq_search(m, v231+int32(12))
	mBase = m.M
	v242 = m.ExcPending
	if v242 != 0 {
		goto L18
	} else {
		goto L77
	}
L77:
	;
	if v241 != 0 {
		goto L78
	} else {
		goto L79
	}
L78:
	;
	v243 = v241
	goto L81
L79:
	;
	goto L80
L80:
	;
	m.G0 = v231 + int32(32)
	v312 = *(*int32)(unsafe.Add(mBase, uint32(v15)+8))
	v313 = *(*int32)(unsafe.Add(mBase, uint32(v15)+80))
	v314 = *(*int32)(unsafe.Add(mBase, uint32(v313)+8))
	F_AtEOSubXact_LargeObject(m, int32(1), v312, v314)
	mBase = m.M
	v316 = m.ExcPending
	if v316 != 0 {
		goto L18
	} else {
		goto L104
	}
L81:
	;
	v252 = *(*int32)(unsafe.Add(mBase, uint32(v243)+64))
	v253 = *(*int32)(unsafe.Add(mBase, uint32(v252)+20))
	if v253 != v224 {
		goto L83
	} else {
		goto L84
	}
L82:
	;
	goto L80
L83:
	;
	v292 = *(*int32)(unsafe.Add(mBase, uint32(v252)+24))
	if v224 == v292 {
		goto L99
	} else {
		goto L100
	}
L84:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v252)+28)) = v227
	*(*int32)(unsafe.Add(mBase, uint32(v252)+20)) = v226
	v257 = *(*int32)(unsafe.Add(mBase, uint32(v252)+12))
	if v257 == int32(0) {
		goto L83
	} else {
		goto L85
	}
L85:
	;
	v262 = *(*int32)(unsafe.Add(mBase, uint32(v257)))
	if v262 == int32(0) {
		goto L87
	} else {
		goto L88
	}
L86:
	;
	goto L83
L87:
	;
	if v228 != 0 {
		goto L96
	} else {
		goto L97
	}
L88:
	;
	v265 = *(*int32)(unsafe.Add(mBase, uint32(v262)+4))
	if v265 == v257 {
		goto L89
	} else {
		goto L90
	}
L89:
	;
	v267 = *(*int32)(unsafe.Add(mBase, uint32(v257)+8))
	*(*int32)(unsafe.Add(mBase, uint32(v262)+4)) = v267
	goto L87
L90:
	;
	goto L91
L91:
	;
	v272 = v265
	goto L92
L92:
	;
	if v272 == int32(0) {
		goto L87
	} else {
		goto L94
	}
L93:
	;
	v277 = *(*int32)(unsafe.Add(mBase, uint32(v257)+8))
	*(*int32)(unsafe.Add(mBase, uint32(v272)+8)) = v277
	goto L87
L94:
	;
	v275 = *(*int32)(unsafe.Add(mBase, uint32(v272)+8))
	if v257 != v275 {
		v272 = v275
		goto L92
	} else {
		goto L95
	}
L95:
	;
	goto L93
L96:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v257))) = v228
	v284 = *(*int32)(unsafe.Add(mBase, uint32(v228)+4))
	*(*int32)(unsafe.Add(mBase, uint32(v257)+8)) = v284
	*(*int32)(unsafe.Add(mBase, uint32(v228)+4)) = v257
	goto L86
L97:
	;
	goto L98
L98:
	;
	v287 = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v257)+8)) = v287
	*(*int32)(unsafe.Add(mBase, uint32(v257))) = v287
	goto L86
L99:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v252)+24)) = v226
	goto L101
L100:
	;
	goto L101
L101:
	;
	v297 = F_hash_seq_search(m, v231+int32(12))
	mBase = m.M
	v298 = m.ExcPending
	if v298 != 0 {
		goto L18
	} else {
		goto L102
	}
L102:
	;
	if v297 != 0 {
		v243 = v297
		goto L81
	} else {
		goto L103
	}
L103:
	;
	goto L82
L104:
	;
	v318 = *(*int32)(unsafe.Add(mBase, _consts[72]))
	v319 = *(*int32)(unsafe.Add(mBase, uint32(v318)+28))
	goto L105
L105:
	;
	v321 = *(*int32)(unsafe.Add(mBase, _consts[168]))
	if v321 == int32(0) {
		goto L106
	} else {
		goto L107
	}
L106:
	;
	v348 = *(*int32)(unsafe.Add(mBase, _consts[169]))
	if v348 == int32(0) {
		goto L116
	} else {
		goto L117
	}
L107:
	;
	v324 = *(*int32)(unsafe.Add(mBase, uint32(v321)))
	if v324 < v319 {
		goto L106
	} else {
		goto L108
	}
L108:
	;
	v326 = *(*int32)(unsafe.Add(mBase, uint32(v321)+8))
	if v326 != 0 {
		goto L110
	} else {
		goto L111
	}
L109:
	;
	*(*int32)(unsafe.Add(mBase, _consts[168])) = v326
	v336 = *(*int32)(unsafe.Add(mBase, uint32(v326)+4))
	v337 = *(*int32)(unsafe.Add(mBase, uint32(v321)+4))
	v338 = F_list_concat(m, v336, v337)
	mBase = m.M
	v339 = m.ExcPending
	if v339 != 0 {
		goto L18
	} else {
		goto L114
	}
L110:
	;
	v327 = *(*int32)(unsafe.Add(mBase, uint32(v326)))
	if v319-int32(1) <= v327 {
		goto L109
	} else {
		goto L113
	}
L111:
	;
	goto L112
L112:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v321))) = v324 - int32(1)
	goto L106
L113:
	;
	goto L112
L114:
	;
	v341 = *(*int32)(unsafe.Add(mBase, _consts[168]))
	*(*int32)(unsafe.Add(mBase, uint32(v341)+4)) = v338
	F_pfree(m, v321)
	mBase = m.M
	v344 = m.ExcPending
	if v344 != 0 {
		goto L18
	} else {
		goto L115
	}
L115:
	;
	goto L106
L116:
	;
	v415 = *(*int32)(unsafe.Add(mBase, _consts[165]))
	if v415 != 0 {
		goto L136
	} else {
		goto L137
	}
L117:
	;
	v351 = *(*int32)(unsafe.Add(mBase, uint32(v348)))
	if v351 < v319 {
		goto L116
	} else {
		goto L118
	}
L118:
	;
	v353 = *(*int32)(unsafe.Add(mBase, uint32(v348)+12))
	if v353 != 0 {
		goto L120
	} else {
		goto L121
	}
L119:
	;
	v361 = int32(0)
	*(*int32)(unsafe.Add(mBase, _consts[169])) = v353
	v364 = *(*int32)(unsafe.Add(mBase, uint32(v348)+4))
	if v364 == v361 {
		goto L124
	} else {
		goto L125
	}
L120:
	;
	v354 = *(*int32)(unsafe.Add(mBase, uint32(v353)))
	if v319-int32(1) <= v354 {
		goto L119
	} else {
		goto L123
	}
L121:
	;
	goto L122
L122:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v348))) = v351 - int32(1)
	goto L116
L123:
	;
	goto L122
L124:
	;
	F_pfree(m, v348)
	mBase = m.M
	v404 = m.ExcPending
	if v404 != 0 {
		goto L18
	} else {
		goto L135
	}
L125:
	;
	v367 = *(*int32)(unsafe.Add(mBase, uint32(v364)+4))
	if v367 <= int32(0) {
		goto L124
	} else {
		goto L126
	}
L126:
	;
	v373 = v361
	goto L127
L127:
	;
	v379 = *(*int32)(unsafe.Add(mBase, uint32(v364)+12))
	v383 = *(*int32)(unsafe.Add(mBase, uint32(v379+v373<<(uint(int32(2))%32))))
	v384 = F_AsyncExistsPendingNotify(m, v383)
	mBase = m.M
	v385 = m.ExcPending
	if v385 != 0 {
		goto L18
	} else {
		goto L129
	}
L128:
	;
	goto L124
L129:
	;
	if v384 == int32(0) {
		goto L130
	} else {
		goto L131
	}
L130:
	;
	F_AddEventToPendingNotifies(m, v383)
	mBase = m.M
	v389 = m.ExcPending
	if v389 != 0 {
		goto L18
	} else {
		goto L133
	}
L131:
	;
	goto L132
L132:
	;
	v391 = v373 + int32(1)
	v392 = *(*int32)(unsafe.Add(mBase, uint32(v364)+4))
	if v391 < v392 {
		v373 = v391
		goto L127
	} else {
		goto L134
	}
L133:
	;
	goto L132
L134:
	;
	goto L128
L135:
	;
	goto L116
L136:
	;
	v416 = *(*int32)(unsafe.Add(mBase, uint32(v15)+8))
	v417 = *(*int32)(unsafe.Add(mBase, uint32(v15)+80))
	v418 = *(*int32)(unsafe.Add(mBase, uint32(v417)+8))
	v419 = v415
	goto L139
L137:
	;
	goto L138
L138:
	;
	v443 = *(*int32)(unsafe.Add(mBase, uint32(v15)+40))
	v444 = int32(1)
	F_ResourceOwnerRelease(m, v443, v444, v444, int32(0))
	mBase = m.M
	v448 = m.ExcPending
	if v448 != 0 {
		goto L18
	} else {
		goto L143
	}
L139:
	;
	v428 = *(*int32)(unsafe.Add(mBase, uint32(v419)))
	v430 = *(*int32)(unsafe.Add(mBase, uint32(v419)+8))
	v431 = *(*int32)(unsafe.Add(mBase, uint32(v419)+4))
	m.T0[v431].(func(*base.Module, int32, int32, int32, int32))(m, int32(1), v416, v418, v430)
	mBase = m.M
	v433 = m.ExcPending
	if v433 != 0 {
		goto L18
	} else {
		goto L141
	}
L140:
	;
	goto L138
L141:
	;
	if v428 != 0 {
		v419 = v428
		goto L139
	} else {
		goto L142
	}
L142:
	;
	goto L140
L143:
	;
	v450 = *(*int32)(unsafe.Add(mBase, uint32(v15)+8))
	v451 = *(*int32)(unsafe.Add(mBase, uint32(v15)+80))
	v452 = *(*int32)(unsafe.Add(mBase, uint32(v451)+8))
	F_AtEOSubXact_RelationCache(m, int32(1), v450, v452)
	mBase = m.M
	v454 = m.ExcPending
	if v454 != 0 {
		goto L18
	} else {
		goto L144
	}
L144:
	;
	F_AtEOXact_TypeCache(m)
	mBase = m.M
	v456 = m.ExcPending
	if v456 != 0 {
		goto L18
	} else {
		goto L145
	}
L145:
	;
	F_AtEOSubXact_Inval(m, int32(1))
	mBase = m.M
	v459 = m.ExcPending
	if v459 != 0 {
		goto L18
	} else {
		goto L146
	}
L146:
	;
	v461 = *(*int32)(unsafe.Add(mBase, _consts[72]))
	v462 = *(*int32)(unsafe.Add(mBase, uint32(v461)+28))
	goto L147
L147:
	;
	v464 = *(*int32)(unsafe.Add(mBase, _consts[138]))
	if v464 != 0 {
		goto L148
	} else {
		goto L149
	}
L148:
	;
	v467 = v464
	goto L151
L149:
	;
	goto L150
L150:
	;
	v490 = *(*int32)(unsafe.Add(mBase, uint32(v15)+40))
	*(*int32)(unsafe.Add(mBase, _consts[170])) = v490
	v492 = *(*int32)(unsafe.Add(mBase, uint32(v15)))
	if v492 != 0 {
		goto L157
	} else {
		goto L158
	}
L151:
	;
	v476 = *(*int32)(unsafe.Add(mBase, uint32(v467)+20))
	if v462 <= v476 {
		goto L153
	} else {
		goto L154
	}
L152:
	;
	goto L150
L153:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v467)+20)) = v462 - int32(1)
	goto L155
L154:
	;
	goto L155
L155:
	;
	v479 = *(*int32)(unsafe.Add(mBase, uint32(v467)+24))
	if v479 != 0 {
		v467 = v479
		goto L151
	} else {
		goto L156
	}
L156:
	;
	goto L152
L157:
	;
	v493 = m.G0
	v495 = v493 - int32(16)
	m.G0 = v495
	*(*int32)(unsafe.Add(mBase, uint32(v495)+12)) = int32(17104896)
	*(*int64)(unsafe.Add(mBase, uint32(v495)+4)) = int64(0)
	*(*int32)(unsafe.Add(mBase, uint32(v495))) = v492
	v504 = F_LockRelease(m, v495, int32(7), int32(0))
	mBase = m.M
	v505 = m.ExcPending
	if v505 != 0 {
		goto L18
	} else {
		goto L160
	}
L158:
	;
	v511 = v490
	goto L159
L159:
	;
	F_ResourceOwnerRelease(m, v511, int32(2), int32(1), int32(0))
	mBase = m.M
	v516 = m.ExcPending
	if v516 != 0 {
		goto L18
	} else {
		goto L161
	}
L160:
	;
	m.G0 = v495 + int32(16)
	v509 = *(*int32)(unsafe.Add(mBase, uint32(v15)+40))
	v511 = v509
	goto L159
L161:
	;
	v517 = *(*int32)(unsafe.Add(mBase, uint32(v15)+40))
	F_ResourceOwnerRelease(m, v517, int32(3), int32(1), int32(0))
	mBase = m.M
	v522 = m.ExcPending
	if v522 != 0 {
		goto L18
	} else {
		goto L162
	}
L162:
	;
	v524 = *(*int32)(unsafe.Add(mBase, uint32(v15)+32))
	F_AtEOXact_GUC(m, int32(1), v524)
	mBase = m.M
	v526 = m.ExcPending
	if v526 != 0 {
		goto L18
	} else {
		goto L163
	}
L163:
	;
	v528 = *(*int32)(unsafe.Add(mBase, uint32(v15)+8))
	F_AtEOSubXact_SPI(m, int32(1), v528)
	mBase = m.M
	v530 = m.ExcPending
	if v530 != 0 {
		goto L18
	} else {
		goto L164
	}
L164:
	;
	v532 = *(*int32)(unsafe.Add(mBase, uint32(v15)+8))
	v533 = *(*int32)(unsafe.Add(mBase, uint32(v15)+80))
	v534 = *(*int32)(unsafe.Add(mBase, uint32(v533)+8))
	F_AtEOSubXact_on_commit_actions(m, int32(1), v532, v534)
	mBase = m.M
	v536 = m.ExcPending
	if v536 != 0 {
		goto L18
	} else {
		goto L165
	}
L165:
	;
	v538 = *(*int32)(unsafe.Add(mBase, uint32(v15)+8))
	v539 = *(*int32)(unsafe.Add(mBase, uint32(v15)+80))
	v540 = *(*int32)(unsafe.Add(mBase, uint32(v539)+8))
	v542 = *(*int32)(unsafe.Add(mBase, _consts[171]))
	if v538 == v542 {
		goto L167
	} else {
		goto L168
	}
L166:
	;
	v566 = *(*int32)(unsafe.Add(mBase, uint32(v15)+8))
	v567 = *(*int32)(unsafe.Add(mBase, uint32(v15)+80))
	v568 = *(*int32)(unsafe.Add(mBase, uint32(v567)+8))
	F_AtEOSubXact_Files(m, int32(1), v566, v568)
	mBase = m.M
	v570 = m.ExcPending
	if v570 != 0 {
		goto L18
	} else {
		goto L173
	}
L167:
	;
	goto L170
L168:
	;
	goto L169
L169:
	;
	goto L166
L170:
	;
	*(*int32)(unsafe.Add(mBase, _consts[171])) = v540
	goto L166
L173:
	;
	v572 = *(*int32)(unsafe.Add(mBase, uint32(v15)+28))
	F_AtEOSubXact_HashTables(m, int32(1), v572)
	mBase = m.M
	v574 = m.ExcPending
	if v574 != 0 {
		goto L18
	} else {
		goto L174
	}
L174:
	;
	v576 = *(*int32)(unsafe.Add(mBase, uint32(v15)+28))
	F_AtEOSubXact_PgStat(m, int32(1), v576)
	mBase = m.M
	v578 = m.ExcPending
	if v578 != 0 {
		goto L18
	} else {
		goto L175
	}
L175:
	;
	v579 = *(*int32)(unsafe.Add(mBase, uint32(v15)+28))
	v581 = *(*int32)(unsafe.Add(mBase, _consts[104]))
	if v581 == int32(0) {
		goto L176
	} else {
		goto L177
	}
L176:
	;
	v609 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v15)+68)))
	*(*uint8)(unsafe.Add(mBase, _consts[172])) = uint8(v609)
	v612 = *(*int32)(unsafe.Add(mBase, uint32(v15)+80))
	v613 = *(*int32)(unsafe.Add(mBase, uint32(v612)+40))
	*(*int32)(unsafe.Add(mBase, _consts[173])) = v613
	*(*int32)(unsafe.Add(mBase, _consts[170])) = v613
	v617 = *(*int32)(unsafe.Add(mBase, uint32(v15)+40))
	F_ResourceOwnerDelete(m, v617)
	mBase = m.M
	v619 = m.ExcPending
	if v619 != 0 {
		goto L18
	} else {
		goto L182
	}
L177:
	;
	v586 = v581
	goto L178
L178:
	;
	v595 = *(*int32)(unsafe.Add(mBase, uint32(v586)+4))
	if v595 < v579 {
		goto L176
	} else {
		goto L180
	}
L179:
	;
	goto L176
L180:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v586)+4)) = v579 - int32(1)
	v598 = *(*int32)(unsafe.Add(mBase, uint32(v586)+8))
	if v598 != 0 {
		v586 = v598
		goto L178
	} else {
		goto L181
	}
L181:
	;
	goto L179
L182:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v15)+40)) = int32(0)
	v624 = *(*int32)(unsafe.Add(mBase, _consts[72]))
	v625 = *(*int32)(unsafe.Add(mBase, uint32(v624)+80))
	v626 = *(*int32)(unsafe.Add(mBase, uint32(v625)+36))
	*(*int32)(unsafe.Add(mBase, _consts[174])) = v626
	*(*int32)(unsafe.Add(mBase, _consts[0])) = v626
	v630 = *(*int32)(unsafe.Add(mBase, uint32(v624)+36))
	v631 = *(*int32)(unsafe.Add(mBase, uint32(v630)+20))
	if v631 != 0 {
		goto L183
	} else {
		goto L184
	}
L183:
	;
	v637 = int32(0)
	goto L185
L184:
	;
	v633 = *(*int32)(unsafe.Add(mBase, uint32(v630)+12))
	v634 = *(*int32)(unsafe.Add(mBase, uint32(v633)+28))
	v635 = m.T0[v634].(func(*base.Module, int32) int32)(m, v630)
	mBase = m.M
	v636 = m.ExcPending
	if v636 != 0 {
		goto L18
	} else {
		goto L186
	}
L185:
	;
	if v637 != 0 {
		goto L187
	} else {
		goto L188
	}
L186:
	;
	v637 = v635
	goto L185
L187:
	;
	v638 = *(*int32)(unsafe.Add(mBase, uint32(v624)+36))
	F_MemoryContextDelete(m, v638)
	mBase = m.M
	v640 = m.ExcPending
	if v640 != 0 {
		goto L18
	} else {
		goto L190
	}
L188:
	;
	goto L189
L189:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v15)+20)) = int32(0)
	F_PopTransaction(m)
	mBase = m.M
	v646 = m.ExcPending
	if v646 != 0 {
		goto L18
	} else {
		goto L191
	}
L190:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v624)+36)) = int32(0)
	goto L189
L191:
	;
	m.G0 = v12 + int32(48)
	return
L192:
	;
	F_errcode(m, int32(261))
	mBase = m.M
	v656 = m.ExcPending
	if v656 != 0 {
		goto L18
	} else {
		goto L193
	}
L193:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v12))) = int32(268435455)
	F_errmsg(m, int32(461112), v12)
	mBase = m.M
	v661 = m.ExcPending
	if v661 != 0 {
		goto L18
	} else {
		goto L194
	}
L194:
	;
	F_errfinish(m, int32(491705), int32(1696), int32(173116))
	mBase = m.M
	v666 = m.ExcPending
	if v666 != 0 {
		goto L18
	} else {
		goto L195
	}
L195:
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
	var v54 int32
	_ = v54
	var v56 int32
	_ = v56
	var v58 int32
	_ = v58
	var v63 int32
	_ = v63
	var v70 int32
	_ = v70
	var v77 int32
	_ = v77
	var v79 int32
	_ = v79
	var v81 int32
	_ = v81
	var v88 int32
	_ = v88
	var v90 int32
	_ = v90
	var v91 int32
	_ = v91
	var v100 int32
	_ = v100
	var v107 int32
	_ = v107
	var v108 int32
	_ = v108
	var v110 int32
	_ = v110
	var v116 int32
	_ = v116
	var v120 int32
	_ = v120
	var v123 int32
	_ = v123
	var v125 int32
	_ = v125
	var v135 int32
	_ = v135
	var v136 int32
	_ = v136
	var v137 int32
	_ = v137
	var v151 int32
	_ = v151
	var v166 int32
	_ = v166
	var v167 int32
	_ = v167
	var v172 int32
	_ = v172
	var v191 int32
	_ = v191
	var v194 int32
	_ = v194
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
	v54 = v46
	v56 = v35 + v48
	v58 = v35 + v50
	v63 = v39
	goto L9
L9:
	;
	v70 = v58 - int32(1)
	if v70 < int32(0) {
		v79 = v54
		goto L11
	} else {
		goto L12
	}
L10:
	;
	if v17 < v18 {
		goto L21
	} else {
		goto L22
	}
L11:
	;
	v81 = v56 - int32(1)
	if v81 < int32(0) {
		v90 = v79
		goto L14
	} else {
		goto L15
	}
L12:
	;
	if v22 <= v70 {
		v79 = v54
		goto L11
	} else {
		goto L13
	}
L13:
	;
	v77 = int32(*(*int16)(unsafe.Add(mBase, uint32(v20+v70<<(uint(int32(1))%32)))))
	v79 = v54 + v77
	goto L11
L14:
	;
	v91 = int32(1)
	if v90 < int32(0) {
		goto L17
	} else {
		goto L18
	}
L15:
	;
	if v27 <= v81 {
		v90 = v79
		goto L14
	} else {
		goto L16
	}
L16:
	;
	v88 = int32(*(*int16)(unsafe.Add(mBase, uint32(v19+v81<<(uint(int32(1))%32)))))
	v90 = v79 - v88
	goto L14
L17:
	;
	v100 = v90 + int32(10000)
	goto L19
L18:
	;
	v100 = v90
	goto L19
L19:
	;
	*(*uint16)(unsafe.Add(mBase, uint32(v44+v63<<(uint(v91)%32)))) = uint16(v100)
	if base.Ui32(int32(1)) < base.Ui32(v63) {
		v54 = v90 >> (uint(int32(31)) % 32)
		v56 = v81
		v58 = v70
		v63 = v63 - v91
		goto L9
	} else {
		goto L20
	}
L20:
	;
	goto L10
L21:
	;
	v107 = v18
	goto L23
L22:
	;
	v107 = v17
	goto L23
L23:
	;
	v108 = *(*int32)(unsafe.Add(mBase, uint32(l2)+16))
	if v108 != 0 {
		goto L24
	} else {
		goto L25
	}
L24:
	;
	F_pfree(m, v108)
	mBase = m.M
	v110 = m.ExcPending
	if v110 != 0 {
		goto L7
	} else {
		goto L27
	}
L25:
	;
	goto L26
L26:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l2)+16)) = v44
	*(*int32)(unsafe.Add(mBase, uint32(l2))) = v39
	*(*int32)(unsafe.Add(mBase, uint32(l2)+12)) = v107
	*(*int32)(unsafe.Add(mBase, uint32(l2)+4)) = v23
	v116 = v44 + int32(2)
	*(*int32)(unsafe.Add(mBase, uint32(l2)+20)) = v116
	v120 = v116
	v123 = v39
	v125 = v23
	goto L31
L27:
	;
	goto L26
L28:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l2))) = v194
	*(*int32)(unsafe.Add(mBase, uint32(l2)+20)) = v191
	return
L29:
	;
	*(*int64)(unsafe.Add(mBase, uint32(l2)+4)) = int64(0)
	v191 = v172
	v194 = int32(0)
	goto L28
L30:
	;
	v151 = v123
	goto L35
L31:
	;
	v135 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v120))))
	if v135 != 0 {
		goto L30
	} else {
		goto L33
	}
L32:
	;
	v172 = v116 + v41
	goto L29
L33:
	;
	v136 = int32(1)
	v137 = v125 - v136
	*(*int32)(unsafe.Add(mBase, uint32(l2)+4)) = v137
	if v136 < v123 {
		v120 = v120 + int32(2)
		v123 = v123 - v136
		v125 = v137
		goto L31
	} else {
		goto L34
	}
L34:
	;
	goto L32
L35:
	;
	v166 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v120-int32(2)+v151<<(uint(int32(1))%32)))))
	if v166 != 0 {
		v191 = v120
		v194 = v151
		goto L28
	} else {
		goto L37
	}
L36:
	;
	v172 = v120
	goto L29
L37:
	;
	v167 = int32(1)
	if v167 < v151 {
		v151 = v151 - v167
		goto L35
	} else {
		goto L38
	}
L38:
	;
	goto L36
}
