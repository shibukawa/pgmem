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
	var v22 int32
	_ = v22
	var v24 int32
	_ = v24
	var v26 int32
	_ = v26
	var v34 int32
	_ = v34
	var v36 int32
	_ = v36
	var v41 int32
	_ = v41
	var v44 int32
	_ = v44
	var v49 int32
	_ = v49
	var v50 int32
	_ = v50
	var v51 float64
	_ = v51
	var v58 float64
	_ = v58
	var v61 int32
	_ = v61
	var v65 int32
	_ = v65
	var v68 int32
	_ = v68
	var v70 int32
	_ = v70
	var v78 int32
	_ = v78
	var v81 int64
	_ = v81
	var v82 int32
	_ = v82
	var v83 int32
	_ = v83
	var v88 int32
	_ = v88
	var v89 int32
	_ = v89
	var v93 int32
	_ = v93
	var v95 int32
	_ = v95
	var v98 int32
	_ = v98
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
	var v137 int32
	_ = v137
	var v139 int32
	_ = v139
	var v142 int32
	_ = v142
	var v147 int32
	_ = v147
	var v152 int32
	_ = v152
	var v153 int32
	_ = v153
	var v163 int32
	_ = v163
	var v164 int32
	_ = v164
	var v165 int32
	_ = v165
	var v168 int32
	_ = v168
	var v172 int32
	_ = v172
	var v177 int32
	_ = v177
	var v178 int32
	_ = v178
	var v179 int32
	_ = v179
	var v181 int32
	_ = v181
	var v184 int32
	_ = v184
	var v189 int32
	_ = v189
	var v190 int32
	_ = v190
	var v194 int32
	_ = v194
	var v196 int32
	_ = v196
	var v197 int32
	_ = v197
	var v198 int32
	_ = v198
	var v203 int32
	_ = v203
	var v220 int32
	_ = v220
	var v221 int32
	_ = v221
	var v224 int32
	_ = v224
	var v227 int32
	_ = v227
	var v228 int32
	_ = v228
	var v232 int32
	_ = v232
	var v235 int32
	_ = v235
	var v236 int32
	_ = v236
	var v266 int32
	_ = v266
	var v284 int32
	_ = v284
	var v285 int32
	_ = v285
	var v293 int32
	_ = v293
	var v300 int32
	_ = v300
	var v309 float64
	_ = v309
	var v313 int32
	_ = v313
	var v317 float64
	_ = v317
	var v320 int32
	_ = v320
	var v324 int32
	_ = v324
	var v327 int32
	_ = v327
	var v336 int32
	_ = v336
	var v341 float64
	_ = v341
	var v348 int32
	_ = v348
	var v351 int32
	_ = v351
	var v360 int32
	_ = v360
	var v365 float32
	_ = v365
	var v368 float64
	_ = v368
	var v378 int32
	_ = v378
	var v379 int64
	_ = v379
	var v380 int64
	_ = v380
	var v383 int64
	_ = v383
	var v384 int64
	_ = v384
	var v385 int64
	_ = v385
	var v386 int64
	_ = v386
	var v387 int64
	_ = v387
	var v388 int64
	_ = v388
	var v389 int64
	_ = v389
	var v390 int64
	_ = v390
	var v391 int64
	_ = v391
	var v392 int64
	_ = v392
	var v393 int64
	_ = v393
	var v394 int64
	_ = v394
	var v395 int64
	_ = v395
	var v396 int64
	_ = v396
	var v397 int64
	_ = v397
	var v398 int64
	_ = v398
	var v399 int64
	_ = v399
	var v400 int64
	_ = v400
	var v401 int64
	_ = v401
	var v402 int64
	_ = v402
	var v403 int64
	_ = v403
	var v404 int64
	_ = v404
	var v405 int64
	_ = v405
	var v406 int64
	_ = v406
	var v407 int64
	_ = v407
	var v408 int64
	_ = v408
	var v409 int64
	_ = v409
	var v410 int64
	_ = v410
	var v411 int64
	_ = v411
	var v412 int64
	_ = v412
	var v413 int64
	_ = v413
	var v445 int64
	_ = v445
	var v446 int32
	_ = v446
	var v447 int32
	_ = v447
	var v448 int32
	_ = v448
	var v450 int32
	_ = v450
	var v452 int32
	_ = v452
	var v455 int32
	_ = v455
	var v456 int32
	_ = v456
	var v458 int32
	_ = v458
	var v459 int32
	_ = v459
	var v461 int32
	_ = v461
	var v464 int32
	_ = v464
	var v466 int32
	_ = v466
	var v481 int32
	_ = v481
	var v489 int32
	_ = v489
	var v491 int32
	_ = v491
	var v492 int32
	_ = v492
	var v494 int32
	_ = v494
	var v497 int32
	_ = v497
	var v498 int32
	_ = v498
	var v500 int32
	_ = v500
	var v505 int32
	_ = v505
	var v522 int32
	_ = v522
	var v523 int32
	_ = v523
	var v531 int32
	_ = v531
	var v536 int32
	_ = v536
	var v542 int32
	_ = v542
	var v548 int32
	_ = v548
	var v549 int32
	_ = v549
	var v551 int32
	_ = v551
	var v552 int32
	_ = v552
	var v553 int32
	_ = v553
	var v554 int32
	_ = v554
	var v555 int32
	_ = v555
	var v556 int32
	_ = v556
	var v563 int32
	_ = v563
	var v564 int32
	_ = v564
	var v565 int32
	_ = v565
	var v567 int32
	_ = v567
	var v570 int32
	_ = v570
	var v571 int32
	_ = v571
	var v574 int32
	_ = v574
	var v575 int32
	_ = v575
	var v576 int32
	_ = v576
	var v581 int32
	_ = v581
	var v602 int32
	_ = v602
	var v604 int32
	_ = v604
	var v605 int32
	_ = v605
	var v606 int32
	_ = v606
	var v607 int32
	_ = v607
	var v608 int32
	_ = v608
	var v612 int32
	_ = v612
	var v618 int32
	_ = v618
	var v643 int32
	_ = v643
	var v659 int32
	_ = v659
	var v663 int32
	_ = v663
	var v665 int32
	_ = v665
	var v667 int32
	_ = v667
	var v676 int32
	_ = v676
	var v705 int32
	_ = v705
	var v709 int32
	_ = v709
	var v714 int32
	_ = v714
	v5 = int32(0)
	v22 = m.G0
	v24 = v22 - int32(112)
	m.G0 = v24
	v26 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	*(*int32)(unsafe.Add(mBase, uint32(v24)+60)) = int32(1281)
	*(*int32)(unsafe.Add(mBase, uint32(v24)+56)) = int32(1282)
	*(*int64)(unsafe.Add(mBase, uint32(v24)+48)) = int64(68719476744)
	v34 = *(*int32)(unsafe.Add(mBase, _c_F_compute_tsvector_stats[0]))
	*(*int32)(unsafe.Add(mBase, uint32(v24)+76)) = v34
	v36 = int32(_a_F_compute_tsvector_stats_0)
	v41 = base.I32_div_s(v26*v36+v36, int32(7))
	v44 = v26 * int32(10)
	v49 = F_hash_create(m, int32(_a_F_compute_tsvector_stats_1), base.I64_extend_i32_s(v44), v24+int32(40), int32(1224))
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
	v51 = float64(0)
	if int32(0) < l2 {
		goto L4
	} else {
		goto L5
	}
L3:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v705 = m.ExcPending
	if v705 != 0 {
		goto L1
	} else {
		goto L110
	}
L4:
	;
	v58 = v51
	v61 = int32(1)
	v65 = v5
	v68 = v5
	v70 = v5
	goto L7
L5:
	;
	v341 = v51
	v348 = v5
	v351 = v5
	goto L6
L6:
	;
	if v348 < l2 {
		goto L62
	} else {
		goto L63
	}
L7:
	;
	F_vacuum_delay_point(m, int32(1))
	mBase = m.M
	v78 = m.ExcPending
	if v78 != 0 {
		goto L1
	} else {
		goto L9
	}
L8:
	;
	v341 = v317
	v348 = v324
	v351 = v327
	goto L6
L9:
	;
	v81 = m.T0[l1].(func(*base.Module, int32, int32, int32) int64)(m, l0, v70, v24+int32(31))
	mBase = m.M
	v82 = m.ExcPending
	if v82 != 0 {
		goto L1
	} else {
		goto L10
	}
L10:
	;
	v83 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v24)+31)))
	if v83 == int32(1) {
		goto L12
	} else {
		goto L13
	}
L11:
	;
	v336 = v70 + int32(1)
	if v336 != l2 {
		v58 = v317
		v61 = v320
		v65 = v324
		v68 = v327
		v70 = v336
		goto L7
	} else {
		goto L60
	}
L12:
	;
	v317 = v58
	v320 = v61
	v324 = v65 + int32(1)
	v327 = v68
	goto L11
L13:
	;
	goto L14
L14:
	;
	v88 = base.I32_wrap_i64(v81)
	v89 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v88))))
	if v89 == int32(1) {
		goto L16
	} else {
		goto L17
	}
L15:
	;
	v115 = F_pg_detoast_datum(m, v88)
	mBase = m.M
	v116 = m.ExcPending
	if v116 != 0 {
		goto L1
	} else {
		goto L26
	}
L16:
	;
	v93 = int32(18)
	v95 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v88)+1)))
	if v95 == v93 {
		goto L19
	} else {
		goto L20
	}
L17:
	;
	goto L18
L18:
	;
	v106 = int32(1)
	if v89&v106 != 0 {
		v114 = int32(base.Ui32(v89) >> (uint(v106) % 32))
		goto L15
	} else {
		goto L25
	}
L19:
	;
	v98 = v93
	goto L21
L20:
	;
	v98 = int32(2)
	goto L21
L21:
	;
	if base.Ui32((v95-int32(1))&int32(255)) < base.Ui32(int32(3)) {
		goto L22
	} else {
		goto L23
	}
L22:
	;
	v105 = int32(6)
	goto L24
L23:
	;
	v105 = v98
	goto L24
L24:
	;
	v114 = v105
	goto L15
L25:
	;
	v110 = *(*int32)(unsafe.Add(mBase, uint32(v88)))
	v114 = int32(base.Ui32(v110) >> (uint(int32(2)) % 32))
	goto L15
L26:
	;
	v117 = *(*int32)(unsafe.Add(mBase, uint32(v115)+4))
	if int32(0) < v117 {
		goto L27
	} else {
		goto L28
	}
L27:
	;
	v121 = v115 + int32(8)
	v132 = v61
	v137 = v121
	v139 = v68
	v142 = int32(0)
	goto L30
L28:
	;
	v293 = v61
	v300 = v68
	goto L29
L29:
	;
	v309 = base.F64_add(v58, base.F64_convert_i32_u(v114))
	if v81 == base.I64_extend_i32_u(v115) {
		v317 = v309
		v320 = v293
		v324 = v65
		v327 = v300
		goto L11
	} else {
		goto L58
	}
L30:
	;
	v147 = *(*int32)(unsafe.Add(mBase, uint32(v137)))
	*(*int32)(unsafe.Add(mBase, uint32(v24)+32)) = v121 + v117<<(uint(int32(2))%32) + int32(base.Ui32(v147)>>(uint(int32(12))%32))
	v152 = *(*int32)(unsafe.Add(mBase, uint32(v137)))
	v153 = int32(1)
	*(*int32)(unsafe.Add(mBase, uint32(v24)+36)) = int32(base.Ui32(v152)>>(uint(v153)%32)) & int32(2047)
	v163 = F_hash_search(m, v49, v24+int32(32), v153, v24+int32(30))
	mBase = m.M
	v164 = m.ExcPending
	if v164 != 0 {
		goto L1
	} else {
		goto L32
	}
L31:
	;
	v293 = v266
	v300 = v189
	goto L29
L32:
	;
	v165 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v24)+30)))
	if v165 == int32(1) {
		goto L34
	} else {
		goto L35
	}
L33:
	;
	v189 = v139 + int32(1)
	v190 = base.I32_rem_s(v189, v41)
	if v190 == int32(0) {
		goto L39
	} else {
		goto L40
	}
L34:
	;
	v168 = *(*int32)(unsafe.Add(mBase, uint32(v163)+8))
	*(*int32)(unsafe.Add(mBase, uint32(v163)+8)) = v168 + int32(1)
	goto L33
L35:
	;
	goto L36
L36:
	;
	v172 = int32(1)
	*(*int32)(unsafe.Add(mBase, uint32(v163)+8)) = v172
	*(*int32)(unsafe.Add(mBase, uint32(v163)+12)) = v132 - v172
	v177 = *(*int32)(unsafe.Add(mBase, uint32(v24)+36))
	v178 = F_palloc(m, v177)
	mBase = m.M
	v179 = m.ExcPending
	if v179 != 0 {
		goto L1
	} else {
		goto L37
	}
L37:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v163))) = v178
	v181 = *(*int32)(unsafe.Add(mBase, uint32(v24)+36))
	if v181 == int32(0) {
		goto L33
	} else {
		goto L38
	}
L38:
	;
	v184 = *(*int32)(unsafe.Add(mBase, uint32(v24)+32))
	base.MemoryCopy(m, v178, v184, v181)
	goto L33
L39:
	;
	v194 = v24 + int32(92)
	F_hash_seq_init(m, v194, v49)
	mBase = m.M
	v196 = m.ExcPending
	if v196 != 0 {
		goto L1
	} else {
		goto L42
	}
L40:
	;
	v266 = v132
	goto L41
L41:
	;
	v284 = v142 + int32(1)
	v285 = *(*int32)(unsafe.Add(mBase, uint32(v115)+4))
	if v284 < v285 {
		v132 = v266
		v137 = v137 + int32(4)
		v139 = v189
		v142 = v284
		goto L30
	} else {
		goto L57
	}
L42:
	;
	v197 = F_hash_seq_search(m, v194)
	mBase = m.M
	v198 = m.ExcPending
	if v198 != 0 {
		goto L1
	} else {
		goto L43
	}
L43:
	;
	if v197 != 0 {
		goto L44
	} else {
		goto L45
	}
L44:
	;
	v203 = v197
	goto L47
L45:
	;
	goto L46
L46:
	;
	v266 = v132 + int32(1)
	goto L41
L47:
	;
	v220 = *(*int32)(unsafe.Add(mBase, uint32(v203)+12))
	v221 = *(*int32)(unsafe.Add(mBase, uint32(v203)+8))
	if v220+v221 <= v132 {
		goto L49
	} else {
		goto L50
	}
L48:
	;
	goto L46
L49:
	;
	v224 = *(*int32)(unsafe.Add(mBase, uint32(v203)))
	v227 = F_hash_search(m, v49, v203, int32(2), int32(0))
	mBase = m.M
	v228 = m.ExcPending
	if v228 != 0 {
		goto L1
	} else {
		goto L52
	}
L50:
	;
	goto L51
L51:
	;
	v235 = F_hash_seq_search(m, v24+int32(92))
	mBase = m.M
	v236 = m.ExcPending
	if v236 != 0 {
		goto L1
	} else {
		goto L55
	}
L52:
	;
	if v227 == int32(0) {
		goto L3
	} else {
		goto L53
	}
L53:
	;
	F_pfree(m, v224)
	mBase = m.M
	v232 = m.ExcPending
	if v232 != 0 {
		goto L1
	} else {
		goto L54
	}
L54:
	;
	goto L51
L55:
	;
	if v235 != 0 {
		v203 = v235
		goto L47
	} else {
		goto L56
	}
L56:
	;
	goto L48
L57:
	;
	goto L31
L58:
	;
	F_pfree(m, v115)
	mBase = m.M
	v313 = m.ExcPending
	if v313 != 0 {
		goto L1
	} else {
		goto L59
	}
L59:
	;
	v317 = v309
	v320 = v293
	v324 = v65
	v327 = v300
	goto L11
L60:
	;
	goto L8
L61:
	;
	m.G0 = v24 + int32(112)
	return
L62:
	;
	v360 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(l0)+36)) = uint8(v360)
	v365 = base.F32_demote_f64(base.F64_div(base.F64_convert_i32_s(v348), base.F64_convert_i32_s(l2)))
	*(*float32)(unsafe.Add(mBase, uint32(l0)+40)) = v365
	v368 = base.F64_convert_i32_s(l2 - v348)
	*(*int32)(unsafe.Add(mBase, uint32(l0)+44)) = base.I32_trunc_sat_f64_s(base.F64_div(v341, v368))
	*(*float32)(unsafe.Add(mBase, uint32(l0)+48)) = base.F32_neg(base.F32_sub(float32(1), v365))
	v378 = *(*int32)(unsafe.Add(mBase, uint32(v49)))
	v379 = *(*int64)(unsafe.Add(mBase, uint32(v378)+8))
	v380 = *(*int64)(unsafe.Add(mBase, uint32(v378)+808))
	if v380 != int64(0) {
		goto L66
	} else {
		goto L67
	}
L63:
	;
	goto L64
L64:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+48)) = int32(0)
	*(*int64)(unsafe.Add(mBase, uint32(l0)+40)) = int64(1065353216)
	v676 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(l0)+36)) = uint8(v676)
	goto L61
L65:
	;
	v446 = base.I32_wrap_i64(v445)
	v447 = F_palloc_mul(m, int32(4), v446)
	mBase = m.M
	v448 = m.ExcPending
	if v448 != 0 {
		goto L1
	} else {
		goto L69
	}
L66:
	;
	v383 = *(*int64)(unsafe.Add(mBase, uint32(v378)+752))
	v384 = *(*int64)(unsafe.Add(mBase, uint32(v378)+728))
	v385 = *(*int64)(unsafe.Add(mBase, uint32(v378)+704))
	v386 = *(*int64)(unsafe.Add(mBase, uint32(v378)+680))
	v387 = *(*int64)(unsafe.Add(mBase, uint32(v378)+656))
	v388 = *(*int64)(unsafe.Add(mBase, uint32(v378)+632))
	v389 = *(*int64)(unsafe.Add(mBase, uint32(v378)+608))
	v390 = *(*int64)(unsafe.Add(mBase, uint32(v378)+584))
	v391 = *(*int64)(unsafe.Add(mBase, uint32(v378)+560))
	v392 = *(*int64)(unsafe.Add(mBase, uint32(v378)+536))
	v393 = *(*int64)(unsafe.Add(mBase, uint32(v378)+512))
	v394 = *(*int64)(unsafe.Add(mBase, uint32(v378)+488))
	v395 = *(*int64)(unsafe.Add(mBase, uint32(v378)+464))
	v396 = *(*int64)(unsafe.Add(mBase, uint32(v378)+440))
	v397 = *(*int64)(unsafe.Add(mBase, uint32(v378)+416))
	v398 = *(*int64)(unsafe.Add(mBase, uint32(v378)+392))
	v399 = *(*int64)(unsafe.Add(mBase, uint32(v378)+368))
	v400 = *(*int64)(unsafe.Add(mBase, uint32(v378)+344))
	v401 = *(*int64)(unsafe.Add(mBase, uint32(v378)+320))
	v402 = *(*int64)(unsafe.Add(mBase, uint32(v378)+296))
	v403 = *(*int64)(unsafe.Add(mBase, uint32(v378)+272))
	v404 = *(*int64)(unsafe.Add(mBase, uint32(v378)+248))
	v405 = *(*int64)(unsafe.Add(mBase, uint32(v378)+224))
	v406 = *(*int64)(unsafe.Add(mBase, uint32(v378)+200))
	v407 = *(*int64)(unsafe.Add(mBase, uint32(v378)+176))
	v408 = *(*int64)(unsafe.Add(mBase, uint32(v378)+152))
	v409 = *(*int64)(unsafe.Add(mBase, uint32(v378)+128))
	v410 = *(*int64)(unsafe.Add(mBase, uint32(v378)+104))
	v411 = *(*int64)(unsafe.Add(mBase, uint32(v378)+80))
	v412 = *(*int64)(unsafe.Add(mBase, uint32(v378)+56))
	v413 = *(*int64)(unsafe.Add(mBase, uint32(v378)+32))
	v445 = v383 + (v384 + (v385 + (v386 + (v387 + (v388 + (v389 + (v390 + (v391 + (v392 + (v393 + (v394 + (v395 + (v396 + (v397 + (v398 + (v399 + (v400 + (v401 + (v402 + (v403 + (v404 + (v405 + (v406 + (v407 + (v408 + (v409 + (v410 + (v411 + (v412 + (v413 + v379))))))))))))))))))))))))))))))
	goto L68
L67:
	;
	v445 = v379
	goto L68
L68:
	;
	goto L65
L69:
	;
	v450 = v24 + int32(92)
	F_hash_seq_init(m, v450, v49)
	mBase = m.M
	v452 = m.ExcPending
	if v452 != 0 {
		goto L1
	} else {
		goto L70
	}
L70:
	;
	v455 = base.I32_div_s(v351*int32(9), v41)
	v456 = int32(0)
	v458 = F_hash_seq_search(m, v450)
	mBase = m.M
	v459 = m.ExcPending
	if v459 != 0 {
		goto L1
	} else {
		goto L71
	}
L71:
	;
	if v458 != 0 {
		goto L72
	} else {
		goto L73
	}
L72:
	;
	v461 = v456
	v464 = v458
	v466 = v456
	goto L75
L73:
	;
	v500 = v456
	v505 = v456
	goto L74
L74:
	;
	v522 = F_errstart(m, int32(12), int32(0))
	mBase = m.M
	v523 = m.ExcPending
	if v523 != 0 {
		goto L1
	} else {
		goto L85
	}
L75:
	;
	v481 = *(*int32)(unsafe.Add(mBase, uint32(v464)+8))
	if v455 < v481 {
		goto L77
	} else {
		goto L78
	}
L76:
	;
	v500 = v492
	v505 = v494
	goto L74
L77:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v447+v466<<(uint(int32(2))%32)))) = v464
	v489 = *(*int32)(unsafe.Add(mBase, uint32(v464)+8))
	if v489 < v461 {
		goto L80
	} else {
		goto L81
	}
L78:
	;
	v492 = v461
	v494 = v466
	goto L79
L79:
	;
	v497 = F_hash_seq_search(m, v24+int32(92))
	mBase = m.M
	v498 = m.ExcPending
	if v498 != 0 {
		goto L1
	} else {
		goto L83
	}
L80:
	;
	v491 = v461
	goto L82
L81:
	;
	v491 = v489
	goto L82
L82:
	;
	v492 = v491
	v494 = v466 + int32(1)
	goto L79
L83:
	;
	if v497 != 0 {
		v461 = v492
		v464 = v497
		v466 = v494
		goto L75
	} else {
		goto L84
	}
L84:
	;
	goto L76
L85:
	;
	if v522 != 0 {
		goto L86
	} else {
		goto L87
	}
L86:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v24)+16)) = v505
	*(*int32)(unsafe.Add(mBase, uint32(v24)+12)) = v446
	*(*int32)(unsafe.Add(mBase, uint32(v24)+8)) = v351
	*(*int32)(unsafe.Add(mBase, uint32(v24)+4)) = v41
	*(*int32)(unsafe.Add(mBase, uint32(v24))) = v44
	F_errmsg_internal(m, int32(_a_F_compute_tsvector_stats_2), v24)
	mBase = m.M
	v531 = m.ExcPending
	if v531 != 0 {
		goto L1
	} else {
		goto L89
	}
L87:
	;
	goto L88
L88:
	;
	if v44 < v505 {
		goto L92
	} else {
		goto L93
	}
L89:
	;
	F_errfinish(m, int32(_a_F_compute_tsvector_stats_3), int32(341), int32(_a_F_compute_tsvector_stats_4))
	mBase = m.M
	v536 = m.ExcPending
	if v536 != 0 {
		goto L1
	} else {
		goto L90
	}
L90:
	;
	goto L88
L91:
	;
	v556 = int32(0)
	if v554 < v556 {
		goto L61
	} else {
		goto L99
	}
L92:
	;
	F_qsort_interruptible(m, v447, v505, int32(4), int32(1283), int32(0))
	mBase = m.M
	v542 = m.ExcPending
	if v542 != 0 {
		goto L1
	} else {
		goto L95
	}
L93:
	;
	goto L94
L94:
	;
	v551 = v455 + int32(1)
	if v505 != 0 {
		goto L96
	} else {
		goto L97
	}
L95:
	;
	v548 = *(*int32)(unsafe.Add(mBase, uint32(v447+v44<<(uint(int32(2))%32)-int32(4))))
	v549 = *(*int32)(unsafe.Add(mBase, uint32(v548)+8))
	v553 = v500
	v554 = v44
	v555 = v549
	goto L91
L96:
	;
	v552 = v500
	goto L98
L97:
	;
	v552 = v551
	goto L98
L98:
	;
	v553 = v552
	v554 = v505
	v555 = v551
	goto L91
L99:
	;
	F_qsort_interruptible(m, v447, v554, int32(4), int32(1284), int32(0))
	mBase = m.M
	v563 = m.ExcPending
	if v563 != 0 {
		goto L1
	} else {
		goto L100
	}
L100:
	;
	v564 = int32(_a_F_compute_tsvector_stats_5)
	v565 = *(*int32)(unsafe.Add(mBase, _c_F_compute_tsvector_stats[0]))
	v567 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	*(*int32)(unsafe.Add(mBase, _c_F_compute_tsvector_stats[0])) = v567
	v570 = F_palloc_mul(m, int32(8), v554)
	mBase = m.M
	v571 = m.ExcPending
	if v571 != 0 {
		goto L1
	} else {
		goto L101
	}
L101:
	;
	v574 = v554 + int32(2)
	v575 = F_palloc_mul(m, int32(4), v574)
	mBase = m.M
	v576 = m.ExcPending
	if v576 != 0 {
		goto L1
	} else {
		goto L102
	}
L102:
	;
	if v554 != 0 {
		goto L103
	} else {
		goto L104
	}
L103:
	;
	v581 = v556
	goto L106
L104:
	;
	goto L105
L105:
	;
	v643 = v575 + v554<<(uint(int32(2))%32)
	*(*float32)(unsafe.Add(mBase, uint32(v643)+4)) = base.F32_demote_f64(base.F64_div(base.F64_convert_i32_s(v553), v368))
	*(*float32)(unsafe.Add(mBase, uint32(v643))) = base.F32_demote_f64(base.F64_div(base.F64_convert_i32_s(v555), v368))
	*(*int32)(unsafe.Add(mBase, _c_F_compute_tsvector_stats[0])) = v565
	*(*int32)(unsafe.Add(mBase, uint32(l0)+124)) = v575
	*(*int32)(unsafe.Add(mBase, uint32(l0)+84)) = int32(100)
	*(*int32)(unsafe.Add(mBase, uint32(l0)+64)) = int32(98)
	v659 = int32(4)
	*(*uint16)(unsafe.Add(mBase, uint32(l0)+52)) = uint16(v659)
	*(*int32)(unsafe.Add(mBase, uint32(l0)+164)) = v570
	*(*int32)(unsafe.Add(mBase, uint32(l0)+104)) = v574
	v663 = int32(105)
	*(*uint8)(unsafe.Add(mBase, uint32(l0)+219)) = uint8(v663)
	v665 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(l0)+214)) = uint8(v665)
	v667 = int32(_a_F_compute_tsvector_stats_6)
	*(*uint16)(unsafe.Add(mBase, uint32(l0)+204)) = uint16(v667)
	*(*int32)(unsafe.Add(mBase, uint32(l0)+184)) = int32(25)
	*(*int32)(unsafe.Add(mBase, uint32(l0)+144)) = v554
	goto L61
L106:
	;
	v602 = v581 << (uint(int32(2)) % 32)
	v604 = *(*int32)(unsafe.Add(mBase, uint32(v447+v602)))
	v605 = *(*int32)(unsafe.Add(mBase, uint32(v604)))
	v606 = *(*int32)(unsafe.Add(mBase, uint32(v604)+4))
	v607 = F_cstring_to_text_with_len(m, v605, v606)
	mBase = m.M
	v608 = m.ExcPending
	if v608 != 0 {
		goto L1
	} else {
		goto L108
	}
L107:
	;
	goto L105
L108:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v570+v581<<(uint(int32(3))%32)))) = base.I64_extend_i32_u(v607)
	v612 = *(*int32)(unsafe.Add(mBase, uint32(v604)+8))
	*(*float32)(unsafe.Add(mBase, uint32(v575+v602))) = base.F32_demote_f64(base.F64_div(base.F64_convert_i32_s(v612), v368))
	v618 = v581 + int32(1)
	if v618 != v554 {
		v581 = v618
		goto L106
	} else {
		goto L109
	}
L109:
	;
	goto L107
L110:
	;
	F_errmsg_internal(m, int32(_a_F_compute_tsvector_stats_7), int32(0))
	mBase = m.M
	v709 = m.ExcPending
	if v709 != 0 {
		goto L1
	} else {
		goto L111
	}
L111:
	;
	F_errfinish(m, int32(_a_F_compute_tsvector_stats_3), int32(484), int32(_a_F_compute_tsvector_stats_8))
	mBase = m.M
	v714 = m.ExcPending
	if v714 != 0 {
		goto L1
	} else {
		goto L112
	}
L112:
	;
	base.Wasm_trap_unreachable()
	for {
	}
}
func F_tsvector_delete_arr(m *base.Module, l0 int32) int64 {
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
	var v61 int32
	_ = v61
	var v63 int32
	_ = v63
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
	var v82 int32
	_ = v82
	var v83 int32
	_ = v83
	var v84 int32
	_ = v84
	var v85 int32
	_ = v85
	var v86 int32
	_ = v86
	var v89 int32
	_ = v89
	var v99 int32
	_ = v99
	var v106 int32
	_ = v106
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
	var v237 int32
	_ = v237
	var v239 int32
	_ = v239
	var v247 int32
	_ = v247
	var v260 int32
	_ = v260
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
	v25 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
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
	return int64(0)
L2:
	;
	v30 = *(*int32)(unsafe.Add(mBase, uint32(l0)+40))
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
	if int32(0) < v47 {
		goto L6
	} else {
		goto L7
	}
L6:
	;
	v51 = v26 + int32(8)
	v61 = v2
	v63 = v2
	v65 = v47
	goto L9
L7:
	;
	v260 = v2
	goto L8
L8:
	;
	v269 = F_tsvector_delete_by_indices(m, v26, v45, v260)
	mBase = m.M
	v270 = m.ExcPending
	if v270 != 0 {
		goto L1
	} else {
		goto L57
	}
L9:
	;
	v72 = *(*int32)(unsafe.Add(mBase, uint32(v23)+4))
	v74 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v72+v61))))
	if v74 != 0 {
		v237 = v63
		v239 = v65
		goto L11
	} else {
		goto L12
	}
L10:
	;
	v260 = v237
	goto L8
L11:
	;
	v247 = v61 + int32(1)
	if v247 < v239 {
		v61 = v247
		v63 = v237
		v65 = v239
		goto L9
	} else {
		goto L56
	}
L12:
	;
	v75 = *(*int32)(unsafe.Add(mBase, uint32(v26)+4))
	if v75 <= int32(0) {
		v237 = v63
		v239 = v65
		goto L11
	} else {
		goto L13
	}
L13:
	;
	v78 = *(*int32)(unsafe.Add(mBase, uint32(v23)+8))
	v82 = *(*int32)(unsafe.Add(mBase, uint32(v78+v61<<(uint(int32(3))%32))))
	v83 = int32(4)
	v84 = v82 + v83
	v85 = *(*int32)(unsafe.Add(mBase, uint32(v82)))
	v86 = int32(2)
	v89 = int32(base.Ui32(v85)>>(uint(v86)%32)) - v83
	v99 = v75
	v106 = int32(0)
	goto L14
L14:
	;
	v116 = v99 + v106
	v117 = int32(2)
	v118 = base.I32_div_s(v116, v117)
	v122 = *(*int32)(unsafe.Add(mBase, uint32(v51+v118<<(uint(v117)%32))))
	v126 = int32(base.Ui32(v122)>>(uint(int32(1))%32)) & int32(2047)
	if v89 == int32(0) {
		goto L21
	} else {
		goto L22
	}
L15:
	;
	if v116 < int32(-1) {
		v237 = v63
		v239 = v65
		goto L11
	} else {
		goto L55
	}
L16:
	;
	goto L15
L17:
	;
	if v214 < v212 {
		v99 = v212
		v106 = v214
		goto L14
	} else {
		goto L54
	}
L18:
	;
	v212 = v99
	v214 = v118 + int32(1)
	goto L17
L19:
	;
	if v206 == int32(0) {
		goto L16
	} else {
		goto L53
	}
L20:
	;
	if int32(0) <= v203 {
		v206 = v203
		goto L19
	} else {
		goto L52
	}
L21:
	;
	if v126 != 0 {
		goto L24
	} else {
		goto L25
	}
L22:
	;
	goto L23
L23:
	;
	if v126 == int32(0) {
		v206 = base.B2i32(base.Ui32(int32(19)) < base.Ui32(v85))
		goto L19
	} else {
		goto L27
	}
L24:
	;
	v131 = int32(-1)
	goto L26
L25:
	;
	v131 = int32(0)
	goto L26
L26:
	;
	v203 = v131
	goto L20
L27:
	;
	v136 = v51 + v75<<(uint(v86)%32) + int32(base.Ui32(v122)>>(uint(int32(12))%32))
	if base.Ui32(v89) < base.Ui32(v126) {
		goto L28
	} else {
		goto L29
	}
L28:
	;
	v138 = v89
	goto L30
L29:
	;
	v138 = v126
	goto L30
L30:
	;
	if base.Ui32(int32(4)) <= base.Ui32(v138) {
		goto L34
	} else {
		goto L35
	}
L31:
	;
	if v200 != 0 {
		v203 = v200
		goto L20
	} else {
		goto L49
	}
L32:
	;
	v200 = int32(0)
	goto L31
L33:
	;
	v174 = v169
	v175 = v170
	v176 = v171
	goto L43
L34:
	;
	if (v84|v136)&int32(3) != 0 {
		v169 = v84
		v170 = v136
		v171 = v138
		goto L33
	} else {
		goto L37
	}
L35:
	;
	v162 = v84
	v163 = v136
	v164 = v138
	goto L36
L36:
	;
	if v164 == int32(0) {
		goto L32
	} else {
		goto L42
	}
L37:
	;
	v146 = v84
	v147 = v136
	v148 = v138
	goto L38
L38:
	;
	v151 = *(*int32)(unsafe.Add(mBase, uint32(v146)))
	v152 = *(*int32)(unsafe.Add(mBase, uint32(v147)))
	if v151 != v152 {
		v169 = v146
		v170 = v147
		v171 = v148
		goto L33
	} else {
		goto L40
	}
L39:
	;
	v162 = v157
	v163 = v155
	v164 = v159
	goto L36
L40:
	;
	v154 = int32(4)
	v155 = v147 + v154
	v157 = v146 + v154
	v159 = v148 - v154
	if base.Ui32(int32(3)) < base.Ui32(v159) {
		v146 = v157
		v147 = v155
		v148 = v159
		goto L38
	} else {
		goto L41
	}
L41:
	;
	goto L39
L42:
	;
	v169 = v162
	v170 = v163
	v171 = v164
	goto L33
L43:
	;
	v179 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v174))))
	v180 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v175))))
	if v179 == v180 {
		goto L45
	} else {
		goto L46
	}
L44:
	;
	v200 = v179 - v180
	goto L31
L45:
	;
	v182 = int32(1)
	v187 = v176 - v182
	if v187 != 0 {
		v174 = v174 + v182
		v175 = v175 + v182
		v176 = v187
		goto L43
	} else {
		goto L48
	}
L46:
	;
	goto L47
L47:
	;
	goto L44
L48:
	;
	goto L32
L49:
	;
	if v126 == v89 {
		goto L16
	} else {
		goto L50
	}
L50:
	;
	if v126 <= v89 {
		goto L18
	} else {
		goto L51
	}
L51:
	;
	v212 = v118
	v214 = v106
	goto L17
L52:
	;
	v212 = v118
	v214 = v106
	goto L17
L53:
	;
	goto L18
L54:
	;
	v237 = v63
	v239 = v65
	goto L11
L55:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v45+v63<<(uint(int32(2))%32)))) = v118
	v225 = *(*int32)(unsafe.Add(mBase, uint32(v23)+12))
	v237 = v63 + int32(1)
	v239 = v225
	goto L11
L56:
	;
	goto L10
L57:
	;
	F_pfree(m, v45)
	mBase = m.M
	v272 = m.ExcPending
	if v272 != 0 {
		goto L1
	} else {
		goto L58
	}
L58:
	;
	v273 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
	if v273 != v26 {
		goto L59
	} else {
		goto L60
	}
L59:
	;
	F_pfree(m, v26)
	mBase = m.M
	v276 = m.ExcPending
	if v276 != 0 {
		goto L1
	} else {
		goto L62
	}
L60:
	;
	goto L61
L61:
	;
	v277 = *(*int32)(unsafe.Add(mBase, uint32(l0)+40))
	if v277 != v31 {
		goto L63
	} else {
		goto L64
	}
L62:
	;
	goto L61
L63:
	;
	F_pfree(m, v31)
	mBase = m.M
	v280 = m.ExcPending
	if v280 != 0 {
		goto L1
	} else {
		goto L66
	}
L64:
	;
	goto L65
L65:
	;
	m.G0 = v23 + int32(16)
	return base.I64_extend_i32_u(v269)
L66:
	;
	goto L65
}
func F_tsvector_gt(m *base.Module, l0 int32) int64 {
	mBase := m.M
	_ = mBase
	var v5 int32
	_ = v5
	var v6 int32
	_ = v6
	var v9 int32
	_ = v9
	var v10 int32
	_ = v10
	var v11 int32
	_ = v11
	var v12 int32
	_ = v12
	var v25 int32
	_ = v25
	var v26 int32
	_ = v26
	var v27 int32
	_ = v27
	var v28 int32
	_ = v28
	var v30 int32
	_ = v30
	var v33 int32
	_ = v33
	var v35 int32
	_ = v35
	var v36 int32
	_ = v36
	var v43 int32
	_ = v43
	var v44 int32
	_ = v44
	var v45 int32
	_ = v45
	var v47 int32
	_ = v47
	var v49 int32
	_ = v49
	var v52 int32
	_ = v52
	var v61 int32
	_ = v61
	var v62 int32
	_ = v62
	var v66 int32
	_ = v66
	var v67 int32
	_ = v67
	var v68 int32
	_ = v68
	var v69 int32
	_ = v69
	var v70 int32
	_ = v70
	var v72 int32
	_ = v72
	var v77 int32
	_ = v77
	var v78 int32
	_ = v78
	var v80 int32
	_ = v80
	var v81 int32
	_ = v81
	var v82 int32
	_ = v82
	var v83 int32
	_ = v83
	var v85 int32
	_ = v85
	var v89 int32
	_ = v89
	var v95 int32
	_ = v95
	var v96 int32
	_ = v96
	var v97 int32
	_ = v97
	var v101 int32
	_ = v101
	var v110 int32
	_ = v110
	var v112 int32
	_ = v112
	var v114 int32
	_ = v114
	var v115 int32
	_ = v115
	var v121 int32
	_ = v121
	var v122 int32
	_ = v122
	var v128 int32
	_ = v128
	var v129 int32
	_ = v129
	var v130 int32
	_ = v130
	var v131 int32
	_ = v131
	var v143 int32
	_ = v143
	var v144 int32
	_ = v144
	var v145 int32
	_ = v145
	var v146 int32
	_ = v146
	var v148 int32
	_ = v148
	var v153 int32
	_ = v153
	var v154 int32
	_ = v154
	var v155 int32
	_ = v155
	var v157 int32
	_ = v157
	var v159 int32
	_ = v159
	var v164 int32
	_ = v164
	var v169 int32
	_ = v169
	var v184 int32
	_ = v184
	var v190 int32
	_ = v190
	var v195 int32
	_ = v195
	var v220 int32
	_ = v220
	var v221 int32
	_ = v221
	var v224 int32
	_ = v224
	var v225 int32
	_ = v225
	var v228 int32
	_ = v228
	v5 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
	v6 = F_pg_detoast_datum(m, v5)
	mBase = m.M
	v9 = m.ExcPending
	if v9 != 0 {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	return int64(0)
L2:
	;
	v10 = *(*int32)(unsafe.Add(mBase, uint32(l0)+40))
	v11 = F_pg_detoast_datum(m, v10)
	mBase = m.M
	v12 = m.ExcPending
	if v12 != 0 {
		goto L1
	} else {
		goto L3
	}
L3:
	;
	v25 = *(*int32)(unsafe.Add(mBase, uint32(v6)))
	v26 = int32(2)
	v27 = int32(base.Ui32(v25) >> (uint(v26) % 32))
	v28 = *(*int32)(unsafe.Add(mBase, uint32(v11)))
	v30 = int32(base.Ui32(v28) >> (uint(v26) % 32))
	if base.Ui32(v27) < base.Ui32(v30) {
		goto L5
	} else {
		goto L6
	}
L4:
	;
	v221 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
	if v221 != v6 {
		goto L67
	} else {
		goto L68
	}
L5:
	;
	v220 = int32(-1)
	goto L4
L6:
	;
	goto L7
L7:
	;
	v33 = int32(1)
	if base.Ui32(v30) < base.Ui32(v27) {
		v195 = v33
		goto L8
	} else {
		goto L9
	}
L8:
	;
	v220 = v195
	goto L4
L9:
	;
	v35 = *(*int32)(unsafe.Add(mBase, uint32(v6)+4))
	v36 = *(*int32)(unsafe.Add(mBase, uint32(v11)+4))
	if v35 < v36 {
		goto L10
	} else {
		goto L11
	}
L10:
	;
	v220 = int32(-1)
	goto L4
L11:
	;
	goto L12
L12:
	;
	if v36 < v35 {
		v195 = v33
		goto L8
	} else {
		goto L13
	}
L13:
	;
	if v35 <= int32(0) {
		goto L14
	} else {
		goto L15
	}
L14:
	;
	v220 = int32(0)
	goto L4
L15:
	;
	goto L16
L16:
	;
	v43 = int32(8)
	v44 = v11 + v43
	v45 = int32(2)
	v47 = v44 + v36<<(uint(v45)%32)
	v49 = v6 + v43
	v52 = v49 + v35<<(uint(v45)%32)
	v61 = v44
	v62 = v49
	v66 = int32(0)
	goto L17
L17:
	;
	v67 = *(*int32)(unsafe.Add(mBase, uint32(v62)))
	v68 = int32(1)
	v69 = v67 & v68
	v70 = *(*int32)(unsafe.Add(mBase, uint32(v61)))
	v72 = v70 & v68
	if v69 != v72 {
		goto L19
	} else {
		goto L20
	}
L18:
	;
	v195 = int32(0)
	goto L8
L19:
	;
	if base.Ui32(v72) < base.Ui32(v69) {
		goto L22
	} else {
		goto L23
	}
L20:
	;
	goto L21
L21:
	;
	v78 = int32(1)
	v80 = int32(2047)
	v81 = int32(base.Ui32(v70)>>(uint(v78)%32)) & v80
	v82 = int32(12)
	v83 = int32(base.Ui32(v70) >> (uint(v82) % 32))
	v85 = int32(base.Ui32(v67) >> (uint(v82) % 32))
	v89 = int32(base.Ui32(v67)>>(uint(v78)%32)) & v80
	if v89 != 0 {
		goto L26
	} else {
		goto L27
	}
L22:
	;
	v77 = int32(-1)
	goto L24
L23:
	;
	v77 = int32(1)
	goto L24
L24:
	;
	v220 = v77
	goto L4
L25:
	;
	if v69 == int32(0) {
		goto L41
	} else {
		goto L42
	}
L26:
	;
	if v81 == int32(0) {
		goto L29
	} else {
		goto L30
	}
L27:
	;
	goto L28
L28:
	;
	if v81 == int32(0) {
		goto L25
	} else {
		goto L40
	}
L29:
	;
	v220 = int32(1)
	goto L4
L30:
	;
	goto L31
L31:
	;
	v95 = base.B2i32(base.Ui32(v89) < base.Ui32(v81))
	if base.Ui32(v89) < base.Ui32(v81) {
		goto L32
	} else {
		goto L33
	}
L32:
	;
	v96 = v89
	goto L34
L33:
	;
	v96 = v81
	goto L34
L34:
	;
	v97 = F_memcmp(m, v85+v52, v83+v47, v96)
	mBase = m.M
	if v97 != 0 {
		v195 = v97
		goto L8
	} else {
		goto L35
	}
L35:
	;
	if v81 == v89 {
		goto L25
	} else {
		goto L36
	}
L36:
	;
	if base.Ui32(v89) < base.Ui32(v81) {
		goto L37
	} else {
		goto L38
	}
L37:
	;
	v101 = int32(-1)
	goto L39
L38:
	;
	v101 = int32(1)
	goto L39
L39:
	;
	v220 = v101
	goto L4
L40:
	;
	v220 = int32(-1)
	goto L4
L41:
	;
	v184 = int32(4)
	v190 = v66 + int32(1)
	if v190 != v35 {
		v61 = v61 + v184
		v62 = v62 + v184
		v66 = v190
		goto L17
	} else {
		goto L66
	}
L42:
	;
	v110 = int32(1)
	v112 = int32(_a_F_tsvector_gt_0)
	v114 = v47 + (v81+v83+v110)&v112
	v115 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v114))))
	v121 = v52 + (v89+v85+v110)&v112
	v122 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v121))))
	if v115 == v122 {
		goto L44
	} else {
		goto L45
	}
L43:
	;
	v129 = v121
	v130 = v114
	v131 = int32(0)
	goto L51
L44:
	;
	if v122 != 0 {
		goto L43
	} else {
		goto L47
	}
L45:
	;
	goto L46
L46:
	;
	if base.Ui32(v115) < base.Ui32(v122) {
		goto L48
	} else {
		goto L49
	}
L47:
	;
	goto L41
L48:
	;
	v128 = int32(-1)
	goto L50
L49:
	;
	v128 = int32(1)
	goto L50
L50:
	;
	v220 = v128
	goto L4
L51:
	;
	v143 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v129)+2)))
	v144 = int32(_a_F_tsvector_gt_1)
	v145 = v143 & v144
	v146 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v130)+2)))
	v148 = v146 & v144
	if v145 != v148 {
		goto L53
	} else {
		goto L54
	}
L52:
	;
	if base.Ui32(v157) < base.Ui32(v155) {
		goto L63
	} else {
		goto L64
	}
L53:
	;
	if base.Ui32(v148) < base.Ui32(v145) {
		goto L56
	} else {
		goto L57
	}
L54:
	;
	goto L55
L55:
	;
	v154 = int32(14)
	v155 = int32(base.Ui32(v143) >> (uint(v154) % 32))
	v157 = int32(base.Ui32(v146) >> (uint(v154) % 32))
	if v155 == v157 {
		goto L59
	} else {
		goto L60
	}
L56:
	;
	v153 = int32(-1)
	goto L58
L57:
	;
	v153 = int32(1)
	goto L58
L58:
	;
	v220 = v153
	goto L4
L59:
	;
	v159 = int32(2)
	v164 = v131 + int32(1)
	if v164 == v122 {
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
	v129 = v129 + v159
	v130 = v130 + v159
	v131 = v164
	goto L51
L63:
	;
	v169 = int32(-1)
	goto L65
L64:
	;
	v169 = int32(1)
	goto L65
L65:
	;
	v195 = v169
	goto L8
L66:
	;
	goto L18
L67:
	;
	F_pfree(m, v6)
	mBase = m.M
	v224 = m.ExcPending
	if v224 != 0 {
		goto L1
	} else {
		goto L70
	}
L68:
	;
	goto L69
L69:
	;
	v225 = *(*int32)(unsafe.Add(mBase, uint32(l0)+40))
	if v225 != v11 {
		goto L71
	} else {
		goto L72
	}
L70:
	;
	goto L69
L71:
	;
	F_pfree(m, v11)
	mBase = m.M
	v228 = m.ExcPending
	if v228 != 0 {
		goto L1
	} else {
		goto L74
	}
L72:
	;
	goto L73
L73:
	;
	return base.I64_extend_i32_u(base.B2i32(int32(0) < v220))
L74:
	;
	goto L73
}
func F_tsvector_setweight_by_filter(m *base.Module, l0 int32) int64 {
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
	var v45 int32
	_ = v45
	var v55 int32
	_ = v55
	var v56 int32
	_ = v56
	var v60 int32
	_ = v60
	var v62 int32
	_ = v62
	var v63 int32
	_ = v63
	var v76 int32
	_ = v76
	var v87 int32
	_ = v87
	var v88 int32
	_ = v88
	var v91 int32
	_ = v91
	var v95 int32
	_ = v95
	var v96 int32
	_ = v96
	var v97 int32
	_ = v97
	var v98 int32
	_ = v98
	var v99 int32
	_ = v99
	var v102 int32
	_ = v102
	var v107 int32
	_ = v107
	var v111 int32
	_ = v111
	var v112 int32
	_ = v112
	var v131 int32
	_ = v131
	var v132 int32
	_ = v132
	var v133 int32
	_ = v133
	var v137 int32
	_ = v137
	var v141 int32
	_ = v141
	var v143 int32
	_ = v143
	var v148 int32
	_ = v148
	var v151 int32
	_ = v151
	var v153 int32
	_ = v153
	var v161 int32
	_ = v161
	var v162 int32
	_ = v162
	var v163 int32
	_ = v163
	var v166 int32
	_ = v166
	var v167 int32
	_ = v167
	var v169 int32
	_ = v169
	var v170 int32
	_ = v170
	var v172 int32
	_ = v172
	var v174 int32
	_ = v174
	var v177 int32
	_ = v177
	var v178 int32
	_ = v178
	var v179 int32
	_ = v179
	var v184 int32
	_ = v184
	var v185 int32
	_ = v185
	var v186 int32
	_ = v186
	var v189 int32
	_ = v189
	var v190 int32
	_ = v190
	var v191 int32
	_ = v191
	var v194 int32
	_ = v194
	var v195 int32
	_ = v195
	var v197 int32
	_ = v197
	var v202 int32
	_ = v202
	var v215 int32
	_ = v215
	var v218 int32
	_ = v218
	var v221 int32
	_ = v221
	var v227 int32
	_ = v227
	var v228 int32
	_ = v228
	var v244 int32
	_ = v244
	var v245 int32
	_ = v245
	var v250 int32
	_ = v250
	var v252 int32
	_ = v252
	var v254 int32
	_ = v254
	var v257 int32
	_ = v257
	var v273 int32
	_ = v273
	var v276 int32
	_ = v276
	var v278 int32
	_ = v278
	var v279 int32
	_ = v279
	var v281 int32
	_ = v281
	var v283 int32
	_ = v283
	var v286 int32
	_ = v286
	var v288 int32
	_ = v288
	var v310 int32
	_ = v310
	var v312 int32
	_ = v312
	var v331 int32
	_ = v331
	var v332 int32
	_ = v332
	var v334 int32
	_ = v334
	var v336 int32
	_ = v336
	var v339 int32
	_ = v339
	var v341 int32
	_ = v341
	var v344 int32
	_ = v344
	var v346 int32
	_ = v346
	var v349 int32
	_ = v349
	var v354 int32
	_ = v354
	var v378 int32
	_ = v378
	var v402 int32
	_ = v402
	var v405 int32
	_ = v405
	var v406 int32
	_ = v406
	var v409 int32
	_ = v409
	v23 = m.G0
	v25 = v23 - int32(16)
	m.G0 = v25
	v27 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
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
	return int64(0)
L2:
	;
	v32 = int32(*(*int8)(unsafe.Add(mBase, uint32(l0)+40)))
	v33 = *(*int32)(unsafe.Add(mBase, uint32(l0)+56))
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
	v36 = F_parse_weight(m, v32)
	mBase = m.M
	v37 = m.ExcPending
	if v37 != 0 {
		goto L1
	} else {
		goto L4
	}
L4:
	;
	v38 = *(*int32)(unsafe.Add(mBase, uint32(v28)))
	v41 = F_palloc(m, int32(base.Ui32(v38)>>(uint(int32(2))%32)))
	mBase = m.M
	v42 = m.ExcPending
	if v42 != 0 {
		goto L1
	} else {
		goto L5
	}
L5:
	;
	v43 = *(*int32)(unsafe.Add(mBase, uint32(v28)))
	v45 = int32(base.Ui32(v43) >> (uint(int32(2)) % 32))
	if v45 != 0 {
		goto L6
	} else {
		goto L7
	}
L6:
	;
	base.MemoryCopy(m, v41, v28, v45)
	goto L8
L7:
	;
	goto L8
L8:
	;
	F_deconstruct_array_builtin(m, v34, int32(25), v25+int32(8), v25+int32(4), v25+int32(12))
	mBase = m.M
	v55 = m.ExcPending
	if v55 != 0 {
		goto L1
	} else {
		goto L9
	}
L9:
	;
	v56 = *(*int32)(unsafe.Add(mBase, uint32(v25)+12))
	if int32(0) < v56 {
		goto L10
	} else {
		goto L11
	}
L10:
	;
	v60 = v41 + int32(8)
	v62 = v36 << (uint(int32(14)) % 32)
	v63 = *(*int32)(unsafe.Add(mBase, uint32(v25)+4))
	v76 = int32(0)
	goto L13
L11:
	;
	goto L12
L12:
	;
	v402 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
	if v402 != v28 {
		goto L72
	} else {
		goto L73
	}
L13:
	;
	v87 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v76+v63))))
	if v87 != 0 {
		goto L15
	} else {
		goto L16
	}
L14:
	;
	goto L12
L15:
	;
	v378 = v76 + int32(1)
	if v378 < v56 {
		v76 = v378
		goto L13
	} else {
		goto L71
	}
L16:
	;
	v88 = *(*int32)(unsafe.Add(mBase, uint32(v41)+4))
	if v88 <= int32(0) {
		goto L15
	} else {
		goto L17
	}
L17:
	;
	v91 = *(*int32)(unsafe.Add(mBase, uint32(v25)+8))
	v95 = *(*int32)(unsafe.Add(mBase, uint32(v91+v76<<(uint(int32(3))%32))))
	v96 = int32(4)
	v97 = v95 + v96
	v98 = *(*int32)(unsafe.Add(mBase, uint32(v95)))
	v99 = int32(2)
	v102 = int32(base.Ui32(v98)>>(uint(v99)%32)) - v96
	v107 = v60 + v88<<(uint(v99)%32)
	v111 = v88
	v112 = int32(0)
	goto L18
L18:
	;
	v131 = v111 + v112
	v132 = int32(2)
	v133 = base.I32_div_s(v131, v132)
	v137 = *(*int32)(unsafe.Add(mBase, uint32(v60+v133<<(uint(v132)%32))))
	v141 = int32(base.Ui32(v137)>>(uint(int32(1))%32)) & int32(2047)
	v143 = int32(base.Ui32(v137) >> (uint(int32(12)) % 32))
	if v102 == int32(0) {
		goto L25
	} else {
		goto L26
	}
L19:
	;
	if base.B2i32(v137&int32(1) == int32(0))|base.B2i32(v131 < int32(-1)) != 0 {
		goto L15
	} else {
		goto L59
	}
L20:
	;
	goto L19
L21:
	;
	if v228 < v227 {
		v111 = v227
		v112 = v228
		goto L18
	} else {
		goto L58
	}
L22:
	;
	v227 = v111
	v228 = v133 + int32(1)
	goto L21
L23:
	;
	if v221 == int32(0) {
		goto L20
	} else {
		goto L57
	}
L24:
	;
	if int32(0) <= v218 {
		v221 = v218
		goto L23
	} else {
		goto L56
	}
L25:
	;
	if v141 != 0 {
		goto L28
	} else {
		goto L29
	}
L26:
	;
	goto L27
L27:
	;
	if v141 == int32(0) {
		v221 = base.B2i32(base.Ui32(int32(19)) < base.Ui32(v98))
		goto L23
	} else {
		goto L31
	}
L28:
	;
	v148 = int32(-1)
	goto L30
L29:
	;
	v148 = int32(0)
	goto L30
L30:
	;
	v218 = v148
	goto L24
L31:
	;
	v151 = v107 + v143
	if base.Ui32(v102) < base.Ui32(v141) {
		goto L32
	} else {
		goto L33
	}
L32:
	;
	v153 = v102
	goto L34
L33:
	;
	v153 = v141
	goto L34
L34:
	;
	if base.Ui32(int32(4)) <= base.Ui32(v153) {
		goto L38
	} else {
		goto L39
	}
L35:
	;
	if v215 != 0 {
		v218 = v215
		goto L24
	} else {
		goto L53
	}
L36:
	;
	v215 = int32(0)
	goto L35
L37:
	;
	v189 = v184
	v190 = v185
	v191 = v186
	goto L47
L38:
	;
	if (v97|v151)&int32(3) != 0 {
		v184 = v97
		v185 = v151
		v186 = v153
		goto L37
	} else {
		goto L41
	}
L39:
	;
	v177 = v97
	v178 = v151
	v179 = v153
	goto L40
L40:
	;
	if v179 == int32(0) {
		goto L36
	} else {
		goto L46
	}
L41:
	;
	v161 = v97
	v162 = v151
	v163 = v153
	goto L42
L42:
	;
	v166 = *(*int32)(unsafe.Add(mBase, uint32(v161)))
	v167 = *(*int32)(unsafe.Add(mBase, uint32(v162)))
	if v166 != v167 {
		v184 = v161
		v185 = v162
		v186 = v163
		goto L37
	} else {
		goto L44
	}
L43:
	;
	v177 = v172
	v178 = v170
	v179 = v174
	goto L40
L44:
	;
	v169 = int32(4)
	v170 = v162 + v169
	v172 = v161 + v169
	v174 = v163 - v169
	if base.Ui32(int32(3)) < base.Ui32(v174) {
		v161 = v172
		v162 = v170
		v163 = v174
		goto L42
	} else {
		goto L45
	}
L45:
	;
	goto L43
L46:
	;
	v184 = v177
	v185 = v178
	v186 = v179
	goto L37
L47:
	;
	v194 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v189))))
	v195 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v190))))
	if v194 == v195 {
		goto L49
	} else {
		goto L50
	}
L48:
	;
	v215 = v194 - v195
	goto L35
L49:
	;
	v197 = int32(1)
	v202 = v191 - v197
	if v202 != 0 {
		v189 = v189 + v197
		v190 = v190 + v197
		v191 = v202
		goto L47
	} else {
		goto L52
	}
L50:
	;
	goto L51
L51:
	;
	goto L48
L52:
	;
	goto L36
L53:
	;
	if v141 == v102 {
		goto L20
	} else {
		goto L54
	}
L54:
	;
	if v141 <= v102 {
		goto L22
	} else {
		goto L55
	}
L55:
	;
	v227 = v133
	v228 = v112
	goto L21
L56:
	;
	v227 = v133
	v228 = v112
	goto L21
L57:
	;
	goto L22
L58:
	;
	goto L15
L59:
	;
	v244 = v107 + (v141+v143+int32(1))&int32(_a_F_tsvector_setweight_by_filter_0)
	v245 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v244))))
	if v245 == int32(0) {
		goto L15
	} else {
		goto L60
	}
L60:
	;
	v250 = v245 & int32(3)
	if v250 != 0 {
		goto L61
	} else {
		goto L62
	}
L61:
	;
	v252 = v244
	v254 = v245
	v257 = int32(0)
	goto L64
L62:
	;
	v286 = v244
	v288 = v245
	goto L63
L63:
	;
	if base.Ui32(v245) < base.Ui32(int32(4)) {
		goto L15
	} else {
		goto L67
	}
L64:
	;
	v273 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v252)+2)))
	v276 = v273&int32(_a_F_tsvector_setweight_by_filter_1) | v62
	*(*uint16)(unsafe.Add(mBase, uint32(v252)+2)) = uint16(v276)
	v278 = int32(1)
	v279 = v254 - v278
	v281 = v252 + int32(2)
	v283 = v257 + v278
	if v283 != v250 {
		v252 = v281
		v254 = v279
		v257 = v283
		goto L64
	} else {
		goto L66
	}
L65:
	;
	v286 = v281
	v288 = v279
	goto L63
L66:
	;
	goto L65
L67:
	;
	v310 = v286
	v312 = v288
	goto L68
L68:
	;
	v331 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v310)+2)))
	v332 = int32(_a_F_tsvector_setweight_by_filter_1)
	v334 = v331&v332 | v62
	*(*uint16)(unsafe.Add(mBase, uint32(v310)+2)) = uint16(v334)
	v336 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v310)+4)))
	v339 = v336&v332 | v62
	*(*uint16)(unsafe.Add(mBase, uint32(v310)+4)) = uint16(v339)
	v341 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v310)+6)))
	v344 = v341&v332 | v62
	*(*uint16)(unsafe.Add(mBase, uint32(v310)+6)) = uint16(v344)
	v346 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v310)+8)))
	v349 = v346&v332 | v62
	*(*uint16)(unsafe.Add(mBase, uint32(v310)+8)) = uint16(v349)
	v354 = v312 - int32(4)
	if v354 != 0 {
		v310 = v310 + int32(8)
		v312 = v354
		goto L68
	} else {
		goto L70
	}
L69:
	;
	goto L15
L70:
	;
	goto L69
L71:
	;
	goto L14
L72:
	;
	F_pfree(m, v28)
	mBase = m.M
	v405 = m.ExcPending
	if v405 != 0 {
		goto L1
	} else {
		goto L75
	}
L73:
	;
	goto L74
L74:
	;
	v406 = *(*int32)(unsafe.Add(mBase, uint32(l0)+56))
	if v406 != v34 {
		goto L76
	} else {
		goto L77
	}
L75:
	;
	goto L74
L76:
	;
	F_pfree(m, v34)
	mBase = m.M
	v409 = m.ExcPending
	if v409 != 0 {
		goto L1
	} else {
		goto L79
	}
L77:
	;
	goto L78
L78:
	;
	m.G0 = v25 + int32(16)
	return base.I64_extend_i32_u(v41)
L79:
	;
	goto L78
}
func F_tsvector_update_trigger_byid(m *base.Module, l0 int32) int64 {
	var v3 int64
	_ = v3
	var v6 int32
	_ = v6
	v3 = F_tsvector_update_trigger(m, l0, int32(0))
	v6 = m.ExcPending
	if v6 != 0 {
		return int64(0)
	} else {
		return v3
	}
}
