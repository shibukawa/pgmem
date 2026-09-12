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
	var v27 float64
	_ = v27
	var v29 float64
	_ = v29
	var v33 int32
	_ = v33
	var v34 float64
	_ = v34
	var v37 int64
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
	var v60 float64
	_ = v60
	var v62 float64
	_ = v62
	var v63 float64
	_ = v63
	var v65 float64
	_ = v65
	var v66 float64
	_ = v66
	var v69 float64
	_ = v69
	var v72 float64
	_ = v72
	var v76 float64
	_ = v76
	var v83 float64
	_ = v83
	var v84 float64
	_ = v84
	var v87 float64
	_ = v87
	var v91 float64
	_ = v91
	var v101 float64
	_ = v101
	var v104 float64
	_ = v104
	v12 = float64(2)
	if base.F64_lt(l2, v12) != 0 {
		v15 = v12
	} else {
		v15 = l2
	}
	v17 = *(*float64)(unsafe.Add(mBase, _consts[384]))
	v20 = base.F64_mul(v15, base.F64_add(base.F64_add(v17, v17), l4))
	v27 = base.F64_convert_i32_u((l3+int32(7))&int32(-8) + int32(24))
	v29 = base.F64_mul(l2, v27)
	v33 = base.F64_lt(l6, v15) & base.F64_gt(l6, float64(0))
	if v33 != 0 {
		v34 = base.F64_mul(l6, v27)
	} else {
		v34 = v29
	}
	v37 = base.I64_extend_i32_s(l5) << (uint(int64(10)) % 64)
	v38 = base.F64_convert_i64_s(v37)
	if base.F64_gt(v34, v38) != 0 {
		v40 = F_log(m, v15)
		mBase = m.M
		v43 = int32(6)
		v45 = base.I64_div_s(v37, int64(278528))
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
		v60 = base.F64_ceil(base.F64_mul(v29, float64(0.0001220703125)))
		v62 = base.F64_div(v29, v38)
		v63 = base.F64_convert_i32_s(v52)
		if base.F64_gt(v62, v63) != 0 {
			v65 = F_log(m, v62)
			mBase = m.M
			v66 = F_log(m, v63)
			mBase = m.M
			v69 = base.F64_ceil(base.F64_div(v65, v66))
		} else {
			v69 = float64(1)
		}
		v72 = *(*float64)(unsafe.Add(mBase, _consts[386]))
		v76 = *(*float64)(unsafe.Add(mBase, _consts[387]))
		v101 = base.F64_add(base.F64_mul(base.F64_mul(base.F64_add(v60, v60), v69), base.F64_add(base.F64_mul(v72, float64(0.75)), base.F64_mul(v76, float64(0.25)))), v55)
	} else {
		if v33 != 0 {
			v83 = l6
		} else {
			v83 = v15
		}
		v84 = base.F64_add(v83, v83)
		if base.F64_gt(v29, v38)|base.F64_gt(v15, v84) != 0 {
			v87 = F_log(m, v84)
			mBase = m.M
			v101 = base.F64_mul(base.F64_div(v87, float64(0.693147180559945)), v20)
		} else {
			v91 = F_log(m, v15)
			mBase = m.M
			v101 = base.F64_mul(base.F64_div(v91, float64(0.693147180559945)), v20)
		}
	}
	*(*float64)(unsafe.Add(mBase, uint32(l0))) = v101
	v104 = *(*float64)(unsafe.Add(mBase, _consts[384]))
	*(*float64)(unsafe.Add(mBase, uint32(l1))) = base.F64_mul(v15, v104)
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
	var v40 int32
	_ = v40
	var v44 int32
	_ = v44
	var v45 int32
	_ = v45
	var v46 int32
	_ = v46
	var v48 int32
	_ = v48
	var v49 int32
	_ = v49
	v10 = m.G0
	v12 = v10 - int32(16)
	m.G0 = v12
	v14 = int32(4443856)
	v15 = *(*int32)(unsafe.Add(mBase, _consts[0]))
	v17 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
	*(*int32)(unsafe.Add(mBase, _consts[0])) = v17
	v19 = *(*int32)(unsafe.Add(mBase, uint32(l0)+60))
	v20 = F_tuplesort_gettuple_common(m, l0, l1, v12)
	mBase = m.M
	v23 = m.ExcPending
	if v23 != 0 {
		return int32(0)
	} else {
		*(*int32)(unsafe.Add(mBase, _consts[0])) = v15
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
					v39 = int32(0)
					v40 = *(*int32)(unsafe.Add(mBase, uint32(v12)))
					if l2 == v39 {
						v48 = v39
						v49 = v40
						*(*int32)(unsafe.Add(mBase, uint32(l3))) = v49
						*(*uint8)(unsafe.Add(mBase, uint32(l4))) = uint8(v48)
						m.G0 = v12 + int32(16)
						return v20
					} else {
						v44 = *(*int32)(unsafe.Add(mBase, uint32(v19)+4))
						v45 = F_datumCopy(m, v40, int32(0), v44)
						mBase = m.M
						v46 = m.ExcPending
						if v46 != 0 {
							return int32(0)
						} else {
							v48 = v39
							v49 = v45
							*(*int32)(unsafe.Add(mBase, uint32(l3))) = v49
							*(*uint8)(unsafe.Add(mBase, uint32(l4))) = uint8(v48)
							m.G0 = v12 + int32(16)
							return v20
						}
					}
				} else {
					v38 = *(*int32)(unsafe.Add(mBase, uint32(v12)+4))
					v48 = v34
					v49 = v38
					*(*int32)(unsafe.Add(mBase, uint32(l3))) = v49
					*(*uint8)(unsafe.Add(mBase, uint32(l4))) = uint8(v48)
					m.G0 = v12 + int32(16)
					return v20
				}
			} else {
				v38 = *(*int32)(unsafe.Add(mBase, uint32(v12)+4))
				v48 = v34
				v49 = v38
				*(*int32)(unsafe.Add(mBase, uint32(l3))) = v49
				*(*uint8)(unsafe.Add(mBase, uint32(l4))) = uint8(v48)
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
	var v197 int64
	_ = v197
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
	var v282 int32
	_ = v282
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
	var v373 int64
	_ = v373
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
	var v467 int32
	_ = v467
	var v468 int32
	_ = v468
	var v469 int64
	_ = v469
	var v471 int64
	_ = v471
	var v474 int32
	_ = v474
	var v479 int32
	_ = v479
	var v481 int32
	_ = v481
	var v482 int32
	_ = v482
	var v483 int32
	_ = v483
	var v484 int32
	_ = v484
	var v485 int32
	_ = v485
	var v494 int32
	_ = v494
	var v496 int32
	_ = v496
	var v498 int32
	_ = v498
	var v509 int32
	_ = v509
	var v511 int32
	_ = v511
	var v517 int32
	_ = v517
	var v518 int32
	_ = v518
	var v519 int32
	_ = v519
	var v522 int32
	_ = v522
	var v523 int32
	_ = v523
	var v526 int32
	_ = v526
	var v527 int32
	_ = v527
	var v528 int32
	_ = v528
	var v529 int32
	_ = v529
	var v534 int32
	_ = v534
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
	var v568 int32
	_ = v568
	var v569 int32
	_ = v569
	var v572 int32
	_ = v572
	var v573 int64
	_ = v573
	var v575 int64
	_ = v575
	var v577 int32
	_ = v577
	var v598 int32
	_ = v598
	var v601 int32
	_ = v601
	var v604 int32
	_ = v604
	var v605 int32
	_ = v605
	var v620 int32
	_ = v620
	var v621 int32
	_ = v621
	var v622 int32
	_ = v622
	var v624 int32
	_ = v624
	var v626 int32
	_ = v626
	var v631 int32
	_ = v631
	var v632 int32
	_ = v632
	var v651 int32
	_ = v651
	var v655 int32
	_ = v655
	var v662 int32
	_ = v662
	var v666 int32
	_ = v666
	var v671 int32
	_ = v671
	var v674 int32
	_ = v674
	var v676 int32
	_ = v676
	var v679 int32
	_ = v679
	var v698 int32
	_ = v698
	var v703 int32
	_ = v703
	var v706 int32
	_ = v706
	var v709 int32
	_ = v709
	var v710 int32
	_ = v710
	var v716 int32
	_ = v716
	var v717 int32
	_ = v717
	var v720 int32
	_ = v720
	var v721 int32
	_ = v721
	var v727 int32
	_ = v727
	var v731 int32
	_ = v731
	var v734 int32
	_ = v734
	var v735 int32
	_ = v735
	var v742 int32
	_ = v742
	var v746 int32
	_ = v746
	var v749 int32
	_ = v749
	var v760 int32
	_ = v760
	var v764 int32
	_ = v764
	var v769 int32
	_ = v769
	v18 = m.G0
	v20 = v18 + int32(-64)
	m.G0 = v20
	v22 = int32(4443856)
	v23 = *(*int32)(unsafe.Add(mBase, _consts[0]))
	v25 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
	*(*int32)(unsafe.Add(mBase, _consts[0])) = v25
	v28 = int32(*(*uint8)(unsafe.Add(mBase, _consts[33])))
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
	F_errmsg_internal(m, int32(193013), v18+int32(-32))
	mBase = m.M
	v48 = m.ExcPending
	if v48 != 0 {
		goto L3
	} else {
		goto L7
	}
L7:
	;
	F_errfinish(m, int32(471425), int32(1369), int32(74925))
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
	v760 = m.ExcPending
	if v760 != 0 {
		goto L3
	} else {
		goto L132
	}
L10:
	;
	v698 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(l0)+228)) = uint8(v698)
	*(*int32)(unsafe.Add(mBase, uint32(l0)+224)) = v698
	v703 = int32(*(*uint8)(unsafe.Add(mBase, _consts[33])))
	if v703 != int32(1) {
		goto L118
	} else {
		goto L119
	}
L11:
	;
	F_dumptuples(m, l0, int32(1))
	mBase = m.M
	v674 = m.ExcPending
	if v674 != 0 {
		goto L3
	} else {
		goto L116
	}
L12:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v662 = m.ExcPending
	if v662 != 0 {
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
	F_s_lock(m, v74, int32(471425), int32(3034), int32(355422))
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
	F_s_lock(m, v56, int32(471425), int32(3079), int32(153368))
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
	v197 = int64(0)
	*(*int64)(unsafe.Add(mBase, uint32(v188)+32)) = v197
	*(*int64)(unsafe.Add(mBase, uint32(v188)+52)) = v197
	*(*int32)(unsafe.Add(mBase, uint32(v188)+48)) = int32(1073741823)
	*(*int64)(unsafe.Add(mBase, uint32(v188)+16)) = v190
	*(*int64)(unsafe.Add(mBase, uint32(v188)+24)) = v190
	*(*int64)(unsafe.Add(mBase, uint32(v188)+40)) = v197
	*(*int64)(unsafe.Add(mBase, uint32(v188)+60)) = v197
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
	*(*int64)(unsafe.Add(mBase, uint32(v188)+32)) = v373
	v375 = int64(1073741823)
	if v375 <= v236 {
		goto L76
	} else {
		goto L77
	}
L56:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v178))) = v234
	v373 = int64(0)
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
	v373 = base.I64_extend_i32_s(v248) << (uint(int64(17)) % 64)
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
	v282 = v275
	goto L70
L70:
	;
	v294 = *(*int32)(unsafe.Add(mBase, uint32(v240)+4))
	v295 = int32(2)
	v298 = *(*int32)(unsafe.Add(mBase, uint32(v234)+4))
	v299 = *(*int32)(unsafe.Add(mBase, uint32(v240)))
	v304 = *(*int32)(unsafe.Add(mBase, uint32(v298+(v282-v299)<<(uint(v295)%32))))
	*(*int32)(unsafe.Add(mBase, uint32(v294+v282<<(uint(v295)%32)))) = v304
	v306 = *(*int32)(unsafe.Add(mBase, uint32(v240)+4))
	v308 = v282 + int32(1)
	v312 = *(*int32)(unsafe.Add(mBase, uint32(v234)+4))
	v313 = *(*int32)(unsafe.Add(mBase, uint32(v240)))
	v318 = *(*int32)(unsafe.Add(mBase, uint32(v312+(v308-v313)<<(uint(v295)%32))))
	*(*int32)(unsafe.Add(mBase, uint32(v306+v308<<(uint(v295)%32)))) = v318
	v321 = v282 + v295
	if v321 != v250 {
		v282 = v321
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
	F_errmsg_internal(m, int32(208313), int32(0))
	mBase = m.M
	v351 = m.ExcPending
	if v351 != 0 {
		goto L3
	} else {
		goto L74
	}
L74:
	;
	F_errfinish(m, int32(477271), int32(912), int32(407779))
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
	*(*int64)(unsafe.Add(mBase, uint32(v178)+32)) = v380 + (v373 - v381)
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
	v598 = *(*int32)(unsafe.Add(mBase, uint32(l0)+40))
	if int32(0) < v598 {
		goto L107
	} else {
		goto L108
	}
L84:
	;
	v467 = v18 + int32(-8)
	v468 = *(*int32)(unsafe.Add(mBase, uint32(l0)+132))
	v469 = *(*int64)(unsafe.Add(mBase, uint32(v468)+8))
	*(*int64)(unsafe.Add(mBase, uint32(v467))) = v469
	v471 = *(*int64)(unsafe.Add(mBase, uint32(v468)))
	*(*int64)(unsafe.Add(mBase, uint32(v20)+48)) = v471
	v474 = v450 - int32(1)
	*(*int32)(unsafe.Add(mBase, uint32(l0)+136)) = v474
	v479 = *(*int32)(unsafe.Add(mBase, _consts[1]))
	if v479 != 0 {
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
	v481 = m.ExcPending
	if v481 != 0 {
		goto L3
	} else {
		goto L89
	}
L87:
	;
	v483 = v474
	goto L88
L88:
	;
	v484 = v474<<(uint(int32(4))%32) + v468
	v485 = int32(0)
	if base.Ui32(v483) < base.Ui32(int32(2)) {
		v545 = v485
		goto L90
	} else {
		goto L91
	}
L89:
	;
	v482 = *(*int32)(unsafe.Add(mBase, uint32(l0)+136))
	v483 = v482
	goto L88
L90:
	;
	v561 = int32(4)
	v563 = v468 + v545<<(uint(v561)%32)
	v564 = *(*int64)(unsafe.Add(mBase, uint32(v484)))
	*(*int64)(unsafe.Add(mBase, uint32(v563))) = v564
	v566 = *(*int64)(unsafe.Add(mBase, uint32(v484)+8))
	*(*int64)(unsafe.Add(mBase, uint32(v563)+8)) = v566
	v568 = *(*int32)(unsafe.Add(mBase, uint32(l0)+132))
	v569 = *(*int32)(unsafe.Add(mBase, uint32(l0)+136))
	v572 = v568 + v569<<(uint(v561)%32)
	v573 = *(*int64)(unsafe.Add(mBase, uint32(v20)+48))
	*(*int64)(unsafe.Add(mBase, uint32(v572))) = v573
	v575 = *(*int64)(unsafe.Add(mBase, uint32(v467)))
	*(*int64)(unsafe.Add(mBase, uint32(v572)+8)) = v575
	v577 = *(*int32)(unsafe.Add(mBase, uint32(l0)+136))
	if int32(1) < v577 {
		v450 = v577
		goto L84
	} else {
		goto L106
	}
L91:
	;
	v494 = v485
	v496 = v485
	v498 = int32(1)
	goto L92
L92:
	;
	v509 = v494 + int32(2)
	if base.Ui32(v483) <= base.Ui32(v509) {
		goto L94
	} else {
		goto L95
	}
L93:
	;
	v545 = v523
	goto L90
L94:
	;
	v523 = v498
	goto L96
L95:
	;
	v511 = int32(4)
	v517 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v518 = m.T0[v517].(func(*base.Module, int32, int32, int32) int32)(m, v468+v498<<(uint(v511)%32), v468+v509<<(uint(v511)%32), l0)
	mBase = m.M
	v519 = m.ExcPending
	if v519 != 0 {
		goto L3
	} else {
		goto L97
	}
L96:
	;
	v526 = v468 + v523<<(uint(int32(4))%32)
	v527 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v528 = m.T0[v527].(func(*base.Module, int32, int32, int32) int32)(m, v484, v526, l0)
	mBase = m.M
	v529 = m.ExcPending
	if v529 != 0 {
		goto L3
	} else {
		goto L101
	}
L97:
	;
	if int32(0) < v518 {
		goto L98
	} else {
		goto L99
	}
L98:
	;
	v522 = v509
	goto L100
L99:
	;
	v522 = v498
	goto L100
L100:
	;
	v523 = v522
	goto L96
L101:
	;
	if v528 <= int32(0) {
		goto L102
	} else {
		goto L103
	}
L102:
	;
	v545 = v496
	goto L90
L103:
	;
	goto L104
L104:
	;
	v534 = v468 + v496<<(uint(int32(4))%32)
	v535 = *(*int64)(unsafe.Add(mBase, uint32(v526)))
	*(*int64)(unsafe.Add(mBase, uint32(v534))) = v535
	v537 = *(*int64)(unsafe.Add(mBase, uint32(v526)+8))
	*(*int64)(unsafe.Add(mBase, uint32(v534)+8)) = v537
	v539 = int32(1)
	v540 = v523 << (uint(v539) % 32)
	v542 = v540 | v539
	if base.Ui32(v542) < base.Ui32(v483) {
		v494 = v540
		v496 = v523
		v498 = v542
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
	v601 = *(*int32)(unsafe.Add(mBase, uint32(l0)+44))
	v604 = v601
	v605 = int32(0)
	goto L110
L108:
	;
	goto L109
L109:
	;
	v651 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(l0)+208)) = uint8(v651)
	*(*int32)(unsafe.Add(mBase, uint32(l0)+204)) = v651
	v655 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(l0)+69)) = uint8(v655)
	*(*int32)(unsafe.Add(mBase, uint32(l0)+64)) = int32(3)
	goto L10
L110:
	;
	v620 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v604)+8)))
	v621 = int32(1)
	v622 = v620 ^ v621
	*(*uint8)(unsafe.Add(mBase, uint32(v604)+8)) = uint8(v622)
	v624 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v604)+9)))
	v626 = v624 ^ v621
	*(*uint8)(unsafe.Add(mBase, uint32(v604)+9)) = uint8(v626)
	v631 = v605 + v621
	v632 = *(*int32)(unsafe.Add(mBase, uint32(l0)+40))
	if v631 < v632 {
		v604 = v604 + int32(36)
		v605 = v631
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
	F_errmsg_internal(m, int32(336171), int32(0))
	mBase = m.M
	v666 = m.ExcPending
	if v666 != 0 {
		goto L3
	} else {
		goto L114
	}
L114:
	;
	F_errfinish(m, int32(471425), int32(1444), int32(74925))
	mBase = m.M
	v671 = m.ExcPending
	if v671 != 0 {
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
	v676 = m.ExcPending
	if v676 != 0 {
		goto L3
	} else {
		goto L117
	}
L117:
	;
	*(*int64)(unsafe.Add(mBase, uint32(l0)+216)) = int64(0)
	v679 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(l0)+208)) = uint8(v679)
	goto L10
L118:
	;
	*(*int32)(unsafe.Add(mBase, _consts[0])) = v23
	m.G0 = v20 - int32(-64)
	return
L119:
	;
	v706 = *(*int32)(unsafe.Add(mBase, uint32(l0)+64))
	v709 = F_errstart(m, int32(15), int32(0))
	mBase = m.M
	v710 = m.ExcPending
	if v710 != 0 {
		goto L3
	} else {
		goto L120
	}
L120:
	;
	if v706 == int32(5) {
		goto L122
	} else {
		goto L123
	}
L121:
	;
	F_errfinish(m, int32(471425), v746, int32(74925))
	mBase = m.M
	v749 = m.ExcPending
	if v749 != 0 {
		goto L3
	} else {
		goto L131
	}
L122:
	;
	if v709 == int32(0) {
		goto L118
	} else {
		goto L125
	}
L123:
	;
	goto L124
L124:
	;
	if v709 == int32(0) {
		goto L118
	} else {
		goto L128
	}
L125:
	;
	v716 = *(*int32)(unsafe.Add(mBase, uint32(l0)+232))
	v717 = *(*int32)(unsafe.Add(mBase, uint32(l0)+176))
	v720 = F_pg_rusage_show(m, l0+int32(256))
	mBase = m.M
	v721 = m.ExcPending
	if v721 != 0 {
		goto L3
	} else {
		goto L126
	}
L126:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v20)+8)) = v720
	*(*int32)(unsafe.Add(mBase, uint32(v20)+4)) = v717
	*(*int32)(unsafe.Add(mBase, uint32(v20))) = v716
	F_errmsg_internal(m, int32(195124), v20)
	mBase = m.M
	v727 = m.ExcPending
	if v727 != 0 {
		goto L3
	} else {
		goto L127
	}
L127:
	;
	v746 = int32(1453)
	goto L121
L128:
	;
	v731 = *(*int32)(unsafe.Add(mBase, uint32(l0)+232))
	v734 = F_pg_rusage_show(m, l0+int32(256))
	mBase = m.M
	v735 = m.ExcPending
	if v735 != 0 {
		goto L3
	} else {
		goto L129
	}
L129:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v20)+20)) = v734
	*(*int32)(unsafe.Add(mBase, uint32(v20)+16)) = v731
	F_errmsg_internal(m, int32(193713), v18+int32(-48))
	mBase = m.M
	v742 = m.ExcPending
	if v742 != 0 {
		goto L3
	} else {
		goto L130
	}
L130:
	;
	v746 = int32(1456)
	goto L121
L131:
	;
	goto L118
L132:
	;
	F_errmsg_internal(m, int32(308004), int32(0))
	mBase = m.M
	v764 = m.ExcPending
	if v764 != 0 {
		goto L3
	} else {
		goto L133
	}
L133:
	;
	F_errfinish(m, int32(471425), int32(3084), int32(153368))
	mBase = m.M
	v769 = m.ExcPending
	if v769 != 0 {
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
	var v6 int32
	_ = v6
	var v8 int32
	_ = v8
	var v10 int32
	_ = v10
	var v11 int32
	_ = v11
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
	var v24 int32
	_ = v24
	var v27 int32
	_ = v27
	var v28 int32
	_ = v28
	var v33 int32
	_ = v33
	var v35 int32
	_ = v35
	var v36 int32
	_ = v36
	var v39 int32
	_ = v39
	var v41 int32
	_ = v41
	v6 = m.G0
	v8 = v6 - int32(16)
	m.G0 = v8
	v10 = int32(4443856)
	v11 = *(*int32)(unsafe.Add(mBase, _consts[0]))
	v13 = *(*int32)(unsafe.Add(mBase, uint32(l0)+32))
	*(*int32)(unsafe.Add(mBase, _consts[0])) = v13
	v15 = F_palloc(m, l2)
	mBase = m.M
	v16 = m.ExcPending
	if v16 != 0 {
		return
	} else {
		if l2 != 0 {
			v17 = F__emscripten_memcpy_bulkmem(m, v15, l1, l2)
			mBase = m.M
			v18 = v17
		} else {
			v18 = v15
		}
		v19 = int32(0)
		*(*uint8)(unsafe.Add(mBase, uint32(v8)+8)) = uint8(v19)
		*(*int32)(unsafe.Add(mBase, uint32(v8)+4)) = v19
		*(*int32)(unsafe.Add(mBase, uint32(v8))) = v18
		v24 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+52)))
		if v24&int32(2) != 0 {
			v27 = F_GetMemoryChunkSpace(m, v18)
			mBase = m.M
			v28 = m.ExcPending
			if v28 != 0 {
				return
			} else {
				v33 = v27
				v35 = *(*int32)(unsafe.Add(mBase, uint32(l0)+44))
				if v35 != 0 {
					v36 = *(*int32)(unsafe.Add(mBase, uint32(v35)+24))
					v39 = base.B2i32(v36 != int32(0))
				} else {
					v39 = int32(0)
				}
				F_tuplesort_puttuple_common(m, l0, v8, v39, v33)
				mBase = m.M
				v41 = m.ExcPending
				if v41 != 0 {
					return
				} else {
					*(*int32)(unsafe.Add(mBase, _consts[0])) = v11
					m.G0 = v8 + int32(16)
					return
				}
			}
		} else {
			v33 = (l2 + int32(7)) & int32(-8)
			v35 = *(*int32)(unsafe.Add(mBase, uint32(l0)+44))
			if v35 != 0 {
				v36 = *(*int32)(unsafe.Add(mBase, uint32(v35)+24))
				v39 = base.B2i32(v36 != int32(0))
			} else {
				v39 = int32(0)
			}
			F_tuplesort_puttuple_common(m, l0, v8, v39, v33)
			mBase = m.M
			v41 = m.ExcPending
			if v41 != 0 {
				return
			} else {
				*(*int32)(unsafe.Add(mBase, _consts[0])) = v11
				m.G0 = v8 + int32(16)
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
	v11 = int32(4443856)
	v12 = *(*int32)(unsafe.Add(mBase, _consts[0]))
	v14 = *(*int32)(unsafe.Add(mBase, uint32(l0)+32))
	*(*int32)(unsafe.Add(mBase, _consts[0])) = v14
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
					*(*int32)(unsafe.Add(mBase, _consts[0])) = v12
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
						*(*int32)(unsafe.Add(mBase, _consts[0])) = v12
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
	var v6 int32
	_ = v6
	var v9 int64
	_ = v9
	var v10 int64
	_ = v10
	var v12 int32
	_ = v12
	var v13 int64
	_ = v13
	var v14 int64
	_ = v14
	var v17 int64
	_ = v17
	var v18 int32
	_ = v18
	var v19 int32
	_ = v19
	var v22 int64
	_ = v22
	var v23 int32
	_ = v23
	var v28 int64
	_ = v28
	var v30 int64
	_ = v30
	var v32 int32
	_ = v32
	var v35 int32
	_ = v35
	var v41 int32
	_ = v41
	var v43 int32
	_ = v43
	var v44 int32
	_ = v44
	v6 = *(*int32)(unsafe.Add(mBase, uint32(l0)+128))
	if v6 == int32(0) {
		v9 = *(*int64)(unsafe.Add(mBase, uint32(l0)+96))
		v10 = *(*int64)(unsafe.Add(mBase, uint32(l0)+88))
		v12 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+120)))
		v22 = v9 - v10
		v23 = v12
		if v23 != base.B2i32(v6 != int32(0)) {
		} else {
			v28 = *(*int64)(unsafe.Add(mBase, uint32(l0)+112))
			if v22 <= v28 {
			} else {
				v30 = v22
				v32 = v23
				*(*uint8)(unsafe.Add(mBase, uint32(l0)+120)) = uint8(v32)
				*(*int64)(unsafe.Add(mBase, uint32(l0)+112)) = v30
				v35 = *(*int32)(unsafe.Add(mBase, uint32(l0)+64))
				*(*int32)(unsafe.Add(mBase, uint32(l0)+124)) = v35
			}
		}
	} else {
		v13 = *(*int64)(unsafe.Add(mBase, uint32(v6)+24))
		v14 = *(*int64)(unsafe.Add(mBase, uint32(v6)+32))
		v17 = (v13 - v14) << (uint(int64(13)) % 64)
		v18 = int32(1)
		v19 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+120)))
		if v19 != v18 {
			v30 = v17
			v32 = v18
			*(*uint8)(unsafe.Add(mBase, uint32(l0)+120)) = uint8(v32)
			*(*int64)(unsafe.Add(mBase, uint32(l0)+112)) = v30
			v35 = *(*int32)(unsafe.Add(mBase, uint32(l0)+64))
			*(*int32)(unsafe.Add(mBase, uint32(l0)+124)) = v35
		} else {
			v22 = v17
			v23 = v19
			if v23 != base.B2i32(v6 != int32(0)) {
			} else {
				v28 = *(*int64)(unsafe.Add(mBase, uint32(l0)+112))
				if v22 <= v28 {
				} else {
					v30 = v22
					v32 = v23
					*(*uint8)(unsafe.Add(mBase, uint32(l0)+120)) = uint8(v32)
					*(*int64)(unsafe.Add(mBase, uint32(l0)+112)) = v30
					v35 = *(*int32)(unsafe.Add(mBase, uint32(l0)+64))
					*(*int32)(unsafe.Add(mBase, uint32(l0)+124)) = v35
				}
			}
		}
	}
	F_tuplesort_free(m, l0)
	mBase = m.M
	v41 = m.ExcPending
	if v41 != 0 {
		return
	} else {
		F_tuplesort_begin_batch(m, l0)
		mBase = m.M
		v43 = m.ExcPending
		if v43 != 0 {
			return
		} else {
			v44 = int32(0)
			*(*int32)(unsafe.Add(mBase, uint32(l0)+156)) = v44
			*(*int64)(unsafe.Add(mBase, uint32(l0)+148)) = int64(0)
			*(*int32)(unsafe.Add(mBase, uint32(l0)+164)) = v44
			return
		}
	}
}
