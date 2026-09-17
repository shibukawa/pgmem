package p0

import (
	base "github.com/shibukawa/pgmem/internal/aot/pgaot/base"
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
	var v10 int32
	_ = v10
	var v12 int32
	_ = v12
	var v14 int32
	_ = v14
	var v15 int32
	_ = v15
	var v17 int32
	_ = v17
	var v19 int32
	_ = v19
	var v20 int32
	_ = v20
	var v23 int32
	_ = v23
	var v28 int32
	_ = v28
	var v29 int32
	_ = v29
	var v32 int32
	_ = v32
	var v34 int32
	_ = v34
	var v37 int32
	_ = v37
	var v38 int32
	_ = v38
	var v39 int32
	_ = v39
	var v43 int32
	_ = v43
	var v44 int32
	_ = v44
	var v45 int32
	_ = v45
	var v47 int32
	_ = v47
	v10 = m.G0
	v12 = v10 - int32(16)
	m.G0 = v12
	v14 = int32(_a_F_tuplesort_getdatum_0)
	v15 = *(*int32)(unsafe.Add(mBase, _c_F_tuplesort_getdatum[0]))
	v17 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
	*(*int32)(unsafe.Add(mBase, _c_F_tuplesort_getdatum[0])) = v17
	v19 = *(*int32)(unsafe.Add(mBase, uint32(l0)+60))
	v20 = F_tuplesort_gettuple_common(m, l0, l1, v12)
	mBase = m.M
	v23 = m.ExcPending
	if v23 != 0 {
		return int32(0)
	} else {
		*(*int32)(unsafe.Add(mBase, _c_F_tuplesort_getdatum[0])) = v15
		if v20 != 0 {
			if l5 == int32(0) {
			} else {
				v28 = *(*int32)(unsafe.Add(mBase, uint32(l0)+44))
				v29 = *(*int32)(unsafe.Add(mBase, uint32(v28)+24))
				if v29 == int32(0) {
				} else {
					v32 = *(*int32)(unsafe.Add(mBase, uint32(v12)+4))
					*(*int32)(unsafe.Add(mBase, uint32(l5))) = v32
				}
			}
			v34 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v12)+8)))
			if v34 == int32(0) {
				v37 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+56)))
				if v37 != 0 {
					v39 = *(*int32)(unsafe.Add(mBase, uint32(v12)))
					if l2 == int32(0) {
						v47 = v39
						*(*int32)(unsafe.Add(mBase, uint32(l3))) = v47
						*(*uint8)(unsafe.Add(mBase, uint32(l4))) = uint8(v34)
						m.G0 = v12 + int32(16)
						return v20
					} else {
						v43 = *(*int32)(unsafe.Add(mBase, uint32(v19)+4))
						v44 = F_datumCopy(m, v39, int32(0), v43)
						mBase = m.M
						v45 = m.ExcPending
						if v45 != 0 {
							return int32(0)
						} else {
							v47 = v44
							*(*int32)(unsafe.Add(mBase, uint32(l3))) = v47
							*(*uint8)(unsafe.Add(mBase, uint32(l4))) = uint8(v34)
							m.G0 = v12 + int32(16)
							return v20
						}
					}
				} else {
					v38 = *(*int32)(unsafe.Add(mBase, uint32(v12)+4))
					v47 = v38
					*(*int32)(unsafe.Add(mBase, uint32(l3))) = v47
					*(*uint8)(unsafe.Add(mBase, uint32(l4))) = uint8(v34)
					m.G0 = v12 + int32(16)
					return v20
				}
			} else {
				v38 = *(*int32)(unsafe.Add(mBase, uint32(v12)+4))
				v47 = v38
				*(*int32)(unsafe.Add(mBase, uint32(l3))) = v47
				*(*uint8)(unsafe.Add(mBase, uint32(l4))) = uint8(v34)
				m.G0 = v12 + int32(16)
				return v20
			}
		} else {
			m.G0 = v12 + int32(16)
			return v20
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
	var v87 int32
	_ = v87
	var v94 int32
	_ = v94
	var v95 int32
	_ = v95
	var v99 int64
	_ = v99
	var v103 int32
	_ = v103
	var v109 int32
	_ = v109
	var v110 int32
	_ = v110
	var v117 int32
	_ = v117
	var v120 int32
	_ = v120
	var v124 int64
	_ = v124
	var v125 int32
	_ = v125
	var v126 int32
	_ = v126
	var v127 int32
	_ = v127
	var v130 int64
	_ = v130
	var v132 int64
	_ = v132
	var v136 int32
	_ = v136
	var v137 int32
	_ = v137
	var v142 int32
	_ = v142
	var v143 int32
	_ = v143
	var v152 int32
	_ = v152
	var v153 int32
	_ = v153
	var v162 int32
	_ = v162
	var v178 int32
	_ = v178
	var v183 int32
	_ = v183
	var v185 int32
	_ = v185
	var v188 int32
	_ = v188
	var v189 int32
	_ = v189
	var v190 int64
	_ = v190
	var v192 int32
	_ = v192
	var v194 int32
	_ = v194
	var v201 int64
	_ = v201
	var v213 int32
	_ = v213
	var v218 int32
	_ = v218
	var v223 int32
	_ = v223
	var v224 int32
	_ = v224
	var v226 int32
	_ = v226
	var v229 int32
	_ = v229
	var v231 int32
	_ = v231
	var v232 int32
	_ = v232
	var v234 int32
	_ = v234
	var v235 int32
	_ = v235
	var v236 int64
	_ = v236
	var v237 int32
	_ = v237
	var v238 int64
	_ = v238
	var v240 int32
	_ = v240
	var v244 int32
	_ = v244
	var v245 int32
	_ = v245
	var v247 int32
	_ = v247
	var v248 int32
	_ = v248
	var v249 int32
	_ = v249
	var v250 int32
	_ = v250
	var v253 int32
	_ = v253
	var v254 int32
	_ = v254
	var v256 int32
	_ = v256
	var v258 int32
	_ = v258
	var v259 int32
	_ = v259
	var v263 int32
	_ = v263
	var v264 int32
	_ = v264
	var v267 int32
	_ = v267
	var v268 int32
	_ = v268
	var v273 int32
	_ = v273
	var v275 int32
	_ = v275
	var v279 int32
	_ = v279
	var v294 int32
	_ = v294
	var v295 int32
	_ = v295
	var v298 int32
	_ = v298
	var v299 int32
	_ = v299
	var v304 int32
	_ = v304
	var v306 int32
	_ = v306
	var v308 int32
	_ = v308
	var v312 int32
	_ = v312
	var v313 int32
	_ = v313
	var v318 int32
	_ = v318
	var v321 int32
	_ = v321
	var v347 int32
	_ = v347
	var v351 int32
	_ = v351
	var v356 int32
	_ = v356
	var v372 int64
	_ = v372
	var v375 int64
	_ = v375
	var v378 int64
	_ = v378
	var v380 int64
	_ = v380
	var v381 int64
	_ = v381
	var v385 int64
	_ = v385
	var v387 int64
	_ = v387
	var v388 int64
	_ = v388
	var v394 int32
	_ = v394
	var v400 int32
	_ = v400
	var v422 int32
	_ = v422
	var v442 int32
	_ = v442
	var v446 int32
	_ = v446
	var v450 int32
	_ = v450
	var v466 int32
	_ = v466
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
	var v494 int32
	_ = v494
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
	var v537 int32
	_ = v537
	var v538 int32
	_ = v538
	var v540 int32
	_ = v540
	var v543 int32
	_ = v543
	var v559 int32
	_ = v559
	var v561 int32
	_ = v561
	var v562 int64
	_ = v562
	var v564 int64
	_ = v564
	var v566 int32
	_ = v566
	var v567 int32
	_ = v567
	var v570 int32
	_ = v570
	var v571 int64
	_ = v571
	var v573 int64
	_ = v573
	var v575 int32
	_ = v575
	var v596 int32
	_ = v596
	var v599 int32
	_ = v599
	var v602 int32
	_ = v602
	var v603 int32
	_ = v603
	var v618 int32
	_ = v618
	var v619 int32
	_ = v619
	var v620 int32
	_ = v620
	var v622 int32
	_ = v622
	var v624 int32
	_ = v624
	var v629 int32
	_ = v629
	var v630 int32
	_ = v630
	var v649 int32
	_ = v649
	var v653 int32
	_ = v653
	var v660 int32
	_ = v660
	var v664 int32
	_ = v664
	var v669 int32
	_ = v669
	var v672 int32
	_ = v672
	var v674 int32
	_ = v674
	var v677 int32
	_ = v677
	var v696 int32
	_ = v696
	var v701 int32
	_ = v701
	var v704 int32
	_ = v704
	var v707 int32
	_ = v707
	var v708 int32
	_ = v708
	var v714 int32
	_ = v714
	var v715 int32
	_ = v715
	var v718 int32
	_ = v718
	var v719 int32
	_ = v719
	var v725 int32
	_ = v725
	var v729 int32
	_ = v729
	var v732 int32
	_ = v732
	var v733 int32
	_ = v733
	var v740 int32
	_ = v740
	var v744 int32
	_ = v744
	var v747 int32
	_ = v747
	var v758 int32
	_ = v758
	var v762 int32
	_ = v762
	var v767 int32
	_ = v767
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
	F_errfinish(m, int32(_a_F_tuplesort_performsort_2), int32(1369), int32(_a_F_tuplesort_performsort_3))
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
	v758 = m.ExcPending
	if v758 != 0 {
		goto L3
	} else {
		goto L132
	}
L10:
	;
	v696 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(l0)+228)) = uint8(v696)
	*(*int32)(unsafe.Add(mBase, uint32(l0)+224)) = v696
	v701 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_tuplesort_performsort[1])))
	if v701 != int32(1) {
		goto L118
	} else {
		goto L119
	}
L11:
	;
	F_dumptuples(m, l0, int32(1))
	mBase = m.M
	v672 = m.ExcPending
	if v672 != 0 {
		goto L3
	} else {
		goto L116
	}
L12:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v660 = m.ExcPending
	if v660 != 0 {
		goto L3
	} else {
		goto L113
	}
L13:
	;
	v446 = *(*int32)(unsafe.Add(mBase, uint32(l0)+136))
	if int32(2) <= v446 {
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
	v442 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(l0)+208)) = uint8(v442)
	*(*int32)(unsafe.Add(mBase, uint32(l0)+204)) = v442
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
	v109 = *(*int32)(unsafe.Add(mBase, uint32(v56)))
	v110 = *(*int32)(unsafe.Add(mBase, uint32(l0)+240))
	*(*int32)(unsafe.Add(mBase, uint32(v56))) = int32(1)
	if v109 != 0 {
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
	F_LogicalTapeFreeze(m, v82, v18+int32(-16))
	mBase = m.M
	v86 = m.ExcPending
	if v86 != 0 {
		goto L3
	} else {
		goto L26
	}
L26:
	;
	v87 = *(*int32)(unsafe.Add(mBase, uint32(v74)))
	*(*int32)(unsafe.Add(mBase, uint32(v74))) = int32(1)
	if v87 != 0 {
		goto L27
	} else {
		goto L28
	}
L27:
	;
	F_s_lock(m, v74, int32(_a_F_tuplesort_performsort_2), int32(3034), int32(_a_F_tuplesort_performsort_4))
	mBase = m.M
	v94 = m.ExcPending
	if v94 != 0 {
		goto L3
	} else {
		goto L30
	}
L28:
	;
	goto L29
L29:
	;
	v95 = *(*int32)(unsafe.Add(mBase, uint32(l0)+232))
	v99 = *(*int64)(unsafe.Add(mBase, uint32(v20)+48))
	*(*int64)(unsafe.Add(mBase, uint32(v74+v95<<(uint(int32(3))%32))+72)) = v99
	*(*int32)(unsafe.Add(mBase, uint32(v74))) = int32(0)
	v103 = *(*int32)(unsafe.Add(mBase, uint32(v74)+8))
	*(*int32)(unsafe.Add(mBase, uint32(v74)+8)) = v103 + int32(1)
	*(*int32)(unsafe.Add(mBase, uint32(l0)+64)) = int32(4)
	goto L15
L30:
	;
	goto L29
L31:
	;
	F_s_lock(m, v56, int32(_a_F_tuplesort_performsort_2), int32(3079), int32(_a_F_tuplesort_performsort_5))
	mBase = m.M
	v117 = m.ExcPending
	if v117 != 0 {
		goto L3
	} else {
		goto L34
	}
L32:
	;
	goto L33
L33:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v56))) = int32(0)
	v120 = *(*int32)(unsafe.Add(mBase, uint32(v56)+8))
	if v110 != v120 {
		goto L9
	} else {
		goto L35
	}
L34:
	;
	goto L33
L35:
	;
	v124 = base.I64_extend_i32_s(v110) << (uint(int64(13)) % 64)
	v125 = *(*int32)(unsafe.Add(mBase, uint32(l0)+132))
	v126 = F_GetMemoryChunkSpace(m, v125)
	mBase = m.M
	v127 = m.ExcPending
	if v127 != 0 {
		goto L3
	} else {
		goto L36
	}
L36:
	;
	v130 = *(*int64)(unsafe.Add(mBase, uint32(l0)+96))
	if v124+base.I64_extend_i32_u(v126) < v130 {
		goto L37
	} else {
		goto L38
	}
L37:
	;
	v132 = *(*int64)(unsafe.Add(mBase, uint32(l0)+88))
	*(*int64)(unsafe.Add(mBase, uint32(l0)+88)) = v132 - v124
	goto L39
L38:
	;
	goto L39
L39:
	;
	F_PrepareTempTablespaces(m)
	mBase = m.M
	v136 = m.ExcPending
	if v136 != 0 {
		goto L3
	} else {
		goto L40
	}
L40:
	;
	v137 = int32(0)
	v142 = F_LogicalTapeSetCreate(m, v137, v56+int32(12), int32(-1))
	mBase = m.M
	v143 = m.ExcPending
	if v143 != 0 {
		goto L3
	} else {
		goto L41
	}
L41:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+180)) = int32(0)
	*(*int64)(unsafe.Add(mBase, uint32(l0)+172)) = int64(0)
	*(*int32)(unsafe.Add(mBase, uint32(l0)+168)) = v110
	*(*int32)(unsafe.Add(mBase, uint32(l0)+128)) = v142
	v152 = F_palloc0(m, v110<<(uint(int32(2))%32))
	mBase = m.M
	v153 = m.ExcPending
	if v153 != 0 {
		goto L3
	} else {
		goto L42
	}
L42:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+192)) = v110
	*(*int32)(unsafe.Add(mBase, uint32(l0)+188)) = v110
	*(*int32)(unsafe.Add(mBase, uint32(l0)+184)) = v152
	if int32(0) < v110 {
		goto L43
	} else {
		goto L44
	}
L43:
	;
	v162 = v137
	goto L46
L44:
	;
	goto L45
L45:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+64)) = int32(2)
	F_mergeruns(m, l0)
	mBase = m.M
	v422 = m.ExcPending
	if v422 != 0 {
		goto L3
	} else {
		goto L80
	}
L46:
	;
	v178 = *(*int32)(unsafe.Add(mBase, uint32(l0)+128))
	v183 = m.G0
	v185 = v183 - int32(1024)
	m.G0 = v185
	v188 = F_palloc(m, int32(72))
	mBase = m.M
	v189 = m.ExcPending
	if v189 != 0 {
		goto L3
	} else {
		goto L48
	}
L47:
	;
	goto L45
L48:
	;
	v190 = int64(-1)
	*(*int64)(unsafe.Add(mBase, uint32(v188)+8)) = v190
	v192 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v188)+6)) = uint8(v192)
	v194 = int32(1)
	*(*uint16)(unsafe.Add(mBase, uint32(v188)+4)) = uint16(v194)
	*(*int32)(unsafe.Add(mBase, uint32(v188))) = v178
	*(*int64)(unsafe.Add(mBase, uint32(v188)+16)) = v190
	*(*int64)(unsafe.Add(mBase, uint32(v188)+24)) = v190
	v201 = int64(0)
	*(*int64)(unsafe.Add(mBase, uint32(v188)+32)) = v201
	*(*int64)(unsafe.Add(mBase, uint32(v188)+40)) = v201
	*(*int64)(unsafe.Add(mBase, uint32(v188)+52)) = v201
	*(*int32)(unsafe.Add(mBase, uint32(v188)+48)) = int32(1073741823)
	*(*int64)(unsafe.Add(mBase, uint32(v188)+60)) = v201
	*(*int32)(unsafe.Add(mBase, uint32(v188)+68)) = v192
	v213 = base.I32_extend16_s(v162)
	if v192 <= v213 {
		goto L50
	} else {
		goto L51
	}
L49:
	;
	v231 = *(*int32)(unsafe.Add(mBase, uint32(v178)+4))
	v232 = int32(0)
	v234 = F_BufFileOpenFileSet(m, v231, v185, v232, v232)
	mBase = m.M
	v235 = m.ExcPending
	if v235 != 0 {
		goto L3
	} else {
		goto L53
	}
L50:
	;
	v223 = v213
	v224 = int32(0)
	goto L52
L51:
	;
	v218 = int32(45)
	*(*uint8)(unsafe.Add(mBase, uint32(v185))) = uint8(v218)
	v223 = int32(0) - v213
	v224 = int32(1)
	goto L52
L52:
	;
	v226 = F_pg_ultoa_n(m, v223, v185+v224)
	mBase = m.M
	v229 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v185+(v226+v224)))) = uint8(v229)
	goto L49
L53:
	;
	v236 = F_BufFileSize(m, v234)
	mBase = m.M
	v237 = m.ExcPending
	if v237 != 0 {
		goto L3
	} else {
		goto L54
	}
L54:
	;
	v238 = *(*int64)(unsafe.Add(mBase, uint32(v56+int32(72)+v162<<(uint(int32(3))%32))))
	*(*int64)(unsafe.Add(mBase, uint32(v188)+8)) = v238
	v240 = *(*int32)(unsafe.Add(mBase, uint32(v178)))
	if v240 == int32(0) {
		goto L56
	} else {
		goto L57
	}
L55:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v188)+32)) = v372
	v375 = int64(1073741823)
	if v375 <= v236 {
		goto L76
	} else {
		goto L77
	}
L56:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v178))) = v234
	v372 = int64(0)
	goto L55
L57:
	;
	goto L58
L58:
	;
	v244 = *(*int32)(unsafe.Add(mBase, uint32(v240)+20))
	v245 = *(*int32)(unsafe.Add(mBase, uint32(v234)+20))
	if v244 == v245 {
		goto L60
	} else {
		goto L61
	}
L59:
	;
	v372 = base.I64_extend_i32_s(v248) << (uint(int64(17)) % 64)
	goto L55
L60:
	;
	v247 = *(*int32)(unsafe.Add(mBase, uint32(v240)+4))
	v248 = *(*int32)(unsafe.Add(mBase, uint32(v240)))
	v249 = *(*int32)(unsafe.Add(mBase, uint32(v234)))
	v250 = v248 + v249
	v253 = F_repalloc(m, v247, v250<<(uint(int32(2))%32))
	mBase = m.M
	v254 = m.ExcPending
	if v254 != 0 {
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
	v347 = m.ExcPending
	if v347 != 0 {
		goto L3
	} else {
		goto L73
	}
L63:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v240)+4)) = v253
	v256 = *(*int32)(unsafe.Add(mBase, uint32(v240)))
	if v250 <= v256 {
		goto L64
	} else {
		goto L65
	}
L64:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v240))) = v250
	goto L59
L65:
	;
	v258 = int32(1)
	v259 = v256 + v258
	if (v250-v256)&v258 != 0 {
		goto L66
	} else {
		goto L67
	}
L66:
	;
	v263 = *(*int32)(unsafe.Add(mBase, uint32(v240)+4))
	v264 = int32(2)
	v267 = *(*int32)(unsafe.Add(mBase, uint32(v234)+4))
	v268 = *(*int32)(unsafe.Add(mBase, uint32(v240)))
	v273 = *(*int32)(unsafe.Add(mBase, uint32(v267+(v256-v268)<<(uint(v264)%32))))
	*(*int32)(unsafe.Add(mBase, uint32(v263+v256<<(uint(v264)%32)))) = v273
	v275 = v259
	goto L68
L67:
	;
	v275 = v256
	goto L68
L68:
	;
	if v259 == v250 {
		goto L64
	} else {
		goto L69
	}
L69:
	;
	v279 = v275
	goto L70
L70:
	;
	v294 = *(*int32)(unsafe.Add(mBase, uint32(v240)+4))
	v295 = int32(2)
	v298 = *(*int32)(unsafe.Add(mBase, uint32(v234)+4))
	v299 = *(*int32)(unsafe.Add(mBase, uint32(v240)))
	v304 = *(*int32)(unsafe.Add(mBase, uint32(v298+(v279-v299)<<(uint(v295)%32))))
	*(*int32)(unsafe.Add(mBase, uint32(v294+v279<<(uint(v295)%32)))) = v304
	v306 = *(*int32)(unsafe.Add(mBase, uint32(v240)+4))
	v308 = v279 + int32(1)
	v312 = *(*int32)(unsafe.Add(mBase, uint32(v234)+4))
	v313 = *(*int32)(unsafe.Add(mBase, uint32(v240)))
	v318 = *(*int32)(unsafe.Add(mBase, uint32(v312+(v308-v313)<<(uint(v295)%32))))
	*(*int32)(unsafe.Add(mBase, uint32(v306+v308<<(uint(v295)%32)))) = v318
	v321 = v279 + v295
	if v321 != v250 {
		v279 = v321
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
	F_errmsg_internal(m, int32(_a_F_tuplesort_performsort_6), int32(0))
	mBase = m.M
	v351 = m.ExcPending
	if v351 != 0 {
		goto L3
	} else {
		goto L74
	}
L74:
	;
	F_errfinish(m, int32(_a_F_tuplesort_performsort_7), int32(912), int32(_a_F_tuplesort_performsort_8))
	mBase = m.M
	v356 = m.ExcPending
	if v356 != 0 {
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
	v378 = v375
	goto L78
L77:
	;
	v378 = v236
	goto L78
L78:
	;
	*(*uint32)(unsafe.Add(mBase, uint32(v188)+48)) = uint32(v378)
	v380 = *(*int64)(unsafe.Add(mBase, uint32(v178)+32))
	v381 = *(*int64)(unsafe.Add(mBase, uint32(v178)+16))
	*(*int64)(unsafe.Add(mBase, uint32(v178)+32)) = v380 + (v372 - v381)
	v385 = *(*int64)(unsafe.Add(mBase, uint32(v188)+32))
	v387 = base.I64_div_s(v236, int64(8192))
	v388 = v385 + v387
	*(*int64)(unsafe.Add(mBase, uint32(v178)+24)) = v388
	*(*int64)(unsafe.Add(mBase, uint32(v178)+16)) = v388
	m.G0 = v185 + int32(1024)
	v394 = *(*int32)(unsafe.Add(mBase, uint32(l0)+184))
	*(*int32)(unsafe.Add(mBase, uint32(v394+v162<<(uint(int32(2))%32)))) = v188
	v400 = v162 + int32(1)
	if v400 != v110 {
		v162 = v400
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
	v450 = v446
	goto L84
L82:
	;
	goto L83
L83:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+136)) = v446
	v596 = *(*int32)(unsafe.Add(mBase, uint32(l0)+40))
	if int32(0) < v596 {
		goto L107
	} else {
		goto L108
	}
L84:
	;
	v466 = *(*int32)(unsafe.Add(mBase, uint32(l0)+132))
	v467 = *(*int64)(unsafe.Add(mBase, uint32(v466)+8))
	*(*int64)(unsafe.Add(mBase, uint32(v20)+56)) = v467
	v469 = *(*int64)(unsafe.Add(mBase, uint32(v466)))
	*(*int64)(unsafe.Add(mBase, uint32(v20)+48)) = v469
	v472 = v450 - int32(1)
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
	v482 = v472<<(uint(int32(4))%32) + v466
	v483 = int32(0)
	if base.Ui32(v481) < base.Ui32(int32(2)) {
		v543 = v483
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
	v559 = int32(4)
	v561 = v466 + v543<<(uint(v559)%32)
	v562 = *(*int64)(unsafe.Add(mBase, uint32(v482)+8))
	*(*int64)(unsafe.Add(mBase, uint32(v561)+8)) = v562
	v564 = *(*int64)(unsafe.Add(mBase, uint32(v482)))
	*(*int64)(unsafe.Add(mBase, uint32(v561))) = v564
	v566 = *(*int32)(unsafe.Add(mBase, uint32(l0)+132))
	v567 = *(*int32)(unsafe.Add(mBase, uint32(l0)+136))
	v570 = v566 + v567<<(uint(v559)%32)
	v571 = *(*int64)(unsafe.Add(mBase, uint32(v20)+56))
	*(*int64)(unsafe.Add(mBase, uint32(v570)+8)) = v571
	v573 = *(*int64)(unsafe.Add(mBase, uint32(v20)+48))
	*(*int64)(unsafe.Add(mBase, uint32(v570))) = v573
	v575 = *(*int32)(unsafe.Add(mBase, uint32(l0)+136))
	if int32(1) < v575 {
		v450 = v575
		goto L84
	} else {
		goto L106
	}
L91:
	;
	v491 = v483
	v494 = int32(1)
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
	v543 = v521
	goto L90
L94:
	;
	v521 = v494
	goto L96
L95:
	;
	v509 = int32(4)
	v515 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v516 = m.T0[v515].(func(*base.Module, int32, int32, int32) int32)(m, v466+v494<<(uint(v509)%32), v466+v507<<(uint(v509)%32), l0)
	mBase = m.M
	v517 = m.ExcPending
	if v517 != 0 {
		goto L3
	} else {
		goto L97
	}
L96:
	;
	v524 = v466 + v521<<(uint(int32(4))%32)
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
	v520 = v494
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
	v543 = v491
	goto L90
L103:
	;
	goto L104
L104:
	;
	v532 = v466 + v491<<(uint(int32(4))%32)
	v533 = *(*int64)(unsafe.Add(mBase, uint32(v524)+8))
	*(*int64)(unsafe.Add(mBase, uint32(v532)+8)) = v533
	v535 = *(*int64)(unsafe.Add(mBase, uint32(v524)))
	*(*int64)(unsafe.Add(mBase, uint32(v532))) = v535
	v537 = int32(1)
	v538 = v521 << (uint(v537) % 32)
	v540 = v538 | v537
	if base.Ui32(v540) < base.Ui32(v481) {
		v491 = v521
		v494 = v540
		v497 = v538
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
	v599 = *(*int32)(unsafe.Add(mBase, uint32(l0)+44))
	v602 = v599
	v603 = int32(0)
	goto L110
L108:
	;
	goto L109
L109:
	;
	v649 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(l0)+208)) = uint8(v649)
	*(*int32)(unsafe.Add(mBase, uint32(l0)+204)) = v649
	v653 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(l0)+69)) = uint8(v653)
	*(*int32)(unsafe.Add(mBase, uint32(l0)+64)) = int32(3)
	goto L10
L110:
	;
	v618 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v602)+8)))
	v619 = int32(1)
	v620 = v618 ^ v619
	*(*uint8)(unsafe.Add(mBase, uint32(v602)+8)) = uint8(v620)
	v622 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v602)+9)))
	v624 = v622 ^ v619
	*(*uint8)(unsafe.Add(mBase, uint32(v602)+9)) = uint8(v624)
	v629 = v603 + v619
	v630 = *(*int32)(unsafe.Add(mBase, uint32(l0)+40))
	if v629 < v630 {
		v602 = v602 + int32(36)
		v603 = v629
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
	F_errmsg_internal(m, int32(_a_F_tuplesort_performsort_9), int32(0))
	mBase = m.M
	v664 = m.ExcPending
	if v664 != 0 {
		goto L3
	} else {
		goto L114
	}
L114:
	;
	F_errfinish(m, int32(_a_F_tuplesort_performsort_2), int32(1444), int32(_a_F_tuplesort_performsort_3))
	mBase = m.M
	v669 = m.ExcPending
	if v669 != 0 {
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
	v674 = m.ExcPending
	if v674 != 0 {
		goto L3
	} else {
		goto L117
	}
L117:
	;
	*(*int64)(unsafe.Add(mBase, uint32(l0)+216)) = int64(0)
	v677 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(l0)+208)) = uint8(v677)
	goto L10
L118:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_tuplesort_performsort[0])) = v23
	m.G0 = v20 - int32(-64)
	return
L119:
	;
	v704 = *(*int32)(unsafe.Add(mBase, uint32(l0)+64))
	v707 = F_errstart(m, int32(15), int32(0))
	mBase = m.M
	v708 = m.ExcPending
	if v708 != 0 {
		goto L3
	} else {
		goto L120
	}
L120:
	;
	if v704 == int32(5) {
		goto L122
	} else {
		goto L123
	}
L121:
	;
	F_errfinish(m, int32(_a_F_tuplesort_performsort_2), v744, int32(_a_F_tuplesort_performsort_3))
	mBase = m.M
	v747 = m.ExcPending
	if v747 != 0 {
		goto L3
	} else {
		goto L131
	}
L122:
	;
	if v707 == int32(0) {
		goto L118
	} else {
		goto L125
	}
L123:
	;
	goto L124
L124:
	;
	if v707 == int32(0) {
		goto L118
	} else {
		goto L128
	}
L125:
	;
	v714 = *(*int32)(unsafe.Add(mBase, uint32(l0)+232))
	v715 = *(*int32)(unsafe.Add(mBase, uint32(l0)+176))
	v718 = F_pg_rusage_show(m, l0+int32(256))
	mBase = m.M
	v719 = m.ExcPending
	if v719 != 0 {
		goto L3
	} else {
		goto L126
	}
L126:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v20)+8)) = v718
	*(*int32)(unsafe.Add(mBase, uint32(v20)+4)) = v715
	*(*int32)(unsafe.Add(mBase, uint32(v20))) = v714
	F_errmsg_internal(m, int32(_a_F_tuplesort_performsort_10), v20)
	mBase = m.M
	v725 = m.ExcPending
	if v725 != 0 {
		goto L3
	} else {
		goto L127
	}
L127:
	;
	v744 = int32(1453)
	goto L121
L128:
	;
	v729 = *(*int32)(unsafe.Add(mBase, uint32(l0)+232))
	v732 = F_pg_rusage_show(m, l0+int32(256))
	mBase = m.M
	v733 = m.ExcPending
	if v733 != 0 {
		goto L3
	} else {
		goto L129
	}
L129:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v20)+20)) = v732
	*(*int32)(unsafe.Add(mBase, uint32(v20)+16)) = v729
	F_errmsg_internal(m, int32(_a_F_tuplesort_performsort_11), v18+int32(-48))
	mBase = m.M
	v740 = m.ExcPending
	if v740 != 0 {
		goto L3
	} else {
		goto L130
	}
L130:
	;
	v744 = int32(1456)
	goto L121
L131:
	;
	goto L118
L132:
	;
	F_errmsg_internal(m, int32(_a_F_tuplesort_performsort_12), int32(0))
	mBase = m.M
	v762 = m.ExcPending
	if v762 != 0 {
		goto L3
	} else {
		goto L133
	}
L133:
	;
	F_errfinish(m, int32(_a_F_tuplesort_performsort_2), int32(3084), int32(_a_F_tuplesort_performsort_5))
	mBase = m.M
	v767 = m.ExcPending
	if v767 != 0 {
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
	var v34 int32
	_ = v34
	var v35 int32
	_ = v35
	var v39 int32
	_ = v39
	var v41 int32
	_ = v41
	v7 = m.G0
	v9 = v7 - int32(16)
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
		*(*uint8)(unsafe.Add(mBase, uint32(v9)+8)) = uint8(v19)
		*(*int32)(unsafe.Add(mBase, uint32(v9)+4)) = v19
		*(*int32)(unsafe.Add(mBase, uint32(v9))) = v16
		v24 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+52)))
		if v24&int32(2) != 0 {
			v27 = F_GetMemoryChunkSpace(m, v16)
			mBase = m.M
			v28 = m.ExcPending
			if v28 != 0 {
				return
			} else {
				v33 = v27
				v34 = *(*int32)(unsafe.Add(mBase, uint32(l0)+44))
				if v34 != 0 {
					v35 = *(*int32)(unsafe.Add(mBase, uint32(v34)+24))
					v39 = base.B2i32(v35 != int32(0))
				} else {
					v39 = int32(0)
				}
				F_tuplesort_puttuple_common(m, l0, v9, v39, v33)
				mBase = m.M
				v41 = m.ExcPending
				if v41 != 0 {
					return
				} else {
					*(*int32)(unsafe.Add(mBase, _c_F_tuplesort_putgintuple[0])) = v12
					m.G0 = v9 + int32(16)
					return
				}
			}
		} else {
			v33 = (l2 + int32(7)) & int32(-8)
			v34 = *(*int32)(unsafe.Add(mBase, uint32(l0)+44))
			if v34 != 0 {
				v35 = *(*int32)(unsafe.Add(mBase, uint32(v34)+24))
				v39 = base.B2i32(v35 != int32(0))
			} else {
				v39 = int32(0)
			}
			F_tuplesort_puttuple_common(m, l0, v9, v39, v33)
			mBase = m.M
			v41 = m.ExcPending
			if v41 != 0 {
				return
			} else {
				*(*int32)(unsafe.Add(mBase, _c_F_tuplesort_putgintuple[0])) = v12
				m.G0 = v9 + int32(16)
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
	var v36 int32
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
		*(*int32)(unsafe.Add(mBase, uint32(v9)+32)) = v20
		v23 = *(*int32)(unsafe.Add(mBase, uint32(v20)))
		v24 = int32(8)
		*(*int32)(unsafe.Add(mBase, uint32(v9)+28)) = v20 - v24
		*(*int32)(unsafe.Add(mBase, uint32(v9)+12)) = v23 + v24
		v32 = *(*int32)(unsafe.Add(mBase, uint32(l0)+44))
		v33 = int32(*(*int16)(unsafe.Add(mBase, uint32(v32)+10)))
		v36 = F_heap_getattr_1(m, v9+int32(12), v33, v16, v9+int32(40))
		mBase = m.M
		v37 = m.ExcPending
		if v37 != 0 {
			return
		} else {
			*(*int32)(unsafe.Add(mBase, uint32(v9)+36)) = v36
			v39 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+52)))
			if v39&int32(2) == int32(0) {
				v44 = *(*int32)(unsafe.Add(mBase, uint32(v20)))
				v51 = (v44 + int32(7)) & int32(-8)
				v54 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v9)+40)))
				v57 = *(*int32)(unsafe.Add(mBase, uint32(l0)+44))
				v58 = *(*int32)(unsafe.Add(mBase, uint32(v57)+24))
				F_tuplesort_puttuple_common(m, l0, v9+int32(32), (v54^int32(1))&base.B2i32(v58 != int32(0)), v51)
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
					F_tuplesort_puttuple_common(m, l0, v9+int32(32), (v54^int32(1))&base.B2i32(v58 != int32(0)), v51)
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
