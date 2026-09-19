package p1

import (
	base "github.com/shibukawa/pgmem/internal/aot/pgaot/base"
	"unsafe"
)

func F_logicalrep_launcher_onexit(m *base.Module, l0 int32, l1 int32) {
	mBase := m.M
	_ = mBase
	var v4 int32
	_ = v4
	v4 = *(*int32)(unsafe.Add(mBase, _c_F_logicalrep_launcher_onexit[0]))
	*(*int32)(unsafe.Add(mBase, uint32(v4))) = int32(0)
	return
}
func F_logicalrep_message_type(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v4 int32
	_ = v4
	var v6 int32
	_ = v6
	var v30 int32
	_ = v30
	var v34 int32
	_ = v34
	var v37 int32
	_ = v37
	var v38 int32
	_ = v38
	v4 = m.G0
	v6 = v4 - int32(16)
	m.G0 = v6
	switch l0 - int32(65) {
	case 0:
		v38 = int32(_a_F_logicalrep_message_type_0)
		m.G0 = v6 + int32(16)
		return v38
	case 1:
		v38 = int32(_a_F_logicalrep_message_type_1)
		m.G0 = v6 + int32(16)
		return v38
	case 2:
		v38 = int32(_a_F_logicalrep_message_type_2)
		m.G0 = v6 + int32(16)
		return v38
	case 3:
		v38 = int32(_a_F_logicalrep_message_type_3)
		m.G0 = v6 + int32(16)
		return v38
	case 4:
		v38 = int32(_a_F_logicalrep_message_type_4)
		m.G0 = v6 + int32(16)
		return v38
	default:
		*(*int32)(unsafe.Add(mBase, uint32(v6))) = l0
		v30 = int32(_a_F_logicalrep_message_type_5)
		v34 = F_pg_snprintf(m, v30, int32(20), int32(_a_F_logicalrep_message_type_6), v6)
		mBase = m.M
		v37 = m.ExcPending
		if v37 != 0 {
			return int32(0)
		} else {
			v38 = v30
			m.G0 = v6 + int32(16)
			return v38
		}
	case 8:
		v38 = int32(_a_F_logicalrep_message_type_7)
		m.G0 = v6 + int32(16)
		return v38
	case 10:
		v38 = int32(_a_F_logicalrep_message_type_8)
		m.G0 = v6 + int32(16)
		return v38
	case 12:
		v38 = int32(_a_F_logicalrep_message_type_9)
		m.G0 = v6 + int32(16)
		return v38
	case 14:
		v38 = int32(_a_F_logicalrep_message_type_10)
		m.G0 = v6 + int32(16)
		return v38
	case 15:
		v38 = int32(_a_F_logicalrep_message_type_11)
		m.G0 = v6 + int32(16)
		return v38
	case 17:
		v38 = int32(_a_F_logicalrep_message_type_12)
		m.G0 = v6 + int32(16)
		return v38
	case 18:
		v38 = int32(_a_F_logicalrep_message_type_13)
		m.G0 = v6 + int32(16)
		return v38
	case 19:
		v38 = int32(_a_F_logicalrep_message_type_14)
		m.G0 = v6 + int32(16)
		return v38
	case 20:
		v38 = int32(_a_F_logicalrep_message_type_15)
		m.G0 = v6 + int32(16)
		return v38
	case 24:
		v38 = int32(_a_F_logicalrep_message_type_16)
		m.G0 = v6 + int32(16)
		return v38
	case 33:
		v38 = int32(_a_F_logicalrep_message_type_17)
		m.G0 = v6 + int32(16)
		return v38
	case 34:
		v38 = int32(_a_F_logicalrep_message_type_18)
		m.G0 = v6 + int32(16)
		return v38
	case 47:
		v38 = int32(_a_F_logicalrep_message_type_19)
		m.G0 = v6 + int32(16)
		return v38
	case 49:
		v38 = int32(_a_F_logicalrep_message_type_20)
		m.G0 = v6 + int32(16)
		return v38
	}
}
func F_logicalrep_rel_close(m *base.Module, l0 int32, l1 int32) {
	mBase := m.M
	_ = mBase
	var v3 int32
	_ = v3
	var v5 int32
	_ = v5
	v3 = *(*int32)(unsafe.Add(mBase, uint32(l0)+40))
	F_relation_close(m, v3, l1)
	mBase = m.M
	v5 = m.ExcPending
	if v5 != 0 {
		return
	} else {
		*(*int32)(unsafe.Add(mBase, uint32(l0)+40)) = int32(0)
		return
	}
}
func F_logicalrep_worker_find(m *base.Module, l0 int32, l1 int32) int32 {
	mBase := m.M
	_ = mBase
	var v3 int32
	_ = v3
	var v8 int32
	_ = v8
	var v12 int32
	_ = v12
	var v18 int32
	_ = v18
	var v23 int32
	_ = v23
	var v24 int32
	_ = v24
	var v27 int32
	_ = v27
	var v30 int32
	_ = v30
	var v32 int32
	_ = v32
	var v35 int32
	_ = v35
	var v40 int32
	_ = v40
	v3 = int32(0)
	v8 = *(*int32)(unsafe.Add(mBase, _c_F_logicalrep_worker_find[0]))
	if v8 <= v3 {
		v40 = v3
		goto L1
	} else {
		goto L2
	}
L1:
	;
	return v40
L2:
	;
	v12 = *(*int32)(unsafe.Add(mBase, _c_F_logicalrep_worker_find[1]))
	v18 = v3
	goto L3
L3:
	;
	v23 = v12 + int32(16) + v18*int32(112)
	v24 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v23)+16)))
	if v24 != int32(1) {
		goto L5
	} else {
		goto L6
	}
L4:
	;
	v40 = int32(0)
	goto L1
L5:
	;
	v35 = v18 + int32(1)
	if v35 != v8 {
		v18 = v35
		goto L3
	} else {
		goto L10
	}
L6:
	;
	v27 = *(*int32)(unsafe.Add(mBase, uint32(v23)))
	if v27 == int32(3) {
		goto L5
	} else {
		goto L7
	}
L7:
	;
	v30 = *(*int32)(unsafe.Add(mBase, uint32(v23)+32))
	if v30 != l0 {
		goto L5
	} else {
		goto L8
	}
L8:
	;
	v32 = *(*int32)(unsafe.Add(mBase, uint32(v23)+36))
	if v32 != l1 {
		goto L5
	} else {
		goto L9
	}
L9:
	;
	v40 = v23
	goto L1
L10:
	;
	goto L4
}
func F_logicalrep_worker_launch(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32, l5 int32, l6 int32) int32 {
	mBase := m.M
	_ = mBase
	var v8 int32
	_ = v8
	var v19 int32
	_ = v19
	var v21 int32
	_ = v21
	var v25 int32
	_ = v25
	var v28 int32
	_ = v28
	var v34 int32
	_ = v34
	var v39 int32
	_ = v39
	var v41 int32
	_ = v41
	var v45 int32
	_ = v45
	var v49 int32
	_ = v49
	var v50 int32
	_ = v50
	var v52 int32
	_ = v52
	var v61 int32
	_ = v61
	var v62 int32
	_ = v62
	var v69 int32
	_ = v69
	var v76 int32
	_ = v76
	var v82 int32
	_ = v82
	var v99 int32
	_ = v99
	var v100 int32
	_ = v100
	var v104 int32
	_ = v104
	var v107 int32
	_ = v107
	var v108 int32
	_ = v108
	var v109 int32
	_ = v109
	var v111 int32
	_ = v111
	var v113 int32
	_ = v113
	var v131 int32
	_ = v131
	var v132 int32
	_ = v132
	var v133 int32
	_ = v133
	var v142 int32
	_ = v142
	var v143 int32
	_ = v143
	var v146 int32
	_ = v146
	var v149 int32
	_ = v149
	var v152 int32
	_ = v152
	var v153 int32
	_ = v153
	var v156 int32
	_ = v156
	var v159 int32
	_ = v159
	var v162 int32
	_ = v162
	var v163 int32
	_ = v163
	var v164 int32
	_ = v164
	var v166 int32
	_ = v166
	var v180 int32
	_ = v180
	var v181 int32
	_ = v181
	var v190 int32
	_ = v190
	var v191 int32
	_ = v191
	var v194 int32
	_ = v194
	var v197 int32
	_ = v197
	var v208 int32
	_ = v208
	var v210 int32
	_ = v210
	var v216 int32
	_ = v216
	var v222 int32
	_ = v222
	var v223 int32
	_ = v223
	var v224 int32
	_ = v224
	var v227 int64
	_ = v227
	var v228 int64
	_ = v228
	var v236 int64
	_ = v236
	var v238 int32
	_ = v238
	var v239 int32
	_ = v239
	var v242 int32
	_ = v242
	var v247 int32
	_ = v247
	var v253 int32
	_ = v253
	var v261 int32
	_ = v261
	var v269 int32
	_ = v269
	var v272 int32
	_ = v272
	var v273 int32
	_ = v273
	var v277 int32
	_ = v277
	var v278 int32
	_ = v278
	var v279 int64
	_ = v279
	var v281 int32
	_ = v281
	var v291 int32
	_ = v291
	var v292 int32
	_ = v292
	var v293 int32
	_ = v293
	var v299 int32
	_ = v299
	var v304 int32
	_ = v304
	var v305 int32
	_ = v305
	var v310 int32
	_ = v310
	var v313 int64
	_ = v313
	var v323 int32
	_ = v323
	var v325 int32
	_ = v325
	var v327 int32
	_ = v327
	var v332 int32
	_ = v332
	var v342 int32
	_ = v342
	var v345 int32
	_ = v345
	var v359 int32
	_ = v359
	var v363 int32
	_ = v363
	var v364 int32
	_ = v364
	var v367 int32
	_ = v367
	var v369 int32
	_ = v369
	var v371 int32
	_ = v371
	var v388 int32
	_ = v388
	var v389 int32
	_ = v389
	var v390 int32
	_ = v390
	var v399 int32
	_ = v399
	var v400 int32
	_ = v400
	var v403 int32
	_ = v403
	var v406 int32
	_ = v406
	var v409 int32
	_ = v409
	var v410 int32
	_ = v410
	var v413 int32
	_ = v413
	var v416 int32
	_ = v416
	var v419 int32
	_ = v419
	var v420 int32
	_ = v420
	var v421 int32
	_ = v421
	var v423 int32
	_ = v423
	var v437 int32
	_ = v437
	var v438 int32
	_ = v438
	var v447 int32
	_ = v447
	var v448 int32
	_ = v448
	var v451 int32
	_ = v451
	var v454 int32
	_ = v454
	var v467 int32
	_ = v467
	var v478 int32
	_ = v478
	var v482 int32
	_ = v482
	var v486 int32
	_ = v486
	var v489 int32
	_ = v489
	var v491 int32
	_ = v491
	var v495 int32
	_ = v495
	var v498 int32
	_ = v498
	var v499 int32
	_ = v499
	var v504 int32
	_ = v504
	var v508 int32
	_ = v508
	var v513 int32
	_ = v513
	var v518 int32
	_ = v518
	var v519 int32
	_ = v519
	var v523 int32
	_ = v523
	var v525 int64
	_ = v525
	var v535 int32
	_ = v535
	var v537 int32
	_ = v537
	var v540 int32
	_ = v540
	var v541 int64
	_ = v541
	var v555 int32
	_ = v555
	var v558 int32
	_ = v558
	var v562 int32
	_ = v562
	var v565 int32
	_ = v565
	var v575 int32
	_ = v575
	var v576 int32
	_ = v576
	var v577 int32
	_ = v577
	var v583 int32
	_ = v583
	var v584 int32
	_ = v584
	var v592 int32
	_ = v592
	var v593 int32
	_ = v593
	var v599 int32
	_ = v599
	var v600 int32
	_ = v600
	var v607 int32
	_ = v607
	var v608 int32
	_ = v608
	var v617 int32
	_ = v617
	var v618 int32
	_ = v618
	var v624 int32
	_ = v624
	var v625 int32
	_ = v625
	var v629 int32
	_ = v629
	var v633 int32
	_ = v633
	var v638 int32
	_ = v638
	var v644 int32
	_ = v644
	var v645 int32
	_ = v645
	var v653 int32
	_ = v653
	var v654 int32
	_ = v654
	var v660 int32
	_ = v660
	var v661 int32
	_ = v661
	var v663 int32
	_ = v663
	var v668 int32
	_ = v668
	var v674 int32
	_ = v674
	var v675 int32
	_ = v675
	var v678 int32
	_ = v678
	var v680 int32
	_ = v680
	var v684 int32
	_ = v684
	var v685 int32
	_ = v685
	var v686 int32
	_ = v686
	var v692 int64
	_ = v692
	var v701 int32
	_ = v701
	var v705 int32
	_ = v705
	var v708 int32
	_ = v708
	var v709 int32
	_ = v709
	var v714 int32
	_ = v714
	var v718 int32
	_ = v718
	var v725 int32
	_ = v725
	var v730 int32
	_ = v730
	var v731 int32
	_ = v731
	var v743 int32
	_ = v743
	var v752 int32
	_ = v752
	var v754 int32
	_ = v754
	var v756 int32
	_ = v756
	var v760 int32
	_ = v760
	var v761 int32
	_ = v761
	var v762 int32
	_ = v762
	var v765 int32
	_ = v765
	var v767 int32
	_ = v767
	var v771 int32
	_ = v771
	var v774 int32
	_ = v774
	var v775 int32
	_ = v775
	var v778 int32
	_ = v778
	var v780 int32
	_ = v780
	var v784 int32
	_ = v784
	var v785 int32
	_ = v785
	var v786 int32
	_ = v786
	var v790 int32
	_ = v790
	var v796 int64
	_ = v796
	var v804 int32
	_ = v804
	var v806 int32
	_ = v806
	var v810 int32
	_ = v810
	var v816 int32
	_ = v816
	var v817 int32
	_ = v817
	var v820 int32
	_ = v820
	var v823 int32
	_ = v823
	var v827 int32
	_ = v827
	var v829 int32
	_ = v829
	var v831 int32
	_ = v831
	var v834 int32
	_ = v834
	var v837 int32
	_ = v837
	var v841 int32
	_ = v841
	var v845 int32
	_ = v845
	var v849 int32
	_ = v849
	var v857 int32
	_ = v857
	var v861 int32
	_ = v861
	var v865 int32
	_ = v865
	var v866 int32
	_ = v866
	var v872 int32
	_ = v872
	var v875 int32
	_ = v875
	var v877 int32
	_ = v877
	var v881 int32
	_ = v881
	var v885 int32
	_ = v885
	var v888 int32
	_ = v888
	var v892 int32
	_ = v892
	var v897 int32
	_ = v897
	var v901 int32
	_ = v901
	v8 = int32(0)
	v19 = m.G0
	v21 = v19 - int32(1568)
	m.G0 = v21
	v25 = F_errstart(m, int32(14), v8)
	mBase = m.M
	v28 = m.ExcPending
	if v28 != 0 {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	return int32(0)
L2:
	;
	if v25 != 0 {
		goto L3
	} else {
		goto L4
	}
L3:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v21)+96)) = l3
	F_errmsg_internal(m, int32(_a_F_logicalrep_worker_launch_0), v21+int32(96))
	mBase = m.M
	v34 = m.ExcPending
	if v34 != 0 {
		goto L1
	} else {
		goto L6
	}
L4:
	;
	goto L5
L5:
	;
	v41 = *(*int32)(unsafe.Add(mBase, _c_F_logicalrep_worker_launch[0]))
	if v41 != 0 {
		goto L9
	} else {
		goto L10
	}
L6:
	;
	F_errfinish(m, int32(_a_F_logicalrep_worker_launch_1), int32(338), int32(_a_F_logicalrep_worker_launch_2))
	mBase = m.M
	v39 = m.ExcPending
	if v39 != 0 {
		goto L1
	} else {
		goto L7
	}
L7:
	;
	goto L5
L8:
	;
	m.G0 = v21 + int32(1568)
	return v901
L9:
	;
	v45 = *(*int32)(unsafe.Add(mBase, _c_F_logicalrep_worker_launch[1]))
	v49 = F_LWLockAcquire(m, v45+int32(_a_F_logicalrep_worker_launch_3), int32(0))
	mBase = m.M
	v50 = m.ExcPending
	if v50 != 0 {
		goto L1
	} else {
		goto L12
	}
L10:
	;
	goto L11
L11:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v885 = m.ExcPending
	if v885 != 0 {
		goto L1
	} else {
		goto L166
	}
L12:
	;
	v52 = *(*int32)(unsafe.Add(mBase, _c_F_logicalrep_worker_launch[2]))
	v61 = v8
	v62 = v52
	v69 = v8
	goto L14
L13:
	;
	if base.B2i32(l0 != int32(1))|base.B2i32(v210 < v345) == int32(0) {
		goto L60
	} else {
		goto L61
	}
L14:
	;
	if v62 <= int32(0) {
		goto L17
	} else {
		goto L18
	}
L15:
	;
	v332 = *(*int32)(unsafe.Add(mBase, _c_F_logicalrep_worker_launch[3]))
	v342 = v327
	v345 = v332
	goto L13
L16:
	;
	v222 = m.G0
	v223 = int32(16)
	v224 = v222 - v223
	m.G0 = v224
	F_gettimeofday(m, v224)
	mBase = m.M
	v227 = *(*int64)(unsafe.Add(mBase, uint32(v224)))
	v228 = int64(*(*int32)(unsafe.Add(mBase, uint32(v224)+8)))
	m.G0 = v224 + v223
	v236 = v228 + v227*int64(1000000) - int64(946684800000000)
	goto L42
L17:
	;
	v208 = v61
	v210 = int32(0)
	v216 = v69
	goto L16
L18:
	;
	goto L19
L19:
	;
	v76 = *(*int32)(unsafe.Add(mBase, _c_F_logicalrep_worker_launch[4]))
	v82 = int32(0)
	goto L20
L20:
	;
	v99 = v76 + int32(16) + v82*int32(112)
	v100 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v99)+16)))
	if v100 != int32(1) {
		goto L23
	} else {
		goto L24
	}
L21:
	;
	v109 = int32(0)
	v111 = *(*int32)(unsafe.Add(mBase, _c_F_logicalrep_worker_launch[4]))
	v113 = v111 + int32(16)
	if v62 != int32(1) {
		goto L27
	} else {
		goto L28
	}
L22:
	;
	goto L21
L23:
	;
	v107 = v99
	v108 = v82
	goto L22
L24:
	;
	goto L25
L25:
	;
	v104 = v82 + int32(1)
	if v104 != v62 {
		v82 = v104
		goto L20
	} else {
		goto L26
	}
L26:
	;
	v107 = v61
	v108 = v69
	goto L22
L27:
	;
	v131 = int32(0)
	v132 = v109
	v133 = v109
	goto L30
L28:
	;
	v180 = v109
	v181 = v109
	goto L29
L29:
	;
	v190 = v113 + v181*int32(112)
	v191 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v190)+16)))
	if v191 != int32(1) {
		v208 = v107
		v210 = v180
		v216 = v108
		goto L16
	} else {
		goto L40
	}
L30:
	;
	v142 = v113 + v133*int32(112)
	v143 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v142)+16)))
	if v143 != int32(1) {
		v152 = v132
		goto L32
	} else {
		goto L33
	}
L31:
	;
	if v62&int32(1) == int32(0) {
		v208 = v107
		v210 = v162
		v216 = v108
		goto L16
	} else {
		goto L39
	}
L32:
	;
	v153 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v142)+128)))
	if v153 != int32(1) {
		v162 = v152
		goto L35
	} else {
		goto L36
	}
L33:
	;
	v146 = *(*int32)(unsafe.Add(mBase, uint32(v142)))
	if v146 != int32(1) {
		v152 = v132
		goto L32
	} else {
		goto L34
	}
L34:
	;
	v149 = *(*int32)(unsafe.Add(mBase, uint32(v142)+32))
	v152 = v132 + base.B2i32(v149 == l2)
	goto L32
L35:
	;
	v163 = int32(2)
	v164 = v133 + v163
	v166 = v131 + v163
	if v166 != v62&int32(2147483646) {
		v131 = v166
		v132 = v162
		v133 = v164
		goto L30
	} else {
		goto L38
	}
L36:
	;
	v156 = *(*int32)(unsafe.Add(mBase, uint32(v142)+112))
	if v156 != int32(1) {
		v162 = v152
		goto L35
	} else {
		goto L37
	}
L37:
	;
	v159 = *(*int32)(unsafe.Add(mBase, uint32(v142)+144))
	v162 = v152 + base.B2i32(v159 == l2)
	goto L35
L38:
	;
	goto L31
L39:
	;
	v180 = v162
	v181 = v164
	goto L29
L40:
	;
	v194 = *(*int32)(unsafe.Add(mBase, uint32(v190)))
	if v194 != int32(1) {
		v208 = v107
		v210 = v180
		v216 = v108
		goto L16
	} else {
		goto L41
	}
L41:
	;
	v197 = *(*int32)(unsafe.Add(mBase, uint32(v190)+32))
	v208 = v107
	v210 = v180 + base.B2i32(v197 == l2)
	v216 = v108
	goto L16
L42:
	;
	v238 = *(*int32)(unsafe.Add(mBase, _c_F_logicalrep_worker_launch[2]))
	v239 = int32(0)
	v242 = *(*int32)(unsafe.Add(mBase, _c_F_logicalrep_worker_launch[3]))
	if base.B2i32(v208 == v239)|base.B2i32(v242 <= v210) == v239 {
		v342 = v238
		v345 = v242
		goto L13
	} else {
		goto L43
	}
L43:
	;
	v247 = int32(0)
	if v238 <= v247 {
		v342 = v238
		v345 = v242
		goto L13
	} else {
		goto L44
	}
L44:
	;
	v253 = int32(0)
	v261 = v247
	goto L45
L45:
	;
	v269 = *(*int32)(unsafe.Add(mBase, _c_F_logicalrep_worker_launch[4]))
	v272 = v269 + v253*int32(112)
	v273 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v272)+32)))
	if v273 != int32(1) {
		v323 = v261
		goto L47
	} else {
		goto L48
	}
L46:
	;
	if v323&int32(1) != 0 {
		v61 = v208
		v62 = v327
		v69 = v216
		goto L14
	} else {
		goto L59
	}
L47:
	;
	v325 = v253 + int32(1)
	v327 = *(*int32)(unsafe.Add(mBase, _c_F_logicalrep_worker_launch[2]))
	if v325 < v327 {
		v253 = v325
		v261 = v323
		goto L45
	} else {
		goto L58
	}
L48:
	;
	v277 = v272 + int32(16)
	v278 = *(*int32)(unsafe.Add(mBase, uint32(v277)+20))
	if v278 != 0 {
		v323 = v261
		goto L47
	} else {
		goto L49
	}
L49:
	;
	v279 = *(*int64)(unsafe.Add(mBase, uint32(v277)+8))
	v281 = *(*int32)(unsafe.Add(mBase, _c_F_logicalrep_worker_launch[5]))
	goto L50
L50:
	;
	if base.B2i32(base.I64_extend_i32_s(v281)*int64(1000) <= v236-v279) == int32(0) {
		v323 = v261
		goto L47
	} else {
		goto L51
	}
L51:
	;
	v291 = F_errstart(m, int32(19), int32(0))
	mBase = m.M
	v292 = m.ExcPending
	if v292 != 0 {
		goto L1
	} else {
		goto L52
	}
L52:
	;
	if v291 != 0 {
		goto L53
	} else {
		goto L54
	}
L53:
	;
	v293 = *(*int32)(unsafe.Add(mBase, uint32(v277)+32))
	*(*int32)(unsafe.Add(mBase, uint32(v21)+80)) = v293
	F_errmsg_internal(m, int32(_a_F_logicalrep_worker_launch_4), v21+int32(80))
	mBase = m.M
	v299 = m.ExcPending
	if v299 != 0 {
		goto L1
	} else {
		goto L56
	}
L54:
	;
	goto L55
L55:
	;
	v305 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v277)+16)) = uint8(v305)
	*(*int32)(unsafe.Add(mBase, uint32(v277))) = v305
	v310 = v272 + int32(36)
	*(*int32)(unsafe.Add(mBase, uint32(v310)+16)) = v305
	v313 = int64(0)
	*(*int64)(unsafe.Add(mBase, uint32(v310)+8)) = v313
	*(*int64)(unsafe.Add(mBase, uint32(v310))) = v313
	*(*uint8)(unsafe.Add(mBase, uint32(v277)+68)) = uint8(v305)
	*(*int32)(unsafe.Add(mBase, uint32(v277)+64)) = int32(-1)
	v323 = int32(1)
	goto L47
L56:
	;
	F_errfinish(m, int32(_a_F_logicalrep_worker_launch_1), int32(393), int32(_a_F_logicalrep_worker_launch_2))
	mBase = m.M
	v304 = m.ExcPending
	if v304 != 0 {
		goto L1
	} else {
		goto L57
	}
L57:
	;
	goto L55
L58:
	;
	goto L46
L59:
	;
	goto L15
L60:
	;
	v359 = *(*int32)(unsafe.Add(mBase, _c_F_logicalrep_worker_launch[1]))
	F_LWLockRelease(m, v359+int32(_a_F_logicalrep_worker_launch_3))
	mBase = m.M
	v363 = m.ExcPending
	if v363 != 0 {
		goto L1
	} else {
		goto L63
	}
L61:
	;
	goto L62
L62:
	;
	v364 = int32(0)
	if v342 <= v364 {
		v467 = v364
		goto L64
	} else {
		goto L65
	}
L63:
	;
	v901 = int32(0)
	goto L8
L64:
	;
	if l0 != int32(3) {
		goto L81
	} else {
		goto L82
	}
L65:
	;
	v367 = int32(0)
	v369 = *(*int32)(unsafe.Add(mBase, _c_F_logicalrep_worker_launch[4]))
	v371 = v369 + int32(16)
	if v342 != int32(1) {
		goto L66
	} else {
		goto L67
	}
L66:
	;
	v388 = int32(0)
	v389 = v364
	v390 = v367
	goto L69
L67:
	;
	v437 = v364
	v438 = v367
	goto L68
L68:
	;
	v447 = v371 + v438*int32(112)
	v448 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v447)+16)))
	if v448 != int32(1) {
		v467 = v437
		goto L64
	} else {
		goto L79
	}
L69:
	;
	v399 = v371 + v390*int32(112)
	v400 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v399)+16)))
	if v400 != int32(1) {
		v409 = v389
		goto L71
	} else {
		goto L72
	}
L70:
	;
	if v342&int32(1) == int32(0) {
		v467 = v419
		goto L64
	} else {
		goto L78
	}
L71:
	;
	v410 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v399)+128)))
	if v410 != int32(1) {
		v419 = v409
		goto L74
	} else {
		goto L75
	}
L72:
	;
	v403 = *(*int32)(unsafe.Add(mBase, uint32(v399)))
	if v403 != int32(3) {
		v409 = v389
		goto L71
	} else {
		goto L73
	}
L73:
	;
	v406 = *(*int32)(unsafe.Add(mBase, uint32(v399)+32))
	v409 = v389 + base.B2i32(v406 == l2)
	goto L71
L74:
	;
	v420 = int32(2)
	v421 = v390 + v420
	v423 = v388 + v420
	if v423 != v342&int32(2147483646) {
		v388 = v423
		v389 = v419
		v390 = v421
		goto L69
	} else {
		goto L77
	}
L75:
	;
	v413 = *(*int32)(unsafe.Add(mBase, uint32(v399)+112))
	if v413 != int32(3) {
		v419 = v409
		goto L74
	} else {
		goto L76
	}
L76:
	;
	v416 = *(*int32)(unsafe.Add(mBase, uint32(v399)+144))
	v419 = v409 + base.B2i32(v416 == l2)
	goto L74
L77:
	;
	goto L70
L78:
	;
	v437 = v419
	v438 = v421
	goto L68
L79:
	;
	v451 = *(*int32)(unsafe.Add(mBase, uint32(v447)))
	if v451 != int32(3) {
		v467 = v437
		goto L64
	} else {
		goto L80
	}
L80:
	;
	v454 = *(*int32)(unsafe.Add(mBase, uint32(v447)+32))
	v467 = v437 + base.B2i32(v454 == l2)
	goto L64
L81:
	;
	if v208 == int32(0) {
		goto L85
	} else {
		goto L86
	}
L82:
	;
	v478 = *(*int32)(unsafe.Add(mBase, _c_F_logicalrep_worker_launch[6]))
	if v467 < v478 {
		goto L81
	} else {
		goto L83
	}
L83:
	;
	v482 = *(*int32)(unsafe.Add(mBase, _c_F_logicalrep_worker_launch[1]))
	F_LWLockRelease(m, v482+int32(_a_F_logicalrep_worker_launch_3))
	mBase = m.M
	v486 = m.ExcPending
	if v486 != 0 {
		goto L1
	} else {
		goto L84
	}
L84:
	;
	v901 = int32(0)
	goto L8
L85:
	;
	v489 = int32(0)
	v491 = *(*int32)(unsafe.Add(mBase, _c_F_logicalrep_worker_launch[1]))
	F_LWLockRelease(m, v491+int32(_a_F_logicalrep_worker_launch_3))
	mBase = m.M
	v495 = m.ExcPending
	if v495 != 0 {
		goto L1
	} else {
		goto L88
	}
L86:
	;
	goto L87
L87:
	;
	v519 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(v208)+16)) = uint8(v519)
	*(*int64)(unsafe.Add(mBase, uint32(v208)+8)) = v236
	*(*int32)(unsafe.Add(mBase, uint32(v208))) = l0
	v523 = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v208)+60)) = v523
	v525 = int64(0)
	*(*int64)(unsafe.Add(mBase, uint32(v208)+48)) = v525
	*(*uint8)(unsafe.Add(mBase, uint32(v208)+40)) = uint8(v523)
	*(*int32)(unsafe.Add(mBase, uint32(v208)+36)) = l5
	*(*int32)(unsafe.Add(mBase, uint32(v208)+32)) = l2
	*(*int32)(unsafe.Add(mBase, uint32(v208)+28)) = l4
	*(*int32)(unsafe.Add(mBase, uint32(v208)+24)) = l1
	*(*int32)(unsafe.Add(mBase, uint32(v208)+20)) = v523
	v535 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v208)+18)))
	v537 = v535 + v519
	*(*uint16)(unsafe.Add(mBase, uint32(v208)+18)) = uint16(v537)
	v540 = *(*int32)(unsafe.Add(mBase, _c_F_logicalrep_worker_launch[7]))
	v541 = int64(-9223372036854775807 - 1)
	*(*int64)(unsafe.Add(mBase, uint32(v208)+104)) = v541
	*(*int64)(unsafe.Add(mBase, uint32(v208)+96)) = v525
	*(*int64)(unsafe.Add(mBase, uint32(v208)+88)) = v541
	*(*int64)(unsafe.Add(mBase, uint32(v208)+80)) = v541
	*(*int64)(unsafe.Add(mBase, uint32(v208)+72)) = v525
	*(*uint8)(unsafe.Add(mBase, uint32(v208)+68)) = uint8(base.B2i32(l0 == int32(3)))
	if l0 != int32(3) {
		goto L95
	} else {
		goto L96
	}
L88:
	;
	v498 = F_errstart(m, int32(19), int32(0))
	mBase = m.M
	v499 = m.ExcPending
	if v499 != 0 {
		goto L1
	} else {
		goto L89
	}
L89:
	;
	if v498 == int32(0) {
		v901 = v489
		goto L8
	} else {
		goto L90
	}
L90:
	;
	F_errcode(m, int32(_a_F_logicalrep_worker_launch_5))
	mBase = m.M
	v504 = m.ExcPending
	if v504 != 0 {
		goto L1
	} else {
		goto L91
	}
L91:
	;
	F_errmsg(m, int32(_a_F_logicalrep_worker_launch_6), int32(0))
	mBase = m.M
	v508 = m.ExcPending
	if v508 != 0 {
		goto L1
	} else {
		goto L92
	}
L92:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v21))) = int32(_a_F_logicalrep_worker_launch_7)
	F_errhint(m, int32(_a_F_logicalrep_worker_launch_8), v21)
	mBase = m.M
	v513 = m.ExcPending
	if v513 != 0 {
		goto L1
	} else {
		goto L93
	}
L93:
	;
	F_errfinish(m, int32(_a_F_logicalrep_worker_launch_1), int32(438), int32(_a_F_logicalrep_worker_launch_2))
	mBase = m.M
	v518 = m.ExcPending
	if v518 != 0 {
		goto L1
	} else {
		goto L94
	}
L94:
	;
	v901 = v489
	goto L8
L95:
	;
	v555 = int32(-1)
	goto L97
L96:
	;
	v555 = v540
	goto L97
L97:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v208)+64)) = v555
	v558 = *(*int32)(unsafe.Add(mBase, _c_F_logicalrep_worker_launch[1]))
	F_LWLockRelease(m, v558+int32(_a_F_logicalrep_worker_launch_3))
	mBase = m.M
	v562 = m.ExcPending
	if v562 != 0 {
		goto L1
	} else {
		goto L98
	}
L98:
	;
	v565 = int32(0)
	base.MemoryFill(m, v21+int32(104), v565, int32(1460))
	*(*int64)(unsafe.Add(mBase, uint32(v21)+296)) = int64(8589934595)
	v575 = F_pg_snprintf(m, v21+int32(308), int32(1024), int32(_a_F_logicalrep_worker_launch_9), v565)
	mBase = m.M
	v576 = m.ExcPending
	if v576 != 0 {
		goto L1
	} else {
		goto L99
	}
L99:
	;
	v577 = *(*int32)(unsafe.Add(mBase, uint32(v208)))
	switch v577 {
	case 0:
		goto L102
	case 1:
		goto L103
	case 2:
		goto L101
	case 3:
		goto L104
	default:
		goto L100
	}
L100:
	;
	v663 = v208 + int32(20)
	*(*int32)(unsafe.Add(mBase, uint32(v21)+304)) = int32(-1)
	*(*int32)(unsafe.Add(mBase, uint32(v21)+1428)) = v216
	v668 = *(*int32)(unsafe.Add(mBase, _c_F_logicalrep_worker_launch[7]))
	*(*int32)(unsafe.Add(mBase, uint32(v21)+1560)) = v668
	v674 = F_RegisterDynamicBackgroundWorker(m, v21+int32(104), v21+int32(100))
	mBase = m.M
	v675 = m.ExcPending
	if v675 != 0 {
		goto L1
	} else {
		goto L117
	}
L101:
	;
	v644 = F_pg_snprintf(m, v21+int32(1332), int32(96), int32(_a_F_logicalrep_worker_launch_10), int32(0))
	mBase = m.M
	v645 = m.ExcPending
	if v645 != 0 {
		goto L1
	} else {
		goto L114
	}
L102:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v629 = m.ExcPending
	if v629 != 0 {
		goto L1
	} else {
		goto L111
	}
L103:
	;
	v607 = F_pg_snprintf(m, v21+int32(1332), int32(96), int32(_a_F_logicalrep_worker_launch_11), int32(0))
	mBase = m.M
	v608 = m.ExcPending
	if v608 != 0 {
		goto L1
	} else {
		goto L108
	}
L104:
	;
	v583 = F_pg_snprintf(m, v21+int32(1332), int32(96), int32(_a_F_logicalrep_worker_launch_12), int32(0))
	mBase = m.M
	v584 = m.ExcPending
	if v584 != 0 {
		goto L1
	} else {
		goto L105
	}
L105:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v21)+48)) = l2
	v592 = F_pg_snprintf(m, v21+int32(104), int32(96), int32(_a_F_logicalrep_worker_launch_13), v21+int32(48))
	mBase = m.M
	v593 = m.ExcPending
	if v593 != 0 {
		goto L1
	} else {
		goto L106
	}
L106:
	;
	v599 = F_pg_snprintf(m, v21+int32(200), int32(96), int32(_a_F_logicalrep_worker_launch_14), int32(0))
	mBase = m.M
	v600 = m.ExcPending
	if v600 != 0 {
		goto L1
	} else {
		goto L107
	}
L107:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v21)+1432)) = l6
	goto L100
L108:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v21)+68)) = l5
	*(*int32)(unsafe.Add(mBase, uint32(v21)+64)) = l2
	v617 = F_pg_snprintf(m, v21+int32(104), int32(96), int32(_a_F_logicalrep_worker_launch_15), v21-int32(-64))
	mBase = m.M
	v618 = m.ExcPending
	if v618 != 0 {
		goto L1
	} else {
		goto L109
	}
L109:
	;
	v624 = F_pg_snprintf(m, v21+int32(200), int32(96), int32(_a_F_logicalrep_worker_launch_16), int32(0))
	mBase = m.M
	v625 = m.ExcPending
	if v625 != 0 {
		goto L1
	} else {
		goto L110
	}
L110:
	;
	goto L100
L111:
	;
	F_errmsg_internal(m, int32(_a_F_logicalrep_worker_launch_17), int32(0))
	mBase = m.M
	v633 = m.ExcPending
	if v633 != 0 {
		goto L1
	} else {
		goto L112
	}
L112:
	;
	F_errfinish(m, int32(_a_F_logicalrep_worker_launch_1), int32(506), int32(_a_F_logicalrep_worker_launch_2))
	mBase = m.M
	v638 = m.ExcPending
	if v638 != 0 {
		goto L1
	} else {
		goto L113
	}
L113:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L114:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v21)+32)) = l2
	v653 = F_pg_snprintf(m, v21+int32(104), int32(96), int32(_a_F_logicalrep_worker_launch_18), v21+int32(32))
	mBase = m.M
	v654 = m.ExcPending
	if v654 != 0 {
		goto L1
	} else {
		goto L115
	}
L115:
	;
	v660 = F_pg_snprintf(m, v21+int32(200), int32(96), int32(_a_F_logicalrep_worker_launch_19), int32(0))
	mBase = m.M
	v661 = m.ExcPending
	if v661 != 0 {
		goto L1
	} else {
		goto L116
	}
L116:
	;
	goto L100
L117:
	;
	if v674 == int32(0) {
		goto L118
	} else {
		goto L119
	}
L118:
	;
	v678 = int32(0)
	v680 = *(*int32)(unsafe.Add(mBase, _c_F_logicalrep_worker_launch[1]))
	v684 = F_LWLockAcquire(m, v680+int32(_a_F_logicalrep_worker_launch_3), v678)
	mBase = m.M
	v685 = m.ExcPending
	if v685 != 0 {
		goto L1
	} else {
		goto L121
	}
L119:
	;
	goto L120
L120:
	;
	v731 = *(*int32)(unsafe.Add(mBase, uint32(v21)+100))
	v743 = int32(0)
	goto L129
L121:
	;
	v686 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v208)+16)) = uint8(v686)
	*(*int32)(unsafe.Add(mBase, uint32(v208))) = v686
	*(*int32)(unsafe.Add(mBase, uint32(v663)+16)) = v686
	v692 = int64(0)
	*(*int64)(unsafe.Add(mBase, uint32(v663)+8)) = v692
	*(*int64)(unsafe.Add(mBase, uint32(v663))) = v692
	*(*uint8)(unsafe.Add(mBase, uint32(v208)+68)) = uint8(v686)
	*(*int32)(unsafe.Add(mBase, uint32(v208)+64)) = int32(-1)
	v701 = *(*int32)(unsafe.Add(mBase, _c_F_logicalrep_worker_launch[1]))
	F_LWLockRelease(m, v701+int32(_a_F_logicalrep_worker_launch_3))
	mBase = m.M
	v705 = m.ExcPending
	if v705 != 0 {
		goto L1
	} else {
		goto L122
	}
L122:
	;
	v708 = F_errstart(m, int32(19), int32(0))
	mBase = m.M
	v709 = m.ExcPending
	if v709 != 0 {
		goto L1
	} else {
		goto L123
	}
L123:
	;
	if v708 == int32(0) {
		v901 = v678
		goto L8
	} else {
		goto L124
	}
L124:
	;
	F_errcode(m, int32(_a_F_logicalrep_worker_launch_5))
	mBase = m.M
	v714 = m.ExcPending
	if v714 != 0 {
		goto L1
	} else {
		goto L125
	}
L125:
	;
	F_errmsg(m, int32(_a_F_logicalrep_worker_launch_20), int32(0))
	mBase = m.M
	v718 = m.ExcPending
	if v718 != 0 {
		goto L1
	} else {
		goto L126
	}
L126:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v21)+16)) = int32(_a_F_logicalrep_worker_launch_21)
	F_errhint(m, int32(_a_F_logicalrep_worker_launch_8), v21+int32(16))
	mBase = m.M
	v725 = m.ExcPending
	if v725 != 0 {
		goto L1
	} else {
		goto L127
	}
L127:
	;
	F_errfinish(m, int32(_a_F_logicalrep_worker_launch_1), int32(524), int32(_a_F_logicalrep_worker_launch_2))
	mBase = m.M
	v730 = m.ExcPending
	if v730 != 0 {
		goto L1
	} else {
		goto L128
	}
L128:
	;
	v901 = v678
	goto L8
L129:
	;
	v752 = *(*int32)(unsafe.Add(mBase, _c_F_logicalrep_worker_launch[8]))
	if v752 != 0 {
		goto L131
	} else {
		goto L132
	}
L131:
	;
	F_ProcessInterrupts(m)
	mBase = m.M
	v754 = m.ExcPending
	if v754 != 0 {
		goto L1
	} else {
		goto L134
	}
L132:
	;
	goto L133
L133:
	;
	v756 = *(*int32)(unsafe.Add(mBase, _c_F_logicalrep_worker_launch[1]))
	v760 = F_LWLockAcquire(m, v756+int32(_a_F_logicalrep_worker_launch_3), int32(1))
	mBase = m.M
	v761 = m.ExcPending
	if v761 != 0 {
		goto L1
	} else {
		goto L135
	}
L134:
	;
	goto L133
L135:
	;
	v762 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v208)+16)))
	if v762 != int32(1) {
		v804 = v762
		goto L137
	} else {
		goto L138
	}
L136:
	;
	v861 = *(*int32)(unsafe.Add(mBase, _c_F_logicalrep_worker_launch[9]))
	v865 = F_WaitLatch(m, v861, int32(41), int32(10), int32(134217734))
	mBase = m.M
	v866 = m.ExcPending
	if v866 != 0 {
		goto L1
	} else {
		goto L161
	}
L137:
	;
	v806 = *(*int32)(unsafe.Add(mBase, _c_F_logicalrep_worker_launch[1]))
	F_LWLockRelease(m, v806+int32(_a_F_logicalrep_worker_launch_3))
	mBase = m.M
	v810 = m.ExcPending
	if v810 != 0 {
		goto L1
	} else {
		goto L145
	}
L138:
	;
	v765 = *(*int32)(unsafe.Add(mBase, uint32(v663)))
	if v765 != 0 {
		v804 = v762
		goto L137
	} else {
		goto L139
	}
L139:
	;
	v767 = *(*int32)(unsafe.Add(mBase, _c_F_logicalrep_worker_launch[1]))
	F_LWLockRelease(m, v767+int32(_a_F_logicalrep_worker_launch_3))
	mBase = m.M
	v771 = m.ExcPending
	if v771 != 0 {
		goto L1
	} else {
		goto L140
	}
L140:
	;
	v774 = F_GetBackgroundWorkerPid(m, v731, v21+int32(1564))
	mBase = m.M
	v775 = m.ExcPending
	if v775 != 0 {
		goto L1
	} else {
		goto L141
	}
L141:
	;
	if v774 != int32(2) {
		goto L136
	} else {
		goto L142
	}
L142:
	;
	v778 = int32(0)
	v780 = *(*int32)(unsafe.Add(mBase, _c_F_logicalrep_worker_launch[1]))
	v784 = F_LWLockAcquire(m, v780+int32(_a_F_logicalrep_worker_launch_3), v778)
	mBase = m.M
	v785 = m.ExcPending
	if v785 != 0 {
		goto L1
	} else {
		goto L143
	}
L143:
	;
	v786 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v208)+18)))
	if v786 != v537&int32(_a_F_logicalrep_worker_launch_22) {
		v804 = v778
		goto L137
	} else {
		goto L144
	}
L144:
	;
	v790 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v208)+16)) = uint8(v790)
	*(*int32)(unsafe.Add(mBase, uint32(v208))) = v790
	*(*int32)(unsafe.Add(mBase, uint32(v663)+16)) = v790
	v796 = int64(0)
	*(*int64)(unsafe.Add(mBase, uint32(v663)+8)) = v796
	*(*int64)(unsafe.Add(mBase, uint32(v663))) = v796
	*(*uint8)(unsafe.Add(mBase, uint32(v208)+68)) = uint8(v790)
	*(*int32)(unsafe.Add(mBase, uint32(v208)+64)) = int32(-1)
	v804 = v778
	goto L137
L145:
	;
	if v743&int32(1) == int32(0) {
		v901 = v804
		goto L8
	} else {
		goto L146
	}
L146:
	;
	v816 = *(*int32)(unsafe.Add(mBase, _c_F_logicalrep_worker_launch[9]))
	v817 = *(*int32)(unsafe.Add(mBase, uint32(v816)))
	if v817 != 0 {
		goto L148
	} else {
		goto L149
	}
L147:
	;
	v901 = v804
	goto L8
L148:
	;
	goto L147
L149:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v816))) = int32(1)
	v820 = *(*int32)(unsafe.Add(mBase, uint32(v816)+4))
	if v820 == int32(0) {
		goto L148
	} else {
		goto L150
	}
L150:
	;
	v823 = *(*int32)(unsafe.Add(mBase, uint32(v816)+12))
	if v823 == int32(0) {
		goto L148
	} else {
		goto L151
	}
L151:
	;
	v827 = *(*int32)(unsafe.Add(mBase, _c_F_logicalrep_worker_launch[7]))
	if v827 == v823 {
		goto L152
	} else {
		goto L153
	}
L152:
	;
	v829 = m.G0
	v831 = v829 - int32(16)
	m.G0 = v831
	v834 = *(*int32)(unsafe.Add(mBase, _c_F_logicalrep_worker_launch[10]))
	if v834 == int32(0) {
		goto L155
	} else {
		goto L156
	}
L153:
	;
	goto L154
L154:
	;
	v857 = F_pgmem_kill(m, v823, int32(23))
	mBase = m.M
	goto L148
L155:
	;
	m.G0 = v831 + int32(16)
	goto L147
L156:
	;
	v837 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v831)+15)) = uint8(v837)
	goto L157
L157:
	;
	v841 = *(*int32)(unsafe.Add(mBase, _c_F_logicalrep_worker_launch[11]))
	v845 = F_write(m, v841, v831+int32(15), int32(1))
	mBase = m.M
	if int32(0) <= v845 {
		goto L155
	} else {
		goto L159
	}
L158:
	;
	goto L155
L159:
	;
	v849 = *(*int32)(unsafe.Add(mBase, _c_F_logicalrep_worker_launch[12]))
	if v849 == int32(27) {
		goto L157
	} else {
		goto L160
	}
L160:
	;
	goto L158
L161:
	;
	if v865&int32(1) == int32(0) {
		goto L129
	} else {
		goto L162
	}
L162:
	;
	v872 = *(*int32)(unsafe.Add(mBase, _c_F_logicalrep_worker_launch[9]))
	*(*int32)(unsafe.Add(mBase, uint32(v872))) = int32(0)
	goto L163
L163:
	;
	v875 = int32(1)
	v877 = *(*int32)(unsafe.Add(mBase, _c_F_logicalrep_worker_launch[8]))
	if v877 == int32(0) {
		v743 = v875
		goto L129
	} else {
		goto L164
	}
L164:
	;
	F_ProcessInterrupts(m)
	mBase = m.M
	v881 = m.ExcPending
	if v881 != 0 {
		goto L1
	} else {
		goto L165
	}
L165:
	;
	v743 = v875
	goto L129
L166:
	;
	F_errcode(m, int32(_a_F_logicalrep_worker_launch_5))
	mBase = m.M
	v888 = m.ExcPending
	if v888 != 0 {
		goto L1
	} else {
		goto L167
	}
L167:
	;
	F_errmsg(m, int32(_a_F_logicalrep_worker_launch_23), int32(0))
	mBase = m.M
	v892 = m.ExcPending
	if v892 != 0 {
		goto L1
	} else {
		goto L168
	}
L168:
	;
	F_errfinish(m, int32(_a_F_logicalrep_worker_launch_1), int32(344), int32(_a_F_logicalrep_worker_launch_2))
	mBase = m.M
	v897 = m.ExcPending
	if v897 != 0 {
		goto L1
	} else {
		goto L169
	}
L169:
	;
	base.Wasm_trap_unreachable()
	for {
	}
}
func F_logicalrep_worker_wakeup(m *base.Module, l0 int32) {
	mBase := m.M
	_ = mBase
	var v7 int32
	_ = v7
	var v11 int32
	_ = v11
	var v12 int32
	_ = v12
	var v14 int32
	_ = v14
	var v18 int32
	_ = v18
	var v23 int32
	_ = v23
	var v28 int32
	_ = v28
	var v29 int32
	_ = v29
	var v32 int32
	_ = v32
	var v35 int32
	_ = v35
	var v37 int32
	_ = v37
	var v38 int32
	_ = v38
	var v41 int32
	_ = v41
	var v44 int32
	_ = v44
	var v45 int32
	_ = v45
	var v48 int32
	_ = v48
	var v51 int32
	_ = v51
	var v55 int32
	_ = v55
	var v57 int32
	_ = v57
	var v59 int32
	_ = v59
	var v62 int32
	_ = v62
	var v65 int32
	_ = v65
	var v69 int32
	_ = v69
	var v73 int32
	_ = v73
	var v77 int32
	_ = v77
	var v85 int32
	_ = v85
	var v94 int32
	_ = v94
	var v98 int32
	_ = v98
	v7 = *(*int32)(unsafe.Add(mBase, _c_F_logicalrep_worker_wakeup[0]))
	v11 = F_LWLockAcquire(m, v7+int32(_a_F_logicalrep_worker_wakeup_0), int32(1))
	mBase = m.M
	v12 = m.ExcPending
	if v12 != 0 {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	return
L2:
	;
	v14 = *(*int32)(unsafe.Add(mBase, _c_F_logicalrep_worker_wakeup[1]))
	if v14 <= int32(0) {
		goto L3
	} else {
		goto L4
	}
L3:
	;
	v94 = *(*int32)(unsafe.Add(mBase, _c_F_logicalrep_worker_wakeup[0]))
	F_LWLockRelease(m, v94+int32(_a_F_logicalrep_worker_wakeup_0))
	mBase = m.M
	v98 = m.ExcPending
	if v98 != 0 {
		goto L1
	} else {
		goto L29
	}
L4:
	;
	v18 = *(*int32)(unsafe.Add(mBase, _c_F_logicalrep_worker_wakeup[2]))
	v23 = int32(0)
	goto L5
L5:
	;
	v28 = v18 + int32(16) + v23*int32(112)
	v29 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v28)+16)))
	if v29 != int32(1) {
		goto L8
	} else {
		goto L9
	}
L6:
	;
	v44 = v38 + int32(20)
	v45 = *(*int32)(unsafe.Add(mBase, uint32(v44)))
	if v45 != 0 {
		goto L16
	} else {
		goto L17
	}
L7:
	;
	goto L6
L8:
	;
	v41 = v23 + int32(1)
	if v41 != v14 {
		v23 = v41
		goto L5
	} else {
		goto L14
	}
L9:
	;
	v32 = *(*int32)(unsafe.Add(mBase, uint32(v28)))
	if v32 == int32(3) {
		goto L8
	} else {
		goto L10
	}
L10:
	;
	v35 = *(*int32)(unsafe.Add(mBase, uint32(v28)+32))
	if v35 != l0 {
		goto L8
	} else {
		goto L11
	}
L11:
	;
	v37 = *(*int32)(unsafe.Add(mBase, uint32(v28)+36))
	if v37 != 0 {
		goto L8
	} else {
		goto L12
	}
L12:
	;
	v38 = *(*int32)(unsafe.Add(mBase, uint32(v28)+20))
	if v38 != 0 {
		goto L7
	} else {
		goto L13
	}
L13:
	;
	goto L8
L14:
	;
	goto L3
L15:
	;
	goto L3
L16:
	;
	goto L15
L17:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v44))) = int32(1)
	v48 = *(*int32)(unsafe.Add(mBase, uint32(v44)+4))
	if v48 == int32(0) {
		goto L16
	} else {
		goto L18
	}
L18:
	;
	v51 = *(*int32)(unsafe.Add(mBase, uint32(v44)+12))
	if v51 == int32(0) {
		goto L16
	} else {
		goto L19
	}
L19:
	;
	v55 = *(*int32)(unsafe.Add(mBase, _c_F_logicalrep_worker_wakeup[3]))
	if v55 == v51 {
		goto L20
	} else {
		goto L21
	}
L20:
	;
	v57 = m.G0
	v59 = v57 - int32(16)
	m.G0 = v59
	v62 = *(*int32)(unsafe.Add(mBase, _c_F_logicalrep_worker_wakeup[4]))
	if v62 == int32(0) {
		goto L23
	} else {
		goto L24
	}
L21:
	;
	goto L22
L22:
	;
	v85 = F_pgmem_kill(m, v51, int32(23))
	mBase = m.M
	goto L16
L23:
	;
	m.G0 = v59 + int32(16)
	goto L15
L24:
	;
	v65 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v59)+15)) = uint8(v65)
	goto L25
L25:
	;
	v69 = *(*int32)(unsafe.Add(mBase, _c_F_logicalrep_worker_wakeup[5]))
	v73 = F_write(m, v69, v59+int32(15), int32(1))
	mBase = m.M
	if int32(0) <= v73 {
		goto L23
	} else {
		goto L27
	}
L26:
	;
	goto L23
L27:
	;
	v77 = *(*int32)(unsafe.Add(mBase, _c_F_logicalrep_worker_wakeup[6]))
	if v77 == int32(27) {
		goto L25
	} else {
		goto L28
	}
L28:
	;
	goto L26
L29:
	;
	return
}
