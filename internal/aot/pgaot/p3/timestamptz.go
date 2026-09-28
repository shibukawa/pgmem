package p3

import (
	base "github.com/shibukawa/pgmem/internal/aot/pgaot/base"
	"unsafe"
)

func F_timestamptz_date(m *base.Module, l0 int32) int64 {
	mBase := m.M
	_ = mBase
	var v4 int64
	_ = v4
	var v5 int32
	_ = v5
	var v6 int32
	_ = v6
	var v9 int32
	_ = v9
	var v10 int32
	_ = v10
	var v13 int32
	_ = v13
	var v16 int32
	_ = v16
	var v19 int32
	_ = v19
	v4 = *(*int64)(unsafe.Add(mBase, uint32(l0)+24))
	v5 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v6 = F_timestamptz2date_safe(m, v4, v5)
	mBase = m.M
	v9 = m.ExcPending
	if v9 != 0 {
		return int64(0)
	} else {
		v10 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
		if v10 == int32(0) {
			return base.I64_extend_i32_s(v6)
		} else {
			v13 = *(*int32)(unsafe.Add(mBase, uint32(v10)))
			if v13 != int32(453) {
				return base.I64_extend_i32_s(v6)
			} else {
				v16 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v10)+4)))
				if v16 != int32(1) {
					return base.I64_extend_i32_s(v6)
				} else {
					v19 = int32(1)
					*(*uint8)(unsafe.Add(mBase, uint32(l0)+16)) = uint8(v19)
					return int64(0)
				}
			}
		}
	}
}
func F_timestamptz_gt_date(m *base.Module, l0 int32) int64 {
	mBase := m.M
	_ = mBase
	var v5 int32
	_ = v5
	var v7 int32
	_ = v7
	var v9 int64
	_ = v9
	var v10 int32
	_ = v10
	var v12 int32
	_ = v12
	var v15 int64
	_ = v15
	var v17 int64
	_ = v17
	var v20 int32
	_ = v20
	var v21 int32
	_ = v21
	var v32 int32
	_ = v32
	var v37 int32
	_ = v37
	var v41 int32
	_ = v41
	v5 = m.G0
	v7 = v5 - int32(16)
	m.G0 = v7
	v9 = *(*int64)(unsafe.Add(mBase, uint32(l0)+24))
	v10 = *(*int32)(unsafe.Add(mBase, uint32(l0)+40))
	v12 = *(*int32)(unsafe.Add(mBase, _c_F_timestamptz_gt_date[0]))
	*(*int32)(unsafe.Add(mBase, uint32(v7)+8)) = v12
	v15 = *(*int64)(unsafe.Add(mBase, _c_F_timestamptz_gt_date[1]))
	*(*int64)(unsafe.Add(mBase, uint32(v7))) = v15
	v17 = F_date2timestamptz_safe(m, v10, v7)
	mBase = m.M
	v20 = m.ExcPending
	if v20 != 0 {
		return int64(0)
	} else {
		v21 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v7)+4)))
		if v21 != int32(1) {
			v41 = base.B2i32(v9 < v17) - base.B2i32(v17 < v9)
		} else {
			if v17 != int64(-9223372036854775807-1) {
				if v17 != int64(9223372036854775807) {
					v41 = base.B2i32(v9 < v17) - base.B2i32(v17 < v9)
				} else {
					if v9 == int64(9223372036854775807) {
						v32 = int32(-1)
					} else {
						v32 = int32(1)
					}
					v41 = v32
				}
			} else {
				if v9 == int64(-9223372036854775807-1) {
					v37 = int32(1)
				} else {
					v37 = int32(-1)
				}
				v41 = v37
			}
		}
		m.G0 = v7 + int32(16)
		return base.I64_extend_i32_u(int32(base.Ui32(v41) >> (uint(int32(31)) % 32)))
	}
}
func F_timestamptz_gt_timestamp(m *base.Module, l0 int32) int64 {
	mBase := m.M
	_ = mBase
	var v5 int32
	_ = v5
	var v7 int32
	_ = v7
	var v9 int64
	_ = v9
	var v10 int64
	_ = v10
	var v12 int32
	_ = v12
	var v15 int64
	_ = v15
	var v17 int64
	_ = v17
	var v20 int32
	_ = v20
	var v21 int32
	_ = v21
	var v33 int32
	_ = v33
	v5 = m.G0
	v7 = v5 - int32(16)
	m.G0 = v7
	v9 = *(*int64)(unsafe.Add(mBase, uint32(l0)+24))
	v10 = *(*int64)(unsafe.Add(mBase, uint32(l0)+40))
	v12 = *(*int32)(unsafe.Add(mBase, _c_F_timestamptz_gt_timestamp[0]))
	*(*int32)(unsafe.Add(mBase, uint32(v7)+8)) = v12
	v15 = *(*int64)(unsafe.Add(mBase, _c_F_timestamptz_gt_timestamp[1]))
	*(*int64)(unsafe.Add(mBase, uint32(v7))) = v15
	v17 = F_timestamp2timestamptz_safe(m, v10, v7)
	mBase = m.M
	v20 = m.ExcPending
	if v20 != 0 {
		return int64(0)
	} else {
		v21 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v7)+4)))
		if v21 != int32(1) {
			v33 = base.B2i32(v17 < v9)
		} else {
			if v17 != int64(-9223372036854775807-1) {
				if v17 != int64(9223372036854775807) {
					v33 = base.B2i32(v17 < v9)
				} else {
					v33 = base.B2i32(v9 == int64(9223372036854775807))
				}
			} else {
				v33 = base.B2i32(v9 != int64(-9223372036854775807-1))
			}
		}
		m.G0 = v7 + int32(16)
		return base.I64_extend_i32_u(v33)
	}
}
func F_timestamptz_le_timestamp(m *base.Module, l0 int32) int64 {
	mBase := m.M
	_ = mBase
	var v5 int32
	_ = v5
	var v7 int32
	_ = v7
	var v9 int64
	_ = v9
	var v10 int64
	_ = v10
	var v12 int32
	_ = v12
	var v15 int64
	_ = v15
	var v17 int64
	_ = v17
	var v20 int32
	_ = v20
	var v21 int32
	_ = v21
	var v33 int32
	_ = v33
	v5 = m.G0
	v7 = v5 - int32(16)
	m.G0 = v7
	v9 = *(*int64)(unsafe.Add(mBase, uint32(l0)+24))
	v10 = *(*int64)(unsafe.Add(mBase, uint32(l0)+40))
	v12 = *(*int32)(unsafe.Add(mBase, _c_F_timestamptz_le_timestamp[0]))
	*(*int32)(unsafe.Add(mBase, uint32(v7)+8)) = v12
	v15 = *(*int64)(unsafe.Add(mBase, _c_F_timestamptz_le_timestamp[1]))
	*(*int64)(unsafe.Add(mBase, uint32(v7))) = v15
	v17 = F_timestamp2timestamptz_safe(m, v10, v7)
	mBase = m.M
	v20 = m.ExcPending
	if v20 != 0 {
		return int64(0)
	} else {
		v21 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v7)+4)))
		if v21 != int32(1) {
			v33 = base.B2i32(v9 <= v17)
		} else {
			if v17 != int64(-9223372036854775807-1) {
				if v17 != int64(9223372036854775807) {
					v33 = base.B2i32(v9 <= v17)
				} else {
					v33 = base.B2i32(v9 != int64(9223372036854775807))
				}
			} else {
				v33 = base.B2i32(v9 == int64(-9223372036854775807-1))
			}
		}
		m.G0 = v7 + int32(16)
		return base.I64_extend_i32_u(v33)
	}
}
func F_timestamptz_out(m *base.Module, l0 int32) int64 {
	mBase := m.M
	_ = mBase
	var v4 int32
	_ = v4
	var v6 int32
	_ = v6
	var v8 int64
	_ = v8
	var v16 int32
	_ = v16
	var v19 int64
	_ = v19
	var v24 int32
	_ = v24
	var v30 int32
	_ = v30
	var v33 int32
	_ = v33
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
	var v50 int32
	_ = v50
	var v54 int32
	_ = v54
	var v59 int32
	_ = v59
	var v61 int32
	_ = v61
	var v64 int64
	_ = v64
	var v67 int32
	_ = v67
	var v68 int32
	_ = v68
	v4 = m.G0
	v6 = v4 - int32(192)
	m.G0 = v6
	v8 = *(*int64)(unsafe.Add(mBase, uint32(l0)+24))
	if base.Ui64(v8-int64(9223372036854775807)) <= base.Ui64(int64(1)) {
		if v8 != int64(-9223372036854775807-1) {
			v61 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_timestamptz_out[0])))
			*(*uint8)(unsafe.Add(mBase, uint32(v6)+8)) = uint8(v61)
			v64 = *(*int64)(unsafe.Add(mBase, _c_F_timestamptz_out[1]))
			*(*int64)(unsafe.Add(mBase, uint32(v6))) = v64
		} else {
			v16 = int32(*(*uint16)(unsafe.Add(mBase, _c_F_timestamptz_out[2])))
			*(*uint16)(unsafe.Add(mBase, uint32(v6)+8)) = uint16(v16)
			v19 = *(*int64)(unsafe.Add(mBase, _c_F_timestamptz_out[3]))
			*(*int64)(unsafe.Add(mBase, uint32(v6))) = v19
		}
		v67 = F_pstrdup(m, v6)
		mBase = m.M
		v68 = m.ExcPending
		if v68 != 0 {
			return int64(0)
		} else {
			m.G0 = v6 + int32(192)
			return base.I64_extend_i32_u(v67)
		}
	} else {
		v24 = v6 + int32(144)
		v30 = F_timestamp2tm(m, v8, v6+int32(188), v24, v6+int32(140), v6+int32(136), int32(0))
		mBase = m.M
		v33 = m.ExcPending
		if v33 != 0 {
			return int64(0)
		} else {
			if v30 == int32(0) {
				v36 = *(*int32)(unsafe.Add(mBase, uint32(v6)+140))
				v38 = *(*int32)(unsafe.Add(mBase, uint32(v6)+188))
				v39 = *(*int32)(unsafe.Add(mBase, uint32(v6)+136))
				v41 = *(*int32)(unsafe.Add(mBase, _c_F_timestamptz_out[4]))
				F_EncodeDateTime(m, v24, v36, int32(1), v38, v39, v41, v6)
				mBase = m.M
				v43 = m.ExcPending
				if v43 != 0 {
					return int64(0)
				} else {
					v67 = F_pstrdup(m, v6)
					mBase = m.M
					v68 = m.ExcPending
					if v68 != 0 {
						return int64(0)
					} else {
						m.G0 = v6 + int32(192)
						return base.I64_extend_i32_u(v67)
					}
				}
			} else {
				F_errstart_cold(m, int32(21), int32(0))
				mBase = m.M
				v47 = m.ExcPending
				if v47 != 0 {
					return int64(0)
				} else {
					F_errcode(m, int32(134217858))
					mBase = m.M
					v50 = m.ExcPending
					if v50 != 0 {
						return int64(0)
					} else {
						F_errmsg(m, int32(_a_F_timestamptz_out_0), int32(0))
						mBase = m.M
						v54 = m.ExcPending
						if v54 != 0 {
							return int64(0)
						} else {
							F_errfinish(m, int32(_a_F_timestamptz_out_1), int32(791), int32(_a_F_timestamptz_out_2))
							mBase = m.M
							v59 = m.ExcPending
							if v59 != 0 {
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
	var v146 int32
	_ = v146
	var v154 int32
	_ = v154
	var v160 int32
	_ = v160
	var v167 int32
	_ = v167
	var v171 int32
	_ = v171
	var v179 int32
	_ = v179
	var v185 int32
	_ = v185
	var v186 int32
	_ = v186
	var v187 int32
	_ = v187
	var v191 int32
	_ = v191
	var v195 int64
	_ = v195
	var v196 int32
	_ = v196
	var v201 int32
	_ = v201
	var v206 int32
	_ = v206
	var v209 int32
	_ = v209
	var v212 int32
	_ = v212
	var v214 int32
	_ = v214
	var v215 int32
	_ = v215
	var v220 int32
	_ = v220
	var v221 int32
	_ = v221
	var v222 int32
	_ = v222
	var v227 int32
	_ = v227
	var v230 int32
	_ = v230
	var v233 int32
	_ = v233
	var v237 int32
	_ = v237
	var v242 int32
	_ = v242
	var v248 int64
	_ = v248
	var v257 int64
	_ = v257
	var v258 int64
	_ = v258
	var v260 int64
	_ = v260
	var v263 int64
	_ = v263
	var v264 int64
	_ = v264
	var v266 int64
	_ = v266
	var v267 int64
	_ = v267
	var v271 int64
	_ = v271
	var v278 int64
	_ = v278
	var v289 int64
	_ = v289
	var v290 int64
	_ = v290
	var v294 int32
	_ = v294
	var v295 int32
	_ = v295
	var v296 int32
	_ = v296
	var v297 int32
	_ = v297
	var v306 int64
	_ = v306
	var v309 int64
	_ = v309
	var v317 int64
	_ = v317
	var v322 int64
	_ = v322
	var v327 int32
	_ = v327
	var v331 int32
	_ = v331
	var v335 int32
	_ = v335
	var v336 int32
	_ = v336
	var v337 int32
	_ = v337
	var v338 int32
	_ = v338
	var v339 int32
	_ = v339
	var v344 int32
	_ = v344
	var v345 int32
	_ = v345
	var v346 int32
	_ = v346
	var v351 int32
	_ = v351
	var v354 int32
	_ = v354
	var v357 int32
	_ = v357
	var v361 int32
	_ = v361
	var v366 int32
	_ = v366
	var v369 int32
	_ = v369
	var v370 int32
	_ = v370
	var v373 int32
	_ = v373
	var v388 int32
	_ = v388
	var v389 int32
	_ = v389
	var v390 int32
	_ = v390
	var v391 int32
	_ = v391
	var v397 int32
	_ = v397
	var v402 int32
	_ = v402
	var v405 int32
	_ = v405
	var v406 int32
	_ = v406
	var v407 int32
	_ = v407
	var v410 int32
	_ = v410
	var v412 int32
	_ = v412
	var v418 int32
	_ = v418
	var v422 int32
	_ = v422
	var v423 int32
	_ = v423
	var v425 int32
	_ = v425
	var v433 int32
	_ = v433
	var v437 int32
	_ = v437
	var v447 int32
	_ = v447
	var v452 int32
	_ = v452
	var v453 int32
	_ = v453
	var v454 int32
	_ = v454
	var v458 int32
	_ = v458
	var v462 int64
	_ = v462
	var v463 int32
	_ = v463
	var v468 int32
	_ = v468
	var v473 int32
	_ = v473
	var v476 int32
	_ = v476
	var v479 int32
	_ = v479
	var v480 int32
	_ = v480
	var v485 int32
	_ = v485
	var v486 int32
	_ = v486
	var v487 int32
	_ = v487
	var v492 int32
	_ = v492
	var v495 int32
	_ = v495
	var v498 int32
	_ = v498
	var v502 int32
	_ = v502
	var v507 int32
	_ = v507
	var v513 int64
	_ = v513
	var v522 int64
	_ = v522
	var v523 int64
	_ = v523
	var v525 int64
	_ = v525
	var v528 int64
	_ = v528
	var v529 int64
	_ = v529
	var v531 int64
	_ = v531
	var v532 int64
	_ = v532
	var v536 int64
	_ = v536
	var v543 int64
	_ = v543
	var v554 int64
	_ = v554
	var v555 int64
	_ = v555
	var v559 int32
	_ = v559
	var v560 int32
	_ = v560
	var v561 int32
	_ = v561
	var v562 int32
	_ = v562
	var v571 int64
	_ = v571
	var v574 int64
	_ = v574
	var v582 int64
	_ = v582
	var v587 int64
	_ = v587
	var v593 int64
	_ = v593
	var v596 int64
	_ = v596
	var v608 int64
	_ = v608
	var v616 int32
	_ = v616
	var v619 int32
	_ = v619
	var v623 int32
	_ = v623
	var v628 int32
	_ = v628
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
	var v715 int32
	_ = v715
	var v718 int32
	_ = v718
	var v722 int32
	_ = v722
	var v727 int32
	_ = v727
	var v734 int32
	_ = v734
	var v737 int32
	_ = v737
	var v741 int32
	_ = v741
	var v746 int32
	_ = v746
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
	v734 = m.ExcPending
	if v734 != 0 {
		goto L18
	} else {
		goto L155
	}
L2:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v715 = m.ExcPending
	if v715 != 0 {
		goto L18
	} else {
		goto L151
	}
L3:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v696 = m.ExcPending
	if v696 != 0 {
		goto L18
	} else {
		goto L147
	}
L4:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v680 = m.ExcPending
	if v680 != 0 {
		goto L18
	} else {
		goto L143
	}
L5:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v664 = m.ExcPending
	if v664 != 0 {
		goto L18
	} else {
		goto L139
	}
L6:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v648 = m.ExcPending
	if v648 != 0 {
		goto L18
	} else {
		goto L135
	}
L7:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v632 = m.ExcPending
	if v632 != 0 {
		goto L18
	} else {
		goto L131
	}
L8:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v616 = m.ExcPending
	if v616 != 0 {
		goto L18
	} else {
		goto L127
	}
L9:
	;
	m.G0 = v11 + int32(96)
	return v608
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
		v608 = v21
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
	F_errmsg(m, int32(_a_F_timestamptz_pl_interval_internal_0), int32(0))
	mBase = m.M
	v39 = m.ExcPending
	if v39 != 0 {
		goto L18
	} else {
		goto L21
	}
L21:
	;
	F_errfinish(m, int32(_a_F_timestamptz_pl_interval_internal_1), int32(3266), int32(_a_F_timestamptz_pl_interval_internal_2))
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
		v608 = v48
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
	F_errmsg(m, int32(_a_F_timestamptz_pl_interval_internal_0), int32(0))
	mBase = m.M
	v64 = m.ExcPending
	if v64 != 0 {
		goto L18
	} else {
		goto L28
	}
L28:
	;
	F_errfinish(m, int32(_a_F_timestamptz_pl_interval_internal_1), int32(3275), int32(_a_F_timestamptz_pl_interval_internal_2))
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
	v608 = l0
	goto L9
L31:
	;
	goto L32
L32:
	;
	v76 = *(*int32)(unsafe.Add(mBase, _c_F_timestamptz_pl_interval_internal[0]))
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
	v322 = l0
	goto L38
L38:
	;
	v327 = *(*int32)(unsafe.Add(mBase, uint32(l1)+8))
	if v327 != 0 {
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
	v185 = m.G0
	v186 = int32(16)
	v187 = v185 - v186
	m.G0 = v187
	v191 = F_DetermineTimeZoneOffsetInternal(m, v11+int32(48), v77, v187+int32(8))
	mBase = m.M
	m.G0 = v187 + v186
	goto L61
L51:
	;
	v179 = *(*int32)(unsafe.Add(mBase, uint32(v171*int32(52)+v131<<(uint(int32(2))%32))+uint32(_c_F_timestamptz_pl_interval_internal[1])))
	*(*int32)(unsafe.Add(mBase, uint32(v11)+60)) = v179
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
	v167 = *(*int32)(unsafe.Add(mBase, uint32(v131<<(uint(int32(2))%32))+uint32(_c_F_timestamptz_pl_interval_internal[1])))
	if v134 <= v167 {
		goto L50
	} else {
		goto L60
	}
L55:
	;
	v144 = base.I32_rem_s(v132, int32(400))
	v146 = base.B2i32(v144 == int32(0))
	v154 = *(*int32)(unsafe.Add(mBase, uint32(v146*int32(52)+v131<<(uint(int32(2))%32))+uint32(_c_F_timestamptz_pl_interval_internal[1])))
	if v134 <= v154 {
		goto L50
	} else {
		goto L58
	}
L56:
	;
	goto L57
L57:
	;
	v160 = *(*int32)(unsafe.Add(mBase, uint32(v131<<(uint(int32(2))%32))+uint32(_c_F_timestamptz_pl_interval_internal[2])))
	if v134 <= v160 {
		goto L50
	} else {
		goto L59
	}
L58:
	;
	v171 = v146
	goto L51
L59:
	;
	v171 = int32(1)
	goto L51
L60:
	;
	v171 = int32(0)
	goto L51
L61:
	;
	v195 = int64(*(*int32)(unsafe.Add(mBase, uint32(v11)+44)))
	v196 = *(*int32)(unsafe.Add(mBase, uint32(v11)+68))
	if v196 <= int32(-4713) {
		goto L63
	} else {
		goto L64
	}
L62:
	;
	v214 = v11 + int32(16)
	v215 = *(*int32)(unsafe.Add(mBase, uint32(v11)+60))
	v220 = base.B2i32(int32(2) < v212)
	if int32(2) < v212 {
		goto L74
	} else {
		goto L75
	}
L63:
	;
	if v196 != int32(-4713) {
		goto L1
	} else {
		goto L66
	}
L64:
	;
	goto L65
L65:
	;
	if v196 <= int32(_a_F_timestamptz_pl_interval_internal_11) {
		goto L68
	} else {
		goto L69
	}
L66:
	;
	v201 = *(*int32)(unsafe.Add(mBase, uint32(v11)+64))
	if int32(10) < v201 {
		v212 = v201
		goto L62
	} else {
		goto L67
	}
L67:
	;
	goto L1
L68:
	;
	v206 = *(*int32)(unsafe.Add(mBase, uint32(v11)+64))
	v212 = v206
	goto L62
L69:
	;
	goto L70
L70:
	;
	if v196 != int32(_a_F_timestamptz_pl_interval_internal_12) {
		goto L1
	} else {
		goto L71
	}
L71:
	;
	v209 = *(*int32)(unsafe.Add(mBase, uint32(v11)+64))
	if int32(5) < v209 {
		goto L1
	} else {
		goto L72
	}
L72:
	;
	v212 = v209
	goto L62
L73:
	;
	v248 = base.I64_extend_i32_s(v215 + v222*int32(365) + v227 + v230 + v233 + v242 - int32(_a_F_timestamptz_pl_interval_internal_6) - int32(_a_F_timestamptz_pl_interval_internal_7))
	v257 = int64(32)
	v258 = int64(20)
	v260 = int64(base.Ui64(v248) >> (uint(v257) % 64))
	v263 = int64(4294967295)
	v264 = int64(500654080)
	v266 = v248 & v263
	v267 = v264 * v266
	v271 = int64(base.Ui64(v267)>>(uint(v257)%64)) + v264*v260
	v278 = v266*v258 + v271&v263
	*(*int64)(unsafe.Add(mBase, uint32(v214)+8)) = v248*int64(0) + v248>>(uint(int64(63))%64)*int64(86400000000) + v258*v260 + int64(base.Ui64(v271)>>(uint(v257)%64)) + int64(base.Ui64(v278)>>(uint(v257)%64))
	*(*int64)(unsafe.Add(mBase, uint32(v214))) = v267&v263 | v278<<(uint(v257)%64)
	goto L80
L74:
	;
	v221 = int32(_a_F_timestamptz_pl_interval_internal_3)
	goto L76
L75:
	;
	v221 = int32(_a_F_timestamptz_pl_interval_internal_4)
	goto L76
L76:
	;
	v222 = v221 + v196
	v227 = base.I32_div_s(v222, int32(4))
	v230 = base.I32_div_s(v222, int32(-100))
	v233 = base.I32_div_s(v222, int32(400))
	if int32(2) < v212 {
		goto L77
	} else {
		goto L78
	}
L77:
	;
	v237 = int32(1)
	goto L79
L78:
	;
	v237 = int32(13)
	goto L79
L79:
	;
	v242 = base.I32_div_s((v237+v212)*int32(_a_F_timestamptz_pl_interval_internal_5), int32(256))
	goto L73
L80:
	;
	v289 = *(*int64)(unsafe.Add(mBase, uint32(v11)+24))
	v290 = *(*int64)(unsafe.Add(mBase, uint32(v11)+16))
	if v289 != v290>>(uint(int64(63))%64) {
		goto L1
	} else {
		goto L81
	}
L81:
	;
	v294 = *(*int32)(unsafe.Add(mBase, uint32(v11)+48))
	v295 = *(*int32)(unsafe.Add(mBase, uint32(v11)+52))
	v296 = *(*int32)(unsafe.Add(mBase, uint32(v11)+56))
	v297 = int32(60)
	v306 = base.I64_extend_i32_s(v294+(v295+v296*v297)*v297)*int64(1000000) + v195
	v309 = v290 + v306
	if base.B2i32(v306 < int64(0))^base.B2i32(v309 < v290) != 0 {
		goto L1
	} else {
		goto L82
	}
L82:
	;
	v317 = v309 + base.I64_extend_i32_s(int32(0)-v191)*int64(-1000000)
	if base.Ui64(int64(-9011559254509551616)) <= base.Ui64(v317+int64(211813488000000000)) {
		goto L1
	} else {
		goto L83
	}
L83:
	;
	v322 = v317
	goto L38
L84:
	;
	v331 = v11 + int32(48)
	v335 = F_timestamp2tm(m, v322, v11+int32(92), v331, v11+int32(44), int32(0), v77)
	mBase = m.M
	v336 = m.ExcPending
	if v336 != 0 {
		goto L18
	} else {
		goto L87
	}
L85:
	;
	v587 = v322
	goto L86
L86:
	;
	v593 = *(*int64)(unsafe.Add(mBase, uint32(l1)))
	v596 = v587 + v593
	if base.B2i32(v593 < int64(0)) != base.B2i32(v596 < v587) {
		goto L4
	} else {
		goto L125
	}
L87:
	;
	if v335 != 0 {
		goto L6
	} else {
		goto L88
	}
L88:
	;
	v337 = *(*int32)(unsafe.Add(mBase, uint32(v11)+68))
	v338 = *(*int32)(unsafe.Add(mBase, uint32(v11)+64))
	v339 = *(*int32)(unsafe.Add(mBase, uint32(v11)+60))
	v344 = base.B2i32(int32(2) < v338)
	if int32(2) < v338 {
		goto L90
	} else {
		goto L91
	}
L89:
	;
	v370 = *(*int32)(unsafe.Add(mBase, uint32(l1)+8))
	v373 = v369 + v370
	if base.B2i32(v370 < int32(0))^base.B2i32(v373 < v369)|base.B2i32(v373 <= int32(-2)) != 0 {
		goto L5
	} else {
		goto L96
	}
L90:
	;
	v345 = int32(_a_F_timestamptz_pl_interval_internal_3)
	goto L92
L91:
	;
	v345 = int32(_a_F_timestamptz_pl_interval_internal_4)
	goto L92
L92:
	;
	v346 = v345 + v337
	v351 = base.I32_div_s(v346, int32(4))
	v354 = base.I32_div_s(v346, int32(-100))
	v357 = base.I32_div_s(v346, int32(400))
	if int32(2) < v338 {
		goto L93
	} else {
		goto L94
	}
L93:
	;
	v361 = int32(1)
	goto L95
L94:
	;
	v361 = int32(13)
	goto L95
L95:
	;
	v366 = base.I32_div_s((v361+v338)*int32(_a_F_timestamptz_pl_interval_internal_5), int32(256))
	v369 = v339 + v346*int32(365) + v351 + v354 + v357 + v366 - int32(_a_F_timestamptz_pl_interval_internal_6)
	goto L89
L96:
	;
	v388 = v373 + int32(_a_F_timestamptz_pl_interval_internal_8)
	v389 = int32(_a_F_timestamptz_pl_interval_internal_9)
	v390 = base.I32_div_u_s(v388, v389)
	v391 = int32(3)
	v397 = int32(2)
	v402 = base.I32_div_u_s((v390*int32(1073595727)+v388)<<(uint(v397)%32)|v391, v389)
	v405 = v373 + v390*v391 + v402 + int32(_a_F_timestamptz_pl_interval_internal_10)
	v406 = int32(1461)
	v407 = base.I32_div_u_s(v405, v406)
	v410 = v407*int32(-1461) + v405
	v412 = v410 << (uint(v397) % 32)
	if base.Ui32(v406) <= base.Ui32(v412) {
		goto L99
	} else {
		goto L100
	}
L97:
	;
	v452 = m.G0
	v453 = int32(16)
	v454 = v452 - v453
	m.G0 = v454
	v458 = F_DetermineTimeZoneOffsetInternal(m, v331, v77, v454+int32(8))
	mBase = m.M
	m.G0 = v454 + v453
	goto L102
L98:
	;
	v425 = base.I32_div_u_s(v412, int32(1461))
	*(*int32)(unsafe.Add(mBase, uint32(v11+int32(68)))) = v425 + v407<<(uint(int32(2))%32) - int32(_a_F_timestamptz_pl_interval_internal_3)
	v433 = v423 + int32(123)
	v437 = int32(base.Ui32(v433*int32(2141)) >> (uint(int32(16)) % 32))
	*(*int32)(unsafe.Add(mBase, uint32(v11+int32(60)))) = v433 - int32(base.Ui32(v437*int32(_a_F_timestamptz_pl_interval_internal_5))>>(uint(int32(8))%32))
	v447 = base.I32_rem_u_s(v437+int32(10), int32(12))
	*(*int32)(unsafe.Add(mBase, uint32(v11-int32(-64)))) = v447 + int32(1)
	goto L97
L99:
	;
	v418 = base.I32_rem_u_s(v410+int32(305), int32(365))
	v423 = v418
	goto L98
L100:
	;
	goto L101
L101:
	;
	v422 = base.I32_rem_u_s(v410+int32(306), int32(366))
	v423 = v422
	goto L98
L102:
	;
	v462 = int64(*(*int32)(unsafe.Add(mBase, uint32(v11)+44)))
	v463 = *(*int32)(unsafe.Add(mBase, uint32(v11)+68))
	if v463 <= int32(-4713) {
		goto L104
	} else {
		goto L105
	}
L103:
	;
	v480 = *(*int32)(unsafe.Add(mBase, uint32(v11)+60))
	v485 = base.B2i32(int32(2) < v479)
	if int32(2) < v479 {
		goto L115
	} else {
		goto L116
	}
L104:
	;
	if v463 != int32(-4713) {
		goto L2
	} else {
		goto L107
	}
L105:
	;
	goto L106
L106:
	;
	if v463 <= int32(_a_F_timestamptz_pl_interval_internal_11) {
		goto L109
	} else {
		goto L110
	}
L107:
	;
	v468 = *(*int32)(unsafe.Add(mBase, uint32(v11)+64))
	if int32(10) < v468 {
		v479 = v468
		goto L103
	} else {
		goto L108
	}
L108:
	;
	goto L2
L109:
	;
	v473 = *(*int32)(unsafe.Add(mBase, uint32(v11)+64))
	v479 = v473
	goto L103
L110:
	;
	goto L111
L111:
	;
	if v463 != int32(_a_F_timestamptz_pl_interval_internal_12) {
		goto L2
	} else {
		goto L112
	}
L112:
	;
	v476 = *(*int32)(unsafe.Add(mBase, uint32(v11)+64))
	if int32(5) < v476 {
		goto L2
	} else {
		goto L113
	}
L113:
	;
	v479 = v476
	goto L103
L114:
	;
	v513 = base.I64_extend_i32_s(v480 + v487*int32(365) + v492 + v495 + v498 + v507 - int32(_a_F_timestamptz_pl_interval_internal_6) - int32(_a_F_timestamptz_pl_interval_internal_7))
	v522 = int64(32)
	v523 = int64(20)
	v525 = int64(base.Ui64(v513) >> (uint(v522) % 64))
	v528 = int64(4294967295)
	v529 = int64(500654080)
	v531 = v513 & v528
	v532 = v529 * v531
	v536 = int64(base.Ui64(v532)>>(uint(v522)%64)) + v529*v525
	v543 = v531*v523 + v536&v528
	*(*int64)(unsafe.Add(mBase, uint32(v11)+8)) = v513*int64(0) + v513>>(uint(int64(63))%64)*int64(86400000000) + v523*v525 + int64(base.Ui64(v536)>>(uint(v522)%64)) + int64(base.Ui64(v543)>>(uint(v522)%64))
	*(*int64)(unsafe.Add(mBase, uint32(v11))) = v532&v528 | v543<<(uint(v522)%64)
	goto L121
L115:
	;
	v486 = int32(_a_F_timestamptz_pl_interval_internal_3)
	goto L117
L116:
	;
	v486 = int32(_a_F_timestamptz_pl_interval_internal_4)
	goto L117
L117:
	;
	v487 = v486 + v463
	v492 = base.I32_div_s(v487, int32(4))
	v495 = base.I32_div_s(v487, int32(-100))
	v498 = base.I32_div_s(v487, int32(400))
	if int32(2) < v479 {
		goto L118
	} else {
		goto L119
	}
L118:
	;
	v502 = int32(1)
	goto L120
L119:
	;
	v502 = int32(13)
	goto L120
L120:
	;
	v507 = base.I32_div_s((v502+v479)*int32(_a_F_timestamptz_pl_interval_internal_5), int32(256))
	goto L114
L121:
	;
	v554 = *(*int64)(unsafe.Add(mBase, uint32(v11)+8))
	v555 = *(*int64)(unsafe.Add(mBase, uint32(v11)))
	if v554 != v555>>(uint(int64(63))%64) {
		goto L2
	} else {
		goto L122
	}
L122:
	;
	v559 = *(*int32)(unsafe.Add(mBase, uint32(v11)+48))
	v560 = *(*int32)(unsafe.Add(mBase, uint32(v11)+52))
	v561 = *(*int32)(unsafe.Add(mBase, uint32(v11)+56))
	v562 = int32(60)
	v571 = base.I64_extend_i32_s(v559+(v560+v561*v562)*v562)*int64(1000000) + v462
	v574 = v555 + v571
	if base.B2i32(v571 < int64(0))^base.B2i32(v574 < v555) != 0 {
		goto L2
	} else {
		goto L123
	}
L123:
	;
	v582 = v574 + base.I64_extend_i32_s(int32(0)-v458)*int64(-1000000)
	if base.Ui64(int64(-9011559254509551616)) <= base.Ui64(v582+int64(211813488000000000)) {
		goto L2
	} else {
		goto L124
	}
L124:
	;
	v587 = v582
	goto L86
L125:
	;
	if base.Ui64(int64(-9011559254509551616)) <= base.Ui64(v596+int64(211813488000000000)) {
		goto L3
	} else {
		goto L126
	}
L126:
	;
	v608 = v596
	goto L9
L127:
	;
	F_errcode(m, int32(134217858))
	mBase = m.M
	v619 = m.ExcPending
	if v619 != 0 {
		goto L18
	} else {
		goto L128
	}
L128:
	;
	F_errmsg(m, int32(_a_F_timestamptz_pl_interval_internal_0), int32(0))
	mBase = m.M
	v623 = m.ExcPending
	if v623 != 0 {
		goto L18
	} else {
		goto L129
	}
L129:
	;
	F_errfinish(m, int32(_a_F_timestamptz_pl_interval_internal_1), int32(3296), int32(_a_F_timestamptz_pl_interval_internal_2))
	mBase = m.M
	v628 = m.ExcPending
	if v628 != 0 {
		goto L18
	} else {
		goto L130
	}
L130:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L131:
	;
	F_errcode(m, int32(134217858))
	mBase = m.M
	v635 = m.ExcPending
	if v635 != 0 {
		goto L18
	} else {
		goto L132
	}
L132:
	;
	F_errmsg(m, int32(_a_F_timestamptz_pl_interval_internal_0), int32(0))
	mBase = m.M
	v639 = m.ExcPending
	if v639 != 0 {
		goto L18
	} else {
		goto L133
	}
L133:
	;
	F_errfinish(m, int32(_a_F_timestamptz_pl_interval_internal_1), int32(3301), int32(_a_F_timestamptz_pl_interval_internal_2))
	mBase = m.M
	v644 = m.ExcPending
	if v644 != 0 {
		goto L18
	} else {
		goto L134
	}
L134:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L135:
	;
	F_errcode(m, int32(134217858))
	mBase = m.M
	v651 = m.ExcPending
	if v651 != 0 {
		goto L18
	} else {
		goto L136
	}
L136:
	;
	F_errmsg(m, int32(_a_F_timestamptz_pl_interval_internal_0), int32(0))
	mBase = m.M
	v655 = m.ExcPending
	if v655 != 0 {
		goto L18
	} else {
		goto L137
	}
L137:
	;
	F_errfinish(m, int32(_a_F_timestamptz_pl_interval_internal_1), int32(3335), int32(_a_F_timestamptz_pl_interval_internal_2))
	mBase = m.M
	v660 = m.ExcPending
	if v660 != 0 {
		goto L18
	} else {
		goto L138
	}
L138:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L139:
	;
	F_errcode(m, int32(134217858))
	mBase = m.M
	v667 = m.ExcPending
	if v667 != 0 {
		goto L18
	} else {
		goto L140
	}
L140:
	;
	F_errmsg(m, int32(_a_F_timestamptz_pl_interval_internal_0), int32(0))
	mBase = m.M
	v671 = m.ExcPending
	if v671 != 0 {
		goto L18
	} else {
		goto L141
	}
L141:
	;
	F_errfinish(m, int32(_a_F_timestamptz_pl_interval_internal_1), int32(3349), int32(_a_F_timestamptz_pl_interval_internal_2))
	mBase = m.M
	v676 = m.ExcPending
	if v676 != 0 {
		goto L18
	} else {
		goto L142
	}
L142:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L143:
	;
	F_errcode(m, int32(134217858))
	mBase = m.M
	v683 = m.ExcPending
	if v683 != 0 {
		goto L18
	} else {
		goto L144
	}
L144:
	;
	F_errmsg(m, int32(_a_F_timestamptz_pl_interval_internal_0), int32(0))
	mBase = m.M
	v687 = m.ExcPending
	if v687 != 0 {
		goto L18
	} else {
		goto L145
	}
L145:
	;
	F_errfinish(m, int32(_a_F_timestamptz_pl_interval_internal_1), int32(3363), int32(_a_F_timestamptz_pl_interval_internal_2))
	mBase = m.M
	v692 = m.ExcPending
	if v692 != 0 {
		goto L18
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
	v699 = m.ExcPending
	if v699 != 0 {
		goto L18
	} else {
		goto L148
	}
L148:
	;
	F_errmsg(m, int32(_a_F_timestamptz_pl_interval_internal_0), int32(0))
	mBase = m.M
	v703 = m.ExcPending
	if v703 != 0 {
		goto L18
	} else {
		goto L149
	}
L149:
	;
	F_errfinish(m, int32(_a_F_timestamptz_pl_interval_internal_1), int32(3368), int32(_a_F_timestamptz_pl_interval_internal_2))
	mBase = m.M
	v708 = m.ExcPending
	if v708 != 0 {
		goto L18
	} else {
		goto L150
	}
L150:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L151:
	;
	F_errcode(m, int32(134217858))
	mBase = m.M
	v718 = m.ExcPending
	if v718 != 0 {
		goto L18
	} else {
		goto L152
	}
L152:
	;
	F_errmsg(m, int32(_a_F_timestamptz_pl_interval_internal_0), int32(0))
	mBase = m.M
	v722 = m.ExcPending
	if v722 != 0 {
		goto L18
	} else {
		goto L153
	}
L153:
	;
	F_errfinish(m, int32(_a_F_timestamptz_pl_interval_internal_1), int32(3357), int32(_a_F_timestamptz_pl_interval_internal_2))
	mBase = m.M
	v727 = m.ExcPending
	if v727 != 0 {
		goto L18
	} else {
		goto L154
	}
L154:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L155:
	;
	F_errcode(m, int32(134217858))
	mBase = m.M
	v737 = m.ExcPending
	if v737 != 0 {
		goto L18
	} else {
		goto L156
	}
L156:
	;
	F_errmsg(m, int32(_a_F_timestamptz_pl_interval_internal_0), int32(0))
	mBase = m.M
	v741 = m.ExcPending
	if v741 != 0 {
		goto L18
	} else {
		goto L157
	}
L157:
	;
	F_errfinish(m, int32(_a_F_timestamptz_pl_interval_internal_1), int32(3322), int32(_a_F_timestamptz_pl_interval_internal_2))
	mBase = m.M
	v746 = m.ExcPending
	if v746 != 0 {
		goto L18
	} else {
		goto L158
	}
L158:
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
	var v17 int32
	_ = v17
	var v19 int32
	_ = v19
	var v20 int32
	_ = v20
	var v27 int32
	_ = v27
	var v30 int32
	_ = v30
	var v37 int32
	_ = v37
	var v38 int32
	_ = v38
	var v42 int32
	_ = v42
	var v48 int32
	_ = v48
	var v50 int32
	_ = v50
	var v53 int32
	_ = v53
	var v59 int32
	_ = v59
	var v66 int32
	_ = v66
	var v74 int32
	_ = v74
	var v77 int32
	_ = v77
	var v79 int32
	_ = v79
	var v80 int32
	_ = v80
	var v85 int32
	_ = v85
	var v90 int32
	_ = v90
	var v91 int32
	_ = v91
	var v99 int32
	_ = v99
	var v100 int32
	_ = v100
	var v101 int32
	_ = v101
	var v102 int32
	_ = v102
	var v105 int32
	_ = v105
	var v106 int32
	_ = v106
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
	var v113 int32
	_ = v113
	var v115 int32
	_ = v115
	var v118 int32
	_ = v118
	var v121 int32
	_ = v121
	var v125 int32
	_ = v125
	var v128 int32
	_ = v128
	var v129 int32
	_ = v129
	var v130 int32
	_ = v130
	var v132 int32
	_ = v132
	var v135 int32
	_ = v135
	var v139 int32
	_ = v139
	var v142 int32
	_ = v142
	var v143 int32
	_ = v143
	var v146 int32
	_ = v146
	var v148 int32
	_ = v148
	var v150 int32
	_ = v150
	var v152 int32
	_ = v152
	var v159 int32
	_ = v159
	var v167 int32
	_ = v167
	var v170 int32
	_ = v170
	var v179 int32
	_ = v179
	var v181 int32
	_ = v181
	var v190 int32
	_ = v190
	var v195 int32
	_ = v195
	var v198 int32
	_ = v198
	var v201 int32
	_ = v201
	var v210 int32
	_ = v210
	var v213 int32
	_ = v213
	var v214 int32
	_ = v214
	var v217 int32
	_ = v217
	var v222 int32
	_ = v222
	var v227 int32
	_ = v227
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
	var v243 int32
	_ = v243
	var v248 int32
	_ = v248
	var v251 int32
	_ = v251
	var v252 int32
	_ = v252
	var v253 int32
	_ = v253
	var v256 int32
	_ = v256
	var v258 int32
	_ = v258
	var v264 int32
	_ = v264
	var v268 int32
	_ = v268
	var v269 int32
	_ = v269
	var v271 int32
	_ = v271
	var v279 int32
	_ = v279
	var v283 int32
	_ = v283
	var v293 int32
	_ = v293
	var v297 int32
	_ = v297
	var v303 int32
	_ = v303
	var v305 int32
	_ = v305
	var v308 int32
	_ = v308
	var v313 int32
	_ = v313
	var v314 int32
	_ = v314
	var v320 int32
	_ = v320
	var v323 int32
	_ = v323
	var v326 int32
	_ = v326
	var v331 int32
	_ = v331
	var v335 int32
	_ = v335
	var v339 int32
	_ = v339
	var v341 int32
	_ = v341
	var v344 int32
	_ = v344
	var v349 int32
	_ = v349
	var v362 int32
	_ = v362
	var v370 int32
	_ = v370
	var v371 int32
	_ = v371
	var v372 int32
	_ = v372
	var v376 int32
	_ = v376
	var v383 int32
	_ = v383
	var v385 int32
	_ = v385
	var v390 int32
	_ = v390
	var v393 int32
	_ = v393
	var v395 int32
	_ = v395
	var v396 int32
	_ = v396
	var v403 int32
	_ = v403
	var v408 int32
	_ = v408
	var v412 int32
	_ = v412
	var v417 int32
	_ = v417
	var v422 int32
	_ = v422
	var v427 int32
	_ = v427
	var v430 int32
	_ = v430
	var v433 int32
	_ = v433
	var v435 int32
	_ = v435
	var v436 int32
	_ = v436
	var v441 int32
	_ = v441
	var v442 int32
	_ = v442
	var v443 int32
	_ = v443
	var v448 int32
	_ = v448
	var v451 int32
	_ = v451
	var v454 int32
	_ = v454
	var v458 int32
	_ = v458
	var v463 int32
	_ = v463
	var v469 int64
	_ = v469
	var v478 int64
	_ = v478
	var v479 int64
	_ = v479
	var v481 int64
	_ = v481
	var v484 int64
	_ = v484
	var v485 int64
	_ = v485
	var v487 int64
	_ = v487
	var v488 int64
	_ = v488
	var v492 int64
	_ = v492
	var v499 int64
	_ = v499
	var v510 int64
	_ = v510
	var v511 int64
	_ = v511
	var v516 int32
	_ = v516
	var v517 int32
	_ = v517
	var v518 int32
	_ = v518
	var v519 int32
	_ = v519
	var v528 int64
	_ = v528
	var v531 int64
	_ = v531
	var v535 int32
	_ = v535
	var v540 int64
	_ = v540
	var v551 int32
	_ = v551
	var v554 int32
	_ = v554
	var v558 int32
	_ = v558
	var v563 int32
	_ = v563
	var v567 int32
	_ = v567
	var v570 int32
	_ = v570
	var v572 int32
	_ = v572
	var v573 int32
	_ = v573
	var v580 int32
	_ = v580
	var v585 int32
	_ = v585
	var v587 int64
	_ = v587
	var v602 int32
	_ = v602
	var v605 int32
	_ = v605
	var v609 int32
	_ = v609
	var v614 int32
	_ = v614
	v11 = m.G0
	v13 = v11 - int32(112)
	m.G0 = v13
	v15 = int32(1)
	v17 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0))))
	v19 = v17 & v15
	if v19 != 0 {
		goto L3
	} else {
		goto L4
	}
L1:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v602 = m.ExcPending
	if v602 != 0 {
		goto L17
	} else {
		goto L128
	}
L2:
	;
	m.G0 = v13 + int32(112)
	return v587
L3:
	;
	v20 = v15
	goto L5
L4:
	;
	v20 = int32(4)
	goto L5
L5:
	;
	if v17 == int32(1) {
		goto L7
	} else {
		goto L8
	}
L6:
	;
	v50 = F_downcase_truncate_identifier(m, l0+v20, v48, int32(0))
	mBase = m.M
	v53 = m.ExcPending
	if v53 != 0 {
		goto L17
	} else {
		goto L18
	}
L7:
	;
	v27 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+1)))
	if v27 == int32(18) {
		goto L10
	} else {
		goto L11
	}
L8:
	;
	goto L9
L9:
	;
	v38 = int32(1)
	if v19 != 0 {
		v48 = int32(base.Ui32(v17)>>(uint(v38)%32)) - v38
		goto L6
	} else {
		goto L16
	}
L10:
	;
	v30 = int32(16)
	goto L12
L11:
	;
	v30 = int32(0)
	goto L12
L12:
	;
	if base.Ui32((v27-int32(1))&int32(255)) < base.Ui32(int32(3)) {
		goto L13
	} else {
		goto L14
	}
L13:
	;
	v37 = int32(4)
	goto L15
L14:
	;
	v37 = v30
	goto L15
L15:
	;
	v48 = v37
	goto L6
L16:
	;
	v42 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v48 = int32(base.Ui32(v42)>>(uint(int32(2))%32)) - int32(4)
	goto L6
L17:
	;
	return int64(0)
L18:
	;
	v59 = Fn14210(m, v50, v13+int32(104), int32(_a_F_timestamptz_trunc_internal_0), int32(_a_F_timestamptz_trunc_internal_1), int32(_a_F_timestamptz_trunc_internal_2))
	mBase = m.M
	goto L19
L19:
	;
	if v59 == int32(17) {
		goto L20
	} else {
		goto L21
	}
L20:
	;
	if base.Ui64(l1-int64(9223372036854775807)) <= base.Ui64(int64(1)) {
		goto L23
	} else {
		goto L24
	}
L21:
	;
	goto L22
L22:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v567 = m.ExcPending
	if v567 != 0 {
		goto L17
	} else {
		goto L123
	}
L23:
	;
	v66 = *(*int32)(unsafe.Add(mBase, uint32(v13)+104))
	if base.Ui32(v66-int32(18)) < base.Ui32(int32(13)) {
		v587 = l1
		goto L2
	} else {
		goto L26
	}
L24:
	;
	goto L25
L25:
	;
	v91 = int32(0)
	v99 = F_timestamp2tm(m, l1, v13+int32(108), v13+int32(56), v13+int32(100), v91, l2)
	mBase = m.M
	v100 = m.ExcPending
	if v100 != 0 {
		goto L17
	} else {
		goto L32
	}
L26:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v74 = m.ExcPending
	if v74 != 0 {
		goto L17
	} else {
		goto L27
	}
L27:
	;
	F_errcode(m, int32(1088))
	mBase = m.M
	v77 = m.ExcPending
	if v77 != 0 {
		goto L17
	} else {
		goto L28
	}
L28:
	;
	v79 = F_format_type_be(m, int32(1184))
	mBase = m.M
	v80 = m.ExcPending
	if v80 != 0 {
		goto L17
	} else {
		goto L29
	}
L29:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v13)+4)) = v79
	*(*int32)(unsafe.Add(mBase, uint32(v13))) = v50
	F_errmsg(m, int32(_a_F_timestamptz_trunc_internal_3), v13)
	mBase = m.M
	v85 = m.ExcPending
	if v85 != 0 {
		goto L17
	} else {
		goto L30
	}
L30:
	;
	F_errfinish(m, int32(_a_F_timestamptz_trunc_internal_4), int32(_a_F_timestamptz_trunc_internal_5), int32(_a_F_timestamptz_trunc_internal_6))
	mBase = m.M
	v90 = m.ExcPending
	if v90 != 0 {
		goto L17
	} else {
		goto L31
	}
L31:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L32:
	;
	if v99 != 0 {
		goto L1
	} else {
		goto L33
	}
L33:
	;
	v101 = int32(1)
	v102 = *(*int32)(unsafe.Add(mBase, uint32(v13)+104))
	switch v102 - int32(18) {
	case 0:
		v412 = v91
		goto L34
	case 1:
		goto L35
	case 2:
		goto L38
	case 3:
		goto L39
	case 4:
		goto L45
	case 5:
		goto L40
	case 6:
		goto L47
	case 7:
		v344 = v101
		goto L41
	case 8:
		goto L42
	case 9:
		goto L46
	case 10:
		goto L44
	case 11:
		goto L37
	case 12:
		goto L48
	default:
		goto L36
	}
L34:
	;
	v417 = *(*int32)(unsafe.Add(mBase, uint32(v13)+76))
	if v417 <= int32(-4713) {
		goto L98
	} else {
		goto L99
	}
L35:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v13)+56)) = int32(0)
	v412 = v91
	goto L34
L36:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v390 = m.ExcPending
	if v390 != 0 {
		goto L17
	} else {
		goto L91
	}
L37:
	;
	v383 = *(*int32)(unsafe.Add(mBase, uint32(v13)+100))
	v385 = base.I32_rem_s(v383, int32(1000))
	v412 = v383 - v385
	goto L34
L38:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v13)+56)) = int64(0)
	v412 = v91
	goto L34
L39:
	;
	v362 = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v13)+64)) = v362
	*(*int64)(unsafe.Add(mBase, uint32(v13)+56)) = int64(0)
	v370 = m.G0
	v371 = int32(16)
	v372 = v370 - v371
	m.G0 = v372
	v376 = F_DetermineTimeZoneOffsetInternal(m, v13+int32(56), l2, v372+int32(8))
	mBase = m.M
	m.G0 = v372 + v371
	goto L90
L40:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v13)+68)) = int32(1)
	goto L39
L41:
	;
	v349 = base.I32_rem_s(v344-int32(1), int32(3))
	*(*int32)(unsafe.Add(mBase, uint32(v13)+72)) = v344 - v349
	goto L40
L42:
	;
	v331 = *(*int32)(unsafe.Add(mBase, uint32(v13)+76))
	if int32(0) < v331 {
		goto L87
	} else {
		goto L88
	}
L43:
	;
	v314 = int32(1)
	if int32(0) < v313 {
		goto L84
	} else {
		goto L85
	}
L44:
	;
	v297 = *(*int32)(unsafe.Add(mBase, uint32(v13)+76))
	if int32(0) < v297 {
		goto L81
	} else {
		goto L82
	}
L45:
	;
	v108 = *(*int32)(unsafe.Add(mBase, uint32(v13)+76))
	v109 = *(*int32)(unsafe.Add(mBase, uint32(v13)+72))
	v110 = *(*int32)(unsafe.Add(mBase, uint32(v13)+68))
	v112 = F_date2j(m, v108, v109, v110)
	mBase = m.M
	v113 = int32(1)
	v115 = F_date2j(m, v108, v113, int32(4))
	mBase = m.M
	v118 = F_j2day(m, v115-v113)
	mBase = m.M
	if v112 < v115-v118 {
		goto L50
	} else {
		goto L51
	}
L46:
	;
	v107 = *(*int32)(unsafe.Add(mBase, uint32(v13)+76))
	v313 = v107
	goto L43
L47:
	;
	v106 = *(*int32)(unsafe.Add(mBase, uint32(v13)+72))
	v344 = v106
	goto L41
L48:
	;
	v105 = *(*int32)(unsafe.Add(mBase, uint32(v13)+100))
	v412 = v105
	goto L34
L49:
	;
	v159 = *(*int32)(unsafe.Add(mBase, uint32(v13)+72))
	if base.B2i32(v159 != int32(1))|base.B2i32(v152 < int32(52)) == int32(0) {
		goto L61
	} else {
		goto L62
	}
L50:
	;
	v121 = int32(1)
	v125 = F_date2j(m, v108-v121, v121, int32(4))
	mBase = m.M
	v128 = F_j2day(m, v125-v121)
	mBase = m.M
	v129 = v125
	v130 = v128
	goto L52
L51:
	;
	v129 = v115
	v130 = v118
	goto L52
L52:
	;
	v132 = v130 - v129 + v112
	if int32(357) <= v132 {
		goto L53
	} else {
		goto L54
	}
L53:
	;
	v135 = int32(1)
	v139 = F_date2j(m, v108+v135, v135, int32(4))
	mBase = m.M
	v142 = F_j2day(m, v139-v135)
	mBase = m.M
	v143 = v139 - v142
	if v112 < v143 {
		goto L56
	} else {
		goto L57
	}
L54:
	;
	v148 = v132
	goto L55
L55:
	;
	v150 = base.I32_div_s(v148, int32(7))
	v152 = v150 + int32(1)
	goto L49
L56:
	;
	v146 = v132
	goto L58
L57:
	;
	v146 = v112 - v143
	goto L58
L58:
	;
	v148 = v146
	goto L55
L59:
	;
	goto L67
L60:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v13)+76)) = v179
	v181 = v179
	goto L59
L61:
	;
	v167 = *(*int32)(unsafe.Add(mBase, uint32(v13)+76))
	v179 = v167 - int32(1)
	goto L60
L62:
	;
	goto L63
L63:
	;
	v170 = *(*int32)(unsafe.Add(mBase, uint32(v13)+76))
	if base.B2i32(v159 != int32(12))|base.B2i32(int32(1) < v152) != 0 {
		v181 = v170
		goto L59
	} else {
		goto L64
	}
L64:
	;
	v179 = v170 + int32(1)
	goto L60
L65:
	;
	v214 = int32(7)
	v217 = int32(1)
	v222 = base.I32_rem_s(v213-v217+v217, v214)
	if v222 < int32(0) {
		goto L73
	} else {
		goto L74
	}
L67:
	;
	goto L68
L68:
	;
	v190 = int32(_a_F_timestamptz_trunc_internal_11) + v181
	v195 = base.I32_div_s(v190, int32(4))
	v198 = base.I32_div_s(v190, int32(-100))
	v201 = base.I32_div_s(v190, int32(400))
	goto L70
L70:
	;
	goto L71
L71:
	;
	v210 = base.I32_div_s(int32(_a_F_timestamptz_trunc_internal_17), int32(256))
	v213 = int32(4) + v190*int32(365) + v195 + v198 + v201 + v210 - int32(_a_F_timestamptz_trunc_internal_13)
	goto L65
L72:
	;
	v230 = v213 + v152*v214 - v227 - int32(7)
	v234 = v230 + int32(_a_F_timestamptz_trunc_internal_18)
	v235 = int32(_a_F_timestamptz_trunc_internal_19)
	v236 = base.I32_div_u_s(v234, v235)
	v237 = int32(3)
	v243 = int32(2)
	v248 = base.I32_div_u_s((v236*int32(1073595727)+v234)<<(uint(v243)%32)|v237, v235)
	v251 = v230 + v236*v237 + v248 + int32(_a_F_timestamptz_trunc_internal_20)
	v252 = int32(1461)
	v253 = base.I32_div_u_s(v251, v252)
	v256 = v253*int32(-1461) + v251
	v258 = v256 << (uint(v243) % 32)
	if base.Ui32(v252) <= base.Ui32(v258) {
		goto L78
	} else {
		goto L79
	}
L73:
	;
	v227 = v222 + v214
	goto L75
L74:
	;
	v227 = v222
	goto L75
L75:
	;
	goto L72
L76:
	;
	goto L39
L77:
	;
	v271 = base.I32_div_u_s(v258, int32(1461))
	*(*int32)(unsafe.Add(mBase, uint32(v13+int32(76)))) = v271 + v253<<(uint(int32(2))%32) - int32(_a_F_timestamptz_trunc_internal_10)
	v279 = v269 + int32(123)
	v283 = int32(base.Ui32(v279*int32(2141)) >> (uint(int32(16)) % 32))
	*(*int32)(unsafe.Add(mBase, uint32(v13+int32(68)))) = v279 - int32(base.Ui32(v283*int32(_a_F_timestamptz_trunc_internal_12))>>(uint(int32(8))%32))
	v293 = base.I32_rem_u_s(v283+int32(10), int32(12))
	*(*int32)(unsafe.Add(mBase, uint32(v13+int32(72)))) = v293 + int32(1)
	goto L76
L78:
	;
	v264 = base.I32_rem_u_s(v256+int32(305), int32(365))
	v269 = v264
	goto L77
L79:
	;
	goto L80
L80:
	;
	v268 = base.I32_rem_u_s(v256+int32(306), int32(366))
	v269 = v268
	goto L77
L81:
	;
	v303 = base.I32_rem_s(v297+int32(999), int32(1000))
	v313 = v297 - v303
	goto L43
L82:
	;
	goto L83
L83:
	;
	v305 = int32(1000)
	v308 = base.I32_rem_s(v305-v297, v305)
	v313 = v297 + v308 - int32(999)
	goto L43
L84:
	;
	v320 = base.I32_rem_s(v313+int32(99), int32(100))
	*(*int32)(unsafe.Add(mBase, uint32(v13)+76)) = v313 - v320
	v344 = v314
	goto L41
L85:
	;
	goto L86
L86:
	;
	v323 = int32(100)
	v326 = base.I32_rem_s(v323-v313, v323)
	*(*int32)(unsafe.Add(mBase, uint32(v13)+76)) = v313 + v326 - int32(99)
	v344 = v314
	goto L41
L87:
	;
	v335 = base.I32_rem_u_s(v331, int32(10))
	*(*int32)(unsafe.Add(mBase, uint32(v13)+76)) = v331 - v335
	v344 = v101
	goto L41
L88:
	;
	goto L89
L89:
	;
	v339 = int32(9) - v331
	v341 = base.I32_rem_s(v339, int32(10))
	*(*int32)(unsafe.Add(mBase, uint32(v13)+76)) = v341 - v339
	v344 = v101
	goto L41
L90:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v13)+108)) = v376
	v412 = v362
	goto L34
L91:
	;
	F_errcode(m, int32(1088))
	mBase = m.M
	v393 = m.ExcPending
	if v393 != 0 {
		goto L17
	} else {
		goto L92
	}
L92:
	;
	v395 = F_format_type_be(m, int32(1184))
	mBase = m.M
	v396 = m.ExcPending
	if v396 != 0 {
		goto L17
	} else {
		goto L93
	}
L93:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v13)+20)) = v395
	*(*int32)(unsafe.Add(mBase, uint32(v13)+16)) = v50
	F_errmsg(m, int32(_a_F_timestamptz_trunc_internal_3), v13+int32(16))
	mBase = m.M
	v403 = m.ExcPending
	if v403 != 0 {
		goto L17
	} else {
		goto L94
	}
L94:
	;
	F_errfinish(m, int32(_a_F_timestamptz_trunc_internal_4), int32(_a_F_timestamptz_trunc_internal_21), int32(_a_F_timestamptz_trunc_internal_6))
	mBase = m.M
	v408 = m.ExcPending
	if v408 != 0 {
		goto L17
	} else {
		goto L95
	}
L95:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L96:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v551 = m.ExcPending
	if v551 != 0 {
		goto L17
	} else {
		goto L119
	}
L97:
	;
	v435 = v13 + int32(32)
	v436 = *(*int32)(unsafe.Add(mBase, uint32(v13)+68))
	v441 = base.B2i32(int32(2) < v433)
	if int32(2) < v433 {
		goto L109
	} else {
		goto L110
	}
L98:
	;
	if v417 != int32(-4713) {
		goto L96
	} else {
		goto L101
	}
L99:
	;
	goto L100
L100:
	;
	if v417 <= int32(_a_F_timestamptz_trunc_internal_15) {
		goto L103
	} else {
		goto L104
	}
L101:
	;
	v422 = *(*int32)(unsafe.Add(mBase, uint32(v13)+72))
	if int32(10) < v422 {
		v433 = v422
		goto L97
	} else {
		goto L102
	}
L102:
	;
	goto L96
L103:
	;
	v427 = *(*int32)(unsafe.Add(mBase, uint32(v13)+72))
	v433 = v427
	goto L97
L104:
	;
	goto L105
L105:
	;
	if v417 != int32(_a_F_timestamptz_trunc_internal_16) {
		goto L96
	} else {
		goto L106
	}
L106:
	;
	v430 = *(*int32)(unsafe.Add(mBase, uint32(v13)+72))
	if int32(5) < v430 {
		goto L96
	} else {
		goto L107
	}
L107:
	;
	v433 = v430
	goto L97
L108:
	;
	v469 = base.I64_extend_i32_s(v436 + v443*int32(365) + v448 + v451 + v454 + v463 - int32(_a_F_timestamptz_trunc_internal_13) - int32(_a_F_timestamptz_trunc_internal_14))
	v478 = int64(32)
	v479 = int64(20)
	v481 = int64(base.Ui64(v469) >> (uint(v478) % 64))
	v484 = int64(4294967295)
	v485 = int64(500654080)
	v487 = v469 & v484
	v488 = v485 * v487
	v492 = int64(base.Ui64(v488)>>(uint(v478)%64)) + v485*v481
	v499 = v487*v479 + v492&v484
	*(*int64)(unsafe.Add(mBase, uint32(v435)+8)) = v469*int64(0) + v469>>(uint(int64(63))%64)*int64(86400000000) + v479*v481 + int64(base.Ui64(v492)>>(uint(v478)%64)) + int64(base.Ui64(v499)>>(uint(v478)%64))
	*(*int64)(unsafe.Add(mBase, uint32(v435))) = v488&v484 | v499<<(uint(v478)%64)
	goto L115
L109:
	;
	v442 = int32(_a_F_timestamptz_trunc_internal_10)
	goto L111
L110:
	;
	v442 = int32(_a_F_timestamptz_trunc_internal_11)
	goto L111
L111:
	;
	v443 = v442 + v417
	v448 = base.I32_div_s(v443, int32(4))
	v451 = base.I32_div_s(v443, int32(-100))
	v454 = base.I32_div_s(v443, int32(400))
	if int32(2) < v433 {
		goto L112
	} else {
		goto L113
	}
L112:
	;
	v458 = int32(1)
	goto L114
L113:
	;
	v458 = int32(13)
	goto L114
L114:
	;
	v463 = base.I32_div_s((v458+v433)*int32(_a_F_timestamptz_trunc_internal_12), int32(256))
	goto L108
L115:
	;
	v510 = *(*int64)(unsafe.Add(mBase, uint32(v13)+40))
	v511 = *(*int64)(unsafe.Add(mBase, uint32(v13)+32))
	if v510 != v511>>(uint(int64(63))%64) {
		goto L96
	} else {
		goto L116
	}
L116:
	;
	v516 = *(*int32)(unsafe.Add(mBase, uint32(v13)+56))
	v517 = *(*int32)(unsafe.Add(mBase, uint32(v13)+60))
	v518 = *(*int32)(unsafe.Add(mBase, uint32(v13)+64))
	v519 = int32(60)
	v528 = base.I64_extend_i32_s(v412) + base.I64_extend_i32_s(v516+(v517+v518*v519)*v519)*int64(1000000)
	v531 = v511 + v528
	if base.B2i32(v528 < int64(0))^base.B2i32(v531 < v511) != 0 {
		goto L96
	} else {
		goto L117
	}
L117:
	;
	v535 = *(*int32)(unsafe.Add(mBase, uint32(v13)+108))
	v540 = base.I64_extend_i32_s(int32(0)-v535)*int64(-1000000) + v531
	if base.Ui64(v540+int64(211813488000000000)) < base.Ui64(int64(-9011559254509551616)) {
		v587 = v540
		goto L2
	} else {
		goto L118
	}
L118:
	;
	goto L96
L119:
	;
	F_errcode(m, int32(134217858))
	mBase = m.M
	v554 = m.ExcPending
	if v554 != 0 {
		goto L17
	} else {
		goto L120
	}
L120:
	;
	F_errmsg(m, int32(_a_F_timestamptz_trunc_internal_7), int32(0))
	mBase = m.M
	v558 = m.ExcPending
	if v558 != 0 {
		goto L17
	} else {
		goto L121
	}
L121:
	;
	F_errfinish(m, int32(_a_F_timestamptz_trunc_internal_4), int32(_a_F_timestamptz_trunc_internal_9), int32(_a_F_timestamptz_trunc_internal_6))
	mBase = m.M
	v563 = m.ExcPending
	if v563 != 0 {
		goto L17
	} else {
		goto L122
	}
L122:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L123:
	;
	F_errcode(m, int32(50856066))
	mBase = m.M
	v570 = m.ExcPending
	if v570 != 0 {
		goto L17
	} else {
		goto L124
	}
L124:
	;
	v572 = F_format_type_be(m, int32(1184))
	mBase = m.M
	v573 = m.ExcPending
	if v573 != 0 {
		goto L17
	} else {
		goto L125
	}
L125:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v13)+52)) = v572
	*(*int32)(unsafe.Add(mBase, uint32(v13)+48)) = v50
	F_errmsg(m, int32(_a_F_timestamptz_trunc_internal_22), v13+int32(48))
	mBase = m.M
	v580 = m.ExcPending
	if v580 != 0 {
		goto L17
	} else {
		goto L126
	}
L126:
	;
	F_errfinish(m, int32(_a_F_timestamptz_trunc_internal_4), int32(_a_F_timestamptz_trunc_internal_23), int32(_a_F_timestamptz_trunc_internal_6))
	mBase = m.M
	v585 = m.ExcPending
	if v585 != 0 {
		goto L17
	} else {
		goto L127
	}
L127:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L128:
	;
	F_errcode(m, int32(134217858))
	mBase = m.M
	v605 = m.ExcPending
	if v605 != 0 {
		goto L17
	} else {
		goto L129
	}
L129:
	;
	F_errmsg(m, int32(_a_F_timestamptz_trunc_internal_7), int32(0))
	mBase = m.M
	v609 = m.ExcPending
	if v609 != 0 {
		goto L17
	} else {
		goto L130
	}
L130:
	;
	F_errfinish(m, int32(_a_F_timestamptz_trunc_internal_4), int32(_a_F_timestamptz_trunc_internal_8), int32(_a_F_timestamptz_trunc_internal_6))
	mBase = m.M
	v614 = m.ExcPending
	if v614 != 0 {
		goto L17
	} else {
		goto L131
	}
L131:
	;
	base.Wasm_trap_unreachable()
	for {
	}
}
func F_timestamptz_zone(m *base.Module, l0 int32) int64 {
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
	var v15 int64
	_ = v15
	var v21 int32
	_ = v21
	var v24 int32
	_ = v24
	var v29 int32
	_ = v29
	var v30 int32
	_ = v30
	var v32 int32
	_ = v32
	var v40 int32
	_ = v40
	var v43 int32
	_ = v43
	var v44 int32
	_ = v44
	var v56 int32
	_ = v56
	var v57 int32
	_ = v57
	var v58 int32
	_ = v58
	var v59 int64
	_ = v59
	var v60 int32
	_ = v60
	var v65 int32
	_ = v65
	var v70 int32
	_ = v70
	var v73 int32
	_ = v73
	var v76 int32
	_ = v76
	var v77 int32
	_ = v77
	var v82 int32
	_ = v82
	var v83 int32
	_ = v83
	var v84 int32
	_ = v84
	var v89 int32
	_ = v89
	var v92 int32
	_ = v92
	var v95 int32
	_ = v95
	var v99 int32
	_ = v99
	var v104 int32
	_ = v104
	var v110 int64
	_ = v110
	var v119 int64
	_ = v119
	var v120 int64
	_ = v120
	var v122 int64
	_ = v122
	var v125 int64
	_ = v125
	var v126 int64
	_ = v126
	var v128 int64
	_ = v128
	var v129 int64
	_ = v129
	var v133 int64
	_ = v133
	var v140 int64
	_ = v140
	var v151 int64
	_ = v151
	var v152 int64
	_ = v152
	var v156 int32
	_ = v156
	var v157 int32
	_ = v157
	var v158 int32
	_ = v158
	var v159 int32
	_ = v159
	var v168 int64
	_ = v168
	var v171 int64
	_ = v171
	var v181 int64
	_ = v181
	var v189 int64
	_ = v189
	var v198 int32
	_ = v198
	var v201 int32
	_ = v201
	var v205 int32
	_ = v205
	var v210 int32
	_ = v210
	var v214 int32
	_ = v214
	var v217 int32
	_ = v217
	var v221 int32
	_ = v221
	var v226 int32
	_ = v226
	var v233 int32
	_ = v233
	var v236 int32
	_ = v236
	var v240 int32
	_ = v240
	var v245 int32
	_ = v245
	v6 = m.G0
	v8 = v6 - int32(352)
	m.G0 = v8
	v10 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
	v11 = F_pg_detoast_datum_packed(m, v10)
	mBase = m.M
	v14 = m.ExcPending
	if v14 != 0 {
		return int64(0)
	} else {
		v15 = *(*int64)(unsafe.Add(mBase, uint32(l0)+40))
		if base.Ui64(int64(2)) <= base.Ui64(v15-int64(9223372036854775807)) {
			v21 = v8 + int32(80)
			F_text_to_cstring_buffer(m, v11, v21, int32(256))
			mBase = m.M
			v24 = m.ExcPending
			if v24 != 0 {
				return int64(0)
			} else {
				v29 = F_DecodeTimezoneName(m, v21, v8+int32(76), v8+int32(72))
				mBase = m.M
				v30 = m.ExcPending
				if v30 != 0 {
					return int64(0)
				} else {
					switch v29 {
					case 0:
						v32 = *(*int32)(unsafe.Add(mBase, uint32(v8)+76))
						v181 = base.I64_extend_i32_s(int32(0)-v32)*int64(-1000000) + v15
						if base.Ui64(int64(-9011559254509551616)) <= base.Ui64(v181+int64(211813488000000000)) {
							F_errstart_cold(m, int32(21), int32(0))
							mBase = m.M
							v214 = m.ExcPending
							if v214 != 0 {
								return int64(0)
							} else {
								F_errcode(m, int32(134217858))
								mBase = m.M
								v217 = m.ExcPending
								if v217 != 0 {
									return int64(0)
								} else {
									F_errmsg(m, int32(_a_F_timestamptz_zone_0), int32(0))
									mBase = m.M
									v221 = m.ExcPending
									if v221 != 0 {
										return int64(0)
									} else {
										F_errfinish(m, int32(_a_F_timestamptz_zone_1), int32(_a_F_timestamptz_zone_2), int32(_a_F_timestamptz_zone_3))
										mBase = m.M
										v226 = m.ExcPending
										if v226 != 0 {
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
							v189 = v181
							m.G0 = v8 + int32(352)
							return v189
						}
					case 1:
						v40 = *(*int32)(unsafe.Add(mBase, uint32(v8)+72))
						v43 = F_DetermineTimeZoneAbbrevOffsetTS(m, v15, v8+int32(80), v40, v8+int32(28))
						mBase = m.M
						v44 = m.ExcPending
						if v44 != 0 {
							return int64(0)
						} else {
							v181 = base.I64_extend_i32_s(v43)*int64(-1000000) + v15
							if base.Ui64(int64(-9011559254509551616)) <= base.Ui64(v181+int64(211813488000000000)) {
								F_errstart_cold(m, int32(21), int32(0))
								mBase = m.M
								v214 = m.ExcPending
								if v214 != 0 {
									return int64(0)
								} else {
									F_errcode(m, int32(134217858))
									mBase = m.M
									v217 = m.ExcPending
									if v217 != 0 {
										return int64(0)
									} else {
										F_errmsg(m, int32(_a_F_timestamptz_zone_0), int32(0))
										mBase = m.M
										v221 = m.ExcPending
										if v221 != 0 {
											return int64(0)
										} else {
											F_errfinish(m, int32(_a_F_timestamptz_zone_1), int32(_a_F_timestamptz_zone_2), int32(_a_F_timestamptz_zone_3))
											mBase = m.M
											v226 = m.ExcPending
											if v226 != 0 {
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
								v189 = v181
								m.G0 = v8 + int32(352)
								return v189
							}
						}
					default:
						v56 = *(*int32)(unsafe.Add(mBase, uint32(v8)+72))
						v57 = F_timestamp2tm(m, v15, v8+int32(348), v8+int32(28), v8+int32(24), int32(0), v56)
						mBase = m.M
						v58 = m.ExcPending
						if v58 != 0 {
							return int64(0)
						} else {
							if v57 != 0 {
								F_errstart_cold(m, int32(21), int32(0))
								mBase = m.M
								v198 = m.ExcPending
								if v198 != 0 {
									return int64(0)
								} else {
									F_errcode(m, int32(134217858))
									mBase = m.M
									v201 = m.ExcPending
									if v201 != 0 {
										return int64(0)
									} else {
										F_errmsg(m, int32(_a_F_timestamptz_zone_0), int32(0))
										mBase = m.M
										v205 = m.ExcPending
										if v205 != 0 {
											return int64(0)
										} else {
											F_errfinish(m, int32(_a_F_timestamptz_zone_1), int32(_a_F_timestamptz_zone_4), int32(_a_F_timestamptz_zone_3))
											mBase = m.M
											v210 = m.ExcPending
											if v210 != 0 {
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
								v59 = int64(*(*int32)(unsafe.Add(mBase, uint32(v8)+24)))
								v60 = *(*int32)(unsafe.Add(mBase, uint32(v8)+48))
								if v60 <= int32(-4713) {
									if v60 != int32(-4713) {
										F_errstart_cold(m, int32(21), int32(0))
										mBase = m.M
										v233 = m.ExcPending
										if v233 != 0 {
											return int64(0)
										} else {
											F_errcode(m, int32(134217858))
											mBase = m.M
											v236 = m.ExcPending
											if v236 != 0 {
												return int64(0)
											} else {
												F_errmsg(m, int32(_a_F_timestamptz_zone_0), int32(0))
												mBase = m.M
												v240 = m.ExcPending
												if v240 != 0 {
													return int64(0)
												} else {
													F_errfinish(m, int32(_a_F_timestamptz_zone_1), int32(_a_F_timestamptz_zone_5), int32(_a_F_timestamptz_zone_3))
													mBase = m.M
													v245 = m.ExcPending
													if v245 != 0 {
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
										v65 = *(*int32)(unsafe.Add(mBase, uint32(v8)+44))
										if int32(10) < v65 {
											v76 = v65
											v77 = *(*int32)(unsafe.Add(mBase, uint32(v8)+40))
											v82 = base.B2i32(int32(2) < v76)
											if int32(2) < v76 {
												v83 = int32(_a_F_timestamptz_zone_6)
											} else {
												v83 = int32(_a_F_timestamptz_zone_7)
											}
											v84 = v83 + v60
											v89 = base.I32_div_s(v84, int32(4))
											v92 = base.I32_div_s(v84, int32(-100))
											v95 = base.I32_div_s(v84, int32(400))
											if int32(2) < v76 {
												v99 = int32(1)
											} else {
												v99 = int32(13)
											}
											v104 = base.I32_div_s((v99+v76)*int32(_a_F_timestamptz_zone_8), int32(256))
											v110 = base.I64_extend_i32_s(v77 + v84*int32(365) + v89 + v92 + v95 + v104 - int32(_a_F_timestamptz_zone_9) - int32(_a_F_timestamptz_zone_10))
											v119 = int64(32)
											v120 = int64(20)
											v122 = int64(base.Ui64(v110) >> (uint(v119) % 64))
											v125 = int64(4294967295)
											v126 = int64(500654080)
											v128 = v110 & v125
											v129 = v126 * v128
											v133 = int64(base.Ui64(v129)>>(uint(v119)%64)) + v126*v122
											v140 = v128*v120 + v133&v125
											*(*int64)(unsafe.Add(mBase, uint32(v8)+8)) = v110*int64(0) + v110>>(uint(int64(63))%64)*int64(86400000000) + v120*v122 + int64(base.Ui64(v133)>>(uint(v119)%64)) + int64(base.Ui64(v140)>>(uint(v119)%64))
											*(*int64)(unsafe.Add(mBase, uint32(v8))) = v129&v125 | v140<<(uint(v119)%64)
											v151 = *(*int64)(unsafe.Add(mBase, uint32(v8)+8))
											v152 = *(*int64)(unsafe.Add(mBase, uint32(v8)))
											if v151 != v152>>(uint(int64(63))%64) {
												F_errstart_cold(m, int32(21), int32(0))
												mBase = m.M
												v233 = m.ExcPending
												if v233 != 0 {
													return int64(0)
												} else {
													F_errcode(m, int32(134217858))
													mBase = m.M
													v236 = m.ExcPending
													if v236 != 0 {
														return int64(0)
													} else {
														F_errmsg(m, int32(_a_F_timestamptz_zone_0), int32(0))
														mBase = m.M
														v240 = m.ExcPending
														if v240 != 0 {
															return int64(0)
														} else {
															F_errfinish(m, int32(_a_F_timestamptz_zone_1), int32(_a_F_timestamptz_zone_5), int32(_a_F_timestamptz_zone_3))
															mBase = m.M
															v245 = m.ExcPending
															if v245 != 0 {
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
												v156 = *(*int32)(unsafe.Add(mBase, uint32(v8)+28))
												v157 = *(*int32)(unsafe.Add(mBase, uint32(v8)+32))
												v158 = *(*int32)(unsafe.Add(mBase, uint32(v8)+36))
												v159 = int32(60)
												v168 = base.I64_extend_i32_s(v156+(v157+v158*v159)*v159)*int64(1000000) + v59
												v171 = v152 + v168
												if base.B2i32(v168 < int64(0))^base.B2i32(v171 < v152)|base.B2i32(base.Ui64(v171-int64(9223371331200000000)) <= base.Ui64(int64(9011559254509551615))) != 0 {
													F_errstart_cold(m, int32(21), int32(0))
													mBase = m.M
													v233 = m.ExcPending
													if v233 != 0 {
														return int64(0)
													} else {
														F_errcode(m, int32(134217858))
														mBase = m.M
														v236 = m.ExcPending
														if v236 != 0 {
															return int64(0)
														} else {
															F_errmsg(m, int32(_a_F_timestamptz_zone_0), int32(0))
															mBase = m.M
															v240 = m.ExcPending
															if v240 != 0 {
																return int64(0)
															} else {
																F_errfinish(m, int32(_a_F_timestamptz_zone_1), int32(_a_F_timestamptz_zone_5), int32(_a_F_timestamptz_zone_3))
																mBase = m.M
																v245 = m.ExcPending
																if v245 != 0 {
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
													v181 = v171
													if base.Ui64(int64(-9011559254509551616)) <= base.Ui64(v181+int64(211813488000000000)) {
														F_errstart_cold(m, int32(21), int32(0))
														mBase = m.M
														v214 = m.ExcPending
														if v214 != 0 {
															return int64(0)
														} else {
															F_errcode(m, int32(134217858))
															mBase = m.M
															v217 = m.ExcPending
															if v217 != 0 {
																return int64(0)
															} else {
																F_errmsg(m, int32(_a_F_timestamptz_zone_0), int32(0))
																mBase = m.M
																v221 = m.ExcPending
																if v221 != 0 {
																	return int64(0)
																} else {
																	F_errfinish(m, int32(_a_F_timestamptz_zone_1), int32(_a_F_timestamptz_zone_2), int32(_a_F_timestamptz_zone_3))
																	mBase = m.M
																	v226 = m.ExcPending
																	if v226 != 0 {
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
														v189 = v181
														m.G0 = v8 + int32(352)
														return v189
													}
												}
											}
										} else {
											F_errstart_cold(m, int32(21), int32(0))
											mBase = m.M
											v233 = m.ExcPending
											if v233 != 0 {
												return int64(0)
											} else {
												F_errcode(m, int32(134217858))
												mBase = m.M
												v236 = m.ExcPending
												if v236 != 0 {
													return int64(0)
												} else {
													F_errmsg(m, int32(_a_F_timestamptz_zone_0), int32(0))
													mBase = m.M
													v240 = m.ExcPending
													if v240 != 0 {
														return int64(0)
													} else {
														F_errfinish(m, int32(_a_F_timestamptz_zone_1), int32(_a_F_timestamptz_zone_5), int32(_a_F_timestamptz_zone_3))
														mBase = m.M
														v245 = m.ExcPending
														if v245 != 0 {
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
									if v60 <= int32(_a_F_timestamptz_zone_11) {
										v70 = *(*int32)(unsafe.Add(mBase, uint32(v8)+44))
										v76 = v70
										v77 = *(*int32)(unsafe.Add(mBase, uint32(v8)+40))
										v82 = base.B2i32(int32(2) < v76)
										if int32(2) < v76 {
											v83 = int32(_a_F_timestamptz_zone_6)
										} else {
											v83 = int32(_a_F_timestamptz_zone_7)
										}
										v84 = v83 + v60
										v89 = base.I32_div_s(v84, int32(4))
										v92 = base.I32_div_s(v84, int32(-100))
										v95 = base.I32_div_s(v84, int32(400))
										if int32(2) < v76 {
											v99 = int32(1)
										} else {
											v99 = int32(13)
										}
										v104 = base.I32_div_s((v99+v76)*int32(_a_F_timestamptz_zone_8), int32(256))
										v110 = base.I64_extend_i32_s(v77 + v84*int32(365) + v89 + v92 + v95 + v104 - int32(_a_F_timestamptz_zone_9) - int32(_a_F_timestamptz_zone_10))
										v119 = int64(32)
										v120 = int64(20)
										v122 = int64(base.Ui64(v110) >> (uint(v119) % 64))
										v125 = int64(4294967295)
										v126 = int64(500654080)
										v128 = v110 & v125
										v129 = v126 * v128
										v133 = int64(base.Ui64(v129)>>(uint(v119)%64)) + v126*v122
										v140 = v128*v120 + v133&v125
										*(*int64)(unsafe.Add(mBase, uint32(v8)+8)) = v110*int64(0) + v110>>(uint(int64(63))%64)*int64(86400000000) + v120*v122 + int64(base.Ui64(v133)>>(uint(v119)%64)) + int64(base.Ui64(v140)>>(uint(v119)%64))
										*(*int64)(unsafe.Add(mBase, uint32(v8))) = v129&v125 | v140<<(uint(v119)%64)
										v151 = *(*int64)(unsafe.Add(mBase, uint32(v8)+8))
										v152 = *(*int64)(unsafe.Add(mBase, uint32(v8)))
										if v151 != v152>>(uint(int64(63))%64) {
											F_errstart_cold(m, int32(21), int32(0))
											mBase = m.M
											v233 = m.ExcPending
											if v233 != 0 {
												return int64(0)
											} else {
												F_errcode(m, int32(134217858))
												mBase = m.M
												v236 = m.ExcPending
												if v236 != 0 {
													return int64(0)
												} else {
													F_errmsg(m, int32(_a_F_timestamptz_zone_0), int32(0))
													mBase = m.M
													v240 = m.ExcPending
													if v240 != 0 {
														return int64(0)
													} else {
														F_errfinish(m, int32(_a_F_timestamptz_zone_1), int32(_a_F_timestamptz_zone_5), int32(_a_F_timestamptz_zone_3))
														mBase = m.M
														v245 = m.ExcPending
														if v245 != 0 {
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
											v156 = *(*int32)(unsafe.Add(mBase, uint32(v8)+28))
											v157 = *(*int32)(unsafe.Add(mBase, uint32(v8)+32))
											v158 = *(*int32)(unsafe.Add(mBase, uint32(v8)+36))
											v159 = int32(60)
											v168 = base.I64_extend_i32_s(v156+(v157+v158*v159)*v159)*int64(1000000) + v59
											v171 = v152 + v168
											if base.B2i32(v168 < int64(0))^base.B2i32(v171 < v152)|base.B2i32(base.Ui64(v171-int64(9223371331200000000)) <= base.Ui64(int64(9011559254509551615))) != 0 {
												F_errstart_cold(m, int32(21), int32(0))
												mBase = m.M
												v233 = m.ExcPending
												if v233 != 0 {
													return int64(0)
												} else {
													F_errcode(m, int32(134217858))
													mBase = m.M
													v236 = m.ExcPending
													if v236 != 0 {
														return int64(0)
													} else {
														F_errmsg(m, int32(_a_F_timestamptz_zone_0), int32(0))
														mBase = m.M
														v240 = m.ExcPending
														if v240 != 0 {
															return int64(0)
														} else {
															F_errfinish(m, int32(_a_F_timestamptz_zone_1), int32(_a_F_timestamptz_zone_5), int32(_a_F_timestamptz_zone_3))
															mBase = m.M
															v245 = m.ExcPending
															if v245 != 0 {
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
												v181 = v171
												if base.Ui64(int64(-9011559254509551616)) <= base.Ui64(v181+int64(211813488000000000)) {
													F_errstart_cold(m, int32(21), int32(0))
													mBase = m.M
													v214 = m.ExcPending
													if v214 != 0 {
														return int64(0)
													} else {
														F_errcode(m, int32(134217858))
														mBase = m.M
														v217 = m.ExcPending
														if v217 != 0 {
															return int64(0)
														} else {
															F_errmsg(m, int32(_a_F_timestamptz_zone_0), int32(0))
															mBase = m.M
															v221 = m.ExcPending
															if v221 != 0 {
																return int64(0)
															} else {
																F_errfinish(m, int32(_a_F_timestamptz_zone_1), int32(_a_F_timestamptz_zone_2), int32(_a_F_timestamptz_zone_3))
																mBase = m.M
																v226 = m.ExcPending
																if v226 != 0 {
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
													v189 = v181
													m.G0 = v8 + int32(352)
													return v189
												}
											}
										}
									} else {
										if v60 != int32(_a_F_timestamptz_zone_12) {
											F_errstart_cold(m, int32(21), int32(0))
											mBase = m.M
											v233 = m.ExcPending
											if v233 != 0 {
												return int64(0)
											} else {
												F_errcode(m, int32(134217858))
												mBase = m.M
												v236 = m.ExcPending
												if v236 != 0 {
													return int64(0)
												} else {
													F_errmsg(m, int32(_a_F_timestamptz_zone_0), int32(0))
													mBase = m.M
													v240 = m.ExcPending
													if v240 != 0 {
														return int64(0)
													} else {
														F_errfinish(m, int32(_a_F_timestamptz_zone_1), int32(_a_F_timestamptz_zone_5), int32(_a_F_timestamptz_zone_3))
														mBase = m.M
														v245 = m.ExcPending
														if v245 != 0 {
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
											v73 = *(*int32)(unsafe.Add(mBase, uint32(v8)+44))
											if int32(5) < v73 {
												F_errstart_cold(m, int32(21), int32(0))
												mBase = m.M
												v233 = m.ExcPending
												if v233 != 0 {
													return int64(0)
												} else {
													F_errcode(m, int32(134217858))
													mBase = m.M
													v236 = m.ExcPending
													if v236 != 0 {
														return int64(0)
													} else {
														F_errmsg(m, int32(_a_F_timestamptz_zone_0), int32(0))
														mBase = m.M
														v240 = m.ExcPending
														if v240 != 0 {
															return int64(0)
														} else {
															F_errfinish(m, int32(_a_F_timestamptz_zone_1), int32(_a_F_timestamptz_zone_5), int32(_a_F_timestamptz_zone_3))
															mBase = m.M
															v245 = m.ExcPending
															if v245 != 0 {
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
												v76 = v73
												v77 = *(*int32)(unsafe.Add(mBase, uint32(v8)+40))
												v82 = base.B2i32(int32(2) < v76)
												if int32(2) < v76 {
													v83 = int32(_a_F_timestamptz_zone_6)
												} else {
													v83 = int32(_a_F_timestamptz_zone_7)
												}
												v84 = v83 + v60
												v89 = base.I32_div_s(v84, int32(4))
												v92 = base.I32_div_s(v84, int32(-100))
												v95 = base.I32_div_s(v84, int32(400))
												if int32(2) < v76 {
													v99 = int32(1)
												} else {
													v99 = int32(13)
												}
												v104 = base.I32_div_s((v99+v76)*int32(_a_F_timestamptz_zone_8), int32(256))
												v110 = base.I64_extend_i32_s(v77 + v84*int32(365) + v89 + v92 + v95 + v104 - int32(_a_F_timestamptz_zone_9) - int32(_a_F_timestamptz_zone_10))
												v119 = int64(32)
												v120 = int64(20)
												v122 = int64(base.Ui64(v110) >> (uint(v119) % 64))
												v125 = int64(4294967295)
												v126 = int64(500654080)
												v128 = v110 & v125
												v129 = v126 * v128
												v133 = int64(base.Ui64(v129)>>(uint(v119)%64)) + v126*v122
												v140 = v128*v120 + v133&v125
												*(*int64)(unsafe.Add(mBase, uint32(v8)+8)) = v110*int64(0) + v110>>(uint(int64(63))%64)*int64(86400000000) + v120*v122 + int64(base.Ui64(v133)>>(uint(v119)%64)) + int64(base.Ui64(v140)>>(uint(v119)%64))
												*(*int64)(unsafe.Add(mBase, uint32(v8))) = v129&v125 | v140<<(uint(v119)%64)
												v151 = *(*int64)(unsafe.Add(mBase, uint32(v8)+8))
												v152 = *(*int64)(unsafe.Add(mBase, uint32(v8)))
												if v151 != v152>>(uint(int64(63))%64) {
													F_errstart_cold(m, int32(21), int32(0))
													mBase = m.M
													v233 = m.ExcPending
													if v233 != 0 {
														return int64(0)
													} else {
														F_errcode(m, int32(134217858))
														mBase = m.M
														v236 = m.ExcPending
														if v236 != 0 {
															return int64(0)
														} else {
															F_errmsg(m, int32(_a_F_timestamptz_zone_0), int32(0))
															mBase = m.M
															v240 = m.ExcPending
															if v240 != 0 {
																return int64(0)
															} else {
																F_errfinish(m, int32(_a_F_timestamptz_zone_1), int32(_a_F_timestamptz_zone_5), int32(_a_F_timestamptz_zone_3))
																mBase = m.M
																v245 = m.ExcPending
																if v245 != 0 {
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
													v156 = *(*int32)(unsafe.Add(mBase, uint32(v8)+28))
													v157 = *(*int32)(unsafe.Add(mBase, uint32(v8)+32))
													v158 = *(*int32)(unsafe.Add(mBase, uint32(v8)+36))
													v159 = int32(60)
													v168 = base.I64_extend_i32_s(v156+(v157+v158*v159)*v159)*int64(1000000) + v59
													v171 = v152 + v168
													if base.B2i32(v168 < int64(0))^base.B2i32(v171 < v152)|base.B2i32(base.Ui64(v171-int64(9223371331200000000)) <= base.Ui64(int64(9011559254509551615))) != 0 {
														F_errstart_cold(m, int32(21), int32(0))
														mBase = m.M
														v233 = m.ExcPending
														if v233 != 0 {
															return int64(0)
														} else {
															F_errcode(m, int32(134217858))
															mBase = m.M
															v236 = m.ExcPending
															if v236 != 0 {
																return int64(0)
															} else {
																F_errmsg(m, int32(_a_F_timestamptz_zone_0), int32(0))
																mBase = m.M
																v240 = m.ExcPending
																if v240 != 0 {
																	return int64(0)
																} else {
																	F_errfinish(m, int32(_a_F_timestamptz_zone_1), int32(_a_F_timestamptz_zone_5), int32(_a_F_timestamptz_zone_3))
																	mBase = m.M
																	v245 = m.ExcPending
																	if v245 != 0 {
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
														v181 = v171
														if base.Ui64(int64(-9011559254509551616)) <= base.Ui64(v181+int64(211813488000000000)) {
															F_errstart_cold(m, int32(21), int32(0))
															mBase = m.M
															v214 = m.ExcPending
															if v214 != 0 {
																return int64(0)
															} else {
																F_errcode(m, int32(134217858))
																mBase = m.M
																v217 = m.ExcPending
																if v217 != 0 {
																	return int64(0)
																} else {
																	F_errmsg(m, int32(_a_F_timestamptz_zone_0), int32(0))
																	mBase = m.M
																	v221 = m.ExcPending
																	if v221 != 0 {
																		return int64(0)
																	} else {
																		F_errfinish(m, int32(_a_F_timestamptz_zone_1), int32(_a_F_timestamptz_zone_2), int32(_a_F_timestamptz_zone_3))
																		mBase = m.M
																		v226 = m.ExcPending
																		if v226 != 0 {
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
															v189 = v181
															m.G0 = v8 + int32(352)
															return v189
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
			v189 = v15
			m.G0 = v8 + int32(352)
			return v189
		}
	}
}
