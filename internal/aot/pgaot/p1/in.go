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
	v2 = int32(*(*uint8)(unsafe.Add(mBase, _consts[308])))
	return int32(base.Ui32(v2&int32(2)) >> (uint(int32(1)) % 32))
}
func F_IsInParallelMode(m *base.Module) int32 {
	mBase := m.M
	_ = mBase
	var v5 int32
	_ = v5
	var v6 int32
	_ = v6
	var v7 int32
	_ = v7
	var v8 int32
	_ = v8
	v5 = *(*int32)(unsafe.Add(mBase, _consts[4]))
	v6 = *(*int32)(unsafe.Add(mBase, uint32(v5)+72))
	if v6 != 0 {
		v8 = int32(1)
	} else {
		v7 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v5)+76)))
		v8 = v7
	}
	return v8 & int32(1)
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
	var v20 int32
	_ = v20
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
	var v51 int32
	_ = v51
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
	v20 = v10
	goto L5
L4:
	;
	return v51
L5:
	;
	if v20 == v13 {
		goto L7
	} else {
		goto L8
	}
L6:
	;
	v51 = int32(0)
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
	v29 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v27+v20))))
	if l3 < v29 {
		v51 = v26
		goto L4
	} else {
		goto L10
	}
L10:
	;
	v31 = v29 - l2
	if v31 < int32(0) {
		v51 = v26
		goto L4
	} else {
		goto L11
	}
L11:
	;
	v37 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l1+int32(base.Ui32(v31)>>(uint(int32(3))%32))))))
	if int32(base.Ui32(v37)>>(uint(v31&int32(7))%32))&int32(1) == int32(0) {
		v51 = v26
		goto L4
	} else {
		goto L12
	}
L12:
	;
	v46 = v20 + int32(1)
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v46
	if l4 != 0 {
		v20 = v46
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
	var v18 int32
	_ = v18
	var v19 int32
	_ = v19
	var v20 int32
	_ = v20
	var v21 int32
	_ = v21
	var v22 int32
	_ = v22
	var v38 int32
	_ = v38
	var v39 float32
	_ = v39
	var v59 float64
	_ = v59
	var v61 float64
	_ = v61
	var v63 float64
	_ = v63
	var v73 int32
	_ = v73
	var v76 int32
	_ = v76
	var v80 int32
	_ = v80
	var v85 int32
	_ = v85
	v9 = *(*int32)(unsafe.Add(mBase, uint32(l0)+36))
	v10 = *(*float64)(unsafe.Add(mBase, uint32(v9)))
	if base.Ui64(int64(9218868437227405312)) < base.Ui64(base.I64_reinterpret_f64(v10)&int64(9223372036854775807)) {
		F_errstart_cold(m, int32(21), int32(0))
		mBase = m.M
		v73 = m.ExcPending
		if v73 != 0 {
			return int32(0)
		} else {
			F_errcode(m, int32(50593922))
			mBase = m.M
			v76 = m.ExcPending
			if v76 != 0 {
				return int32(0)
			} else {
				F_errmsg(m, int32(242417), int32(0))
				mBase = m.M
				v80 = m.ExcPending
				if v80 != 0 {
					return int32(0)
				} else {
					F_errfinish(m, int32(476769), int32(1119), int32(530803))
					mBase = m.M
					v85 = m.ExcPending
					if v85 != 0 {
						return int32(0)
					} else {
						base.Wasm_trap_unreachable()
						for {
						}
					}
				}
			}
		}
	} else {
		if base.F64_lt(v10, float64(0)) != 0 {
			F_errstart_cold(m, int32(21), int32(0))
			mBase = m.M
			v73 = m.ExcPending
			if v73 != 0 {
				return int32(0)
			} else {
				F_errcode(m, int32(50593922))
				mBase = m.M
				v76 = m.ExcPending
				if v76 != 0 {
					return int32(0)
				} else {
					F_errmsg(m, int32(242417), int32(0))
					mBase = m.M
					v80 = m.ExcPending
					if v80 != 0 {
						return int32(0)
					} else {
						F_errfinish(m, int32(476769), int32(1119), int32(530803))
						mBase = m.M
						v85 = m.ExcPending
						if v85 != 0 {
							return int32(0)
						} else {
							base.Wasm_trap_unreachable()
							for {
							}
						}
					}
				}
			}
		} else {
			v18 = *(*int32)(unsafe.Add(mBase, uint32(l0)+52))
			v19 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
			v20 = int32(2147483647)
			v21 = v19 & v20
			v22 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
			if base.Ui32(int32(2139095041)) <= base.Ui32(v22&v20) {
				return base.B2i32(v18 == int32(0)) | base.B2i32(base.Ui32(int32(2139095040)) < base.Ui32(v21))
			} else {
				if base.Ui32(int32(2139095040)) < base.Ui32(v21) {
					return base.B2i32(v18 != int32(0))
				} else {
					v38 = *(*int32)(unsafe.Add(mBase, uint32(l0)+44))
					v39 = base.F32_reinterpret_i32(v19)
					if base.F64_ne(base.F64_abs(v10), math.Float64frombits(uint64(0x7ff0000000000000))) != 0 {
						if v38 != 0 {
							v59 = base.F64_neg(v10)
						} else {
							v59 = v10
						}
						v61 = base.F64_add(v59, base.F64_promote_f32(v39))
						v63 = base.F64_promote_f32(base.F32_reinterpret_i32(v22))
						if v18 != 0 {
							return base.F64_ge(v61, v63)
						} else {
							return base.F64_le(v61, v63)
						}
					} else {
						if base.F32_ne(base.F32_abs(v39), math.Float32frombits(uint32(0x7f800000))) != 0 {
							if v38 != 0 {
								v59 = base.F64_neg(v10)
							} else {
								v59 = v10
							}
							v61 = base.F64_add(v59, base.F64_promote_f32(v39))
							v63 = base.F64_promote_f32(base.F32_reinterpret_i32(v22))
							if v18 != 0 {
								return base.F64_ge(v61, v63)
							} else {
								return base.F64_le(v61, v63)
							}
						} else {
							if v38 != 0 {
								if base.F32_gt(v39, float32(0)) == int32(0) {
									if v38 != 0 {
										v59 = base.F64_neg(v10)
									} else {
										v59 = v10
									}
									v61 = base.F64_add(v59, base.F64_promote_f32(v39))
									v63 = base.F64_promote_f32(base.F32_reinterpret_i32(v22))
									if v18 != 0 {
										return base.F64_ge(v61, v63)
									} else {
										return base.F64_le(v61, v63)
									}
								} else {
									return int32(1)
								}
							} else {
								if base.F32_lt(v39, float32(0)) == int32(0) {
									if v38 != 0 {
										v59 = base.F64_neg(v10)
									} else {
										v59 = v10
									}
									v61 = base.F64_add(v59, base.F64_promote_f32(v39))
									v63 = base.F64_promote_f32(base.F32_reinterpret_i32(v22))
									if v18 != 0 {
										return base.F64_ge(v61, v63)
									} else {
										return base.F64_le(v61, v63)
									}
								} else {
									return int32(1)
								}
							}
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
				F_errmsg(m, int32(242417), int32(0))
				mBase = m.M
				v43 = m.ExcPending
				if v43 != 0 {
					return int32(0)
				} else {
					F_errfinish(m, int32(476339), int32(746), int32(533223))
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
	v9 = F_DirectFunctionCall5Coll(m, int32(1319), int32(0), v4, v5, v6, v7, v8)
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
	var v34 int32
	_ = v34
	var v35 int32
	_ = v35
	var v36 int32
	_ = v36
	var v41 int32
	_ = v41
	var v45 int32
	_ = v45
	var v50 int32
	_ = v50
	var v51 int32
	_ = v51
	var v52 int32
	_ = v52
	var v53 int32
	_ = v53
	var v54 int32
	_ = v54
	var v55 int32
	_ = v55
	var v58 int32
	_ = v58
	var v60 int32
	_ = v60
	var v61 int32
	_ = v61
	var v66 int32
	_ = v66
	var v67 int32
	_ = v67
	var v68 int32
	_ = v68
	var v70 int32
	_ = v70
	var v71 int32
	_ = v71
	var v75 int32
	_ = v75
	var v76 int32
	_ = v76
	var v78 int32
	_ = v78
	var v81 int32
	_ = v81
	var v83 int32
	_ = v83
	var v95 int32
	_ = v95
	var v97 int32
	_ = v97
	var v99 int32
	_ = v99
	var v106 int32
	_ = v106
	var v109 int32
	_ = v109
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
	var v140 int32
	_ = v140
	var v149 int32
	_ = v149
	var v150 int32
	_ = v150
	var v155 int32
	_ = v155
	var v159 int32
	_ = v159
	var v163 int32
	_ = v163
	var v168 int32
	_ = v168
	var v170 int32
	_ = v170
	var v174 int32
	_ = v174
	var v181 int32
	_ = v181
	var v184 int32
	_ = v184
	var v186 int32
	_ = v186
	var v189 int32
	_ = v189
	var v200 int32
	_ = v200
	var v206 int32
	_ = v206
	var v207 int32
	_ = v207
	var v216 int32
	_ = v216
	var v218 int32
	_ = v218
	var v221 int32
	_ = v221
	var v226 int32
	_ = v226
	var v227 int32
	_ = v227
	var v229 int32
	_ = v229
	var v232 int32
	_ = v232
	var v233 int32
	_ = v233
	var v238 int32
	_ = v238
	var v244 int32
	_ = v244
	var v248 int32
	_ = v248
	var v253 int32
	_ = v253
	var v255 int32
	_ = v255
	var v256 int32
	_ = v256
	var v257 int32
	_ = v257
	var v264 int32
	_ = v264
	var v270 int32
	_ = v270
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
	var v384 int32
	_ = v384
	var v385 int32
	_ = v385
	var v392 int32
	_ = v392
	var v393 int32
	_ = v393
	var v396 int32
	_ = v396
	var v397 int32
	_ = v397
	var v399 int32
	_ = v399
	var v404 int32
	_ = v404
	var v405 int32
	_ = v405
	var v409 int32
	_ = v409
	var v410 int32
	_ = v410
	var v414 int32
	_ = v414
	var v419 int32
	_ = v419
	var v431 int32
	_ = v431
	var v434 int32
	_ = v434
	var v450 int32
	_ = v450
	var v454 int32
	_ = v454
	var v471 int32
	_ = v471
	var v472 int32
	_ = v472
	var v481 int32
	_ = v481
	var v482 int32
	_ = v482
	var v487 int32
	_ = v487
	var v493 int32
	_ = v493
	var v497 int32
	_ = v497
	var v502 int32
	_ = v502
	var v503 int32
	_ = v503
	var v504 int32
	_ = v504
	var v509 int32
	_ = v509
	var v515 int32
	_ = v515
	var v519 int32
	_ = v519
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
	var v531 int32
	_ = v531
	var v532 int32
	_ = v532
	var v533 int32
	_ = v533
	var v535 int32
	_ = v535
	var v536 int32
	_ = v536
	var v538 int32
	_ = v538
	var v540 int32
	_ = v540
	var v542 int32
	_ = v542
	var v543 int32
	_ = v543
	var v547 int32
	_ = v547
	var v548 int32
	_ = v548
	var v549 int32
	_ = v549
	var v567 int32
	_ = v567
	var v568 int32
	_ = v568
	var v572 int32
	_ = v572
	var v573 int32
	_ = v573
	var v576 int32
	_ = v576
	v18 = m.G0
	v20 = v18 - int32(112)
	m.G0 = v20
	v22 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
	v23 = *(*int32)(unsafe.Add(mBase, uint32(l0)+36))
	v24 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v25 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
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
	return int32(0)
L2:
	;
	if v22 != int32(2249) {
		goto L4
	} else {
		goto L5
	}
L3:
	;
	m.G0 = v20 + int32(112)
	return v576
L4:
	;
	v51 = F_lookup_rowtype_tupdesc(m, v22, v23)
	mBase = m.M
	v52 = m.ExcPending
	if v52 != 0 {
		goto L1
	} else {
		goto L12
	}
L5:
	;
	if int32(0) <= v23 {
		goto L4
	} else {
		goto L6
	}
L6:
	;
	v34 = int32(0)
	v35 = F_errsave_start(m, v24)
	mBase = m.M
	v36 = m.ExcPending
	if v36 != 0 {
		goto L1
	} else {
		goto L7
	}
L7:
	;
	if v35 == int32(0) {
		v576 = v34
		goto L3
	} else {
		goto L8
	}
L8:
	;
	F_errcode(m, int32(1088))
	mBase = m.M
	v41 = m.ExcPending
	if v41 != 0 {
		goto L1
	} else {
		goto L9
	}
L9:
	;
	F_errmsg(m, int32(430224), int32(0))
	mBase = m.M
	v45 = m.ExcPending
	if v45 != 0 {
		goto L1
	} else {
		goto L10
	}
L10:
	;
	F_errsave_finish(m, v24, int32(477369), int32(105), int32(269572))
	mBase = m.M
	v50 = m.ExcPending
	if v50 != 0 {
		goto L1
	} else {
		goto L11
	}
L11:
	;
	v576 = v34
	goto L3
L12:
	;
	v53 = *(*int32)(unsafe.Add(mBase, uint32(v51)))
	v54 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v55 = *(*int32)(unsafe.Add(mBase, uint32(v54)+16))
	if v55 == int32(0) {
		goto L15
	} else {
		goto L16
	}
L13:
	;
	v119 = F_palloc(m, v53<<(uint(int32(2))%32))
	mBase = m.M
	v120 = m.ExcPending
	if v120 != 0 {
		goto L1
	} else {
		goto L33
	}
L14:
	;
	if v76 == v22 {
		goto L19
	} else {
		goto L20
	}
L15:
	;
	v61 = *(*int32)(unsafe.Add(mBase, uint32(v54)+20))
	v66 = F_MemoryContextAlloc(m, v61, v53*int32(44)+int32(12))
	mBase = m.M
	v67 = m.ExcPending
	if v67 != 0 {
		goto L1
	} else {
		goto L18
	}
L16:
	;
	v58 = *(*int32)(unsafe.Add(mBase, uint32(v55)+8))
	if v58 != v53 {
		goto L15
	} else {
		goto L17
	}
L17:
	;
	v60 = *(*int32)(unsafe.Add(mBase, uint32(v55)))
	v75 = v55
	v76 = v60
	goto L14
L18:
	;
	v68 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	*(*int32)(unsafe.Add(mBase, uint32(v68)+16)) = v66
	v70 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v71 = *(*int32)(unsafe.Add(mBase, uint32(v70)+16))
	*(*int64)(unsafe.Add(mBase, uint32(v71))) = int64(0)
	v75 = v71
	v76 = int32(0)
	goto L14
L19:
	;
	v78 = *(*int32)(unsafe.Add(mBase, uint32(v75)+4))
	if v78 == v23 {
		goto L13
	} else {
		goto L22
	}
L20:
	;
	goto L21
L21:
	;
	v81 = v53 * int32(44)
	v83 = v81 + int32(12)
	if v75&int32(3) != 0 {
		goto L24
	} else {
		goto L25
	}
L22:
	;
	goto L21
L23:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v75)+8)) = v53
	*(*int32)(unsafe.Add(mBase, uint32(v75)+4)) = v23
	*(*int32)(unsafe.Add(mBase, uint32(v75))) = v22
	goto L13
L24:
	;
	v109 = F__emscripten_memset_bulkmem(m, v75, base.I32_extend8_s(int32(0)), v83)
	mBase = m.M
	goto L32
L25:
	;
	if base.Ui32(int32(1024)) < base.Ui32(v83) {
		goto L24
	} else {
		goto L26
	}
L26:
	;
	if base.Ui32(v75+v83) <= base.Ui32(v75) {
		goto L23
	} else {
		goto L27
	}
L27:
	;
	v95 = v75 + v81 + int32(12)
	v97 = v75 + int32(4)
	if base.Ui32(v97) < base.Ui32(v95) {
		goto L28
	} else {
		goto L29
	}
L28:
	;
	v99 = v95
	goto L30
L29:
	;
	v99 = v97
	goto L30
L30:
	;
	v106 = F__emscripten_memset_bulkmem(m, v75, base.I32_extend8_s(int32(0)), (v75^int32(-1)+v99)&int32(-4)+int32(4))
	mBase = m.M
	goto L31
L31:
	;
	goto L23
L32:
	;
	goto L23
L33:
	;
	v121 = F_palloc(m, v53)
	mBase = m.M
	v122 = m.ExcPending
	if v122 != 0 {
		goto L1
	} else {
		goto L34
	}
L34:
	;
	v124 = v25
	goto L37
L35:
	;
	v567 = int32(0)
	v568 = *(*int32)(unsafe.Add(mBase, uint32(v51)+12))
	if v567 <= v568 {
		goto L147
	} else {
		goto L148
	}
L36:
	;
	v170 = v124 + int32(1)
	F_initStringInfo(m, v20+int32(96))
	mBase = m.M
	v174 = m.ExcPending
	if v174 != 0 {
		goto L1
	} else {
		goto L48
	}
L37:
	;
	v140 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v124))))
	if base.Ui32(v140-int32(9)) < base.Ui32(int32(5)) {
		goto L40
	} else {
		goto L41
	}
L38:
	;
	v149 = F_errsave_start(m, v24)
	mBase = m.M
	v150 = m.ExcPending
	if v150 != 0 {
		goto L1
	} else {
		goto L42
	}
L39:
	;
	goto L38
L40:
	;
	v124 = v124 + int32(1)
	goto L37
L41:
	;
	switch v140 - int32(32) {
	case 0:
		goto L40
	default:
		goto L39
	case 8:
		goto L36
	}
L42:
	;
	if v149 == int32(0) {
		goto L35
	} else {
		goto L43
	}
L43:
	;
	F_errcode(m, int32(33685634))
	mBase = m.M
	v155 = m.ExcPending
	if v155 != 0 {
		goto L1
	} else {
		goto L44
	}
L44:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v20))) = v25
	F_errmsg(m, int32(691636), v20)
	mBase = m.M
	v159 = m.ExcPending
	if v159 != 0 {
		goto L1
	} else {
		goto L45
	}
L45:
	;
	F_errdetail(m, int32(561085), int32(0))
	mBase = m.M
	v163 = m.ExcPending
	if v163 != 0 {
		goto L1
	} else {
		goto L46
	}
L46:
	;
	F_errsave_finish(m, v24, int32(477369), int32(159), int32(269572))
	mBase = m.M
	v168 = m.ExcPending
	if v168 != 0 {
		goto L1
	} else {
		goto L47
	}
L47:
	;
	goto L35
L48:
	;
	if int32(0) < v53 {
		goto L49
	} else {
		goto L50
	}
L49:
	;
	v181 = int32(0)
	v184 = v170
	v186 = v181
	v189 = v181
	goto L52
L50:
	;
	v434 = v170
	goto L51
L51:
	;
	v450 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v434))))
	if v450 == int32(41) {
		goto L112
	} else {
		goto L113
	}
L52:
	;
	v200 = *(*int32)(unsafe.Add(mBase, uint32(v51)))
	v206 = v51 + int32(20) + v200<<(uint(int32(4))%32) + v186*int32(100)
	v207 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v206)+91)))
	if v207 == int32(1) {
		goto L55
	} else {
		goto L56
	}
L53:
	;
	v434 = v414
	goto L51
L54:
	;
	v431 = v186 + int32(1)
	if v431 != v53 {
		v184 = v414
		v186 = v431
		v189 = v419
		goto L52
	} else {
		goto L110
	}
L55:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v119+v186<<(uint(int32(2))%32)))) = int32(0)
	v216 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(v186+v121))) = uint8(v216)
	v414 = v184
	v419 = v189
	goto L54
L56:
	;
	goto L57
L57:
	;
	v218 = *(*int32)(unsafe.Add(mBase, uint32(v206)+68))
	if v189&int32(1) != 0 {
		goto L61
	} else {
		goto L62
	}
L58:
	;
	*(*uint8)(unsafe.Add(mBase, uint32(v186+v121))) = uint8(v365)
	v384 = v75 + int32(12) + v186*int32(44)
	v385 = *(*int32)(unsafe.Add(mBase, uint32(v384)))
	if v218 != v385 {
		goto L103
	} else {
		goto L104
	}
L59:
	;
	v255 = v20 + int32(96)
	v256 = *(*int32)(unsafe.Add(mBase, uint32(v255)))
	v257 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v256))) = uint8(v257)
	*(*int32)(unsafe.Add(mBase, uint32(v255)+12)) = v257
	*(*int32)(unsafe.Add(mBase, uint32(v255)+4)) = v257
	goto L71
L60:
	;
	v232 = F_errsave_start(m, v24)
	mBase = m.M
	v233 = m.ExcPending
	if v233 != 0 {
		goto L1
	} else {
		goto L65
	}
L61:
	;
	v221 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v184))))
	if v221 != int32(44) {
		goto L60
	} else {
		goto L64
	}
L62:
	;
	v226 = v184
	goto L63
L63:
	;
	v227 = int32(0)
	v229 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v226))))
	switch v229 - int32(41) {
	case 0, 3:
		v364 = v226
		v365 = int32(1)
		v370 = v227
		goto L58
	default:
		goto L59
	}
L64:
	;
	v226 = v184 + int32(1)
	goto L63
L65:
	;
	if v232 == int32(0) {
		goto L35
	} else {
		goto L66
	}
L66:
	;
	F_errcode(m, int32(33685634))
	mBase = m.M
	v238 = m.ExcPending
	if v238 != 0 {
		goto L1
	} else {
		goto L67
	}
L67:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v20)+80)) = v25
	F_errmsg(m, int32(691636), v20+int32(80))
	mBase = m.M
	v244 = m.ExcPending
	if v244 != 0 {
		goto L1
	} else {
		goto L68
	}
L68:
	;
	F_errdetail(m, int32(558637), int32(0))
	mBase = m.M
	v248 = m.ExcPending
	if v248 != 0 {
		goto L1
	} else {
		goto L69
	}
L69:
	;
	F_errsave_finish(m, v24, int32(477369), int32(191), int32(269572))
	mBase = m.M
	v253 = m.ExcPending
	if v253 != 0 {
		goto L1
	} else {
		goto L70
	}
L70:
	;
	goto L35
L71:
	;
	v264 = v226
	v270 = v227
	goto L72
L72:
	;
	v280 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v264))))
	if v270&int32(1) != 0 {
		goto L75
	} else {
		goto L76
	}
L73:
	;
	v362 = *(*int32)(unsafe.Add(mBase, uint32(v20)+96))
	v364 = v264
	v365 = int32(0)
	v370 = v362
	goto L58
L74:
	;
	goto L73
L75:
	;
	v286 = v264 + int32(1)
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
	F_appendStringInfoChar(m, v20+int32(96), base.I32_extend8_s(v354))
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
	if v270&v340 == int32(0) {
		v264 = v286
		v270 = v340
		goto L72
	} else {
		goto L100
	}
L81:
	;
	if v280 != 0 {
		v353 = v286
		v354 = v280
		v355 = v270
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
	v291 = F_errsave_start(m, v24)
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
	*(*int32)(unsafe.Add(mBase, uint32(v20)+48)) = v25
	F_errmsg(m, int32(691636), v20+int32(48))
	mBase = m.M
	v303 = m.ExcPending
	if v303 != 0 {
		goto L1
	} else {
		goto L88
	}
L88:
	;
	F_errdetail(m, int32(544616), int32(0))
	mBase = m.M
	v307 = m.ExcPending
	if v307 != 0 {
		goto L1
	} else {
		goto L89
	}
L89:
	;
	F_errsave_finish(m, v24, int32(477369), int32(218), int32(269572))
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
	v316 = F_errsave_start(m, v24)
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
	v353 = v264 + int32(2)
	v354 = v313
	v355 = v270
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
	*(*int32)(unsafe.Add(mBase, uint32(v20)+64)) = v25
	F_errmsg(m, int32(691636), v20-int32(-64))
	mBase = m.M
	v328 = m.ExcPending
	if v328 != 0 {
		goto L1
	} else {
		goto L97
	}
L97:
	;
	F_errdetail(m, int32(544616), int32(0))
	mBase = m.M
	v332 = m.ExcPending
	if v332 != 0 {
		goto L1
	} else {
		goto L98
	}
L98:
	;
	F_errsave_finish(m, v24, int32(477369), int32(229), int32(269572))
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
		v264 = v286
		v270 = int32(0)
		goto L72
	} else {
		goto L101
	}
L101:
	;
	v353 = v264 + int32(2)
	v354 = int32(34)
	v355 = int32(1)
	goto L77
L102:
	;
	v264 = v353
	v270 = v355
	goto L72
L103:
	;
	F_getTypeInputInfo(m, v218, v384+int32(4), v384+int32(8))
	mBase = m.M
	v392 = m.ExcPending
	if v392 != 0 {
		goto L1
	} else {
		goto L106
	}
L104:
	;
	goto L105
L105:
	;
	v404 = *(*int32)(unsafe.Add(mBase, uint32(v384)+8))
	v405 = *(*int32)(unsafe.Add(mBase, uint32(v206)+76))
	v409 = F_InputFunctionCallSafe(m, v384+int32(16), v370, v404, v405, v24, v119+v186<<(uint(int32(2))%32))
	mBase = m.M
	v410 = m.ExcPending
	if v410 != 0 {
		goto L1
	} else {
		goto L108
	}
L106:
	;
	v393 = *(*int32)(unsafe.Add(mBase, uint32(v384)+4))
	v396 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v397 = *(*int32)(unsafe.Add(mBase, uint32(v396)+20))
	F_fmgr_info_cxt(m, v393, v384+int32(16), v397)
	mBase = m.M
	v399 = m.ExcPending
	if v399 != 0 {
		goto L1
	} else {
		goto L107
	}
L107:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v384))) = v218
	goto L105
L108:
	;
	if v409 == int32(0) {
		goto L35
	} else {
		goto L109
	}
L109:
	;
	v414 = v364
	v419 = int32(1)
	goto L54
L110:
	;
	goto L53
L111:
	;
	v525 = F_heap_form_tuple(m, v51, v119, v121)
	mBase = m.M
	v526 = m.ExcPending
	if v526 != 0 {
		goto L1
	} else {
		goto L132
	}
L112:
	;
	v454 = v434
	goto L115
L113:
	;
	goto L114
L114:
	;
	v503 = F_errsave_start(m, v24)
	mBase = m.M
	v504 = m.ExcPending
	if v504 != 0 {
		goto L1
	} else {
		goto L126
	}
L115:
	;
	v471 = v454 + int32(1)
	v472 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v471))))
	if base.Ui32(v472-int32(9)) < base.Ui32(int32(5)) {
		v454 = v471
		goto L115
	} else {
		goto L117
	}
L116:
	;
	if v472 == int32(0) {
		goto L111
	} else {
		goto L119
	}
L117:
	;
	if v472 == int32(32) {
		v454 = v471
		goto L115
	} else {
		goto L118
	}
L118:
	;
	goto L116
L119:
	;
	v481 = F_errsave_start(m, v24)
	mBase = m.M
	v482 = m.ExcPending
	if v482 != 0 {
		goto L1
	} else {
		goto L120
	}
L120:
	;
	if v481 == int32(0) {
		goto L35
	} else {
		goto L121
	}
L121:
	;
	F_errcode(m, int32(33685634))
	mBase = m.M
	v487 = m.ExcPending
	if v487 != 0 {
		goto L1
	} else {
		goto L122
	}
L122:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v20)+16)) = v25
	F_errmsg(m, int32(691636), v20+int32(16))
	mBase = m.M
	v493 = m.ExcPending
	if v493 != 0 {
		goto L1
	} else {
		goto L123
	}
L123:
	;
	F_errdetail(m, int32(561055), int32(0))
	mBase = m.M
	v497 = m.ExcPending
	if v497 != 0 {
		goto L1
	} else {
		goto L124
	}
L124:
	;
	F_errsave_finish(m, v24, int32(477369), int32(297), int32(269572))
	mBase = m.M
	v502 = m.ExcPending
	if v502 != 0 {
		goto L1
	} else {
		goto L125
	}
L125:
	;
	goto L35
L126:
	;
	if v503 == int32(0) {
		goto L35
	} else {
		goto L127
	}
L127:
	;
	F_errcode(m, int32(33685634))
	mBase = m.M
	v509 = m.ExcPending
	if v509 != 0 {
		goto L1
	} else {
		goto L128
	}
L128:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v20)+32)) = v25
	F_errmsg(m, int32(691636), v20+int32(32))
	mBase = m.M
	v515 = m.ExcPending
	if v515 != 0 {
		goto L1
	} else {
		goto L129
	}
L129:
	;
	F_errdetail(m, int32(558524), int32(0))
	mBase = m.M
	v519 = m.ExcPending
	if v519 != 0 {
		goto L1
	} else {
		goto L130
	}
L130:
	;
	F_errsave_finish(m, v24, int32(477369), int32(286), int32(269572))
	mBase = m.M
	v524 = m.ExcPending
	if v524 != 0 {
		goto L1
	} else {
		goto L131
	}
L131:
	;
	goto L35
L132:
	;
	v527 = *(*int32)(unsafe.Add(mBase, uint32(v525)))
	v528 = F_palloc(m, v527)
	mBase = m.M
	v529 = m.ExcPending
	if v529 != 0 {
		goto L1
	} else {
		goto L133
	}
L133:
	;
	v530 = *(*int32)(unsafe.Add(mBase, uint32(v525)+16))
	v531 = *(*int32)(unsafe.Add(mBase, uint32(v525)))
	if v531 != 0 {
		goto L135
	} else {
		goto L136
	}
L134:
	;
	F_pfree(m, v525)
	mBase = m.M
	v535 = m.ExcPending
	if v535 != 0 {
		goto L1
	} else {
		goto L138
	}
L135:
	;
	v532 = F__emscripten_memcpy_bulkmem(m, v528, v530, v531)
	mBase = m.M
	v533 = v532
	goto L137
L136:
	;
	v533 = v528
	goto L137
L137:
	;
	goto L134
L138:
	;
	v536 = *(*int32)(unsafe.Add(mBase, uint32(v20)+96))
	F_pfree(m, v536)
	mBase = m.M
	v538 = m.ExcPending
	if v538 != 0 {
		goto L1
	} else {
		goto L139
	}
L139:
	;
	F_pfree(m, v119)
	mBase = m.M
	v540 = m.ExcPending
	if v540 != 0 {
		goto L1
	} else {
		goto L140
	}
L140:
	;
	F_pfree(m, v121)
	mBase = m.M
	v542 = m.ExcPending
	if v542 != 0 {
		goto L1
	} else {
		goto L141
	}
L141:
	;
	v543 = *(*int32)(unsafe.Add(mBase, uint32(v51)+12))
	if int32(0) <= v543 {
		goto L142
	} else {
		goto L143
	}
L142:
	;
	F_DecrTupleDescRefCount(m, v51)
	mBase = m.M
	v547 = m.ExcPending
	if v547 != 0 {
		goto L1
	} else {
		goto L145
	}
L143:
	;
	goto L144
L144:
	;
	v548 = F_HeapTupleHeaderGetDatum(m, v533)
	mBase = m.M
	v549 = m.ExcPending
	if v549 != 0 {
		goto L1
	} else {
		goto L146
	}
L145:
	;
	goto L144
L146:
	;
	v576 = v548
	goto L3
L147:
	;
	F_DecrTupleDescRefCount(m, v51)
	mBase = m.M
	v572 = m.ExcPending
	if v572 != 0 {
		goto L1
	} else {
		goto L150
	}
L148:
	;
	goto L149
L149:
	;
	v573 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(l0)+16)) = uint8(v573)
	v576 = v567
	goto L3
L150:
	;
	goto L149
}
