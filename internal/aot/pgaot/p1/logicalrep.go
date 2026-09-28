package p1

import (
	base "github.com/shibukawa/pgmem/internal/aot/pgaot/base"
	"unsafe"
)

func F_logicalrep_launcher_onexit(m *base.Module, l0 int32, l1 int64) {
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
func F_logicalrep_worker_find(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) int32 {
	mBase := m.M
	_ = mBase
	var v5 int32
	_ = v5
	var v11 int32
	_ = v11
	var v15 int32
	_ = v15
	var v23 int32
	_ = v23
	var v29 int32
	_ = v29
	var v30 int32
	_ = v30
	var v33 int32
	_ = v33
	var v36 int32
	_ = v36
	var v38 int32
	_ = v38
	var v44 int32
	_ = v44
	var v47 int32
	_ = v47
	var v54 int32
	_ = v54
	v5 = int32(0)
	v11 = *(*int32)(unsafe.Add(mBase, _c_F_logicalrep_worker_find[0]))
	if v11 <= v5 {
		v54 = v5
		goto L1
	} else {
		goto L2
	}
L1:
	;
	return v54
L2:
	;
	v15 = *(*int32)(unsafe.Add(mBase, _c_F_logicalrep_worker_find[1]))
	v23 = v5
	goto L3
L3:
	;
	v29 = v15 + int32(16) + v23<<(uint(int32(7))%32)
	v30 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v29)+16)))
	if v30 != int32(1) {
		goto L5
	} else {
		goto L6
	}
L4:
	;
	v54 = int32(0)
	goto L1
L5:
	;
	v47 = v23 + int32(1)
	if v47 != v11 {
		v23 = v47
		goto L3
	} else {
		goto L12
	}
L6:
	;
	v33 = *(*int32)(unsafe.Add(mBase, uint32(v29)))
	if v33 == int32(4) {
		goto L5
	} else {
		goto L7
	}
L7:
	;
	v36 = *(*int32)(unsafe.Add(mBase, uint32(v29)+32))
	if v36 != l1 {
		goto L5
	} else {
		goto L8
	}
L8:
	;
	v38 = *(*int32)(unsafe.Add(mBase, uint32(v29)+36))
	if base.B2i32(v38 != l2)|base.B2i32(l0 != v33) != 0 {
		goto L5
	} else {
		goto L9
	}
L9:
	;
	if l3 == int32(0) {
		v54 = v29
		goto L1
	} else {
		goto L10
	}
L10:
	;
	v44 = *(*int32)(unsafe.Add(mBase, uint32(v29)+20))
	if v44 != 0 {
		v54 = v29
		goto L1
	} else {
		goto L11
	}
L11:
	;
	goto L5
L12:
	;
	goto L4
}
func F_logicalrep_worker_launch(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32, l5 int32, l6 int32, l7 int32) int32 {
	mBase := m.M
	_ = mBase
	var v9 int32
	_ = v9
	var v20 int32
	_ = v20
	var v22 int32
	_ = v22
	var v26 int32
	_ = v26
	var v29 int32
	_ = v29
	var v35 int32
	_ = v35
	var v40 int32
	_ = v40
	var v42 int32
	_ = v42
	var v46 int32
	_ = v46
	var v50 int32
	_ = v50
	var v51 int32
	_ = v51
	var v53 int32
	_ = v53
	var v63 int32
	_ = v63
	var v64 int32
	_ = v64
	var v71 int32
	_ = v71
	var v78 int32
	_ = v78
	var v84 int32
	_ = v84
	var v102 int32
	_ = v102
	var v103 int32
	_ = v103
	var v107 int32
	_ = v107
	var v110 int32
	_ = v110
	var v111 int32
	_ = v111
	var v112 int32
	_ = v112
	var v114 int32
	_ = v114
	var v116 int32
	_ = v116
	var v135 int32
	_ = v135
	var v136 int32
	_ = v136
	var v137 int32
	_ = v137
	var v146 int32
	_ = v146
	var v147 int32
	_ = v147
	var v149 int32
	_ = v149
	var v152 int32
	_ = v152
	var v158 int32
	_ = v158
	var v159 int32
	_ = v159
	var v161 int32
	_ = v161
	var v164 int32
	_ = v164
	var v170 int32
	_ = v170
	var v171 int32
	_ = v171
	var v172 int32
	_ = v172
	var v174 int32
	_ = v174
	var v189 int32
	_ = v189
	var v190 int32
	_ = v190
	var v199 int32
	_ = v199
	var v200 int32
	_ = v200
	var v202 int32
	_ = v202
	var v205 int32
	_ = v205
	var v220 int32
	_ = v220
	var v222 int32
	_ = v222
	var v228 int32
	_ = v228
	var v234 int32
	_ = v234
	var v235 int32
	_ = v235
	var v236 int32
	_ = v236
	var v239 int64
	_ = v239
	var v240 int64
	_ = v240
	var v248 int64
	_ = v248
	var v250 int32
	_ = v250
	var v251 int32
	_ = v251
	var v254 int32
	_ = v254
	var v259 int32
	_ = v259
	var v265 int32
	_ = v265
	var v274 int32
	_ = v274
	var v282 int32
	_ = v282
	var v285 int32
	_ = v285
	var v286 int32
	_ = v286
	var v290 int32
	_ = v290
	var v291 int32
	_ = v291
	var v292 int64
	_ = v292
	var v294 int32
	_ = v294
	var v304 int32
	_ = v304
	var v305 int32
	_ = v305
	var v306 int32
	_ = v306
	var v312 int32
	_ = v312
	var v317 int32
	_ = v317
	var v318 int32
	_ = v318
	var v323 int32
	_ = v323
	var v326 int64
	_ = v326
	var v336 int32
	_ = v336
	var v338 int32
	_ = v338
	var v340 int32
	_ = v340
	var v345 int32
	_ = v345
	var v356 int32
	_ = v356
	var v359 int32
	_ = v359
	var v375 int32
	_ = v375
	var v379 int32
	_ = v379
	var v380 int32
	_ = v380
	var v383 int32
	_ = v383
	var v385 int32
	_ = v385
	var v387 int32
	_ = v387
	var v405 int32
	_ = v405
	var v406 int32
	_ = v406
	var v407 int32
	_ = v407
	var v416 int32
	_ = v416
	var v417 int32
	_ = v417
	var v420 int32
	_ = v420
	var v423 int32
	_ = v423
	var v426 int32
	_ = v426
	var v427 int32
	_ = v427
	var v430 int32
	_ = v430
	var v433 int32
	_ = v433
	var v436 int32
	_ = v436
	var v437 int32
	_ = v437
	var v438 int32
	_ = v438
	var v440 int32
	_ = v440
	var v455 int32
	_ = v455
	var v456 int32
	_ = v456
	var v465 int32
	_ = v465
	var v466 int32
	_ = v466
	var v469 int32
	_ = v469
	var v472 int32
	_ = v472
	var v486 int32
	_ = v486
	var v497 int32
	_ = v497
	var v501 int32
	_ = v501
	var v505 int32
	_ = v505
	var v508 int32
	_ = v508
	var v510 int32
	_ = v510
	var v514 int32
	_ = v514
	var v517 int32
	_ = v517
	var v518 int32
	_ = v518
	var v523 int32
	_ = v523
	var v527 int32
	_ = v527
	var v532 int32
	_ = v532
	var v537 int32
	_ = v537
	var v538 int32
	_ = v538
	var v542 int32
	_ = v542
	var v554 int32
	_ = v554
	var v556 int32
	_ = v556
	var v559 int32
	_ = v559
	var v564 int32
	_ = v564
	var v568 int32
	_ = v568
	var v569 int32
	_ = v569
	var v570 int32
	_ = v570
	var v571 int64
	_ = v571
	var v573 int64
	_ = v573
	var v585 int32
	_ = v585
	var v589 int32
	_ = v589
	var v592 int32
	_ = v592
	var v602 int32
	_ = v602
	var v603 int32
	_ = v603
	var v604 int32
	_ = v604
	var v610 int32
	_ = v610
	var v611 int32
	_ = v611
	var v619 int32
	_ = v619
	var v620 int32
	_ = v620
	var v626 int32
	_ = v626
	var v627 int32
	_ = v627
	var v634 int32
	_ = v634
	var v635 int32
	_ = v635
	var v643 int32
	_ = v643
	var v644 int32
	_ = v644
	var v650 int32
	_ = v650
	var v651 int32
	_ = v651
	var v657 int32
	_ = v657
	var v658 int32
	_ = v658
	var v667 int32
	_ = v667
	var v668 int32
	_ = v668
	var v674 int32
	_ = v674
	var v675 int32
	_ = v675
	var v679 int32
	_ = v679
	var v683 int32
	_ = v683
	var v688 int32
	_ = v688
	var v694 int32
	_ = v694
	var v695 int32
	_ = v695
	var v703 int32
	_ = v703
	var v704 int32
	_ = v704
	var v710 int32
	_ = v710
	var v711 int32
	_ = v711
	var v713 int32
	_ = v713
	var v716 int32
	_ = v716
	var v718 int32
	_ = v718
	var v726 int32
	_ = v726
	var v727 int32
	_ = v727
	var v731 int32
	_ = v731
	var v735 int32
	_ = v735
	var v736 int32
	_ = v736
	var v737 int32
	_ = v737
	var v743 int64
	_ = v743
	var v752 int32
	_ = v752
	var v756 int32
	_ = v756
	var v759 int32
	_ = v759
	var v760 int32
	_ = v760
	var v765 int32
	_ = v765
	var v769 int32
	_ = v769
	var v776 int32
	_ = v776
	var v781 int32
	_ = v781
	var v782 int32
	_ = v782
	var v786 int32
	_ = v786
	var v804 int32
	_ = v804
	var v806 int32
	_ = v806
	var v808 int32
	_ = v808
	var v812 int32
	_ = v812
	var v813 int32
	_ = v813
	var v814 int32
	_ = v814
	var v817 int32
	_ = v817
	var v819 int32
	_ = v819
	var v823 int32
	_ = v823
	var v826 int32
	_ = v826
	var v827 int32
	_ = v827
	var v830 int32
	_ = v830
	var v832 int32
	_ = v832
	var v836 int32
	_ = v836
	var v837 int32
	_ = v837
	var v838 int32
	_ = v838
	var v842 int32
	_ = v842
	var v848 int64
	_ = v848
	var v856 int32
	_ = v856
	var v858 int32
	_ = v858
	var v862 int32
	_ = v862
	var v868 int32
	_ = v868
	var v869 int32
	_ = v869
	var v872 int32
	_ = v872
	var v873 int32
	_ = v873
	var v876 int32
	_ = v876
	var v879 int32
	_ = v879
	var v880 int32
	_ = v880
	var v883 int32
	_ = v883
	var v887 int32
	_ = v887
	var v889 int32
	_ = v889
	var v891 int32
	_ = v891
	var v894 int32
	_ = v894
	var v897 int32
	_ = v897
	var v901 int32
	_ = v901
	var v905 int32
	_ = v905
	var v909 int32
	_ = v909
	var v917 int32
	_ = v917
	var v921 int32
	_ = v921
	var v925 int32
	_ = v925
	var v926 int32
	_ = v926
	var v932 int32
	_ = v932
	var v933 int32
	_ = v933
	var v938 int32
	_ = v938
	var v939 int32
	_ = v939
	var v941 int32
	_ = v941
	var v945 int32
	_ = v945
	var v949 int32
	_ = v949
	var v952 int32
	_ = v952
	var v956 int32
	_ = v956
	var v961 int32
	_ = v961
	var v965 int32
	_ = v965
	v9 = int32(0)
	v20 = m.G0
	v22 = v20 - int32(1600)
	m.G0 = v22
	v26 = F_errstart(m, int32(14), v9)
	mBase = m.M
	v29 = m.ExcPending
	if v29 != 0 {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	return int32(0)
L2:
	;
	if v26 != 0 {
		goto L3
	} else {
		goto L4
	}
L3:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v22)+112)) = l3
	F_errmsg_internal(m, int32(_a_F_logicalrep_worker_launch_0), v22+int32(112))
	mBase = m.M
	v35 = m.ExcPending
	if v35 != 0 {
		goto L1
	} else {
		goto L6
	}
L4:
	;
	goto L5
L5:
	;
	v42 = *(*int32)(unsafe.Add(mBase, _c_F_logicalrep_worker_launch[0]))
	if v42 != 0 {
		goto L9
	} else {
		goto L10
	}
L6:
	;
	F_errfinish(m, int32(_a_F_logicalrep_worker_launch_1), int32(369), int32(_a_F_logicalrep_worker_launch_2))
	mBase = m.M
	v40 = m.ExcPending
	if v40 != 0 {
		goto L1
	} else {
		goto L7
	}
L7:
	;
	goto L5
L8:
	;
	m.G0 = v22 + int32(1600)
	return v965
L9:
	;
	v46 = *(*int32)(unsafe.Add(mBase, _c_F_logicalrep_worker_launch[1]))
	v50 = F_LWLockAcquire(m, v46+int32(_a_F_logicalrep_worker_launch_3), int32(0))
	mBase = m.M
	v51 = m.ExcPending
	if v51 != 0 {
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
	v949 = m.ExcPending
	if v949 != 0 {
		goto L1
	} else {
		goto L173
	}
L12:
	;
	v53 = *(*int32)(unsafe.Add(mBase, _c_F_logicalrep_worker_launch[2]))
	v63 = v9
	v64 = v53
	v71 = v9
	goto L14
L13:
	;
	if base.B2i32(v222 < v359)|base.B2i32(base.Ui32(l0-int32(3)) < base.Ui32(int32(-2))) == int32(0) {
		goto L60
	} else {
		goto L61
	}
L14:
	;
	if v64 <= int32(0) {
		goto L17
	} else {
		goto L18
	}
L15:
	;
	v345 = *(*int32)(unsafe.Add(mBase, _c_F_logicalrep_worker_launch[3]))
	v356 = v340
	v359 = v345
	goto L13
L16:
	;
	v234 = m.G0
	v235 = int32(16)
	v236 = v234 - v235
	m.G0 = v236
	F_gettimeofday(m, v236)
	mBase = m.M
	v239 = *(*int64)(unsafe.Add(mBase, uint32(v236)))
	v240 = int64(*(*int32)(unsafe.Add(mBase, uint32(v236)+8)))
	m.G0 = v236 + v235
	v248 = v240 + v239*int64(1000000) - int64(946684800000000)
	goto L42
L17:
	;
	v220 = v63
	v222 = int32(0)
	v228 = v71
	goto L16
L18:
	;
	goto L19
L19:
	;
	v78 = *(*int32)(unsafe.Add(mBase, _c_F_logicalrep_worker_launch[4]))
	v84 = int32(0)
	goto L20
L20:
	;
	v102 = v78 + int32(16) + v84<<(uint(int32(7))%32)
	v103 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v102)+16)))
	if v103 != int32(1) {
		goto L23
	} else {
		goto L24
	}
L21:
	;
	v112 = int32(0)
	v114 = *(*int32)(unsafe.Add(mBase, _c_F_logicalrep_worker_launch[4]))
	v116 = v114 + int32(16)
	if v64 != int32(1) {
		goto L27
	} else {
		goto L28
	}
L22:
	;
	goto L21
L23:
	;
	v110 = v102
	v111 = v84
	goto L22
L24:
	;
	goto L25
L25:
	;
	v107 = v84 + int32(1)
	if v107 != v64 {
		v84 = v107
		goto L20
	} else {
		goto L26
	}
L26:
	;
	v110 = v63
	v111 = v71
	goto L22
L27:
	;
	v135 = int32(0)
	v136 = v112
	v137 = v112
	goto L30
L28:
	;
	v189 = v112
	v190 = v112
	goto L29
L29:
	;
	v199 = v116 + v190<<(uint(int32(7))%32)
	v200 = *(*int32)(unsafe.Add(mBase, uint32(v199)+32))
	if v200 != l2 {
		v220 = v110
		v222 = v189
		v228 = v111
		goto L16
	} else {
		goto L40
	}
L30:
	;
	v146 = v116 + v137<<(uint(int32(7))%32)
	v147 = *(*int32)(unsafe.Add(mBase, uint32(v146)+32))
	if v147 != l2 {
		v158 = v136
		goto L32
	} else {
		goto L33
	}
L31:
	;
	if v64&int32(1) == int32(0) {
		v220 = v110
		v222 = v170
		v228 = v111
		goto L16
	} else {
		goto L39
	}
L32:
	;
	v159 = *(*int32)(unsafe.Add(mBase, uint32(v146)+160))
	if v159 != l2 {
		v170 = v158
		goto L35
	} else {
		goto L36
	}
L33:
	;
	v149 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v146)+16)))
	if v149 != int32(1) {
		v158 = v136
		goto L32
	} else {
		goto L34
	}
L34:
	;
	v152 = *(*int32)(unsafe.Add(mBase, uint32(v146)))
	v158 = v136 + base.B2i32(base.Ui32(v152-int32(1)) < base.Ui32(int32(2)))
	goto L32
L35:
	;
	v171 = int32(2)
	v172 = v137 + v171
	v174 = v135 + v171
	if v174 != v64&int32(2147483646) {
		v135 = v174
		v136 = v170
		v137 = v172
		goto L30
	} else {
		goto L38
	}
L36:
	;
	v161 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v146)+144)))
	if v161 != int32(1) {
		v170 = v158
		goto L35
	} else {
		goto L37
	}
L37:
	;
	v164 = *(*int32)(unsafe.Add(mBase, uint32(v146)+128))
	v170 = v158 + base.B2i32(base.Ui32(v164-int32(1)) < base.Ui32(int32(2)))
	goto L35
L38:
	;
	goto L31
L39:
	;
	v189 = v170
	v190 = v172
	goto L29
L40:
	;
	v202 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v199)+16)))
	if v202 != int32(1) {
		v220 = v110
		v222 = v189
		v228 = v111
		goto L16
	} else {
		goto L41
	}
L41:
	;
	v205 = *(*int32)(unsafe.Add(mBase, uint32(v199)))
	v220 = v110
	v222 = v189 + base.B2i32(base.Ui32(v205-int32(1)) < base.Ui32(int32(2)))
	v228 = v111
	goto L16
L42:
	;
	v250 = *(*int32)(unsafe.Add(mBase, _c_F_logicalrep_worker_launch[2]))
	v251 = int32(0)
	v254 = *(*int32)(unsafe.Add(mBase, _c_F_logicalrep_worker_launch[3]))
	if base.B2i32(v220 == v251)|base.B2i32(v254 <= v222) == v251 {
		v356 = v250
		v359 = v254
		goto L13
	} else {
		goto L43
	}
L43:
	;
	v259 = int32(0)
	if v250 <= v259 {
		v356 = v250
		v359 = v254
		goto L13
	} else {
		goto L44
	}
L44:
	;
	v265 = int32(0)
	v274 = v259
	goto L45
L45:
	;
	v282 = *(*int32)(unsafe.Add(mBase, _c_F_logicalrep_worker_launch[4]))
	v285 = v282 + v265<<(uint(int32(7))%32)
	v286 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v285)+32)))
	if v286 != int32(1) {
		v336 = v274
		goto L47
	} else {
		goto L48
	}
L46:
	;
	if v336&int32(1) != 0 {
		v63 = v220
		v64 = v340
		v71 = v228
		goto L14
	} else {
		goto L59
	}
L47:
	;
	v338 = v265 + int32(1)
	v340 = *(*int32)(unsafe.Add(mBase, _c_F_logicalrep_worker_launch[2]))
	if v338 < v340 {
		v265 = v338
		v274 = v336
		goto L45
	} else {
		goto L58
	}
L48:
	;
	v290 = v285 + int32(16)
	v291 = *(*int32)(unsafe.Add(mBase, uint32(v290)+20))
	if v291 != 0 {
		v336 = v274
		goto L47
	} else {
		goto L49
	}
L49:
	;
	v292 = *(*int64)(unsafe.Add(mBase, uint32(v290)+8))
	v294 = *(*int32)(unsafe.Add(mBase, _c_F_logicalrep_worker_launch[5]))
	goto L50
L50:
	;
	if base.B2i32(base.I64_extend_i32_s(v294)*int64(1000) <= v248-v292) == int32(0) {
		v336 = v274
		goto L47
	} else {
		goto L51
	}
L51:
	;
	v304 = F_errstart(m, int32(19), int32(0))
	mBase = m.M
	v305 = m.ExcPending
	if v305 != 0 {
		goto L1
	} else {
		goto L52
	}
L52:
	;
	if v304 != 0 {
		goto L53
	} else {
		goto L54
	}
L53:
	;
	v306 = *(*int32)(unsafe.Add(mBase, uint32(v290)+32))
	*(*int32)(unsafe.Add(mBase, uint32(v22)+96)) = v306
	F_errmsg_internal(m, int32(_a_F_logicalrep_worker_launch_4), v22+int32(96))
	mBase = m.M
	v312 = m.ExcPending
	if v312 != 0 {
		goto L1
	} else {
		goto L56
	}
L54:
	;
	goto L55
L55:
	;
	v318 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v290)+16)) = uint8(v318)
	*(*int32)(unsafe.Add(mBase, uint32(v290))) = v318
	v323 = v285 + int32(36)
	*(*int32)(unsafe.Add(mBase, uint32(v323)+16)) = v318
	v326 = int64(0)
	*(*int64)(unsafe.Add(mBase, uint32(v323)+8)) = v326
	*(*int64)(unsafe.Add(mBase, uint32(v323))) = v326
	*(*uint8)(unsafe.Add(mBase, uint32(v290)+68)) = uint8(v318)
	*(*int32)(unsafe.Add(mBase, uint32(v290)+64)) = int32(-1)
	v336 = int32(1)
	goto L47
L56:
	;
	F_errfinish(m, int32(_a_F_logicalrep_worker_launch_1), int32(424), int32(_a_F_logicalrep_worker_launch_2))
	mBase = m.M
	v317 = m.ExcPending
	if v317 != 0 {
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
	v375 = *(*int32)(unsafe.Add(mBase, _c_F_logicalrep_worker_launch[1]))
	F_LWLockRelease(m, v375+int32(_a_F_logicalrep_worker_launch_3))
	mBase = m.M
	v379 = m.ExcPending
	if v379 != 0 {
		goto L1
	} else {
		goto L63
	}
L61:
	;
	goto L62
L62:
	;
	v380 = int32(0)
	if v356 <= v380 {
		v486 = v380
		goto L64
	} else {
		goto L65
	}
L63:
	;
	v965 = int32(0)
	goto L8
L64:
	;
	if l0 != int32(4) {
		goto L81
	} else {
		goto L82
	}
L65:
	;
	v383 = int32(0)
	v385 = *(*int32)(unsafe.Add(mBase, _c_F_logicalrep_worker_launch[4]))
	v387 = v385 + int32(16)
	if v356 != int32(1) {
		goto L66
	} else {
		goto L67
	}
L66:
	;
	v405 = int32(0)
	v406 = v380
	v407 = v383
	goto L69
L67:
	;
	v455 = v380
	v456 = v383
	goto L68
L68:
	;
	v465 = v387 + v456<<(uint(int32(7))%32)
	v466 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v465)+16)))
	if v466 != int32(1) {
		v486 = v455
		goto L64
	} else {
		goto L79
	}
L69:
	;
	v416 = v387 + v407<<(uint(int32(7))%32)
	v417 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v416)+16)))
	if v417 != int32(1) {
		v426 = v406
		goto L71
	} else {
		goto L72
	}
L70:
	;
	if v356&int32(1) == int32(0) {
		v486 = v436
		goto L64
	} else {
		goto L78
	}
L71:
	;
	v427 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v416)+144)))
	if v427 != int32(1) {
		v436 = v426
		goto L74
	} else {
		goto L75
	}
L72:
	;
	v420 = *(*int32)(unsafe.Add(mBase, uint32(v416)))
	if v420 != int32(4) {
		v426 = v406
		goto L71
	} else {
		goto L73
	}
L73:
	;
	v423 = *(*int32)(unsafe.Add(mBase, uint32(v416)+32))
	v426 = v406 + base.B2i32(v423 == l2)
	goto L71
L74:
	;
	v437 = int32(2)
	v438 = v407 + v437
	v440 = v405 + v437
	if v440 != v356&int32(2147483646) {
		v405 = v440
		v406 = v436
		v407 = v438
		goto L69
	} else {
		goto L77
	}
L75:
	;
	v430 = *(*int32)(unsafe.Add(mBase, uint32(v416)+128))
	if v430 != int32(4) {
		v436 = v426
		goto L74
	} else {
		goto L76
	}
L76:
	;
	v433 = *(*int32)(unsafe.Add(mBase, uint32(v416)+160))
	v436 = v426 + base.B2i32(v433 == l2)
	goto L74
L77:
	;
	goto L70
L78:
	;
	v455 = v436
	v456 = v438
	goto L68
L79:
	;
	v469 = *(*int32)(unsafe.Add(mBase, uint32(v465)))
	if v469 != int32(4) {
		v486 = v455
		goto L64
	} else {
		goto L80
	}
L80:
	;
	v472 = *(*int32)(unsafe.Add(mBase, uint32(v465)+32))
	v486 = v455 + base.B2i32(v472 == l2)
	goto L64
L81:
	;
	if v220 == int32(0) {
		goto L85
	} else {
		goto L86
	}
L82:
	;
	v497 = *(*int32)(unsafe.Add(mBase, _c_F_logicalrep_worker_launch[6]))
	if v486 < v497 {
		goto L81
	} else {
		goto L83
	}
L83:
	;
	v501 = *(*int32)(unsafe.Add(mBase, _c_F_logicalrep_worker_launch[1]))
	F_LWLockRelease(m, v501+int32(_a_F_logicalrep_worker_launch_3))
	mBase = m.M
	v505 = m.ExcPending
	if v505 != 0 {
		goto L1
	} else {
		goto L84
	}
L84:
	;
	v965 = int32(0)
	goto L8
L85:
	;
	v508 = int32(0)
	v510 = *(*int32)(unsafe.Add(mBase, _c_F_logicalrep_worker_launch[1]))
	F_LWLockRelease(m, v510+int32(_a_F_logicalrep_worker_launch_3))
	mBase = m.M
	v514 = m.ExcPending
	if v514 != 0 {
		goto L1
	} else {
		goto L88
	}
L86:
	;
	goto L87
L87:
	;
	v538 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(v220)+16)) = uint8(v538)
	*(*int64)(unsafe.Add(mBase, uint32(v220)+8)) = v248
	*(*int32)(unsafe.Add(mBase, uint32(v220))) = l0
	v542 = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v220)+60)) = v542
	*(*int64)(unsafe.Add(mBase, uint32(v220)+48)) = int64(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v220)+40)) = uint8(v542)
	*(*int32)(unsafe.Add(mBase, uint32(v220)+36)) = l5
	*(*int32)(unsafe.Add(mBase, uint32(v220)+32)) = l2
	*(*int32)(unsafe.Add(mBase, uint32(v220)+28)) = l4
	*(*int32)(unsafe.Add(mBase, uint32(v220)+24)) = l1
	*(*int32)(unsafe.Add(mBase, uint32(v220)+20)) = v542
	v554 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v220)+18)))
	v556 = v554 + v538
	*(*uint16)(unsafe.Add(mBase, uint32(v220)+18)) = uint16(v556)
	v559 = *(*int32)(unsafe.Add(mBase, _c_F_logicalrep_worker_launch[7]))
	*(*uint8)(unsafe.Add(mBase, uint32(v220)+68)) = uint8(base.B2i32(l0 == int32(4)))
	if l0 != int32(4) {
		goto L95
	} else {
		goto L96
	}
L88:
	;
	v517 = F_errstart(m, int32(19), int32(0))
	mBase = m.M
	v518 = m.ExcPending
	if v518 != 0 {
		goto L1
	} else {
		goto L89
	}
L89:
	;
	if v517 == int32(0) {
		v965 = v508
		goto L8
	} else {
		goto L90
	}
L90:
	;
	F_errcode(m, int32(_a_F_logicalrep_worker_launch_5))
	mBase = m.M
	v523 = m.ExcPending
	if v523 != 0 {
		goto L1
	} else {
		goto L91
	}
L91:
	;
	F_errmsg(m, int32(_a_F_logicalrep_worker_launch_6), int32(0))
	mBase = m.M
	v527 = m.ExcPending
	if v527 != 0 {
		goto L1
	} else {
		goto L92
	}
L92:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v22))) = int32(_a_F_logicalrep_worker_launch_7)
	F_errhint(m, int32(_a_F_logicalrep_worker_launch_8), v22)
	mBase = m.M
	v532 = m.ExcPending
	if v532 != 0 {
		goto L1
	} else {
		goto L93
	}
L93:
	;
	F_errfinish(m, int32(_a_F_logicalrep_worker_launch_1), int32(470), int32(_a_F_logicalrep_worker_launch_2))
	mBase = m.M
	v537 = m.ExcPending
	if v537 != 0 {
		goto L1
	} else {
		goto L94
	}
L94:
	;
	v965 = v508
	goto L8
L95:
	;
	v564 = int32(-1)
	goto L97
L96:
	;
	v564 = v559
	goto L97
L97:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v220)+64)) = v564
	if l7 != 0 {
		goto L98
	} else {
		goto L99
	}
L98:
	;
	v568 = *(*int32)(unsafe.Add(mBase, _c_F_logicalrep_worker_launch[8]))
	v569 = *(*int32)(unsafe.Add(mBase, uint32(v568)+96))
	v570 = v569
	goto L100
L99:
	;
	v570 = int32(0)
	goto L100
L100:
	;
	v571 = int64(0)
	*(*int64)(unsafe.Add(mBase, uint32(v220)+120)) = v571
	v573 = int64(-9223372036854775807 - 1)
	*(*int64)(unsafe.Add(mBase, uint32(v220)+112)) = v573
	*(*int64)(unsafe.Add(mBase, uint32(v220)+104)) = v571
	*(*int64)(unsafe.Add(mBase, uint32(v220)+96)) = v573
	*(*int64)(unsafe.Add(mBase, uint32(v220)+88)) = v573
	*(*int64)(unsafe.Add(mBase, uint32(v220)+80)) = v571
	*(*int32)(unsafe.Add(mBase, uint32(v220)+72)) = v570
	v585 = *(*int32)(unsafe.Add(mBase, _c_F_logicalrep_worker_launch[1]))
	F_LWLockRelease(m, v585+int32(_a_F_logicalrep_worker_launch_3))
	mBase = m.M
	v589 = m.ExcPending
	if v589 != 0 {
		goto L1
	} else {
		goto L101
	}
L101:
	;
	v592 = int32(0)
	base.MemoryFill(m, v22+int32(120), v592, int32(1472))
	*(*int64)(unsafe.Add(mBase, uint32(v22)+312)) = int64(8589934595)
	v602 = F_pg_snprintf(m, v22+int32(324), int32(1024), int32(_a_F_logicalrep_worker_launch_9), v592)
	mBase = m.M
	v603 = m.ExcPending
	if v603 != 0 {
		goto L1
	} else {
		goto L102
	}
L102:
	;
	v604 = *(*int32)(unsafe.Add(mBase, uint32(v220)))
	switch v604 {
	case 0:
		goto L105
	case 1:
		goto L106
	case 2:
		goto L107
	case 3:
		goto L104
	case 4:
		goto L108
	default:
		goto L103
	}
L103:
	;
	v713 = v220 + int32(20)
	*(*int32)(unsafe.Add(mBase, uint32(v22)+320)) = int32(-1)
	v716 = int32(0)
	v718 = *(*int32)(unsafe.Add(mBase, _c_F_logicalrep_worker_launch[7]))
	*(*int32)(unsafe.Add(mBase, uint32(v22)+1584)) = v718
	*(*int64)(unsafe.Add(mBase, uint32(v22)+1448)) = base.I64_extend_i32_u(v228)
	v726 = F_RegisterDynamicBackgroundWorker(m, v22+int32(120), v22+int32(116))
	mBase = m.M
	v727 = m.ExcPending
	if v727 != 0 {
		goto L1
	} else {
		goto L124
	}
L104:
	;
	v694 = F_pg_snprintf(m, v22+int32(1348), int32(96), int32(_a_F_logicalrep_worker_launch_10), int32(0))
	mBase = m.M
	v695 = m.ExcPending
	if v695 != 0 {
		goto L1
	} else {
		goto L121
	}
L105:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v679 = m.ExcPending
	if v679 != 0 {
		goto L1
	} else {
		goto L118
	}
L106:
	;
	v657 = F_pg_snprintf(m, v22+int32(1348), int32(96), int32(_a_F_logicalrep_worker_launch_11), int32(0))
	mBase = m.M
	v658 = m.ExcPending
	if v658 != 0 {
		goto L1
	} else {
		goto L115
	}
L107:
	;
	v634 = F_pg_snprintf(m, v22+int32(1348), int32(96), int32(_a_F_logicalrep_worker_launch_12), int32(0))
	mBase = m.M
	v635 = m.ExcPending
	if v635 != 0 {
		goto L1
	} else {
		goto L112
	}
L108:
	;
	v610 = F_pg_snprintf(m, v22+int32(1348), int32(96), int32(_a_F_logicalrep_worker_launch_13), int32(0))
	mBase = m.M
	v611 = m.ExcPending
	if v611 != 0 {
		goto L1
	} else {
		goto L109
	}
L109:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v22)+48)) = l2
	v619 = F_pg_snprintf(m, v22+int32(120), int32(96), int32(_a_F_logicalrep_worker_launch_14), v22+int32(48))
	mBase = m.M
	v620 = m.ExcPending
	if v620 != 0 {
		goto L1
	} else {
		goto L110
	}
L110:
	;
	v626 = F_pg_snprintf(m, v22+int32(216), int32(96), int32(_a_F_logicalrep_worker_launch_15), int32(0))
	mBase = m.M
	v627 = m.ExcPending
	if v627 != 0 {
		goto L1
	} else {
		goto L111
	}
L111:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v22)+1456)) = l6
	goto L103
L112:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v22)+64)) = l2
	v643 = F_pg_snprintf(m, v22+int32(120), int32(96), int32(_a_F_logicalrep_worker_launch_16), v22-int32(-64))
	mBase = m.M
	v644 = m.ExcPending
	if v644 != 0 {
		goto L1
	} else {
		goto L113
	}
L113:
	;
	v650 = F_pg_snprintf(m, v22+int32(216), int32(96), int32(_a_F_logicalrep_worker_launch_17), int32(0))
	mBase = m.M
	v651 = m.ExcPending
	if v651 != 0 {
		goto L1
	} else {
		goto L114
	}
L114:
	;
	goto L103
L115:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v22)+84)) = l5
	*(*int32)(unsafe.Add(mBase, uint32(v22)+80)) = l2
	v667 = F_pg_snprintf(m, v22+int32(120), int32(96), int32(_a_F_logicalrep_worker_launch_18), v22+int32(80))
	mBase = m.M
	v668 = m.ExcPending
	if v668 != 0 {
		goto L1
	} else {
		goto L116
	}
L116:
	;
	v674 = F_pg_snprintf(m, v22+int32(216), int32(96), int32(_a_F_logicalrep_worker_launch_19), int32(0))
	mBase = m.M
	v675 = m.ExcPending
	if v675 != 0 {
		goto L1
	} else {
		goto L117
	}
L117:
	;
	goto L103
L118:
	;
	F_errmsg_internal(m, int32(_a_F_logicalrep_worker_launch_20), int32(0))
	mBase = m.M
	v683 = m.ExcPending
	if v683 != 0 {
		goto L1
	} else {
		goto L119
	}
L119:
	;
	F_errfinish(m, int32(_a_F_logicalrep_worker_launch_1), int32(550), int32(_a_F_logicalrep_worker_launch_2))
	mBase = m.M
	v688 = m.ExcPending
	if v688 != 0 {
		goto L1
	} else {
		goto L120
	}
L120:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L121:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v22)+32)) = l2
	v703 = F_pg_snprintf(m, v22+int32(120), int32(96), int32(_a_F_logicalrep_worker_launch_21), v22+int32(32))
	mBase = m.M
	v704 = m.ExcPending
	if v704 != 0 {
		goto L1
	} else {
		goto L122
	}
L122:
	;
	v710 = F_pg_snprintf(m, v22+int32(216), int32(96), int32(_a_F_logicalrep_worker_launch_22), int32(0))
	mBase = m.M
	v711 = m.ExcPending
	if v711 != 0 {
		goto L1
	} else {
		goto L123
	}
L123:
	;
	goto L103
L124:
	;
	if v726 == int32(0) {
		goto L125
	} else {
		goto L126
	}
L125:
	;
	v731 = *(*int32)(unsafe.Add(mBase, _c_F_logicalrep_worker_launch[1]))
	v735 = F_LWLockAcquire(m, v731+int32(_a_F_logicalrep_worker_launch_3), int32(0))
	mBase = m.M
	v736 = m.ExcPending
	if v736 != 0 {
		goto L1
	} else {
		goto L128
	}
L126:
	;
	goto L127
L127:
	;
	v782 = *(*int32)(unsafe.Add(mBase, uint32(v22)+116))
	v786 = int32(0)
	goto L136
L128:
	;
	v737 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v220)+16)) = uint8(v737)
	*(*int32)(unsafe.Add(mBase, uint32(v220))) = v737
	*(*int32)(unsafe.Add(mBase, uint32(v713)+16)) = v737
	v743 = int64(0)
	*(*int64)(unsafe.Add(mBase, uint32(v713)+8)) = v743
	*(*int64)(unsafe.Add(mBase, uint32(v713))) = v743
	*(*uint8)(unsafe.Add(mBase, uint32(v220)+68)) = uint8(v737)
	*(*int32)(unsafe.Add(mBase, uint32(v220)+64)) = int32(-1)
	v752 = *(*int32)(unsafe.Add(mBase, _c_F_logicalrep_worker_launch[1]))
	F_LWLockRelease(m, v752+int32(_a_F_logicalrep_worker_launch_3))
	mBase = m.M
	v756 = m.ExcPending
	if v756 != 0 {
		goto L1
	} else {
		goto L129
	}
L129:
	;
	v759 = F_errstart(m, int32(19), int32(0))
	mBase = m.M
	v760 = m.ExcPending
	if v760 != 0 {
		goto L1
	} else {
		goto L130
	}
L130:
	;
	if v759 == int32(0) {
		v965 = v716
		goto L8
	} else {
		goto L131
	}
L131:
	;
	F_errcode(m, int32(_a_F_logicalrep_worker_launch_5))
	mBase = m.M
	v765 = m.ExcPending
	if v765 != 0 {
		goto L1
	} else {
		goto L132
	}
L132:
	;
	F_errmsg(m, int32(_a_F_logicalrep_worker_launch_23), int32(0))
	mBase = m.M
	v769 = m.ExcPending
	if v769 != 0 {
		goto L1
	} else {
		goto L133
	}
L133:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v22)+16)) = int32(_a_F_logicalrep_worker_launch_24)
	F_errhint(m, int32(_a_F_logicalrep_worker_launch_8), v22+int32(16))
	mBase = m.M
	v776 = m.ExcPending
	if v776 != 0 {
		goto L1
	} else {
		goto L134
	}
L134:
	;
	F_errfinish(m, int32(_a_F_logicalrep_worker_launch_1), int32(568), int32(_a_F_logicalrep_worker_launch_2))
	mBase = m.M
	v781 = m.ExcPending
	if v781 != 0 {
		goto L1
	} else {
		goto L135
	}
L135:
	;
	v965 = v716
	goto L8
L136:
	;
	v804 = *(*int32)(unsafe.Add(mBase, _c_F_logicalrep_worker_launch[9]))
	if v804 != 0 {
		goto L138
	} else {
		goto L139
	}
L138:
	;
	F_ProcessInterrupts(m)
	mBase = m.M
	v806 = m.ExcPending
	if v806 != 0 {
		goto L1
	} else {
		goto L141
	}
L139:
	;
	goto L140
L140:
	;
	v808 = *(*int32)(unsafe.Add(mBase, _c_F_logicalrep_worker_launch[1]))
	v812 = F_LWLockAcquire(m, v808+int32(_a_F_logicalrep_worker_launch_3), int32(1))
	mBase = m.M
	v813 = m.ExcPending
	if v813 != 0 {
		goto L1
	} else {
		goto L142
	}
L141:
	;
	goto L140
L142:
	;
	v814 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v220)+16)))
	if v814 != int32(1) {
		v856 = v814
		goto L144
	} else {
		goto L145
	}
L143:
	;
	v921 = *(*int32)(unsafe.Add(mBase, _c_F_logicalrep_worker_launch[10]))
	v925 = F_WaitLatch(m, v921, int32(41), int32(10), int32(134217734))
	mBase = m.M
	v926 = m.ExcPending
	if v926 != 0 {
		goto L1
	} else {
		goto L168
	}
L144:
	;
	v858 = *(*int32)(unsafe.Add(mBase, _c_F_logicalrep_worker_launch[1]))
	F_LWLockRelease(m, v858+int32(_a_F_logicalrep_worker_launch_3))
	mBase = m.M
	v862 = m.ExcPending
	if v862 != 0 {
		goto L1
	} else {
		goto L152
	}
L145:
	;
	v817 = *(*int32)(unsafe.Add(mBase, uint32(v713)))
	if v817 != 0 {
		v856 = v814
		goto L144
	} else {
		goto L146
	}
L146:
	;
	v819 = *(*int32)(unsafe.Add(mBase, _c_F_logicalrep_worker_launch[1]))
	F_LWLockRelease(m, v819+int32(_a_F_logicalrep_worker_launch_3))
	mBase = m.M
	v823 = m.ExcPending
	if v823 != 0 {
		goto L1
	} else {
		goto L147
	}
L147:
	;
	v826 = F_GetBackgroundWorkerPid(m, v782, v22+int32(1596))
	mBase = m.M
	v827 = m.ExcPending
	if v827 != 0 {
		goto L1
	} else {
		goto L148
	}
L148:
	;
	if v826 != int32(2) {
		goto L143
	} else {
		goto L149
	}
L149:
	;
	v830 = int32(0)
	v832 = *(*int32)(unsafe.Add(mBase, _c_F_logicalrep_worker_launch[1]))
	v836 = F_LWLockAcquire(m, v832+int32(_a_F_logicalrep_worker_launch_3), v830)
	mBase = m.M
	v837 = m.ExcPending
	if v837 != 0 {
		goto L1
	} else {
		goto L150
	}
L150:
	;
	v838 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v220)+18)))
	if v838 != v556&int32(_a_F_logicalrep_worker_launch_25) {
		v856 = v830
		goto L144
	} else {
		goto L151
	}
L151:
	;
	v842 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v220)+16)) = uint8(v842)
	*(*int32)(unsafe.Add(mBase, uint32(v220))) = v842
	*(*int32)(unsafe.Add(mBase, uint32(v713)+16)) = v842
	v848 = int64(0)
	*(*int64)(unsafe.Add(mBase, uint32(v713)+8)) = v848
	*(*int64)(unsafe.Add(mBase, uint32(v713))) = v848
	*(*uint8)(unsafe.Add(mBase, uint32(v220)+68)) = uint8(v842)
	*(*int32)(unsafe.Add(mBase, uint32(v220)+64)) = int32(-1)
	v856 = v830
	goto L144
L152:
	;
	if v786&int32(1) == int32(0) {
		v965 = v856
		goto L8
	} else {
		goto L153
	}
L153:
	;
	v868 = *(*int32)(unsafe.Add(mBase, _c_F_logicalrep_worker_launch[10]))
	v869 = int32(0)
	v872 = base.AtomicRmwOr32(m, v869, int32(_a_F_logicalrep_worker_launch_26), v869)
	v873 = *(*int32)(unsafe.Add(mBase, uint32(v868)))
	if v873 != 0 {
		goto L155
	} else {
		goto L156
	}
L154:
	;
	v965 = v856
	goto L8
L155:
	;
	goto L154
L156:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v868))) = int32(1)
	v876 = int32(0)
	v879 = base.AtomicRmwOr32(m, v876, int32(_a_F_logicalrep_worker_launch_26), v876)
	v880 = *(*int32)(unsafe.Add(mBase, uint32(v868)+4))
	if v880 == v876 {
		goto L155
	} else {
		goto L157
	}
L157:
	;
	v883 = *(*int32)(unsafe.Add(mBase, uint32(v868)+12))
	if v883 == int32(0) {
		goto L155
	} else {
		goto L158
	}
L158:
	;
	v887 = *(*int32)(unsafe.Add(mBase, _c_F_logicalrep_worker_launch[7]))
	if v887 == v883 {
		goto L159
	} else {
		goto L160
	}
L159:
	;
	v889 = m.G0
	v891 = v889 - int32(16)
	m.G0 = v891
	v894 = *(*int32)(unsafe.Add(mBase, _c_F_logicalrep_worker_launch[11]))
	if v894 == int32(0) {
		goto L162
	} else {
		goto L163
	}
L160:
	;
	goto L161
L161:
	;
	v917 = F_pgmem_kill(m, v883, int32(23))
	mBase = m.M
	goto L155
L162:
	;
	m.G0 = v891 + int32(16)
	goto L154
L163:
	;
	v897 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v891)+15)) = uint8(v897)
	goto L164
L164:
	;
	v901 = *(*int32)(unsafe.Add(mBase, _c_F_logicalrep_worker_launch[12]))
	v905 = F_write(m, v901, v891+int32(15), int32(1))
	mBase = m.M
	if int32(0) <= v905 {
		goto L162
	} else {
		goto L166
	}
L165:
	;
	goto L162
L166:
	;
	v909 = *(*int32)(unsafe.Add(mBase, _c_F_logicalrep_worker_launch[13]))
	if v909 == int32(27) {
		goto L164
	} else {
		goto L167
	}
L167:
	;
	goto L165
L168:
	;
	if v925&int32(1) == int32(0) {
		goto L136
	} else {
		goto L169
	}
L169:
	;
	v932 = *(*int32)(unsafe.Add(mBase, _c_F_logicalrep_worker_launch[10]))
	v933 = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v932))) = v933
	v938 = base.AtomicRmwOr32(m, v933, int32(_a_F_logicalrep_worker_launch_26), v933)
	goto L170
L170:
	;
	v939 = int32(1)
	v941 = *(*int32)(unsafe.Add(mBase, _c_F_logicalrep_worker_launch[9]))
	if v941 == int32(0) {
		v786 = v939
		goto L136
	} else {
		goto L171
	}
L171:
	;
	F_ProcessInterrupts(m)
	mBase = m.M
	v945 = m.ExcPending
	if v945 != 0 {
		goto L1
	} else {
		goto L172
	}
L172:
	;
	v786 = v939
	goto L136
L173:
	;
	F_errcode(m, int32(_a_F_logicalrep_worker_launch_5))
	mBase = m.M
	v952 = m.ExcPending
	if v952 != 0 {
		goto L1
	} else {
		goto L174
	}
L174:
	;
	F_errmsg(m, int32(_a_F_logicalrep_worker_launch_27), int32(0))
	mBase = m.M
	v956 = m.ExcPending
	if v956 != 0 {
		goto L1
	} else {
		goto L175
	}
L175:
	;
	F_errfinish(m, int32(_a_F_logicalrep_worker_launch_1), int32(375), int32(_a_F_logicalrep_worker_launch_2))
	mBase = m.M
	v961 = m.ExcPending
	if v961 != 0 {
		goto L1
	} else {
		goto L176
	}
L176:
	;
	base.Wasm_trap_unreachable()
	for {
	}
}
func F_logicalrep_worker_wakeup(m *base.Module, l0 int32) {
	mBase := m.M
	_ = mBase
	var v8 int32
	_ = v8
	var v12 int32
	_ = v12
	var v13 int32
	_ = v13
	var v15 int32
	_ = v15
	var v19 int32
	_ = v19
	var v24 int32
	_ = v24
	var v30 int32
	_ = v30
	var v31 int32
	_ = v31
	var v34 int32
	_ = v34
	var v37 int32
	_ = v37
	var v42 int32
	_ = v42
	var v43 int32
	_ = v43
	var v47 int32
	_ = v47
	var v50 int32
	_ = v50
	var v51 int32
	_ = v51
	var v54 int32
	_ = v54
	var v55 int32
	_ = v55
	var v58 int32
	_ = v58
	var v61 int32
	_ = v61
	var v62 int32
	_ = v62
	var v65 int32
	_ = v65
	var v69 int32
	_ = v69
	var v71 int32
	_ = v71
	var v73 int32
	_ = v73
	var v76 int32
	_ = v76
	var v79 int32
	_ = v79
	var v83 int32
	_ = v83
	var v87 int32
	_ = v87
	var v91 int32
	_ = v91
	var v99 int32
	_ = v99
	var v109 int32
	_ = v109
	var v113 int32
	_ = v113
	v8 = *(*int32)(unsafe.Add(mBase, _c_F_logicalrep_worker_wakeup[0]))
	v12 = F_LWLockAcquire(m, v8+int32(_a_F_logicalrep_worker_wakeup_0), int32(1))
	mBase = m.M
	v13 = m.ExcPending
	if v13 != 0 {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	return
L2:
	;
	v15 = *(*int32)(unsafe.Add(mBase, _c_F_logicalrep_worker_wakeup[1]))
	if v15 <= int32(0) {
		goto L3
	} else {
		goto L4
	}
L3:
	;
	v109 = *(*int32)(unsafe.Add(mBase, _c_F_logicalrep_worker_wakeup[0]))
	F_LWLockRelease(m, v109+int32(_a_F_logicalrep_worker_wakeup_0))
	mBase = m.M
	v113 = m.ExcPending
	if v113 != 0 {
		goto L1
	} else {
		goto L29
	}
L4:
	;
	v19 = *(*int32)(unsafe.Add(mBase, _c_F_logicalrep_worker_wakeup[2]))
	v24 = int32(0)
	goto L5
L5:
	;
	v30 = v19 + int32(16) + v24<<(uint(int32(7))%32)
	v31 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v30)+16)))
	if v31 != int32(1) {
		goto L8
	} else {
		goto L9
	}
L6:
	;
	v50 = v43 + int32(316)
	v51 = int32(0)
	v54 = base.AtomicRmwOr32(m, v51, int32(_a_F_logicalrep_worker_wakeup_1), v51)
	v55 = *(*int32)(unsafe.Add(mBase, uint32(v50)))
	if v55 != 0 {
		goto L16
	} else {
		goto L17
	}
L7:
	;
	goto L6
L8:
	;
	v47 = v24 + int32(1)
	if v47 != v15 {
		v24 = v47
		goto L5
	} else {
		goto L14
	}
L9:
	;
	v34 = *(*int32)(unsafe.Add(mBase, uint32(v30)))
	if v34 == int32(4) {
		goto L8
	} else {
		goto L10
	}
L10:
	;
	v37 = *(*int32)(unsafe.Add(mBase, uint32(v30)+32))
	if base.B2i32(v37 != l0)|base.B2i32(v34 != int32(3)) != 0 {
		goto L8
	} else {
		goto L11
	}
L11:
	;
	v42 = *(*int32)(unsafe.Add(mBase, uint32(v30)+36))
	if v42 != 0 {
		goto L8
	} else {
		goto L12
	}
L12:
	;
	v43 = *(*int32)(unsafe.Add(mBase, uint32(v30)+20))
	if v43 != 0 {
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
	*(*int32)(unsafe.Add(mBase, uint32(v50))) = int32(1)
	v58 = int32(0)
	v61 = base.AtomicRmwOr32(m, v58, int32(_a_F_logicalrep_worker_wakeup_1), v58)
	v62 = *(*int32)(unsafe.Add(mBase, uint32(v50)+4))
	if v62 == v58 {
		goto L16
	} else {
		goto L18
	}
L18:
	;
	v65 = *(*int32)(unsafe.Add(mBase, uint32(v50)+12))
	if v65 == int32(0) {
		goto L16
	} else {
		goto L19
	}
L19:
	;
	v69 = *(*int32)(unsafe.Add(mBase, _c_F_logicalrep_worker_wakeup[3]))
	if v69 == v65 {
		goto L20
	} else {
		goto L21
	}
L20:
	;
	v71 = m.G0
	v73 = v71 - int32(16)
	m.G0 = v73
	v76 = *(*int32)(unsafe.Add(mBase, _c_F_logicalrep_worker_wakeup[4]))
	if v76 == int32(0) {
		goto L23
	} else {
		goto L24
	}
L21:
	;
	goto L22
L22:
	;
	v99 = F_pgmem_kill(m, v65, int32(23))
	mBase = m.M
	goto L16
L23:
	;
	m.G0 = v73 + int32(16)
	goto L15
L24:
	;
	v79 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v73)+15)) = uint8(v79)
	goto L25
L25:
	;
	v83 = *(*int32)(unsafe.Add(mBase, _c_F_logicalrep_worker_wakeup[5]))
	v87 = F_write(m, v83, v73+int32(15), int32(1))
	mBase = m.M
	if int32(0) <= v87 {
		goto L23
	} else {
		goto L27
	}
L26:
	;
	goto L23
L27:
	;
	v91 = *(*int32)(unsafe.Add(mBase, _c_F_logicalrep_worker_wakeup[6]))
	if v91 == int32(27) {
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
