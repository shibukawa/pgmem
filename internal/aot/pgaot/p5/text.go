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
	var v9 int32
	_ = v9
	v4 = F_cstring_to_text(m, l2)
	v5 = m.ExcPending
	if v5 != 0 {
		return
	} else {
		F_assign_simple_var(m, l0, l1, v4, int32(0), int32(1))
		v9 = m.ExcPending
		if v9 != 0 {
			return
		} else {
			return
		}
	}
}
func F_text_concat_ws(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v6 int32
	_ = v6
	var v9 int32
	_ = v9
	var v10 int32
	_ = v10
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
	var v48 int32
	_ = v48
	var v49 int32
	_ = v49
	var v50 int32
	_ = v50
	var v52 int32
	_ = v52
	var v55 int32
	_ = v55
	var v59 int32
	_ = v59
	var v63 int32
	_ = v63
	var v65 int32
	_ = v65
	var v66 int32
	_ = v66
	var v71 int32
	_ = v71
	var v74 int32
	_ = v74
	v6 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+24)))
	if v6 == int32(1) {
		v71 = int32(1)
		*(*uint8)(unsafe.Add(mBase, uint32(l0)+16)) = uint8(v71)
		v74 = int32(0)
		return v74
	} else {
		v9 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
		v10 = F_pg_detoast_datum_packed(m, v9)
		mBase = m.M
		v13 = m.ExcPending
		if v13 != 0 {
			return int32(0)
		} else {
			v14 = F_pg_detoast_datum_packed(m, v10)
			mBase = m.M
			v15 = m.ExcPending
			if v15 != 0 {
				return int32(0)
			} else {
				v16 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v14))))
				if v16 == int32(1) {
					v22 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v14)+1)))
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
						v39 = *(*int32)(unsafe.Add(mBase, uint32(v14)))
						v45 = int32(base.Ui32(v39)>>(uint(int32(2))%32)) - int32(4)
					}
				}
				v48 = F_palloc(m, v45+int32(1))
				mBase = m.M
				v49 = m.ExcPending
				if v49 != 0 {
					return int32(0)
				} else {
					if v45 != 0 {
						v50 = int32(1)
						v52 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v14))))
						if v52&v50 != 0 {
							v55 = v50
						} else {
							v55 = int32(4)
						}
						base.MemoryCopy(m, v48, v14+v55, v45)
					} else {
					}
					v59 = int32(0)
					*(*uint8)(unsafe.Add(mBase, uint32(v45+v48))) = uint8(v59)
					if v14 != v10 {
						F_pfree(m, v14)
						mBase = m.M
						v63 = m.ExcPending
						if v63 != 0 {
							return int32(0)
						} else {
							v65 = F_concat_internal(m, v48, int32(1), l0)
							mBase = m.M
							v66 = m.ExcPending
							if v66 != 0 {
								return int32(0)
							} else {
								if v65 != 0 {
									v74 = v65
								} else {
									v71 = int32(1)
									*(*uint8)(unsafe.Add(mBase, uint32(l0)+16)) = uint8(v71)
									v74 = int32(0)
								}
								return v74
							}
						}
					} else {
						v65 = F_concat_internal(m, v48, int32(1), l0)
						mBase = m.M
						v66 = m.ExcPending
						if v66 != 0 {
							return int32(0)
						} else {
							if v65 != 0 {
								v74 = v65
							} else {
								v71 = int32(1)
								*(*uint8)(unsafe.Add(mBase, uint32(l0)+16)) = uint8(v71)
								v74 = int32(0)
							}
							return v74
						}
					}
				}
			}
		}
	}
}
func F_text_format(m *base.Module, l0 int32) int32 {
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
	var v32 int32
	_ = v32
	var v33 int32
	_ = v33
	var v34 int32
	_ = v34
	var v37 int32
	_ = v37
	var v40 int32
	_ = v40
	var v43 int32
	_ = v43
	var v45 int32
	_ = v45
	var v47 int32
	_ = v47
	var v49 int32
	_ = v49
	var v50 int32
	_ = v50
	var v51 int32
	_ = v51
	var v54 int32
	_ = v54
	var v55 int32
	_ = v55
	var v63 int32
	_ = v63
	var v64 int32
	_ = v64
	var v65 int32
	_ = v65
	var v66 int32
	_ = v66
	var v74 int32
	_ = v74
	var v75 int32
	_ = v75
	var v78 int32
	_ = v78
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
	var v85 int32
	_ = v85
	var v86 int32
	_ = v86
	var v89 int32
	_ = v89
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
	var v139 int32
	_ = v139
	var v144 int32
	_ = v144
	var v151 int32
	_ = v151
	var v152 int32
	_ = v152
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
	var v226 int32
	_ = v226
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
	var v289 int32
	_ = v289
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
	var v469 int32
	_ = v469
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
	var v501 int32
	_ = v501
	var v504 int32
	_ = v504
	var v513 int32
	_ = v513
	var v528 int32
	_ = v528
	var v530 int32
	_ = v530
	var v533 int32
	_ = v533
	var v535 int32
	_ = v535
	var v536 int32
	_ = v536
	var v537 int32
	_ = v537
	var v540 int32
	_ = v540
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
	var v632 int32
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
	var v645 int32
	_ = v645
	var v647 int32
	_ = v647
	var v648 int32
	_ = v648
	var v649 int32
	_ = v649
	var v652 int32
	_ = v652
	var v653 int32
	_ = v653
	var v666 int32
	_ = v666
	var v667 int32
	_ = v667
	var v671 int32
	_ = v671
	var v672 int32
	_ = v672
	var v675 int32
	_ = v675
	var v676 int32
	_ = v676
	var v677 int32
	_ = v677
	var v678 int32
	_ = v678
	var v680 int32
	_ = v680
	var v683 int32
	_ = v683
	var v684 int32
	_ = v684
	var v686 int32
	_ = v686
	var v689 int32
	_ = v689
	var v695 int32
	_ = v695
	var v696 int32
	_ = v696
	var v697 int32
	_ = v697
	var v698 int32
	_ = v698
	var v699 int32
	_ = v699
	var v700 int32
	_ = v700
	var v702 int32
	_ = v702
	var v703 int32
	_ = v703
	var v705 int32
	_ = v705
	var v706 int32
	_ = v706
	var v710 int32
	_ = v710
	var v711 int32
	_ = v711
	var v712 int32
	_ = v712
	var v713 int32
	_ = v713
	var v722 int32
	_ = v722
	var v723 int32
	_ = v723
	var v727 int32
	_ = v727
	var v728 int32
	_ = v728
	var v729 int32
	_ = v729
	var v731 int32
	_ = v731
	var v734 int32
	_ = v734
	var v735 int32
	_ = v735
	var v744 int32
	_ = v744
	var v749 int32
	_ = v749
	var v753 int32
	_ = v753
	var v756 int32
	_ = v756
	var v760 int32
	_ = v760
	var v765 int32
	_ = v765
	var v768 int32
	_ = v768
	var v769 int32
	_ = v769
	var v772 int32
	_ = v772
	var v773 int32
	_ = v773
	var v775 int32
	_ = v775
	var v777 int32
	_ = v777
	var v780 int32
	_ = v780
	var v781 int32
	_ = v781
	var v783 int32
	_ = v783
	var v785 int32
	_ = v785
	var v787 int32
	_ = v787
	var v791 int32
	_ = v791
	var v793 int32
	_ = v793
	var v797 int32
	_ = v797
	var v800 int32
	_ = v800
	var v801 int32
	_ = v801
	var v802 int32
	_ = v802
	var v809 int32
	_ = v809
	var v813 int32
	_ = v813
	var v818 int32
	_ = v818
	var v822 int32
	_ = v822
	var v827 int32
	_ = v827
	var v834 int32
	_ = v834
	var v835 int32
	_ = v835
	var v837 int32
	_ = v837
	var v856 int32
	_ = v856
	var v858 int32
	_ = v858
	var v859 int32
	_ = v859
	var v861 int32
	_ = v861
	var v862 int32
	_ = v862
	var v863 int32
	_ = v863
	var v865 int32
	_ = v865
	var v866 int32
	_ = v866
	var v867 int32
	_ = v867
	var v874 int32
	_ = v874
	var v876 int32
	_ = v876
	var v880 int32
	_ = v880
	var v901 int32
	_ = v901
	var v904 int32
	_ = v904
	var v905 int32
	_ = v905
	var v906 int32
	_ = v906
	var v911 int32
	_ = v911
	var v915 int32
	_ = v915
	var v920 int32
	_ = v920
	var v924 int32
	_ = v924
	var v927 int32
	_ = v927
	var v931 int32
	_ = v931
	var v936 int32
	_ = v936
	var v940 int32
	_ = v940
	var v944 int32
	_ = v944
	var v949 int32
	_ = v949
	var v953 int32
	_ = v953
	var v956 int32
	_ = v956
	var v960 int32
	_ = v960
	var v965 int32
	_ = v965
	var v969 int32
	_ = v969
	var v973 int32
	_ = v973
	var v978 int32
	_ = v978
	v2 = int32(0)
	v18 = m.G0
	v20 = v18 - int32(112)
	m.G0 = v20
	*(*int32)(unsafe.Add(mBase, uint32(v20)+84)) = v2
	*(*int32)(unsafe.Add(mBase, uint32(v20)+80)) = v2
	v26 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+24)))
	if v26 == int32(1) {
		goto L7
	} else {
		goto L8
	}
L1:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v969 = m.ExcPending
	if v969 != 0 {
		goto L20
	} else {
		goto L273
	}
L2:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v953 = m.ExcPending
	if v953 != 0 {
		goto L20
	} else {
		goto L269
	}
L3:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v940 = m.ExcPending
	if v940 != 0 {
		goto L20
	} else {
		goto L266
	}
L4:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v924 = m.ExcPending
	if v924 != 0 {
		goto L20
	} else {
		goto L262
	}
L5:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v901 = m.ExcPending
	if v901 != 0 {
		goto L20
	} else {
		goto L256
	}
L6:
	;
	m.G0 = v20 + int32(112)
	return v880
L7:
	;
	v29 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(l0)+16)) = uint8(v29)
	v880 = v2
	goto L6
L8:
	;
	goto L9
L9:
	;
	v32 = l0 + int32(20)
	v33 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v34 = int32(0)
	if v33 == v34 {
		v45 = v34
		goto L12
	} else {
		goto L13
	}
L10:
	;
	v82 = *(*int32)(unsafe.Add(mBase, uint32(v32)))
	v83 = F_pg_detoast_datum_packed(m, v82)
	mBase = m.M
	v84 = m.ExcPending
	if v84 != 0 {
		goto L20
	} else {
		goto L24
	}
L11:
	;
	if v47 != 0 {
		goto L16
	} else {
		goto L17
	}
L12:
	;
	v47 = v45 & int32(1)
	goto L11
L13:
	;
	v37 = *(*int32)(unsafe.Add(mBase, uint32(v33)+24))
	if v37 == int32(0) {
		v45 = v34
		goto L12
	} else {
		goto L14
	}
L14:
	;
	v40 = *(*int32)(unsafe.Add(mBase, uint32(v37)))
	if v40 != int32(15) {
		v45 = v34
		goto L12
	} else {
		goto L15
	}
L15:
	;
	v43 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v37)+13)))
	v45 = v43
	goto L12
L16:
	;
	v49 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+32)))
	if v49 != 0 {
		v80 = v2
		v81 = int32(1)
		goto L10
	} else {
		goto L19
	}
L17:
	;
	goto L18
L18:
	;
	v78 = int32(*(*int16)(unsafe.Add(mBase, uint32(l0)+18)))
	v80 = v2
	v81 = v78
	goto L10
L19:
	;
	v50 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
	v51 = F_pg_detoast_datum(m, v50)
	mBase = m.M
	v54 = m.ExcPending
	if v54 != 0 {
		goto L20
	} else {
		goto L21
	}
L20:
	;
	return int32(0)
L21:
	;
	v55 = *(*int32)(unsafe.Add(mBase, uint32(v51)+12))
	F_get_typlenbyvalalign(m, v55, v20+int32(24), v20+int32(88), v20+int32(108))
	mBase = m.M
	v63 = m.ExcPending
	if v63 != 0 {
		goto L20
	} else {
		goto L22
	}
L22:
	;
	v64 = int32(*(*int16)(unsafe.Add(mBase, uint32(v20)+24)))
	v65 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v20)+88)))
	v66 = int32(*(*int8)(unsafe.Add(mBase, uint32(v20)+108)))
	F_deconstruct_array(m, v51, v64, v65, v66, v20+int32(84), v20+int32(80), v20+int32(52))
	mBase = m.M
	v74 = m.ExcPending
	if v74 != 0 {
		goto L20
	} else {
		goto L23
	}
L23:
	;
	v75 = *(*int32)(unsafe.Add(mBase, uint32(v20)+52))
	v80 = v55
	v81 = v75 + int32(1)
	goto L10
L24:
	;
	v85 = int32(1)
	v86 = v83 + v85
	v89 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v83))))
	v91 = v89 & v85
	if v91 != 0 {
		goto L25
	} else {
		goto L26
	}
L25:
	;
	v92 = v86
	goto L27
L26:
	;
	v92 = v83 + int32(4)
	goto L27
L27:
	;
	if v89 == int32(1) {
		goto L33
	} else {
		goto L34
	}
L28:
	;
	v856 = *(*int32)(unsafe.Add(mBase, uint32(v20)+84))
	if v856 != 0 {
		goto L243
	} else {
		goto L244
	}
L29:
	;
	v139 = v92
	v144 = int32(1)
	v151 = v2
	v152 = v2
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
	v119 = *(*int32)(unsafe.Add(mBase, uint32(v83)))
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
	v95 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v86))))
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
	if v91 == int32(0) {
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
	v125 = int32(base.Ui32(v89)>>(uint(v109)%32)) - v109
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
	v153 = int32(*(*int8)(unsafe.Add(mBase, uint32(v139))))
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
	v837 = v822 + int32(1)
	if base.Ui32(v837) < base.Ui32(v134) {
		v139 = v837
		v144 = v827
		v151 = v834
		v152 = v835
		goto L44
	} else {
		goto L242
	}
L47:
	;
	v513 = int32(*(*int8)(unsafe.Add(mBase, uint32(v498))))
	goto L144
L48:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v20)+108)) = v469
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
	v822 = v139
	v827 = v144
	v834 = v151
	v835 = v152
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
	v166 = v139 + int32(1)
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
	v822 = v139
	v827 = v144
	v834 = v151
	v835 = v152
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
	v822 = v166
	v827 = v144
	v834 = v151
	v835 = v152
	goto L46
L69:
	;
	v822 = v166
	v827 = v144
	v834 = v151
	v835 = v152
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
		v289 = v215
		v295 = v216
		goto L79
	default:
		v469 = v215
		v475 = v216
		goto L48
	case 3:
		goto L80
	}
L74:
	;
	v207 = int32(-1)
	v498 = v201
	v499 = v207
	v500 = v207
	v501 = v202
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
	v304 = v289 + int32(1)
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
	v226 = v221
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
	v240 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v226))))
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
		v289 = v226
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
	v247 = v226 + int32(1)
	if base.Ui32(v247) < base.Ui32(v134) {
		v226 = v247
		goto L84
	} else {
		goto L90
	}
L89:
	;
	v469 = v226
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
	F_errmsg(m, int32(_a_F_text_format_0), int32(0))
	mBase = m.M
	v276 = m.ExcPending
	if v276 != 0 {
		goto L20
	} else {
		goto L93
	}
L93:
	;
	F_errhint(m, int32(_a_F_text_format_1), int32(0))
	mBase = m.M
	v280 = m.ExcPending
	if v280 != 0 {
		goto L20
	} else {
		goto L94
	}
L94:
	;
	F_errfinish(m, int32(_a_F_text_format_2), int32(_a_F_text_format_3), int32(_a_F_text_format_4))
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
	v498 = v314
	v499 = v214
	v500 = int32(0)
	v501 = v307
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
	v498 = v325
	v499 = v214
	v500 = v321
	v501 = v307
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
	F_errmsg(m, int32(_a_F_text_format_0), int32(0))
	mBase = m.M
	v348 = m.ExcPending
	if v348 != 0 {
		goto L20
	} else {
		goto L106
	}
L106:
	;
	F_errhint(m, int32(_a_F_text_format_1), int32(0))
	mBase = m.M
	v352 = m.ExcPending
	if v352 != 0 {
		goto L20
	} else {
		goto L107
	}
L107:
	;
	F_errfinish(m, int32(_a_F_text_format_2), int32(_a_F_text_format_5), int32(_a_F_text_format_6))
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
	F_errmsg(m, int32(_a_F_text_format_7), int32(0))
	mBase = m.M
	v368 = m.ExcPending
	if v368 != 0 {
		goto L20
	} else {
		goto L111
	}
L111:
	;
	F_errfinish(m, int32(_a_F_text_format_2), int32(_a_F_text_format_8), int32(_a_F_text_format_4))
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
	F_errmsg(m, int32(_a_F_text_format_0), int32(0))
	mBase = m.M
	v384 = m.ExcPending
	if v384 != 0 {
		goto L20
	} else {
		goto L115
	}
L115:
	;
	F_errhint(m, int32(_a_F_text_format_1), int32(0))
	mBase = m.M
	v388 = m.ExcPending
	if v388 != 0 {
		goto L20
	} else {
		goto L116
	}
L116:
	;
	F_errfinish(m, int32(_a_F_text_format_2), int32(_a_F_text_format_9), int32(_a_F_text_format_4))
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
	F_errmsg(m, int32(_a_F_text_format_0), int32(0))
	mBase = m.M
	v404 = m.ExcPending
	if v404 != 0 {
		goto L20
	} else {
		goto L120
	}
L120:
	;
	F_errhint(m, int32(_a_F_text_format_1), int32(0))
	mBase = m.M
	v408 = m.ExcPending
	if v408 != 0 {
		goto L20
	} else {
		goto L121
	}
L121:
	;
	F_errfinish(m, int32(_a_F_text_format_2), int32(_a_F_text_format_10), int32(_a_F_text_format_4))
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
	F_errmsg(m, int32(_a_F_text_format_11), int32(0))
	mBase = m.M
	v424 = m.ExcPending
	if v424 != 0 {
		goto L20
	} else {
		goto L125
	}
L125:
	;
	F_errfinish(m, int32(_a_F_text_format_2), int32(_a_F_text_format_12), int32(_a_F_text_format_4))
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
	F_errmsg(m, int32(_a_F_text_format_7), int32(0))
	mBase = m.M
	v440 = m.ExcPending
	if v440 != 0 {
		goto L20
	} else {
		goto L129
	}
L129:
	;
	F_errfinish(m, int32(_a_F_text_format_2), int32(_a_F_text_format_13), int32(_a_F_text_format_4))
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
	F_errmsg(m, int32(_a_F_text_format_0), int32(0))
	mBase = m.M
	v456 = m.ExcPending
	if v456 != 0 {
		goto L20
	} else {
		goto L133
	}
L133:
	;
	F_errhint(m, int32(_a_F_text_format_1), int32(0))
	mBase = m.M
	v460 = m.ExcPending
	if v460 != 0 {
		goto L20
	} else {
		goto L134
	}
L134:
	;
	F_errfinish(m, int32(_a_F_text_format_2), int32(_a_F_text_format_14), int32(_a_F_text_format_4))
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
	v498 = v494
	v499 = v214
	v500 = int32(-1)
	v501 = v492
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
	if base.B2i32(v536 != v537) == int32(0) {
		goto L141
	} else {
		goto L150
	}
L144:
	;
	v528 = int32(_a_F_text_format_15)
	v530 = int32(4)
	goto L145
L145:
	;
	v533 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v528))))
	if v533 == v513&int32(255) {
		v590 = v528
		v592 = v530
		goto L142
	} else {
		goto L147
	}
L146:
	;
	goto L143
L147:
	;
	v535 = int32(1)
	v536 = v530 - v535
	v537 = int32(0)
	v540 = v528 + v535
	if v540&int32(3) == v537 {
		goto L143
	} else {
		goto L148
	}
L148:
	;
	if v536 != 0 {
		v528 = v540
		v530 = v536
		goto L145
	} else {
		goto L149
	}
L149:
	;
	goto L146
L150:
	;
	v553 = v513 & int32(255)
	v554 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v540))))
	if base.B2i32(v553 == v554)|base.B2i32(base.Ui32(v536) < base.Ui32(int32(4))) == int32(0) {
		goto L151
	} else {
		goto L152
	}
L151:
	;
	v563 = v540
	v565 = v536
	goto L154
L152:
	;
	v583 = v540
	v585 = v536
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
	if v500 < int32(0) {
		v683 = v501
		v684 = v144
		v686 = v151
		goto L166
	} else {
		goto L167
	}
L166:
	;
	if int32(0) < v499 {
		goto L191
	} else {
		goto L192
	}
L167:
	;
	if v500 != 0 {
		goto L168
	} else {
		goto L169
	}
L168:
	;
	v624 = v500
	goto L170
L169:
	;
	v624 = v144
	goto L170
L170:
	;
	if v81 <= v624 {
		goto L4
	} else {
		goto L171
	}
L171:
	;
	if v47 == int32(0) {
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
	v630 = v32 + v624<<(uint(int32(3))%32)
	v631 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v630)+4)))
	v632 = *(*int32)(unsafe.Add(mBase, uint32(v630)))
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
	v645 = *(*int32)(unsafe.Add(mBase, uint32(v641+v637<<(uint(int32(2))%32))))
	v647 = v645
	v648 = v640
	v649 = v80
	goto L172
L176:
	;
	v647 = v632
	v648 = v631
	v649 = v634
	goto L172
L177:
	;
	v652 = int32(1)
	v653 = v624 + v652
	if v648&v652 != 0 {
		goto L178
	} else {
		goto L179
	}
L178:
	;
	v683 = int32(0)
	v684 = v653
	v686 = v151
	goto L166
L179:
	;
	goto L180
L180:
	;
	switch v649 - int32(21) {
	case 0:
		goto L182
	default:
		goto L181
	case 2:
		v683 = v647
		v684 = v653
		v686 = v151
		goto L166
	}
L181:
	;
	if v649 != v151 {
		goto L183
	} else {
		goto L184
	}
L182:
	;
	v683 = base.I32_extend16_s(v647)
	v684 = v653
	v686 = v151
	goto L166
L183:
	;
	F_getTypeOutputInfo(m, v649, v20+int32(108), v20+int32(104))
	mBase = m.M
	v666 = m.ExcPending
	if v666 != 0 {
		goto L20
	} else {
		goto L186
	}
L184:
	;
	v672 = v151
	goto L185
L185:
	;
	v675 = F_OutputFunctionCall(m, v20+int32(24), v647)
	mBase = m.M
	v676 = m.ExcPending
	if v676 != 0 {
		goto L20
	} else {
		goto L188
	}
L186:
	;
	v667 = *(*int32)(unsafe.Add(mBase, uint32(v20)+108))
	F_fmgr_info(m, v667, v20+int32(24))
	mBase = m.M
	v671 = m.ExcPending
	if v671 != 0 {
		goto L20
	} else {
		goto L187
	}
L187:
	;
	v672 = v649
	goto L185
L188:
	;
	v677 = F_pg_strtoint32(m, v675)
	mBase = m.M
	v678 = m.ExcPending
	if v678 != 0 {
		goto L20
	} else {
		goto L189
	}
L189:
	;
	F_pfree(m, v675)
	mBase = m.M
	v680 = m.ExcPending
	if v680 != 0 {
		goto L20
	} else {
		goto L190
	}
L190:
	;
	v683 = v677
	v684 = v653
	v686 = v672
	goto L166
L191:
	;
	v689 = v499
	goto L193
L192:
	;
	v689 = v684
	goto L193
L193:
	;
	if v81 <= v689 {
		goto L2
	} else {
		goto L194
	}
L194:
	;
	if v47 == int32(0) {
		goto L196
	} else {
		goto L197
	}
L195:
	;
	if v713 == int32(0) {
		goto L1
	} else {
		goto L200
	}
L196:
	;
	v695 = v32 + v689<<(uint(int32(3))%32)
	v696 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v695)+4)))
	v697 = *(*int32)(unsafe.Add(mBase, uint32(v695)))
	v698 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v699 = F_get_fn_expr_argtype(m, v698, v689)
	mBase = m.M
	v700 = m.ExcPending
	if v700 != 0 {
		goto L20
	} else {
		goto L199
	}
L197:
	;
	goto L198
L198:
	;
	v702 = v689 - int32(1)
	v703 = *(*int32)(unsafe.Add(mBase, uint32(v20)+80))
	v705 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v702+v703))))
	v706 = *(*int32)(unsafe.Add(mBase, uint32(v20)+84))
	v710 = *(*int32)(unsafe.Add(mBase, uint32(v706+v702<<(uint(int32(2))%32))))
	v711 = v710
	v712 = v705
	v713 = v80
	goto L195
L199:
	;
	v711 = v697
	v712 = v696
	v713 = v699
	goto L195
L200:
	;
	if v713 != v152 {
		goto L201
	} else {
		goto L202
	}
L201:
	;
	F_getTypeOutputInfo(m, v713, v20+int32(108), v20+int32(104))
	mBase = m.M
	v722 = m.ExcPending
	if v722 != 0 {
		goto L20
	} else {
		goto L204
	}
L202:
	;
	v728 = v152
	goto L203
L203:
	;
	v729 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v498))))
	v731 = v729 - int32(73)
	switch v731 {
	case 0, 3:
		goto L207
	case 1, 2:
		goto L206
	default:
		goto L208
	}
L204:
	;
	v723 = *(*int32)(unsafe.Add(mBase, uint32(v20)+108))
	F_fmgr_info(m, v723, v20+int32(52))
	mBase = m.M
	v727 = m.ExcPending
	if v727 != 0 {
		goto L20
	} else {
		goto L205
	}
L205:
	;
	v728 = v713
	goto L203
L206:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v797 = m.ExcPending
	if v797 != 0 {
		goto L20
	} else {
		goto L236
	}
L207:
	;
	v734 = int32(1)
	v735 = v689 + v734
	if v712&v734 != 0 {
		goto L210
	} else {
		goto L211
	}
L208:
	;
	if v729 != int32(115) {
		goto L206
	} else {
		goto L209
	}
L209:
	;
	goto L207
L210:
	;
	switch v731 {
	case 0:
		goto L213
	case 1, 2:
		v822 = v498
		v827 = v735
		v834 = v686
		v835 = v728
		goto L46
	case 3:
		goto L214
	default:
		goto L215
	}
L211:
	;
	goto L212
L212:
	;
	v768 = F_OutputFunctionCall(m, v20+int32(52), v711)
	mBase = m.M
	v769 = m.ExcPending
	if v769 != 0 {
		goto L20
	} else {
		goto L223
	}
L213:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v753 = m.ExcPending
	if v753 != 0 {
		goto L20
	} else {
		goto L219
	}
L214:
	;
	F_text_format_append_string(m, v20+int32(88), int32(_a_F_text_format_16), v504, v683)
	mBase = m.M
	v749 = m.ExcPending
	if v749 != 0 {
		goto L20
	} else {
		goto L218
	}
L215:
	;
	if v729 != int32(115) {
		v822 = v498
		v827 = v735
		v834 = v686
		v835 = v728
		goto L46
	} else {
		goto L216
	}
L216:
	;
	F_text_format_append_string(m, v20+int32(88), int32(_a_F_text_format_17), v504, v683)
	mBase = m.M
	v744 = m.ExcPending
	if v744 != 0 {
		goto L20
	} else {
		goto L217
	}
L217:
	;
	v822 = v498
	v827 = v735
	v834 = v686
	v835 = v728
	goto L46
L218:
	;
	v822 = v498
	v827 = v735
	v834 = v686
	v835 = v728
	goto L46
L219:
	;
	F_errcode(m, int32(67108994))
	mBase = m.M
	v756 = m.ExcPending
	if v756 != 0 {
		goto L20
	} else {
		goto L220
	}
L220:
	;
	F_errmsg(m, int32(_a_F_text_format_18), int32(0))
	mBase = m.M
	v760 = m.ExcPending
	if v760 != 0 {
		goto L20
	} else {
		goto L221
	}
L221:
	;
	F_errfinish(m, int32(_a_F_text_format_2), int32(_a_F_text_format_19), int32(_a_F_text_format_20))
	mBase = m.M
	v765 = m.ExcPending
	if v765 != 0 {
		goto L20
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
	switch v731 {
	case 0:
		goto L226
	default:
		goto L224
	case 3:
		goto L225
	}
L224:
	;
	F_text_format_append_string(m, v20+int32(88), v768, v504, v683)
	mBase = m.M
	v791 = m.ExcPending
	if v791 != 0 {
		goto L20
	} else {
		goto L234
	}
L225:
	;
	v780 = F_quote_literal_cstr(m, v768)
	mBase = m.M
	v781 = m.ExcPending
	if v781 != 0 {
		goto L20
	} else {
		goto L230
	}
L226:
	;
	v772 = F_quote_identifier(m, v768)
	mBase = m.M
	v773 = m.ExcPending
	if v773 != 0 {
		goto L20
	} else {
		goto L227
	}
L227:
	;
	F_text_format_append_string(m, v20+int32(88), v772, v504, v683)
	mBase = m.M
	v775 = m.ExcPending
	if v775 != 0 {
		goto L20
	} else {
		goto L228
	}
L228:
	;
	F_pfree(m, v768)
	mBase = m.M
	v777 = m.ExcPending
	if v777 != 0 {
		goto L20
	} else {
		goto L229
	}
L229:
	;
	v822 = v498
	v827 = v735
	v834 = v686
	v835 = v728
	goto L46
L230:
	;
	F_text_format_append_string(m, v20+int32(88), v780, v504, v683)
	mBase = m.M
	v783 = m.ExcPending
	if v783 != 0 {
		goto L20
	} else {
		goto L231
	}
L231:
	;
	F_pfree(m, v780)
	mBase = m.M
	v785 = m.ExcPending
	if v785 != 0 {
		goto L20
	} else {
		goto L232
	}
L232:
	;
	F_pfree(m, v768)
	mBase = m.M
	v787 = m.ExcPending
	if v787 != 0 {
		goto L20
	} else {
		goto L233
	}
L233:
	;
	v822 = v498
	v827 = v735
	v834 = v686
	v835 = v728
	goto L46
L234:
	;
	F_pfree(m, v768)
	mBase = m.M
	v793 = m.ExcPending
	if v793 != 0 {
		goto L20
	} else {
		goto L235
	}
L235:
	;
	v822 = v498
	v827 = v735
	v834 = v686
	v835 = v728
	goto L46
L236:
	;
	F_errcode(m, int32(50856066))
	mBase = m.M
	v800 = m.ExcPending
	if v800 != 0 {
		goto L20
	} else {
		goto L237
	}
L237:
	;
	v801 = F_pg_mblen_range(m, v498, v134)
	mBase = m.M
	v802 = m.ExcPending
	if v802 != 0 {
		goto L20
	} else {
		goto L238
	}
L238:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v20)+20)) = v498
	*(*int32)(unsafe.Add(mBase, uint32(v20)+16)) = v801
	F_errmsg(m, int32(_a_F_text_format_21), v20+int32(16))
	mBase = m.M
	v809 = m.ExcPending
	if v809 != 0 {
		goto L20
	} else {
		goto L239
	}
L239:
	;
	F_errhint(m, int32(_a_F_text_format_1), int32(0))
	mBase = m.M
	v813 = m.ExcPending
	if v813 != 0 {
		goto L20
	} else {
		goto L240
	}
L240:
	;
	F_errfinish(m, int32(_a_F_text_format_2), int32(_a_F_text_format_22), int32(_a_F_text_format_6))
	mBase = m.M
	v818 = m.ExcPending
	if v818 != 0 {
		goto L20
	} else {
		goto L241
	}
L241:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L242:
	;
	goto L45
L243:
	;
	F_pfree(m, v856)
	mBase = m.M
	v858 = m.ExcPending
	if v858 != 0 {
		goto L20
	} else {
		goto L246
	}
L244:
	;
	goto L245
L245:
	;
	v859 = *(*int32)(unsafe.Add(mBase, uint32(v20)+80))
	if v859 != 0 {
		goto L247
	} else {
		goto L248
	}
L246:
	;
	goto L245
L247:
	;
	F_pfree(m, v859)
	mBase = m.M
	v861 = m.ExcPending
	if v861 != 0 {
		goto L20
	} else {
		goto L250
	}
L248:
	;
	goto L249
L249:
	;
	v862 = *(*int32)(unsafe.Add(mBase, uint32(v20)+88))
	v863 = *(*int32)(unsafe.Add(mBase, uint32(v20)+92))
	v865 = v863 + int32(4)
	v866 = F_palloc(m, v865)
	mBase = m.M
	v867 = m.ExcPending
	if v867 != 0 {
		goto L20
	} else {
		goto L251
	}
L250:
	;
	goto L249
L251:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v866))) = v865 << (uint(int32(2)) % 32)
	if v863 != 0 {
		goto L252
	} else {
		goto L253
	}
L252:
	;
	base.MemoryCopy(m, v866+int32(4), v862, v863)
	goto L254
L253:
	;
	goto L254
L254:
	;
	v874 = *(*int32)(unsafe.Add(mBase, uint32(v20)+88))
	F_pfree(m, v874)
	mBase = m.M
	v876 = m.ExcPending
	if v876 != 0 {
		goto L20
	} else {
		goto L255
	}
L255:
	;
	v880 = v866
	goto L6
L256:
	;
	F_errcode(m, int32(50856066))
	mBase = m.M
	v904 = m.ExcPending
	if v904 != 0 {
		goto L20
	} else {
		goto L257
	}
L257:
	;
	v905 = F_pg_mblen_range(m, v498, v134)
	mBase = m.M
	v906 = m.ExcPending
	if v906 != 0 {
		goto L20
	} else {
		goto L258
	}
L258:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v20)+4)) = v498
	*(*int32)(unsafe.Add(mBase, uint32(v20))) = v905
	F_errmsg(m, int32(_a_F_text_format_21), v20)
	mBase = m.M
	v911 = m.ExcPending
	if v911 != 0 {
		goto L20
	} else {
		goto L259
	}
L259:
	;
	F_errhint(m, int32(_a_F_text_format_1), int32(0))
	mBase = m.M
	v915 = m.ExcPending
	if v915 != 0 {
		goto L20
	} else {
		goto L260
	}
L260:
	;
	F_errfinish(m, int32(_a_F_text_format_2), int32(_a_F_text_format_23), int32(_a_F_text_format_6))
	mBase = m.M
	v920 = m.ExcPending
	if v920 != 0 {
		goto L20
	} else {
		goto L261
	}
L261:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L262:
	;
	F_errcode(m, int32(50856066))
	mBase = m.M
	v927 = m.ExcPending
	if v927 != 0 {
		goto L20
	} else {
		goto L263
	}
L263:
	;
	F_errmsg(m, int32(_a_F_text_format_24), int32(0))
	mBase = m.M
	v931 = m.ExcPending
	if v931 != 0 {
		goto L20
	} else {
		goto L264
	}
L264:
	;
	F_errfinish(m, int32(_a_F_text_format_2), int32(_a_F_text_format_25), int32(_a_F_text_format_6))
	mBase = m.M
	v936 = m.ExcPending
	if v936 != 0 {
		goto L20
	} else {
		goto L265
	}
L265:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L266:
	;
	F_errmsg_internal(m, int32(_a_F_text_format_26), int32(0))
	mBase = m.M
	v944 = m.ExcPending
	if v944 != 0 {
		goto L20
	} else {
		goto L267
	}
L267:
	;
	F_errfinish(m, int32(_a_F_text_format_2), int32(_a_F_text_format_27), int32(_a_F_text_format_6))
	mBase = m.M
	v949 = m.ExcPending
	if v949 != 0 {
		goto L20
	} else {
		goto L268
	}
L268:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L269:
	;
	F_errcode(m, int32(50856066))
	mBase = m.M
	v956 = m.ExcPending
	if v956 != 0 {
		goto L20
	} else {
		goto L270
	}
L270:
	;
	F_errmsg(m, int32(_a_F_text_format_24), int32(0))
	mBase = m.M
	v960 = m.ExcPending
	if v960 != 0 {
		goto L20
	} else {
		goto L271
	}
L271:
	;
	F_errfinish(m, int32(_a_F_text_format_2), int32(_a_F_text_format_28), int32(_a_F_text_format_6))
	mBase = m.M
	v965 = m.ExcPending
	if v965 != 0 {
		goto L20
	} else {
		goto L272
	}
L272:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L273:
	;
	F_errmsg_internal(m, int32(_a_F_text_format_26), int32(0))
	mBase = m.M
	v973 = m.ExcPending
	if v973 != 0 {
		goto L20
	} else {
		goto L274
	}
L274:
	;
	F_errfinish(m, int32(_a_F_text_format_2), int32(_a_F_text_format_29), int32(_a_F_text_format_6))
	mBase = m.M
	v978 = m.ExcPending
	if v978 != 0 {
		goto L20
	} else {
		goto L275
	}
L275:
	;
	base.Wasm_trap_unreachable()
	for {
	}
}
func F_text_larger(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v8 int32
	_ = v8
	var v9 int32
	_ = v9
	var v12 int32
	_ = v12
	var v14 int32
	_ = v14
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
	var v22 int32
	_ = v22
	var v24 int32
	_ = v24
	var v25 int32
	_ = v25
	var v31 int32
	_ = v31
	var v34 int32
	_ = v34
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
	var v59 int32
	_ = v59
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
	var v93 int32
	_ = v93
	v8 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v9 = F_pg_detoast_datum_packed(m, v8)
	mBase = m.M
	v12 = m.ExcPending
	if v12 != 0 {
		return int32(0)
	} else {
		v14 = v9 + int32(1)
		v15 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
		v16 = F_pg_detoast_datum_packed(m, v15)
		mBase = m.M
		v17 = m.ExcPending
		if v17 != 0 {
			return int32(0)
		} else {
			v18 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v16))))
			v19 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
			v22 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v9))))
			v24 = v22 & int32(1)
			if v24 != 0 {
				v25 = v14
			} else {
				v25 = v9 + int32(4)
			}
			if v22 == int32(1) {
				v31 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v14))))
				if v31 == int32(18) {
					v34 = int32(16)
				} else {
					v34 = int32(0)
				}
				if base.Ui32((v31-int32(1))&int32(255)) < base.Ui32(int32(3)) {
					v41 = int32(4)
				} else {
					v41 = v34
				}
				v52 = v41
			} else {
				v42 = int32(1)
				if v24 != 0 {
					v52 = int32(base.Ui32(v22)>>(uint(v42)%32)) - v42
				} else {
					v46 = *(*int32)(unsafe.Add(mBase, uint32(v9)))
					v52 = int32(base.Ui32(v46)>>(uint(int32(2))%32)) - int32(4)
				}
			}
			v53 = int32(1)
			v54 = v16 + v53
			if v18&v53 != 0 {
				v59 = v54
			} else {
				v59 = v16 + int32(4)
			}
			if v18 == int32(1) {
				v65 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v54))))
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
				if v18&v76 != 0 {
					v88 = int32(base.Ui32(v18)>>(uint(v76)%32)) - v76
				} else {
					v82 = *(*int32)(unsafe.Add(mBase, uint32(v16)))
					v88 = int32(base.Ui32(v82)>>(uint(int32(2))%32)) - int32(4)
				}
			}
			v89 = F_varstr_cmp(m, v25, v52, v59, v88, v19)
			mBase = m.M
			v90 = m.ExcPending
			if v90 != 0 {
				return int32(0)
			} else {
				if int32(0) < v89 {
					v93 = v9
				} else {
					v93 = v16
				}
				return v93
			}
		}
	}
}
func F_text_le(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v9 int32
	_ = v9
	var v10 int32
	_ = v10
	var v13 int32
	_ = v13
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
	var v20 int32
	_ = v20
	var v23 int32
	_ = v23
	var v25 int32
	_ = v25
	var v26 int32
	_ = v26
	var v32 int32
	_ = v32
	var v35 int32
	_ = v35
	var v42 int32
	_ = v42
	var v43 int32
	_ = v43
	var v47 int32
	_ = v47
	var v53 int32
	_ = v53
	var v54 int32
	_ = v54
	var v55 int32
	_ = v55
	var v60 int32
	_ = v60
	var v66 int32
	_ = v66
	var v69 int32
	_ = v69
	var v76 int32
	_ = v76
	var v77 int32
	_ = v77
	var v83 int32
	_ = v83
	var v89 int32
	_ = v89
	var v90 int32
	_ = v90
	var v91 int32
	_ = v91
	var v92 int32
	_ = v92
	var v95 int32
	_ = v95
	var v96 int32
	_ = v96
	var v99 int32
	_ = v99
	v9 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v10 = F_pg_detoast_datum_packed(m, v9)
	mBase = m.M
	v13 = m.ExcPending
	if v13 != 0 {
		return int32(0)
	} else {
		v15 = v10 + int32(1)
		v16 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
		v17 = F_pg_detoast_datum_packed(m, v16)
		mBase = m.M
		v18 = m.ExcPending
		if v18 != 0 {
			return int32(0)
		} else {
			v19 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v17))))
			v20 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
			v23 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v10))))
			v25 = v23 & int32(1)
			if v25 != 0 {
				v26 = v15
			} else {
				v26 = v10 + int32(4)
			}
			if v23 == int32(1) {
				v32 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v15))))
				if v32 == int32(18) {
					v35 = int32(16)
				} else {
					v35 = int32(0)
				}
				if base.Ui32((v32-int32(1))&int32(255)) < base.Ui32(int32(3)) {
					v42 = int32(4)
				} else {
					v42 = v35
				}
				v53 = v42
			} else {
				v43 = int32(1)
				if v25 != 0 {
					v53 = int32(base.Ui32(v23)>>(uint(v43)%32)) - v43
				} else {
					v47 = *(*int32)(unsafe.Add(mBase, uint32(v10)))
					v53 = int32(base.Ui32(v47)>>(uint(int32(2))%32)) - int32(4)
				}
			}
			v54 = int32(1)
			v55 = v17 + v54
			if v19&v54 != 0 {
				v60 = v55
			} else {
				v60 = v17 + int32(4)
			}
			if v19 == int32(1) {
				v66 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v55))))
				if v66 == int32(18) {
					v69 = int32(16)
				} else {
					v69 = int32(0)
				}
				if base.Ui32((v66-int32(1))&int32(255)) < base.Ui32(int32(3)) {
					v76 = int32(4)
				} else {
					v76 = v69
				}
				v89 = v76
			} else {
				v77 = int32(1)
				if v19&v77 != 0 {
					v89 = int32(base.Ui32(v19)>>(uint(v77)%32)) - v77
				} else {
					v83 = *(*int32)(unsafe.Add(mBase, uint32(v17)))
					v89 = int32(base.Ui32(v83)>>(uint(int32(2))%32)) - int32(4)
				}
			}
			v90 = F_varstr_cmp(m, v26, v53, v60, v89, v20)
			mBase = m.M
			v91 = m.ExcPending
			if v91 != 0 {
				return int32(0)
			} else {
				v92 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
				if v92 != v10 {
					F_pfree(m, v10)
					mBase = m.M
					v95 = m.ExcPending
					if v95 != 0 {
						return int32(0)
					} else {
						v96 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
						if v96 != v17 {
							F_pfree(m, v17)
							mBase = m.M
							v99 = m.ExcPending
							if v99 != 0 {
								return int32(0)
							} else {
								return base.B2i32(v90 <= int32(0))
							}
						} else {
							return base.B2i32(v90 <= int32(0))
						}
					}
				} else {
					v96 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
					if v96 != v17 {
						F_pfree(m, v17)
						mBase = m.M
						v99 = m.ExcPending
						if v99 != 0 {
							return int32(0)
						} else {
							return base.B2i32(v90 <= int32(0))
						}
					} else {
						return base.B2i32(v90 <= int32(0))
					}
				}
			}
		}
	}
}
func F_text_left(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v7 int32
	_ = v7
	var v8 int32
	_ = v8
	var v11 int32
	_ = v11
	var v14 int32
	_ = v14
	var v15 int32
	_ = v15
	var v16 int32
	_ = v16
	var v19 int32
	_ = v19
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
	var v70 int32
	_ = v70
	var v71 int32
	_ = v71
	var v72 int32
	_ = v72
	v7 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v8 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
	if v8 < int32(0) {
		v11 = F_pg_detoast_datum_packed(m, v7)
		mBase = m.M
		v14 = m.ExcPending
		if v14 != 0 {
			return int32(0)
		} else {
			v15 = int32(1)
			v16 = v11 + v15
			v19 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v11))))
			v21 = v19 & v15
			if v21 != 0 {
				v22 = v16
			} else {
				v22 = v11 + int32(4)
			}
			if v19 == int32(1) {
				v28 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v16))))
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
				if v21 != 0 {
					v49 = int32(base.Ui32(v19)>>(uint(v39)%32)) - v39
				} else {
					v43 = *(*int32)(unsafe.Add(mBase, uint32(v11)))
					v49 = int32(base.Ui32(v43)>>(uint(int32(2))%32)) - int32(4)
				}
			}
			v50 = F_pg_mbstrlen_with_len(m, v22, v49)
			mBase = m.M
			v51 = m.ExcPending
			if v51 != 0 {
				return int32(0)
			} else {
				v53 = F_pg_mbcharcliplen(m, v22, v49, v50+v8)
				mBase = m.M
				v54 = m.ExcPending
				if v54 != 0 {
					return int32(0)
				} else {
					v56 = v53 + int32(4)
					v57 = F_palloc(m, v56)
					mBase = m.M
					v58 = m.ExcPending
					if v58 != 0 {
						return int32(0)
					} else {
						*(*int32)(unsafe.Add(mBase, uint32(v57))) = v56 << (uint(int32(2)) % 32)
						if v53 == int32(0) {
							v72 = v57
							return v72
						} else {
							base.MemoryCopy(m, v57+int32(4), v22, v53)
							return v57
						}
					}
				}
			}
		}
	} else {
		v70 = F_text_substring(m, v7, int32(1), v8, int32(0))
		mBase = m.M
		v71 = m.ExcPending
		if v71 != 0 {
			return int32(0)
		} else {
			v72 = v70
			return v72
		}
	}
}
func F_text_name(m *base.Module, l0 int32) int32 {
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
	v4 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v5 = F_pg_detoast_datum_packed(m, v4)
	mBase = m.M
	v8 = m.ExcPending
	if v8 != 0 {
		return int32(0)
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
				return int32(0)
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
				return v51
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
					return int32(0)
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
					return v51
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
					return int32(0)
				} else {
					v49 = v47
					v51 = F_palloc0(m, int32(64))
					mBase = m.M
					v52 = m.ExcPending
					if v52 != 0 {
						return int32(0)
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
						return v51
					}
				}
			}
		}
	}
}
func F_text_substr_no_len(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v2 int32
	_ = v2
	var v3 int32
	_ = v3
	var v6 int32
	_ = v6
	var v9 int32
	_ = v9
	v2 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v3 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
	v6 = F_text_substring(m, v2, v3, int32(-1), int32(1))
	mBase = m.M
	v9 = m.ExcPending
	if v9 != 0 {
		return int32(0)
	} else {
		return v6
	}
}
