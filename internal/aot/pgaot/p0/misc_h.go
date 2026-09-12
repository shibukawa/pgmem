package p0

import (
	base "github.com/shibukawa/pgmem/internal/aot/pgaot/base"
	"unsafe"
)

func F_HaveVirtualXIDsDelayingChkpt(m *base.Module, l0 int32, l1 int32, l2 int32) int32 {
	mBase := m.M
	_ = mBase
	var v13 int32
	_ = v13
	var v15 int32
	_ = v15
	var v19 int32
	_ = v19
	var v22 int32
	_ = v22
	var v23 int32
	_ = v23
	var v24 int32
	_ = v24
	var v30 int32
	_ = v30
	var v35 int32
	_ = v35
	var v45 int32
	_ = v45
	var v48 int32
	_ = v48
	var v49 int32
	_ = v49
	var v53 int32
	_ = v53
	var v58 int32
	_ = v58
	var v63 int32
	_ = v63
	var v73 int32
	_ = v73
	var v74 int32
	_ = v74
	var v76 int32
	_ = v76
	var v80 int32
	_ = v80
	var v94 int32
	_ = v94
	var v108 int32
	_ = v108
	var v110 int32
	_ = v110
	var v114 int32
	_ = v114
	v13 = *(*int32)(unsafe.Add(mBase, _c_F_HaveVirtualXIDsDelayingChkpt[0]))
	v15 = *(*int32)(unsafe.Add(mBase, _c_F_HaveVirtualXIDsDelayingChkpt[1]))
	v19 = F_LWLockAcquire(m, v15+int32(512), int32(1))
	mBase = m.M
	v22 = m.ExcPending
	if v22 != 0 {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	return int32(0)
L2:
	;
	v23 = int32(0)
	v24 = *(*int32)(unsafe.Add(mBase, uint32(v13)))
	if v24 <= v23 {
		v108 = v23
		goto L3
	} else {
		goto L4
	}
L3:
	;
	v110 = *(*int32)(unsafe.Add(mBase, _c_F_HaveVirtualXIDsDelayingChkpt[1]))
	F_LWLockRelease(m, v110+int32(512))
	mBase = m.M
	v114 = m.ExcPending
	if v114 != 0 {
		goto L1
	} else {
		goto L18
	}
L4:
	;
	v30 = *(*int32)(unsafe.Add(mBase, _c_F_HaveVirtualXIDsDelayingChkpt[2]))
	v35 = int32(0)
	goto L5
L5:
	;
	v45 = *(*int32)(unsafe.Add(mBase, uint32(v13+int32(36)+v35<<(uint(int32(2))%32))))
	v48 = v30 + v45*int32(640)
	v49 = *(*int32)(unsafe.Add(mBase, uint32(v48)+120))
	if v49&l2 == int32(0) {
		goto L7
	} else {
		goto L8
	}
L6:
	;
	v108 = int32(0)
	goto L3
L7:
	;
	v94 = v35 + int32(1)
	if v94 != v24 {
		v35 = v94
		goto L5
	} else {
		goto L17
	}
L8:
	;
	v53 = *(*int32)(unsafe.Add(mBase, uint32(v48)+56))
	if v53 == int32(0) {
		goto L7
	} else {
		goto L9
	}
L9:
	;
	if l1 <= int32(0) {
		goto L7
	} else {
		goto L10
	}
L10:
	;
	v58 = *(*int32)(unsafe.Add(mBase, uint32(v48)+52))
	v63 = int32(0)
	goto L11
L11:
	;
	v73 = l0 + v63<<(uint(int32(3))%32)
	v74 = *(*int32)(unsafe.Add(mBase, uint32(v73)))
	if v58 != v74 {
		goto L13
	} else {
		goto L14
	}
L12:
	;
	goto L7
L13:
	;
	v80 = v63 + int32(1)
	if v80 != l1 {
		v63 = v80
		goto L11
	} else {
		goto L16
	}
L14:
	;
	v76 = *(*int32)(unsafe.Add(mBase, uint32(v73)+4))
	if v53 != v76 {
		goto L13
	} else {
		goto L15
	}
L15:
	;
	v108 = int32(1)
	goto L3
L16:
	;
	goto L12
L17:
	;
	goto L6
L18:
	;
	return v108
}
func F_handle_pm_shutdown_request_signal(m *base.Module, l0 int32) {
	mBase := m.M
	_ = mBase
	var v7 int32
	_ = v7
	var v16 int32
	_ = v16
	var v18 int32
	_ = v18
	switch l0 - int32(2) {
	case 0:
		v7 = int32(_a_F_handle_pm_shutdown_request_signal_0)
		*(*int32)(unsafe.Add(mBase, uint32(v7))) = int32(1)
		*(*int32)(unsafe.Add(mBase, _c_F_handle_pm_shutdown_request_signal[0])) = int32(1)
	case 1:
		v7 = int32(_a_F_handle_pm_shutdown_request_signal_1)
		*(*int32)(unsafe.Add(mBase, uint32(v7))) = int32(1)
		*(*int32)(unsafe.Add(mBase, _c_F_handle_pm_shutdown_request_signal[0])) = int32(1)
	default:
	case 13:
		*(*int32)(unsafe.Add(mBase, _c_F_handle_pm_shutdown_request_signal[0])) = int32(1)
	}
	v16 = *(*int32)(unsafe.Add(mBase, _c_F_handle_pm_shutdown_request_signal[1]))
	F_SetLatch(m, v16)
	mBase = m.M
	v18 = m.ExcPending
	if v18 != 0 {
		return
	} else {
		return
	}
}
func F_has_bypassrls_privilege(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v4 int32
	_ = v4
	var v7 int32
	_ = v7
	var v11 int32
	_ = v11
	var v12 int32
	_ = v12
	var v17 int32
	_ = v17
	var v18 int32
	_ = v18
	var v20 int32
	_ = v20
	var v22 int32
	_ = v22
	var v24 int32
	_ = v24
	v4 = F_superuser_arg(m, l0)
	mBase = m.M
	v7 = m.ExcPending
	if v7 != 0 {
		return int32(0)
	} else {
		if v4 == int32(0) {
			v11 = F_SearchSysCache1(m, int32(11), l0)
			mBase = m.M
			v12 = m.ExcPending
			if v12 != 0 {
				return int32(0)
			} else {
				if v11 == int32(0) {
					return int32(0)
				} else {
					v17 = *(*int32)(unsafe.Add(mBase, uint32(v11)+16))
					v18 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v17)+22)))
					v20 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v17+v18)+74)))
					F_ReleaseCatCache(m, v11)
					mBase = m.M
					v22 = m.ExcPending
					if v22 != 0 {
						return int32(0)
					} else {
						v24 = v20
						return v24 & int32(1)
					}
				}
			}
		} else {
			v24 = int32(1)
			return v24 & int32(1)
		}
	}
}
func F_has_superclass(m *base.Module, l0 int32) int32 {
	var v5 int32
	_ = v5
	var v7 int32
	_ = v7
	var v11 int32
	_ = v11
	var v14 int32
	_ = v14
	var v19 int32
	_ = v19
	var v21 int32
	_ = v21
	var v24 int32
	_ = v24
	var v25 int32
	_ = v25
	var v26 int32
	_ = v26
	var v27 int32
	_ = v27
	var v29 int32
	_ = v29
	var v32 int32
	_ = v32
	v5 = m.G0
	v7 = v5 - int32(48)
	m.G0 = v7
	v11 = F_table_open(m, int32(2611), int32(1))
	v14 = m.ExcPending
	if v14 != 0 {
		return int32(0)
	} else {
		F_ScanKeyInit(m, v7, int32(1), int32(3), int32(184), l0)
		v19 = m.ExcPending
		if v19 != 0 {
			return int32(0)
		} else {
			v21 = int32(1)
			v24 = F_systable_beginscan(m, v11, int32(2680), v21, int32(0), v21, v7)
			v25 = m.ExcPending
			if v25 != 0 {
				return int32(0)
			} else {
				v26 = F_systable_getnext(m, v24)
				v27 = m.ExcPending
				if v27 != 0 {
					return int32(0)
				} else {
					F_systable_endscan(m, v24)
					v29 = m.ExcPending
					if v29 != 0 {
						return int32(0)
					} else {
						F_sequence_close(m, v11, int32(1))
						v32 = m.ExcPending
						if v32 != 0 {
							return int32(0)
						} else {
							m.G0 = v7 + int32(48)
							return base.B2i32(v26 != int32(0))
						}
					}
				}
			}
		}
	}
}
func F_hashadjustmembers(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) {
	mBase := m.M
	_ = mBase
	var v9 int32
	_ = v9
	var v10 int32
	_ = v10
	var v11 int32
	_ = v11
	var v13 int32
	_ = v13
	var v14 int32
	_ = v14
	var v15 int32
	_ = v15
	var v18 int32
	_ = v18
	var v19 int32
	_ = v19
	var v23 int32
	_ = v23
	var v24 int32
	_ = v24
	var v26 int32
	_ = v26
	var v29 int32
	_ = v29
	var v33 int32
	_ = v33
	var v34 int32
	_ = v34
	var v37 int32
	_ = v37
	var v40 int32
	_ = v40
	var v41 int32
	_ = v41
	var v45 int32
	_ = v45
	var v46 int32
	_ = v46
	var v47 int32
	_ = v47
	var v48 int32
	_ = v48
	var v50 int32
	_ = v50
	var v53 int32
	_ = v53
	var v58 int32
	_ = v58
	var v60 int32
	_ = v60
	var v61 int32
	_ = v61
	var v64 int32
	_ = v64
	var v65 int32
	_ = v65
	if l1 != 0 {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	F_CommandCounterIncrement(m)
	mBase = m.M
	v9 = m.ExcPending
	if v9 != 0 {
		goto L4
	} else {
		goto L5
	}
L2:
	;
	v13 = int32(0)
	goto L3
L3:
	;
	v14 = F_list_concat_copy(m, l2, l3)
	mBase = m.M
	v15 = m.ExcPending
	if v15 != 0 {
		goto L4
	} else {
		goto L8
	}
L4:
	;
	return
L5:
	;
	v10 = F_get_opclass_input_type(m, l1)
	mBase = m.M
	v11 = m.ExcPending
	if v11 != 0 {
		goto L4
	} else {
		goto L6
	}
L6:
	;
	v13 = v10
	goto L3
L7:
	;
	return
L8:
	;
	if v14 == int32(0) {
		goto L7
	} else {
		goto L9
	}
L9:
	;
	v18 = int32(0)
	v19 = *(*int32)(unsafe.Add(mBase, uint32(v14)+4))
	if v19 <= v18 {
		goto L7
	} else {
		goto L10
	}
L10:
	;
	v23 = l1
	v24 = v18
	v26 = v13
	goto L11
L11:
	;
	v29 = *(*int32)(unsafe.Add(mBase, uint32(v14)+12))
	v33 = *(*int32)(unsafe.Add(mBase, uint32(v29+v24<<(uint(int32(2))%32))))
	v34 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v33))))
	if v34 == int32(1) {
		goto L15
	} else {
		goto L16
	}
L12:
	;
	goto L7
L13:
	;
	v64 = v24 + int32(1)
	v65 = *(*int32)(unsafe.Add(mBase, uint32(v14)+4))
	if v64 < v65 {
		v23 = v60
		v24 = v64
		v26 = v61
		goto L11
	} else {
		goto L27
	}
L14:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v33)+28)) = l0
	v58 = int32(256)
	*(*uint16)(unsafe.Add(mBase, uint32(v33)+24)) = uint16(v58)
	v60 = v23
	v61 = v26
	goto L13
L15:
	;
	v37 = *(*int32)(unsafe.Add(mBase, uint32(v33)+8))
	if v37 != int32(1) {
		goto L14
	} else {
		goto L18
	}
L16:
	;
	goto L17
L17:
	;
	v40 = *(*int32)(unsafe.Add(mBase, uint32(v33)+12))
	v41 = *(*int32)(unsafe.Add(mBase, uint32(v33)+16))
	if v40 != v41 {
		goto L14
	} else {
		goto L19
	}
L18:
	;
	goto L17
L19:
	;
	if v26 != v40 {
		goto L20
	} else {
		goto L21
	}
L20:
	;
	v45 = F_opclass_for_family_datatype(m, int32(405), l0, v40)
	mBase = m.M
	v46 = m.ExcPending
	if v46 != 0 {
		goto L4
	} else {
		goto L23
	}
L21:
	;
	v47 = v23
	v48 = v26
	goto L22
L22:
	;
	if v47 != 0 {
		goto L24
	} else {
		goto L25
	}
L23:
	;
	v47 = v45
	v48 = v40
	goto L22
L24:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v33)+28)) = v47
	v50 = int32(1)
	*(*uint16)(unsafe.Add(mBase, uint32(v33)+24)) = uint16(v50)
	v60 = v47
	v61 = v48
	goto L13
L25:
	;
	goto L26
L26:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v33)+28)) = l0
	v53 = int32(256)
	*(*uint16)(unsafe.Add(mBase, uint32(v33)+24)) = uint16(v53)
	v60 = int32(0)
	v61 = v48
	goto L13
L27:
	;
	goto L12
}
func F_hashbool(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v2 int32
	_ = v2
	var v9 int32
	_ = v9
	var v12 int32
	_ = v12
	var v17 int32
	_ = v17
	var v22 int32
	_ = v22
	var v26 int32
	_ = v26
	var v30 int32
	_ = v30
	v2 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v9 = int32(711645284)
	v12 = base.B2i32(v2 != int32(0)) - int32(1636608428) ^ v9 - int32(1455628627)
	v17 = v12 ^ int32(-1636608428) - base.I32_rotl(v12, int32(25))
	v22 = v17 ^ v9 - base.I32_rotl(v17, int32(16))
	v26 = v22 ^ v12 - base.I32_rotl(v22, int32(4))
	v30 = v26 ^ v17 - base.I32_rotl(v26, int32(14))
	return v30 ^ v22 - base.I32_rotl(v30, int32(24))
}
func F_hashbucketcleanup(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32, l5 int32, l6 int32, l7 int32, l8 int32, l9 int32, l10 int32, l11 int32, l12 int32) {
	mBase := m.M
	_ = mBase
	var v14 int32
	_ = v14
	var v35 int32
	_ = v35
	var v37 int32
	_ = v37
	var v54 int32
	_ = v54
	var v60 int32
	_ = v60
	var v62 int32
	_ = v62
	var v75 int32
	_ = v75
	var v76 int32
	_ = v76
	var v77 int32
	_ = v77
	var v81 int32
	_ = v81
	var v87 int32
	_ = v87
	var v89 int32
	_ = v89
	var v95 int32
	_ = v95
	var v96 int32
	_ = v96
	var v97 int32
	_ = v97
	var v98 int32
	_ = v98
	var v102 int32
	_ = v102
	var v106 int32
	_ = v106
	var v129 int32
	_ = v129
	var v133 int32
	_ = v133
	var v154 int32
	_ = v154
	var v157 int32
	_ = v157
	var v160 int32
	_ = v160
	var v161 int32
	_ = v161
	var v166 float64
	_ = v166
	var v174 int32
	_ = v174
	var v177 int32
	_ = v177
	var v179 int32
	_ = v179
	var v181 int32
	_ = v181
	var v183 int32
	_ = v183
	var v188 int32
	_ = v188
	var v196 float64
	_ = v196
	var v200 int32
	_ = v200
	var v204 int32
	_ = v204
	var v207 int32
	_ = v207
	var v208 int32
	_ = v208
	var v210 int32
	_ = v210
	var v217 int32
	_ = v217
	var v220 float64
	_ = v220
	var v225 int32
	_ = v225
	var v231 int32
	_ = v231
	var v234 int32
	_ = v234
	var v237 int32
	_ = v237
	var v238 int32
	_ = v238
	var v239 int32
	_ = v239
	var v243 int32
	_ = v243
	var v246 int32
	_ = v246
	var v247 int32
	_ = v247
	var v252 int32
	_ = v252
	var v257 int32
	_ = v257
	var v258 int32
	_ = v258
	var v264 int32
	_ = v264
	var v268 int32
	_ = v268
	var v269 int32
	_ = v269
	var v275 int32
	_ = v275
	var v278 int64
	_ = v278
	var v279 int32
	_ = v279
	var v283 int32
	_ = v283
	var v289 int32
	_ = v289
	var v291 int32
	_ = v291
	var v297 int32
	_ = v297
	var v302 int32
	_ = v302
	var v304 int32
	_ = v304
	var v305 int32
	_ = v305
	var v322 int32
	_ = v322
	var v332 int32
	_ = v332
	var v346 int32
	_ = v346
	var v347 int32
	_ = v347
	var v351 int32
	_ = v351
	var v353 int32
	_ = v353
	var v356 int32
	_ = v356
	var v359 int32
	_ = v359
	var v363 int32
	_ = v363
	var v369 int32
	_ = v369
	var v371 int32
	_ = v371
	var v377 int32
	_ = v377
	var v378 int32
	_ = v378
	var v379 int32
	_ = v379
	var v381 int32
	_ = v381
	var v385 int32
	_ = v385
	var v386 int32
	_ = v386
	var v388 int32
	_ = v388
	var v391 int32
	_ = v391
	var v392 int32
	_ = v392
	var v393 int32
	_ = v393
	var v397 int32
	_ = v397
	var v400 int32
	_ = v400
	var v401 int32
	_ = v401
	var v403 int32
	_ = v403
	var v407 int32
	_ = v407
	var v410 int64
	_ = v410
	var v411 int32
	_ = v411
	var v415 int32
	_ = v415
	var v417 int32
	_ = v417
	var v425 int32
	_ = v425
	var v426 int32
	_ = v426
	var v429 int32
	_ = v429
	var v431 int32
	_ = v431
	var v436 int32
	_ = v436
	var v442 int32
	_ = v442
	var v444 int32
	_ = v444
	var v450 int32
	_ = v450
	var v451 int32
	_ = v451
	var v452 int32
	_ = v452
	var v453 int32
	_ = v453
	var v465 int32
	_ = v465
	var v473 int32
	_ = v473
	var v491 int32
	_ = v491
	var v493 int32
	_ = v493
	var v494 int32
	_ = v494
	var v498 int32
	_ = v498
	var v504 int32
	_ = v504
	var v506 int32
	_ = v506
	var v512 int32
	_ = v512
	var v513 int32
	_ = v513
	var v514 int32
	_ = v514
	var v515 int32
	_ = v515
	var v518 int32
	_ = v518
	var v519 int32
	_ = v519
	var v520 int32
	_ = v520
	var v521 int32
	_ = v521
	var v523 int32
	_ = v523
	var v524 int32
	_ = v524
	var v526 int32
	_ = v526
	var v531 int32
	_ = v531
	var v535 int32
	_ = v535
	var v536 int32
	_ = v536
	var v537 int32
	_ = v537
	var v542 int32
	_ = v542
	var v543 int32
	_ = v543
	var v544 int32
	_ = v544
	var v552 int32
	_ = v552
	var v559 int32
	_ = v559
	var v573 int32
	_ = v573
	var v574 int32
	_ = v574
	var v582 int32
	_ = v582
	var v598 int32
	_ = v598
	var v599 int32
	_ = v599
	var v608 int32
	_ = v608
	var v613 int32
	_ = v613
	var v614 int32
	_ = v614
	var v622 int32
	_ = v622
	var v633 int32
	_ = v633
	var v635 int32
	_ = v635
	var v636 int32
	_ = v636
	var v638 int32
	_ = v638
	var v639 int32
	_ = v639
	var v654 int32
	_ = v654
	var v655 int32
	_ = v655
	var v661 int32
	_ = v661
	var v662 int32
	_ = v662
	var v668 int32
	_ = v668
	var v669 int32
	_ = v669
	var v670 int32
	_ = v670
	var v678 int32
	_ = v678
	var v689 int32
	_ = v689
	var v691 int32
	_ = v691
	var v694 int32
	_ = v694
	var v695 int32
	_ = v695
	var v704 int32
	_ = v704
	var v707 int32
	_ = v707
	var v708 int32
	_ = v708
	var v709 int32
	_ = v709
	var v711 int32
	_ = v711
	var v715 int32
	_ = v715
	var v716 int32
	_ = v716
	var v719 int32
	_ = v719
	var v722 int32
	_ = v722
	var v723 int32
	_ = v723
	var v724 int32
	_ = v724
	var v725 int32
	_ = v725
	var v726 int32
	_ = v726
	var v730 int32
	_ = v730
	var v733 int32
	_ = v733
	var v734 int32
	_ = v734
	var v739 int32
	_ = v739
	var v740 int32
	_ = v740
	var v742 int32
	_ = v742
	var v751 int32
	_ = v751
	var v753 int32
	_ = v753
	var v757 int32
	_ = v757
	var v759 int32
	_ = v759
	var v761 int32
	_ = v761
	var v762 int32
	_ = v762
	var v763 int32
	_ = v763
	var v767 int32
	_ = v767
	var v770 int32
	_ = v770
	var v771 int32
	_ = v771
	var v776 int32
	_ = v776
	var v781 int32
	_ = v781
	var v782 int32
	_ = v782
	var v788 int32
	_ = v788
	var v792 int32
	_ = v792
	var v793 int32
	_ = v793
	var v799 int32
	_ = v799
	var v800 int32
	_ = v800
	var v803 int32
	_ = v803
	var v814 int32
	_ = v814
	var v841 int32
	_ = v841
	var v845 int32
	_ = v845
	var v849 int32
	_ = v849
	var v851 int32
	_ = v851
	var v853 int32
	_ = v853
	var v858 int32
	_ = v858
	var v865 int32
	_ = v865
	var v868 int64
	_ = v868
	var v869 int32
	_ = v869
	var v873 int32
	_ = v873
	var v879 int32
	_ = v879
	var v881 int32
	_ = v881
	var v887 int32
	_ = v887
	var v894 int32
	_ = v894
	var v896 int32
	_ = v896
	var v898 int32
	_ = v898
	var v902 int32
	_ = v902
	var v905 int64
	_ = v905
	var v941 int32
	_ = v941
	var v943 int32
	_ = v943
	var v984 int32
	_ = v984
	var v986 int32
	_ = v986
	var v989 int32
	_ = v989
	var v993 int32
	_ = v993
	var v999 int32
	_ = v999
	var v1001 int32
	_ = v1001
	var v1007 int32
	_ = v1007
	var v1008 int32
	_ = v1008
	var v1009 int32
	_ = v1009
	var v1010 int32
	_ = v1010
	var v1024 int32
	_ = v1024
	var v1054 int32
	_ = v1054
	var v1056 int32
	_ = v1056
	var v1058 int32
	_ = v1058
	var v1060 int32
	_ = v1060
	var v1068 int32
	_ = v1068
	var v1079 int32
	_ = v1079
	var v1080 int32
	_ = v1080
	var v1082 int32
	_ = v1082
	var v1091 int32
	_ = v1091
	var v1095 int32
	_ = v1095
	var v1096 int32
	_ = v1096
	var v1104 int32
	_ = v1104
	var v1115 int32
	_ = v1115
	var v1117 int32
	_ = v1117
	var v1120 int32
	_ = v1120
	var v1121 int32
	_ = v1121
	var v1130 int32
	_ = v1130
	var v1134 int32
	_ = v1134
	var v1135 int32
	_ = v1135
	var v1154 int32
	_ = v1154
	var v1159 int32
	_ = v1159
	var v1160 int32
	_ = v1160
	var v1168 int32
	_ = v1168
	var v1170 int32
	_ = v1170
	var v1172 int32
	_ = v1172
	var v1176 int32
	_ = v1176
	var v1178 int32
	_ = v1178
	var v1180 int32
	_ = v1180
	var v1184 int32
	_ = v1184
	var v1188 int32
	_ = v1188
	var v1194 int32
	_ = v1194
	var v1196 int32
	_ = v1196
	var v1202 int32
	_ = v1202
	var v1203 int32
	_ = v1203
	var v1207 int32
	_ = v1207
	var v1213 int32
	_ = v1213
	var v1215 int32
	_ = v1215
	var v1221 int32
	_ = v1221
	var v1222 int32
	_ = v1222
	var v1223 int32
	_ = v1223
	var v1224 int32
	_ = v1224
	var v1225 int32
	_ = v1225
	var v1229 int32
	_ = v1229
	var v1235 int32
	_ = v1235
	var v1237 int32
	_ = v1237
	var v1243 int32
	_ = v1243
	var v1244 int32
	_ = v1244
	var v1250 int32
	_ = v1250
	var v1251 int32
	_ = v1251
	var v1252 int32
	_ = v1252
	var v1256 int32
	_ = v1256
	var v1257 int32
	_ = v1257
	var v1258 int32
	_ = v1258
	var v1262 int32
	_ = v1262
	var v1263 int32
	_ = v1263
	var v1267 int32
	_ = v1267
	var v1273 int32
	_ = v1273
	var v1275 int32
	_ = v1275
	var v1281 int32
	_ = v1281
	var v1284 int32
	_ = v1284
	var v1286 int32
	_ = v1286
	var v1288 int32
	_ = v1288
	var v1303 int32
	_ = v1303
	var v1335 int32
	_ = v1335
	var v1336 int32
	_ = v1336
	var v1338 int32
	_ = v1338
	var v1349 int32
	_ = v1349
	var v1358 int32
	_ = v1358
	var v1359 int32
	_ = v1359
	var v1361 int32
	_ = v1361
	var v1372 int32
	_ = v1372
	var v1373 int32
	_ = v1373
	var v1376 int32
	_ = v1376
	var v1379 int32
	_ = v1379
	var v1381 int32
	_ = v1381
	var v1384 int32
	_ = v1384
	var v1423 int32
	_ = v1423
	var v1426 int32
	_ = v1426
	var v1430 int32
	_ = v1430
	var v1435 int32
	_ = v1435
	var v1439 int32
	_ = v1439
	var v1440 int32
	_ = v1440
	var v1441 int32
	_ = v1441
	var v1442 int32
	_ = v1442
	var v1449 int32
	_ = v1449
	var v1450 int32
	_ = v1450
	var v1455 int32
	_ = v1455
	var v1458 int32
	_ = v1458
	var v1461 int32
	_ = v1461
	var v1462 int32
	_ = v1462
	var v1466 int32
	_ = v1466
	var v1472 int32
	_ = v1472
	var v1474 int32
	_ = v1474
	var v1480 int32
	_ = v1480
	var v1483 int32
	_ = v1483
	var v1484 int32
	_ = v1484
	var v1485 int32
	_ = v1485
	var v1489 int32
	_ = v1489
	var v1492 int32
	_ = v1492
	var v1493 int32
	_ = v1493
	var v1498 int32
	_ = v1498
	var v1499 int32
	_ = v1499
	var v1501 int32
	_ = v1501
	var v1506 int32
	_ = v1506
	var v1508 int32
	_ = v1508
	var v1512 int32
	_ = v1512
	var v1513 int32
	_ = v1513
	var v1519 int32
	_ = v1519
	var v1523 int32
	_ = v1523
	var v1529 int32
	_ = v1529
	var v1531 int32
	_ = v1531
	var v1537 int32
	_ = v1537
	var v1538 int32
	_ = v1538
	var v1542 int32
	_ = v1542
	var v1549 int32
	_ = v1549
	var v1555 int32
	_ = v1555
	var v1557 int32
	_ = v1557
	var v1563 int32
	_ = v1563
	var v1564 int32
	_ = v1564
	var v1568 int32
	_ = v1568
	var v1570 int32
	_ = v1570
	var v1572 int32
	_ = v1572
	var v1575 int32
	_ = v1575
	var v1576 int32
	_ = v1576
	var v1582 int32
	_ = v1582
	var v1584 int32
	_ = v1584
	var v1585 int32
	_ = v1585
	var v1586 int32
	_ = v1586
	var v1591 int32
	_ = v1591
	var v1592 int32
	_ = v1592
	var v1593 int32
	_ = v1593
	var v1597 int32
	_ = v1597
	var v1600 int32
	_ = v1600
	var v1601 int32
	_ = v1601
	var v1610 int32
	_ = v1610
	var v1615 int32
	_ = v1615
	var v1616 int32
	_ = v1616
	var v1622 int32
	_ = v1622
	var v1623 int32
	_ = v1623
	var v1624 int32
	_ = v1624
	var v1628 int32
	_ = v1628
	var v1629 int32
	_ = v1629
	var v1633 int32
	_ = v1633
	var v1644 int32
	_ = v1644
	var v1673 int32
	_ = v1673
	var v1675 int32
	_ = v1675
	var v1677 int32
	_ = v1677
	var v1679 int32
	_ = v1679
	var v1681 int32
	_ = v1681
	var v1683 int32
	_ = v1683
	var v1685 int32
	_ = v1685
	var v1688 int32
	_ = v1688
	var v1697 int32
	_ = v1697
	var v1698 int32
	_ = v1698
	var v1700 int32
	_ = v1700
	var v1702 int32
	_ = v1702
	var v1712 int32
	_ = v1712
	var v1740 int32
	_ = v1740
	var v1743 int32
	_ = v1743
	var v1747 int32
	_ = v1747
	var v1751 int32
	_ = v1751
	var v1755 int32
	_ = v1755
	var v1761 int32
	_ = v1761
	var v1767 int32
	_ = v1767
	var v1771 int32
	_ = v1771
	var v1774 int64
	_ = v1774
	var v1775 int32
	_ = v1775
	var v1781 int32
	_ = v1781
	var v1787 int32
	_ = v1787
	var v1789 int32
	_ = v1789
	var v1795 int32
	_ = v1795
	var v1802 int32
	_ = v1802
	var v1808 int32
	_ = v1808
	var v1810 int32
	_ = v1810
	var v1816 int32
	_ = v1816
	var v1817 int64
	_ = v1817
	var v1822 int32
	_ = v1822
	var v1823 int32
	_ = v1823
	var v1826 int32
	_ = v1826
	var v1830 int32
	_ = v1830
	var v1836 int32
	_ = v1836
	var v1838 int32
	_ = v1838
	var v1844 int32
	_ = v1844
	var v1851 int32
	_ = v1851
	var v1857 int32
	_ = v1857
	var v1859 int32
	_ = v1859
	var v1865 int32
	_ = v1865
	var v1872 int32
	_ = v1872
	var v1878 int32
	_ = v1878
	var v1880 int32
	_ = v1880
	var v1886 int32
	_ = v1886
	var v1892 int32
	_ = v1892
	var v1898 int32
	_ = v1898
	var v1900 int32
	_ = v1900
	var v1906 int32
	_ = v1906
	var v1943 int32
	_ = v1943
	var v1945 int32
	_ = v1945
	var v1953 int32
	_ = v1953
	var v1955 int32
	_ = v1955
	var v1957 int32
	_ = v1957
	var v1959 int32
	_ = v1959
	var v1961 int32
	_ = v1961
	var v1968 int32
	_ = v1968
	var v1972 int32
	_ = v1972
	var v1977 int32
	_ = v1977
	var v1988 int32
	_ = v1988
	var v2018 int32
	_ = v2018
	var v2020 int32
	_ = v2020
	var v2022 int32
	_ = v2022
	var v2062 int32
	_ = v2062
	var v2064 int32
	_ = v2064
	var v2065 int32
	_ = v2065
	var v2069 int32
	_ = v2069
	var v2075 int32
	_ = v2075
	var v2077 int32
	_ = v2077
	var v2083 int32
	_ = v2083
	var v2084 int32
	_ = v2084
	var v2087 int32
	_ = v2087
	var v2090 int32
	_ = v2090
	var v2110 int32
	_ = v2110
	var v2115 int32
	_ = v2115
	var v2130 int32
	_ = v2130
	var v2155 int32
	_ = v2155
	v14 = int32(0)
	v35 = m.G0
	v37 = v35 - int32(_a_F_hashbucketcleanup_0)
	m.G0 = v37
	v54 = l2
	v60 = l3
	v62 = v14
	goto L1
L1:
	;
	F_vacuum_delay_point(m, int32(0))
	mBase = m.M
	v75 = m.ExcPending
	if v75 != 0 {
		goto L3
	} else {
		goto L4
	}
L3:
	;
	return
L4:
	;
	v76 = int32(0)
	v77 = base.B2i32(v76 <= v54)
	if v77 == v76 {
		goto L6
	} else {
		goto L7
	}
L5:
	;
	v96 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v95)+16)))
	v97 = v96 + v95
	v98 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v95)+12)))
	if base.Ui32(int32(25)) <= base.Ui32(v98) {
		goto L11
	} else {
		goto L12
	}
L6:
	;
	v81 = *(*int32)(unsafe.Add(mBase, _c_F_hashbucketcleanup[0]))
	v87 = *(*int32)(unsafe.Add(mBase, uint32(v81+(v54^int32(-1))<<(uint(int32(2))%32))))
	v95 = v87
	goto L5
L7:
	;
	goto L8
L8:
	;
	v89 = *(*int32)(unsafe.Add(mBase, _c_F_hashbucketcleanup[1]))
	v95 = v89 + v54<<(uint(int32(13))%32) + int32(-8192)
	goto L5
L9:
	;
	if v322 != int32(-1) {
		goto L65
	} else {
		goto L66
	}
L10:
	;
	v129 = int32(1)
	v133 = int32(0)
	goto L15
L11:
	;
	v102 = v98 + int32(_a_F_hashbucketcleanup_1)
	if v102&int32(_a_F_hashbucketcleanup_2) != 0 {
		goto L10
	} else {
		goto L14
	}
L12:
	;
	goto L13
L13:
	;
	v106 = *(*int32)(unsafe.Add(mBase, uint32(v97)+4))
	v322 = v106
	v332 = v62
	goto L9
L14:
	;
	goto L13
L15:
	;
	v154 = *(*int32)(unsafe.Add(mBase, uint32(v129<<(uint(int32(2))%32)+(v95+int32(24))-int32(4))))
	v157 = v95 + v154&int32(_a_F_hashbucketcleanup_3)
	if l11 == int32(0) {
		goto L20
	} else {
		goto L21
	}
L16:
	;
	v204 = *(*int32)(unsafe.Add(mBase, uint32(v97)+4))
	if v200 <= int32(0) {
		v322 = v204
		v332 = v62
		goto L9
	} else {
		goto L37
	}
L17:
	;
	if v129 != int32(base.Ui32(v102)>>(uint(int32(2))%32))&int32(_a_F_hashbucketcleanup_4) {
		v129 = v129 + int32(1)
		v133 = v200
		goto L15
	} else {
		goto L36
	}
L18:
	;
	if l9 == int32(0) {
		v200 = v133
		goto L17
	} else {
		goto L35
	}
L19:
	;
	v188 = int32(1)
	*(*uint16)(unsafe.Add(mBase, uint32(v37+int32(16)+v133<<(uint(v188)%32)))) = uint16(v129)
	v200 = v133 + v188
	goto L17
L20:
	;
	if l10 == int32(0) {
		goto L18
	} else {
		goto L25
	}
L21:
	;
	v160 = m.T0[l11].(func(*base.Module, int32, int32) int32)(m, v157, l12)
	mBase = m.M
	v161 = m.ExcPending
	if v161 != 0 {
		goto L3
	} else {
		goto L22
	}
L22:
	;
	if v160 == int32(0) {
		goto L20
	} else {
		goto L23
	}
L23:
	;
	if l8 == int32(0) {
		goto L19
	} else {
		goto L24
	}
L24:
	;
	v166 = *(*float64)(unsafe.Add(mBase, uint32(l8)))
	*(*float64)(unsafe.Add(mBase, uint32(l8))) = base.F64_add(v166, float64(1))
	goto L19
L25:
	;
	v174 = int32(*(*int16)(unsafe.Add(mBase, uint32(v157)+6)))
	if int32(0) <= v174 {
		goto L27
	} else {
		goto L28
	}
L26:
	;
	v181 = v179 & l6
	if base.Ui32(v181) <= base.Ui32(l5) {
		goto L31
	} else {
		goto L32
	}
L27:
	;
	v177 = int32(8)
	goto L29
L28:
	;
	v177 = int32(16)
	goto L29
L29:
	;
	v179 = *(*int32)(unsafe.Add(mBase, uint32(v157+v177)))
	goto L26
L30:
	;
	if v183&v181 == l1 {
		goto L18
	} else {
		goto L34
	}
L31:
	;
	v183 = int32(-1)
	goto L33
L32:
	;
	v183 = l7
	goto L33
L33:
	;
	goto L30
L34:
	;
	goto L19
L35:
	;
	v196 = *(*float64)(unsafe.Add(mBase, uint32(l9)))
	*(*float64)(unsafe.Add(mBase, uint32(l9))) = base.F64_add(v196, float64(1))
	v200 = v133
	goto L17
L36:
	;
	goto L16
L37:
	;
	v207 = int32(0)
	v208 = int32(_a_F_hashbucketcleanup_5)
	v210 = *(*int32)(unsafe.Add(mBase, _c_F_hashbucketcleanup[2]))
	*(*int32)(unsafe.Add(mBase, _c_F_hashbucketcleanup[2])) = v210 + int32(1)
	F_PageIndexMultiDelete(m, v95, v37+int32(16), v200)
	mBase = m.M
	v217 = m.ExcPending
	if v217 != 0 {
		goto L3
	} else {
		goto L38
	}
L38:
	;
	if l8 == int32(0) {
		v234 = v207
		goto L39
	} else {
		goto L40
	}
L39:
	;
	F_MarkBufferDirty(m, v54)
	mBase = m.M
	v237 = m.ExcPending
	if v237 != 0 {
		goto L3
	} else {
		goto L43
	}
L40:
	;
	v220 = *(*float64)(unsafe.Add(mBase, uint32(l8)))
	if base.F64_gt(v220, float64(0)) == int32(0) {
		v234 = v207
		goto L39
	} else {
		goto L41
	}
L41:
	;
	v225 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v97)+12)))
	if v225&int32(128) == int32(0) {
		v234 = v207
		goto L39
	} else {
		goto L42
	}
L42:
	;
	v231 = v225 & int32(_a_F_hashbucketcleanup_6)
	*(*uint16)(unsafe.Add(mBase, uint32(v97)+12)) = uint16(v231)
	v234 = int32(1)
	goto L39
L43:
	;
	v238 = *(*int32)(unsafe.Add(mBase, uint32(l0)+48))
	v239 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v238)+118)))
	if v239 != int32(112) {
		goto L44
	} else {
		goto L45
	}
L44:
	;
	v302 = int32(_a_F_hashbucketcleanup_5)
	v304 = *(*int32)(unsafe.Add(mBase, _c_F_hashbucketcleanup[2]))
	v305 = int32(1)
	*(*int32)(unsafe.Add(mBase, _c_F_hashbucketcleanup[2])) = v304 - v305
	v322 = v204
	v332 = v305
	goto L9
L45:
	;
	v243 = *(*int32)(unsafe.Add(mBase, _c_F_hashbucketcleanup[3]))
	if v243 <= int32(0) {
		goto L46
	} else {
		goto L47
	}
L46:
	;
	v246 = *(*int32)(unsafe.Add(mBase, uint32(l0)+32))
	if v246 != 0 {
		goto L44
	} else {
		goto L49
	}
L47:
	;
	goto L48
L48:
	;
	*(*uint8)(unsafe.Add(mBase, uint32(v37)+14)) = uint8(v234)
	*(*uint8)(unsafe.Add(mBase, uint32(v37)+15)) = uint8(base.B2i32(l2 == v54))
	F_XLogBeginInsert(m)
	mBase = m.M
	v252 = m.ExcPending
	if v252 != 0 {
		goto L3
	} else {
		goto L51
	}
L49:
	;
	v247 = *(*int32)(unsafe.Add(mBase, uint32(l0)+40))
	if v247 != 0 {
		goto L44
	} else {
		goto L50
	}
L50:
	;
	goto L48
L51:
	;
	F_XLogRegisterData(m, v37+int32(14), int32(2))
	mBase = m.M
	v257 = m.ExcPending
	if v257 != 0 {
		goto L3
	} else {
		goto L52
	}
L52:
	;
	v258 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v37)+15)))
	if v258 == int32(0) {
		goto L53
	} else {
		goto L54
	}
L53:
	;
	F_XLogRegisterBuffer(m, int32(0), l2, int32(42))
	mBase = m.M
	v264 = m.ExcPending
	if v264 != 0 {
		goto L3
	} else {
		goto L56
	}
L54:
	;
	goto L55
L55:
	;
	F_XLogRegisterBuffer(m, int32(1), v54, int32(8))
	mBase = m.M
	v268 = m.ExcPending
	if v268 != 0 {
		goto L3
	} else {
		goto L57
	}
L56:
	;
	goto L55
L57:
	;
	v269 = int32(1)
	F_XLogRegisterBufData(m, v269, v37+int32(16), v200<<(uint(v269)%32))
	mBase = m.M
	v275 = m.ExcPending
	if v275 != 0 {
		goto L3
	} else {
		goto L58
	}
L58:
	;
	v278 = F_XLogInsert(m, int32(12), int32(144))
	mBase = m.M
	v279 = m.ExcPending
	if v279 != 0 {
		goto L3
	} else {
		goto L59
	}
L59:
	;
	if v77 == int32(0) {
		goto L61
	} else {
		goto L62
	}
L60:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v297))) = base.I64_rotr(v278, int64(32))
	goto L44
L61:
	;
	v283 = *(*int32)(unsafe.Add(mBase, _c_F_hashbucketcleanup[0]))
	v289 = *(*int32)(unsafe.Add(mBase, uint32(v283+(v54^int32(-1))<<(uint(int32(2))%32))))
	v297 = v289
	goto L60
L62:
	;
	goto L63
L63:
	;
	v291 = *(*int32)(unsafe.Add(mBase, _c_F_hashbucketcleanup[1]))
	v297 = v291 + v54<<(uint(int32(13))%32) + int32(-8192)
	goto L60
L64:
	;
	v54 = v346
	v60 = v322
	v62 = v332
	goto L1
L65:
	;
	v346 = F__hash_getbuf_with_strategy(m, l0, v322, int32(1), l4)
	mBase = m.M
	v347 = m.ExcPending
	if v347 != 0 {
		goto L3
	} else {
		goto L68
	}
L66:
	;
	goto L67
L67:
	;
	if l2 != v54 {
		goto L74
	} else {
		goto L75
	}
L68:
	;
	if l3 == v60 {
		goto L69
	} else {
		goto L70
	}
L69:
	;
	F_LockBuffer(m, v54, int32(0))
	mBase = m.M
	v351 = m.ExcPending
	if v351 != 0 {
		goto L3
	} else {
		goto L72
	}
L70:
	;
	goto L71
L71:
	;
	F_UnlockReleaseBuffer(m, v54)
	mBase = m.M
	v353 = m.ExcPending
	if v353 != 0 {
		goto L3
	} else {
		goto L73
	}
L72:
	;
	goto L64
L73:
	;
	goto L64
L74:
	;
	F_UnlockReleaseBuffer(m, v54)
	mBase = m.M
	v356 = m.ExcPending
	if v356 != 0 {
		goto L3
	} else {
		goto L77
	}
L75:
	;
	goto L76
L76:
	;
	if l10 != 0 {
		goto L79
	} else {
		goto L80
	}
L77:
	;
	F_LockBuffer(m, l2, int32(2))
	mBase = m.M
	v359 = m.ExcPending
	if v359 != 0 {
		goto L3
	} else {
		goto L78
	}
L78:
	;
	goto L76
L79:
	;
	if l2 < int32(0) {
		goto L83
	} else {
		goto L84
	}
L80:
	;
	goto L81
L81:
	;
	if v332 == int32(0) {
		goto L98
	} else {
		goto L99
	}
L82:
	;
	v378 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v377)+16)))
	v379 = int32(_a_F_hashbucketcleanup_5)
	v381 = *(*int32)(unsafe.Add(mBase, _c_F_hashbucketcleanup[2]))
	*(*int32)(unsafe.Add(mBase, _c_F_hashbucketcleanup[2])) = v381 + int32(1)
	v385 = v377 + v378
	v386 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v385)+12)))
	v388 = v386 & int32(_a_F_hashbucketcleanup_7)
	*(*uint16)(unsafe.Add(mBase, uint32(v385)+12)) = uint16(v388)
	F_MarkBufferDirty(m, l2)
	mBase = m.M
	v391 = m.ExcPending
	if v391 != 0 {
		goto L3
	} else {
		goto L86
	}
L83:
	;
	v363 = *(*int32)(unsafe.Add(mBase, _c_F_hashbucketcleanup[0]))
	v369 = *(*int32)(unsafe.Add(mBase, uint32(v363+(l2^int32(-1))<<(uint(int32(2))%32))))
	v377 = v369
	goto L82
L84:
	;
	goto L85
L85:
	;
	v371 = *(*int32)(unsafe.Add(mBase, _c_F_hashbucketcleanup[1]))
	v377 = v371 + l2<<(uint(int32(13))%32) + int32(-8192)
	goto L82
L86:
	;
	v392 = *(*int32)(unsafe.Add(mBase, uint32(l0)+48))
	v393 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v392)+118)))
	if v393 != int32(112) {
		goto L87
	} else {
		goto L88
	}
L87:
	;
	v415 = int32(_a_F_hashbucketcleanup_5)
	v417 = *(*int32)(unsafe.Add(mBase, _c_F_hashbucketcleanup[2]))
	*(*int32)(unsafe.Add(mBase, _c_F_hashbucketcleanup[2])) = v417 - int32(1)
	goto L81
L88:
	;
	v397 = *(*int32)(unsafe.Add(mBase, _c_F_hashbucketcleanup[3]))
	if v397 <= int32(0) {
		goto L89
	} else {
		goto L90
	}
L89:
	;
	v400 = *(*int32)(unsafe.Add(mBase, uint32(l0)+32))
	if v400 != 0 {
		goto L87
	} else {
		goto L92
	}
L90:
	;
	goto L91
L91:
	;
	F_XLogBeginInsert(m)
	mBase = m.M
	v403 = m.ExcPending
	if v403 != 0 {
		goto L3
	} else {
		goto L94
	}
L92:
	;
	v401 = *(*int32)(unsafe.Add(mBase, uint32(l0)+40))
	if v401 != 0 {
		goto L87
	} else {
		goto L93
	}
L93:
	;
	goto L91
L94:
	;
	F_XLogRegisterBuffer(m, int32(0), l2, int32(8))
	mBase = m.M
	v407 = m.ExcPending
	if v407 != 0 {
		goto L3
	} else {
		goto L95
	}
L95:
	;
	v410 = F_XLogInsert(m, int32(12), int32(160))
	mBase = m.M
	v411 = m.ExcPending
	if v411 != 0 {
		goto L3
	} else {
		goto L96
	}
L96:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v377))) = base.I64_rotr(v410, int64(32))
	goto L87
L97:
	;
	m.G0 = v2155 + int32(_a_F_hashbucketcleanup_0)
	return
L98:
	;
	F_LockBuffer(m, l2, int32(0))
	mBase = m.M
	v2130 = m.ExcPending
	if v2130 != 0 {
		goto L3
	} else {
		goto L440
	}
L99:
	;
	v425 = F_IsBufferCleanupOK(m, l2)
	mBase = m.M
	v426 = m.ExcPending
	if v426 != 0 {
		goto L3
	} else {
		goto L100
	}
L100:
	;
	if v425 == int32(0) {
		goto L98
	} else {
		goto L101
	}
L101:
	;
	v429 = m.G0
	v431 = v429 + int32(-8192)
	m.G0 = v431
	if l2 < int32(0) {
		goto L104
	} else {
		goto L105
	}
L102:
	;
	m.G0 = v2110 - int32(-8192)
	v2155 = v2115
	goto L97
L103:
	;
	v451 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v450)+16)))
	v452 = v451 + v450
	v453 = *(*int32)(unsafe.Add(mBase, uint32(v452)+4))
	if v453 != int32(-1) {
		goto L107
	} else {
		goto L108
	}
L104:
	;
	v436 = *(*int32)(unsafe.Add(mBase, _c_F_hashbucketcleanup[0]))
	v442 = *(*int32)(unsafe.Add(mBase, uint32(v436+(l2^int32(-1))<<(uint(int32(2))%32))))
	v450 = v442
	goto L103
L105:
	;
	goto L106
L106:
	;
	v444 = *(*int32)(unsafe.Add(mBase, _c_F_hashbucketcleanup[1]))
	v450 = v444 + l2<<(uint(int32(13))%32) + int32(-8192)
	goto L103
L107:
	;
	v465 = v453
	v473 = v14
	goto L110
L108:
	;
	goto L109
L109:
	;
	F_LockBuffer(m, l2, int32(0))
	mBase = m.M
	v2090 = m.ExcPending
	if v2090 != 0 {
		goto L3
	} else {
		goto L439
	}
L110:
	;
	if v473 != 0 {
		goto L112
	} else {
		goto L113
	}
L111:
	;
	v518 = l3
	v519 = l2
	v520 = l2
	v521 = l3
	v523 = l4
	v524 = v465
	v526 = l0
	v531 = v514
	v535 = v493
	v536 = v512
	v537 = v431
	v542 = v37
	v543 = v450
	v544 = v452
	goto L122
L112:
	;
	F_UnlockReleaseBuffer(m, v473)
	mBase = m.M
	v491 = m.ExcPending
	if v491 != 0 {
		goto L3
	} else {
		goto L115
	}
L113:
	;
	goto L114
L114:
	;
	v493 = F__hash_getbuf_with_strategy(m, l0, v465, int32(1), l4)
	mBase = m.M
	v494 = m.ExcPending
	if v494 != 0 {
		goto L3
	} else {
		goto L117
	}
L115:
	;
	goto L114
L116:
	;
	v513 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v512)+16)))
	v514 = v513 + v512
	v515 = *(*int32)(unsafe.Add(mBase, uint32(v514)+4))
	if v515 != int32(-1) {
		v465 = v515
		v473 = v493
		goto L110
	} else {
		goto L121
	}
L117:
	;
	if v493 < int32(0) {
		goto L118
	} else {
		goto L119
	}
L118:
	;
	v498 = *(*int32)(unsafe.Add(mBase, _c_F_hashbucketcleanup[0]))
	v504 = *(*int32)(unsafe.Add(mBase, uint32(v498+(v493^int32(-1))<<(uint(int32(2))%32))))
	v512 = v504
	goto L116
L119:
	;
	goto L120
L120:
	;
	v506 = *(*int32)(unsafe.Add(mBase, _c_F_hashbucketcleanup[1]))
	v512 = v506 + v493<<(uint(int32(13))%32) + int32(-8192)
	goto L116
L121:
	;
	goto L111
L122:
	;
	v552 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v536)+12)))
	if base.Ui32(v552) < base.Ui32(int32(25)) {
		goto L125
	} else {
		goto L126
	}
L123:
	;
	F_UnlockReleaseBuffer(m, v1135)
	mBase = m.M
	v2087 = m.ExcPending
	if v2087 != 0 {
		goto L3
	} else {
		goto L438
	}
L124:
	;
	v1168 = *(*int32)(unsafe.Add(mBase, uint32(v531)))
	v1170 = v537 + int32(2464)
	v1172 = v537 + int32(16)
	v1176 = v1154 & int32(_a_F_hashbucketcleanup_4)
	v1178 = m.G0
	v1180 = v1178 - int32(32)
	m.G0 = v1180
	F__hash_checkpage(m, v526, v535, int32(1))
	mBase = m.M
	v1184 = m.ExcPending
	if v1184 != 0 {
		goto L3
	} else {
		goto L223
	}
L125:
	;
	v1134 = v518
	v1135 = v519
	v1154 = int32(0)
	v1159 = v543
	v1160 = v544
	goto L124
L126:
	;
	goto L127
L127:
	;
	v559 = int32(base.Ui32(v552+int32(_a_F_hashbucketcleanup_1)) >> (uint(int32(2)) % 32))
	if v559&int32(_a_F_hashbucketcleanup_4) == int32(0) {
		goto L128
	} else {
		goto L129
	}
L128:
	;
	v1134 = v518
	v1135 = v519
	v1154 = int32(0)
	v1159 = v543
	v1160 = v544
	goto L124
L129:
	;
	goto L130
L130:
	;
	v573 = v518
	v574 = v519
	v582 = v559
	v598 = v543
	v599 = v544
	goto L131
L131:
	;
	v608 = int32(0)
	v613 = v573
	v614 = v574
	v622 = v608
	v633 = v608
	v635 = v608
	v636 = int32(1)
	v638 = v598
	v639 = v599
	goto L133
L132:
	;
	v1134 = v1095
	v1135 = v1096
	v1154 = v1115
	v1159 = v1120
	v1160 = v1121
	goto L124
L133:
	;
	v654 = *(*int32)(unsafe.Add(mBase, uint32(v636&int32(_a_F_hashbucketcleanup_4)<<(uint(int32(2))%32)+(v536+int32(24))-int32(4))))
	v655 = int32(_a_F_hashbucketcleanup_8)
	if v654&v655 != v655 {
		goto L135
	} else {
		goto L136
	}
L134:
	;
	goto L132
L135:
	;
	v661 = v536 + v654&int32(_a_F_hashbucketcleanup_3)
	v662 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v661)+6)))
	v668 = (v662&int32(_a_F_hashbucketcleanup_9) + int32(7)) & int32(_a_F_hashbucketcleanup_10)
	v669 = v613
	v670 = v614
	v678 = v622
	v689 = v633
	v691 = v635
	v694 = v638
	v695 = v639
	goto L139
L136:
	;
	v1095 = v613
	v1096 = v614
	v1104 = v622
	v1115 = v633
	v1117 = v635
	v1120 = v638
	v1121 = v639
	goto L137
L137:
	;
	v1130 = v636 + int32(1)
	if base.Ui32(v1130&int32(_a_F_hashbucketcleanup_4)) <= base.Ui32(v582&int32(_a_F_hashbucketcleanup_4)) {
		v613 = v1095
		v614 = v1096
		v622 = v1104
		v633 = v1115
		v635 = v1117
		v636 = v1130
		v638 = v1120
		v639 = v1121
		goto L133
	} else {
		goto L222
	}
L138:
	;
	*(*uint16)(unsafe.Add(mBase, uint32(v537+int32(_a_F_hashbucketcleanup_11)+v678&int32(_a_F_hashbucketcleanup_4)<<(uint(int32(1))%32)))) = uint16(v636)
	v1079 = F_CopyIndexTuple(m, v661)
	mBase = m.M
	v1080 = m.ExcPending
	if v1080 != 0 {
		goto L3
	} else {
		goto L221
	}
L139:
	;
	v704 = v689 & int32(_a_F_hashbucketcleanup_4)
	v707 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v694)+14)))
	v708 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v694)+12)))
	v709 = v707 - v708
	v711 = (v704 + int32(1)) << (uint(int32(2)) % 32)
	if v711 <= v709 {
		goto L142
	} else {
		goto L143
	}
L140:
	;
	v1024 = v1010
	goto L213
L141:
	;
	v716 = v668 + v691
	if base.Ui32(v716) <= base.Ui32(v715) {
		goto L138
	} else {
		goto L145
	}
L142:
	;
	v715 = v709 - v711
	goto L144
L143:
	;
	v715 = int32(0)
	goto L144
L144:
	;
	goto L141
L145:
	;
	v719 = *(*int32)(unsafe.Add(mBase, uint32(v695)+4))
	if v524 != v719 {
		goto L146
	} else {
		goto L147
	}
L146:
	;
	v722 = F__hash_getbuf_with_strategy(m, v526, v719, int32(1), v523)
	mBase = m.M
	v723 = m.ExcPending
	if v723 != 0 {
		goto L3
	} else {
		goto L149
	}
L147:
	;
	v724 = int32(0)
	goto L148
L148:
	;
	if v704 != 0 {
		goto L150
	} else {
		goto L151
	}
L149:
	;
	v724 = v722
	goto L148
L150:
	;
	v725 = *(*int32)(unsafe.Add(mBase, uint32(v526)+48))
	v726 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v725)+118)))
	if v726 != int32(112) {
		goto L153
	} else {
		goto L154
	}
L151:
	;
	goto L152
L152:
	;
	if v669 == v521 {
		goto L199
	} else {
		goto L200
	}
L153:
	;
	v740 = int32(_a_F_hashbucketcleanup_5)
	v742 = *(*int32)(unsafe.Add(mBase, _c_F_hashbucketcleanup[2]))
	*(*int32)(unsafe.Add(mBase, _c_F_hashbucketcleanup[2])) = v742 + int32(1)
	F__hash_pgaddmultitup(m, v526, v670, v537+int32(2464), v537+int32(16), v704)
	mBase = m.M
	v751 = m.ExcPending
	if v751 != 0 {
		goto L3
	} else {
		goto L161
	}
L154:
	;
	v730 = *(*int32)(unsafe.Add(mBase, _c_F_hashbucketcleanup[3]))
	if v730 <= int32(0) {
		goto L155
	} else {
		goto L156
	}
L155:
	;
	v733 = *(*int32)(unsafe.Add(mBase, uint32(v526)+32))
	if v733 != 0 {
		goto L153
	} else {
		goto L158
	}
L156:
	;
	goto L157
L157:
	;
	F_XLogEnsureRecordSpace(m, int32(0), v704+int32(3))
	mBase = m.M
	v739 = m.ExcPending
	if v739 != 0 {
		goto L3
	} else {
		goto L160
	}
L158:
	;
	v734 = *(*int32)(unsafe.Add(mBase, uint32(v526)+40))
	if v734 != 0 {
		goto L153
	} else {
		goto L159
	}
L159:
	;
	goto L157
L160:
	;
	goto L153
L161:
	;
	F_MarkBufferDirty(m, v670)
	mBase = m.M
	v753 = m.ExcPending
	if v753 != 0 {
		goto L3
	} else {
		goto L162
	}
L162:
	;
	v757 = v678 & int32(_a_F_hashbucketcleanup_4)
	F_PageIndexMultiDelete(m, v536, v537+int32(_a_F_hashbucketcleanup_11), v757)
	mBase = m.M
	v759 = m.ExcPending
	if v759 != 0 {
		goto L3
	} else {
		goto L163
	}
L163:
	;
	F_MarkBufferDirty(m, v535)
	mBase = m.M
	v761 = m.ExcPending
	if v761 != 0 {
		goto L3
	} else {
		goto L164
	}
L164:
	;
	v762 = *(*int32)(unsafe.Add(mBase, uint32(v526)+48))
	v763 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v762)+118)))
	if v763 != int32(112) {
		goto L165
	} else {
		goto L166
	}
L165:
	;
	v941 = int32(_a_F_hashbucketcleanup_5)
	v943 = *(*int32)(unsafe.Add(mBase, _c_F_hashbucketcleanup[2]))
	*(*int32)(unsafe.Add(mBase, _c_F_hashbucketcleanup[2])) = v943 - int32(1)
	goto L152
L166:
	;
	v767 = *(*int32)(unsafe.Add(mBase, _c_F_hashbucketcleanup[3]))
	if v767 <= int32(0) {
		goto L167
	} else {
		goto L168
	}
L167:
	;
	v770 = *(*int32)(unsafe.Add(mBase, uint32(v526)+32))
	if v770 != 0 {
		goto L165
	} else {
		goto L170
	}
L168:
	;
	goto L169
L169:
	;
	*(*uint16)(unsafe.Add(mBase, uint32(v537)+12)) = uint16(v689)
	*(*uint8)(unsafe.Add(mBase, uint32(v537)+14)) = uint8(base.B2i32(v670 == v520))
	F_XLogBeginInsert(m)
	mBase = m.M
	v776 = m.ExcPending
	if v776 != 0 {
		goto L3
	} else {
		goto L172
	}
L170:
	;
	v771 = *(*int32)(unsafe.Add(mBase, uint32(v526)+40))
	if v771 != 0 {
		goto L165
	} else {
		goto L171
	}
L171:
	;
	goto L169
L172:
	;
	F_XLogRegisterData(m, v537+int32(12), int32(3))
	mBase = m.M
	v781 = m.ExcPending
	if v781 != 0 {
		goto L3
	} else {
		goto L173
	}
L173:
	;
	v782 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v537)+14)))
	if v782 == int32(0) {
		goto L174
	} else {
		goto L175
	}
L174:
	;
	F_XLogRegisterBuffer(m, int32(0), v520, int32(42))
	mBase = m.M
	v788 = m.ExcPending
	if v788 != 0 {
		goto L3
	} else {
		goto L177
	}
L175:
	;
	goto L176
L176:
	;
	F_XLogRegisterBuffer(m, int32(1), v670, int32(8))
	mBase = m.M
	v792 = m.ExcPending
	if v792 != 0 {
		goto L3
	} else {
		goto L178
	}
L177:
	;
	goto L176
L178:
	;
	v793 = int32(1)
	F_XLogRegisterBufData(m, v793, v537+int32(16), v704<<(uint(v793)%32))
	mBase = m.M
	v799 = m.ExcPending
	if v799 != 0 {
		goto L3
	} else {
		goto L179
	}
L179:
	;
	v800 = int32(1)
	if base.Ui32(v704) <= base.Ui32(v800) {
		goto L180
	} else {
		goto L181
	}
L180:
	;
	v803 = v800
	goto L182
L181:
	;
	v803 = v704
	goto L182
L182:
	;
	v814 = int32(0)
	goto L183
L183:
	;
	v841 = v814 << (uint(int32(2)) % 32)
	v845 = *(*int32)(unsafe.Add(mBase, uint32(v841+(v537+int32(2464)))))
	v849 = *(*int32)(unsafe.Add(mBase, uint32(v537+int32(832)+v841)))
	F_XLogRegisterBufData(m, int32(1), v845, v849)
	mBase = m.M
	v851 = m.ExcPending
	if v851 != 0 {
		goto L3
	} else {
		goto L185
	}
L184:
	;
	F_XLogRegisterBuffer(m, int32(2), v535, int32(8))
	mBase = m.M
	v858 = m.ExcPending
	if v858 != 0 {
		goto L3
	} else {
		goto L187
	}
L185:
	;
	v853 = v814 + int32(1)
	if v853 != v803 {
		v814 = v853
		goto L183
	} else {
		goto L186
	}
L186:
	;
	goto L184
L187:
	;
	F_XLogRegisterBufData(m, int32(2), v537+int32(_a_F_hashbucketcleanup_11), v757<<(uint(int32(1))%32))
	mBase = m.M
	v865 = m.ExcPending
	if v865 != 0 {
		goto L3
	} else {
		goto L188
	}
L188:
	;
	v868 = F_XLogInsert(m, int32(12), int32(112))
	mBase = m.M
	v869 = m.ExcPending
	if v869 != 0 {
		goto L3
	} else {
		goto L189
	}
L189:
	;
	if v670 < int32(0) {
		goto L191
	} else {
		goto L192
	}
L190:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v887))) = base.I64_rotr(v868, int64(32))
	if v535 < int32(0) {
		goto L195
	} else {
		goto L196
	}
L191:
	;
	v873 = *(*int32)(unsafe.Add(mBase, _c_F_hashbucketcleanup[0]))
	v879 = *(*int32)(unsafe.Add(mBase, uint32(v873+(v670^int32(-1))<<(uint(int32(2))%32))))
	v887 = v879
	goto L190
L192:
	;
	goto L193
L193:
	;
	v881 = *(*int32)(unsafe.Add(mBase, _c_F_hashbucketcleanup[1]))
	v887 = v881 + v670<<(uint(int32(13))%32) + int32(-8192)
	goto L190
L194:
	;
	*(*uint32)(unsafe.Add(mBase, uint32(v902)+4)) = uint32(v868)
	v905 = int64(base.Ui64(v868) >> (uint(int64(32)) % 64))
	*(*uint32)(unsafe.Add(mBase, uint32(v902))) = uint32(v905)
	goto L165
L195:
	;
	v894 = *(*int32)(unsafe.Add(mBase, _c_F_hashbucketcleanup[0]))
	v896 = *(*int32)(unsafe.Add(mBase, uint32(v894+(v535^int32(-1))<<(uint(int32(2))%32))))
	v902 = v896
	goto L194
L196:
	;
	goto L197
L197:
	;
	v898 = *(*int32)(unsafe.Add(mBase, _c_F_hashbucketcleanup[1]))
	v902 = v898 + v535<<(uint(int32(13))%32) + int32(-8192)
	goto L194
L198:
	;
	if v719 == v524 {
		goto L204
	} else {
		goto L205
	}
L199:
	;
	F_LockBuffer(m, v670, int32(0))
	mBase = m.M
	v984 = m.ExcPending
	if v984 != 0 {
		goto L3
	} else {
		goto L202
	}
L200:
	;
	goto L201
L201:
	;
	F_UnlockReleaseBuffer(m, v670)
	mBase = m.M
	v986 = m.ExcPending
	if v986 != 0 {
		goto L3
	} else {
		goto L203
	}
L202:
	;
	goto L198
L203:
	;
	goto L198
L204:
	;
	F_UnlockReleaseBuffer(m, v535)
	mBase = m.M
	v989 = m.ExcPending
	if v989 != 0 {
		goto L3
	} else {
		goto L207
	}
L205:
	;
	goto L206
L206:
	;
	if v724 < int32(0) {
		goto L209
	} else {
		goto L210
	}
L207:
	;
	v2110 = v537
	v2115 = v542
	goto L102
L208:
	;
	v1008 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v1007)+16)))
	v1009 = v1008 + v1007
	v1010 = int32(0)
	if v704 == v1010 {
		v669 = v719
		v670 = v724
		v678 = v1010
		v689 = v1010
		v691 = v1010
		v694 = v1007
		v695 = v1009
		goto L139
	} else {
		goto L212
	}
L209:
	;
	v993 = *(*int32)(unsafe.Add(mBase, _c_F_hashbucketcleanup[0]))
	v999 = *(*int32)(unsafe.Add(mBase, uint32(v993+(v724^int32(-1))<<(uint(int32(2))%32))))
	v1007 = v999
	goto L208
L210:
	;
	goto L211
L211:
	;
	v1001 = *(*int32)(unsafe.Add(mBase, _c_F_hashbucketcleanup[1]))
	v1007 = v1001 + v724<<(uint(int32(13))%32) + int32(-8192)
	goto L208
L212:
	;
	goto L140
L213:
	;
	v1054 = *(*int32)(unsafe.Add(mBase, uint32(v537+int32(2464)+v1024<<(uint(int32(2))%32))))
	F_pfree(m, v1054)
	mBase = m.M
	v1056 = m.ExcPending
	if v1056 != 0 {
		goto L3
	} else {
		goto L215
	}
L214:
	;
	v1060 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v536)+12)))
	if base.Ui32(int32(25)) <= base.Ui32(v1060) {
		goto L217
	} else {
		goto L218
	}
L215:
	;
	v1058 = v1024 + int32(1)
	if v1058 != v704 {
		v1024 = v1058
		goto L213
	} else {
		goto L216
	}
L216:
	;
	goto L214
L217:
	;
	v1068 = int32(base.Ui32(v1060+int32(_a_F_hashbucketcleanup_1)) >> (uint(int32(2)) % 32))
	goto L219
L218:
	;
	v1068 = int32(0)
	goto L219
L219:
	;
	if v1068&int32(_a_F_hashbucketcleanup_4) != 0 {
		v573 = v719
		v574 = v724
		v582 = v1068
		v598 = v1007
		v599 = v1009
		goto L131
	} else {
		goto L220
	}
L220:
	;
	v1134 = v719
	v1135 = v724
	v1154 = v1010
	v1159 = v1007
	v1160 = v1009
	goto L124
L221:
	;
	v1082 = v704 << (uint(int32(2)) % 32)
	*(*int32)(unsafe.Add(mBase, uint32(v1082+(v537+int32(832))))) = v668
	*(*int32)(unsafe.Add(mBase, uint32(v537+int32(2464)+v1082))) = v1079
	v1091 = int32(1)
	v1095 = v669
	v1096 = v670
	v1104 = v678 + v1091
	v1115 = v689 + v1091
	v1117 = v716
	v1120 = v694
	v1121 = v695
	goto L137
L222:
	;
	goto L134
L223:
	;
	if v535 < int32(0) {
		goto L225
	} else {
		goto L226
	}
L224:
	;
	if v535 < int32(0) {
		goto L229
	} else {
		goto L230
	}
L225:
	;
	v1188 = *(*int32)(unsafe.Add(mBase, _c_F_hashbucketcleanup[4]))
	v1194 = *(*int32)(unsafe.Add(mBase, uint32(v1188+(v535^int32(-1))<<(uint(int32(6))%32))+16))
	v1203 = v1194
	goto L224
L226:
	;
	goto L227
L227:
	;
	v1196 = *(*int32)(unsafe.Add(mBase, _c_F_hashbucketcleanup[5]))
	v1202 = *(*int32)(unsafe.Add(mBase, uint32(v1196+v535<<(uint(int32(6))%32)+int32(-64))+16))
	v1203 = v1202
	goto L224
L228:
	;
	v1222 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v1221)+16)))
	v1223 = v1222 + v1221
	v1224 = *(*int32)(unsafe.Add(mBase, uint32(v1223)))
	v1225 = *(*int32)(unsafe.Add(mBase, uint32(v1223)+4))
	if v1135 < int32(0) {
		goto L233
	} else {
		goto L234
	}
L229:
	;
	v1207 = *(*int32)(unsafe.Add(mBase, _c_F_hashbucketcleanup[0]))
	v1213 = *(*int32)(unsafe.Add(mBase, uint32(v1207+(v535^int32(-1))<<(uint(int32(2))%32))))
	v1221 = v1213
	goto L228
L230:
	;
	goto L231
L231:
	;
	v1215 = *(*int32)(unsafe.Add(mBase, _c_F_hashbucketcleanup[1]))
	v1221 = v1215 + v535<<(uint(int32(13))%32) + int32(-8192)
	goto L228
L232:
	;
	if v1224 == int32(-1) {
		v1252 = int32(0)
		goto L236
	} else {
		goto L237
	}
L233:
	;
	v1229 = *(*int32)(unsafe.Add(mBase, _c_F_hashbucketcleanup[4]))
	v1235 = *(*int32)(unsafe.Add(mBase, uint32(v1229+(v1135^int32(-1))<<(uint(int32(6))%32))+16))
	v1244 = v1235
	goto L232
L234:
	;
	goto L235
L235:
	;
	v1237 = *(*int32)(unsafe.Add(mBase, _c_F_hashbucketcleanup[5]))
	v1243 = *(*int32)(unsafe.Add(mBase, uint32(v1237+v1135<<(uint(int32(6))%32)+int32(-64))+16))
	v1244 = v1243
	goto L232
L236:
	;
	if v1225 != int32(-1) {
		goto L240
	} else {
		goto L241
	}
L237:
	;
	if v1224 == v1244 {
		v1252 = v1135
		goto L236
	} else {
		goto L238
	}
L238:
	;
	v1250 = F__hash_getbuf_with_strategy(m, v526, v1224, int32(3), v523)
	mBase = m.M
	v1251 = m.ExcPending
	if v1251 != 0 {
		goto L3
	} else {
		goto L239
	}
L239:
	;
	v1252 = v1250
	goto L236
L240:
	;
	v1256 = F__hash_getbuf_with_strategy(m, v526, v1225, int32(1), v523)
	mBase = m.M
	v1257 = m.ExcPending
	if v1257 != 0 {
		goto L3
	} else {
		goto L243
	}
L241:
	;
	v1258 = int32(0)
	goto L242
L242:
	;
	v1262 = F__hash_getbuf(m, v526, int32(0), int32(1), int32(8))
	mBase = m.M
	v1263 = m.ExcPending
	if v1263 != 0 {
		goto L3
	} else {
		goto L245
	}
L243:
	;
	v1258 = v1256
	goto L242
L244:
	;
	v1284 = m.G0
	v1286 = v1284 - int32(16)
	m.G0 = v1286
	v1288 = *(*int32)(unsafe.Add(mBase, uint32(v1281+int32(24))+36))
	if v1288 == int32(0) {
		goto L250
	} else {
		goto L251
	}
L245:
	;
	if v1262 < int32(0) {
		goto L246
	} else {
		goto L247
	}
L246:
	;
	v1267 = *(*int32)(unsafe.Add(mBase, _c_F_hashbucketcleanup[0]))
	v1273 = *(*int32)(unsafe.Add(mBase, uint32(v1267+(v1262^int32(-1))<<(uint(int32(2))%32))))
	v1281 = v1273
	goto L244
L247:
	;
	goto L248
L248:
	;
	v1275 = *(*int32)(unsafe.Add(mBase, _c_F_hashbucketcleanup[1]))
	v1281 = v1275 + v1262<<(uint(int32(13))%32) + int32(-8192)
	goto L244
L249:
	;
	m.G0 = v1286 + int32(16)
	v1439 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v1281)+46)))
	v1440 = int32(1)
	v1441 = v1373 - v1440
	v1442 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v1281)+44)))
	*(*int32)(unsafe.Add(mBase, uint32(v1180)+28)) = v1441 & (v1442<<(uint(int32(3))%32) - v1440)
	v1449 = int32(base.Ui32(v1441) >> (uint(v1439) % 32))
	v1450 = *(*int32)(unsafe.Add(mBase, uint32(v1281)+68))
	if base.Ui32(v1449) < base.Ui32(v1450) {
		goto L273
	} else {
		goto L274
	}
L250:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v1423 = m.ExcPending
	if v1423 != 0 {
		goto L3
	} else {
		goto L268
	}
L251:
	;
	v1303 = int32(1)
	goto L252
L252:
	;
	if base.Ui32(v1303) <= base.Ui32(int32(9)) {
		goto L255
	} else {
		goto L256
	}
L253:
	;
	goto L250
L254:
	;
	if base.Ui32(v1203) <= base.Ui32(v1349) {
		goto L250
	} else {
		goto L258
	}
L255:
	;
	v1349 = int32(1) << (uint(v1303) % 32)
	goto L254
L256:
	;
	goto L257
L257:
	;
	v1335 = v1303 - int32(10)
	v1336 = int32(2)
	v1338 = int32(512) << (uint(int32(base.Ui32(v1335)>>(uint(v1336)%32))) % 32)
	v1349 = v1338>>(uint(v1336)%32)*(v1335&int32(3)+int32(1)) + v1338
	goto L254
L258:
	;
	if base.Ui32(v1303) <= base.Ui32(int32(9)) {
		goto L260
	} else {
		goto L261
	}
L259:
	;
	v1373 = v1203 - v1372
	v1376 = v1303<<(uint(int32(2))%32) + (v1281 + int32(76))
	v1379 = *(*int32)(unsafe.Add(mBase, uint32(v1376-int32(4))))
	if base.Ui32(v1379) < base.Ui32(v1373) {
		goto L263
	} else {
		goto L264
	}
L260:
	;
	v1372 = int32(1) << (uint(v1303) % 32)
	goto L259
L261:
	;
	goto L262
L262:
	;
	v1358 = v1303 - int32(10)
	v1359 = int32(2)
	v1361 = int32(512) << (uint(int32(base.Ui32(v1358)>>(uint(v1359)%32))) % 32)
	v1372 = v1361>>(uint(v1359)%32)*(v1358&int32(3)+int32(1)) + v1361
	goto L259
L263:
	;
	v1381 = *(*int32)(unsafe.Add(mBase, uint32(v1376)))
	if base.Ui32(v1373) <= base.Ui32(v1381) {
		goto L249
	} else {
		goto L266
	}
L264:
	;
	goto L265
L265:
	;
	v1384 = v1303 + int32(1)
	if base.Ui32(v1384) <= base.Ui32(v1288) {
		v1303 = v1384
		goto L252
	} else {
		goto L267
	}
L266:
	;
	goto L265
L267:
	;
	goto L253
L268:
	;
	F_errcode(m, int32(50856066))
	mBase = m.M
	v1426 = m.ExcPending
	if v1426 != 0 {
		goto L3
	} else {
		goto L269
	}
L269:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v1286))) = v1203
	F_errmsg(m, int32(_a_F_hashbucketcleanup_12), v1286)
	mBase = m.M
	v1430 = m.ExcPending
	if v1430 != 0 {
		goto L3
	} else {
		goto L270
	}
L270:
	;
	F_errfinish(m, int32(_a_F_hashbucketcleanup_13), int32(88), int32(_a_F_hashbucketcleanup_14))
	mBase = m.M
	v1435 = m.ExcPending
	if v1435 != 0 {
		goto L3
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
	if v1176 != 0 {
		goto L420
	} else {
		goto L421
	}
L273:
	;
	v1455 = *(*int32)(unsafe.Add(mBase, uint32(v1281+v1449<<(uint(int32(2))%32))+468))
	F_LockBuffer(m, v1262, int32(0))
	mBase = m.M
	v1458 = m.ExcPending
	if v1458 != 0 {
		goto L3
	} else {
		goto L276
	}
L274:
	;
	goto L275
L275:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v1968 = m.ExcPending
	if v1968 != 0 {
		goto L3
	} else {
		goto L417
	}
L276:
	;
	v1461 = F__hash_getbuf(m, v526, v1455, int32(2), int32(4))
	mBase = m.M
	v1462 = m.ExcPending
	if v1462 != 0 {
		goto L3
	} else {
		goto L278
	}
L277:
	;
	F_LockBuffer(m, v1262, int32(2))
	mBase = m.M
	v1483 = m.ExcPending
	if v1483 != 0 {
		goto L3
	} else {
		goto L282
	}
L278:
	;
	if v1461 < int32(0) {
		goto L279
	} else {
		goto L280
	}
L279:
	;
	v1466 = *(*int32)(unsafe.Add(mBase, _c_F_hashbucketcleanup[0]))
	v1472 = *(*int32)(unsafe.Add(mBase, uint32(v1466+(v1461^int32(-1))<<(uint(int32(2))%32))))
	v1480 = v1472
	goto L277
L280:
	;
	goto L281
L281:
	;
	v1474 = *(*int32)(unsafe.Add(mBase, _c_F_hashbucketcleanup[1]))
	v1480 = v1474 + v1461<<(uint(int32(13))%32) + int32(-8192)
	goto L277
L282:
	;
	v1484 = *(*int32)(unsafe.Add(mBase, uint32(v526)+48))
	v1485 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1484)+118)))
	if v1485 != int32(112) {
		goto L283
	} else {
		goto L284
	}
L283:
	;
	v1499 = int32(_a_F_hashbucketcleanup_5)
	v1501 = *(*int32)(unsafe.Add(mBase, _c_F_hashbucketcleanup[2]))
	*(*int32)(unsafe.Add(mBase, _c_F_hashbucketcleanup[2])) = v1501 + int32(1)
	if v1176 != 0 {
		goto L291
	} else {
		goto L292
	}
L284:
	;
	v1489 = *(*int32)(unsafe.Add(mBase, _c_F_hashbucketcleanup[3]))
	if v1489 <= int32(0) {
		goto L285
	} else {
		goto L286
	}
L285:
	;
	v1492 = *(*int32)(unsafe.Add(mBase, uint32(v526)+32))
	if v1492 != 0 {
		goto L283
	} else {
		goto L288
	}
L286:
	;
	goto L287
L287:
	;
	F_XLogEnsureRecordSpace(m, int32(6), v1176+int32(4))
	mBase = m.M
	v1498 = m.ExcPending
	if v1498 != 0 {
		goto L3
	} else {
		goto L290
	}
L288:
	;
	v1493 = *(*int32)(unsafe.Add(mBase, uint32(v526)+40))
	if v1493 != 0 {
		goto L283
	} else {
		goto L289
	}
L289:
	;
	goto L287
L290:
	;
	goto L283
L291:
	;
	F__hash_pgaddmultitup(m, v526, v1135, v1170, v1172, v1176)
	mBase = m.M
	v1506 = m.ExcPending
	if v1506 != 0 {
		goto L3
	} else {
		goto L294
	}
L292:
	;
	goto L293
L293:
	;
	F_PageInit(m, v1221, int32(_a_F_hashbucketcleanup_15), int32(16))
	mBase = m.M
	goto L296
L294:
	;
	F_MarkBufferDirty(m, v1135)
	mBase = m.M
	v1508 = m.ExcPending
	if v1508 != 0 {
		goto L3
	} else {
		goto L295
	}
L295:
	;
	goto L293
L296:
	;
	v1512 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v1221)+16)))
	v1513 = v1221 + v1512
	*(*int64)(unsafe.Add(mBase, uint32(v1513)+8)) = int64(-36028792723996673)
	*(*int64)(unsafe.Add(mBase, uint32(v1513))) = int64(-1)
	F_MarkBufferDirty(m, v535)
	mBase = m.M
	v1519 = m.ExcPending
	if v1519 != 0 {
		goto L3
	} else {
		goto L297
	}
L297:
	;
	if v1252 != 0 {
		goto L298
	} else {
		goto L299
	}
L298:
	;
	if v1252 < int32(0) {
		goto L302
	} else {
		goto L303
	}
L299:
	;
	goto L300
L300:
	;
	if v1258 != 0 {
		goto L306
	} else {
		goto L307
	}
L301:
	;
	v1538 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v1537)+16)))
	*(*int32)(unsafe.Add(mBase, uint32(v1538+v1537)+4)) = v1225
	F_MarkBufferDirty(m, v1252)
	mBase = m.M
	v1542 = m.ExcPending
	if v1542 != 0 {
		goto L3
	} else {
		goto L305
	}
L302:
	;
	v1523 = *(*int32)(unsafe.Add(mBase, _c_F_hashbucketcleanup[0]))
	v1529 = *(*int32)(unsafe.Add(mBase, uint32(v1523+(v1252^int32(-1))<<(uint(int32(2))%32))))
	v1537 = v1529
	goto L301
L303:
	;
	goto L304
L304:
	;
	v1531 = *(*int32)(unsafe.Add(mBase, _c_F_hashbucketcleanup[1]))
	v1537 = v1531 + v1252<<(uint(int32(13))%32) + int32(-8192)
	goto L301
L305:
	;
	goto L300
L306:
	;
	if v1258 < int32(0) {
		goto L310
	} else {
		goto L311
	}
L307:
	;
	goto L308
L308:
	;
	v1570 = *(*int32)(unsafe.Add(mBase, uint32(v1180)+28))
	v1572 = base.I32_div_s(v1570, int32(32))
	v1575 = v1480 + int32(24) + v1572<<(uint(int32(2))%32)
	v1576 = *(*int32)(unsafe.Add(mBase, uint32(v1575)))
	*(*int32)(unsafe.Add(mBase, uint32(v1575))) = v1576 & base.I32_rotl(int32(-2), v1570)
	F_MarkBufferDirty(m, v1461)
	mBase = m.M
	v1582 = m.ExcPending
	if v1582 != 0 {
		goto L3
	} else {
		goto L314
	}
L309:
	;
	v1564 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v1563)+16)))
	*(*int32)(unsafe.Add(mBase, uint32(v1564+v1563))) = v1224
	F_MarkBufferDirty(m, v1258)
	mBase = m.M
	v1568 = m.ExcPending
	if v1568 != 0 {
		goto L3
	} else {
		goto L313
	}
L310:
	;
	v1549 = *(*int32)(unsafe.Add(mBase, _c_F_hashbucketcleanup[0]))
	v1555 = *(*int32)(unsafe.Add(mBase, uint32(v1549+(v1258^int32(-1))<<(uint(int32(2))%32))))
	v1563 = v1555
	goto L309
L311:
	;
	goto L312
L312:
	;
	v1557 = *(*int32)(unsafe.Add(mBase, _c_F_hashbucketcleanup[1]))
	v1563 = v1557 + v1258<<(uint(int32(13))%32) + int32(-8192)
	goto L309
L313:
	;
	goto L308
L314:
	;
	v1584 = v1281 - int32(-64)
	v1585 = *(*int32)(unsafe.Add(mBase, uint32(v1584)))
	v1586 = base.B2i32(base.Ui32(v1585) <= base.Ui32(v1441))
	if v1586 == int32(0) {
		goto L315
	} else {
		goto L316
	}
L315:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v1584))) = v1441
	F_MarkBufferDirty(m, v1262)
	mBase = m.M
	v1591 = m.ExcPending
	if v1591 != 0 {
		goto L3
	} else {
		goto L318
	}
L316:
	;
	goto L317
L317:
	;
	v1592 = *(*int32)(unsafe.Add(mBase, uint32(v526)+48))
	v1593 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1592)+118)))
	if v1593 != int32(112) {
		goto L319
	} else {
		goto L320
	}
L318:
	;
	goto L317
L319:
	;
	v1943 = int32(_a_F_hashbucketcleanup_5)
	v1945 = *(*int32)(unsafe.Add(mBase, _c_F_hashbucketcleanup[2]))
	*(*int32)(unsafe.Add(mBase, _c_F_hashbucketcleanup[2])) = v1945 - int32(1)
	if v1252 == int32(0) {
		goto L403
	} else {
		goto L404
	}
L320:
	;
	v1597 = *(*int32)(unsafe.Add(mBase, _c_F_hashbucketcleanup[3]))
	if v1597 <= int32(0) {
		goto L321
	} else {
		goto L322
	}
L321:
	;
	v1600 = *(*int32)(unsafe.Add(mBase, uint32(v526)+32))
	if v1600 != 0 {
		goto L319
	} else {
		goto L324
	}
L322:
	;
	goto L323
L323:
	;
	*(*uint16)(unsafe.Add(mBase, uint32(v1180)+24)) = uint16(v1176)
	*(*int32)(unsafe.Add(mBase, uint32(v1180)+20)) = v1225
	*(*int32)(unsafe.Add(mBase, uint32(v1180)+16)) = v1224
	*(*uint8)(unsafe.Add(mBase, uint32(v1180)+27)) = uint8(base.B2i32(v1135 == v1252))
	*(*uint8)(unsafe.Add(mBase, uint32(v1180)+26)) = uint8(base.B2i32(v1135 == v520))
	F_XLogBeginInsert(m)
	mBase = m.M
	v1610 = m.ExcPending
	if v1610 != 0 {
		goto L3
	} else {
		goto L326
	}
L324:
	;
	v1601 = *(*int32)(unsafe.Add(mBase, uint32(v526)+40))
	if v1601 != 0 {
		goto L319
	} else {
		goto L325
	}
L325:
	;
	goto L323
L326:
	;
	F_XLogRegisterData(m, v1180+int32(16), int32(12))
	mBase = m.M
	v1615 = m.ExcPending
	if v1615 != 0 {
		goto L3
	} else {
		goto L327
	}
L327:
	;
	v1616 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1180)+26)))
	if v1616 == int32(0) {
		goto L328
	} else {
		goto L329
	}
L328:
	;
	F_XLogRegisterBuffer(m, int32(0), v520, int32(42))
	mBase = m.M
	v1622 = m.ExcPending
	if v1622 != 0 {
		goto L3
	} else {
		goto L331
	}
L329:
	;
	goto L330
L330:
	;
	v1623 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v1180)+24)))
	if v1623 != 0 {
		goto L333
	} else {
		goto L334
	}
L331:
	;
	goto L330
L332:
	;
	F_XLogRegisterBuffer(m, int32(2), v535, int32(8))
	mBase = m.M
	v1740 = m.ExcPending
	if v1740 != 0 {
		goto L3
	} else {
		goto L352
	}
L333:
	;
	v1624 = int32(1)
	F_XLogRegisterBuffer(m, v1624, v1135, int32(8))
	mBase = m.M
	v1628 = m.ExcPending
	if v1628 != 0 {
		goto L3
	} else {
		goto L336
	}
L334:
	;
	goto L335
L335:
	;
	v1683 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1180)+27)))
	v1685 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1180)+26)))
	if v1685 == int32(0) {
		goto L344
	} else {
		goto L345
	}
L336:
	;
	v1629 = int32(1)
	F_XLogRegisterBufData(m, v1629, v1172, v1176<<(uint(v1629)%32))
	mBase = m.M
	v1633 = m.ExcPending
	if v1633 != 0 {
		goto L3
	} else {
		goto L337
	}
L337:
	;
	if v1176 == int32(0) {
		v1712 = v1624
		goto L332
	} else {
		goto L338
	}
L338:
	;
	v1644 = int32(0)
	goto L339
L339:
	;
	v1673 = v1644 << (uint(int32(2)) % 32)
	v1675 = *(*int32)(unsafe.Add(mBase, uint32(v1170+v1673)))
	v1677 = *(*int32)(unsafe.Add(mBase, uint32(v1673+(v537+int32(832)))))
	F_XLogRegisterBufData(m, int32(1), v1675, v1677)
	mBase = m.M
	v1679 = m.ExcPending
	if v1679 != 0 {
		goto L3
	} else {
		goto L341
	}
L340:
	;
	v1712 = v1624
	goto L332
L341:
	;
	v1681 = v1644 + int32(1)
	if v1681 != v1176 {
		v1644 = v1681
		goto L339
	} else {
		goto L342
	}
L342:
	;
	goto L340
L343:
	;
	F_XLogRegisterBuffer(m, int32(1), v1135, v1700)
	mBase = m.M
	v1702 = m.ExcPending
	if v1702 != 0 {
		goto L3
	} else {
		goto L351
	}
L344:
	;
	v1688 = int32(1)
	if v1683&v1688 != 0 {
		v1698 = v1688
		v1700 = int32(8)
		goto L343
	} else {
		goto L347
	}
L345:
	;
	goto L346
L346:
	;
	if v1683&int32(1) != 0 {
		goto L348
	} else {
		goto L349
	}
L347:
	;
	v1712 = int32(0)
	goto L332
L348:
	;
	v1697 = int32(8)
	goto L350
L349:
	;
	v1697 = int32(40)
	goto L350
L350:
	;
	v1698 = v1683
	v1700 = v1697
	goto L343
L351:
	;
	v1712 = v1698
	goto L332
L352:
	;
	if v1252 == int32(0) {
		goto L353
	} else {
		goto L354
	}
L353:
	;
	if v1258 != 0 {
		goto L357
	} else {
		goto L358
	}
L354:
	;
	v1743 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1180)+27)))
	if v1743 != 0 {
		goto L353
	} else {
		goto L355
	}
L355:
	;
	F_XLogRegisterBuffer(m, int32(3), v1252, int32(8))
	mBase = m.M
	v1747 = m.ExcPending
	if v1747 != 0 {
		goto L3
	} else {
		goto L356
	}
L356:
	;
	goto L353
L357:
	;
	F_XLogRegisterBuffer(m, int32(4), v1258, int32(8))
	mBase = m.M
	v1751 = m.ExcPending
	if v1751 != 0 {
		goto L3
	} else {
		goto L360
	}
L358:
	;
	goto L359
L359:
	;
	F_XLogRegisterBuffer(m, int32(5), v1461, int32(8))
	mBase = m.M
	v1755 = m.ExcPending
	if v1755 != 0 {
		goto L3
	} else {
		goto L361
	}
L360:
	;
	goto L359
L361:
	;
	F_XLogRegisterBufData(m, int32(5), v1180+int32(28), int32(4))
	mBase = m.M
	v1761 = m.ExcPending
	if v1761 != 0 {
		goto L3
	} else {
		goto L362
	}
L362:
	;
	if v1586 == int32(0) {
		goto L363
	} else {
		goto L364
	}
L363:
	;
	F_XLogRegisterBuffer(m, int32(6), v1262, int32(8))
	mBase = m.M
	v1767 = m.ExcPending
	if v1767 != 0 {
		goto L3
	} else {
		goto L366
	}
L364:
	;
	goto L365
L365:
	;
	v1774 = F_XLogInsert(m, int32(12), int32(128))
	mBase = m.M
	v1775 = m.ExcPending
	if v1775 != 0 {
		goto L3
	} else {
		goto L368
	}
L366:
	;
	F_XLogRegisterBufData(m, int32(6), v1584, int32(4))
	mBase = m.M
	v1771 = m.ExcPending
	if v1771 != 0 {
		goto L3
	} else {
		goto L367
	}
L367:
	;
	goto L365
L368:
	;
	if v1712&int32(1) != 0 {
		goto L369
	} else {
		goto L370
	}
L369:
	;
	if v1135 < int32(0) {
		goto L373
	} else {
		goto L374
	}
L370:
	;
	goto L371
L371:
	;
	if v535 < int32(0) {
		goto L377
	} else {
		goto L378
	}
L372:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v1795))) = base.I64_rotr(v1774, int64(32))
	goto L371
L373:
	;
	v1781 = *(*int32)(unsafe.Add(mBase, _c_F_hashbucketcleanup[0]))
	v1787 = *(*int32)(unsafe.Add(mBase, uint32(v1781+(v1135^int32(-1))<<(uint(int32(2))%32))))
	v1795 = v1787
	goto L372
L374:
	;
	goto L375
L375:
	;
	v1789 = *(*int32)(unsafe.Add(mBase, _c_F_hashbucketcleanup[1]))
	v1795 = v1789 + v1135<<(uint(int32(13))%32) + int32(-8192)
	goto L372
L376:
	;
	v1817 = int64(32)
	*(*int64)(unsafe.Add(mBase, uint32(v1816))) = base.I64_rotr(v1774, v1817)
	v1822 = base.I32_wrap_i64(int64(base.Ui64(v1774) >> (uint(v1817) % 64)))
	v1823 = base.I32_wrap_i64(v1774)
	if v1252 == int32(0) {
		goto L380
	} else {
		goto L381
	}
L377:
	;
	v1802 = *(*int32)(unsafe.Add(mBase, _c_F_hashbucketcleanup[0]))
	v1808 = *(*int32)(unsafe.Add(mBase, uint32(v1802+(v535^int32(-1))<<(uint(int32(2))%32))))
	v1816 = v1808
	goto L376
L378:
	;
	goto L379
L379:
	;
	v1810 = *(*int32)(unsafe.Add(mBase, _c_F_hashbucketcleanup[1]))
	v1816 = v1810 + v535<<(uint(int32(13))%32) + int32(-8192)
	goto L376
L380:
	;
	if v1258 != 0 {
		goto L387
	} else {
		goto L388
	}
L381:
	;
	v1826 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1180)+27)))
	if v1826 != 0 {
		goto L380
	} else {
		goto L382
	}
L382:
	;
	if v1252 < int32(0) {
		goto L384
	} else {
		goto L385
	}
L383:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v1844)+4)) = v1823
	*(*int32)(unsafe.Add(mBase, uint32(v1844))) = v1822
	goto L380
L384:
	;
	v1830 = *(*int32)(unsafe.Add(mBase, _c_F_hashbucketcleanup[0]))
	v1836 = *(*int32)(unsafe.Add(mBase, uint32(v1830+(v1252^int32(-1))<<(uint(int32(2))%32))))
	v1844 = v1836
	goto L383
L385:
	;
	goto L386
L386:
	;
	v1838 = *(*int32)(unsafe.Add(mBase, _c_F_hashbucketcleanup[1]))
	v1844 = v1838 + v1252<<(uint(int32(13))%32) + int32(-8192)
	goto L383
L387:
	;
	if v1258 < int32(0) {
		goto L391
	} else {
		goto L392
	}
L388:
	;
	goto L389
L389:
	;
	if v1461 < int32(0) {
		goto L395
	} else {
		goto L396
	}
L390:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v1865)+4)) = v1823
	*(*int32)(unsafe.Add(mBase, uint32(v1865))) = v1822
	goto L389
L391:
	;
	v1851 = *(*int32)(unsafe.Add(mBase, _c_F_hashbucketcleanup[0]))
	v1857 = *(*int32)(unsafe.Add(mBase, uint32(v1851+(v1258^int32(-1))<<(uint(int32(2))%32))))
	v1865 = v1857
	goto L390
L392:
	;
	goto L393
L393:
	;
	v1859 = *(*int32)(unsafe.Add(mBase, _c_F_hashbucketcleanup[1]))
	v1865 = v1859 + v1258<<(uint(int32(13))%32) + int32(-8192)
	goto L390
L394:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v1886)+4)) = v1823
	*(*int32)(unsafe.Add(mBase, uint32(v1886))) = v1822
	if base.Ui32(v1585) <= base.Ui32(v1441) {
		goto L319
	} else {
		goto L398
	}
L395:
	;
	v1872 = *(*int32)(unsafe.Add(mBase, _c_F_hashbucketcleanup[0]))
	v1878 = *(*int32)(unsafe.Add(mBase, uint32(v1872+(v1461^int32(-1))<<(uint(int32(2))%32))))
	v1886 = v1878
	goto L394
L396:
	;
	goto L397
L397:
	;
	v1880 = *(*int32)(unsafe.Add(mBase, _c_F_hashbucketcleanup[1]))
	v1886 = v1880 + v1461<<(uint(int32(13))%32) + int32(-8192)
	goto L394
L398:
	;
	if v1262 < int32(0) {
		goto L400
	} else {
		goto L401
	}
L399:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v1906)+4)) = v1823
	*(*int32)(unsafe.Add(mBase, uint32(v1906))) = v1822
	goto L319
L400:
	;
	v1892 = *(*int32)(unsafe.Add(mBase, _c_F_hashbucketcleanup[0]))
	v1898 = *(*int32)(unsafe.Add(mBase, uint32(v1892+(v1262^int32(-1))<<(uint(int32(2))%32))))
	v1906 = v1898
	goto L399
L401:
	;
	goto L402
L402:
	;
	v1900 = *(*int32)(unsafe.Add(mBase, _c_F_hashbucketcleanup[1]))
	v1906 = v1900 + v1262<<(uint(int32(13))%32) + int32(-8192)
	goto L399
L403:
	;
	if v535 != 0 {
		goto L407
	} else {
		goto L408
	}
L404:
	;
	if v1224 == v1244 {
		goto L403
	} else {
		goto L405
	}
L405:
	;
	F_UnlockReleaseBuffer(m, v1252)
	mBase = m.M
	v1953 = m.ExcPending
	if v1953 != 0 {
		goto L3
	} else {
		goto L406
	}
L406:
	;
	goto L403
L407:
	;
	F_UnlockReleaseBuffer(m, v535)
	mBase = m.M
	v1955 = m.ExcPending
	if v1955 != 0 {
		goto L3
	} else {
		goto L410
	}
L408:
	;
	goto L409
L409:
	;
	if v1258 != 0 {
		goto L411
	} else {
		goto L412
	}
L410:
	;
	goto L409
L411:
	;
	F_UnlockReleaseBuffer(m, v1258)
	mBase = m.M
	v1957 = m.ExcPending
	if v1957 != 0 {
		goto L3
	} else {
		goto L414
	}
L412:
	;
	goto L413
L413:
	;
	F_UnlockReleaseBuffer(m, v1461)
	mBase = m.M
	v1959 = m.ExcPending
	if v1959 != 0 {
		goto L3
	} else {
		goto L415
	}
L414:
	;
	goto L413
L415:
	;
	F_UnlockReleaseBuffer(m, v1262)
	mBase = m.M
	v1961 = m.ExcPending
	if v1961 != 0 {
		goto L3
	} else {
		goto L416
	}
L416:
	;
	m.G0 = v1180 + int32(32)
	goto L272
L417:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v1180))) = v1441
	F_errmsg_internal(m, int32(_a_F_hashbucketcleanup_16), v1180)
	mBase = m.M
	v1972 = m.ExcPending
	if v1972 != 0 {
		goto L3
	} else {
		goto L418
	}
L418:
	;
	F_errfinish(m, int32(_a_F_hashbucketcleanup_13), int32(562), int32(_a_F_hashbucketcleanup_17))
	mBase = m.M
	v1977 = m.ExcPending
	if v1977 != 0 {
		goto L3
	} else {
		goto L419
	}
L419:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L420:
	;
	v1988 = int32(0)
	goto L423
L421:
	;
	goto L422
L422:
	;
	if v1134 == v1168 {
		goto L428
	} else {
		goto L429
	}
L423:
	;
	v2018 = *(*int32)(unsafe.Add(mBase, uint32(v537+int32(2464)+v1988<<(uint(int32(2))%32))))
	F_pfree(m, v2018)
	mBase = m.M
	v2020 = m.ExcPending
	if v2020 != 0 {
		goto L3
	} else {
		goto L425
	}
L424:
	;
	goto L422
L425:
	;
	v2022 = v1988 + int32(1)
	if v2022 != v1176 {
		v1988 = v2022
		goto L423
	} else {
		goto L426
	}
L426:
	;
	goto L424
L427:
	;
	goto L123
L428:
	;
	if v1134 != v521 {
		goto L427
	} else {
		goto L431
	}
L429:
	;
	goto L430
L430:
	;
	v2064 = F__hash_getbuf_with_strategy(m, v526, v1168, int32(1), v523)
	mBase = m.M
	v2065 = m.ExcPending
	if v2065 != 0 {
		goto L3
	} else {
		goto L434
	}
L431:
	;
	F_LockBuffer(m, v1135, int32(0))
	mBase = m.M
	v2062 = m.ExcPending
	if v2062 != 0 {
		goto L3
	} else {
		goto L432
	}
L432:
	;
	v2110 = v537
	v2115 = v542
	goto L102
L433:
	;
	v2084 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v2083)+16)))
	v518 = v1134
	v519 = v1135
	v524 = v1168
	v531 = v2084 + v2083
	v535 = v2064
	v536 = v2083
	v543 = v1159
	v544 = v1160
	goto L122
L434:
	;
	if v2064 < int32(0) {
		goto L435
	} else {
		goto L436
	}
L435:
	;
	v2069 = *(*int32)(unsafe.Add(mBase, _c_F_hashbucketcleanup[0]))
	v2075 = *(*int32)(unsafe.Add(mBase, uint32(v2069+(v2064^int32(-1))<<(uint(int32(2))%32))))
	v2083 = v2075
	goto L433
L436:
	;
	goto L437
L437:
	;
	v2077 = *(*int32)(unsafe.Add(mBase, _c_F_hashbucketcleanup[1]))
	v2083 = v2077 + v2064<<(uint(int32(13))%32) + int32(-8192)
	goto L433
L438:
	;
	v2110 = v537
	v2115 = v542
	goto L102
L439:
	;
	v2110 = v431
	v2115 = v37
	goto L102
L440:
	;
	v2155 = v37
	goto L97
}
func F_hashbyteaextended(m *base.Module, l0 int32) int32 {
	var v2 int32
	_ = v2
	var v5 int32
	_ = v5
	v2 = F_hashvarlenaextended(m, l0)
	v5 = m.ExcPending
	if v5 != 0 {
		return int32(0)
	} else {
		return v2
	}
}
func F_hashfloat8(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v4 int32
	_ = v4
	var v6 int32
	_ = v6
	var v8 int32
	_ = v8
	var v9 float64
	_ = v9
	var v22 int32
	_ = v22
	var v29 int32
	_ = v29
	var v92 int32
	_ = v92
	var v96 int32
	_ = v96
	var v97 int32
	_ = v97
	var v101 int32
	_ = v101
	var v102 int32
	_ = v102
	var v106 int32
	_ = v106
	var v107 int32
	_ = v107
	var v109 int32
	_ = v109
	var v111 int32
	_ = v111
	var v115 int32
	_ = v115
	var v116 int32
	_ = v116
	var v120 int32
	_ = v120
	var v121 int32
	_ = v121
	var v123 int32
	_ = v123
	var v124 int32
	_ = v124
	var v126 int32
	_ = v126
	var v130 int32
	_ = v130
	var v131 int32
	_ = v131
	var v135 int32
	_ = v135
	var v136 int32
	_ = v136
	var v194 int32
	_ = v194
	var v198 int32
	_ = v198
	var v199 int32
	_ = v199
	var v203 int32
	_ = v203
	var v204 int32
	_ = v204
	var v208 int32
	_ = v208
	var v209 int32
	_ = v209
	var v213 int32
	_ = v213
	var v214 int32
	_ = v214
	var v215 int32
	_ = v215
	var v219 int32
	_ = v219
	var v220 int32
	_ = v220
	var v221 int32
	_ = v221
	var v225 int32
	_ = v225
	var v226 int32
	_ = v226
	var v227 int32
	_ = v227
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
	var v238 int32
	_ = v238
	var v242 int32
	_ = v242
	var v243 int32
	_ = v243
	var v244 int32
	_ = v244
	var v245 int32
	_ = v245
	var v249 int32
	_ = v249
	var v250 int32
	_ = v250
	var v251 int32
	_ = v251
	var v252 int32
	_ = v252
	var v256 int32
	_ = v256
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
	var v271 int32
	_ = v271
	var v275 int32
	_ = v275
	var v279 int32
	_ = v279
	var v283 int32
	_ = v283
	var v288 int32
	_ = v288
	v4 = m.G0
	v6 = v4 - int32(16)
	m.G0 = v6
	v8 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v9 = *(*float64)(unsafe.Add(mBase, uint32(v8)))
	*(*float64)(unsafe.Add(mBase, uint32(v6)+8)) = v9
	if base.F64_ne(v9, float64(0)) != 0 {
		if base.Ui64(int64(9218868437227405313)) <= base.Ui64(base.I64_reinterpret_f64(v9)&int64(9223372036854775807)) {
			*(*int64)(unsafe.Add(mBase, uint32(v6)+8)) = int64(9221120237041090560)
		} else {
		}
		v22 = v6 + int32(8)
		v29 = int32(-1636608424)
		if v22&int32(3) != 0 {
			switch int32(7) {
			case 0:
				v249 = v29
				v250 = v29
				v251 = v29
				v252 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v22))))
				v256 = v249 + v252
				v257 = v250
				v258 = v251
			case 1:
				v242 = v29
				v243 = v29
				v244 = v29
				v245 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v22)+1)))
				v249 = v245<<(uint(int32(8))%32) + v242
				v250 = v243
				v251 = v244
				v252 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v22))))
				v256 = v249 + v252
				v257 = v250
				v258 = v251
			case 2:
				v235 = v29
				v236 = v29
				v237 = v29
				v238 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v22)+2)))
				v242 = v238<<(uint(int32(16))%32) + v235
				v243 = v236
				v244 = v237
				v245 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v22)+1)))
				v249 = v245<<(uint(int32(8))%32) + v242
				v250 = v243
				v251 = v244
				v252 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v22))))
				v256 = v249 + v252
				v257 = v250
				v258 = v251
			case 3:
				v229 = v29
				v230 = v29
				v231 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v22)+3)))
				v235 = v231<<(uint(int32(24))%32) + v29
				v236 = v229
				v237 = v230
				v238 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v22)+2)))
				v242 = v238<<(uint(int32(16))%32) + v235
				v243 = v236
				v244 = v237
				v245 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v22)+1)))
				v249 = v245<<(uint(int32(8))%32) + v242
				v250 = v243
				v251 = v244
				v252 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v22))))
				v256 = v249 + v252
				v257 = v250
				v258 = v251
			case 4:
				v225 = v29
				v226 = v29
				v227 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v22)+4)))
				v229 = v225 + v227
				v230 = v226
				v231 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v22)+3)))
				v235 = v231<<(uint(int32(24))%32) + v29
				v236 = v229
				v237 = v230
				v238 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v22)+2)))
				v242 = v238<<(uint(int32(16))%32) + v235
				v243 = v236
				v244 = v237
				v245 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v22)+1)))
				v249 = v245<<(uint(int32(8))%32) + v242
				v250 = v243
				v251 = v244
				v252 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v22))))
				v256 = v249 + v252
				v257 = v250
				v258 = v251
			case 5:
				v219 = v29
				v220 = v29
				v221 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v22)+5)))
				v225 = v221<<(uint(int32(8))%32) + v219
				v226 = v220
				v227 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v22)+4)))
				v229 = v225 + v227
				v230 = v226
				v231 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v22)+3)))
				v235 = v231<<(uint(int32(24))%32) + v29
				v236 = v229
				v237 = v230
				v238 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v22)+2)))
				v242 = v238<<(uint(int32(16))%32) + v235
				v243 = v236
				v244 = v237
				v245 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v22)+1)))
				v249 = v245<<(uint(int32(8))%32) + v242
				v250 = v243
				v251 = v244
				v252 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v22))))
				v256 = v249 + v252
				v257 = v250
				v258 = v251
			case 6:
				v213 = v29
				v214 = v29
				v215 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v22)+6)))
				v219 = v215<<(uint(int32(16))%32) + v213
				v220 = v214
				v221 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v22)+5)))
				v225 = v221<<(uint(int32(8))%32) + v219
				v226 = v220
				v227 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v22)+4)))
				v229 = v225 + v227
				v230 = v226
				v231 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v22)+3)))
				v235 = v231<<(uint(int32(24))%32) + v29
				v236 = v229
				v237 = v230
				v238 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v22)+2)))
				v242 = v238<<(uint(int32(16))%32) + v235
				v243 = v236
				v244 = v237
				v245 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v22)+1)))
				v249 = v245<<(uint(int32(8))%32) + v242
				v250 = v243
				v251 = v244
				v252 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v22))))
				v256 = v249 + v252
				v257 = v250
				v258 = v251
			case 7:
				v208 = v29
				v209 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v22)+7)))
				v213 = v209<<(uint(int32(24))%32) + v29
				v214 = v208
				v215 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v22)+6)))
				v219 = v215<<(uint(int32(16))%32) + v213
				v220 = v214
				v221 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v22)+5)))
				v225 = v221<<(uint(int32(8))%32) + v219
				v226 = v220
				v227 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v22)+4)))
				v229 = v225 + v227
				v230 = v226
				v231 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v22)+3)))
				v235 = v231<<(uint(int32(24))%32) + v29
				v236 = v229
				v237 = v230
				v238 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v22)+2)))
				v242 = v238<<(uint(int32(16))%32) + v235
				v243 = v236
				v244 = v237
				v245 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v22)+1)))
				v249 = v245<<(uint(int32(8))%32) + v242
				v250 = v243
				v251 = v244
				v252 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v22))))
				v256 = v249 + v252
				v257 = v250
				v258 = v251
			case 8:
				v203 = v29
				v204 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v22)+8)))
				v208 = v204<<(uint(int32(8))%32) + v203
				v209 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v22)+7)))
				v213 = v209<<(uint(int32(24))%32) + v29
				v214 = v208
				v215 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v22)+6)))
				v219 = v215<<(uint(int32(16))%32) + v213
				v220 = v214
				v221 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v22)+5)))
				v225 = v221<<(uint(int32(8))%32) + v219
				v226 = v220
				v227 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v22)+4)))
				v229 = v225 + v227
				v230 = v226
				v231 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v22)+3)))
				v235 = v231<<(uint(int32(24))%32) + v29
				v236 = v229
				v237 = v230
				v238 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v22)+2)))
				v242 = v238<<(uint(int32(16))%32) + v235
				v243 = v236
				v244 = v237
				v245 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v22)+1)))
				v249 = v245<<(uint(int32(8))%32) + v242
				v250 = v243
				v251 = v244
				v252 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v22))))
				v256 = v249 + v252
				v257 = v250
				v258 = v251
			case 9:
				v198 = v29
				v199 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v22)+9)))
				v203 = v199<<(uint(int32(16))%32) + v198
				v204 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v22)+8)))
				v208 = v204<<(uint(int32(8))%32) + v203
				v209 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v22)+7)))
				v213 = v209<<(uint(int32(24))%32) + v29
				v214 = v208
				v215 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v22)+6)))
				v219 = v215<<(uint(int32(16))%32) + v213
				v220 = v214
				v221 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v22)+5)))
				v225 = v221<<(uint(int32(8))%32) + v219
				v226 = v220
				v227 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v22)+4)))
				v229 = v225 + v227
				v230 = v226
				v231 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v22)+3)))
				v235 = v231<<(uint(int32(24))%32) + v29
				v236 = v229
				v237 = v230
				v238 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v22)+2)))
				v242 = v238<<(uint(int32(16))%32) + v235
				v243 = v236
				v244 = v237
				v245 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v22)+1)))
				v249 = v245<<(uint(int32(8))%32) + v242
				v250 = v243
				v251 = v244
				v252 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v22))))
				v256 = v249 + v252
				v257 = v250
				v258 = v251
			case 10:
				v194 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v22)+10)))
				v198 = v194<<(uint(int32(24))%32) + v29
				v199 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v22)+9)))
				v203 = v199<<(uint(int32(16))%32) + v198
				v204 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v22)+8)))
				v208 = v204<<(uint(int32(8))%32) + v203
				v209 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v22)+7)))
				v213 = v209<<(uint(int32(24))%32) + v29
				v214 = v208
				v215 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v22)+6)))
				v219 = v215<<(uint(int32(16))%32) + v213
				v220 = v214
				v221 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v22)+5)))
				v225 = v221<<(uint(int32(8))%32) + v219
				v226 = v220
				v227 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v22)+4)))
				v229 = v225 + v227
				v230 = v226
				v231 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v22)+3)))
				v235 = v231<<(uint(int32(24))%32) + v29
				v236 = v229
				v237 = v230
				v238 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v22)+2)))
				v242 = v238<<(uint(int32(16))%32) + v235
				v243 = v236
				v244 = v237
				v245 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v22)+1)))
				v249 = v245<<(uint(int32(8))%32) + v242
				v250 = v243
				v251 = v244
				v252 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v22))))
				v256 = v249 + v252
				v257 = v250
				v258 = v251
			default:
				v256 = v29
				v257 = v29
				v258 = v29
			}
		} else {
			switch int32(7) {
			case 0:
				v135 = v29
				v136 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v22))))
				v256 = v135 + v136
				v257 = v29
				v258 = v29
			case 1:
				v130 = v29
				v131 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v22)+1)))
				v135 = v131<<(uint(int32(8))%32) + v130
				v136 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v22))))
				v256 = v135 + v136
				v257 = v29
				v258 = v29
			case 2:
				v126 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v22)+2)))
				v130 = v126<<(uint(int32(16))%32) + v29
				v131 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v22)+1)))
				v135 = v131<<(uint(int32(8))%32) + v130
				v136 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v22))))
				v256 = v135 + v136
				v257 = v29
				v258 = v29
			case 3:
				v123 = v29
				v124 = *(*int32)(unsafe.Add(mBase, uint32(v22)))
				v256 = v124 + v29
				v257 = v123
				v258 = v29
			case 4:
				v120 = v29
				v121 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v22)+4)))
				v123 = v120 + v121
				v124 = *(*int32)(unsafe.Add(mBase, uint32(v22)))
				v256 = v124 + v29
				v257 = v123
				v258 = v29
			case 5:
				v115 = v29
				v116 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v22)+5)))
				v120 = v116<<(uint(int32(8))%32) + v115
				v121 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v22)+4)))
				v123 = v120 + v121
				v124 = *(*int32)(unsafe.Add(mBase, uint32(v22)))
				v256 = v124 + v29
				v257 = v123
				v258 = v29
			case 6:
				v111 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v22)+6)))
				v115 = v111<<(uint(int32(16))%32) + v29
				v116 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v22)+5)))
				v120 = v116<<(uint(int32(8))%32) + v115
				v121 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v22)+4)))
				v123 = v120 + v121
				v124 = *(*int32)(unsafe.Add(mBase, uint32(v22)))
				v256 = v124 + v29
				v257 = v123
				v258 = v29
			case 7:
				v106 = v29
				v107 = *(*int32)(unsafe.Add(mBase, uint32(v22)))
				v109 = *(*int32)(unsafe.Add(mBase, uint32(v22)+4))
				v256 = v107 + v29
				v257 = v109 + v29
				v258 = v106
			case 8:
				v101 = v29
				v102 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v22)+8)))
				v106 = v102<<(uint(int32(8))%32) + v101
				v107 = *(*int32)(unsafe.Add(mBase, uint32(v22)))
				v109 = *(*int32)(unsafe.Add(mBase, uint32(v22)+4))
				v256 = v107 + v29
				v257 = v109 + v29
				v258 = v106
			case 9:
				v96 = v29
				v97 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v22)+9)))
				v101 = v97<<(uint(int32(16))%32) + v96
				v102 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v22)+8)))
				v106 = v102<<(uint(int32(8))%32) + v101
				v107 = *(*int32)(unsafe.Add(mBase, uint32(v22)))
				v109 = *(*int32)(unsafe.Add(mBase, uint32(v22)+4))
				v256 = v107 + v29
				v257 = v109 + v29
				v258 = v106
			case 10:
				v92 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v22)+10)))
				v96 = v92<<(uint(int32(24))%32) + v29
				v97 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v22)+9)))
				v101 = v97<<(uint(int32(16))%32) + v96
				v102 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v22)+8)))
				v106 = v102<<(uint(int32(8))%32) + v101
				v107 = *(*int32)(unsafe.Add(mBase, uint32(v22)))
				v109 = *(*int32)(unsafe.Add(mBase, uint32(v22)+4))
				v256 = v107 + v29
				v257 = v109 + v29
				v258 = v106
			default:
				v256 = v29
				v257 = v29
				v258 = v29
			}
		}
		v261 = int32(14)
		v263 = v257 ^ v258 - base.I32_rotl(v257, v261)
		v267 = v263 ^ v256 - base.I32_rotl(v263, int32(11))
		v271 = v267 ^ v257 - base.I32_rotl(v267, int32(25))
		v275 = v271 ^ v263 - base.I32_rotl(v271, int32(16))
		v279 = v275 ^ v267 - base.I32_rotl(v275, int32(4))
		v283 = v279 ^ v271 - base.I32_rotl(v279, v261)
		v288 = v283 ^ v275 - base.I32_rotl(v283, int32(24))
	} else {
		v288 = int32(0)
	}
	m.G0 = v6 + int32(16)
	return v288
}
func F_hashgettuple(m *base.Module, l0 int32, l1 int32) int32 {
	mBase := m.M
	_ = mBase
	var v6 int32
	_ = v6
	var v8 int32
	_ = v8
	var v9 int32
	_ = v9
	var v12 int32
	_ = v12
	var v15 int32
	_ = v15
	var v17 int32
	_ = v17
	var v20 int32
	_ = v20
	var v24 int32
	_ = v24
	var v25 int32
	_ = v25
	var v27 int32
	_ = v27
	var v28 int32
	_ = v28
	var v37 int32
	_ = v37
	var v41 int32
	_ = v41
	var v42 int32
	_ = v42
	v6 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(l0)+72)) = uint8(v6)
	v8 = *(*int32)(unsafe.Add(mBase, uint32(l0)+36))
	v9 = *(*int32)(unsafe.Add(mBase, uint32(v8)+28))
	if v9 == int32(-1) {
		v12 = F__hash_first(m, l0, l1)
		mBase = m.M
		v15 = m.ExcPending
		if v15 != 0 {
			return int32(0)
		} else {
			return v12
		}
	} else {
		v17 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+30)))
		if v17 != int32(1) {
			v41 = F__hash_next(m, l0, l1)
			mBase = m.M
			v42 = m.ExcPending
			if v42 != 0 {
				return int32(0)
			} else {
				return v41
			}
		} else {
			v20 = *(*int32)(unsafe.Add(mBase, uint32(v8)+16))
			if v20 == int32(0) {
				v24 = F_palloc(m, int32(1632))
				mBase = m.M
				v25 = m.ExcPending
				if v25 != 0 {
					return int32(0)
				} else {
					*(*int32)(unsafe.Add(mBase, uint32(v8)+16)) = v24
					v27 = v24
					v28 = *(*int32)(unsafe.Add(mBase, uint32(v8)+20))
					if int32(407) < v28 {
					} else {
						*(*int32)(unsafe.Add(mBase, uint32(v8)+20)) = v28 + int32(1)
						v37 = *(*int32)(unsafe.Add(mBase, uint32(v8)+48))
						*(*int32)(unsafe.Add(mBase, uint32(v27+v28<<(uint(int32(2))%32)))) = v37
					}
					v41 = F__hash_next(m, l0, l1)
					mBase = m.M
					v42 = m.ExcPending
					if v42 != 0 {
						return int32(0)
					} else {
						return v41
					}
				}
			} else {
				v27 = v20
				v28 = *(*int32)(unsafe.Add(mBase, uint32(v8)+20))
				if int32(407) < v28 {
				} else {
					*(*int32)(unsafe.Add(mBase, uint32(v8)+20)) = v28 + int32(1)
					v37 = *(*int32)(unsafe.Add(mBase, uint32(v8)+48))
					*(*int32)(unsafe.Add(mBase, uint32(v27+v28<<(uint(int32(2))%32)))) = v37
				}
				v41 = F__hash_next(m, l0, l1)
				mBase = m.M
				v42 = m.ExcPending
				if v42 != 0 {
					return int32(0)
				} else {
					return v41
				}
			}
		}
	}
}
func F_hashhandler(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v3 int32
	_ = v3
	var v6 int32
	_ = v6
	var v7 int32
	_ = v7
	var v9 int32
	_ = v9
	var v13 int32
	_ = v13
	var v19 int32
	_ = v19
	var v35 int64
	_ = v35
	v3 = F_palloc0(m, int32(140))
	mBase = m.M
	v6 = m.ExcPending
	if v6 != 0 {
		return int32(0)
	} else {
		v7 = int32(0)
		*(*int32)(unsafe.Add(mBase, uint32(v3)+16)) = v7
		v9 = int32(256)
		*(*uint16)(unsafe.Add(mBase, uint32(v3)+14)) = uint16(v9)
		*(*int32)(unsafe.Add(mBase, uint32(v3)+10)) = int32(16842752)
		v13 = int32(3)
		*(*uint16)(unsafe.Add(mBase, uint32(v3)+8)) = uint16(v13)
		*(*int64)(unsafe.Add(mBase, uint32(v3))) = int64(844429225099702)
		*(*int32)(unsafe.Add(mBase, uint32(v3)+24)) = v7
		v19 = int32(1)
		*(*uint8)(unsafe.Add(mBase, uint32(v3)+23)) = uint8(v19)
		*(*int32)(unsafe.Add(mBase, uint32(v3)+108)) = int32(121)
		*(*int32)(unsafe.Add(mBase, uint32(v3)+104)) = int32(122)
		*(*int32)(unsafe.Add(mBase, uint32(v3)+100)) = int32(123)
		*(*int32)(unsafe.Add(mBase, uint32(v3)+96)) = int32(124)
		*(*int32)(unsafe.Add(mBase, uint32(v3)+92)) = int32(125)
		*(*int32)(unsafe.Add(mBase, uint32(v3)+88)) = int32(126)
		*(*int32)(unsafe.Add(mBase, uint32(v3)+84)) = int32(127)
		v35 = int64(0)
		*(*int64)(unsafe.Add(mBase, uint32(v3)+76)) = v35
		*(*int32)(unsafe.Add(mBase, uint32(v3)+72)) = int32(128)
		*(*int32)(unsafe.Add(mBase, uint32(v3)+68)) = v7
		*(*int32)(unsafe.Add(mBase, uint32(v3)+64)) = int32(129)
		*(*int32)(unsafe.Add(mBase, uint32(v3)+60)) = v7
		*(*int32)(unsafe.Add(mBase, uint32(v3)+56)) = int32(130)
		*(*int32)(unsafe.Add(mBase, uint32(v3)+52)) = int32(131)
		*(*int32)(unsafe.Add(mBase, uint32(v3)+48)) = v7
		*(*int32)(unsafe.Add(mBase, uint32(v3)+44)) = int32(132)
		*(*int32)(unsafe.Add(mBase, uint32(v3)+40)) = int32(133)
		*(*int32)(unsafe.Add(mBase, uint32(v3)+36)) = int32(134)
		*(*int32)(unsafe.Add(mBase, uint32(v3)+32)) = int32(23)
		*(*int32)(unsafe.Add(mBase, uint32(v3)+19)) = v7
		*(*uint16)(unsafe.Add(mBase, uint32(v3)+28)) = uint16(v9)
		*(*int32)(unsafe.Add(mBase, uint32(v3)+128)) = v7
		*(*int64)(unsafe.Add(mBase, uint32(v3)+120)) = v35
		*(*int64)(unsafe.Add(mBase, uint32(v3)+112)) = v35
		*(*int32)(unsafe.Add(mBase, uint32(v3)+132)) = int32(135)
		*(*int32)(unsafe.Add(mBase, uint32(v3)+136)) = int32(136)
		return v3
	}
}
func F_hashint2extended(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v2 int32
	_ = v2
	var v3 int32
	_ = v3
	var v4 int64
	_ = v4
	var v12 int32
	_ = v12
	var v15 int32
	_ = v15
	var v20 int32
	_ = v20
	var v26 int32
	_ = v26
	var v28 int32
	_ = v28
	var v29 int32
	_ = v29
	var v30 int32
	_ = v30
	var v34 int32
	_ = v34
	var v38 int32
	_ = v38
	var v42 int32
	_ = v42
	var v43 int32
	_ = v43
	var v44 int32
	_ = v44
	var v50 int32
	_ = v50
	var v52 int32
	_ = v52
	var v55 int32
	_ = v55
	var v57 int32
	_ = v57
	var v59 int32
	_ = v59
	var v64 int32
	_ = v64
	var v68 int32
	_ = v68
	var v72 int32
	_ = v72
	var v76 int32
	_ = v76
	var v80 int32
	_ = v80
	var v90 int32
	_ = v90
	var v93 int32
	_ = v93
	v2 = int32(*(*int16)(unsafe.Add(mBase, uint32(l0)+20)))
	v3 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
	v4 = *(*int64)(unsafe.Add(mBase, uint32(v3)))
	if v4 == int64(0) {
		v12 = int32(-1636608428)
		v50 = v12
		v52 = v12
		v55 = v12
	} else {
		v15 = base.I32_wrap_i64(v4)
		v20 = base.I32_wrap_i64(int64(base.Ui64(v4)>>(uint(int64(32))%64))) ^ int32(-415931063)
		v26 = v15 - v20 - int32(1636608428) ^ base.I32_rotl(v20, int32(6))
		v28 = v15 + int32(1021750440)
		v29 = v20 + v28
		v30 = v26 + v29
		v34 = v28 - v26 ^ base.I32_rotl(v26, int32(8))
		v38 = v29 - v34 ^ base.I32_rotl(v34, int32(16))
		v42 = v30 - v38 ^ base.I32_rotl(v38, int32(19))
		v43 = v34 + v30
		v44 = v38 + v43
		v50 = v42 + v44
		v52 = v44
		v55 = v43 - v42 ^ base.I32_rotl(v42, int32(4))
	}
	v57 = int32(14)
	v59 = v50 ^ v55 - base.I32_rotl(v50, v57)
	v64 = v59 ^ (v2 + v52) - base.I32_rotl(v59, int32(11))
	v68 = v64 ^ v50 - base.I32_rotl(v64, int32(25))
	v72 = v68 ^ v59 - base.I32_rotl(v68, int32(16))
	v76 = v72 ^ v64 - base.I32_rotl(v72, int32(4))
	v80 = v76 ^ v68 - base.I32_rotl(v76, v57)
	v90 = F_Int64GetDatum(m, base.I64_extend_i32_u(v80)<<(uint(int64(32))%64)|base.I64_extend_i32_u(v80^v72-base.I32_rotl(v80, int32(24))))
	mBase = m.M
	v93 = m.ExcPending
	if v93 != 0 {
		return int32(0)
	} else {
		return v90
	}
}
func F_hashint8extended(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v3 int32
	_ = v3
	var v4 int64
	_ = v4
	var v12 int32
	_ = v12
	var v13 int64
	_ = v13
	var v21 int32
	_ = v21
	var v24 int32
	_ = v24
	var v29 int32
	_ = v29
	var v35 int32
	_ = v35
	var v37 int32
	_ = v37
	var v38 int32
	_ = v38
	var v39 int32
	_ = v39
	var v43 int32
	_ = v43
	var v47 int32
	_ = v47
	var v51 int32
	_ = v51
	var v52 int32
	_ = v52
	var v53 int32
	_ = v53
	var v59 int32
	_ = v59
	var v61 int32
	_ = v61
	var v64 int32
	_ = v64
	var v66 int32
	_ = v66
	var v68 int32
	_ = v68
	var v73 int32
	_ = v73
	var v77 int32
	_ = v77
	var v81 int32
	_ = v81
	var v85 int32
	_ = v85
	var v89 int32
	_ = v89
	var v99 int32
	_ = v99
	var v102 int32
	_ = v102
	v3 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v4 = *(*int64)(unsafe.Add(mBase, uint32(v3)))
	v12 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
	v13 = *(*int64)(unsafe.Add(mBase, uint32(v12)))
	if v13 == int64(0) {
		v21 = int32(-1636608428)
		v59 = v21
		v61 = v21
		v64 = v21
	} else {
		v24 = base.I32_wrap_i64(v13)
		v29 = base.I32_wrap_i64(int64(base.Ui64(v13)>>(uint(int64(32))%64))) ^ int32(-415931063)
		v35 = v24 - v29 - int32(1636608428) ^ base.I32_rotl(v29, int32(6))
		v37 = v24 + int32(1021750440)
		v38 = v29 + v37
		v39 = v35 + v38
		v43 = v37 - v35 ^ base.I32_rotl(v35, int32(8))
		v47 = v38 - v43 ^ base.I32_rotl(v43, int32(16))
		v51 = v39 - v47 ^ base.I32_rotl(v47, int32(19))
		v52 = v43 + v39
		v53 = v47 + v52
		v59 = v51 + v53
		v61 = v53
		v64 = v52 - v51 ^ base.I32_rotl(v51, int32(4))
	}
	v66 = int32(14)
	v68 = v59 ^ v64 - base.I32_rotl(v59, v66)
	v73 = v68 ^ (base.I32_wrap_i64(v4>>(uint(int64(63))%64)^int64(base.Ui64(v4)>>(uint(int64(32))%64))^v4) + v61) - base.I32_rotl(v68, int32(11))
	v77 = v73 ^ v59 - base.I32_rotl(v73, int32(25))
	v81 = v77 ^ v68 - base.I32_rotl(v77, int32(16))
	v85 = v81 ^ v73 - base.I32_rotl(v81, int32(4))
	v89 = v85 ^ v77 - base.I32_rotl(v85, v66)
	v99 = F_Int64GetDatum(m, base.I64_extend_i32_u(v89)<<(uint(int64(32))%64)|base.I64_extend_i32_u(v89^v81-base.I32_rotl(v89, int32(24))))
	mBase = m.M
	v102 = m.ExcPending
	if v102 != 0 {
		return int32(0)
	} else {
		return v99
	}
}
func F_hashmacaddr(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v2 int32
	_ = v2
	var v9 int32
	_ = v9
	var v72 int32
	_ = v72
	var v76 int32
	_ = v76
	var v77 int32
	_ = v77
	var v81 int32
	_ = v81
	var v82 int32
	_ = v82
	var v86 int32
	_ = v86
	var v87 int32
	_ = v87
	var v89 int32
	_ = v89
	var v91 int32
	_ = v91
	var v95 int32
	_ = v95
	var v96 int32
	_ = v96
	var v100 int32
	_ = v100
	var v101 int32
	_ = v101
	var v103 int32
	_ = v103
	var v104 int32
	_ = v104
	var v106 int32
	_ = v106
	var v110 int32
	_ = v110
	var v111 int32
	_ = v111
	var v115 int32
	_ = v115
	var v116 int32
	_ = v116
	var v174 int32
	_ = v174
	var v178 int32
	_ = v178
	var v179 int32
	_ = v179
	var v183 int32
	_ = v183
	var v184 int32
	_ = v184
	var v188 int32
	_ = v188
	var v189 int32
	_ = v189
	var v193 int32
	_ = v193
	var v194 int32
	_ = v194
	var v195 int32
	_ = v195
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
	var v210 int32
	_ = v210
	var v211 int32
	_ = v211
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
	var v223 int32
	_ = v223
	var v224 int32
	_ = v224
	var v225 int32
	_ = v225
	var v229 int32
	_ = v229
	var v230 int32
	_ = v230
	var v231 int32
	_ = v231
	var v232 int32
	_ = v232
	var v236 int32
	_ = v236
	var v237 int32
	_ = v237
	var v238 int32
	_ = v238
	var v241 int32
	_ = v241
	var v243 int32
	_ = v243
	var v247 int32
	_ = v247
	var v251 int32
	_ = v251
	var v255 int32
	_ = v255
	var v259 int32
	_ = v259
	var v263 int32
	_ = v263
	v2 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v9 = int32(-1636608426)
	if v2&int32(3) != 0 {
		switch int32(5) {
		case 0:
			v229 = v9
			v230 = v9
			v231 = v9
			v232 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v2))))
			v236 = v229 + v232
			v237 = v230
			v238 = v231
		case 1:
			v222 = v9
			v223 = v9
			v224 = v9
			v225 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v2)+1)))
			v229 = v225<<(uint(int32(8))%32) + v222
			v230 = v223
			v231 = v224
			v232 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v2))))
			v236 = v229 + v232
			v237 = v230
			v238 = v231
		case 2:
			v215 = v9
			v216 = v9
			v217 = v9
			v218 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v2)+2)))
			v222 = v218<<(uint(int32(16))%32) + v215
			v223 = v216
			v224 = v217
			v225 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v2)+1)))
			v229 = v225<<(uint(int32(8))%32) + v222
			v230 = v223
			v231 = v224
			v232 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v2))))
			v236 = v229 + v232
			v237 = v230
			v238 = v231
		case 3:
			v209 = v9
			v210 = v9
			v211 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v2)+3)))
			v215 = v211<<(uint(int32(24))%32) + v9
			v216 = v209
			v217 = v210
			v218 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v2)+2)))
			v222 = v218<<(uint(int32(16))%32) + v215
			v223 = v216
			v224 = v217
			v225 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v2)+1)))
			v229 = v225<<(uint(int32(8))%32) + v222
			v230 = v223
			v231 = v224
			v232 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v2))))
			v236 = v229 + v232
			v237 = v230
			v238 = v231
		case 4:
			v205 = v9
			v206 = v9
			v207 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v2)+4)))
			v209 = v205 + v207
			v210 = v206
			v211 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v2)+3)))
			v215 = v211<<(uint(int32(24))%32) + v9
			v216 = v209
			v217 = v210
			v218 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v2)+2)))
			v222 = v218<<(uint(int32(16))%32) + v215
			v223 = v216
			v224 = v217
			v225 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v2)+1)))
			v229 = v225<<(uint(int32(8))%32) + v222
			v230 = v223
			v231 = v224
			v232 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v2))))
			v236 = v229 + v232
			v237 = v230
			v238 = v231
		case 5:
			v199 = v9
			v200 = v9
			v201 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v2)+5)))
			v205 = v201<<(uint(int32(8))%32) + v199
			v206 = v200
			v207 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v2)+4)))
			v209 = v205 + v207
			v210 = v206
			v211 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v2)+3)))
			v215 = v211<<(uint(int32(24))%32) + v9
			v216 = v209
			v217 = v210
			v218 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v2)+2)))
			v222 = v218<<(uint(int32(16))%32) + v215
			v223 = v216
			v224 = v217
			v225 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v2)+1)))
			v229 = v225<<(uint(int32(8))%32) + v222
			v230 = v223
			v231 = v224
			v232 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v2))))
			v236 = v229 + v232
			v237 = v230
			v238 = v231
		case 6:
			v193 = v9
			v194 = v9
			v195 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v2)+6)))
			v199 = v195<<(uint(int32(16))%32) + v193
			v200 = v194
			v201 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v2)+5)))
			v205 = v201<<(uint(int32(8))%32) + v199
			v206 = v200
			v207 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v2)+4)))
			v209 = v205 + v207
			v210 = v206
			v211 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v2)+3)))
			v215 = v211<<(uint(int32(24))%32) + v9
			v216 = v209
			v217 = v210
			v218 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v2)+2)))
			v222 = v218<<(uint(int32(16))%32) + v215
			v223 = v216
			v224 = v217
			v225 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v2)+1)))
			v229 = v225<<(uint(int32(8))%32) + v222
			v230 = v223
			v231 = v224
			v232 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v2))))
			v236 = v229 + v232
			v237 = v230
			v238 = v231
		case 7:
			v188 = v9
			v189 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v2)+7)))
			v193 = v189<<(uint(int32(24))%32) + v9
			v194 = v188
			v195 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v2)+6)))
			v199 = v195<<(uint(int32(16))%32) + v193
			v200 = v194
			v201 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v2)+5)))
			v205 = v201<<(uint(int32(8))%32) + v199
			v206 = v200
			v207 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v2)+4)))
			v209 = v205 + v207
			v210 = v206
			v211 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v2)+3)))
			v215 = v211<<(uint(int32(24))%32) + v9
			v216 = v209
			v217 = v210
			v218 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v2)+2)))
			v222 = v218<<(uint(int32(16))%32) + v215
			v223 = v216
			v224 = v217
			v225 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v2)+1)))
			v229 = v225<<(uint(int32(8))%32) + v222
			v230 = v223
			v231 = v224
			v232 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v2))))
			v236 = v229 + v232
			v237 = v230
			v238 = v231
		case 8:
			v183 = v9
			v184 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v2)+8)))
			v188 = v184<<(uint(int32(8))%32) + v183
			v189 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v2)+7)))
			v193 = v189<<(uint(int32(24))%32) + v9
			v194 = v188
			v195 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v2)+6)))
			v199 = v195<<(uint(int32(16))%32) + v193
			v200 = v194
			v201 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v2)+5)))
			v205 = v201<<(uint(int32(8))%32) + v199
			v206 = v200
			v207 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v2)+4)))
			v209 = v205 + v207
			v210 = v206
			v211 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v2)+3)))
			v215 = v211<<(uint(int32(24))%32) + v9
			v216 = v209
			v217 = v210
			v218 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v2)+2)))
			v222 = v218<<(uint(int32(16))%32) + v215
			v223 = v216
			v224 = v217
			v225 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v2)+1)))
			v229 = v225<<(uint(int32(8))%32) + v222
			v230 = v223
			v231 = v224
			v232 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v2))))
			v236 = v229 + v232
			v237 = v230
			v238 = v231
		case 9:
			v178 = v9
			v179 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v2)+9)))
			v183 = v179<<(uint(int32(16))%32) + v178
			v184 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v2)+8)))
			v188 = v184<<(uint(int32(8))%32) + v183
			v189 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v2)+7)))
			v193 = v189<<(uint(int32(24))%32) + v9
			v194 = v188
			v195 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v2)+6)))
			v199 = v195<<(uint(int32(16))%32) + v193
			v200 = v194
			v201 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v2)+5)))
			v205 = v201<<(uint(int32(8))%32) + v199
			v206 = v200
			v207 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v2)+4)))
			v209 = v205 + v207
			v210 = v206
			v211 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v2)+3)))
			v215 = v211<<(uint(int32(24))%32) + v9
			v216 = v209
			v217 = v210
			v218 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v2)+2)))
			v222 = v218<<(uint(int32(16))%32) + v215
			v223 = v216
			v224 = v217
			v225 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v2)+1)))
			v229 = v225<<(uint(int32(8))%32) + v222
			v230 = v223
			v231 = v224
			v232 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v2))))
			v236 = v229 + v232
			v237 = v230
			v238 = v231
		case 10:
			v174 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v2)+10)))
			v178 = v174<<(uint(int32(24))%32) + v9
			v179 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v2)+9)))
			v183 = v179<<(uint(int32(16))%32) + v178
			v184 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v2)+8)))
			v188 = v184<<(uint(int32(8))%32) + v183
			v189 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v2)+7)))
			v193 = v189<<(uint(int32(24))%32) + v9
			v194 = v188
			v195 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v2)+6)))
			v199 = v195<<(uint(int32(16))%32) + v193
			v200 = v194
			v201 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v2)+5)))
			v205 = v201<<(uint(int32(8))%32) + v199
			v206 = v200
			v207 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v2)+4)))
			v209 = v205 + v207
			v210 = v206
			v211 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v2)+3)))
			v215 = v211<<(uint(int32(24))%32) + v9
			v216 = v209
			v217 = v210
			v218 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v2)+2)))
			v222 = v218<<(uint(int32(16))%32) + v215
			v223 = v216
			v224 = v217
			v225 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v2)+1)))
			v229 = v225<<(uint(int32(8))%32) + v222
			v230 = v223
			v231 = v224
			v232 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v2))))
			v236 = v229 + v232
			v237 = v230
			v238 = v231
		default:
			v236 = v9
			v237 = v9
			v238 = v9
		}
	} else {
		switch int32(5) {
		case 0:
			v115 = v9
			v116 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v2))))
			v236 = v115 + v116
			v237 = v9
			v238 = v9
		case 1:
			v110 = v9
			v111 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v2)+1)))
			v115 = v111<<(uint(int32(8))%32) + v110
			v116 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v2))))
			v236 = v115 + v116
			v237 = v9
			v238 = v9
		case 2:
			v106 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v2)+2)))
			v110 = v106<<(uint(int32(16))%32) + v9
			v111 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v2)+1)))
			v115 = v111<<(uint(int32(8))%32) + v110
			v116 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v2))))
			v236 = v115 + v116
			v237 = v9
			v238 = v9
		case 3:
			v103 = v9
			v104 = *(*int32)(unsafe.Add(mBase, uint32(v2)))
			v236 = v104 + v9
			v237 = v103
			v238 = v9
		case 4:
			v100 = v9
			v101 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v2)+4)))
			v103 = v100 + v101
			v104 = *(*int32)(unsafe.Add(mBase, uint32(v2)))
			v236 = v104 + v9
			v237 = v103
			v238 = v9
		case 5:
			v95 = v9
			v96 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v2)+5)))
			v100 = v96<<(uint(int32(8))%32) + v95
			v101 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v2)+4)))
			v103 = v100 + v101
			v104 = *(*int32)(unsafe.Add(mBase, uint32(v2)))
			v236 = v104 + v9
			v237 = v103
			v238 = v9
		case 6:
			v91 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v2)+6)))
			v95 = v91<<(uint(int32(16))%32) + v9
			v96 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v2)+5)))
			v100 = v96<<(uint(int32(8))%32) + v95
			v101 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v2)+4)))
			v103 = v100 + v101
			v104 = *(*int32)(unsafe.Add(mBase, uint32(v2)))
			v236 = v104 + v9
			v237 = v103
			v238 = v9
		case 7:
			v86 = v9
			v87 = *(*int32)(unsafe.Add(mBase, uint32(v2)))
			v89 = *(*int32)(unsafe.Add(mBase, uint32(v2)+4))
			v236 = v87 + v9
			v237 = v89 + v9
			v238 = v86
		case 8:
			v81 = v9
			v82 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v2)+8)))
			v86 = v82<<(uint(int32(8))%32) + v81
			v87 = *(*int32)(unsafe.Add(mBase, uint32(v2)))
			v89 = *(*int32)(unsafe.Add(mBase, uint32(v2)+4))
			v236 = v87 + v9
			v237 = v89 + v9
			v238 = v86
		case 9:
			v76 = v9
			v77 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v2)+9)))
			v81 = v77<<(uint(int32(16))%32) + v76
			v82 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v2)+8)))
			v86 = v82<<(uint(int32(8))%32) + v81
			v87 = *(*int32)(unsafe.Add(mBase, uint32(v2)))
			v89 = *(*int32)(unsafe.Add(mBase, uint32(v2)+4))
			v236 = v87 + v9
			v237 = v89 + v9
			v238 = v86
		case 10:
			v72 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v2)+10)))
			v76 = v72<<(uint(int32(24))%32) + v9
			v77 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v2)+9)))
			v81 = v77<<(uint(int32(16))%32) + v76
			v82 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v2)+8)))
			v86 = v82<<(uint(int32(8))%32) + v81
			v87 = *(*int32)(unsafe.Add(mBase, uint32(v2)))
			v89 = *(*int32)(unsafe.Add(mBase, uint32(v2)+4))
			v236 = v87 + v9
			v237 = v89 + v9
			v238 = v86
		default:
			v236 = v9
			v237 = v9
			v238 = v9
		}
	}
	v241 = int32(14)
	v243 = v237 ^ v238 - base.I32_rotl(v237, v241)
	v247 = v243 ^ v236 - base.I32_rotl(v243, int32(11))
	v251 = v247 ^ v237 - base.I32_rotl(v247, int32(25))
	v255 = v251 ^ v243 - base.I32_rotl(v251, int32(16))
	v259 = v255 ^ v247 - base.I32_rotl(v255, int32(4))
	v263 = v259 ^ v251 - base.I32_rotl(v259, v241)
	return v263 ^ v255 - base.I32_rotl(v263, int32(24))
}
func F_hashmacaddr8(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v2 int32
	_ = v2
	var v9 int32
	_ = v9
	var v72 int32
	_ = v72
	var v76 int32
	_ = v76
	var v77 int32
	_ = v77
	var v81 int32
	_ = v81
	var v82 int32
	_ = v82
	var v86 int32
	_ = v86
	var v87 int32
	_ = v87
	var v89 int32
	_ = v89
	var v91 int32
	_ = v91
	var v95 int32
	_ = v95
	var v96 int32
	_ = v96
	var v100 int32
	_ = v100
	var v101 int32
	_ = v101
	var v103 int32
	_ = v103
	var v104 int32
	_ = v104
	var v106 int32
	_ = v106
	var v110 int32
	_ = v110
	var v111 int32
	_ = v111
	var v115 int32
	_ = v115
	var v116 int32
	_ = v116
	var v174 int32
	_ = v174
	var v178 int32
	_ = v178
	var v179 int32
	_ = v179
	var v183 int32
	_ = v183
	var v184 int32
	_ = v184
	var v188 int32
	_ = v188
	var v189 int32
	_ = v189
	var v193 int32
	_ = v193
	var v194 int32
	_ = v194
	var v195 int32
	_ = v195
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
	var v210 int32
	_ = v210
	var v211 int32
	_ = v211
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
	var v223 int32
	_ = v223
	var v224 int32
	_ = v224
	var v225 int32
	_ = v225
	var v229 int32
	_ = v229
	var v230 int32
	_ = v230
	var v231 int32
	_ = v231
	var v232 int32
	_ = v232
	var v236 int32
	_ = v236
	var v237 int32
	_ = v237
	var v238 int32
	_ = v238
	var v241 int32
	_ = v241
	var v243 int32
	_ = v243
	var v247 int32
	_ = v247
	var v251 int32
	_ = v251
	var v255 int32
	_ = v255
	var v259 int32
	_ = v259
	var v263 int32
	_ = v263
	v2 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v9 = int32(-1636608424)
	if v2&int32(3) != 0 {
		switch int32(7) {
		case 0:
			v229 = v9
			v230 = v9
			v231 = v9
			v232 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v2))))
			v236 = v229 + v232
			v237 = v230
			v238 = v231
		case 1:
			v222 = v9
			v223 = v9
			v224 = v9
			v225 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v2)+1)))
			v229 = v225<<(uint(int32(8))%32) + v222
			v230 = v223
			v231 = v224
			v232 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v2))))
			v236 = v229 + v232
			v237 = v230
			v238 = v231
		case 2:
			v215 = v9
			v216 = v9
			v217 = v9
			v218 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v2)+2)))
			v222 = v218<<(uint(int32(16))%32) + v215
			v223 = v216
			v224 = v217
			v225 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v2)+1)))
			v229 = v225<<(uint(int32(8))%32) + v222
			v230 = v223
			v231 = v224
			v232 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v2))))
			v236 = v229 + v232
			v237 = v230
			v238 = v231
		case 3:
			v209 = v9
			v210 = v9
			v211 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v2)+3)))
			v215 = v211<<(uint(int32(24))%32) + v9
			v216 = v209
			v217 = v210
			v218 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v2)+2)))
			v222 = v218<<(uint(int32(16))%32) + v215
			v223 = v216
			v224 = v217
			v225 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v2)+1)))
			v229 = v225<<(uint(int32(8))%32) + v222
			v230 = v223
			v231 = v224
			v232 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v2))))
			v236 = v229 + v232
			v237 = v230
			v238 = v231
		case 4:
			v205 = v9
			v206 = v9
			v207 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v2)+4)))
			v209 = v205 + v207
			v210 = v206
			v211 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v2)+3)))
			v215 = v211<<(uint(int32(24))%32) + v9
			v216 = v209
			v217 = v210
			v218 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v2)+2)))
			v222 = v218<<(uint(int32(16))%32) + v215
			v223 = v216
			v224 = v217
			v225 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v2)+1)))
			v229 = v225<<(uint(int32(8))%32) + v222
			v230 = v223
			v231 = v224
			v232 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v2))))
			v236 = v229 + v232
			v237 = v230
			v238 = v231
		case 5:
			v199 = v9
			v200 = v9
			v201 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v2)+5)))
			v205 = v201<<(uint(int32(8))%32) + v199
			v206 = v200
			v207 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v2)+4)))
			v209 = v205 + v207
			v210 = v206
			v211 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v2)+3)))
			v215 = v211<<(uint(int32(24))%32) + v9
			v216 = v209
			v217 = v210
			v218 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v2)+2)))
			v222 = v218<<(uint(int32(16))%32) + v215
			v223 = v216
			v224 = v217
			v225 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v2)+1)))
			v229 = v225<<(uint(int32(8))%32) + v222
			v230 = v223
			v231 = v224
			v232 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v2))))
			v236 = v229 + v232
			v237 = v230
			v238 = v231
		case 6:
			v193 = v9
			v194 = v9
			v195 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v2)+6)))
			v199 = v195<<(uint(int32(16))%32) + v193
			v200 = v194
			v201 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v2)+5)))
			v205 = v201<<(uint(int32(8))%32) + v199
			v206 = v200
			v207 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v2)+4)))
			v209 = v205 + v207
			v210 = v206
			v211 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v2)+3)))
			v215 = v211<<(uint(int32(24))%32) + v9
			v216 = v209
			v217 = v210
			v218 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v2)+2)))
			v222 = v218<<(uint(int32(16))%32) + v215
			v223 = v216
			v224 = v217
			v225 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v2)+1)))
			v229 = v225<<(uint(int32(8))%32) + v222
			v230 = v223
			v231 = v224
			v232 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v2))))
			v236 = v229 + v232
			v237 = v230
			v238 = v231
		case 7:
			v188 = v9
			v189 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v2)+7)))
			v193 = v189<<(uint(int32(24))%32) + v9
			v194 = v188
			v195 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v2)+6)))
			v199 = v195<<(uint(int32(16))%32) + v193
			v200 = v194
			v201 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v2)+5)))
			v205 = v201<<(uint(int32(8))%32) + v199
			v206 = v200
			v207 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v2)+4)))
			v209 = v205 + v207
			v210 = v206
			v211 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v2)+3)))
			v215 = v211<<(uint(int32(24))%32) + v9
			v216 = v209
			v217 = v210
			v218 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v2)+2)))
			v222 = v218<<(uint(int32(16))%32) + v215
			v223 = v216
			v224 = v217
			v225 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v2)+1)))
			v229 = v225<<(uint(int32(8))%32) + v222
			v230 = v223
			v231 = v224
			v232 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v2))))
			v236 = v229 + v232
			v237 = v230
			v238 = v231
		case 8:
			v183 = v9
			v184 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v2)+8)))
			v188 = v184<<(uint(int32(8))%32) + v183
			v189 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v2)+7)))
			v193 = v189<<(uint(int32(24))%32) + v9
			v194 = v188
			v195 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v2)+6)))
			v199 = v195<<(uint(int32(16))%32) + v193
			v200 = v194
			v201 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v2)+5)))
			v205 = v201<<(uint(int32(8))%32) + v199
			v206 = v200
			v207 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v2)+4)))
			v209 = v205 + v207
			v210 = v206
			v211 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v2)+3)))
			v215 = v211<<(uint(int32(24))%32) + v9
			v216 = v209
			v217 = v210
			v218 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v2)+2)))
			v222 = v218<<(uint(int32(16))%32) + v215
			v223 = v216
			v224 = v217
			v225 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v2)+1)))
			v229 = v225<<(uint(int32(8))%32) + v222
			v230 = v223
			v231 = v224
			v232 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v2))))
			v236 = v229 + v232
			v237 = v230
			v238 = v231
		case 9:
			v178 = v9
			v179 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v2)+9)))
			v183 = v179<<(uint(int32(16))%32) + v178
			v184 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v2)+8)))
			v188 = v184<<(uint(int32(8))%32) + v183
			v189 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v2)+7)))
			v193 = v189<<(uint(int32(24))%32) + v9
			v194 = v188
			v195 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v2)+6)))
			v199 = v195<<(uint(int32(16))%32) + v193
			v200 = v194
			v201 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v2)+5)))
			v205 = v201<<(uint(int32(8))%32) + v199
			v206 = v200
			v207 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v2)+4)))
			v209 = v205 + v207
			v210 = v206
			v211 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v2)+3)))
			v215 = v211<<(uint(int32(24))%32) + v9
			v216 = v209
			v217 = v210
			v218 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v2)+2)))
			v222 = v218<<(uint(int32(16))%32) + v215
			v223 = v216
			v224 = v217
			v225 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v2)+1)))
			v229 = v225<<(uint(int32(8))%32) + v222
			v230 = v223
			v231 = v224
			v232 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v2))))
			v236 = v229 + v232
			v237 = v230
			v238 = v231
		case 10:
			v174 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v2)+10)))
			v178 = v174<<(uint(int32(24))%32) + v9
			v179 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v2)+9)))
			v183 = v179<<(uint(int32(16))%32) + v178
			v184 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v2)+8)))
			v188 = v184<<(uint(int32(8))%32) + v183
			v189 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v2)+7)))
			v193 = v189<<(uint(int32(24))%32) + v9
			v194 = v188
			v195 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v2)+6)))
			v199 = v195<<(uint(int32(16))%32) + v193
			v200 = v194
			v201 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v2)+5)))
			v205 = v201<<(uint(int32(8))%32) + v199
			v206 = v200
			v207 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v2)+4)))
			v209 = v205 + v207
			v210 = v206
			v211 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v2)+3)))
			v215 = v211<<(uint(int32(24))%32) + v9
			v216 = v209
			v217 = v210
			v218 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v2)+2)))
			v222 = v218<<(uint(int32(16))%32) + v215
			v223 = v216
			v224 = v217
			v225 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v2)+1)))
			v229 = v225<<(uint(int32(8))%32) + v222
			v230 = v223
			v231 = v224
			v232 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v2))))
			v236 = v229 + v232
			v237 = v230
			v238 = v231
		default:
			v236 = v9
			v237 = v9
			v238 = v9
		}
	} else {
		switch int32(7) {
		case 0:
			v115 = v9
			v116 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v2))))
			v236 = v115 + v116
			v237 = v9
			v238 = v9
		case 1:
			v110 = v9
			v111 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v2)+1)))
			v115 = v111<<(uint(int32(8))%32) + v110
			v116 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v2))))
			v236 = v115 + v116
			v237 = v9
			v238 = v9
		case 2:
			v106 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v2)+2)))
			v110 = v106<<(uint(int32(16))%32) + v9
			v111 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v2)+1)))
			v115 = v111<<(uint(int32(8))%32) + v110
			v116 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v2))))
			v236 = v115 + v116
			v237 = v9
			v238 = v9
		case 3:
			v103 = v9
			v104 = *(*int32)(unsafe.Add(mBase, uint32(v2)))
			v236 = v104 + v9
			v237 = v103
			v238 = v9
		case 4:
			v100 = v9
			v101 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v2)+4)))
			v103 = v100 + v101
			v104 = *(*int32)(unsafe.Add(mBase, uint32(v2)))
			v236 = v104 + v9
			v237 = v103
			v238 = v9
		case 5:
			v95 = v9
			v96 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v2)+5)))
			v100 = v96<<(uint(int32(8))%32) + v95
			v101 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v2)+4)))
			v103 = v100 + v101
			v104 = *(*int32)(unsafe.Add(mBase, uint32(v2)))
			v236 = v104 + v9
			v237 = v103
			v238 = v9
		case 6:
			v91 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v2)+6)))
			v95 = v91<<(uint(int32(16))%32) + v9
			v96 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v2)+5)))
			v100 = v96<<(uint(int32(8))%32) + v95
			v101 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v2)+4)))
			v103 = v100 + v101
			v104 = *(*int32)(unsafe.Add(mBase, uint32(v2)))
			v236 = v104 + v9
			v237 = v103
			v238 = v9
		case 7:
			v86 = v9
			v87 = *(*int32)(unsafe.Add(mBase, uint32(v2)))
			v89 = *(*int32)(unsafe.Add(mBase, uint32(v2)+4))
			v236 = v87 + v9
			v237 = v89 + v9
			v238 = v86
		case 8:
			v81 = v9
			v82 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v2)+8)))
			v86 = v82<<(uint(int32(8))%32) + v81
			v87 = *(*int32)(unsafe.Add(mBase, uint32(v2)))
			v89 = *(*int32)(unsafe.Add(mBase, uint32(v2)+4))
			v236 = v87 + v9
			v237 = v89 + v9
			v238 = v86
		case 9:
			v76 = v9
			v77 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v2)+9)))
			v81 = v77<<(uint(int32(16))%32) + v76
			v82 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v2)+8)))
			v86 = v82<<(uint(int32(8))%32) + v81
			v87 = *(*int32)(unsafe.Add(mBase, uint32(v2)))
			v89 = *(*int32)(unsafe.Add(mBase, uint32(v2)+4))
			v236 = v87 + v9
			v237 = v89 + v9
			v238 = v86
		case 10:
			v72 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v2)+10)))
			v76 = v72<<(uint(int32(24))%32) + v9
			v77 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v2)+9)))
			v81 = v77<<(uint(int32(16))%32) + v76
			v82 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v2)+8)))
			v86 = v82<<(uint(int32(8))%32) + v81
			v87 = *(*int32)(unsafe.Add(mBase, uint32(v2)))
			v89 = *(*int32)(unsafe.Add(mBase, uint32(v2)+4))
			v236 = v87 + v9
			v237 = v89 + v9
			v238 = v86
		default:
			v236 = v9
			v237 = v9
			v238 = v9
		}
	}
	v241 = int32(14)
	v243 = v237 ^ v238 - base.I32_rotl(v237, v241)
	v247 = v243 ^ v236 - base.I32_rotl(v243, int32(11))
	v251 = v247 ^ v237 - base.I32_rotl(v247, int32(25))
	v255 = v251 ^ v243 - base.I32_rotl(v251, int32(16))
	v259 = v255 ^ v247 - base.I32_rotl(v255, int32(4))
	v263 = v259 ^ v251 - base.I32_rotl(v259, v241)
	return v263 ^ v255 - base.I32_rotl(v263, int32(24))
}
func F_hashoptions(m *base.Module, l0 int32, l1 int32) int32 {
	var v3 int32
	_ = v3
	var v7 int32
	_ = v7
	var v10 int32
	_ = v10
	v3 = int32(8)
	v7 = F_build_reloptions(m, l0, l1, v3, v3, int32(_a_F_hashoptions_0), int32(1))
	v10 = m.ExcPending
	if v10 != 0 {
		return int32(0)
	} else {
		return v7
	}
}
func F_hashrescan(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32) {
	mBase := m.M
	_ = mBase
	var v7 int32
	_ = v7
	var v8 int32
	_ = v8
	var v11 int32
	_ = v11
	var v15 int32
	_ = v15
	var v17 int32
	_ = v17
	var v18 int32
	_ = v18
	var v28 int32
	_ = v28
	var v31 int32
	_ = v31
	var v33 int32
	_ = v33
	var v34 int32
	_ = v34
	var v37 int32
	_ = v37
	v7 = *(*int32)(unsafe.Add(mBase, uint32(l0)+36))
	v8 = *(*int32)(unsafe.Add(mBase, uint32(v7)+28))
	if v8 == int32(-1) {
		F__hash_dropscanbuf(m, v7)
		mBase = m.M
		v17 = m.ExcPending
		if v17 != 0 {
			return
		} else {
			v18 = int32(0)
			*(*int32)(unsafe.Add(mBase, uint32(v7)+48)) = v18
			*(*int64)(unsafe.Add(mBase, uint32(v7)+40)) = int64(0)
			*(*int64)(unsafe.Add(mBase, uint32(v7)+32)) = int64(-1)
			*(*int64)(unsafe.Add(mBase, uint32(v7)+24)) = int64(-4294967296)
			if l1 == v18 {
			} else {
				v28 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
				if v28 <= int32(0) {
				} else {
					v31 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
					v33 = v28 * int32(48)
					if v33 != 0 {
						v34 = F__emscripten_memcpy_bulkmem(m, v31, l1, v33)
						mBase = m.M
					} else {
					}
				}
			}
			v37 = int32(0)
			*(*uint16)(unsafe.Add(mBase, uint32(v7)+12)) = uint16(v37)
			return
		}
	} else {
		v11 = *(*int32)(unsafe.Add(mBase, uint32(v7)+20))
		if v11 <= int32(0) {
			F__hash_dropscanbuf(m, v7)
			mBase = m.M
			v17 = m.ExcPending
			if v17 != 0 {
				return
			} else {
				v18 = int32(0)
				*(*int32)(unsafe.Add(mBase, uint32(v7)+48)) = v18
				*(*int64)(unsafe.Add(mBase, uint32(v7)+40)) = int64(0)
				*(*int64)(unsafe.Add(mBase, uint32(v7)+32)) = int64(-1)
				*(*int64)(unsafe.Add(mBase, uint32(v7)+24)) = int64(-4294967296)
				if l1 == v18 {
				} else {
					v28 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
					if v28 <= int32(0) {
					} else {
						v31 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
						v33 = v28 * int32(48)
						if v33 != 0 {
							v34 = F__emscripten_memcpy_bulkmem(m, v31, l1, v33)
							mBase = m.M
						} else {
						}
					}
				}
				v37 = int32(0)
				*(*uint16)(unsafe.Add(mBase, uint32(v7)+12)) = uint16(v37)
				return
			}
		} else {
			F__hash_kill_items(m, l0)
			mBase = m.M
			v15 = m.ExcPending
			if v15 != 0 {
				return
			} else {
				F__hash_dropscanbuf(m, v7)
				mBase = m.M
				v17 = m.ExcPending
				if v17 != 0 {
					return
				} else {
					v18 = int32(0)
					*(*int32)(unsafe.Add(mBase, uint32(v7)+48)) = v18
					*(*int64)(unsafe.Add(mBase, uint32(v7)+40)) = int64(0)
					*(*int64)(unsafe.Add(mBase, uint32(v7)+32)) = int64(-1)
					*(*int64)(unsafe.Add(mBase, uint32(v7)+24)) = int64(-4294967296)
					if l1 == v18 {
					} else {
						v28 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
						if v28 <= int32(0) {
						} else {
							v31 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
							v33 = v28 * int32(48)
							if v33 != 0 {
								v34 = F__emscripten_memcpy_bulkmem(m, v31, l1, v33)
								mBase = m.M
							} else {
							}
						}
					}
					v37 = int32(0)
					*(*uint16)(unsafe.Add(mBase, uint32(v7)+12)) = uint16(v37)
					return
				}
			}
		}
	}
}
func F_hashtextextended(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
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
	var v17 int32
	_ = v17
	var v18 int32
	_ = v18
	var v21 int32
	_ = v21
	var v23 int32
	_ = v23
	var v24 int32
	_ = v24
	var v25 int32
	_ = v25
	var v30 int32
	_ = v30
	var v32 int32
	_ = v32
	var v41 int32
	_ = v41
	var v44 int32
	_ = v44
	var v45 int32
	_ = v45
	var v49 int32
	_ = v49
	var v55 int32
	_ = v55
	var v56 int32
	_ = v56
	var v57 int64
	_ = v57
	var v63 int32
	_ = v63
	var v67 int32
	_ = v67
	var v68 int32
	_ = v68
	var v72 int32
	_ = v72
	var v74 int32
	_ = v74
	var v78 int32
	_ = v78
	var v82 int32
	_ = v82
	var v83 int32
	_ = v83
	var v84 int32
	_ = v84
	var v85 int32
	_ = v85
	var v89 int32
	_ = v89
	var v93 int32
	_ = v93
	var v98 int32
	_ = v98
	var v100 int32
	_ = v100
	var v102 int32
	_ = v102
	var v104 int32
	_ = v104
	var v109 int32
	_ = v109
	var v110 int32
	_ = v110
	var v112 int32
	_ = v112
	var v113 int32
	_ = v113
	var v114 int32
	_ = v114
	var v116 int32
	_ = v116
	var v117 int32
	_ = v117
	var v118 int32
	_ = v118
	var v120 int32
	_ = v120
	var v121 int32
	_ = v121
	var v123 int32
	_ = v123
	var v125 int32
	_ = v125
	var v129 int32
	_ = v129
	var v130 int32
	_ = v130
	var v131 int32
	_ = v131
	var v132 int32
	_ = v132
	var v136 int32
	_ = v136
	var v140 int32
	_ = v140
	var v144 int32
	_ = v144
	var v145 int32
	_ = v145
	var v146 int32
	_ = v146
	var v147 int32
	_ = v147
	var v151 int32
	_ = v151
	var v152 int32
	_ = v152
	var v153 int32
	_ = v153
	var v155 int32
	_ = v155
	var v158 int32
	_ = v158
	var v159 int32
	_ = v159
	var v161 int32
	_ = v161
	var v162 int32
	_ = v162
	var v163 int32
	_ = v163
	var v169 int32
	_ = v169
	var v170 int32
	_ = v170
	var v172 int32
	_ = v172
	var v173 int32
	_ = v173
	var v174 int32
	_ = v174
	var v176 int32
	_ = v176
	var v177 int32
	_ = v177
	var v178 int32
	_ = v178
	var v180 int32
	_ = v180
	var v181 int32
	_ = v181
	var v183 int32
	_ = v183
	var v185 int32
	_ = v185
	var v189 int32
	_ = v189
	var v190 int32
	_ = v190
	var v191 int32
	_ = v191
	var v192 int32
	_ = v192
	var v196 int32
	_ = v196
	var v200 int32
	_ = v200
	var v204 int32
	_ = v204
	var v205 int32
	_ = v205
	var v206 int32
	_ = v206
	var v207 int32
	_ = v207
	var v211 int32
	_ = v211
	var v212 int32
	_ = v212
	var v213 int32
	_ = v213
	var v215 int32
	_ = v215
	var v218 int32
	_ = v218
	var v219 int32
	_ = v219
	var v221 int32
	_ = v221
	var v222 int32
	_ = v222
	var v223 int32
	_ = v223
	var v227 int32
	_ = v227
	var v231 int32
	_ = v231
	var v232 int32
	_ = v232
	var v236 int32
	_ = v236
	var v237 int32
	_ = v237
	var v241 int32
	_ = v241
	var v242 int32
	_ = v242
	var v244 int32
	_ = v244
	var v246 int32
	_ = v246
	var v250 int32
	_ = v250
	var v251 int32
	_ = v251
	var v255 int32
	_ = v255
	var v256 int32
	_ = v256
	var v258 int32
	_ = v258
	var v259 int32
	_ = v259
	var v261 int32
	_ = v261
	var v265 int32
	_ = v265
	var v266 int32
	_ = v266
	var v270 int32
	_ = v270
	var v271 int32
	_ = v271
	var v273 int32
	_ = v273
	var v277 int32
	_ = v277
	var v278 int32
	_ = v278
	var v282 int32
	_ = v282
	var v283 int32
	_ = v283
	var v287 int32
	_ = v287
	var v288 int32
	_ = v288
	var v292 int32
	_ = v292
	var v293 int32
	_ = v293
	var v294 int32
	_ = v294
	var v298 int32
	_ = v298
	var v299 int32
	_ = v299
	var v300 int32
	_ = v300
	var v304 int32
	_ = v304
	var v305 int32
	_ = v305
	var v306 int32
	_ = v306
	var v308 int32
	_ = v308
	var v309 int32
	_ = v309
	var v310 int32
	_ = v310
	var v314 int32
	_ = v314
	var v315 int32
	_ = v315
	var v316 int32
	_ = v316
	var v317 int32
	_ = v317
	var v321 int32
	_ = v321
	var v322 int32
	_ = v322
	var v323 int32
	_ = v323
	var v324 int32
	_ = v324
	var v328 int32
	_ = v328
	var v329 int32
	_ = v329
	var v330 int32
	_ = v330
	var v331 int32
	_ = v331
	var v336 int32
	_ = v336
	var v337 int32
	_ = v337
	var v338 int32
	_ = v338
	var v341 int32
	_ = v341
	var v343 int32
	_ = v343
	var v347 int32
	_ = v347
	var v351 int32
	_ = v351
	var v355 int32
	_ = v355
	var v359 int32
	_ = v359
	var v363 int32
	_ = v363
	var v373 int32
	_ = v373
	var v374 int32
	_ = v374
	var v375 int32
	_ = v375
	var v379 int32
	_ = v379
	var v381 int32
	_ = v381
	var v390 int32
	_ = v390
	var v393 int32
	_ = v393
	var v394 int32
	_ = v394
	var v398 int32
	_ = v398
	var v404 int32
	_ = v404
	var v405 int32
	_ = v405
	var v406 int32
	_ = v406
	var v408 int32
	_ = v408
	var v409 int32
	_ = v409
	var v410 int32
	_ = v410
	var v411 int32
	_ = v411
	var v412 int32
	_ = v412
	var v414 int32
	_ = v414
	var v415 int64
	_ = v415
	var v421 int32
	_ = v421
	var v425 int32
	_ = v425
	var v426 int32
	_ = v426
	var v430 int32
	_ = v430
	var v432 int32
	_ = v432
	var v436 int32
	_ = v436
	var v440 int32
	_ = v440
	var v441 int32
	_ = v441
	var v442 int32
	_ = v442
	var v443 int32
	_ = v443
	var v447 int32
	_ = v447
	var v451 int32
	_ = v451
	var v456 int32
	_ = v456
	var v458 int32
	_ = v458
	var v460 int32
	_ = v460
	var v462 int32
	_ = v462
	var v467 int32
	_ = v467
	var v468 int32
	_ = v468
	var v470 int32
	_ = v470
	var v471 int32
	_ = v471
	var v472 int32
	_ = v472
	var v474 int32
	_ = v474
	var v475 int32
	_ = v475
	var v476 int32
	_ = v476
	var v478 int32
	_ = v478
	var v479 int32
	_ = v479
	var v481 int32
	_ = v481
	var v483 int32
	_ = v483
	var v487 int32
	_ = v487
	var v488 int32
	_ = v488
	var v489 int32
	_ = v489
	var v490 int32
	_ = v490
	var v494 int32
	_ = v494
	var v498 int32
	_ = v498
	var v502 int32
	_ = v502
	var v503 int32
	_ = v503
	var v504 int32
	_ = v504
	var v505 int32
	_ = v505
	var v509 int32
	_ = v509
	var v510 int32
	_ = v510
	var v511 int32
	_ = v511
	var v513 int32
	_ = v513
	var v516 int32
	_ = v516
	var v517 int32
	_ = v517
	var v519 int32
	_ = v519
	var v520 int32
	_ = v520
	var v521 int32
	_ = v521
	var v527 int32
	_ = v527
	var v528 int32
	_ = v528
	var v530 int32
	_ = v530
	var v531 int32
	_ = v531
	var v532 int32
	_ = v532
	var v534 int32
	_ = v534
	var v535 int32
	_ = v535
	var v536 int32
	_ = v536
	var v538 int32
	_ = v538
	var v539 int32
	_ = v539
	var v541 int32
	_ = v541
	var v543 int32
	_ = v543
	var v547 int32
	_ = v547
	var v548 int32
	_ = v548
	var v549 int32
	_ = v549
	var v550 int32
	_ = v550
	var v554 int32
	_ = v554
	var v558 int32
	_ = v558
	var v562 int32
	_ = v562
	var v563 int32
	_ = v563
	var v564 int32
	_ = v564
	var v565 int32
	_ = v565
	var v569 int32
	_ = v569
	var v570 int32
	_ = v570
	var v571 int32
	_ = v571
	var v573 int32
	_ = v573
	var v576 int32
	_ = v576
	var v577 int32
	_ = v577
	var v579 int32
	_ = v579
	var v580 int32
	_ = v580
	var v581 int32
	_ = v581
	var v585 int32
	_ = v585
	var v589 int32
	_ = v589
	var v590 int32
	_ = v590
	var v594 int32
	_ = v594
	var v595 int32
	_ = v595
	var v599 int32
	_ = v599
	var v600 int32
	_ = v600
	var v602 int32
	_ = v602
	var v604 int32
	_ = v604
	var v608 int32
	_ = v608
	var v609 int32
	_ = v609
	var v613 int32
	_ = v613
	var v614 int32
	_ = v614
	var v616 int32
	_ = v616
	var v617 int32
	_ = v617
	var v619 int32
	_ = v619
	var v623 int32
	_ = v623
	var v624 int32
	_ = v624
	var v628 int32
	_ = v628
	var v629 int32
	_ = v629
	var v631 int32
	_ = v631
	var v635 int32
	_ = v635
	var v636 int32
	_ = v636
	var v640 int32
	_ = v640
	var v641 int32
	_ = v641
	var v645 int32
	_ = v645
	var v646 int32
	_ = v646
	var v650 int32
	_ = v650
	var v651 int32
	_ = v651
	var v652 int32
	_ = v652
	var v656 int32
	_ = v656
	var v657 int32
	_ = v657
	var v658 int32
	_ = v658
	var v662 int32
	_ = v662
	var v663 int32
	_ = v663
	var v664 int32
	_ = v664
	var v666 int32
	_ = v666
	var v667 int32
	_ = v667
	var v668 int32
	_ = v668
	var v672 int32
	_ = v672
	var v673 int32
	_ = v673
	var v674 int32
	_ = v674
	var v675 int32
	_ = v675
	var v679 int32
	_ = v679
	var v680 int32
	_ = v680
	var v681 int32
	_ = v681
	var v682 int32
	_ = v682
	var v686 int32
	_ = v686
	var v687 int32
	_ = v687
	var v688 int32
	_ = v688
	var v689 int32
	_ = v689
	var v694 int32
	_ = v694
	var v695 int32
	_ = v695
	var v696 int32
	_ = v696
	var v699 int32
	_ = v699
	var v701 int32
	_ = v701
	var v705 int32
	_ = v705
	var v709 int32
	_ = v709
	var v713 int32
	_ = v713
	var v717 int32
	_ = v717
	var v721 int32
	_ = v721
	var v731 int32
	_ = v731
	var v732 int32
	_ = v732
	var v734 int32
	_ = v734
	var v735 int32
	_ = v735
	var v740 int32
	_ = v740
	var v743 int32
	_ = v743
	var v748 int32
	_ = v748
	var v751 int32
	_ = v751
	var v755 int32
	_ = v755
	var v759 int32
	_ = v759
	var v764 int32
	_ = v764
	var v768 int32
	_ = v768
	var v772 int32
	_ = v772
	var v777 int32
	_ = v777
	v9 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v10 = F_pg_detoast_datum_packed(m, v9)
	mBase = m.M
	v13 = m.ExcPending
	if v13 != 0 {
		return int32(0)
	} else {
		v14 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
		if v14 != 0 {
			v15 = F_pg_newlocale_from_collation(m, v14)
			mBase = m.M
			v16 = m.ExcPending
			if v16 != 0 {
				return int32(0)
			} else {
				v17 = int32(1)
				v18 = v10 + v17
				v21 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v10))))
				v23 = v21 & v17
				if v23 != 0 {
					v24 = v18
				} else {
					v24 = v10 + int32(4)
				}
				v25 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v15)+1)))
				if v25 == int32(1) {
					if v21 == int32(1) {
						v30 = int32(4)
						v32 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v18))))
						if v32&int32(254) == int32(2) {
							v41 = v30
						} else {
							v41 = base.B2i32(v32 == int32(18)) << (uint(v30) % 32)
						}
						if v32 == int32(1) {
							v44 = v30
						} else {
							v44 = v41
						}
						v55 = v44
					} else {
						v45 = int32(1)
						if v23 != 0 {
							v55 = int32(base.Ui32(v21)>>(uint(v45)%32)) - v45
						} else {
							v49 = *(*int32)(unsafe.Add(mBase, uint32(v10)))
							v55 = int32(base.Ui32(v49)>>(uint(int32(2))%32)) - int32(4)
						}
					}
					v56 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
					v57 = *(*int64)(unsafe.Add(mBase, uint32(v56)))
					v63 = v55 - int32(1636608432)
					if v57 == int64(0) {
						v100 = v63
						v102 = v63
						v104 = v63
					} else {
						v67 = v63 + base.I32_wrap_i64(v57)
						v68 = v67 + v63
						v72 = int32(4)
						v74 = base.I32_wrap_i64(int64(base.Ui64(v57)>>(uint(int64(32))%64))) ^ base.I32_rotl(v63, v72)
						v78 = v67 - v74 ^ base.I32_rotl(v74, int32(6))
						v82 = v68 - v78 ^ base.I32_rotl(v78, int32(8))
						v83 = v74 + v68
						v84 = v78 + v83
						v85 = v82 + v84
						v89 = v83 - v82 ^ base.I32_rotl(v82, int32(16))
						v93 = v84 - v89 ^ base.I32_rotl(v89, int32(19))
						v98 = v89 + v85
						v100 = v98
						v102 = v85 - v93 ^ base.I32_rotl(v93, v72)
						v104 = v93 + v98
					}
					if v24&int32(3) != 0 {
						if base.Ui32(int32(11)) < base.Ui32(v55) {
							v109 = v24
							v110 = v55
							v112 = v100
							v113 = v104
							v114 = v102
							for {
								v116 = *(*int32)(unsafe.Add(mBase, uint32(v109)+4))
								v117 = v116 + v113
								v118 = *(*int32)(unsafe.Add(mBase, uint32(v109)))
								v120 = *(*int32)(unsafe.Add(mBase, uint32(v109)+8))
								v121 = v120 + v114
								v123 = int32(4)
								v125 = v118 + v112 - v121 ^ base.I32_rotl(v121, v123)
								v129 = v117 - v125 ^ base.I32_rotl(v125, int32(6))
								v130 = v121 + v117
								v131 = v125 + v130
								v132 = v129 + v131
								v136 = v130 - v129 ^ base.I32_rotl(v129, int32(8))
								v140 = v131 - v136 ^ base.I32_rotl(v136, int32(16))
								v144 = v132 - v140 ^ base.I32_rotl(v140, int32(19))
								v145 = v136 + v132
								v146 = v140 + v145
								v147 = v144 + v146
								v151 = v145 - v144 ^ base.I32_rotl(v144, v123)
								v152 = int32(12)
								v153 = v109 + v152
								v155 = v110 - v152
								if base.Ui32(int32(11)) < base.Ui32(v155) {
									v109 = v153
									v110 = v155
									v112 = v146
									v113 = v147
									v114 = v151
									continue
								} else {
									break
								}
								break
							}
							v158 = v153
							v159 = v155
							v161 = v146
							v162 = v147
							v163 = v151
						} else {
							v158 = v24
							v159 = v55
							v161 = v100
							v162 = v104
							v163 = v102
						}
						switch v159 - int32(1) {
						case 0:
							v328 = v161
							v329 = v162
							v330 = v163
							v331 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v158))))
							v336 = v328 + v331
							v337 = v329
							v338 = v330
						case 1:
							v321 = v161
							v322 = v162
							v323 = v163
							v324 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v158)+1)))
							v328 = v324<<(uint(int32(8))%32) + v321
							v329 = v322
							v330 = v323
							v331 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v158))))
							v336 = v328 + v331
							v337 = v329
							v338 = v330
						case 2:
							v314 = v161
							v315 = v162
							v316 = v163
							v317 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v158)+2)))
							v321 = v317<<(uint(int32(16))%32) + v314
							v322 = v315
							v323 = v316
							v324 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v158)+1)))
							v328 = v324<<(uint(int32(8))%32) + v321
							v329 = v322
							v330 = v323
							v331 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v158))))
							v336 = v328 + v331
							v337 = v329
							v338 = v330
						case 3:
							v308 = v162
							v309 = v163
							v310 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v158)+3)))
							v314 = v310<<(uint(int32(24))%32) + v161
							v315 = v308
							v316 = v309
							v317 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v158)+2)))
							v321 = v317<<(uint(int32(16))%32) + v314
							v322 = v315
							v323 = v316
							v324 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v158)+1)))
							v328 = v324<<(uint(int32(8))%32) + v321
							v329 = v322
							v330 = v323
							v331 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v158))))
							v336 = v328 + v331
							v337 = v329
							v338 = v330
						case 4:
							v304 = v162
							v305 = v163
							v306 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v158)+4)))
							v308 = v304 + v306
							v309 = v305
							v310 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v158)+3)))
							v314 = v310<<(uint(int32(24))%32) + v161
							v315 = v308
							v316 = v309
							v317 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v158)+2)))
							v321 = v317<<(uint(int32(16))%32) + v314
							v322 = v315
							v323 = v316
							v324 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v158)+1)))
							v328 = v324<<(uint(int32(8))%32) + v321
							v329 = v322
							v330 = v323
							v331 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v158))))
							v336 = v328 + v331
							v337 = v329
							v338 = v330
						case 5:
							v298 = v162
							v299 = v163
							v300 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v158)+5)))
							v304 = v300<<(uint(int32(8))%32) + v298
							v305 = v299
							v306 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v158)+4)))
							v308 = v304 + v306
							v309 = v305
							v310 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v158)+3)))
							v314 = v310<<(uint(int32(24))%32) + v161
							v315 = v308
							v316 = v309
							v317 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v158)+2)))
							v321 = v317<<(uint(int32(16))%32) + v314
							v322 = v315
							v323 = v316
							v324 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v158)+1)))
							v328 = v324<<(uint(int32(8))%32) + v321
							v329 = v322
							v330 = v323
							v331 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v158))))
							v336 = v328 + v331
							v337 = v329
							v338 = v330
						case 6:
							v292 = v162
							v293 = v163
							v294 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v158)+6)))
							v298 = v294<<(uint(int32(16))%32) + v292
							v299 = v293
							v300 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v158)+5)))
							v304 = v300<<(uint(int32(8))%32) + v298
							v305 = v299
							v306 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v158)+4)))
							v308 = v304 + v306
							v309 = v305
							v310 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v158)+3)))
							v314 = v310<<(uint(int32(24))%32) + v161
							v315 = v308
							v316 = v309
							v317 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v158)+2)))
							v321 = v317<<(uint(int32(16))%32) + v314
							v322 = v315
							v323 = v316
							v324 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v158)+1)))
							v328 = v324<<(uint(int32(8))%32) + v321
							v329 = v322
							v330 = v323
							v331 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v158))))
							v336 = v328 + v331
							v337 = v329
							v338 = v330
						case 7:
							v287 = v163
							v288 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v158)+7)))
							v292 = v288<<(uint(int32(24))%32) + v162
							v293 = v287
							v294 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v158)+6)))
							v298 = v294<<(uint(int32(16))%32) + v292
							v299 = v293
							v300 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v158)+5)))
							v304 = v300<<(uint(int32(8))%32) + v298
							v305 = v299
							v306 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v158)+4)))
							v308 = v304 + v306
							v309 = v305
							v310 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v158)+3)))
							v314 = v310<<(uint(int32(24))%32) + v161
							v315 = v308
							v316 = v309
							v317 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v158)+2)))
							v321 = v317<<(uint(int32(16))%32) + v314
							v322 = v315
							v323 = v316
							v324 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v158)+1)))
							v328 = v324<<(uint(int32(8))%32) + v321
							v329 = v322
							v330 = v323
							v331 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v158))))
							v336 = v328 + v331
							v337 = v329
							v338 = v330
						case 8:
							v282 = v163
							v283 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v158)+8)))
							v287 = v283<<(uint(int32(8))%32) + v282
							v288 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v158)+7)))
							v292 = v288<<(uint(int32(24))%32) + v162
							v293 = v287
							v294 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v158)+6)))
							v298 = v294<<(uint(int32(16))%32) + v292
							v299 = v293
							v300 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v158)+5)))
							v304 = v300<<(uint(int32(8))%32) + v298
							v305 = v299
							v306 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v158)+4)))
							v308 = v304 + v306
							v309 = v305
							v310 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v158)+3)))
							v314 = v310<<(uint(int32(24))%32) + v161
							v315 = v308
							v316 = v309
							v317 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v158)+2)))
							v321 = v317<<(uint(int32(16))%32) + v314
							v322 = v315
							v323 = v316
							v324 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v158)+1)))
							v328 = v324<<(uint(int32(8))%32) + v321
							v329 = v322
							v330 = v323
							v331 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v158))))
							v336 = v328 + v331
							v337 = v329
							v338 = v330
						case 9:
							v277 = v163
							v278 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v158)+9)))
							v282 = v278<<(uint(int32(16))%32) + v277
							v283 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v158)+8)))
							v287 = v283<<(uint(int32(8))%32) + v282
							v288 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v158)+7)))
							v292 = v288<<(uint(int32(24))%32) + v162
							v293 = v287
							v294 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v158)+6)))
							v298 = v294<<(uint(int32(16))%32) + v292
							v299 = v293
							v300 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v158)+5)))
							v304 = v300<<(uint(int32(8))%32) + v298
							v305 = v299
							v306 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v158)+4)))
							v308 = v304 + v306
							v309 = v305
							v310 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v158)+3)))
							v314 = v310<<(uint(int32(24))%32) + v161
							v315 = v308
							v316 = v309
							v317 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v158)+2)))
							v321 = v317<<(uint(int32(16))%32) + v314
							v322 = v315
							v323 = v316
							v324 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v158)+1)))
							v328 = v324<<(uint(int32(8))%32) + v321
							v329 = v322
							v330 = v323
							v331 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v158))))
							v336 = v328 + v331
							v337 = v329
							v338 = v330
						case 10:
							v273 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v158)+10)))
							v277 = v273<<(uint(int32(24))%32) + v163
							v278 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v158)+9)))
							v282 = v278<<(uint(int32(16))%32) + v277
							v283 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v158)+8)))
							v287 = v283<<(uint(int32(8))%32) + v282
							v288 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v158)+7)))
							v292 = v288<<(uint(int32(24))%32) + v162
							v293 = v287
							v294 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v158)+6)))
							v298 = v294<<(uint(int32(16))%32) + v292
							v299 = v293
							v300 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v158)+5)))
							v304 = v300<<(uint(int32(8))%32) + v298
							v305 = v299
							v306 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v158)+4)))
							v308 = v304 + v306
							v309 = v305
							v310 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v158)+3)))
							v314 = v310<<(uint(int32(24))%32) + v161
							v315 = v308
							v316 = v309
							v317 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v158)+2)))
							v321 = v317<<(uint(int32(16))%32) + v314
							v322 = v315
							v323 = v316
							v324 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v158)+1)))
							v328 = v324<<(uint(int32(8))%32) + v321
							v329 = v322
							v330 = v323
							v331 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v158))))
							v336 = v328 + v331
							v337 = v329
							v338 = v330
						default:
							v336 = v161
							v337 = v162
							v338 = v163
						}
					} else {
						if base.Ui32(int32(12)) <= base.Ui32(v55) {
							v169 = v24
							v170 = v55
							v172 = v100
							v173 = v104
							v174 = v102
							for {
								v176 = *(*int32)(unsafe.Add(mBase, uint32(v169)+4))
								v177 = v176 + v173
								v178 = *(*int32)(unsafe.Add(mBase, uint32(v169)))
								v180 = *(*int32)(unsafe.Add(mBase, uint32(v169)+8))
								v181 = v180 + v174
								v183 = int32(4)
								v185 = v178 + v172 - v181 ^ base.I32_rotl(v181, v183)
								v189 = v177 - v185 ^ base.I32_rotl(v185, int32(6))
								v190 = v181 + v177
								v191 = v185 + v190
								v192 = v189 + v191
								v196 = v190 - v189 ^ base.I32_rotl(v189, int32(8))
								v200 = v191 - v196 ^ base.I32_rotl(v196, int32(16))
								v204 = v192 - v200 ^ base.I32_rotl(v200, int32(19))
								v205 = v196 + v192
								v206 = v200 + v205
								v207 = v204 + v206
								v211 = v205 - v204 ^ base.I32_rotl(v204, v183)
								v212 = int32(12)
								v213 = v169 + v212
								v215 = v170 - v212
								if base.Ui32(int32(11)) < base.Ui32(v215) {
									v169 = v213
									v170 = v215
									v172 = v206
									v173 = v207
									v174 = v211
									continue
								} else {
									break
								}
								break
							}
							v218 = v213
							v219 = v215
							v221 = v206
							v222 = v207
							v223 = v211
						} else {
							v218 = v24
							v219 = v55
							v221 = v100
							v222 = v104
							v223 = v102
						}
						switch v219 - int32(1) {
						case 0:
							v270 = v221
							v271 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v218))))
							v336 = v270 + v271
							v337 = v222
							v338 = v223
						case 1:
							v265 = v221
							v266 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v218)+1)))
							v270 = v266<<(uint(int32(8))%32) + v265
							v271 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v218))))
							v336 = v270 + v271
							v337 = v222
							v338 = v223
						case 2:
							v261 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v218)+2)))
							v265 = v261<<(uint(int32(16))%32) + v221
							v266 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v218)+1)))
							v270 = v266<<(uint(int32(8))%32) + v265
							v271 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v218))))
							v336 = v270 + v271
							v337 = v222
							v338 = v223
						case 3:
							v258 = v222
							v259 = *(*int32)(unsafe.Add(mBase, uint32(v218)))
							v336 = v259 + v221
							v337 = v258
							v338 = v223
						case 4:
							v255 = v222
							v256 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v218)+4)))
							v258 = v255 + v256
							v259 = *(*int32)(unsafe.Add(mBase, uint32(v218)))
							v336 = v259 + v221
							v337 = v258
							v338 = v223
						case 5:
							v250 = v222
							v251 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v218)+5)))
							v255 = v251<<(uint(int32(8))%32) + v250
							v256 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v218)+4)))
							v258 = v255 + v256
							v259 = *(*int32)(unsafe.Add(mBase, uint32(v218)))
							v336 = v259 + v221
							v337 = v258
							v338 = v223
						case 6:
							v246 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v218)+6)))
							v250 = v246<<(uint(int32(16))%32) + v222
							v251 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v218)+5)))
							v255 = v251<<(uint(int32(8))%32) + v250
							v256 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v218)+4)))
							v258 = v255 + v256
							v259 = *(*int32)(unsafe.Add(mBase, uint32(v218)))
							v336 = v259 + v221
							v337 = v258
							v338 = v223
						case 7:
							v241 = v223
							v242 = *(*int32)(unsafe.Add(mBase, uint32(v218)))
							v244 = *(*int32)(unsafe.Add(mBase, uint32(v218)+4))
							v336 = v242 + v221
							v337 = v244 + v222
							v338 = v241
						case 8:
							v236 = v223
							v237 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v218)+8)))
							v241 = v237<<(uint(int32(8))%32) + v236
							v242 = *(*int32)(unsafe.Add(mBase, uint32(v218)))
							v244 = *(*int32)(unsafe.Add(mBase, uint32(v218)+4))
							v336 = v242 + v221
							v337 = v244 + v222
							v338 = v241
						case 9:
							v231 = v223
							v232 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v218)+9)))
							v236 = v232<<(uint(int32(16))%32) + v231
							v237 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v218)+8)))
							v241 = v237<<(uint(int32(8))%32) + v236
							v242 = *(*int32)(unsafe.Add(mBase, uint32(v218)))
							v244 = *(*int32)(unsafe.Add(mBase, uint32(v218)+4))
							v336 = v242 + v221
							v337 = v244 + v222
							v338 = v241
						case 10:
							v227 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v218)+10)))
							v231 = v227<<(uint(int32(24))%32) + v223
							v232 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v218)+9)))
							v236 = v232<<(uint(int32(16))%32) + v231
							v237 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v218)+8)))
							v241 = v237<<(uint(int32(8))%32) + v236
							v242 = *(*int32)(unsafe.Add(mBase, uint32(v218)))
							v244 = *(*int32)(unsafe.Add(mBase, uint32(v218)+4))
							v336 = v242 + v221
							v337 = v244 + v222
							v338 = v241
						default:
							v336 = v221
							v337 = v222
							v338 = v223
						}
					}
					v341 = int32(14)
					v343 = v337 ^ v338 - base.I32_rotl(v337, v341)
					v347 = v343 ^ v336 - base.I32_rotl(v343, int32(11))
					v351 = v347 ^ v337 - base.I32_rotl(v347, int32(25))
					v355 = v351 ^ v343 - base.I32_rotl(v351, int32(16))
					v359 = v355 ^ v347 - base.I32_rotl(v355, int32(4))
					v363 = v359 ^ v351 - base.I32_rotl(v359, v341)
					v373 = F_Int64GetDatum(m, base.I64_extend_i32_u(v363)<<(uint(int64(32))%64)|base.I64_extend_i32_u(v355^v363-base.I32_rotl(v363, int32(24))))
					mBase = m.M
					v374 = m.ExcPending
					if v374 != 0 {
						return int32(0)
					} else {
						v735 = v373
						v740 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
						if v740 != v10 {
							F_pfree(m, v10)
							mBase = m.M
							v743 = m.ExcPending
							if v743 != 0 {
								return int32(0)
							} else {
								return v735
							}
						} else {
							return v735
						}
					}
				} else {
					v375 = int32(0)
					if v21 == int32(1) {
						v379 = int32(4)
						v381 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v18))))
						if v381&int32(254) == int32(2) {
							v390 = v379
						} else {
							v390 = base.B2i32(v381 == int32(18)) << (uint(v379) % 32)
						}
						if v381 == int32(1) {
							v393 = v379
						} else {
							v393 = v390
						}
						v404 = v393
					} else {
						v394 = int32(1)
						if v23 != 0 {
							v404 = int32(base.Ui32(v21)>>(uint(v394)%32)) - v394
						} else {
							v398 = *(*int32)(unsafe.Add(mBase, uint32(v10)))
							v404 = int32(base.Ui32(v398)>>(uint(int32(2))%32)) - int32(4)
						}
					}
					v405 = F_pg_strnxfrm(m, v375, v375, v24, v404, v15)
					mBase = m.M
					v406 = m.ExcPending
					if v406 != 0 {
						return int32(0)
					} else {
						v408 = v405 + int32(1)
						v409 = F_palloc(m, v408)
						mBase = m.M
						v410 = m.ExcPending
						if v410 != 0 {
							return int32(0)
						} else {
							v411 = F_pg_strnxfrm(m, v409, v408, v24, v404, v15)
							mBase = m.M
							v412 = m.ExcPending
							if v412 != 0 {
								return int32(0)
							} else {
								if base.Ui32(v405) < base.Ui32(v411) {
									F_errstart_cold(m, int32(21), int32(0))
									mBase = m.M
									v768 = m.ExcPending
									if v768 != 0 {
										return int32(0)
									} else {
										F_errmsg_internal(m, int32(_a_F_hashtextextended_0), int32(0))
										mBase = m.M
										v772 = m.ExcPending
										if v772 != 0 {
											return int32(0)
										} else {
											F_errfinish(m, int32(_a_F_hashtextextended_1), int32(361), int32(_a_F_hashtextextended_2))
											mBase = m.M
											v777 = m.ExcPending
											if v777 != 0 {
												return int32(0)
											} else {
												base.Wasm_trap_unreachable()
												for {
												}
											}
										}
									}
								} else {
									v414 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
									v415 = *(*int64)(unsafe.Add(mBase, uint32(v414)))
									v421 = v408 - int32(1636608432)
									if v415 == int64(0) {
										v458 = v421
										v460 = v421
										v462 = v421
									} else {
										v425 = v421 + base.I32_wrap_i64(v415)
										v426 = v425 + v421
										v430 = int32(4)
										v432 = base.I32_wrap_i64(int64(base.Ui64(v415)>>(uint(int64(32))%64))) ^ base.I32_rotl(v421, v430)
										v436 = v425 - v432 ^ base.I32_rotl(v432, int32(6))
										v440 = v426 - v436 ^ base.I32_rotl(v436, int32(8))
										v441 = v432 + v426
										v442 = v436 + v441
										v443 = v440 + v442
										v447 = v441 - v440 ^ base.I32_rotl(v440, int32(16))
										v451 = v442 - v447 ^ base.I32_rotl(v447, int32(19))
										v456 = v447 + v443
										v458 = v456
										v460 = v443 - v451 ^ base.I32_rotl(v451, v430)
										v462 = v451 + v456
									}
									if v409&int32(3) != 0 {
										if base.Ui32(int32(11)) < base.Ui32(v408) {
											v467 = v409
											v468 = v408
											v470 = v458
											v471 = v462
											v472 = v460
											for {
												v474 = *(*int32)(unsafe.Add(mBase, uint32(v467)+4))
												v475 = v474 + v471
												v476 = *(*int32)(unsafe.Add(mBase, uint32(v467)))
												v478 = *(*int32)(unsafe.Add(mBase, uint32(v467)+8))
												v479 = v478 + v472
												v481 = int32(4)
												v483 = v476 + v470 - v479 ^ base.I32_rotl(v479, v481)
												v487 = v475 - v483 ^ base.I32_rotl(v483, int32(6))
												v488 = v479 + v475
												v489 = v483 + v488
												v490 = v487 + v489
												v494 = v488 - v487 ^ base.I32_rotl(v487, int32(8))
												v498 = v489 - v494 ^ base.I32_rotl(v494, int32(16))
												v502 = v490 - v498 ^ base.I32_rotl(v498, int32(19))
												v503 = v494 + v490
												v504 = v498 + v503
												v505 = v502 + v504
												v509 = v503 - v502 ^ base.I32_rotl(v502, v481)
												v510 = int32(12)
												v511 = v467 + v510
												v513 = v468 - v510
												if base.Ui32(int32(11)) < base.Ui32(v513) {
													v467 = v511
													v468 = v513
													v470 = v504
													v471 = v505
													v472 = v509
													continue
												} else {
													break
												}
												break
											}
											v516 = v511
											v517 = v513
											v519 = v504
											v520 = v505
											v521 = v509
										} else {
											v516 = v409
											v517 = v408
											v519 = v458
											v520 = v462
											v521 = v460
										}
										switch v517 - int32(1) {
										case 0:
											v686 = v519
											v687 = v520
											v688 = v521
											v689 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v516))))
											v694 = v686 + v689
											v695 = v687
											v696 = v688
										case 1:
											v679 = v519
											v680 = v520
											v681 = v521
											v682 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v516)+1)))
											v686 = v682<<(uint(int32(8))%32) + v679
											v687 = v680
											v688 = v681
											v689 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v516))))
											v694 = v686 + v689
											v695 = v687
											v696 = v688
										case 2:
											v672 = v519
											v673 = v520
											v674 = v521
											v675 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v516)+2)))
											v679 = v675<<(uint(int32(16))%32) + v672
											v680 = v673
											v681 = v674
											v682 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v516)+1)))
											v686 = v682<<(uint(int32(8))%32) + v679
											v687 = v680
											v688 = v681
											v689 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v516))))
											v694 = v686 + v689
											v695 = v687
											v696 = v688
										case 3:
											v666 = v520
											v667 = v521
											v668 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v516)+3)))
											v672 = v668<<(uint(int32(24))%32) + v519
											v673 = v666
											v674 = v667
											v675 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v516)+2)))
											v679 = v675<<(uint(int32(16))%32) + v672
											v680 = v673
											v681 = v674
											v682 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v516)+1)))
											v686 = v682<<(uint(int32(8))%32) + v679
											v687 = v680
											v688 = v681
											v689 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v516))))
											v694 = v686 + v689
											v695 = v687
											v696 = v688
										case 4:
											v662 = v520
											v663 = v521
											v664 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v516)+4)))
											v666 = v662 + v664
											v667 = v663
											v668 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v516)+3)))
											v672 = v668<<(uint(int32(24))%32) + v519
											v673 = v666
											v674 = v667
											v675 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v516)+2)))
											v679 = v675<<(uint(int32(16))%32) + v672
											v680 = v673
											v681 = v674
											v682 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v516)+1)))
											v686 = v682<<(uint(int32(8))%32) + v679
											v687 = v680
											v688 = v681
											v689 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v516))))
											v694 = v686 + v689
											v695 = v687
											v696 = v688
										case 5:
											v656 = v520
											v657 = v521
											v658 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v516)+5)))
											v662 = v658<<(uint(int32(8))%32) + v656
											v663 = v657
											v664 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v516)+4)))
											v666 = v662 + v664
											v667 = v663
											v668 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v516)+3)))
											v672 = v668<<(uint(int32(24))%32) + v519
											v673 = v666
											v674 = v667
											v675 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v516)+2)))
											v679 = v675<<(uint(int32(16))%32) + v672
											v680 = v673
											v681 = v674
											v682 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v516)+1)))
											v686 = v682<<(uint(int32(8))%32) + v679
											v687 = v680
											v688 = v681
											v689 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v516))))
											v694 = v686 + v689
											v695 = v687
											v696 = v688
										case 6:
											v650 = v520
											v651 = v521
											v652 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v516)+6)))
											v656 = v652<<(uint(int32(16))%32) + v650
											v657 = v651
											v658 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v516)+5)))
											v662 = v658<<(uint(int32(8))%32) + v656
											v663 = v657
											v664 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v516)+4)))
											v666 = v662 + v664
											v667 = v663
											v668 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v516)+3)))
											v672 = v668<<(uint(int32(24))%32) + v519
											v673 = v666
											v674 = v667
											v675 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v516)+2)))
											v679 = v675<<(uint(int32(16))%32) + v672
											v680 = v673
											v681 = v674
											v682 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v516)+1)))
											v686 = v682<<(uint(int32(8))%32) + v679
											v687 = v680
											v688 = v681
											v689 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v516))))
											v694 = v686 + v689
											v695 = v687
											v696 = v688
										case 7:
											v645 = v521
											v646 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v516)+7)))
											v650 = v646<<(uint(int32(24))%32) + v520
											v651 = v645
											v652 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v516)+6)))
											v656 = v652<<(uint(int32(16))%32) + v650
											v657 = v651
											v658 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v516)+5)))
											v662 = v658<<(uint(int32(8))%32) + v656
											v663 = v657
											v664 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v516)+4)))
											v666 = v662 + v664
											v667 = v663
											v668 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v516)+3)))
											v672 = v668<<(uint(int32(24))%32) + v519
											v673 = v666
											v674 = v667
											v675 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v516)+2)))
											v679 = v675<<(uint(int32(16))%32) + v672
											v680 = v673
											v681 = v674
											v682 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v516)+1)))
											v686 = v682<<(uint(int32(8))%32) + v679
											v687 = v680
											v688 = v681
											v689 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v516))))
											v694 = v686 + v689
											v695 = v687
											v696 = v688
										case 8:
											v640 = v521
											v641 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v516)+8)))
											v645 = v641<<(uint(int32(8))%32) + v640
											v646 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v516)+7)))
											v650 = v646<<(uint(int32(24))%32) + v520
											v651 = v645
											v652 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v516)+6)))
											v656 = v652<<(uint(int32(16))%32) + v650
											v657 = v651
											v658 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v516)+5)))
											v662 = v658<<(uint(int32(8))%32) + v656
											v663 = v657
											v664 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v516)+4)))
											v666 = v662 + v664
											v667 = v663
											v668 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v516)+3)))
											v672 = v668<<(uint(int32(24))%32) + v519
											v673 = v666
											v674 = v667
											v675 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v516)+2)))
											v679 = v675<<(uint(int32(16))%32) + v672
											v680 = v673
											v681 = v674
											v682 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v516)+1)))
											v686 = v682<<(uint(int32(8))%32) + v679
											v687 = v680
											v688 = v681
											v689 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v516))))
											v694 = v686 + v689
											v695 = v687
											v696 = v688
										case 9:
											v635 = v521
											v636 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v516)+9)))
											v640 = v636<<(uint(int32(16))%32) + v635
											v641 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v516)+8)))
											v645 = v641<<(uint(int32(8))%32) + v640
											v646 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v516)+7)))
											v650 = v646<<(uint(int32(24))%32) + v520
											v651 = v645
											v652 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v516)+6)))
											v656 = v652<<(uint(int32(16))%32) + v650
											v657 = v651
											v658 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v516)+5)))
											v662 = v658<<(uint(int32(8))%32) + v656
											v663 = v657
											v664 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v516)+4)))
											v666 = v662 + v664
											v667 = v663
											v668 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v516)+3)))
											v672 = v668<<(uint(int32(24))%32) + v519
											v673 = v666
											v674 = v667
											v675 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v516)+2)))
											v679 = v675<<(uint(int32(16))%32) + v672
											v680 = v673
											v681 = v674
											v682 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v516)+1)))
											v686 = v682<<(uint(int32(8))%32) + v679
											v687 = v680
											v688 = v681
											v689 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v516))))
											v694 = v686 + v689
											v695 = v687
											v696 = v688
										case 10:
											v631 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v516)+10)))
											v635 = v631<<(uint(int32(24))%32) + v521
											v636 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v516)+9)))
											v640 = v636<<(uint(int32(16))%32) + v635
											v641 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v516)+8)))
											v645 = v641<<(uint(int32(8))%32) + v640
											v646 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v516)+7)))
											v650 = v646<<(uint(int32(24))%32) + v520
											v651 = v645
											v652 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v516)+6)))
											v656 = v652<<(uint(int32(16))%32) + v650
											v657 = v651
											v658 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v516)+5)))
											v662 = v658<<(uint(int32(8))%32) + v656
											v663 = v657
											v664 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v516)+4)))
											v666 = v662 + v664
											v667 = v663
											v668 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v516)+3)))
											v672 = v668<<(uint(int32(24))%32) + v519
											v673 = v666
											v674 = v667
											v675 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v516)+2)))
											v679 = v675<<(uint(int32(16))%32) + v672
											v680 = v673
											v681 = v674
											v682 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v516)+1)))
											v686 = v682<<(uint(int32(8))%32) + v679
											v687 = v680
											v688 = v681
											v689 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v516))))
											v694 = v686 + v689
											v695 = v687
											v696 = v688
										default:
											v694 = v519
											v695 = v520
											v696 = v521
										}
									} else {
										if base.Ui32(int32(12)) <= base.Ui32(v408) {
											v527 = v409
											v528 = v408
											v530 = v458
											v531 = v462
											v532 = v460
											for {
												v534 = *(*int32)(unsafe.Add(mBase, uint32(v527)+4))
												v535 = v534 + v531
												v536 = *(*int32)(unsafe.Add(mBase, uint32(v527)))
												v538 = *(*int32)(unsafe.Add(mBase, uint32(v527)+8))
												v539 = v538 + v532
												v541 = int32(4)
												v543 = v536 + v530 - v539 ^ base.I32_rotl(v539, v541)
												v547 = v535 - v543 ^ base.I32_rotl(v543, int32(6))
												v548 = v539 + v535
												v549 = v543 + v548
												v550 = v547 + v549
												v554 = v548 - v547 ^ base.I32_rotl(v547, int32(8))
												v558 = v549 - v554 ^ base.I32_rotl(v554, int32(16))
												v562 = v550 - v558 ^ base.I32_rotl(v558, int32(19))
												v563 = v554 + v550
												v564 = v558 + v563
												v565 = v562 + v564
												v569 = v563 - v562 ^ base.I32_rotl(v562, v541)
												v570 = int32(12)
												v571 = v527 + v570
												v573 = v528 - v570
												if base.Ui32(int32(11)) < base.Ui32(v573) {
													v527 = v571
													v528 = v573
													v530 = v564
													v531 = v565
													v532 = v569
													continue
												} else {
													break
												}
												break
											}
											v576 = v571
											v577 = v573
											v579 = v564
											v580 = v565
											v581 = v569
										} else {
											v576 = v409
											v577 = v408
											v579 = v458
											v580 = v462
											v581 = v460
										}
										switch v577 - int32(1) {
										case 0:
											v628 = v579
											v629 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v576))))
											v694 = v628 + v629
											v695 = v580
											v696 = v581
										case 1:
											v623 = v579
											v624 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v576)+1)))
											v628 = v624<<(uint(int32(8))%32) + v623
											v629 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v576))))
											v694 = v628 + v629
											v695 = v580
											v696 = v581
										case 2:
											v619 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v576)+2)))
											v623 = v619<<(uint(int32(16))%32) + v579
											v624 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v576)+1)))
											v628 = v624<<(uint(int32(8))%32) + v623
											v629 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v576))))
											v694 = v628 + v629
											v695 = v580
											v696 = v581
										case 3:
											v616 = v580
											v617 = *(*int32)(unsafe.Add(mBase, uint32(v576)))
											v694 = v617 + v579
											v695 = v616
											v696 = v581
										case 4:
											v613 = v580
											v614 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v576)+4)))
											v616 = v613 + v614
											v617 = *(*int32)(unsafe.Add(mBase, uint32(v576)))
											v694 = v617 + v579
											v695 = v616
											v696 = v581
										case 5:
											v608 = v580
											v609 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v576)+5)))
											v613 = v609<<(uint(int32(8))%32) + v608
											v614 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v576)+4)))
											v616 = v613 + v614
											v617 = *(*int32)(unsafe.Add(mBase, uint32(v576)))
											v694 = v617 + v579
											v695 = v616
											v696 = v581
										case 6:
											v604 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v576)+6)))
											v608 = v604<<(uint(int32(16))%32) + v580
											v609 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v576)+5)))
											v613 = v609<<(uint(int32(8))%32) + v608
											v614 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v576)+4)))
											v616 = v613 + v614
											v617 = *(*int32)(unsafe.Add(mBase, uint32(v576)))
											v694 = v617 + v579
											v695 = v616
											v696 = v581
										case 7:
											v599 = v581
											v600 = *(*int32)(unsafe.Add(mBase, uint32(v576)))
											v602 = *(*int32)(unsafe.Add(mBase, uint32(v576)+4))
											v694 = v600 + v579
											v695 = v602 + v580
											v696 = v599
										case 8:
											v594 = v581
											v595 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v576)+8)))
											v599 = v595<<(uint(int32(8))%32) + v594
											v600 = *(*int32)(unsafe.Add(mBase, uint32(v576)))
											v602 = *(*int32)(unsafe.Add(mBase, uint32(v576)+4))
											v694 = v600 + v579
											v695 = v602 + v580
											v696 = v599
										case 9:
											v589 = v581
											v590 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v576)+9)))
											v594 = v590<<(uint(int32(16))%32) + v589
											v595 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v576)+8)))
											v599 = v595<<(uint(int32(8))%32) + v594
											v600 = *(*int32)(unsafe.Add(mBase, uint32(v576)))
											v602 = *(*int32)(unsafe.Add(mBase, uint32(v576)+4))
											v694 = v600 + v579
											v695 = v602 + v580
											v696 = v599
										case 10:
											v585 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v576)+10)))
											v589 = v585<<(uint(int32(24))%32) + v581
											v590 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v576)+9)))
											v594 = v590<<(uint(int32(16))%32) + v589
											v595 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v576)+8)))
											v599 = v595<<(uint(int32(8))%32) + v594
											v600 = *(*int32)(unsafe.Add(mBase, uint32(v576)))
											v602 = *(*int32)(unsafe.Add(mBase, uint32(v576)+4))
											v694 = v600 + v579
											v695 = v602 + v580
											v696 = v599
										default:
											v694 = v579
											v695 = v580
											v696 = v581
										}
									}
									v699 = int32(14)
									v701 = v695 ^ v696 - base.I32_rotl(v695, v699)
									v705 = v701 ^ v694 - base.I32_rotl(v701, int32(11))
									v709 = v705 ^ v695 - base.I32_rotl(v705, int32(25))
									v713 = v709 ^ v701 - base.I32_rotl(v709, int32(16))
									v717 = v713 ^ v705 - base.I32_rotl(v713, int32(4))
									v721 = v717 ^ v709 - base.I32_rotl(v717, v699)
									v731 = F_Int64GetDatum(m, base.I64_extend_i32_u(v721)<<(uint(int64(32))%64)|base.I64_extend_i32_u(v713^v721-base.I32_rotl(v721, int32(24))))
									mBase = m.M
									v732 = m.ExcPending
									if v732 != 0 {
										return int32(0)
									} else {
										F_pfree(m, v409)
										mBase = m.M
										v734 = m.ExcPending
										if v734 != 0 {
											return int32(0)
										} else {
											v735 = v731
											v740 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
											if v740 != v10 {
												F_pfree(m, v10)
												mBase = m.M
												v743 = m.ExcPending
												if v743 != 0 {
													return int32(0)
												} else {
													return v735
												}
											} else {
												return v735
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
			F_errstart_cold(m, int32(21), int32(0))
			mBase = m.M
			v748 = m.ExcPending
			if v748 != 0 {
				return int32(0)
			} else {
				F_errcode(m, int32(34209924))
				mBase = m.M
				v751 = m.ExcPending
				if v751 != 0 {
					return int32(0)
				} else {
					F_errmsg(m, int32(_a_F_hashtextextended_3), int32(0))
					mBase = m.M
					v755 = m.ExcPending
					if v755 != 0 {
						return int32(0)
					} else {
						F_errhint(m, int32(_a_F_hashtextextended_4), int32(0))
						mBase = m.M
						v759 = m.ExcPending
						if v759 != 0 {
							return int32(0)
						} else {
							F_errfinish(m, int32(_a_F_hashtextextended_1), int32(336), int32(_a_F_hashtextextended_2))
							mBase = m.M
							v764 = m.ExcPending
							if v764 != 0 {
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
		}
	}
}
func F_hashvalidate(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v2 int32
	_ = v2
	var v15 int32
	_ = v15
	var v17 int32
	_ = v17
	var v20 int32
	_ = v20
	var v23 int32
	_ = v23
	var v24 int32
	_ = v24
	var v25 int32
	_ = v25
	var v26 int32
	_ = v26
	var v27 int32
	_ = v27
	var v28 int32
	_ = v28
	var v29 int32
	_ = v29
	var v30 int32
	_ = v30
	var v33 int32
	_ = v33
	var v35 int32
	_ = v35
	var v36 int32
	_ = v36
	var v37 int32
	_ = v37
	var v40 int32
	_ = v40
	var v42 int32
	_ = v42
	var v43 int32
	_ = v43
	var v44 int32
	_ = v44
	var v52 int32
	_ = v52
	var v53 int32
	_ = v53
	var v57 int32
	_ = v57
	var v66 int32
	_ = v66
	var v67 int32
	_ = v67
	var v68 int32
	_ = v68
	var v69 int32
	_ = v69
	var v70 int32
	_ = v70
	var v71 int32
	_ = v71
	var v73 int32
	_ = v73
	var v76 int32
	_ = v76
	var v77 int32
	_ = v77
	var v82 int32
	_ = v82
	var v83 int32
	_ = v83
	var v84 int32
	_ = v84
	var v85 int32
	_ = v85
	var v94 int32
	_ = v94
	var v99 int32
	_ = v99
	var v100 int32
	_ = v100
	var v101 int32
	_ = v101
	var v104 int32
	_ = v104
	var v105 int32
	_ = v105
	var v106 int32
	_ = v106
	var v109 int32
	_ = v109
	var v112 int32
	_ = v112
	var v113 int32
	_ = v113
	var v118 int32
	_ = v118
	var v119 int32
	_ = v119
	var v120 int32
	_ = v120
	var v121 int32
	_ = v121
	var v122 int32
	_ = v122
	var v132 int32
	_ = v132
	var v137 int32
	_ = v137
	var v138 int32
	_ = v138
	var v139 int32
	_ = v139
	var v142 int32
	_ = v142
	var v147 int32
	_ = v147
	var v148 int32
	_ = v148
	var v149 int32
	_ = v149
	var v150 int32
	_ = v150
	var v151 int32
	_ = v151
	var v156 int32
	_ = v156
	var v160 int32
	_ = v160
	var v161 int32
	_ = v161
	var v166 int32
	_ = v166
	var v167 int32
	_ = v167
	var v173 int32
	_ = v173
	var v174 int32
	_ = v174
	var v175 int32
	_ = v175
	var v178 int32
	_ = v178
	var v181 int32
	_ = v181
	var v182 int32
	_ = v182
	var v187 int32
	_ = v187
	var v188 int32
	_ = v188
	var v189 int32
	_ = v189
	var v190 int32
	_ = v190
	var v191 int32
	_ = v191
	var v201 int32
	_ = v201
	var v206 int32
	_ = v206
	var v208 int32
	_ = v208
	var v209 int32
	_ = v209
	var v212 int32
	_ = v212
	var v213 int32
	_ = v213
	var v218 int32
	_ = v218
	var v222 int32
	_ = v222
	var v227 int32
	_ = v227
	var v232 int32
	_ = v232
	var v236 int32
	_ = v236
	var v242 int32
	_ = v242
	var v251 int32
	_ = v251
	var v252 int32
	_ = v252
	var v265 int32
	_ = v265
	var v266 int32
	_ = v266
	var v267 int32
	_ = v267
	var v268 int32
	_ = v268
	var v269 int32
	_ = v269
	var v272 int32
	_ = v272
	var v275 int32
	_ = v275
	var v276 int32
	_ = v276
	var v281 int32
	_ = v281
	var v282 int32
	_ = v282
	var v283 int32
	_ = v283
	var v284 int32
	_ = v284
	var v285 int32
	_ = v285
	var v295 int32
	_ = v295
	var v300 int32
	_ = v300
	var v302 int32
	_ = v302
	var v303 int32
	_ = v303
	var v306 int32
	_ = v306
	var v309 int32
	_ = v309
	var v312 int32
	_ = v312
	var v313 int32
	_ = v313
	var v318 int32
	_ = v318
	var v319 int32
	_ = v319
	var v320 int32
	_ = v320
	var v321 int32
	_ = v321
	var v330 int32
	_ = v330
	var v335 int32
	_ = v335
	var v336 int32
	_ = v336
	var v337 int32
	_ = v337
	var v339 int32
	_ = v339
	var v340 int32
	_ = v340
	var v341 int32
	_ = v341
	var v342 int32
	_ = v342
	var v343 int32
	_ = v343
	var v346 int32
	_ = v346
	var v347 int32
	_ = v347
	var v352 int32
	_ = v352
	var v353 int32
	_ = v353
	var v354 int32
	_ = v354
	var v355 int32
	_ = v355
	var v364 int32
	_ = v364
	var v369 int32
	_ = v369
	var v370 int32
	_ = v370
	var v371 int32
	_ = v371
	var v372 int32
	_ = v372
	var v378 int32
	_ = v378
	var v381 int32
	_ = v381
	var v384 int32
	_ = v384
	var v385 int32
	_ = v385
	var v387 int32
	_ = v387
	var v395 int32
	_ = v395
	var v396 int32
	_ = v396
	var v398 int32
	_ = v398
	var v403 int32
	_ = v403
	var v410 int32
	_ = v410
	var v411 int32
	_ = v411
	var v412 int32
	_ = v412
	var v418 int32
	_ = v418
	var v421 int32
	_ = v421
	var v424 int32
	_ = v424
	var v425 int32
	_ = v425
	var v427 int32
	_ = v427
	var v435 int32
	_ = v435
	var v436 int32
	_ = v436
	var v438 int32
	_ = v438
	var v443 int32
	_ = v443
	var v450 int32
	_ = v450
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
	var v462 int32
	_ = v462
	var v463 int32
	_ = v463
	var v472 int32
	_ = v472
	var v477 int32
	_ = v477
	var v478 int32
	_ = v478
	var v480 int32
	_ = v480
	var v481 int32
	_ = v481
	var v487 int32
	_ = v487
	var v497 int32
	_ = v497
	var v498 int32
	_ = v498
	var v499 int32
	_ = v499
	var v500 int32
	_ = v500
	var v507 int32
	_ = v507
	var v508 int32
	_ = v508
	var v509 int32
	_ = v509
	var v519 int32
	_ = v519
	var v523 int32
	_ = v523
	var v524 int32
	_ = v524
	var v526 int32
	_ = v526
	var v528 int32
	_ = v528
	var v529 int32
	_ = v529
	var v530 int64
	_ = v530
	var v533 int32
	_ = v533
	var v536 int32
	_ = v536
	var v537 int32
	_ = v537
	var v542 int32
	_ = v542
	var v543 int32
	_ = v543
	var v544 int32
	_ = v544
	var v545 int32
	_ = v545
	var v546 int32
	_ = v546
	var v547 int32
	_ = v547
	var v548 int32
	_ = v548
	var v558 int32
	_ = v558
	var v563 int32
	_ = v563
	var v564 int32
	_ = v564
	var v567 int32
	_ = v567
	var v568 int32
	_ = v568
	var v576 int32
	_ = v576
	var v586 int32
	_ = v586
	var v603 int32
	_ = v603
	var v606 int32
	_ = v606
	var v607 int32
	_ = v607
	var v610 int32
	_ = v610
	var v620 int32
	_ = v620
	var v625 int32
	_ = v625
	var v626 int32
	_ = v626
	var v633 int32
	_ = v633
	var v643 int32
	_ = v643
	var v648 int32
	_ = v648
	var v658 int32
	_ = v658
	var v659 int32
	_ = v659
	var v661 int32
	_ = v661
	var v664 int32
	_ = v664
	var v667 int32
	_ = v667
	var v668 int32
	_ = v668
	var v673 int32
	_ = v673
	var v681 int32
	_ = v681
	var v686 int32
	_ = v686
	var v687 int32
	_ = v687
	var v689 int32
	_ = v689
	var v691 int32
	_ = v691
	var v693 int32
	_ = v693
	v2 = int32(0)
	v15 = m.G0
	v17 = v15 - int32(208)
	m.G0 = v17
	v20 = F_SearchSysCache1(m, int32(14), l0)
	mBase = m.M
	v23 = m.ExcPending
	if v23 != 0 {
		goto L2
	} else {
		goto L3
	}
L1:
	;
	v242 = *(*int32)(unsafe.Add(mBase, uint32(v35)+40))
	if int32(0) < v242 {
		goto L52
	} else {
		goto L53
	}
L2:
	;
	return int32(0)
L3:
	;
	if v20 != 0 {
		goto L4
	} else {
		goto L5
	}
L4:
	;
	v24 = *(*int32)(unsafe.Add(mBase, uint32(v20)+16))
	v25 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v24)+22)))
	v26 = v24 + v25
	v27 = *(*int32)(unsafe.Add(mBase, uint32(v26)+84))
	v28 = *(*int32)(unsafe.Add(mBase, uint32(v26)+80))
	v29 = F_get_opfamily_name(m, v28)
	mBase = m.M
	v30 = m.ExcPending
	if v30 != 0 {
		goto L2
	} else {
		goto L7
	}
L5:
	;
	goto L6
L6:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v218 = m.ExcPending
	if v218 != 0 {
		goto L2
	} else {
		goto L49
	}
L7:
	;
	v33 = int32(0)
	v35 = F_SearchSysCacheList(m, int32(4), int32(1), v28, v33, v33)
	mBase = m.M
	v36 = m.ExcPending
	if v36 != 0 {
		goto L2
	} else {
		goto L8
	}
L8:
	;
	v37 = int32(1)
	v40 = int32(0)
	v42 = F_SearchSysCacheList(m, int32(5), v37, v28, v40, v40)
	mBase = m.M
	v43 = m.ExcPending
	if v43 != 0 {
		goto L2
	} else {
		goto L9
	}
L9:
	;
	v44 = *(*int32)(unsafe.Add(mBase, uint32(v42)+40))
	if v44 <= int32(0) {
		v232 = v37
		v236 = v2
		goto L1
	} else {
		goto L10
	}
L10:
	;
	v52 = v2
	v53 = v37
	v57 = v2
	goto L11
L11:
	;
	v66 = *(*int32)(unsafe.Add(mBase, uint32(v42+int32(48)+v52<<(uint(int32(2))%32))))
	v67 = *(*int32)(unsafe.Add(mBase, uint32(v66)+56))
	v68 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v67)+22)))
	v69 = v67 + v68
	v70 = *(*int32)(unsafe.Add(mBase, uint32(v69)+8))
	v71 = *(*int32)(unsafe.Add(mBase, uint32(v69)+12))
	if v70 == v71 {
		v100 = v53
		goto L13
	} else {
		goto L14
	}
L12:
	;
	v232 = v208
	v236 = v209
	goto L1
L13:
	;
	v101 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v69)+16)))
	switch v101 - int32(1) {
	case 0:
		goto L25
	case 1:
		goto L24
	case 2:
		goto L27
	default:
		goto L26
	}
L14:
	;
	v73 = int32(0)
	v76 = F_errstart(m, int32(17), v73)
	mBase = m.M
	v77 = m.ExcPending
	if v77 != 0 {
		goto L2
	} else {
		goto L15
	}
L15:
	;
	if v76 == int32(0) {
		v100 = v73
		goto L13
	} else {
		goto L16
	}
L16:
	;
	F_errcode(m, int32(117833860))
	mBase = m.M
	v82 = m.ExcPending
	if v82 != 0 {
		goto L2
	} else {
		goto L17
	}
L17:
	;
	v83 = *(*int32)(unsafe.Add(mBase, uint32(v69)+20))
	v84 = F_format_procedure(m, v83)
	mBase = m.M
	v85 = m.ExcPending
	if v85 != 0 {
		goto L2
	} else {
		goto L18
	}
L18:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v17)+200)) = v84
	*(*int32)(unsafe.Add(mBase, uint32(v17)+196)) = int32(_a_F_hashvalidate_0)
	*(*int32)(unsafe.Add(mBase, uint32(v17)+192)) = v29
	F_errmsg(m, int32(_a_F_hashvalidate_1), v17+int32(192))
	mBase = m.M
	v94 = m.ExcPending
	if v94 != 0 {
		goto L2
	} else {
		goto L19
	}
L19:
	;
	F_errfinish(m, int32(_a_F_hashvalidate_2), int32(91), int32(_a_F_hashvalidate_3))
	mBase = m.M
	v99 = m.ExcPending
	if v99 != 0 {
		goto L2
	} else {
		goto L20
	}
L20:
	;
	v100 = v73
	goto L13
L21:
	;
	v212 = v52 + int32(1)
	v213 = *(*int32)(unsafe.Add(mBase, uint32(v42)+40))
	if v212 < v213 {
		v52 = v212
		v53 = v208
		v57 = v209
		goto L11
	} else {
		goto L48
	}
L22:
	;
	v178 = int32(0)
	v181 = F_errstart(m, int32(17), v178)
	mBase = m.M
	v182 = m.ExcPending
	if v182 != 0 {
		goto L2
	} else {
		goto L42
	}
L23:
	;
	v166 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v69)+16)))
	v167 = int32(1)
	if base.Ui32(v167) < base.Ui32((v166-v167)&int32(_a_F_hashvalidate_4)) {
		v208 = v100
		v209 = v57
		goto L21
	} else {
		goto L40
	}
L24:
	;
	v149 = *(*int32)(unsafe.Add(mBase, uint32(v69)+20))
	v150 = *(*int32)(unsafe.Add(mBase, uint32(v69)+8))
	v151 = int32(20)
	*(*int32)(unsafe.Add(mBase, uint32(v17)+180)) = v151
	*(*int32)(unsafe.Add(mBase, uint32(v17)+176)) = v150
	v156 = int32(2)
	v160 = F_check_amproc_signature(m, v149, v151, int32(1), v156, v156, v17+int32(176))
	mBase = m.M
	v161 = m.ExcPending
	if v161 != 0 {
		goto L2
	} else {
		goto L38
	}
L25:
	;
	v138 = *(*int32)(unsafe.Add(mBase, uint32(v69)+20))
	v139 = *(*int32)(unsafe.Add(mBase, uint32(v69)+8))
	*(*int32)(unsafe.Add(mBase, uint32(v17)+160)) = v139
	v142 = int32(1)
	v147 = F_check_amproc_signature(m, v138, int32(23), v142, v142, v142, v17+int32(160))
	mBase = m.M
	v148 = m.ExcPending
	if v148 != 0 {
		goto L2
	} else {
		goto L36
	}
L26:
	;
	v109 = int32(0)
	v112 = F_errstart(m, int32(17), v109)
	mBase = m.M
	v113 = m.ExcPending
	if v113 != 0 {
		goto L2
	} else {
		goto L30
	}
L27:
	;
	v104 = *(*int32)(unsafe.Add(mBase, uint32(v69)+20))
	v105 = F_check_amoptsproc_signature(m, v104)
	mBase = m.M
	v106 = m.ExcPending
	if v106 != 0 {
		goto L2
	} else {
		goto L28
	}
L28:
	;
	if v105 == int32(0) {
		goto L22
	} else {
		goto L29
	}
L29:
	;
	goto L23
L30:
	;
	if v112 == int32(0) {
		v208 = v109
		v209 = v57
		goto L21
	} else {
		goto L31
	}
L31:
	;
	F_errcode(m, int32(117833860))
	mBase = m.M
	v118 = m.ExcPending
	if v118 != 0 {
		goto L2
	} else {
		goto L32
	}
L32:
	;
	v119 = *(*int32)(unsafe.Add(mBase, uint32(v69)+20))
	v120 = F_format_procedure(m, v119)
	mBase = m.M
	v121 = m.ExcPending
	if v121 != 0 {
		goto L2
	} else {
		goto L33
	}
L33:
	;
	v122 = int32(*(*int16)(unsafe.Add(mBase, uint32(v69)+16)))
	*(*int32)(unsafe.Add(mBase, uint32(v17)+140)) = v122
	*(*int32)(unsafe.Add(mBase, uint32(v17)+136)) = v120
	*(*int32)(unsafe.Add(mBase, uint32(v17)+132)) = int32(_a_F_hashvalidate_0)
	*(*int32)(unsafe.Add(mBase, uint32(v17)+128)) = v29
	F_errmsg(m, int32(_a_F_hashvalidate_5), v17+int32(128))
	mBase = m.M
	v132 = m.ExcPending
	if v132 != 0 {
		goto L2
	} else {
		goto L34
	}
L34:
	;
	F_errfinish(m, int32(_a_F_hashvalidate_2), int32(115), int32(_a_F_hashvalidate_3))
	mBase = m.M
	v137 = m.ExcPending
	if v137 != 0 {
		goto L2
	} else {
		goto L35
	}
L35:
	;
	v208 = v109
	v209 = v57
	goto L21
L36:
	;
	if v147 != 0 {
		goto L23
	} else {
		goto L37
	}
L37:
	;
	goto L22
L38:
	;
	if v160 == int32(0) {
		goto L22
	} else {
		goto L39
	}
L39:
	;
	goto L23
L40:
	;
	v173 = *(*int32)(unsafe.Add(mBase, uint32(v69)+8))
	v174 = F_list_append_unique_oid(m, v57, v173)
	mBase = m.M
	v175 = m.ExcPending
	if v175 != 0 {
		goto L2
	} else {
		goto L41
	}
L41:
	;
	v208 = v100
	v209 = v174
	goto L21
L42:
	;
	if v181 == int32(0) {
		v208 = v178
		v209 = v57
		goto L21
	} else {
		goto L43
	}
L43:
	;
	F_errcode(m, int32(117833860))
	mBase = m.M
	v187 = m.ExcPending
	if v187 != 0 {
		goto L2
	} else {
		goto L44
	}
L44:
	;
	v188 = *(*int32)(unsafe.Add(mBase, uint32(v69)+20))
	v189 = F_format_procedure(m, v188)
	mBase = m.M
	v190 = m.ExcPending
	if v190 != 0 {
		goto L2
	} else {
		goto L45
	}
L45:
	;
	v191 = int32(*(*int16)(unsafe.Add(mBase, uint32(v69)+16)))
	*(*int32)(unsafe.Add(mBase, uint32(v17)+156)) = v191
	*(*int32)(unsafe.Add(mBase, uint32(v17)+152)) = v189
	*(*int32)(unsafe.Add(mBase, uint32(v17)+148)) = int32(_a_F_hashvalidate_0)
	*(*int32)(unsafe.Add(mBase, uint32(v17)+144)) = v29
	F_errmsg(m, int32(_a_F_hashvalidate_6), v17+int32(144))
	mBase = m.M
	v201 = m.ExcPending
	if v201 != 0 {
		goto L2
	} else {
		goto L46
	}
L46:
	;
	F_errfinish(m, int32(_a_F_hashvalidate_2), int32(127), int32(_a_F_hashvalidate_3))
	mBase = m.M
	v206 = m.ExcPending
	if v206 != 0 {
		goto L2
	} else {
		goto L47
	}
L47:
	;
	v208 = v178
	v209 = v57
	goto L21
L48:
	;
	goto L12
L49:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v17))) = l0
	F_errmsg_internal(m, int32(_a_F_hashvalidate_7), v17)
	mBase = m.M
	v222 = m.ExcPending
	if v222 != 0 {
		goto L2
	} else {
		goto L50
	}
L50:
	;
	F_errfinish(m, int32(_a_F_hashvalidate_2), int32(60), int32(_a_F_hashvalidate_3))
	mBase = m.M
	v227 = m.ExcPending
	if v227 != 0 {
		goto L2
	} else {
		goto L51
	}
L51:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L52:
	;
	v251 = int32(0)
	v252 = v232
	goto L55
L53:
	;
	v487 = v232
	goto L54
L54:
	;
	v497 = F_identify_opfamily_groups(m, v35, v42)
	mBase = m.M
	v498 = m.ExcPending
	if v498 != 0 {
		goto L2
	} else {
		goto L125
	}
L55:
	;
	v265 = *(*int32)(unsafe.Add(mBase, uint32(v35+int32(48)+v251<<(uint(int32(2))%32))))
	v266 = *(*int32)(unsafe.Add(mBase, uint32(v265)+56))
	v267 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v266)+22)))
	v268 = v266 + v267
	v269 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v268)+16)))
	if v269 == int32(1) {
		v302 = v252
		goto L57
	} else {
		goto L58
	}
L56:
	;
	v487 = v478
	goto L54
L57:
	;
	v303 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v268)+18)))
	if v303 == int32(115) {
		goto L66
	} else {
		goto L67
	}
L58:
	;
	v272 = int32(0)
	v275 = F_errstart(m, int32(17), v272)
	mBase = m.M
	v276 = m.ExcPending
	if v276 != 0 {
		goto L2
	} else {
		goto L59
	}
L59:
	;
	if v275 == int32(0) {
		v302 = v272
		goto L57
	} else {
		goto L60
	}
L60:
	;
	F_errcode(m, int32(117833860))
	mBase = m.M
	v281 = m.ExcPending
	if v281 != 0 {
		goto L2
	} else {
		goto L61
	}
L61:
	;
	v282 = *(*int32)(unsafe.Add(mBase, uint32(v268)+20))
	v283 = F_format_operator(m, v282)
	mBase = m.M
	v284 = m.ExcPending
	if v284 != 0 {
		goto L2
	} else {
		goto L62
	}
L62:
	;
	v285 = int32(*(*int16)(unsafe.Add(mBase, uint32(v268)+16)))
	*(*int32)(unsafe.Add(mBase, uint32(v17)+124)) = v285
	*(*int32)(unsafe.Add(mBase, uint32(v17)+120)) = v283
	*(*int32)(unsafe.Add(mBase, uint32(v17)+116)) = int32(_a_F_hashvalidate_0)
	*(*int32)(unsafe.Add(mBase, uint32(v17)+112)) = v29
	F_errmsg(m, int32(_a_F_hashvalidate_8), v17+int32(112))
	mBase = m.M
	v295 = m.ExcPending
	if v295 != 0 {
		goto L2
	} else {
		goto L63
	}
L63:
	;
	F_errfinish(m, int32(_a_F_hashvalidate_2), int32(153), int32(_a_F_hashvalidate_3))
	mBase = m.M
	v300 = m.ExcPending
	if v300 != 0 {
		goto L2
	} else {
		goto L64
	}
L64:
	;
	v302 = v272
	goto L57
L65:
	;
	v337 = *(*int32)(unsafe.Add(mBase, uint32(v268)+20))
	v339 = *(*int32)(unsafe.Add(mBase, uint32(v268)+8))
	v340 = *(*int32)(unsafe.Add(mBase, uint32(v268)+12))
	v341 = F_check_amop_signature(m, v337, int32(16), v339, v340)
	mBase = m.M
	v342 = m.ExcPending
	if v342 != 0 {
		goto L2
	} else {
		goto L77
	}
L66:
	;
	v306 = *(*int32)(unsafe.Add(mBase, uint32(v268)+28))
	if v306 == int32(0) {
		v336 = v302
		goto L65
	} else {
		goto L69
	}
L67:
	;
	goto L68
L68:
	;
	v309 = int32(0)
	v312 = F_errstart(m, int32(17), v309)
	mBase = m.M
	v313 = m.ExcPending
	if v313 != 0 {
		goto L2
	} else {
		goto L70
	}
L69:
	;
	goto L68
L70:
	;
	if v312 == int32(0) {
		v336 = v309
		goto L65
	} else {
		goto L71
	}
L71:
	;
	F_errcode(m, int32(117833860))
	mBase = m.M
	v318 = m.ExcPending
	if v318 != 0 {
		goto L2
	} else {
		goto L72
	}
L72:
	;
	v319 = *(*int32)(unsafe.Add(mBase, uint32(v268)+20))
	v320 = F_format_operator(m, v319)
	mBase = m.M
	v321 = m.ExcPending
	if v321 != 0 {
		goto L2
	} else {
		goto L73
	}
L73:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v17)+104)) = v320
	*(*int32)(unsafe.Add(mBase, uint32(v17)+100)) = int32(_a_F_hashvalidate_0)
	*(*int32)(unsafe.Add(mBase, uint32(v17)+96)) = v29
	F_errmsg(m, int32(_a_F_hashvalidate_9), v17+int32(96))
	mBase = m.M
	v330 = m.ExcPending
	if v330 != 0 {
		goto L2
	} else {
		goto L74
	}
L74:
	;
	F_errfinish(m, int32(_a_F_hashvalidate_2), int32(165), int32(_a_F_hashvalidate_3))
	mBase = m.M
	v335 = m.ExcPending
	if v335 != 0 {
		goto L2
	} else {
		goto L75
	}
L75:
	;
	v336 = v309
	goto L65
L76:
	;
	v371 = *(*int32)(unsafe.Add(mBase, uint32(v268)+8))
	v372 = int32(0)
	if v236 == v372 {
		goto L87
	} else {
		goto L88
	}
L77:
	;
	if v341 != 0 {
		v370 = v336
		goto L76
	} else {
		goto L78
	}
L78:
	;
	v343 = int32(0)
	v346 = F_errstart(m, int32(17), v343)
	mBase = m.M
	v347 = m.ExcPending
	if v347 != 0 {
		goto L2
	} else {
		goto L79
	}
L79:
	;
	if v346 == int32(0) {
		v370 = v343
		goto L76
	} else {
		goto L80
	}
L80:
	;
	F_errcode(m, int32(117833860))
	mBase = m.M
	v352 = m.ExcPending
	if v352 != 0 {
		goto L2
	} else {
		goto L81
	}
L81:
	;
	v353 = *(*int32)(unsafe.Add(mBase, uint32(v268)+20))
	v354 = F_format_operator(m, v353)
	mBase = m.M
	v355 = m.ExcPending
	if v355 != 0 {
		goto L2
	} else {
		goto L82
	}
L82:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v17)+88)) = v354
	*(*int32)(unsafe.Add(mBase, uint32(v17)+84)) = int32(_a_F_hashvalidate_0)
	*(*int32)(unsafe.Add(mBase, uint32(v17)+80)) = v29
	F_errmsg(m, int32(_a_F_hashvalidate_10), v17+int32(80))
	mBase = m.M
	v364 = m.ExcPending
	if v364 != 0 {
		goto L2
	} else {
		goto L83
	}
L83:
	;
	F_errfinish(m, int32(_a_F_hashvalidate_2), int32(178), int32(_a_F_hashvalidate_3))
	mBase = m.M
	v369 = m.ExcPending
	if v369 != 0 {
		goto L2
	} else {
		goto L84
	}
L84:
	;
	v370 = v343
	goto L76
L85:
	;
	v480 = v251 + int32(1)
	v481 = *(*int32)(unsafe.Add(mBase, uint32(v35)+40))
	if v480 < v481 {
		v251 = v480
		v252 = v478
		goto L55
	} else {
		goto L122
	}
L86:
	;
	if v410 != 0 {
		goto L99
	} else {
		goto L100
	}
L87:
	;
	v410 = int32(0)
	goto L86
L88:
	;
	goto L89
L89:
	;
	v378 = *(*int32)(unsafe.Add(mBase, uint32(v236)+4))
	if v378 <= int32(0) {
		v403 = v372
		goto L90
	} else {
		goto L91
	}
L90:
	;
	v410 = v403
	goto L86
L91:
	;
	v381 = int32(0)
	if v381 < v378 {
		goto L92
	} else {
		goto L93
	}
L92:
	;
	v384 = v378
	goto L94
L93:
	;
	v384 = v381
	goto L94
L94:
	;
	v385 = *(*int32)(unsafe.Add(mBase, uint32(v236)+12))
	v387 = int32(0)
	goto L95
L95:
	;
	v395 = *(*int32)(unsafe.Add(mBase, uint32(v385+v387<<(uint(int32(2))%32))))
	v396 = base.B2i32(v395 == v371)
	if v395 == v371 {
		v403 = v396
		goto L90
	} else {
		goto L97
	}
L96:
	;
	v403 = v396
	goto L90
L97:
	;
	v398 = v387 + int32(1)
	if v398 != v384 {
		v387 = v398
		goto L95
	} else {
		goto L98
	}
L98:
	;
	goto L96
L99:
	;
	v411 = *(*int32)(unsafe.Add(mBase, uint32(v268)+12))
	v412 = int32(0)
	if v236 == v412 {
		goto L103
	} else {
		goto L104
	}
L100:
	;
	goto L101
L101:
	;
	v451 = int32(0)
	v454 = F_errstart(m, int32(17), v451)
	mBase = m.M
	v455 = m.ExcPending
	if v455 != 0 {
		goto L2
	} else {
		goto L116
	}
L102:
	;
	if v450 != 0 {
		v478 = v370
		goto L85
	} else {
		goto L115
	}
L103:
	;
	v450 = int32(0)
	goto L102
L104:
	;
	goto L105
L105:
	;
	v418 = *(*int32)(unsafe.Add(mBase, uint32(v236)+4))
	if v418 <= int32(0) {
		v443 = v412
		goto L106
	} else {
		goto L107
	}
L106:
	;
	v450 = v443
	goto L102
L107:
	;
	v421 = int32(0)
	if v421 < v418 {
		goto L108
	} else {
		goto L109
	}
L108:
	;
	v424 = v418
	goto L110
L109:
	;
	v424 = v421
	goto L110
L110:
	;
	v425 = *(*int32)(unsafe.Add(mBase, uint32(v236)+12))
	v427 = int32(0)
	goto L111
L111:
	;
	v435 = *(*int32)(unsafe.Add(mBase, uint32(v425+v427<<(uint(int32(2))%32))))
	v436 = base.B2i32(v435 == v411)
	if v435 == v411 {
		v443 = v436
		goto L106
	} else {
		goto L113
	}
L112:
	;
	v443 = v436
	goto L106
L113:
	;
	v438 = v427 + int32(1)
	if v438 != v424 {
		v427 = v438
		goto L111
	} else {
		goto L114
	}
L114:
	;
	goto L112
L115:
	;
	goto L101
L116:
	;
	if v454 == int32(0) {
		v478 = v451
		goto L85
	} else {
		goto L117
	}
L117:
	;
	F_errcode(m, int32(117833860))
	mBase = m.M
	v460 = m.ExcPending
	if v460 != 0 {
		goto L2
	} else {
		goto L118
	}
L118:
	;
	v461 = *(*int32)(unsafe.Add(mBase, uint32(v268)+20))
	v462 = F_format_operator(m, v461)
	mBase = m.M
	v463 = m.ExcPending
	if v463 != 0 {
		goto L2
	} else {
		goto L119
	}
L119:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v17)+72)) = v462
	*(*int32)(unsafe.Add(mBase, uint32(v17)+68)) = int32(_a_F_hashvalidate_0)
	*(*int32)(unsafe.Add(mBase, uint32(v17)+64)) = v29
	F_errmsg(m, int32(_a_F_hashvalidate_11), v17-int32(-64))
	mBase = m.M
	v472 = m.ExcPending
	if v472 != 0 {
		goto L2
	} else {
		goto L120
	}
L120:
	;
	F_errfinish(m, int32(_a_F_hashvalidate_2), int32(190), int32(_a_F_hashvalidate_3))
	mBase = m.M
	v477 = m.ExcPending
	if v477 != 0 {
		goto L2
	} else {
		goto L121
	}
L121:
	;
	v478 = v451
	goto L85
L122:
	;
	goto L56
L123:
	;
	if v236 != 0 {
		goto L160
	} else {
		goto L161
	}
L124:
	;
	v643 = *(*int32)(unsafe.Add(mBase, uint32(v497)+4))
	v648 = v633
	v658 = v643
	goto L123
L125:
	;
	if v497 != 0 {
		goto L126
	} else {
		goto L127
	}
L126:
	;
	v499 = int32(0)
	v500 = *(*int32)(unsafe.Add(mBase, uint32(v497)+4))
	if v500 <= v499 {
		goto L129
	} else {
		goto L130
	}
L127:
	;
	goto L128
L128:
	;
	v603 = int32(0)
	v606 = F_errstart(m, int32(17), v603)
	mBase = m.M
	v607 = m.ExcPending
	if v607 != 0 {
		goto L2
	} else {
		goto L151
	}
L129:
	;
	v576 = v487
	v586 = int32(1)
	goto L131
L130:
	;
	v507 = int32(0)
	v508 = v499
	v509 = v487
	goto L132
L131:
	;
	if v586 == int32(0) {
		v633 = v576
		goto L124
	} else {
		goto L150
	}
L132:
	;
	v519 = *(*int32)(unsafe.Add(mBase, uint32(v497)+12))
	v523 = *(*int32)(unsafe.Add(mBase, uint32(v519+v508<<(uint(int32(2))%32))))
	v524 = *(*int32)(unsafe.Add(mBase, uint32(v523)))
	if v27 == v524 {
		goto L134
	} else {
		goto L135
	}
L133:
	;
	v576 = v564
	v586 = base.B2i32(v529 == int32(0))
	goto L131
L134:
	;
	v526 = *(*int32)(unsafe.Add(mBase, uint32(v523)+4))
	if v526 == v27 {
		goto L137
	} else {
		goto L138
	}
L135:
	;
	v529 = v507
	goto L136
L136:
	;
	v530 = *(*int64)(unsafe.Add(mBase, uint32(v523)+8))
	if v530 == int64(2) {
		v564 = v509
		goto L140
	} else {
		goto L141
	}
L137:
	;
	v528 = v523
	goto L139
L138:
	;
	v528 = v507
	goto L139
L139:
	;
	v529 = v528
	goto L136
L140:
	;
	v567 = v508 + int32(1)
	v568 = *(*int32)(unsafe.Add(mBase, uint32(v497)+4))
	if v567 < v568 {
		v507 = v529
		v508 = v567
		v509 = v564
		goto L132
	} else {
		goto L149
	}
L141:
	;
	v533 = int32(0)
	v536 = F_errstart(m, int32(17), v533)
	mBase = m.M
	v537 = m.ExcPending
	if v537 != 0 {
		goto L2
	} else {
		goto L142
	}
L142:
	;
	if v536 == int32(0) {
		v564 = v533
		goto L140
	} else {
		goto L143
	}
L143:
	;
	F_errcode(m, int32(117833860))
	mBase = m.M
	v542 = m.ExcPending
	if v542 != 0 {
		goto L2
	} else {
		goto L144
	}
L144:
	;
	v543 = *(*int32)(unsafe.Add(mBase, uint32(v523)))
	v544 = F_format_type_be(m, v543)
	mBase = m.M
	v545 = m.ExcPending
	if v545 != 0 {
		goto L2
	} else {
		goto L145
	}
L145:
	;
	v546 = *(*int32)(unsafe.Add(mBase, uint32(v523)+4))
	v547 = F_format_type_be(m, v546)
	mBase = m.M
	v548 = m.ExcPending
	if v548 != 0 {
		goto L2
	} else {
		goto L146
	}
L146:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v17)+60)) = v547
	*(*int32)(unsafe.Add(mBase, uint32(v17)+56)) = v544
	*(*int32)(unsafe.Add(mBase, uint32(v17)+52)) = int32(_a_F_hashvalidate_0)
	*(*int32)(unsafe.Add(mBase, uint32(v17)+48)) = v29
	F_errmsg(m, int32(_a_F_hashvalidate_12), v17+int32(48))
	mBase = m.M
	v558 = m.ExcPending
	if v558 != 0 {
		goto L2
	} else {
		goto L147
	}
L147:
	;
	F_errfinish(m, int32(_a_F_hashvalidate_2), int32(219), int32(_a_F_hashvalidate_3))
	mBase = m.M
	v563 = m.ExcPending
	if v563 != 0 {
		goto L2
	} else {
		goto L148
	}
L148:
	;
	v564 = v533
	goto L140
L149:
	;
	goto L133
L150:
	;
	goto L128
L151:
	;
	if v606 != 0 {
		goto L152
	} else {
		goto L153
	}
L152:
	;
	F_errcode(m, int32(117833860))
	mBase = m.M
	v610 = m.ExcPending
	if v610 != 0 {
		goto L2
	} else {
		goto L155
	}
L153:
	;
	goto L154
L154:
	;
	v626 = int32(0)
	if v497 == v626 {
		v648 = v603
		v658 = v626
		goto L123
	} else {
		goto L158
	}
L155:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v17)+36)) = int32(_a_F_hashvalidate_0)
	*(*int32)(unsafe.Add(mBase, uint32(v17)+32)) = v26 + int32(8)
	F_errmsg(m, int32(_a_F_hashvalidate_13), v17+int32(32))
	mBase = m.M
	v620 = m.ExcPending
	if v620 != 0 {
		goto L2
	} else {
		goto L156
	}
L156:
	;
	F_errfinish(m, int32(_a_F_hashvalidate_2), int32(231), int32(_a_F_hashvalidate_3))
	mBase = m.M
	v625 = m.ExcPending
	if v625 != 0 {
		goto L2
	} else {
		goto L157
	}
L157:
	;
	goto L154
L158:
	;
	v633 = v603
	goto L124
L159:
	;
	F_ReleaseCatCacheList(m, v42)
	mBase = m.M
	v689 = m.ExcPending
	if v689 != 0 {
		goto L2
	} else {
		goto L169
	}
L160:
	;
	v659 = *(*int32)(unsafe.Add(mBase, uint32(v236)+4))
	v661 = v659
	goto L162
L161:
	;
	v661 = int32(0)
	goto L162
L162:
	;
	if v658 == v661*v661 {
		v687 = v648
		goto L159
	} else {
		goto L163
	}
L163:
	;
	v664 = int32(0)
	v667 = F_errstart(m, int32(17), v664)
	mBase = m.M
	v668 = m.ExcPending
	if v668 != 0 {
		goto L2
	} else {
		goto L164
	}
L164:
	;
	if v667 == int32(0) {
		v687 = v664
		goto L159
	} else {
		goto L165
	}
L165:
	;
	F_errcode(m, int32(117833860))
	mBase = m.M
	v673 = m.ExcPending
	if v673 != 0 {
		goto L2
	} else {
		goto L166
	}
L166:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v17)+20)) = int32(_a_F_hashvalidate_0)
	*(*int32)(unsafe.Add(mBase, uint32(v17)+16)) = v29
	F_errmsg(m, int32(_a_F_hashvalidate_14), v17+int32(16))
	mBase = m.M
	v681 = m.ExcPending
	if v681 != 0 {
		goto L2
	} else {
		goto L167
	}
L167:
	;
	F_errfinish(m, int32(_a_F_hashvalidate_2), int32(247), int32(_a_F_hashvalidate_3))
	mBase = m.M
	v686 = m.ExcPending
	if v686 != 0 {
		goto L2
	} else {
		goto L168
	}
L168:
	;
	v687 = v664
	goto L159
L169:
	;
	F_ReleaseCatCacheList(m, v35)
	mBase = m.M
	v691 = m.ExcPending
	if v691 != 0 {
		goto L2
	} else {
		goto L170
	}
L170:
	;
	F_ReleaseCatCache(m, v20)
	mBase = m.M
	v693 = m.ExcPending
	if v693 != 0 {
		goto L2
	} else {
		goto L171
	}
L171:
	;
	m.G0 = v17 + int32(208)
	return v687
}
func F_hashvarlenaextended(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v6 int32
	_ = v6
	var v7 int32
	_ = v7
	var v10 int32
	_ = v10
	var v11 int32
	_ = v11
	var v12 int32
	_ = v12
	var v15 int32
	_ = v15
	var v17 int32
	_ = v17
	var v18 int32
	_ = v18
	var v21 int32
	_ = v21
	var v23 int32
	_ = v23
	var v32 int32
	_ = v32
	var v35 int32
	_ = v35
	var v36 int32
	_ = v36
	var v40 int32
	_ = v40
	var v46 int32
	_ = v46
	var v47 int32
	_ = v47
	var v48 int64
	_ = v48
	var v54 int32
	_ = v54
	var v58 int32
	_ = v58
	var v59 int32
	_ = v59
	var v63 int32
	_ = v63
	var v65 int32
	_ = v65
	var v69 int32
	_ = v69
	var v73 int32
	_ = v73
	var v74 int32
	_ = v74
	var v75 int32
	_ = v75
	var v76 int32
	_ = v76
	var v80 int32
	_ = v80
	var v84 int32
	_ = v84
	var v89 int32
	_ = v89
	var v91 int32
	_ = v91
	var v93 int32
	_ = v93
	var v95 int32
	_ = v95
	var v100 int32
	_ = v100
	var v101 int32
	_ = v101
	var v103 int32
	_ = v103
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
	var v111 int32
	_ = v111
	var v112 int32
	_ = v112
	var v114 int32
	_ = v114
	var v116 int32
	_ = v116
	var v120 int32
	_ = v120
	var v121 int32
	_ = v121
	var v122 int32
	_ = v122
	var v123 int32
	_ = v123
	var v127 int32
	_ = v127
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
	var v143 int32
	_ = v143
	var v144 int32
	_ = v144
	var v146 int32
	_ = v146
	var v149 int32
	_ = v149
	var v150 int32
	_ = v150
	var v152 int32
	_ = v152
	var v153 int32
	_ = v153
	var v154 int32
	_ = v154
	var v160 int32
	_ = v160
	var v161 int32
	_ = v161
	var v163 int32
	_ = v163
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
	var v174 int32
	_ = v174
	var v176 int32
	_ = v176
	var v180 int32
	_ = v180
	var v181 int32
	_ = v181
	var v182 int32
	_ = v182
	var v183 int32
	_ = v183
	var v187 int32
	_ = v187
	var v191 int32
	_ = v191
	var v195 int32
	_ = v195
	var v196 int32
	_ = v196
	var v197 int32
	_ = v197
	var v198 int32
	_ = v198
	var v202 int32
	_ = v202
	var v203 int32
	_ = v203
	var v204 int32
	_ = v204
	var v206 int32
	_ = v206
	var v209 int32
	_ = v209
	var v210 int32
	_ = v210
	var v212 int32
	_ = v212
	var v213 int32
	_ = v213
	var v214 int32
	_ = v214
	var v218 int32
	_ = v218
	var v222 int32
	_ = v222
	var v223 int32
	_ = v223
	var v227 int32
	_ = v227
	var v228 int32
	_ = v228
	var v232 int32
	_ = v232
	var v233 int32
	_ = v233
	var v235 int32
	_ = v235
	var v237 int32
	_ = v237
	var v241 int32
	_ = v241
	var v242 int32
	_ = v242
	var v246 int32
	_ = v246
	var v247 int32
	_ = v247
	var v249 int32
	_ = v249
	var v250 int32
	_ = v250
	var v252 int32
	_ = v252
	var v256 int32
	_ = v256
	var v257 int32
	_ = v257
	var v261 int32
	_ = v261
	var v262 int32
	_ = v262
	var v264 int32
	_ = v264
	var v268 int32
	_ = v268
	var v269 int32
	_ = v269
	var v273 int32
	_ = v273
	var v274 int32
	_ = v274
	var v278 int32
	_ = v278
	var v279 int32
	_ = v279
	var v283 int32
	_ = v283
	var v284 int32
	_ = v284
	var v285 int32
	_ = v285
	var v289 int32
	_ = v289
	var v290 int32
	_ = v290
	var v291 int32
	_ = v291
	var v295 int32
	_ = v295
	var v296 int32
	_ = v296
	var v297 int32
	_ = v297
	var v299 int32
	_ = v299
	var v300 int32
	_ = v300
	var v301 int32
	_ = v301
	var v305 int32
	_ = v305
	var v306 int32
	_ = v306
	var v307 int32
	_ = v307
	var v308 int32
	_ = v308
	var v312 int32
	_ = v312
	var v313 int32
	_ = v313
	var v314 int32
	_ = v314
	var v315 int32
	_ = v315
	var v319 int32
	_ = v319
	var v320 int32
	_ = v320
	var v321 int32
	_ = v321
	var v322 int32
	_ = v322
	var v327 int32
	_ = v327
	var v328 int32
	_ = v328
	var v329 int32
	_ = v329
	var v332 int32
	_ = v332
	var v334 int32
	_ = v334
	var v338 int32
	_ = v338
	var v342 int32
	_ = v342
	var v346 int32
	_ = v346
	var v350 int32
	_ = v350
	var v354 int32
	_ = v354
	var v364 int32
	_ = v364
	var v365 int32
	_ = v365
	var v366 int32
	_ = v366
	var v369 int32
	_ = v369
	v6 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v7 = F_pg_detoast_datum_packed(m, v6)
	mBase = m.M
	v10 = m.ExcPending
	if v10 != 0 {
		return int32(0)
	} else {
		v11 = int32(1)
		v12 = v7 + v11
		v15 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v7))))
		v17 = v15 & v11
		if v17 != 0 {
			v18 = v12
		} else {
			v18 = v7 + int32(4)
		}
		if v15 == int32(1) {
			v21 = int32(4)
			v23 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v12))))
			if v23&int32(254) == int32(2) {
				v32 = v21
			} else {
				v32 = base.B2i32(v23 == int32(18)) << (uint(v21) % 32)
			}
			if v23 == int32(1) {
				v35 = v21
			} else {
				v35 = v32
			}
			v46 = v35
		} else {
			v36 = int32(1)
			if v17 != 0 {
				v46 = int32(base.Ui32(v15)>>(uint(v36)%32)) - v36
			} else {
				v40 = *(*int32)(unsafe.Add(mBase, uint32(v7)))
				v46 = int32(base.Ui32(v40)>>(uint(int32(2))%32)) - int32(4)
			}
		}
		v47 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
		v48 = *(*int64)(unsafe.Add(mBase, uint32(v47)))
		v54 = v46 - int32(1636608432)
		if v48 == int64(0) {
			v91 = v54
			v93 = v54
			v95 = v54
		} else {
			v58 = v54 + base.I32_wrap_i64(v48)
			v59 = v58 + v54
			v63 = int32(4)
			v65 = base.I32_wrap_i64(int64(base.Ui64(v48)>>(uint(int64(32))%64))) ^ base.I32_rotl(v54, v63)
			v69 = v58 - v65 ^ base.I32_rotl(v65, int32(6))
			v73 = v59 - v69 ^ base.I32_rotl(v69, int32(8))
			v74 = v65 + v59
			v75 = v69 + v74
			v76 = v73 + v75
			v80 = v74 - v73 ^ base.I32_rotl(v73, int32(16))
			v84 = v75 - v80 ^ base.I32_rotl(v80, int32(19))
			v89 = v80 + v76
			v91 = v89
			v93 = v76 - v84 ^ base.I32_rotl(v84, v63)
			v95 = v84 + v89
		}
		if v18&int32(3) != 0 {
			if base.Ui32(int32(11)) < base.Ui32(v46) {
				v100 = v18
				v101 = v46
				v103 = v91
				v104 = v95
				v105 = v93
				for {
					v107 = *(*int32)(unsafe.Add(mBase, uint32(v100)+4))
					v108 = v107 + v104
					v109 = *(*int32)(unsafe.Add(mBase, uint32(v100)))
					v111 = *(*int32)(unsafe.Add(mBase, uint32(v100)+8))
					v112 = v111 + v105
					v114 = int32(4)
					v116 = v109 + v103 - v112 ^ base.I32_rotl(v112, v114)
					v120 = v108 - v116 ^ base.I32_rotl(v116, int32(6))
					v121 = v112 + v108
					v122 = v116 + v121
					v123 = v120 + v122
					v127 = v121 - v120 ^ base.I32_rotl(v120, int32(8))
					v131 = v122 - v127 ^ base.I32_rotl(v127, int32(16))
					v135 = v123 - v131 ^ base.I32_rotl(v131, int32(19))
					v136 = v127 + v123
					v137 = v131 + v136
					v138 = v135 + v137
					v142 = v136 - v135 ^ base.I32_rotl(v135, v114)
					v143 = int32(12)
					v144 = v100 + v143
					v146 = v101 - v143
					if base.Ui32(int32(11)) < base.Ui32(v146) {
						v100 = v144
						v101 = v146
						v103 = v137
						v104 = v138
						v105 = v142
						continue
					} else {
						break
					}
					break
				}
				v149 = v144
				v150 = v146
				v152 = v137
				v153 = v138
				v154 = v142
			} else {
				v149 = v18
				v150 = v46
				v152 = v91
				v153 = v95
				v154 = v93
			}
			switch v150 - int32(1) {
			case 0:
				v319 = v152
				v320 = v153
				v321 = v154
				v322 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v149))))
				v327 = v319 + v322
				v328 = v320
				v329 = v321
			case 1:
				v312 = v152
				v313 = v153
				v314 = v154
				v315 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v149)+1)))
				v319 = v315<<(uint(int32(8))%32) + v312
				v320 = v313
				v321 = v314
				v322 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v149))))
				v327 = v319 + v322
				v328 = v320
				v329 = v321
			case 2:
				v305 = v152
				v306 = v153
				v307 = v154
				v308 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v149)+2)))
				v312 = v308<<(uint(int32(16))%32) + v305
				v313 = v306
				v314 = v307
				v315 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v149)+1)))
				v319 = v315<<(uint(int32(8))%32) + v312
				v320 = v313
				v321 = v314
				v322 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v149))))
				v327 = v319 + v322
				v328 = v320
				v329 = v321
			case 3:
				v299 = v153
				v300 = v154
				v301 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v149)+3)))
				v305 = v301<<(uint(int32(24))%32) + v152
				v306 = v299
				v307 = v300
				v308 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v149)+2)))
				v312 = v308<<(uint(int32(16))%32) + v305
				v313 = v306
				v314 = v307
				v315 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v149)+1)))
				v319 = v315<<(uint(int32(8))%32) + v312
				v320 = v313
				v321 = v314
				v322 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v149))))
				v327 = v319 + v322
				v328 = v320
				v329 = v321
			case 4:
				v295 = v153
				v296 = v154
				v297 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v149)+4)))
				v299 = v295 + v297
				v300 = v296
				v301 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v149)+3)))
				v305 = v301<<(uint(int32(24))%32) + v152
				v306 = v299
				v307 = v300
				v308 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v149)+2)))
				v312 = v308<<(uint(int32(16))%32) + v305
				v313 = v306
				v314 = v307
				v315 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v149)+1)))
				v319 = v315<<(uint(int32(8))%32) + v312
				v320 = v313
				v321 = v314
				v322 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v149))))
				v327 = v319 + v322
				v328 = v320
				v329 = v321
			case 5:
				v289 = v153
				v290 = v154
				v291 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v149)+5)))
				v295 = v291<<(uint(int32(8))%32) + v289
				v296 = v290
				v297 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v149)+4)))
				v299 = v295 + v297
				v300 = v296
				v301 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v149)+3)))
				v305 = v301<<(uint(int32(24))%32) + v152
				v306 = v299
				v307 = v300
				v308 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v149)+2)))
				v312 = v308<<(uint(int32(16))%32) + v305
				v313 = v306
				v314 = v307
				v315 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v149)+1)))
				v319 = v315<<(uint(int32(8))%32) + v312
				v320 = v313
				v321 = v314
				v322 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v149))))
				v327 = v319 + v322
				v328 = v320
				v329 = v321
			case 6:
				v283 = v153
				v284 = v154
				v285 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v149)+6)))
				v289 = v285<<(uint(int32(16))%32) + v283
				v290 = v284
				v291 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v149)+5)))
				v295 = v291<<(uint(int32(8))%32) + v289
				v296 = v290
				v297 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v149)+4)))
				v299 = v295 + v297
				v300 = v296
				v301 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v149)+3)))
				v305 = v301<<(uint(int32(24))%32) + v152
				v306 = v299
				v307 = v300
				v308 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v149)+2)))
				v312 = v308<<(uint(int32(16))%32) + v305
				v313 = v306
				v314 = v307
				v315 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v149)+1)))
				v319 = v315<<(uint(int32(8))%32) + v312
				v320 = v313
				v321 = v314
				v322 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v149))))
				v327 = v319 + v322
				v328 = v320
				v329 = v321
			case 7:
				v278 = v154
				v279 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v149)+7)))
				v283 = v279<<(uint(int32(24))%32) + v153
				v284 = v278
				v285 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v149)+6)))
				v289 = v285<<(uint(int32(16))%32) + v283
				v290 = v284
				v291 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v149)+5)))
				v295 = v291<<(uint(int32(8))%32) + v289
				v296 = v290
				v297 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v149)+4)))
				v299 = v295 + v297
				v300 = v296
				v301 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v149)+3)))
				v305 = v301<<(uint(int32(24))%32) + v152
				v306 = v299
				v307 = v300
				v308 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v149)+2)))
				v312 = v308<<(uint(int32(16))%32) + v305
				v313 = v306
				v314 = v307
				v315 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v149)+1)))
				v319 = v315<<(uint(int32(8))%32) + v312
				v320 = v313
				v321 = v314
				v322 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v149))))
				v327 = v319 + v322
				v328 = v320
				v329 = v321
			case 8:
				v273 = v154
				v274 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v149)+8)))
				v278 = v274<<(uint(int32(8))%32) + v273
				v279 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v149)+7)))
				v283 = v279<<(uint(int32(24))%32) + v153
				v284 = v278
				v285 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v149)+6)))
				v289 = v285<<(uint(int32(16))%32) + v283
				v290 = v284
				v291 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v149)+5)))
				v295 = v291<<(uint(int32(8))%32) + v289
				v296 = v290
				v297 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v149)+4)))
				v299 = v295 + v297
				v300 = v296
				v301 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v149)+3)))
				v305 = v301<<(uint(int32(24))%32) + v152
				v306 = v299
				v307 = v300
				v308 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v149)+2)))
				v312 = v308<<(uint(int32(16))%32) + v305
				v313 = v306
				v314 = v307
				v315 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v149)+1)))
				v319 = v315<<(uint(int32(8))%32) + v312
				v320 = v313
				v321 = v314
				v322 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v149))))
				v327 = v319 + v322
				v328 = v320
				v329 = v321
			case 9:
				v268 = v154
				v269 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v149)+9)))
				v273 = v269<<(uint(int32(16))%32) + v268
				v274 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v149)+8)))
				v278 = v274<<(uint(int32(8))%32) + v273
				v279 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v149)+7)))
				v283 = v279<<(uint(int32(24))%32) + v153
				v284 = v278
				v285 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v149)+6)))
				v289 = v285<<(uint(int32(16))%32) + v283
				v290 = v284
				v291 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v149)+5)))
				v295 = v291<<(uint(int32(8))%32) + v289
				v296 = v290
				v297 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v149)+4)))
				v299 = v295 + v297
				v300 = v296
				v301 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v149)+3)))
				v305 = v301<<(uint(int32(24))%32) + v152
				v306 = v299
				v307 = v300
				v308 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v149)+2)))
				v312 = v308<<(uint(int32(16))%32) + v305
				v313 = v306
				v314 = v307
				v315 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v149)+1)))
				v319 = v315<<(uint(int32(8))%32) + v312
				v320 = v313
				v321 = v314
				v322 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v149))))
				v327 = v319 + v322
				v328 = v320
				v329 = v321
			case 10:
				v264 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v149)+10)))
				v268 = v264<<(uint(int32(24))%32) + v154
				v269 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v149)+9)))
				v273 = v269<<(uint(int32(16))%32) + v268
				v274 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v149)+8)))
				v278 = v274<<(uint(int32(8))%32) + v273
				v279 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v149)+7)))
				v283 = v279<<(uint(int32(24))%32) + v153
				v284 = v278
				v285 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v149)+6)))
				v289 = v285<<(uint(int32(16))%32) + v283
				v290 = v284
				v291 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v149)+5)))
				v295 = v291<<(uint(int32(8))%32) + v289
				v296 = v290
				v297 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v149)+4)))
				v299 = v295 + v297
				v300 = v296
				v301 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v149)+3)))
				v305 = v301<<(uint(int32(24))%32) + v152
				v306 = v299
				v307 = v300
				v308 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v149)+2)))
				v312 = v308<<(uint(int32(16))%32) + v305
				v313 = v306
				v314 = v307
				v315 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v149)+1)))
				v319 = v315<<(uint(int32(8))%32) + v312
				v320 = v313
				v321 = v314
				v322 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v149))))
				v327 = v319 + v322
				v328 = v320
				v329 = v321
			default:
				v327 = v152
				v328 = v153
				v329 = v154
			}
		} else {
			if base.Ui32(int32(12)) <= base.Ui32(v46) {
				v160 = v18
				v161 = v46
				v163 = v91
				v164 = v95
				v165 = v93
				for {
					v167 = *(*int32)(unsafe.Add(mBase, uint32(v160)+4))
					v168 = v167 + v164
					v169 = *(*int32)(unsafe.Add(mBase, uint32(v160)))
					v171 = *(*int32)(unsafe.Add(mBase, uint32(v160)+8))
					v172 = v171 + v165
					v174 = int32(4)
					v176 = v169 + v163 - v172 ^ base.I32_rotl(v172, v174)
					v180 = v168 - v176 ^ base.I32_rotl(v176, int32(6))
					v181 = v172 + v168
					v182 = v176 + v181
					v183 = v180 + v182
					v187 = v181 - v180 ^ base.I32_rotl(v180, int32(8))
					v191 = v182 - v187 ^ base.I32_rotl(v187, int32(16))
					v195 = v183 - v191 ^ base.I32_rotl(v191, int32(19))
					v196 = v187 + v183
					v197 = v191 + v196
					v198 = v195 + v197
					v202 = v196 - v195 ^ base.I32_rotl(v195, v174)
					v203 = int32(12)
					v204 = v160 + v203
					v206 = v161 - v203
					if base.Ui32(int32(11)) < base.Ui32(v206) {
						v160 = v204
						v161 = v206
						v163 = v197
						v164 = v198
						v165 = v202
						continue
					} else {
						break
					}
					break
				}
				v209 = v204
				v210 = v206
				v212 = v197
				v213 = v198
				v214 = v202
			} else {
				v209 = v18
				v210 = v46
				v212 = v91
				v213 = v95
				v214 = v93
			}
			switch v210 - int32(1) {
			case 0:
				v261 = v212
				v262 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v209))))
				v327 = v261 + v262
				v328 = v213
				v329 = v214
			case 1:
				v256 = v212
				v257 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v209)+1)))
				v261 = v257<<(uint(int32(8))%32) + v256
				v262 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v209))))
				v327 = v261 + v262
				v328 = v213
				v329 = v214
			case 2:
				v252 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v209)+2)))
				v256 = v252<<(uint(int32(16))%32) + v212
				v257 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v209)+1)))
				v261 = v257<<(uint(int32(8))%32) + v256
				v262 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v209))))
				v327 = v261 + v262
				v328 = v213
				v329 = v214
			case 3:
				v249 = v213
				v250 = *(*int32)(unsafe.Add(mBase, uint32(v209)))
				v327 = v250 + v212
				v328 = v249
				v329 = v214
			case 4:
				v246 = v213
				v247 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v209)+4)))
				v249 = v246 + v247
				v250 = *(*int32)(unsafe.Add(mBase, uint32(v209)))
				v327 = v250 + v212
				v328 = v249
				v329 = v214
			case 5:
				v241 = v213
				v242 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v209)+5)))
				v246 = v242<<(uint(int32(8))%32) + v241
				v247 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v209)+4)))
				v249 = v246 + v247
				v250 = *(*int32)(unsafe.Add(mBase, uint32(v209)))
				v327 = v250 + v212
				v328 = v249
				v329 = v214
			case 6:
				v237 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v209)+6)))
				v241 = v237<<(uint(int32(16))%32) + v213
				v242 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v209)+5)))
				v246 = v242<<(uint(int32(8))%32) + v241
				v247 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v209)+4)))
				v249 = v246 + v247
				v250 = *(*int32)(unsafe.Add(mBase, uint32(v209)))
				v327 = v250 + v212
				v328 = v249
				v329 = v214
			case 7:
				v232 = v214
				v233 = *(*int32)(unsafe.Add(mBase, uint32(v209)))
				v235 = *(*int32)(unsafe.Add(mBase, uint32(v209)+4))
				v327 = v233 + v212
				v328 = v235 + v213
				v329 = v232
			case 8:
				v227 = v214
				v228 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v209)+8)))
				v232 = v228<<(uint(int32(8))%32) + v227
				v233 = *(*int32)(unsafe.Add(mBase, uint32(v209)))
				v235 = *(*int32)(unsafe.Add(mBase, uint32(v209)+4))
				v327 = v233 + v212
				v328 = v235 + v213
				v329 = v232
			case 9:
				v222 = v214
				v223 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v209)+9)))
				v227 = v223<<(uint(int32(16))%32) + v222
				v228 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v209)+8)))
				v232 = v228<<(uint(int32(8))%32) + v227
				v233 = *(*int32)(unsafe.Add(mBase, uint32(v209)))
				v235 = *(*int32)(unsafe.Add(mBase, uint32(v209)+4))
				v327 = v233 + v212
				v328 = v235 + v213
				v329 = v232
			case 10:
				v218 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v209)+10)))
				v222 = v218<<(uint(int32(24))%32) + v214
				v223 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v209)+9)))
				v227 = v223<<(uint(int32(16))%32) + v222
				v228 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v209)+8)))
				v232 = v228<<(uint(int32(8))%32) + v227
				v233 = *(*int32)(unsafe.Add(mBase, uint32(v209)))
				v235 = *(*int32)(unsafe.Add(mBase, uint32(v209)+4))
				v327 = v233 + v212
				v328 = v235 + v213
				v329 = v232
			default:
				v327 = v212
				v328 = v213
				v329 = v214
			}
		}
		v332 = int32(14)
		v334 = v328 ^ v329 - base.I32_rotl(v328, v332)
		v338 = v334 ^ v327 - base.I32_rotl(v334, int32(11))
		v342 = v338 ^ v328 - base.I32_rotl(v338, int32(25))
		v346 = v342 ^ v334 - base.I32_rotl(v342, int32(16))
		v350 = v346 ^ v338 - base.I32_rotl(v346, int32(4))
		v354 = v350 ^ v342 - base.I32_rotl(v350, v332)
		v364 = F_Int64GetDatum(m, base.I64_extend_i32_u(v354)<<(uint(int64(32))%64)|base.I64_extend_i32_u(v346^v354-base.I32_rotl(v354, int32(24))))
		mBase = m.M
		v365 = m.ExcPending
		if v365 != 0 {
			return int32(0)
		} else {
			v366 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
			if v366 != v7 {
				F_pfree(m, v7)
				mBase = m.M
				v369 = m.ExcPending
				if v369 != 0 {
					return int32(0)
				} else {
					return v364
				}
			} else {
				return v364
			}
		}
	}
}
func F_heapgettup_pagemode(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) {
	mBase := m.M
	_ = mBase
	var v5 int32
	_ = v5
	var v17 int32
	_ = v17
	var v19 int32
	_ = v19
	var v23 int32
	_ = v23
	var v26 int32
	_ = v26
	var v30 int32
	_ = v30
	var v36 int32
	_ = v36
	var v38 int32
	_ = v38
	var v44 int32
	_ = v44
	var v45 int32
	_ = v45
	var v46 int32
	_ = v46
	var v49 int32
	_ = v49
	var v52 int32
	_ = v52
	var v55 int32
	_ = v55
	var v56 int32
	_ = v56
	var v57 int32
	_ = v57
	var v58 int32
	_ = v58
	var v63 int32
	_ = v63
	var v65 int32
	_ = v65
	var v66 int32
	_ = v66
	var v69 int32
	_ = v69
	var v78 int32
	_ = v78
	var v79 int32
	_ = v79
	var v83 int32
	_ = v83
	var v84 int32
	_ = v84
	var v88 int32
	_ = v88
	var v94 int32
	_ = v94
	var v96 int32
	_ = v96
	var v102 int32
	_ = v102
	var v103 int32
	_ = v103
	var v107 int32
	_ = v107
	var v108 int32
	_ = v108
	var v113 int32
	_ = v113
	var v125 int32
	_ = v125
	var v127 int32
	_ = v127
	var v138 int32
	_ = v138
	var v143 int32
	_ = v143
	var v144 int32
	_ = v144
	var v149 int32
	_ = v149
	var v158 int32
	_ = v158
	var v159 int32
	_ = v159
	var v164 int32
	_ = v164
	var v165 int32
	_ = v165
	var v176 int32
	_ = v176
	var v179 int32
	_ = v179
	var v182 int32
	_ = v182
	var v183 int32
	_ = v183
	var v184 int32
	_ = v184
	var v187 int32
	_ = v187
	var v188 int32
	_ = v188
	var v189 int32
	_ = v189
	var v190 int32
	_ = v190
	var v196 int32
	_ = v196
	var v204 int32
	_ = v204
	var v217 int32
	_ = v217
	var v224 int32
	_ = v224
	var v235 int32
	_ = v235
	var v239 int32
	_ = v239
	v5 = int32(0)
	v17 = m.G0
	v19 = v17 - int32(16)
	m.G0 = v19
	v23 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+48)))
	if v23 == int32(1) {
		goto L3
	} else {
		goto L4
	}
L1:
	;
	v63 = v58
	v65 = v55
	v66 = v56
	v69 = v57
	goto L12
L2:
	;
	v55 = v52
	v56 = v46
	v57 = v44
	v58 = int32(1)
	goto L1
L3:
	;
	v26 = *(*int32)(unsafe.Add(mBase, uint32(l0)+56))
	if v26 < int32(0) {
		goto L7
	} else {
		goto L8
	}
L4:
	;
	goto L5
L5:
	;
	v55 = v5
	v56 = v5
	v57 = v5
	v58 = int32(0)
	goto L1
L6:
	;
	v45 = *(*int32)(unsafe.Add(mBase, uint32(l0)+100))
	v46 = v45 + l1
	if l1 != int32(1) {
		v52 = v45
		goto L2
	} else {
		goto L10
	}
L7:
	;
	v30 = *(*int32)(unsafe.Add(mBase, _c_F_heapgettup_pagemode[0]))
	v36 = *(*int32)(unsafe.Add(mBase, uint32(v30+(v26^int32(-1))<<(uint(int32(2))%32))))
	v44 = v36
	goto L6
L8:
	;
	goto L9
L9:
	;
	v38 = *(*int32)(unsafe.Add(mBase, _c_F_heapgettup_pagemode[1]))
	v44 = v38 + v26<<(uint(int32(13))%32) + int32(-8192)
	goto L6
L10:
	;
	v49 = *(*int32)(unsafe.Add(mBase, uint32(l0)+104))
	v52 = v49 - v46
	goto L2
L11:
	;
	m.G0 = v19 + int32(16)
	return
L12:
	;
	if v63 == int32(0) {
		goto L15
	} else {
		goto L16
	}
L13:
	;
	v235 = int32(-1)
	*(*int32)(unsafe.Add(mBase, uint32(l0)+92)) = v235
	*(*int32)(unsafe.Add(mBase, uint32(l0)+52)) = v235
	v239 = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(l0)+80)) = v239
	*(*uint8)(unsafe.Add(mBase, uint32(l0)+48)) = uint8(v239)
	goto L11
L14:
	;
	goto L13
L15:
	;
	F_heap_fetch_next_buffer(m, l0, l1)
	mBase = m.M
	v78 = m.ExcPending
	if v78 != 0 {
		goto L18
	} else {
		goto L19
	}
L16:
	;
	goto L17
L17:
	;
	if v65 != 0 {
		goto L29
	} else {
		goto L30
	}
L18:
	;
	return
L19:
	;
	v79 = *(*int32)(unsafe.Add(mBase, uint32(l0)+56))
	if v79 == int32(0) {
		goto L14
	} else {
		goto L20
	}
L20:
	;
	F_heap_prepare_pagescan(m, l0)
	mBase = m.M
	v83 = m.ExcPending
	if v83 != 0 {
		goto L18
	} else {
		goto L21
	}
L21:
	;
	v84 = *(*int32)(unsafe.Add(mBase, uint32(l0)+56))
	if v84 < int32(0) {
		goto L23
	} else {
		goto L24
	}
L22:
	;
	v103 = *(*int32)(unsafe.Add(mBase, uint32(l0)+52))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+68)) = base.I32_rotr(v103, int32(16))
	v107 = *(*int32)(unsafe.Add(mBase, uint32(l0)+104))
	v108 = int32(1)
	if l1 != v108 {
		goto L26
	} else {
		goto L27
	}
L23:
	;
	v88 = *(*int32)(unsafe.Add(mBase, _c_F_heapgettup_pagemode[0]))
	v94 = *(*int32)(unsafe.Add(mBase, uint32(v88+(v84^int32(-1))<<(uint(int32(2))%32))))
	v102 = v94
	goto L22
L24:
	;
	goto L25
L25:
	;
	v96 = *(*int32)(unsafe.Add(mBase, _c_F_heapgettup_pagemode[1]))
	v102 = v96 + v84<<(uint(int32(13))%32) + int32(-8192)
	goto L22
L26:
	;
	v113 = v107 - v108
	goto L28
L27:
	;
	v113 = int32(0)
	goto L28
L28:
	;
	v63 = int32(1)
	v65 = v107
	v66 = v113
	v69 = v102
	goto L12
L29:
	;
	v125 = v65
	v127 = v66
	goto L32
L30:
	;
	v224 = v65
	goto L31
L31:
	;
	v63 = int32(0)
	v65 = v224
	goto L12
L32:
	;
	v138 = int32(*(*uint16)(unsafe.Add(mBase, uint32(l0+int32(108)+v127<<(uint(int32(1))%32)))))
	v143 = v138<<(uint(int32(2))%32) + (v69 + int32(24)) - int32(4)
	v144 = *(*int32)(unsafe.Add(mBase, uint32(v143)))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+80)) = v69 + v144&int32(_a_F_heapgettup_pagemode_0)
	v149 = *(*int32)(unsafe.Add(mBase, uint32(v143)))
	*(*uint16)(unsafe.Add(mBase, uint32(l0)+72)) = uint16(v138)
	*(*int32)(unsafe.Add(mBase, uint32(l0)+64)) = int32(base.Ui32(v149) >> (uint(int32(17)) % 32))
	if l3 == int32(0) {
		v204 = v66
		goto L35
	} else {
		goto L36
	}
L33:
	;
	v224 = v217
	goto L31
L34:
	;
	v217 = v125 - int32(1)
	if v217 != 0 {
		v125 = v217
		v127 = l1 + v127
		goto L32
	} else {
		goto L46
	}
L35:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+100)) = v204
	goto L11
L36:
	;
	if l2 == int32(0) {
		v204 = v66
		goto L35
	} else {
		goto L37
	}
L37:
	;
	v158 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v159 = *(*int32)(unsafe.Add(mBase, uint32(v158)+52))
	v164 = l3
	v165 = l2
	goto L38
L38:
	;
	v176 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v164))))
	if v176&int32(1) != 0 {
		goto L34
	} else {
		goto L40
	}
L39:
	;
	v204 = v127
	goto L35
L40:
	;
	v179 = int32(*(*int16)(unsafe.Add(mBase, uint32(v164)+4)))
	v182 = F_heap_getattr_1(m, l0-int32(-64), v179, v159, v19+int32(15))
	mBase = m.M
	v183 = m.ExcPending
	if v183 != 0 {
		goto L18
	} else {
		goto L41
	}
L41:
	;
	v184 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v19)+15)))
	if v184 != 0 {
		goto L34
	} else {
		goto L42
	}
L42:
	;
	v187 = *(*int32)(unsafe.Add(mBase, uint32(v164)+12))
	v188 = *(*int32)(unsafe.Add(mBase, uint32(v164)+44))
	v189 = F_FunctionCall2Coll(m, v164+int32(16), v187, v182, v188)
	mBase = m.M
	v190 = m.ExcPending
	if v190 != 0 {
		goto L18
	} else {
		goto L43
	}
L43:
	;
	if v189 == int32(0) {
		goto L34
	} else {
		goto L44
	}
L44:
	;
	v196 = v165 - int32(1)
	if v196 != 0 {
		v164 = v164 + int32(48)
		v165 = v196
		goto L38
	} else {
		goto L45
	}
L45:
	;
	goto L39
L46:
	;
	goto L33
}
func F_hemdist_1(m *base.Module, l0 int32, l1 int32, l2 int32) int32 {
	mBase := m.M
	_ = mBase
	var v4 int32
	_ = v4
	var v10 int32
	_ = v10
	var v11 int32
	_ = v11
	var v12 int32
	_ = v12
	var v13 int32
	_ = v13
	var v19 int32
	_ = v19
	var v26 int32
	_ = v26
	var v29 int32
	_ = v29
	var v31 int32
	_ = v31
	var v35 int32
	_ = v35
	var v36 int32
	_ = v36
	var v71 int32
	_ = v71
	var v75 int32
	_ = v75
	var v80 int32
	_ = v80
	var v87 int32
	_ = v87
	var v90 int32
	_ = v90
	var v92 int32
	_ = v92
	var v96 int32
	_ = v96
	var v97 int32
	_ = v97
	var v132 int32
	_ = v132
	var v136 int32
	_ = v136
	var v144 int32
	_ = v144
	var v145 int32
	_ = v145
	var v147 int32
	_ = v147
	var v148 int32
	_ = v148
	var v150 int32
	_ = v150
	var v153 int32
	_ = v153
	var v154 int32
	_ = v154
	var v158 int32
	_ = v158
	var v162 int32
	_ = v162
	var v165 int32
	_ = v165
	var v166 int32
	_ = v166
	var v167 int32
	_ = v167
	var v175 int32
	_ = v175
	var v177 int32
	_ = v177
	var v179 int32
	_ = v179
	var v181 int32
	_ = v181
	var v183 int32
	_ = v183
	var v184 int32
	_ = v184
	var v193 int32
	_ = v193
	var v194 int32
	_ = v194
	var v195 int32
	_ = v195
	var v197 int32
	_ = v197
	var v199 int32
	_ = v199
	var v200 int32
	_ = v200
	var v209 int32
	_ = v209
	var v211 int32
	_ = v211
	var v213 int32
	_ = v213
	var v223 int32
	_ = v223
	v4 = int32(0)
	v10 = *(*int32)(unsafe.Add(mBase, uint32(l1)+4))
	v11 = int32(4)
	v12 = v10 & v11
	v13 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+4)))
	if v13&v11 != 0 {
		if v12 != 0 {
			return int32(0)
		} else {
			v19 = l2 << (uint(int32(3)) % 32)
			if l2 <= int32(0) {
				return v19
			} else {
				v26 = l1 + int32(8)
				v29 = int32(0)
				v31 = v4
				for {
					v35 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v26))))
					v36 = int32(1)
					v71 = v35&v36 + v29 + int32(base.Ui32(v35)>>(uint(int32(7))%32)) + int32(base.Ui32(v35)>>(uint(v36)%32))&v36 + int32(base.Ui32(v35)>>(uint(int32(2))%32))&v36 + int32(base.Ui32(v35)>>(uint(int32(3))%32))&v36 + int32(base.Ui32(v35)>>(uint(int32(4))%32))&v36 + int32(base.Ui32(v35)>>(uint(int32(5))%32))&v36 + int32(base.Ui32(v35)>>(uint(int32(6))%32))&v36
					v75 = v31 + v36
					if v75 != l2 {
						v26 = v26 + v36
						v29 = v71
						v31 = v75
						continue
					} else {
						break
					}
					break
				}
				return v19 - v71
			}
		}
	} else {
		if v12 != 0 {
			v80 = l2 << (uint(int32(3)) % 32)
			if l2 <= int32(0) {
				return v80
			} else {
				v87 = l0 + int32(8)
				v90 = int32(0)
				v92 = v4
				for {
					v96 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v87))))
					v97 = int32(1)
					v132 = v96&v97 + v90 + int32(base.Ui32(v96)>>(uint(int32(7))%32)) + int32(base.Ui32(v96)>>(uint(v97)%32))&v97 + int32(base.Ui32(v96)>>(uint(int32(2))%32))&v97 + int32(base.Ui32(v96)>>(uint(int32(3))%32))&v97 + int32(base.Ui32(v96)>>(uint(int32(4))%32))&v97 + int32(base.Ui32(v96)>>(uint(int32(5))%32))&v97 + int32(base.Ui32(v96)>>(uint(int32(6))%32))&v97
					v136 = v92 + v97
					if v136 != l2 {
						v87 = v87 + v97
						v90 = v132
						v92 = v136
						continue
					} else {
						break
					}
					break
				}
				return v80 - v132
			}
		} else {
			if l2 <= int32(0) {
				return int32(0)
			} else {
				v144 = int32(8)
				v145 = l1 + v144
				v147 = l0 + v144
				v148 = int32(1)
				v150 = l2 << (uint(int32(3)) % 32)
				if v150 <= v148 {
					v153 = v148
				} else {
					v153 = v150
				}
				v154 = int32(1)
				if v153 == v154 {
					v158 = int32(0)
					v199 = v158
					v200 = v158
				} else {
					v162 = int32(0)
					v165 = v162
					v166 = v162
					v167 = v162
					for {
						v175 = int32(base.Ui32(v166) >> (uint(int32(3)) % 32))
						v177 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v145+v175))))
						v179 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v175+v147))))
						v181 = base.I32_extend8_s(v177 ^ v179)
						v183 = v166 & int32(6)
						v184 = int32(1)
						v193 = int32(base.Ui32(v181)>>(uint(v183|v184)%32))&v184 + (int32(base.Ui32(v181)>>(uint(v183)%32))&v184 + v165)
						v194 = int32(2)
						v195 = v166 + v194
						v197 = v167 + v194
						if v197 != v153&int32(2147483640) {
							v165 = v193
							v166 = v195
							v167 = v197
							continue
						} else {
							break
						}
						break
					}
					v199 = v193
					v200 = v195
				}
				if v153&v154 != 0 {
					v209 = int32(base.Ui32(v200) >> (uint(int32(3)) % 32))
					v211 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v145+v209))))
					v213 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v209+v147))))
					v223 = int32(base.Ui32(base.I32_extend8_s(v211^v213))>>(uint(v200&int32(7))%32))&int32(1) + v199
				} else {
					v223 = v199
				}
				return v223
			}
		}
	}
}
func F_hex_decode_safe(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) int64 {
	mBase := m.M
	_ = mBase
	var v10 int64
	_ = v10
	var v11 int32
	_ = v11
	var v13 int32
	_ = v13
	var v15 int32
	_ = v15
	var v17 int32
	_ = v17
	var v23 int32
	_ = v23
	var v27 int32
	_ = v27
	var v29 int32
	_ = v29
	var v39 int32
	_ = v39
	var v45 int32
	_ = v45
	var v49 int32
	_ = v49
	var v52 int32
	_ = v52
	var v57 int32
	_ = v57
	var v58 int32
	_ = v58
	var v59 int32
	_ = v59
	var v66 int32
	_ = v66
	var v71 int32
	_ = v71
	var v73 int32
	_ = v73
	var v75 int32
	_ = v75
	var v76 int32
	_ = v76
	var v81 int32
	_ = v81
	var v85 int32
	_ = v85
	var v90 int32
	_ = v90
	var v91 int32
	_ = v91
	var v96 int32
	_ = v96
	var v100 int32
	_ = v100
	var v101 int32
	_ = v101
	var v106 int32
	_ = v106
	var v107 int32
	_ = v107
	var v108 int32
	_ = v108
	var v113 int32
	_ = v113
	var v118 int32
	_ = v118
	var v121 int32
	_ = v121
	var v124 int32
	_ = v124
	var v126 int32
	_ = v126
	var v134 int32
	_ = v134
	var v149 int64
	_ = v149
	v10 = int64(0)
	v11 = m.G0
	v13 = v11 - int32(32)
	m.G0 = v13
	v15 = l0 + l1
	if base.Ui32(v15) <= base.Ui32(l0) {
		v134 = l2
		goto L2
	} else {
		goto L3
	}
L1:
	;
	m.G0 = v13 + int32(32)
	return v149
L2:
	;
	v149 = base.I64_extend_i32_s(v134 - l2)
	goto L1
L3:
	;
	v17 = l0
	v23 = l2
	goto L4
L4:
	;
	v27 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v17))))
	v29 = v27 - int32(9)
	if base.Ui32(int32(23)) < base.Ui32(v29) {
		goto L6
	} else {
		goto L7
	}
L5:
	;
	v134 = v124
	goto L2
L6:
	;
	if base.Ui32(v27) <= base.Ui32(int32(126)) {
		goto L11
	} else {
		goto L12
	}
L7:
	;
	if int32(1)<<(uint(v29)%32)&int32(_a_F_hex_decode_safe_0) == int32(0) {
		goto L6
	} else {
		goto L8
	}
L8:
	;
	v39 = v17 + int32(1)
	if base.Ui32(v39) < base.Ui32(v15) {
		v17 = v39
		goto L4
	} else {
		goto L9
	}
L9:
	;
	v134 = v23
	goto L2
L10:
	;
	v73 = v17 + int32(1)
	if base.Ui32(v15) <= base.Ui32(v73) {
		goto L22
	} else {
		goto L23
	}
L11:
	;
	v45 = int32(*(*int8)(unsafe.Add(mBase, uint32(v27)+uint32(_c_F_hex_decode_safe[0]))))
	if int32(0) <= v45 {
		goto L10
	} else {
		goto L14
	}
L12:
	;
	goto L13
L13:
	;
	v49 = F_errsave_start(m, l3)
	mBase = m.M
	v52 = m.ExcPending
	if v52 != 0 {
		goto L15
	} else {
		goto L16
	}
L14:
	;
	goto L13
L15:
	;
	return int64(0)
L16:
	;
	if v49 == int32(0) {
		v149 = v10
		goto L1
	} else {
		goto L17
	}
L17:
	;
	F_errcode(m, int32(50856066))
	mBase = m.M
	v57 = m.ExcPending
	if v57 != 0 {
		goto L15
	} else {
		goto L18
	}
L18:
	;
	v58 = F_pg_mblen_range(m, v17, v15)
	mBase = m.M
	v59 = m.ExcPending
	if v59 != 0 {
		goto L15
	} else {
		goto L19
	}
L19:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v13)+20)) = v17
	*(*int32)(unsafe.Add(mBase, uint32(v13)+16)) = v58
	F_errmsg(m, int32(_a_F_hex_decode_safe_1), v13+int32(16))
	mBase = m.M
	v66 = m.ExcPending
	if v66 != 0 {
		goto L15
	} else {
		goto L20
	}
L20:
	;
	F_errsave_finish(m, l3, int32(_a_F_hex_decode_safe_2), int32(239), int32(_a_F_hex_decode_safe_3))
	mBase = m.M
	v71 = m.ExcPending
	if v71 != 0 {
		goto L15
	} else {
		goto L21
	}
L21:
	;
	v149 = v10
	goto L1
L22:
	;
	v75 = F_errsave_start(m, l3)
	mBase = m.M
	v76 = m.ExcPending
	if v76 != 0 {
		goto L15
	} else {
		goto L25
	}
L23:
	;
	goto L24
L24:
	;
	v91 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v73))))
	if base.Ui32(v91) <= base.Ui32(int32(126)) {
		goto L31
	} else {
		goto L32
	}
L25:
	;
	if v75 == int32(0) {
		v149 = v10
		goto L1
	} else {
		goto L26
	}
L26:
	;
	F_errcode(m, int32(50856066))
	mBase = m.M
	v81 = m.ExcPending
	if v81 != 0 {
		goto L15
	} else {
		goto L27
	}
L27:
	;
	F_errmsg(m, int32(_a_F_hex_decode_safe_4), int32(0))
	mBase = m.M
	v85 = m.ExcPending
	if v85 != 0 {
		goto L15
	} else {
		goto L28
	}
L28:
	;
	F_errsave_finish(m, l3, int32(_a_F_hex_decode_safe_2), int32(244), int32(_a_F_hex_decode_safe_3))
	mBase = m.M
	v90 = m.ExcPending
	if v90 != 0 {
		goto L15
	} else {
		goto L29
	}
L29:
	;
	v149 = v10
	goto L1
L30:
	;
	v121 = v96 | v45<<(uint(int32(4))%32)
	*(*uint8)(unsafe.Add(mBase, uint32(v23))) = uint8(v121)
	v124 = v23 + int32(1)
	v126 = v17 + int32(2)
	if base.Ui32(v126) < base.Ui32(v15) {
		v17 = v126
		v23 = v124
		goto L4
	} else {
		goto L41
	}
L31:
	;
	v96 = int32(*(*int8)(unsafe.Add(mBase, uint32(v91)+uint32(_c_F_hex_decode_safe[0]))))
	if int32(0) <= v96 {
		goto L30
	} else {
		goto L34
	}
L32:
	;
	goto L33
L33:
	;
	v100 = F_errsave_start(m, l3)
	mBase = m.M
	v101 = m.ExcPending
	if v101 != 0 {
		goto L15
	} else {
		goto L35
	}
L34:
	;
	goto L33
L35:
	;
	if v100 == int32(0) {
		v149 = v10
		goto L1
	} else {
		goto L36
	}
L36:
	;
	F_errcode(m, int32(50856066))
	mBase = m.M
	v106 = m.ExcPending
	if v106 != 0 {
		goto L15
	} else {
		goto L37
	}
L37:
	;
	v107 = F_pg_mblen_range(m, v73, v15)
	mBase = m.M
	v108 = m.ExcPending
	if v108 != 0 {
		goto L15
	} else {
		goto L38
	}
L38:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v13)+4)) = v73
	*(*int32)(unsafe.Add(mBase, uint32(v13))) = v107
	F_errmsg(m, int32(_a_F_hex_decode_safe_1), v13)
	mBase = m.M
	v113 = m.ExcPending
	if v113 != 0 {
		goto L15
	} else {
		goto L39
	}
L39:
	;
	F_errsave_finish(m, l3, int32(_a_F_hex_decode_safe_2), int32(249), int32(_a_F_hex_decode_safe_3))
	mBase = m.M
	v118 = m.ExcPending
	if v118 != 0 {
		goto L15
	} else {
		goto L40
	}
L40:
	;
	v149 = v10
	goto L1
L41:
	;
	goto L5
}
func F_hex_enc_len(m *base.Module, l0 int32, l1 int32) int64 {
	return base.I64_extend_i32_u(l1) << (uint(int64(1)) % 64)
}
func F_hex_encode(m *base.Module, l0 int32, l1 int32, l2 int32) int64 {
	mBase := m.M
	_ = mBase
	var v7 int32
	_ = v7
	var v10 int32
	_ = v10
	var v11 int32
	_ = v11
	var v13 int32
	_ = v13
	var v14 int32
	_ = v14
	var v17 int32
	_ = v17
	var v18 int32
	_ = v18
	var v22 int32
	_ = v22
	var v25 int32
	_ = v25
	var v27 int32
	_ = v27
	var v29 int32
	_ = v29
	var v31 int32
	_ = v31
	var v33 int32
	_ = v33
	var v41 int32
	_ = v41
	var v43 int32
	_ = v43
	var v47 int32
	_ = v47
	var v48 int32
	_ = v48
	var v52 int32
	_ = v52
	var v54 int32
	_ = v54
	var v59 int32
	_ = v59
	var v61 int32
	_ = v61
	var v66 int32
	_ = v66
	var v68 int32
	_ = v68
	var v73 int32
	_ = v73
	var v78 int32
	_ = v78
	v7 = l0 + l1
	if base.Ui32(v7) <= base.Ui32(l0) {
	} else {
		v10 = l1 & int32(3)
		if v10 != 0 {
			v11 = l0
			v13 = l2
			v14 = int32(0)
			for {
				v17 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v11))))
				v18 = int32(1)
				v22 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v17<<(uint(v18)%32))+uint32(_c_F_hex_encode[0]))))
				*(*uint16)(unsafe.Add(mBase, uint32(v13))) = uint16(v22)
				v25 = v13 + int32(2)
				v27 = v11 + v18
				v29 = v14 + v18
				if v29 != v10 {
					v11 = v27
					v13 = v25
					v14 = v29
					continue
				} else {
					break
				}
				break
			}
			v31 = v27
			v33 = v25
		} else {
			v31 = l0
			v33 = l2
		}
		if base.Ui32(l1-int32(1)) < base.Ui32(int32(3)) {
		} else {
			v41 = v31
			v43 = v33
			for {
				v47 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v41))))
				v48 = int32(1)
				v52 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v47<<(uint(v48)%32))+uint32(_c_F_hex_encode[0]))))
				*(*uint16)(unsafe.Add(mBase, uint32(v43))) = uint16(v52)
				v54 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v41)+1)))
				v59 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v54<<(uint(v48)%32))+uint32(_c_F_hex_encode[0]))))
				*(*uint16)(unsafe.Add(mBase, uint32(v43)+2)) = uint16(v59)
				v61 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v41)+2)))
				v66 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v61<<(uint(v48)%32))+uint32(_c_F_hex_encode[0]))))
				*(*uint16)(unsafe.Add(mBase, uint32(v43)+4)) = uint16(v66)
				v68 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v41)+3)))
				v73 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v68<<(uint(v48)%32))+uint32(_c_F_hex_encode[0]))))
				*(*uint16)(unsafe.Add(mBase, uint32(v43)+6)) = uint16(v73)
				v78 = v41 + int32(4)
				if v78 != v7 {
					v41 = v78
					v43 = v43 + int32(8)
					continue
				} else {
					break
				}
				break
			}
		}
	}
	return base.I64_extend_i32_u(l1) << (uint(int64(1)) % 64)
}
func F_hindi_UTF_8_create_env(m *base.Module) int32 {
	var v1 int32
	_ = v1
	var v3 int32
	_ = v3
	var v6 int32
	_ = v6
	v1 = int32(0)
	v3 = F_SN_create_env(m, v1, v1)
	v6 = m.ExcPending
	if v6 != 0 {
		return int32(0)
	} else {
		return v3
	}
}
func F_hmac_free(m *base.Module, l0 int32) {
	mBase := m.M
	_ = mBase
	var v4 int32
	_ = v4
	var v5 int32
	_ = v5
	var v6 int32
	_ = v6
	var v7 int32
	_ = v7
	var v8 int32
	_ = v8
	var v9 int32
	_ = v9
	var v11 int32
	_ = v11
	var v12 int32
	_ = v12
	var v14 int32
	_ = v14
	var v15 int32
	_ = v15
	var v17 int32
	_ = v17
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
	v4 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
	v5 = *(*int32)(unsafe.Add(mBase, uint32(v4)+4))
	v6 = m.T0[v5].(func(*base.Module, int32) int32)(m, v4)
	mBase = m.M
	v7 = m.ExcPending
	if v7 != 0 {
		return
	} else {
		v8 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
		v9 = *(*int32)(unsafe.Add(mBase, uint32(v8)+20))
		m.T0[v9].(func(*base.Module, int32))(m, v8)
		mBase = m.M
		v11 = m.ExcPending
		if v11 != 0 {
			return
		} else {
			v12 = *(*int32)(unsafe.Add(mBase, uint32(l0)+32))
			v14 = F___memset(m, v12, int32(0), v6)
			mBase = m.M
			v15 = *(*int32)(unsafe.Add(mBase, uint32(l0)+36))
			v17 = F___memset(m, v15, int32(0), v6)
			mBase = m.M
			v18 = *(*int32)(unsafe.Add(mBase, uint32(l0)+32))
			F_pfree(m, v18)
			mBase = m.M
			v20 = m.ExcPending
			if v20 != 0 {
				return
			} else {
				v21 = *(*int32)(unsafe.Add(mBase, uint32(l0)+36))
				F_pfree(m, v21)
				mBase = m.M
				v23 = m.ExcPending
				if v23 != 0 {
					return
				} else {
					F_pfree(m, l0)
					mBase = m.M
					v25 = m.ExcPending
					if v25 != 0 {
						return
					} else {
						return
					}
				}
			}
		}
	}
}
func F_htons(m *base.Module, l0 int32) int32 {
	var v2 int32
	_ = v2
	v2 = int32(8)
	return (l0<<(uint(v2)%32) | int32(base.Ui32(l0)>>(uint(v2)%32))) & int32(_a_F_htons_0)
}
func F_hyphenate(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) int32 {
	mBase := m.M
	_ = mBase
	var v5 int32
	_ = v5
	var v16 int32
	_ = v16
	var v18 int32
	_ = v18
	var v22 int32
	_ = v22
	var v23 int32
	_ = v23
	var v25 int32
	_ = v25
	var v26 int32
	_ = v26
	var v38 int32
	_ = v38
	var v41 int32
	_ = v41
	var v42 int32
	_ = v42
	var v48 int32
	_ = v48
	var v63 int32
	_ = v63
	var v64 int32
	_ = v64
	var v69 int32
	_ = v69
	var v72 int32
	_ = v72
	var v74 int32
	_ = v74
	var v81 int32
	_ = v81
	var v82 int32
	_ = v82
	var v83 int32
	_ = v83
	var v85 int32
	_ = v85
	var v86 int32
	_ = v86
	var v89 int32
	_ = v89
	var v90 int32
	_ = v90
	var v91 int32
	_ = v91
	var v97 int32
	_ = v97
	var v99 int32
	_ = v99
	var v100 int32
	_ = v100
	var v101 int32
	_ = v101
	var v103 int32
	_ = v103
	var v104 int32
	_ = v104
	var v105 int32
	_ = v105
	var v106 int32
	_ = v106
	var v107 int32
	_ = v107
	var v109 int32
	_ = v109
	var v114 int32
	_ = v114
	var v119 int32
	_ = v119
	var v120 int32
	_ = v120
	var v121 int32
	_ = v121
	var v122 int32
	_ = v122
	var v124 int32
	_ = v124
	var v125 int32
	_ = v125
	var v128 int32
	_ = v128
	var v129 int32
	_ = v129
	var v132 int32
	_ = v132
	var v136 int32
	_ = v136
	var v140 int32
	_ = v140
	var v141 int32
	_ = v141
	var v145 int32
	_ = v145
	var v146 int32
	_ = v146
	var v157 int32
	_ = v157
	var v161 int32
	_ = v161
	var v167 int32
	_ = v167
	var v168 int32
	_ = v168
	var v169 int32
	_ = v169
	var v172 int32
	_ = v172
	var v173 int32
	_ = v173
	var v176 int32
	_ = v176
	var v177 int32
	_ = v177
	var v178 int32
	_ = v178
	var v180 int32
	_ = v180
	var v182 int32
	_ = v182
	var v183 int32
	_ = v183
	var v184 int32
	_ = v184
	var v186 int32
	_ = v186
	var v187 int32
	_ = v187
	var v188 int32
	_ = v188
	var v189 int32
	_ = v189
	var v190 int32
	_ = v190
	var v192 int32
	_ = v192
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
	var v208 int32
	_ = v208
	var v209 int32
	_ = v209
	var v219 int32
	_ = v219
	var v220 int32
	_ = v220
	var v224 int32
	_ = v224
	var v225 int32
	_ = v225
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
	var v235 int32
	_ = v235
	var v236 int32
	_ = v236
	var v239 int32
	_ = v239
	var v243 int32
	_ = v243
	var v244 int32
	_ = v244
	var v258 int32
	_ = v258
	var v259 int32
	_ = v259
	var v261 int32
	_ = v261
	v5 = int32(0)
	v16 = int32(*(*int8)(unsafe.Add(mBase, uint32(l1))))
	if l3 != 0 {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	v18 = l2
	goto L3
L2:
	;
	v18 = v5
	goto L3
L3:
	;
	if v18 == int32(0) {
		goto L4
	} else {
		goto L5
	}
L4:
	;
	if v16 != 0 {
		goto L7
	} else {
		goto L8
	}
L5:
	;
	goto L6
L6:
	;
	v69 = l3 + v16<<(uint(int32(3))%32)
	v72 = *(*int32)(unsafe.Add(mBase, uint32(v69-int32(380))))
	v74 = v72 + int32(1)
	if base.Ui32(v74) < base.Ui32(int32(2)) {
		goto L13
	} else {
		goto L14
	}
L7:
	;
	v22 = l0
	v23 = l1
	v25 = int32(0)
	v26 = v16
	goto L10
L8:
	;
	v48 = l0
	v63 = int32(1)
	goto L9
L9:
	;
	v64 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v48))) = uint8(v64)
	return v63
L10:
	;
	*(*uint8)(unsafe.Add(mBase, uint32(v22))) = uint8(v26)
	v38 = int32(1)
	v41 = v22 + v38
	v42 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v23)+1)))
	if v42 != 0 {
		v22 = v41
		v23 = v23 + v38
		v25 = v25 + v38
		v26 = v42
		goto L10
	} else {
		goto L12
	}
L11:
	;
	v48 = v41
	v63 = v25 + int32(2)
	goto L9
L12:
	;
	goto L11
L13:
	;
	return int32(0)
L14:
	;
	goto L15
L15:
	;
	v81 = *(*int32)(unsafe.Add(mBase, uint32(v69-int32(384))))
	v82 = int32(1)
	v83 = v81 - v82
	v85 = int32(base.Ui32(v74) >> (uint(v82) % 32))
	v86 = v83 + v85
	v89 = l2 + v86<<(uint(int32(3))%32)
	v90 = *(*int32)(unsafe.Add(mBase, uint32(v89)+4))
	v91 = *(*int32)(unsafe.Add(mBase, uint32(v89)))
	v97 = v91
	v99 = int32(0)
	v100 = v90
	v101 = v5
	v103 = v86
	v104 = l1
	v105 = v83
	v106 = v85
	v107 = v81 + v72
	goto L17
L16:
	;
	v201 = *(*int32)(unsafe.Add(mBase, uint32(l2+v103<<(uint(int32(3))%32))))
	v202 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v201))))
	if v202 != 0 {
		goto L44
	} else {
		goto L45
	}
L17:
	;
	v109 = int32(*(*int8)(unsafe.Add(mBase, uint32(v104))))
	if v99&int32(1) == int32(0) {
		goto L21
	} else {
		goto L22
	}
L18:
	;
	return int32(0)
L19:
	;
	if v189 != 0 {
		v97 = v192
		v99 = v182
		v100 = v183
		v101 = v184
		v103 = v186
		v104 = v187
		v105 = v188
		v106 = v189
		v107 = v190
		goto L17
	} else {
		goto L43
	}
L20:
	;
	v167 = (v99 | base.B2i32(base.I32_extend8_s(v161) <= v109)) & int32(1)
	if v167 != 0 {
		goto L37
	} else {
		goto L38
	}
L21:
	;
	v114 = int32(*(*int8)(unsafe.Add(mBase, uint32(v97))))
	if v109 < v114 {
		v161 = v114
		goto L20
	} else {
		goto L24
	}
L22:
	;
	goto L23
L23:
	;
	if v101&int32(1) != 0 {
		goto L26
	} else {
		goto L27
	}
L24:
	;
	goto L23
L25:
	;
	v128 = v99 | base.B2i32(base.I32_extend8_s(v124) < v109)
	v129 = int32(1)
	v132 = base.B2i32(v109 < v125) | v101
	if v128&v129&(v132&v129) != 0 {
		goto L16
	} else {
		goto L30
	}
L26:
	;
	v119 = int32(*(*int8)(unsafe.Add(mBase, uint32(v100))))
	v120 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v97))))
	v124 = v120
	v125 = v119
	goto L25
L27:
	;
	goto L28
L28:
	;
	v121 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v97))))
	v122 = int32(*(*int8)(unsafe.Add(mBase, uint32(v100))))
	if v122 < v109 {
		v161 = v121
		goto L20
	} else {
		goto L29
	}
L29:
	;
	v124 = v121
	v125 = v122
	goto L25
L30:
	;
	v136 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v97)+1)))
	if v136 == int32(0) {
		goto L16
	} else {
		goto L31
	}
L31:
	;
	v140 = v100 + int32(1)
	v141 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v140))))
	if v141 == int32(0) {
		goto L16
	} else {
		goto L32
	}
L32:
	;
	v145 = v104 + int32(1)
	v146 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v145))))
	if v146 == int32(0) {
		goto L16
	} else {
		goto L33
	}
L33:
	;
	if base.Ui32((v136-int32(48))&int32(255)) < base.Ui32(int32(10)) {
		goto L34
	} else {
		goto L35
	}
L34:
	;
	v182 = v128
	v183 = v140
	v184 = v132
	v186 = v103
	v187 = v145
	v188 = v105
	v189 = v106
	v190 = v107
	v192 = v97 + int32(1)
	goto L19
L35:
	;
	goto L36
L36:
	;
	v157 = int32(2)
	v182 = v128
	v183 = v100 + v157
	v184 = v132
	v186 = v103
	v187 = v145
	v188 = v105
	v189 = v106
	v190 = v107
	v192 = v97 + v157
	goto L19
L37:
	;
	v168 = v107
	goto L39
L38:
	;
	v168 = v103
	goto L39
L39:
	;
	if v167 != 0 {
		goto L40
	} else {
		goto L41
	}
L40:
	;
	v169 = v103
	goto L42
L41:
	;
	v169 = v105
	goto L42
L42:
	;
	v172 = int32(base.Ui32(v168-v169) >> (uint(int32(1)) % 32))
	v173 = v172 + v169
	v176 = l2 + v173<<(uint(int32(3))%32)
	v177 = *(*int32)(unsafe.Add(mBase, uint32(v176)+4))
	v178 = int32(0)
	v180 = *(*int32)(unsafe.Add(mBase, uint32(v176)))
	v182 = v178
	v183 = v177
	v184 = v178
	v186 = v173
	v187 = l1
	v188 = v169
	v189 = v172
	v190 = v168
	v192 = v180
	goto L19
L43:
	;
	goto L18
L44:
	;
	v204 = l0
	v205 = l1
	v207 = v202
	v208 = v201
	v209 = int32(0)
	goto L47
L45:
	;
	v243 = l0
	v244 = l1
	v258 = int32(1)
	goto L46
L46:
	;
	v259 = int32(45)
	*(*uint8)(unsafe.Add(mBase, uint32(v243))) = uint8(v259)
	v261 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v244))))
	*(*uint8)(unsafe.Add(mBase, uint32(v243)+1)) = uint8(v261)
	return v258
L47:
	;
	v219 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v205))))
	if v219 != 0 {
		goto L49
	} else {
		goto L50
	}
L48:
	;
	v243 = v235
	v244 = v236
	v258 = v239 + int32(1)
	goto L46
L49:
	;
	v220 = int32(45)
	v224 = base.B2i32(v207&int32(255) != v220)
	if v207&int32(255) != v220 {
		goto L52
	} else {
		goto L53
	}
L50:
	;
	v235 = v204
	v236 = v205
	v239 = v209
	goto L51
L51:
	;
	goto L48
L52:
	;
	v225 = v219
	goto L54
L53:
	;
	v225 = v220
	goto L54
L54:
	;
	*(*uint8)(unsafe.Add(mBase, uint32(v204))) = uint8(v225)
	v227 = int32(1)
	v228 = v209 + v227
	v230 = v204 + v227
	v231 = v205 + v224
	v232 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v208)+1)))
	if v232 != 0 {
		v204 = v230
		v205 = v231
		v207 = v232
		v208 = v208 + v227
		v209 = v228
		goto L47
	} else {
		goto L55
	}
L55:
	;
	v235 = v230
	v236 = v231
	v239 = v228
	goto L51
}
func F_hypothetical_percent_rank_final(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v5 int32
	_ = v5
	var v7 int32
	_ = v7
	var v12 int64
	_ = v12
	var v15 int32
	_ = v15
	var v16 int64
	_ = v16
	var v25 float64
	_ = v25
	var v26 int32
	_ = v26
	var v27 int32
	_ = v27
	v5 = m.G0
	v7 = v5 - int32(16)
	m.G0 = v7
	v12 = F_hypothetical_rank_common(m, l0, int32(-1), v7+int32(8))
	mBase = m.M
	v15 = m.ExcPending
	if v15 != 0 {
		return int32(0)
	} else {
		v16 = *(*int64)(unsafe.Add(mBase, uint32(v7)+8))
		if v16 == int64(0) {
			v25 = float64(0)
		} else {
			v25 = base.F64_div(base.F64_convert_i64_s(v12-int64(1)), base.F64_convert_i64_s(v16))
		}
		v26 = F_Float8GetDatum(m, v25)
		mBase = m.M
		v27 = m.ExcPending
		if v27 != 0 {
			return int32(0)
		} else {
			m.G0 = v7 + int32(16)
			return v26
		}
	}
}
func F_hypothetical_rank_final(m *base.Module, l0 int32) int32 {
	var v3 int32
	_ = v3
	var v5 int32
	_ = v5
	var v10 int64
	_ = v10
	var v13 int32
	_ = v13
	var v14 int32
	_ = v14
	var v15 int32
	_ = v15
	v3 = m.G0
	v5 = v3 - int32(16)
	m.G0 = v5
	v10 = F_hypothetical_rank_common(m, l0, int32(-1), v5+int32(8))
	v13 = m.ExcPending
	if v13 != 0 {
		return int32(0)
	} else {
		v14 = F_Int64GetDatum(m, v10)
		v15 = m.ExcPending
		if v15 != 0 {
			return int32(0)
		} else {
			m.G0 = v5 + int32(16)
			return v14
		}
	}
}
