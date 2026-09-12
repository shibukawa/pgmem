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
	var v21 int32
	_ = v21
	var v23 int32
	_ = v23
	var v25 int32
	_ = v25
	var v27 int32
	_ = v27
	var v30 int32
	_ = v30
	var v31 int32
	_ = v31
	var v32 int32
	_ = v32
	var v35 int32
	_ = v35
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
	var v47 int32
	_ = v47
	var v49 int32
	_ = v49
	var v51 int32
	_ = v51
	var v53 int32
	_ = v53
	var v58 int32
	_ = v58
	v4 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	switch v4 {
	case 0:
		v5 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
		v6 = *(*int32)(unsafe.Add(mBase, uint32(v5)+4))
		v9 = *(*int32)(unsafe.Add(mBase, uint32(v5)+8))
		if v9 <= v6+int32(1) {
			F_appendStringInfoChar(m, v5, int32(10))
			mBase = m.M
			v13 = m.ExcPending
			if v13 != 0 {
				return
			} else {
				F_CopySendEndOfRow(m, l0)
				mBase = m.M
				v15 = m.ExcPending
				if v15 != 0 {
					return
				} else {
					return
				}
			}
		} else {
			v16 = *(*int32)(unsafe.Add(mBase, uint32(v5)))
			v18 = int32(10)
			*(*uint8)(unsafe.Add(mBase, uint32(v16+v6))) = uint8(v18)
			v20 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
			v21 = *(*int32)(unsafe.Add(mBase, uint32(v20)+4))
			v23 = v21 + int32(1)
			*(*int32)(unsafe.Add(mBase, uint32(v20)+4)) = v23
			v25 = *(*int32)(unsafe.Add(mBase, uint32(v20)))
			v27 = int32(0)
			*(*uint8)(unsafe.Add(mBase, uint32(v25+v23))) = uint8(v27)
			F_CopySendEndOfRow(m, l0)
			mBase = m.M
			v30 = m.ExcPending
			if v30 != 0 {
				return
			} else {
				return
			}
		}
	case 1:
		v31 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
		v32 = *(*int32)(unsafe.Add(mBase, uint32(v31)+4))
		v35 = *(*int32)(unsafe.Add(mBase, uint32(v31)+8))
		if v35 <= v32+int32(1) {
			F_appendStringInfoChar(m, v31, int32(10))
			mBase = m.M
			v39 = m.ExcPending
			if v39 != 0 {
				return
			} else {
				F_CopySendEndOfRow(m, l0)
				mBase = m.M
				v41 = m.ExcPending
				if v41 != 0 {
					return
				} else {
					return
				}
			}
		} else {
			v42 = *(*int32)(unsafe.Add(mBase, uint32(v31)))
			v44 = int32(10)
			*(*uint8)(unsafe.Add(mBase, uint32(v42+v32))) = uint8(v44)
			v46 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
			v47 = *(*int32)(unsafe.Add(mBase, uint32(v46)+4))
			v49 = v47 + int32(1)
			*(*int32)(unsafe.Add(mBase, uint32(v46)+4)) = v49
			v51 = *(*int32)(unsafe.Add(mBase, uint32(v46)))
			v53 = int32(0)
			*(*uint8)(unsafe.Add(mBase, uint32(v51+v49))) = uint8(v53)
			F_CopySendEndOfRow(m, l0)
			mBase = m.M
			v58 = m.ExcPending
			if v58 != 0 {
				return
			} else {
				return
			}
		}
	default:
		F_CopySendEndOfRow(m, l0)
		mBase = m.M
		v58 = m.ExcPending
		if v58 != 0 {
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
	var v5 int32
	_ = v5
	var v8 int32
	_ = v8
	var v10 int32
	_ = v10
	var v11 int32
	_ = v11
	var v14 int32
	_ = v14
	var v15 int32
	_ = v15
	var v16 int32
	_ = v16
	var v17 int32
	_ = v17
	var v20 int32
	_ = v20
	var v22 int32
	_ = v22
	var v31 int32
	_ = v31
	var v34 int32
	_ = v34
	var v35 int32
	_ = v35
	var v41 int32
	_ = v41
	var v47 int32
	_ = v47
	var v50 int32
	_ = v50
	var v51 int32
	_ = v51
	var v52 int32
	_ = v52
	var v54 int32
	_ = v54
	var v57 int32
	_ = v57
	var v59 int32
	_ = v59
	var v60 int32
	_ = v60
	var v62 int32
	_ = v62
	var v66 int32
	_ = v66
	var v68 int32
	_ = v68
	var v69 int32
	_ = v69
	var v70 int32
	_ = v70
	var v77 int32
	_ = v77
	v5 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+24)))
	if v5 == int32(1) {
		v8 = int32(1)
		*(*uint8)(unsafe.Add(mBase, uint32(l0)+16)) = uint8(v8)
		v77 = int32(0)
		return v77
	} else {
		v10 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
		v11 = F_pg_detoast_datum_packed(m, v10)
		mBase = m.M
		v14 = m.ExcPending
		if v14 != 0 {
			return int32(0)
		} else {
			v15 = F_pg_detoast_datum_packed(m, v11)
			mBase = m.M
			v16 = m.ExcPending
			if v16 != 0 {
				return int32(0)
			} else {
				v17 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v15))))
				if v17 == int32(1) {
					v20 = int32(4)
					v22 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v15)+1)))
					if v22&int32(254) == int32(2) {
						v31 = v20
					} else {
						v31 = base.B2i32(v22 == int32(18)) << (uint(v20) % 32)
					}
					if v22 == int32(1) {
						v34 = v20
					} else {
						v34 = v31
					}
					v47 = v34
				} else {
					v35 = int32(1)
					if v17&v35 != 0 {
						v47 = int32(base.Ui32(v17)>>(uint(v35)%32)) - v35
					} else {
						v41 = *(*int32)(unsafe.Add(mBase, uint32(v15)))
						v47 = int32(base.Ui32(v41)>>(uint(int32(2))%32)) - int32(4)
					}
				}
				v50 = F_palloc(m, v47+int32(1))
				mBase = m.M
				v51 = m.ExcPending
				if v51 != 0 {
					return int32(0)
				} else {
					v52 = int32(1)
					v54 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v15))))
					if v54&v52 != 0 {
						v57 = v52
					} else {
						v57 = int32(4)
					}
					if v47 != 0 {
						v59 = F__emscripten_memcpy_bulkmem(m, v50, v15+v57, v47)
						mBase = m.M
						v60 = v59
					} else {
						v60 = v50
					}
					v62 = int32(0)
					*(*uint8)(unsafe.Add(mBase, uint32(v47+v60))) = uint8(v62)
					if v15 != v11 {
						F_pfree(m, v15)
						mBase = m.M
						v66 = m.ExcPending
						if v66 != 0 {
							return int32(0)
						} else {
							v68 = F_concat_internal(m, v60, int32(1), l0)
							mBase = m.M
							v69 = m.ExcPending
							if v69 != 0 {
								return int32(0)
							} else {
								if v68 != 0 {
									v77 = v68
								} else {
									v70 = int32(1)
									*(*uint8)(unsafe.Add(mBase, uint32(l0)+16)) = uint8(v70)
									v77 = int32(0)
								}
								return v77
							}
						}
					} else {
						v68 = F_concat_internal(m, v60, int32(1), l0)
						mBase = m.M
						v69 = m.ExcPending
						if v69 != 0 {
							return int32(0)
						} else {
							if v68 != 0 {
								v77 = v68
							} else {
								v70 = int32(1)
								*(*uint8)(unsafe.Add(mBase, uint32(l0)+16)) = uint8(v70)
								v77 = int32(0)
							}
							return v77
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
	var v21 int32
	_ = v21
	var v23 int32
	_ = v23
	var v29 int32
	_ = v29
	var v32 int32
	_ = v32
	var v35 int32
	_ = v35
	var v36 int32
	_ = v36
	var v37 int32
	_ = v37
	var v40 int32
	_ = v40
	var v43 int32
	_ = v43
	var v46 int32
	_ = v46
	var v48 int32
	_ = v48
	var v50 int32
	_ = v50
	var v51 int32
	_ = v51
	var v53 int32
	_ = v53
	var v54 int32
	_ = v54
	var v57 int32
	_ = v57
	var v58 int32
	_ = v58
	var v66 int32
	_ = v66
	var v67 int32
	_ = v67
	var v68 int32
	_ = v68
	var v69 int32
	_ = v69
	var v77 int32
	_ = v77
	var v78 int32
	_ = v78
	var v81 int32
	_ = v81
	var v83 int32
	_ = v83
	var v87 int32
	_ = v87
	var v88 int32
	_ = v88
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
	var v97 int32
	_ = v97
	var v98 int32
	_ = v98
	var v101 int32
	_ = v101
	var v103 int32
	_ = v103
	var v112 int32
	_ = v112
	var v115 int32
	_ = v115
	var v116 int32
	_ = v116
	var v120 int32
	_ = v120
	var v126 int32
	_ = v126
	var v130 int32
	_ = v130
	var v131 int32
	_ = v131
	var v136 int32
	_ = v136
	var v140 int32
	_ = v140
	var v147 int32
	_ = v147
	var v148 int32
	_ = v148
	var v154 int32
	_ = v154
	var v157 int32
	_ = v157
	var v158 int32
	_ = v158
	var v165 int32
	_ = v165
	var v167 int32
	_ = v167
	var v169 int32
	_ = v169
	var v172 int32
	_ = v172
	var v173 int32
	_ = v173
	var v181 int32
	_ = v181
	var v182 int32
	_ = v182
	var v184 int32
	_ = v184
	var v186 int32
	_ = v186
	var v188 int32
	_ = v188
	var v190 int32
	_ = v190
	var v192 int32
	_ = v192
	var v200 int32
	_ = v200
	var v201 int32
	_ = v201
	var v202 int32
	_ = v202
	var v203 int32
	_ = v203
	var v204 int32
	_ = v204
	var v208 int32
	_ = v208
	var v213 int32
	_ = v213
	var v215 int32
	_ = v215
	var v216 int32
	_ = v216
	var v217 int32
	_ = v217
	var v218 int32
	_ = v218
	var v222 int32
	_ = v222
	var v226 int32
	_ = v226
	var v244 int32
	_ = v244
	var v247 int32
	_ = v247
	var v251 int32
	_ = v251
	var v276 int32
	_ = v276
	var v279 int32
	_ = v279
	var v283 int32
	_ = v283
	var v287 int32
	_ = v287
	var v292 int32
	_ = v292
	var v295 int32
	_ = v295
	var v302 int32
	_ = v302
	var v314 int32
	_ = v314
	var v317 int32
	_ = v317
	var v322 int32
	_ = v322
	var v323 int32
	_ = v323
	var v324 int32
	_ = v324
	var v328 int32
	_ = v328
	var v331 int32
	_ = v331
	var v335 int32
	_ = v335
	var v337 int32
	_ = v337
	var v340 int32
	_ = v340
	var v342 int32
	_ = v342
	var v344 int32
	_ = v344
	var v346 int32
	_ = v346
	var v351 int32
	_ = v351
	var v354 int32
	_ = v354
	var v358 int32
	_ = v358
	var v362 int32
	_ = v362
	var v367 int32
	_ = v367
	var v371 int32
	_ = v371
	var v374 int32
	_ = v374
	var v378 int32
	_ = v378
	var v383 int32
	_ = v383
	var v387 int32
	_ = v387
	var v390 int32
	_ = v390
	var v394 int32
	_ = v394
	var v398 int32
	_ = v398
	var v403 int32
	_ = v403
	var v407 int32
	_ = v407
	var v410 int32
	_ = v410
	var v414 int32
	_ = v414
	var v418 int32
	_ = v418
	var v423 int32
	_ = v423
	var v427 int32
	_ = v427
	var v430 int32
	_ = v430
	var v434 int32
	_ = v434
	var v439 int32
	_ = v439
	var v443 int32
	_ = v443
	var v446 int32
	_ = v446
	var v450 int32
	_ = v450
	var v455 int32
	_ = v455
	var v459 int32
	_ = v459
	var v462 int32
	_ = v462
	var v466 int32
	_ = v466
	var v470 int32
	_ = v470
	var v475 int32
	_ = v475
	var v478 int32
	_ = v478
	var v485 int32
	_ = v485
	var v501 int32
	_ = v501
	var v502 int32
	_ = v502
	var v503 int32
	_ = v503
	var v505 int32
	_ = v505
	var v507 int32
	_ = v507
	var v510 int32
	_ = v510
	var v511 int32
	_ = v511
	var v513 int32
	_ = v513
	var v515 int32
	_ = v515
	var v517 int32
	_ = v517
	var v529 int32
	_ = v529
	var v543 int32
	_ = v543
	var v545 int32
	_ = v545
	var v548 int32
	_ = v548
	var v550 int32
	_ = v550
	var v551 int32
	_ = v551
	var v552 int32
	_ = v552
	var v555 int32
	_ = v555
	var v567 int32
	_ = v567
	var v577 int32
	_ = v577
	var v579 int32
	_ = v579
	var v583 int32
	_ = v583
	var v584 int32
	_ = v584
	var v587 int32
	_ = v587
	var v591 int32
	_ = v591
	var v592 int32
	_ = v592
	var v594 int32
	_ = v594
	var v597 int32
	_ = v597
	var v599 int32
	_ = v599
	var v604 int32
	_ = v604
	var v606 int32
	_ = v606
	var v611 int32
	_ = v611
	var v613 int32
	_ = v613
	var v616 int32
	_ = v616
	var v618 int32
	_ = v618
	var v621 int32
	_ = v621
	var v633 int32
	_ = v633
	var v638 int32
	_ = v638
	var v644 int32
	_ = v644
	var v645 int32
	_ = v645
	var v646 int32
	_ = v646
	var v647 int32
	_ = v647
	var v648 int32
	_ = v648
	var v649 int32
	_ = v649
	var v651 int32
	_ = v651
	var v652 int32
	_ = v652
	var v654 int32
	_ = v654
	var v655 int32
	_ = v655
	var v659 int32
	_ = v659
	var v660 int32
	_ = v660
	var v662 int32
	_ = v662
	var v663 int32
	_ = v663
	var v666 int32
	_ = v666
	var v667 int32
	_ = v667
	var v680 int32
	_ = v680
	var v681 int32
	_ = v681
	var v685 int32
	_ = v685
	var v686 int32
	_ = v686
	var v689 int32
	_ = v689
	var v690 int32
	_ = v690
	var v691 int32
	_ = v691
	var v692 int32
	_ = v692
	var v694 int32
	_ = v694
	var v696 int32
	_ = v696
	var v697 int32
	_ = v697
	var v700 int32
	_ = v700
	var v703 int32
	_ = v703
	var v709 int32
	_ = v709
	var v710 int32
	_ = v710
	var v711 int32
	_ = v711
	var v712 int32
	_ = v712
	var v713 int32
	_ = v713
	var v714 int32
	_ = v714
	var v716 int32
	_ = v716
	var v717 int32
	_ = v717
	var v719 int32
	_ = v719
	var v720 int32
	_ = v720
	var v724 int32
	_ = v724
	var v726 int32
	_ = v726
	var v727 int32
	_ = v727
	var v728 int32
	_ = v728
	var v737 int32
	_ = v737
	var v738 int32
	_ = v738
	var v742 int32
	_ = v742
	var v743 int32
	_ = v743
	var v744 int32
	_ = v744
	var v746 int32
	_ = v746
	var v749 int32
	_ = v749
	var v750 int32
	_ = v750
	var v759 int32
	_ = v759
	var v764 int32
	_ = v764
	var v768 int32
	_ = v768
	var v771 int32
	_ = v771
	var v775 int32
	_ = v775
	var v780 int32
	_ = v780
	var v783 int32
	_ = v783
	var v784 int32
	_ = v784
	var v787 int32
	_ = v787
	var v788 int32
	_ = v788
	var v790 int32
	_ = v790
	var v792 int32
	_ = v792
	var v795 int32
	_ = v795
	var v796 int32
	_ = v796
	var v798 int32
	_ = v798
	var v800 int32
	_ = v800
	var v802 int32
	_ = v802
	var v806 int32
	_ = v806
	var v808 int32
	_ = v808
	var v812 int32
	_ = v812
	var v815 int32
	_ = v815
	var v816 int32
	_ = v816
	var v817 int32
	_ = v817
	var v824 int32
	_ = v824
	var v828 int32
	_ = v828
	var v833 int32
	_ = v833
	var v836 int32
	_ = v836
	var v840 int32
	_ = v840
	var v847 int32
	_ = v847
	var v848 int32
	_ = v848
	var v855 int32
	_ = v855
	var v877 int32
	_ = v877
	var v879 int32
	_ = v879
	var v880 int32
	_ = v880
	var v882 int32
	_ = v882
	var v883 int32
	_ = v883
	var v884 int32
	_ = v884
	var v886 int32
	_ = v886
	var v887 int32
	_ = v887
	var v888 int32
	_ = v888
	var v894 int32
	_ = v894
	var v896 int32
	_ = v896
	var v898 int32
	_ = v898
	var v901 int32
	_ = v901
	var v926 int32
	_ = v926
	var v929 int32
	_ = v929
	var v930 int32
	_ = v930
	var v931 int32
	_ = v931
	var v936 int32
	_ = v936
	var v940 int32
	_ = v940
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
	var v969 int32
	_ = v969
	var v974 int32
	_ = v974
	var v978 int32
	_ = v978
	var v981 int32
	_ = v981
	var v985 int32
	_ = v985
	var v990 int32
	_ = v990
	var v994 int32
	_ = v994
	var v998 int32
	_ = v998
	var v1003 int32
	_ = v1003
	v2 = int32(0)
	v21 = m.G0
	v23 = v21 - int32(112)
	m.G0 = v23
	*(*int32)(unsafe.Add(mBase, uint32(v23)+84)) = v2
	*(*int32)(unsafe.Add(mBase, uint32(v23)+80)) = v2
	v29 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+24)))
	if v29 == int32(1) {
		goto L7
	} else {
		goto L8
	}
L1:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v994 = m.ExcPending
	if v994 != 0 {
		goto L22
	} else {
		goto L276
	}
L2:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v978 = m.ExcPending
	if v978 != 0 {
		goto L22
	} else {
		goto L272
	}
L3:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v965 = m.ExcPending
	if v965 != 0 {
		goto L22
	} else {
		goto L269
	}
L4:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v949 = m.ExcPending
	if v949 != 0 {
		goto L22
	} else {
		goto L265
	}
L5:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v926 = m.ExcPending
	if v926 != 0 {
		goto L22
	} else {
		goto L259
	}
L6:
	;
	m.G0 = v23 + int32(112)
	return v901
L7:
	;
	v32 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(l0)+16)) = uint8(v32)
	v901 = v2
	goto L6
L8:
	;
	goto L9
L9:
	;
	v35 = l0 + int32(20)
	v36 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v37 = int32(0)
	if v36 == v37 {
		v48 = v37
		goto L12
	} else {
		goto L13
	}
L10:
	;
	v88 = *(*int32)(unsafe.Add(mBase, uint32(v35)))
	v89 = F_pg_detoast_datum_packed(m, v88)
	mBase = m.M
	v90 = m.ExcPending
	if v90 != 0 {
		goto L22
	} else {
		goto L26
	}
L11:
	;
	if v50 != 0 {
		goto L16
	} else {
		goto L17
	}
L12:
	;
	v50 = v48 & int32(1)
	goto L11
L13:
	;
	v40 = *(*int32)(unsafe.Add(mBase, uint32(v36)+24))
	if v40 == int32(0) {
		v48 = v37
		goto L12
	} else {
		goto L14
	}
L14:
	;
	v43 = *(*int32)(unsafe.Add(mBase, uint32(v40)))
	if v43 != int32(15) {
		v48 = v37
		goto L12
	} else {
		goto L15
	}
L15:
	;
	v46 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v40)+13)))
	v48 = v46
	goto L12
L16:
	;
	v51 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+32)))
	if v51 != 0 {
		goto L19
	} else {
		goto L20
	}
L17:
	;
	goto L18
L18:
	;
	v81 = int32(*(*int16)(unsafe.Add(mBase, uint32(l0)+18)))
	v83 = v2
	v87 = v81
	goto L10
L19:
	;
	v83 = v2
	v87 = int32(1)
	goto L10
L20:
	;
	goto L21
L21:
	;
	v53 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
	v54 = F_pg_detoast_datum(m, v53)
	mBase = m.M
	v57 = m.ExcPending
	if v57 != 0 {
		goto L22
	} else {
		goto L23
	}
L22:
	;
	return int32(0)
L23:
	;
	v58 = *(*int32)(unsafe.Add(mBase, uint32(v54)+12))
	F_get_typlenbyvalalign(m, v58, v23+int32(24), v23+int32(88), v23+int32(108))
	mBase = m.M
	v66 = m.ExcPending
	if v66 != 0 {
		goto L22
	} else {
		goto L24
	}
L24:
	;
	v67 = int32(*(*int16)(unsafe.Add(mBase, uint32(v23)+24)))
	v68 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v23)+88)))
	v69 = int32(*(*int8)(unsafe.Add(mBase, uint32(v23)+108)))
	F_deconstruct_array(m, v54, v67, v68, v69, v23+int32(84), v23+int32(80), v23+int32(52))
	mBase = m.M
	v77 = m.ExcPending
	if v77 != 0 {
		goto L22
	} else {
		goto L25
	}
L25:
	;
	v78 = *(*int32)(unsafe.Add(mBase, uint32(v23)+52))
	v83 = v58
	v87 = v78 + int32(1)
	goto L10
L26:
	;
	v91 = int32(1)
	v92 = v89 + v91
	v95 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v89))))
	v97 = v95 & v91
	if v97 != 0 {
		goto L27
	} else {
		goto L28
	}
L27:
	;
	v98 = v92
	goto L29
L28:
	;
	v98 = v89 + int32(4)
	goto L29
L29:
	;
	if v95 == int32(1) {
		goto L31
	} else {
		goto L32
	}
L30:
	;
	F_initStringInfo(m, v23+int32(88))
	mBase = m.M
	v130 = m.ExcPending
	if v130 != 0 {
		goto L22
	} else {
		goto L41
	}
L31:
	;
	v101 = int32(4)
	v103 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v92))))
	if v103&int32(254) == int32(2) {
		goto L34
	} else {
		goto L35
	}
L32:
	;
	goto L33
L33:
	;
	v116 = int32(1)
	if v97 != 0 {
		v126 = int32(base.Ui32(v95)>>(uint(v116)%32)) - v116
		goto L30
	} else {
		goto L40
	}
L34:
	;
	v112 = v101
	goto L36
L35:
	;
	v112 = base.B2i32(v103 == int32(18)) << (uint(v101) % 32)
	goto L36
L36:
	;
	if v103 == int32(1) {
		goto L37
	} else {
		goto L38
	}
L37:
	;
	v115 = v101
	goto L39
L38:
	;
	v115 = v112
	goto L39
L39:
	;
	v126 = v115
	goto L30
L40:
	;
	v120 = *(*int32)(unsafe.Add(mBase, uint32(v89)))
	v126 = int32(base.Ui32(v120)>>(uint(int32(2))%32)) - int32(4)
	goto L30
L41:
	;
	v131 = v98 + v126
	if base.Ui32(v98) < base.Ui32(v131) {
		goto L42
	} else {
		goto L43
	}
L42:
	;
	v136 = v98
	v140 = int32(1)
	v147 = v2
	v148 = v2
	goto L45
L43:
	;
	goto L44
L44:
	;
	v877 = *(*int32)(unsafe.Add(mBase, uint32(v23)+84))
	if v877 != 0 {
		goto L245
	} else {
		goto L246
	}
L45:
	;
	v154 = int32(*(*int8)(unsafe.Add(mBase, uint32(v136))))
	if v154 != int32(37) {
		goto L58
	} else {
		goto L59
	}
L46:
	;
	goto L44
L47:
	;
	v855 = v836 + int32(1)
	if base.Ui32(v855) < base.Ui32(v131) {
		v136 = v855
		v140 = v840
		v147 = v847
		v148 = v848
		goto L45
	} else {
		goto L244
	}
L48:
	;
	v529 = int32(*(*int8)(unsafe.Add(mBase, uint32(v510))))
	goto L145
L49:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v23)+108)) = v478
	v501 = F_text_format_parse_digits(m, v23+int32(108), v131, v23+int32(104))
	mBase = m.M
	v502 = m.ExcPending
	if v502 != 0 {
		goto L22
	} else {
		goto L137
	}
L50:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v459 = m.ExcPending
	if v459 != 0 {
		goto L22
	} else {
		goto L132
	}
L51:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v443 = m.ExcPending
	if v443 != 0 {
		goto L22
	} else {
		goto L128
	}
L52:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v427 = m.ExcPending
	if v427 != 0 {
		goto L22
	} else {
		goto L124
	}
L53:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v407 = m.ExcPending
	if v407 != 0 {
		goto L22
	} else {
		goto L119
	}
L54:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v387 = m.ExcPending
	if v387 != 0 {
		goto L22
	} else {
		goto L114
	}
L55:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v371 = m.ExcPending
	if v371 != 0 {
		goto L22
	} else {
		goto L110
	}
L56:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v351 = m.ExcPending
	if v351 != 0 {
		goto L22
	} else {
		goto L105
	}
L57:
	;
	v337 = *(*int32)(unsafe.Add(mBase, uint32(v23)+88))
	*(*uint8)(unsafe.Add(mBase, uint32(v337+v158))) = uint8(v154)
	v340 = *(*int32)(unsafe.Add(mBase, uint32(v23)+92))
	v342 = v340 + int32(1)
	*(*int32)(unsafe.Add(mBase, uint32(v23)+92)) = v342
	v344 = *(*int32)(unsafe.Add(mBase, uint32(v23)+88))
	v346 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v344+v342))) = uint8(v346)
	v836 = v136
	v840 = v140
	v847 = v147
	v848 = v148
	goto L47
L58:
	;
	v157 = *(*int32)(unsafe.Add(mBase, uint32(v23)+96))
	v158 = *(*int32)(unsafe.Add(mBase, uint32(v23)+92))
	if v158+int32(1) < v157 {
		goto L57
	} else {
		goto L61
	}
L59:
	;
	goto L60
L60:
	;
	v167 = v136 + int32(1)
	if base.Ui32(v131) <= base.Ui32(v167) {
		goto L56
	} else {
		goto L63
	}
L61:
	;
	F_appendStringInfoChar(m, v23+int32(88), v154)
	mBase = m.M
	v165 = m.ExcPending
	if v165 != 0 {
		goto L22
	} else {
		goto L62
	}
L62:
	;
	v836 = v136
	v840 = v140
	v847 = v147
	v848 = v148
	goto L47
L63:
	;
	v169 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v167))))
	if v169 == int32(37) {
		goto L64
	} else {
		goto L65
	}
L64:
	;
	v172 = *(*int32)(unsafe.Add(mBase, uint32(v23)+96))
	v173 = *(*int32)(unsafe.Add(mBase, uint32(v23)+92))
	if v172 <= v173+int32(1) {
		goto L67
	} else {
		goto L68
	}
L65:
	;
	goto L66
L66:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v23)+108)) = v167
	v200 = F_text_format_parse_digits(m, v23+int32(108), v131, v23+int32(104))
	mBase = m.M
	v201 = m.ExcPending
	if v201 != 0 {
		goto L22
	} else {
		goto L71
	}
L67:
	;
	F_appendStringInfoChar(m, v23+int32(88), int32(37))
	mBase = m.M
	v181 = m.ExcPending
	if v181 != 0 {
		goto L22
	} else {
		goto L70
	}
L68:
	;
	goto L69
L69:
	;
	v182 = *(*int32)(unsafe.Add(mBase, uint32(v23)+88))
	v184 = int32(37)
	*(*uint8)(unsafe.Add(mBase, uint32(v182+v173))) = uint8(v184)
	v186 = *(*int32)(unsafe.Add(mBase, uint32(v23)+92))
	v188 = v186 + int32(1)
	*(*int32)(unsafe.Add(mBase, uint32(v23)+92)) = v188
	v190 = *(*int32)(unsafe.Add(mBase, uint32(v23)+88))
	v192 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v190+v188))) = uint8(v192)
	v836 = v167
	v840 = v140
	v847 = v147
	v848 = v148
	goto L47
L70:
	;
	v836 = v167
	v840 = v140
	v847 = v147
	v848 = v148
	goto L47
L71:
	;
	v202 = *(*int32)(unsafe.Add(mBase, uint32(v23)+108))
	if v200 != 0 {
		goto L72
	} else {
		goto L73
	}
L72:
	;
	v203 = *(*int32)(unsafe.Add(mBase, uint32(v23)+104))
	v204 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v202))))
	if v204 != int32(36) {
		goto L75
	} else {
		goto L76
	}
L73:
	;
	v215 = v202
	v216 = int32(-1)
	goto L74
L74:
	;
	v217 = int32(0)
	v218 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v215))))
	switch v218 - int32(42) {
	case 0:
		v295 = v215
		v302 = v217
		goto L80
	default:
		v478 = v215
		v485 = v217
		goto L49
	case 3:
		goto L81
	}
L75:
	;
	v208 = int32(-1)
	v510 = v202
	v511 = v208
	v513 = v203
	v515 = v208
	v517 = int32(0)
	goto L48
L76:
	;
	goto L77
L77:
	;
	if v203 == int32(0) {
		goto L55
	} else {
		goto L78
	}
L78:
	;
	v213 = v202 + int32(1)
	if base.Ui32(v131) <= base.Ui32(v213) {
		goto L54
	} else {
		goto L79
	}
L79:
	;
	v215 = v213
	v216 = v203
	goto L74
L80:
	;
	v314 = v295 + int32(1)
	*(*int32)(unsafe.Add(mBase, uint32(v23)+108)) = v314
	if base.Ui32(v131) <= base.Ui32(v314) {
		goto L53
	} else {
		goto L97
	}
L81:
	;
	v222 = v215 + int32(1)
	if base.Ui32(v222) < base.Ui32(v131) {
		goto L82
	} else {
		goto L83
	}
L82:
	;
	v226 = v222
	goto L85
L83:
	;
	goto L84
L84:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v276 = m.ExcPending
	if v276 != 0 {
		goto L22
	} else {
		goto L92
	}
L85:
	;
	v244 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v226))))
	if v244 != int32(45) {
		goto L87
	} else {
		goto L88
	}
L86:
	;
	goto L84
L87:
	;
	v247 = int32(1)
	if v244 == int32(42) {
		v295 = v226
		v302 = v247
		goto L80
	} else {
		goto L90
	}
L88:
	;
	goto L89
L89:
	;
	v251 = v226 + int32(1)
	if base.Ui32(v251) < base.Ui32(v131) {
		v226 = v251
		goto L85
	} else {
		goto L91
	}
L90:
	;
	v478 = v226
	v485 = v247
	goto L49
L91:
	;
	goto L86
L92:
	;
	F_errcode(m, int32(50856066))
	mBase = m.M
	v279 = m.ExcPending
	if v279 != 0 {
		goto L22
	} else {
		goto L93
	}
L93:
	;
	F_errmsg(m, int32(212336), int32(0))
	mBase = m.M
	v283 = m.ExcPending
	if v283 != 0 {
		goto L22
	} else {
		goto L94
	}
L94:
	;
	F_errhint(m, int32(630142), int32(0))
	mBase = m.M
	v287 = m.ExcPending
	if v287 != 0 {
		goto L22
	} else {
		goto L95
	}
L95:
	;
	F_errfinish(m, int32(478374), int32(6261), int32(104825))
	mBase = m.M
	v292 = m.ExcPending
	if v292 != 0 {
		goto L22
	} else {
		goto L96
	}
L96:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L97:
	;
	v317 = int32(0)
	v322 = F_text_format_parse_digits(m, v23+int32(108), v131, v23+int32(104))
	mBase = m.M
	v323 = m.ExcPending
	if v323 != 0 {
		goto L22
	} else {
		goto L98
	}
L98:
	;
	v324 = *(*int32)(unsafe.Add(mBase, uint32(v23)+108))
	if v322 == int32(0) {
		goto L99
	} else {
		goto L100
	}
L99:
	;
	v510 = v324
	v511 = int32(0)
	v513 = v317
	v515 = v216
	v517 = v302
	goto L48
L100:
	;
	goto L101
L101:
	;
	v328 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v324))))
	if v328 != int32(36) {
		goto L52
	} else {
		goto L102
	}
L102:
	;
	v331 = *(*int32)(unsafe.Add(mBase, uint32(v23)+104))
	if v331 == int32(0) {
		goto L51
	} else {
		goto L103
	}
L103:
	;
	v335 = v324 + int32(1)
	if base.Ui32(v131) <= base.Ui32(v335) {
		goto L50
	} else {
		goto L104
	}
L104:
	;
	v510 = v335
	v511 = v331
	v513 = v317
	v515 = v216
	v517 = v302
	goto L48
L105:
	;
	F_errcode(m, int32(50856066))
	mBase = m.M
	v354 = m.ExcPending
	if v354 != 0 {
		goto L22
	} else {
		goto L106
	}
L106:
	;
	F_errmsg(m, int32(212336), int32(0))
	mBase = m.M
	v358 = m.ExcPending
	if v358 != 0 {
		goto L22
	} else {
		goto L107
	}
L107:
	;
	F_errhint(m, int32(630142), int32(0))
	mBase = m.M
	v362 = m.ExcPending
	if v362 != 0 {
		goto L22
	} else {
		goto L108
	}
L108:
	;
	F_errfinish(m, int32(478374), int32(5999), int32(104765))
	mBase = m.M
	v367 = m.ExcPending
	if v367 != 0 {
		goto L22
	} else {
		goto L109
	}
L109:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L110:
	;
	F_errcode(m, int32(50856066))
	mBase = m.M
	v374 = m.ExcPending
	if v374 != 0 {
		goto L22
	} else {
		goto L111
	}
L111:
	;
	F_errmsg(m, int32(531099), int32(0))
	mBase = m.M
	v378 = m.ExcPending
	if v378 != 0 {
		goto L22
	} else {
		goto L112
	}
L112:
	;
	F_errfinish(m, int32(478374), int32(6253), int32(104825))
	mBase = m.M
	v383 = m.ExcPending
	if v383 != 0 {
		goto L22
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
	F_errcode(m, int32(50856066))
	mBase = m.M
	v390 = m.ExcPending
	if v390 != 0 {
		goto L22
	} else {
		goto L115
	}
L115:
	;
	F_errmsg(m, int32(212336), int32(0))
	mBase = m.M
	v394 = m.ExcPending
	if v394 != 0 {
		goto L22
	} else {
		goto L116
	}
L116:
	;
	F_errhint(m, int32(630142), int32(0))
	mBase = m.M
	v398 = m.ExcPending
	if v398 != 0 {
		goto L22
	} else {
		goto L117
	}
L117:
	;
	F_errfinish(m, int32(478374), int32(6254), int32(104825))
	mBase = m.M
	v403 = m.ExcPending
	if v403 != 0 {
		goto L22
	} else {
		goto L118
	}
L118:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L119:
	;
	F_errcode(m, int32(50856066))
	mBase = m.M
	v410 = m.ExcPending
	if v410 != 0 {
		goto L22
	} else {
		goto L120
	}
L120:
	;
	F_errmsg(m, int32(212336), int32(0))
	mBase = m.M
	v414 = m.ExcPending
	if v414 != 0 {
		goto L22
	} else {
		goto L121
	}
L121:
	;
	F_errhint(m, int32(630142), int32(0))
	mBase = m.M
	v418 = m.ExcPending
	if v418 != 0 {
		goto L22
	} else {
		goto L122
	}
L122:
	;
	F_errfinish(m, int32(478374), int32(6267), int32(104825))
	mBase = m.M
	v423 = m.ExcPending
	if v423 != 0 {
		goto L22
	} else {
		goto L123
	}
L123:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L124:
	;
	F_errcode(m, int32(50856066))
	mBase = m.M
	v430 = m.ExcPending
	if v430 != 0 {
		goto L22
	} else {
		goto L125
	}
L125:
	;
	F_errmsg(m, int32(690360), int32(0))
	mBase = m.M
	v434 = m.ExcPending
	if v434 != 0 {
		goto L22
	} else {
		goto L126
	}
L126:
	;
	F_errfinish(m, int32(478374), int32(6274), int32(104825))
	mBase = m.M
	v439 = m.ExcPending
	if v439 != 0 {
		goto L22
	} else {
		goto L127
	}
L127:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L128:
	;
	F_errcode(m, int32(50856066))
	mBase = m.M
	v446 = m.ExcPending
	if v446 != 0 {
		goto L22
	} else {
		goto L129
	}
L129:
	;
	F_errmsg(m, int32(531099), int32(0))
	mBase = m.M
	v450 = m.ExcPending
	if v450 != 0 {
		goto L22
	} else {
		goto L130
	}
L130:
	;
	F_errfinish(m, int32(478374), int32(6281), int32(104825))
	mBase = m.M
	v455 = m.ExcPending
	if v455 != 0 {
		goto L22
	} else {
		goto L131
	}
L131:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L132:
	;
	F_errcode(m, int32(50856066))
	mBase = m.M
	v462 = m.ExcPending
	if v462 != 0 {
		goto L22
	} else {
		goto L133
	}
L133:
	;
	F_errmsg(m, int32(212336), int32(0))
	mBase = m.M
	v466 = m.ExcPending
	if v466 != 0 {
		goto L22
	} else {
		goto L134
	}
L134:
	;
	F_errhint(m, int32(630142), int32(0))
	mBase = m.M
	v470 = m.ExcPending
	if v470 != 0 {
		goto L22
	} else {
		goto L135
	}
L135:
	;
	F_errfinish(m, int32(478374), int32(6282), int32(104825))
	mBase = m.M
	v475 = m.ExcPending
	if v475 != 0 {
		goto L22
	} else {
		goto L136
	}
L136:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L137:
	;
	v503 = *(*int32)(unsafe.Add(mBase, uint32(v23)+104))
	if v501 != 0 {
		goto L138
	} else {
		goto L139
	}
L138:
	;
	v505 = v503
	goto L140
L139:
	;
	v505 = int32(0)
	goto L140
L140:
	;
	v507 = *(*int32)(unsafe.Add(mBase, uint32(v23)+108))
	v510 = v507
	v511 = int32(-1)
	v513 = v505
	v515 = v216
	v517 = v485
	goto L48
L141:
	;
	if v633 == int32(0) {
		goto L5
	} else {
		goto L167
	}
L142:
	;
	v633 = int32(0)
	goto L141
L143:
	;
	v611 = v604
	v613 = v606
	goto L161
L144:
	;
	if base.B2i32(v551 != v552) == int32(0) {
		goto L142
	} else {
		goto L152
	}
L145:
	;
	goto L146
L146:
	;
	v543 = int32(511289)
	v545 = int32(4)
	goto L147
L147:
	;
	v548 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v543))))
	if v548 == v529&int32(255) {
		v604 = v543
		v606 = v545
		goto L143
	} else {
		goto L149
	}
L148:
	;
	goto L144
L149:
	;
	v550 = int32(1)
	v551 = v545 - v550
	v552 = int32(0)
	v555 = v543 + v550
	if v555&int32(3) == v552 {
		goto L144
	} else {
		goto L150
	}
L150:
	;
	if v551 != 0 {
		v543 = v555
		v545 = v551
		goto L147
	} else {
		goto L151
	}
L151:
	;
	goto L148
L152:
	;
	v567 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v555))))
	if v567 == v529&int32(255) {
		v597 = v555
		v599 = v551
		goto L153
	} else {
		goto L154
	}
L153:
	;
	if v599 == int32(0) {
		goto L142
	} else {
		goto L160
	}
L154:
	;
	if base.Ui32(v551) < base.Ui32(int32(4)) {
		v597 = v555
		v599 = v551
		goto L153
	} else {
		goto L155
	}
L155:
	;
	v577 = v555
	v579 = v551
	goto L156
L156:
	;
	v583 = *(*int32)(unsafe.Add(mBase, uint32(v577)))
	v584 = v583 ^ v529&int32(255)*int32(16843009)
	v587 = int32(-2139062144)
	if (int32(16843008)-v584|v584)&v587 != v587 {
		v604 = v577
		v606 = v579
		goto L143
	} else {
		goto L158
	}
L157:
	;
	v597 = v592
	v599 = v594
	goto L153
L158:
	;
	v591 = int32(4)
	v592 = v577 + v591
	v594 = v579 - v591
	if base.Ui32(int32(3)) < base.Ui32(v594) {
		v577 = v592
		v579 = v594
		goto L156
	} else {
		goto L159
	}
L159:
	;
	goto L157
L160:
	;
	v604 = v597
	v606 = v599
	goto L143
L161:
	;
	v616 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v611))))
	if v529&int32(255) == v616 {
		goto L163
	} else {
		goto L164
	}
L162:
	;
	goto L142
L163:
	;
	v633 = v611
	goto L141
L164:
	;
	goto L165
L165:
	;
	v618 = int32(1)
	v621 = v613 - v618
	if v621 != 0 {
		v611 = v611 + v618
		v613 = v621
		goto L161
	} else {
		goto L166
	}
L166:
	;
	goto L162
L167:
	;
	if v511 < int32(0) {
		v696 = v513
		v697 = v140
		v700 = v147
		goto L168
	} else {
		goto L169
	}
L168:
	;
	if int32(0) < v515 {
		goto L193
	} else {
		goto L194
	}
L169:
	;
	if v511 != 0 {
		goto L170
	} else {
		goto L171
	}
L170:
	;
	v638 = v511
	goto L172
L171:
	;
	v638 = v140
	goto L172
L172:
	;
	if v87 <= v638 {
		goto L4
	} else {
		goto L173
	}
L173:
	;
	if v50 == int32(0) {
		goto L175
	} else {
		goto L176
	}
L174:
	;
	if v663 == int32(0) {
		goto L3
	} else {
		goto L179
	}
L175:
	;
	v644 = v35 + v638<<(uint(int32(3))%32)
	v645 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v644)+4)))
	v646 = *(*int32)(unsafe.Add(mBase, uint32(v644)))
	v647 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v648 = F_get_fn_expr_argtype(m, v647, v638)
	mBase = m.M
	v649 = m.ExcPending
	if v649 != 0 {
		goto L22
	} else {
		goto L178
	}
L176:
	;
	goto L177
L177:
	;
	v651 = v638 - int32(1)
	v652 = *(*int32)(unsafe.Add(mBase, uint32(v23)+80))
	v654 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v651+v652))))
	v655 = *(*int32)(unsafe.Add(mBase, uint32(v23)+84))
	v659 = *(*int32)(unsafe.Add(mBase, uint32(v655+v651<<(uint(int32(2))%32))))
	v660 = v659
	v662 = v654
	v663 = v83
	goto L174
L178:
	;
	v660 = v646
	v662 = v645
	v663 = v648
	goto L174
L179:
	;
	v666 = int32(1)
	v667 = v638 + v666
	if v662&v666 != 0 {
		goto L180
	} else {
		goto L181
	}
L180:
	;
	v696 = int32(0)
	v697 = v667
	v700 = v147
	goto L168
L181:
	;
	goto L182
L182:
	;
	switch v663 - int32(21) {
	case 0:
		goto L184
	default:
		goto L183
	case 2:
		v696 = v660
		v697 = v667
		v700 = v147
		goto L168
	}
L183:
	;
	if v663 != v147 {
		goto L185
	} else {
		goto L186
	}
L184:
	;
	v696 = base.I32_extend16_s(v660)
	v697 = v667
	v700 = v147
	goto L168
L185:
	;
	F_getTypeOutputInfo(m, v663, v23+int32(108), v23+int32(104))
	mBase = m.M
	v680 = m.ExcPending
	if v680 != 0 {
		goto L22
	} else {
		goto L188
	}
L186:
	;
	v686 = v147
	goto L187
L187:
	;
	v689 = F_OutputFunctionCall(m, v23+int32(24), v660)
	mBase = m.M
	v690 = m.ExcPending
	if v690 != 0 {
		goto L22
	} else {
		goto L190
	}
L188:
	;
	v681 = *(*int32)(unsafe.Add(mBase, uint32(v23)+108))
	F_fmgr_info(m, v681, v23+int32(24))
	mBase = m.M
	v685 = m.ExcPending
	if v685 != 0 {
		goto L22
	} else {
		goto L189
	}
L189:
	;
	v686 = v663
	goto L187
L190:
	;
	v691 = F_pg_strtoint32(m, v689)
	mBase = m.M
	v692 = m.ExcPending
	if v692 != 0 {
		goto L22
	} else {
		goto L191
	}
L191:
	;
	F_pfree(m, v689)
	mBase = m.M
	v694 = m.ExcPending
	if v694 != 0 {
		goto L22
	} else {
		goto L192
	}
L192:
	;
	v696 = v691
	v697 = v667
	v700 = v686
	goto L168
L193:
	;
	v703 = v515
	goto L195
L194:
	;
	v703 = v697
	goto L195
L195:
	;
	if v87 <= v703 {
		goto L2
	} else {
		goto L196
	}
L196:
	;
	if v50 == int32(0) {
		goto L198
	} else {
		goto L199
	}
L197:
	;
	if v728 == int32(0) {
		goto L1
	} else {
		goto L202
	}
L198:
	;
	v709 = v35 + v703<<(uint(int32(3))%32)
	v710 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v709)+4)))
	v711 = *(*int32)(unsafe.Add(mBase, uint32(v709)))
	v712 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v713 = F_get_fn_expr_argtype(m, v712, v703)
	mBase = m.M
	v714 = m.ExcPending
	if v714 != 0 {
		goto L22
	} else {
		goto L201
	}
L199:
	;
	goto L200
L200:
	;
	v716 = v703 - int32(1)
	v717 = *(*int32)(unsafe.Add(mBase, uint32(v23)+80))
	v719 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v716+v717))))
	v720 = *(*int32)(unsafe.Add(mBase, uint32(v23)+84))
	v724 = *(*int32)(unsafe.Add(mBase, uint32(v720+v716<<(uint(int32(2))%32))))
	v726 = v719
	v727 = v724
	v728 = v83
	goto L197
L201:
	;
	v726 = v710
	v727 = v711
	v728 = v713
	goto L197
L202:
	;
	if v728 != v148 {
		goto L203
	} else {
		goto L204
	}
L203:
	;
	F_getTypeOutputInfo(m, v728, v23+int32(108), v23+int32(104))
	mBase = m.M
	v737 = m.ExcPending
	if v737 != 0 {
		goto L22
	} else {
		goto L206
	}
L204:
	;
	v743 = v148
	goto L205
L205:
	;
	v744 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v510))))
	v746 = v744 - int32(73)
	switch v746 {
	case 0, 3:
		goto L209
	case 1, 2:
		goto L208
	default:
		goto L210
	}
L206:
	;
	v738 = *(*int32)(unsafe.Add(mBase, uint32(v23)+108))
	F_fmgr_info(m, v738, v23+int32(52))
	mBase = m.M
	v742 = m.ExcPending
	if v742 != 0 {
		goto L22
	} else {
		goto L207
	}
L207:
	;
	v743 = v728
	goto L205
L208:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v812 = m.ExcPending
	if v812 != 0 {
		goto L22
	} else {
		goto L238
	}
L209:
	;
	v749 = int32(1)
	v750 = v703 + v749
	if v726&v749 != 0 {
		goto L212
	} else {
		goto L213
	}
L210:
	;
	if v744 != int32(115) {
		goto L208
	} else {
		goto L211
	}
L211:
	;
	goto L209
L212:
	;
	switch v746 {
	case 0:
		goto L215
	case 1, 2:
		v836 = v510
		v840 = v750
		v847 = v700
		v848 = v743
		goto L47
	case 3:
		goto L216
	default:
		goto L217
	}
L213:
	;
	goto L214
L214:
	;
	v783 = F_OutputFunctionCall(m, v23+int32(52), v727)
	mBase = m.M
	v784 = m.ExcPending
	if v784 != 0 {
		goto L22
	} else {
		goto L225
	}
L215:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v768 = m.ExcPending
	if v768 != 0 {
		goto L22
	} else {
		goto L221
	}
L216:
	;
	F_text_format_append_string(m, v23+int32(88), int32(510768), v517, v696)
	mBase = m.M
	v764 = m.ExcPending
	if v764 != 0 {
		goto L22
	} else {
		goto L220
	}
L217:
	;
	if v744 != int32(115) {
		v836 = v510
		v840 = v750
		v847 = v700
		v848 = v743
		goto L47
	} else {
		goto L218
	}
L218:
	;
	F_text_format_append_string(m, v23+int32(88), int32(717063), v517, v696)
	mBase = m.M
	v759 = m.ExcPending
	if v759 != 0 {
		goto L22
	} else {
		goto L219
	}
L219:
	;
	v836 = v510
	v840 = v750
	v847 = v700
	v848 = v743
	goto L47
L220:
	;
	v836 = v510
	v840 = v750
	v847 = v700
	v848 = v743
	goto L47
L221:
	;
	F_errcode(m, int32(67108994))
	mBase = m.M
	v771 = m.ExcPending
	if v771 != 0 {
		goto L22
	} else {
		goto L222
	}
L222:
	;
	F_errmsg(m, int32(212136), int32(0))
	mBase = m.M
	v775 = m.ExcPending
	if v775 != 0 {
		goto L22
	} else {
		goto L223
	}
L223:
	;
	F_errfinish(m, int32(478374), int32(6319), int32(259238))
	mBase = m.M
	v780 = m.ExcPending
	if v780 != 0 {
		goto L22
	} else {
		goto L224
	}
L224:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L225:
	;
	switch v746 {
	case 0:
		goto L228
	default:
		goto L226
	case 3:
		goto L227
	}
L226:
	;
	F_text_format_append_string(m, v23+int32(88), v783, v517, v696)
	mBase = m.M
	v806 = m.ExcPending
	if v806 != 0 {
		goto L22
	} else {
		goto L236
	}
L227:
	;
	v795 = F_quote_literal_cstr(m, v783)
	mBase = m.M
	v796 = m.ExcPending
	if v796 != 0 {
		goto L22
	} else {
		goto L232
	}
L228:
	;
	v787 = F_quote_identifier(m, v783)
	mBase = m.M
	v788 = m.ExcPending
	if v788 != 0 {
		goto L22
	} else {
		goto L229
	}
L229:
	;
	F_text_format_append_string(m, v23+int32(88), v787, v517, v696)
	mBase = m.M
	v790 = m.ExcPending
	if v790 != 0 {
		goto L22
	} else {
		goto L230
	}
L230:
	;
	F_pfree(m, v783)
	mBase = m.M
	v792 = m.ExcPending
	if v792 != 0 {
		goto L22
	} else {
		goto L231
	}
L231:
	;
	v836 = v510
	v840 = v750
	v847 = v700
	v848 = v743
	goto L47
L232:
	;
	F_text_format_append_string(m, v23+int32(88), v795, v517, v696)
	mBase = m.M
	v798 = m.ExcPending
	if v798 != 0 {
		goto L22
	} else {
		goto L233
	}
L233:
	;
	F_pfree(m, v795)
	mBase = m.M
	v800 = m.ExcPending
	if v800 != 0 {
		goto L22
	} else {
		goto L234
	}
L234:
	;
	F_pfree(m, v783)
	mBase = m.M
	v802 = m.ExcPending
	if v802 != 0 {
		goto L22
	} else {
		goto L235
	}
L235:
	;
	v836 = v510
	v840 = v750
	v847 = v700
	v848 = v743
	goto L47
L236:
	;
	F_pfree(m, v783)
	mBase = m.M
	v808 = m.ExcPending
	if v808 != 0 {
		goto L22
	} else {
		goto L237
	}
L237:
	;
	v836 = v510
	v840 = v750
	v847 = v700
	v848 = v743
	goto L47
L238:
	;
	F_errcode(m, int32(50856066))
	mBase = m.M
	v815 = m.ExcPending
	if v815 != 0 {
		goto L22
	} else {
		goto L239
	}
L239:
	;
	v816 = F_pg_mblen_range(m, v510, v131)
	mBase = m.M
	v817 = m.ExcPending
	if v817 != 0 {
		goto L22
	} else {
		goto L240
	}
L240:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v23)+20)) = v510
	*(*int32)(unsafe.Add(mBase, uint32(v23)+16)) = v816
	F_errmsg(m, int32(650127), v23+int32(16))
	mBase = m.M
	v824 = m.ExcPending
	if v824 != 0 {
		goto L22
	} else {
		goto L241
	}
L241:
	;
	F_errhint(m, int32(630142), int32(0))
	mBase = m.M
	v828 = m.ExcPending
	if v828 != 0 {
		goto L22
	} else {
		goto L242
	}
L242:
	;
	F_errfinish(m, int32(478374), int32(6147), int32(104765))
	mBase = m.M
	v833 = m.ExcPending
	if v833 != 0 {
		goto L22
	} else {
		goto L243
	}
L243:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L244:
	;
	goto L46
L245:
	;
	F_pfree(m, v877)
	mBase = m.M
	v879 = m.ExcPending
	if v879 != 0 {
		goto L22
	} else {
		goto L248
	}
L246:
	;
	goto L247
L247:
	;
	v880 = *(*int32)(unsafe.Add(mBase, uint32(v23)+80))
	if v880 != 0 {
		goto L249
	} else {
		goto L250
	}
L248:
	;
	goto L247
L249:
	;
	F_pfree(m, v880)
	mBase = m.M
	v882 = m.ExcPending
	if v882 != 0 {
		goto L22
	} else {
		goto L252
	}
L250:
	;
	goto L251
L251:
	;
	v883 = *(*int32)(unsafe.Add(mBase, uint32(v23)+88))
	v884 = *(*int32)(unsafe.Add(mBase, uint32(v23)+92))
	v886 = v884 + int32(4)
	v887 = F_palloc(m, v886)
	mBase = m.M
	v888 = m.ExcPending
	if v888 != 0 {
		goto L22
	} else {
		goto L253
	}
L252:
	;
	goto L251
L253:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v887))) = v886 << (uint(int32(2)) % 32)
	if v884 != 0 {
		goto L255
	} else {
		goto L256
	}
L254:
	;
	v896 = *(*int32)(unsafe.Add(mBase, uint32(v23)+88))
	F_pfree(m, v896)
	mBase = m.M
	v898 = m.ExcPending
	if v898 != 0 {
		goto L22
	} else {
		goto L258
	}
L255:
	;
	v894 = F__emscripten_memcpy_bulkmem(m, v887+int32(4), v883, v884)
	mBase = m.M
	goto L257
L256:
	;
	goto L257
L257:
	;
	goto L254
L258:
	;
	v901 = v887
	goto L6
L259:
	;
	F_errcode(m, int32(50856066))
	mBase = m.M
	v929 = m.ExcPending
	if v929 != 0 {
		goto L22
	} else {
		goto L260
	}
L260:
	;
	v930 = F_pg_mblen_range(m, v510, v131)
	mBase = m.M
	v931 = m.ExcPending
	if v931 != 0 {
		goto L22
	} else {
		goto L261
	}
L261:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v23)+4)) = v510
	*(*int32)(unsafe.Add(mBase, uint32(v23))) = v930
	F_errmsg(m, int32(650127), v23)
	mBase = m.M
	v936 = m.ExcPending
	if v936 != 0 {
		goto L22
	} else {
		goto L262
	}
L262:
	;
	F_errhint(m, int32(630142), int32(0))
	mBase = m.M
	v940 = m.ExcPending
	if v940 != 0 {
		goto L22
	} else {
		goto L263
	}
L263:
	;
	F_errfinish(m, int32(478374), int32(6026), int32(104765))
	mBase = m.M
	v945 = m.ExcPending
	if v945 != 0 {
		goto L22
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
	F_errcode(m, int32(50856066))
	mBase = m.M
	v952 = m.ExcPending
	if v952 != 0 {
		goto L22
	} else {
		goto L266
	}
L266:
	;
	F_errmsg(m, int32(643328), int32(0))
	mBase = m.M
	v956 = m.ExcPending
	if v956 != 0 {
		goto L22
	} else {
		goto L267
	}
L267:
	;
	F_errfinish(m, int32(478374), int32(6037), int32(104765))
	mBase = m.M
	v961 = m.ExcPending
	if v961 != 0 {
		goto L22
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
	F_errmsg_internal(m, int32(62342), int32(0))
	mBase = m.M
	v969 = m.ExcPending
	if v969 != 0 {
		goto L22
	} else {
		goto L270
	}
L270:
	;
	F_errfinish(m, int32(478374), int32(6053), int32(104765))
	mBase = m.M
	v974 = m.ExcPending
	if v974 != 0 {
		goto L22
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
	F_errcode(m, int32(50856066))
	mBase = m.M
	v981 = m.ExcPending
	if v981 != 0 {
		goto L22
	} else {
		goto L273
	}
L273:
	;
	F_errmsg(m, int32(643328), int32(0))
	mBase = m.M
	v985 = m.ExcPending
	if v985 != 0 {
		goto L22
	} else {
		goto L274
	}
L274:
	;
	F_errfinish(m, int32(478374), int32(6094), int32(104765))
	mBase = m.M
	v990 = m.ExcPending
	if v990 != 0 {
		goto L22
	} else {
		goto L275
	}
L275:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L276:
	;
	F_errmsg_internal(m, int32(62342), int32(0))
	mBase = m.M
	v998 = m.ExcPending
	if v998 != 0 {
		goto L22
	} else {
		goto L277
	}
L277:
	;
	F_errfinish(m, int32(478374), int32(6110), int32(104765))
	mBase = m.M
	v1003 = m.ExcPending
	if v1003 != 0 {
		goto L22
	} else {
		goto L278
	}
L278:
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
	var v28 int32
	_ = v28
	var v30 int32
	_ = v30
	var v39 int32
	_ = v39
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
	var v63 int32
	_ = v63
	var v65 int32
	_ = v65
	var v74 int32
	_ = v74
	var v77 int32
	_ = v77
	var v78 int32
	_ = v78
	var v84 int32
	_ = v84
	var v90 int32
	_ = v90
	var v91 int32
	_ = v91
	var v92 int32
	_ = v92
	var v95 int32
	_ = v95
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
				v28 = int32(4)
				v30 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v14))))
				if v30&int32(254) == int32(2) {
					v39 = v28
				} else {
					v39 = base.B2i32(v30 == int32(18)) << (uint(v28) % 32)
				}
				if v30 == int32(1) {
					v42 = v28
				} else {
					v42 = v39
				}
				v53 = v42
			} else {
				v43 = int32(1)
				if v24 != 0 {
					v53 = int32(base.Ui32(v22)>>(uint(v43)%32)) - v43
				} else {
					v47 = *(*int32)(unsafe.Add(mBase, uint32(v9)))
					v53 = int32(base.Ui32(v47)>>(uint(int32(2))%32)) - int32(4)
				}
			}
			v54 = int32(1)
			v55 = v16 + v54
			if v18&v54 != 0 {
				v60 = v55
			} else {
				v60 = v16 + int32(4)
			}
			if v18 == int32(1) {
				v63 = int32(4)
				v65 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v55))))
				if v65&int32(254) == int32(2) {
					v74 = v63
				} else {
					v74 = base.B2i32(v65 == int32(18)) << (uint(v63) % 32)
				}
				if v65 == int32(1) {
					v77 = v63
				} else {
					v77 = v74
				}
				v90 = v77
			} else {
				v78 = int32(1)
				if v18&v78 != 0 {
					v90 = int32(base.Ui32(v18)>>(uint(v78)%32)) - v78
				} else {
					v84 = *(*int32)(unsafe.Add(mBase, uint32(v16)))
					v90 = int32(base.Ui32(v84)>>(uint(int32(2))%32)) - int32(4)
				}
			}
			v91 = F_varstr_cmp(m, v25, v53, v60, v90, v19)
			mBase = m.M
			v92 = m.ExcPending
			if v92 != 0 {
				return int32(0)
			} else {
				if int32(0) < v91 {
					v95 = v9
				} else {
					v95 = v16
				}
				return v95
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
	var v29 int32
	_ = v29
	var v31 int32
	_ = v31
	var v40 int32
	_ = v40
	var v43 int32
	_ = v43
	var v44 int32
	_ = v44
	var v48 int32
	_ = v48
	var v54 int32
	_ = v54
	var v55 int32
	_ = v55
	var v56 int32
	_ = v56
	var v61 int32
	_ = v61
	var v64 int32
	_ = v64
	var v66 int32
	_ = v66
	var v75 int32
	_ = v75
	var v78 int32
	_ = v78
	var v79 int32
	_ = v79
	var v85 int32
	_ = v85
	var v91 int32
	_ = v91
	var v92 int32
	_ = v92
	var v93 int32
	_ = v93
	var v94 int32
	_ = v94
	var v97 int32
	_ = v97
	var v98 int32
	_ = v98
	var v101 int32
	_ = v101
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
				v29 = int32(4)
				v31 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v15))))
				if v31&int32(254) == int32(2) {
					v40 = v29
				} else {
					v40 = base.B2i32(v31 == int32(18)) << (uint(v29) % 32)
				}
				if v31 == int32(1) {
					v43 = v29
				} else {
					v43 = v40
				}
				v54 = v43
			} else {
				v44 = int32(1)
				if v25 != 0 {
					v54 = int32(base.Ui32(v23)>>(uint(v44)%32)) - v44
				} else {
					v48 = *(*int32)(unsafe.Add(mBase, uint32(v10)))
					v54 = int32(base.Ui32(v48)>>(uint(int32(2))%32)) - int32(4)
				}
			}
			v55 = int32(1)
			v56 = v17 + v55
			if v19&v55 != 0 {
				v61 = v56
			} else {
				v61 = v17 + int32(4)
			}
			if v19 == int32(1) {
				v64 = int32(4)
				v66 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v56))))
				if v66&int32(254) == int32(2) {
					v75 = v64
				} else {
					v75 = base.B2i32(v66 == int32(18)) << (uint(v64) % 32)
				}
				if v66 == int32(1) {
					v78 = v64
				} else {
					v78 = v75
				}
				v91 = v78
			} else {
				v79 = int32(1)
				if v19&v79 != 0 {
					v91 = int32(base.Ui32(v19)>>(uint(v79)%32)) - v79
				} else {
					v85 = *(*int32)(unsafe.Add(mBase, uint32(v17)))
					v91 = int32(base.Ui32(v85)>>(uint(int32(2))%32)) - int32(4)
				}
			}
			v92 = F_varstr_cmp(m, v26, v54, v61, v91, v20)
			mBase = m.M
			v93 = m.ExcPending
			if v93 != 0 {
				return int32(0)
			} else {
				v94 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
				if v94 != v10 {
					F_pfree(m, v10)
					mBase = m.M
					v97 = m.ExcPending
					if v97 != 0 {
						return int32(0)
					} else {
						v98 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
						if v98 != v17 {
							F_pfree(m, v17)
							mBase = m.M
							v101 = m.ExcPending
							if v101 != 0 {
								return int32(0)
							} else {
								return base.B2i32(v92 <= int32(0))
							}
						} else {
							return base.B2i32(v92 <= int32(0))
						}
					}
				} else {
					v98 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
					if v98 != v17 {
						F_pfree(m, v17)
						mBase = m.M
						v101 = m.ExcPending
						if v101 != 0 {
							return int32(0)
						} else {
							return base.B2i32(v92 <= int32(0))
						}
					} else {
						return base.B2i32(v92 <= int32(0))
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
	var v25 int32
	_ = v25
	var v27 int32
	_ = v27
	var v36 int32
	_ = v36
	var v39 int32
	_ = v39
	var v40 int32
	_ = v40
	var v44 int32
	_ = v44
	var v50 int32
	_ = v50
	var v51 int32
	_ = v51
	var v52 int32
	_ = v52
	var v54 int32
	_ = v54
	var v55 int32
	_ = v55
	var v57 int32
	_ = v57
	var v58 int32
	_ = v58
	var v59 int32
	_ = v59
	var v65 int32
	_ = v65
	var v70 int32
	_ = v70
	var v71 int32
	_ = v71
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
				v25 = int32(4)
				v27 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v16))))
				if v27&int32(254) == int32(2) {
					v36 = v25
				} else {
					v36 = base.B2i32(v27 == int32(18)) << (uint(v25) % 32)
				}
				if v27 == int32(1) {
					v39 = v25
				} else {
					v39 = v36
				}
				v50 = v39
			} else {
				v40 = int32(1)
				if v21 != 0 {
					v50 = int32(base.Ui32(v19)>>(uint(v40)%32)) - v40
				} else {
					v44 = *(*int32)(unsafe.Add(mBase, uint32(v11)))
					v50 = int32(base.Ui32(v44)>>(uint(int32(2))%32)) - int32(4)
				}
			}
			v51 = F_pg_mbstrlen_with_len(m, v22, v50)
			mBase = m.M
			v52 = m.ExcPending
			if v52 != 0 {
				return int32(0)
			} else {
				v54 = F_pg_mbcharcliplen(m, v22, v50, v51+v8)
				mBase = m.M
				v55 = m.ExcPending
				if v55 != 0 {
					return int32(0)
				} else {
					v57 = v54 + int32(4)
					v58 = F_palloc(m, v57)
					mBase = m.M
					v59 = m.ExcPending
					if v59 != 0 {
						return int32(0)
					} else {
						*(*int32)(unsafe.Add(mBase, uint32(v58))) = v57 << (uint(int32(2)) % 32)
						if v54 != 0 {
							v65 = F__emscripten_memcpy_bulkmem(m, v58+int32(4), v22, v54)
							mBase = m.M
						} else {
						}
						return v58
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
			return v70
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
	var v12 int32
	_ = v12
	var v14 int32
	_ = v14
	var v23 int32
	_ = v23
	var v26 int32
	_ = v26
	var v29 int32
	_ = v29
	var v33 int32
	_ = v33
	var v38 int32
	_ = v38
	var v41 int32
	_ = v41
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
	var v53 int32
	_ = v53
	var v54 int32
	_ = v54
	var v56 int32
	_ = v56
	var v59 int32
	_ = v59
	var v61 int32
	_ = v61
	var v62 int32
	_ = v62
	v4 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v5 = F_pg_detoast_datum_packed(m, v4)
	mBase = m.M
	v8 = m.ExcPending
	if v8 != 0 {
		return int32(0)
	} else {
		v9 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v5))))
		if v9 == int32(1) {
			v12 = int32(4)
			v14 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v5)+1)))
			if v14&int32(254) == int32(2) {
				v23 = v12
			} else {
				v23 = base.B2i32(v14 == int32(18)) << (uint(v12) % 32)
			}
			if v14 == int32(1) {
				v26 = v12
			} else {
				v26 = v23
			}
			v50 = v26
			v52 = F_palloc0(m, int32(64))
			mBase = m.M
			v53 = m.ExcPending
			if v53 != 0 {
				return int32(0)
			} else {
				v54 = int32(1)
				v56 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v5))))
				if v56&v54 != 0 {
					v59 = v54
				} else {
					v59 = int32(4)
				}
				if v50 != 0 {
					v61 = F__emscripten_memcpy_bulkmem(m, v52, v5+v59, v50)
					mBase = m.M
					v62 = v61
				} else {
					v62 = v52
				}
				return v62
			}
		} else {
			if v9&int32(1) != 0 {
				v29 = int32(1)
				v38 = int32(base.Ui32(v9)>>(uint(v29)%32)) - v29
			} else {
				v33 = *(*int32)(unsafe.Add(mBase, uint32(v5)))
				v38 = int32(base.Ui32(v33)>>(uint(int32(2))%32)) - int32(4)
			}
			if v38 < int32(64) {
				v50 = v38
				v52 = F_palloc0(m, int32(64))
				mBase = m.M
				v53 = m.ExcPending
				if v53 != 0 {
					return int32(0)
				} else {
					v54 = int32(1)
					v56 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v5))))
					if v56&v54 != 0 {
						v59 = v54
					} else {
						v59 = int32(4)
					}
					if v50 != 0 {
						v61 = F__emscripten_memcpy_bulkmem(m, v52, v5+v59, v50)
						mBase = m.M
						v62 = v61
					} else {
						v62 = v52
					}
					return v62
				}
			} else {
				v41 = int32(1)
				if v9&v41 != 0 {
					v45 = v41
				} else {
					v45 = int32(4)
				}
				v48 = F_pg_mbcliplen(m, v5+v45, v38, int32(63))
				mBase = m.M
				v49 = m.ExcPending
				if v49 != 0 {
					return int32(0)
				} else {
					v50 = v48
					v52 = F_palloc0(m, int32(64))
					mBase = m.M
					v53 = m.ExcPending
					if v53 != 0 {
						return int32(0)
					} else {
						v54 = int32(1)
						v56 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v5))))
						if v56&v54 != 0 {
							v59 = v54
						} else {
							v59 = int32(4)
						}
						if v50 != 0 {
							v61 = F__emscripten_memcpy_bulkmem(m, v52, v5+v59, v50)
							mBase = m.M
							v62 = v61
						} else {
							v62 = v52
						}
						return v62
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
