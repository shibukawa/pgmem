package p3

import (
	base "github.com/shibukawa/pgmem/internal/aot/pgaot/base"
	"unsafe"
)

func F_CLOGPagePrecedes(m *base.Module, l0 int64, l1 int64) int32 {
	var v3 int32
	_ = v3
	var v7 int32
	_ = v7
	var v9 int32
	_ = v9
	var v10 int32
	_ = v10
	var v13 int32
	_ = v13
	var v15 int32
	_ = v15
	var v27 int32
	_ = v27
	var v29 int32
	_ = v29
	var v41 int32
	_ = v41
	var v42 int32
	_ = v42
	v3 = int32(0)
	v7 = int32(15)
	v9 = int32(4)
	v10 = base.I32_wrap_i64(l0)<<(uint(v7)%32) | v9
	v13 = base.I32_wrap_i64(l1) << (uint(v7) % 32)
	v15 = v13 | v9
	if base.B2i32(base.Ui32(int32(2)) < base.Ui32(v15))&base.B2i32(base.Ui32(int32(3)) <= base.Ui32(v10)) == v3 {
		v27 = base.B2i32(base.Ui32(v10) < base.Ui32(v15))
	} else {
		v27 = int32(base.Ui32(v10-v15) >> (uint(int32(31)) % 32))
	}
	if v27 != 0 {
		v29 = v13 + int32(_a_F_CLOGPagePrecedes_0)
		if base.B2i32(base.Ui32(int32(2)) < base.Ui32(v29))&base.B2i32(base.Ui32(int32(3)) <= base.Ui32(v10)) == int32(0) {
			v41 = base.B2i32(base.Ui32(v10) < base.Ui32(v29))
		} else {
			v41 = int32(base.Ui32(v10-v29) >> (uint(int32(31)) % 32))
		}
		v42 = v41
	} else {
		v42 = v3
	}
	return v42
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
							F_errfinish(m, int32(_a_F_ClosePipeToProgram_1), int32(572), int32(_a_F_ClosePipeToProgram_2))
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
									F_errfinish(m, int32(_a_F_ClosePipeToProgram_1), int32(579), int32(_a_F_ClosePipeToProgram_2))
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
	var v7 int32
	_ = v7
	var v21 int32
	_ = v21
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
	var v56 int32
	_ = v56
	var v58 int32
	_ = v58
	var v61 int32
	_ = v61
	var v62 int32
	_ = v62
	var v63 int32
	_ = v63
	var v64 int32
	_ = v64
	var v65 int32
	_ = v65
	var v66 int32
	_ = v66
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
	var v82 int32
	_ = v82
	var v83 float64
	_ = v83
	var v84 int32
	_ = v84
	var v85 int32
	_ = v85
	var v86 int32
	_ = v86
	var v90 int32
	_ = v90
	var v94 int32
	_ = v94
	var v95 float64
	_ = v95
	var v96 int32
	_ = v96
	var v98 int32
	_ = v98
	var v99 int32
	_ = v99
	var v100 int32
	_ = v100
	var v101 float64
	_ = v101
	var v102 int32
	_ = v102
	var v104 int32
	_ = v104
	var v105 float64
	_ = v105
	var v106 int32
	_ = v106
	var v107 int32
	_ = v107
	var v108 float64
	_ = v108
	var v109 int32
	_ = v109
	var v111 int32
	_ = v111
	var v115 int32
	_ = v115
	var v116 int32
	_ = v116
	var v121 int32
	_ = v121
	var v122 int32
	_ = v122
	var v128 float64
	_ = v128
	var v129 int32
	_ = v129
	var v130 float64
	_ = v130
	var v133 int32
	_ = v133
	var v134 int32
	_ = v134
	var v139 int32
	_ = v139
	var v147 int32
	_ = v147
	var v150 float64
	_ = v150
	var v154 int32
	_ = v154
	var v156 int32
	_ = v156
	var v157 int32
	_ = v157
	var v158 int32
	_ = v158
	var v159 int32
	_ = v159
	var v165 int32
	_ = v165
	var v166 float64
	_ = v166
	var v167 int32
	_ = v167
	var v171 float64
	_ = v171
	var v174 int32
	_ = v174
	var v175 int32
	_ = v175
	var v189 float64
	_ = v189
	var v196 int32
	_ = v196
	var v199 int32
	_ = v199
	var v202 int32
	_ = v202
	var v203 int32
	_ = v203
	var v206 int32
	_ = v206
	var v207 int32
	_ = v207
	var v208 float64
	_ = v208
	var v209 int32
	_ = v209
	var v210 int32
	_ = v210
	var v211 int32
	_ = v211
	var v212 int32
	_ = v212
	var v215 int32
	_ = v215
	var v218 float64
	_ = v218
	var v219 int32
	_ = v219
	var v220 int32
	_ = v220
	var v221 int32
	_ = v221
	var v224 int32
	_ = v224
	var v225 float64
	_ = v225
	var v226 int32
	_ = v226
	var v227 int32
	_ = v227
	var v230 int32
	_ = v230
	var v233 float64
	_ = v233
	var v234 int32
	_ = v234
	var v235 int32
	_ = v235
	var v236 int32
	_ = v236
	var v239 int32
	_ = v239
	var v240 float64
	_ = v240
	var v241 int32
	_ = v241
	var v242 int32
	_ = v242
	var v244 int32
	_ = v244
	var v246 int32
	_ = v246
	var v247 int32
	_ = v247
	var v248 int32
	_ = v248
	var v249 int32
	_ = v249
	var v250 int32
	_ = v250
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
	var v264 int32
	_ = v264
	var v265 int32
	_ = v265
	var v268 int32
	_ = v268
	var v269 int32
	_ = v269
	var v272 float64
	_ = v272
	var v273 int32
	_ = v273
	var v274 float64
	_ = v274
	var v275 int32
	_ = v275
	var v276 float64
	_ = v276
	var v280 int32
	_ = v280
	var v281 int32
	_ = v281
	var v282 float64
	_ = v282
	var v283 int32
	_ = v283
	var v284 int32
	_ = v284
	var v285 int32
	_ = v285
	var v286 int32
	_ = v286
	var v288 int32
	_ = v288
	var v293 int32
	_ = v293
	var v294 int32
	_ = v294
	var v295 int32
	_ = v295
	var v296 int32
	_ = v296
	var v298 float32
	_ = v298
	var v299 float64
	_ = v299
	var v305 int32
	_ = v305
	var v306 int32
	_ = v306
	var v309 int32
	_ = v309
	var v312 float64
	_ = v312
	var v313 int32
	_ = v313
	var v314 float32
	_ = v314
	var v315 float64
	_ = v315
	var v319 int32
	_ = v319
	var v320 int32
	_ = v320
	var v321 float64
	_ = v321
	var v322 float64
	_ = v322
	var v323 float64
	_ = v323
	var v329 int32
	_ = v329
	var v335 int32
	_ = v335
	var v340 int32
	_ = v340
	var v343 float64
	_ = v343
	var v347 int32
	_ = v347
	var v359 int32
	_ = v359
	var v365 int32
	_ = v365
	var v370 int32
	_ = v370
	var v376 float64
	_ = v376
	var v377 int32
	_ = v377
	var v382 int32
	_ = v382
	var v386 int32
	_ = v386
	var v391 int32
	_ = v391
	var v392 float64
	_ = v392
	var v393 int32
	_ = v393
	var v395 float64
	_ = v395
	var v399 int32
	_ = v399
	var v400 int32
	_ = v400
	var v402 int32
	_ = v402
	var v403 float64
	_ = v403
	var v413 float64
	_ = v413
	var v420 int32
	_ = v420
	var v421 int32
	_ = v421
	var v422 int32
	_ = v422
	var v423 float64
	_ = v423
	var v428 float64
	_ = v428
	var v429 int32
	_ = v429
	var v430 float64
	_ = v430
	var v431 int32
	_ = v431
	var v432 int32
	_ = v432
	var v433 float64
	_ = v433
	var v434 int32
	_ = v434
	var v435 int32
	_ = v435
	var v436 int32
	_ = v436
	var v437 float64
	_ = v437
	var v438 int32
	_ = v438
	var v439 float64
	_ = v439
	var v441 int32
	_ = v441
	var v444 float64
	_ = v444
	var v446 int32
	_ = v446
	var v448 int32
	_ = v448
	var v451 int32
	_ = v451
	var v452 int32
	_ = v452
	var v457 int32
	_ = v457
	var v458 int32
	_ = v458
	var v462 float64
	_ = v462
	var v463 int32
	_ = v463
	var v464 int32
	_ = v464
	var v467 int32
	_ = v467
	var v469 int32
	_ = v469
	var v471 float64
	_ = v471
	var v487 float64
	_ = v487
	var v495 int32
	_ = v495
	var v510 float64
	_ = v510
	v7 = int32(0)
	if l1 == v7 {
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
	v21 = *(*int32)(unsafe.Add(mBase, uint32(l1)))
	if v21 != int32(318) {
		goto L6
	} else {
		goto L7
	}
L4:
	;
	return v510
L5:
	;
	switch v65 - int32(6) {
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
	v64 = l1
	v65 = v21
	v66 = v7
	v68 = int32(0)
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
		v56 = v7
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
		v56 = v7
		goto L12
	} else {
		goto L18
	}
L18:
	;
	goto L13
L19:
	;
	v56 = int32(1)
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
		v510 = v51
		goto L4
	} else {
		goto L24
	}
L23:
	;
	v510 = v46
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
	v64 = v62
	v65 = v63
	v66 = v56
	v68 = l1
	goto L5
L28:
	;
	if v66 == int32(0) {
		v510 = v487
		goto L4
	} else {
		goto L189
	}
L29:
	;
	v446 = m.G0
	v448 = v446 - int32(32)
	m.G0 = v448
	F_examine_variable(m, l0, v64, l2, v448)
	mBase = m.M
	v451 = m.ExcPending
	if v451 != 0 {
		goto L16
	} else {
		goto L181
	}
L30:
	;
	v441 = *(*int32)(unsafe.Add(mBase, uint32(v64)))
	if v441 == int32(18) {
		goto L178
	} else {
		goto L179
	}
L31:
	;
	v435 = *(*int32)(unsafe.Add(mBase, uint32(v64)+28))
	v436 = *(*int32)(unsafe.Add(mBase, uint32(v64)+24))
	v437 = F_restriction_selectivity(m, l0, v196, v435, v436, l2)
	mBase = m.M
	v438 = m.ExcPending
	if v438 != 0 {
		goto L16
	} else {
		goto L177
	}
L32:
	;
	v432 = *(*int32)(unsafe.Add(mBase, uint32(v64)+4))
	v433 = F_clause_selectivity_ext(m, l0, v432, l2, l3, l4, l5)
	mBase = m.M
	v434 = m.ExcPending
	if v434 != 0 {
		goto L16
	} else {
		goto L176
	}
L33:
	;
	v429 = *(*int32)(unsafe.Add(mBase, uint32(v64)+4))
	v430 = F_clause_selectivity_ext(m, l0, v429, l2, l3, l4, l5)
	mBase = m.M
	v431 = m.ExcPending
	if v431 != 0 {
		goto L16
	} else {
		goto L175
	}
L34:
	;
	v420 = *(*int32)(unsafe.Add(mBase, uint32(v64)+4))
	v421 = F_find_base_rel(m, l0, v420)
	mBase = m.M
	v422 = m.ExcPending
	if v422 != 0 {
		goto L16
	} else {
		goto L171
	}
L35:
	;
	v284 = *(*int32)(unsafe.Add(mBase, uint32(v64)+8))
	v285 = *(*int32)(unsafe.Add(mBase, uint32(v64)+4))
	v286 = m.G0
	v288 = v286 - int32(112)
	m.G0 = v288
	F_examine_variable(m, l0, v285, l2, v288+int32(80))
	mBase = m.M
	v293 = m.ExcPending
	if v293 != 0 {
		goto L16
	} else {
		goto L125
	}
L36:
	;
	v280 = *(*int32)(unsafe.Add(mBase, uint32(v64)+8))
	v281 = *(*int32)(unsafe.Add(mBase, uint32(v64)+4))
	v282 = F_nulltestsel(m, l0, v280, v281, l2)
	mBase = m.M
	v283 = m.ExcPending
	if v283 != 0 {
		goto L16
	} else {
		goto L124
	}
L37:
	;
	v242 = m.G0
	v244 = v242 - int32(16)
	m.G0 = v244
	v246 = *(*int32)(unsafe.Add(mBase, uint32(v64)+16))
	v247 = *(*int32)(unsafe.Add(mBase, uint32(v246)+12))
	v248 = *(*int32)(unsafe.Add(mBase, uint32(v247)))
	v249 = *(*int32)(unsafe.Add(mBase, uint32(v64)+8))
	v250 = *(*int32)(unsafe.Add(mBase, uint32(v249)+12))
	v251 = *(*int32)(unsafe.Add(mBase, uint32(v250)))
	v252 = *(*int32)(unsafe.Add(mBase, uint32(v64)+20))
	v253 = *(*int32)(unsafe.Add(mBase, uint32(v252)+12))
	v254 = *(*int32)(unsafe.Add(mBase, uint32(v253)))
	*(*int32)(unsafe.Add(mBase, uint32(v244)+12)) = v254
	v256 = *(*int32)(unsafe.Add(mBase, uint32(v64)+24))
	v257 = *(*int32)(unsafe.Add(mBase, uint32(v256)+12))
	v258 = *(*int32)(unsafe.Add(mBase, uint32(v257)))
	*(*int32)(unsafe.Add(mBase, uint32(v244)+8)) = v258
	*(*int32)(unsafe.Add(mBase, uint32(v244)+4)) = v254
	*(*int32)(unsafe.Add(mBase, uint32(v244))) = v258
	v264 = F_list_make2_impl(m, v244+int32(4), v244)
	mBase = m.M
	v265 = m.ExcPending
	if v265 != 0 {
		goto L16
	} else {
		goto L115
	}
L38:
	;
	v227 = int32(0)
	if l2 != 0 {
		v239 = v227
		goto L106
	} else {
		goto L107
	}
L39:
	;
	v210 = *(*int32)(unsafe.Add(mBase, uint32(v64)+24))
	v211 = *(*int32)(unsafe.Add(mBase, uint32(v64)+28))
	v212 = *(*int32)(unsafe.Add(mBase, uint32(v64)+4))
	if l2 != 0 {
		v224 = v7
		goto L97
	} else {
		goto L98
	}
L40:
	;
	v196 = *(*int32)(unsafe.Add(mBase, uint32(v64)+4))
	if l2 != 0 {
		goto L31
	} else {
		goto L87
	}
L41:
	;
	v96 = *(*int32)(unsafe.Add(mBase, uint32(v64)+4))
	switch v96 {
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
	v84 = F_estimate_expression_value(m, l0, v64)
	mBase = m.M
	v85 = m.ExcPending
	if v85 != 0 {
		goto L16
	} else {
		goto L54
	}
L43:
	;
	v78 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v64)+24)))
	if v78 != 0 {
		goto L48
	} else {
		goto L49
	}
L44:
	;
	v72 = float64(0.5)
	v73 = *(*int32)(unsafe.Add(mBase, uint32(v64)+28))
	if v73 != 0 {
		v487 = v72
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
	v76 = *(*int32)(unsafe.Add(mBase, uint32(v64)+4))
	if l2 != v76 {
		v487 = v72
		goto L28
	} else {
		goto L47
	}
L47:
	;
	goto L29
L48:
	;
	v487 = float64(0)
	goto L28
L49:
	;
	goto L50
L50:
	;
	v82 = *(*int32)(unsafe.Add(mBase, uint32(v64)+20))
	if v82 != 0 {
		goto L51
	} else {
		goto L52
	}
L51:
	;
	v83 = float64(1)
	goto L53
L52:
	;
	v83 = float64(0)
	goto L53
L53:
	;
	v487 = v83
	goto L28
L54:
	;
	v86 = *(*int32)(unsafe.Add(mBase, uint32(v84)))
	if v86 != int32(7) {
		goto L55
	} else {
		goto L56
	}
L55:
	;
	v487 = float64(0.5)
	goto L28
L56:
	;
	goto L57
L57:
	;
	v90 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v84)+24)))
	if v90 != 0 {
		goto L58
	} else {
		goto L59
	}
L58:
	;
	v487 = float64(0)
	goto L28
L59:
	;
	goto L60
L60:
	;
	v94 = *(*int32)(unsafe.Add(mBase, uint32(v84)+20))
	if v94 != 0 {
		goto L61
	} else {
		goto L62
	}
L61:
	;
	v95 = float64(1)
	goto L63
L62:
	;
	v95 = float64(0)
	goto L63
L63:
	;
	v487 = v95
	goto L28
L64:
	;
	v107 = *(*int32)(unsafe.Add(mBase, uint32(v64)+8))
	v108 = float64(0)
	v109 = m.G0
	v111 = v109 - int32(16)
	m.G0 = v111
	*(*int32)(unsafe.Add(mBase, uint32(v111)+12)) = int32(0)
	v115 = F_find_single_rel_for_clauses(m, l0, v107)
	mBase = m.M
	v116 = m.ExcPending
	if v116 != 0 {
		goto L16
	} else {
		goto L69
	}
L65:
	;
	v104 = *(*int32)(unsafe.Add(mBase, uint32(v64)+8))
	v105 = F_clauselist_selectivity_ext(m, l0, v104, l2, l3, l4, l5)
	mBase = m.M
	v106 = m.ExcPending
	if v106 != 0 {
		goto L16
	} else {
		goto L68
	}
L66:
	;
	v98 = *(*int32)(unsafe.Add(mBase, uint32(v64)+8))
	v99 = *(*int32)(unsafe.Add(mBase, uint32(v98)+12))
	v100 = *(*int32)(unsafe.Add(mBase, uint32(v99)))
	v101 = F_clause_selectivity_ext(m, l0, v100, l2, l3, l4, l5)
	mBase = m.M
	v102 = m.ExcPending
	if v102 != 0 {
		goto L16
	} else {
		goto L67
	}
L67:
	;
	v487 = base.F64_sub(float64(1), v101)
	goto L28
L68:
	;
	v487 = v105
	goto L28
L69:
	;
	if l5 == int32(0) {
		v130 = v108
		goto L70
	} else {
		goto L71
	}
L70:
	;
	if v107 == int32(0) {
		v189 = v130
		goto L76
	} else {
		goto L77
	}
L71:
	;
	if v115 == int32(0) {
		v130 = v108
		goto L70
	} else {
		goto L72
	}
L72:
	;
	v121 = *(*int32)(unsafe.Add(mBase, uint32(v115)+76))
	if v121 != 0 {
		v130 = v108
		goto L70
	} else {
		goto L73
	}
L73:
	;
	v122 = *(*int32)(unsafe.Add(mBase, uint32(v115)+112))
	if v122 == int32(0) {
		v130 = v108
		goto L70
	} else {
		goto L74
	}
L74:
	;
	v128 = F_statext_clauselist_selectivity(m, l0, v107, l2, l3, l4, v115, v111+int32(12), int32(1))
	mBase = m.M
	v129 = m.ExcPending
	if v129 != 0 {
		goto L16
	} else {
		goto L75
	}
L75:
	;
	v130 = v128
	goto L70
L76:
	;
	m.G0 = v111 + int32(16)
	v487 = v189
	goto L28
L77:
	;
	v133 = int32(0)
	v134 = *(*int32)(unsafe.Add(mBase, uint32(v107)+4))
	if v134 <= v133 {
		v189 = v130
		goto L76
	} else {
		goto L78
	}
L78:
	;
	v139 = v133
	v147 = int32(-1)
	v150 = v130
	goto L79
L79:
	;
	v154 = *(*int32)(unsafe.Add(mBase, uint32(v107)+12))
	v156 = v147 + int32(1)
	v157 = *(*int32)(unsafe.Add(mBase, uint32(v111)+12))
	v158 = F_bms_is_member(m, v156, v157)
	mBase = m.M
	v159 = m.ExcPending
	if v159 != 0 {
		goto L16
	} else {
		goto L81
	}
L80:
	;
	v189 = v171
	goto L76
L81:
	;
	if v158 == int32(0) {
		goto L82
	} else {
		goto L83
	}
L82:
	;
	v165 = *(*int32)(unsafe.Add(mBase, uint32(v154+v139<<(uint(int32(2))%32))))
	v166 = F_clause_selectivity_ext(m, l0, v165, l2, l3, l4, l5)
	mBase = m.M
	v167 = m.ExcPending
	if v167 != 0 {
		goto L16
	} else {
		goto L85
	}
L83:
	;
	v171 = v150
	goto L84
L84:
	;
	v174 = v139 + int32(1)
	v175 = *(*int32)(unsafe.Add(mBase, uint32(v107)+4))
	if v174 < v175 {
		v139 = v174
		v147 = v156
		v150 = v171
		goto L79
	} else {
		goto L86
	}
L85:
	;
	v171 = base.F64_sub(base.F64_add(v150, v166), base.F64_mul(v150, v166))
	goto L84
L86:
	;
	goto L80
L87:
	;
	if l4 == int32(0) {
		goto L31
	} else {
		goto L88
	}
L88:
	;
	if v68 != 0 {
		goto L90
	} else {
		goto L91
	}
L89:
	;
	v206 = *(*int32)(unsafe.Add(mBase, uint32(v64)+28))
	v207 = *(*int32)(unsafe.Add(mBase, uint32(v64)+24))
	v208 = F_join_selectivity(m, l0, v196, v206, v207, l3, l4)
	mBase = m.M
	v209 = m.ExcPending
	if v209 != 0 {
		goto L16
	} else {
		goto L96
	}
L90:
	;
	v199 = *(*int32)(unsafe.Add(mBase, uint32(v68)+24))
	if int32(1) < v199 {
		goto L89
	} else {
		goto L93
	}
L91:
	;
	goto L92
L92:
	;
	v202 = F_NumRelids(m, l0, v64)
	mBase = m.M
	v203 = m.ExcPending
	if v203 != 0 {
		goto L16
	} else {
		goto L94
	}
L93:
	;
	goto L31
L94:
	;
	if v202 < int32(2) {
		goto L31
	} else {
		goto L95
	}
L95:
	;
	goto L89
L96:
	;
	v439 = v208
	goto L30
L97:
	;
	v225 = F_function_selectivity(m, l0, v212, v211, v210, v224, l2, l3, l4)
	mBase = m.M
	v226 = m.ExcPending
	if v226 != 0 {
		goto L16
	} else {
		goto L105
	}
L98:
	;
	if l4 == int32(0) {
		v224 = v7
		goto L97
	} else {
		goto L99
	}
L99:
	;
	if v68 != 0 {
		goto L100
	} else {
		goto L101
	}
L100:
	;
	v215 = *(*int32)(unsafe.Add(mBase, uint32(v68)+24))
	v218 = F_function_selectivity(m, l0, v212, v211, v210, base.B2i32(int32(1) < v215), l2, l3, l4)
	mBase = m.M
	v219 = m.ExcPending
	if v219 != 0 {
		goto L16
	} else {
		goto L103
	}
L101:
	;
	goto L102
L102:
	;
	v220 = F_NumRelids(m, l0, v64)
	mBase = m.M
	v221 = m.ExcPending
	if v221 != 0 {
		goto L16
	} else {
		goto L104
	}
L103:
	;
	v487 = v218
	goto L28
L104:
	;
	v224 = base.B2i32(int32(1) < v220)
	goto L97
L105:
	;
	v487 = v225
	goto L28
L106:
	;
	v240 = F_scalararraysel(m, l0, v64, v239, l2, l3, l4)
	mBase = m.M
	v241 = m.ExcPending
	if v241 != 0 {
		goto L16
	} else {
		goto L114
	}
L107:
	;
	if l4 == int32(0) {
		v239 = v227
		goto L106
	} else {
		goto L108
	}
L108:
	;
	if v68 != 0 {
		goto L109
	} else {
		goto L110
	}
L109:
	;
	v230 = *(*int32)(unsafe.Add(mBase, uint32(v68)+24))
	v233 = F_scalararraysel(m, l0, v64, base.B2i32(int32(1) < v230), l2, l3, l4)
	mBase = m.M
	v234 = m.ExcPending
	if v234 != 0 {
		goto L16
	} else {
		goto L112
	}
L110:
	;
	goto L111
L111:
	;
	v235 = F_NumRelids(m, l0, v64)
	mBase = m.M
	v236 = m.ExcPending
	if v236 != 0 {
		goto L16
	} else {
		goto L113
	}
L112:
	;
	v487 = v233
	goto L28
L113:
	;
	v239 = base.B2i32(int32(1) < v235)
	goto L106
L114:
	;
	v487 = v240
	goto L28
L115:
	;
	if l2 != 0 {
		goto L117
	} else {
		goto L118
	}
L116:
	;
	m.G0 = v244 + int32(16)
	v487 = v276
	goto L28
L117:
	;
	v274 = F_restriction_selectivity(m, l0, v251, v264, v248, l2)
	mBase = m.M
	v275 = m.ExcPending
	if v275 != 0 {
		goto L16
	} else {
		goto L123
	}
L118:
	;
	if l4 == int32(0) {
		goto L117
	} else {
		goto L119
	}
L119:
	;
	v268 = F_NumRelids(m, l0, v264)
	mBase = m.M
	v269 = m.ExcPending
	if v269 != 0 {
		goto L16
	} else {
		goto L120
	}
L120:
	;
	if v268 < int32(2) {
		goto L117
	} else {
		goto L121
	}
L121:
	;
	v272 = F_join_selectivity(m, l0, v251, v264, v248, l3, l4)
	mBase = m.M
	v273 = m.ExcPending
	if v273 != 0 {
		goto L16
	} else {
		goto L122
	}
L122:
	;
	v276 = v272
	goto L116
L123:
	;
	v276 = v274
	goto L116
L124:
	;
	v487 = v282
	goto L28
L125:
	;
	v294 = *(*int32)(unsafe.Add(mBase, uint32(v288)+88))
	if v294 != 0 {
		goto L128
	} else {
		goto L129
	}
L126:
	;
	m.G0 = v288 + int32(112)
	v487 = v413
	goto L28
L127:
	;
	v399 = *(*int32)(unsafe.Add(mBase, uint32(v288)+88))
	if v399 != 0 {
		goto L165
	} else {
		goto L166
	}
L128:
	;
	v295 = *(*int32)(unsafe.Add(mBase, uint32(v294)+16))
	v296 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v295)+22)))
	v298 = *(*float32)(unsafe.Add(mBase, uint32(v295+v296)+8))
	v299 = base.F64_promote_f32(v298)
	v305 = F_get_attstatsslot(m, v288+int32(44), v294, int32(1), int32(0), int32(3))
	mBase = m.M
	v306 = m.ExcPending
	if v306 != 0 {
		goto L16
	} else {
		goto L132
	}
L129:
	;
	goto L130
L130:
	;
	switch v284 {
	case 0, 3:
		goto L156
	case 1, 2:
		goto L158
	case 4:
		v413 = float64(0.005)
		goto L126
	case 5:
		goto L159
	default:
		goto L157
	}
L131:
	;
	switch v284 {
	case 0, 2:
		goto L152
	case 1, 3:
		goto L151
	case 4:
		v395 = v299
		goto L127
	case 5:
		goto L149
	default:
		goto L150
	}
L132:
	;
	if v305 == int32(0) {
		goto L131
	} else {
		goto L133
	}
L133:
	;
	v309 = *(*int32)(unsafe.Add(mBase, uint32(v288)+68))
	if v309 <= int32(0) {
		goto L131
	} else {
		goto L134
	}
L134:
	;
	v312 = float64(1)
	v313 = *(*int32)(unsafe.Add(mBase, uint32(v288)+64))
	v314 = *(*float32)(unsafe.Add(mBase, uint32(v313)))
	v315 = base.F64_promote_f32(v314)
	v319 = *(*int32)(unsafe.Add(mBase, uint32(v288)+56))
	v320 = *(*int32)(unsafe.Add(mBase, uint32(v319)))
	if v320 != 0 {
		goto L135
	} else {
		goto L136
	}
L135:
	;
	v321 = v315
	goto L137
L136:
	;
	v321 = base.F64_sub(base.F64_sub(v312, v315), v299)
	goto L137
L137:
	;
	v322 = base.F64_sub(v312, v321)
	v323 = base.F64_sub(v322, v299)
	switch v284 {
	case 0:
		goto L144
	case 1:
		goto L143
	case 2:
		goto L142
	case 3:
		goto L141
	case 4:
		v343 = v299
		goto L138
	case 5:
		goto L139
	default:
		goto L140
	}
L138:
	;
	F_free_attstatsslot(m, v288+int32(44))
	mBase = m.M
	v347 = m.ExcPending
	if v347 != 0 {
		goto L16
	} else {
		goto L148
	}
L139:
	;
	v343 = base.F64_sub(float64(1), v299)
	goto L138
L140:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v329 = m.ExcPending
	if v329 != 0 {
		goto L16
	} else {
		goto L145
	}
L141:
	;
	v343 = base.F64_sub(float64(1), v323)
	goto L138
L142:
	;
	v343 = v323
	goto L138
L143:
	;
	v343 = v322
	goto L138
L144:
	;
	v343 = v321
	goto L138
L145:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v288)+16)) = v284
	F_errmsg_internal(m, int32(_a_F_clause_selectivity_ext_0), v288+int32(16))
	mBase = m.M
	v335 = m.ExcPending
	if v335 != 0 {
		goto L16
	} else {
		goto L146
	}
L146:
	;
	F_errfinish(m, int32(_a_F_clause_selectivity_ext_1), int32(1615), int32(_a_F_clause_selectivity_ext_2))
	mBase = m.M
	v340 = m.ExcPending
	if v340 != 0 {
		goto L16
	} else {
		goto L147
	}
L147:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L148:
	;
	v395 = v343
	goto L127
L149:
	;
	v395 = base.F64_sub(float64(1), v299)
	goto L127
L150:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v359 = m.ExcPending
	if v359 != 0 {
		goto L16
	} else {
		goto L153
	}
L151:
	;
	v395 = base.F64_mul(base.F64_add(v299, float64(1)), float64(0.5))
	goto L127
L152:
	;
	v395 = base.F64_mul(base.F64_sub(float64(1), v299), float64(0.5))
	goto L127
L153:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v288)+32)) = v284
	F_errmsg_internal(m, int32(_a_F_clause_selectivity_ext_0), v288+int32(32))
	mBase = m.M
	v365 = m.ExcPending
	if v365 != 0 {
		goto L16
	} else {
		goto L154
	}
L154:
	;
	F_errfinish(m, int32(_a_F_clause_selectivity_ext_1), int32(1652), int32(_a_F_clause_selectivity_ext_2))
	mBase = m.M
	v370 = m.ExcPending
	if v370 != 0 {
		goto L16
	} else {
		goto L155
	}
L155:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L156:
	;
	v392 = F_clause_selectivity(m, l0, v285, l2, l3, l4)
	mBase = m.M
	v393 = m.ExcPending
	if v393 != 0 {
		goto L16
	} else {
		goto L164
	}
L157:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v382 = m.ExcPending
	if v382 != 0 {
		goto L16
	} else {
		goto L161
	}
L158:
	;
	v376 = F_clause_selectivity(m, l0, v285, l2, l3, l4)
	mBase = m.M
	v377 = m.ExcPending
	if v377 != 0 {
		goto L16
	} else {
		goto L160
	}
L159:
	;
	v413 = float64(0.995)
	goto L126
L160:
	;
	v395 = base.F64_sub(float64(1), v376)
	goto L127
L161:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v288))) = v284
	F_errmsg_internal(m, int32(_a_F_clause_selectivity_ext_0), v288)
	mBase = m.M
	v386 = m.ExcPending
	if v386 != 0 {
		goto L16
	} else {
		goto L162
	}
L162:
	;
	F_errfinish(m, int32(_a_F_clause_selectivity_ext_1), int32(1688), int32(_a_F_clause_selectivity_ext_2))
	mBase = m.M
	v391 = m.ExcPending
	if v391 != 0 {
		goto L16
	} else {
		goto L163
	}
L163:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L164:
	;
	v395 = v392
	goto L127
L165:
	;
	v400 = *(*int32)(unsafe.Add(mBase, uint32(v288)+92))
	m.T0[v400].(func(*base.Module, int32))(m, v399)
	mBase = m.M
	v402 = m.ExcPending
	if v402 != 0 {
		goto L16
	} else {
		goto L168
	}
L166:
	;
	goto L167
L167:
	;
	v403 = float64(0)
	if base.F64_lt(v395, v403) != 0 {
		v413 = v403
		goto L126
	} else {
		goto L169
	}
L168:
	;
	goto L167
L169:
	;
	if base.F64_gt(v395, float64(1)) == int32(0) {
		v413 = v395
		goto L126
	} else {
		goto L170
	}
L170:
	;
	v413 = float64(1)
	goto L126
L171:
	;
	v423 = *(*float64)(unsafe.Add(mBase, uint32(v421)+120))
	if base.F64_gt(v423, float64(0)) != 0 {
		goto L172
	} else {
		goto L173
	}
L172:
	;
	v428 = base.F64_div(float64(1), v423)
	goto L174
L173:
	;
	v428 = float64(0.5)
	goto L174
L174:
	;
	v487 = v428
	goto L28
L175:
	;
	v487 = v430
	goto L28
L176:
	;
	v487 = v433
	goto L28
L177:
	;
	v439 = v437
	goto L30
L178:
	;
	v444 = base.F64_sub(float64(1), v439)
	goto L180
L179:
	;
	v444 = v439
	goto L180
L180:
	;
	v487 = v444
	goto L28
L181:
	;
	v452 = *(*int32)(unsafe.Add(mBase, uint32(v448)+8))
	if v452 == int32(0) {
		goto L183
	} else {
		goto L184
	}
L182:
	;
	m.G0 = v448 + int32(32)
	v487 = v471
	goto L28
L183:
	;
	v471 = float64(0.5)
	goto L182
L184:
	;
	goto L185
L185:
	;
	v457 = int32(0)
	v458 = int32(1)
	v462 = F_var_eq_const(m, v448, int32(91), v457, v458, v457, v458, v457)
	mBase = m.M
	v463 = m.ExcPending
	if v463 != 0 {
		goto L16
	} else {
		goto L186
	}
L186:
	;
	v464 = *(*int32)(unsafe.Add(mBase, uint32(v448)+8))
	if v464 == int32(0) {
		v471 = v462
		goto L182
	} else {
		goto L187
	}
L187:
	;
	v467 = *(*int32)(unsafe.Add(mBase, uint32(v448)+12))
	m.T0[v467].(func(*base.Module, int32))(m, v464)
	mBase = m.M
	v469 = m.ExcPending
	if v469 != 0 {
		goto L16
	} else {
		goto L188
	}
L188:
	;
	v471 = v462
	goto L182
L189:
	;
	if l3 != 0 {
		goto L190
	} else {
		goto L191
	}
L190:
	;
	v495 = int32(88)
	goto L192
L191:
	;
	v495 = int32(80)
	goto L192
L192:
	;
	*(*float64)(unsafe.Add(mBase, uint32(v68+v495))) = v487
	v510 = v487
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
func F_clock_timestamp(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v3 int32
	_ = v3
	var v5 int32
	_ = v5
	var v8 int64
	_ = v8
	var v9 int64
	_ = v9
	var v15 int32
	_ = v15
	var v18 int32
	_ = v18
	v3 = m.G0
	v5 = v3 - int32(16)
	m.G0 = v5
	F___gettimeofday(m, v5)
	mBase = m.M
	v8 = int64(*(*int32)(unsafe.Add(mBase, uint32(v5)+8)))
	v9 = *(*int64)(unsafe.Add(mBase, uint32(v5)))
	v15 = F_Int64GetDatum(m, v8+v9*int64(1000000)-int64(946684800000000))
	mBase = m.M
	v18 = m.ExcPending
	if v18 != 0 {
		return int32(0)
	} else {
		m.G0 = v5 + int32(16)
		return v15
	}
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
	var v15 int64
	_ = v15
	var v19 int32
	_ = v19
	var v20 int64
	_ = v20
	var v21 int32
	_ = v21
	var v28 int32
	_ = v28
	v6 = m.G0
	v8 = v6 - int32(32)
	m.G0 = v8
	v10 = *(*int32)(unsafe.Add(mBase, uint32(l1)+96))
	v11 = *(*int32)(unsafe.Add(mBase, uint32(v10)+64))
	v12 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v10)+48)))
	switch v12 & int32(240) {
	case 0:
		v15 = *(*int64)(unsafe.Add(mBase, uint32(v11)))
		*(*int64)(unsafe.Add(mBase, uint32(v8))) = v15
		F_appendStringInfo(m, l0, int32(_a_F_clog_desc_0), v8)
		mBase = m.M
		v19 = m.ExcPending
		if v19 != 0 {
			return
		} else {
			m.G0 = v8 + int32(32)
			return
		}
	default:
		m.G0 = v8 + int32(32)
		return
	case 16:
		v20 = *(*int64)(unsafe.Add(mBase, uint32(v11)))
		v21 = *(*int32)(unsafe.Add(mBase, uint32(v11)+8))
		*(*int32)(unsafe.Add(mBase, uint32(v8)+24)) = v21
		*(*int64)(unsafe.Add(mBase, uint32(v8)+16)) = v20
		F_appendStringInfo(m, l0, int32(_a_F_clog_desc_1), v8+int32(16))
		mBase = m.M
		v28 = m.ExcPending
		if v28 != 0 {
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
	var v16 int32
	_ = v16
	var v20 int32
	_ = v20
	var v21 int32
	_ = v21
	var v24 int32
	_ = v24
	var v25 int32
	_ = v25
	var v28 int32
	_ = v28
	var v35 int32
	_ = v35
	var v36 int32
	_ = v36
	var v43 int32
	_ = v43
	var v45 int32
	_ = v45
	var v46 int32
	_ = v46
	var v55 int32
	_ = v55
	var v60 int32
	_ = v60
	var v64 int32
	_ = v64
	var v70 int32
	_ = v70
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
	var v90 int32
	_ = v90
	var v95 int32
	_ = v95
	var v97 int32
	_ = v97
	var v100 int32
	_ = v100
	var v110 int32
	_ = v110
	var v114 int32
	_ = v114
	var v119 int32
	_ = v119
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
	v110 = m.ExcPending
	if v110 != 0 {
		goto L18
	} else {
		goto L33
	}
L2:
	;
	v80 = F_errstart(m, int32(11), int32(0))
	mBase = m.M
	v81 = m.ExcPending
	if v81 != 0 {
		goto L18
	} else {
		goto L26
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
		goto L25
	}
L6:
	;
	v10 = *(*int32)(unsafe.Add(mBase, uint32(v9)+48))
	v12 = v10 + int32(4)
	v15 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0))))
	v16 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v12))))
	if v16 == int32(0) {
		v35 = v15
		v36 = v16
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
	v64 = m.ExcPending
	if v64 != 0 {
		goto L18
	} else {
		goto L22
	}
L9:
	;
	if v36-v35 == int32(0) {
		goto L2
	} else {
		goto L17
	}
L10:
	;
	goto L9
L11:
	;
	if v15 != v16 {
		v35 = v15
		v36 = v16
		goto L10
	} else {
		goto L12
	}
L12:
	;
	v20 = v12
	v21 = l0
	goto L13
L13:
	;
	v24 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v21)+1)))
	v25 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v20)+1)))
	if v25 == int32(0) {
		v35 = v24
		v36 = v25
		goto L10
	} else {
		goto L15
	}
L14:
	;
	v35 = v24
	v36 = v25
	goto L10
L15:
	;
	v28 = int32(1)
	if v24 == v25 {
		v20 = v20 + v28
		v21 = v21 + v28
		goto L13
	} else {
		goto L16
	}
L16:
	;
	goto L14
L17:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v43 = m.ExcPending
	if v43 != 0 {
		goto L18
	} else {
		goto L19
	}
L18:
	;
	return
L19:
	;
	v45 = *(*int32)(unsafe.Add(mBase, _c_F_closerel[0]))
	v46 = *(*int32)(unsafe.Add(mBase, uint32(v45)+48))
	*(*int32)(unsafe.Add(mBase, uint32(v6)+32)) = l0
	*(*int32)(unsafe.Add(mBase, uint32(v6)+36)) = v46 + int32(4)
	F_errmsg_internal(m, int32(_a_F_closerel_0), v6+int32(32))
	mBase = m.M
	v55 = m.ExcPending
	if v55 != 0 {
		goto L18
	} else {
		goto L20
	}
L20:
	;
	F_errfinish(m, int32(_a_F_closerel_1), int32(493), int32(_a_F_closerel_2))
	mBase = m.M
	v60 = m.ExcPending
	if v60 != 0 {
		goto L18
	} else {
		goto L21
	}
L21:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L22:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v6)+16)) = l0
	F_errmsg_internal(m, int32(_a_F_closerel_3), v6+int32(16))
	mBase = m.M
	v70 = m.ExcPending
	if v70 != 0 {
		goto L18
	} else {
		goto L23
	}
L23:
	;
	F_errfinish(m, int32(_a_F_closerel_1), int32(497), int32(_a_F_closerel_2))
	mBase = m.M
	v75 = m.ExcPending
	if v75 != 0 {
		goto L18
	} else {
		goto L24
	}
L24:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L25:
	;
	goto L2
L26:
	;
	if v80 != 0 {
		goto L27
	} else {
		goto L28
	}
L27:
	;
	v83 = *(*int32)(unsafe.Add(mBase, _c_F_closerel[0]))
	v84 = *(*int32)(unsafe.Add(mBase, uint32(v83)+48))
	*(*int32)(unsafe.Add(mBase, uint32(v6))) = v84 + int32(4)
	F_errmsg_internal(m, int32(_a_F_closerel_4), v6)
	mBase = m.M
	v90 = m.ExcPending
	if v90 != 0 {
		goto L18
	} else {
		goto L30
	}
L28:
	;
	goto L29
L29:
	;
	v97 = *(*int32)(unsafe.Add(mBase, _c_F_closerel[0]))
	F_sequence_close(m, v97, int32(0))
	mBase = m.M
	v100 = m.ExcPending
	if v100 != 0 {
		goto L18
	} else {
		goto L32
	}
L30:
	;
	F_errfinish(m, int32(_a_F_closerel_1), int32(505), int32(_a_F_closerel_2))
	mBase = m.M
	v95 = m.ExcPending
	if v95 != 0 {
		goto L18
	} else {
		goto L31
	}
L31:
	;
	goto L29
L32:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_closerel[0])) = int32(0)
	m.G0 = v6 + int32(48)
	return
L33:
	;
	F_errmsg_internal(m, int32(_a_F_closerel_5), int32(0))
	mBase = m.M
	v114 = m.ExcPending
	if v114 != 0 {
		goto L18
	} else {
		goto L34
	}
L34:
	;
	F_errfinish(m, int32(_a_F_closerel_1), int32(501), int32(_a_F_closerel_2))
	mBase = m.M
	v119 = m.ExcPending
	if v119 != 0 {
		goto L18
	} else {
		goto L35
	}
L35:
	;
	base.Wasm_trap_unreachable()
	for {
	}
}
