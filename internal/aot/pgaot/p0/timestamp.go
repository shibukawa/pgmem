package p0

import (
	base "github.com/shibukawa/pgmem/internal/aot/pgaot/base"
	"unsafe"
)

func F_TimestampDifferenceExceeds(m *base.Module, l0 int64, l1 int64, l2 int32) int32 {
	return base.B2i32(base.I64_extend_i32_s(l2)*int64(1000) <= l1-l0)
}
func F_extract_timestamp(m *base.Module, l0 int32) int32 {
	var v3 int32
	_ = v3
	var v6 int32
	_ = v6
	v3 = F_timestamp_part_common(m, l0, int32(1))
	v6 = m.ExcPending
	if v6 != 0 {
		return int32(0)
	} else {
		return v3
	}
}
func F_timestamp_age(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v21 int32
	_ = v21
	var v23 int32
	_ = v23
	var v25 int32
	_ = v25
	var v26 int64
	_ = v26
	var v28 int32
	_ = v28
	var v29 int64
	_ = v29
	var v31 int32
	_ = v31
	var v34 int32
	_ = v34
	var v37 int64
	_ = v37
	var v46 int32
	_ = v46
	var v49 int32
	_ = v49
	var v53 int32
	_ = v53
	var v58 int32
	_ = v58
	var v65 int32
	_ = v65
	var v68 int32
	_ = v68
	var v72 int32
	_ = v72
	var v77 int32
	_ = v77
	var v85 int64
	_ = v85
	var v86 int32
	_ = v86
	var v91 int64
	_ = v91
	var v99 int64
	_ = v99
	var v100 int64
	_ = v100
	var v103 int64
	_ = v103
	var v106 int32
	_ = v106
	var v118 int32
	_ = v118
	var v119 int32
	_ = v119
	var v120 int32
	_ = v120
	var v121 int32
	_ = v121
	var v127 int32
	_ = v127
	var v132 int32
	_ = v132
	var v135 int32
	_ = v135
	var v136 int32
	_ = v136
	var v137 int32
	_ = v137
	var v140 int32
	_ = v140
	var v142 int32
	_ = v142
	var v148 int32
	_ = v148
	var v152 int32
	_ = v152
	var v153 int32
	_ = v153
	var v155 int32
	_ = v155
	var v163 int32
	_ = v163
	var v167 int32
	_ = v167
	var v177 int32
	_ = v177
	var v189 int64
	_ = v189
	var v191 int64
	_ = v191
	var v196 int64
	_ = v196
	var v198 int64
	_ = v198
	var v203 int64
	_ = v203
	var v205 int64
	_ = v205
	var v208 int64
	_ = v208
	var v216 int64
	_ = v216
	var v217 int64
	_ = v217
	var v220 int64
	_ = v220
	var v223 int32
	_ = v223
	var v235 int32
	_ = v235
	var v236 int32
	_ = v236
	var v237 int32
	_ = v237
	var v238 int32
	_ = v238
	var v244 int32
	_ = v244
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
	var v259 int32
	_ = v259
	var v265 int32
	_ = v265
	var v269 int32
	_ = v269
	var v270 int32
	_ = v270
	var v272 int32
	_ = v272
	var v280 int32
	_ = v280
	var v284 int32
	_ = v284
	var v294 int32
	_ = v294
	var v310 int64
	_ = v310
	var v312 int64
	_ = v312
	var v316 int64
	_ = v316
	var v318 int64
	_ = v318
	var v322 int64
	_ = v322
	var v324 int64
	_ = v324
	var v329 int32
	_ = v329
	var v330 int32
	_ = v330
	var v332 int32
	_ = v332
	var v333 int32
	_ = v333
	var v334 int32
	_ = v334
	var v336 int32
	_ = v336
	var v337 int32
	_ = v337
	var v338 int32
	_ = v338
	var v340 int32
	_ = v340
	var v342 int64
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
	var v351 int32
	_ = v351
	var v352 int32
	_ = v352
	var v355 int32
	_ = v355
	var v369 int64
	_ = v369
	var v370 int32
	_ = v370
	var v371 int32
	_ = v371
	var v372 int32
	_ = v372
	var v373 int32
	_ = v373
	var v374 int32
	_ = v374
	var v375 int32
	_ = v375
	var v378 int32
	_ = v378
	var v381 int32
	_ = v381
	var v383 int32
	_ = v383
	var v386 int32
	_ = v386
	var v387 int32
	_ = v387
	var v388 int32
	_ = v388
	var v397 int32
	_ = v397
	var v398 int32
	_ = v398
	var v402 int32
	_ = v402
	var v405 int32
	_ = v405
	var v407 int32
	_ = v407
	var v410 int32
	_ = v410
	var v411 int32
	_ = v411
	var v412 int32
	_ = v412
	var v421 int32
	_ = v421
	var v422 int32
	_ = v422
	var v426 int32
	_ = v426
	var v429 int32
	_ = v429
	var v431 int32
	_ = v431
	var v434 int32
	_ = v434
	var v435 int32
	_ = v435
	var v436 int32
	_ = v436
	var v446 int64
	_ = v446
	var v447 int32
	_ = v447
	var v451 int64
	_ = v451
	var v454 int64
	_ = v454
	var v457 int64
	_ = v457
	var v460 int64
	_ = v460
	var v461 int64
	_ = v461
	var v462 int64
	_ = v462
	var v472 int64
	_ = v472
	var v474 int32
	_ = v474
	var v478 int32
	_ = v478
	var v480 int32
	_ = v480
	var v484 int32
	_ = v484
	var v498 int32
	_ = v498
	var v500 int32
	_ = v500
	var v504 int32
	_ = v504
	var v518 int32
	_ = v518
	var v521 int32
	_ = v521
	var v530 int32
	_ = v530
	var v534 int32
	_ = v534
	var v543 int32
	_ = v543
	var v544 int32
	_ = v544
	var v555 int32
	_ = v555
	var v559 int32
	_ = v559
	var v569 int32
	_ = v569
	var v572 int32
	_ = v572
	var v574 int32
	_ = v574
	var v577 int32
	_ = v577
	var v578 int32
	_ = v578
	var v579 int32
	_ = v579
	var v588 int32
	_ = v588
	var v590 int32
	_ = v590
	var v593 int32
	_ = v593
	var v607 int64
	_ = v607
	var v608 int32
	_ = v608
	var v609 int32
	_ = v609
	var v610 int32
	_ = v610
	var v611 int32
	_ = v611
	var v612 int32
	_ = v612
	var v613 int32
	_ = v613
	var v618 int64
	_ = v618
	var v626 int32
	_ = v626
	var v629 int64
	_ = v629
	var v630 int64
	_ = v630
	var v635 int64
	_ = v635
	var v638 int64
	_ = v638
	var v641 int64
	_ = v641
	var v644 int64
	_ = v644
	var v645 int64
	_ = v645
	var v649 int64
	_ = v649
	var v656 int64
	_ = v656
	var v667 int64
	_ = v667
	var v669 int64
	_ = v669
	var v675 int64
	_ = v675
	var v676 int64
	_ = v676
	var v684 int64
	_ = v684
	var v685 int64
	_ = v685
	var v691 int64
	_ = v691
	var v692 int64
	_ = v692
	var v710 int32
	_ = v710
	var v713 int32
	_ = v713
	var v717 int32
	_ = v717
	var v722 int32
	_ = v722
	var v730 int32
	_ = v730
	var v733 int32
	_ = v733
	var v737 int32
	_ = v737
	var v742 int32
	_ = v742
	v21 = m.G0
	v23 = v21 - int32(112)
	m.G0 = v23
	v25 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v26 = *(*int64)(unsafe.Add(mBase, uint32(v25)))
	v28 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
	v29 = *(*int64)(unsafe.Add(mBase, uint32(v28)))
	v31 = F_palloc(m, int32(16))
	mBase = m.M
	v34 = m.ExcPending
	if v34 != 0 {
		return int32(0)
	} else {
		if v26 != int64(9223372036854775807) {
			v37 = int64(-9223372036854775807 - 1)
			if v26 != v37 {
				if v29 == int64(-9223372036854775807-1) {
					v85 = int64(9223372036854775807)
					v86 = int32(2147483647)
					*(*int32)(unsafe.Add(mBase, uint32(v31)+12)) = v86
					*(*int32)(unsafe.Add(mBase, uint32(v31)+8)) = v86
					*(*int64)(unsafe.Add(mBase, uint32(v31))) = v85
					m.G0 = v23 + int32(112)
					return v31
				} else {
					if v29 != int64(9223372036854775807) {
						v91 = base.I64_div_s(v26, int64(86400000000))
						if base.Ui64(int64(172799999999)) <= base.Ui64(v26+int64(86399999999)) {
							v99 = v91 * int64(-86400000000)
						} else {
							v99 = int64(0)
						}
						v100 = v99 + v26
						v103 = v100>>(uint(int64(63))%64) + v91
						if v103 < int64(-2451545) {
							F_errstart_cold(m, int32(21), int32(0))
							mBase = m.M
							v730 = m.ExcPending
							if v730 != 0 {
								return int32(0)
							} else {
								F_errcode(m, int32(134217858))
								mBase = m.M
								v733 = m.ExcPending
								if v733 != 0 {
									return int32(0)
								} else {
									F_errmsg(m, int32(418494), int32(0))
									mBase = m.M
									v737 = m.ExcPending
									if v737 != 0 {
										return int32(0)
									} else {
										F_errfinish(m, int32(515856), int32(4439), int32(424820))
										mBase = m.M
										v742 = m.ExcPending
										if v742 != 0 {
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
							v106 = base.I32_wrap_i64(v103)
							v118 = v106 + int32(2483589)
							v119 = int32(146097)
							v120 = base.I32_div_u_s(v118, v119)
							v121 = int32(3)
							v127 = int32(2)
							v132 = base.I32_div_u_s((v120*int32(1073595727)+v118)<<(uint(v127)%32)|v121, v119)
							v135 = v106 + int32(2451545) + v120*v121 + v132 + int32(32104)
							v136 = int32(1461)
							v137 = base.I32_div_u_s(v135, v136)
							v140 = v137*int32(-1461) + v135
							v142 = v140 << (uint(v127) % 32)
							if base.Ui32(v136) <= base.Ui32(v142) {
								v148 = base.I32_rem_u_s(v140+int32(305), int32(365))
								v153 = v148
							} else {
								v152 = base.I32_rem_u_s(v140+int32(306), int32(366))
								v153 = v152
							}
							v155 = base.I32_div_u_s(v142, int32(1461))
							*(*int32)(unsafe.Add(mBase, uint32(v23+int32(88)))) = v155 + v137<<(uint(int32(2))%32) - int32(4800)
							v163 = v153 + int32(123)
							v167 = int32(base.Ui32(v163*int32(2141)) >> (uint(int32(16)) % 32))
							*(*int32)(unsafe.Add(mBase, uint32(v23+int32(80)))) = v163 - int32(base.Ui32(v167*int32(7834))>>(uint(int32(8))%32))
							v177 = base.I32_rem_u_s(v167+int32(10), int32(12))
							*(*int32)(unsafe.Add(mBase, uint32(v23+int32(84)))) = v177 + int32(1)
							*(*int32)(unsafe.Add(mBase, uint32(v23)+108)) = int32(0)
							*(*int64)(unsafe.Add(mBase, uint32(v23)+100)) = int64(4294967295)
							if v100 < int64(0) {
								v189 = v100 + int64(86400000000)
							} else {
								v189 = v100
							}
							v191 = base.I64_div_s(v189, int64(3600000000))
							*(*uint32)(unsafe.Add(mBase, uint32(v23)+76)) = uint32(v191)
							v196 = base.I64_extend32_s(v191)*int64(-3600000000) + v189
							v198 = base.I64_div_s(v196, int64(60000000))
							*(*uint32)(unsafe.Add(mBase, uint32(v23)+72)) = uint32(v198)
							v203 = base.I64_extend32_s(v198)*int64(-60000000) + v196
							v205 = base.I64_div_s(v203, int64(1000000))
							*(*uint32)(unsafe.Add(mBase, uint32(v23)+68)) = uint32(v205)
							v208 = base.I64_div_s(v29, int64(86400000000))
							if base.Ui64(int64(172799999999)) <= base.Ui64(v29+int64(86399999999)) {
								v216 = v208 * int64(-86400000000)
							} else {
								v216 = int64(0)
							}
							v217 = v216 + v29
							v220 = v217>>(uint(int64(63))%64) + v208
							if v220 < int64(-2451545) {
								F_errstart_cold(m, int32(21), int32(0))
								mBase = m.M
								v730 = m.ExcPending
								if v730 != 0 {
									return int32(0)
								} else {
									F_errcode(m, int32(134217858))
									mBase = m.M
									v733 = m.ExcPending
									if v733 != 0 {
										return int32(0)
									} else {
										F_errmsg(m, int32(418494), int32(0))
										mBase = m.M
										v737 = m.ExcPending
										if v737 != 0 {
											return int32(0)
										} else {
											F_errfinish(m, int32(515856), int32(4439), int32(424820))
											mBase = m.M
											v742 = m.ExcPending
											if v742 != 0 {
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
								v223 = base.I32_wrap_i64(v220)
								v235 = v223 + int32(2483589)
								v236 = int32(146097)
								v237 = base.I32_div_u_s(v235, v236)
								v238 = int32(3)
								v244 = int32(2)
								v249 = base.I32_div_u_s((v237*int32(1073595727)+v235)<<(uint(v244)%32)|v238, v236)
								v252 = v223 + int32(2451545) + v237*v238 + v249 + int32(32104)
								v253 = int32(1461)
								v254 = base.I32_div_u_s(v252, v253)
								v257 = v254*int32(-1461) + v252
								v259 = v257 << (uint(v244) % 32)
								if base.Ui32(v253) <= base.Ui32(v259) {
									v265 = base.I32_rem_u_s(v257+int32(305), int32(365))
									v270 = v265
								} else {
									v269 = base.I32_rem_u_s(v257+int32(306), int32(366))
									v270 = v269
								}
								v272 = base.I32_div_u_s(v259, int32(1461))
								*(*int32)(unsafe.Add(mBase, uint32(v23+int32(44)))) = v272 + v254<<(uint(int32(2))%32) - int32(4800)
								v280 = v270 + int32(123)
								v284 = int32(base.Ui32(v280*int32(2141)) >> (uint(int32(16)) % 32))
								*(*int32)(unsafe.Add(mBase, uint32(v23+int32(36)))) = v280 - int32(base.Ui32(v284*int32(7834))>>(uint(int32(8))%32))
								v294 = base.I32_rem_u_s(v284+int32(10), int32(12))
								*(*int32)(unsafe.Add(mBase, uint32(v23+int32(40)))) = v294 + int32(1)
								*(*int32)(unsafe.Add(mBase, uint32(v23)+64)) = int32(0)
								*(*int64)(unsafe.Add(mBase, uint32(v23)+56)) = int64(4294967295)
								if v217 < int64(0) {
									v310 = v217 + int64(86400000000)
								} else {
									v310 = v217
								}
								v312 = base.I64_div_s(v310, int64(3600000000))
								v316 = base.I64_extend32_s(v312)*int64(-3600000000) + v310
								v318 = base.I64_div_s(v316, int64(60000000))
								v322 = base.I64_extend32_s(v318)*int64(-60000000) + v316
								v324 = base.I64_div_s(v322, int64(1000000))
								v329 = base.I32_wrap_i64(v205*int64(4293967296)+v203) - base.I32_wrap_i64(v324*int64(4293967296)+v322)
								v330 = base.I32_wrap_i64(v324)
								*(*int32)(unsafe.Add(mBase, uint32(v23)+24)) = v330
								v332 = *(*int32)(unsafe.Add(mBase, uint32(v23)+68))
								v333 = v332 - v330
								v334 = base.I32_wrap_i64(v318)
								*(*int32)(unsafe.Add(mBase, uint32(v23)+28)) = v334
								v336 = *(*int32)(unsafe.Add(mBase, uint32(v23)+72))
								v337 = v336 - v334
								v338 = base.I32_wrap_i64(v312)
								*(*int32)(unsafe.Add(mBase, uint32(v23)+32)) = v338
								v340 = *(*int32)(unsafe.Add(mBase, uint32(v23)+76))
								v342 = base.I64_extend_i32_s(v340 - v338)
								v343 = *(*int32)(unsafe.Add(mBase, uint32(v23)+88))
								v344 = *(*int32)(unsafe.Add(mBase, uint32(v23)+44))
								v345 = v343 - v344
								v346 = *(*int32)(unsafe.Add(mBase, uint32(v23)+84))
								v347 = *(*int32)(unsafe.Add(mBase, uint32(v23)+40))
								v348 = v346 - v347
								v349 = *(*int32)(unsafe.Add(mBase, uint32(v23)+80))
								v350 = *(*int32)(unsafe.Add(mBase, uint32(v23)+36))
								v351 = v349 - v350
								v352 = base.B2i32(v29 <= v26)
								if v352 == int32(0) {
									v355 = int32(0)
									v369 = int64(0) - v342
									v370 = v355 - v348
									v371 = v355 - v333
									v372 = v355 - v337
									v373 = v355 - v329
									v374 = v355 - v351
									v375 = v355 - v345
								} else {
									v369 = v342
									v370 = v348
									v371 = v333
									v372 = v337
									v373 = v329
									v374 = v351
									v375 = v345
								}
								if v373 < int32(0) {
									v378 = int32(-1000000)
									if base.Ui32(v373) <= base.Ui32(v378) {
										v381 = v378
									} else {
										v381 = v373
									}
									v383 = base.B2i32(base.Ui32(v373) < base.Ui32(int32(-1000000)))
									v386 = int32(1000000)
									v387 = base.I32_div_u_s(v381-(v373+v383), v386)
									v388 = v387 + v383
									v397 = v371 + (v388 ^ int32(-1))
									v398 = v373 + v388*v386 + v386
								} else {
									v397 = v371
									v398 = v373
								}
								if v397 < int32(0) {
									v402 = int32(-60)
									if base.Ui32(v397) <= base.Ui32(v402) {
										v405 = v402
									} else {
										v405 = v397
									}
									v407 = base.B2i32(base.Ui32(v397) < base.Ui32(int32(-60)))
									v410 = int32(60)
									v411 = base.I32_div_u_s(v405-(v397+v407), v410)
									v412 = v411 + v407
									v421 = v397 + v412*v410 + v410
									v422 = v372 + (v412 ^ int32(-1))
								} else {
									v421 = v397
									v422 = v372
								}
								if v422 < int32(0) {
									v426 = int32(-60)
									if base.Ui32(v422) <= base.Ui32(v426) {
										v429 = v426
									} else {
										v429 = v422
									}
									v431 = base.B2i32(base.Ui32(v422) < base.Ui32(int32(-60)))
									v434 = int32(60)
									v435 = base.I32_div_u_s(v429-(v422+v431), v434)
									v436 = v435 + v431
									v446 = v369 + base.I64_extend_i32_s(v436^int32(-1))
									v447 = v422 + v436*v434 + v434
								} else {
									v446 = v369
									v447 = v422
								}
								if v446 < int64(0) {
									v451 = int64(-24)
									if base.Ui64(v446) <= base.Ui64(v451) {
										v454 = v451
									} else {
										v454 = v446
									}
									v457 = base.I64_extend_i32_u(base.B2i32(base.Ui64(v446) < base.Ui64(int64(-24))))
									v460 = int64(24)
									v461 = base.I64_div_u_s(v454-(v446+v457), v460)
									v462 = v461 + v457
									v472 = v446 + v462*v460 + v460
									v474 = v374 + (base.I32_wrap_i64(v462) ^ int32(-1))
								} else {
									v472 = v446
									v474 = v374
								}
								if v474 < int32(0) {
									v478 = base.I32_rem_s(v343, int32(400))
									if v478 != 0 {
										v480 = base.I32_rem_s(v343, int32(100))
										v484 = base.B2i32(v480 != int32(0))
									} else {
										v484 = int32(1)
									}
									v498 = base.I32_rem_s(v344, int32(400))
									if v498 != 0 {
										v500 = base.I32_rem_s(v344, int32(100))
										v504 = base.B2i32(v500 != int32(0))
									} else {
										v504 = int32(1)
									}
									if v26 < v29 {
										v518 = v484&base.B2i32(v343&int32(3) == int32(0))*int32(52) + int32(1686864) + v346<<(uint(int32(2))%32)
									} else {
										v518 = v504&base.B2i32(v344&int32(3) == int32(0))*int32(52) + int32(1686864) + v347<<(uint(int32(2))%32)
									}
									v521 = *(*int32)(unsafe.Add(mBase, uint32(v518-int32(4))))
									v530 = v370
									v534 = v474
									for {
										v543 = v530 - int32(1)
										v544 = v534 + v521
										if v544 < int32(0) {
											v530 = v543
											v534 = v544
											continue
										} else {
											break
										}
										break
									}
									v555 = v543
									v559 = v544
								} else {
									v555 = v370
									v559 = v474
								}
								if v555 < int32(0) {
									v569 = int32(-12)
									if base.Ui32(v555) <= base.Ui32(v569) {
										v572 = v569
									} else {
										v572 = v555
									}
									v574 = base.B2i32(base.Ui32(v555) < base.Ui32(int32(-12)))
									v577 = int32(12)
									v578 = base.I32_div_u_s(v572-(v555+v574), v577)
									v579 = v578 + v574
									v588 = v555 + v579*v577 + v577
									v590 = v375 + (v579 ^ int32(-1))
								} else {
									v588 = v555
									v590 = v375
								}
								if v352 == int32(0) {
									v593 = int32(0)
									v607 = int64(0) - v472
									v608 = v593 - v588
									v609 = v593 - v421
									v610 = v593 - v447
									v611 = v593 - v398
									v612 = v593 - v559
									v613 = v593 - v590
								} else {
									v607 = v472
									v608 = v588
									v609 = v421
									v610 = v447
									v611 = v398
									v612 = v559
									v613 = v590
								}
								v618 = base.I64_extend_i32_s(v608) + base.I64_extend_i32_s(v613)*int64(12)
								if base.Ui64(v618-int64(2147483648)) < base.Ui64(int64(-4294967296)) {
									F_errstart_cold(m, int32(21), int32(0))
									mBase = m.M
									v710 = m.ExcPending
									if v710 != 0 {
										return int32(0)
									} else {
										F_errcode(m, int32(134217858))
										mBase = m.M
										v713 = m.ExcPending
										if v713 != 0 {
											return int32(0)
										} else {
											F_errmsg(m, int32(418557), int32(0))
											mBase = m.M
											v717 = m.ExcPending
											if v717 != 0 {
												return int32(0)
											} else {
												F_errfinish(m, int32(515856), int32(4434), int32(424820))
												mBase = m.M
												v722 = m.ExcPending
												if v722 != 0 {
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
									*(*int32)(unsafe.Add(mBase, uint32(v31)+8)) = v612
									*(*uint32)(unsafe.Add(mBase, uint32(v31)+12)) = uint32(v618)
									v626 = v23 + int32(8)
									v629 = int64(3600000000)
									v630 = int64(0)
									v635 = int64(32)
									v638 = int64(base.Ui64(v607) >> (uint(v635) % 64))
									v641 = int64(4294967295)
									v644 = v607 & v641
									v645 = v629 * v644
									v649 = int64(base.Ui64(v645)>>(uint(v635)%64)) + v629*v638
									v656 = v644*v630 + v649&v641
									*(*int64)(unsafe.Add(mBase, uint32(v626)+8)) = v607*v630 + v607>>(uint(int64(63))%64)*v629 + v630*v638 + int64(base.Ui64(v649)>>(uint(v635)%64)) + int64(base.Ui64(v656)>>(uint(v635)%64))
									*(*int64)(unsafe.Add(mBase, uint32(v626))) = v645&v641 | v656<<(uint(v635)%64)
									v667 = *(*int64)(unsafe.Add(mBase, uint32(v23)+8))
									*(*int64)(unsafe.Add(mBase, uint32(v31))) = v667
									v669 = *(*int64)(unsafe.Add(mBase, uint32(v23)+16))
									if v669 != v667>>(uint(int64(63))%64) {
										F_errstart_cold(m, int32(21), int32(0))
										mBase = m.M
										v710 = m.ExcPending
										if v710 != 0 {
											return int32(0)
										} else {
											F_errcode(m, int32(134217858))
											mBase = m.M
											v713 = m.ExcPending
											if v713 != 0 {
												return int32(0)
											} else {
												F_errmsg(m, int32(418557), int32(0))
												mBase = m.M
												v717 = m.ExcPending
												if v717 != 0 {
													return int32(0)
												} else {
													F_errfinish(m, int32(515856), int32(4434), int32(424820))
													mBase = m.M
													v722 = m.ExcPending
													if v722 != 0 {
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
										v675 = base.I64_extend_i32_s(v610) * int64(60000000)
										v676 = v667 + v675
										*(*int64)(unsafe.Add(mBase, uint32(v31))) = v676
										if base.B2i32(v675 < int64(0))^base.B2i32(v676 < v667) != 0 {
											F_errstart_cold(m, int32(21), int32(0))
											mBase = m.M
											v710 = m.ExcPending
											if v710 != 0 {
												return int32(0)
											} else {
												F_errcode(m, int32(134217858))
												mBase = m.M
												v713 = m.ExcPending
												if v713 != 0 {
													return int32(0)
												} else {
													F_errmsg(m, int32(418557), int32(0))
													mBase = m.M
													v717 = m.ExcPending
													if v717 != 0 {
														return int32(0)
													} else {
														F_errfinish(m, int32(515856), int32(4434), int32(424820))
														mBase = m.M
														v722 = m.ExcPending
														if v722 != 0 {
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
											v684 = base.I64_extend_i32_s(v609) * int64(1000000)
											v685 = v676 + v684
											*(*int64)(unsafe.Add(mBase, uint32(v31))) = v685
											if base.B2i32(v684 < int64(0))^base.B2i32(v685 < v676) != 0 {
												F_errstart_cold(m, int32(21), int32(0))
												mBase = m.M
												v710 = m.ExcPending
												if v710 != 0 {
													return int32(0)
												} else {
													F_errcode(m, int32(134217858))
													mBase = m.M
													v713 = m.ExcPending
													if v713 != 0 {
														return int32(0)
													} else {
														F_errmsg(m, int32(418557), int32(0))
														mBase = m.M
														v717 = m.ExcPending
														if v717 != 0 {
															return int32(0)
														} else {
															F_errfinish(m, int32(515856), int32(4434), int32(424820))
															mBase = m.M
															v722 = m.ExcPending
															if v722 != 0 {
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
												v691 = base.I64_extend_i32_s(v611)
												v692 = v685 + v691
												*(*int64)(unsafe.Add(mBase, uint32(v31))) = v692
												if base.B2i32(v691 < int64(0))^base.B2i32(v692 < v685) != 0 {
													F_errstart_cold(m, int32(21), int32(0))
													mBase = m.M
													v710 = m.ExcPending
													if v710 != 0 {
														return int32(0)
													} else {
														F_errcode(m, int32(134217858))
														mBase = m.M
														v713 = m.ExcPending
														if v713 != 0 {
															return int32(0)
														} else {
															F_errmsg(m, int32(418557), int32(0))
															mBase = m.M
															v717 = m.ExcPending
															if v717 != 0 {
																return int32(0)
															} else {
																F_errfinish(m, int32(515856), int32(4434), int32(424820))
																mBase = m.M
																v722 = m.ExcPending
																if v722 != 0 {
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
													if v618 != int64(2147483647) {
														m.G0 = v23 + int32(112)
														return v31
													} else {
														if v612 != int32(2147483647) {
															m.G0 = v23 + int32(112)
															return v31
														} else {
															if v692 != int64(9223372036854775807) {
																m.G0 = v23 + int32(112)
																return v31
															} else {
																F_errstart_cold(m, int32(21), int32(0))
																mBase = m.M
																v710 = m.ExcPending
																if v710 != 0 {
																	return int32(0)
																} else {
																	F_errcode(m, int32(134217858))
																	mBase = m.M
																	v713 = m.ExcPending
																	if v713 != 0 {
																		return int32(0)
																	} else {
																		F_errmsg(m, int32(418557), int32(0))
																		mBase = m.M
																		v717 = m.ExcPending
																		if v717 != 0 {
																			return int32(0)
																		} else {
																			F_errfinish(m, int32(515856), int32(4434), int32(424820))
																			mBase = m.M
																			v722 = m.ExcPending
																			if v722 != 0 {
																				return int32(0)
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
							}
						}
					} else {
						v85 = v37
						v86 = int32(-2147483648)
						*(*int32)(unsafe.Add(mBase, uint32(v31)+12)) = v86
						*(*int32)(unsafe.Add(mBase, uint32(v31)+8)) = v86
						*(*int64)(unsafe.Add(mBase, uint32(v31))) = v85
						m.G0 = v23 + int32(112)
						return v31
					}
				}
			} else {
				if v29 != int64(-9223372036854775807-1) {
					v85 = v37
					v86 = int32(-2147483648)
					*(*int32)(unsafe.Add(mBase, uint32(v31)+12)) = v86
					*(*int32)(unsafe.Add(mBase, uint32(v31)+8)) = v86
					*(*int64)(unsafe.Add(mBase, uint32(v31))) = v85
					m.G0 = v23 + int32(112)
					return v31
				} else {
					F_errstart_cold(m, int32(21), int32(0))
					mBase = m.M
					v46 = m.ExcPending
					if v46 != 0 {
						return int32(0)
					} else {
						F_errcode(m, int32(134217858))
						mBase = m.M
						v49 = m.ExcPending
						if v49 != 0 {
							return int32(0)
						} else {
							F_errmsg(m, int32(418557), int32(0))
							mBase = m.M
							v53 = m.ExcPending
							if v53 != 0 {
								return int32(0)
							} else {
								F_errfinish(m, int32(515856), int32(4333), int32(424820))
								mBase = m.M
								v58 = m.ExcPending
								if v58 != 0 {
									return int32(0)
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
			if v29 != int64(9223372036854775807) {
				v85 = int64(9223372036854775807)
				v86 = int32(2147483647)
				*(*int32)(unsafe.Add(mBase, uint32(v31)+12)) = v86
				*(*int32)(unsafe.Add(mBase, uint32(v31)+8)) = v86
				*(*int64)(unsafe.Add(mBase, uint32(v31))) = v85
				m.G0 = v23 + int32(112)
				return v31
			} else {
				F_errstart_cold(m, int32(21), int32(0))
				mBase = m.M
				v65 = m.ExcPending
				if v65 != 0 {
					return int32(0)
				} else {
					F_errcode(m, int32(134217858))
					mBase = m.M
					v68 = m.ExcPending
					if v68 != 0 {
						return int32(0)
					} else {
						F_errmsg(m, int32(418557), int32(0))
						mBase = m.M
						v72 = m.ExcPending
						if v72 != 0 {
							return int32(0)
						} else {
							F_errfinish(m, int32(515856), int32(4342), int32(424820))
							mBase = m.M
							v77 = m.ExcPending
							if v77 != 0 {
								return int32(0)
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
func F_timestamp_bin(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v8 int32
	_ = v8
	var v10 int32
	_ = v10
	var v12 int32
	_ = v12
	var v13 int64
	_ = v13
	var v18 int32
	_ = v18
	var v19 int64
	_ = v19
	var v24 int32
	_ = v24
	var v25 int32
	_ = v25
	var v30 int32
	_ = v30
	var v33 int64
	_ = v33
	var v36 int32
	_ = v36
	var v39 int64
	_ = v39
	var v47 int32
	_ = v47
	var v50 int32
	_ = v50
	var v54 int32
	_ = v54
	var v59 int32
	_ = v59
	var v60 int64
	_ = v60
	var v69 int64
	_ = v69
	var v70 int64
	_ = v70
	var v72 int64
	_ = v72
	var v75 int64
	_ = v75
	var v76 int64
	_ = v76
	var v78 int64
	_ = v78
	var v79 int64
	_ = v79
	var v83 int64
	_ = v83
	var v90 int64
	_ = v90
	var v101 int64
	_ = v101
	var v102 int64
	_ = v102
	var v106 int64
	_ = v106
	var v109 int64
	_ = v109
	var v114 int64
	_ = v114
	var v119 int64
	_ = v119
	var v121 int64
	_ = v121
	var v124 int64
	_ = v124
	var v134 int64
	_ = v134
	var v139 int32
	_ = v139
	var v140 int32
	_ = v140
	var v148 int32
	_ = v148
	var v151 int32
	_ = v151
	var v155 int32
	_ = v155
	var v160 int32
	_ = v160
	var v165 int32
	_ = v165
	var v168 int32
	_ = v168
	var v172 int32
	_ = v172
	var v177 int32
	_ = v177
	var v181 int32
	_ = v181
	var v184 int32
	_ = v184
	var v188 int32
	_ = v188
	var v193 int32
	_ = v193
	var v197 int32
	_ = v197
	var v200 int32
	_ = v200
	var v204 int32
	_ = v204
	var v209 int32
	_ = v209
	var v213 int32
	_ = v213
	var v216 int32
	_ = v216
	var v220 int32
	_ = v220
	var v225 int32
	_ = v225
	var v229 int32
	_ = v229
	var v232 int32
	_ = v232
	var v236 int32
	_ = v236
	var v241 int32
	_ = v241
	v8 = m.G0
	v10 = v8 - int32(16)
	m.G0 = v10
	v12 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
	v13 = *(*int64)(unsafe.Add(mBase, uint32(v12)))
	if base.Ui64(v13-int64(9223372036854775807)) < base.Ui64(int64(2)) {
		v134 = v13
		v139 = F_Int64GetDatum(m, v134)
		mBase = m.M
		v140 = m.ExcPending
		if v140 != 0 {
			return int32(0)
		} else {
			m.G0 = v10 + int32(16)
			return v139
		}
	} else {
		v18 = *(*int32)(unsafe.Add(mBase, uint32(l0)+36))
		v19 = *(*int64)(unsafe.Add(mBase, uint32(v18)))
		if base.Ui64(v19-int64(9223372036854775807)) <= base.Ui64(int64(1)) {
			F_errstart_cold(m, int32(21), int32(0))
			mBase = m.M
			v148 = m.ExcPending
			if v148 != 0 {
				return int32(0)
			} else {
				F_errcode(m, int32(134217858))
				mBase = m.M
				v151 = m.ExcPending
				if v151 != 0 {
					return int32(0)
				} else {
					F_errmsg(m, int32(418537), int32(0))
					mBase = m.M
					v155 = m.ExcPending
					if v155 != 0 {
						return int32(0)
					} else {
						F_errfinish(m, int32(515856), int32(4623), int32(288541))
						mBase = m.M
						v160 = m.ExcPending
						if v160 != 0 {
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
			v24 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
			v25 = *(*int32)(unsafe.Add(mBase, uint32(v24)+12))
			if v25 != 0 {
				if v25 != int32(2147483647) {
					if v25 != int32(-2147483648) {
						F_errstart_cold(m, int32(21), int32(0))
						mBase = m.M
						v47 = m.ExcPending
						if v47 != 0 {
							return int32(0)
						} else {
							F_errcode(m, int32(1088))
							mBase = m.M
							v50 = m.ExcPending
							if v50 != 0 {
								return int32(0)
							} else {
								F_errmsg(m, int32(143604), int32(0))
								mBase = m.M
								v54 = m.ExcPending
								if v54 != 0 {
									return int32(0)
								} else {
									F_errfinish(m, int32(515856), int32(4633), int32(288541))
									mBase = m.M
									v59 = m.ExcPending
									if v59 != 0 {
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
						v30 = *(*int32)(unsafe.Add(mBase, uint32(v24)+8))
						if v30 != int32(-2147483648) {
							F_errstart_cold(m, int32(21), int32(0))
							mBase = m.M
							v47 = m.ExcPending
							if v47 != 0 {
								return int32(0)
							} else {
								F_errcode(m, int32(1088))
								mBase = m.M
								v50 = m.ExcPending
								if v50 != 0 {
									return int32(0)
								} else {
									F_errmsg(m, int32(143604), int32(0))
									mBase = m.M
									v54 = m.ExcPending
									if v54 != 0 {
										return int32(0)
									} else {
										F_errfinish(m, int32(515856), int32(4633), int32(288541))
										mBase = m.M
										v59 = m.ExcPending
										if v59 != 0 {
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
							v33 = *(*int64)(unsafe.Add(mBase, uint32(v24)))
							if v33 != int64(-9223372036854775807-1) {
								F_errstart_cold(m, int32(21), int32(0))
								mBase = m.M
								v47 = m.ExcPending
								if v47 != 0 {
									return int32(0)
								} else {
									F_errcode(m, int32(1088))
									mBase = m.M
									v50 = m.ExcPending
									if v50 != 0 {
										return int32(0)
									} else {
										F_errmsg(m, int32(143604), int32(0))
										mBase = m.M
										v54 = m.ExcPending
										if v54 != 0 {
											return int32(0)
										} else {
											F_errfinish(m, int32(515856), int32(4633), int32(288541))
											mBase = m.M
											v59 = m.ExcPending
											if v59 != 0 {
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
								F_errstart_cold(m, int32(21), int32(0))
								mBase = m.M
								v229 = m.ExcPending
								if v229 != 0 {
									return int32(0)
								} else {
									F_errcode(m, int32(134217858))
									mBase = m.M
									v232 = m.ExcPending
									if v232 != 0 {
										return int32(0)
									} else {
										F_errmsg(m, int32(161560), int32(0))
										mBase = m.M
										v236 = m.ExcPending
										if v236 != 0 {
											return int32(0)
										} else {
											F_errfinish(m, int32(515856), int32(4628), int32(288541))
											mBase = m.M
											v241 = m.ExcPending
											if v241 != 0 {
												return int32(0)
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
					v36 = *(*int32)(unsafe.Add(mBase, uint32(v24)+8))
					if v36 != int32(2147483647) {
						F_errstart_cold(m, int32(21), int32(0))
						mBase = m.M
						v47 = m.ExcPending
						if v47 != 0 {
							return int32(0)
						} else {
							F_errcode(m, int32(1088))
							mBase = m.M
							v50 = m.ExcPending
							if v50 != 0 {
								return int32(0)
							} else {
								F_errmsg(m, int32(143604), int32(0))
								mBase = m.M
								v54 = m.ExcPending
								if v54 != 0 {
									return int32(0)
								} else {
									F_errfinish(m, int32(515856), int32(4633), int32(288541))
									mBase = m.M
									v59 = m.ExcPending
									if v59 != 0 {
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
						v39 = *(*int64)(unsafe.Add(mBase, uint32(v24)))
						if v39 == int64(9223372036854775807) {
							F_errstart_cold(m, int32(21), int32(0))
							mBase = m.M
							v229 = m.ExcPending
							if v229 != 0 {
								return int32(0)
							} else {
								F_errcode(m, int32(134217858))
								mBase = m.M
								v232 = m.ExcPending
								if v232 != 0 {
									return int32(0)
								} else {
									F_errmsg(m, int32(161560), int32(0))
									mBase = m.M
									v236 = m.ExcPending
									if v236 != 0 {
										return int32(0)
									} else {
										F_errfinish(m, int32(515856), int32(4628), int32(288541))
										mBase = m.M
										v241 = m.ExcPending
										if v241 != 0 {
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
							F_errstart_cold(m, int32(21), int32(0))
							mBase = m.M
							v47 = m.ExcPending
							if v47 != 0 {
								return int32(0)
							} else {
								F_errcode(m, int32(1088))
								mBase = m.M
								v50 = m.ExcPending
								if v50 != 0 {
									return int32(0)
								} else {
									F_errmsg(m, int32(143604), int32(0))
									mBase = m.M
									v54 = m.ExcPending
									if v54 != 0 {
										return int32(0)
									} else {
										F_errfinish(m, int32(515856), int32(4633), int32(288541))
										mBase = m.M
										v59 = m.ExcPending
										if v59 != 0 {
											return int32(0)
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
				v60 = int64(*(*int32)(unsafe.Add(mBase, uint32(v24)+8)))
				v69 = int64(32)
				v70 = int64(20)
				v72 = int64(base.Ui64(v60) >> (uint(v69) % 64))
				v75 = int64(4294967295)
				v76 = int64(500654080)
				v78 = v60 & v75
				v79 = v76 * v78
				v83 = int64(base.Ui64(v79)>>(uint(v69)%64)) + v76*v72
				v90 = v78*v70 + v83&v75
				*(*int64)(unsafe.Add(mBase, uint32(v10)+8)) = v60*int64(0) + v60>>(uint(int64(63))%64)*int64(86400000000) + v70*v72 + int64(base.Ui64(v83)>>(uint(v69)%64)) + int64(base.Ui64(v90)>>(uint(v69)%64))
				*(*int64)(unsafe.Add(mBase, uint32(v10))) = v79&v75 | v90<<(uint(v69)%64)
				v101 = *(*int64)(unsafe.Add(mBase, uint32(v10)+8))
				v102 = *(*int64)(unsafe.Add(mBase, uint32(v10)))
				if v101 != v102>>(uint(int64(63))%64) {
					F_errstart_cold(m, int32(21), int32(0))
					mBase = m.M
					v165 = m.ExcPending
					if v165 != 0 {
						return int32(0)
					} else {
						F_errcode(m, int32(134217858))
						mBase = m.M
						v168 = m.ExcPending
						if v168 != 0 {
							return int32(0)
						} else {
							F_errmsg(m, int32(418557), int32(0))
							mBase = m.M
							v172 = m.ExcPending
							if v172 != 0 {
								return int32(0)
							} else {
								F_errfinish(m, int32(515856), int32(4639), int32(288541))
								mBase = m.M
								v177 = m.ExcPending
								if v177 != 0 {
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
					v106 = *(*int64)(unsafe.Add(mBase, uint32(v24)))
					v109 = v106 + v102
					if base.B2i32(v106 < int64(0)) != base.B2i32(v109 < v102) {
						F_errstart_cold(m, int32(21), int32(0))
						mBase = m.M
						v165 = m.ExcPending
						if v165 != 0 {
							return int32(0)
						} else {
							F_errcode(m, int32(134217858))
							mBase = m.M
							v168 = m.ExcPending
							if v168 != 0 {
								return int32(0)
							} else {
								F_errmsg(m, int32(418557), int32(0))
								mBase = m.M
								v172 = m.ExcPending
								if v172 != 0 {
									return int32(0)
								} else {
									F_errfinish(m, int32(515856), int32(4639), int32(288541))
									mBase = m.M
									v177 = m.ExcPending
									if v177 != 0 {
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
						if v109 <= int64(0) {
							F_errstart_cold(m, int32(21), int32(0))
							mBase = m.M
							v181 = m.ExcPending
							if v181 != 0 {
								return int32(0)
							} else {
								F_errcode(m, int32(134217858))
								mBase = m.M
								v184 = m.ExcPending
								if v184 != 0 {
									return int32(0)
								} else {
									F_errmsg(m, int32(250132), int32(0))
									mBase = m.M
									v188 = m.ExcPending
									if v188 != 0 {
										return int32(0)
									} else {
										F_errfinish(m, int32(515856), int32(4644), int32(288541))
										mBase = m.M
										v193 = m.ExcPending
										if v193 != 0 {
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
							v114 = v13 - v19
							if base.B2i32(v114 < v13) != base.B2i32(int64(0) < v19) {
								F_errstart_cold(m, int32(21), int32(0))
								mBase = m.M
								v197 = m.ExcPending
								if v197 != 0 {
									return int32(0)
								} else {
									F_errcode(m, int32(134217858))
									mBase = m.M
									v200 = m.ExcPending
									if v200 != 0 {
										return int32(0)
									} else {
										F_errmsg(m, int32(418557), int32(0))
										mBase = m.M
										v204 = m.ExcPending
										if v204 != 0 {
											return int32(0)
										} else {
											F_errfinish(m, int32(515856), int32(4649), int32(288541))
											mBase = m.M
											v209 = m.ExcPending
											if v209 != 0 {
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
								v119 = base.I64_rem_s(v114, v109)
								v121 = v114 - v119 + v19
								if int64(0) <= v119 {
									v134 = v121
									v139 = F_Int64GetDatum(m, v134)
									mBase = m.M
									v140 = m.ExcPending
									if v140 != 0 {
										return int32(0)
									} else {
										m.G0 = v10 + int32(16)
										return v139
									}
								} else {
									v124 = v121 - v109
									if base.B2i32(v124 < v121)^base.B2i32(int64(0) < v109) != 0 {
										F_errstart_cold(m, int32(21), int32(0))
										mBase = m.M
										v213 = m.ExcPending
										if v213 != 0 {
											return int32(0)
										} else {
											F_errcode(m, int32(134217858))
											mBase = m.M
											v216 = m.ExcPending
											if v216 != 0 {
												return int32(0)
											} else {
												F_errmsg(m, int32(418494), int32(0))
												mBase = m.M
												v220 = m.ExcPending
												if v220 != 0 {
													return int32(0)
												} else {
													F_errfinish(m, int32(515856), int32(4667), int32(288541))
													mBase = m.M
													v225 = m.ExcPending
													if v225 != 0 {
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
										if base.Ui64(v124-int64(9223371331200000000)) <= base.Ui64(int64(9011559254509551615)) {
											F_errstart_cold(m, int32(21), int32(0))
											mBase = m.M
											v213 = m.ExcPending
											if v213 != 0 {
												return int32(0)
											} else {
												F_errcode(m, int32(134217858))
												mBase = m.M
												v216 = m.ExcPending
												if v216 != 0 {
													return int32(0)
												} else {
													F_errmsg(m, int32(418494), int32(0))
													mBase = m.M
													v220 = m.ExcPending
													if v220 != 0 {
														return int32(0)
													} else {
														F_errfinish(m, int32(515856), int32(4667), int32(288541))
														mBase = m.M
														v225 = m.ExcPending
														if v225 != 0 {
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
											v134 = v124
											v139 = F_Int64GetDatum(m, v134)
											mBase = m.M
											v140 = m.ExcPending
											if v140 != 0 {
												return int32(0)
											} else {
												m.G0 = v10 + int32(16)
												return v139
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
func F_timestamp_lt_date(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v3 int32
	_ = v3
	var v4 int64
	_ = v4
	var v6 int32
	_ = v6
	var v17 int64
	_ = v17
	var v25 int32
	_ = v25
	var v26 int32
	_ = v26
	v3 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v4 = *(*int64)(unsafe.Add(mBase, uint32(v3)))
	v6 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
	if v6 == int32(-2147483648) {
		v17 = int64(-9223372036854775807 - 1)
		v26 = base.B2i32(v4 < v17) - base.B2i32(v17 < v4)
	} else {
		if v6 == int32(2147483647) {
			v17 = int64(9223372036854775807)
			v26 = base.B2i32(v4 < v17) - base.B2i32(v17 < v4)
		} else {
			if int32(106751982) < v6 {
				if v4 == int64(9223372036854775807) {
					v25 = int32(-1)
				} else {
					v25 = int32(1)
				}
				v26 = v25
			} else {
				v17 = base.I64_extend_i32_s(v6) * int64(86400000000)
				v26 = base.B2i32(v4 < v17) - base.B2i32(v17 < v4)
			}
		}
	}
	return base.B2i32(int32(0) < v26)
}
func F_timestamp_part(m *base.Module, l0 int32) int32 {
	var v3 int32
	_ = v3
	var v6 int32
	_ = v6
	v3 = F_timestamp_part_common(m, l0, int32(0))
	v6 = m.ExcPending
	if v6 != 0 {
		return int32(0)
	} else {
		return v3
	}
}
