package p5

import (
	base "github.com/shibukawa/pgmem/internal/aot/pgaot/base"
	"unsafe"
)

func F_TimestampDifference(m *base.Module, l0 int64, l1 int64, l2 int32, l3 int32) {
	mBase := m.M
	_ = mBase
	var v7 int64
	_ = v7
	var v11 int64
	_ = v11
	var v12 int64
	_ = v12
	var v19 int32
	_ = v19
	var v20 int32
	_ = v20
	v7 = l1 - l0
	if v7 <= int64(0) {
		v19 = int32(0)
		v20 = int32(0)
	} else {
		v11 = int64(1000000)
		v12 = base.I64_div_u_s(v7, v11)
		v19 = base.I32_wrap_i64(v12)
		v20 = base.I32_wrap_i64(v7 - v12*v11)
	}
	*(*int32)(unsafe.Add(mBase, uint32(l2))) = v19
	*(*int32)(unsafe.Add(mBase, uint32(l3))) = v20
	return
}
func F_TimestampDifferenceMilliseconds(m *base.Module, l0 int64, l1 int64) int32 {
	var v8 int64
	_ = v8
	var v17 int64
	_ = v17
	var v20 int32
	_ = v20
	if l1 <= l0 {
		v20 = int32(0)
	} else {
		v8 = l1 - l0
		if base.B2i32(int64(0) < l0)^base.B2i32(v8 < l1)|base.B2i32(int64(2147483646000) < v8) != 0 {
			v20 = int32(2147483647)
		} else {
			v17 = base.I64_div_s(v8+int64(999), int64(1000))
			v20 = base.I32_wrap_i64(v17)
		}
	}
	return v20
}
func F_timestamp_gt_timestamptz(m *base.Module, l0 int32) int32 {
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
	v9 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
	v10 = *(*int64)(unsafe.Add(mBase, uint32(v9)))
	v11 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
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
				v29 = base.B2i32(v10 < v15)
			}
		}
		m.G0 = v7 + int32(16)
		return v29
	}
}
func F_timestamp_le_date(m *base.Module, l0 int32) int32 {
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
	return int32(base.Ui32(v26^int32(-1)) >> (uint(int32(31)) % 32))
}
func F_timestamp_ne_timestamptz(m *base.Module, l0 int32) int32 {
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
	var v19 int32
	_ = v19
	v5 = m.G0
	v7 = v5 - int32(16)
	m.G0 = v7
	v9 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
	v10 = *(*int64)(unsafe.Add(mBase, uint32(v9)))
	v11 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v12 = *(*int64)(unsafe.Add(mBase, uint32(v11)))
	v15 = F_timestamp2timestamptz_opt_overflow(m, v12, v7+int32(12))
	mBase = m.M
	v18 = m.ExcPending
	if v18 != 0 {
		return int32(0)
	} else {
		v19 = *(*int32)(unsafe.Add(mBase, uint32(v7)+12))
		m.G0 = v7 + int32(16)
		return base.B2i32(v19 != int32(0)) | base.B2i32(v10 != v15)
	}
}
func F_timestamp_pl_interval(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v11 int32
	_ = v11
	var v13 int32
	_ = v13
	var v15 int32
	_ = v15
	var v16 int64
	_ = v16
	var v17 int32
	_ = v17
	var v18 int32
	_ = v18
	var v23 int32
	_ = v23
	var v26 int64
	_ = v26
	var v27 int64
	_ = v27
	var v37 int32
	_ = v37
	var v40 int32
	_ = v40
	var v44 int32
	_ = v44
	var v49 int32
	_ = v49
	var v50 int32
	_ = v50
	var v53 int64
	_ = v53
	var v54 int64
	_ = v54
	var v62 int32
	_ = v62
	var v65 int32
	_ = v65
	var v69 int32
	_ = v69
	var v74 int32
	_ = v74
	var v81 int64
	_ = v81
	var v89 int64
	_ = v89
	var v90 int64
	_ = v90
	var v93 int64
	_ = v93
	var v96 int32
	_ = v96
	var v108 int32
	_ = v108
	var v109 int32
	_ = v109
	var v110 int32
	_ = v110
	var v111 int32
	_ = v111
	var v117 int32
	_ = v117
	var v122 int32
	_ = v122
	var v125 int32
	_ = v125
	var v126 int32
	_ = v126
	var v127 int32
	_ = v127
	var v130 int32
	_ = v130
	var v132 int32
	_ = v132
	var v138 int32
	_ = v138
	var v142 int32
	_ = v142
	var v143 int32
	_ = v143
	var v145 int32
	_ = v145
	var v153 int32
	_ = v153
	var v157 int32
	_ = v157
	var v167 int32
	_ = v167
	var v177 int64
	_ = v177
	var v179 int64
	_ = v179
	var v184 int64
	_ = v184
	var v186 int64
	_ = v186
	var v191 int64
	_ = v191
	var v193 int64
	_ = v193
	var v195 int32
	_ = v195
	var v197 int32
	_ = v197
	var v198 int32
	_ = v198
	var v199 int32
	_ = v199
	var v207 int32
	_ = v207
	var v208 int32
	_ = v208
	var v209 int32
	_ = v209
	var v210 int32
	_ = v210
	var v211 int32
	_ = v211
	var v212 int32
	_ = v212
	var v219 int32
	_ = v219
	var v224 int32
	_ = v224
	var v225 int32
	_ = v225
	var v228 int32
	_ = v228
	var v235 int32
	_ = v235
	var v238 int32
	_ = v238
	var v240 int32
	_ = v240
	var v241 int32
	_ = v241
	var v243 int32
	_ = v243
	var v249 int32
	_ = v249
	var v253 int32
	_ = v253
	var v255 int32
	_ = v255
	var v263 int32
	_ = v263
	var v269 int32
	_ = v269
	var v276 int32
	_ = v276
	var v280 int32
	_ = v280
	var v288 int32
	_ = v288
	var v290 int32
	_ = v290
	var v306 int32
	_ = v306
	var v311 int32
	_ = v311
	var v312 int32
	_ = v312
	var v313 int32
	_ = v313
	var v318 int32
	_ = v318
	var v321 int32
	_ = v321
	var v324 int32
	_ = v324
	var v328 int32
	_ = v328
	var v333 int32
	_ = v333
	var v339 int64
	_ = v339
	var v348 int64
	_ = v348
	var v349 int64
	_ = v349
	var v351 int64
	_ = v351
	var v354 int64
	_ = v354
	var v355 int64
	_ = v355
	var v357 int64
	_ = v357
	var v358 int64
	_ = v358
	var v362 int64
	_ = v362
	var v369 int64
	_ = v369
	var v380 int64
	_ = v380
	var v381 int64
	_ = v381
	var v389 int32
	_ = v389
	var v390 int32
	_ = v390
	var v391 int32
	_ = v391
	var v392 int32
	_ = v392
	var v401 int64
	_ = v401
	var v404 int64
	_ = v404
	var v416 int64
	_ = v416
	var v419 int32
	_ = v419
	var v421 int64
	_ = v421
	var v429 int64
	_ = v429
	var v430 int64
	_ = v430
	var v433 int64
	_ = v433
	var v436 int32
	_ = v436
	var v440 int32
	_ = v440
	var v442 int32
	_ = v442
	var v444 int32
	_ = v444
	var v448 int32
	_ = v448
	var v449 int32
	_ = v449
	var v450 int32
	_ = v450
	var v451 int32
	_ = v451
	var v457 int32
	_ = v457
	var v462 int32
	_ = v462
	var v465 int32
	_ = v465
	var v466 int32
	_ = v466
	var v467 int32
	_ = v467
	var v470 int32
	_ = v470
	var v472 int32
	_ = v472
	var v478 int32
	_ = v478
	var v482 int32
	_ = v482
	var v483 int32
	_ = v483
	var v485 int32
	_ = v485
	var v493 int32
	_ = v493
	var v497 int32
	_ = v497
	var v507 int32
	_ = v507
	var v517 int64
	_ = v517
	var v519 int64
	_ = v519
	var v524 int64
	_ = v524
	var v526 int64
	_ = v526
	var v531 int64
	_ = v531
	var v533 int64
	_ = v533
	var v537 int32
	_ = v537
	var v538 int32
	_ = v538
	var v539 int32
	_ = v539
	var v544 int32
	_ = v544
	var v545 int32
	_ = v545
	var v546 int32
	_ = v546
	var v551 int32
	_ = v551
	var v554 int32
	_ = v554
	var v557 int32
	_ = v557
	var v561 int32
	_ = v561
	var v566 int32
	_ = v566
	var v569 int32
	_ = v569
	var v570 int32
	_ = v570
	var v571 int32
	_ = v571
	var v573 int32
	_ = v573
	var v582 int32
	_ = v582
	var v583 int32
	_ = v583
	var v584 int32
	_ = v584
	var v585 int32
	_ = v585
	var v591 int32
	_ = v591
	var v596 int32
	_ = v596
	var v599 int32
	_ = v599
	var v600 int32
	_ = v600
	var v601 int32
	_ = v601
	var v604 int32
	_ = v604
	var v606 int32
	_ = v606
	var v612 int32
	_ = v612
	var v616 int32
	_ = v616
	var v617 int32
	_ = v617
	var v619 int32
	_ = v619
	var v627 int32
	_ = v627
	var v631 int32
	_ = v631
	var v641 int32
	_ = v641
	var v645 int32
	_ = v645
	var v650 int32
	_ = v650
	var v655 int32
	_ = v655
	var v658 int32
	_ = v658
	var v661 int32
	_ = v661
	var v662 int32
	_ = v662
	var v667 int32
	_ = v667
	var v668 int32
	_ = v668
	var v669 int32
	_ = v669
	var v674 int32
	_ = v674
	var v677 int32
	_ = v677
	var v680 int32
	_ = v680
	var v684 int32
	_ = v684
	var v689 int32
	_ = v689
	var v695 int64
	_ = v695
	var v704 int64
	_ = v704
	var v705 int64
	_ = v705
	var v707 int64
	_ = v707
	var v710 int64
	_ = v710
	var v711 int64
	_ = v711
	var v713 int64
	_ = v713
	var v714 int64
	_ = v714
	var v718 int64
	_ = v718
	var v725 int64
	_ = v725
	var v736 int64
	_ = v736
	var v737 int64
	_ = v737
	var v745 int32
	_ = v745
	var v746 int32
	_ = v746
	var v747 int32
	_ = v747
	var v748 int32
	_ = v748
	var v757 int64
	_ = v757
	var v760 int64
	_ = v760
	var v773 int64
	_ = v773
	var v776 int64
	_ = v776
	var v779 int64
	_ = v779
	var v792 int64
	_ = v792
	var v794 int32
	_ = v794
	var v795 int32
	_ = v795
	var v803 int32
	_ = v803
	var v806 int32
	_ = v806
	var v810 int32
	_ = v810
	var v815 int32
	_ = v815
	var v819 int32
	_ = v819
	var v822 int32
	_ = v822
	var v826 int32
	_ = v826
	var v831 int32
	_ = v831
	var v835 int32
	_ = v835
	var v838 int32
	_ = v838
	var v842 int32
	_ = v842
	var v847 int32
	_ = v847
	var v851 int32
	_ = v851
	var v854 int32
	_ = v854
	var v858 int32
	_ = v858
	var v863 int32
	_ = v863
	var v867 int32
	_ = v867
	var v870 int32
	_ = v870
	var v874 int32
	_ = v874
	var v879 int32
	_ = v879
	var v883 int32
	_ = v883
	var v886 int32
	_ = v886
	var v890 int32
	_ = v890
	var v895 int32
	_ = v895
	var v902 int32
	_ = v902
	var v905 int32
	_ = v905
	var v909 int32
	_ = v909
	var v914 int32
	_ = v914
	var v920 int32
	_ = v920
	var v923 int32
	_ = v923
	var v927 int32
	_ = v927
	var v932 int32
	_ = v932
	v11 = m.G0
	v13 = v11 - int32(80)
	m.G0 = v13
	v15 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v16 = *(*int64)(unsafe.Add(mBase, uint32(v15)))
	v17 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
	v18 = *(*int32)(unsafe.Add(mBase, uint32(v17)+12))
	if v18 != int32(2147483647) {
		goto L11
	} else {
		goto L12
	}
L1:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v920 = m.ExcPending
	if v920 != 0 {
		goto L18
	} else {
		goto L164
	}
L2:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v902 = m.ExcPending
	if v902 != 0 {
		goto L18
	} else {
		goto L160
	}
L3:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v883 = m.ExcPending
	if v883 != 0 {
		goto L18
	} else {
		goto L156
	}
L4:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v867 = m.ExcPending
	if v867 != 0 {
		goto L18
	} else {
		goto L152
	}
L5:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v851 = m.ExcPending
	if v851 != 0 {
		goto L18
	} else {
		goto L148
	}
L6:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v835 = m.ExcPending
	if v835 != 0 {
		goto L18
	} else {
		goto L144
	}
L7:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v819 = m.ExcPending
	if v819 != 0 {
		goto L18
	} else {
		goto L140
	}
L8:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v803 = m.ExcPending
	if v803 != 0 {
		goto L18
	} else {
		goto L136
	}
L9:
	;
	v794 = F_Int64GetDatum(m, v792)
	mBase = m.M
	v795 = m.ExcPending
	if v795 != 0 {
		goto L18
	} else {
		goto L135
	}
L10:
	;
	if base.Ui64(v16-int64(9223372036854775807)) < base.Ui64(int64(2)) {
		goto L30
	} else {
		goto L31
	}
L11:
	;
	if v18 != int32(-2147483648) {
		goto L10
	} else {
		goto L14
	}
L12:
	;
	goto L13
L13:
	;
	v50 = *(*int32)(unsafe.Add(mBase, uint32(v17)+8))
	if v50 != int32(2147483647) {
		goto L10
	} else {
		goto L23
	}
L14:
	;
	v23 = *(*int32)(unsafe.Add(mBase, uint32(v17)+8))
	if v23 != int32(-2147483648) {
		goto L10
	} else {
		goto L15
	}
L15:
	;
	v26 = int64(-9223372036854775807 - 1)
	v27 = *(*int64)(unsafe.Add(mBase, uint32(v17)))
	if v27 != v26 {
		goto L10
	} else {
		goto L16
	}
L16:
	;
	if v16 != int64(9223372036854775807) {
		v792 = v26
		goto L9
	} else {
		goto L17
	}
L17:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v37 = m.ExcPending
	if v37 != 0 {
		goto L18
	} else {
		goto L19
	}
L18:
	;
	return int32(0)
L19:
	;
	F_errcode(m, int32(134217858))
	mBase = m.M
	v40 = m.ExcPending
	if v40 != 0 {
		goto L18
	} else {
		goto L20
	}
L20:
	;
	F_errmsg(m, int32(_a_F_timestamp_pl_interval_0), int32(0))
	mBase = m.M
	v44 = m.ExcPending
	if v44 != 0 {
		goto L18
	} else {
		goto L21
	}
L21:
	;
	F_errfinish(m, int32(_a_F_timestamp_pl_interval_1), int32(3125), int32(_a_F_timestamp_pl_interval_2))
	mBase = m.M
	v49 = m.ExcPending
	if v49 != 0 {
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
	v53 = int64(9223372036854775807)
	v54 = *(*int64)(unsafe.Add(mBase, uint32(v17)))
	if v54 != v53 {
		goto L10
	} else {
		goto L24
	}
L24:
	;
	if v16 != int64(-9223372036854775807-1) {
		v792 = v53
		goto L9
	} else {
		goto L25
	}
L25:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v62 = m.ExcPending
	if v62 != 0 {
		goto L18
	} else {
		goto L26
	}
L26:
	;
	F_errcode(m, int32(134217858))
	mBase = m.M
	v65 = m.ExcPending
	if v65 != 0 {
		goto L18
	} else {
		goto L27
	}
L27:
	;
	F_errmsg(m, int32(_a_F_timestamp_pl_interval_0), int32(0))
	mBase = m.M
	v69 = m.ExcPending
	if v69 != 0 {
		goto L18
	} else {
		goto L28
	}
L28:
	;
	F_errfinish(m, int32(_a_F_timestamp_pl_interval_1), int32(3134), int32(_a_F_timestamp_pl_interval_2))
	mBase = m.M
	v74 = m.ExcPending
	if v74 != 0 {
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
	v792 = v16
	goto L9
L31:
	;
	goto L32
L32:
	;
	if v18 != 0 {
		goto L33
	} else {
		goto L34
	}
L33:
	;
	v81 = base.I64_div_s(v16, int64(86400000000))
	if base.Ui64(int64(172799999999)) <= base.Ui64(v16+int64(86399999999)) {
		goto L36
	} else {
		goto L37
	}
L34:
	;
	v416 = v16
	goto L35
L35:
	;
	v419 = *(*int32)(unsafe.Add(mBase, uint32(v17)+8))
	if v419 != 0 {
		goto L84
	} else {
		goto L85
	}
L36:
	;
	v89 = v81 * int64(-86400000000)
	goto L38
L37:
	;
	v89 = int64(0)
	goto L38
L38:
	;
	v90 = v89 + v16
	v93 = v90>>(uint(int64(63))%64) + v81
	if v93 <= int64(-2451546) {
		goto L8
	} else {
		goto L39
	}
L39:
	;
	v96 = base.I32_wrap_i64(v93)
	v108 = v96 + int32(_a_F_timestamp_pl_interval_3)
	v109 = int32(_a_F_timestamp_pl_interval_4)
	v110 = base.I32_div_u_s(v108, v109)
	v111 = int32(3)
	v117 = int32(2)
	v122 = base.I32_div_u_s((v110*int32(1073595727)+v108)<<(uint(v117)%32)|v111, v109)
	v125 = v96 + int32(_a_F_timestamp_pl_interval_5) + v110*v111 + v122 + int32(_a_F_timestamp_pl_interval_6)
	v126 = int32(1461)
	v127 = base.I32_div_u_s(v125, v126)
	v130 = v127*int32(-1461) + v125
	v132 = v130 << (uint(v117) % 32)
	if base.Ui32(v126) <= base.Ui32(v132) {
		goto L42
	} else {
		goto L43
	}
L40:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v13)+68)) = int64(4294967295)
	if v90 < int64(0) {
		goto L45
	} else {
		goto L46
	}
L41:
	;
	v145 = base.I32_div_u_s(v132, int32(1461))
	*(*int32)(unsafe.Add(mBase, uint32(v13+int32(56)))) = v145 + v127<<(uint(int32(2))%32) - int32(_a_F_timestamp_pl_interval_7)
	v153 = v143 + int32(123)
	v157 = int32(base.Ui32(v153*int32(2141)) >> (uint(int32(16)) % 32))
	*(*int32)(unsafe.Add(mBase, uint32(v13+int32(48)))) = v153 - int32(base.Ui32(v157*int32(_a_F_timestamp_pl_interval_8))>>(uint(int32(8))%32))
	v167 = base.I32_rem_u_s(v157+int32(10), int32(12))
	*(*int32)(unsafe.Add(mBase, uint32(v13+int32(52)))) = v167 + int32(1)
	goto L40
L42:
	;
	v138 = base.I32_rem_u_s(v130+int32(305), int32(365))
	v143 = v138
	goto L41
L43:
	;
	goto L44
L44:
	;
	v142 = base.I32_rem_u_s(v130+int32(306), int32(366))
	v143 = v142
	goto L41
L45:
	;
	v177 = v90 + int64(86400000000)
	goto L47
L46:
	;
	v177 = v90
	goto L47
L47:
	;
	v179 = base.I64_div_s(v177, int64(3600000000))
	*(*uint32)(unsafe.Add(mBase, uint32(v13)+44)) = uint32(v179)
	v184 = base.I64_extend32_s(v179)*int64(-3600000000) + v177
	v186 = base.I64_div_s(v184, int64(60000000))
	*(*uint32)(unsafe.Add(mBase, uint32(v13)+40)) = uint32(v186)
	v191 = base.I64_extend32_s(v186)*int64(-60000000) + v184
	v193 = base.I64_div_s(v191, int64(1000000))
	*(*uint32)(unsafe.Add(mBase, uint32(v13)+36)) = uint32(v193)
	v195 = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v13)+76)) = v195
	v197 = *(*int32)(unsafe.Add(mBase, uint32(v13)+52))
	v198 = *(*int32)(unsafe.Add(mBase, uint32(v17)+12))
	v199 = v197 + v198
	*(*int32)(unsafe.Add(mBase, uint32(v13)+52)) = v199
	if base.B2i32(v198 < v195) != base.B2i32(v199 < v197) {
		goto L7
	} else {
		goto L48
	}
L48:
	;
	if int32(13) <= v199 {
		goto L51
	} else {
		goto L52
	}
L49:
	;
	v243 = *(*int32)(unsafe.Add(mBase, uint32(v13)+48))
	if v240&int32(3) == int32(0) {
		goto L57
	} else {
		goto L58
	}
L50:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v13)+52)) = v238
	v240 = v235
	v241 = v238
	goto L49
L51:
	;
	v207 = int32(1)
	v208 = v199 - v207
	v209 = int32(12)
	v210 = base.I32_div_u_s(v208, v209)
	v211 = *(*int32)(unsafe.Add(mBase, uint32(v13)+56))
	v212 = v210 + v211
	*(*int32)(unsafe.Add(mBase, uint32(v13)+56)) = v212
	v235 = v212
	v238 = v208 - v210*v209 + v207
	goto L50
L52:
	;
	goto L53
L53:
	;
	v219 = *(*int32)(unsafe.Add(mBase, uint32(v13)+56))
	if int32(0) < v199 {
		v240 = v219
		v241 = v199
		goto L49
	} else {
		goto L54
	}
L54:
	;
	v224 = int32(12)
	v225 = base.I32_div_u_s(int32(0)-v199, v224)
	v228 = v219 + (v225 ^ int32(-1))
	*(*int32)(unsafe.Add(mBase, uint32(v13)+56)) = v228
	v235 = v228
	v238 = v225*v224 + v199 + v224
	goto L50
L55:
	;
	if v240 <= int32(-4713) {
		goto L67
	} else {
		goto L68
	}
L56:
	;
	v288 = *(*int32)(unsafe.Add(mBase, uint32(v280*int32(52)+v241<<(uint(int32(2))%32))+uint32(_c_F_timestamp_pl_interval[0])))
	*(*int32)(unsafe.Add(mBase, uint32(v13)+48)) = v288
	v290 = v288
	goto L55
L57:
	;
	v249 = base.I32_rem_s(v240, int32(100))
	if v249 == int32(0) {
		goto L60
	} else {
		goto L61
	}
L58:
	;
	goto L59
L59:
	;
	v276 = *(*int32)(unsafe.Add(mBase, uint32(v241<<(uint(int32(2))%32))+uint32(_c_F_timestamp_pl_interval[0])))
	if v243 <= v276 {
		v290 = v243
		goto L55
	} else {
		goto L65
	}
L60:
	;
	v253 = base.I32_rem_s(v240, int32(400))
	v255 = base.B2i32(v253 == int32(0))
	v263 = *(*int32)(unsafe.Add(mBase, uint32(v255*int32(52)+v241<<(uint(int32(2))%32))+uint32(_c_F_timestamp_pl_interval[0])))
	if v243 <= v263 {
		v290 = v243
		goto L55
	} else {
		goto L63
	}
L61:
	;
	goto L62
L62:
	;
	v269 = *(*int32)(unsafe.Add(mBase, uint32(v241<<(uint(int32(2))%32))+uint32(_c_F_timestamp_pl_interval[1])))
	if v243 <= v269 {
		v290 = v243
		goto L55
	} else {
		goto L64
	}
L63:
	;
	v280 = v255
	goto L56
L64:
	;
	v280 = int32(1)
	goto L56
L65:
	;
	v280 = int32(0)
	goto L56
L66:
	;
	v306 = v13 + int32(16)
	v311 = base.B2i32(int32(2) < v241)
	if int32(2) < v241 {
		goto L75
	} else {
		goto L76
	}
L67:
	;
	if v240 != int32(-4713) {
		goto L1
	} else {
		goto L70
	}
L68:
	;
	goto L69
L69:
	;
	if v240 < int32(_a_F_timestamp_pl_interval_13) {
		goto L66
	} else {
		goto L72
	}
L70:
	;
	if base.Ui32(int32(10)) < base.Ui32(v241) {
		goto L66
	} else {
		goto L71
	}
L71:
	;
	goto L1
L72:
	;
	if base.B2i32(v240 != int32(_a_F_timestamp_pl_interval_13))|base.B2i32(base.Ui32(int32(5)) < base.Ui32(v241)) != 0 {
		goto L1
	} else {
		goto L73
	}
L73:
	;
	goto L66
L74:
	;
	v339 = base.I64_extend_i32_s(v290 + v313*int32(365) + v318 + v321 + v324 + v333 - int32(_a_F_timestamp_pl_interval_10) - int32(_a_F_timestamp_pl_interval_5))
	v348 = int64(32)
	v349 = int64(20)
	v351 = int64(base.Ui64(v339) >> (uint(v348) % 64))
	v354 = int64(4294967295)
	v355 = int64(500654080)
	v357 = v339 & v354
	v358 = v355 * v357
	v362 = int64(base.Ui64(v358)>>(uint(v348)%64)) + v355*v351
	v369 = v357*v349 + v362&v354
	*(*int64)(unsafe.Add(mBase, uint32(v306)+8)) = v339*int64(0) + v339>>(uint(int64(63))%64)*int64(86400000000) + v349*v351 + int64(base.Ui64(v362)>>(uint(v348)%64)) + int64(base.Ui64(v369)>>(uint(v348)%64))
	*(*int64)(unsafe.Add(mBase, uint32(v306))) = v358&v354 | v369<<(uint(v348)%64)
	goto L81
L75:
	;
	v312 = int32(_a_F_timestamp_pl_interval_7)
	goto L77
L76:
	;
	v312 = int32(_a_F_timestamp_pl_interval_9)
	goto L77
L77:
	;
	v313 = v312 + v240
	v318 = base.I32_div_s(v313, int32(4))
	v321 = base.I32_div_s(v313, int32(-100))
	v324 = base.I32_div_s(v313, int32(400))
	if int32(2) < v241 {
		goto L78
	} else {
		goto L79
	}
L78:
	;
	v328 = int32(1)
	goto L80
L79:
	;
	v328 = int32(13)
	goto L80
L80:
	;
	v333 = base.I32_div_s((v328+v241)*int32(_a_F_timestamp_pl_interval_8), int32(256))
	goto L74
L81:
	;
	v380 = *(*int64)(unsafe.Add(mBase, uint32(v13)+24))
	v381 = *(*int64)(unsafe.Add(mBase, uint32(v13)+16))
	if v380 != v381>>(uint(int64(63))%64) {
		goto L1
	} else {
		goto L82
	}
L82:
	;
	v389 = *(*int32)(unsafe.Add(mBase, uint32(v13)+36))
	v390 = *(*int32)(unsafe.Add(mBase, uint32(v13)+40))
	v391 = *(*int32)(unsafe.Add(mBase, uint32(v13)+44))
	v392 = int32(60)
	v401 = base.I64_extend32_s(v193*int64(4293967296)+v191) + base.I64_extend_i32_s(v389+(v390+v391*v392)*v392)*int64(1000000)
	v404 = v381 + v401
	if base.B2i32(v401 < int64(0))^base.B2i32(v404 < v381)|base.B2i32(base.Ui64(v404-int64(9223371331200000000)) <= base.Ui64(int64(9011559254509551615))) != 0 {
		goto L1
	} else {
		goto L83
	}
L83:
	;
	v416 = v404
	goto L35
L84:
	;
	v421 = base.I64_div_s(v416, int64(86400000000))
	if base.Ui64(int64(172799999999)) <= base.Ui64(v416+int64(86399999999)) {
		goto L87
	} else {
		goto L88
	}
L85:
	;
	v773 = v416
	goto L86
L86:
	;
	v776 = *(*int64)(unsafe.Add(mBase, uint32(v17)))
	v779 = v773 + v776
	if base.B2i32(v776 < int64(0)) != base.B2i32(v779 < v773) {
		goto L4
	} else {
		goto L133
	}
L87:
	;
	v429 = v421 * int64(-86400000000)
	goto L89
L88:
	;
	v429 = int64(0)
	goto L89
L89:
	;
	v430 = v429 + v416
	v433 = v430>>(uint(int64(63))%64) + v421
	if v433 <= int64(-2451546) {
		goto L6
	} else {
		goto L90
	}
L90:
	;
	v436 = base.I32_wrap_i64(v433)
	v440 = v13 + int32(56)
	v442 = v13 + int32(52)
	v444 = v13 + int32(48)
	v448 = v436 + int32(_a_F_timestamp_pl_interval_3)
	v449 = int32(_a_F_timestamp_pl_interval_4)
	v450 = base.I32_div_u_s(v448, v449)
	v451 = int32(3)
	v457 = int32(2)
	v462 = base.I32_div_u_s((v450*int32(1073595727)+v448)<<(uint(v457)%32)|v451, v449)
	v465 = v436 + int32(_a_F_timestamp_pl_interval_5) + v450*v451 + v462 + int32(_a_F_timestamp_pl_interval_6)
	v466 = int32(1461)
	v467 = base.I32_div_u_s(v465, v466)
	v470 = v467*int32(-1461) + v465
	v472 = v470 << (uint(v457) % 32)
	if base.Ui32(v466) <= base.Ui32(v472) {
		goto L93
	} else {
		goto L94
	}
L91:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v13)+68)) = int64(4294967295)
	if v430 < int64(0) {
		goto L96
	} else {
		goto L97
	}
L92:
	;
	v485 = base.I32_div_u_s(v472, int32(1461))
	*(*int32)(unsafe.Add(mBase, uint32(v440))) = v485 + v467<<(uint(int32(2))%32) - int32(_a_F_timestamp_pl_interval_7)
	v493 = v483 + int32(123)
	v497 = int32(base.Ui32(v493*int32(2141)) >> (uint(int32(16)) % 32))
	*(*int32)(unsafe.Add(mBase, uint32(v444))) = v493 - int32(base.Ui32(v497*int32(_a_F_timestamp_pl_interval_8))>>(uint(int32(8))%32))
	v507 = base.I32_rem_u_s(v497+int32(10), int32(12))
	*(*int32)(unsafe.Add(mBase, uint32(v442))) = v507 + int32(1)
	goto L91
L93:
	;
	v478 = base.I32_rem_u_s(v470+int32(305), int32(365))
	v483 = v478
	goto L92
L94:
	;
	goto L95
L95:
	;
	v482 = base.I32_rem_u_s(v470+int32(306), int32(366))
	v483 = v482
	goto L92
L96:
	;
	v517 = v430 + int64(86400000000)
	goto L98
L97:
	;
	v517 = v430
	goto L98
L98:
	;
	v519 = base.I64_div_s(v517, int64(3600000000))
	*(*uint32)(unsafe.Add(mBase, uint32(v13)+44)) = uint32(v519)
	v524 = base.I64_extend32_s(v519)*int64(-3600000000) + v517
	v526 = base.I64_div_s(v524, int64(60000000))
	*(*uint32)(unsafe.Add(mBase, uint32(v13)+40)) = uint32(v526)
	v531 = base.I64_extend32_s(v526)*int64(-60000000) + v524
	v533 = base.I64_div_s(v531, int64(1000000))
	*(*uint32)(unsafe.Add(mBase, uint32(v13)+36)) = uint32(v533)
	*(*int32)(unsafe.Add(mBase, uint32(v13)+76)) = int32(0)
	v537 = *(*int32)(unsafe.Add(mBase, uint32(v13)+56))
	v538 = *(*int32)(unsafe.Add(mBase, uint32(v13)+52))
	v539 = *(*int32)(unsafe.Add(mBase, uint32(v13)+48))
	v544 = base.B2i32(int32(2) < v538)
	if int32(2) < v538 {
		goto L100
	} else {
		goto L101
	}
L99:
	;
	v570 = *(*int32)(unsafe.Add(mBase, uint32(v17)+8))
	v571 = int32(0)
	v573 = v569 + v570
	if base.B2i32(v570 < v571)^base.B2i32(v573 < v569)|base.B2i32(v573 < v571) != 0 {
		goto L5
	} else {
		goto L106
	}
L100:
	;
	v545 = int32(_a_F_timestamp_pl_interval_7)
	goto L102
L101:
	;
	v545 = int32(_a_F_timestamp_pl_interval_9)
	goto L102
L102:
	;
	v546 = v545 + v537
	v551 = base.I32_div_s(v546, int32(4))
	v554 = base.I32_div_s(v546, int32(-100))
	v557 = base.I32_div_s(v546, int32(400))
	if int32(2) < v538 {
		goto L103
	} else {
		goto L104
	}
L103:
	;
	v561 = int32(1)
	goto L105
L104:
	;
	v561 = int32(13)
	goto L105
L105:
	;
	v566 = base.I32_div_s((v561+v538)*int32(_a_F_timestamp_pl_interval_8), int32(256))
	v569 = v539 + v546*int32(365) + v551 + v554 + v557 + v566 - int32(_a_F_timestamp_pl_interval_10)
	goto L99
L106:
	;
	v582 = v573 + int32(_a_F_timestamp_pl_interval_11)
	v583 = int32(_a_F_timestamp_pl_interval_4)
	v584 = base.I32_div_u_s(v582, v583)
	v585 = int32(3)
	v591 = int32(2)
	v596 = base.I32_div_u_s((v584*int32(1073595727)+v582)<<(uint(v591)%32)|v585, v583)
	v599 = v573 + v584*v585 + v596 + int32(_a_F_timestamp_pl_interval_6)
	v600 = int32(1461)
	v601 = base.I32_div_u_s(v599, v600)
	v604 = v601*int32(-1461) + v599
	v606 = v604 << (uint(v591) % 32)
	if base.Ui32(v600) <= base.Ui32(v606) {
		goto L109
	} else {
		goto L110
	}
L107:
	;
	v645 = *(*int32)(unsafe.Add(mBase, uint32(v13)+56))
	if v645 <= int32(-4713) {
		goto L113
	} else {
		goto L114
	}
L108:
	;
	v619 = base.I32_div_u_s(v606, int32(1461))
	*(*int32)(unsafe.Add(mBase, uint32(v440))) = v619 + v601<<(uint(int32(2))%32) - int32(_a_F_timestamp_pl_interval_7)
	v627 = v617 + int32(123)
	v631 = int32(base.Ui32(v627*int32(2141)) >> (uint(int32(16)) % 32))
	*(*int32)(unsafe.Add(mBase, uint32(v444))) = v627 - int32(base.Ui32(v631*int32(_a_F_timestamp_pl_interval_8))>>(uint(int32(8))%32))
	v641 = base.I32_rem_u_s(v631+int32(10), int32(12))
	*(*int32)(unsafe.Add(mBase, uint32(v442))) = v641 + int32(1)
	goto L107
L109:
	;
	v612 = base.I32_rem_u_s(v604+int32(305), int32(365))
	v617 = v612
	goto L108
L110:
	;
	goto L111
L111:
	;
	v616 = base.I32_rem_u_s(v604+int32(306), int32(366))
	v617 = v616
	goto L108
L112:
	;
	v662 = *(*int32)(unsafe.Add(mBase, uint32(v13)+48))
	v667 = base.B2i32(int32(2) < v661)
	if int32(2) < v661 {
		goto L124
	} else {
		goto L125
	}
L113:
	;
	if v645 != int32(-4713) {
		goto L2
	} else {
		goto L116
	}
L114:
	;
	goto L115
L115:
	;
	if v645 <= int32(_a_F_timestamp_pl_interval_12) {
		goto L118
	} else {
		goto L119
	}
L116:
	;
	v650 = *(*int32)(unsafe.Add(mBase, uint32(v13)+52))
	if int32(10) < v650 {
		v661 = v650
		goto L112
	} else {
		goto L117
	}
L117:
	;
	goto L2
L118:
	;
	v655 = *(*int32)(unsafe.Add(mBase, uint32(v13)+52))
	v661 = v655
	goto L112
L119:
	;
	goto L120
L120:
	;
	if v645 != int32(_a_F_timestamp_pl_interval_13) {
		goto L2
	} else {
		goto L121
	}
L121:
	;
	v658 = *(*int32)(unsafe.Add(mBase, uint32(v13)+52))
	if int32(5) < v658 {
		goto L2
	} else {
		goto L122
	}
L122:
	;
	v661 = v658
	goto L112
L123:
	;
	v695 = base.I64_extend_i32_s(v662 + v669*int32(365) + v674 + v677 + v680 + v689 - int32(_a_F_timestamp_pl_interval_10) - int32(_a_F_timestamp_pl_interval_5))
	v704 = int64(32)
	v705 = int64(20)
	v707 = int64(base.Ui64(v695) >> (uint(v704) % 64))
	v710 = int64(4294967295)
	v711 = int64(500654080)
	v713 = v695 & v710
	v714 = v711 * v713
	v718 = int64(base.Ui64(v714)>>(uint(v704)%64)) + v711*v707
	v725 = v713*v705 + v718&v710
	*(*int64)(unsafe.Add(mBase, uint32(v13)+8)) = v695*int64(0) + v695>>(uint(int64(63))%64)*int64(86400000000) + v705*v707 + int64(base.Ui64(v718)>>(uint(v704)%64)) + int64(base.Ui64(v725)>>(uint(v704)%64))
	*(*int64)(unsafe.Add(mBase, uint32(v13))) = v714&v710 | v725<<(uint(v704)%64)
	goto L130
L124:
	;
	v668 = int32(_a_F_timestamp_pl_interval_7)
	goto L126
L125:
	;
	v668 = int32(_a_F_timestamp_pl_interval_9)
	goto L126
L126:
	;
	v669 = v668 + v645
	v674 = base.I32_div_s(v669, int32(4))
	v677 = base.I32_div_s(v669, int32(-100))
	v680 = base.I32_div_s(v669, int32(400))
	if int32(2) < v661 {
		goto L127
	} else {
		goto L128
	}
L127:
	;
	v684 = int32(1)
	goto L129
L128:
	;
	v684 = int32(13)
	goto L129
L129:
	;
	v689 = base.I32_div_s((v684+v661)*int32(_a_F_timestamp_pl_interval_8), int32(256))
	goto L123
L130:
	;
	v736 = *(*int64)(unsafe.Add(mBase, uint32(v13)+8))
	v737 = *(*int64)(unsafe.Add(mBase, uint32(v13)))
	if v736 != v737>>(uint(int64(63))%64) {
		goto L2
	} else {
		goto L131
	}
L131:
	;
	v745 = *(*int32)(unsafe.Add(mBase, uint32(v13)+36))
	v746 = *(*int32)(unsafe.Add(mBase, uint32(v13)+40))
	v747 = *(*int32)(unsafe.Add(mBase, uint32(v13)+44))
	v748 = int32(60)
	v757 = base.I64_extend32_s(v533*int64(4293967296)+v531) + base.I64_extend_i32_s(v745+(v746+v747*v748)*v748)*int64(1000000)
	v760 = v737 + v757
	if base.B2i32(v757 < int64(0))^base.B2i32(v760 < v737)|base.B2i32(base.Ui64(v760-int64(9223371331200000000)) <= base.Ui64(int64(9011559254509551615))) != 0 {
		goto L2
	} else {
		goto L132
	}
L132:
	;
	v773 = v760
	goto L86
L133:
	;
	if base.Ui64(int64(-9011559254509551616)) <= base.Ui64(v779+int64(211813488000000000)) {
		goto L3
	} else {
		goto L134
	}
L134:
	;
	v792 = v779
	goto L9
L135:
	;
	m.G0 = v13 + int32(80)
	return v794
L136:
	;
	F_errcode(m, int32(134217858))
	mBase = m.M
	v806 = m.ExcPending
	if v806 != 0 {
		goto L18
	} else {
		goto L137
	}
L137:
	;
	F_errmsg(m, int32(_a_F_timestamp_pl_interval_0), int32(0))
	mBase = m.M
	v810 = m.ExcPending
	if v810 != 0 {
		goto L18
	} else {
		goto L138
	}
L138:
	;
	F_errfinish(m, int32(_a_F_timestamp_pl_interval_1), int32(3151), int32(_a_F_timestamp_pl_interval_2))
	mBase = m.M
	v815 = m.ExcPending
	if v815 != 0 {
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
	v822 = m.ExcPending
	if v822 != 0 {
		goto L18
	} else {
		goto L141
	}
L141:
	;
	F_errmsg(m, int32(_a_F_timestamp_pl_interval_0), int32(0))
	mBase = m.M
	v826 = m.ExcPending
	if v826 != 0 {
		goto L18
	} else {
		goto L142
	}
L142:
	;
	F_errfinish(m, int32(_a_F_timestamp_pl_interval_1), int32(3156), int32(_a_F_timestamp_pl_interval_2))
	mBase = m.M
	v831 = m.ExcPending
	if v831 != 0 {
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
	v838 = m.ExcPending
	if v838 != 0 {
		goto L18
	} else {
		goto L145
	}
L145:
	;
	F_errmsg(m, int32(_a_F_timestamp_pl_interval_0), int32(0))
	mBase = m.M
	v842 = m.ExcPending
	if v842 != 0 {
		goto L18
	} else {
		goto L146
	}
L146:
	;
	F_errfinish(m, int32(_a_F_timestamp_pl_interval_1), int32(3188), int32(_a_F_timestamp_pl_interval_2))
	mBase = m.M
	v847 = m.ExcPending
	if v847 != 0 {
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
	v854 = m.ExcPending
	if v854 != 0 {
		goto L18
	} else {
		goto L149
	}
L149:
	;
	F_errmsg(m, int32(_a_F_timestamp_pl_interval_0), int32(0))
	mBase = m.M
	v858 = m.ExcPending
	if v858 != 0 {
		goto L18
	} else {
		goto L150
	}
L150:
	;
	F_errfinish(m, int32(_a_F_timestamp_pl_interval_1), int32(3199), int32(_a_F_timestamp_pl_interval_2))
	mBase = m.M
	v863 = m.ExcPending
	if v863 != 0 {
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
	v870 = m.ExcPending
	if v870 != 0 {
		goto L18
	} else {
		goto L153
	}
L153:
	;
	F_errmsg(m, int32(_a_F_timestamp_pl_interval_0), int32(0))
	mBase = m.M
	v874 = m.ExcPending
	if v874 != 0 {
		goto L18
	} else {
		goto L154
	}
L154:
	;
	F_errfinish(m, int32(_a_F_timestamp_pl_interval_1), int32(3211), int32(_a_F_timestamp_pl_interval_2))
	mBase = m.M
	v879 = m.ExcPending
	if v879 != 0 {
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
	v886 = m.ExcPending
	if v886 != 0 {
		goto L18
	} else {
		goto L157
	}
L157:
	;
	F_errmsg(m, int32(_a_F_timestamp_pl_interval_0), int32(0))
	mBase = m.M
	v890 = m.ExcPending
	if v890 != 0 {
		goto L18
	} else {
		goto L158
	}
L158:
	;
	F_errfinish(m, int32(_a_F_timestamp_pl_interval_1), int32(3216), int32(_a_F_timestamp_pl_interval_2))
	mBase = m.M
	v895 = m.ExcPending
	if v895 != 0 {
		goto L18
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
	F_errcode(m, int32(134217858))
	mBase = m.M
	v905 = m.ExcPending
	if v905 != 0 {
		goto L18
	} else {
		goto L161
	}
L161:
	;
	F_errmsg(m, int32(_a_F_timestamp_pl_interval_0), int32(0))
	mBase = m.M
	v909 = m.ExcPending
	if v909 != 0 {
		goto L18
	} else {
		goto L162
	}
L162:
	;
	F_errfinish(m, int32(_a_F_timestamp_pl_interval_1), int32(3205), int32(_a_F_timestamp_pl_interval_2))
	mBase = m.M
	v914 = m.ExcPending
	if v914 != 0 {
		goto L18
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
	F_errcode(m, int32(134217858))
	mBase = m.M
	v923 = m.ExcPending
	if v923 != 0 {
		goto L18
	} else {
		goto L165
	}
L165:
	;
	F_errmsg(m, int32(_a_F_timestamp_pl_interval_0), int32(0))
	mBase = m.M
	v927 = m.ExcPending
	if v927 != 0 {
		goto L18
	} else {
		goto L166
	}
L166:
	;
	F_errfinish(m, int32(_a_F_timestamp_pl_interval_1), int32(3175), int32(_a_F_timestamp_pl_interval_2))
	mBase = m.M
	v932 = m.ExcPending
	if v932 != 0 {
		goto L18
	} else {
		goto L167
	}
L167:
	;
	base.Wasm_trap_unreachable()
	for {
	}
}
func F_timestamp_scale(m *base.Module, l0 int32) int32 {
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
	var v10 int64
	_ = v10
	var v18 int32
	_ = v18
	var v19 int64
	_ = v19
	var v20 int32
	_ = v20
	var v21 int32
	_ = v21
	v4 = m.G0
	v6 = v4 - int32(16)
	m.G0 = v6
	v8 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
	v9 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v10 = *(*int64)(unsafe.Add(mBase, uint32(v9)))
	*(*int64)(unsafe.Add(mBase, uint32(v6)+8)) = v10
	F_AdjustTimestampForTypmod(m, v6+int32(8), v8, int32(0))
	mBase = m.M
	v18 = m.ExcPending
	if v18 != 0 {
		return int32(0)
	} else {
		v19 = *(*int64)(unsafe.Add(mBase, uint32(v6)+8))
		v20 = F_Int64GetDatum(m, v19)
		mBase = m.M
		v21 = m.ExcPending
		if v21 != 0 {
			return int32(0)
		} else {
			m.G0 = v6 + int32(16)
			return v20
		}
	}
}
func F_timestamp_timestamptz(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v2 int32
	_ = v2
	var v3 int64
	_ = v3
	var v5 int64
	_ = v5
	var v8 int32
	_ = v8
	var v9 int32
	_ = v9
	var v10 int32
	_ = v10
	v2 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v3 = *(*int64)(unsafe.Add(mBase, uint32(v2)))
	v5 = F_timestamp2timestamptz_opt_overflow(m, v3, int32(0))
	mBase = m.M
	v8 = m.ExcPending
	if v8 != 0 {
		return int32(0)
	} else {
		v9 = F_Int64GetDatum(m, v5)
		mBase = m.M
		v10 = m.ExcPending
		if v10 != 0 {
			return int32(0)
		} else {
			return v9
		}
	}
}
func F_timestamp_to_char(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v7 int32
	_ = v7
	var v9 int32
	_ = v9
	var v11 int32
	_ = v11
	var v12 int64
	_ = v12
	var v13 int32
	_ = v13
	var v14 int32
	_ = v14
	var v17 int32
	_ = v17
	var v18 int32
	_ = v18
	var v24 int32
	_ = v24
	var v27 int32
	_ = v27
	var v34 int32
	_ = v34
	var v35 int32
	_ = v35
	var v41 int32
	_ = v41
	var v47 int32
	_ = v47
	var v53 int32
	_ = v53
	var v56 int64
	_ = v56
	var v68 int32
	_ = v68
	var v75 int32
	_ = v75
	var v76 int32
	_ = v76
	var v77 int32
	_ = v77
	var v78 int32
	_ = v78
	var v79 int32
	_ = v79
	var v84 int32
	_ = v84
	var v85 int32
	_ = v85
	var v86 int32
	_ = v86
	var v91 int32
	_ = v91
	var v94 int32
	_ = v94
	var v97 int32
	_ = v97
	var v101 int32
	_ = v101
	var v106 int32
	_ = v106
	var v109 int32
	_ = v109
	var v110 int32
	_ = v110
	var v113 int32
	_ = v113
	var v115 int32
	_ = v115
	var v124 int32
	_ = v124
	var v129 int32
	_ = v129
	var v132 int32
	_ = v132
	var v135 int32
	_ = v135
	var v144 int32
	_ = v144
	var v148 int64
	_ = v148
	var v150 int64
	_ = v150
	var v152 int64
	_ = v152
	var v154 int64
	_ = v154
	var v156 int32
	_ = v156
	var v160 int32
	_ = v160
	var v166 int32
	_ = v166
	var v167 int32
	_ = v167
	var v168 int32
	_ = v168
	var v172 int32
	_ = v172
	var v176 int32
	_ = v176
	var v185 int32
	_ = v185
	var v188 int32
	_ = v188
	var v192 int32
	_ = v192
	var v197 int32
	_ = v197
	v7 = m.G0
	v9 = v7 - int32(96)
	m.G0 = v9
	v11 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v12 = *(*int64)(unsafe.Add(mBase, uint32(v11)))
	v13 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
	v14 = F_pg_detoast_datum_packed(m, v13)
	mBase = m.M
	v17 = m.ExcPending
	if v17 != 0 {
		return int32(0)
	} else {
		v18 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v14))))
		if v18 == int32(1) {
			v24 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v14)+1)))
			if v24 == int32(18) {
				v27 = int32(16)
			} else {
				v27 = int32(0)
			}
			if base.Ui32((v24-int32(1))&int32(255)) < base.Ui32(int32(3)) {
				v34 = int32(4)
			} else {
				v34 = v27
			}
			v47 = v34
		} else {
			v35 = int32(1)
			if v18&v35 != 0 {
				v47 = int32(base.Ui32(v18)>>(uint(v35)%32)) - v35
			} else {
				v41 = *(*int32)(unsafe.Add(mBase, uint32(v14)))
				v47 = int32(base.Ui32(v41)>>(uint(int32(2))%32)) - int32(4)
			}
		}
		if base.Ui64(int64(1)) < base.Ui64(v12-int64(9223372036854775807)) {
			v53 = v47
		} else {
			v53 = int32(0)
		}
		if v53 == int32(0) {
			v172 = int32(1)
			*(*uint8)(unsafe.Add(mBase, uint32(l0)+16)) = uint8(v172)
			v176 = int32(0)
			m.G0 = v9 + int32(96)
			return v176
		} else {
			v56 = int64(0)
			*(*int64)(unsafe.Add(mBase, uint32(v9)+80)) = v56
			*(*int64)(unsafe.Add(mBase, uint32(v9)+72)) = v56
			*(*int64)(unsafe.Add(mBase, uint32(v9)+56)) = v56
			*(*int64)(unsafe.Add(mBase, uint32(v9)+48)) = v56
			*(*int64)(unsafe.Add(mBase, uint32(v9)+88)) = v56
			*(*int64)(unsafe.Add(mBase, uint32(v9)+64)) = int64(4294967297)
			v68 = int32(0)
			v75 = F_timestamp2tm(m, v12, v68, v9+int32(4), v9+int32(88), v68, v68)
			mBase = m.M
			v76 = m.ExcPending
			if v76 != 0 {
				return int32(0)
			} else {
				if v75 != 0 {
					F_errstart_cold(m, int32(21), int32(0))
					mBase = m.M
					v185 = m.ExcPending
					if v185 != 0 {
						return int32(0)
					} else {
						F_errcode(m, int32(134217858))
						mBase = m.M
						v188 = m.ExcPending
						if v188 != 0 {
							return int32(0)
						} else {
							F_errmsg(m, int32(_a_F_timestamp_to_char_0), int32(0))
							mBase = m.M
							v192 = m.ExcPending
							if v192 != 0 {
								return int32(0)
							} else {
								F_errfinish(m, int32(_a_F_timestamp_to_char_1), int32(4017), int32(_a_F_timestamp_to_char_2))
								mBase = m.M
								v197 = m.ExcPending
								if v197 != 0 {
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
					v77 = *(*int32)(unsafe.Add(mBase, uint32(v9)+24))
					v78 = *(*int32)(unsafe.Add(mBase, uint32(v9)+20))
					v79 = *(*int32)(unsafe.Add(mBase, uint32(v9)+16))
					v84 = base.B2i32(int32(2) < v78)
					if int32(2) < v78 {
						v85 = int32(_a_F_timestamp_to_char_3)
					} else {
						v85 = int32(_a_F_timestamp_to_char_4)
					}
					v86 = v85 + v77
					v91 = base.I32_div_s(v86, int32(4))
					v94 = base.I32_div_s(v86, int32(-100))
					v97 = base.I32_div_s(v86, int32(400))
					if int32(2) < v78 {
						v101 = int32(1)
					} else {
						v101 = int32(13)
					}
					v106 = base.I32_div_s((v101+v78)*int32(_a_F_timestamp_to_char_5), int32(256))
					v109 = v79 + v86*int32(365) + v91 + v94 + v97 + v106 - int32(_a_F_timestamp_to_char_6)
					v110 = int32(1)
					v113 = base.I32_rem_s(v109+v110, int32(7))
					*(*int32)(unsafe.Add(mBase, uint32(v9)+28)) = v113
					v115 = *(*int32)(unsafe.Add(mBase, uint32(v9)+24))
					v124 = int32(_a_F_timestamp_to_char_4) + v115
					v129 = base.I32_div_s(v124, int32(4))
					v132 = base.I32_div_s(v124, int32(-100))
					v135 = base.I32_div_s(v124, int32(400))
					v144 = base.I32_div_s(int32(_a_F_timestamp_to_char_7), int32(256))
					v148 = *(*int64)(unsafe.Add(mBase, uint32(v9)+4))
					*(*int64)(unsafe.Add(mBase, uint32(v9)+48)) = v148
					v150 = int64(*(*int32)(unsafe.Add(mBase, uint32(v9)+12)))
					*(*int64)(unsafe.Add(mBase, uint32(v9)+56)) = v150
					v152 = *(*int64)(unsafe.Add(mBase, uint32(v9)+16))
					*(*int64)(unsafe.Add(mBase, uint32(v9)+64)) = v152
					v154 = *(*int64)(unsafe.Add(mBase, uint32(v9)+24))
					*(*int64)(unsafe.Add(mBase, uint32(v9)+72)) = v154
					v156 = *(*int32)(unsafe.Add(mBase, uint32(v9)+40))
					*(*int32)(unsafe.Add(mBase, uint32(v9)+84)) = v156
					v160 = v109 - (v110 + v124*int32(365) + v129 + v132 + v135 + v144 - int32(_a_F_timestamp_to_char_6)) + int32(1)
					*(*int32)(unsafe.Add(mBase, uint32(v9)+32)) = v160
					*(*int32)(unsafe.Add(mBase, uint32(v9)+80)) = v160
					v166 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
					v167 = F_datetime_to_char_body(m, v9+int32(48), v14, int32(0), v166)
					mBase = m.M
					v168 = m.ExcPending
					if v168 != 0 {
						return int32(0)
					} else {
						if v167 != 0 {
							v176 = v167
						} else {
							v172 = int32(1)
							*(*uint8)(unsafe.Add(mBase, uint32(l0)+16)) = uint8(v172)
							v176 = int32(0)
						}
						m.G0 = v9 + int32(96)
						return v176
					}
				}
			}
		}
	}
}
func F_timestamp_zone(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v8 int32
	_ = v8
	var v10 int32
	_ = v10
	var v12 int32
	_ = v12
	var v13 int32
	_ = v13
	var v16 int32
	_ = v16
	var v17 int32
	_ = v17
	var v18 int64
	_ = v18
	var v23 int32
	_ = v23
	var v24 int32
	_ = v24
	var v26 int32
	_ = v26
	var v29 int32
	_ = v29
	var v34 int32
	_ = v34
	var v35 int32
	_ = v35
	var v37 int64
	_ = v37
	var v45 int64
	_ = v45
	var v46 int64
	_ = v46
	var v49 int64
	_ = v49
	var v52 int32
	_ = v52
	var v64 int32
	_ = v64
	var v65 int32
	_ = v65
	var v66 int32
	_ = v66
	var v67 int32
	_ = v67
	var v73 int32
	_ = v73
	var v78 int32
	_ = v78
	var v81 int32
	_ = v81
	var v82 int32
	_ = v82
	var v83 int32
	_ = v83
	var v86 int32
	_ = v86
	var v88 int32
	_ = v88
	var v94 int32
	_ = v94
	var v98 int32
	_ = v98
	var v99 int32
	_ = v99
	var v101 int32
	_ = v101
	var v109 int32
	_ = v109
	var v113 int32
	_ = v113
	var v123 int32
	_ = v123
	var v133 int64
	_ = v133
	var v135 int64
	_ = v135
	var v140 int64
	_ = v140
	var v142 int64
	_ = v142
	var v149 int64
	_ = v149
	var v151 int32
	_ = v151
	var v155 int32
	_ = v155
	var v158 int32
	_ = v158
	var v163 int32
	_ = v163
	var v165 int32
	_ = v165
	var v169 int32
	_ = v169
	var v171 int32
	_ = v171
	var v173 int32
	_ = v173
	var v174 int32
	_ = v174
	var v176 int32
	_ = v176
	var v180 int32
	_ = v180
	var v182 int32
	_ = v182
	var v184 int32
	_ = v184
	var v202 int32
	_ = v202
	var v203 int32
	_ = v203
	var v204 int32
	_ = v204
	var v209 int32
	_ = v209
	var v219 int64
	_ = v219
	var v227 int64
	_ = v227
	var v228 int64
	_ = v228
	var v231 int64
	_ = v231
	var v234 int32
	_ = v234
	var v246 int32
	_ = v246
	var v247 int32
	_ = v247
	var v248 int32
	_ = v248
	var v249 int32
	_ = v249
	var v255 int32
	_ = v255
	var v260 int32
	_ = v260
	var v263 int32
	_ = v263
	var v264 int32
	_ = v264
	var v265 int32
	_ = v265
	var v268 int32
	_ = v268
	var v270 int32
	_ = v270
	var v276 int32
	_ = v276
	var v280 int32
	_ = v280
	var v281 int32
	_ = v281
	var v283 int32
	_ = v283
	var v291 int32
	_ = v291
	var v295 int32
	_ = v295
	var v305 int32
	_ = v305
	var v317 int64
	_ = v317
	var v319 int64
	_ = v319
	var v324 int64
	_ = v324
	var v326 int64
	_ = v326
	var v331 int64
	_ = v331
	var v333 int64
	_ = v333
	var v337 int32
	_ = v337
	var v339 int32
	_ = v339
	var v340 int32
	_ = v340
	var v341 int32
	_ = v341
	var v345 int32
	_ = v345
	var v349 int32
	_ = v349
	var v354 int32
	_ = v354
	var v359 int32
	_ = v359
	var v362 int32
	_ = v362
	var v365 int32
	_ = v365
	var v366 int32
	_ = v366
	var v371 int32
	_ = v371
	var v372 int32
	_ = v372
	var v373 int32
	_ = v373
	var v378 int32
	_ = v378
	var v381 int32
	_ = v381
	var v384 int32
	_ = v384
	var v388 int32
	_ = v388
	var v393 int32
	_ = v393
	var v399 int64
	_ = v399
	var v408 int64
	_ = v408
	var v409 int64
	_ = v409
	var v411 int64
	_ = v411
	var v414 int64
	_ = v414
	var v415 int64
	_ = v415
	var v417 int64
	_ = v417
	var v418 int64
	_ = v418
	var v422 int64
	_ = v422
	var v429 int64
	_ = v429
	var v440 int64
	_ = v440
	var v441 int64
	_ = v441
	var v449 int32
	_ = v449
	var v450 int32
	_ = v450
	var v451 int32
	_ = v451
	var v452 int32
	_ = v452
	var v461 int64
	_ = v461
	var v464 int64
	_ = v464
	var v472 int64
	_ = v472
	var v483 int32
	_ = v483
	var v486 int32
	_ = v486
	var v490 int32
	_ = v490
	var v495 int32
	_ = v495
	var v496 int64
	_ = v496
	var v503 int64
	_ = v503
	var v510 int32
	_ = v510
	var v511 int32
	_ = v511
	var v518 int32
	_ = v518
	var v526 int32
	_ = v526
	var v529 int32
	_ = v529
	var v533 int32
	_ = v533
	var v538 int32
	_ = v538
	var v542 int32
	_ = v542
	var v545 int32
	_ = v545
	var v549 int32
	_ = v549
	var v554 int32
	_ = v554
	var v558 int32
	_ = v558
	var v561 int32
	_ = v561
	var v565 int32
	_ = v565
	var v570 int32
	_ = v570
	v8 = m.G0
	v10 = v8 - int32(336)
	m.G0 = v10
	v12 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v13 = F_pg_detoast_datum_packed(m, v12)
	mBase = m.M
	v16 = m.ExcPending
	if v16 != 0 {
		return int32(0)
	} else {
		v17 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
		v18 = *(*int64)(unsafe.Add(mBase, uint32(v17)))
		if base.Ui64(v18-int64(9223372036854775807)) <= base.Ui64(int64(1)) {
			v23 = F_Int64GetDatum(m, v18)
			mBase = m.M
			v24 = m.ExcPending
			if v24 != 0 {
				return int32(0)
			} else {
				v518 = v23
				m.G0 = v10 + int32(336)
				return v518
			}
		} else {
			v26 = v10 + int32(80)
			F_text_to_cstring_buffer(m, v13, v26, int32(256))
			mBase = m.M
			v29 = m.ExcPending
			if v29 != 0 {
				return int32(0)
			} else {
				v34 = F_DecodeTimezoneName(m, v26, v10+int32(76), v10+int32(72))
				mBase = m.M
				v35 = m.ExcPending
				if v35 != 0 {
					return int32(0)
				} else {
					switch v34 {
					case 0:
						v496 = int64(*(*int32)(unsafe.Add(mBase, uint32(v10)+76)))
						v503 = v496*int64(-1000000) + v18
						if base.Ui64(int64(-9011559254509551616)) <= base.Ui64(v503+int64(211813488000000000)) {
							F_errstart_cold(m, int32(21), int32(0))
							mBase = m.M
							v558 = m.ExcPending
							if v558 != 0 {
								return int32(0)
							} else {
								F_errcode(m, int32(134217858))
								mBase = m.M
								v561 = m.ExcPending
								if v561 != 0 {
									return int32(0)
								} else {
									F_errmsg(m, int32(_a_F_timestamp_zone_0), int32(0))
									mBase = m.M
									v565 = m.ExcPending
									if v565 != 0 {
										return int32(0)
									} else {
										F_errfinish(m, int32(_a_F_timestamp_zone_1), int32(_a_F_timestamp_zone_2), int32(_a_F_timestamp_zone_3))
										mBase = m.M
										v570 = m.ExcPending
										if v570 != 0 {
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
							v510 = F_Int64GetDatum(m, v503)
							mBase = m.M
							v511 = m.ExcPending
							if v511 != 0 {
								return int32(0)
							} else {
								v518 = v510
								m.G0 = v10 + int32(336)
								return v518
							}
						}
					case 1:
						v37 = base.I64_div_s(v18, int64(86400000000))
						if base.Ui64(int64(172799999999)) <= base.Ui64(v18+int64(86399999999)) {
							v45 = v37 * int64(-86400000000)
						} else {
							v45 = int64(0)
						}
						v46 = v45 + v18
						v49 = v46>>(uint(int64(63))%64) + v37
						if v49 <= int64(-2451546) {
							F_errstart_cold(m, int32(21), int32(0))
							mBase = m.M
							v526 = m.ExcPending
							if v526 != 0 {
								return int32(0)
							} else {
								F_errcode(m, int32(134217858))
								mBase = m.M
								v529 = m.ExcPending
								if v529 != 0 {
									return int32(0)
								} else {
									F_errmsg(m, int32(_a_F_timestamp_zone_0), int32(0))
									mBase = m.M
									v533 = m.ExcPending
									if v533 != 0 {
										return int32(0)
									} else {
										F_errfinish(m, int32(_a_F_timestamp_zone_1), int32(_a_F_timestamp_zone_4), int32(_a_F_timestamp_zone_3))
										mBase = m.M
										v538 = m.ExcPending
										if v538 != 0 {
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
							v52 = base.I32_wrap_i64(v49)
							v64 = v52 + int32(_a_F_timestamp_zone_5)
							v65 = int32(_a_F_timestamp_zone_6)
							v66 = base.I32_div_u_s(v64, v65)
							v67 = int32(3)
							v73 = int32(2)
							v78 = base.I32_div_u_s((v66*int32(1073595727)+v64)<<(uint(v73)%32)|v67, v65)
							v81 = v52 + int32(_a_F_timestamp_zone_7) + v66*v67 + v78 + int32(_a_F_timestamp_zone_8)
							v82 = int32(1461)
							v83 = base.I32_div_u_s(v81, v82)
							v86 = v83*int32(-1461) + v81
							v88 = v86 << (uint(v73) % 32)
							if base.Ui32(v82) <= base.Ui32(v88) {
								v94 = base.I32_rem_u_s(v86+int32(305), int32(365))
								v99 = v94
							} else {
								v98 = base.I32_rem_u_s(v86+int32(306), int32(366))
								v99 = v98
							}
							v101 = base.I32_div_u_s(v88, int32(1461))
							*(*int32)(unsafe.Add(mBase, uint32(v10+int32(48)))) = v101 + v83<<(uint(int32(2))%32) - int32(_a_F_timestamp_zone_9)
							v109 = v99 + int32(123)
							v113 = int32(base.Ui32(v109*int32(2141)) >> (uint(int32(16)) % 32))
							*(*int32)(unsafe.Add(mBase, uint32(v10+int32(40)))) = v109 - int32(base.Ui32(v113*int32(_a_F_timestamp_zone_10))>>(uint(int32(8))%32))
							v123 = base.I32_rem_u_s(v113+int32(10), int32(12))
							*(*int32)(unsafe.Add(mBase, uint32(v10+int32(44)))) = v123 + int32(1)
							*(*int64)(unsafe.Add(mBase, uint32(v10)+60)) = int64(4294967295)
							if v46 < int64(0) {
								v133 = v46 + int64(86400000000)
							} else {
								v133 = v46
							}
							v135 = base.I64_div_s(v133, int64(3600000000))
							*(*uint32)(unsafe.Add(mBase, uint32(v10)+36)) = uint32(v135)
							v140 = base.I64_extend32_s(v135)*int64(-3600000000) + v133
							v142 = base.I64_div_s(v140, int64(60000000))
							*(*uint32)(unsafe.Add(mBase, uint32(v10)+32)) = uint32(v142)
							v149 = base.I64_div_s(base.I64_extend32_s(v142)*int64(-60000000)+v140, int64(1000000))
							*(*uint32)(unsafe.Add(mBase, uint32(v10)+28)) = uint32(v149)
							v151 = int32(0)
							*(*int32)(unsafe.Add(mBase, uint32(v10)+68)) = v151
							v155 = v10 + int32(28)
							v158 = *(*int32)(unsafe.Add(mBase, uint32(v10)+72))
							v163 = m.G0
							v165 = v163 - int32(288)
							m.G0 = v165
							v169 = F_DetermineTimeZoneOffsetInternal(m, v155, v158, v165+int32(280))
							mBase = m.M
							v171 = v165 + int32(16)
							v173 = F_strlcpy(m, v171, v10+int32(80), int32(256))
							mBase = m.M
							v174 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v165)+16)))
							if v174 != 0 {
								v176 = v171
								v180 = v174
								for {
									v182 = F_pg_toupper(m, v180)
									mBase = m.M
									*(*uint8)(unsafe.Add(mBase, uint32(v176))) = uint8(v182)
									v184 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v176)+1)))
									if v184 != 0 {
										v176 = v176 + int32(1)
										v180 = v184
										continue
									} else {
										break
									}
									break
								}
							} else {
							}
							v202 = F_pg_interpret_timezone_abbrev(m, v165+int32(16), v165+int32(280), v165+int32(12), v165+int32(8), v158)
							mBase = m.M
							if v202 != 0 {
								v203 = *(*int32)(unsafe.Add(mBase, uint32(v165)+12))
								v204 = *(*int32)(unsafe.Add(mBase, uint32(v165)+8))
								*(*int32)(unsafe.Add(mBase, uint32(v155)+32)) = v204
								v209 = int32(0) - v203
							} else {
								v209 = v169
							}
							m.G0 = v165 + int32(288)
							v503 = base.I64_extend_i32_s(v151-v209)*int64(-1000000) + v18
							if base.Ui64(int64(-9011559254509551616)) <= base.Ui64(v503+int64(211813488000000000)) {
								F_errstart_cold(m, int32(21), int32(0))
								mBase = m.M
								v558 = m.ExcPending
								if v558 != 0 {
									return int32(0)
								} else {
									F_errcode(m, int32(134217858))
									mBase = m.M
									v561 = m.ExcPending
									if v561 != 0 {
										return int32(0)
									} else {
										F_errmsg(m, int32(_a_F_timestamp_zone_0), int32(0))
										mBase = m.M
										v565 = m.ExcPending
										if v565 != 0 {
											return int32(0)
										} else {
											F_errfinish(m, int32(_a_F_timestamp_zone_1), int32(_a_F_timestamp_zone_2), int32(_a_F_timestamp_zone_3))
											mBase = m.M
											v570 = m.ExcPending
											if v570 != 0 {
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
								v510 = F_Int64GetDatum(m, v503)
								mBase = m.M
								v511 = m.ExcPending
								if v511 != 0 {
									return int32(0)
								} else {
									v518 = v510
									m.G0 = v10 + int32(336)
									return v518
								}
							}
						}
					default:
						v219 = base.I64_div_s(v18, int64(86400000000))
						if base.Ui64(int64(172799999999)) <= base.Ui64(v18+int64(86399999999)) {
							v227 = v219 * int64(-86400000000)
						} else {
							v227 = int64(0)
						}
						v228 = v227 + v18
						v231 = v228>>(uint(int64(63))%64) + v219
						if v231 <= int64(-2451546) {
							F_errstart_cold(m, int32(21), int32(0))
							mBase = m.M
							v542 = m.ExcPending
							if v542 != 0 {
								return int32(0)
							} else {
								F_errcode(m, int32(134217858))
								mBase = m.M
								v545 = m.ExcPending
								if v545 != 0 {
									return int32(0)
								} else {
									F_errmsg(m, int32(_a_F_timestamp_zone_0), int32(0))
									mBase = m.M
									v549 = m.ExcPending
									if v549 != 0 {
										return int32(0)
									} else {
										F_errfinish(m, int32(_a_F_timestamp_zone_1), int32(_a_F_timestamp_zone_11), int32(_a_F_timestamp_zone_3))
										mBase = m.M
										v554 = m.ExcPending
										if v554 != 0 {
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
							v234 = base.I32_wrap_i64(v231)
							v246 = v234 + int32(_a_F_timestamp_zone_5)
							v247 = int32(_a_F_timestamp_zone_6)
							v248 = base.I32_div_u_s(v246, v247)
							v249 = int32(3)
							v255 = int32(2)
							v260 = base.I32_div_u_s((v248*int32(1073595727)+v246)<<(uint(v255)%32)|v249, v247)
							v263 = v234 + int32(_a_F_timestamp_zone_7) + v248*v249 + v260 + int32(_a_F_timestamp_zone_8)
							v264 = int32(1461)
							v265 = base.I32_div_u_s(v263, v264)
							v268 = v265*int32(-1461) + v263
							v270 = v268 << (uint(v255) % 32)
							if base.Ui32(v264) <= base.Ui32(v270) {
								v276 = base.I32_rem_u_s(v268+int32(305), int32(365))
								v281 = v276
							} else {
								v280 = base.I32_rem_u_s(v268+int32(306), int32(366))
								v281 = v280
							}
							v283 = base.I32_div_u_s(v270, int32(1461))
							*(*int32)(unsafe.Add(mBase, uint32(v10+int32(48)))) = v283 + v265<<(uint(int32(2))%32) - int32(_a_F_timestamp_zone_9)
							v291 = v281 + int32(123)
							v295 = int32(base.Ui32(v291*int32(2141)) >> (uint(int32(16)) % 32))
							*(*int32)(unsafe.Add(mBase, uint32(v10+int32(40)))) = v291 - int32(base.Ui32(v295*int32(_a_F_timestamp_zone_10))>>(uint(int32(8))%32))
							v305 = base.I32_rem_u_s(v295+int32(10), int32(12))
							*(*int32)(unsafe.Add(mBase, uint32(v10+int32(44)))) = v305 + int32(1)
							*(*int32)(unsafe.Add(mBase, uint32(v10)+68)) = int32(0)
							*(*int64)(unsafe.Add(mBase, uint32(v10)+60)) = int64(4294967295)
							if v228 < int64(0) {
								v317 = v228 + int64(86400000000)
							} else {
								v317 = v228
							}
							v319 = base.I64_div_s(v317, int64(3600000000))
							*(*uint32)(unsafe.Add(mBase, uint32(v10)+36)) = uint32(v319)
							v324 = base.I64_extend32_s(v319)*int64(-3600000000) + v317
							v326 = base.I64_div_s(v324, int64(60000000))
							*(*uint32)(unsafe.Add(mBase, uint32(v10)+32)) = uint32(v326)
							v331 = base.I64_extend32_s(v326)*int64(-60000000) + v324
							v333 = base.I64_div_s(v331, int64(1000000))
							*(*uint32)(unsafe.Add(mBase, uint32(v10)+28)) = uint32(v333)
							v337 = *(*int32)(unsafe.Add(mBase, uint32(v10)+72))
							v339 = m.G0
							v340 = int32(16)
							v341 = v339 - v340
							m.G0 = v341
							v345 = F_DetermineTimeZoneOffsetInternal(m, v10+int32(28), v337, v341+int32(8))
							mBase = m.M
							m.G0 = v341 + v340
							v349 = *(*int32)(unsafe.Add(mBase, uint32(v10)+48))
							if v349 <= int32(-4713) {
								if v349 != int32(-4713) {
									F_errstart_cold(m, int32(21), int32(0))
									mBase = m.M
									v483 = m.ExcPending
									if v483 != 0 {
										return int32(0)
									} else {
										F_errcode(m, int32(134217858))
										mBase = m.M
										v486 = m.ExcPending
										if v486 != 0 {
											return int32(0)
										} else {
											F_errmsg(m, int32(_a_F_timestamp_zone_0), int32(0))
											mBase = m.M
											v490 = m.ExcPending
											if v490 != 0 {
												return int32(0)
											} else {
												F_errfinish(m, int32(_a_F_timestamp_zone_1), int32(_a_F_timestamp_zone_12), int32(_a_F_timestamp_zone_3))
												mBase = m.M
												v495 = m.ExcPending
												if v495 != 0 {
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
									v354 = *(*int32)(unsafe.Add(mBase, uint32(v10)+44))
									if int32(10) < v354 {
										v365 = v354
										v366 = *(*int32)(unsafe.Add(mBase, uint32(v10)+40))
										v371 = base.B2i32(int32(2) < v365)
										if int32(2) < v365 {
											v372 = int32(_a_F_timestamp_zone_9)
										} else {
											v372 = int32(_a_F_timestamp_zone_13)
										}
										v373 = v372 + v349
										v378 = base.I32_div_s(v373, int32(4))
										v381 = base.I32_div_s(v373, int32(-100))
										v384 = base.I32_div_s(v373, int32(400))
										if int32(2) < v365 {
											v388 = int32(1)
										} else {
											v388 = int32(13)
										}
										v393 = base.I32_div_s((v388+v365)*int32(_a_F_timestamp_zone_10), int32(256))
										v399 = base.I64_extend_i32_s(v366 + v373*int32(365) + v378 + v381 + v384 + v393 - int32(_a_F_timestamp_zone_14) - int32(_a_F_timestamp_zone_7))
										v408 = int64(32)
										v409 = int64(20)
										v411 = int64(base.Ui64(v399) >> (uint(v408) % 64))
										v414 = int64(4294967295)
										v415 = int64(500654080)
										v417 = v399 & v414
										v418 = v415 * v417
										v422 = int64(base.Ui64(v418)>>(uint(v408)%64)) + v415*v411
										v429 = v417*v409 + v422&v414
										*(*int64)(unsafe.Add(mBase, uint32(v10)+8)) = v399*int64(0) + v399>>(uint(int64(63))%64)*int64(86400000000) + v409*v411 + int64(base.Ui64(v422)>>(uint(v408)%64)) + int64(base.Ui64(v429)>>(uint(v408)%64))
										*(*int64)(unsafe.Add(mBase, uint32(v10))) = v418&v414 | v429<<(uint(v408)%64)
										v440 = *(*int64)(unsafe.Add(mBase, uint32(v10)+8))
										v441 = *(*int64)(unsafe.Add(mBase, uint32(v10)))
										if v440 != v441>>(uint(int64(63))%64) {
											F_errstart_cold(m, int32(21), int32(0))
											mBase = m.M
											v483 = m.ExcPending
											if v483 != 0 {
												return int32(0)
											} else {
												F_errcode(m, int32(134217858))
												mBase = m.M
												v486 = m.ExcPending
												if v486 != 0 {
													return int32(0)
												} else {
													F_errmsg(m, int32(_a_F_timestamp_zone_0), int32(0))
													mBase = m.M
													v490 = m.ExcPending
													if v490 != 0 {
														return int32(0)
													} else {
														F_errfinish(m, int32(_a_F_timestamp_zone_1), int32(_a_F_timestamp_zone_12), int32(_a_F_timestamp_zone_3))
														mBase = m.M
														v495 = m.ExcPending
														if v495 != 0 {
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
											v449 = *(*int32)(unsafe.Add(mBase, uint32(v10)+28))
											v450 = *(*int32)(unsafe.Add(mBase, uint32(v10)+32))
											v451 = *(*int32)(unsafe.Add(mBase, uint32(v10)+36))
											v452 = int32(60)
											v461 = base.I64_extend32_s(v333*int64(4293967296)+v331) + base.I64_extend_i32_s(v449+(v450+v451*v452)*v452)*int64(1000000)
											v464 = v441 + v461
											if base.B2i32(v461 < int64(0))^base.B2i32(v464 < v441) != 0 {
												F_errstart_cold(m, int32(21), int32(0))
												mBase = m.M
												v483 = m.ExcPending
												if v483 != 0 {
													return int32(0)
												} else {
													F_errcode(m, int32(134217858))
													mBase = m.M
													v486 = m.ExcPending
													if v486 != 0 {
														return int32(0)
													} else {
														F_errmsg(m, int32(_a_F_timestamp_zone_0), int32(0))
														mBase = m.M
														v490 = m.ExcPending
														if v490 != 0 {
															return int32(0)
														} else {
															F_errfinish(m, int32(_a_F_timestamp_zone_1), int32(_a_F_timestamp_zone_12), int32(_a_F_timestamp_zone_3))
															mBase = m.M
															v495 = m.ExcPending
															if v495 != 0 {
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
												v472 = v464 + base.I64_extend_i32_s(int32(0)-v345)*int64(-1000000)
												if base.Ui64(v472+int64(211813488000000000)) < base.Ui64(int64(-9011559254509551616)) {
													v503 = v472
													if base.Ui64(int64(-9011559254509551616)) <= base.Ui64(v503+int64(211813488000000000)) {
														F_errstart_cold(m, int32(21), int32(0))
														mBase = m.M
														v558 = m.ExcPending
														if v558 != 0 {
															return int32(0)
														} else {
															F_errcode(m, int32(134217858))
															mBase = m.M
															v561 = m.ExcPending
															if v561 != 0 {
																return int32(0)
															} else {
																F_errmsg(m, int32(_a_F_timestamp_zone_0), int32(0))
																mBase = m.M
																v565 = m.ExcPending
																if v565 != 0 {
																	return int32(0)
																} else {
																	F_errfinish(m, int32(_a_F_timestamp_zone_1), int32(_a_F_timestamp_zone_2), int32(_a_F_timestamp_zone_3))
																	mBase = m.M
																	v570 = m.ExcPending
																	if v570 != 0 {
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
														v510 = F_Int64GetDatum(m, v503)
														mBase = m.M
														v511 = m.ExcPending
														if v511 != 0 {
															return int32(0)
														} else {
															v518 = v510
															m.G0 = v10 + int32(336)
															return v518
														}
													}
												} else {
													F_errstart_cold(m, int32(21), int32(0))
													mBase = m.M
													v483 = m.ExcPending
													if v483 != 0 {
														return int32(0)
													} else {
														F_errcode(m, int32(134217858))
														mBase = m.M
														v486 = m.ExcPending
														if v486 != 0 {
															return int32(0)
														} else {
															F_errmsg(m, int32(_a_F_timestamp_zone_0), int32(0))
															mBase = m.M
															v490 = m.ExcPending
															if v490 != 0 {
																return int32(0)
															} else {
																F_errfinish(m, int32(_a_F_timestamp_zone_1), int32(_a_F_timestamp_zone_12), int32(_a_F_timestamp_zone_3))
																mBase = m.M
																v495 = m.ExcPending
																if v495 != 0 {
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
										F_errstart_cold(m, int32(21), int32(0))
										mBase = m.M
										v483 = m.ExcPending
										if v483 != 0 {
											return int32(0)
										} else {
											F_errcode(m, int32(134217858))
											mBase = m.M
											v486 = m.ExcPending
											if v486 != 0 {
												return int32(0)
											} else {
												F_errmsg(m, int32(_a_F_timestamp_zone_0), int32(0))
												mBase = m.M
												v490 = m.ExcPending
												if v490 != 0 {
													return int32(0)
												} else {
													F_errfinish(m, int32(_a_F_timestamp_zone_1), int32(_a_F_timestamp_zone_12), int32(_a_F_timestamp_zone_3))
													mBase = m.M
													v495 = m.ExcPending
													if v495 != 0 {
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
								if v349 <= int32(_a_F_timestamp_zone_15) {
									v359 = *(*int32)(unsafe.Add(mBase, uint32(v10)+44))
									v365 = v359
									v366 = *(*int32)(unsafe.Add(mBase, uint32(v10)+40))
									v371 = base.B2i32(int32(2) < v365)
									if int32(2) < v365 {
										v372 = int32(_a_F_timestamp_zone_9)
									} else {
										v372 = int32(_a_F_timestamp_zone_13)
									}
									v373 = v372 + v349
									v378 = base.I32_div_s(v373, int32(4))
									v381 = base.I32_div_s(v373, int32(-100))
									v384 = base.I32_div_s(v373, int32(400))
									if int32(2) < v365 {
										v388 = int32(1)
									} else {
										v388 = int32(13)
									}
									v393 = base.I32_div_s((v388+v365)*int32(_a_F_timestamp_zone_10), int32(256))
									v399 = base.I64_extend_i32_s(v366 + v373*int32(365) + v378 + v381 + v384 + v393 - int32(_a_F_timestamp_zone_14) - int32(_a_F_timestamp_zone_7))
									v408 = int64(32)
									v409 = int64(20)
									v411 = int64(base.Ui64(v399) >> (uint(v408) % 64))
									v414 = int64(4294967295)
									v415 = int64(500654080)
									v417 = v399 & v414
									v418 = v415 * v417
									v422 = int64(base.Ui64(v418)>>(uint(v408)%64)) + v415*v411
									v429 = v417*v409 + v422&v414
									*(*int64)(unsafe.Add(mBase, uint32(v10)+8)) = v399*int64(0) + v399>>(uint(int64(63))%64)*int64(86400000000) + v409*v411 + int64(base.Ui64(v422)>>(uint(v408)%64)) + int64(base.Ui64(v429)>>(uint(v408)%64))
									*(*int64)(unsafe.Add(mBase, uint32(v10))) = v418&v414 | v429<<(uint(v408)%64)
									v440 = *(*int64)(unsafe.Add(mBase, uint32(v10)+8))
									v441 = *(*int64)(unsafe.Add(mBase, uint32(v10)))
									if v440 != v441>>(uint(int64(63))%64) {
										F_errstart_cold(m, int32(21), int32(0))
										mBase = m.M
										v483 = m.ExcPending
										if v483 != 0 {
											return int32(0)
										} else {
											F_errcode(m, int32(134217858))
											mBase = m.M
											v486 = m.ExcPending
											if v486 != 0 {
												return int32(0)
											} else {
												F_errmsg(m, int32(_a_F_timestamp_zone_0), int32(0))
												mBase = m.M
												v490 = m.ExcPending
												if v490 != 0 {
													return int32(0)
												} else {
													F_errfinish(m, int32(_a_F_timestamp_zone_1), int32(_a_F_timestamp_zone_12), int32(_a_F_timestamp_zone_3))
													mBase = m.M
													v495 = m.ExcPending
													if v495 != 0 {
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
										v449 = *(*int32)(unsafe.Add(mBase, uint32(v10)+28))
										v450 = *(*int32)(unsafe.Add(mBase, uint32(v10)+32))
										v451 = *(*int32)(unsafe.Add(mBase, uint32(v10)+36))
										v452 = int32(60)
										v461 = base.I64_extend32_s(v333*int64(4293967296)+v331) + base.I64_extend_i32_s(v449+(v450+v451*v452)*v452)*int64(1000000)
										v464 = v441 + v461
										if base.B2i32(v461 < int64(0))^base.B2i32(v464 < v441) != 0 {
											F_errstart_cold(m, int32(21), int32(0))
											mBase = m.M
											v483 = m.ExcPending
											if v483 != 0 {
												return int32(0)
											} else {
												F_errcode(m, int32(134217858))
												mBase = m.M
												v486 = m.ExcPending
												if v486 != 0 {
													return int32(0)
												} else {
													F_errmsg(m, int32(_a_F_timestamp_zone_0), int32(0))
													mBase = m.M
													v490 = m.ExcPending
													if v490 != 0 {
														return int32(0)
													} else {
														F_errfinish(m, int32(_a_F_timestamp_zone_1), int32(_a_F_timestamp_zone_12), int32(_a_F_timestamp_zone_3))
														mBase = m.M
														v495 = m.ExcPending
														if v495 != 0 {
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
											v472 = v464 + base.I64_extend_i32_s(int32(0)-v345)*int64(-1000000)
											if base.Ui64(v472+int64(211813488000000000)) < base.Ui64(int64(-9011559254509551616)) {
												v503 = v472
												if base.Ui64(int64(-9011559254509551616)) <= base.Ui64(v503+int64(211813488000000000)) {
													F_errstart_cold(m, int32(21), int32(0))
													mBase = m.M
													v558 = m.ExcPending
													if v558 != 0 {
														return int32(0)
													} else {
														F_errcode(m, int32(134217858))
														mBase = m.M
														v561 = m.ExcPending
														if v561 != 0 {
															return int32(0)
														} else {
															F_errmsg(m, int32(_a_F_timestamp_zone_0), int32(0))
															mBase = m.M
															v565 = m.ExcPending
															if v565 != 0 {
																return int32(0)
															} else {
																F_errfinish(m, int32(_a_F_timestamp_zone_1), int32(_a_F_timestamp_zone_2), int32(_a_F_timestamp_zone_3))
																mBase = m.M
																v570 = m.ExcPending
																if v570 != 0 {
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
													v510 = F_Int64GetDatum(m, v503)
													mBase = m.M
													v511 = m.ExcPending
													if v511 != 0 {
														return int32(0)
													} else {
														v518 = v510
														m.G0 = v10 + int32(336)
														return v518
													}
												}
											} else {
												F_errstart_cold(m, int32(21), int32(0))
												mBase = m.M
												v483 = m.ExcPending
												if v483 != 0 {
													return int32(0)
												} else {
													F_errcode(m, int32(134217858))
													mBase = m.M
													v486 = m.ExcPending
													if v486 != 0 {
														return int32(0)
													} else {
														F_errmsg(m, int32(_a_F_timestamp_zone_0), int32(0))
														mBase = m.M
														v490 = m.ExcPending
														if v490 != 0 {
															return int32(0)
														} else {
															F_errfinish(m, int32(_a_F_timestamp_zone_1), int32(_a_F_timestamp_zone_12), int32(_a_F_timestamp_zone_3))
															mBase = m.M
															v495 = m.ExcPending
															if v495 != 0 {
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
									if v349 != int32(_a_F_timestamp_zone_16) {
										F_errstart_cold(m, int32(21), int32(0))
										mBase = m.M
										v483 = m.ExcPending
										if v483 != 0 {
											return int32(0)
										} else {
											F_errcode(m, int32(134217858))
											mBase = m.M
											v486 = m.ExcPending
											if v486 != 0 {
												return int32(0)
											} else {
												F_errmsg(m, int32(_a_F_timestamp_zone_0), int32(0))
												mBase = m.M
												v490 = m.ExcPending
												if v490 != 0 {
													return int32(0)
												} else {
													F_errfinish(m, int32(_a_F_timestamp_zone_1), int32(_a_F_timestamp_zone_12), int32(_a_F_timestamp_zone_3))
													mBase = m.M
													v495 = m.ExcPending
													if v495 != 0 {
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
										v362 = *(*int32)(unsafe.Add(mBase, uint32(v10)+44))
										if int32(5) < v362 {
											F_errstart_cold(m, int32(21), int32(0))
											mBase = m.M
											v483 = m.ExcPending
											if v483 != 0 {
												return int32(0)
											} else {
												F_errcode(m, int32(134217858))
												mBase = m.M
												v486 = m.ExcPending
												if v486 != 0 {
													return int32(0)
												} else {
													F_errmsg(m, int32(_a_F_timestamp_zone_0), int32(0))
													mBase = m.M
													v490 = m.ExcPending
													if v490 != 0 {
														return int32(0)
													} else {
														F_errfinish(m, int32(_a_F_timestamp_zone_1), int32(_a_F_timestamp_zone_12), int32(_a_F_timestamp_zone_3))
														mBase = m.M
														v495 = m.ExcPending
														if v495 != 0 {
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
											v365 = v362
											v366 = *(*int32)(unsafe.Add(mBase, uint32(v10)+40))
											v371 = base.B2i32(int32(2) < v365)
											if int32(2) < v365 {
												v372 = int32(_a_F_timestamp_zone_9)
											} else {
												v372 = int32(_a_F_timestamp_zone_13)
											}
											v373 = v372 + v349
											v378 = base.I32_div_s(v373, int32(4))
											v381 = base.I32_div_s(v373, int32(-100))
											v384 = base.I32_div_s(v373, int32(400))
											if int32(2) < v365 {
												v388 = int32(1)
											} else {
												v388 = int32(13)
											}
											v393 = base.I32_div_s((v388+v365)*int32(_a_F_timestamp_zone_10), int32(256))
											v399 = base.I64_extend_i32_s(v366 + v373*int32(365) + v378 + v381 + v384 + v393 - int32(_a_F_timestamp_zone_14) - int32(_a_F_timestamp_zone_7))
											v408 = int64(32)
											v409 = int64(20)
											v411 = int64(base.Ui64(v399) >> (uint(v408) % 64))
											v414 = int64(4294967295)
											v415 = int64(500654080)
											v417 = v399 & v414
											v418 = v415 * v417
											v422 = int64(base.Ui64(v418)>>(uint(v408)%64)) + v415*v411
											v429 = v417*v409 + v422&v414
											*(*int64)(unsafe.Add(mBase, uint32(v10)+8)) = v399*int64(0) + v399>>(uint(int64(63))%64)*int64(86400000000) + v409*v411 + int64(base.Ui64(v422)>>(uint(v408)%64)) + int64(base.Ui64(v429)>>(uint(v408)%64))
											*(*int64)(unsafe.Add(mBase, uint32(v10))) = v418&v414 | v429<<(uint(v408)%64)
											v440 = *(*int64)(unsafe.Add(mBase, uint32(v10)+8))
											v441 = *(*int64)(unsafe.Add(mBase, uint32(v10)))
											if v440 != v441>>(uint(int64(63))%64) {
												F_errstart_cold(m, int32(21), int32(0))
												mBase = m.M
												v483 = m.ExcPending
												if v483 != 0 {
													return int32(0)
												} else {
													F_errcode(m, int32(134217858))
													mBase = m.M
													v486 = m.ExcPending
													if v486 != 0 {
														return int32(0)
													} else {
														F_errmsg(m, int32(_a_F_timestamp_zone_0), int32(0))
														mBase = m.M
														v490 = m.ExcPending
														if v490 != 0 {
															return int32(0)
														} else {
															F_errfinish(m, int32(_a_F_timestamp_zone_1), int32(_a_F_timestamp_zone_12), int32(_a_F_timestamp_zone_3))
															mBase = m.M
															v495 = m.ExcPending
															if v495 != 0 {
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
												v449 = *(*int32)(unsafe.Add(mBase, uint32(v10)+28))
												v450 = *(*int32)(unsafe.Add(mBase, uint32(v10)+32))
												v451 = *(*int32)(unsafe.Add(mBase, uint32(v10)+36))
												v452 = int32(60)
												v461 = base.I64_extend32_s(v333*int64(4293967296)+v331) + base.I64_extend_i32_s(v449+(v450+v451*v452)*v452)*int64(1000000)
												v464 = v441 + v461
												if base.B2i32(v461 < int64(0))^base.B2i32(v464 < v441) != 0 {
													F_errstart_cold(m, int32(21), int32(0))
													mBase = m.M
													v483 = m.ExcPending
													if v483 != 0 {
														return int32(0)
													} else {
														F_errcode(m, int32(134217858))
														mBase = m.M
														v486 = m.ExcPending
														if v486 != 0 {
															return int32(0)
														} else {
															F_errmsg(m, int32(_a_F_timestamp_zone_0), int32(0))
															mBase = m.M
															v490 = m.ExcPending
															if v490 != 0 {
																return int32(0)
															} else {
																F_errfinish(m, int32(_a_F_timestamp_zone_1), int32(_a_F_timestamp_zone_12), int32(_a_F_timestamp_zone_3))
																mBase = m.M
																v495 = m.ExcPending
																if v495 != 0 {
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
													v472 = v464 + base.I64_extend_i32_s(int32(0)-v345)*int64(-1000000)
													if base.Ui64(v472+int64(211813488000000000)) < base.Ui64(int64(-9011559254509551616)) {
														v503 = v472
														if base.Ui64(int64(-9011559254509551616)) <= base.Ui64(v503+int64(211813488000000000)) {
															F_errstart_cold(m, int32(21), int32(0))
															mBase = m.M
															v558 = m.ExcPending
															if v558 != 0 {
																return int32(0)
															} else {
																F_errcode(m, int32(134217858))
																mBase = m.M
																v561 = m.ExcPending
																if v561 != 0 {
																	return int32(0)
																} else {
																	F_errmsg(m, int32(_a_F_timestamp_zone_0), int32(0))
																	mBase = m.M
																	v565 = m.ExcPending
																	if v565 != 0 {
																		return int32(0)
																	} else {
																		F_errfinish(m, int32(_a_F_timestamp_zone_1), int32(_a_F_timestamp_zone_2), int32(_a_F_timestamp_zone_3))
																		mBase = m.M
																		v570 = m.ExcPending
																		if v570 != 0 {
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
															v510 = F_Int64GetDatum(m, v503)
															mBase = m.M
															v511 = m.ExcPending
															if v511 != 0 {
																return int32(0)
															} else {
																v518 = v510
																m.G0 = v10 + int32(336)
																return v518
															}
														}
													} else {
														F_errstart_cold(m, int32(21), int32(0))
														mBase = m.M
														v483 = m.ExcPending
														if v483 != 0 {
															return int32(0)
														} else {
															F_errcode(m, int32(134217858))
															mBase = m.M
															v486 = m.ExcPending
															if v486 != 0 {
																return int32(0)
															} else {
																F_errmsg(m, int32(_a_F_timestamp_zone_0), int32(0))
																mBase = m.M
																v490 = m.ExcPending
																if v490 != 0 {
																	return int32(0)
																} else {
																	F_errfinish(m, int32(_a_F_timestamp_zone_1), int32(_a_F_timestamp_zone_12), int32(_a_F_timestamp_zone_3))
																	mBase = m.M
																	v495 = m.ExcPending
																	if v495 != 0 {
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
			}
		}
	}
}
