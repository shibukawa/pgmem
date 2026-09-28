package p0

import (
	base "github.com/shibukawa/pgmem/internal/aot/pgaot/base"
	"unsafe"
)

func F_TimestampDifferenceExceeds(m *base.Module, l0 int64, l1 int64, l2 int32) int32 {
	return base.B2i32(base.I64_extend_i32_s(l2)*int64(1000) <= l1-l0)
}
func F_extract_timestamp(m *base.Module, l0 int32) int64 {
	var v3 int64
	_ = v3
	var v6 int32
	_ = v6
	v3 = F_timestamp_part_common(m, l0, int32(1))
	v6 = m.ExcPending
	if v6 != 0 {
		return int64(0)
	} else {
		return v3
	}
}
func F_timestamp_age(m *base.Module, l0 int32) int64 {
	mBase := m.M
	_ = mBase
	var v21 int32
	_ = v21
	var v23 int32
	_ = v23
	var v25 int64
	_ = v25
	var v27 int64
	_ = v27
	var v29 int32
	_ = v29
	var v32 int32
	_ = v32
	var v35 int64
	_ = v35
	var v44 int32
	_ = v44
	var v47 int32
	_ = v47
	var v51 int32
	_ = v51
	var v56 int32
	_ = v56
	var v63 int32
	_ = v63
	var v66 int32
	_ = v66
	var v70 int32
	_ = v70
	var v75 int32
	_ = v75
	var v83 int64
	_ = v83
	var v84 int32
	_ = v84
	var v89 int64
	_ = v89
	var v97 int64
	_ = v97
	var v98 int64
	_ = v98
	var v101 int64
	_ = v101
	var v104 int32
	_ = v104
	var v116 int32
	_ = v116
	var v117 int32
	_ = v117
	var v118 int32
	_ = v118
	var v119 int32
	_ = v119
	var v125 int32
	_ = v125
	var v130 int32
	_ = v130
	var v133 int32
	_ = v133
	var v134 int32
	_ = v134
	var v135 int32
	_ = v135
	var v138 int32
	_ = v138
	var v140 int32
	_ = v140
	var v146 int32
	_ = v146
	var v150 int32
	_ = v150
	var v151 int32
	_ = v151
	var v153 int32
	_ = v153
	var v161 int32
	_ = v161
	var v165 int32
	_ = v165
	var v175 int32
	_ = v175
	var v187 int64
	_ = v187
	var v189 int64
	_ = v189
	var v194 int64
	_ = v194
	var v196 int64
	_ = v196
	var v201 int64
	_ = v201
	var v203 int64
	_ = v203
	var v206 int64
	_ = v206
	var v214 int64
	_ = v214
	var v215 int64
	_ = v215
	var v218 int64
	_ = v218
	var v221 int32
	_ = v221
	var v233 int32
	_ = v233
	var v234 int32
	_ = v234
	var v235 int32
	_ = v235
	var v236 int32
	_ = v236
	var v242 int32
	_ = v242
	var v247 int32
	_ = v247
	var v250 int32
	_ = v250
	var v251 int32
	_ = v251
	var v252 int32
	_ = v252
	var v255 int32
	_ = v255
	var v257 int32
	_ = v257
	var v263 int32
	_ = v263
	var v267 int32
	_ = v267
	var v268 int32
	_ = v268
	var v270 int32
	_ = v270
	var v278 int32
	_ = v278
	var v282 int32
	_ = v282
	var v292 int32
	_ = v292
	var v308 int64
	_ = v308
	var v310 int64
	_ = v310
	var v314 int64
	_ = v314
	var v316 int64
	_ = v316
	var v320 int64
	_ = v320
	var v322 int64
	_ = v322
	var v327 int32
	_ = v327
	var v328 int32
	_ = v328
	var v330 int32
	_ = v330
	var v331 int32
	_ = v331
	var v332 int32
	_ = v332
	var v334 int32
	_ = v334
	var v335 int32
	_ = v335
	var v336 int32
	_ = v336
	var v338 int32
	_ = v338
	var v340 int64
	_ = v340
	var v341 int32
	_ = v341
	var v342 int32
	_ = v342
	var v343 int32
	_ = v343
	var v344 int32
	_ = v344
	var v345 int32
	_ = v345
	var v346 int32
	_ = v346
	var v347 int32
	_ = v347
	var v348 int32
	_ = v348
	var v349 int32
	_ = v349
	var v350 int32
	_ = v350
	var v353 int32
	_ = v353
	var v367 int32
	_ = v367
	var v368 int64
	_ = v368
	var v369 int32
	_ = v369
	var v370 int32
	_ = v370
	var v371 int32
	_ = v371
	var v372 int32
	_ = v372
	var v373 int32
	_ = v373
	var v376 int32
	_ = v376
	var v379 int32
	_ = v379
	var v381 int32
	_ = v381
	var v384 int32
	_ = v384
	var v385 int32
	_ = v385
	var v386 int32
	_ = v386
	var v395 int32
	_ = v395
	var v396 int32
	_ = v396
	var v400 int32
	_ = v400
	var v403 int32
	_ = v403
	var v405 int32
	_ = v405
	var v408 int32
	_ = v408
	var v409 int32
	_ = v409
	var v410 int32
	_ = v410
	var v419 int32
	_ = v419
	var v420 int32
	_ = v420
	var v424 int32
	_ = v424
	var v427 int32
	_ = v427
	var v429 int32
	_ = v429
	var v432 int32
	_ = v432
	var v433 int32
	_ = v433
	var v434 int32
	_ = v434
	var v444 int64
	_ = v444
	var v445 int32
	_ = v445
	var v449 int64
	_ = v449
	var v452 int64
	_ = v452
	var v455 int64
	_ = v455
	var v458 int64
	_ = v458
	var v459 int64
	_ = v459
	var v460 int64
	_ = v460
	var v470 int64
	_ = v470
	var v472 int32
	_ = v472
	var v481 int32
	_ = v481
	var v484 int32
	_ = v484
	var v489 int32
	_ = v489
	var v490 int32
	_ = v490
	var v501 int32
	_ = v501
	var v504 int32
	_ = v504
	var v509 int32
	_ = v509
	var v510 int32
	_ = v510
	var v516 int32
	_ = v516
	var v519 int32
	_ = v519
	var v529 int32
	_ = v529
	var v531 int32
	_ = v531
	var v541 int32
	_ = v541
	var v542 int32
	_ = v542
	var v554 int32
	_ = v554
	var v556 int32
	_ = v556
	var v567 int32
	_ = v567
	var v570 int32
	_ = v570
	var v572 int32
	_ = v572
	var v575 int32
	_ = v575
	var v576 int32
	_ = v576
	var v577 int32
	_ = v577
	var v586 int32
	_ = v586
	var v588 int32
	_ = v588
	var v591 int32
	_ = v591
	var v605 int32
	_ = v605
	var v606 int64
	_ = v606
	var v607 int32
	_ = v607
	var v608 int32
	_ = v608
	var v609 int32
	_ = v609
	var v610 int32
	_ = v610
	var v611 int32
	_ = v611
	var v616 int64
	_ = v616
	var v625 int64
	_ = v625
	var v626 int64
	_ = v626
	var v631 int64
	_ = v631
	var v634 int64
	_ = v634
	var v637 int64
	_ = v637
	var v640 int64
	_ = v640
	var v641 int64
	_ = v641
	var v645 int64
	_ = v645
	var v652 int64
	_ = v652
	var v663 int64
	_ = v663
	var v665 int64
	_ = v665
	var v671 int64
	_ = v671
	var v672 int64
	_ = v672
	var v680 int64
	_ = v680
	var v681 int64
	_ = v681
	var v687 int64
	_ = v687
	var v688 int64
	_ = v688
	var v708 int32
	_ = v708
	var v711 int32
	_ = v711
	var v715 int32
	_ = v715
	var v720 int32
	_ = v720
	var v728 int32
	_ = v728
	var v731 int32
	_ = v731
	var v735 int32
	_ = v735
	var v740 int32
	_ = v740
	v21 = m.G0
	v23 = v21 - int32(112)
	m.G0 = v23
	v25 = *(*int64)(unsafe.Add(mBase, uint32(l0)+24))
	v27 = *(*int64)(unsafe.Add(mBase, uint32(l0)+40))
	v29 = F_palloc(m, int32(16))
	mBase = m.M
	v32 = m.ExcPending
	if v32 != 0 {
		return int64(0)
	} else {
		if v25 != int64(9223372036854775807) {
			v35 = int64(-9223372036854775807 - 1)
			if v25 != v35 {
				if v27 == int64(-9223372036854775807-1) {
					v83 = int64(9223372036854775807)
					v84 = int32(2147483647)
					*(*int32)(unsafe.Add(mBase, uint32(v29)+12)) = v84
					*(*int32)(unsafe.Add(mBase, uint32(v29)+8)) = v84
					*(*int64)(unsafe.Add(mBase, uint32(v29))) = v83
					m.G0 = v23 + int32(112)
					return base.I64_extend_i32_u(v29)
				} else {
					if v27 != int64(9223372036854775807) {
						v89 = base.I64_div_s(v25, int64(86400000000))
						if base.Ui64(int64(172799999999)) <= base.Ui64(v25+int64(86399999999)) {
							v97 = v89 * int64(-86400000000)
						} else {
							v97 = int64(0)
						}
						v98 = v97 + v25
						v101 = v98>>(uint(int64(63))%64) + v89
						if v101 < int64(-2451545) {
							F_errstart_cold(m, int32(21), int32(0))
							mBase = m.M
							v728 = m.ExcPending
							if v728 != 0 {
								return int64(0)
							} else {
								F_errcode(m, int32(134217858))
								mBase = m.M
								v731 = m.ExcPending
								if v731 != 0 {
									return int64(0)
								} else {
									F_errmsg(m, int32(_a_F_timestamp_age_0), int32(0))
									mBase = m.M
									v735 = m.ExcPending
									if v735 != 0 {
										return int64(0)
									} else {
										F_errfinish(m, int32(_a_F_timestamp_age_1), int32(_a_F_timestamp_age_2), int32(_a_F_timestamp_age_3))
										mBase = m.M
										v740 = m.ExcPending
										if v740 != 0 {
											return int64(0)
										} else {
											base.Wasm_trap_unreachable()
											for {
											}
										}
									}
								}
							}
						} else {
							v104 = base.I32_wrap_i64(v101)
							v116 = v104 + int32(_a_F_timestamp_age_4)
							v117 = int32(_a_F_timestamp_age_5)
							v118 = base.I32_div_u_s(v116, v117)
							v119 = int32(3)
							v125 = int32(2)
							v130 = base.I32_div_u_s((v118*int32(1073595727)+v116)<<(uint(v125)%32)|v119, v117)
							v133 = v104 + int32(_a_F_timestamp_age_6) + v118*v119 + v130 + int32(_a_F_timestamp_age_7)
							v134 = int32(1461)
							v135 = base.I32_div_u_s(v133, v134)
							v138 = v135*int32(-1461) + v133
							v140 = v138 << (uint(v125) % 32)
							if base.Ui32(v134) <= base.Ui32(v140) {
								v146 = base.I32_rem_u_s(v138+int32(305), int32(365))
								v151 = v146
							} else {
								v150 = base.I32_rem_u_s(v138+int32(306), int32(366))
								v151 = v150
							}
							v153 = base.I32_div_u_s(v140, int32(1461))
							*(*int32)(unsafe.Add(mBase, uint32(v23+int32(88)))) = v153 + v135<<(uint(int32(2))%32) - int32(_a_F_timestamp_age_8)
							v161 = v151 + int32(123)
							v165 = int32(base.Ui32(v161*int32(2141)) >> (uint(int32(16)) % 32))
							*(*int32)(unsafe.Add(mBase, uint32(v23+int32(80)))) = v161 - int32(base.Ui32(v165*int32(_a_F_timestamp_age_9))>>(uint(int32(8))%32))
							v175 = base.I32_rem_u_s(v165+int32(10), int32(12))
							*(*int32)(unsafe.Add(mBase, uint32(v23+int32(84)))) = v175 + int32(1)
							*(*int32)(unsafe.Add(mBase, uint32(v23)+108)) = int32(0)
							*(*int64)(unsafe.Add(mBase, uint32(v23)+100)) = int64(4294967295)
							if v98 < int64(0) {
								v187 = v98 + int64(86400000000)
							} else {
								v187 = v98
							}
							v189 = base.I64_div_s(v187, int64(3600000000))
							*(*uint32)(unsafe.Add(mBase, uint32(v23)+76)) = uint32(v189)
							v194 = base.I64_extend32_s(v189)*int64(-3600000000) + v187
							v196 = base.I64_div_s(v194, int64(60000000))
							*(*uint32)(unsafe.Add(mBase, uint32(v23)+72)) = uint32(v196)
							v201 = base.I64_extend32_s(v196)*int64(-60000000) + v194
							v203 = base.I64_div_s(v201, int64(1000000))
							*(*uint32)(unsafe.Add(mBase, uint32(v23)+68)) = uint32(v203)
							v206 = base.I64_div_s(v27, int64(86400000000))
							if base.Ui64(int64(172799999999)) <= base.Ui64(v27+int64(86399999999)) {
								v214 = v206 * int64(-86400000000)
							} else {
								v214 = int64(0)
							}
							v215 = v214 + v27
							v218 = v215>>(uint(int64(63))%64) + v206
							if v218 < int64(-2451545) {
								F_errstart_cold(m, int32(21), int32(0))
								mBase = m.M
								v728 = m.ExcPending
								if v728 != 0 {
									return int64(0)
								} else {
									F_errcode(m, int32(134217858))
									mBase = m.M
									v731 = m.ExcPending
									if v731 != 0 {
										return int64(0)
									} else {
										F_errmsg(m, int32(_a_F_timestamp_age_0), int32(0))
										mBase = m.M
										v735 = m.ExcPending
										if v735 != 0 {
											return int64(0)
										} else {
											F_errfinish(m, int32(_a_F_timestamp_age_1), int32(_a_F_timestamp_age_2), int32(_a_F_timestamp_age_3))
											mBase = m.M
											v740 = m.ExcPending
											if v740 != 0 {
												return int64(0)
											} else {
												base.Wasm_trap_unreachable()
												for {
												}
											}
										}
									}
								}
							} else {
								v221 = base.I32_wrap_i64(v218)
								v233 = v221 + int32(_a_F_timestamp_age_4)
								v234 = int32(_a_F_timestamp_age_5)
								v235 = base.I32_div_u_s(v233, v234)
								v236 = int32(3)
								v242 = int32(2)
								v247 = base.I32_div_u_s((v235*int32(1073595727)+v233)<<(uint(v242)%32)|v236, v234)
								v250 = v221 + int32(_a_F_timestamp_age_6) + v235*v236 + v247 + int32(_a_F_timestamp_age_7)
								v251 = int32(1461)
								v252 = base.I32_div_u_s(v250, v251)
								v255 = v252*int32(-1461) + v250
								v257 = v255 << (uint(v242) % 32)
								if base.Ui32(v251) <= base.Ui32(v257) {
									v263 = base.I32_rem_u_s(v255+int32(305), int32(365))
									v268 = v263
								} else {
									v267 = base.I32_rem_u_s(v255+int32(306), int32(366))
									v268 = v267
								}
								v270 = base.I32_div_u_s(v257, int32(1461))
								*(*int32)(unsafe.Add(mBase, uint32(v23+int32(44)))) = v270 + v252<<(uint(int32(2))%32) - int32(_a_F_timestamp_age_8)
								v278 = v268 + int32(123)
								v282 = int32(base.Ui32(v278*int32(2141)) >> (uint(int32(16)) % 32))
								*(*int32)(unsafe.Add(mBase, uint32(v23+int32(36)))) = v278 - int32(base.Ui32(v282*int32(_a_F_timestamp_age_9))>>(uint(int32(8))%32))
								v292 = base.I32_rem_u_s(v282+int32(10), int32(12))
								*(*int32)(unsafe.Add(mBase, uint32(v23+int32(40)))) = v292 + int32(1)
								*(*int32)(unsafe.Add(mBase, uint32(v23)+64)) = int32(0)
								*(*int64)(unsafe.Add(mBase, uint32(v23)+56)) = int64(4294967295)
								if v215 < int64(0) {
									v308 = v215 + int64(86400000000)
								} else {
									v308 = v215
								}
								v310 = base.I64_div_s(v308, int64(3600000000))
								v314 = base.I64_extend32_s(v310)*int64(-3600000000) + v308
								v316 = base.I64_div_s(v314, int64(60000000))
								v320 = base.I64_extend32_s(v316)*int64(-60000000) + v314
								v322 = base.I64_div_s(v320, int64(1000000))
								v327 = base.I32_wrap_i64(v203*int64(4293967296)+v201) - base.I32_wrap_i64(v322*int64(4293967296)+v320)
								v328 = base.I32_wrap_i64(v322)
								*(*int32)(unsafe.Add(mBase, uint32(v23)+24)) = v328
								v330 = *(*int32)(unsafe.Add(mBase, uint32(v23)+68))
								v331 = v330 - v328
								v332 = base.I32_wrap_i64(v316)
								*(*int32)(unsafe.Add(mBase, uint32(v23)+28)) = v332
								v334 = *(*int32)(unsafe.Add(mBase, uint32(v23)+72))
								v335 = v334 - v332
								v336 = base.I32_wrap_i64(v310)
								*(*int32)(unsafe.Add(mBase, uint32(v23)+32)) = v336
								v338 = *(*int32)(unsafe.Add(mBase, uint32(v23)+76))
								v340 = base.I64_extend_i32_s(v338 - v336)
								v341 = *(*int32)(unsafe.Add(mBase, uint32(v23)+88))
								v342 = *(*int32)(unsafe.Add(mBase, uint32(v23)+44))
								v343 = v341 - v342
								v344 = *(*int32)(unsafe.Add(mBase, uint32(v23)+84))
								v345 = *(*int32)(unsafe.Add(mBase, uint32(v23)+40))
								v346 = v344 - v345
								v347 = *(*int32)(unsafe.Add(mBase, uint32(v23)+80))
								v348 = *(*int32)(unsafe.Add(mBase, uint32(v23)+36))
								v349 = v347 - v348
								v350 = base.B2i32(v27 <= v25)
								if v350 == int32(0) {
									v353 = int32(0)
									v367 = v353 - v331
									v368 = int64(0) - v340
									v369 = v353 - v335
									v370 = v353 - v346
									v371 = v353 - v327
									v372 = v353 - v349
									v373 = v353 - v343
								} else {
									v367 = v331
									v368 = v340
									v369 = v335
									v370 = v346
									v371 = v327
									v372 = v349
									v373 = v343
								}
								if v371 < int32(0) {
									v376 = int32(-1000000)
									if base.Ui32(v371) <= base.Ui32(v376) {
										v379 = v376
									} else {
										v379 = v371
									}
									v381 = base.B2i32(base.Ui32(v371) < base.Ui32(int32(-1000000)))
									v384 = int32(_a_F_timestamp_age_10)
									v385 = base.I32_div_u_s(v379-(v371+v381), v384)
									v386 = v385 + v381
									v395 = v367 + (v386 ^ int32(-1))
									v396 = v371 + v386*v384 + v384
								} else {
									v395 = v367
									v396 = v371
								}
								if v395 < int32(0) {
									v400 = int32(-60)
									if base.Ui32(v395) <= base.Ui32(v400) {
										v403 = v400
									} else {
										v403 = v395
									}
									v405 = base.B2i32(base.Ui32(v395) < base.Ui32(int32(-60)))
									v408 = int32(60)
									v409 = base.I32_div_u_s(v403-(v395+v405), v408)
									v410 = v409 + v405
									v419 = v395 + v410*v408 + v408
									v420 = v369 + (v410 ^ int32(-1))
								} else {
									v419 = v395
									v420 = v369
								}
								if v420 < int32(0) {
									v424 = int32(-60)
									if base.Ui32(v420) <= base.Ui32(v424) {
										v427 = v424
									} else {
										v427 = v420
									}
									v429 = base.B2i32(base.Ui32(v420) < base.Ui32(int32(-60)))
									v432 = int32(60)
									v433 = base.I32_div_u_s(v427-(v420+v429), v432)
									v434 = v433 + v429
									v444 = v368 + base.I64_extend_i32_s(v434^int32(-1))
									v445 = v420 + v434*v432 + v432
								} else {
									v444 = v368
									v445 = v420
								}
								if v444 < int64(0) {
									v449 = int64(-24)
									if base.Ui64(v444) <= base.Ui64(v449) {
										v452 = v449
									} else {
										v452 = v444
									}
									v455 = base.I64_extend_i32_u(base.B2i32(base.Ui64(v444) < base.Ui64(int64(-24))))
									v458 = int64(24)
									v459 = base.I64_div_u_s(v452-(v444+v455), v458)
									v460 = v459 + v455
									v470 = v444 + v460*v458 + v458
									v472 = v372 + (base.I32_wrap_i64(v460) ^ int32(-1))
								} else {
									v470 = v444
									v472 = v372
								}
								if v472 < int32(0) {
									v481 = int32(0)
									v484 = base.I32_rem_s(v341, int32(100))
									v489 = base.I32_rem_s(v341, int32(400))
									if v489 != 0 {
										v490 = base.B2i32(v484 != v481)
									} else {
										v490 = int32(1)
									}
									v501 = int32(0)
									v504 = base.I32_rem_s(v342, int32(100))
									v509 = base.I32_rem_s(v342, int32(400))
									if v509 != 0 {
										v510 = base.B2i32(v504 != v501)
									} else {
										v510 = int32(1)
									}
									if v25 < v27 {
										v516 = v344<<(uint(int32(2))%32) + int32(_a_F_timestamp_age_11) + base.B2i32(v341&int32(3) == v481)&v490*int32(52)
									} else {
										v516 = v345<<(uint(int32(2))%32) + int32(_a_F_timestamp_age_11) + base.B2i32(v342&int32(3) == v501)&v510*int32(52)
									}
									v519 = *(*int32)(unsafe.Add(mBase, uint32(v516-int32(4))))
									v529 = v370
									v531 = v472
									for {
										v541 = v529 - int32(1)
										v542 = v531 + v519
										if v542 < int32(0) {
											v529 = v541
											v531 = v542
											continue
										} else {
											break
										}
										break
									}
									v554 = v541
									v556 = v542
								} else {
									v554 = v370
									v556 = v472
								}
								if v554 < int32(0) {
									v567 = int32(-12)
									if base.Ui32(v554) <= base.Ui32(v567) {
										v570 = v567
									} else {
										v570 = v554
									}
									v572 = base.B2i32(base.Ui32(v554) < base.Ui32(int32(-12)))
									v575 = int32(12)
									v576 = base.I32_div_u_s(v570-(v554+v572), v575)
									v577 = v576 + v572
									v586 = v554 + v577*v575 + v575
									v588 = v373 + (v577 ^ int32(-1))
								} else {
									v586 = v554
									v588 = v373
								}
								if v350 == int32(0) {
									v591 = int32(0)
									v605 = v591 - v419
									v606 = int64(0) - v470
									v607 = v591 - v445
									v608 = v591 - v586
									v609 = v591 - v396
									v610 = v591 - v556
									v611 = v591 - v588
								} else {
									v605 = v419
									v606 = v470
									v607 = v445
									v608 = v586
									v609 = v396
									v610 = v556
									v611 = v588
								}
								v616 = base.I64_extend_i32_s(v608) + base.I64_extend_i32_s(v611)*int64(12)
								if base.Ui64(v616-int64(2147483648)) < base.Ui64(int64(-4294967296)) {
									F_errstart_cold(m, int32(21), int32(0))
									mBase = m.M
									v708 = m.ExcPending
									if v708 != 0 {
										return int64(0)
									} else {
										F_errcode(m, int32(134217858))
										mBase = m.M
										v711 = m.ExcPending
										if v711 != 0 {
											return int64(0)
										} else {
											F_errmsg(m, int32(_a_F_timestamp_age_12), int32(0))
											mBase = m.M
											v715 = m.ExcPending
											if v715 != 0 {
												return int64(0)
											} else {
												F_errfinish(m, int32(_a_F_timestamp_age_1), int32(_a_F_timestamp_age_13), int32(_a_F_timestamp_age_3))
												mBase = m.M
												v720 = m.ExcPending
												if v720 != 0 {
													return int64(0)
												} else {
													base.Wasm_trap_unreachable()
													for {
													}
												}
											}
										}
									}
								} else {
									*(*int32)(unsafe.Add(mBase, uint32(v29)+8)) = v610
									*(*uint32)(unsafe.Add(mBase, uint32(v29)+12)) = uint32(v616)
									v625 = int64(3600000000)
									v626 = int64(0)
									v631 = int64(32)
									v634 = int64(base.Ui64(v606) >> (uint(v631) % 64))
									v637 = int64(4294967295)
									v640 = v606 & v637
									v641 = v625 * v640
									v645 = int64(base.Ui64(v641)>>(uint(v631)%64)) + v625*v634
									v652 = v640*v626 + v645&v637
									*(*int64)(unsafe.Add(mBase, uint32(v23)+8)) = v606*v626 + v606>>(uint(int64(63))%64)*v625 + v626*v634 + int64(base.Ui64(v645)>>(uint(v631)%64)) + int64(base.Ui64(v652)>>(uint(v631)%64))
									*(*int64)(unsafe.Add(mBase, uint32(v23))) = v641&v637 | v652<<(uint(v631)%64)
									v663 = *(*int64)(unsafe.Add(mBase, uint32(v23)))
									*(*int64)(unsafe.Add(mBase, uint32(v29))) = v663
									v665 = *(*int64)(unsafe.Add(mBase, uint32(v23)+8))
									if v665 != v663>>(uint(int64(63))%64) {
										F_errstart_cold(m, int32(21), int32(0))
										mBase = m.M
										v708 = m.ExcPending
										if v708 != 0 {
											return int64(0)
										} else {
											F_errcode(m, int32(134217858))
											mBase = m.M
											v711 = m.ExcPending
											if v711 != 0 {
												return int64(0)
											} else {
												F_errmsg(m, int32(_a_F_timestamp_age_12), int32(0))
												mBase = m.M
												v715 = m.ExcPending
												if v715 != 0 {
													return int64(0)
												} else {
													F_errfinish(m, int32(_a_F_timestamp_age_1), int32(_a_F_timestamp_age_13), int32(_a_F_timestamp_age_3))
													mBase = m.M
													v720 = m.ExcPending
													if v720 != 0 {
														return int64(0)
													} else {
														base.Wasm_trap_unreachable()
														for {
														}
													}
												}
											}
										}
									} else {
										v671 = base.I64_extend_i32_s(v607) * int64(60000000)
										v672 = v663 + v671
										*(*int64)(unsafe.Add(mBase, uint32(v29))) = v672
										if base.B2i32(v671 < int64(0))^base.B2i32(v672 < v663) != 0 {
											F_errstart_cold(m, int32(21), int32(0))
											mBase = m.M
											v708 = m.ExcPending
											if v708 != 0 {
												return int64(0)
											} else {
												F_errcode(m, int32(134217858))
												mBase = m.M
												v711 = m.ExcPending
												if v711 != 0 {
													return int64(0)
												} else {
													F_errmsg(m, int32(_a_F_timestamp_age_12), int32(0))
													mBase = m.M
													v715 = m.ExcPending
													if v715 != 0 {
														return int64(0)
													} else {
														F_errfinish(m, int32(_a_F_timestamp_age_1), int32(_a_F_timestamp_age_13), int32(_a_F_timestamp_age_3))
														mBase = m.M
														v720 = m.ExcPending
														if v720 != 0 {
															return int64(0)
														} else {
															base.Wasm_trap_unreachable()
															for {
															}
														}
													}
												}
											}
										} else {
											v680 = base.I64_extend_i32_s(v605) * int64(1000000)
											v681 = v672 + v680
											*(*int64)(unsafe.Add(mBase, uint32(v29))) = v681
											if base.B2i32(v680 < int64(0))^base.B2i32(v681 < v672) != 0 {
												F_errstart_cold(m, int32(21), int32(0))
												mBase = m.M
												v708 = m.ExcPending
												if v708 != 0 {
													return int64(0)
												} else {
													F_errcode(m, int32(134217858))
													mBase = m.M
													v711 = m.ExcPending
													if v711 != 0 {
														return int64(0)
													} else {
														F_errmsg(m, int32(_a_F_timestamp_age_12), int32(0))
														mBase = m.M
														v715 = m.ExcPending
														if v715 != 0 {
															return int64(0)
														} else {
															F_errfinish(m, int32(_a_F_timestamp_age_1), int32(_a_F_timestamp_age_13), int32(_a_F_timestamp_age_3))
															mBase = m.M
															v720 = m.ExcPending
															if v720 != 0 {
																return int64(0)
															} else {
																base.Wasm_trap_unreachable()
																for {
																}
															}
														}
													}
												}
											} else {
												v687 = base.I64_extend_i32_s(v609)
												v688 = v681 + v687
												*(*int64)(unsafe.Add(mBase, uint32(v29))) = v688
												if base.B2i32(v687 < int64(0))^base.B2i32(v688 < v681) != 0 {
													F_errstart_cold(m, int32(21), int32(0))
													mBase = m.M
													v708 = m.ExcPending
													if v708 != 0 {
														return int64(0)
													} else {
														F_errcode(m, int32(134217858))
														mBase = m.M
														v711 = m.ExcPending
														if v711 != 0 {
															return int64(0)
														} else {
															F_errmsg(m, int32(_a_F_timestamp_age_12), int32(0))
															mBase = m.M
															v715 = m.ExcPending
															if v715 != 0 {
																return int64(0)
															} else {
																F_errfinish(m, int32(_a_F_timestamp_age_1), int32(_a_F_timestamp_age_13), int32(_a_F_timestamp_age_3))
																mBase = m.M
																v720 = m.ExcPending
																if v720 != 0 {
																	return int64(0)
																} else {
																	base.Wasm_trap_unreachable()
																	for {
																	}
																}
															}
														}
													}
												} else {
													if base.B2i32(v610 != int32(2147483647))|base.B2i32(v616 != int64(2147483647))|base.B2i32(v688 != int64(9223372036854775807)) != 0 {
														m.G0 = v23 + int32(112)
														return base.I64_extend_i32_u(v29)
													} else {
														F_errstart_cold(m, int32(21), int32(0))
														mBase = m.M
														v708 = m.ExcPending
														if v708 != 0 {
															return int64(0)
														} else {
															F_errcode(m, int32(134217858))
															mBase = m.M
															v711 = m.ExcPending
															if v711 != 0 {
																return int64(0)
															} else {
																F_errmsg(m, int32(_a_F_timestamp_age_12), int32(0))
																mBase = m.M
																v715 = m.ExcPending
																if v715 != 0 {
																	return int64(0)
																} else {
																	F_errfinish(m, int32(_a_F_timestamp_age_1), int32(_a_F_timestamp_age_13), int32(_a_F_timestamp_age_3))
																	mBase = m.M
																	v720 = m.ExcPending
																	if v720 != 0 {
																		return int64(0)
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
										}
									}
								}
							}
						}
					} else {
						v83 = v35
						v84 = int32(-2147483648)
						*(*int32)(unsafe.Add(mBase, uint32(v29)+12)) = v84
						*(*int32)(unsafe.Add(mBase, uint32(v29)+8)) = v84
						*(*int64)(unsafe.Add(mBase, uint32(v29))) = v83
						m.G0 = v23 + int32(112)
						return base.I64_extend_i32_u(v29)
					}
				}
			} else {
				if v27 != int64(-9223372036854775807-1) {
					v83 = v35
					v84 = int32(-2147483648)
					*(*int32)(unsafe.Add(mBase, uint32(v29)+12)) = v84
					*(*int32)(unsafe.Add(mBase, uint32(v29)+8)) = v84
					*(*int64)(unsafe.Add(mBase, uint32(v29))) = v83
					m.G0 = v23 + int32(112)
					return base.I64_extend_i32_u(v29)
				} else {
					F_errstart_cold(m, int32(21), int32(0))
					mBase = m.M
					v44 = m.ExcPending
					if v44 != 0 {
						return int64(0)
					} else {
						F_errcode(m, int32(134217858))
						mBase = m.M
						v47 = m.ExcPending
						if v47 != 0 {
							return int64(0)
						} else {
							F_errmsg(m, int32(_a_F_timestamp_age_12), int32(0))
							mBase = m.M
							v51 = m.ExcPending
							if v51 != 0 {
								return int64(0)
							} else {
								F_errfinish(m, int32(_a_F_timestamp_age_1), int32(_a_F_timestamp_age_14), int32(_a_F_timestamp_age_3))
								mBase = m.M
								v56 = m.ExcPending
								if v56 != 0 {
									return int64(0)
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
		} else {
			if v27 != int64(9223372036854775807) {
				v83 = int64(9223372036854775807)
				v84 = int32(2147483647)
				*(*int32)(unsafe.Add(mBase, uint32(v29)+12)) = v84
				*(*int32)(unsafe.Add(mBase, uint32(v29)+8)) = v84
				*(*int64)(unsafe.Add(mBase, uint32(v29))) = v83
				m.G0 = v23 + int32(112)
				return base.I64_extend_i32_u(v29)
			} else {
				F_errstart_cold(m, int32(21), int32(0))
				mBase = m.M
				v63 = m.ExcPending
				if v63 != 0 {
					return int64(0)
				} else {
					F_errcode(m, int32(134217858))
					mBase = m.M
					v66 = m.ExcPending
					if v66 != 0 {
						return int64(0)
					} else {
						F_errmsg(m, int32(_a_F_timestamp_age_12), int32(0))
						mBase = m.M
						v70 = m.ExcPending
						if v70 != 0 {
							return int64(0)
						} else {
							F_errfinish(m, int32(_a_F_timestamp_age_1), int32(_a_F_timestamp_age_15), int32(_a_F_timestamp_age_3))
							mBase = m.M
							v75 = m.ExcPending
							if v75 != 0 {
								return int64(0)
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
}
func F_timestamp_bin(m *base.Module, l0 int32) int64 {
	var v10 int64
	_ = v10
	var v13 int32
	_ = v13
	v10 = Fn14381(m, l0, int32(_a_F_timestamp_bin_0), int32(_a_F_timestamp_bin_1), int32(_a_F_timestamp_bin_2), int32(_a_F_timestamp_bin_3), int32(_a_F_timestamp_bin_4), int32(_a_F_timestamp_bin_5), int32(_a_F_timestamp_bin_6), int32(_a_F_timestamp_bin_7))
	v13 = m.ExcPending
	if v13 != 0 {
		return int64(0)
	} else {
		return v10
	}
}
func F_timestamp_lt_date(m *base.Module, l0 int32) int64 {
	mBase := m.M
	_ = mBase
	var v2 int32
	_ = v2
	var v3 int64
	_ = v3
	var v4 int32
	_ = v4
	var v7 int32
	_ = v7
	v2 = *(*int32)(unsafe.Add(mBase, uint32(l0)+40))
	v3 = *(*int64)(unsafe.Add(mBase, uint32(l0)+24))
	v4 = F_date_cmp_timestamp_internal(m, v2, v3)
	mBase = m.M
	v7 = m.ExcPending
	if v7 != 0 {
		return int64(0)
	} else {
		return base.I64_extend_i32_u(base.B2i32(int32(0) < v4))
	}
}
func F_timestamp_part(m *base.Module, l0 int32) int64 {
	var v3 int64
	_ = v3
	var v6 int32
	_ = v6
	v3 = F_timestamp_part_common(m, l0, int32(0))
	v6 = m.ExcPending
	if v6 != 0 {
		return int64(0)
	} else {
		return v3
	}
}
func F_timestamp_random(m *base.Module, l0 int32) int64 {
	var v5 int64
	_ = v5
	var v8 int32
	_ = v8
	v5 = Fn14383(m, l0, int32(_a_F_timestamp_random_0), int32(240), int32(235))
	v8 = m.ExcPending
	if v8 != 0 {
		return int64(0)
	} else {
		return v5
	}
}
