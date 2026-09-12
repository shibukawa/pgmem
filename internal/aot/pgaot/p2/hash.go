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
	v11 = *(*int32)(unsafe.Add(mBase, _consts[1]))
	if v11 != 0 {
		F_ProcessInterrupts(m)
		mBase = m.M
		v15 = m.ExcPending
		if v15 != 0 {
			return int32(0)
		} else {
			v16 = int32(8)
			v20 = F_BufFileReadMaybeEOF(m, l0, v8+v16, v16, int32(1))
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
		v20 = F_BufFileReadMaybeEOF(m, l0, v8+v16, v16, int32(1))
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
		v15 = int32(4549024)
		v16 = *(*int32)(unsafe.Add(mBase, _consts[0]))
		v18 = *(*int32)(unsafe.Add(mBase, uint32(l3)+124))
		*(*int32)(unsafe.Add(mBase, _consts[0])) = v18
		v21 = F_BufFileCreateTemp(m, int32(0))
		mBase = m.M
		v22 = m.ExcPending
		if v22 != 0 {
			return
		} else {
			*(*int32)(unsafe.Add(mBase, uint32(l2))) = v21
			*(*int32)(unsafe.Add(mBase, _consts[0])) = v16
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
	var v203 int32
	_ = v203
	var v207 int32
	_ = v207
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
	var v260 int32
	_ = v260
	var v264 int32
	_ = v264
	var v267 int32
	_ = v267
	var v268 int32
	_ = v268
	var v269 int32
	_ = v269
	var v273 int32
	_ = v273
	var v274 int32
	_ = v274
	var v285 int32
	_ = v285
	var v287 int32
	_ = v287
	var v292 int32
	_ = v292
	var v309 int32
	_ = v309
	var v312 int32
	_ = v312
	var v313 int32
	_ = v313
	var v319 int32
	_ = v319
	var v325 int32
	_ = v325
	var v326 int32
	_ = v326
	var v327 int32
	_ = v327
	var v333 int32
	_ = v333
	var v337 int32
	_ = v337
	var v341 int32
	_ = v341
	var v342 int32
	_ = v342
	var v343 int32
	_ = v343
	var v345 int32
	_ = v345
	var v352 int32
	_ = v352
	var v354 int32
	_ = v354
	var v381 int32
	_ = v381
	var v382 int32
	_ = v382
	var v388 int32
	_ = v388
	var v394 int32
	_ = v394
	var v399 int32
	_ = v399
	var v400 int32
	_ = v400
	var v403 int32
	_ = v403
	var v404 int32
	_ = v404
	var v406 int32
	_ = v406
	var v413 int32
	_ = v413
	var v414 int32
	_ = v414
	var v415 int32
	_ = v415
	var v416 int32
	_ = v416
	var v418 int32
	_ = v418
	var v420 int32
	_ = v420
	var v421 int32
	_ = v421
	var v425 int32
	_ = v425
	var v429 int32
	_ = v429
	var v431 int32
	_ = v431
	var v435 int32
	_ = v435
	var v437 int32
	_ = v437
	var v438 float64
	_ = v438
	var v443 int32
	_ = v443
	var v445 int32
	_ = v445
	var v446 int32
	_ = v446
	var v447 int32
	_ = v447
	var v451 int32
	_ = v451
	var v454 int32
	_ = v454
	var v455 int32
	_ = v455
	var v460 int32
	_ = v460
	var v461 int32
	_ = v461
	var v466 int32
	_ = v466
	var v469 int32
	_ = v469
	var v474 int32
	_ = v474
	var v475 int32
	_ = v475
	var v478 int32
	_ = v478
	var v479 int32
	_ = v479
	var v485 int32
	_ = v485
	var v489 int32
	_ = v489
	var v494 int32
	_ = v494
	var v500 int32
	_ = v500
	var v504 int32
	_ = v504
	var v507 int64
	_ = v507
	var v508 int32
	_ = v508
	var v512 int32
	_ = v512
	var v518 int32
	_ = v518
	var v520 int32
	_ = v520
	var v526 int32
	_ = v526
	var v533 int32
	_ = v533
	var v537 int32
	_ = v537
	var v539 int32
	_ = v539
	var v543 int32
	_ = v543
	var v546 int64
	_ = v546
	var v551 int32
	_ = v551
	var v553 int32
	_ = v553
	var v559 int32
	_ = v559
	var v587 int32
	_ = v587
	var v588 int32
	_ = v588
	var v589 int32
	_ = v589
	var v590 int32
	_ = v590
	var v593 int32
	_ = v593
	var v624 int32
	_ = v624
	var v629 int32
	_ = v629
	var v632 int32
	_ = v632
	var v635 int32
	_ = v635
	var v636 int32
	_ = v636
	var v640 int32
	_ = v640
	var v646 int32
	_ = v646
	var v648 int32
	_ = v648
	var v656 int32
	_ = v656
	var v658 int32
	_ = v658
	var v659 int32
	_ = v659
	var v663 int32
	_ = v663
	var v669 int32
	_ = v669
	var v671 int32
	_ = v671
	var v677 int32
	_ = v677
	var v678 int32
	_ = v678
	var v679 int32
	_ = v679
	var v681 int32
	_ = v681
	var v682 int32
	_ = v682
	var v683 int32
	_ = v683
	var v684 int32
	_ = v684
	var v687 int32
	_ = v687
	var v695 int32
	_ = v695
	var v720 int32
	_ = v720
	var v721 int32
	_ = v721
	var v723 int32
	_ = v723
	var v727 int32
	_ = v727
	var v729 int32
	_ = v729
	var v733 int32
	_ = v733
	var v737 int32
	_ = v737
	var v743 int32
	_ = v743
	var v745 int32
	_ = v745
	var v751 int32
	_ = v751
	var v752 int32
	_ = v752
	var v753 int32
	_ = v753
	var v762 int32
	_ = v762
	var v765 int32
	_ = v765
	var v768 int32
	_ = v768
	var v770 int32
	_ = v770
	var v776 int32
	_ = v776
	var v778 int32
	_ = v778
	var v787 int32
	_ = v787
	var v796 int32
	_ = v796
	var v797 int32
	_ = v797
	var v801 int32
	_ = v801
	var v807 int32
	_ = v807
	var v813 int32
	_ = v813
	var v816 int32
	_ = v816
	var v819 int32
	_ = v819
	var v822 int32
	_ = v822
	var v824 int32
	_ = v824
	var v825 int32
	_ = v825
	var v826 int32
	_ = v826
	var v831 int32
	_ = v831
	var v837 int32
	_ = v837
	var v845 int32
	_ = v845
	var v847 int32
	_ = v847
	var v849 int32
	_ = v849
	var v850 int32
	_ = v850
	var v856 int32
	_ = v856
	var v857 int32
	_ = v857
	var v863 int32
	_ = v863
	var v868 int32
	_ = v868
	var v873 int32
	_ = v873
	var v875 int32
	_ = v875
	var v876 float64
	_ = v876
	var v878 float64
	_ = v878
	var v880 int32
	_ = v880
	var v881 int32
	_ = v881
	var v883 int32
	_ = v883
	var v888 int32
	_ = v888
	var v889 int32
	_ = v889
	var v893 int32
	_ = v893
	var v896 int32
	_ = v896
	var v897 int32
	_ = v897
	var v900 int32
	_ = v900
	var v905 int32
	_ = v905
	var v909 int32
	_ = v909
	var v913 int32
	_ = v913
	var v915 int32
	_ = v915
	var v919 int32
	_ = v919
	var v922 int64
	_ = v922
	var v923 int32
	_ = v923
	var v927 int32
	_ = v927
	var v933 int32
	_ = v933
	var v935 int32
	_ = v935
	var v941 int32
	_ = v941
	var v942 int64
	_ = v942
	var v952 int32
	_ = v952
	var v958 int32
	_ = v958
	var v960 int32
	_ = v960
	var v966 int32
	_ = v966
	var v974 int32
	_ = v974
	var v976 int32
	_ = v976
	var v982 int32
	_ = v982
	var v984 int32
	_ = v984
	var v987 int32
	_ = v987
	var v989 int32
	_ = v989
	var v993 int32
	_ = v993
	var v996 int32
	_ = v996
	var v998 int32
	_ = v998
	var v1028 int32
	_ = v1028
	var v1031 int32
	_ = v1031
	var v1035 int32
	_ = v1035
	var v1039 int32
	_ = v1039
	var v1041 int32
	_ = v1041
	var v1045 int32
	_ = v1045
	var v1047 int32
	_ = v1047
	var v1048 int32
	_ = v1048
	var v1051 float64
	_ = v1051
	var v1052 int32
	_ = v1052
	var v1055 int32
	_ = v1055
	var v1060 int32
	_ = v1060
	var v1061 int32
	_ = v1061
	var v1062 int32
	_ = v1062
	var v1063 int32
	_ = v1063
	var v1064 int32
	_ = v1064
	var v1068 int32
	_ = v1068
	var v1074 int32
	_ = v1074
	var v1077 int32
	_ = v1077
	var v1087 int32
	_ = v1087
	var v1091 int32
	_ = v1091
	var v1095 int32
	_ = v1095
	var v1098 int32
	_ = v1098
	var v1099 int32
	_ = v1099
	var v1100 int32
	_ = v1100
	var v1101 int32
	_ = v1101
	var v1105 int32
	_ = v1105
	var v1109 int32
	_ = v1109
	var v1113 int32
	_ = v1113
	var v1117 int32
	_ = v1117
	var v1122 int32
	_ = v1122
	var v1123 int32
	_ = v1123
	var v1129 int32
	_ = v1129
	var v1135 int32
	_ = v1135
	var v1137 int32
	_ = v1137
	var v1143 int32
	_ = v1143
	var v1144 int32
	_ = v1144
	var v1146 int32
	_ = v1146
	var v1149 int32
	_ = v1149
	var v1150 int32
	_ = v1150
	var v1151 int32
	_ = v1151
	var v1154 int32
	_ = v1154
	var v1157 int32
	_ = v1157
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
	var v1169 int32
	_ = v1169
	var v1170 int32
	_ = v1170
	var v1177 int32
	_ = v1177
	var v1179 int32
	_ = v1179
	var v1180 int32
	_ = v1180
	var v1181 int32
	_ = v1181
	var v1185 int32
	_ = v1185
	var v1191 int32
	_ = v1191
	var v1194 int32
	_ = v1194
	var v1204 int32
	_ = v1204
	var v1205 int32
	_ = v1205
	var v1208 int32
	_ = v1208
	var v1212 int32
	_ = v1212
	var v1213 int32
	_ = v1213
	var v1214 int32
	_ = v1214
	var v1215 int32
	_ = v1215
	var v1219 int32
	_ = v1219
	var v1225 int32
	_ = v1225
	var v1228 int32
	_ = v1228
	var v1238 int32
	_ = v1238
	var v1240 int32
	_ = v1240
	var v1241 int32
	_ = v1241
	var v1250 int32
	_ = v1250
	var v1251 int32
	_ = v1251
	var v1253 int32
	_ = v1253
	var v1264 int32
	_ = v1264
	var v1265 int32
	_ = v1265
	var v1272 int32
	_ = v1272
	var v1299 int32
	_ = v1299
	var v1305 int32
	_ = v1305
	var v1311 int32
	_ = v1311
	var v1314 int32
	_ = v1314
	var v1317 int32
	_ = v1317
	var v1322 int32
	_ = v1322
	var v1323 int32
	_ = v1323
	var v1327 int32
	_ = v1327
	var v1330 int32
	_ = v1330
	var v1331 int32
	_ = v1331
	var v1337 int32
	_ = v1337
	var v1340 int32
	_ = v1340
	var v1343 int32
	_ = v1343
	var v1346 int32
	_ = v1346
	var v1348 int32
	_ = v1348
	var v1349 int32
	_ = v1349
	var v1352 int32
	_ = v1352
	var v1354 int64
	_ = v1354
	var v1358 int32
	_ = v1358
	var v1359 int32
	_ = v1359
	var v1362 int32
	_ = v1362
	var v1363 int32
	_ = v1363
	var v1364 int32
	_ = v1364
	var v1366 int32
	_ = v1366
	var v1368 int32
	_ = v1368
	var v1370 int32
	_ = v1370
	var v1374 int32
	_ = v1374
	var v1376 int32
	_ = v1376
	var v1377 int32
	_ = v1377
	var v1382 int32
	_ = v1382
	var v1387 int32
	_ = v1387
	var v1388 int32
	_ = v1388
	var v1389 int32
	_ = v1389
	var v1390 int32
	_ = v1390
	var v1394 int32
	_ = v1394
	var v1396 int32
	_ = v1396
	var v1397 int32
	_ = v1397
	var v1399 int32
	_ = v1399
	var v1405 int32
	_ = v1405
	var v1406 int32
	_ = v1406
	var v1407 int32
	_ = v1407
	var v1413 int32
	_ = v1413
	var v1414 int32
	_ = v1414
	var v1417 int32
	_ = v1417
	var v1423 int32
	_ = v1423
	var v1427 int32
	_ = v1427
	var v1432 int32
	_ = v1432
	var v1433 int32
	_ = v1433
	var v1434 int32
	_ = v1434
	var v1438 int32
	_ = v1438
	var v1444 int32
	_ = v1444
	var v1446 int32
	_ = v1446
	var v1452 int32
	_ = v1452
	var v1453 int32
	_ = v1453
	var v1454 int32
	_ = v1454
	var v1456 int32
	_ = v1456
	var v1458 int32
	_ = v1458
	var v1461 int32
	_ = v1461
	var v1465 int32
	_ = v1465
	var v1471 int32
	_ = v1471
	var v1473 int32
	_ = v1473
	var v1479 int32
	_ = v1479
	var v1480 int32
	_ = v1480
	var v1481 int32
	_ = v1481
	var v1489 int32
	_ = v1489
	var v1490 int32
	_ = v1490
	var v1491 int32
	_ = v1491
	var v1495 int32
	_ = v1495
	var v1498 int32
	_ = v1498
	var v1499 int32
	_ = v1499
	var v1501 int32
	_ = v1501
	var v1503 int32
	_ = v1503
	var v1504 int32
	_ = v1504
	var v1508 int32
	_ = v1508
	var v1512 int32
	_ = v1512
	var v1516 int32
	_ = v1516
	var v1520 int32
	_ = v1520
	var v1523 int32
	_ = v1523
	var v1528 int32
	_ = v1528
	var v1532 int32
	_ = v1532
	var v1535 int32
	_ = v1535
	var v1536 int32
	_ = v1536
	var v1537 int32
	_ = v1537
	var v1542 int32
	_ = v1542
	var v1543 int32
	_ = v1543
	var v1544 int32
	_ = v1544
	var v1550 int32
	_ = v1550
	var v1555 int32
	_ = v1555
	var v1558 int64
	_ = v1558
	var v1559 int32
	_ = v1559
	var v1563 int32
	_ = v1563
	var v1569 int32
	_ = v1569
	var v1571 int32
	_ = v1571
	var v1577 int32
	_ = v1577
	var v1578 int64
	_ = v1578
	var v1583 int32
	_ = v1583
	var v1584 int32
	_ = v1584
	var v1588 int32
	_ = v1588
	var v1594 int32
	_ = v1594
	var v1596 int32
	_ = v1596
	var v1602 int32
	_ = v1602
	var v1608 int32
	_ = v1608
	var v1612 int32
	_ = v1612
	var v1614 int32
	_ = v1614
	var v1618 int32
	_ = v1618
	var v1625 int32
	_ = v1625
	var v1627 int32
	_ = v1627
	var v1633 int32
	_ = v1633
	var v1636 int32
	_ = v1636
	var v1638 int32
	_ = v1638
	var v1640 int32
	_ = v1640
	var v1643 int32
	_ = v1643
	var v1657 int32
	_ = v1657
	var v1703 int32
	_ = v1703
	var v1710 int32
	_ = v1710
	var v1713 int32
	_ = v1713
	var v1714 int32
	_ = v1714
	var v1725 int32
	_ = v1725
	var v1729 int32
	_ = v1729
	var v1734 int32
	_ = v1734
	v5 = int32(0)
	v28 = m.G0
	v30 = v28 - int32(4128)
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
	v48 = (v42&int32(8191) + int32(7)) & int32(16376)
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
	v1710 = m.ExcPending
	if v1710 != 0 {
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
	v84 = *(*int32)(unsafe.Add(mBase, _consts[5]))
	v90 = *(*int32)(unsafe.Add(mBase, uint32(v84+(v79^int32(-1))<<(uint(int32(2))%32))))
	v98 = v90
	goto L8
L12:
	;
	goto L13
L13:
	;
	v92 = *(*int32)(unsafe.Add(mBase, _consts[6]))
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
	v116 = *(*int32)(unsafe.Add(mBase, _consts[8]))
	v122 = *(*int32)(unsafe.Add(mBase, uint32(v116+(v111^int32(-1))<<(uint(int32(6))%32))+16))
	v131 = v122
	goto L16
L18:
	;
	goto L19
L19:
	;
	v124 = *(*int32)(unsafe.Add(mBase, _consts[9]))
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
	v137 = *(*int32)(unsafe.Add(mBase, _consts[5]))
	v143 = *(*int32)(unsafe.Add(mBase, uint32(v137+(v111^int32(-1))<<(uint(int32(2))%32))))
	v151 = v143
	goto L22
L24:
	;
	goto L25
L25:
	;
	v145 = *(*int32)(unsafe.Add(mBase, _consts[6]))
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
	v720 = m.ExcPending
	if v720 != 0 {
		goto L9
	} else {
		goto L144
	}
L34:
	;
	if base.Ui32(v48) <= base.Ui32(v184-int32(4)) {
		v695 = v111
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
	v203 = v151
	v207 = v153
	goto L39
L39:
	;
	v219 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v207)+12)))
	if v219&int32(128) == int32(0) {
		goto L41
	} else {
		goto L42
	}
L40:
	;
	v695 = v677
	goto L33
L41:
	;
	v624 = *(*int32)(unsafe.Add(mBase, uint32(v207)+4))
	if v624 != int32(-1) {
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
	v587 = int32(4)
	v588 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v203)+14)))
	v589 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v203)+12)))
	v590 = v588 - v589
	if v590 <= v587 {
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
	v233 = *(*int32)(unsafe.Add(mBase, _consts[5]))
	v239 = *(*int32)(unsafe.Add(mBase, uint32(v233+(v196^int32(-1))<<(uint(int32(2))%32))))
	v247 = v239
	goto L46
L48:
	;
	goto L49
L49:
	;
	v241 = *(*int32)(unsafe.Add(mBase, _consts[6]))
	v247 = v241 + v196<<(uint(int32(13))%32) + int32(-8192)
	goto L46
L50:
	;
	v254 = int32(base.Ui32(v248+int32(262120)) >> (uint(int32(2)) % 32))
	if v254&int32(65535) == int32(0) {
		goto L45
	} else {
		goto L51
	}
L51:
	;
	v259 = int32(1)
	v260 = int32(2)
	v264 = (v254 + v259) & int32(65535)
	if base.Ui32(v264) <= base.Ui32(v260) {
		goto L52
	} else {
		goto L53
	}
L52:
	;
	v267 = v260
	goto L54
L53:
	;
	v267 = v264
	goto L54
L54:
	;
	v268 = int32(1)
	v269 = v267 - v268
	v273 = v247 + int32(24)
	v274 = int32(0)
	if base.Ui32(int32(3)) <= base.Ui32(v264) {
		goto L55
	} else {
		goto L56
	}
L55:
	;
	v285 = v259
	v287 = v274
	v292 = int32(0)
	goto L58
L56:
	;
	v352 = v259
	v354 = v274
	goto L57
L57:
	;
	if v269&v268 == int32(0) {
		v394 = v354
		goto L67
	} else {
		goto L68
	}
L58:
	;
	v309 = v285<<(uint(int32(2))%32) + v273
	v312 = *(*int32)(unsafe.Add(mBase, uint32(v309-int32(4))))
	v313 = int32(98304)
	if v312&v313 == v313 {
		goto L60
	} else {
		goto L61
	}
L59:
	;
	v352 = v343
	v354 = v341
	goto L57
L60:
	;
	v319 = int32(1)
	*(*uint16)(unsafe.Add(mBase, uint32(v30+int32(32)+v287<<(uint(v319)%32)))) = uint16(v285)
	v325 = v287 + v319
	goto L62
L61:
	;
	v325 = v287
	goto L62
L62:
	;
	v326 = *(*int32)(unsafe.Add(mBase, uint32(v309)))
	v327 = int32(98304)
	if v326&v327 == v327 {
		goto L63
	} else {
		goto L64
	}
L63:
	;
	v333 = int32(1)
	v337 = v285 + v333
	*(*uint16)(unsafe.Add(mBase, uint32(v30+int32(32)+v325<<(uint(v333)%32)))) = uint16(v337)
	v341 = v325 + v333
	goto L65
L64:
	;
	v341 = v325
	goto L65
L65:
	;
	v342 = int32(2)
	v343 = v285 + v342
	v345 = v292 + v342
	if v345 != v269&int32(-2) {
		v285 = v343
		v287 = v341
		v292 = v345
		goto L58
	} else {
		goto L66
	}
L66:
	;
	goto L59
L67:
	;
	if v394 <= int32(0) {
		goto L45
	} else {
		goto L70
	}
L68:
	;
	v381 = *(*int32)(unsafe.Add(mBase, uint32(v352<<(uint(int32(2))%32)+v273-int32(4))))
	v382 = int32(98304)
	if v381&v382 != v382 {
		v394 = v354
		goto L67
	} else {
		goto L69
	}
L69:
	;
	v388 = int32(1)
	*(*uint16)(unsafe.Add(mBase, uint32(v30+int32(32)+v354<<(uint(v388)%32)))) = uint16(v352)
	v394 = v354 + v388
	goto L67
L70:
	;
	v399 = F_index_compute_xid_horizon_for_tuples(m, l0, l2, v196, v30+int32(32), v394)
	mBase = m.M
	v400 = m.ExcPending
	if v400 != 0 {
		goto L9
	} else {
		goto L71
	}
L71:
	;
	F_LockBuffer(m, v79, int32(2))
	mBase = m.M
	v403 = m.ExcPending
	if v403 != 0 {
		goto L9
	} else {
		goto L72
	}
L72:
	;
	v404 = int32(4543684)
	v406 = *(*int32)(unsafe.Add(mBase, _consts[7]))
	*(*int32)(unsafe.Add(mBase, _consts[7])) = v406 + int32(1)
	F_PageIndexMultiDelete(m, v247, v30+int32(32), v394)
	mBase = m.M
	v413 = m.ExcPending
	if v413 != 0 {
		goto L9
	} else {
		goto L73
	}
L73:
	;
	v414 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v247)+16)))
	v415 = v247 + v414
	v416 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v415)+12)))
	v418 = v416 & int32(65407)
	*(*uint16)(unsafe.Add(mBase, uint32(v415)+12)) = uint16(v418)
	v420 = int32(0)
	v421 = base.B2i32(v420 <= v79)
	if v421 == v420 {
		goto L75
	} else {
		goto L76
	}
L74:
	;
	v437 = v435 + int32(32)
	v438 = *(*float64)(unsafe.Add(mBase, uint32(v437)))
	*(*float64)(unsafe.Add(mBase, uint32(v437))) = base.F64_sub(v438, base.F64_convert_i32_u(v394))
	F_MarkBufferDirty(m, v196)
	mBase = m.M
	v443 = m.ExcPending
	if v443 != 0 {
		goto L9
	} else {
		goto L78
	}
L75:
	;
	v425 = *(*int32)(unsafe.Add(mBase, _consts[5]))
	v429 = *(*int32)(unsafe.Add(mBase, uint32(v425+v189<<(uint(int32(2))%32))))
	v435 = v429
	goto L74
L76:
	;
	goto L77
L77:
	;
	v431 = *(*int32)(unsafe.Add(mBase, _consts[6]))
	v435 = v431 + v191 + int32(-8192)
	goto L74
L78:
	;
	F_MarkBufferDirty(m, v79)
	mBase = m.M
	v445 = m.ExcPending
	if v445 != 0 {
		goto L9
	} else {
		goto L79
	}
L79:
	;
	v446 = *(*int32)(unsafe.Add(mBase, uint32(l0)+48))
	v447 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v446)+118)))
	if v447 != int32(112) {
		goto L80
	} else {
		goto L81
	}
L80:
	;
	v551 = int32(4543684)
	v553 = *(*int32)(unsafe.Add(mBase, _consts[7]))
	*(*int32)(unsafe.Add(mBase, _consts[7])) = v553 - int32(1)
	F_LockBuffer(m, v79, int32(0))
	mBase = m.M
	v559 = m.ExcPending
	if v559 != 0 {
		goto L9
	} else {
		goto L114
	}
L81:
	;
	v451 = *(*int32)(unsafe.Add(mBase, _consts[10]))
	if v451 <= int32(0) {
		goto L83
	} else {
		goto L84
	}
L82:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v30)+24)) = v399
	*(*uint8)(unsafe.Add(mBase, uint32(v30)+30)) = uint8(v479)
	*(*uint16)(unsafe.Add(mBase, uint32(v30)+28)) = uint16(v394)
	F_XLogBeginInsert(m)
	mBase = m.M
	v485 = m.ExcPending
	if v485 != 0 {
		goto L9
	} else {
		goto L100
	}
L83:
	;
	v454 = *(*int32)(unsafe.Add(mBase, uint32(l0)+32))
	if v454 != 0 {
		goto L80
	} else {
		goto L86
	}
L84:
	;
	goto L85
L85:
	;
	if v451 == int32(1) {
		goto L88
	} else {
		goto L89
	}
L86:
	;
	v455 = *(*int32)(unsafe.Add(mBase, uint32(l0)+40))
	if v455 != 0 {
		goto L80
	} else {
		goto L87
	}
L87:
	;
	v479 = int32(0)
	goto L82
L88:
	;
	v479 = int32(0)
	goto L82
L89:
	;
	goto L90
L90:
	;
	v460 = *(*int32)(unsafe.Add(mBase, uint32(l2)+48))
	v461 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v460)+118)))
	if v461 != int32(112) {
		goto L91
	} else {
		goto L92
	}
L91:
	;
	v479 = int32(0)
	goto L82
L92:
	;
	goto L93
L93:
	;
	v466 = *(*int32)(unsafe.Add(mBase, uint32(l2)+56))
	goto L94
L94:
	;
	if base.Ui32(v466) < base.Ui32(int32(12000)) {
		v479 = int32(1)
		goto L82
	} else {
		goto L95
	}
L95:
	;
	v469 = *(*int32)(unsafe.Add(mBase, uint32(l2)+180))
	if v469 == int32(0) {
		goto L96
	} else {
		goto L97
	}
L96:
	;
	v479 = int32(0)
	goto L82
L97:
	;
	goto L98
L98:
	;
	v474 = *(*int32)(unsafe.Add(mBase, uint32(l2)+48))
	v475 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v474)+119)))
	switch v475 - int32(109) {
	case 0, 5:
		goto L99
	default:
		v479 = int32(0)
		goto L82
	}
L99:
	;
	v478 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v469)+104)))
	v479 = v478
	goto L82
L100:
	;
	F_XLogRegisterBuffer(m, int32(0), v196, int32(8))
	mBase = m.M
	v489 = m.ExcPending
	if v489 != 0 {
		goto L9
	} else {
		goto L101
	}
L101:
	;
	F_XLogRegisterData(m, v30+int32(24), int32(8))
	mBase = m.M
	v494 = m.ExcPending
	if v494 != 0 {
		goto L9
	} else {
		goto L102
	}
L102:
	;
	F_XLogRegisterData(m, v30+int32(32), v394<<(uint(int32(1))%32))
	mBase = m.M
	v500 = m.ExcPending
	if v500 != 0 {
		goto L9
	} else {
		goto L103
	}
L103:
	;
	F_XLogRegisterBuffer(m, int32(1), v79, int32(8))
	mBase = m.M
	v504 = m.ExcPending
	if v504 != 0 {
		goto L9
	} else {
		goto L104
	}
L104:
	;
	v507 = F_XLogInsert(m, int32(12), int32(192))
	mBase = m.M
	v508 = m.ExcPending
	if v508 != 0 {
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
	*(*int64)(unsafe.Add(mBase, uint32(v526))) = base.I64_rotr(v507, int64(32))
	if v421 == int32(0) {
		goto L111
	} else {
		goto L112
	}
L107:
	;
	v512 = *(*int32)(unsafe.Add(mBase, _consts[5]))
	v518 = *(*int32)(unsafe.Add(mBase, uint32(v512+(v196^int32(-1))<<(uint(int32(2))%32))))
	v526 = v518
	goto L106
L108:
	;
	goto L109
L109:
	;
	v520 = *(*int32)(unsafe.Add(mBase, _consts[6]))
	v526 = v520 + v196<<(uint(int32(13))%32) + int32(-8192)
	goto L106
L110:
	;
	*(*uint32)(unsafe.Add(mBase, uint32(v543)+4)) = uint32(v507)
	v546 = int64(base.Ui64(v507) >> (uint(int64(32)) % 64))
	*(*uint32)(unsafe.Add(mBase, uint32(v543))) = uint32(v546)
	goto L80
L111:
	;
	v533 = *(*int32)(unsafe.Add(mBase, _consts[5]))
	v537 = *(*int32)(unsafe.Add(mBase, uint32(v533+v189<<(uint(int32(2))%32))))
	v543 = v537
	goto L110
L112:
	;
	goto L113
L113:
	;
	v539 = *(*int32)(unsafe.Add(mBase, _consts[6]))
	v543 = v539 + v191 + int32(-8192)
	goto L110
L114:
	;
	goto L45
L115:
	;
	if base.Ui32(v48) <= base.Ui32(v593-int32(4)) {
		v695 = v196
		goto L33
	} else {
		goto L119
	}
L116:
	;
	v593 = v587
	goto L118
L117:
	;
	v593 = v590
	goto L118
L118:
	;
	goto L115
L119:
	;
	goto L41
L120:
	;
	v679 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v678)+16)))
	v681 = int32(4)
	v682 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v678)+14)))
	v683 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v678)+12)))
	v684 = v682 - v683
	if v684 <= v681 {
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
	v656 = m.ExcPending
	if v656 != 0 {
		goto L9
	} else {
		goto L134
	}
L124:
	;
	v635 = F__hash_getbuf(m, l0, v624, int32(2), int32(1))
	mBase = m.M
	v636 = m.ExcPending
	if v636 != 0 {
		goto L9
	} else {
		goto L130
	}
L125:
	;
	F_UnlockReleaseBuffer(m, v196)
	mBase = m.M
	v629 = m.ExcPending
	if v629 != 0 {
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
	v632 = m.ExcPending
	if v632 != 0 {
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
	if v635 < int32(0) {
		goto L131
	} else {
		goto L132
	}
L131:
	;
	v640 = *(*int32)(unsafe.Add(mBase, _consts[5]))
	v646 = *(*int32)(unsafe.Add(mBase, uint32(v640+(v635^int32(-1))<<(uint(int32(2))%32))))
	v677 = v635
	v678 = v646
	goto L120
L132:
	;
	goto L133
L133:
	;
	v648 = *(*int32)(unsafe.Add(mBase, _consts[6]))
	v677 = v635
	v678 = v648 + v635<<(uint(int32(13))%32) + int32(-8192)
	goto L120
L134:
	;
	v658 = F__hash_addovflpage(m, l0, v79, v196, base.B2i32(v196 == v111))
	mBase = m.M
	v659 = m.ExcPending
	if v659 != 0 {
		goto L9
	} else {
		goto L135
	}
L135:
	;
	if v658 < int32(0) {
		goto L136
	} else {
		goto L137
	}
L136:
	;
	v663 = *(*int32)(unsafe.Add(mBase, _consts[5]))
	v669 = *(*int32)(unsafe.Add(mBase, uint32(v663+(v658^int32(-1))<<(uint(int32(2))%32))))
	v677 = v658
	v678 = v669
	goto L120
L137:
	;
	goto L138
L138:
	;
	v671 = *(*int32)(unsafe.Add(mBase, _consts[6]))
	v677 = v658
	v678 = v671 + v658<<(uint(int32(13))%32) + int32(-8192)
	goto L120
L139:
	;
	if base.Ui32(v687-int32(4)) < base.Ui32(v48) {
		v196 = v677
		v203 = v678
		v207 = v679 + v678
		goto L39
	} else {
		goto L143
	}
L140:
	;
	v687 = v681
	goto L142
L141:
	;
	v687 = v684
	goto L142
L142:
	;
	goto L139
L143:
	;
	goto L40
L144:
	;
	v721 = int32(4543684)
	v723 = *(*int32)(unsafe.Add(mBase, _consts[7]))
	*(*int32)(unsafe.Add(mBase, _consts[7])) = v723 + int32(1)
	v727 = m.G0
	v729 = v727 - int32(16)
	m.G0 = v729
	F__hash_checkpage(m, l0, v695, int32(3))
	mBase = m.M
	v733 = m.ExcPending
	if v733 != 0 {
		goto L9
	} else {
		goto L145
	}
L145:
	;
	if v695 < int32(0) {
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
	v737 = *(*int32)(unsafe.Add(mBase, _consts[5]))
	v743 = *(*int32)(unsafe.Add(mBase, uint32(v737+(v695^int32(-1))<<(uint(int32(2))%32))))
	v751 = v743
	goto L146
L148:
	;
	goto L149
L149:
	;
	v745 = *(*int32)(unsafe.Add(mBase, _consts[6]))
	v751 = v745 + v695<<(uint(int32(13))%32) + int32(-8192)
	goto L146
L150:
	;
	v847 = v845 & int32(65535)
	v849 = F_PageAddItemExtended(m, v751, l1, v48, v847, int32(0))
	mBase = m.M
	v850 = m.ExcPending
	if v850 != 0 {
		goto L9
	} else {
		goto L180
	}
L151:
	;
	v752 = int32(1)
	v753 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v751)+12)))
	if base.Ui32(v753) < base.Ui32(int32(25)) {
		goto L154
	} else {
		goto L155
	}
L152:
	;
	goto L153
L153:
	;
	v765 = int32(*(*int16)(unsafe.Add(mBase, uint32(l1)+6)))
	if int32(0) <= v765 {
		goto L158
	} else {
		goto L159
	}
L154:
	;
	v762 = v752
	goto L156
L155:
	;
	v762 = int32(base.Ui32(v753+int32(262120))>>(uint(int32(2))%32)) + v752
	goto L156
L156:
	;
	v845 = v762
	goto L150
L157:
	;
	v776 = int32(1)
	v778 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v751)+12)))
	if base.Ui32(v778) < base.Ui32(int32(25)) {
		goto L162
	} else {
		goto L163
	}
L158:
	;
	v768 = int32(8)
	goto L160
L159:
	;
	v768 = int32(16)
	goto L160
L160:
	;
	v770 = *(*int32)(unsafe.Add(mBase, uint32(l1+v768)))
	goto L157
L161:
	;
	v845 = v837 & int32(65535)
	goto L150
L162:
	;
	v787 = v776
	goto L164
L163:
	;
	v787 = int32(base.Ui32(v778+int32(262120))>>(uint(int32(2))%32)) + v776
	goto L164
L164:
	;
	if base.Ui32(int32(2)) <= base.Ui32(v787&int32(65535)) {
		goto L165
	} else {
		goto L166
	}
L165:
	;
	v796 = v776
	v797 = v787
	goto L168
L166:
	;
	v837 = v776
	goto L167
L167:
	;
	goto L161
L168:
	;
	v801 = int32(65535)
	v807 = int32(base.Ui32(v797&v801+v796&v801) >> (uint(int32(1)) % 32))
	v813 = *(*int32)(unsafe.Add(mBase, uint32(v807<<(uint(int32(2))%32)+(v751+int32(24))-int32(4))))
	v816 = v751 + v813&int32(32767)
	v819 = int32(*(*int16)(unsafe.Add(mBase, uint32(v816)+6)))
	if int32(0) <= v819 {
		goto L170
	} else {
		goto L171
	}
L169:
	;
	v837 = v831
	goto L167
L170:
	;
	v822 = int32(8)
	goto L172
L171:
	;
	v822 = int32(16)
	goto L172
L172:
	;
	v824 = *(*int32)(unsafe.Add(mBase, uint32(v816+v822)))
	v825 = base.B2i32(base.Ui32(v824) < base.Ui32(v770))
	if base.Ui32(v824) < base.Ui32(v770) {
		goto L173
	} else {
		goto L174
	}
L173:
	;
	v826 = v797
	goto L175
L174:
	;
	v826 = v807
	goto L175
L175:
	;
	if base.Ui32(v824) < base.Ui32(v770) {
		goto L176
	} else {
		goto L177
	}
L176:
	;
	v831 = v807 + int32(1)
	goto L178
L177:
	;
	v831 = v796
	goto L178
L178:
	;
	if base.Ui32(v831&int32(65535)) < base.Ui32(v826&int32(65535)) {
		v796 = v831
		v797 = v826
		goto L168
	} else {
		goto L179
	}
L179:
	;
	goto L169
L180:
	;
	if v849 == int32(0) {
		goto L181
	} else {
		goto L182
	}
L181:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v856 = m.ExcPending
	if v856 != 0 {
		goto L9
	} else {
		goto L184
	}
L182:
	;
	goto L183
L183:
	;
	m.G0 = v729 + int32(16)
	F_MarkBufferDirty(m, v695)
	mBase = m.M
	v873 = m.ExcPending
	if v873 != 0 {
		goto L9
	} else {
		goto L187
	}
L184:
	;
	v857 = *(*int32)(unsafe.Add(mBase, uint32(l0)+48))
	*(*int32)(unsafe.Add(mBase, uint32(v729))) = v857 + int32(4)
	F_errmsg_internal(m, int32(730156), v729)
	mBase = m.M
	v863 = m.ExcPending
	if v863 != 0 {
		goto L9
	} else {
		goto L185
	}
L185:
	;
	F_errfinish(m, int32(512279), int32(316), int32(242376))
	mBase = m.M
	v868 = m.ExcPending
	if v868 != 0 {
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
	v875 = v98 + int32(32)
	v876 = *(*float64)(unsafe.Add(mBase, uint32(v875)))
	v878 = base.F64_add(v876, float64(1))
	*(*float64)(unsafe.Add(mBase, uint32(v875))) = v878
	v880 = *(*int32)(unsafe.Add(mBase, uint32(v98)+48))
	v881 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v98)+40)))
	F_MarkBufferDirty(m, v79)
	mBase = m.M
	v883 = m.ExcPending
	if v883 != 0 {
		goto L9
	} else {
		goto L188
	}
L188:
	;
	v888 = *(*int32)(unsafe.Add(mBase, uint32(l0)+48))
	v889 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v888)+118)))
	if v889 != int32(112) {
		goto L189
	} else {
		goto L190
	}
L189:
	;
	v974 = int32(4543684)
	v976 = *(*int32)(unsafe.Add(mBase, _consts[7]))
	*(*int32)(unsafe.Add(mBase, _consts[7])) = v976 - int32(1)
	F_LockBuffer(m, v79, int32(0))
	mBase = m.M
	v982 = m.ExcPending
	if v982 != 0 {
		goto L9
	} else {
		goto L210
	}
L190:
	;
	v893 = *(*int32)(unsafe.Add(mBase, _consts[10]))
	if v893 <= int32(0) {
		goto L191
	} else {
		goto L192
	}
L191:
	;
	v896 = *(*int32)(unsafe.Add(mBase, uint32(l0)+32))
	if v896 != 0 {
		goto L189
	} else {
		goto L194
	}
L192:
	;
	goto L193
L193:
	;
	*(*uint16)(unsafe.Add(mBase, uint32(v30)+32)) = uint16(v847)
	F_XLogBeginInsert(m)
	mBase = m.M
	v900 = m.ExcPending
	if v900 != 0 {
		goto L9
	} else {
		goto L196
	}
L194:
	;
	v897 = *(*int32)(unsafe.Add(mBase, uint32(l0)+40))
	if v897 != 0 {
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
	v905 = m.ExcPending
	if v905 != 0 {
		goto L9
	} else {
		goto L197
	}
L197:
	;
	F_XLogRegisterBuffer(m, int32(1), v79, int32(8))
	mBase = m.M
	v909 = m.ExcPending
	if v909 != 0 {
		goto L9
	} else {
		goto L198
	}
L198:
	;
	F_XLogRegisterBuffer(m, int32(0), v695, int32(8))
	mBase = m.M
	v913 = m.ExcPending
	if v913 != 0 {
		goto L9
	} else {
		goto L199
	}
L199:
	;
	v915 = int32(*(*uint16)(unsafe.Add(mBase, uint32(l1)+6)))
	F_XLogRegisterBufData(m, int32(0), l1, v915&int32(8191))
	mBase = m.M
	v919 = m.ExcPending
	if v919 != 0 {
		goto L9
	} else {
		goto L200
	}
L200:
	;
	v922 = F_XLogInsert(m, int32(12), int32(32))
	mBase = m.M
	v923 = m.ExcPending
	if v923 != 0 {
		goto L9
	} else {
		goto L201
	}
L201:
	;
	if v695 < int32(0) {
		goto L203
	} else {
		goto L204
	}
L202:
	;
	v942 = int64(32)
	*(*int64)(unsafe.Add(mBase, uint32(v941))) = base.I64_rotr(v922, v942)
	if v79 < int32(0) {
		goto L207
	} else {
		goto L208
	}
L203:
	;
	v927 = *(*int32)(unsafe.Add(mBase, _consts[5]))
	v933 = *(*int32)(unsafe.Add(mBase, uint32(v927+(v695^int32(-1))<<(uint(int32(2))%32))))
	v941 = v933
	goto L202
L204:
	;
	goto L205
L205:
	;
	v935 = *(*int32)(unsafe.Add(mBase, _consts[6]))
	v941 = v935 + v695<<(uint(int32(13))%32) + int32(-8192)
	goto L202
L206:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v966)+4)) = base.I32_wrap_i64(v922)
	*(*int32)(unsafe.Add(mBase, uint32(v966))) = base.I32_wrap_i64(int64(base.Ui64(v922) >> (uint(v942) % 64)))
	goto L189
L207:
	;
	v952 = *(*int32)(unsafe.Add(mBase, _consts[5]))
	v958 = *(*int32)(unsafe.Add(mBase, uint32(v952+(v79^int32(-1))<<(uint(int32(2))%32))))
	v966 = v958
	goto L206
L208:
	;
	goto L209
L209:
	;
	v960 = *(*int32)(unsafe.Add(mBase, _consts[6]))
	v966 = v960 + v79<<(uint(int32(13))%32) + int32(-8192)
	goto L206
L210:
	;
	F_UnlockReleaseBuffer(m, v695)
	mBase = m.M
	v984 = m.ExcPending
	if v984 != 0 {
		goto L9
	} else {
		goto L211
	}
L211:
	;
	if v695 != v111 {
		goto L212
	} else {
		goto L213
	}
L212:
	;
	F_ReleaseBuffer(m, v111)
	mBase = m.M
	v987 = m.ExcPending
	if v987 != 0 {
		goto L9
	} else {
		goto L215
	}
L213:
	;
	goto L214
L214:
	;
	if base.F64_lt(base.F64_mul(base.F64_convert_i32_u(v881), base.F64_convert_i32_u(v880+int32(1))), v878) != 0 {
		goto L216
	} else {
		goto L217
	}
L215:
	;
	goto L214
L216:
	;
	v989 = m.G0
	v993 = (v989 - int32(12288)) & int32(-4096)
	m.G0 = v993
	v996 = v79 ^ int32(-1)
	v998 = v79 << (uint(int32(13)) % 32)
	goto L221
L217:
	;
	goto L218
L218:
	;
	F_ReleaseBuffer(m, v79)
	mBase = m.M
	v1703 = m.ExcPending
	if v1703 != 0 {
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
	v1657 = m.ExcPending
	if v1657 != 0 {
		goto L9
	} else {
		goto L392
	}
L221:
	;
	F_LockBuffer(m, v79, int32(2))
	mBase = m.M
	v1028 = m.ExcPending
	if v1028 != 0 {
		goto L9
	} else {
		goto L223
	}
L222:
	;
	v1180 = int32(2)
	v1181 = v1048 + v1180
	v1185 = v1181 - int32(1)
	if base.Ui32(v1180) <= base.Ui32(v1181) {
		goto L274
	} else {
		goto L275
	}
L223:
	;
	F__hash_checkpage(m, l0, v79, int32(8))
	mBase = m.M
	v1031 = m.ExcPending
	if v1031 != 0 {
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
	v1047 = v1045 + int32(48)
	v1048 = *(*int32)(unsafe.Add(mBase, uint32(v1047)))
	if base.Ui32(int32(2147483645)) < base.Ui32(v1048) {
		goto L220
	} else {
		goto L229
	}
L226:
	;
	v1035 = *(*int32)(unsafe.Add(mBase, _consts[5]))
	v1039 = *(*int32)(unsafe.Add(mBase, uint32(v1035+v996<<(uint(int32(2))%32))))
	v1045 = v1039
	goto L225
L227:
	;
	goto L228
L228:
	;
	v1041 = *(*int32)(unsafe.Add(mBase, _consts[6]))
	v1045 = v1041 + v998 + int32(-8192)
	goto L225
L229:
	;
	v1051 = *(*float64)(unsafe.Add(mBase, uint32(v1045)+32))
	v1052 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v1045)+40)))
	v1055 = v1048 + int32(1)
	if base.F64_le(v1051, base.F64_mul(base.F64_convert_i32_u(v1052), base.F64_convert_i32_u(v1055))) != 0 {
		goto L220
	} else {
		goto L230
	}
L230:
	;
	v1060 = v1045 + int32(56)
	v1061 = *(*int32)(unsafe.Add(mBase, uint32(v1060)))
	v1062 = v1061 & v1055
	if v1062 != 0 {
		goto L232
	} else {
		goto L233
	}
L231:
	;
	if v1123 == int32(0) {
		goto L220
	} else {
		goto L255
	}
L232:
	;
	v1063 = int32(1)
	v1064 = v1062 + v1063
	v1068 = v1064 - v1063
	if base.Ui32(int32(2)) <= base.Ui32(v1064) {
		goto L236
	} else {
		goto L237
	}
L233:
	;
	v1095 = int32(1)
	goto L234
L234:
	;
	if v1095 != int32(-1) {
		goto L242
	} else {
		goto L243
	}
L235:
	;
	v1091 = *(*int32)(unsafe.Add(mBase, uint32(v1087<<(uint(int32(2))%32)+v1045)+72))
	v1095 = v1091 + v1064
	goto L234
L236:
	;
	v1074 = int32(32) - base.I32_clz(v1068)
	goto L238
L237:
	;
	v1074 = int32(0)
	goto L238
L238:
	;
	if base.Ui32(int32(10)) <= base.Ui32(v1074) {
		goto L239
	} else {
		goto L240
	}
L239:
	;
	v1077 = int32(3)
	v1087 = int32(base.Ui32(v1068)>>(uint(v1074-v1077)%32))&v1077 | v1074<<(uint(int32(2))%32) - int32(30)
	goto L241
L240:
	;
	v1087 = v1074
	goto L241
L241:
	;
	goto L235
L242:
	;
	v1098 = F_ReadBuffer(m, l0, v1095)
	mBase = m.M
	v1099 = m.ExcPending
	if v1099 != 0 {
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
	v1113 = m.ExcPending
	if v1113 != 0 {
		goto L9
	} else {
		goto L252
	}
L245:
	;
	v1100 = F_ConditionalLockBufferForCleanup(m, v1098)
	mBase = m.M
	v1101 = m.ExcPending
	if v1101 != 0 {
		goto L9
	} else {
		goto L246
	}
L246:
	;
	if v1100 == int32(0) {
		goto L247
	} else {
		goto L248
	}
L247:
	;
	F_ReleaseBuffer(m, v1098)
	mBase = m.M
	v1105 = m.ExcPending
	if v1105 != 0 {
		goto L9
	} else {
		goto L250
	}
L248:
	;
	goto L249
L249:
	;
	F__hash_checkpage(m, l0, v1098, int32(2))
	mBase = m.M
	v1109 = m.ExcPending
	if v1109 != 0 {
		goto L9
	} else {
		goto L251
	}
L250:
	;
	v1123 = int32(0)
	goto L231
L251:
	;
	v1123 = v1098
	goto L231
L252:
	;
	F_errmsg_internal(m, int32(537435), int32(0))
	mBase = m.M
	v1117 = m.ExcPending
	if v1117 != 0 {
		goto L9
	} else {
		goto L253
	}
L253:
	;
	F_errfinish(m, int32(519802), int32(101), int32(242617))
	mBase = m.M
	v1122 = m.ExcPending
	if v1122 != 0 {
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
	if v1123 < int32(0) {
		goto L257
	} else {
		goto L258
	}
L256:
	;
	v1144 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v1143)+16)))
	v1146 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v1144+v1143)+12)))
	if v1146&int32(32) != 0 {
		goto L260
	} else {
		goto L261
	}
L257:
	;
	v1129 = *(*int32)(unsafe.Add(mBase, _consts[5]))
	v1135 = *(*int32)(unsafe.Add(mBase, uint32(v1129+(v1123^int32(-1))<<(uint(int32(2))%32))))
	v1143 = v1135
	goto L256
L258:
	;
	goto L259
L259:
	;
	v1137 = *(*int32)(unsafe.Add(mBase, _consts[6]))
	v1143 = v1137 + v1123<<(uint(int32(13))%32) + int32(-8192)
	goto L256
L260:
	;
	v1149 = *(*int32)(unsafe.Add(mBase, uint32(v1060)))
	v1150 = *(*int32)(unsafe.Add(mBase, uint32(v1045)+52))
	v1151 = *(*int32)(unsafe.Add(mBase, uint32(v1047)))
	F_LockBuffer(m, v79, int32(0))
	mBase = m.M
	v1154 = m.ExcPending
	if v1154 != 0 {
		goto L9
	} else {
		goto L263
	}
L261:
	;
	goto L262
L262:
	;
	if v1146&int32(64) != 0 {
		goto L267
	} else {
		goto L268
	}
L263:
	;
	F_LockBuffer(m, v1123, int32(0))
	mBase = m.M
	v1157 = m.ExcPending
	if v1157 != 0 {
		goto L9
	} else {
		goto L264
	}
L264:
	;
	F__hash_finish_split(m, l0, v79, v1123, v1062, v1151, v1150, v1149)
	mBase = m.M
	v1159 = m.ExcPending
	if v1159 != 0 {
		goto L9
	} else {
		goto L265
	}
L265:
	;
	F_ReleaseBuffer(m, v1123)
	mBase = m.M
	v1161 = m.ExcPending
	if v1161 != 0 {
		goto L9
	} else {
		goto L266
	}
L266:
	;
	goto L221
L267:
	;
	v1164 = *(*int32)(unsafe.Add(mBase, uint32(v1060)))
	v1165 = *(*int32)(unsafe.Add(mBase, uint32(v1045)+52))
	v1166 = *(*int32)(unsafe.Add(mBase, uint32(v1047)))
	F_LockBuffer(m, v79, int32(0))
	mBase = m.M
	v1169 = m.ExcPending
	if v1169 != 0 {
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
	v1170 = int32(0)
	F_hashbucketcleanup(m, l0, v1062, v1123, v1095, v1170, v1166, v1165, v1164, v1170, v1170, int32(1), v1170, v1170)
	mBase = m.M
	v1177 = m.ExcPending
	if v1177 != 0 {
		goto L9
	} else {
		goto L271
	}
L271:
	;
	F_ReleaseBuffer(m, v1123)
	mBase = m.M
	v1179 = m.ExcPending
	if v1179 != 0 {
		goto L9
	} else {
		goto L272
	}
L272:
	;
	goto L221
L273:
	;
	v1205 = int32(2)
	v1208 = v1045 + int32(76)
	v1212 = *(*int32)(unsafe.Add(mBase, uint32(v1204<<(uint(v1205)%32)+v1208-int32(4))))
	v1213 = v1212 + v1055
	v1214 = int32(1)
	v1215 = v1213 + v1214
	v1219 = v1181 - v1214
	if base.Ui32(v1205) <= base.Ui32(v1181) {
		goto L282
	} else {
		goto L283
	}
L274:
	;
	v1191 = int32(32) - base.I32_clz(v1185)
	goto L276
L275:
	;
	v1191 = int32(0)
	goto L276
L276:
	;
	if base.Ui32(int32(10)) <= base.Ui32(v1191) {
		goto L277
	} else {
		goto L278
	}
L277:
	;
	v1194 = int32(3)
	v1204 = int32(base.Ui32(v1185)>>(uint(v1191-v1194)%32))&v1194 | v1191<<(uint(int32(2))%32) - int32(30)
	goto L279
L278:
	;
	v1204 = v1191
	goto L279
L279:
	;
	goto L273
L280:
	;
	F_UnlockReleaseBuffer(m, v1123)
	mBase = m.M
	v1643 = m.ExcPending
	if v1643 != 0 {
		goto L9
	} else {
		goto L391
	}
L281:
	;
	v1240 = v1045 + int32(60)
	v1241 = *(*int32)(unsafe.Add(mBase, uint32(v1240)))
	if base.Ui32(v1241) < base.Ui32(v1238) {
		goto L288
	} else {
		goto L289
	}
L282:
	;
	v1225 = int32(32) - base.I32_clz(v1219)
	goto L284
L283:
	;
	v1225 = int32(0)
	goto L284
L284:
	;
	if base.Ui32(int32(10)) <= base.Ui32(v1225) {
		goto L285
	} else {
		goto L286
	}
L285:
	;
	v1228 = int32(3)
	v1238 = int32(base.Ui32(v1219)>>(uint(v1225-v1228)%32))&v1228 | v1225<<(uint(int32(2))%32) - int32(30)
	goto L287
L286:
	;
	v1238 = v1225
	goto L287
L287:
	;
	goto L281
L288:
	;
	if base.Ui32(v1238) <= base.Ui32(int32(9)) {
		goto L292
	} else {
		goto L293
	}
L289:
	;
	goto L290
L290:
	;
	v1387 = F__hash_getnewbuf(m, l0, v1215, int32(0))
	mBase = m.M
	v1388 = m.ExcPending
	if v1388 != 0 {
		goto L9
	} else {
		goto L328
	}
L291:
	;
	v1265 = v1264 + v1212
	if base.Ui32(v1265) < base.Ui32(v1215) {
		goto L280
	} else {
		goto L295
	}
L292:
	;
	v1264 = int32(1) << (uint(v1238) % 32)
	goto L291
L293:
	;
	goto L294
L294:
	;
	v1250 = v1238 - int32(10)
	v1251 = int32(2)
	v1253 = int32(512) << (uint(int32(base.Ui32(v1250)>>(uint(v1251)%32))) % 32)
	v1264 = v1253>>(uint(v1251)%32)*(v1250&int32(3)+int32(1)) + v1253
	goto L291
L295:
	;
	if v1264-v1055^v1213 == int32(-1) {
		goto L280
	} else {
		goto L296
	}
L296:
	;
	v1272 = v993 + int32(4096)
	if v1272&int32(3) != 0 {
		goto L299
	} else {
		goto L300
	}
L297:
	;
	v1314 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v993)+uint32(_consts[56]))))
	v1317 = v1314 + (v993 + int32(4096))
	*(*int64)(unsafe.Add(mBase, uint32(v1317)+8)) = int64(-36028792723996673)
	*(*int64)(unsafe.Add(mBase, uint32(v1317))) = int64(-1)
	v1322 = *(*int32)(unsafe.Add(mBase, uint32(l0)+48))
	v1323 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1322)+118)))
	if v1323 != int32(112) {
		goto L307
	} else {
		goto L308
	}
L298:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v993)+uint32(_consts[57]))) = int32(1572864)
	v1305 = int32(8196)
	*(*uint16)(unsafe.Add(mBase, uint32(v993)+uint32(_consts[58]))) = uint16(v1305)
	v1311 = int32(8176)
	*(*uint16)(unsafe.Add(mBase, uint32(v993)+uint32(_consts[56]))) = uint16(v1311)
	*(*uint16)(unsafe.Add(mBase, uint32(v993)+uint32(_consts[59]))) = uint16(v1311)
	goto L297
L299:
	;
	v1299 = F___memset(m, v1272, int32(0), int32(8192))
	mBase = m.M
	goto L298
L300:
	;
	goto L299
L307:
	;
	v1340 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v993)+uint32(_consts[59]))))
	if v1340 == int32(0) {
		goto L316
	} else {
		goto L317
	}
L308:
	;
	v1327 = *(*int32)(unsafe.Add(mBase, _consts[10]))
	if v1327 <= int32(0) {
		goto L309
	} else {
		goto L310
	}
L309:
	;
	v1330 = *(*int32)(unsafe.Add(mBase, uint32(l0)+32))
	if v1330 != 0 {
		goto L307
	} else {
		goto L312
	}
L310:
	;
	goto L311
L311:
	;
	F_log_newpage(m, l0, int32(0), v1265, v993+int32(4096), int32(1))
	mBase = m.M
	v1337 = m.ExcPending
	if v1337 != 0 {
		goto L9
	} else {
		goto L314
	}
L312:
	;
	v1331 = *(*int32)(unsafe.Add(mBase, uint32(l0)+40))
	if v1331 != 0 {
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
	v1348 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	if v1348 != 0 {
		goto L319
	} else {
		goto L320
	}
L316:
	;
	goto L315
L317:
	;
	v1343 = F_DataChecksumsEnabled(m)
	mBase = m.M
	if v1343 == int32(0) {
		goto L316
	} else {
		goto L318
	}
L318:
	;
	v1346 = F_pg_checksum_page(m, v993+int32(4096), v1265)
	mBase = m.M
	*(*uint16)(unsafe.Add(mBase, uint32(v993)+uint32(_consts[60]))) = uint16(v1346)
	goto L316
L319:
	;
	v1376 = v1348
	goto L321
L320:
	;
	v1349 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v1352 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	*(*int32)(unsafe.Add(mBase, uint32(v993+int32(4088)))) = v1352
	v1354 = *(*int64)(unsafe.Add(mBase, uint32(l0)))
	*(*int64)(unsafe.Add(mBase, uint32(v993)+4080)) = v1354
	v1358 = F_smgropen(m, v993+int32(4080), v1349)
	mBase = m.M
	v1359 = m.ExcPending
	if v1359 != 0 {
		goto L9
	} else {
		goto L322
	}
L321:
	;
	v1377 = int32(0)
	F_smgrextend(m, v1376, v1377, v1265, v993+int32(4096), v1377)
	mBase = m.M
	v1382 = m.ExcPending
	if v1382 != 0 {
		goto L9
	} else {
		goto L327
	}
L322:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+12)) = v1358
	v1362 = *(*int32)(unsafe.Add(mBase, uint32(v1358)+72))
	if v1362 != 0 {
		goto L324
	} else {
		goto L325
	}
L323:
	;
	v1374 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	v1376 = v1374
	goto L321
L324:
	;
	v1370 = v1362
	goto L326
L325:
	;
	v1363 = *(*int32)(unsafe.Add(mBase, uint32(v1358)+76))
	v1364 = *(*int32)(unsafe.Add(mBase, uint32(v1358)+80))
	*(*int32)(unsafe.Add(mBase, uint32(v1363)+4)) = v1364
	v1366 = *(*int32)(unsafe.Add(mBase, uint32(v1358)+76))
	*(*int32)(unsafe.Add(mBase, uint32(v1364))) = v1366
	v1368 = *(*int32)(unsafe.Add(mBase, uint32(v1358)+72))
	v1370 = v1368
	goto L326
L326:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v1358)+72)) = v1370 + int32(1)
	goto L323
L327:
	;
	goto L290
L328:
	;
	v1389 = F_IsBufferCleanupOK(m, v1387)
	mBase = m.M
	v1390 = m.ExcPending
	if v1390 != 0 {
		goto L9
	} else {
		goto L329
	}
L329:
	;
	if v1389 == int32(0) {
		goto L330
	} else {
		goto L331
	}
L330:
	;
	F_UnlockReleaseBuffer(m, v1123)
	mBase = m.M
	v1394 = m.ExcPending
	if v1394 != 0 {
		goto L9
	} else {
		goto L333
	}
L331:
	;
	goto L332
L332:
	;
	v1397 = int32(4543684)
	v1399 = *(*int32)(unsafe.Add(mBase, _consts[7]))
	*(*int32)(unsafe.Add(mBase, _consts[7])) = v1399 + int32(1)
	*(*int32)(unsafe.Add(mBase, uint32(v1045)+48)) = v1055
	v1405 = v1045 + int32(52)
	v1406 = *(*int32)(unsafe.Add(mBase, uint32(v1405)))
	v1407 = base.B2i32(base.Ui32(v1055) <= base.Ui32(v1406))
	if v1407 == int32(0) {
		goto L335
	} else {
		goto L336
	}
L333:
	;
	F_UnlockReleaseBuffer(m, v1387)
	mBase = m.M
	v1396 = m.ExcPending
	if v1396 != 0 {
		goto L9
	} else {
		goto L334
	}
L334:
	;
	goto L220
L335:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v1060))) = v1406
	*(*int32)(unsafe.Add(mBase, uint32(v1405))) = v1406 | v1055
	goto L337
L336:
	;
	goto L337
L337:
	;
	v1413 = *(*int32)(unsafe.Add(mBase, uint32(v1240)))
	v1414 = base.B2i32(base.Ui32(v1238) <= base.Ui32(v1413))
	if v1414 == int32(0) {
		goto L338
	} else {
		goto L339
	}
L338:
	;
	v1417 = int32(2)
	v1423 = *(*int32)(unsafe.Add(mBase, uint32(v1208+v1413<<(uint(v1417)%32))))
	*(*int32)(unsafe.Add(mBase, uint32(v1208+v1238<<(uint(v1417)%32)))) = v1423
	*(*int32)(unsafe.Add(mBase, uint32(v1240))) = v1238
	goto L340
L339:
	;
	goto L340
L340:
	;
	F_MarkBufferDirty(m, v79)
	mBase = m.M
	v1427 = m.ExcPending
	if v1427 != 0 {
		goto L9
	} else {
		goto L341
	}
L341:
	;
	v1432 = *(*int32)(unsafe.Add(mBase, uint32(v1045)+48))
	v1433 = *(*int32)(unsafe.Add(mBase, uint32(v1045+int32(56))))
	v1434 = *(*int32)(unsafe.Add(mBase, uint32(v1045+int32(52))))
	if v1123 < int32(0) {
		goto L343
	} else {
		goto L344
	}
L342:
	;
	v1453 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v1452)+16)))
	v1454 = v1453 + v1452
	*(*int32)(unsafe.Add(mBase, uint32(v1454))) = v1432
	v1456 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v1454)+12)))
	v1458 = v1456 | int32(32)
	*(*uint16)(unsafe.Add(mBase, uint32(v1454)+12)) = uint16(v1458)
	F_MarkBufferDirty(m, v1123)
	mBase = m.M
	v1461 = m.ExcPending
	if v1461 != 0 {
		goto L9
	} else {
		goto L346
	}
L343:
	;
	v1438 = *(*int32)(unsafe.Add(mBase, _consts[5]))
	v1444 = *(*int32)(unsafe.Add(mBase, uint32(v1438+(v1123^int32(-1))<<(uint(int32(2))%32))))
	v1452 = v1444
	goto L342
L344:
	;
	goto L345
L345:
	;
	v1446 = *(*int32)(unsafe.Add(mBase, _consts[6]))
	v1452 = v1446 + v1123<<(uint(int32(13))%32) + int32(-8192)
	goto L342
L346:
	;
	if v1387 < int32(0) {
		goto L348
	} else {
		goto L349
	}
L347:
	;
	v1480 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v1479)+16)))
	v1481 = v1480 + v1479
	*(*int32)(unsafe.Add(mBase, uint32(v1481)+12)) = int32(-8388590)
	*(*int32)(unsafe.Add(mBase, uint32(v1481)+8)) = v1055
	*(*int32)(unsafe.Add(mBase, uint32(v1481)+4)) = int32(-1)
	*(*int32)(unsafe.Add(mBase, uint32(v1481))) = v1432
	F_MarkBufferDirty(m, v1387)
	mBase = m.M
	v1489 = m.ExcPending
	if v1489 != 0 {
		goto L9
	} else {
		goto L351
	}
L348:
	;
	v1465 = *(*int32)(unsafe.Add(mBase, _consts[5]))
	v1471 = *(*int32)(unsafe.Add(mBase, uint32(v1465+(v1387^int32(-1))<<(uint(int32(2))%32))))
	v1479 = v1471
	goto L347
L349:
	;
	goto L350
L350:
	;
	v1473 = *(*int32)(unsafe.Add(mBase, _consts[6]))
	v1479 = v1473 + v1387<<(uint(int32(13))%32) + int32(-8192)
	goto L347
L351:
	;
	v1490 = *(*int32)(unsafe.Add(mBase, uint32(l0)+48))
	v1491 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1490)+118)))
	if v1491 != int32(112) {
		goto L352
	} else {
		goto L353
	}
L352:
	;
	v1625 = int32(4543684)
	v1627 = *(*int32)(unsafe.Add(mBase, _consts[7]))
	*(*int32)(unsafe.Add(mBase, _consts[7])) = v1627 - int32(1)
	F_LockBuffer(m, v79, int32(0))
	mBase = m.M
	v1633 = m.ExcPending
	if v1633 != 0 {
		goto L9
	} else {
		goto L387
	}
L353:
	;
	v1495 = *(*int32)(unsafe.Add(mBase, _consts[10]))
	if v1495 <= int32(0) {
		goto L354
	} else {
		goto L355
	}
L354:
	;
	v1498 = *(*int32)(unsafe.Add(mBase, uint32(l0)+32))
	if v1498 != 0 {
		goto L352
	} else {
		goto L357
	}
L355:
	;
	goto L356
L356:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v993)+uint32(_consts[61]))) = v1432
	v1501 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v1454)+12)))
	*(*uint16)(unsafe.Add(mBase, uint32(v993)+uint32(_consts[62]))) = uint16(v1501)
	v1503 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v1481)+12)))
	v1504 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v993)+uint32(_consts[60]))) = uint8(v1504)
	*(*uint16)(unsafe.Add(mBase, uint32(v993)+uint32(_consts[63]))) = uint16(v1503)
	F_XLogBeginInsert(m)
	mBase = m.M
	v1508 = m.ExcPending
	if v1508 != 0 {
		goto L9
	} else {
		goto L359
	}
L357:
	;
	v1499 = *(*int32)(unsafe.Add(mBase, uint32(l0)+40))
	if v1499 != 0 {
		goto L352
	} else {
		goto L358
	}
L358:
	;
	goto L356
L359:
	;
	F_XLogRegisterBuffer(m, int32(0), v1123, int32(8))
	mBase = m.M
	v1512 = m.ExcPending
	if v1512 != 0 {
		goto L9
	} else {
		goto L360
	}
L360:
	;
	F_XLogRegisterBuffer(m, int32(1), v1387, int32(6))
	mBase = m.M
	v1516 = m.ExcPending
	if v1516 != 0 {
		goto L9
	} else {
		goto L361
	}
L361:
	;
	F_XLogRegisterBuffer(m, int32(2), v79, int32(8))
	mBase = m.M
	v1520 = m.ExcPending
	if v1520 != 0 {
		goto L9
	} else {
		goto L362
	}
L362:
	;
	if v1407 == int32(0) {
		goto L363
	} else {
		goto L364
	}
L363:
	;
	v1523 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(v993)+uint32(_consts[60]))) = uint8(v1523)
	F_XLogRegisterBufData(m, int32(2), v1060, int32(4))
	mBase = m.M
	v1528 = m.ExcPending
	if v1528 != 0 {
		goto L9
	} else {
		goto L366
	}
L364:
	;
	goto L365
L365:
	;
	if v1414 == int32(0) {
		goto L368
	} else {
		goto L369
	}
L366:
	;
	F_XLogRegisterBufData(m, int32(2), v1405, int32(4))
	mBase = m.M
	v1532 = m.ExcPending
	if v1532 != 0 {
		goto L9
	} else {
		goto L367
	}
L367:
	;
	goto L365
L368:
	;
	v1535 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v993)+uint32(_consts[60]))))
	v1536 = int32(2)
	v1537 = v1535 | v1536
	*(*uint8)(unsafe.Add(mBase, uint32(v993)+uint32(_consts[60]))) = uint8(v1537)
	F_XLogRegisterBufData(m, v1536, v1240, int32(4))
	mBase = m.M
	v1542 = m.ExcPending
	if v1542 != 0 {
		goto L9
	} else {
		goto L371
	}
L369:
	;
	goto L370
L370:
	;
	F_XLogRegisterData(m, v993+int32(4096), int32(9))
	mBase = m.M
	v1555 = m.ExcPending
	if v1555 != 0 {
		goto L9
	} else {
		goto L373
	}
L371:
	;
	v1543 = int32(2)
	v1544 = *(*int32)(unsafe.Add(mBase, uint32(v1240)))
	F_XLogRegisterBufData(m, v1543, v1208+v1544<<(uint(v1543)%32), int32(4))
	mBase = m.M
	v1550 = m.ExcPending
	if v1550 != 0 {
		goto L9
	} else {
		goto L372
	}
L372:
	;
	goto L370
L373:
	;
	v1558 = F_XLogInsert(m, int32(12), int32(64))
	mBase = m.M
	v1559 = m.ExcPending
	if v1559 != 0 {
		goto L9
	} else {
		goto L374
	}
L374:
	;
	if v1123 < int32(0) {
		goto L376
	} else {
		goto L377
	}
L375:
	;
	v1578 = int64(32)
	*(*int64)(unsafe.Add(mBase, uint32(v1577))) = base.I64_rotr(v1558, v1578)
	v1583 = base.I32_wrap_i64(int64(base.Ui64(v1558) >> (uint(v1578) % 64)))
	v1584 = base.I32_wrap_i64(v1558)
	if v1387 < int32(0) {
		goto L380
	} else {
		goto L381
	}
L376:
	;
	v1563 = *(*int32)(unsafe.Add(mBase, _consts[5]))
	v1569 = *(*int32)(unsafe.Add(mBase, uint32(v1563+(v1123^int32(-1))<<(uint(int32(2))%32))))
	v1577 = v1569
	goto L375
L377:
	;
	goto L378
L378:
	;
	v1571 = *(*int32)(unsafe.Add(mBase, _consts[6]))
	v1577 = v1571 + v1123<<(uint(int32(13))%32) + int32(-8192)
	goto L375
L379:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v1602)+4)) = v1584
	*(*int32)(unsafe.Add(mBase, uint32(v1602))) = v1583
	if v79 < int32(0) {
		goto L384
	} else {
		goto L385
	}
L380:
	;
	v1588 = *(*int32)(unsafe.Add(mBase, _consts[5]))
	v1594 = *(*int32)(unsafe.Add(mBase, uint32(v1588+(v1387^int32(-1))<<(uint(int32(2))%32))))
	v1602 = v1594
	goto L379
L381:
	;
	goto L382
L382:
	;
	v1596 = *(*int32)(unsafe.Add(mBase, _consts[6]))
	v1602 = v1596 + v1387<<(uint(int32(13))%32) + int32(-8192)
	goto L379
L383:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v1618)+4)) = v1584
	*(*int32)(unsafe.Add(mBase, uint32(v1618))) = v1583
	goto L352
L384:
	;
	v1608 = *(*int32)(unsafe.Add(mBase, _consts[5]))
	v1612 = *(*int32)(unsafe.Add(mBase, uint32(v1608+v996<<(uint(int32(2))%32))))
	v1618 = v1612
	goto L383
L385:
	;
	goto L386
L386:
	;
	v1614 = *(*int32)(unsafe.Add(mBase, _consts[6]))
	v1618 = v1614 + v998 + int32(-8192)
	goto L383
L387:
	;
	F__hash_splitbucket(m, l0, v79, v1062, v1055, v1123, v1387, int32(0), v1432, v1434, v1433)
	mBase = m.M
	v1636 = m.ExcPending
	if v1636 != 0 {
		goto L9
	} else {
		goto L388
	}
L388:
	;
	F_ReleaseBuffer(m, v1123)
	mBase = m.M
	v1638 = m.ExcPending
	if v1638 != 0 {
		goto L9
	} else {
		goto L389
	}
L389:
	;
	F_ReleaseBuffer(m, v1387)
	mBase = m.M
	v1640 = m.ExcPending
	if v1640 != 0 {
		goto L9
	} else {
		goto L390
	}
L390:
	;
	m.G0 = v989
	goto L219
L391:
	;
	goto L220
L392:
	;
	m.G0 = v989
	goto L219
L393:
	;
	m.G0 = v30 + int32(4128)
	return
L394:
	;
	F_errcode(m, int32(261))
	mBase = m.M
	v1713 = m.ExcPending
	if v1713 != 0 {
		goto L9
	} else {
		goto L395
	}
L395:
	;
	v1714 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v98)+19)))
	*(*int32)(unsafe.Add(mBase, uint32(v30))) = v48
	*(*int32)(unsafe.Add(mBase, uint32(v30)+4)) = (v1714<<(uint(int32(8))%32) - int32(44)) & int32(-48)
	F_errmsg(m, int32(38202), v30)
	mBase = m.M
	v1725 = m.ExcPending
	if v1725 != 0 {
		goto L9
	} else {
		goto L396
	}
L396:
	;
	F_errhint(m, int32(670104), int32(0))
	mBase = m.M
	v1729 = m.ExcPending
	if v1729 != 0 {
		goto L9
	} else {
		goto L397
	}
L397:
	;
	F_errfinish(m, int32(512279), int32(86), int32(87186))
	mBase = m.M
	v1734 = m.ExcPending
	if v1734 != 0 {
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
	var v99 int32
	_ = v99
	var v102 int32
	_ = v102
	var v110 int32
	_ = v110
	var v113 int32
	_ = v113
	var v123 int32
	_ = v123
	var v124 int32
	_ = v124
	var v127 int32
	_ = v127
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
	var v153 int32
	_ = v153
	var v157 int32
	_ = v157
	var v176 int32
	_ = v176
	var v192 int32
	_ = v192
	var v199 int32
	_ = v199
	var v200 int32
	_ = v200
	var v222 int32
	_ = v222
	var v226 int32
	_ = v226
	var v228 int32
	_ = v228
	var v231 int32
	_ = v231
	var v232 int32
	_ = v232
	var v235 int32
	_ = v235
	var v236 int32
	_ = v236
	var v241 int32
	_ = v241
	var v245 int32
	_ = v245
	var v251 int32
	_ = v251
	var v253 int32
	_ = v253
	var v259 int32
	_ = v259
	var v260 int32
	_ = v260
	var v262 int32
	_ = v262
	var v264 int32
	_ = v264
	var v266 int32
	_ = v266
	var v269 int32
	_ = v269
	var v276 int32
	_ = v276
	var v280 int32
	_ = v280
	var v285 int32
	_ = v285
	v19 = m.G0
	v21 = v19 + int32(-64)
	m.G0 = v21
	*(*int64)(unsafe.Add(mBase, uint32(v21)+32)) = int64(25769803782)
	v26 = *(*int32)(unsafe.Add(mBase, _consts[0]))
	*(*int32)(unsafe.Add(mBase, uint32(v21)+56)) = v26
	v33 = F_hash_create(m, int32(182093), int32(256), v19+int32(-48), int32(1064))
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
	v43 = *(*int32)(unsafe.Add(mBase, _consts[5]))
	v49 = *(*int32)(unsafe.Add(mBase, uint32(v43+(v38^int32(-1))<<(uint(int32(2))%32))))
	v57 = v49
	goto L3
L6:
	;
	goto L7
L7:
	;
	v51 = *(*int32)(unsafe.Add(mBase, _consts[6]))
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
	v99 = v58
	goto L13
L13:
	;
	F_UnlockReleaseBuffer(m, v38)
	mBase = m.M
	v102 = m.ExcPending
	if v102 != 0 {
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
	v99 = v96 + v76
	goto L13
L20:
	;
	v110 = v99
	v113 = int32(0)
	goto L22
L21:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v276 = m.ExcPending
	if v276 != 0 {
		goto L1
	} else {
		goto L64
	}
L22:
	;
	if v110 == int32(-1) {
		goto L21
	} else {
		goto L24
	}
L23:
	;
	v231 = F_ConditionalLockBufferForCleanup(m, l2)
	mBase = m.M
	v232 = m.ExcPending
	if v232 != 0 {
		goto L1
	} else {
		goto L50
	}
L24:
	;
	v123 = F_ReadBuffer(m, l0, v110)
	mBase = m.M
	v124 = m.ExcPending
	if v124 != 0 {
		goto L1
	} else {
		goto L25
	}
L25:
	;
	F_LockBuffer(m, v123, int32(1))
	mBase = m.M
	v127 = m.ExcPending
	if v127 != 0 {
		goto L1
	} else {
		goto L26
	}
L26:
	;
	F__hash_checkpage(m, l0, v123, int32(3))
	mBase = m.M
	v130 = m.ExcPending
	if v130 != 0 {
		goto L1
	} else {
		goto L27
	}
L27:
	;
	if v110 == v99 {
		goto L28
	} else {
		goto L29
	}
L28:
	;
	v132 = v123
	goto L30
L29:
	;
	v132 = v113
	goto L30
L30:
	;
	if v123 < int32(0) {
		goto L32
	} else {
		goto L33
	}
L31:
	;
	v151 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v150)+16)))
	v153 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v150)+12)))
	if base.Ui32(v153) < base.Ui32(int32(25)) {
		goto L35
	} else {
		goto L36
	}
L32:
	;
	v136 = *(*int32)(unsafe.Add(mBase, _consts[5]))
	v142 = *(*int32)(unsafe.Add(mBase, uint32(v136+(v123^int32(-1))<<(uint(int32(2))%32))))
	v150 = v142
	goto L31
L33:
	;
	goto L34
L34:
	;
	v144 = *(*int32)(unsafe.Add(mBase, _consts[6]))
	v150 = v144 + v123<<(uint(int32(13))%32) + int32(-8192)
	goto L31
L35:
	;
	v222 = *(*int32)(unsafe.Add(mBase, uint32(v151+v150)+4))
	if v123 == v132 {
		goto L43
	} else {
		goto L44
	}
L36:
	;
	v157 = v153 + int32(262120)
	if v157&int32(262140) == int32(0) {
		goto L35
	} else {
		goto L37
	}
L37:
	;
	v176 = int32(1)
	goto L38
L38:
	;
	v192 = *(*int32)(unsafe.Add(mBase, uint32(v176<<(uint(int32(2))%32)+(v150+int32(24))-int32(4))))
	v199 = F_hash_search(m, v33, v150+v192&int32(32767), int32(1), v19+int32(-49))
	mBase = m.M
	v200 = m.ExcPending
	if v200 != 0 {
		goto L1
	} else {
		goto L40
	}
L39:
	;
	goto L35
L40:
	;
	if v176 != int32(base.Ui32(v157)>>(uint(int32(2))%32))&int32(65535) {
		v176 = v176 + int32(1)
		goto L38
	} else {
		goto L41
	}
L41:
	;
	goto L39
L42:
	;
	if v222 != int32(-1) {
		v110 = v222
		v113 = v132
		goto L22
	} else {
		goto L48
	}
L43:
	;
	F_LockBuffer(m, v123, int32(0))
	mBase = m.M
	v226 = m.ExcPending
	if v226 != 0 {
		goto L1
	} else {
		goto L46
	}
L44:
	;
	goto L45
L45:
	;
	F_UnlockReleaseBuffer(m, v123)
	mBase = m.M
	v228 = m.ExcPending
	if v228 != 0 {
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
	v269 = m.ExcPending
	if v269 != 0 {
		goto L1
	} else {
		goto L63
	}
L50:
	;
	if v231 == int32(0) {
		goto L49
	} else {
		goto L51
	}
L51:
	;
	v235 = F_ConditionalLockBufferForCleanup(m, v132)
	mBase = m.M
	v236 = m.ExcPending
	if v236 != 0 {
		goto L1
	} else {
		goto L52
	}
L52:
	;
	if v235 == int32(0) {
		goto L53
	} else {
		goto L54
	}
L53:
	;
	F_LockBuffer(m, l2, int32(0))
	mBase = m.M
	v241 = m.ExcPending
	if v241 != 0 {
		goto L1
	} else {
		goto L56
	}
L54:
	;
	goto L55
L55:
	;
	if v132 < int32(0) {
		goto L58
	} else {
		goto L59
	}
L56:
	;
	goto L49
L57:
	;
	v260 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v259)+16)))
	v262 = *(*int32)(unsafe.Add(mBase, uint32(v260+v259)+8))
	F__hash_splitbucket(m, l0, l1, l3, v262, l2, v132, v33, l4, l5, l6)
	mBase = m.M
	v264 = m.ExcPending
	if v264 != 0 {
		goto L1
	} else {
		goto L61
	}
L58:
	;
	v245 = *(*int32)(unsafe.Add(mBase, _consts[5]))
	v251 = *(*int32)(unsafe.Add(mBase, uint32(v245+(v132^int32(-1))<<(uint(int32(2))%32))))
	v259 = v251
	goto L57
L59:
	;
	goto L60
L60:
	;
	v253 = *(*int32)(unsafe.Add(mBase, _consts[6]))
	v259 = v253 + v132<<(uint(int32(13))%32) + int32(-8192)
	goto L57
L61:
	;
	F_ReleaseBuffer(m, v132)
	mBase = m.M
	v266 = m.ExcPending
	if v266 != 0 {
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
	F_errmsg_internal(m, int32(537435), int32(0))
	mBase = m.M
	v280 = m.ExcPending
	if v280 != 0 {
		goto L1
	} else {
		goto L65
	}
L65:
	;
	F_errfinish(m, int32(519802), int32(75), int32(352350))
	mBase = m.M
	v285 = m.ExcPending
	if v285 != 0 {
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
	var v102 int32
	_ = v102
	var v106 int32
	_ = v106
	var v107 int32
	_ = v107
	var v112 int32
	_ = v112
	var v118 int32
	_ = v118
	var v120 int32
	_ = v120
	var v126 int32
	_ = v126
	var v127 int32
	_ = v127
	var v128 int32
	_ = v128
	var v130 int32
	_ = v130
	var v134 int32
	_ = v134
	var v140 int32
	_ = v140
	var v142 int32
	_ = v142
	var v148 int32
	_ = v148
	var v149 int32
	_ = v149
	var v150 int32
	_ = v150
	var v152 int32
	_ = v152
	var v154 int32
	_ = v154
	var v159 int32
	_ = v159
	var v166 int32
	_ = v166
	var v170 int32
	_ = v170
	var v171 int32
	_ = v171
	var v175 int32
	_ = v175
	var v181 int32
	_ = v181
	var v183 int32
	_ = v183
	var v189 int32
	_ = v189
	var v193 int32
	_ = v193
	var v194 int32
	_ = v194
	var v207 int32
	_ = v207
	var v213 int32
	_ = v213
	var v218 int32
	_ = v218
	var v220 int32
	_ = v220
	var v226 int32
	_ = v226
	var v229 int32
	_ = v229
	var v230 int32
	_ = v230
	var v234 int32
	_ = v234
	var v237 int32
	_ = v237
	var v241 int32
	_ = v241
	var v247 int32
	_ = v247
	var v249 int32
	_ = v249
	var v255 int32
	_ = v255
	var v256 int32
	_ = v256
	var v257 int32
	_ = v257
	var v259 int32
	_ = v259
	var v262 int32
	_ = v262
	var v264 int32
	_ = v264
	var v266 int32
	_ = v266
	var v269 int32
	_ = v269
	var v278 int32
	_ = v278
	var v287 int32
	_ = v287
	var v290 int32
	_ = v290
	var v293 int32
	_ = v293
	var v301 int32
	_ = v301
	var v302 int32
	_ = v302
	var v303 int32
	_ = v303
	var v315 int32
	_ = v315
	var v320 int32
	_ = v320
	var v321 int32
	_ = v321
	var v326 int32
	_ = v326
	var v329 int32
	_ = v329
	var v330 int32
	_ = v330
	var v332 int32
	_ = v332
	var v337 int32
	_ = v337
	var v353 int32
	_ = v353
	var v356 int32
	_ = v356
	var v360 int32
	_ = v360
	var v365 int32
	_ = v365
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
		v337 = v40
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
	v353 = m.ExcPending
	if v353 != 0 {
		goto L6
	} else {
		goto L82
	}
L14:
	;
	m.G0 = v14 + int32(16)
	return v337
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
	*(*int32)(unsafe.Add(mBase, uint32(v16))) = v102
	v106 = F__hash_getbucketbuf_from_hashkey(m, v17, v102, int32(1), int32(0))
	mBase = m.M
	v107 = m.ExcPending
	if v107 != 0 {
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
	v102 = v59
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
	F_errmsg_internal(m, int32(719333), v63)
	mBase = m.M
	v86 = m.ExcPending
	if v86 != 0 {
		goto L6
	} else {
		goto L29
	}
L29:
	;
	F_errfinish(m, int32(517837), int32(115), int32(381446))
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
	v102 = v94
	goto L16
L32:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v14)+12)) = v106
	if v106 < int32(0) {
		goto L34
	} else {
		goto L35
	}
L33:
	;
	v128 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	F_PredicateLockPage(m, v17, v127, v128)
	mBase = m.M
	v130 = m.ExcPending
	if v130 != 0 {
		goto L6
	} else {
		goto L37
	}
L34:
	;
	v112 = *(*int32)(unsafe.Add(mBase, _consts[8]))
	v118 = *(*int32)(unsafe.Add(mBase, uint32(v112+(v106^int32(-1))<<(uint(int32(6))%32))+16))
	v127 = v118
	goto L33
L35:
	;
	goto L36
L36:
	;
	v120 = *(*int32)(unsafe.Add(mBase, _consts[9]))
	v126 = *(*int32)(unsafe.Add(mBase, uint32(v120+v106<<(uint(int32(6))%32)+int32(-64))+16))
	v127 = v126
	goto L33
L37:
	;
	if v106 < int32(0) {
		goto L39
	} else {
		goto L40
	}
L38:
	;
	v149 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v148)+16)))
	v150 = v149 + v148
	*(*int32)(unsafe.Add(mBase, uint32(v14)+4)) = v150
	v152 = *(*int32)(unsafe.Add(mBase, uint32(v150)+8))
	*(*int32)(unsafe.Add(mBase, uint32(v16)+4)) = v106
	v154 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v150)+12)))
	if v154&int32(16) == int32(0) {
		v269 = v150
		goto L42
	} else {
		goto L43
	}
L39:
	;
	v134 = *(*int32)(unsafe.Add(mBase, _consts[5]))
	v140 = *(*int32)(unsafe.Add(mBase, uint32(v134+(v106^int32(-1))<<(uint(int32(2))%32))))
	v148 = v140
	goto L38
L40:
	;
	goto L41
L41:
	;
	v142 = *(*int32)(unsafe.Add(mBase, _consts[6]))
	v148 = v142 + v106<<(uint(int32(13))%32) + int32(-8192)
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
	v159 = int32(-1)
	v166 = v152 & (v159<<(uint(int32(31)-base.I32_clz(v152))%32) ^ v159)
	v170 = F__hash_getbuf(m, v17, int32(0), int32(1), int32(8))
	mBase = m.M
	v171 = m.ExcPending
	if v171 != 0 {
		goto L6
	} else {
		goto L45
	}
L44:
	;
	if v166 != 0 {
		goto L49
	} else {
		goto L50
	}
L45:
	;
	if v170 < int32(0) {
		goto L46
	} else {
		goto L47
	}
L46:
	;
	v175 = *(*int32)(unsafe.Add(mBase, _consts[5]))
	v181 = *(*int32)(unsafe.Add(mBase, uint32(v175+(v170^int32(-1))<<(uint(int32(2))%32))))
	v189 = v181
	goto L44
L47:
	;
	goto L48
L48:
	;
	v183 = *(*int32)(unsafe.Add(mBase, _consts[6]))
	v189 = v183 + v170<<(uint(int32(13))%32) + int32(-8192)
	goto L44
L49:
	;
	v193 = base.I32_clz(v166)
	v194 = int32(32) - v193
	if base.Ui32(int32(512)) <= base.Ui32(v166) {
		goto L52
	} else {
		goto L53
	}
L50:
	;
	v218 = int32(0)
	goto L51
L51:
	;
	F_UnlockReleaseBuffer(m, v170)
	mBase = m.M
	v220 = m.ExcPending
	if v220 != 0 {
		goto L6
	} else {
		goto L55
	}
L52:
	;
	v207 = int32(base.Ui32(v166)>>(uint(int32(29)-v193)%32))&int32(3) | v194<<(uint(int32(2))%32) - int32(30)
	goto L54
L53:
	;
	v207 = v194
	goto L54
L54:
	;
	v213 = *(*int32)(unsafe.Add(mBase, uint32(v207<<(uint(int32(2))%32)+(v189+int32(76))-int32(4))))
	v218 = v213
	goto L51
L55:
	;
	F_LockBuffer(m, v106, int32(0))
	mBase = m.M
	v226 = m.ExcPending
	if v226 != 0 {
		goto L6
	} else {
		goto L56
	}
L56:
	;
	v229 = F__hash_getbuf(m, v17, v218+v166+int32(1), int32(1), int32(2))
	mBase = m.M
	v230 = m.ExcPending
	if v230 != 0 {
		goto L6
	} else {
		goto L57
	}
L57:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v16)+8)) = v229
	F_LockBuffer(m, v229, int32(0))
	mBase = m.M
	v234 = m.ExcPending
	if v234 != 0 {
		goto L6
	} else {
		goto L58
	}
L58:
	;
	F_LockBuffer(m, v106, int32(1))
	mBase = m.M
	v237 = m.ExcPending
	if v237 != 0 {
		goto L6
	} else {
		goto L59
	}
L59:
	;
	if v106 < int32(0) {
		goto L61
	} else {
		goto L62
	}
L60:
	;
	v256 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v255)+16)))
	v257 = v256 + v255
	*(*int32)(unsafe.Add(mBase, uint32(v14)+4)) = v257
	v259 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v257)+12)))
	if v259&int32(16) != 0 {
		goto L64
	} else {
		goto L65
	}
L61:
	;
	v241 = *(*int32)(unsafe.Add(mBase, _consts[5]))
	v247 = *(*int32)(unsafe.Add(mBase, uint32(v241+(v106^int32(-1))<<(uint(int32(2))%32))))
	v255 = v247
	goto L60
L62:
	;
	goto L63
L63:
	;
	v249 = *(*int32)(unsafe.Add(mBase, _consts[6]))
	v255 = v249 + v106<<(uint(int32(13))%32) + int32(-8192)
	goto L60
L64:
	;
	v262 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(v16)+12)) = uint8(v262)
	v269 = v257
	goto L42
L65:
	;
	goto L66
L66:
	;
	v264 = *(*int32)(unsafe.Add(mBase, uint32(v16)+8))
	F_ReleaseBuffer(m, v264)
	mBase = m.M
	v266 = m.ExcPending
	if v266 != 0 {
		goto L6
	} else {
		goto L67
	}
L67:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v16)+8)) = int32(0)
	v269 = v257
	goto L42
L68:
	;
	v278 = v269
	goto L71
L69:
	;
	v315 = v106
	goto L70
L70:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v16)+24)) = v315
	v320 = F__hash_readpage(m, l0, v14+int32(12), l1)
	mBase = m.M
	v321 = m.ExcPending
	if v321 != 0 {
		goto L6
	} else {
		goto L80
	}
L71:
	;
	v287 = *(*int32)(unsafe.Add(mBase, uint32(v278)+4))
	if v287 == int32(-1) {
		goto L74
	} else {
		goto L75
	}
L72:
	;
	v303 = *(*int32)(unsafe.Add(mBase, uint32(v14)+12))
	v315 = v303
	goto L70
L73:
	;
	goto L72
L74:
	;
	v290 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v16)+12)))
	if v290 != int32(1) {
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
	v301 = m.ExcPending
	if v301 != 0 {
		goto L6
	} else {
		goto L79
	}
L77:
	;
	v293 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v16)+13)))
	if v293 != 0 {
		goto L73
	} else {
		goto L78
	}
L78:
	;
	goto L76
L79:
	;
	v302 = *(*int32)(unsafe.Add(mBase, uint32(v14)+4))
	v278 = v302
	goto L71
L80:
	;
	if v320 == int32(0) {
		v337 = int32(0)
		goto L14
	} else {
		goto L81
	}
L81:
	;
	v326 = *(*int32)(unsafe.Add(mBase, uint32(v16)+48))
	v329 = v16 + v326<<(uint(int32(3))%32)
	v330 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v329)+56)))
	*(*uint16)(unsafe.Add(mBase, uint32(l0-int32(-64)))) = uint16(v330)
	v332 = *(*int32)(unsafe.Add(mBase, uint32(v329)+52))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+60)) = v332
	v337 = int32(1)
	goto L14
L82:
	;
	F_errcode(m, int32(1088))
	mBase = m.M
	v356 = m.ExcPending
	if v356 != 0 {
		goto L6
	} else {
		goto L83
	}
L83:
	;
	F_errmsg(m, int32(158071), int32(0))
	mBase = m.M
	v360 = m.ExcPending
	if v360 != 0 {
		goto L6
	} else {
		goto L84
	}
L84:
	;
	F_errfinish(m, int32(518446), int32(313), int32(73313))
	mBase = m.M
	v365 = m.ExcPending
	if v365 != 0 {
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
			F_errmsg_internal(m, int32(537435), int32(0))
			v26 = m.ExcPending
			if v26 != 0 {
				return int32(0)
			} else {
				F_errfinish(m, int32(519802), int32(246), int32(20773))
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
	var v107 int32
	_ = v107
	var v110 int32
	_ = v110
	var v113 int32
	_ = v113
	var v116 int32
	_ = v116
	var v118 int32
	_ = v118
	var v119 int32
	_ = v119
	var v120 int32
	_ = v120
	var v125 int32
	_ = v125
	var v131 int32
	_ = v131
	var v137 int32
	_ = v137
	var v139 int32
	_ = v139
	var v147 int32
	_ = v147
	var v148 int32
	_ = v148
	var v152 int32
	_ = v152
	var v157 int32
	_ = v157
	var v158 int32
	_ = v158
	var v164 int32
	_ = v164
	var v169 int32
	_ = v169
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
	v21 = *(*int32)(unsafe.Add(mBase, _consts[5]))
	v27 = *(*int32)(unsafe.Add(mBase, uint32(v21+(l1^int32(-1))<<(uint(int32(2))%32))))
	v35 = v27
	goto L3
L5:
	;
	goto L6
L6:
	;
	v29 = *(*int32)(unsafe.Add(mBase, _consts[6]))
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
	v157 = m.ExcPending
	if v157 != 0 {
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
	*(*uint16)(unsafe.Add(mBase, uint32(l3+v40<<(uint(int32(1))%32)))) = uint16(v137)
	v139 = *(*int32)(unsafe.Add(mBase, uint32(v51)))
	v147 = F_PageAddItemExtended(m, v35, v139, (v53&int32(8191)+int32(7))&int32(16376), v137, int32(0))
	mBase = m.M
	v148 = m.ExcPending
	if v148 != 0 {
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
	v81 = int32(base.Ui32(v72+int32(262120))>>(uint(int32(2))%32)) + v70
	goto L19
L19:
	;
	if base.Ui32(int32(2)) <= base.Ui32(v81&int32(65535)) {
		goto L20
	} else {
		goto L21
	}
L20:
	;
	v90 = v70
	v91 = v81
	goto L23
L21:
	;
	v131 = v70
	goto L22
L22:
	;
	v137 = v131 & int32(65535)
	goto L16
L23:
	;
	v95 = int32(65535)
	v101 = int32(base.Ui32(v91&v95+v90&v95) >> (uint(int32(1)) % 32))
	v107 = *(*int32)(unsafe.Add(mBase, uint32(v101<<(uint(int32(2))%32)+(v35+int32(24))-int32(4))))
	v110 = v35 + v107&int32(32767)
	v113 = int32(*(*int16)(unsafe.Add(mBase, uint32(v110)+6)))
	if int32(0) <= v113 {
		goto L25
	} else {
		goto L26
	}
L24:
	;
	v131 = v125
	goto L22
L25:
	;
	v116 = int32(8)
	goto L27
L26:
	;
	v116 = int32(16)
	goto L27
L27:
	;
	v118 = *(*int32)(unsafe.Add(mBase, uint32(v110+v116)))
	v119 = base.B2i32(base.Ui32(v118) < base.Ui32(v64))
	if base.Ui32(v118) < base.Ui32(v64) {
		goto L28
	} else {
		goto L29
	}
L28:
	;
	v120 = v91
	goto L30
L29:
	;
	v120 = v101
	goto L30
L30:
	;
	if base.Ui32(v118) < base.Ui32(v64) {
		goto L31
	} else {
		goto L32
	}
L31:
	;
	v125 = v101 + int32(1)
	goto L33
L32:
	;
	v125 = v90
	goto L33
L33:
	;
	if base.Ui32(v125&int32(65535)) < base.Ui32(v120&int32(65535)) {
		v90 = v125
		v91 = v120
		goto L23
	} else {
		goto L34
	}
L34:
	;
	goto L24
L35:
	;
	if v147 == int32(0) {
		goto L11
	} else {
		goto L36
	}
L36:
	;
	v152 = v40 + int32(1)
	if l4 != v152 {
		v40 = v152
		goto L9
	} else {
		goto L37
	}
L37:
	;
	goto L7
L38:
	;
	v158 = *(*int32)(unsafe.Add(mBase, uint32(l0)+48))
	*(*int32)(unsafe.Add(mBase, uint32(v13))) = v158 + int32(4)
	F_errmsg_internal(m, int32(730156), v13)
	mBase = m.M
	v164 = m.ExcPending
	if v164 != 0 {
		goto L1
	} else {
		goto L39
	}
L39:
	;
	F_errfinish(m, int32(512279), int32(358), int32(242047))
	mBase = m.M
	v169 = m.ExcPending
	if v169 != 0 {
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
	var v155 float64
	_ = v155
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
								switch v79 - int32(65530) {
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
									switch v79 - int32(65530) {
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
						v155 = float64(1e+100)
						if base.F64_gt(v153, v155) != 0 {
							v167 = v155
						} else {
							if base.Ui64(int64(9218868437227405312)) < base.Ui64(base.I64_reinterpret_f64(v153)&int64(9223372036854775807)) {
								v167 = v155
							} else {
								v163 = float64(1)
								if base.F64_le(v153, v163) != 0 {
									v167 = v163
								} else {
									v167 = base.F64_nearest(v153)
								}
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
										switch v79 - int32(65530) {
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
											switch v79 - int32(65530) {
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
								v155 = float64(1e+100)
								if base.F64_gt(v153, v155) != 0 {
									v167 = v155
								} else {
									if base.Ui64(int64(9218868437227405312)) < base.Ui64(base.I64_reinterpret_f64(v153)&int64(9223372036854775807)) {
										v167 = v155
									} else {
										v163 = float64(1)
										if base.F64_le(v153, v163) != 0 {
											v167 = v163
										} else {
											v167 = base.F64_nearest(v153)
										}
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
											switch v79 - int32(65530) {
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
												switch v79 - int32(65530) {
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
									v155 = float64(1e+100)
									if base.F64_gt(v153, v155) != 0 {
										v167 = v155
									} else {
										if base.Ui64(int64(9218868437227405312)) < base.Ui64(base.I64_reinterpret_f64(v153)&int64(9223372036854775807)) {
											v167 = v155
										} else {
											v163 = float64(1)
											if base.F64_le(v153, v163) != 0 {
												v167 = v163
											} else {
												v167 = base.F64_nearest(v153)
											}
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
		v29 = v20 + v14
		v30 = v24 + v29
		v31 = v28 + v30
		v35 = v29 - v28 ^ base.I32_rotl(v28, int32(16))
		v39 = v30 - v35 ^ base.I32_rotl(v35, int32(19))
		v44 = v35 + v31
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
	return base.I64_extend_i32_u(v309)<<(uint(int64(32))%64) | base.I64_extend_i32_u(v301^v309-base.I32_rotl(v309, int32(24)))
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
		F_appendStringInfo(m, l0, int32(488701), v10)
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
		F_appendStringInfo(m, l0, int32(493195), v10+int32(16))
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
		F_appendStringInfo(m, l0, int32(53039), v10+int32(32))
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
		F_appendStringInfo(m, l0, int32(522764), v10+int32(48))
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
		F_appendStringInfo(m, l0, int32(522823), v10-int32(-64))
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
		F_appendStringInfo(m, l0, int32(52481), v10+int32(80))
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
		F_appendStringInfo(m, l0, int32(521537), v10+int32(96))
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
		F_appendStringInfo(m, l0, int32(521509), v10+int32(112))
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
		F_appendStringInfo(m, l0, int32(521561), v10+int32(128))
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
		F_appendStringInfo(m, l0, int32(352168), v10+int32(144))
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
		F_appendStringInfo(m, l0, int32(521770), v10+int32(160))
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
	var v42 int32
	_ = v42
	var v44 int32
	_ = v44
	var v47 int32
	_ = v47
	var v49 int32
	_ = v49
	var v51 int32
	_ = v51
	var v54 int32
	_ = v54
	var v55 int32
	_ = v55
	var v57 int32
	_ = v57
	var v60 int32
	_ = v60
	var v62 int32
	_ = v62
	var v65 int32
	_ = v65
	var v76 int32
	_ = v76
	var v77 int32
	_ = v77
	var v80 int32
	_ = v80
	var v82 int32
	_ = v82
	var v83 int32
	_ = v83
	var v85 int32
	_ = v85
	var v88 int32
	_ = v88
	var v92 int32
	_ = v92
	var v94 int32
	_ = v94
	var v101 int32
	_ = v101
	var v102 int32
	_ = v102
	var v106 int32
	_ = v106
	var v108 int32
	_ = v108
	var v110 int32
	_ = v110
	var v115 int32
	_ = v115
	var v116 int32
	_ = v116
	var v122 int32
	_ = v122
	var v128 int32
	_ = v128
	var v132 int32
	_ = v132
	var v133 int32
	_ = v133
	var v137 int32
	_ = v137
	var v139 int32
	_ = v139
	var v143 int32
	_ = v143
	var v147 int32
	_ = v147
	var v148 int32
	_ = v148
	var v149 int32
	_ = v149
	var v150 int32
	_ = v150
	var v154 int32
	_ = v154
	var v158 int32
	_ = v158
	var v163 int32
	_ = v163
	var v165 int32
	_ = v165
	var v167 int32
	_ = v167
	var v169 int32
	_ = v169
	var v174 int32
	_ = v174
	var v175 int32
	_ = v175
	var v177 int32
	_ = v177
	var v178 int32
	_ = v178
	var v179 int32
	_ = v179
	var v181 int32
	_ = v181
	var v182 int32
	_ = v182
	var v183 int32
	_ = v183
	var v185 int32
	_ = v185
	var v186 int32
	_ = v186
	var v188 int32
	_ = v188
	var v190 int32
	_ = v190
	var v194 int32
	_ = v194
	var v195 int32
	_ = v195
	var v196 int32
	_ = v196
	var v197 int32
	_ = v197
	var v201 int32
	_ = v201
	var v205 int32
	_ = v205
	var v209 int32
	_ = v209
	var v210 int32
	_ = v210
	var v211 int32
	_ = v211
	var v212 int32
	_ = v212
	var v216 int32
	_ = v216
	var v217 int32
	_ = v217
	var v218 int32
	_ = v218
	var v220 int32
	_ = v220
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
	var v234 int32
	_ = v234
	var v235 int32
	_ = v235
	var v237 int32
	_ = v237
	var v238 int32
	_ = v238
	var v239 int32
	_ = v239
	var v241 int32
	_ = v241
	var v242 int32
	_ = v242
	var v243 int32
	_ = v243
	var v245 int32
	_ = v245
	var v246 int32
	_ = v246
	var v248 int32
	_ = v248
	var v250 int32
	_ = v250
	var v254 int32
	_ = v254
	var v255 int32
	_ = v255
	var v256 int32
	_ = v256
	var v257 int32
	_ = v257
	var v261 int32
	_ = v261
	var v265 int32
	_ = v265
	var v269 int32
	_ = v269
	var v270 int32
	_ = v270
	var v271 int32
	_ = v271
	var v272 int32
	_ = v272
	var v276 int32
	_ = v276
	var v277 int32
	_ = v277
	var v278 int32
	_ = v278
	var v280 int32
	_ = v280
	var v283 int32
	_ = v283
	var v284 int32
	_ = v284
	var v286 int32
	_ = v286
	var v287 int32
	_ = v287
	var v288 int32
	_ = v288
	var v292 int32
	_ = v292
	var v296 int32
	_ = v296
	var v297 int32
	_ = v297
	var v301 int32
	_ = v301
	var v302 int32
	_ = v302
	var v306 int32
	_ = v306
	var v307 int32
	_ = v307
	var v309 int32
	_ = v309
	var v311 int32
	_ = v311
	var v315 int32
	_ = v315
	var v316 int32
	_ = v316
	var v320 int32
	_ = v320
	var v321 int32
	_ = v321
	var v323 int32
	_ = v323
	var v324 int32
	_ = v324
	var v326 int32
	_ = v326
	var v330 int32
	_ = v330
	var v331 int32
	_ = v331
	var v335 int32
	_ = v335
	var v336 int32
	_ = v336
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
	var v352 int32
	_ = v352
	var v353 int32
	_ = v353
	var v357 int32
	_ = v357
	var v358 int32
	_ = v358
	var v359 int32
	_ = v359
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
	var v373 int32
	_ = v373
	var v374 int32
	_ = v374
	var v375 int32
	_ = v375
	var v379 int32
	_ = v379
	var v380 int32
	_ = v380
	var v381 int32
	_ = v381
	var v382 int32
	_ = v382
	var v386 int32
	_ = v386
	var v387 int32
	_ = v387
	var v388 int32
	_ = v388
	var v389 int32
	_ = v389
	var v393 int32
	_ = v393
	var v394 int32
	_ = v394
	var v395 int32
	_ = v395
	var v396 int32
	_ = v396
	var v401 int32
	_ = v401
	var v402 int32
	_ = v402
	var v403 int32
	_ = v403
	var v406 int32
	_ = v406
	var v408 int32
	_ = v408
	var v412 int32
	_ = v412
	var v416 int32
	_ = v416
	var v420 int32
	_ = v420
	var v424 int32
	_ = v424
	var v428 int32
	_ = v428
	var v438 int32
	_ = v438
	var v439 int32
	_ = v439
	var v440 int64
	_ = v440
	var v454 int64
	_ = v454
	var v455 int32
	_ = v455
	var v456 int32
	_ = v456
	var v471 int32
	_ = v471
	var v472 int32
	_ = v472
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
	v20 = int32(49152)
	if v19&v20 != v20 {
		goto L4
	} else {
		goto L5
	}
L3:
	;
	v471 = F_Int64GetDatum(m, v18-int64(1))
	mBase = m.M
	v472 = m.ExcPending
	if v472 != 0 {
		goto L1
	} else {
		goto L80
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
	v454 = v18
	goto L6
L6:
	;
	v455 = F_Int64GetDatum(m, v454)
	mBase = m.M
	v456 = m.ExcPending
	if v456 != 0 {
		goto L1
	} else {
		goto L79
	}
L7:
	;
	v41 = *(*int32)(unsafe.Add(mBase, uint32(v13)))
	v42 = int32(2)
	v44 = v40 + int32(base.Ui32(v41)>>(uint(v42)%32))
	if base.Ui32(v44) < base.Ui32(v42) {
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
	v47 = int32(0)
	v49 = v13 + int32(6)
	v51 = v13 + int32(8)
	if v24 < v47 {
		goto L12
	} else {
		goto L13
	}
L12:
	;
	v54 = v49
	goto L14
L13:
	;
	v54 = v51
	goto L14
L14:
	;
	v55 = int32(1)
	v57 = int32(base.Ui32(v44) >> (uint(v55) % 32))
	if base.Ui32(v57) <= base.Ui32(v55) {
		goto L15
	} else {
		goto L16
	}
L15:
	;
	v60 = v55
	goto L17
L16:
	;
	v60 = v57
	goto L17
L17:
	;
	v62 = v47
	v65 = v39
	goto L19
L18:
	;
	if v82 == v57 {
		goto L3
	} else {
		goto L23
	}
L19:
	;
	v76 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v54+v62<<(uint(int32(1))%32)))))
	if v76 != 0 {
		v82 = v62
		v83 = v65
		goto L18
	} else {
		goto L21
	}
L20:
	;
	v82 = v60
	v83 = v39 - v60
	goto L18
L21:
	;
	v77 = int32(1)
	v80 = v62 + v77
	if v80 != v60 {
		v62 = v80
		v65 = v65 - v77
		goto L19
	} else {
		goto L22
	}
L22:
	;
	goto L20
L23:
	;
	v85 = int32(1)
	if v57 <= v85 {
		goto L24
	} else {
		goto L25
	}
L24:
	;
	v88 = v85
	goto L26
L25:
	;
	v88 = v57
	goto L26
L26:
	;
	v92 = int32(0)
	v94 = v57
	goto L28
L27:
	;
	if int32(0) <= v24 {
		goto L32
	} else {
		goto L33
	}
L28:
	;
	v101 = int32(1)
	v102 = v94 - v101
	v106 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v54+v102<<(uint(v101)%32)))))
	if v106 != 0 {
		v110 = v92
		goto L27
	} else {
		goto L30
	}
L29:
	;
	v110 = v88
	goto L27
L30:
	;
	v108 = v92 + int32(1)
	if v108 != v88 {
		v92 = v108
		v94 = v102
		goto L28
	} else {
		goto L31
	}
L31:
	;
	goto L29
L32:
	;
	v115 = v51
	goto L34
L33:
	;
	v115 = v49
	goto L34
L34:
	;
	v116 = v82<<(uint(int32(1))%32) + v115
	v122 = (v44 - (v82+v110)<<(uint(int32(1))%32)) & int32(-2)
	v128 = v122 - int32(1636608432)
	if v18 == int64(0) {
		goto L37
	} else {
		goto L38
	}
L35:
	;
	v438 = F_Int64GetDatum(m, base.I64_extend_i32_u(v428)<<(uint(int64(32))%64)|base.I64_extend_i32_u(v420^v428-base.I32_rotl(v428, int32(24))))
	mBase = m.M
	v439 = m.ExcPending
	if v439 != 0 {
		goto L1
	} else {
		goto L78
	}
L36:
	;
	if v116&int32(3) != 0 {
		goto L52
	} else {
		goto L53
	}
L37:
	;
	v165 = v128
	v167 = v128
	v169 = v128
	goto L36
L38:
	;
	goto L39
L39:
	;
	v132 = v128 + base.I32_wrap_i64(v18)
	v133 = v132 + v128
	v137 = int32(4)
	v139 = base.I32_wrap_i64(int64(base.Ui64(v18)>>(uint(int64(32))%64))) ^ base.I32_rotl(v128, v137)
	v143 = v132 - v139 ^ base.I32_rotl(v139, int32(6))
	v147 = v133 - v143 ^ base.I32_rotl(v143, int32(8))
	v148 = v139 + v133
	v149 = v143 + v148
	v150 = v147 + v149
	v154 = v148 - v147 ^ base.I32_rotl(v147, int32(16))
	v158 = v149 - v154 ^ base.I32_rotl(v154, int32(19))
	v163 = v154 + v150
	v165 = v163
	v167 = v150 - v158 ^ base.I32_rotl(v158, v137)
	v169 = v158 + v163
	goto L36
L40:
	;
	v406 = int32(14)
	v408 = v402 ^ v403 - base.I32_rotl(v402, v406)
	v412 = v408 ^ v401 - base.I32_rotl(v408, int32(11))
	v416 = v412 ^ v402 - base.I32_rotl(v412, int32(25))
	v420 = v416 ^ v408 - base.I32_rotl(v416, int32(16))
	v424 = v420 ^ v412 - base.I32_rotl(v420, int32(4))
	v428 = v424 ^ v416 - base.I32_rotl(v424, v406)
	goto L35
L41:
	;
	v396 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v223))))
	v401 = v393 + v396
	v402 = v394
	v403 = v395
	goto L40
L42:
	;
	v389 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v223)+1)))
	v393 = v389<<(uint(int32(8))%32) + v386
	v394 = v387
	v395 = v388
	goto L41
L43:
	;
	v382 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v223)+2)))
	v386 = v382<<(uint(int32(16))%32) + v379
	v387 = v380
	v388 = v381
	goto L42
L44:
	;
	v375 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v223)+3)))
	v379 = v375<<(uint(int32(24))%32) + v226
	v380 = v373
	v381 = v374
	goto L43
L45:
	;
	v371 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v223)+4)))
	v373 = v369 + v371
	v374 = v370
	goto L44
L46:
	;
	v365 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v223)+5)))
	v369 = v365<<(uint(int32(8))%32) + v363
	v370 = v364
	goto L45
L47:
	;
	v359 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v223)+6)))
	v363 = v359<<(uint(int32(16))%32) + v357
	v364 = v358
	goto L46
L48:
	;
	v353 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v223)+7)))
	v357 = v353<<(uint(int32(24))%32) + v227
	v358 = v352
	goto L47
L49:
	;
	v348 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v223)+8)))
	v352 = v348<<(uint(int32(8))%32) + v347
	goto L48
L50:
	;
	v343 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v223)+9)))
	v347 = v343<<(uint(int32(16))%32) + v342
	goto L49
L51:
	;
	v338 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v223)+10)))
	v342 = v338<<(uint(int32(24))%32) + v228
	goto L50
L52:
	;
	if base.Ui32(int32(11)) < base.Ui32(v122) {
		goto L55
	} else {
		goto L56
	}
L53:
	;
	goto L54
L54:
	;
	if base.Ui32(int32(12)) <= base.Ui32(v122) {
		goto L61
	} else {
		goto L62
	}
L55:
	;
	v174 = v116
	v175 = v122
	v177 = v165
	v178 = v169
	v179 = v167
	goto L58
L56:
	;
	v223 = v116
	v224 = v122
	v226 = v165
	v227 = v169
	v228 = v167
	goto L57
L57:
	;
	switch v224 - int32(1) {
	case 0:
		v393 = v226
		v394 = v227
		v395 = v228
		goto L41
	case 1:
		v386 = v226
		v387 = v227
		v388 = v228
		goto L42
	case 2:
		v379 = v226
		v380 = v227
		v381 = v228
		goto L43
	case 3:
		v373 = v227
		v374 = v228
		goto L44
	case 4:
		v369 = v227
		v370 = v228
		goto L45
	case 5:
		v363 = v227
		v364 = v228
		goto L46
	case 6:
		v357 = v227
		v358 = v228
		goto L47
	case 7:
		v352 = v228
		goto L48
	case 8:
		v347 = v228
		goto L49
	case 9:
		v342 = v228
		goto L50
	case 10:
		goto L51
	default:
		v401 = v226
		v402 = v227
		v403 = v228
		goto L40
	}
L58:
	;
	v181 = *(*int32)(unsafe.Add(mBase, uint32(v174)+4))
	v182 = v181 + v178
	v183 = *(*int32)(unsafe.Add(mBase, uint32(v174)))
	v185 = *(*int32)(unsafe.Add(mBase, uint32(v174)+8))
	v186 = v185 + v179
	v188 = int32(4)
	v190 = v183 + v177 - v186 ^ base.I32_rotl(v186, v188)
	v194 = v182 - v190 ^ base.I32_rotl(v190, int32(6))
	v195 = v186 + v182
	v196 = v190 + v195
	v197 = v194 + v196
	v201 = v195 - v194 ^ base.I32_rotl(v194, int32(8))
	v205 = v196 - v201 ^ base.I32_rotl(v201, int32(16))
	v209 = v197 - v205 ^ base.I32_rotl(v205, int32(19))
	v210 = v201 + v197
	v211 = v205 + v210
	v212 = v209 + v211
	v216 = v210 - v209 ^ base.I32_rotl(v209, v188)
	v217 = int32(12)
	v218 = v174 + v217
	v220 = v175 - v217
	if base.Ui32(int32(11)) < base.Ui32(v220) {
		v174 = v218
		v175 = v220
		v177 = v211
		v178 = v212
		v179 = v216
		goto L58
	} else {
		goto L60
	}
L59:
	;
	v223 = v218
	v224 = v220
	v226 = v211
	v227 = v212
	v228 = v216
	goto L57
L60:
	;
	goto L59
L61:
	;
	v234 = v116
	v235 = v122
	v237 = v165
	v238 = v169
	v239 = v167
	goto L64
L62:
	;
	v283 = v116
	v284 = v122
	v286 = v165
	v287 = v169
	v288 = v167
	goto L63
L63:
	;
	switch v284 - int32(1) {
	case 0:
		v335 = v286
		goto L67
	case 1:
		v330 = v286
		goto L68
	case 2:
		goto L69
	case 3:
		v323 = v287
		goto L70
	case 4:
		v320 = v287
		goto L71
	case 5:
		v315 = v287
		goto L72
	case 6:
		goto L73
	case 7:
		v306 = v288
		goto L74
	case 8:
		v301 = v288
		goto L75
	case 9:
		v296 = v288
		goto L76
	case 10:
		goto L77
	default:
		v401 = v286
		v402 = v287
		v403 = v288
		goto L40
	}
L64:
	;
	v241 = *(*int32)(unsafe.Add(mBase, uint32(v234)+4))
	v242 = v241 + v238
	v243 = *(*int32)(unsafe.Add(mBase, uint32(v234)))
	v245 = *(*int32)(unsafe.Add(mBase, uint32(v234)+8))
	v246 = v245 + v239
	v248 = int32(4)
	v250 = v243 + v237 - v246 ^ base.I32_rotl(v246, v248)
	v254 = v242 - v250 ^ base.I32_rotl(v250, int32(6))
	v255 = v246 + v242
	v256 = v250 + v255
	v257 = v254 + v256
	v261 = v255 - v254 ^ base.I32_rotl(v254, int32(8))
	v265 = v256 - v261 ^ base.I32_rotl(v261, int32(16))
	v269 = v257 - v265 ^ base.I32_rotl(v265, int32(19))
	v270 = v261 + v257
	v271 = v265 + v270
	v272 = v269 + v271
	v276 = v270 - v269 ^ base.I32_rotl(v269, v248)
	v277 = int32(12)
	v278 = v234 + v277
	v280 = v235 - v277
	if base.Ui32(int32(11)) < base.Ui32(v280) {
		v234 = v278
		v235 = v280
		v237 = v271
		v238 = v272
		v239 = v276
		goto L64
	} else {
		goto L66
	}
L65:
	;
	v283 = v278
	v284 = v280
	v286 = v271
	v287 = v272
	v288 = v276
	goto L63
L66:
	;
	goto L65
L67:
	;
	v336 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v283))))
	v401 = v335 + v336
	v402 = v287
	v403 = v288
	goto L40
L68:
	;
	v331 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v283)+1)))
	v335 = v331<<(uint(int32(8))%32) + v330
	goto L67
L69:
	;
	v326 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v283)+2)))
	v330 = v326<<(uint(int32(16))%32) + v286
	goto L68
L70:
	;
	v324 = *(*int32)(unsafe.Add(mBase, uint32(v283)))
	v401 = v324 + v286
	v402 = v323
	v403 = v288
	goto L40
L71:
	;
	v321 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v283)+4)))
	v323 = v320 + v321
	goto L70
L72:
	;
	v316 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v283)+5)))
	v320 = v316<<(uint(int32(8))%32) + v315
	goto L71
L73:
	;
	v311 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v283)+6)))
	v315 = v311<<(uint(int32(16))%32) + v287
	goto L72
L74:
	;
	v307 = *(*int32)(unsafe.Add(mBase, uint32(v283)))
	v309 = *(*int32)(unsafe.Add(mBase, uint32(v283)+4))
	v401 = v307 + v286
	v402 = v309 + v287
	v403 = v306
	goto L40
L75:
	;
	v302 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v283)+8)))
	v306 = v302<<(uint(int32(8))%32) + v301
	goto L74
L76:
	;
	v297 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v283)+9)))
	v301 = v297<<(uint(int32(16))%32) + v296
	goto L75
L77:
	;
	v292 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v283)+10)))
	v296 = v292<<(uint(int32(24))%32) + v288
	goto L76
L78:
	;
	v440 = *(*int64)(unsafe.Add(mBase, uint32(v438)))
	v454 = v440 ^ base.I64_extend_i32_s(v83)
	goto L6
L79:
	;
	return v455
L80:
	;
	return v471
}
func F_hash_object_field_start(m *base.Module, l0 int32, l1 int32, l2 int32) int32 {
	mBase := m.M
	_ = mBase
	var v5 int32
	_ = v5
	var v6 int32
	_ = v6
	var v9 int32
	_ = v9
	var v14 int32
	_ = v14
	var v15 int32
	_ = v15
	v5 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v6 = *(*int32)(unsafe.Add(mBase, uint32(v5)+32))
	if v6 <= int32(1) {
		v9 = *(*int32)(unsafe.Add(mBase, uint32(v5)+28))
		*(*int32)(unsafe.Add(mBase, uint32(l0)+20)) = v9
		switch v9 - int32(3) {
		case 0, 2:
			v14 = *(*int32)(unsafe.Add(mBase, uint32(v5)+12))
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
							F_errmsg_internal(m, int32(46326), v7)
							mBase = m.M
							v64 = m.ExcPending
							if v64 != 0 {
								return int32(0)
							} else {
								F_errfinish(m, int32(512850), int32(855), int32(218447))
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
	var v4 int32
	_ = v4
	var v6 int32
	_ = v6
	var v10 int32
	_ = v10
	var v13 int32
	_ = v13
	var v14 int32
	_ = v14
	var v27 int32
	_ = v27
	var v28 int32
	_ = v28
	var v35 int32
	_ = v35
	var v37 int32
	_ = v37
	v4 = int32(1)
	v6 = int32(32)
	v10 = int32(1073741823)
	if v10 <= l0 {
		v13 = v10
	} else {
		v13 = l0
	}
	v14 = int32(1)
	if base.Ui32(v13) <= base.Ui32(v14) {
		v27 = v4
	} else {
		v27 = int32(base.Ui32(int32(-1)<<(uint(v6-base.I32_clz(v13-v14))%32)^int32(-1))>>(uint(int32(8))%32)) + v14
	}
	v28 = int32(1)
	if base.Ui32(v27) <= base.Ui32(v28) {
		v35 = v4
	} else {
		v35 = v4 << (uint(v6-base.I32_clz(v27-v28)) % 32)
	}
	v37 = int32(256)
	for {
		if v37 < v35 {
			v37 = v37 << (uint(int32(1)) % 32)
			continue
		} else {
			break
		}
		break
	}
	return v37
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
	var v26 int32
	_ = v26
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
	var v63 int32
	_ = v63
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
	var v160 int32
	_ = v160
	var v163 int32
	_ = v163
	var v164 int32
	_ = v164
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
	v26 = v2
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
	v36 = v33 + v26*int32(52)
	v37 = *(*int32)(unsafe.Add(mBase, uint32(v36)+16))
	v38 = *(*int32)(unsafe.Add(mBase, uint32(v36)))
	v39 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v14)+11)) = uint8(v39)
	*(*int32)(unsafe.Add(mBase, uint32(l0)+188)) = v26
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
	v63 = int32(0)
	goto L18
L16:
	;
	goto L17
L17:
	;
	v111 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v37)+4)))
	v113 = v111 & int32(65533)
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
	v82 = int32(*(*int16)(unsafe.Add(mBase, uint32(v78+v63<<(uint(v79)%32)))))
	v84 = v82 - v79
	v88 = *(*int32)(unsafe.Add(mBase, uint32(v77+v84<<(uint(v74)%32))))
	*(*int32)(unsafe.Add(mBase, uint32(v73+v63<<(uint(v74)%32)))) = v88
	v90 = *(*int32)(unsafe.Add(mBase, uint32(v37)+20))
	v92 = *(*int32)(unsafe.Add(mBase, uint32(v21)+20))
	v94 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v92+v84))))
	*(*uint8)(unsafe.Add(mBase, uint32(v90+v63))) = uint8(v94)
	v97 = v63 + v79
	v98 = *(*int32)(unsafe.Add(mBase, uint32(v36)+32))
	if v97 < v98 {
		v63 = v97
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
	*(*int32)(unsafe.Add(mBase, uint32(v19+v26<<(uint(int32(2))%32)))) = v160
	v163 = v26 + int32(1)
	v164 = *(*int32)(unsafe.Add(mBase, uint32(l0)+244))
	if v163 < v164 {
		v26 = v163
		goto L4
	} else {
		goto L38
	}
L23:
	;
	v160 = int32(0)
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
	v140 = v137 + v26*int32(24)
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
	v160 = v133 - v130
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
