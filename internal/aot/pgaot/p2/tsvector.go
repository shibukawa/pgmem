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
	var v65 int32
	_ = v65
	var v70 int32
	_ = v70
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
	var v103 int32
	_ = v103
	var v104 int32
	_ = v104
	var v108 int32
	_ = v108
	var v112 int32
	_ = v112
	var v113 int32
	_ = v113
	var v114 int32
	_ = v114
	var v115 int32
	_ = v115
	var v119 int32
	_ = v119
	var v130 int32
	_ = v130
	var v134 int32
	_ = v134
	var v137 int32
	_ = v137
	var v141 int32
	_ = v141
	var v146 int32
	_ = v146
	var v151 int32
	_ = v151
	var v152 int32
	_ = v152
	var v162 int32
	_ = v162
	var v163 int32
	_ = v163
	var v164 int32
	_ = v164
	var v167 int32
	_ = v167
	var v171 int32
	_ = v171
	var v176 int32
	_ = v176
	var v177 int32
	_ = v177
	var v178 int32
	_ = v178
	var v180 int32
	_ = v180
	var v183 int32
	_ = v183
	var v188 int32
	_ = v188
	var v189 int32
	_ = v189
	var v193 int32
	_ = v193
	var v195 int32
	_ = v195
	var v196 int32
	_ = v196
	var v197 int32
	_ = v197
	var v202 int32
	_ = v202
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
	var v267 int32
	_ = v267
	var v286 int32
	_ = v286
	var v287 int32
	_ = v287
	var v295 int32
	_ = v295
	var v299 int32
	_ = v299
	var v312 float64
	_ = v312
	var v315 int32
	_ = v315
	var v322 int32
	_ = v322
	var v326 int32
	_ = v326
	var v327 int32
	_ = v327
	var v336 float64
	_ = v336
	var v339 int32
	_ = v339
	var v351 int32
	_ = v351
	var v352 int32
	_ = v352
	var v361 float64
	_ = v361
	var v364 int32
	_ = v364
	var v369 float32
	_ = v369
	var v372 float64
	_ = v372
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
	var v446 int32
	_ = v446
	var v449 int32
	_ = v449
	var v450 int32
	_ = v450
	var v452 int32
	_ = v452
	var v454 int32
	_ = v454
	var v457 int32
	_ = v457
	var v458 int32
	_ = v458
	var v459 int32
	_ = v459
	var v464 int32
	_ = v464
	var v467 int32
	_ = v467
	var v470 int32
	_ = v470
	var v472 int32
	_ = v472
	var v473 int32
	_ = v473
	var v488 int32
	_ = v488
	var v494 int32
	_ = v494
	var v496 int32
	_ = v496
	var v500 int32
	_ = v500
	var v501 int32
	_ = v501
	var v503 int32
	_ = v503
	var v504 int32
	_ = v504
	var v507 int32
	_ = v507
	var v508 int32
	_ = v508
	var v511 int32
	_ = v511
	var v516 int32
	_ = v516
	var v532 float64
	_ = v532
	var v535 int32
	_ = v535
	var v536 int32
	_ = v536
	var v544 int32
	_ = v544
	var v549 int32
	_ = v549
	var v555 int32
	_ = v555
	var v561 int32
	_ = v561
	var v562 int32
	_ = v562
	var v563 int32
	_ = v563
	var v564 int32
	_ = v564
	var v567 int32
	_ = v567
	var v572 int32
	_ = v572
	var v573 int32
	_ = v573
	var v574 int32
	_ = v574
	var v576 int32
	_ = v576
	var v580 int32
	_ = v580
	var v581 int32
	_ = v581
	var v582 int32
	_ = v582
	var v583 int32
	_ = v583
	var v586 int32
	_ = v586
	var v587 int32
	_ = v587
	var v592 int32
	_ = v592
	var v611 int32
	_ = v611
	var v614 int32
	_ = v614
	var v615 int32
	_ = v615
	var v616 int32
	_ = v616
	var v617 int32
	_ = v617
	var v618 int32
	_ = v618
	var v621 int32
	_ = v621
	var v627 int32
	_ = v627
	var v631 int32
	_ = v631
	var v646 int32
	_ = v646
	var v650 int32
	_ = v650
	var v652 int32
	_ = v652
	var v654 int32
	_ = v654
	var v663 int32
	_ = v663
	var v693 int32
	_ = v693
	var v697 int32
	_ = v697
	var v702 int32
	_ = v702
	v5 = int32(0)
	v21 = float64(0)
	v23 = m.G0
	v25 = v23 - int32(112)
	m.G0 = v25
	v27 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	*(*int32)(unsafe.Add(mBase, uint32(v25)+72)) = int32(1166)
	*(*int32)(unsafe.Add(mBase, uint32(v25)+68)) = int32(1167)
	*(*int64)(unsafe.Add(mBase, uint32(v25)+60)) = int64(68719476744)
	v35 = *(*int32)(unsafe.Add(mBase, _c_F_compute_tsvector_stats[0]))
	*(*int32)(unsafe.Add(mBase, uint32(v25)+84)) = v35
	v37 = int32(_a_F_compute_tsvector_stats_0)
	v42 = base.I32_div_s(v27*v37+v37, int32(7))
	v45 = v27 * int32(10)
	v49 = F_hash_create(m, int32(_a_F_compute_tsvector_stats_1), v45, v25+int32(44), int32(1224))
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
	if int32(0) < l2 {
		goto L4
	} else {
		goto L5
	}
L3:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v693 = m.ExcPending
	if v693 != 0 {
		goto L1
	} else {
		goto L106
	}
L4:
	;
	v60 = int32(1)
	v64 = v5
	v65 = v5
	v70 = v5
	v74 = v21
	goto L7
L5:
	;
	v351 = v5
	v352 = v5
	v361 = v21
	goto L6
L6:
	;
	if v352 < l2 {
		goto L60
	} else {
		goto L61
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
	v351 = v326
	v352 = v327
	v361 = v336
	goto L6
L9:
	;
	v81 = m.T0[l1].(func(*base.Module, int32, int32, int32) int32)(m, l0, v70, v25+int32(35))
	mBase = m.M
	v82 = m.ExcPending
	if v82 != 0 {
		goto L1
	} else {
		goto L10
	}
L10:
	;
	v83 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v25)+35)))
	if v83 == int32(1) {
		goto L12
	} else {
		goto L13
	}
L11:
	;
	v339 = v70 + int32(1)
	if v339 != l2 {
		v60 = v322
		v64 = v326
		v65 = v327
		v70 = v339
		v74 = v336
		goto L7
	} else {
		goto L58
	}
L12:
	;
	v322 = v60
	v326 = v64
	v327 = v65 + int32(1)
	v336 = v74
	goto L11
L13:
	;
	goto L14
L14:
	;
	v88 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v81))))
	if v88 == int32(1) {
		goto L16
	} else {
		goto L17
	}
L15:
	;
	v113 = F_pg_detoast_datum(m, v81)
	mBase = m.M
	v114 = m.ExcPending
	if v114 != 0 {
		goto L1
	} else {
		goto L24
	}
L16:
	;
	v92 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v81)+1)))
	if base.Ui32((v92-int32(1))&int32(255)) < base.Ui32(int32(3)) {
		v112 = int32(6)
		goto L15
	} else {
		goto L19
	}
L17:
	;
	goto L18
L18:
	;
	v104 = int32(1)
	if v88&v104 != 0 {
		v112 = int32(base.Ui32(v88) >> (uint(v104) % 32))
		goto L15
	} else {
		goto L23
	}
L19:
	;
	v99 = int32(18)
	if v92 == v99 {
		goto L20
	} else {
		goto L21
	}
L20:
	;
	v103 = v99
	goto L22
L21:
	;
	v103 = int32(2)
	goto L22
L22:
	;
	v112 = v103
	goto L15
L23:
	;
	v108 = *(*int32)(unsafe.Add(mBase, uint32(v81)))
	v112 = int32(base.Ui32(v108) >> (uint(int32(2)) % 32))
	goto L15
L24:
	;
	v115 = *(*int32)(unsafe.Add(mBase, uint32(v113)+4))
	if int32(0) < v115 {
		goto L25
	} else {
		goto L26
	}
L25:
	;
	v119 = v113 + int32(8)
	v130 = v60
	v134 = v64
	v137 = v119
	v141 = int32(0)
	goto L28
L26:
	;
	v295 = v60
	v299 = v64
	goto L27
L27:
	;
	v312 = base.F64_add(v74, base.F64_convert_i32_u(v112))
	if v81 == v113 {
		v322 = v295
		v326 = v299
		v327 = v65
		v336 = v312
		goto L11
	} else {
		goto L56
	}
L28:
	;
	v146 = *(*int32)(unsafe.Add(mBase, uint32(v137)))
	*(*int32)(unsafe.Add(mBase, uint32(v25)+36)) = v119 + v115<<(uint(int32(2))%32) + int32(base.Ui32(v146)>>(uint(int32(12))%32))
	v151 = *(*int32)(unsafe.Add(mBase, uint32(v137)))
	v152 = int32(1)
	*(*int32)(unsafe.Add(mBase, uint32(v25)+40)) = int32(base.Ui32(v151)>>(uint(v152)%32)) & int32(2047)
	v162 = F_hash_search(m, v49, v25+int32(36), v152, v25+int32(34))
	mBase = m.M
	v163 = m.ExcPending
	if v163 != 0 {
		goto L1
	} else {
		goto L30
	}
L29:
	;
	v295 = v267
	v299 = v188
	goto L27
L30:
	;
	v164 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v25)+34)))
	if v164 == int32(1) {
		goto L32
	} else {
		goto L33
	}
L31:
	;
	v188 = v134 + int32(1)
	v189 = base.I32_rem_s(v188, v42)
	if v189 == int32(0) {
		goto L37
	} else {
		goto L38
	}
L32:
	;
	v167 = *(*int32)(unsafe.Add(mBase, uint32(v162)+8))
	*(*int32)(unsafe.Add(mBase, uint32(v162)+8)) = v167 + int32(1)
	goto L31
L33:
	;
	goto L34
L34:
	;
	v171 = int32(1)
	*(*int32)(unsafe.Add(mBase, uint32(v162)+8)) = v171
	*(*int32)(unsafe.Add(mBase, uint32(v162)+12)) = v130 - v171
	v176 = *(*int32)(unsafe.Add(mBase, uint32(v25)+40))
	v177 = F_palloc(m, v176)
	mBase = m.M
	v178 = m.ExcPending
	if v178 != 0 {
		goto L1
	} else {
		goto L35
	}
L35:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v162))) = v177
	v180 = *(*int32)(unsafe.Add(mBase, uint32(v25)+40))
	if v180 == int32(0) {
		goto L31
	} else {
		goto L36
	}
L36:
	;
	v183 = *(*int32)(unsafe.Add(mBase, uint32(v25)+36))
	base.MemoryCopy(m, v177, v183, v180)
	goto L31
L37:
	;
	v193 = v25 + int32(92)
	F_hash_seq_init(m, v193, v49)
	mBase = m.M
	v195 = m.ExcPending
	if v195 != 0 {
		goto L1
	} else {
		goto L40
	}
L38:
	;
	v267 = v130
	goto L39
L39:
	;
	v286 = v141 + int32(1)
	v287 = *(*int32)(unsafe.Add(mBase, uint32(v113)+4))
	if v286 < v287 {
		v130 = v267
		v134 = v188
		v137 = v137 + int32(4)
		v141 = v286
		goto L28
	} else {
		goto L55
	}
L40:
	;
	v196 = F_hash_seq_search(m, v193)
	mBase = m.M
	v197 = m.ExcPending
	if v197 != 0 {
		goto L1
	} else {
		goto L41
	}
L41:
	;
	if v196 != 0 {
		goto L42
	} else {
		goto L43
	}
L42:
	;
	v202 = v196
	goto L45
L43:
	;
	goto L44
L44:
	;
	v267 = v130 + int32(1)
	goto L39
L45:
	;
	v220 = *(*int32)(unsafe.Add(mBase, uint32(v202)+12))
	v221 = *(*int32)(unsafe.Add(mBase, uint32(v202)+8))
	if v220+v221 <= v130 {
		goto L47
	} else {
		goto L48
	}
L46:
	;
	goto L44
L47:
	;
	v224 = *(*int32)(unsafe.Add(mBase, uint32(v202)))
	v227 = F_hash_search(m, v49, v202, int32(2), int32(0))
	mBase = m.M
	v228 = m.ExcPending
	if v228 != 0 {
		goto L1
	} else {
		goto L50
	}
L48:
	;
	goto L49
L49:
	;
	v235 = F_hash_seq_search(m, v25+int32(92))
	mBase = m.M
	v236 = m.ExcPending
	if v236 != 0 {
		goto L1
	} else {
		goto L53
	}
L50:
	;
	if v227 == int32(0) {
		goto L3
	} else {
		goto L51
	}
L51:
	;
	F_pfree(m, v224)
	mBase = m.M
	v232 = m.ExcPending
	if v232 != 0 {
		goto L1
	} else {
		goto L52
	}
L52:
	;
	goto L49
L53:
	;
	if v235 != 0 {
		v202 = v235
		goto L45
	} else {
		goto L54
	}
L54:
	;
	goto L46
L55:
	;
	goto L29
L56:
	;
	F_pfree(m, v113)
	mBase = m.M
	v315 = m.ExcPending
	if v315 != 0 {
		goto L1
	} else {
		goto L57
	}
L57:
	;
	v322 = v295
	v326 = v299
	v327 = v65
	v336 = v312
	goto L11
L58:
	;
	goto L8
L59:
	;
	m.G0 = v25 + int32(112)
	return
L60:
	;
	v364 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(l0)+36)) = uint8(v364)
	v369 = base.F32_demote_f64(base.F64_div(base.F64_convert_i32_s(v352), base.F64_convert_i32_s(l2)))
	*(*float32)(unsafe.Add(mBase, uint32(l0)+40)) = v369
	v372 = base.F64_convert_i32_s(l2 - v352)
	*(*int32)(unsafe.Add(mBase, uint32(l0)+44)) = base.I32_trunc_sat_f64_s(base.F64_div(v361, v372))
	*(*float32)(unsafe.Add(mBase, uint32(l0)+48)) = base.F32_neg(base.F32_sub(float32(1), v369))
	v381 = *(*int32)(unsafe.Add(mBase, uint32(v49)))
	v382 = *(*int32)(unsafe.Add(mBase, uint32(v381)+4))
	v383 = *(*int32)(unsafe.Add(mBase, uint32(v381)+412))
	if v383 != 0 {
		goto L64
	} else {
		goto L65
	}
L61:
	;
	goto L62
L62:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+48)) = int32(0)
	*(*int64)(unsafe.Add(mBase, uint32(l0)+40)) = int64(1065353216)
	v663 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(l0)+36)) = uint8(v663)
	goto L59
L63:
	;
	v449 = F_palloc(m, v446<<(uint(int32(2))%32))
	mBase = m.M
	v450 = m.ExcPending
	if v450 != 0 {
		goto L1
	} else {
		goto L67
	}
L64:
	;
	v384 = *(*int32)(unsafe.Add(mBase, uint32(v381)+376))
	v385 = *(*int32)(unsafe.Add(mBase, uint32(v381)+364))
	v386 = *(*int32)(unsafe.Add(mBase, uint32(v381)+352))
	v387 = *(*int32)(unsafe.Add(mBase, uint32(v381)+340))
	v388 = *(*int32)(unsafe.Add(mBase, uint32(v381)+328))
	v389 = *(*int32)(unsafe.Add(mBase, uint32(v381)+316))
	v390 = *(*int32)(unsafe.Add(mBase, uint32(v381)+304))
	v391 = *(*int32)(unsafe.Add(mBase, uint32(v381)+292))
	v392 = *(*int32)(unsafe.Add(mBase, uint32(v381)+280))
	v393 = *(*int32)(unsafe.Add(mBase, uint32(v381)+268))
	v394 = *(*int32)(unsafe.Add(mBase, uint32(v381)+256))
	v395 = *(*int32)(unsafe.Add(mBase, uint32(v381)+244))
	v396 = *(*int32)(unsafe.Add(mBase, uint32(v381)+232))
	v397 = *(*int32)(unsafe.Add(mBase, uint32(v381)+220))
	v398 = *(*int32)(unsafe.Add(mBase, uint32(v381)+208))
	v399 = *(*int32)(unsafe.Add(mBase, uint32(v381)+196))
	v400 = *(*int32)(unsafe.Add(mBase, uint32(v381)+184))
	v401 = *(*int32)(unsafe.Add(mBase, uint32(v381)+172))
	v402 = *(*int32)(unsafe.Add(mBase, uint32(v381)+160))
	v403 = *(*int32)(unsafe.Add(mBase, uint32(v381)+148))
	v404 = *(*int32)(unsafe.Add(mBase, uint32(v381)+136))
	v405 = *(*int32)(unsafe.Add(mBase, uint32(v381)+124))
	v406 = *(*int32)(unsafe.Add(mBase, uint32(v381)+112))
	v407 = *(*int32)(unsafe.Add(mBase, uint32(v381)+100))
	v408 = *(*int32)(unsafe.Add(mBase, uint32(v381)+88))
	v409 = *(*int32)(unsafe.Add(mBase, uint32(v381)+76))
	v410 = *(*int32)(unsafe.Add(mBase, uint32(v381)+64))
	v411 = *(*int32)(unsafe.Add(mBase, uint32(v381)+52))
	v412 = *(*int32)(unsafe.Add(mBase, uint32(v381)+40))
	v413 = *(*int32)(unsafe.Add(mBase, uint32(v381)+28))
	v414 = *(*int32)(unsafe.Add(mBase, uint32(v381)+16))
	v446 = v384 + (v385 + (v386 + (v387 + (v388 + (v389 + (v390 + (v391 + (v392 + (v393 + (v394 + (v395 + (v396 + (v397 + (v398 + (v399 + (v400 + (v401 + (v402 + (v403 + (v404 + (v405 + (v406 + (v407 + (v408 + (v409 + (v410 + (v411 + (v412 + (v413 + (v414 + v382))))))))))))))))))))))))))))))
	goto L66
L65:
	;
	v446 = v382
	goto L66
L66:
	;
	goto L63
L67:
	;
	v452 = v25 + int32(92)
	F_hash_seq_init(m, v452, v49)
	mBase = m.M
	v454 = m.ExcPending
	if v454 != 0 {
		goto L1
	} else {
		goto L68
	}
L68:
	;
	v457 = base.I32_div_s(v351*int32(9), v42)
	v458 = F_hash_seq_search(m, v452)
	mBase = m.M
	v459 = m.ExcPending
	if v459 != 0 {
		goto L1
	} else {
		goto L70
	}
L69:
	;
	v535 = F_errstart(m, int32(12), int32(0))
	mBase = m.M
	v536 = m.ExcPending
	if v536 != 0 {
		goto L1
	} else {
		goto L87
	}
L70:
	;
	if v458 == int32(0) {
		goto L71
	} else {
		goto L72
	}
L71:
	;
	v511 = v351
	v516 = int32(0)
	v532 = float64(0)
	goto L69
L72:
	;
	goto L73
L73:
	;
	v464 = int32(0)
	v467 = v351
	v470 = v458
	v472 = v464
	v473 = v464
	goto L74
L74:
	;
	v488 = *(*int32)(unsafe.Add(mBase, uint32(v470)+8))
	if v457 < v488 {
		goto L76
	} else {
		goto L77
	}
L75:
	;
	v511 = v501
	v516 = v503
	v532 = base.F64_convert_i32_u(v504)
	goto L69
L76:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v449+v472<<(uint(int32(2))%32)))) = v470
	v494 = *(*int32)(unsafe.Add(mBase, uint32(v470)+8))
	if v494 < v473 {
		goto L79
	} else {
		goto L80
	}
L77:
	;
	v501 = v467
	v503 = v472
	v504 = v473
	goto L78
L78:
	;
	v507 = F_hash_seq_search(m, v25+int32(92))
	mBase = m.M
	v508 = m.ExcPending
	if v508 != 0 {
		goto L1
	} else {
		goto L85
	}
L79:
	;
	v496 = v473
	goto L81
L80:
	;
	v496 = v494
	goto L81
L81:
	;
	if v467 < v494 {
		goto L82
	} else {
		goto L83
	}
L82:
	;
	v500 = v467
	goto L84
L83:
	;
	v500 = v494
	goto L84
L84:
	;
	v501 = v500
	v503 = v472 + int32(1)
	v504 = v496
	goto L78
L85:
	;
	if v507 != 0 {
		v467 = v501
		v470 = v507
		v472 = v503
		v473 = v504
		goto L74
	} else {
		goto L86
	}
L86:
	;
	goto L75
L87:
	;
	if v535 != 0 {
		goto L88
	} else {
		goto L89
	}
L88:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v25)+16)) = v516
	*(*int32)(unsafe.Add(mBase, uint32(v25)+12)) = v446
	*(*int32)(unsafe.Add(mBase, uint32(v25)+8)) = v351
	*(*int32)(unsafe.Add(mBase, uint32(v25)+4)) = v42
	*(*int32)(unsafe.Add(mBase, uint32(v25))) = v45
	F_errmsg_internal(m, int32(_a_F_compute_tsvector_stats_2), v25)
	mBase = m.M
	v544 = m.ExcPending
	if v544 != 0 {
		goto L1
	} else {
		goto L91
	}
L89:
	;
	goto L90
L90:
	;
	if v516 <= v45 {
		goto L94
	} else {
		goto L95
	}
L91:
	;
	F_errfinish(m, int32(_a_F_compute_tsvector_stats_3), int32(343), int32(_a_F_compute_tsvector_stats_4))
	mBase = m.M
	v549 = m.ExcPending
	if v549 != 0 {
		goto L1
	} else {
		goto L92
	}
L92:
	;
	goto L90
L93:
	;
	if v564 <= int32(0) {
		goto L59
	} else {
		goto L98
	}
L94:
	;
	v563 = v511
	v564 = v516
	goto L93
L95:
	;
	goto L96
L96:
	;
	F_qsort_interruptible(m, v449, v516, int32(4), int32(1168), int32(0))
	mBase = m.M
	v555 = m.ExcPending
	if v555 != 0 {
		goto L1
	} else {
		goto L97
	}
L97:
	;
	v561 = *(*int32)(unsafe.Add(mBase, uint32(v449+v45<<(uint(int32(2))%32)-int32(4))))
	v562 = *(*int32)(unsafe.Add(mBase, uint32(v561)+8))
	v563 = v562
	v564 = v45
	goto L93
L98:
	;
	v567 = int32(0)
	F_qsort_interruptible(m, v449, v564, int32(4), int32(1169), v567)
	mBase = m.M
	v572 = m.ExcPending
	if v572 != 0 {
		goto L1
	} else {
		goto L99
	}
L99:
	;
	v573 = int32(_a_F_compute_tsvector_stats_5)
	v574 = *(*int32)(unsafe.Add(mBase, _c_F_compute_tsvector_stats[0]))
	v576 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	*(*int32)(unsafe.Add(mBase, _c_F_compute_tsvector_stats[0])) = v576
	v580 = F_palloc(m, v564<<(uint(int32(2))%32))
	mBase = m.M
	v581 = m.ExcPending
	if v581 != 0 {
		goto L1
	} else {
		goto L100
	}
L100:
	;
	v582 = int32(2)
	v583 = v564 + v582
	v586 = F_palloc(m, v583<<(uint(v582)%32))
	mBase = m.M
	v587 = m.ExcPending
	if v587 != 0 {
		goto L1
	} else {
		goto L101
	}
L101:
	;
	v592 = v567
	goto L102
L102:
	;
	v611 = v592 << (uint(int32(2)) % 32)
	v614 = *(*int32)(unsafe.Add(mBase, uint32(v449+v611)))
	v615 = *(*int32)(unsafe.Add(mBase, uint32(v614)))
	v616 = *(*int32)(unsafe.Add(mBase, uint32(v614)+4))
	v617 = F_cstring_to_text_with_len(m, v615, v616)
	mBase = m.M
	v618 = m.ExcPending
	if v618 != 0 {
		goto L1
	} else {
		goto L104
	}
L103:
	;
	v631 = v586 + v564<<(uint(int32(2))%32)
	*(*float32)(unsafe.Add(mBase, uint32(v631)+4)) = base.F32_demote_f64(base.F64_div(v532, v372))
	*(*float32)(unsafe.Add(mBase, uint32(v631))) = base.F32_demote_f64(base.F64_div(base.F64_convert_i32_s(v563), v372))
	*(*int32)(unsafe.Add(mBase, _c_F_compute_tsvector_stats[0])) = v574
	*(*int32)(unsafe.Add(mBase, uint32(l0)+124)) = v586
	*(*int32)(unsafe.Add(mBase, uint32(l0)+84)) = int32(100)
	*(*int32)(unsafe.Add(mBase, uint32(l0)+64)) = int32(98)
	v646 = int32(4)
	*(*uint16)(unsafe.Add(mBase, uint32(l0)+52)) = uint16(v646)
	*(*int32)(unsafe.Add(mBase, uint32(l0)+164)) = v580
	*(*int32)(unsafe.Add(mBase, uint32(l0)+104)) = v583
	v650 = int32(105)
	*(*uint8)(unsafe.Add(mBase, uint32(l0)+219)) = uint8(v650)
	v652 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(l0)+214)) = uint8(v652)
	v654 = int32(_a_F_compute_tsvector_stats_6)
	*(*uint16)(unsafe.Add(mBase, uint32(l0)+204)) = uint16(v654)
	*(*int32)(unsafe.Add(mBase, uint32(l0)+184)) = int32(25)
	*(*int32)(unsafe.Add(mBase, uint32(l0)+144)) = v564
	goto L59
L104:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v580+v611))) = v617
	v621 = *(*int32)(unsafe.Add(mBase, uint32(v614)+8))
	*(*float32)(unsafe.Add(mBase, uint32(v611+v586))) = base.F32_demote_f64(base.F64_div(base.F64_convert_i32_s(v621), v372))
	v627 = v592 + int32(1)
	if v627 != v564 {
		v592 = v627
		goto L102
	} else {
		goto L105
	}
L105:
	;
	goto L103
L106:
	;
	F_errmsg_internal(m, int32(_a_F_compute_tsvector_stats_7), int32(0))
	mBase = m.M
	v697 = m.ExcPending
	if v697 != 0 {
		goto L1
	} else {
		goto L107
	}
L107:
	;
	F_errfinish(m, int32(_a_F_compute_tsvector_stats_3), int32(467), int32(_a_F_compute_tsvector_stats_8))
	mBase = m.M
	v702 = m.ExcPending
	if v702 != 0 {
		goto L1
	} else {
		goto L108
	}
L108:
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
	v79 = int32(2)
	v82 = *(*int32)(unsafe.Add(mBase, uint32(v78+v61<<(uint(v79)%32))))
	v83 = int32(4)
	v84 = v82 + v83
	v85 = *(*int32)(unsafe.Add(mBase, uint32(v82)))
	v89 = int32(base.Ui32(v85)>>(uint(v79)%32)) - v83
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
	v136 = v51 + v75<<(uint(v79)%32) + int32(base.Ui32(v122)>>(uint(int32(12))%32))
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
	v273 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
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
	v277 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
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
	return v269
L66:
	;
	goto L65
}
func F_tsvector_gt(m *base.Module, l0 int32) int32 {
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
	v5 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
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
	return int32(0)
L2:
	;
	v10 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
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
	v221 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
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
	v225 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
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
	return base.B2i32(int32(0) < v220)
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
	var v76 int32
	_ = v76
	var v77 int32
	_ = v77
	var v81 int32
	_ = v81
	var v82 int32
	_ = v82
	var v95 int32
	_ = v95
	var v106 int32
	_ = v106
	var v107 int32
	_ = v107
	var v110 int32
	_ = v110
	var v111 int32
	_ = v111
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
	var v126 int32
	_ = v126
	var v130 int32
	_ = v130
	var v133 int32
	_ = v133
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
	var v162 int32
	_ = v162
	var v167 int32
	_ = v167
	var v170 int32
	_ = v170
	var v172 int32
	_ = v172
	var v180 int32
	_ = v180
	var v181 int32
	_ = v181
	var v182 int32
	_ = v182
	var v185 int32
	_ = v185
	var v186 int32
	_ = v186
	var v188 int32
	_ = v188
	var v189 int32
	_ = v189
	var v191 int32
	_ = v191
	var v193 int32
	_ = v193
	var v196 int32
	_ = v196
	var v197 int32
	_ = v197
	var v198 int32
	_ = v198
	var v203 int32
	_ = v203
	var v204 int32
	_ = v204
	var v205 int32
	_ = v205
	var v208 int32
	_ = v208
	var v209 int32
	_ = v209
	var v210 int32
	_ = v210
	var v213 int32
	_ = v213
	var v214 int32
	_ = v214
	var v216 int32
	_ = v216
	var v221 int32
	_ = v221
	var v234 int32
	_ = v234
	var v237 int32
	_ = v237
	var v240 int32
	_ = v240
	var v246 int32
	_ = v246
	var v247 int32
	_ = v247
	var v263 int32
	_ = v263
	var v264 int32
	_ = v264
	var v269 int32
	_ = v269
	var v271 int32
	_ = v271
	var v275 int32
	_ = v275
	var v277 int32
	_ = v277
	var v292 int32
	_ = v292
	var v295 int32
	_ = v295
	var v297 int32
	_ = v297
	var v298 int32
	_ = v298
	var v300 int32
	_ = v300
	var v302 int32
	_ = v302
	var v305 int32
	_ = v305
	var v309 int32
	_ = v309
	var v329 int32
	_ = v329
	var v333 int32
	_ = v333
	var v350 int32
	_ = v350
	var v351 int32
	_ = v351
	var v353 int32
	_ = v353
	var v355 int32
	_ = v355
	var v358 int32
	_ = v358
	var v360 int32
	_ = v360
	var v363 int32
	_ = v363
	var v365 int32
	_ = v365
	var v368 int32
	_ = v368
	var v373 int32
	_ = v373
	var v397 int32
	_ = v397
	var v421 int32
	_ = v421
	var v424 int32
	_ = v424
	var v425 int32
	_ = v425
	var v428 int32
	_ = v428
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
		v58 = int32(_a_F_tsvector_setweight_by_filter_0)
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
	v58 = int32(_a_F_tsvector_setweight_by_filter_1)
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
	v58 = int32(_a_F_tsvector_setweight_by_filter_2)
	goto L4
L9:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v25))) = base.I32_extend8_s(v32)
	F_errmsg_internal(m, int32(_a_F_tsvector_setweight_by_filter_3), v25)
	mBase = m.M
	v51 = m.ExcPending
	if v51 != 0 {
		goto L1
	} else {
		goto L10
	}
L10:
	;
	F_errfinish(m, int32(_a_F_tsvector_setweight_by_filter_4), int32(308), int32(_a_F_tsvector_setweight_by_filter_5))
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
		goto L13
	} else {
		goto L14
	}
L13:
	;
	base.MemoryCopy(m, v62, v28, v66)
	goto L15
L14:
	;
	goto L15
L15:
	;
	F_deconstruct_array_builtin(m, v34, int32(25), v25+int32(8), v25+int32(4), v25+int32(12))
	mBase = m.M
	v76 = m.ExcPending
	if v76 != 0 {
		goto L1
	} else {
		goto L16
	}
L16:
	;
	v77 = *(*int32)(unsafe.Add(mBase, uint32(v25)+12))
	if int32(0) < v77 {
		goto L17
	} else {
		goto L18
	}
L17:
	;
	v81 = v62 + int32(8)
	v82 = *(*int32)(unsafe.Add(mBase, uint32(v25)+4))
	v95 = int32(0)
	goto L20
L18:
	;
	goto L19
L19:
	;
	v421 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	if v421 != v28 {
		goto L79
	} else {
		goto L80
	}
L20:
	;
	v106 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v95+v82))))
	if v106 != 0 {
		goto L22
	} else {
		goto L23
	}
L21:
	;
	goto L19
L22:
	;
	v397 = v95 + int32(1)
	if v397 < v77 {
		v95 = v397
		goto L20
	} else {
		goto L78
	}
L23:
	;
	v107 = *(*int32)(unsafe.Add(mBase, uint32(v62)+4))
	if v107 <= int32(0) {
		goto L22
	} else {
		goto L24
	}
L24:
	;
	v110 = *(*int32)(unsafe.Add(mBase, uint32(v25)+8))
	v111 = int32(2)
	v114 = *(*int32)(unsafe.Add(mBase, uint32(v110+v95<<(uint(v111)%32))))
	v115 = int32(4)
	v116 = v114 + v115
	v117 = *(*int32)(unsafe.Add(mBase, uint32(v114)))
	v121 = int32(base.Ui32(v117)>>(uint(v111)%32)) - v115
	v126 = v81 + v107<<(uint(v111)%32)
	v130 = v107
	v133 = int32(0)
	goto L25
L25:
	;
	v150 = v130 + v133
	v151 = int32(2)
	v152 = base.I32_div_s(v150, v151)
	v156 = *(*int32)(unsafe.Add(mBase, uint32(v81+v152<<(uint(v151)%32))))
	v160 = int32(base.Ui32(v156)>>(uint(int32(1))%32)) & int32(2047)
	v162 = int32(base.Ui32(v156) >> (uint(int32(12)) % 32))
	if v121 == int32(0) {
		goto L32
	} else {
		goto L33
	}
L26:
	;
	if base.B2i32(v156&int32(1) == int32(0))|base.B2i32(v150 < int32(-1)) != 0 {
		goto L22
	} else {
		goto L66
	}
L27:
	;
	goto L26
L28:
	;
	if v247 < v246 {
		v130 = v246
		v133 = v247
		goto L25
	} else {
		goto L65
	}
L29:
	;
	v246 = v130
	v247 = v152 + int32(1)
	goto L28
L30:
	;
	if v240 == int32(0) {
		goto L27
	} else {
		goto L64
	}
L31:
	;
	if int32(0) <= v237 {
		v240 = v237
		goto L30
	} else {
		goto L63
	}
L32:
	;
	if v160 != 0 {
		goto L35
	} else {
		goto L36
	}
L33:
	;
	goto L34
L34:
	;
	if v160 == int32(0) {
		v240 = base.B2i32(base.Ui32(int32(19)) < base.Ui32(v117))
		goto L30
	} else {
		goto L38
	}
L35:
	;
	v167 = int32(-1)
	goto L37
L36:
	;
	v167 = int32(0)
	goto L37
L37:
	;
	v237 = v167
	goto L31
L38:
	;
	v170 = v126 + v162
	if base.Ui32(v121) < base.Ui32(v160) {
		goto L39
	} else {
		goto L40
	}
L39:
	;
	v172 = v121
	goto L41
L40:
	;
	v172 = v160
	goto L41
L41:
	;
	if base.Ui32(int32(4)) <= base.Ui32(v172) {
		goto L45
	} else {
		goto L46
	}
L42:
	;
	if v234 != 0 {
		v237 = v234
		goto L31
	} else {
		goto L60
	}
L43:
	;
	v234 = int32(0)
	goto L42
L44:
	;
	v208 = v203
	v209 = v204
	v210 = v205
	goto L54
L45:
	;
	if (v116|v170)&int32(3) != 0 {
		v203 = v116
		v204 = v170
		v205 = v172
		goto L44
	} else {
		goto L48
	}
L46:
	;
	v196 = v116
	v197 = v170
	v198 = v172
	goto L47
L47:
	;
	if v198 == int32(0) {
		goto L43
	} else {
		goto L53
	}
L48:
	;
	v180 = v116
	v181 = v170
	v182 = v172
	goto L49
L49:
	;
	v185 = *(*int32)(unsafe.Add(mBase, uint32(v180)))
	v186 = *(*int32)(unsafe.Add(mBase, uint32(v181)))
	if v185 != v186 {
		v203 = v180
		v204 = v181
		v205 = v182
		goto L44
	} else {
		goto L51
	}
L50:
	;
	v196 = v191
	v197 = v189
	v198 = v193
	goto L47
L51:
	;
	v188 = int32(4)
	v189 = v181 + v188
	v191 = v180 + v188
	v193 = v182 - v188
	if base.Ui32(int32(3)) < base.Ui32(v193) {
		v180 = v191
		v181 = v189
		v182 = v193
		goto L49
	} else {
		goto L52
	}
L52:
	;
	goto L50
L53:
	;
	v203 = v196
	v204 = v197
	v205 = v198
	goto L44
L54:
	;
	v213 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v208))))
	v214 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v209))))
	if v213 == v214 {
		goto L56
	} else {
		goto L57
	}
L55:
	;
	v234 = v213 - v214
	goto L42
L56:
	;
	v216 = int32(1)
	v221 = v210 - v216
	if v221 != 0 {
		v208 = v208 + v216
		v209 = v209 + v216
		v210 = v221
		goto L54
	} else {
		goto L59
	}
L57:
	;
	goto L58
L58:
	;
	goto L55
L59:
	;
	goto L43
L60:
	;
	if v160 == v121 {
		goto L27
	} else {
		goto L61
	}
L61:
	;
	if v160 <= v121 {
		goto L29
	} else {
		goto L62
	}
L62:
	;
	v246 = v152
	v247 = v133
	goto L28
L63:
	;
	v246 = v152
	v247 = v133
	goto L28
L64:
	;
	goto L29
L65:
	;
	goto L22
L66:
	;
	v263 = v126 + (v160+v162+int32(1))&int32(_a_F_tsvector_setweight_by_filter_6)
	v264 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v263))))
	if v264 == int32(0) {
		goto L22
	} else {
		goto L67
	}
L67:
	;
	v269 = v264 & int32(3)
	if v269 != 0 {
		goto L68
	} else {
		goto L69
	}
L68:
	;
	v271 = v263
	v275 = v264
	v277 = int32(0)
	goto L71
L69:
	;
	v305 = v263
	v309 = v264
	goto L70
L70:
	;
	if base.Ui32(v264) < base.Ui32(int32(4)) {
		goto L22
	} else {
		goto L74
	}
L71:
	;
	v292 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v271)+2)))
	v295 = v292&int32(_a_F_tsvector_setweight_by_filter_7) | v58
	*(*uint16)(unsafe.Add(mBase, uint32(v271)+2)) = uint16(v295)
	v297 = int32(1)
	v298 = v275 - v297
	v300 = v271 + int32(2)
	v302 = v277 + v297
	if v302 != v269 {
		v271 = v300
		v275 = v298
		v277 = v302
		goto L71
	} else {
		goto L73
	}
L72:
	;
	v305 = v300
	v309 = v298
	goto L70
L73:
	;
	goto L72
L74:
	;
	v329 = v305
	v333 = v309
	goto L75
L75:
	;
	v350 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v329)+2)))
	v351 = int32(_a_F_tsvector_setweight_by_filter_7)
	v353 = v350&v351 | v58
	*(*uint16)(unsafe.Add(mBase, uint32(v329)+2)) = uint16(v353)
	v355 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v329)+4)))
	v358 = v355&v351 | v58
	*(*uint16)(unsafe.Add(mBase, uint32(v329)+4)) = uint16(v358)
	v360 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v329)+6)))
	v363 = v360&v351 | v58
	*(*uint16)(unsafe.Add(mBase, uint32(v329)+6)) = uint16(v363)
	v365 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v329)+8)))
	v368 = v365&v351 | v58
	*(*uint16)(unsafe.Add(mBase, uint32(v329)+8)) = uint16(v368)
	v373 = v333 - int32(4)
	if v373 != 0 {
		v329 = v329 + int32(8)
		v333 = v373
		goto L75
	} else {
		goto L77
	}
L76:
	;
	goto L22
L77:
	;
	goto L76
L78:
	;
	goto L21
L79:
	;
	F_pfree(m, v28)
	mBase = m.M
	v424 = m.ExcPending
	if v424 != 0 {
		goto L1
	} else {
		goto L82
	}
L80:
	;
	goto L81
L81:
	;
	v425 = *(*int32)(unsafe.Add(mBase, uint32(l0)+36))
	if v425 != v34 {
		goto L83
	} else {
		goto L84
	}
L82:
	;
	goto L81
L83:
	;
	F_pfree(m, v34)
	mBase = m.M
	v428 = m.ExcPending
	if v428 != 0 {
		goto L1
	} else {
		goto L86
	}
L84:
	;
	goto L85
L85:
	;
	m.G0 = v25 + int32(16)
	return v62
L86:
	;
	goto L85
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
