package p4

import (
	base "github.com/shibukawa/pgmem/internal/aot/pgaot/base"
	"unsafe"
)

func F_hstoreArrayToPairs(m *base.Module, l0 int32, l1 int32) int32 {
	mBase := m.M
	_ = mBase
	var v10 int32
	_ = v10
	var v12 int32
	_ = v12
	var v24 int32
	_ = v24
	var v25 int32
	_ = v25
	var v33 int32
	_ = v33
	var v34 int32
	_ = v34
	var v35 int32
	_ = v35
	var v39 int32
	_ = v39
	var v40 int32
	_ = v40
	var v41 int32
	_ = v41
	var v43 int32
	_ = v43
	var v44 int32
	_ = v44
	var v53 int32
	_ = v53
	var v58 int32
	_ = v58
	var v59 int32
	_ = v59
	var v62 int32
	_ = v62
	var v63 int32
	_ = v63
	var v66 int32
	_ = v66
	var v67 int32
	_ = v67
	var v71 int32
	_ = v71
	var v80 int32
	_ = v80
	var v84 int32
	_ = v84
	var v87 int32
	_ = v87
	var v97 int32
	_ = v97
	var v98 int32
	_ = v98
	var v102 int32
	_ = v102
	var v108 int32
	_ = v108
	var v117 int32
	_ = v117
	var v120 int32
	_ = v120
	var v123 int32
	_ = v123
	var v129 int32
	_ = v129
	var v136 int32
	_ = v136
	v10 = m.G0
	v12 = v10 - int32(32)
	m.G0 = v12
	F_deconstruct_array_builtin(m, l0, int32(25), v12+int32(28), v12+int32(24), v12+int32(20))
	mBase = m.M
	v24 = m.ExcPending
	if v24 != 0 {
		return int32(0)
	} else {
		v25 = *(*int32)(unsafe.Add(mBase, uint32(v12)+20))
		if v25 == int32(0) {
			v102 = int32(0)
			v108 = int32(0)
			*(*int32)(unsafe.Add(mBase, uint32(l1))) = v108
			m.G0 = v12 + int32(32)
			return v102
		} else {
			if base.Ui32(int32(53687092)) <= base.Ui32(v25) {
				F_errstart_cold(m, int32(21), int32(0))
				mBase = m.M
				v117 = m.ExcPending
				if v117 != 0 {
					return int32(0)
				} else {
					F_errcode(m, int32(261))
					mBase = m.M
					v120 = m.ExcPending
					if v120 != 0 {
						return int32(0)
					} else {
						*(*int32)(unsafe.Add(mBase, uint32(v12)+4)) = int32(53687091)
						v123 = *(*int32)(unsafe.Add(mBase, uint32(v12)+20))
						*(*int32)(unsafe.Add(mBase, uint32(v12))) = v123
						F_errmsg(m, int32(655860), v12)
						mBase = m.M
						v129 = m.ExcPending
						if v129 != 0 {
							return int32(0)
						} else {
							F_errfinish(m, int32(488210), int32(102), int32(129648))
							mBase = m.M
							v136 = m.ExcPending
							if v136 != 0 {
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
				v33 = F_palloc(m, v25*int32(20))
				mBase = m.M
				v34 = m.ExcPending
				if v34 != 0 {
					return int32(0)
				} else {
					v35 = *(*int32)(unsafe.Add(mBase, uint32(v12)+20))
					if v35 <= int32(0) {
						v87 = int32(0)
					} else {
						v39 = *(*int32)(unsafe.Add(mBase, uint32(v12)+28))
						v40 = *(*int32)(unsafe.Add(mBase, uint32(v12)+24))
						v41 = int32(0)
						v43 = v41
						v44 = v41
						for {
							v53 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v43+v40))))
							if v53 == int32(0) {
								v58 = v33 + v44*int32(20)
								v59 = int32(2)
								v62 = *(*int32)(unsafe.Add(mBase, uint32(v39+v43<<(uint(v59)%32))))
								v63 = int32(4)
								*(*int32)(unsafe.Add(mBase, uint32(v58))) = v62 + v63
								v66 = *(*int32)(unsafe.Add(mBase, uint32(v62)))
								v67 = int32(0)
								*(*int32)(unsafe.Add(mBase, uint32(v58)+12)) = v67
								*(*int32)(unsafe.Add(mBase, uint32(v58)+4)) = v67
								v71 = int32(1)
								*(*uint16)(unsafe.Add(mBase, uint32(v58)+16)) = uint16(v71)
								*(*int32)(unsafe.Add(mBase, uint32(v58)+8)) = int32(base.Ui32(v66)>>(uint(v59)%32)) - v63
								v80 = v44 + v71
							} else {
								v80 = v44
							}
							v84 = v43 + int32(1)
							if v84 != v35 {
								v43 = v84
								v44 = v80
								continue
							} else {
								break
							}
							break
						}
						v87 = v80
					}
					v97 = F_hstoreUniquePairs(m, v33, v87, v12+int32(16))
					mBase = m.M
					v98 = m.ExcPending
					if v98 != 0 {
						return int32(0)
					} else {
						v102 = v33
						v108 = v97
						*(*int32)(unsafe.Add(mBase, uint32(l1))) = v108
						m.G0 = v12 + int32(32)
						return v102
					}
				}
			}
		}
	}
}
func F_hstoreValidNewFormat(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v12 int32
	_ = v12
	var v13 int32
	_ = v13
	var v15 int32
	_ = v15
	var v21 int32
	_ = v21
	var v22 int32
	_ = v22
	var v27 int32
	_ = v27
	var v29 int32
	_ = v29
	var v33 int32
	_ = v33
	var v38 int32
	_ = v38
	var v39 int32
	_ = v39
	var v41 int32
	_ = v41
	var v43 int32
	_ = v43
	var v47 int32
	_ = v47
	var v59 int32
	_ = v59
	var v60 int32
	_ = v60
	var v63 int32
	_ = v63
	var v67 int32
	_ = v67
	var v72 int32
	_ = v72
	var v76 int32
	_ = v76
	var v79 int32
	_ = v79
	var v84 int32
	_ = v84
	var v93 int32
	_ = v93
	var v94 int32
	_ = v94
	var v95 int32
	_ = v95
	var v97 int32
	_ = v97
	var v102 int32
	_ = v102
	var v106 int32
	_ = v106
	var v107 int32
	_ = v107
	var v108 int32
	_ = v108
	var v110 int32
	_ = v110
	var v115 int32
	_ = v115
	var v119 int32
	_ = v119
	var v120 int32
	_ = v120
	var v125 int32
	_ = v125
	var v141 int32
	_ = v141
	var v144 int32
	_ = v144
	v12 = int32(2)
	v13 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v15 = v13 & int32(268435455)
	if v15 == int32(0) {
		v144 = v12
		goto L1
	} else {
		goto L2
	}
L1:
	;
	return v144
L2:
	;
	if v13 < int32(0) {
		v144 = v12
		goto L1
	} else {
		goto L3
	}
L3:
	;
	v21 = l0 + int32(8)
	v22 = *(*int32)(unsafe.Add(mBase, uint32(v21)))
	if int32(0) <= v22 {
		goto L4
	} else {
		goto L5
	}
L4:
	;
	return int32(0)
L5:
	;
	goto L6
L6:
	;
	v27 = int32(0)
	v29 = v15 << (uint(int32(3)) % 32)
	v33 = *(*int32)(unsafe.Add(mBase, uint32(v29+v21-int32(4))))
	v38 = v29 + v33&int32(1073741823) + int32(8)
	v39 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v41 = int32(base.Ui32(v39) >> (uint(int32(2)) % 32))
	if base.Ui32(v41) < base.Ui32(v38) {
		v144 = v27
		goto L1
	} else {
		goto L7
	}
L7:
	;
	v43 = int32(1)
	v47 = v43
	goto L8
L8:
	;
	v59 = v21 + v47<<(uint(int32(2))%32)
	v60 = *(*int32)(unsafe.Add(mBase, uint32(v59)))
	if v60 < int32(0) {
		v144 = v27
		goto L1
	} else {
		goto L10
	}
L9:
	;
	if v15 != int32(1) {
		goto L13
	} else {
		goto L14
	}
L10:
	;
	v63 = int32(1073741823)
	v67 = *(*int32)(unsafe.Add(mBase, uint32(v59-int32(4))))
	if base.Ui32(v60&v63) < base.Ui32(v67&v63) {
		v144 = v27
		goto L1
	} else {
		goto L11
	}
L11:
	;
	v72 = v47 + int32(1)
	if v72 != v15<<(uint(v43)%32) {
		v47 = v72
		goto L8
	} else {
		goto L12
	}
L12:
	;
	goto L9
L13:
	;
	v76 = int32(2)
	if base.Ui32(v15) <= base.Ui32(v76) {
		goto L16
	} else {
		goto L17
	}
L14:
	;
	goto L15
L15:
	;
	if v38 == v41 {
		goto L30
	} else {
		goto L31
	}
L16:
	;
	v79 = v76
	goto L18
L17:
	;
	v79 = v15
	goto L18
L18:
	;
	v84 = int32(1)
	goto L19
L19:
	;
	v93 = v84 << (uint(int32(3)) % 32)
	v94 = v21 + v93
	v95 = *(*int32)(unsafe.Add(mBase, uint32(v94)))
	v97 = v95 & int32(1073741823)
	if int32(0) <= v95 {
		goto L21
	} else {
		goto L22
	}
L20:
	;
	goto L15
L21:
	;
	v102 = *(*int32)(unsafe.Add(mBase, uint32(v94-int32(4))))
	v106 = v97 - v102&int32(1073741823)
	goto L23
L22:
	;
	v106 = v97
	goto L23
L23:
	;
	v107 = l0 + v93
	v108 = *(*int32)(unsafe.Add(mBase, uint32(v107)))
	v110 = v108 & int32(1073741823)
	if int32(0) <= v108 {
		goto L24
	} else {
		goto L25
	}
L24:
	;
	v115 = *(*int32)(unsafe.Add(mBase, uint32(v107-int32(4))))
	v119 = v110 - v115&int32(1073741823)
	goto L26
L25:
	;
	v119 = v110
	goto L26
L26:
	;
	v120 = int32(0)
	if v95&int32(1073741824) != 0 {
		v144 = v120
		goto L1
	} else {
		goto L27
	}
L27:
	;
	if base.Ui32(v106) < base.Ui32(v119) {
		v144 = v120
		goto L1
	} else {
		goto L28
	}
L28:
	;
	v125 = v84 + int32(1)
	if v125 != v79 {
		v84 = v125
		goto L19
	} else {
		goto L29
	}
L29:
	;
	goto L20
L30:
	;
	v141 = int32(2)
	goto L32
L31:
	;
	v141 = int32(1)
	goto L32
L32:
	;
	v144 = v141
	goto L1
}
func F_hstore_concat(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v23 int32
	_ = v23
	var v24 int32
	_ = v24
	var v27 int32
	_ = v27
	var v28 int32
	_ = v28
	var v29 int32
	_ = v29
	var v30 int32
	_ = v30
	var v31 int32
	_ = v31
	var v32 int32
	_ = v32
	var v34 int32
	_ = v34
	var v38 int32
	_ = v38
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
	var v48 int32
	_ = v48
	var v52 int32
	_ = v52
	var v61 int32
	_ = v61
	var v64 int32
	_ = v64
	var v66 int32
	_ = v66
	var v67 int32
	_ = v67
	var v68 int32
	_ = v68
	var v70 int32
	_ = v70
	var v72 int32
	_ = v72
	var v80 int32
	_ = v80
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
	var v96 int32
	_ = v96
	var v107 int32
	_ = v107
	var v108 int32
	_ = v108
	var v109 int32
	_ = v109
	var v110 int32
	_ = v110
	var v112 int32
	_ = v112
	var v114 int32
	_ = v114
	var v116 int32
	_ = v116
	var v117 int32
	_ = v117
	var v122 int32
	_ = v122
	var v123 int32
	_ = v123
	var v126 int32
	_ = v126
	var v131 int32
	_ = v131
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
	var v140 int32
	_ = v140
	var v143 int32
	_ = v143
	var v145 int32
	_ = v145
	var v157 int32
	_ = v157
	var v166 int32
	_ = v166
	var v168 int32
	_ = v168
	var v169 int32
	_ = v169
	var v170 int32
	_ = v170
	var v171 int32
	_ = v171
	var v174 int32
	_ = v174
	var v175 int32
	_ = v175
	var v177 int32
	_ = v177
	var v182 int32
	_ = v182
	var v186 int32
	_ = v186
	var v191 int32
	_ = v191
	var v195 int32
	_ = v195
	var v197 int32
	_ = v197
	var v203 int32
	_ = v203
	var v206 int32
	_ = v206
	var v207 int32
	_ = v207
	var v212 int32
	_ = v212
	var v215 int32
	_ = v215
	var v216 int32
	_ = v216
	var v224 int32
	_ = v224
	var v225 int32
	_ = v225
	var v226 int32
	_ = v226
	var v229 int32
	_ = v229
	var v230 int32
	_ = v230
	var v232 int32
	_ = v232
	var v233 int32
	_ = v233
	var v235 int32
	_ = v235
	var v237 int32
	_ = v237
	var v240 int32
	_ = v240
	var v241 int32
	_ = v241
	var v242 int32
	_ = v242
	var v247 int32
	_ = v247
	var v248 int32
	_ = v248
	var v249 int32
	_ = v249
	var v252 int32
	_ = v252
	var v253 int32
	_ = v253
	var v254 int32
	_ = v254
	var v257 int32
	_ = v257
	var v258 int32
	_ = v258
	var v260 int32
	_ = v260
	var v265 int32
	_ = v265
	var v278 int32
	_ = v278
	var v292 int32
	_ = v292
	var v293 int32
	_ = v293
	var v298 int32
	_ = v298
	var v302 int32
	_ = v302
	var v304 int32
	_ = v304
	var v308 int32
	_ = v308
	var v309 int32
	_ = v309
	var v314 int32
	_ = v314
	var v315 int32
	_ = v315
	var v317 int32
	_ = v317
	var v322 int32
	_ = v322
	var v326 int32
	_ = v326
	var v327 int32
	_ = v327
	var v328 int32
	_ = v328
	var v329 int32
	_ = v329
	var v330 int32
	_ = v330
	var v332 int32
	_ = v332
	var v337 int32
	_ = v337
	var v341 int32
	_ = v341
	var v342 int32
	_ = v342
	var v345 int32
	_ = v345
	var v348 int32
	_ = v348
	var v349 int32
	_ = v349
	var v351 int32
	_ = v351
	var v352 int32
	_ = v352
	var v354 int32
	_ = v354
	var v357 int32
	_ = v357
	var v368 int32
	_ = v368
	var v374 int32
	_ = v374
	var v375 int32
	_ = v375
	var v380 int32
	_ = v380
	var v384 int32
	_ = v384
	var v386 int32
	_ = v386
	var v390 int32
	_ = v390
	var v391 int32
	_ = v391
	var v396 int32
	_ = v396
	var v397 int32
	_ = v397
	var v399 int32
	_ = v399
	var v404 int32
	_ = v404
	var v408 int32
	_ = v408
	var v409 int32
	_ = v409
	var v410 int32
	_ = v410
	var v411 int32
	_ = v411
	var v412 int32
	_ = v412
	var v414 int32
	_ = v414
	var v419 int32
	_ = v419
	var v423 int32
	_ = v423
	var v424 int32
	_ = v424
	var v427 int32
	_ = v427
	var v430 int32
	_ = v430
	var v431 int32
	_ = v431
	var v433 int32
	_ = v433
	var v434 int32
	_ = v434
	var v436 int32
	_ = v436
	var v439 int32
	_ = v439
	var v451 int32
	_ = v451
	var v454 int32
	_ = v454
	var v459 int32
	_ = v459
	var v460 int32
	_ = v460
	var v462 int32
	_ = v462
	var v464 int32
	_ = v464
	var v465 int32
	_ = v465
	var v466 int32
	_ = v466
	var v468 int32
	_ = v468
	var v472 int32
	_ = v472
	var v473 int32
	_ = v473
	var v484 int32
	_ = v484
	var v488 int32
	_ = v488
	var v495 int32
	_ = v495
	var v498 int32
	_ = v498
	var v505 int32
	_ = v505
	var v506 int32
	_ = v506
	var v507 int32
	_ = v507
	var v511 int32
	_ = v511
	var v513 int32
	_ = v513
	var v514 int32
	_ = v514
	var v516 int32
	_ = v516
	var v518 int32
	_ = v518
	var v525 int32
	_ = v525
	var v531 int32
	_ = v531
	var v532 int32
	_ = v532
	var v534 int32
	_ = v534
	var v540 int32
	_ = v540
	var v547 int32
	_ = v547
	var v551 int32
	_ = v551
	var v554 int32
	_ = v554
	var v560 int32
	_ = v560
	var v567 int32
	_ = v567
	var v571 int32
	_ = v571
	var v574 int32
	_ = v574
	var v577 int32
	_ = v577
	var v578 int32
	_ = v578
	var v579 int32
	_ = v579
	var v584 int32
	_ = v584
	var v585 int32
	_ = v585
	var v586 int32
	_ = v586
	var v588 int32
	_ = v588
	var v590 int32
	_ = v590
	var v591 int32
	_ = v591
	var v593 int32
	_ = v593
	var v595 int32
	_ = v595
	var v599 int32
	_ = v599
	var v600 int32
	_ = v600
	var v601 int32
	_ = v601
	var v606 int32
	_ = v606
	var v607 int32
	_ = v607
	var v608 int32
	_ = v608
	var v610 int32
	_ = v610
	var v612 int32
	_ = v612
	var v617 int32
	_ = v617
	v23 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v24 = F_hstoreUpgrade(m, v23)
	mBase = m.M
	v27 = m.ExcPending
	if v27 != 0 {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	return int32(0)
L2:
	;
	v28 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
	v29 = F_hstoreUpgrade(m, v28)
	mBase = m.M
	v30 = m.ExcPending
	if v30 != 0 {
		goto L1
	} else {
		goto L3
	}
L3:
	;
	v31 = *(*int32)(unsafe.Add(mBase, uint32(v29)))
	v32 = int32(2)
	v34 = *(*int32)(unsafe.Add(mBase, uint32(v24)))
	v38 = F_palloc(m, int32(base.Ui32(v31)>>(uint(v32)%32))+int32(base.Ui32(v34)>>(uint(v32)%32)))
	mBase = m.M
	v39 = m.ExcPending
	if v39 != 0 {
		goto L1
	} else {
		goto L4
	}
L4:
	;
	v40 = *(*int32)(unsafe.Add(mBase, uint32(v29)))
	v41 = *(*int32)(unsafe.Add(mBase, uint32(v24)))
	v42 = *(*int32)(unsafe.Add(mBase, uint32(v29)+4))
	v43 = int32(268435455)
	v44 = v42 & v43
	v45 = *(*int32)(unsafe.Add(mBase, uint32(v24)+4))
	v47 = v45 & v43
	v48 = v44 + v47
	*(*int32)(unsafe.Add(mBase, uint32(v38)+4)) = v48 | int32(-2147483648)
	v52 = int32(-4)
	*(*int32)(unsafe.Add(mBase, uint32(v38))) = (v40+v41&v52)&v52 - int32(32)
	v61 = v38 + int32(4)
	if v47 == int32(0) {
		goto L5
	} else {
		goto L6
	}
L5:
	;
	v64 = *(*int32)(unsafe.Add(mBase, uint32(v29)))
	v66 = int32(base.Ui32(v64) >> (uint(int32(2)) % 32))
	if v66 != 0 {
		goto L9
	} else {
		goto L10
	}
L6:
	;
	goto L7
L7:
	;
	if v44 == int32(0) {
		goto L15
	} else {
		goto L16
	}
L8:
	;
	if v44 != 0 {
		goto L12
	} else {
		goto L13
	}
L9:
	;
	v67 = F__emscripten_memcpy_bulkmem(m, v38, v29, v66)
	mBase = m.M
	v68 = v67
	goto L11
L10:
	;
	v68 = v38
	goto L11
L11:
	;
	goto L8
L12:
	;
	v70 = v44 << (uint(int32(3)) % 32)
	v72 = *(*int32)(unsafe.Add(mBase, uint32(v70+v61)))
	v80 = (v70+v72)<<(uint(int32(2))%32) + int32(32)
	goto L14
L13:
	;
	v80 = int32(32)
	goto L14
L14:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v68))) = v80
	*(*int32)(unsafe.Add(mBase, uint32(v68)+4)) = v44 | int32(-2147483648)
	return v68
L15:
	;
	v88 = *(*int32)(unsafe.Add(mBase, uint32(v24)))
	v90 = int32(base.Ui32(v88) >> (uint(int32(2)) % 32))
	if v90 != 0 {
		goto L19
	} else {
		goto L20
	}
L16:
	;
	goto L17
L17:
	;
	v107 = int32(8)
	v108 = v29 + v107
	v109 = *(*int32)(unsafe.Add(mBase, uint32(v29)+4))
	v110 = int32(3)
	v112 = int32(2147483640)
	v114 = v108 + v109<<(uint(v110)%32)&v112
	v116 = v24 + v107
	v117 = *(*int32)(unsafe.Add(mBase, uint32(v24)+4))
	v122 = v116 + v117<<(uint(v110)%32)&v112
	v123 = int32(1)
	v126 = v38 + v107
	v131 = v126 + v48<<(uint(v110)%32)&v112
	v132 = int32(0)
	v135 = v126
	v136 = v123
	v137 = v132
	v138 = v123
	v140 = v131
	v143 = v132
	v145 = v132
	goto L22
L18:
	;
	v94 = v47 << (uint(int32(3)) % 32)
	v96 = *(*int32)(unsafe.Add(mBase, uint32(v61+v94)))
	*(*int32)(unsafe.Add(mBase, uint32(v92)+4)) = v47 | int32(-2147483648)
	*(*int32)(unsafe.Add(mBase, uint32(v92))) = (v94+v96)<<(uint(int32(2))%32) + int32(32)
	return v92
L19:
	;
	v91 = F__emscripten_memcpy_bulkmem(m, v38, v24, v90)
	mBase = m.M
	v92 = v91
	goto L21
L20:
	;
	v92 = v38
	goto L21
L21:
	;
	goto L18
L22:
	;
	v157 = int32(1)
	if v136&v157 == int32(0) {
		v368 = v157
		goto L25
	} else {
		goto L26
	}
L23:
	;
	v468 = *(*int32)(unsafe.Add(mBase, uint32(v126)))
	*(*int32)(unsafe.Add(mBase, uint32(v126))) = v468 | int32(-2147483648)
	v472 = v454 - v131
	v473 = *(*int32)(unsafe.Add(mBase, uint32(v61)))
	if v473&int32(268435455) != v462 {
		goto L100
	} else {
		goto L101
	}
L24:
	;
	v460 = base.B2i32(v451 < v44)
	v462 = v145 + int32(1)
	v464 = v135 + int32(8)
	v465 = v459 + v143
	v466 = base.B2i32(base.Ui32(v465) < base.Ui32(v47))
	if base.Ui32(v465) < base.Ui32(v47) {
		v135 = v464
		v136 = v466
		v137 = v451
		v138 = v460
		v140 = v454
		v143 = v465
		v145 = v462
		goto L22
	} else {
		goto L98
	}
L25:
	;
	v374 = v108 + v137<<(uint(int32(3))%32)
	v375 = *(*int32)(unsafe.Add(mBase, uint32(v374)))
	if v375 < int32(0) {
		goto L82
	} else {
		goto L83
	}
L26:
	;
	if v138&int32(1) == int32(0) {
		goto L27
	} else {
		goto L28
	}
L27:
	;
	v292 = v116 + v143<<(uint(int32(3))%32)
	v293 = *(*int32)(unsafe.Add(mBase, uint32(v292)))
	if v293 < int32(0) {
		goto L65
	} else {
		goto L66
	}
L28:
	;
	v166 = int32(3)
	v168 = v116 + v143<<(uint(v166)%32)
	v169 = *(*int32)(unsafe.Add(mBase, uint32(v168)))
	v170 = int32(1073741823)
	v171 = v169 & v170
	v174 = v108 + v137<<(uint(v166)%32)
	v175 = *(*int32)(unsafe.Add(mBase, uint32(v174)))
	v177 = v175 & v170
	if int32(0) <= v169 {
		goto L29
	} else {
		goto L30
	}
L29:
	;
	v182 = *(*int32)(unsafe.Add(mBase, uint32(v168-int32(4))))
	v186 = v171 - v182&int32(1073741823)
	goto L31
L30:
	;
	v186 = v171
	goto L31
L31:
	;
	if int32(0) <= v175 {
		goto L32
	} else {
		goto L33
	}
L32:
	;
	v191 = *(*int32)(unsafe.Add(mBase, uint32(v174-int32(4))))
	v195 = v177 - v191&int32(1073741823)
	goto L34
L33:
	;
	v195 = v177
	goto L34
L34:
	;
	if v186 == v195 {
		goto L35
	} else {
		goto L36
	}
L35:
	;
	v197 = int32(0)
	if v197 <= v169 {
		goto L38
	} else {
		goto L39
	}
L36:
	;
	goto L37
L37:
	;
	if v195 < v186 {
		v368 = int32(1)
		goto L25
	} else {
		goto L63
	}
L38:
	;
	v203 = *(*int32)(unsafe.Add(mBase, uint32(v168-int32(4))))
	v206 = v203 & int32(1073741823)
	goto L40
L39:
	;
	v206 = v197
	goto L40
L40:
	;
	v207 = v206 + v122
	if int32(0) <= v175 {
		goto L41
	} else {
		goto L42
	}
L41:
	;
	v212 = *(*int32)(unsafe.Add(mBase, uint32(v174-int32(4))))
	v215 = v212 & int32(1073741823)
	goto L43
L42:
	;
	v215 = v197
	goto L43
L43:
	;
	v216 = v215 + v114
	if base.Ui32(int32(4)) <= base.Ui32(v186) {
		goto L47
	} else {
		goto L48
	}
L44:
	;
	if v278 < int32(0) {
		goto L27
	} else {
		goto L62
	}
L45:
	;
	v278 = int32(0)
	goto L44
L46:
	;
	v252 = v247
	v253 = v248
	v254 = v249
	goto L56
L47:
	;
	if (v207|v216)&int32(3) != 0 {
		v247 = v207
		v248 = v216
		v249 = v186
		goto L46
	} else {
		goto L50
	}
L48:
	;
	v240 = v207
	v241 = v216
	v242 = v186
	goto L49
L49:
	;
	if v242 == int32(0) {
		goto L45
	} else {
		goto L55
	}
L50:
	;
	v224 = v207
	v225 = v216
	v226 = v186
	goto L51
L51:
	;
	v229 = *(*int32)(unsafe.Add(mBase, uint32(v224)))
	v230 = *(*int32)(unsafe.Add(mBase, uint32(v225)))
	if v229 != v230 {
		v247 = v224
		v248 = v225
		v249 = v226
		goto L46
	} else {
		goto L53
	}
L52:
	;
	v240 = v235
	v241 = v233
	v242 = v237
	goto L49
L53:
	;
	v232 = int32(4)
	v233 = v225 + v232
	v235 = v224 + v232
	v237 = v226 - v232
	if base.Ui32(int32(3)) < base.Ui32(v237) {
		v224 = v235
		v225 = v233
		v226 = v237
		goto L51
	} else {
		goto L54
	}
L54:
	;
	goto L52
L55:
	;
	v247 = v240
	v248 = v241
	v249 = v242
	goto L46
L56:
	;
	v257 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v252))))
	v258 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v253))))
	if v257 == v258 {
		goto L58
	} else {
		goto L59
	}
L57:
	;
	v278 = v257 - v258
	goto L44
L58:
	;
	v260 = int32(1)
	v265 = v254 - v260
	if v265 != 0 {
		v252 = v252 + v260
		v253 = v253 + v260
		v254 = v265
		goto L56
	} else {
		goto L61
	}
L59:
	;
	goto L60
L60:
	;
	goto L57
L61:
	;
	goto L45
L62:
	;
	v368 = v278
	goto L25
L63:
	;
	goto L27
L64:
	;
	v314 = v116 + v143<<(uint(int32(3))%32) + int32(4)
	v315 = *(*int32)(unsafe.Add(mBase, uint32(v314)))
	v317 = v315 & int32(1073741823)
	if int32(0) <= v315 {
		goto L68
	} else {
		goto L69
	}
L65:
	;
	v308 = v293 & int32(1073741823)
	v309 = v122
	goto L64
L66:
	;
	goto L67
L67:
	;
	v298 = int32(1073741823)
	v302 = *(*int32)(unsafe.Add(mBase, uint32(v292-int32(4))))
	v304 = v302 & v298
	v308 = v293&v298 - v304
	v309 = v304 + v122
	goto L64
L68:
	;
	v322 = *(*int32)(unsafe.Add(mBase, uint32(v314-int32(4))))
	v326 = v317 - v322&int32(1073741823)
	goto L70
L69:
	;
	v326 = v317
	goto L70
L70:
	;
	v327 = v326 + v308
	if v327 != 0 {
		goto L72
	} else {
		goto L73
	}
L71:
	;
	v330 = *(*int32)(unsafe.Add(mBase, uint32(v292)))
	v332 = v330 & int32(1073741823)
	if int32(0) <= v330 {
		goto L75
	} else {
		goto L76
	}
L72:
	;
	v328 = F__emscripten_memcpy_bulkmem(m, v140, v309, v327)
	mBase = m.M
	v329 = v328
	goto L74
L73:
	;
	v329 = v140
	goto L74
L74:
	;
	goto L71
L75:
	;
	v337 = *(*int32)(unsafe.Add(mBase, uint32(v292-int32(4))))
	v341 = v332 - v337&int32(1073741823)
	goto L77
L76:
	;
	v341 = v332
	goto L77
L77:
	;
	v342 = *(*int32)(unsafe.Add(mBase, uint32(v314)))
	v345 = int32(0)
	if v345 <= v342 {
		goto L78
	} else {
		goto L79
	}
L78:
	;
	v348 = v332
	goto L80
L79:
	;
	v348 = v345
	goto L80
L80:
	;
	v349 = v342&int32(1073741823) - v348
	v351 = v349 + (v341 + v329)
	v352 = v351 - v131
	v354 = int32(1073741823)
	*(*int32)(unsafe.Add(mBase, uint32(v135))) = (v352 - v349) & v354
	v357 = *(*int32)(unsafe.Add(mBase, uint32(v314)))
	*(*int32)(unsafe.Add(mBase, uint32(v135)+4)) = v357&int32(1073741824) | v352&v354
	v451 = v137
	v454 = v351
	v459 = int32(1)
	goto L24
L81:
	;
	v396 = v108 + v137<<(uint(int32(3))%32) + int32(4)
	v397 = *(*int32)(unsafe.Add(mBase, uint32(v396)))
	v399 = v397 & int32(1073741823)
	if int32(0) <= v397 {
		goto L85
	} else {
		goto L86
	}
L82:
	;
	v390 = v375 & int32(1073741823)
	v391 = v114
	goto L81
L83:
	;
	goto L84
L84:
	;
	v380 = int32(1073741823)
	v384 = *(*int32)(unsafe.Add(mBase, uint32(v374-int32(4))))
	v386 = v384 & v380
	v390 = v375&v380 - v386
	v391 = v386 + v114
	goto L81
L85:
	;
	v404 = *(*int32)(unsafe.Add(mBase, uint32(v396-int32(4))))
	v408 = v399 - v404&int32(1073741823)
	goto L87
L86:
	;
	v408 = v399
	goto L87
L87:
	;
	v409 = v408 + v390
	if v409 != 0 {
		goto L89
	} else {
		goto L90
	}
L88:
	;
	v412 = *(*int32)(unsafe.Add(mBase, uint32(v374)))
	v414 = v412 & int32(1073741823)
	if int32(0) <= v412 {
		goto L92
	} else {
		goto L93
	}
L89:
	;
	v410 = F__emscripten_memcpy_bulkmem(m, v140, v391, v409)
	mBase = m.M
	v411 = v410
	goto L91
L90:
	;
	v411 = v140
	goto L91
L91:
	;
	goto L88
L92:
	;
	v419 = *(*int32)(unsafe.Add(mBase, uint32(v374-int32(4))))
	v423 = v414 - v419&int32(1073741823)
	goto L94
L93:
	;
	v423 = v414
	goto L94
L94:
	;
	v424 = *(*int32)(unsafe.Add(mBase, uint32(v396)))
	v427 = int32(0)
	if v427 <= v424 {
		goto L95
	} else {
		goto L96
	}
L95:
	;
	v430 = v414
	goto L97
L96:
	;
	v430 = v427
	goto L97
L97:
	;
	v431 = v424&int32(1073741823) - v430
	v433 = v431 + (v423 + v411)
	v434 = v433 - v131
	v436 = int32(1073741823)
	*(*int32)(unsafe.Add(mBase, uint32(v135))) = (v434 - v431) & v436
	v439 = *(*int32)(unsafe.Add(mBase, uint32(v396)))
	*(*int32)(unsafe.Add(mBase, uint32(v135)+4)) = v439&int32(1073741824) | v434&v436
	v451 = v137 + int32(1)
	v454 = v433
	v459 = base.B2i32(v368 == int32(0))
	goto L24
L98:
	;
	if v451 < v44 {
		v135 = v464
		v136 = v466
		v137 = v451
		v138 = v460
		v140 = v454
		v143 = v465
		v145 = v462
		goto L22
	} else {
		goto L99
	}
L99:
	;
	goto L23
L100:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v61))) = v462 | int32(-2147483648)
	v484 = v126 + v462<<(uint(int32(3))%32)&int32(2147483640)
	if v484 == v131 {
		goto L104
	} else {
		goto L105
	}
L101:
	;
	goto L102
L102:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v38))) = (v462<<(uint(int32(3))%32)+v472)<<(uint(int32(2))%32) + int32(32)
	return v38
L103:
	;
	goto L102
L104:
	;
	goto L103
L105:
	;
	v488 = v484 + v472
	if base.Ui32(v131-v488) <= base.Ui32(int32(0)-v472<<(uint(int32(1))%32)) {
		goto L106
	} else {
		goto L107
	}
L106:
	;
	v495 = F___memcpy(m, v484, v131, v472)
	mBase = m.M
	goto L103
L107:
	;
	goto L108
L108:
	;
	v498 = (v484 ^ v131) & int32(3)
	if base.Ui32(v484) < base.Ui32(v131) {
		goto L111
	} else {
		goto L112
	}
L109:
	;
	if v600 == int32(0) {
		goto L104
	} else {
		goto L145
	}
L110:
	;
	if base.Ui32(v578) <= base.Ui32(int32(3)) {
		v599 = v577
		v600 = v578
		v601 = v579
		goto L109
	} else {
		goto L141
	}
L111:
	;
	if v498 != 0 {
		goto L114
	} else {
		goto L115
	}
L112:
	;
	goto L113
L113:
	;
	if v498 != 0 {
		v560 = v472
		goto L124
	} else {
		goto L125
	}
L114:
	;
	v599 = v131
	v600 = v472
	v601 = v484
	goto L109
L115:
	;
	goto L116
L116:
	;
	if v484&int32(3) == int32(0) {
		goto L117
	} else {
		goto L118
	}
L117:
	;
	v577 = v131
	v578 = v472
	v579 = v484
	goto L110
L118:
	;
	goto L119
L119:
	;
	v505 = v131
	v506 = v472
	v507 = v484
	goto L120
L120:
	;
	if v506 == int32(0) {
		goto L104
	} else {
		goto L122
	}
L121:
	;
	v577 = v514
	v578 = v516
	v579 = v518
	goto L110
L122:
	;
	v511 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v505))))
	*(*uint8)(unsafe.Add(mBase, uint32(v507))) = uint8(v511)
	v513 = int32(1)
	v514 = v505 + v513
	v516 = v506 - v513
	v518 = v507 + v513
	if v518&int32(3) != 0 {
		v505 = v514
		v506 = v516
		v507 = v518
		goto L120
	} else {
		goto L123
	}
L123:
	;
	goto L121
L124:
	;
	if v560 == int32(0) {
		goto L104
	} else {
		goto L137
	}
L125:
	;
	if v488&int32(3) != 0 {
		goto L126
	} else {
		goto L127
	}
L126:
	;
	v525 = v472
	goto L129
L127:
	;
	v540 = v472
	goto L128
L128:
	;
	if base.Ui32(v540) <= base.Ui32(int32(3)) {
		v560 = v540
		goto L124
	} else {
		goto L133
	}
L129:
	;
	if v525 == int32(0) {
		goto L104
	} else {
		goto L131
	}
L130:
	;
	v540 = v531
	goto L128
L131:
	;
	v531 = v525 - int32(1)
	v532 = v484 + v531
	v534 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v131+v531))))
	*(*uint8)(unsafe.Add(mBase, uint32(v532))) = uint8(v534)
	if v532&int32(3) != 0 {
		v525 = v531
		goto L129
	} else {
		goto L132
	}
L132:
	;
	goto L130
L133:
	;
	v547 = v540
	goto L134
L134:
	;
	v551 = v547 - int32(4)
	v554 = *(*int32)(unsafe.Add(mBase, uint32(v131+v551)))
	*(*int32)(unsafe.Add(mBase, uint32(v484+v551))) = v554
	if base.Ui32(int32(3)) < base.Ui32(v551) {
		v547 = v551
		goto L134
	} else {
		goto L136
	}
L135:
	;
	v560 = v551
	goto L124
L136:
	;
	goto L135
L137:
	;
	v567 = v560
	goto L138
L138:
	;
	v571 = v567 - int32(1)
	v574 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v131+v571))))
	*(*uint8)(unsafe.Add(mBase, uint32(v484+v571))) = uint8(v574)
	if v571 != 0 {
		v567 = v571
		goto L138
	} else {
		goto L140
	}
L139:
	;
	goto L104
L140:
	;
	goto L139
L141:
	;
	v584 = v577
	v585 = v578
	v586 = v579
	goto L142
L142:
	;
	v588 = *(*int32)(unsafe.Add(mBase, uint32(v584)))
	*(*int32)(unsafe.Add(mBase, uint32(v586))) = v588
	v590 = int32(4)
	v591 = v584 + v590
	v593 = v586 + v590
	v595 = v585 - v590
	if base.Ui32(int32(3)) < base.Ui32(v595) {
		v584 = v591
		v585 = v595
		v586 = v593
		goto L142
	} else {
		goto L144
	}
L143:
	;
	v599 = v591
	v600 = v595
	v601 = v593
	goto L109
L144:
	;
	goto L143
L145:
	;
	v606 = v599
	v607 = v600
	v608 = v601
	goto L146
L146:
	;
	v610 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v606))))
	*(*uint8)(unsafe.Add(mBase, uint32(v608))) = uint8(v610)
	v612 = int32(1)
	v617 = v607 - v612
	if v617 != 0 {
		v606 = v606 + v612
		v607 = v617
		v608 = v608 + v612
		goto L146
	} else {
		goto L148
	}
L147:
	;
	goto L104
L148:
	;
	goto L147
}
func F_hstore_eq(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v4 int32
	_ = v4
	var v5 int32
	_ = v5
	var v6 int32
	_ = v6
	var v9 int32
	_ = v9
	v4 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v5 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
	v6 = F_DirectFunctionCall2Coll(m, int32(5466), int32(0), v4, v5)
	mBase = m.M
	v9 = m.ExcPending
	if v9 != 0 {
		return int32(0)
	} else {
		return base.B2i32(v6 == int32(0))
	}
}
func F_hstore_exists_all(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v16 int32
	_ = v16
	var v18 int32
	_ = v18
	var v20 int32
	_ = v20
	var v21 int32
	_ = v21
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
	var v32 int32
	_ = v32
	var v33 int32
	_ = v33
	var v37 int32
	_ = v37
	var v38 int32
	_ = v38
	var v40 int32
	_ = v40
	var v45 int32
	_ = v45
	var v51 int32
	_ = v51
	var v64 int32
	_ = v64
	var v65 int32
	_ = v65
	var v66 int32
	_ = v66
	var v67 int32
	_ = v67
	var v74 int32
	_ = v74
	var v84 int32
	_ = v84
	var v85 int32
	_ = v85
	var v88 int32
	_ = v88
	var v89 int32
	_ = v89
	var v91 int32
	_ = v91
	var v96 int32
	_ = v96
	var v98 int32
	_ = v98
	var v99 int32
	_ = v99
	var v103 int32
	_ = v103
	var v108 int32
	_ = v108
	var v110 int32
	_ = v110
	var v111 int32
	_ = v111
	var v119 int32
	_ = v119
	var v120 int32
	_ = v120
	var v121 int32
	_ = v121
	var v124 int32
	_ = v124
	var v125 int32
	_ = v125
	var v127 int32
	_ = v127
	var v128 int32
	_ = v128
	var v130 int32
	_ = v130
	var v132 int32
	_ = v132
	var v135 int32
	_ = v135
	var v136 int32
	_ = v136
	var v137 int32
	_ = v137
	var v142 int32
	_ = v142
	var v143 int32
	_ = v143
	var v144 int32
	_ = v144
	var v147 int32
	_ = v147
	var v148 int32
	_ = v148
	var v149 int32
	_ = v149
	var v152 int32
	_ = v152
	var v153 int32
	_ = v153
	var v155 int32
	_ = v155
	var v160 int32
	_ = v160
	var v173 int32
	_ = v173
	var v174 int32
	_ = v174
	var v177 int32
	_ = v177
	var v181 int32
	_ = v181
	var v183 int32
	_ = v183
	var v185 int32
	_ = v185
	var v189 int32
	_ = v189
	var v190 int32
	_ = v190
	var v191 int32
	_ = v191
	var v195 int32
	_ = v195
	v16 = m.G0
	v18 = v16 - int32(16)
	m.G0 = v18
	v20 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v21 = F_hstoreUpgrade(m, v20)
	mBase = m.M
	v24 = m.ExcPending
	if v24 != 0 {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	return int32(0)
L2:
	;
	v26 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
	v27 = F_pg_detoast_datum(m, v26)
	mBase = m.M
	v28 = m.ExcPending
	if v28 != 0 {
		goto L1
	} else {
		goto L3
	}
L3:
	;
	v31 = F_hstoreArrayToPairs(m, v27, v18+int32(12))
	mBase = m.M
	v32 = m.ExcPending
	if v32 != 0 {
		goto L1
	} else {
		goto L4
	}
L4:
	;
	v33 = *(*int32)(unsafe.Add(mBase, uint32(v18)+12))
	if v33 <= int32(0) {
		v195 = int32(1)
		goto L5
	} else {
		goto L6
	}
L5:
	;
	m.G0 = v18 + int32(16)
	return v195
L6:
	;
	v37 = v21 + int32(8)
	v38 = *(*int32)(unsafe.Add(mBase, uint32(v21)+4))
	v40 = v38 & int32(268435455)
	v45 = int32(0)
	v51 = int32(0)
	goto L7
L7:
	;
	if v40 <= v45 {
		goto L9
	} else {
		goto L10
	}
L8:
	;
	v195 = v185
	goto L5
L9:
	;
	v195 = int32(0)
	goto L5
L10:
	;
	goto L11
L11:
	;
	v64 = v31 + v51*int32(20)
	v65 = *(*int32)(unsafe.Add(mBase, uint32(v64)+8))
	v66 = *(*int32)(unsafe.Add(mBase, uint32(v64)))
	v67 = v45
	v74 = v40
	goto L12
L12:
	;
	v84 = base.I32_div_s(v74-v67, int32(2))
	v85 = v84 + v67
	v88 = v37 + v85<<(uint(int32(3))%32)
	v89 = *(*int32)(unsafe.Add(mBase, uint32(v88)))
	v91 = v89 & int32(1073741823)
	if int32(0) <= v89 {
		goto L17
	} else {
		goto L18
	}
L13:
	;
	goto L8
L14:
	;
	v185 = int32(0)
	v189 = base.B2i32(v183 < v185)
	if v183 < v185 {
		goto L46
	} else {
		goto L47
	}
L15:
	;
	v111 = v110 + (v37 + v40<<(uint(int32(3))%32))
	if base.Ui32(int32(4)) <= base.Ui32(v65) {
		goto L28
	} else {
		goto L29
	}
L16:
	;
	if base.Ui32(v65) < base.Ui32(v103) {
		goto L22
	} else {
		goto L23
	}
L17:
	;
	v96 = *(*int32)(unsafe.Add(mBase, uint32(v88-int32(4))))
	v98 = v96 & int32(1073741823)
	v99 = v91 - v98
	if v99 != v65 {
		v103 = v99
		goto L16
	} else {
		goto L20
	}
L18:
	;
	goto L19
L19:
	;
	if v91 == v65 {
		v110 = int32(0)
		goto L15
	} else {
		goto L21
	}
L20:
	;
	v110 = v98
	goto L15
L21:
	;
	v103 = v91
	goto L16
L22:
	;
	v108 = int32(1)
	goto L24
L23:
	;
	v108 = int32(-1)
	goto L24
L24:
	;
	v183 = v108
	goto L14
L25:
	;
	if v173 != 0 {
		v183 = v173
		goto L14
	} else {
		goto L43
	}
L26:
	;
	v173 = int32(0)
	goto L25
L27:
	;
	v147 = v142
	v148 = v143
	v149 = v144
	goto L37
L28:
	;
	if (v111|v66)&int32(3) != 0 {
		v142 = v111
		v143 = v66
		v144 = v65
		goto L27
	} else {
		goto L31
	}
L29:
	;
	v135 = v111
	v136 = v66
	v137 = v65
	goto L30
L30:
	;
	if v137 == int32(0) {
		goto L26
	} else {
		goto L36
	}
L31:
	;
	v119 = v111
	v120 = v66
	v121 = v65
	goto L32
L32:
	;
	v124 = *(*int32)(unsafe.Add(mBase, uint32(v119)))
	v125 = *(*int32)(unsafe.Add(mBase, uint32(v120)))
	if v124 != v125 {
		v142 = v119
		v143 = v120
		v144 = v121
		goto L27
	} else {
		goto L34
	}
L33:
	;
	v135 = v130
	v136 = v128
	v137 = v132
	goto L30
L34:
	;
	v127 = int32(4)
	v128 = v120 + v127
	v130 = v119 + v127
	v132 = v121 - v127
	if base.Ui32(int32(3)) < base.Ui32(v132) {
		v119 = v130
		v120 = v128
		v121 = v132
		goto L32
	} else {
		goto L35
	}
L35:
	;
	goto L33
L36:
	;
	v142 = v135
	v143 = v136
	v144 = v137
	goto L27
L37:
	;
	v152 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v147))))
	v153 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v148))))
	if v152 == v153 {
		goto L39
	} else {
		goto L40
	}
L38:
	;
	v173 = v152 - v153
	goto L25
L39:
	;
	v155 = int32(1)
	v160 = v149 - v155
	if v160 != 0 {
		v147 = v147 + v155
		v148 = v148 + v155
		v149 = v160
		goto L37
	} else {
		goto L42
	}
L40:
	;
	goto L41
L41:
	;
	goto L38
L42:
	;
	goto L26
L43:
	;
	v174 = int32(0)
	if v85 < v174 {
		v195 = v174
		goto L5
	} else {
		goto L44
	}
L44:
	;
	v177 = int32(1)
	v181 = v51 + v177
	if v181 != v33 {
		v45 = v85 + v177
		v51 = v181
		goto L7
	} else {
		goto L45
	}
L45:
	;
	v195 = v177
	goto L5
L46:
	;
	v190 = v85 + int32(1)
	goto L48
L47:
	;
	v190 = v67
	goto L48
L48:
	;
	if v183 < v185 {
		goto L49
	} else {
		goto L50
	}
L49:
	;
	v191 = v74
	goto L51
L50:
	;
	v191 = v85
	goto L51
L51:
	;
	if v190 < v191 {
		v67 = v190
		v74 = v191
		goto L12
	} else {
		goto L52
	}
L52:
	;
	goto L13
}
func F_hstore_from_record(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v2 int32
	_ = v2
	var v15 int32
	_ = v15
	var v17 int32
	_ = v17
	var v19 int32
	_ = v19
	var v23 int32
	_ = v23
	var v25 int32
	_ = v25
	var v28 int32
	_ = v28
	var v29 int32
	_ = v29
	var v30 int32
	_ = v30
	var v31 int32
	_ = v31
	var v32 int32
	_ = v32
	var v33 int32
	_ = v33
	var v34 int32
	_ = v34
	var v35 int32
	_ = v35
	var v36 int32
	_ = v36
	var v37 int32
	_ = v37
	var v38 int32
	_ = v38
	var v39 int32
	_ = v39
	var v40 int32
	_ = v40
	var v41 int32
	_ = v41
	var v44 int32
	_ = v44
	var v46 int32
	_ = v46
	var v47 int32
	_ = v47
	var v52 int32
	_ = v52
	var v53 int32
	_ = v53
	var v54 int32
	_ = v54
	var v58 int32
	_ = v58
	var v59 int32
	_ = v59
	var v61 int32
	_ = v61
	var v64 int32
	_ = v64
	var v66 int32
	_ = v66
	var v78 int32
	_ = v78
	var v80 int32
	_ = v80
	var v82 int32
	_ = v82
	var v89 int32
	_ = v89
	var v92 int32
	_ = v92
	var v100 int32
	_ = v100
	var v103 int32
	_ = v103
	var v104 int32
	_ = v104
	var v105 int32
	_ = v105
	var v107 int32
	_ = v107
	var v113 int32
	_ = v113
	var v120 int32
	_ = v120
	var v121 int32
	_ = v121
	var v122 int32
	_ = v122
	var v123 int32
	_ = v123
	var v125 int32
	_ = v125
	var v127 int32
	_ = v127
	var v128 int32
	_ = v128
	var v137 int32
	_ = v137
	var v141 int32
	_ = v141
	var v150 int32
	_ = v150
	var v156 int32
	_ = v156
	var v157 int32
	_ = v157
	var v160 int32
	_ = v160
	var v163 int32
	_ = v163
	var v165 int32
	_ = v165
	var v174 int32
	_ = v174
	var v179 int32
	_ = v179
	var v183 int32
	_ = v183
	var v188 int32
	_ = v188
	var v190 int32
	_ = v190
	var v194 int32
	_ = v194
	var v200 int32
	_ = v200
	var v203 int32
	_ = v203
	var v209 int32
	_ = v209
	var v213 int32
	_ = v213
	var v215 int32
	_ = v215
	var v223 int32
	_ = v223
	var v228 int32
	_ = v228
	var v237 int32
	_ = v237
	var v238 int32
	_ = v238
	var v245 int32
	_ = v245
	var v246 int32
	_ = v246
	var v249 int32
	_ = v249
	var v250 int32
	_ = v250
	var v252 int32
	_ = v252
	var v259 int32
	_ = v259
	var v260 int32
	_ = v260
	var v261 int32
	_ = v261
	var v263 int32
	_ = v263
	var v271 int32
	_ = v271
	var v276 int32
	_ = v276
	var v280 int32
	_ = v280
	var v285 int32
	_ = v285
	var v287 int32
	_ = v287
	var v291 int32
	_ = v291
	var v297 int32
	_ = v297
	var v300 int32
	_ = v300
	var v306 int32
	_ = v306
	var v310 int32
	_ = v310
	var v312 int32
	_ = v312
	var v320 int32
	_ = v320
	var v323 int32
	_ = v323
	var v324 int32
	_ = v324
	var v325 int32
	_ = v325
	var v334 int32
	_ = v334
	var v336 int32
	_ = v336
	var v343 int32
	_ = v343
	var v354 int32
	_ = v354
	var v355 int32
	_ = v355
	var v356 int32
	_ = v356
	var v357 int32
	_ = v357
	var v358 int32
	_ = v358
	var v359 int32
	_ = v359
	var v363 int32
	_ = v363
	var v371 int32
	_ = v371
	var v374 int32
	_ = v374
	var v380 int32
	_ = v380
	var v387 int32
	_ = v387
	var v391 int32
	_ = v391
	var v394 int32
	_ = v394
	var v400 int32
	_ = v400
	var v407 int32
	_ = v407
	v2 = int32(0)
	v15 = m.G0
	v17 = v15 - int32(32)
	m.G0 = v17
	v19 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+24)))
	if v19 == int32(1) {
		goto L2
	} else {
		goto L3
	}
L1:
	;
	v37 = F_lookup_rowtype_tupdesc_domain(m, v36, v35)
	mBase = m.M
	v38 = m.ExcPending
	if v38 != 0 {
		goto L5
	} else {
		goto L8
	}
L2:
	;
	v23 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v25 = F_get_fn_expr_argtype(m, v23, int32(0))
	mBase = m.M
	v28 = m.ExcPending
	if v28 != 0 {
		goto L5
	} else {
		goto L6
	}
L3:
	;
	goto L4
L4:
	;
	v29 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v30 = F_pg_detoast_datum(m, v29)
	mBase = m.M
	v31 = m.ExcPending
	if v31 != 0 {
		goto L5
	} else {
		goto L7
	}
L5:
	;
	return int32(0)
L6:
	;
	v34 = v2
	v35 = int32(-1)
	v36 = v25
	goto L1
L7:
	;
	v32 = *(*int32)(unsafe.Add(mBase, uint32(v30)+4))
	v33 = *(*int32)(unsafe.Add(mBase, uint32(v30)+8))
	v34 = v30
	v35 = v32
	v36 = v33
	goto L1
L8:
	;
	v39 = *(*int32)(unsafe.Add(mBase, uint32(v37)))
	v40 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v41 = *(*int32)(unsafe.Add(mBase, uint32(v40)+16))
	if v41 == int32(0) {
		goto L10
	} else {
		goto L11
	}
L9:
	;
	if v36 == v59 {
		goto L15
	} else {
		goto L16
	}
L10:
	;
	v47 = *(*int32)(unsafe.Add(mBase, uint32(v40)+20))
	v52 = F_MemoryContextAlloc(m, v47, v39*int32(40)+int32(16))
	mBase = m.M
	v53 = m.ExcPending
	if v53 != 0 {
		goto L5
	} else {
		goto L13
	}
L11:
	;
	v44 = *(*int32)(unsafe.Add(mBase, uint32(v41)+12))
	if v44 != v39 {
		goto L10
	} else {
		goto L12
	}
L12:
	;
	v46 = *(*int32)(unsafe.Add(mBase, uint32(v41)))
	v58 = v41
	v59 = v46
	goto L9
L13:
	;
	v54 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	*(*int32)(unsafe.Add(mBase, uint32(v54)+16)) = v52
	*(*int64)(unsafe.Add(mBase, uint32(v52))) = int64(0)
	v58 = v52
	v59 = v2
	goto L9
L14:
	;
	v100 = int32(0)
	v103 = F_palloc(m, v39*int32(20))
	mBase = m.M
	v104 = m.ExcPending
	if v104 != 0 {
		goto L5
	} else {
		goto L29
	}
L15:
	;
	v61 = *(*int32)(unsafe.Add(mBase, uint32(v58)+4))
	if v61 == v35 {
		goto L14
	} else {
		goto L18
	}
L16:
	;
	goto L17
L17:
	;
	v64 = v39 * int32(40)
	v66 = v64 + int32(16)
	if v58&int32(3) != 0 {
		goto L20
	} else {
		goto L21
	}
L18:
	;
	goto L17
L19:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v58)+12)) = v39
	*(*int32)(unsafe.Add(mBase, uint32(v58)+4)) = v35
	*(*int32)(unsafe.Add(mBase, uint32(v58))) = v36
	goto L14
L20:
	;
	v92 = F__emscripten_memset_bulkmem(m, v58, base.I32_extend8_s(int32(0)), v66)
	mBase = m.M
	goto L28
L21:
	;
	if base.Ui32(int32(1024)) < base.Ui32(v66) {
		goto L20
	} else {
		goto L22
	}
L22:
	;
	if v66 == int32(0) {
		goto L19
	} else {
		goto L23
	}
L23:
	;
	v78 = v58 + v64 + int32(16)
	v80 = v58 + int32(4)
	if base.Ui32(v80) < base.Ui32(v78) {
		goto L24
	} else {
		goto L25
	}
L24:
	;
	v82 = v78
	goto L26
L25:
	;
	v82 = v80
	goto L26
L26:
	;
	v89 = F__emscripten_memset_bulkmem(m, v58, base.I32_extend8_s(int32(0)), (v58^int32(-1)+v82)&int32(-4)+int32(4))
	mBase = m.M
	goto L27
L27:
	;
	goto L19
L28:
	;
	goto L19
L29:
	;
	if v34 != 0 {
		goto L30
	} else {
		goto L31
	}
L30:
	;
	v105 = *(*int32)(unsafe.Add(mBase, uint32(v34)))
	*(*int32)(unsafe.Add(mBase, uint32(v17)+24)) = v34
	v107 = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v17)+20)) = v107
	*(*uint16)(unsafe.Add(mBase, uint32(v17)+16)) = uint16(v107)
	*(*int32)(unsafe.Add(mBase, uint32(v17)+12)) = int32(-1)
	v113 = int32(2)
	*(*int32)(unsafe.Add(mBase, uint32(v17)+8)) = int32(base.Ui32(v105) >> (uint(v113) % 32))
	v120 = F_palloc(m, v39<<(uint(v113)%32))
	mBase = m.M
	v121 = m.ExcPending
	if v121 != 0 {
		goto L5
	} else {
		goto L33
	}
L31:
	;
	v127 = v2
	v128 = v2
	goto L32
L32:
	;
	if int32(0) < v39 {
		goto L38
	} else {
		goto L39
	}
L33:
	;
	v122 = F_palloc(m, v39)
	mBase = m.M
	v123 = m.ExcPending
	if v123 != 0 {
		goto L5
	} else {
		goto L34
	}
L34:
	;
	F_heap_deform_tuple(m, v17+int32(8), v37, v120, v122)
	mBase = m.M
	v125 = m.ExcPending
	if v125 != 0 {
		goto L5
	} else {
		goto L35
	}
L35:
	;
	v127 = v122
	v128 = v120
	goto L32
L36:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v391 = m.ExcPending
	if v391 != 0 {
		goto L5
	} else {
		goto L105
	}
L37:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v371 = m.ExcPending
	if v371 != 0 {
		goto L5
	} else {
		goto L101
	}
L38:
	;
	v137 = int32(0)
	v141 = v100
	goto L41
L39:
	;
	v343 = v100
	goto L40
L40:
	;
	v354 = F_hstoreUniquePairs(m, v103, v343, v17+int32(28))
	mBase = m.M
	v355 = m.ExcPending
	if v355 != 0 {
		goto L5
	} else {
		goto L95
	}
L41:
	;
	v150 = *(*int32)(unsafe.Add(mBase, uint32(v37)))
	v156 = v37 + int32(20) + v150<<(uint(int32(4))%32) + v137*int32(100)
	v157 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v156)+91)))
	if v157 == int32(0) {
		goto L43
	} else {
		goto L44
	}
L42:
	;
	v343 = v334
	goto L40
L43:
	;
	v160 = *(*int32)(unsafe.Add(mBase, uint32(v156)+68))
	v163 = v103 + v141*int32(20)
	v165 = v156 + int32(4)
	*(*int32)(unsafe.Add(mBase, uint32(v163))) = v165
	if v165&int32(3) == int32(0) {
		v190 = v165
		goto L48
	} else {
		goto L49
	}
L44:
	;
	v334 = v141
	goto L45
L45:
	;
	v336 = v137 + int32(1)
	if v336 != v39 {
		v137 = v336
		v141 = v334
		goto L41
	} else {
		goto L94
	}
L46:
	;
	if base.Ui32(int32(1073741824)) <= base.Ui32(v223) {
		goto L37
	} else {
		goto L63
	}
L47:
	;
	v223 = v215 - v165
	goto L46
L48:
	;
	v194 = v190
	goto L57
L49:
	;
	v174 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v165))))
	if v174 == int32(0) {
		goto L50
	} else {
		goto L51
	}
L50:
	;
	v223 = int32(0)
	goto L46
L51:
	;
	goto L52
L52:
	;
	v179 = v165
	goto L53
L53:
	;
	v183 = v179 + int32(1)
	if v183&int32(3) == int32(0) {
		v190 = v183
		goto L48
	} else {
		goto L55
	}
L54:
	;
	v215 = v183
	goto L47
L55:
	;
	v188 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v183))))
	if v188 != 0 {
		v179 = v183
		goto L53
	} else {
		goto L56
	}
L56:
	;
	goto L54
L57:
	;
	v200 = *(*int32)(unsafe.Add(mBase, uint32(v194)))
	v203 = int32(-2139062144)
	if (int32(16843008)-v200|v200)&v203 == v203 {
		v194 = v194 + int32(4)
		goto L57
	} else {
		goto L59
	}
L58:
	;
	v209 = v194
	goto L60
L59:
	;
	goto L58
L60:
	;
	v213 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v209))))
	if v213 != 0 {
		v209 = v209 + int32(1)
		goto L60
	} else {
		goto L62
	}
L61:
	;
	v215 = v209
	goto L47
L62:
	;
	goto L61
L63:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v163)+8)) = v223
	if v127 != 0 {
		goto L66
	} else {
		goto L67
	}
L64:
	;
	v325 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v163)+17)) = uint8(v325)
	*(*uint8)(unsafe.Add(mBase, uint32(v163)+16)) = uint8(v323)
	*(*int32)(unsafe.Add(mBase, uint32(v163)+12)) = v324
	v334 = v141 + int32(1)
	goto L45
L65:
	;
	v237 = v58 + int32(16) + v137*int32(40)
	v238 = *(*int32)(unsafe.Add(mBase, uint32(v237)))
	if v160 != v238 {
		goto L70
	} else {
		goto L71
	}
L66:
	;
	v228 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v137+v127))))
	if v228 != int32(1) {
		goto L65
	} else {
		goto L69
	}
L67:
	;
	goto L68
L68:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v163)+4)) = int32(0)
	v323 = int32(1)
	v324 = int32(4)
	goto L64
L69:
	;
	goto L68
L70:
	;
	F_getTypeOutputInfo(m, v160, v237+int32(4), v17+int32(28))
	mBase = m.M
	v245 = m.ExcPending
	if v245 != 0 {
		goto L5
	} else {
		goto L73
	}
L71:
	;
	goto L72
L72:
	;
	v259 = *(*int32)(unsafe.Add(mBase, uint32(v128+v137<<(uint(int32(2))%32))))
	v260 = F_OutputFunctionCall(m, v237+int32(12), v259)
	mBase = m.M
	v261 = m.ExcPending
	if v261 != 0 {
		goto L5
	} else {
		goto L75
	}
L73:
	;
	v246 = *(*int32)(unsafe.Add(mBase, uint32(v237)+4))
	v249 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v250 = *(*int32)(unsafe.Add(mBase, uint32(v249)+20))
	F_fmgr_info_cxt(m, v246, v237+int32(12), v250)
	mBase = m.M
	v252 = m.ExcPending
	if v252 != 0 {
		goto L5
	} else {
		goto L74
	}
L74:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v237))) = v160
	goto L72
L75:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v163)+4)) = v260
	v263 = int32(0)
	if v260&int32(3) == v263 {
		v287 = v260
		goto L78
	} else {
		goto L79
	}
L76:
	;
	if base.Ui32(int32(1073741824)) <= base.Ui32(v320) {
		goto L36
	} else {
		goto L93
	}
L77:
	;
	v320 = v312 - v260
	goto L76
L78:
	;
	v291 = v287
	goto L87
L79:
	;
	v271 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v260))))
	if v271 == int32(0) {
		goto L80
	} else {
		goto L81
	}
L80:
	;
	v320 = int32(0)
	goto L76
L81:
	;
	goto L82
L82:
	;
	v276 = v260
	goto L83
L83:
	;
	v280 = v276 + int32(1)
	if v280&int32(3) == int32(0) {
		v287 = v280
		goto L78
	} else {
		goto L85
	}
L84:
	;
	v312 = v280
	goto L77
L85:
	;
	v285 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v280))))
	if v285 != 0 {
		v276 = v280
		goto L83
	} else {
		goto L86
	}
L86:
	;
	goto L84
L87:
	;
	v297 = *(*int32)(unsafe.Add(mBase, uint32(v291)))
	v300 = int32(-2139062144)
	if (int32(16843008)-v297|v297)&v300 == v300 {
		v291 = v291 + int32(4)
		goto L87
	} else {
		goto L89
	}
L88:
	;
	v306 = v291
	goto L90
L89:
	;
	goto L88
L90:
	;
	v310 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v306))))
	if v310 != 0 {
		v306 = v306 + int32(1)
		goto L90
	} else {
		goto L92
	}
L91:
	;
	v312 = v306
	goto L77
L92:
	;
	goto L91
L93:
	;
	v323 = v263
	v324 = v320
	goto L64
L94:
	;
	goto L42
L95:
	;
	v356 = *(*int32)(unsafe.Add(mBase, uint32(v17)+28))
	v357 = F_hstorePairs(m, v103, v354, v356)
	mBase = m.M
	v358 = m.ExcPending
	if v358 != 0 {
		goto L5
	} else {
		goto L96
	}
L96:
	;
	v359 = *(*int32)(unsafe.Add(mBase, uint32(v37)+12))
	if int32(0) <= v359 {
		goto L97
	} else {
		goto L98
	}
L97:
	;
	F_DecrTupleDescRefCount(m, v37)
	mBase = m.M
	v363 = m.ExcPending
	if v363 != 0 {
		goto L5
	} else {
		goto L100
	}
L98:
	;
	goto L99
L99:
	;
	m.G0 = v17 + int32(32)
	return v357
L100:
	;
	goto L99
L101:
	;
	F_errcode(m, int32(16777346))
	mBase = m.M
	v374 = m.ExcPending
	if v374 != 0 {
		goto L5
	} else {
		goto L102
	}
L102:
	;
	F_errmsg(m, int32(22170), int32(0))
	mBase = m.M
	v380 = m.ExcPending
	if v380 != 0 {
		goto L5
	} else {
		goto L103
	}
L103:
	;
	F_errfinish(m, int32(488526), int32(413), int32(278300))
	mBase = m.M
	v387 = m.ExcPending
	if v387 != 0 {
		goto L5
	} else {
		goto L104
	}
L104:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L105:
	;
	F_errcode(m, int32(16777346))
	mBase = m.M
	v394 = m.ExcPending
	if v394 != 0 {
		goto L5
	} else {
		goto L106
	}
L106:
	;
	F_errmsg(m, int32(340730), int32(0))
	mBase = m.M
	v400 = m.ExcPending
	if v400 != 0 {
		goto L5
	} else {
		goto L107
	}
L107:
	;
	F_errfinish(m, int32(488526), int32(433), int32(278318))
	mBase = m.M
	v407 = m.ExcPending
	if v407 != 0 {
		goto L5
	} else {
		goto L108
	}
L108:
	;
	base.Wasm_trap_unreachable()
	for {
	}
}
func F_hstore_hash(m *base.Module, l0 int32) int32 {
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
	var v10 int32
	_ = v10
	var v11 int32
	_ = v11
	var v15 int32
	_ = v15
	var v21 int32
	_ = v21
	var v28 int32
	_ = v28
	var v29 int32
	_ = v29
	var v30 int32
	_ = v30
	var v31 int32
	_ = v31
	var v32 int32
	_ = v32
	var v34 int32
	_ = v34
	var v35 int32
	_ = v35
	var v36 int32
	_ = v36
	var v38 int32
	_ = v38
	var v39 int32
	_ = v39
	var v41 int32
	_ = v41
	var v43 int32
	_ = v43
	var v47 int32
	_ = v47
	var v48 int32
	_ = v48
	var v49 int32
	_ = v49
	var v50 int32
	_ = v50
	var v54 int32
	_ = v54
	var v58 int32
	_ = v58
	var v62 int32
	_ = v62
	var v63 int32
	_ = v63
	var v64 int32
	_ = v64
	var v65 int32
	_ = v65
	var v69 int32
	_ = v69
	var v70 int32
	_ = v70
	var v71 int32
	_ = v71
	var v73 int32
	_ = v73
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
	var v84 int32
	_ = v84
	var v88 int32
	_ = v88
	var v89 int32
	_ = v89
	var v93 int32
	_ = v93
	var v94 int32
	_ = v94
	var v98 int32
	_ = v98
	var v99 int32
	_ = v99
	var v101 int32
	_ = v101
	var v103 int32
	_ = v103
	var v107 int32
	_ = v107
	var v108 int32
	_ = v108
	var v112 int32
	_ = v112
	var v113 int32
	_ = v113
	var v115 int32
	_ = v115
	var v116 int32
	_ = v116
	var v118 int32
	_ = v118
	var v122 int32
	_ = v122
	var v123 int32
	_ = v123
	var v127 int32
	_ = v127
	var v128 int32
	_ = v128
	var v130 int32
	_ = v130
	var v131 int32
	_ = v131
	var v132 int32
	_ = v132
	var v133 int32
	_ = v133
	var v134 int32
	_ = v134
	var v136 int32
	_ = v136
	var v137 int32
	_ = v137
	var v138 int32
	_ = v138
	var v140 int32
	_ = v140
	var v141 int32
	_ = v141
	var v143 int32
	_ = v143
	var v145 int32
	_ = v145
	var v149 int32
	_ = v149
	var v150 int32
	_ = v150
	var v151 int32
	_ = v151
	var v152 int32
	_ = v152
	var v156 int32
	_ = v156
	var v160 int32
	_ = v160
	var v164 int32
	_ = v164
	var v165 int32
	_ = v165
	var v166 int32
	_ = v166
	var v167 int32
	_ = v167
	var v171 int32
	_ = v171
	var v172 int32
	_ = v172
	var v173 int32
	_ = v173
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
	var v182 int32
	_ = v182
	var v186 int32
	_ = v186
	var v190 int32
	_ = v190
	var v191 int32
	_ = v191
	var v195 int32
	_ = v195
	var v196 int32
	_ = v196
	var v200 int32
	_ = v200
	var v201 int32
	_ = v201
	var v205 int32
	_ = v205
	var v206 int32
	_ = v206
	var v207 int32
	_ = v207
	var v211 int32
	_ = v211
	var v212 int32
	_ = v212
	var v213 int32
	_ = v213
	var v217 int32
	_ = v217
	var v218 int32
	_ = v218
	var v219 int32
	_ = v219
	var v221 int32
	_ = v221
	var v222 int32
	_ = v222
	var v223 int32
	_ = v223
	var v227 int32
	_ = v227
	var v228 int32
	_ = v228
	var v229 int32
	_ = v229
	var v230 int32
	_ = v230
	var v234 int32
	_ = v234
	var v235 int32
	_ = v235
	var v236 int32
	_ = v236
	var v237 int32
	_ = v237
	var v241 int32
	_ = v241
	var v242 int32
	_ = v242
	var v243 int32
	_ = v243
	var v244 int32
	_ = v244
	var v248 int32
	_ = v248
	var v249 int32
	_ = v249
	var v250 int32
	_ = v250
	var v253 int32
	_ = v253
	var v255 int32
	_ = v255
	var v259 int32
	_ = v259
	var v263 int32
	_ = v263
	var v267 int32
	_ = v267
	var v271 int32
	_ = v271
	var v275 int32
	_ = v275
	var v280 int32
	_ = v280
	var v283 int32
	_ = v283
	v4 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v5 = F_hstoreUpgrade(m, v4)
	mBase = m.M
	v8 = m.ExcPending
	if v8 != 0 {
		return int32(0)
	} else {
		v9 = int32(4)
		v10 = v5 + v9
		v11 = *(*int32)(unsafe.Add(mBase, uint32(v5)))
		v15 = int32(base.Ui32(v11)>>(uint(int32(2))%32)) - v9
		v21 = v15 - int32(1636608432)
		if v10&int32(3) != 0 {
			if base.Ui32(int32(11)) < base.Ui32(v15) {
				v130 = v10
				v131 = v15
				v132 = v21
				v133 = v21
				v134 = v21
				for {
					v136 = *(*int32)(unsafe.Add(mBase, uint32(v130)+4))
					v137 = v136 + v133
					v138 = *(*int32)(unsafe.Add(mBase, uint32(v130)))
					v140 = *(*int32)(unsafe.Add(mBase, uint32(v130)+8))
					v141 = v140 + v134
					v143 = int32(4)
					v145 = v138 + v132 - v141 ^ base.I32_rotl(v141, v143)
					v149 = v137 - v145 ^ base.I32_rotl(v145, int32(6))
					v150 = v141 + v137
					v151 = v145 + v150
					v152 = v149 + v151
					v156 = v150 - v149 ^ base.I32_rotl(v149, int32(8))
					v160 = v151 - v156 ^ base.I32_rotl(v156, int32(16))
					v164 = v152 - v160 ^ base.I32_rotl(v160, int32(19))
					v165 = v156 + v152
					v166 = v160 + v165
					v167 = v164 + v166
					v171 = v165 - v164 ^ base.I32_rotl(v164, v143)
					v172 = int32(12)
					v173 = v130 + v172
					v175 = v131 - v172
					if base.Ui32(int32(11)) < base.Ui32(v175) {
						v130 = v173
						v131 = v175
						v132 = v166
						v133 = v167
						v134 = v171
						continue
					} else {
						break
					}
					break
				}
				v178 = v173
				v179 = v175
				v180 = v166
				v181 = v167
				v182 = v171
			} else {
				v178 = v10
				v179 = v15
				v180 = v21
				v181 = v21
				v182 = v21
			}
			switch v179 - int32(1) {
			case 0:
				v241 = v180
				v242 = v181
				v243 = v182
				v244 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v178))))
				v248 = v241 + v244
				v249 = v242
				v250 = v243
			case 1:
				v234 = v180
				v235 = v181
				v236 = v182
				v237 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v178)+1)))
				v241 = v237<<(uint(int32(8))%32) + v234
				v242 = v235
				v243 = v236
				v244 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v178))))
				v248 = v241 + v244
				v249 = v242
				v250 = v243
			case 2:
				v227 = v180
				v228 = v181
				v229 = v182
				v230 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v178)+2)))
				v234 = v230<<(uint(int32(16))%32) + v227
				v235 = v228
				v236 = v229
				v237 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v178)+1)))
				v241 = v237<<(uint(int32(8))%32) + v234
				v242 = v235
				v243 = v236
				v244 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v178))))
				v248 = v241 + v244
				v249 = v242
				v250 = v243
			case 3:
				v221 = v181
				v222 = v182
				v223 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v178)+3)))
				v227 = v223<<(uint(int32(24))%32) + v180
				v228 = v221
				v229 = v222
				v230 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v178)+2)))
				v234 = v230<<(uint(int32(16))%32) + v227
				v235 = v228
				v236 = v229
				v237 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v178)+1)))
				v241 = v237<<(uint(int32(8))%32) + v234
				v242 = v235
				v243 = v236
				v244 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v178))))
				v248 = v241 + v244
				v249 = v242
				v250 = v243
			case 4:
				v217 = v181
				v218 = v182
				v219 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v178)+4)))
				v221 = v217 + v219
				v222 = v218
				v223 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v178)+3)))
				v227 = v223<<(uint(int32(24))%32) + v180
				v228 = v221
				v229 = v222
				v230 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v178)+2)))
				v234 = v230<<(uint(int32(16))%32) + v227
				v235 = v228
				v236 = v229
				v237 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v178)+1)))
				v241 = v237<<(uint(int32(8))%32) + v234
				v242 = v235
				v243 = v236
				v244 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v178))))
				v248 = v241 + v244
				v249 = v242
				v250 = v243
			case 5:
				v211 = v181
				v212 = v182
				v213 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v178)+5)))
				v217 = v213<<(uint(int32(8))%32) + v211
				v218 = v212
				v219 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v178)+4)))
				v221 = v217 + v219
				v222 = v218
				v223 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v178)+3)))
				v227 = v223<<(uint(int32(24))%32) + v180
				v228 = v221
				v229 = v222
				v230 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v178)+2)))
				v234 = v230<<(uint(int32(16))%32) + v227
				v235 = v228
				v236 = v229
				v237 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v178)+1)))
				v241 = v237<<(uint(int32(8))%32) + v234
				v242 = v235
				v243 = v236
				v244 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v178))))
				v248 = v241 + v244
				v249 = v242
				v250 = v243
			case 6:
				v205 = v181
				v206 = v182
				v207 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v178)+6)))
				v211 = v207<<(uint(int32(16))%32) + v205
				v212 = v206
				v213 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v178)+5)))
				v217 = v213<<(uint(int32(8))%32) + v211
				v218 = v212
				v219 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v178)+4)))
				v221 = v217 + v219
				v222 = v218
				v223 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v178)+3)))
				v227 = v223<<(uint(int32(24))%32) + v180
				v228 = v221
				v229 = v222
				v230 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v178)+2)))
				v234 = v230<<(uint(int32(16))%32) + v227
				v235 = v228
				v236 = v229
				v237 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v178)+1)))
				v241 = v237<<(uint(int32(8))%32) + v234
				v242 = v235
				v243 = v236
				v244 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v178))))
				v248 = v241 + v244
				v249 = v242
				v250 = v243
			case 7:
				v200 = v182
				v201 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v178)+7)))
				v205 = v201<<(uint(int32(24))%32) + v181
				v206 = v200
				v207 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v178)+6)))
				v211 = v207<<(uint(int32(16))%32) + v205
				v212 = v206
				v213 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v178)+5)))
				v217 = v213<<(uint(int32(8))%32) + v211
				v218 = v212
				v219 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v178)+4)))
				v221 = v217 + v219
				v222 = v218
				v223 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v178)+3)))
				v227 = v223<<(uint(int32(24))%32) + v180
				v228 = v221
				v229 = v222
				v230 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v178)+2)))
				v234 = v230<<(uint(int32(16))%32) + v227
				v235 = v228
				v236 = v229
				v237 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v178)+1)))
				v241 = v237<<(uint(int32(8))%32) + v234
				v242 = v235
				v243 = v236
				v244 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v178))))
				v248 = v241 + v244
				v249 = v242
				v250 = v243
			case 8:
				v195 = v182
				v196 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v178)+8)))
				v200 = v196<<(uint(int32(8))%32) + v195
				v201 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v178)+7)))
				v205 = v201<<(uint(int32(24))%32) + v181
				v206 = v200
				v207 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v178)+6)))
				v211 = v207<<(uint(int32(16))%32) + v205
				v212 = v206
				v213 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v178)+5)))
				v217 = v213<<(uint(int32(8))%32) + v211
				v218 = v212
				v219 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v178)+4)))
				v221 = v217 + v219
				v222 = v218
				v223 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v178)+3)))
				v227 = v223<<(uint(int32(24))%32) + v180
				v228 = v221
				v229 = v222
				v230 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v178)+2)))
				v234 = v230<<(uint(int32(16))%32) + v227
				v235 = v228
				v236 = v229
				v237 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v178)+1)))
				v241 = v237<<(uint(int32(8))%32) + v234
				v242 = v235
				v243 = v236
				v244 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v178))))
				v248 = v241 + v244
				v249 = v242
				v250 = v243
			case 9:
				v190 = v182
				v191 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v178)+9)))
				v195 = v191<<(uint(int32(16))%32) + v190
				v196 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v178)+8)))
				v200 = v196<<(uint(int32(8))%32) + v195
				v201 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v178)+7)))
				v205 = v201<<(uint(int32(24))%32) + v181
				v206 = v200
				v207 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v178)+6)))
				v211 = v207<<(uint(int32(16))%32) + v205
				v212 = v206
				v213 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v178)+5)))
				v217 = v213<<(uint(int32(8))%32) + v211
				v218 = v212
				v219 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v178)+4)))
				v221 = v217 + v219
				v222 = v218
				v223 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v178)+3)))
				v227 = v223<<(uint(int32(24))%32) + v180
				v228 = v221
				v229 = v222
				v230 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v178)+2)))
				v234 = v230<<(uint(int32(16))%32) + v227
				v235 = v228
				v236 = v229
				v237 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v178)+1)))
				v241 = v237<<(uint(int32(8))%32) + v234
				v242 = v235
				v243 = v236
				v244 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v178))))
				v248 = v241 + v244
				v249 = v242
				v250 = v243
			case 10:
				v186 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v178)+10)))
				v190 = v186<<(uint(int32(24))%32) + v182
				v191 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v178)+9)))
				v195 = v191<<(uint(int32(16))%32) + v190
				v196 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v178)+8)))
				v200 = v196<<(uint(int32(8))%32) + v195
				v201 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v178)+7)))
				v205 = v201<<(uint(int32(24))%32) + v181
				v206 = v200
				v207 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v178)+6)))
				v211 = v207<<(uint(int32(16))%32) + v205
				v212 = v206
				v213 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v178)+5)))
				v217 = v213<<(uint(int32(8))%32) + v211
				v218 = v212
				v219 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v178)+4)))
				v221 = v217 + v219
				v222 = v218
				v223 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v178)+3)))
				v227 = v223<<(uint(int32(24))%32) + v180
				v228 = v221
				v229 = v222
				v230 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v178)+2)))
				v234 = v230<<(uint(int32(16))%32) + v227
				v235 = v228
				v236 = v229
				v237 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v178)+1)))
				v241 = v237<<(uint(int32(8))%32) + v234
				v242 = v235
				v243 = v236
				v244 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v178))))
				v248 = v241 + v244
				v249 = v242
				v250 = v243
			default:
				v248 = v180
				v249 = v181
				v250 = v182
			}
		} else {
			if base.Ui32(v15) < base.Ui32(int32(12)) {
				v76 = v10
				v77 = v15
				v78 = v21
				v79 = v21
				v80 = v21
			} else {
				v28 = v10
				v29 = v15
				v30 = v21
				v31 = v21
				v32 = v21
				for {
					v34 = *(*int32)(unsafe.Add(mBase, uint32(v28)+4))
					v35 = v34 + v31
					v36 = *(*int32)(unsafe.Add(mBase, uint32(v28)))
					v38 = *(*int32)(unsafe.Add(mBase, uint32(v28)+8))
					v39 = v38 + v32
					v41 = int32(4)
					v43 = v36 + v30 - v39 ^ base.I32_rotl(v39, v41)
					v47 = v35 - v43 ^ base.I32_rotl(v43, int32(6))
					v48 = v39 + v35
					v49 = v43 + v48
					v50 = v47 + v49
					v54 = v48 - v47 ^ base.I32_rotl(v47, int32(8))
					v58 = v49 - v54 ^ base.I32_rotl(v54, int32(16))
					v62 = v50 - v58 ^ base.I32_rotl(v58, int32(19))
					v63 = v54 + v50
					v64 = v58 + v63
					v65 = v62 + v64
					v69 = v63 - v62 ^ base.I32_rotl(v62, v41)
					v70 = int32(12)
					v71 = v28 + v70
					v73 = v29 - v70
					if base.Ui32(int32(11)) < base.Ui32(v73) {
						v28 = v71
						v29 = v73
						v30 = v64
						v31 = v65
						v32 = v69
						continue
					} else {
						break
					}
					break
				}
				v76 = v71
				v77 = v73
				v78 = v64
				v79 = v65
				v80 = v69
			}
			switch v77 - int32(1) {
			case 0:
				v127 = v78
				v128 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v76))))
				v248 = v127 + v128
				v249 = v79
				v250 = v80
			case 1:
				v122 = v78
				v123 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v76)+1)))
				v127 = v123<<(uint(int32(8))%32) + v122
				v128 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v76))))
				v248 = v127 + v128
				v249 = v79
				v250 = v80
			case 2:
				v118 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v76)+2)))
				v122 = v118<<(uint(int32(16))%32) + v78
				v123 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v76)+1)))
				v127 = v123<<(uint(int32(8))%32) + v122
				v128 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v76))))
				v248 = v127 + v128
				v249 = v79
				v250 = v80
			case 3:
				v115 = v79
				v116 = *(*int32)(unsafe.Add(mBase, uint32(v76)))
				v248 = v116 + v78
				v249 = v115
				v250 = v80
			case 4:
				v112 = v79
				v113 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v76)+4)))
				v115 = v112 + v113
				v116 = *(*int32)(unsafe.Add(mBase, uint32(v76)))
				v248 = v116 + v78
				v249 = v115
				v250 = v80
			case 5:
				v107 = v79
				v108 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v76)+5)))
				v112 = v108<<(uint(int32(8))%32) + v107
				v113 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v76)+4)))
				v115 = v112 + v113
				v116 = *(*int32)(unsafe.Add(mBase, uint32(v76)))
				v248 = v116 + v78
				v249 = v115
				v250 = v80
			case 6:
				v103 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v76)+6)))
				v107 = v103<<(uint(int32(16))%32) + v79
				v108 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v76)+5)))
				v112 = v108<<(uint(int32(8))%32) + v107
				v113 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v76)+4)))
				v115 = v112 + v113
				v116 = *(*int32)(unsafe.Add(mBase, uint32(v76)))
				v248 = v116 + v78
				v249 = v115
				v250 = v80
			case 7:
				v98 = v80
				v99 = *(*int32)(unsafe.Add(mBase, uint32(v76)))
				v101 = *(*int32)(unsafe.Add(mBase, uint32(v76)+4))
				v248 = v99 + v78
				v249 = v101 + v79
				v250 = v98
			case 8:
				v93 = v80
				v94 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v76)+8)))
				v98 = v94<<(uint(int32(8))%32) + v93
				v99 = *(*int32)(unsafe.Add(mBase, uint32(v76)))
				v101 = *(*int32)(unsafe.Add(mBase, uint32(v76)+4))
				v248 = v99 + v78
				v249 = v101 + v79
				v250 = v98
			case 9:
				v88 = v80
				v89 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v76)+9)))
				v93 = v89<<(uint(int32(16))%32) + v88
				v94 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v76)+8)))
				v98 = v94<<(uint(int32(8))%32) + v93
				v99 = *(*int32)(unsafe.Add(mBase, uint32(v76)))
				v101 = *(*int32)(unsafe.Add(mBase, uint32(v76)+4))
				v248 = v99 + v78
				v249 = v101 + v79
				v250 = v98
			case 10:
				v84 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v76)+10)))
				v88 = v84<<(uint(int32(24))%32) + v80
				v89 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v76)+9)))
				v93 = v89<<(uint(int32(16))%32) + v88
				v94 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v76)+8)))
				v98 = v94<<(uint(int32(8))%32) + v93
				v99 = *(*int32)(unsafe.Add(mBase, uint32(v76)))
				v101 = *(*int32)(unsafe.Add(mBase, uint32(v76)+4))
				v248 = v99 + v78
				v249 = v101 + v79
				v250 = v98
			default:
				v248 = v78
				v249 = v79
				v250 = v80
			}
		}
		v253 = int32(14)
		v255 = v249 ^ v250 - base.I32_rotl(v249, v253)
		v259 = v255 ^ v248 - base.I32_rotl(v255, int32(11))
		v263 = v259 ^ v249 - base.I32_rotl(v259, int32(25))
		v267 = v263 ^ v255 - base.I32_rotl(v263, int32(16))
		v271 = v267 ^ v259 - base.I32_rotl(v267, int32(4))
		v275 = v271 ^ v263 - base.I32_rotl(v271, v253)
		v280 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
		if v280 != v5 {
			F_pfree(m, v5)
			mBase = m.M
			v283 = m.ExcPending
			if v283 != 0 {
				return int32(0)
			} else {
				return v275 ^ v267 - base.I32_rotl(v275, int32(24))
			}
		} else {
			return v275 ^ v267 - base.I32_rotl(v275, int32(24))
		}
	}
}
func F_hstore_lt(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v4 int32
	_ = v4
	var v5 int32
	_ = v5
	var v6 int32
	_ = v6
	var v9 int32
	_ = v9
	v4 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v5 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
	v6 = F_DirectFunctionCall2Coll(m, int32(5466), int32(0), v4, v5)
	mBase = m.M
	v9 = m.ExcPending
	if v9 != 0 {
		return int32(0)
	} else {
		return int32(base.Ui32(v6) >> (uint(int32(31)) % 32))
	}
}
func F_hstore_out(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v2 int32
	_ = v2
	var v13 int32
	_ = v13
	var v14 int32
	_ = v14
	var v17 int32
	_ = v17
	var v18 int32
	_ = v18
	var v20 int32
	_ = v20
	var v22 int32
	_ = v22
	var v25 int32
	_ = v25
	var v27 int32
	_ = v27
	var v29 int32
	_ = v29
	var v41 int32
	_ = v41
	var v42 int32
	_ = v42
	var v44 int32
	_ = v44
	var v49 int32
	_ = v49
	var v53 int32
	_ = v53
	var v56 int32
	_ = v56
	var v61 int32
	_ = v61
	var v62 int32
	_ = v62
	var v71 int32
	_ = v71
	var v75 int32
	_ = v75
	var v80 int32
	_ = v80
	var v84 int32
	_ = v84
	var v86 int32
	_ = v86
	var v88 int32
	_ = v88
	var v89 int32
	_ = v89
	var v90 int32
	_ = v90
	var v96 int32
	_ = v96
	var v102 int32
	_ = v102
	var v106 int32
	_ = v106
	var v107 int32
	_ = v107
	var v112 int32
	_ = v112
	var v116 int32
	_ = v116
	var v118 int32
	_ = v118
	var v122 int32
	_ = v122
	var v123 int32
	_ = v123
	var v125 int32
	_ = v125
	var v130 int32
	_ = v130
	var v131 int32
	_ = v131
	var v132 int32
	_ = v132
	var v135 int32
	_ = v135
	var v143 int32
	_ = v143
	var v151 int32
	_ = v151
	var v153 int32
	_ = v153
	var v156 int32
	_ = v156
	var v157 int32
	_ = v157
	var v159 int32
	_ = v159
	var v160 int32
	_ = v160
	var v162 int32
	_ = v162
	var v164 int32
	_ = v164
	var v166 int32
	_ = v166
	var v167 int32
	_ = v167
	var v180 int32
	_ = v180
	var v181 int32
	_ = v181
	var v192 int32
	_ = v192
	var v200 int32
	_ = v200
	var v202 int32
	_ = v202
	var v205 int32
	_ = v205
	var v206 int32
	_ = v206
	var v208 int32
	_ = v208
	var v215 int32
	_ = v215
	var v217 int32
	_ = v217
	var v220 int32
	_ = v220
	var v221 int32
	_ = v221
	var v223 int32
	_ = v223
	var v230 int32
	_ = v230
	var v232 int32
	_ = v232
	var v235 int32
	_ = v235
	var v236 int32
	_ = v236
	var v238 int32
	_ = v238
	var v245 int32
	_ = v245
	var v247 int32
	_ = v247
	var v250 int32
	_ = v250
	var v251 int32
	_ = v251
	var v254 int32
	_ = v254
	var v256 int32
	_ = v256
	var v260 int32
	_ = v260
	var v271 int32
	_ = v271
	var v273 int32
	_ = v273
	var v277 int32
	_ = v277
	var v278 int32
	_ = v278
	var v285 int32
	_ = v285
	var v288 int32
	_ = v288
	var v289 int32
	_ = v289
	var v294 int32
	_ = v294
	var v298 int32
	_ = v298
	var v300 int32
	_ = v300
	var v304 int32
	_ = v304
	var v305 int32
	_ = v305
	var v307 int32
	_ = v307
	var v312 int32
	_ = v312
	var v313 int32
	_ = v313
	var v314 int32
	_ = v314
	var v317 int32
	_ = v317
	var v325 int32
	_ = v325
	var v333 int32
	_ = v333
	var v335 int32
	_ = v335
	var v338 int32
	_ = v338
	var v339 int32
	_ = v339
	var v341 int32
	_ = v341
	var v342 int32
	_ = v342
	var v344 int32
	_ = v344
	var v346 int32
	_ = v346
	var v348 int32
	_ = v348
	var v349 int32
	_ = v349
	var v362 int32
	_ = v362
	var v363 int32
	_ = v363
	var v374 int32
	_ = v374
	var v382 int32
	_ = v382
	var v384 int32
	_ = v384
	var v387 int32
	_ = v387
	var v388 int32
	_ = v388
	var v390 int32
	_ = v390
	var v397 int32
	_ = v397
	var v399 int32
	_ = v399
	var v402 int32
	_ = v402
	var v403 int32
	_ = v403
	var v405 int32
	_ = v405
	var v412 int32
	_ = v412
	var v414 int32
	_ = v414
	var v417 int32
	_ = v417
	var v418 int32
	_ = v418
	var v420 int32
	_ = v420
	var v427 int32
	_ = v427
	var v429 int32
	_ = v429
	var v432 int32
	_ = v432
	var v433 int32
	_ = v433
	var v436 int32
	_ = v436
	var v438 int32
	_ = v438
	var v442 int32
	_ = v442
	var v453 int32
	_ = v453
	var v469 int32
	_ = v469
	var v471 int32
	_ = v471
	var v473 int32
	_ = v473
	var v477 int32
	_ = v477
	var v479 int32
	_ = v479
	var v483 int32
	_ = v483
	var v484 int32
	_ = v484
	v2 = int32(0)
	v13 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v14 = F_hstoreUpgrade(m, v13)
	mBase = m.M
	v17 = m.ExcPending
	if v17 != 0 {
		return int32(0)
	} else {
		v18 = *(*int32)(unsafe.Add(mBase, uint32(v14)+4))
		v20 = v18 & int32(268435455)
		if v20 != 0 {
			v22 = v14 + int32(8)
			v25 = v22 + v20<<(uint(int32(3))%32)
			v27 = int32(0)
			v29 = v2
			for {
				v41 = v22 + v27<<(uint(int32(3))%32)
				v42 = *(*int32)(unsafe.Add(mBase, uint32(v41)))
				v44 = v42 & int32(1073741823)
				if int32(0) <= v42 {
					v49 = *(*int32)(unsafe.Add(mBase, uint32(v41-int32(4))))
					v53 = v44 - v49&int32(1073741823)
				} else {
					v53 = v44
				}
				v56 = int32(4)
				v61 = v22 + v27<<(uint(int32(3))%32) + v56
				v62 = *(*int32)(unsafe.Add(mBase, uint32(v61)))
				if v62&int32(1073741824) != 0 {
					v80 = v56
				} else {
					if v62 < int32(0) {
						v75 = v62 & int32(1073741823)
					} else {
						v71 = *(*int32)(unsafe.Add(mBase, uint32(v61-int32(4))))
						v75 = v62 - v71&int32(1073741823)
					}
					v80 = v75<<(uint(int32(1))%32) + int32(2)
				}
				v84 = v80 + (v29 + v53<<(uint(int32(1))%32)) + int32(6)
				v86 = v27 + int32(1)
				if v86 != v20 {
					v27 = v86
					v29 = v84
					continue
				} else {
					break
				}
				break
			}
			v88 = F_palloc(m, v84)
			mBase = m.M
			v89 = m.ExcPending
			if v89 != 0 {
				return int32(0)
			} else {
				v90 = v88
				v96 = v2
				for {
					v102 = int32(34)
					*(*uint8)(unsafe.Add(mBase, uint32(v90))) = uint8(v102)
					v106 = v22 + v96<<(uint(int32(3))%32)
					v107 = *(*int32)(unsafe.Add(mBase, uint32(v106)))
					if v107 < int32(0) {
						v122 = v107 & int32(1073741823)
						v123 = v25
					} else {
						v112 = int32(1073741823)
						v116 = *(*int32)(unsafe.Add(mBase, uint32(v106-int32(4))))
						v118 = v116 & v112
						v122 = v107&v112 - v118
						v123 = v118 + v25
					}
					v125 = v90 + int32(1)
					if v122 <= int32(0) {
						v260 = v125
					} else {
						v130 = v122 & int32(3)
						if v130 != 0 {
							v131 = v123
							v132 = v125
							v135 = int32(0)
							for {
								v143 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v131))))
								if base.B2i32(v143 != int32(92))&base.B2i32(v143 != int32(34)) == int32(0) {
									v151 = int32(92)
									*(*uint8)(unsafe.Add(mBase, uint32(v132))) = uint8(v151)
									v153 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v131))))
									v156 = v132 + int32(1)
									v157 = v153
								} else {
									v156 = v132
									v157 = v143
								}
								*(*uint8)(unsafe.Add(mBase, uint32(v156))) = uint8(v157)
								v159 = int32(1)
								v160 = v156 + v159
								v162 = v131 + v159
								v164 = v135 + v159
								if v164 != v130 {
									v131 = v162
									v132 = v160
									v135 = v164
									continue
								} else {
									break
								}
								break
							}
							v166 = v162
							v167 = v160
						} else {
							v166 = v123
							v167 = v125
						}
						if base.Ui32(v122) < base.Ui32(int32(4)) {
							v260 = v167
						} else {
							v180 = v166
							v181 = v167
							for {
								v192 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v180))))
								if base.B2i32(v192 != int32(92))&base.B2i32(v192 != int32(34)) == int32(0) {
									v200 = int32(92)
									*(*uint8)(unsafe.Add(mBase, uint32(v181))) = uint8(v200)
									v202 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v180))))
									v205 = v181 + int32(1)
									v206 = v202
								} else {
									v205 = v181
									v206 = v192
								}
								*(*uint8)(unsafe.Add(mBase, uint32(v205))) = uint8(v206)
								v208 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v180)+1)))
								if v208 == int32(92) {
									v215 = int32(92)
									*(*uint8)(unsafe.Add(mBase, uint32(v205)+1)) = uint8(v215)
									v217 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v180)+1)))
									v220 = v217
									v221 = v205 + int32(2)
								} else {
									if v208 == int32(34) {
										v215 = int32(92)
										*(*uint8)(unsafe.Add(mBase, uint32(v205)+1)) = uint8(v215)
										v217 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v180)+1)))
										v220 = v217
										v221 = v205 + int32(2)
									} else {
										v220 = v208
										v221 = v205 + int32(1)
									}
								}
								*(*uint8)(unsafe.Add(mBase, uint32(v221))) = uint8(v220)
								v223 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v180)+2)))
								if v223 == int32(92) {
									v230 = int32(92)
									*(*uint8)(unsafe.Add(mBase, uint32(v221)+1)) = uint8(v230)
									v232 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v180)+2)))
									v235 = v232
									v236 = v221 + int32(2)
								} else {
									if v223 == int32(34) {
										v230 = int32(92)
										*(*uint8)(unsafe.Add(mBase, uint32(v221)+1)) = uint8(v230)
										v232 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v180)+2)))
										v235 = v232
										v236 = v221 + int32(2)
									} else {
										v235 = v223
										v236 = v221 + int32(1)
									}
								}
								*(*uint8)(unsafe.Add(mBase, uint32(v236))) = uint8(v235)
								v238 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v180)+3)))
								if v238 == int32(92) {
									v245 = int32(92)
									*(*uint8)(unsafe.Add(mBase, uint32(v236)+1)) = uint8(v245)
									v247 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v180)+3)))
									v250 = v247
									v251 = v236 + int32(2)
								} else {
									if v238 == int32(34) {
										v245 = int32(92)
										*(*uint8)(unsafe.Add(mBase, uint32(v236)+1)) = uint8(v245)
										v247 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v180)+3)))
										v250 = v247
										v251 = v236 + int32(2)
									} else {
										v250 = v238
										v251 = v236 + int32(1)
									}
								}
								*(*uint8)(unsafe.Add(mBase, uint32(v251))) = uint8(v250)
								v254 = v251 + int32(1)
								v256 = v180 + int32(4)
								if v256-v123 < v122 {
									v180 = v256
									v181 = v254
									continue
								} else {
									break
								}
								break
							}
							v260 = v254
						}
					}
					v271 = int32(62)
					*(*uint8)(unsafe.Add(mBase, uint32(v260)+2)) = uint8(v271)
					v273 = int32(15650)
					*(*uint16)(unsafe.Add(mBase, uint32(v260))) = uint16(v273)
					v277 = v22 + v96<<(uint(int32(3))%32)
					v278 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v277)+7)))
					if v278&int32(64) != 0 {
						*(*int32)(unsafe.Add(mBase, uint32(v260)+3)) = int32(1280070990)
						v469 = v260 + int32(7)
					} else {
						v285 = int32(34)
						*(*uint8)(unsafe.Add(mBase, uint32(v260)+3)) = uint8(v285)
						v288 = v277 + int32(4)
						v289 = *(*int32)(unsafe.Add(mBase, uint32(v288)))
						if v289 < int32(0) {
							v304 = v289 & int32(1073741823)
							v305 = v25
						} else {
							v294 = int32(1073741823)
							v298 = *(*int32)(unsafe.Add(mBase, uint32(v288-int32(4))))
							v300 = v298 & v294
							v304 = v289&v294 - v300
							v305 = v300 + v25
						}
						v307 = v260 + int32(4)
						if v304 <= int32(0) {
							v442 = v307
						} else {
							v312 = v304 & int32(3)
							if v312 != 0 {
								v313 = v305
								v314 = v307
								v317 = int32(0)
								for {
									v325 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v313))))
									if base.B2i32(v325 != int32(92))&base.B2i32(v325 != int32(34)) == int32(0) {
										v333 = int32(92)
										*(*uint8)(unsafe.Add(mBase, uint32(v314))) = uint8(v333)
										v335 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v313))))
										v338 = v314 + int32(1)
										v339 = v335
									} else {
										v338 = v314
										v339 = v325
									}
									*(*uint8)(unsafe.Add(mBase, uint32(v338))) = uint8(v339)
									v341 = int32(1)
									v342 = v338 + v341
									v344 = v313 + v341
									v346 = v317 + v341
									if v346 != v312 {
										v313 = v344
										v314 = v342
										v317 = v346
										continue
									} else {
										break
									}
									break
								}
								v348 = v344
								v349 = v342
							} else {
								v348 = v305
								v349 = v307
							}
							if base.Ui32(v304) < base.Ui32(int32(4)) {
								v442 = v349
							} else {
								v362 = v348
								v363 = v349
								for {
									v374 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v362))))
									if base.B2i32(v374 != int32(92))&base.B2i32(v374 != int32(34)) == int32(0) {
										v382 = int32(92)
										*(*uint8)(unsafe.Add(mBase, uint32(v363))) = uint8(v382)
										v384 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v362))))
										v387 = v363 + int32(1)
										v388 = v384
									} else {
										v387 = v363
										v388 = v374
									}
									*(*uint8)(unsafe.Add(mBase, uint32(v387))) = uint8(v388)
									v390 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v362)+1)))
									if v390 == int32(92) {
										v397 = int32(92)
										*(*uint8)(unsafe.Add(mBase, uint32(v387)+1)) = uint8(v397)
										v399 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v362)+1)))
										v402 = v399
										v403 = v387 + int32(2)
									} else {
										if v390 == int32(34) {
											v397 = int32(92)
											*(*uint8)(unsafe.Add(mBase, uint32(v387)+1)) = uint8(v397)
											v399 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v362)+1)))
											v402 = v399
											v403 = v387 + int32(2)
										} else {
											v402 = v390
											v403 = v387 + int32(1)
										}
									}
									*(*uint8)(unsafe.Add(mBase, uint32(v403))) = uint8(v402)
									v405 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v362)+2)))
									if v405 == int32(92) {
										v412 = int32(92)
										*(*uint8)(unsafe.Add(mBase, uint32(v403)+1)) = uint8(v412)
										v414 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v362)+2)))
										v417 = v414
										v418 = v403 + int32(2)
									} else {
										if v405 == int32(34) {
											v412 = int32(92)
											*(*uint8)(unsafe.Add(mBase, uint32(v403)+1)) = uint8(v412)
											v414 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v362)+2)))
											v417 = v414
											v418 = v403 + int32(2)
										} else {
											v417 = v405
											v418 = v403 + int32(1)
										}
									}
									*(*uint8)(unsafe.Add(mBase, uint32(v418))) = uint8(v417)
									v420 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v362)+3)))
									if v420 == int32(92) {
										v427 = int32(92)
										*(*uint8)(unsafe.Add(mBase, uint32(v418)+1)) = uint8(v427)
										v429 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v362)+3)))
										v432 = v429
										v433 = v418 + int32(2)
									} else {
										if v420 == int32(34) {
											v427 = int32(92)
											*(*uint8)(unsafe.Add(mBase, uint32(v418)+1)) = uint8(v427)
											v429 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v362)+3)))
											v432 = v429
											v433 = v418 + int32(2)
										} else {
											v432 = v420
											v433 = v418 + int32(1)
										}
									}
									*(*uint8)(unsafe.Add(mBase, uint32(v433))) = uint8(v432)
									v436 = v433 + int32(1)
									v438 = v362 + int32(4)
									if v438-v305 < v304 {
										v362 = v438
										v363 = v436
										continue
									} else {
										break
									}
									break
								}
								v442 = v436
							}
						}
						v453 = int32(34)
						*(*uint8)(unsafe.Add(mBase, uint32(v442))) = uint8(v453)
						v469 = v442 + int32(1)
					}
					v471 = v96 + int32(1)
					if v20 != v471 {
						v473 = int32(8236)
						*(*uint16)(unsafe.Add(mBase, uint32(v469))) = uint16(v473)
						v477 = v469 + int32(2)
					} else {
						v477 = v469
					}
					if v471 != v20 {
						v90 = v477
						v96 = v471
						continue
					} else {
						break
					}
					break
				}
				v479 = int32(0)
				*(*uint8)(unsafe.Add(mBase, uint32(v477))) = uint8(v479)
				return v88
			}
		} else {
			v483 = F_pstrdup(m, int32(733277))
			mBase = m.M
			v484 = m.ExcPending
			if v484 != 0 {
				return int32(0)
			} else {
				return v483
			}
		}
	}
}
func F_hstore_recv(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v9 int32
	_ = v9
	var v11 int32
	_ = v11
	var v13 int32
	_ = v13
	var v15 int32
	_ = v15
	var v18 int32
	_ = v18
	var v22 int32
	_ = v22
	var v23 int32
	_ = v23
	var v30 int32
	_ = v30
	var v31 int32
	_ = v31
	var v37 int32
	_ = v37
	var v41 int32
	_ = v41
	var v42 int32
	_ = v42
	var v47 int32
	_ = v47
	var v50 int32
	_ = v50
	var v51 int32
	_ = v51
	var v53 int32
	_ = v53
	var v56 int32
	_ = v56
	var v60 int32
	_ = v60
	var v61 int32
	_ = v61
	var v64 int32
	_ = v64
	var v70 int32
	_ = v70
	var v71 int32
	_ = v71
	var v74 int32
	_ = v74
	var v77 int32
	_ = v77
	var v78 int32
	_ = v78
	var v82 int32
	_ = v82
	var v86 int32
	_ = v86
	var v87 int32
	_ = v87
	var v88 int32
	_ = v88
	var v89 int32
	_ = v89
	var v90 int32
	_ = v90
	var v91 int32
	_ = v91
	var v106 int32
	_ = v106
	var v109 int32
	_ = v109
	var v117 int32
	_ = v117
	var v124 int32
	_ = v124
	var v128 int32
	_ = v128
	var v131 int32
	_ = v131
	var v137 int32
	_ = v137
	var v144 int32
	_ = v144
	var v148 int32
	_ = v148
	var v151 int32
	_ = v151
	var v157 int32
	_ = v157
	var v164 int32
	_ = v164
	var v168 int32
	_ = v168
	var v171 int32
	_ = v171
	var v177 int32
	_ = v177
	var v184 int32
	_ = v184
	v9 = m.G0
	v11 = v9 - int32(16)
	m.G0 = v11
	v13 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v15 = F_pq_getmsgint(m, v13, int32(4))
	mBase = m.M
	v18 = m.ExcPending
	if v18 != 0 {
		goto L6
	} else {
		goto L7
	}
L1:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v168 = m.ExcPending
	if v168 != 0 {
		goto L6
	} else {
		goto L42
	}
L2:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v148 = m.ExcPending
	if v148 != 0 {
		goto L6
	} else {
		goto L38
	}
L3:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v128 = m.ExcPending
	if v128 != 0 {
		goto L6
	} else {
		goto L34
	}
L4:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v106 = m.ExcPending
	if v106 != 0 {
		goto L6
	} else {
		goto L30
	}
L5:
	;
	m.G0 = v11 + int32(16)
	return v91
L6:
	;
	return int32(0)
L7:
	;
	if v15 == int32(0) {
		goto L8
	} else {
		goto L9
	}
L8:
	;
	v22 = F_palloc(m, int32(8))
	mBase = m.M
	v23 = m.ExcPending
	if v23 != 0 {
		goto L6
	} else {
		goto L11
	}
L9:
	;
	goto L10
L10:
	;
	if base.Ui32(int32(53687092)) <= base.Ui32(v15) {
		goto L4
	} else {
		goto L12
	}
L11:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v22))) = int64(-9223372036854775776)
	v91 = v22
	goto L5
L12:
	;
	v30 = F_palloc(m, v15*int32(20))
	mBase = m.M
	v31 = m.ExcPending
	if v31 != 0 {
		goto L6
	} else {
		goto L13
	}
L13:
	;
	v37 = int32(0)
	goto L14
L14:
	;
	v41 = F_pq_getmsgint(m, v13, int32(4))
	mBase = m.M
	v42 = m.ExcPending
	if v42 != 0 {
		goto L6
	} else {
		goto L16
	}
L15:
	;
	v86 = F_hstoreUniquePairs(m, v30, v15, v11+int32(12))
	mBase = m.M
	v87 = m.ExcPending
	if v87 != 0 {
		goto L6
	} else {
		goto L28
	}
L16:
	;
	if v41 < int32(0) {
		goto L3
	} else {
		goto L17
	}
L17:
	;
	v47 = v30 + v37*int32(20)
	v50 = F_pq_getmsgtext(m, v13, v41, v11+int32(12))
	mBase = m.M
	v51 = m.ExcPending
	if v51 != 0 {
		goto L6
	} else {
		goto L18
	}
L18:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v47))) = v50
	v53 = *(*int32)(unsafe.Add(mBase, uint32(v11)+12))
	if base.Ui32(int32(1073741824)) <= base.Ui32(v53) {
		goto L2
	} else {
		goto L19
	}
L19:
	;
	v56 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(v47)+17)) = uint8(v56)
	*(*int32)(unsafe.Add(mBase, uint32(v47)+8)) = v53
	v60 = F_pq_getmsgint(m, v13, int32(4))
	mBase = m.M
	v61 = m.ExcPending
	if v61 != 0 {
		goto L6
	} else {
		goto L21
	}
L20:
	;
	*(*uint8)(unsafe.Add(mBase, uint32(v47)+16)) = uint8(v78)
	*(*int32)(unsafe.Add(mBase, uint32(v47)+12)) = v77
	v82 = v37 + int32(1)
	if v82 != v15 {
		v37 = v82
		goto L14
	} else {
		goto L27
	}
L21:
	;
	if v60 < int32(0) {
		goto L22
	} else {
		goto L23
	}
L22:
	;
	v64 = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v47)+4)) = v64
	v77 = v64
	v78 = int32(1)
	goto L20
L23:
	;
	goto L24
L24:
	;
	v70 = F_pq_getmsgtext(m, v13, v60, v11+int32(12))
	mBase = m.M
	v71 = m.ExcPending
	if v71 != 0 {
		goto L6
	} else {
		goto L25
	}
L25:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v47)+4)) = v70
	v74 = *(*int32)(unsafe.Add(mBase, uint32(v11)+12))
	if base.Ui32(int32(1073741824)) <= base.Ui32(v74) {
		goto L1
	} else {
		goto L26
	}
L26:
	;
	v77 = v74
	v78 = int32(0)
	goto L20
L27:
	;
	goto L15
L28:
	;
	v88 = *(*int32)(unsafe.Add(mBase, uint32(v11)+12))
	v89 = F_hstorePairs(m, v30, v86, v88)
	mBase = m.M
	v90 = m.ExcPending
	if v90 != 0 {
		goto L6
	} else {
		goto L29
	}
L29:
	;
	v91 = v89
	goto L5
L30:
	;
	F_errcode(m, int32(261))
	mBase = m.M
	v109 = m.ExcPending
	if v109 != 0 {
		goto L6
	} else {
		goto L31
	}
L31:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v11)+4)) = int32(53687091)
	*(*int32)(unsafe.Add(mBase, uint32(v11))) = v15
	F_errmsg(m, int32(655860), v11)
	mBase = m.M
	v117 = m.ExcPending
	if v117 != 0 {
		goto L6
	} else {
		goto L32
	}
L32:
	;
	F_errfinish(m, int32(488526), int32(525), int32(36182))
	mBase = m.M
	v124 = m.ExcPending
	if v124 != 0 {
		goto L6
	} else {
		goto L33
	}
L33:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L34:
	;
	F_errcode(m, int32(67108994))
	mBase = m.M
	v131 = m.ExcPending
	if v131 != 0 {
		goto L6
	} else {
		goto L35
	}
L35:
	;
	F_errmsg(m, int32(22201), int32(0))
	mBase = m.M
	v137 = m.ExcPending
	if v137 != 0 {
		goto L6
	} else {
		goto L36
	}
L36:
	;
	F_errfinish(m, int32(488526), int32(536), int32(36182))
	mBase = m.M
	v144 = m.ExcPending
	if v144 != 0 {
		goto L6
	} else {
		goto L37
	}
L37:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L38:
	;
	F_errcode(m, int32(16777346))
	mBase = m.M
	v151 = m.ExcPending
	if v151 != 0 {
		goto L6
	} else {
		goto L39
	}
L39:
	;
	F_errmsg(m, int32(22170), int32(0))
	mBase = m.M
	v157 = m.ExcPending
	if v157 != 0 {
		goto L6
	} else {
		goto L40
	}
L40:
	;
	F_errfinish(m, int32(488526), int32(413), int32(278300))
	mBase = m.M
	v164 = m.ExcPending
	if v164 != 0 {
		goto L6
	} else {
		goto L41
	}
L41:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L42:
	;
	F_errcode(m, int32(16777346))
	mBase = m.M
	v171 = m.ExcPending
	if v171 != 0 {
		goto L6
	} else {
		goto L43
	}
L43:
	;
	F_errmsg(m, int32(340730), int32(0))
	mBase = m.M
	v177 = m.ExcPending
	if v177 != 0 {
		goto L6
	} else {
		goto L44
	}
L44:
	;
	F_errfinish(m, int32(488526), int32(433), int32(278318))
	mBase = m.M
	v184 = m.ExcPending
	if v184 != 0 {
		goto L6
	} else {
		goto L45
	}
L45:
	;
	base.Wasm_trap_unreachable()
	for {
	}
}
func F_hstore_slice_to_array(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v20 int32
	_ = v20
	var v22 int32
	_ = v22
	var v24 int32
	_ = v24
	var v25 int32
	_ = v25
	var v28 int32
	_ = v28
	var v29 int32
	_ = v29
	var v30 int32
	_ = v30
	var v31 int32
	_ = v31
	var v32 int32
	_ = v32
	var v41 int32
	_ = v41
	var v42 int32
	_ = v42
	var v46 int32
	_ = v46
	var v47 int32
	_ = v47
	var v50 int32
	_ = v50
	var v51 int32
	_ = v51
	var v52 int32
	_ = v52
	var v53 int32
	_ = v53
	var v54 int32
	_ = v54
	var v55 int32
	_ = v55
	var v59 int32
	_ = v59
	var v64 int32
	_ = v64
	var v68 int32
	_ = v68
	var v77 int32
	_ = v77
	var v84 int32
	_ = v84
	var v86 int32
	_ = v86
	var v87 int32
	_ = v87
	var v89 int32
	_ = v89
	var v95 int32
	_ = v95
	var v96 int32
	_ = v96
	var v97 int32
	_ = v97
	var v99 int32
	_ = v99
	var v100 int32
	_ = v100
	var v101 int32
	_ = v101
	var v102 int32
	_ = v102
	var v106 int32
	_ = v106
	var v108 int32
	_ = v108
	var v112 int32
	_ = v112
	var v129 int32
	_ = v129
	var v130 int32
	_ = v130
	var v133 int32
	_ = v133
	var v134 int32
	_ = v134
	var v136 int32
	_ = v136
	var v141 int32
	_ = v141
	var v143 int32
	_ = v143
	var v144 int32
	_ = v144
	var v148 int32
	_ = v148
	var v153 int32
	_ = v153
	var v155 int32
	_ = v155
	var v156 int32
	_ = v156
	var v164 int32
	_ = v164
	var v165 int32
	_ = v165
	var v166 int32
	_ = v166
	var v169 int32
	_ = v169
	var v170 int32
	_ = v170
	var v172 int32
	_ = v172
	var v173 int32
	_ = v173
	var v175 int32
	_ = v175
	var v177 int32
	_ = v177
	var v180 int32
	_ = v180
	var v181 int32
	_ = v181
	var v182 int32
	_ = v182
	var v187 int32
	_ = v187
	var v188 int32
	_ = v188
	var v189 int32
	_ = v189
	var v192 int32
	_ = v192
	var v193 int32
	_ = v193
	var v194 int32
	_ = v194
	var v197 int32
	_ = v197
	var v198 int32
	_ = v198
	var v200 int32
	_ = v200
	var v205 int32
	_ = v205
	var v218 int32
	_ = v218
	var v222 int32
	_ = v222
	var v223 int32
	_ = v223
	var v230 int32
	_ = v230
	var v235 int32
	_ = v235
	var v236 int32
	_ = v236
	var v237 int32
	_ = v237
	var v259 int32
	_ = v259
	var v268 int32
	_ = v268
	var v270 int32
	_ = v270
	var v273 int32
	_ = v273
	var v275 int32
	_ = v275
	var v277 int32
	_ = v277
	var v278 int32
	_ = v278
	var v281 int32
	_ = v281
	var v283 int32
	_ = v283
	var v296 int32
	_ = v296
	var v304 int32
	_ = v304
	var v325 int32
	_ = v325
	var v327 int32
	_ = v327
	var v335 int32
	_ = v335
	var v336 int32
	_ = v336
	var v356 int32
	_ = v356
	v20 = m.G0
	v22 = v20 - int32(16)
	m.G0 = v22
	v24 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v25 = F_hstoreUpgrade(m, v24)
	mBase = m.M
	v28 = m.ExcPending
	if v28 != 0 {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	return int32(0)
L2:
	;
	v29 = *(*int32)(unsafe.Add(mBase, uint32(v25)+4))
	v30 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
	v31 = F_pg_detoast_datum(m, v30)
	mBase = m.M
	v32 = m.ExcPending
	if v32 != 0 {
		goto L1
	} else {
		goto L3
	}
L3:
	;
	F_deconstruct_array_builtin(m, v31, int32(25), v22+int32(12), v22+int32(8), v22+int32(4))
	mBase = m.M
	v41 = m.ExcPending
	if v41 != 0 {
		goto L1
	} else {
		goto L4
	}
L4:
	;
	v42 = *(*int32)(unsafe.Add(mBase, uint32(v22)+4))
	if v42 == int32(0) {
		goto L6
	} else {
		goto L7
	}
L5:
	;
	m.G0 = v22 + int32(16)
	return v356
L6:
	;
	v46 = F_construct_empty_array(m, int32(25))
	mBase = m.M
	v47 = m.ExcPending
	if v47 != 0 {
		goto L1
	} else {
		goto L9
	}
L7:
	;
	goto L8
L8:
	;
	v50 = F_palloc(m, v42<<(uint(int32(2))%32))
	mBase = m.M
	v51 = m.ExcPending
	if v51 != 0 {
		goto L1
	} else {
		goto L10
	}
L9:
	;
	v356 = v46
	goto L5
L10:
	;
	v52 = *(*int32)(unsafe.Add(mBase, uint32(v22)+4))
	v53 = F_palloc(m, v52)
	mBase = m.M
	v54 = m.ExcPending
	if v54 != 0 {
		goto L1
	} else {
		goto L11
	}
L11:
	;
	v55 = *(*int32)(unsafe.Add(mBase, uint32(v22)+4))
	if int32(0) < v55 {
		goto L12
	} else {
		goto L13
	}
L12:
	;
	v59 = v25 + int32(8)
	v64 = v59 + v29<<(uint(int32(3))%32)&int32(2147483640)
	v68 = int32(0)
	v77 = v55
	goto L15
L13:
	;
	goto L14
L14:
	;
	v325 = *(*int32)(unsafe.Add(mBase, uint32(v31)+4))
	v327 = v31 + int32(16)
	v335 = F_construct_md_array(m, v50, v53, v325, v327, v327+v325<<(uint(int32(2))%32), int32(25), int32(-1), int32(0), int32(105))
	mBase = m.M
	v336 = m.ExcPending
	if v336 != 0 {
		goto L1
	} else {
		goto L67
	}
L15:
	;
	v84 = *(*int32)(unsafe.Add(mBase, uint32(v22)+8))
	v86 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v84+v68))))
	if v86 != 0 {
		goto L20
	} else {
		goto L21
	}
L16:
	;
	goto L14
L17:
	;
	v304 = v68 + int32(1)
	if v304 < v296 {
		v68 = v304
		v77 = v296
		goto L15
	} else {
		goto L66
	}
L18:
	;
	v277 = F_cstring_to_text_with_len(m, v275, v273)
	mBase = m.M
	v278 = m.ExcPending
	if v278 != 0 {
		goto L1
	} else {
		goto L65
	}
L19:
	;
	v268 = *(*int32)(unsafe.Add(mBase, uint32(v222-int32(4))))
	v270 = v268 & int32(1073741823)
	v273 = v223 - v270
	v275 = v270 + v64
	goto L18
L20:
	;
	v259 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(v68+v53))) = uint8(v259)
	*(*int32)(unsafe.Add(mBase, uint32(v50+v68<<(uint(int32(2))%32)))) = int32(0)
	v296 = v77
	goto L17
L21:
	;
	v87 = *(*int32)(unsafe.Add(mBase, uint32(v25)+4))
	v89 = v87 & int32(268435455)
	if v89 == int32(0) {
		goto L20
	} else {
		goto L22
	}
L22:
	;
	v95 = int32(2)
	v96 = v68 << (uint(v95) % 32)
	v97 = *(*int32)(unsafe.Add(mBase, uint32(v22)+12))
	v99 = *(*int32)(unsafe.Add(mBase, uint32(v96+v97)))
	v100 = int32(4)
	v101 = v99 + v100
	v102 = *(*int32)(unsafe.Add(mBase, uint32(v99)))
	v106 = int32(base.Ui32(v102)>>(uint(v95)%32)) - v100
	v108 = int32(0)
	v112 = v89
	goto L23
L23:
	;
	v129 = base.I32_div_s(v112-v108, int32(2))
	v130 = v129 + v108
	v133 = v59 + v130<<(uint(int32(3))%32)
	v134 = *(*int32)(unsafe.Add(mBase, uint32(v133)))
	v136 = v134 & int32(1073741823)
	if int32(0) <= v134 {
		goto L28
	} else {
		goto L29
	}
L24:
	;
	goto L20
L25:
	;
	v235 = base.B2i32(v230 < int32(0))
	if v230 < int32(0) {
		goto L58
	} else {
		goto L59
	}
L26:
	;
	v156 = v155 + (v59 + v89<<(uint(int32(3))%32))
	if base.Ui32(int32(4)) <= base.Ui32(v106) {
		goto L39
	} else {
		goto L40
	}
L27:
	;
	if base.Ui32(v106) < base.Ui32(v148) {
		goto L33
	} else {
		goto L34
	}
L28:
	;
	v141 = *(*int32)(unsafe.Add(mBase, uint32(v133-int32(4))))
	v143 = v141 & int32(1073741823)
	v144 = v136 - v143
	if v144 != v106 {
		v148 = v144
		goto L27
	} else {
		goto L31
	}
L29:
	;
	goto L30
L30:
	;
	if v136 == v106 {
		v155 = int32(0)
		goto L26
	} else {
		goto L32
	}
L31:
	;
	v155 = v143
	goto L26
L32:
	;
	v148 = v136
	goto L27
L33:
	;
	v153 = int32(1)
	goto L35
L34:
	;
	v153 = int32(-1)
	goto L35
L35:
	;
	v230 = v153
	goto L25
L36:
	;
	if v218 != 0 {
		v230 = v218
		goto L25
	} else {
		goto L54
	}
L37:
	;
	v218 = int32(0)
	goto L36
L38:
	;
	v192 = v187
	v193 = v188
	v194 = v189
	goto L48
L39:
	;
	if (v156|v101)&int32(3) != 0 {
		v187 = v156
		v188 = v101
		v189 = v106
		goto L38
	} else {
		goto L42
	}
L40:
	;
	v180 = v156
	v181 = v101
	v182 = v106
	goto L41
L41:
	;
	if v182 == int32(0) {
		goto L37
	} else {
		goto L47
	}
L42:
	;
	v164 = v156
	v165 = v101
	v166 = v106
	goto L43
L43:
	;
	v169 = *(*int32)(unsafe.Add(mBase, uint32(v164)))
	v170 = *(*int32)(unsafe.Add(mBase, uint32(v165)))
	if v169 != v170 {
		v187 = v164
		v188 = v165
		v189 = v166
		goto L38
	} else {
		goto L45
	}
L44:
	;
	v180 = v175
	v181 = v173
	v182 = v177
	goto L41
L45:
	;
	v172 = int32(4)
	v173 = v165 + v172
	v175 = v164 + v172
	v177 = v166 - v172
	if base.Ui32(int32(3)) < base.Ui32(v177) {
		v164 = v175
		v165 = v173
		v166 = v177
		goto L43
	} else {
		goto L46
	}
L46:
	;
	goto L44
L47:
	;
	v187 = v180
	v188 = v181
	v189 = v182
	goto L38
L48:
	;
	v197 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v192))))
	v198 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v193))))
	if v197 == v198 {
		goto L50
	} else {
		goto L51
	}
L49:
	;
	v218 = v197 - v198
	goto L36
L50:
	;
	v200 = int32(1)
	v205 = v194 - v200
	if v205 != 0 {
		v192 = v192 + v200
		v193 = v193 + v200
		v194 = v205
		goto L48
	} else {
		goto L53
	}
L51:
	;
	goto L52
L52:
	;
	goto L49
L53:
	;
	goto L37
L54:
	;
	if v130 < int32(0) {
		goto L20
	} else {
		goto L55
	}
L55:
	;
	v222 = v133 + int32(4)
	v223 = *(*int32)(unsafe.Add(mBase, uint32(v222)))
	if v223&int32(1073741824) != 0 {
		goto L20
	} else {
		goto L56
	}
L56:
	;
	if int32(0) <= v223 {
		goto L19
	} else {
		goto L57
	}
L57:
	;
	v273 = v223 & int32(1073741823)
	v275 = v64
	goto L18
L58:
	;
	v236 = v130 + int32(1)
	goto L60
L59:
	;
	v236 = v108
	goto L60
L60:
	;
	if v230 < int32(0) {
		goto L61
	} else {
		goto L62
	}
L61:
	;
	v237 = v112
	goto L63
L62:
	;
	v237 = v130
	goto L63
L63:
	;
	if v236 < v237 {
		v108 = v236
		v112 = v237
		goto L23
	} else {
		goto L64
	}
L64:
	;
	goto L24
L65:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v50+v96))) = v277
	v281 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v68+v53))) = uint8(v281)
	v283 = *(*int32)(unsafe.Add(mBase, uint32(v22)+4))
	v296 = v283
	goto L17
L66:
	;
	goto L16
L67:
	;
	v356 = v335
	goto L5
}
func F_hstore_slice_to_hstore(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v2 int32
	_ = v2
	var v20 int32
	_ = v20
	var v22 int32
	_ = v22
	var v24 int32
	_ = v24
	var v25 int32
	_ = v25
	var v28 int32
	_ = v28
	var v29 int32
	_ = v29
	var v30 int32
	_ = v30
	var v31 int32
	_ = v31
	var v32 int32
	_ = v32
	var v35 int32
	_ = v35
	var v36 int32
	_ = v36
	var v37 int32
	_ = v37
	var v40 int32
	_ = v40
	var v43 int32
	_ = v43
	var v44 int32
	_ = v44
	var v47 int32
	_ = v47
	var v48 int32
	_ = v48
	var v52 int32
	_ = v52
	var v53 int32
	_ = v53
	var v57 int32
	_ = v57
	var v58 int32
	_ = v58
	var v60 int32
	_ = v60
	var v65 int32
	_ = v65
	var v71 int32
	_ = v71
	var v72 int32
	_ = v72
	var v75 int32
	_ = v75
	var v87 int32
	_ = v87
	var v88 int32
	_ = v88
	var v89 int32
	_ = v89
	var v90 int32
	_ = v90
	var v93 int32
	_ = v93
	var v111 int32
	_ = v111
	var v112 int32
	_ = v112
	var v115 int32
	_ = v115
	var v116 int32
	_ = v116
	var v118 int32
	_ = v118
	var v123 int32
	_ = v123
	var v125 int32
	_ = v125
	var v126 int32
	_ = v126
	var v130 int32
	_ = v130
	var v135 int32
	_ = v135
	var v137 int32
	_ = v137
	var v138 int32
	_ = v138
	var v146 int32
	_ = v146
	var v147 int32
	_ = v147
	var v148 int32
	_ = v148
	var v151 int32
	_ = v151
	var v152 int32
	_ = v152
	var v154 int32
	_ = v154
	var v155 int32
	_ = v155
	var v157 int32
	_ = v157
	var v159 int32
	_ = v159
	var v162 int32
	_ = v162
	var v163 int32
	_ = v163
	var v164 int32
	_ = v164
	var v169 int32
	_ = v169
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
	var v179 int32
	_ = v179
	var v180 int32
	_ = v180
	var v182 int32
	_ = v182
	var v187 int32
	_ = v187
	var v200 int32
	_ = v200
	var v202 int32
	_ = v202
	var v207 int32
	_ = v207
	var v212 int32
	_ = v212
	var v213 int32
	_ = v213
	var v218 int32
	_ = v218
	var v223 int32
	_ = v223
	var v224 int32
	_ = v224
	var v225 int32
	_ = v225
	var v227 int32
	_ = v227
	var v231 int32
	_ = v231
	var v233 int32
	_ = v233
	var v236 int32
	_ = v236
	var v238 int32
	_ = v238
	var v239 int32
	_ = v239
	var v245 int32
	_ = v245
	var v246 int32
	_ = v246
	var v251 int32
	_ = v251
	var v257 int32
	_ = v257
	var v258 int32
	_ = v258
	var v271 int32
	_ = v271
	var v279 int32
	_ = v279
	var v280 int32
	_ = v280
	var v292 int32
	_ = v292
	var v293 int32
	_ = v293
	var v313 int32
	_ = v313
	v2 = int32(0)
	v20 = m.G0
	v22 = v20 - int32(16)
	m.G0 = v22
	v24 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v25 = F_hstoreUpgrade(m, v24)
	mBase = m.M
	v28 = m.ExcPending
	if v28 != 0 {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	return int32(0)
L2:
	;
	v29 = *(*int32)(unsafe.Add(mBase, uint32(v25)+4))
	v30 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
	v31 = F_pg_detoast_datum(m, v30)
	mBase = m.M
	v32 = m.ExcPending
	if v32 != 0 {
		goto L1
	} else {
		goto L3
	}
L3:
	;
	v35 = F_hstoreArrayToPairs(m, v31, v22+int32(12))
	mBase = m.M
	v36 = m.ExcPending
	if v36 != 0 {
		goto L1
	} else {
		goto L4
	}
L4:
	;
	v37 = *(*int32)(unsafe.Add(mBase, uint32(v22)+12))
	if v37 == int32(0) {
		goto L6
	} else {
		goto L7
	}
L5:
	;
	m.G0 = v22 + int32(16)
	return v313
L6:
	;
	v40 = int32(0)
	v43 = F_hstorePairs(m, v40, v40, v40)
	mBase = m.M
	v44 = m.ExcPending
	if v44 != 0 {
		goto L1
	} else {
		goto L9
	}
L7:
	;
	goto L8
L8:
	;
	v47 = F_palloc(m, v37*int32(20))
	mBase = m.M
	v48 = m.ExcPending
	if v48 != 0 {
		goto L1
	} else {
		goto L10
	}
L9:
	;
	v313 = v43
	goto L5
L10:
	;
	if v37 <= int32(0) {
		goto L12
	} else {
		goto L13
	}
L11:
	;
	v292 = F_hstorePairs(m, v47, v279, v280)
	mBase = m.M
	v293 = m.ExcPending
	if v293 != 0 {
		goto L1
	} else {
		goto L63
	}
L12:
	;
	v279 = v2
	v280 = v2
	goto L11
L13:
	;
	goto L14
L14:
	;
	v52 = v25 + int32(8)
	v53 = int32(3)
	v57 = v52 + v29<<(uint(v53)%32)&int32(2147483640)
	v58 = *(*int32)(unsafe.Add(mBase, uint32(v25)+4))
	v60 = v58 & int32(268435455)
	v65 = int32(0)
	v71 = v2
	v72 = v2
	v75 = v2
	goto L15
L15:
	;
	if v60 <= v65 {
		v251 = v65
		v257 = v71
		v258 = v72
		goto L17
	} else {
		goto L18
	}
L16:
	;
	v279 = v257
	v280 = v258
	goto L11
L17:
	;
	v271 = v75 + int32(1)
	if v271 != v37 {
		v65 = v251
		v71 = v257
		v72 = v258
		v75 = v271
		goto L15
	} else {
		goto L62
	}
L18:
	;
	v87 = v35 + v75*int32(20)
	v88 = *(*int32)(unsafe.Add(mBase, uint32(v87)+8))
	v89 = *(*int32)(unsafe.Add(mBase, uint32(v87)))
	v90 = v65
	v93 = v60
	goto L20
L19:
	;
	v239 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v207)+17)) = uint8(v239)
	*(*int32)(unsafe.Add(mBase, uint32(v207)+12)) = v236
	*(*int32)(unsafe.Add(mBase, uint32(v207)+4)) = v238
	v245 = int32(1)
	v246 = int32(base.Ui32(v213)>>(uint(int32(30))%32)) & v245
	*(*uint8)(unsafe.Add(mBase, uint32(v207)+16)) = uint8(v246)
	v251 = v202
	v257 = v71 + v245
	v258 = v236 + (v88 + v72)
	goto L17
L20:
	;
	v111 = base.I32_div_s(v93-v90, int32(2))
	v112 = v111 + v90
	v115 = v52 + v112<<(uint(int32(3))%32)
	v116 = *(*int32)(unsafe.Add(mBase, uint32(v115)))
	v118 = v116 & int32(1073741823)
	if int32(0) <= v116 {
		goto L26
	} else {
		goto L27
	}
L21:
	;
	v227 = int32(1073741823)
	v231 = *(*int32)(unsafe.Add(mBase, uint32(v212-int32(4))))
	v233 = v231 & v227
	v236 = v213&v227 - v233
	v238 = v233 + v57
	goto L19
L22:
	;
	goto L21
L23:
	;
	v223 = base.B2i32(v218 < int32(0))
	if v218 < int32(0) {
		goto L55
	} else {
		goto L56
	}
L24:
	;
	v138 = v137 + (v52 + v60<<(uint(v53)%32))
	if base.Ui32(int32(4)) <= base.Ui32(v88) {
		goto L37
	} else {
		goto L38
	}
L25:
	;
	if base.Ui32(v88) < base.Ui32(v130) {
		goto L31
	} else {
		goto L32
	}
L26:
	;
	v123 = *(*int32)(unsafe.Add(mBase, uint32(v115-int32(4))))
	v125 = v123 & int32(1073741823)
	v126 = v118 - v125
	if v126 != v88 {
		v130 = v126
		goto L25
	} else {
		goto L29
	}
L27:
	;
	goto L28
L28:
	;
	if v118 == v88 {
		v137 = int32(0)
		goto L24
	} else {
		goto L30
	}
L29:
	;
	v137 = v125
	goto L24
L30:
	;
	v130 = v118
	goto L25
L31:
	;
	v135 = int32(1)
	goto L33
L32:
	;
	v135 = int32(-1)
	goto L33
L33:
	;
	v218 = v135
	goto L23
L34:
	;
	if v200 != 0 {
		v218 = v200
		goto L23
	} else {
		goto L52
	}
L35:
	;
	v200 = int32(0)
	goto L34
L36:
	;
	v174 = v169
	v175 = v170
	v176 = v171
	goto L46
L37:
	;
	if (v138|v89)&int32(3) != 0 {
		v169 = v138
		v170 = v89
		v171 = v88
		goto L36
	} else {
		goto L40
	}
L38:
	;
	v162 = v138
	v163 = v89
	v164 = v88
	goto L39
L39:
	;
	if v164 == int32(0) {
		goto L35
	} else {
		goto L45
	}
L40:
	;
	v146 = v138
	v147 = v89
	v148 = v88
	goto L41
L41:
	;
	v151 = *(*int32)(unsafe.Add(mBase, uint32(v146)))
	v152 = *(*int32)(unsafe.Add(mBase, uint32(v147)))
	if v151 != v152 {
		v169 = v146
		v170 = v147
		v171 = v148
		goto L36
	} else {
		goto L43
	}
L42:
	;
	v162 = v157
	v163 = v155
	v164 = v159
	goto L39
L43:
	;
	v154 = int32(4)
	v155 = v147 + v154
	v157 = v146 + v154
	v159 = v148 - v154
	if base.Ui32(int32(3)) < base.Ui32(v159) {
		v146 = v157
		v147 = v155
		v148 = v159
		goto L41
	} else {
		goto L44
	}
L44:
	;
	goto L42
L45:
	;
	v169 = v162
	v170 = v163
	v171 = v164
	goto L36
L46:
	;
	v179 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v174))))
	v180 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v175))))
	if v179 == v180 {
		goto L48
	} else {
		goto L49
	}
L47:
	;
	v200 = v179 - v180
	goto L34
L48:
	;
	v182 = int32(1)
	v187 = v176 - v182
	if v187 != 0 {
		v174 = v174 + v182
		v175 = v175 + v182
		v176 = v187
		goto L46
	} else {
		goto L51
	}
L49:
	;
	goto L50
L50:
	;
	goto L47
L51:
	;
	goto L35
L52:
	;
	v202 = v112 + int32(1)
	if v112 < int32(0) {
		v251 = v202
		v257 = v71
		v258 = v72
		goto L17
	} else {
		goto L53
	}
L53:
	;
	v207 = v47 + v71*int32(20)
	*(*int32)(unsafe.Add(mBase, uint32(v207)+8)) = v88
	*(*int32)(unsafe.Add(mBase, uint32(v207))) = v89
	v212 = v115 + int32(4)
	v213 = *(*int32)(unsafe.Add(mBase, uint32(v212)))
	if int32(0) <= v213 {
		goto L22
	} else {
		goto L54
	}
L54:
	;
	v236 = v213 & int32(1073741823)
	v238 = v57
	goto L19
L55:
	;
	v224 = v112 + int32(1)
	goto L57
L56:
	;
	v224 = v90
	goto L57
L57:
	;
	if v218 < int32(0) {
		goto L58
	} else {
		goto L59
	}
L58:
	;
	v225 = v93
	goto L60
L59:
	;
	v225 = v112
	goto L60
L60:
	;
	if v224 < v225 {
		v90 = v224
		v93 = v225
		goto L20
	} else {
		goto L61
	}
L61:
	;
	v251 = v224
	v257 = v71
	v258 = v72
	goto L17
L62:
	;
	goto L16
L63:
	;
	v313 = v292
	goto L5
}
func F_hstore_subscript_fetch(m *base.Module, l0 int32, l1 int32, l2 int32) {
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
	var v21 int32
	_ = v21
	var v22 int32
	_ = v22
	var v23 int32
	_ = v23
	var v24 int32
	_ = v24
	var v25 int32
	_ = v25
	var v28 int32
	_ = v28
	var v30 int32
	_ = v30
	var v31 int32
	_ = v31
	var v34 int32
	_ = v34
	var v36 int32
	_ = v36
	var v45 int32
	_ = v45
	var v48 int32
	_ = v48
	var v49 int32
	_ = v49
	var v53 int32
	_ = v53
	var v59 int32
	_ = v59
	var v68 int32
	_ = v68
	var v70 int32
	_ = v70
	var v72 int32
	_ = v72
	var v75 int32
	_ = v75
	var v84 int32
	_ = v84
	var v85 int32
	_ = v85
	var v92 int32
	_ = v92
	var v93 int32
	_ = v93
	var v96 int32
	_ = v96
	var v97 int32
	_ = v97
	var v99 int32
	_ = v99
	var v104 int32
	_ = v104
	var v106 int32
	_ = v106
	var v107 int32
	_ = v107
	var v111 int32
	_ = v111
	var v116 int32
	_ = v116
	var v118 int32
	_ = v118
	var v120 int32
	_ = v120
	var v125 int32
	_ = v125
	var v130 int32
	_ = v130
	var v131 int32
	_ = v131
	var v132 int32
	_ = v132
	var v160 int32
	_ = v160
	var v174 int32
	_ = v174
	var v179 int32
	_ = v179
	var v180 int32
	_ = v180
	var v188 int32
	_ = v188
	var v189 int32
	_ = v189
	var v191 int32
	_ = v191
	var v196 int32
	_ = v196
	var v203 int32
	_ = v203
	var v205 int32
	_ = v205
	var v208 int32
	_ = v208
	var v209 int32
	_ = v209
	var v211 int32
	_ = v211
	var v212 int32
	_ = v212
	var v213 int32
	_ = v213
	v7 = *(*int32)(unsafe.Add(mBase, uint32(l1)+20))
	v8 = *(*int32)(unsafe.Add(mBase, uint32(v7)+20))
	v9 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v8))))
	if v9 == int32(1) {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	v12 = *(*int32)(unsafe.Add(mBase, uint32(l1)+8))
	v13 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(v12))) = uint8(v13)
	return
L2:
	;
	goto L3
L3:
	;
	v15 = *(*int32)(unsafe.Add(mBase, uint32(l1)+4))
	v16 = *(*int32)(unsafe.Add(mBase, uint32(v15)))
	v17 = F_hstoreUpgrade(m, v16)
	mBase = m.M
	v18 = m.ExcPending
	if v18 != 0 {
		goto L5
	} else {
		goto L6
	}
L4:
	;
	v191 = *(*int32)(unsafe.Add(mBase, uint32(v17)+4))
	v196 = v174 + v191<<(uint(int32(3))%32)&int32(2147483640)
	if v180 < int32(0) {
		goto L59
	} else {
		goto L60
	}
L5:
	;
	return
L6:
	;
	v20 = *(*int32)(unsafe.Add(mBase, uint32(v7)+16))
	v21 = *(*int32)(unsafe.Add(mBase, uint32(v20)))
	v22 = F_pg_detoast_datum_packed(m, v21)
	mBase = m.M
	v23 = m.ExcPending
	if v23 != 0 {
		goto L5
	} else {
		goto L7
	}
L7:
	;
	v24 = int32(1)
	v25 = v22 + v24
	v28 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v22))))
	v30 = v28 & v24
	if v30 != 0 {
		goto L8
	} else {
		goto L9
	}
L8:
	;
	v31 = v25
	goto L10
L9:
	;
	v31 = v22 + int32(4)
	goto L10
L10:
	;
	if v28 == int32(1) {
		goto L12
	} else {
		goto L13
	}
L11:
	;
	goto L26
L12:
	;
	v34 = int32(4)
	v36 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v25))))
	if v36&int32(254) == int32(2) {
		goto L15
	} else {
		goto L16
	}
L13:
	;
	goto L14
L14:
	;
	v49 = int32(1)
	if v30 != 0 {
		v59 = int32(base.Ui32(v28)>>(uint(v49)%32)) - v49
		goto L11
	} else {
		goto L21
	}
L15:
	;
	v45 = v34
	goto L17
L16:
	;
	v45 = base.B2i32(v36 == int32(18)) << (uint(v34) % 32)
	goto L17
L17:
	;
	if v36 == int32(1) {
		goto L18
	} else {
		goto L19
	}
L18:
	;
	v48 = v34
	goto L20
L19:
	;
	v48 = v45
	goto L20
L20:
	;
	v59 = v48
	goto L11
L21:
	;
	v53 = *(*int32)(unsafe.Add(mBase, uint32(v22)))
	v59 = int32(base.Ui32(v53)>>(uint(int32(2))%32)) - int32(4)
	goto L11
L22:
	;
	if int32(0) <= v160 {
		goto L54
	} else {
		goto L55
	}
L23:
	;
	goto L22
L26:
	;
	v68 = int32(0)
	goto L27
L27:
	;
	v70 = *(*int32)(unsafe.Add(mBase, uint32(v17)+4))
	v72 = v70 & int32(268435455)
	if v68 < v72 {
		goto L28
	} else {
		goto L29
	}
L28:
	;
	v75 = v17 + int32(8)
	v84 = v68
	v85 = v72
	goto L31
L29:
	;
	goto L30
L30:
	;
	v160 = int32(-1)
	goto L23
L31:
	;
	v92 = base.I32_div_s(v85-v84, int32(2))
	v93 = v92 + v84
	v96 = v75 + v93<<(uint(int32(3))%32)
	v97 = *(*int32)(unsafe.Add(mBase, uint32(v96)))
	v99 = v97 & int32(1073741823)
	if int32(0) <= v97 {
		goto L36
	} else {
		goto L37
	}
L32:
	;
	goto L30
L33:
	;
	v130 = base.B2i32(v125 < int32(0))
	if v125 < int32(0) {
		goto L46
	} else {
		goto L47
	}
L34:
	;
	v120 = F_memcmp(m, v118+(v75+v72<<(uint(int32(3))%32)), v31, v59)
	mBase = m.M
	if v120 != 0 {
		v125 = v120
		goto L33
	} else {
		goto L44
	}
L35:
	;
	if base.Ui32(v59) < base.Ui32(v111) {
		goto L41
	} else {
		goto L42
	}
L36:
	;
	v104 = *(*int32)(unsafe.Add(mBase, uint32(v96-int32(4))))
	v106 = v104 & int32(1073741823)
	v107 = v99 - v106
	if v107 != v59 {
		v111 = v107
		goto L35
	} else {
		goto L39
	}
L37:
	;
	goto L38
L38:
	;
	if v59 == v99 {
		v118 = int32(0)
		goto L34
	} else {
		goto L40
	}
L39:
	;
	v118 = v106
	goto L34
L40:
	;
	v111 = v99
	goto L35
L41:
	;
	v116 = int32(1)
	goto L43
L42:
	;
	v116 = int32(-1)
	goto L43
L43:
	;
	v125 = v116
	goto L33
L44:
	;
	v160 = v93
	goto L23
L46:
	;
	v131 = v93 + int32(1)
	goto L48
L47:
	;
	v131 = v84
	goto L48
L48:
	;
	if v125 < int32(0) {
		goto L49
	} else {
		goto L50
	}
L49:
	;
	v132 = v85
	goto L51
L50:
	;
	v132 = v93
	goto L51
L51:
	;
	if v131 < v132 {
		v84 = v131
		v85 = v132
		goto L31
	} else {
		goto L52
	}
L52:
	;
	goto L32
L54:
	;
	v174 = v17 + int32(8)
	v179 = v174 + v160<<(uint(int32(3))%32) + int32(4)
	v180 = *(*int32)(unsafe.Add(mBase, uint32(v179)))
	if v180&int32(1073741824) == int32(0) {
		goto L4
	} else {
		goto L57
	}
L55:
	;
	goto L56
L56:
	;
	v188 = *(*int32)(unsafe.Add(mBase, uint32(l1)+8))
	v189 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(v188))) = uint8(v189)
	return
L57:
	;
	goto L56
L58:
	;
	v211 = F_cstring_to_text_with_len(m, v209, v208)
	mBase = m.M
	v212 = m.ExcPending
	if v212 != 0 {
		goto L5
	} else {
		goto L62
	}
L59:
	;
	v208 = v180 & int32(1073741823)
	v209 = v196
	goto L58
L60:
	;
	goto L61
L61:
	;
	v203 = *(*int32)(unsafe.Add(mBase, uint32(v179-int32(4))))
	v205 = v203 & int32(1073741823)
	v208 = v180 - v205
	v209 = v196 + v205
	goto L58
L62:
	;
	v213 = *(*int32)(unsafe.Add(mBase, uint32(l1)+4))
	*(*int32)(unsafe.Add(mBase, uint32(v213))) = v211
	return
}
func F_hstore_to_json(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v8 int32
	_ = v8
	var v10 int32
	_ = v10
	var v12 int32
	_ = v12
	var v13 int32
	_ = v13
	var v16 int32
	_ = v16
	var v17 int32
	_ = v17
	var v19 int32
	_ = v19
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
	var v35 int32
	_ = v35
	var v37 int32
	_ = v37
	var v46 int32
	_ = v46
	var v47 int32
	_ = v47
	var v52 int32
	_ = v52
	var v56 int32
	_ = v56
	var v58 int32
	_ = v58
	var v61 int32
	_ = v61
	var v63 int32
	_ = v63
	var v65 int32
	_ = v65
	var v68 int32
	_ = v68
	var v73 int32
	_ = v73
	var v74 int32
	_ = v74
	var v79 int32
	_ = v79
	var v86 int32
	_ = v86
	var v88 int32
	_ = v88
	var v91 int32
	_ = v91
	var v93 int32
	_ = v93
	var v95 int32
	_ = v95
	var v99 int32
	_ = v99
	var v103 int32
	_ = v103
	var v107 int32
	_ = v107
	var v108 int32
	_ = v108
	var v109 int32
	_ = v109
	var v110 int32
	_ = v110
	var v111 int32
	_ = v111
	var v119 int32
	_ = v119
	v8 = m.G0
	v10 = v8 - int32(16)
	m.G0 = v10
	v12 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v13 = F_hstoreUpgrade(m, v12)
	mBase = m.M
	v16 = m.ExcPending
	if v16 != 0 {
		goto L2
	} else {
		goto L3
	}
L1:
	;
	m.G0 = v10 + int32(16)
	return v119
L2:
	;
	return int32(0)
L3:
	;
	v17 = *(*int32)(unsafe.Add(mBase, uint32(v13)+4))
	v19 = v17 & int32(268435455)
	if v19 == int32(0) {
		goto L4
	} else {
		goto L5
	}
L4:
	;
	v24 = F_cstring_to_text_with_len(m, int32(4103), int32(2))
	mBase = m.M
	v25 = m.ExcPending
	if v25 != 0 {
		goto L2
	} else {
		goto L7
	}
L5:
	;
	goto L6
L6:
	;
	v27 = v13 + int32(8)
	v30 = v27 + v19<<(uint(int32(3))%32)
	F_initStringInfo(m, v10)
	mBase = m.M
	v32 = m.ExcPending
	if v32 != 0 {
		goto L2
	} else {
		goto L8
	}
L7:
	;
	v119 = v24
	goto L1
L8:
	;
	F_appendStringInfoChar(m, v10, int32(123))
	mBase = m.M
	v35 = m.ExcPending
	if v35 != 0 {
		goto L2
	} else {
		goto L9
	}
L9:
	;
	v37 = int32(0)
	goto L10
L10:
	;
	v46 = v27 + v37<<(uint(int32(3))%32)
	v47 = *(*int32)(unsafe.Add(mBase, uint32(v46)))
	if v47 < int32(0) {
		goto L13
	} else {
		goto L14
	}
L11:
	;
	F_appendStringInfoChar(m, v10, int32(125))
	mBase = m.M
	v107 = m.ExcPending
	if v107 != 0 {
		goto L2
	} else {
		goto L33
	}
L12:
	;
	F_escape_json_with_len(m, v10, v63, v61)
	mBase = m.M
	v65 = m.ExcPending
	if v65 != 0 {
		goto L2
	} else {
		goto L16
	}
L13:
	;
	v61 = v47 & int32(1073741823)
	v63 = v30
	goto L12
L14:
	;
	goto L15
L15:
	;
	v52 = int32(1073741823)
	v56 = *(*int32)(unsafe.Add(mBase, uint32(v46-int32(4))))
	v58 = v56 & v52
	v61 = v47&v52 - v58
	v63 = v58 + v30
	goto L12
L16:
	;
	F_appendStringInfoString(m, v10, int32(721798))
	mBase = m.M
	v68 = m.ExcPending
	if v68 != 0 {
		goto L2
	} else {
		goto L17
	}
L17:
	;
	v73 = v27 + v37<<(uint(int32(3))%32) + int32(4)
	v74 = *(*int32)(unsafe.Add(mBase, uint32(v73)))
	if v74&int32(1073741824) != 0 {
		goto L19
	} else {
		goto L20
	}
L18:
	;
	v99 = v37 + int32(1)
	if v19 != v99 {
		goto L28
	} else {
		goto L29
	}
L19:
	;
	F_appendStringInfoString(m, v10, int32(298999))
	mBase = m.M
	v79 = m.ExcPending
	if v79 != 0 {
		goto L2
	} else {
		goto L22
	}
L20:
	;
	goto L21
L21:
	;
	if v74 < int32(0) {
		goto L24
	} else {
		goto L25
	}
L22:
	;
	goto L18
L23:
	;
	F_escape_json_with_len(m, v10, v93, v91)
	mBase = m.M
	v95 = m.ExcPending
	if v95 != 0 {
		goto L2
	} else {
		goto L27
	}
L24:
	;
	v91 = v74 & int32(1073741823)
	v93 = v30
	goto L23
L25:
	;
	goto L26
L26:
	;
	v86 = *(*int32)(unsafe.Add(mBase, uint32(v73-int32(4))))
	v88 = v86 & int32(1073741823)
	v91 = v74 - v88
	v93 = v88 + v30
	goto L23
L27:
	;
	goto L18
L28:
	;
	F_appendStringInfoString(m, v10, int32(722041))
	mBase = m.M
	v103 = m.ExcPending
	if v103 != 0 {
		goto L2
	} else {
		goto L31
	}
L29:
	;
	goto L30
L30:
	;
	if v99 != v19 {
		v37 = v99
		goto L10
	} else {
		goto L32
	}
L31:
	;
	goto L30
L32:
	;
	goto L11
L33:
	;
	v108 = *(*int32)(unsafe.Add(mBase, uint32(v10)))
	v109 = *(*int32)(unsafe.Add(mBase, uint32(v10)+4))
	v110 = F_cstring_to_text_with_len(m, v108, v109)
	mBase = m.M
	v111 = m.ExcPending
	if v111 != 0 {
		goto L2
	} else {
		goto L34
	}
L34:
	;
	v119 = v110
	goto L1
}
func F_hstore_to_jsonb(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v9 int32
	_ = v9
	var v11 int32
	_ = v11
	var v13 int32
	_ = v13
	var v14 int32
	_ = v14
	var v17 int32
	_ = v17
	var v18 int32
	_ = v18
	var v19 int32
	_ = v19
	var v25 int32
	_ = v25
	var v26 int32
	_ = v26
	var v28 int32
	_ = v28
	var v30 int32
	_ = v30
	var v33 int32
	_ = v33
	var v35 int32
	_ = v35
	var v47 int32
	_ = v47
	var v48 int32
	_ = v48
	var v50 int32
	_ = v50
	var v56 int32
	_ = v56
	var v57 int32
	_ = v57
	var v58 int32
	_ = v58
	var v62 int32
	_ = v62
	var v66 int32
	_ = v66
	var v74 int32
	_ = v74
	var v75 int32
	_ = v75
	var v80 int32
	_ = v80
	var v81 int32
	_ = v81
	var v95 int32
	_ = v95
	var v96 int32
	_ = v96
	var v97 int32
	_ = v97
	var v101 int32
	_ = v101
	var v105 int32
	_ = v105
	var v114 int32
	_ = v114
	var v115 int32
	_ = v115
	var v117 int32
	_ = v117
	var v131 int32
	_ = v131
	var v132 int32
	_ = v132
	var v133 int32
	_ = v133
	var v134 int32
	_ = v134
	v9 = m.G0
	v11 = v9 - int32(48)
	m.G0 = v11
	v13 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v14 = F_hstoreUpgrade(m, v13)
	mBase = m.M
	v17 = m.ExcPending
	if v17 != 0 {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	return int32(0)
L2:
	;
	v18 = *(*int32)(unsafe.Add(mBase, uint32(v14)+4))
	v19 = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v11)+44)) = v19
	v25 = F_pushJsonbValue(m, v11+int32(44), int32(6), v19)
	mBase = m.M
	v26 = m.ExcPending
	if v26 != 0 {
		goto L1
	} else {
		goto L3
	}
L3:
	;
	v28 = v18 & int32(268435455)
	if v28 != 0 {
		goto L4
	} else {
		goto L5
	}
L4:
	;
	v30 = v14 + int32(8)
	v33 = v30 + v28<<(uint(int32(3))%32)
	v35 = int32(0)
	goto L7
L5:
	;
	goto L6
L6:
	;
	v131 = F_pushJsonbValue(m, v11+int32(44), int32(7), int32(0))
	mBase = m.M
	v132 = m.ExcPending
	if v132 != 0 {
		goto L1
	} else {
		goto L24
	}
L7:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v11)+24)) = int32(1)
	v47 = v30 + v35<<(uint(int32(3))%32)
	v48 = *(*int32)(unsafe.Add(mBase, uint32(v47)))
	v50 = v48 & int32(1073741823)
	if v48 < int32(0) {
		goto L10
	} else {
		goto L11
	}
L8:
	;
	goto L6
L9:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v11)+32)) = v66 + v33
	v74 = F_pushJsonbValue(m, v11+int32(44), int32(1), v11+int32(24))
	mBase = m.M
	v75 = m.ExcPending
	if v75 != 0 {
		goto L1
	} else {
		goto L13
	}
L10:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v11)+28)) = v50
	v66 = int32(0)
	goto L9
L11:
	;
	goto L12
L12:
	;
	v56 = v47 - int32(4)
	v57 = *(*int32)(unsafe.Add(mBase, uint32(v56)))
	v58 = int32(1073741823)
	*(*int32)(unsafe.Add(mBase, uint32(v11)+28)) = v50 - v57&v58
	v62 = *(*int32)(unsafe.Add(mBase, uint32(v56)))
	v66 = v62 & v58
	goto L9
L13:
	;
	v80 = v30 + v35<<(uint(int32(3))%32) + int32(4)
	v81 = *(*int32)(unsafe.Add(mBase, uint32(v80)))
	if v81&int32(1073741824) != 0 {
		goto L15
	} else {
		goto L16
	}
L14:
	;
	v114 = F_pushJsonbValue(m, v11+int32(44), int32(2), v11+int32(4))
	mBase = m.M
	v115 = m.ExcPending
	if v115 != 0 {
		goto L1
	} else {
		goto L22
	}
L15:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v11)+4)) = int32(0)
	goto L14
L16:
	;
	goto L17
L17:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v11)+4)) = int32(1)
	if v81 < int32(0) {
		goto L19
	} else {
		goto L20
	}
L18:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v11)+12)) = v105 + v33
	goto L14
L19:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v11)+8)) = v81 & int32(1073741823)
	v105 = int32(0)
	goto L18
L20:
	;
	goto L21
L21:
	;
	v95 = v80 - int32(4)
	v96 = *(*int32)(unsafe.Add(mBase, uint32(v95)))
	v97 = int32(1073741823)
	*(*int32)(unsafe.Add(mBase, uint32(v11)+8)) = v81 - v96&v97
	v101 = *(*int32)(unsafe.Add(mBase, uint32(v95)))
	v105 = v101 & v97
	goto L18
L22:
	;
	v117 = v35 + int32(1)
	if v117 != v28 {
		v35 = v117
		goto L7
	} else {
		goto L23
	}
L23:
	;
	goto L8
L24:
	;
	v133 = F_JsonbValueToJsonb(m, v131)
	mBase = m.M
	v134 = m.ExcPending
	if v134 != 0 {
		goto L1
	} else {
		goto L25
	}
L25:
	;
	m.G0 = v11 + int32(48)
	return v133
}
