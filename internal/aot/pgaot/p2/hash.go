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
	var v166 int32
	_ = v166
	var v167 int32
	_ = v167
	var v168 int32
	_ = v168
	var v169 int32
	_ = v169
	var v170 int32
	_ = v170
	var v172 int32
	_ = v172
	var v174 int32
	_ = v174
	var v176 int32
	_ = v176
	var v178 int32
	_ = v178
	var v179 int32
	_ = v179
	var v180 int32
	_ = v180
	var v181 int32
	_ = v181
	var v184 int32
	_ = v184
	var v189 int32
	_ = v189
	var v191 int32
	_ = v191
	var v196 int32
	_ = v196
	var v201 int32
	_ = v201
	var v208 int32
	_ = v208
	var v219 int32
	_ = v219
	var v224 int32
	_ = v224
	var v225 int32
	_ = v225
	var v228 int32
	_ = v228
	var v229 int32
	_ = v229
	var v233 int32
	_ = v233
	var v239 int32
	_ = v239
	var v241 int32
	_ = v241
	var v247 int32
	_ = v247
	var v248 int32
	_ = v248
	var v254 int32
	_ = v254
	var v259 int32
	_ = v259
	var v261 int32
	_ = v261
	var v262 int32
	_ = v262
	var v266 int32
	_ = v266
	var v269 int32
	_ = v269
	var v272 int32
	_ = v272
	var v273 int32
	_ = v273
	var v274 int32
	_ = v274
	var v286 int32
	_ = v286
	var v289 int32
	_ = v289
	var v294 int32
	_ = v294
	var v311 int32
	_ = v311
	var v312 int32
	_ = v312
	var v318 int32
	_ = v318
	var v324 int32
	_ = v324
	var v326 int32
	_ = v326
	var v330 int32
	_ = v330
	var v331 int32
	_ = v331
	var v337 int32
	_ = v337
	var v343 int32
	_ = v343
	var v344 int32
	_ = v344
	var v345 int32
	_ = v345
	var v347 int32
	_ = v347
	var v356 int32
	_ = v356
	var v359 int32
	_ = v359
	var v381 int32
	_ = v381
	var v382 int32
	_ = v382
	var v388 int32
	_ = v388
	var v402 int32
	_ = v402
	var v424 int32
	_ = v424
	var v425 int32
	_ = v425
	var v426 int32
	_ = v426
	var v429 int32
	_ = v429
	var v430 int32
	_ = v430
	var v432 int32
	_ = v432
	var v437 int32
	_ = v437
	var v438 int32
	_ = v438
	var v439 int32
	_ = v439
	var v440 int32
	_ = v440
	var v442 int32
	_ = v442
	var v444 int32
	_ = v444
	var v445 int32
	_ = v445
	var v449 int32
	_ = v449
	var v453 int32
	_ = v453
	var v455 int32
	_ = v455
	var v459 int32
	_ = v459
	var v460 float64
	_ = v460
	var v465 int32
	_ = v465
	var v467 int32
	_ = v467
	var v468 int32
	_ = v468
	var v469 int32
	_ = v469
	var v473 int32
	_ = v473
	var v476 int32
	_ = v476
	var v477 int32
	_ = v477
	var v482 int32
	_ = v482
	var v483 int32
	_ = v483
	var v488 int32
	_ = v488
	var v491 int32
	_ = v491
	var v496 int32
	_ = v496
	var v497 int32
	_ = v497
	var v500 int32
	_ = v500
	var v501 int32
	_ = v501
	var v507 int32
	_ = v507
	var v511 int32
	_ = v511
	var v516 int32
	_ = v516
	var v522 int32
	_ = v522
	var v526 int32
	_ = v526
	var v529 int64
	_ = v529
	var v530 int32
	_ = v530
	var v534 int32
	_ = v534
	var v540 int32
	_ = v540
	var v542 int32
	_ = v542
	var v548 int32
	_ = v548
	var v549 int64
	_ = v549
	var v559 int32
	_ = v559
	var v563 int32
	_ = v563
	var v565 int32
	_ = v565
	var v569 int32
	_ = v569
	var v577 int32
	_ = v577
	var v579 int32
	_ = v579
	var v585 int32
	_ = v585
	var v613 int32
	_ = v613
	var v614 int32
	_ = v614
	var v615 int32
	_ = v615
	var v616 int32
	_ = v616
	var v619 int32
	_ = v619
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
	var v666 int32
	_ = v666
	var v672 int32
	_ = v672
	var v674 int32
	_ = v674
	var v682 int32
	_ = v682
	var v684 int32
	_ = v684
	var v685 int32
	_ = v685
	var v689 int32
	_ = v689
	var v695 int32
	_ = v695
	var v697 int32
	_ = v697
	var v703 int32
	_ = v703
	var v704 int32
	_ = v704
	var v705 int32
	_ = v705
	var v707 int32
	_ = v707
	var v708 int32
	_ = v708
	var v709 int32
	_ = v709
	var v710 int32
	_ = v710
	var v713 int32
	_ = v713
	var v721 int32
	_ = v721
	var v746 int32
	_ = v746
	var v747 int32
	_ = v747
	var v749 int32
	_ = v749
	var v753 int32
	_ = v753
	var v755 int32
	_ = v755
	var v759 int32
	_ = v759
	var v763 int32
	_ = v763
	var v769 int32
	_ = v769
	var v771 int32
	_ = v771
	var v777 int32
	_ = v777
	var v778 int32
	_ = v778
	var v779 int32
	_ = v779
	var v788 int32
	_ = v788
	var v791 int32
	_ = v791
	var v794 int32
	_ = v794
	var v796 int32
	_ = v796
	var v802 int32
	_ = v802
	var v804 int32
	_ = v804
	var v813 int32
	_ = v813
	var v822 int32
	_ = v822
	var v823 int32
	_ = v823
	var v827 int32
	_ = v827
	var v833 int32
	_ = v833
	var v837 int32
	_ = v837
	var v840 int32
	_ = v840
	var v843 int32
	_ = v843
	var v846 int32
	_ = v846
	var v848 int32
	_ = v848
	var v849 int32
	_ = v849
	var v850 int32
	_ = v850
	var v855 int32
	_ = v855
	var v862 int32
	_ = v862
	var v869 int32
	_ = v869
	var v871 int32
	_ = v871
	var v873 int32
	_ = v873
	var v874 int32
	_ = v874
	var v880 int32
	_ = v880
	var v881 int32
	_ = v881
	var v887 int32
	_ = v887
	var v892 int32
	_ = v892
	var v897 int32
	_ = v897
	var v898 float64
	_ = v898
	var v900 float64
	_ = v900
	var v902 int32
	_ = v902
	var v903 int32
	_ = v903
	var v905 int32
	_ = v905
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
	var v949 int32
	_ = v949
	var v955 int32
	_ = v955
	var v957 int32
	_ = v957
	var v963 int32
	_ = v963
	var v964 int64
	_ = v964
	var v974 int32
	_ = v974
	var v980 int32
	_ = v980
	var v982 int32
	_ = v982
	var v988 int32
	_ = v988
	var v996 int32
	_ = v996
	var v998 int32
	_ = v998
	var v1004 int32
	_ = v1004
	var v1006 int32
	_ = v1006
	var v1009 int32
	_ = v1009
	var v1011 int32
	_ = v1011
	var v1015 int32
	_ = v1015
	var v1018 int32
	_ = v1018
	var v1020 int32
	_ = v1020
	var v1050 int32
	_ = v1050
	var v1053 int32
	_ = v1053
	var v1057 int32
	_ = v1057
	var v1061 int32
	_ = v1061
	var v1063 int32
	_ = v1063
	var v1067 int32
	_ = v1067
	var v1068 int32
	_ = v1068
	var v1071 float64
	_ = v1071
	var v1072 int32
	_ = v1072
	var v1075 int32
	_ = v1075
	var v1079 int32
	_ = v1079
	var v1080 int32
	_ = v1080
	var v1081 int32
	_ = v1081
	var v1082 int32
	_ = v1082
	var v1086 int32
	_ = v1086
	var v1092 int32
	_ = v1092
	var v1095 int32
	_ = v1095
	var v1105 int32
	_ = v1105
	var v1109 int32
	_ = v1109
	var v1113 int32
	_ = v1113
	var v1116 int32
	_ = v1116
	var v1117 int32
	_ = v1117
	var v1118 int32
	_ = v1118
	var v1119 int32
	_ = v1119
	var v1123 int32
	_ = v1123
	var v1127 int32
	_ = v1127
	var v1131 int32
	_ = v1131
	var v1135 int32
	_ = v1135
	var v1140 int32
	_ = v1140
	var v1141 int32
	_ = v1141
	var v1147 int32
	_ = v1147
	var v1153 int32
	_ = v1153
	var v1155 int32
	_ = v1155
	var v1161 int32
	_ = v1161
	var v1162 int32
	_ = v1162
	var v1164 int32
	_ = v1164
	var v1167 int32
	_ = v1167
	var v1168 int32
	_ = v1168
	var v1169 int32
	_ = v1169
	var v1172 int32
	_ = v1172
	var v1175 int32
	_ = v1175
	var v1177 int32
	_ = v1177
	var v1179 int32
	_ = v1179
	var v1182 int32
	_ = v1182
	var v1183 int32
	_ = v1183
	var v1184 int32
	_ = v1184
	var v1187 int32
	_ = v1187
	var v1188 int32
	_ = v1188
	var v1195 int32
	_ = v1195
	var v1197 int32
	_ = v1197
	var v1198 int32
	_ = v1198
	var v1199 int32
	_ = v1199
	var v1203 int32
	_ = v1203
	var v1209 int32
	_ = v1209
	var v1212 int32
	_ = v1212
	var v1222 int32
	_ = v1222
	var v1223 int32
	_ = v1223
	var v1226 int32
	_ = v1226
	var v1227 int32
	_ = v1227
	var v1228 int32
	_ = v1228
	var v1229 int32
	_ = v1229
	var v1233 int32
	_ = v1233
	var v1239 int32
	_ = v1239
	var v1242 int32
	_ = v1242
	var v1252 int32
	_ = v1252
	var v1253 int32
	_ = v1253
	var v1262 int32
	_ = v1262
	var v1263 int32
	_ = v1263
	var v1265 int32
	_ = v1265
	var v1276 int32
	_ = v1276
	var v1277 int32
	_ = v1277
	var v1285 int32
	_ = v1285
	var v1288 int32
	_ = v1288
	var v1304 int32
	_ = v1304
	var v1306 int32
	_ = v1306
	var v1308 int32
	_ = v1308
	var v1313 int32
	_ = v1313
	var v1327 int32
	_ = v1327
	var v1333 int32
	_ = v1333
	var v1336 int32
	_ = v1336
	var v1337 int32
	_ = v1337
	var v1342 int32
	_ = v1342
	var v1343 int32
	_ = v1343
	var v1347 int32
	_ = v1347
	var v1350 int32
	_ = v1350
	var v1351 int32
	_ = v1351
	var v1357 int32
	_ = v1357
	var v1360 int32
	_ = v1360
	var v1363 int32
	_ = v1363
	var v1366 int32
	_ = v1366
	var v1368 int32
	_ = v1368
	var v1369 int32
	_ = v1369
	var v1370 int32
	_ = v1370
	var v1372 int64
	_ = v1372
	var v1376 int32
	_ = v1376
	var v1377 int32
	_ = v1377
	var v1380 int32
	_ = v1380
	var v1381 int32
	_ = v1381
	var v1382 int32
	_ = v1382
	var v1384 int32
	_ = v1384
	var v1386 int32
	_ = v1386
	var v1388 int32
	_ = v1388
	var v1392 int32
	_ = v1392
	var v1394 int32
	_ = v1394
	var v1395 int32
	_ = v1395
	var v1400 int32
	_ = v1400
	var v1406 int32
	_ = v1406
	var v1407 int32
	_ = v1407
	var v1408 int32
	_ = v1408
	var v1409 int32
	_ = v1409
	var v1413 int32
	_ = v1413
	var v1415 int32
	_ = v1415
	var v1417 int32
	_ = v1417
	var v1418 int32
	_ = v1418
	var v1420 int32
	_ = v1420
	var v1426 int32
	_ = v1426
	var v1427 int32
	_ = v1427
	var v1428 int32
	_ = v1428
	var v1435 int32
	_ = v1435
	var v1436 int32
	_ = v1436
	var v1437 int32
	_ = v1437
	var v1441 int32
	_ = v1441
	var v1442 int32
	_ = v1442
	var v1448 int32
	_ = v1448
	var v1453 int32
	_ = v1453
	var v1454 int32
	_ = v1454
	var v1455 int32
	_ = v1455
	var v1456 int32
	_ = v1456
	var v1460 int32
	_ = v1460
	var v1466 int32
	_ = v1466
	var v1468 int32
	_ = v1468
	var v1474 int32
	_ = v1474
	var v1475 int32
	_ = v1475
	var v1476 int32
	_ = v1476
	var v1478 int32
	_ = v1478
	var v1480 int32
	_ = v1480
	var v1483 int32
	_ = v1483
	var v1487 int32
	_ = v1487
	var v1493 int32
	_ = v1493
	var v1495 int32
	_ = v1495
	var v1501 int32
	_ = v1501
	var v1502 int32
	_ = v1502
	var v1503 int32
	_ = v1503
	var v1511 int32
	_ = v1511
	var v1512 int32
	_ = v1512
	var v1513 int32
	_ = v1513
	var v1517 int32
	_ = v1517
	var v1520 int32
	_ = v1520
	var v1521 int32
	_ = v1521
	var v1523 int32
	_ = v1523
	var v1525 int32
	_ = v1525
	var v1526 int32
	_ = v1526
	var v1530 int32
	_ = v1530
	var v1534 int32
	_ = v1534
	var v1538 int32
	_ = v1538
	var v1539 int32
	_ = v1539
	var v1543 int32
	_ = v1543
	var v1546 int32
	_ = v1546
	var v1551 int32
	_ = v1551
	var v1555 int32
	_ = v1555
	var v1557 int32
	_ = v1557
	var v1564 int32
	_ = v1564
	var v1565 int32
	_ = v1565
	var v1566 int32
	_ = v1566
	var v1574 int32
	_ = v1574
	var v1579 int32
	_ = v1579
	var v1582 int64
	_ = v1582
	var v1583 int32
	_ = v1583
	var v1587 int32
	_ = v1587
	var v1593 int32
	_ = v1593
	var v1595 int32
	_ = v1595
	var v1601 int32
	_ = v1601
	var v1602 int64
	_ = v1602
	var v1607 int32
	_ = v1607
	var v1608 int32
	_ = v1608
	var v1612 int32
	_ = v1612
	var v1618 int32
	_ = v1618
	var v1620 int32
	_ = v1620
	var v1626 int32
	_ = v1626
	var v1632 int32
	_ = v1632
	var v1636 int32
	_ = v1636
	var v1638 int32
	_ = v1638
	var v1642 int32
	_ = v1642
	var v1649 int32
	_ = v1649
	var v1651 int32
	_ = v1651
	var v1657 int32
	_ = v1657
	var v1660 int32
	_ = v1660
	var v1662 int32
	_ = v1662
	var v1664 int32
	_ = v1664
	var v1667 int32
	_ = v1667
	var v1679 int32
	_ = v1679
	var v1727 int32
	_ = v1727
	var v1734 int32
	_ = v1734
	var v1737 int32
	_ = v1737
	var v1738 int32
	_ = v1738
	var v1749 int32
	_ = v1749
	var v1753 int32
	_ = v1753
	var v1758 int32
	_ = v1758
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
	v1734 = m.ExcPending
	if v1734 != 0 {
		goto L9
	} else {
		goto L394
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
	v178 = int32(4)
	v179 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v151)+14)))
	v180 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v151)+12)))
	v181 = v179 - v180
	if v181 <= v178 {
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
	v111 = F__hash_getbucketbuf_from_hashkey(m, l0, v41, int32(2), v30+int32(20))
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
	v122 = *(*int32)(unsafe.Add(mBase, uint32(v116+(v111^int32(-1))<<(uint(int32(6))%32))+16))
	v131 = v122
	goto L16
L18:
	;
	goto L19
L19:
	;
	v124 = *(*int32)(unsafe.Add(mBase, _c_F__hash_doinsert[3]))
	v130 = *(*int32)(unsafe.Add(mBase, uint32(v124+v111<<(uint(int32(6))%32)+int32(-64))+16))
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
	F_LockBuffer(m, v111, int32(0))
	mBase = m.M
	v166 = m.ExcPending
	if v166 != 0 {
		goto L9
	} else {
		goto L29
	}
L29:
	;
	v167 = *(*int32)(unsafe.Add(mBase, uint32(v30)+20))
	v168 = *(*int32)(unsafe.Add(mBase, uint32(v167)+24))
	v169 = *(*int32)(unsafe.Add(mBase, uint32(v167)+28))
	v170 = *(*int32)(unsafe.Add(mBase, uint32(v167)+32))
	F__hash_finish_split(m, l0, v79, v111, v159, v168, v169, v170)
	mBase = m.M
	v172 = m.ExcPending
	if v172 != 0 {
		goto L9
	} else {
		goto L30
	}
L30:
	;
	F_ReleaseBuffer(m, v111)
	mBase = m.M
	v174 = m.ExcPending
	if v174 != 0 {
		goto L9
	} else {
		goto L31
	}
L31:
	;
	F_ReleaseBuffer(m, v79)
	mBase = m.M
	v176 = m.ExcPending
	if v176 != 0 {
		goto L9
	} else {
		goto L32
	}
L32:
	;
	goto L6
L33:
	;
	F_LockBuffer(m, v79, int32(2))
	mBase = m.M
	v746 = m.ExcPending
	if v746 != 0 {
		goto L9
	} else {
		goto L144
	}
L34:
	;
	if base.Ui32(v48) <= base.Ui32(v184-int32(4)) {
		v721 = v111
		goto L33
	} else {
		goto L38
	}
L35:
	;
	v184 = v178
	goto L37
L36:
	;
	v184 = v181
	goto L37
L37:
	;
	goto L34
L38:
	;
	v189 = v79 ^ int32(-1)
	v191 = v79 << (uint(int32(13)) % 32)
	v196 = v111
	v201 = v151
	v208 = v153
	goto L39
L39:
	;
	v219 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v208)+12)))
	if v219&int32(128) == int32(0) {
		goto L41
	} else {
		goto L42
	}
L40:
	;
	v721 = v703
	goto L33
L41:
	;
	v650 = *(*int32)(unsafe.Add(mBase, uint32(v208)+4))
	if v650 != int32(-1) {
		goto L121
	} else {
		goto L122
	}
L42:
	;
	v224 = F_IsBufferCleanupOK(m, v196)
	mBase = m.M
	v225 = m.ExcPending
	if v225 != 0 {
		goto L9
	} else {
		goto L43
	}
L43:
	;
	if v224 == int32(0) {
		goto L41
	} else {
		goto L44
	}
L44:
	;
	v228 = int32(0)
	v229 = base.B2i32(v228 <= v196)
	if v229 == v228 {
		goto L47
	} else {
		goto L48
	}
L45:
	;
	v613 = int32(4)
	v614 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v201)+14)))
	v615 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v201)+12)))
	v616 = v614 - v615
	if v616 <= v613 {
		goto L116
	} else {
		goto L117
	}
L46:
	;
	v248 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v247)+12)))
	if base.Ui32(v248) < base.Ui32(int32(25)) {
		goto L45
	} else {
		goto L50
	}
L47:
	;
	v233 = *(*int32)(unsafe.Add(mBase, _c_F__hash_doinsert[0]))
	v239 = *(*int32)(unsafe.Add(mBase, uint32(v233+(v196^int32(-1))<<(uint(int32(2))%32))))
	v247 = v239
	goto L46
L48:
	;
	goto L49
L49:
	;
	v241 = *(*int32)(unsafe.Add(mBase, _c_F__hash_doinsert[1]))
	v247 = v241 + v196<<(uint(int32(13))%32) + int32(-8192)
	goto L46
L50:
	;
	v254 = int32(base.Ui32(v248+int32(_a_F__hash_doinsert_3)) >> (uint(int32(2)) % 32))
	if v254&int32(_a_F__hash_doinsert_4) == int32(0) {
		goto L45
	} else {
		goto L51
	}
L51:
	;
	v259 = int32(1)
	v261 = v247 + int32(20)
	v262 = int32(0)
	v266 = (v254 + v259) & int32(_a_F__hash_doinsert_4)
	if base.Ui32(int32(3)) <= base.Ui32(v266) {
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
	v269 = int32(2)
	if base.Ui32(v266) <= base.Ui32(v269) {
		goto L56
	} else {
		goto L57
	}
L54:
	;
	v356 = v259
	v359 = v262
	goto L55
L55:
	;
	v381 = *(*int32)(unsafe.Add(mBase, uint32(v261+v356<<(uint(int32(2))%32))))
	v382 = int32(_a_F__hash_doinsert_5)
	if v381&v382 != v382 {
		v402 = v359
		goto L52
	} else {
		goto L69
	}
L56:
	;
	v272 = v269
	goto L58
L57:
	;
	v272 = v266
	goto L58
L58:
	;
	v273 = int32(1)
	v274 = v272 - v273
	v286 = v273
	v289 = v262
	v294 = int32(0)
	goto L59
L59:
	;
	v311 = *(*int32)(unsafe.Add(mBase, uint32(v261+v286<<(uint(int32(2))%32))))
	v312 = int32(_a_F__hash_doinsert_5)
	if v311&v312 == v312 {
		goto L61
	} else {
		goto L62
	}
L60:
	;
	if v274&v273 == int32(0) {
		v402 = v343
		goto L52
	} else {
		goto L68
	}
L61:
	;
	v318 = int32(1)
	*(*uint16)(unsafe.Add(mBase, uint32(v30+int32(32)+v289<<(uint(v318)%32)))) = uint16(v286)
	v324 = v289 + v318
	goto L63
L62:
	;
	v324 = v289
	goto L63
L63:
	;
	v326 = v286 + int32(1)
	v330 = *(*int32)(unsafe.Add(mBase, uint32(v261+v326<<(uint(int32(2))%32))))
	v331 = int32(_a_F__hash_doinsert_5)
	if v330&v331 == v331 {
		goto L64
	} else {
		goto L65
	}
L64:
	;
	v337 = int32(1)
	*(*uint16)(unsafe.Add(mBase, uint32(v30+int32(32)+v324<<(uint(v337)%32)))) = uint16(v326)
	v343 = v324 + v337
	goto L66
L65:
	;
	v343 = v324
	goto L66
L66:
	;
	v344 = int32(2)
	v345 = v286 + v344
	v347 = v294 + v344
	if v347 != v274&int32(-2) {
		v286 = v345
		v289 = v343
		v294 = v347
		goto L59
	} else {
		goto L67
	}
L67:
	;
	goto L60
L68:
	;
	v356 = v345
	v359 = v343
	goto L55
L69:
	;
	v388 = int32(1)
	*(*uint16)(unsafe.Add(mBase, uint32(v30+int32(32)+v359<<(uint(v388)%32)))) = uint16(v356)
	v402 = v359 + v388
	goto L52
L70:
	;
	v424 = v30 + int32(32)
	v425 = F_index_compute_xid_horizon_for_tuples(m, l0, l2, v196, v424, v402)
	mBase = m.M
	v426 = m.ExcPending
	if v426 != 0 {
		goto L9
	} else {
		goto L71
	}
L71:
	;
	F_LockBuffer(m, v79, int32(2))
	mBase = m.M
	v429 = m.ExcPending
	if v429 != 0 {
		goto L9
	} else {
		goto L72
	}
L72:
	;
	v430 = int32(_a_F__hash_doinsert_6)
	v432 = *(*int32)(unsafe.Add(mBase, _c_F__hash_doinsert[4]))
	*(*int32)(unsafe.Add(mBase, _c_F__hash_doinsert[4])) = v432 + int32(1)
	F_PageIndexMultiDelete(m, v247, v424, v402)
	mBase = m.M
	v437 = m.ExcPending
	if v437 != 0 {
		goto L9
	} else {
		goto L73
	}
L73:
	;
	v438 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v247)+16)))
	v439 = v247 + v438
	v440 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v439)+12)))
	v442 = v440 & int32(_a_F__hash_doinsert_7)
	*(*uint16)(unsafe.Add(mBase, uint32(v439)+12)) = uint16(v442)
	v444 = int32(0)
	v445 = base.B2i32(v444 <= v79)
	if v445 == v444 {
		goto L75
	} else {
		goto L76
	}
L74:
	;
	v460 = *(*float64)(unsafe.Add(mBase, uint32(v459)+32))
	*(*float64)(unsafe.Add(mBase, uint32(v459)+32)) = base.F64_sub(v460, base.F64_convert_i32_u(v402))
	F_MarkBufferDirty(m, v196)
	mBase = m.M
	v465 = m.ExcPending
	if v465 != 0 {
		goto L9
	} else {
		goto L78
	}
L75:
	;
	v449 = *(*int32)(unsafe.Add(mBase, _c_F__hash_doinsert[0]))
	v453 = *(*int32)(unsafe.Add(mBase, uint32(v449+v189<<(uint(int32(2))%32))))
	v459 = v453
	goto L74
L76:
	;
	goto L77
L77:
	;
	v455 = *(*int32)(unsafe.Add(mBase, _c_F__hash_doinsert[1]))
	v459 = v455 + v191 + int32(-8192)
	goto L74
L78:
	;
	F_MarkBufferDirty(m, v79)
	mBase = m.M
	v467 = m.ExcPending
	if v467 != 0 {
		goto L9
	} else {
		goto L79
	}
L79:
	;
	v468 = *(*int32)(unsafe.Add(mBase, uint32(l0)+48))
	v469 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v468)+118)))
	if v469 != int32(112) {
		goto L80
	} else {
		goto L81
	}
L80:
	;
	v577 = int32(_a_F__hash_doinsert_6)
	v579 = *(*int32)(unsafe.Add(mBase, _c_F__hash_doinsert[4]))
	*(*int32)(unsafe.Add(mBase, _c_F__hash_doinsert[4])) = v579 - int32(1)
	F_LockBuffer(m, v79, int32(0))
	mBase = m.M
	v585 = m.ExcPending
	if v585 != 0 {
		goto L9
	} else {
		goto L114
	}
L81:
	;
	v473 = *(*int32)(unsafe.Add(mBase, _c_F__hash_doinsert[5]))
	if v473 <= int32(0) {
		goto L83
	} else {
		goto L84
	}
L82:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v30)+24)) = v425
	*(*uint8)(unsafe.Add(mBase, uint32(v30)+30)) = uint8(v501)
	*(*uint16)(unsafe.Add(mBase, uint32(v30)+28)) = uint16(v402)
	F_XLogBeginInsert(m)
	mBase = m.M
	v507 = m.ExcPending
	if v507 != 0 {
		goto L9
	} else {
		goto L100
	}
L83:
	;
	v476 = *(*int32)(unsafe.Add(mBase, uint32(l0)+32))
	if v476 != 0 {
		goto L80
	} else {
		goto L86
	}
L84:
	;
	goto L85
L85:
	;
	if v473 == int32(1) {
		goto L88
	} else {
		goto L89
	}
L86:
	;
	v477 = *(*int32)(unsafe.Add(mBase, uint32(l0)+40))
	if v477 != 0 {
		goto L80
	} else {
		goto L87
	}
L87:
	;
	v501 = int32(0)
	goto L82
L88:
	;
	v501 = int32(0)
	goto L82
L89:
	;
	goto L90
L90:
	;
	v482 = *(*int32)(unsafe.Add(mBase, uint32(l2)+48))
	v483 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v482)+118)))
	if v483 != int32(112) {
		goto L91
	} else {
		goto L92
	}
L91:
	;
	v501 = int32(0)
	goto L82
L92:
	;
	goto L93
L93:
	;
	v488 = *(*int32)(unsafe.Add(mBase, uint32(l2)+56))
	goto L94
L94:
	;
	if base.Ui32(v488) < base.Ui32(int32(_a_F__hash_doinsert_8)) {
		v501 = int32(1)
		goto L82
	} else {
		goto L95
	}
L95:
	;
	v491 = *(*int32)(unsafe.Add(mBase, uint32(l2)+180))
	if v491 == int32(0) {
		goto L96
	} else {
		goto L97
	}
L96:
	;
	v501 = int32(0)
	goto L82
L97:
	;
	goto L98
L98:
	;
	v496 = *(*int32)(unsafe.Add(mBase, uint32(l2)+48))
	v497 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v496)+119)))
	switch v497 - int32(109) {
	case 0, 5:
		goto L99
	default:
		v501 = int32(0)
		goto L82
	}
L99:
	;
	v500 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v491)+104)))
	v501 = v500
	goto L82
L100:
	;
	F_XLogRegisterBuffer(m, int32(0), v196, int32(8))
	mBase = m.M
	v511 = m.ExcPending
	if v511 != 0 {
		goto L9
	} else {
		goto L101
	}
L101:
	;
	F_XLogRegisterData(m, v30+int32(24), int32(8))
	mBase = m.M
	v516 = m.ExcPending
	if v516 != 0 {
		goto L9
	} else {
		goto L102
	}
L102:
	;
	F_XLogRegisterData(m, v30+int32(32), v402<<(uint(int32(1))%32))
	mBase = m.M
	v522 = m.ExcPending
	if v522 != 0 {
		goto L9
	} else {
		goto L103
	}
L103:
	;
	F_XLogRegisterBuffer(m, int32(1), v79, int32(8))
	mBase = m.M
	v526 = m.ExcPending
	if v526 != 0 {
		goto L9
	} else {
		goto L104
	}
L104:
	;
	v529 = F_XLogInsert(m, int32(12), int32(192))
	mBase = m.M
	v530 = m.ExcPending
	if v530 != 0 {
		goto L9
	} else {
		goto L105
	}
L105:
	;
	if v229 == int32(0) {
		goto L107
	} else {
		goto L108
	}
L106:
	;
	v549 = int64(32)
	*(*int64)(unsafe.Add(mBase, uint32(v548))) = base.I64_rotr(v529, v549)
	if v445 == int32(0) {
		goto L111
	} else {
		goto L112
	}
L107:
	;
	v534 = *(*int32)(unsafe.Add(mBase, _c_F__hash_doinsert[0]))
	v540 = *(*int32)(unsafe.Add(mBase, uint32(v534+(v196^int32(-1))<<(uint(int32(2))%32))))
	v548 = v540
	goto L106
L108:
	;
	goto L109
L109:
	;
	v542 = *(*int32)(unsafe.Add(mBase, _c_F__hash_doinsert[1]))
	v548 = v542 + v196<<(uint(int32(13))%32) + int32(-8192)
	goto L106
L110:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v569)+4)) = base.I32_wrap_i64(v529)
	*(*int32)(unsafe.Add(mBase, uint32(v569))) = base.I32_wrap_i64(int64(base.Ui64(v529) >> (uint(v549) % 64)))
	goto L80
L111:
	;
	v559 = *(*int32)(unsafe.Add(mBase, _c_F__hash_doinsert[0]))
	v563 = *(*int32)(unsafe.Add(mBase, uint32(v559+v189<<(uint(int32(2))%32))))
	v569 = v563
	goto L110
L112:
	;
	goto L113
L113:
	;
	v565 = *(*int32)(unsafe.Add(mBase, _c_F__hash_doinsert[1]))
	v569 = v565 + v191 + int32(-8192)
	goto L110
L114:
	;
	goto L45
L115:
	;
	if base.Ui32(v48) <= base.Ui32(v619-int32(4)) {
		v721 = v196
		goto L33
	} else {
		goto L119
	}
L116:
	;
	v619 = v613
	goto L118
L117:
	;
	v619 = v616
	goto L118
L118:
	;
	goto L115
L119:
	;
	goto L41
L120:
	;
	v705 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v704)+16)))
	v707 = int32(4)
	v708 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v704)+14)))
	v709 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v704)+12)))
	v710 = v708 - v709
	if v710 <= v707 {
		goto L140
	} else {
		goto L141
	}
L121:
	;
	if v196 != v111 {
		goto L125
	} else {
		goto L126
	}
L122:
	;
	goto L123
L123:
	;
	F_LockBuffer(m, v196, int32(0))
	mBase = m.M
	v682 = m.ExcPending
	if v682 != 0 {
		goto L9
	} else {
		goto L134
	}
L124:
	;
	v661 = F__hash_getbuf(m, l0, v650, int32(2), int32(1))
	mBase = m.M
	v662 = m.ExcPending
	if v662 != 0 {
		goto L9
	} else {
		goto L130
	}
L125:
	;
	F_UnlockReleaseBuffer(m, v196)
	mBase = m.M
	v655 = m.ExcPending
	if v655 != 0 {
		goto L9
	} else {
		goto L128
	}
L126:
	;
	goto L127
L127:
	;
	F_LockBuffer(m, v111, int32(0))
	mBase = m.M
	v658 = m.ExcPending
	if v658 != 0 {
		goto L9
	} else {
		goto L129
	}
L128:
	;
	goto L124
L129:
	;
	goto L124
L130:
	;
	if v661 < int32(0) {
		goto L131
	} else {
		goto L132
	}
L131:
	;
	v666 = *(*int32)(unsafe.Add(mBase, _c_F__hash_doinsert[0]))
	v672 = *(*int32)(unsafe.Add(mBase, uint32(v666+(v661^int32(-1))<<(uint(int32(2))%32))))
	v703 = v661
	v704 = v672
	goto L120
L132:
	;
	goto L133
L133:
	;
	v674 = *(*int32)(unsafe.Add(mBase, _c_F__hash_doinsert[1]))
	v703 = v661
	v704 = v674 + v661<<(uint(int32(13))%32) + int32(-8192)
	goto L120
L134:
	;
	v684 = F__hash_addovflpage(m, l0, v79, v196, base.B2i32(v196 == v111))
	mBase = m.M
	v685 = m.ExcPending
	if v685 != 0 {
		goto L9
	} else {
		goto L135
	}
L135:
	;
	if v684 < int32(0) {
		goto L136
	} else {
		goto L137
	}
L136:
	;
	v689 = *(*int32)(unsafe.Add(mBase, _c_F__hash_doinsert[0]))
	v695 = *(*int32)(unsafe.Add(mBase, uint32(v689+(v684^int32(-1))<<(uint(int32(2))%32))))
	v703 = v684
	v704 = v695
	goto L120
L137:
	;
	goto L138
L138:
	;
	v697 = *(*int32)(unsafe.Add(mBase, _c_F__hash_doinsert[1]))
	v703 = v684
	v704 = v697 + v684<<(uint(int32(13))%32) + int32(-8192)
	goto L120
L139:
	;
	if base.Ui32(v713-int32(4)) < base.Ui32(v48) {
		v196 = v703
		v201 = v704
		v208 = v705 + v704
		goto L39
	} else {
		goto L143
	}
L140:
	;
	v713 = v707
	goto L142
L141:
	;
	v713 = v710
	goto L142
L142:
	;
	goto L139
L143:
	;
	goto L40
L144:
	;
	v747 = int32(_a_F__hash_doinsert_6)
	v749 = *(*int32)(unsafe.Add(mBase, _c_F__hash_doinsert[4]))
	*(*int32)(unsafe.Add(mBase, _c_F__hash_doinsert[4])) = v749 + int32(1)
	v753 = m.G0
	v755 = v753 - int32(16)
	m.G0 = v755
	F__hash_checkpage(m, l0, v721, int32(3))
	mBase = m.M
	v759 = m.ExcPending
	if v759 != 0 {
		goto L9
	} else {
		goto L145
	}
L145:
	;
	if v721 < int32(0) {
		goto L147
	} else {
		goto L148
	}
L146:
	;
	if l3 != 0 {
		goto L151
	} else {
		goto L152
	}
L147:
	;
	v763 = *(*int32)(unsafe.Add(mBase, _c_F__hash_doinsert[0]))
	v769 = *(*int32)(unsafe.Add(mBase, uint32(v763+(v721^int32(-1))<<(uint(int32(2))%32))))
	v777 = v769
	goto L146
L148:
	;
	goto L149
L149:
	;
	v771 = *(*int32)(unsafe.Add(mBase, _c_F__hash_doinsert[1]))
	v777 = v771 + v721<<(uint(int32(13))%32) + int32(-8192)
	goto L146
L150:
	;
	v871 = v869 & int32(_a_F__hash_doinsert_4)
	v873 = F_PageAddItemExtended(m, v777, l1, v48, v871, int32(0))
	mBase = m.M
	v874 = m.ExcPending
	if v874 != 0 {
		goto L9
	} else {
		goto L180
	}
L151:
	;
	v778 = int32(1)
	v779 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v777)+12)))
	if base.Ui32(v779) < base.Ui32(int32(25)) {
		goto L154
	} else {
		goto L155
	}
L152:
	;
	goto L153
L153:
	;
	v791 = int32(*(*int16)(unsafe.Add(mBase, uint32(l1)+6)))
	if int32(0) <= v791 {
		goto L158
	} else {
		goto L159
	}
L154:
	;
	v788 = v778
	goto L156
L155:
	;
	v788 = int32(base.Ui32(v779+int32(_a_F__hash_doinsert_3))>>(uint(int32(2))%32)) + v778
	goto L156
L156:
	;
	v869 = v788
	goto L150
L157:
	;
	v802 = int32(1)
	v804 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v777)+12)))
	if base.Ui32(v804) < base.Ui32(int32(25)) {
		goto L162
	} else {
		goto L163
	}
L158:
	;
	v794 = int32(8)
	goto L160
L159:
	;
	v794 = int32(16)
	goto L160
L160:
	;
	v796 = *(*int32)(unsafe.Add(mBase, uint32(l1+v794)))
	goto L157
L161:
	;
	v869 = v862 & int32(_a_F__hash_doinsert_4)
	goto L150
L162:
	;
	v813 = v802
	goto L164
L163:
	;
	v813 = int32(base.Ui32(v804+int32(_a_F__hash_doinsert_3))>>(uint(int32(2))%32)) + v802
	goto L164
L164:
	;
	if base.Ui32(int32(2)) <= base.Ui32(v813&int32(_a_F__hash_doinsert_4)) {
		goto L165
	} else {
		goto L166
	}
L165:
	;
	v822 = v813
	v823 = v802
	goto L168
L166:
	;
	v862 = v802
	goto L167
L167:
	;
	goto L161
L168:
	;
	v827 = int32(_a_F__hash_doinsert_4)
	v833 = int32(base.Ui32(v822&v827+v823&v827) >> (uint(int32(1)) % 32))
	v837 = *(*int32)(unsafe.Add(mBase, uint32(v777+int32(20)+v833<<(uint(int32(2))%32))))
	v840 = v777 + v837&int32(_a_F__hash_doinsert_9)
	v843 = int32(*(*int16)(unsafe.Add(mBase, uint32(v840)+6)))
	if int32(0) <= v843 {
		goto L170
	} else {
		goto L171
	}
L169:
	;
	v862 = v855
	goto L167
L170:
	;
	v846 = int32(8)
	goto L172
L171:
	;
	v846 = int32(16)
	goto L172
L172:
	;
	v848 = *(*int32)(unsafe.Add(mBase, uint32(v840+v846)))
	v849 = base.B2i32(base.Ui32(v848) < base.Ui32(v796))
	if base.Ui32(v848) < base.Ui32(v796) {
		goto L173
	} else {
		goto L174
	}
L173:
	;
	v850 = v822
	goto L175
L174:
	;
	v850 = v833
	goto L175
L175:
	;
	if base.Ui32(v848) < base.Ui32(v796) {
		goto L176
	} else {
		goto L177
	}
L176:
	;
	v855 = v833 + int32(1)
	goto L178
L177:
	;
	v855 = v823
	goto L178
L178:
	;
	if base.Ui32(v855&int32(_a_F__hash_doinsert_4)) < base.Ui32(v850&int32(_a_F__hash_doinsert_4)) {
		v822 = v850
		v823 = v855
		goto L168
	} else {
		goto L179
	}
L179:
	;
	goto L169
L180:
	;
	if v873 == int32(0) {
		goto L181
	} else {
		goto L182
	}
L181:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v880 = m.ExcPending
	if v880 != 0 {
		goto L9
	} else {
		goto L184
	}
L182:
	;
	goto L183
L183:
	;
	m.G0 = v755 + int32(16)
	F_MarkBufferDirty(m, v721)
	mBase = m.M
	v897 = m.ExcPending
	if v897 != 0 {
		goto L9
	} else {
		goto L187
	}
L184:
	;
	v881 = *(*int32)(unsafe.Add(mBase, uint32(l0)+48))
	*(*int32)(unsafe.Add(mBase, uint32(v755))) = v881 + int32(4)
	F_errmsg_internal(m, int32(_a_F__hash_doinsert_10), v755)
	mBase = m.M
	v887 = m.ExcPending
	if v887 != 0 {
		goto L9
	} else {
		goto L185
	}
L185:
	;
	F_errfinish(m, int32(_a_F__hash_doinsert_11), int32(316), int32(_a_F__hash_doinsert_12))
	mBase = m.M
	v892 = m.ExcPending
	if v892 != 0 {
		goto L9
	} else {
		goto L186
	}
L186:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L187:
	;
	v898 = *(*float64)(unsafe.Add(mBase, uint32(v98)+32))
	v900 = base.F64_add(v898, float64(1))
	*(*float64)(unsafe.Add(mBase, uint32(v98)+32)) = v900
	v902 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v98)+40)))
	v903 = *(*int32)(unsafe.Add(mBase, uint32(v98)+48))
	F_MarkBufferDirty(m, v79)
	mBase = m.M
	v905 = m.ExcPending
	if v905 != 0 {
		goto L9
	} else {
		goto L188
	}
L188:
	;
	v910 = *(*int32)(unsafe.Add(mBase, uint32(l0)+48))
	v911 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v910)+118)))
	if v911 != int32(112) {
		goto L189
	} else {
		goto L190
	}
L189:
	;
	v996 = int32(_a_F__hash_doinsert_6)
	v998 = *(*int32)(unsafe.Add(mBase, _c_F__hash_doinsert[4]))
	*(*int32)(unsafe.Add(mBase, _c_F__hash_doinsert[4])) = v998 - int32(1)
	F_LockBuffer(m, v79, int32(0))
	mBase = m.M
	v1004 = m.ExcPending
	if v1004 != 0 {
		goto L9
	} else {
		goto L210
	}
L190:
	;
	v915 = *(*int32)(unsafe.Add(mBase, _c_F__hash_doinsert[5]))
	if v915 <= int32(0) {
		goto L191
	} else {
		goto L192
	}
L191:
	;
	v918 = *(*int32)(unsafe.Add(mBase, uint32(l0)+32))
	if v918 != 0 {
		goto L189
	} else {
		goto L194
	}
L192:
	;
	goto L193
L193:
	;
	*(*uint16)(unsafe.Add(mBase, uint32(v30)+32)) = uint16(v871)
	F_XLogBeginInsert(m)
	mBase = m.M
	v922 = m.ExcPending
	if v922 != 0 {
		goto L9
	} else {
		goto L196
	}
L194:
	;
	v919 = *(*int32)(unsafe.Add(mBase, uint32(l0)+40))
	if v919 != 0 {
		goto L189
	} else {
		goto L195
	}
L195:
	;
	goto L193
L196:
	;
	F_XLogRegisterData(m, v30+int32(32), int32(2))
	mBase = m.M
	v927 = m.ExcPending
	if v927 != 0 {
		goto L9
	} else {
		goto L197
	}
L197:
	;
	F_XLogRegisterBuffer(m, int32(1), v79, int32(8))
	mBase = m.M
	v931 = m.ExcPending
	if v931 != 0 {
		goto L9
	} else {
		goto L198
	}
L198:
	;
	F_XLogRegisterBuffer(m, int32(0), v721, int32(8))
	mBase = m.M
	v935 = m.ExcPending
	if v935 != 0 {
		goto L9
	} else {
		goto L199
	}
L199:
	;
	v937 = int32(*(*uint16)(unsafe.Add(mBase, uint32(l1)+6)))
	F_XLogRegisterBufData(m, int32(0), l1, v937&int32(_a_F__hash_doinsert_1))
	mBase = m.M
	v941 = m.ExcPending
	if v941 != 0 {
		goto L9
	} else {
		goto L200
	}
L200:
	;
	v944 = F_XLogInsert(m, int32(12), int32(32))
	mBase = m.M
	v945 = m.ExcPending
	if v945 != 0 {
		goto L9
	} else {
		goto L201
	}
L201:
	;
	if v721 < int32(0) {
		goto L203
	} else {
		goto L204
	}
L202:
	;
	v964 = int64(32)
	*(*int64)(unsafe.Add(mBase, uint32(v963))) = base.I64_rotr(v944, v964)
	if v79 < int32(0) {
		goto L207
	} else {
		goto L208
	}
L203:
	;
	v949 = *(*int32)(unsafe.Add(mBase, _c_F__hash_doinsert[0]))
	v955 = *(*int32)(unsafe.Add(mBase, uint32(v949+(v721^int32(-1))<<(uint(int32(2))%32))))
	v963 = v955
	goto L202
L204:
	;
	goto L205
L205:
	;
	v957 = *(*int32)(unsafe.Add(mBase, _c_F__hash_doinsert[1]))
	v963 = v957 + v721<<(uint(int32(13))%32) + int32(-8192)
	goto L202
L206:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v988)+4)) = base.I32_wrap_i64(v944)
	*(*int32)(unsafe.Add(mBase, uint32(v988))) = base.I32_wrap_i64(int64(base.Ui64(v944) >> (uint(v964) % 64)))
	goto L189
L207:
	;
	v974 = *(*int32)(unsafe.Add(mBase, _c_F__hash_doinsert[0]))
	v980 = *(*int32)(unsafe.Add(mBase, uint32(v974+(v79^int32(-1))<<(uint(int32(2))%32))))
	v988 = v980
	goto L206
L208:
	;
	goto L209
L209:
	;
	v982 = *(*int32)(unsafe.Add(mBase, _c_F__hash_doinsert[1]))
	v988 = v982 + v79<<(uint(int32(13))%32) + int32(-8192)
	goto L206
L210:
	;
	F_UnlockReleaseBuffer(m, v721)
	mBase = m.M
	v1006 = m.ExcPending
	if v1006 != 0 {
		goto L9
	} else {
		goto L211
	}
L211:
	;
	if v721 != v111 {
		goto L212
	} else {
		goto L213
	}
L212:
	;
	F_ReleaseBuffer(m, v111)
	mBase = m.M
	v1009 = m.ExcPending
	if v1009 != 0 {
		goto L9
	} else {
		goto L215
	}
L213:
	;
	goto L214
L214:
	;
	if base.F64_lt(base.F64_mul(base.F64_convert_i32_u(v902), base.F64_convert_i32_u(v903+int32(1))), v900) != 0 {
		goto L216
	} else {
		goto L217
	}
L215:
	;
	goto L214
L216:
	;
	v1011 = m.G0
	v1015 = (v1011 - int32(_a_F__hash_doinsert_13)) & int32(-4096)
	m.G0 = v1015
	v1018 = v79 ^ int32(-1)
	v1020 = v79 << (uint(int32(13)) % 32)
	goto L221
L217:
	;
	goto L218
L218:
	;
	F_ReleaseBuffer(m, v79)
	mBase = m.M
	v1727 = m.ExcPending
	if v1727 != 0 {
		goto L9
	} else {
		goto L393
	}
L219:
	;
	goto L218
L220:
	;
	F_LockBuffer(m, v79, int32(0))
	mBase = m.M
	v1679 = m.ExcPending
	if v1679 != 0 {
		goto L9
	} else {
		goto L392
	}
L221:
	;
	F_LockBuffer(m, v79, int32(2))
	mBase = m.M
	v1050 = m.ExcPending
	if v1050 != 0 {
		goto L9
	} else {
		goto L223
	}
L222:
	;
	v1198 = int32(2)
	v1199 = v1068 + v1198
	v1203 = v1199 - int32(1)
	if base.Ui32(v1198) <= base.Ui32(v1199) {
		goto L274
	} else {
		goto L275
	}
L223:
	;
	F__hash_checkpage(m, l0, v79, int32(8))
	mBase = m.M
	v1053 = m.ExcPending
	if v1053 != 0 {
		goto L9
	} else {
		goto L224
	}
L224:
	;
	if v79 < int32(0) {
		goto L226
	} else {
		goto L227
	}
L225:
	;
	v1068 = *(*int32)(unsafe.Add(mBase, uint32(v1067)+48))
	if base.Ui32(int32(2147483645)) < base.Ui32(v1068) {
		goto L220
	} else {
		goto L229
	}
L226:
	;
	v1057 = *(*int32)(unsafe.Add(mBase, _c_F__hash_doinsert[0]))
	v1061 = *(*int32)(unsafe.Add(mBase, uint32(v1057+v1018<<(uint(int32(2))%32))))
	v1067 = v1061
	goto L225
L227:
	;
	goto L228
L228:
	;
	v1063 = *(*int32)(unsafe.Add(mBase, _c_F__hash_doinsert[1]))
	v1067 = v1063 + v1020 + int32(-8192)
	goto L225
L229:
	;
	v1071 = *(*float64)(unsafe.Add(mBase, uint32(v1067)+32))
	v1072 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v1067)+40)))
	v1075 = v1068 + int32(1)
	if base.F64_le(v1071, base.F64_mul(base.F64_convert_i32_u(v1072), base.F64_convert_i32_u(v1075))) != 0 {
		goto L220
	} else {
		goto L230
	}
L230:
	;
	v1079 = *(*int32)(unsafe.Add(mBase, uint32(v1067)+56))
	v1080 = v1079 & v1075
	if v1080 != 0 {
		goto L232
	} else {
		goto L233
	}
L231:
	;
	if v1141 == int32(0) {
		goto L220
	} else {
		goto L255
	}
L232:
	;
	v1081 = int32(1)
	v1082 = v1080 + v1081
	v1086 = v1082 - v1081
	if base.Ui32(int32(2)) <= base.Ui32(v1082) {
		goto L236
	} else {
		goto L237
	}
L233:
	;
	v1113 = int32(1)
	goto L234
L234:
	;
	if v1113 != int32(-1) {
		goto L242
	} else {
		goto L243
	}
L235:
	;
	v1109 = *(*int32)(unsafe.Add(mBase, uint32(v1067+v1105<<(uint(int32(2))%32))+72))
	v1113 = v1109 + v1082
	goto L234
L236:
	;
	v1092 = int32(32) - base.I32_clz(v1086)
	goto L238
L237:
	;
	v1092 = int32(0)
	goto L238
L238:
	;
	if base.Ui32(int32(10)) <= base.Ui32(v1092) {
		goto L239
	} else {
		goto L240
	}
L239:
	;
	v1095 = int32(3)
	v1105 = int32(base.Ui32(v1086)>>(uint(v1092-v1095)%32))&v1095 | v1092<<(uint(int32(2))%32) - int32(30)
	goto L241
L240:
	;
	v1105 = v1092
	goto L241
L241:
	;
	goto L235
L242:
	;
	v1116 = F_ReadBuffer(m, l0, v1113)
	mBase = m.M
	v1117 = m.ExcPending
	if v1117 != 0 {
		goto L9
	} else {
		goto L245
	}
L243:
	;
	goto L244
L244:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v1131 = m.ExcPending
	if v1131 != 0 {
		goto L9
	} else {
		goto L252
	}
L245:
	;
	v1118 = F_ConditionalLockBufferForCleanup(m, v1116)
	mBase = m.M
	v1119 = m.ExcPending
	if v1119 != 0 {
		goto L9
	} else {
		goto L246
	}
L246:
	;
	if v1118 == int32(0) {
		goto L247
	} else {
		goto L248
	}
L247:
	;
	F_ReleaseBuffer(m, v1116)
	mBase = m.M
	v1123 = m.ExcPending
	if v1123 != 0 {
		goto L9
	} else {
		goto L250
	}
L248:
	;
	goto L249
L249:
	;
	F__hash_checkpage(m, l0, v1116, int32(2))
	mBase = m.M
	v1127 = m.ExcPending
	if v1127 != 0 {
		goto L9
	} else {
		goto L251
	}
L250:
	;
	v1141 = int32(0)
	goto L231
L251:
	;
	v1141 = v1116
	goto L231
L252:
	;
	F_errmsg_internal(m, int32(_a_F__hash_doinsert_14), int32(0))
	mBase = m.M
	v1135 = m.ExcPending
	if v1135 != 0 {
		goto L9
	} else {
		goto L253
	}
L253:
	;
	F_errfinish(m, int32(_a_F__hash_doinsert_15), int32(101), int32(_a_F__hash_doinsert_16))
	mBase = m.M
	v1140 = m.ExcPending
	if v1140 != 0 {
		goto L9
	} else {
		goto L254
	}
L254:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L255:
	;
	if v1141 < int32(0) {
		goto L257
	} else {
		goto L258
	}
L256:
	;
	v1162 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v1161)+16)))
	v1164 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v1162+v1161)+12)))
	if v1164&int32(32) != 0 {
		goto L260
	} else {
		goto L261
	}
L257:
	;
	v1147 = *(*int32)(unsafe.Add(mBase, _c_F__hash_doinsert[0]))
	v1153 = *(*int32)(unsafe.Add(mBase, uint32(v1147+(v1141^int32(-1))<<(uint(int32(2))%32))))
	v1161 = v1153
	goto L256
L258:
	;
	goto L259
L259:
	;
	v1155 = *(*int32)(unsafe.Add(mBase, _c_F__hash_doinsert[1]))
	v1161 = v1155 + v1141<<(uint(int32(13))%32) + int32(-8192)
	goto L256
L260:
	;
	v1167 = *(*int32)(unsafe.Add(mBase, uint32(v1067)+56))
	v1168 = *(*int32)(unsafe.Add(mBase, uint32(v1067)+52))
	v1169 = *(*int32)(unsafe.Add(mBase, uint32(v1067)+48))
	F_LockBuffer(m, v79, int32(0))
	mBase = m.M
	v1172 = m.ExcPending
	if v1172 != 0 {
		goto L9
	} else {
		goto L263
	}
L261:
	;
	goto L262
L262:
	;
	if v1164&int32(64) != 0 {
		goto L267
	} else {
		goto L268
	}
L263:
	;
	F_LockBuffer(m, v1141, int32(0))
	mBase = m.M
	v1175 = m.ExcPending
	if v1175 != 0 {
		goto L9
	} else {
		goto L264
	}
L264:
	;
	F__hash_finish_split(m, l0, v79, v1141, v1080, v1169, v1168, v1167)
	mBase = m.M
	v1177 = m.ExcPending
	if v1177 != 0 {
		goto L9
	} else {
		goto L265
	}
L265:
	;
	F_ReleaseBuffer(m, v1141)
	mBase = m.M
	v1179 = m.ExcPending
	if v1179 != 0 {
		goto L9
	} else {
		goto L266
	}
L266:
	;
	goto L221
L267:
	;
	v1182 = *(*int32)(unsafe.Add(mBase, uint32(v1067)+56))
	v1183 = *(*int32)(unsafe.Add(mBase, uint32(v1067)+52))
	v1184 = *(*int32)(unsafe.Add(mBase, uint32(v1067)+48))
	F_LockBuffer(m, v79, int32(0))
	mBase = m.M
	v1187 = m.ExcPending
	if v1187 != 0 {
		goto L9
	} else {
		goto L270
	}
L268:
	;
	goto L269
L269:
	;
	goto L222
L270:
	;
	v1188 = int32(0)
	F_hashbucketcleanup(m, l0, v1080, v1141, v1113, v1188, v1184, v1183, v1182, v1188, v1188, int32(1), v1188, v1188)
	mBase = m.M
	v1195 = m.ExcPending
	if v1195 != 0 {
		goto L9
	} else {
		goto L271
	}
L271:
	;
	F_ReleaseBuffer(m, v1141)
	mBase = m.M
	v1197 = m.ExcPending
	if v1197 != 0 {
		goto L9
	} else {
		goto L272
	}
L272:
	;
	goto L221
L273:
	;
	v1223 = int32(2)
	v1226 = *(*int32)(unsafe.Add(mBase, uint32(v1067+v1222<<(uint(v1223)%32))+72))
	v1227 = v1226 + v1075
	v1228 = int32(1)
	v1229 = v1227 + v1228
	v1233 = v1199 - v1228
	if base.Ui32(v1223) <= base.Ui32(v1199) {
		goto L282
	} else {
		goto L283
	}
L274:
	;
	v1209 = int32(32) - base.I32_clz(v1203)
	goto L276
L275:
	;
	v1209 = int32(0)
	goto L276
L276:
	;
	if base.Ui32(int32(10)) <= base.Ui32(v1209) {
		goto L277
	} else {
		goto L278
	}
L277:
	;
	v1212 = int32(3)
	v1222 = int32(base.Ui32(v1203)>>(uint(v1209-v1212)%32))&v1212 | v1209<<(uint(int32(2))%32) - int32(30)
	goto L279
L278:
	;
	v1222 = v1209
	goto L279
L279:
	;
	goto L273
L280:
	;
	F_UnlockReleaseBuffer(m, v1141)
	mBase = m.M
	v1667 = m.ExcPending
	if v1667 != 0 {
		goto L9
	} else {
		goto L391
	}
L281:
	;
	v1253 = *(*int32)(unsafe.Add(mBase, uint32(v1067)+60))
	if base.Ui32(v1253) < base.Ui32(v1252) {
		goto L288
	} else {
		goto L289
	}
L282:
	;
	v1239 = int32(32) - base.I32_clz(v1233)
	goto L284
L283:
	;
	v1239 = int32(0)
	goto L284
L284:
	;
	if base.Ui32(int32(10)) <= base.Ui32(v1239) {
		goto L285
	} else {
		goto L286
	}
L285:
	;
	v1242 = int32(3)
	v1252 = int32(base.Ui32(v1233)>>(uint(v1239-v1242)%32))&v1242 | v1239<<(uint(int32(2))%32) - int32(30)
	goto L287
L286:
	;
	v1252 = v1239
	goto L287
L287:
	;
	goto L281
L288:
	;
	if base.Ui32(v1252) <= base.Ui32(int32(9)) {
		goto L292
	} else {
		goto L293
	}
L289:
	;
	goto L290
L290:
	;
	v1406 = F__hash_getnewbuf(m, l0, v1229, int32(0))
	mBase = m.M
	v1407 = m.ExcPending
	if v1407 != 0 {
		goto L9
	} else {
		goto L328
	}
L291:
	;
	v1277 = v1276 + v1226
	if base.B2i32(base.Ui32(v1277) < base.Ui32(v1229))|base.B2i32(v1276-v1075^v1227 == int32(-1)) != 0 {
		goto L280
	} else {
		goto L295
	}
L292:
	;
	v1276 = int32(1) << (uint(v1252) % 32)
	goto L291
L293:
	;
	goto L294
L294:
	;
	v1262 = v1252 - int32(10)
	v1263 = int32(2)
	v1265 = int32(512) << (uint(int32(base.Ui32(v1262)>>(uint(v1263)%32))) % 32)
	v1276 = v1265>>(uint(v1263)%32)*(v1262&int32(3)+int32(1)) + v1265
	goto L291
L295:
	;
	v1285 = v1015 + int32(_a_F__hash_doinsert_17)
	v1288 = int32(0)
	if v1288|(v1285&int32(3)|int32(1)) == v1288 {
		goto L298
	} else {
		goto L299
	}
L296:
	;
	v1336 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v1015)+uint32(_c_F__hash_doinsert[6]))))
	v1337 = v1285 + v1336
	*(*int64)(unsafe.Add(mBase, uint32(v1337)+8)) = int64(-36028792723996673)
	*(*int64)(unsafe.Add(mBase, uint32(v1337))) = int64(-1)
	v1342 = *(*int32)(unsafe.Add(mBase, uint32(l0)+48))
	v1343 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1342)+118)))
	if v1343 != int32(112) {
		goto L307
	} else {
		goto L308
	}
L297:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v1015)+uint32(_c_F__hash_doinsert[7]))) = int32(_a_F__hash_doinsert_18)
	v1327 = int32(_a_F__hash_doinsert_19)
	*(*uint16)(unsafe.Add(mBase, uint32(v1015)+uint32(_c_F__hash_doinsert[8]))) = uint16(v1327)
	v1333 = int32(_a_F__hash_doinsert_20)
	*(*uint16)(unsafe.Add(mBase, uint32(v1015)+uint32(_c_F__hash_doinsert[6]))) = uint16(v1333)
	*(*uint16)(unsafe.Add(mBase, uint32(v1015)+uint32(_c_F__hash_doinsert[9]))) = uint16(v1333)
	goto L296
L298:
	;
	goto L301
L299:
	;
	goto L300
L300:
	;
	goto L306
L301:
	;
	v1304 = v1015 + int32(_a_F__hash_doinsert_13)
	v1306 = v1015 + int32(_a_F__hash_doinsert_21)
	if base.Ui32(v1306) < base.Ui32(v1304) {
		goto L302
	} else {
		goto L303
	}
L302:
	;
	v1308 = v1304
	goto L304
L303:
	;
	v1308 = v1306
	goto L304
L304:
	;
	v1313 = (v1285^int32(-1)+v1308)&int32(-4) + int32(4)
	if v1313 == int32(0) {
		goto L297
	} else {
		goto L305
	}
L305:
	;
	base.MemoryFill(m, v1285, int32(0), v1313)
	goto L297
L306:
	;
	base.MemoryFill(m, v1285, int32(0), int32(_a_F__hash_doinsert_22))
	goto L297
L307:
	;
	v1360 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v1015)+uint32(_c_F__hash_doinsert[9]))))
	if v1360 == int32(0) {
		goto L316
	} else {
		goto L317
	}
L308:
	;
	v1347 = *(*int32)(unsafe.Add(mBase, _c_F__hash_doinsert[5]))
	if v1347 <= int32(0) {
		goto L309
	} else {
		goto L310
	}
L309:
	;
	v1350 = *(*int32)(unsafe.Add(mBase, uint32(l0)+32))
	if v1350 != 0 {
		goto L307
	} else {
		goto L312
	}
L310:
	;
	goto L311
L311:
	;
	F_log_newpage(m, l0, int32(0), v1277, v1015+int32(_a_F__hash_doinsert_17), int32(1))
	mBase = m.M
	v1357 = m.ExcPending
	if v1357 != 0 {
		goto L9
	} else {
		goto L314
	}
L312:
	;
	v1351 = *(*int32)(unsafe.Add(mBase, uint32(l0)+40))
	if v1351 != 0 {
		goto L307
	} else {
		goto L313
	}
L313:
	;
	goto L311
L314:
	;
	goto L307
L315:
	;
	v1368 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	if v1368 != 0 {
		goto L319
	} else {
		goto L320
	}
L316:
	;
	goto L315
L317:
	;
	v1363 = F_DataChecksumsEnabled(m)
	mBase = m.M
	if v1363 == int32(0) {
		goto L316
	} else {
		goto L318
	}
L318:
	;
	v1366 = F_pg_checksum_page(m, v1015+int32(_a_F__hash_doinsert_17), v1277)
	mBase = m.M
	*(*uint16)(unsafe.Add(mBase, uint32(v1015)+uint32(_c_F__hash_doinsert[10]))) = uint16(v1366)
	goto L316
L319:
	;
	v1394 = v1368
	goto L321
L320:
	;
	v1369 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v1370 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	*(*int32)(unsafe.Add(mBase, uint32(v1015)+4088)) = v1370
	v1372 = *(*int64)(unsafe.Add(mBase, uint32(l0)))
	*(*int64)(unsafe.Add(mBase, uint32(v1015)+4080)) = v1372
	v1376 = F_smgropen(m, v1015+int32(4080), v1369)
	mBase = m.M
	v1377 = m.ExcPending
	if v1377 != 0 {
		goto L9
	} else {
		goto L322
	}
L321:
	;
	v1395 = int32(0)
	F_smgrextend(m, v1394, v1395, v1277, v1015+int32(_a_F__hash_doinsert_17), v1395)
	mBase = m.M
	v1400 = m.ExcPending
	if v1400 != 0 {
		goto L9
	} else {
		goto L327
	}
L322:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+12)) = v1376
	v1380 = *(*int32)(unsafe.Add(mBase, uint32(v1376)+72))
	if v1380 != 0 {
		goto L324
	} else {
		goto L325
	}
L323:
	;
	v1392 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	v1394 = v1392
	goto L321
L324:
	;
	v1388 = v1380
	goto L326
L325:
	;
	v1381 = *(*int32)(unsafe.Add(mBase, uint32(v1376)+76))
	v1382 = *(*int32)(unsafe.Add(mBase, uint32(v1376)+80))
	*(*int32)(unsafe.Add(mBase, uint32(v1381)+4)) = v1382
	v1384 = *(*int32)(unsafe.Add(mBase, uint32(v1376)+76))
	*(*int32)(unsafe.Add(mBase, uint32(v1382))) = v1384
	v1386 = *(*int32)(unsafe.Add(mBase, uint32(v1376)+72))
	v1388 = v1386
	goto L326
L326:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v1376)+72)) = v1388 + int32(1)
	goto L323
L327:
	;
	goto L290
L328:
	;
	v1408 = F_IsBufferCleanupOK(m, v1406)
	mBase = m.M
	v1409 = m.ExcPending
	if v1409 != 0 {
		goto L9
	} else {
		goto L329
	}
L329:
	;
	if v1408 == int32(0) {
		goto L330
	} else {
		goto L331
	}
L330:
	;
	F_UnlockReleaseBuffer(m, v1141)
	mBase = m.M
	v1413 = m.ExcPending
	if v1413 != 0 {
		goto L9
	} else {
		goto L333
	}
L331:
	;
	goto L332
L332:
	;
	v1417 = v1067 + int32(56)
	v1418 = int32(_a_F__hash_doinsert_6)
	v1420 = *(*int32)(unsafe.Add(mBase, _c_F__hash_doinsert[4]))
	*(*int32)(unsafe.Add(mBase, _c_F__hash_doinsert[4])) = v1420 + int32(1)
	*(*int32)(unsafe.Add(mBase, uint32(v1067)+48)) = v1075
	v1426 = v1067 + int32(52)
	v1427 = *(*int32)(unsafe.Add(mBase, uint32(v1067)+52))
	v1428 = base.B2i32(base.Ui32(v1075) <= base.Ui32(v1427))
	if v1428 == int32(0) {
		goto L335
	} else {
		goto L336
	}
L333:
	;
	F_UnlockReleaseBuffer(m, v1406)
	mBase = m.M
	v1415 = m.ExcPending
	if v1415 != 0 {
		goto L9
	} else {
		goto L334
	}
L334:
	;
	goto L220
L335:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v1417))) = v1427
	*(*int32)(unsafe.Add(mBase, uint32(v1426))) = v1075 | v1427
	goto L337
L336:
	;
	goto L337
L337:
	;
	v1435 = v1067 + int32(60)
	v1436 = *(*int32)(unsafe.Add(mBase, uint32(v1435)))
	v1437 = base.B2i32(base.Ui32(v1252) <= base.Ui32(v1436))
	if v1437 == int32(0) {
		goto L338
	} else {
		goto L339
	}
L338:
	;
	v1441 = v1067 + int32(76)
	v1442 = int32(2)
	v1448 = *(*int32)(unsafe.Add(mBase, uint32(v1441+v1436<<(uint(v1442)%32))))
	*(*int32)(unsafe.Add(mBase, uint32(v1441+v1252<<(uint(v1442)%32)))) = v1448
	*(*int32)(unsafe.Add(mBase, uint32(v1067)+60)) = v1252
	goto L340
L339:
	;
	goto L340
L340:
	;
	F_MarkBufferDirty(m, v79)
	mBase = m.M
	v1453 = m.ExcPending
	if v1453 != 0 {
		goto L9
	} else {
		goto L341
	}
L341:
	;
	v1454 = *(*int32)(unsafe.Add(mBase, uint32(v1067)+48))
	v1455 = *(*int32)(unsafe.Add(mBase, uint32(v1067)+56))
	v1456 = *(*int32)(unsafe.Add(mBase, uint32(v1067)+52))
	if v1141 < int32(0) {
		goto L343
	} else {
		goto L344
	}
L342:
	;
	v1475 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v1474)+16)))
	v1476 = v1475 + v1474
	*(*int32)(unsafe.Add(mBase, uint32(v1476))) = v1454
	v1478 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v1476)+12)))
	v1480 = v1478 | int32(32)
	*(*uint16)(unsafe.Add(mBase, uint32(v1476)+12)) = uint16(v1480)
	F_MarkBufferDirty(m, v1141)
	mBase = m.M
	v1483 = m.ExcPending
	if v1483 != 0 {
		goto L9
	} else {
		goto L346
	}
L343:
	;
	v1460 = *(*int32)(unsafe.Add(mBase, _c_F__hash_doinsert[0]))
	v1466 = *(*int32)(unsafe.Add(mBase, uint32(v1460+(v1141^int32(-1))<<(uint(int32(2))%32))))
	v1474 = v1466
	goto L342
L344:
	;
	goto L345
L345:
	;
	v1468 = *(*int32)(unsafe.Add(mBase, _c_F__hash_doinsert[1]))
	v1474 = v1468 + v1141<<(uint(int32(13))%32) + int32(-8192)
	goto L342
L346:
	;
	if v1406 < int32(0) {
		goto L348
	} else {
		goto L349
	}
L347:
	;
	v1502 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v1501)+16)))
	v1503 = v1502 + v1501
	*(*int32)(unsafe.Add(mBase, uint32(v1503)+12)) = int32(-8388590)
	*(*int32)(unsafe.Add(mBase, uint32(v1503)+8)) = v1075
	*(*int32)(unsafe.Add(mBase, uint32(v1503)+4)) = int32(-1)
	*(*int32)(unsafe.Add(mBase, uint32(v1503))) = v1454
	F_MarkBufferDirty(m, v1406)
	mBase = m.M
	v1511 = m.ExcPending
	if v1511 != 0 {
		goto L9
	} else {
		goto L351
	}
L348:
	;
	v1487 = *(*int32)(unsafe.Add(mBase, _c_F__hash_doinsert[0]))
	v1493 = *(*int32)(unsafe.Add(mBase, uint32(v1487+(v1406^int32(-1))<<(uint(int32(2))%32))))
	v1501 = v1493
	goto L347
L349:
	;
	goto L350
L350:
	;
	v1495 = *(*int32)(unsafe.Add(mBase, _c_F__hash_doinsert[1]))
	v1501 = v1495 + v1406<<(uint(int32(13))%32) + int32(-8192)
	goto L347
L351:
	;
	v1512 = *(*int32)(unsafe.Add(mBase, uint32(l0)+48))
	v1513 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1512)+118)))
	if v1513 != int32(112) {
		goto L352
	} else {
		goto L353
	}
L352:
	;
	v1649 = int32(_a_F__hash_doinsert_6)
	v1651 = *(*int32)(unsafe.Add(mBase, _c_F__hash_doinsert[4]))
	*(*int32)(unsafe.Add(mBase, _c_F__hash_doinsert[4])) = v1651 - int32(1)
	F_LockBuffer(m, v79, int32(0))
	mBase = m.M
	v1657 = m.ExcPending
	if v1657 != 0 {
		goto L9
	} else {
		goto L387
	}
L353:
	;
	v1517 = *(*int32)(unsafe.Add(mBase, _c_F__hash_doinsert[5]))
	if v1517 <= int32(0) {
		goto L354
	} else {
		goto L355
	}
L354:
	;
	v1520 = *(*int32)(unsafe.Add(mBase, uint32(l0)+32))
	if v1520 != 0 {
		goto L352
	} else {
		goto L357
	}
L355:
	;
	goto L356
L356:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v1015)+uint32(_c_F__hash_doinsert[11]))) = v1454
	v1523 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v1476)+12)))
	*(*uint16)(unsafe.Add(mBase, uint32(v1015)+uint32(_c_F__hash_doinsert[12]))) = uint16(v1523)
	v1525 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v1503)+12)))
	v1526 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v1015)+uint32(_c_F__hash_doinsert[10]))) = uint8(v1526)
	*(*uint16)(unsafe.Add(mBase, uint32(v1015)+uint32(_c_F__hash_doinsert[13]))) = uint16(v1525)
	F_XLogBeginInsert(m)
	mBase = m.M
	v1530 = m.ExcPending
	if v1530 != 0 {
		goto L9
	} else {
		goto L359
	}
L357:
	;
	v1521 = *(*int32)(unsafe.Add(mBase, uint32(l0)+40))
	if v1521 != 0 {
		goto L352
	} else {
		goto L358
	}
L358:
	;
	goto L356
L359:
	;
	F_XLogRegisterBuffer(m, int32(0), v1141, int32(8))
	mBase = m.M
	v1534 = m.ExcPending
	if v1534 != 0 {
		goto L9
	} else {
		goto L360
	}
L360:
	;
	F_XLogRegisterBuffer(m, int32(1), v1406, int32(6))
	mBase = m.M
	v1538 = m.ExcPending
	if v1538 != 0 {
		goto L9
	} else {
		goto L361
	}
L361:
	;
	v1539 = int32(2)
	F_XLogRegisterBuffer(m, v1539, v79, int32(8))
	mBase = m.M
	v1543 = m.ExcPending
	if v1543 != 0 {
		goto L9
	} else {
		goto L362
	}
L362:
	;
	if v1428 == int32(0) {
		goto L363
	} else {
		goto L364
	}
L363:
	;
	v1546 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(v1015)+uint32(_c_F__hash_doinsert[10]))) = uint8(v1546)
	F_XLogRegisterBufData(m, int32(2), v1417, int32(4))
	mBase = m.M
	v1551 = m.ExcPending
	if v1551 != 0 {
		goto L9
	} else {
		goto L366
	}
L364:
	;
	v1557 = v1539
	goto L365
L365:
	;
	if v1437 == int32(0) {
		goto L368
	} else {
		goto L369
	}
L366:
	;
	F_XLogRegisterBufData(m, int32(2), v1426, int32(4))
	mBase = m.M
	v1555 = m.ExcPending
	if v1555 != 0 {
		goto L9
	} else {
		goto L367
	}
L367:
	;
	v1557 = int32(3)
	goto L365
L368:
	;
	*(*uint8)(unsafe.Add(mBase, uint32(v1015)+uint32(_c_F__hash_doinsert[10]))) = uint8(v1557)
	F_XLogRegisterBufData(m, int32(2), v1435, int32(4))
	mBase = m.M
	v1564 = m.ExcPending
	if v1564 != 0 {
		goto L9
	} else {
		goto L371
	}
L369:
	;
	goto L370
L370:
	;
	F_XLogRegisterData(m, v1015+int32(_a_F__hash_doinsert_17), int32(9))
	mBase = m.M
	v1579 = m.ExcPending
	if v1579 != 0 {
		goto L9
	} else {
		goto L373
	}
L371:
	;
	v1565 = int32(2)
	v1566 = *(*int32)(unsafe.Add(mBase, uint32(v1067)+60))
	F_XLogRegisterBufData(m, v1565, v1067+v1566<<(uint(v1565)%32)+int32(76), int32(4))
	mBase = m.M
	v1574 = m.ExcPending
	if v1574 != 0 {
		goto L9
	} else {
		goto L372
	}
L372:
	;
	goto L370
L373:
	;
	v1582 = F_XLogInsert(m, int32(12), int32(64))
	mBase = m.M
	v1583 = m.ExcPending
	if v1583 != 0 {
		goto L9
	} else {
		goto L374
	}
L374:
	;
	if v1141 < int32(0) {
		goto L376
	} else {
		goto L377
	}
L375:
	;
	v1602 = int64(32)
	*(*int64)(unsafe.Add(mBase, uint32(v1601))) = base.I64_rotr(v1582, v1602)
	v1607 = base.I32_wrap_i64(int64(base.Ui64(v1582) >> (uint(v1602) % 64)))
	v1608 = base.I32_wrap_i64(v1582)
	if v1406 < int32(0) {
		goto L380
	} else {
		goto L381
	}
L376:
	;
	v1587 = *(*int32)(unsafe.Add(mBase, _c_F__hash_doinsert[0]))
	v1593 = *(*int32)(unsafe.Add(mBase, uint32(v1587+(v1141^int32(-1))<<(uint(int32(2))%32))))
	v1601 = v1593
	goto L375
L377:
	;
	goto L378
L378:
	;
	v1595 = *(*int32)(unsafe.Add(mBase, _c_F__hash_doinsert[1]))
	v1601 = v1595 + v1141<<(uint(int32(13))%32) + int32(-8192)
	goto L375
L379:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v1626)+4)) = v1608
	*(*int32)(unsafe.Add(mBase, uint32(v1626))) = v1607
	if v79 < int32(0) {
		goto L384
	} else {
		goto L385
	}
L380:
	;
	v1612 = *(*int32)(unsafe.Add(mBase, _c_F__hash_doinsert[0]))
	v1618 = *(*int32)(unsafe.Add(mBase, uint32(v1612+(v1406^int32(-1))<<(uint(int32(2))%32))))
	v1626 = v1618
	goto L379
L381:
	;
	goto L382
L382:
	;
	v1620 = *(*int32)(unsafe.Add(mBase, _c_F__hash_doinsert[1]))
	v1626 = v1620 + v1406<<(uint(int32(13))%32) + int32(-8192)
	goto L379
L383:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v1642)+4)) = v1608
	*(*int32)(unsafe.Add(mBase, uint32(v1642))) = v1607
	goto L352
L384:
	;
	v1632 = *(*int32)(unsafe.Add(mBase, _c_F__hash_doinsert[0]))
	v1636 = *(*int32)(unsafe.Add(mBase, uint32(v1632+v1018<<(uint(int32(2))%32))))
	v1642 = v1636
	goto L383
L385:
	;
	goto L386
L386:
	;
	v1638 = *(*int32)(unsafe.Add(mBase, _c_F__hash_doinsert[1]))
	v1642 = v1638 + v1020 + int32(-8192)
	goto L383
L387:
	;
	F__hash_splitbucket(m, l0, v79, v1080, v1075, v1141, v1406, int32(0), v1454, v1456, v1455)
	mBase = m.M
	v1660 = m.ExcPending
	if v1660 != 0 {
		goto L9
	} else {
		goto L388
	}
L388:
	;
	F_ReleaseBuffer(m, v1141)
	mBase = m.M
	v1662 = m.ExcPending
	if v1662 != 0 {
		goto L9
	} else {
		goto L389
	}
L389:
	;
	F_ReleaseBuffer(m, v1406)
	mBase = m.M
	v1664 = m.ExcPending
	if v1664 != 0 {
		goto L9
	} else {
		goto L390
	}
L390:
	;
	m.G0 = v1011
	goto L219
L391:
	;
	goto L220
L392:
	;
	m.G0 = v1011
	goto L219
L393:
	;
	m.G0 = v30 + int32(_a_F__hash_doinsert_0)
	return
L394:
	;
	F_errcode(m, int32(261))
	mBase = m.M
	v1737 = m.ExcPending
	if v1737 != 0 {
		goto L9
	} else {
		goto L395
	}
L395:
	;
	v1738 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v98)+19)))
	*(*int32)(unsafe.Add(mBase, uint32(v30))) = v48
	*(*int32)(unsafe.Add(mBase, uint32(v30)+4)) = (v1738<<(uint(int32(8))%32) - int32(44)) & int32(-48)
	F_errmsg(m, int32(_a_F__hash_doinsert_23), v30)
	mBase = m.M
	v1749 = m.ExcPending
	if v1749 != 0 {
		goto L9
	} else {
		goto L396
	}
L396:
	;
	F_errhint(m, int32(_a_F__hash_doinsert_24), int32(0))
	mBase = m.M
	v1753 = m.ExcPending
	if v1753 != 0 {
		goto L9
	} else {
		goto L397
	}
L397:
	;
	F_errfinish(m, int32(_a_F__hash_doinsert_11), int32(86), int32(_a_F__hash_doinsert_25))
	mBase = m.M
	v1758 = m.ExcPending
	if v1758 != 0 {
		goto L9
	} else {
		goto L398
	}
L398:
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
	var v223 int32
	_ = v223
	var v225 int32
	_ = v225
	var v228 int32
	_ = v228
	var v229 int32
	_ = v229
	var v232 int32
	_ = v232
	var v233 int32
	_ = v233
	var v238 int32
	_ = v238
	var v242 int32
	_ = v242
	var v248 int32
	_ = v248
	var v250 int32
	_ = v250
	var v256 int32
	_ = v256
	var v257 int32
	_ = v257
	var v259 int32
	_ = v259
	var v261 int32
	_ = v261
	var v263 int32
	_ = v263
	var v266 int32
	_ = v266
	var v273 int32
	_ = v273
	var v277 int32
	_ = v277
	var v282 int32
	_ = v282
	v19 = m.G0
	v21 = v19 + int32(-64)
	m.G0 = v21
	*(*int64)(unsafe.Add(mBase, uint32(v21)+32)) = int64(25769803782)
	v26 = *(*int32)(unsafe.Add(mBase, _c_F__hash_finish_split[0]))
	*(*int32)(unsafe.Add(mBase, uint32(v21)+56)) = v26
	v33 = F_hash_create(m, int32(_a_F__hash_finish_split_0), int32(256), v19+int32(-48), int32(1064))
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
	v273 = m.ExcPending
	if v273 != 0 {
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
	v228 = F_ConditionalLockBufferForCleanup(m, l2)
	mBase = m.M
	v229 = m.ExcPending
	if v229 != 0 {
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
	F_LockBuffer(m, v122, int32(1))
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
	F_LockBuffer(m, v122, int32(0))
	mBase = m.M
	v223 = m.ExcPending
	if v223 != 0 {
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
	v225 = m.ExcPending
	if v225 != 0 {
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
	v266 = m.ExcPending
	if v266 != 0 {
		goto L1
	} else {
		goto L63
	}
L50:
	;
	if v228 == int32(0) {
		goto L49
	} else {
		goto L51
	}
L51:
	;
	v232 = F_ConditionalLockBufferForCleanup(m, v131)
	mBase = m.M
	v233 = m.ExcPending
	if v233 != 0 {
		goto L1
	} else {
		goto L52
	}
L52:
	;
	if v232 == int32(0) {
		goto L53
	} else {
		goto L54
	}
L53:
	;
	F_LockBuffer(m, l2, int32(0))
	mBase = m.M
	v238 = m.ExcPending
	if v238 != 0 {
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
	v257 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v256)+16)))
	v259 = *(*int32)(unsafe.Add(mBase, uint32(v257+v256)+8))
	F__hash_splitbucket(m, l0, l1, l3, v259, l2, v131, v33, l4, l5, l6)
	mBase = m.M
	v261 = m.ExcPending
	if v261 != 0 {
		goto L1
	} else {
		goto L61
	}
L58:
	;
	v242 = *(*int32)(unsafe.Add(mBase, _c_F__hash_finish_split[1]))
	v248 = *(*int32)(unsafe.Add(mBase, uint32(v242+(v131^int32(-1))<<(uint(int32(2))%32))))
	v256 = v248
	goto L57
L59:
	;
	goto L60
L60:
	;
	v250 = *(*int32)(unsafe.Add(mBase, _c_F__hash_finish_split[2]))
	v256 = v250 + v131<<(uint(int32(13))%32) + int32(-8192)
	goto L57
L61:
	;
	F_ReleaseBuffer(m, v131)
	mBase = m.M
	v263 = m.ExcPending
	if v263 != 0 {
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
	v277 = m.ExcPending
	if v277 != 0 {
		goto L1
	} else {
		goto L65
	}
L65:
	;
	F_errfinish(m, int32(_a_F__hash_finish_split_6), int32(75), int32(_a_F__hash_finish_split_7))
	mBase = m.M
	v282 = m.ExcPending
	if v282 != 0 {
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
	var v12 int32
	_ = v12
	var v14 int32
	_ = v14
	var v16 int32
	_ = v16
	var v17 int32
	_ = v17
	var v18 int32
	_ = v18
	var v21 int32
	_ = v21
	var v27 int32
	_ = v27
	var v28 int32
	_ = v28
	var v29 int32
	_ = v29
	var v30 int64
	_ = v30
	var v35 int32
	_ = v35
	var v36 int64
	_ = v36
	var v40 int32
	_ = v40
	var v41 int32
	_ = v41
	var v44 int32
	_ = v44
	var v45 int32
	_ = v45
	var v48 int32
	_ = v48
	var v49 int32
	_ = v49
	var v50 int32
	_ = v50
	var v51 int32
	_ = v51
	var v53 int32
	_ = v53
	var v55 int32
	_ = v55
	var v56 int32
	_ = v56
	var v57 int32
	_ = v57
	var v58 int32
	_ = v58
	var v59 int32
	_ = v59
	var v60 int32
	_ = v60
	var v61 int32
	_ = v61
	var v63 int32
	_ = v63
	var v65 int32
	_ = v65
	var v66 int32
	_ = v66
	var v68 int32
	_ = v68
	var v69 int32
	_ = v69
	var v75 int32
	_ = v75
	var v76 int32
	_ = v76
	var v86 int32
	_ = v86
	var v91 int32
	_ = v91
	var v92 int32
	_ = v92
	var v93 int32
	_ = v93
	var v94 int32
	_ = v94
	var v95 int32
	_ = v95
	var v101 int32
	_ = v101
	var v105 int32
	_ = v105
	var v106 int32
	_ = v106
	var v111 int32
	_ = v111
	var v117 int32
	_ = v117
	var v119 int32
	_ = v119
	var v125 int32
	_ = v125
	var v126 int32
	_ = v126
	var v127 int32
	_ = v127
	var v129 int32
	_ = v129
	var v133 int32
	_ = v133
	var v139 int32
	_ = v139
	var v141 int32
	_ = v141
	var v147 int32
	_ = v147
	var v148 int32
	_ = v148
	var v149 int32
	_ = v149
	var v151 int32
	_ = v151
	var v153 int32
	_ = v153
	var v158 int32
	_ = v158
	var v165 int32
	_ = v165
	var v169 int32
	_ = v169
	var v170 int32
	_ = v170
	var v174 int32
	_ = v174
	var v180 int32
	_ = v180
	var v182 int32
	_ = v182
	var v188 int32
	_ = v188
	var v190 int32
	_ = v190
	var v191 int32
	_ = v191
	var v204 int32
	_ = v204
	var v208 int32
	_ = v208
	var v212 int32
	_ = v212
	var v214 int32
	_ = v214
	var v220 int32
	_ = v220
	var v223 int32
	_ = v223
	var v224 int32
	_ = v224
	var v228 int32
	_ = v228
	var v231 int32
	_ = v231
	var v235 int32
	_ = v235
	var v241 int32
	_ = v241
	var v243 int32
	_ = v243
	var v249 int32
	_ = v249
	var v250 int32
	_ = v250
	var v251 int32
	_ = v251
	var v253 int32
	_ = v253
	var v256 int32
	_ = v256
	var v258 int32
	_ = v258
	var v260 int32
	_ = v260
	var v263 int32
	_ = v263
	var v273 int32
	_ = v273
	var v282 int32
	_ = v282
	var v285 int32
	_ = v285
	var v288 int32
	_ = v288
	var v296 int32
	_ = v296
	var v297 int32
	_ = v297
	var v298 int32
	_ = v298
	var v310 int32
	_ = v310
	var v315 int32
	_ = v315
	var v316 int32
	_ = v316
	var v319 int32
	_ = v319
	var v322 int32
	_ = v322
	var v323 int32
	_ = v323
	var v325 int32
	_ = v325
	var v330 int32
	_ = v330
	var v346 int32
	_ = v346
	var v349 int32
	_ = v349
	var v353 int32
	_ = v353
	var v358 int32
	_ = v358
	v12 = m.G0
	v14 = v12 - int32(16)
	m.G0 = v14
	v16 = *(*int32)(unsafe.Add(mBase, uint32(l0)+36))
	v17 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v18 = *(*int32)(unsafe.Add(mBase, uint32(v17)+272))
	if v18 == int32(0) {
		goto L2
	} else {
		goto L3
	}
L1:
	;
	v35 = *(*int32)(unsafe.Add(mBase, uint32(l0)+40))
	if v35 != 0 {
		goto L8
	} else {
		goto L9
	}
L2:
	;
	v21 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v17)+268)))
	if v21 != int32(1) {
		goto L1
	} else {
		goto L5
	}
L3:
	;
	v29 = v18
	goto L4
L4:
	;
	v30 = *(*int64)(unsafe.Add(mBase, uint32(v29)+16))
	*(*int64)(unsafe.Add(mBase, uint32(v29)+16)) = v30 + int64(1)
	goto L1
L5:
	;
	F_pgstat_assoc_relation(m, v17)
	mBase = m.M
	v27 = m.ExcPending
	if v27 != 0 {
		goto L6
	} else {
		goto L7
	}
L6:
	;
	return int32(0)
L7:
	;
	v28 = *(*int32)(unsafe.Add(mBase, uint32(v17)+272))
	v29 = v28
	goto L4
L8:
	;
	v36 = *(*int64)(unsafe.Add(mBase, uint32(v35)))
	*(*int64)(unsafe.Add(mBase, uint32(v35))) = v36 + int64(1)
	goto L10
L9:
	;
	goto L10
L10:
	;
	v40 = int32(0)
	v41 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	if v40 < v41 {
		goto L11
	} else {
		goto L12
	}
L11:
	;
	v44 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v45 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v44))))
	if v45&int32(1) != 0 {
		v330 = v40
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
	v346 = m.ExcPending
	if v346 != 0 {
		goto L6
	} else {
		goto L82
	}
L14:
	;
	m.G0 = v14 + int32(16)
	return v330
L15:
	;
	v48 = *(*int32)(unsafe.Add(mBase, uint32(v44)+44))
	v49 = *(*int32)(unsafe.Add(mBase, uint32(v44)+8))
	if v49 != 0 {
		goto L18
	} else {
		goto L19
	}
L16:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v16))) = v101
	v105 = F__hash_getbucketbuf_from_hashkey(m, v17, v101, int32(1), int32(0))
	mBase = m.M
	v106 = m.ExcPending
	if v106 != 0 {
		goto L6
	} else {
		goto L32
	}
L17:
	;
	v61 = m.G0
	v63 = v61 - int32(16)
	m.G0 = v63
	v65 = *(*int32)(unsafe.Add(mBase, uint32(v17)+208))
	v66 = *(*int32)(unsafe.Add(mBase, uint32(v65)))
	v68 = F_get_opfamily_proc(m, v66, v49, v49, int32(1))
	mBase = m.M
	v69 = m.ExcPending
	if v69 != 0 {
		goto L6
	} else {
		goto L24
	}
L18:
	;
	v50 = *(*int32)(unsafe.Add(mBase, uint32(v17)+212))
	v51 = *(*int32)(unsafe.Add(mBase, uint32(v50)))
	if v49 != v51 {
		goto L17
	} else {
		goto L21
	}
L19:
	;
	goto L20
L20:
	;
	v53 = int32(1)
	v55 = F_index_getprocinfo(m, v17, v53, v53)
	mBase = m.M
	v56 = m.ExcPending
	if v56 != 0 {
		goto L6
	} else {
		goto L22
	}
L21:
	;
	goto L20
L22:
	;
	v57 = *(*int32)(unsafe.Add(mBase, uint32(v17)+248))
	v58 = *(*int32)(unsafe.Add(mBase, uint32(v57)))
	v59 = F_FunctionCall1Coll(m, v55, v58, v48)
	mBase = m.M
	v60 = m.ExcPending
	if v60 != 0 {
		goto L6
	} else {
		goto L23
	}
L23:
	;
	v101 = v59
	goto L16
L24:
	;
	if v68 == int32(0) {
		goto L25
	} else {
		goto L26
	}
L25:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v75 = m.ExcPending
	if v75 != 0 {
		goto L6
	} else {
		goto L28
	}
L26:
	;
	goto L27
L27:
	;
	v92 = *(*int32)(unsafe.Add(mBase, uint32(v17)+248))
	v93 = *(*int32)(unsafe.Add(mBase, uint32(v92)))
	v94 = F_OidFunctionCall1Coll(m, v68, v93, v48)
	mBase = m.M
	v95 = m.ExcPending
	if v95 != 0 {
		goto L6
	} else {
		goto L31
	}
L28:
	;
	v76 = *(*int32)(unsafe.Add(mBase, uint32(v17)+48))
	*(*int32)(unsafe.Add(mBase, uint32(v63)+8)) = v49
	*(*int32)(unsafe.Add(mBase, uint32(v63)+4)) = v49
	*(*int32)(unsafe.Add(mBase, uint32(v63))) = int32(1)
	*(*int32)(unsafe.Add(mBase, uint32(v63)+12)) = v76 + int32(4)
	F_errmsg_internal(m, int32(_a_F__hash_first_0), v63)
	mBase = m.M
	v86 = m.ExcPending
	if v86 != 0 {
		goto L6
	} else {
		goto L29
	}
L29:
	;
	F_errfinish(m, int32(_a_F__hash_first_1), int32(115), int32(_a_F__hash_first_2))
	mBase = m.M
	v91 = m.ExcPending
	if v91 != 0 {
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
	m.G0 = v63 + int32(16)
	v101 = v94
	goto L16
L32:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v14)+12)) = v105
	if v105 < int32(0) {
		goto L34
	} else {
		goto L35
	}
L33:
	;
	v127 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	F_PredicateLockPage(m, v17, v126, v127)
	mBase = m.M
	v129 = m.ExcPending
	if v129 != 0 {
		goto L6
	} else {
		goto L37
	}
L34:
	;
	v111 = *(*int32)(unsafe.Add(mBase, _c_F__hash_first[0]))
	v117 = *(*int32)(unsafe.Add(mBase, uint32(v111+(v105^int32(-1))<<(uint(int32(6))%32))+16))
	v126 = v117
	goto L33
L35:
	;
	goto L36
L36:
	;
	v119 = *(*int32)(unsafe.Add(mBase, _c_F__hash_first[1]))
	v125 = *(*int32)(unsafe.Add(mBase, uint32(v119+v105<<(uint(int32(6))%32)+int32(-64))+16))
	v126 = v125
	goto L33
L37:
	;
	if v105 < int32(0) {
		goto L39
	} else {
		goto L40
	}
L38:
	;
	v148 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v147)+16)))
	v149 = v148 + v147
	*(*int32)(unsafe.Add(mBase, uint32(v14)+4)) = v149
	v151 = *(*int32)(unsafe.Add(mBase, uint32(v149)+8))
	*(*int32)(unsafe.Add(mBase, uint32(v16)+4)) = v105
	v153 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v149)+12)))
	if v153&int32(16) == int32(0) {
		v263 = v149
		goto L42
	} else {
		goto L43
	}
L39:
	;
	v133 = *(*int32)(unsafe.Add(mBase, _c_F__hash_first[2]))
	v139 = *(*int32)(unsafe.Add(mBase, uint32(v133+(v105^int32(-1))<<(uint(int32(2))%32))))
	v147 = v139
	goto L38
L40:
	;
	goto L41
L41:
	;
	v141 = *(*int32)(unsafe.Add(mBase, _c_F__hash_first[3]))
	v147 = v141 + v105<<(uint(int32(13))%32) + int32(-8192)
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
	v158 = int32(-1)
	v165 = v151 & (v158<<(uint(int32(31)-base.I32_clz(v151))%32) ^ v158)
	v169 = F__hash_getbuf(m, v17, int32(0), int32(1), int32(8))
	mBase = m.M
	v170 = m.ExcPending
	if v170 != 0 {
		goto L6
	} else {
		goto L45
	}
L44:
	;
	if v165 != 0 {
		goto L49
	} else {
		goto L50
	}
L45:
	;
	if v169 < int32(0) {
		goto L46
	} else {
		goto L47
	}
L46:
	;
	v174 = *(*int32)(unsafe.Add(mBase, _c_F__hash_first[2]))
	v180 = *(*int32)(unsafe.Add(mBase, uint32(v174+(v169^int32(-1))<<(uint(int32(2))%32))))
	v188 = v180
	goto L44
L47:
	;
	goto L48
L48:
	;
	v182 = *(*int32)(unsafe.Add(mBase, _c_F__hash_first[3]))
	v188 = v182 + v169<<(uint(int32(13))%32) + int32(-8192)
	goto L44
L49:
	;
	v190 = base.I32_clz(v165)
	v191 = int32(32) - v190
	if base.Ui32(int32(512)) <= base.Ui32(v165) {
		goto L52
	} else {
		goto L53
	}
L50:
	;
	v212 = int32(0)
	goto L51
L51:
	;
	F_UnlockReleaseBuffer(m, v169)
	mBase = m.M
	v214 = m.ExcPending
	if v214 != 0 {
		goto L6
	} else {
		goto L55
	}
L52:
	;
	v204 = int32(base.Ui32(v165)>>(uint(int32(29)-v190)%32))&int32(3) | v191<<(uint(int32(2))%32) - int32(30)
	goto L54
L53:
	;
	v204 = v191
	goto L54
L54:
	;
	v208 = *(*int32)(unsafe.Add(mBase, uint32(v188+v204<<(uint(int32(2))%32))+72))
	v212 = v208
	goto L51
L55:
	;
	F_LockBuffer(m, v105, int32(0))
	mBase = m.M
	v220 = m.ExcPending
	if v220 != 0 {
		goto L6
	} else {
		goto L56
	}
L56:
	;
	v223 = F__hash_getbuf(m, v17, v165+v212+int32(1), int32(1), int32(2))
	mBase = m.M
	v224 = m.ExcPending
	if v224 != 0 {
		goto L6
	} else {
		goto L57
	}
L57:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v16)+8)) = v223
	F_LockBuffer(m, v223, int32(0))
	mBase = m.M
	v228 = m.ExcPending
	if v228 != 0 {
		goto L6
	} else {
		goto L58
	}
L58:
	;
	F_LockBuffer(m, v105, int32(1))
	mBase = m.M
	v231 = m.ExcPending
	if v231 != 0 {
		goto L6
	} else {
		goto L59
	}
L59:
	;
	if v105 < int32(0) {
		goto L61
	} else {
		goto L62
	}
L60:
	;
	v250 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v249)+16)))
	v251 = v250 + v249
	*(*int32)(unsafe.Add(mBase, uint32(v14)+4)) = v251
	v253 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v251)+12)))
	if v253&int32(16) != 0 {
		goto L64
	} else {
		goto L65
	}
L61:
	;
	v235 = *(*int32)(unsafe.Add(mBase, _c_F__hash_first[2]))
	v241 = *(*int32)(unsafe.Add(mBase, uint32(v235+(v105^int32(-1))<<(uint(int32(2))%32))))
	v249 = v241
	goto L60
L62:
	;
	goto L63
L63:
	;
	v243 = *(*int32)(unsafe.Add(mBase, _c_F__hash_first[3]))
	v249 = v243 + v105<<(uint(int32(13))%32) + int32(-8192)
	goto L60
L64:
	;
	v256 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(v16)+12)) = uint8(v256)
	v263 = v251
	goto L42
L65:
	;
	goto L66
L66:
	;
	v258 = *(*int32)(unsafe.Add(mBase, uint32(v16)+8))
	F_ReleaseBuffer(m, v258)
	mBase = m.M
	v260 = m.ExcPending
	if v260 != 0 {
		goto L6
	} else {
		goto L67
	}
L67:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v16)+8)) = int32(0)
	v263 = v251
	goto L42
L68:
	;
	v273 = v263
	goto L71
L69:
	;
	v310 = v105
	goto L70
L70:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v16)+24)) = v310
	v315 = F__hash_readpage(m, l0, v14+int32(12), l1)
	mBase = m.M
	v316 = m.ExcPending
	if v316 != 0 {
		goto L6
	} else {
		goto L80
	}
L71:
	;
	v282 = *(*int32)(unsafe.Add(mBase, uint32(v273)+4))
	if v282 == int32(-1) {
		goto L74
	} else {
		goto L75
	}
L72:
	;
	v298 = *(*int32)(unsafe.Add(mBase, uint32(v14)+12))
	v310 = v298
	goto L70
L73:
	;
	goto L72
L74:
	;
	v285 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v16)+12)))
	if v285 != int32(1) {
		goto L73
	} else {
		goto L77
	}
L75:
	;
	goto L76
L76:
	;
	F__hash_readnext(m, l0, v14+int32(12), v14+int32(8), v14+int32(4))
	mBase = m.M
	v296 = m.ExcPending
	if v296 != 0 {
		goto L6
	} else {
		goto L79
	}
L77:
	;
	v288 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v16)+13)))
	if v288 != 0 {
		goto L73
	} else {
		goto L78
	}
L78:
	;
	goto L76
L79:
	;
	v297 = *(*int32)(unsafe.Add(mBase, uint32(v14)+4))
	v273 = v297
	goto L71
L80:
	;
	if v315 == int32(0) {
		v330 = int32(0)
		goto L14
	} else {
		goto L81
	}
L81:
	;
	v319 = *(*int32)(unsafe.Add(mBase, uint32(v16)+48))
	v322 = v16 + v319<<(uint(int32(3))%32)
	v323 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v322)+56)))
	*(*uint16)(unsafe.Add(mBase, uint32(l0)+64)) = uint16(v323)
	v325 = *(*int32)(unsafe.Add(mBase, uint32(v322)+52))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+60)) = v325
	v330 = int32(1)
	goto L14
L82:
	;
	F_errcode(m, int32(1088))
	mBase = m.M
	v349 = m.ExcPending
	if v349 != 0 {
		goto L6
	} else {
		goto L83
	}
L83:
	;
	F_errmsg(m, int32(_a_F__hash_first_3), int32(0))
	mBase = m.M
	v353 = m.ExcPending
	if v353 != 0 {
		goto L6
	} else {
		goto L84
	}
L84:
	;
	F_errfinish(m, int32(_a_F__hash_first_4), int32(313), int32(_a_F__hash_first_5))
	mBase = m.M
	v358 = m.ExcPending
	if v358 != 0 {
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
			F_LockBuffer(m, v9, int32(2))
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
	F_errfinish(m, int32(_a_F__hash_pgaddmultitup_6), int32(358), int32(_a_F__hash_pgaddmultitup_7))
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
	var v10 int32
	_ = v10
	var v12 int32
	_ = v12
	var v18 int32
	_ = v18
	var v21 int32
	_ = v21
	var v29 int32
	_ = v29
	var v30 int32
	_ = v30
	var v33 int32
	_ = v33
	var v36 int32
	_ = v36
	var v37 float32
	_ = v37
	var v43 int32
	_ = v43
	var v45 int32
	_ = v45
	var v47 int32
	_ = v47
	var v48 int32
	_ = v48
	var v49 float64
	_ = v49
	var v53 int32
	_ = v53
	var v54 int32
	_ = v54
	var v55 int32
	_ = v55
	var v56 int32
	_ = v56
	var v57 float32
	_ = v57
	var v59 float32
	_ = v59
	var v61 int32
	_ = v61
	var v65 int32
	_ = v65
	var v68 int32
	_ = v68
	var v72 int32
	_ = v72
	var v75 int32
	_ = v75
	var v79 int32
	_ = v79
	var v86 float64
	_ = v86
	var v87 float64
	_ = v87
	var v91 int32
	_ = v91
	var v92 float64
	_ = v92
	var v95 float64
	_ = v95
	var v96 int32
	_ = v96
	var v99 int32
	_ = v99
	var v102 float64
	_ = v102
	var v105 int32
	_ = v105
	var v112 float64
	_ = v112
	var v115 float64
	_ = v115
	var v116 int32
	_ = v116
	var v121 float64
	_ = v121
	var v122 int32
	_ = v122
	var v125 float64
	_ = v125
	var v126 float64
	_ = v126
	var v129 float64
	_ = v129
	var v132 int32
	_ = v132
	var v133 int32
	_ = v133
	var v134 int32
	_ = v134
	var v136 float32
	_ = v136
	var v140 float64
	_ = v140
	var v142 float64
	_ = v142
	var v143 int32
	_ = v143
	var v146 float64
	_ = v146
	var v151 float64
	_ = v151
	var v153 float64
	_ = v153
	var v154 float64
	_ = v154
	var v163 float64
	_ = v163
	var v167 float64
	_ = v167
	var v168 float64
	_ = v168
	var v172 float64
	_ = v172
	var v173 float64
	_ = v173
	var v178 float64
	_ = v178
	var v185 float64
	_ = v185
	var v186 float64
	_ = v186
	var v194 float64
	_ = v194
	var v201 int32
	_ = v201
	var v202 int32
	_ = v202
	var v204 int32
	_ = v204
	v10 = m.G0
	v12 = v10 - int32(80)
	m.G0 = v12
	F_examine_variable(m, l0, l1, int32(0), v12+int32(48))
	mBase = m.M
	v18 = m.ExcPending
	if v18 != 0 {
		return
	} else {
		*(*int64)(unsafe.Add(mBase, uint32(l3))) = int64(0)
		v21 = *(*int32)(unsafe.Add(mBase, uint32(v12)+56))
		if v21 == int32(0) {
			v45 = v12 + int32(48)
			v47 = v12 + int32(47)
			v48 = int32(0)
			v49 = float64(0)
			*(*uint8)(unsafe.Add(mBase, uint32(v47))) = uint8(v48)
			v53 = *(*int32)(unsafe.Add(mBase, uint32(v45)+8))
			if v53 != 0 {
				v54 = *(*int32)(unsafe.Add(mBase, uint32(v53)+16))
				v55 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v54)+22)))
				v56 = v54 + v55
				v57 = *(*float32)(unsafe.Add(mBase, uint32(v56)+8))
				v59 = *(*float32)(unsafe.Add(mBase, uint32(v56)+16))
				v86 = base.F64_promote_f32(v59)
				v87 = base.F64_promote_f32(v57)
			} else {
				v61 = *(*int32)(unsafe.Add(mBase, uint32(v45)+16))
				if v61 == int32(16) {
					v86 = float64(2)
					v87 = v49
				} else {
					v65 = *(*int32)(unsafe.Add(mBase, uint32(v45)+4))
					if v65 == int32(0) {
						v72 = *(*int32)(unsafe.Add(mBase, uint32(v45)))
						if v72 == int32(0) {
							v86 = float64(0)
							v87 = v49
						} else {
							v75 = *(*int32)(unsafe.Add(mBase, uint32(v72)))
							if v75 != int32(6) {
								v86 = float64(0)
								v87 = v49
							} else {
								v79 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v72)+8)))
								switch v79 - int32(_a_F_estimate_hash_bucket_stats_0) {
								case 0:
									v86 = float64(1)
									v87 = v49
								default:
									v86 = float64(0)
									v87 = v49
								case 5:
									v86 = float64(-1)
									v87 = v49
								}
							}
						}
					} else {
						v68 = *(*int32)(unsafe.Add(mBase, uint32(v65)+76))
						if v68 != int32(5) {
							v72 = *(*int32)(unsafe.Add(mBase, uint32(v45)))
							if v72 == int32(0) {
								v86 = float64(0)
								v87 = v49
							} else {
								v75 = *(*int32)(unsafe.Add(mBase, uint32(v72)))
								if v75 != int32(6) {
									v86 = float64(0)
									v87 = v49
								} else {
									v79 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v72)+8)))
									switch v79 - int32(_a_F_estimate_hash_bucket_stats_0) {
									case 0:
										v86 = float64(1)
										v87 = v49
									default:
										v86 = float64(0)
										v87 = v49
									case 5:
										v86 = float64(-1)
										v87 = v49
									}
								}
							}
						} else {
							v86 = float64(-1)
							v87 = v49
						}
					}
				}
			}
			v91 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v45)+28)))
			if v91 != 0 {
				v92 = base.F64_neg(base.F64_sub(float64(1), v87))
			} else {
				v92 = v86
			}
			if base.F64_gt(v92, float64(0)) != 0 {
				v95 = F_clamp_row_est(m, v92)
				mBase = m.M
				v121 = v95
			} else {
				v96 = *(*int32)(unsafe.Add(mBase, uint32(v45)+4))
				if v96 == int32(0) {
					v99 = int32(1)
					*(*uint8)(unsafe.Add(mBase, uint32(v47))) = uint8(v99)
					v121 = float64(200)
				} else {
					v102 = *(*float64)(unsafe.Add(mBase, uint32(v96)+120))
					if base.F64_le(v102, float64(0)) != 0 {
						v105 = int32(1)
						*(*uint8)(unsafe.Add(mBase, uint32(v47))) = uint8(v105)
						v121 = float64(200)
					} else {
						if base.F64_lt(v92, float64(0)) != 0 {
							v112 = F_clamp_row_est(m, base.F64_mul(v102, base.F64_neg(v92)))
							mBase = m.M
							v121 = v112
						} else {
							if base.F64_lt(v102, float64(200)) != 0 {
								v115 = F_clamp_row_est(m, v102)
								mBase = m.M
								v121 = v115
							} else {
								v116 = int32(1)
								*(*uint8)(unsafe.Add(mBase, uint32(v47))) = uint8(v116)
								v121 = float64(200)
							}
						}
					}
				}
			}
			v122 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v12)+47)))
			if v122 == int32(1) {
				v125 = float64(0.1)
				v126 = *(*float64)(unsafe.Add(mBase, uint32(l3)))
				if base.F64_lt(v126, v125) != 0 {
					v129 = v125
				} else {
					v129 = v126
				}
				*(*float64)(unsafe.Add(mBase, uint32(l4))) = v129
			} else {
				v132 = *(*int32)(unsafe.Add(mBase, uint32(v12)+56))
				if v132 != 0 {
					v133 = *(*int32)(unsafe.Add(mBase, uint32(v132)+16))
					v134 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v133)+22)))
					v136 = *(*float32)(unsafe.Add(mBase, uint32(v133+v134)+8))
					v140 = base.F64_promote_f32(v136)
				} else {
					v140 = float64(0)
				}
				v142 = base.F64_div(base.F64_sub(float64(1), v140), v121)
				v143 = *(*int32)(unsafe.Add(mBase, uint32(v12)+52))
				if v143 == int32(0) {
					v168 = v121
				} else {
					v146 = *(*float64)(unsafe.Add(mBase, uint32(v143)+120))
					if base.F64_gt(v146, float64(0)) == int32(0) {
						v168 = v121
					} else {
						v151 = *(*float64)(unsafe.Add(mBase, uint32(v143)+16))
						v153 = base.F64_mul(v121, base.F64_div(v151, v146))
						v154 = float64(1e+100)
						if base.B2i32(base.Ui64(int64(9218868437227405312)) < base.Ui64(base.I64_reinterpret_f64(v153)&int64(9223372036854775807)))|base.F64_gt(v153, v154) != 0 {
							v167 = v154
						} else {
							v163 = float64(1)
							if base.F64_le(v153, v163) != 0 {
								v167 = v163
							} else {
								v167 = base.F64_nearest(v153)
							}
						}
						v168 = v167
					}
				}
				if base.F64_lt(l2, v168) != 0 {
					v172 = l2
				} else {
					v172 = v168
				}
				v173 = base.F64_div(float64(1), v172)
				if base.F64_gt(v142, float64(0)) == int32(0) {
					v185 = v173
				} else {
					v178 = *(*float64)(unsafe.Add(mBase, uint32(l3)))
					if base.F64_gt(v178, v142) == int32(0) {
						v185 = v173
					} else {
						v185 = base.F64_mul(v173, base.F64_div(v178, v142))
					}
				}
				v186 = float64(1e-06)
				if base.F64_lt(v185, v186) != 0 {
					v194 = v186
				} else {
					if base.F64_gt(v185, float64(1)) == int32(0) {
						v194 = v185
					} else {
						v194 = float64(1)
					}
				}
				*(*float64)(unsafe.Add(mBase, uint32(l4))) = v194
			}
			v201 = *(*int32)(unsafe.Add(mBase, uint32(v12)+56))
			if v201 != 0 {
				v202 = *(*int32)(unsafe.Add(mBase, uint32(v12)+60))
				m.T0[v202].(func(*base.Module, int32))(m, v201)
				mBase = m.M
				v204 = m.ExcPending
				if v204 != 0 {
					return
				} else {
					m.G0 = v12 + int32(80)
					return
				}
			} else {
				m.G0 = v12 + int32(80)
				return
			}
		} else {
			v29 = F_get_attstatsslot(m, v12+int32(8), v21, int32(1), int32(0), int32(2))
			mBase = m.M
			v30 = m.ExcPending
			if v30 != 0 {
				return
			} else {
				if v29 == int32(0) {
					v45 = v12 + int32(48)
					v47 = v12 + int32(47)
					v48 = int32(0)
					v49 = float64(0)
					*(*uint8)(unsafe.Add(mBase, uint32(v47))) = uint8(v48)
					v53 = *(*int32)(unsafe.Add(mBase, uint32(v45)+8))
					if v53 != 0 {
						v54 = *(*int32)(unsafe.Add(mBase, uint32(v53)+16))
						v55 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v54)+22)))
						v56 = v54 + v55
						v57 = *(*float32)(unsafe.Add(mBase, uint32(v56)+8))
						v59 = *(*float32)(unsafe.Add(mBase, uint32(v56)+16))
						v86 = base.F64_promote_f32(v59)
						v87 = base.F64_promote_f32(v57)
					} else {
						v61 = *(*int32)(unsafe.Add(mBase, uint32(v45)+16))
						if v61 == int32(16) {
							v86 = float64(2)
							v87 = v49
						} else {
							v65 = *(*int32)(unsafe.Add(mBase, uint32(v45)+4))
							if v65 == int32(0) {
								v72 = *(*int32)(unsafe.Add(mBase, uint32(v45)))
								if v72 == int32(0) {
									v86 = float64(0)
									v87 = v49
								} else {
									v75 = *(*int32)(unsafe.Add(mBase, uint32(v72)))
									if v75 != int32(6) {
										v86 = float64(0)
										v87 = v49
									} else {
										v79 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v72)+8)))
										switch v79 - int32(_a_F_estimate_hash_bucket_stats_0) {
										case 0:
											v86 = float64(1)
											v87 = v49
										default:
											v86 = float64(0)
											v87 = v49
										case 5:
											v86 = float64(-1)
											v87 = v49
										}
									}
								}
							} else {
								v68 = *(*int32)(unsafe.Add(mBase, uint32(v65)+76))
								if v68 != int32(5) {
									v72 = *(*int32)(unsafe.Add(mBase, uint32(v45)))
									if v72 == int32(0) {
										v86 = float64(0)
										v87 = v49
									} else {
										v75 = *(*int32)(unsafe.Add(mBase, uint32(v72)))
										if v75 != int32(6) {
											v86 = float64(0)
											v87 = v49
										} else {
											v79 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v72)+8)))
											switch v79 - int32(_a_F_estimate_hash_bucket_stats_0) {
											case 0:
												v86 = float64(1)
												v87 = v49
											default:
												v86 = float64(0)
												v87 = v49
											case 5:
												v86 = float64(-1)
												v87 = v49
											}
										}
									}
								} else {
									v86 = float64(-1)
									v87 = v49
								}
							}
						}
					}
					v91 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v45)+28)))
					if v91 != 0 {
						v92 = base.F64_neg(base.F64_sub(float64(1), v87))
					} else {
						v92 = v86
					}
					if base.F64_gt(v92, float64(0)) != 0 {
						v95 = F_clamp_row_est(m, v92)
						mBase = m.M
						v121 = v95
					} else {
						v96 = *(*int32)(unsafe.Add(mBase, uint32(v45)+4))
						if v96 == int32(0) {
							v99 = int32(1)
							*(*uint8)(unsafe.Add(mBase, uint32(v47))) = uint8(v99)
							v121 = float64(200)
						} else {
							v102 = *(*float64)(unsafe.Add(mBase, uint32(v96)+120))
							if base.F64_le(v102, float64(0)) != 0 {
								v105 = int32(1)
								*(*uint8)(unsafe.Add(mBase, uint32(v47))) = uint8(v105)
								v121 = float64(200)
							} else {
								if base.F64_lt(v92, float64(0)) != 0 {
									v112 = F_clamp_row_est(m, base.F64_mul(v102, base.F64_neg(v92)))
									mBase = m.M
									v121 = v112
								} else {
									if base.F64_lt(v102, float64(200)) != 0 {
										v115 = F_clamp_row_est(m, v102)
										mBase = m.M
										v121 = v115
									} else {
										v116 = int32(1)
										*(*uint8)(unsafe.Add(mBase, uint32(v47))) = uint8(v116)
										v121 = float64(200)
									}
								}
							}
						}
					}
					v122 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v12)+47)))
					if v122 == int32(1) {
						v125 = float64(0.1)
						v126 = *(*float64)(unsafe.Add(mBase, uint32(l3)))
						if base.F64_lt(v126, v125) != 0 {
							v129 = v125
						} else {
							v129 = v126
						}
						*(*float64)(unsafe.Add(mBase, uint32(l4))) = v129
					} else {
						v132 = *(*int32)(unsafe.Add(mBase, uint32(v12)+56))
						if v132 != 0 {
							v133 = *(*int32)(unsafe.Add(mBase, uint32(v132)+16))
							v134 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v133)+22)))
							v136 = *(*float32)(unsafe.Add(mBase, uint32(v133+v134)+8))
							v140 = base.F64_promote_f32(v136)
						} else {
							v140 = float64(0)
						}
						v142 = base.F64_div(base.F64_sub(float64(1), v140), v121)
						v143 = *(*int32)(unsafe.Add(mBase, uint32(v12)+52))
						if v143 == int32(0) {
							v168 = v121
						} else {
							v146 = *(*float64)(unsafe.Add(mBase, uint32(v143)+120))
							if base.F64_gt(v146, float64(0)) == int32(0) {
								v168 = v121
							} else {
								v151 = *(*float64)(unsafe.Add(mBase, uint32(v143)+16))
								v153 = base.F64_mul(v121, base.F64_div(v151, v146))
								v154 = float64(1e+100)
								if base.B2i32(base.Ui64(int64(9218868437227405312)) < base.Ui64(base.I64_reinterpret_f64(v153)&int64(9223372036854775807)))|base.F64_gt(v153, v154) != 0 {
									v167 = v154
								} else {
									v163 = float64(1)
									if base.F64_le(v153, v163) != 0 {
										v167 = v163
									} else {
										v167 = base.F64_nearest(v153)
									}
								}
								v168 = v167
							}
						}
						if base.F64_lt(l2, v168) != 0 {
							v172 = l2
						} else {
							v172 = v168
						}
						v173 = base.F64_div(float64(1), v172)
						if base.F64_gt(v142, float64(0)) == int32(0) {
							v185 = v173
						} else {
							v178 = *(*float64)(unsafe.Add(mBase, uint32(l3)))
							if base.F64_gt(v178, v142) == int32(0) {
								v185 = v173
							} else {
								v185 = base.F64_mul(v173, base.F64_div(v178, v142))
							}
						}
						v186 = float64(1e-06)
						if base.F64_lt(v185, v186) != 0 {
							v194 = v186
						} else {
							if base.F64_gt(v185, float64(1)) == int32(0) {
								v194 = v185
							} else {
								v194 = float64(1)
							}
						}
						*(*float64)(unsafe.Add(mBase, uint32(l4))) = v194
					}
					v201 = *(*int32)(unsafe.Add(mBase, uint32(v12)+56))
					if v201 != 0 {
						v202 = *(*int32)(unsafe.Add(mBase, uint32(v12)+60))
						m.T0[v202].(func(*base.Module, int32))(m, v201)
						mBase = m.M
						v204 = m.ExcPending
						if v204 != 0 {
							return
						} else {
							m.G0 = v12 + int32(80)
							return
						}
					} else {
						m.G0 = v12 + int32(80)
						return
					}
				} else {
					v33 = *(*int32)(unsafe.Add(mBase, uint32(v12)+32))
					if int32(0) < v33 {
						v36 = *(*int32)(unsafe.Add(mBase, uint32(v12)+28))
						v37 = *(*float32)(unsafe.Add(mBase, uint32(v36)))
						*(*float64)(unsafe.Add(mBase, uint32(l3))) = base.F64_promote_f32(v37)
					} else {
					}
					F_free_attstatsslot(m, v12+int32(8))
					mBase = m.M
					v43 = m.ExcPending
					if v43 != 0 {
						return
					} else {
						v45 = v12 + int32(48)
						v47 = v12 + int32(47)
						v48 = int32(0)
						v49 = float64(0)
						*(*uint8)(unsafe.Add(mBase, uint32(v47))) = uint8(v48)
						v53 = *(*int32)(unsafe.Add(mBase, uint32(v45)+8))
						if v53 != 0 {
							v54 = *(*int32)(unsafe.Add(mBase, uint32(v53)+16))
							v55 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v54)+22)))
							v56 = v54 + v55
							v57 = *(*float32)(unsafe.Add(mBase, uint32(v56)+8))
							v59 = *(*float32)(unsafe.Add(mBase, uint32(v56)+16))
							v86 = base.F64_promote_f32(v59)
							v87 = base.F64_promote_f32(v57)
						} else {
							v61 = *(*int32)(unsafe.Add(mBase, uint32(v45)+16))
							if v61 == int32(16) {
								v86 = float64(2)
								v87 = v49
							} else {
								v65 = *(*int32)(unsafe.Add(mBase, uint32(v45)+4))
								if v65 == int32(0) {
									v72 = *(*int32)(unsafe.Add(mBase, uint32(v45)))
									if v72 == int32(0) {
										v86 = float64(0)
										v87 = v49
									} else {
										v75 = *(*int32)(unsafe.Add(mBase, uint32(v72)))
										if v75 != int32(6) {
											v86 = float64(0)
											v87 = v49
										} else {
											v79 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v72)+8)))
											switch v79 - int32(_a_F_estimate_hash_bucket_stats_0) {
											case 0:
												v86 = float64(1)
												v87 = v49
											default:
												v86 = float64(0)
												v87 = v49
											case 5:
												v86 = float64(-1)
												v87 = v49
											}
										}
									}
								} else {
									v68 = *(*int32)(unsafe.Add(mBase, uint32(v65)+76))
									if v68 != int32(5) {
										v72 = *(*int32)(unsafe.Add(mBase, uint32(v45)))
										if v72 == int32(0) {
											v86 = float64(0)
											v87 = v49
										} else {
											v75 = *(*int32)(unsafe.Add(mBase, uint32(v72)))
											if v75 != int32(6) {
												v86 = float64(0)
												v87 = v49
											} else {
												v79 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v72)+8)))
												switch v79 - int32(_a_F_estimate_hash_bucket_stats_0) {
												case 0:
													v86 = float64(1)
													v87 = v49
												default:
													v86 = float64(0)
													v87 = v49
												case 5:
													v86 = float64(-1)
													v87 = v49
												}
											}
										}
									} else {
										v86 = float64(-1)
										v87 = v49
									}
								}
							}
						}
						v91 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v45)+28)))
						if v91 != 0 {
							v92 = base.F64_neg(base.F64_sub(float64(1), v87))
						} else {
							v92 = v86
						}
						if base.F64_gt(v92, float64(0)) != 0 {
							v95 = F_clamp_row_est(m, v92)
							mBase = m.M
							v121 = v95
						} else {
							v96 = *(*int32)(unsafe.Add(mBase, uint32(v45)+4))
							if v96 == int32(0) {
								v99 = int32(1)
								*(*uint8)(unsafe.Add(mBase, uint32(v47))) = uint8(v99)
								v121 = float64(200)
							} else {
								v102 = *(*float64)(unsafe.Add(mBase, uint32(v96)+120))
								if base.F64_le(v102, float64(0)) != 0 {
									v105 = int32(1)
									*(*uint8)(unsafe.Add(mBase, uint32(v47))) = uint8(v105)
									v121 = float64(200)
								} else {
									if base.F64_lt(v92, float64(0)) != 0 {
										v112 = F_clamp_row_est(m, base.F64_mul(v102, base.F64_neg(v92)))
										mBase = m.M
										v121 = v112
									} else {
										if base.F64_lt(v102, float64(200)) != 0 {
											v115 = F_clamp_row_est(m, v102)
											mBase = m.M
											v121 = v115
										} else {
											v116 = int32(1)
											*(*uint8)(unsafe.Add(mBase, uint32(v47))) = uint8(v116)
											v121 = float64(200)
										}
									}
								}
							}
						}
						v122 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v12)+47)))
						if v122 == int32(1) {
							v125 = float64(0.1)
							v126 = *(*float64)(unsafe.Add(mBase, uint32(l3)))
							if base.F64_lt(v126, v125) != 0 {
								v129 = v125
							} else {
								v129 = v126
							}
							*(*float64)(unsafe.Add(mBase, uint32(l4))) = v129
						} else {
							v132 = *(*int32)(unsafe.Add(mBase, uint32(v12)+56))
							if v132 != 0 {
								v133 = *(*int32)(unsafe.Add(mBase, uint32(v132)+16))
								v134 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v133)+22)))
								v136 = *(*float32)(unsafe.Add(mBase, uint32(v133+v134)+8))
								v140 = base.F64_promote_f32(v136)
							} else {
								v140 = float64(0)
							}
							v142 = base.F64_div(base.F64_sub(float64(1), v140), v121)
							v143 = *(*int32)(unsafe.Add(mBase, uint32(v12)+52))
							if v143 == int32(0) {
								v168 = v121
							} else {
								v146 = *(*float64)(unsafe.Add(mBase, uint32(v143)+120))
								if base.F64_gt(v146, float64(0)) == int32(0) {
									v168 = v121
								} else {
									v151 = *(*float64)(unsafe.Add(mBase, uint32(v143)+16))
									v153 = base.F64_mul(v121, base.F64_div(v151, v146))
									v154 = float64(1e+100)
									if base.B2i32(base.Ui64(int64(9218868437227405312)) < base.Ui64(base.I64_reinterpret_f64(v153)&int64(9223372036854775807)))|base.F64_gt(v153, v154) != 0 {
										v167 = v154
									} else {
										v163 = float64(1)
										if base.F64_le(v153, v163) != 0 {
											v167 = v163
										} else {
											v167 = base.F64_nearest(v153)
										}
									}
									v168 = v167
								}
							}
							if base.F64_lt(l2, v168) != 0 {
								v172 = l2
							} else {
								v172 = v168
							}
							v173 = base.F64_div(float64(1), v172)
							if base.F64_gt(v142, float64(0)) == int32(0) {
								v185 = v173
							} else {
								v178 = *(*float64)(unsafe.Add(mBase, uint32(l3)))
								if base.F64_gt(v178, v142) == int32(0) {
									v185 = v173
								} else {
									v185 = base.F64_mul(v173, base.F64_div(v178, v142))
								}
							}
							v186 = float64(1e-06)
							if base.F64_lt(v185, v186) != 0 {
								v194 = v186
							} else {
								if base.F64_gt(v185, float64(1)) == int32(0) {
									v194 = v185
								} else {
									v194 = float64(1)
								}
							}
							*(*float64)(unsafe.Add(mBase, uint32(l4))) = v194
						}
						v201 = *(*int32)(unsafe.Add(mBase, uint32(v12)+56))
						if v201 != 0 {
							v202 = *(*int32)(unsafe.Add(mBase, uint32(v12)+60))
							m.T0[v202].(func(*base.Module, int32))(m, v201)
							mBase = m.M
							v204 = m.ExcPending
							if v204 != 0 {
								return
							} else {
								m.G0 = v12 + int32(80)
								return
							}
						} else {
							m.G0 = v12 + int32(80)
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
	var v8 int32
	_ = v8
	var v10 int32
	_ = v10
	var v12 int32
	_ = v12
	var v13 int32
	_ = v13
	var v14 int32
	_ = v14
	var v17 float64
	_ = v17
	var v18 int32
	_ = v18
	var v23 int32
	_ = v23
	var v24 int32
	_ = v24
	var v30 int32
	_ = v30
	var v31 int32
	_ = v31
	var v37 int32
	_ = v37
	var v38 int32
	_ = v38
	var v39 int32
	_ = v39
	var v43 int32
	_ = v43
	var v49 int32
	_ = v49
	var v50 int32
	_ = v50
	var v51 int32
	_ = v51
	var v57 int32
	_ = v57
	var v63 int32
	_ = v63
	var v69 int32
	_ = v69
	var v70 int32
	_ = v70
	var v71 int32
	_ = v71
	var v78 int32
	_ = v78
	var v79 int32
	_ = v79
	var v80 int32
	_ = v80
	var v84 int32
	_ = v84
	var v90 int32
	_ = v90
	var v91 int32
	_ = v91
	var v92 int64
	_ = v92
	var v93 int32
	_ = v93
	var v98 int32
	_ = v98
	var v104 int32
	_ = v104
	var v105 int32
	_ = v105
	var v108 int32
	_ = v108
	var v109 int32
	_ = v109
	var v113 int32
	_ = v113
	var v119 int32
	_ = v119
	var v120 float64
	_ = v120
	var v126 int32
	_ = v126
	var v127 int32
	_ = v127
	var v128 int32
	_ = v128
	var v129 int32
	_ = v129
	var v134 int32
	_ = v134
	var v140 int32
	_ = v140
	v8 = m.G0
	v10 = v8 - int32(176)
	m.G0 = v10
	v12 = *(*int32)(unsafe.Add(mBase, uint32(l1)+96))
	v13 = *(*int32)(unsafe.Add(mBase, uint32(v12)+64))
	v14 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v12)+48)))
	switch int32(base.Ui32(v14) >> (uint(int32(4)) % 32)) {
	case 0:
		v17 = *(*float64)(unsafe.Add(mBase, uint32(v13)))
		v18 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v13)+12)))
		*(*int32)(unsafe.Add(mBase, uint32(v10)+8)) = v18
		*(*float64)(unsafe.Add(mBase, uint32(v10))) = v17
		F_appendStringInfo(m, l0, int32(_a_F_hash_desc_0), v10)
		mBase = m.M
		v23 = m.ExcPending
		if v23 != 0 {
			return
		} else {
			m.G0 = v10 + int32(176)
			return
		}
	case 1:
		v24 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v13))))
		*(*int32)(unsafe.Add(mBase, uint32(v10)+16)) = v24
		F_appendStringInfo(m, l0, int32(_a_F_hash_desc_1), v10+int32(16))
		mBase = m.M
		v30 = m.ExcPending
		if v30 != 0 {
			return
		} else {
			m.G0 = v10 + int32(176)
			return
		}
	case 2:
		v31 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v13))))
		*(*int32)(unsafe.Add(mBase, uint32(v10)+32)) = v31
		F_appendStringInfo(m, l0, int32(_a_F_hash_desc_2), v10+int32(32))
		mBase = m.M
		v37 = m.ExcPending
		if v37 != 0 {
			return
		} else {
			m.G0 = v10 + int32(176)
			return
		}
	case 3:
		v38 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v13)+2)))
		v39 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v13))))
		*(*int32)(unsafe.Add(mBase, uint32(v10)+48)) = v39
		if v38 != 0 {
			v43 = int32(84)
		} else {
			v43 = int32(70)
		}
		*(*int32)(unsafe.Add(mBase, uint32(v10)+52)) = v43
		F_appendStringInfo(m, l0, int32(_a_F_hash_desc_3), v10+int32(48))
		mBase = m.M
		v49 = m.ExcPending
		if v49 != 0 {
			return
		} else {
			m.G0 = v10 + int32(176)
			return
		}
	case 4:
		v50 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v13)+8)))
		v51 = *(*int32)(unsafe.Add(mBase, uint32(v13)))
		*(*int32)(unsafe.Add(mBase, uint32(v10)+64)) = v51
		if v50&int32(2) != 0 {
			v57 = int32(84)
		} else {
			v57 = int32(70)
		}
		*(*int32)(unsafe.Add(mBase, uint32(v10)+72)) = v57
		if v50&int32(1) != 0 {
			v63 = int32(84)
		} else {
			v63 = int32(70)
		}
		*(*int32)(unsafe.Add(mBase, uint32(v10)+68)) = v63
		F_appendStringInfo(m, l0, int32(_a_F_hash_desc_4), v10-int32(-64))
		mBase = m.M
		v69 = m.ExcPending
		if v69 != 0 {
			return
		} else {
			m.G0 = v10 + int32(176)
			return
		}
	default:
		m.G0 = v10 + int32(176)
		return
	case 6:
		v70 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v13))))
		v71 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v13)+2)))
		*(*int32)(unsafe.Add(mBase, uint32(v10)+84)) = v71
		*(*int32)(unsafe.Add(mBase, uint32(v10)+80)) = v70
		F_appendStringInfo(m, l0, int32(_a_F_hash_desc_5), v10+int32(80))
		mBase = m.M
		v78 = m.ExcPending
		if v78 != 0 {
			return
		} else {
			m.G0 = v10 + int32(176)
			return
		}
	case 7:
		v79 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v13)+2)))
		v80 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v13))))
		*(*int32)(unsafe.Add(mBase, uint32(v10)+96)) = v80
		if v79 != 0 {
			v84 = int32(84)
		} else {
			v84 = int32(70)
		}
		*(*int32)(unsafe.Add(mBase, uint32(v10)+100)) = v84
		F_appendStringInfo(m, l0, int32(_a_F_hash_desc_6), v10+int32(96))
		mBase = m.M
		v90 = m.ExcPending
		if v90 != 0 {
			return
		} else {
			m.G0 = v10 + int32(176)
			return
		}
	case 8:
		v91 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v13)+10)))
		v92 = *(*int64)(unsafe.Add(mBase, uint32(v13)))
		v93 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v13)+8)))
		*(*int32)(unsafe.Add(mBase, uint32(v10)+120)) = v93
		*(*int64)(unsafe.Add(mBase, uint32(v10)+112)) = v92
		if v91 != 0 {
			v98 = int32(84)
		} else {
			v98 = int32(70)
		}
		*(*int32)(unsafe.Add(mBase, uint32(v10)+124)) = v98
		F_appendStringInfo(m, l0, int32(_a_F_hash_desc_7), v10+int32(112))
		mBase = m.M
		v104 = m.ExcPending
		if v104 != 0 {
			return
		} else {
			m.G0 = v10 + int32(176)
			return
		}
	case 9:
		v105 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v13))))
		v108 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v13)+1)))
		if v108 != 0 {
			v109 = int32(84)
		} else {
			v109 = int32(70)
		}
		*(*int32)(unsafe.Add(mBase, uint32(v10)+132)) = v109
		if v105 != 0 {
			v113 = int32(84)
		} else {
			v113 = int32(70)
		}
		*(*int32)(unsafe.Add(mBase, uint32(v10)+128)) = v113
		F_appendStringInfo(m, l0, int32(_a_F_hash_desc_8), v10+int32(128))
		mBase = m.M
		v119 = m.ExcPending
		if v119 != 0 {
			return
		} else {
			m.G0 = v10 + int32(176)
			return
		}
	case 11:
		v120 = *(*float64)(unsafe.Add(mBase, uint32(v13)))
		*(*float64)(unsafe.Add(mBase, uint32(v10)+144)) = v120
		F_appendStringInfo(m, l0, int32(_a_F_hash_desc_9), v10+int32(144))
		mBase = m.M
		v126 = m.ExcPending
		if v126 != 0 {
			return
		} else {
			m.G0 = v10 + int32(176)
			return
		}
	case 12:
		v127 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v13)+6)))
		v128 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v13)+4)))
		v129 = *(*int32)(unsafe.Add(mBase, uint32(v13)))
		*(*int32)(unsafe.Add(mBase, uint32(v10)+164)) = v129
		*(*int32)(unsafe.Add(mBase, uint32(v10)+160)) = v128
		if v127 != 0 {
			v134 = int32(84)
		} else {
			v134 = int32(70)
		}
		*(*int32)(unsafe.Add(mBase, uint32(v10)+168)) = v134
		F_appendStringInfo(m, l0, int32(_a_F_hash_desc_10), v10+int32(160))
		mBase = m.M
		v140 = m.ExcPending
		if v140 != 0 {
			return
		} else {
			m.G0 = v10 + int32(176)
			return
		}
	}
}
func F_hash_numeric_extended(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
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
	var v19 int32
	_ = v19
	var v20 int32
	_ = v20
	var v24 int32
	_ = v24
	var v37 int32
	_ = v37
	var v39 int32
	_ = v39
	var v40 int32
	_ = v40
	var v41 int32
	_ = v41
	var v44 int32
	_ = v44
	var v46 int32
	_ = v46
	var v49 int32
	_ = v49
	var v51 int32
	_ = v51
	var v53 int32
	_ = v53
	var v56 int32
	_ = v56
	var v57 int32
	_ = v57
	var v60 int32
	_ = v60
	var v71 int32
	_ = v71
	var v74 int32
	_ = v74
	var v77 int32
	_ = v77
	var v81 int32
	_ = v81
	var v84 int32
	_ = v84
	var v91 int32
	_ = v91
	var v92 int32
	_ = v92
	var v96 int32
	_ = v96
	var v98 int32
	_ = v98
	var v100 int32
	_ = v100
	var v105 int32
	_ = v105
	var v106 int32
	_ = v106
	var v112 int32
	_ = v112
	var v118 int32
	_ = v118
	var v122 int32
	_ = v122
	var v123 int32
	_ = v123
	var v127 int32
	_ = v127
	var v129 int32
	_ = v129
	var v133 int32
	_ = v133
	var v137 int32
	_ = v137
	var v138 int32
	_ = v138
	var v139 int32
	_ = v139
	var v140 int32
	_ = v140
	var v144 int32
	_ = v144
	var v148 int32
	_ = v148
	var v153 int32
	_ = v153
	var v155 int32
	_ = v155
	var v157 int32
	_ = v157
	var v159 int32
	_ = v159
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
	var v171 int32
	_ = v171
	var v172 int32
	_ = v172
	var v173 int32
	_ = v173
	var v175 int32
	_ = v175
	var v176 int32
	_ = v176
	var v178 int32
	_ = v178
	var v180 int32
	_ = v180
	var v184 int32
	_ = v184
	var v185 int32
	_ = v185
	var v186 int32
	_ = v186
	var v187 int32
	_ = v187
	var v191 int32
	_ = v191
	var v195 int32
	_ = v195
	var v199 int32
	_ = v199
	var v200 int32
	_ = v200
	var v201 int32
	_ = v201
	var v202 int32
	_ = v202
	var v206 int32
	_ = v206
	var v207 int32
	_ = v207
	var v208 int32
	_ = v208
	var v210 int32
	_ = v210
	var v213 int32
	_ = v213
	var v214 int32
	_ = v214
	var v216 int32
	_ = v216
	var v217 int32
	_ = v217
	var v218 int32
	_ = v218
	var v224 int32
	_ = v224
	var v225 int32
	_ = v225
	var v227 int32
	_ = v227
	var v228 int32
	_ = v228
	var v229 int32
	_ = v229
	var v231 int32
	_ = v231
	var v232 int32
	_ = v232
	var v233 int32
	_ = v233
	var v235 int32
	_ = v235
	var v236 int32
	_ = v236
	var v238 int32
	_ = v238
	var v240 int32
	_ = v240
	var v244 int32
	_ = v244
	var v245 int32
	_ = v245
	var v246 int32
	_ = v246
	var v247 int32
	_ = v247
	var v251 int32
	_ = v251
	var v255 int32
	_ = v255
	var v259 int32
	_ = v259
	var v260 int32
	_ = v260
	var v261 int32
	_ = v261
	var v262 int32
	_ = v262
	var v266 int32
	_ = v266
	var v267 int32
	_ = v267
	var v268 int32
	_ = v268
	var v270 int32
	_ = v270
	var v273 int32
	_ = v273
	var v274 int32
	_ = v274
	var v276 int32
	_ = v276
	var v277 int32
	_ = v277
	var v278 int32
	_ = v278
	var v282 int32
	_ = v282
	var v286 int32
	_ = v286
	var v287 int32
	_ = v287
	var v291 int32
	_ = v291
	var v292 int32
	_ = v292
	var v296 int32
	_ = v296
	var v297 int32
	_ = v297
	var v299 int32
	_ = v299
	var v301 int32
	_ = v301
	var v305 int32
	_ = v305
	var v306 int32
	_ = v306
	var v310 int32
	_ = v310
	var v311 int32
	_ = v311
	var v313 int32
	_ = v313
	var v314 int32
	_ = v314
	var v316 int32
	_ = v316
	var v320 int32
	_ = v320
	var v321 int32
	_ = v321
	var v325 int32
	_ = v325
	var v326 int32
	_ = v326
	var v328 int32
	_ = v328
	var v332 int32
	_ = v332
	var v333 int32
	_ = v333
	var v337 int32
	_ = v337
	var v338 int32
	_ = v338
	var v342 int32
	_ = v342
	var v343 int32
	_ = v343
	var v347 int32
	_ = v347
	var v348 int32
	_ = v348
	var v349 int32
	_ = v349
	var v353 int32
	_ = v353
	var v354 int32
	_ = v354
	var v355 int32
	_ = v355
	var v359 int32
	_ = v359
	var v360 int32
	_ = v360
	var v361 int32
	_ = v361
	var v363 int32
	_ = v363
	var v364 int32
	_ = v364
	var v365 int32
	_ = v365
	var v369 int32
	_ = v369
	var v370 int32
	_ = v370
	var v371 int32
	_ = v371
	var v372 int32
	_ = v372
	var v376 int32
	_ = v376
	var v377 int32
	_ = v377
	var v378 int32
	_ = v378
	var v379 int32
	_ = v379
	var v383 int32
	_ = v383
	var v384 int32
	_ = v384
	var v385 int32
	_ = v385
	var v386 int32
	_ = v386
	var v391 int32
	_ = v391
	var v392 int32
	_ = v392
	var v393 int32
	_ = v393
	var v396 int32
	_ = v396
	var v398 int32
	_ = v398
	var v402 int32
	_ = v402
	var v406 int32
	_ = v406
	var v410 int32
	_ = v410
	var v414 int32
	_ = v414
	var v418 int32
	_ = v418
	var v428 int32
	_ = v428
	var v429 int32
	_ = v429
	var v430 int64
	_ = v430
	var v444 int64
	_ = v444
	var v445 int32
	_ = v445
	var v446 int32
	_ = v446
	var v461 int32
	_ = v461
	var v462 int32
	_ = v462
	v12 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
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
	return int32(0)
L2:
	;
	v17 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
	v18 = *(*int64)(unsafe.Add(mBase, uint32(v17)))
	v19 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v13)+4)))
	v20 = int32(_a_F_hash_numeric_extended_0)
	if v19&v20 != v20 {
		goto L4
	} else {
		goto L5
	}
L3:
	;
	v461 = F_Int64GetDatum(m, v18-int64(1))
	mBase = m.M
	v462 = m.ExcPending
	if v462 != 0 {
		goto L1
	} else {
		goto L75
	}
L4:
	;
	v24 = base.I32_extend16_s(v19)
	if v24 < int32(0) {
		goto L8
	} else {
		goto L9
	}
L5:
	;
	v444 = v18
	goto L6
L6:
	;
	v445 = F_Int64GetDatum(m, v444)
	mBase = m.M
	v446 = m.ExcPending
	if v446 != 0 {
		goto L1
	} else {
		goto L74
	}
L7:
	;
	v41 = *(*int32)(unsafe.Add(mBase, uint32(v13)))
	v44 = v40 + int32(base.Ui32(v41)>>(uint(int32(2))%32))
	v46 = int32(base.Ui32(v44) >> (uint(int32(1)) % 32))
	if v46 == int32(0) {
		goto L3
	} else {
		goto L11
	}
L8:
	;
	v39 = v19<<(uint(int32(25))%32)>>(uint(int32(31))%32)&int32(-64) | v19&int32(63)
	v40 = int32(-6)
	goto L7
L9:
	;
	goto L10
L10:
	;
	v37 = int32(*(*int16)(unsafe.Add(mBase, uint32(v13)+6)))
	v39 = v37
	v40 = int32(-8)
	goto L7
L11:
	;
	v49 = int32(0)
	v51 = v13 + int32(6)
	v53 = v13 + int32(8)
	if v24 < v49 {
		goto L12
	} else {
		goto L13
	}
L12:
	;
	v56 = v51
	goto L14
L13:
	;
	v56 = v53
	goto L14
L14:
	;
	v57 = v49
	v60 = v39
	goto L15
L15:
	;
	v71 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v56+v57<<(uint(int32(1))%32)))))
	if v71 == int32(0) {
		goto L17
	} else {
		goto L18
	}
L16:
	;
	if v57 == v46 {
		goto L3
	} else {
		goto L21
	}
L17:
	;
	v74 = int32(1)
	v77 = v57 + v74
	if v77 != v46 {
		v57 = v77
		v60 = v60 - v74
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
	v81 = v46
	v84 = int32(0)
	goto L23
L22:
	;
	if int32(0) <= v24 {
		goto L27
	} else {
		goto L28
	}
L23:
	;
	v91 = int32(1)
	v92 = v81 - v91
	v96 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v56+v92<<(uint(v91)%32)))))
	if v96 != 0 {
		v100 = v84
		goto L22
	} else {
		goto L25
	}
L24:
	;
	v100 = v46
	goto L22
L25:
	;
	v98 = v84 + int32(1)
	if v98 != v46 {
		v81 = v92
		v84 = v98
		goto L23
	} else {
		goto L26
	}
L26:
	;
	goto L24
L27:
	;
	v105 = v53
	goto L29
L28:
	;
	v105 = v51
	goto L29
L29:
	;
	v106 = v57<<(uint(int32(1))%32) + v105
	v112 = (v44 - (v57+v100)<<(uint(int32(1))%32)) & int32(-2)
	v118 = v112 - int32(1636608432)
	if v18 == int64(0) {
		goto L32
	} else {
		goto L33
	}
L30:
	;
	v428 = F_Int64GetDatum(m, base.I64_extend_i32_u(v418)<<(uint(int64(32))%64)|base.I64_extend_i32_u(v418^v410-base.I32_rotl(v418, int32(24))))
	mBase = m.M
	v429 = m.ExcPending
	if v429 != 0 {
		goto L1
	} else {
		goto L73
	}
L31:
	;
	if v106&int32(3) != 0 {
		goto L47
	} else {
		goto L48
	}
L32:
	;
	v155 = v118
	v157 = v118
	v159 = v118
	goto L31
L33:
	;
	goto L34
L34:
	;
	v122 = v118 + base.I32_wrap_i64(v18)
	v123 = v122 + v118
	v127 = int32(4)
	v129 = base.I32_wrap_i64(int64(base.Ui64(v18)>>(uint(int64(32))%64))) ^ base.I32_rotl(v118, v127)
	v133 = v122 - v129 ^ base.I32_rotl(v129, int32(6))
	v137 = v123 - v133 ^ base.I32_rotl(v133, int32(8))
	v138 = v123 + v129
	v139 = v133 + v138
	v140 = v137 + v139
	v144 = v138 - v137 ^ base.I32_rotl(v137, int32(16))
	v148 = v139 - v144 ^ base.I32_rotl(v144, int32(19))
	v153 = v140 + v144
	v155 = v153
	v157 = v140 - v148 ^ base.I32_rotl(v148, v127)
	v159 = v148 + v153
	goto L31
L35:
	;
	v396 = int32(14)
	v398 = v392 ^ v393 - base.I32_rotl(v392, v396)
	v402 = v398 ^ v391 - base.I32_rotl(v398, int32(11))
	v406 = v402 ^ v392 - base.I32_rotl(v402, int32(25))
	v410 = v406 ^ v398 - base.I32_rotl(v406, int32(16))
	v414 = v410 ^ v402 - base.I32_rotl(v410, int32(4))
	v418 = v414 ^ v406 - base.I32_rotl(v414, v396)
	goto L30
L36:
	;
	v386 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v213))))
	v391 = v383 + v386
	v392 = v384
	v393 = v385
	goto L35
L37:
	;
	v379 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v213)+1)))
	v383 = v379<<(uint(int32(8))%32) + v376
	v384 = v377
	v385 = v378
	goto L36
L38:
	;
	v372 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v213)+2)))
	v376 = v372<<(uint(int32(16))%32) + v369
	v377 = v370
	v378 = v371
	goto L37
L39:
	;
	v365 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v213)+3)))
	v369 = v365<<(uint(int32(24))%32) + v216
	v370 = v363
	v371 = v364
	goto L38
L40:
	;
	v361 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v213)+4)))
	v363 = v359 + v361
	v364 = v360
	goto L39
L41:
	;
	v355 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v213)+5)))
	v359 = v355<<(uint(int32(8))%32) + v353
	v360 = v354
	goto L40
L42:
	;
	v349 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v213)+6)))
	v353 = v349<<(uint(int32(16))%32) + v347
	v354 = v348
	goto L41
L43:
	;
	v343 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v213)+7)))
	v347 = v343<<(uint(int32(24))%32) + v217
	v348 = v342
	goto L42
L44:
	;
	v338 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v213)+8)))
	v342 = v338<<(uint(int32(8))%32) + v337
	goto L43
L45:
	;
	v333 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v213)+9)))
	v337 = v333<<(uint(int32(16))%32) + v332
	goto L44
L46:
	;
	v328 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v213)+10)))
	v332 = v328<<(uint(int32(24))%32) + v218
	goto L45
L47:
	;
	if base.Ui32(int32(11)) < base.Ui32(v112) {
		goto L50
	} else {
		goto L51
	}
L48:
	;
	goto L49
L49:
	;
	if base.Ui32(int32(12)) <= base.Ui32(v112) {
		goto L56
	} else {
		goto L57
	}
L50:
	;
	v164 = v106
	v165 = v112
	v167 = v155
	v168 = v159
	v169 = v157
	goto L53
L51:
	;
	v213 = v106
	v214 = v112
	v216 = v155
	v217 = v159
	v218 = v157
	goto L52
L52:
	;
	switch v214 - int32(1) {
	case 0:
		v383 = v216
		v384 = v217
		v385 = v218
		goto L36
	case 1:
		v376 = v216
		v377 = v217
		v378 = v218
		goto L37
	case 2:
		v369 = v216
		v370 = v217
		v371 = v218
		goto L38
	case 3:
		v363 = v217
		v364 = v218
		goto L39
	case 4:
		v359 = v217
		v360 = v218
		goto L40
	case 5:
		v353 = v217
		v354 = v218
		goto L41
	case 6:
		v347 = v217
		v348 = v218
		goto L42
	case 7:
		v342 = v218
		goto L43
	case 8:
		v337 = v218
		goto L44
	case 9:
		v332 = v218
		goto L45
	case 10:
		goto L46
	default:
		v391 = v216
		v392 = v217
		v393 = v218
		goto L35
	}
L53:
	;
	v171 = *(*int32)(unsafe.Add(mBase, uint32(v164)+4))
	v172 = v171 + v168
	v173 = *(*int32)(unsafe.Add(mBase, uint32(v164)))
	v175 = *(*int32)(unsafe.Add(mBase, uint32(v164)+8))
	v176 = v175 + v169
	v178 = int32(4)
	v180 = v173 + v167 - v176 ^ base.I32_rotl(v176, v178)
	v184 = v172 - v180 ^ base.I32_rotl(v180, int32(6))
	v185 = v176 + v172
	v186 = v180 + v185
	v187 = v184 + v186
	v191 = v185 - v184 ^ base.I32_rotl(v184, int32(8))
	v195 = v186 - v191 ^ base.I32_rotl(v191, int32(16))
	v199 = v187 - v195 ^ base.I32_rotl(v195, int32(19))
	v200 = v191 + v187
	v201 = v195 + v200
	v202 = v199 + v201
	v206 = v200 - v199 ^ base.I32_rotl(v199, v178)
	v207 = int32(12)
	v208 = v164 + v207
	v210 = v165 - v207
	if base.Ui32(int32(11)) < base.Ui32(v210) {
		v164 = v208
		v165 = v210
		v167 = v201
		v168 = v202
		v169 = v206
		goto L53
	} else {
		goto L55
	}
L54:
	;
	v213 = v208
	v214 = v210
	v216 = v201
	v217 = v202
	v218 = v206
	goto L52
L55:
	;
	goto L54
L56:
	;
	v224 = v106
	v225 = v112
	v227 = v155
	v228 = v159
	v229 = v157
	goto L59
L57:
	;
	v273 = v106
	v274 = v112
	v276 = v155
	v277 = v159
	v278 = v157
	goto L58
L58:
	;
	switch v274 - int32(1) {
	case 0:
		v325 = v276
		goto L62
	case 1:
		v320 = v276
		goto L63
	case 2:
		goto L64
	case 3:
		v313 = v277
		goto L65
	case 4:
		v310 = v277
		goto L66
	case 5:
		v305 = v277
		goto L67
	case 6:
		goto L68
	case 7:
		v296 = v278
		goto L69
	case 8:
		v291 = v278
		goto L70
	case 9:
		v286 = v278
		goto L71
	case 10:
		goto L72
	default:
		v391 = v276
		v392 = v277
		v393 = v278
		goto L35
	}
L59:
	;
	v231 = *(*int32)(unsafe.Add(mBase, uint32(v224)+4))
	v232 = v231 + v228
	v233 = *(*int32)(unsafe.Add(mBase, uint32(v224)))
	v235 = *(*int32)(unsafe.Add(mBase, uint32(v224)+8))
	v236 = v235 + v229
	v238 = int32(4)
	v240 = v233 + v227 - v236 ^ base.I32_rotl(v236, v238)
	v244 = v232 - v240 ^ base.I32_rotl(v240, int32(6))
	v245 = v236 + v232
	v246 = v240 + v245
	v247 = v244 + v246
	v251 = v245 - v244 ^ base.I32_rotl(v244, int32(8))
	v255 = v246 - v251 ^ base.I32_rotl(v251, int32(16))
	v259 = v247 - v255 ^ base.I32_rotl(v255, int32(19))
	v260 = v251 + v247
	v261 = v255 + v260
	v262 = v259 + v261
	v266 = v260 - v259 ^ base.I32_rotl(v259, v238)
	v267 = int32(12)
	v268 = v224 + v267
	v270 = v225 - v267
	if base.Ui32(int32(11)) < base.Ui32(v270) {
		v224 = v268
		v225 = v270
		v227 = v261
		v228 = v262
		v229 = v266
		goto L59
	} else {
		goto L61
	}
L60:
	;
	v273 = v268
	v274 = v270
	v276 = v261
	v277 = v262
	v278 = v266
	goto L58
L61:
	;
	goto L60
L62:
	;
	v326 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v273))))
	v391 = v325 + v326
	v392 = v277
	v393 = v278
	goto L35
L63:
	;
	v321 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v273)+1)))
	v325 = v321<<(uint(int32(8))%32) + v320
	goto L62
L64:
	;
	v316 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v273)+2)))
	v320 = v316<<(uint(int32(16))%32) + v276
	goto L63
L65:
	;
	v314 = *(*int32)(unsafe.Add(mBase, uint32(v273)))
	v391 = v314 + v276
	v392 = v313
	v393 = v278
	goto L35
L66:
	;
	v311 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v273)+4)))
	v313 = v310 + v311
	goto L65
L67:
	;
	v306 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v273)+5)))
	v310 = v306<<(uint(int32(8))%32) + v305
	goto L66
L68:
	;
	v301 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v273)+6)))
	v305 = v301<<(uint(int32(16))%32) + v277
	goto L67
L69:
	;
	v297 = *(*int32)(unsafe.Add(mBase, uint32(v273)))
	v299 = *(*int32)(unsafe.Add(mBase, uint32(v273)+4))
	v391 = v297 + v276
	v392 = v299 + v277
	v393 = v296
	goto L35
L70:
	;
	v292 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v273)+8)))
	v296 = v292<<(uint(int32(8))%32) + v291
	goto L69
L71:
	;
	v287 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v273)+9)))
	v291 = v287<<(uint(int32(16))%32) + v286
	goto L70
L72:
	;
	v282 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v273)+10)))
	v286 = v282<<(uint(int32(24))%32) + v278
	goto L71
L73:
	;
	v430 = *(*int64)(unsafe.Add(mBase, uint32(v428)))
	v444 = v430 ^ base.I64_extend_i32_s(v60)
	goto L6
L74:
	;
	return v445
L75:
	;
	return v461
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
	var v23 int32
	_ = v23
	var v24 int32
	_ = v24
	var v25 int32
	_ = v25
	var v28 int32
	_ = v28
	var v29 int32
	_ = v29
	var v30 int32
	_ = v30
	var v32 int32
	_ = v32
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
	var v43 int32
	_ = v43
	var v44 int32
	_ = v44
	var v46 int32
	_ = v46
	var v48 int32
	_ = v48
	var v52 int32
	_ = v52
	var v60 int32
	_ = v60
	var v64 int32
	_ = v64
	var v69 int32
	_ = v69
	v2 = int32(0)
	v5 = m.G0
	v7 = v5 - int32(16)
	m.G0 = v7
	v9 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
	if v9 == v2 {
		v52 = v2
		m.G0 = v7 + int32(16)
		return v52
	} else {
		v12 = *(*int32)(unsafe.Add(mBase, uint32(v9)+4))
		if v12 != int32(2) {
			v52 = v2
			m.G0 = v7 + int32(16)
			return v52
		} else {
			v15 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
			if base.B2i32(v15 != int32(2988))&base.B2i32(v15 != int32(1070)) == int32(0) {
				v23 = *(*int32)(unsafe.Add(mBase, uint32(v9)+12))
				v24 = *(*int32)(unsafe.Add(mBase, uint32(v23)))
				v25 = F_exprType(m, v24)
				mBase = m.M
				v28 = m.ExcPending
				if v28 != 0 {
					return int32(0)
				} else {
					v29 = F_op_hashjoinable(m, v15, v25)
					mBase = m.M
					v30 = m.ExcPending
					if v30 != 0 {
						return int32(0)
					} else {
						v52 = v29
						m.G0 = v7 + int32(16)
						return v52
					}
				}
			} else {
				v32 = F_SearchSysCache1(m, int32(40), v15)
				mBase = m.M
				v33 = m.ExcPending
				if v33 != 0 {
					return int32(0)
				} else {
					if v32 == int32(0) {
						F_errstart_cold(m, int32(21), int32(0))
						mBase = m.M
						v60 = m.ExcPending
						if v60 != 0 {
							return int32(0)
						} else {
							*(*int32)(unsafe.Add(mBase, uint32(v7))) = v15
							F_errmsg_internal(m, int32(_a_F_hash_ok_operator_0), v7)
							mBase = m.M
							v64 = m.ExcPending
							if v64 != 0 {
								return int32(0)
							} else {
								F_errfinish(m, int32(_a_F_hash_ok_operator_1), int32(855), int32(_a_F_hash_ok_operator_2))
								mBase = m.M
								v69 = m.ExcPending
								if v69 != 0 {
									return int32(0)
								} else {
									base.Wasm_trap_unreachable()
									for {
									}
								}
							}
						}
					} else {
						v36 = *(*int32)(unsafe.Add(mBase, uint32(v32)+16))
						v37 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v36)+22)))
						v38 = v36 + v37
						v39 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v38)+78)))
						if v39 == int32(1) {
							v42 = *(*int32)(unsafe.Add(mBase, uint32(v38)+100))
							v43 = F_func_strict(m, v42)
							mBase = m.M
							v44 = m.ExcPending
							if v44 != 0 {
								return int32(0)
							} else {
								if v43 != 0 {
									F_ReleaseCatCache(m, v32)
									mBase = m.M
									v48 = m.ExcPending
									if v48 != 0 {
										return int32(0)
									} else {
										v52 = int32(1)
										m.G0 = v7 + int32(16)
										return v52
									}
								} else {
									F_ReleaseCatCache(m, v32)
									mBase = m.M
									v46 = m.ExcPending
									if v46 != 0 {
										return int32(0)
									} else {
										v52 = v2
										m.G0 = v7 + int32(16)
										return v52
									}
								}
							}
						} else {
							F_ReleaseCatCache(m, v32)
							mBase = m.M
							v46 = m.ExcPending
							if v46 != 0 {
								return int32(0)
							} else {
								v52 = v2
								m.G0 = v7 + int32(16)
								return v52
							}
						}
					}
				}
			}
		}
	}
}
func F_hash_select_dirsize(m *base.Module, l0 int32) int32 {
	var v2 int32
	_ = v2
	var v8 int32
	_ = v8
	var v11 int32
	_ = v11
	var v25 int32
	_ = v25
	v2 = int32(0)
	v8 = int32(1073741823)
	if v8 <= l0 {
		v11 = v8
	} else {
		v11 = l0
	}
	v25 = int32(256)
	for {
		if v25 < int32(1)<<(uint(v2-base.I32_clz(int32(base.Ui32(int32(-1)<<(uint(v2-base.I32_clz(v11-int32(1)))%32)^int32(-1))>>(uint(int32(8))%32))))%32) {
			v25 = v25 << (uint(int32(1)) % 32)
			continue
		} else {
			break
		}
		break
	}
	return v25
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
	var v27 int32
	_ = v27
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
	var v49 int32
	_ = v49
	var v53 int32
	_ = v53
	var v54 int32
	_ = v54
	var v55 int32
	_ = v55
	var v57 int32
	_ = v57
	var v58 int32
	_ = v58
	var v64 int32
	_ = v64
	var v73 int32
	_ = v73
	var v74 int32
	_ = v74
	var v77 int32
	_ = v77
	var v78 int32
	_ = v78
	var v79 int32
	_ = v79
	var v82 int32
	_ = v82
	var v84 int32
	_ = v84
	var v88 int32
	_ = v88
	var v90 int32
	_ = v90
	var v92 int32
	_ = v92
	var v94 int32
	_ = v94
	var v97 int32
	_ = v97
	var v98 int32
	_ = v98
	var v111 int32
	_ = v111
	var v113 int32
	_ = v113
	var v115 int32
	_ = v115
	var v116 int32
	_ = v116
	var v123 int32
	_ = v123
	var v124 int32
	_ = v124
	var v125 int32
	_ = v125
	var v129 int32
	_ = v129
	var v130 int32
	_ = v130
	var v133 int32
	_ = v133
	var v135 int32
	_ = v135
	var v136 int32
	_ = v136
	var v137 int32
	_ = v137
	var v140 int32
	_ = v140
	var v141 int32
	_ = v141
	var v144 int32
	_ = v144
	var v146 int32
	_ = v146
	var v147 int32
	_ = v147
	var v149 float64
	_ = v149
	var v151 int32
	_ = v151
	var v152 int32
	_ = v152
	var v154 int32
	_ = v154
	var v162 int32
	_ = v162
	var v165 int32
	_ = v165
	var v166 int32
	_ = v166
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
	v27 = v2
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
	v36 = v33 + v27*int32(52)
	v37 = *(*int32)(unsafe.Add(mBase, uint32(v36)+16))
	v38 = *(*int32)(unsafe.Add(mBase, uint32(v36)))
	v39 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v14)+11)) = uint8(v39)
	*(*int32)(unsafe.Add(mBase, uint32(l0)+188)) = v27
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
	F_slot_getsomeattrs_int(m, v21, v45)
	mBase = m.M
	v49 = m.ExcPending
	if v49 != 0 {
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
	v53 = int32(0)
	goto L13
L12:
	;
	v53 = v14 + int32(11)
	goto L13
L13:
	;
	v54 = *(*int32)(unsafe.Add(mBase, uint32(v37)+8))
	v55 = *(*int32)(unsafe.Add(mBase, uint32(v54)+12))
	m.T0[v55].(func(*base.Module, int32))(m, v37)
	mBase = m.M
	v57 = m.ExcPending
	if v57 != 0 {
		goto L9
	} else {
		goto L14
	}
L14:
	;
	v58 = *(*int32)(unsafe.Add(mBase, uint32(v36)+32))
	if int32(0) < v58 {
		goto L15
	} else {
		goto L16
	}
L15:
	;
	v64 = int32(0)
	goto L18
L16:
	;
	goto L17
L17:
	;
	v111 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v37)+4)))
	v113 = v111 & int32(_a_F_lookup_hash_entries_0)
	*(*uint16)(unsafe.Add(mBase, uint32(v37)+4)) = uint16(v113)
	v115 = *(*int32)(unsafe.Add(mBase, uint32(v37)+12))
	v116 = *(*int32)(unsafe.Add(mBase, uint32(v115)))
	*(*uint16)(unsafe.Add(mBase, uint32(v37)+6)) = uint16(v116)
	goto L21
L18:
	;
	v73 = *(*int32)(unsafe.Add(mBase, uint32(v37)+16))
	v74 = int32(2)
	v77 = *(*int32)(unsafe.Add(mBase, uint32(v21)+16))
	v78 = *(*int32)(unsafe.Add(mBase, uint32(v36)+40))
	v79 = int32(1)
	v82 = int32(*(*int16)(unsafe.Add(mBase, uint32(v78+v64<<(uint(v79)%32)))))
	v84 = v82 - v79
	v88 = *(*int32)(unsafe.Add(mBase, uint32(v77+v84<<(uint(v74)%32))))
	*(*int32)(unsafe.Add(mBase, uint32(v73+v64<<(uint(v74)%32)))) = v88
	v90 = *(*int32)(unsafe.Add(mBase, uint32(v37)+20))
	v92 = *(*int32)(unsafe.Add(mBase, uint32(v21)+20))
	v94 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v92+v84))))
	*(*uint8)(unsafe.Add(mBase, uint32(v90+v64))) = uint8(v94)
	v97 = v64 + v79
	v98 = *(*int32)(unsafe.Add(mBase, uint32(v36)+32))
	if v97 < v98 {
		v64 = v97
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
	v123 = F_LookupTupleHashEntry(m, v38, v37, v53, v14+int32(12))
	mBase = m.M
	v124 = m.ExcPending
	if v124 != 0 {
		goto L9
	} else {
		goto L24
	}
L22:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v19+v27<<(uint(int32(2))%32)))) = v162
	v165 = v27 + int32(1)
	v166 = *(*int32)(unsafe.Add(mBase, uint32(l0)+244))
	if v165 < v166 {
		v27 = v165
		goto L4
	} else {
		goto L38
	}
L23:
	;
	v162 = int32(0)
	goto L22
L24:
	;
	if v123 != 0 {
		goto L25
	} else {
		goto L26
	}
L25:
	;
	v125 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v14)+11)))
	if v125 == int32(1) {
		goto L28
	} else {
		goto L29
	}
L26:
	;
	goto L27
L27:
	;
	v135 = *(*int32)(unsafe.Add(mBase, uint32(l0)+164))
	v136 = *(*int32)(unsafe.Add(mBase, uint32(v135)+12))
	v137 = *(*int32)(unsafe.Add(mBase, uint32(l0)+260))
	v140 = v137 + v27*int32(24)
	v141 = *(*int32)(unsafe.Add(mBase, uint32(v140)+4))
	if v141 == int32(0) {
		goto L33
	} else {
		goto L34
	}
L28:
	;
	F_initialize_hash_entry(m, l0, v38, v123)
	mBase = m.M
	v129 = m.ExcPending
	if v129 != 0 {
		goto L9
	} else {
		goto L31
	}
L29:
	;
	goto L30
L30:
	;
	v130 = *(*int32)(unsafe.Add(mBase, uint32(v38)+32))
	if v130 == int32(0) {
		goto L23
	} else {
		goto L32
	}
L31:
	;
	goto L30
L32:
	;
	v133 = *(*int32)(unsafe.Add(mBase, uint32(v123)))
	v162 = v133 - v130
	goto L22
L33:
	;
	v144 = *(*int32)(unsafe.Add(mBase, uint32(l0)+256))
	v146 = *(*int32)(unsafe.Add(mBase, uint32(v36)+48))
	v147 = *(*int32)(unsafe.Add(mBase, uint32(v146)+96))
	v149 = *(*float64)(unsafe.Add(mBase, uint32(l0)+304))
	F_hashagg_spill_init(m, v140, v144, int32(0), base.F64_convert_i32_s(v147), v149)
	mBase = m.M
	v151 = m.ExcPending
	if v151 != 0 {
		goto L9
	} else {
		goto L36
	}
L34:
	;
	goto L35
L35:
	;
	v152 = *(*int32)(unsafe.Add(mBase, uint32(v14)+12))
	F_hashagg_spill_tuple(m, l0, v140, v136, v152)
	mBase = m.M
	v154 = m.ExcPending
	if v154 != 0 {
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
