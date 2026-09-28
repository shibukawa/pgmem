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
func F_in_range_float4_float8(m *base.Module, l0 int32) int64 {
	mBase := m.M
	_ = mBase
	var v10 int64
	_ = v10
	var v15 float64
	_ = v15
	var v18 int64
	_ = v18
	var v19 int32
	_ = v19
	var v20 int32
	_ = v20
	var v21 int32
	_ = v21
	var v22 int32
	_ = v22
	var v40 int64
	_ = v40
	var v41 float32
	_ = v41
	var v66 float64
	_ = v66
	var v68 float64
	_ = v68
	var v70 float64
	_ = v70
	var v85 int32
	_ = v85
	var v88 int32
	_ = v88
	var v92 int32
	_ = v92
	var v97 int32
	_ = v97
	v10 = *(*int64)(unsafe.Add(mBase, uint32(l0)+56))
	if base.Ui64(int64(9218868437227405312)) < base.Ui64(v10&int64(9223372036854775807)) {
		F_errstart_cold(m, int32(21), int32(0))
		mBase = m.M
		v85 = m.ExcPending
		if v85 != 0 {
			return int64(0)
		} else {
			F_errcode(m, int32(50593922))
			mBase = m.M
			v88 = m.ExcPending
			if v88 != 0 {
				return int64(0)
			} else {
				F_errmsg(m, int32(_a_F_in_range_float4_float8_0), int32(0))
				mBase = m.M
				v92 = m.ExcPending
				if v92 != 0 {
					return int64(0)
				} else {
					F_errfinish(m, int32(_a_F_in_range_float4_float8_1), int32(1160), int32(_a_F_in_range_float4_float8_2))
					mBase = m.M
					v97 = m.ExcPending
					if v97 != 0 {
						return int64(0)
					} else {
						base.Wasm_trap_unreachable()
						for {
						}
					}
				}
			}
		}
	} else {
		v15 = base.F64_reinterpret_i64(v10)
		if base.F64_lt(v15, float64(0)) != 0 {
			F_errstart_cold(m, int32(21), int32(0))
			mBase = m.M
			v85 = m.ExcPending
			if v85 != 0 {
				return int64(0)
			} else {
				F_errcode(m, int32(50593922))
				mBase = m.M
				v88 = m.ExcPending
				if v88 != 0 {
					return int64(0)
				} else {
					F_errmsg(m, int32(_a_F_in_range_float4_float8_0), int32(0))
					mBase = m.M
					v92 = m.ExcPending
					if v92 != 0 {
						return int64(0)
					} else {
						F_errfinish(m, int32(_a_F_in_range_float4_float8_1), int32(1160), int32(_a_F_in_range_float4_float8_2))
						mBase = m.M
						v97 = m.ExcPending
						if v97 != 0 {
							return int64(0)
						} else {
							base.Wasm_trap_unreachable()
							for {
							}
						}
					}
				}
			}
		} else {
			v18 = *(*int64)(unsafe.Add(mBase, uint32(l0)+88))
			v19 = *(*int32)(unsafe.Add(mBase, uint32(l0)+40))
			v20 = int32(2147483647)
			v21 = v19 & v20
			v22 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
			if base.Ui32(int32(2139095041)) <= base.Ui32(v22&v20) {
				return base.I64_extend_i32_u(base.B2i32(v18 == int64(0)) | base.B2i32(base.Ui32(int32(2139095040)) < base.Ui32(v21)))
			} else {
				if base.Ui32(int32(2139095040)) < base.Ui32(v21) {
					return base.I64_extend_i32_u(base.B2i32(v18 != int64(0)))
				} else {
					v40 = *(*int64)(unsafe.Add(mBase, uint32(l0)+72))
					v41 = base.F32_reinterpret_i32(v19)
					if base.F32_ne(base.F32_abs(v41), math.Float32frombits(uint32(0x7f800000)))|base.F64_ne(base.F64_abs(v15), math.Float64frombits(uint64(0x7ff0000000000000))) != 0 {
						if v40 == int64(0) {
							v66 = v15
						} else {
							v66 = base.F64_neg(v15)
						}
						v68 = base.F64_add(v66, base.F64_promote_f32(v41))
						v70 = base.F64_promote_f32(base.F32_reinterpret_i32(v22))
						if v18 != int64(0) {
							return base.I64_extend_i32_u(base.F64_ge(v68, v70))
						} else {
							return base.I64_extend_i32_u(base.F64_le(v68, v70))
						}
					} else {
						if v40 != int64(0) {
							if base.F32_gt(v41, float32(0)) == int32(0) {
								if v40 == int64(0) {
									v66 = v15
								} else {
									v66 = base.F64_neg(v15)
								}
								v68 = base.F64_add(v66, base.F64_promote_f32(v41))
								v70 = base.F64_promote_f32(base.F32_reinterpret_i32(v22))
								if v18 != int64(0) {
									return base.I64_extend_i32_u(base.F64_ge(v68, v70))
								} else {
									return base.I64_extend_i32_u(base.F64_le(v68, v70))
								}
							} else {
								return int64(1)
							}
						} else {
							if base.F32_lt(v41, float32(0)) == int32(0) {
								if v40 == int64(0) {
									v66 = v15
								} else {
									v66 = base.F64_neg(v15)
								}
								v68 = base.F64_add(v66, base.F64_promote_f32(v41))
								v70 = base.F64_promote_f32(base.F32_reinterpret_i32(v22))
								if v18 != int64(0) {
									return base.I64_extend_i32_u(base.F64_ge(v68, v70))
								} else {
									return base.I64_extend_i32_u(base.F64_le(v68, v70))
								}
							} else {
								return int64(1)
							}
						}
					}
				}
			}
		}
	}
}
func F_in_range_int2_int4(m *base.Module, l0 int32) int64 {
	mBase := m.M
	_ = mBase
	var v6 int32
	_ = v6
	var v9 int64
	_ = v9
	var v12 int64
	_ = v12
	var v14 int32
	_ = v14
	var v15 int32
	_ = v15
	var v18 int32
	_ = v18
	var v19 int32
	_ = v19
	var v27 int32
	_ = v27
	var v41 int32
	_ = v41
	var v44 int32
	_ = v44
	var v48 int32
	_ = v48
	var v53 int32
	_ = v53
	v6 = *(*int32)(unsafe.Add(mBase, uint32(l0)+56))
	if int32(0) <= v6 {
		v9 = *(*int64)(unsafe.Add(mBase, uint32(l0)+88))
		v12 = *(*int64)(unsafe.Add(mBase, uint32(l0)+72))
		v14 = base.B2i32(v12 != int64(0))
		if v12 != int64(0) {
			v15 = int32(0) - v6
		} else {
			v15 = v6
		}
		v18 = int32(*(*int16)(unsafe.Add(mBase, uint32(l0)+40)))
		v19 = v18 + v15
		if base.B2i32(v15 < int32(0)) != base.B2i32(v19 < v18) {
			return base.I64_extend_i32_u(v14 ^ base.B2i32(v9 != int64(0)))
		} else {
			v27 = int32(*(*int16)(unsafe.Add(mBase, uint32(l0)+24)))
			if v9 != int64(0) {
				return base.I64_extend_i32_u(base.B2i32(v27 <= v19))
			} else {
				return base.I64_extend_i32_u(base.B2i32(v19 <= v27))
			}
		}
	} else {
		F_errstart_cold(m, int32(21), int32(0))
		mBase = m.M
		v41 = m.ExcPending
		if v41 != 0 {
			return int64(0)
		} else {
			F_errcode(m, int32(50593922))
			mBase = m.M
			v44 = m.ExcPending
			if v44 != 0 {
				return int64(0)
			} else {
				F_errmsg(m, int32(_a_F_in_range_int2_int4_0), int32(0))
				mBase = m.M
				v48 = m.ExcPending
				if v48 != 0 {
					return int64(0)
				} else {
					F_errfinish(m, int32(_a_F_in_range_int2_int4_1), int32(746), int32(_a_F_in_range_int2_int4_2))
					mBase = m.M
					v53 = m.ExcPending
					if v53 != 0 {
						return int64(0)
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
func F_in_range_int2_int8(m *base.Module, l0 int32) int64 {
	mBase := m.M
	_ = mBase
	var v4 int64
	_ = v4
	var v5 int64
	_ = v5
	var v6 int64
	_ = v6
	var v7 int64
	_ = v7
	var v8 int64
	_ = v8
	var v9 int64
	_ = v9
	var v12 int32
	_ = v12
	v4 = int64(*(*int16)(unsafe.Add(mBase, uint32(l0)+24)))
	v5 = int64(*(*int16)(unsafe.Add(mBase, uint32(l0)+40)))
	v6 = *(*int64)(unsafe.Add(mBase, uint32(l0)+56))
	v7 = *(*int64)(unsafe.Add(mBase, uint32(l0)+72))
	v8 = *(*int64)(unsafe.Add(mBase, uint32(l0)+88))
	v9 = F_DirectFunctionCall5Coll(m, int32(1445), int32(0), v4, v5, v6, v7, v8)
	mBase = m.M
	v12 = m.ExcPending
	if v12 != 0 {
		return int64(0)
	} else {
		return v9
	}
}
func F_record_in(m *base.Module, l0 int32) int64 {
	mBase := m.M
	_ = mBase
	var v17 int64
	_ = v17
	var v18 int32
	_ = v18
	var v20 int32
	_ = v20
	var v22 int32
	_ = v22
	var v23 int32
	_ = v23
	var v24 int32
	_ = v24
	var v25 int32
	_ = v25
	var v29 int32
	_ = v29
	var v32 int32
	_ = v32
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
	var v125 int32
	_ = v125
	var v126 int32
	_ = v126
	var v128 int32
	_ = v128
	var v129 int32
	_ = v129
	var v131 int32
	_ = v131
	var v147 int32
	_ = v147
	var v156 int32
	_ = v156
	var v157 int32
	_ = v157
	var v162 int32
	_ = v162
	var v166 int32
	_ = v166
	var v169 int32
	_ = v169
	var v170 int32
	_ = v170
	var v175 int32
	_ = v175
	var v177 int32
	_ = v177
	var v181 int32
	_ = v181
	var v186 int32
	_ = v186
	var v189 int32
	_ = v189
	var v190 int32
	_ = v190
	var v191 int32
	_ = v191
	var v205 int32
	_ = v205
	var v211 int32
	_ = v211
	var v212 int32
	_ = v212
	var v221 int32
	_ = v221
	var v223 int32
	_ = v223
	var v224 int32
	_ = v224
	var v229 int32
	_ = v229
	var v230 int32
	_ = v230
	var v232 int32
	_ = v232
	var v235 int32
	_ = v235
	var v236 int32
	_ = v236
	var v241 int32
	_ = v241
	var v247 int32
	_ = v247
	var v250 int32
	_ = v250
	var v251 int32
	_ = v251
	var v256 int32
	_ = v256
	var v258 int32
	_ = v258
	var v259 int32
	_ = v259
	var v260 int32
	_ = v260
	var v267 int32
	_ = v267
	var v273 int32
	_ = v273
	var v283 int32
	_ = v283
	var v289 int32
	_ = v289
	var v294 int32
	_ = v294
	var v295 int32
	_ = v295
	var v300 int32
	_ = v300
	var v306 int32
	_ = v306
	var v309 int32
	_ = v309
	var v310 int32
	_ = v310
	var v315 int32
	_ = v315
	var v316 int32
	_ = v316
	var v319 int32
	_ = v319
	var v320 int32
	_ = v320
	var v325 int32
	_ = v325
	var v331 int32
	_ = v331
	var v334 int32
	_ = v334
	var v335 int32
	_ = v335
	var v340 int32
	_ = v340
	var v343 int32
	_ = v343
	var v349 int32
	_ = v349
	var v356 int32
	_ = v356
	var v357 int32
	_ = v357
	var v358 int32
	_ = v358
	var v363 int32
	_ = v363
	var v365 int32
	_ = v365
	var v367 int32
	_ = v367
	var v368 int32
	_ = v368
	var v373 int32
	_ = v373
	var v387 int32
	_ = v387
	var v388 int32
	_ = v388
	var v395 int32
	_ = v395
	var v396 int32
	_ = v396
	var v399 int32
	_ = v399
	var v400 int32
	_ = v400
	var v402 int32
	_ = v402
	var v407 int32
	_ = v407
	var v410 int32
	_ = v410
	var v414 int32
	_ = v414
	var v415 int32
	_ = v415
	var v419 int32
	_ = v419
	var v420 int32
	_ = v420
	var v436 int32
	_ = v436
	var v439 int32
	_ = v439
	var v455 int32
	_ = v455
	var v459 int32
	_ = v459
	var v477 int32
	_ = v477
	var v487 int32
	_ = v487
	var v488 int32
	_ = v488
	var v493 int32
	_ = v493
	var v499 int32
	_ = v499
	var v502 int32
	_ = v502
	var v503 int32
	_ = v503
	var v508 int32
	_ = v508
	var v509 int32
	_ = v509
	var v510 int32
	_ = v510
	var v515 int32
	_ = v515
	var v521 int32
	_ = v521
	var v524 int32
	_ = v524
	var v525 int32
	_ = v525
	var v530 int32
	_ = v530
	var v531 int32
	_ = v531
	var v532 int32
	_ = v532
	var v533 int32
	_ = v533
	var v534 int32
	_ = v534
	var v535 int32
	_ = v535
	var v536 int32
	_ = v536
	var v537 int32
	_ = v537
	var v540 int32
	_ = v540
	var v541 int32
	_ = v541
	var v543 int32
	_ = v543
	var v545 int32
	_ = v545
	var v547 int32
	_ = v547
	var v548 int32
	_ = v548
	var v552 int32
	_ = v552
	var v553 int64
	_ = v553
	var v554 int32
	_ = v554
	var v572 int32
	_ = v572
	var v576 int32
	_ = v576
	var v577 int32
	_ = v577
	var v595 int64
	_ = v595
	v17 = int64(0)
	v18 = m.G0
	v20 = v18 - int32(112)
	m.G0 = v20
	v22 = *(*int32)(unsafe.Add(mBase, uint32(l0)+40))
	v23 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v24 = *(*int32)(unsafe.Add(mBase, uint32(l0)+56))
	v25 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
	F_check_stack_depth(m)
	mBase = m.M
	v29 = m.ExcPending
	if v29 != 0 {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	return int64(0)
L2:
	;
	v32 = int32(0)
	if base.B2i32(v22 != int32(2249))|base.B2i32(v32 <= v24) == v32 {
		goto L4
	} else {
		goto L5
	}
L3:
	;
	m.G0 = v20 + int32(112)
	return v595
L4:
	;
	v37 = F_errsave_start(m, v23)
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
	v53 = F_lookup_rowtype_tupdesc(m, v22, v24)
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
		v595 = v17
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
	F_errsave_finish(m, v23, int32(_a_F_record_in_1), int32(105), int32(_a_F_record_in_2))
	mBase = m.M
	v52 = m.ExcPending
	if v52 != 0 {
		goto L1
	} else {
		goto L11
	}
L11:
	;
	v595 = v17
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
	v125 = F_palloc_mul(m, int32(8), v55)
	mBase = m.M
	v126 = m.ExcPending
	if v126 != 0 {
		goto L1
	} else {
		goto L33
	}
L14:
	;
	if v78 == v22 {
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
	if v80 == v24 {
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
	*(*int32)(unsafe.Add(mBase, uint32(v77)+4)) = v24
	*(*int32)(unsafe.Add(mBase, uint32(v77))) = v22
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
	v128 = F_palloc_mul(m, int32(1), v55)
	mBase = m.M
	v129 = m.ExcPending
	if v129 != 0 {
		goto L1
	} else {
		goto L34
	}
L34:
	;
	v131 = v25
	goto L37
L35:
	;
	v572 = *(*int32)(unsafe.Add(mBase, uint32(v53)+12))
	if int32(0) <= v572 {
		goto L145
	} else {
		goto L146
	}
L36:
	;
	v177 = v131 + int32(1)
	F_initStringInfo(m, v20+int32(96))
	mBase = m.M
	v181 = m.ExcPending
	if v181 != 0 {
		goto L1
	} else {
		goto L48
	}
L37:
	;
	v147 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v131))))
	if base.Ui32(v147-int32(9)) < base.Ui32(int32(5)) {
		goto L40
	} else {
		goto L41
	}
L38:
	;
	v156 = F_errsave_start(m, v23)
	mBase = m.M
	v157 = m.ExcPending
	if v157 != 0 {
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
	switch v147 - int32(32) {
	case 0:
		goto L40
	default:
		goto L39
	case 8:
		goto L36
	}
L42:
	;
	if v156 == int32(0) {
		goto L35
	} else {
		goto L43
	}
L43:
	;
	F_errcode(m, int32(33685634))
	mBase = m.M
	v162 = m.ExcPending
	if v162 != 0 {
		goto L1
	} else {
		goto L44
	}
L44:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v20))) = v25
	F_errmsg(m, int32(_a_F_record_in_3), v20)
	mBase = m.M
	v166 = m.ExcPending
	if v166 != 0 {
		goto L1
	} else {
		goto L45
	}
L45:
	;
	v169 = F_errdetail(m, int32(_a_F_record_in_4), int32(0))
	mBase = m.M
	v170 = m.ExcPending
	if v170 != 0 {
		goto L1
	} else {
		goto L46
	}
L46:
	;
	F_errsave_finish(m, v23, int32(_a_F_record_in_1), int32(159), int32(_a_F_record_in_2))
	mBase = m.M
	v175 = m.ExcPending
	if v175 != 0 {
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
	v186 = int32(0)
	v189 = v177
	v190 = v186
	v191 = v186
	goto L52
L50:
	;
	v439 = v177
	goto L51
L51:
	;
	v455 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v439))))
	if v455 == int32(41) {
		goto L112
	} else {
		goto L113
	}
L52:
	;
	v205 = *(*int32)(unsafe.Add(mBase, uint32(v53)))
	v211 = v53 + v205<<(uint(int32(3))%32) + v191*int32(100)
	v212 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v211)+119)))
	if v212 == int32(1) {
		goto L55
	} else {
		goto L56
	}
L53:
	;
	v439 = v419
	goto L51
L54:
	;
	v436 = v191 + int32(1)
	if v436 != v55 {
		v189 = v419
		v190 = v420
		v191 = v436
		goto L52
	} else {
		goto L110
	}
L55:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v125+v191<<(uint(int32(3))%32)))) = int64(0)
	v221 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(v191+v128))) = uint8(v221)
	v419 = v189
	v420 = v190
	goto L54
L56:
	;
	goto L57
L57:
	;
	v223 = *(*int32)(unsafe.Add(mBase, uint32(v211)+96))
	if v190 != 0 {
		goto L61
	} else {
		goto L62
	}
L58:
	;
	*(*uint8)(unsafe.Add(mBase, uint32(v191+v128))) = uint8(v368)
	v387 = v77 + int32(12) + v191*int32(44)
	v388 = *(*int32)(unsafe.Add(mBase, uint32(v387)))
	if v223 != v388 {
		goto L103
	} else {
		goto L104
	}
L59:
	;
	v258 = v20 + int32(96)
	v259 = *(*int32)(unsafe.Add(mBase, uint32(v258)))
	v260 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v259))) = uint8(v260)
	*(*int32)(unsafe.Add(mBase, uint32(v258)+12)) = v260
	*(*int32)(unsafe.Add(mBase, uint32(v258)+4)) = v260
	goto L71
L60:
	;
	v235 = F_errsave_start(m, v23)
	mBase = m.M
	v236 = m.ExcPending
	if v236 != 0 {
		goto L1
	} else {
		goto L65
	}
L61:
	;
	v224 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v189))))
	if v224 != int32(44) {
		goto L60
	} else {
		goto L64
	}
L62:
	;
	v229 = v189
	goto L63
L63:
	;
	v230 = int32(0)
	v232 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v229))))
	switch v232 - int32(41) {
	case 0, 3:
		v367 = v229
		v368 = int32(1)
		v373 = v230
		goto L58
	default:
		goto L59
	}
L64:
	;
	v229 = v189 + int32(1)
	goto L63
L65:
	;
	if v235 == int32(0) {
		goto L35
	} else {
		goto L66
	}
L66:
	;
	F_errcode(m, int32(33685634))
	mBase = m.M
	v241 = m.ExcPending
	if v241 != 0 {
		goto L1
	} else {
		goto L67
	}
L67:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v20)+80)) = v25
	F_errmsg(m, int32(_a_F_record_in_3), v20+int32(80))
	mBase = m.M
	v247 = m.ExcPending
	if v247 != 0 {
		goto L1
	} else {
		goto L68
	}
L68:
	;
	v250 = F_errdetail(m, int32(_a_F_record_in_5), int32(0))
	mBase = m.M
	v251 = m.ExcPending
	if v251 != 0 {
		goto L1
	} else {
		goto L69
	}
L69:
	;
	F_errsave_finish(m, v23, int32(_a_F_record_in_1), int32(191), int32(_a_F_record_in_2))
	mBase = m.M
	v256 = m.ExcPending
	if v256 != 0 {
		goto L1
	} else {
		goto L70
	}
L70:
	;
	goto L35
L71:
	;
	v267 = v229
	v273 = v230
	goto L72
L72:
	;
	v283 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v267))))
	if v273&int32(1) != 0 {
		goto L75
	} else {
		goto L76
	}
L73:
	;
	v365 = *(*int32)(unsafe.Add(mBase, uint32(v20)+96))
	v367 = v267
	v368 = int32(0)
	v373 = v365
	goto L58
L74:
	;
	goto L73
L75:
	;
	v289 = v267 + int32(1)
	if v283 != int32(34) {
		goto L78
	} else {
		goto L79
	}
L76:
	;
	switch v283 - int32(41) {
	case 0, 3:
		goto L74
	default:
		goto L75
	}
L77:
	;
	F_appendStringInfoChar(m, v20+int32(96), base.I32_extend8_s(v357))
	mBase = m.M
	v363 = m.ExcPending
	if v363 != 0 {
		goto L1
	} else {
		goto L102
	}
L78:
	;
	if v283 != int32(92) {
		goto L81
	} else {
		goto L82
	}
L79:
	;
	goto L80
L80:
	;
	v343 = int32(1)
	if v273&v343 == int32(0) {
		v267 = v289
		v273 = v343
		goto L72
	} else {
		goto L100
	}
L81:
	;
	if v283 != 0 {
		v356 = v289
		v357 = v283
		v358 = v273
		goto L77
	} else {
		goto L84
	}
L82:
	;
	goto L83
L83:
	;
	v316 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v289))))
	if v316 == int32(0) {
		goto L91
	} else {
		goto L92
	}
L84:
	;
	v294 = F_errsave_start(m, v23)
	mBase = m.M
	v295 = m.ExcPending
	if v295 != 0 {
		goto L1
	} else {
		goto L85
	}
L85:
	;
	if v294 == int32(0) {
		goto L35
	} else {
		goto L86
	}
L86:
	;
	F_errcode(m, int32(33685634))
	mBase = m.M
	v300 = m.ExcPending
	if v300 != 0 {
		goto L1
	} else {
		goto L87
	}
L87:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v20)+48)) = v25
	F_errmsg(m, int32(_a_F_record_in_3), v20+int32(48))
	mBase = m.M
	v306 = m.ExcPending
	if v306 != 0 {
		goto L1
	} else {
		goto L88
	}
L88:
	;
	v309 = F_errdetail(m, int32(_a_F_record_in_6), int32(0))
	mBase = m.M
	v310 = m.ExcPending
	if v310 != 0 {
		goto L1
	} else {
		goto L89
	}
L89:
	;
	F_errsave_finish(m, v23, int32(_a_F_record_in_1), int32(218), int32(_a_F_record_in_2))
	mBase = m.M
	v315 = m.ExcPending
	if v315 != 0 {
		goto L1
	} else {
		goto L90
	}
L90:
	;
	goto L35
L91:
	;
	v319 = F_errsave_start(m, v23)
	mBase = m.M
	v320 = m.ExcPending
	if v320 != 0 {
		goto L1
	} else {
		goto L94
	}
L92:
	;
	goto L93
L93:
	;
	v356 = v267 + int32(2)
	v357 = v316
	v358 = v273
	goto L77
L94:
	;
	if v319 == int32(0) {
		goto L35
	} else {
		goto L95
	}
L95:
	;
	F_errcode(m, int32(33685634))
	mBase = m.M
	v325 = m.ExcPending
	if v325 != 0 {
		goto L1
	} else {
		goto L96
	}
L96:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v20)+64)) = v25
	F_errmsg(m, int32(_a_F_record_in_3), v20-int32(-64))
	mBase = m.M
	v331 = m.ExcPending
	if v331 != 0 {
		goto L1
	} else {
		goto L97
	}
L97:
	;
	v334 = F_errdetail(m, int32(_a_F_record_in_6), int32(0))
	mBase = m.M
	v335 = m.ExcPending
	if v335 != 0 {
		goto L1
	} else {
		goto L98
	}
L98:
	;
	F_errsave_finish(m, v23, int32(_a_F_record_in_1), int32(229), int32(_a_F_record_in_2))
	mBase = m.M
	v340 = m.ExcPending
	if v340 != 0 {
		goto L1
	} else {
		goto L99
	}
L99:
	;
	goto L35
L100:
	;
	v349 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v289))))
	if v349 != int32(34) {
		v267 = v289
		v273 = int32(0)
		goto L72
	} else {
		goto L101
	}
L101:
	;
	v356 = v267 + int32(2)
	v357 = int32(34)
	v358 = int32(1)
	goto L77
L102:
	;
	v267 = v356
	v273 = v358
	goto L72
L103:
	;
	F_getTypeInputInfo(m, v223, v387+int32(4), v387+int32(8))
	mBase = m.M
	v395 = m.ExcPending
	if v395 != 0 {
		goto L1
	} else {
		goto L106
	}
L104:
	;
	goto L105
L105:
	;
	v407 = *(*int32)(unsafe.Add(mBase, uint32(v387)+8))
	v410 = *(*int32)(unsafe.Add(mBase, uint32(v211+int32(28))+76))
	v414 = F_InputFunctionCallSafe(m, v387+int32(16), v373, v407, v410, v23, v125+v191<<(uint(int32(3))%32))
	mBase = m.M
	v415 = m.ExcPending
	if v415 != 0 {
		goto L1
	} else {
		goto L108
	}
L106:
	;
	v396 = *(*int32)(unsafe.Add(mBase, uint32(v387)+4))
	v399 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v400 = *(*int32)(unsafe.Add(mBase, uint32(v399)+20))
	F_fmgr_info_cxt(m, v396, v387+int32(16), v400)
	mBase = m.M
	v402 = m.ExcPending
	if v402 != 0 {
		goto L1
	} else {
		goto L107
	}
L107:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v387))) = v223
	goto L105
L108:
	;
	if v414 == int32(0) {
		goto L35
	} else {
		goto L109
	}
L109:
	;
	v419 = v367
	v420 = int32(1)
	goto L54
L110:
	;
	goto L53
L111:
	;
	v531 = F_heap_form_tuple(m, v53, v125, v128)
	mBase = m.M
	v532 = m.ExcPending
	if v532 != 0 {
		goto L1
	} else {
		goto L131
	}
L112:
	;
	v459 = v439
	goto L115
L113:
	;
	goto L114
L114:
	;
	v509 = F_errsave_start(m, v23)
	mBase = m.M
	v510 = m.ExcPending
	if v510 != 0 {
		goto L1
	} else {
		goto L125
	}
L115:
	;
	v477 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v459)+1)))
	if base.B2i32(base.Ui32(v477-int32(9)) < base.Ui32(int32(5)))|base.B2i32(v477 == int32(32)) != 0 {
		v459 = v459 + int32(1)
		goto L115
	} else {
		goto L117
	}
L116:
	;
	if v477 == int32(0) {
		goto L111
	} else {
		goto L118
	}
L117:
	;
	goto L116
L118:
	;
	v487 = F_errsave_start(m, v23)
	mBase = m.M
	v488 = m.ExcPending
	if v488 != 0 {
		goto L1
	} else {
		goto L119
	}
L119:
	;
	if v487 == int32(0) {
		goto L35
	} else {
		goto L120
	}
L120:
	;
	F_errcode(m, int32(33685634))
	mBase = m.M
	v493 = m.ExcPending
	if v493 != 0 {
		goto L1
	} else {
		goto L121
	}
L121:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v20)+16)) = v25
	F_errmsg(m, int32(_a_F_record_in_3), v20+int32(16))
	mBase = m.M
	v499 = m.ExcPending
	if v499 != 0 {
		goto L1
	} else {
		goto L122
	}
L122:
	;
	v502 = F_errdetail(m, int32(_a_F_record_in_7), int32(0))
	mBase = m.M
	v503 = m.ExcPending
	if v503 != 0 {
		goto L1
	} else {
		goto L123
	}
L123:
	;
	F_errsave_finish(m, v23, int32(_a_F_record_in_1), int32(297), int32(_a_F_record_in_2))
	mBase = m.M
	v508 = m.ExcPending
	if v508 != 0 {
		goto L1
	} else {
		goto L124
	}
L124:
	;
	goto L35
L125:
	;
	if v509 == int32(0) {
		goto L35
	} else {
		goto L126
	}
L126:
	;
	F_errcode(m, int32(33685634))
	mBase = m.M
	v515 = m.ExcPending
	if v515 != 0 {
		goto L1
	} else {
		goto L127
	}
L127:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v20)+32)) = v25
	F_errmsg(m, int32(_a_F_record_in_3), v20+int32(32))
	mBase = m.M
	v521 = m.ExcPending
	if v521 != 0 {
		goto L1
	} else {
		goto L128
	}
L128:
	;
	v524 = F_errdetail(m, int32(_a_F_record_in_8), int32(0))
	mBase = m.M
	v525 = m.ExcPending
	if v525 != 0 {
		goto L1
	} else {
		goto L129
	}
L129:
	;
	F_errsave_finish(m, v23, int32(_a_F_record_in_1), int32(286), int32(_a_F_record_in_2))
	mBase = m.M
	v530 = m.ExcPending
	if v530 != 0 {
		goto L1
	} else {
		goto L130
	}
L130:
	;
	goto L35
L131:
	;
	v533 = *(*int32)(unsafe.Add(mBase, uint32(v531)))
	v534 = F_palloc(m, v533)
	mBase = m.M
	v535 = m.ExcPending
	if v535 != 0 {
		goto L1
	} else {
		goto L132
	}
L132:
	;
	v536 = *(*int32)(unsafe.Add(mBase, uint32(v531)))
	if v536 != 0 {
		goto L133
	} else {
		goto L134
	}
L133:
	;
	v537 = *(*int32)(unsafe.Add(mBase, uint32(v531)+16))
	base.MemoryCopy(m, v534, v537, v536)
	goto L135
L134:
	;
	goto L135
L135:
	;
	F_pfree(m, v531)
	mBase = m.M
	v540 = m.ExcPending
	if v540 != 0 {
		goto L1
	} else {
		goto L136
	}
L136:
	;
	v541 = *(*int32)(unsafe.Add(mBase, uint32(v20)+96))
	F_pfree(m, v541)
	mBase = m.M
	v543 = m.ExcPending
	if v543 != 0 {
		goto L1
	} else {
		goto L137
	}
L137:
	;
	F_pfree(m, v125)
	mBase = m.M
	v545 = m.ExcPending
	if v545 != 0 {
		goto L1
	} else {
		goto L138
	}
L138:
	;
	F_pfree(m, v128)
	mBase = m.M
	v547 = m.ExcPending
	if v547 != 0 {
		goto L1
	} else {
		goto L139
	}
L139:
	;
	v548 = *(*int32)(unsafe.Add(mBase, uint32(v53)+12))
	if int32(0) <= v548 {
		goto L140
	} else {
		goto L141
	}
L140:
	;
	F_DecrTupleDescRefCount(m, v53)
	mBase = m.M
	v552 = m.ExcPending
	if v552 != 0 {
		goto L1
	} else {
		goto L143
	}
L141:
	;
	goto L142
L142:
	;
	v553 = F_HeapTupleHeaderGetDatum(m, v534)
	mBase = m.M
	v554 = m.ExcPending
	if v554 != 0 {
		goto L1
	} else {
		goto L144
	}
L143:
	;
	goto L142
L144:
	;
	v595 = v553
	goto L3
L145:
	;
	F_DecrTupleDescRefCount(m, v53)
	mBase = m.M
	v576 = m.ExcPending
	if v576 != 0 {
		goto L1
	} else {
		goto L148
	}
L146:
	;
	goto L147
L147:
	;
	v577 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(l0)+16)) = uint8(v577)
	v595 = v17
	goto L3
L148:
	;
	goto L147
}
