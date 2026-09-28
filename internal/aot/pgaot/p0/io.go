package p0

import (
	base "github.com/shibukawa/pgmem/internal/aot/pgaot/base"
	"unsafe"
)

func F_check_io_max_concurrency(m *base.Module, l0 int32, l1 int32, l2 int32) int32 {
	mBase := m.M
	_ = mBase
	var v4 int32
	_ = v4
	var v8 int32
	_ = v8
	var v14 int32
	_ = v14
	var v17 int32
	_ = v17
	v4 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	if v4 == int32(0) {
		v8 = *(*int32)(unsafe.Add(mBase, _c_F_check_io_max_concurrency[0]))
		*(*int32)(unsafe.Add(mBase, _c_F_check_io_max_concurrency[1])) = v8
		v14 = F_format_elog_string(m, int32(_a_F_check_io_max_concurrency_0), int32(0))
		mBase = m.M
		v17 = m.ExcPending
		if v17 != 0 {
			return int32(0)
		} else {
			*(*int32)(unsafe.Add(mBase, _c_F_check_io_max_concurrency[2])) = v14
			return base.B2i32(v4 != int32(0))
		}
	} else {
		return base.B2i32(v4 != int32(0))
	}
}
func F_print_io_usage(m *base.Module, l0 int32, l1 int32) {
	mBase := m.M
	_ = mBase
	var v10 int32
	_ = v10
	var v12 int32
	_ = v12
	var v14 int64
	_ = v14
	var v17 int32
	_ = v17
	var v21 int32
	_ = v21
	var v22 int32
	_ = v22
	var v23 int64
	_ = v23
	var v24 int64
	_ = v24
	var v25 int32
	_ = v25
	var v26 int32
	_ = v26
	var v37 int32
	_ = v37
	var v38 int64
	_ = v38
	var v42 int32
	_ = v42
	var v43 int32
	_ = v43
	var v44 int64
	_ = v44
	var v45 int64
	_ = v45
	var v46 int64
	_ = v46
	var v48 int64
	_ = v48
	var v49 float64
	_ = v49
	var v59 int32
	_ = v59
	var v62 int64
	_ = v62
	var v68 int32
	_ = v68
	var v71 int64
	_ = v71
	var v73 int32
	_ = v73
	var v76 int64
	_ = v76
	var v78 int32
	_ = v78
	var v81 int64
	_ = v81
	var v83 int32
	_ = v83
	var v86 int64
	_ = v86
	var v88 int32
	_ = v88
	var v91 int64
	_ = v91
	var v94 int64
	_ = v94
	var v98 float64
	_ = v98
	var v102 int32
	_ = v102
	var v105 int64
	_ = v105
	var v108 int64
	_ = v108
	var v112 float64
	_ = v112
	var v116 int32
	_ = v116
	v10 = m.G0
	v12 = v10 - int32(48)
	m.G0 = v12
	v14 = *(*int64)(unsafe.Add(mBase, uint32(l1)))
	if v14 == int64(0) {
		m.G0 = v12 + int32(48)
		return
	} else {
		v17 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
		if v17 == int32(0) {
			F_ExplainIndentText(m, l0)
			mBase = m.M
			v21 = m.ExcPending
			if v21 != 0 {
				return
			} else {
				v22 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
				v23 = *(*int64)(unsafe.Add(mBase, uint32(l1)+8))
				v24 = *(*int64)(unsafe.Add(mBase, uint32(l1)))
				v25 = int32(*(*int16)(unsafe.Add(mBase, uint32(l1)+16)))
				v26 = int32(*(*int16)(unsafe.Add(mBase, uint32(l1)+18)))
				*(*int32)(unsafe.Add(mBase, uint32(v12)+44)) = v26
				*(*int32)(unsafe.Add(mBase, uint32(v12)+40)) = v25
				*(*float64)(unsafe.Add(mBase, uint32(v12)+32)) = base.F64_div(base.F64_convert_i64_u(v23), base.F64_convert_i64_u(v24))
				F_appendStringInfo(m, v22, int32(_a_F_print_io_usage_0), v12+int32(32))
				mBase = m.M
				v37 = m.ExcPending
				if v37 != 0 {
					return
				} else {
					v38 = *(*int64)(unsafe.Add(mBase, uint32(l1)+32))
					if v38 == int64(0) {
						m.G0 = v12 + int32(48)
						return
					} else {
						F_ExplainIndentText(m, l0)
						mBase = m.M
						v42 = m.ExcPending
						if v42 != 0 {
							return
						} else {
							v43 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
							v44 = *(*int64)(unsafe.Add(mBase, uint32(l1)+24))
							v45 = *(*int64)(unsafe.Add(mBase, uint32(l1)+40))
							v46 = *(*int64)(unsafe.Add(mBase, uint32(l1)+48))
							v48 = *(*int64)(unsafe.Add(mBase, uint32(l1)+32))
							v49 = base.F64_convert_i64_u(v48)
							*(*float64)(unsafe.Add(mBase, uint32(v12)+24)) = base.F64_div(base.F64_convert_i64_u(v46), v49)
							*(*float64)(unsafe.Add(mBase, uint32(v12)+16)) = base.F64_div(base.F64_convert_i64_u(v45), v49)
							*(*int64)(unsafe.Add(mBase, uint32(v12)+8)) = v44
							*(*int64)(unsafe.Add(mBase, uint32(v12))) = v48
							F_appendStringInfo(m, v43, int32(_a_F_print_io_usage_1), v12)
							mBase = m.M
							v59 = m.ExcPending
							if v59 != 0 {
								return
							} else {
								m.G0 = v12 + int32(48)
								return
							}
						}
					}
				}
			}
		} else {
			v62 = *(*int64)(unsafe.Add(mBase, uint32(l1)+8))
			F_ExplainPropertyFloat(m, int32(_a_F_print_io_usage_2), int32(0), base.F64_div(base.F64_convert_i64_u(v62), base.F64_convert_i64_u(v14)), int32(3), l0)
			mBase = m.M
			v68 = m.ExcPending
			if v68 != 0 {
				return
			} else {
				v71 = int64(*(*int16)(unsafe.Add(mBase, uint32(l1)+16)))
				F_ExplainPropertyInteger(m, int32(_a_F_print_io_usage_3), int32(0), v71, l0)
				mBase = m.M
				v73 = m.ExcPending
				if v73 != 0 {
					return
				} else {
					v76 = int64(*(*int16)(unsafe.Add(mBase, uint32(l1)+18)))
					F_ExplainPropertyInteger(m, int32(_a_F_print_io_usage_4), int32(0), v76, l0)
					mBase = m.M
					v78 = m.ExcPending
					if v78 != 0 {
						return
					} else {
						v81 = *(*int64)(unsafe.Add(mBase, uint32(l1)+32))
						F_ExplainPropertyUInteger(m, int32(_a_F_print_io_usage_5), int32(0), v81, l0)
						mBase = m.M
						v83 = m.ExcPending
						if v83 != 0 {
							return
						} else {
							v86 = *(*int64)(unsafe.Add(mBase, uint32(l1)+24))
							F_ExplainPropertyUInteger(m, int32(_a_F_print_io_usage_6), int32(0), v86, l0)
							mBase = m.M
							v88 = m.ExcPending
							if v88 != 0 {
								return
							} else {
								v91 = *(*int64)(unsafe.Add(mBase, uint32(l1)+40))
								v94 = *(*int64)(unsafe.Add(mBase, uint32(l1)+32))
								if v94 == int64(0) {
									v98 = float64(1)
								} else {
									v98 = base.F64_convert_i64_u(v94)
								}
								F_ExplainPropertyFloat(m, int32(_a_F_print_io_usage_7), int32(0), base.F64_div(base.F64_convert_i64_u(v91), v98), int32(3), l0)
								mBase = m.M
								v102 = m.ExcPending
								if v102 != 0 {
									return
								} else {
									v105 = *(*int64)(unsafe.Add(mBase, uint32(l1)+48))
									v108 = *(*int64)(unsafe.Add(mBase, uint32(l1)+32))
									if v108 == int64(0) {
										v112 = float64(1)
									} else {
										v112 = base.F64_convert_i64_u(v108)
									}
									F_ExplainPropertyFloat(m, int32(_a_F_print_io_usage_8), int32(0), base.F64_div(base.F64_convert_i64_u(v105), v112), int32(3), l0)
									mBase = m.M
									v116 = m.ExcPending
									if v116 != 0 {
										return
									} else {
										m.G0 = v12 + int32(48)
										return
									}
								}
							}
						}
					}
				}
			}
		}
	}
}
func F_show_scan_io_usage(m *base.Module, l0 int32, l1 int32) {
	mBase := m.M
	_ = mBase
	var v3 int32
	_ = v3
	var v16 int64
	_ = v16
	var v28 int32
	_ = v28
	var v30 int32
	_ = v30
	var v32 int32
	_ = v32
	var v47 int32
	_ = v47
	var v50 int32
	_ = v50
	var v53 int32
	_ = v53
	var v56 int64
	_ = v56
	var v58 int64
	_ = v58
	var v60 int64
	_ = v60
	var v62 int64
	_ = v62
	var v64 int64
	_ = v64
	var v66 int64
	_ = v66
	var v68 int64
	_ = v68
	var v71 int32
	_ = v71
	var v74 int32
	_ = v74
	var v77 int64
	_ = v77
	var v78 int32
	_ = v78
	var v81 int64
	_ = v81
	var v82 int64
	_ = v82
	var v83 int64
	_ = v83
	var v84 int64
	_ = v84
	var v85 int32
	_ = v85
	var v86 int32
	_ = v86
	var v87 int64
	_ = v87
	var v91 int32
	_ = v91
	var v92 int32
	_ = v92
	var v94 int32
	_ = v94
	var v95 int32
	_ = v95
	var v103 int64
	_ = v103
	var v104 int64
	_ = v104
	var v105 int64
	_ = v105
	var v106 int64
	_ = v106
	var v107 int64
	_ = v107
	var v108 int64
	_ = v108
	var v117 int32
	_ = v117
	var v118 int32
	_ = v118
	var v119 int32
	_ = v119
	var v121 int32
	_ = v121
	var v122 int32
	_ = v122
	var v124 int64
	_ = v124
	var v127 int64
	_ = v127
	var v128 int64
	_ = v128
	var v129 int64
	_ = v129
	var v130 int64
	_ = v130
	var v132 int32
	_ = v132
	var v133 int64
	_ = v133
	var v134 int32
	_ = v134
	var v136 int32
	_ = v136
	var v138 int32
	_ = v138
	var v139 int32
	_ = v139
	var v140 int32
	_ = v140
	var v145 int32
	_ = v145
	var v146 int32
	_ = v146
	var v149 int32
	_ = v149
	var v150 int32
	_ = v150
	var v155 int32
	_ = v155
	var v158 int32
	_ = v158
	var v159 int32
	_ = v159
	var v182 int32
	_ = v182
	var v186 int32
	_ = v186
	var v190 int32
	_ = v190
	var v193 int32
	_ = v193
	var v195 int32
	_ = v195
	var v198 int32
	_ = v198
	var v228 int32
	_ = v228
	var v259 int32
	_ = v259
	var v261 int32
	_ = v261
	var v269 int32
	_ = v269
	var v289 int32
	_ = v289
	var v290 int32
	_ = v290
	var v291 int64
	_ = v291
	var v292 int64
	_ = v292
	var v293 int64
	_ = v293
	var v294 int64
	_ = v294
	var v295 int64
	_ = v295
	var v296 int64
	_ = v296
	var v298 int32
	_ = v298
	var v307 int32
	_ = v307
	var v310 int64
	_ = v310
	var v311 int32
	_ = v311
	var v316 int64
	_ = v316
	var v317 int64
	_ = v317
	var v318 int64
	_ = v318
	var v319 int64
	_ = v319
	var v320 int32
	_ = v320
	var v321 int32
	_ = v321
	var v322 int64
	_ = v322
	var v326 int32
	_ = v326
	var v327 int32
	_ = v327
	var v328 int32
	_ = v328
	var v329 int32
	_ = v329
	var v338 int64
	_ = v338
	var v339 int64
	_ = v339
	var v340 int64
	_ = v340
	var v341 int64
	_ = v341
	var v342 int64
	_ = v342
	var v343 int64
	_ = v343
	var v352 int32
	_ = v352
	var v353 int32
	_ = v353
	var v354 int32
	_ = v354
	var v356 int32
	_ = v356
	var v357 int32
	_ = v357
	var v359 int64
	_ = v359
	var v360 int64
	_ = v360
	var v361 int64
	_ = v361
	var v362 int64
	_ = v362
	var v363 int64
	_ = v363
	var v364 int64
	_ = v364
	var v365 int32
	_ = v365
	var v367 int32
	_ = v367
	var v369 int32
	_ = v369
	var v370 int32
	_ = v370
	var v371 int32
	_ = v371
	var v376 int32
	_ = v376
	var v377 int32
	_ = v377
	var v380 int32
	_ = v380
	var v381 int32
	_ = v381
	var v386 int32
	_ = v386
	var v389 int32
	_ = v389
	var v390 int32
	_ = v390
	var v413 int32
	_ = v413
	var v417 int32
	_ = v417
	var v421 int32
	_ = v421
	var v424 int32
	_ = v424
	var v426 int32
	_ = v426
	var v429 int32
	_ = v429
	var v459 int32
	_ = v459
	var v490 int32
	_ = v490
	var v492 int32
	_ = v492
	var v498 int32
	_ = v498
	var v520 int32
	_ = v520
	var v521 int32
	_ = v521
	var v522 int64
	_ = v522
	var v523 int64
	_ = v523
	var v524 int64
	_ = v524
	var v525 int64
	_ = v525
	var v526 int64
	_ = v526
	var v527 int64
	_ = v527
	var v529 int32
	_ = v529
	var v538 int32
	_ = v538
	var v541 int64
	_ = v541
	var v542 int32
	_ = v542
	var v547 int64
	_ = v547
	var v548 int64
	_ = v548
	var v549 int64
	_ = v549
	var v550 int64
	_ = v550
	var v551 int32
	_ = v551
	var v552 int32
	_ = v552
	var v553 int64
	_ = v553
	var v557 int32
	_ = v557
	var v558 int32
	_ = v558
	var v559 int32
	_ = v559
	var v560 int32
	_ = v560
	var v569 int64
	_ = v569
	var v570 int64
	_ = v570
	var v571 int64
	_ = v571
	var v572 int64
	_ = v572
	var v573 int64
	_ = v573
	var v574 int64
	_ = v574
	var v583 int32
	_ = v583
	var v584 int32
	_ = v584
	var v585 int32
	_ = v585
	var v587 int32
	_ = v587
	var v588 int32
	_ = v588
	var v590 int64
	_ = v590
	var v591 int64
	_ = v591
	var v592 int64
	_ = v592
	var v593 int64
	_ = v593
	var v594 int64
	_ = v594
	var v595 int64
	_ = v595
	var v596 int32
	_ = v596
	var v598 int32
	_ = v598
	var v600 int32
	_ = v600
	var v601 int32
	_ = v601
	var v602 int32
	_ = v602
	var v607 int32
	_ = v607
	var v608 int32
	_ = v608
	var v611 int32
	_ = v611
	var v612 int32
	_ = v612
	var v617 int32
	_ = v617
	var v620 int32
	_ = v620
	var v621 int32
	_ = v621
	var v644 int32
	_ = v644
	var v648 int32
	_ = v648
	var v652 int32
	_ = v652
	var v655 int32
	_ = v655
	var v657 int32
	_ = v657
	var v660 int32
	_ = v660
	var v690 int32
	_ = v690
	var v721 int32
	_ = v721
	var v723 int32
	_ = v723
	var v729 int32
	_ = v729
	var v751 int32
	_ = v751
	var v752 int32
	_ = v752
	var v753 int64
	_ = v753
	var v754 int64
	_ = v754
	var v755 int64
	_ = v755
	var v756 int64
	_ = v756
	var v757 int64
	_ = v757
	var v758 int64
	_ = v758
	var v760 int32
	_ = v760
	var v789 int64
	_ = v789
	var v827 int32
	_ = v827
	v3 = int32(0)
	v16 = int64(0)
	v28 = m.G0
	v30 = v28 + int32(-64)
	m.G0 = v30
	v32 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	*(*int64)(unsafe.Add(mBase, uint32(v30)+56)) = v16
	*(*int64)(unsafe.Add(mBase, uint32(v30)+48)) = v16
	*(*int64)(unsafe.Add(mBase, uint32(v30)+40)) = v16
	*(*int64)(unsafe.Add(mBase, uint32(v30)+32)) = v16
	*(*int64)(unsafe.Add(mBase, uint32(v30)+24)) = v16
	*(*int64)(unsafe.Add(mBase, uint32(v30)+16)) = v16
	*(*int64)(unsafe.Add(mBase, uint32(v30)+8)) = v16
	v47 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l1)+13)))
	if v47 != int32(1) {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	m.G0 = v30 - int32(-64)
	return
L2:
	;
	v50 = *(*int32)(unsafe.Add(mBase, uint32(l0)+108))
	if v50 == int32(0) {
		goto L3
	} else {
		goto L4
	}
L3:
	;
	v71 = *(*int32)(unsafe.Add(mBase, uint32(v32)))
	switch v71 - int32(343) {
	case 0:
		goto L9
	default:
		goto L1
	case 5:
		goto L10
	case 7:
		goto L8
	}
L4:
	;
	v53 = *(*int32)(unsafe.Add(mBase, uint32(v50)+36))
	if v53 == int32(0) {
		goto L3
	} else {
		goto L5
	}
L5:
	;
	v56 = *(*int64)(unsafe.Add(mBase, uint32(v53)+48))
	*(*int64)(unsafe.Add(mBase, uint32(v30)+56)) = v56
	v58 = *(*int64)(unsafe.Add(mBase, uint32(v53)+40))
	*(*int64)(unsafe.Add(mBase, uint32(v30)+48)) = v58
	v60 = *(*int64)(unsafe.Add(mBase, uint32(v53)+32))
	*(*int64)(unsafe.Add(mBase, uint32(v30)+40)) = v60
	v62 = *(*int64)(unsafe.Add(mBase, uint32(v53)+24))
	*(*int64)(unsafe.Add(mBase, uint32(v30)+32)) = v62
	v64 = *(*int64)(unsafe.Add(mBase, uint32(v53)+16))
	*(*int64)(unsafe.Add(mBase, uint32(v30)+24)) = v64
	v66 = *(*int64)(unsafe.Add(mBase, uint32(v53)+8))
	*(*int64)(unsafe.Add(mBase, uint32(v30)+16)) = v66
	v68 = *(*int64)(unsafe.Add(mBase, uint32(v53)))
	*(*int64)(unsafe.Add(mBase, uint32(v30)+8)) = v68
	goto L3
L6:
	;
	F_print_io_usage(m, l1, v28+int32(-56))
	mBase = m.M
	v827 = m.ExcPending
	if v827 != 0 {
		goto L18
	} else {
		goto L90
	}
L7:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v30)+8)) = v789
	goto L6
L8:
	;
	v538 = *(*int32)(unsafe.Add(mBase, uint32(l0)+140))
	if v538 == int32(0) {
		goto L6
	} else {
		goto L64
	}
L9:
	;
	v307 = *(*int32)(unsafe.Add(mBase, uint32(l0)+120))
	if v307 == int32(0) {
		goto L6
	} else {
		goto L38
	}
L10:
	;
	v74 = *(*int32)(unsafe.Add(mBase, uint32(l0)+208))
	if v74 == int32(0) {
		goto L6
	} else {
		goto L11
	}
L11:
	;
	v77 = *(*int64)(unsafe.Add(mBase, uint32(v30)+8))
	v78 = *(*int32)(unsafe.Add(mBase, uint32(v74)))
	if v78 <= int32(0) {
		v789 = v77
		goto L7
	} else {
		goto L12
	}
L12:
	;
	v81 = *(*int64)(unsafe.Add(mBase, uint32(v30)+56))
	v82 = *(*int64)(unsafe.Add(mBase, uint32(v30)+48))
	v83 = *(*int64)(unsafe.Add(mBase, uint32(v30)+40))
	v84 = *(*int64)(unsafe.Add(mBase, uint32(v30)+32))
	v85 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v30)+26)))
	v86 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v30)+24)))
	v87 = *(*int64)(unsafe.Add(mBase, uint32(v30)+16))
	v91 = v85
	v92 = v86
	v94 = v3
	v95 = v78
	v103 = v81
	v104 = v82
	v105 = v83
	v106 = v84
	v107 = v87
	v108 = v77
	goto L13
L13:
	;
	v117 = v74 + v94*int32(72)
	v118 = int32(*(*int16)(unsafe.Add(mBase, uint32(v117)+42)))
	v119 = base.I32_extend16_s(v91)
	v121 = int32(*(*int16)(unsafe.Add(mBase, uint32(v117)+40)))
	v122 = base.I32_extend16_s(v92)
	v124 = *(*int64)(unsafe.Add(mBase, uint32(v117)+72))
	v127 = *(*int64)(unsafe.Add(mBase, uint32(v117-int32(-64))))
	v128 = *(*int64)(unsafe.Add(mBase, uint32(v117)+56))
	v129 = *(*int64)(unsafe.Add(mBase, uint32(v117)+48))
	v130 = *(*int64)(unsafe.Add(mBase, uint32(v117)+32))
	v132 = v117 + int32(24)
	v133 = *(*int64)(unsafe.Add(mBase, uint32(v132)))
	v134 = *(*int32)(unsafe.Add(mBase, uint32(l1)+60))
	if v134 != 0 {
		goto L15
	} else {
		goto L16
	}
L14:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v30)+56)) = v291
	*(*int64)(unsafe.Add(mBase, uint32(v30)+48)) = v292
	*(*int64)(unsafe.Add(mBase, uint32(v30)+40)) = v293
	*(*int64)(unsafe.Add(mBase, uint32(v30)+32)) = v294
	*(*uint16)(unsafe.Add(mBase, uint32(v30)+26)) = uint16(v289)
	*(*uint16)(unsafe.Add(mBase, uint32(v30)+24)) = uint16(v290)
	*(*int64)(unsafe.Add(mBase, uint32(v30)+16)) = v295
	v789 = v296
	goto L7
L15:
	;
	F_ExplainOpenWorker(m, v94, l1)
	mBase = m.M
	v136 = m.ExcPending
	if v136 != 0 {
		goto L18
	} else {
		goto L19
	}
L16:
	;
	v269 = v95
	goto L17
L17:
	;
	if v119 < v118 {
		goto L31
	} else {
		goto L32
	}
L18:
	;
	return
L19:
	;
	F_print_io_usage(m, l1, v132)
	mBase = m.M
	v138 = m.ExcPending
	if v138 != 0 {
		goto L18
	} else {
		goto L20
	}
L20:
	;
	v139 = *(*int32)(unsafe.Add(mBase, uint32(l1)+60))
	v140 = *(*int32)(unsafe.Add(mBase, uint32(v139)+12))
	F_ExplainSaveGroup(m, l1, v140+v94<<(uint(int32(2))%32))
	mBase = m.M
	v145 = m.ExcPending
	if v145 != 0 {
		goto L18
	} else {
		goto L21
	}
L21:
	;
	v146 = *(*int32)(unsafe.Add(mBase, uint32(l1)+20))
	if v146 == int32(0) {
		goto L22
	} else {
		goto L23
	}
L22:
	;
	v149 = *(*int32)(unsafe.Add(mBase, uint32(l1)))
	v150 = *(*int32)(unsafe.Add(mBase, uint32(v149)+4))
	if v150 <= int32(0) {
		goto L25
	} else {
		goto L26
	}
L23:
	;
	goto L24
L24:
	;
	v259 = *(*int32)(unsafe.Add(mBase, uint32(v139)+16))
	*(*int32)(unsafe.Add(mBase, uint32(l1))) = v259
	v261 = *(*int32)(unsafe.Add(mBase, uint32(v74)))
	v269 = v261
	goto L17
L25:
	;
	v228 = *(*int32)(unsafe.Add(mBase, uint32(l1)+24))
	*(*int32)(unsafe.Add(mBase, uint32(l1)+24)) = v228 - int32(1)
	goto L24
L26:
	;
	v155 = v149
	v158 = v150
	v159 = v149 + int32(4)
	goto L27
L27:
	;
	v182 = *(*int32)(unsafe.Add(mBase, uint32(v155)))
	v186 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v182+v158-int32(1)))))
	if v186 == int32(10) {
		goto L25
	} else {
		goto L29
	}
L28:
	;
	goto L25
L29:
	;
	v190 = v158 - int32(1)
	*(*int32)(unsafe.Add(mBase, uint32(v159))) = v190
	v193 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v182+v190))) = uint8(v193)
	v195 = *(*int32)(unsafe.Add(mBase, uint32(l1)))
	v198 = *(*int32)(unsafe.Add(mBase, uint32(v195)+4))
	if v193 < v198 {
		v155 = v195
		v158 = v198
		v159 = v195 + int32(4)
		goto L27
	} else {
		goto L30
	}
L30:
	;
	goto L28
L31:
	;
	v289 = v118
	goto L33
L32:
	;
	v289 = v119
	goto L33
L33:
	;
	if v122 < v121 {
		goto L34
	} else {
		goto L35
	}
L34:
	;
	v290 = v121
	goto L36
L35:
	;
	v290 = v122
	goto L36
L36:
	;
	v291 = v103 + v124
	v292 = v104 + v127
	v293 = v105 + v128
	v294 = v106 + v129
	v295 = v107 + v130
	v296 = v108 + v133
	v298 = v94 + int32(1)
	if v298 < v269 {
		v91 = v289
		v92 = v290
		v94 = v298
		v95 = v269
		v103 = v291
		v104 = v292
		v105 = v293
		v106 = v294
		v107 = v295
		v108 = v296
		goto L13
	} else {
		goto L37
	}
L37:
	;
	goto L14
L38:
	;
	v310 = *(*int64)(unsafe.Add(mBase, uint32(v30)+8))
	v311 = *(*int32)(unsafe.Add(mBase, uint32(v307)))
	if v311 <= int32(0) {
		v789 = v310
		goto L7
	} else {
		goto L39
	}
L39:
	;
	v316 = *(*int64)(unsafe.Add(mBase, uint32(v30)+56))
	v317 = *(*int64)(unsafe.Add(mBase, uint32(v30)+48))
	v318 = *(*int64)(unsafe.Add(mBase, uint32(v30)+40))
	v319 = *(*int64)(unsafe.Add(mBase, uint32(v30)+32))
	v320 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v30)+26)))
	v321 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v30)+24)))
	v322 = *(*int64)(unsafe.Add(mBase, uint32(v30)+16))
	v326 = v320
	v327 = v321
	v328 = v311
	v329 = v3
	v338 = v316
	v339 = v317
	v340 = v318
	v341 = v319
	v342 = v322
	v343 = v310
	goto L40
L40:
	;
	v352 = v307 + int32(8) + v329*int32(56)
	v353 = int32(*(*int16)(unsafe.Add(mBase, uint32(v352)+18)))
	v354 = base.I32_extend16_s(v326)
	v356 = int32(*(*int16)(unsafe.Add(mBase, uint32(v352)+16)))
	v357 = base.I32_extend16_s(v327)
	v359 = *(*int64)(unsafe.Add(mBase, uint32(v352)+48))
	v360 = *(*int64)(unsafe.Add(mBase, uint32(v352)+40))
	v361 = *(*int64)(unsafe.Add(mBase, uint32(v352)+32))
	v362 = *(*int64)(unsafe.Add(mBase, uint32(v352)+24))
	v363 = *(*int64)(unsafe.Add(mBase, uint32(v352)+8))
	v364 = *(*int64)(unsafe.Add(mBase, uint32(v352)))
	v365 = *(*int32)(unsafe.Add(mBase, uint32(l1)+60))
	if v365 != 0 {
		goto L42
	} else {
		goto L43
	}
L41:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v30)+56)) = v522
	*(*int64)(unsafe.Add(mBase, uint32(v30)+48)) = v523
	*(*int64)(unsafe.Add(mBase, uint32(v30)+40)) = v524
	*(*int64)(unsafe.Add(mBase, uint32(v30)+32)) = v525
	*(*uint16)(unsafe.Add(mBase, uint32(v30)+26)) = uint16(v520)
	*(*uint16)(unsafe.Add(mBase, uint32(v30)+24)) = uint16(v521)
	*(*int64)(unsafe.Add(mBase, uint32(v30)+16)) = v526
	v789 = v527
	goto L7
L42:
	;
	F_ExplainOpenWorker(m, v329, l1)
	mBase = m.M
	v367 = m.ExcPending
	if v367 != 0 {
		goto L18
	} else {
		goto L45
	}
L43:
	;
	v498 = v328
	goto L44
L44:
	;
	if v354 < v353 {
		goto L57
	} else {
		goto L58
	}
L45:
	;
	F_print_io_usage(m, l1, v352)
	mBase = m.M
	v369 = m.ExcPending
	if v369 != 0 {
		goto L18
	} else {
		goto L46
	}
L46:
	;
	v370 = *(*int32)(unsafe.Add(mBase, uint32(l1)+60))
	v371 = *(*int32)(unsafe.Add(mBase, uint32(v370)+12))
	F_ExplainSaveGroup(m, l1, v371+v329<<(uint(int32(2))%32))
	mBase = m.M
	v376 = m.ExcPending
	if v376 != 0 {
		goto L18
	} else {
		goto L47
	}
L47:
	;
	v377 = *(*int32)(unsafe.Add(mBase, uint32(l1)+20))
	if v377 == int32(0) {
		goto L48
	} else {
		goto L49
	}
L48:
	;
	v380 = *(*int32)(unsafe.Add(mBase, uint32(l1)))
	v381 = *(*int32)(unsafe.Add(mBase, uint32(v380)+4))
	if v381 <= int32(0) {
		goto L51
	} else {
		goto L52
	}
L49:
	;
	goto L50
L50:
	;
	v490 = *(*int32)(unsafe.Add(mBase, uint32(v370)+16))
	*(*int32)(unsafe.Add(mBase, uint32(l1))) = v490
	v492 = *(*int32)(unsafe.Add(mBase, uint32(v307)))
	v498 = v492
	goto L44
L51:
	;
	v459 = *(*int32)(unsafe.Add(mBase, uint32(l1)+24))
	*(*int32)(unsafe.Add(mBase, uint32(l1)+24)) = v459 - int32(1)
	goto L50
L52:
	;
	v386 = v380
	v389 = v381
	v390 = v380 + int32(4)
	goto L53
L53:
	;
	v413 = *(*int32)(unsafe.Add(mBase, uint32(v386)))
	v417 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v413+v389-int32(1)))))
	if v417 == int32(10) {
		goto L51
	} else {
		goto L55
	}
L54:
	;
	goto L51
L55:
	;
	v421 = v389 - int32(1)
	*(*int32)(unsafe.Add(mBase, uint32(v390))) = v421
	v424 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v413+v421))) = uint8(v424)
	v426 = *(*int32)(unsafe.Add(mBase, uint32(l1)))
	v429 = *(*int32)(unsafe.Add(mBase, uint32(v426)+4))
	if v424 < v429 {
		v386 = v426
		v389 = v429
		v390 = v426 + int32(4)
		goto L53
	} else {
		goto L56
	}
L56:
	;
	goto L54
L57:
	;
	v520 = v353
	goto L59
L58:
	;
	v520 = v354
	goto L59
L59:
	;
	if v357 < v356 {
		goto L60
	} else {
		goto L61
	}
L60:
	;
	v521 = v356
	goto L62
L61:
	;
	v521 = v357
	goto L62
L62:
	;
	v522 = v338 + v359
	v523 = v339 + v360
	v524 = v340 + v361
	v525 = v341 + v362
	v526 = v342 + v363
	v527 = v343 + v364
	v529 = v329 + int32(1)
	if v529 < v498 {
		v326 = v520
		v327 = v521
		v328 = v498
		v329 = v529
		v338 = v522
		v339 = v523
		v340 = v524
		v341 = v525
		v342 = v526
		v343 = v527
		goto L40
	} else {
		goto L63
	}
L63:
	;
	goto L41
L64:
	;
	v541 = *(*int64)(unsafe.Add(mBase, uint32(v30)+8))
	v542 = *(*int32)(unsafe.Add(mBase, uint32(v538)))
	if v542 <= int32(0) {
		v789 = v541
		goto L7
	} else {
		goto L65
	}
L65:
	;
	v547 = *(*int64)(unsafe.Add(mBase, uint32(v30)+56))
	v548 = *(*int64)(unsafe.Add(mBase, uint32(v30)+48))
	v549 = *(*int64)(unsafe.Add(mBase, uint32(v30)+40))
	v550 = *(*int64)(unsafe.Add(mBase, uint32(v30)+32))
	v551 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v30)+26)))
	v552 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v30)+24)))
	v553 = *(*int64)(unsafe.Add(mBase, uint32(v30)+16))
	v557 = v551
	v558 = v552
	v559 = v542
	v560 = v3
	v569 = v547
	v570 = v548
	v571 = v549
	v572 = v550
	v573 = v553
	v574 = v541
	goto L66
L66:
	;
	v583 = v538 + int32(8) + v560*int32(56)
	v584 = int32(*(*int16)(unsafe.Add(mBase, uint32(v583)+18)))
	v585 = base.I32_extend16_s(v557)
	v587 = int32(*(*int16)(unsafe.Add(mBase, uint32(v583)+16)))
	v588 = base.I32_extend16_s(v558)
	v590 = *(*int64)(unsafe.Add(mBase, uint32(v583)+48))
	v591 = *(*int64)(unsafe.Add(mBase, uint32(v583)+40))
	v592 = *(*int64)(unsafe.Add(mBase, uint32(v583)+32))
	v593 = *(*int64)(unsafe.Add(mBase, uint32(v583)+24))
	v594 = *(*int64)(unsafe.Add(mBase, uint32(v583)+8))
	v595 = *(*int64)(unsafe.Add(mBase, uint32(v583)))
	v596 = *(*int32)(unsafe.Add(mBase, uint32(l1)+60))
	if v596 != 0 {
		goto L68
	} else {
		goto L69
	}
L67:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v30)+56)) = v753
	*(*int64)(unsafe.Add(mBase, uint32(v30)+48)) = v754
	*(*int64)(unsafe.Add(mBase, uint32(v30)+40)) = v755
	*(*int64)(unsafe.Add(mBase, uint32(v30)+32)) = v756
	*(*uint16)(unsafe.Add(mBase, uint32(v30)+26)) = uint16(v751)
	*(*uint16)(unsafe.Add(mBase, uint32(v30)+24)) = uint16(v752)
	*(*int64)(unsafe.Add(mBase, uint32(v30)+16)) = v757
	v789 = v758
	goto L7
L68:
	;
	F_ExplainOpenWorker(m, v560, l1)
	mBase = m.M
	v598 = m.ExcPending
	if v598 != 0 {
		goto L18
	} else {
		goto L71
	}
L69:
	;
	v729 = v559
	goto L70
L70:
	;
	if v585 < v584 {
		goto L83
	} else {
		goto L84
	}
L71:
	;
	F_print_io_usage(m, l1, v583)
	mBase = m.M
	v600 = m.ExcPending
	if v600 != 0 {
		goto L18
	} else {
		goto L72
	}
L72:
	;
	v601 = *(*int32)(unsafe.Add(mBase, uint32(l1)+60))
	v602 = *(*int32)(unsafe.Add(mBase, uint32(v601)+12))
	F_ExplainSaveGroup(m, l1, v602+v560<<(uint(int32(2))%32))
	mBase = m.M
	v607 = m.ExcPending
	if v607 != 0 {
		goto L18
	} else {
		goto L73
	}
L73:
	;
	v608 = *(*int32)(unsafe.Add(mBase, uint32(l1)+20))
	if v608 == int32(0) {
		goto L74
	} else {
		goto L75
	}
L74:
	;
	v611 = *(*int32)(unsafe.Add(mBase, uint32(l1)))
	v612 = *(*int32)(unsafe.Add(mBase, uint32(v611)+4))
	if v612 <= int32(0) {
		goto L77
	} else {
		goto L78
	}
L75:
	;
	goto L76
L76:
	;
	v721 = *(*int32)(unsafe.Add(mBase, uint32(v601)+16))
	*(*int32)(unsafe.Add(mBase, uint32(l1))) = v721
	v723 = *(*int32)(unsafe.Add(mBase, uint32(v538)))
	v729 = v723
	goto L70
L77:
	;
	v690 = *(*int32)(unsafe.Add(mBase, uint32(l1)+24))
	*(*int32)(unsafe.Add(mBase, uint32(l1)+24)) = v690 - int32(1)
	goto L76
L78:
	;
	v617 = v611
	v620 = v612
	v621 = v611 + int32(4)
	goto L79
L79:
	;
	v644 = *(*int32)(unsafe.Add(mBase, uint32(v617)))
	v648 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v644+v620-int32(1)))))
	if v648 == int32(10) {
		goto L77
	} else {
		goto L81
	}
L80:
	;
	goto L77
L81:
	;
	v652 = v620 - int32(1)
	*(*int32)(unsafe.Add(mBase, uint32(v621))) = v652
	v655 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v644+v652))) = uint8(v655)
	v657 = *(*int32)(unsafe.Add(mBase, uint32(l1)))
	v660 = *(*int32)(unsafe.Add(mBase, uint32(v657)+4))
	if v655 < v660 {
		v617 = v657
		v620 = v660
		v621 = v657 + int32(4)
		goto L79
	} else {
		goto L82
	}
L82:
	;
	goto L80
L83:
	;
	v751 = v584
	goto L85
L84:
	;
	v751 = v585
	goto L85
L85:
	;
	if v588 < v587 {
		goto L86
	} else {
		goto L87
	}
L86:
	;
	v752 = v587
	goto L88
L87:
	;
	v752 = v588
	goto L88
L88:
	;
	v753 = v569 + v590
	v754 = v570 + v591
	v755 = v571 + v592
	v756 = v572 + v593
	v757 = v573 + v594
	v758 = v574 + v595
	v760 = v560 + int32(1)
	if v760 < v729 {
		v557 = v751
		v558 = v752
		v559 = v729
		v560 = v760
		v569 = v753
		v570 = v754
		v571 = v755
		v572 = v756
		v573 = v757
		v574 = v758
		goto L66
	} else {
		goto L89
	}
L89:
	;
	goto L67
L90:
	;
	goto L1
}
