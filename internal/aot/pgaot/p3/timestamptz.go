package p3

import (
	base "github.com/shibukawa/pgmem/internal/aot/pgaot/base"
	"unsafe"
)

func F_timestamptz_date(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v4 int32
	_ = v4
	var v6 int32
	_ = v6
	var v9 int32
	_ = v9
	var v10 int64
	_ = v10
	var v22 int32
	_ = v22
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
	var v35 int32
	_ = v35
	var v36 int32
	_ = v36
	var v37 int32
	_ = v37
	var v42 int32
	_ = v42
	var v45 int32
	_ = v45
	var v48 int32
	_ = v48
	var v52 int32
	_ = v52
	var v57 int32
	_ = v57
	var v63 int32
	_ = v63
	var v71 int32
	_ = v71
	var v74 int32
	_ = v74
	var v78 int32
	_ = v78
	var v83 int32
	_ = v83
	v4 = m.G0
	v6 = v4 + int32(-64)
	m.G0 = v6
	v9 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v10 = *(*int64)(unsafe.Add(mBase, uint32(v9)))
	if v10 == int64(-9223372036854775807-1) {
		v63 = int32(-2147483648)
		m.G0 = v6 - int32(-64)
		return v63
	} else {
		if v10 == int64(9223372036854775807) {
			v63 = int32(2147483647)
			m.G0 = v6 - int32(-64)
			return v63
		} else {
			v22 = int32(0)
			v24 = F_timestamp2tm(m, v10, v4+int32(-52), v4+int32(-44), v4+int32(-48), v22, v22)
			mBase = m.M
			v27 = m.ExcPending
			if v27 != 0 {
				return int32(0)
			} else {
				if v24 != 0 {
					F_errstart_cold(m, int32(21), int32(0))
					mBase = m.M
					v71 = m.ExcPending
					if v71 != 0 {
						return int32(0)
					} else {
						F_errcode(m, int32(134217858))
						mBase = m.M
						v74 = m.ExcPending
						if v74 != 0 {
							return int32(0)
						} else {
							F_errmsg(m, int32(418494), int32(0))
							mBase = m.M
							v78 = m.ExcPending
							if v78 != 0 {
								return int32(0)
							} else {
								F_errfinish(m, int32(519090), int32(1425), int32(371823))
								mBase = m.M
								v83 = m.ExcPending
								if v83 != 0 {
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
					v28 = *(*int32)(unsafe.Add(mBase, uint32(v6)+40))
					v29 = *(*int32)(unsafe.Add(mBase, uint32(v6)+36))
					v30 = *(*int32)(unsafe.Add(mBase, uint32(v6)+32))
					v35 = base.B2i32(int32(2) < v29)
					if int32(2) < v29 {
						v36 = int32(4800)
					} else {
						v36 = int32(4799)
					}
					v37 = v36 + v28
					v42 = base.I32_div_s(v37, int32(4))
					v45 = base.I32_div_s(v37, int32(-100))
					v48 = base.I32_div_s(v37, int32(400))
					if int32(2) < v29 {
						v52 = int32(1)
					} else {
						v52 = int32(13)
					}
					v57 = base.I32_div_s((v52+v29)*int32(7834), int32(256))
					v63 = v30 + v37*int32(365) + v42 + v45 + v48 + v57 - int32(32167) - int32(2451545)
					m.G0 = v6 - int32(-64)
					return v63
				}
			}
		}
	}
}
func F_timestamptz_gt_date(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v2 int32
	_ = v2
	var v3 int32
	_ = v3
	var v4 int64
	_ = v4
	var v7 int32
	_ = v7
	var v9 int32
	_ = v9
	var v14 int32
	_ = v14
	var v36 int32
	_ = v36
	var v37 int32
	_ = v37
	var v44 int64
	_ = v44
	var v56 int32
	_ = v56
	var v61 int32
	_ = v61
	var v62 int64
	_ = v62
	var v63 int32
	_ = v63
	var v65 int32
	_ = v65
	v2 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
	v3 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v4 = *(*int64)(unsafe.Add(mBase, uint32(v3)))
	v7 = m.G0
	v9 = v7 - int32(48)
	m.G0 = v9
	if v2 == int32(-2147483648) {
		v14 = F_timestamp_cmp_internal(m, int64(-9223372036854775807-1), v4)
		mBase = m.M
		v65 = v14
	} else {
		if v2 == int32(2147483647) {
			v62 = int64(9223372036854775807)
			v63 = F_timestamp_cmp_internal(m, v62, v4)
			mBase = m.M
			v65 = v63
		} else {
			if v2 <= int32(106751982) {
				F_j2date(m, v2+int32(2451545), v9+int32(24), v9+int32(20), v9+int32(16))
				mBase = m.M
				*(*int64)(unsafe.Add(mBase, uint32(v9)+4)) = int64(0)
				*(*int32)(unsafe.Add(mBase, uint32(v9)+12)) = int32(0)
				v36 = *(*int32)(unsafe.Add(mBase, _consts[526]))
				v37 = F_DetermineTimeZoneOffset(m, v9+int32(4), v36)
				mBase = m.M
				v44 = base.I64_extend_i32_s(v37)*int64(1000000) + base.I64_extend_i32_s(v2)*int64(86400000000)
				if base.Ui64(v44+int64(211813488000000000)) < base.Ui64(int64(-9011559254509551616)) {
					v62 = v44
					v63 = F_timestamp_cmp_internal(m, v62, v4)
					mBase = m.M
					v65 = v63
				} else {
					if v44 < int64(-211813488000000000) {
						if v4 == int64(-9223372036854775807-1) {
							v61 = int32(1)
						} else {
							v61 = int32(-1)
						}
						v65 = v61
					} else {
						if v4 == int64(9223372036854775807) {
							v56 = int32(-1)
						} else {
							v56 = int32(1)
						}
						v65 = v56
					}
				}
			} else {
				if v4 == int64(9223372036854775807) {
					v56 = int32(-1)
				} else {
					v56 = int32(1)
				}
				v65 = v56
			}
		}
	}
	m.G0 = v9 + int32(48)
	return int32(base.Ui32(v65) >> (uint(int32(31)) % 32))
}
func F_timestamptz_gt_timestamp(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v5 int32
	_ = v5
	var v7 int32
	_ = v7
	var v9 int32
	_ = v9
	var v10 int64
	_ = v10
	var v11 int32
	_ = v11
	var v12 int64
	_ = v12
	var v15 int64
	_ = v15
	var v18 int32
	_ = v18
	var v21 int32
	_ = v21
	var v29 int32
	_ = v29
	v5 = m.G0
	v7 = v5 - int32(16)
	m.G0 = v7
	v9 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v10 = *(*int64)(unsafe.Add(mBase, uint32(v9)))
	v11 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
	v12 = *(*int64)(unsafe.Add(mBase, uint32(v11)))
	v15 = F_timestamp2timestamptz_opt_overflow(m, v12, v7+int32(12))
	mBase = m.M
	v18 = m.ExcPending
	if v18 != 0 {
		return int32(0)
	} else {
		v21 = *(*int32)(unsafe.Add(mBase, uint32(v7)+12))
		if int32(0) < v21 {
			v29 = base.B2i32(v10 == int64(9223372036854775807))
		} else {
			if v21 < int32(0) {
				v29 = base.B2i32(v10 != int64(-9223372036854775807-1))
			} else {
				v29 = base.B2i32(v15 < v10)
			}
		}
		m.G0 = v7 + int32(16)
		return v29
	}
}
func F_timestamptz_le_timestamp(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v5 int32
	_ = v5
	var v7 int32
	_ = v7
	var v9 int32
	_ = v9
	var v10 int64
	_ = v10
	var v11 int32
	_ = v11
	var v12 int64
	_ = v12
	var v15 int64
	_ = v15
	var v18 int32
	_ = v18
	var v21 int32
	_ = v21
	var v29 int32
	_ = v29
	v5 = m.G0
	v7 = v5 - int32(16)
	m.G0 = v7
	v9 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v10 = *(*int64)(unsafe.Add(mBase, uint32(v9)))
	v11 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
	v12 = *(*int64)(unsafe.Add(mBase, uint32(v11)))
	v15 = F_timestamp2timestamptz_opt_overflow(m, v12, v7+int32(12))
	mBase = m.M
	v18 = m.ExcPending
	if v18 != 0 {
		return int32(0)
	} else {
		v21 = *(*int32)(unsafe.Add(mBase, uint32(v7)+12))
		if int32(0) < v21 {
			v29 = base.B2i32(v10 != int64(9223372036854775807))
		} else {
			if v21 < int32(0) {
				v29 = base.B2i32(v10 == int64(-9223372036854775807-1))
			} else {
				v29 = base.B2i32(v10 <= v15)
			}
		}
		m.G0 = v7 + int32(16)
		return v29
	}
}
func F_timestamptz_out(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v4 int32
	_ = v4
	var v6 int32
	_ = v6
	var v8 int32
	_ = v8
	var v9 int64
	_ = v9
	var v17 int32
	_ = v17
	var v20 int64
	_ = v20
	var v31 int32
	_ = v31
	var v34 int32
	_ = v34
	var v39 int32
	_ = v39
	var v41 int32
	_ = v41
	var v42 int32
	_ = v42
	var v44 int32
	_ = v44
	var v46 int32
	_ = v46
	var v50 int32
	_ = v50
	var v53 int32
	_ = v53
	var v57 int32
	_ = v57
	var v62 int32
	_ = v62
	var v64 int32
	_ = v64
	var v67 int64
	_ = v67
	var v69 int32
	_ = v69
	var v70 int32
	_ = v70
	v4 = m.G0
	v6 = v4 - int32(192)
	m.G0 = v6
	v8 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v9 = *(*int64)(unsafe.Add(mBase, uint32(v8)))
	if base.Ui64(v9-int64(9223372036854775807)) <= base.Ui64(int64(1)) {
		if v9 != int64(-9223372036854775807-1) {
			v64 = int32(*(*uint8)(unsafe.Add(mBase, _consts[1058])))
			*(*uint8)(unsafe.Add(mBase, uint32(v6)+8)) = uint8(v64)
			v67 = *(*int64)(unsafe.Add(mBase, _consts[1059]))
			*(*int64)(unsafe.Add(mBase, uint32(v6))) = v67
		} else {
			v17 = int32(*(*uint16)(unsafe.Add(mBase, _consts[1056])))
			*(*uint16)(unsafe.Add(mBase, uint32(v6)+8)) = uint16(v17)
			v20 = *(*int64)(unsafe.Add(mBase, _consts[1057]))
			*(*int64)(unsafe.Add(mBase, uint32(v6))) = v20
		}
		v69 = F_pstrdup(m, v6)
		mBase = m.M
		v70 = m.ExcPending
		if v70 != 0 {
			return int32(0)
		} else {
			m.G0 = v6 + int32(192)
			return v69
		}
	} else {
		v31 = F_timestamp2tm(m, v9, v6+int32(188), v6+int32(144), v6+int32(140), v6+int32(136), int32(0))
		mBase = m.M
		v34 = m.ExcPending
		if v34 != 0 {
			return int32(0)
		} else {
			if v31 == int32(0) {
				v39 = *(*int32)(unsafe.Add(mBase, uint32(v6)+140))
				v41 = *(*int32)(unsafe.Add(mBase, uint32(v6)+188))
				v42 = *(*int32)(unsafe.Add(mBase, uint32(v6)+136))
				v44 = *(*int32)(unsafe.Add(mBase, _consts[524]))
				F_EncodeDateTime(m, v6+int32(144), v39, int32(1), v41, v42, v44, v6)
				mBase = m.M
				v46 = m.ExcPending
				if v46 != 0 {
					return int32(0)
				} else {
					v69 = F_pstrdup(m, v6)
					mBase = m.M
					v70 = m.ExcPending
					if v70 != 0 {
						return int32(0)
					} else {
						m.G0 = v6 + int32(192)
						return v69
					}
				}
			} else {
				F_errstart_cold(m, int32(21), int32(0))
				mBase = m.M
				v50 = m.ExcPending
				if v50 != 0 {
					return int32(0)
				} else {
					F_errcode(m, int32(134217858))
					mBase = m.M
					v53 = m.ExcPending
					if v53 != 0 {
						return int32(0)
					} else {
						F_errmsg(m, int32(418494), int32(0))
						mBase = m.M
						v57 = m.ExcPending
						if v57 != 0 {
							return int32(0)
						} else {
							F_errfinish(m, int32(515856), int32(794), int32(71770))
							mBase = m.M
							v62 = m.ExcPending
							if v62 != 0 {
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
func F_timestamptz_pl_interval_internal(m *base.Module, l0 int64, l1 int32, l2 int32) int64 {
	mBase := m.M
	_ = mBase
	var v9 int32
	_ = v9
	var v11 int32
	_ = v11
	var v13 int32
	_ = v13
	var v18 int32
	_ = v18
	var v21 int64
	_ = v21
	var v22 int64
	_ = v22
	var v32 int32
	_ = v32
	var v35 int32
	_ = v35
	var v39 int32
	_ = v39
	var v44 int32
	_ = v44
	var v45 int32
	_ = v45
	var v48 int64
	_ = v48
	var v49 int64
	_ = v49
	var v57 int32
	_ = v57
	var v60 int32
	_ = v60
	var v64 int32
	_ = v64
	var v69 int32
	_ = v69
	var v76 int32
	_ = v76
	var v77 int32
	_ = v77
	var v85 int32
	_ = v85
	var v86 int32
	_ = v86
	var v87 int32
	_ = v87
	var v88 int32
	_ = v88
	var v89 int32
	_ = v89
	var v97 int32
	_ = v97
	var v98 int32
	_ = v98
	var v99 int32
	_ = v99
	var v100 int32
	_ = v100
	var v101 int32
	_ = v101
	var v102 int32
	_ = v102
	var v111 int32
	_ = v111
	var v112 int32
	_ = v112
	var v115 int32
	_ = v115
	var v116 int32
	_ = v116
	var v119 int32
	_ = v119
	var v127 int32
	_ = v127
	var v129 int32
	_ = v129
	var v131 int32
	_ = v131
	var v132 int32
	_ = v132
	var v134 int32
	_ = v134
	var v140 int32
	_ = v140
	var v144 int32
	_ = v144
	var v152 int32
	_ = v152
	var v156 int32
	_ = v156
	var v159 int32
	_ = v159
	var v163 int32
	_ = v163
	var v168 int32
	_ = v168
	var v172 int32
	_ = v172
	var v177 int32
	_ = v177
	var v180 int32
	_ = v180
	var v181 int32
	_ = v181
	var v189 int32
	_ = v189
	var v195 int32
	_ = v195
	var v196 int32
	_ = v196
	var v197 int32
	_ = v197
	var v201 int32
	_ = v201
	var v206 int64
	_ = v206
	var v207 int32
	_ = v207
	var v212 int32
	_ = v212
	var v217 int32
	_ = v217
	var v220 int32
	_ = v220
	var v223 int32
	_ = v223
	var v225 int32
	_ = v225
	var v226 int32
	_ = v226
	var v231 int32
	_ = v231
	var v232 int32
	_ = v232
	var v233 int32
	_ = v233
	var v238 int32
	_ = v238
	var v241 int32
	_ = v241
	var v244 int32
	_ = v244
	var v248 int32
	_ = v248
	var v253 int32
	_ = v253
	var v259 int64
	_ = v259
	var v268 int64
	_ = v268
	var v269 int64
	_ = v269
	var v271 int64
	_ = v271
	var v274 int64
	_ = v274
	var v275 int64
	_ = v275
	var v277 int64
	_ = v277
	var v278 int64
	_ = v278
	var v282 int64
	_ = v282
	var v289 int64
	_ = v289
	var v300 int64
	_ = v300
	var v301 int64
	_ = v301
	var v305 int32
	_ = v305
	var v306 int32
	_ = v306
	var v307 int32
	_ = v307
	var v308 int32
	_ = v308
	var v317 int64
	_ = v317
	var v320 int64
	_ = v320
	var v324 int32
	_ = v324
	var v329 int64
	_ = v329
	var v334 int64
	_ = v334
	var v339 int32
	_ = v339
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
	var v356 int32
	_ = v356
	var v357 int32
	_ = v357
	var v358 int32
	_ = v358
	var v363 int32
	_ = v363
	var v366 int32
	_ = v366
	var v369 int32
	_ = v369
	var v373 int32
	_ = v373
	var v378 int32
	_ = v378
	var v381 int32
	_ = v381
	var v382 int32
	_ = v382
	var v385 int32
	_ = v385
	var v399 int32
	_ = v399
	var v400 int32
	_ = v400
	var v401 int32
	_ = v401
	var v402 int32
	_ = v402
	var v408 int32
	_ = v408
	var v413 int32
	_ = v413
	var v416 int32
	_ = v416
	var v417 int32
	_ = v417
	var v418 int32
	_ = v418
	var v421 int32
	_ = v421
	var v423 int32
	_ = v423
	var v429 int32
	_ = v429
	var v433 int32
	_ = v433
	var v434 int32
	_ = v434
	var v436 int32
	_ = v436
	var v444 int32
	_ = v444
	var v448 int32
	_ = v448
	var v458 int32
	_ = v458
	var v465 int32
	_ = v465
	var v466 int32
	_ = v466
	var v467 int32
	_ = v467
	var v471 int32
	_ = v471
	var v476 int64
	_ = v476
	var v477 int32
	_ = v477
	var v482 int32
	_ = v482
	var v487 int32
	_ = v487
	var v490 int32
	_ = v490
	var v493 int32
	_ = v493
	var v495 int32
	_ = v495
	var v496 int32
	_ = v496
	var v501 int32
	_ = v501
	var v502 int32
	_ = v502
	var v503 int32
	_ = v503
	var v508 int32
	_ = v508
	var v511 int32
	_ = v511
	var v514 int32
	_ = v514
	var v518 int32
	_ = v518
	var v523 int32
	_ = v523
	var v529 int64
	_ = v529
	var v538 int64
	_ = v538
	var v539 int64
	_ = v539
	var v541 int64
	_ = v541
	var v544 int64
	_ = v544
	var v545 int64
	_ = v545
	var v547 int64
	_ = v547
	var v548 int64
	_ = v548
	var v552 int64
	_ = v552
	var v559 int64
	_ = v559
	var v570 int64
	_ = v570
	var v571 int64
	_ = v571
	var v575 int32
	_ = v575
	var v576 int32
	_ = v576
	var v577 int32
	_ = v577
	var v578 int32
	_ = v578
	var v587 int64
	_ = v587
	var v590 int64
	_ = v590
	var v594 int32
	_ = v594
	var v599 int64
	_ = v599
	var v604 int64
	_ = v604
	var v609 int64
	_ = v609
	var v612 int64
	_ = v612
	var v624 int64
	_ = v624
	var v632 int32
	_ = v632
	var v635 int32
	_ = v635
	var v639 int32
	_ = v639
	var v644 int32
	_ = v644
	var v648 int32
	_ = v648
	var v651 int32
	_ = v651
	var v655 int32
	_ = v655
	var v660 int32
	_ = v660
	var v664 int32
	_ = v664
	var v667 int32
	_ = v667
	var v671 int32
	_ = v671
	var v676 int32
	_ = v676
	var v680 int32
	_ = v680
	var v683 int32
	_ = v683
	var v687 int32
	_ = v687
	var v692 int32
	_ = v692
	var v696 int32
	_ = v696
	var v699 int32
	_ = v699
	var v703 int32
	_ = v703
	var v708 int32
	_ = v708
	var v712 int32
	_ = v712
	var v715 int32
	_ = v715
	var v719 int32
	_ = v719
	var v724 int32
	_ = v724
	var v731 int32
	_ = v731
	var v734 int32
	_ = v734
	var v738 int32
	_ = v738
	var v743 int32
	_ = v743
	var v750 int32
	_ = v750
	var v753 int32
	_ = v753
	var v757 int32
	_ = v757
	var v762 int32
	_ = v762
	v9 = m.G0
	v11 = v9 - int32(96)
	m.G0 = v11
	v13 = *(*int32)(unsafe.Add(mBase, uint32(l1)+12))
	if v13 != int32(2147483647) {
		goto L11
	} else {
		goto L12
	}
L1:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v750 = m.ExcPending
	if v750 != 0 {
		goto L18
	} else {
		goto L156
	}
L2:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v731 = m.ExcPending
	if v731 != 0 {
		goto L18
	} else {
		goto L152
	}
L3:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v712 = m.ExcPending
	if v712 != 0 {
		goto L18
	} else {
		goto L148
	}
L4:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v696 = m.ExcPending
	if v696 != 0 {
		goto L18
	} else {
		goto L144
	}
L5:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v680 = m.ExcPending
	if v680 != 0 {
		goto L18
	} else {
		goto L140
	}
L6:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v664 = m.ExcPending
	if v664 != 0 {
		goto L18
	} else {
		goto L136
	}
L7:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v648 = m.ExcPending
	if v648 != 0 {
		goto L18
	} else {
		goto L132
	}
L8:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v632 = m.ExcPending
	if v632 != 0 {
		goto L18
	} else {
		goto L128
	}
L9:
	;
	m.G0 = v11 + int32(96)
	return v624
L10:
	;
	if base.Ui64(l0-int64(9223372036854775807)) < base.Ui64(int64(2)) {
		goto L30
	} else {
		goto L31
	}
L11:
	;
	if v13 != int32(-2147483648) {
		goto L10
	} else {
		goto L14
	}
L12:
	;
	goto L13
L13:
	;
	v45 = *(*int32)(unsafe.Add(mBase, uint32(l1)+8))
	if v45 != int32(2147483647) {
		goto L10
	} else {
		goto L23
	}
L14:
	;
	v18 = *(*int32)(unsafe.Add(mBase, uint32(l1)+8))
	if v18 != int32(-2147483648) {
		goto L10
	} else {
		goto L15
	}
L15:
	;
	v21 = int64(-9223372036854775807 - 1)
	v22 = *(*int64)(unsafe.Add(mBase, uint32(l1)))
	if v22 != v21 {
		goto L10
	} else {
		goto L16
	}
L16:
	;
	if l0 != int64(9223372036854775807) {
		v624 = v21
		goto L9
	} else {
		goto L17
	}
L17:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v32 = m.ExcPending
	if v32 != 0 {
		goto L18
	} else {
		goto L19
	}
L18:
	;
	return int64(0)
L19:
	;
	F_errcode(m, int32(134217858))
	mBase = m.M
	v35 = m.ExcPending
	if v35 != 0 {
		goto L18
	} else {
		goto L20
	}
L20:
	;
	F_errmsg(m, int32(418494), int32(0))
	mBase = m.M
	v39 = m.ExcPending
	if v39 != 0 {
		goto L18
	} else {
		goto L21
	}
L21:
	;
	F_errfinish(m, int32(515856), int32(3269), int32(324081))
	mBase = m.M
	v44 = m.ExcPending
	if v44 != 0 {
		goto L18
	} else {
		goto L22
	}
L22:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L23:
	;
	v48 = int64(9223372036854775807)
	v49 = *(*int64)(unsafe.Add(mBase, uint32(l1)))
	if v49 != v48 {
		goto L10
	} else {
		goto L24
	}
L24:
	;
	if l0 != int64(-9223372036854775807-1) {
		v624 = v48
		goto L9
	} else {
		goto L25
	}
L25:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v57 = m.ExcPending
	if v57 != 0 {
		goto L18
	} else {
		goto L26
	}
L26:
	;
	F_errcode(m, int32(134217858))
	mBase = m.M
	v60 = m.ExcPending
	if v60 != 0 {
		goto L18
	} else {
		goto L27
	}
L27:
	;
	F_errmsg(m, int32(418494), int32(0))
	mBase = m.M
	v64 = m.ExcPending
	if v64 != 0 {
		goto L18
	} else {
		goto L28
	}
L28:
	;
	F_errfinish(m, int32(515856), int32(3278), int32(324081))
	mBase = m.M
	v69 = m.ExcPending
	if v69 != 0 {
		goto L18
	} else {
		goto L29
	}
L29:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L30:
	;
	v624 = l0
	goto L9
L31:
	;
	goto L32
L32:
	;
	v76 = *(*int32)(unsafe.Add(mBase, _consts[526]))
	if l2 != 0 {
		goto L33
	} else {
		goto L34
	}
L33:
	;
	v77 = l2
	goto L35
L34:
	;
	v77 = v76
	goto L35
L35:
	;
	if v13 != 0 {
		goto L36
	} else {
		goto L37
	}
L36:
	;
	v85 = F_timestamp2tm(m, l0, v11+int32(92), v11+int32(48), v11+int32(44), int32(0), v77)
	mBase = m.M
	v86 = m.ExcPending
	if v86 != 0 {
		goto L18
	} else {
		goto L39
	}
L37:
	;
	v334 = l0
	goto L38
L38:
	;
	v339 = *(*int32)(unsafe.Add(mBase, uint32(l1)+8))
	if v339 != 0 {
		goto L84
	} else {
		goto L85
	}
L39:
	;
	if v85 != 0 {
		goto L8
	} else {
		goto L40
	}
L40:
	;
	v87 = *(*int32)(unsafe.Add(mBase, uint32(v11)+64))
	v88 = *(*int32)(unsafe.Add(mBase, uint32(l1)+12))
	v89 = v87 + v88
	*(*int32)(unsafe.Add(mBase, uint32(v11)+64)) = v89
	if base.B2i32(v88 < int32(0)) != base.B2i32(v89 < v87) {
		goto L7
	} else {
		goto L41
	}
L41:
	;
	if int32(13) <= v89 {
		goto L44
	} else {
		goto L45
	}
L42:
	;
	v134 = *(*int32)(unsafe.Add(mBase, uint32(v11)+60))
	if v132&int32(3) == int32(0) {
		goto L52
	} else {
		goto L53
	}
L43:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v11)+64)) = v129
	v131 = v129
	v132 = v127
	goto L42
L44:
	;
	v97 = int32(1)
	v98 = v89 - v97
	v99 = int32(12)
	v100 = base.I32_div_u_s(v98, v99)
	v101 = *(*int32)(unsafe.Add(mBase, uint32(v11)+68))
	v102 = v100 + v101
	*(*int32)(unsafe.Add(mBase, uint32(v11)+68)) = v102
	v127 = v102
	v129 = v98 - v100*v99 + v97
	goto L43
L45:
	;
	goto L46
L46:
	;
	if int32(0) < v89 {
		goto L47
	} else {
		goto L48
	}
L47:
	;
	v111 = *(*int32)(unsafe.Add(mBase, uint32(v11)+68))
	v131 = v89
	v132 = v111
	goto L42
L48:
	;
	goto L49
L49:
	;
	v112 = *(*int32)(unsafe.Add(mBase, uint32(v11)+68))
	v115 = int32(12)
	v116 = base.I32_div_u_s(int32(0)-v89, v115)
	v119 = v112 + (v116 ^ int32(-1))
	*(*int32)(unsafe.Add(mBase, uint32(v11)+68)) = v119
	v127 = v119
	v129 = v116*v115 + v89 + v115
	goto L43
L50:
	;
	v195 = m.G0
	v196 = int32(16)
	v197 = v195 - v196
	m.G0 = v197
	v201 = F_DetermineTimeZoneOffsetInternal(m, v11+int32(48), v77, v197+int32(8))
	mBase = m.M
	m.G0 = v197 + v196
	goto L61
L51:
	;
	v189 = *(*int32)(unsafe.Add(mBase, uint32(v181*int32(52)+int32(1686864)+v180<<(uint(int32(2))%32))))
	*(*int32)(unsafe.Add(mBase, uint32(v11)+60)) = v189
	goto L50
L52:
	;
	v140 = base.I32_rem_s(v132, int32(100))
	if v140 == int32(0) {
		goto L55
	} else {
		goto L56
	}
L53:
	;
	goto L54
L54:
	;
	v172 = v131 - int32(1)
	v177 = *(*int32)(unsafe.Add(mBase, uint32(v172<<(uint(int32(2))%32))+uint32(_consts[1062])))
	if v134 <= v177 {
		goto L50
	} else {
		goto L60
	}
L55:
	;
	v144 = base.I32_rem_s(v132, int32(400))
	v152 = v131 - int32(1)
	v156 = *(*int32)(unsafe.Add(mBase, uint32(base.B2i32(v144 == int32(0))*int32(52)+int32(1686864)+v152<<(uint(int32(2))%32))))
	if v134 <= v156 {
		goto L50
	} else {
		goto L58
	}
L56:
	;
	goto L57
L57:
	;
	v163 = v131 - int32(1)
	v168 = *(*int32)(unsafe.Add(mBase, uint32(v163<<(uint(int32(2))%32))+uint32(_consts[1061])))
	if v134 <= v168 {
		goto L50
	} else {
		goto L59
	}
L58:
	;
	v159 = base.I32_rem_s(v132, int32(400))
	v180 = v152
	v181 = base.B2i32(v159 == int32(0))
	goto L51
L59:
	;
	v180 = v163
	v181 = int32(1)
	goto L51
L60:
	;
	v180 = v172
	v181 = int32(0)
	goto L51
L61:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v11)+92)) = v201
	v206 = int64(*(*int32)(unsafe.Add(mBase, uint32(v11)+44)))
	v207 = *(*int32)(unsafe.Add(mBase, uint32(v11)+68))
	if v207 <= int32(-4713) {
		goto L63
	} else {
		goto L64
	}
L62:
	;
	v225 = v11 + int32(24)
	v226 = *(*int32)(unsafe.Add(mBase, uint32(v11)+60))
	v231 = base.B2i32(int32(2) < v223)
	if int32(2) < v223 {
		goto L74
	} else {
		goto L75
	}
L63:
	;
	if v207 != int32(-4713) {
		goto L1
	} else {
		goto L66
	}
L64:
	;
	goto L65
L65:
	;
	if v207 <= int32(5874897) {
		goto L68
	} else {
		goto L69
	}
L66:
	;
	v212 = *(*int32)(unsafe.Add(mBase, uint32(v11)+64))
	if int32(10) < v212 {
		v223 = v212
		goto L62
	} else {
		goto L67
	}
L67:
	;
	goto L1
L68:
	;
	v217 = *(*int32)(unsafe.Add(mBase, uint32(v11)+64))
	v223 = v217
	goto L62
L69:
	;
	goto L70
L70:
	;
	if v207 != int32(5874898) {
		goto L1
	} else {
		goto L71
	}
L71:
	;
	v220 = *(*int32)(unsafe.Add(mBase, uint32(v11)+64))
	if int32(5) < v220 {
		goto L1
	} else {
		goto L72
	}
L72:
	;
	v223 = v220
	goto L62
L73:
	;
	v259 = base.I64_extend_i32_s(v226 + v233*int32(365) + v238 + v241 + v244 + v253 - int32(32167) - int32(2451545))
	v268 = int64(32)
	v269 = int64(20)
	v271 = int64(base.Ui64(v259) >> (uint(v268) % 64))
	v274 = int64(4294967295)
	v275 = int64(500654080)
	v277 = v259 & v274
	v278 = v275 * v277
	v282 = int64(base.Ui64(v278)>>(uint(v268)%64)) + v275*v271
	v289 = v277*v269 + v282&v274
	*(*int64)(unsafe.Add(mBase, uint32(v225)+8)) = v259*int64(0) + v259>>(uint(int64(63))%64)*int64(86400000000) + v269*v271 + int64(base.Ui64(v282)>>(uint(v268)%64)) + int64(base.Ui64(v289)>>(uint(v268)%64))
	*(*int64)(unsafe.Add(mBase, uint32(v225))) = v278&v274 | v289<<(uint(v268)%64)
	goto L80
L74:
	;
	v232 = int32(4800)
	goto L76
L75:
	;
	v232 = int32(4799)
	goto L76
L76:
	;
	v233 = v232 + v207
	v238 = base.I32_div_s(v233, int32(4))
	v241 = base.I32_div_s(v233, int32(-100))
	v244 = base.I32_div_s(v233, int32(400))
	if int32(2) < v223 {
		goto L77
	} else {
		goto L78
	}
L77:
	;
	v248 = int32(1)
	goto L79
L78:
	;
	v248 = int32(13)
	goto L79
L79:
	;
	v253 = base.I32_div_s((v248+v223)*int32(7834), int32(256))
	goto L73
L80:
	;
	v300 = *(*int64)(unsafe.Add(mBase, uint32(v11)+32))
	v301 = *(*int64)(unsafe.Add(mBase, uint32(v11)+24))
	if v300 != v301>>(uint(int64(63))%64) {
		goto L1
	} else {
		goto L81
	}
L81:
	;
	v305 = *(*int32)(unsafe.Add(mBase, uint32(v11)+48))
	v306 = *(*int32)(unsafe.Add(mBase, uint32(v11)+52))
	v307 = *(*int32)(unsafe.Add(mBase, uint32(v11)+56))
	v308 = int32(60)
	v317 = base.I64_extend_i32_s(v305+(v306+v307*v308)*v308)*int64(1000000) + v206
	v320 = v301 + v317
	if base.B2i32(v317 < int64(0))^base.B2i32(v320 < v301) != 0 {
		goto L1
	} else {
		goto L82
	}
L82:
	;
	v324 = *(*int32)(unsafe.Add(mBase, uint32(v11)+92))
	v329 = base.I64_extend_i32_s(int32(0)-v324)*int64(-1000000) + v320
	if base.Ui64(int64(-9011559254509551616)) <= base.Ui64(v329+int64(211813488000000000)) {
		goto L1
	} else {
		goto L83
	}
L83:
	;
	v334 = v329
	goto L38
L84:
	;
	v347 = F_timestamp2tm(m, v334, v11+int32(92), v11+int32(48), v11+int32(44), int32(0), v77)
	mBase = m.M
	v348 = m.ExcPending
	if v348 != 0 {
		goto L18
	} else {
		goto L87
	}
L85:
	;
	v604 = v334
	goto L86
L86:
	;
	v609 = *(*int64)(unsafe.Add(mBase, uint32(l1)))
	v612 = v604 + v609
	if base.B2i32(v609 < int64(0)) != base.B2i32(v612 < v604) {
		goto L4
	} else {
		goto L126
	}
L87:
	;
	if v347 != 0 {
		goto L6
	} else {
		goto L88
	}
L88:
	;
	v349 = *(*int32)(unsafe.Add(mBase, uint32(v11)+68))
	v350 = *(*int32)(unsafe.Add(mBase, uint32(v11)+64))
	v351 = *(*int32)(unsafe.Add(mBase, uint32(v11)+60))
	v356 = base.B2i32(int32(2) < v350)
	if int32(2) < v350 {
		goto L90
	} else {
		goto L91
	}
L89:
	;
	v382 = *(*int32)(unsafe.Add(mBase, uint32(l1)+8))
	v385 = v381 + v382
	if base.B2i32(v382 < int32(0))^base.B2i32(v385 < v381) != 0 {
		goto L5
	} else {
		goto L96
	}
L90:
	;
	v357 = int32(4800)
	goto L92
L91:
	;
	v357 = int32(4799)
	goto L92
L92:
	;
	v358 = v357 + v349
	v363 = base.I32_div_s(v358, int32(4))
	v366 = base.I32_div_s(v358, int32(-100))
	v369 = base.I32_div_s(v358, int32(400))
	if int32(2) < v350 {
		goto L93
	} else {
		goto L94
	}
L93:
	;
	v373 = int32(1)
	goto L95
L94:
	;
	v373 = int32(13)
	goto L95
L95:
	;
	v378 = base.I32_div_s((v373+v350)*int32(7834), int32(256))
	v381 = v351 + v358*int32(365) + v363 + v366 + v369 + v378 - int32(32167)
	goto L89
L96:
	;
	if v385 <= int32(-2) {
		goto L5
	} else {
		goto L97
	}
L97:
	;
	v399 = v385 + int32(32044)
	v400 = int32(146097)
	v401 = base.I32_div_u_s(v399, v400)
	v402 = int32(3)
	v408 = int32(2)
	v413 = base.I32_div_u_s((v401*int32(1073595727)+v399)<<(uint(v408)%32)|v402, v400)
	v416 = v385 + v401*v402 + v413 + int32(32104)
	v417 = int32(1461)
	v418 = base.I32_div_u_s(v416, v417)
	v421 = v418*int32(-1461) + v416
	v423 = v421 << (uint(v408) % 32)
	if base.Ui32(v417) <= base.Ui32(v423) {
		goto L100
	} else {
		goto L101
	}
L98:
	;
	v465 = m.G0
	v466 = int32(16)
	v467 = v465 - v466
	m.G0 = v467
	v471 = F_DetermineTimeZoneOffsetInternal(m, v11+int32(48), v77, v467+int32(8))
	mBase = m.M
	m.G0 = v467 + v466
	goto L103
L99:
	;
	v436 = base.I32_div_u_s(v423, int32(1461))
	*(*int32)(unsafe.Add(mBase, uint32(v11+int32(68)))) = v436 + v418<<(uint(int32(2))%32) - int32(4800)
	v444 = v434 + int32(123)
	v448 = int32(base.Ui32(v444*int32(2141)) >> (uint(int32(16)) % 32))
	*(*int32)(unsafe.Add(mBase, uint32(v11+int32(60)))) = v444 - int32(base.Ui32(v448*int32(7834))>>(uint(int32(8))%32))
	v458 = base.I32_rem_u_s(v448+int32(10), int32(12))
	*(*int32)(unsafe.Add(mBase, uint32(v11-int32(-64)))) = v458 + int32(1)
	goto L98
L100:
	;
	v429 = base.I32_rem_u_s(v421+int32(305), int32(365))
	v434 = v429
	goto L99
L101:
	;
	goto L102
L102:
	;
	v433 = base.I32_rem_u_s(v421+int32(306), int32(366))
	v434 = v433
	goto L99
L103:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v11)+92)) = v471
	v476 = int64(*(*int32)(unsafe.Add(mBase, uint32(v11)+44)))
	v477 = *(*int32)(unsafe.Add(mBase, uint32(v11)+68))
	if v477 <= int32(-4713) {
		goto L105
	} else {
		goto L106
	}
L104:
	;
	v495 = v11 + int32(8)
	v496 = *(*int32)(unsafe.Add(mBase, uint32(v11)+60))
	v501 = base.B2i32(int32(2) < v493)
	if int32(2) < v493 {
		goto L116
	} else {
		goto L117
	}
L105:
	;
	if v477 != int32(-4713) {
		goto L2
	} else {
		goto L108
	}
L106:
	;
	goto L107
L107:
	;
	if v477 <= int32(5874897) {
		goto L110
	} else {
		goto L111
	}
L108:
	;
	v482 = *(*int32)(unsafe.Add(mBase, uint32(v11)+64))
	if int32(10) < v482 {
		v493 = v482
		goto L104
	} else {
		goto L109
	}
L109:
	;
	goto L2
L110:
	;
	v487 = *(*int32)(unsafe.Add(mBase, uint32(v11)+64))
	v493 = v487
	goto L104
L111:
	;
	goto L112
L112:
	;
	if v477 != int32(5874898) {
		goto L2
	} else {
		goto L113
	}
L113:
	;
	v490 = *(*int32)(unsafe.Add(mBase, uint32(v11)+64))
	if int32(5) < v490 {
		goto L2
	} else {
		goto L114
	}
L114:
	;
	v493 = v490
	goto L104
L115:
	;
	v529 = base.I64_extend_i32_s(v496 + v503*int32(365) + v508 + v511 + v514 + v523 - int32(32167) - int32(2451545))
	v538 = int64(32)
	v539 = int64(20)
	v541 = int64(base.Ui64(v529) >> (uint(v538) % 64))
	v544 = int64(4294967295)
	v545 = int64(500654080)
	v547 = v529 & v544
	v548 = v545 * v547
	v552 = int64(base.Ui64(v548)>>(uint(v538)%64)) + v545*v541
	v559 = v547*v539 + v552&v544
	*(*int64)(unsafe.Add(mBase, uint32(v495)+8)) = v529*int64(0) + v529>>(uint(int64(63))%64)*int64(86400000000) + v539*v541 + int64(base.Ui64(v552)>>(uint(v538)%64)) + int64(base.Ui64(v559)>>(uint(v538)%64))
	*(*int64)(unsafe.Add(mBase, uint32(v495))) = v548&v544 | v559<<(uint(v538)%64)
	goto L122
L116:
	;
	v502 = int32(4800)
	goto L118
L117:
	;
	v502 = int32(4799)
	goto L118
L118:
	;
	v503 = v502 + v477
	v508 = base.I32_div_s(v503, int32(4))
	v511 = base.I32_div_s(v503, int32(-100))
	v514 = base.I32_div_s(v503, int32(400))
	if int32(2) < v493 {
		goto L119
	} else {
		goto L120
	}
L119:
	;
	v518 = int32(1)
	goto L121
L120:
	;
	v518 = int32(13)
	goto L121
L121:
	;
	v523 = base.I32_div_s((v518+v493)*int32(7834), int32(256))
	goto L115
L122:
	;
	v570 = *(*int64)(unsafe.Add(mBase, uint32(v11)+16))
	v571 = *(*int64)(unsafe.Add(mBase, uint32(v11)+8))
	if v570 != v571>>(uint(int64(63))%64) {
		goto L2
	} else {
		goto L123
	}
L123:
	;
	v575 = *(*int32)(unsafe.Add(mBase, uint32(v11)+48))
	v576 = *(*int32)(unsafe.Add(mBase, uint32(v11)+52))
	v577 = *(*int32)(unsafe.Add(mBase, uint32(v11)+56))
	v578 = int32(60)
	v587 = base.I64_extend_i32_s(v575+(v576+v577*v578)*v578)*int64(1000000) + v476
	v590 = v571 + v587
	if base.B2i32(v587 < int64(0))^base.B2i32(v590 < v571) != 0 {
		goto L2
	} else {
		goto L124
	}
L124:
	;
	v594 = *(*int32)(unsafe.Add(mBase, uint32(v11)+92))
	v599 = base.I64_extend_i32_s(int32(0)-v594)*int64(-1000000) + v590
	if base.Ui64(int64(-9011559254509551616)) <= base.Ui64(v599+int64(211813488000000000)) {
		goto L2
	} else {
		goto L125
	}
L125:
	;
	v604 = v599
	goto L86
L126:
	;
	if base.Ui64(int64(-9011559254509551616)) <= base.Ui64(v612+int64(211813488000000000)) {
		goto L3
	} else {
		goto L127
	}
L127:
	;
	v624 = v612
	goto L9
L128:
	;
	F_errcode(m, int32(134217858))
	mBase = m.M
	v635 = m.ExcPending
	if v635 != 0 {
		goto L18
	} else {
		goto L129
	}
L129:
	;
	F_errmsg(m, int32(418494), int32(0))
	mBase = m.M
	v639 = m.ExcPending
	if v639 != 0 {
		goto L18
	} else {
		goto L130
	}
L130:
	;
	F_errfinish(m, int32(515856), int32(3299), int32(324081))
	mBase = m.M
	v644 = m.ExcPending
	if v644 != 0 {
		goto L18
	} else {
		goto L131
	}
L131:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L132:
	;
	F_errcode(m, int32(134217858))
	mBase = m.M
	v651 = m.ExcPending
	if v651 != 0 {
		goto L18
	} else {
		goto L133
	}
L133:
	;
	F_errmsg(m, int32(418494), int32(0))
	mBase = m.M
	v655 = m.ExcPending
	if v655 != 0 {
		goto L18
	} else {
		goto L134
	}
L134:
	;
	F_errfinish(m, int32(515856), int32(3304), int32(324081))
	mBase = m.M
	v660 = m.ExcPending
	if v660 != 0 {
		goto L18
	} else {
		goto L135
	}
L135:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L136:
	;
	F_errcode(m, int32(134217858))
	mBase = m.M
	v667 = m.ExcPending
	if v667 != 0 {
		goto L18
	} else {
		goto L137
	}
L137:
	;
	F_errmsg(m, int32(418494), int32(0))
	mBase = m.M
	v671 = m.ExcPending
	if v671 != 0 {
		goto L18
	} else {
		goto L138
	}
L138:
	;
	F_errfinish(m, int32(515856), int32(3338), int32(324081))
	mBase = m.M
	v676 = m.ExcPending
	if v676 != 0 {
		goto L18
	} else {
		goto L139
	}
L139:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L140:
	;
	F_errcode(m, int32(134217858))
	mBase = m.M
	v683 = m.ExcPending
	if v683 != 0 {
		goto L18
	} else {
		goto L141
	}
L141:
	;
	F_errmsg(m, int32(418494), int32(0))
	mBase = m.M
	v687 = m.ExcPending
	if v687 != 0 {
		goto L18
	} else {
		goto L142
	}
L142:
	;
	F_errfinish(m, int32(515856), int32(3352), int32(324081))
	mBase = m.M
	v692 = m.ExcPending
	if v692 != 0 {
		goto L18
	} else {
		goto L143
	}
L143:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L144:
	;
	F_errcode(m, int32(134217858))
	mBase = m.M
	v699 = m.ExcPending
	if v699 != 0 {
		goto L18
	} else {
		goto L145
	}
L145:
	;
	F_errmsg(m, int32(418494), int32(0))
	mBase = m.M
	v703 = m.ExcPending
	if v703 != 0 {
		goto L18
	} else {
		goto L146
	}
L146:
	;
	F_errfinish(m, int32(515856), int32(3366), int32(324081))
	mBase = m.M
	v708 = m.ExcPending
	if v708 != 0 {
		goto L18
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
	F_errcode(m, int32(134217858))
	mBase = m.M
	v715 = m.ExcPending
	if v715 != 0 {
		goto L18
	} else {
		goto L149
	}
L149:
	;
	F_errmsg(m, int32(418494), int32(0))
	mBase = m.M
	v719 = m.ExcPending
	if v719 != 0 {
		goto L18
	} else {
		goto L150
	}
L150:
	;
	F_errfinish(m, int32(515856), int32(3371), int32(324081))
	mBase = m.M
	v724 = m.ExcPending
	if v724 != 0 {
		goto L18
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
	F_errcode(m, int32(134217858))
	mBase = m.M
	v734 = m.ExcPending
	if v734 != 0 {
		goto L18
	} else {
		goto L153
	}
L153:
	;
	F_errmsg(m, int32(418494), int32(0))
	mBase = m.M
	v738 = m.ExcPending
	if v738 != 0 {
		goto L18
	} else {
		goto L154
	}
L154:
	;
	F_errfinish(m, int32(515856), int32(3360), int32(324081))
	mBase = m.M
	v743 = m.ExcPending
	if v743 != 0 {
		goto L18
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
	F_errcode(m, int32(134217858))
	mBase = m.M
	v753 = m.ExcPending
	if v753 != 0 {
		goto L18
	} else {
		goto L157
	}
L157:
	;
	F_errmsg(m, int32(418494), int32(0))
	mBase = m.M
	v757 = m.ExcPending
	if v757 != 0 {
		goto L18
	} else {
		goto L158
	}
L158:
	;
	F_errfinish(m, int32(515856), int32(3325), int32(324081))
	mBase = m.M
	v762 = m.ExcPending
	if v762 != 0 {
		goto L18
	} else {
		goto L159
	}
L159:
	;
	base.Wasm_trap_unreachable()
	for {
	}
}
func F_timestamptz_trunc_internal(m *base.Module, l0 int32, l1 int64, l2 int32) int64 {
	mBase := m.M
	_ = mBase
	var v11 int32
	_ = v11
	var v13 int32
	_ = v13
	var v15 int32
	_ = v15
	var v16 int32
	_ = v16
	var v19 int32
	_ = v19
	var v21 int32
	_ = v21
	var v22 int32
	_ = v22
	var v25 int32
	_ = v25
	var v27 int32
	_ = v27
	var v36 int32
	_ = v36
	var v39 int32
	_ = v39
	var v40 int32
	_ = v40
	var v44 int32
	_ = v44
	var v50 int32
	_ = v50
	var v52 int32
	_ = v52
	var v55 int32
	_ = v55
	var v57 int32
	_ = v57
	var v64 int32
	_ = v64
	var v66 int32
	_ = v66
	var v69 int32
	_ = v69
	var v75 int32
	_ = v75
	var v77 int32
	_ = v77
	var v84 int32
	_ = v84
	var v85 int32
	_ = v85
	var v86 int32
	_ = v86
	var v90 int32
	_ = v90
	var v93 int32
	_ = v93
	var v97 int32
	_ = v97
	var v98 int32
	_ = v98
	var v101 int32
	_ = v101
	var v108 int32
	_ = v108
	var v115 int32
	_ = v115
	var v116 int32
	_ = v116
	var v125 int32
	_ = v125
	var v132 int32
	_ = v132
	var v140 int32
	_ = v140
	var v143 int32
	_ = v143
	var v145 int32
	_ = v145
	var v146 int32
	_ = v146
	var v151 int32
	_ = v151
	var v156 int32
	_ = v156
	var v157 int32
	_ = v157
	var v165 int32
	_ = v165
	var v166 int32
	_ = v166
	var v167 int32
	_ = v167
	var v168 int32
	_ = v168
	var v171 int32
	_ = v171
	var v172 int32
	_ = v172
	var v173 int32
	_ = v173
	var v174 int32
	_ = v174
	var v175 int32
	_ = v175
	var v176 int32
	_ = v176
	var v178 int32
	_ = v178
	var v179 int32
	_ = v179
	var v181 int32
	_ = v181
	var v184 int32
	_ = v184
	var v187 int32
	_ = v187
	var v191 int32
	_ = v191
	var v194 int32
	_ = v194
	var v195 int32
	_ = v195
	var v196 int32
	_ = v196
	var v198 int32
	_ = v198
	var v201 int32
	_ = v201
	var v205 int32
	_ = v205
	var v208 int32
	_ = v208
	var v209 int32
	_ = v209
	var v212 int32
	_ = v212
	var v214 int32
	_ = v214
	var v216 int32
	_ = v216
	var v218 int32
	_ = v218
	var v225 int32
	_ = v225
	var v230 int32
	_ = v230
	var v233 int32
	_ = v233
	var v241 int32
	_ = v241
	var v243 int32
	_ = v243
	var v254 int32
	_ = v254
	var v259 int32
	_ = v259
	var v262 int32
	_ = v262
	var v265 int32
	_ = v265
	var v274 int32
	_ = v274
	var v277 int32
	_ = v277
	var v279 int32
	_ = v279
	var v283 int32
	_ = v283
	var v284 int32
	_ = v284
	var v289 int32
	_ = v289
	var v292 int32
	_ = v292
	var v296 int32
	_ = v296
	var v297 int32
	_ = v297
	var v298 int32
	_ = v298
	var v299 int32
	_ = v299
	var v305 int32
	_ = v305
	var v310 int32
	_ = v310
	var v313 int32
	_ = v313
	var v314 int32
	_ = v314
	var v315 int32
	_ = v315
	var v318 int32
	_ = v318
	var v320 int32
	_ = v320
	var v326 int32
	_ = v326
	var v330 int32
	_ = v330
	var v331 int32
	_ = v331
	var v333 int32
	_ = v333
	var v341 int32
	_ = v341
	var v345 int32
	_ = v345
	var v355 int32
	_ = v355
	var v359 int32
	_ = v359
	var v365 int32
	_ = v365
	var v367 int32
	_ = v367
	var v370 int32
	_ = v370
	var v375 int32
	_ = v375
	var v376 int32
	_ = v376
	var v382 int32
	_ = v382
	var v385 int32
	_ = v385
	var v388 int32
	_ = v388
	var v393 int32
	_ = v393
	var v397 int32
	_ = v397
	var v401 int32
	_ = v401
	var v403 int32
	_ = v403
	var v406 int32
	_ = v406
	var v411 int32
	_ = v411
	var v424 int32
	_ = v424
	var v432 int32
	_ = v432
	var v433 int32
	_ = v433
	var v434 int32
	_ = v434
	var v438 int32
	_ = v438
	var v443 int32
	_ = v443
	var v445 int32
	_ = v445
	var v450 int32
	_ = v450
	var v453 int32
	_ = v453
	var v455 int32
	_ = v455
	var v456 int32
	_ = v456
	var v463 int32
	_ = v463
	var v468 int32
	_ = v468
	var v474 int32
	_ = v474
	var v479 int32
	_ = v479
	var v484 int32
	_ = v484
	var v489 int32
	_ = v489
	var v492 int32
	_ = v492
	var v495 int32
	_ = v495
	var v497 int32
	_ = v497
	var v498 int32
	_ = v498
	var v503 int32
	_ = v503
	var v504 int32
	_ = v504
	var v505 int32
	_ = v505
	var v510 int32
	_ = v510
	var v513 int32
	_ = v513
	var v516 int32
	_ = v516
	var v520 int32
	_ = v520
	var v525 int32
	_ = v525
	var v531 int64
	_ = v531
	var v540 int64
	_ = v540
	var v541 int64
	_ = v541
	var v543 int64
	_ = v543
	var v546 int64
	_ = v546
	var v547 int64
	_ = v547
	var v549 int64
	_ = v549
	var v550 int64
	_ = v550
	var v554 int64
	_ = v554
	var v561 int64
	_ = v561
	var v572 int64
	_ = v572
	var v573 int64
	_ = v573
	var v578 int32
	_ = v578
	var v579 int32
	_ = v579
	var v580 int32
	_ = v580
	var v581 int32
	_ = v581
	var v590 int64
	_ = v590
	var v593 int64
	_ = v593
	var v597 int32
	_ = v597
	var v602 int64
	_ = v602
	var v613 int32
	_ = v613
	var v616 int32
	_ = v616
	var v620 int32
	_ = v620
	var v625 int32
	_ = v625
	var v629 int32
	_ = v629
	var v632 int32
	_ = v632
	var v634 int32
	_ = v634
	var v635 int32
	_ = v635
	var v642 int32
	_ = v642
	var v647 int32
	_ = v647
	var v649 int64
	_ = v649
	var v663 int32
	_ = v663
	var v666 int32
	_ = v666
	var v670 int32
	_ = v670
	var v675 int32
	_ = v675
	v11 = m.G0
	v13 = v11 - int32(112)
	m.G0 = v13
	v15 = int32(1)
	v16 = l0 + v15
	v19 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0))))
	v21 = v19 & v15
	if v21 != 0 {
		goto L3
	} else {
		goto L4
	}
L1:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v663 = m.ExcPending
	if v663 != 0 {
		goto L17
	} else {
		goto L147
	}
L2:
	;
	m.G0 = v13 + int32(112)
	return v649
L3:
	;
	v22 = v16
	goto L5
L4:
	;
	v22 = l0 + int32(4)
	goto L5
L5:
	;
	if v19 == int32(1) {
		goto L7
	} else {
		goto L8
	}
L6:
	;
	v52 = F_downcase_truncate_identifier(m, v22, v50, int32(0))
	mBase = m.M
	v55 = m.ExcPending
	if v55 != 0 {
		goto L17
	} else {
		goto L18
	}
L7:
	;
	v25 = int32(4)
	v27 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v16))))
	if v27&int32(254) == int32(2) {
		goto L10
	} else {
		goto L11
	}
L8:
	;
	goto L9
L9:
	;
	v40 = int32(1)
	if v21 != 0 {
		v50 = int32(base.Ui32(v19)>>(uint(v40)%32)) - v40
		goto L6
	} else {
		goto L16
	}
L10:
	;
	v36 = v25
	goto L12
L11:
	;
	v36 = base.B2i32(v27 == int32(18)) << (uint(v25) % 32)
	goto L12
L12:
	;
	if v27 == int32(1) {
		goto L13
	} else {
		goto L14
	}
L13:
	;
	v39 = v25
	goto L15
L14:
	;
	v39 = v36
	goto L15
L15:
	;
	v50 = v39
	goto L6
L16:
	;
	v44 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v50 = int32(base.Ui32(v44)>>(uint(int32(2))%32)) - int32(4)
	goto L6
L17:
	;
	return int64(0)
L18:
	;
	v57 = v13 + int32(104)
	v64 = *(*int32)(unsafe.Add(mBase, _consts[1063]))
	if v64 != 0 {
		goto L21
	} else {
		goto L22
	}
L19:
	;
	if v125 == int32(17) {
		goto L38
	} else {
		goto L39
	}
L20:
	;
	*(*int32)(unsafe.Add(mBase, _consts[1063])) = v108
	v115 = int32(*(*int8)(unsafe.Add(mBase, uint32(v108)+11)))
	v116 = *(*int32)(unsafe.Add(mBase, uint32(v108)+12))
	*(*int32)(unsafe.Add(mBase, uint32(v57))) = v116
	v125 = v115
	goto L19
L21:
	;
	v66 = F_strncmp(m, v52, v64, int32(10))
	mBase = m.M
	if v66 == int32(0) {
		v108 = v64
		goto L20
	} else {
		goto L24
	}
L22:
	;
	goto L23
L23:
	;
	v69 = int32(*(*int8)(unsafe.Add(mBase, uint32(v52))))
	v75 = int32(1688224)
	v77 = int32(1689184)
	goto L25
L24:
	;
	goto L23
L25:
	;
	v84 = v75 + (v77-v75)>>(uint(int32(5))%32)<<(uint(int32(4))%32)
	v85 = int32(*(*int8)(unsafe.Add(mBase, uint32(v84))))
	v86 = v69 - v85
	if v86 == int32(0) {
		goto L27
	} else {
		goto L28
	}
L26:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v57))) = int32(0)
	v125 = int32(31)
	goto L19
L27:
	;
	v90 = F_strncmp(m, v52, v84, int32(10))
	mBase = m.M
	if v90 == int32(0) {
		v108 = v84
		goto L20
	} else {
		goto L30
	}
L28:
	;
	v93 = v86
	goto L29
L29:
	;
	v97 = base.B2i32(v93 < int32(0))
	if v93 < int32(0) {
		goto L31
	} else {
		goto L32
	}
L30:
	;
	v93 = v90
	goto L29
L31:
	;
	v98 = v84 - int32(16)
	goto L33
L32:
	;
	v98 = v77
	goto L33
L33:
	;
	if v93 < int32(0) {
		goto L34
	} else {
		goto L35
	}
L34:
	;
	v101 = v75
	goto L36
L35:
	;
	v101 = v84 + int32(16)
	goto L36
L36:
	;
	if base.Ui32(v101) <= base.Ui32(v98) {
		v75 = v101
		v77 = v98
		goto L25
	} else {
		goto L37
	}
L37:
	;
	goto L26
L38:
	;
	if base.Ui64(l1-int64(9223372036854775807)) <= base.Ui64(int64(1)) {
		goto L41
	} else {
		goto L42
	}
L39:
	;
	goto L40
L40:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v629 = m.ExcPending
	if v629 != 0 {
		goto L17
	} else {
		goto L142
	}
L41:
	;
	v132 = *(*int32)(unsafe.Add(mBase, uint32(v13)+104))
	if base.Ui32(v132-int32(18)) < base.Ui32(int32(13)) {
		v649 = l1
		goto L2
	} else {
		goto L44
	}
L42:
	;
	goto L43
L43:
	;
	v157 = int32(0)
	v165 = F_timestamp2tm(m, l1, v13+int32(108), v13+int32(56), v13+int32(100), v157, l2)
	mBase = m.M
	v166 = m.ExcPending
	if v166 != 0 {
		goto L17
	} else {
		goto L50
	}
L44:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v140 = m.ExcPending
	if v140 != 0 {
		goto L17
	} else {
		goto L45
	}
L45:
	;
	F_errcode(m, int32(1088))
	mBase = m.M
	v143 = m.ExcPending
	if v143 != 0 {
		goto L17
	} else {
		goto L46
	}
L46:
	;
	v145 = F_format_type_be(m, int32(1184))
	mBase = m.M
	v146 = m.ExcPending
	if v146 != 0 {
		goto L17
	} else {
		goto L47
	}
L47:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v13)+4)) = v145
	*(*int32)(unsafe.Add(mBase, uint32(v13))) = v52
	F_errmsg(m, int32(199023), v13)
	mBase = m.M
	v151 = m.ExcPending
	if v151 != 0 {
		goto L17
	} else {
		goto L48
	}
L48:
	;
	F_errfinish(m, int32(515856), int32(4964), int32(325052))
	mBase = m.M
	v156 = m.ExcPending
	if v156 != 0 {
		goto L17
	} else {
		goto L49
	}
L49:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L50:
	;
	if v165 != 0 {
		goto L1
	} else {
		goto L51
	}
L51:
	;
	v167 = int32(1)
	v168 = *(*int32)(unsafe.Add(mBase, uint32(v13)+104))
	switch v168 - int32(18) {
	case 0:
		v474 = v157
		goto L52
	case 1:
		goto L54
	case 2:
		goto L53
	case 3:
		goto L57
	case 4:
		goto L63
	case 5:
		goto L58
	case 6:
		goto L65
	case 7:
		v406 = v167
		goto L59
	case 8:
		goto L60
	case 9:
		goto L64
	case 10:
		goto L62
	case 11:
		goto L56
	case 12:
		goto L66
	default:
		goto L55
	}
L52:
	;
	v479 = *(*int32)(unsafe.Add(mBase, uint32(v13)+76))
	if v479 <= int32(-4713) {
		goto L117
	} else {
		goto L118
	}
L53:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v13)+56)) = int64(0)
	v474 = v157
	goto L52
L54:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v13)+56)) = int32(0)
	v474 = v157
	goto L52
L55:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v450 = m.ExcPending
	if v450 != 0 {
		goto L17
	} else {
		goto L110
	}
L56:
	;
	v443 = *(*int32)(unsafe.Add(mBase, uint32(v13)+100))
	v445 = base.I32_rem_s(v443, int32(1000))
	v474 = v443 - v445
	goto L52
L57:
	;
	v424 = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v13)+64)) = v424
	*(*int64)(unsafe.Add(mBase, uint32(v13)+56)) = int64(0)
	v432 = m.G0
	v433 = int32(16)
	v434 = v432 - v433
	m.G0 = v434
	v438 = F_DetermineTimeZoneOffsetInternal(m, v13+int32(56), l2, v434+int32(8))
	mBase = m.M
	m.G0 = v434 + v433
	goto L109
L58:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v13)+68)) = int32(1)
	goto L57
L59:
	;
	v411 = base.I32_rem_s(v406-int32(1), int32(3))
	*(*int32)(unsafe.Add(mBase, uint32(v13)+72)) = v406 - v411
	goto L58
L60:
	;
	v393 = *(*int32)(unsafe.Add(mBase, uint32(v13)+76))
	if int32(0) < v393 {
		goto L106
	} else {
		goto L107
	}
L61:
	;
	v376 = int32(1)
	if int32(0) < v375 {
		goto L103
	} else {
		goto L104
	}
L62:
	;
	v359 = *(*int32)(unsafe.Add(mBase, uint32(v13)+76))
	if int32(0) < v359 {
		goto L100
	} else {
		goto L101
	}
L63:
	;
	v174 = *(*int32)(unsafe.Add(mBase, uint32(v13)+76))
	v175 = *(*int32)(unsafe.Add(mBase, uint32(v13)+72))
	v176 = *(*int32)(unsafe.Add(mBase, uint32(v13)+68))
	v178 = F_date2j(m, v174, v175, v176)
	mBase = m.M
	v179 = int32(1)
	v181 = F_date2j(m, v174, v179, int32(4))
	mBase = m.M
	v184 = F_j2day(m, v181-v179)
	mBase = m.M
	if v178 < v181-v184 {
		goto L68
	} else {
		goto L69
	}
L64:
	;
	v173 = *(*int32)(unsafe.Add(mBase, uint32(v13)+76))
	v375 = v173
	goto L61
L65:
	;
	v172 = *(*int32)(unsafe.Add(mBase, uint32(v13)+72))
	v406 = v172
	goto L59
L66:
	;
	v171 = *(*int32)(unsafe.Add(mBase, uint32(v13)+100))
	v474 = v171
	goto L52
L67:
	;
	v225 = *(*int32)(unsafe.Add(mBase, uint32(v13)+72))
	if v218 < int32(52) {
		goto L79
	} else {
		goto L80
	}
L68:
	;
	v187 = int32(1)
	v191 = F_date2j(m, v174-v187, v187, int32(4))
	mBase = m.M
	v194 = F_j2day(m, v191-v187)
	mBase = m.M
	v195 = v191
	v196 = v194
	goto L70
L69:
	;
	v195 = v181
	v196 = v184
	goto L70
L70:
	;
	v198 = v196 - v195 + v178
	if int32(357) <= v198 {
		goto L71
	} else {
		goto L72
	}
L71:
	;
	v201 = int32(1)
	v205 = F_date2j(m, v174+v201, v201, int32(4))
	mBase = m.M
	v208 = F_j2day(m, v205-v201)
	mBase = m.M
	v209 = v205 - v208
	if v178 < v209 {
		goto L74
	} else {
		goto L75
	}
L72:
	;
	v214 = v198
	goto L73
L73:
	;
	v216 = base.I32_div_s(v214, int32(7))
	v218 = v216 + int32(1)
	goto L67
L74:
	;
	v212 = v198
	goto L76
L75:
	;
	v212 = v178 - v209
	goto L76
L76:
	;
	v214 = v212
	goto L73
L77:
	;
	goto L86
L78:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v13)+76)) = v241
	v243 = v241
	goto L77
L79:
	;
	v233 = *(*int32)(unsafe.Add(mBase, uint32(v13)+76))
	if int32(1) < v218 {
		v243 = v233
		goto L77
	} else {
		goto L82
	}
L80:
	;
	if v225 != int32(1) {
		goto L79
	} else {
		goto L81
	}
L81:
	;
	v230 = *(*int32)(unsafe.Add(mBase, uint32(v13)+76))
	v241 = v230 - int32(1)
	goto L78
L82:
	;
	if v225 != int32(12) {
		v243 = v233
		goto L77
	} else {
		goto L83
	}
L83:
	;
	v241 = v233 + int32(1)
	goto L78
L84:
	;
	v279 = int32(1)
	v283 = int32(7)
	v284 = base.I32_rem_s(v277-v279+v279, v283)
	if v284 < int32(0) {
		goto L92
	} else {
		goto L93
	}
L86:
	;
	goto L87
L87:
	;
	v254 = int32(4799) + v243
	v259 = base.I32_div_s(v254, int32(4))
	v262 = base.I32_div_s(v254, int32(-100))
	v265 = base.I32_div_s(v254, int32(400))
	goto L89
L89:
	;
	goto L90
L90:
	;
	v274 = base.I32_div_s(int32(109676), int32(256))
	v277 = int32(4) + v254*int32(365) + v259 + v262 + v265 + v274 - int32(32167)
	goto L84
L91:
	;
	v292 = v218*int32(7) + v277 - v289 - int32(7)
	v296 = v292 + int32(32044)
	v297 = int32(146097)
	v298 = base.I32_div_u_s(v296, v297)
	v299 = int32(3)
	v305 = int32(2)
	v310 = base.I32_div_u_s((v298*int32(1073595727)+v296)<<(uint(v305)%32)|v299, v297)
	v313 = v292 + v298*v299 + v310 + int32(32104)
	v314 = int32(1461)
	v315 = base.I32_div_u_s(v313, v314)
	v318 = v315*int32(-1461) + v313
	v320 = v318 << (uint(v305) % 32)
	if base.Ui32(v314) <= base.Ui32(v320) {
		goto L97
	} else {
		goto L98
	}
L92:
	;
	v289 = v284 + v283
	goto L94
L93:
	;
	v289 = v284
	goto L94
L94:
	;
	goto L91
L95:
	;
	goto L57
L96:
	;
	v333 = base.I32_div_u_s(v320, int32(1461))
	*(*int32)(unsafe.Add(mBase, uint32(v13+int32(76)))) = v333 + v315<<(uint(int32(2))%32) - int32(4800)
	v341 = v331 + int32(123)
	v345 = int32(base.Ui32(v341*int32(2141)) >> (uint(int32(16)) % 32))
	*(*int32)(unsafe.Add(mBase, uint32(v13+int32(68)))) = v341 - int32(base.Ui32(v345*int32(7834))>>(uint(int32(8))%32))
	v355 = base.I32_rem_u_s(v345+int32(10), int32(12))
	*(*int32)(unsafe.Add(mBase, uint32(v13+int32(72)))) = v355 + int32(1)
	goto L95
L97:
	;
	v326 = base.I32_rem_u_s(v318+int32(305), int32(365))
	v331 = v326
	goto L96
L98:
	;
	goto L99
L99:
	;
	v330 = base.I32_rem_u_s(v318+int32(306), int32(366))
	v331 = v330
	goto L96
L100:
	;
	v365 = base.I32_rem_s(v359+int32(999), int32(1000))
	v375 = v359 - v365
	goto L61
L101:
	;
	goto L102
L102:
	;
	v367 = int32(1000)
	v370 = base.I32_rem_s(v367-v359, v367)
	v375 = v359 + v370 - int32(999)
	goto L61
L103:
	;
	v382 = base.I32_rem_s(v375+int32(99), int32(100))
	*(*int32)(unsafe.Add(mBase, uint32(v13)+76)) = v375 - v382
	v406 = v376
	goto L59
L104:
	;
	goto L105
L105:
	;
	v385 = int32(100)
	v388 = base.I32_rem_s(v385-v375, v385)
	*(*int32)(unsafe.Add(mBase, uint32(v13)+76)) = v375 + v388 - int32(99)
	v406 = v376
	goto L59
L106:
	;
	v397 = base.I32_rem_u_s(v393, int32(10))
	*(*int32)(unsafe.Add(mBase, uint32(v13)+76)) = v393 - v397
	v406 = v167
	goto L59
L107:
	;
	goto L108
L108:
	;
	v401 = int32(9) - v393
	v403 = base.I32_rem_s(v401, int32(10))
	*(*int32)(unsafe.Add(mBase, uint32(v13)+76)) = v403 - v401
	v406 = v167
	goto L59
L109:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v13)+108)) = v438
	v474 = v424
	goto L52
L110:
	;
	F_errcode(m, int32(1088))
	mBase = m.M
	v453 = m.ExcPending
	if v453 != 0 {
		goto L17
	} else {
		goto L111
	}
L111:
	;
	v455 = F_format_type_be(m, int32(1184))
	mBase = m.M
	v456 = m.ExcPending
	if v456 != 0 {
		goto L17
	} else {
		goto L112
	}
L112:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v13)+20)) = v455
	*(*int32)(unsafe.Add(mBase, uint32(v13)+16)) = v52
	F_errmsg(m, int32(199023), v13+int32(16))
	mBase = m.M
	v463 = m.ExcPending
	if v463 != 0 {
		goto L17
	} else {
		goto L113
	}
L113:
	;
	F_errfinish(m, int32(515856), int32(5065), int32(325052))
	mBase = m.M
	v468 = m.ExcPending
	if v468 != 0 {
		goto L17
	} else {
		goto L114
	}
L114:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L115:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v613 = m.ExcPending
	if v613 != 0 {
		goto L17
	} else {
		goto L138
	}
L116:
	;
	v497 = v13 + int32(32)
	v498 = *(*int32)(unsafe.Add(mBase, uint32(v13)+68))
	v503 = base.B2i32(int32(2) < v495)
	if int32(2) < v495 {
		goto L128
	} else {
		goto L129
	}
L117:
	;
	if v479 != int32(-4713) {
		goto L115
	} else {
		goto L120
	}
L118:
	;
	goto L119
L119:
	;
	if v479 <= int32(5874897) {
		goto L122
	} else {
		goto L123
	}
L120:
	;
	v484 = *(*int32)(unsafe.Add(mBase, uint32(v13)+72))
	if int32(10) < v484 {
		v495 = v484
		goto L116
	} else {
		goto L121
	}
L121:
	;
	goto L115
L122:
	;
	v489 = *(*int32)(unsafe.Add(mBase, uint32(v13)+72))
	v495 = v489
	goto L116
L123:
	;
	goto L124
L124:
	;
	if v479 != int32(5874898) {
		goto L115
	} else {
		goto L125
	}
L125:
	;
	v492 = *(*int32)(unsafe.Add(mBase, uint32(v13)+72))
	if int32(5) < v492 {
		goto L115
	} else {
		goto L126
	}
L126:
	;
	v495 = v492
	goto L116
L127:
	;
	v531 = base.I64_extend_i32_s(v498 + v505*int32(365) + v510 + v513 + v516 + v525 - int32(32167) - int32(2451545))
	v540 = int64(32)
	v541 = int64(20)
	v543 = int64(base.Ui64(v531) >> (uint(v540) % 64))
	v546 = int64(4294967295)
	v547 = int64(500654080)
	v549 = v531 & v546
	v550 = v547 * v549
	v554 = int64(base.Ui64(v550)>>(uint(v540)%64)) + v547*v543
	v561 = v549*v541 + v554&v546
	*(*int64)(unsafe.Add(mBase, uint32(v497)+8)) = v531*int64(0) + v531>>(uint(int64(63))%64)*int64(86400000000) + v541*v543 + int64(base.Ui64(v554)>>(uint(v540)%64)) + int64(base.Ui64(v561)>>(uint(v540)%64))
	*(*int64)(unsafe.Add(mBase, uint32(v497))) = v550&v546 | v561<<(uint(v540)%64)
	goto L134
L128:
	;
	v504 = int32(4800)
	goto L130
L129:
	;
	v504 = int32(4799)
	goto L130
L130:
	;
	v505 = v504 + v479
	v510 = base.I32_div_s(v505, int32(4))
	v513 = base.I32_div_s(v505, int32(-100))
	v516 = base.I32_div_s(v505, int32(400))
	if int32(2) < v495 {
		goto L131
	} else {
		goto L132
	}
L131:
	;
	v520 = int32(1)
	goto L133
L132:
	;
	v520 = int32(13)
	goto L133
L133:
	;
	v525 = base.I32_div_s((v520+v495)*int32(7834), int32(256))
	goto L127
L134:
	;
	v572 = *(*int64)(unsafe.Add(mBase, uint32(v13)+40))
	v573 = *(*int64)(unsafe.Add(mBase, uint32(v13)+32))
	if v572 != v573>>(uint(int64(63))%64) {
		goto L115
	} else {
		goto L135
	}
L135:
	;
	v578 = *(*int32)(unsafe.Add(mBase, uint32(v13)+56))
	v579 = *(*int32)(unsafe.Add(mBase, uint32(v13)+60))
	v580 = *(*int32)(unsafe.Add(mBase, uint32(v13)+64))
	v581 = int32(60)
	v590 = base.I64_extend_i32_s(v474) + base.I64_extend_i32_s(v578+(v579+v580*v581)*v581)*int64(1000000)
	v593 = v573 + v590
	if base.B2i32(v590 < int64(0))^base.B2i32(v593 < v573) != 0 {
		goto L115
	} else {
		goto L136
	}
L136:
	;
	v597 = *(*int32)(unsafe.Add(mBase, uint32(v13)+108))
	v602 = base.I64_extend_i32_s(int32(0)-v597)*int64(-1000000) + v593
	if base.Ui64(v602+int64(211813488000000000)) < base.Ui64(int64(-9011559254509551616)) {
		v649 = v602
		goto L2
	} else {
		goto L137
	}
L137:
	;
	goto L115
L138:
	;
	F_errcode(m, int32(134217858))
	mBase = m.M
	v616 = m.ExcPending
	if v616 != 0 {
		goto L17
	} else {
		goto L139
	}
L139:
	;
	F_errmsg(m, int32(418494), int32(0))
	mBase = m.M
	v620 = m.ExcPending
	if v620 != 0 {
		goto L17
	} else {
		goto L140
	}
L140:
	;
	F_errfinish(m, int32(515856), int32(5075), int32(325052))
	mBase = m.M
	v625 = m.ExcPending
	if v625 != 0 {
		goto L17
	} else {
		goto L141
	}
L141:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L142:
	;
	F_errcode(m, int32(50856066))
	mBase = m.M
	v632 = m.ExcPending
	if v632 != 0 {
		goto L17
	} else {
		goto L143
	}
L143:
	;
	v634 = F_format_type_be(m, int32(1184))
	mBase = m.M
	v635 = m.ExcPending
	if v635 != 0 {
		goto L17
	} else {
		goto L144
	}
L144:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v13)+52)) = v634
	*(*int32)(unsafe.Add(mBase, uint32(v13)+48)) = v52
	F_errmsg(m, int32(198986), v13+int32(48))
	mBase = m.M
	v642 = m.ExcPending
	if v642 != 0 {
		goto L17
	} else {
		goto L145
	}
L145:
	;
	F_errfinish(m, int32(515856), int32(5082), int32(325052))
	mBase = m.M
	v647 = m.ExcPending
	if v647 != 0 {
		goto L17
	} else {
		goto L146
	}
L146:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L147:
	;
	F_errcode(m, int32(134217858))
	mBase = m.M
	v666 = m.ExcPending
	if v666 != 0 {
		goto L17
	} else {
		goto L148
	}
L148:
	;
	F_errmsg(m, int32(418494), int32(0))
	mBase = m.M
	v670 = m.ExcPending
	if v670 != 0 {
		goto L17
	} else {
		goto L149
	}
L149:
	;
	F_errfinish(m, int32(515856), int32(4972), int32(325052))
	mBase = m.M
	v675 = m.ExcPending
	if v675 != 0 {
		goto L17
	} else {
		goto L150
	}
L150:
	;
	base.Wasm_trap_unreachable()
	for {
	}
}
func F_timestamptz_zone(m *base.Module, l0 int32) int32 {
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
	var v14 int32
	_ = v14
	var v15 int32
	_ = v15
	var v16 int64
	_ = v16
	var v21 int32
	_ = v21
	var v22 int32
	_ = v22
	var v27 int32
	_ = v27
	var v34 int32
	_ = v34
	var v35 int32
	_ = v35
	var v37 int32
	_ = v37
	var v38 int32
	_ = v38
	var v46 int32
	_ = v46
	var v49 int32
	_ = v49
	var v50 int32
	_ = v50
	var v63 int32
	_ = v63
	var v64 int32
	_ = v64
	var v65 int32
	_ = v65
	var v66 int64
	_ = v66
	var v67 int32
	_ = v67
	var v72 int32
	_ = v72
	var v77 int32
	_ = v77
	var v80 int32
	_ = v80
	var v83 int32
	_ = v83
	var v85 int32
	_ = v85
	var v86 int32
	_ = v86
	var v91 int32
	_ = v91
	var v92 int32
	_ = v92
	var v93 int32
	_ = v93
	var v98 int32
	_ = v98
	var v101 int32
	_ = v101
	var v104 int32
	_ = v104
	var v108 int32
	_ = v108
	var v113 int32
	_ = v113
	var v119 int64
	_ = v119
	var v128 int64
	_ = v128
	var v129 int64
	_ = v129
	var v131 int64
	_ = v131
	var v134 int64
	_ = v134
	var v135 int64
	_ = v135
	var v137 int64
	_ = v137
	var v138 int64
	_ = v138
	var v142 int64
	_ = v142
	var v149 int64
	_ = v149
	var v160 int64
	_ = v160
	var v161 int64
	_ = v161
	var v165 int32
	_ = v165
	var v166 int32
	_ = v166
	var v167 int32
	_ = v167
	var v168 int32
	_ = v168
	var v177 int64
	_ = v177
	var v180 int64
	_ = v180
	var v189 int64
	_ = v189
	var v195 int32
	_ = v195
	var v196 int32
	_ = v196
	var v201 int32
	_ = v201
	var v209 int32
	_ = v209
	var v212 int32
	_ = v212
	var v216 int32
	_ = v216
	var v221 int32
	_ = v221
	var v225 int32
	_ = v225
	var v228 int32
	_ = v228
	var v232 int32
	_ = v232
	var v237 int32
	_ = v237
	var v244 int32
	_ = v244
	var v247 int32
	_ = v247
	var v251 int32
	_ = v251
	var v256 int32
	_ = v256
	v6 = m.G0
	v8 = v6 - int32(352)
	m.G0 = v8
	v10 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v11 = F_pg_detoast_datum_packed(m, v10)
	mBase = m.M
	v14 = m.ExcPending
	if v14 != 0 {
		return int32(0)
	} else {
		v15 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
		v16 = *(*int64)(unsafe.Add(mBase, uint32(v15)))
		if base.Ui64(v16-int64(9223372036854775807)) <= base.Ui64(int64(1)) {
			v21 = F_Int64GetDatum(m, v16)
			mBase = m.M
			v22 = m.ExcPending
			if v22 != 0 {
				return int32(0)
			} else {
				v201 = v21
				m.G0 = v8 + int32(352)
				return v201
			}
		} else {
			F_text_to_cstring_buffer(m, v11, v8+int32(80), int32(256))
			mBase = m.M
			v27 = m.ExcPending
			if v27 != 0 {
				return int32(0)
			} else {
				v34 = F_DecodeTimezoneName(m, v8+int32(80), v8+int32(76), v8+int32(72))
				mBase = m.M
				v35 = m.ExcPending
				if v35 != 0 {
					return int32(0)
				} else {
					switch v34 {
					case 0:
						v37 = *(*int32)(unsafe.Add(mBase, uint32(v8)+76))
						v38 = int32(0) - v37
						*(*int32)(unsafe.Add(mBase, uint32(v8)+348)) = v38
						v189 = base.I64_extend_i32_s(v38)*int64(-1000000) + v16
						if base.Ui64(int64(-9011559254509551616)) <= base.Ui64(v189+int64(211813488000000000)) {
							F_errstart_cold(m, int32(21), int32(0))
							mBase = m.M
							v225 = m.ExcPending
							if v225 != 0 {
								return int32(0)
							} else {
								F_errcode(m, int32(134217858))
								mBase = m.M
								v228 = m.ExcPending
								if v228 != 0 {
									return int32(0)
								} else {
									F_errmsg(m, int32(418494), int32(0))
									mBase = m.M
									v232 = m.ExcPending
									if v232 != 0 {
										return int32(0)
									} else {
										F_errfinish(m, int32(515856), int32(6618), int32(387606))
										mBase = m.M
										v237 = m.ExcPending
										if v237 != 0 {
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
							v195 = F_Int64GetDatum(m, v189)
							mBase = m.M
							v196 = m.ExcPending
							if v196 != 0 {
								return int32(0)
							} else {
								v201 = v195
								m.G0 = v8 + int32(352)
								return v201
							}
						}
					case 1:
						v46 = *(*int32)(unsafe.Add(mBase, uint32(v8)+72))
						v49 = F_DetermineTimeZoneAbbrevOffsetTS(m, v16, v8+int32(80), v46, v8+int32(28))
						mBase = m.M
						v50 = m.ExcPending
						if v50 != 0 {
							return int32(0)
						} else {
							*(*int32)(unsafe.Add(mBase, uint32(v8)+348)) = v49
							v189 = base.I64_extend_i32_s(v49)*int64(-1000000) + v16
							if base.Ui64(int64(-9011559254509551616)) <= base.Ui64(v189+int64(211813488000000000)) {
								F_errstart_cold(m, int32(21), int32(0))
								mBase = m.M
								v225 = m.ExcPending
								if v225 != 0 {
									return int32(0)
								} else {
									F_errcode(m, int32(134217858))
									mBase = m.M
									v228 = m.ExcPending
									if v228 != 0 {
										return int32(0)
									} else {
										F_errmsg(m, int32(418494), int32(0))
										mBase = m.M
										v232 = m.ExcPending
										if v232 != 0 {
											return int32(0)
										} else {
											F_errfinish(m, int32(515856), int32(6618), int32(387606))
											mBase = m.M
											v237 = m.ExcPending
											if v237 != 0 {
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
								v195 = F_Int64GetDatum(m, v189)
								mBase = m.M
								v196 = m.ExcPending
								if v196 != 0 {
									return int32(0)
								} else {
									v201 = v195
									m.G0 = v8 + int32(352)
									return v201
								}
							}
						}
					default:
						v63 = *(*int32)(unsafe.Add(mBase, uint32(v8)+72))
						v64 = F_timestamp2tm(m, v16, v8+int32(348), v8+int32(28), v8+int32(24), int32(0), v63)
						mBase = m.M
						v65 = m.ExcPending
						if v65 != 0 {
							return int32(0)
						} else {
							if v64 != 0 {
								F_errstart_cold(m, int32(21), int32(0))
								mBase = m.M
								v209 = m.ExcPending
								if v209 != 0 {
									return int32(0)
								} else {
									F_errcode(m, int32(134217858))
									mBase = m.M
									v212 = m.ExcPending
									if v212 != 0 {
										return int32(0)
									} else {
										F_errmsg(m, int32(418494), int32(0))
										mBase = m.M
										v216 = m.ExcPending
										if v216 != 0 {
											return int32(0)
										} else {
											F_errfinish(m, int32(515856), int32(6608), int32(387606))
											mBase = m.M
											v221 = m.ExcPending
											if v221 != 0 {
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
								v66 = int64(*(*int32)(unsafe.Add(mBase, uint32(v8)+24)))
								v67 = *(*int32)(unsafe.Add(mBase, uint32(v8)+48))
								if v67 <= int32(-4713) {
									if v67 != int32(-4713) {
										F_errstart_cold(m, int32(21), int32(0))
										mBase = m.M
										v244 = m.ExcPending
										if v244 != 0 {
											return int32(0)
										} else {
											F_errcode(m, int32(134217858))
											mBase = m.M
											v247 = m.ExcPending
											if v247 != 0 {
												return int32(0)
											} else {
												F_errmsg(m, int32(418494), int32(0))
												mBase = m.M
												v251 = m.ExcPending
												if v251 != 0 {
													return int32(0)
												} else {
													F_errfinish(m, int32(515856), int32(6612), int32(387606))
													mBase = m.M
													v256 = m.ExcPending
													if v256 != 0 {
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
										v72 = *(*int32)(unsafe.Add(mBase, uint32(v8)+44))
										if int32(10) < v72 {
											v83 = v72
											v85 = v8 + int32(8)
											v86 = *(*int32)(unsafe.Add(mBase, uint32(v8)+40))
											v91 = base.B2i32(int32(2) < v83)
											if int32(2) < v83 {
												v92 = int32(4800)
											} else {
												v92 = int32(4799)
											}
											v93 = v92 + v67
											v98 = base.I32_div_s(v93, int32(4))
											v101 = base.I32_div_s(v93, int32(-100))
											v104 = base.I32_div_s(v93, int32(400))
											if int32(2) < v83 {
												v108 = int32(1)
											} else {
												v108 = int32(13)
											}
											v113 = base.I32_div_s((v108+v83)*int32(7834), int32(256))
											v119 = base.I64_extend_i32_s(v86 + v93*int32(365) + v98 + v101 + v104 + v113 - int32(32167) - int32(2451545))
											v128 = int64(32)
											v129 = int64(20)
											v131 = int64(base.Ui64(v119) >> (uint(v128) % 64))
											v134 = int64(4294967295)
											v135 = int64(500654080)
											v137 = v119 & v134
											v138 = v135 * v137
											v142 = int64(base.Ui64(v138)>>(uint(v128)%64)) + v135*v131
											v149 = v137*v129 + v142&v134
											*(*int64)(unsafe.Add(mBase, uint32(v85)+8)) = v119*int64(0) + v119>>(uint(int64(63))%64)*int64(86400000000) + v129*v131 + int64(base.Ui64(v142)>>(uint(v128)%64)) + int64(base.Ui64(v149)>>(uint(v128)%64))
											*(*int64)(unsafe.Add(mBase, uint32(v85))) = v138&v134 | v149<<(uint(v128)%64)
											v160 = *(*int64)(unsafe.Add(mBase, uint32(v8)+16))
											v161 = *(*int64)(unsafe.Add(mBase, uint32(v8)+8))
											if v160 != v161>>(uint(int64(63))%64) {
												F_errstart_cold(m, int32(21), int32(0))
												mBase = m.M
												v244 = m.ExcPending
												if v244 != 0 {
													return int32(0)
												} else {
													F_errcode(m, int32(134217858))
													mBase = m.M
													v247 = m.ExcPending
													if v247 != 0 {
														return int32(0)
													} else {
														F_errmsg(m, int32(418494), int32(0))
														mBase = m.M
														v251 = m.ExcPending
														if v251 != 0 {
															return int32(0)
														} else {
															F_errfinish(m, int32(515856), int32(6612), int32(387606))
															mBase = m.M
															v256 = m.ExcPending
															if v256 != 0 {
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
												v165 = *(*int32)(unsafe.Add(mBase, uint32(v8)+28))
												v166 = *(*int32)(unsafe.Add(mBase, uint32(v8)+32))
												v167 = *(*int32)(unsafe.Add(mBase, uint32(v8)+36))
												v168 = int32(60)
												v177 = base.I64_extend_i32_s(v165+(v166+v167*v168)*v168)*int64(1000000) + v66
												v180 = v177 + v161
												if base.B2i32(v177 < int64(0))^base.B2i32(v180 < v161) != 0 {
													F_errstart_cold(m, int32(21), int32(0))
													mBase = m.M
													v244 = m.ExcPending
													if v244 != 0 {
														return int32(0)
													} else {
														F_errcode(m, int32(134217858))
														mBase = m.M
														v247 = m.ExcPending
														if v247 != 0 {
															return int32(0)
														} else {
															F_errmsg(m, int32(418494), int32(0))
															mBase = m.M
															v251 = m.ExcPending
															if v251 != 0 {
																return int32(0)
															} else {
																F_errfinish(m, int32(515856), int32(6612), int32(387606))
																mBase = m.M
																v256 = m.ExcPending
																if v256 != 0 {
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
													if base.Ui64(v180-int64(9223371331200000000)) <= base.Ui64(int64(9011559254509551615)) {
														F_errstart_cold(m, int32(21), int32(0))
														mBase = m.M
														v244 = m.ExcPending
														if v244 != 0 {
															return int32(0)
														} else {
															F_errcode(m, int32(134217858))
															mBase = m.M
															v247 = m.ExcPending
															if v247 != 0 {
																return int32(0)
															} else {
																F_errmsg(m, int32(418494), int32(0))
																mBase = m.M
																v251 = m.ExcPending
																if v251 != 0 {
																	return int32(0)
																} else {
																	F_errfinish(m, int32(515856), int32(6612), int32(387606))
																	mBase = m.M
																	v256 = m.ExcPending
																	if v256 != 0 {
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
														v189 = v180
														if base.Ui64(int64(-9011559254509551616)) <= base.Ui64(v189+int64(211813488000000000)) {
															F_errstart_cold(m, int32(21), int32(0))
															mBase = m.M
															v225 = m.ExcPending
															if v225 != 0 {
																return int32(0)
															} else {
																F_errcode(m, int32(134217858))
																mBase = m.M
																v228 = m.ExcPending
																if v228 != 0 {
																	return int32(0)
																} else {
																	F_errmsg(m, int32(418494), int32(0))
																	mBase = m.M
																	v232 = m.ExcPending
																	if v232 != 0 {
																		return int32(0)
																	} else {
																		F_errfinish(m, int32(515856), int32(6618), int32(387606))
																		mBase = m.M
																		v237 = m.ExcPending
																		if v237 != 0 {
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
															v195 = F_Int64GetDatum(m, v189)
															mBase = m.M
															v196 = m.ExcPending
															if v196 != 0 {
																return int32(0)
															} else {
																v201 = v195
																m.G0 = v8 + int32(352)
																return v201
															}
														}
													}
												}
											}
										} else {
											F_errstart_cold(m, int32(21), int32(0))
											mBase = m.M
											v244 = m.ExcPending
											if v244 != 0 {
												return int32(0)
											} else {
												F_errcode(m, int32(134217858))
												mBase = m.M
												v247 = m.ExcPending
												if v247 != 0 {
													return int32(0)
												} else {
													F_errmsg(m, int32(418494), int32(0))
													mBase = m.M
													v251 = m.ExcPending
													if v251 != 0 {
														return int32(0)
													} else {
														F_errfinish(m, int32(515856), int32(6612), int32(387606))
														mBase = m.M
														v256 = m.ExcPending
														if v256 != 0 {
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
									if v67 <= int32(5874897) {
										v77 = *(*int32)(unsafe.Add(mBase, uint32(v8)+44))
										v83 = v77
										v85 = v8 + int32(8)
										v86 = *(*int32)(unsafe.Add(mBase, uint32(v8)+40))
										v91 = base.B2i32(int32(2) < v83)
										if int32(2) < v83 {
											v92 = int32(4800)
										} else {
											v92 = int32(4799)
										}
										v93 = v92 + v67
										v98 = base.I32_div_s(v93, int32(4))
										v101 = base.I32_div_s(v93, int32(-100))
										v104 = base.I32_div_s(v93, int32(400))
										if int32(2) < v83 {
											v108 = int32(1)
										} else {
											v108 = int32(13)
										}
										v113 = base.I32_div_s((v108+v83)*int32(7834), int32(256))
										v119 = base.I64_extend_i32_s(v86 + v93*int32(365) + v98 + v101 + v104 + v113 - int32(32167) - int32(2451545))
										v128 = int64(32)
										v129 = int64(20)
										v131 = int64(base.Ui64(v119) >> (uint(v128) % 64))
										v134 = int64(4294967295)
										v135 = int64(500654080)
										v137 = v119 & v134
										v138 = v135 * v137
										v142 = int64(base.Ui64(v138)>>(uint(v128)%64)) + v135*v131
										v149 = v137*v129 + v142&v134
										*(*int64)(unsafe.Add(mBase, uint32(v85)+8)) = v119*int64(0) + v119>>(uint(int64(63))%64)*int64(86400000000) + v129*v131 + int64(base.Ui64(v142)>>(uint(v128)%64)) + int64(base.Ui64(v149)>>(uint(v128)%64))
										*(*int64)(unsafe.Add(mBase, uint32(v85))) = v138&v134 | v149<<(uint(v128)%64)
										v160 = *(*int64)(unsafe.Add(mBase, uint32(v8)+16))
										v161 = *(*int64)(unsafe.Add(mBase, uint32(v8)+8))
										if v160 != v161>>(uint(int64(63))%64) {
											F_errstart_cold(m, int32(21), int32(0))
											mBase = m.M
											v244 = m.ExcPending
											if v244 != 0 {
												return int32(0)
											} else {
												F_errcode(m, int32(134217858))
												mBase = m.M
												v247 = m.ExcPending
												if v247 != 0 {
													return int32(0)
												} else {
													F_errmsg(m, int32(418494), int32(0))
													mBase = m.M
													v251 = m.ExcPending
													if v251 != 0 {
														return int32(0)
													} else {
														F_errfinish(m, int32(515856), int32(6612), int32(387606))
														mBase = m.M
														v256 = m.ExcPending
														if v256 != 0 {
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
											v165 = *(*int32)(unsafe.Add(mBase, uint32(v8)+28))
											v166 = *(*int32)(unsafe.Add(mBase, uint32(v8)+32))
											v167 = *(*int32)(unsafe.Add(mBase, uint32(v8)+36))
											v168 = int32(60)
											v177 = base.I64_extend_i32_s(v165+(v166+v167*v168)*v168)*int64(1000000) + v66
											v180 = v177 + v161
											if base.B2i32(v177 < int64(0))^base.B2i32(v180 < v161) != 0 {
												F_errstart_cold(m, int32(21), int32(0))
												mBase = m.M
												v244 = m.ExcPending
												if v244 != 0 {
													return int32(0)
												} else {
													F_errcode(m, int32(134217858))
													mBase = m.M
													v247 = m.ExcPending
													if v247 != 0 {
														return int32(0)
													} else {
														F_errmsg(m, int32(418494), int32(0))
														mBase = m.M
														v251 = m.ExcPending
														if v251 != 0 {
															return int32(0)
														} else {
															F_errfinish(m, int32(515856), int32(6612), int32(387606))
															mBase = m.M
															v256 = m.ExcPending
															if v256 != 0 {
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
												if base.Ui64(v180-int64(9223371331200000000)) <= base.Ui64(int64(9011559254509551615)) {
													F_errstart_cold(m, int32(21), int32(0))
													mBase = m.M
													v244 = m.ExcPending
													if v244 != 0 {
														return int32(0)
													} else {
														F_errcode(m, int32(134217858))
														mBase = m.M
														v247 = m.ExcPending
														if v247 != 0 {
															return int32(0)
														} else {
															F_errmsg(m, int32(418494), int32(0))
															mBase = m.M
															v251 = m.ExcPending
															if v251 != 0 {
																return int32(0)
															} else {
																F_errfinish(m, int32(515856), int32(6612), int32(387606))
																mBase = m.M
																v256 = m.ExcPending
																if v256 != 0 {
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
													v189 = v180
													if base.Ui64(int64(-9011559254509551616)) <= base.Ui64(v189+int64(211813488000000000)) {
														F_errstart_cold(m, int32(21), int32(0))
														mBase = m.M
														v225 = m.ExcPending
														if v225 != 0 {
															return int32(0)
														} else {
															F_errcode(m, int32(134217858))
															mBase = m.M
															v228 = m.ExcPending
															if v228 != 0 {
																return int32(0)
															} else {
																F_errmsg(m, int32(418494), int32(0))
																mBase = m.M
																v232 = m.ExcPending
																if v232 != 0 {
																	return int32(0)
																} else {
																	F_errfinish(m, int32(515856), int32(6618), int32(387606))
																	mBase = m.M
																	v237 = m.ExcPending
																	if v237 != 0 {
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
														v195 = F_Int64GetDatum(m, v189)
														mBase = m.M
														v196 = m.ExcPending
														if v196 != 0 {
															return int32(0)
														} else {
															v201 = v195
															m.G0 = v8 + int32(352)
															return v201
														}
													}
												}
											}
										}
									} else {
										if v67 != int32(5874898) {
											F_errstart_cold(m, int32(21), int32(0))
											mBase = m.M
											v244 = m.ExcPending
											if v244 != 0 {
												return int32(0)
											} else {
												F_errcode(m, int32(134217858))
												mBase = m.M
												v247 = m.ExcPending
												if v247 != 0 {
													return int32(0)
												} else {
													F_errmsg(m, int32(418494), int32(0))
													mBase = m.M
													v251 = m.ExcPending
													if v251 != 0 {
														return int32(0)
													} else {
														F_errfinish(m, int32(515856), int32(6612), int32(387606))
														mBase = m.M
														v256 = m.ExcPending
														if v256 != 0 {
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
											v80 = *(*int32)(unsafe.Add(mBase, uint32(v8)+44))
											if int32(5) < v80 {
												F_errstart_cold(m, int32(21), int32(0))
												mBase = m.M
												v244 = m.ExcPending
												if v244 != 0 {
													return int32(0)
												} else {
													F_errcode(m, int32(134217858))
													mBase = m.M
													v247 = m.ExcPending
													if v247 != 0 {
														return int32(0)
													} else {
														F_errmsg(m, int32(418494), int32(0))
														mBase = m.M
														v251 = m.ExcPending
														if v251 != 0 {
															return int32(0)
														} else {
															F_errfinish(m, int32(515856), int32(6612), int32(387606))
															mBase = m.M
															v256 = m.ExcPending
															if v256 != 0 {
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
												v83 = v80
												v85 = v8 + int32(8)
												v86 = *(*int32)(unsafe.Add(mBase, uint32(v8)+40))
												v91 = base.B2i32(int32(2) < v83)
												if int32(2) < v83 {
													v92 = int32(4800)
												} else {
													v92 = int32(4799)
												}
												v93 = v92 + v67
												v98 = base.I32_div_s(v93, int32(4))
												v101 = base.I32_div_s(v93, int32(-100))
												v104 = base.I32_div_s(v93, int32(400))
												if int32(2) < v83 {
													v108 = int32(1)
												} else {
													v108 = int32(13)
												}
												v113 = base.I32_div_s((v108+v83)*int32(7834), int32(256))
												v119 = base.I64_extend_i32_s(v86 + v93*int32(365) + v98 + v101 + v104 + v113 - int32(32167) - int32(2451545))
												v128 = int64(32)
												v129 = int64(20)
												v131 = int64(base.Ui64(v119) >> (uint(v128) % 64))
												v134 = int64(4294967295)
												v135 = int64(500654080)
												v137 = v119 & v134
												v138 = v135 * v137
												v142 = int64(base.Ui64(v138)>>(uint(v128)%64)) + v135*v131
												v149 = v137*v129 + v142&v134
												*(*int64)(unsafe.Add(mBase, uint32(v85)+8)) = v119*int64(0) + v119>>(uint(int64(63))%64)*int64(86400000000) + v129*v131 + int64(base.Ui64(v142)>>(uint(v128)%64)) + int64(base.Ui64(v149)>>(uint(v128)%64))
												*(*int64)(unsafe.Add(mBase, uint32(v85))) = v138&v134 | v149<<(uint(v128)%64)
												v160 = *(*int64)(unsafe.Add(mBase, uint32(v8)+16))
												v161 = *(*int64)(unsafe.Add(mBase, uint32(v8)+8))
												if v160 != v161>>(uint(int64(63))%64) {
													F_errstart_cold(m, int32(21), int32(0))
													mBase = m.M
													v244 = m.ExcPending
													if v244 != 0 {
														return int32(0)
													} else {
														F_errcode(m, int32(134217858))
														mBase = m.M
														v247 = m.ExcPending
														if v247 != 0 {
															return int32(0)
														} else {
															F_errmsg(m, int32(418494), int32(0))
															mBase = m.M
															v251 = m.ExcPending
															if v251 != 0 {
																return int32(0)
															} else {
																F_errfinish(m, int32(515856), int32(6612), int32(387606))
																mBase = m.M
																v256 = m.ExcPending
																if v256 != 0 {
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
													v165 = *(*int32)(unsafe.Add(mBase, uint32(v8)+28))
													v166 = *(*int32)(unsafe.Add(mBase, uint32(v8)+32))
													v167 = *(*int32)(unsafe.Add(mBase, uint32(v8)+36))
													v168 = int32(60)
													v177 = base.I64_extend_i32_s(v165+(v166+v167*v168)*v168)*int64(1000000) + v66
													v180 = v177 + v161
													if base.B2i32(v177 < int64(0))^base.B2i32(v180 < v161) != 0 {
														F_errstart_cold(m, int32(21), int32(0))
														mBase = m.M
														v244 = m.ExcPending
														if v244 != 0 {
															return int32(0)
														} else {
															F_errcode(m, int32(134217858))
															mBase = m.M
															v247 = m.ExcPending
															if v247 != 0 {
																return int32(0)
															} else {
																F_errmsg(m, int32(418494), int32(0))
																mBase = m.M
																v251 = m.ExcPending
																if v251 != 0 {
																	return int32(0)
																} else {
																	F_errfinish(m, int32(515856), int32(6612), int32(387606))
																	mBase = m.M
																	v256 = m.ExcPending
																	if v256 != 0 {
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
														if base.Ui64(v180-int64(9223371331200000000)) <= base.Ui64(int64(9011559254509551615)) {
															F_errstart_cold(m, int32(21), int32(0))
															mBase = m.M
															v244 = m.ExcPending
															if v244 != 0 {
																return int32(0)
															} else {
																F_errcode(m, int32(134217858))
																mBase = m.M
																v247 = m.ExcPending
																if v247 != 0 {
																	return int32(0)
																} else {
																	F_errmsg(m, int32(418494), int32(0))
																	mBase = m.M
																	v251 = m.ExcPending
																	if v251 != 0 {
																		return int32(0)
																	} else {
																		F_errfinish(m, int32(515856), int32(6612), int32(387606))
																		mBase = m.M
																		v256 = m.ExcPending
																		if v256 != 0 {
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
															v189 = v180
															if base.Ui64(int64(-9011559254509551616)) <= base.Ui64(v189+int64(211813488000000000)) {
																F_errstart_cold(m, int32(21), int32(0))
																mBase = m.M
																v225 = m.ExcPending
																if v225 != 0 {
																	return int32(0)
																} else {
																	F_errcode(m, int32(134217858))
																	mBase = m.M
																	v228 = m.ExcPending
																	if v228 != 0 {
																		return int32(0)
																	} else {
																		F_errmsg(m, int32(418494), int32(0))
																		mBase = m.M
																		v232 = m.ExcPending
																		if v232 != 0 {
																			return int32(0)
																		} else {
																			F_errfinish(m, int32(515856), int32(6618), int32(387606))
																			mBase = m.M
																			v237 = m.ExcPending
																			if v237 != 0 {
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
																v195 = F_Int64GetDatum(m, v189)
																mBase = m.M
																v196 = m.ExcPending
																if v196 != 0 {
																	return int32(0)
																} else {
																	v201 = v195
																	m.G0 = v8 + int32(352)
																	return v201
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
	}
}
