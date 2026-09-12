package p2

import (
	base "github.com/shibukawa/pgmem/internal/aot/pgaot/base"
	"unsafe"
)

func F_compute_tsvector_stats(m *base.Module, l0 int32, l1 int32, l2 int32, l3 float64) {
	mBase := m.M
	_ = mBase
	var v5 int32
	_ = v5
	var v21 float64
	_ = v21
	var v23 int32
	_ = v23
	var v25 int32
	_ = v25
	var v27 int32
	_ = v27
	var v35 int32
	_ = v35
	var v37 int32
	_ = v37
	var v42 int32
	_ = v42
	var v45 int32
	_ = v45
	var v49 int32
	_ = v49
	var v50 int32
	_ = v50
	var v60 int32
	_ = v60
	var v64 int32
	_ = v64
	var v68 int32
	_ = v68
	var v71 int32
	_ = v71
	var v74 float64
	_ = v74
	var v78 int32
	_ = v78
	var v81 int32
	_ = v81
	var v82 int32
	_ = v82
	var v83 int32
	_ = v83
	var v88 int32
	_ = v88
	var v92 int32
	_ = v92
	var v99 int32
	_ = v99
	var v105 int32
	_ = v105
	var v106 int32
	_ = v106
	var v110 int32
	_ = v110
	var v114 int32
	_ = v114
	var v115 int32
	_ = v115
	var v116 int32
	_ = v116
	var v117 int32
	_ = v117
	var v121 int32
	_ = v121
	var v132 int32
	_ = v132
	var v136 int32
	_ = v136
	var v137 int32
	_ = v137
	var v139 int32
	_ = v139
	var v148 int32
	_ = v148
	var v153 int32
	_ = v153
	var v154 int32
	_ = v154
	var v164 int32
	_ = v164
	var v165 int32
	_ = v165
	var v166 int32
	_ = v166
	var v169 int32
	_ = v169
	var v173 int32
	_ = v173
	var v178 int32
	_ = v178
	var v179 int32
	_ = v179
	var v180 int32
	_ = v180
	var v182 int32
	_ = v182
	var v183 int32
	_ = v183
	var v184 int32
	_ = v184
	var v188 int32
	_ = v188
	var v189 int32
	_ = v189
	var v195 int32
	_ = v195
	var v198 int32
	_ = v198
	var v199 int32
	_ = v199
	var v204 int32
	_ = v204
	var v222 int32
	_ = v222
	var v223 int32
	_ = v223
	var v226 int32
	_ = v226
	var v229 int32
	_ = v229
	var v230 int32
	_ = v230
	var v234 int32
	_ = v234
	var v238 int32
	_ = v238
	var v239 int32
	_ = v239
	var v270 int32
	_ = v270
	var v289 int32
	_ = v289
	var v290 int32
	_ = v290
	var v298 int32
	_ = v298
	var v302 int32
	_ = v302
	var v315 float64
	_ = v315
	var v318 int32
	_ = v318
	var v325 int32
	_ = v325
	var v329 int32
	_ = v329
	var v333 int32
	_ = v333
	var v339 float64
	_ = v339
	var v342 int32
	_ = v342
	var v354 int32
	_ = v354
	var v358 int32
	_ = v358
	var v364 float64
	_ = v364
	var v367 int32
	_ = v367
	var v372 float32
	_ = v372
	var v375 float64
	_ = v375
	var v376 float64
	_ = v376
	var v380 int32
	_ = v380
	var v382 int32
	_ = v382
	var v389 int32
	_ = v389
	var v390 int32
	_ = v390
	var v391 int32
	_ = v391
	var v392 int32
	_ = v392
	var v393 int32
	_ = v393
	var v394 int32
	_ = v394
	var v395 int32
	_ = v395
	var v396 int32
	_ = v396
	var v397 int32
	_ = v397
	var v398 int32
	_ = v398
	var v399 int32
	_ = v399
	var v400 int32
	_ = v400
	var v401 int32
	_ = v401
	var v402 int32
	_ = v402
	var v403 int32
	_ = v403
	var v404 int32
	_ = v404
	var v405 int32
	_ = v405
	var v406 int32
	_ = v406
	var v407 int32
	_ = v407
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
	var v413 int32
	_ = v413
	var v414 int32
	_ = v414
	var v415 int32
	_ = v415
	var v416 int32
	_ = v416
	var v417 int32
	_ = v417
	var v420 int32
	_ = v420
	var v421 int32
	_ = v421
	var v422 int32
	_ = v422
	var v423 int32
	_ = v423
	var v424 int32
	_ = v424
	var v456 int32
	_ = v456
	var v459 int32
	_ = v459
	var v460 int32
	_ = v460
	var v464 int32
	_ = v464
	var v467 int32
	_ = v467
	var v470 int32
	_ = v470
	var v471 int32
	_ = v471
	var v476 int32
	_ = v476
	var v482 int32
	_ = v482
	var v484 int32
	_ = v484
	var v485 int32
	_ = v485
	var v487 int32
	_ = v487
	var v500 int32
	_ = v500
	var v506 int32
	_ = v506
	var v508 int32
	_ = v508
	var v512 int32
	_ = v512
	var v514 int32
	_ = v514
	var v515 int32
	_ = v515
	var v516 int32
	_ = v516
	var v519 int32
	_ = v519
	var v520 int32
	_ = v520
	var v528 int32
	_ = v528
	var v529 int32
	_ = v529
	var v544 float64
	_ = v544
	var v547 int32
	_ = v547
	var v548 int32
	_ = v548
	var v556 int32
	_ = v556
	var v561 int32
	_ = v561
	var v567 int32
	_ = v567
	var v573 int32
	_ = v573
	var v574 int32
	_ = v574
	var v575 int32
	_ = v575
	var v576 int32
	_ = v576
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
	var v592 int32
	_ = v592
	var v593 int32
	_ = v593
	var v594 int32
	_ = v594
	var v595 int32
	_ = v595
	var v598 int32
	_ = v598
	var v599 int32
	_ = v599
	var v604 int32
	_ = v604
	var v623 int32
	_ = v623
	var v626 int32
	_ = v626
	var v627 int32
	_ = v627
	var v628 int32
	_ = v628
	var v629 int32
	_ = v629
	var v630 int32
	_ = v630
	var v633 int32
	_ = v633
	var v639 int32
	_ = v639
	var v643 int32
	_ = v643
	var v658 int32
	_ = v658
	var v662 int32
	_ = v662
	var v664 int32
	_ = v664
	var v666 int32
	_ = v666
	var v675 int32
	_ = v675
	var v705 int32
	_ = v705
	var v709 int32
	_ = v709
	var v714 int32
	_ = v714
	v5 = int32(0)
	v21 = float64(0)
	v23 = m.G0
	v25 = v23 - int32(112)
	m.G0 = v25
	v27 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	*(*int32)(unsafe.Add(mBase, uint32(v25)+72)) = int32(1182)
	*(*int32)(unsafe.Add(mBase, uint32(v25)+68)) = int32(1183)
	*(*int64)(unsafe.Add(mBase, uint32(v25)+60)) = int64(68719476744)
	v35 = *(*int32)(unsafe.Add(mBase, _consts[0]))
	*(*int32)(unsafe.Add(mBase, uint32(v25)+84)) = v35
	v37 = int32(10000)
	v42 = base.I32_div_s(v27*v37+v37, int32(7))
	v45 = v27 * int32(10)
	v49 = F_hash_create(m, int32(389041), v45, v25+int32(44), int32(1224))
	mBase = m.M
	v50 = m.ExcPending
	if v50 != 0 {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	return
L2:
	;
	if l2 <= int32(0) {
		goto L5
	} else {
		goto L6
	}
L3:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v705 = m.ExcPending
	if v705 != 0 {
		goto L1
	} else {
		goto L114
	}
L4:
	;
	if v358 < l2 {
		goto L64
	} else {
		goto L65
	}
L5:
	;
	v354 = v5
	v358 = v5
	v364 = v21
	goto L4
L6:
	;
	goto L7
L7:
	;
	v60 = int32(1)
	v64 = v5
	v68 = v5
	v71 = v5
	v74 = v21
	goto L8
L8:
	;
	F_vacuum_delay_point(m, int32(1))
	mBase = m.M
	v78 = m.ExcPending
	if v78 != 0 {
		goto L1
	} else {
		goto L10
	}
L9:
	;
	v354 = v329
	v358 = v333
	v364 = v339
	goto L4
L10:
	;
	v81 = m.T0[l1].(func(*base.Module, int32, int32, int32) int32)(m, l0, v71, v25+int32(35))
	mBase = m.M
	v82 = m.ExcPending
	if v82 != 0 {
		goto L1
	} else {
		goto L11
	}
L11:
	;
	v83 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v25)+35)))
	if v83 == int32(1) {
		goto L13
	} else {
		goto L14
	}
L12:
	;
	v342 = v71 + int32(1)
	if v342 != l2 {
		v60 = v325
		v64 = v329
		v68 = v333
		v71 = v342
		v74 = v339
		goto L8
	} else {
		goto L62
	}
L13:
	;
	v325 = v60
	v329 = v64
	v333 = v68 + int32(1)
	v339 = v74
	goto L12
L14:
	;
	goto L15
L15:
	;
	v88 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v81))))
	if v88 == int32(1) {
		goto L17
	} else {
		goto L18
	}
L16:
	;
	v115 = F_pg_detoast_datum(m, v81)
	mBase = m.M
	v116 = m.ExcPending
	if v116 != 0 {
		goto L1
	} else {
		goto L25
	}
L17:
	;
	v92 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v81)+1)))
	if base.Ui32((v92-int32(1))&int32(255)) < base.Ui32(int32(3)) {
		v114 = int32(6)
		goto L16
	} else {
		goto L20
	}
L18:
	;
	goto L19
L19:
	;
	v106 = int32(1)
	if v88&v106 != 0 {
		v114 = int32(base.Ui32(v88) >> (uint(v106) % 32))
		goto L16
	} else {
		goto L24
	}
L20:
	;
	v99 = int32(18)
	if v92&int32(255) == v99 {
		goto L21
	} else {
		goto L22
	}
L21:
	;
	v105 = v99
	goto L23
L22:
	;
	v105 = int32(2)
	goto L23
L23:
	;
	v114 = v105
	goto L16
L24:
	;
	v110 = *(*int32)(unsafe.Add(mBase, uint32(v81)))
	v114 = int32(base.Ui32(v110) >> (uint(int32(2)) % 32))
	goto L16
L25:
	;
	v117 = *(*int32)(unsafe.Add(mBase, uint32(v115)+4))
	if int32(0) < v117 {
		goto L26
	} else {
		goto L27
	}
L26:
	;
	v121 = v115 + int32(8)
	v132 = v60
	v136 = v64
	v137 = v121
	v139 = int32(0)
	goto L29
L27:
	;
	v298 = v60
	v302 = v64
	goto L28
L28:
	;
	v315 = base.F64_add(v74, base.F64_convert_i32_u(v114))
	if v115 == v81 {
		v325 = v298
		v329 = v302
		v333 = v68
		v339 = v315
		goto L12
	} else {
		goto L60
	}
L29:
	;
	v148 = *(*int32)(unsafe.Add(mBase, uint32(v137)))
	*(*int32)(unsafe.Add(mBase, uint32(v25)+36)) = v121 + v117<<(uint(int32(2))%32) + int32(base.Ui32(v148)>>(uint(int32(12))%32))
	v153 = *(*int32)(unsafe.Add(mBase, uint32(v137)))
	v154 = int32(1)
	*(*int32)(unsafe.Add(mBase, uint32(v25)+40)) = int32(base.Ui32(v153)>>(uint(v154)%32)) & int32(2047)
	v164 = F_hash_search(m, v49, v25+int32(36), v154, v25+int32(34))
	mBase = m.M
	v165 = m.ExcPending
	if v165 != 0 {
		goto L1
	} else {
		goto L31
	}
L30:
	;
	v298 = v270
	v302 = v188
	goto L28
L31:
	;
	v166 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v25)+34)))
	if v166 == int32(1) {
		goto L33
	} else {
		goto L34
	}
L32:
	;
	v188 = v136 + int32(1)
	v189 = base.I32_rem_s(v188, v42)
	if v189 == int32(0) {
		goto L41
	} else {
		goto L42
	}
L33:
	;
	v169 = *(*int32)(unsafe.Add(mBase, uint32(v164)+8))
	*(*int32)(unsafe.Add(mBase, uint32(v164)+8)) = v169 + int32(1)
	goto L32
L34:
	;
	goto L35
L35:
	;
	v173 = int32(1)
	*(*int32)(unsafe.Add(mBase, uint32(v164)+8)) = v173
	*(*int32)(unsafe.Add(mBase, uint32(v164)+12)) = v132 - v173
	v178 = *(*int32)(unsafe.Add(mBase, uint32(v25)+40))
	v179 = F_palloc(m, v178)
	mBase = m.M
	v180 = m.ExcPending
	if v180 != 0 {
		goto L1
	} else {
		goto L36
	}
L36:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v164))) = v179
	v182 = *(*int32)(unsafe.Add(mBase, uint32(v25)+36))
	v183 = *(*int32)(unsafe.Add(mBase, uint32(v25)+40))
	if v183 != 0 {
		goto L38
	} else {
		goto L39
	}
L37:
	;
	goto L32
L38:
	;
	v184 = F__emscripten_memcpy_bulkmem(m, v179, v182, v183)
	mBase = m.M
	goto L40
L39:
	;
	goto L40
L40:
	;
	goto L37
L41:
	;
	F_hash_seq_init(m, v25+int32(92), v49)
	mBase = m.M
	v195 = m.ExcPending
	if v195 != 0 {
		goto L1
	} else {
		goto L44
	}
L42:
	;
	v270 = v132
	goto L43
L43:
	;
	v289 = v139 + int32(1)
	v290 = *(*int32)(unsafe.Add(mBase, uint32(v115)+4))
	if v289 < v290 {
		v132 = v270
		v136 = v188
		v137 = v137 + int32(4)
		v139 = v289
		goto L29
	} else {
		goto L59
	}
L44:
	;
	v198 = F_hash_seq_search(m, v25+int32(92))
	mBase = m.M
	v199 = m.ExcPending
	if v199 != 0 {
		goto L1
	} else {
		goto L45
	}
L45:
	;
	if v198 != 0 {
		goto L46
	} else {
		goto L47
	}
L46:
	;
	v204 = v198
	goto L49
L47:
	;
	goto L48
L48:
	;
	v270 = v132 + int32(1)
	goto L43
L49:
	;
	v222 = *(*int32)(unsafe.Add(mBase, uint32(v204)+12))
	v223 = *(*int32)(unsafe.Add(mBase, uint32(v204)+8))
	if v222+v223 <= v132 {
		goto L51
	} else {
		goto L52
	}
L50:
	;
	goto L48
L51:
	;
	v226 = *(*int32)(unsafe.Add(mBase, uint32(v204)))
	v229 = F_hash_search(m, v49, v204, int32(2), int32(0))
	mBase = m.M
	v230 = m.ExcPending
	if v230 != 0 {
		goto L1
	} else {
		goto L54
	}
L52:
	;
	goto L53
L53:
	;
	v238 = F_hash_seq_search(m, v25+int32(92))
	mBase = m.M
	v239 = m.ExcPending
	if v239 != 0 {
		goto L1
	} else {
		goto L57
	}
L54:
	;
	if v229 == int32(0) {
		goto L3
	} else {
		goto L55
	}
L55:
	;
	F_pfree(m, v226)
	mBase = m.M
	v234 = m.ExcPending
	if v234 != 0 {
		goto L1
	} else {
		goto L56
	}
L56:
	;
	goto L53
L57:
	;
	if v238 != 0 {
		v204 = v238
		goto L49
	} else {
		goto L58
	}
L58:
	;
	goto L50
L59:
	;
	goto L30
L60:
	;
	F_pfree(m, v115)
	mBase = m.M
	v318 = m.ExcPending
	if v318 != 0 {
		goto L1
	} else {
		goto L61
	}
L61:
	;
	v325 = v298
	v329 = v302
	v333 = v68
	v339 = v315
	goto L12
L62:
	;
	goto L9
L63:
	;
	m.G0 = v25 + int32(112)
	return
L64:
	;
	v367 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(l0)+36)) = uint8(v367)
	v372 = base.F32_demote_f64(base.F64_div(base.F64_convert_i32_s(v358), base.F64_convert_i32_s(l2)))
	*(*float32)(unsafe.Add(mBase, uint32(l0)+40)) = v372
	v375 = base.F64_convert_i32_s(l2 - v358)
	v376 = base.F64_div(v364, v375)
	if base.F64_lt(base.F64_abs(v376), float64(2.147483648e+09)) != 0 {
		goto L68
	} else {
		goto L69
	}
L65:
	;
	goto L66
L66:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+48)) = int32(0)
	*(*int64)(unsafe.Add(mBase, uint32(l0)+40)) = int64(1065353216)
	v675 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(l0)+36)) = uint8(v675)
	goto L63
L67:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+44)) = v382
	*(*float32)(unsafe.Add(mBase, uint32(l0)+48)) = base.F32_neg(base.F32_sub(float32(1), v372))
	v389 = *(*int32)(unsafe.Add(mBase, uint32(v49)))
	v390 = *(*int32)(unsafe.Add(mBase, uint32(v389)+4))
	v391 = *(*int32)(unsafe.Add(mBase, uint32(v389)+412))
	if v391 != 0 {
		goto L72
	} else {
		goto L73
	}
L68:
	;
	v380 = base.I32_trunc_f64_s(v376)
	v382 = v380
	goto L67
L69:
	;
	goto L70
L70:
	;
	v382 = int32(-2147483648)
	goto L67
L71:
	;
	v459 = F_palloc(m, v456<<(uint(int32(2))%32))
	mBase = m.M
	v460 = m.ExcPending
	if v460 != 0 {
		goto L1
	} else {
		goto L75
	}
L72:
	;
	v392 = *(*int32)(unsafe.Add(mBase, uint32(v389)+376))
	v393 = *(*int32)(unsafe.Add(mBase, uint32(v389)+364))
	v394 = *(*int32)(unsafe.Add(mBase, uint32(v389)+352))
	v395 = *(*int32)(unsafe.Add(mBase, uint32(v389)+340))
	v396 = *(*int32)(unsafe.Add(mBase, uint32(v389)+328))
	v397 = *(*int32)(unsafe.Add(mBase, uint32(v389)+316))
	v398 = *(*int32)(unsafe.Add(mBase, uint32(v389)+304))
	v399 = *(*int32)(unsafe.Add(mBase, uint32(v389)+292))
	v400 = *(*int32)(unsafe.Add(mBase, uint32(v389)+280))
	v401 = *(*int32)(unsafe.Add(mBase, uint32(v389)+268))
	v402 = *(*int32)(unsafe.Add(mBase, uint32(v389)+256))
	v403 = *(*int32)(unsafe.Add(mBase, uint32(v389)+244))
	v404 = *(*int32)(unsafe.Add(mBase, uint32(v389)+232))
	v405 = *(*int32)(unsafe.Add(mBase, uint32(v389)+220))
	v406 = *(*int32)(unsafe.Add(mBase, uint32(v389)+208))
	v407 = *(*int32)(unsafe.Add(mBase, uint32(v389)+196))
	v408 = *(*int32)(unsafe.Add(mBase, uint32(v389)+184))
	v409 = *(*int32)(unsafe.Add(mBase, uint32(v389)+172))
	v410 = *(*int32)(unsafe.Add(mBase, uint32(v389)+160))
	v411 = *(*int32)(unsafe.Add(mBase, uint32(v389)+148))
	v412 = *(*int32)(unsafe.Add(mBase, uint32(v389)+136))
	v413 = *(*int32)(unsafe.Add(mBase, uint32(v389)+124))
	v414 = *(*int32)(unsafe.Add(mBase, uint32(v389)+112))
	v415 = *(*int32)(unsafe.Add(mBase, uint32(v389)+100))
	v416 = *(*int32)(unsafe.Add(mBase, uint32(v389)+88))
	v417 = *(*int32)(unsafe.Add(mBase, uint32(v389)+76))
	v420 = *(*int32)(unsafe.Add(mBase, uint32(v389-int32(-64))))
	v421 = *(*int32)(unsafe.Add(mBase, uint32(v389)+52))
	v422 = *(*int32)(unsafe.Add(mBase, uint32(v389)+40))
	v423 = *(*int32)(unsafe.Add(mBase, uint32(v389)+28))
	v424 = *(*int32)(unsafe.Add(mBase, uint32(v389)+16))
	v456 = v392 + (v393 + (v394 + (v395 + (v396 + (v397 + (v398 + (v399 + (v400 + (v401 + (v402 + (v403 + (v404 + (v405 + (v406 + (v407 + (v408 + (v409 + (v410 + (v411 + (v412 + (v413 + (v414 + (v415 + (v416 + (v417 + (v420 + (v421 + (v422 + (v423 + (v424 + v390))))))))))))))))))))))))))))))
	goto L74
L73:
	;
	v456 = v390
	goto L74
L74:
	;
	goto L71
L75:
	;
	F_hash_seq_init(m, v25+int32(92), v49)
	mBase = m.M
	v464 = m.ExcPending
	if v464 != 0 {
		goto L1
	} else {
		goto L76
	}
L76:
	;
	v467 = base.I32_div_s(v354*int32(9), v42)
	v470 = F_hash_seq_search(m, v25+int32(92))
	mBase = m.M
	v471 = m.ExcPending
	if v471 != 0 {
		goto L1
	} else {
		goto L78
	}
L77:
	;
	v547 = F_errstart(m, int32(12), int32(0))
	mBase = m.M
	v548 = m.ExcPending
	if v548 != 0 {
		goto L1
	} else {
		goto L95
	}
L78:
	;
	if v470 == int32(0) {
		goto L79
	} else {
		goto L80
	}
L79:
	;
	v528 = int32(0)
	v529 = v354
	v544 = float64(0)
	goto L77
L80:
	;
	goto L81
L81:
	;
	v476 = int32(0)
	v482 = v470
	v484 = v476
	v485 = v354
	v487 = v476
	goto L82
L82:
	;
	v500 = *(*int32)(unsafe.Add(mBase, uint32(v482)+8))
	if v467 < v500 {
		goto L84
	} else {
		goto L85
	}
L83:
	;
	v528 = v514
	v529 = v515
	v544 = base.F64_convert_i32_u(v516)
	goto L77
L84:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v459+v484<<(uint(int32(2))%32)))) = v482
	v506 = *(*int32)(unsafe.Add(mBase, uint32(v482)+8))
	if v506 < v487 {
		goto L87
	} else {
		goto L88
	}
L85:
	;
	v514 = v484
	v515 = v485
	v516 = v487
	goto L86
L86:
	;
	v519 = F_hash_seq_search(m, v25+int32(92))
	mBase = m.M
	v520 = m.ExcPending
	if v520 != 0 {
		goto L1
	} else {
		goto L93
	}
L87:
	;
	v508 = v487
	goto L89
L88:
	;
	v508 = v506
	goto L89
L89:
	;
	if v485 < v506 {
		goto L90
	} else {
		goto L91
	}
L90:
	;
	v512 = v485
	goto L92
L91:
	;
	v512 = v506
	goto L92
L92:
	;
	v514 = v484 + int32(1)
	v515 = v512
	v516 = v508
	goto L86
L93:
	;
	if v519 != 0 {
		v482 = v519
		v484 = v514
		v485 = v515
		v487 = v516
		goto L82
	} else {
		goto L94
	}
L94:
	;
	goto L83
L95:
	;
	if v547 != 0 {
		goto L96
	} else {
		goto L97
	}
L96:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v25)+16)) = v528
	*(*int32)(unsafe.Add(mBase, uint32(v25)+12)) = v456
	*(*int32)(unsafe.Add(mBase, uint32(v25)+8)) = v354
	*(*int32)(unsafe.Add(mBase, uint32(v25)+4)) = v42
	*(*int32)(unsafe.Add(mBase, uint32(v25))) = v45
	F_errmsg_internal(m, int32(475292), v25)
	mBase = m.M
	v556 = m.ExcPending
	if v556 != 0 {
		goto L1
	} else {
		goto L99
	}
L97:
	;
	goto L98
L98:
	;
	if v528 <= v45 {
		goto L102
	} else {
		goto L103
	}
L99:
	;
	F_errfinish(m, int32(493349), int32(343), int32(124659))
	mBase = m.M
	v561 = m.ExcPending
	if v561 != 0 {
		goto L1
	} else {
		goto L100
	}
L100:
	;
	goto L98
L101:
	;
	if v576 <= int32(0) {
		goto L63
	} else {
		goto L106
	}
L102:
	;
	v575 = v529
	v576 = v528
	goto L101
L103:
	;
	goto L104
L104:
	;
	F_qsort_interruptible(m, v459, v528, int32(4), int32(1184), int32(0))
	mBase = m.M
	v567 = m.ExcPending
	if v567 != 0 {
		goto L1
	} else {
		goto L105
	}
L105:
	;
	v573 = *(*int32)(unsafe.Add(mBase, uint32(v459+v45<<(uint(int32(2))%32)-int32(4))))
	v574 = *(*int32)(unsafe.Add(mBase, uint32(v573)+8))
	v575 = v574
	v576 = v45
	goto L101
L106:
	;
	v579 = int32(0)
	F_qsort_interruptible(m, v459, v576, int32(4), int32(1185), v579)
	mBase = m.M
	v584 = m.ExcPending
	if v584 != 0 {
		goto L1
	} else {
		goto L107
	}
L107:
	;
	v585 = int32(4480304)
	v586 = *(*int32)(unsafe.Add(mBase, _consts[0]))
	v588 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	*(*int32)(unsafe.Add(mBase, _consts[0])) = v588
	v592 = F_palloc(m, v576<<(uint(int32(2))%32))
	mBase = m.M
	v593 = m.ExcPending
	if v593 != 0 {
		goto L1
	} else {
		goto L108
	}
L108:
	;
	v594 = int32(2)
	v595 = v576 + v594
	v598 = F_palloc(m, v595<<(uint(v594)%32))
	mBase = m.M
	v599 = m.ExcPending
	if v599 != 0 {
		goto L1
	} else {
		goto L109
	}
L109:
	;
	v604 = v579
	goto L110
L110:
	;
	v623 = v604 << (uint(int32(2)) % 32)
	v626 = *(*int32)(unsafe.Add(mBase, uint32(v623+v459)))
	v627 = *(*int32)(unsafe.Add(mBase, uint32(v626)))
	v628 = *(*int32)(unsafe.Add(mBase, uint32(v626)+4))
	v629 = F_cstring_to_text_with_len(m, v627, v628)
	mBase = m.M
	v630 = m.ExcPending
	if v630 != 0 {
		goto L1
	} else {
		goto L112
	}
L111:
	;
	v643 = v598 + v576<<(uint(int32(2))%32)
	*(*float32)(unsafe.Add(mBase, uint32(v643)+4)) = base.F32_demote_f64(base.F64_div(v544, v375))
	*(*float32)(unsafe.Add(mBase, uint32(v643))) = base.F32_demote_f64(base.F64_div(base.F64_convert_i32_s(v575), v375))
	*(*int32)(unsafe.Add(mBase, _consts[0])) = v586
	*(*int32)(unsafe.Add(mBase, uint32(l0)+124)) = v598
	*(*int32)(unsafe.Add(mBase, uint32(l0)+84)) = int32(100)
	*(*int32)(unsafe.Add(mBase, uint32(l0)+64)) = int32(98)
	v658 = int32(4)
	*(*uint16)(unsafe.Add(mBase, uint32(l0)+52)) = uint16(v658)
	*(*int32)(unsafe.Add(mBase, uint32(l0)+164)) = v592
	*(*int32)(unsafe.Add(mBase, uint32(l0)+104)) = v595
	v662 = int32(105)
	*(*uint8)(unsafe.Add(mBase, uint32(l0)+219)) = uint8(v662)
	v664 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(l0)+214)) = uint8(v664)
	v666 = int32(65535)
	*(*uint16)(unsafe.Add(mBase, uint32(l0)+204)) = uint16(v666)
	*(*int32)(unsafe.Add(mBase, uint32(l0)+184)) = int32(25)
	*(*int32)(unsafe.Add(mBase, uint32(l0)+144)) = v576
	goto L63
L112:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v592+v623))) = v629
	v633 = *(*int32)(unsafe.Add(mBase, uint32(v626)+8))
	*(*float32)(unsafe.Add(mBase, uint32(v623+v598))) = base.F32_demote_f64(base.F64_div(base.F64_convert_i32_s(v633), v375))
	v639 = v604 + int32(1)
	if v639 != v576 {
		v604 = v639
		goto L110
	} else {
		goto L113
	}
L113:
	;
	goto L111
L114:
	;
	F_errmsg_internal(m, int32(440684), int32(0))
	mBase = m.M
	v709 = m.ExcPending
	if v709 != 0 {
		goto L1
	} else {
		goto L115
	}
L115:
	;
	F_errfinish(m, int32(493349), int32(467), int32(388064))
	mBase = m.M
	v714 = m.ExcPending
	if v714 != 0 {
		goto L1
	} else {
		goto L116
	}
L116:
	;
	base.Wasm_trap_unreachable()
	for {
	}
}
func F_tsvector_delete_arr(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v2 int32
	_ = v2
	var v21 int32
	_ = v21
	var v23 int32
	_ = v23
	var v25 int32
	_ = v25
	var v26 int32
	_ = v26
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
	var v45 int32
	_ = v45
	var v46 int32
	_ = v46
	var v47 int32
	_ = v47
	var v51 int32
	_ = v51
	var v56 int32
	_ = v56
	var v59 int32
	_ = v59
	var v65 int32
	_ = v65
	var v72 int32
	_ = v72
	var v74 int32
	_ = v74
	var v75 int32
	_ = v75
	var v78 int32
	_ = v78
	var v79 int32
	_ = v79
	var v82 int32
	_ = v82
	var v83 int32
	_ = v83
	var v84 int32
	_ = v84
	var v85 int32
	_ = v85
	var v89 int32
	_ = v89
	var v101 int32
	_ = v101
	var v107 int32
	_ = v107
	var v116 int32
	_ = v116
	var v117 int32
	_ = v117
	var v118 int32
	_ = v118
	var v122 int32
	_ = v122
	var v126 int32
	_ = v126
	var v131 int32
	_ = v131
	var v136 int32
	_ = v136
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
	var v203 int32
	_ = v203
	var v206 int32
	_ = v206
	var v212 int32
	_ = v212
	var v214 int32
	_ = v214
	var v225 int32
	_ = v225
	var v230 int32
	_ = v230
	var v239 int32
	_ = v239
	var v247 int32
	_ = v247
	var v253 int32
	_ = v253
	var v269 int32
	_ = v269
	var v270 int32
	_ = v270
	var v272 int32
	_ = v272
	var v273 int32
	_ = v273
	var v276 int32
	_ = v276
	var v277 int32
	_ = v277
	var v280 int32
	_ = v280
	v2 = int32(0)
	v21 = m.G0
	v23 = v21 - int32(16)
	m.G0 = v23
	v25 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v26 = F_pg_detoast_datum(m, v25)
	mBase = m.M
	v29 = m.ExcPending
	if v29 != 0 {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	return int32(0)
L2:
	;
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
	F_deconstruct_array_builtin(m, v31, int32(25), v23+int32(8), v23+int32(4), v23+int32(12))
	mBase = m.M
	v41 = m.ExcPending
	if v41 != 0 {
		goto L1
	} else {
		goto L4
	}
L4:
	;
	v42 = *(*int32)(unsafe.Add(mBase, uint32(v23)+12))
	v45 = F_palloc0(m, v42<<(uint(int32(2))%32))
	mBase = m.M
	v46 = m.ExcPending
	if v46 != 0 {
		goto L1
	} else {
		goto L5
	}
L5:
	;
	v47 = *(*int32)(unsafe.Add(mBase, uint32(v23)+12))
	if v47 <= int32(0) {
		goto L7
	} else {
		goto L8
	}
L6:
	;
	v269 = F_tsvector_delete_by_indices(m, v26, v45, v253)
	mBase = m.M
	v270 = m.ExcPending
	if v270 != 0 {
		goto L1
	} else {
		goto L58
	}
L7:
	;
	v253 = v2
	goto L6
L8:
	;
	goto L9
L9:
	;
	v51 = v26 + int32(8)
	v56 = v2
	v59 = v2
	v65 = v47
	goto L10
L10:
	;
	v72 = *(*int32)(unsafe.Add(mBase, uint32(v23)+4))
	v74 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v72+v59))))
	if v74 != 0 {
		v230 = v56
		v239 = v65
		goto L12
	} else {
		goto L13
	}
L11:
	;
	v253 = v230
	goto L6
L12:
	;
	v247 = v59 + int32(1)
	if v247 < v239 {
		v56 = v230
		v59 = v247
		v65 = v239
		goto L10
	} else {
		goto L57
	}
L13:
	;
	v75 = *(*int32)(unsafe.Add(mBase, uint32(v26)+4))
	if v75 <= int32(0) {
		v230 = v56
		v239 = v65
		goto L12
	} else {
		goto L14
	}
L14:
	;
	v78 = *(*int32)(unsafe.Add(mBase, uint32(v23)+8))
	v79 = int32(2)
	v82 = *(*int32)(unsafe.Add(mBase, uint32(v78+v59<<(uint(v79)%32))))
	v83 = int32(4)
	v84 = v82 + v83
	v85 = *(*int32)(unsafe.Add(mBase, uint32(v82)))
	v89 = int32(base.Ui32(v85)>>(uint(v79)%32)) - v83
	v101 = v75
	v107 = int32(0)
	goto L15
L15:
	;
	v116 = v101 + v107
	v117 = int32(2)
	v118 = base.I32_div_s(v116, v117)
	v122 = *(*int32)(unsafe.Add(mBase, uint32(v51+v118<<(uint(v117)%32))))
	v126 = int32(base.Ui32(v122)>>(uint(int32(1))%32)) & int32(2047)
	if v89 == int32(0) {
		goto L22
	} else {
		goto L23
	}
L16:
	;
	if v116 < int32(-1) {
		v230 = v56
		v239 = v65
		goto L12
	} else {
		goto L56
	}
L17:
	;
	goto L16
L18:
	;
	if v214 < v212 {
		v101 = v212
		v107 = v214
		goto L15
	} else {
		goto L55
	}
L19:
	;
	v212 = v101
	v214 = v118 + int32(1)
	goto L18
L20:
	;
	if v206 == int32(0) {
		goto L17
	} else {
		goto L54
	}
L21:
	;
	if int32(0) <= v203 {
		v206 = v203
		goto L20
	} else {
		goto L53
	}
L22:
	;
	if v126 != 0 {
		goto L25
	} else {
		goto L26
	}
L23:
	;
	goto L24
L24:
	;
	if v126 == int32(0) {
		v206 = base.B2i32(base.Ui32(int32(19)) < base.Ui32(v85))
		goto L20
	} else {
		goto L28
	}
L25:
	;
	v131 = int32(-1)
	goto L27
L26:
	;
	v131 = int32(0)
	goto L27
L27:
	;
	v203 = v131
	goto L21
L28:
	;
	v136 = v51 + v75<<(uint(v79)%32) + int32(base.Ui32(v122)>>(uint(int32(12))%32))
	if base.Ui32(v89) < base.Ui32(v126) {
		goto L29
	} else {
		goto L30
	}
L29:
	;
	v138 = v89
	goto L31
L30:
	;
	v138 = v126
	goto L31
L31:
	;
	if base.Ui32(int32(4)) <= base.Ui32(v138) {
		goto L35
	} else {
		goto L36
	}
L32:
	;
	if v200 != 0 {
		v203 = v200
		goto L21
	} else {
		goto L50
	}
L33:
	;
	v200 = int32(0)
	goto L32
L34:
	;
	v174 = v169
	v175 = v170
	v176 = v171
	goto L44
L35:
	;
	if (v84|v136)&int32(3) != 0 {
		v169 = v84
		v170 = v136
		v171 = v138
		goto L34
	} else {
		goto L38
	}
L36:
	;
	v162 = v84
	v163 = v136
	v164 = v138
	goto L37
L37:
	;
	if v164 == int32(0) {
		goto L33
	} else {
		goto L43
	}
L38:
	;
	v146 = v84
	v147 = v136
	v148 = v138
	goto L39
L39:
	;
	v151 = *(*int32)(unsafe.Add(mBase, uint32(v146)))
	v152 = *(*int32)(unsafe.Add(mBase, uint32(v147)))
	if v151 != v152 {
		v169 = v146
		v170 = v147
		v171 = v148
		goto L34
	} else {
		goto L41
	}
L40:
	;
	v162 = v157
	v163 = v155
	v164 = v159
	goto L37
L41:
	;
	v154 = int32(4)
	v155 = v147 + v154
	v157 = v146 + v154
	v159 = v148 - v154
	if base.Ui32(int32(3)) < base.Ui32(v159) {
		v146 = v157
		v147 = v155
		v148 = v159
		goto L39
	} else {
		goto L42
	}
L42:
	;
	goto L40
L43:
	;
	v169 = v162
	v170 = v163
	v171 = v164
	goto L34
L44:
	;
	v179 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v174))))
	v180 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v175))))
	if v179 == v180 {
		goto L46
	} else {
		goto L47
	}
L45:
	;
	v200 = v179 - v180
	goto L32
L46:
	;
	v182 = int32(1)
	v187 = v176 - v182
	if v187 != 0 {
		v174 = v174 + v182
		v175 = v175 + v182
		v176 = v187
		goto L44
	} else {
		goto L49
	}
L47:
	;
	goto L48
L48:
	;
	goto L45
L49:
	;
	goto L33
L50:
	;
	if v126 == v89 {
		goto L17
	} else {
		goto L51
	}
L51:
	;
	if v126 <= v89 {
		goto L19
	} else {
		goto L52
	}
L52:
	;
	v212 = v118
	v214 = v107
	goto L18
L53:
	;
	v212 = v118
	v214 = v107
	goto L18
L54:
	;
	goto L19
L55:
	;
	v230 = v56
	v239 = v65
	goto L12
L56:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v45+v56<<(uint(int32(2))%32)))) = v118
	v225 = *(*int32)(unsafe.Add(mBase, uint32(v23)+12))
	v230 = v56 + int32(1)
	v239 = v225
	goto L12
L57:
	;
	goto L11
L58:
	;
	F_pfree(m, v45)
	mBase = m.M
	v272 = m.ExcPending
	if v272 != 0 {
		goto L1
	} else {
		goto L59
	}
L59:
	;
	v273 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	if v273 != v26 {
		goto L60
	} else {
		goto L61
	}
L60:
	;
	F_pfree(m, v26)
	mBase = m.M
	v276 = m.ExcPending
	if v276 != 0 {
		goto L1
	} else {
		goto L63
	}
L61:
	;
	goto L62
L62:
	;
	v277 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
	if v277 != v31 {
		goto L64
	} else {
		goto L65
	}
L63:
	;
	goto L62
L64:
	;
	F_pfree(m, v31)
	mBase = m.M
	v280 = m.ExcPending
	if v280 != 0 {
		goto L1
	} else {
		goto L67
	}
L65:
	;
	goto L66
L66:
	;
	m.G0 = v23 + int32(16)
	return v269
L67:
	;
	goto L66
}
func F_tsvector_gt(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v6 int32
	_ = v6
	var v7 int32
	_ = v7
	var v10 int32
	_ = v10
	var v12 int32
	_ = v12
	var v13 int32
	_ = v13
	var v14 int32
	_ = v14
	var v15 int32
	_ = v15
	var v28 int32
	_ = v28
	var v29 int32
	_ = v29
	var v30 int32
	_ = v30
	var v31 int32
	_ = v31
	var v33 int32
	_ = v33
	var v36 int32
	_ = v36
	var v38 int32
	_ = v38
	var v39 int32
	_ = v39
	var v46 int32
	_ = v46
	var v47 int32
	_ = v47
	var v48 int32
	_ = v48
	var v50 int32
	_ = v50
	var v52 int32
	_ = v52
	var v55 int32
	_ = v55
	var v63 int32
	_ = v63
	var v64 int32
	_ = v64
	var v68 int32
	_ = v68
	var v70 int32
	_ = v70
	var v71 int32
	_ = v71
	var v72 int32
	_ = v72
	var v73 int32
	_ = v73
	var v75 int32
	_ = v75
	var v80 int32
	_ = v80
	var v81 int32
	_ = v81
	var v83 int32
	_ = v83
	var v84 int32
	_ = v84
	var v85 int32
	_ = v85
	var v86 int32
	_ = v86
	var v88 int32
	_ = v88
	var v92 int32
	_ = v92
	var v98 int32
	_ = v98
	var v99 int32
	_ = v99
	var v100 int32
	_ = v100
	var v104 int32
	_ = v104
	var v113 int32
	_ = v113
	var v115 int32
	_ = v115
	var v117 int32
	_ = v117
	var v118 int32
	_ = v118
	var v124 int32
	_ = v124
	var v125 int32
	_ = v125
	var v131 int32
	_ = v131
	var v132 int32
	_ = v132
	var v133 int32
	_ = v133
	var v134 int32
	_ = v134
	var v146 int32
	_ = v146
	var v147 int32
	_ = v147
	var v148 int32
	_ = v148
	var v149 int32
	_ = v149
	var v150 int32
	_ = v150
	var v152 int32
	_ = v152
	var v153 int32
	_ = v153
	var v155 int32
	_ = v155
	var v160 int32
	_ = v160
	var v161 int32
	_ = v161
	var v162 int32
	_ = v162
	var v164 int32
	_ = v164
	var v167 int32
	_ = v167
	var v172 int32
	_ = v172
	var v187 int32
	_ = v187
	var v193 int32
	_ = v193
	var v199 int32
	_ = v199
	var v223 int32
	_ = v223
	var v224 int32
	_ = v224
	var v227 int32
	_ = v227
	var v228 int32
	_ = v228
	var v231 int32
	_ = v231
	v6 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v7 = F_pg_detoast_datum(m, v6)
	mBase = m.M
	v10 = m.ExcPending
	if v10 != 0 {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	return int32(0)
L2:
	;
	v12 = l0 + int32(28)
	v13 = *(*int32)(unsafe.Add(mBase, uint32(v12)))
	v14 = F_pg_detoast_datum(m, v13)
	mBase = m.M
	v15 = m.ExcPending
	if v15 != 0 {
		goto L1
	} else {
		goto L3
	}
L3:
	;
	v28 = *(*int32)(unsafe.Add(mBase, uint32(v7)))
	v29 = int32(2)
	v30 = int32(base.Ui32(v28) >> (uint(v29) % 32))
	v31 = *(*int32)(unsafe.Add(mBase, uint32(v14)))
	v33 = int32(base.Ui32(v31) >> (uint(v29) % 32))
	if base.Ui32(v30) < base.Ui32(v33) {
		goto L5
	} else {
		goto L6
	}
L4:
	;
	v224 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	if v224 != v7 {
		goto L67
	} else {
		goto L68
	}
L5:
	;
	v223 = int32(-1)
	goto L4
L6:
	;
	goto L7
L7:
	;
	v36 = int32(1)
	if base.Ui32(v33) < base.Ui32(v30) {
		v199 = v36
		goto L8
	} else {
		goto L9
	}
L8:
	;
	v223 = v199
	goto L4
L9:
	;
	v38 = *(*int32)(unsafe.Add(mBase, uint32(v7)+4))
	v39 = *(*int32)(unsafe.Add(mBase, uint32(v14)+4))
	if v38 < v39 {
		goto L10
	} else {
		goto L11
	}
L10:
	;
	v223 = int32(-1)
	goto L4
L11:
	;
	goto L12
L12:
	;
	if v39 < v38 {
		v199 = v36
		goto L8
	} else {
		goto L13
	}
L13:
	;
	if v38 <= int32(0) {
		goto L14
	} else {
		goto L15
	}
L14:
	;
	v223 = int32(0)
	goto L4
L15:
	;
	goto L16
L16:
	;
	v46 = int32(8)
	v47 = v14 + v46
	v48 = int32(2)
	v50 = v47 + v39<<(uint(v48)%32)
	v52 = v7 + v46
	v55 = v52 + v38<<(uint(v48)%32)
	v63 = v47
	v64 = v52
	v68 = int32(0)
	goto L17
L17:
	;
	v70 = *(*int32)(unsafe.Add(mBase, uint32(v64)))
	v71 = int32(1)
	v72 = v70 & v71
	v73 = *(*int32)(unsafe.Add(mBase, uint32(v63)))
	v75 = v73 & v71
	if v72 != v75 {
		goto L19
	} else {
		goto L20
	}
L18:
	;
	v199 = int32(0)
	goto L8
L19:
	;
	if base.Ui32(v75) < base.Ui32(v72) {
		goto L22
	} else {
		goto L23
	}
L20:
	;
	goto L21
L21:
	;
	v81 = int32(1)
	v83 = int32(2047)
	v84 = int32(base.Ui32(v73)>>(uint(v81)%32)) & v83
	v85 = int32(12)
	v86 = int32(base.Ui32(v73) >> (uint(v85) % 32))
	v88 = int32(base.Ui32(v70) >> (uint(v85) % 32))
	v92 = int32(base.Ui32(v70)>>(uint(v81)%32)) & v83
	if v92 != 0 {
		goto L26
	} else {
		goto L27
	}
L22:
	;
	v80 = int32(-1)
	goto L24
L23:
	;
	v80 = int32(1)
	goto L24
L24:
	;
	v223 = v80
	goto L4
L25:
	;
	if v72 == int32(0) {
		goto L41
	} else {
		goto L42
	}
L26:
	;
	if v84 == int32(0) {
		goto L29
	} else {
		goto L30
	}
L27:
	;
	goto L28
L28:
	;
	if v84 == int32(0) {
		goto L25
	} else {
		goto L40
	}
L29:
	;
	v223 = int32(1)
	goto L4
L30:
	;
	goto L31
L31:
	;
	v98 = base.B2i32(base.Ui32(v92) < base.Ui32(v84))
	if base.Ui32(v92) < base.Ui32(v84) {
		goto L32
	} else {
		goto L33
	}
L32:
	;
	v99 = v92
	goto L34
L33:
	;
	v99 = v84
	goto L34
L34:
	;
	v100 = F_memcmp(m, v88+v55, v86+v50, v99)
	mBase = m.M
	if v100 != 0 {
		v199 = v100
		goto L8
	} else {
		goto L35
	}
L35:
	;
	if v92 == v84 {
		goto L25
	} else {
		goto L36
	}
L36:
	;
	if base.Ui32(v92) < base.Ui32(v84) {
		goto L37
	} else {
		goto L38
	}
L37:
	;
	v104 = int32(-1)
	goto L39
L38:
	;
	v104 = int32(1)
	goto L39
L39:
	;
	v223 = v104
	goto L4
L40:
	;
	v223 = int32(-1)
	goto L4
L41:
	;
	v187 = int32(4)
	v193 = v68 + int32(1)
	if v193 != v38 {
		v63 = v63 + v187
		v64 = v64 + v187
		v68 = v193
		goto L17
	} else {
		goto L66
	}
L42:
	;
	v113 = int32(1)
	v115 = int32(4194302)
	v117 = v50 + (v84+v86+v113)&v115
	v118 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v117))))
	v124 = v55 + (v92+v88+v113)&v115
	v125 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v124))))
	if v118 == v125 {
		goto L44
	} else {
		goto L45
	}
L43:
	;
	v132 = v124
	v133 = v117
	v134 = int32(0)
	goto L51
L44:
	;
	if v125 != 0 {
		goto L43
	} else {
		goto L47
	}
L45:
	;
	goto L46
L46:
	;
	if base.Ui32(v118) < base.Ui32(v125) {
		goto L48
	} else {
		goto L49
	}
L47:
	;
	goto L41
L48:
	;
	v131 = int32(-1)
	goto L50
L49:
	;
	v131 = int32(1)
	goto L50
L50:
	;
	v223 = v131
	goto L4
L51:
	;
	v146 = int32(2)
	v147 = v132 + v146
	v148 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v147))))
	v149 = int32(16383)
	v150 = v148 & v149
	v152 = v133 + v146
	v153 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v152))))
	v155 = v153 & v149
	if v150 != v155 {
		goto L53
	} else {
		goto L54
	}
L52:
	;
	if base.Ui32(v164) < base.Ui32(v162) {
		goto L63
	} else {
		goto L64
	}
L53:
	;
	if base.Ui32(v155) < base.Ui32(v150) {
		goto L56
	} else {
		goto L57
	}
L54:
	;
	goto L55
L55:
	;
	v161 = int32(14)
	v162 = int32(base.Ui32(v148) >> (uint(v161) % 32))
	v164 = int32(base.Ui32(v153) >> (uint(v161) % 32))
	if v162 == v164 {
		goto L59
	} else {
		goto L60
	}
L56:
	;
	v160 = int32(-1)
	goto L58
L57:
	;
	v160 = int32(1)
	goto L58
L58:
	;
	v223 = v160
	goto L4
L59:
	;
	v167 = v134 + int32(1)
	if v167 == v125 {
		goto L41
	} else {
		goto L62
	}
L60:
	;
	goto L61
L61:
	;
	goto L52
L62:
	;
	v132 = v147
	v133 = v152
	v134 = v167
	goto L51
L63:
	;
	v172 = int32(-1)
	goto L65
L64:
	;
	v172 = int32(1)
	goto L65
L65:
	;
	v199 = v172
	goto L8
L66:
	;
	goto L18
L67:
	;
	F_pfree(m, v7)
	mBase = m.M
	v227 = m.ExcPending
	if v227 != 0 {
		goto L1
	} else {
		goto L70
	}
L68:
	;
	goto L69
L69:
	;
	v228 = *(*int32)(unsafe.Add(mBase, uint32(v12)))
	if v228 != v14 {
		goto L71
	} else {
		goto L72
	}
L70:
	;
	goto L69
L71:
	;
	F_pfree(m, v14)
	mBase = m.M
	v231 = m.ExcPending
	if v231 != 0 {
		goto L1
	} else {
		goto L74
	}
L72:
	;
	goto L73
L73:
	;
	return base.B2i32(int32(0) < v223)
L74:
	;
	goto L73
}
func F_tsvector_setweight_by_filter(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v23 int32
	_ = v23
	var v25 int32
	_ = v25
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
	var v34 int32
	_ = v34
	var v35 int32
	_ = v35
	var v46 int32
	_ = v46
	var v51 int32
	_ = v51
	var v56 int32
	_ = v56
	var v58 int32
	_ = v58
	var v59 int32
	_ = v59
	var v62 int32
	_ = v62
	var v63 int32
	_ = v63
	var v64 int32
	_ = v64
	var v66 int32
	_ = v66
	var v67 int32
	_ = v67
	var v68 int32
	_ = v68
	var v77 int32
	_ = v77
	var v78 int32
	_ = v78
	var v82 int32
	_ = v82
	var v83 int32
	_ = v83
	var v93 int32
	_ = v93
	var v107 int32
	_ = v107
	var v108 int32
	_ = v108
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
	var v118 int32
	_ = v118
	var v122 int32
	_ = v122
	var v127 int32
	_ = v127
	var v135 int32
	_ = v135
	var v136 int32
	_ = v136
	var v151 int32
	_ = v151
	var v152 int32
	_ = v152
	var v153 int32
	_ = v153
	var v157 int32
	_ = v157
	var v161 int32
	_ = v161
	var v163 int32
	_ = v163
	var v168 int32
	_ = v168
	var v171 int32
	_ = v171
	var v173 int32
	_ = v173
	var v181 int32
	_ = v181
	var v182 int32
	_ = v182
	var v183 int32
	_ = v183
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
	var v194 int32
	_ = v194
	var v197 int32
	_ = v197
	var v198 int32
	_ = v198
	var v199 int32
	_ = v199
	var v204 int32
	_ = v204
	var v205 int32
	_ = v205
	var v206 int32
	_ = v206
	var v209 int32
	_ = v209
	var v210 int32
	_ = v210
	var v211 int32
	_ = v211
	var v214 int32
	_ = v214
	var v215 int32
	_ = v215
	var v217 int32
	_ = v217
	var v222 int32
	_ = v222
	var v235 int32
	_ = v235
	var v238 int32
	_ = v238
	var v241 int32
	_ = v241
	var v248 int32
	_ = v248
	var v249 int32
	_ = v249
	var v263 int32
	_ = v263
	var v264 int32
	_ = v264
	var v269 int32
	_ = v269
	var v271 int32
	_ = v271
	var v272 int32
	_ = v272
	var v276 int32
	_ = v276
	var v293 int32
	_ = v293
	var v294 int32
	_ = v294
	var v297 int32
	_ = v297
	var v299 int32
	_ = v299
	var v300 int32
	_ = v300
	var v302 int32
	_ = v302
	var v305 int32
	_ = v305
	var v310 int32
	_ = v310
	var v329 int32
	_ = v329
	var v334 int32
	_ = v334
	var v351 int32
	_ = v351
	var v352 int32
	_ = v352
	var v353 int32
	_ = v353
	var v355 int32
	_ = v355
	var v357 int32
	_ = v357
	var v358 int32
	_ = v358
	var v359 int32
	_ = v359
	var v362 int32
	_ = v362
	var v365 int32
	_ = v365
	var v366 int32
	_ = v366
	var v369 int32
	_ = v369
	var v372 int32
	_ = v372
	var v373 int32
	_ = v373
	var v376 int32
	_ = v376
	var v379 int32
	_ = v379
	var v403 int32
	_ = v403
	var v427 int32
	_ = v427
	var v430 int32
	_ = v430
	var v431 int32
	_ = v431
	var v434 int32
	_ = v434
	v23 = m.G0
	v25 = v23 - int32(16)
	m.G0 = v25
	v27 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v28 = F_pg_detoast_datum(m, v27)
	mBase = m.M
	v31 = m.ExcPending
	if v31 != 0 {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	return int32(0)
L2:
	;
	v32 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
	v33 = *(*int32)(unsafe.Add(mBase, uint32(l0)+36))
	v34 = F_pg_detoast_datum(m, v33)
	mBase = m.M
	v35 = m.ExcPending
	if v35 != 0 {
		goto L1
	} else {
		goto L3
	}
L3:
	;
	switch v32&int32(255) - int32(65) {
	case 0, 32:
		v58 = int32(49152)
		goto L4
	case 1, 33:
		goto L5
	case 2, 34:
		goto L8
	case 3, 35:
		goto L7
	default:
		goto L6
	}
L4:
	;
	v59 = *(*int32)(unsafe.Add(mBase, uint32(v28)))
	v62 = F_palloc(m, int32(base.Ui32(v59)>>(uint(int32(2))%32)))
	mBase = m.M
	v63 = m.ExcPending
	if v63 != 0 {
		goto L1
	} else {
		goto L12
	}
L5:
	;
	v58 = int32(32768)
	goto L4
L6:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v46 = m.ExcPending
	if v46 != 0 {
		goto L1
	} else {
		goto L9
	}
L7:
	;
	v58 = int32(0)
	goto L4
L8:
	;
	v58 = int32(16384)
	goto L4
L9:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v25))) = base.I32_extend8_s(v32)
	F_errmsg_internal(m, int32(497047), v25)
	mBase = m.M
	v51 = m.ExcPending
	if v51 != 0 {
		goto L1
	} else {
		goto L10
	}
L10:
	;
	F_errfinish(m, int32(490818), int32(308), int32(214022))
	mBase = m.M
	v56 = m.ExcPending
	if v56 != 0 {
		goto L1
	} else {
		goto L11
	}
L11:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L12:
	;
	v64 = *(*int32)(unsafe.Add(mBase, uint32(v28)))
	v66 = int32(base.Ui32(v64) >> (uint(int32(2)) % 32))
	if v66 != 0 {
		goto L14
	} else {
		goto L15
	}
L13:
	;
	F_deconstruct_array_builtin(m, v34, int32(25), v25+int32(8), v25+int32(4), v25+int32(12))
	mBase = m.M
	v77 = m.ExcPending
	if v77 != 0 {
		goto L1
	} else {
		goto L17
	}
L14:
	;
	v67 = F__emscripten_memcpy_bulkmem(m, v62, v28, v66)
	mBase = m.M
	v68 = v67
	goto L16
L15:
	;
	v68 = v62
	goto L16
L16:
	;
	goto L13
L17:
	;
	v78 = *(*int32)(unsafe.Add(mBase, uint32(v25)+12))
	if int32(0) < v78 {
		goto L18
	} else {
		goto L19
	}
L18:
	;
	v82 = v68 + int32(8)
	v83 = *(*int32)(unsafe.Add(mBase, uint32(v25)+4))
	v93 = int32(0)
	goto L21
L19:
	;
	goto L20
L20:
	;
	v427 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	if v427 != v28 {
		goto L81
	} else {
		goto L82
	}
L21:
	;
	v107 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v93+v83))))
	if v107 != 0 {
		goto L23
	} else {
		goto L24
	}
L22:
	;
	goto L20
L23:
	;
	v403 = v93 + int32(1)
	if v403 < v78 {
		v93 = v403
		goto L21
	} else {
		goto L80
	}
L24:
	;
	v108 = *(*int32)(unsafe.Add(mBase, uint32(v68)+4))
	if v108 <= int32(0) {
		goto L23
	} else {
		goto L25
	}
L25:
	;
	v111 = *(*int32)(unsafe.Add(mBase, uint32(v25)+8))
	v112 = int32(2)
	v115 = *(*int32)(unsafe.Add(mBase, uint32(v111+v93<<(uint(v112)%32))))
	v116 = int32(4)
	v117 = v115 + v116
	v118 = *(*int32)(unsafe.Add(mBase, uint32(v115)))
	v122 = int32(base.Ui32(v118)>>(uint(v112)%32)) - v116
	v127 = v82 + v108<<(uint(v112)%32)
	v135 = int32(0)
	v136 = v108
	goto L26
L26:
	;
	v151 = v135 + v136
	v152 = int32(2)
	v153 = base.I32_div_s(v151, v152)
	v157 = *(*int32)(unsafe.Add(mBase, uint32(v82+v153<<(uint(v152)%32))))
	v161 = int32(base.Ui32(v157)>>(uint(int32(1))%32)) & int32(2047)
	v163 = int32(base.Ui32(v157) >> (uint(int32(12)) % 32))
	if v122 == int32(0) {
		goto L33
	} else {
		goto L34
	}
L27:
	;
	if v151 < int32(-1) {
		goto L23
	} else {
		goto L67
	}
L28:
	;
	goto L27
L29:
	;
	if v248 < v249 {
		v135 = v248
		v136 = v249
		goto L26
	} else {
		goto L66
	}
L30:
	;
	v248 = v153 + int32(1)
	v249 = v136
	goto L29
L31:
	;
	if v241 == int32(0) {
		goto L28
	} else {
		goto L65
	}
L32:
	;
	if int32(0) <= v238 {
		v241 = v238
		goto L31
	} else {
		goto L64
	}
L33:
	;
	if v161 != 0 {
		goto L36
	} else {
		goto L37
	}
L34:
	;
	goto L35
L35:
	;
	if v161 == int32(0) {
		v241 = base.B2i32(base.Ui32(int32(19)) < base.Ui32(v118))
		goto L31
	} else {
		goto L39
	}
L36:
	;
	v168 = int32(-1)
	goto L38
L37:
	;
	v168 = int32(0)
	goto L38
L38:
	;
	v238 = v168
	goto L32
L39:
	;
	v171 = v127 + v163
	if base.Ui32(v122) < base.Ui32(v161) {
		goto L40
	} else {
		goto L41
	}
L40:
	;
	v173 = v122
	goto L42
L41:
	;
	v173 = v161
	goto L42
L42:
	;
	if base.Ui32(int32(4)) <= base.Ui32(v173) {
		goto L46
	} else {
		goto L47
	}
L43:
	;
	if v235 != 0 {
		v238 = v235
		goto L32
	} else {
		goto L61
	}
L44:
	;
	v235 = int32(0)
	goto L43
L45:
	;
	v209 = v204
	v210 = v205
	v211 = v206
	goto L55
L46:
	;
	if (v117|v171)&int32(3) != 0 {
		v204 = v117
		v205 = v171
		v206 = v173
		goto L45
	} else {
		goto L49
	}
L47:
	;
	v197 = v117
	v198 = v171
	v199 = v173
	goto L48
L48:
	;
	if v199 == int32(0) {
		goto L44
	} else {
		goto L54
	}
L49:
	;
	v181 = v117
	v182 = v171
	v183 = v173
	goto L50
L50:
	;
	v186 = *(*int32)(unsafe.Add(mBase, uint32(v181)))
	v187 = *(*int32)(unsafe.Add(mBase, uint32(v182)))
	if v186 != v187 {
		v204 = v181
		v205 = v182
		v206 = v183
		goto L45
	} else {
		goto L52
	}
L51:
	;
	v197 = v192
	v198 = v190
	v199 = v194
	goto L48
L52:
	;
	v189 = int32(4)
	v190 = v182 + v189
	v192 = v181 + v189
	v194 = v183 - v189
	if base.Ui32(int32(3)) < base.Ui32(v194) {
		v181 = v192
		v182 = v190
		v183 = v194
		goto L50
	} else {
		goto L53
	}
L53:
	;
	goto L51
L54:
	;
	v204 = v197
	v205 = v198
	v206 = v199
	goto L45
L55:
	;
	v214 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v209))))
	v215 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v210))))
	if v214 == v215 {
		goto L57
	} else {
		goto L58
	}
L56:
	;
	v235 = v214 - v215
	goto L43
L57:
	;
	v217 = int32(1)
	v222 = v211 - v217
	if v222 != 0 {
		v209 = v209 + v217
		v210 = v210 + v217
		v211 = v222
		goto L55
	} else {
		goto L60
	}
L58:
	;
	goto L59
L59:
	;
	goto L56
L60:
	;
	goto L44
L61:
	;
	if v161 == v122 {
		goto L28
	} else {
		goto L62
	}
L62:
	;
	if v161 <= v122 {
		goto L30
	} else {
		goto L63
	}
L63:
	;
	v248 = v135
	v249 = v153
	goto L29
L64:
	;
	v248 = v135
	v249 = v153
	goto L29
L65:
	;
	goto L30
L66:
	;
	goto L23
L67:
	;
	if v157&int32(1) == int32(0) {
		goto L23
	} else {
		goto L68
	}
L68:
	;
	v263 = v127 + (v161+v163+int32(1))&int32(4194302)
	v264 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v263))))
	if v264 == int32(0) {
		goto L23
	} else {
		goto L69
	}
L69:
	;
	v269 = v264 & int32(3)
	if v269 != 0 {
		goto L70
	} else {
		goto L71
	}
L70:
	;
	v271 = v263
	v272 = int32(0)
	v276 = v264
	goto L73
L71:
	;
	v305 = v263
	v310 = v264
	goto L72
L72:
	;
	if base.Ui32(v264) < base.Ui32(int32(4)) {
		goto L23
	} else {
		goto L76
	}
L73:
	;
	v293 = v271 + int32(2)
	v294 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v293))))
	v297 = v294&int32(16383) | v58
	*(*uint16)(unsafe.Add(mBase, uint32(v293))) = uint16(v297)
	v299 = int32(1)
	v300 = v276 - v299
	v302 = v272 + v299
	if v302 != v269 {
		v271 = v293
		v272 = v302
		v276 = v300
		goto L73
	} else {
		goto L75
	}
L74:
	;
	v305 = v293
	v310 = v300
	goto L72
L75:
	;
	goto L74
L76:
	;
	v329 = v305
	v334 = v310
	goto L77
L77:
	;
	v351 = v329 + int32(2)
	v352 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v351))))
	v353 = int32(16383)
	v355 = v352&v353 | v58
	*(*uint16)(unsafe.Add(mBase, uint32(v351))) = uint16(v355)
	v357 = int32(4)
	v358 = v329 + v357
	v359 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v358))))
	v362 = v359&v353 | v58
	*(*uint16)(unsafe.Add(mBase, uint32(v358))) = uint16(v362)
	v365 = v329 + int32(6)
	v366 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v365))))
	v369 = v366&v353 | v58
	*(*uint16)(unsafe.Add(mBase, uint32(v365))) = uint16(v369)
	v372 = v329 + int32(8)
	v373 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v372))))
	v376 = v373&v353 | v58
	*(*uint16)(unsafe.Add(mBase, uint32(v372))) = uint16(v376)
	v379 = v334 - v357
	if v379 != 0 {
		v329 = v372
		v334 = v379
		goto L77
	} else {
		goto L79
	}
L78:
	;
	goto L23
L79:
	;
	goto L78
L80:
	;
	goto L22
L81:
	;
	F_pfree(m, v28)
	mBase = m.M
	v430 = m.ExcPending
	if v430 != 0 {
		goto L1
	} else {
		goto L84
	}
L82:
	;
	goto L83
L83:
	;
	v431 = *(*int32)(unsafe.Add(mBase, uint32(l0)+36))
	if v431 != v34 {
		goto L85
	} else {
		goto L86
	}
L84:
	;
	goto L83
L85:
	;
	F_pfree(m, v34)
	mBase = m.M
	v434 = m.ExcPending
	if v434 != 0 {
		goto L1
	} else {
		goto L88
	}
L86:
	;
	goto L87
L87:
	;
	m.G0 = v25 + int32(16)
	return v68
L88:
	;
	goto L87
}
func F_tsvector_update_trigger_byid(m *base.Module, l0 int32) int32 {
	var v3 int32
	_ = v3
	var v6 int32
	_ = v6
	v3 = F_tsvector_update_trigger(m, l0, int32(0))
	v6 = m.ExcPending
	if v6 != 0 {
		return int32(0)
	} else {
		return v3
	}
}
