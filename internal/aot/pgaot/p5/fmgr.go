package p5

import (
	base "github.com/shibukawa/pgmem/internal/aot/pgaot/base"
	"unsafe"
)

func F_fmgr_c_validator(m *base.Module, l0 int32) int64 {
	mBase := m.M
	_ = mBase
	var v5 int32
	_ = v5
	var v7 int32
	_ = v7
	var v9 int32
	_ = v9
	var v10 int32
	_ = v10
	var v11 int64
	_ = v11
	var v12 int32
	_ = v12
	var v13 int32
	_ = v13
	var v16 int32
	_ = v16
	var v20 int32
	_ = v20
	var v21 int32
	_ = v21
	var v26 int64
	_ = v26
	var v27 int32
	_ = v27
	var v29 int32
	_ = v29
	var v30 int32
	_ = v30
	var v33 int64
	_ = v33
	var v34 int32
	_ = v34
	var v36 int32
	_ = v36
	var v37 int32
	_ = v37
	var v41 int32
	_ = v41
	var v42 int32
	_ = v42
	var v43 int32
	_ = v43
	var v44 int32
	_ = v44
	var v45 int32
	_ = v45
	var v47 int32
	_ = v47
	var v58 int32
	_ = v58
	var v62 int32
	_ = v62
	var v67 int32
	_ = v67
	v5 = m.G0
	v7 = v5 - int32(16)
	m.G0 = v7
	v9 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v10 = *(*int32)(unsafe.Add(mBase, uint32(v9)+4))
	v11 = *(*int64)(unsafe.Add(mBase, uint32(l0)+24))
	v12 = base.I32_wrap_i64(v11)
	v13 = F_CheckFunctionValidatorAccess(m, v10, v12)
	mBase = m.M
	v16 = m.ExcPending
	if v16 != 0 {
		return int64(0)
	} else {
		if v13 != 0 {
			v20 = F_SearchSysCache1(m, int32(47), v11&int64(4294967295))
			mBase = m.M
			v21 = m.ExcPending
			if v21 != 0 {
				return int64(0)
			} else {
				if v20 == int32(0) {
					F_errstart_cold(m, int32(21), int32(0))
					mBase = m.M
					v58 = m.ExcPending
					if v58 != 0 {
						return int64(0)
					} else {
						*(*int32)(unsafe.Add(mBase, uint32(v7))) = v12
						F_errmsg_internal(m, int32(_a_F_fmgr_c_validator_0), v7)
						mBase = m.M
						v62 = m.ExcPending
						if v62 != 0 {
							return int64(0)
						} else {
							F_errfinish(m, int32(_a_F_fmgr_c_validator_1), int32(840), int32(_a_F_fmgr_c_validator_2))
							mBase = m.M
							v67 = m.ExcPending
							if v67 != 0 {
								return int64(0)
							} else {
								base.Wasm_trap_unreachable()
								for {
								}
							}
						}
					}
				} else {
					v26 = F_SysCacheGetAttrNotNull(m, int32(47), v20, int32(26))
					mBase = m.M
					v27 = m.ExcPending
					if v27 != 0 {
						return int64(0)
					} else {
						v29 = F_text_to_cstring(m, base.I32_wrap_i64(v26))
						mBase = m.M
						v30 = m.ExcPending
						if v30 != 0 {
							return int64(0)
						} else {
							v33 = F_SysCacheGetAttrNotNull(m, int32(47), v20, int32(27))
							mBase = m.M
							v34 = m.ExcPending
							if v34 != 0 {
								return int64(0)
							} else {
								v36 = F_text_to_cstring(m, base.I32_wrap_i64(v33))
								mBase = m.M
								v37 = m.ExcPending
								if v37 != 0 {
									return int64(0)
								} else {
									v41 = F_load_external_function(m, v36, v29, int32(1), v7+int32(12))
									mBase = m.M
									v42 = m.ExcPending
									if v42 != 0 {
										return int64(0)
									} else {
										v43 = *(*int32)(unsafe.Add(mBase, uint32(v7)+12))
										v44 = F_fetch_finfo_record(m, v43, v29)
										mBase = m.M
										v45 = m.ExcPending
										if v45 != 0 {
											return int64(0)
										} else {
											F_ReleaseCatCache(m, v20)
											mBase = m.M
											v47 = m.ExcPending
											if v47 != 0 {
												return int64(0)
											} else {
												m.G0 = v7 + int32(16)
												return int64(0)
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
			m.G0 = v7 + int32(16)
			return int64(0)
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
func F_fmgr_security_definer(m *base.Module, l0 int32) int64 {
	mBase := m.M
	_ = mBase
	var v2 int32
	_ = v2
	var v19 int32
	_ = v19
	var v21 int32
	_ = v21
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
	var v33 int32
	_ = v33
	var v44 int32
	_ = v44
	var v45 int32
	_ = v45
	var v48 int32
	_ = v48
	var v56 int32
	_ = v56
	var v57 int32
	_ = v57
	var v59 int32
	_ = v59
	var v60 int32
	_ = v60
	var v62 int32
	_ = v62
	var v71 int32
	_ = v71
	var v72 int32
	_ = v72
	var v73 int32
	_ = v73
	var v74 int32
	_ = v74
	var v76 int32
	_ = v76
	var v77 int64
	_ = v77
	var v85 int32
	_ = v85
	var v86 int32
	_ = v86
	var v98 int32
	_ = v98
	var v99 int32
	_ = v99
	var v100 int32
	_ = v100
	var v110 int32
	_ = v110
	var v121 int32
	_ = v121
	var v122 int32
	_ = v122
	var v123 int32
	_ = v123
	var v124 int32
	_ = v124
	var v125 int32
	_ = v125
	var v128 int32
	_ = v128
	var v129 int32
	_ = v129
	var v141 int64
	_ = v141
	var v142 int32
	_ = v142
	var v143 int32
	_ = v143
	var v146 int32
	_ = v146
	var v147 int32
	_ = v147
	var v149 int32
	_ = v149
	var v150 int32
	_ = v150
	var v152 int32
	_ = v152
	var v153 int32
	_ = v153
	var v162 int32
	_ = v162
	var v163 int32
	_ = v163
	var v164 int32
	_ = v164
	var v165 int32
	_ = v165
	var v172 int32
	_ = v172
	var v179 int32
	_ = v179
	var v180 int32
	_ = v180
	var v181 int32
	_ = v181
	var v183 int32
	_ = v183
	var v184 int32
	_ = v184
	var v187 int32
	_ = v187
	var v188 int32
	_ = v188
	var v193 int32
	_ = v193
	var v209 int32
	_ = v209
	var v213 int32
	_ = v213
	var v214 int32
	_ = v214
	var v215 int32
	_ = v215
	var v222 int32
	_ = v222
	var v225 int32
	_ = v225
	var v226 int32
	_ = v226
	var v227 int32
	_ = v227
	var v233 int32
	_ = v233
	var v240 int32
	_ = v240
	var v241 int32
	_ = v241
	var v242 int32
	_ = v242
	var v245 int32
	_ = v245
	var v246 int32
	_ = v246
	var v276 int32
	_ = v276
	var v293 int32
	_ = v293
	var v294 int32
	_ = v294
	var v295 int32
	_ = v295
	var v306 int32
	_ = v306
	var v327 int32
	_ = v327
	var v330 int32
	_ = v330
	var v332 int32
	_ = v332
	var v333 int32
	_ = v333
	var v341 int32
	_ = v341
	var v343 int32
	_ = v343
	var v345 int32
	_ = v345
	var v348 int32
	_ = v348
	var v349 int32
	_ = v349
	var v350 int32
	_ = v350
	var v351 int32
	_ = v351
	var v358 int32
	_ = v358
	var v365 int32
	_ = v365
	var v366 int32
	_ = v366
	var v368 int32
	_ = v368
	var v370 int32
	_ = v370
	var v374 int32
	_ = v374
	var v390 int32
	_ = v390
	var v394 int32
	_ = v394
	var v396 int32
	_ = v396
	var v400 int32
	_ = v400
	var v401 int32
	_ = v401
	var v404 int32
	_ = v404
	var v406 int32
	_ = v406
	var v410 int32
	_ = v410
	var v413 int32
	_ = v413
	var v417 int32
	_ = v417
	var v421 int32
	_ = v421
	var v424 int32
	_ = v424
	var v432 int32
	_ = v432
	var v437 int32
	_ = v437
	var v440 int32
	_ = v440
	var v442 int32
	_ = v442
	var v443 int32
	_ = v443
	var v445 int32
	_ = v445
	var v457 int32
	_ = v457
	var v458 int32
	_ = v458
	var v462 int32
	_ = v462
	var v463 int32
	_ = v463
	var v464 int32
	_ = v464
	var v472 int32
	_ = v472
	var v481 int32
	_ = v481
	var v485 int32
	_ = v485
	var v487 int32
	_ = v487
	var v488 int32
	_ = v488
	var v493 int32
	_ = v493
	var v494 int32
	_ = v494
	var v495 int32
	_ = v495
	var v496 int32
	_ = v496
	var v497 int32
	_ = v497
	var v498 int32
	_ = v498
	var v499 int32
	_ = v499
	var v515 int32
	_ = v515
	var v526 int32
	_ = v526
	var v527 int32
	_ = v527
	var v528 int32
	_ = v528
	var v535 int64
	_ = v535
	var v536 int32
	_ = v536
	var v537 int32
	_ = v537
	var v538 int32
	_ = v538
	var v541 int32
	_ = v541
	var v544 int32
	_ = v544
	var v547 int32
	_ = v547
	var v555 int32
	_ = v555
	var v562 int32
	_ = v562
	var v564 int32
	_ = v564
	var v566 int32
	_ = v566
	var v569 int32
	_ = v569
	var v570 int64
	_ = v570
	var v572 int64
	_ = v572
	var v573 int64
	_ = v573
	var v574 int64
	_ = v574
	var v578 int64
	_ = v578
	var v579 int64
	_ = v579
	var v582 int64
	_ = v582
	var v584 int64
	_ = v584
	var v589 int64
	_ = v589
	var v606 int32
	_ = v606
	var v607 int32
	_ = v607
	var v616 int32
	_ = v616
	var v617 int32
	_ = v617
	var v618 int32
	_ = v618
	var v625 int32
	_ = v625
	var v626 int32
	_ = v626
	var v632 int32
	_ = v632
	var v640 int32
	_ = v640
	var v645 int32
	_ = v645
	var v656 int32
	_ = v656
	var v664 int32
	_ = v664
	var v669 int32
	_ = v669
	var v677 int32
	_ = v677
	var v696 int32
	_ = v696
	var v697 int64
	_ = v697
	var v701 int32
	_ = v701
	var v703 int32
	_ = v703
	var v704 int32
	_ = v704
	var v707 int32
	_ = v707
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
	var v715 int32
	_ = v715
	var v716 int32
	_ = v716
	var v717 int32
	_ = v717
	var v721 int32
	_ = v721
	v2 = int32(0)
	v19 = m.G0
	v21 = v19 - int32(256)
	m.G0 = v21
	v26 = v2
	v27 = v2
	v28 = v2
	v29 = v2
	v30 = v2
	v31 = v2
	v32 = v2
	v33 = int32(-1)
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
	if v33 == int32(1) {
		v493 = v26
		v494 = v27
		v495 = v28
		v496 = v29
		v497 = v30
		v498 = v31
		v499 = v32
		goto L7
	} else {
		goto L8
	}
L5:
	;
	goto L3
L6:
	;
	v696 = int32(m.ExcTag)
	v697 = int64(m.ExcVals[0])
	m.ExcPending = 0
	if v696 == int32(0) {
		goto L115
	} else {
		goto L116
	}
L7:
	;
	if v493 == int32(0) {
		goto L83
	} else {
		goto L84
	}
L8:
	;
	v44 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v45 = *(*int32)(unsafe.Add(mBase, uint32(v44)+16))
	if v45 == int32(0) {
		goto L10
	} else {
		goto L11
	}
L9:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v21)+236)) = v27
	*(*int32)(unsafe.Add(mBase, uint32(v21)+232)) = v28
	*(*int32)(unsafe.Add(mBase, uint32(v21)+240)) = v29
	*(*int32)(unsafe.Add(mBase, uint32(v21)+244)) = v30
	*(*int32)(unsafe.Add(mBase, uint32(v21)+248)) = v31
	*(*int32)(unsafe.Add(mBase, uint32(v21)+252)) = v306
	v327 = *(*int32)(unsafe.Add(mBase, _c_F_fmgr_security_definer[0]))
	*(*int32)(unsafe.Add(mBase, uint32(v21+int32(224)))) = v327
	v330 = *(*int32)(unsafe.Add(mBase, _c_F_fmgr_security_definer[1]))
	*(*int32)(unsafe.Add(mBase, uint32(v21+int32(220)))) = v330
	goto L48
L10:
	;
	v48 = *(*int32)(unsafe.Add(mBase, uint32(v44)+20))
	*(*int32)(unsafe.Add(mBase, uint32(v21)+236)) = v27
	*(*int32)(unsafe.Add(mBase, uint32(v21)+232)) = v28
	*(*int32)(unsafe.Add(mBase, uint32(v21)+240)) = v29
	*(*int32)(unsafe.Add(mBase, uint32(v21)+244)) = v30
	*(*int32)(unsafe.Add(mBase, uint32(v21)+248)) = v31
	*(*int32)(unsafe.Add(mBase, uint32(v21)+252)) = v32
	v56 = F_MemoryContextAllocZero(m, v48, int32(56))
	mBase = m.M
	v57 = m.ExcPending
	if v57 != 0 {
		goto L6
	} else {
		goto L13
	}
L11:
	;
	goto L12
L12:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v21)+228)) = v45
	v306 = v32
	goto L9
L13:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v21)+228)) = v56
	v59 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v60 = *(*int32)(unsafe.Add(mBase, uint32(v59)+4))
	v62 = *(*int32)(unsafe.Add(mBase, uint32(v59)+20))
	*(*int32)(unsafe.Add(mBase, uint32(v21)+236)) = v27
	*(*int32)(unsafe.Add(mBase, uint32(v21)+232)) = v28
	*(*int32)(unsafe.Add(mBase, uint32(v21)+240)) = v29
	*(*int32)(unsafe.Add(mBase, uint32(v21)+244)) = v30
	*(*int32)(unsafe.Add(mBase, uint32(v21)+248)) = v31
	*(*int32)(unsafe.Add(mBase, uint32(v21)+252)) = v32
	F_fmgr_info_cxt_security(m, v60, v56, v62, int32(1))
	mBase = m.M
	v71 = m.ExcPending
	if v71 != 0 {
		goto L6
	} else {
		goto L14
	}
L14:
	;
	v72 = *(*int32)(unsafe.Add(mBase, uint32(v21)+228))
	v73 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v74 = *(*int32)(unsafe.Add(mBase, uint32(v73)+24))
	*(*int32)(unsafe.Add(mBase, uint32(v72)+24)) = v74
	v76 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v77 = int64(*(*uint32)(unsafe.Add(mBase, uint32(v76)+4)))
	*(*int32)(unsafe.Add(mBase, uint32(v21)+236)) = v27
	*(*int32)(unsafe.Add(mBase, uint32(v21)+232)) = v28
	*(*int32)(unsafe.Add(mBase, uint32(v21)+240)) = v29
	*(*int32)(unsafe.Add(mBase, uint32(v21)+244)) = v30
	*(*int32)(unsafe.Add(mBase, uint32(v21)+248)) = v31
	*(*int32)(unsafe.Add(mBase, uint32(v21)+252)) = v32
	v85 = F_SearchSysCache1(m, int32(47), v77)
	mBase = m.M
	v86 = m.ExcPending
	if v86 != 0 {
		goto L6
	} else {
		goto L15
	}
L15:
	;
	if v85 == int32(0) {
		goto L16
	} else {
		goto L17
	}
L16:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v21)+236)) = v27
	*(*int32)(unsafe.Add(mBase, uint32(v21)+232)) = v28
	*(*int32)(unsafe.Add(mBase, uint32(v21)+240)) = v29
	*(*int32)(unsafe.Add(mBase, uint32(v21)+244)) = v30
	*(*int32)(unsafe.Add(mBase, uint32(v21)+248)) = v31
	*(*int32)(unsafe.Add(mBase, uint32(v21)+252)) = v32
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v98 = m.ExcPending
	if v98 != 0 {
		goto L6
	} else {
		goto L19
	}
L17:
	;
	goto L18
L18:
	;
	v122 = *(*int32)(unsafe.Add(mBase, uint32(v85)+16))
	v123 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v122)+22)))
	v124 = v122 + v123
	v125 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v124)+97)))
	if v125 == int32(1) {
		goto L22
	} else {
		goto L23
	}
L19:
	;
	v99 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v100 = *(*int32)(unsafe.Add(mBase, uint32(v99)+4))
	*(*int32)(unsafe.Add(mBase, uint32(v21)+236)) = v27
	*(*int32)(unsafe.Add(mBase, uint32(v21)+232)) = v28
	*(*int32)(unsafe.Add(mBase, uint32(v21)+240)) = v29
	*(*int32)(unsafe.Add(mBase, uint32(v21)+244)) = v30
	*(*int32)(unsafe.Add(mBase, uint32(v21)+248)) = v31
	*(*int32)(unsafe.Add(mBase, uint32(v21)+252)) = v32
	*(*int32)(unsafe.Add(mBase, uint32(v21))) = v100
	F_errmsg_internal(m, int32(_a_F_fmgr_security_definer_0), v21)
	mBase = m.M
	v110 = m.ExcPending
	if v110 != 0 {
		goto L6
	} else {
		goto L20
	}
L20:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v21)+236)) = v27
	*(*int32)(unsafe.Add(mBase, uint32(v21)+232)) = v28
	*(*int32)(unsafe.Add(mBase, uint32(v21)+240)) = v29
	*(*int32)(unsafe.Add(mBase, uint32(v21)+244)) = v30
	*(*int32)(unsafe.Add(mBase, uint32(v21)+248)) = v31
	*(*int32)(unsafe.Add(mBase, uint32(v21)+252)) = v32
	F_errfinish(m, int32(_a_F_fmgr_security_definer_1), int32(666), int32(_a_F_fmgr_security_definer_2))
	mBase = m.M
	v121 = m.ExcPending
	if v121 != 0 {
		goto L6
	} else {
		goto L21
	}
L21:
	;
	goto L3
L22:
	;
	v128 = *(*int32)(unsafe.Add(mBase, uint32(v21)+228))
	v129 = *(*int32)(unsafe.Add(mBase, uint32(v124)+72))
	*(*int32)(unsafe.Add(mBase, uint32(v128)+28)) = v129
	goto L24
L23:
	;
	goto L24
L24:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v21)+236)) = v27
	*(*int32)(unsafe.Add(mBase, uint32(v21)+232)) = v28
	*(*int32)(unsafe.Add(mBase, uint32(v21)+240)) = v29
	*(*int32)(unsafe.Add(mBase, uint32(v21)+244)) = v30
	*(*int32)(unsafe.Add(mBase, uint32(v21)+248)) = v31
	*(*int32)(unsafe.Add(mBase, uint32(v21)+252)) = v32
	v141 = F_SysCacheGetAttr(m, int32(47), v85, int32(29), v21+int32(183))
	mBase = m.M
	v142 = m.ExcPending
	if v142 != 0 {
		goto L6
	} else {
		goto L25
	}
L25:
	;
	v143 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v21)+183)))
	if v143 == int32(0) {
		goto L26
	} else {
		goto L27
	}
L26:
	;
	v146 = int32(_a_F_fmgr_security_definer_3)
	v147 = *(*int32)(unsafe.Add(mBase, _c_F_fmgr_security_definer[2]))
	v149 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v150 = *(*int32)(unsafe.Add(mBase, uint32(v149)+20))
	*(*int32)(unsafe.Add(mBase, _c_F_fmgr_security_definer[2])) = v150
	v152 = base.I32_wrap_i64(v141)
	v153 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v152))))
	if v153&int32(3) != 0 {
		goto L29
	} else {
		goto L30
	}
L27:
	;
	v276 = v32
	goto L28
L28:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v21)+236)) = v27
	*(*int32)(unsafe.Add(mBase, uint32(v21)+232)) = v28
	*(*int32)(unsafe.Add(mBase, uint32(v21)+240)) = v29
	*(*int32)(unsafe.Add(mBase, uint32(v21)+244)) = v30
	*(*int32)(unsafe.Add(mBase, uint32(v21)+248)) = v31
	*(*int32)(unsafe.Add(mBase, uint32(v21)+252)) = v276
	F_ReleaseCatCache(m, v85)
	mBase = m.M
	v293 = m.ExcPending
	if v293 != 0 {
		goto L6
	} else {
		goto L47
	}
L29:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v21)+236)) = v27
	*(*int32)(unsafe.Add(mBase, uint32(v21)+232)) = v28
	*(*int32)(unsafe.Add(mBase, uint32(v21)+240)) = v29
	*(*int32)(unsafe.Add(mBase, uint32(v21)+244)) = v30
	*(*int32)(unsafe.Add(mBase, uint32(v21)+248)) = v31
	*(*int32)(unsafe.Add(mBase, uint32(v21)+252)) = v32
	v162 = F_detoast_attr(m, v152)
	mBase = m.M
	v163 = m.ExcPending
	if v163 != 0 {
		goto L6
	} else {
		goto L32
	}
L30:
	;
	v164 = v152
	v165 = v32
	goto L31
L31:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v21)+232)) = v28
	*(*int32)(unsafe.Add(mBase, uint32(v21)+236)) = v27
	*(*int32)(unsafe.Add(mBase, uint32(v21)+240)) = v29
	*(*int32)(unsafe.Add(mBase, uint32(v21)+244)) = v30
	*(*int32)(unsafe.Add(mBase, uint32(v21)+248)) = v31
	*(*int32)(unsafe.Add(mBase, uint32(v21)+252)) = v165
	v172 = *(*int32)(unsafe.Add(mBase, uint32(v21)+228))
	F_TransformGUCArray(m, v164, v172+int32(32), v172+int32(40))
	mBase = m.M
	v179 = m.ExcPending
	if v179 != 0 {
		goto L6
	} else {
		goto L33
	}
L32:
	;
	v164 = v162
	v165 = v162
	goto L31
L33:
	;
	v180 = *(*int32)(unsafe.Add(mBase, uint32(v21)+228))
	v181 = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v180)+36)) = v181
	v183 = *(*int32)(unsafe.Add(mBase, uint32(v21)+228))
	v184 = *(*int32)(unsafe.Add(mBase, uint32(v183)+32))
	if v184 == v181 {
		goto L34
	} else {
		goto L35
	}
L34:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_fmgr_security_definer[2])) = v147
	v276 = v165
	goto L28
L35:
	;
	v187 = int32(0)
	v188 = *(*int32)(unsafe.Add(mBase, uint32(v184)+4))
	if v188 <= v187 {
		goto L34
	} else {
		goto L36
	}
L36:
	;
	v193 = v187
	goto L37
L37:
	;
	v209 = *(*int32)(unsafe.Add(mBase, uint32(v184)+12))
	v213 = *(*int32)(unsafe.Add(mBase, uint32(v209+v193<<(uint(int32(2))%32))))
	v214 = *(*int32)(unsafe.Add(mBase, uint32(v21)+228))
	v215 = *(*int32)(unsafe.Add(mBase, uint32(v214)+36))
	*(*int32)(unsafe.Add(mBase, uint32(v21)+236)) = v27
	*(*int32)(unsafe.Add(mBase, uint32(v21)+232)) = v28
	*(*int32)(unsafe.Add(mBase, uint32(v21)+240)) = v29
	*(*int32)(unsafe.Add(mBase, uint32(v21)+244)) = v30
	*(*int32)(unsafe.Add(mBase, uint32(v21)+248)) = v31
	*(*int32)(unsafe.Add(mBase, uint32(v21)+252)) = v165
	v222 = int32(0)
	v225 = F_find_option(m, v213, v222, v222, v222)
	mBase = m.M
	v226 = m.ExcPending
	if v226 != 0 {
		goto L6
	} else {
		goto L40
	}
L38:
	;
	goto L34
L39:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v21)+236)) = v27
	*(*int32)(unsafe.Add(mBase, uint32(v21)+232)) = v28
	*(*int32)(unsafe.Add(mBase, uint32(v21)+240)) = v29
	*(*int32)(unsafe.Add(mBase, uint32(v21)+244)) = v30
	*(*int32)(unsafe.Add(mBase, uint32(v21)+248)) = v31
	*(*int32)(unsafe.Add(mBase, uint32(v21)+252)) = v165
	v240 = F_lappend(m, v215, v233)
	mBase = m.M
	v241 = m.ExcPending
	if v241 != 0 {
		goto L6
	} else {
		goto L45
	}
L40:
	;
	if v225 != 0 {
		goto L41
	} else {
		goto L42
	}
L41:
	;
	v227 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v225)+21)))
	if v227&int32(2) == int32(0) {
		v233 = v225
		goto L39
	} else {
		goto L44
	}
L42:
	;
	goto L43
L43:
	;
	v233 = int32(0)
	goto L39
L44:
	;
	goto L43
L45:
	;
	v242 = *(*int32)(unsafe.Add(mBase, uint32(v21)+228))
	*(*int32)(unsafe.Add(mBase, uint32(v242)+36)) = v240
	v245 = v193 + int32(1)
	v246 = *(*int32)(unsafe.Add(mBase, uint32(v184)+4))
	if v245 < v246 {
		v193 = v245
		goto L37
	} else {
		goto L46
	}
L46:
	;
	goto L38
L47:
	;
	v294 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v295 = *(*int32)(unsafe.Add(mBase, uint32(v21)+228))
	*(*int32)(unsafe.Add(mBase, uint32(v294)+16)) = v295
	v306 = v276
	goto L9
L48:
	;
	v332 = *(*int32)(unsafe.Add(mBase, uint32(v21)+228))
	v333 = *(*int32)(unsafe.Add(mBase, uint32(v332)+32))
	if v333 != 0 {
		goto L49
	} else {
		goto L50
	}
L49:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v21)+236)) = v27
	*(*int32)(unsafe.Add(mBase, uint32(v21)+232)) = v28
	*(*int32)(unsafe.Add(mBase, uint32(v21)+240)) = v29
	*(*int32)(unsafe.Add(mBase, uint32(v21)+244)) = v30
	*(*int32)(unsafe.Add(mBase, uint32(v21)+248)) = v31
	*(*int32)(unsafe.Add(mBase, uint32(v21)+252)) = v306
	v341 = int32(_a_F_fmgr_security_definer_4)
	v343 = *(*int32)(unsafe.Add(mBase, _c_F_fmgr_security_definer[3]))
	v345 = v343 + int32(1)
	*(*int32)(unsafe.Add(mBase, _c_F_fmgr_security_definer[3])) = v345
	goto L52
L50:
	;
	v348 = v31
	v349 = int32(0)
	goto L51
L51:
	;
	v350 = *(*int32)(unsafe.Add(mBase, uint32(v21)+228))
	v351 = *(*int32)(unsafe.Add(mBase, uint32(v350)+28))
	if v351 != 0 {
		goto L53
	} else {
		goto L54
	}
L52:
	;
	v348 = v345
	v349 = v345
	goto L51
L53:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v21)+232)) = v28
	*(*int32)(unsafe.Add(mBase, uint32(v21)+236)) = v27
	*(*int32)(unsafe.Add(mBase, uint32(v21)+240)) = v29
	*(*int32)(unsafe.Add(mBase, uint32(v21)+244)) = v349
	*(*int32)(unsafe.Add(mBase, uint32(v21)+248)) = v348
	*(*int32)(unsafe.Add(mBase, uint32(v21)+252)) = v306
	v358 = *(*int32)(unsafe.Add(mBase, uint32(v21)+220))
	*(*int32)(unsafe.Add(mBase, _c_F_fmgr_security_definer[1])) = v358 | int32(1)
	*(*int32)(unsafe.Add(mBase, _c_F_fmgr_security_definer[0])) = v351
	goto L56
L54:
	;
	goto L55
L55:
	;
	v365 = *(*int32)(unsafe.Add(mBase, uint32(v21)+228))
	v366 = *(*int32)(unsafe.Add(mBase, uint32(v365)+32))
	v368 = *(*int32)(unsafe.Add(mBase, uint32(v365)+36))
	v370 = *(*int32)(unsafe.Add(mBase, uint32(v365)+40))
	v374 = int32(0)
	goto L57
L56:
	;
	goto L55
L57:
	;
	v390 = int32(0)
	if v366 == v390 {
		v400 = v390
		goto L59
	} else {
		goto L60
	}
L59:
	;
	v401 = int32(0)
	if v368 == v401 {
		v410 = v401
		goto L62
	} else {
		goto L63
	}
L60:
	;
	v394 = *(*int32)(unsafe.Add(mBase, uint32(v366)+4))
	if v394 <= v374 {
		v400 = int32(0)
		goto L59
	} else {
		goto L61
	}
L61:
	;
	v396 = *(*int32)(unsafe.Add(mBase, uint32(v366)+12))
	v400 = v396 + v374<<(uint(int32(2))%32)
	goto L59
L62:
	;
	if v370 == int32(0) {
		goto L66
	} else {
		goto L67
	}
L63:
	;
	v404 = *(*int32)(unsafe.Add(mBase, uint32(v368)+4))
	if v404 <= v374 {
		v410 = v401
		goto L62
	} else {
		goto L64
	}
L64:
	;
	v406 = *(*int32)(unsafe.Add(mBase, uint32(v368)+12))
	v410 = v406 + v374<<(uint(int32(2))%32)
	goto L62
L65:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v21)+236)) = v27
	*(*int32)(unsafe.Add(mBase, uint32(v21)+232)) = v28
	*(*int32)(unsafe.Add(mBase, uint32(v21)+240)) = v29
	*(*int32)(unsafe.Add(mBase, uint32(v21)+244)) = v349
	*(*int32)(unsafe.Add(mBase, uint32(v21)+248)) = v348
	*(*int32)(unsafe.Add(mBase, uint32(v21)+252)) = v306
	v457 = F_superuser(m)
	mBase = m.M
	v458 = m.ExcPending
	if v458 != 0 {
		goto L6
	} else {
		goto L78
	}
L66:
	;
	v424 = *(*int32)(unsafe.Add(mBase, _c_F_fmgr_security_definer[4]))
	if v424 != 0 {
		goto L70
	} else {
		goto L71
	}
L67:
	;
	v413 = int32(0)
	v417 = *(*int32)(unsafe.Add(mBase, uint32(v370)+4))
	if base.B2i32(v410 == v413)|(base.B2i32(v400 == v413)|base.B2i32(v417 <= v374)) != 0 {
		goto L66
	} else {
		goto L68
	}
L68:
	;
	v421 = *(*int32)(unsafe.Add(mBase, uint32(v370)+12))
	if v421 != 0 {
		goto L65
	} else {
		goto L69
	}
L69:
	;
	goto L66
L70:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v21)+232)) = v28
	*(*int32)(unsafe.Add(mBase, uint32(v21)+236)) = v27
	*(*int32)(unsafe.Add(mBase, uint32(v21)+240)) = v29
	*(*int32)(unsafe.Add(mBase, uint32(v21)+244)) = v349
	*(*int32)(unsafe.Add(mBase, uint32(v21)+248)) = v348
	*(*int32)(unsafe.Add(mBase, uint32(v21)+252)) = v306
	v432 = *(*int32)(unsafe.Add(mBase, uint32(v21)+228))
	m.T0[v424].(func(*base.Module, int32, int32, int32))(m, int32(0), v432, v432+int32(48))
	mBase = m.M
	v437 = m.ExcPending
	if v437 != 0 {
		goto L6
	} else {
		goto L73
	}
L71:
	;
	goto L72
L72:
	;
	v440 = *(*int32)(unsafe.Add(mBase, _c_F_fmgr_security_definer[5]))
	v442 = *(*int32)(unsafe.Add(mBase, _c_F_fmgr_security_definer[6]))
	v443 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	goto L74
L73:
	;
	goto L72
L74:
	;
	v445 = v21 + int32(16)
	*(*int32)(unsafe.Add(mBase, uint32(v445)+4)) = int32(1)
	*(*int32)(unsafe.Add(mBase, uint32(v445))) = v21 + int32(12)
	goto L77
L75:
	;
	v493 = int32(0)
	v494 = v442
	v495 = v440
	v496 = v443
	v497 = v349
	v498 = v348
	v499 = v306
	goto L7
L77:
	;
	goto L75
L78:
	;
	v462 = *(*int32)(unsafe.Add(mBase, uint32(v421+v374<<(uint(int32(2))%32))))
	v463 = *(*int32)(unsafe.Add(mBase, uint32(v410)))
	v464 = *(*int32)(unsafe.Add(mBase, uint32(v400)))
	*(*int32)(unsafe.Add(mBase, uint32(v21)+236)) = v27
	*(*int32)(unsafe.Add(mBase, uint32(v21)+232)) = v28
	*(*int32)(unsafe.Add(mBase, uint32(v21)+240)) = v29
	*(*int32)(unsafe.Add(mBase, uint32(v21)+244)) = v349
	*(*int32)(unsafe.Add(mBase, uint32(v21)+248)) = v348
	*(*int32)(unsafe.Add(mBase, uint32(v21)+252)) = v306
	v472 = *(*int32)(unsafe.Add(mBase, _c_F_fmgr_security_definer[0]))
	*(*int32)(unsafe.Add(mBase, uint32(v21)+236)) = v27
	*(*int32)(unsafe.Add(mBase, uint32(v21)+232)) = v28
	*(*int32)(unsafe.Add(mBase, uint32(v21)+240)) = v29
	*(*int32)(unsafe.Add(mBase, uint32(v21)+244)) = v349
	*(*int32)(unsafe.Add(mBase, uint32(v21)+248)) = v348
	*(*int32)(unsafe.Add(mBase, uint32(v21)+252)) = v306
	if v457 != 0 {
		goto L79
	} else {
		goto L80
	}
L79:
	;
	v481 = int32(5)
	goto L81
L80:
	;
	v481 = int32(6)
	goto L81
L81:
	;
	v485 = int32(0)
	v487 = F_set_config_with_handle(m, v464, v463, v462, v481, int32(13), v472, int32(2), int32(1), v485, v485)
	mBase = m.M
	v488 = m.ExcPending
	if v488 != 0 {
		goto L6
	} else {
		goto L82
	}
L82:
	;
	v374 = v374 + int32(1)
	goto L57
L83:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_fmgr_security_definer[6])) = v21 + int32(16)
	v515 = *(*int32)(unsafe.Add(mBase, uint32(v21)+228))
	*(*int32)(unsafe.Add(mBase, uint32(l0))) = v515
	*(*int32)(unsafe.Add(mBase, uint32(v21)+236)) = v494
	*(*int32)(unsafe.Add(mBase, uint32(v21)+232)) = v495
	*(*int32)(unsafe.Add(mBase, uint32(v21)+240)) = v496
	*(*int32)(unsafe.Add(mBase, uint32(v21)+244)) = v497
	*(*int32)(unsafe.Add(mBase, uint32(v21)+248)) = v498
	*(*int32)(unsafe.Add(mBase, uint32(v21)+252)) = v499
	F_pgstat_init_function_usage(m, l0, v21+int32(184))
	mBase = m.M
	v526 = m.ExcPending
	if v526 != 0 {
		goto L6
	} else {
		goto L86
	}
L84:
	;
	goto L85
L85:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_fmgr_security_definer[5])) = v495
	*(*int32)(unsafe.Add(mBase, _c_F_fmgr_security_definer[6])) = v494
	*(*int32)(unsafe.Add(mBase, uint32(l0))) = v496
	v656 = *(*int32)(unsafe.Add(mBase, _c_F_fmgr_security_definer[4]))
	if v656 != 0 {
		goto L110
	} else {
		goto L111
	}
L86:
	;
	v527 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v528 = *(*int32)(unsafe.Add(mBase, uint32(v527)))
	*(*int32)(unsafe.Add(mBase, uint32(v21)+236)) = v494
	*(*int32)(unsafe.Add(mBase, uint32(v21)+232)) = v495
	*(*int32)(unsafe.Add(mBase, uint32(v21)+240)) = v496
	*(*int32)(unsafe.Add(mBase, uint32(v21)+244)) = v497
	*(*int32)(unsafe.Add(mBase, uint32(v21)+248)) = v498
	*(*int32)(unsafe.Add(mBase, uint32(v21)+252)) = v499
	v535 = m.T0[v528].(func(*base.Module, int32) int64)(m, l0)
	mBase = m.M
	v536 = m.ExcPending
	if v536 != 0 {
		goto L6
	} else {
		goto L87
	}
L87:
	;
	v537 = int32(1)
	v538 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	if v538 == int32(0) {
		v547 = v537
		goto L88
	} else {
		goto L89
	}
L88:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v21)+236)) = v494
	*(*int32)(unsafe.Add(mBase, uint32(v21)+232)) = v495
	*(*int32)(unsafe.Add(mBase, uint32(v21)+240)) = v496
	*(*int32)(unsafe.Add(mBase, uint32(v21)+244)) = v497
	*(*int32)(unsafe.Add(mBase, uint32(v21)+248)) = v498
	*(*int32)(unsafe.Add(mBase, uint32(v21)+252)) = v499
	v555 = v21 + int32(184)
	v562 = m.G0
	v564 = v562 - int32(16)
	m.G0 = v564
	v566 = *(*int32)(unsafe.Add(mBase, uint32(v555)))
	if v566 != 0 {
		goto L92
	} else {
		goto L93
	}
L89:
	;
	v541 = *(*int32)(unsafe.Add(mBase, uint32(v538)))
	if v541 != int32(389) {
		v547 = v537
		goto L88
	} else {
		goto L90
	}
L90:
	;
	v544 = *(*int32)(unsafe.Add(mBase, uint32(v538)+20))
	v547 = base.B2i32(v544 != int32(1))
	goto L88
L91:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_fmgr_security_definer[5])) = v495
	*(*int32)(unsafe.Add(mBase, _c_F_fmgr_security_definer[6])) = v494
	*(*int32)(unsafe.Add(mBase, uint32(l0))) = v496
	v606 = *(*int32)(unsafe.Add(mBase, uint32(v21)+228))
	v607 = *(*int32)(unsafe.Add(mBase, uint32(v606)+32))
	if v607 != 0 {
		goto L98
	} else {
		goto L99
	}
L92:
	;
	F___clock_gettime(m, int32(1), v564)
	mBase = m.M
	v569 = int32(_a_F_fmgr_security_definer_5)
	v570 = *(*int64)(unsafe.Add(mBase, _c_F_fmgr_security_definer[7]))
	v572 = *(*int64)(unsafe.Add(mBase, uint32(v555)+16))
	v573 = int64(*(*int32)(unsafe.Add(mBase, uint32(v564)+8)))
	v574 = *(*int64)(unsafe.Add(mBase, uint32(v564)))
	v578 = *(*int64)(unsafe.Add(mBase, uint32(v555)+24))
	v579 = v573 + v574*int64(1000000000) - v578
	*(*int64)(unsafe.Add(mBase, _c_F_fmgr_security_definer[7])) = v572 + v579
	v582 = *(*int64)(unsafe.Add(mBase, uint32(v555)+8))
	if v547 != 0 {
		goto L95
	} else {
		goto L96
	}
L93:
	;
	goto L94
L94:
	;
	m.G0 = v564 + int32(16)
	goto L91
L95:
	;
	v584 = *(*int64)(unsafe.Add(mBase, uint32(v566)))
	*(*int64)(unsafe.Add(mBase, uint32(v566))) = v584 + int64(1)
	goto L97
L96:
	;
	goto L97
L97:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v566)+8)) = v582 + v579
	v589 = *(*int64)(unsafe.Add(mBase, uint32(v566)+16))
	*(*int64)(unsafe.Add(mBase, uint32(v566)+16)) = v589 + (v579 - v570 + v572)
	goto L94
L98:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v21)+236)) = v494
	*(*int32)(unsafe.Add(mBase, uint32(v21)+232)) = v495
	*(*int32)(unsafe.Add(mBase, uint32(v21)+240)) = v496
	*(*int32)(unsafe.Add(mBase, uint32(v21)+244)) = v497
	*(*int32)(unsafe.Add(mBase, uint32(v21)+248)) = v498
	*(*int32)(unsafe.Add(mBase, uint32(v21)+252)) = v499
	F_AtEOXact_GUC(m, int32(1), v497)
	mBase = m.M
	v616 = m.ExcPending
	if v616 != 0 {
		goto L6
	} else {
		goto L101
	}
L99:
	;
	goto L100
L100:
	;
	v617 = *(*int32)(unsafe.Add(mBase, uint32(v21)+228))
	v618 = *(*int32)(unsafe.Add(mBase, uint32(v617)+28))
	if v618 != 0 {
		goto L102
	} else {
		goto L103
	}
L101:
	;
	goto L100
L102:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v21)+232)) = v495
	*(*int32)(unsafe.Add(mBase, uint32(v21)+236)) = v494
	*(*int32)(unsafe.Add(mBase, uint32(v21)+240)) = v496
	*(*int32)(unsafe.Add(mBase, uint32(v21)+244)) = v497
	*(*int32)(unsafe.Add(mBase, uint32(v21)+248)) = v498
	*(*int32)(unsafe.Add(mBase, uint32(v21)+252)) = v499
	v625 = *(*int32)(unsafe.Add(mBase, uint32(v21)+224))
	v626 = *(*int32)(unsafe.Add(mBase, uint32(v21)+220))
	*(*int32)(unsafe.Add(mBase, _c_F_fmgr_security_definer[1])) = v626
	*(*int32)(unsafe.Add(mBase, _c_F_fmgr_security_definer[0])) = v625
	goto L105
L103:
	;
	goto L104
L104:
	;
	v632 = *(*int32)(unsafe.Add(mBase, _c_F_fmgr_security_definer[4]))
	if v632 != 0 {
		goto L106
	} else {
		goto L107
	}
L105:
	;
	goto L104
L106:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v21)+232)) = v495
	*(*int32)(unsafe.Add(mBase, uint32(v21)+236)) = v494
	*(*int32)(unsafe.Add(mBase, uint32(v21)+240)) = v496
	*(*int32)(unsafe.Add(mBase, uint32(v21)+244)) = v497
	*(*int32)(unsafe.Add(mBase, uint32(v21)+248)) = v498
	*(*int32)(unsafe.Add(mBase, uint32(v21)+252)) = v499
	v640 = *(*int32)(unsafe.Add(mBase, uint32(v21)+228))
	m.T0[v632].(func(*base.Module, int32, int32, int32))(m, int32(1), v640, v640+int32(48))
	mBase = m.M
	v645 = m.ExcPending
	if v645 != 0 {
		goto L6
	} else {
		goto L109
	}
L107:
	;
	goto L108
L108:
	;
	m.G0 = v21 + int32(256)
	return v535
L109:
	;
	goto L108
L110:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v21)+232)) = v495
	*(*int32)(unsafe.Add(mBase, uint32(v21)+236)) = v494
	*(*int32)(unsafe.Add(mBase, uint32(v21)+240)) = v496
	*(*int32)(unsafe.Add(mBase, uint32(v21)+244)) = v497
	*(*int32)(unsafe.Add(mBase, uint32(v21)+248)) = v498
	*(*int32)(unsafe.Add(mBase, uint32(v21)+252)) = v499
	v664 = *(*int32)(unsafe.Add(mBase, uint32(v21)+228))
	m.T0[v656].(func(*base.Module, int32, int32, int32))(m, int32(2), v664, v664+int32(48))
	mBase = m.M
	v669 = m.ExcPending
	if v669 != 0 {
		goto L6
	} else {
		goto L113
	}
L111:
	;
	goto L112
L112:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v21)+236)) = v494
	*(*int32)(unsafe.Add(mBase, uint32(v21)+232)) = v495
	*(*int32)(unsafe.Add(mBase, uint32(v21)+240)) = v496
	*(*int32)(unsafe.Add(mBase, uint32(v21)+244)) = v497
	*(*int32)(unsafe.Add(mBase, uint32(v21)+248)) = v498
	*(*int32)(unsafe.Add(mBase, uint32(v21)+252)) = v499
	F_pg_re_throw(m)
	mBase = m.M
	v677 = m.ExcPending
	if v677 != 0 {
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
	v701 = int32(v697)
	m.G0 = v21
	v703 = *(*int32)(unsafe.Add(mBase, uint32(v701)+4))
	v704 = *(*int32)(unsafe.Add(mBase, uint32(v701)))
	v707 = *(*int32)(unsafe.Add(mBase, uint32(v704)))
	if v21+int32(12) == v707 {
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
	if v711 != 0 {
		goto L121
	} else {
		goto L122
	}
L118:
	;
	v709 = *(*int32)(unsafe.Add(mBase, uint32(v704)+4))
	v711 = v709
	goto L120
L119:
	;
	v711 = int32(0)
	goto L120
L120:
	;
	goto L117
L121:
	;
	v712 = *(*int32)(unsafe.Add(mBase, uint32(v21)+252))
	v713 = *(*int32)(unsafe.Add(mBase, uint32(v21)+248))
	v714 = *(*int32)(unsafe.Add(mBase, uint32(v21)+244))
	v715 = *(*int32)(unsafe.Add(mBase, uint32(v21)+240))
	v716 = *(*int32)(unsafe.Add(mBase, uint32(v21)+236))
	v717 = *(*int32)(unsafe.Add(mBase, uint32(v21)+232))
	v26 = v703
	v27 = v716
	v28 = v717
	v29 = v715
	v30 = v714
	v31 = v713
	v32 = v712
	v33 = v711
	goto L1
L122:
	;
	goto L123
L123:
	;
	F___wasm_longjmp(m, v704, v703)
	mBase = m.M
	v721 = m.ExcPending
	if v721 != 0 {
		goto L124
	} else {
		goto L125
	}
L124:
	;
	return int64(0)
L125:
	;
	base.Wasm_trap_unreachable()
	for {
	}
}
