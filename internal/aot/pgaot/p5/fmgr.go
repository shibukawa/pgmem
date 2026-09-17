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
						F_errmsg_internal(m, int32(_a_F_fmgr_c_validator_0), v6)
						mBase = m.M
						v56 = m.ExcPending
						if v56 != 0 {
							return int32(0)
						} else {
							F_errfinish(m, int32(_a_F_fmgr_c_validator_1), int32(809), int32(_a_F_fmgr_c_validator_2))
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
	v4 = *(*int32)(unsafe.Add(mBase, _c_F_fmgr_info[0]))
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
	var v18 int32
	_ = v18
	var v20 int32
	_ = v20
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
	var v31 int32
	_ = v31
	var v32 int32
	_ = v32
	var v42 int32
	_ = v42
	var v43 int32
	_ = v43
	var v46 int32
	_ = v46
	var v54 int32
	_ = v54
	var v55 int32
	_ = v55
	var v57 int32
	_ = v57
	var v58 int32
	_ = v58
	var v60 int32
	_ = v60
	var v69 int32
	_ = v69
	var v70 int32
	_ = v70
	var v71 int32
	_ = v71
	var v72 int32
	_ = v72
	var v74 int32
	_ = v74
	var v75 int32
	_ = v75
	var v83 int32
	_ = v83
	var v84 int32
	_ = v84
	var v96 int32
	_ = v96
	var v97 int32
	_ = v97
	var v98 int32
	_ = v98
	var v108 int32
	_ = v108
	var v119 int32
	_ = v119
	var v120 int32
	_ = v120
	var v121 int32
	_ = v121
	var v122 int32
	_ = v122
	var v123 int32
	_ = v123
	var v126 int32
	_ = v126
	var v127 int32
	_ = v127
	var v139 int32
	_ = v139
	var v140 int32
	_ = v140
	var v141 int32
	_ = v141
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
	var v159 int32
	_ = v159
	var v160 int32
	_ = v160
	var v161 int32
	_ = v161
	var v162 int32
	_ = v162
	var v169 int32
	_ = v169
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
	var v184 int32
	_ = v184
	var v185 int32
	_ = v185
	var v190 int32
	_ = v190
	var v205 int32
	_ = v205
	var v209 int32
	_ = v209
	var v210 int32
	_ = v210
	var v211 int32
	_ = v211
	var v218 int32
	_ = v218
	var v221 int32
	_ = v221
	var v222 int32
	_ = v222
	var v223 int32
	_ = v223
	var v229 int32
	_ = v229
	var v236 int32
	_ = v236
	var v237 int32
	_ = v237
	var v238 int32
	_ = v238
	var v241 int32
	_ = v241
	var v242 int32
	_ = v242
	var v271 int32
	_ = v271
	var v287 int32
	_ = v287
	var v288 int32
	_ = v288
	var v289 int32
	_ = v289
	var v300 int32
	_ = v300
	var v320 int32
	_ = v320
	var v323 int32
	_ = v323
	var v325 int32
	_ = v325
	var v326 int32
	_ = v326
	var v334 int32
	_ = v334
	var v336 int32
	_ = v336
	var v338 int32
	_ = v338
	var v341 int32
	_ = v341
	var v342 int32
	_ = v342
	var v343 int32
	_ = v343
	var v344 int32
	_ = v344
	var v351 int32
	_ = v351
	var v358 int32
	_ = v358
	var v359 int32
	_ = v359
	var v361 int32
	_ = v361
	var v363 int32
	_ = v363
	var v367 int32
	_ = v367
	var v382 int32
	_ = v382
	var v386 int32
	_ = v386
	var v388 int32
	_ = v388
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
	var v405 int32
	_ = v405
	var v409 int32
	_ = v409
	var v413 int32
	_ = v413
	var v416 int32
	_ = v416
	var v424 int32
	_ = v424
	var v429 int32
	_ = v429
	var v432 int32
	_ = v432
	var v434 int32
	_ = v434
	var v435 int32
	_ = v435
	var v437 int32
	_ = v437
	var v449 int32
	_ = v449
	var v450 int32
	_ = v450
	var v454 int32
	_ = v454
	var v455 int32
	_ = v455
	var v456 int32
	_ = v456
	var v464 int32
	_ = v464
	var v473 int32
	_ = v473
	var v477 int32
	_ = v477
	var v479 int32
	_ = v479
	var v480 int32
	_ = v480
	var v485 int32
	_ = v485
	var v486 int32
	_ = v486
	var v487 int32
	_ = v487
	var v488 int32
	_ = v488
	var v489 int32
	_ = v489
	var v490 int32
	_ = v490
	var v491 int32
	_ = v491
	var v506 int32
	_ = v506
	var v517 int32
	_ = v517
	var v518 int32
	_ = v518
	var v519 int32
	_ = v519
	var v526 int32
	_ = v526
	var v527 int32
	_ = v527
	var v528 int32
	_ = v528
	var v529 int32
	_ = v529
	var v532 int32
	_ = v532
	var v535 int32
	_ = v535
	var v538 int32
	_ = v538
	var v546 int32
	_ = v546
	var v553 int32
	_ = v553
	var v555 int32
	_ = v555
	var v557 int32
	_ = v557
	var v560 int32
	_ = v560
	var v561 int64
	_ = v561
	var v563 int64
	_ = v563
	var v564 int64
	_ = v564
	var v565 int64
	_ = v565
	var v569 int64
	_ = v569
	var v570 int64
	_ = v570
	var v573 int64
	_ = v573
	var v575 int64
	_ = v575
	var v580 int64
	_ = v580
	var v597 int32
	_ = v597
	var v598 int32
	_ = v598
	var v607 int32
	_ = v607
	var v608 int32
	_ = v608
	var v609 int32
	_ = v609
	var v616 int32
	_ = v616
	var v617 int32
	_ = v617
	var v623 int32
	_ = v623
	var v631 int32
	_ = v631
	var v636 int32
	_ = v636
	var v647 int32
	_ = v647
	var v655 int32
	_ = v655
	var v660 int32
	_ = v660
	var v668 int32
	_ = v668
	var v686 int32
	_ = v686
	var v687 int64
	_ = v687
	var v691 int32
	_ = v691
	var v693 int32
	_ = v693
	var v694 int32
	_ = v694
	var v697 int32
	_ = v697
	var v699 int32
	_ = v699
	var v701 int32
	_ = v701
	var v702 int32
	_ = v702
	var v703 int32
	_ = v703
	var v704 int32
	_ = v704
	var v705 int32
	_ = v705
	var v706 int32
	_ = v706
	var v707 int32
	_ = v707
	var v711 int32
	_ = v711
	v2 = int32(0)
	v18 = m.G0
	v20 = v18 - int32(256)
	m.G0 = v20
	v25 = v2
	v26 = v2
	v27 = v2
	v28 = v2
	v29 = v2
	v30 = v2
	v31 = v2
	v32 = int32(-1)
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
	if v32 == int32(1) {
		v485 = v25
		v486 = v26
		v487 = v27
		v488 = v28
		v489 = v29
		v490 = v30
		v491 = v31
		goto L7
	} else {
		goto L8
	}
L5:
	;
	goto L3
L6:
	;
	v686 = int32(m.ExcTag)
	v687 = int64(m.ExcVals[0])
	m.ExcPending = 0
	if v686 == int32(0) {
		goto L115
	} else {
		goto L116
	}
L7:
	;
	if v485 == int32(0) {
		goto L83
	} else {
		goto L84
	}
L8:
	;
	v42 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v43 = *(*int32)(unsafe.Add(mBase, uint32(v42)+16))
	if v43 == int32(0) {
		goto L10
	} else {
		goto L11
	}
L9:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v20)+236)) = v26
	*(*int32)(unsafe.Add(mBase, uint32(v20)+232)) = v27
	*(*int32)(unsafe.Add(mBase, uint32(v20)+240)) = v28
	*(*int32)(unsafe.Add(mBase, uint32(v20)+244)) = v29
	*(*int32)(unsafe.Add(mBase, uint32(v20)+248)) = v30
	*(*int32)(unsafe.Add(mBase, uint32(v20)+252)) = v300
	v320 = *(*int32)(unsafe.Add(mBase, _c_F_fmgr_security_definer[0]))
	*(*int32)(unsafe.Add(mBase, uint32(v20+int32(224)))) = v320
	v323 = *(*int32)(unsafe.Add(mBase, _c_F_fmgr_security_definer[1]))
	*(*int32)(unsafe.Add(mBase, uint32(v20+int32(220)))) = v323
	goto L48
L10:
	;
	v46 = *(*int32)(unsafe.Add(mBase, uint32(v42)+20))
	*(*int32)(unsafe.Add(mBase, uint32(v20)+236)) = v26
	*(*int32)(unsafe.Add(mBase, uint32(v20)+232)) = v27
	*(*int32)(unsafe.Add(mBase, uint32(v20)+240)) = v28
	*(*int32)(unsafe.Add(mBase, uint32(v20)+244)) = v29
	*(*int32)(unsafe.Add(mBase, uint32(v20)+248)) = v30
	*(*int32)(unsafe.Add(mBase, uint32(v20)+252)) = v31
	v54 = F_MemoryContextAllocZero(m, v46, int32(48))
	mBase = m.M
	v55 = m.ExcPending
	if v55 != 0 {
		goto L6
	} else {
		goto L13
	}
L11:
	;
	goto L12
L12:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v20)+228)) = v43
	v300 = v31
	goto L9
L13:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v20)+228)) = v54
	v57 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v58 = *(*int32)(unsafe.Add(mBase, uint32(v57)+4))
	v60 = *(*int32)(unsafe.Add(mBase, uint32(v57)+20))
	*(*int32)(unsafe.Add(mBase, uint32(v20)+236)) = v26
	*(*int32)(unsafe.Add(mBase, uint32(v20)+232)) = v27
	*(*int32)(unsafe.Add(mBase, uint32(v20)+240)) = v28
	*(*int32)(unsafe.Add(mBase, uint32(v20)+244)) = v29
	*(*int32)(unsafe.Add(mBase, uint32(v20)+248)) = v30
	*(*int32)(unsafe.Add(mBase, uint32(v20)+252)) = v31
	F_fmgr_info_cxt_security(m, v58, v54, v60, int32(1))
	mBase = m.M
	v69 = m.ExcPending
	if v69 != 0 {
		goto L6
	} else {
		goto L14
	}
L14:
	;
	v70 = *(*int32)(unsafe.Add(mBase, uint32(v20)+228))
	v71 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v72 = *(*int32)(unsafe.Add(mBase, uint32(v71)+24))
	*(*int32)(unsafe.Add(mBase, uint32(v70)+24)) = v72
	v74 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v75 = *(*int32)(unsafe.Add(mBase, uint32(v74)+4))
	*(*int32)(unsafe.Add(mBase, uint32(v20)+236)) = v26
	*(*int32)(unsafe.Add(mBase, uint32(v20)+232)) = v27
	*(*int32)(unsafe.Add(mBase, uint32(v20)+240)) = v28
	*(*int32)(unsafe.Add(mBase, uint32(v20)+244)) = v29
	*(*int32)(unsafe.Add(mBase, uint32(v20)+248)) = v30
	*(*int32)(unsafe.Add(mBase, uint32(v20)+252)) = v31
	v83 = F_SearchSysCache1(m, int32(47), v75)
	mBase = m.M
	v84 = m.ExcPending
	if v84 != 0 {
		goto L6
	} else {
		goto L15
	}
L15:
	;
	if v83 == int32(0) {
		goto L16
	} else {
		goto L17
	}
L16:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v20)+236)) = v26
	*(*int32)(unsafe.Add(mBase, uint32(v20)+232)) = v27
	*(*int32)(unsafe.Add(mBase, uint32(v20)+240)) = v28
	*(*int32)(unsafe.Add(mBase, uint32(v20)+244)) = v29
	*(*int32)(unsafe.Add(mBase, uint32(v20)+248)) = v30
	*(*int32)(unsafe.Add(mBase, uint32(v20)+252)) = v31
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v96 = m.ExcPending
	if v96 != 0 {
		goto L6
	} else {
		goto L19
	}
L17:
	;
	goto L18
L18:
	;
	v120 = *(*int32)(unsafe.Add(mBase, uint32(v83)+16))
	v121 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v120)+22)))
	v122 = v120 + v121
	v123 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v122)+97)))
	if v123 == int32(1) {
		goto L22
	} else {
		goto L23
	}
L19:
	;
	v97 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v98 = *(*int32)(unsafe.Add(mBase, uint32(v97)+4))
	*(*int32)(unsafe.Add(mBase, uint32(v20)+236)) = v26
	*(*int32)(unsafe.Add(mBase, uint32(v20)+232)) = v27
	*(*int32)(unsafe.Add(mBase, uint32(v20)+240)) = v28
	*(*int32)(unsafe.Add(mBase, uint32(v20)+244)) = v29
	*(*int32)(unsafe.Add(mBase, uint32(v20)+248)) = v30
	*(*int32)(unsafe.Add(mBase, uint32(v20)+252)) = v31
	*(*int32)(unsafe.Add(mBase, uint32(v20))) = v98
	F_errmsg_internal(m, int32(_a_F_fmgr_security_definer_0), v20)
	mBase = m.M
	v108 = m.ExcPending
	if v108 != 0 {
		goto L6
	} else {
		goto L20
	}
L20:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v20)+236)) = v26
	*(*int32)(unsafe.Add(mBase, uint32(v20)+232)) = v27
	*(*int32)(unsafe.Add(mBase, uint32(v20)+240)) = v28
	*(*int32)(unsafe.Add(mBase, uint32(v20)+244)) = v29
	*(*int32)(unsafe.Add(mBase, uint32(v20)+248)) = v30
	*(*int32)(unsafe.Add(mBase, uint32(v20)+252)) = v31
	F_errfinish(m, int32(_a_F_fmgr_security_definer_1), int32(664), int32(_a_F_fmgr_security_definer_2))
	mBase = m.M
	v119 = m.ExcPending
	if v119 != 0 {
		goto L6
	} else {
		goto L21
	}
L21:
	;
	goto L3
L22:
	;
	v126 = *(*int32)(unsafe.Add(mBase, uint32(v20)+228))
	v127 = *(*int32)(unsafe.Add(mBase, uint32(v122)+72))
	*(*int32)(unsafe.Add(mBase, uint32(v126)+28)) = v127
	goto L24
L23:
	;
	goto L24
L24:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v20)+236)) = v26
	*(*int32)(unsafe.Add(mBase, uint32(v20)+232)) = v27
	*(*int32)(unsafe.Add(mBase, uint32(v20)+240)) = v28
	*(*int32)(unsafe.Add(mBase, uint32(v20)+244)) = v29
	*(*int32)(unsafe.Add(mBase, uint32(v20)+248)) = v30
	*(*int32)(unsafe.Add(mBase, uint32(v20)+252)) = v31
	v139 = F_SysCacheGetAttr(m, int32(47), v83, int32(29), v20+int32(183))
	mBase = m.M
	v140 = m.ExcPending
	if v140 != 0 {
		goto L6
	} else {
		goto L25
	}
L25:
	;
	v141 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v20)+183)))
	if v141 == int32(0) {
		goto L26
	} else {
		goto L27
	}
L26:
	;
	v144 = int32(_a_F_fmgr_security_definer_3)
	v145 = *(*int32)(unsafe.Add(mBase, _c_F_fmgr_security_definer[2]))
	v147 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v148 = *(*int32)(unsafe.Add(mBase, uint32(v147)+20))
	*(*int32)(unsafe.Add(mBase, _c_F_fmgr_security_definer[2])) = v148
	v150 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v139))))
	if v150&int32(3) != 0 {
		goto L29
	} else {
		goto L30
	}
L27:
	;
	v271 = v31
	goto L28
L28:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v20)+236)) = v26
	*(*int32)(unsafe.Add(mBase, uint32(v20)+232)) = v27
	*(*int32)(unsafe.Add(mBase, uint32(v20)+240)) = v28
	*(*int32)(unsafe.Add(mBase, uint32(v20)+244)) = v29
	*(*int32)(unsafe.Add(mBase, uint32(v20)+248)) = v30
	*(*int32)(unsafe.Add(mBase, uint32(v20)+252)) = v271
	F_ReleaseCatCache(m, v83)
	mBase = m.M
	v287 = m.ExcPending
	if v287 != 0 {
		goto L6
	} else {
		goto L47
	}
L29:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v20)+236)) = v26
	*(*int32)(unsafe.Add(mBase, uint32(v20)+232)) = v27
	*(*int32)(unsafe.Add(mBase, uint32(v20)+240)) = v28
	*(*int32)(unsafe.Add(mBase, uint32(v20)+244)) = v29
	*(*int32)(unsafe.Add(mBase, uint32(v20)+248)) = v30
	*(*int32)(unsafe.Add(mBase, uint32(v20)+252)) = v31
	v159 = F_detoast_attr(m, v139)
	mBase = m.M
	v160 = m.ExcPending
	if v160 != 0 {
		goto L6
	} else {
		goto L32
	}
L30:
	;
	v161 = v139
	v162 = v31
	goto L31
L31:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v20)+232)) = v27
	*(*int32)(unsafe.Add(mBase, uint32(v20)+236)) = v26
	*(*int32)(unsafe.Add(mBase, uint32(v20)+240)) = v28
	*(*int32)(unsafe.Add(mBase, uint32(v20)+244)) = v29
	*(*int32)(unsafe.Add(mBase, uint32(v20)+248)) = v30
	*(*int32)(unsafe.Add(mBase, uint32(v20)+252)) = v162
	v169 = *(*int32)(unsafe.Add(mBase, uint32(v20)+228))
	F_TransformGUCArray(m, v161, v169+int32(32), v169+int32(40))
	mBase = m.M
	v176 = m.ExcPending
	if v176 != 0 {
		goto L6
	} else {
		goto L33
	}
L32:
	;
	v161 = v159
	v162 = v159
	goto L31
L33:
	;
	v177 = *(*int32)(unsafe.Add(mBase, uint32(v20)+228))
	v178 = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v177)+36)) = v178
	v180 = *(*int32)(unsafe.Add(mBase, uint32(v20)+228))
	v181 = *(*int32)(unsafe.Add(mBase, uint32(v180)+32))
	if v181 == v178 {
		goto L34
	} else {
		goto L35
	}
L34:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_fmgr_security_definer[2])) = v145
	v271 = v162
	goto L28
L35:
	;
	v184 = int32(0)
	v185 = *(*int32)(unsafe.Add(mBase, uint32(v181)+4))
	if v185 <= v184 {
		goto L34
	} else {
		goto L36
	}
L36:
	;
	v190 = v184
	goto L37
L37:
	;
	v205 = *(*int32)(unsafe.Add(mBase, uint32(v181)+12))
	v209 = *(*int32)(unsafe.Add(mBase, uint32(v205+v190<<(uint(int32(2))%32))))
	v210 = *(*int32)(unsafe.Add(mBase, uint32(v20)+228))
	v211 = *(*int32)(unsafe.Add(mBase, uint32(v210)+36))
	*(*int32)(unsafe.Add(mBase, uint32(v20)+236)) = v26
	*(*int32)(unsafe.Add(mBase, uint32(v20)+232)) = v27
	*(*int32)(unsafe.Add(mBase, uint32(v20)+240)) = v28
	*(*int32)(unsafe.Add(mBase, uint32(v20)+244)) = v29
	*(*int32)(unsafe.Add(mBase, uint32(v20)+248)) = v30
	*(*int32)(unsafe.Add(mBase, uint32(v20)+252)) = v162
	v218 = int32(0)
	v221 = F_find_option(m, v209, v218, v218, v218)
	mBase = m.M
	v222 = m.ExcPending
	if v222 != 0 {
		goto L6
	} else {
		goto L40
	}
L38:
	;
	goto L34
L39:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v20)+236)) = v26
	*(*int32)(unsafe.Add(mBase, uint32(v20)+232)) = v27
	*(*int32)(unsafe.Add(mBase, uint32(v20)+240)) = v28
	*(*int32)(unsafe.Add(mBase, uint32(v20)+244)) = v29
	*(*int32)(unsafe.Add(mBase, uint32(v20)+248)) = v30
	*(*int32)(unsafe.Add(mBase, uint32(v20)+252)) = v162
	v236 = F_lappend(m, v211, v229)
	mBase = m.M
	v237 = m.ExcPending
	if v237 != 0 {
		goto L6
	} else {
		goto L45
	}
L40:
	;
	if v221 != 0 {
		goto L41
	} else {
		goto L42
	}
L41:
	;
	v223 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v221)+21)))
	if v223&int32(2) == int32(0) {
		v229 = v221
		goto L39
	} else {
		goto L44
	}
L42:
	;
	goto L43
L43:
	;
	v229 = int32(0)
	goto L39
L44:
	;
	goto L43
L45:
	;
	v238 = *(*int32)(unsafe.Add(mBase, uint32(v20)+228))
	*(*int32)(unsafe.Add(mBase, uint32(v238)+36)) = v236
	v241 = v190 + int32(1)
	v242 = *(*int32)(unsafe.Add(mBase, uint32(v181)+4))
	if v241 < v242 {
		v190 = v241
		goto L37
	} else {
		goto L46
	}
L46:
	;
	goto L38
L47:
	;
	v288 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v289 = *(*int32)(unsafe.Add(mBase, uint32(v20)+228))
	*(*int32)(unsafe.Add(mBase, uint32(v288)+16)) = v289
	v300 = v271
	goto L9
L48:
	;
	v325 = *(*int32)(unsafe.Add(mBase, uint32(v20)+228))
	v326 = *(*int32)(unsafe.Add(mBase, uint32(v325)+32))
	if v326 != 0 {
		goto L49
	} else {
		goto L50
	}
L49:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v20)+236)) = v26
	*(*int32)(unsafe.Add(mBase, uint32(v20)+232)) = v27
	*(*int32)(unsafe.Add(mBase, uint32(v20)+240)) = v28
	*(*int32)(unsafe.Add(mBase, uint32(v20)+244)) = v29
	*(*int32)(unsafe.Add(mBase, uint32(v20)+248)) = v30
	*(*int32)(unsafe.Add(mBase, uint32(v20)+252)) = v300
	v334 = int32(_a_F_fmgr_security_definer_4)
	v336 = *(*int32)(unsafe.Add(mBase, _c_F_fmgr_security_definer[3]))
	v338 = v336 + int32(1)
	*(*int32)(unsafe.Add(mBase, _c_F_fmgr_security_definer[3])) = v338
	goto L52
L50:
	;
	v341 = v30
	v342 = int32(0)
	goto L51
L51:
	;
	v343 = *(*int32)(unsafe.Add(mBase, uint32(v20)+228))
	v344 = *(*int32)(unsafe.Add(mBase, uint32(v343)+28))
	if v344 != 0 {
		goto L53
	} else {
		goto L54
	}
L52:
	;
	v341 = v338
	v342 = v338
	goto L51
L53:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v20)+232)) = v27
	*(*int32)(unsafe.Add(mBase, uint32(v20)+236)) = v26
	*(*int32)(unsafe.Add(mBase, uint32(v20)+240)) = v28
	*(*int32)(unsafe.Add(mBase, uint32(v20)+244)) = v342
	*(*int32)(unsafe.Add(mBase, uint32(v20)+248)) = v341
	*(*int32)(unsafe.Add(mBase, uint32(v20)+252)) = v300
	v351 = *(*int32)(unsafe.Add(mBase, uint32(v20)+220))
	*(*int32)(unsafe.Add(mBase, _c_F_fmgr_security_definer[1])) = v351 | int32(1)
	*(*int32)(unsafe.Add(mBase, _c_F_fmgr_security_definer[0])) = v344
	goto L56
L54:
	;
	goto L55
L55:
	;
	v358 = *(*int32)(unsafe.Add(mBase, uint32(v20)+228))
	v359 = *(*int32)(unsafe.Add(mBase, uint32(v358)+32))
	v361 = *(*int32)(unsafe.Add(mBase, uint32(v358)+36))
	v363 = *(*int32)(unsafe.Add(mBase, uint32(v358)+40))
	v367 = int32(0)
	goto L57
L56:
	;
	goto L55
L57:
	;
	v382 = int32(0)
	if v359 == v382 {
		v392 = v382
		goto L59
	} else {
		goto L60
	}
L59:
	;
	v393 = int32(0)
	if v361 == v393 {
		v402 = v393
		goto L62
	} else {
		goto L63
	}
L60:
	;
	v386 = *(*int32)(unsafe.Add(mBase, uint32(v359)+4))
	if v386 <= v367 {
		v392 = int32(0)
		goto L59
	} else {
		goto L61
	}
L61:
	;
	v388 = *(*int32)(unsafe.Add(mBase, uint32(v359)+12))
	v392 = v388 + v367<<(uint(int32(2))%32)
	goto L59
L62:
	;
	if v363 == int32(0) {
		goto L66
	} else {
		goto L67
	}
L63:
	;
	v396 = *(*int32)(unsafe.Add(mBase, uint32(v361)+4))
	if v396 <= v367 {
		v402 = v393
		goto L62
	} else {
		goto L64
	}
L64:
	;
	v398 = *(*int32)(unsafe.Add(mBase, uint32(v361)+12))
	v402 = v398 + v367<<(uint(int32(2))%32)
	goto L62
L65:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v20)+236)) = v26
	*(*int32)(unsafe.Add(mBase, uint32(v20)+232)) = v27
	*(*int32)(unsafe.Add(mBase, uint32(v20)+240)) = v28
	*(*int32)(unsafe.Add(mBase, uint32(v20)+244)) = v342
	*(*int32)(unsafe.Add(mBase, uint32(v20)+248)) = v341
	*(*int32)(unsafe.Add(mBase, uint32(v20)+252)) = v300
	v449 = F_superuser(m)
	mBase = m.M
	v450 = m.ExcPending
	if v450 != 0 {
		goto L6
	} else {
		goto L78
	}
L66:
	;
	v416 = *(*int32)(unsafe.Add(mBase, _c_F_fmgr_security_definer[4]))
	if v416 != 0 {
		goto L70
	} else {
		goto L71
	}
L67:
	;
	v405 = int32(0)
	v409 = *(*int32)(unsafe.Add(mBase, uint32(v363)+4))
	if base.B2i32(v402 == v405)|(base.B2i32(v392 == v405)|base.B2i32(v409 <= v367)) != 0 {
		goto L66
	} else {
		goto L68
	}
L68:
	;
	v413 = *(*int32)(unsafe.Add(mBase, uint32(v363)+12))
	if v413 != 0 {
		goto L65
	} else {
		goto L69
	}
L69:
	;
	goto L66
L70:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v20)+232)) = v27
	*(*int32)(unsafe.Add(mBase, uint32(v20)+236)) = v26
	*(*int32)(unsafe.Add(mBase, uint32(v20)+240)) = v28
	*(*int32)(unsafe.Add(mBase, uint32(v20)+244)) = v342
	*(*int32)(unsafe.Add(mBase, uint32(v20)+248)) = v341
	*(*int32)(unsafe.Add(mBase, uint32(v20)+252)) = v300
	v424 = *(*int32)(unsafe.Add(mBase, uint32(v20)+228))
	m.T0[v416].(func(*base.Module, int32, int32, int32))(m, int32(0), v424, v424+int32(44))
	mBase = m.M
	v429 = m.ExcPending
	if v429 != 0 {
		goto L6
	} else {
		goto L73
	}
L71:
	;
	goto L72
L72:
	;
	v432 = *(*int32)(unsafe.Add(mBase, _c_F_fmgr_security_definer[5]))
	v434 = *(*int32)(unsafe.Add(mBase, _c_F_fmgr_security_definer[6]))
	v435 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	goto L74
L73:
	;
	goto L72
L74:
	;
	v437 = v20 + int32(16)
	*(*int32)(unsafe.Add(mBase, uint32(v437)+4)) = int32(1)
	*(*int32)(unsafe.Add(mBase, uint32(v437))) = v20 + int32(12)
	goto L77
L75:
	;
	v485 = int32(0)
	v486 = v434
	v487 = v432
	v488 = v435
	v489 = v342
	v490 = v341
	v491 = v300
	goto L7
L77:
	;
	goto L75
L78:
	;
	v454 = *(*int32)(unsafe.Add(mBase, uint32(v413+v367<<(uint(int32(2))%32))))
	v455 = *(*int32)(unsafe.Add(mBase, uint32(v402)))
	v456 = *(*int32)(unsafe.Add(mBase, uint32(v392)))
	*(*int32)(unsafe.Add(mBase, uint32(v20)+236)) = v26
	*(*int32)(unsafe.Add(mBase, uint32(v20)+232)) = v27
	*(*int32)(unsafe.Add(mBase, uint32(v20)+240)) = v28
	*(*int32)(unsafe.Add(mBase, uint32(v20)+244)) = v342
	*(*int32)(unsafe.Add(mBase, uint32(v20)+248)) = v341
	*(*int32)(unsafe.Add(mBase, uint32(v20)+252)) = v300
	v464 = *(*int32)(unsafe.Add(mBase, _c_F_fmgr_security_definer[0]))
	*(*int32)(unsafe.Add(mBase, uint32(v20)+236)) = v26
	*(*int32)(unsafe.Add(mBase, uint32(v20)+232)) = v27
	*(*int32)(unsafe.Add(mBase, uint32(v20)+240)) = v28
	*(*int32)(unsafe.Add(mBase, uint32(v20)+244)) = v342
	*(*int32)(unsafe.Add(mBase, uint32(v20)+248)) = v341
	*(*int32)(unsafe.Add(mBase, uint32(v20)+252)) = v300
	if v449 != 0 {
		goto L79
	} else {
		goto L80
	}
L79:
	;
	v473 = int32(5)
	goto L81
L80:
	;
	v473 = int32(6)
	goto L81
L81:
	;
	v477 = int32(0)
	v479 = F_set_config_with_handle(m, v456, v455, v454, v473, int32(13), v464, int32(2), int32(1), v477, v477)
	mBase = m.M
	v480 = m.ExcPending
	if v480 != 0 {
		goto L6
	} else {
		goto L82
	}
L82:
	;
	v367 = v367 + int32(1)
	goto L57
L83:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_fmgr_security_definer[6])) = v20 + int32(16)
	v506 = *(*int32)(unsafe.Add(mBase, uint32(v20)+228))
	*(*int32)(unsafe.Add(mBase, uint32(l0))) = v506
	*(*int32)(unsafe.Add(mBase, uint32(v20)+236)) = v486
	*(*int32)(unsafe.Add(mBase, uint32(v20)+232)) = v487
	*(*int32)(unsafe.Add(mBase, uint32(v20)+240)) = v488
	*(*int32)(unsafe.Add(mBase, uint32(v20)+244)) = v489
	*(*int32)(unsafe.Add(mBase, uint32(v20)+248)) = v490
	*(*int32)(unsafe.Add(mBase, uint32(v20)+252)) = v491
	F_pgstat_init_function_usage(m, l0, v20+int32(184))
	mBase = m.M
	v517 = m.ExcPending
	if v517 != 0 {
		goto L6
	} else {
		goto L86
	}
L84:
	;
	goto L85
L85:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_fmgr_security_definer[5])) = v487
	*(*int32)(unsafe.Add(mBase, _c_F_fmgr_security_definer[6])) = v486
	*(*int32)(unsafe.Add(mBase, uint32(l0))) = v488
	v647 = *(*int32)(unsafe.Add(mBase, _c_F_fmgr_security_definer[4]))
	if v647 != 0 {
		goto L110
	} else {
		goto L111
	}
L86:
	;
	v518 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v519 = *(*int32)(unsafe.Add(mBase, uint32(v518)))
	*(*int32)(unsafe.Add(mBase, uint32(v20)+236)) = v486
	*(*int32)(unsafe.Add(mBase, uint32(v20)+232)) = v487
	*(*int32)(unsafe.Add(mBase, uint32(v20)+240)) = v488
	*(*int32)(unsafe.Add(mBase, uint32(v20)+244)) = v489
	*(*int32)(unsafe.Add(mBase, uint32(v20)+248)) = v490
	*(*int32)(unsafe.Add(mBase, uint32(v20)+252)) = v491
	v526 = m.T0[v519].(func(*base.Module, int32) int32)(m, l0)
	mBase = m.M
	v527 = m.ExcPending
	if v527 != 0 {
		goto L6
	} else {
		goto L87
	}
L87:
	;
	v528 = int32(1)
	v529 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	if v529 == int32(0) {
		v538 = v528
		goto L88
	} else {
		goto L89
	}
L88:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v20)+236)) = v486
	*(*int32)(unsafe.Add(mBase, uint32(v20)+232)) = v487
	*(*int32)(unsafe.Add(mBase, uint32(v20)+240)) = v488
	*(*int32)(unsafe.Add(mBase, uint32(v20)+244)) = v489
	*(*int32)(unsafe.Add(mBase, uint32(v20)+248)) = v490
	*(*int32)(unsafe.Add(mBase, uint32(v20)+252)) = v491
	v546 = v20 + int32(184)
	v553 = m.G0
	v555 = v553 - int32(16)
	m.G0 = v555
	v557 = *(*int32)(unsafe.Add(mBase, uint32(v546)))
	if v557 != 0 {
		goto L92
	} else {
		goto L93
	}
L89:
	;
	v532 = *(*int32)(unsafe.Add(mBase, uint32(v529)))
	if v532 != int32(383) {
		v538 = v528
		goto L88
	} else {
		goto L90
	}
L90:
	;
	v535 = *(*int32)(unsafe.Add(mBase, uint32(v529)+20))
	v538 = base.B2i32(v535 != int32(1))
	goto L88
L91:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_fmgr_security_definer[5])) = v487
	*(*int32)(unsafe.Add(mBase, _c_F_fmgr_security_definer[6])) = v486
	*(*int32)(unsafe.Add(mBase, uint32(l0))) = v488
	v597 = *(*int32)(unsafe.Add(mBase, uint32(v20)+228))
	v598 = *(*int32)(unsafe.Add(mBase, uint32(v597)+32))
	if v598 != 0 {
		goto L98
	} else {
		goto L99
	}
L92:
	;
	F___clock_gettime(m, int32(1), v555)
	mBase = m.M
	v560 = int32(_a_F_fmgr_security_definer_5)
	v561 = *(*int64)(unsafe.Add(mBase, _c_F_fmgr_security_definer[7]))
	v563 = *(*int64)(unsafe.Add(mBase, uint32(v546)+16))
	v564 = int64(*(*int32)(unsafe.Add(mBase, uint32(v555)+8)))
	v565 = *(*int64)(unsafe.Add(mBase, uint32(v555)))
	v569 = *(*int64)(unsafe.Add(mBase, uint32(v546)+24))
	v570 = v564 + v565*int64(1000000000) - v569
	*(*int64)(unsafe.Add(mBase, _c_F_fmgr_security_definer[7])) = v563 + v570
	v573 = *(*int64)(unsafe.Add(mBase, uint32(v546)+8))
	if v538 != 0 {
		goto L95
	} else {
		goto L96
	}
L93:
	;
	goto L94
L94:
	;
	m.G0 = v555 + int32(16)
	goto L91
L95:
	;
	v575 = *(*int64)(unsafe.Add(mBase, uint32(v557)))
	*(*int64)(unsafe.Add(mBase, uint32(v557))) = v575 + int64(1)
	goto L97
L96:
	;
	goto L97
L97:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v557)+8)) = v573 + v570
	v580 = *(*int64)(unsafe.Add(mBase, uint32(v557)+16))
	*(*int64)(unsafe.Add(mBase, uint32(v557)+16)) = v580 + (v570 - v561 + v563)
	goto L94
L98:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v20)+236)) = v486
	*(*int32)(unsafe.Add(mBase, uint32(v20)+232)) = v487
	*(*int32)(unsafe.Add(mBase, uint32(v20)+240)) = v488
	*(*int32)(unsafe.Add(mBase, uint32(v20)+244)) = v489
	*(*int32)(unsafe.Add(mBase, uint32(v20)+248)) = v490
	*(*int32)(unsafe.Add(mBase, uint32(v20)+252)) = v491
	F_AtEOXact_GUC(m, int32(1), v489)
	mBase = m.M
	v607 = m.ExcPending
	if v607 != 0 {
		goto L6
	} else {
		goto L101
	}
L99:
	;
	goto L100
L100:
	;
	v608 = *(*int32)(unsafe.Add(mBase, uint32(v20)+228))
	v609 = *(*int32)(unsafe.Add(mBase, uint32(v608)+28))
	if v609 != 0 {
		goto L102
	} else {
		goto L103
	}
L101:
	;
	goto L100
L102:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v20)+232)) = v487
	*(*int32)(unsafe.Add(mBase, uint32(v20)+236)) = v486
	*(*int32)(unsafe.Add(mBase, uint32(v20)+240)) = v488
	*(*int32)(unsafe.Add(mBase, uint32(v20)+244)) = v489
	*(*int32)(unsafe.Add(mBase, uint32(v20)+248)) = v490
	*(*int32)(unsafe.Add(mBase, uint32(v20)+252)) = v491
	v616 = *(*int32)(unsafe.Add(mBase, uint32(v20)+224))
	v617 = *(*int32)(unsafe.Add(mBase, uint32(v20)+220))
	*(*int32)(unsafe.Add(mBase, _c_F_fmgr_security_definer[1])) = v617
	*(*int32)(unsafe.Add(mBase, _c_F_fmgr_security_definer[0])) = v616
	goto L105
L103:
	;
	goto L104
L104:
	;
	v623 = *(*int32)(unsafe.Add(mBase, _c_F_fmgr_security_definer[4]))
	if v623 != 0 {
		goto L106
	} else {
		goto L107
	}
L105:
	;
	goto L104
L106:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v20)+232)) = v487
	*(*int32)(unsafe.Add(mBase, uint32(v20)+236)) = v486
	*(*int32)(unsafe.Add(mBase, uint32(v20)+240)) = v488
	*(*int32)(unsafe.Add(mBase, uint32(v20)+244)) = v489
	*(*int32)(unsafe.Add(mBase, uint32(v20)+248)) = v490
	*(*int32)(unsafe.Add(mBase, uint32(v20)+252)) = v491
	v631 = *(*int32)(unsafe.Add(mBase, uint32(v20)+228))
	m.T0[v623].(func(*base.Module, int32, int32, int32))(m, int32(1), v631, v631+int32(44))
	mBase = m.M
	v636 = m.ExcPending
	if v636 != 0 {
		goto L6
	} else {
		goto L109
	}
L107:
	;
	goto L108
L108:
	;
	m.G0 = v20 + int32(256)
	return v526
L109:
	;
	goto L108
L110:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v20)+232)) = v487
	*(*int32)(unsafe.Add(mBase, uint32(v20)+236)) = v486
	*(*int32)(unsafe.Add(mBase, uint32(v20)+240)) = v488
	*(*int32)(unsafe.Add(mBase, uint32(v20)+244)) = v489
	*(*int32)(unsafe.Add(mBase, uint32(v20)+248)) = v490
	*(*int32)(unsafe.Add(mBase, uint32(v20)+252)) = v491
	v655 = *(*int32)(unsafe.Add(mBase, uint32(v20)+228))
	m.T0[v647].(func(*base.Module, int32, int32, int32))(m, int32(2), v655, v655+int32(44))
	mBase = m.M
	v660 = m.ExcPending
	if v660 != 0 {
		goto L6
	} else {
		goto L113
	}
L111:
	;
	goto L112
L112:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v20)+236)) = v486
	*(*int32)(unsafe.Add(mBase, uint32(v20)+232)) = v487
	*(*int32)(unsafe.Add(mBase, uint32(v20)+240)) = v488
	*(*int32)(unsafe.Add(mBase, uint32(v20)+244)) = v489
	*(*int32)(unsafe.Add(mBase, uint32(v20)+248)) = v490
	*(*int32)(unsafe.Add(mBase, uint32(v20)+252)) = v491
	F_pg_re_throw(m)
	mBase = m.M
	v668 = m.ExcPending
	if v668 != 0 {
		goto L6
	} else {
		goto L114
	}
L113:
	;
	goto L112
L114:
	;
	goto L5
L115:
	;
	v691 = int32(v687)
	m.G0 = v20
	v693 = *(*int32)(unsafe.Add(mBase, uint32(v691)+4))
	v694 = *(*int32)(unsafe.Add(mBase, uint32(v691)))
	v697 = *(*int32)(unsafe.Add(mBase, uint32(v694)))
	if v20+int32(12) == v697 {
		goto L118
	} else {
		goto L119
	}
L116:
	;
	m.ExcPending = 1
	goto L124
L117:
	;
	if v701 != 0 {
		goto L121
	} else {
		goto L122
	}
L118:
	;
	v699 = *(*int32)(unsafe.Add(mBase, uint32(v694)+4))
	v701 = v699
	goto L120
L119:
	;
	v701 = int32(0)
	goto L120
L120:
	;
	goto L117
L121:
	;
	v702 = *(*int32)(unsafe.Add(mBase, uint32(v20)+252))
	v703 = *(*int32)(unsafe.Add(mBase, uint32(v20)+248))
	v704 = *(*int32)(unsafe.Add(mBase, uint32(v20)+244))
	v705 = *(*int32)(unsafe.Add(mBase, uint32(v20)+240))
	v706 = *(*int32)(unsafe.Add(mBase, uint32(v20)+236))
	v707 = *(*int32)(unsafe.Add(mBase, uint32(v20)+232))
	v25 = v693
	v26 = v706
	v27 = v707
	v28 = v705
	v29 = v704
	v30 = v703
	v31 = v702
	v32 = v701
	goto L1
L122:
	;
	goto L123
L123:
	;
	F___wasm_longjmp(m, v694, v693)
	mBase = m.M
	v711 = m.ExcPending
	if v711 != 0 {
		goto L124
	} else {
		goto L125
	}
L124:
	;
	return int32(0)
L125:
	;
	base.Wasm_trap_unreachable()
	for {
	}
}
