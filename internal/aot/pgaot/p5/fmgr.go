package p5

import (
	base "github.com/shibukawa/pgmem/internal/aot/pgaot/base"
	"unsafe"
)

func F_fmgr_c_validator(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v4 int32
	_ = v4
	var v6 int32
	_ = v6
	var v8 int32
	_ = v8
	var v9 int32
	_ = v9
	var v10 int32
	_ = v10
	var v11 int32
	_ = v11
	var v14 int32
	_ = v14
	var v16 int32
	_ = v16
	var v17 int32
	_ = v17
	var v22 int32
	_ = v22
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
	var v31 int32
	_ = v31
	var v35 int32
	_ = v35
	var v36 int32
	_ = v36
	var v37 int32
	_ = v37
	var v38 int32
	_ = v38
	var v39 int32
	_ = v39
	var v41 int32
	_ = v41
	var v52 int32
	_ = v52
	var v56 int32
	_ = v56
	var v61 int32
	_ = v61
	v4 = m.G0
	v6 = v4 - int32(16)
	m.G0 = v6
	v8 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v9 = *(*int32)(unsafe.Add(mBase, uint32(v8)+4))
	v10 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v11 = F_CheckFunctionValidatorAccess(m, v9, v10)
	mBase = m.M
	v14 = m.ExcPending
	if v14 != 0 {
		return int32(0)
	} else {
		if v11 != 0 {
			v16 = F_SearchSysCache1(m, int32(47), v10)
			mBase = m.M
			v17 = m.ExcPending
			if v17 != 0 {
				return int32(0)
			} else {
				if v16 == int32(0) {
					F_errstart_cold(m, int32(21), int32(0))
					mBase = m.M
					v52 = m.ExcPending
					if v52 != 0 {
						return int32(0)
					} else {
						*(*int32)(unsafe.Add(mBase, uint32(v6))) = v10
						F_errmsg_internal(m, int32(42248), v6)
						mBase = m.M
						v56 = m.ExcPending
						if v56 != 0 {
							return int32(0)
						} else {
							F_errfinish(m, int32(476774), int32(809), int32(199265))
							mBase = m.M
							v61 = m.ExcPending
							if v61 != 0 {
								return int32(0)
							} else {
								base.Wasm_trap_unreachable()
								for {
								}
							}
						}
					}
				} else {
					v22 = F_SysCacheGetAttrNotNull(m, int32(47), v16, int32(26))
					mBase = m.M
					v23 = m.ExcPending
					if v23 != 0 {
						return int32(0)
					} else {
						v24 = F_text_to_cstring(m, v22)
						mBase = m.M
						v25 = m.ExcPending
						if v25 != 0 {
							return int32(0)
						} else {
							v28 = F_SysCacheGetAttrNotNull(m, int32(47), v16, int32(27))
							mBase = m.M
							v29 = m.ExcPending
							if v29 != 0 {
								return int32(0)
							} else {
								v30 = F_text_to_cstring(m, v28)
								mBase = m.M
								v31 = m.ExcPending
								if v31 != 0 {
									return int32(0)
								} else {
									v35 = F_load_external_function(m, v30, v24, int32(1), v6+int32(12))
									mBase = m.M
									v36 = m.ExcPending
									if v36 != 0 {
										return int32(0)
									} else {
										v37 = *(*int32)(unsafe.Add(mBase, uint32(v6)+12))
										v38 = F_fetch_finfo_record(m, v37, v24)
										mBase = m.M
										v39 = m.ExcPending
										if v39 != 0 {
											return int32(0)
										} else {
											F_ReleaseCatCache(m, v16)
											mBase = m.M
											v41 = m.ExcPending
											if v41 != 0 {
												return int32(0)
											} else {
												m.G0 = v6 + int32(16)
												return int32(0)
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
			m.G0 = v6 + int32(16)
			return int32(0)
		}
	}
}
func F_fmgr_info(m *base.Module, l0 int32, l1 int32) {
	mBase := m.M
	_ = mBase
	var v4 int32
	_ = v4
	var v7 int32
	_ = v7
	v4 = *(*int32)(unsafe.Add(mBase, _consts[9]))
	F_fmgr_info_cxt_security(m, l0, l1, v4, int32(0))
	mBase = m.M
	v7 = m.ExcPending
	if v7 != 0 {
		return
	} else {
		return
	}
}
func F_fmgr_info_cxt(m *base.Module, l0 int32, l1 int32, l2 int32) {
	var v6 int32
	_ = v6
	F_fmgr_info_cxt_security(m, l0, l1, l2, int32(0))
	v6 = m.ExcPending
	if v6 != 0 {
		return
	} else {
		return
	}
}
func F_fmgr_security_definer(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v2 int32
	_ = v2
	var v24 int32
	_ = v24
	var v26 int32
	_ = v26
	var v31 int32
	_ = v31
	var v32 int32
	_ = v32
	var v33 int32
	_ = v33
	var v34 int32
	_ = v34
	var v35 int32
	_ = v35
	var v36 int32
	_ = v36
	var v38 int32
	_ = v38
	var v39 int32
	_ = v39
	var v40 int32
	_ = v40
	var v41 int32
	_ = v41
	var v42 int32
	_ = v42
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
	var v58 int32
	_ = v58
	var v61 int32
	_ = v61
	var v64 int32
	_ = v64
	var v67 int32
	_ = v67
	var v70 int32
	_ = v70
	var v72 int32
	_ = v72
	var v73 int32
	_ = v73
	var v76 int32
	_ = v76
	var v89 int32
	_ = v89
	var v90 int32
	_ = v90
	var v92 int32
	_ = v92
	var v93 int32
	_ = v93
	var v95 int32
	_ = v95
	var v109 int32
	_ = v109
	var v110 int32
	_ = v110
	var v111 int32
	_ = v111
	var v112 int32
	_ = v112
	var v114 int32
	_ = v114
	var v115 int32
	_ = v115
	var v128 int32
	_ = v128
	var v129 int32
	_ = v129
	var v146 int32
	_ = v146
	var v147 int32
	_ = v147
	var v148 int32
	_ = v148
	var v163 int32
	_ = v163
	var v179 int32
	_ = v179
	var v180 int32
	_ = v180
	var v181 int32
	_ = v181
	var v182 int32
	_ = v182
	var v183 int32
	_ = v183
	var v186 int32
	_ = v186
	var v187 int32
	_ = v187
	var v202 int32
	_ = v202
	var v203 int32
	_ = v203
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
	var v211 int32
	_ = v211
	var v225 int32
	_ = v225
	var v226 int32
	_ = v226
	var v227 int32
	_ = v227
	var v228 int32
	_ = v228
	var v229 int32
	_ = v229
	var v247 int32
	_ = v247
	var v248 int32
	_ = v248
	var v249 int32
	_ = v249
	var v251 int32
	_ = v251
	var v252 int32
	_ = v252
	var v255 int32
	_ = v255
	var v256 int32
	_ = v256
	var v261 int32
	_ = v261
	var v282 int32
	_ = v282
	var v286 int32
	_ = v286
	var v287 int32
	_ = v287
	var v288 int32
	_ = v288
	var v300 int32
	_ = v300
	var v303 int32
	_ = v303
	var v304 int32
	_ = v304
	var v305 int32
	_ = v305
	var v311 int32
	_ = v311
	var v323 int32
	_ = v323
	var v324 int32
	_ = v324
	var v325 int32
	_ = v325
	var v328 int32
	_ = v328
	var v329 int32
	_ = v329
	var v371 int32
	_ = v371
	var v392 int32
	_ = v392
	var v393 int32
	_ = v393
	var v394 int32
	_ = v394
	var v410 int32
	_ = v410
	var v431 int32
	_ = v431
	var v434 int32
	_ = v434
	var v436 int32
	_ = v436
	var v437 int32
	_ = v437
	var v450 int32
	_ = v450
	var v452 int32
	_ = v452
	var v454 int32
	_ = v454
	var v457 int32
	_ = v457
	var v458 int32
	_ = v458
	var v459 int32
	_ = v459
	var v460 int32
	_ = v460
	var v461 int32
	_ = v461
	var v462 int32
	_ = v462
	var v463 int32
	_ = v463
	var v483 int32
	_ = v483
	var v484 int32
	_ = v484
	var v486 int32
	_ = v486
	var v488 int32
	_ = v488
	var v492 int32
	_ = v492
	var v513 int32
	_ = v513
	var v517 int32
	_ = v517
	var v519 int32
	_ = v519
	var v523 int32
	_ = v523
	var v524 int32
	_ = v524
	var v527 int32
	_ = v527
	var v529 int32
	_ = v529
	var v533 int32
	_ = v533
	var v536 int32
	_ = v536
	var v542 int32
	_ = v542
	var v545 int32
	_ = v545
	var v548 int32
	_ = v548
	var v549 int32
	_ = v549
	var v566 int32
	_ = v566
	var v571 int32
	_ = v571
	var v573 int32
	_ = v573
	var v574 int32
	_ = v574
	var v591 int32
	_ = v591
	var v592 int32
	_ = v592
	var v593 int32
	_ = v593
	var v594 int32
	_ = v594
	var v595 int32
	_ = v595
	var v608 int32
	_ = v608
	var v622 int32
	_ = v622
	var v626 int32
	_ = v626
	var v628 int32
	_ = v628
	var v629 int32
	_ = v629
	var v634 int32
	_ = v634
	var v635 int32
	_ = v635
	var v636 int32
	_ = v636
	var v637 int32
	_ = v637
	var v638 int32
	_ = v638
	var v639 int32
	_ = v639
	var v641 int32
	_ = v641
	var v642 int32
	_ = v642
	var v643 int32
	_ = v643
	var v644 int32
	_ = v644
	var v645 int32
	_ = v645
	var v646 int32
	_ = v646
	var v651 int32
	_ = v651
	var v659 int32
	_ = v659
	var v673 int32
	_ = v673
	var v674 int32
	_ = v674
	var v675 int32
	_ = v675
	var v687 int32
	_ = v687
	var v688 int32
	_ = v688
	var v689 int32
	_ = v689
	var v690 int32
	_ = v690
	var v693 int32
	_ = v693
	var v696 int32
	_ = v696
	var v699 int32
	_ = v699
	var v717 int32
	_ = v717
	var v719 int32
	_ = v719
	var v721 int32
	_ = v721
	var v724 int32
	_ = v724
	var v725 int64
	_ = v725
	var v727 int64
	_ = v727
	var v728 int64
	_ = v728
	var v729 int64
	_ = v729
	var v733 int64
	_ = v733
	var v734 int64
	_ = v734
	var v737 int64
	_ = v737
	var v739 int64
	_ = v739
	var v744 int64
	_ = v744
	var v761 int32
	_ = v761
	var v762 int32
	_ = v762
	var v776 int32
	_ = v776
	var v777 int32
	_ = v777
	var v778 int32
	_ = v778
	var v779 int32
	_ = v779
	var v780 int32
	_ = v780
	var v799 int32
	_ = v799
	var v800 int32
	_ = v800
	var v817 int32
	_ = v817
	var v830 int32
	_ = v830
	var v831 int32
	_ = v831
	var v848 int32
	_ = v848
	var v863 int32
	_ = v863
	var v883 int32
	_ = v883
	var v887 int32
	_ = v887
	var v888 int64
	_ = v888
	var v892 int32
	_ = v892
	var v894 int32
	_ = v894
	var v895 int32
	_ = v895
	var v899 int32
	_ = v899
	var v901 int32
	_ = v901
	var v902 int32
	_ = v902
	var v903 int32
	_ = v903
	var v904 int32
	_ = v904
	var v905 int32
	_ = v905
	var v906 int32
	_ = v906
	var v907 int32
	_ = v907
	var v908 int32
	_ = v908
	var v909 int32
	_ = v909
	var v910 int32
	_ = v910
	var v911 int32
	_ = v911
	var v912 int32
	_ = v912
	var v913 int32
	_ = v913
	var v917 int32
	_ = v917
	v2 = int32(0)
	v24 = m.G0
	v26 = v24 + int32(-64)
	m.G0 = v26
	v31 = v2
	v32 = v2
	v33 = v2
	v34 = v2
	v35 = v2
	v36 = v2
	v38 = v2
	v39 = v2
	v40 = v2
	v41 = v2
	v42 = v2
	v43 = v2
	v44 = int32(-1)
	v48 = v26
	goto L1
L1:
	;
	goto L4
L2:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L3:
	;
	goto L2
L4:
	;
	if v44 == int32(1) {
		v634 = v31
		v635 = v32
		v636 = v33
		v637 = v34
		v638 = v35
		v639 = v36
		v641 = v38
		v642 = v39
		v643 = v40
		v644 = v41
		v645 = v42
		v646 = v43
		v651 = v48
		goto L7
	} else {
		goto L8
	}
L5:
	;
	goto L3
L6:
	;
	v887 = int32(m.ExcTag)
	v888 = int64(m.ExcVals[0])
	m.ExcPending = 0
	if v887 == int32(0) {
		goto L116
	} else {
		goto L117
	}
L7:
	;
	if v634 == int32(0) {
		goto L84
	} else {
		goto L85
	}
L8:
	;
	v54 = int32(16)
	v55 = v48 - v54
	m.G0 = v55
	v58 = v55 - v54
	m.G0 = v58
	v61 = v58 - v54
	m.G0 = v61
	v64 = v61 - int32(32)
	m.G0 = v64
	v67 = v64 - v54
	m.G0 = v67
	v70 = v67 - int32(160)
	m.G0 = v70
	v72 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v73 = *(*int32)(unsafe.Add(mBase, uint32(v72)+16))
	if v73 == int32(0) {
		goto L12
	} else {
		goto L13
	}
L9:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v26)+24)) = v38
	*(*int32)(unsafe.Add(mBase, uint32(v26)+20)) = v39
	*(*int32)(unsafe.Add(mBase, uint32(v26)+28)) = v40
	*(*int32)(unsafe.Add(mBase, uint32(v26)+32)) = v41
	*(*int32)(unsafe.Add(mBase, uint32(v26)+36)) = v42
	*(*int32)(unsafe.Add(mBase, uint32(v26)+40)) = v410
	*(*int32)(unsafe.Add(mBase, uint32(v26)+44)) = v70
	*(*int32)(unsafe.Add(mBase, uint32(v26)+48)) = v64
	*(*int32)(unsafe.Add(mBase, uint32(v26)+52)) = v61
	*(*int32)(unsafe.Add(mBase, uint32(v26)+56)) = v58
	*(*int32)(unsafe.Add(mBase, uint32(v26)+60)) = v55
	v431 = *(*int32)(unsafe.Add(mBase, _consts[239]))
	*(*int32)(unsafe.Add(mBase, uint32(v58))) = v431
	v434 = *(*int32)(unsafe.Add(mBase, _consts[240]))
	*(*int32)(unsafe.Add(mBase, uint32(v61))) = v434
	goto L47
L10:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v26)+24)) = v38
	*(*int32)(unsafe.Add(mBase, uint32(v26)+20)) = v39
	*(*int32)(unsafe.Add(mBase, uint32(v26)+28)) = v40
	*(*int32)(unsafe.Add(mBase, uint32(v26)+32)) = v41
	*(*int32)(unsafe.Add(mBase, uint32(v26)+36)) = v42
	*(*int32)(unsafe.Add(mBase, uint32(v26)+40)) = v371
	*(*int32)(unsafe.Add(mBase, uint32(v26)+44)) = v70
	*(*int32)(unsafe.Add(mBase, uint32(v26)+48)) = v64
	*(*int32)(unsafe.Add(mBase, uint32(v26)+52)) = v61
	*(*int32)(unsafe.Add(mBase, uint32(v26)+56)) = v58
	*(*int32)(unsafe.Add(mBase, uint32(v26)+60)) = v55
	F_ReleaseCatCache(m, v128)
	mBase = m.M
	v392 = m.ExcPending
	if v392 != 0 {
		v883 = v70
		goto L6
	} else {
		goto L46
	}
L11:
	;
	*(*int32)(unsafe.Add(mBase, _consts[9])) = v206
	v371 = v228
	goto L10
L12:
	;
	v76 = *(*int32)(unsafe.Add(mBase, uint32(v72)+20))
	*(*int32)(unsafe.Add(mBase, uint32(v26)+24)) = v38
	*(*int32)(unsafe.Add(mBase, uint32(v26)+20)) = v39
	*(*int32)(unsafe.Add(mBase, uint32(v26)+28)) = v40
	*(*int32)(unsafe.Add(mBase, uint32(v26)+32)) = v41
	*(*int32)(unsafe.Add(mBase, uint32(v26)+36)) = v42
	*(*int32)(unsafe.Add(mBase, uint32(v26)+40)) = v43
	*(*int32)(unsafe.Add(mBase, uint32(v26)+44)) = v70
	*(*int32)(unsafe.Add(mBase, uint32(v26)+48)) = v64
	*(*int32)(unsafe.Add(mBase, uint32(v26)+52)) = v61
	*(*int32)(unsafe.Add(mBase, uint32(v26)+56)) = v58
	*(*int32)(unsafe.Add(mBase, uint32(v26)+60)) = v55
	v89 = F_MemoryContextAllocZero(m, v76, int32(48))
	mBase = m.M
	v90 = m.ExcPending
	if v90 != 0 {
		v883 = v70
		goto L6
	} else {
		goto L15
	}
L13:
	;
	goto L14
L14:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v55))) = v73
	v410 = v43
	goto L9
L15:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v55))) = v89
	v92 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v93 = *(*int32)(unsafe.Add(mBase, uint32(v92)+4))
	v95 = *(*int32)(unsafe.Add(mBase, uint32(v92)+20))
	*(*int32)(unsafe.Add(mBase, uint32(v26)+24)) = v38
	*(*int32)(unsafe.Add(mBase, uint32(v26)+20)) = v39
	*(*int32)(unsafe.Add(mBase, uint32(v26)+28)) = v40
	*(*int32)(unsafe.Add(mBase, uint32(v26)+32)) = v41
	*(*int32)(unsafe.Add(mBase, uint32(v26)+36)) = v42
	*(*int32)(unsafe.Add(mBase, uint32(v26)+40)) = v43
	*(*int32)(unsafe.Add(mBase, uint32(v26)+44)) = v70
	*(*int32)(unsafe.Add(mBase, uint32(v26)+48)) = v64
	*(*int32)(unsafe.Add(mBase, uint32(v26)+52)) = v61
	*(*int32)(unsafe.Add(mBase, uint32(v26)+56)) = v58
	*(*int32)(unsafe.Add(mBase, uint32(v26)+60)) = v55
	F_fmgr_info_cxt_security(m, v93, v89, v95, int32(1))
	mBase = m.M
	v109 = m.ExcPending
	if v109 != 0 {
		v883 = v70
		goto L6
	} else {
		goto L16
	}
L16:
	;
	v110 = *(*int32)(unsafe.Add(mBase, uint32(v55)))
	v111 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v112 = *(*int32)(unsafe.Add(mBase, uint32(v111)+24))
	*(*int32)(unsafe.Add(mBase, uint32(v110)+24)) = v112
	v114 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v115 = *(*int32)(unsafe.Add(mBase, uint32(v114)+4))
	*(*int32)(unsafe.Add(mBase, uint32(v26)+24)) = v38
	*(*int32)(unsafe.Add(mBase, uint32(v26)+20)) = v39
	*(*int32)(unsafe.Add(mBase, uint32(v26)+28)) = v40
	*(*int32)(unsafe.Add(mBase, uint32(v26)+32)) = v41
	*(*int32)(unsafe.Add(mBase, uint32(v26)+36)) = v42
	*(*int32)(unsafe.Add(mBase, uint32(v26)+40)) = v43
	*(*int32)(unsafe.Add(mBase, uint32(v26)+44)) = v70
	*(*int32)(unsafe.Add(mBase, uint32(v26)+48)) = v64
	*(*int32)(unsafe.Add(mBase, uint32(v26)+52)) = v61
	*(*int32)(unsafe.Add(mBase, uint32(v26)+56)) = v58
	*(*int32)(unsafe.Add(mBase, uint32(v26)+60)) = v55
	v128 = F_SearchSysCache1(m, int32(47), v115)
	mBase = m.M
	v129 = m.ExcPending
	if v129 != 0 {
		v883 = v70
		goto L6
	} else {
		goto L17
	}
L17:
	;
	if v128 == int32(0) {
		goto L18
	} else {
		goto L19
	}
L18:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v26)+24)) = v38
	*(*int32)(unsafe.Add(mBase, uint32(v26)+20)) = v39
	*(*int32)(unsafe.Add(mBase, uint32(v26)+28)) = v40
	*(*int32)(unsafe.Add(mBase, uint32(v26)+32)) = v41
	*(*int32)(unsafe.Add(mBase, uint32(v26)+36)) = v42
	*(*int32)(unsafe.Add(mBase, uint32(v26)+40)) = v43
	*(*int32)(unsafe.Add(mBase, uint32(v26)+44)) = v70
	*(*int32)(unsafe.Add(mBase, uint32(v26)+48)) = v64
	*(*int32)(unsafe.Add(mBase, uint32(v26)+52)) = v61
	*(*int32)(unsafe.Add(mBase, uint32(v26)+56)) = v58
	*(*int32)(unsafe.Add(mBase, uint32(v26)+60)) = v55
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v146 = m.ExcPending
	if v146 != 0 {
		v883 = v70
		goto L6
	} else {
		goto L21
	}
L19:
	;
	goto L20
L20:
	;
	v180 = *(*int32)(unsafe.Add(mBase, uint32(v128)+16))
	v181 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v180)+22)))
	v182 = v180 + v181
	v183 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v182)+97)))
	if v183 == int32(1) {
		goto L24
	} else {
		goto L25
	}
L21:
	;
	v147 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v148 = *(*int32)(unsafe.Add(mBase, uint32(v147)+4))
	*(*int32)(unsafe.Add(mBase, uint32(v26)+24)) = v38
	*(*int32)(unsafe.Add(mBase, uint32(v26)+20)) = v39
	*(*int32)(unsafe.Add(mBase, uint32(v26)+28)) = v40
	*(*int32)(unsafe.Add(mBase, uint32(v26)+32)) = v41
	*(*int32)(unsafe.Add(mBase, uint32(v26)+36)) = v42
	*(*int32)(unsafe.Add(mBase, uint32(v26)+40)) = v43
	*(*int32)(unsafe.Add(mBase, uint32(v26)+44)) = v70
	*(*int32)(unsafe.Add(mBase, uint32(v26)+48)) = v64
	*(*int32)(unsafe.Add(mBase, uint32(v26)+52)) = v61
	*(*int32)(unsafe.Add(mBase, uint32(v26)+56)) = v58
	*(*int32)(unsafe.Add(mBase, uint32(v26)+60)) = v55
	*(*int32)(unsafe.Add(mBase, uint32(v26))) = v148
	F_errmsg_internal(m, int32(42248), v26)
	mBase = m.M
	v163 = m.ExcPending
	if v163 != 0 {
		v883 = v70
		goto L6
	} else {
		goto L22
	}
L22:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v26)+24)) = v38
	*(*int32)(unsafe.Add(mBase, uint32(v26)+20)) = v39
	*(*int32)(unsafe.Add(mBase, uint32(v26)+28)) = v40
	*(*int32)(unsafe.Add(mBase, uint32(v26)+32)) = v41
	*(*int32)(unsafe.Add(mBase, uint32(v26)+36)) = v42
	*(*int32)(unsafe.Add(mBase, uint32(v26)+40)) = v43
	*(*int32)(unsafe.Add(mBase, uint32(v26)+44)) = v70
	*(*int32)(unsafe.Add(mBase, uint32(v26)+48)) = v64
	*(*int32)(unsafe.Add(mBase, uint32(v26)+52)) = v61
	*(*int32)(unsafe.Add(mBase, uint32(v26)+56)) = v58
	*(*int32)(unsafe.Add(mBase, uint32(v26)+60)) = v55
	F_errfinish(m, int32(472279), int32(664), int32(208148))
	mBase = m.M
	v179 = m.ExcPending
	if v179 != 0 {
		v883 = v70
		goto L6
	} else {
		goto L23
	}
L23:
	;
	goto L3
L24:
	;
	v186 = *(*int32)(unsafe.Add(mBase, uint32(v55)))
	v187 = *(*int32)(unsafe.Add(mBase, uint32(v182)+72))
	*(*int32)(unsafe.Add(mBase, uint32(v186)+28)) = v187
	goto L26
L25:
	;
	goto L26
L26:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v26)+24)) = v38
	*(*int32)(unsafe.Add(mBase, uint32(v26)+20)) = v39
	*(*int32)(unsafe.Add(mBase, uint32(v26)+28)) = v40
	*(*int32)(unsafe.Add(mBase, uint32(v26)+32)) = v41
	*(*int32)(unsafe.Add(mBase, uint32(v26)+36)) = v42
	*(*int32)(unsafe.Add(mBase, uint32(v26)+40)) = v43
	*(*int32)(unsafe.Add(mBase, uint32(v26)+44)) = v70
	*(*int32)(unsafe.Add(mBase, uint32(v26)+48)) = v64
	*(*int32)(unsafe.Add(mBase, uint32(v26)+52)) = v61
	*(*int32)(unsafe.Add(mBase, uint32(v26)+56)) = v58
	*(*int32)(unsafe.Add(mBase, uint32(v26)+60)) = v55
	v202 = F_SysCacheGetAttr(m, int32(47), v128, int32(29), v67)
	mBase = m.M
	v203 = m.ExcPending
	if v203 != 0 {
		v883 = v70
		goto L6
	} else {
		goto L27
	}
L27:
	;
	v204 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v67))))
	if v204 != 0 {
		v371 = v43
		goto L10
	} else {
		goto L28
	}
L28:
	;
	v205 = int32(4442992)
	v206 = *(*int32)(unsafe.Add(mBase, _consts[9]))
	v208 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v209 = *(*int32)(unsafe.Add(mBase, uint32(v208)+20))
	*(*int32)(unsafe.Add(mBase, _consts[9])) = v209
	v211 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v202))))
	if v211&int32(3) != 0 {
		goto L29
	} else {
		goto L30
	}
L29:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v26)+24)) = v38
	*(*int32)(unsafe.Add(mBase, uint32(v26)+20)) = v39
	*(*int32)(unsafe.Add(mBase, uint32(v26)+28)) = v40
	*(*int32)(unsafe.Add(mBase, uint32(v26)+32)) = v41
	*(*int32)(unsafe.Add(mBase, uint32(v26)+36)) = v42
	*(*int32)(unsafe.Add(mBase, uint32(v26)+40)) = v43
	*(*int32)(unsafe.Add(mBase, uint32(v26)+44)) = v70
	*(*int32)(unsafe.Add(mBase, uint32(v26)+48)) = v64
	*(*int32)(unsafe.Add(mBase, uint32(v26)+52)) = v61
	*(*int32)(unsafe.Add(mBase, uint32(v26)+56)) = v58
	*(*int32)(unsafe.Add(mBase, uint32(v26)+60)) = v55
	v225 = F_detoast_attr(m, v202)
	mBase = m.M
	v226 = m.ExcPending
	if v226 != 0 {
		v883 = v70
		goto L6
	} else {
		goto L32
	}
L30:
	;
	v227 = v202
	v228 = v43
	goto L31
L31:
	;
	v229 = *(*int32)(unsafe.Add(mBase, uint32(v55)))
	*(*int32)(unsafe.Add(mBase, uint32(v26)+24)) = v38
	*(*int32)(unsafe.Add(mBase, uint32(v26)+20)) = v39
	*(*int32)(unsafe.Add(mBase, uint32(v26)+28)) = v40
	*(*int32)(unsafe.Add(mBase, uint32(v26)+32)) = v41
	*(*int32)(unsafe.Add(mBase, uint32(v26)+36)) = v42
	*(*int32)(unsafe.Add(mBase, uint32(v26)+40)) = v228
	*(*int32)(unsafe.Add(mBase, uint32(v26)+44)) = v70
	*(*int32)(unsafe.Add(mBase, uint32(v26)+48)) = v64
	*(*int32)(unsafe.Add(mBase, uint32(v26)+52)) = v61
	*(*int32)(unsafe.Add(mBase, uint32(v26)+56)) = v58
	*(*int32)(unsafe.Add(mBase, uint32(v26)+60)) = v55
	F_TransformGUCArray(m, v227, v229+int32(32), v229+int32(40))
	mBase = m.M
	v247 = m.ExcPending
	if v247 != 0 {
		v883 = v70
		goto L6
	} else {
		goto L33
	}
L32:
	;
	v227 = v225
	v228 = v225
	goto L31
L33:
	;
	v248 = *(*int32)(unsafe.Add(mBase, uint32(v55)))
	v249 = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v248)+36)) = v249
	v251 = *(*int32)(unsafe.Add(mBase, uint32(v55)))
	v252 = *(*int32)(unsafe.Add(mBase, uint32(v251)+32))
	if v252 == v249 {
		goto L11
	} else {
		goto L34
	}
L34:
	;
	v255 = int32(0)
	v256 = *(*int32)(unsafe.Add(mBase, uint32(v252)+4))
	if v256 <= v255 {
		goto L11
	} else {
		goto L35
	}
L35:
	;
	v261 = v255
	goto L36
L36:
	;
	v282 = *(*int32)(unsafe.Add(mBase, uint32(v252)+12))
	v286 = *(*int32)(unsafe.Add(mBase, uint32(v282+v261<<(uint(int32(2))%32))))
	v287 = *(*int32)(unsafe.Add(mBase, uint32(v55)))
	v288 = *(*int32)(unsafe.Add(mBase, uint32(v287)+36))
	*(*int32)(unsafe.Add(mBase, uint32(v26)+24)) = v38
	*(*int32)(unsafe.Add(mBase, uint32(v26)+20)) = v39
	*(*int32)(unsafe.Add(mBase, uint32(v26)+28)) = v40
	*(*int32)(unsafe.Add(mBase, uint32(v26)+32)) = v41
	*(*int32)(unsafe.Add(mBase, uint32(v26)+36)) = v42
	*(*int32)(unsafe.Add(mBase, uint32(v26)+40)) = v228
	*(*int32)(unsafe.Add(mBase, uint32(v26)+44)) = v70
	*(*int32)(unsafe.Add(mBase, uint32(v26)+48)) = v64
	*(*int32)(unsafe.Add(mBase, uint32(v26)+52)) = v61
	*(*int32)(unsafe.Add(mBase, uint32(v26)+56)) = v58
	*(*int32)(unsafe.Add(mBase, uint32(v26)+60)) = v55
	v300 = int32(0)
	v303 = F_find_option(m, v286, v300, v300, v300)
	mBase = m.M
	v304 = m.ExcPending
	if v304 != 0 {
		v883 = v70
		goto L6
	} else {
		goto L39
	}
L37:
	;
	goto L11
L38:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v26)+24)) = v38
	*(*int32)(unsafe.Add(mBase, uint32(v26)+20)) = v39
	*(*int32)(unsafe.Add(mBase, uint32(v26)+28)) = v40
	*(*int32)(unsafe.Add(mBase, uint32(v26)+32)) = v41
	*(*int32)(unsafe.Add(mBase, uint32(v26)+36)) = v42
	*(*int32)(unsafe.Add(mBase, uint32(v26)+40)) = v228
	*(*int32)(unsafe.Add(mBase, uint32(v26)+44)) = v70
	*(*int32)(unsafe.Add(mBase, uint32(v26)+48)) = v64
	*(*int32)(unsafe.Add(mBase, uint32(v26)+52)) = v61
	*(*int32)(unsafe.Add(mBase, uint32(v26)+56)) = v58
	*(*int32)(unsafe.Add(mBase, uint32(v26)+60)) = v55
	v323 = F_lappend(m, v288, v311)
	mBase = m.M
	v324 = m.ExcPending
	if v324 != 0 {
		v883 = v70
		goto L6
	} else {
		goto L44
	}
L39:
	;
	if v303 != 0 {
		goto L40
	} else {
		goto L41
	}
L40:
	;
	v305 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v303)+21)))
	if v305&int32(2) == int32(0) {
		v311 = v303
		goto L38
	} else {
		goto L43
	}
L41:
	;
	goto L42
L42:
	;
	v311 = int32(0)
	goto L38
L43:
	;
	goto L42
L44:
	;
	v325 = *(*int32)(unsafe.Add(mBase, uint32(v55)))
	*(*int32)(unsafe.Add(mBase, uint32(v325)+36)) = v323
	v328 = v261 + int32(1)
	v329 = *(*int32)(unsafe.Add(mBase, uint32(v252)+4))
	if v328 < v329 {
		v261 = v328
		goto L36
	} else {
		goto L45
	}
L45:
	;
	goto L37
L46:
	;
	v393 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v394 = *(*int32)(unsafe.Add(mBase, uint32(v55)))
	*(*int32)(unsafe.Add(mBase, uint32(v393)+16)) = v394
	v410 = v371
	goto L9
L47:
	;
	v436 = *(*int32)(unsafe.Add(mBase, uint32(v55)))
	v437 = *(*int32)(unsafe.Add(mBase, uint32(v436)+32))
	if v437 != 0 {
		goto L48
	} else {
		goto L49
	}
L48:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v26)+24)) = v38
	*(*int32)(unsafe.Add(mBase, uint32(v26)+20)) = v39
	*(*int32)(unsafe.Add(mBase, uint32(v26)+28)) = v40
	*(*int32)(unsafe.Add(mBase, uint32(v26)+32)) = v41
	*(*int32)(unsafe.Add(mBase, uint32(v26)+36)) = v42
	*(*int32)(unsafe.Add(mBase, uint32(v26)+40)) = v410
	*(*int32)(unsafe.Add(mBase, uint32(v26)+44)) = v70
	*(*int32)(unsafe.Add(mBase, uint32(v26)+48)) = v64
	*(*int32)(unsafe.Add(mBase, uint32(v26)+52)) = v61
	*(*int32)(unsafe.Add(mBase, uint32(v26)+56)) = v58
	*(*int32)(unsafe.Add(mBase, uint32(v26)+60)) = v55
	v450 = int32(4441032)
	v452 = *(*int32)(unsafe.Add(mBase, _consts[241]))
	v454 = v452 + int32(1)
	*(*int32)(unsafe.Add(mBase, _consts[241])) = v454
	goto L51
L49:
	;
	v457 = v42
	v458 = int32(0)
	goto L50
L50:
	;
	v459 = *(*int32)(unsafe.Add(mBase, uint32(v55)))
	v460 = *(*int32)(unsafe.Add(mBase, uint32(v459)+28))
	if v460 != 0 {
		goto L52
	} else {
		goto L53
	}
L51:
	;
	v457 = v454
	v458 = v454
	goto L50
L52:
	;
	v461 = *(*int32)(unsafe.Add(mBase, uint32(v55)))
	v462 = *(*int32)(unsafe.Add(mBase, uint32(v461)+28))
	v463 = *(*int32)(unsafe.Add(mBase, uint32(v61)))
	*(*int32)(unsafe.Add(mBase, uint32(v26)+24)) = v38
	*(*int32)(unsafe.Add(mBase, uint32(v26)+20)) = v39
	*(*int32)(unsafe.Add(mBase, uint32(v26)+28)) = v40
	*(*int32)(unsafe.Add(mBase, uint32(v26)+32)) = v458
	*(*int32)(unsafe.Add(mBase, uint32(v26)+36)) = v457
	*(*int32)(unsafe.Add(mBase, uint32(v26)+40)) = v410
	*(*int32)(unsafe.Add(mBase, uint32(v26)+44)) = v70
	*(*int32)(unsafe.Add(mBase, uint32(v26)+48)) = v64
	*(*int32)(unsafe.Add(mBase, uint32(v26)+52)) = v61
	*(*int32)(unsafe.Add(mBase, uint32(v26)+56)) = v58
	*(*int32)(unsafe.Add(mBase, uint32(v26)+60)) = v55
	*(*int32)(unsafe.Add(mBase, _consts[240])) = v463 | int32(1)
	*(*int32)(unsafe.Add(mBase, _consts[239])) = v462
	goto L55
L53:
	;
	goto L54
L54:
	;
	v483 = *(*int32)(unsafe.Add(mBase, uint32(v55)))
	v484 = *(*int32)(unsafe.Add(mBase, uint32(v483)+32))
	v486 = *(*int32)(unsafe.Add(mBase, uint32(v483)+36))
	v488 = *(*int32)(unsafe.Add(mBase, uint32(v483)+40))
	v492 = int32(0)
	goto L56
L55:
	;
	goto L54
L56:
	;
	v513 = int32(0)
	if v484 == v513 {
		v523 = v513
		goto L58
	} else {
		goto L59
	}
L58:
	;
	v524 = int32(0)
	if v486 == v524 {
		v533 = v524
		goto L61
	} else {
		goto L62
	}
L59:
	;
	v517 = *(*int32)(unsafe.Add(mBase, uint32(v484)+4))
	if v517 <= v492 {
		v523 = int32(0)
		goto L58
	} else {
		goto L60
	}
L60:
	;
	v519 = *(*int32)(unsafe.Add(mBase, uint32(v484)+12))
	v523 = v519 + v492<<(uint(int32(2))%32)
	goto L58
L61:
	;
	if v488 == int32(0) {
		goto L65
	} else {
		goto L66
	}
L62:
	;
	v527 = *(*int32)(unsafe.Add(mBase, uint32(v486)+4))
	if v527 <= v492 {
		v533 = v524
		goto L61
	} else {
		goto L63
	}
L63:
	;
	v529 = *(*int32)(unsafe.Add(mBase, uint32(v486)+12))
	v533 = v529 + v492<<(uint(int32(2))%32)
	goto L61
L64:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v26)+24)) = v38
	*(*int32)(unsafe.Add(mBase, uint32(v26)+20)) = v39
	*(*int32)(unsafe.Add(mBase, uint32(v26)+28)) = v40
	*(*int32)(unsafe.Add(mBase, uint32(v26)+32)) = v458
	*(*int32)(unsafe.Add(mBase, uint32(v26)+36)) = v457
	*(*int32)(unsafe.Add(mBase, uint32(v26)+40)) = v410
	*(*int32)(unsafe.Add(mBase, uint32(v26)+44)) = v70
	*(*int32)(unsafe.Add(mBase, uint32(v26)+48)) = v64
	*(*int32)(unsafe.Add(mBase, uint32(v26)+52)) = v61
	*(*int32)(unsafe.Add(mBase, uint32(v26)+56)) = v58
	*(*int32)(unsafe.Add(mBase, uint32(v26)+60)) = v55
	v591 = F_superuser(m)
	mBase = m.M
	v592 = m.ExcPending
	if v592 != 0 {
		v883 = v70
		goto L6
	} else {
		goto L79
	}
L65:
	;
	v548 = *(*int32)(unsafe.Add(mBase, _consts[1216]))
	if v548 != 0 {
		goto L71
	} else {
		goto L72
	}
L66:
	;
	v536 = *(*int32)(unsafe.Add(mBase, uint32(v488)+4))
	if v536 <= v492 {
		goto L65
	} else {
		goto L67
	}
L67:
	;
	if v523 == int32(0) {
		goto L65
	} else {
		goto L68
	}
L68:
	;
	if v533 == int32(0) {
		goto L65
	} else {
		goto L69
	}
L69:
	;
	v542 = *(*int32)(unsafe.Add(mBase, uint32(v488)+12))
	v545 = v542 + v492<<(uint(int32(2))%32)
	if v545 != 0 {
		goto L64
	} else {
		goto L70
	}
L70:
	;
	goto L65
L71:
	;
	v549 = *(*int32)(unsafe.Add(mBase, uint32(v55)))
	*(*int32)(unsafe.Add(mBase, uint32(v26)+24)) = v38
	*(*int32)(unsafe.Add(mBase, uint32(v26)+20)) = v39
	*(*int32)(unsafe.Add(mBase, uint32(v26)+28)) = v40
	*(*int32)(unsafe.Add(mBase, uint32(v26)+32)) = v458
	*(*int32)(unsafe.Add(mBase, uint32(v26)+36)) = v457
	*(*int32)(unsafe.Add(mBase, uint32(v26)+40)) = v410
	*(*int32)(unsafe.Add(mBase, uint32(v26)+44)) = v70
	*(*int32)(unsafe.Add(mBase, uint32(v26)+48)) = v64
	*(*int32)(unsafe.Add(mBase, uint32(v26)+52)) = v61
	*(*int32)(unsafe.Add(mBase, uint32(v26)+56)) = v58
	*(*int32)(unsafe.Add(mBase, uint32(v26)+60)) = v55
	m.T0[v548].(func(*base.Module, int32, int32, int32))(m, int32(0), v549, v549+int32(44))
	mBase = m.M
	v566 = m.ExcPending
	if v566 != 0 {
		v883 = v70
		goto L6
	} else {
		goto L74
	}
L72:
	;
	goto L73
L73:
	;
	v571 = *(*int32)(unsafe.Add(mBase, _consts[141]))
	v573 = *(*int32)(unsafe.Add(mBase, _consts[142]))
	v574 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	goto L75
L74:
	;
	goto L73
L75:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v70)+4)) = int32(1)
	*(*int32)(unsafe.Add(mBase, uint32(v70))) = v24 + int32(-48)
	goto L78
L76:
	;
	v634 = int32(0)
	v635 = v55
	v636 = v61
	v637 = v58
	v638 = v64
	v639 = v70
	v641 = v573
	v642 = v571
	v643 = v574
	v644 = v458
	v645 = v457
	v646 = v410
	v651 = v70
	goto L7
L78:
	;
	goto L76
L79:
	;
	v593 = *(*int32)(unsafe.Add(mBase, uint32(v545)))
	v594 = *(*int32)(unsafe.Add(mBase, uint32(v533)))
	v595 = *(*int32)(unsafe.Add(mBase, uint32(v523)))
	*(*int32)(unsafe.Add(mBase, uint32(v26)+24)) = v38
	*(*int32)(unsafe.Add(mBase, uint32(v26)+20)) = v39
	*(*int32)(unsafe.Add(mBase, uint32(v26)+28)) = v40
	*(*int32)(unsafe.Add(mBase, uint32(v26)+32)) = v458
	*(*int32)(unsafe.Add(mBase, uint32(v26)+36)) = v457
	*(*int32)(unsafe.Add(mBase, uint32(v26)+40)) = v410
	*(*int32)(unsafe.Add(mBase, uint32(v26)+44)) = v70
	*(*int32)(unsafe.Add(mBase, uint32(v26)+48)) = v64
	*(*int32)(unsafe.Add(mBase, uint32(v26)+52)) = v61
	*(*int32)(unsafe.Add(mBase, uint32(v26)+56)) = v58
	*(*int32)(unsafe.Add(mBase, uint32(v26)+60)) = v55
	v608 = *(*int32)(unsafe.Add(mBase, _consts[239]))
	*(*int32)(unsafe.Add(mBase, uint32(v26)+24)) = v38
	*(*int32)(unsafe.Add(mBase, uint32(v26)+20)) = v39
	*(*int32)(unsafe.Add(mBase, uint32(v26)+28)) = v40
	*(*int32)(unsafe.Add(mBase, uint32(v26)+32)) = v458
	*(*int32)(unsafe.Add(mBase, uint32(v26)+36)) = v457
	*(*int32)(unsafe.Add(mBase, uint32(v26)+40)) = v410
	*(*int32)(unsafe.Add(mBase, uint32(v26)+44)) = v70
	*(*int32)(unsafe.Add(mBase, uint32(v26)+48)) = v64
	*(*int32)(unsafe.Add(mBase, uint32(v26)+52)) = v61
	*(*int32)(unsafe.Add(mBase, uint32(v26)+56)) = v58
	*(*int32)(unsafe.Add(mBase, uint32(v26)+60)) = v55
	if v591 != 0 {
		goto L80
	} else {
		goto L81
	}
L80:
	;
	v622 = int32(5)
	goto L82
L81:
	;
	v622 = int32(6)
	goto L82
L82:
	;
	v626 = int32(0)
	v628 = F_set_config_with_handle(m, v595, v594, v593, v622, int32(13), v608, int32(2), int32(1), v626, v626)
	mBase = m.M
	v629 = m.ExcPending
	if v629 != 0 {
		v883 = v70
		goto L6
	} else {
		goto L83
	}
L83:
	;
	v492 = v492 + int32(1)
	goto L56
L84:
	;
	*(*int32)(unsafe.Add(mBase, _consts[142])) = v639
	v659 = *(*int32)(unsafe.Add(mBase, uint32(v635)))
	*(*int32)(unsafe.Add(mBase, uint32(l0))) = v659
	*(*int32)(unsafe.Add(mBase, uint32(v26)+24)) = v641
	*(*int32)(unsafe.Add(mBase, uint32(v26)+20)) = v642
	*(*int32)(unsafe.Add(mBase, uint32(v26)+28)) = v643
	*(*int32)(unsafe.Add(mBase, uint32(v26)+32)) = v644
	*(*int32)(unsafe.Add(mBase, uint32(v26)+36)) = v645
	*(*int32)(unsafe.Add(mBase, uint32(v26)+40)) = v646
	*(*int32)(unsafe.Add(mBase, uint32(v26)+44)) = v639
	*(*int32)(unsafe.Add(mBase, uint32(v26)+48)) = v638
	*(*int32)(unsafe.Add(mBase, uint32(v26)+52)) = v636
	*(*int32)(unsafe.Add(mBase, uint32(v26)+56)) = v637
	*(*int32)(unsafe.Add(mBase, uint32(v26)+60)) = v635
	F_pgstat_init_function_usage(m, l0, v638)
	mBase = m.M
	v673 = m.ExcPending
	if v673 != 0 {
		v883 = v651
		goto L6
	} else {
		goto L87
	}
L85:
	;
	goto L86
L86:
	;
	*(*int32)(unsafe.Add(mBase, _consts[141])) = v642
	*(*int32)(unsafe.Add(mBase, _consts[142])) = v641
	*(*int32)(unsafe.Add(mBase, uint32(l0))) = v643
	v830 = *(*int32)(unsafe.Add(mBase, _consts[1216]))
	if v830 != 0 {
		goto L111
	} else {
		goto L112
	}
L87:
	;
	v674 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v675 = *(*int32)(unsafe.Add(mBase, uint32(v674)))
	*(*int32)(unsafe.Add(mBase, uint32(v26)+24)) = v641
	*(*int32)(unsafe.Add(mBase, uint32(v26)+20)) = v642
	*(*int32)(unsafe.Add(mBase, uint32(v26)+28)) = v643
	*(*int32)(unsafe.Add(mBase, uint32(v26)+32)) = v644
	*(*int32)(unsafe.Add(mBase, uint32(v26)+36)) = v645
	*(*int32)(unsafe.Add(mBase, uint32(v26)+40)) = v646
	*(*int32)(unsafe.Add(mBase, uint32(v26)+44)) = v639
	*(*int32)(unsafe.Add(mBase, uint32(v26)+48)) = v638
	*(*int32)(unsafe.Add(mBase, uint32(v26)+52)) = v636
	*(*int32)(unsafe.Add(mBase, uint32(v26)+56)) = v637
	*(*int32)(unsafe.Add(mBase, uint32(v26)+60)) = v635
	v687 = m.T0[v675].(func(*base.Module, int32) int32)(m, l0)
	mBase = m.M
	v688 = m.ExcPending
	if v688 != 0 {
		v883 = v651
		goto L6
	} else {
		goto L88
	}
L88:
	;
	v689 = int32(1)
	v690 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	if v690 == int32(0) {
		v699 = v689
		goto L89
	} else {
		goto L90
	}
L89:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v26)+24)) = v641
	*(*int32)(unsafe.Add(mBase, uint32(v26)+20)) = v642
	*(*int32)(unsafe.Add(mBase, uint32(v26)+28)) = v643
	*(*int32)(unsafe.Add(mBase, uint32(v26)+32)) = v644
	*(*int32)(unsafe.Add(mBase, uint32(v26)+36)) = v645
	*(*int32)(unsafe.Add(mBase, uint32(v26)+40)) = v646
	*(*int32)(unsafe.Add(mBase, uint32(v26)+44)) = v639
	*(*int32)(unsafe.Add(mBase, uint32(v26)+48)) = v638
	*(*int32)(unsafe.Add(mBase, uint32(v26)+52)) = v636
	*(*int32)(unsafe.Add(mBase, uint32(v26)+56)) = v637
	*(*int32)(unsafe.Add(mBase, uint32(v26)+60)) = v635
	v717 = m.G0
	v719 = v717 - int32(16)
	m.G0 = v719
	v721 = *(*int32)(unsafe.Add(mBase, uint32(v638)))
	if v721 != 0 {
		goto L93
	} else {
		goto L94
	}
L90:
	;
	v693 = *(*int32)(unsafe.Add(mBase, uint32(v690)))
	if v693 != int32(383) {
		v699 = v689
		goto L89
	} else {
		goto L91
	}
L91:
	;
	v696 = *(*int32)(unsafe.Add(mBase, uint32(v690)+20))
	v699 = base.B2i32(v696 != int32(1))
	goto L89
L92:
	;
	*(*int32)(unsafe.Add(mBase, _consts[141])) = v642
	*(*int32)(unsafe.Add(mBase, _consts[142])) = v641
	*(*int32)(unsafe.Add(mBase, uint32(l0))) = v643
	v761 = *(*int32)(unsafe.Add(mBase, uint32(v635)))
	v762 = *(*int32)(unsafe.Add(mBase, uint32(v761)+32))
	if v762 != 0 {
		goto L99
	} else {
		goto L100
	}
L93:
	;
	F___clock_gettime(m, int32(1), v719)
	mBase = m.M
	v724 = int32(4422768)
	v725 = *(*int64)(unsafe.Add(mBase, _consts[316]))
	v727 = *(*int64)(unsafe.Add(mBase, uint32(v638)+16))
	v728 = int64(*(*int32)(unsafe.Add(mBase, uint32(v719)+8)))
	v729 = *(*int64)(unsafe.Add(mBase, uint32(v719)))
	v733 = *(*int64)(unsafe.Add(mBase, uint32(v638)+24))
	v734 = v728 + v729*int64(1000000000) - v733
	*(*int64)(unsafe.Add(mBase, _consts[316])) = v727 + v734
	v737 = *(*int64)(unsafe.Add(mBase, uint32(v638)+8))
	if v699 != 0 {
		goto L96
	} else {
		goto L97
	}
L94:
	;
	goto L95
L95:
	;
	m.G0 = v719 + int32(16)
	goto L92
L96:
	;
	v739 = *(*int64)(unsafe.Add(mBase, uint32(v721)))
	*(*int64)(unsafe.Add(mBase, uint32(v721))) = v739 + int64(1)
	goto L98
L97:
	;
	goto L98
L98:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v721)+8)) = v737 + v734
	v744 = *(*int64)(unsafe.Add(mBase, uint32(v721)+16))
	*(*int64)(unsafe.Add(mBase, uint32(v721)+16)) = v744 + (v734 - v725 + v727)
	goto L95
L99:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v26)+24)) = v641
	*(*int32)(unsafe.Add(mBase, uint32(v26)+20)) = v642
	*(*int32)(unsafe.Add(mBase, uint32(v26)+28)) = v643
	*(*int32)(unsafe.Add(mBase, uint32(v26)+32)) = v644
	*(*int32)(unsafe.Add(mBase, uint32(v26)+36)) = v645
	*(*int32)(unsafe.Add(mBase, uint32(v26)+40)) = v646
	*(*int32)(unsafe.Add(mBase, uint32(v26)+44)) = v639
	*(*int32)(unsafe.Add(mBase, uint32(v26)+48)) = v638
	*(*int32)(unsafe.Add(mBase, uint32(v26)+52)) = v636
	*(*int32)(unsafe.Add(mBase, uint32(v26)+56)) = v637
	*(*int32)(unsafe.Add(mBase, uint32(v26)+60)) = v635
	F_AtEOXact_GUC(m, int32(1), v644)
	mBase = m.M
	v776 = m.ExcPending
	if v776 != 0 {
		v883 = v651
		goto L6
	} else {
		goto L102
	}
L100:
	;
	goto L101
L101:
	;
	v777 = *(*int32)(unsafe.Add(mBase, uint32(v635)))
	v778 = *(*int32)(unsafe.Add(mBase, uint32(v777)+28))
	if v778 != 0 {
		goto L103
	} else {
		goto L104
	}
L102:
	;
	goto L101
L103:
	;
	v779 = *(*int32)(unsafe.Add(mBase, uint32(v636)))
	v780 = *(*int32)(unsafe.Add(mBase, uint32(v637)))
	*(*int32)(unsafe.Add(mBase, uint32(v26)+24)) = v641
	*(*int32)(unsafe.Add(mBase, uint32(v26)+20)) = v642
	*(*int32)(unsafe.Add(mBase, uint32(v26)+28)) = v643
	*(*int32)(unsafe.Add(mBase, uint32(v26)+32)) = v644
	*(*int32)(unsafe.Add(mBase, uint32(v26)+36)) = v645
	*(*int32)(unsafe.Add(mBase, uint32(v26)+40)) = v646
	*(*int32)(unsafe.Add(mBase, uint32(v26)+44)) = v639
	*(*int32)(unsafe.Add(mBase, uint32(v26)+48)) = v638
	*(*int32)(unsafe.Add(mBase, uint32(v26)+52)) = v636
	*(*int32)(unsafe.Add(mBase, uint32(v26)+56)) = v637
	*(*int32)(unsafe.Add(mBase, uint32(v26)+60)) = v635
	*(*int32)(unsafe.Add(mBase, _consts[240])) = v779
	*(*int32)(unsafe.Add(mBase, _consts[239])) = v780
	goto L106
L104:
	;
	goto L105
L105:
	;
	v799 = *(*int32)(unsafe.Add(mBase, _consts[1216]))
	if v799 != 0 {
		goto L107
	} else {
		goto L108
	}
L106:
	;
	goto L105
L107:
	;
	v800 = *(*int32)(unsafe.Add(mBase, uint32(v635)))
	*(*int32)(unsafe.Add(mBase, uint32(v26)+24)) = v641
	*(*int32)(unsafe.Add(mBase, uint32(v26)+20)) = v642
	*(*int32)(unsafe.Add(mBase, uint32(v26)+28)) = v643
	*(*int32)(unsafe.Add(mBase, uint32(v26)+32)) = v644
	*(*int32)(unsafe.Add(mBase, uint32(v26)+36)) = v645
	*(*int32)(unsafe.Add(mBase, uint32(v26)+40)) = v646
	*(*int32)(unsafe.Add(mBase, uint32(v26)+44)) = v639
	*(*int32)(unsafe.Add(mBase, uint32(v26)+48)) = v638
	*(*int32)(unsafe.Add(mBase, uint32(v26)+52)) = v636
	*(*int32)(unsafe.Add(mBase, uint32(v26)+56)) = v637
	*(*int32)(unsafe.Add(mBase, uint32(v26)+60)) = v635
	m.T0[v799].(func(*base.Module, int32, int32, int32))(m, int32(1), v800, v800+int32(44))
	mBase = m.M
	v817 = m.ExcPending
	if v817 != 0 {
		v883 = v651
		goto L6
	} else {
		goto L110
	}
L108:
	;
	goto L109
L109:
	;
	m.G0 = v26 - int32(-64)
	return v687
L110:
	;
	goto L109
L111:
	;
	v831 = *(*int32)(unsafe.Add(mBase, uint32(v635)))
	*(*int32)(unsafe.Add(mBase, uint32(v26)+24)) = v641
	*(*int32)(unsafe.Add(mBase, uint32(v26)+20)) = v642
	*(*int32)(unsafe.Add(mBase, uint32(v26)+28)) = v643
	*(*int32)(unsafe.Add(mBase, uint32(v26)+32)) = v644
	*(*int32)(unsafe.Add(mBase, uint32(v26)+36)) = v645
	*(*int32)(unsafe.Add(mBase, uint32(v26)+40)) = v646
	*(*int32)(unsafe.Add(mBase, uint32(v26)+44)) = v639
	*(*int32)(unsafe.Add(mBase, uint32(v26)+48)) = v638
	*(*int32)(unsafe.Add(mBase, uint32(v26)+52)) = v636
	*(*int32)(unsafe.Add(mBase, uint32(v26)+56)) = v637
	*(*int32)(unsafe.Add(mBase, uint32(v26)+60)) = v635
	m.T0[v830].(func(*base.Module, int32, int32, int32))(m, int32(2), v831, v831+int32(44))
	mBase = m.M
	v848 = m.ExcPending
	if v848 != 0 {
		v883 = v651
		goto L6
	} else {
		goto L114
	}
L112:
	;
	goto L113
L113:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v26)+24)) = v641
	*(*int32)(unsafe.Add(mBase, uint32(v26)+20)) = v642
	*(*int32)(unsafe.Add(mBase, uint32(v26)+28)) = v643
	*(*int32)(unsafe.Add(mBase, uint32(v26)+32)) = v644
	*(*int32)(unsafe.Add(mBase, uint32(v26)+36)) = v645
	*(*int32)(unsafe.Add(mBase, uint32(v26)+40)) = v646
	*(*int32)(unsafe.Add(mBase, uint32(v26)+44)) = v639
	*(*int32)(unsafe.Add(mBase, uint32(v26)+48)) = v638
	*(*int32)(unsafe.Add(mBase, uint32(v26)+52)) = v636
	*(*int32)(unsafe.Add(mBase, uint32(v26)+56)) = v637
	*(*int32)(unsafe.Add(mBase, uint32(v26)+60)) = v635
	F_pg_re_throw(m)
	mBase = m.M
	v863 = m.ExcPending
	if v863 != 0 {
		v883 = v651
		goto L6
	} else {
		goto L115
	}
L114:
	;
	goto L113
L115:
	;
	goto L5
L116:
	;
	v892 = int32(v888)
	m.G0 = v883
	v894 = *(*int32)(unsafe.Add(mBase, uint32(v892)+4))
	v895 = *(*int32)(unsafe.Add(mBase, uint32(v892)))
	v899 = *(*int32)(unsafe.Add(mBase, uint32(v895)))
	if v24+int32(-48) == v899 {
		goto L119
	} else {
		goto L120
	}
L117:
	;
	m.ExcPending = 1
	goto L125
L118:
	;
	if v902 != 0 {
		goto L122
	} else {
		goto L123
	}
L119:
	;
	v901 = *(*int32)(unsafe.Add(mBase, uint32(v895)+4))
	v902 = v901
	goto L121
L120:
	;
	v902 = int32(0)
	goto L121
L121:
	;
	goto L118
L122:
	;
	v903 = *(*int32)(unsafe.Add(mBase, uint32(v26)+60))
	v904 = *(*int32)(unsafe.Add(mBase, uint32(v26)+56))
	v905 = *(*int32)(unsafe.Add(mBase, uint32(v26)+52))
	v906 = *(*int32)(unsafe.Add(mBase, uint32(v26)+48))
	v907 = *(*int32)(unsafe.Add(mBase, uint32(v26)+44))
	v908 = *(*int32)(unsafe.Add(mBase, uint32(v26)+40))
	v909 = *(*int32)(unsafe.Add(mBase, uint32(v26)+36))
	v910 = *(*int32)(unsafe.Add(mBase, uint32(v26)+32))
	v911 = *(*int32)(unsafe.Add(mBase, uint32(v26)+28))
	v912 = *(*int32)(unsafe.Add(mBase, uint32(v26)+24))
	v913 = *(*int32)(unsafe.Add(mBase, uint32(v26)+20))
	v31 = v894
	v32 = v903
	v33 = v905
	v34 = v904
	v35 = v906
	v36 = v907
	v38 = v912
	v39 = v913
	v40 = v911
	v41 = v910
	v42 = v909
	v43 = v908
	v44 = v902
	v48 = v883
	goto L1
L123:
	;
	goto L124
L124:
	;
	F___wasm_longjmp(m, v895, v894)
	mBase = m.M
	v917 = m.ExcPending
	if v917 != 0 {
		goto L125
	} else {
		goto L126
	}
L125:
	;
	return int32(0)
L126:
	;
	base.Wasm_trap_unreachable()
	for {
	}
}
