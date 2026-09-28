package p0

import (
	base "github.com/shibukawa/pgmem/internal/aot/pgaot/base"
	"sync/atomic"
	"unsafe"
)

func F_cost_tuplesort(m *base.Module, l0 int32, l1 int32, l2 float64, l3 int32, l4 float64, l5 int32, l6 float64) {
	mBase := m.M
	_ = mBase
	var v12 float64
	_ = v12
	var v15 float64
	_ = v15
	var v17 float64
	_ = v17
	var v20 float64
	_ = v20
	var v23 int64
	_ = v23
	var v24 float64
	_ = v24
	var v31 float64
	_ = v31
	var v33 float64
	_ = v33
	var v37 int32
	_ = v37
	var v38 float64
	_ = v38
	var v40 float64
	_ = v40
	var v43 int32
	_ = v43
	var v45 int64
	_ = v45
	var v46 int32
	_ = v46
	var v49 int32
	_ = v49
	var v52 int32
	_ = v52
	var v55 float64
	_ = v55
	var v59 float64
	_ = v59
	var v61 float64
	_ = v61
	var v62 float64
	_ = v62
	var v64 float64
	_ = v64
	var v65 float64
	_ = v65
	var v69 float64
	_ = v69
	var v72 float64
	_ = v72
	var v76 float64
	_ = v76
	var v82 float64
	_ = v82
	var v83 float64
	_ = v83
	var v87 float64
	_ = v87
	var v91 float64
	_ = v91
	var v99 float64
	_ = v99
	var v102 float64
	_ = v102
	v12 = float64(2)
	if base.F64_lt(l2, v12) != 0 {
		v15 = v12
	} else {
		v15 = l2
	}
	v17 = *(*float64)(unsafe.Add(mBase, _c_F_cost_tuplesort[0]))
	v20 = base.F64_mul(v15, base.F64_add(base.F64_add(v17, v17), l4))
	v23 = base.I64_extend_i32_s(l5) << (uint(int64(10)) % 64)
	v24 = base.F64_convert_i64_s(v23)
	v31 = base.F64_convert_i32_u((l3+int32(7))&int32(-8) + int32(24))
	v33 = base.F64_mul(l2, v31)
	v37 = base.F64_lt(l6, v15) & base.F64_gt(l6, float64(0))
	if v37 != 0 {
		v38 = base.F64_mul(l6, v31)
	} else {
		v38 = v33
	}
	if base.F64_lt(v24, v38) != 0 {
		v40 = F_log(m, v15)
		mBase = m.M
		v43 = int32(6)
		v45 = base.I64_div_s(v23, int64(278528))
		v46 = base.I32_wrap_i64(v45)
		if v46 <= v43 {
			v49 = v43
		} else {
			v49 = v46
		}
		if int32(500) <= v49 {
			v52 = int32(500)
		} else {
			v52 = v49
		}
		v55 = base.F64_mul(base.F64_div(v40, float64(0.693147180559945)), v20)
		*(*float64)(unsafe.Add(mBase, uint32(l0))) = v55
		v59 = base.F64_ceil(base.F64_mul(v33, float64(0.0001220703125)))
		v61 = base.F64_div(v33, v24)
		v62 = base.F64_convert_i32_s(v52)
		if base.F64_gt(v61, v62) != 0 {
			v64 = F_log(m, v61)
			mBase = m.M
			v65 = F_log(m, v62)
			mBase = m.M
			v69 = base.F64_ceil(base.F64_div(v64, v65))
		} else {
			v69 = float64(1)
		}
		v72 = *(*float64)(unsafe.Add(mBase, _c_F_cost_tuplesort[1]))
		v76 = *(*float64)(unsafe.Add(mBase, _c_F_cost_tuplesort[2]))
		v99 = base.F64_add(base.F64_mul(base.F64_mul(base.F64_add(v59, v59), v69), base.F64_add(base.F64_mul(v72, float64(0.75)), base.F64_mul(v76, float64(0.25)))), v55)
	} else {
		if v37 != 0 {
			v82 = l6
		} else {
			v82 = v15
		}
		v83 = base.F64_add(v82, v82)
		if base.F64_gt(v15, v83)|base.F64_gt(v33, v24) != 0 {
			v87 = F_log(m, v83)
			mBase = m.M
			v99 = base.F64_mul(base.F64_div(v87, float64(0.693147180559945)), v20)
		} else {
			v91 = F_log(m, v15)
			mBase = m.M
			v99 = base.F64_mul(base.F64_div(v91, float64(0.693147180559945)), v20)
		}
	}
	*(*float64)(unsafe.Add(mBase, uint32(l0))) = v99
	v102 = *(*float64)(unsafe.Add(mBase, _c_F_cost_tuplesort[0]))
	*(*float64)(unsafe.Add(mBase, uint32(l1))) = base.F64_mul(v15, v102)
	return
}
func F_tuplesort_attach_shared(m *base.Module, l0 int32, l1 int32) {
	var v6 int32
	_ = v6
	F_SharedFileSetAttach(m, l0+int32(12), l1)
	v6 = m.ExcPending
	if v6 != 0 {
		return
	} else {
		return
	}
}
func F_tuplesort_end(m *base.Module, l0 int32) {
	mBase := m.M
	_ = mBase
	var v3 int32
	_ = v3
	var v4 int32
	_ = v4
	var v6 int32
	_ = v6
	F_tuplesort_free(m, l0)
	mBase = m.M
	v3 = m.ExcPending
	if v3 != 0 {
		return
	} else {
		v4 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
		F_MemoryContextDelete(m, v4)
		mBase = m.M
		v6 = m.ExcPending
		if v6 != 0 {
			return
		} else {
			return
		}
	}
}
func F_tuplesort_getdatum(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32, l5 int32) int32 {
	mBase := m.M
	_ = mBase
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
	var v23 int32
	_ = v23
	var v26 int32
	_ = v26
	var v31 int32
	_ = v31
	var v32 int32
	_ = v32
	var v35 int64
	_ = v35
	var v37 int32
	_ = v37
	var v40 int32
	_ = v40
	var v41 int64
	_ = v41
	var v42 int64
	_ = v42
	var v46 int32
	_ = v46
	var v47 int64
	_ = v47
	var v48 int32
	_ = v48
	var v50 int64
	_ = v50
	v11 = m.G0
	v13 = v11 - int32(32)
	m.G0 = v13
	v15 = int32(_a_F_tuplesort_getdatum_0)
	v16 = *(*int32)(unsafe.Add(mBase, _c_F_tuplesort_getdatum[0]))
	v18 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
	*(*int32)(unsafe.Add(mBase, _c_F_tuplesort_getdatum[0])) = v18
	v20 = *(*int32)(unsafe.Add(mBase, uint32(l0)+60))
	v23 = F_tuplesort_gettuple_common(m, l0, l1, v13+int32(8))
	mBase = m.M
	v26 = m.ExcPending
	if v26 != 0 {
		return int32(0)
	} else {
		*(*int32)(unsafe.Add(mBase, _c_F_tuplesort_getdatum[0])) = v16
		if v23 != 0 {
			if l5 == int32(0) {
			} else {
				v31 = *(*int32)(unsafe.Add(mBase, uint32(l0)+44))
				v32 = *(*int32)(unsafe.Add(mBase, uint32(v31)+24))
				if v32 == int32(0) {
				} else {
					v35 = *(*int64)(unsafe.Add(mBase, uint32(v13)+16))
					*(*int64)(unsafe.Add(mBase, uint32(l5))) = v35
				}
			}
			v37 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v13)+24)))
			if v37 == int32(0) {
				v40 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+56)))
				if v40 != 0 {
					v42 = int64(*(*uint32)(unsafe.Add(mBase, uint32(v13)+8)))
					if l2 == int32(0) {
						v50 = v42
						*(*int64)(unsafe.Add(mBase, uint32(l3))) = v50
						*(*uint8)(unsafe.Add(mBase, uint32(l4))) = uint8(v37)
						m.G0 = v13 + int32(32)
						return v23
					} else {
						v46 = *(*int32)(unsafe.Add(mBase, uint32(v20)+4))
						v47 = F_datumCopy(m, v42, int32(0), v46)
						mBase = m.M
						v48 = m.ExcPending
						if v48 != 0 {
							return int32(0)
						} else {
							v50 = v47
							*(*int64)(unsafe.Add(mBase, uint32(l3))) = v50
							*(*uint8)(unsafe.Add(mBase, uint32(l4))) = uint8(v37)
							m.G0 = v13 + int32(32)
							return v23
						}
					}
				} else {
					v41 = *(*int64)(unsafe.Add(mBase, uint32(v13)+16))
					v50 = v41
					*(*int64)(unsafe.Add(mBase, uint32(l3))) = v50
					*(*uint8)(unsafe.Add(mBase, uint32(l4))) = uint8(v37)
					m.G0 = v13 + int32(32)
					return v23
				}
			} else {
				v41 = *(*int64)(unsafe.Add(mBase, uint32(v13)+16))
				v50 = v41
				*(*int64)(unsafe.Add(mBase, uint32(l3))) = v50
				*(*uint8)(unsafe.Add(mBase, uint32(l4))) = uint8(v37)
				m.G0 = v13 + int32(32)
				return v23
			}
		} else {
			m.G0 = v13 + int32(32)
			return v23
		}
	}
}
func F_tuplesort_performsort(m *base.Module, l0 int32) {
	mBase := m.M
	_ = mBase
	var v18 int32
	_ = v18
	var v20 int32
	_ = v20
	var v22 int32
	_ = v22
	var v23 int32
	_ = v23
	var v25 int32
	_ = v25
	var v28 int32
	_ = v28
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
	var v48 int32
	_ = v48
	var v53 int32
	_ = v53
	var v55 int32
	_ = v55
	var v56 int32
	_ = v56
	var v60 int32
	_ = v60
	var v63 int32
	_ = v63
	var v68 int32
	_ = v68
	var v71 int32
	_ = v71
	var v72 int32
	_ = v72
	var v74 int32
	_ = v74
	var v75 int32
	_ = v75
	var v77 int32
	_ = v77
	var v78 int32
	_ = v78
	var v82 int32
	_ = v82
	var v86 int32
	_ = v86
	var v89 int32
	_ = v89
	var v92 int32
	_ = v92
	var v93 int32
	_ = v93
	var v97 int64
	_ = v97
	var v99 int32
	_ = v99
	var v103 int32
	_ = v103
	var v108 int32
	_ = v108
	var v111 int32
	_ = v111
	var v114 int32
	_ = v114
	var v115 int32
	_ = v115
	var v116 int32
	_ = v116
	var v122 int64
	_ = v122
	var v123 int32
	_ = v123
	var v124 int32
	_ = v124
	var v125 int32
	_ = v125
	var v128 int64
	_ = v128
	var v130 int64
	_ = v130
	var v134 int32
	_ = v134
	var v135 int32
	_ = v135
	var v140 int32
	_ = v140
	var v141 int32
	_ = v141
	var v150 int32
	_ = v150
	var v151 int32
	_ = v151
	var v160 int32
	_ = v160
	var v176 int32
	_ = v176
	var v181 int32
	_ = v181
	var v183 int32
	_ = v183
	var v186 int32
	_ = v186
	var v187 int32
	_ = v187
	var v188 int64
	_ = v188
	var v190 int32
	_ = v190
	var v192 int32
	_ = v192
	var v199 int64
	_ = v199
	var v211 int32
	_ = v211
	var v216 int32
	_ = v216
	var v221 int32
	_ = v221
	var v222 int32
	_ = v222
	var v224 int32
	_ = v224
	var v227 int32
	_ = v227
	var v229 int32
	_ = v229
	var v230 int32
	_ = v230
	var v232 int32
	_ = v232
	var v233 int32
	_ = v233
	var v234 int64
	_ = v234
	var v235 int32
	_ = v235
	var v236 int64
	_ = v236
	var v238 int32
	_ = v238
	var v242 int32
	_ = v242
	var v243 int32
	_ = v243
	var v245 int32
	_ = v245
	var v246 int32
	_ = v246
	var v247 int32
	_ = v247
	var v248 int32
	_ = v248
	var v251 int32
	_ = v251
	var v252 int32
	_ = v252
	var v254 int32
	_ = v254
	var v256 int32
	_ = v256
	var v257 int32
	_ = v257
	var v261 int32
	_ = v261
	var v262 int32
	_ = v262
	var v265 int32
	_ = v265
	var v266 int32
	_ = v266
	var v271 int32
	_ = v271
	var v273 int32
	_ = v273
	var v277 int32
	_ = v277
	var v292 int32
	_ = v292
	var v293 int32
	_ = v293
	var v296 int32
	_ = v296
	var v297 int32
	_ = v297
	var v302 int32
	_ = v302
	var v304 int32
	_ = v304
	var v306 int32
	_ = v306
	var v310 int32
	_ = v310
	var v311 int32
	_ = v311
	var v316 int32
	_ = v316
	var v319 int32
	_ = v319
	var v345 int32
	_ = v345
	var v349 int32
	_ = v349
	var v354 int32
	_ = v354
	var v370 int64
	_ = v370
	var v373 int64
	_ = v373
	var v376 int64
	_ = v376
	var v378 int64
	_ = v378
	var v379 int64
	_ = v379
	var v383 int64
	_ = v383
	var v385 int64
	_ = v385
	var v386 int64
	_ = v386
	var v392 int32
	_ = v392
	var v398 int32
	_ = v398
	var v420 int32
	_ = v420
	var v440 int32
	_ = v440
	var v444 int32
	_ = v444
	var v448 int32
	_ = v448
	var v464 int32
	_ = v464
	var v465 int64
	_ = v465
	var v467 int64
	_ = v467
	var v469 int64
	_ = v469
	var v472 int32
	_ = v472
	var v477 int32
	_ = v477
	var v479 int32
	_ = v479
	var v480 int32
	_ = v480
	var v481 int32
	_ = v481
	var v482 int32
	_ = v482
	var v483 int32
	_ = v483
	var v491 int32
	_ = v491
	var v495 int32
	_ = v495
	var v497 int32
	_ = v497
	var v507 int32
	_ = v507
	var v509 int32
	_ = v509
	var v515 int32
	_ = v515
	var v516 int32
	_ = v516
	var v517 int32
	_ = v517
	var v520 int32
	_ = v520
	var v521 int32
	_ = v521
	var v524 int32
	_ = v524
	var v525 int32
	_ = v525
	var v526 int32
	_ = v526
	var v527 int32
	_ = v527
	var v532 int32
	_ = v532
	var v533 int64
	_ = v533
	var v535 int64
	_ = v535
	var v537 int64
	_ = v537
	var v539 int32
	_ = v539
	var v540 int32
	_ = v540
	var v542 int32
	_ = v542
	var v545 int32
	_ = v545
	var v561 int32
	_ = v561
	var v563 int32
	_ = v563
	var v564 int64
	_ = v564
	var v566 int64
	_ = v566
	var v568 int64
	_ = v568
	var v570 int32
	_ = v570
	var v571 int32
	_ = v571
	var v574 int32
	_ = v574
	var v575 int64
	_ = v575
	var v577 int64
	_ = v577
	var v579 int64
	_ = v579
	var v581 int32
	_ = v581
	var v602 int32
	_ = v602
	var v605 int32
	_ = v605
	var v608 int32
	_ = v608
	var v609 int32
	_ = v609
	var v624 int32
	_ = v624
	var v625 int32
	_ = v625
	var v626 int32
	_ = v626
	var v628 int32
	_ = v628
	var v630 int32
	_ = v630
	var v635 int32
	_ = v635
	var v636 int32
	_ = v636
	var v655 int32
	_ = v655
	var v659 int32
	_ = v659
	var v666 int32
	_ = v666
	var v670 int32
	_ = v670
	var v675 int32
	_ = v675
	var v678 int32
	_ = v678
	var v680 int32
	_ = v680
	var v683 int32
	_ = v683
	var v702 int32
	_ = v702
	var v707 int32
	_ = v707
	var v710 int32
	_ = v710
	var v713 int32
	_ = v713
	var v714 int32
	_ = v714
	var v720 int32
	_ = v720
	var v721 int32
	_ = v721
	var v724 int32
	_ = v724
	var v725 int32
	_ = v725
	var v731 int32
	_ = v731
	var v735 int32
	_ = v735
	var v738 int32
	_ = v738
	var v739 int32
	_ = v739
	var v746 int32
	_ = v746
	var v750 int32
	_ = v750
	var v753 int32
	_ = v753
	var v764 int32
	_ = v764
	var v768 int32
	_ = v768
	var v773 int32
	_ = v773
	v18 = m.G0
	v20 = v18 + int32(-64)
	m.G0 = v20
	v22 = int32(_a_F_tuplesort_performsort_0)
	v23 = *(*int32)(unsafe.Add(mBase, _c_F_tuplesort_performsort[0]))
	v25 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
	*(*int32)(unsafe.Add(mBase, _c_F_tuplesort_performsort[0])) = v25
	v28 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_tuplesort_performsort[1])))
	if v28 != int32(1) {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	v55 = *(*int32)(unsafe.Add(mBase, uint32(l0)+64))
	switch v55 {
	case 0:
		goto L14
	case 1:
		goto L13
	case 2:
		goto L11
	default:
		goto L12
	}
L2:
	;
	v33 = F_errstart(m, int32(15), int32(0))
	mBase = m.M
	v34 = m.ExcPending
	if v34 != 0 {
		goto L3
	} else {
		goto L4
	}
L3:
	;
	return
L4:
	;
	if v33 == int32(0) {
		goto L1
	} else {
		goto L5
	}
L5:
	;
	v37 = *(*int32)(unsafe.Add(mBase, uint32(l0)+232))
	v40 = F_pg_rusage_show(m, l0+int32(256))
	mBase = m.M
	v41 = m.ExcPending
	if v41 != 0 {
		goto L3
	} else {
		goto L6
	}
L6:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v20)+36)) = v40
	*(*int32)(unsafe.Add(mBase, uint32(v20)+32)) = v37
	F_errmsg_internal(m, int32(_a_F_tuplesort_performsort_1), v18+int32(-32))
	mBase = m.M
	v48 = m.ExcPending
	if v48 != 0 {
		goto L3
	} else {
		goto L7
	}
L7:
	;
	F_errfinish(m, int32(_a_F_tuplesort_performsort_2), int32(1266), int32(_a_F_tuplesort_performsort_3))
	mBase = m.M
	v53 = m.ExcPending
	if v53 != 0 {
		goto L3
	} else {
		goto L8
	}
L8:
	;
	goto L1
L9:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v764 = m.ExcPending
	if v764 != 0 {
		goto L3
	} else {
		goto L132
	}
L10:
	;
	v702 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(l0)+228)) = uint8(v702)
	*(*int32)(unsafe.Add(mBase, uint32(l0)+224)) = v702
	v707 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_tuplesort_performsort[1])))
	if v707 != int32(1) {
		goto L118
	} else {
		goto L119
	}
L11:
	;
	F_dumptuples(m, l0, int32(1))
	mBase = m.M
	v678 = m.ExcPending
	if v678 != 0 {
		goto L3
	} else {
		goto L116
	}
L12:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v666 = m.ExcPending
	if v666 != 0 {
		goto L3
	} else {
		goto L113
	}
L13:
	;
	v444 = *(*int32)(unsafe.Add(mBase, uint32(l0)+136))
	if int32(2) <= v444 {
		goto L81
	} else {
		goto L82
	}
L14:
	;
	v56 = *(*int32)(unsafe.Add(mBase, uint32(l0)+236))
	if v56 == int32(0) {
		goto L16
	} else {
		goto L17
	}
L15:
	;
	*(*int64)(unsafe.Add(mBase, uint32(l0)+216)) = int64(0)
	v440 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(l0)+208)) = uint8(v440)
	*(*int32)(unsafe.Add(mBase, uint32(l0)+204)) = v440
	goto L10
L16:
	;
	F_tuplesort_sort_memtuples(m, l0)
	mBase = m.M
	v60 = m.ExcPending
	if v60 != 0 {
		goto L3
	} else {
		goto L19
	}
L17:
	;
	goto L18
L18:
	;
	v63 = *(*int32)(unsafe.Add(mBase, uint32(l0)+232))
	if v63 != int32(-1) {
		goto L20
	} else {
		goto L21
	}
L19:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+64)) = int32(3)
	goto L15
L20:
	;
	F_inittapes(m, l0, int32(0))
	mBase = m.M
	v68 = m.ExcPending
	if v68 != 0 {
		goto L3
	} else {
		goto L23
	}
L21:
	;
	goto L22
L22:
	;
	v108 = *(*int32)(unsafe.Add(mBase, uint32(l0)+240))
	v111 = base.AtomicRmwXchg32(m, v56, int32(0), int32(1))
	if v111 != 0 {
		goto L31
	} else {
		goto L32
	}
L23:
	;
	F_dumptuples(m, l0, int32(1))
	mBase = m.M
	v71 = m.ExcPending
	if v71 != 0 {
		goto L3
	} else {
		goto L24
	}
L24:
	;
	v72 = *(*int32)(unsafe.Add(mBase, uint32(l0)+196))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+200)) = v72
	v74 = *(*int32)(unsafe.Add(mBase, uint32(l0)+236))
	v75 = *(*int32)(unsafe.Add(mBase, uint32(l0)+132))
	F_pfree(m, v75)
	mBase = m.M
	v77 = m.ExcPending
	if v77 != 0 {
		goto L3
	} else {
		goto L25
	}
L25:
	;
	v78 = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(l0)+140)) = v78
	*(*int32)(unsafe.Add(mBase, uint32(l0)+132)) = v78
	v82 = *(*int32)(unsafe.Add(mBase, uint32(l0)+200))
	F_LogicalTapeFreeze(m, v82, v18+int32(-24))
	mBase = m.M
	v86 = m.ExcPending
	if v86 != 0 {
		goto L3
	} else {
		goto L26
	}
L26:
	;
	v89 = base.AtomicRmwXchg32(m, v74, int32(0), int32(1))
	if v89 != 0 {
		goto L27
	} else {
		goto L28
	}
L27:
	;
	F_s_lock(m, v74, int32(_a_F_tuplesort_performsort_4))
	mBase = m.M
	v92 = m.ExcPending
	if v92 != 0 {
		goto L3
	} else {
		goto L30
	}
L28:
	;
	goto L29
L29:
	;
	v93 = *(*int32)(unsafe.Add(mBase, uint32(l0)+232))
	v97 = *(*int64)(unsafe.Add(mBase, uint32(v20)+40))
	*(*int64)(unsafe.Add(mBase, uint32(v74+v93<<(uint(int32(3))%32))+72)) = v97
	v99 = *(*int32)(unsafe.Add(mBase, uint32(v74)+8))
	*(*int32)(unsafe.Add(mBase, uint32(v74)+8)) = v99 + int32(1)
	v103 = int32(0)
	atomic.StoreUint32((*uint32)(unsafe.Add(mBase, uint32(v74))), uint32(v103))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+64)) = int32(4)
	goto L15
L30:
	;
	goto L29
L31:
	;
	F_s_lock(m, v56, int32(_a_F_tuplesort_performsort_4))
	mBase = m.M
	v114 = m.ExcPending
	if v114 != 0 {
		goto L3
	} else {
		goto L34
	}
L32:
	;
	goto L33
L33:
	;
	v115 = *(*int32)(unsafe.Add(mBase, uint32(v56)+8))
	v116 = int32(0)
	atomic.StoreUint32((*uint32)(unsafe.Add(mBase, uint32(v56))), uint32(v116))
	if v115 != v108 {
		goto L9
	} else {
		goto L35
	}
L34:
	;
	goto L33
L35:
	;
	v122 = base.I64_extend_i32_s(v108) << (uint(int64(13)) % 64)
	v123 = *(*int32)(unsafe.Add(mBase, uint32(l0)+132))
	v124 = F_GetMemoryChunkSpace(m, v123)
	mBase = m.M
	v125 = m.ExcPending
	if v125 != 0 {
		goto L3
	} else {
		goto L36
	}
L36:
	;
	v128 = *(*int64)(unsafe.Add(mBase, uint32(l0)+96))
	if v122+base.I64_extend_i32_u(v124) < v128 {
		goto L37
	} else {
		goto L38
	}
L37:
	;
	v130 = *(*int64)(unsafe.Add(mBase, uint32(l0)+88))
	*(*int64)(unsafe.Add(mBase, uint32(l0)+88)) = v130 - v122
	goto L39
L38:
	;
	goto L39
L39:
	;
	F_PrepareTempTablespaces(m)
	mBase = m.M
	v134 = m.ExcPending
	if v134 != 0 {
		goto L3
	} else {
		goto L40
	}
L40:
	;
	v135 = int32(0)
	v140 = F_LogicalTapeSetCreate(m, v135, v56+int32(12), int32(-1))
	mBase = m.M
	v141 = m.ExcPending
	if v141 != 0 {
		goto L3
	} else {
		goto L41
	}
L41:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+180)) = int32(0)
	*(*int64)(unsafe.Add(mBase, uint32(l0)+172)) = int64(0)
	*(*int32)(unsafe.Add(mBase, uint32(l0)+168)) = v108
	*(*int32)(unsafe.Add(mBase, uint32(l0)+128)) = v140
	v150 = F_palloc0(m, v108<<(uint(int32(2))%32))
	mBase = m.M
	v151 = m.ExcPending
	if v151 != 0 {
		goto L3
	} else {
		goto L42
	}
L42:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+192)) = v108
	*(*int32)(unsafe.Add(mBase, uint32(l0)+188)) = v108
	*(*int32)(unsafe.Add(mBase, uint32(l0)+184)) = v150
	if int32(0) < v108 {
		goto L43
	} else {
		goto L44
	}
L43:
	;
	v160 = v135
	goto L46
L44:
	;
	goto L45
L45:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+64)) = int32(2)
	F_mergeruns(m, l0)
	mBase = m.M
	v420 = m.ExcPending
	if v420 != 0 {
		goto L3
	} else {
		goto L80
	}
L46:
	;
	v176 = *(*int32)(unsafe.Add(mBase, uint32(l0)+128))
	v181 = m.G0
	v183 = v181 - int32(1024)
	m.G0 = v183
	v186 = F_palloc(m, int32(72))
	mBase = m.M
	v187 = m.ExcPending
	if v187 != 0 {
		goto L3
	} else {
		goto L48
	}
L47:
	;
	goto L45
L48:
	;
	v188 = int64(-1)
	*(*int64)(unsafe.Add(mBase, uint32(v186)+8)) = v188
	v190 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v186)+6)) = uint8(v190)
	v192 = int32(1)
	*(*uint16)(unsafe.Add(mBase, uint32(v186)+4)) = uint16(v192)
	*(*int32)(unsafe.Add(mBase, uint32(v186))) = v176
	*(*int64)(unsafe.Add(mBase, uint32(v186)+16)) = v188
	*(*int64)(unsafe.Add(mBase, uint32(v186)+24)) = v188
	v199 = int64(0)
	*(*int64)(unsafe.Add(mBase, uint32(v186)+32)) = v199
	*(*int64)(unsafe.Add(mBase, uint32(v186)+40)) = v199
	*(*int64)(unsafe.Add(mBase, uint32(v186)+52)) = v199
	*(*int32)(unsafe.Add(mBase, uint32(v186)+48)) = int32(1073741823)
	*(*int64)(unsafe.Add(mBase, uint32(v186)+60)) = v199
	*(*int32)(unsafe.Add(mBase, uint32(v186)+68)) = v190
	v211 = base.I32_extend16_s(v160)
	if v190 <= v211 {
		goto L50
	} else {
		goto L51
	}
L49:
	;
	v229 = *(*int32)(unsafe.Add(mBase, uint32(v176)+4))
	v230 = int32(0)
	v232 = F_BufFileOpenFileSet(m, v229, v183, v230, v230)
	mBase = m.M
	v233 = m.ExcPending
	if v233 != 0 {
		goto L3
	} else {
		goto L53
	}
L50:
	;
	v221 = v211
	v222 = int32(0)
	goto L52
L51:
	;
	v216 = int32(45)
	*(*uint8)(unsafe.Add(mBase, uint32(v183))) = uint8(v216)
	v221 = int32(0) - v211
	v222 = int32(1)
	goto L52
L52:
	;
	v224 = F_pg_ultoa_n(m, v221, v183+v222)
	mBase = m.M
	v227 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v183+(v224+v222)))) = uint8(v227)
	goto L49
L53:
	;
	v234 = F_BufFileSize(m, v232)
	mBase = m.M
	v235 = m.ExcPending
	if v235 != 0 {
		goto L3
	} else {
		goto L54
	}
L54:
	;
	v236 = *(*int64)(unsafe.Add(mBase, uint32(v56+int32(72)+v160<<(uint(int32(3))%32))))
	*(*int64)(unsafe.Add(mBase, uint32(v186)+8)) = v236
	v238 = *(*int32)(unsafe.Add(mBase, uint32(v176)))
	if v238 == int32(0) {
		goto L56
	} else {
		goto L57
	}
L55:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v186)+32)) = v370
	v373 = int64(1073741823)
	if v373 <= v234 {
		goto L76
	} else {
		goto L77
	}
L56:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v176))) = v232
	v370 = int64(0)
	goto L55
L57:
	;
	goto L58
L58:
	;
	v242 = *(*int32)(unsafe.Add(mBase, uint32(v238)+20))
	v243 = *(*int32)(unsafe.Add(mBase, uint32(v232)+20))
	if v242 == v243 {
		goto L60
	} else {
		goto L61
	}
L59:
	;
	v370 = base.I64_extend_i32_s(v246) << (uint(int64(17)) % 64)
	goto L55
L60:
	;
	v245 = *(*int32)(unsafe.Add(mBase, uint32(v238)+4))
	v246 = *(*int32)(unsafe.Add(mBase, uint32(v238)))
	v247 = *(*int32)(unsafe.Add(mBase, uint32(v232)))
	v248 = v246 + v247
	v251 = F_repalloc(m, v245, v248<<(uint(int32(2))%32))
	mBase = m.M
	v252 = m.ExcPending
	if v252 != 0 {
		goto L3
	} else {
		goto L63
	}
L61:
	;
	goto L62
L62:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v345 = m.ExcPending
	if v345 != 0 {
		goto L3
	} else {
		goto L73
	}
L63:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v238)+4)) = v251
	v254 = *(*int32)(unsafe.Add(mBase, uint32(v238)))
	if v248 <= v254 {
		goto L64
	} else {
		goto L65
	}
L64:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v238))) = v248
	goto L59
L65:
	;
	v256 = int32(1)
	v257 = v254 + v256
	if (v248-v254)&v256 != 0 {
		goto L66
	} else {
		goto L67
	}
L66:
	;
	v261 = *(*int32)(unsafe.Add(mBase, uint32(v238)+4))
	v262 = int32(2)
	v265 = *(*int32)(unsafe.Add(mBase, uint32(v232)+4))
	v266 = *(*int32)(unsafe.Add(mBase, uint32(v238)))
	v271 = *(*int32)(unsafe.Add(mBase, uint32(v265+(v254-v266)<<(uint(v262)%32))))
	*(*int32)(unsafe.Add(mBase, uint32(v261+v254<<(uint(v262)%32)))) = v271
	v273 = v257
	goto L68
L67:
	;
	v273 = v254
	goto L68
L68:
	;
	if v257 == v248 {
		goto L64
	} else {
		goto L69
	}
L69:
	;
	v277 = v273
	goto L70
L70:
	;
	v292 = *(*int32)(unsafe.Add(mBase, uint32(v238)+4))
	v293 = int32(2)
	v296 = *(*int32)(unsafe.Add(mBase, uint32(v232)+4))
	v297 = *(*int32)(unsafe.Add(mBase, uint32(v238)))
	v302 = *(*int32)(unsafe.Add(mBase, uint32(v296+(v277-v297)<<(uint(v293)%32))))
	*(*int32)(unsafe.Add(mBase, uint32(v292+v277<<(uint(v293)%32)))) = v302
	v304 = *(*int32)(unsafe.Add(mBase, uint32(v238)+4))
	v306 = v277 + int32(1)
	v310 = *(*int32)(unsafe.Add(mBase, uint32(v232)+4))
	v311 = *(*int32)(unsafe.Add(mBase, uint32(v238)))
	v316 = *(*int32)(unsafe.Add(mBase, uint32(v310+(v306-v311)<<(uint(v293)%32))))
	*(*int32)(unsafe.Add(mBase, uint32(v304+v306<<(uint(v293)%32)))) = v316
	v319 = v277 + v293
	if v319 != v248 {
		v277 = v319
		goto L70
	} else {
		goto L72
	}
L71:
	;
	goto L64
L72:
	;
	goto L71
L73:
	;
	F_errmsg_internal(m, int32(_a_F_tuplesort_performsort_5), int32(0))
	mBase = m.M
	v349 = m.ExcPending
	if v349 != 0 {
		goto L3
	} else {
		goto L74
	}
L74:
	;
	F_errfinish(m, int32(_a_F_tuplesort_performsort_6), int32(912), int32(_a_F_tuplesort_performsort_7))
	mBase = m.M
	v354 = m.ExcPending
	if v354 != 0 {
		goto L3
	} else {
		goto L75
	}
L75:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L76:
	;
	v376 = v373
	goto L78
L77:
	;
	v376 = v234
	goto L78
L78:
	;
	*(*uint32)(unsafe.Add(mBase, uint32(v186)+48)) = uint32(v376)
	v378 = *(*int64)(unsafe.Add(mBase, uint32(v176)+32))
	v379 = *(*int64)(unsafe.Add(mBase, uint32(v176)+16))
	*(*int64)(unsafe.Add(mBase, uint32(v176)+32)) = v378 + (v370 - v379)
	v383 = *(*int64)(unsafe.Add(mBase, uint32(v186)+32))
	v385 = base.I64_div_s(v234, int64(8192))
	v386 = v383 + v385
	*(*int64)(unsafe.Add(mBase, uint32(v176)+24)) = v386
	*(*int64)(unsafe.Add(mBase, uint32(v176)+16)) = v386
	m.G0 = v183 + int32(1024)
	v392 = *(*int32)(unsafe.Add(mBase, uint32(l0)+184))
	*(*int32)(unsafe.Add(mBase, uint32(v392+v160<<(uint(int32(2))%32)))) = v186
	v398 = v160 + int32(1)
	if v398 != v108 {
		v160 = v398
		goto L46
	} else {
		goto L79
	}
L79:
	;
	goto L47
L80:
	;
	goto L15
L81:
	;
	v448 = v444
	goto L84
L82:
	;
	goto L83
L83:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+136)) = v444
	v602 = *(*int32)(unsafe.Add(mBase, uint32(l0)+40))
	if int32(0) < v602 {
		goto L107
	} else {
		goto L108
	}
L84:
	;
	v464 = *(*int32)(unsafe.Add(mBase, uint32(l0)+132))
	v465 = *(*int64)(unsafe.Add(mBase, uint32(v464)+16))
	*(*int64)(unsafe.Add(mBase, uint32(v20)+56)) = v465
	v467 = *(*int64)(unsafe.Add(mBase, uint32(v464)+8))
	*(*int64)(unsafe.Add(mBase, uint32(v20)+48)) = v467
	v469 = *(*int64)(unsafe.Add(mBase, uint32(v464)))
	*(*int64)(unsafe.Add(mBase, uint32(v20)+40)) = v469
	v472 = v448 - int32(1)
	*(*int32)(unsafe.Add(mBase, uint32(l0)+136)) = v472
	v477 = *(*int32)(unsafe.Add(mBase, _c_F_tuplesort_performsort[2]))
	if v477 != 0 {
		goto L86
	} else {
		goto L87
	}
L85:
	;
	goto L83
L86:
	;
	F_ProcessInterrupts(m)
	mBase = m.M
	v479 = m.ExcPending
	if v479 != 0 {
		goto L3
	} else {
		goto L89
	}
L87:
	;
	v481 = v472
	goto L88
L88:
	;
	v482 = v472*int32(24) + v464
	v483 = int32(0)
	if base.Ui32(v481) < base.Ui32(int32(2)) {
		v545 = v483
		goto L90
	} else {
		goto L91
	}
L89:
	;
	v480 = *(*int32)(unsafe.Add(mBase, uint32(l0)+136))
	v481 = v480
	goto L88
L90:
	;
	v561 = int32(24)
	v563 = v464 + v545*v561
	v564 = *(*int64)(unsafe.Add(mBase, uint32(v482)+16))
	*(*int64)(unsafe.Add(mBase, uint32(v563)+16)) = v564
	v566 = *(*int64)(unsafe.Add(mBase, uint32(v482)+8))
	*(*int64)(unsafe.Add(mBase, uint32(v563)+8)) = v566
	v568 = *(*int64)(unsafe.Add(mBase, uint32(v482)))
	*(*int64)(unsafe.Add(mBase, uint32(v563))) = v568
	v570 = *(*int32)(unsafe.Add(mBase, uint32(l0)+132))
	v571 = *(*int32)(unsafe.Add(mBase, uint32(l0)+136))
	v574 = v570 + v571*v561
	v575 = *(*int64)(unsafe.Add(mBase, uint32(v20)+56))
	*(*int64)(unsafe.Add(mBase, uint32(v574)+16)) = v575
	v577 = *(*int64)(unsafe.Add(mBase, uint32(v20)+48))
	*(*int64)(unsafe.Add(mBase, uint32(v574)+8)) = v577
	v579 = *(*int64)(unsafe.Add(mBase, uint32(v20)+40))
	*(*int64)(unsafe.Add(mBase, uint32(v574))) = v579
	v581 = *(*int32)(unsafe.Add(mBase, uint32(l0)+136))
	if int32(1) < v581 {
		v448 = v581
		goto L84
	} else {
		goto L106
	}
L91:
	;
	v491 = int32(1)
	v495 = v483
	v497 = v483
	goto L92
L92:
	;
	v507 = v497 + int32(2)
	if base.Ui32(v481) <= base.Ui32(v507) {
		goto L94
	} else {
		goto L95
	}
L93:
	;
	v545 = v521
	goto L90
L94:
	;
	v521 = v491
	goto L96
L95:
	;
	v509 = int32(24)
	v515 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v516 = m.T0[v515].(func(*base.Module, int32, int32, int32) int32)(m, v464+v491*v509, v464+v507*v509, l0)
	mBase = m.M
	v517 = m.ExcPending
	if v517 != 0 {
		goto L3
	} else {
		goto L97
	}
L96:
	;
	v524 = v464 + v521*int32(24)
	v525 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v526 = m.T0[v525].(func(*base.Module, int32, int32, int32) int32)(m, v482, v524, l0)
	mBase = m.M
	v527 = m.ExcPending
	if v527 != 0 {
		goto L3
	} else {
		goto L101
	}
L97:
	;
	if int32(0) < v516 {
		goto L98
	} else {
		goto L99
	}
L98:
	;
	v520 = v507
	goto L100
L99:
	;
	v520 = v491
	goto L100
L100:
	;
	v521 = v520
	goto L96
L101:
	;
	if v526 <= int32(0) {
		goto L102
	} else {
		goto L103
	}
L102:
	;
	v545 = v495
	goto L90
L103:
	;
	goto L104
L104:
	;
	v532 = v464 + v495*int32(24)
	v533 = *(*int64)(unsafe.Add(mBase, uint32(v524)+16))
	*(*int64)(unsafe.Add(mBase, uint32(v532)+16)) = v533
	v535 = *(*int64)(unsafe.Add(mBase, uint32(v524)+8))
	*(*int64)(unsafe.Add(mBase, uint32(v532)+8)) = v535
	v537 = *(*int64)(unsafe.Add(mBase, uint32(v524)))
	*(*int64)(unsafe.Add(mBase, uint32(v532))) = v537
	v539 = int32(1)
	v540 = v521 << (uint(v539) % 32)
	v542 = v540 | v539
	if base.Ui32(v542) < base.Ui32(v481) {
		v491 = v542
		v495 = v521
		v497 = v540
		goto L92
	} else {
		goto L105
	}
L105:
	;
	goto L93
L106:
	;
	goto L85
L107:
	;
	v605 = *(*int32)(unsafe.Add(mBase, uint32(l0)+44))
	v608 = v605
	v609 = int32(0)
	goto L110
L108:
	;
	goto L109
L109:
	;
	v655 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(l0)+208)) = uint8(v655)
	*(*int32)(unsafe.Add(mBase, uint32(l0)+204)) = v655
	v659 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(l0)+69)) = uint8(v659)
	*(*int32)(unsafe.Add(mBase, uint32(l0)+64)) = int32(3)
	goto L10
L110:
	;
	v624 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v608)+8)))
	v625 = int32(1)
	v626 = v624 ^ v625
	*(*uint8)(unsafe.Add(mBase, uint32(v608)+8)) = uint8(v626)
	v628 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v608)+9)))
	v630 = v628 ^ v625
	*(*uint8)(unsafe.Add(mBase, uint32(v608)+9)) = uint8(v630)
	v635 = v609 + v625
	v636 = *(*int32)(unsafe.Add(mBase, uint32(l0)+40))
	if v635 < v636 {
		v608 = v608 + int32(36)
		v609 = v635
		goto L110
	} else {
		goto L112
	}
L111:
	;
	goto L109
L112:
	;
	goto L111
L113:
	;
	F_errmsg_internal(m, int32(_a_F_tuplesort_performsort_8), int32(0))
	mBase = m.M
	v670 = m.ExcPending
	if v670 != 0 {
		goto L3
	} else {
		goto L114
	}
L114:
	;
	F_errfinish(m, int32(_a_F_tuplesort_performsort_2), int32(1341), int32(_a_F_tuplesort_performsort_3))
	mBase = m.M
	v675 = m.ExcPending
	if v675 != 0 {
		goto L3
	} else {
		goto L115
	}
L115:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L116:
	;
	F_mergeruns(m, l0)
	mBase = m.M
	v680 = m.ExcPending
	if v680 != 0 {
		goto L3
	} else {
		goto L117
	}
L117:
	;
	*(*int64)(unsafe.Add(mBase, uint32(l0)+216)) = int64(0)
	v683 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(l0)+208)) = uint8(v683)
	goto L10
L118:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_tuplesort_performsort[0])) = v23
	m.G0 = v20 - int32(-64)
	return
L119:
	;
	v710 = *(*int32)(unsafe.Add(mBase, uint32(l0)+64))
	v713 = F_errstart(m, int32(15), int32(0))
	mBase = m.M
	v714 = m.ExcPending
	if v714 != 0 {
		goto L3
	} else {
		goto L120
	}
L120:
	;
	if v710 == int32(5) {
		goto L122
	} else {
		goto L123
	}
L121:
	;
	F_errfinish(m, int32(_a_F_tuplesort_performsort_2), v750, int32(_a_F_tuplesort_performsort_3))
	mBase = m.M
	v753 = m.ExcPending
	if v753 != 0 {
		goto L3
	} else {
		goto L131
	}
L122:
	;
	if v713 == int32(0) {
		goto L118
	} else {
		goto L125
	}
L123:
	;
	goto L124
L124:
	;
	if v713 == int32(0) {
		goto L118
	} else {
		goto L128
	}
L125:
	;
	v720 = *(*int32)(unsafe.Add(mBase, uint32(l0)+232))
	v721 = *(*int32)(unsafe.Add(mBase, uint32(l0)+176))
	v724 = F_pg_rusage_show(m, l0+int32(256))
	mBase = m.M
	v725 = m.ExcPending
	if v725 != 0 {
		goto L3
	} else {
		goto L126
	}
L126:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v20)+8)) = v724
	*(*int32)(unsafe.Add(mBase, uint32(v20)+4)) = v721
	*(*int32)(unsafe.Add(mBase, uint32(v20))) = v720
	F_errmsg_internal(m, int32(_a_F_tuplesort_performsort_9), v20)
	mBase = m.M
	v731 = m.ExcPending
	if v731 != 0 {
		goto L3
	} else {
		goto L127
	}
L127:
	;
	v750 = int32(1350)
	goto L121
L128:
	;
	v735 = *(*int32)(unsafe.Add(mBase, uint32(l0)+232))
	v738 = F_pg_rusage_show(m, l0+int32(256))
	mBase = m.M
	v739 = m.ExcPending
	if v739 != 0 {
		goto L3
	} else {
		goto L129
	}
L129:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v20)+20)) = v738
	*(*int32)(unsafe.Add(mBase, uint32(v20)+16)) = v735
	F_errmsg_internal(m, int32(_a_F_tuplesort_performsort_10), v18+int32(-48))
	mBase = m.M
	v746 = m.ExcPending
	if v746 != 0 {
		goto L3
	} else {
		goto L130
	}
L130:
	;
	v750 = int32(1353)
	goto L121
L131:
	;
	goto L118
L132:
	;
	F_errmsg_internal(m, int32(_a_F_tuplesort_performsort_11), int32(0))
	mBase = m.M
	v768 = m.ExcPending
	if v768 != 0 {
		goto L3
	} else {
		goto L133
	}
L133:
	;
	F_errfinish(m, int32(_a_F_tuplesort_performsort_2), int32(3395), int32(_a_F_tuplesort_performsort_12))
	mBase = m.M
	v773 = m.ExcPending
	if v773 != 0 {
		goto L3
	} else {
		goto L134
	}
L134:
	;
	base.Wasm_trap_unreachable()
	for {
	}
}
func F_tuplesort_putgintuple(m *base.Module, l0 int32, l1 int32, l2 int32) {
	mBase := m.M
	_ = mBase
	var v7 int32
	_ = v7
	var v9 int32
	_ = v9
	var v11 int32
	_ = v11
	var v12 int32
	_ = v12
	var v14 int32
	_ = v14
	var v16 int32
	_ = v16
	var v17 int32
	_ = v17
	var v19 int32
	_ = v19
	var v24 int32
	_ = v24
	var v27 int32
	_ = v27
	var v28 int32
	_ = v28
	var v33 int32
	_ = v33
	var v36 int32
	_ = v36
	var v37 int32
	_ = v37
	var v41 int32
	_ = v41
	var v43 int32
	_ = v43
	v7 = m.G0
	v9 = v7 - int32(32)
	m.G0 = v9
	v11 = int32(_a_F_tuplesort_putgintuple_0)
	v12 = *(*int32)(unsafe.Add(mBase, _c_F_tuplesort_putgintuple[0]))
	v14 = *(*int32)(unsafe.Add(mBase, uint32(l0)+32))
	*(*int32)(unsafe.Add(mBase, _c_F_tuplesort_putgintuple[0])) = v14
	v16 = F_palloc(m, l2)
	mBase = m.M
	v17 = m.ExcPending
	if v17 != 0 {
		return
	} else {
		if l2 != 0 {
			base.MemoryCopy(m, v16, l1, l2)
		} else {
		}
		v19 = int32(0)
		*(*uint8)(unsafe.Add(mBase, uint32(v9)+24)) = uint8(v19)
		*(*int64)(unsafe.Add(mBase, uint32(v9)+16)) = int64(0)
		*(*int32)(unsafe.Add(mBase, uint32(v9)+8)) = v16
		v24 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+52)))
		if v24&int32(2) != 0 {
			v27 = F_GetMemoryChunkSpace(m, v16)
			mBase = m.M
			v28 = m.ExcPending
			if v28 != 0 {
				return
			} else {
				v33 = v27
				v36 = *(*int32)(unsafe.Add(mBase, uint32(l0)+44))
				if v36 != 0 {
					v37 = *(*int32)(unsafe.Add(mBase, uint32(v36)+24))
					v41 = base.B2i32(v37 != int32(0))
				} else {
					v41 = int32(0)
				}
				F_tuplesort_puttuple_common(m, l0, v9+int32(8), v41, v33)
				mBase = m.M
				v43 = m.ExcPending
				if v43 != 0 {
					return
				} else {
					*(*int32)(unsafe.Add(mBase, _c_F_tuplesort_putgintuple[0])) = v12
					m.G0 = v9 + int32(32)
					return
				}
			}
		} else {
			v33 = (l2 + int32(7)) & int32(-8)
			v36 = *(*int32)(unsafe.Add(mBase, uint32(l0)+44))
			if v36 != 0 {
				v37 = *(*int32)(unsafe.Add(mBase, uint32(v36)+24))
				v41 = base.B2i32(v37 != int32(0))
			} else {
				v41 = int32(0)
			}
			F_tuplesort_puttuple_common(m, l0, v9+int32(8), v41, v33)
			mBase = m.M
			v43 = m.ExcPending
			if v43 != 0 {
				return
			} else {
				*(*int32)(unsafe.Add(mBase, _c_F_tuplesort_putgintuple[0])) = v12
				m.G0 = v9 + int32(32)
				return
			}
		}
	}
}
func F_tuplesort_puttupleslot(m *base.Module, l0 int32, l1 int32) {
	mBase := m.M
	_ = mBase
	var v7 int32
	_ = v7
	var v9 int32
	_ = v9
	var v11 int32
	_ = v11
	var v12 int32
	_ = v12
	var v14 int32
	_ = v14
	var v16 int32
	_ = v16
	var v18 int32
	_ = v18
	var v19 int32
	_ = v19
	var v20 int32
	_ = v20
	var v21 int32
	_ = v21
	var v23 int32
	_ = v23
	var v24 int32
	_ = v24
	var v32 int32
	_ = v32
	var v33 int32
	_ = v33
	var v36 int64
	_ = v36
	var v37 int32
	_ = v37
	var v39 int32
	_ = v39
	var v44 int32
	_ = v44
	var v49 int32
	_ = v49
	var v50 int32
	_ = v50
	var v51 int32
	_ = v51
	var v54 int32
	_ = v54
	var v57 int32
	_ = v57
	var v58 int32
	_ = v58
	var v63 int32
	_ = v63
	v7 = m.G0
	v9 = v7 - int32(48)
	m.G0 = v9
	v11 = int32(_a_F_tuplesort_puttupleslot_0)
	v12 = *(*int32)(unsafe.Add(mBase, _c_F_tuplesort_puttupleslot[0]))
	v14 = *(*int32)(unsafe.Add(mBase, uint32(l0)+32))
	*(*int32)(unsafe.Add(mBase, _c_F_tuplesort_puttupleslot[0])) = v14
	v16 = *(*int32)(unsafe.Add(mBase, uint32(l0)+60))
	v18 = *(*int32)(unsafe.Add(mBase, uint32(l1)+8))
	v19 = *(*int32)(unsafe.Add(mBase, uint32(v18)+48))
	v20 = m.T0[v19].(func(*base.Module, int32, int32) int32)(m, l1, int32(0))
	mBase = m.M
	v21 = m.ExcPending
	if v21 != 0 {
		return
	} else {
		*(*int32)(unsafe.Add(mBase, uint32(v9)+24)) = v20
		v23 = *(*int32)(unsafe.Add(mBase, uint32(v20)))
		v24 = int32(8)
		*(*int32)(unsafe.Add(mBase, uint32(v9)+20)) = v20 - v24
		*(*int32)(unsafe.Add(mBase, uint32(v9)+4)) = v23 + v24
		v32 = *(*int32)(unsafe.Add(mBase, uint32(l0)+44))
		v33 = int32(*(*int16)(unsafe.Add(mBase, uint32(v32)+10)))
		v36 = F_heap_getattr_1(m, v9+int32(4), v33, v16, v9+int32(40))
		mBase = m.M
		v37 = m.ExcPending
		if v37 != 0 {
			return
		} else {
			*(*int64)(unsafe.Add(mBase, uint32(v9)+32)) = v36
			v39 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+52)))
			if v39&int32(2) == int32(0) {
				v44 = *(*int32)(unsafe.Add(mBase, uint32(v20)))
				v51 = (v44 + int32(7)) & int32(-8)
				v54 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v9)+40)))
				v57 = *(*int32)(unsafe.Add(mBase, uint32(l0)+44))
				v58 = *(*int32)(unsafe.Add(mBase, uint32(v57)+24))
				F_tuplesort_puttuple_common(m, l0, v9+int32(24), (v54^int32(1))&base.B2i32(v58 != int32(0)), v51)
				mBase = m.M
				v63 = m.ExcPending
				if v63 != 0 {
					return
				} else {
					*(*int32)(unsafe.Add(mBase, _c_F_tuplesort_puttupleslot[0])) = v12
					m.G0 = v9 + int32(48)
					return
				}
			} else {
				v49 = F_GetMemoryChunkSpace(m, v20)
				mBase = m.M
				v50 = m.ExcPending
				if v50 != 0 {
					return
				} else {
					v51 = v49
					v54 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v9)+40)))
					v57 = *(*int32)(unsafe.Add(mBase, uint32(l0)+44))
					v58 = *(*int32)(unsafe.Add(mBase, uint32(v57)+24))
					F_tuplesort_puttuple_common(m, l0, v9+int32(24), (v54^int32(1))&base.B2i32(v58 != int32(0)), v51)
					mBase = m.M
					v63 = m.ExcPending
					if v63 != 0 {
						return
					} else {
						*(*int32)(unsafe.Add(mBase, _c_F_tuplesort_puttupleslot[0])) = v12
						m.G0 = v9 + int32(48)
						return
					}
				}
			}
		}
	}
}
func F_tuplesort_reset(m *base.Module, l0 int32) {
	mBase := m.M
	_ = mBase
	var v5 int32
	_ = v5
	var v8 int64
	_ = v8
	var v9 int64
	_ = v9
	var v11 int32
	_ = v11
	var v12 int64
	_ = v12
	var v13 int64
	_ = v13
	var v16 int64
	_ = v16
	var v17 int32
	_ = v17
	var v18 int32
	_ = v18
	var v21 int64
	_ = v21
	var v22 int32
	_ = v22
	var v26 int64
	_ = v26
	var v28 int64
	_ = v28
	var v29 int32
	_ = v29
	var v32 int32
	_ = v32
	var v37 int32
	_ = v37
	var v39 int32
	_ = v39
	var v40 int32
	_ = v40
	v5 = *(*int32)(unsafe.Add(mBase, uint32(l0)+128))
	if v5 == int32(0) {
		v8 = *(*int64)(unsafe.Add(mBase, uint32(l0)+96))
		v9 = *(*int64)(unsafe.Add(mBase, uint32(l0)+88))
		v11 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+120)))
		v21 = v8 - v9
		v22 = v11
		if v22 != base.B2i32(v5 != int32(0)) {
		} else {
			v26 = *(*int64)(unsafe.Add(mBase, uint32(l0)+112))
			if v21 <= v26 {
			} else {
				v28 = v21
				v29 = v22
				*(*uint8)(unsafe.Add(mBase, uint32(l0)+120)) = uint8(v29)
				*(*int64)(unsafe.Add(mBase, uint32(l0)+112)) = v28
				v32 = *(*int32)(unsafe.Add(mBase, uint32(l0)+64))
				*(*int32)(unsafe.Add(mBase, uint32(l0)+124)) = v32
			}
		}
	} else {
		v12 = *(*int64)(unsafe.Add(mBase, uint32(v5)+24))
		v13 = *(*int64)(unsafe.Add(mBase, uint32(v5)+32))
		v16 = (v12 - v13) << (uint(int64(13)) % 64)
		v17 = int32(1)
		v18 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+120)))
		if v18 != v17 {
			v28 = v16
			v29 = v17
			*(*uint8)(unsafe.Add(mBase, uint32(l0)+120)) = uint8(v29)
			*(*int64)(unsafe.Add(mBase, uint32(l0)+112)) = v28
			v32 = *(*int32)(unsafe.Add(mBase, uint32(l0)+64))
			*(*int32)(unsafe.Add(mBase, uint32(l0)+124)) = v32
		} else {
			v21 = v16
			v22 = v17
			if v22 != base.B2i32(v5 != int32(0)) {
			} else {
				v26 = *(*int64)(unsafe.Add(mBase, uint32(l0)+112))
				if v21 <= v26 {
				} else {
					v28 = v21
					v29 = v22
					*(*uint8)(unsafe.Add(mBase, uint32(l0)+120)) = uint8(v29)
					*(*int64)(unsafe.Add(mBase, uint32(l0)+112)) = v28
					v32 = *(*int32)(unsafe.Add(mBase, uint32(l0)+64))
					*(*int32)(unsafe.Add(mBase, uint32(l0)+124)) = v32
				}
			}
		}
	}
	F_tuplesort_free(m, l0)
	mBase = m.M
	v37 = m.ExcPending
	if v37 != 0 {
		return
	} else {
		F_tuplesort_begin_batch(m, l0)
		mBase = m.M
		v39 = m.ExcPending
		if v39 != 0 {
			return
		} else {
			v40 = int32(0)
			*(*int32)(unsafe.Add(mBase, uint32(l0)+156)) = v40
			*(*int64)(unsafe.Add(mBase, uint32(l0)+148)) = int64(0)
			*(*int32)(unsafe.Add(mBase, uint32(l0)+164)) = v40
			return
		}
	}
}
