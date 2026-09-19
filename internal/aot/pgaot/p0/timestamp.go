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
	var v369 int32
	_ = v369
	var v370 int64
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
	var v483 int32
	_ = v483
	var v486 int32
	_ = v486
	var v491 int32
	_ = v491
	var v492 int32
	_ = v492
	var v503 int32
	_ = v503
	var v506 int32
	_ = v506
	var v511 int32
	_ = v511
	var v512 int32
	_ = v512
	var v518 int32
	_ = v518
	var v521 int32
	_ = v521
	var v531 int32
	_ = v531
	var v533 int32
	_ = v533
	var v543 int32
	_ = v543
	var v544 int32
	_ = v544
	var v556 int32
	_ = v556
	var v558 int32
	_ = v558
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
	var v607 int32
	_ = v607
	var v608 int64
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
	var v627 int64
	_ = v627
	var v628 int64
	_ = v628
	var v633 int64
	_ = v633
	var v636 int64
	_ = v636
	var v639 int64
	_ = v639
	var v642 int64
	_ = v642
	var v643 int64
	_ = v643
	var v647 int64
	_ = v647
	var v654 int64
	_ = v654
	var v665 int64
	_ = v665
	var v667 int64
	_ = v667
	var v673 int64
	_ = v673
	var v674 int64
	_ = v674
	var v682 int64
	_ = v682
	var v683 int64
	_ = v683
	var v689 int64
	_ = v689
	var v690 int64
	_ = v690
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
									F_errmsg(m, int32(_a_F_timestamp_age_0), int32(0))
									mBase = m.M
									v737 = m.ExcPending
									if v737 != 0 {
										return int32(0)
									} else {
										F_errfinish(m, int32(_a_F_timestamp_age_1), int32(_a_F_timestamp_age_2), int32(_a_F_timestamp_age_3))
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
							v118 = v106 + int32(_a_F_timestamp_age_4)
							v119 = int32(_a_F_timestamp_age_5)
							v120 = base.I32_div_u_s(v118, v119)
							v121 = int32(3)
							v127 = int32(2)
							v132 = base.I32_div_u_s((v120*int32(1073595727)+v118)<<(uint(v127)%32)|v121, v119)
							v135 = v106 + int32(_a_F_timestamp_age_6) + v120*v121 + v132 + int32(_a_F_timestamp_age_7)
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
							*(*int32)(unsafe.Add(mBase, uint32(v23+int32(88)))) = v155 + v137<<(uint(int32(2))%32) - int32(_a_F_timestamp_age_8)
							v163 = v153 + int32(123)
							v167 = int32(base.Ui32(v163*int32(2141)) >> (uint(int32(16)) % 32))
							*(*int32)(unsafe.Add(mBase, uint32(v23+int32(80)))) = v163 - int32(base.Ui32(v167*int32(_a_F_timestamp_age_9))>>(uint(int32(8))%32))
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
										F_errmsg(m, int32(_a_F_timestamp_age_0), int32(0))
										mBase = m.M
										v737 = m.ExcPending
										if v737 != 0 {
											return int32(0)
										} else {
											F_errfinish(m, int32(_a_F_timestamp_age_1), int32(_a_F_timestamp_age_2), int32(_a_F_timestamp_age_3))
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
								v235 = v223 + int32(_a_F_timestamp_age_4)
								v236 = int32(_a_F_timestamp_age_5)
								v237 = base.I32_div_u_s(v235, v236)
								v238 = int32(3)
								v244 = int32(2)
								v249 = base.I32_div_u_s((v237*int32(1073595727)+v235)<<(uint(v244)%32)|v238, v236)
								v252 = v223 + int32(_a_F_timestamp_age_6) + v237*v238 + v249 + int32(_a_F_timestamp_age_7)
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
								*(*int32)(unsafe.Add(mBase, uint32(v23+int32(44)))) = v272 + v254<<(uint(int32(2))%32) - int32(_a_F_timestamp_age_8)
								v280 = v270 + int32(123)
								v284 = int32(base.Ui32(v280*int32(2141)) >> (uint(int32(16)) % 32))
								*(*int32)(unsafe.Add(mBase, uint32(v23+int32(36)))) = v280 - int32(base.Ui32(v284*int32(_a_F_timestamp_age_9))>>(uint(int32(8))%32))
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
									v369 = v355 - v333
									v370 = int64(0) - v342
									v371 = v355 - v337
									v372 = v355 - v348
									v373 = v355 - v329
									v374 = v355 - v351
									v375 = v355 - v345
								} else {
									v369 = v333
									v370 = v342
									v371 = v337
									v372 = v348
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
									v386 = int32(_a_F_timestamp_age_10)
									v387 = base.I32_div_u_s(v381-(v373+v383), v386)
									v388 = v387 + v383
									v397 = v369 + (v388 ^ int32(-1))
									v398 = v373 + v388*v386 + v386
								} else {
									v397 = v369
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
									v422 = v371 + (v412 ^ int32(-1))
								} else {
									v421 = v397
									v422 = v371
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
									v446 = v370 + base.I64_extend_i32_s(v436^int32(-1))
									v447 = v422 + v436*v434 + v434
								} else {
									v446 = v370
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
									v483 = int32(0)
									v486 = base.I32_rem_s(v343, int32(100))
									v491 = base.I32_rem_s(v343, int32(400))
									if v491 != 0 {
										v492 = base.B2i32(v486 != v483)
									} else {
										v492 = int32(1)
									}
									v503 = int32(0)
									v506 = base.I32_rem_s(v344, int32(100))
									v511 = base.I32_rem_s(v344, int32(400))
									if v511 != 0 {
										v512 = base.B2i32(v506 != v503)
									} else {
										v512 = int32(1)
									}
									if v26 < v29 {
										v518 = v346<<(uint(int32(2))%32) + int32(_a_F_timestamp_age_11) + base.B2i32(v343&int32(3) == v483)&v492*int32(52)
									} else {
										v518 = v347<<(uint(int32(2))%32) + int32(_a_F_timestamp_age_11) + base.B2i32(v344&int32(3) == v503)&v512*int32(52)
									}
									v521 = *(*int32)(unsafe.Add(mBase, uint32(v518-int32(4))))
									v531 = v372
									v533 = v474
									for {
										v543 = v531 - int32(1)
										v544 = v533 + v521
										if v544 < int32(0) {
											v531 = v543
											v533 = v544
											continue
										} else {
											break
										}
										break
									}
									v556 = v543
									v558 = v544
								} else {
									v556 = v372
									v558 = v474
								}
								if v556 < int32(0) {
									v569 = int32(-12)
									if base.Ui32(v556) <= base.Ui32(v569) {
										v572 = v569
									} else {
										v572 = v556
									}
									v574 = base.B2i32(base.Ui32(v556) < base.Ui32(int32(-12)))
									v577 = int32(12)
									v578 = base.I32_div_u_s(v572-(v556+v574), v577)
									v579 = v578 + v574
									v588 = v556 + v579*v577 + v577
									v590 = v375 + (v579 ^ int32(-1))
								} else {
									v588 = v556
									v590 = v375
								}
								if v352 == int32(0) {
									v593 = int32(0)
									v607 = v593 - v421
									v608 = int64(0) - v472
									v609 = v593 - v447
									v610 = v593 - v588
									v611 = v593 - v398
									v612 = v593 - v558
									v613 = v593 - v590
								} else {
									v607 = v421
									v608 = v472
									v609 = v447
									v610 = v588
									v611 = v398
									v612 = v558
									v613 = v590
								}
								v618 = base.I64_extend_i32_s(v610) + base.I64_extend_i32_s(v613)*int64(12)
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
											F_errmsg(m, int32(_a_F_timestamp_age_12), int32(0))
											mBase = m.M
											v717 = m.ExcPending
											if v717 != 0 {
												return int32(0)
											} else {
												F_errfinish(m, int32(_a_F_timestamp_age_1), int32(_a_F_timestamp_age_13), int32(_a_F_timestamp_age_3))
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
									v627 = int64(3600000000)
									v628 = int64(0)
									v633 = int64(32)
									v636 = int64(base.Ui64(v608) >> (uint(v633) % 64))
									v639 = int64(4294967295)
									v642 = v608 & v639
									v643 = v627 * v642
									v647 = int64(base.Ui64(v643)>>(uint(v633)%64)) + v627*v636
									v654 = v642*v628 + v647&v639
									*(*int64)(unsafe.Add(mBase, uint32(v23)+8)) = v608*v628 + v608>>(uint(int64(63))%64)*v627 + v628*v636 + int64(base.Ui64(v647)>>(uint(v633)%64)) + int64(base.Ui64(v654)>>(uint(v633)%64))
									*(*int64)(unsafe.Add(mBase, uint32(v23))) = v643&v639 | v654<<(uint(v633)%64)
									v665 = *(*int64)(unsafe.Add(mBase, uint32(v23)))
									*(*int64)(unsafe.Add(mBase, uint32(v31))) = v665
									v667 = *(*int64)(unsafe.Add(mBase, uint32(v23)+8))
									if v667 != v665>>(uint(int64(63))%64) {
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
												F_errmsg(m, int32(_a_F_timestamp_age_12), int32(0))
												mBase = m.M
												v717 = m.ExcPending
												if v717 != 0 {
													return int32(0)
												} else {
													F_errfinish(m, int32(_a_F_timestamp_age_1), int32(_a_F_timestamp_age_13), int32(_a_F_timestamp_age_3))
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
										v673 = base.I64_extend_i32_s(v609) * int64(60000000)
										v674 = v665 + v673
										*(*int64)(unsafe.Add(mBase, uint32(v31))) = v674
										if base.B2i32(v673 < int64(0))^base.B2i32(v674 < v665) != 0 {
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
													F_errmsg(m, int32(_a_F_timestamp_age_12), int32(0))
													mBase = m.M
													v717 = m.ExcPending
													if v717 != 0 {
														return int32(0)
													} else {
														F_errfinish(m, int32(_a_F_timestamp_age_1), int32(_a_F_timestamp_age_13), int32(_a_F_timestamp_age_3))
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
											v682 = base.I64_extend_i32_s(v607) * int64(1000000)
											v683 = v674 + v682
											*(*int64)(unsafe.Add(mBase, uint32(v31))) = v683
											if base.B2i32(v682 < int64(0))^base.B2i32(v683 < v674) != 0 {
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
														F_errmsg(m, int32(_a_F_timestamp_age_12), int32(0))
														mBase = m.M
														v717 = m.ExcPending
														if v717 != 0 {
															return int32(0)
														} else {
															F_errfinish(m, int32(_a_F_timestamp_age_1), int32(_a_F_timestamp_age_13), int32(_a_F_timestamp_age_3))
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
												v689 = base.I64_extend_i32_s(v611)
												v690 = v683 + v689
												*(*int64)(unsafe.Add(mBase, uint32(v31))) = v690
												if base.B2i32(v689 < int64(0))^base.B2i32(v690 < v683) != 0 {
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
															F_errmsg(m, int32(_a_F_timestamp_age_12), int32(0))
															mBase = m.M
															v717 = m.ExcPending
															if v717 != 0 {
																return int32(0)
															} else {
																F_errfinish(m, int32(_a_F_timestamp_age_1), int32(_a_F_timestamp_age_13), int32(_a_F_timestamp_age_3))
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
													if base.B2i32(v612 != int32(2147483647))|base.B2i32(v618 != int64(2147483647))|base.B2i32(v690 != int64(9223372036854775807)) != 0 {
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
																F_errmsg(m, int32(_a_F_timestamp_age_12), int32(0))
																mBase = m.M
																v717 = m.ExcPending
																if v717 != 0 {
																	return int32(0)
																} else {
																	F_errfinish(m, int32(_a_F_timestamp_age_1), int32(_a_F_timestamp_age_13), int32(_a_F_timestamp_age_3))
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
							F_errmsg(m, int32(_a_F_timestamp_age_12), int32(0))
							mBase = m.M
							v53 = m.ExcPending
							if v53 != 0 {
								return int32(0)
							} else {
								F_errfinish(m, int32(_a_F_timestamp_age_1), int32(_a_F_timestamp_age_14), int32(_a_F_timestamp_age_3))
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
						F_errmsg(m, int32(_a_F_timestamp_age_12), int32(0))
						mBase = m.M
						v72 = m.ExcPending
						if v72 != 0 {
							return int32(0)
						} else {
							F_errfinish(m, int32(_a_F_timestamp_age_1), int32(_a_F_timestamp_age_15), int32(_a_F_timestamp_age_3))
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
	var v10 int32
	_ = v10
	var v13 int32
	_ = v13
	v10 = Fn14000(m, l0, int32(_a_F_timestamp_bin_0), int32(_a_F_timestamp_bin_1), int32(_a_F_timestamp_bin_2), int32(_a_F_timestamp_bin_3), int32(_a_F_timestamp_bin_4), int32(_a_F_timestamp_bin_5), int32(_a_F_timestamp_bin_6), int32(_a_F_timestamp_bin_7))
	v13 = m.ExcPending
	if v13 != 0 {
		return int32(0)
	} else {
		return v10
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
