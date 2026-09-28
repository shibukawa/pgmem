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
func F_timestamp_gt_timestamptz(m *base.Module, l0 int32) int64 {
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
	v9 = *(*int64)(unsafe.Add(mBase, uint32(l0)+40))
	v10 = *(*int64)(unsafe.Add(mBase, uint32(l0)+24))
	v12 = *(*int32)(unsafe.Add(mBase, _c_F_timestamp_gt_timestamptz[0]))
	*(*int32)(unsafe.Add(mBase, uint32(v7)+8)) = v12
	v15 = *(*int64)(unsafe.Add(mBase, _c_F_timestamp_gt_timestamptz[1]))
	*(*int64)(unsafe.Add(mBase, uint32(v7))) = v15
	v17 = F_timestamp2timestamptz_safe(m, v10, v7)
	mBase = m.M
	v20 = m.ExcPending
	if v20 != 0 {
		return int64(0)
	} else {
		v21 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v7)+4)))
		if v21 != int32(1) {
			v33 = base.B2i32(v9 < v17)
		} else {
			if v17 != int64(-9223372036854775807-1) {
				if v17 != int64(9223372036854775807) {
					v33 = base.B2i32(v9 < v17)
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
func F_timestamp_le_date(m *base.Module, l0 int32) int64 {
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
		return base.I64_extend_i32_u(base.B2i32(int32(0) <= v4))
	}
}
func F_timestamp_ne_timestamptz(m *base.Module, l0 int32) int64 {
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
	v5 = m.G0
	v7 = v5 - int32(16)
	m.G0 = v7
	v9 = *(*int64)(unsafe.Add(mBase, uint32(l0)+24))
	v10 = *(*int64)(unsafe.Add(mBase, uint32(l0)+40))
	v12 = *(*int32)(unsafe.Add(mBase, _c_F_timestamp_ne_timestamptz[0]))
	*(*int32)(unsafe.Add(mBase, uint32(v7)+8)) = v12
	v15 = *(*int64)(unsafe.Add(mBase, _c_F_timestamp_ne_timestamptz[1]))
	*(*int64)(unsafe.Add(mBase, uint32(v7))) = v15
	v17 = F_timestamp2timestamptz_safe(m, v9, v7)
	mBase = m.M
	v20 = m.ExcPending
	if v20 != 0 {
		return int64(0)
	} else {
		v21 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v7)+4)))
		m.G0 = v7 + int32(16)
		return base.I64_extend_i32_u(v21&base.B2i32(base.Ui64(v17-int64(9223372036854775807)) < base.Ui64(int64(2))) | base.B2i32(v17 != v10))
	}
}
func F_timestamp_pl_interval(m *base.Module, l0 int32) int64 {
	mBase := m.M
	_ = mBase
	var v11 int32
	_ = v11
	var v13 int32
	_ = v13
	var v15 int64
	_ = v15
	var v16 int32
	_ = v16
	var v17 int32
	_ = v17
	var v22 int32
	_ = v22
	var v25 int64
	_ = v25
	var v26 int64
	_ = v26
	var v36 int32
	_ = v36
	var v39 int32
	_ = v39
	var v43 int32
	_ = v43
	var v48 int32
	_ = v48
	var v49 int32
	_ = v49
	var v52 int64
	_ = v52
	var v53 int64
	_ = v53
	var v61 int32
	_ = v61
	var v64 int32
	_ = v64
	var v68 int32
	_ = v68
	var v73 int32
	_ = v73
	var v80 int64
	_ = v80
	var v88 int64
	_ = v88
	var v89 int64
	_ = v89
	var v92 int64
	_ = v92
	var v95 int32
	_ = v95
	var v107 int32
	_ = v107
	var v108 int32
	_ = v108
	var v109 int32
	_ = v109
	var v110 int32
	_ = v110
	var v116 int32
	_ = v116
	var v121 int32
	_ = v121
	var v124 int32
	_ = v124
	var v125 int32
	_ = v125
	var v126 int32
	_ = v126
	var v129 int32
	_ = v129
	var v131 int32
	_ = v131
	var v137 int32
	_ = v137
	var v141 int32
	_ = v141
	var v142 int32
	_ = v142
	var v144 int32
	_ = v144
	var v152 int32
	_ = v152
	var v156 int32
	_ = v156
	var v166 int32
	_ = v166
	var v176 int64
	_ = v176
	var v178 int64
	_ = v178
	var v183 int64
	_ = v183
	var v185 int64
	_ = v185
	var v190 int64
	_ = v190
	var v192 int64
	_ = v192
	var v194 int32
	_ = v194
	var v196 int32
	_ = v196
	var v197 int32
	_ = v197
	var v198 int32
	_ = v198
	var v206 int32
	_ = v206
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
	var v218 int32
	_ = v218
	var v223 int32
	_ = v223
	var v224 int32
	_ = v224
	var v227 int32
	_ = v227
	var v234 int32
	_ = v234
	var v237 int32
	_ = v237
	var v239 int32
	_ = v239
	var v240 int32
	_ = v240
	var v242 int32
	_ = v242
	var v248 int32
	_ = v248
	var v252 int32
	_ = v252
	var v254 int32
	_ = v254
	var v262 int32
	_ = v262
	var v268 int32
	_ = v268
	var v275 int32
	_ = v275
	var v279 int32
	_ = v279
	var v287 int32
	_ = v287
	var v289 int32
	_ = v289
	var v305 int32
	_ = v305
	var v310 int32
	_ = v310
	var v311 int32
	_ = v311
	var v312 int32
	_ = v312
	var v317 int32
	_ = v317
	var v320 int32
	_ = v320
	var v323 int32
	_ = v323
	var v327 int32
	_ = v327
	var v332 int32
	_ = v332
	var v338 int64
	_ = v338
	var v347 int64
	_ = v347
	var v348 int64
	_ = v348
	var v350 int64
	_ = v350
	var v353 int64
	_ = v353
	var v354 int64
	_ = v354
	var v356 int64
	_ = v356
	var v357 int64
	_ = v357
	var v361 int64
	_ = v361
	var v368 int64
	_ = v368
	var v379 int64
	_ = v379
	var v380 int64
	_ = v380
	var v388 int32
	_ = v388
	var v389 int32
	_ = v389
	var v390 int32
	_ = v390
	var v391 int32
	_ = v391
	var v400 int64
	_ = v400
	var v403 int64
	_ = v403
	var v415 int64
	_ = v415
	var v418 int32
	_ = v418
	var v420 int64
	_ = v420
	var v428 int64
	_ = v428
	var v429 int64
	_ = v429
	var v432 int64
	_ = v432
	var v435 int32
	_ = v435
	var v439 int32
	_ = v439
	var v441 int32
	_ = v441
	var v443 int32
	_ = v443
	var v447 int32
	_ = v447
	var v448 int32
	_ = v448
	var v449 int32
	_ = v449
	var v450 int32
	_ = v450
	var v456 int32
	_ = v456
	var v461 int32
	_ = v461
	var v464 int32
	_ = v464
	var v465 int32
	_ = v465
	var v466 int32
	_ = v466
	var v469 int32
	_ = v469
	var v471 int32
	_ = v471
	var v477 int32
	_ = v477
	var v481 int32
	_ = v481
	var v482 int32
	_ = v482
	var v484 int32
	_ = v484
	var v492 int32
	_ = v492
	var v496 int32
	_ = v496
	var v506 int32
	_ = v506
	var v516 int64
	_ = v516
	var v518 int64
	_ = v518
	var v523 int64
	_ = v523
	var v525 int64
	_ = v525
	var v530 int64
	_ = v530
	var v532 int64
	_ = v532
	var v536 int32
	_ = v536
	var v537 int32
	_ = v537
	var v538 int32
	_ = v538
	var v543 int32
	_ = v543
	var v544 int32
	_ = v544
	var v545 int32
	_ = v545
	var v550 int32
	_ = v550
	var v553 int32
	_ = v553
	var v556 int32
	_ = v556
	var v560 int32
	_ = v560
	var v565 int32
	_ = v565
	var v568 int32
	_ = v568
	var v569 int32
	_ = v569
	var v570 int32
	_ = v570
	var v572 int32
	_ = v572
	var v581 int32
	_ = v581
	var v582 int32
	_ = v582
	var v583 int32
	_ = v583
	var v584 int32
	_ = v584
	var v590 int32
	_ = v590
	var v595 int32
	_ = v595
	var v598 int32
	_ = v598
	var v599 int32
	_ = v599
	var v600 int32
	_ = v600
	var v603 int32
	_ = v603
	var v605 int32
	_ = v605
	var v611 int32
	_ = v611
	var v615 int32
	_ = v615
	var v616 int32
	_ = v616
	var v618 int32
	_ = v618
	var v626 int32
	_ = v626
	var v630 int32
	_ = v630
	var v640 int32
	_ = v640
	var v644 int32
	_ = v644
	var v649 int32
	_ = v649
	var v654 int32
	_ = v654
	var v657 int32
	_ = v657
	var v660 int32
	_ = v660
	var v661 int32
	_ = v661
	var v666 int32
	_ = v666
	var v667 int32
	_ = v667
	var v668 int32
	_ = v668
	var v673 int32
	_ = v673
	var v676 int32
	_ = v676
	var v679 int32
	_ = v679
	var v683 int32
	_ = v683
	var v688 int32
	_ = v688
	var v694 int64
	_ = v694
	var v703 int64
	_ = v703
	var v704 int64
	_ = v704
	var v706 int64
	_ = v706
	var v709 int64
	_ = v709
	var v710 int64
	_ = v710
	var v712 int64
	_ = v712
	var v713 int64
	_ = v713
	var v717 int64
	_ = v717
	var v724 int64
	_ = v724
	var v735 int64
	_ = v735
	var v736 int64
	_ = v736
	var v744 int32
	_ = v744
	var v745 int32
	_ = v745
	var v746 int32
	_ = v746
	var v747 int32
	_ = v747
	var v756 int64
	_ = v756
	var v759 int64
	_ = v759
	var v772 int64
	_ = v772
	var v775 int64
	_ = v775
	var v778 int64
	_ = v778
	var v791 int64
	_ = v791
	var v800 int32
	_ = v800
	var v803 int32
	_ = v803
	var v807 int32
	_ = v807
	var v812 int32
	_ = v812
	var v816 int32
	_ = v816
	var v819 int32
	_ = v819
	var v823 int32
	_ = v823
	var v828 int32
	_ = v828
	var v832 int32
	_ = v832
	var v835 int32
	_ = v835
	var v839 int32
	_ = v839
	var v844 int32
	_ = v844
	var v848 int32
	_ = v848
	var v851 int32
	_ = v851
	var v855 int32
	_ = v855
	var v860 int32
	_ = v860
	var v864 int32
	_ = v864
	var v867 int32
	_ = v867
	var v871 int32
	_ = v871
	var v876 int32
	_ = v876
	var v880 int32
	_ = v880
	var v883 int32
	_ = v883
	var v887 int32
	_ = v887
	var v892 int32
	_ = v892
	var v899 int32
	_ = v899
	var v902 int32
	_ = v902
	var v906 int32
	_ = v906
	var v911 int32
	_ = v911
	var v917 int32
	_ = v917
	var v920 int32
	_ = v920
	var v924 int32
	_ = v924
	var v929 int32
	_ = v929
	v11 = m.G0
	v13 = v11 - int32(80)
	m.G0 = v13
	v15 = *(*int64)(unsafe.Add(mBase, uint32(l0)+24))
	v16 = *(*int32)(unsafe.Add(mBase, uint32(l0)+40))
	v17 = *(*int32)(unsafe.Add(mBase, uint32(v16)+12))
	if v17 != int32(2147483647) {
		goto L11
	} else {
		goto L12
	}
L1:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v917 = m.ExcPending
	if v917 != 0 {
		goto L18
	} else {
		goto L163
	}
L2:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v899 = m.ExcPending
	if v899 != 0 {
		goto L18
	} else {
		goto L159
	}
L3:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v880 = m.ExcPending
	if v880 != 0 {
		goto L18
	} else {
		goto L155
	}
L4:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v864 = m.ExcPending
	if v864 != 0 {
		goto L18
	} else {
		goto L151
	}
L5:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v848 = m.ExcPending
	if v848 != 0 {
		goto L18
	} else {
		goto L147
	}
L6:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v832 = m.ExcPending
	if v832 != 0 {
		goto L18
	} else {
		goto L143
	}
L7:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v816 = m.ExcPending
	if v816 != 0 {
		goto L18
	} else {
		goto L139
	}
L8:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v800 = m.ExcPending
	if v800 != 0 {
		goto L18
	} else {
		goto L135
	}
L9:
	;
	m.G0 = v13 + int32(80)
	return v791
L10:
	;
	if base.Ui64(v15-int64(9223372036854775807)) < base.Ui64(int64(2)) {
		goto L30
	} else {
		goto L31
	}
L11:
	;
	if v17 != int32(-2147483648) {
		goto L10
	} else {
		goto L14
	}
L12:
	;
	goto L13
L13:
	;
	v49 = *(*int32)(unsafe.Add(mBase, uint32(v16)+8))
	if v49 != int32(2147483647) {
		goto L10
	} else {
		goto L23
	}
L14:
	;
	v22 = *(*int32)(unsafe.Add(mBase, uint32(v16)+8))
	if v22 != int32(-2147483648) {
		goto L10
	} else {
		goto L15
	}
L15:
	;
	v25 = int64(-9223372036854775807 - 1)
	v26 = *(*int64)(unsafe.Add(mBase, uint32(v16)))
	if v26 != v25 {
		goto L10
	} else {
		goto L16
	}
L16:
	;
	if v15 != int64(9223372036854775807) {
		v791 = v25
		goto L9
	} else {
		goto L17
	}
L17:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v36 = m.ExcPending
	if v36 != 0 {
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
	v39 = m.ExcPending
	if v39 != 0 {
		goto L18
	} else {
		goto L20
	}
L20:
	;
	F_errmsg(m, int32(_a_F_timestamp_pl_interval_0), int32(0))
	mBase = m.M
	v43 = m.ExcPending
	if v43 != 0 {
		goto L18
	} else {
		goto L21
	}
L21:
	;
	F_errfinish(m, int32(_a_F_timestamp_pl_interval_1), int32(3121), int32(_a_F_timestamp_pl_interval_2))
	mBase = m.M
	v48 = m.ExcPending
	if v48 != 0 {
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
	v52 = int64(9223372036854775807)
	v53 = *(*int64)(unsafe.Add(mBase, uint32(v16)))
	if v53 != v52 {
		goto L10
	} else {
		goto L24
	}
L24:
	;
	if v15 != int64(-9223372036854775807-1) {
		v791 = v52
		goto L9
	} else {
		goto L25
	}
L25:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v61 = m.ExcPending
	if v61 != 0 {
		goto L18
	} else {
		goto L26
	}
L26:
	;
	F_errcode(m, int32(134217858))
	mBase = m.M
	v64 = m.ExcPending
	if v64 != 0 {
		goto L18
	} else {
		goto L27
	}
L27:
	;
	F_errmsg(m, int32(_a_F_timestamp_pl_interval_0), int32(0))
	mBase = m.M
	v68 = m.ExcPending
	if v68 != 0 {
		goto L18
	} else {
		goto L28
	}
L28:
	;
	F_errfinish(m, int32(_a_F_timestamp_pl_interval_1), int32(3130), int32(_a_F_timestamp_pl_interval_2))
	mBase = m.M
	v73 = m.ExcPending
	if v73 != 0 {
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
	v791 = v15
	goto L9
L31:
	;
	goto L32
L32:
	;
	if v17 != 0 {
		goto L33
	} else {
		goto L34
	}
L33:
	;
	v80 = base.I64_div_s(v15, int64(86400000000))
	if base.Ui64(int64(172799999999)) <= base.Ui64(v15+int64(86399999999)) {
		goto L36
	} else {
		goto L37
	}
L34:
	;
	v415 = v15
	goto L35
L35:
	;
	v418 = *(*int32)(unsafe.Add(mBase, uint32(v16)+8))
	if v418 != 0 {
		goto L84
	} else {
		goto L85
	}
L36:
	;
	v88 = v80 * int64(-86400000000)
	goto L38
L37:
	;
	v88 = int64(0)
	goto L38
L38:
	;
	v89 = v88 + v15
	v92 = v89>>(uint(int64(63))%64) + v80
	if v92 <= int64(-2451546) {
		goto L8
	} else {
		goto L39
	}
L39:
	;
	v95 = base.I32_wrap_i64(v92)
	v107 = v95 + int32(_a_F_timestamp_pl_interval_3)
	v108 = int32(_a_F_timestamp_pl_interval_4)
	v109 = base.I32_div_u_s(v107, v108)
	v110 = int32(3)
	v116 = int32(2)
	v121 = base.I32_div_u_s((v109*int32(1073595727)+v107)<<(uint(v116)%32)|v110, v108)
	v124 = v95 + int32(_a_F_timestamp_pl_interval_5) + v109*v110 + v121 + int32(_a_F_timestamp_pl_interval_6)
	v125 = int32(1461)
	v126 = base.I32_div_u_s(v124, v125)
	v129 = v126*int32(-1461) + v124
	v131 = v129 << (uint(v116) % 32)
	if base.Ui32(v125) <= base.Ui32(v131) {
		goto L42
	} else {
		goto L43
	}
L40:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v13)+68)) = int64(4294967295)
	if v89 < int64(0) {
		goto L45
	} else {
		goto L46
	}
L41:
	;
	v144 = base.I32_div_u_s(v131, int32(1461))
	*(*int32)(unsafe.Add(mBase, uint32(v13+int32(56)))) = v144 + v126<<(uint(int32(2))%32) - int32(_a_F_timestamp_pl_interval_7)
	v152 = v142 + int32(123)
	v156 = int32(base.Ui32(v152*int32(2141)) >> (uint(int32(16)) % 32))
	*(*int32)(unsafe.Add(mBase, uint32(v13+int32(48)))) = v152 - int32(base.Ui32(v156*int32(_a_F_timestamp_pl_interval_8))>>(uint(int32(8))%32))
	v166 = base.I32_rem_u_s(v156+int32(10), int32(12))
	*(*int32)(unsafe.Add(mBase, uint32(v13+int32(52)))) = v166 + int32(1)
	goto L40
L42:
	;
	v137 = base.I32_rem_u_s(v129+int32(305), int32(365))
	v142 = v137
	goto L41
L43:
	;
	goto L44
L44:
	;
	v141 = base.I32_rem_u_s(v129+int32(306), int32(366))
	v142 = v141
	goto L41
L45:
	;
	v176 = v89 + int64(86400000000)
	goto L47
L46:
	;
	v176 = v89
	goto L47
L47:
	;
	v178 = base.I64_div_s(v176, int64(3600000000))
	*(*uint32)(unsafe.Add(mBase, uint32(v13)+44)) = uint32(v178)
	v183 = base.I64_extend32_s(v178)*int64(-3600000000) + v176
	v185 = base.I64_div_s(v183, int64(60000000))
	*(*uint32)(unsafe.Add(mBase, uint32(v13)+40)) = uint32(v185)
	v190 = base.I64_extend32_s(v185)*int64(-60000000) + v183
	v192 = base.I64_div_s(v190, int64(1000000))
	*(*uint32)(unsafe.Add(mBase, uint32(v13)+36)) = uint32(v192)
	v194 = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v13)+76)) = v194
	v196 = *(*int32)(unsafe.Add(mBase, uint32(v13)+52))
	v197 = *(*int32)(unsafe.Add(mBase, uint32(v16)+12))
	v198 = v196 + v197
	*(*int32)(unsafe.Add(mBase, uint32(v13)+52)) = v198
	if base.B2i32(v197 < v194) != base.B2i32(v198 < v196) {
		goto L7
	} else {
		goto L48
	}
L48:
	;
	if int32(13) <= v198 {
		goto L51
	} else {
		goto L52
	}
L49:
	;
	v242 = *(*int32)(unsafe.Add(mBase, uint32(v13)+48))
	if v239&int32(3) == int32(0) {
		goto L57
	} else {
		goto L58
	}
L50:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v13)+52)) = v237
	v239 = v234
	v240 = v237
	goto L49
L51:
	;
	v206 = int32(1)
	v207 = v198 - v206
	v208 = int32(12)
	v209 = base.I32_div_u_s(v207, v208)
	v210 = *(*int32)(unsafe.Add(mBase, uint32(v13)+56))
	v211 = v209 + v210
	*(*int32)(unsafe.Add(mBase, uint32(v13)+56)) = v211
	v234 = v211
	v237 = v207 - v209*v208 + v206
	goto L50
L52:
	;
	goto L53
L53:
	;
	v218 = *(*int32)(unsafe.Add(mBase, uint32(v13)+56))
	if int32(0) < v198 {
		v239 = v218
		v240 = v198
		goto L49
	} else {
		goto L54
	}
L54:
	;
	v223 = int32(12)
	v224 = base.I32_div_u_s(int32(0)-v198, v223)
	v227 = v218 + (v224 ^ int32(-1))
	*(*int32)(unsafe.Add(mBase, uint32(v13)+56)) = v227
	v234 = v227
	v237 = v224*v223 + v198 + v223
	goto L50
L55:
	;
	if v239 <= int32(-4713) {
		goto L67
	} else {
		goto L68
	}
L56:
	;
	v287 = *(*int32)(unsafe.Add(mBase, uint32(v279*int32(52)+v240<<(uint(int32(2))%32))+uint32(_c_F_timestamp_pl_interval[0])))
	*(*int32)(unsafe.Add(mBase, uint32(v13)+48)) = v287
	v289 = v287
	goto L55
L57:
	;
	v248 = base.I32_rem_s(v239, int32(100))
	if v248 == int32(0) {
		goto L60
	} else {
		goto L61
	}
L58:
	;
	goto L59
L59:
	;
	v275 = *(*int32)(unsafe.Add(mBase, uint32(v240<<(uint(int32(2))%32))+uint32(_c_F_timestamp_pl_interval[0])))
	if v242 <= v275 {
		v289 = v242
		goto L55
	} else {
		goto L65
	}
L60:
	;
	v252 = base.I32_rem_s(v239, int32(400))
	v254 = base.B2i32(v252 == int32(0))
	v262 = *(*int32)(unsafe.Add(mBase, uint32(v254*int32(52)+v240<<(uint(int32(2))%32))+uint32(_c_F_timestamp_pl_interval[0])))
	if v242 <= v262 {
		v289 = v242
		goto L55
	} else {
		goto L63
	}
L61:
	;
	goto L62
L62:
	;
	v268 = *(*int32)(unsafe.Add(mBase, uint32(v240<<(uint(int32(2))%32))+uint32(_c_F_timestamp_pl_interval[1])))
	if v242 <= v268 {
		v289 = v242
		goto L55
	} else {
		goto L64
	}
L63:
	;
	v279 = v254
	goto L56
L64:
	;
	v279 = int32(1)
	goto L56
L65:
	;
	v279 = int32(0)
	goto L56
L66:
	;
	v305 = v13 + int32(16)
	v310 = base.B2i32(int32(2) < v240)
	if int32(2) < v240 {
		goto L75
	} else {
		goto L76
	}
L67:
	;
	if v239 != int32(-4713) {
		goto L1
	} else {
		goto L70
	}
L68:
	;
	goto L69
L69:
	;
	if v239 < int32(_a_F_timestamp_pl_interval_13) {
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
	if base.B2i32(v239 != int32(_a_F_timestamp_pl_interval_13))|base.B2i32(base.Ui32(int32(5)) < base.Ui32(v240)) != 0 {
		goto L1
	} else {
		goto L73
	}
L73:
	;
	goto L66
L74:
	;
	v338 = base.I64_extend_i32_s(v289 + v312*int32(365) + v317 + v320 + v323 + v332 - int32(_a_F_timestamp_pl_interval_10) - int32(_a_F_timestamp_pl_interval_5))
	v347 = int64(32)
	v348 = int64(20)
	v350 = int64(base.Ui64(v338) >> (uint(v347) % 64))
	v353 = int64(4294967295)
	v354 = int64(500654080)
	v356 = v338 & v353
	v357 = v354 * v356
	v361 = int64(base.Ui64(v357)>>(uint(v347)%64)) + v354*v350
	v368 = v356*v348 + v361&v353
	*(*int64)(unsafe.Add(mBase, uint32(v305)+8)) = v338*int64(0) + v338>>(uint(int64(63))%64)*int64(86400000000) + v348*v350 + int64(base.Ui64(v361)>>(uint(v347)%64)) + int64(base.Ui64(v368)>>(uint(v347)%64))
	*(*int64)(unsafe.Add(mBase, uint32(v305))) = v357&v353 | v368<<(uint(v347)%64)
	goto L81
L75:
	;
	v311 = int32(_a_F_timestamp_pl_interval_7)
	goto L77
L76:
	;
	v311 = int32(_a_F_timestamp_pl_interval_9)
	goto L77
L77:
	;
	v312 = v311 + v239
	v317 = base.I32_div_s(v312, int32(4))
	v320 = base.I32_div_s(v312, int32(-100))
	v323 = base.I32_div_s(v312, int32(400))
	if int32(2) < v240 {
		goto L78
	} else {
		goto L79
	}
L78:
	;
	v327 = int32(1)
	goto L80
L79:
	;
	v327 = int32(13)
	goto L80
L80:
	;
	v332 = base.I32_div_s((v327+v240)*int32(_a_F_timestamp_pl_interval_8), int32(256))
	goto L74
L81:
	;
	v379 = *(*int64)(unsafe.Add(mBase, uint32(v13)+24))
	v380 = *(*int64)(unsafe.Add(mBase, uint32(v13)+16))
	if v379 != v380>>(uint(int64(63))%64) {
		goto L1
	} else {
		goto L82
	}
L82:
	;
	v388 = *(*int32)(unsafe.Add(mBase, uint32(v13)+36))
	v389 = *(*int32)(unsafe.Add(mBase, uint32(v13)+40))
	v390 = *(*int32)(unsafe.Add(mBase, uint32(v13)+44))
	v391 = int32(60)
	v400 = base.I64_extend32_s(v192*int64(4293967296)+v190) + base.I64_extend_i32_s(v388+(v389+v390*v391)*v391)*int64(1000000)
	v403 = v380 + v400
	if base.B2i32(v400 < int64(0))^base.B2i32(v403 < v380)|base.B2i32(base.Ui64(v403-int64(9223371331200000000)) <= base.Ui64(int64(9011559254509551615))) != 0 {
		goto L1
	} else {
		goto L83
	}
L83:
	;
	v415 = v403
	goto L35
L84:
	;
	v420 = base.I64_div_s(v415, int64(86400000000))
	if base.Ui64(int64(172799999999)) <= base.Ui64(v415+int64(86399999999)) {
		goto L87
	} else {
		goto L88
	}
L85:
	;
	v772 = v415
	goto L86
L86:
	;
	v775 = *(*int64)(unsafe.Add(mBase, uint32(v16)))
	v778 = v772 + v775
	if base.B2i32(v775 < int64(0)) != base.B2i32(v778 < v772) {
		goto L4
	} else {
		goto L133
	}
L87:
	;
	v428 = v420 * int64(-86400000000)
	goto L89
L88:
	;
	v428 = int64(0)
	goto L89
L89:
	;
	v429 = v428 + v415
	v432 = v429>>(uint(int64(63))%64) + v420
	if v432 <= int64(-2451546) {
		goto L6
	} else {
		goto L90
	}
L90:
	;
	v435 = base.I32_wrap_i64(v432)
	v439 = v13 + int32(56)
	v441 = v13 + int32(52)
	v443 = v13 + int32(48)
	v447 = v435 + int32(_a_F_timestamp_pl_interval_3)
	v448 = int32(_a_F_timestamp_pl_interval_4)
	v449 = base.I32_div_u_s(v447, v448)
	v450 = int32(3)
	v456 = int32(2)
	v461 = base.I32_div_u_s((v449*int32(1073595727)+v447)<<(uint(v456)%32)|v450, v448)
	v464 = v435 + int32(_a_F_timestamp_pl_interval_5) + v449*v450 + v461 + int32(_a_F_timestamp_pl_interval_6)
	v465 = int32(1461)
	v466 = base.I32_div_u_s(v464, v465)
	v469 = v466*int32(-1461) + v464
	v471 = v469 << (uint(v456) % 32)
	if base.Ui32(v465) <= base.Ui32(v471) {
		goto L93
	} else {
		goto L94
	}
L91:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v13)+68)) = int64(4294967295)
	if v429 < int64(0) {
		goto L96
	} else {
		goto L97
	}
L92:
	;
	v484 = base.I32_div_u_s(v471, int32(1461))
	*(*int32)(unsafe.Add(mBase, uint32(v439))) = v484 + v466<<(uint(int32(2))%32) - int32(_a_F_timestamp_pl_interval_7)
	v492 = v482 + int32(123)
	v496 = int32(base.Ui32(v492*int32(2141)) >> (uint(int32(16)) % 32))
	*(*int32)(unsafe.Add(mBase, uint32(v443))) = v492 - int32(base.Ui32(v496*int32(_a_F_timestamp_pl_interval_8))>>(uint(int32(8))%32))
	v506 = base.I32_rem_u_s(v496+int32(10), int32(12))
	*(*int32)(unsafe.Add(mBase, uint32(v441))) = v506 + int32(1)
	goto L91
L93:
	;
	v477 = base.I32_rem_u_s(v469+int32(305), int32(365))
	v482 = v477
	goto L92
L94:
	;
	goto L95
L95:
	;
	v481 = base.I32_rem_u_s(v469+int32(306), int32(366))
	v482 = v481
	goto L92
L96:
	;
	v516 = v429 + int64(86400000000)
	goto L98
L97:
	;
	v516 = v429
	goto L98
L98:
	;
	v518 = base.I64_div_s(v516, int64(3600000000))
	*(*uint32)(unsafe.Add(mBase, uint32(v13)+44)) = uint32(v518)
	v523 = base.I64_extend32_s(v518)*int64(-3600000000) + v516
	v525 = base.I64_div_s(v523, int64(60000000))
	*(*uint32)(unsafe.Add(mBase, uint32(v13)+40)) = uint32(v525)
	v530 = base.I64_extend32_s(v525)*int64(-60000000) + v523
	v532 = base.I64_div_s(v530, int64(1000000))
	*(*uint32)(unsafe.Add(mBase, uint32(v13)+36)) = uint32(v532)
	*(*int32)(unsafe.Add(mBase, uint32(v13)+76)) = int32(0)
	v536 = *(*int32)(unsafe.Add(mBase, uint32(v13)+56))
	v537 = *(*int32)(unsafe.Add(mBase, uint32(v13)+52))
	v538 = *(*int32)(unsafe.Add(mBase, uint32(v13)+48))
	v543 = base.B2i32(int32(2) < v537)
	if int32(2) < v537 {
		goto L100
	} else {
		goto L101
	}
L99:
	;
	v569 = *(*int32)(unsafe.Add(mBase, uint32(v16)+8))
	v570 = int32(0)
	v572 = v568 + v569
	if base.B2i32(v569 < v570)^base.B2i32(v572 < v568)|base.B2i32(v572 < v570) != 0 {
		goto L5
	} else {
		goto L106
	}
L100:
	;
	v544 = int32(_a_F_timestamp_pl_interval_7)
	goto L102
L101:
	;
	v544 = int32(_a_F_timestamp_pl_interval_9)
	goto L102
L102:
	;
	v545 = v544 + v536
	v550 = base.I32_div_s(v545, int32(4))
	v553 = base.I32_div_s(v545, int32(-100))
	v556 = base.I32_div_s(v545, int32(400))
	if int32(2) < v537 {
		goto L103
	} else {
		goto L104
	}
L103:
	;
	v560 = int32(1)
	goto L105
L104:
	;
	v560 = int32(13)
	goto L105
L105:
	;
	v565 = base.I32_div_s((v560+v537)*int32(_a_F_timestamp_pl_interval_8), int32(256))
	v568 = v538 + v545*int32(365) + v550 + v553 + v556 + v565 - int32(_a_F_timestamp_pl_interval_10)
	goto L99
L106:
	;
	v581 = v572 + int32(_a_F_timestamp_pl_interval_11)
	v582 = int32(_a_F_timestamp_pl_interval_4)
	v583 = base.I32_div_u_s(v581, v582)
	v584 = int32(3)
	v590 = int32(2)
	v595 = base.I32_div_u_s((v583*int32(1073595727)+v581)<<(uint(v590)%32)|v584, v582)
	v598 = v572 + v583*v584 + v595 + int32(_a_F_timestamp_pl_interval_6)
	v599 = int32(1461)
	v600 = base.I32_div_u_s(v598, v599)
	v603 = v600*int32(-1461) + v598
	v605 = v603 << (uint(v590) % 32)
	if base.Ui32(v599) <= base.Ui32(v605) {
		goto L109
	} else {
		goto L110
	}
L107:
	;
	v644 = *(*int32)(unsafe.Add(mBase, uint32(v13)+56))
	if v644 <= int32(-4713) {
		goto L113
	} else {
		goto L114
	}
L108:
	;
	v618 = base.I32_div_u_s(v605, int32(1461))
	*(*int32)(unsafe.Add(mBase, uint32(v439))) = v618 + v600<<(uint(int32(2))%32) - int32(_a_F_timestamp_pl_interval_7)
	v626 = v616 + int32(123)
	v630 = int32(base.Ui32(v626*int32(2141)) >> (uint(int32(16)) % 32))
	*(*int32)(unsafe.Add(mBase, uint32(v443))) = v626 - int32(base.Ui32(v630*int32(_a_F_timestamp_pl_interval_8))>>(uint(int32(8))%32))
	v640 = base.I32_rem_u_s(v630+int32(10), int32(12))
	*(*int32)(unsafe.Add(mBase, uint32(v441))) = v640 + int32(1)
	goto L107
L109:
	;
	v611 = base.I32_rem_u_s(v603+int32(305), int32(365))
	v616 = v611
	goto L108
L110:
	;
	goto L111
L111:
	;
	v615 = base.I32_rem_u_s(v603+int32(306), int32(366))
	v616 = v615
	goto L108
L112:
	;
	v661 = *(*int32)(unsafe.Add(mBase, uint32(v13)+48))
	v666 = base.B2i32(int32(2) < v660)
	if int32(2) < v660 {
		goto L124
	} else {
		goto L125
	}
L113:
	;
	if v644 != int32(-4713) {
		goto L2
	} else {
		goto L116
	}
L114:
	;
	goto L115
L115:
	;
	if v644 <= int32(_a_F_timestamp_pl_interval_12) {
		goto L118
	} else {
		goto L119
	}
L116:
	;
	v649 = *(*int32)(unsafe.Add(mBase, uint32(v13)+52))
	if int32(10) < v649 {
		v660 = v649
		goto L112
	} else {
		goto L117
	}
L117:
	;
	goto L2
L118:
	;
	v654 = *(*int32)(unsafe.Add(mBase, uint32(v13)+52))
	v660 = v654
	goto L112
L119:
	;
	goto L120
L120:
	;
	if v644 != int32(_a_F_timestamp_pl_interval_13) {
		goto L2
	} else {
		goto L121
	}
L121:
	;
	v657 = *(*int32)(unsafe.Add(mBase, uint32(v13)+52))
	if int32(5) < v657 {
		goto L2
	} else {
		goto L122
	}
L122:
	;
	v660 = v657
	goto L112
L123:
	;
	v694 = base.I64_extend_i32_s(v661 + v668*int32(365) + v673 + v676 + v679 + v688 - int32(_a_F_timestamp_pl_interval_10) - int32(_a_F_timestamp_pl_interval_5))
	v703 = int64(32)
	v704 = int64(20)
	v706 = int64(base.Ui64(v694) >> (uint(v703) % 64))
	v709 = int64(4294967295)
	v710 = int64(500654080)
	v712 = v694 & v709
	v713 = v710 * v712
	v717 = int64(base.Ui64(v713)>>(uint(v703)%64)) + v710*v706
	v724 = v712*v704 + v717&v709
	*(*int64)(unsafe.Add(mBase, uint32(v13)+8)) = v694*int64(0) + v694>>(uint(int64(63))%64)*int64(86400000000) + v704*v706 + int64(base.Ui64(v717)>>(uint(v703)%64)) + int64(base.Ui64(v724)>>(uint(v703)%64))
	*(*int64)(unsafe.Add(mBase, uint32(v13))) = v713&v709 | v724<<(uint(v703)%64)
	goto L130
L124:
	;
	v667 = int32(_a_F_timestamp_pl_interval_7)
	goto L126
L125:
	;
	v667 = int32(_a_F_timestamp_pl_interval_9)
	goto L126
L126:
	;
	v668 = v667 + v644
	v673 = base.I32_div_s(v668, int32(4))
	v676 = base.I32_div_s(v668, int32(-100))
	v679 = base.I32_div_s(v668, int32(400))
	if int32(2) < v660 {
		goto L127
	} else {
		goto L128
	}
L127:
	;
	v683 = int32(1)
	goto L129
L128:
	;
	v683 = int32(13)
	goto L129
L129:
	;
	v688 = base.I32_div_s((v683+v660)*int32(_a_F_timestamp_pl_interval_8), int32(256))
	goto L123
L130:
	;
	v735 = *(*int64)(unsafe.Add(mBase, uint32(v13)+8))
	v736 = *(*int64)(unsafe.Add(mBase, uint32(v13)))
	if v735 != v736>>(uint(int64(63))%64) {
		goto L2
	} else {
		goto L131
	}
L131:
	;
	v744 = *(*int32)(unsafe.Add(mBase, uint32(v13)+36))
	v745 = *(*int32)(unsafe.Add(mBase, uint32(v13)+40))
	v746 = *(*int32)(unsafe.Add(mBase, uint32(v13)+44))
	v747 = int32(60)
	v756 = base.I64_extend32_s(v532*int64(4293967296)+v530) + base.I64_extend_i32_s(v744+(v745+v746*v747)*v747)*int64(1000000)
	v759 = v736 + v756
	if base.B2i32(v756 < int64(0))^base.B2i32(v759 < v736)|base.B2i32(base.Ui64(v759-int64(9223371331200000000)) <= base.Ui64(int64(9011559254509551615))) != 0 {
		goto L2
	} else {
		goto L132
	}
L132:
	;
	v772 = v759
	goto L86
L133:
	;
	if base.Ui64(int64(-9011559254509551616)) <= base.Ui64(v778+int64(211813488000000000)) {
		goto L3
	} else {
		goto L134
	}
L134:
	;
	v791 = v778
	goto L9
L135:
	;
	F_errcode(m, int32(134217858))
	mBase = m.M
	v803 = m.ExcPending
	if v803 != 0 {
		goto L18
	} else {
		goto L136
	}
L136:
	;
	F_errmsg(m, int32(_a_F_timestamp_pl_interval_0), int32(0))
	mBase = m.M
	v807 = m.ExcPending
	if v807 != 0 {
		goto L18
	} else {
		goto L137
	}
L137:
	;
	F_errfinish(m, int32(_a_F_timestamp_pl_interval_1), int32(3147), int32(_a_F_timestamp_pl_interval_2))
	mBase = m.M
	v812 = m.ExcPending
	if v812 != 0 {
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
	v819 = m.ExcPending
	if v819 != 0 {
		goto L18
	} else {
		goto L140
	}
L140:
	;
	F_errmsg(m, int32(_a_F_timestamp_pl_interval_0), int32(0))
	mBase = m.M
	v823 = m.ExcPending
	if v823 != 0 {
		goto L18
	} else {
		goto L141
	}
L141:
	;
	F_errfinish(m, int32(_a_F_timestamp_pl_interval_1), int32(3152), int32(_a_F_timestamp_pl_interval_2))
	mBase = m.M
	v828 = m.ExcPending
	if v828 != 0 {
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
	v835 = m.ExcPending
	if v835 != 0 {
		goto L18
	} else {
		goto L144
	}
L144:
	;
	F_errmsg(m, int32(_a_F_timestamp_pl_interval_0), int32(0))
	mBase = m.M
	v839 = m.ExcPending
	if v839 != 0 {
		goto L18
	} else {
		goto L145
	}
L145:
	;
	F_errfinish(m, int32(_a_F_timestamp_pl_interval_1), int32(3184), int32(_a_F_timestamp_pl_interval_2))
	mBase = m.M
	v844 = m.ExcPending
	if v844 != 0 {
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
	v851 = m.ExcPending
	if v851 != 0 {
		goto L18
	} else {
		goto L148
	}
L148:
	;
	F_errmsg(m, int32(_a_F_timestamp_pl_interval_0), int32(0))
	mBase = m.M
	v855 = m.ExcPending
	if v855 != 0 {
		goto L18
	} else {
		goto L149
	}
L149:
	;
	F_errfinish(m, int32(_a_F_timestamp_pl_interval_1), int32(3195), int32(_a_F_timestamp_pl_interval_2))
	mBase = m.M
	v860 = m.ExcPending
	if v860 != 0 {
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
	v867 = m.ExcPending
	if v867 != 0 {
		goto L18
	} else {
		goto L152
	}
L152:
	;
	F_errmsg(m, int32(_a_F_timestamp_pl_interval_0), int32(0))
	mBase = m.M
	v871 = m.ExcPending
	if v871 != 0 {
		goto L18
	} else {
		goto L153
	}
L153:
	;
	F_errfinish(m, int32(_a_F_timestamp_pl_interval_1), int32(3207), int32(_a_F_timestamp_pl_interval_2))
	mBase = m.M
	v876 = m.ExcPending
	if v876 != 0 {
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
	v883 = m.ExcPending
	if v883 != 0 {
		goto L18
	} else {
		goto L156
	}
L156:
	;
	F_errmsg(m, int32(_a_F_timestamp_pl_interval_0), int32(0))
	mBase = m.M
	v887 = m.ExcPending
	if v887 != 0 {
		goto L18
	} else {
		goto L157
	}
L157:
	;
	F_errfinish(m, int32(_a_F_timestamp_pl_interval_1), int32(3212), int32(_a_F_timestamp_pl_interval_2))
	mBase = m.M
	v892 = m.ExcPending
	if v892 != 0 {
		goto L18
	} else {
		goto L158
	}
L158:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L159:
	;
	F_errcode(m, int32(134217858))
	mBase = m.M
	v902 = m.ExcPending
	if v902 != 0 {
		goto L18
	} else {
		goto L160
	}
L160:
	;
	F_errmsg(m, int32(_a_F_timestamp_pl_interval_0), int32(0))
	mBase = m.M
	v906 = m.ExcPending
	if v906 != 0 {
		goto L18
	} else {
		goto L161
	}
L161:
	;
	F_errfinish(m, int32(_a_F_timestamp_pl_interval_1), int32(3201), int32(_a_F_timestamp_pl_interval_2))
	mBase = m.M
	v911 = m.ExcPending
	if v911 != 0 {
		goto L18
	} else {
		goto L162
	}
L162:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L163:
	;
	F_errcode(m, int32(134217858))
	mBase = m.M
	v920 = m.ExcPending
	if v920 != 0 {
		goto L18
	} else {
		goto L164
	}
L164:
	;
	F_errmsg(m, int32(_a_F_timestamp_pl_interval_0), int32(0))
	mBase = m.M
	v924 = m.ExcPending
	if v924 != 0 {
		goto L18
	} else {
		goto L165
	}
L165:
	;
	F_errfinish(m, int32(_a_F_timestamp_pl_interval_1), int32(3171), int32(_a_F_timestamp_pl_interval_2))
	mBase = m.M
	v929 = m.ExcPending
	if v929 != 0 {
		goto L18
	} else {
		goto L166
	}
L166:
	;
	base.Wasm_trap_unreachable()
	for {
	}
}
func F_timestamp_scale(m *base.Module, l0 int32) int64 {
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
	var v14 int32
	_ = v14
	var v15 int32
	_ = v15
	var v18 int32
	_ = v18
	var v21 int32
	_ = v21
	var v24 int64
	_ = v24
	var v25 int64
	_ = v25
	v5 = m.G0
	v7 = v5 - int32(16)
	m.G0 = v7
	v9 = *(*int32)(unsafe.Add(mBase, uint32(l0)+40))
	v10 = *(*int64)(unsafe.Add(mBase, uint32(l0)+24))
	*(*int64)(unsafe.Add(mBase, uint32(v7)+8)) = v10
	v14 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v15 = F_AdjustTimestampForTypmod(m, v7+int32(8), v9, v14)
	mBase = m.M
	v18 = m.ExcPending
	if v18 != 0 {
		return int64(0)
	} else {
		if v15 == int32(0) {
			v21 = int32(1)
			*(*uint8)(unsafe.Add(mBase, uint32(l0)+16)) = uint8(v21)
			v25 = int64(0)
		} else {
			v24 = *(*int64)(unsafe.Add(mBase, uint32(v7)+8))
			v25 = v24
		}
		m.G0 = v7 + int32(16)
		return v25
	}
}
func F_timestamp_timestamptz(m *base.Module, l0 int32) int64 {
	mBase := m.M
	_ = mBase
	var v4 int64
	_ = v4
	var v5 int32
	_ = v5
	var v6 int64
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
	var v22 int64
	_ = v22
	v4 = *(*int64)(unsafe.Add(mBase, uint32(l0)+24))
	v5 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v6 = F_timestamp2timestamptz_safe(m, v4, v5)
	mBase = m.M
	v9 = m.ExcPending
	if v9 != 0 {
		return int64(0)
	} else {
		v10 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
		if v10 == int32(0) {
			v22 = v6
		} else {
			v13 = *(*int32)(unsafe.Add(mBase, uint32(v10)))
			if v13 != int32(453) {
				v22 = v6
			} else {
				v16 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v10)+4)))
				if v16 != int32(1) {
					v22 = v6
				} else {
					v19 = int32(1)
					*(*uint8)(unsafe.Add(mBase, uint32(l0)+16)) = uint8(v19)
					v22 = int64(0)
				}
			}
		}
		return v22
	}
}
func F_timestamp_to_char(m *base.Module, l0 int32) int64 {
	mBase := m.M
	_ = mBase
	var v7 int32
	_ = v7
	var v9 int32
	_ = v9
	var v11 int64
	_ = v11
	var v12 int32
	_ = v12
	var v13 int32
	_ = v13
	var v16 int32
	_ = v16
	var v17 int32
	_ = v17
	var v23 int32
	_ = v23
	var v26 int32
	_ = v26
	var v33 int32
	_ = v33
	var v34 int32
	_ = v34
	var v40 int32
	_ = v40
	var v46 int32
	_ = v46
	var v52 int32
	_ = v52
	var v55 int32
	_ = v55
	var v58 int64
	_ = v58
	var v70 int32
	_ = v70
	var v77 int32
	_ = v77
	var v78 int32
	_ = v78
	var v79 int32
	_ = v79
	var v80 int32
	_ = v80
	var v81 int32
	_ = v81
	var v86 int32
	_ = v86
	var v87 int32
	_ = v87
	var v88 int32
	_ = v88
	var v93 int32
	_ = v93
	var v96 int32
	_ = v96
	var v99 int32
	_ = v99
	var v103 int32
	_ = v103
	var v108 int32
	_ = v108
	var v111 int32
	_ = v111
	var v112 int32
	_ = v112
	var v115 int32
	_ = v115
	var v117 int32
	_ = v117
	var v126 int32
	_ = v126
	var v131 int32
	_ = v131
	var v134 int32
	_ = v134
	var v137 int32
	_ = v137
	var v146 int32
	_ = v146
	var v150 int64
	_ = v150
	var v152 int64
	_ = v152
	var v154 int64
	_ = v154
	var v156 int64
	_ = v156
	var v158 int32
	_ = v158
	var v162 int32
	_ = v162
	var v168 int32
	_ = v168
	var v169 int32
	_ = v169
	var v170 int32
	_ = v170
	var v173 int32
	_ = v173
	var v180 int64
	_ = v180
	var v188 int32
	_ = v188
	var v191 int32
	_ = v191
	var v195 int32
	_ = v195
	var v200 int32
	_ = v200
	v7 = m.G0
	v9 = v7 - int32(96)
	m.G0 = v9
	v11 = *(*int64)(unsafe.Add(mBase, uint32(l0)+24))
	v12 = *(*int32)(unsafe.Add(mBase, uint32(l0)+40))
	v13 = F_pg_detoast_datum_packed(m, v12)
	mBase = m.M
	v16 = m.ExcPending
	if v16 != 0 {
		return int64(0)
	} else {
		v17 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v13))))
		if v17 == int32(1) {
			v23 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v13)+1)))
			if v23 == int32(18) {
				v26 = int32(16)
			} else {
				v26 = int32(0)
			}
			if base.Ui32((v23-int32(1))&int32(255)) < base.Ui32(int32(3)) {
				v33 = int32(4)
			} else {
				v33 = v26
			}
			v46 = v33
		} else {
			v34 = int32(1)
			if v17&v34 != 0 {
				v46 = int32(base.Ui32(v17)>>(uint(v34)%32)) - v34
			} else {
				v40 = *(*int32)(unsafe.Add(mBase, uint32(v13)))
				v46 = int32(base.Ui32(v40)>>(uint(int32(2))%32)) - int32(4)
			}
		}
		if base.Ui64(int64(1)) < base.Ui64(v11-int64(9223372036854775807)) {
			v52 = v46
		} else {
			v52 = int32(0)
		}
		if v52 == int32(0) {
			v55 = int32(1)
			*(*uint8)(unsafe.Add(mBase, uint32(l0)+16)) = uint8(v55)
			v180 = int64(0)
			m.G0 = v9 + int32(96)
			return v180
		} else {
			v58 = int64(0)
			*(*int64)(unsafe.Add(mBase, uint32(v9)+80)) = v58
			*(*int64)(unsafe.Add(mBase, uint32(v9)+72)) = v58
			*(*int64)(unsafe.Add(mBase, uint32(v9)+56)) = v58
			*(*int64)(unsafe.Add(mBase, uint32(v9)+48)) = v58
			*(*int64)(unsafe.Add(mBase, uint32(v9)+88)) = v58
			*(*int64)(unsafe.Add(mBase, uint32(v9)+64)) = int64(4294967297)
			v70 = int32(0)
			v77 = F_timestamp2tm(m, v11, v70, v9+int32(4), v9+int32(88), v70, v70)
			mBase = m.M
			v78 = m.ExcPending
			if v78 != 0 {
				return int64(0)
			} else {
				if v77 != 0 {
					F_errstart_cold(m, int32(21), int32(0))
					mBase = m.M
					v188 = m.ExcPending
					if v188 != 0 {
						return int64(0)
					} else {
						F_errcode(m, int32(134217858))
						mBase = m.M
						v191 = m.ExcPending
						if v191 != 0 {
							return int64(0)
						} else {
							F_errmsg(m, int32(_a_F_timestamp_to_char_0), int32(0))
							mBase = m.M
							v195 = m.ExcPending
							if v195 != 0 {
								return int64(0)
							} else {
								F_errfinish(m, int32(_a_F_timestamp_to_char_1), int32(_a_F_timestamp_to_char_2), int32(_a_F_timestamp_to_char_3))
								mBase = m.M
								v200 = m.ExcPending
								if v200 != 0 {
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
					v79 = *(*int32)(unsafe.Add(mBase, uint32(v9)+24))
					v80 = *(*int32)(unsafe.Add(mBase, uint32(v9)+20))
					v81 = *(*int32)(unsafe.Add(mBase, uint32(v9)+16))
					v86 = base.B2i32(int32(2) < v80)
					if int32(2) < v80 {
						v87 = int32(_a_F_timestamp_to_char_4)
					} else {
						v87 = int32(_a_F_timestamp_to_char_5)
					}
					v88 = v87 + v79
					v93 = base.I32_div_s(v88, int32(4))
					v96 = base.I32_div_s(v88, int32(-100))
					v99 = base.I32_div_s(v88, int32(400))
					if int32(2) < v80 {
						v103 = int32(1)
					} else {
						v103 = int32(13)
					}
					v108 = base.I32_div_s((v103+v80)*int32(_a_F_timestamp_to_char_6), int32(256))
					v111 = v81 + v88*int32(365) + v93 + v96 + v99 + v108 - int32(_a_F_timestamp_to_char_7)
					v112 = int32(1)
					v115 = base.I32_rem_s(v111+v112, int32(7))
					*(*int32)(unsafe.Add(mBase, uint32(v9)+28)) = v115
					v117 = *(*int32)(unsafe.Add(mBase, uint32(v9)+24))
					v126 = int32(_a_F_timestamp_to_char_5) + v117
					v131 = base.I32_div_s(v126, int32(4))
					v134 = base.I32_div_s(v126, int32(-100))
					v137 = base.I32_div_s(v126, int32(400))
					v146 = base.I32_div_s(int32(_a_F_timestamp_to_char_8), int32(256))
					v150 = *(*int64)(unsafe.Add(mBase, uint32(v9)+4))
					*(*int64)(unsafe.Add(mBase, uint32(v9)+48)) = v150
					v152 = int64(*(*int32)(unsafe.Add(mBase, uint32(v9)+12)))
					*(*int64)(unsafe.Add(mBase, uint32(v9)+56)) = v152
					v154 = *(*int64)(unsafe.Add(mBase, uint32(v9)+16))
					*(*int64)(unsafe.Add(mBase, uint32(v9)+64)) = v154
					v156 = *(*int64)(unsafe.Add(mBase, uint32(v9)+24))
					*(*int64)(unsafe.Add(mBase, uint32(v9)+72)) = v156
					v158 = *(*int32)(unsafe.Add(mBase, uint32(v9)+40))
					*(*int32)(unsafe.Add(mBase, uint32(v9)+84)) = v158
					v162 = v111 - (v112 + v126*int32(365) + v131 + v134 + v137 + v146 - int32(_a_F_timestamp_to_char_7)) + int32(1)
					*(*int32)(unsafe.Add(mBase, uint32(v9)+32)) = v162
					*(*int32)(unsafe.Add(mBase, uint32(v9)+80)) = v162
					v168 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
					v169 = F_datetime_to_char_body(m, v9+int32(48), v13, int32(0), v168)
					mBase = m.M
					v170 = m.ExcPending
					if v170 != 0 {
						return int64(0)
					} else {
						if v169 == int32(0) {
							v173 = int32(1)
							*(*uint8)(unsafe.Add(mBase, uint32(l0)+16)) = uint8(v173)
							v180 = int64(0)
						} else {
							v180 = base.I64_extend_i32_u(v169)
						}
						m.G0 = v9 + int32(96)
						return v180
					}
				}
			}
		}
	}
}
func F_timestamp_zone(m *base.Module, l0 int32) int64 {
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
	var v17 int64
	_ = v17
	var v23 int32
	_ = v23
	var v26 int32
	_ = v26
	var v31 int32
	_ = v31
	var v32 int32
	_ = v32
	var v34 int64
	_ = v34
	var v42 int64
	_ = v42
	var v43 int64
	_ = v43
	var v46 int64
	_ = v46
	var v49 int32
	_ = v49
	var v61 int32
	_ = v61
	var v62 int32
	_ = v62
	var v63 int32
	_ = v63
	var v64 int32
	_ = v64
	var v70 int32
	_ = v70
	var v75 int32
	_ = v75
	var v78 int32
	_ = v78
	var v79 int32
	_ = v79
	var v80 int32
	_ = v80
	var v83 int32
	_ = v83
	var v85 int32
	_ = v85
	var v91 int32
	_ = v91
	var v95 int32
	_ = v95
	var v96 int32
	_ = v96
	var v98 int32
	_ = v98
	var v106 int32
	_ = v106
	var v110 int32
	_ = v110
	var v120 int32
	_ = v120
	var v130 int64
	_ = v130
	var v132 int64
	_ = v132
	var v137 int64
	_ = v137
	var v139 int64
	_ = v139
	var v146 int64
	_ = v146
	var v148 int32
	_ = v148
	var v152 int32
	_ = v152
	var v155 int32
	_ = v155
	var v160 int32
	_ = v160
	var v162 int32
	_ = v162
	var v166 int32
	_ = v166
	var v168 int32
	_ = v168
	var v170 int32
	_ = v170
	var v171 int32
	_ = v171
	var v173 int32
	_ = v173
	var v177 int32
	_ = v177
	var v179 int32
	_ = v179
	var v181 int32
	_ = v181
	var v199 int32
	_ = v199
	var v200 int32
	_ = v200
	var v201 int32
	_ = v201
	var v206 int32
	_ = v206
	var v216 int64
	_ = v216
	var v224 int64
	_ = v224
	var v225 int64
	_ = v225
	var v228 int64
	_ = v228
	var v231 int32
	_ = v231
	var v243 int32
	_ = v243
	var v244 int32
	_ = v244
	var v245 int32
	_ = v245
	var v246 int32
	_ = v246
	var v252 int32
	_ = v252
	var v257 int32
	_ = v257
	var v260 int32
	_ = v260
	var v261 int32
	_ = v261
	var v262 int32
	_ = v262
	var v265 int32
	_ = v265
	var v267 int32
	_ = v267
	var v273 int32
	_ = v273
	var v277 int32
	_ = v277
	var v278 int32
	_ = v278
	var v280 int32
	_ = v280
	var v288 int32
	_ = v288
	var v292 int32
	_ = v292
	var v302 int32
	_ = v302
	var v314 int64
	_ = v314
	var v316 int64
	_ = v316
	var v321 int64
	_ = v321
	var v323 int64
	_ = v323
	var v328 int64
	_ = v328
	var v330 int64
	_ = v330
	var v334 int32
	_ = v334
	var v336 int32
	_ = v336
	var v337 int32
	_ = v337
	var v338 int32
	_ = v338
	var v342 int32
	_ = v342
	var v346 int32
	_ = v346
	var v351 int32
	_ = v351
	var v356 int32
	_ = v356
	var v359 int32
	_ = v359
	var v362 int32
	_ = v362
	var v363 int32
	_ = v363
	var v368 int32
	_ = v368
	var v369 int32
	_ = v369
	var v370 int32
	_ = v370
	var v375 int32
	_ = v375
	var v378 int32
	_ = v378
	var v381 int32
	_ = v381
	var v385 int32
	_ = v385
	var v390 int32
	_ = v390
	var v396 int64
	_ = v396
	var v405 int64
	_ = v405
	var v406 int64
	_ = v406
	var v408 int64
	_ = v408
	var v411 int64
	_ = v411
	var v412 int64
	_ = v412
	var v414 int64
	_ = v414
	var v415 int64
	_ = v415
	var v419 int64
	_ = v419
	var v426 int64
	_ = v426
	var v437 int64
	_ = v437
	var v438 int64
	_ = v438
	var v446 int32
	_ = v446
	var v447 int32
	_ = v447
	var v448 int32
	_ = v448
	var v449 int32
	_ = v449
	var v458 int64
	_ = v458
	var v461 int64
	_ = v461
	var v469 int64
	_ = v469
	var v480 int32
	_ = v480
	var v483 int32
	_ = v483
	var v487 int32
	_ = v487
	var v492 int32
	_ = v492
	var v493 int64
	_ = v493
	var v500 int64
	_ = v500
	var v510 int64
	_ = v510
	var v520 int32
	_ = v520
	var v523 int32
	_ = v523
	var v527 int32
	_ = v527
	var v532 int32
	_ = v532
	var v536 int32
	_ = v536
	var v539 int32
	_ = v539
	var v543 int32
	_ = v543
	var v548 int32
	_ = v548
	var v552 int32
	_ = v552
	var v555 int32
	_ = v555
	var v559 int32
	_ = v559
	var v564 int32
	_ = v564
	v8 = m.G0
	v10 = v8 - int32(336)
	m.G0 = v10
	v12 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
	v13 = F_pg_detoast_datum_packed(m, v12)
	mBase = m.M
	v16 = m.ExcPending
	if v16 != 0 {
		return int64(0)
	} else {
		v17 = *(*int64)(unsafe.Add(mBase, uint32(l0)+40))
		if base.Ui64(int64(2)) <= base.Ui64(v17-int64(9223372036854775807)) {
			v23 = v10 + int32(80)
			F_text_to_cstring_buffer(m, v13, v23, int32(256))
			mBase = m.M
			v26 = m.ExcPending
			if v26 != 0 {
				return int64(0)
			} else {
				v31 = F_DecodeTimezoneName(m, v23, v10+int32(76), v10+int32(72))
				mBase = m.M
				v32 = m.ExcPending
				if v32 != 0 {
					return int64(0)
				} else {
					switch v31 {
					case 0:
						v493 = int64(*(*int32)(unsafe.Add(mBase, uint32(v10)+76)))
						v500 = v493*int64(-1000000) + v17
						if base.Ui64(int64(-9011559254509551616)) <= base.Ui64(v500+int64(211813488000000000)) {
							F_errstart_cold(m, int32(21), int32(0))
							mBase = m.M
							v552 = m.ExcPending
							if v552 != 0 {
								return int64(0)
							} else {
								F_errcode(m, int32(134217858))
								mBase = m.M
								v555 = m.ExcPending
								if v555 != 0 {
									return int64(0)
								} else {
									F_errmsg(m, int32(_a_F_timestamp_zone_0), int32(0))
									mBase = m.M
									v559 = m.ExcPending
									if v559 != 0 {
										return int64(0)
									} else {
										F_errfinish(m, int32(_a_F_timestamp_zone_1), int32(_a_F_timestamp_zone_2), int32(_a_F_timestamp_zone_3))
										mBase = m.M
										v564 = m.ExcPending
										if v564 != 0 {
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
							v510 = v500
							m.G0 = v10 + int32(336)
							return v510
						}
					case 1:
						v34 = base.I64_div_s(v17, int64(86400000000))
						if base.Ui64(int64(172799999999)) <= base.Ui64(v17+int64(86399999999)) {
							v42 = v34 * int64(-86400000000)
						} else {
							v42 = int64(0)
						}
						v43 = v42 + v17
						v46 = v43>>(uint(int64(63))%64) + v34
						if v46 <= int64(-2451546) {
							F_errstart_cold(m, int32(21), int32(0))
							mBase = m.M
							v520 = m.ExcPending
							if v520 != 0 {
								return int64(0)
							} else {
								F_errcode(m, int32(134217858))
								mBase = m.M
								v523 = m.ExcPending
								if v523 != 0 {
									return int64(0)
								} else {
									F_errmsg(m, int32(_a_F_timestamp_zone_0), int32(0))
									mBase = m.M
									v527 = m.ExcPending
									if v527 != 0 {
										return int64(0)
									} else {
										F_errfinish(m, int32(_a_F_timestamp_zone_1), int32(_a_F_timestamp_zone_4), int32(_a_F_timestamp_zone_3))
										mBase = m.M
										v532 = m.ExcPending
										if v532 != 0 {
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
							v49 = base.I32_wrap_i64(v46)
							v61 = v49 + int32(_a_F_timestamp_zone_5)
							v62 = int32(_a_F_timestamp_zone_6)
							v63 = base.I32_div_u_s(v61, v62)
							v64 = int32(3)
							v70 = int32(2)
							v75 = base.I32_div_u_s((v63*int32(1073595727)+v61)<<(uint(v70)%32)|v64, v62)
							v78 = v49 + int32(_a_F_timestamp_zone_7) + v63*v64 + v75 + int32(_a_F_timestamp_zone_8)
							v79 = int32(1461)
							v80 = base.I32_div_u_s(v78, v79)
							v83 = v80*int32(-1461) + v78
							v85 = v83 << (uint(v70) % 32)
							if base.Ui32(v79) <= base.Ui32(v85) {
								v91 = base.I32_rem_u_s(v83+int32(305), int32(365))
								v96 = v91
							} else {
								v95 = base.I32_rem_u_s(v83+int32(306), int32(366))
								v96 = v95
							}
							v98 = base.I32_div_u_s(v85, int32(1461))
							*(*int32)(unsafe.Add(mBase, uint32(v10+int32(48)))) = v98 + v80<<(uint(int32(2))%32) - int32(_a_F_timestamp_zone_9)
							v106 = v96 + int32(123)
							v110 = int32(base.Ui32(v106*int32(2141)) >> (uint(int32(16)) % 32))
							*(*int32)(unsafe.Add(mBase, uint32(v10+int32(40)))) = v106 - int32(base.Ui32(v110*int32(_a_F_timestamp_zone_10))>>(uint(int32(8))%32))
							v120 = base.I32_rem_u_s(v110+int32(10), int32(12))
							*(*int32)(unsafe.Add(mBase, uint32(v10+int32(44)))) = v120 + int32(1)
							*(*int64)(unsafe.Add(mBase, uint32(v10)+60)) = int64(4294967295)
							if v43 < int64(0) {
								v130 = v43 + int64(86400000000)
							} else {
								v130 = v43
							}
							v132 = base.I64_div_s(v130, int64(3600000000))
							*(*uint32)(unsafe.Add(mBase, uint32(v10)+36)) = uint32(v132)
							v137 = base.I64_extend32_s(v132)*int64(-3600000000) + v130
							v139 = base.I64_div_s(v137, int64(60000000))
							*(*uint32)(unsafe.Add(mBase, uint32(v10)+32)) = uint32(v139)
							v146 = base.I64_div_s(base.I64_extend32_s(v139)*int64(-60000000)+v137, int64(1000000))
							*(*uint32)(unsafe.Add(mBase, uint32(v10)+28)) = uint32(v146)
							v148 = int32(0)
							*(*int32)(unsafe.Add(mBase, uint32(v10)+68)) = v148
							v152 = v10 + int32(28)
							v155 = *(*int32)(unsafe.Add(mBase, uint32(v10)+72))
							v160 = m.G0
							v162 = v160 - int32(288)
							m.G0 = v162
							v166 = F_DetermineTimeZoneOffsetInternal(m, v152, v155, v162+int32(280))
							mBase = m.M
							v168 = v162 + int32(16)
							v170 = F_strlcpy(m, v168, v10+int32(80), int32(256))
							mBase = m.M
							v171 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v162)+16)))
							if v171 != 0 {
								v173 = v168
								v177 = v171
								for {
									v179 = F_pg_toupper(m, v177)
									mBase = m.M
									*(*uint8)(unsafe.Add(mBase, uint32(v173))) = uint8(v179)
									v181 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v173)+1)))
									if v181 != 0 {
										v173 = v173 + int32(1)
										v177 = v181
										continue
									} else {
										break
									}
									break
								}
							} else {
							}
							v199 = F_pg_interpret_timezone_abbrev(m, v162+int32(16), v162+int32(280), v162+int32(12), v162+int32(8), v155)
							mBase = m.M
							if v199 != 0 {
								v200 = *(*int32)(unsafe.Add(mBase, uint32(v162)+12))
								v201 = *(*int32)(unsafe.Add(mBase, uint32(v162)+8))
								*(*int32)(unsafe.Add(mBase, uint32(v152)+32)) = v201
								v206 = int32(0) - v200
							} else {
								v206 = v166
							}
							m.G0 = v162 + int32(288)
							v500 = base.I64_extend_i32_s(v148-v206)*int64(-1000000) + v17
							if base.Ui64(int64(-9011559254509551616)) <= base.Ui64(v500+int64(211813488000000000)) {
								F_errstart_cold(m, int32(21), int32(0))
								mBase = m.M
								v552 = m.ExcPending
								if v552 != 0 {
									return int64(0)
								} else {
									F_errcode(m, int32(134217858))
									mBase = m.M
									v555 = m.ExcPending
									if v555 != 0 {
										return int64(0)
									} else {
										F_errmsg(m, int32(_a_F_timestamp_zone_0), int32(0))
										mBase = m.M
										v559 = m.ExcPending
										if v559 != 0 {
											return int64(0)
										} else {
											F_errfinish(m, int32(_a_F_timestamp_zone_1), int32(_a_F_timestamp_zone_2), int32(_a_F_timestamp_zone_3))
											mBase = m.M
											v564 = m.ExcPending
											if v564 != 0 {
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
								v510 = v500
								m.G0 = v10 + int32(336)
								return v510
							}
						}
					default:
						v216 = base.I64_div_s(v17, int64(86400000000))
						if base.Ui64(int64(172799999999)) <= base.Ui64(v17+int64(86399999999)) {
							v224 = v216 * int64(-86400000000)
						} else {
							v224 = int64(0)
						}
						v225 = v224 + v17
						v228 = v225>>(uint(int64(63))%64) + v216
						if v228 <= int64(-2451546) {
							F_errstart_cold(m, int32(21), int32(0))
							mBase = m.M
							v536 = m.ExcPending
							if v536 != 0 {
								return int64(0)
							} else {
								F_errcode(m, int32(134217858))
								mBase = m.M
								v539 = m.ExcPending
								if v539 != 0 {
									return int64(0)
								} else {
									F_errmsg(m, int32(_a_F_timestamp_zone_0), int32(0))
									mBase = m.M
									v543 = m.ExcPending
									if v543 != 0 {
										return int64(0)
									} else {
										F_errfinish(m, int32(_a_F_timestamp_zone_1), int32(_a_F_timestamp_zone_11), int32(_a_F_timestamp_zone_3))
										mBase = m.M
										v548 = m.ExcPending
										if v548 != 0 {
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
							v231 = base.I32_wrap_i64(v228)
							v243 = v231 + int32(_a_F_timestamp_zone_5)
							v244 = int32(_a_F_timestamp_zone_6)
							v245 = base.I32_div_u_s(v243, v244)
							v246 = int32(3)
							v252 = int32(2)
							v257 = base.I32_div_u_s((v245*int32(1073595727)+v243)<<(uint(v252)%32)|v246, v244)
							v260 = v231 + int32(_a_F_timestamp_zone_7) + v245*v246 + v257 + int32(_a_F_timestamp_zone_8)
							v261 = int32(1461)
							v262 = base.I32_div_u_s(v260, v261)
							v265 = v262*int32(-1461) + v260
							v267 = v265 << (uint(v252) % 32)
							if base.Ui32(v261) <= base.Ui32(v267) {
								v273 = base.I32_rem_u_s(v265+int32(305), int32(365))
								v278 = v273
							} else {
								v277 = base.I32_rem_u_s(v265+int32(306), int32(366))
								v278 = v277
							}
							v280 = base.I32_div_u_s(v267, int32(1461))
							*(*int32)(unsafe.Add(mBase, uint32(v10+int32(48)))) = v280 + v262<<(uint(int32(2))%32) - int32(_a_F_timestamp_zone_9)
							v288 = v278 + int32(123)
							v292 = int32(base.Ui32(v288*int32(2141)) >> (uint(int32(16)) % 32))
							*(*int32)(unsafe.Add(mBase, uint32(v10+int32(40)))) = v288 - int32(base.Ui32(v292*int32(_a_F_timestamp_zone_10))>>(uint(int32(8))%32))
							v302 = base.I32_rem_u_s(v292+int32(10), int32(12))
							*(*int32)(unsafe.Add(mBase, uint32(v10+int32(44)))) = v302 + int32(1)
							*(*int32)(unsafe.Add(mBase, uint32(v10)+68)) = int32(0)
							*(*int64)(unsafe.Add(mBase, uint32(v10)+60)) = int64(4294967295)
							if v225 < int64(0) {
								v314 = v225 + int64(86400000000)
							} else {
								v314 = v225
							}
							v316 = base.I64_div_s(v314, int64(3600000000))
							*(*uint32)(unsafe.Add(mBase, uint32(v10)+36)) = uint32(v316)
							v321 = base.I64_extend32_s(v316)*int64(-3600000000) + v314
							v323 = base.I64_div_s(v321, int64(60000000))
							*(*uint32)(unsafe.Add(mBase, uint32(v10)+32)) = uint32(v323)
							v328 = base.I64_extend32_s(v323)*int64(-60000000) + v321
							v330 = base.I64_div_s(v328, int64(1000000))
							*(*uint32)(unsafe.Add(mBase, uint32(v10)+28)) = uint32(v330)
							v334 = *(*int32)(unsafe.Add(mBase, uint32(v10)+72))
							v336 = m.G0
							v337 = int32(16)
							v338 = v336 - v337
							m.G0 = v338
							v342 = F_DetermineTimeZoneOffsetInternal(m, v10+int32(28), v334, v338+int32(8))
							mBase = m.M
							m.G0 = v338 + v337
							v346 = *(*int32)(unsafe.Add(mBase, uint32(v10)+48))
							if v346 <= int32(-4713) {
								if v346 != int32(-4713) {
									F_errstart_cold(m, int32(21), int32(0))
									mBase = m.M
									v480 = m.ExcPending
									if v480 != 0 {
										return int64(0)
									} else {
										F_errcode(m, int32(134217858))
										mBase = m.M
										v483 = m.ExcPending
										if v483 != 0 {
											return int64(0)
										} else {
											F_errmsg(m, int32(_a_F_timestamp_zone_0), int32(0))
											mBase = m.M
											v487 = m.ExcPending
											if v487 != 0 {
												return int64(0)
											} else {
												F_errfinish(m, int32(_a_F_timestamp_zone_1), int32(_a_F_timestamp_zone_12), int32(_a_F_timestamp_zone_3))
												mBase = m.M
												v492 = m.ExcPending
												if v492 != 0 {
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
									v351 = *(*int32)(unsafe.Add(mBase, uint32(v10)+44))
									if int32(10) < v351 {
										v362 = v351
										v363 = *(*int32)(unsafe.Add(mBase, uint32(v10)+40))
										v368 = base.B2i32(int32(2) < v362)
										if int32(2) < v362 {
											v369 = int32(_a_F_timestamp_zone_9)
										} else {
											v369 = int32(_a_F_timestamp_zone_13)
										}
										v370 = v369 + v346
										v375 = base.I32_div_s(v370, int32(4))
										v378 = base.I32_div_s(v370, int32(-100))
										v381 = base.I32_div_s(v370, int32(400))
										if int32(2) < v362 {
											v385 = int32(1)
										} else {
											v385 = int32(13)
										}
										v390 = base.I32_div_s((v385+v362)*int32(_a_F_timestamp_zone_10), int32(256))
										v396 = base.I64_extend_i32_s(v363 + v370*int32(365) + v375 + v378 + v381 + v390 - int32(_a_F_timestamp_zone_14) - int32(_a_F_timestamp_zone_7))
										v405 = int64(32)
										v406 = int64(20)
										v408 = int64(base.Ui64(v396) >> (uint(v405) % 64))
										v411 = int64(4294967295)
										v412 = int64(500654080)
										v414 = v396 & v411
										v415 = v412 * v414
										v419 = int64(base.Ui64(v415)>>(uint(v405)%64)) + v412*v408
										v426 = v414*v406 + v419&v411
										*(*int64)(unsafe.Add(mBase, uint32(v10)+8)) = v396*int64(0) + v396>>(uint(int64(63))%64)*int64(86400000000) + v406*v408 + int64(base.Ui64(v419)>>(uint(v405)%64)) + int64(base.Ui64(v426)>>(uint(v405)%64))
										*(*int64)(unsafe.Add(mBase, uint32(v10))) = v415&v411 | v426<<(uint(v405)%64)
										v437 = *(*int64)(unsafe.Add(mBase, uint32(v10)+8))
										v438 = *(*int64)(unsafe.Add(mBase, uint32(v10)))
										if v437 != v438>>(uint(int64(63))%64) {
											F_errstart_cold(m, int32(21), int32(0))
											mBase = m.M
											v480 = m.ExcPending
											if v480 != 0 {
												return int64(0)
											} else {
												F_errcode(m, int32(134217858))
												mBase = m.M
												v483 = m.ExcPending
												if v483 != 0 {
													return int64(0)
												} else {
													F_errmsg(m, int32(_a_F_timestamp_zone_0), int32(0))
													mBase = m.M
													v487 = m.ExcPending
													if v487 != 0 {
														return int64(0)
													} else {
														F_errfinish(m, int32(_a_F_timestamp_zone_1), int32(_a_F_timestamp_zone_12), int32(_a_F_timestamp_zone_3))
														mBase = m.M
														v492 = m.ExcPending
														if v492 != 0 {
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
											v446 = *(*int32)(unsafe.Add(mBase, uint32(v10)+28))
											v447 = *(*int32)(unsafe.Add(mBase, uint32(v10)+32))
											v448 = *(*int32)(unsafe.Add(mBase, uint32(v10)+36))
											v449 = int32(60)
											v458 = base.I64_extend32_s(v330*int64(4293967296)+v328) + base.I64_extend_i32_s(v446+(v447+v448*v449)*v449)*int64(1000000)
											v461 = v438 + v458
											if base.B2i32(v458 < int64(0))^base.B2i32(v461 < v438) != 0 {
												F_errstart_cold(m, int32(21), int32(0))
												mBase = m.M
												v480 = m.ExcPending
												if v480 != 0 {
													return int64(0)
												} else {
													F_errcode(m, int32(134217858))
													mBase = m.M
													v483 = m.ExcPending
													if v483 != 0 {
														return int64(0)
													} else {
														F_errmsg(m, int32(_a_F_timestamp_zone_0), int32(0))
														mBase = m.M
														v487 = m.ExcPending
														if v487 != 0 {
															return int64(0)
														} else {
															F_errfinish(m, int32(_a_F_timestamp_zone_1), int32(_a_F_timestamp_zone_12), int32(_a_F_timestamp_zone_3))
															mBase = m.M
															v492 = m.ExcPending
															if v492 != 0 {
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
												v469 = v461 + base.I64_extend_i32_s(int32(0)-v342)*int64(-1000000)
												if base.Ui64(v469+int64(211813488000000000)) < base.Ui64(int64(-9011559254509551616)) {
													v500 = v469
													if base.Ui64(int64(-9011559254509551616)) <= base.Ui64(v500+int64(211813488000000000)) {
														F_errstart_cold(m, int32(21), int32(0))
														mBase = m.M
														v552 = m.ExcPending
														if v552 != 0 {
															return int64(0)
														} else {
															F_errcode(m, int32(134217858))
															mBase = m.M
															v555 = m.ExcPending
															if v555 != 0 {
																return int64(0)
															} else {
																F_errmsg(m, int32(_a_F_timestamp_zone_0), int32(0))
																mBase = m.M
																v559 = m.ExcPending
																if v559 != 0 {
																	return int64(0)
																} else {
																	F_errfinish(m, int32(_a_F_timestamp_zone_1), int32(_a_F_timestamp_zone_2), int32(_a_F_timestamp_zone_3))
																	mBase = m.M
																	v564 = m.ExcPending
																	if v564 != 0 {
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
														v510 = v500
														m.G0 = v10 + int32(336)
														return v510
													}
												} else {
													F_errstart_cold(m, int32(21), int32(0))
													mBase = m.M
													v480 = m.ExcPending
													if v480 != 0 {
														return int64(0)
													} else {
														F_errcode(m, int32(134217858))
														mBase = m.M
														v483 = m.ExcPending
														if v483 != 0 {
															return int64(0)
														} else {
															F_errmsg(m, int32(_a_F_timestamp_zone_0), int32(0))
															mBase = m.M
															v487 = m.ExcPending
															if v487 != 0 {
																return int64(0)
															} else {
																F_errfinish(m, int32(_a_F_timestamp_zone_1), int32(_a_F_timestamp_zone_12), int32(_a_F_timestamp_zone_3))
																mBase = m.M
																v492 = m.ExcPending
																if v492 != 0 {
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
									} else {
										F_errstart_cold(m, int32(21), int32(0))
										mBase = m.M
										v480 = m.ExcPending
										if v480 != 0 {
											return int64(0)
										} else {
											F_errcode(m, int32(134217858))
											mBase = m.M
											v483 = m.ExcPending
											if v483 != 0 {
												return int64(0)
											} else {
												F_errmsg(m, int32(_a_F_timestamp_zone_0), int32(0))
												mBase = m.M
												v487 = m.ExcPending
												if v487 != 0 {
													return int64(0)
												} else {
													F_errfinish(m, int32(_a_F_timestamp_zone_1), int32(_a_F_timestamp_zone_12), int32(_a_F_timestamp_zone_3))
													mBase = m.M
													v492 = m.ExcPending
													if v492 != 0 {
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
								if v346 <= int32(_a_F_timestamp_zone_15) {
									v356 = *(*int32)(unsafe.Add(mBase, uint32(v10)+44))
									v362 = v356
									v363 = *(*int32)(unsafe.Add(mBase, uint32(v10)+40))
									v368 = base.B2i32(int32(2) < v362)
									if int32(2) < v362 {
										v369 = int32(_a_F_timestamp_zone_9)
									} else {
										v369 = int32(_a_F_timestamp_zone_13)
									}
									v370 = v369 + v346
									v375 = base.I32_div_s(v370, int32(4))
									v378 = base.I32_div_s(v370, int32(-100))
									v381 = base.I32_div_s(v370, int32(400))
									if int32(2) < v362 {
										v385 = int32(1)
									} else {
										v385 = int32(13)
									}
									v390 = base.I32_div_s((v385+v362)*int32(_a_F_timestamp_zone_10), int32(256))
									v396 = base.I64_extend_i32_s(v363 + v370*int32(365) + v375 + v378 + v381 + v390 - int32(_a_F_timestamp_zone_14) - int32(_a_F_timestamp_zone_7))
									v405 = int64(32)
									v406 = int64(20)
									v408 = int64(base.Ui64(v396) >> (uint(v405) % 64))
									v411 = int64(4294967295)
									v412 = int64(500654080)
									v414 = v396 & v411
									v415 = v412 * v414
									v419 = int64(base.Ui64(v415)>>(uint(v405)%64)) + v412*v408
									v426 = v414*v406 + v419&v411
									*(*int64)(unsafe.Add(mBase, uint32(v10)+8)) = v396*int64(0) + v396>>(uint(int64(63))%64)*int64(86400000000) + v406*v408 + int64(base.Ui64(v419)>>(uint(v405)%64)) + int64(base.Ui64(v426)>>(uint(v405)%64))
									*(*int64)(unsafe.Add(mBase, uint32(v10))) = v415&v411 | v426<<(uint(v405)%64)
									v437 = *(*int64)(unsafe.Add(mBase, uint32(v10)+8))
									v438 = *(*int64)(unsafe.Add(mBase, uint32(v10)))
									if v437 != v438>>(uint(int64(63))%64) {
										F_errstart_cold(m, int32(21), int32(0))
										mBase = m.M
										v480 = m.ExcPending
										if v480 != 0 {
											return int64(0)
										} else {
											F_errcode(m, int32(134217858))
											mBase = m.M
											v483 = m.ExcPending
											if v483 != 0 {
												return int64(0)
											} else {
												F_errmsg(m, int32(_a_F_timestamp_zone_0), int32(0))
												mBase = m.M
												v487 = m.ExcPending
												if v487 != 0 {
													return int64(0)
												} else {
													F_errfinish(m, int32(_a_F_timestamp_zone_1), int32(_a_F_timestamp_zone_12), int32(_a_F_timestamp_zone_3))
													mBase = m.M
													v492 = m.ExcPending
													if v492 != 0 {
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
										v446 = *(*int32)(unsafe.Add(mBase, uint32(v10)+28))
										v447 = *(*int32)(unsafe.Add(mBase, uint32(v10)+32))
										v448 = *(*int32)(unsafe.Add(mBase, uint32(v10)+36))
										v449 = int32(60)
										v458 = base.I64_extend32_s(v330*int64(4293967296)+v328) + base.I64_extend_i32_s(v446+(v447+v448*v449)*v449)*int64(1000000)
										v461 = v438 + v458
										if base.B2i32(v458 < int64(0))^base.B2i32(v461 < v438) != 0 {
											F_errstart_cold(m, int32(21), int32(0))
											mBase = m.M
											v480 = m.ExcPending
											if v480 != 0 {
												return int64(0)
											} else {
												F_errcode(m, int32(134217858))
												mBase = m.M
												v483 = m.ExcPending
												if v483 != 0 {
													return int64(0)
												} else {
													F_errmsg(m, int32(_a_F_timestamp_zone_0), int32(0))
													mBase = m.M
													v487 = m.ExcPending
													if v487 != 0 {
														return int64(0)
													} else {
														F_errfinish(m, int32(_a_F_timestamp_zone_1), int32(_a_F_timestamp_zone_12), int32(_a_F_timestamp_zone_3))
														mBase = m.M
														v492 = m.ExcPending
														if v492 != 0 {
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
											v469 = v461 + base.I64_extend_i32_s(int32(0)-v342)*int64(-1000000)
											if base.Ui64(v469+int64(211813488000000000)) < base.Ui64(int64(-9011559254509551616)) {
												v500 = v469
												if base.Ui64(int64(-9011559254509551616)) <= base.Ui64(v500+int64(211813488000000000)) {
													F_errstart_cold(m, int32(21), int32(0))
													mBase = m.M
													v552 = m.ExcPending
													if v552 != 0 {
														return int64(0)
													} else {
														F_errcode(m, int32(134217858))
														mBase = m.M
														v555 = m.ExcPending
														if v555 != 0 {
															return int64(0)
														} else {
															F_errmsg(m, int32(_a_F_timestamp_zone_0), int32(0))
															mBase = m.M
															v559 = m.ExcPending
															if v559 != 0 {
																return int64(0)
															} else {
																F_errfinish(m, int32(_a_F_timestamp_zone_1), int32(_a_F_timestamp_zone_2), int32(_a_F_timestamp_zone_3))
																mBase = m.M
																v564 = m.ExcPending
																if v564 != 0 {
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
													v510 = v500
													m.G0 = v10 + int32(336)
													return v510
												}
											} else {
												F_errstart_cold(m, int32(21), int32(0))
												mBase = m.M
												v480 = m.ExcPending
												if v480 != 0 {
													return int64(0)
												} else {
													F_errcode(m, int32(134217858))
													mBase = m.M
													v483 = m.ExcPending
													if v483 != 0 {
														return int64(0)
													} else {
														F_errmsg(m, int32(_a_F_timestamp_zone_0), int32(0))
														mBase = m.M
														v487 = m.ExcPending
														if v487 != 0 {
															return int64(0)
														} else {
															F_errfinish(m, int32(_a_F_timestamp_zone_1), int32(_a_F_timestamp_zone_12), int32(_a_F_timestamp_zone_3))
															mBase = m.M
															v492 = m.ExcPending
															if v492 != 0 {
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
								} else {
									if v346 != int32(_a_F_timestamp_zone_16) {
										F_errstart_cold(m, int32(21), int32(0))
										mBase = m.M
										v480 = m.ExcPending
										if v480 != 0 {
											return int64(0)
										} else {
											F_errcode(m, int32(134217858))
											mBase = m.M
											v483 = m.ExcPending
											if v483 != 0 {
												return int64(0)
											} else {
												F_errmsg(m, int32(_a_F_timestamp_zone_0), int32(0))
												mBase = m.M
												v487 = m.ExcPending
												if v487 != 0 {
													return int64(0)
												} else {
													F_errfinish(m, int32(_a_F_timestamp_zone_1), int32(_a_F_timestamp_zone_12), int32(_a_F_timestamp_zone_3))
													mBase = m.M
													v492 = m.ExcPending
													if v492 != 0 {
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
										v359 = *(*int32)(unsafe.Add(mBase, uint32(v10)+44))
										if int32(5) < v359 {
											F_errstart_cold(m, int32(21), int32(0))
											mBase = m.M
											v480 = m.ExcPending
											if v480 != 0 {
												return int64(0)
											} else {
												F_errcode(m, int32(134217858))
												mBase = m.M
												v483 = m.ExcPending
												if v483 != 0 {
													return int64(0)
												} else {
													F_errmsg(m, int32(_a_F_timestamp_zone_0), int32(0))
													mBase = m.M
													v487 = m.ExcPending
													if v487 != 0 {
														return int64(0)
													} else {
														F_errfinish(m, int32(_a_F_timestamp_zone_1), int32(_a_F_timestamp_zone_12), int32(_a_F_timestamp_zone_3))
														mBase = m.M
														v492 = m.ExcPending
														if v492 != 0 {
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
											v362 = v359
											v363 = *(*int32)(unsafe.Add(mBase, uint32(v10)+40))
											v368 = base.B2i32(int32(2) < v362)
											if int32(2) < v362 {
												v369 = int32(_a_F_timestamp_zone_9)
											} else {
												v369 = int32(_a_F_timestamp_zone_13)
											}
											v370 = v369 + v346
											v375 = base.I32_div_s(v370, int32(4))
											v378 = base.I32_div_s(v370, int32(-100))
											v381 = base.I32_div_s(v370, int32(400))
											if int32(2) < v362 {
												v385 = int32(1)
											} else {
												v385 = int32(13)
											}
											v390 = base.I32_div_s((v385+v362)*int32(_a_F_timestamp_zone_10), int32(256))
											v396 = base.I64_extend_i32_s(v363 + v370*int32(365) + v375 + v378 + v381 + v390 - int32(_a_F_timestamp_zone_14) - int32(_a_F_timestamp_zone_7))
											v405 = int64(32)
											v406 = int64(20)
											v408 = int64(base.Ui64(v396) >> (uint(v405) % 64))
											v411 = int64(4294967295)
											v412 = int64(500654080)
											v414 = v396 & v411
											v415 = v412 * v414
											v419 = int64(base.Ui64(v415)>>(uint(v405)%64)) + v412*v408
											v426 = v414*v406 + v419&v411
											*(*int64)(unsafe.Add(mBase, uint32(v10)+8)) = v396*int64(0) + v396>>(uint(int64(63))%64)*int64(86400000000) + v406*v408 + int64(base.Ui64(v419)>>(uint(v405)%64)) + int64(base.Ui64(v426)>>(uint(v405)%64))
											*(*int64)(unsafe.Add(mBase, uint32(v10))) = v415&v411 | v426<<(uint(v405)%64)
											v437 = *(*int64)(unsafe.Add(mBase, uint32(v10)+8))
											v438 = *(*int64)(unsafe.Add(mBase, uint32(v10)))
											if v437 != v438>>(uint(int64(63))%64) {
												F_errstart_cold(m, int32(21), int32(0))
												mBase = m.M
												v480 = m.ExcPending
												if v480 != 0 {
													return int64(0)
												} else {
													F_errcode(m, int32(134217858))
													mBase = m.M
													v483 = m.ExcPending
													if v483 != 0 {
														return int64(0)
													} else {
														F_errmsg(m, int32(_a_F_timestamp_zone_0), int32(0))
														mBase = m.M
														v487 = m.ExcPending
														if v487 != 0 {
															return int64(0)
														} else {
															F_errfinish(m, int32(_a_F_timestamp_zone_1), int32(_a_F_timestamp_zone_12), int32(_a_F_timestamp_zone_3))
															mBase = m.M
															v492 = m.ExcPending
															if v492 != 0 {
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
												v446 = *(*int32)(unsafe.Add(mBase, uint32(v10)+28))
												v447 = *(*int32)(unsafe.Add(mBase, uint32(v10)+32))
												v448 = *(*int32)(unsafe.Add(mBase, uint32(v10)+36))
												v449 = int32(60)
												v458 = base.I64_extend32_s(v330*int64(4293967296)+v328) + base.I64_extend_i32_s(v446+(v447+v448*v449)*v449)*int64(1000000)
												v461 = v438 + v458
												if base.B2i32(v458 < int64(0))^base.B2i32(v461 < v438) != 0 {
													F_errstart_cold(m, int32(21), int32(0))
													mBase = m.M
													v480 = m.ExcPending
													if v480 != 0 {
														return int64(0)
													} else {
														F_errcode(m, int32(134217858))
														mBase = m.M
														v483 = m.ExcPending
														if v483 != 0 {
															return int64(0)
														} else {
															F_errmsg(m, int32(_a_F_timestamp_zone_0), int32(0))
															mBase = m.M
															v487 = m.ExcPending
															if v487 != 0 {
																return int64(0)
															} else {
																F_errfinish(m, int32(_a_F_timestamp_zone_1), int32(_a_F_timestamp_zone_12), int32(_a_F_timestamp_zone_3))
																mBase = m.M
																v492 = m.ExcPending
																if v492 != 0 {
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
													v469 = v461 + base.I64_extend_i32_s(int32(0)-v342)*int64(-1000000)
													if base.Ui64(v469+int64(211813488000000000)) < base.Ui64(int64(-9011559254509551616)) {
														v500 = v469
														if base.Ui64(int64(-9011559254509551616)) <= base.Ui64(v500+int64(211813488000000000)) {
															F_errstart_cold(m, int32(21), int32(0))
															mBase = m.M
															v552 = m.ExcPending
															if v552 != 0 {
																return int64(0)
															} else {
																F_errcode(m, int32(134217858))
																mBase = m.M
																v555 = m.ExcPending
																if v555 != 0 {
																	return int64(0)
																} else {
																	F_errmsg(m, int32(_a_F_timestamp_zone_0), int32(0))
																	mBase = m.M
																	v559 = m.ExcPending
																	if v559 != 0 {
																		return int64(0)
																	} else {
																		F_errfinish(m, int32(_a_F_timestamp_zone_1), int32(_a_F_timestamp_zone_2), int32(_a_F_timestamp_zone_3))
																		mBase = m.M
																		v564 = m.ExcPending
																		if v564 != 0 {
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
															v510 = v500
															m.G0 = v10 + int32(336)
															return v510
														}
													} else {
														F_errstart_cold(m, int32(21), int32(0))
														mBase = m.M
														v480 = m.ExcPending
														if v480 != 0 {
															return int64(0)
														} else {
															F_errcode(m, int32(134217858))
															mBase = m.M
															v483 = m.ExcPending
															if v483 != 0 {
																return int64(0)
															} else {
																F_errmsg(m, int32(_a_F_timestamp_zone_0), int32(0))
																mBase = m.M
																v487 = m.ExcPending
																if v487 != 0 {
																	return int64(0)
																} else {
																	F_errfinish(m, int32(_a_F_timestamp_zone_1), int32(_a_F_timestamp_zone_12), int32(_a_F_timestamp_zone_3))
																	mBase = m.M
																	v492 = m.ExcPending
																	if v492 != 0 {
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
					}
				}
			}
		} else {
			v510 = v17
			m.G0 = v10 + int32(336)
			return v510
		}
	}
}
