package p5

import (
	base "github.com/shibukawa/pgmem/internal/aot/pgaot/base"
	"unsafe"
)

func F_CopySendTextLikeEndOfRow(m *base.Module, l0 int32) {
	mBase := m.M
	_ = mBase
	var v4 int32
	_ = v4
	var v5 int32
	_ = v5
	var v6 int32
	_ = v6
	var v9 int32
	_ = v9
	var v11 int32
	_ = v11
	var v13 int32
	_ = v13
	var v15 int32
	_ = v15
	var v16 int32
	_ = v16
	var v18 int32
	_ = v18
	var v20 int32
	_ = v20
	var v22 int32
	_ = v22
	var v25 int32
	_ = v25
	var v26 int32
	_ = v26
	var v27 int32
	_ = v27
	var v30 int32
	_ = v30
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
	var v43 int32
	_ = v43
	var v48 int32
	_ = v48
	var v49 int32
	_ = v49
	var v53 int32
	_ = v53
	var v55 int32
	_ = v55
	v4 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	switch v4 {
	case 0:
		v5 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
		v6 = *(*int32)(unsafe.Add(mBase, uint32(v5)+4))
		v9 = *(*int32)(unsafe.Add(mBase, uint32(v5)+8))
		if v9 <= v6+int32(1) {
			v49 = v5
			F_appendStringInfoChar(m, v49, int32(10))
			mBase = m.M
			v53 = m.ExcPending
			if v53 != 0 {
				return
			} else {
				F_CopySendEndOfRow(m, l0)
				mBase = m.M
				v55 = m.ExcPending
				if v55 != 0 {
					return
				} else {
					return
				}
			}
		} else {
			v11 = *(*int32)(unsafe.Add(mBase, uint32(v5)))
			v13 = int32(10)
			*(*uint8)(unsafe.Add(mBase, uint32(v11+v6))) = uint8(v13)
			v15 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
			v16 = *(*int32)(unsafe.Add(mBase, uint32(v15)+4))
			v18 = v16 + int32(1)
			*(*int32)(unsafe.Add(mBase, uint32(v15)+4)) = v18
			v20 = *(*int32)(unsafe.Add(mBase, uint32(v15)))
			v22 = int32(0)
			*(*uint8)(unsafe.Add(mBase, uint32(v20+v18))) = uint8(v22)
			F_CopySendEndOfRow(m, l0)
			mBase = m.M
			v25 = m.ExcPending
			if v25 != 0 {
				return
			} else {
				return
			}
		}
	case 1:
		v26 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
		v27 = *(*int32)(unsafe.Add(mBase, uint32(v26)+4))
		v30 = *(*int32)(unsafe.Add(mBase, uint32(v26)+8))
		if v30 <= v27+int32(1) {
			v49 = v26
			F_appendStringInfoChar(m, v49, int32(10))
			mBase = m.M
			v53 = m.ExcPending
			if v53 != 0 {
				return
			} else {
				F_CopySendEndOfRow(m, l0)
				mBase = m.M
				v55 = m.ExcPending
				if v55 != 0 {
					return
				} else {
					return
				}
			}
		} else {
			v32 = *(*int32)(unsafe.Add(mBase, uint32(v26)))
			v34 = int32(10)
			*(*uint8)(unsafe.Add(mBase, uint32(v32+v27))) = uint8(v34)
			v36 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
			v37 = *(*int32)(unsafe.Add(mBase, uint32(v36)+4))
			v39 = v37 + int32(1)
			*(*int32)(unsafe.Add(mBase, uint32(v36)+4)) = v39
			v41 = *(*int32)(unsafe.Add(mBase, uint32(v36)))
			v43 = int32(0)
			*(*uint8)(unsafe.Add(mBase, uint32(v41+v39))) = uint8(v43)
			F_CopySendEndOfRow(m, l0)
			mBase = m.M
			v48 = m.ExcPending
			if v48 != 0 {
				return
			} else {
				return
			}
		}
	default:
		F_CopySendEndOfRow(m, l0)
		mBase = m.M
		v48 = m.ExcPending
		if v48 != 0 {
			return
		} else {
			return
		}
	}
}
func F_assign_text_var(m *base.Module, l0 int32, l1 int32, l2 int32) {
	var v4 int32
	_ = v4
	var v5 int32
	_ = v5
	var v10 int32
	_ = v10
	v4 = F_cstring_to_text(m, l2)
	v5 = m.ExcPending
	if v5 != 0 {
		return
	} else {
		F_assign_simple_var(m, l0, l1, base.I64_extend_i32_u(v4), int32(0), int32(1))
		v10 = m.ExcPending
		if v10 != 0 {
			return
		} else {
			return
		}
	}
}
func F_text_concat_ws(m *base.Module, l0 int32) int64 {
	mBase := m.M
	_ = mBase
	var v6 int32
	_ = v6
	var v9 int32
	_ = v9
	var v13 int32
	_ = v13
	var v14 int32
	_ = v14
	var v17 int32
	_ = v17
	var v18 int32
	_ = v18
	var v19 int32
	_ = v19
	var v20 int32
	_ = v20
	var v26 int32
	_ = v26
	var v29 int32
	_ = v29
	var v36 int32
	_ = v36
	var v37 int32
	_ = v37
	var v43 int32
	_ = v43
	var v49 int32
	_ = v49
	var v52 int32
	_ = v52
	var v53 int32
	_ = v53
	var v54 int32
	_ = v54
	var v56 int32
	_ = v56
	var v59 int32
	_ = v59
	var v63 int32
	_ = v63
	var v67 int32
	_ = v67
	var v69 int32
	_ = v69
	var v70 int32
	_ = v70
	var v73 int32
	_ = v73
	v6 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+32)))
	if v6 == int32(1) {
		v9 = int32(1)
		*(*uint8)(unsafe.Add(mBase, uint32(l0)+16)) = uint8(v9)
		return int64(0)
	} else {
		v13 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
		v14 = F_pg_detoast_datum_packed(m, v13)
		mBase = m.M
		v17 = m.ExcPending
		if v17 != 0 {
			return int64(0)
		} else {
			v18 = F_pg_detoast_datum_packed(m, v14)
			mBase = m.M
			v19 = m.ExcPending
			if v19 != 0 {
				return int64(0)
			} else {
				v20 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v18))))
				if v20 == int32(1) {
					v26 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v18)+1)))
					if v26 == int32(18) {
						v29 = int32(16)
					} else {
						v29 = int32(0)
					}
					if base.Ui32((v26-int32(1))&int32(255)) < base.Ui32(int32(3)) {
						v36 = int32(4)
					} else {
						v36 = v29
					}
					v49 = v36
				} else {
					v37 = int32(1)
					if v20&v37 != 0 {
						v49 = int32(base.Ui32(v20)>>(uint(v37)%32)) - v37
					} else {
						v43 = *(*int32)(unsafe.Add(mBase, uint32(v18)))
						v49 = int32(base.Ui32(v43)>>(uint(int32(2))%32)) - int32(4)
					}
				}
				v52 = F_palloc(m, v49+int32(1))
				mBase = m.M
				v53 = m.ExcPending
				if v53 != 0 {
					return int64(0)
				} else {
					if v49 != 0 {
						v54 = int32(1)
						v56 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v18))))
						if v56&v54 != 0 {
							v59 = v54
						} else {
							v59 = int32(4)
						}
						base.MemoryCopy(m, v52, v18+v59, v49)
					} else {
					}
					v63 = int32(0)
					*(*uint8)(unsafe.Add(mBase, uint32(v49+v52))) = uint8(v63)
					if v18 != v14 {
						F_pfree(m, v18)
						mBase = m.M
						v67 = m.ExcPending
						if v67 != 0 {
							return int64(0)
						} else {
							v69 = F_concat_internal(m, v52, int32(1), l0)
							mBase = m.M
							v70 = m.ExcPending
							if v70 != 0 {
								return int64(0)
							} else {
								if v69 == int32(0) {
									v73 = int32(1)
									*(*uint8)(unsafe.Add(mBase, uint32(l0)+16)) = uint8(v73)
									return int64(0)
								} else {
									return base.I64_extend_i32_u(v69)
								}
							}
						}
					} else {
						v69 = F_concat_internal(m, v52, int32(1), l0)
						mBase = m.M
						v70 = m.ExcPending
						if v70 != 0 {
							return int64(0)
						} else {
							if v69 == int32(0) {
								v73 = int32(1)
								*(*uint8)(unsafe.Add(mBase, uint32(l0)+16)) = uint8(v73)
								return int64(0)
							} else {
								return base.I64_extend_i32_u(v69)
							}
						}
					}
				}
			}
		}
	}
}
func F_text_format(m *base.Module, l0 int32) int64 {
	mBase := m.M
	_ = mBase
	var v2 int32
	_ = v2
	var v18 int32
	_ = v18
	var v20 int32
	_ = v20
	var v26 int32
	_ = v26
	var v29 int32
	_ = v29
	var v33 int32
	_ = v33
	var v34 int32
	_ = v34
	var v35 int32
	_ = v35
	var v38 int32
	_ = v38
	var v41 int32
	_ = v41
	var v44 int32
	_ = v44
	var v46 int32
	_ = v46
	var v48 int32
	_ = v48
	var v50 int32
	_ = v50
	var v51 int32
	_ = v51
	var v52 int32
	_ = v52
	var v55 int32
	_ = v55
	var v56 int32
	_ = v56
	var v64 int32
	_ = v64
	var v65 int32
	_ = v65
	var v66 int32
	_ = v66
	var v67 int32
	_ = v67
	var v75 int32
	_ = v75
	var v76 int32
	_ = v76
	var v79 int32
	_ = v79
	var v81 int32
	_ = v81
	var v82 int32
	_ = v82
	var v83 int32
	_ = v83
	var v84 int32
	_ = v84
	var v85 int32
	_ = v85
	var v86 int32
	_ = v86
	var v88 int32
	_ = v88
	var v90 int32
	_ = v90
	var v91 int32
	_ = v91
	var v92 int32
	_ = v92
	var v95 int32
	_ = v95
	var v106 int32
	_ = v106
	var v109 int32
	_ = v109
	var v118 int32
	_ = v118
	var v119 int32
	_ = v119
	var v125 int32
	_ = v125
	var v129 int32
	_ = v129
	var v134 int32
	_ = v134
	var v140 int32
	_ = v140
	var v144 int32
	_ = v144
	var v146 int32
	_ = v146
	var v151 int32
	_ = v151
	var v153 int32
	_ = v153
	var v156 int32
	_ = v156
	var v157 int32
	_ = v157
	var v164 int32
	_ = v164
	var v166 int32
	_ = v166
	var v168 int32
	_ = v168
	var v171 int32
	_ = v171
	var v172 int32
	_ = v172
	var v180 int32
	_ = v180
	var v181 int32
	_ = v181
	var v183 int32
	_ = v183
	var v185 int32
	_ = v185
	var v187 int32
	_ = v187
	var v189 int32
	_ = v189
	var v191 int32
	_ = v191
	var v199 int32
	_ = v199
	var v200 int32
	_ = v200
	var v201 int32
	_ = v201
	var v202 int32
	_ = v202
	var v203 int32
	_ = v203
	var v207 int32
	_ = v207
	var v212 int32
	_ = v212
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
	var v227 int32
	_ = v227
	var v240 int32
	_ = v240
	var v243 int32
	_ = v243
	var v247 int32
	_ = v247
	var v269 int32
	_ = v269
	var v272 int32
	_ = v272
	var v276 int32
	_ = v276
	var v280 int32
	_ = v280
	var v285 int32
	_ = v285
	var v290 int32
	_ = v290
	var v295 int32
	_ = v295
	var v304 int32
	_ = v304
	var v307 int32
	_ = v307
	var v312 int32
	_ = v312
	var v313 int32
	_ = v313
	var v314 int32
	_ = v314
	var v318 int32
	_ = v318
	var v321 int32
	_ = v321
	var v325 int32
	_ = v325
	var v327 int32
	_ = v327
	var v330 int32
	_ = v330
	var v332 int32
	_ = v332
	var v334 int32
	_ = v334
	var v336 int32
	_ = v336
	var v341 int32
	_ = v341
	var v344 int32
	_ = v344
	var v348 int32
	_ = v348
	var v352 int32
	_ = v352
	var v357 int32
	_ = v357
	var v361 int32
	_ = v361
	var v364 int32
	_ = v364
	var v368 int32
	_ = v368
	var v373 int32
	_ = v373
	var v377 int32
	_ = v377
	var v380 int32
	_ = v380
	var v384 int32
	_ = v384
	var v388 int32
	_ = v388
	var v393 int32
	_ = v393
	var v397 int32
	_ = v397
	var v400 int32
	_ = v400
	var v404 int32
	_ = v404
	var v408 int32
	_ = v408
	var v413 int32
	_ = v413
	var v417 int32
	_ = v417
	var v420 int32
	_ = v420
	var v424 int32
	_ = v424
	var v429 int32
	_ = v429
	var v433 int32
	_ = v433
	var v436 int32
	_ = v436
	var v440 int32
	_ = v440
	var v445 int32
	_ = v445
	var v449 int32
	_ = v449
	var v452 int32
	_ = v452
	var v456 int32
	_ = v456
	var v460 int32
	_ = v460
	var v465 int32
	_ = v465
	var v470 int32
	_ = v470
	var v475 int32
	_ = v475
	var v488 int32
	_ = v488
	var v489 int32
	_ = v489
	var v490 int32
	_ = v490
	var v492 int32
	_ = v492
	var v494 int32
	_ = v494
	var v498 int32
	_ = v498
	var v499 int32
	_ = v499
	var v500 int32
	_ = v500
	var v502 int32
	_ = v502
	var v504 int32
	_ = v504
	var v512 int32
	_ = v512
	var v513 int32
	_ = v513
	var v514 int32
	_ = v514
	var v553 int32
	_ = v553
	var v554 int32
	_ = v554
	var v563 int32
	_ = v563
	var v565 int32
	_ = v565
	var v569 int32
	_ = v569
	var v570 int32
	_ = v570
	var v573 int32
	_ = v573
	var v577 int32
	_ = v577
	var v578 int32
	_ = v578
	var v580 int32
	_ = v580
	var v583 int32
	_ = v583
	var v585 int32
	_ = v585
	var v590 int32
	_ = v590
	var v592 int32
	_ = v592
	var v597 int32
	_ = v597
	var v599 int32
	_ = v599
	var v602 int32
	_ = v602
	var v604 int32
	_ = v604
	var v607 int32
	_ = v607
	var v619 int32
	_ = v619
	var v624 int32
	_ = v624
	var v630 int32
	_ = v630
	var v631 int32
	_ = v631
	var v632 int64
	_ = v632
	var v633 int32
	_ = v633
	var v634 int32
	_ = v634
	var v635 int32
	_ = v635
	var v637 int32
	_ = v637
	var v638 int32
	_ = v638
	var v640 int32
	_ = v640
	var v641 int32
	_ = v641
	var v645 int64
	_ = v645
	var v647 int32
	_ = v647
	var v648 int64
	_ = v648
	var v649 int32
	_ = v649
	var v652 int32
	_ = v652
	var v653 int32
	_ = v653
	var v668 int32
	_ = v668
	var v669 int32
	_ = v669
	var v673 int32
	_ = v673
	var v674 int32
	_ = v674
	var v677 int32
	_ = v677
	var v678 int32
	_ = v678
	var v679 int32
	_ = v679
	var v680 int32
	_ = v680
	var v682 int32
	_ = v682
	var v684 int32
	_ = v684
	var v686 int32
	_ = v686
	var v687 int32
	_ = v687
	var v691 int32
	_ = v691
	var v697 int32
	_ = v697
	var v698 int32
	_ = v698
	var v699 int64
	_ = v699
	var v700 int32
	_ = v700
	var v701 int32
	_ = v701
	var v702 int32
	_ = v702
	var v704 int32
	_ = v704
	var v705 int32
	_ = v705
	var v707 int32
	_ = v707
	var v708 int32
	_ = v708
	var v712 int64
	_ = v712
	var v714 int32
	_ = v714
	var v715 int64
	_ = v715
	var v716 int32
	_ = v716
	var v725 int32
	_ = v725
	var v726 int32
	_ = v726
	var v730 int32
	_ = v730
	var v731 int32
	_ = v731
	var v732 int32
	_ = v732
	var v734 int32
	_ = v734
	var v737 int32
	_ = v737
	var v738 int32
	_ = v738
	var v747 int32
	_ = v747
	var v752 int32
	_ = v752
	var v756 int32
	_ = v756
	var v759 int32
	_ = v759
	var v763 int32
	_ = v763
	var v768 int32
	_ = v768
	var v771 int32
	_ = v771
	var v772 int32
	_ = v772
	var v775 int32
	_ = v775
	var v776 int32
	_ = v776
	var v778 int32
	_ = v778
	var v780 int32
	_ = v780
	var v783 int32
	_ = v783
	var v784 int32
	_ = v784
	var v786 int32
	_ = v786
	var v788 int32
	_ = v788
	var v790 int32
	_ = v790
	var v794 int32
	_ = v794
	var v796 int32
	_ = v796
	var v800 int32
	_ = v800
	var v803 int32
	_ = v803
	var v804 int32
	_ = v804
	var v805 int32
	_ = v805
	var v812 int32
	_ = v812
	var v816 int32
	_ = v816
	var v821 int32
	_ = v821
	var v826 int32
	_ = v826
	var v830 int32
	_ = v830
	var v832 int32
	_ = v832
	var v837 int32
	_ = v837
	var v840 int32
	_ = v840
	var v859 int32
	_ = v859
	var v861 int32
	_ = v861
	var v862 int32
	_ = v862
	var v864 int32
	_ = v864
	var v865 int32
	_ = v865
	var v866 int32
	_ = v866
	var v868 int32
	_ = v868
	var v869 int32
	_ = v869
	var v870 int32
	_ = v870
	var v877 int32
	_ = v877
	var v879 int32
	_ = v879
	var v898 int64
	_ = v898
	var v906 int32
	_ = v906
	var v909 int32
	_ = v909
	var v910 int32
	_ = v910
	var v911 int32
	_ = v911
	var v916 int32
	_ = v916
	var v920 int32
	_ = v920
	var v925 int32
	_ = v925
	var v929 int32
	_ = v929
	var v932 int32
	_ = v932
	var v936 int32
	_ = v936
	var v941 int32
	_ = v941
	var v945 int32
	_ = v945
	var v949 int32
	_ = v949
	var v954 int32
	_ = v954
	var v958 int32
	_ = v958
	var v961 int32
	_ = v961
	var v965 int32
	_ = v965
	var v970 int32
	_ = v970
	var v974 int32
	_ = v974
	var v978 int32
	_ = v978
	var v983 int32
	_ = v983
	v2 = int32(0)
	v18 = m.G0
	v20 = v18 - int32(112)
	m.G0 = v20
	*(*int32)(unsafe.Add(mBase, uint32(v20)+84)) = v2
	*(*int32)(unsafe.Add(mBase, uint32(v20)+80)) = v2
	v26 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+32)))
	if v26 == int32(1) {
		goto L7
	} else {
		goto L8
	}
L1:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v974 = m.ExcPending
	if v974 != 0 {
		goto L20
	} else {
		goto L272
	}
L2:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v958 = m.ExcPending
	if v958 != 0 {
		goto L20
	} else {
		goto L268
	}
L3:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v945 = m.ExcPending
	if v945 != 0 {
		goto L20
	} else {
		goto L265
	}
L4:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v929 = m.ExcPending
	if v929 != 0 {
		goto L20
	} else {
		goto L261
	}
L5:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v906 = m.ExcPending
	if v906 != 0 {
		goto L20
	} else {
		goto L255
	}
L6:
	;
	m.G0 = v20 + int32(112)
	return v898
L7:
	;
	v29 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(l0)+16)) = uint8(v29)
	v898 = int64(0)
	goto L6
L8:
	;
	goto L9
L9:
	;
	v33 = l0 + int32(24)
	v34 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v35 = int32(0)
	if v34 == v35 {
		v46 = v35
		goto L12
	} else {
		goto L13
	}
L10:
	;
	v83 = *(*int32)(unsafe.Add(mBase, uint32(v33)))
	v84 = F_pg_detoast_datum_packed(m, v83)
	mBase = m.M
	v85 = m.ExcPending
	if v85 != 0 {
		goto L20
	} else {
		goto L24
	}
L11:
	;
	if v48 != 0 {
		goto L16
	} else {
		goto L17
	}
L12:
	;
	v48 = v46 & int32(1)
	goto L11
L13:
	;
	v38 = *(*int32)(unsafe.Add(mBase, uint32(v34)+24))
	if v38 == int32(0) {
		v46 = v35
		goto L12
	} else {
		goto L14
	}
L14:
	;
	v41 = *(*int32)(unsafe.Add(mBase, uint32(v38)))
	if v41 != int32(15) {
		v46 = v35
		goto L12
	} else {
		goto L15
	}
L15:
	;
	v44 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v38)+13)))
	v46 = v44
	goto L12
L16:
	;
	v50 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+48)))
	if v50 != 0 {
		v81 = v2
		v82 = int32(1)
		goto L10
	} else {
		goto L19
	}
L17:
	;
	goto L18
L18:
	;
	v79 = int32(*(*int16)(unsafe.Add(mBase, uint32(l0)+18)))
	v81 = v2
	v82 = v79
	goto L10
L19:
	;
	v51 = *(*int32)(unsafe.Add(mBase, uint32(l0)+40))
	v52 = F_pg_detoast_datum(m, v51)
	mBase = m.M
	v55 = m.ExcPending
	if v55 != 0 {
		goto L20
	} else {
		goto L21
	}
L20:
	;
	return int64(0)
L21:
	;
	v56 = *(*int32)(unsafe.Add(mBase, uint32(v52)+12))
	F_get_typlenbyvalalign(m, v56, v20+int32(24), v20+int32(88), v20+int32(108))
	mBase = m.M
	v64 = m.ExcPending
	if v64 != 0 {
		goto L20
	} else {
		goto L22
	}
L22:
	;
	v65 = int32(*(*int16)(unsafe.Add(mBase, uint32(v20)+24)))
	v66 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v20)+88)))
	v67 = int32(*(*int8)(unsafe.Add(mBase, uint32(v20)+108)))
	F_deconstruct_array(m, v52, v65, v66, v67, v20+int32(84), v20+int32(80), v20+int32(52))
	mBase = m.M
	v75 = m.ExcPending
	if v75 != 0 {
		goto L20
	} else {
		goto L23
	}
L23:
	;
	v76 = *(*int32)(unsafe.Add(mBase, uint32(v20)+52))
	v81 = v56
	v82 = v76 + int32(1)
	goto L10
L24:
	;
	v86 = int32(1)
	v88 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v84))))
	v90 = v88 & v86
	if v90 != 0 {
		goto L25
	} else {
		goto L26
	}
L25:
	;
	v91 = v86
	goto L27
L26:
	;
	v91 = int32(4)
	goto L27
L27:
	;
	v92 = v84 + v91
	if v88 == int32(1) {
		goto L33
	} else {
		goto L34
	}
L28:
	;
	v859 = *(*int32)(unsafe.Add(mBase, uint32(v20)+84))
	if v859 != 0 {
		goto L242
	} else {
		goto L243
	}
L29:
	;
	v140 = v92
	v144 = int32(1)
	v146 = v2
	v151 = v2
	goto L44
L30:
	;
	F_initStringInfo(m, v20+int32(88))
	mBase = m.M
	v129 = m.ExcPending
	if v129 != 0 {
		goto L20
	} else {
		goto L42
	}
L31:
	;
	v119 = *(*int32)(unsafe.Add(mBase, uint32(v84)))
	v125 = int32(base.Ui32(v119)>>(uint(int32(2))%32)) - int32(4)
	goto L30
L32:
	;
	F_initStringInfo(m, v20+int32(88))
	mBase = m.M
	v118 = m.ExcPending
	if v118 != 0 {
		goto L20
	} else {
		goto L41
	}
L33:
	;
	v95 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v84)+1)))
	if base.Ui32((v95-int32(1))&int32(255)) < base.Ui32(int32(3)) {
		goto L32
	} else {
		goto L36
	}
L34:
	;
	goto L35
L35:
	;
	if v90 == int32(0) {
		goto L31
	} else {
		goto L40
	}
L36:
	;
	if v95 == int32(18) {
		goto L37
	} else {
		goto L38
	}
L37:
	;
	v106 = int32(16)
	goto L39
L38:
	;
	v106 = int32(0)
	goto L39
L39:
	;
	v125 = v106
	goto L30
L40:
	;
	v109 = int32(1)
	v125 = int32(base.Ui32(v88)>>(uint(v109)%32)) - v109
	goto L30
L41:
	;
	v134 = v92 + int32(4)
	goto L29
L42:
	;
	if v125 == int32(0) {
		goto L28
	} else {
		goto L43
	}
L43:
	;
	v134 = v125 + v92
	goto L29
L44:
	;
	v153 = int32(*(*int8)(unsafe.Add(mBase, uint32(v140))))
	if v153 != int32(37) {
		goto L57
	} else {
		goto L58
	}
L45:
	;
	goto L28
L46:
	;
	v840 = v826 + int32(1)
	if base.Ui32(v840) < base.Ui32(v134) {
		v140 = v840
		v144 = v830
		v146 = v832
		v151 = v837
		goto L44
	} else {
		goto L241
	}
L47:
	;
	v512 = int32(_a_F_text_format_0)
	v513 = int32(*(*int8)(unsafe.Add(mBase, uint32(v499))))
	v514 = int32(4)
	goto L143
L48:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v20)+108)) = v470
	v488 = F_text_format_parse_digits(m, v20+int32(108), v134, v20+int32(104))
	mBase = m.M
	v489 = m.ExcPending
	if v489 != 0 {
		goto L20
	} else {
		goto L136
	}
L49:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v449 = m.ExcPending
	if v449 != 0 {
		goto L20
	} else {
		goto L131
	}
L50:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v433 = m.ExcPending
	if v433 != 0 {
		goto L20
	} else {
		goto L127
	}
L51:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v417 = m.ExcPending
	if v417 != 0 {
		goto L20
	} else {
		goto L123
	}
L52:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v397 = m.ExcPending
	if v397 != 0 {
		goto L20
	} else {
		goto L118
	}
L53:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v377 = m.ExcPending
	if v377 != 0 {
		goto L20
	} else {
		goto L113
	}
L54:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v361 = m.ExcPending
	if v361 != 0 {
		goto L20
	} else {
		goto L109
	}
L55:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v341 = m.ExcPending
	if v341 != 0 {
		goto L20
	} else {
		goto L104
	}
L56:
	;
	v327 = *(*int32)(unsafe.Add(mBase, uint32(v20)+88))
	*(*uint8)(unsafe.Add(mBase, uint32(v327+v157))) = uint8(v153)
	v330 = *(*int32)(unsafe.Add(mBase, uint32(v20)+92))
	v332 = v330 + int32(1)
	*(*int32)(unsafe.Add(mBase, uint32(v20)+92)) = v332
	v334 = *(*int32)(unsafe.Add(mBase, uint32(v20)+88))
	v336 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v334+v332))) = uint8(v336)
	v826 = v140
	v830 = v144
	v832 = v146
	v837 = v151
	goto L46
L57:
	;
	v156 = *(*int32)(unsafe.Add(mBase, uint32(v20)+96))
	v157 = *(*int32)(unsafe.Add(mBase, uint32(v20)+92))
	if v157+int32(1) < v156 {
		goto L56
	} else {
		goto L60
	}
L58:
	;
	goto L59
L59:
	;
	v166 = v140 + int32(1)
	if base.Ui32(v134) <= base.Ui32(v166) {
		goto L55
	} else {
		goto L62
	}
L60:
	;
	F_appendStringInfoChar(m, v20+int32(88), v153)
	mBase = m.M
	v164 = m.ExcPending
	if v164 != 0 {
		goto L20
	} else {
		goto L61
	}
L61:
	;
	v826 = v140
	v830 = v144
	v832 = v146
	v837 = v151
	goto L46
L62:
	;
	v168 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v166))))
	if v168 == int32(37) {
		goto L63
	} else {
		goto L64
	}
L63:
	;
	v171 = *(*int32)(unsafe.Add(mBase, uint32(v20)+96))
	v172 = *(*int32)(unsafe.Add(mBase, uint32(v20)+92))
	if v171 <= v172+int32(1) {
		goto L66
	} else {
		goto L67
	}
L64:
	;
	goto L65
L65:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v20)+108)) = v166
	v199 = F_text_format_parse_digits(m, v20+int32(108), v134, v20+int32(104))
	mBase = m.M
	v200 = m.ExcPending
	if v200 != 0 {
		goto L20
	} else {
		goto L70
	}
L66:
	;
	F_appendStringInfoChar(m, v20+int32(88), int32(37))
	mBase = m.M
	v180 = m.ExcPending
	if v180 != 0 {
		goto L20
	} else {
		goto L69
	}
L67:
	;
	goto L68
L68:
	;
	v181 = *(*int32)(unsafe.Add(mBase, uint32(v20)+88))
	v183 = int32(37)
	*(*uint8)(unsafe.Add(mBase, uint32(v181+v172))) = uint8(v183)
	v185 = *(*int32)(unsafe.Add(mBase, uint32(v20)+92))
	v187 = v185 + int32(1)
	*(*int32)(unsafe.Add(mBase, uint32(v20)+92)) = v187
	v189 = *(*int32)(unsafe.Add(mBase, uint32(v20)+88))
	v191 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v189+v187))) = uint8(v191)
	v826 = v166
	v830 = v144
	v832 = v146
	v837 = v151
	goto L46
L69:
	;
	v826 = v166
	v830 = v144
	v832 = v146
	v837 = v151
	goto L46
L70:
	;
	v201 = *(*int32)(unsafe.Add(mBase, uint32(v20)+108))
	if v199 != 0 {
		goto L71
	} else {
		goto L72
	}
L71:
	;
	v202 = *(*int32)(unsafe.Add(mBase, uint32(v20)+104))
	v203 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v201))))
	if v203 != int32(36) {
		goto L74
	} else {
		goto L75
	}
L72:
	;
	v214 = int32(-1)
	v215 = v201
	goto L73
L73:
	;
	v216 = int32(0)
	v217 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v215))))
	switch v217 - int32(42) {
	case 0:
		v290 = v215
		v295 = v216
		goto L79
	default:
		v470 = v215
		v475 = v216
		goto L48
	case 3:
		goto L80
	}
L74:
	;
	v207 = int32(-1)
	v498 = v207
	v499 = v201
	v500 = v202
	v502 = v207
	v504 = int32(0)
	goto L47
L75:
	;
	goto L76
L76:
	;
	if v202 == int32(0) {
		goto L54
	} else {
		goto L77
	}
L77:
	;
	v212 = v201 + int32(1)
	if base.Ui32(v134) <= base.Ui32(v212) {
		goto L53
	} else {
		goto L78
	}
L78:
	;
	v214 = v202
	v215 = v212
	goto L73
L79:
	;
	v304 = v290 + int32(1)
	*(*int32)(unsafe.Add(mBase, uint32(v20)+108)) = v304
	if base.Ui32(v134) <= base.Ui32(v304) {
		goto L52
	} else {
		goto L96
	}
L80:
	;
	v221 = v215 + int32(1)
	if base.Ui32(v221) < base.Ui32(v134) {
		goto L81
	} else {
		goto L82
	}
L81:
	;
	v227 = v221
	goto L84
L82:
	;
	goto L83
L83:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v269 = m.ExcPending
	if v269 != 0 {
		goto L20
	} else {
		goto L91
	}
L84:
	;
	v240 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v227))))
	if v240 != int32(45) {
		goto L86
	} else {
		goto L87
	}
L85:
	;
	goto L83
L86:
	;
	v243 = int32(1)
	if v240 == int32(42) {
		v290 = v227
		v295 = v243
		goto L79
	} else {
		goto L89
	}
L87:
	;
	goto L88
L88:
	;
	v247 = v227 + int32(1)
	if base.Ui32(v247) < base.Ui32(v134) {
		v227 = v247
		goto L84
	} else {
		goto L90
	}
L89:
	;
	v470 = v227
	v475 = v243
	goto L48
L90:
	;
	goto L85
L91:
	;
	F_errcode(m, int32(50856066))
	mBase = m.M
	v272 = m.ExcPending
	if v272 != 0 {
		goto L20
	} else {
		goto L92
	}
L92:
	;
	F_errmsg(m, int32(_a_F_text_format_1), int32(0))
	mBase = m.M
	v276 = m.ExcPending
	if v276 != 0 {
		goto L20
	} else {
		goto L93
	}
L93:
	;
	F_errhint(m, int32(_a_F_text_format_2), int32(0))
	mBase = m.M
	v280 = m.ExcPending
	if v280 != 0 {
		goto L20
	} else {
		goto L94
	}
L94:
	;
	F_errfinish(m, int32(_a_F_text_format_3), int32(_a_F_text_format_4), int32(_a_F_text_format_5))
	mBase = m.M
	v285 = m.ExcPending
	if v285 != 0 {
		goto L20
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
	v307 = int32(0)
	v312 = F_text_format_parse_digits(m, v20+int32(108), v134, v20+int32(104))
	mBase = m.M
	v313 = m.ExcPending
	if v313 != 0 {
		goto L20
	} else {
		goto L97
	}
L97:
	;
	v314 = *(*int32)(unsafe.Add(mBase, uint32(v20)+108))
	if v312 == int32(0) {
		goto L98
	} else {
		goto L99
	}
L98:
	;
	v498 = v214
	v499 = v314
	v500 = v307
	v502 = int32(0)
	v504 = v295
	goto L47
L99:
	;
	goto L100
L100:
	;
	v318 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v314))))
	if v318 != int32(36) {
		goto L51
	} else {
		goto L101
	}
L101:
	;
	v321 = *(*int32)(unsafe.Add(mBase, uint32(v20)+104))
	if v321 == int32(0) {
		goto L50
	} else {
		goto L102
	}
L102:
	;
	v325 = v314 + int32(1)
	if base.Ui32(v134) <= base.Ui32(v325) {
		goto L49
	} else {
		goto L103
	}
L103:
	;
	v498 = v214
	v499 = v325
	v500 = v307
	v502 = v321
	v504 = v295
	goto L47
L104:
	;
	F_errcode(m, int32(50856066))
	mBase = m.M
	v344 = m.ExcPending
	if v344 != 0 {
		goto L20
	} else {
		goto L105
	}
L105:
	;
	F_errmsg(m, int32(_a_F_text_format_1), int32(0))
	mBase = m.M
	v348 = m.ExcPending
	if v348 != 0 {
		goto L20
	} else {
		goto L106
	}
L106:
	;
	F_errhint(m, int32(_a_F_text_format_2), int32(0))
	mBase = m.M
	v352 = m.ExcPending
	if v352 != 0 {
		goto L20
	} else {
		goto L107
	}
L107:
	;
	F_errfinish(m, int32(_a_F_text_format_3), int32(_a_F_text_format_6), int32(_a_F_text_format_7))
	mBase = m.M
	v357 = m.ExcPending
	if v357 != 0 {
		goto L20
	} else {
		goto L108
	}
L108:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L109:
	;
	F_errcode(m, int32(50856066))
	mBase = m.M
	v364 = m.ExcPending
	if v364 != 0 {
		goto L20
	} else {
		goto L110
	}
L110:
	;
	F_errmsg(m, int32(_a_F_text_format_8), int32(0))
	mBase = m.M
	v368 = m.ExcPending
	if v368 != 0 {
		goto L20
	} else {
		goto L111
	}
L111:
	;
	F_errfinish(m, int32(_a_F_text_format_3), int32(_a_F_text_format_9), int32(_a_F_text_format_5))
	mBase = m.M
	v373 = m.ExcPending
	if v373 != 0 {
		goto L20
	} else {
		goto L112
	}
L112:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L113:
	;
	F_errcode(m, int32(50856066))
	mBase = m.M
	v380 = m.ExcPending
	if v380 != 0 {
		goto L20
	} else {
		goto L114
	}
L114:
	;
	F_errmsg(m, int32(_a_F_text_format_1), int32(0))
	mBase = m.M
	v384 = m.ExcPending
	if v384 != 0 {
		goto L20
	} else {
		goto L115
	}
L115:
	;
	F_errhint(m, int32(_a_F_text_format_2), int32(0))
	mBase = m.M
	v388 = m.ExcPending
	if v388 != 0 {
		goto L20
	} else {
		goto L116
	}
L116:
	;
	F_errfinish(m, int32(_a_F_text_format_3), int32(_a_F_text_format_10), int32(_a_F_text_format_5))
	mBase = m.M
	v393 = m.ExcPending
	if v393 != 0 {
		goto L20
	} else {
		goto L117
	}
L117:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L118:
	;
	F_errcode(m, int32(50856066))
	mBase = m.M
	v400 = m.ExcPending
	if v400 != 0 {
		goto L20
	} else {
		goto L119
	}
L119:
	;
	F_errmsg(m, int32(_a_F_text_format_1), int32(0))
	mBase = m.M
	v404 = m.ExcPending
	if v404 != 0 {
		goto L20
	} else {
		goto L120
	}
L120:
	;
	F_errhint(m, int32(_a_F_text_format_2), int32(0))
	mBase = m.M
	v408 = m.ExcPending
	if v408 != 0 {
		goto L20
	} else {
		goto L121
	}
L121:
	;
	F_errfinish(m, int32(_a_F_text_format_3), int32(_a_F_text_format_11), int32(_a_F_text_format_5))
	mBase = m.M
	v413 = m.ExcPending
	if v413 != 0 {
		goto L20
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
	v420 = m.ExcPending
	if v420 != 0 {
		goto L20
	} else {
		goto L124
	}
L124:
	;
	F_errmsg(m, int32(_a_F_text_format_12), int32(0))
	mBase = m.M
	v424 = m.ExcPending
	if v424 != 0 {
		goto L20
	} else {
		goto L125
	}
L125:
	;
	F_errfinish(m, int32(_a_F_text_format_3), int32(_a_F_text_format_13), int32(_a_F_text_format_5))
	mBase = m.M
	v429 = m.ExcPending
	if v429 != 0 {
		goto L20
	} else {
		goto L126
	}
L126:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L127:
	;
	F_errcode(m, int32(50856066))
	mBase = m.M
	v436 = m.ExcPending
	if v436 != 0 {
		goto L20
	} else {
		goto L128
	}
L128:
	;
	F_errmsg(m, int32(_a_F_text_format_8), int32(0))
	mBase = m.M
	v440 = m.ExcPending
	if v440 != 0 {
		goto L20
	} else {
		goto L129
	}
L129:
	;
	F_errfinish(m, int32(_a_F_text_format_3), int32(_a_F_text_format_14), int32(_a_F_text_format_5))
	mBase = m.M
	v445 = m.ExcPending
	if v445 != 0 {
		goto L20
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
	F_errcode(m, int32(50856066))
	mBase = m.M
	v452 = m.ExcPending
	if v452 != 0 {
		goto L20
	} else {
		goto L132
	}
L132:
	;
	F_errmsg(m, int32(_a_F_text_format_1), int32(0))
	mBase = m.M
	v456 = m.ExcPending
	if v456 != 0 {
		goto L20
	} else {
		goto L133
	}
L133:
	;
	F_errhint(m, int32(_a_F_text_format_2), int32(0))
	mBase = m.M
	v460 = m.ExcPending
	if v460 != 0 {
		goto L20
	} else {
		goto L134
	}
L134:
	;
	F_errfinish(m, int32(_a_F_text_format_3), int32(_a_F_text_format_15), int32(_a_F_text_format_5))
	mBase = m.M
	v465 = m.ExcPending
	if v465 != 0 {
		goto L20
	} else {
		goto L135
	}
L135:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L136:
	;
	v490 = *(*int32)(unsafe.Add(mBase, uint32(v20)+104))
	if v488 != 0 {
		goto L137
	} else {
		goto L138
	}
L137:
	;
	v492 = v490
	goto L139
L138:
	;
	v492 = int32(0)
	goto L139
L139:
	;
	v494 = *(*int32)(unsafe.Add(mBase, uint32(v20)+108))
	v498 = v214
	v499 = v494
	v500 = v492
	v502 = int32(-1)
	v504 = v475
	goto L47
L140:
	;
	if v619 == int32(0) {
		goto L5
	} else {
		goto L165
	}
L141:
	;
	v619 = int32(0)
	goto L140
L142:
	;
	v597 = v590
	v599 = v592
	goto L159
L143:
	;
	goto L150
L150:
	;
	v553 = v513 & int32(255)
	v554 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_text_format[0])))
	if base.B2i32(v553 == v554)|int32(0) == int32(0) {
		goto L151
	} else {
		goto L152
	}
L151:
	;
	v563 = v512
	v565 = v514
	goto L154
L152:
	;
	v583 = v512
	v585 = v514
	goto L153
L153:
	;
	if v585 == int32(0) {
		goto L141
	} else {
		goto L158
	}
L154:
	;
	v569 = *(*int32)(unsafe.Add(mBase, uint32(v563)))
	v570 = v569 ^ v553*int32(16843009)
	v573 = int32(-2139062144)
	if (int32(16843008)-v570|v570)&v573 != v573 {
		v590 = v563
		v592 = v565
		goto L142
	} else {
		goto L156
	}
L155:
	;
	v583 = v578
	v585 = v580
	goto L153
L156:
	;
	v577 = int32(4)
	v578 = v563 + v577
	v580 = v565 - v577
	if base.Ui32(int32(3)) < base.Ui32(v580) {
		v563 = v578
		v565 = v580
		goto L154
	} else {
		goto L157
	}
L157:
	;
	goto L155
L158:
	;
	v590 = v583
	v592 = v585
	goto L142
L159:
	;
	v602 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v597))))
	if v513&int32(255) == v602 {
		goto L161
	} else {
		goto L162
	}
L160:
	;
	goto L141
L161:
	;
	v619 = v597
	goto L140
L162:
	;
	goto L163
L163:
	;
	v604 = int32(1)
	v607 = v599 - v604
	if v607 != 0 {
		v597 = v597 + v604
		v599 = v607
		goto L159
	} else {
		goto L164
	}
L164:
	;
	goto L160
L165:
	;
	if v502 < int32(0) {
		v684 = v500
		v686 = v144
		v687 = v151
		goto L166
	} else {
		goto L167
	}
L166:
	;
	if int32(0) < v498 {
		goto L190
	} else {
		goto L191
	}
L167:
	;
	if v502 != 0 {
		goto L168
	} else {
		goto L169
	}
L168:
	;
	v624 = v502
	goto L170
L169:
	;
	v624 = v144
	goto L170
L170:
	;
	if v82 <= v624 {
		goto L4
	} else {
		goto L171
	}
L171:
	;
	if v48 == int32(0) {
		goto L173
	} else {
		goto L174
	}
L172:
	;
	if v649 == int32(0) {
		goto L3
	} else {
		goto L177
	}
L173:
	;
	v630 = v33 + v624<<(uint(int32(4))%32)
	v631 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v630)+8)))
	v632 = *(*int64)(unsafe.Add(mBase, uint32(v630)))
	v633 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v634 = F_get_fn_expr_argtype(m, v633, v624)
	mBase = m.M
	v635 = m.ExcPending
	if v635 != 0 {
		goto L20
	} else {
		goto L176
	}
L174:
	;
	goto L175
L175:
	;
	v637 = v624 - int32(1)
	v638 = *(*int32)(unsafe.Add(mBase, uint32(v20)+80))
	v640 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v637+v638))))
	v641 = *(*int32)(unsafe.Add(mBase, uint32(v20)+84))
	v645 = *(*int64)(unsafe.Add(mBase, uint32(v641+v637<<(uint(int32(3))%32))))
	v647 = v640
	v648 = v645
	v649 = v81
	goto L172
L176:
	;
	v647 = v631
	v648 = v632
	v649 = v634
	goto L172
L177:
	;
	v652 = int32(1)
	v653 = v624 + v652
	if v647&v652 != 0 {
		v684 = int32(0)
		v686 = v653
		v687 = v151
		goto L166
	} else {
		goto L178
	}
L178:
	;
	switch v649 - int32(21) {
	case 0:
		goto L180
	default:
		goto L179
	case 2:
		goto L181
	}
L179:
	;
	if v649 != v151 {
		goto L182
	} else {
		goto L183
	}
L180:
	;
	v684 = base.I32_extend16_s(base.I32_wrap_i64(v648))
	v686 = v653
	v687 = v151
	goto L166
L181:
	;
	v684 = base.I32_wrap_i64(v648)
	v686 = v653
	v687 = v151
	goto L166
L182:
	;
	F_getTypeOutputInfo(m, v649, v20+int32(108), v20+int32(104))
	mBase = m.M
	v668 = m.ExcPending
	if v668 != 0 {
		goto L20
	} else {
		goto L185
	}
L183:
	;
	v674 = v151
	goto L184
L184:
	;
	v677 = F_OutputFunctionCall(m, v20+int32(24), v648)
	mBase = m.M
	v678 = m.ExcPending
	if v678 != 0 {
		goto L20
	} else {
		goto L187
	}
L185:
	;
	v669 = *(*int32)(unsafe.Add(mBase, uint32(v20)+108))
	F_fmgr_info(m, v669, v20+int32(24))
	mBase = m.M
	v673 = m.ExcPending
	if v673 != 0 {
		goto L20
	} else {
		goto L186
	}
L186:
	;
	v674 = v649
	goto L184
L187:
	;
	v679 = F_pg_strtoint32(m, v677)
	mBase = m.M
	v680 = m.ExcPending
	if v680 != 0 {
		goto L20
	} else {
		goto L188
	}
L188:
	;
	F_pfree(m, v677)
	mBase = m.M
	v682 = m.ExcPending
	if v682 != 0 {
		goto L20
	} else {
		goto L189
	}
L189:
	;
	v684 = v679
	v686 = v653
	v687 = v674
	goto L166
L190:
	;
	v691 = v498
	goto L192
L191:
	;
	v691 = v686
	goto L192
L192:
	;
	if v82 <= v691 {
		goto L2
	} else {
		goto L193
	}
L193:
	;
	if v48 == int32(0) {
		goto L195
	} else {
		goto L196
	}
L194:
	;
	if v716 == int32(0) {
		goto L1
	} else {
		goto L199
	}
L195:
	;
	v697 = v33 + v691<<(uint(int32(4))%32)
	v698 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v697)+8)))
	v699 = *(*int64)(unsafe.Add(mBase, uint32(v697)))
	v700 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v701 = F_get_fn_expr_argtype(m, v700, v691)
	mBase = m.M
	v702 = m.ExcPending
	if v702 != 0 {
		goto L20
	} else {
		goto L198
	}
L196:
	;
	goto L197
L197:
	;
	v704 = v691 - int32(1)
	v705 = *(*int32)(unsafe.Add(mBase, uint32(v20)+80))
	v707 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v704+v705))))
	v708 = *(*int32)(unsafe.Add(mBase, uint32(v20)+84))
	v712 = *(*int64)(unsafe.Add(mBase, uint32(v708+v704<<(uint(int32(3))%32))))
	v714 = v707
	v715 = v712
	v716 = v81
	goto L194
L198:
	;
	v714 = v698
	v715 = v699
	v716 = v701
	goto L194
L199:
	;
	if v716 != v146 {
		goto L200
	} else {
		goto L201
	}
L200:
	;
	F_getTypeOutputInfo(m, v716, v20+int32(108), v20+int32(104))
	mBase = m.M
	v725 = m.ExcPending
	if v725 != 0 {
		goto L20
	} else {
		goto L203
	}
L201:
	;
	v731 = v146
	goto L202
L202:
	;
	v732 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v499))))
	v734 = v732 - int32(73)
	switch v734 {
	case 0, 3:
		goto L206
	case 1, 2:
		goto L205
	default:
		goto L207
	}
L203:
	;
	v726 = *(*int32)(unsafe.Add(mBase, uint32(v20)+108))
	F_fmgr_info(m, v726, v20+int32(52))
	mBase = m.M
	v730 = m.ExcPending
	if v730 != 0 {
		goto L20
	} else {
		goto L204
	}
L204:
	;
	v731 = v716
	goto L202
L205:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v800 = m.ExcPending
	if v800 != 0 {
		goto L20
	} else {
		goto L235
	}
L206:
	;
	v737 = int32(1)
	v738 = v691 + v737
	if v714&v737 != 0 {
		goto L209
	} else {
		goto L210
	}
L207:
	;
	if v732 != int32(115) {
		goto L205
	} else {
		goto L208
	}
L208:
	;
	goto L206
L209:
	;
	switch v734 {
	case 0:
		goto L212
	case 1, 2:
		v826 = v499
		v830 = v738
		v832 = v731
		v837 = v687
		goto L46
	case 3:
		goto L213
	default:
		goto L214
	}
L210:
	;
	goto L211
L211:
	;
	v771 = F_OutputFunctionCall(m, v20+int32(52), v715)
	mBase = m.M
	v772 = m.ExcPending
	if v772 != 0 {
		goto L20
	} else {
		goto L222
	}
L212:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v756 = m.ExcPending
	if v756 != 0 {
		goto L20
	} else {
		goto L218
	}
L213:
	;
	F_text_format_append_string(m, v20+int32(88), int32(_a_F_text_format_16), v504, v684)
	mBase = m.M
	v752 = m.ExcPending
	if v752 != 0 {
		goto L20
	} else {
		goto L217
	}
L214:
	;
	if v732 != int32(115) {
		v826 = v499
		v830 = v738
		v832 = v731
		v837 = v687
		goto L46
	} else {
		goto L215
	}
L215:
	;
	F_text_format_append_string(m, v20+int32(88), int32(_a_F_text_format_17), v504, v684)
	mBase = m.M
	v747 = m.ExcPending
	if v747 != 0 {
		goto L20
	} else {
		goto L216
	}
L216:
	;
	v826 = v499
	v830 = v738
	v832 = v731
	v837 = v687
	goto L46
L217:
	;
	v826 = v499
	v830 = v738
	v832 = v731
	v837 = v687
	goto L46
L218:
	;
	F_errcode(m, int32(67108994))
	mBase = m.M
	v759 = m.ExcPending
	if v759 != 0 {
		goto L20
	} else {
		goto L219
	}
L219:
	;
	F_errmsg(m, int32(_a_F_text_format_18), int32(0))
	mBase = m.M
	v763 = m.ExcPending
	if v763 != 0 {
		goto L20
	} else {
		goto L220
	}
L220:
	;
	F_errfinish(m, int32(_a_F_text_format_3), int32(_a_F_text_format_19), int32(_a_F_text_format_20))
	mBase = m.M
	v768 = m.ExcPending
	if v768 != 0 {
		goto L20
	} else {
		goto L221
	}
L221:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L222:
	;
	switch v734 {
	case 0:
		goto L225
	default:
		goto L223
	case 3:
		goto L224
	}
L223:
	;
	F_text_format_append_string(m, v20+int32(88), v771, v504, v684)
	mBase = m.M
	v794 = m.ExcPending
	if v794 != 0 {
		goto L20
	} else {
		goto L233
	}
L224:
	;
	v783 = F_quote_literal_cstr(m, v771)
	mBase = m.M
	v784 = m.ExcPending
	if v784 != 0 {
		goto L20
	} else {
		goto L229
	}
L225:
	;
	v775 = F_quote_identifier(m, v771)
	mBase = m.M
	v776 = m.ExcPending
	if v776 != 0 {
		goto L20
	} else {
		goto L226
	}
L226:
	;
	F_text_format_append_string(m, v20+int32(88), v775, v504, v684)
	mBase = m.M
	v778 = m.ExcPending
	if v778 != 0 {
		goto L20
	} else {
		goto L227
	}
L227:
	;
	F_pfree(m, v771)
	mBase = m.M
	v780 = m.ExcPending
	if v780 != 0 {
		goto L20
	} else {
		goto L228
	}
L228:
	;
	v826 = v499
	v830 = v738
	v832 = v731
	v837 = v687
	goto L46
L229:
	;
	F_text_format_append_string(m, v20+int32(88), v783, v504, v684)
	mBase = m.M
	v786 = m.ExcPending
	if v786 != 0 {
		goto L20
	} else {
		goto L230
	}
L230:
	;
	F_pfree(m, v783)
	mBase = m.M
	v788 = m.ExcPending
	if v788 != 0 {
		goto L20
	} else {
		goto L231
	}
L231:
	;
	F_pfree(m, v771)
	mBase = m.M
	v790 = m.ExcPending
	if v790 != 0 {
		goto L20
	} else {
		goto L232
	}
L232:
	;
	v826 = v499
	v830 = v738
	v832 = v731
	v837 = v687
	goto L46
L233:
	;
	F_pfree(m, v771)
	mBase = m.M
	v796 = m.ExcPending
	if v796 != 0 {
		goto L20
	} else {
		goto L234
	}
L234:
	;
	v826 = v499
	v830 = v738
	v832 = v731
	v837 = v687
	goto L46
L235:
	;
	F_errcode(m, int32(50856066))
	mBase = m.M
	v803 = m.ExcPending
	if v803 != 0 {
		goto L20
	} else {
		goto L236
	}
L236:
	;
	v804 = F_pg_mblen_range(m, v499, v134)
	mBase = m.M
	v805 = m.ExcPending
	if v805 != 0 {
		goto L20
	} else {
		goto L237
	}
L237:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v20)+20)) = v499
	*(*int32)(unsafe.Add(mBase, uint32(v20)+16)) = v804
	F_errmsg(m, int32(_a_F_text_format_21), v20+int32(16))
	mBase = m.M
	v812 = m.ExcPending
	if v812 != 0 {
		goto L20
	} else {
		goto L238
	}
L238:
	;
	F_errhint(m, int32(_a_F_text_format_2), int32(0))
	mBase = m.M
	v816 = m.ExcPending
	if v816 != 0 {
		goto L20
	} else {
		goto L239
	}
L239:
	;
	F_errfinish(m, int32(_a_F_text_format_3), int32(_a_F_text_format_22), int32(_a_F_text_format_7))
	mBase = m.M
	v821 = m.ExcPending
	if v821 != 0 {
		goto L20
	} else {
		goto L240
	}
L240:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L241:
	;
	goto L45
L242:
	;
	F_pfree(m, v859)
	mBase = m.M
	v861 = m.ExcPending
	if v861 != 0 {
		goto L20
	} else {
		goto L245
	}
L243:
	;
	goto L244
L244:
	;
	v862 = *(*int32)(unsafe.Add(mBase, uint32(v20)+80))
	if v862 != 0 {
		goto L246
	} else {
		goto L247
	}
L245:
	;
	goto L244
L246:
	;
	F_pfree(m, v862)
	mBase = m.M
	v864 = m.ExcPending
	if v864 != 0 {
		goto L20
	} else {
		goto L249
	}
L247:
	;
	goto L248
L248:
	;
	v865 = *(*int32)(unsafe.Add(mBase, uint32(v20)+88))
	v866 = *(*int32)(unsafe.Add(mBase, uint32(v20)+92))
	v868 = v866 + int32(4)
	v869 = F_palloc(m, v868)
	mBase = m.M
	v870 = m.ExcPending
	if v870 != 0 {
		goto L20
	} else {
		goto L250
	}
L249:
	;
	goto L248
L250:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v869))) = v868 << (uint(int32(2)) % 32)
	if v866 != 0 {
		goto L251
	} else {
		goto L252
	}
L251:
	;
	base.MemoryCopy(m, v869+int32(4), v865, v866)
	goto L253
L252:
	;
	goto L253
L253:
	;
	v877 = *(*int32)(unsafe.Add(mBase, uint32(v20)+88))
	F_pfree(m, v877)
	mBase = m.M
	v879 = m.ExcPending
	if v879 != 0 {
		goto L20
	} else {
		goto L254
	}
L254:
	;
	v898 = base.I64_extend_i32_u(v869)
	goto L6
L255:
	;
	F_errcode(m, int32(50856066))
	mBase = m.M
	v909 = m.ExcPending
	if v909 != 0 {
		goto L20
	} else {
		goto L256
	}
L256:
	;
	v910 = F_pg_mblen_range(m, v499, v134)
	mBase = m.M
	v911 = m.ExcPending
	if v911 != 0 {
		goto L20
	} else {
		goto L257
	}
L257:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v20)+4)) = v499
	*(*int32)(unsafe.Add(mBase, uint32(v20))) = v910
	F_errmsg(m, int32(_a_F_text_format_21), v20)
	mBase = m.M
	v916 = m.ExcPending
	if v916 != 0 {
		goto L20
	} else {
		goto L258
	}
L258:
	;
	F_errhint(m, int32(_a_F_text_format_2), int32(0))
	mBase = m.M
	v920 = m.ExcPending
	if v920 != 0 {
		goto L20
	} else {
		goto L259
	}
L259:
	;
	F_errfinish(m, int32(_a_F_text_format_3), int32(_a_F_text_format_23), int32(_a_F_text_format_7))
	mBase = m.M
	v925 = m.ExcPending
	if v925 != 0 {
		goto L20
	} else {
		goto L260
	}
L260:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L261:
	;
	F_errcode(m, int32(50856066))
	mBase = m.M
	v932 = m.ExcPending
	if v932 != 0 {
		goto L20
	} else {
		goto L262
	}
L262:
	;
	F_errmsg(m, int32(_a_F_text_format_24), int32(0))
	mBase = m.M
	v936 = m.ExcPending
	if v936 != 0 {
		goto L20
	} else {
		goto L263
	}
L263:
	;
	F_errfinish(m, int32(_a_F_text_format_3), int32(_a_F_text_format_25), int32(_a_F_text_format_7))
	mBase = m.M
	v941 = m.ExcPending
	if v941 != 0 {
		goto L20
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
	F_errmsg_internal(m, int32(_a_F_text_format_26), int32(0))
	mBase = m.M
	v949 = m.ExcPending
	if v949 != 0 {
		goto L20
	} else {
		goto L266
	}
L266:
	;
	F_errfinish(m, int32(_a_F_text_format_3), int32(_a_F_text_format_27), int32(_a_F_text_format_7))
	mBase = m.M
	v954 = m.ExcPending
	if v954 != 0 {
		goto L20
	} else {
		goto L267
	}
L267:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L268:
	;
	F_errcode(m, int32(50856066))
	mBase = m.M
	v961 = m.ExcPending
	if v961 != 0 {
		goto L20
	} else {
		goto L269
	}
L269:
	;
	F_errmsg(m, int32(_a_F_text_format_24), int32(0))
	mBase = m.M
	v965 = m.ExcPending
	if v965 != 0 {
		goto L20
	} else {
		goto L270
	}
L270:
	;
	F_errfinish(m, int32(_a_F_text_format_3), int32(_a_F_text_format_28), int32(_a_F_text_format_7))
	mBase = m.M
	v970 = m.ExcPending
	if v970 != 0 {
		goto L20
	} else {
		goto L271
	}
L271:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L272:
	;
	F_errmsg_internal(m, int32(_a_F_text_format_26), int32(0))
	mBase = m.M
	v978 = m.ExcPending
	if v978 != 0 {
		goto L20
	} else {
		goto L273
	}
L273:
	;
	F_errfinish(m, int32(_a_F_text_format_3), int32(_a_F_text_format_29), int32(_a_F_text_format_7))
	mBase = m.M
	v983 = m.ExcPending
	if v983 != 0 {
		goto L20
	} else {
		goto L274
	}
L274:
	;
	base.Wasm_trap_unreachable()
	for {
	}
}
func F_text_larger(m *base.Module, l0 int32) int64 {
	mBase := m.M
	_ = mBase
	var v7 int32
	_ = v7
	var v8 int32
	_ = v8
	var v11 int32
	_ = v11
	var v12 int32
	_ = v12
	var v13 int32
	_ = v13
	var v14 int32
	_ = v14
	var v15 int32
	_ = v15
	var v16 int32
	_ = v16
	var v22 int32
	_ = v22
	var v25 int32
	_ = v25
	var v32 int32
	_ = v32
	var v33 int32
	_ = v33
	var v39 int32
	_ = v39
	var v45 int32
	_ = v45
	var v46 int32
	_ = v46
	var v47 int32
	_ = v47
	var v51 int32
	_ = v51
	var v53 int32
	_ = v53
	var v57 int32
	_ = v57
	var v64 int32
	_ = v64
	var v67 int32
	_ = v67
	var v74 int32
	_ = v74
	var v75 int32
	_ = v75
	var v81 int32
	_ = v81
	var v87 int32
	_ = v87
	var v88 int32
	_ = v88
	var v89 int32
	_ = v89
	var v92 int32
	_ = v92
	v7 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
	v8 = F_pg_detoast_datum_packed(m, v7)
	mBase = m.M
	v11 = m.ExcPending
	if v11 != 0 {
		return int64(0)
	} else {
		v12 = *(*int32)(unsafe.Add(mBase, uint32(l0)+40))
		v13 = F_pg_detoast_datum_packed(m, v12)
		mBase = m.M
		v14 = m.ExcPending
		if v14 != 0 {
			return int64(0)
		} else {
			v15 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v13))))
			v16 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v8))))
			if v16 == int32(1) {
				v22 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v8)+1)))
				if v22 == int32(18) {
					v25 = int32(16)
				} else {
					v25 = int32(0)
				}
				if base.Ui32((v22-int32(1))&int32(255)) < base.Ui32(int32(3)) {
					v32 = int32(4)
				} else {
					v32 = v25
				}
				v45 = v32
			} else {
				v33 = int32(1)
				if v16&v33 != 0 {
					v45 = int32(base.Ui32(v16)>>(uint(v33)%32)) - v33
				} else {
					v39 = *(*int32)(unsafe.Add(mBase, uint32(v8)))
					v45 = int32(base.Ui32(v39)>>(uint(int32(2))%32)) - int32(4)
				}
			}
			v46 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
			v47 = int32(1)
			if v16&v47 != 0 {
				v51 = v47
			} else {
				v51 = int32(4)
			}
			v53 = int32(1)
			if v15&v53 != 0 {
				v57 = v53
			} else {
				v57 = int32(4)
			}
			if v15 == int32(1) {
				v64 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v13)+1)))
				if v64 == int32(18) {
					v67 = int32(16)
				} else {
					v67 = int32(0)
				}
				if base.Ui32((v64-int32(1))&int32(255)) < base.Ui32(int32(3)) {
					v74 = int32(4)
				} else {
					v74 = v67
				}
				v87 = v74
			} else {
				v75 = int32(1)
				if v15&v75 != 0 {
					v87 = int32(base.Ui32(v15)>>(uint(v75)%32)) - v75
				} else {
					v81 = *(*int32)(unsafe.Add(mBase, uint32(v13)))
					v87 = int32(base.Ui32(v81)>>(uint(int32(2))%32)) - int32(4)
				}
			}
			v88 = F_varstr_cmp(m, v8+v51, v45, v13+v57, v87, v46)
			mBase = m.M
			v89 = m.ExcPending
			if v89 != 0 {
				return int64(0)
			} else {
				if int32(0) < v88 {
					v92 = v8
				} else {
					v92 = v13
				}
				return base.I64_extend_i32_u(v92)
			}
		}
	}
}
func F_text_le(m *base.Module, l0 int32) int64 {
	mBase := m.M
	_ = mBase
	var v8 int32
	_ = v8
	var v9 int32
	_ = v9
	var v12 int32
	_ = v12
	var v13 int32
	_ = v13
	var v14 int32
	_ = v14
	var v15 int32
	_ = v15
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
	var v47 int32
	_ = v47
	var v48 int32
	_ = v48
	var v52 int32
	_ = v52
	var v54 int32
	_ = v54
	var v58 int32
	_ = v58
	var v65 int32
	_ = v65
	var v68 int32
	_ = v68
	var v75 int32
	_ = v75
	var v76 int32
	_ = v76
	var v82 int32
	_ = v82
	var v88 int32
	_ = v88
	var v89 int32
	_ = v89
	var v90 int32
	_ = v90
	var v91 int32
	_ = v91
	var v94 int32
	_ = v94
	var v95 int32
	_ = v95
	var v98 int32
	_ = v98
	v8 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
	v9 = F_pg_detoast_datum_packed(m, v8)
	mBase = m.M
	v12 = m.ExcPending
	if v12 != 0 {
		return int64(0)
	} else {
		v13 = *(*int32)(unsafe.Add(mBase, uint32(l0)+40))
		v14 = F_pg_detoast_datum_packed(m, v13)
		mBase = m.M
		v15 = m.ExcPending
		if v15 != 0 {
			return int64(0)
		} else {
			v16 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v14))))
			v17 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v9))))
			if v17 == int32(1) {
				v23 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v9)+1)))
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
					v40 = *(*int32)(unsafe.Add(mBase, uint32(v9)))
					v46 = int32(base.Ui32(v40)>>(uint(int32(2))%32)) - int32(4)
				}
			}
			v47 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
			v48 = int32(1)
			if v17&v48 != 0 {
				v52 = v48
			} else {
				v52 = int32(4)
			}
			v54 = int32(1)
			if v16&v54 != 0 {
				v58 = v54
			} else {
				v58 = int32(4)
			}
			if v16 == int32(1) {
				v65 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v14)+1)))
				if v65 == int32(18) {
					v68 = int32(16)
				} else {
					v68 = int32(0)
				}
				if base.Ui32((v65-int32(1))&int32(255)) < base.Ui32(int32(3)) {
					v75 = int32(4)
				} else {
					v75 = v68
				}
				v88 = v75
			} else {
				v76 = int32(1)
				if v16&v76 != 0 {
					v88 = int32(base.Ui32(v16)>>(uint(v76)%32)) - v76
				} else {
					v82 = *(*int32)(unsafe.Add(mBase, uint32(v14)))
					v88 = int32(base.Ui32(v82)>>(uint(int32(2))%32)) - int32(4)
				}
			}
			v89 = F_varstr_cmp(m, v9+v52, v46, v14+v58, v88, v47)
			mBase = m.M
			v90 = m.ExcPending
			if v90 != 0 {
				return int64(0)
			} else {
				v91 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
				if v91 != v9 {
					F_pfree(m, v9)
					mBase = m.M
					v94 = m.ExcPending
					if v94 != 0 {
						return int64(0)
					} else {
						v95 = *(*int32)(unsafe.Add(mBase, uint32(l0)+40))
						if v95 != v14 {
							F_pfree(m, v14)
							mBase = m.M
							v98 = m.ExcPending
							if v98 != 0 {
								return int64(0)
							} else {
								return base.I64_extend_i32_u(base.B2i32(v89 <= int32(0)))
							}
						} else {
							return base.I64_extend_i32_u(base.B2i32(v89 <= int32(0)))
						}
					}
				} else {
					v95 = *(*int32)(unsafe.Add(mBase, uint32(l0)+40))
					if v95 != v14 {
						F_pfree(m, v14)
						mBase = m.M
						v98 = m.ExcPending
						if v98 != 0 {
							return int64(0)
						} else {
							return base.I64_extend_i32_u(base.B2i32(v89 <= int32(0)))
						}
					} else {
						return base.I64_extend_i32_u(base.B2i32(v89 <= int32(0)))
					}
				}
			}
		}
	}
}
func F_text_left(m *base.Module, l0 int32) int64 {
	mBase := m.M
	_ = mBase
	var v7 int64
	_ = v7
	var v8 int32
	_ = v8
	var v12 int32
	_ = v12
	var v15 int32
	_ = v15
	var v16 int32
	_ = v16
	var v18 int32
	_ = v18
	var v20 int32
	_ = v20
	var v21 int32
	_ = v21
	var v22 int32
	_ = v22
	var v28 int32
	_ = v28
	var v31 int32
	_ = v31
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
	var v53 int32
	_ = v53
	var v54 int32
	_ = v54
	var v56 int32
	_ = v56
	var v57 int32
	_ = v57
	var v58 int32
	_ = v58
	var v71 int32
	_ = v71
	var v72 int32
	_ = v72
	var v73 int32
	_ = v73
	v7 = *(*int64)(unsafe.Add(mBase, uint32(l0)+24))
	v8 = *(*int32)(unsafe.Add(mBase, uint32(l0)+40))
	if v8 < int32(0) {
		v12 = F_pg_detoast_datum_packed(m, base.I32_wrap_i64(v7))
		mBase = m.M
		v15 = m.ExcPending
		if v15 != 0 {
			return int64(0)
		} else {
			v16 = int32(1)
			v18 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v12))))
			v20 = v18 & v16
			if v20 != 0 {
				v21 = v16
			} else {
				v21 = int32(4)
			}
			v22 = v12 + v21
			if v18 == int32(1) {
				v28 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v12)+1)))
				if v28 == int32(18) {
					v31 = int32(16)
				} else {
					v31 = int32(0)
				}
				if base.Ui32((v28-int32(1))&int32(255)) < base.Ui32(int32(3)) {
					v38 = int32(4)
				} else {
					v38 = v31
				}
				v49 = v38
			} else {
				v39 = int32(1)
				if v20 != 0 {
					v49 = int32(base.Ui32(v18)>>(uint(v39)%32)) - v39
				} else {
					v43 = *(*int32)(unsafe.Add(mBase, uint32(v12)))
					v49 = int32(base.Ui32(v43)>>(uint(int32(2))%32)) - int32(4)
				}
			}
			v50 = F_pg_mbstrlen_with_len(m, v22, v49)
			mBase = m.M
			v51 = m.ExcPending
			if v51 != 0 {
				return int64(0)
			} else {
				v53 = F_pg_mbcharcliplen(m, v22, v49, v50+v8)
				mBase = m.M
				v54 = m.ExcPending
				if v54 != 0 {
					return int64(0)
				} else {
					v56 = v53 + int32(4)
					v57 = F_palloc(m, v56)
					mBase = m.M
					v58 = m.ExcPending
					if v58 != 0 {
						return int64(0)
					} else {
						*(*int32)(unsafe.Add(mBase, uint32(v57))) = v56 << (uint(int32(2)) % 32)
						if v53 == int32(0) {
							v73 = v57
							return base.I64_extend_i32_u(v73)
						} else {
							base.MemoryCopy(m, v57+int32(4), v22, v53)
							return base.I64_extend_i32_u(v57)
						}
					}
				}
			}
		}
	} else {
		v71 = F_text_substring(m, v7, int32(1), v8, int32(0))
		mBase = m.M
		v72 = m.ExcPending
		if v72 != 0 {
			return int64(0)
		} else {
			v73 = v71
			return base.I64_extend_i32_u(v73)
		}
	}
}
func F_text_name(m *base.Module, l0 int32) int64 {
	mBase := m.M
	_ = mBase
	var v4 int32
	_ = v4
	var v5 int32
	_ = v5
	var v8 int32
	_ = v8
	var v9 int32
	_ = v9
	var v15 int32
	_ = v15
	var v18 int32
	_ = v18
	var v25 int32
	_ = v25
	var v28 int32
	_ = v28
	var v32 int32
	_ = v32
	var v37 int32
	_ = v37
	var v40 int32
	_ = v40
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
	var v52 int32
	_ = v52
	var v53 int32
	_ = v53
	var v55 int32
	_ = v55
	var v58 int32
	_ = v58
	v4 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
	v5 = F_pg_detoast_datum_packed(m, v4)
	mBase = m.M
	v8 = m.ExcPending
	if v8 != 0 {
		return int64(0)
	} else {
		v9 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v5))))
		if v9 == int32(1) {
			v15 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v5)+1)))
			if v15 == int32(18) {
				v18 = int32(16)
			} else {
				v18 = int32(0)
			}
			if base.Ui32((v15-int32(1))&int32(255)) < base.Ui32(int32(3)) {
				v25 = int32(4)
			} else {
				v25 = v18
			}
			v49 = v25
			v51 = F_palloc0(m, int32(64))
			mBase = m.M
			v52 = m.ExcPending
			if v52 != 0 {
				return int64(0)
			} else {
				if v49 != 0 {
					v53 = int32(1)
					v55 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v5))))
					if v55&v53 != 0 {
						v58 = v53
					} else {
						v58 = int32(4)
					}
					base.MemoryCopy(m, v51, v5+v58, v49)
				} else {
				}
				return base.I64_extend_i32_u(v51)
			}
		} else {
			if v9&int32(1) != 0 {
				v28 = int32(1)
				v37 = int32(base.Ui32(v9)>>(uint(v28)%32)) - v28
			} else {
				v32 = *(*int32)(unsafe.Add(mBase, uint32(v5)))
				v37 = int32(base.Ui32(v32)>>(uint(int32(2))%32)) - int32(4)
			}
			if v37 < int32(64) {
				v49 = v37
				v51 = F_palloc0(m, int32(64))
				mBase = m.M
				v52 = m.ExcPending
				if v52 != 0 {
					return int64(0)
				} else {
					if v49 != 0 {
						v53 = int32(1)
						v55 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v5))))
						if v55&v53 != 0 {
							v58 = v53
						} else {
							v58 = int32(4)
						}
						base.MemoryCopy(m, v51, v5+v58, v49)
					} else {
					}
					return base.I64_extend_i32_u(v51)
				}
			} else {
				v40 = int32(1)
				if v9&v40 != 0 {
					v44 = v40
				} else {
					v44 = int32(4)
				}
				v47 = F_pg_mbcliplen(m, v5+v44, v37, int32(63))
				mBase = m.M
				v48 = m.ExcPending
				if v48 != 0 {
					return int64(0)
				} else {
					v49 = v47
					v51 = F_palloc0(m, int32(64))
					mBase = m.M
					v52 = m.ExcPending
					if v52 != 0 {
						return int64(0)
					} else {
						if v49 != 0 {
							v53 = int32(1)
							v55 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v5))))
							if v55&v53 != 0 {
								v58 = v53
							} else {
								v58 = int32(4)
							}
							base.MemoryCopy(m, v51, v5+v58, v49)
						} else {
						}
						return base.I64_extend_i32_u(v51)
					}
				}
			}
		}
	}
}
func F_text_substr_no_len(m *base.Module, l0 int32) int64 {
	mBase := m.M
	_ = mBase
	var v2 int64
	_ = v2
	var v3 int32
	_ = v3
	var v6 int32
	_ = v6
	var v9 int32
	_ = v9
	v2 = *(*int64)(unsafe.Add(mBase, uint32(l0)+24))
	v3 = *(*int32)(unsafe.Add(mBase, uint32(l0)+40))
	v6 = F_text_substring(m, v2, v3, int32(-1), int32(1))
	mBase = m.M
	v9 = m.ExcPending
	if v9 != 0 {
		return int64(0)
	} else {
		return base.I64_extend_i32_u(v6)
	}
}
