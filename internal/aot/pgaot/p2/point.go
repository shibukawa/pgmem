package p2

import (
	base "github.com/shibukawa/pgmem/internal/aot/pgaot/base"
	"math"
	"sync/atomic"
	"unsafe"
)

func F_CreateCheckPoint(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v2 int32
	_ = v2
	var v16 int32
	_ = v16
	var v18 int32
	_ = v18
	var v21 int32
	_ = v21
	var v23 int32
	_ = v23
	var v31 int32
	_ = v31
	var v32 int32
	_ = v32
	var v33 int32
	_ = v33
	var v36 int32
	_ = v36
	var v38 int32
	_ = v38
	var v40 int32
	_ = v40
	var v42 int32
	_ = v42
	var v43 int32
	_ = v43
	var v51 int32
	_ = v51
	var v52 int32
	_ = v52
	var v53 int32
	_ = v53
	var v56 int64
	_ = v56
	var v57 int64
	_ = v57
	var v70 int32
	_ = v70
	var v71 int32
	_ = v71
	var v73 int32
	_ = v73
	var v74 int32
	_ = v74
	var v75 int32
	_ = v75
	var v77 int32
	_ = v77
	var v79 int32
	_ = v79
	var v84 int32
	_ = v84
	var v86 int32
	_ = v86
	var v90 int32
	_ = v90
	var v91 int32
	_ = v91
	var v93 int32
	_ = v93
	var v97 int32
	_ = v97
	var v99 int32
	_ = v99
	var v101 int32
	_ = v101
	var v105 int32
	_ = v105
	var v112 int64
	_ = v112
	var v115 int32
	_ = v115
	var v120 int32
	_ = v120
	var v121 int32
	_ = v121
	var v122 int32
	_ = v122
	var v124 int64
	_ = v124
	var v125 int32
	_ = v125
	var v128 int32
	_ = v128
	var v130 int32
	_ = v130
	var v131 int64
	_ = v131
	var v133 int32
	_ = v133
	var v135 int32
	_ = v135
	var v141 int32
	_ = v141
	var v142 int32
	_ = v142
	var v148 int32
	_ = v148
	var v153 int32
	_ = v153
	var v157 int32
	_ = v157
	var v158 int32
	_ = v158
	var v163 int32
	_ = v163
	var v164 int32
	_ = v164
	var v166 int32
	_ = v166
	var v169 int32
	_ = v169
	var v170 int32
	_ = v170
	var v172 int32
	_ = v172
	var v173 int32
	_ = v173
	var v176 int32
	_ = v176
	var v177 int32
	_ = v177
	var v180 int32
	_ = v180
	var v182 int64
	_ = v182
	var v184 int64
	_ = v184
	var v185 int64
	_ = v185
	var v187 int64
	_ = v187
	var v193 int64
	_ = v193
	var v194 int64
	_ = v194
	var v195 int64
	_ = v195
	var v205 int64
	_ = v205
	var v207 int32
	_ = v207
	var v212 int64
	_ = v212
	var v227 int64
	_ = v227
	var v230 int32
	_ = v230
	var v235 int32
	_ = v235
	var v237 int32
	_ = v237
	var v239 int32
	_ = v239
	var v244 int32
	_ = v244
	var v246 int32
	_ = v246
	var v248 int32
	_ = v248
	var v252 int32
	_ = v252
	var v257 int32
	_ = v257
	var v258 int32
	_ = v258
	var v261 int32
	_ = v261
	var v263 int32
	_ = v263
	var v267 int32
	_ = v267
	var v270 int32
	_ = v270
	var v273 int32
	_ = v273
	var v275 int32
	_ = v275
	var v280 int32
	_ = v280
	var v283 int64
	_ = v283
	var v284 int32
	_ = v284
	var v286 int64
	_ = v286
	var v292 int32
	_ = v292
	var v295 int32
	_ = v295
	var v300 int32
	_ = v300
	var v302 int32
	_ = v302
	var v303 int64
	_ = v303
	var v305 int32
	_ = v305
	var v309 int32
	_ = v309
	var v314 int32
	_ = v314
	var v319 int32
	_ = v319
	var v325 int32
	_ = v325
	var v327 int32
	_ = v327
	var v328 int32
	_ = v328
	var v331 int32
	_ = v331
	var v332 int32
	_ = v332
	var v333 int32
	_ = v333
	var v336 int32
	_ = v336
	var v340 int32
	_ = v340
	var v341 int32
	_ = v341
	var v343 int32
	_ = v343
	var v344 int64
	_ = v344
	var v346 int32
	_ = v346
	var v348 int32
	_ = v348
	var v351 int32
	_ = v351
	var v355 int32
	_ = v355
	var v357 int32
	_ = v357
	var v361 int32
	_ = v361
	var v362 int32
	_ = v362
	var v364 int32
	_ = v364
	var v365 int32
	_ = v365
	var v367 int32
	_ = v367
	var v370 int32
	_ = v370
	var v374 int32
	_ = v374
	var v376 int32
	_ = v376
	var v380 int32
	_ = v380
	var v381 int32
	_ = v381
	var v383 int32
	_ = v383
	var v384 int32
	_ = v384
	var v388 int32
	_ = v388
	var v392 int32
	_ = v392
	var v396 int32
	_ = v396
	var v397 int32
	_ = v397
	var v398 int32
	_ = v398
	var v409 int32
	_ = v409
	var v413 int32
	_ = v413
	var v414 int32
	_ = v414
	var v415 int32
	_ = v415
	var v416 int32
	_ = v416
	var v417 int32
	_ = v417
	var v420 int32
	_ = v420
	var v421 int64
	_ = v421
	var v423 int32
	_ = v423
	var v426 int32
	_ = v426
	var v427 int32
	_ = v427
	var v430 int32
	_ = v430
	var v434 int32
	_ = v434
	var v435 int32
	_ = v435
	var v437 int32
	_ = v437
	var v438 int32
	_ = v438
	var v444 int32
	_ = v444
	var v445 int32
	_ = v445
	var v446 int32
	_ = v446
	var v465 int32
	_ = v465
	var v466 int32
	_ = v466
	var v467 int32
	_ = v467
	var v473 int32
	_ = v473
	var v476 int32
	_ = v476
	var v478 int32
	_ = v478
	var v479 int32
	_ = v479
	var v496 int32
	_ = v496
	var v497 int64
	_ = v497
	var v499 int32
	_ = v499
	var v503 int32
	_ = v503
	var v504 int32
	_ = v504
	var v505 int32
	_ = v505
	var v524 int32
	_ = v524
	var v525 int32
	_ = v525
	var v526 int32
	_ = v526
	var v532 int32
	_ = v532
	var v535 int32
	_ = v535
	var v537 int32
	_ = v537
	var v538 int32
	_ = v538
	var v555 int32
	_ = v555
	var v557 int32
	_ = v557
	var v560 int64
	_ = v560
	var v561 int32
	_ = v561
	var v562 int32
	_ = v562
	var v564 int32
	_ = v564
	var v569 int32
	_ = v569
	var v574 int32
	_ = v574
	var v575 int32
	_ = v575
	var v578 int32
	_ = v578
	var v579 int64
	_ = v579
	var v580 int32
	_ = v580
	var v582 int32
	_ = v582
	var v585 int64
	_ = v585
	var v587 int64
	_ = v587
	var v590 int32
	_ = v590
	var v591 int64
	_ = v591
	var v593 int32
	_ = v593
	var v597 int32
	_ = v597
	var v598 int32
	_ = v598
	var v600 int32
	_ = v600
	var v604 int64
	_ = v604
	var v614 int64
	_ = v614
	var v617 int32
	_ = v617
	var v620 int64
	_ = v620
	var v622 int32
	_ = v622
	var v625 int32
	_ = v625
	var v627 int32
	_ = v627
	var v629 int32
	_ = v629
	var v633 int32
	_ = v633
	var v634 int32
	_ = v634
	var v636 int32
	_ = v636
	var v641 int32
	_ = v641
	var v645 int32
	_ = v645
	var v649 int32
	_ = v649
	var v650 int32
	_ = v650
	var v652 int32
	_ = v652
	var v653 int32
	_ = v653
	var v655 int32
	_ = v655
	var v659 int32
	_ = v659
	var v663 int32
	_ = v663
	var v664 int32
	_ = v664
	var v669 int32
	_ = v669
	var v670 int32
	_ = v670
	var v673 int32
	_ = v673
	var v674 int32
	_ = v674
	var v677 int32
	_ = v677
	var v680 int32
	_ = v680
	var v681 int32
	_ = v681
	var v684 int32
	_ = v684
	var v688 int32
	_ = v688
	var v690 int32
	_ = v690
	var v692 int32
	_ = v692
	var v695 int32
	_ = v695
	var v698 int32
	_ = v698
	var v702 int32
	_ = v702
	var v706 int32
	_ = v706
	var v710 int32
	_ = v710
	var v718 int32
	_ = v718
	var v723 int32
	_ = v723
	var v725 int32
	_ = v725
	var v728 int32
	_ = v728
	var v729 int32
	_ = v729
	var v734 int32
	_ = v734
	var v739 int32
	_ = v739
	var v748 int32
	_ = v748
	var v751 int32
	_ = v751
	var v752 int32
	_ = v752
	var v753 int32
	_ = v753
	var v754 int32
	_ = v754
	var v756 int32
	_ = v756
	var v759 int32
	_ = v759
	var v760 int32
	_ = v760
	var v765 int32
	_ = v765
	var v766 int32
	_ = v766
	var v767 int32
	_ = v767
	var v771 int32
	_ = v771
	var v776 int32
	_ = v776
	var v777 int32
	_ = v777
	var v781 int32
	_ = v781
	var v785 int32
	_ = v785
	var v790 int32
	_ = v790
	var v791 int32
	_ = v791
	var v798 int32
	_ = v798
	var v800 int32
	_ = v800
	var v803 int32
	_ = v803
	var v804 int32
	_ = v804
	var v822 int32
	_ = v822
	var v839 int32
	_ = v839
	var v841 int32
	_ = v841
	var v842 int32
	_ = v842
	var v844 int32
	_ = v844
	var v845 int32
	_ = v845
	var v848 int32
	_ = v848
	var v852 int32
	_ = v852
	var v867 int32
	_ = v867
	var v868 int32
	_ = v868
	var v872 int32
	_ = v872
	var v874 int32
	_ = v874
	var v876 int32
	_ = v876
	var v879 int32
	_ = v879
	var v895 int32
	_ = v895
	var v896 int32
	_ = v896
	var v897 int32
	_ = v897
	var v907 int32
	_ = v907
	var v919 int64
	_ = v919
	var v924 float64
	_ = v924
	var v926 int32
	_ = v926
	var v928 float64
	_ = v928
	var v935 float64
	_ = v935
	var v940 int64
	_ = v940
	var v941 int64
	_ = v941
	var v944 int32
	_ = v944
	var v946 int32
	_ = v946
	var v948 int64
	_ = v948
	var v949 int32
	_ = v949
	var v951 int32
	_ = v951
	var v952 int32
	_ = v952
	var v954 int64
	_ = v954
	var v956 int64
	_ = v956
	var v957 int64
	_ = v957
	var v960 int32
	_ = v960
	var v961 int64
	_ = v961
	var v962 int64
	_ = v962
	var v966 int64
	_ = v966
	var v967 int32
	_ = v967
	var v969 int32
	_ = v969
	var v971 int32
	_ = v971
	var v972 int32
	_ = v972
	var v976 int64
	_ = v976
	var v978 int32
	_ = v978
	var v990 int64
	_ = v990
	var v993 int32
	_ = v993
	var v998 int32
	_ = v998
	var v999 int32
	_ = v999
	var v1002 int32
	_ = v1002
	var v1003 int32
	_ = v1003
	var v1006 int32
	_ = v1006
	var v1008 int32
	_ = v1008
	var v1015 int32
	_ = v1015
	var v1020 int32
	_ = v1020
	var v1021 int32
	_ = v1021
	var v1023 int32
	_ = v1023
	var v1026 int32
	_ = v1026
	var v1027 int32
	_ = v1027
	var v1029 int32
	_ = v1029
	var v1033 int32
	_ = v1033
	var v1034 int32
	_ = v1034
	var v1038 int32
	_ = v1038
	var v1059 int32
	_ = v1059
	var v1063 int32
	_ = v1063
	var v1068 int32
	_ = v1068
	var v1072 int32
	_ = v1072
	var v1076 int32
	_ = v1076
	var v1081 int32
	_ = v1081
	v2 = int32(0)
	v16 = m.G0
	v18 = v16 - int32(1168)
	m.G0 = v18
	v21 = *(*int32)(unsafe.Add(mBase, _c_F_CreateCheckPoint[0]))
	v23 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_CreateCheckPoint[1])))
	if v23 == v2 {
		goto L4
	} else {
		goto L5
	}
L1:
	;
	F_errstart_cold(m, int32(24), int32(0))
	mBase = m.M
	v1072 = m.ExcPending
	if v1072 != 0 {
		goto L10
	} else {
		goto L223
	}
L2:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v1059 = m.ExcPending
	if v1059 != 0 {
		goto L10
	} else {
		goto L220
	}
L3:
	;
	v42 = int32(0)
	v43 = int32(_a_F_CreateCheckPoint_0)
	base.MemoryFill(m, v43, v42, int32(80))
	v51 = m.G0
	v52 = int32(16)
	v53 = v51 - v52
	m.G0 = v53
	F_gettimeofday(m, v53)
	mBase = m.M
	v56 = *(*int64)(unsafe.Add(mBase, uint32(v53)))
	v57 = int64(*(*int32)(unsafe.Add(mBase, uint32(v53)+8)))
	m.G0 = v53 + v52
	goto L9
L4:
	;
	v40 = base.B2i32(l0&int32(2) == int32(0))
	goto L3
L5:
	;
	goto L6
L6:
	;
	v31 = *(*int32)(unsafe.Add(mBase, uint32(v21)+308))
	v32 = int32(2)
	v33 = base.B2i32(v31 != v32)
	*(*uint8)(unsafe.Add(mBase, _c_F_CreateCheckPoint[1])) = uint8(v33)
	v36 = l0 & v32
	v38 = base.B2i32(v36 == int32(0))
	if v36 != 0 {
		v40 = v38
		goto L3
	} else {
		goto L7
	}
L7:
	;
	if v31 != v32 {
		goto L2
	} else {
		goto L8
	}
L8:
	;
	v40 = v38
	goto L3
L9:
	;
	*(*int64)(unsafe.Add(mBase, _c_F_CreateCheckPoint[2])) = v57 + v56*int64(1000000) - int64(946684800000000)
	F_AbsorbSyncRequests(m)
	mBase = m.M
	v70 = m.ExcPending
	if v70 != 0 {
		goto L10
	} else {
		goto L11
	}
L10:
	;
	return int32(0)
L11:
	;
	v71 = int32(_a_F_CreateCheckPoint_1)
	v73 = int32(*(*uint16)(unsafe.Add(mBase, _c_F_CreateCheckPoint[3])))
	v74 = int32(1)
	v75 = v73 + v74
	*(*uint16)(unsafe.Add(mBase, _c_F_CreateCheckPoint[3])) = uint16(v75)
	v77 = int32(_a_F_CreateCheckPoint_2)
	v79 = *(*int32)(unsafe.Add(mBase, _c_F_CreateCheckPoint[4]))
	*(*int32)(unsafe.Add(mBase, _c_F_CreateCheckPoint[4])) = v79 + v74
	v84 = l0 & int32(3)
	if v84 != 0 {
		goto L12
	} else {
		goto L13
	}
L12:
	;
	v86 = *(*int32)(unsafe.Add(mBase, _c_F_CreateCheckPoint[5]))
	v90 = F_LWLockAcquire(m, v86+int32(1152), int32(0))
	mBase = m.M
	v91 = m.ExcPending
	if v91 != 0 {
		goto L10
	} else {
		goto L15
	}
L13:
	;
	goto L14
L14:
	;
	base.MemoryFill(m, v18+int32(32), int32(0), int32(96))
	v112 = F_time(m)
	mBase = m.M
	*(*int64)(unsafe.Add(mBase, uint32(v18)+104)) = v112
	if v84 != 0 {
		v122 = v42
		goto L18
	} else {
		goto L19
	}
L15:
	;
	v93 = *(*int32)(unsafe.Add(mBase, _c_F_CreateCheckPoint[6]))
	*(*int32)(unsafe.Add(mBase, uint32(v93)+16)) = int32(3)
	v97 = *(*int32)(unsafe.Add(mBase, _c_F_CreateCheckPoint[7]))
	F_update_controlfile(m, v97, v93)
	mBase = m.M
	v99 = m.ExcPending
	if v99 != 0 {
		goto L10
	} else {
		goto L16
	}
L16:
	;
	v101 = *(*int32)(unsafe.Add(mBase, _c_F_CreateCheckPoint[5]))
	F_LWLockRelease(m, v101+int32(1152))
	mBase = m.M
	v105 = m.ExcPending
	if v105 != 0 {
		goto L10
	} else {
		goto L17
	}
L17:
	;
	goto L14
L18:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v18)+120)) = v122
	v124 = F_GetLastImportantRecPtr(m)
	mBase = m.M
	v125 = m.ExcPending
	if v125 != 0 {
		goto L10
	} else {
		goto L22
	}
L19:
	;
	v115 = *(*int32)(unsafe.Add(mBase, _c_F_CreateCheckPoint[8]))
	if v115 <= int32(0) {
		v122 = v42
		goto L18
	} else {
		goto L20
	}
L20:
	;
	v120 = F_GetOldestActiveTransactionId(m, int32(0), int32(1))
	mBase = m.M
	v121 = m.ExcPending
	if v121 != 0 {
		goto L10
	} else {
		goto L21
	}
L21:
	;
	v122 = v120
	goto L18
L22:
	;
	if l0&int32(11) != 0 {
		goto L24
	} else {
		goto L25
	}
L23:
	;
	m.G0 = v18 + int32(1168)
	return v1038
L24:
	;
	if v40 == int32(0) {
		goto L32
	} else {
		goto L33
	}
L25:
	;
	v128 = int32(0)
	v130 = *(*int32)(unsafe.Add(mBase, _c_F_CreateCheckPoint[6]))
	v131 = *(*int64)(unsafe.Add(mBase, uint32(v130)+32))
	if v124 != v131 {
		goto L24
	} else {
		goto L26
	}
L26:
	;
	v133 = int32(_a_F_CreateCheckPoint_2)
	v135 = *(*int32)(unsafe.Add(mBase, _c_F_CreateCheckPoint[4]))
	*(*int32)(unsafe.Add(mBase, _c_F_CreateCheckPoint[4])) = v135 - int32(1)
	v141 = F_errstart(m, int32(14), int32(0))
	mBase = m.M
	v142 = m.ExcPending
	if v142 != 0 {
		goto L10
	} else {
		goto L27
	}
L27:
	;
	if v141 == int32(0) {
		v1038 = v128
		goto L23
	} else {
		goto L28
	}
L28:
	;
	F_errmsg_internal(m, int32(_a_F_CreateCheckPoint_3), int32(0))
	mBase = m.M
	v148 = m.ExcPending
	if v148 != 0 {
		goto L10
	} else {
		goto L29
	}
L29:
	;
	F_errfinish(m, int32(_a_F_CreateCheckPoint_4), int32(_a_F_CreateCheckPoint_5), int32(_a_F_CreateCheckPoint_6))
	mBase = m.M
	v153 = m.ExcPending
	if v153 != 0 {
		goto L10
	} else {
		goto L30
	}
L30:
	;
	v1038 = v128
	goto L23
L31:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v18)+44)) = v172
	F_WALInsertLockAcquireExclusive(m)
	mBase = m.M
	v176 = m.ExcPending
	if v176 != 0 {
		goto L10
	} else {
		goto L35
	}
L32:
	;
	v157 = int32(_a_F_CreateCheckPoint_7)
	v158 = *(*int32)(unsafe.Add(mBase, _c_F_CreateCheckPoint[9]))
	*(*int32)(unsafe.Add(mBase, _c_F_CreateCheckPoint[9])) = int32(1)
	v163 = *(*int32)(unsafe.Add(mBase, _c_F_CreateCheckPoint[0]))
	v164 = *(*int32)(unsafe.Add(mBase, uint32(v163)+300))
	*(*int32)(unsafe.Add(mBase, uint32(v18)+40)) = v164
	v166 = *(*int32)(unsafe.Add(mBase, uint32(v163)+304))
	v172 = v166
	v173 = v158
	goto L31
L33:
	;
	goto L34
L34:
	;
	v169 = *(*int32)(unsafe.Add(mBase, _c_F_CreateCheckPoint[0]))
	v170 = *(*int32)(unsafe.Add(mBase, uint32(v169)+300))
	*(*int32)(unsafe.Add(mBase, uint32(v18)+40)) = v170
	v172 = v170
	v173 = int32(0)
	goto L31
L35:
	;
	v177 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v21)+160)))
	*(*uint8)(unsafe.Add(mBase, uint32(v18)+48)) = uint8(v177)
	v180 = *(*int32)(unsafe.Add(mBase, _c_F_CreateCheckPoint[8]))
	*(*int32)(unsafe.Add(mBase, uint32(v18)+52)) = v180
	if v84 != 0 {
		goto L37
	} else {
		goto L38
	}
L36:
	;
	v292 = *(*int32)(unsafe.Add(mBase, _c_F_CreateCheckPoint[0]))
	v295 = base.AtomicRmwXchg32(m, v292, int32(440), int32(1))
	if v295 != 0 {
		goto L62
	} else {
		goto L63
	}
L37:
	;
	v182 = *(*int64)(unsafe.Add(mBase, uint32(v21)+8))
	v184 = int64(*(*int32)(unsafe.Add(mBase, _c_F_CreateCheckPoint[10])))
	v185 = base.I64_div_u_s(v182, v184)
	v187 = v182 - v185*v184
	if base.Ui64(v187) <= base.Ui64(int64(8151)) {
		goto L41
	} else {
		goto L42
	}
L38:
	;
	goto L39
L39:
	;
	F_WALInsertLockRelease(m)
	mBase = m.M
	v237 = m.ExcPending
	if v237 != 0 {
		goto L10
	} else {
		goto L50
	}
L40:
	;
	v207 = *(*int32)(unsafe.Add(mBase, _c_F_CreateCheckPoint[11]))
	v212 = v185*base.I64_extend_i32_s(v207) + v205&int64(4294967295)
	if v212&int64(8191) != int64(0) {
		v227 = v212
		goto L44
	} else {
		goto L45
	}
L41:
	;
	v205 = v187 + int64(40)
	goto L40
L42:
	;
	goto L43
L43:
	;
	v193 = v187 - int64(8152)
	v194 = int64(8168)
	v195 = base.I64_div_u_s(v193, v194)
	v205 = v193 - v195*v194 + v195<<(uint(int64(13))%64) + int64(8216)
	goto L40
L44:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v18)+32)) = v227
	v230 = *(*int32)(unsafe.Add(mBase, _c_F_CreateCheckPoint[0]))
	*(*int64)(unsafe.Add(mBase, uint32(v230)+152)) = v227
	*(*int64)(unsafe.Add(mBase, _c_F_CreateCheckPoint[12])) = v227
	F_WALInsertLockRelease(m)
	mBase = m.M
	v235 = m.ExcPending
	if v235 != 0 {
		goto L10
	} else {
		goto L49
	}
L45:
	;
	if v212&base.I64_extend_i32_s(v207-int32(1)) == int64(0) {
		goto L46
	} else {
		goto L47
	}
L46:
	;
	v227 = v212 | int64(40)
	goto L44
L47:
	;
	goto L48
L48:
	;
	v227 = v212 | int64(24)
	goto L44
L49:
	;
	goto L36
L50:
	;
	v239 = *(*int32)(unsafe.Add(mBase, _c_F_CreateCheckPoint[13]))
	if v239 == int32(-1) {
		goto L51
	} else {
		goto L52
	}
L51:
	;
	v244 = *(*int32)(unsafe.Add(mBase, _c_F_CreateCheckPoint[14]))
	v246 = base.I32_rem_s(v244, int32(8))
	*(*int32)(unsafe.Add(mBase, _c_F_CreateCheckPoint[13])) = v246
	v248 = v246
	goto L53
L52:
	;
	v248 = v239
	goto L53
L53:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_CreateCheckPoint[15])) = v248
	v252 = *(*int32)(unsafe.Add(mBase, _c_F_CreateCheckPoint[16]))
	v257 = F_LWLockAcquire(m, v252+v248<<(uint(int32(7))%32), int32(0))
	mBase = m.M
	v258 = m.ExcPending
	if v258 != 0 {
		goto L10
	} else {
		goto L54
	}
L54:
	;
	if v257 == int32(0) {
		goto L55
	} else {
		goto L56
	}
L55:
	;
	v261 = int32(_a_F_CreateCheckPoint_8)
	v263 = *(*int32)(unsafe.Add(mBase, _c_F_CreateCheckPoint[13]))
	v267 = base.I32_rem_s(v263+int32(1), int32(8))
	*(*int32)(unsafe.Add(mBase, _c_F_CreateCheckPoint[13])) = v267
	goto L57
L56:
	;
	goto L57
L57:
	;
	v270 = *(*int32)(unsafe.Add(mBase, _c_F_CreateCheckPoint[8]))
	*(*int32)(unsafe.Add(mBase, uint32(v18)+128)) = v270
	F_WALInsertLockRelease(m)
	mBase = m.M
	v273 = m.ExcPending
	if v273 != 0 {
		goto L10
	} else {
		goto L58
	}
L58:
	;
	F_XLogBeginInsert(m)
	mBase = m.M
	v275 = m.ExcPending
	if v275 != 0 {
		goto L10
	} else {
		goto L59
	}
L59:
	;
	F_XLogRegisterData(m, v18+int32(128), int32(4))
	mBase = m.M
	v280 = m.ExcPending
	if v280 != 0 {
		goto L10
	} else {
		goto L60
	}
L60:
	;
	v283 = F_XLogInsert(m, int32(0), int32(224))
	mBase = m.M
	v284 = m.ExcPending
	if v284 != 0 {
		goto L10
	} else {
		goto L61
	}
L61:
	;
	v286 = *(*int64)(unsafe.Add(mBase, _c_F_CreateCheckPoint[12]))
	*(*int64)(unsafe.Add(mBase, uint32(v18)+32)) = v286
	goto L36
L62:
	;
	F_s_lock(m, v292+int32(440), int32(_a_F_CreateCheckPoint_9))
	mBase = m.M
	v300 = m.ExcPending
	if v300 != 0 {
		goto L10
	} else {
		goto L65
	}
L63:
	;
	goto L64
L64:
	;
	v302 = *(*int32)(unsafe.Add(mBase, _c_F_CreateCheckPoint[0]))
	v303 = *(*int64)(unsafe.Add(mBase, uint32(v18)+32))
	*(*int64)(unsafe.Add(mBase, uint32(v302)+200)) = v303
	v305 = int32(0)
	atomic.StoreUint32((*uint32)(unsafe.Add(mBase, uint32(v302)+440)), uint32(v305))
	v309 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_CreateCheckPoint[17])))
	if v309 == int32(1) {
		goto L66
	} else {
		goto L67
	}
L65:
	;
	goto L64
L66:
	;
	F_LogCheckpointStart(m, l0, int32(0))
	mBase = m.M
	v314 = m.ExcPending
	if v314 != 0 {
		goto L10
	} else {
		goto L69
	}
L67:
	;
	goto L68
L68:
	;
	if v84 != 0 {
		goto L70
	} else {
		goto L71
	}
L69:
	;
	goto L68
L70:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v18)+8)) = int32(_a_F_CreateCheckPoint_10)
	if v40 != 0 {
		goto L73
	} else {
		goto L74
	}
L71:
	;
	goto L72
L72:
	;
	v336 = *(*int32)(unsafe.Add(mBase, _c_F_CreateCheckPoint[5]))
	v340 = F_LWLockAcquire(m, v336+int32(384), int32(1))
	mBase = m.M
	v341 = m.ExcPending
	if v341 != 0 {
		goto L10
	} else {
		goto L80
	}
L73:
	;
	v319 = int32(_a_F_CreateCheckPoint_11)
	goto L75
L74:
	;
	v319 = int32(_a_F_CreateCheckPoint_12)
	goto L75
L75:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v18))) = v319
	if l0&int32(1) != 0 {
		goto L76
	} else {
		goto L77
	}
L76:
	;
	v325 = int32(_a_F_CreateCheckPoint_13)
	goto L78
L77:
	;
	v325 = int32(_a_F_CreateCheckPoint_11)
	goto L78
L78:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v18)+4)) = v325
	v327 = int32(128)
	v328 = v18 + v327
	v331 = F_pg_snprintf(m, v328, v327, int32(_a_F_CreateCheckPoint_14), v18)
	mBase = m.M
	v332 = m.ExcPending
	if v332 != 0 {
		goto L10
	} else {
		goto L79
	}
L79:
	;
	v333 = F_strlen(m, v328)
	mBase = m.M
	goto L72
L80:
	;
	v343 = *(*int32)(unsafe.Add(mBase, _c_F_CreateCheckPoint[18]))
	v344 = *(*int64)(unsafe.Add(mBase, uint32(v343)+8))
	*(*int64)(unsafe.Add(mBase, uint32(v18)+64)) = v344
	v346 = *(*int32)(unsafe.Add(mBase, uint32(v343)+16))
	*(*int32)(unsafe.Add(mBase, uint32(v18)+88)) = v346
	v348 = *(*int32)(unsafe.Add(mBase, uint32(v343)+36))
	*(*int32)(unsafe.Add(mBase, uint32(v18)+92)) = v348
	v351 = *(*int32)(unsafe.Add(mBase, _c_F_CreateCheckPoint[5]))
	F_LWLockRelease(m, v351+int32(384))
	mBase = m.M
	v355 = m.ExcPending
	if v355 != 0 {
		goto L10
	} else {
		goto L81
	}
L81:
	;
	v357 = *(*int32)(unsafe.Add(mBase, _c_F_CreateCheckPoint[5]))
	v361 = F_LWLockAcquire(m, v357+int32(_a_F_CreateCheckPoint_15), int32(1))
	mBase = m.M
	v362 = m.ExcPending
	if v362 != 0 {
		goto L10
	} else {
		goto L82
	}
L82:
	;
	v364 = *(*int32)(unsafe.Add(mBase, _c_F_CreateCheckPoint[18]))
	v365 = *(*int32)(unsafe.Add(mBase, uint32(v364)+40))
	*(*int32)(unsafe.Add(mBase, uint32(v18)+112)) = v365
	v367 = *(*int32)(unsafe.Add(mBase, uint32(v364)+44))
	*(*int32)(unsafe.Add(mBase, uint32(v18)+116)) = v367
	v370 = *(*int32)(unsafe.Add(mBase, _c_F_CreateCheckPoint[5]))
	F_LWLockRelease(m, v370+int32(_a_F_CreateCheckPoint_15))
	mBase = m.M
	v374 = m.ExcPending
	if v374 != 0 {
		goto L10
	} else {
		goto L83
	}
L83:
	;
	v376 = *(*int32)(unsafe.Add(mBase, _c_F_CreateCheckPoint[5]))
	v380 = F_LWLockAcquire(m, v376+int32(256), int32(1))
	mBase = m.M
	v381 = m.ExcPending
	if v381 != 0 {
		goto L10
	} else {
		goto L84
	}
L84:
	;
	v383 = *(*int32)(unsafe.Add(mBase, _c_F_CreateCheckPoint[18]))
	v384 = *(*int32)(unsafe.Add(mBase, uint32(v383)))
	*(*int32)(unsafe.Add(mBase, uint32(v18)+72)) = v384
	if v84 == int32(0) {
		goto L85
	} else {
		goto L86
	}
L85:
	;
	v388 = *(*int32)(unsafe.Add(mBase, uint32(v383)+4))
	*(*int32)(unsafe.Add(mBase, uint32(v18)+72)) = v388 + v384
	goto L87
L86:
	;
	goto L87
L87:
	;
	v392 = *(*int32)(unsafe.Add(mBase, _c_F_CreateCheckPoint[5]))
	F_LWLockRelease(m, v392+int32(256))
	mBase = m.M
	v396 = m.ExcPending
	if v396 != 0 {
		goto L10
	} else {
		goto L88
	}
L88:
	;
	v397 = F_IsLogicalDecodingEnabled(m)
	mBase = m.M
	v398 = m.ExcPending
	if v398 != 0 {
		goto L10
	} else {
		goto L89
	}
L89:
	;
	*(*uint8)(unsafe.Add(mBase, uint32(v18)+56)) = uint8(v397)
	v409 = *(*int32)(unsafe.Add(mBase, _c_F_CreateCheckPoint[5]))
	v413 = F_LWLockAcquire(m, v409+int32(1664), int32(1))
	mBase = m.M
	v414 = m.ExcPending
	if v414 != 0 {
		goto L10
	} else {
		goto L90
	}
L90:
	;
	v415 = int32(_a_F_CreateCheckPoint_16)
	v416 = *(*int32)(unsafe.Add(mBase, _c_F_CreateCheckPoint[19]))
	v417 = *(*int32)(unsafe.Add(mBase, uint32(v416)))
	*(*int32)(unsafe.Add(mBase, uint32(v18+int32(76)))) = v417
	v420 = *(*int32)(unsafe.Add(mBase, _c_F_CreateCheckPoint[19]))
	v421 = *(*int64)(unsafe.Add(mBase, uint32(v420)+8))
	*(*int64)(unsafe.Add(mBase, uint32(v18+int32(80)))) = v421
	v423 = *(*int32)(unsafe.Add(mBase, uint32(v420)+20))
	*(*int32)(unsafe.Add(mBase, uint32(v18+int32(96)))) = v423
	v426 = *(*int32)(unsafe.Add(mBase, _c_F_CreateCheckPoint[19]))
	v427 = *(*int32)(unsafe.Add(mBase, uint32(v426)+24))
	*(*int32)(unsafe.Add(mBase, uint32(v18+int32(100)))) = v427
	v430 = *(*int32)(unsafe.Add(mBase, _c_F_CreateCheckPoint[5]))
	F_LWLockRelease(m, v430+int32(1664))
	mBase = m.M
	v434 = m.ExcPending
	if v434 != 0 {
		goto L10
	} else {
		goto L91
	}
L91:
	;
	v435 = int32(_a_F_CreateCheckPoint_2)
	v437 = *(*int32)(unsafe.Add(mBase, _c_F_CreateCheckPoint[4]))
	v438 = int32(1)
	*(*int32)(unsafe.Add(mBase, _c_F_CreateCheckPoint[4])) = v437 - v438
	v444 = F_GetVirtualXIDsDelayingChkpt(m, v18+int32(20), v438)
	mBase = m.M
	v445 = m.ExcPending
	if v445 != 0 {
		goto L10
	} else {
		goto L92
	}
L92:
	;
	v446 = *(*int32)(unsafe.Add(mBase, uint32(v18)+20))
	if int32(0) < v446 {
		goto L93
	} else {
		goto L94
	}
L93:
	;
	goto L96
L94:
	;
	goto L95
L95:
	;
	F_pfree(m, v444)
	mBase = m.M
	v496 = m.ExcPending
	if v496 != 0 {
		goto L10
	} else {
		goto L101
	}
L96:
	;
	F_AbsorbSyncRequests(m)
	mBase = m.M
	v465 = m.ExcPending
	if v465 != 0 {
		goto L10
	} else {
		goto L98
	}
L97:
	;
	goto L95
L98:
	;
	v466 = int32(_a_F_CreateCheckPoint_17)
	v467 = *(*int32)(unsafe.Add(mBase, _c_F_CreateCheckPoint[20]))
	*(*int32)(unsafe.Add(mBase, uint32(v467))) = int32(134217738)
	F_pg_usleep(m, int32(_a_F_CreateCheckPoint_18))
	mBase = m.M
	v473 = *(*int32)(unsafe.Add(mBase, _c_F_CreateCheckPoint[20]))
	*(*int32)(unsafe.Add(mBase, uint32(v473))) = int32(0)
	v476 = *(*int32)(unsafe.Add(mBase, uint32(v18)+20))
	v478 = F_HaveVirtualXIDsDelayingChkpt(m, v444, v476, int32(1))
	mBase = m.M
	v479 = m.ExcPending
	if v479 != 0 {
		goto L10
	} else {
		goto L99
	}
L99:
	;
	if v478 != 0 {
		goto L96
	} else {
		goto L100
	}
L100:
	;
	goto L97
L101:
	;
	v497 = *(*int64)(unsafe.Add(mBase, uint32(v18)+32))
	F_CheckPointGuts(m, v497, l0)
	mBase = m.M
	v499 = m.ExcPending
	if v499 != 0 {
		goto L10
	} else {
		goto L102
	}
L102:
	;
	v503 = F_GetVirtualXIDsDelayingChkpt(m, v18+int32(20), int32(2))
	mBase = m.M
	v504 = m.ExcPending
	if v504 != 0 {
		goto L10
	} else {
		goto L103
	}
L103:
	;
	v505 = *(*int32)(unsafe.Add(mBase, uint32(v18)+20))
	if int32(0) < v505 {
		goto L104
	} else {
		goto L105
	}
L104:
	;
	goto L107
L105:
	;
	goto L106
L106:
	;
	F_pfree(m, v503)
	mBase = m.M
	v555 = m.ExcPending
	if v555 != 0 {
		goto L10
	} else {
		goto L112
	}
L107:
	;
	F_AbsorbSyncRequests(m)
	mBase = m.M
	v524 = m.ExcPending
	if v524 != 0 {
		goto L10
	} else {
		goto L109
	}
L108:
	;
	goto L106
L109:
	;
	v525 = int32(_a_F_CreateCheckPoint_17)
	v526 = *(*int32)(unsafe.Add(mBase, _c_F_CreateCheckPoint[20]))
	*(*int32)(unsafe.Add(mBase, uint32(v526))) = int32(134217737)
	F_pg_usleep(m, int32(_a_F_CreateCheckPoint_18))
	mBase = m.M
	v532 = *(*int32)(unsafe.Add(mBase, _c_F_CreateCheckPoint[20]))
	*(*int32)(unsafe.Add(mBase, uint32(v532))) = int32(0)
	v535 = *(*int32)(unsafe.Add(mBase, uint32(v18)+20))
	v537 = F_HaveVirtualXIDsDelayingChkpt(m, v503, v535, int32(2))
	mBase = m.M
	v538 = m.ExcPending
	if v538 != 0 {
		goto L10
	} else {
		goto L110
	}
L110:
	;
	if v537 != 0 {
		goto L107
	} else {
		goto L111
	}
L111:
	;
	goto L108
L112:
	;
	if v84 != 0 {
		goto L113
	} else {
		goto L114
	}
L113:
	;
	v562 = int32(_a_F_CreateCheckPoint_2)
	v564 = *(*int32)(unsafe.Add(mBase, _c_F_CreateCheckPoint[4]))
	*(*int32)(unsafe.Add(mBase, _c_F_CreateCheckPoint[4])) = v564 + int32(1)
	F_XLogBeginInsert(m)
	mBase = m.M
	v569 = m.ExcPending
	if v569 != 0 {
		goto L10
	} else {
		goto L117
	}
L114:
	;
	v557 = *(*int32)(unsafe.Add(mBase, _c_F_CreateCheckPoint[8]))
	if v557 <= int32(0) {
		goto L113
	} else {
		goto L115
	}
L115:
	;
	v560 = F_LogStandbySnapshot(m)
	mBase = m.M
	v561 = m.ExcPending
	if v561 != 0 {
		goto L10
	} else {
		goto L116
	}
L116:
	;
	goto L113
L117:
	;
	F_XLogRegisterData(m, v18+int32(32), int32(96))
	mBase = m.M
	v574 = m.ExcPending
	if v574 != 0 {
		goto L10
	} else {
		goto L118
	}
L118:
	;
	v575 = int32(0)
	if v84 != 0 {
		goto L119
	} else {
		goto L120
	}
L119:
	;
	v578 = v575
	goto L121
L120:
	;
	v578 = int32(16)
	goto L121
L121:
	;
	v579 = F_XLogInsert(m, v575, v578)
	mBase = m.M
	v580 = m.ExcPending
	if v580 != 0 {
		goto L10
	} else {
		goto L122
	}
L122:
	;
	F_XLogFlush(m, v579)
	mBase = m.M
	v582 = m.ExcPending
	if v582 != 0 {
		goto L10
	} else {
		goto L123
	}
L123:
	;
	if v84 != 0 {
		goto L124
	} else {
		goto L125
	}
L124:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_CreateCheckPoint[9])) = v173
	v585 = *(*int64)(unsafe.Add(mBase, uint32(v18)+32))
	v587 = *(*int64)(unsafe.Add(mBase, _c_F_CreateCheckPoint[21]))
	if v585 != v587 {
		goto L1
	} else {
		goto L127
	}
L125:
	;
	goto L126
L126:
	;
	v590 = *(*int32)(unsafe.Add(mBase, _c_F_CreateCheckPoint[6]))
	v591 = *(*int64)(unsafe.Add(mBase, uint32(v590)+40))
	v593 = *(*int32)(unsafe.Add(mBase, _c_F_CreateCheckPoint[5]))
	v597 = F_LWLockAcquire(m, v593+int32(1152), int32(0))
	mBase = m.M
	v598 = m.ExcPending
	if v598 != 0 {
		goto L10
	} else {
		goto L128
	}
L127:
	;
	goto L126
L128:
	;
	v600 = *(*int32)(unsafe.Add(mBase, _c_F_CreateCheckPoint[6]))
	if v84 != 0 {
		goto L129
	} else {
		goto L130
	}
L129:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v600)+16)) = int32(1)
	goto L131
L130:
	;
	goto L131
L131:
	;
	v604 = *(*int64)(unsafe.Add(mBase, _c_F_CreateCheckPoint[21]))
	*(*int64)(unsafe.Add(mBase, uint32(v600)+32)) = v604
	base.MemoryCopy(m, v600+int32(40), v18+int32(32), int32(96))
	*(*int32)(unsafe.Add(mBase, uint32(v600)+152)) = int32(0)
	v614 = int64(0)
	*(*int64)(unsafe.Add(mBase, uint32(v600)+144)) = v614
	v617 = *(*int32)(unsafe.Add(mBase, _c_F_CreateCheckPoint[0]))
	v620 = base.AtomicRmwOr64(m, v617, int32(232), v614)
	v622 = *(*int32)(unsafe.Add(mBase, _c_F_CreateCheckPoint[6]))
	*(*int64)(unsafe.Add(mBase, uint32(v622)+136)) = v620
	v625 = *(*int32)(unsafe.Add(mBase, _c_F_CreateCheckPoint[7]))
	F_update_controlfile(m, v625, v622)
	mBase = m.M
	v627 = m.ExcPending
	if v627 != 0 {
		goto L10
	} else {
		goto L132
	}
L132:
	;
	v629 = *(*int32)(unsafe.Add(mBase, _c_F_CreateCheckPoint[5]))
	F_LWLockRelease(m, v629+int32(1152))
	mBase = m.M
	v633 = m.ExcPending
	if v633 != 0 {
		goto L10
	} else {
		goto L133
	}
L133:
	;
	v634 = int32(_a_F_CreateCheckPoint_2)
	v636 = *(*int32)(unsafe.Add(mBase, _c_F_CreateCheckPoint[4]))
	*(*int32)(unsafe.Add(mBase, _c_F_CreateCheckPoint[4])) = v636 - int32(1)
	v641 = *(*int32)(unsafe.Add(mBase, _c_F_CreateCheckPoint[22]))
	if v641 == int32(0) {
		goto L134
	} else {
		goto L135
	}
L134:
	;
	v723 = m.G0
	v725 = v723 - int32(1040)
	m.G0 = v725
	v728 = *(*int32)(unsafe.Add(mBase, _c_F_CreateCheckPoint[23]))
	if v728 != 0 {
		goto L155
	} else {
		goto L156
	}
L135:
	;
	v645 = *(*int32)(unsafe.Add(mBase, _c_F_CreateCheckPoint[5]))
	v649 = F_LWLockAcquire(m, v645+int32(_a_F_CreateCheckPoint_19), int32(1))
	mBase = m.M
	v650 = m.ExcPending
	if v650 != 0 {
		goto L10
	} else {
		goto L136
	}
L136:
	;
	v652 = *(*int32)(unsafe.Add(mBase, _c_F_CreateCheckPoint[22]))
	v653 = *(*int32)(unsafe.Add(mBase, uint32(v652)+20))
	v655 = *(*int32)(unsafe.Add(mBase, _c_F_CreateCheckPoint[5]))
	F_LWLockRelease(m, v655+int32(_a_F_CreateCheckPoint_19))
	mBase = m.M
	v659 = m.ExcPending
	if v659 != 0 {
		goto L10
	} else {
		goto L137
	}
L137:
	;
	if v653 == int32(-1) {
		goto L134
	} else {
		goto L138
	}
L138:
	;
	v663 = *(*int32)(unsafe.Add(mBase, _c_F_CreateCheckPoint[24]))
	v664 = *(*int32)(unsafe.Add(mBase, uint32(v663)))
	v669 = v664 + v653*int32(768) + int32(316)
	v670 = int32(0)
	v673 = base.AtomicRmwOr32(m, v670, int32(_a_F_CreateCheckPoint_20), v670)
	v674 = *(*int32)(unsafe.Add(mBase, uint32(v669)))
	if v674 != 0 {
		goto L140
	} else {
		goto L141
	}
L139:
	;
	goto L134
L140:
	;
	goto L139
L141:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v669))) = int32(1)
	v677 = int32(0)
	v680 = base.AtomicRmwOr32(m, v677, int32(_a_F_CreateCheckPoint_20), v677)
	v681 = *(*int32)(unsafe.Add(mBase, uint32(v669)+4))
	if v681 == v677 {
		goto L140
	} else {
		goto L142
	}
L142:
	;
	v684 = *(*int32)(unsafe.Add(mBase, uint32(v669)+12))
	if v684 == int32(0) {
		goto L140
	} else {
		goto L143
	}
L143:
	;
	v688 = *(*int32)(unsafe.Add(mBase, _c_F_CreateCheckPoint[25]))
	if v688 == v684 {
		goto L144
	} else {
		goto L145
	}
L144:
	;
	v690 = m.G0
	v692 = v690 - int32(16)
	m.G0 = v692
	v695 = *(*int32)(unsafe.Add(mBase, _c_F_CreateCheckPoint[26]))
	if v695 == int32(0) {
		goto L147
	} else {
		goto L148
	}
L145:
	;
	goto L146
L146:
	;
	v718 = F_pgmem_kill(m, v684, int32(23))
	mBase = m.M
	goto L140
L147:
	;
	m.G0 = v692 + int32(16)
	goto L139
L148:
	;
	v698 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v692)+15)) = uint8(v698)
	goto L149
L149:
	;
	v702 = *(*int32)(unsafe.Add(mBase, _c_F_CreateCheckPoint[27]))
	v706 = F_write(m, v702, v692+int32(15), int32(1))
	mBase = m.M
	if int32(0) <= v706 {
		goto L147
	} else {
		goto L151
	}
L150:
	;
	goto L147
L151:
	;
	v710 = *(*int32)(unsafe.Add(mBase, _c_F_CreateCheckPoint[28]))
	if v710 == int32(27) {
		goto L149
	} else {
		goto L152
	}
L152:
	;
	goto L150
L153:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_CreateCheckPoint[23])) = v907
	m.G0 = v725 + int32(1040)
	v919 = *(*int64)(unsafe.Add(mBase, _c_F_CreateCheckPoint[12]))
	if v591 != int64(0) {
		goto L189
	} else {
		goto L190
	}
L154:
	;
	v842 = int32(0)
	v844 = *(*int32)(unsafe.Add(mBase, _c_F_CreateCheckPoint[23]))
	v845 = *(*int32)(unsafe.Add(mBase, uint32(v844)+12))
	v848 = (v751 - v845) >> (uint(int32(2)) % 32)
	if v842 < v848 {
		goto L181
	} else {
		goto L182
	}
L155:
	;
	v729 = *(*int32)(unsafe.Add(mBase, uint32(v728)+4))
	if int32(0) < v729 {
		goto L158
	} else {
		goto L159
	}
L156:
	;
	v839 = int32(0)
	goto L157
L157:
	;
	F_list_free_deep(m, v839)
	mBase = m.M
	v841 = m.ExcPending
	if v841 != 0 {
		goto L10
	} else {
		goto L180
	}
L158:
	;
	v734 = int32(0)
	v739 = int32(10)
	goto L161
L159:
	;
	goto L160
L160:
	;
	v822 = *(*int32)(unsafe.Add(mBase, _c_F_CreateCheckPoint[23]))
	v839 = v822
	goto L157
L161:
	;
	v748 = *(*int32)(unsafe.Add(mBase, uint32(v728)+12))
	v751 = v748 + v734<<(uint(int32(2))%32)
	v752 = *(*int32)(unsafe.Add(mBase, uint32(v751)))
	v753 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v752)+26)))
	if v753 != 0 {
		v800 = v739
		goto L163
	} else {
		goto L164
	}
L162:
	;
	goto L160
L163:
	;
	v803 = v734 + int32(1)
	v804 = *(*int32)(unsafe.Add(mBase, uint32(v728)+4))
	if v803 < v804 {
		v734 = v803
		v739 = v800
		goto L161
	} else {
		goto L179
	}
L164:
	;
	v754 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v752)+24)))
	v756 = int32(*(*uint16)(unsafe.Add(mBase, _c_F_CreateCheckPoint[3])))
	if v754 == v756 {
		goto L154
	} else {
		goto L165
	}
L165:
	;
	v759 = v725 + int32(16)
	v760 = int32(*(*int16)(unsafe.Add(mBase, uint32(v752))))
	v765 = *(*int32)(unsafe.Add(mBase, uint32(v760*int32(12))+uint32(_c_F_CreateCheckPoint[29])))
	v766 = m.T0[v765].(func(*base.Module, int32, int32) int32)(m, v752, v759)
	mBase = m.M
	v767 = m.ExcPending
	if v767 != 0 {
		goto L10
	} else {
		goto L167
	}
L166:
	;
	v791 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(v752)+26)) = uint8(v791)
	if v791 < v739 {
		goto L175
	} else {
		goto L176
	}
L167:
	;
	if int32(0) <= v766 {
		goto L166
	} else {
		goto L168
	}
L168:
	;
	v771 = *(*int32)(unsafe.Add(mBase, _c_F_CreateCheckPoint[28]))
	if v771 == int32(44) {
		goto L166
	} else {
		goto L169
	}
L169:
	;
	v776 = F_errstart(m, int32(19), int32(0))
	mBase = m.M
	v777 = m.ExcPending
	if v777 != 0 {
		goto L10
	} else {
		goto L170
	}
L170:
	;
	if v776 == int32(0) {
		goto L166
	} else {
		goto L171
	}
L171:
	;
	F_errcode_for_file_access(m)
	mBase = m.M
	v781 = m.ExcPending
	if v781 != 0 {
		goto L10
	} else {
		goto L172
	}
L172:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v725))) = v759
	F_errmsg(m, int32(_a_F_CreateCheckPoint_21), v725)
	mBase = m.M
	v785 = m.ExcPending
	if v785 != 0 {
		goto L10
	} else {
		goto L173
	}
L173:
	;
	F_errfinish(m, int32(_a_F_CreateCheckPoint_22), int32(244), int32(_a_F_CreateCheckPoint_23))
	mBase = m.M
	v790 = m.ExcPending
	if v790 != 0 {
		goto L10
	} else {
		goto L174
	}
L174:
	;
	goto L166
L175:
	;
	v800 = v739 - int32(1)
	goto L163
L176:
	;
	goto L177
L177:
	;
	F_AbsorbSyncRequests(m)
	mBase = m.M
	v798 = m.ExcPending
	if v798 != 0 {
		goto L10
	} else {
		goto L178
	}
L178:
	;
	v800 = int32(10)
	goto L163
L179:
	;
	goto L162
L180:
	;
	v907 = v2
	goto L153
L181:
	;
	v852 = v842
	goto L184
L182:
	;
	v895 = v844
	goto L183
L183:
	;
	v896 = F_list_delete_first_n(m, v895, v848)
	mBase = m.M
	v897 = m.ExcPending
	if v897 != 0 {
		goto L10
	} else {
		goto L188
	}
L184:
	;
	v867 = *(*int32)(unsafe.Add(mBase, _c_F_CreateCheckPoint[23]))
	v868 = *(*int32)(unsafe.Add(mBase, uint32(v867)+12))
	v872 = *(*int32)(unsafe.Add(mBase, uint32(v868+v852<<(uint(int32(2))%32))))
	F_pfree(m, v872)
	mBase = m.M
	v874 = m.ExcPending
	if v874 != 0 {
		goto L10
	} else {
		goto L186
	}
L185:
	;
	v879 = *(*int32)(unsafe.Add(mBase, _c_F_CreateCheckPoint[23]))
	v895 = v879
	goto L183
L186:
	;
	v876 = v852 + int32(1)
	if v876 != v848 {
		v852 = v876
		goto L184
	} else {
		goto L187
	}
L187:
	;
	goto L185
L188:
	;
	v907 = v896
	goto L153
L189:
	;
	v924 = base.F64_convert_i64_u(v919 - v591)
	*(*float64)(unsafe.Add(mBase, _c_F_CreateCheckPoint[30])) = v924
	v926 = int32(_a_F_CreateCheckPoint_24)
	v928 = *(*float64)(unsafe.Add(mBase, _c_F_CreateCheckPoint[31]))
	if base.F64_gt(v924, v928) != 0 {
		goto L192
	} else {
		goto L193
	}
L190:
	;
	goto L191
L191:
	;
	v940 = int64(*(*int32)(unsafe.Add(mBase, _c_F_CreateCheckPoint[11])))
	v941 = base.I64_div_u_s(v919, v940)
	*(*int64)(unsafe.Add(mBase, uint32(v18)+24)) = v941
	v944 = v18 + int32(24)
	F_KeepLogSeg(m, v579, v944)
	mBase = m.M
	v946 = m.ExcPending
	if v946 != 0 {
		goto L10
	} else {
		goto L195
	}
L192:
	;
	v935 = v924
	goto L194
L193:
	;
	v935 = base.F64_add(base.F64_mul(v928, float64(0.9)), base.F64_mul(v924, float64(0.1)))
	goto L194
L194:
	;
	*(*float64)(unsafe.Add(mBase, _c_F_CreateCheckPoint[31])) = v935
	goto L191
L195:
	;
	v948 = *(*int64)(unsafe.Add(mBase, uint32(v18)+24))
	v949 = int32(0)
	v951 = F_InvalidateObsoleteReplicationSlots(m, int32(9), v948, v949, v949)
	mBase = m.M
	v952 = m.ExcPending
	if v952 != 0 {
		goto L10
	} else {
		goto L196
	}
L196:
	;
	if v951 != 0 {
		goto L197
	} else {
		goto L198
	}
L197:
	;
	v954 = *(*int64)(unsafe.Add(mBase, _c_F_CreateCheckPoint[12]))
	v956 = int64(*(*int32)(unsafe.Add(mBase, _c_F_CreateCheckPoint[11])))
	v957 = base.I64_div_u_s(v954, v956)
	*(*int64)(unsafe.Add(mBase, uint32(v18)+24)) = v957
	F_KeepLogSeg(m, v579, v944)
	mBase = m.M
	v960 = m.ExcPending
	if v960 != 0 {
		goto L10
	} else {
		goto L200
	}
L198:
	;
	v962 = v948
	goto L199
L199:
	;
	v966 = *(*int64)(unsafe.Add(mBase, _c_F_CreateCheckPoint[12]))
	v967 = *(*int32)(unsafe.Add(mBase, uint32(v18)+40))
	F_RemoveOldXlogFiles(m, v962-int64(1), v966, v579, v967)
	mBase = m.M
	v969 = m.ExcPending
	if v969 != 0 {
		goto L10
	} else {
		goto L201
	}
L200:
	;
	v961 = *(*int64)(unsafe.Add(mBase, uint32(v18)+24))
	v962 = v961
	goto L199
L201:
	;
	if v84 != 0 {
		goto L202
	} else {
		goto L203
	}
L202:
	;
	v1015 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_CreateCheckPoint[1])))
	if v1015 == int32(1) {
		goto L212
	} else {
		goto L213
	}
L203:
	;
	v971 = *(*int32)(unsafe.Add(mBase, _c_F_CreateCheckPoint[0]))
	v972 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v971)+312)))
	if v972 != int32(1) {
		goto L202
	} else {
		goto L204
	}
L204:
	;
	v976 = v579 - int64(1)
	v978 = *(*int32)(unsafe.Add(mBase, _c_F_CreateCheckPoint[11]))
	if base.Ui64(v976&base.I64_extend_i32_s(v978-int32(1))) < base.Ui64(base.I64_extend_i32_u(base.I32_trunc_sat_f64_u(base.F64_mul(base.F64_convert_i32_s(v978), float64(0.75))))) {
		goto L202
	} else {
		goto L205
	}
L205:
	;
	v990 = base.I64_div_u_s(v976, base.I64_extend_i32_s(v978))
	v993 = *(*int32)(unsafe.Add(mBase, uint32(v18)+40))
	v998 = F_XLogFileInitInternal(m, v990+int64(1), v993, v18+int32(1167), v18+int32(128))
	mBase = m.M
	v999 = m.ExcPending
	if v999 != 0 {
		goto L10
	} else {
		goto L206
	}
L206:
	;
	if int32(0) <= v998 {
		goto L207
	} else {
		goto L208
	}
L207:
	;
	v1002 = F_close(m, v998)
	mBase = m.M
	goto L209
L208:
	;
	goto L209
L209:
	;
	v1003 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v18)+1167)))
	if v1003 != int32(1) {
		goto L202
	} else {
		goto L210
	}
L210:
	;
	v1006 = int32(_a_F_CreateCheckPoint_25)
	v1008 = *(*int32)(unsafe.Add(mBase, _c_F_CreateCheckPoint[32]))
	*(*int32)(unsafe.Add(mBase, _c_F_CreateCheckPoint[32])) = v1008 + int32(1)
	goto L202
L211:
	;
	F_LogCheckpointEnd(m, int32(0), l0)
	mBase = m.M
	v1033 = m.ExcPending
	if v1033 != 0 {
		goto L10
	} else {
		goto L218
	}
L212:
	;
	v1020 = *(*int32)(unsafe.Add(mBase, _c_F_CreateCheckPoint[0]))
	v1021 = *(*int32)(unsafe.Add(mBase, uint32(v1020)+308))
	v1023 = base.B2i32(v1021 != int32(2))
	*(*uint8)(unsafe.Add(mBase, _c_F_CreateCheckPoint[1])) = uint8(v1023)
	if v1021 != int32(2) {
		goto L211
	} else {
		goto L215
	}
L213:
	;
	goto L214
L214:
	;
	v1026 = F_GetOldestTransactionIdConsideredRunning(m)
	mBase = m.M
	v1027 = m.ExcPending
	if v1027 != 0 {
		goto L10
	} else {
		goto L216
	}
L215:
	;
	goto L214
L216:
	;
	F_TruncateSUBTRANS(m, v1026)
	mBase = m.M
	v1029 = m.ExcPending
	if v1029 != 0 {
		goto L10
	} else {
		goto L217
	}
L217:
	;
	goto L211
L218:
	;
	v1034 = int32(1)
	if v84 == int32(0) {
		v1038 = v1034
		goto L23
	} else {
		goto L219
	}
L219:
	;
	v1038 = v1034
	goto L23
L220:
	;
	F_errmsg_internal(m, int32(_a_F_CreateCheckPoint_26), int32(0))
	mBase = m.M
	v1063 = m.ExcPending
	if v1063 != 0 {
		goto L10
	} else {
		goto L221
	}
L221:
	;
	F_errfinish(m, int32(_a_F_CreateCheckPoint_4), int32(_a_F_CreateCheckPoint_27), int32(_a_F_CreateCheckPoint_6))
	mBase = m.M
	v1068 = m.ExcPending
	if v1068 != 0 {
		goto L10
	} else {
		goto L222
	}
L222:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L223:
	;
	F_errmsg(m, int32(_a_F_CreateCheckPoint_28), int32(0))
	mBase = m.M
	v1076 = m.ExcPending
	if v1076 != 0 {
		goto L10
	} else {
		goto L224
	}
L224:
	;
	F_errfinish(m, int32(_a_F_CreateCheckPoint_4), int32(_a_F_CreateCheckPoint_29), int32(_a_F_CreateCheckPoint_6))
	mBase = m.M
	v1081 = m.ExcPending
	if v1081 != 0 {
		goto L10
	} else {
		goto L225
	}
L225:
	;
	base.Wasm_trap_unreachable()
	for {
	}
}
func F_construct_point(m *base.Module, l0 int32) int64 {
	mBase := m.M
	_ = mBase
	var v4 float64
	_ = v4
	var v5 float64
	_ = v5
	var v7 int32
	_ = v7
	var v10 int32
	_ = v10
	v4 = *(*float64)(unsafe.Add(mBase, uint32(l0)+24))
	v5 = *(*float64)(unsafe.Add(mBase, uint32(l0)+40))
	v7 = F_palloc(m, int32(16))
	mBase = m.M
	v10 = m.ExcPending
	if v10 != 0 {
		return int64(0)
	} else {
		*(*float64)(unsafe.Add(mBase, uint32(v7)+8)) = v5
		*(*float64)(unsafe.Add(mBase, uint32(v7))) = v4
		return base.I64_extend_i32_u(v7)
	}
}
func F_point_above(m *base.Module, l0 int32) int64 {
	mBase := m.M
	_ = mBase
	var v2 int32
	_ = v2
	var v3 float64
	_ = v3
	var v4 int32
	_ = v4
	var v5 float64
	_ = v5
	v2 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
	v3 = *(*float64)(unsafe.Add(mBase, uint32(v2)+8))
	v4 = *(*int32)(unsafe.Add(mBase, uint32(l0)+40))
	v5 = *(*float64)(unsafe.Add(mBase, uint32(v4)+8))
	return base.I64_extend_i32_u(base.F64_gt(v3, base.F64_add(v5, float64(1e-06))))
}
func F_point_distance(m *base.Module, l0 int32) int64 {
	mBase := m.M
	_ = mBase
	var v6 int32
	_ = v6
	var v7 float64
	_ = v7
	var v8 int32
	_ = v8
	var v9 float64
	_ = v9
	var v10 float64
	_ = v10
	var v12 float64
	_ = v12
	var v23 float64
	_ = v23
	var v26 int32
	_ = v26
	var v27 float64
	_ = v27
	var v28 float64
	_ = v28
	var v29 float64
	_ = v29
	var v30 float64
	_ = v30
	var v32 float64
	_ = v32
	var v43 float64
	_ = v43
	var v44 int32
	_ = v44
	var v45 float64
	_ = v45
	var v54 int32
	_ = v54
	var v56 int32
	_ = v56
	var v58 float64
	_ = v58
	var v59 float64
	_ = v59
	var v62 int32
	_ = v62
	var v63 float64
	_ = v63
	var v64 int64
	_ = v64
	var v66 int64
	_ = v66
	var v69 float64
	_ = v69
	var v72 int64
	_ = v72
	var v74 int64
	_ = v74
	var v85 float64
	_ = v85
	var v93 float64
	_ = v93
	var v98 float64
	_ = v98
	var v99 float64
	_ = v99
	var v100 float64
	_ = v100
	var v109 float64
	_ = v109
	var v110 float64
	_ = v110
	var v112 float64
	_ = v112
	var v114 float64
	_ = v114
	var v121 float64
	_ = v121
	v6 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
	v7 = *(*float64)(unsafe.Add(mBase, uint32(v6)))
	v8 = *(*int32)(unsafe.Add(mBase, uint32(l0)+40))
	v9 = *(*float64)(unsafe.Add(mBase, uint32(v8)))
	v10 = base.F64_sub(v7, v9)
	v12 = math.Float64frombits(uint64(0x7ff0000000000000))
	if base.F64_ne(base.F64_abs(v10), v12)|base.F64_eq(base.F64_abs(v7), v12)|base.F64_eq(base.F64_abs(v9), v12) != 0 {
		v27 = v10
		v28 = *(*float64)(unsafe.Add(mBase, uint32(v6)+8))
		v29 = *(*float64)(unsafe.Add(mBase, uint32(v8)+8))
		v30 = base.F64_sub(v28, v29)
		v32 = math.Float64frombits(uint64(0x7ff0000000000000))
		if base.F64_ne(base.F64_abs(v30), v32)|base.F64_eq(base.F64_abs(v28), v32)|base.F64_eq(base.F64_abs(v29), v32) != 0 {
			v45 = v30
			v54 = m.G0
			v56 = v54 - int32(32)
			m.G0 = v56
			v58 = base.F64_abs(v27)
			v59 = base.F64_abs(v45)
			v62 = base.B2i32(base.Ui64(base.I64_reinterpret_f64(v58)) < base.Ui64(base.I64_reinterpret_f64(v59)))
			if base.Ui64(base.I64_reinterpret_f64(v58)) < base.Ui64(base.I64_reinterpret_f64(v59)) {
				v63 = v58
			} else {
				v63 = v59
			}
			v64 = base.I64_reinterpret_f64(v63)
			v66 = int64(base.Ui64(v64) >> (uint(int64(52)) % 64))
			if v66 == int64(2047) {
				v121 = v63
			} else {
				if base.Ui64(base.I64_reinterpret_f64(v58)) < base.Ui64(base.I64_reinterpret_f64(v59)) {
					v69 = v59
				} else {
					v69 = v58
				}
				if v64 == int64(0) {
					v121 = v69
				} else {
					v72 = base.I64_reinterpret_f64(v69)
					v74 = int64(base.Ui64(v72) >> (uint(int64(52)) % 64))
					if v74 == int64(2047) {
						v121 = v69
					} else {
						if int32(65) <= base.I32_wrap_i64(v74)-base.I32_wrap_i64(v66) {
							v121 = base.F64_add(v58, v59)
						} else {
							if base.Ui64(int64(6908521828386340864)) <= base.Ui64(v72) {
								v85 = float64(1.90109156629516e-211)
								v98 = base.F64_mul(v69, v85)
								v99 = base.F64_mul(v63, v85)
								v100 = float64(5.260135901548374e+210)
							} else {
								if base.Ui64(int64(2580562586483294207)) < base.Ui64(v64) {
									v98 = v69
									v99 = v63
									v100 = float64(1)
								} else {
									v93 = float64(5.260135901548374e+210)
									v98 = base.F64_mul(v69, v93)
									v99 = base.F64_mul(v63, v93)
									v100 = float64(1.90109156629516e-211)
								}
							}
							F_sq(m, v56+int32(24), v56+int32(16), v98)
							mBase = m.M
							F_sq(m, v56+int32(8), v56, v99)
							mBase = m.M
							v109 = *(*float64)(unsafe.Add(mBase, uint32(v56)))
							v110 = *(*float64)(unsafe.Add(mBase, uint32(v56)+16))
							v112 = *(*float64)(unsafe.Add(mBase, uint32(v56)+8))
							v114 = *(*float64)(unsafe.Add(mBase, uint32(v56)+24))
							v121 = base.F64_mul(v100, base.F64_sqrt(base.F64_add(base.F64_add(base.F64_add(v109, v110), v112), v114)))
						}
					}
				}
			}
			m.G0 = v56 + int32(32)
			return base.I64_reinterpret_f64(v121)
		} else {
			v43 = F_float_overflow_error_ext(m, int32(0))
			mBase = m.M
			v44 = m.ExcPending
			if v44 != 0 {
				return int64(0)
			} else {
				v45 = v43
				v54 = m.G0
				v56 = v54 - int32(32)
				m.G0 = v56
				v58 = base.F64_abs(v27)
				v59 = base.F64_abs(v45)
				v62 = base.B2i32(base.Ui64(base.I64_reinterpret_f64(v58)) < base.Ui64(base.I64_reinterpret_f64(v59)))
				if base.Ui64(base.I64_reinterpret_f64(v58)) < base.Ui64(base.I64_reinterpret_f64(v59)) {
					v63 = v58
				} else {
					v63 = v59
				}
				v64 = base.I64_reinterpret_f64(v63)
				v66 = int64(base.Ui64(v64) >> (uint(int64(52)) % 64))
				if v66 == int64(2047) {
					v121 = v63
				} else {
					if base.Ui64(base.I64_reinterpret_f64(v58)) < base.Ui64(base.I64_reinterpret_f64(v59)) {
						v69 = v59
					} else {
						v69 = v58
					}
					if v64 == int64(0) {
						v121 = v69
					} else {
						v72 = base.I64_reinterpret_f64(v69)
						v74 = int64(base.Ui64(v72) >> (uint(int64(52)) % 64))
						if v74 == int64(2047) {
							v121 = v69
						} else {
							if int32(65) <= base.I32_wrap_i64(v74)-base.I32_wrap_i64(v66) {
								v121 = base.F64_add(v58, v59)
							} else {
								if base.Ui64(int64(6908521828386340864)) <= base.Ui64(v72) {
									v85 = float64(1.90109156629516e-211)
									v98 = base.F64_mul(v69, v85)
									v99 = base.F64_mul(v63, v85)
									v100 = float64(5.260135901548374e+210)
								} else {
									if base.Ui64(int64(2580562586483294207)) < base.Ui64(v64) {
										v98 = v69
										v99 = v63
										v100 = float64(1)
									} else {
										v93 = float64(5.260135901548374e+210)
										v98 = base.F64_mul(v69, v93)
										v99 = base.F64_mul(v63, v93)
										v100 = float64(1.90109156629516e-211)
									}
								}
								F_sq(m, v56+int32(24), v56+int32(16), v98)
								mBase = m.M
								F_sq(m, v56+int32(8), v56, v99)
								mBase = m.M
								v109 = *(*float64)(unsafe.Add(mBase, uint32(v56)))
								v110 = *(*float64)(unsafe.Add(mBase, uint32(v56)+16))
								v112 = *(*float64)(unsafe.Add(mBase, uint32(v56)+8))
								v114 = *(*float64)(unsafe.Add(mBase, uint32(v56)+24))
								v121 = base.F64_mul(v100, base.F64_sqrt(base.F64_add(base.F64_add(base.F64_add(v109, v110), v112), v114)))
							}
						}
					}
				}
				m.G0 = v56 + int32(32)
				return base.I64_reinterpret_f64(v121)
			}
		}
	} else {
		v23 = F_float_overflow_error_ext(m, int32(0))
		mBase = m.M
		v26 = m.ExcPending
		if v26 != 0 {
			return int64(0)
		} else {
			v27 = v23
			v28 = *(*float64)(unsafe.Add(mBase, uint32(v6)+8))
			v29 = *(*float64)(unsafe.Add(mBase, uint32(v8)+8))
			v30 = base.F64_sub(v28, v29)
			v32 = math.Float64frombits(uint64(0x7ff0000000000000))
			if base.F64_ne(base.F64_abs(v30), v32)|base.F64_eq(base.F64_abs(v28), v32)|base.F64_eq(base.F64_abs(v29), v32) != 0 {
				v45 = v30
				v54 = m.G0
				v56 = v54 - int32(32)
				m.G0 = v56
				v58 = base.F64_abs(v27)
				v59 = base.F64_abs(v45)
				v62 = base.B2i32(base.Ui64(base.I64_reinterpret_f64(v58)) < base.Ui64(base.I64_reinterpret_f64(v59)))
				if base.Ui64(base.I64_reinterpret_f64(v58)) < base.Ui64(base.I64_reinterpret_f64(v59)) {
					v63 = v58
				} else {
					v63 = v59
				}
				v64 = base.I64_reinterpret_f64(v63)
				v66 = int64(base.Ui64(v64) >> (uint(int64(52)) % 64))
				if v66 == int64(2047) {
					v121 = v63
				} else {
					if base.Ui64(base.I64_reinterpret_f64(v58)) < base.Ui64(base.I64_reinterpret_f64(v59)) {
						v69 = v59
					} else {
						v69 = v58
					}
					if v64 == int64(0) {
						v121 = v69
					} else {
						v72 = base.I64_reinterpret_f64(v69)
						v74 = int64(base.Ui64(v72) >> (uint(int64(52)) % 64))
						if v74 == int64(2047) {
							v121 = v69
						} else {
							if int32(65) <= base.I32_wrap_i64(v74)-base.I32_wrap_i64(v66) {
								v121 = base.F64_add(v58, v59)
							} else {
								if base.Ui64(int64(6908521828386340864)) <= base.Ui64(v72) {
									v85 = float64(1.90109156629516e-211)
									v98 = base.F64_mul(v69, v85)
									v99 = base.F64_mul(v63, v85)
									v100 = float64(5.260135901548374e+210)
								} else {
									if base.Ui64(int64(2580562586483294207)) < base.Ui64(v64) {
										v98 = v69
										v99 = v63
										v100 = float64(1)
									} else {
										v93 = float64(5.260135901548374e+210)
										v98 = base.F64_mul(v69, v93)
										v99 = base.F64_mul(v63, v93)
										v100 = float64(1.90109156629516e-211)
									}
								}
								F_sq(m, v56+int32(24), v56+int32(16), v98)
								mBase = m.M
								F_sq(m, v56+int32(8), v56, v99)
								mBase = m.M
								v109 = *(*float64)(unsafe.Add(mBase, uint32(v56)))
								v110 = *(*float64)(unsafe.Add(mBase, uint32(v56)+16))
								v112 = *(*float64)(unsafe.Add(mBase, uint32(v56)+8))
								v114 = *(*float64)(unsafe.Add(mBase, uint32(v56)+24))
								v121 = base.F64_mul(v100, base.F64_sqrt(base.F64_add(base.F64_add(base.F64_add(v109, v110), v112), v114)))
							}
						}
					}
				}
				m.G0 = v56 + int32(32)
				return base.I64_reinterpret_f64(v121)
			} else {
				v43 = F_float_overflow_error_ext(m, int32(0))
				mBase = m.M
				v44 = m.ExcPending
				if v44 != 0 {
					return int64(0)
				} else {
					v45 = v43
					v54 = m.G0
					v56 = v54 - int32(32)
					m.G0 = v56
					v58 = base.F64_abs(v27)
					v59 = base.F64_abs(v45)
					v62 = base.B2i32(base.Ui64(base.I64_reinterpret_f64(v58)) < base.Ui64(base.I64_reinterpret_f64(v59)))
					if base.Ui64(base.I64_reinterpret_f64(v58)) < base.Ui64(base.I64_reinterpret_f64(v59)) {
						v63 = v58
					} else {
						v63 = v59
					}
					v64 = base.I64_reinterpret_f64(v63)
					v66 = int64(base.Ui64(v64) >> (uint(int64(52)) % 64))
					if v66 == int64(2047) {
						v121 = v63
					} else {
						if base.Ui64(base.I64_reinterpret_f64(v58)) < base.Ui64(base.I64_reinterpret_f64(v59)) {
							v69 = v59
						} else {
							v69 = v58
						}
						if v64 == int64(0) {
							v121 = v69
						} else {
							v72 = base.I64_reinterpret_f64(v69)
							v74 = int64(base.Ui64(v72) >> (uint(int64(52)) % 64))
							if v74 == int64(2047) {
								v121 = v69
							} else {
								if int32(65) <= base.I32_wrap_i64(v74)-base.I32_wrap_i64(v66) {
									v121 = base.F64_add(v58, v59)
								} else {
									if base.Ui64(int64(6908521828386340864)) <= base.Ui64(v72) {
										v85 = float64(1.90109156629516e-211)
										v98 = base.F64_mul(v69, v85)
										v99 = base.F64_mul(v63, v85)
										v100 = float64(5.260135901548374e+210)
									} else {
										if base.Ui64(int64(2580562586483294207)) < base.Ui64(v64) {
											v98 = v69
											v99 = v63
											v100 = float64(1)
										} else {
											v93 = float64(5.260135901548374e+210)
											v98 = base.F64_mul(v69, v93)
											v99 = base.F64_mul(v63, v93)
											v100 = float64(1.90109156629516e-211)
										}
									}
									F_sq(m, v56+int32(24), v56+int32(16), v98)
									mBase = m.M
									F_sq(m, v56+int32(8), v56, v99)
									mBase = m.M
									v109 = *(*float64)(unsafe.Add(mBase, uint32(v56)))
									v110 = *(*float64)(unsafe.Add(mBase, uint32(v56)+16))
									v112 = *(*float64)(unsafe.Add(mBase, uint32(v56)+8))
									v114 = *(*float64)(unsafe.Add(mBase, uint32(v56)+24))
									v121 = base.F64_mul(v100, base.F64_sqrt(base.F64_add(base.F64_add(base.F64_add(v109, v110), v112), v114)))
								}
							}
						}
					}
					m.G0 = v56 + int32(32)
					return base.I64_reinterpret_f64(v121)
				}
			}
		}
	}
}
func F_point_left(m *base.Module, l0 int32) int64 {
	mBase := m.M
	_ = mBase
	var v2 int32
	_ = v2
	var v3 float64
	_ = v3
	var v4 int32
	_ = v4
	var v5 float64
	_ = v5
	v2 = *(*int32)(unsafe.Add(mBase, uint32(l0)+40))
	v3 = *(*float64)(unsafe.Add(mBase, uint32(v2)))
	v4 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
	v5 = *(*float64)(unsafe.Add(mBase, uint32(v4)))
	return base.I64_extend_i32_u(base.F64_gt(v3, base.F64_add(v5, float64(1e-06))))
}
func F_point_mul_point(m *base.Module, l0 int32, l1 int32, l2 int32) {
	mBase := m.M
	_ = mBase
	var v9 float64
	_ = v9
	var v10 float64
	_ = v10
	var v11 float64
	_ = v11
	var v13 float64
	_ = v13
	var v26 float64
	_ = v26
	var v27 int32
	_ = v27
	var v28 float64
	_ = v28
	var v37 float64
	_ = v37
	var v38 int32
	_ = v38
	var v39 float64
	_ = v39
	var v40 float64
	_ = v40
	var v41 float64
	_ = v41
	var v42 float64
	_ = v42
	var v44 float64
	_ = v44
	var v57 float64
	_ = v57
	var v58 int32
	_ = v58
	var v59 float64
	_ = v59
	var v68 float64
	_ = v68
	var v69 int32
	_ = v69
	var v70 float64
	_ = v70
	var v72 float64
	_ = v72
	var v74 float64
	_ = v74
	var v86 float64
	_ = v86
	var v87 int32
	_ = v87
	var v88 float64
	_ = v88
	var v89 float64
	_ = v89
	var v90 float64
	_ = v90
	var v91 float64
	_ = v91
	var v93 float64
	_ = v93
	var v106 float64
	_ = v106
	var v107 int32
	_ = v107
	var v108 float64
	_ = v108
	var v117 float64
	_ = v117
	var v118 int32
	_ = v118
	var v119 float64
	_ = v119
	var v120 float64
	_ = v120
	var v121 float64
	_ = v121
	var v122 float64
	_ = v122
	var v124 float64
	_ = v124
	var v137 float64
	_ = v137
	var v138 int32
	_ = v138
	var v139 float64
	_ = v139
	var v148 float64
	_ = v148
	var v149 int32
	_ = v149
	var v150 float64
	_ = v150
	var v152 float64
	_ = v152
	var v154 float64
	_ = v154
	var v164 float64
	_ = v164
	var v165 int32
	_ = v165
	var v166 float64
	_ = v166
	v9 = *(*float64)(unsafe.Add(mBase, uint32(l1)))
	v10 = *(*float64)(unsafe.Add(mBase, uint32(l2)))
	v11 = base.F64_mul(v9, v10)
	v13 = math.Float64frombits(uint64(0x7ff0000000000000))
	if base.F64_ne(base.F64_abs(v11), v13)|base.F64_eq(base.F64_abs(v9), v13)|base.F64_eq(base.F64_abs(v10), v13) == int32(0) {
		goto L2
	} else {
		goto L3
	}
L1:
	;
	v40 = *(*float64)(unsafe.Add(mBase, uint32(l1)+8))
	v41 = *(*float64)(unsafe.Add(mBase, uint32(l2)+8))
	v42 = base.F64_mul(v40, v41)
	v44 = math.Float64frombits(uint64(0x7ff0000000000000))
	if base.F64_ne(base.F64_abs(v42), v44)|base.F64_eq(base.F64_abs(v40), v44)|base.F64_eq(base.F64_abs(v41), v44) == int32(0) {
		goto L10
	} else {
		goto L11
	}
L2:
	;
	v26 = F_float_overflow_error_ext(m, int32(0))
	mBase = m.M
	v27 = m.ExcPending
	if v27 != 0 {
		goto L5
	} else {
		goto L6
	}
L3:
	;
	goto L4
L4:
	;
	v28 = float64(0)
	if base.F64_eq(v9, v28)|base.F64_ne(v11, v28)|base.F64_eq(v10, v28) != 0 {
		v39 = v11
		goto L1
	} else {
		goto L7
	}
L5:
	;
	return
L6:
	;
	v39 = v26
	goto L1
L7:
	;
	v37 = F_float_underflow_error_ext(m, int32(0))
	mBase = m.M
	v38 = m.ExcPending
	if v38 != 0 {
		goto L5
	} else {
		goto L8
	}
L8:
	;
	v39 = v37
	goto L1
L9:
	;
	v72 = math.Float64frombits(uint64(0x7ff0000000000000))
	v74 = base.F64_sub(v39, v70)
	if base.F64_eq(base.F64_abs(v39), v72)|base.F64_ne(base.F64_abs(v74), v72)|base.F64_eq(base.F64_abs(v70), v72) == int32(0) {
		goto L16
	} else {
		goto L17
	}
L10:
	;
	v57 = F_float_overflow_error_ext(m, int32(0))
	mBase = m.M
	v58 = m.ExcPending
	if v58 != 0 {
		goto L5
	} else {
		goto L13
	}
L11:
	;
	goto L12
L12:
	;
	v59 = float64(0)
	if base.F64_eq(v40, v59)|base.F64_ne(v42, v59)|base.F64_eq(v41, v59) != 0 {
		v70 = v42
		goto L9
	} else {
		goto L14
	}
L13:
	;
	v70 = v57
	goto L9
L14:
	;
	v68 = F_float_underflow_error_ext(m, int32(0))
	mBase = m.M
	v69 = m.ExcPending
	if v69 != 0 {
		goto L5
	} else {
		goto L15
	}
L15:
	;
	v70 = v68
	goto L9
L16:
	;
	v86 = F_float_overflow_error_ext(m, int32(0))
	mBase = m.M
	v87 = m.ExcPending
	if v87 != 0 {
		goto L5
	} else {
		goto L19
	}
L17:
	;
	v88 = v74
	goto L18
L18:
	;
	v89 = *(*float64)(unsafe.Add(mBase, uint32(l1)))
	v90 = *(*float64)(unsafe.Add(mBase, uint32(l2)+8))
	v91 = base.F64_mul(v89, v90)
	v93 = math.Float64frombits(uint64(0x7ff0000000000000))
	if base.F64_ne(base.F64_abs(v91), v93)|base.F64_eq(base.F64_abs(v89), v93)|base.F64_eq(base.F64_abs(v90), v93) == int32(0) {
		goto L21
	} else {
		goto L22
	}
L19:
	;
	v88 = v86
	goto L18
L20:
	;
	v120 = *(*float64)(unsafe.Add(mBase, uint32(l1)+8))
	v121 = *(*float64)(unsafe.Add(mBase, uint32(l2)))
	v122 = base.F64_mul(v120, v121)
	v124 = math.Float64frombits(uint64(0x7ff0000000000000))
	if base.F64_ne(base.F64_abs(v122), v124)|base.F64_eq(base.F64_abs(v120), v124)|base.F64_eq(base.F64_abs(v121), v124) == int32(0) {
		goto L28
	} else {
		goto L29
	}
L21:
	;
	v106 = F_float_overflow_error_ext(m, int32(0))
	mBase = m.M
	v107 = m.ExcPending
	if v107 != 0 {
		goto L5
	} else {
		goto L24
	}
L22:
	;
	goto L23
L23:
	;
	v108 = float64(0)
	if base.F64_eq(v89, v108)|base.F64_ne(v91, v108)|base.F64_eq(v90, v108) != 0 {
		v119 = v91
		goto L20
	} else {
		goto L25
	}
L24:
	;
	v119 = v106
	goto L20
L25:
	;
	v117 = F_float_underflow_error_ext(m, int32(0))
	mBase = m.M
	v118 = m.ExcPending
	if v118 != 0 {
		goto L5
	} else {
		goto L26
	}
L26:
	;
	v119 = v117
	goto L20
L27:
	;
	v152 = math.Float64frombits(uint64(0x7ff0000000000000))
	v154 = base.F64_add(v119, v150)
	if base.F64_eq(base.F64_abs(v119), v152)|base.F64_ne(base.F64_abs(v154), v152)|base.F64_eq(base.F64_abs(v150), v152) != 0 {
		goto L34
	} else {
		goto L35
	}
L28:
	;
	v137 = F_float_overflow_error_ext(m, int32(0))
	mBase = m.M
	v138 = m.ExcPending
	if v138 != 0 {
		goto L5
	} else {
		goto L31
	}
L29:
	;
	goto L30
L30:
	;
	v139 = float64(0)
	if base.F64_eq(v120, v139)|base.F64_ne(v122, v139)|base.F64_eq(v121, v139) != 0 {
		v150 = v122
		goto L27
	} else {
		goto L32
	}
L31:
	;
	v150 = v137
	goto L27
L32:
	;
	v148 = F_float_underflow_error_ext(m, int32(0))
	mBase = m.M
	v149 = m.ExcPending
	if v149 != 0 {
		goto L5
	} else {
		goto L33
	}
L33:
	;
	v150 = v148
	goto L27
L34:
	;
	v166 = v154
	goto L36
L35:
	;
	v164 = F_float_overflow_error_ext(m, int32(0))
	mBase = m.M
	v165 = m.ExcPending
	if v165 != 0 {
		goto L5
	} else {
		goto L37
	}
L36:
	;
	*(*float64)(unsafe.Add(mBase, uint32(l0)+8)) = v166
	*(*float64)(unsafe.Add(mBase, uint32(l0))) = v88
	return
L37:
	;
	v166 = v164
	goto L36
}
func F_point_ne(m *base.Module, l0 int32) int64 {
	mBase := m.M
	_ = mBase
	var v10 int32
	_ = v10
	var v11 int32
	_ = v11
	var v12 float64
	_ = v12
	var v18 float64
	_ = v18
	var v20 int64
	_ = v20
	var v21 int64
	_ = v21
	var v22 float64
	_ = v22
	var v25 int64
	_ = v25
	var v30 int32
	_ = v30
	var v33 float64
	_ = v33
	var v41 int64
	_ = v41
	var v46 float64
	_ = v46
	var v65 int32
	_ = v65
	var v67 int32
	_ = v67
	var v74 float64
	_ = v74
	var v78 int64
	_ = v78
	var v79 float64
	_ = v79
	var v82 int64
	_ = v82
	var v91 int32
	_ = v91
	v10 = *(*int32)(unsafe.Add(mBase, uint32(l0)+40))
	v11 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
	v12 = *(*float64)(unsafe.Add(mBase, uint32(v11)))
	if base.Ui64(base.I64_reinterpret_f64(v12)&int64(9223372036854775807)) <= base.Ui64(int64(9218868437227405312)) {
		v18 = *(*float64)(unsafe.Add(mBase, uint32(v10)))
		v20 = int64(9223372036854775807)
		v21 = base.I64_reinterpret_f64(v18) & v20
		v22 = *(*float64)(unsafe.Add(mBase, uint32(v11)+8))
		v25 = base.I64_reinterpret_f64(v22) & v20
		if base.Ui64(int64(9218868437227405313)) <= base.Ui64(v25) {
			v65 = base.B2i32(base.Ui64(v21) < base.Ui64(int64(9218868437227405313)))
			v67 = int32(0)
			if base.B2i32(v65 == v67)|base.F64_ne(v12, v18) != 0 {
				v91 = v67
			} else {
				v74 = v22
				v78 = v25
				v79 = *(*float64)(unsafe.Add(mBase, uint32(v10)+8))
				v82 = base.I64_reinterpret_f64(v79) & int64(9223372036854775807)
				if base.Ui64(int64(9218868437227405313)) <= base.Ui64(v78) {
					v91 = base.B2i32(base.Ui64(int64(9218868437227405312)) < base.Ui64(v82))
				} else {
					v91 = base.B2i32(base.Ui64(v82) < base.Ui64(int64(9218868437227405313))) & base.F64_eq(v79, v74)
				}
			}
		} else {
			v30 = int32(0)
			if base.Ui64(int64(9218868437227405312)) < base.Ui64(v21) {
				v91 = v30
			} else {
				v33 = *(*float64)(unsafe.Add(mBase, uint32(v10)+8))
				if base.Ui64(base.I64_reinterpret_f64(v33)&int64(9223372036854775807)) <= base.Ui64(int64(9218868437227405312)) {
					if base.B2i32(base.F64_le(base.F64_abs(base.F64_sub(v12, v18)), float64(1e-06)) == int32(0))&base.F64_ne(v12, v18) != 0 {
						v91 = v30
					} else {
						v91 = base.F64_eq(v22, v33) | base.F64_le(base.F64_abs(base.F64_sub(v22, v33)), float64(1e-06))
					}
				} else {
					v65 = int32(1)
					v67 = int32(0)
					if base.B2i32(v65 == v67)|base.F64_ne(v12, v18) != 0 {
						v91 = v67
					} else {
						v74 = v22
						v78 = v25
						v79 = *(*float64)(unsafe.Add(mBase, uint32(v10)+8))
						v82 = base.I64_reinterpret_f64(v79) & int64(9223372036854775807)
						if base.Ui64(int64(9218868437227405313)) <= base.Ui64(v78) {
							v91 = base.B2i32(base.Ui64(int64(9218868437227405312)) < base.Ui64(v82))
						} else {
							v91 = base.B2i32(base.Ui64(v82) < base.Ui64(int64(9218868437227405313))) & base.F64_eq(v79, v74)
						}
					}
				}
			}
		}
	} else {
		v41 = *(*int64)(unsafe.Add(mBase, uint32(v10)))
		if base.Ui64(v41&int64(9223372036854775807)) < base.Ui64(int64(9218868437227405313)) {
			v91 = int32(0)
		} else {
			v46 = *(*float64)(unsafe.Add(mBase, uint32(v11)+8))
			v74 = v46
			v78 = base.I64_reinterpret_f64(v46) & int64(9223372036854775807)
			v79 = *(*float64)(unsafe.Add(mBase, uint32(v10)+8))
			v82 = base.I64_reinterpret_f64(v79) & int64(9223372036854775807)
			if base.Ui64(int64(9218868437227405313)) <= base.Ui64(v78) {
				v91 = base.B2i32(base.Ui64(int64(9218868437227405312)) < base.Ui64(v82))
			} else {
				v91 = base.B2i32(base.Ui64(v82) < base.Ui64(int64(9218868437227405313))) & base.F64_eq(v79, v74)
			}
		}
	}
	return base.I64_extend_i32_u(v91 ^ int32(1))
}
func F_point_slope(m *base.Module, l0 int32) int64 {
	mBase := m.M
	_ = mBase
	var v2 int32
	_ = v2
	var v3 int32
	_ = v3
	var v4 float64
	_ = v4
	var v7 int32
	_ = v7
	v2 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
	v3 = *(*int32)(unsafe.Add(mBase, uint32(l0)+40))
	v4 = F_point_sl(m, v2, v3)
	mBase = m.M
	v7 = m.ExcPending
	if v7 != 0 {
		return int64(0)
	} else {
		return base.I64_reinterpret_f64(v4)
	}
}
