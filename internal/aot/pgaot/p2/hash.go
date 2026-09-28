package p2

import (
	base "github.com/shibukawa/pgmem/internal/aot/pgaot/base"
	"unsafe"
)

func F_ExecHashJoinGetSavedTuple(m *base.Module, l0 int32, l1 int32, l2 int32) int32 {
	mBase := m.M
	_ = mBase
	var v6 int32
	_ = v6
	var v8 int32
	_ = v8
	var v11 int32
	_ = v11
	var v15 int32
	_ = v15
	var v16 int32
	_ = v16
	var v20 int32
	_ = v20
	var v21 int32
	_ = v21
	var v24 int32
	_ = v24
	var v25 int32
	_ = v25
	var v27 int32
	_ = v27
	var v29 int32
	_ = v29
	var v31 int32
	_ = v31
	var v32 int32
	_ = v32
	var v33 int32
	_ = v33
	var v34 int32
	_ = v34
	var v36 int32
	_ = v36
	var v41 int32
	_ = v41
	var v44 int32
	_ = v44
	var v46 int32
	_ = v46
	v6 = m.G0
	v8 = v6 - int32(16)
	m.G0 = v8
	v11 = *(*int32)(unsafe.Add(mBase, _c_F_ExecHashJoinGetSavedTuple[0]))
	if v11 != 0 {
		F_ProcessInterrupts(m)
		mBase = m.M
		v15 = m.ExcPending
		if v15 != 0 {
			return int32(0)
		} else {
			v16 = int32(8)
			v20 = F_BufFileReadCommon(m, l0, v8+v16, v16, int32(1))
			mBase = m.M
			v21 = m.ExcPending
			if v21 != 0 {
				return int32(0)
			} else {
				if v20 == int32(0) {
					v24 = *(*int32)(unsafe.Add(mBase, uint32(l2)+8))
					v25 = *(*int32)(unsafe.Add(mBase, uint32(v24)+12))
					m.T0[v25].(func(*base.Module, int32))(m, l2)
					mBase = m.M
					v27 = m.ExcPending
					if v27 != 0 {
						return int32(0)
					} else {
						v46 = int32(0)
						m.G0 = v8 + int32(16)
						return v46
					}
				} else {
					v29 = *(*int32)(unsafe.Add(mBase, uint32(v8)+8))
					*(*int32)(unsafe.Add(mBase, uint32(l1))) = v29
					v31 = *(*int32)(unsafe.Add(mBase, uint32(v8)+12))
					v32 = F_palloc(m, v31)
					mBase = m.M
					v33 = m.ExcPending
					if v33 != 0 {
						return int32(0)
					} else {
						v34 = *(*int32)(unsafe.Add(mBase, uint32(v8)+12))
						*(*int32)(unsafe.Add(mBase, uint32(v32))) = v34
						v36 = int32(4)
						F_BufFileReadExact(m, l0, v32+v36, v34-v36)
						mBase = m.M
						v41 = m.ExcPending
						if v41 != 0 {
							return int32(0)
						} else {
							F_ExecForceStoreMinimalTuple(m, v32, l2, int32(1))
							mBase = m.M
							v44 = m.ExcPending
							if v44 != 0 {
								return int32(0)
							} else {
								v46 = l2
								m.G0 = v8 + int32(16)
								return v46
							}
						}
					}
				}
			}
		}
	} else {
		v16 = int32(8)
		v20 = F_BufFileReadCommon(m, l0, v8+v16, v16, int32(1))
		mBase = m.M
		v21 = m.ExcPending
		if v21 != 0 {
			return int32(0)
		} else {
			if v20 == int32(0) {
				v24 = *(*int32)(unsafe.Add(mBase, uint32(l2)+8))
				v25 = *(*int32)(unsafe.Add(mBase, uint32(v24)+12))
				m.T0[v25].(func(*base.Module, int32))(m, l2)
				mBase = m.M
				v27 = m.ExcPending
				if v27 != 0 {
					return int32(0)
				} else {
					v46 = int32(0)
					m.G0 = v8 + int32(16)
					return v46
				}
			} else {
				v29 = *(*int32)(unsafe.Add(mBase, uint32(v8)+8))
				*(*int32)(unsafe.Add(mBase, uint32(l1))) = v29
				v31 = *(*int32)(unsafe.Add(mBase, uint32(v8)+12))
				v32 = F_palloc(m, v31)
				mBase = m.M
				v33 = m.ExcPending
				if v33 != 0 {
					return int32(0)
				} else {
					v34 = *(*int32)(unsafe.Add(mBase, uint32(v8)+12))
					*(*int32)(unsafe.Add(mBase, uint32(v32))) = v34
					v36 = int32(4)
					F_BufFileReadExact(m, l0, v32+v36, v34-v36)
					mBase = m.M
					v41 = m.ExcPending
					if v41 != 0 {
						return int32(0)
					} else {
						F_ExecForceStoreMinimalTuple(m, v32, l2, int32(1))
						mBase = m.M
						v44 = m.ExcPending
						if v44 != 0 {
							return int32(0)
						} else {
							v46 = l2
							m.G0 = v8 + int32(16)
							return v46
						}
					}
				}
			}
		}
	}
}
func F_ExecHashJoinSaveTuple(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) {
	mBase := m.M
	_ = mBase
	var v7 int32
	_ = v7
	var v9 int32
	_ = v9
	var v12 int32
	_ = v12
	var v15 int32
	_ = v15
	var v16 int32
	_ = v16
	var v18 int32
	_ = v18
	var v21 int32
	_ = v21
	var v22 int32
	_ = v22
	var v26 int32
	_ = v26
	var v32 int32
	_ = v32
	var v33 int32
	_ = v33
	var v35 int32
	_ = v35
	v7 = m.G0
	v9 = v7 - int32(16)
	m.G0 = v9
	*(*int32)(unsafe.Add(mBase, uint32(v9)+12)) = l1
	v12 = *(*int32)(unsafe.Add(mBase, uint32(l2)))
	if v12 == int32(0) {
		v15 = int32(_a_F_ExecHashJoinSaveTuple_0)
		v16 = *(*int32)(unsafe.Add(mBase, _c_F_ExecHashJoinSaveTuple[0]))
		v18 = *(*int32)(unsafe.Add(mBase, uint32(l3)+124))
		*(*int32)(unsafe.Add(mBase, _c_F_ExecHashJoinSaveTuple[0])) = v18
		v21 = F_BufFileCreateTemp(m, int32(0))
		mBase = m.M
		v22 = m.ExcPending
		if v22 != 0 {
			return
		} else {
			*(*int32)(unsafe.Add(mBase, uint32(l2))) = v21
			*(*int32)(unsafe.Add(mBase, _c_F_ExecHashJoinSaveTuple[0])) = v16
			v26 = v21
			F_BufFileWrite(m, v26, v9+int32(12), int32(4))
			mBase = m.M
			v32 = m.ExcPending
			if v32 != 0 {
				return
			} else {
				v33 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
				F_BufFileWrite(m, v26, l0, v33)
				mBase = m.M
				v35 = m.ExcPending
				if v35 != 0 {
					return
				} else {
					m.G0 = v9 + int32(16)
					return
				}
			}
		}
	} else {
		v26 = v12
		F_BufFileWrite(m, v26, v9+int32(12), int32(4))
		mBase = m.M
		v32 = m.ExcPending
		if v32 != 0 {
			return
		} else {
			v33 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
			F_BufFileWrite(m, v26, l0, v33)
			mBase = m.M
			v35 = m.ExcPending
			if v35 != 0 {
				return
			} else {
				m.G0 = v9 + int32(16)
				return
			}
		}
	}
}
func F_ExecHashTableDetach(m *base.Module, l0 int32) {
	mBase := m.M
	_ = mBase
	var v2 int32
	_ = v2
	var v6 int32
	_ = v6
	var v10 int32
	_ = v10
	var v11 int32
	_ = v11
	var v14 int32
	_ = v14
	var v17 int32
	_ = v17
	var v23 int32
	_ = v23
	var v26 int32
	_ = v26
	var v27 int32
	_ = v27
	var v29 int32
	_ = v29
	var v31 int32
	_ = v31
	var v32 int32
	_ = v32
	var v34 int32
	_ = v34
	var v36 int32
	_ = v36
	var v37 int32
	_ = v37
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
	var v48 int32
	_ = v48
	var v49 int32
	_ = v49
	var v56 int32
	_ = v56
	var v57 int32
	_ = v57
	var v60 int32
	_ = v60
	var v63 int32
	_ = v63
	var v65 int32
	_ = v65
	v2 = int32(0)
	v6 = *(*int32)(unsafe.Add(mBase, uint32(l0)+140))
	if v6 == v2 {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+140)) = int32(0)
	return
L2:
	;
	v10 = v6 + int32(56)
	v11 = *(*int32)(unsafe.Add(mBase, uint32(v10)+4))
	if v11 != int32(4) {
		goto L1
	} else {
		goto L3
	}
L3:
	;
	v14 = *(*int32)(unsafe.Add(mBase, uint32(l0)+144))
	if v14 == int32(0) {
		goto L4
	} else {
		goto L5
	}
L4:
	;
	v56 = F_BarrierArriveAndDetach(m, v10)
	mBase = m.M
	v57 = m.ExcPending
	if v57 != 0 {
		goto L9
	} else {
		goto L15
	}
L5:
	;
	v17 = *(*int32)(unsafe.Add(mBase, uint32(l0)+44))
	if v17 <= int32(0) {
		goto L4
	} else {
		goto L6
	}
L6:
	;
	v23 = v2
	goto L7
L7:
	;
	v26 = v23 * int32(36)
	v27 = *(*int32)(unsafe.Add(mBase, uint32(l0)+144))
	v29 = *(*int32)(unsafe.Add(mBase, uint32(v26+v27)+28))
	F_sts_end_write(m, v29)
	mBase = m.M
	v31 = m.ExcPending
	if v31 != 0 {
		goto L9
	} else {
		goto L10
	}
L8:
	;
	goto L4
L9:
	;
	return
L10:
	;
	v32 = *(*int32)(unsafe.Add(mBase, uint32(l0)+144))
	v34 = *(*int32)(unsafe.Add(mBase, uint32(v32+v26)+32))
	F_sts_end_write(m, v34)
	mBase = m.M
	v36 = m.ExcPending
	if v36 != 0 {
		goto L9
	} else {
		goto L11
	}
L11:
	;
	v37 = *(*int32)(unsafe.Add(mBase, uint32(l0)+144))
	v39 = *(*int32)(unsafe.Add(mBase, uint32(v37+v26)+28))
	F_sts_end_parallel_scan(m, v39)
	mBase = m.M
	v41 = m.ExcPending
	if v41 != 0 {
		goto L9
	} else {
		goto L12
	}
L12:
	;
	v42 = *(*int32)(unsafe.Add(mBase, uint32(l0)+144))
	v44 = *(*int32)(unsafe.Add(mBase, uint32(v42+v26)+32))
	F_sts_end_parallel_scan(m, v44)
	mBase = m.M
	v46 = m.ExcPending
	if v46 != 0 {
		goto L9
	} else {
		goto L13
	}
L13:
	;
	v48 = v23 + int32(1)
	v49 = *(*int32)(unsafe.Add(mBase, uint32(l0)+44))
	if v48 < v49 {
		v23 = v48
		goto L7
	} else {
		goto L14
	}
L14:
	;
	goto L8
L15:
	;
	if v56 == int32(0) {
		goto L1
	} else {
		goto L16
	}
L16:
	;
	v60 = *(*int32)(unsafe.Add(mBase, uint32(v6)))
	if v60 == int32(0) {
		goto L1
	} else {
		goto L17
	}
L17:
	;
	v63 = *(*int32)(unsafe.Add(mBase, uint32(l0)+136))
	F_dsa_free(m, v63, v60)
	mBase = m.M
	v65 = m.ExcPending
	if v65 != 0 {
		goto L9
	} else {
		goto L18
	}
L18:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v6))) = int32(0)
	goto L1
}
func F__hash_doinsert(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) {
	mBase := m.M
	_ = mBase
	var v5 int32
	_ = v5
	var v28 int32
	_ = v28
	var v30 int32
	_ = v30
	var v36 int32
	_ = v36
	var v39 int32
	_ = v39
	var v41 int32
	_ = v41
	var v42 int32
	_ = v42
	var v48 int32
	_ = v48
	var v79 int32
	_ = v79
	var v80 int32
	_ = v80
	var v84 int32
	_ = v84
	var v90 int32
	_ = v90
	var v92 int32
	_ = v92
	var v98 int32
	_ = v98
	var v99 int32
	_ = v99
	var v111 int32
	_ = v111
	var v112 int32
	_ = v112
	var v116 int32
	_ = v116
	var v122 int32
	_ = v122
	var v124 int32
	_ = v124
	var v125 int32
	_ = v125
	var v130 int32
	_ = v130
	var v131 int32
	_ = v131
	var v133 int32
	_ = v133
	var v137 int32
	_ = v137
	var v143 int32
	_ = v143
	var v145 int32
	_ = v145
	var v151 int32
	_ = v151
	var v152 int32
	_ = v152
	var v153 int32
	_ = v153
	var v154 int32
	_ = v154
	var v159 int32
	_ = v159
	var v160 int32
	_ = v160
	var v161 int32
	_ = v161
	var v165 int32
	_ = v165
	var v166 int32
	_ = v166
	var v167 int32
	_ = v167
	var v168 int32
	_ = v168
	var v169 int32
	_ = v169
	var v171 int32
	_ = v171
	var v173 int32
	_ = v173
	var v175 int32
	_ = v175
	var v177 int32
	_ = v177
	var v178 int32
	_ = v178
	var v179 int32
	_ = v179
	var v180 int32
	_ = v180
	var v183 int32
	_ = v183
	var v188 int32
	_ = v188
	var v190 int32
	_ = v190
	var v195 int32
	_ = v195
	var v199 int32
	_ = v199
	var v207 int32
	_ = v207
	var v218 int32
	_ = v218
	var v223 int32
	_ = v223
	var v224 int32
	_ = v224
	var v227 int32
	_ = v227
	var v228 int32
	_ = v228
	var v232 int32
	_ = v232
	var v238 int32
	_ = v238
	var v240 int32
	_ = v240
	var v246 int32
	_ = v246
	var v247 int32
	_ = v247
	var v253 int32
	_ = v253
	var v258 int32
	_ = v258
	var v260 int32
	_ = v260
	var v261 int32
	_ = v261
	var v265 int32
	_ = v265
	var v268 int32
	_ = v268
	var v271 int32
	_ = v271
	var v272 int32
	_ = v272
	var v273 int32
	_ = v273
	var v285 int32
	_ = v285
	var v289 int32
	_ = v289
	var v292 int32
	_ = v292
	var v310 int32
	_ = v310
	var v311 int32
	_ = v311
	var v317 int32
	_ = v317
	var v323 int32
	_ = v323
	var v325 int32
	_ = v325
	var v329 int32
	_ = v329
	var v330 int32
	_ = v330
	var v336 int32
	_ = v336
	var v342 int32
	_ = v342
	var v343 int32
	_ = v343
	var v344 int32
	_ = v344
	var v346 int32
	_ = v346
	var v355 int32
	_ = v355
	var v359 int32
	_ = v359
	var v380 int32
	_ = v380
	var v381 int32
	_ = v381
	var v387 int32
	_ = v387
	var v402 int32
	_ = v402
	var v423 int32
	_ = v423
	var v424 int32
	_ = v424
	var v425 int32
	_ = v425
	var v428 int32
	_ = v428
	var v429 int32
	_ = v429
	var v431 int32
	_ = v431
	var v436 int32
	_ = v436
	var v437 int32
	_ = v437
	var v438 int32
	_ = v438
	var v439 int32
	_ = v439
	var v441 int32
	_ = v441
	var v443 int32
	_ = v443
	var v444 int32
	_ = v444
	var v448 int32
	_ = v448
	var v452 int32
	_ = v452
	var v454 int32
	_ = v454
	var v458 int32
	_ = v458
	var v459 float64
	_ = v459
	var v464 int32
	_ = v464
	var v466 int32
	_ = v466
	var v467 int32
	_ = v467
	var v468 int32
	_ = v468
	var v472 int32
	_ = v472
	var v473 int32
	_ = v473
	var v474 int32
	_ = v474
	var v477 int32
	_ = v477
	var v478 int32
	_ = v478
	var v482 int32
	_ = v482
	var v486 int32
	_ = v486
	var v487 int32
	_ = v487
	var v491 int32
	_ = v491
	var v493 int32
	_ = v493
	var v498 int32
	_ = v498
	var v501 int32
	_ = v501
	var v506 int32
	_ = v506
	var v507 int32
	_ = v507
	var v510 int32
	_ = v510
	var v511 int32
	_ = v511
	var v517 int32
	_ = v517
	var v521 int32
	_ = v521
	var v526 int32
	_ = v526
	var v532 int32
	_ = v532
	var v536 int32
	_ = v536
	var v539 int64
	_ = v539
	var v540 int32
	_ = v540
	var v543 int64
	_ = v543
	var v544 int32
	_ = v544
	var v547 int64
	_ = v547
	var v551 int32
	_ = v551
	var v557 int32
	_ = v557
	var v559 int32
	_ = v559
	var v565 int32
	_ = v565
	var v567 int64
	_ = v567
	var v572 int32
	_ = v572
	var v576 int32
	_ = v576
	var v578 int32
	_ = v578
	var v582 int32
	_ = v582
	var v584 int32
	_ = v584
	var v586 int32
	_ = v586
	var v591 int32
	_ = v591
	var v619 int32
	_ = v619
	var v620 int32
	_ = v620
	var v621 int32
	_ = v621
	var v622 int32
	_ = v622
	var v625 int32
	_ = v625
	var v656 int32
	_ = v656
	var v661 int32
	_ = v661
	var v663 int32
	_ = v663
	var v666 int32
	_ = v666
	var v667 int32
	_ = v667
	var v671 int32
	_ = v671
	var v677 int32
	_ = v677
	var v679 int32
	_ = v679
	var v686 int32
	_ = v686
	var v688 int32
	_ = v688
	var v689 int32
	_ = v689
	var v693 int32
	_ = v693
	var v699 int32
	_ = v699
	var v701 int32
	_ = v701
	var v707 int32
	_ = v707
	var v708 int32
	_ = v708
	var v709 int32
	_ = v709
	var v711 int32
	_ = v711
	var v712 int32
	_ = v712
	var v713 int32
	_ = v713
	var v714 int32
	_ = v714
	var v717 int32
	_ = v717
	var v725 int32
	_ = v725
	var v750 int32
	_ = v750
	var v751 int32
	_ = v751
	var v753 int32
	_ = v753
	var v757 int32
	_ = v757
	var v759 int32
	_ = v759
	var v763 int32
	_ = v763
	var v767 int32
	_ = v767
	var v773 int32
	_ = v773
	var v775 int32
	_ = v775
	var v781 int32
	_ = v781
	var v782 int32
	_ = v782
	var v783 int32
	_ = v783
	var v792 int32
	_ = v792
	var v795 int32
	_ = v795
	var v798 int32
	_ = v798
	var v800 int32
	_ = v800
	var v806 int32
	_ = v806
	var v808 int32
	_ = v808
	var v817 int32
	_ = v817
	var v826 int32
	_ = v826
	var v827 int32
	_ = v827
	var v831 int32
	_ = v831
	var v837 int32
	_ = v837
	var v841 int32
	_ = v841
	var v844 int32
	_ = v844
	var v847 int32
	_ = v847
	var v850 int32
	_ = v850
	var v852 int32
	_ = v852
	var v853 int32
	_ = v853
	var v854 int32
	_ = v854
	var v859 int32
	_ = v859
	var v866 int32
	_ = v866
	var v873 int32
	_ = v873
	var v875 int32
	_ = v875
	var v877 int32
	_ = v877
	var v878 int32
	_ = v878
	var v884 int32
	_ = v884
	var v885 int32
	_ = v885
	var v891 int32
	_ = v891
	var v896 int32
	_ = v896
	var v901 int32
	_ = v901
	var v902 float64
	_ = v902
	var v904 float64
	_ = v904
	var v906 int32
	_ = v906
	var v907 int32
	_ = v907
	var v909 int32
	_ = v909
	var v910 int32
	_ = v910
	var v911 int32
	_ = v911
	var v915 int32
	_ = v915
	var v918 int32
	_ = v918
	var v919 int32
	_ = v919
	var v922 int32
	_ = v922
	var v927 int32
	_ = v927
	var v931 int32
	_ = v931
	var v935 int32
	_ = v935
	var v937 int32
	_ = v937
	var v941 int32
	_ = v941
	var v944 int64
	_ = v944
	var v945 int32
	_ = v945
	var v946 int64
	_ = v946
	var v947 int32
	_ = v947
	var v948 int64
	_ = v948
	var v956 int32
	_ = v956
	var v962 int32
	_ = v962
	var v964 int32
	_ = v964
	var v970 int32
	_ = v970
	var v972 int64
	_ = v972
	var v978 int32
	_ = v978
	var v984 int32
	_ = v984
	var v986 int32
	_ = v986
	var v992 int32
	_ = v992
	var v994 int32
	_ = v994
	var v996 int32
	_ = v996
	var v1001 int32
	_ = v1001
	var v1003 int32
	_ = v1003
	var v1006 int32
	_ = v1006
	var v1008 int32
	_ = v1008
	var v1012 int32
	_ = v1012
	var v1015 int32
	_ = v1015
	var v1017 int32
	_ = v1017
	var v1049 int32
	_ = v1049
	var v1052 int32
	_ = v1052
	var v1056 int32
	_ = v1056
	var v1058 int32
	_ = v1058
	var v1060 int32
	_ = v1060
	var v1064 int32
	_ = v1064
	var v1065 int32
	_ = v1065
	var v1068 float64
	_ = v1068
	var v1069 int32
	_ = v1069
	var v1072 int32
	_ = v1072
	var v1076 int32
	_ = v1076
	var v1077 int32
	_ = v1077
	var v1078 int32
	_ = v1078
	var v1079 int32
	_ = v1079
	var v1083 int32
	_ = v1083
	var v1089 int32
	_ = v1089
	var v1092 int32
	_ = v1092
	var v1102 int32
	_ = v1102
	var v1106 int32
	_ = v1106
	var v1110 int32
	_ = v1110
	var v1113 int32
	_ = v1113
	var v1114 int32
	_ = v1114
	var v1115 int32
	_ = v1115
	var v1116 int32
	_ = v1116
	var v1120 int32
	_ = v1120
	var v1124 int32
	_ = v1124
	var v1128 int32
	_ = v1128
	var v1132 int32
	_ = v1132
	var v1137 int32
	_ = v1137
	var v1138 int32
	_ = v1138
	var v1144 int32
	_ = v1144
	var v1150 int32
	_ = v1150
	var v1152 int32
	_ = v1152
	var v1158 int32
	_ = v1158
	var v1159 int32
	_ = v1159
	var v1161 int32
	_ = v1161
	var v1164 int32
	_ = v1164
	var v1165 int32
	_ = v1165
	var v1166 int32
	_ = v1166
	var v1168 int32
	_ = v1168
	var v1170 int32
	_ = v1170
	var v1172 int32
	_ = v1172
	var v1174 int32
	_ = v1174
	var v1177 int32
	_ = v1177
	var v1178 int32
	_ = v1178
	var v1179 int32
	_ = v1179
	var v1181 int32
	_ = v1181
	var v1182 int32
	_ = v1182
	var v1189 int32
	_ = v1189
	var v1191 int32
	_ = v1191
	var v1192 int32
	_ = v1192
	var v1193 int32
	_ = v1193
	var v1197 int32
	_ = v1197
	var v1203 int32
	_ = v1203
	var v1206 int32
	_ = v1206
	var v1216 int32
	_ = v1216
	var v1217 int32
	_ = v1217
	var v1220 int32
	_ = v1220
	var v1221 int32
	_ = v1221
	var v1222 int32
	_ = v1222
	var v1223 int32
	_ = v1223
	var v1227 int32
	_ = v1227
	var v1233 int32
	_ = v1233
	var v1236 int32
	_ = v1236
	var v1246 int32
	_ = v1246
	var v1247 int32
	_ = v1247
	var v1256 int32
	_ = v1256
	var v1257 int32
	_ = v1257
	var v1259 int32
	_ = v1259
	var v1270 int32
	_ = v1270
	var v1271 int32
	_ = v1271
	var v1279 int32
	_ = v1279
	var v1282 int32
	_ = v1282
	var v1298 int32
	_ = v1298
	var v1300 int32
	_ = v1300
	var v1302 int32
	_ = v1302
	var v1307 int32
	_ = v1307
	var v1321 int32
	_ = v1321
	var v1327 int32
	_ = v1327
	var v1330 int32
	_ = v1330
	var v1331 int32
	_ = v1331
	var v1336 int32
	_ = v1336
	var v1337 int32
	_ = v1337
	var v1341 int32
	_ = v1341
	var v1344 int32
	_ = v1344
	var v1345 int32
	_ = v1345
	var v1351 int32
	_ = v1351
	var v1355 int32
	_ = v1355
	var v1356 int32
	_ = v1356
	var v1357 int32
	_ = v1357
	var v1358 int32
	_ = v1358
	var v1360 int64
	_ = v1360
	var v1364 int32
	_ = v1364
	var v1365 int32
	_ = v1365
	var v1368 int32
	_ = v1368
	var v1369 int32
	_ = v1369
	var v1370 int32
	_ = v1370
	var v1372 int32
	_ = v1372
	var v1374 int32
	_ = v1374
	var v1376 int32
	_ = v1376
	var v1380 int32
	_ = v1380
	var v1382 int32
	_ = v1382
	var v1383 int32
	_ = v1383
	var v1388 int32
	_ = v1388
	var v1394 int32
	_ = v1394
	var v1395 int32
	_ = v1395
	var v1396 int32
	_ = v1396
	var v1397 int32
	_ = v1397
	var v1401 int32
	_ = v1401
	var v1403 int32
	_ = v1403
	var v1405 int32
	_ = v1405
	var v1406 int32
	_ = v1406
	var v1408 int32
	_ = v1408
	var v1414 int32
	_ = v1414
	var v1415 int32
	_ = v1415
	var v1416 int32
	_ = v1416
	var v1423 int32
	_ = v1423
	var v1424 int32
	_ = v1424
	var v1425 int32
	_ = v1425
	var v1429 int32
	_ = v1429
	var v1430 int32
	_ = v1430
	var v1436 int32
	_ = v1436
	var v1441 int32
	_ = v1441
	var v1442 int32
	_ = v1442
	var v1443 int32
	_ = v1443
	var v1444 int32
	_ = v1444
	var v1448 int32
	_ = v1448
	var v1454 int32
	_ = v1454
	var v1456 int32
	_ = v1456
	var v1462 int32
	_ = v1462
	var v1463 int32
	_ = v1463
	var v1464 int32
	_ = v1464
	var v1466 int32
	_ = v1466
	var v1468 int32
	_ = v1468
	var v1471 int32
	_ = v1471
	var v1475 int32
	_ = v1475
	var v1481 int32
	_ = v1481
	var v1483 int32
	_ = v1483
	var v1489 int32
	_ = v1489
	var v1490 int32
	_ = v1490
	var v1491 int32
	_ = v1491
	var v1499 int32
	_ = v1499
	var v1500 int32
	_ = v1500
	var v1501 int32
	_ = v1501
	var v1505 int32
	_ = v1505
	var v1508 int32
	_ = v1508
	var v1509 int32
	_ = v1509
	var v1511 int32
	_ = v1511
	var v1513 int32
	_ = v1513
	var v1514 int32
	_ = v1514
	var v1518 int32
	_ = v1518
	var v1522 int32
	_ = v1522
	var v1526 int32
	_ = v1526
	var v1527 int32
	_ = v1527
	var v1531 int32
	_ = v1531
	var v1534 int32
	_ = v1534
	var v1539 int32
	_ = v1539
	var v1543 int32
	_ = v1543
	var v1545 int32
	_ = v1545
	var v1552 int32
	_ = v1552
	var v1553 int32
	_ = v1553
	var v1554 int32
	_ = v1554
	var v1562 int32
	_ = v1562
	var v1567 int32
	_ = v1567
	var v1570 int64
	_ = v1570
	var v1571 int32
	_ = v1571
	var v1572 int64
	_ = v1572
	var v1573 int32
	_ = v1573
	var v1575 int64
	_ = v1575
	var v1579 int32
	_ = v1579
	var v1585 int32
	_ = v1585
	var v1587 int32
	_ = v1587
	var v1593 int32
	_ = v1593
	var v1595 int64
	_ = v1595
	var v1600 int32
	_ = v1600
	var v1606 int32
	_ = v1606
	var v1608 int32
	_ = v1608
	var v1614 int32
	_ = v1614
	var v1619 int32
	_ = v1619
	var v1623 int32
	_ = v1623
	var v1625 int32
	_ = v1625
	var v1629 int32
	_ = v1629
	var v1631 int32
	_ = v1631
	var v1633 int32
	_ = v1633
	var v1638 int32
	_ = v1638
	var v1641 int32
	_ = v1641
	var v1643 int32
	_ = v1643
	var v1645 int32
	_ = v1645
	var v1648 int32
	_ = v1648
	var v1659 int32
	_ = v1659
	var v1705 int32
	_ = v1705
	var v1712 int32
	_ = v1712
	var v1715 int32
	_ = v1715
	var v1716 int32
	_ = v1716
	var v1727 int32
	_ = v1727
	var v1731 int32
	_ = v1731
	var v1736 int32
	_ = v1736
	v5 = int32(0)
	v28 = m.G0
	v30 = v28 - int32(_a_F__hash_doinsert_0)
	m.G0 = v30
	*(*int32)(unsafe.Add(mBase, uint32(v30)+20)) = v5
	v36 = int32(*(*int16)(unsafe.Add(mBase, uint32(l1)+6)))
	if v5 <= v36 {
		goto L2
	} else {
		goto L3
	}
L1:
	;
	v42 = int32(*(*uint16)(unsafe.Add(mBase, uint32(l1)+6)))
	v48 = (v42&int32(_a_F__hash_doinsert_1) + int32(7)) & int32(_a_F__hash_doinsert_2)
	goto L6
L2:
	;
	v39 = int32(8)
	goto L4
L3:
	;
	v39 = int32(16)
	goto L4
L4:
	;
	v41 = *(*int32)(unsafe.Add(mBase, uint32(l1+v39)))
	goto L1
L5:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v1712 = m.ExcPending
	if v1712 != 0 {
		goto L9
	} else {
		goto L403
	}
L6:
	;
	v79 = F__hash_getbuf(m, l0, int32(0), int32(-1), int32(8))
	mBase = m.M
	v80 = m.ExcPending
	if v80 != 0 {
		goto L9
	} else {
		goto L10
	}
L7:
	;
	v177 = int32(4)
	v178 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v151)+14)))
	v179 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v151)+12)))
	v180 = v178 - v179
	if v180 <= v177 {
		goto L35
	} else {
		goto L36
	}
L8:
	;
	v99 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v98)+19)))
	if base.Ui32((v99<<(uint(int32(8))%32)-int32(44))&int32(-48)) < base.Ui32(v48) {
		goto L5
	} else {
		goto L14
	}
L9:
	;
	return
L10:
	;
	if v79 < int32(0) {
		goto L11
	} else {
		goto L12
	}
L11:
	;
	v84 = *(*int32)(unsafe.Add(mBase, _c_F__hash_doinsert[0]))
	v90 = *(*int32)(unsafe.Add(mBase, uint32(v84+(v79^int32(-1))<<(uint(int32(2))%32))))
	v98 = v90
	goto L8
L12:
	;
	goto L13
L13:
	;
	v92 = *(*int32)(unsafe.Add(mBase, _c_F__hash_doinsert[1]))
	v98 = v92 + v79<<(uint(int32(13))%32) + int32(-8192)
	goto L8
L14:
	;
	v111 = F__hash_getbucketbuf_from_hashkey(m, l0, v41, int32(3), v30+int32(20))
	mBase = m.M
	v112 = m.ExcPending
	if v112 != 0 {
		goto L9
	} else {
		goto L15
	}
L15:
	;
	if v111 < int32(0) {
		goto L17
	} else {
		goto L18
	}
L16:
	;
	F_CheckForSerializableConflictIn(m, l0, int32(0), v131)
	mBase = m.M
	v133 = m.ExcPending
	if v133 != 0 {
		goto L9
	} else {
		goto L20
	}
L17:
	;
	v116 = *(*int32)(unsafe.Add(mBase, _c_F__hash_doinsert[2]))
	v122 = *(*int32)(unsafe.Add(mBase, uint32(v116+(v111^int32(-1))*int32(56))+16))
	v131 = v122
	goto L16
L18:
	;
	goto L19
L19:
	;
	v124 = *(*int32)(unsafe.Add(mBase, _c_F__hash_doinsert[3]))
	v125 = int32(56)
	v130 = *(*int32)(unsafe.Add(mBase, uint32(v124+v111*v125-v125)+16))
	v131 = v130
	goto L16
L20:
	;
	if v111 < int32(0) {
		goto L23
	} else {
		goto L24
	}
L21:
	;
	goto L7
L22:
	;
	v152 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v151)+16)))
	v153 = v152 + v151
	v154 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v153)+12)))
	if v154&int32(32) == int32(0) {
		goto L21
	} else {
		goto L26
	}
L23:
	;
	v137 = *(*int32)(unsafe.Add(mBase, _c_F__hash_doinsert[0]))
	v143 = *(*int32)(unsafe.Add(mBase, uint32(v137+(v111^int32(-1))<<(uint(int32(2))%32))))
	v151 = v143
	goto L22
L24:
	;
	goto L25
L25:
	;
	v145 = *(*int32)(unsafe.Add(mBase, _c_F__hash_doinsert[1]))
	v151 = v145 + v111<<(uint(int32(13))%32) + int32(-8192)
	goto L22
L26:
	;
	v159 = *(*int32)(unsafe.Add(mBase, uint32(v153)+8))
	v160 = F_IsBufferCleanupOK(m, v111)
	mBase = m.M
	v161 = m.ExcPending
	if v161 != 0 {
		goto L9
	} else {
		goto L27
	}
L27:
	;
	if v160 == int32(0) {
		goto L21
	} else {
		goto L28
	}
L28:
	;
	F_UnlockBuffer(m, v111)
	mBase = m.M
	v165 = m.ExcPending
	if v165 != 0 {
		goto L9
	} else {
		goto L29
	}
L29:
	;
	v166 = *(*int32)(unsafe.Add(mBase, uint32(v30)+20))
	v167 = *(*int32)(unsafe.Add(mBase, uint32(v166)+24))
	v168 = *(*int32)(unsafe.Add(mBase, uint32(v166)+28))
	v169 = *(*int32)(unsafe.Add(mBase, uint32(v166)+32))
	F__hash_finish_split(m, l0, v79, v111, v159, v167, v168, v169)
	mBase = m.M
	v171 = m.ExcPending
	if v171 != 0 {
		goto L9
	} else {
		goto L30
	}
L30:
	;
	F_ReleaseBuffer(m, v111)
	mBase = m.M
	v173 = m.ExcPending
	if v173 != 0 {
		goto L9
	} else {
		goto L31
	}
L31:
	;
	F_ReleaseBuffer(m, v79)
	mBase = m.M
	v175 = m.ExcPending
	if v175 != 0 {
		goto L9
	} else {
		goto L32
	}
L32:
	;
	goto L6
L33:
	;
	F_LockBufferInternal(m, v79, int32(3))
	mBase = m.M
	v750 = m.ExcPending
	if v750 != 0 {
		goto L9
	} else {
		goto L152
	}
L34:
	;
	if base.Ui32(v48) <= base.Ui32(v183-int32(4)) {
		v725 = v111
		goto L33
	} else {
		goto L38
	}
L35:
	;
	v183 = v177
	goto L37
L36:
	;
	v183 = v180
	goto L37
L37:
	;
	goto L34
L38:
	;
	v188 = v79 ^ int32(-1)
	v190 = v79 << (uint(int32(13)) % 32)
	v195 = v111
	v199 = v151
	v207 = v153
	goto L39
L39:
	;
	v218 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v207)+12)))
	if v218&int32(128) == int32(0) {
		goto L41
	} else {
		goto L42
	}
L40:
	;
	v725 = v707
	goto L33
L41:
	;
	v656 = *(*int32)(unsafe.Add(mBase, uint32(v207)+4))
	if v656 != int32(-1) {
		goto L129
	} else {
		goto L130
	}
L42:
	;
	v223 = F_IsBufferCleanupOK(m, v195)
	mBase = m.M
	v224 = m.ExcPending
	if v224 != 0 {
		goto L9
	} else {
		goto L43
	}
L43:
	;
	if v223 == int32(0) {
		goto L41
	} else {
		goto L44
	}
L44:
	;
	v227 = int32(0)
	v228 = base.B2i32(v227 <= v195)
	if v228 == v227 {
		goto L47
	} else {
		goto L48
	}
L45:
	;
	v619 = int32(4)
	v620 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v199)+14)))
	v621 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v199)+12)))
	v622 = v620 - v621
	if v622 <= v619 {
		goto L124
	} else {
		goto L125
	}
L46:
	;
	v247 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v246)+12)))
	if base.Ui32(v247) < base.Ui32(int32(25)) {
		goto L45
	} else {
		goto L50
	}
L47:
	;
	v232 = *(*int32)(unsafe.Add(mBase, _c_F__hash_doinsert[0]))
	v238 = *(*int32)(unsafe.Add(mBase, uint32(v232+(v195^int32(-1))<<(uint(int32(2))%32))))
	v246 = v238
	goto L46
L48:
	;
	goto L49
L49:
	;
	v240 = *(*int32)(unsafe.Add(mBase, _c_F__hash_doinsert[1]))
	v246 = v240 + v195<<(uint(int32(13))%32) + int32(-8192)
	goto L46
L50:
	;
	v253 = int32(base.Ui32(v247+int32(_a_F__hash_doinsert_3)) >> (uint(int32(2)) % 32))
	if v253&int32(_a_F__hash_doinsert_4) == int32(0) {
		goto L45
	} else {
		goto L51
	}
L51:
	;
	v258 = int32(1)
	v260 = v246 + int32(20)
	v261 = int32(0)
	v265 = (v253 + v258) & int32(_a_F__hash_doinsert_4)
	if base.Ui32(int32(3)) <= base.Ui32(v265) {
		goto L53
	} else {
		goto L54
	}
L52:
	;
	if v402 <= int32(0) {
		goto L45
	} else {
		goto L70
	}
L53:
	;
	v268 = int32(2)
	if base.Ui32(v265) <= base.Ui32(v268) {
		goto L56
	} else {
		goto L57
	}
L54:
	;
	v355 = v258
	v359 = v261
	goto L55
L55:
	;
	v380 = *(*int32)(unsafe.Add(mBase, uint32(v260+v355<<(uint(int32(2))%32))))
	v381 = int32(_a_F__hash_doinsert_5)
	if v380&v381 != v381 {
		v402 = v359
		goto L52
	} else {
		goto L69
	}
L56:
	;
	v271 = v268
	goto L58
L57:
	;
	v271 = v265
	goto L58
L58:
	;
	v272 = int32(1)
	v273 = v271 - v272
	v285 = v272
	v289 = v261
	v292 = int32(0)
	goto L59
L59:
	;
	v310 = *(*int32)(unsafe.Add(mBase, uint32(v260+v285<<(uint(int32(2))%32))))
	v311 = int32(_a_F__hash_doinsert_5)
	if v310&v311 == v311 {
		goto L61
	} else {
		goto L62
	}
L60:
	;
	if v273&v272 == int32(0) {
		v402 = v342
		goto L52
	} else {
		goto L68
	}
L61:
	;
	v317 = int32(1)
	*(*uint16)(unsafe.Add(mBase, uint32(v30+int32(32)+v289<<(uint(v317)%32)))) = uint16(v285)
	v323 = v289 + v317
	goto L63
L62:
	;
	v323 = v289
	goto L63
L63:
	;
	v325 = v285 + int32(1)
	v329 = *(*int32)(unsafe.Add(mBase, uint32(v260+v325<<(uint(int32(2))%32))))
	v330 = int32(_a_F__hash_doinsert_5)
	if v329&v330 == v330 {
		goto L64
	} else {
		goto L65
	}
L64:
	;
	v336 = int32(1)
	*(*uint16)(unsafe.Add(mBase, uint32(v30+int32(32)+v323<<(uint(v336)%32)))) = uint16(v325)
	v342 = v323 + v336
	goto L66
L65:
	;
	v342 = v323
	goto L66
L66:
	;
	v343 = int32(2)
	v344 = v285 + v343
	v346 = v292 + v343
	if v346 != v273&int32(-2) {
		v285 = v344
		v289 = v342
		v292 = v346
		goto L59
	} else {
		goto L67
	}
L67:
	;
	goto L60
L68:
	;
	v355 = v344
	v359 = v342
	goto L55
L69:
	;
	v387 = int32(1)
	*(*uint16)(unsafe.Add(mBase, uint32(v30+int32(32)+v359<<(uint(v387)%32)))) = uint16(v355)
	v402 = v359 + v387
	goto L52
L70:
	;
	v423 = v30 + int32(32)
	v424 = F_index_compute_xid_horizon_for_tuples(m, l0, l2, v195, v423, v402)
	mBase = m.M
	v425 = m.ExcPending
	if v425 != 0 {
		goto L9
	} else {
		goto L71
	}
L71:
	;
	F_LockBufferInternal(m, v79, int32(3))
	mBase = m.M
	v428 = m.ExcPending
	if v428 != 0 {
		goto L9
	} else {
		goto L72
	}
L72:
	;
	v429 = int32(_a_F__hash_doinsert_6)
	v431 = *(*int32)(unsafe.Add(mBase, _c_F__hash_doinsert[4]))
	*(*int32)(unsafe.Add(mBase, _c_F__hash_doinsert[4])) = v431 + int32(1)
	F_PageIndexMultiDelete(m, v246, v423, v402)
	mBase = m.M
	v436 = m.ExcPending
	if v436 != 0 {
		goto L9
	} else {
		goto L73
	}
L73:
	;
	v437 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v246)+16)))
	v438 = v246 + v437
	v439 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v438)+12)))
	v441 = v439 & int32(_a_F__hash_doinsert_7)
	*(*uint16)(unsafe.Add(mBase, uint32(v438)+12)) = uint16(v441)
	v443 = int32(0)
	v444 = base.B2i32(v443 <= v79)
	if v444 == v443 {
		goto L75
	} else {
		goto L76
	}
L74:
	;
	v459 = *(*float64)(unsafe.Add(mBase, uint32(v458)+32))
	*(*float64)(unsafe.Add(mBase, uint32(v458)+32)) = base.F64_sub(v459, base.F64_convert_i32_u(v402))
	F_MarkBufferDirty(m, v195)
	mBase = m.M
	v464 = m.ExcPending
	if v464 != 0 {
		goto L9
	} else {
		goto L78
	}
L75:
	;
	v448 = *(*int32)(unsafe.Add(mBase, _c_F__hash_doinsert[0]))
	v452 = *(*int32)(unsafe.Add(mBase, uint32(v448+v188<<(uint(int32(2))%32))))
	v458 = v452
	goto L74
L76:
	;
	goto L77
L77:
	;
	v454 = *(*int32)(unsafe.Add(mBase, _c_F__hash_doinsert[1]))
	v458 = v454 + v190 + int32(-8192)
	goto L74
L78:
	;
	F_MarkBufferDirty(m, v79)
	mBase = m.M
	v466 = m.ExcPending
	if v466 != 0 {
		goto L9
	} else {
		goto L79
	}
L79:
	;
	v467 = *(*int32)(unsafe.Add(mBase, uint32(l0)+48))
	v468 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v467)+118)))
	if v468 != int32(112) {
		goto L81
	} else {
		goto L82
	}
L80:
	;
	if v228 == int32(0) {
		goto L115
	} else {
		goto L116
	}
L81:
	;
	v543 = F_XLogGetFakeLSN(m, l0)
	mBase = m.M
	v544 = m.ExcPending
	if v544 != 0 {
		goto L9
	} else {
		goto L113
	}
L82:
	;
	v472 = *(*int32)(unsafe.Add(mBase, _c_F__hash_doinsert[5]))
	v473 = int32(0)
	v474 = base.B2i32(v473 < v472)
	if v474 == v473 {
		goto L83
	} else {
		goto L84
	}
L83:
	;
	v477 = *(*int32)(unsafe.Add(mBase, uint32(l0)+32))
	if v477 != 0 {
		goto L81
	} else {
		goto L86
	}
L84:
	;
	goto L85
L85:
	;
	if int32(1) < v472 {
		goto L89
	} else {
		goto L90
	}
L86:
	;
	v478 = *(*int32)(unsafe.Add(mBase, uint32(l0)+40))
	if v478 != 0 {
		goto L81
	} else {
		goto L87
	}
L87:
	;
	goto L85
L88:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v30)+24)) = v424
	*(*uint8)(unsafe.Add(mBase, uint32(v30)+30)) = uint8(v511)
	*(*uint16)(unsafe.Add(mBase, uint32(v30)+28)) = uint16(v402)
	F_XLogBeginInsert(m)
	mBase = m.M
	v517 = m.ExcPending
	if v517 != 0 {
		goto L9
	} else {
		goto L107
	}
L89:
	;
	v486 = *(*int32)(unsafe.Add(mBase, uint32(l2)+48))
	v487 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v486)+118)))
	if v487 != int32(112) {
		goto L92
	} else {
		goto L93
	}
L90:
	;
	v482 = int32(*(*uint8)(unsafe.Add(mBase, _c_F__hash_doinsert[6])))
	if v482&int32(1) != 0 {
		goto L89
	} else {
		goto L91
	}
L91:
	;
	v511 = int32(0)
	goto L88
L92:
	;
	v511 = int32(0)
	goto L88
L93:
	;
	goto L94
L94:
	;
	if v473 < v472 {
		goto L95
	} else {
		goto L96
	}
L95:
	;
	v498 = *(*int32)(unsafe.Add(mBase, uint32(l2)+56))
	goto L101
L96:
	;
	v491 = *(*int32)(unsafe.Add(mBase, uint32(l2)+32))
	if v491 != 0 {
		goto L97
	} else {
		goto L98
	}
L97:
	;
	v511 = int32(0)
	goto L88
L98:
	;
	goto L99
L99:
	;
	v493 = *(*int32)(unsafe.Add(mBase, uint32(l2)+40))
	if v493 == int32(0) {
		goto L95
	} else {
		goto L100
	}
L100:
	;
	v511 = int32(0)
	goto L88
L101:
	;
	if base.Ui32(v498) < base.Ui32(int32(_a_F__hash_doinsert_8)) {
		v511 = int32(1)
		goto L88
	} else {
		goto L102
	}
L102:
	;
	v501 = *(*int32)(unsafe.Add(mBase, uint32(l2)+180))
	if v501 == int32(0) {
		goto L103
	} else {
		goto L104
	}
L103:
	;
	v511 = int32(0)
	goto L88
L104:
	;
	goto L105
L105:
	;
	v506 = *(*int32)(unsafe.Add(mBase, uint32(l2)+48))
	v507 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v506)+119)))
	switch v507 - int32(109) {
	case 0, 5:
		goto L106
	default:
		v511 = int32(0)
		goto L88
	}
L106:
	;
	v510 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v501)+112)))
	v511 = v510
	goto L88
L107:
	;
	F_XLogRegisterBuffer(m, int32(0), v195, int32(8))
	mBase = m.M
	v521 = m.ExcPending
	if v521 != 0 {
		goto L9
	} else {
		goto L108
	}
L108:
	;
	F_XLogRegisterData(m, v30+int32(24), int32(8))
	mBase = m.M
	v526 = m.ExcPending
	if v526 != 0 {
		goto L9
	} else {
		goto L109
	}
L109:
	;
	F_XLogRegisterData(m, v30+int32(32), v402<<(uint(int32(1))%32))
	mBase = m.M
	v532 = m.ExcPending
	if v532 != 0 {
		goto L9
	} else {
		goto L110
	}
L110:
	;
	F_XLogRegisterBuffer(m, int32(1), v79, int32(8))
	mBase = m.M
	v536 = m.ExcPending
	if v536 != 0 {
		goto L9
	} else {
		goto L111
	}
L111:
	;
	v539 = F_XLogInsert(m, int32(12), int32(192))
	mBase = m.M
	v540 = m.ExcPending
	if v540 != 0 {
		goto L9
	} else {
		goto L112
	}
L112:
	;
	v547 = v539
	goto L80
L113:
	;
	v547 = v543
	goto L80
L114:
	;
	v567 = base.I64_rotl(v547, int64(32))
	*(*int64)(unsafe.Add(mBase, uint32(v565))) = v567
	if v444 == int32(0) {
		goto L119
	} else {
		goto L120
	}
L115:
	;
	v551 = *(*int32)(unsafe.Add(mBase, _c_F__hash_doinsert[0]))
	v557 = *(*int32)(unsafe.Add(mBase, uint32(v551+(v195^int32(-1))<<(uint(int32(2))%32))))
	v565 = v557
	goto L114
L116:
	;
	goto L117
L117:
	;
	v559 = *(*int32)(unsafe.Add(mBase, _c_F__hash_doinsert[1]))
	v565 = v559 + v195<<(uint(int32(13))%32) + int32(-8192)
	goto L114
L118:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v582))) = v567
	v584 = int32(_a_F__hash_doinsert_6)
	v586 = *(*int32)(unsafe.Add(mBase, _c_F__hash_doinsert[4]))
	*(*int32)(unsafe.Add(mBase, _c_F__hash_doinsert[4])) = v586 - int32(1)
	F_UnlockBuffer(m, v79)
	mBase = m.M
	v591 = m.ExcPending
	if v591 != 0 {
		goto L9
	} else {
		goto L122
	}
L119:
	;
	v572 = *(*int32)(unsafe.Add(mBase, _c_F__hash_doinsert[0]))
	v576 = *(*int32)(unsafe.Add(mBase, uint32(v572+v188<<(uint(int32(2))%32))))
	v582 = v576
	goto L118
L120:
	;
	goto L121
L121:
	;
	v578 = *(*int32)(unsafe.Add(mBase, _c_F__hash_doinsert[1]))
	v582 = v578 + v190 + int32(-8192)
	goto L118
L122:
	;
	goto L45
L123:
	;
	if base.Ui32(v48) <= base.Ui32(v625-int32(4)) {
		v725 = v195
		goto L33
	} else {
		goto L127
	}
L124:
	;
	v625 = v619
	goto L126
L125:
	;
	v625 = v622
	goto L126
L126:
	;
	goto L123
L127:
	;
	goto L41
L128:
	;
	v709 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v708)+16)))
	v711 = int32(4)
	v712 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v708)+14)))
	v713 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v708)+12)))
	v714 = v712 - v713
	if v714 <= v711 {
		goto L148
	} else {
		goto L149
	}
L129:
	;
	if v195 != v111 {
		goto L133
	} else {
		goto L134
	}
L130:
	;
	goto L131
L131:
	;
	F_UnlockBuffer(m, v195)
	mBase = m.M
	v686 = m.ExcPending
	if v686 != 0 {
		goto L9
	} else {
		goto L142
	}
L132:
	;
	v666 = F__hash_getbuf(m, l0, v656, int32(3), int32(1))
	mBase = m.M
	v667 = m.ExcPending
	if v667 != 0 {
		goto L9
	} else {
		goto L138
	}
L133:
	;
	F_UnlockReleaseBuffer(m, v195)
	mBase = m.M
	v661 = m.ExcPending
	if v661 != 0 {
		goto L9
	} else {
		goto L136
	}
L134:
	;
	goto L135
L135:
	;
	F_UnlockBuffer(m, v111)
	mBase = m.M
	v663 = m.ExcPending
	if v663 != 0 {
		goto L9
	} else {
		goto L137
	}
L136:
	;
	goto L132
L137:
	;
	goto L132
L138:
	;
	if v666 < int32(0) {
		goto L139
	} else {
		goto L140
	}
L139:
	;
	v671 = *(*int32)(unsafe.Add(mBase, _c_F__hash_doinsert[0]))
	v677 = *(*int32)(unsafe.Add(mBase, uint32(v671+(v666^int32(-1))<<(uint(int32(2))%32))))
	v707 = v666
	v708 = v677
	goto L128
L140:
	;
	goto L141
L141:
	;
	v679 = *(*int32)(unsafe.Add(mBase, _c_F__hash_doinsert[1]))
	v707 = v666
	v708 = v679 + v666<<(uint(int32(13))%32) + int32(-8192)
	goto L128
L142:
	;
	v688 = F__hash_addovflpage(m, l0, v79, v195, base.B2i32(v195 == v111))
	mBase = m.M
	v689 = m.ExcPending
	if v689 != 0 {
		goto L9
	} else {
		goto L143
	}
L143:
	;
	if v688 < int32(0) {
		goto L144
	} else {
		goto L145
	}
L144:
	;
	v693 = *(*int32)(unsafe.Add(mBase, _c_F__hash_doinsert[0]))
	v699 = *(*int32)(unsafe.Add(mBase, uint32(v693+(v688^int32(-1))<<(uint(int32(2))%32))))
	v707 = v688
	v708 = v699
	goto L128
L145:
	;
	goto L146
L146:
	;
	v701 = *(*int32)(unsafe.Add(mBase, _c_F__hash_doinsert[1]))
	v707 = v688
	v708 = v701 + v688<<(uint(int32(13))%32) + int32(-8192)
	goto L128
L147:
	;
	if base.Ui32(v717-int32(4)) < base.Ui32(v48) {
		v195 = v707
		v199 = v708
		v207 = v709 + v708
		goto L39
	} else {
		goto L151
	}
L148:
	;
	v717 = v711
	goto L150
L149:
	;
	v717 = v714
	goto L150
L150:
	;
	goto L147
L151:
	;
	goto L40
L152:
	;
	v751 = int32(_a_F__hash_doinsert_6)
	v753 = *(*int32)(unsafe.Add(mBase, _c_F__hash_doinsert[4]))
	*(*int32)(unsafe.Add(mBase, _c_F__hash_doinsert[4])) = v753 + int32(1)
	v757 = m.G0
	v759 = v757 - int32(16)
	m.G0 = v759
	F__hash_checkpage(m, l0, v725, int32(3))
	mBase = m.M
	v763 = m.ExcPending
	if v763 != 0 {
		goto L9
	} else {
		goto L153
	}
L153:
	;
	if v725 < int32(0) {
		goto L155
	} else {
		goto L156
	}
L154:
	;
	if l3 != 0 {
		goto L159
	} else {
		goto L160
	}
L155:
	;
	v767 = *(*int32)(unsafe.Add(mBase, _c_F__hash_doinsert[0]))
	v773 = *(*int32)(unsafe.Add(mBase, uint32(v767+(v725^int32(-1))<<(uint(int32(2))%32))))
	v781 = v773
	goto L154
L156:
	;
	goto L157
L157:
	;
	v775 = *(*int32)(unsafe.Add(mBase, _c_F__hash_doinsert[1]))
	v781 = v775 + v725<<(uint(int32(13))%32) + int32(-8192)
	goto L154
L158:
	;
	v875 = v873 & int32(_a_F__hash_doinsert_4)
	v877 = F_PageAddItemExtended(m, v781, l1, v48, v875, int32(0))
	mBase = m.M
	v878 = m.ExcPending
	if v878 != 0 {
		goto L9
	} else {
		goto L188
	}
L159:
	;
	v782 = int32(1)
	v783 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v781)+12)))
	if base.Ui32(v783) < base.Ui32(int32(25)) {
		goto L162
	} else {
		goto L163
	}
L160:
	;
	goto L161
L161:
	;
	v795 = int32(*(*int16)(unsafe.Add(mBase, uint32(l1)+6)))
	if int32(0) <= v795 {
		goto L166
	} else {
		goto L167
	}
L162:
	;
	v792 = v782
	goto L164
L163:
	;
	v792 = int32(base.Ui32(v783+int32(_a_F__hash_doinsert_3))>>(uint(int32(2))%32)) + v782
	goto L164
L164:
	;
	v873 = v792
	goto L158
L165:
	;
	v806 = int32(1)
	v808 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v781)+12)))
	if base.Ui32(v808) < base.Ui32(int32(25)) {
		goto L170
	} else {
		goto L171
	}
L166:
	;
	v798 = int32(8)
	goto L168
L167:
	;
	v798 = int32(16)
	goto L168
L168:
	;
	v800 = *(*int32)(unsafe.Add(mBase, uint32(l1+v798)))
	goto L165
L169:
	;
	v873 = v866 & int32(_a_F__hash_doinsert_4)
	goto L158
L170:
	;
	v817 = v806
	goto L172
L171:
	;
	v817 = int32(base.Ui32(v808+int32(_a_F__hash_doinsert_3))>>(uint(int32(2))%32)) + v806
	goto L172
L172:
	;
	if base.Ui32(int32(2)) <= base.Ui32(v817&int32(_a_F__hash_doinsert_4)) {
		goto L173
	} else {
		goto L174
	}
L173:
	;
	v826 = v817
	v827 = v806
	goto L176
L174:
	;
	v866 = v806
	goto L175
L175:
	;
	goto L169
L176:
	;
	v831 = int32(_a_F__hash_doinsert_4)
	v837 = int32(base.Ui32(v826&v831+v827&v831) >> (uint(int32(1)) % 32))
	v841 = *(*int32)(unsafe.Add(mBase, uint32(v781+int32(20)+v837<<(uint(int32(2))%32))))
	v844 = v781 + v841&int32(_a_F__hash_doinsert_9)
	v847 = int32(*(*int16)(unsafe.Add(mBase, uint32(v844)+6)))
	if int32(0) <= v847 {
		goto L178
	} else {
		goto L179
	}
L177:
	;
	v866 = v859
	goto L175
L178:
	;
	v850 = int32(8)
	goto L180
L179:
	;
	v850 = int32(16)
	goto L180
L180:
	;
	v852 = *(*int32)(unsafe.Add(mBase, uint32(v844+v850)))
	v853 = base.B2i32(base.Ui32(v852) < base.Ui32(v800))
	if base.Ui32(v852) < base.Ui32(v800) {
		goto L181
	} else {
		goto L182
	}
L181:
	;
	v854 = v826
	goto L183
L182:
	;
	v854 = v837
	goto L183
L183:
	;
	if base.Ui32(v852) < base.Ui32(v800) {
		goto L184
	} else {
		goto L185
	}
L184:
	;
	v859 = v837 + int32(1)
	goto L186
L185:
	;
	v859 = v827
	goto L186
L186:
	;
	if base.Ui32(v859&int32(_a_F__hash_doinsert_4)) < base.Ui32(v854&int32(_a_F__hash_doinsert_4)) {
		v826 = v854
		v827 = v859
		goto L176
	} else {
		goto L187
	}
L187:
	;
	goto L177
L188:
	;
	if v877 == int32(0) {
		goto L189
	} else {
		goto L190
	}
L189:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v884 = m.ExcPending
	if v884 != 0 {
		goto L9
	} else {
		goto L192
	}
L190:
	;
	goto L191
L191:
	;
	m.G0 = v759 + int32(16)
	F_MarkBufferDirty(m, v725)
	mBase = m.M
	v901 = m.ExcPending
	if v901 != 0 {
		goto L9
	} else {
		goto L195
	}
L192:
	;
	v885 = *(*int32)(unsafe.Add(mBase, uint32(l0)+48))
	*(*int32)(unsafe.Add(mBase, uint32(v759))) = v885 + int32(4)
	F_errmsg_internal(m, int32(_a_F__hash_doinsert_10), v759)
	mBase = m.M
	v891 = m.ExcPending
	if v891 != 0 {
		goto L9
	} else {
		goto L193
	}
L193:
	;
	F_errfinish(m, int32(_a_F__hash_doinsert_11), int32(316), int32(_a_F__hash_doinsert_12))
	mBase = m.M
	v896 = m.ExcPending
	if v896 != 0 {
		goto L9
	} else {
		goto L194
	}
L194:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L195:
	;
	v902 = *(*float64)(unsafe.Add(mBase, uint32(v98)+32))
	v904 = base.F64_add(v902, float64(1))
	*(*float64)(unsafe.Add(mBase, uint32(v98)+32)) = v904
	v906 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v98)+40)))
	v907 = *(*int32)(unsafe.Add(mBase, uint32(v98)+48))
	F_MarkBufferDirty(m, v79)
	mBase = m.M
	v909 = m.ExcPending
	if v909 != 0 {
		goto L9
	} else {
		goto L196
	}
L196:
	;
	v910 = *(*int32)(unsafe.Add(mBase, uint32(l0)+48))
	v911 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v910)+118)))
	if v911 != int32(112) {
		goto L198
	} else {
		goto L199
	}
L197:
	;
	if v725 < int32(0) {
		goto L213
	} else {
		goto L214
	}
L198:
	;
	v946 = F_XLogGetFakeLSN(m, l0)
	mBase = m.M
	v947 = m.ExcPending
	if v947 != 0 {
		goto L9
	} else {
		goto L211
	}
L199:
	;
	v915 = *(*int32)(unsafe.Add(mBase, _c_F__hash_doinsert[5]))
	if v915 <= int32(0) {
		goto L200
	} else {
		goto L201
	}
L200:
	;
	v918 = *(*int32)(unsafe.Add(mBase, uint32(l0)+32))
	if v918 != 0 {
		goto L198
	} else {
		goto L203
	}
L201:
	;
	goto L202
L202:
	;
	*(*uint16)(unsafe.Add(mBase, uint32(v30)+32)) = uint16(v875)
	F_XLogBeginInsert(m)
	mBase = m.M
	v922 = m.ExcPending
	if v922 != 0 {
		goto L9
	} else {
		goto L205
	}
L203:
	;
	v919 = *(*int32)(unsafe.Add(mBase, uint32(l0)+40))
	if v919 != 0 {
		goto L198
	} else {
		goto L204
	}
L204:
	;
	goto L202
L205:
	;
	F_XLogRegisterData(m, v30+int32(32), int32(2))
	mBase = m.M
	v927 = m.ExcPending
	if v927 != 0 {
		goto L9
	} else {
		goto L206
	}
L206:
	;
	F_XLogRegisterBuffer(m, int32(1), v79, int32(8))
	mBase = m.M
	v931 = m.ExcPending
	if v931 != 0 {
		goto L9
	} else {
		goto L207
	}
L207:
	;
	F_XLogRegisterBuffer(m, int32(0), v725, int32(8))
	mBase = m.M
	v935 = m.ExcPending
	if v935 != 0 {
		goto L9
	} else {
		goto L208
	}
L208:
	;
	v937 = int32(*(*uint16)(unsafe.Add(mBase, uint32(l1)+6)))
	F_XLogRegisterBufData(m, int32(0), l1, v937&int32(_a_F__hash_doinsert_1))
	mBase = m.M
	v941 = m.ExcPending
	if v941 != 0 {
		goto L9
	} else {
		goto L209
	}
L209:
	;
	v944 = F_XLogInsert(m, int32(12), int32(32))
	mBase = m.M
	v945 = m.ExcPending
	if v945 != 0 {
		goto L9
	} else {
		goto L210
	}
L210:
	;
	v948 = v944
	goto L197
L211:
	;
	v948 = v946
	goto L197
L212:
	;
	v972 = base.I64_rotl(v948, int64(32))
	*(*int64)(unsafe.Add(mBase, uint32(v970))) = v972
	if v79 < int32(0) {
		goto L217
	} else {
		goto L218
	}
L213:
	;
	v956 = *(*int32)(unsafe.Add(mBase, _c_F__hash_doinsert[0]))
	v962 = *(*int32)(unsafe.Add(mBase, uint32(v956+(v725^int32(-1))<<(uint(int32(2))%32))))
	v970 = v962
	goto L212
L214:
	;
	goto L215
L215:
	;
	v964 = *(*int32)(unsafe.Add(mBase, _c_F__hash_doinsert[1]))
	v970 = v964 + v725<<(uint(int32(13))%32) + int32(-8192)
	goto L212
L216:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v992))) = v972
	v994 = int32(_a_F__hash_doinsert_6)
	v996 = *(*int32)(unsafe.Add(mBase, _c_F__hash_doinsert[4]))
	*(*int32)(unsafe.Add(mBase, _c_F__hash_doinsert[4])) = v996 - int32(1)
	F_UnlockBuffer(m, v79)
	mBase = m.M
	v1001 = m.ExcPending
	if v1001 != 0 {
		goto L9
	} else {
		goto L220
	}
L217:
	;
	v978 = *(*int32)(unsafe.Add(mBase, _c_F__hash_doinsert[0]))
	v984 = *(*int32)(unsafe.Add(mBase, uint32(v978+(v79^int32(-1))<<(uint(int32(2))%32))))
	v992 = v984
	goto L216
L218:
	;
	goto L219
L219:
	;
	v986 = *(*int32)(unsafe.Add(mBase, _c_F__hash_doinsert[1]))
	v992 = v986 + v79<<(uint(int32(13))%32) + int32(-8192)
	goto L216
L220:
	;
	F_UnlockReleaseBuffer(m, v725)
	mBase = m.M
	v1003 = m.ExcPending
	if v1003 != 0 {
		goto L9
	} else {
		goto L221
	}
L221:
	;
	if v725 != v111 {
		goto L222
	} else {
		goto L223
	}
L222:
	;
	F_ReleaseBuffer(m, v111)
	mBase = m.M
	v1006 = m.ExcPending
	if v1006 != 0 {
		goto L9
	} else {
		goto L225
	}
L223:
	;
	goto L224
L224:
	;
	if base.F64_lt(base.F64_mul(base.F64_convert_i32_u(v906), base.F64_convert_i32_u(v907+int32(1))), v904) != 0 {
		goto L226
	} else {
		goto L227
	}
L225:
	;
	goto L224
L226:
	;
	v1008 = m.G0
	v1012 = (v1008 - int32(_a_F__hash_doinsert_13)) & int32(-4096)
	m.G0 = v1012
	v1015 = v79 << (uint(int32(13)) % 32)
	v1017 = v79 ^ int32(-1)
	goto L231
L227:
	;
	goto L228
L228:
	;
	F_ReleaseBuffer(m, v79)
	mBase = m.M
	v1705 = m.ExcPending
	if v1705 != 0 {
		goto L9
	} else {
		goto L402
	}
L229:
	;
	goto L228
L230:
	;
	F_UnlockBuffer(m, v79)
	mBase = m.M
	v1659 = m.ExcPending
	if v1659 != 0 {
		goto L9
	} else {
		goto L401
	}
L231:
	;
	F_LockBufferInternal(m, v79, int32(3))
	mBase = m.M
	v1049 = m.ExcPending
	if v1049 != 0 {
		goto L9
	} else {
		goto L233
	}
L232:
	;
	v1192 = int32(2)
	v1193 = v1065 + v1192
	v1197 = v1193 - int32(1)
	if base.Ui32(v1192) <= base.Ui32(v1193) {
		goto L284
	} else {
		goto L285
	}
L233:
	;
	F__hash_checkpage(m, l0, v79, int32(8))
	mBase = m.M
	v1052 = m.ExcPending
	if v1052 != 0 {
		goto L9
	} else {
		goto L234
	}
L234:
	;
	if v79 < int32(0) {
		goto L236
	} else {
		goto L237
	}
L235:
	;
	v1065 = *(*int32)(unsafe.Add(mBase, uint32(v1064)+48))
	if base.Ui32(int32(2147483645)) < base.Ui32(v1065) {
		goto L230
	} else {
		goto L239
	}
L236:
	;
	v1056 = *(*int32)(unsafe.Add(mBase, _c_F__hash_doinsert[0]))
	v1058 = *(*int32)(unsafe.Add(mBase, uint32(v1056+v1017<<(uint(int32(2))%32))))
	v1064 = v1058
	goto L235
L237:
	;
	goto L238
L238:
	;
	v1060 = *(*int32)(unsafe.Add(mBase, _c_F__hash_doinsert[1]))
	v1064 = v1060 + v1015 + int32(-8192)
	goto L235
L239:
	;
	v1068 = *(*float64)(unsafe.Add(mBase, uint32(v1064)+32))
	v1069 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v1064)+40)))
	v1072 = v1065 + int32(1)
	if base.F64_le(v1068, base.F64_mul(base.F64_convert_i32_u(v1069), base.F64_convert_i32_u(v1072))) != 0 {
		goto L230
	} else {
		goto L240
	}
L240:
	;
	v1076 = *(*int32)(unsafe.Add(mBase, uint32(v1064)+56))
	v1077 = v1076 & v1072
	if v1077 != 0 {
		goto L242
	} else {
		goto L243
	}
L241:
	;
	if v1138 == int32(0) {
		goto L230
	} else {
		goto L265
	}
L242:
	;
	v1078 = int32(1)
	v1079 = v1077 + v1078
	v1083 = v1079 - v1078
	if base.Ui32(int32(2)) <= base.Ui32(v1079) {
		goto L246
	} else {
		goto L247
	}
L243:
	;
	v1110 = int32(1)
	goto L244
L244:
	;
	if v1110 != int32(-1) {
		goto L252
	} else {
		goto L253
	}
L245:
	;
	v1106 = *(*int32)(unsafe.Add(mBase, uint32(v1064+v1102<<(uint(int32(2))%32))+72))
	v1110 = v1106 + v1079
	goto L244
L246:
	;
	v1089 = int32(32) - base.I32_clz(v1083)
	goto L248
L247:
	;
	v1089 = int32(0)
	goto L248
L248:
	;
	if base.Ui32(int32(10)) <= base.Ui32(v1089) {
		goto L249
	} else {
		goto L250
	}
L249:
	;
	v1092 = int32(3)
	v1102 = int32(base.Ui32(v1083)>>(uint(v1089-v1092)%32))&v1092 | v1089<<(uint(int32(2))%32) - int32(30)
	goto L251
L250:
	;
	v1102 = v1089
	goto L251
L251:
	;
	goto L245
L252:
	;
	v1113 = F_ReadBuffer(m, l0, v1110)
	mBase = m.M
	v1114 = m.ExcPending
	if v1114 != 0 {
		goto L9
	} else {
		goto L255
	}
L253:
	;
	goto L254
L254:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v1128 = m.ExcPending
	if v1128 != 0 {
		goto L9
	} else {
		goto L262
	}
L255:
	;
	v1115 = F_ConditionalLockBufferForCleanup(m, v1113)
	mBase = m.M
	v1116 = m.ExcPending
	if v1116 != 0 {
		goto L9
	} else {
		goto L256
	}
L256:
	;
	if v1115 == int32(0) {
		goto L257
	} else {
		goto L258
	}
L257:
	;
	F_ReleaseBuffer(m, v1113)
	mBase = m.M
	v1120 = m.ExcPending
	if v1120 != 0 {
		goto L9
	} else {
		goto L260
	}
L258:
	;
	goto L259
L259:
	;
	F__hash_checkpage(m, l0, v1113, int32(2))
	mBase = m.M
	v1124 = m.ExcPending
	if v1124 != 0 {
		goto L9
	} else {
		goto L261
	}
L260:
	;
	v1138 = int32(0)
	goto L241
L261:
	;
	v1138 = v1113
	goto L241
L262:
	;
	F_errmsg_internal(m, int32(_a_F__hash_doinsert_14), int32(0))
	mBase = m.M
	v1132 = m.ExcPending
	if v1132 != 0 {
		goto L9
	} else {
		goto L263
	}
L263:
	;
	F_errfinish(m, int32(_a_F__hash_doinsert_15), int32(101), int32(_a_F__hash_doinsert_16))
	mBase = m.M
	v1137 = m.ExcPending
	if v1137 != 0 {
		goto L9
	} else {
		goto L264
	}
L264:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L265:
	;
	if v1138 < int32(0) {
		goto L267
	} else {
		goto L268
	}
L266:
	;
	v1159 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v1158)+16)))
	v1161 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v1159+v1158)+12)))
	if v1161&int32(32) != 0 {
		goto L270
	} else {
		goto L271
	}
L267:
	;
	v1144 = *(*int32)(unsafe.Add(mBase, _c_F__hash_doinsert[0]))
	v1150 = *(*int32)(unsafe.Add(mBase, uint32(v1144+(v1138^int32(-1))<<(uint(int32(2))%32))))
	v1158 = v1150
	goto L266
L268:
	;
	goto L269
L269:
	;
	v1152 = *(*int32)(unsafe.Add(mBase, _c_F__hash_doinsert[1]))
	v1158 = v1152 + v1138<<(uint(int32(13))%32) + int32(-8192)
	goto L266
L270:
	;
	v1164 = *(*int32)(unsafe.Add(mBase, uint32(v1064)+56))
	v1165 = *(*int32)(unsafe.Add(mBase, uint32(v1064)+52))
	v1166 = *(*int32)(unsafe.Add(mBase, uint32(v1064)+48))
	F_UnlockBuffer(m, v79)
	mBase = m.M
	v1168 = m.ExcPending
	if v1168 != 0 {
		goto L9
	} else {
		goto L273
	}
L271:
	;
	goto L272
L272:
	;
	if v1161&int32(64) != 0 {
		goto L277
	} else {
		goto L278
	}
L273:
	;
	F_UnlockBuffer(m, v1138)
	mBase = m.M
	v1170 = m.ExcPending
	if v1170 != 0 {
		goto L9
	} else {
		goto L274
	}
L274:
	;
	F__hash_finish_split(m, l0, v79, v1138, v1077, v1166, v1165, v1164)
	mBase = m.M
	v1172 = m.ExcPending
	if v1172 != 0 {
		goto L9
	} else {
		goto L275
	}
L275:
	;
	F_ReleaseBuffer(m, v1138)
	mBase = m.M
	v1174 = m.ExcPending
	if v1174 != 0 {
		goto L9
	} else {
		goto L276
	}
L276:
	;
	goto L231
L277:
	;
	v1177 = *(*int32)(unsafe.Add(mBase, uint32(v1064)+56))
	v1178 = *(*int32)(unsafe.Add(mBase, uint32(v1064)+52))
	v1179 = *(*int32)(unsafe.Add(mBase, uint32(v1064)+48))
	F_UnlockBuffer(m, v79)
	mBase = m.M
	v1181 = m.ExcPending
	if v1181 != 0 {
		goto L9
	} else {
		goto L280
	}
L278:
	;
	goto L279
L279:
	;
	goto L232
L280:
	;
	v1182 = int32(0)
	F_hashbucketcleanup(m, l0, v1077, v1138, v1110, v1182, v1179, v1178, v1177, v1182, v1182, int32(1), v1182, v1182)
	mBase = m.M
	v1189 = m.ExcPending
	if v1189 != 0 {
		goto L9
	} else {
		goto L281
	}
L281:
	;
	F_ReleaseBuffer(m, v1138)
	mBase = m.M
	v1191 = m.ExcPending
	if v1191 != 0 {
		goto L9
	} else {
		goto L282
	}
L282:
	;
	goto L231
L283:
	;
	v1217 = int32(2)
	v1220 = *(*int32)(unsafe.Add(mBase, uint32(v1064+v1216<<(uint(v1217)%32))+72))
	v1221 = v1220 + v1072
	v1222 = int32(1)
	v1223 = v1221 + v1222
	v1227 = v1193 - v1222
	if base.Ui32(v1217) <= base.Ui32(v1193) {
		goto L292
	} else {
		goto L293
	}
L284:
	;
	v1203 = int32(32) - base.I32_clz(v1197)
	goto L286
L285:
	;
	v1203 = int32(0)
	goto L286
L286:
	;
	if base.Ui32(int32(10)) <= base.Ui32(v1203) {
		goto L287
	} else {
		goto L288
	}
L287:
	;
	v1206 = int32(3)
	v1216 = int32(base.Ui32(v1197)>>(uint(v1203-v1206)%32))&v1206 | v1203<<(uint(int32(2))%32) - int32(30)
	goto L289
L288:
	;
	v1216 = v1203
	goto L289
L289:
	;
	goto L283
L290:
	;
	F_UnlockReleaseBuffer(m, v1138)
	mBase = m.M
	v1648 = m.ExcPending
	if v1648 != 0 {
		goto L9
	} else {
		goto L400
	}
L291:
	;
	v1247 = *(*int32)(unsafe.Add(mBase, uint32(v1064)+60))
	if base.Ui32(v1247) < base.Ui32(v1246) {
		goto L298
	} else {
		goto L299
	}
L292:
	;
	v1233 = int32(32) - base.I32_clz(v1227)
	goto L294
L293:
	;
	v1233 = int32(0)
	goto L294
L294:
	;
	if base.Ui32(int32(10)) <= base.Ui32(v1233) {
		goto L295
	} else {
		goto L296
	}
L295:
	;
	v1236 = int32(3)
	v1246 = int32(base.Ui32(v1227)>>(uint(v1233-v1236)%32))&v1236 | v1233<<(uint(int32(2))%32) - int32(30)
	goto L297
L296:
	;
	v1246 = v1233
	goto L297
L297:
	;
	goto L291
L298:
	;
	if base.Ui32(v1246) <= base.Ui32(int32(9)) {
		goto L302
	} else {
		goto L303
	}
L299:
	;
	goto L300
L300:
	;
	v1394 = F__hash_getnewbuf(m, l0, v1223, int32(0))
	mBase = m.M
	v1395 = m.ExcPending
	if v1395 != 0 {
		goto L9
	} else {
		goto L335
	}
L301:
	;
	v1271 = v1270 + v1220
	if base.B2i32(base.Ui32(v1271) < base.Ui32(v1223))|base.B2i32(v1270-v1072^v1221 == int32(-1)) != 0 {
		goto L290
	} else {
		goto L305
	}
L302:
	;
	v1270 = int32(1) << (uint(v1246) % 32)
	goto L301
L303:
	;
	goto L304
L304:
	;
	v1256 = v1246 - int32(10)
	v1257 = int32(2)
	v1259 = int32(512) << (uint(int32(base.Ui32(v1256)>>(uint(v1257)%32))) % 32)
	v1270 = v1259>>(uint(v1257)%32)*(v1256&int32(3)+int32(1)) + v1259
	goto L301
L305:
	;
	v1279 = v1012 + int32(_a_F__hash_doinsert_17)
	v1282 = int32(0)
	if v1282|(v1279&int32(3)|int32(1)) == v1282 {
		goto L308
	} else {
		goto L309
	}
L306:
	;
	v1330 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v1012)+uint32(_c_F__hash_doinsert[7]))))
	v1331 = v1279 + v1330
	*(*int64)(unsafe.Add(mBase, uint32(v1331)+8)) = int64(-36028792723996673)
	*(*int64)(unsafe.Add(mBase, uint32(v1331))) = int64(-1)
	v1336 = *(*int32)(unsafe.Add(mBase, uint32(l0)+48))
	v1337 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1336)+118)))
	if v1337 != int32(112) {
		goto L317
	} else {
		goto L318
	}
L307:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v1012)+uint32(_c_F__hash_doinsert[8]))) = int32(_a_F__hash_doinsert_18)
	v1321 = int32(_a_F__hash_doinsert_19)
	*(*uint16)(unsafe.Add(mBase, uint32(v1012)+uint32(_c_F__hash_doinsert[9]))) = uint16(v1321)
	v1327 = int32(_a_F__hash_doinsert_20)
	*(*uint16)(unsafe.Add(mBase, uint32(v1012)+uint32(_c_F__hash_doinsert[7]))) = uint16(v1327)
	*(*uint16)(unsafe.Add(mBase, uint32(v1012)+uint32(_c_F__hash_doinsert[10]))) = uint16(v1327)
	goto L306
L308:
	;
	goto L311
L309:
	;
	goto L310
L310:
	;
	goto L316
L311:
	;
	v1298 = v1012 + int32(_a_F__hash_doinsert_13)
	v1300 = v1012 + int32(_a_F__hash_doinsert_21)
	if base.Ui32(v1300) < base.Ui32(v1298) {
		goto L312
	} else {
		goto L313
	}
L312:
	;
	v1302 = v1298
	goto L314
L313:
	;
	v1302 = v1300
	goto L314
L314:
	;
	v1307 = (v1279^int32(-1)+v1302)&int32(-4) + int32(4)
	if v1307 == int32(0) {
		goto L307
	} else {
		goto L315
	}
L315:
	;
	base.MemoryFill(m, v1279, int32(0), v1307)
	goto L307
L316:
	;
	base.MemoryFill(m, v1279, int32(0), int32(_a_F__hash_doinsert_22))
	goto L307
L317:
	;
	F_PageSetChecksum(m, v1012+int32(_a_F__hash_doinsert_17), v1271)
	mBase = m.M
	v1355 = m.ExcPending
	if v1355 != 0 {
		goto L9
	} else {
		goto L325
	}
L318:
	;
	v1341 = *(*int32)(unsafe.Add(mBase, _c_F__hash_doinsert[5]))
	if v1341 <= int32(0) {
		goto L319
	} else {
		goto L320
	}
L319:
	;
	v1344 = *(*int32)(unsafe.Add(mBase, uint32(l0)+32))
	if v1344 != 0 {
		goto L317
	} else {
		goto L322
	}
L320:
	;
	goto L321
L321:
	;
	F_log_newpage(m, l0, int32(0), v1271, v1012+int32(_a_F__hash_doinsert_17), int32(1))
	mBase = m.M
	v1351 = m.ExcPending
	if v1351 != 0 {
		goto L9
	} else {
		goto L324
	}
L322:
	;
	v1345 = *(*int32)(unsafe.Add(mBase, uint32(l0)+40))
	if v1345 != 0 {
		goto L317
	} else {
		goto L323
	}
L323:
	;
	goto L321
L324:
	;
	goto L317
L325:
	;
	v1356 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	if v1356 != 0 {
		goto L326
	} else {
		goto L327
	}
L326:
	;
	v1382 = v1356
	goto L328
L327:
	;
	v1357 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v1358 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	*(*int32)(unsafe.Add(mBase, uint32(v1012)+4088)) = v1358
	v1360 = *(*int64)(unsafe.Add(mBase, uint32(l0)))
	*(*int64)(unsafe.Add(mBase, uint32(v1012)+4080)) = v1360
	v1364 = F_smgropen(m, v1012+int32(4080), v1357)
	mBase = m.M
	v1365 = m.ExcPending
	if v1365 != 0 {
		goto L9
	} else {
		goto L329
	}
L328:
	;
	v1383 = int32(0)
	F_smgrextend(m, v1382, v1383, v1271, v1012+int32(_a_F__hash_doinsert_17), v1383)
	mBase = m.M
	v1388 = m.ExcPending
	if v1388 != 0 {
		goto L9
	} else {
		goto L334
	}
L329:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+12)) = v1364
	v1368 = *(*int32)(unsafe.Add(mBase, uint32(v1364)+72))
	if v1368 != 0 {
		goto L331
	} else {
		goto L332
	}
L330:
	;
	v1380 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	v1382 = v1380
	goto L328
L331:
	;
	v1376 = v1368
	goto L333
L332:
	;
	v1369 = *(*int32)(unsafe.Add(mBase, uint32(v1364)+76))
	v1370 = *(*int32)(unsafe.Add(mBase, uint32(v1364)+80))
	*(*int32)(unsafe.Add(mBase, uint32(v1369)+4)) = v1370
	v1372 = *(*int32)(unsafe.Add(mBase, uint32(v1364)+76))
	*(*int32)(unsafe.Add(mBase, uint32(v1370))) = v1372
	v1374 = *(*int32)(unsafe.Add(mBase, uint32(v1364)+72))
	v1376 = v1374
	goto L333
L333:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v1364)+72)) = v1376 + int32(1)
	goto L330
L334:
	;
	goto L300
L335:
	;
	v1396 = F_IsBufferCleanupOK(m, v1394)
	mBase = m.M
	v1397 = m.ExcPending
	if v1397 != 0 {
		goto L9
	} else {
		goto L336
	}
L336:
	;
	if v1396 == int32(0) {
		goto L337
	} else {
		goto L338
	}
L337:
	;
	F_UnlockReleaseBuffer(m, v1138)
	mBase = m.M
	v1401 = m.ExcPending
	if v1401 != 0 {
		goto L9
	} else {
		goto L340
	}
L338:
	;
	goto L339
L339:
	;
	v1405 = v1064 + int32(56)
	v1406 = int32(_a_F__hash_doinsert_6)
	v1408 = *(*int32)(unsafe.Add(mBase, _c_F__hash_doinsert[4]))
	*(*int32)(unsafe.Add(mBase, _c_F__hash_doinsert[4])) = v1408 + int32(1)
	*(*int32)(unsafe.Add(mBase, uint32(v1064)+48)) = v1072
	v1414 = v1064 + int32(52)
	v1415 = *(*int32)(unsafe.Add(mBase, uint32(v1064)+52))
	v1416 = base.B2i32(base.Ui32(v1072) <= base.Ui32(v1415))
	if v1416 == int32(0) {
		goto L342
	} else {
		goto L343
	}
L340:
	;
	F_UnlockReleaseBuffer(m, v1394)
	mBase = m.M
	v1403 = m.ExcPending
	if v1403 != 0 {
		goto L9
	} else {
		goto L341
	}
L341:
	;
	goto L230
L342:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v1405))) = v1415
	*(*int32)(unsafe.Add(mBase, uint32(v1414))) = v1072 | v1415
	goto L344
L343:
	;
	goto L344
L344:
	;
	v1423 = v1064 + int32(60)
	v1424 = *(*int32)(unsafe.Add(mBase, uint32(v1423)))
	v1425 = base.B2i32(base.Ui32(v1246) <= base.Ui32(v1424))
	if v1425 == int32(0) {
		goto L345
	} else {
		goto L346
	}
L345:
	;
	v1429 = v1064 + int32(76)
	v1430 = int32(2)
	v1436 = *(*int32)(unsafe.Add(mBase, uint32(v1429+v1424<<(uint(v1430)%32))))
	*(*int32)(unsafe.Add(mBase, uint32(v1429+v1246<<(uint(v1430)%32)))) = v1436
	*(*int32)(unsafe.Add(mBase, uint32(v1064)+60)) = v1246
	goto L347
L346:
	;
	goto L347
L347:
	;
	F_MarkBufferDirty(m, v79)
	mBase = m.M
	v1441 = m.ExcPending
	if v1441 != 0 {
		goto L9
	} else {
		goto L348
	}
L348:
	;
	v1442 = *(*int32)(unsafe.Add(mBase, uint32(v1064)+48))
	v1443 = *(*int32)(unsafe.Add(mBase, uint32(v1064)+56))
	v1444 = *(*int32)(unsafe.Add(mBase, uint32(v1064)+52))
	if v1138 < int32(0) {
		goto L350
	} else {
		goto L351
	}
L349:
	;
	v1463 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v1462)+16)))
	v1464 = v1463 + v1462
	*(*int32)(unsafe.Add(mBase, uint32(v1464))) = v1442
	v1466 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v1464)+12)))
	v1468 = v1466 | int32(32)
	*(*uint16)(unsafe.Add(mBase, uint32(v1464)+12)) = uint16(v1468)
	F_MarkBufferDirty(m, v1138)
	mBase = m.M
	v1471 = m.ExcPending
	if v1471 != 0 {
		goto L9
	} else {
		goto L353
	}
L350:
	;
	v1448 = *(*int32)(unsafe.Add(mBase, _c_F__hash_doinsert[0]))
	v1454 = *(*int32)(unsafe.Add(mBase, uint32(v1448+(v1138^int32(-1))<<(uint(int32(2))%32))))
	v1462 = v1454
	goto L349
L351:
	;
	goto L352
L352:
	;
	v1456 = *(*int32)(unsafe.Add(mBase, _c_F__hash_doinsert[1]))
	v1462 = v1456 + v1138<<(uint(int32(13))%32) + int32(-8192)
	goto L349
L353:
	;
	if v1394 < int32(0) {
		goto L355
	} else {
		goto L356
	}
L354:
	;
	v1490 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v1489)+16)))
	v1491 = v1490 + v1489
	*(*int32)(unsafe.Add(mBase, uint32(v1491)+12)) = int32(-8388590)
	*(*int32)(unsafe.Add(mBase, uint32(v1491)+8)) = v1072
	*(*int32)(unsafe.Add(mBase, uint32(v1491)+4)) = int32(-1)
	*(*int32)(unsafe.Add(mBase, uint32(v1491))) = v1442
	F_MarkBufferDirty(m, v1394)
	mBase = m.M
	v1499 = m.ExcPending
	if v1499 != 0 {
		goto L9
	} else {
		goto L358
	}
L355:
	;
	v1475 = *(*int32)(unsafe.Add(mBase, _c_F__hash_doinsert[0]))
	v1481 = *(*int32)(unsafe.Add(mBase, uint32(v1475+(v1394^int32(-1))<<(uint(int32(2))%32))))
	v1489 = v1481
	goto L354
L356:
	;
	goto L357
L357:
	;
	v1483 = *(*int32)(unsafe.Add(mBase, _c_F__hash_doinsert[1]))
	v1489 = v1483 + v1394<<(uint(int32(13))%32) + int32(-8192)
	goto L354
L358:
	;
	v1500 = *(*int32)(unsafe.Add(mBase, uint32(l0)+48))
	v1501 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1500)+118)))
	if v1501 != int32(112) {
		goto L360
	} else {
		goto L361
	}
L359:
	;
	if v1138 < int32(0) {
		goto L385
	} else {
		goto L386
	}
L360:
	;
	v1572 = F_XLogGetFakeLSN(m, l0)
	mBase = m.M
	v1573 = m.ExcPending
	if v1573 != 0 {
		goto L9
	} else {
		goto L383
	}
L361:
	;
	v1505 = *(*int32)(unsafe.Add(mBase, _c_F__hash_doinsert[5]))
	if v1505 <= int32(0) {
		goto L362
	} else {
		goto L363
	}
L362:
	;
	v1508 = *(*int32)(unsafe.Add(mBase, uint32(l0)+32))
	if v1508 != 0 {
		goto L360
	} else {
		goto L365
	}
L363:
	;
	goto L364
L364:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v1012)+uint32(_c_F__hash_doinsert[11]))) = v1442
	v1511 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v1464)+12)))
	*(*uint16)(unsafe.Add(mBase, uint32(v1012)+uint32(_c_F__hash_doinsert[12]))) = uint16(v1511)
	v1513 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v1491)+12)))
	v1514 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v1012)+uint32(_c_F__hash_doinsert[13]))) = uint8(v1514)
	*(*uint16)(unsafe.Add(mBase, uint32(v1012)+uint32(_c_F__hash_doinsert[14]))) = uint16(v1513)
	F_XLogBeginInsert(m)
	mBase = m.M
	v1518 = m.ExcPending
	if v1518 != 0 {
		goto L9
	} else {
		goto L367
	}
L365:
	;
	v1509 = *(*int32)(unsafe.Add(mBase, uint32(l0)+40))
	if v1509 != 0 {
		goto L360
	} else {
		goto L366
	}
L366:
	;
	goto L364
L367:
	;
	F_XLogRegisterBuffer(m, int32(0), v1138, int32(8))
	mBase = m.M
	v1522 = m.ExcPending
	if v1522 != 0 {
		goto L9
	} else {
		goto L368
	}
L368:
	;
	F_XLogRegisterBuffer(m, int32(1), v1394, int32(6))
	mBase = m.M
	v1526 = m.ExcPending
	if v1526 != 0 {
		goto L9
	} else {
		goto L369
	}
L369:
	;
	v1527 = int32(2)
	F_XLogRegisterBuffer(m, v1527, v79, int32(8))
	mBase = m.M
	v1531 = m.ExcPending
	if v1531 != 0 {
		goto L9
	} else {
		goto L370
	}
L370:
	;
	if v1416 == int32(0) {
		goto L371
	} else {
		goto L372
	}
L371:
	;
	v1534 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(v1012)+uint32(_c_F__hash_doinsert[13]))) = uint8(v1534)
	F_XLogRegisterBufData(m, int32(2), v1405, int32(4))
	mBase = m.M
	v1539 = m.ExcPending
	if v1539 != 0 {
		goto L9
	} else {
		goto L374
	}
L372:
	;
	v1545 = v1527
	goto L373
L373:
	;
	if v1425 == int32(0) {
		goto L376
	} else {
		goto L377
	}
L374:
	;
	F_XLogRegisterBufData(m, int32(2), v1414, int32(4))
	mBase = m.M
	v1543 = m.ExcPending
	if v1543 != 0 {
		goto L9
	} else {
		goto L375
	}
L375:
	;
	v1545 = int32(3)
	goto L373
L376:
	;
	*(*uint8)(unsafe.Add(mBase, uint32(v1012)+uint32(_c_F__hash_doinsert[13]))) = uint8(v1545)
	F_XLogRegisterBufData(m, int32(2), v1423, int32(4))
	mBase = m.M
	v1552 = m.ExcPending
	if v1552 != 0 {
		goto L9
	} else {
		goto L379
	}
L377:
	;
	goto L378
L378:
	;
	F_XLogRegisterData(m, v1012+int32(_a_F__hash_doinsert_17), int32(9))
	mBase = m.M
	v1567 = m.ExcPending
	if v1567 != 0 {
		goto L9
	} else {
		goto L381
	}
L379:
	;
	v1553 = int32(2)
	v1554 = *(*int32)(unsafe.Add(mBase, uint32(v1064)+60))
	F_XLogRegisterBufData(m, v1553, v1064+v1554<<(uint(v1553)%32)+int32(76), int32(4))
	mBase = m.M
	v1562 = m.ExcPending
	if v1562 != 0 {
		goto L9
	} else {
		goto L380
	}
L380:
	;
	goto L378
L381:
	;
	v1570 = F_XLogInsert(m, int32(12), int32(64))
	mBase = m.M
	v1571 = m.ExcPending
	if v1571 != 0 {
		goto L9
	} else {
		goto L382
	}
L382:
	;
	v1575 = v1570
	goto L359
L383:
	;
	v1575 = v1572
	goto L359
L384:
	;
	v1595 = base.I64_rotl(v1575, int64(32))
	*(*int64)(unsafe.Add(mBase, uint32(v1593))) = v1595
	if v1394 < int32(0) {
		goto L389
	} else {
		goto L390
	}
L385:
	;
	v1579 = *(*int32)(unsafe.Add(mBase, _c_F__hash_doinsert[0]))
	v1585 = *(*int32)(unsafe.Add(mBase, uint32(v1579+(v1138^int32(-1))<<(uint(int32(2))%32))))
	v1593 = v1585
	goto L384
L386:
	;
	goto L387
L387:
	;
	v1587 = *(*int32)(unsafe.Add(mBase, _c_F__hash_doinsert[1]))
	v1593 = v1587 + v1138<<(uint(int32(13))%32) + int32(-8192)
	goto L384
L388:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v1614))) = v1595
	if v79 < int32(0) {
		goto L393
	} else {
		goto L394
	}
L389:
	;
	v1600 = *(*int32)(unsafe.Add(mBase, _c_F__hash_doinsert[0]))
	v1606 = *(*int32)(unsafe.Add(mBase, uint32(v1600+(v1394^int32(-1))<<(uint(int32(2))%32))))
	v1614 = v1606
	goto L388
L390:
	;
	goto L391
L391:
	;
	v1608 = *(*int32)(unsafe.Add(mBase, _c_F__hash_doinsert[1]))
	v1614 = v1608 + v1394<<(uint(int32(13))%32) + int32(-8192)
	goto L388
L392:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v1629))) = v1595
	v1631 = int32(_a_F__hash_doinsert_6)
	v1633 = *(*int32)(unsafe.Add(mBase, _c_F__hash_doinsert[4]))
	*(*int32)(unsafe.Add(mBase, _c_F__hash_doinsert[4])) = v1633 - int32(1)
	F_UnlockBuffer(m, v79)
	mBase = m.M
	v1638 = m.ExcPending
	if v1638 != 0 {
		goto L9
	} else {
		goto L396
	}
L393:
	;
	v1619 = *(*int32)(unsafe.Add(mBase, _c_F__hash_doinsert[0]))
	v1623 = *(*int32)(unsafe.Add(mBase, uint32(v1619+v1017<<(uint(int32(2))%32))))
	v1629 = v1623
	goto L392
L394:
	;
	goto L395
L395:
	;
	v1625 = *(*int32)(unsafe.Add(mBase, _c_F__hash_doinsert[1]))
	v1629 = v1625 + v1015 + int32(-8192)
	goto L392
L396:
	;
	F__hash_splitbucket(m, l0, v79, v1077, v1072, v1138, v1394, int32(0), v1442, v1444, v1443)
	mBase = m.M
	v1641 = m.ExcPending
	if v1641 != 0 {
		goto L9
	} else {
		goto L397
	}
L397:
	;
	F_ReleaseBuffer(m, v1138)
	mBase = m.M
	v1643 = m.ExcPending
	if v1643 != 0 {
		goto L9
	} else {
		goto L398
	}
L398:
	;
	F_ReleaseBuffer(m, v1394)
	mBase = m.M
	v1645 = m.ExcPending
	if v1645 != 0 {
		goto L9
	} else {
		goto L399
	}
L399:
	;
	m.G0 = v1008
	goto L229
L400:
	;
	goto L230
L401:
	;
	m.G0 = v1008
	goto L229
L402:
	;
	m.G0 = v30 + int32(_a_F__hash_doinsert_0)
	return
L403:
	;
	F_errcode(m, int32(261))
	mBase = m.M
	v1715 = m.ExcPending
	if v1715 != 0 {
		goto L9
	} else {
		goto L404
	}
L404:
	;
	v1716 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v98)+19)))
	*(*int32)(unsafe.Add(mBase, uint32(v30))) = v48
	*(*int32)(unsafe.Add(mBase, uint32(v30)+4)) = (v1716<<(uint(int32(8))%32) - int32(44)) & int32(-48)
	F_errmsg(m, int32(_a_F__hash_doinsert_23), v30)
	mBase = m.M
	v1727 = m.ExcPending
	if v1727 != 0 {
		goto L9
	} else {
		goto L405
	}
L405:
	;
	F_errhint(m, int32(_a_F__hash_doinsert_24), int32(0))
	mBase = m.M
	v1731 = m.ExcPending
	if v1731 != 0 {
		goto L9
	} else {
		goto L406
	}
L406:
	;
	F_errfinish(m, int32(_a_F__hash_doinsert_11), int32(87), int32(_a_F__hash_doinsert_25))
	mBase = m.M
	v1736 = m.ExcPending
	if v1736 != 0 {
		goto L9
	} else {
		goto L407
	}
L407:
	;
	base.Wasm_trap_unreachable()
	for {
	}
}
func F__hash_finish_split(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32, l5 int32, l6 int32) {
	mBase := m.M
	_ = mBase
	var v19 int32
	_ = v19
	var v21 int32
	_ = v21
	var v26 int32
	_ = v26
	var v33 int32
	_ = v33
	var v34 int32
	_ = v34
	var v38 int32
	_ = v38
	var v39 int32
	_ = v39
	var v43 int32
	_ = v43
	var v49 int32
	_ = v49
	var v51 int32
	_ = v51
	var v57 int32
	_ = v57
	var v58 int32
	_ = v58
	var v59 int32
	_ = v59
	var v67 int32
	_ = v67
	var v68 int32
	_ = v68
	var v70 int32
	_ = v70
	var v76 int32
	_ = v76
	var v79 int32
	_ = v79
	var v82 int32
	_ = v82
	var v92 int32
	_ = v92
	var v96 int32
	_ = v96
	var v98 int32
	_ = v98
	var v101 int32
	_ = v101
	var v109 int32
	_ = v109
	var v113 int32
	_ = v113
	var v122 int32
	_ = v122
	var v123 int32
	_ = v123
	var v126 int32
	_ = v126
	var v129 int32
	_ = v129
	var v131 int32
	_ = v131
	var v135 int32
	_ = v135
	var v141 int32
	_ = v141
	var v143 int32
	_ = v143
	var v149 int32
	_ = v149
	var v150 int32
	_ = v150
	var v152 int32
	_ = v152
	var v156 int32
	_ = v156
	var v175 int32
	_ = v175
	var v189 int32
	_ = v189
	var v196 int32
	_ = v196
	var v197 int32
	_ = v197
	var v219 int32
	_ = v219
	var v222 int32
	_ = v222
	var v224 int32
	_ = v224
	var v227 int32
	_ = v227
	var v228 int32
	_ = v228
	var v231 int32
	_ = v231
	var v232 int32
	_ = v232
	var v236 int32
	_ = v236
	var v240 int32
	_ = v240
	var v246 int32
	_ = v246
	var v248 int32
	_ = v248
	var v254 int32
	_ = v254
	var v255 int32
	_ = v255
	var v257 int32
	_ = v257
	var v259 int32
	_ = v259
	var v261 int32
	_ = v261
	var v264 int32
	_ = v264
	var v271 int32
	_ = v271
	var v275 int32
	_ = v275
	var v280 int32
	_ = v280
	v19 = m.G0
	v21 = v19 + int32(-64)
	m.G0 = v21
	*(*int64)(unsafe.Add(mBase, uint32(v21)+24)) = int64(25769803782)
	v26 = *(*int32)(unsafe.Add(mBase, _c_F__hash_finish_split[0]))
	*(*int32)(unsafe.Add(mBase, uint32(v21)+52)) = v26
	v33 = F_hash_create(m, int32(_a_F__hash_finish_split_0), int64(256), v19+int32(-48), int32(1064))
	mBase = m.M
	v34 = m.ExcPending
	if v34 != 0 {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	return
L2:
	;
	v38 = F__hash_getbuf(m, l0, int32(0), int32(1), int32(8))
	mBase = m.M
	v39 = m.ExcPending
	if v39 != 0 {
		goto L1
	} else {
		goto L4
	}
L3:
	;
	v58 = int32(1)
	v59 = *(*int32)(unsafe.Add(mBase, uint32(v57)+56))
	v67 = v59 + v58 | l3
	v68 = *(*int32)(unsafe.Add(mBase, uint32(v57)+48))
	if base.Ui32(v68) < base.Ui32(v67) {
		goto L8
	} else {
		goto L9
	}
L4:
	;
	if v38 < int32(0) {
		goto L5
	} else {
		goto L6
	}
L5:
	;
	v43 = *(*int32)(unsafe.Add(mBase, _c_F__hash_finish_split[1]))
	v49 = *(*int32)(unsafe.Add(mBase, uint32(v43+(v38^int32(-1))<<(uint(int32(2))%32))))
	v57 = v49
	goto L3
L6:
	;
	goto L7
L7:
	;
	v51 = *(*int32)(unsafe.Add(mBase, _c_F__hash_finish_split[2]))
	v57 = v51 + v38<<(uint(int32(13))%32) + int32(-8192)
	goto L3
L8:
	;
	v70 = int32(base.Ui32(v59)>>(uint(v58)%32)) + v58 | l3
	goto L10
L9:
	;
	v70 = v67
	goto L10
L10:
	;
	if v70 != 0 {
		goto L11
	} else {
		goto L12
	}
L11:
	;
	v76 = v70 + int32(1)
	if base.Ui32(int32(2)) <= base.Ui32(v76) {
		goto L14
	} else {
		goto L15
	}
L12:
	;
	v98 = v58
	goto L13
L13:
	;
	F_UnlockReleaseBuffer(m, v38)
	mBase = m.M
	v101 = m.ExcPending
	if v101 != 0 {
		goto L1
	} else {
		goto L20
	}
L14:
	;
	v79 = int32(32) - base.I32_clz(v70)
	goto L16
L15:
	;
	v79 = int32(0)
	goto L16
L16:
	;
	if base.Ui32(int32(10)) <= base.Ui32(v79) {
		goto L17
	} else {
		goto L18
	}
L17:
	;
	v82 = int32(3)
	v92 = int32(base.Ui32(v70)>>(uint(v79-v82)%32))&v82 | v79<<(uint(int32(2))%32) - int32(30)
	goto L19
L18:
	;
	v92 = v79
	goto L19
L19:
	;
	v96 = *(*int32)(unsafe.Add(mBase, uint32(v57+v92<<(uint(int32(2))%32))+72))
	v98 = v96 + v76
	goto L13
L20:
	;
	v109 = v98
	v113 = int32(0)
	goto L22
L21:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v271 = m.ExcPending
	if v271 != 0 {
		goto L1
	} else {
		goto L64
	}
L22:
	;
	if v109 == int32(-1) {
		goto L21
	} else {
		goto L24
	}
L23:
	;
	v227 = F_ConditionalLockBufferForCleanup(m, l2)
	mBase = m.M
	v228 = m.ExcPending
	if v228 != 0 {
		goto L1
	} else {
		goto L50
	}
L24:
	;
	v122 = F_ReadBuffer(m, l0, v109)
	mBase = m.M
	v123 = m.ExcPending
	if v123 != 0 {
		goto L1
	} else {
		goto L25
	}
L25:
	;
	F_LockBufferInternal(m, v122, int32(1))
	mBase = m.M
	v126 = m.ExcPending
	if v126 != 0 {
		goto L1
	} else {
		goto L26
	}
L26:
	;
	F__hash_checkpage(m, l0, v122, int32(3))
	mBase = m.M
	v129 = m.ExcPending
	if v129 != 0 {
		goto L1
	} else {
		goto L27
	}
L27:
	;
	if v109 == v98 {
		goto L28
	} else {
		goto L29
	}
L28:
	;
	v131 = v122
	goto L30
L29:
	;
	v131 = v113
	goto L30
L30:
	;
	if v122 < int32(0) {
		goto L32
	} else {
		goto L33
	}
L31:
	;
	v150 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v149)+16)))
	v152 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v149)+12)))
	if base.Ui32(v152) < base.Ui32(int32(25)) {
		goto L35
	} else {
		goto L36
	}
L32:
	;
	v135 = *(*int32)(unsafe.Add(mBase, _c_F__hash_finish_split[1]))
	v141 = *(*int32)(unsafe.Add(mBase, uint32(v135+(v122^int32(-1))<<(uint(int32(2))%32))))
	v149 = v141
	goto L31
L33:
	;
	goto L34
L34:
	;
	v143 = *(*int32)(unsafe.Add(mBase, _c_F__hash_finish_split[2]))
	v149 = v143 + v122<<(uint(int32(13))%32) + int32(-8192)
	goto L31
L35:
	;
	v219 = *(*int32)(unsafe.Add(mBase, uint32(v150+v149)+4))
	if v122 == v131 {
		goto L43
	} else {
		goto L44
	}
L36:
	;
	v156 = v152 + int32(_a_F__hash_finish_split_1)
	if v156&int32(_a_F__hash_finish_split_2) == int32(0) {
		goto L35
	} else {
		goto L37
	}
L37:
	;
	v175 = int32(1)
	goto L38
L38:
	;
	v189 = *(*int32)(unsafe.Add(mBase, uint32(v149+int32(20)+v175<<(uint(int32(2))%32))))
	v196 = F_hash_search(m, v33, v149+v189&int32(_a_F__hash_finish_split_3), int32(1), v19+int32(-49))
	mBase = m.M
	v197 = m.ExcPending
	if v197 != 0 {
		goto L1
	} else {
		goto L40
	}
L39:
	;
	goto L35
L40:
	;
	if v175 != int32(base.Ui32(v156)>>(uint(int32(2))%32))&int32(_a_F__hash_finish_split_4) {
		v175 = v175 + int32(1)
		goto L38
	} else {
		goto L41
	}
L41:
	;
	goto L39
L42:
	;
	if v219 != int32(-1) {
		v109 = v219
		v113 = v131
		goto L22
	} else {
		goto L48
	}
L43:
	;
	F_UnlockBuffer(m, v122)
	mBase = m.M
	v222 = m.ExcPending
	if v222 != 0 {
		goto L1
	} else {
		goto L46
	}
L44:
	;
	goto L45
L45:
	;
	F_UnlockReleaseBuffer(m, v122)
	mBase = m.M
	v224 = m.ExcPending
	if v224 != 0 {
		goto L1
	} else {
		goto L47
	}
L46:
	;
	goto L42
L47:
	;
	goto L42
L48:
	;
	goto L23
L49:
	;
	F_hash_destroy(m, v33)
	mBase = m.M
	v264 = m.ExcPending
	if v264 != 0 {
		goto L1
	} else {
		goto L63
	}
L50:
	;
	if v227 == int32(0) {
		goto L49
	} else {
		goto L51
	}
L51:
	;
	v231 = F_ConditionalLockBufferForCleanup(m, v131)
	mBase = m.M
	v232 = m.ExcPending
	if v232 != 0 {
		goto L1
	} else {
		goto L52
	}
L52:
	;
	if v231 == int32(0) {
		goto L53
	} else {
		goto L54
	}
L53:
	;
	F_UnlockBuffer(m, l2)
	mBase = m.M
	v236 = m.ExcPending
	if v236 != 0 {
		goto L1
	} else {
		goto L56
	}
L54:
	;
	goto L55
L55:
	;
	if v131 < int32(0) {
		goto L58
	} else {
		goto L59
	}
L56:
	;
	goto L49
L57:
	;
	v255 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v254)+16)))
	v257 = *(*int32)(unsafe.Add(mBase, uint32(v255+v254)+8))
	F__hash_splitbucket(m, l0, l1, l3, v257, l2, v131, v33, l4, l5, l6)
	mBase = m.M
	v259 = m.ExcPending
	if v259 != 0 {
		goto L1
	} else {
		goto L61
	}
L58:
	;
	v240 = *(*int32)(unsafe.Add(mBase, _c_F__hash_finish_split[1]))
	v246 = *(*int32)(unsafe.Add(mBase, uint32(v240+(v131^int32(-1))<<(uint(int32(2))%32))))
	v254 = v246
	goto L57
L59:
	;
	goto L60
L60:
	;
	v248 = *(*int32)(unsafe.Add(mBase, _c_F__hash_finish_split[2]))
	v254 = v248 + v131<<(uint(int32(13))%32) + int32(-8192)
	goto L57
L61:
	;
	F_ReleaseBuffer(m, v131)
	mBase = m.M
	v261 = m.ExcPending
	if v261 != 0 {
		goto L1
	} else {
		goto L62
	}
L62:
	;
	goto L49
L63:
	;
	m.G0 = v21 - int32(-64)
	return
L64:
	;
	F_errmsg_internal(m, int32(_a_F__hash_finish_split_5), int32(0))
	mBase = m.M
	v275 = m.ExcPending
	if v275 != 0 {
		goto L1
	} else {
		goto L65
	}
L65:
	;
	F_errfinish(m, int32(_a_F__hash_finish_split_6), int32(75), int32(_a_F__hash_finish_split_7))
	mBase = m.M
	v280 = m.ExcPending
	if v280 != 0 {
		goto L1
	} else {
		goto L66
	}
L66:
	;
	base.Wasm_trap_unreachable()
	for {
	}
}
func F__hash_first(m *base.Module, l0 int32, l1 int32) int32 {
	mBase := m.M
	_ = mBase
	var v13 int32
	_ = v13
	var v15 int32
	_ = v15
	var v17 int32
	_ = v17
	var v18 int32
	_ = v18
	var v19 int32
	_ = v19
	var v22 int32
	_ = v22
	var v28 int32
	_ = v28
	var v29 int32
	_ = v29
	var v30 int32
	_ = v30
	var v31 int64
	_ = v31
	var v36 int32
	_ = v36
	var v37 int64
	_ = v37
	var v41 int32
	_ = v41
	var v42 int32
	_ = v42
	var v45 int32
	_ = v45
	var v46 int32
	_ = v46
	var v49 int64
	_ = v49
	var v50 int32
	_ = v50
	var v51 int32
	_ = v51
	var v52 int32
	_ = v52
	var v54 int32
	_ = v54
	var v56 int32
	_ = v56
	var v57 int32
	_ = v57
	var v58 int32
	_ = v58
	var v59 int32
	_ = v59
	var v60 int64
	_ = v60
	var v61 int32
	_ = v61
	var v63 int32
	_ = v63
	var v65 int32
	_ = v65
	var v67 int32
	_ = v67
	var v68 int32
	_ = v68
	var v70 int32
	_ = v70
	var v71 int32
	_ = v71
	var v77 int32
	_ = v77
	var v78 int32
	_ = v78
	var v88 int32
	_ = v88
	var v93 int32
	_ = v93
	var v94 int32
	_ = v94
	var v95 int32
	_ = v95
	var v96 int64
	_ = v96
	var v97 int32
	_ = v97
	var v104 int32
	_ = v104
	var v108 int32
	_ = v108
	var v109 int32
	_ = v109
	var v114 int32
	_ = v114
	var v120 int32
	_ = v120
	var v122 int32
	_ = v122
	var v123 int32
	_ = v123
	var v128 int32
	_ = v128
	var v129 int32
	_ = v129
	var v130 int32
	_ = v130
	var v132 int32
	_ = v132
	var v136 int32
	_ = v136
	var v142 int32
	_ = v142
	var v144 int32
	_ = v144
	var v150 int32
	_ = v150
	var v151 int32
	_ = v151
	var v152 int32
	_ = v152
	var v154 int32
	_ = v154
	var v156 int32
	_ = v156
	var v161 int32
	_ = v161
	var v168 int32
	_ = v168
	var v172 int32
	_ = v172
	var v173 int32
	_ = v173
	var v177 int32
	_ = v177
	var v183 int32
	_ = v183
	var v185 int32
	_ = v185
	var v191 int32
	_ = v191
	var v193 int32
	_ = v193
	var v194 int32
	_ = v194
	var v207 int32
	_ = v207
	var v211 int32
	_ = v211
	var v215 int32
	_ = v215
	var v217 int32
	_ = v217
	var v222 int32
	_ = v222
	var v225 int32
	_ = v225
	var v226 int32
	_ = v226
	var v229 int32
	_ = v229
	var v232 int32
	_ = v232
	var v236 int32
	_ = v236
	var v242 int32
	_ = v242
	var v244 int32
	_ = v244
	var v250 int32
	_ = v250
	var v251 int32
	_ = v251
	var v252 int32
	_ = v252
	var v254 int32
	_ = v254
	var v257 int32
	_ = v257
	var v259 int32
	_ = v259
	var v261 int32
	_ = v261
	var v264 int32
	_ = v264
	var v273 int32
	_ = v273
	var v283 int32
	_ = v283
	var v286 int32
	_ = v286
	var v289 int32
	_ = v289
	var v297 int32
	_ = v297
	var v298 int32
	_ = v298
	var v299 int32
	_ = v299
	var v312 int32
	_ = v312
	var v317 int32
	_ = v317
	var v318 int32
	_ = v318
	var v321 int32
	_ = v321
	var v324 int32
	_ = v324
	var v325 int32
	_ = v325
	var v327 int32
	_ = v327
	var v332 int32
	_ = v332
	var v349 int32
	_ = v349
	var v352 int32
	_ = v352
	var v356 int32
	_ = v356
	var v361 int32
	_ = v361
	v13 = m.G0
	v15 = v13 - int32(16)
	m.G0 = v15
	v17 = *(*int32)(unsafe.Add(mBase, uint32(l0)+36))
	v18 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v19 = *(*int32)(unsafe.Add(mBase, uint32(v18)+272))
	if v19 == int32(0) {
		goto L2
	} else {
		goto L3
	}
L1:
	;
	v36 = *(*int32)(unsafe.Add(mBase, uint32(l0)+40))
	if v36 != 0 {
		goto L8
	} else {
		goto L9
	}
L2:
	;
	v22 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v18)+268)))
	if v22 != int32(1) {
		goto L1
	} else {
		goto L5
	}
L3:
	;
	v30 = v19
	goto L4
L4:
	;
	v31 = *(*int64)(unsafe.Add(mBase, uint32(v30)+16))
	*(*int64)(unsafe.Add(mBase, uint32(v30)+16)) = v31 + int64(1)
	goto L1
L5:
	;
	F_pgstat_assoc_relation(m, v18)
	mBase = m.M
	v28 = m.ExcPending
	if v28 != 0 {
		goto L6
	} else {
		goto L7
	}
L6:
	;
	return int32(0)
L7:
	;
	v29 = *(*int32)(unsafe.Add(mBase, uint32(v18)+272))
	v30 = v29
	goto L4
L8:
	;
	v37 = *(*int64)(unsafe.Add(mBase, uint32(v36)))
	*(*int64)(unsafe.Add(mBase, uint32(v36))) = v37 + int64(1)
	goto L10
L9:
	;
	goto L10
L10:
	;
	v41 = int32(0)
	v42 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	if v41 < v42 {
		goto L11
	} else {
		goto L12
	}
L11:
	;
	v45 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v46 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v45))))
	if v46&int32(1) != 0 {
		v332 = v41
		goto L14
	} else {
		goto L15
	}
L12:
	;
	goto L13
L13:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v349 = m.ExcPending
	if v349 != 0 {
		goto L6
	} else {
		goto L82
	}
L14:
	;
	m.G0 = v15 + int32(16)
	return v332
L15:
	;
	v49 = *(*int64)(unsafe.Add(mBase, uint32(v45)+48))
	v50 = *(*int32)(unsafe.Add(mBase, uint32(v45)+8))
	if v50 != 0 {
		goto L18
	} else {
		goto L19
	}
L16:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v17))) = v104
	v108 = F__hash_getbucketbuf_from_hashkey(m, v18, v104, int32(1), int32(0))
	mBase = m.M
	v109 = m.ExcPending
	if v109 != 0 {
		goto L6
	} else {
		goto L32
	}
L17:
	;
	v63 = m.G0
	v65 = v63 - int32(16)
	m.G0 = v65
	v67 = *(*int32)(unsafe.Add(mBase, uint32(v18)+208))
	v68 = *(*int32)(unsafe.Add(mBase, uint32(v67)))
	v70 = F_get_opfamily_proc(m, v68, v50, v50, int32(1))
	mBase = m.M
	v71 = m.ExcPending
	if v71 != 0 {
		goto L6
	} else {
		goto L24
	}
L18:
	;
	v51 = *(*int32)(unsafe.Add(mBase, uint32(v18)+212))
	v52 = *(*int32)(unsafe.Add(mBase, uint32(v51)))
	if v50 != v52 {
		goto L17
	} else {
		goto L21
	}
L19:
	;
	goto L20
L20:
	;
	v54 = int32(1)
	v56 = F_index_getprocinfo(m, v18, v54, v54)
	mBase = m.M
	v57 = m.ExcPending
	if v57 != 0 {
		goto L6
	} else {
		goto L22
	}
L21:
	;
	goto L20
L22:
	;
	v58 = *(*int32)(unsafe.Add(mBase, uint32(v18)+248))
	v59 = *(*int32)(unsafe.Add(mBase, uint32(v58)))
	v60 = F_FunctionCall1Coll(m, v56, v59, v49)
	mBase = m.M
	v61 = m.ExcPending
	if v61 != 0 {
		goto L6
	} else {
		goto L23
	}
L23:
	;
	v104 = base.I32_wrap_i64(v60)
	goto L16
L24:
	;
	if v70 == int32(0) {
		goto L25
	} else {
		goto L26
	}
L25:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v77 = m.ExcPending
	if v77 != 0 {
		goto L6
	} else {
		goto L28
	}
L26:
	;
	goto L27
L27:
	;
	v94 = *(*int32)(unsafe.Add(mBase, uint32(v18)+248))
	v95 = *(*int32)(unsafe.Add(mBase, uint32(v94)))
	v96 = F_OidFunctionCall1Coll(m, v70, v95, v49)
	mBase = m.M
	v97 = m.ExcPending
	if v97 != 0 {
		goto L6
	} else {
		goto L31
	}
L28:
	;
	v78 = *(*int32)(unsafe.Add(mBase, uint32(v18)+48))
	*(*int32)(unsafe.Add(mBase, uint32(v65)+8)) = v50
	*(*int32)(unsafe.Add(mBase, uint32(v65)+4)) = v50
	*(*int32)(unsafe.Add(mBase, uint32(v65))) = int32(1)
	*(*int32)(unsafe.Add(mBase, uint32(v65)+12)) = v78 + int32(4)
	F_errmsg_internal(m, int32(_a_F__hash_first_0), v65)
	mBase = m.M
	v88 = m.ExcPending
	if v88 != 0 {
		goto L6
	} else {
		goto L29
	}
L29:
	;
	F_errfinish(m, int32(_a_F__hash_first_1), int32(115), int32(_a_F__hash_first_2))
	mBase = m.M
	v93 = m.ExcPending
	if v93 != 0 {
		goto L6
	} else {
		goto L30
	}
L30:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L31:
	;
	m.G0 = v65 + int32(16)
	v104 = base.I32_wrap_i64(v96)
	goto L16
L32:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v15)+12)) = v108
	if v108 < int32(0) {
		goto L34
	} else {
		goto L35
	}
L33:
	;
	v130 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	F_PredicateLockPage(m, v18, v129, v130)
	mBase = m.M
	v132 = m.ExcPending
	if v132 != 0 {
		goto L6
	} else {
		goto L37
	}
L34:
	;
	v114 = *(*int32)(unsafe.Add(mBase, _c_F__hash_first[0]))
	v120 = *(*int32)(unsafe.Add(mBase, uint32(v114+(v108^int32(-1))*int32(56))+16))
	v129 = v120
	goto L33
L35:
	;
	goto L36
L36:
	;
	v122 = *(*int32)(unsafe.Add(mBase, _c_F__hash_first[1]))
	v123 = int32(56)
	v128 = *(*int32)(unsafe.Add(mBase, uint32(v122+v108*v123-v123)+16))
	v129 = v128
	goto L33
L37:
	;
	if v108 < int32(0) {
		goto L39
	} else {
		goto L40
	}
L38:
	;
	v151 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v150)+16)))
	v152 = v151 + v150
	*(*int32)(unsafe.Add(mBase, uint32(v15)+4)) = v152
	v154 = *(*int32)(unsafe.Add(mBase, uint32(v152)+8))
	*(*int32)(unsafe.Add(mBase, uint32(v17)+4)) = v108
	v156 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v152)+12)))
	if v156&int32(16) == int32(0) {
		v264 = v152
		goto L42
	} else {
		goto L43
	}
L39:
	;
	v136 = *(*int32)(unsafe.Add(mBase, _c_F__hash_first[2]))
	v142 = *(*int32)(unsafe.Add(mBase, uint32(v136+(v108^int32(-1))<<(uint(int32(2))%32))))
	v150 = v142
	goto L38
L40:
	;
	goto L41
L41:
	;
	v144 = *(*int32)(unsafe.Add(mBase, _c_F__hash_first[3]))
	v150 = v144 + v108<<(uint(int32(13))%32) + int32(-8192)
	goto L38
L42:
	;
	if l1 == int32(-1) {
		goto L68
	} else {
		goto L69
	}
L43:
	;
	v161 = int32(-1)
	v168 = v154 & (v161<<(uint(int32(31)-base.I32_clz(v154))%32) ^ v161)
	v172 = F__hash_getbuf(m, v18, int32(0), int32(1), int32(8))
	mBase = m.M
	v173 = m.ExcPending
	if v173 != 0 {
		goto L6
	} else {
		goto L45
	}
L44:
	;
	if v168 != 0 {
		goto L49
	} else {
		goto L50
	}
L45:
	;
	if v172 < int32(0) {
		goto L46
	} else {
		goto L47
	}
L46:
	;
	v177 = *(*int32)(unsafe.Add(mBase, _c_F__hash_first[2]))
	v183 = *(*int32)(unsafe.Add(mBase, uint32(v177+(v172^int32(-1))<<(uint(int32(2))%32))))
	v191 = v183
	goto L44
L47:
	;
	goto L48
L48:
	;
	v185 = *(*int32)(unsafe.Add(mBase, _c_F__hash_first[3]))
	v191 = v185 + v172<<(uint(int32(13))%32) + int32(-8192)
	goto L44
L49:
	;
	v193 = base.I32_clz(v168)
	v194 = int32(32) - v193
	if base.Ui32(int32(512)) <= base.Ui32(v168) {
		goto L52
	} else {
		goto L53
	}
L50:
	;
	v215 = int32(0)
	goto L51
L51:
	;
	F_UnlockReleaseBuffer(m, v172)
	mBase = m.M
	v217 = m.ExcPending
	if v217 != 0 {
		goto L6
	} else {
		goto L55
	}
L52:
	;
	v207 = int32(base.Ui32(v168)>>(uint(int32(29)-v193)%32))&int32(3) | v194<<(uint(int32(2))%32) - int32(30)
	goto L54
L53:
	;
	v207 = v194
	goto L54
L54:
	;
	v211 = *(*int32)(unsafe.Add(mBase, uint32(v191+v207<<(uint(int32(2))%32))+72))
	v215 = v211
	goto L51
L55:
	;
	F_UnlockBuffer(m, v108)
	mBase = m.M
	v222 = m.ExcPending
	if v222 != 0 {
		goto L6
	} else {
		goto L56
	}
L56:
	;
	v225 = F__hash_getbuf(m, v18, v168+v215+int32(1), int32(1), int32(2))
	mBase = m.M
	v226 = m.ExcPending
	if v226 != 0 {
		goto L6
	} else {
		goto L57
	}
L57:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v17)+8)) = v225
	F_UnlockBuffer(m, v225)
	mBase = m.M
	v229 = m.ExcPending
	if v229 != 0 {
		goto L6
	} else {
		goto L58
	}
L58:
	;
	F_LockBufferInternal(m, v108, int32(1))
	mBase = m.M
	v232 = m.ExcPending
	if v232 != 0 {
		goto L6
	} else {
		goto L59
	}
L59:
	;
	if v108 < int32(0) {
		goto L61
	} else {
		goto L62
	}
L60:
	;
	v251 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v250)+16)))
	v252 = v251 + v250
	*(*int32)(unsafe.Add(mBase, uint32(v15)+4)) = v252
	v254 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v252)+12)))
	if v254&int32(16) != 0 {
		goto L64
	} else {
		goto L65
	}
L61:
	;
	v236 = *(*int32)(unsafe.Add(mBase, _c_F__hash_first[2]))
	v242 = *(*int32)(unsafe.Add(mBase, uint32(v236+(v108^int32(-1))<<(uint(int32(2))%32))))
	v250 = v242
	goto L60
L62:
	;
	goto L63
L63:
	;
	v244 = *(*int32)(unsafe.Add(mBase, _c_F__hash_first[3]))
	v250 = v244 + v108<<(uint(int32(13))%32) + int32(-8192)
	goto L60
L64:
	;
	v257 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(v17)+12)) = uint8(v257)
	v264 = v252
	goto L42
L65:
	;
	goto L66
L66:
	;
	v259 = *(*int32)(unsafe.Add(mBase, uint32(v17)+8))
	F_ReleaseBuffer(m, v259)
	mBase = m.M
	v261 = m.ExcPending
	if v261 != 0 {
		goto L6
	} else {
		goto L67
	}
L67:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v17)+8)) = int32(0)
	v264 = v252
	goto L42
L68:
	;
	v273 = v264
	goto L71
L69:
	;
	v312 = v108
	goto L70
L70:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v17)+24)) = v312
	v317 = F__hash_readpage(m, l0, v15+int32(12), l1)
	mBase = m.M
	v318 = m.ExcPending
	if v318 != 0 {
		goto L6
	} else {
		goto L80
	}
L71:
	;
	v283 = *(*int32)(unsafe.Add(mBase, uint32(v273)+4))
	if v283 == int32(-1) {
		goto L74
	} else {
		goto L75
	}
L72:
	;
	v299 = *(*int32)(unsafe.Add(mBase, uint32(v15)+12))
	v312 = v299
	goto L70
L73:
	;
	goto L72
L74:
	;
	v286 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v17)+12)))
	if v286 != int32(1) {
		goto L73
	} else {
		goto L77
	}
L75:
	;
	goto L76
L76:
	;
	F__hash_readnext(m, l0, v15+int32(12), v15+int32(8), v15+int32(4))
	mBase = m.M
	v297 = m.ExcPending
	if v297 != 0 {
		goto L6
	} else {
		goto L79
	}
L77:
	;
	v289 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v17)+13)))
	if v289 != 0 {
		goto L73
	} else {
		goto L78
	}
L78:
	;
	goto L76
L79:
	;
	v298 = *(*int32)(unsafe.Add(mBase, uint32(v15)+4))
	v273 = v298
	goto L71
L80:
	;
	if v317 == int32(0) {
		v332 = int32(0)
		goto L14
	} else {
		goto L81
	}
L81:
	;
	v321 = *(*int32)(unsafe.Add(mBase, uint32(v17)+48))
	v324 = v17 + v321<<(uint(int32(3))%32)
	v325 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v324)+56)))
	*(*uint16)(unsafe.Add(mBase, uint32(l0)+64)) = uint16(v325)
	v327 = *(*int32)(unsafe.Add(mBase, uint32(v324)+52))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+60)) = v327
	v332 = int32(1)
	goto L14
L82:
	;
	F_errcode(m, int32(1088))
	mBase = m.M
	v352 = m.ExcPending
	if v352 != 0 {
		goto L6
	} else {
		goto L83
	}
L83:
	;
	F_errmsg(m, int32(_a_F__hash_first_3), int32(0))
	mBase = m.M
	v356 = m.ExcPending
	if v356 != 0 {
		goto L6
	} else {
		goto L84
	}
L84:
	;
	F_errfinish(m, int32(_a_F__hash_first_4), int32(314), int32(_a_F__hash_first_5))
	mBase = m.M
	v361 = m.ExcPending
	if v361 != 0 {
		goto L6
	} else {
		goto L85
	}
L85:
	;
	base.Wasm_trap_unreachable()
	for {
	}
}
func F__hash_getbuf_with_strategy(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) int32 {
	var v7 int32
	_ = v7
	var v9 int32
	_ = v9
	var v12 int32
	_ = v12
	var v15 int32
	_ = v15
	var v17 int32
	_ = v17
	var v22 int32
	_ = v22
	var v26 int32
	_ = v26
	var v31 int32
	_ = v31
	if l1 != int32(-1) {
		v7 = int32(0)
		v9 = F_ReadBufferExtended(m, l0, v7, l1, v7, l3)
		v12 = m.ExcPending
		if v12 != 0 {
			return int32(0)
		} else {
			F_LockBufferInternal(m, v9, int32(3))
			v15 = m.ExcPending
			if v15 != 0 {
				return int32(0)
			} else {
				F__hash_checkpage(m, l0, v9, l2)
				v17 = m.ExcPending
				if v17 != 0 {
					return int32(0)
				} else {
					return v9
				}
			}
		}
	} else {
		F_errstart_cold(m, int32(21), int32(0))
		v22 = m.ExcPending
		if v22 != 0 {
			return int32(0)
		} else {
			F_errmsg_internal(m, int32(_a_F__hash_getbuf_with_strategy_0), int32(0))
			v26 = m.ExcPending
			if v26 != 0 {
				return int32(0)
			} else {
				F_errfinish(m, int32(_a_F__hash_getbuf_with_strategy_1), int32(246), int32(_a_F__hash_getbuf_with_strategy_2))
				v31 = m.ExcPending
				if v31 != 0 {
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
func F__hash_pgaddmultitup(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32) {
	mBase := m.M
	_ = mBase
	var v11 int32
	_ = v11
	var v13 int32
	_ = v13
	var v17 int32
	_ = v17
	var v21 int32
	_ = v21
	var v27 int32
	_ = v27
	var v29 int32
	_ = v29
	var v35 int32
	_ = v35
	var v40 int32
	_ = v40
	var v51 int32
	_ = v51
	var v52 int32
	_ = v52
	var v53 int32
	_ = v53
	var v59 int32
	_ = v59
	var v62 int32
	_ = v62
	var v64 int32
	_ = v64
	var v70 int32
	_ = v70
	var v72 int32
	_ = v72
	var v81 int32
	_ = v81
	var v90 int32
	_ = v90
	var v91 int32
	_ = v91
	var v95 int32
	_ = v95
	var v101 int32
	_ = v101
	var v105 int32
	_ = v105
	var v108 int32
	_ = v108
	var v111 int32
	_ = v111
	var v114 int32
	_ = v114
	var v116 int32
	_ = v116
	var v117 int32
	_ = v117
	var v118 int32
	_ = v118
	var v123 int32
	_ = v123
	var v130 int32
	_ = v130
	var v135 int32
	_ = v135
	var v137 int32
	_ = v137
	var v145 int32
	_ = v145
	var v146 int32
	_ = v146
	var v150 int32
	_ = v150
	var v155 int32
	_ = v155
	var v156 int32
	_ = v156
	var v162 int32
	_ = v162
	var v167 int32
	_ = v167
	v11 = m.G0
	v13 = v11 - int32(16)
	m.G0 = v13
	F__hash_checkpage(m, l0, l1, int32(3))
	mBase = m.M
	v17 = m.ExcPending
	if v17 != 0 {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	return
L2:
	;
	if l1 < int32(0) {
		goto L4
	} else {
		goto L5
	}
L3:
	;
	if l4 == int32(0) {
		goto L7
	} else {
		goto L8
	}
L4:
	;
	v21 = *(*int32)(unsafe.Add(mBase, _c_F__hash_pgaddmultitup[0]))
	v27 = *(*int32)(unsafe.Add(mBase, uint32(v21+(l1^int32(-1))<<(uint(int32(2))%32))))
	v35 = v27
	goto L3
L5:
	;
	goto L6
L6:
	;
	v29 = *(*int32)(unsafe.Add(mBase, _c_F__hash_pgaddmultitup[1]))
	v35 = v29 + l1<<(uint(int32(13))%32) + int32(-8192)
	goto L3
L7:
	;
	m.G0 = v13 + int32(16)
	return
L8:
	;
	v40 = int32(0)
	goto L9
L9:
	;
	v51 = l2 + v40<<(uint(int32(2))%32)
	v52 = *(*int32)(unsafe.Add(mBase, uint32(v51)))
	v53 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v52)+6)))
	v59 = int32(*(*int16)(unsafe.Add(mBase, uint32(v52)+6)))
	if int32(0) <= v59 {
		goto L13
	} else {
		goto L14
	}
L10:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v155 = m.ExcPending
	if v155 != 0 {
		goto L1
	} else {
		goto L38
	}
L11:
	;
	goto L10
L12:
	;
	v70 = int32(1)
	v72 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v35)+12)))
	if base.Ui32(v72) < base.Ui32(int32(25)) {
		goto L17
	} else {
		goto L18
	}
L13:
	;
	v62 = int32(8)
	goto L15
L14:
	;
	v62 = int32(16)
	goto L15
L15:
	;
	v64 = *(*int32)(unsafe.Add(mBase, uint32(v52+v62)))
	goto L12
L16:
	;
	*(*uint16)(unsafe.Add(mBase, uint32(l3+v40<<(uint(int32(1))%32)))) = uint16(v135)
	v137 = *(*int32)(unsafe.Add(mBase, uint32(v51)))
	v145 = F_PageAddItemExtended(m, v35, v137, (v53&int32(_a_F__hash_pgaddmultitup_0)+int32(7))&int32(_a_F__hash_pgaddmultitup_1), v135, int32(0))
	mBase = m.M
	v146 = m.ExcPending
	if v146 != 0 {
		goto L1
	} else {
		goto L35
	}
L17:
	;
	v81 = v70
	goto L19
L18:
	;
	v81 = int32(base.Ui32(v72+int32(_a_F__hash_pgaddmultitup_2))>>(uint(int32(2))%32)) + v70
	goto L19
L19:
	;
	if base.Ui32(int32(2)) <= base.Ui32(v81&int32(_a_F__hash_pgaddmultitup_3)) {
		goto L20
	} else {
		goto L21
	}
L20:
	;
	v90 = v81
	v91 = v70
	goto L23
L21:
	;
	v130 = v70
	goto L22
L22:
	;
	v135 = v130 & int32(_a_F__hash_pgaddmultitup_3)
	goto L16
L23:
	;
	v95 = int32(_a_F__hash_pgaddmultitup_3)
	v101 = int32(base.Ui32(v90&v95+v91&v95) >> (uint(int32(1)) % 32))
	v105 = *(*int32)(unsafe.Add(mBase, uint32(v35+int32(20)+v101<<(uint(int32(2))%32))))
	v108 = v35 + v105&int32(_a_F__hash_pgaddmultitup_4)
	v111 = int32(*(*int16)(unsafe.Add(mBase, uint32(v108)+6)))
	if int32(0) <= v111 {
		goto L25
	} else {
		goto L26
	}
L24:
	;
	v130 = v123
	goto L22
L25:
	;
	v114 = int32(8)
	goto L27
L26:
	;
	v114 = int32(16)
	goto L27
L27:
	;
	v116 = *(*int32)(unsafe.Add(mBase, uint32(v108+v114)))
	v117 = base.B2i32(base.Ui32(v116) < base.Ui32(v64))
	if base.Ui32(v116) < base.Ui32(v64) {
		goto L28
	} else {
		goto L29
	}
L28:
	;
	v118 = v90
	goto L30
L29:
	;
	v118 = v101
	goto L30
L30:
	;
	if base.Ui32(v116) < base.Ui32(v64) {
		goto L31
	} else {
		goto L32
	}
L31:
	;
	v123 = v101 + int32(1)
	goto L33
L32:
	;
	v123 = v91
	goto L33
L33:
	;
	if base.Ui32(v123&int32(_a_F__hash_pgaddmultitup_3)) < base.Ui32(v118&int32(_a_F__hash_pgaddmultitup_3)) {
		v90 = v118
		v91 = v123
		goto L23
	} else {
		goto L34
	}
L34:
	;
	goto L24
L35:
	;
	if v145 == int32(0) {
		goto L11
	} else {
		goto L36
	}
L36:
	;
	v150 = v40 + int32(1)
	if l4 != v150 {
		v40 = v150
		goto L9
	} else {
		goto L37
	}
L37:
	;
	goto L7
L38:
	;
	v156 = *(*int32)(unsafe.Add(mBase, uint32(l0)+48))
	*(*int32)(unsafe.Add(mBase, uint32(v13))) = v156 + int32(4)
	F_errmsg_internal(m, int32(_a_F__hash_pgaddmultitup_5), v13)
	mBase = m.M
	v162 = m.ExcPending
	if v162 != 0 {
		goto L1
	} else {
		goto L39
	}
L39:
	;
	F_errfinish(m, int32(_a_F__hash_pgaddmultitup_6), int32(356), int32(_a_F__hash_pgaddmultitup_7))
	mBase = m.M
	v167 = m.ExcPending
	if v167 != 0 {
		goto L1
	} else {
		goto L40
	}
L40:
	;
	base.Wasm_trap_unreachable()
	for {
	}
}
func F_estimate_hash_bucket_stats(m *base.Module, l0 int32, l1 int32, l2 float64, l3 int32, l4 int32) {
	mBase := m.M
	_ = mBase
	var v9 int32
	_ = v9
	var v11 int32
	_ = v11
	var v17 int32
	_ = v17
	var v20 int32
	_ = v20
	var v28 int32
	_ = v28
	var v29 int32
	_ = v29
	var v30 int32
	_ = v30
	var v33 int32
	_ = v33
	var v34 float32
	_ = v34
	var v40 int32
	_ = v40
	var v43 int32
	_ = v43
	var v45 int32
	_ = v45
	var v47 int32
	_ = v47
	var v48 int32
	_ = v48
	var v51 int32
	_ = v51
	var v54 float64
	_ = v54
	var v65 int32
	_ = v65
	var v67 int32
	_ = v67
	var v68 int32
	_ = v68
	var v69 float64
	_ = v69
	var v73 int32
	_ = v73
	var v74 int32
	_ = v74
	var v75 int32
	_ = v75
	var v76 int32
	_ = v76
	var v77 float32
	_ = v77
	var v79 float32
	_ = v79
	var v81 int32
	_ = v81
	var v85 int32
	_ = v85
	var v88 int32
	_ = v88
	var v92 int32
	_ = v92
	var v95 int32
	_ = v95
	var v99 int32
	_ = v99
	var v106 float64
	_ = v106
	var v107 float64
	_ = v107
	var v111 int32
	_ = v111
	var v112 float64
	_ = v112
	var v115 float64
	_ = v115
	var v116 int32
	_ = v116
	var v119 int32
	_ = v119
	var v122 float64
	_ = v122
	var v125 int32
	_ = v125
	var v132 float64
	_ = v132
	var v135 float64
	_ = v135
	var v136 int32
	_ = v136
	var v141 float64
	_ = v141
	var v142 int32
	_ = v142
	var v145 float64
	_ = v145
	var v146 float64
	_ = v146
	var v149 float64
	_ = v149
	var v151 int32
	_ = v151
	var v154 float64
	_ = v154
	var v159 float64
	_ = v159
	var v161 float64
	_ = v161
	var v162 float64
	_ = v162
	var v171 float64
	_ = v171
	var v175 float64
	_ = v175
	var v176 float64
	_ = v176
	var v180 float64
	_ = v180
	var v181 float64
	_ = v181
	var v182 float64
	_ = v182
	var v184 float64
	_ = v184
	var v190 int32
	_ = v190
	var v191 int32
	_ = v191
	var v193 int32
	_ = v193
	v9 = m.G0
	v11 = v9 - int32(80)
	m.G0 = v11
	F_examine_variable(m, l0, l1, int32(0), v11+int32(48))
	mBase = m.M
	v17 = m.ExcPending
	if v17 != 0 {
		return
	} else {
		*(*int64)(unsafe.Add(mBase, uint32(l3))) = int64(0)
		v20 = *(*int32)(unsafe.Add(mBase, uint32(v11)+56))
		if v20 == int32(0) {
			v65 = v11 + int32(48)
			v67 = v11 + int32(47)
			v68 = int32(0)
			v69 = float64(0)
			*(*uint8)(unsafe.Add(mBase, uint32(v67))) = uint8(v68)
			v73 = *(*int32)(unsafe.Add(mBase, uint32(v65)+8))
			if v73 != 0 {
				v74 = *(*int32)(unsafe.Add(mBase, uint32(v73)+16))
				v75 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v74)+22)))
				v76 = v74 + v75
				v77 = *(*float32)(unsafe.Add(mBase, uint32(v76)+8))
				v79 = *(*float32)(unsafe.Add(mBase, uint32(v76)+16))
				v106 = base.F64_promote_f32(v79)
				v107 = base.F64_promote_f32(v77)
			} else {
				v81 = *(*int32)(unsafe.Add(mBase, uint32(v65)+16))
				if v81 == int32(16) {
					v106 = float64(2)
					v107 = v69
				} else {
					v85 = *(*int32)(unsafe.Add(mBase, uint32(v65)+4))
					if v85 == int32(0) {
						v92 = *(*int32)(unsafe.Add(mBase, uint32(v65)))
						if v92 == int32(0) {
							v106 = float64(0)
							v107 = v69
						} else {
							v95 = *(*int32)(unsafe.Add(mBase, uint32(v92)))
							if v95 != int32(6) {
								v106 = float64(0)
								v107 = v69
							} else {
								v99 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v92)+8)))
								switch v99 - int32(_a_F_estimate_hash_bucket_stats_0) {
								case 0:
									v106 = float64(1)
									v107 = v69
								default:
									v106 = float64(0)
									v107 = v69
								case 5:
									v106 = float64(-1)
									v107 = v69
								}
							}
						}
					} else {
						v88 = *(*int32)(unsafe.Add(mBase, uint32(v85)+84))
						if v88 != int32(5) {
							v92 = *(*int32)(unsafe.Add(mBase, uint32(v65)))
							if v92 == int32(0) {
								v106 = float64(0)
								v107 = v69
							} else {
								v95 = *(*int32)(unsafe.Add(mBase, uint32(v92)))
								if v95 != int32(6) {
									v106 = float64(0)
									v107 = v69
								} else {
									v99 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v92)+8)))
									switch v99 - int32(_a_F_estimate_hash_bucket_stats_0) {
									case 0:
										v106 = float64(1)
										v107 = v69
									default:
										v106 = float64(0)
										v107 = v69
									case 5:
										v106 = float64(-1)
										v107 = v69
									}
								}
							}
						} else {
							v106 = float64(-1)
							v107 = v69
						}
					}
				}
			}
			v111 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v65)+28)))
			if v111 != 0 {
				v112 = base.F64_neg(base.F64_sub(float64(1), v107))
			} else {
				v112 = v106
			}
			if base.F64_gt(v112, float64(0)) != 0 {
				v115 = F_clamp_row_est(m, v112)
				mBase = m.M
				v141 = v115
			} else {
				v116 = *(*int32)(unsafe.Add(mBase, uint32(v65)+4))
				if v116 == int32(0) {
					v119 = int32(1)
					*(*uint8)(unsafe.Add(mBase, uint32(v67))) = uint8(v119)
					v141 = float64(200)
				} else {
					v122 = *(*float64)(unsafe.Add(mBase, uint32(v116)+128))
					if base.F64_le(v122, float64(0)) != 0 {
						v125 = int32(1)
						*(*uint8)(unsafe.Add(mBase, uint32(v67))) = uint8(v125)
						v141 = float64(200)
					} else {
						if base.F64_lt(v112, float64(0)) != 0 {
							v132 = F_clamp_row_est(m, base.F64_mul(v122, base.F64_neg(v112)))
							mBase = m.M
							v141 = v132
						} else {
							if base.F64_lt(v122, float64(200)) != 0 {
								v135 = F_clamp_row_est(m, v122)
								mBase = m.M
								v141 = v135
							} else {
								v136 = int32(1)
								*(*uint8)(unsafe.Add(mBase, uint32(v67))) = uint8(v136)
								v141 = float64(200)
							}
						}
					}
				}
			}
			v142 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v11)+47)))
			if v142 == int32(1) {
				v145 = float64(0.1)
				v146 = *(*float64)(unsafe.Add(mBase, uint32(l3)))
				if base.F64_lt(v146, v145) != 0 {
					v149 = v145
				} else {
					v149 = v146
				}
				*(*float64)(unsafe.Add(mBase, uint32(l4))) = v149
			} else {
				v151 = *(*int32)(unsafe.Add(mBase, uint32(v11)+52))
				if v151 == int32(0) {
					v176 = v141
				} else {
					v154 = *(*float64)(unsafe.Add(mBase, uint32(v151)+128))
					if base.F64_gt(v154, float64(0)) == int32(0) {
						v176 = v141
					} else {
						v159 = *(*float64)(unsafe.Add(mBase, uint32(v151)+16))
						v161 = base.F64_mul(v141, base.F64_div(v159, v154))
						v162 = float64(1e+100)
						if base.B2i32(base.Ui64(int64(9218868437227405312)) < base.Ui64(base.I64_reinterpret_f64(v161)&int64(9223372036854775807)))|base.F64_gt(v161, v162) != 0 {
							v175 = v162
						} else {
							v171 = float64(1)
							if base.F64_le(v161, v171) != 0 {
								v175 = v171
							} else {
								v175 = base.F64_nearest(v161)
							}
						}
						v176 = v175
					}
				}
				if base.F64_lt(l2, v176) != 0 {
					v180 = l2
				} else {
					v180 = v176
				}
				v181 = base.F64_div(float64(1), v180)
				v182 = *(*float64)(unsafe.Add(mBase, uint32(l3)))
				if base.F64_gt(v181, v182) != 0 {
					v184 = v181
				} else {
					v184 = v182
				}
				*(*float64)(unsafe.Add(mBase, uint32(l4))) = v184
			}
			v190 = *(*int32)(unsafe.Add(mBase, uint32(v11)+56))
			if v190 != 0 {
				v191 = *(*int32)(unsafe.Add(mBase, uint32(v11)+60))
				m.T0[v191].(func(*base.Module, int32))(m, v190)
				mBase = m.M
				v193 = m.ExcPending
				if v193 != 0 {
					return
				} else {
					m.G0 = v11 + int32(80)
					return
				}
			} else {
				m.G0 = v11 + int32(80)
				return
			}
		} else {
			v28 = F_get_attstatsslot(m, v11+int32(8), v20, int32(1), int32(0), int32(2))
			mBase = m.M
			v29 = m.ExcPending
			if v29 != 0 {
				return
			} else {
				if v28 != 0 {
					v30 = *(*int32)(unsafe.Add(mBase, uint32(v11)+32))
					if int32(0) < v30 {
						v33 = *(*int32)(unsafe.Add(mBase, uint32(v11)+28))
						v34 = *(*float32)(unsafe.Add(mBase, uint32(v33)))
						*(*float64)(unsafe.Add(mBase, uint32(l3))) = base.F64_promote_f32(v34)
					} else {
					}
					F_free_attstatsslot(m, v11+int32(8))
					mBase = m.M
					v40 = m.ExcPending
					if v40 != 0 {
						return
					} else {
						v65 = v11 + int32(48)
						v67 = v11 + int32(47)
						v68 = int32(0)
						v69 = float64(0)
						*(*uint8)(unsafe.Add(mBase, uint32(v67))) = uint8(v68)
						v73 = *(*int32)(unsafe.Add(mBase, uint32(v65)+8))
						if v73 != 0 {
							v74 = *(*int32)(unsafe.Add(mBase, uint32(v73)+16))
							v75 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v74)+22)))
							v76 = v74 + v75
							v77 = *(*float32)(unsafe.Add(mBase, uint32(v76)+8))
							v79 = *(*float32)(unsafe.Add(mBase, uint32(v76)+16))
							v106 = base.F64_promote_f32(v79)
							v107 = base.F64_promote_f32(v77)
						} else {
							v81 = *(*int32)(unsafe.Add(mBase, uint32(v65)+16))
							if v81 == int32(16) {
								v106 = float64(2)
								v107 = v69
							} else {
								v85 = *(*int32)(unsafe.Add(mBase, uint32(v65)+4))
								if v85 == int32(0) {
									v92 = *(*int32)(unsafe.Add(mBase, uint32(v65)))
									if v92 == int32(0) {
										v106 = float64(0)
										v107 = v69
									} else {
										v95 = *(*int32)(unsafe.Add(mBase, uint32(v92)))
										if v95 != int32(6) {
											v106 = float64(0)
											v107 = v69
										} else {
											v99 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v92)+8)))
											switch v99 - int32(_a_F_estimate_hash_bucket_stats_0) {
											case 0:
												v106 = float64(1)
												v107 = v69
											default:
												v106 = float64(0)
												v107 = v69
											case 5:
												v106 = float64(-1)
												v107 = v69
											}
										}
									}
								} else {
									v88 = *(*int32)(unsafe.Add(mBase, uint32(v85)+84))
									if v88 != int32(5) {
										v92 = *(*int32)(unsafe.Add(mBase, uint32(v65)))
										if v92 == int32(0) {
											v106 = float64(0)
											v107 = v69
										} else {
											v95 = *(*int32)(unsafe.Add(mBase, uint32(v92)))
											if v95 != int32(6) {
												v106 = float64(0)
												v107 = v69
											} else {
												v99 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v92)+8)))
												switch v99 - int32(_a_F_estimate_hash_bucket_stats_0) {
												case 0:
													v106 = float64(1)
													v107 = v69
												default:
													v106 = float64(0)
													v107 = v69
												case 5:
													v106 = float64(-1)
													v107 = v69
												}
											}
										}
									} else {
										v106 = float64(-1)
										v107 = v69
									}
								}
							}
						}
						v111 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v65)+28)))
						if v111 != 0 {
							v112 = base.F64_neg(base.F64_sub(float64(1), v107))
						} else {
							v112 = v106
						}
						if base.F64_gt(v112, float64(0)) != 0 {
							v115 = F_clamp_row_est(m, v112)
							mBase = m.M
							v141 = v115
						} else {
							v116 = *(*int32)(unsafe.Add(mBase, uint32(v65)+4))
							if v116 == int32(0) {
								v119 = int32(1)
								*(*uint8)(unsafe.Add(mBase, uint32(v67))) = uint8(v119)
								v141 = float64(200)
							} else {
								v122 = *(*float64)(unsafe.Add(mBase, uint32(v116)+128))
								if base.F64_le(v122, float64(0)) != 0 {
									v125 = int32(1)
									*(*uint8)(unsafe.Add(mBase, uint32(v67))) = uint8(v125)
									v141 = float64(200)
								} else {
									if base.F64_lt(v112, float64(0)) != 0 {
										v132 = F_clamp_row_est(m, base.F64_mul(v122, base.F64_neg(v112)))
										mBase = m.M
										v141 = v132
									} else {
										if base.F64_lt(v122, float64(200)) != 0 {
											v135 = F_clamp_row_est(m, v122)
											mBase = m.M
											v141 = v135
										} else {
											v136 = int32(1)
											*(*uint8)(unsafe.Add(mBase, uint32(v67))) = uint8(v136)
											v141 = float64(200)
										}
									}
								}
							}
						}
						v142 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v11)+47)))
						if v142 == int32(1) {
							v145 = float64(0.1)
							v146 = *(*float64)(unsafe.Add(mBase, uint32(l3)))
							if base.F64_lt(v146, v145) != 0 {
								v149 = v145
							} else {
								v149 = v146
							}
							*(*float64)(unsafe.Add(mBase, uint32(l4))) = v149
						} else {
							v151 = *(*int32)(unsafe.Add(mBase, uint32(v11)+52))
							if v151 == int32(0) {
								v176 = v141
							} else {
								v154 = *(*float64)(unsafe.Add(mBase, uint32(v151)+128))
								if base.F64_gt(v154, float64(0)) == int32(0) {
									v176 = v141
								} else {
									v159 = *(*float64)(unsafe.Add(mBase, uint32(v151)+16))
									v161 = base.F64_mul(v141, base.F64_div(v159, v154))
									v162 = float64(1e+100)
									if base.B2i32(base.Ui64(int64(9218868437227405312)) < base.Ui64(base.I64_reinterpret_f64(v161)&int64(9223372036854775807)))|base.F64_gt(v161, v162) != 0 {
										v175 = v162
									} else {
										v171 = float64(1)
										if base.F64_le(v161, v171) != 0 {
											v175 = v171
										} else {
											v175 = base.F64_nearest(v161)
										}
									}
									v176 = v175
								}
							}
							if base.F64_lt(l2, v176) != 0 {
								v180 = l2
							} else {
								v180 = v176
							}
							v181 = base.F64_div(float64(1), v180)
							v182 = *(*float64)(unsafe.Add(mBase, uint32(l3)))
							if base.F64_gt(v181, v182) != 0 {
								v184 = v181
							} else {
								v184 = v182
							}
							*(*float64)(unsafe.Add(mBase, uint32(l4))) = v184
						}
						v190 = *(*int32)(unsafe.Add(mBase, uint32(v11)+56))
						if v190 != 0 {
							v191 = *(*int32)(unsafe.Add(mBase, uint32(v11)+60))
							m.T0[v191].(func(*base.Module, int32))(m, v190)
							mBase = m.M
							v193 = m.ExcPending
							if v193 != 0 {
								return
							} else {
								m.G0 = v11 + int32(80)
								return
							}
						} else {
							m.G0 = v11 + int32(80)
							return
						}
					}
				} else {
					v43 = *(*int32)(unsafe.Add(mBase, uint32(v11)+56))
					v45 = int32(0)
					v47 = F_get_attstatsslot(m, v11+int32(8), v43, int32(2), v45, v45)
					mBase = m.M
					v48 = m.ExcPending
					if v48 != 0 {
						return
					} else {
						if v47 == int32(0) {
						} else {
							v51 = *(*int32)(unsafe.Add(mBase, uint32(v11)+52))
							if v51 == int32(0) {
							} else {
								v54 = *(*float64)(unsafe.Add(mBase, uint32(v51)+128))
								if base.F64_gt(v54, float64(0)) == int32(0) {
								} else {
									*(*float64)(unsafe.Add(mBase, uint32(l3))) = base.F64_div(float64(1), v54)
								}
							}
						}
						v65 = v11 + int32(48)
						v67 = v11 + int32(47)
						v68 = int32(0)
						v69 = float64(0)
						*(*uint8)(unsafe.Add(mBase, uint32(v67))) = uint8(v68)
						v73 = *(*int32)(unsafe.Add(mBase, uint32(v65)+8))
						if v73 != 0 {
							v74 = *(*int32)(unsafe.Add(mBase, uint32(v73)+16))
							v75 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v74)+22)))
							v76 = v74 + v75
							v77 = *(*float32)(unsafe.Add(mBase, uint32(v76)+8))
							v79 = *(*float32)(unsafe.Add(mBase, uint32(v76)+16))
							v106 = base.F64_promote_f32(v79)
							v107 = base.F64_promote_f32(v77)
						} else {
							v81 = *(*int32)(unsafe.Add(mBase, uint32(v65)+16))
							if v81 == int32(16) {
								v106 = float64(2)
								v107 = v69
							} else {
								v85 = *(*int32)(unsafe.Add(mBase, uint32(v65)+4))
								if v85 == int32(0) {
									v92 = *(*int32)(unsafe.Add(mBase, uint32(v65)))
									if v92 == int32(0) {
										v106 = float64(0)
										v107 = v69
									} else {
										v95 = *(*int32)(unsafe.Add(mBase, uint32(v92)))
										if v95 != int32(6) {
											v106 = float64(0)
											v107 = v69
										} else {
											v99 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v92)+8)))
											switch v99 - int32(_a_F_estimate_hash_bucket_stats_0) {
											case 0:
												v106 = float64(1)
												v107 = v69
											default:
												v106 = float64(0)
												v107 = v69
											case 5:
												v106 = float64(-1)
												v107 = v69
											}
										}
									}
								} else {
									v88 = *(*int32)(unsafe.Add(mBase, uint32(v85)+84))
									if v88 != int32(5) {
										v92 = *(*int32)(unsafe.Add(mBase, uint32(v65)))
										if v92 == int32(0) {
											v106 = float64(0)
											v107 = v69
										} else {
											v95 = *(*int32)(unsafe.Add(mBase, uint32(v92)))
											if v95 != int32(6) {
												v106 = float64(0)
												v107 = v69
											} else {
												v99 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v92)+8)))
												switch v99 - int32(_a_F_estimate_hash_bucket_stats_0) {
												case 0:
													v106 = float64(1)
													v107 = v69
												default:
													v106 = float64(0)
													v107 = v69
												case 5:
													v106 = float64(-1)
													v107 = v69
												}
											}
										}
									} else {
										v106 = float64(-1)
										v107 = v69
									}
								}
							}
						}
						v111 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v65)+28)))
						if v111 != 0 {
							v112 = base.F64_neg(base.F64_sub(float64(1), v107))
						} else {
							v112 = v106
						}
						if base.F64_gt(v112, float64(0)) != 0 {
							v115 = F_clamp_row_est(m, v112)
							mBase = m.M
							v141 = v115
						} else {
							v116 = *(*int32)(unsafe.Add(mBase, uint32(v65)+4))
							if v116 == int32(0) {
								v119 = int32(1)
								*(*uint8)(unsafe.Add(mBase, uint32(v67))) = uint8(v119)
								v141 = float64(200)
							} else {
								v122 = *(*float64)(unsafe.Add(mBase, uint32(v116)+128))
								if base.F64_le(v122, float64(0)) != 0 {
									v125 = int32(1)
									*(*uint8)(unsafe.Add(mBase, uint32(v67))) = uint8(v125)
									v141 = float64(200)
								} else {
									if base.F64_lt(v112, float64(0)) != 0 {
										v132 = F_clamp_row_est(m, base.F64_mul(v122, base.F64_neg(v112)))
										mBase = m.M
										v141 = v132
									} else {
										if base.F64_lt(v122, float64(200)) != 0 {
											v135 = F_clamp_row_est(m, v122)
											mBase = m.M
											v141 = v135
										} else {
											v136 = int32(1)
											*(*uint8)(unsafe.Add(mBase, uint32(v67))) = uint8(v136)
											v141 = float64(200)
										}
									}
								}
							}
						}
						v142 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v11)+47)))
						if v142 == int32(1) {
							v145 = float64(0.1)
							v146 = *(*float64)(unsafe.Add(mBase, uint32(l3)))
							if base.F64_lt(v146, v145) != 0 {
								v149 = v145
							} else {
								v149 = v146
							}
							*(*float64)(unsafe.Add(mBase, uint32(l4))) = v149
						} else {
							v151 = *(*int32)(unsafe.Add(mBase, uint32(v11)+52))
							if v151 == int32(0) {
								v176 = v141
							} else {
								v154 = *(*float64)(unsafe.Add(mBase, uint32(v151)+128))
								if base.F64_gt(v154, float64(0)) == int32(0) {
									v176 = v141
								} else {
									v159 = *(*float64)(unsafe.Add(mBase, uint32(v151)+16))
									v161 = base.F64_mul(v141, base.F64_div(v159, v154))
									v162 = float64(1e+100)
									if base.B2i32(base.Ui64(int64(9218868437227405312)) < base.Ui64(base.I64_reinterpret_f64(v161)&int64(9223372036854775807)))|base.F64_gt(v161, v162) != 0 {
										v175 = v162
									} else {
										v171 = float64(1)
										if base.F64_le(v161, v171) != 0 {
											v175 = v171
										} else {
											v175 = base.F64_nearest(v161)
										}
									}
									v176 = v175
								}
							}
							if base.F64_lt(l2, v176) != 0 {
								v180 = l2
							} else {
								v180 = v176
							}
							v181 = base.F64_div(float64(1), v180)
							v182 = *(*float64)(unsafe.Add(mBase, uint32(l3)))
							if base.F64_gt(v181, v182) != 0 {
								v184 = v181
							} else {
								v184 = v182
							}
							*(*float64)(unsafe.Add(mBase, uint32(l4))) = v184
						}
						v190 = *(*int32)(unsafe.Add(mBase, uint32(v11)+56))
						if v190 != 0 {
							v191 = *(*int32)(unsafe.Add(mBase, uint32(v11)+60))
							m.T0[v191].(func(*base.Module, int32))(m, v190)
							mBase = m.M
							v193 = m.ExcPending
							if v193 != 0 {
								return
							} else {
								m.G0 = v11 + int32(80)
								return
							}
						} else {
							m.G0 = v11 + int32(80)
							return
						}
					}
				}
			}
		}
	}
}
func F_hash_bytes(m *base.Module, l0 int32, l1 int32) int32 {
	mBase := m.M
	_ = mBase
	var v8 int32
	_ = v8
	var v15 int32
	_ = v15
	var v16 int32
	_ = v16
	var v17 int32
	_ = v17
	var v18 int32
	_ = v18
	var v19 int32
	_ = v19
	var v21 int32
	_ = v21
	var v22 int32
	_ = v22
	var v23 int32
	_ = v23
	var v25 int32
	_ = v25
	var v26 int32
	_ = v26
	var v28 int32
	_ = v28
	var v30 int32
	_ = v30
	var v34 int32
	_ = v34
	var v35 int32
	_ = v35
	var v36 int32
	_ = v36
	var v37 int32
	_ = v37
	var v41 int32
	_ = v41
	var v45 int32
	_ = v45
	var v49 int32
	_ = v49
	var v50 int32
	_ = v50
	var v51 int32
	_ = v51
	var v52 int32
	_ = v52
	var v56 int32
	_ = v56
	var v57 int32
	_ = v57
	var v58 int32
	_ = v58
	var v60 int32
	_ = v60
	var v63 int32
	_ = v63
	var v64 int32
	_ = v64
	var v65 int32
	_ = v65
	var v66 int32
	_ = v66
	var v67 int32
	_ = v67
	var v71 int32
	_ = v71
	var v75 int32
	_ = v75
	var v76 int32
	_ = v76
	var v80 int32
	_ = v80
	var v81 int32
	_ = v81
	var v85 int32
	_ = v85
	var v86 int32
	_ = v86
	var v88 int32
	_ = v88
	var v90 int32
	_ = v90
	var v94 int32
	_ = v94
	var v95 int32
	_ = v95
	var v99 int32
	_ = v99
	var v100 int32
	_ = v100
	var v102 int32
	_ = v102
	var v103 int32
	_ = v103
	var v105 int32
	_ = v105
	var v109 int32
	_ = v109
	var v110 int32
	_ = v110
	var v114 int32
	_ = v114
	var v115 int32
	_ = v115
	var v117 int32
	_ = v117
	var v118 int32
	_ = v118
	var v119 int32
	_ = v119
	var v120 int32
	_ = v120
	var v121 int32
	_ = v121
	var v123 int32
	_ = v123
	var v124 int32
	_ = v124
	var v125 int32
	_ = v125
	var v127 int32
	_ = v127
	var v128 int32
	_ = v128
	var v130 int32
	_ = v130
	var v132 int32
	_ = v132
	var v136 int32
	_ = v136
	var v137 int32
	_ = v137
	var v138 int32
	_ = v138
	var v139 int32
	_ = v139
	var v143 int32
	_ = v143
	var v147 int32
	_ = v147
	var v151 int32
	_ = v151
	var v152 int32
	_ = v152
	var v153 int32
	_ = v153
	var v154 int32
	_ = v154
	var v158 int32
	_ = v158
	var v159 int32
	_ = v159
	var v160 int32
	_ = v160
	var v162 int32
	_ = v162
	var v165 int32
	_ = v165
	var v166 int32
	_ = v166
	var v167 int32
	_ = v167
	var v168 int32
	_ = v168
	var v169 int32
	_ = v169
	var v173 int32
	_ = v173
	var v177 int32
	_ = v177
	var v178 int32
	_ = v178
	var v182 int32
	_ = v182
	var v183 int32
	_ = v183
	var v187 int32
	_ = v187
	var v188 int32
	_ = v188
	var v192 int32
	_ = v192
	var v193 int32
	_ = v193
	var v194 int32
	_ = v194
	var v198 int32
	_ = v198
	var v199 int32
	_ = v199
	var v200 int32
	_ = v200
	var v204 int32
	_ = v204
	var v205 int32
	_ = v205
	var v206 int32
	_ = v206
	var v208 int32
	_ = v208
	var v209 int32
	_ = v209
	var v210 int32
	_ = v210
	var v214 int32
	_ = v214
	var v215 int32
	_ = v215
	var v216 int32
	_ = v216
	var v217 int32
	_ = v217
	var v221 int32
	_ = v221
	var v222 int32
	_ = v222
	var v223 int32
	_ = v223
	var v224 int32
	_ = v224
	var v228 int32
	_ = v228
	var v229 int32
	_ = v229
	var v230 int32
	_ = v230
	var v231 int32
	_ = v231
	var v235 int32
	_ = v235
	var v236 int32
	_ = v236
	var v237 int32
	_ = v237
	var v240 int32
	_ = v240
	var v242 int32
	_ = v242
	var v246 int32
	_ = v246
	var v250 int32
	_ = v250
	var v254 int32
	_ = v254
	var v258 int32
	_ = v258
	var v262 int32
	_ = v262
	v8 = l1 - int32(1636608432)
	if l0&int32(3) != 0 {
		if base.Ui32(int32(11)) < base.Ui32(l1) {
			v117 = l0
			v118 = l1
			v119 = v8
			v120 = v8
			v121 = v8
			for {
				v123 = *(*int32)(unsafe.Add(mBase, uint32(v117)+4))
				v124 = v123 + v120
				v125 = *(*int32)(unsafe.Add(mBase, uint32(v117)))
				v127 = *(*int32)(unsafe.Add(mBase, uint32(v117)+8))
				v128 = v127 + v121
				v130 = int32(4)
				v132 = v125 + v119 - v128 ^ base.I32_rotl(v128, v130)
				v136 = v124 - v132 ^ base.I32_rotl(v132, int32(6))
				v137 = v128 + v124
				v138 = v132 + v137
				v139 = v136 + v138
				v143 = v137 - v136 ^ base.I32_rotl(v136, int32(8))
				v147 = v138 - v143 ^ base.I32_rotl(v143, int32(16))
				v151 = v139 - v147 ^ base.I32_rotl(v147, int32(19))
				v152 = v143 + v139
				v153 = v147 + v152
				v154 = v151 + v153
				v158 = v152 - v151 ^ base.I32_rotl(v151, v130)
				v159 = int32(12)
				v160 = v117 + v159
				v162 = v118 - v159
				if base.Ui32(int32(11)) < base.Ui32(v162) {
					v117 = v160
					v118 = v162
					v119 = v153
					v120 = v154
					v121 = v158
					continue
				} else {
					break
				}
				break
			}
			v165 = v160
			v166 = v162
			v167 = v153
			v168 = v154
			v169 = v158
		} else {
			v165 = l0
			v166 = l1
			v167 = v8
			v168 = v8
			v169 = v8
		}
		switch v166 - int32(1) {
		case 0:
			v228 = v167
			v229 = v168
			v230 = v169
			v231 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v165))))
			v235 = v228 + v231
			v236 = v229
			v237 = v230
		case 1:
			v221 = v167
			v222 = v168
			v223 = v169
			v224 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v165)+1)))
			v228 = v224<<(uint(int32(8))%32) + v221
			v229 = v222
			v230 = v223
			v231 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v165))))
			v235 = v228 + v231
			v236 = v229
			v237 = v230
		case 2:
			v214 = v167
			v215 = v168
			v216 = v169
			v217 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v165)+2)))
			v221 = v217<<(uint(int32(16))%32) + v214
			v222 = v215
			v223 = v216
			v224 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v165)+1)))
			v228 = v224<<(uint(int32(8))%32) + v221
			v229 = v222
			v230 = v223
			v231 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v165))))
			v235 = v228 + v231
			v236 = v229
			v237 = v230
		case 3:
			v208 = v168
			v209 = v169
			v210 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v165)+3)))
			v214 = v210<<(uint(int32(24))%32) + v167
			v215 = v208
			v216 = v209
			v217 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v165)+2)))
			v221 = v217<<(uint(int32(16))%32) + v214
			v222 = v215
			v223 = v216
			v224 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v165)+1)))
			v228 = v224<<(uint(int32(8))%32) + v221
			v229 = v222
			v230 = v223
			v231 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v165))))
			v235 = v228 + v231
			v236 = v229
			v237 = v230
		case 4:
			v204 = v168
			v205 = v169
			v206 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v165)+4)))
			v208 = v204 + v206
			v209 = v205
			v210 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v165)+3)))
			v214 = v210<<(uint(int32(24))%32) + v167
			v215 = v208
			v216 = v209
			v217 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v165)+2)))
			v221 = v217<<(uint(int32(16))%32) + v214
			v222 = v215
			v223 = v216
			v224 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v165)+1)))
			v228 = v224<<(uint(int32(8))%32) + v221
			v229 = v222
			v230 = v223
			v231 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v165))))
			v235 = v228 + v231
			v236 = v229
			v237 = v230
		case 5:
			v198 = v168
			v199 = v169
			v200 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v165)+5)))
			v204 = v200<<(uint(int32(8))%32) + v198
			v205 = v199
			v206 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v165)+4)))
			v208 = v204 + v206
			v209 = v205
			v210 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v165)+3)))
			v214 = v210<<(uint(int32(24))%32) + v167
			v215 = v208
			v216 = v209
			v217 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v165)+2)))
			v221 = v217<<(uint(int32(16))%32) + v214
			v222 = v215
			v223 = v216
			v224 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v165)+1)))
			v228 = v224<<(uint(int32(8))%32) + v221
			v229 = v222
			v230 = v223
			v231 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v165))))
			v235 = v228 + v231
			v236 = v229
			v237 = v230
		case 6:
			v192 = v168
			v193 = v169
			v194 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v165)+6)))
			v198 = v194<<(uint(int32(16))%32) + v192
			v199 = v193
			v200 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v165)+5)))
			v204 = v200<<(uint(int32(8))%32) + v198
			v205 = v199
			v206 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v165)+4)))
			v208 = v204 + v206
			v209 = v205
			v210 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v165)+3)))
			v214 = v210<<(uint(int32(24))%32) + v167
			v215 = v208
			v216 = v209
			v217 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v165)+2)))
			v221 = v217<<(uint(int32(16))%32) + v214
			v222 = v215
			v223 = v216
			v224 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v165)+1)))
			v228 = v224<<(uint(int32(8))%32) + v221
			v229 = v222
			v230 = v223
			v231 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v165))))
			v235 = v228 + v231
			v236 = v229
			v237 = v230
		case 7:
			v187 = v169
			v188 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v165)+7)))
			v192 = v188<<(uint(int32(24))%32) + v168
			v193 = v187
			v194 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v165)+6)))
			v198 = v194<<(uint(int32(16))%32) + v192
			v199 = v193
			v200 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v165)+5)))
			v204 = v200<<(uint(int32(8))%32) + v198
			v205 = v199
			v206 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v165)+4)))
			v208 = v204 + v206
			v209 = v205
			v210 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v165)+3)))
			v214 = v210<<(uint(int32(24))%32) + v167
			v215 = v208
			v216 = v209
			v217 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v165)+2)))
			v221 = v217<<(uint(int32(16))%32) + v214
			v222 = v215
			v223 = v216
			v224 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v165)+1)))
			v228 = v224<<(uint(int32(8))%32) + v221
			v229 = v222
			v230 = v223
			v231 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v165))))
			v235 = v228 + v231
			v236 = v229
			v237 = v230
		case 8:
			v182 = v169
			v183 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v165)+8)))
			v187 = v183<<(uint(int32(8))%32) + v182
			v188 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v165)+7)))
			v192 = v188<<(uint(int32(24))%32) + v168
			v193 = v187
			v194 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v165)+6)))
			v198 = v194<<(uint(int32(16))%32) + v192
			v199 = v193
			v200 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v165)+5)))
			v204 = v200<<(uint(int32(8))%32) + v198
			v205 = v199
			v206 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v165)+4)))
			v208 = v204 + v206
			v209 = v205
			v210 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v165)+3)))
			v214 = v210<<(uint(int32(24))%32) + v167
			v215 = v208
			v216 = v209
			v217 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v165)+2)))
			v221 = v217<<(uint(int32(16))%32) + v214
			v222 = v215
			v223 = v216
			v224 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v165)+1)))
			v228 = v224<<(uint(int32(8))%32) + v221
			v229 = v222
			v230 = v223
			v231 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v165))))
			v235 = v228 + v231
			v236 = v229
			v237 = v230
		case 9:
			v177 = v169
			v178 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v165)+9)))
			v182 = v178<<(uint(int32(16))%32) + v177
			v183 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v165)+8)))
			v187 = v183<<(uint(int32(8))%32) + v182
			v188 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v165)+7)))
			v192 = v188<<(uint(int32(24))%32) + v168
			v193 = v187
			v194 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v165)+6)))
			v198 = v194<<(uint(int32(16))%32) + v192
			v199 = v193
			v200 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v165)+5)))
			v204 = v200<<(uint(int32(8))%32) + v198
			v205 = v199
			v206 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v165)+4)))
			v208 = v204 + v206
			v209 = v205
			v210 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v165)+3)))
			v214 = v210<<(uint(int32(24))%32) + v167
			v215 = v208
			v216 = v209
			v217 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v165)+2)))
			v221 = v217<<(uint(int32(16))%32) + v214
			v222 = v215
			v223 = v216
			v224 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v165)+1)))
			v228 = v224<<(uint(int32(8))%32) + v221
			v229 = v222
			v230 = v223
			v231 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v165))))
			v235 = v228 + v231
			v236 = v229
			v237 = v230
		case 10:
			v173 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v165)+10)))
			v177 = v173<<(uint(int32(24))%32) + v169
			v178 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v165)+9)))
			v182 = v178<<(uint(int32(16))%32) + v177
			v183 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v165)+8)))
			v187 = v183<<(uint(int32(8))%32) + v182
			v188 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v165)+7)))
			v192 = v188<<(uint(int32(24))%32) + v168
			v193 = v187
			v194 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v165)+6)))
			v198 = v194<<(uint(int32(16))%32) + v192
			v199 = v193
			v200 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v165)+5)))
			v204 = v200<<(uint(int32(8))%32) + v198
			v205 = v199
			v206 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v165)+4)))
			v208 = v204 + v206
			v209 = v205
			v210 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v165)+3)))
			v214 = v210<<(uint(int32(24))%32) + v167
			v215 = v208
			v216 = v209
			v217 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v165)+2)))
			v221 = v217<<(uint(int32(16))%32) + v214
			v222 = v215
			v223 = v216
			v224 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v165)+1)))
			v228 = v224<<(uint(int32(8))%32) + v221
			v229 = v222
			v230 = v223
			v231 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v165))))
			v235 = v228 + v231
			v236 = v229
			v237 = v230
		default:
			v235 = v167
			v236 = v168
			v237 = v169
		}
	} else {
		if base.Ui32(l1) < base.Ui32(int32(12)) {
			v63 = l0
			v64 = l1
			v65 = v8
			v66 = v8
			v67 = v8
		} else {
			v15 = l0
			v16 = l1
			v17 = v8
			v18 = v8
			v19 = v8
			for {
				v21 = *(*int32)(unsafe.Add(mBase, uint32(v15)+4))
				v22 = v21 + v18
				v23 = *(*int32)(unsafe.Add(mBase, uint32(v15)))
				v25 = *(*int32)(unsafe.Add(mBase, uint32(v15)+8))
				v26 = v25 + v19
				v28 = int32(4)
				v30 = v23 + v17 - v26 ^ base.I32_rotl(v26, v28)
				v34 = v22 - v30 ^ base.I32_rotl(v30, int32(6))
				v35 = v26 + v22
				v36 = v30 + v35
				v37 = v34 + v36
				v41 = v35 - v34 ^ base.I32_rotl(v34, int32(8))
				v45 = v36 - v41 ^ base.I32_rotl(v41, int32(16))
				v49 = v37 - v45 ^ base.I32_rotl(v45, int32(19))
				v50 = v41 + v37
				v51 = v45 + v50
				v52 = v49 + v51
				v56 = v50 - v49 ^ base.I32_rotl(v49, v28)
				v57 = int32(12)
				v58 = v15 + v57
				v60 = v16 - v57
				if base.Ui32(int32(11)) < base.Ui32(v60) {
					v15 = v58
					v16 = v60
					v17 = v51
					v18 = v52
					v19 = v56
					continue
				} else {
					break
				}
				break
			}
			v63 = v58
			v64 = v60
			v65 = v51
			v66 = v52
			v67 = v56
		}
		switch v64 - int32(1) {
		case 0:
			v114 = v65
			v115 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v63))))
			v235 = v114 + v115
			v236 = v66
			v237 = v67
		case 1:
			v109 = v65
			v110 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v63)+1)))
			v114 = v110<<(uint(int32(8))%32) + v109
			v115 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v63))))
			v235 = v114 + v115
			v236 = v66
			v237 = v67
		case 2:
			v105 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v63)+2)))
			v109 = v105<<(uint(int32(16))%32) + v65
			v110 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v63)+1)))
			v114 = v110<<(uint(int32(8))%32) + v109
			v115 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v63))))
			v235 = v114 + v115
			v236 = v66
			v237 = v67
		case 3:
			v102 = v66
			v103 = *(*int32)(unsafe.Add(mBase, uint32(v63)))
			v235 = v103 + v65
			v236 = v102
			v237 = v67
		case 4:
			v99 = v66
			v100 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v63)+4)))
			v102 = v99 + v100
			v103 = *(*int32)(unsafe.Add(mBase, uint32(v63)))
			v235 = v103 + v65
			v236 = v102
			v237 = v67
		case 5:
			v94 = v66
			v95 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v63)+5)))
			v99 = v95<<(uint(int32(8))%32) + v94
			v100 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v63)+4)))
			v102 = v99 + v100
			v103 = *(*int32)(unsafe.Add(mBase, uint32(v63)))
			v235 = v103 + v65
			v236 = v102
			v237 = v67
		case 6:
			v90 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v63)+6)))
			v94 = v90<<(uint(int32(16))%32) + v66
			v95 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v63)+5)))
			v99 = v95<<(uint(int32(8))%32) + v94
			v100 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v63)+4)))
			v102 = v99 + v100
			v103 = *(*int32)(unsafe.Add(mBase, uint32(v63)))
			v235 = v103 + v65
			v236 = v102
			v237 = v67
		case 7:
			v85 = v67
			v86 = *(*int32)(unsafe.Add(mBase, uint32(v63)))
			v88 = *(*int32)(unsafe.Add(mBase, uint32(v63)+4))
			v235 = v86 + v65
			v236 = v88 + v66
			v237 = v85
		case 8:
			v80 = v67
			v81 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v63)+8)))
			v85 = v81<<(uint(int32(8))%32) + v80
			v86 = *(*int32)(unsafe.Add(mBase, uint32(v63)))
			v88 = *(*int32)(unsafe.Add(mBase, uint32(v63)+4))
			v235 = v86 + v65
			v236 = v88 + v66
			v237 = v85
		case 9:
			v75 = v67
			v76 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v63)+9)))
			v80 = v76<<(uint(int32(16))%32) + v75
			v81 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v63)+8)))
			v85 = v81<<(uint(int32(8))%32) + v80
			v86 = *(*int32)(unsafe.Add(mBase, uint32(v63)))
			v88 = *(*int32)(unsafe.Add(mBase, uint32(v63)+4))
			v235 = v86 + v65
			v236 = v88 + v66
			v237 = v85
		case 10:
			v71 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v63)+10)))
			v75 = v71<<(uint(int32(24))%32) + v67
			v76 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v63)+9)))
			v80 = v76<<(uint(int32(16))%32) + v75
			v81 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v63)+8)))
			v85 = v81<<(uint(int32(8))%32) + v80
			v86 = *(*int32)(unsafe.Add(mBase, uint32(v63)))
			v88 = *(*int32)(unsafe.Add(mBase, uint32(v63)+4))
			v235 = v86 + v65
			v236 = v88 + v66
			v237 = v85
		default:
			v235 = v65
			v236 = v66
			v237 = v67
		}
	}
	v240 = int32(14)
	v242 = v236 ^ v237 - base.I32_rotl(v236, v240)
	v246 = v242 ^ v235 - base.I32_rotl(v242, int32(11))
	v250 = v246 ^ v236 - base.I32_rotl(v246, int32(25))
	v254 = v250 ^ v242 - base.I32_rotl(v250, int32(16))
	v258 = v254 ^ v246 - base.I32_rotl(v254, int32(4))
	v262 = v258 ^ v250 - base.I32_rotl(v258, v240)
	return v262 ^ v254 - base.I32_rotl(v262, int32(24))
}
func F_hash_bytes_extended(m *base.Module, l0 int32, l1 int32, l2 int64) int64 {
	mBase := m.M
	_ = mBase
	var v9 int32
	_ = v9
	var v13 int32
	_ = v13
	var v14 int32
	_ = v14
	var v18 int32
	_ = v18
	var v20 int32
	_ = v20
	var v24 int32
	_ = v24
	var v28 int32
	_ = v28
	var v29 int32
	_ = v29
	var v30 int32
	_ = v30
	var v31 int32
	_ = v31
	var v35 int32
	_ = v35
	var v39 int32
	_ = v39
	var v44 int32
	_ = v44
	var v46 int32
	_ = v46
	var v48 int32
	_ = v48
	var v50 int32
	_ = v50
	var v55 int32
	_ = v55
	var v56 int32
	_ = v56
	var v58 int32
	_ = v58
	var v59 int32
	_ = v59
	var v60 int32
	_ = v60
	var v62 int32
	_ = v62
	var v63 int32
	_ = v63
	var v64 int32
	_ = v64
	var v66 int32
	_ = v66
	var v67 int32
	_ = v67
	var v69 int32
	_ = v69
	var v71 int32
	_ = v71
	var v75 int32
	_ = v75
	var v76 int32
	_ = v76
	var v77 int32
	_ = v77
	var v78 int32
	_ = v78
	var v82 int32
	_ = v82
	var v86 int32
	_ = v86
	var v90 int32
	_ = v90
	var v91 int32
	_ = v91
	var v92 int32
	_ = v92
	var v93 int32
	_ = v93
	var v97 int32
	_ = v97
	var v98 int32
	_ = v98
	var v99 int32
	_ = v99
	var v101 int32
	_ = v101
	var v104 int32
	_ = v104
	var v105 int32
	_ = v105
	var v107 int32
	_ = v107
	var v108 int32
	_ = v108
	var v109 int32
	_ = v109
	var v115 int32
	_ = v115
	var v116 int32
	_ = v116
	var v118 int32
	_ = v118
	var v119 int32
	_ = v119
	var v120 int32
	_ = v120
	var v122 int32
	_ = v122
	var v123 int32
	_ = v123
	var v124 int32
	_ = v124
	var v126 int32
	_ = v126
	var v127 int32
	_ = v127
	var v129 int32
	_ = v129
	var v131 int32
	_ = v131
	var v135 int32
	_ = v135
	var v136 int32
	_ = v136
	var v137 int32
	_ = v137
	var v138 int32
	_ = v138
	var v142 int32
	_ = v142
	var v146 int32
	_ = v146
	var v150 int32
	_ = v150
	var v151 int32
	_ = v151
	var v152 int32
	_ = v152
	var v153 int32
	_ = v153
	var v157 int32
	_ = v157
	var v158 int32
	_ = v158
	var v159 int32
	_ = v159
	var v161 int32
	_ = v161
	var v164 int32
	_ = v164
	var v165 int32
	_ = v165
	var v167 int32
	_ = v167
	var v168 int32
	_ = v168
	var v169 int32
	_ = v169
	var v173 int32
	_ = v173
	var v177 int32
	_ = v177
	var v178 int32
	_ = v178
	var v182 int32
	_ = v182
	var v183 int32
	_ = v183
	var v187 int32
	_ = v187
	var v188 int32
	_ = v188
	var v190 int32
	_ = v190
	var v192 int32
	_ = v192
	var v196 int32
	_ = v196
	var v197 int32
	_ = v197
	var v201 int32
	_ = v201
	var v202 int32
	_ = v202
	var v204 int32
	_ = v204
	var v205 int32
	_ = v205
	var v207 int32
	_ = v207
	var v211 int32
	_ = v211
	var v212 int32
	_ = v212
	var v216 int32
	_ = v216
	var v217 int32
	_ = v217
	var v219 int32
	_ = v219
	var v223 int32
	_ = v223
	var v224 int32
	_ = v224
	var v228 int32
	_ = v228
	var v229 int32
	_ = v229
	var v233 int32
	_ = v233
	var v234 int32
	_ = v234
	var v238 int32
	_ = v238
	var v239 int32
	_ = v239
	var v240 int32
	_ = v240
	var v244 int32
	_ = v244
	var v245 int32
	_ = v245
	var v246 int32
	_ = v246
	var v250 int32
	_ = v250
	var v251 int32
	_ = v251
	var v252 int32
	_ = v252
	var v254 int32
	_ = v254
	var v255 int32
	_ = v255
	var v256 int32
	_ = v256
	var v260 int32
	_ = v260
	var v261 int32
	_ = v261
	var v262 int32
	_ = v262
	var v263 int32
	_ = v263
	var v267 int32
	_ = v267
	var v268 int32
	_ = v268
	var v269 int32
	_ = v269
	var v270 int32
	_ = v270
	var v274 int32
	_ = v274
	var v275 int32
	_ = v275
	var v276 int32
	_ = v276
	var v277 int32
	_ = v277
	var v282 int32
	_ = v282
	var v283 int32
	_ = v283
	var v284 int32
	_ = v284
	var v287 int32
	_ = v287
	var v289 int32
	_ = v289
	var v293 int32
	_ = v293
	var v297 int32
	_ = v297
	var v301 int32
	_ = v301
	var v305 int32
	_ = v305
	var v309 int32
	_ = v309
	v9 = l1 - int32(1636608432)
	if l2 == int64(0) {
		v46 = v9
		v48 = v9
		v50 = v9
	} else {
		v13 = v9 + base.I32_wrap_i64(l2)
		v14 = v13 + v9
		v18 = int32(4)
		v20 = base.I32_wrap_i64(int64(base.Ui64(l2)>>(uint(int64(32))%64))) ^ base.I32_rotl(v9, v18)
		v24 = v13 - v20 ^ base.I32_rotl(v20, int32(6))
		v28 = v14 - v24 ^ base.I32_rotl(v24, int32(8))
		v29 = v14 + v20
		v30 = v24 + v29
		v31 = v28 + v30
		v35 = v29 - v28 ^ base.I32_rotl(v28, int32(16))
		v39 = v30 - v35 ^ base.I32_rotl(v35, int32(19))
		v44 = v31 + v35
		v46 = v44
		v48 = v31 - v39 ^ base.I32_rotl(v39, v18)
		v50 = v39 + v44
	}
	if l0&int32(3) != 0 {
		if base.Ui32(int32(11)) < base.Ui32(l1) {
			v55 = l0
			v56 = l1
			v58 = v46
			v59 = v50
			v60 = v48
			for {
				v62 = *(*int32)(unsafe.Add(mBase, uint32(v55)+4))
				v63 = v62 + v59
				v64 = *(*int32)(unsafe.Add(mBase, uint32(v55)))
				v66 = *(*int32)(unsafe.Add(mBase, uint32(v55)+8))
				v67 = v66 + v60
				v69 = int32(4)
				v71 = v64 + v58 - v67 ^ base.I32_rotl(v67, v69)
				v75 = v63 - v71 ^ base.I32_rotl(v71, int32(6))
				v76 = v67 + v63
				v77 = v71 + v76
				v78 = v75 + v77
				v82 = v76 - v75 ^ base.I32_rotl(v75, int32(8))
				v86 = v77 - v82 ^ base.I32_rotl(v82, int32(16))
				v90 = v78 - v86 ^ base.I32_rotl(v86, int32(19))
				v91 = v82 + v78
				v92 = v86 + v91
				v93 = v90 + v92
				v97 = v91 - v90 ^ base.I32_rotl(v90, v69)
				v98 = int32(12)
				v99 = v55 + v98
				v101 = v56 - v98
				if base.Ui32(int32(11)) < base.Ui32(v101) {
					v55 = v99
					v56 = v101
					v58 = v92
					v59 = v93
					v60 = v97
					continue
				} else {
					break
				}
				break
			}
			v104 = v99
			v105 = v101
			v107 = v92
			v108 = v93
			v109 = v97
		} else {
			v104 = l0
			v105 = l1
			v107 = v46
			v108 = v50
			v109 = v48
		}
		switch v105 - int32(1) {
		case 0:
			v274 = v107
			v275 = v108
			v276 = v109
			v277 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v104))))
			v282 = v274 + v277
			v283 = v275
			v284 = v276
		case 1:
			v267 = v107
			v268 = v108
			v269 = v109
			v270 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v104)+1)))
			v274 = v270<<(uint(int32(8))%32) + v267
			v275 = v268
			v276 = v269
			v277 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v104))))
			v282 = v274 + v277
			v283 = v275
			v284 = v276
		case 2:
			v260 = v107
			v261 = v108
			v262 = v109
			v263 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v104)+2)))
			v267 = v263<<(uint(int32(16))%32) + v260
			v268 = v261
			v269 = v262
			v270 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v104)+1)))
			v274 = v270<<(uint(int32(8))%32) + v267
			v275 = v268
			v276 = v269
			v277 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v104))))
			v282 = v274 + v277
			v283 = v275
			v284 = v276
		case 3:
			v254 = v108
			v255 = v109
			v256 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v104)+3)))
			v260 = v256<<(uint(int32(24))%32) + v107
			v261 = v254
			v262 = v255
			v263 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v104)+2)))
			v267 = v263<<(uint(int32(16))%32) + v260
			v268 = v261
			v269 = v262
			v270 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v104)+1)))
			v274 = v270<<(uint(int32(8))%32) + v267
			v275 = v268
			v276 = v269
			v277 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v104))))
			v282 = v274 + v277
			v283 = v275
			v284 = v276
		case 4:
			v250 = v108
			v251 = v109
			v252 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v104)+4)))
			v254 = v250 + v252
			v255 = v251
			v256 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v104)+3)))
			v260 = v256<<(uint(int32(24))%32) + v107
			v261 = v254
			v262 = v255
			v263 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v104)+2)))
			v267 = v263<<(uint(int32(16))%32) + v260
			v268 = v261
			v269 = v262
			v270 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v104)+1)))
			v274 = v270<<(uint(int32(8))%32) + v267
			v275 = v268
			v276 = v269
			v277 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v104))))
			v282 = v274 + v277
			v283 = v275
			v284 = v276
		case 5:
			v244 = v108
			v245 = v109
			v246 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v104)+5)))
			v250 = v246<<(uint(int32(8))%32) + v244
			v251 = v245
			v252 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v104)+4)))
			v254 = v250 + v252
			v255 = v251
			v256 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v104)+3)))
			v260 = v256<<(uint(int32(24))%32) + v107
			v261 = v254
			v262 = v255
			v263 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v104)+2)))
			v267 = v263<<(uint(int32(16))%32) + v260
			v268 = v261
			v269 = v262
			v270 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v104)+1)))
			v274 = v270<<(uint(int32(8))%32) + v267
			v275 = v268
			v276 = v269
			v277 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v104))))
			v282 = v274 + v277
			v283 = v275
			v284 = v276
		case 6:
			v238 = v108
			v239 = v109
			v240 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v104)+6)))
			v244 = v240<<(uint(int32(16))%32) + v238
			v245 = v239
			v246 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v104)+5)))
			v250 = v246<<(uint(int32(8))%32) + v244
			v251 = v245
			v252 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v104)+4)))
			v254 = v250 + v252
			v255 = v251
			v256 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v104)+3)))
			v260 = v256<<(uint(int32(24))%32) + v107
			v261 = v254
			v262 = v255
			v263 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v104)+2)))
			v267 = v263<<(uint(int32(16))%32) + v260
			v268 = v261
			v269 = v262
			v270 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v104)+1)))
			v274 = v270<<(uint(int32(8))%32) + v267
			v275 = v268
			v276 = v269
			v277 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v104))))
			v282 = v274 + v277
			v283 = v275
			v284 = v276
		case 7:
			v233 = v109
			v234 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v104)+7)))
			v238 = v234<<(uint(int32(24))%32) + v108
			v239 = v233
			v240 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v104)+6)))
			v244 = v240<<(uint(int32(16))%32) + v238
			v245 = v239
			v246 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v104)+5)))
			v250 = v246<<(uint(int32(8))%32) + v244
			v251 = v245
			v252 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v104)+4)))
			v254 = v250 + v252
			v255 = v251
			v256 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v104)+3)))
			v260 = v256<<(uint(int32(24))%32) + v107
			v261 = v254
			v262 = v255
			v263 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v104)+2)))
			v267 = v263<<(uint(int32(16))%32) + v260
			v268 = v261
			v269 = v262
			v270 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v104)+1)))
			v274 = v270<<(uint(int32(8))%32) + v267
			v275 = v268
			v276 = v269
			v277 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v104))))
			v282 = v274 + v277
			v283 = v275
			v284 = v276
		case 8:
			v228 = v109
			v229 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v104)+8)))
			v233 = v229<<(uint(int32(8))%32) + v228
			v234 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v104)+7)))
			v238 = v234<<(uint(int32(24))%32) + v108
			v239 = v233
			v240 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v104)+6)))
			v244 = v240<<(uint(int32(16))%32) + v238
			v245 = v239
			v246 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v104)+5)))
			v250 = v246<<(uint(int32(8))%32) + v244
			v251 = v245
			v252 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v104)+4)))
			v254 = v250 + v252
			v255 = v251
			v256 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v104)+3)))
			v260 = v256<<(uint(int32(24))%32) + v107
			v261 = v254
			v262 = v255
			v263 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v104)+2)))
			v267 = v263<<(uint(int32(16))%32) + v260
			v268 = v261
			v269 = v262
			v270 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v104)+1)))
			v274 = v270<<(uint(int32(8))%32) + v267
			v275 = v268
			v276 = v269
			v277 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v104))))
			v282 = v274 + v277
			v283 = v275
			v284 = v276
		case 9:
			v223 = v109
			v224 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v104)+9)))
			v228 = v224<<(uint(int32(16))%32) + v223
			v229 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v104)+8)))
			v233 = v229<<(uint(int32(8))%32) + v228
			v234 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v104)+7)))
			v238 = v234<<(uint(int32(24))%32) + v108
			v239 = v233
			v240 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v104)+6)))
			v244 = v240<<(uint(int32(16))%32) + v238
			v245 = v239
			v246 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v104)+5)))
			v250 = v246<<(uint(int32(8))%32) + v244
			v251 = v245
			v252 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v104)+4)))
			v254 = v250 + v252
			v255 = v251
			v256 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v104)+3)))
			v260 = v256<<(uint(int32(24))%32) + v107
			v261 = v254
			v262 = v255
			v263 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v104)+2)))
			v267 = v263<<(uint(int32(16))%32) + v260
			v268 = v261
			v269 = v262
			v270 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v104)+1)))
			v274 = v270<<(uint(int32(8))%32) + v267
			v275 = v268
			v276 = v269
			v277 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v104))))
			v282 = v274 + v277
			v283 = v275
			v284 = v276
		case 10:
			v219 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v104)+10)))
			v223 = v219<<(uint(int32(24))%32) + v109
			v224 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v104)+9)))
			v228 = v224<<(uint(int32(16))%32) + v223
			v229 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v104)+8)))
			v233 = v229<<(uint(int32(8))%32) + v228
			v234 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v104)+7)))
			v238 = v234<<(uint(int32(24))%32) + v108
			v239 = v233
			v240 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v104)+6)))
			v244 = v240<<(uint(int32(16))%32) + v238
			v245 = v239
			v246 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v104)+5)))
			v250 = v246<<(uint(int32(8))%32) + v244
			v251 = v245
			v252 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v104)+4)))
			v254 = v250 + v252
			v255 = v251
			v256 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v104)+3)))
			v260 = v256<<(uint(int32(24))%32) + v107
			v261 = v254
			v262 = v255
			v263 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v104)+2)))
			v267 = v263<<(uint(int32(16))%32) + v260
			v268 = v261
			v269 = v262
			v270 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v104)+1)))
			v274 = v270<<(uint(int32(8))%32) + v267
			v275 = v268
			v276 = v269
			v277 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v104))))
			v282 = v274 + v277
			v283 = v275
			v284 = v276
		default:
			v282 = v107
			v283 = v108
			v284 = v109
		}
	} else {
		if base.Ui32(int32(12)) <= base.Ui32(l1) {
			v115 = l0
			v116 = l1
			v118 = v46
			v119 = v50
			v120 = v48
			for {
				v122 = *(*int32)(unsafe.Add(mBase, uint32(v115)+4))
				v123 = v122 + v119
				v124 = *(*int32)(unsafe.Add(mBase, uint32(v115)))
				v126 = *(*int32)(unsafe.Add(mBase, uint32(v115)+8))
				v127 = v126 + v120
				v129 = int32(4)
				v131 = v124 + v118 - v127 ^ base.I32_rotl(v127, v129)
				v135 = v123 - v131 ^ base.I32_rotl(v131, int32(6))
				v136 = v127 + v123
				v137 = v131 + v136
				v138 = v135 + v137
				v142 = v136 - v135 ^ base.I32_rotl(v135, int32(8))
				v146 = v137 - v142 ^ base.I32_rotl(v142, int32(16))
				v150 = v138 - v146 ^ base.I32_rotl(v146, int32(19))
				v151 = v142 + v138
				v152 = v146 + v151
				v153 = v150 + v152
				v157 = v151 - v150 ^ base.I32_rotl(v150, v129)
				v158 = int32(12)
				v159 = v115 + v158
				v161 = v116 - v158
				if base.Ui32(int32(11)) < base.Ui32(v161) {
					v115 = v159
					v116 = v161
					v118 = v152
					v119 = v153
					v120 = v157
					continue
				} else {
					break
				}
				break
			}
			v164 = v159
			v165 = v161
			v167 = v152
			v168 = v153
			v169 = v157
		} else {
			v164 = l0
			v165 = l1
			v167 = v46
			v168 = v50
			v169 = v48
		}
		switch v165 - int32(1) {
		case 0:
			v216 = v167
			v217 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v164))))
			v282 = v216 + v217
			v283 = v168
			v284 = v169
		case 1:
			v211 = v167
			v212 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v164)+1)))
			v216 = v212<<(uint(int32(8))%32) + v211
			v217 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v164))))
			v282 = v216 + v217
			v283 = v168
			v284 = v169
		case 2:
			v207 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v164)+2)))
			v211 = v207<<(uint(int32(16))%32) + v167
			v212 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v164)+1)))
			v216 = v212<<(uint(int32(8))%32) + v211
			v217 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v164))))
			v282 = v216 + v217
			v283 = v168
			v284 = v169
		case 3:
			v204 = v168
			v205 = *(*int32)(unsafe.Add(mBase, uint32(v164)))
			v282 = v205 + v167
			v283 = v204
			v284 = v169
		case 4:
			v201 = v168
			v202 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v164)+4)))
			v204 = v201 + v202
			v205 = *(*int32)(unsafe.Add(mBase, uint32(v164)))
			v282 = v205 + v167
			v283 = v204
			v284 = v169
		case 5:
			v196 = v168
			v197 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v164)+5)))
			v201 = v197<<(uint(int32(8))%32) + v196
			v202 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v164)+4)))
			v204 = v201 + v202
			v205 = *(*int32)(unsafe.Add(mBase, uint32(v164)))
			v282 = v205 + v167
			v283 = v204
			v284 = v169
		case 6:
			v192 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v164)+6)))
			v196 = v192<<(uint(int32(16))%32) + v168
			v197 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v164)+5)))
			v201 = v197<<(uint(int32(8))%32) + v196
			v202 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v164)+4)))
			v204 = v201 + v202
			v205 = *(*int32)(unsafe.Add(mBase, uint32(v164)))
			v282 = v205 + v167
			v283 = v204
			v284 = v169
		case 7:
			v187 = v169
			v188 = *(*int32)(unsafe.Add(mBase, uint32(v164)))
			v190 = *(*int32)(unsafe.Add(mBase, uint32(v164)+4))
			v282 = v188 + v167
			v283 = v190 + v168
			v284 = v187
		case 8:
			v182 = v169
			v183 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v164)+8)))
			v187 = v183<<(uint(int32(8))%32) + v182
			v188 = *(*int32)(unsafe.Add(mBase, uint32(v164)))
			v190 = *(*int32)(unsafe.Add(mBase, uint32(v164)+4))
			v282 = v188 + v167
			v283 = v190 + v168
			v284 = v187
		case 9:
			v177 = v169
			v178 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v164)+9)))
			v182 = v178<<(uint(int32(16))%32) + v177
			v183 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v164)+8)))
			v187 = v183<<(uint(int32(8))%32) + v182
			v188 = *(*int32)(unsafe.Add(mBase, uint32(v164)))
			v190 = *(*int32)(unsafe.Add(mBase, uint32(v164)+4))
			v282 = v188 + v167
			v283 = v190 + v168
			v284 = v187
		case 10:
			v173 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v164)+10)))
			v177 = v173<<(uint(int32(24))%32) + v169
			v178 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v164)+9)))
			v182 = v178<<(uint(int32(16))%32) + v177
			v183 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v164)+8)))
			v187 = v183<<(uint(int32(8))%32) + v182
			v188 = *(*int32)(unsafe.Add(mBase, uint32(v164)))
			v190 = *(*int32)(unsafe.Add(mBase, uint32(v164)+4))
			v282 = v188 + v167
			v283 = v190 + v168
			v284 = v187
		default:
			v282 = v167
			v283 = v168
			v284 = v169
		}
	}
	v287 = int32(14)
	v289 = v283 ^ v284 - base.I32_rotl(v283, v287)
	v293 = v289 ^ v282 - base.I32_rotl(v289, int32(11))
	v297 = v293 ^ v283 - base.I32_rotl(v293, int32(25))
	v301 = v297 ^ v289 - base.I32_rotl(v297, int32(16))
	v305 = v301 ^ v293 - base.I32_rotl(v301, int32(4))
	v309 = v305 ^ v297 - base.I32_rotl(v305, v287)
	return base.I64_extend_i32_u(v309)<<(uint(int64(32))%64) | base.I64_extend_i32_u(v309^v301-base.I32_rotl(v309, int32(24)))
}
func F_hash_desc(m *base.Module, l0 int32, l1 int32) {
	mBase := m.M
	_ = mBase
	var v9 int32
	_ = v9
	var v11 int32
	_ = v11
	var v13 int32
	_ = v13
	var v14 int32
	_ = v14
	var v15 int32
	_ = v15
	var v18 float64
	_ = v18
	var v19 int32
	_ = v19
	var v20 int32
	_ = v20
	var v26 int32
	_ = v26
	var v27 int32
	_ = v27
	var v33 int32
	_ = v33
	var v34 int32
	_ = v34
	var v40 int32
	_ = v40
	var v41 int32
	_ = v41
	var v42 int32
	_ = v42
	var v46 int32
	_ = v46
	var v52 int32
	_ = v52
	var v53 int32
	_ = v53
	var v54 int32
	_ = v54
	var v55 int32
	_ = v55
	var v58 int32
	_ = v58
	var v61 int32
	_ = v61
	var v70 int32
	_ = v70
	var v76 int32
	_ = v76
	var v77 int32
	_ = v77
	var v78 int32
	_ = v78
	var v85 int32
	_ = v85
	var v86 int32
	_ = v86
	var v87 int32
	_ = v87
	var v91 int32
	_ = v91
	var v97 int32
	_ = v97
	var v98 int32
	_ = v98
	var v99 int64
	_ = v99
	var v100 int32
	_ = v100
	var v103 int32
	_ = v103
	var v104 int32
	_ = v104
	var v110 int32
	_ = v110
	var v116 int32
	_ = v116
	var v117 int32
	_ = v117
	var v120 int32
	_ = v120
	var v121 int32
	_ = v121
	var v125 int32
	_ = v125
	var v131 int32
	_ = v131
	var v132 float64
	_ = v132
	var v138 int32
	_ = v138
	var v139 int32
	_ = v139
	var v140 int32
	_ = v140
	var v141 int32
	_ = v141
	var v146 int32
	_ = v146
	var v152 int32
	_ = v152
	v9 = m.G0
	v11 = v9 - int32(208)
	m.G0 = v11
	v13 = *(*int32)(unsafe.Add(mBase, uint32(l1)+96))
	v14 = *(*int32)(unsafe.Add(mBase, uint32(v13)+64))
	v15 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v13)+48)))
	switch int32(base.Ui32(v15) >> (uint(int32(4)) % 32)) {
	case 0:
		v18 = *(*float64)(unsafe.Add(mBase, uint32(v14)))
		v19 = *(*int32)(unsafe.Add(mBase, uint32(v14)+8))
		v20 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v14)+12)))
		*(*int32)(unsafe.Add(mBase, uint32(v11)+12)) = v20
		*(*int32)(unsafe.Add(mBase, uint32(v11)+8)) = v19
		*(*float64)(unsafe.Add(mBase, uint32(v11))) = v18
		F_appendStringInfo(m, l0, int32(_a_F_hash_desc_0), v11)
		mBase = m.M
		v26 = m.ExcPending
		if v26 != 0 {
			return
		} else {
			m.G0 = v11 + int32(208)
			return
		}
	case 1:
		v27 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v14))))
		*(*int32)(unsafe.Add(mBase, uint32(v11)+16)) = v27
		F_appendStringInfo(m, l0, int32(_a_F_hash_desc_1), v11+int32(16))
		mBase = m.M
		v33 = m.ExcPending
		if v33 != 0 {
			return
		} else {
			m.G0 = v11 + int32(208)
			return
		}
	case 2:
		v34 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v14))))
		*(*int32)(unsafe.Add(mBase, uint32(v11)+32)) = v34
		F_appendStringInfo(m, l0, int32(_a_F_hash_desc_2), v11+int32(32))
		mBase = m.M
		v40 = m.ExcPending
		if v40 != 0 {
			return
		} else {
			m.G0 = v11 + int32(208)
			return
		}
	case 3:
		v41 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v14)+2)))
		v42 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v14))))
		*(*int32)(unsafe.Add(mBase, uint32(v11)+48)) = v42
		if v41 != 0 {
			v46 = int32(84)
		} else {
			v46 = int32(70)
		}
		*(*int32)(unsafe.Add(mBase, uint32(v11)+52)) = v46
		F_appendStringInfo(m, l0, int32(_a_F_hash_desc_3), v11+int32(48))
		mBase = m.M
		v52 = m.ExcPending
		if v52 != 0 {
			return
		} else {
			m.G0 = v11 + int32(208)
			return
		}
	case 4:
		v53 = *(*int32)(unsafe.Add(mBase, uint32(v14)))
		v54 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v14)+4)))
		v55 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v14)+6)))
		v58 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v14)+8)))
		if v58&int32(2) != 0 {
			v61 = int32(84)
		} else {
			v61 = int32(70)
		}
		*(*int32)(unsafe.Add(mBase, uint32(v11)+80)) = v61
		*(*int32)(unsafe.Add(mBase, uint32(v11)+72)) = v55
		*(*int32)(unsafe.Add(mBase, uint32(v11)+68)) = v54
		*(*int32)(unsafe.Add(mBase, uint32(v11)+64)) = v53
		if v58&int32(1) != 0 {
			v70 = int32(84)
		} else {
			v70 = int32(70)
		}
		*(*int32)(unsafe.Add(mBase, uint32(v11)+76)) = v70
		F_appendStringInfo(m, l0, int32(_a_F_hash_desc_4), v11-int32(-64))
		mBase = m.M
		v76 = m.ExcPending
		if v76 != 0 {
			return
		} else {
			m.G0 = v11 + int32(208)
			return
		}
	default:
		m.G0 = v11 + int32(208)
		return
	case 6:
		v77 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v14))))
		v78 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v14)+2)))
		*(*int32)(unsafe.Add(mBase, uint32(v11)+100)) = v78
		*(*int32)(unsafe.Add(mBase, uint32(v11)+96)) = v77
		F_appendStringInfo(m, l0, int32(_a_F_hash_desc_5), v11+int32(96))
		mBase = m.M
		v85 = m.ExcPending
		if v85 != 0 {
			return
		} else {
			m.G0 = v11 + int32(208)
			return
		}
	case 7:
		v86 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v14)+2)))
		v87 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v14))))
		*(*int32)(unsafe.Add(mBase, uint32(v11)+112)) = v87
		if v86 != 0 {
			v91 = int32(84)
		} else {
			v91 = int32(70)
		}
		*(*int32)(unsafe.Add(mBase, uint32(v11)+116)) = v91
		F_appendStringInfo(m, l0, int32(_a_F_hash_desc_6), v11+int32(112))
		mBase = m.M
		v97 = m.ExcPending
		if v97 != 0 {
			return
		} else {
			m.G0 = v11 + int32(208)
			return
		}
	case 8:
		v98 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v14)+10)))
		v99 = *(*int64)(unsafe.Add(mBase, uint32(v14)))
		v100 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v14)+8)))
		v103 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v14)+11)))
		if v103 != 0 {
			v104 = int32(84)
		} else {
			v104 = int32(70)
		}
		*(*int32)(unsafe.Add(mBase, uint32(v11)+144)) = v104
		*(*int32)(unsafe.Add(mBase, uint32(v11)+136)) = v100
		*(*int64)(unsafe.Add(mBase, uint32(v11)+128)) = v99
		if v98 != 0 {
			v110 = int32(84)
		} else {
			v110 = int32(70)
		}
		*(*int32)(unsafe.Add(mBase, uint32(v11)+140)) = v110
		F_appendStringInfo(m, l0, int32(_a_F_hash_desc_7), v11+int32(128))
		mBase = m.M
		v116 = m.ExcPending
		if v116 != 0 {
			return
		} else {
			m.G0 = v11 + int32(208)
			return
		}
	case 9:
		v117 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v14))))
		v120 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v14)+1)))
		if v120 != 0 {
			v121 = int32(84)
		} else {
			v121 = int32(70)
		}
		*(*int32)(unsafe.Add(mBase, uint32(v11)+164)) = v121
		if v117 != 0 {
			v125 = int32(84)
		} else {
			v125 = int32(70)
		}
		*(*int32)(unsafe.Add(mBase, uint32(v11)+160)) = v125
		F_appendStringInfo(m, l0, int32(_a_F_hash_desc_8), v11+int32(160))
		mBase = m.M
		v131 = m.ExcPending
		if v131 != 0 {
			return
		} else {
			m.G0 = v11 + int32(208)
			return
		}
	case 11:
		v132 = *(*float64)(unsafe.Add(mBase, uint32(v14)))
		*(*float64)(unsafe.Add(mBase, uint32(v11)+176)) = v132
		F_appendStringInfo(m, l0, int32(_a_F_hash_desc_9), v11+int32(176))
		mBase = m.M
		v138 = m.ExcPending
		if v138 != 0 {
			return
		} else {
			m.G0 = v11 + int32(208)
			return
		}
	case 12:
		v139 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v14)+6)))
		v140 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v14)+4)))
		v141 = *(*int32)(unsafe.Add(mBase, uint32(v14)))
		*(*int32)(unsafe.Add(mBase, uint32(v11)+196)) = v141
		*(*int32)(unsafe.Add(mBase, uint32(v11)+192)) = v140
		if v139 != 0 {
			v146 = int32(84)
		} else {
			v146 = int32(70)
		}
		*(*int32)(unsafe.Add(mBase, uint32(v11)+200)) = v146
		F_appendStringInfo(m, l0, int32(_a_F_hash_desc_10), v11+int32(192))
		mBase = m.M
		v152 = m.ExcPending
		if v152 != 0 {
			return
		} else {
			m.G0 = v11 + int32(208)
			return
		}
	}
}
func F_hash_numeric_extended(m *base.Module, l0 int32) int64 {
	mBase := m.M
	_ = mBase
	var v12 int32
	_ = v12
	var v13 int32
	_ = v13
	var v16 int32
	_ = v16
	var v17 int64
	_ = v17
	var v18 int32
	_ = v18
	var v19 int32
	_ = v19
	var v23 int32
	_ = v23
	var v36 int32
	_ = v36
	var v38 int32
	_ = v38
	var v39 int32
	_ = v39
	var v40 int32
	_ = v40
	var v43 int32
	_ = v43
	var v45 int32
	_ = v45
	var v48 int32
	_ = v48
	var v50 int32
	_ = v50
	var v52 int32
	_ = v52
	var v55 int32
	_ = v55
	var v56 int32
	_ = v56
	var v59 int32
	_ = v59
	var v70 int32
	_ = v70
	var v73 int32
	_ = v73
	var v76 int32
	_ = v76
	var v80 int32
	_ = v80
	var v83 int32
	_ = v83
	var v90 int32
	_ = v90
	var v91 int32
	_ = v91
	var v95 int32
	_ = v95
	var v97 int32
	_ = v97
	var v99 int32
	_ = v99
	var v104 int32
	_ = v104
	var v105 int32
	_ = v105
	var v111 int32
	_ = v111
	var v117 int32
	_ = v117
	var v121 int32
	_ = v121
	var v122 int32
	_ = v122
	var v126 int32
	_ = v126
	var v128 int32
	_ = v128
	var v132 int32
	_ = v132
	var v136 int32
	_ = v136
	var v137 int32
	_ = v137
	var v138 int32
	_ = v138
	var v139 int32
	_ = v139
	var v143 int32
	_ = v143
	var v147 int32
	_ = v147
	var v152 int32
	_ = v152
	var v154 int32
	_ = v154
	var v156 int32
	_ = v156
	var v158 int32
	_ = v158
	var v163 int32
	_ = v163
	var v164 int32
	_ = v164
	var v166 int32
	_ = v166
	var v167 int32
	_ = v167
	var v168 int32
	_ = v168
	var v170 int32
	_ = v170
	var v171 int32
	_ = v171
	var v172 int32
	_ = v172
	var v174 int32
	_ = v174
	var v175 int32
	_ = v175
	var v177 int32
	_ = v177
	var v179 int32
	_ = v179
	var v183 int32
	_ = v183
	var v184 int32
	_ = v184
	var v185 int32
	_ = v185
	var v186 int32
	_ = v186
	var v190 int32
	_ = v190
	var v194 int32
	_ = v194
	var v198 int32
	_ = v198
	var v199 int32
	_ = v199
	var v200 int32
	_ = v200
	var v201 int32
	_ = v201
	var v205 int32
	_ = v205
	var v206 int32
	_ = v206
	var v207 int32
	_ = v207
	var v209 int32
	_ = v209
	var v212 int32
	_ = v212
	var v213 int32
	_ = v213
	var v215 int32
	_ = v215
	var v216 int32
	_ = v216
	var v217 int32
	_ = v217
	var v223 int32
	_ = v223
	var v224 int32
	_ = v224
	var v226 int32
	_ = v226
	var v227 int32
	_ = v227
	var v228 int32
	_ = v228
	var v230 int32
	_ = v230
	var v231 int32
	_ = v231
	var v232 int32
	_ = v232
	var v234 int32
	_ = v234
	var v235 int32
	_ = v235
	var v237 int32
	_ = v237
	var v239 int32
	_ = v239
	var v243 int32
	_ = v243
	var v244 int32
	_ = v244
	var v245 int32
	_ = v245
	var v246 int32
	_ = v246
	var v250 int32
	_ = v250
	var v254 int32
	_ = v254
	var v258 int32
	_ = v258
	var v259 int32
	_ = v259
	var v260 int32
	_ = v260
	var v261 int32
	_ = v261
	var v265 int32
	_ = v265
	var v266 int32
	_ = v266
	var v267 int32
	_ = v267
	var v269 int32
	_ = v269
	var v272 int32
	_ = v272
	var v273 int32
	_ = v273
	var v275 int32
	_ = v275
	var v276 int32
	_ = v276
	var v277 int32
	_ = v277
	var v281 int32
	_ = v281
	var v285 int32
	_ = v285
	var v286 int32
	_ = v286
	var v290 int32
	_ = v290
	var v291 int32
	_ = v291
	var v295 int32
	_ = v295
	var v296 int32
	_ = v296
	var v298 int32
	_ = v298
	var v300 int32
	_ = v300
	var v304 int32
	_ = v304
	var v305 int32
	_ = v305
	var v309 int32
	_ = v309
	var v310 int32
	_ = v310
	var v312 int32
	_ = v312
	var v313 int32
	_ = v313
	var v315 int32
	_ = v315
	var v319 int32
	_ = v319
	var v320 int32
	_ = v320
	var v324 int32
	_ = v324
	var v325 int32
	_ = v325
	var v327 int32
	_ = v327
	var v331 int32
	_ = v331
	var v332 int32
	_ = v332
	var v336 int32
	_ = v336
	var v337 int32
	_ = v337
	var v341 int32
	_ = v341
	var v342 int32
	_ = v342
	var v346 int32
	_ = v346
	var v347 int32
	_ = v347
	var v348 int32
	_ = v348
	var v352 int32
	_ = v352
	var v353 int32
	_ = v353
	var v354 int32
	_ = v354
	var v358 int32
	_ = v358
	var v359 int32
	_ = v359
	var v360 int32
	_ = v360
	var v362 int32
	_ = v362
	var v363 int32
	_ = v363
	var v364 int32
	_ = v364
	var v368 int32
	_ = v368
	var v369 int32
	_ = v369
	var v370 int32
	_ = v370
	var v371 int32
	_ = v371
	var v375 int32
	_ = v375
	var v376 int32
	_ = v376
	var v377 int32
	_ = v377
	var v378 int32
	_ = v378
	var v382 int32
	_ = v382
	var v383 int32
	_ = v383
	var v384 int32
	_ = v384
	var v385 int32
	_ = v385
	var v390 int32
	_ = v390
	var v391 int32
	_ = v391
	var v392 int32
	_ = v392
	var v395 int32
	_ = v395
	var v397 int32
	_ = v397
	var v401 int32
	_ = v401
	var v405 int32
	_ = v405
	var v409 int32
	_ = v409
	var v413 int32
	_ = v413
	var v417 int32
	_ = v417
	var v440 int64
	_ = v440
	v12 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
	v13 = F_pg_detoast_datum(m, v12)
	mBase = m.M
	v16 = m.ExcPending
	if v16 != 0 {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	return int64(0)
L2:
	;
	v17 = *(*int64)(unsafe.Add(mBase, uint32(l0)+40))
	v18 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v13)+4)))
	v19 = int32(_a_F_hash_numeric_extended_0)
	if v18&v19 != v19 {
		goto L4
	} else {
		goto L5
	}
L3:
	;
	return v17 - int64(1)
L4:
	;
	v23 = base.I32_extend16_s(v18)
	if v23 < int32(0) {
		goto L8
	} else {
		goto L9
	}
L5:
	;
	v440 = v17
	goto L6
L6:
	;
	return v440
L7:
	;
	v40 = *(*int32)(unsafe.Add(mBase, uint32(v13)))
	v43 = v39 + int32(base.Ui32(v40)>>(uint(int32(2))%32))
	v45 = int32(base.Ui32(v43) >> (uint(int32(1)) % 32))
	if v45 == int32(0) {
		goto L3
	} else {
		goto L11
	}
L8:
	;
	v38 = v18<<(uint(int32(25))%32)>>(uint(int32(31))%32)&int32(-64) | v18&int32(63)
	v39 = int32(-6)
	goto L7
L9:
	;
	goto L10
L10:
	;
	v36 = int32(*(*int16)(unsafe.Add(mBase, uint32(v13)+6)))
	v38 = v36
	v39 = int32(-8)
	goto L7
L11:
	;
	v48 = int32(0)
	v50 = v13 + int32(6)
	v52 = v13 + int32(8)
	if v23 < v48 {
		goto L12
	} else {
		goto L13
	}
L12:
	;
	v55 = v50
	goto L14
L13:
	;
	v55 = v52
	goto L14
L14:
	;
	v56 = v48
	v59 = v38
	goto L15
L15:
	;
	v70 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v55+v56<<(uint(int32(1))%32)))))
	if v70 == int32(0) {
		goto L17
	} else {
		goto L18
	}
L16:
	;
	if v56 == v45 {
		goto L3
	} else {
		goto L21
	}
L17:
	;
	v73 = int32(1)
	v76 = v56 + v73
	if v76 != v45 {
		v56 = v76
		v59 = v59 - v73
		goto L15
	} else {
		goto L20
	}
L18:
	;
	goto L19
L19:
	;
	goto L16
L20:
	;
	goto L3
L21:
	;
	v80 = v45
	v83 = int32(0)
	goto L23
L22:
	;
	if int32(0) <= v23 {
		goto L27
	} else {
		goto L28
	}
L23:
	;
	v90 = int32(1)
	v91 = v80 - v90
	v95 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v55+v91<<(uint(v90)%32)))))
	if v95 != 0 {
		v99 = v83
		goto L22
	} else {
		goto L25
	}
L24:
	;
	v99 = v45
	goto L22
L25:
	;
	v97 = v83 + int32(1)
	if v97 != v45 {
		v80 = v91
		v83 = v97
		goto L23
	} else {
		goto L26
	}
L26:
	;
	goto L24
L27:
	;
	v104 = v52
	goto L29
L28:
	;
	v104 = v50
	goto L29
L29:
	;
	v105 = v56<<(uint(int32(1))%32) + v104
	v111 = (v43 - (v56+v99)<<(uint(int32(1))%32)) & int32(-2)
	v117 = v111 - int32(1636608432)
	if v17 == int64(0) {
		goto L32
	} else {
		goto L33
	}
L30:
	;
	v440 = base.I64_extend_i32_u(v417)<<(uint(int64(32))%64) | base.I64_extend_i32_u(v417^v409-base.I32_rotl(v417, int32(24))) ^ base.I64_extend_i32_s(v59)
	goto L6
L31:
	;
	if v105&int32(3) != 0 {
		goto L47
	} else {
		goto L48
	}
L32:
	;
	v154 = v117
	v156 = v117
	v158 = v117
	goto L31
L33:
	;
	goto L34
L34:
	;
	v121 = v117 + base.I32_wrap_i64(v17)
	v122 = v121 + v117
	v126 = int32(4)
	v128 = base.I32_wrap_i64(int64(base.Ui64(v17)>>(uint(int64(32))%64))) ^ base.I32_rotl(v117, v126)
	v132 = v121 - v128 ^ base.I32_rotl(v128, int32(6))
	v136 = v122 - v132 ^ base.I32_rotl(v132, int32(8))
	v137 = v122 + v128
	v138 = v132 + v137
	v139 = v136 + v138
	v143 = v137 - v136 ^ base.I32_rotl(v136, int32(16))
	v147 = v138 - v143 ^ base.I32_rotl(v143, int32(19))
	v152 = v139 + v143
	v154 = v152
	v156 = v139 - v147 ^ base.I32_rotl(v147, v126)
	v158 = v147 + v152
	goto L31
L35:
	;
	v395 = int32(14)
	v397 = v391 ^ v392 - base.I32_rotl(v391, v395)
	v401 = v397 ^ v390 - base.I32_rotl(v397, int32(11))
	v405 = v401 ^ v391 - base.I32_rotl(v401, int32(25))
	v409 = v405 ^ v397 - base.I32_rotl(v405, int32(16))
	v413 = v409 ^ v401 - base.I32_rotl(v409, int32(4))
	v417 = v413 ^ v405 - base.I32_rotl(v413, v395)
	goto L30
L36:
	;
	v385 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v212))))
	v390 = v382 + v385
	v391 = v383
	v392 = v384
	goto L35
L37:
	;
	v378 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v212)+1)))
	v382 = v378<<(uint(int32(8))%32) + v375
	v383 = v376
	v384 = v377
	goto L36
L38:
	;
	v371 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v212)+2)))
	v375 = v371<<(uint(int32(16))%32) + v368
	v376 = v369
	v377 = v370
	goto L37
L39:
	;
	v364 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v212)+3)))
	v368 = v364<<(uint(int32(24))%32) + v215
	v369 = v362
	v370 = v363
	goto L38
L40:
	;
	v360 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v212)+4)))
	v362 = v358 + v360
	v363 = v359
	goto L39
L41:
	;
	v354 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v212)+5)))
	v358 = v354<<(uint(int32(8))%32) + v352
	v359 = v353
	goto L40
L42:
	;
	v348 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v212)+6)))
	v352 = v348<<(uint(int32(16))%32) + v346
	v353 = v347
	goto L41
L43:
	;
	v342 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v212)+7)))
	v346 = v342<<(uint(int32(24))%32) + v216
	v347 = v341
	goto L42
L44:
	;
	v337 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v212)+8)))
	v341 = v337<<(uint(int32(8))%32) + v336
	goto L43
L45:
	;
	v332 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v212)+9)))
	v336 = v332<<(uint(int32(16))%32) + v331
	goto L44
L46:
	;
	v327 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v212)+10)))
	v331 = v327<<(uint(int32(24))%32) + v217
	goto L45
L47:
	;
	if base.Ui32(int32(11)) < base.Ui32(v111) {
		goto L50
	} else {
		goto L51
	}
L48:
	;
	goto L49
L49:
	;
	if base.Ui32(int32(12)) <= base.Ui32(v111) {
		goto L56
	} else {
		goto L57
	}
L50:
	;
	v163 = v105
	v164 = v111
	v166 = v154
	v167 = v158
	v168 = v156
	goto L53
L51:
	;
	v212 = v105
	v213 = v111
	v215 = v154
	v216 = v158
	v217 = v156
	goto L52
L52:
	;
	switch v213 - int32(1) {
	case 0:
		v382 = v215
		v383 = v216
		v384 = v217
		goto L36
	case 1:
		v375 = v215
		v376 = v216
		v377 = v217
		goto L37
	case 2:
		v368 = v215
		v369 = v216
		v370 = v217
		goto L38
	case 3:
		v362 = v216
		v363 = v217
		goto L39
	case 4:
		v358 = v216
		v359 = v217
		goto L40
	case 5:
		v352 = v216
		v353 = v217
		goto L41
	case 6:
		v346 = v216
		v347 = v217
		goto L42
	case 7:
		v341 = v217
		goto L43
	case 8:
		v336 = v217
		goto L44
	case 9:
		v331 = v217
		goto L45
	case 10:
		goto L46
	default:
		v390 = v215
		v391 = v216
		v392 = v217
		goto L35
	}
L53:
	;
	v170 = *(*int32)(unsafe.Add(mBase, uint32(v163)+4))
	v171 = v170 + v167
	v172 = *(*int32)(unsafe.Add(mBase, uint32(v163)))
	v174 = *(*int32)(unsafe.Add(mBase, uint32(v163)+8))
	v175 = v174 + v168
	v177 = int32(4)
	v179 = v172 + v166 - v175 ^ base.I32_rotl(v175, v177)
	v183 = v171 - v179 ^ base.I32_rotl(v179, int32(6))
	v184 = v175 + v171
	v185 = v179 + v184
	v186 = v183 + v185
	v190 = v184 - v183 ^ base.I32_rotl(v183, int32(8))
	v194 = v185 - v190 ^ base.I32_rotl(v190, int32(16))
	v198 = v186 - v194 ^ base.I32_rotl(v194, int32(19))
	v199 = v190 + v186
	v200 = v194 + v199
	v201 = v198 + v200
	v205 = v199 - v198 ^ base.I32_rotl(v198, v177)
	v206 = int32(12)
	v207 = v163 + v206
	v209 = v164 - v206
	if base.Ui32(int32(11)) < base.Ui32(v209) {
		v163 = v207
		v164 = v209
		v166 = v200
		v167 = v201
		v168 = v205
		goto L53
	} else {
		goto L55
	}
L54:
	;
	v212 = v207
	v213 = v209
	v215 = v200
	v216 = v201
	v217 = v205
	goto L52
L55:
	;
	goto L54
L56:
	;
	v223 = v105
	v224 = v111
	v226 = v154
	v227 = v158
	v228 = v156
	goto L59
L57:
	;
	v272 = v105
	v273 = v111
	v275 = v154
	v276 = v158
	v277 = v156
	goto L58
L58:
	;
	switch v273 - int32(1) {
	case 0:
		v324 = v275
		goto L62
	case 1:
		v319 = v275
		goto L63
	case 2:
		goto L64
	case 3:
		v312 = v276
		goto L65
	case 4:
		v309 = v276
		goto L66
	case 5:
		v304 = v276
		goto L67
	case 6:
		goto L68
	case 7:
		v295 = v277
		goto L69
	case 8:
		v290 = v277
		goto L70
	case 9:
		v285 = v277
		goto L71
	case 10:
		goto L72
	default:
		v390 = v275
		v391 = v276
		v392 = v277
		goto L35
	}
L59:
	;
	v230 = *(*int32)(unsafe.Add(mBase, uint32(v223)+4))
	v231 = v230 + v227
	v232 = *(*int32)(unsafe.Add(mBase, uint32(v223)))
	v234 = *(*int32)(unsafe.Add(mBase, uint32(v223)+8))
	v235 = v234 + v228
	v237 = int32(4)
	v239 = v232 + v226 - v235 ^ base.I32_rotl(v235, v237)
	v243 = v231 - v239 ^ base.I32_rotl(v239, int32(6))
	v244 = v235 + v231
	v245 = v239 + v244
	v246 = v243 + v245
	v250 = v244 - v243 ^ base.I32_rotl(v243, int32(8))
	v254 = v245 - v250 ^ base.I32_rotl(v250, int32(16))
	v258 = v246 - v254 ^ base.I32_rotl(v254, int32(19))
	v259 = v250 + v246
	v260 = v254 + v259
	v261 = v258 + v260
	v265 = v259 - v258 ^ base.I32_rotl(v258, v237)
	v266 = int32(12)
	v267 = v223 + v266
	v269 = v224 - v266
	if base.Ui32(int32(11)) < base.Ui32(v269) {
		v223 = v267
		v224 = v269
		v226 = v260
		v227 = v261
		v228 = v265
		goto L59
	} else {
		goto L61
	}
L60:
	;
	v272 = v267
	v273 = v269
	v275 = v260
	v276 = v261
	v277 = v265
	goto L58
L61:
	;
	goto L60
L62:
	;
	v325 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v272))))
	v390 = v324 + v325
	v391 = v276
	v392 = v277
	goto L35
L63:
	;
	v320 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v272)+1)))
	v324 = v320<<(uint(int32(8))%32) + v319
	goto L62
L64:
	;
	v315 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v272)+2)))
	v319 = v315<<(uint(int32(16))%32) + v275
	goto L63
L65:
	;
	v313 = *(*int32)(unsafe.Add(mBase, uint32(v272)))
	v390 = v313 + v275
	v391 = v312
	v392 = v277
	goto L35
L66:
	;
	v310 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v272)+4)))
	v312 = v309 + v310
	goto L65
L67:
	;
	v305 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v272)+5)))
	v309 = v305<<(uint(int32(8))%32) + v304
	goto L66
L68:
	;
	v300 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v272)+6)))
	v304 = v300<<(uint(int32(16))%32) + v276
	goto L67
L69:
	;
	v296 = *(*int32)(unsafe.Add(mBase, uint32(v272)))
	v298 = *(*int32)(unsafe.Add(mBase, uint32(v272)+4))
	v390 = v296 + v275
	v391 = v298 + v276
	v392 = v295
	goto L35
L70:
	;
	v291 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v272)+8)))
	v295 = v291<<(uint(int32(8))%32) + v290
	goto L69
L71:
	;
	v286 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v272)+9)))
	v290 = v286<<(uint(int32(16))%32) + v285
	goto L70
L72:
	;
	v281 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v272)+10)))
	v285 = v281<<(uint(int32(24))%32) + v277
	goto L71
}
func F_hash_object_field_start(m *base.Module, l0 int32, l1 int32, l2 int32) int32 {
	mBase := m.M
	_ = mBase
	var v6 int32
	_ = v6
	var v7 int32
	_ = v7
	var v9 int32
	_ = v9
	var v14 int32
	_ = v14
	var v15 int32
	_ = v15
	v6 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v7 = *(*int32)(unsafe.Add(mBase, uint32(v6)+32))
	if v7 <= int32(1) {
		v9 = *(*int32)(unsafe.Add(mBase, uint32(v6)+28))
		*(*int32)(unsafe.Add(mBase, uint32(l0)+20)) = v9
		switch v9 - int32(3) {
		case 0, 2:
			v14 = *(*int32)(unsafe.Add(mBase, uint32(v6)+12))
			v15 = v14
		default:
			v15 = int32(0)
		}
		*(*int32)(unsafe.Add(mBase, uint32(l0)+16)) = v15
	} else {
	}
	return int32(0)
}
func F_hash_ok_operator(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v2 int32
	_ = v2
	var v5 int32
	_ = v5
	var v7 int32
	_ = v7
	var v9 int32
	_ = v9
	var v12 int32
	_ = v12
	var v15 int32
	_ = v15
	var v27 int32
	_ = v27
	var v28 int32
	_ = v28
	var v29 int32
	_ = v29
	var v32 int32
	_ = v32
	var v33 int32
	_ = v33
	var v34 int32
	_ = v34
	var v37 int32
	_ = v37
	var v38 int32
	_ = v38
	var v41 int32
	_ = v41
	var v42 int32
	_ = v42
	var v43 int32
	_ = v43
	var v44 int32
	_ = v44
	var v47 int32
	_ = v47
	var v48 int32
	_ = v48
	var v49 int32
	_ = v49
	var v51 int32
	_ = v51
	var v53 int32
	_ = v53
	var v57 int32
	_ = v57
	var v65 int32
	_ = v65
	var v69 int32
	_ = v69
	var v74 int32
	_ = v74
	v2 = int32(0)
	v5 = m.G0
	v7 = v5 - int32(16)
	m.G0 = v7
	v9 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
	if v9 == v2 {
		v57 = v2
		m.G0 = v7 + int32(16)
		return v57
	} else {
		v12 = *(*int32)(unsafe.Add(mBase, uint32(v9)+4))
		if v12 != int32(2) {
			v57 = v2
			m.G0 = v7 + int32(16)
			return v57
		} else {
			v15 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
			if v15 <= int32(2987) {
				if base.B2i32(v15 == int32(1070))|base.B2i32(v15 == int32(2860)) != 0 {
					v27 = *(*int32)(unsafe.Add(mBase, uint32(v9)+12))
					v28 = *(*int32)(unsafe.Add(mBase, uint32(v27)))
					v29 = F_exprType(m, v28)
					mBase = m.M
					v32 = m.ExcPending
					if v32 != 0 {
						return int32(0)
					} else {
						v33 = F_op_hashjoinable(m, v15, v29)
						mBase = m.M
						v34 = m.ExcPending
						if v34 != 0 {
							return int32(0)
						} else {
							v57 = v33
							m.G0 = v7 + int32(16)
							return v57
						}
					}
				} else {
					v37 = F_SearchSysCache1(m, int32(40), base.I64_extend_i32_u(v15))
					mBase = m.M
					v38 = m.ExcPending
					if v38 != 0 {
						return int32(0)
					} else {
						if v37 == int32(0) {
							F_errstart_cold(m, int32(21), int32(0))
							mBase = m.M
							v65 = m.ExcPending
							if v65 != 0 {
								return int32(0)
							} else {
								*(*int32)(unsafe.Add(mBase, uint32(v7))) = v15
								F_errmsg_internal(m, int32(_a_F_hash_ok_operator_0), v7)
								mBase = m.M
								v69 = m.ExcPending
								if v69 != 0 {
									return int32(0)
								} else {
									F_errfinish(m, int32(_a_F_hash_ok_operator_1), int32(866), int32(_a_F_hash_ok_operator_2))
									mBase = m.M
									v74 = m.ExcPending
									if v74 != 0 {
										return int32(0)
									} else {
										base.Wasm_trap_unreachable()
										for {
										}
									}
								}
							}
						} else {
							v41 = *(*int32)(unsafe.Add(mBase, uint32(v37)+16))
							v42 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v41)+22)))
							v43 = v41 + v42
							v44 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v43)+78)))
							if v44 == int32(1) {
								v47 = *(*int32)(unsafe.Add(mBase, uint32(v43)+100))
								v48 = F_func_strict(m, v47)
								mBase = m.M
								v49 = m.ExcPending
								if v49 != 0 {
									return int32(0)
								} else {
									if v48 != 0 {
										F_ReleaseCatCache(m, v37)
										mBase = m.M
										v53 = m.ExcPending
										if v53 != 0 {
											return int32(0)
										} else {
											v57 = int32(1)
											m.G0 = v7 + int32(16)
											return v57
										}
									} else {
										F_ReleaseCatCache(m, v37)
										mBase = m.M
										v51 = m.ExcPending
										if v51 != 0 {
											return int32(0)
										} else {
											v57 = v2
											m.G0 = v7 + int32(16)
											return v57
										}
									}
								}
							} else {
								F_ReleaseCatCache(m, v37)
								mBase = m.M
								v51 = m.ExcPending
								if v51 != 0 {
									return int32(0)
								} else {
									v57 = v2
									m.G0 = v7 + int32(16)
									return v57
								}
							}
						}
					}
				}
			} else {
				if v15 == int32(3882) {
					v27 = *(*int32)(unsafe.Add(mBase, uint32(v9)+12))
					v28 = *(*int32)(unsafe.Add(mBase, uint32(v27)))
					v29 = F_exprType(m, v28)
					mBase = m.M
					v32 = m.ExcPending
					if v32 != 0 {
						return int32(0)
					} else {
						v33 = F_op_hashjoinable(m, v15, v29)
						mBase = m.M
						v34 = m.ExcPending
						if v34 != 0 {
							return int32(0)
						} else {
							v57 = v33
							m.G0 = v7 + int32(16)
							return v57
						}
					}
				} else {
					if v15 != int32(2988) {
						v37 = F_SearchSysCache1(m, int32(40), base.I64_extend_i32_u(v15))
						mBase = m.M
						v38 = m.ExcPending
						if v38 != 0 {
							return int32(0)
						} else {
							if v37 == int32(0) {
								F_errstart_cold(m, int32(21), int32(0))
								mBase = m.M
								v65 = m.ExcPending
								if v65 != 0 {
									return int32(0)
								} else {
									*(*int32)(unsafe.Add(mBase, uint32(v7))) = v15
									F_errmsg_internal(m, int32(_a_F_hash_ok_operator_0), v7)
									mBase = m.M
									v69 = m.ExcPending
									if v69 != 0 {
										return int32(0)
									} else {
										F_errfinish(m, int32(_a_F_hash_ok_operator_1), int32(866), int32(_a_F_hash_ok_operator_2))
										mBase = m.M
										v74 = m.ExcPending
										if v74 != 0 {
											return int32(0)
										} else {
											base.Wasm_trap_unreachable()
											for {
											}
										}
									}
								}
							} else {
								v41 = *(*int32)(unsafe.Add(mBase, uint32(v37)+16))
								v42 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v41)+22)))
								v43 = v41 + v42
								v44 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v43)+78)))
								if v44 == int32(1) {
									v47 = *(*int32)(unsafe.Add(mBase, uint32(v43)+100))
									v48 = F_func_strict(m, v47)
									mBase = m.M
									v49 = m.ExcPending
									if v49 != 0 {
										return int32(0)
									} else {
										if v48 != 0 {
											F_ReleaseCatCache(m, v37)
											mBase = m.M
											v53 = m.ExcPending
											if v53 != 0 {
												return int32(0)
											} else {
												v57 = int32(1)
												m.G0 = v7 + int32(16)
												return v57
											}
										} else {
											F_ReleaseCatCache(m, v37)
											mBase = m.M
											v51 = m.ExcPending
											if v51 != 0 {
												return int32(0)
											} else {
												v57 = v2
												m.G0 = v7 + int32(16)
												return v57
											}
										}
									}
								} else {
									F_ReleaseCatCache(m, v37)
									mBase = m.M
									v51 = m.ExcPending
									if v51 != 0 {
										return int32(0)
									} else {
										v57 = v2
										m.G0 = v7 + int32(16)
										return v57
									}
								}
							}
						}
					} else {
						v27 = *(*int32)(unsafe.Add(mBase, uint32(v9)+12))
						v28 = *(*int32)(unsafe.Add(mBase, uint32(v27)))
						v29 = F_exprType(m, v28)
						mBase = m.M
						v32 = m.ExcPending
						if v32 != 0 {
							return int32(0)
						} else {
							v33 = F_op_hashjoinable(m, v15, v29)
							mBase = m.M
							v34 = m.ExcPending
							if v34 != 0 {
								return int32(0)
							} else {
								v57 = v33
								m.G0 = v7 + int32(16)
								return v57
							}
						}
					}
				}
			}
		}
	}
}
func F_lookup_hash_entries(m *base.Module, l0 int32) {
	mBase := m.M
	_ = mBase
	var v2 int32
	_ = v2
	var v12 int32
	_ = v12
	var v14 int32
	_ = v14
	var v16 int32
	_ = v16
	var v19 int32
	_ = v19
	var v20 int32
	_ = v20
	var v21 int32
	_ = v21
	var v28 int32
	_ = v28
	var v33 int32
	_ = v33
	var v36 int32
	_ = v36
	var v37 int32
	_ = v37
	var v38 int32
	_ = v38
	var v39 int32
	_ = v39
	var v42 int32
	_ = v42
	var v44 int32
	_ = v44
	var v45 int32
	_ = v45
	var v46 int32
	_ = v46
	var v48 int32
	_ = v48
	var v49 int32
	_ = v49
	var v51 int32
	_ = v51
	var v55 int32
	_ = v55
	var v56 int32
	_ = v56
	var v57 int32
	_ = v57
	var v59 int32
	_ = v59
	var v60 int32
	_ = v60
	var v66 int32
	_ = v66
	var v75 int32
	_ = v75
	var v76 int32
	_ = v76
	var v79 int32
	_ = v79
	var v80 int32
	_ = v80
	var v81 int32
	_ = v81
	var v84 int32
	_ = v84
	var v86 int32
	_ = v86
	var v90 int64
	_ = v90
	var v92 int32
	_ = v92
	var v94 int32
	_ = v94
	var v96 int32
	_ = v96
	var v99 int32
	_ = v99
	var v100 int32
	_ = v100
	var v113 int32
	_ = v113
	var v115 int32
	_ = v115
	var v117 int32
	_ = v117
	var v118 int32
	_ = v118
	var v125 int32
	_ = v125
	var v126 int32
	_ = v126
	var v127 int32
	_ = v127
	var v131 int32
	_ = v131
	var v132 int32
	_ = v132
	var v135 int32
	_ = v135
	var v137 int32
	_ = v137
	var v138 int32
	_ = v138
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
	var v149 float64
	_ = v149
	var v150 float64
	_ = v150
	var v152 int32
	_ = v152
	var v153 int32
	_ = v153
	var v155 int32
	_ = v155
	var v163 int32
	_ = v163
	var v166 int32
	_ = v166
	var v167 int32
	_ = v167
	v2 = int32(0)
	v12 = m.G0
	v14 = v12 - int32(16)
	m.G0 = v14
	v16 = *(*int32)(unsafe.Add(mBase, uint32(l0)+244))
	if v2 < v16 {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	v19 = *(*int32)(unsafe.Add(mBase, uint32(l0)+344))
	v20 = *(*int32)(unsafe.Add(mBase, uint32(l0)+164))
	v21 = *(*int32)(unsafe.Add(mBase, uint32(v20)+12))
	v28 = v2
	goto L4
L2:
	;
	goto L3
L3:
	;
	m.G0 = v14 + int32(16)
	return
L4:
	;
	v33 = *(*int32)(unsafe.Add(mBase, uint32(l0)+340))
	v36 = v33 + v28*int32(52)
	v37 = *(*int32)(unsafe.Add(mBase, uint32(v36)+16))
	v38 = *(*int32)(unsafe.Add(mBase, uint32(v36)))
	v39 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v14)+11)) = uint8(v39)
	*(*int32)(unsafe.Add(mBase, uint32(l0)+188)) = v28
	v42 = *(*int32)(unsafe.Add(mBase, uint32(l0)+156))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+168)) = v42
	v44 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+277)))
	v45 = *(*int32)(unsafe.Add(mBase, uint32(v36)+36))
	v46 = int32(*(*int16)(unsafe.Add(mBase, uint32(v21)+6)))
	if v46 < v45 {
		goto L6
	} else {
		goto L7
	}
L5:
	;
	goto L3
L6:
	;
	v48 = *(*int32)(unsafe.Add(mBase, uint32(v21)+8))
	v49 = *(*int32)(unsafe.Add(mBase, uint32(v48)+16))
	m.T0[v49].(func(*base.Module, int32, int32))(m, v21, v45)
	mBase = m.M
	v51 = m.ExcPending
	if v51 != 0 {
		goto L9
	} else {
		goto L10
	}
L7:
	;
	goto L8
L8:
	;
	if v44 != 0 {
		goto L11
	} else {
		goto L12
	}
L9:
	;
	return
L10:
	;
	goto L8
L11:
	;
	v55 = int32(0)
	goto L13
L12:
	;
	v55 = v14 + int32(11)
	goto L13
L13:
	;
	v56 = *(*int32)(unsafe.Add(mBase, uint32(v37)+8))
	v57 = *(*int32)(unsafe.Add(mBase, uint32(v56)+12))
	m.T0[v57].(func(*base.Module, int32))(m, v37)
	mBase = m.M
	v59 = m.ExcPending
	if v59 != 0 {
		goto L9
	} else {
		goto L14
	}
L14:
	;
	v60 = *(*int32)(unsafe.Add(mBase, uint32(v36)+32))
	if int32(0) < v60 {
		goto L15
	} else {
		goto L16
	}
L15:
	;
	v66 = int32(0)
	goto L18
L16:
	;
	goto L17
L17:
	;
	v113 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v37)+4)))
	v115 = v113 & int32(_a_F_lookup_hash_entries_0)
	*(*uint16)(unsafe.Add(mBase, uint32(v37)+4)) = uint16(v115)
	v117 = *(*int32)(unsafe.Add(mBase, uint32(v37)+12))
	v118 = *(*int32)(unsafe.Add(mBase, uint32(v117)))
	*(*uint16)(unsafe.Add(mBase, uint32(v37)+6)) = uint16(v118)
	goto L21
L18:
	;
	v75 = *(*int32)(unsafe.Add(mBase, uint32(v37)+16))
	v76 = int32(3)
	v79 = *(*int32)(unsafe.Add(mBase, uint32(v21)+16))
	v80 = *(*int32)(unsafe.Add(mBase, uint32(v36)+40))
	v81 = int32(1)
	v84 = int32(*(*int16)(unsafe.Add(mBase, uint32(v80+v66<<(uint(v81)%32)))))
	v86 = v84 - v81
	v90 = *(*int64)(unsafe.Add(mBase, uint32(v79+v86<<(uint(v76)%32))))
	*(*int64)(unsafe.Add(mBase, uint32(v75+v66<<(uint(v76)%32)))) = v90
	v92 = *(*int32)(unsafe.Add(mBase, uint32(v37)+20))
	v94 = *(*int32)(unsafe.Add(mBase, uint32(v21)+20))
	v96 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v94+v86))))
	*(*uint8)(unsafe.Add(mBase, uint32(v92+v66))) = uint8(v96)
	v99 = v66 + v81
	v100 = *(*int32)(unsafe.Add(mBase, uint32(v36)+32))
	if v99 < v100 {
		v66 = v99
		goto L18
	} else {
		goto L20
	}
L19:
	;
	goto L17
L20:
	;
	goto L19
L21:
	;
	v125 = F_LookupTupleHashEntry(m, v38, v37, v55, v14+int32(12))
	mBase = m.M
	v126 = m.ExcPending
	if v126 != 0 {
		goto L9
	} else {
		goto L24
	}
L22:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v19+v28<<(uint(int32(2))%32)))) = v163
	v166 = v28 + int32(1)
	v167 = *(*int32)(unsafe.Add(mBase, uint32(l0)+244))
	if v166 < v167 {
		v28 = v166
		goto L4
	} else {
		goto L38
	}
L23:
	;
	v163 = int32(0)
	goto L22
L24:
	;
	if v125 != 0 {
		goto L25
	} else {
		goto L26
	}
L25:
	;
	v127 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v14)+11)))
	if v127 == int32(1) {
		goto L28
	} else {
		goto L29
	}
L26:
	;
	goto L27
L27:
	;
	v137 = *(*int32)(unsafe.Add(mBase, uint32(l0)+164))
	v138 = *(*int32)(unsafe.Add(mBase, uint32(v137)+12))
	v139 = *(*int32)(unsafe.Add(mBase, uint32(l0)+260))
	v142 = v139 + v28*int32(24)
	v143 = *(*int32)(unsafe.Add(mBase, uint32(v142)+4))
	if v143 == int32(0) {
		goto L33
	} else {
		goto L34
	}
L28:
	;
	F_initialize_hash_entry(m, l0, v38, v125)
	mBase = m.M
	v131 = m.ExcPending
	if v131 != 0 {
		goto L9
	} else {
		goto L31
	}
L29:
	;
	goto L30
L30:
	;
	v132 = *(*int32)(unsafe.Add(mBase, uint32(v38)+32))
	if v132 == int32(0) {
		goto L23
	} else {
		goto L32
	}
L31:
	;
	goto L30
L32:
	;
	v135 = *(*int32)(unsafe.Add(mBase, uint32(v125)))
	v163 = v135 - v132
	goto L22
L33:
	;
	v146 = *(*int32)(unsafe.Add(mBase, uint32(l0)+256))
	v148 = *(*int32)(unsafe.Add(mBase, uint32(v36)+48))
	v149 = *(*float64)(unsafe.Add(mBase, uint32(v148)+96))
	v150 = *(*float64)(unsafe.Add(mBase, uint32(l0)+304))
	F_hashagg_spill_init(m, v142, v146, int32(0), v149, v150)
	mBase = m.M
	v152 = m.ExcPending
	if v152 != 0 {
		goto L9
	} else {
		goto L36
	}
L34:
	;
	goto L35
L35:
	;
	v153 = *(*int32)(unsafe.Add(mBase, uint32(v14)+12))
	F_hashagg_spill_tuple(m, l0, v142, v138, v153)
	mBase = m.M
	v155 = m.ExcPending
	if v155 != 0 {
		goto L9
	} else {
		goto L37
	}
L36:
	;
	goto L35
L37:
	;
	goto L23
L38:
	;
	goto L5
}
