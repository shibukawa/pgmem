package p3

import (
	base "github.com/shibukawa/pgmem/internal/aot/pgaot/base"
	"unsafe"
)

func F_CLOGPagePrecedes(m *base.Module, l0 int64, l1 int64) int32 {
	var v5 int32
	_ = v5
	var v10 int32
	_ = v10
	v5 = int32(15)
	v10 = base.I32_wrap_i64(l0)<<(uint(v5)%32) - base.I32_wrap_i64(l1)<<(uint(v5)%32)
	return int32(base.Ui32(v10&(v10-int32(_a_F_CLOGPagePrecedes_0))) >> (uint(int32(31)) % 32))
}
func F_ClosePipeToProgram(m *base.Module, l0 int32) {
	mBase := m.M
	_ = mBase
	var v4 int32
	_ = v4
	var v6 int32
	_ = v6
	var v8 int32
	_ = v8
	var v9 int32
	_ = v9
	var v10 int32
	_ = v10
	var v16 int32
	_ = v16
	var v19 int32
	_ = v19
	var v20 int32
	_ = v20
	var v26 int32
	_ = v26
	var v27 int32
	_ = v27
	var v28 int32
	_ = v28
	var v32 int32
	_ = v32
	var v37 int32
	_ = v37
	var v44 int32
	_ = v44
	var v46 int32
	_ = v46
	var v50 int32
	_ = v50
	var v55 int32
	_ = v55
	v4 = m.G0
	v6 = v4 - int32(32)
	m.G0 = v6
	v8 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	v9 = F_ClosePipeStream(m, v8)
	mBase = m.M
	v10 = m.ExcPending
	if v10 != 0 {
		return
	} else {
		if v9 != 0 {
			if v9 == int32(-1) {
				F_errstart_cold(m, int32(21), int32(0))
				mBase = m.M
				v44 = m.ExcPending
				if v44 != 0 {
					return
				} else {
					F_errcode_for_file_access(m)
					mBase = m.M
					v46 = m.ExcPending
					if v46 != 0 {
						return
					} else {
						F_errmsg(m, int32(_a_F_ClosePipeToProgram_0), int32(0))
						mBase = m.M
						v50 = m.ExcPending
						if v50 != 0 {
							return
						} else {
							F_errfinish(m, int32(_a_F_ClosePipeToProgram_1), int32(732), int32(_a_F_ClosePipeToProgram_2))
							mBase = m.M
							v55 = m.ExcPending
							if v55 != 0 {
								return
							} else {
								base.Wasm_trap_unreachable()
								for {
								}
							}
						}
					}
				}
			} else {
				F_errstart_cold(m, int32(21), int32(0))
				mBase = m.M
				v16 = m.ExcPending
				if v16 != 0 {
					return
				} else {
					F_errcode(m, int32(515))
					mBase = m.M
					v19 = m.ExcPending
					if v19 != 0 {
						return
					} else {
						v20 = *(*int32)(unsafe.Add(mBase, uint32(l0)+36))
						*(*int32)(unsafe.Add(mBase, uint32(v6)+16)) = v20
						F_errmsg(m, int32(_a_F_ClosePipeToProgram_3), v6+int32(16))
						mBase = m.M
						v26 = m.ExcPending
						if v26 != 0 {
							return
						} else {
							v27 = F_wait_result_to_str(m, v9)
							mBase = m.M
							v28 = m.ExcPending
							if v28 != 0 {
								return
							} else {
								*(*int32)(unsafe.Add(mBase, uint32(v6))) = v27
								F_errdetail_internal(m, int32(_a_F_ClosePipeToProgram_4), v6)
								mBase = m.M
								v32 = m.ExcPending
								if v32 != 0 {
									return
								} else {
									F_errfinish(m, int32(_a_F_ClosePipeToProgram_1), int32(739), int32(_a_F_ClosePipeToProgram_2))
									mBase = m.M
									v37 = m.ExcPending
									if v37 != 0 {
										return
									} else {
										base.Wasm_trap_unreachable()
										for {
										}
									}
								}
							}
						}
					}
				}
			}
		} else {
			m.G0 = v6 + int32(32)
			return
		}
	}
}
func F_clause_selectivity_ext(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32, l5 int32) float64 {
	mBase := m.M
	_ = mBase
	var v11 int32
	_ = v11
	var v22 int32
	_ = v22
	var v25 int32
	_ = v25
	var v28 int32
	_ = v28
	var v29 int32
	_ = v29
	var v36 int32
	_ = v36
	var v37 int32
	_ = v37
	var v38 int32
	_ = v38
	var v41 int32
	_ = v41
	var v46 float64
	_ = v46
	var v51 float64
	_ = v51
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
	var v65 int32
	_ = v65
	var v66 int32
	_ = v66
	var v67 int32
	_ = v67
	var v68 int32
	_ = v68
	var v72 float64
	_ = v72
	var v73 int32
	_ = v73
	var v76 int32
	_ = v76
	var v78 int32
	_ = v78
	var v82 int64
	_ = v82
	var v85 float64
	_ = v85
	var v86 int32
	_ = v86
	var v87 int32
	_ = v87
	var v88 int32
	_ = v88
	var v92 int32
	_ = v92
	var v96 int64
	_ = v96
	var v99 float64
	_ = v99
	var v100 int32
	_ = v100
	var v102 int32
	_ = v102
	var v103 int32
	_ = v103
	var v104 int32
	_ = v104
	var v105 float64
	_ = v105
	var v106 int32
	_ = v106
	var v108 int32
	_ = v108
	var v109 float64
	_ = v109
	var v110 int32
	_ = v110
	var v111 int32
	_ = v111
	var v112 float64
	_ = v112
	var v113 int32
	_ = v113
	var v115 int32
	_ = v115
	var v117 int32
	_ = v117
	var v121 int32
	_ = v121
	var v122 int32
	_ = v122
	var v126 int32
	_ = v126
	var v127 int32
	_ = v127
	var v133 float64
	_ = v133
	var v134 int32
	_ = v134
	var v135 float64
	_ = v135
	var v138 int32
	_ = v138
	var v144 int32
	_ = v144
	var v149 float64
	_ = v149
	var v154 int32
	_ = v154
	var v160 int32
	_ = v160
	var v162 int32
	_ = v162
	var v163 int32
	_ = v163
	var v164 int32
	_ = v164
	var v165 int32
	_ = v165
	var v171 int32
	_ = v171
	var v172 float64
	_ = v172
	var v173 int32
	_ = v173
	var v177 float64
	_ = v177
	var v180 int32
	_ = v180
	var v181 int32
	_ = v181
	var v189 float64
	_ = v189
	var v203 int32
	_ = v203
	var v209 int32
	_ = v209
	var v210 int32
	_ = v210
	var v213 int32
	_ = v213
	var v216 int32
	_ = v216
	var v217 int32
	_ = v217
	var v218 float64
	_ = v218
	var v219 int32
	_ = v219
	var v220 int32
	_ = v220
	var v221 int32
	_ = v221
	var v222 int32
	_ = v222
	var v223 int32
	_ = v223
	var v227 int32
	_ = v227
	var v230 int32
	_ = v230
	var v231 int32
	_ = v231
	var v234 int32
	_ = v234
	var v235 int32
	_ = v235
	var v237 int32
	_ = v237
	var v239 float64
	_ = v239
	var v240 int32
	_ = v240
	var v241 int32
	_ = v241
	var v258 int32
	_ = v258
	var v260 int64
	_ = v260
	var v261 int32
	_ = v261
	var v264 float64
	_ = v264
	var v275 int32
	_ = v275
	var v276 float64
	_ = v276
	var v280 int32
	_ = v280
	var v285 int32
	_ = v285
	var v287 float64
	_ = v287
	var v298 int32
	_ = v298
	var v301 float64
	_ = v301
	var v302 int32
	_ = v302
	var v303 int32
	_ = v303
	var v304 int32
	_ = v304
	var v307 int32
	_ = v307
	var v308 float64
	_ = v308
	var v309 int32
	_ = v309
	var v310 int32
	_ = v310
	var v312 int32
	_ = v312
	var v314 int32
	_ = v314
	var v315 int32
	_ = v315
	var v316 int32
	_ = v316
	var v317 int32
	_ = v317
	var v318 int32
	_ = v318
	var v319 int32
	_ = v319
	var v320 int32
	_ = v320
	var v321 int32
	_ = v321
	var v322 int32
	_ = v322
	var v324 int32
	_ = v324
	var v325 int32
	_ = v325
	var v326 int32
	_ = v326
	var v332 int32
	_ = v332
	var v333 int32
	_ = v333
	var v337 int32
	_ = v337
	var v338 int32
	_ = v338
	var v341 float64
	_ = v341
	var v342 int32
	_ = v342
	var v343 float64
	_ = v343
	var v344 int32
	_ = v344
	var v345 float64
	_ = v345
	var v349 int32
	_ = v349
	var v350 int32
	_ = v350
	var v351 float64
	_ = v351
	var v352 int32
	_ = v352
	var v353 int32
	_ = v353
	var v354 int32
	_ = v354
	var v355 int32
	_ = v355
	var v357 int32
	_ = v357
	var v362 int32
	_ = v362
	var v363 int32
	_ = v363
	var v364 int32
	_ = v364
	var v365 int32
	_ = v365
	var v367 float32
	_ = v367
	var v368 float64
	_ = v368
	var v374 int32
	_ = v374
	var v375 int32
	_ = v375
	var v378 int32
	_ = v378
	var v381 float64
	_ = v381
	var v383 int32
	_ = v383
	var v384 float32
	_ = v384
	var v385 float64
	_ = v385
	var v388 int32
	_ = v388
	var v389 int64
	_ = v389
	var v392 float64
	_ = v392
	var v393 float64
	_ = v393
	var v394 float64
	_ = v394
	var v400 int32
	_ = v400
	var v406 int32
	_ = v406
	var v411 int32
	_ = v411
	var v414 float64
	_ = v414
	var v418 int32
	_ = v418
	var v430 int32
	_ = v430
	var v436 int32
	_ = v436
	var v441 int32
	_ = v441
	var v447 float64
	_ = v447
	var v448 int32
	_ = v448
	var v453 int32
	_ = v453
	var v457 int32
	_ = v457
	var v462 int32
	_ = v462
	var v463 float64
	_ = v463
	var v464 int32
	_ = v464
	var v466 float64
	_ = v466
	var v470 int32
	_ = v470
	var v471 int32
	_ = v471
	var v473 int32
	_ = v473
	var v474 float64
	_ = v474
	var v484 float64
	_ = v484
	var v491 int32
	_ = v491
	var v492 int32
	_ = v492
	var v493 int32
	_ = v493
	var v494 float64
	_ = v494
	var v499 float64
	_ = v499
	var v500 int32
	_ = v500
	var v501 float64
	_ = v501
	var v502 int32
	_ = v502
	var v503 int32
	_ = v503
	var v504 float64
	_ = v504
	var v505 int32
	_ = v505
	var v506 int32
	_ = v506
	var v507 int32
	_ = v507
	var v508 float64
	_ = v508
	var v509 int32
	_ = v509
	var v510 float64
	_ = v510
	var v512 int32
	_ = v512
	var v515 float64
	_ = v515
	var v523 int32
	_ = v523
	var v525 int32
	_ = v525
	var v528 int32
	_ = v528
	var v529 int32
	_ = v529
	var v537 int32
	_ = v537
	var v540 float64
	_ = v540
	var v542 int32
	_ = v542
	var v547 float64
	_ = v547
	var v548 int32
	_ = v548
	var v549 int32
	_ = v549
	var v552 int32
	_ = v552
	var v554 int32
	_ = v554
	var v556 float64
	_ = v556
	var v566 float64
	_ = v566
	var v581 int32
	_ = v581
	var v590 float64
	_ = v590
	v11 = int32(0)
	if l1 == v11 {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	return float64(0.5)
L2:
	;
	goto L3
L3:
	;
	v22 = *(*int32)(unsafe.Add(mBase, uint32(l1)))
	if v22 != int32(320) {
		goto L6
	} else {
		goto L7
	}
L4:
	;
	return v590
L5:
	;
	switch v66 - int32(6) {
	case 0:
		goto L44
	case 1:
		goto L43
	case 2:
		goto L42
	default:
		goto L29
	case 9:
		goto L39
	case 11, 12:
		goto L40
	case 14:
		goto L38
	case 15:
		goto L41
	case 21:
		goto L33
	case 31:
		goto L37
	case 46:
		goto L36
	case 47:
		goto L35
	case 49:
		goto L32
	case 52:
		goto L34
	}
L6:
	;
	v65 = l1
	v66 = v22
	v67 = v11
	v68 = v11
	goto L5
L7:
	;
	goto L8
L8:
	;
	v25 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l1)+10)))
	if v25 != int32(1) {
		goto L9
	} else {
		goto L10
	}
L9:
	;
	if l2 == int32(0) {
		goto L13
	} else {
		goto L14
	}
L10:
	;
	v28 = *(*int32)(unsafe.Add(mBase, uint32(l1)+4))
	v29 = *(*int32)(unsafe.Add(mBase, uint32(v28)))
	if v29 == int32(7) {
		goto L9
	} else {
		goto L11
	}
L11:
	;
	return float64(1)
L12:
	;
	v58 = *(*int32)(unsafe.Add(mBase, uint32(l1)+52))
	if v58 == int32(0) {
		goto L25
	} else {
		goto L26
	}
L13:
	;
	if l3 == int32(0) {
		goto L20
	} else {
		goto L21
	}
L14:
	;
	v36 = *(*int32)(unsafe.Add(mBase, uint32(l1)+24))
	switch v36 {
	case 0:
		goto L13
	case 1:
		goto L15
	default:
		v57 = v11
		goto L12
	}
L15:
	;
	v37 = *(*int32)(unsafe.Add(mBase, uint32(l1)+28))
	v38 = F_bms_is_member(m, l2, v37)
	mBase = m.M
	v41 = m.ExcPending
	if v41 != 0 {
		goto L16
	} else {
		goto L17
	}
L16:
	;
	return float64(0)
L17:
	;
	if v38 == int32(0) {
		v57 = v11
		goto L12
	} else {
		goto L18
	}
L18:
	;
	goto L13
L19:
	;
	v57 = int32(1)
	goto L12
L20:
	;
	v46 = *(*float64)(unsafe.Add(mBase, uint32(l1)+80))
	if base.F64_ge(v46, float64(0)) == int32(0) {
		goto L19
	} else {
		goto L23
	}
L21:
	;
	goto L22
L22:
	;
	v51 = *(*float64)(unsafe.Add(mBase, uint32(l1)+88))
	if base.F64_ge(v51, float64(0)) != 0 {
		v590 = v51
		goto L4
	} else {
		goto L24
	}
L23:
	;
	v590 = v46
	goto L4
L24:
	;
	goto L19
L25:
	;
	v61 = *(*int32)(unsafe.Add(mBase, uint32(l1)+4))
	v62 = v61
	goto L27
L26:
	;
	v62 = v58
	goto L27
L27:
	;
	v63 = *(*int32)(unsafe.Add(mBase, uint32(v62)))
	v65 = v62
	v66 = v63
	v67 = l1
	v68 = v57
	goto L5
L28:
	;
	if v68 == int32(0) {
		v590 = v566
		goto L4
	} else {
		goto L199
	}
L29:
	;
	v523 = m.G0
	v525 = v523 - int32(32)
	m.G0 = v525
	F_examine_variable(m, l0, v65, l2, v525)
	mBase = m.M
	v528 = m.ExcPending
	if v528 != 0 {
		goto L16
	} else {
		goto L185
	}
L30:
	;
	v512 = *(*int32)(unsafe.Add(mBase, uint32(v65)))
	if v512 == int32(18) {
		goto L182
	} else {
		goto L183
	}
L31:
	;
	v506 = *(*int32)(unsafe.Add(mBase, uint32(v65)+28))
	v507 = *(*int32)(unsafe.Add(mBase, uint32(v65)+24))
	v508 = F_restriction_selectivity(m, l0, v203, v506, v507, l2)
	mBase = m.M
	v509 = m.ExcPending
	if v509 != 0 {
		goto L16
	} else {
		goto L181
	}
L32:
	;
	v503 = *(*int32)(unsafe.Add(mBase, uint32(v65)+4))
	v504 = F_clause_selectivity_ext(m, l0, v503, l2, l3, l4, l5)
	mBase = m.M
	v505 = m.ExcPending
	if v505 != 0 {
		goto L16
	} else {
		goto L180
	}
L33:
	;
	v500 = *(*int32)(unsafe.Add(mBase, uint32(v65)+4))
	v501 = F_clause_selectivity_ext(m, l0, v500, l2, l3, l4, l5)
	mBase = m.M
	v502 = m.ExcPending
	if v502 != 0 {
		goto L16
	} else {
		goto L179
	}
L34:
	;
	v491 = *(*int32)(unsafe.Add(mBase, uint32(v65)+4))
	v492 = F_find_base_rel(m, l0, v491)
	mBase = m.M
	v493 = m.ExcPending
	if v493 != 0 {
		goto L16
	} else {
		goto L175
	}
L35:
	;
	v353 = *(*int32)(unsafe.Add(mBase, uint32(v65)+8))
	v354 = *(*int32)(unsafe.Add(mBase, uint32(v65)+4))
	v355 = m.G0
	v357 = v355 - int32(112)
	m.G0 = v357
	F_examine_variable(m, l0, v354, l2, v357+int32(80))
	mBase = m.M
	v362 = m.ExcPending
	if v362 != 0 {
		goto L16
	} else {
		goto L129
	}
L36:
	;
	v349 = *(*int32)(unsafe.Add(mBase, uint32(v65)+8))
	v350 = *(*int32)(unsafe.Add(mBase, uint32(v65)+4))
	v351 = F_nulltestsel(m, l0, v349, v350, l2)
	mBase = m.M
	v352 = m.ExcPending
	if v352 != 0 {
		goto L16
	} else {
		goto L128
	}
L37:
	;
	v310 = m.G0
	v312 = v310 - int32(16)
	m.G0 = v312
	v314 = *(*int32)(unsafe.Add(mBase, uint32(v65)+16))
	v315 = *(*int32)(unsafe.Add(mBase, uint32(v314)+12))
	v316 = *(*int32)(unsafe.Add(mBase, uint32(v315)))
	v317 = *(*int32)(unsafe.Add(mBase, uint32(v65)+8))
	v318 = *(*int32)(unsafe.Add(mBase, uint32(v317)+12))
	v319 = *(*int32)(unsafe.Add(mBase, uint32(v318)))
	v320 = *(*int32)(unsafe.Add(mBase, uint32(v65)+20))
	v321 = *(*int32)(unsafe.Add(mBase, uint32(v320)+12))
	v322 = *(*int32)(unsafe.Add(mBase, uint32(v321)))
	*(*int32)(unsafe.Add(mBase, uint32(v312)+12)) = v322
	v324 = *(*int32)(unsafe.Add(mBase, uint32(v65)+24))
	v325 = *(*int32)(unsafe.Add(mBase, uint32(v324)+12))
	v326 = *(*int32)(unsafe.Add(mBase, uint32(v325)))
	*(*int32)(unsafe.Add(mBase, uint32(v312)+8)) = v326
	*(*int32)(unsafe.Add(mBase, uint32(v312)+4)) = v322
	*(*int32)(unsafe.Add(mBase, uint32(v312))) = v326
	v332 = F_list_make2_impl(m, v312+int32(4), v312)
	mBase = m.M
	v333 = m.ExcPending
	if v333 != 0 {
		goto L16
	} else {
		goto L120
	}
L38:
	;
	if l2|base.B2i32(l4 == int32(0)) != 0 {
		goto L111
	} else {
		goto L112
	}
L39:
	;
	v220 = *(*int32)(unsafe.Add(mBase, uint32(v65)+24))
	v221 = *(*int32)(unsafe.Add(mBase, uint32(v65)+28))
	v222 = *(*int32)(unsafe.Add(mBase, uint32(v65)+4))
	v223 = int32(0)
	if l2|base.B2i32(l4 == v223) != 0 {
		v234 = v223
		goto L95
	} else {
		goto L96
	}
L40:
	;
	v203 = *(*int32)(unsafe.Add(mBase, uint32(v65)+4))
	if l2|base.B2i32(l4 == int32(0)) != 0 {
		goto L31
	} else {
		goto L86
	}
L41:
	;
	v100 = *(*int32)(unsafe.Add(mBase, uint32(v65)+4))
	switch v100 {
	case 0:
		goto L65
	case 1:
		goto L64
	case 2:
		goto L66
	default:
		goto L29
	}
L42:
	;
	v86 = F_estimate_expression_value(m, l0, v65)
	mBase = m.M
	v87 = m.ExcPending
	if v87 != 0 {
		goto L16
	} else {
		goto L54
	}
L43:
	;
	v78 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v65)+32)))
	if v78 != 0 {
		goto L48
	} else {
		goto L49
	}
L44:
	;
	v72 = float64(0.5)
	v73 = *(*int32)(unsafe.Add(mBase, uint32(v65)+28))
	if v73 != 0 {
		v566 = v72
		goto L28
	} else {
		goto L45
	}
L45:
	;
	if l2 == int32(0) {
		goto L29
	} else {
		goto L46
	}
L46:
	;
	v76 = *(*int32)(unsafe.Add(mBase, uint32(v65)+4))
	if l2 != v76 {
		v566 = v72
		goto L28
	} else {
		goto L47
	}
L47:
	;
	goto L29
L48:
	;
	v566 = float64(0)
	goto L28
L49:
	;
	goto L50
L50:
	;
	v82 = *(*int64)(unsafe.Add(mBase, uint32(v65)+24))
	if v82 == int64(0) {
		goto L51
	} else {
		goto L52
	}
L51:
	;
	v85 = float64(0)
	goto L53
L52:
	;
	v85 = float64(1)
	goto L53
L53:
	;
	v566 = v85
	goto L28
L54:
	;
	v88 = *(*int32)(unsafe.Add(mBase, uint32(v86)))
	if v88 != int32(7) {
		goto L55
	} else {
		goto L56
	}
L55:
	;
	v566 = float64(0.5)
	goto L28
L56:
	;
	goto L57
L57:
	;
	v92 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v86)+32)))
	if v92 != 0 {
		goto L58
	} else {
		goto L59
	}
L58:
	;
	v566 = float64(0)
	goto L28
L59:
	;
	goto L60
L60:
	;
	v96 = *(*int64)(unsafe.Add(mBase, uint32(v86)+24))
	if v96 == int64(0) {
		goto L61
	} else {
		goto L62
	}
L61:
	;
	v99 = float64(0)
	goto L63
L62:
	;
	v99 = float64(1)
	goto L63
L63:
	;
	v566 = v99
	goto L28
L64:
	;
	v111 = *(*int32)(unsafe.Add(mBase, uint32(v65)+8))
	v112 = float64(0)
	v113 = m.G0
	v115 = v113 - int32(16)
	m.G0 = v115
	v117 = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v115)+12)) = v117
	v121 = F_find_single_rel_for_clauses(m, l0, v111)
	mBase = m.M
	v122 = m.ExcPending
	if v122 != 0 {
		goto L16
	} else {
		goto L70
	}
L65:
	;
	v108 = *(*int32)(unsafe.Add(mBase, uint32(v65)+8))
	v109 = F_clauselist_selectivity_ext(m, l0, v108, l2, l3, l4, l5)
	mBase = m.M
	v110 = m.ExcPending
	if v110 != 0 {
		goto L16
	} else {
		goto L68
	}
L66:
	;
	v102 = *(*int32)(unsafe.Add(mBase, uint32(v65)+8))
	v103 = *(*int32)(unsafe.Add(mBase, uint32(v102)+12))
	v104 = *(*int32)(unsafe.Add(mBase, uint32(v103)))
	v105 = F_clause_selectivity_ext(m, l0, v104, l2, l3, l4, l5)
	mBase = m.M
	v106 = m.ExcPending
	if v106 != 0 {
		goto L16
	} else {
		goto L67
	}
L67:
	;
	v566 = base.F64_sub(float64(1), v105)
	goto L28
L68:
	;
	v566 = v109
	goto L28
L69:
	;
	if v111 == int32(0) {
		v189 = v135
		goto L75
	} else {
		goto L76
	}
L70:
	;
	if base.B2i32(l5 == v117)|base.B2i32(v121 == int32(0)) != 0 {
		v135 = v112
		goto L69
	} else {
		goto L71
	}
L71:
	;
	v126 = *(*int32)(unsafe.Add(mBase, uint32(v121)+84))
	if v126 != 0 {
		v135 = v112
		goto L69
	} else {
		goto L72
	}
L72:
	;
	v127 = *(*int32)(unsafe.Add(mBase, uint32(v121)+120))
	if v127 == int32(0) {
		v135 = v112
		goto L69
	} else {
		goto L73
	}
L73:
	;
	v133 = F_statext_clauselist_selectivity(m, l0, v111, l2, l3, l4, v121, v115+int32(12), int32(1))
	mBase = m.M
	v134 = m.ExcPending
	if v134 != 0 {
		goto L16
	} else {
		goto L74
	}
L74:
	;
	v135 = v133
	goto L69
L75:
	;
	m.G0 = v115 + int32(16)
	v566 = v189
	goto L28
L76:
	;
	v138 = *(*int32)(unsafe.Add(mBase, uint32(v111)+4))
	if v138 <= int32(0) {
		v189 = v135
		goto L75
	} else {
		goto L77
	}
L77:
	;
	v144 = int32(0)
	v149 = v135
	v154 = int32(-1)
	goto L78
L78:
	;
	v160 = *(*int32)(unsafe.Add(mBase, uint32(v111)+12))
	v162 = v154 + int32(1)
	v163 = *(*int32)(unsafe.Add(mBase, uint32(v115)+12))
	v164 = F_bms_is_member(m, v162, v163)
	mBase = m.M
	v165 = m.ExcPending
	if v165 != 0 {
		goto L16
	} else {
		goto L80
	}
L79:
	;
	v189 = v177
	goto L75
L80:
	;
	if v164 == int32(0) {
		goto L81
	} else {
		goto L82
	}
L81:
	;
	v171 = *(*int32)(unsafe.Add(mBase, uint32(v160+v144<<(uint(int32(2))%32))))
	v172 = F_clause_selectivity_ext(m, l0, v171, l2, l3, l4, l5)
	mBase = m.M
	v173 = m.ExcPending
	if v173 != 0 {
		goto L16
	} else {
		goto L84
	}
L82:
	;
	v177 = v149
	goto L83
L83:
	;
	v180 = v144 + int32(1)
	v181 = *(*int32)(unsafe.Add(mBase, uint32(v111)+4))
	if v180 < v181 {
		v144 = v180
		v149 = v177
		v154 = v162
		goto L78
	} else {
		goto L85
	}
L84:
	;
	v177 = base.F64_sub(base.F64_add(v149, v172), base.F64_mul(v149, v172))
	goto L83
L85:
	;
	goto L79
L86:
	;
	if v67 == int32(0) {
		goto L88
	} else {
		goto L89
	}
L87:
	;
	v216 = *(*int32)(unsafe.Add(mBase, uint32(v65)+28))
	v217 = *(*int32)(unsafe.Add(mBase, uint32(v65)+24))
	v218 = F_join_selectivity(m, l0, v203, v216, v217, l3, l4)
	mBase = m.M
	v219 = m.ExcPending
	if v219 != 0 {
		goto L16
	} else {
		goto L94
	}
L88:
	;
	v209 = F_NumRelids(m, l0, v65)
	mBase = m.M
	v210 = m.ExcPending
	if v210 != 0 {
		goto L16
	} else {
		goto L91
	}
L89:
	;
	goto L90
L90:
	;
	v213 = *(*int32)(unsafe.Add(mBase, uint32(v67)+24))
	if v213 < int32(2) {
		goto L31
	} else {
		goto L93
	}
L91:
	;
	if int32(1) < v209 {
		goto L87
	} else {
		goto L92
	}
L92:
	;
	goto L31
L93:
	;
	goto L87
L94:
	;
	v510 = v218
	goto L30
L95:
	;
	v235 = m.G0
	v237 = v235 + int32(-64)
	m.G0 = v237
	v239 = float64(-1)
	v240 = F_get_func_support(m, v222)
	mBase = m.M
	v241 = m.ExcPending
	if v241 != 0 {
		goto L16
	} else {
		goto L102
	}
L96:
	;
	if v67 != 0 {
		goto L97
	} else {
		goto L98
	}
L97:
	;
	v227 = *(*int32)(unsafe.Add(mBase, uint32(v67)+24))
	v234 = base.B2i32(int32(1) < v227)
	goto L95
L98:
	;
	goto L99
L99:
	;
	v230 = F_NumRelids(m, l0, v65)
	mBase = m.M
	v231 = m.ExcPending
	if v231 != 0 {
		goto L16
	} else {
		goto L100
	}
L100:
	;
	v234 = base.B2i32(int32(1) < v230)
	goto L95
L101:
	;
	m.G0 = v237 - int32(-64)
	if base.F64_lt(v287, float64(0)) != 0 {
		goto L29
	} else {
		goto L110
	}
L102:
	;
	if v240 == int32(0) {
		v287 = v239
		goto L101
	} else {
		goto L103
	}
L103:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v237)+56)) = int64(-4616189618054758400)
	*(*int32)(unsafe.Add(mBase, uint32(v237)+48)) = l4
	*(*int32)(unsafe.Add(mBase, uint32(v237)+44)) = l3
	*(*int32)(unsafe.Add(mBase, uint32(v237)+40)) = l2
	*(*uint8)(unsafe.Add(mBase, uint32(v237)+36)) = uint8(v234)
	*(*int32)(unsafe.Add(mBase, uint32(v237)+32)) = v220
	*(*int32)(unsafe.Add(mBase, uint32(v237)+28)) = v221
	*(*int32)(unsafe.Add(mBase, uint32(v237)+24)) = v222
	*(*int32)(unsafe.Add(mBase, uint32(v237)+20)) = l0
	*(*int32)(unsafe.Add(mBase, uint32(v237)+16)) = int32(466)
	v258 = v235 + int32(-48)
	v260 = F_OidFunctionCall1Coll(m, v240, int32(0), base.I64_extend_i32_u(v258))
	mBase = m.M
	v261 = m.ExcPending
	if v261 != 0 {
		goto L16
	} else {
		goto L104
	}
L104:
	;
	if base.I32_wrap_i64(v260) != v258 {
		v287 = v239
		goto L101
	} else {
		goto L105
	}
L105:
	;
	v264 = *(*float64)(unsafe.Add(mBase, uint32(v237)+56))
	if base.F64_lt(v264, float64(0))|base.F64_gt(v264, float64(1)) == int32(0) {
		v287 = v264
		goto L101
	} else {
		goto L106
	}
L106:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v275 = m.ExcPending
	if v275 != 0 {
		goto L16
	} else {
		goto L107
	}
L107:
	;
	v276 = *(*float64)(unsafe.Add(mBase, uint32(v237)+56))
	*(*float64)(unsafe.Add(mBase, uint32(v237))) = v276
	F_errmsg_internal(m, int32(_a_F_clause_selectivity_ext_0), v237)
	mBase = m.M
	v280 = m.ExcPending
	if v280 != 0 {
		goto L16
	} else {
		goto L108
	}
L108:
	;
	F_errfinish(m, int32(_a_F_clause_selectivity_ext_1), int32(2316), int32(_a_F_clause_selectivity_ext_2))
	mBase = m.M
	v285 = m.ExcPending
	if v285 != 0 {
		goto L16
	} else {
		goto L109
	}
L109:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L110:
	;
	v566 = v287
	goto L28
L111:
	;
	v307 = int32(0)
	goto L113
L112:
	;
	if v67 != 0 {
		goto L114
	} else {
		goto L115
	}
L113:
	;
	v308 = F_scalararraysel(m, l0, v65, v307, l2, l3, l4)
	mBase = m.M
	v309 = m.ExcPending
	if v309 != 0 {
		goto L16
	} else {
		goto L119
	}
L114:
	;
	v298 = *(*int32)(unsafe.Add(mBase, uint32(v67)+24))
	v301 = F_scalararraysel(m, l0, v65, base.B2i32(int32(1) < v298), l2, l3, l4)
	mBase = m.M
	v302 = m.ExcPending
	if v302 != 0 {
		goto L16
	} else {
		goto L117
	}
L115:
	;
	goto L116
L116:
	;
	v303 = F_NumRelids(m, l0, v65)
	mBase = m.M
	v304 = m.ExcPending
	if v304 != 0 {
		goto L16
	} else {
		goto L118
	}
L117:
	;
	v566 = v301
	goto L28
L118:
	;
	v307 = base.B2i32(int32(1) < v303)
	goto L113
L119:
	;
	v566 = v308
	goto L28
L120:
	;
	if l2|base.B2i32(l4 == int32(0)) != 0 {
		goto L122
	} else {
		goto L123
	}
L121:
	;
	m.G0 = v312 + int32(16)
	v566 = v345
	goto L28
L122:
	;
	v343 = F_restriction_selectivity(m, l0, v319, v332, v316, l2)
	mBase = m.M
	v344 = m.ExcPending
	if v344 != 0 {
		goto L16
	} else {
		goto L127
	}
L123:
	;
	v337 = F_NumRelids(m, l0, v332)
	mBase = m.M
	v338 = m.ExcPending
	if v338 != 0 {
		goto L16
	} else {
		goto L124
	}
L124:
	;
	if v337 < int32(2) {
		goto L122
	} else {
		goto L125
	}
L125:
	;
	v341 = F_join_selectivity(m, l0, v319, v332, v316, l3, l4)
	mBase = m.M
	v342 = m.ExcPending
	if v342 != 0 {
		goto L16
	} else {
		goto L126
	}
L126:
	;
	v345 = v341
	goto L121
L127:
	;
	v345 = v343
	goto L121
L128:
	;
	v566 = v351
	goto L28
L129:
	;
	v363 = *(*int32)(unsafe.Add(mBase, uint32(v357)+88))
	if v363 != 0 {
		goto L132
	} else {
		goto L133
	}
L130:
	;
	m.G0 = v357 + int32(112)
	v566 = v484
	goto L28
L131:
	;
	v470 = *(*int32)(unsafe.Add(mBase, uint32(v357)+88))
	if v470 != 0 {
		goto L169
	} else {
		goto L170
	}
L132:
	;
	v364 = *(*int32)(unsafe.Add(mBase, uint32(v363)+16))
	v365 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v364)+22)))
	v367 = *(*float32)(unsafe.Add(mBase, uint32(v364+v365)+8))
	v368 = base.F64_promote_f32(v367)
	v374 = F_get_attstatsslot(m, v357+int32(44), v363, int32(1), int32(0), int32(3))
	mBase = m.M
	v375 = m.ExcPending
	if v375 != 0 {
		goto L16
	} else {
		goto L136
	}
L133:
	;
	goto L134
L134:
	;
	switch v353 {
	case 0, 3:
		goto L160
	case 1, 2:
		goto L162
	case 4:
		v484 = float64(0.005)
		goto L130
	case 5:
		goto L163
	default:
		goto L161
	}
L135:
	;
	switch v353 {
	case 0, 2:
		goto L156
	case 1, 3:
		goto L155
	case 4:
		v466 = v368
		goto L131
	case 5:
		goto L153
	default:
		goto L154
	}
L136:
	;
	if v374 == int32(0) {
		goto L135
	} else {
		goto L137
	}
L137:
	;
	v378 = *(*int32)(unsafe.Add(mBase, uint32(v357)+68))
	if v378 <= int32(0) {
		goto L135
	} else {
		goto L138
	}
L138:
	;
	v381 = float64(1)
	v383 = *(*int32)(unsafe.Add(mBase, uint32(v357)+64))
	v384 = *(*float32)(unsafe.Add(mBase, uint32(v383)))
	v385 = base.F64_promote_f32(v384)
	v388 = *(*int32)(unsafe.Add(mBase, uint32(v357)+56))
	v389 = *(*int64)(unsafe.Add(mBase, uint32(v388)))
	if v389 == int64(0) {
		goto L139
	} else {
		goto L140
	}
L139:
	;
	v392 = base.F64_sub(base.F64_sub(v381, v385), v368)
	goto L141
L140:
	;
	v392 = v385
	goto L141
L141:
	;
	v393 = base.F64_sub(v381, v392)
	v394 = base.F64_sub(v393, v368)
	switch v353 {
	case 0:
		goto L148
	case 1:
		goto L147
	case 2:
		goto L146
	case 3:
		goto L145
	case 4:
		v414 = v368
		goto L142
	case 5:
		goto L143
	default:
		goto L144
	}
L142:
	;
	F_free_attstatsslot(m, v357+int32(44))
	mBase = m.M
	v418 = m.ExcPending
	if v418 != 0 {
		goto L16
	} else {
		goto L152
	}
L143:
	;
	v414 = base.F64_sub(float64(1), v368)
	goto L142
L144:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v400 = m.ExcPending
	if v400 != 0 {
		goto L16
	} else {
		goto L149
	}
L145:
	;
	v414 = base.F64_sub(float64(1), v394)
	goto L142
L146:
	;
	v414 = v394
	goto L142
L147:
	;
	v414 = v393
	goto L142
L148:
	;
	v414 = v392
	goto L142
L149:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v357)+16)) = v353
	F_errmsg_internal(m, int32(_a_F_clause_selectivity_ext_3), v357+int32(16))
	mBase = m.M
	v406 = m.ExcPending
	if v406 != 0 {
		goto L16
	} else {
		goto L150
	}
L150:
	;
	F_errfinish(m, int32(_a_F_clause_selectivity_ext_4), int32(1692), int32(_a_F_clause_selectivity_ext_5))
	mBase = m.M
	v411 = m.ExcPending
	if v411 != 0 {
		goto L16
	} else {
		goto L151
	}
L151:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L152:
	;
	v466 = v414
	goto L131
L153:
	;
	v466 = base.F64_sub(float64(1), v368)
	goto L131
L154:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v430 = m.ExcPending
	if v430 != 0 {
		goto L16
	} else {
		goto L157
	}
L155:
	;
	v466 = base.F64_mul(base.F64_add(v368, float64(1)), float64(0.5))
	goto L131
L156:
	;
	v466 = base.F64_mul(base.F64_sub(float64(1), v368), float64(0.5))
	goto L131
L157:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v357)+32)) = v353
	F_errmsg_internal(m, int32(_a_F_clause_selectivity_ext_3), v357+int32(32))
	mBase = m.M
	v436 = m.ExcPending
	if v436 != 0 {
		goto L16
	} else {
		goto L158
	}
L158:
	;
	F_errfinish(m, int32(_a_F_clause_selectivity_ext_4), int32(1729), int32(_a_F_clause_selectivity_ext_5))
	mBase = m.M
	v441 = m.ExcPending
	if v441 != 0 {
		goto L16
	} else {
		goto L159
	}
L159:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L160:
	;
	v463 = F_clause_selectivity(m, l0, v354, l2, l3, l4)
	mBase = m.M
	v464 = m.ExcPending
	if v464 != 0 {
		goto L16
	} else {
		goto L168
	}
L161:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v453 = m.ExcPending
	if v453 != 0 {
		goto L16
	} else {
		goto L165
	}
L162:
	;
	v447 = F_clause_selectivity(m, l0, v354, l2, l3, l4)
	mBase = m.M
	v448 = m.ExcPending
	if v448 != 0 {
		goto L16
	} else {
		goto L164
	}
L163:
	;
	v484 = float64(0.995)
	goto L130
L164:
	;
	v466 = base.F64_sub(float64(1), v447)
	goto L131
L165:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v357))) = v353
	F_errmsg_internal(m, int32(_a_F_clause_selectivity_ext_3), v357)
	mBase = m.M
	v457 = m.ExcPending
	if v457 != 0 {
		goto L16
	} else {
		goto L166
	}
L166:
	;
	F_errfinish(m, int32(_a_F_clause_selectivity_ext_4), int32(1765), int32(_a_F_clause_selectivity_ext_5))
	mBase = m.M
	v462 = m.ExcPending
	if v462 != 0 {
		goto L16
	} else {
		goto L167
	}
L167:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L168:
	;
	v466 = v463
	goto L131
L169:
	;
	v471 = *(*int32)(unsafe.Add(mBase, uint32(v357)+92))
	m.T0[v471].(func(*base.Module, int32))(m, v470)
	mBase = m.M
	v473 = m.ExcPending
	if v473 != 0 {
		goto L16
	} else {
		goto L172
	}
L170:
	;
	goto L171
L171:
	;
	v474 = float64(0)
	if base.F64_lt(v466, v474) != 0 {
		v484 = v474
		goto L130
	} else {
		goto L173
	}
L172:
	;
	goto L171
L173:
	;
	if base.F64_gt(v466, float64(1)) == int32(0) {
		v484 = v466
		goto L130
	} else {
		goto L174
	}
L174:
	;
	v484 = float64(1)
	goto L130
L175:
	;
	v494 = *(*float64)(unsafe.Add(mBase, uint32(v492)+128))
	if base.F64_gt(v494, float64(0)) != 0 {
		goto L176
	} else {
		goto L177
	}
L176:
	;
	v499 = base.F64_div(float64(1), v494)
	goto L178
L177:
	;
	v499 = float64(0.5)
	goto L178
L178:
	;
	v566 = v499
	goto L28
L179:
	;
	v566 = v501
	goto L28
L180:
	;
	v566 = v504
	goto L28
L181:
	;
	v510 = v508
	goto L30
L182:
	;
	v515 = base.F64_sub(float64(1), v510)
	goto L184
L183:
	;
	v515 = v510
	goto L184
L184:
	;
	v566 = v515
	goto L28
L185:
	;
	v529 = *(*int32)(unsafe.Add(mBase, uint32(v525)+8))
	if v529 == int32(0) {
		goto L187
	} else {
		goto L188
	}
L186:
	;
	m.G0 = v525 + int32(32)
	v566 = v556
	goto L28
L187:
	;
	if v65 == int32(0) {
		goto L190
	} else {
		goto L191
	}
L188:
	;
	goto L189
L189:
	;
	v542 = int32(0)
	v547 = F_var_eq_const(m, v525, int32(91), v542, int64(1), v542, int32(1), v542)
	mBase = m.M
	v548 = m.ExcPending
	if v548 != 0 {
		goto L16
	} else {
		goto L196
	}
L190:
	;
	v556 = float64(0.5)
	goto L186
L191:
	;
	goto L192
L192:
	;
	v537 = *(*int32)(unsafe.Add(mBase, uint32(v65)))
	if v537 == int32(15) {
		goto L193
	} else {
		goto L194
	}
L193:
	;
	v540 = float64(0.3333333)
	goto L195
L194:
	;
	v540 = float64(0.5)
	goto L195
L195:
	;
	v556 = v540
	goto L186
L196:
	;
	v549 = *(*int32)(unsafe.Add(mBase, uint32(v525)+8))
	if v549 == int32(0) {
		v556 = v547
		goto L186
	} else {
		goto L197
	}
L197:
	;
	v552 = *(*int32)(unsafe.Add(mBase, uint32(v525)+12))
	m.T0[v552].(func(*base.Module, int32))(m, v549)
	mBase = m.M
	v554 = m.ExcPending
	if v554 != 0 {
		goto L16
	} else {
		goto L198
	}
L198:
	;
	v556 = v547
	goto L186
L199:
	;
	if l3 != 0 {
		goto L200
	} else {
		goto L201
	}
L200:
	;
	v581 = int32(88)
	goto L202
L201:
	;
	v581 = int32(80)
	goto L202
L202:
	;
	*(*float64)(unsafe.Add(mBase, uint32(v67+v581))) = v566
	v590 = v566
	goto L4
}
func F_clean_NOT_intree(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v7 int32
	_ = v7
	var v8 int32
	_ = v8
	var v9 int32
	_ = v9
	var v13 int32
	_ = v13
	var v16 int32
	_ = v16
	var v17 int32
	_ = v17
	var v18 int32
	_ = v18
	var v22 int32
	_ = v22
	var v23 int32
	_ = v23
	var v24 int32
	_ = v24
	var v31 int32
	_ = v31
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
	var v40 int32
	_ = v40
	var v42 int32
	_ = v42
	var v47 int32
	_ = v47
	var v53 int32
	_ = v53
	var v57 int32
	_ = v57
	F_check_stack_depth(m)
	mBase = m.M
	v7 = m.ExcPending
	if v7 != 0 {
		return int32(0)
	} else {
		v8 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
		v9 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v8))))
		if v9 == int32(1) {
			return l0
		} else {
			v13 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v8)+1)))
			switch v13 - int32(1) {
			case 0:
				F_freetree(m, l0)
				mBase = m.M
				v31 = m.ExcPending
				if v31 != 0 {
					return int32(0)
				} else {
					return int32(0)
				}
			default:
				v34 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
				v35 = F_clean_NOT_intree(m, v34)
				mBase = m.M
				v36 = m.ExcPending
				if v36 != 0 {
					return int32(0)
				} else {
					*(*int32)(unsafe.Add(mBase, uint32(l0))) = v35
					v38 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
					v39 = F_clean_NOT_intree(m, v38)
					mBase = m.M
					v40 = m.ExcPending
					if v40 != 0 {
						return int32(0)
					} else {
						*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v39
						v42 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
						if v39|v42 == int32(0) {
							F_pfree(m, l0)
							mBase = m.M
							v47 = m.ExcPending
							if v47 != 0 {
								return int32(0)
							} else {
								return int32(0)
							}
						} else {
							if v42 == int32(0) {
								F_pfree(m, l0)
								mBase = m.M
								v53 = m.ExcPending
								if v53 != 0 {
									return int32(0)
								} else {
									return v39
								}
							} else {
								if v39 != 0 {
									return l0
								} else {
									F_pfree(m, l0)
									mBase = m.M
									v57 = m.ExcPending
									if v57 != 0 {
										return int32(0)
									} else {
										return v42
									}
								}
							}
						}
					}
				}
			case 2:
				v16 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
				v17 = F_clean_NOT_intree(m, v16)
				mBase = m.M
				v18 = m.ExcPending
				if v18 != 0 {
					return int32(0)
				} else {
					*(*int32)(unsafe.Add(mBase, uint32(l0))) = v17
					if v17 == int32(0) {
						F_freetree(m, l0)
						mBase = m.M
						v31 = m.ExcPending
						if v31 != 0 {
							return int32(0)
						} else {
							return int32(0)
						}
					} else {
						v22 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
						v23 = F_clean_NOT_intree(m, v22)
						mBase = m.M
						v24 = m.ExcPending
						if v24 != 0 {
							return int32(0)
						} else {
							*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v23
							if v23 == int32(0) {
								F_freetree(m, l0)
								mBase = m.M
								v31 = m.ExcPending
								if v31 != 0 {
									return int32(0)
								} else {
									return int32(0)
								}
							} else {
								return l0
							}
						}
					}
				}
			}
		}
	}
}
func F_clock_timestamp(m *base.Module, l0 int32) int64 {
	mBase := m.M
	_ = mBase
	var v4 int32
	_ = v4
	var v5 int32
	_ = v5
	var v6 int32
	_ = v6
	var v9 int64
	_ = v9
	var v10 int64
	_ = v10
	v4 = m.G0
	v5 = int32(16)
	v6 = v4 - v5
	m.G0 = v6
	F_gettimeofday(m, v6)
	mBase = m.M
	v9 = *(*int64)(unsafe.Add(mBase, uint32(v6)))
	v10 = int64(*(*int32)(unsafe.Add(mBase, uint32(v6)+8)))
	m.G0 = v6 + v5
	return v10 + v9*int64(1000000) - int64(946684800000000)
}
func F_clog_desc(m *base.Module, l0 int32, l1 int32) {
	mBase := m.M
	_ = mBase
	var v6 int32
	_ = v6
	var v8 int32
	_ = v8
	var v10 int32
	_ = v10
	var v11 int32
	_ = v11
	var v12 int32
	_ = v12
	var v14 int32
	_ = v14
	var v17 int64
	_ = v17
	var v21 int32
	_ = v21
	var v22 int64
	_ = v22
	var v23 int32
	_ = v23
	var v30 int32
	_ = v30
	v6 = m.G0
	v8 = v6 - int32(32)
	m.G0 = v8
	v10 = *(*int32)(unsafe.Add(mBase, uint32(l1)+96))
	v11 = *(*int32)(unsafe.Add(mBase, uint32(v10)+64))
	v12 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v10)+48)))
	v14 = v12 & int32(240)
	if v14 != 0 {
		if v14 == int32(16) {
			v22 = *(*int64)(unsafe.Add(mBase, uint32(v11)))
			v23 = *(*int32)(unsafe.Add(mBase, uint32(v11)+8))
			*(*int32)(unsafe.Add(mBase, uint32(v8)+24)) = v23
			*(*int64)(unsafe.Add(mBase, uint32(v8)+16)) = v22
			F_appendStringInfo(m, l0, int32(_a_F_clog_desc_0), v8+int32(16))
			mBase = m.M
			v30 = m.ExcPending
			if v30 != 0 {
				return
			} else {
				m.G0 = v8 + int32(32)
				return
			}
		} else {
			m.G0 = v8 + int32(32)
			return
		}
	} else {
		v17 = *(*int64)(unsafe.Add(mBase, uint32(v11)))
		*(*int64)(unsafe.Add(mBase, uint32(v8))) = v17
		F_appendStringInfo(m, l0, int32(_a_F_clog_desc_1), v8)
		mBase = m.M
		v21 = m.ExcPending
		if v21 != 0 {
			return
		} else {
			m.G0 = v8 + int32(32)
			return
		}
	}
}
func F_closedir(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v3 int32
	_ = v3
	var v4 int32
	_ = v4
	v3 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	v4 = F_close(m, v3)
	mBase = m.M
	F_emscripten_builtin_free(m, l0)
	mBase = m.M
	return v4
}
func F_closerel(m *base.Module, l0 int32) {
	mBase := m.M
	_ = mBase
	var v4 int32
	_ = v4
	var v6 int32
	_ = v6
	var v9 int32
	_ = v9
	var v10 int32
	_ = v10
	var v12 int32
	_ = v12
	var v15 int32
	_ = v15
	var v18 int32
	_ = v18
	var v21 int32
	_ = v21
	var v22 int32
	_ = v22
	var v25 int32
	_ = v25
	var v26 int32
	_ = v26
	var v29 int32
	_ = v29
	var v36 int32
	_ = v36
	var v37 int32
	_ = v37
	var v44 int32
	_ = v44
	var v46 int32
	_ = v46
	var v47 int32
	_ = v47
	var v56 int32
	_ = v56
	var v61 int32
	_ = v61
	var v65 int32
	_ = v65
	var v71 int32
	_ = v71
	var v76 int32
	_ = v76
	var v81 int32
	_ = v81
	var v82 int32
	_ = v82
	var v84 int32
	_ = v84
	var v85 int32
	_ = v85
	var v91 int32
	_ = v91
	var v96 int32
	_ = v96
	var v98 int32
	_ = v98
	var v101 int32
	_ = v101
	var v111 int32
	_ = v111
	var v115 int32
	_ = v115
	var v120 int32
	_ = v120
	v4 = m.G0
	v6 = v4 - int32(48)
	m.G0 = v6
	v9 = *(*int32)(unsafe.Add(mBase, _c_F_closerel[0]))
	if l0 != 0 {
		goto L3
	} else {
		goto L4
	}
L1:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v111 = m.ExcPending
	if v111 != 0 {
		goto L17
	} else {
		goto L32
	}
L2:
	;
	v81 = F_errstart(m, int32(11), int32(0))
	mBase = m.M
	v82 = m.ExcPending
	if v82 != 0 {
		goto L17
	} else {
		goto L25
	}
L3:
	;
	if v9 != 0 {
		goto L6
	} else {
		goto L7
	}
L4:
	;
	goto L5
L5:
	;
	if v9 == int32(0) {
		goto L1
	} else {
		goto L24
	}
L6:
	;
	v10 = *(*int32)(unsafe.Add(mBase, uint32(v9)+48))
	v12 = v10 + int32(4)
	v15 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v12))))
	v18 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0))))
	if base.B2i32(v15 == int32(0))|base.B2i32(v15 != v18) != 0 {
		v36 = v15
		v37 = v18
		goto L10
	} else {
		goto L11
	}
L7:
	;
	goto L8
L8:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v65 = m.ExcPending
	if v65 != 0 {
		goto L17
	} else {
		goto L21
	}
L9:
	;
	if v36-v37 == int32(0) {
		goto L2
	} else {
		goto L16
	}
L10:
	;
	goto L9
L11:
	;
	v21 = v12
	v22 = l0
	goto L12
L12:
	;
	v25 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v22)+1)))
	v26 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v21)+1)))
	if v26 == int32(0) {
		v36 = v26
		v37 = v25
		goto L10
	} else {
		goto L14
	}
L13:
	;
	v36 = v26
	v37 = v25
	goto L10
L14:
	;
	v29 = int32(1)
	if v26 == v25 {
		v21 = v21 + v29
		v22 = v22 + v29
		goto L12
	} else {
		goto L15
	}
L15:
	;
	goto L13
L16:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v44 = m.ExcPending
	if v44 != 0 {
		goto L17
	} else {
		goto L18
	}
L17:
	;
	return
L18:
	;
	v46 = *(*int32)(unsafe.Add(mBase, _c_F_closerel[0]))
	v47 = *(*int32)(unsafe.Add(mBase, uint32(v46)+48))
	*(*int32)(unsafe.Add(mBase, uint32(v6)+32)) = l0
	*(*int32)(unsafe.Add(mBase, uint32(v6)+36)) = v47 + int32(4)
	F_errmsg_internal(m, int32(_a_F_closerel_0), v6+int32(32))
	mBase = m.M
	v56 = m.ExcPending
	if v56 != 0 {
		goto L17
	} else {
		goto L19
	}
L19:
	;
	F_errfinish(m, int32(_a_F_closerel_1), int32(537), int32(_a_F_closerel_2))
	mBase = m.M
	v61 = m.ExcPending
	if v61 != 0 {
		goto L17
	} else {
		goto L20
	}
L20:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L21:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v6)+16)) = l0
	F_errmsg_internal(m, int32(_a_F_closerel_3), v6+int32(16))
	mBase = m.M
	v71 = m.ExcPending
	if v71 != 0 {
		goto L17
	} else {
		goto L22
	}
L22:
	;
	F_errfinish(m, int32(_a_F_closerel_1), int32(541), int32(_a_F_closerel_2))
	mBase = m.M
	v76 = m.ExcPending
	if v76 != 0 {
		goto L17
	} else {
		goto L23
	}
L23:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L24:
	;
	goto L2
L25:
	;
	if v81 != 0 {
		goto L26
	} else {
		goto L27
	}
L26:
	;
	v84 = *(*int32)(unsafe.Add(mBase, _c_F_closerel[0]))
	v85 = *(*int32)(unsafe.Add(mBase, uint32(v84)+48))
	*(*int32)(unsafe.Add(mBase, uint32(v6))) = v85 + int32(4)
	F_errmsg_internal(m, int32(_a_F_closerel_4), v6)
	mBase = m.M
	v91 = m.ExcPending
	if v91 != 0 {
		goto L17
	} else {
		goto L29
	}
L27:
	;
	goto L28
L28:
	;
	v98 = *(*int32)(unsafe.Add(mBase, _c_F_closerel[0]))
	F_relation_close(m, v98, int32(0))
	mBase = m.M
	v101 = m.ExcPending
	if v101 != 0 {
		goto L17
	} else {
		goto L31
	}
L29:
	;
	F_errfinish(m, int32(_a_F_closerel_1), int32(549), int32(_a_F_closerel_2))
	mBase = m.M
	v96 = m.ExcPending
	if v96 != 0 {
		goto L17
	} else {
		goto L30
	}
L30:
	;
	goto L28
L31:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_closerel[0])) = int32(0)
	m.G0 = v6 + int32(48)
	return
L32:
	;
	F_errmsg_internal(m, int32(_a_F_closerel_5), int32(0))
	mBase = m.M
	v115 = m.ExcPending
	if v115 != 0 {
		goto L17
	} else {
		goto L33
	}
L33:
	;
	F_errfinish(m, int32(_a_F_closerel_1), int32(545), int32(_a_F_closerel_2))
	mBase = m.M
	v120 = m.ExcPending
	if v120 != 0 {
		goto L17
	} else {
		goto L34
	}
L34:
	;
	base.Wasm_trap_unreachable()
	for {
	}
}
