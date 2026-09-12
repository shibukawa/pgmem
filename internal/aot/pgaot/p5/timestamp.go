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
	var v5 int32
	_ = v5
	var v8 int64
	_ = v8
	var v16 int64
	_ = v16
	var v19 int32
	_ = v19
	if l1 <= l0 {
		v19 = int32(0)
	} else {
		v5 = int32(2147483647)
		v8 = l1 - l0
		if base.B2i32(int64(0) < l0)^base.B2i32(v8 < l1) != 0 {
			v19 = v5
		} else {
			if int64(2147483646000) < v8 {
				v19 = v5
			} else {
				v16 = base.I64_div_s(v8+int64(999), int64(1000))
				v19 = base.I32_wrap_i64(v16)
			}
		}
	}
	return v19
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
	var v236 int32
	_ = v236
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
	var v261 int32
	_ = v261
	var v265 int32
	_ = v265
	var v268 int32
	_ = v268
	var v272 int32
	_ = v272
	var v277 int32
	_ = v277
	var v281 int32
	_ = v281
	var v286 int32
	_ = v286
	var v289 int32
	_ = v289
	var v290 int32
	_ = v290
	var v298 int32
	_ = v298
	var v300 int32
	_ = v300
	var v315 int32
	_ = v315
	var v320 int32
	_ = v320
	var v321 int32
	_ = v321
	var v322 int32
	_ = v322
	var v327 int32
	_ = v327
	var v330 int32
	_ = v330
	var v333 int32
	_ = v333
	var v337 int32
	_ = v337
	var v342 int32
	_ = v342
	var v348 int64
	_ = v348
	var v357 int64
	_ = v357
	var v358 int64
	_ = v358
	var v360 int64
	_ = v360
	var v363 int64
	_ = v363
	var v364 int64
	_ = v364
	var v366 int64
	_ = v366
	var v367 int64
	_ = v367
	var v371 int64
	_ = v371
	var v378 int64
	_ = v378
	var v389 int64
	_ = v389
	var v390 int64
	_ = v390
	var v398 int32
	_ = v398
	var v399 int32
	_ = v399
	var v400 int32
	_ = v400
	var v401 int32
	_ = v401
	var v410 int64
	_ = v410
	var v413 int64
	_ = v413
	var v424 int64
	_ = v424
	var v427 int32
	_ = v427
	var v429 int64
	_ = v429
	var v437 int64
	_ = v437
	var v438 int64
	_ = v438
	var v441 int64
	_ = v441
	var v444 int32
	_ = v444
	var v448 int32
	_ = v448
	var v450 int32
	_ = v450
	var v452 int32
	_ = v452
	var v456 int32
	_ = v456
	var v457 int32
	_ = v457
	var v458 int32
	_ = v458
	var v459 int32
	_ = v459
	var v465 int32
	_ = v465
	var v470 int32
	_ = v470
	var v473 int32
	_ = v473
	var v474 int32
	_ = v474
	var v475 int32
	_ = v475
	var v478 int32
	_ = v478
	var v480 int32
	_ = v480
	var v486 int32
	_ = v486
	var v490 int32
	_ = v490
	var v491 int32
	_ = v491
	var v493 int32
	_ = v493
	var v501 int32
	_ = v501
	var v505 int32
	_ = v505
	var v515 int32
	_ = v515
	var v525 int64
	_ = v525
	var v527 int64
	_ = v527
	var v532 int64
	_ = v532
	var v534 int64
	_ = v534
	var v539 int64
	_ = v539
	var v541 int64
	_ = v541
	var v545 int32
	_ = v545
	var v546 int32
	_ = v546
	var v547 int32
	_ = v547
	var v552 int32
	_ = v552
	var v553 int32
	_ = v553
	var v554 int32
	_ = v554
	var v559 int32
	_ = v559
	var v562 int32
	_ = v562
	var v565 int32
	_ = v565
	var v569 int32
	_ = v569
	var v574 int32
	_ = v574
	var v577 int32
	_ = v577
	var v578 int32
	_ = v578
	var v581 int32
	_ = v581
	var v589 int32
	_ = v589
	var v590 int32
	_ = v590
	var v591 int32
	_ = v591
	var v592 int32
	_ = v592
	var v598 int32
	_ = v598
	var v603 int32
	_ = v603
	var v606 int32
	_ = v606
	var v607 int32
	_ = v607
	var v608 int32
	_ = v608
	var v611 int32
	_ = v611
	var v613 int32
	_ = v613
	var v619 int32
	_ = v619
	var v623 int32
	_ = v623
	var v624 int32
	_ = v624
	var v626 int32
	_ = v626
	var v634 int32
	_ = v634
	var v638 int32
	_ = v638
	var v648 int32
	_ = v648
	var v652 int32
	_ = v652
	var v657 int32
	_ = v657
	var v662 int32
	_ = v662
	var v665 int32
	_ = v665
	var v668 int32
	_ = v668
	var v669 int32
	_ = v669
	var v674 int32
	_ = v674
	var v675 int32
	_ = v675
	var v676 int32
	_ = v676
	var v681 int32
	_ = v681
	var v684 int32
	_ = v684
	var v687 int32
	_ = v687
	var v691 int32
	_ = v691
	var v696 int32
	_ = v696
	var v702 int64
	_ = v702
	var v711 int64
	_ = v711
	var v712 int64
	_ = v712
	var v714 int64
	_ = v714
	var v717 int64
	_ = v717
	var v718 int64
	_ = v718
	var v720 int64
	_ = v720
	var v721 int64
	_ = v721
	var v725 int64
	_ = v725
	var v732 int64
	_ = v732
	var v743 int64
	_ = v743
	var v744 int64
	_ = v744
	var v752 int32
	_ = v752
	var v753 int32
	_ = v753
	var v754 int32
	_ = v754
	var v755 int32
	_ = v755
	var v764 int64
	_ = v764
	var v767 int64
	_ = v767
	var v779 int64
	_ = v779
	var v782 int64
	_ = v782
	var v785 int64
	_ = v785
	var v798 int64
	_ = v798
	var v800 int32
	_ = v800
	var v801 int32
	_ = v801
	var v809 int32
	_ = v809
	var v812 int32
	_ = v812
	var v816 int32
	_ = v816
	var v821 int32
	_ = v821
	var v825 int32
	_ = v825
	var v828 int32
	_ = v828
	var v832 int32
	_ = v832
	var v837 int32
	_ = v837
	var v841 int32
	_ = v841
	var v844 int32
	_ = v844
	var v848 int32
	_ = v848
	var v853 int32
	_ = v853
	var v857 int32
	_ = v857
	var v860 int32
	_ = v860
	var v864 int32
	_ = v864
	var v869 int32
	_ = v869
	var v873 int32
	_ = v873
	var v876 int32
	_ = v876
	var v880 int32
	_ = v880
	var v885 int32
	_ = v885
	var v889 int32
	_ = v889
	var v892 int32
	_ = v892
	var v896 int32
	_ = v896
	var v901 int32
	_ = v901
	var v908 int32
	_ = v908
	var v911 int32
	_ = v911
	var v915 int32
	_ = v915
	var v920 int32
	_ = v920
	var v926 int32
	_ = v926
	var v929 int32
	_ = v929
	var v933 int32
	_ = v933
	var v938 int32
	_ = v938
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
	v926 = m.ExcPending
	if v926 != 0 {
		goto L18
	} else {
		goto L168
	}
L2:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v908 = m.ExcPending
	if v908 != 0 {
		goto L18
	} else {
		goto L164
	}
L3:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v889 = m.ExcPending
	if v889 != 0 {
		goto L18
	} else {
		goto L160
	}
L4:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v873 = m.ExcPending
	if v873 != 0 {
		goto L18
	} else {
		goto L156
	}
L5:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v857 = m.ExcPending
	if v857 != 0 {
		goto L18
	} else {
		goto L152
	}
L6:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v841 = m.ExcPending
	if v841 != 0 {
		goto L18
	} else {
		goto L148
	}
L7:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v825 = m.ExcPending
	if v825 != 0 {
		goto L18
	} else {
		goto L144
	}
L8:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v809 = m.ExcPending
	if v809 != 0 {
		goto L18
	} else {
		goto L140
	}
L9:
	;
	v800 = F_Int64GetDatum(m, v798)
	mBase = m.M
	v801 = m.ExcPending
	if v801 != 0 {
		goto L18
	} else {
		goto L139
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
		v798 = v26
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
	F_errmsg(m, int32(422894), int32(0))
	mBase = m.M
	v44 = m.ExcPending
	if v44 != 0 {
		goto L18
	} else {
		goto L21
	}
L21:
	;
	F_errfinish(m, int32(521665), int32(3125), int32(324995))
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
		v798 = v53
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
	F_errmsg(m, int32(422894), int32(0))
	mBase = m.M
	v69 = m.ExcPending
	if v69 != 0 {
		goto L18
	} else {
		goto L28
	}
L28:
	;
	F_errfinish(m, int32(521665), int32(3134), int32(324995))
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
	v798 = v16
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
	v424 = v16
	goto L35
L35:
	;
	v427 = *(*int32)(unsafe.Add(mBase, uint32(v17)+8))
	if v427 != 0 {
		goto L86
	} else {
		goto L87
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
	v108 = v96 + int32(2483589)
	v109 = int32(146097)
	v110 = base.I32_div_u_s(v108, v109)
	v111 = int32(3)
	v117 = int32(2)
	v122 = base.I32_div_u_s((v110*int32(1073595727)+v108)<<(uint(v117)%32)|v111, v109)
	v125 = v96 + int32(2451545) + v110*v111 + v122 + int32(32104)
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
	*(*int32)(unsafe.Add(mBase, uint32(v13+int32(56)))) = v145 + v127<<(uint(int32(2))%32) - int32(4800)
	v153 = v143 + int32(123)
	v157 = int32(base.Ui32(v153*int32(2141)) >> (uint(int32(16)) % 32))
	*(*int32)(unsafe.Add(mBase, uint32(v13+int32(48)))) = v153 - int32(base.Ui32(v157*int32(7834))>>(uint(int32(8))%32))
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
	if v241&int32(3) == int32(0) {
		goto L57
	} else {
		goto L58
	}
L50:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v13)+52)) = v238
	v240 = v238
	v241 = v236
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
	v236 = v212
	v238 = v208 - v210*v209 + v207
	goto L50
L52:
	;
	goto L53
L53:
	;
	v219 = *(*int32)(unsafe.Add(mBase, uint32(v13)+56))
	if int32(0) < v199 {
		v240 = v199
		v241 = v219
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
	v236 = v228
	v238 = v225*v224 + v199 + v224
	goto L50
L55:
	;
	if v241 <= int32(-4713) {
		goto L67
	} else {
		goto L68
	}
L56:
	;
	v298 = *(*int32)(unsafe.Add(mBase, uint32(v290*int32(52)+int32(1695056)+v289<<(uint(int32(2))%32))))
	*(*int32)(unsafe.Add(mBase, uint32(v13)+48)) = v298
	v300 = v298
	goto L55
L57:
	;
	v249 = base.I32_rem_s(v241, int32(100))
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
	v281 = v240 - int32(1)
	v286 = *(*int32)(unsafe.Add(mBase, uint32(v281<<(uint(int32(2))%32))+uint32(_consts[1104])))
	if v243 <= v286 {
		v300 = v243
		goto L55
	} else {
		goto L65
	}
L60:
	;
	v253 = base.I32_rem_s(v241, int32(400))
	v261 = v240 - int32(1)
	v265 = *(*int32)(unsafe.Add(mBase, uint32(base.B2i32(v253 == int32(0))*int32(52)+int32(1695056)+v261<<(uint(int32(2))%32))))
	if v243 <= v265 {
		v300 = v243
		goto L55
	} else {
		goto L63
	}
L61:
	;
	goto L62
L62:
	;
	v272 = v240 - int32(1)
	v277 = *(*int32)(unsafe.Add(mBase, uint32(v272<<(uint(int32(2))%32))+uint32(_consts[1103])))
	if v243 <= v277 {
		v300 = v243
		goto L55
	} else {
		goto L64
	}
L63:
	;
	v268 = base.I32_rem_s(v241, int32(400))
	v289 = v261
	v290 = base.B2i32(v268 == int32(0))
	goto L56
L64:
	;
	v289 = v272
	v290 = int32(1)
	goto L56
L65:
	;
	v289 = v281
	v290 = int32(0)
	goto L56
L66:
	;
	v315 = v13 + int32(16)
	v320 = base.B2i32(int32(2) < v240)
	if int32(2) < v240 {
		goto L76
	} else {
		goto L77
	}
L67:
	;
	if v241 != int32(-4713) {
		goto L1
	} else {
		goto L70
	}
L68:
	;
	goto L69
L69:
	;
	if v241 < int32(5874898) {
		goto L66
	} else {
		goto L72
	}
L70:
	;
	if base.Ui32(int32(10)) < base.Ui32(v240) {
		goto L66
	} else {
		goto L71
	}
L71:
	;
	goto L1
L72:
	;
	if v241 != int32(5874898) {
		goto L1
	} else {
		goto L73
	}
L73:
	;
	if base.Ui32(int32(5)) < base.Ui32(v240) {
		goto L1
	} else {
		goto L74
	}
L74:
	;
	goto L66
L75:
	;
	v348 = base.I64_extend_i32_s(v300 + v322*int32(365) + v327 + v330 + v333 + v342 - int32(32167) - int32(2451545))
	v357 = int64(32)
	v358 = int64(20)
	v360 = int64(base.Ui64(v348) >> (uint(v357) % 64))
	v363 = int64(4294967295)
	v364 = int64(500654080)
	v366 = v348 & v363
	v367 = v364 * v366
	v371 = int64(base.Ui64(v367)>>(uint(v357)%64)) + v364*v360
	v378 = v366*v358 + v371&v363
	*(*int64)(unsafe.Add(mBase, uint32(v315)+8)) = v348*int64(0) + v348>>(uint(int64(63))%64)*int64(86400000000) + v358*v360 + int64(base.Ui64(v371)>>(uint(v357)%64)) + int64(base.Ui64(v378)>>(uint(v357)%64))
	*(*int64)(unsafe.Add(mBase, uint32(v315))) = v367&v363 | v378<<(uint(v357)%64)
	goto L82
L76:
	;
	v321 = int32(4800)
	goto L78
L77:
	;
	v321 = int32(4799)
	goto L78
L78:
	;
	v322 = v321 + v241
	v327 = base.I32_div_s(v322, int32(4))
	v330 = base.I32_div_s(v322, int32(-100))
	v333 = base.I32_div_s(v322, int32(400))
	if int32(2) < v240 {
		goto L79
	} else {
		goto L80
	}
L79:
	;
	v337 = int32(1)
	goto L81
L80:
	;
	v337 = int32(13)
	goto L81
L81:
	;
	v342 = base.I32_div_s((v337+v240)*int32(7834), int32(256))
	goto L75
L82:
	;
	v389 = *(*int64)(unsafe.Add(mBase, uint32(v13)+24))
	v390 = *(*int64)(unsafe.Add(mBase, uint32(v13)+16))
	if v389 != v390>>(uint(int64(63))%64) {
		goto L1
	} else {
		goto L83
	}
L83:
	;
	v398 = *(*int32)(unsafe.Add(mBase, uint32(v13)+36))
	v399 = *(*int32)(unsafe.Add(mBase, uint32(v13)+40))
	v400 = *(*int32)(unsafe.Add(mBase, uint32(v13)+44))
	v401 = int32(60)
	v410 = base.I64_extend32_s(v193*int64(4293967296)+v191) + base.I64_extend_i32_s(v398+(v399+v400*v401)*v401)*int64(1000000)
	v413 = v410 + v390
	if base.B2i32(v410 < int64(0))^base.B2i32(v413 < v390) != 0 {
		goto L1
	} else {
		goto L84
	}
L84:
	;
	if base.Ui64(v413-int64(9223371331200000000)) <= base.Ui64(int64(9011559254509551615)) {
		goto L1
	} else {
		goto L85
	}
L85:
	;
	v424 = v413
	goto L35
L86:
	;
	v429 = base.I64_div_s(v424, int64(86400000000))
	if base.Ui64(int64(172799999999)) <= base.Ui64(v424+int64(86399999999)) {
		goto L89
	} else {
		goto L90
	}
L87:
	;
	v779 = v424
	goto L88
L88:
	;
	v782 = *(*int64)(unsafe.Add(mBase, uint32(v17)))
	v785 = v779 + v782
	if base.B2i32(v782 < int64(0)) != base.B2i32(v785 < v779) {
		goto L4
	} else {
		goto L137
	}
L89:
	;
	v437 = v429 * int64(-86400000000)
	goto L91
L90:
	;
	v437 = int64(0)
	goto L91
L91:
	;
	v438 = v437 + v424
	v441 = v438>>(uint(int64(63))%64) + v429
	if v441 <= int64(-2451546) {
		goto L6
	} else {
		goto L92
	}
L92:
	;
	v444 = base.I32_wrap_i64(v441)
	v448 = v13 + int32(56)
	v450 = v13 + int32(52)
	v452 = v13 + int32(48)
	v456 = v444 + int32(2483589)
	v457 = int32(146097)
	v458 = base.I32_div_u_s(v456, v457)
	v459 = int32(3)
	v465 = int32(2)
	v470 = base.I32_div_u_s((v458*int32(1073595727)+v456)<<(uint(v465)%32)|v459, v457)
	v473 = v444 + int32(2451545) + v458*v459 + v470 + int32(32104)
	v474 = int32(1461)
	v475 = base.I32_div_u_s(v473, v474)
	v478 = v475*int32(-1461) + v473
	v480 = v478 << (uint(v465) % 32)
	if base.Ui32(v474) <= base.Ui32(v480) {
		goto L95
	} else {
		goto L96
	}
L93:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v13)+68)) = int64(4294967295)
	if v438 < int64(0) {
		goto L98
	} else {
		goto L99
	}
L94:
	;
	v493 = base.I32_div_u_s(v480, int32(1461))
	*(*int32)(unsafe.Add(mBase, uint32(v448))) = v493 + v475<<(uint(int32(2))%32) - int32(4800)
	v501 = v491 + int32(123)
	v505 = int32(base.Ui32(v501*int32(2141)) >> (uint(int32(16)) % 32))
	*(*int32)(unsafe.Add(mBase, uint32(v452))) = v501 - int32(base.Ui32(v505*int32(7834))>>(uint(int32(8))%32))
	v515 = base.I32_rem_u_s(v505+int32(10), int32(12))
	*(*int32)(unsafe.Add(mBase, uint32(v450))) = v515 + int32(1)
	goto L93
L95:
	;
	v486 = base.I32_rem_u_s(v478+int32(305), int32(365))
	v491 = v486
	goto L94
L96:
	;
	goto L97
L97:
	;
	v490 = base.I32_rem_u_s(v478+int32(306), int32(366))
	v491 = v490
	goto L94
L98:
	;
	v525 = v438 + int64(86400000000)
	goto L100
L99:
	;
	v525 = v438
	goto L100
L100:
	;
	v527 = base.I64_div_s(v525, int64(3600000000))
	*(*uint32)(unsafe.Add(mBase, uint32(v13)+44)) = uint32(v527)
	v532 = base.I64_extend32_s(v527)*int64(-3600000000) + v525
	v534 = base.I64_div_s(v532, int64(60000000))
	*(*uint32)(unsafe.Add(mBase, uint32(v13)+40)) = uint32(v534)
	v539 = base.I64_extend32_s(v534)*int64(-60000000) + v532
	v541 = base.I64_div_s(v539, int64(1000000))
	*(*uint32)(unsafe.Add(mBase, uint32(v13)+36)) = uint32(v541)
	*(*int32)(unsafe.Add(mBase, uint32(v13)+76)) = int32(0)
	v545 = *(*int32)(unsafe.Add(mBase, uint32(v13)+56))
	v546 = *(*int32)(unsafe.Add(mBase, uint32(v13)+52))
	v547 = *(*int32)(unsafe.Add(mBase, uint32(v13)+48))
	v552 = base.B2i32(int32(2) < v546)
	if int32(2) < v546 {
		goto L102
	} else {
		goto L103
	}
L101:
	;
	v578 = *(*int32)(unsafe.Add(mBase, uint32(v17)+8))
	v581 = v577 + v578
	if base.B2i32(v578 < int32(0))^base.B2i32(v581 < v577) != 0 {
		goto L5
	} else {
		goto L108
	}
L102:
	;
	v553 = int32(4800)
	goto L104
L103:
	;
	v553 = int32(4799)
	goto L104
L104:
	;
	v554 = v553 + v545
	v559 = base.I32_div_s(v554, int32(4))
	v562 = base.I32_div_s(v554, int32(-100))
	v565 = base.I32_div_s(v554, int32(400))
	if int32(2) < v546 {
		goto L105
	} else {
		goto L106
	}
L105:
	;
	v569 = int32(1)
	goto L107
L106:
	;
	v569 = int32(13)
	goto L107
L107:
	;
	v574 = base.I32_div_s((v569+v546)*int32(7834), int32(256))
	v577 = v547 + v554*int32(365) + v559 + v562 + v565 + v574 - int32(32167)
	goto L101
L108:
	;
	if v581 < int32(0) {
		goto L5
	} else {
		goto L109
	}
L109:
	;
	v589 = v581 + int32(32044)
	v590 = int32(146097)
	v591 = base.I32_div_u_s(v589, v590)
	v592 = int32(3)
	v598 = int32(2)
	v603 = base.I32_div_u_s((v591*int32(1073595727)+v589)<<(uint(v598)%32)|v592, v590)
	v606 = v581 + v591*v592 + v603 + int32(32104)
	v607 = int32(1461)
	v608 = base.I32_div_u_s(v606, v607)
	v611 = v608*int32(-1461) + v606
	v613 = v611 << (uint(v598) % 32)
	if base.Ui32(v607) <= base.Ui32(v613) {
		goto L112
	} else {
		goto L113
	}
L110:
	;
	v652 = *(*int32)(unsafe.Add(mBase, uint32(v13)+56))
	if v652 <= int32(-4713) {
		goto L116
	} else {
		goto L117
	}
L111:
	;
	v626 = base.I32_div_u_s(v613, int32(1461))
	*(*int32)(unsafe.Add(mBase, uint32(v448))) = v626 + v608<<(uint(int32(2))%32) - int32(4800)
	v634 = v624 + int32(123)
	v638 = int32(base.Ui32(v634*int32(2141)) >> (uint(int32(16)) % 32))
	*(*int32)(unsafe.Add(mBase, uint32(v452))) = v634 - int32(base.Ui32(v638*int32(7834))>>(uint(int32(8))%32))
	v648 = base.I32_rem_u_s(v638+int32(10), int32(12))
	*(*int32)(unsafe.Add(mBase, uint32(v450))) = v648 + int32(1)
	goto L110
L112:
	;
	v619 = base.I32_rem_u_s(v611+int32(305), int32(365))
	v624 = v619
	goto L111
L113:
	;
	goto L114
L114:
	;
	v623 = base.I32_rem_u_s(v611+int32(306), int32(366))
	v624 = v623
	goto L111
L115:
	;
	v669 = *(*int32)(unsafe.Add(mBase, uint32(v13)+48))
	v674 = base.B2i32(int32(2) < v668)
	if int32(2) < v668 {
		goto L127
	} else {
		goto L128
	}
L116:
	;
	if v652 != int32(-4713) {
		goto L2
	} else {
		goto L119
	}
L117:
	;
	goto L118
L118:
	;
	if v652 <= int32(5874897) {
		goto L121
	} else {
		goto L122
	}
L119:
	;
	v657 = *(*int32)(unsafe.Add(mBase, uint32(v13)+52))
	if int32(10) < v657 {
		v668 = v657
		goto L115
	} else {
		goto L120
	}
L120:
	;
	goto L2
L121:
	;
	v662 = *(*int32)(unsafe.Add(mBase, uint32(v13)+52))
	v668 = v662
	goto L115
L122:
	;
	goto L123
L123:
	;
	if v652 != int32(5874898) {
		goto L2
	} else {
		goto L124
	}
L124:
	;
	v665 = *(*int32)(unsafe.Add(mBase, uint32(v13)+52))
	if int32(5) < v665 {
		goto L2
	} else {
		goto L125
	}
L125:
	;
	v668 = v665
	goto L115
L126:
	;
	v702 = base.I64_extend_i32_s(v669 + v676*int32(365) + v681 + v684 + v687 + v696 - int32(32167) - int32(2451545))
	v711 = int64(32)
	v712 = int64(20)
	v714 = int64(base.Ui64(v702) >> (uint(v711) % 64))
	v717 = int64(4294967295)
	v718 = int64(500654080)
	v720 = v702 & v717
	v721 = v718 * v720
	v725 = int64(base.Ui64(v721)>>(uint(v711)%64)) + v718*v714
	v732 = v720*v712 + v725&v717
	*(*int64)(unsafe.Add(mBase, uint32(v13)+8)) = v702*int64(0) + v702>>(uint(int64(63))%64)*int64(86400000000) + v712*v714 + int64(base.Ui64(v725)>>(uint(v711)%64)) + int64(base.Ui64(v732)>>(uint(v711)%64))
	*(*int64)(unsafe.Add(mBase, uint32(v13))) = v721&v717 | v732<<(uint(v711)%64)
	goto L133
L127:
	;
	v675 = int32(4800)
	goto L129
L128:
	;
	v675 = int32(4799)
	goto L129
L129:
	;
	v676 = v675 + v652
	v681 = base.I32_div_s(v676, int32(4))
	v684 = base.I32_div_s(v676, int32(-100))
	v687 = base.I32_div_s(v676, int32(400))
	if int32(2) < v668 {
		goto L130
	} else {
		goto L131
	}
L130:
	;
	v691 = int32(1)
	goto L132
L131:
	;
	v691 = int32(13)
	goto L132
L132:
	;
	v696 = base.I32_div_s((v691+v668)*int32(7834), int32(256))
	goto L126
L133:
	;
	v743 = *(*int64)(unsafe.Add(mBase, uint32(v13)+8))
	v744 = *(*int64)(unsafe.Add(mBase, uint32(v13)))
	if v743 != v744>>(uint(int64(63))%64) {
		goto L2
	} else {
		goto L134
	}
L134:
	;
	v752 = *(*int32)(unsafe.Add(mBase, uint32(v13)+36))
	v753 = *(*int32)(unsafe.Add(mBase, uint32(v13)+40))
	v754 = *(*int32)(unsafe.Add(mBase, uint32(v13)+44))
	v755 = int32(60)
	v764 = base.I64_extend32_s(v541*int64(4293967296)+v539) + base.I64_extend_i32_s(v752+(v753+v754*v755)*v755)*int64(1000000)
	v767 = v764 + v744
	if base.B2i32(v764 < int64(0))^base.B2i32(v767 < v744) != 0 {
		goto L2
	} else {
		goto L135
	}
L135:
	;
	if base.Ui64(v767-int64(9223371331200000000)) <= base.Ui64(int64(9011559254509551615)) {
		goto L2
	} else {
		goto L136
	}
L136:
	;
	v779 = v767
	goto L88
L137:
	;
	if base.Ui64(int64(-9011559254509551616)) <= base.Ui64(v785+int64(211813488000000000)) {
		goto L3
	} else {
		goto L138
	}
L138:
	;
	v798 = v785
	goto L9
L139:
	;
	m.G0 = v13 + int32(80)
	return v800
L140:
	;
	F_errcode(m, int32(134217858))
	mBase = m.M
	v812 = m.ExcPending
	if v812 != 0 {
		goto L18
	} else {
		goto L141
	}
L141:
	;
	F_errmsg(m, int32(422894), int32(0))
	mBase = m.M
	v816 = m.ExcPending
	if v816 != 0 {
		goto L18
	} else {
		goto L142
	}
L142:
	;
	F_errfinish(m, int32(521665), int32(3151), int32(324995))
	mBase = m.M
	v821 = m.ExcPending
	if v821 != 0 {
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
	v828 = m.ExcPending
	if v828 != 0 {
		goto L18
	} else {
		goto L145
	}
L145:
	;
	F_errmsg(m, int32(422894), int32(0))
	mBase = m.M
	v832 = m.ExcPending
	if v832 != 0 {
		goto L18
	} else {
		goto L146
	}
L146:
	;
	F_errfinish(m, int32(521665), int32(3156), int32(324995))
	mBase = m.M
	v837 = m.ExcPending
	if v837 != 0 {
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
	v844 = m.ExcPending
	if v844 != 0 {
		goto L18
	} else {
		goto L149
	}
L149:
	;
	F_errmsg(m, int32(422894), int32(0))
	mBase = m.M
	v848 = m.ExcPending
	if v848 != 0 {
		goto L18
	} else {
		goto L150
	}
L150:
	;
	F_errfinish(m, int32(521665), int32(3188), int32(324995))
	mBase = m.M
	v853 = m.ExcPending
	if v853 != 0 {
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
	v860 = m.ExcPending
	if v860 != 0 {
		goto L18
	} else {
		goto L153
	}
L153:
	;
	F_errmsg(m, int32(422894), int32(0))
	mBase = m.M
	v864 = m.ExcPending
	if v864 != 0 {
		goto L18
	} else {
		goto L154
	}
L154:
	;
	F_errfinish(m, int32(521665), int32(3199), int32(324995))
	mBase = m.M
	v869 = m.ExcPending
	if v869 != 0 {
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
	v876 = m.ExcPending
	if v876 != 0 {
		goto L18
	} else {
		goto L157
	}
L157:
	;
	F_errmsg(m, int32(422894), int32(0))
	mBase = m.M
	v880 = m.ExcPending
	if v880 != 0 {
		goto L18
	} else {
		goto L158
	}
L158:
	;
	F_errfinish(m, int32(521665), int32(3211), int32(324995))
	mBase = m.M
	v885 = m.ExcPending
	if v885 != 0 {
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
	v892 = m.ExcPending
	if v892 != 0 {
		goto L18
	} else {
		goto L161
	}
L161:
	;
	F_errmsg(m, int32(422894), int32(0))
	mBase = m.M
	v896 = m.ExcPending
	if v896 != 0 {
		goto L18
	} else {
		goto L162
	}
L162:
	;
	F_errfinish(m, int32(521665), int32(3216), int32(324995))
	mBase = m.M
	v901 = m.ExcPending
	if v901 != 0 {
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
	v911 = m.ExcPending
	if v911 != 0 {
		goto L18
	} else {
		goto L165
	}
L165:
	;
	F_errmsg(m, int32(422894), int32(0))
	mBase = m.M
	v915 = m.ExcPending
	if v915 != 0 {
		goto L18
	} else {
		goto L166
	}
L166:
	;
	F_errfinish(m, int32(521665), int32(3205), int32(324995))
	mBase = m.M
	v920 = m.ExcPending
	if v920 != 0 {
		goto L18
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
	F_errcode(m, int32(134217858))
	mBase = m.M
	v929 = m.ExcPending
	if v929 != 0 {
		goto L18
	} else {
		goto L169
	}
L169:
	;
	F_errmsg(m, int32(422894), int32(0))
	mBase = m.M
	v933 = m.ExcPending
	if v933 != 0 {
		goto L18
	} else {
		goto L170
	}
L170:
	;
	F_errfinish(m, int32(521665), int32(3175), int32(324995))
	mBase = m.M
	v938 = m.ExcPending
	if v938 != 0 {
		goto L18
	} else {
		goto L171
	}
L171:
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
	var v21 int32
	_ = v21
	var v23 int32
	_ = v23
	var v32 int32
	_ = v32
	var v35 int32
	_ = v35
	var v36 int32
	_ = v36
	var v42 int32
	_ = v42
	var v48 int32
	_ = v48
	var v54 int32
	_ = v54
	var v57 int32
	_ = v57
	var v59 int64
	_ = v59
	var v73 int32
	_ = v73
	var v80 int32
	_ = v80
	var v81 int32
	_ = v81
	var v82 int32
	_ = v82
	var v83 int32
	_ = v83
	var v84 int32
	_ = v84
	var v89 int32
	_ = v89
	var v90 int32
	_ = v90
	var v91 int32
	_ = v91
	var v96 int32
	_ = v96
	var v99 int32
	_ = v99
	var v102 int32
	_ = v102
	var v106 int32
	_ = v106
	var v111 int32
	_ = v111
	var v114 int32
	_ = v114
	var v115 int32
	_ = v115
	var v118 int32
	_ = v118
	var v120 int32
	_ = v120
	var v129 int32
	_ = v129
	var v134 int32
	_ = v134
	var v137 int32
	_ = v137
	var v140 int32
	_ = v140
	var v149 int32
	_ = v149
	var v153 int64
	_ = v153
	var v155 int64
	_ = v155
	var v157 int64
	_ = v157
	var v159 int64
	_ = v159
	var v161 int32
	_ = v161
	var v165 int32
	_ = v165
	var v171 int32
	_ = v171
	var v172 int32
	_ = v172
	var v173 int32
	_ = v173
	var v174 int32
	_ = v174
	var v181 int32
	_ = v181
	var v190 int32
	_ = v190
	var v193 int32
	_ = v193
	var v197 int32
	_ = v197
	var v202 int32
	_ = v202
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
			v21 = int32(4)
			v23 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v14)+1)))
			if v23&int32(254) == int32(2) {
				v32 = v21
			} else {
				v32 = base.B2i32(v23 == int32(18)) << (uint(v21) % 32)
			}
			if v23 == int32(1) {
				v35 = v21
			} else {
				v35 = v32
			}
			v48 = v35
		} else {
			v36 = int32(1)
			if v18&v36 != 0 {
				v48 = int32(base.Ui32(v18)>>(uint(v36)%32)) - v36
			} else {
				v42 = *(*int32)(unsafe.Add(mBase, uint32(v14)))
				v48 = int32(base.Ui32(v42)>>(uint(int32(2))%32)) - int32(4)
			}
		}
		if base.Ui64(int64(1)) < base.Ui64(v12-int64(9223372036854775807)) {
			v54 = v48
		} else {
			v54 = int32(0)
		}
		if v54 == int32(0) {
			v57 = int32(1)
			*(*uint8)(unsafe.Add(mBase, uint32(l0)+16)) = uint8(v57)
			v181 = int32(0)
			m.G0 = v9 + int32(96)
			return v181
		} else {
			v59 = int64(0)
			*(*int64)(unsafe.Add(mBase, uint32(v9)+80)) = v59
			*(*int64)(unsafe.Add(mBase, uint32(v9)+72)) = v59
			*(*int64)(unsafe.Add(mBase, uint32(v9)+56)) = v59
			*(*int64)(unsafe.Add(mBase, uint32(v9-int32(-64)))) = int64(4294967297)
			*(*int64)(unsafe.Add(mBase, uint32(v9)+48)) = v59
			*(*int64)(unsafe.Add(mBase, uint32(v9)+88)) = v59
			v73 = int32(0)
			v80 = F_timestamp2tm(m, v12, v73, v9+int32(4), v9+int32(88), v73, v73)
			mBase = m.M
			v81 = m.ExcPending
			if v81 != 0 {
				return int32(0)
			} else {
				if v80 != 0 {
					F_errstart_cold(m, int32(21), int32(0))
					mBase = m.M
					v190 = m.ExcPending
					if v190 != 0 {
						return int32(0)
					} else {
						F_errcode(m, int32(134217858))
						mBase = m.M
						v193 = m.ExcPending
						if v193 != 0 {
							return int32(0)
						} else {
							F_errmsg(m, int32(422894), int32(0))
							mBase = m.M
							v197 = m.ExcPending
							if v197 != 0 {
								return int32(0)
							} else {
								F_errfinish(m, int32(524465), int32(4017), int32(242517))
								mBase = m.M
								v202 = m.ExcPending
								if v202 != 0 {
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
					v82 = *(*int32)(unsafe.Add(mBase, uint32(v9)+24))
					v83 = *(*int32)(unsafe.Add(mBase, uint32(v9)+20))
					v84 = *(*int32)(unsafe.Add(mBase, uint32(v9)+16))
					v89 = base.B2i32(int32(2) < v83)
					if int32(2) < v83 {
						v90 = int32(4800)
					} else {
						v90 = int32(4799)
					}
					v91 = v90 + v82
					v96 = base.I32_div_s(v91, int32(4))
					v99 = base.I32_div_s(v91, int32(-100))
					v102 = base.I32_div_s(v91, int32(400))
					if int32(2) < v83 {
						v106 = int32(1)
					} else {
						v106 = int32(13)
					}
					v111 = base.I32_div_s((v106+v83)*int32(7834), int32(256))
					v114 = v84 + v91*int32(365) + v96 + v99 + v102 + v111 - int32(32167)
					v115 = int32(1)
					v118 = base.I32_rem_s(v114+v115, int32(7))
					*(*int32)(unsafe.Add(mBase, uint32(v9)+28)) = v118
					v120 = *(*int32)(unsafe.Add(mBase, uint32(v9)+24))
					v129 = int32(4799) + v120
					v134 = base.I32_div_s(v129, int32(4))
					v137 = base.I32_div_s(v129, int32(-100))
					v140 = base.I32_div_s(v129, int32(400))
					v149 = base.I32_div_s(int32(109676), int32(256))
					v153 = *(*int64)(unsafe.Add(mBase, uint32(v9)+4))
					*(*int64)(unsafe.Add(mBase, uint32(v9)+48)) = v153
					v155 = int64(*(*int32)(unsafe.Add(mBase, uint32(v9)+12)))
					*(*int64)(unsafe.Add(mBase, uint32(v9)+56)) = v155
					v157 = *(*int64)(unsafe.Add(mBase, uint32(v9)+16))
					*(*int64)(unsafe.Add(mBase, uint32(v9)+64)) = v157
					v159 = *(*int64)(unsafe.Add(mBase, uint32(v9)+24))
					*(*int64)(unsafe.Add(mBase, uint32(v9)+72)) = v159
					v161 = *(*int32)(unsafe.Add(mBase, uint32(v9)+40))
					*(*int32)(unsafe.Add(mBase, uint32(v9)+84)) = v161
					v165 = v114 - (v115 + v129*int32(365) + v134 + v137 + v140 + v149 - int32(32167)) + int32(1)
					*(*int32)(unsafe.Add(mBase, uint32(v9)+32)) = v165
					*(*int32)(unsafe.Add(mBase, uint32(v9)+80)) = v165
					v171 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
					v172 = F_datetime_to_char_body(m, v9+int32(48), v14, int32(0), v171)
					mBase = m.M
					v173 = m.ExcPending
					if v173 != 0 {
						return int32(0)
					} else {
						if v172 != 0 {
							v181 = v172
						} else {
							v174 = int32(1)
							*(*uint8)(unsafe.Add(mBase, uint32(l0)+16)) = uint8(v174)
							v181 = int32(0)
						}
						m.G0 = v9 + int32(96)
						return v181
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
	var v29 int32
	_ = v29
	var v36 int32
	_ = v36
	var v37 int32
	_ = v37
	var v39 int64
	_ = v39
	var v47 int64
	_ = v47
	var v48 int64
	_ = v48
	var v51 int64
	_ = v51
	var v54 int32
	_ = v54
	var v66 int32
	_ = v66
	var v67 int32
	_ = v67
	var v68 int32
	_ = v68
	var v69 int32
	_ = v69
	var v75 int32
	_ = v75
	var v80 int32
	_ = v80
	var v83 int32
	_ = v83
	var v84 int32
	_ = v84
	var v85 int32
	_ = v85
	var v88 int32
	_ = v88
	var v90 int32
	_ = v90
	var v96 int32
	_ = v96
	var v100 int32
	_ = v100
	var v101 int32
	_ = v101
	var v103 int32
	_ = v103
	var v111 int32
	_ = v111
	var v115 int32
	_ = v115
	var v125 int32
	_ = v125
	var v135 int64
	_ = v135
	var v137 int64
	_ = v137
	var v142 int64
	_ = v142
	var v144 int64
	_ = v144
	var v151 int64
	_ = v151
	var v153 int32
	_ = v153
	var v157 int32
	_ = v157
	var v160 int32
	_ = v160
	var v164 int32
	_ = v164
	var v166 int32
	_ = v166
	var v170 int32
	_ = v170
	var v174 int32
	_ = v174
	var v175 int32
	_ = v175
	var v179 int32
	_ = v179
	var v183 int32
	_ = v183
	var v184 int32
	_ = v184
	var v187 int32
	_ = v187
	var v188 int32
	_ = v188
	var v203 int32
	_ = v203
	var v204 int32
	_ = v204
	var v205 int32
	_ = v205
	var v210 int32
	_ = v210
	var v220 int64
	_ = v220
	var v228 int64
	_ = v228
	var v229 int64
	_ = v229
	var v232 int64
	_ = v232
	var v235 int32
	_ = v235
	var v247 int32
	_ = v247
	var v248 int32
	_ = v248
	var v249 int32
	_ = v249
	var v250 int32
	_ = v250
	var v256 int32
	_ = v256
	var v261 int32
	_ = v261
	var v264 int32
	_ = v264
	var v265 int32
	_ = v265
	var v266 int32
	_ = v266
	var v269 int32
	_ = v269
	var v271 int32
	_ = v271
	var v277 int32
	_ = v277
	var v281 int32
	_ = v281
	var v282 int32
	_ = v282
	var v284 int32
	_ = v284
	var v292 int32
	_ = v292
	var v296 int32
	_ = v296
	var v306 int32
	_ = v306
	var v318 int64
	_ = v318
	var v320 int64
	_ = v320
	var v325 int64
	_ = v325
	var v327 int64
	_ = v327
	var v332 int64
	_ = v332
	var v334 int64
	_ = v334
	var v338 int32
	_ = v338
	var v340 int32
	_ = v340
	var v341 int32
	_ = v341
	var v342 int32
	_ = v342
	var v346 int32
	_ = v346
	var v350 int32
	_ = v350
	var v355 int32
	_ = v355
	var v360 int32
	_ = v360
	var v363 int32
	_ = v363
	var v366 int32
	_ = v366
	var v368 int32
	_ = v368
	var v369 int32
	_ = v369
	var v374 int32
	_ = v374
	var v375 int32
	_ = v375
	var v376 int32
	_ = v376
	var v381 int32
	_ = v381
	var v384 int32
	_ = v384
	var v387 int32
	_ = v387
	var v391 int32
	_ = v391
	var v396 int32
	_ = v396
	var v402 int64
	_ = v402
	var v411 int64
	_ = v411
	var v412 int64
	_ = v412
	var v414 int64
	_ = v414
	var v417 int64
	_ = v417
	var v418 int64
	_ = v418
	var v420 int64
	_ = v420
	var v421 int64
	_ = v421
	var v425 int64
	_ = v425
	var v432 int64
	_ = v432
	var v443 int64
	_ = v443
	var v444 int64
	_ = v444
	var v452 int32
	_ = v452
	var v453 int32
	_ = v453
	var v454 int32
	_ = v454
	var v455 int32
	_ = v455
	var v464 int64
	_ = v464
	var v467 int64
	_ = v467
	var v475 int64
	_ = v475
	var v486 int32
	_ = v486
	var v489 int32
	_ = v489
	var v493 int32
	_ = v493
	var v498 int32
	_ = v498
	var v499 int64
	_ = v499
	var v506 int64
	_ = v506
	var v513 int32
	_ = v513
	var v514 int32
	_ = v514
	var v521 int32
	_ = v521
	var v529 int32
	_ = v529
	var v532 int32
	_ = v532
	var v536 int32
	_ = v536
	var v541 int32
	_ = v541
	var v545 int32
	_ = v545
	var v548 int32
	_ = v548
	var v552 int32
	_ = v552
	var v557 int32
	_ = v557
	var v561 int32
	_ = v561
	var v564 int32
	_ = v564
	var v568 int32
	_ = v568
	var v573 int32
	_ = v573
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
				v521 = v23
				m.G0 = v10 + int32(336)
				return v521
			}
		} else {
			F_text_to_cstring_buffer(m, v13, v10+int32(80), int32(256))
			mBase = m.M
			v29 = m.ExcPending
			if v29 != 0 {
				return int32(0)
			} else {
				v36 = F_DecodeTimezoneName(m, v10+int32(80), v10+int32(76), v10+int32(72))
				mBase = m.M
				v37 = m.ExcPending
				if v37 != 0 {
					return int32(0)
				} else {
					switch v36 {
					case 0:
						v499 = int64(*(*int32)(unsafe.Add(mBase, uint32(v10)+76)))
						v506 = v499*int64(-1000000) + v18
						if base.Ui64(int64(-9011559254509551616)) <= base.Ui64(v506+int64(211813488000000000)) {
							F_errstart_cold(m, int32(21), int32(0))
							mBase = m.M
							v561 = m.ExcPending
							if v561 != 0 {
								return int32(0)
							} else {
								F_errcode(m, int32(134217858))
								mBase = m.M
								v564 = m.ExcPending
								if v564 != 0 {
									return int32(0)
								} else {
									F_errmsg(m, int32(422894), int32(0))
									mBase = m.M
									v568 = m.ExcPending
									if v568 != 0 {
										return int32(0)
									} else {
										F_errfinish(m, int32(521665), int32(6382), int32(392006))
										mBase = m.M
										v573 = m.ExcPending
										if v573 != 0 {
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
							v513 = F_Int64GetDatum(m, v506)
							mBase = m.M
							v514 = m.ExcPending
							if v514 != 0 {
								return int32(0)
							} else {
								v521 = v513
								m.G0 = v10 + int32(336)
								return v521
							}
						}
					case 1:
						v39 = base.I64_div_s(v18, int64(86400000000))
						if base.Ui64(int64(172799999999)) <= base.Ui64(v18+int64(86399999999)) {
							v47 = v39 * int64(-86400000000)
						} else {
							v47 = int64(0)
						}
						v48 = v47 + v18
						v51 = v48>>(uint(int64(63))%64) + v39
						if v51 <= int64(-2451546) {
							F_errstart_cold(m, int32(21), int32(0))
							mBase = m.M
							v529 = m.ExcPending
							if v529 != 0 {
								return int32(0)
							} else {
								F_errcode(m, int32(134217858))
								mBase = m.M
								v532 = m.ExcPending
								if v532 != 0 {
									return int32(0)
								} else {
									F_errmsg(m, int32(422894), int32(0))
									mBase = m.M
									v536 = m.ExcPending
									if v536 != 0 {
										return int32(0)
									} else {
										F_errfinish(m, int32(521665), int32(6361), int32(392006))
										mBase = m.M
										v541 = m.ExcPending
										if v541 != 0 {
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
							v54 = base.I32_wrap_i64(v51)
							v66 = v54 + int32(2483589)
							v67 = int32(146097)
							v68 = base.I32_div_u_s(v66, v67)
							v69 = int32(3)
							v75 = int32(2)
							v80 = base.I32_div_u_s((v68*int32(1073595727)+v66)<<(uint(v75)%32)|v69, v67)
							v83 = v54 + int32(2451545) + v68*v69 + v80 + int32(32104)
							v84 = int32(1461)
							v85 = base.I32_div_u_s(v83, v84)
							v88 = v85*int32(-1461) + v83
							v90 = v88 << (uint(v75) % 32)
							if base.Ui32(v84) <= base.Ui32(v90) {
								v96 = base.I32_rem_u_s(v88+int32(305), int32(365))
								v101 = v96
							} else {
								v100 = base.I32_rem_u_s(v88+int32(306), int32(366))
								v101 = v100
							}
							v103 = base.I32_div_u_s(v90, int32(1461))
							*(*int32)(unsafe.Add(mBase, uint32(v10+int32(48)))) = v103 + v85<<(uint(int32(2))%32) - int32(4800)
							v111 = v101 + int32(123)
							v115 = int32(base.Ui32(v111*int32(2141)) >> (uint(int32(16)) % 32))
							*(*int32)(unsafe.Add(mBase, uint32(v10+int32(40)))) = v111 - int32(base.Ui32(v115*int32(7834))>>(uint(int32(8))%32))
							v125 = base.I32_rem_u_s(v115+int32(10), int32(12))
							*(*int32)(unsafe.Add(mBase, uint32(v10+int32(44)))) = v125 + int32(1)
							*(*int64)(unsafe.Add(mBase, uint32(v10)+60)) = int64(4294967295)
							if v48 < int64(0) {
								v135 = v48 + int64(86400000000)
							} else {
								v135 = v48
							}
							v137 = base.I64_div_s(v135, int64(3600000000))
							*(*uint32)(unsafe.Add(mBase, uint32(v10)+36)) = uint32(v137)
							v142 = base.I64_extend32_s(v137)*int64(-3600000000) + v135
							v144 = base.I64_div_s(v142, int64(60000000))
							*(*uint32)(unsafe.Add(mBase, uint32(v10)+32)) = uint32(v144)
							v151 = base.I64_div_s(base.I64_extend32_s(v144)*int64(-60000000)+v142, int64(1000000))
							*(*uint32)(unsafe.Add(mBase, uint32(v10)+28)) = uint32(v151)
							v153 = int32(0)
							*(*int32)(unsafe.Add(mBase, uint32(v10)+68)) = v153
							v157 = v10 + int32(28)
							v160 = *(*int32)(unsafe.Add(mBase, uint32(v10)+72))
							v164 = m.G0
							v166 = v164 - int32(288)
							m.G0 = v166
							v170 = F_DetermineTimeZoneOffsetInternal(m, v157, v160, v166+int32(280))
							mBase = m.M
							v174 = F_strlcpy(m, v166+int32(16), v10+int32(80), int32(256))
							mBase = m.M
							v175 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v166)+16)))
							if v175 != 0 {
								v179 = v166 + int32(16)
								v183 = v175
								for {
									v184 = F_pg_toupper(m, v183)
									mBase = m.M
									*(*uint8)(unsafe.Add(mBase, uint32(v179))) = uint8(v184)
									v187 = v179 + int32(1)
									v188 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v187))))
									if v188 != 0 {
										v179 = v187
										v183 = v188
										continue
									} else {
										break
									}
									break
								}
							} else {
							}
							v203 = F_pg_interpret_timezone_abbrev(m, v166+int32(16), v166+int32(280), v166+int32(12), v166+int32(8), v160)
							mBase = m.M
							if v203 != 0 {
								v204 = *(*int32)(unsafe.Add(mBase, uint32(v166)+12))
								v205 = *(*int32)(unsafe.Add(mBase, uint32(v166)+8))
								*(*int32)(unsafe.Add(mBase, uint32(v157)+32)) = v205
								v210 = int32(0) - v204
							} else {
								v210 = v170
							}
							m.G0 = v166 + int32(288)
							v506 = base.I64_extend_i32_s(v153-v210)*int64(-1000000) + v18
							if base.Ui64(int64(-9011559254509551616)) <= base.Ui64(v506+int64(211813488000000000)) {
								F_errstart_cold(m, int32(21), int32(0))
								mBase = m.M
								v561 = m.ExcPending
								if v561 != 0 {
									return int32(0)
								} else {
									F_errcode(m, int32(134217858))
									mBase = m.M
									v564 = m.ExcPending
									if v564 != 0 {
										return int32(0)
									} else {
										F_errmsg(m, int32(422894), int32(0))
										mBase = m.M
										v568 = m.ExcPending
										if v568 != 0 {
											return int32(0)
										} else {
											F_errfinish(m, int32(521665), int32(6382), int32(392006))
											mBase = m.M
											v573 = m.ExcPending
											if v573 != 0 {
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
								v513 = F_Int64GetDatum(m, v506)
								mBase = m.M
								v514 = m.ExcPending
								if v514 != 0 {
									return int32(0)
								} else {
									v521 = v513
									m.G0 = v10 + int32(336)
									return v521
								}
							}
						}
					default:
						v220 = base.I64_div_s(v18, int64(86400000000))
						if base.Ui64(int64(172799999999)) <= base.Ui64(v18+int64(86399999999)) {
							v228 = v220 * int64(-86400000000)
						} else {
							v228 = int64(0)
						}
						v229 = v228 + v18
						v232 = v229>>(uint(int64(63))%64) + v220
						if v232 <= int64(-2451546) {
							F_errstart_cold(m, int32(21), int32(0))
							mBase = m.M
							v545 = m.ExcPending
							if v545 != 0 {
								return int32(0)
							} else {
								F_errcode(m, int32(134217858))
								mBase = m.M
								v548 = m.ExcPending
								if v548 != 0 {
									return int32(0)
								} else {
									F_errmsg(m, int32(422894), int32(0))
									mBase = m.M
									v552 = m.ExcPending
									if v552 != 0 {
										return int32(0)
									} else {
										F_errfinish(m, int32(521665), int32(6371), int32(392006))
										mBase = m.M
										v557 = m.ExcPending
										if v557 != 0 {
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
							v235 = base.I32_wrap_i64(v232)
							v247 = v235 + int32(2483589)
							v248 = int32(146097)
							v249 = base.I32_div_u_s(v247, v248)
							v250 = int32(3)
							v256 = int32(2)
							v261 = base.I32_div_u_s((v249*int32(1073595727)+v247)<<(uint(v256)%32)|v250, v248)
							v264 = v235 + int32(2451545) + v249*v250 + v261 + int32(32104)
							v265 = int32(1461)
							v266 = base.I32_div_u_s(v264, v265)
							v269 = v266*int32(-1461) + v264
							v271 = v269 << (uint(v256) % 32)
							if base.Ui32(v265) <= base.Ui32(v271) {
								v277 = base.I32_rem_u_s(v269+int32(305), int32(365))
								v282 = v277
							} else {
								v281 = base.I32_rem_u_s(v269+int32(306), int32(366))
								v282 = v281
							}
							v284 = base.I32_div_u_s(v271, int32(1461))
							*(*int32)(unsafe.Add(mBase, uint32(v10+int32(48)))) = v284 + v266<<(uint(int32(2))%32) - int32(4800)
							v292 = v282 + int32(123)
							v296 = int32(base.Ui32(v292*int32(2141)) >> (uint(int32(16)) % 32))
							*(*int32)(unsafe.Add(mBase, uint32(v10+int32(40)))) = v292 - int32(base.Ui32(v296*int32(7834))>>(uint(int32(8))%32))
							v306 = base.I32_rem_u_s(v296+int32(10), int32(12))
							*(*int32)(unsafe.Add(mBase, uint32(v10+int32(44)))) = v306 + int32(1)
							*(*int32)(unsafe.Add(mBase, uint32(v10)+68)) = int32(0)
							*(*int64)(unsafe.Add(mBase, uint32(v10)+60)) = int64(4294967295)
							if v229 < int64(0) {
								v318 = v229 + int64(86400000000)
							} else {
								v318 = v229
							}
							v320 = base.I64_div_s(v318, int64(3600000000))
							*(*uint32)(unsafe.Add(mBase, uint32(v10)+36)) = uint32(v320)
							v325 = base.I64_extend32_s(v320)*int64(-3600000000) + v318
							v327 = base.I64_div_s(v325, int64(60000000))
							*(*uint32)(unsafe.Add(mBase, uint32(v10)+32)) = uint32(v327)
							v332 = base.I64_extend32_s(v327)*int64(-60000000) + v325
							v334 = base.I64_div_s(v332, int64(1000000))
							*(*uint32)(unsafe.Add(mBase, uint32(v10)+28)) = uint32(v334)
							v338 = *(*int32)(unsafe.Add(mBase, uint32(v10)+72))
							v340 = m.G0
							v341 = int32(16)
							v342 = v340 - v341
							m.G0 = v342
							v346 = F_DetermineTimeZoneOffsetInternal(m, v10+int32(28), v338, v342+int32(8))
							mBase = m.M
							m.G0 = v342 + v341
							v350 = *(*int32)(unsafe.Add(mBase, uint32(v10)+48))
							if v350 <= int32(-4713) {
								if v350 != int32(-4713) {
									F_errstart_cold(m, int32(21), int32(0))
									mBase = m.M
									v486 = m.ExcPending
									if v486 != 0 {
										return int32(0)
									} else {
										F_errcode(m, int32(134217858))
										mBase = m.M
										v489 = m.ExcPending
										if v489 != 0 {
											return int32(0)
										} else {
											F_errmsg(m, int32(422894), int32(0))
											mBase = m.M
											v493 = m.ExcPending
											if v493 != 0 {
												return int32(0)
											} else {
												F_errfinish(m, int32(521665), int32(6376), int32(392006))
												mBase = m.M
												v498 = m.ExcPending
												if v498 != 0 {
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
									v355 = *(*int32)(unsafe.Add(mBase, uint32(v10)+44))
									if int32(10) < v355 {
										v366 = v355
										v368 = v10 + int32(8)
										v369 = *(*int32)(unsafe.Add(mBase, uint32(v10)+40))
										v374 = base.B2i32(int32(2) < v366)
										if int32(2) < v366 {
											v375 = int32(4800)
										} else {
											v375 = int32(4799)
										}
										v376 = v375 + v350
										v381 = base.I32_div_s(v376, int32(4))
										v384 = base.I32_div_s(v376, int32(-100))
										v387 = base.I32_div_s(v376, int32(400))
										if int32(2) < v366 {
											v391 = int32(1)
										} else {
											v391 = int32(13)
										}
										v396 = base.I32_div_s((v391+v366)*int32(7834), int32(256))
										v402 = base.I64_extend_i32_s(v369 + v376*int32(365) + v381 + v384 + v387 + v396 - int32(32167) - int32(2451545))
										v411 = int64(32)
										v412 = int64(20)
										v414 = int64(base.Ui64(v402) >> (uint(v411) % 64))
										v417 = int64(4294967295)
										v418 = int64(500654080)
										v420 = v402 & v417
										v421 = v418 * v420
										v425 = int64(base.Ui64(v421)>>(uint(v411)%64)) + v418*v414
										v432 = v420*v412 + v425&v417
										*(*int64)(unsafe.Add(mBase, uint32(v368)+8)) = v402*int64(0) + v402>>(uint(int64(63))%64)*int64(86400000000) + v412*v414 + int64(base.Ui64(v425)>>(uint(v411)%64)) + int64(base.Ui64(v432)>>(uint(v411)%64))
										*(*int64)(unsafe.Add(mBase, uint32(v368))) = v421&v417 | v432<<(uint(v411)%64)
										v443 = *(*int64)(unsafe.Add(mBase, uint32(v10)+16))
										v444 = *(*int64)(unsafe.Add(mBase, uint32(v10)+8))
										if v443 != v444>>(uint(int64(63))%64) {
											F_errstart_cold(m, int32(21), int32(0))
											mBase = m.M
											v486 = m.ExcPending
											if v486 != 0 {
												return int32(0)
											} else {
												F_errcode(m, int32(134217858))
												mBase = m.M
												v489 = m.ExcPending
												if v489 != 0 {
													return int32(0)
												} else {
													F_errmsg(m, int32(422894), int32(0))
													mBase = m.M
													v493 = m.ExcPending
													if v493 != 0 {
														return int32(0)
													} else {
														F_errfinish(m, int32(521665), int32(6376), int32(392006))
														mBase = m.M
														v498 = m.ExcPending
														if v498 != 0 {
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
											v452 = *(*int32)(unsafe.Add(mBase, uint32(v10)+28))
											v453 = *(*int32)(unsafe.Add(mBase, uint32(v10)+32))
											v454 = *(*int32)(unsafe.Add(mBase, uint32(v10)+36))
											v455 = int32(60)
											v464 = base.I64_extend32_s(v334*int64(4293967296)+v332) + base.I64_extend_i32_s(v452+(v453+v454*v455)*v455)*int64(1000000)
											v467 = v444 + v464
											if base.B2i32(v464 < int64(0))^base.B2i32(v467 < v444) != 0 {
												F_errstart_cold(m, int32(21), int32(0))
												mBase = m.M
												v486 = m.ExcPending
												if v486 != 0 {
													return int32(0)
												} else {
													F_errcode(m, int32(134217858))
													mBase = m.M
													v489 = m.ExcPending
													if v489 != 0 {
														return int32(0)
													} else {
														F_errmsg(m, int32(422894), int32(0))
														mBase = m.M
														v493 = m.ExcPending
														if v493 != 0 {
															return int32(0)
														} else {
															F_errfinish(m, int32(521665), int32(6376), int32(392006))
															mBase = m.M
															v498 = m.ExcPending
															if v498 != 0 {
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
												v475 = v467 + base.I64_extend_i32_s(int32(0)-v346)*int64(-1000000)
												if base.Ui64(v475+int64(211813488000000000)) < base.Ui64(int64(-9011559254509551616)) {
													v506 = v475
													if base.Ui64(int64(-9011559254509551616)) <= base.Ui64(v506+int64(211813488000000000)) {
														F_errstart_cold(m, int32(21), int32(0))
														mBase = m.M
														v561 = m.ExcPending
														if v561 != 0 {
															return int32(0)
														} else {
															F_errcode(m, int32(134217858))
															mBase = m.M
															v564 = m.ExcPending
															if v564 != 0 {
																return int32(0)
															} else {
																F_errmsg(m, int32(422894), int32(0))
																mBase = m.M
																v568 = m.ExcPending
																if v568 != 0 {
																	return int32(0)
																} else {
																	F_errfinish(m, int32(521665), int32(6382), int32(392006))
																	mBase = m.M
																	v573 = m.ExcPending
																	if v573 != 0 {
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
														v513 = F_Int64GetDatum(m, v506)
														mBase = m.M
														v514 = m.ExcPending
														if v514 != 0 {
															return int32(0)
														} else {
															v521 = v513
															m.G0 = v10 + int32(336)
															return v521
														}
													}
												} else {
													F_errstart_cold(m, int32(21), int32(0))
													mBase = m.M
													v486 = m.ExcPending
													if v486 != 0 {
														return int32(0)
													} else {
														F_errcode(m, int32(134217858))
														mBase = m.M
														v489 = m.ExcPending
														if v489 != 0 {
															return int32(0)
														} else {
															F_errmsg(m, int32(422894), int32(0))
															mBase = m.M
															v493 = m.ExcPending
															if v493 != 0 {
																return int32(0)
															} else {
																F_errfinish(m, int32(521665), int32(6376), int32(392006))
																mBase = m.M
																v498 = m.ExcPending
																if v498 != 0 {
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
										v486 = m.ExcPending
										if v486 != 0 {
											return int32(0)
										} else {
											F_errcode(m, int32(134217858))
											mBase = m.M
											v489 = m.ExcPending
											if v489 != 0 {
												return int32(0)
											} else {
												F_errmsg(m, int32(422894), int32(0))
												mBase = m.M
												v493 = m.ExcPending
												if v493 != 0 {
													return int32(0)
												} else {
													F_errfinish(m, int32(521665), int32(6376), int32(392006))
													mBase = m.M
													v498 = m.ExcPending
													if v498 != 0 {
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
								if v350 <= int32(5874897) {
									v360 = *(*int32)(unsafe.Add(mBase, uint32(v10)+44))
									v366 = v360
									v368 = v10 + int32(8)
									v369 = *(*int32)(unsafe.Add(mBase, uint32(v10)+40))
									v374 = base.B2i32(int32(2) < v366)
									if int32(2) < v366 {
										v375 = int32(4800)
									} else {
										v375 = int32(4799)
									}
									v376 = v375 + v350
									v381 = base.I32_div_s(v376, int32(4))
									v384 = base.I32_div_s(v376, int32(-100))
									v387 = base.I32_div_s(v376, int32(400))
									if int32(2) < v366 {
										v391 = int32(1)
									} else {
										v391 = int32(13)
									}
									v396 = base.I32_div_s((v391+v366)*int32(7834), int32(256))
									v402 = base.I64_extend_i32_s(v369 + v376*int32(365) + v381 + v384 + v387 + v396 - int32(32167) - int32(2451545))
									v411 = int64(32)
									v412 = int64(20)
									v414 = int64(base.Ui64(v402) >> (uint(v411) % 64))
									v417 = int64(4294967295)
									v418 = int64(500654080)
									v420 = v402 & v417
									v421 = v418 * v420
									v425 = int64(base.Ui64(v421)>>(uint(v411)%64)) + v418*v414
									v432 = v420*v412 + v425&v417
									*(*int64)(unsafe.Add(mBase, uint32(v368)+8)) = v402*int64(0) + v402>>(uint(int64(63))%64)*int64(86400000000) + v412*v414 + int64(base.Ui64(v425)>>(uint(v411)%64)) + int64(base.Ui64(v432)>>(uint(v411)%64))
									*(*int64)(unsafe.Add(mBase, uint32(v368))) = v421&v417 | v432<<(uint(v411)%64)
									v443 = *(*int64)(unsafe.Add(mBase, uint32(v10)+16))
									v444 = *(*int64)(unsafe.Add(mBase, uint32(v10)+8))
									if v443 != v444>>(uint(int64(63))%64) {
										F_errstart_cold(m, int32(21), int32(0))
										mBase = m.M
										v486 = m.ExcPending
										if v486 != 0 {
											return int32(0)
										} else {
											F_errcode(m, int32(134217858))
											mBase = m.M
											v489 = m.ExcPending
											if v489 != 0 {
												return int32(0)
											} else {
												F_errmsg(m, int32(422894), int32(0))
												mBase = m.M
												v493 = m.ExcPending
												if v493 != 0 {
													return int32(0)
												} else {
													F_errfinish(m, int32(521665), int32(6376), int32(392006))
													mBase = m.M
													v498 = m.ExcPending
													if v498 != 0 {
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
										v452 = *(*int32)(unsafe.Add(mBase, uint32(v10)+28))
										v453 = *(*int32)(unsafe.Add(mBase, uint32(v10)+32))
										v454 = *(*int32)(unsafe.Add(mBase, uint32(v10)+36))
										v455 = int32(60)
										v464 = base.I64_extend32_s(v334*int64(4293967296)+v332) + base.I64_extend_i32_s(v452+(v453+v454*v455)*v455)*int64(1000000)
										v467 = v444 + v464
										if base.B2i32(v464 < int64(0))^base.B2i32(v467 < v444) != 0 {
											F_errstart_cold(m, int32(21), int32(0))
											mBase = m.M
											v486 = m.ExcPending
											if v486 != 0 {
												return int32(0)
											} else {
												F_errcode(m, int32(134217858))
												mBase = m.M
												v489 = m.ExcPending
												if v489 != 0 {
													return int32(0)
												} else {
													F_errmsg(m, int32(422894), int32(0))
													mBase = m.M
													v493 = m.ExcPending
													if v493 != 0 {
														return int32(0)
													} else {
														F_errfinish(m, int32(521665), int32(6376), int32(392006))
														mBase = m.M
														v498 = m.ExcPending
														if v498 != 0 {
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
											v475 = v467 + base.I64_extend_i32_s(int32(0)-v346)*int64(-1000000)
											if base.Ui64(v475+int64(211813488000000000)) < base.Ui64(int64(-9011559254509551616)) {
												v506 = v475
												if base.Ui64(int64(-9011559254509551616)) <= base.Ui64(v506+int64(211813488000000000)) {
													F_errstart_cold(m, int32(21), int32(0))
													mBase = m.M
													v561 = m.ExcPending
													if v561 != 0 {
														return int32(0)
													} else {
														F_errcode(m, int32(134217858))
														mBase = m.M
														v564 = m.ExcPending
														if v564 != 0 {
															return int32(0)
														} else {
															F_errmsg(m, int32(422894), int32(0))
															mBase = m.M
															v568 = m.ExcPending
															if v568 != 0 {
																return int32(0)
															} else {
																F_errfinish(m, int32(521665), int32(6382), int32(392006))
																mBase = m.M
																v573 = m.ExcPending
																if v573 != 0 {
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
													v513 = F_Int64GetDatum(m, v506)
													mBase = m.M
													v514 = m.ExcPending
													if v514 != 0 {
														return int32(0)
													} else {
														v521 = v513
														m.G0 = v10 + int32(336)
														return v521
													}
												}
											} else {
												F_errstart_cold(m, int32(21), int32(0))
												mBase = m.M
												v486 = m.ExcPending
												if v486 != 0 {
													return int32(0)
												} else {
													F_errcode(m, int32(134217858))
													mBase = m.M
													v489 = m.ExcPending
													if v489 != 0 {
														return int32(0)
													} else {
														F_errmsg(m, int32(422894), int32(0))
														mBase = m.M
														v493 = m.ExcPending
														if v493 != 0 {
															return int32(0)
														} else {
															F_errfinish(m, int32(521665), int32(6376), int32(392006))
															mBase = m.M
															v498 = m.ExcPending
															if v498 != 0 {
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
									if v350 != int32(5874898) {
										F_errstart_cold(m, int32(21), int32(0))
										mBase = m.M
										v486 = m.ExcPending
										if v486 != 0 {
											return int32(0)
										} else {
											F_errcode(m, int32(134217858))
											mBase = m.M
											v489 = m.ExcPending
											if v489 != 0 {
												return int32(0)
											} else {
												F_errmsg(m, int32(422894), int32(0))
												mBase = m.M
												v493 = m.ExcPending
												if v493 != 0 {
													return int32(0)
												} else {
													F_errfinish(m, int32(521665), int32(6376), int32(392006))
													mBase = m.M
													v498 = m.ExcPending
													if v498 != 0 {
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
										v363 = *(*int32)(unsafe.Add(mBase, uint32(v10)+44))
										if int32(5) < v363 {
											F_errstart_cold(m, int32(21), int32(0))
											mBase = m.M
											v486 = m.ExcPending
											if v486 != 0 {
												return int32(0)
											} else {
												F_errcode(m, int32(134217858))
												mBase = m.M
												v489 = m.ExcPending
												if v489 != 0 {
													return int32(0)
												} else {
													F_errmsg(m, int32(422894), int32(0))
													mBase = m.M
													v493 = m.ExcPending
													if v493 != 0 {
														return int32(0)
													} else {
														F_errfinish(m, int32(521665), int32(6376), int32(392006))
														mBase = m.M
														v498 = m.ExcPending
														if v498 != 0 {
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
											v366 = v363
											v368 = v10 + int32(8)
											v369 = *(*int32)(unsafe.Add(mBase, uint32(v10)+40))
											v374 = base.B2i32(int32(2) < v366)
											if int32(2) < v366 {
												v375 = int32(4800)
											} else {
												v375 = int32(4799)
											}
											v376 = v375 + v350
											v381 = base.I32_div_s(v376, int32(4))
											v384 = base.I32_div_s(v376, int32(-100))
											v387 = base.I32_div_s(v376, int32(400))
											if int32(2) < v366 {
												v391 = int32(1)
											} else {
												v391 = int32(13)
											}
											v396 = base.I32_div_s((v391+v366)*int32(7834), int32(256))
											v402 = base.I64_extend_i32_s(v369 + v376*int32(365) + v381 + v384 + v387 + v396 - int32(32167) - int32(2451545))
											v411 = int64(32)
											v412 = int64(20)
											v414 = int64(base.Ui64(v402) >> (uint(v411) % 64))
											v417 = int64(4294967295)
											v418 = int64(500654080)
											v420 = v402 & v417
											v421 = v418 * v420
											v425 = int64(base.Ui64(v421)>>(uint(v411)%64)) + v418*v414
											v432 = v420*v412 + v425&v417
											*(*int64)(unsafe.Add(mBase, uint32(v368)+8)) = v402*int64(0) + v402>>(uint(int64(63))%64)*int64(86400000000) + v412*v414 + int64(base.Ui64(v425)>>(uint(v411)%64)) + int64(base.Ui64(v432)>>(uint(v411)%64))
											*(*int64)(unsafe.Add(mBase, uint32(v368))) = v421&v417 | v432<<(uint(v411)%64)
											v443 = *(*int64)(unsafe.Add(mBase, uint32(v10)+16))
											v444 = *(*int64)(unsafe.Add(mBase, uint32(v10)+8))
											if v443 != v444>>(uint(int64(63))%64) {
												F_errstart_cold(m, int32(21), int32(0))
												mBase = m.M
												v486 = m.ExcPending
												if v486 != 0 {
													return int32(0)
												} else {
													F_errcode(m, int32(134217858))
													mBase = m.M
													v489 = m.ExcPending
													if v489 != 0 {
														return int32(0)
													} else {
														F_errmsg(m, int32(422894), int32(0))
														mBase = m.M
														v493 = m.ExcPending
														if v493 != 0 {
															return int32(0)
														} else {
															F_errfinish(m, int32(521665), int32(6376), int32(392006))
															mBase = m.M
															v498 = m.ExcPending
															if v498 != 0 {
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
												v452 = *(*int32)(unsafe.Add(mBase, uint32(v10)+28))
												v453 = *(*int32)(unsafe.Add(mBase, uint32(v10)+32))
												v454 = *(*int32)(unsafe.Add(mBase, uint32(v10)+36))
												v455 = int32(60)
												v464 = base.I64_extend32_s(v334*int64(4293967296)+v332) + base.I64_extend_i32_s(v452+(v453+v454*v455)*v455)*int64(1000000)
												v467 = v444 + v464
												if base.B2i32(v464 < int64(0))^base.B2i32(v467 < v444) != 0 {
													F_errstart_cold(m, int32(21), int32(0))
													mBase = m.M
													v486 = m.ExcPending
													if v486 != 0 {
														return int32(0)
													} else {
														F_errcode(m, int32(134217858))
														mBase = m.M
														v489 = m.ExcPending
														if v489 != 0 {
															return int32(0)
														} else {
															F_errmsg(m, int32(422894), int32(0))
															mBase = m.M
															v493 = m.ExcPending
															if v493 != 0 {
																return int32(0)
															} else {
																F_errfinish(m, int32(521665), int32(6376), int32(392006))
																mBase = m.M
																v498 = m.ExcPending
																if v498 != 0 {
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
													v475 = v467 + base.I64_extend_i32_s(int32(0)-v346)*int64(-1000000)
													if base.Ui64(v475+int64(211813488000000000)) < base.Ui64(int64(-9011559254509551616)) {
														v506 = v475
														if base.Ui64(int64(-9011559254509551616)) <= base.Ui64(v506+int64(211813488000000000)) {
															F_errstart_cold(m, int32(21), int32(0))
															mBase = m.M
															v561 = m.ExcPending
															if v561 != 0 {
																return int32(0)
															} else {
																F_errcode(m, int32(134217858))
																mBase = m.M
																v564 = m.ExcPending
																if v564 != 0 {
																	return int32(0)
																} else {
																	F_errmsg(m, int32(422894), int32(0))
																	mBase = m.M
																	v568 = m.ExcPending
																	if v568 != 0 {
																		return int32(0)
																	} else {
																		F_errfinish(m, int32(521665), int32(6382), int32(392006))
																		mBase = m.M
																		v573 = m.ExcPending
																		if v573 != 0 {
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
															v513 = F_Int64GetDatum(m, v506)
															mBase = m.M
															v514 = m.ExcPending
															if v514 != 0 {
																return int32(0)
															} else {
																v521 = v513
																m.G0 = v10 + int32(336)
																return v521
															}
														}
													} else {
														F_errstart_cold(m, int32(21), int32(0))
														mBase = m.M
														v486 = m.ExcPending
														if v486 != 0 {
															return int32(0)
														} else {
															F_errcode(m, int32(134217858))
															mBase = m.M
															v489 = m.ExcPending
															if v489 != 0 {
																return int32(0)
															} else {
																F_errmsg(m, int32(422894), int32(0))
																mBase = m.M
																v493 = m.ExcPending
																if v493 != 0 {
																	return int32(0)
																} else {
																	F_errfinish(m, int32(521665), int32(6376), int32(392006))
																	mBase = m.M
																	v498 = m.ExcPending
																	if v498 != 0 {
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
