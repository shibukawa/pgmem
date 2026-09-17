package p1

import (
	base "github.com/shibukawa/pgmem/internal/aot/pgaot/base"
	"math"
	"unsafe"
)

func F_InSecurityRestrictedOperation(m *base.Module) int32 {
	mBase := m.M
	_ = mBase
	var v2 int32
	_ = v2
	v2 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_InSecurityRestrictedOperation[0])))
	return int32(base.Ui32(v2&int32(2)) >> (uint(int32(1)) % 32))
}
func F_IsInParallelMode(m *base.Module) int32 {
	mBase := m.M
	_ = mBase
	var v3 int32
	_ = v3
	var v4 int32
	_ = v4
	var v6 int32
	_ = v6
	var v7 int32
	_ = v7
	v3 = *(*int32)(unsafe.Add(mBase, _c_F_IsInParallelMode[0]))
	v4 = *(*int32)(unsafe.Add(mBase, uint32(v3)+72))
	if v4 != 0 {
		v7 = int32(1)
	} else {
		v6 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v3)+76)))
		v7 = v6
	}
	return v7 & int32(1)
}
func F_in_grouping(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32) int32 {
	mBase := m.M
	_ = mBase
	var v10 int32
	_ = v10
	var v11 int32
	_ = v11
	var v13 int32
	_ = v13
	var v19 int32
	_ = v19
	var v26 int32
	_ = v26
	var v27 int32
	_ = v27
	var v29 int32
	_ = v29
	var v31 int32
	_ = v31
	var v37 int32
	_ = v37
	var v46 int32
	_ = v46
	var v50 int32
	_ = v50
	v10 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v11 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	if v11 < v10 {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	v13 = v10
	goto L3
L2:
	;
	v13 = v11
	goto L3
L3:
	;
	v19 = v10
	goto L5
L4:
	;
	return v50
L5:
	;
	if v19 == v13 {
		goto L7
	} else {
		goto L8
	}
L6:
	;
	v50 = int32(0)
	goto L4
L7:
	;
	return int32(-1)
L8:
	;
	goto L9
L9:
	;
	v26 = int32(1)
	v27 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v29 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v27+v19))))
	if l3 < v29 {
		v50 = v26
		goto L4
	} else {
		goto L10
	}
L10:
	;
	v31 = v29 - l2
	if v31 < int32(0) {
		v50 = v26
		goto L4
	} else {
		goto L11
	}
L11:
	;
	v37 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l1+int32(base.Ui32(v31)>>(uint(int32(3))%32))))))
	if int32(base.Ui32(v37)>>(uint(v31&int32(7))%32))&int32(1) == int32(0) {
		v50 = v26
		goto L4
	} else {
		goto L12
	}
L12:
	;
	v46 = v19 + int32(1)
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v46
	if l4 != 0 {
		v19 = v46
		goto L5
	} else {
		goto L13
	}
L13:
	;
	goto L6
}
func F_in_range_float4_float8(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v9 int32
	_ = v9
	var v10 float64
	_ = v10
	var v21 int32
	_ = v21
	var v22 int32
	_ = v22
	var v23 int32
	_ = v23
	var v24 int32
	_ = v24
	var v25 int32
	_ = v25
	var v41 int32
	_ = v41
	var v42 float32
	_ = v42
	var v63 float64
	_ = v63
	var v65 float64
	_ = v65
	var v67 float64
	_ = v67
	var v77 int32
	_ = v77
	var v80 int32
	_ = v80
	var v84 int32
	_ = v84
	var v89 int32
	_ = v89
	v9 = *(*int32)(unsafe.Add(mBase, uint32(l0)+36))
	v10 = *(*float64)(unsafe.Add(mBase, uint32(v9)))
	if base.B2i32(base.Ui64(int64(9218868437227405312)) < base.Ui64(base.I64_reinterpret_f64(v10)&int64(9223372036854775807)))|base.F64_lt(v10, float64(0)) == int32(0) {
		v21 = *(*int32)(unsafe.Add(mBase, uint32(l0)+52))
		v22 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
		v23 = int32(2147483647)
		v24 = v22 & v23
		v25 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
		if base.Ui32(int32(2139095041)) <= base.Ui32(v25&v23) {
			return base.B2i32(v21 == int32(0)) | base.B2i32(base.Ui32(int32(2139095040)) < base.Ui32(v24))
		} else {
			if base.Ui32(int32(2139095040)) < base.Ui32(v24) {
				return base.B2i32(v21 != int32(0))
			} else {
				v41 = *(*int32)(unsafe.Add(mBase, uint32(l0)+44))
				v42 = base.F32_reinterpret_i32(v22)
				if base.F32_ne(base.F32_abs(v42), math.Float32frombits(uint32(0x7f800000)))|base.F64_ne(base.F64_abs(v10), math.Float64frombits(uint64(0x7ff0000000000000))) != 0 {
					if v41 != 0 {
						v63 = base.F64_neg(v10)
					} else {
						v63 = v10
					}
					v65 = base.F64_add(v63, base.F64_promote_f32(v42))
					v67 = base.F64_promote_f32(base.F32_reinterpret_i32(v25))
					if v21 != 0 {
						return base.F64_ge(v65, v67)
					} else {
						return base.F64_le(v65, v67)
					}
				} else {
					if v41 != 0 {
						if base.F32_gt(v42, float32(0)) == int32(0) {
							if v41 != 0 {
								v63 = base.F64_neg(v10)
							} else {
								v63 = v10
							}
							v65 = base.F64_add(v63, base.F64_promote_f32(v42))
							v67 = base.F64_promote_f32(base.F32_reinterpret_i32(v25))
							if v21 != 0 {
								return base.F64_ge(v65, v67)
							} else {
								return base.F64_le(v65, v67)
							}
						} else {
							return int32(1)
						}
					} else {
						if base.F32_lt(v42, float32(0)) == int32(0) {
							if v41 != 0 {
								v63 = base.F64_neg(v10)
							} else {
								v63 = v10
							}
							v65 = base.F64_add(v63, base.F64_promote_f32(v42))
							v67 = base.F64_promote_f32(base.F32_reinterpret_i32(v25))
							if v21 != 0 {
								return base.F64_ge(v65, v67)
							} else {
								return base.F64_le(v65, v67)
							}
						} else {
							return int32(1)
						}
					}
				}
			}
		}
	} else {
		F_errstart_cold(m, int32(21), int32(0))
		mBase = m.M
		v77 = m.ExcPending
		if v77 != 0 {
			return int32(0)
		} else {
			F_errcode(m, int32(50593922))
			mBase = m.M
			v80 = m.ExcPending
			if v80 != 0 {
				return int32(0)
			} else {
				F_errmsg(m, int32(_a_F_in_range_float4_float8_0), int32(0))
				mBase = m.M
				v84 = m.ExcPending
				if v84 != 0 {
					return int32(0)
				} else {
					F_errfinish(m, int32(_a_F_in_range_float4_float8_1), int32(1119), int32(_a_F_in_range_float4_float8_2))
					mBase = m.M
					v89 = m.ExcPending
					if v89 != 0 {
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
func F_in_range_int2_int4(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v6 int32
	_ = v6
	var v9 int32
	_ = v9
	var v12 int32
	_ = v12
	var v13 int32
	_ = v13
	var v16 int32
	_ = v16
	var v17 int32
	_ = v17
	var v20 int32
	_ = v20
	var v26 int32
	_ = v26
	var v36 int32
	_ = v36
	var v39 int32
	_ = v39
	var v43 int32
	_ = v43
	var v48 int32
	_ = v48
	v6 = *(*int32)(unsafe.Add(mBase, uint32(l0)+36))
	if int32(0) <= v6 {
		v9 = *(*int32)(unsafe.Add(mBase, uint32(l0)+52))
		v12 = *(*int32)(unsafe.Add(mBase, uint32(l0)+44))
		if v12 != 0 {
			v13 = int32(0) - v6
		} else {
			v13 = v6
		}
		v16 = int32(*(*int16)(unsafe.Add(mBase, uint32(l0)+28)))
		v17 = v16 + v13
		if base.B2i32(v13 < int32(0)) != base.B2i32(v17 < v16) {
			v20 = int32(0)
			return base.B2i32(v12 != v20) ^ base.B2i32(v9 != v20)
		} else {
			v26 = int32(*(*int16)(unsafe.Add(mBase, uint32(l0)+20)))
			if v9 != 0 {
				return base.B2i32(v26 <= v17)
			} else {
				return base.B2i32(v17 <= v26)
			}
		}
	} else {
		F_errstart_cold(m, int32(21), int32(0))
		mBase = m.M
		v36 = m.ExcPending
		if v36 != 0 {
			return int32(0)
		} else {
			F_errcode(m, int32(50593922))
			mBase = m.M
			v39 = m.ExcPending
			if v39 != 0 {
				return int32(0)
			} else {
				F_errmsg(m, int32(_a_F_in_range_int2_int4_0), int32(0))
				mBase = m.M
				v43 = m.ExcPending
				if v43 != 0 {
					return int32(0)
				} else {
					F_errfinish(m, int32(_a_F_in_range_int2_int4_1), int32(746), int32(_a_F_in_range_int2_int4_2))
					mBase = m.M
					v48 = m.ExcPending
					if v48 != 0 {
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
func F_in_range_int2_int8(m *base.Module, l0 int32) int32 {
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
	var v12 int32
	_ = v12
	v4 = int32(*(*int16)(unsafe.Add(mBase, uint32(l0)+20)))
	v5 = int32(*(*int16)(unsafe.Add(mBase, uint32(l0)+28)))
	v6 = *(*int32)(unsafe.Add(mBase, uint32(l0)+36))
	v7 = *(*int32)(unsafe.Add(mBase, uint32(l0)+44))
	v8 = *(*int32)(unsafe.Add(mBase, uint32(l0)+52))
	v9 = F_DirectFunctionCall5Coll(m, int32(1304), int32(0), v4, v5, v6, v7, v8)
	mBase = m.M
	v12 = m.ExcPending
	if v12 != 0 {
		return int32(0)
	} else {
		return v9
	}
}
func F_record_in(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v17 int32
	_ = v17
	var v19 int32
	_ = v19
	var v21 int32
	_ = v21
	var v22 int32
	_ = v22
	var v23 int32
	_ = v23
	var v24 int32
	_ = v24
	var v28 int32
	_ = v28
	var v31 int32
	_ = v31
	var v36 int32
	_ = v36
	var v37 int32
	_ = v37
	var v38 int32
	_ = v38
	var v43 int32
	_ = v43
	var v47 int32
	_ = v47
	var v52 int32
	_ = v52
	var v53 int32
	_ = v53
	var v54 int32
	_ = v54
	var v55 int32
	_ = v55
	var v56 int32
	_ = v56
	var v57 int32
	_ = v57
	var v60 int32
	_ = v60
	var v62 int32
	_ = v62
	var v63 int32
	_ = v63
	var v68 int32
	_ = v68
	var v69 int32
	_ = v69
	var v70 int32
	_ = v70
	var v72 int32
	_ = v72
	var v73 int32
	_ = v73
	var v77 int32
	_ = v77
	var v78 int32
	_ = v78
	var v80 int32
	_ = v80
	var v85 int32
	_ = v85
	var v87 int32
	_ = v87
	var v99 int32
	_ = v99
	var v101 int32
	_ = v101
	var v103 int32
	_ = v103
	var v108 int32
	_ = v108
	var v126 int32
	_ = v126
	var v127 int32
	_ = v127
	var v128 int32
	_ = v128
	var v129 int32
	_ = v129
	var v131 int32
	_ = v131
	var v146 int32
	_ = v146
	var v155 int32
	_ = v155
	var v156 int32
	_ = v156
	var v161 int32
	_ = v161
	var v165 int32
	_ = v165
	var v169 int32
	_ = v169
	var v174 int32
	_ = v174
	var v176 int32
	_ = v176
	var v180 int32
	_ = v180
	var v185 int32
	_ = v185
	var v188 int32
	_ = v188
	var v189 int32
	_ = v189
	var v190 int32
	_ = v190
	var v203 int32
	_ = v203
	var v209 int32
	_ = v209
	var v210 int32
	_ = v210
	var v219 int32
	_ = v219
	var v221 int32
	_ = v221
	var v222 int32
	_ = v222
	var v227 int32
	_ = v227
	var v228 int32
	_ = v228
	var v230 int32
	_ = v230
	var v233 int32
	_ = v233
	var v234 int32
	_ = v234
	var v239 int32
	_ = v239
	var v245 int32
	_ = v245
	var v249 int32
	_ = v249
	var v254 int32
	_ = v254
	var v256 int32
	_ = v256
	var v257 int32
	_ = v257
	var v258 int32
	_ = v258
	var v265 int32
	_ = v265
	var v271 int32
	_ = v271
	var v280 int32
	_ = v280
	var v286 int32
	_ = v286
	var v291 int32
	_ = v291
	var v292 int32
	_ = v292
	var v297 int32
	_ = v297
	var v303 int32
	_ = v303
	var v307 int32
	_ = v307
	var v312 int32
	_ = v312
	var v313 int32
	_ = v313
	var v316 int32
	_ = v316
	var v317 int32
	_ = v317
	var v322 int32
	_ = v322
	var v328 int32
	_ = v328
	var v332 int32
	_ = v332
	var v337 int32
	_ = v337
	var v340 int32
	_ = v340
	var v346 int32
	_ = v346
	var v353 int32
	_ = v353
	var v354 int32
	_ = v354
	var v355 int32
	_ = v355
	var v360 int32
	_ = v360
	var v362 int32
	_ = v362
	var v364 int32
	_ = v364
	var v365 int32
	_ = v365
	var v370 int32
	_ = v370
	var v383 int32
	_ = v383
	var v384 int32
	_ = v384
	var v391 int32
	_ = v391
	var v392 int32
	_ = v392
	var v395 int32
	_ = v395
	var v396 int32
	_ = v396
	var v398 int32
	_ = v398
	var v403 int32
	_ = v403
	var v406 int32
	_ = v406
	var v410 int32
	_ = v410
	var v411 int32
	_ = v411
	var v415 int32
	_ = v415
	var v416 int32
	_ = v416
	var v431 int32
	_ = v431
	var v434 int32
	_ = v434
	var v449 int32
	_ = v449
	var v453 int32
	_ = v453
	var v470 int32
	_ = v470
	var v480 int32
	_ = v480
	var v481 int32
	_ = v481
	var v486 int32
	_ = v486
	var v492 int32
	_ = v492
	var v496 int32
	_ = v496
	var v501 int32
	_ = v501
	var v502 int32
	_ = v502
	var v503 int32
	_ = v503
	var v508 int32
	_ = v508
	var v514 int32
	_ = v514
	var v518 int32
	_ = v518
	var v523 int32
	_ = v523
	var v524 int32
	_ = v524
	var v525 int32
	_ = v525
	var v526 int32
	_ = v526
	var v527 int32
	_ = v527
	var v528 int32
	_ = v528
	var v529 int32
	_ = v529
	var v530 int32
	_ = v530
	var v533 int32
	_ = v533
	var v534 int32
	_ = v534
	var v536 int32
	_ = v536
	var v538 int32
	_ = v538
	var v540 int32
	_ = v540
	var v541 int32
	_ = v541
	var v545 int32
	_ = v545
	var v546 int32
	_ = v546
	var v547 int32
	_ = v547
	var v564 int32
	_ = v564
	var v565 int32
	_ = v565
	var v569 int32
	_ = v569
	var v570 int32
	_ = v570
	var v573 int32
	_ = v573
	v17 = m.G0
	v19 = v17 - int32(112)
	m.G0 = v19
	v21 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
	v22 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v23 = *(*int32)(unsafe.Add(mBase, uint32(l0)+36))
	v24 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	F_check_stack_depth(m)
	mBase = m.M
	v28 = m.ExcPending
	if v28 != 0 {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	return int32(0)
L2:
	;
	v31 = int32(0)
	if base.B2i32(v21 != int32(2249))|base.B2i32(v31 <= v23) == v31 {
		goto L4
	} else {
		goto L5
	}
L3:
	;
	m.G0 = v19 + int32(112)
	return v573
L4:
	;
	v36 = int32(0)
	v37 = F_errsave_start(m, v22)
	mBase = m.M
	v38 = m.ExcPending
	if v38 != 0 {
		goto L1
	} else {
		goto L7
	}
L5:
	;
	goto L6
L6:
	;
	v53 = F_lookup_rowtype_tupdesc(m, v21, v23)
	mBase = m.M
	v54 = m.ExcPending
	if v54 != 0 {
		goto L1
	} else {
		goto L12
	}
L7:
	;
	if v37 == int32(0) {
		v573 = v36
		goto L3
	} else {
		goto L8
	}
L8:
	;
	F_errcode(m, int32(1088))
	mBase = m.M
	v43 = m.ExcPending
	if v43 != 0 {
		goto L1
	} else {
		goto L9
	}
L9:
	;
	F_errmsg(m, int32(_a_F_record_in_0), int32(0))
	mBase = m.M
	v47 = m.ExcPending
	if v47 != 0 {
		goto L1
	} else {
		goto L10
	}
L10:
	;
	F_errsave_finish(m, v22, int32(_a_F_record_in_1), int32(105), int32(_a_F_record_in_2))
	mBase = m.M
	v52 = m.ExcPending
	if v52 != 0 {
		goto L1
	} else {
		goto L11
	}
L11:
	;
	v573 = v36
	goto L3
L12:
	;
	v55 = *(*int32)(unsafe.Add(mBase, uint32(v53)))
	v56 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v57 = *(*int32)(unsafe.Add(mBase, uint32(v56)+16))
	if v57 == int32(0) {
		goto L15
	} else {
		goto L16
	}
L13:
	;
	v126 = F_palloc(m, v55<<(uint(int32(2))%32))
	mBase = m.M
	v127 = m.ExcPending
	if v127 != 0 {
		goto L1
	} else {
		goto L33
	}
L14:
	;
	if v78 == v21 {
		goto L19
	} else {
		goto L20
	}
L15:
	;
	v63 = *(*int32)(unsafe.Add(mBase, uint32(v56)+20))
	v68 = F_MemoryContextAlloc(m, v63, v55*int32(44)+int32(12))
	mBase = m.M
	v69 = m.ExcPending
	if v69 != 0 {
		goto L1
	} else {
		goto L18
	}
L16:
	;
	v60 = *(*int32)(unsafe.Add(mBase, uint32(v57)+8))
	if v60 != v55 {
		goto L15
	} else {
		goto L17
	}
L17:
	;
	v62 = *(*int32)(unsafe.Add(mBase, uint32(v57)))
	v77 = v57
	v78 = v62
	goto L14
L18:
	;
	v70 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	*(*int32)(unsafe.Add(mBase, uint32(v70)+16)) = v68
	v72 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v73 = *(*int32)(unsafe.Add(mBase, uint32(v72)+16))
	*(*int64)(unsafe.Add(mBase, uint32(v73))) = int64(0)
	v77 = v73
	v78 = int32(0)
	goto L14
L19:
	;
	v80 = *(*int32)(unsafe.Add(mBase, uint32(v77)+4))
	if v80 == v23 {
		goto L13
	} else {
		goto L22
	}
L20:
	;
	goto L21
L21:
	;
	v85 = v55 * int32(44)
	v87 = v85 + int32(12)
	if v77&int32(3)|base.B2i32(base.Ui32(int32(1024)) < base.Ui32(v87)) == int32(0) {
		goto L24
	} else {
		goto L25
	}
L22:
	;
	goto L21
L23:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v77)+8)) = v55
	*(*int32)(unsafe.Add(mBase, uint32(v77)+4)) = v23
	*(*int32)(unsafe.Add(mBase, uint32(v77))) = v21
	goto L13
L24:
	;
	if v87 == int32(0) {
		goto L23
	} else {
		goto L27
	}
L25:
	;
	goto L26
L26:
	;
	if v87 == int32(0) {
		goto L23
	} else {
		goto L32
	}
L27:
	;
	v99 = v77 + v85 + int32(12)
	v101 = v77 + int32(4)
	if base.Ui32(v101) < base.Ui32(v99) {
		goto L28
	} else {
		goto L29
	}
L28:
	;
	v103 = v99
	goto L30
L29:
	;
	v103 = v101
	goto L30
L30:
	;
	v108 = (v77^int32(-1)+v103)&int32(-4) + int32(4)
	if v108 == int32(0) {
		goto L23
	} else {
		goto L31
	}
L31:
	;
	base.MemoryFill(m, v77, int32(0), v108)
	goto L23
L32:
	;
	base.MemoryFill(m, v77, int32(0), v87)
	goto L23
L33:
	;
	v128 = F_palloc(m, v55)
	mBase = m.M
	v129 = m.ExcPending
	if v129 != 0 {
		goto L1
	} else {
		goto L34
	}
L34:
	;
	v131 = v24
	goto L37
L35:
	;
	v564 = int32(0)
	v565 = *(*int32)(unsafe.Add(mBase, uint32(v53)+12))
	if v564 <= v565 {
		goto L145
	} else {
		goto L146
	}
L36:
	;
	v176 = v131 + int32(1)
	F_initStringInfo(m, v19+int32(96))
	mBase = m.M
	v180 = m.ExcPending
	if v180 != 0 {
		goto L1
	} else {
		goto L48
	}
L37:
	;
	v146 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v131))))
	if base.Ui32(v146-int32(9)) < base.Ui32(int32(5)) {
		goto L40
	} else {
		goto L41
	}
L38:
	;
	v155 = F_errsave_start(m, v22)
	mBase = m.M
	v156 = m.ExcPending
	if v156 != 0 {
		goto L1
	} else {
		goto L42
	}
L39:
	;
	goto L38
L40:
	;
	v131 = v131 + int32(1)
	goto L37
L41:
	;
	switch v146 - int32(32) {
	case 0:
		goto L40
	default:
		goto L39
	case 8:
		goto L36
	}
L42:
	;
	if v155 == int32(0) {
		goto L35
	} else {
		goto L43
	}
L43:
	;
	F_errcode(m, int32(33685634))
	mBase = m.M
	v161 = m.ExcPending
	if v161 != 0 {
		goto L1
	} else {
		goto L44
	}
L44:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v19))) = v24
	F_errmsg(m, int32(_a_F_record_in_3), v19)
	mBase = m.M
	v165 = m.ExcPending
	if v165 != 0 {
		goto L1
	} else {
		goto L45
	}
L45:
	;
	F_errdetail(m, int32(_a_F_record_in_4), int32(0))
	mBase = m.M
	v169 = m.ExcPending
	if v169 != 0 {
		goto L1
	} else {
		goto L46
	}
L46:
	;
	F_errsave_finish(m, v22, int32(_a_F_record_in_1), int32(159), int32(_a_F_record_in_2))
	mBase = m.M
	v174 = m.ExcPending
	if v174 != 0 {
		goto L1
	} else {
		goto L47
	}
L47:
	;
	goto L35
L48:
	;
	if int32(0) < v55 {
		goto L49
	} else {
		goto L50
	}
L49:
	;
	v185 = int32(0)
	v188 = v176
	v189 = v185
	v190 = v185
	goto L52
L50:
	;
	v434 = v176
	goto L51
L51:
	;
	v449 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v434))))
	if v449 == int32(41) {
		goto L112
	} else {
		goto L113
	}
L52:
	;
	v203 = *(*int32)(unsafe.Add(mBase, uint32(v53)))
	v209 = v53 + v203<<(uint(int32(4))%32) + v190*int32(100)
	v210 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v209)+111)))
	if v210 == int32(1) {
		goto L55
	} else {
		goto L56
	}
L53:
	;
	v434 = v415
	goto L51
L54:
	;
	v431 = v190 + int32(1)
	if v431 != v55 {
		v188 = v415
		v189 = v416
		v190 = v431
		goto L52
	} else {
		goto L110
	}
L55:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v126+v190<<(uint(int32(2))%32)))) = int32(0)
	v219 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(v190+v128))) = uint8(v219)
	v415 = v188
	v416 = v189
	goto L54
L56:
	;
	goto L57
L57:
	;
	v221 = *(*int32)(unsafe.Add(mBase, uint32(v209)+88))
	if v189 != 0 {
		goto L61
	} else {
		goto L62
	}
L58:
	;
	*(*uint8)(unsafe.Add(mBase, uint32(v190+v128))) = uint8(v365)
	v383 = v77 + int32(12) + v190*int32(44)
	v384 = *(*int32)(unsafe.Add(mBase, uint32(v383)))
	if v221 != v384 {
		goto L103
	} else {
		goto L104
	}
L59:
	;
	v256 = v19 + int32(96)
	v257 = *(*int32)(unsafe.Add(mBase, uint32(v256)))
	v258 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v257))) = uint8(v258)
	*(*int32)(unsafe.Add(mBase, uint32(v256)+12)) = v258
	*(*int32)(unsafe.Add(mBase, uint32(v256)+4)) = v258
	goto L71
L60:
	;
	v233 = F_errsave_start(m, v22)
	mBase = m.M
	v234 = m.ExcPending
	if v234 != 0 {
		goto L1
	} else {
		goto L65
	}
L61:
	;
	v222 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v188))))
	if v222 != int32(44) {
		goto L60
	} else {
		goto L64
	}
L62:
	;
	v227 = v188
	goto L63
L63:
	;
	v228 = int32(0)
	v230 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v227))))
	switch v230 - int32(41) {
	case 0, 3:
		v364 = v227
		v365 = int32(1)
		v370 = v228
		goto L58
	default:
		goto L59
	}
L64:
	;
	v227 = v188 + int32(1)
	goto L63
L65:
	;
	if v233 == int32(0) {
		goto L35
	} else {
		goto L66
	}
L66:
	;
	F_errcode(m, int32(33685634))
	mBase = m.M
	v239 = m.ExcPending
	if v239 != 0 {
		goto L1
	} else {
		goto L67
	}
L67:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v19)+80)) = v24
	F_errmsg(m, int32(_a_F_record_in_3), v19+int32(80))
	mBase = m.M
	v245 = m.ExcPending
	if v245 != 0 {
		goto L1
	} else {
		goto L68
	}
L68:
	;
	F_errdetail(m, int32(_a_F_record_in_5), int32(0))
	mBase = m.M
	v249 = m.ExcPending
	if v249 != 0 {
		goto L1
	} else {
		goto L69
	}
L69:
	;
	F_errsave_finish(m, v22, int32(_a_F_record_in_1), int32(191), int32(_a_F_record_in_2))
	mBase = m.M
	v254 = m.ExcPending
	if v254 != 0 {
		goto L1
	} else {
		goto L70
	}
L70:
	;
	goto L35
L71:
	;
	v265 = v227
	v271 = v228
	goto L72
L72:
	;
	v280 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v265))))
	if v271&int32(1) != 0 {
		goto L75
	} else {
		goto L76
	}
L73:
	;
	v362 = *(*int32)(unsafe.Add(mBase, uint32(v19)+96))
	v364 = v265
	v365 = int32(0)
	v370 = v362
	goto L58
L74:
	;
	goto L73
L75:
	;
	v286 = v265 + int32(1)
	if v280 != int32(34) {
		goto L78
	} else {
		goto L79
	}
L76:
	;
	switch v280 - int32(41) {
	case 0, 3:
		goto L74
	default:
		goto L75
	}
L77:
	;
	F_appendStringInfoChar(m, v19+int32(96), base.I32_extend8_s(v354))
	mBase = m.M
	v360 = m.ExcPending
	if v360 != 0 {
		goto L1
	} else {
		goto L102
	}
L78:
	;
	if v280 != int32(92) {
		goto L81
	} else {
		goto L82
	}
L79:
	;
	goto L80
L80:
	;
	v340 = int32(1)
	if v271&v340 == int32(0) {
		v265 = v286
		v271 = v340
		goto L72
	} else {
		goto L100
	}
L81:
	;
	if v280 != 0 {
		v353 = v286
		v354 = v280
		v355 = v271
		goto L77
	} else {
		goto L84
	}
L82:
	;
	goto L83
L83:
	;
	v313 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v286))))
	if v313 == int32(0) {
		goto L91
	} else {
		goto L92
	}
L84:
	;
	v291 = F_errsave_start(m, v22)
	mBase = m.M
	v292 = m.ExcPending
	if v292 != 0 {
		goto L1
	} else {
		goto L85
	}
L85:
	;
	if v291 == int32(0) {
		goto L35
	} else {
		goto L86
	}
L86:
	;
	F_errcode(m, int32(33685634))
	mBase = m.M
	v297 = m.ExcPending
	if v297 != 0 {
		goto L1
	} else {
		goto L87
	}
L87:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v19)+48)) = v24
	F_errmsg(m, int32(_a_F_record_in_3), v19+int32(48))
	mBase = m.M
	v303 = m.ExcPending
	if v303 != 0 {
		goto L1
	} else {
		goto L88
	}
L88:
	;
	F_errdetail(m, int32(_a_F_record_in_6), int32(0))
	mBase = m.M
	v307 = m.ExcPending
	if v307 != 0 {
		goto L1
	} else {
		goto L89
	}
L89:
	;
	F_errsave_finish(m, v22, int32(_a_F_record_in_1), int32(218), int32(_a_F_record_in_2))
	mBase = m.M
	v312 = m.ExcPending
	if v312 != 0 {
		goto L1
	} else {
		goto L90
	}
L90:
	;
	goto L35
L91:
	;
	v316 = F_errsave_start(m, v22)
	mBase = m.M
	v317 = m.ExcPending
	if v317 != 0 {
		goto L1
	} else {
		goto L94
	}
L92:
	;
	goto L93
L93:
	;
	v353 = v265 + int32(2)
	v354 = v313
	v355 = v271
	goto L77
L94:
	;
	if v316 == int32(0) {
		goto L35
	} else {
		goto L95
	}
L95:
	;
	F_errcode(m, int32(33685634))
	mBase = m.M
	v322 = m.ExcPending
	if v322 != 0 {
		goto L1
	} else {
		goto L96
	}
L96:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v19)+64)) = v24
	F_errmsg(m, int32(_a_F_record_in_3), v19-int32(-64))
	mBase = m.M
	v328 = m.ExcPending
	if v328 != 0 {
		goto L1
	} else {
		goto L97
	}
L97:
	;
	F_errdetail(m, int32(_a_F_record_in_6), int32(0))
	mBase = m.M
	v332 = m.ExcPending
	if v332 != 0 {
		goto L1
	} else {
		goto L98
	}
L98:
	;
	F_errsave_finish(m, v22, int32(_a_F_record_in_1), int32(229), int32(_a_F_record_in_2))
	mBase = m.M
	v337 = m.ExcPending
	if v337 != 0 {
		goto L1
	} else {
		goto L99
	}
L99:
	;
	goto L35
L100:
	;
	v346 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v286))))
	if v346 != int32(34) {
		v265 = v286
		v271 = int32(0)
		goto L72
	} else {
		goto L101
	}
L101:
	;
	v353 = v265 + int32(2)
	v354 = int32(34)
	v355 = int32(1)
	goto L77
L102:
	;
	v265 = v353
	v271 = v355
	goto L72
L103:
	;
	F_getTypeInputInfo(m, v221, v383+int32(4), v383+int32(8))
	mBase = m.M
	v391 = m.ExcPending
	if v391 != 0 {
		goto L1
	} else {
		goto L106
	}
L104:
	;
	goto L105
L105:
	;
	v403 = *(*int32)(unsafe.Add(mBase, uint32(v383)+8))
	v406 = *(*int32)(unsafe.Add(mBase, uint32(v209+int32(20))+76))
	v410 = F_InputFunctionCallSafe(m, v383+int32(16), v370, v403, v406, v22, v126+v190<<(uint(int32(2))%32))
	mBase = m.M
	v411 = m.ExcPending
	if v411 != 0 {
		goto L1
	} else {
		goto L108
	}
L106:
	;
	v392 = *(*int32)(unsafe.Add(mBase, uint32(v383)+4))
	v395 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v396 = *(*int32)(unsafe.Add(mBase, uint32(v395)+20))
	F_fmgr_info_cxt(m, v392, v383+int32(16), v396)
	mBase = m.M
	v398 = m.ExcPending
	if v398 != 0 {
		goto L1
	} else {
		goto L107
	}
L107:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v383))) = v221
	goto L105
L108:
	;
	if v410 == int32(0) {
		goto L35
	} else {
		goto L109
	}
L109:
	;
	v415 = v364
	v416 = int32(1)
	goto L54
L110:
	;
	goto L53
L111:
	;
	v524 = F_heap_form_tuple(m, v53, v126, v128)
	mBase = m.M
	v525 = m.ExcPending
	if v525 != 0 {
		goto L1
	} else {
		goto L131
	}
L112:
	;
	v453 = v434
	goto L115
L113:
	;
	goto L114
L114:
	;
	v502 = F_errsave_start(m, v22)
	mBase = m.M
	v503 = m.ExcPending
	if v503 != 0 {
		goto L1
	} else {
		goto L125
	}
L115:
	;
	v470 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v453)+1)))
	if base.B2i32(base.Ui32(v470-int32(9)) < base.Ui32(int32(5)))|base.B2i32(v470 == int32(32)) != 0 {
		v453 = v453 + int32(1)
		goto L115
	} else {
		goto L117
	}
L116:
	;
	if v470 == int32(0) {
		goto L111
	} else {
		goto L118
	}
L117:
	;
	goto L116
L118:
	;
	v480 = F_errsave_start(m, v22)
	mBase = m.M
	v481 = m.ExcPending
	if v481 != 0 {
		goto L1
	} else {
		goto L119
	}
L119:
	;
	if v480 == int32(0) {
		goto L35
	} else {
		goto L120
	}
L120:
	;
	F_errcode(m, int32(33685634))
	mBase = m.M
	v486 = m.ExcPending
	if v486 != 0 {
		goto L1
	} else {
		goto L121
	}
L121:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v19)+16)) = v24
	F_errmsg(m, int32(_a_F_record_in_3), v19+int32(16))
	mBase = m.M
	v492 = m.ExcPending
	if v492 != 0 {
		goto L1
	} else {
		goto L122
	}
L122:
	;
	F_errdetail(m, int32(_a_F_record_in_7), int32(0))
	mBase = m.M
	v496 = m.ExcPending
	if v496 != 0 {
		goto L1
	} else {
		goto L123
	}
L123:
	;
	F_errsave_finish(m, v22, int32(_a_F_record_in_1), int32(297), int32(_a_F_record_in_2))
	mBase = m.M
	v501 = m.ExcPending
	if v501 != 0 {
		goto L1
	} else {
		goto L124
	}
L124:
	;
	goto L35
L125:
	;
	if v502 == int32(0) {
		goto L35
	} else {
		goto L126
	}
L126:
	;
	F_errcode(m, int32(33685634))
	mBase = m.M
	v508 = m.ExcPending
	if v508 != 0 {
		goto L1
	} else {
		goto L127
	}
L127:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v19)+32)) = v24
	F_errmsg(m, int32(_a_F_record_in_3), v19+int32(32))
	mBase = m.M
	v514 = m.ExcPending
	if v514 != 0 {
		goto L1
	} else {
		goto L128
	}
L128:
	;
	F_errdetail(m, int32(_a_F_record_in_8), int32(0))
	mBase = m.M
	v518 = m.ExcPending
	if v518 != 0 {
		goto L1
	} else {
		goto L129
	}
L129:
	;
	F_errsave_finish(m, v22, int32(_a_F_record_in_1), int32(286), int32(_a_F_record_in_2))
	mBase = m.M
	v523 = m.ExcPending
	if v523 != 0 {
		goto L1
	} else {
		goto L130
	}
L130:
	;
	goto L35
L131:
	;
	v526 = *(*int32)(unsafe.Add(mBase, uint32(v524)))
	v527 = F_palloc(m, v526)
	mBase = m.M
	v528 = m.ExcPending
	if v528 != 0 {
		goto L1
	} else {
		goto L132
	}
L132:
	;
	v529 = *(*int32)(unsafe.Add(mBase, uint32(v524)))
	if v529 != 0 {
		goto L133
	} else {
		goto L134
	}
L133:
	;
	v530 = *(*int32)(unsafe.Add(mBase, uint32(v524)+16))
	base.MemoryCopy(m, v527, v530, v529)
	goto L135
L134:
	;
	goto L135
L135:
	;
	F_pfree(m, v524)
	mBase = m.M
	v533 = m.ExcPending
	if v533 != 0 {
		goto L1
	} else {
		goto L136
	}
L136:
	;
	v534 = *(*int32)(unsafe.Add(mBase, uint32(v19)+96))
	F_pfree(m, v534)
	mBase = m.M
	v536 = m.ExcPending
	if v536 != 0 {
		goto L1
	} else {
		goto L137
	}
L137:
	;
	F_pfree(m, v126)
	mBase = m.M
	v538 = m.ExcPending
	if v538 != 0 {
		goto L1
	} else {
		goto L138
	}
L138:
	;
	F_pfree(m, v128)
	mBase = m.M
	v540 = m.ExcPending
	if v540 != 0 {
		goto L1
	} else {
		goto L139
	}
L139:
	;
	v541 = *(*int32)(unsafe.Add(mBase, uint32(v53)+12))
	if int32(0) <= v541 {
		goto L140
	} else {
		goto L141
	}
L140:
	;
	F_DecrTupleDescRefCount(m, v53)
	mBase = m.M
	v545 = m.ExcPending
	if v545 != 0 {
		goto L1
	} else {
		goto L143
	}
L141:
	;
	goto L142
L142:
	;
	v546 = F_HeapTupleHeaderGetDatum(m, v527)
	mBase = m.M
	v547 = m.ExcPending
	if v547 != 0 {
		goto L1
	} else {
		goto L144
	}
L143:
	;
	goto L142
L144:
	;
	v573 = v546
	goto L3
L145:
	;
	F_DecrTupleDescRefCount(m, v53)
	mBase = m.M
	v569 = m.ExcPending
	if v569 != 0 {
		goto L1
	} else {
		goto L148
	}
L146:
	;
	goto L147
L147:
	;
	v570 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(l0)+16)) = uint8(v570)
	v573 = v564
	goto L3
L148:
	;
	goto L147
}
