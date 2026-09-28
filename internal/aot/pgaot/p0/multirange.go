package p0

import (
	base "github.com/shibukawa/pgmem/internal/aot/pgaot/base"
	"unsafe"
)

func F_multirange_after_multirange(m *base.Module, l0 int32) int64 {
	mBase := m.M
	_ = mBase
	var v7 int64
	_ = v7
	var v8 int32
	_ = v8
	var v10 int32
	_ = v10
	var v12 int32
	_ = v12
	var v13 int32
	_ = v13
	var v16 int32
	_ = v16
	var v17 int32
	_ = v17
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
	var v23 int32
	_ = v23
	var v26 int32
	_ = v26
	var v27 int32
	_ = v27
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
	var v40 int32
	_ = v40
	var v46 int32
	_ = v46
	var v48 int32
	_ = v48
	var v51 int32
	_ = v51
	var v55 int32
	_ = v55
	var v56 int32
	_ = v56
	var v57 int32
	_ = v57
	var v64 int64
	_ = v64
	var v72 int32
	_ = v72
	var v76 int32
	_ = v76
	var v81 int32
	_ = v81
	v7 = int64(0)
	v8 = m.G0
	v10 = v8 - int32(80)
	m.G0 = v10
	v12 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
	v13 = F_pg_detoast_datum(m, v12)
	mBase = m.M
	v16 = m.ExcPending
	if v16 != 0 {
		return int64(0)
	} else {
		v17 = *(*int32)(unsafe.Add(mBase, uint32(l0)+40))
		v18 = F_pg_detoast_datum(m, v17)
		mBase = m.M
		v19 = m.ExcPending
		if v19 != 0 {
			return int64(0)
		} else {
			v20 = *(*int32)(unsafe.Add(mBase, uint32(v13)+4))
			v21 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
			v22 = *(*int32)(unsafe.Add(mBase, uint32(v21)+16))
			if v22 != 0 {
				v23 = *(*int32)(unsafe.Add(mBase, uint32(v22)))
				if v23 == v20 {
					v33 = v22
					v34 = *(*int32)(unsafe.Add(mBase, uint32(v18)+8))
					if v34 == int32(0) {
						v64 = v7
						m.G0 = v10 + int32(80)
						return v64
					} else {
						v37 = *(*int32)(unsafe.Add(mBase, uint32(v13)+8))
						if v37 == int32(0) {
							v64 = v7
							m.G0 = v10 + int32(80)
							return v64
						} else {
							v40 = *(*int32)(unsafe.Add(mBase, uint32(v33)+296))
							v46 = v10 + int32(48)
							F_multirange_get_bounds(m, v40, v18, v34-int32(1), v10-int32(-64), v46)
							mBase = m.M
							v48 = m.ExcPending
							if v48 != 0 {
								return int64(0)
							} else {
								v51 = v10 + int32(32)
								F_multirange_get_bounds(m, v40, v13, int32(0), v51, v10+int32(16))
								mBase = m.M
								v55 = m.ExcPending
								if v55 != 0 {
									return int64(0)
								} else {
									v56 = F_range_cmp_bounds(m, v40, v46, v51)
									mBase = m.M
									v57 = m.ExcPending
									if v57 != 0 {
										return int64(0)
									} else {
										v64 = base.I64_extend_i32_u(int32(base.Ui32(v56) >> (uint(int32(31)) % 32)))
										m.G0 = v10 + int32(80)
										return v64
									}
								}
							}
						}
					}
				} else {
					v26 = F_lookup_type_cache(m, v20, int32(_a_F_multirange_after_multirange_0))
					mBase = m.M
					v27 = m.ExcPending
					if v27 != 0 {
						return int64(0)
					} else {
						v28 = *(*int32)(unsafe.Add(mBase, uint32(v26)+296))
						if v28 == int32(0) {
							F_errstart_cold(m, int32(21), int32(0))
							mBase = m.M
							v72 = m.ExcPending
							if v72 != 0 {
								return int64(0)
							} else {
								*(*int32)(unsafe.Add(mBase, uint32(v10))) = v20
								F_errmsg_internal(m, int32(_a_F_multirange_after_multirange_1), v10)
								mBase = m.M
								v76 = m.ExcPending
								if v76 != 0 {
									return int64(0)
								} else {
									F_errfinish(m, int32(_a_F_multirange_after_multirange_2), int32(561), int32(_a_F_multirange_after_multirange_3))
									mBase = m.M
									v81 = m.ExcPending
									if v81 != 0 {
										return int64(0)
									} else {
										base.Wasm_trap_unreachable()
										for {
										}
									}
								}
							}
						} else {
							v31 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
							*(*int32)(unsafe.Add(mBase, uint32(v31)+16)) = v26
							v33 = v26
							v34 = *(*int32)(unsafe.Add(mBase, uint32(v18)+8))
							if v34 == int32(0) {
								v64 = v7
								m.G0 = v10 + int32(80)
								return v64
							} else {
								v37 = *(*int32)(unsafe.Add(mBase, uint32(v13)+8))
								if v37 == int32(0) {
									v64 = v7
									m.G0 = v10 + int32(80)
									return v64
								} else {
									v40 = *(*int32)(unsafe.Add(mBase, uint32(v33)+296))
									v46 = v10 + int32(48)
									F_multirange_get_bounds(m, v40, v18, v34-int32(1), v10-int32(-64), v46)
									mBase = m.M
									v48 = m.ExcPending
									if v48 != 0 {
										return int64(0)
									} else {
										v51 = v10 + int32(32)
										F_multirange_get_bounds(m, v40, v13, int32(0), v51, v10+int32(16))
										mBase = m.M
										v55 = m.ExcPending
										if v55 != 0 {
											return int64(0)
										} else {
											v56 = F_range_cmp_bounds(m, v40, v46, v51)
											mBase = m.M
											v57 = m.ExcPending
											if v57 != 0 {
												return int64(0)
											} else {
												v64 = base.I64_extend_i32_u(int32(base.Ui32(v56) >> (uint(int32(31)) % 32)))
												m.G0 = v10 + int32(80)
												return v64
											}
										}
									}
								}
							}
						}
					}
				}
			} else {
				v26 = F_lookup_type_cache(m, v20, int32(_a_F_multirange_after_multirange_0))
				mBase = m.M
				v27 = m.ExcPending
				if v27 != 0 {
					return int64(0)
				} else {
					v28 = *(*int32)(unsafe.Add(mBase, uint32(v26)+296))
					if v28 == int32(0) {
						F_errstart_cold(m, int32(21), int32(0))
						mBase = m.M
						v72 = m.ExcPending
						if v72 != 0 {
							return int64(0)
						} else {
							*(*int32)(unsafe.Add(mBase, uint32(v10))) = v20
							F_errmsg_internal(m, int32(_a_F_multirange_after_multirange_1), v10)
							mBase = m.M
							v76 = m.ExcPending
							if v76 != 0 {
								return int64(0)
							} else {
								F_errfinish(m, int32(_a_F_multirange_after_multirange_2), int32(561), int32(_a_F_multirange_after_multirange_3))
								mBase = m.M
								v81 = m.ExcPending
								if v81 != 0 {
									return int64(0)
								} else {
									base.Wasm_trap_unreachable()
									for {
									}
								}
							}
						}
					} else {
						v31 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
						*(*int32)(unsafe.Add(mBase, uint32(v31)+16)) = v26
						v33 = v26
						v34 = *(*int32)(unsafe.Add(mBase, uint32(v18)+8))
						if v34 == int32(0) {
							v64 = v7
							m.G0 = v10 + int32(80)
							return v64
						} else {
							v37 = *(*int32)(unsafe.Add(mBase, uint32(v13)+8))
							if v37 == int32(0) {
								v64 = v7
								m.G0 = v10 + int32(80)
								return v64
							} else {
								v40 = *(*int32)(unsafe.Add(mBase, uint32(v33)+296))
								v46 = v10 + int32(48)
								F_multirange_get_bounds(m, v40, v18, v34-int32(1), v10-int32(-64), v46)
								mBase = m.M
								v48 = m.ExcPending
								if v48 != 0 {
									return int64(0)
								} else {
									v51 = v10 + int32(32)
									F_multirange_get_bounds(m, v40, v13, int32(0), v51, v10+int32(16))
									mBase = m.M
									v55 = m.ExcPending
									if v55 != 0 {
										return int64(0)
									} else {
										v56 = F_range_cmp_bounds(m, v40, v46, v51)
										mBase = m.M
										v57 = m.ExcPending
										if v57 != 0 {
											return int64(0)
										} else {
											v64 = base.I64_extend_i32_u(int32(base.Ui32(v56) >> (uint(int32(31)) % 32)))
											m.G0 = v10 + int32(80)
											return v64
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
}
func F_multirange_contains_elem_internal(m *base.Module, l0 int32, l1 int32, l2 int64) int32 {
	mBase := m.M
	_ = mBase
	var v10 int32
	_ = v10
	var v12 int32
	_ = v12
	var v14 int32
	_ = v14
	var v16 int32
	_ = v16
	var v21 int32
	_ = v21
	var v22 int32
	_ = v22
	var v28 int32
	_ = v28
	var v34 int32
	_ = v34
	var v35 int32
	_ = v35
	var v36 int32
	_ = v36
	var v37 int64
	_ = v37
	var v38 int64
	_ = v38
	var v39 int32
	_ = v39
	var v40 int32
	_ = v40
	var v43 int32
	_ = v43
	var v47 int32
	_ = v47
	var v48 int32
	_ = v48
	var v49 int64
	_ = v49
	var v50 int64
	_ = v50
	var v51 int32
	_ = v51
	var v52 int32
	_ = v52
	var v55 int32
	_ = v55
	var v60 int32
	_ = v60
	var v61 int32
	_ = v61
	var v85 int32
	_ = v85
	v10 = m.G0
	v12 = v10 - int32(32)
	m.G0 = v12
	v14 = *(*int32)(unsafe.Add(mBase, uint32(l1)+8))
	if v14 != 0 {
		goto L3
	} else {
		goto L4
	}
L1:
	;
	m.G0 = v12 + int32(32)
	return v85
L2:
	;
	v85 = int32(1)
	goto L1
L3:
	;
	v16 = l0 + int32(212)
	v21 = v14
	v22 = int32(0)
	goto L6
L4:
	;
	goto L5
L5:
	;
	v85 = int32(0)
	goto L1
L6:
	;
	v28 = int32(base.Ui32(v21+v22) >> (uint(int32(1)) % 32))
	F_multirange_get_bounds(m, l0, l1, v28, v12+int32(16), v12)
	mBase = m.M
	v34 = m.ExcPending
	if v34 != 0 {
		goto L8
	} else {
		goto L9
	}
L7:
	;
	goto L5
L8:
	;
	return int32(0)
L9:
	;
	v35 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v12)+24)))
	if v35 != 0 {
		goto L11
	} else {
		goto L12
	}
L10:
	;
	if base.Ui32(v61) < base.Ui32(v60) {
		v21 = v60
		v22 = v61
		goto L6
	} else {
		goto L26
	}
L11:
	;
	v47 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v12)+8)))
	if v47 != 0 {
		goto L2
	} else {
		goto L19
	}
L12:
	;
	v36 = *(*int32)(unsafe.Add(mBase, uint32(l0)+208))
	v37 = *(*int64)(unsafe.Add(mBase, uint32(v12)+16))
	v38 = F_FunctionCall2Coll(m, v16, v36, v37, l2)
	mBase = m.M
	v39 = m.ExcPending
	if v39 != 0 {
		goto L8
	} else {
		goto L13
	}
L13:
	;
	v40 = base.I32_wrap_i64(v38)
	if int32(0) < v40 {
		goto L14
	} else {
		goto L15
	}
L14:
	;
	v60 = v28
	v61 = v22
	goto L10
L15:
	;
	goto L16
L16:
	;
	if v40 != 0 {
		goto L11
	} else {
		goto L17
	}
L17:
	;
	v43 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v12)+25)))
	if v43&int32(1) != 0 {
		goto L11
	} else {
		goto L18
	}
L18:
	;
	v60 = v28
	v61 = v22
	goto L10
L19:
	;
	v48 = *(*int32)(unsafe.Add(mBase, uint32(l0)+208))
	v49 = *(*int64)(unsafe.Add(mBase, uint32(v12)))
	v50 = F_FunctionCall2Coll(m, v16, v48, v49, l2)
	mBase = m.M
	v51 = m.ExcPending
	if v51 != 0 {
		goto L8
	} else {
		goto L20
	}
L20:
	;
	v52 = base.I32_wrap_i64(v50)
	if int32(0) <= v52 {
		goto L21
	} else {
		goto L22
	}
L21:
	;
	if v52 != 0 {
		goto L2
	} else {
		goto L24
	}
L22:
	;
	goto L23
L23:
	;
	v60 = v21
	v61 = v28 + int32(1)
	goto L10
L24:
	;
	v55 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v12)+9)))
	if v55&int32(1) != 0 {
		goto L2
	} else {
		goto L25
	}
L25:
	;
	goto L23
L26:
	;
	goto L7
}
func F_multirange_gt(m *base.Module, l0 int32) int64 {
	var v2 int64
	_ = v2
	var v5 int32
	_ = v5
	v2 = F_multirange_cmp(m, l0)
	v5 = m.ExcPending
	if v5 != 0 {
		return int64(0)
	} else {
		return base.I64_extend_i32_u(base.B2i32(int64(0) < v2))
	}
}
func F_multirange_in(m *base.Module, l0 int32) int64 {
	mBase := m.M
	_ = mBase
	var v2 int32
	_ = v2
	var v21 int64
	_ = v21
	var v22 int32
	_ = v22
	var v24 int32
	_ = v24
	var v26 int32
	_ = v26
	var v27 int32
	_ = v27
	var v28 int32
	_ = v28
	var v29 int32
	_ = v29
	var v32 int32
	_ = v32
	var v35 int32
	_ = v35
	var v37 int32
	_ = v37
	var v38 int32
	_ = v38
	var v39 int32
	_ = v39
	var v40 int32
	_ = v40
	var v43 int32
	_ = v43
	var v62 int32
	_ = v62
	var v77 int32
	_ = v77
	var v78 int32
	_ = v78
	var v83 int32
	_ = v83
	var v87 int32
	_ = v87
	var v90 int32
	_ = v90
	var v91 int32
	_ = v91
	var v96 int32
	_ = v96
	var v99 int32
	_ = v99
	var v100 int32
	_ = v100
	var v106 int32
	_ = v106
	var v107 int32
	_ = v107
	var v108 int32
	_ = v108
	var v109 int32
	_ = v109
	var v112 int32
	_ = v112
	var v119 int32
	_ = v119
	var v120 int32
	_ = v120
	var v130 int32
	_ = v130
	var v131 int32
	_ = v131
	var v136 int32
	_ = v136
	var v142 int32
	_ = v142
	var v145 int32
	_ = v145
	var v146 int32
	_ = v146
	var v151 int32
	_ = v151
	var v152 int32
	_ = v152
	var v169 int32
	_ = v169
	var v170 int32
	_ = v170
	var v171 int32
	_ = v171
	var v174 int32
	_ = v174
	var v175 int32
	_ = v175
	var v185 int32
	_ = v185
	var v194 int32
	_ = v194
	var v197 int32
	_ = v197
	var v199 int32
	_ = v199
	var v216 int32
	_ = v216
	var v219 int32
	_ = v219
	var v224 int32
	_ = v224
	var v225 int32
	_ = v225
	var v230 int32
	_ = v230
	var v236 int32
	_ = v236
	var v239 int32
	_ = v239
	var v240 int32
	_ = v240
	var v245 int32
	_ = v245
	var v255 int32
	_ = v255
	var v256 int32
	_ = v256
	var v262 int32
	_ = v262
	var v263 int32
	_ = v263
	var v264 int32
	_ = v264
	var v265 int32
	_ = v265
	var v266 int32
	_ = v266
	var v269 int32
	_ = v269
	var v270 int32
	_ = v270
	var v273 int32
	_ = v273
	var v276 int32
	_ = v276
	var v277 int32
	_ = v277
	var v278 int32
	_ = v278
	var v279 int32
	_ = v279
	var v280 int32
	_ = v280
	var v281 int32
	_ = v281
	var v287 int32
	_ = v287
	var v303 int32
	_ = v303
	var v305 int32
	_ = v305
	var v306 int32
	_ = v306
	var v309 int32
	_ = v309
	var v316 int32
	_ = v316
	var v317 int32
	_ = v317
	var v322 int32
	_ = v322
	var v328 int32
	_ = v328
	var v331 int32
	_ = v331
	var v332 int32
	_ = v332
	var v337 int32
	_ = v337
	var v340 int32
	_ = v340
	var v341 int32
	_ = v341
	var v343 int32
	_ = v343
	var v344 int32
	_ = v344
	var v346 int32
	_ = v346
	var v347 int32
	_ = v347
	var v352 int32
	_ = v352
	var v373 int32
	_ = v373
	var v381 int32
	_ = v381
	var v382 int32
	_ = v382
	var v387 int32
	_ = v387
	var v393 int32
	_ = v393
	var v396 int32
	_ = v396
	var v397 int32
	_ = v397
	var v402 int32
	_ = v402
	var v403 int32
	_ = v403
	var v404 int32
	_ = v404
	var v426 int64
	_ = v426
	v2 = int32(0)
	v21 = int64(0)
	v22 = m.G0
	v24 = v22 - int32(80)
	m.G0 = v24
	v26 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v27 = *(*int32)(unsafe.Add(mBase, uint32(l0)+56))
	v28 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
	v29 = *(*int32)(unsafe.Add(mBase, uint32(l0)+40))
	v32 = F_palloc_mul(m, int32(4), int32(8))
	mBase = m.M
	v35 = m.ExcPending
	if v35 != 0 {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	return int64(0)
L2:
	;
	v37 = F_get_multirange_io_data(m, l0, v29, int32(0))
	mBase = m.M
	v38 = m.ExcPending
	if v38 != 0 {
		goto L1
	} else {
		goto L3
	}
L3:
	;
	v39 = *(*int32)(unsafe.Add(mBase, uint32(v37)))
	v40 = *(*int32)(unsafe.Add(mBase, uint32(v39)+296))
	v43 = v28
	goto L6
L4:
	;
	m.G0 = v24 + int32(80)
	return v426
L5:
	;
	v99 = v43
	v100 = v2
	v106 = v2
	v107 = v32
	v108 = v2
	v109 = v2
	v112 = int32(8)
	goto L20
L6:
	;
	v62 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v43))))
	if base.B2i32(base.Ui32(v62-int32(9)) < base.Ui32(int32(5)))|base.B2i32(v62 == int32(32)) != 0 {
		goto L8
	} else {
		goto L9
	}
L7:
	;
	v77 = F_errsave_start(m, v26)
	mBase = m.M
	v78 = m.ExcPending
	if v78 != 0 {
		goto L1
	} else {
		goto L14
	}
L8:
	;
	v43 = v43 + int32(1)
	goto L6
L9:
	;
	if v62 == int32(123) {
		goto L11
	} else {
		goto L12
	}
L10:
	;
	goto L7
L11:
	;
	goto L5
L12:
	;
	goto L13
L13:
	;
	goto L10
L14:
	;
	if v77 == int32(0) {
		v426 = v21
		goto L4
	} else {
		goto L15
	}
L15:
	;
	F_errcode(m, int32(33685634))
	mBase = m.M
	v83 = m.ExcPending
	if v83 != 0 {
		goto L1
	} else {
		goto L16
	}
L16:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v24))) = v28
	F_errmsg(m, int32(_a_F_multirange_in_0), v24)
	mBase = m.M
	v87 = m.ExcPending
	if v87 != 0 {
		goto L1
	} else {
		goto L17
	}
L17:
	;
	v90 = F_errdetail(m, int32(_a_F_multirange_in_1), int32(0))
	mBase = m.M
	v91 = m.ExcPending
	if v91 != 0 {
		goto L1
	} else {
		goto L18
	}
L18:
	;
	F_errsave_finish(m, v26, int32(_a_F_multirange_in_2), int32(152), int32(_a_F_multirange_in_3))
	mBase = m.M
	v96 = m.ExcPending
	if v96 != 0 {
		goto L1
	} else {
		goto L19
	}
L19:
	;
	v426 = v21
	goto L4
L20:
	;
	v119 = v99 + int32(1)
	v120 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v99)+1)))
	if base.B2i32(base.Ui32(v120-int32(9)) < base.Ui32(int32(5)))|base.B2i32(v120 == int32(32)) != 0 {
		v99 = v119
		goto L20
	} else {
		goto L22
	}
L21:
	;
	v352 = v119
	goto L100
L22:
	;
	if v120 == int32(0) {
		goto L23
	} else {
		goto L24
	}
L23:
	;
	v130 = F_errsave_start(m, v26)
	mBase = m.M
	v131 = m.ExcPending
	if v131 != 0 {
		goto L1
	} else {
		goto L26
	}
L24:
	;
	goto L25
L25:
	;
	v152 = int32(1)
	switch v100 - v152 {
	case 0:
		goto L38
	case 1:
		v99 = v119
		v100 = v152
		goto L20
	case 2:
		goto L37
	case 3:
		goto L35
	case 4:
		goto L36
	default:
		goto L39
	}
L26:
	;
	if v130 == int32(0) {
		v426 = v21
		goto L4
	} else {
		goto L27
	}
L27:
	;
	F_errcode(m, int32(33685634))
	mBase = m.M
	v136 = m.ExcPending
	if v136 != 0 {
		goto L1
	} else {
		goto L28
	}
L28:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v24)+64)) = v28
	F_errmsg(m, int32(_a_F_multirange_in_0), v24-int32(-64))
	mBase = m.M
	v142 = m.ExcPending
	if v142 != 0 {
		goto L1
	} else {
		goto L29
	}
L29:
	;
	v145 = F_errdetail(m, int32(_a_F_multirange_in_4), int32(0))
	mBase = m.M
	v146 = m.ExcPending
	if v146 != 0 {
		goto L1
	} else {
		goto L30
	}
L30:
	;
	F_errsave_finish(m, v26, int32(_a_F_multirange_in_2), int32(165), int32(_a_F_multirange_in_3))
	mBase = m.M
	v151 = m.ExcPending
	if v151 != 0 {
		goto L1
	} else {
		goto L31
	}
L31:
	;
	v426 = v21
	goto L4
L32:
	;
	goto L21
L33:
	;
	v99 = v340
	v100 = v341
	v106 = v343
	v107 = v344
	v109 = v346
	v112 = v347
	goto L20
L34:
	;
	v340 = v119
	v341 = int32(1)
	v343 = v106
	v344 = v107
	v346 = v109
	v347 = v112
	goto L33
L35:
	;
	v340 = v119
	v341 = int32(3)
	v343 = v106
	v344 = v107
	v346 = v109
	v347 = v112
	goto L33
L36:
	;
	if v120 == int32(44) {
		v99 = v119
		v100 = int32(0)
		goto L20
	} else {
		goto L92
	}
L37:
	;
	if v120 != int32(92) {
		goto L82
	} else {
		goto L83
	}
L38:
	;
	switch v120 - int32(34) {
	case 0:
		v99 = v119
		v100 = int32(3)
		goto L20
	case 1, 2, 3, 4, 5, 6:
		goto L34
	case 7:
		goto L67
	default:
		goto L68
	}
L39:
	;
	if base.B2i32(v120 == int32(40))|base.B2i32(v120 == int32(91)) != 0 {
		v99 = v119
		v100 = v152
		v108 = v119
		goto L20
	} else {
		goto L40
	}
L40:
	;
	if base.B2i32(v109 == int32(0))&base.B2i32(v120 == int32(125)) != 0 {
		goto L32
	} else {
		goto L41
	}
L41:
	;
	v169 = v119
	v170 = int32(_a_F_multirange_in_5)
	v171 = int32(5)
	goto L43
L42:
	;
	if v216 == int32(0) {
		goto L58
	} else {
		goto L59
	}
L43:
	;
	if v171 != 0 {
		goto L45
	} else {
		goto L46
	}
L44:
	;
	v216 = int32(0)
	goto L42
L45:
	;
	v174 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v169))))
	v175 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v170))))
	if v174 == v175 {
		v197 = v174
		goto L48
	} else {
		goto L49
	}
L46:
	;
	goto L47
L47:
	;
	goto L44
L48:
	;
	v199 = int32(1)
	if v197 != 0 {
		v169 = v169 + v199
		v170 = v170 + v199
		v171 = v171 - v199
		goto L43
	} else {
		goto L57
	}
L49:
	;
	if base.Ui32((v174-int32(65))&int32(255)) < base.Ui32(int32(26)) {
		goto L50
	} else {
		goto L51
	}
L50:
	;
	v185 = v174 | int32(32)
	goto L52
L51:
	;
	v185 = v174
	goto L52
L52:
	;
	if base.Ui32((v175-int32(65))&int32(255)) < base.Ui32(int32(26)) {
		goto L53
	} else {
		goto L54
	}
L53:
	;
	v194 = v175 | int32(32)
	goto L55
L54:
	;
	v194 = v175
	goto L55
L55:
	;
	if v185 == v194 {
		v197 = v185
		goto L48
	} else {
		goto L56
	}
L56:
	;
	v216 = v185 - v194
	goto L42
L57:
	;
	goto L47
L58:
	;
	v219 = int32(5)
	v340 = v99 + v219
	v341 = v219
	v343 = v106
	v344 = v107
	v346 = v109 + int32(1)
	v347 = v112
	goto L33
L59:
	;
	goto L60
L60:
	;
	v224 = F_errsave_start(m, v26)
	mBase = m.M
	v225 = m.ExcPending
	if v225 != 0 {
		goto L1
	} else {
		goto L61
	}
L61:
	;
	if v224 == int32(0) {
		v426 = v21
		goto L4
	} else {
		goto L62
	}
L62:
	;
	F_errcode(m, int32(33685634))
	mBase = m.M
	v230 = m.ExcPending
	if v230 != 0 {
		goto L1
	} else {
		goto L63
	}
L63:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v24)+32)) = v28
	F_errmsg(m, int32(_a_F_multirange_in_0), v24+int32(32))
	mBase = m.M
	v236 = m.ExcPending
	if v236 != 0 {
		goto L1
	} else {
		goto L64
	}
L64:
	;
	v239 = F_errdetail(m, int32(_a_F_multirange_in_6), int32(0))
	mBase = m.M
	v240 = m.ExcPending
	if v240 != 0 {
		goto L1
	} else {
		goto L65
	}
L65:
	;
	F_errsave_finish(m, v26, int32(_a_F_multirange_in_2), int32(194), int32(_a_F_multirange_in_3))
	mBase = m.M
	v245 = m.ExcPending
	if v245 != 0 {
		goto L1
	} else {
		goto L66
	}
L66:
	;
	v426 = v21
	goto L4
L67:
	;
	v255 = F_pnstrdup(m, v108, v119-v108+int32(1))
	mBase = m.M
	v256 = m.ExcPending
	if v256 != 0 {
		goto L1
	} else {
		goto L70
	}
L68:
	;
	switch v120 - int32(92) {
	case 0:
		goto L69
	case 1:
		goto L67
	default:
		goto L34
	}
L69:
	;
	v340 = v119
	v341 = int32(2)
	v343 = v106
	v344 = v107
	v346 = v109
	v347 = v112
	goto L33
L70:
	;
	if v106 == v112 {
		goto L71
	} else {
		goto L72
	}
L71:
	;
	v262 = F_repalloc(m, v107, v106<<(uint(int32(3))%32))
	mBase = m.M
	v263 = m.ExcPending
	if v263 != 0 {
		goto L1
	} else {
		goto L74
	}
L72:
	;
	v264 = v107
	v265 = v112
	goto L73
L73:
	;
	v266 = *(*int32)(unsafe.Add(mBase, uint32(v37)+32))
	v269 = F_InputFunctionCallSafe(m, v37+int32(4), v255, v266, v27, v26, v24+int32(72))
	mBase = m.M
	v270 = m.ExcPending
	if v270 != 0 {
		goto L1
	} else {
		goto L75
	}
L74:
	;
	v264 = v262
	v265 = v106 << (uint(int32(1)) % 32)
	goto L73
L75:
	;
	if v269 == int32(0) {
		goto L76
	} else {
		goto L77
	}
L76:
	;
	v273 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(l0)+16)) = uint8(v273)
	v426 = v21
	goto L4
L77:
	;
	goto L78
L78:
	;
	v276 = v109 + int32(1)
	v277 = int32(5)
	v278 = *(*int32)(unsafe.Add(mBase, uint32(v24)+72))
	v279 = F_pg_detoast_datum(m, v278)
	mBase = m.M
	v280 = m.ExcPending
	if v280 != 0 {
		goto L1
	} else {
		goto L79
	}
L79:
	;
	v281 = *(*int32)(unsafe.Add(mBase, uint32(v279)))
	v287 = int32(*(*int8)(unsafe.Add(mBase, uint32(v279+int32(base.Ui32(v281)>>(uint(int32(2))%32))-int32(1)))))
	goto L80
L80:
	;
	if v287&int32(1) != 0 {
		v99 = v119
		v100 = v277
		v107 = v264
		v109 = v276
		v112 = v265
		goto L20
	} else {
		goto L81
	}
L81:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v264+v106<<(uint(int32(2))%32)))) = v279
	v340 = v119
	v341 = v277
	v343 = v106 + int32(1)
	v344 = v264
	v346 = v276
	v347 = v265
	goto L33
L82:
	;
	if v120 != int32(34) {
		v99 = v119
		v100 = int32(3)
		goto L20
	} else {
		goto L85
	}
L83:
	;
	goto L84
L84:
	;
	v340 = v119
	v341 = int32(4)
	v343 = v106
	v344 = v107
	v346 = v109
	v347 = v112
	goto L33
L85:
	;
	v303 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v99)+2)))
	v305 = base.B2i32(v303 == int32(34))
	if v303 == int32(34) {
		goto L86
	} else {
		goto L87
	}
L86:
	;
	v306 = v99 + int32(2)
	goto L88
L87:
	;
	v306 = v119
	goto L88
L88:
	;
	if v303 == int32(34) {
		goto L89
	} else {
		goto L90
	}
L89:
	;
	v309 = int32(3)
	goto L91
L90:
	;
	v309 = int32(1)
	goto L91
L91:
	;
	v340 = v306
	v341 = v309
	v343 = v106
	v344 = v107
	v346 = v109
	v347 = v112
	goto L33
L92:
	;
	if v120 == int32(125) {
		goto L32
	} else {
		goto L93
	}
L93:
	;
	v316 = F_errsave_start(m, v26)
	mBase = m.M
	v317 = m.ExcPending
	if v317 != 0 {
		goto L1
	} else {
		goto L94
	}
L94:
	;
	if v316 == int32(0) {
		v426 = v21
		goto L4
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
	*(*int32)(unsafe.Add(mBase, uint32(v24)+48)) = v28
	F_errmsg(m, int32(_a_F_multirange_in_0), v24+int32(48))
	mBase = m.M
	v328 = m.ExcPending
	if v328 != 0 {
		goto L1
	} else {
		goto L97
	}
L97:
	;
	v331 = F_errdetail(m, int32(_a_F_multirange_in_7), int32(0))
	mBase = m.M
	v332 = m.ExcPending
	if v332 != 0 {
		goto L1
	} else {
		goto L98
	}
L98:
	;
	F_errsave_finish(m, v26, int32(_a_F_multirange_in_2), int32(268), int32(_a_F_multirange_in_3))
	mBase = m.M
	v337 = m.ExcPending
	if v337 != 0 {
		goto L1
	} else {
		goto L99
	}
L99:
	;
	v426 = v21
	goto L4
L100:
	;
	v373 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v352)+1)))
	if base.B2i32(base.Ui32(v373-int32(9)) < base.Ui32(int32(5)))|base.B2i32(v373 == int32(32)) != 0 {
		v352 = v352 + int32(1)
		goto L100
	} else {
		goto L102
	}
L101:
	;
	if v373 != 0 {
		goto L103
	} else {
		goto L104
	}
L102:
	;
	goto L101
L103:
	;
	v381 = F_errsave_start(m, v26)
	mBase = m.M
	v382 = m.ExcPending
	if v382 != 0 {
		goto L1
	} else {
		goto L106
	}
L104:
	;
	goto L105
L105:
	;
	v403 = F_make_multirange(m, v29, v40, v106, v107)
	mBase = m.M
	v404 = m.ExcPending
	if v404 != 0 {
		goto L1
	} else {
		goto L112
	}
L106:
	;
	if v381 == int32(0) {
		v426 = v21
		goto L4
	} else {
		goto L107
	}
L107:
	;
	F_errcode(m, int32(33685634))
	mBase = m.M
	v387 = m.ExcPending
	if v387 != 0 {
		goto L1
	} else {
		goto L108
	}
L108:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v24)+16)) = v28
	F_errmsg(m, int32(_a_F_multirange_in_0), v24+int32(16))
	mBase = m.M
	v393 = m.ExcPending
	if v393 != 0 {
		goto L1
	} else {
		goto L109
	}
L109:
	;
	v396 = F_errdetail(m, int32(_a_F_multirange_in_8), int32(0))
	mBase = m.M
	v397 = m.ExcPending
	if v397 != 0 {
		goto L1
	} else {
		goto L110
	}
L110:
	;
	F_errsave_finish(m, v26, int32(_a_F_multirange_in_2), int32(292), int32(_a_F_multirange_in_3))
	mBase = m.M
	v402 = m.ExcPending
	if v402 != 0 {
		goto L1
	} else {
		goto L111
	}
L111:
	;
	v426 = v21
	goto L4
L112:
	;
	v426 = base.I64_extend_i32_u(v403)
	goto L4
}
func F_multirange_le(m *base.Module, l0 int32) int64 {
	var v2 int64
	_ = v2
	var v5 int32
	_ = v5
	v2 = F_multirange_cmp(m, l0)
	v5 = m.ExcPending
	if v5 != 0 {
		return int64(0)
	} else {
		return base.I64_extend_i32_u(base.B2i32(v2 <= int64(0)))
	}
}
func F_multirange_ne(m *base.Module, l0 int32) int64 {
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
	var v20 int32
	_ = v20
	var v21 int32
	_ = v21
	var v22 int32
	_ = v22
	var v25 int32
	_ = v25
	var v26 int32
	_ = v26
	var v27 int32
	_ = v27
	var v30 int32
	_ = v30
	var v32 int32
	_ = v32
	var v33 int32
	_ = v33
	var v34 int32
	_ = v34
	var v35 int32
	_ = v35
	var v46 int32
	_ = v46
	var v50 int32
	_ = v50
	var v55 int32
	_ = v55
	v7 = m.G0
	v9 = v7 - int32(16)
	m.G0 = v9
	v11 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
	v12 = F_pg_detoast_datum(m, v11)
	mBase = m.M
	v15 = m.ExcPending
	if v15 != 0 {
		return int64(0)
	} else {
		v16 = *(*int32)(unsafe.Add(mBase, uint32(l0)+40))
		v17 = F_pg_detoast_datum(m, v16)
		mBase = m.M
		v18 = m.ExcPending
		if v18 != 0 {
			return int64(0)
		} else {
			v19 = *(*int32)(unsafe.Add(mBase, uint32(v12)+4))
			v20 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
			v21 = *(*int32)(unsafe.Add(mBase, uint32(v20)+16))
			if v21 != 0 {
				v22 = *(*int32)(unsafe.Add(mBase, uint32(v21)))
				if v22 == v19 {
					v32 = v21
					v33 = *(*int32)(unsafe.Add(mBase, uint32(v32)+296))
					v34 = F_multirange_eq_internal(m, v33, v12, v17)
					mBase = m.M
					v35 = m.ExcPending
					if v35 != 0 {
						return int64(0)
					} else {
						m.G0 = v9 + int32(16)
						return base.I64_extend_i32_u(v34 ^ int32(1))
					}
				} else {
					v25 = F_lookup_type_cache(m, v19, int32(_a_F_multirange_ne_0))
					mBase = m.M
					v26 = m.ExcPending
					if v26 != 0 {
						return int64(0)
					} else {
						v27 = *(*int32)(unsafe.Add(mBase, uint32(v25)+296))
						if v27 == int32(0) {
							F_errstart_cold(m, int32(21), int32(0))
							mBase = m.M
							v46 = m.ExcPending
							if v46 != 0 {
								return int64(0)
							} else {
								*(*int32)(unsafe.Add(mBase, uint32(v9))) = v19
								F_errmsg_internal(m, int32(_a_F_multirange_ne_1), v9)
								mBase = m.M
								v50 = m.ExcPending
								if v50 != 0 {
									return int64(0)
								} else {
									F_errfinish(m, int32(_a_F_multirange_ne_2), int32(561), int32(_a_F_multirange_ne_3))
									mBase = m.M
									v55 = m.ExcPending
									if v55 != 0 {
										return int64(0)
									} else {
										base.Wasm_trap_unreachable()
										for {
										}
									}
								}
							}
						} else {
							v30 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
							*(*int32)(unsafe.Add(mBase, uint32(v30)+16)) = v25
							v32 = v25
							v33 = *(*int32)(unsafe.Add(mBase, uint32(v32)+296))
							v34 = F_multirange_eq_internal(m, v33, v12, v17)
							mBase = m.M
							v35 = m.ExcPending
							if v35 != 0 {
								return int64(0)
							} else {
								m.G0 = v9 + int32(16)
								return base.I64_extend_i32_u(v34 ^ int32(1))
							}
						}
					}
				}
			} else {
				v25 = F_lookup_type_cache(m, v19, int32(_a_F_multirange_ne_0))
				mBase = m.M
				v26 = m.ExcPending
				if v26 != 0 {
					return int64(0)
				} else {
					v27 = *(*int32)(unsafe.Add(mBase, uint32(v25)+296))
					if v27 == int32(0) {
						F_errstart_cold(m, int32(21), int32(0))
						mBase = m.M
						v46 = m.ExcPending
						if v46 != 0 {
							return int64(0)
						} else {
							*(*int32)(unsafe.Add(mBase, uint32(v9))) = v19
							F_errmsg_internal(m, int32(_a_F_multirange_ne_1), v9)
							mBase = m.M
							v50 = m.ExcPending
							if v50 != 0 {
								return int64(0)
							} else {
								F_errfinish(m, int32(_a_F_multirange_ne_2), int32(561), int32(_a_F_multirange_ne_3))
								mBase = m.M
								v55 = m.ExcPending
								if v55 != 0 {
									return int64(0)
								} else {
									base.Wasm_trap_unreachable()
									for {
									}
								}
							}
						}
					} else {
						v30 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
						*(*int32)(unsafe.Add(mBase, uint32(v30)+16)) = v25
						v32 = v25
						v33 = *(*int32)(unsafe.Add(mBase, uint32(v32)+296))
						v34 = F_multirange_eq_internal(m, v33, v12, v17)
						mBase = m.M
						v35 = m.ExcPending
						if v35 != 0 {
							return int64(0)
						} else {
							m.G0 = v9 + int32(16)
							return base.I64_extend_i32_u(v34 ^ int32(1))
						}
					}
				}
			}
		}
	}
}
func F_multirange_overleft_range(m *base.Module, l0 int32) int64 {
	mBase := m.M
	_ = mBase
	var v7 int64
	_ = v7
	var v8 int32
	_ = v8
	var v10 int32
	_ = v10
	var v12 int32
	_ = v12
	var v13 int32
	_ = v13
	var v16 int32
	_ = v16
	var v17 int32
	_ = v17
	var v18 int32
	_ = v18
	var v19 int32
	_ = v19
	var v20 int32
	_ = v20
	var v23 int32
	_ = v23
	var v29 int32
	_ = v29
	var v32 int32
	_ = v32
	var v33 int32
	_ = v33
	var v34 int32
	_ = v34
	var v35 int32
	_ = v35
	var v38 int32
	_ = v38
	var v39 int32
	_ = v39
	var v40 int32
	_ = v40
	var v43 int32
	_ = v43
	var v45 int32
	_ = v45
	var v46 int32
	_ = v46
	var v47 int32
	_ = v47
	var v53 int32
	_ = v53
	var v55 int32
	_ = v55
	var v56 int32
	_ = v56
	var v60 int32
	_ = v60
	var v64 int32
	_ = v64
	var v65 int32
	_ = v65
	var v66 int32
	_ = v66
	var v67 int32
	_ = v67
	var v75 int64
	_ = v75
	var v83 int32
	_ = v83
	var v87 int32
	_ = v87
	var v92 int32
	_ = v92
	v7 = int64(0)
	v8 = m.G0
	v10 = v8 - int32(80)
	m.G0 = v10
	v12 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
	v13 = F_pg_detoast_datum(m, v12)
	mBase = m.M
	v16 = m.ExcPending
	if v16 != 0 {
		return int64(0)
	} else {
		v17 = *(*int32)(unsafe.Add(mBase, uint32(l0)+40))
		v18 = F_pg_detoast_datum(m, v17)
		mBase = m.M
		v19 = m.ExcPending
		if v19 != 0 {
			return int64(0)
		} else {
			v20 = *(*int32)(unsafe.Add(mBase, uint32(v13)+8))
			if v20 == int32(0) {
				v75 = v7
				m.G0 = v10 + int32(80)
				return v75
			} else {
				v23 = *(*int32)(unsafe.Add(mBase, uint32(v18)))
				v29 = int32(*(*int8)(unsafe.Add(mBase, uint32(v18+int32(base.Ui32(v23)>>(uint(int32(2))%32))-int32(1)))))
				if v29&int32(1) != 0 {
					v75 = v7
					m.G0 = v10 + int32(80)
					return v75
				} else {
					v32 = *(*int32)(unsafe.Add(mBase, uint32(v13)+4))
					v33 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
					v34 = *(*int32)(unsafe.Add(mBase, uint32(v33)+16))
					if v34 != 0 {
						v35 = *(*int32)(unsafe.Add(mBase, uint32(v34)))
						if v35 == v32 {
							v45 = v34
							v46 = *(*int32)(unsafe.Add(mBase, uint32(v45)+296))
							v47 = *(*int32)(unsafe.Add(mBase, uint32(v13)+8))
							v53 = v10 + int32(48)
							F_multirange_get_bounds(m, v46, v13, v47-int32(1), v10-int32(-64), v53)
							mBase = m.M
							v55 = m.ExcPending
							if v55 != 0 {
								return int64(0)
							} else {
								v56 = *(*int32)(unsafe.Add(mBase, uint32(v45)+296))
								v60 = v10 + int32(16)
								F_range_deserialize(m, v56, v18, v10+int32(32), v60, v10+int32(15))
								mBase = m.M
								v64 = m.ExcPending
								if v64 != 0 {
									return int64(0)
								} else {
									v65 = *(*int32)(unsafe.Add(mBase, uint32(v45)+296))
									v66 = F_range_cmp_bounds(m, v65, v53, v60)
									mBase = m.M
									v67 = m.ExcPending
									if v67 != 0 {
										return int64(0)
									} else {
										v75 = base.I64_extend_i32_u(base.B2i32(v66 <= int32(0)))
										m.G0 = v10 + int32(80)
										return v75
									}
								}
							}
						} else {
							v38 = F_lookup_type_cache(m, v32, int32(_a_F_multirange_overleft_range_0))
							mBase = m.M
							v39 = m.ExcPending
							if v39 != 0 {
								return int64(0)
							} else {
								v40 = *(*int32)(unsafe.Add(mBase, uint32(v38)+296))
								if v40 == int32(0) {
									F_errstart_cold(m, int32(21), int32(0))
									mBase = m.M
									v83 = m.ExcPending
									if v83 != 0 {
										return int64(0)
									} else {
										*(*int32)(unsafe.Add(mBase, uint32(v10))) = v32
										F_errmsg_internal(m, int32(_a_F_multirange_overleft_range_1), v10)
										mBase = m.M
										v87 = m.ExcPending
										if v87 != 0 {
											return int64(0)
										} else {
											F_errfinish(m, int32(_a_F_multirange_overleft_range_2), int32(561), int32(_a_F_multirange_overleft_range_3))
											mBase = m.M
											v92 = m.ExcPending
											if v92 != 0 {
												return int64(0)
											} else {
												base.Wasm_trap_unreachable()
												for {
												}
											}
										}
									}
								} else {
									v43 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
									*(*int32)(unsafe.Add(mBase, uint32(v43)+16)) = v38
									v45 = v38
									v46 = *(*int32)(unsafe.Add(mBase, uint32(v45)+296))
									v47 = *(*int32)(unsafe.Add(mBase, uint32(v13)+8))
									v53 = v10 + int32(48)
									F_multirange_get_bounds(m, v46, v13, v47-int32(1), v10-int32(-64), v53)
									mBase = m.M
									v55 = m.ExcPending
									if v55 != 0 {
										return int64(0)
									} else {
										v56 = *(*int32)(unsafe.Add(mBase, uint32(v45)+296))
										v60 = v10 + int32(16)
										F_range_deserialize(m, v56, v18, v10+int32(32), v60, v10+int32(15))
										mBase = m.M
										v64 = m.ExcPending
										if v64 != 0 {
											return int64(0)
										} else {
											v65 = *(*int32)(unsafe.Add(mBase, uint32(v45)+296))
											v66 = F_range_cmp_bounds(m, v65, v53, v60)
											mBase = m.M
											v67 = m.ExcPending
											if v67 != 0 {
												return int64(0)
											} else {
												v75 = base.I64_extend_i32_u(base.B2i32(v66 <= int32(0)))
												m.G0 = v10 + int32(80)
												return v75
											}
										}
									}
								}
							}
						}
					} else {
						v38 = F_lookup_type_cache(m, v32, int32(_a_F_multirange_overleft_range_0))
						mBase = m.M
						v39 = m.ExcPending
						if v39 != 0 {
							return int64(0)
						} else {
							v40 = *(*int32)(unsafe.Add(mBase, uint32(v38)+296))
							if v40 == int32(0) {
								F_errstart_cold(m, int32(21), int32(0))
								mBase = m.M
								v83 = m.ExcPending
								if v83 != 0 {
									return int64(0)
								} else {
									*(*int32)(unsafe.Add(mBase, uint32(v10))) = v32
									F_errmsg_internal(m, int32(_a_F_multirange_overleft_range_1), v10)
									mBase = m.M
									v87 = m.ExcPending
									if v87 != 0 {
										return int64(0)
									} else {
										F_errfinish(m, int32(_a_F_multirange_overleft_range_2), int32(561), int32(_a_F_multirange_overleft_range_3))
										mBase = m.M
										v92 = m.ExcPending
										if v92 != 0 {
											return int64(0)
										} else {
											base.Wasm_trap_unreachable()
											for {
											}
										}
									}
								}
							} else {
								v43 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
								*(*int32)(unsafe.Add(mBase, uint32(v43)+16)) = v38
								v45 = v38
								v46 = *(*int32)(unsafe.Add(mBase, uint32(v45)+296))
								v47 = *(*int32)(unsafe.Add(mBase, uint32(v13)+8))
								v53 = v10 + int32(48)
								F_multirange_get_bounds(m, v46, v13, v47-int32(1), v10-int32(-64), v53)
								mBase = m.M
								v55 = m.ExcPending
								if v55 != 0 {
									return int64(0)
								} else {
									v56 = *(*int32)(unsafe.Add(mBase, uint32(v45)+296))
									v60 = v10 + int32(16)
									F_range_deserialize(m, v56, v18, v10+int32(32), v60, v10+int32(15))
									mBase = m.M
									v64 = m.ExcPending
									if v64 != 0 {
										return int64(0)
									} else {
										v65 = *(*int32)(unsafe.Add(mBase, uint32(v45)+296))
										v66 = F_range_cmp_bounds(m, v65, v53, v60)
										mBase = m.M
										v67 = m.ExcPending
										if v67 != 0 {
											return int64(0)
										} else {
											v75 = base.I64_extend_i32_u(base.B2i32(v66 <= int32(0)))
											m.G0 = v10 + int32(80)
											return v75
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
}
func F_multirange_overright_multirange(m *base.Module, l0 int32) int64 {
	mBase := m.M
	_ = mBase
	var v7 int64
	_ = v7
	var v8 int32
	_ = v8
	var v10 int32
	_ = v10
	var v12 int32
	_ = v12
	var v13 int32
	_ = v13
	var v16 int32
	_ = v16
	var v17 int32
	_ = v17
	var v18 int32
	_ = v18
	var v19 int32
	_ = v19
	var v20 int32
	_ = v20
	var v23 int32
	_ = v23
	var v26 int32
	_ = v26
	var v27 int32
	_ = v27
	var v28 int32
	_ = v28
	var v29 int32
	_ = v29
	var v32 int32
	_ = v32
	var v33 int32
	_ = v33
	var v34 int32
	_ = v34
	var v37 int32
	_ = v37
	var v39 int32
	_ = v39
	var v40 int32
	_ = v40
	var v43 int32
	_ = v43
	var v47 int32
	_ = v47
	var v48 int32
	_ = v48
	var v51 int32
	_ = v51
	var v55 int32
	_ = v55
	var v56 int32
	_ = v56
	var v57 int32
	_ = v57
	var v58 int32
	_ = v58
	var v66 int64
	_ = v66
	var v74 int32
	_ = v74
	var v78 int32
	_ = v78
	var v83 int32
	_ = v83
	v7 = int64(0)
	v8 = m.G0
	v10 = v8 - int32(80)
	m.G0 = v10
	v12 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
	v13 = F_pg_detoast_datum(m, v12)
	mBase = m.M
	v16 = m.ExcPending
	if v16 != 0 {
		return int64(0)
	} else {
		v17 = *(*int32)(unsafe.Add(mBase, uint32(l0)+40))
		v18 = F_pg_detoast_datum(m, v17)
		mBase = m.M
		v19 = m.ExcPending
		if v19 != 0 {
			return int64(0)
		} else {
			v20 = *(*int32)(unsafe.Add(mBase, uint32(v13)+8))
			if v20 == int32(0) {
				v66 = v7
				m.G0 = v10 + int32(80)
				return v66
			} else {
				v23 = *(*int32)(unsafe.Add(mBase, uint32(v18)+8))
				if v23 == int32(0) {
					v66 = v7
					m.G0 = v10 + int32(80)
					return v66
				} else {
					v26 = *(*int32)(unsafe.Add(mBase, uint32(v13)+4))
					v27 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
					v28 = *(*int32)(unsafe.Add(mBase, uint32(v27)+16))
					if v28 != 0 {
						v29 = *(*int32)(unsafe.Add(mBase, uint32(v28)))
						if v29 == v26 {
							v39 = v28
							v40 = *(*int32)(unsafe.Add(mBase, uint32(v39)+296))
							v43 = v10 - int32(-64)
							F_multirange_get_bounds(m, v40, v13, int32(0), v43, v10+int32(48))
							mBase = m.M
							v47 = m.ExcPending
							if v47 != 0 {
								return int64(0)
							} else {
								v48 = *(*int32)(unsafe.Add(mBase, uint32(v39)+296))
								v51 = v10 + int32(32)
								F_multirange_get_bounds(m, v48, v18, int32(0), v51, v10+int32(16))
								mBase = m.M
								v55 = m.ExcPending
								if v55 != 0 {
									return int64(0)
								} else {
									v56 = *(*int32)(unsafe.Add(mBase, uint32(v39)+296))
									v57 = F_range_cmp_bounds(m, v56, v43, v51)
									mBase = m.M
									v58 = m.ExcPending
									if v58 != 0 {
										return int64(0)
									} else {
										v66 = base.I64_extend_i32_u(base.B2i32(int32(0) <= v57))
										m.G0 = v10 + int32(80)
										return v66
									}
								}
							}
						} else {
							v32 = F_lookup_type_cache(m, v26, int32(_a_F_multirange_overright_multirange_0))
							mBase = m.M
							v33 = m.ExcPending
							if v33 != 0 {
								return int64(0)
							} else {
								v34 = *(*int32)(unsafe.Add(mBase, uint32(v32)+296))
								if v34 == int32(0) {
									F_errstart_cold(m, int32(21), int32(0))
									mBase = m.M
									v74 = m.ExcPending
									if v74 != 0 {
										return int64(0)
									} else {
										*(*int32)(unsafe.Add(mBase, uint32(v10))) = v26
										F_errmsg_internal(m, int32(_a_F_multirange_overright_multirange_1), v10)
										mBase = m.M
										v78 = m.ExcPending
										if v78 != 0 {
											return int64(0)
										} else {
											F_errfinish(m, int32(_a_F_multirange_overright_multirange_2), int32(561), int32(_a_F_multirange_overright_multirange_3))
											mBase = m.M
											v83 = m.ExcPending
											if v83 != 0 {
												return int64(0)
											} else {
												base.Wasm_trap_unreachable()
												for {
												}
											}
										}
									}
								} else {
									v37 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
									*(*int32)(unsafe.Add(mBase, uint32(v37)+16)) = v32
									v39 = v32
									v40 = *(*int32)(unsafe.Add(mBase, uint32(v39)+296))
									v43 = v10 - int32(-64)
									F_multirange_get_bounds(m, v40, v13, int32(0), v43, v10+int32(48))
									mBase = m.M
									v47 = m.ExcPending
									if v47 != 0 {
										return int64(0)
									} else {
										v48 = *(*int32)(unsafe.Add(mBase, uint32(v39)+296))
										v51 = v10 + int32(32)
										F_multirange_get_bounds(m, v48, v18, int32(0), v51, v10+int32(16))
										mBase = m.M
										v55 = m.ExcPending
										if v55 != 0 {
											return int64(0)
										} else {
											v56 = *(*int32)(unsafe.Add(mBase, uint32(v39)+296))
											v57 = F_range_cmp_bounds(m, v56, v43, v51)
											mBase = m.M
											v58 = m.ExcPending
											if v58 != 0 {
												return int64(0)
											} else {
												v66 = base.I64_extend_i32_u(base.B2i32(int32(0) <= v57))
												m.G0 = v10 + int32(80)
												return v66
											}
										}
									}
								}
							}
						}
					} else {
						v32 = F_lookup_type_cache(m, v26, int32(_a_F_multirange_overright_multirange_0))
						mBase = m.M
						v33 = m.ExcPending
						if v33 != 0 {
							return int64(0)
						} else {
							v34 = *(*int32)(unsafe.Add(mBase, uint32(v32)+296))
							if v34 == int32(0) {
								F_errstart_cold(m, int32(21), int32(0))
								mBase = m.M
								v74 = m.ExcPending
								if v74 != 0 {
									return int64(0)
								} else {
									*(*int32)(unsafe.Add(mBase, uint32(v10))) = v26
									F_errmsg_internal(m, int32(_a_F_multirange_overright_multirange_1), v10)
									mBase = m.M
									v78 = m.ExcPending
									if v78 != 0 {
										return int64(0)
									} else {
										F_errfinish(m, int32(_a_F_multirange_overright_multirange_2), int32(561), int32(_a_F_multirange_overright_multirange_3))
										mBase = m.M
										v83 = m.ExcPending
										if v83 != 0 {
											return int64(0)
										} else {
											base.Wasm_trap_unreachable()
											for {
											}
										}
									}
								}
							} else {
								v37 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
								*(*int32)(unsafe.Add(mBase, uint32(v37)+16)) = v32
								v39 = v32
								v40 = *(*int32)(unsafe.Add(mBase, uint32(v39)+296))
								v43 = v10 - int32(-64)
								F_multirange_get_bounds(m, v40, v13, int32(0), v43, v10+int32(48))
								mBase = m.M
								v47 = m.ExcPending
								if v47 != 0 {
									return int64(0)
								} else {
									v48 = *(*int32)(unsafe.Add(mBase, uint32(v39)+296))
									v51 = v10 + int32(32)
									F_multirange_get_bounds(m, v48, v18, int32(0), v51, v10+int32(16))
									mBase = m.M
									v55 = m.ExcPending
									if v55 != 0 {
										return int64(0)
									} else {
										v56 = *(*int32)(unsafe.Add(mBase, uint32(v39)+296))
										v57 = F_range_cmp_bounds(m, v56, v43, v51)
										mBase = m.M
										v58 = m.ExcPending
										if v58 != 0 {
											return int64(0)
										} else {
											v66 = base.I64_extend_i32_u(base.B2i32(int32(0) <= v57))
											m.G0 = v10 + int32(80)
											return v66
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
}
func F_multirange_typanalyze(m *base.Module, l0 int32) int64 {
	mBase := m.M
	_ = mBase
	var v4 int32
	_ = v4
	var v5 int32
	_ = v5
	var v6 int32
	_ = v6
	var v9 int32
	_ = v9
	var v10 int32
	_ = v10
	var v11 int32
	_ = v11
	var v12 int32
	_ = v12
	var v16 int32
	_ = v16
	var v18 int32
	_ = v18
	v4 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
	v5 = *(*int32)(unsafe.Add(mBase, uint32(v4)+4))
	v6 = F_getBaseType(m, v5)
	mBase = m.M
	v9 = m.ExcPending
	if v9 != 0 {
		return int64(0)
	} else {
		v10 = F_multirange_get_typcache(m, l0, v6)
		mBase = m.M
		v11 = m.ExcPending
		if v11 != 0 {
			return int64(0)
		} else {
			v12 = *(*int32)(unsafe.Add(mBase, uint32(v4)))
			if v12 < int32(0) {
				v16 = *(*int32)(unsafe.Add(mBase, _c_F_multirange_typanalyze[0]))
				*(*int32)(unsafe.Add(mBase, uint32(v4))) = v16
				v18 = v16
			} else {
				v18 = v12
			}
			*(*int32)(unsafe.Add(mBase, uint32(v4)+32)) = v10
			*(*int32)(unsafe.Add(mBase, uint32(v4)+24)) = int32(1686)
			*(*int32)(unsafe.Add(mBase, uint32(v4)+28)) = v18 * int32(300)
			return int64(1)
		}
	}
}
func F_multirange_union(m *base.Module, l0 int32) int64 {
	mBase := m.M
	_ = mBase
	var v2 int32
	_ = v2
	var v10 int32
	_ = v10
	var v12 int32
	_ = v12
	var v14 int32
	_ = v14
	var v15 int32
	_ = v15
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
	var v34 int32
	_ = v34
	var v35 int32
	_ = v35
	var v36 int32
	_ = v36
	var v39 int32
	_ = v39
	var v41 int32
	_ = v41
	var v42 int32
	_ = v42
	var v43 int32
	_ = v43
	var v44 int32
	_ = v44
	var v49 int32
	_ = v49
	var v50 int32
	_ = v50
	var v51 int32
	_ = v51
	var v63 int32
	_ = v63
	var v64 int32
	_ = v64
	var v67 int32
	_ = v67
	var v69 int32
	_ = v69
	var v72 int32
	_ = v72
	var v77 int32
	_ = v77
	var v79 int32
	_ = v79
	var v84 int32
	_ = v84
	var v85 int32
	_ = v85
	var v86 int32
	_ = v86
	var v98 int32
	_ = v98
	var v99 int32
	_ = v99
	var v102 int32
	_ = v102
	var v112 int32
	_ = v112
	var v113 int32
	_ = v113
	var v116 int32
	_ = v116
	var v117 int32
	_ = v117
	var v119 int32
	_ = v119
	var v122 int32
	_ = v122
	var v125 int32
	_ = v125
	var v126 int32
	_ = v126
	var v127 int32
	_ = v127
	var v128 int32
	_ = v128
	var v132 int32
	_ = v132
	var v146 int32
	_ = v146
	var v150 int32
	_ = v150
	var v155 int32
	_ = v155
	v2 = int32(0)
	v10 = m.G0
	v12 = v10 - int32(16)
	m.G0 = v12
	v14 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
	v15 = F_pg_detoast_datum(m, v14)
	mBase = m.M
	v18 = m.ExcPending
	if v18 != 0 {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	return int64(0)
L2:
	;
	v19 = *(*int32)(unsafe.Add(mBase, uint32(l0)+40))
	v20 = F_pg_detoast_datum(m, v19)
	mBase = m.M
	v21 = m.ExcPending
	if v21 != 0 {
		goto L1
	} else {
		goto L3
	}
L3:
	;
	v22 = *(*int32)(unsafe.Add(mBase, uint32(v15)+8))
	if v22 == int32(0) {
		v132 = v20
		goto L5
	} else {
		goto L6
	}
L4:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v146 = m.ExcPending
	if v146 != 0 {
		goto L1
	} else {
		goto L41
	}
L5:
	;
	m.G0 = v12 + int32(16)
	return base.I64_extend_i32_u(v132)
L6:
	;
	v25 = *(*int32)(unsafe.Add(mBase, uint32(v20)+8))
	if v25 == int32(0) {
		goto L7
	} else {
		goto L8
	}
L7:
	;
	v132 = v15
	goto L5
L8:
	;
	goto L9
L9:
	;
	v28 = *(*int32)(unsafe.Add(mBase, uint32(v15)+4))
	v29 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v30 = *(*int32)(unsafe.Add(mBase, uint32(v29)+16))
	if v30 != 0 {
		goto L11
	} else {
		goto L12
	}
L10:
	;
	v44 = *(*int32)(unsafe.Add(mBase, uint32(v42)+296))
	if int32(0) < v43 {
		goto L17
	} else {
		goto L18
	}
L11:
	;
	v31 = *(*int32)(unsafe.Add(mBase, uint32(v30)))
	if v31 == v28 {
		v42 = v30
		v43 = v22
		goto L10
	} else {
		goto L14
	}
L12:
	;
	goto L13
L13:
	;
	v34 = F_lookup_type_cache(m, v28, int32(_a_F_multirange_union_0))
	mBase = m.M
	v35 = m.ExcPending
	if v35 != 0 {
		goto L1
	} else {
		goto L15
	}
L14:
	;
	goto L13
L15:
	;
	v36 = *(*int32)(unsafe.Add(mBase, uint32(v34)+296))
	if v36 == int32(0) {
		goto L4
	} else {
		goto L16
	}
L16:
	;
	v39 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	*(*int32)(unsafe.Add(mBase, uint32(v39)+16)) = v34
	v41 = *(*int32)(unsafe.Add(mBase, uint32(v15)+8))
	v42 = v34
	v43 = v41
	goto L10
L17:
	;
	v49 = F_palloc_mul(m, int32(4), v43)
	mBase = m.M
	v50 = m.ExcPending
	if v50 != 0 {
		goto L1
	} else {
		goto L20
	}
L18:
	;
	v72 = v44
	v77 = v2
	goto L19
L19:
	;
	v79 = *(*int32)(unsafe.Add(mBase, uint32(v20)+8))
	if int32(0) < v79 {
		goto L25
	} else {
		goto L26
	}
L20:
	;
	v51 = int32(0)
	goto L21
L21:
	;
	v63 = F_multirange_get_range(m, v44, v15, v51)
	mBase = m.M
	v64 = m.ExcPending
	if v64 != 0 {
		goto L1
	} else {
		goto L23
	}
L22:
	;
	v69 = *(*int32)(unsafe.Add(mBase, uint32(v42)+296))
	v72 = v69
	v77 = v49
	goto L19
L23:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v49+v51<<(uint(int32(2))%32)))) = v63
	v67 = v51 + int32(1)
	if v67 != v43 {
		v51 = v67
		goto L21
	} else {
		goto L24
	}
L24:
	;
	goto L22
L25:
	;
	v84 = F_palloc_mul(m, int32(4), v79)
	mBase = m.M
	v85 = m.ExcPending
	if v85 != 0 {
		goto L1
	} else {
		goto L28
	}
L26:
	;
	v112 = v2
	goto L27
L27:
	;
	v113 = v79 + v43
	v116 = F_palloc0(m, v113<<(uint(int32(2))%32))
	mBase = m.M
	v117 = m.ExcPending
	if v117 != 0 {
		goto L1
	} else {
		goto L33
	}
L28:
	;
	v86 = int32(0)
	goto L29
L29:
	;
	v98 = F_multirange_get_range(m, v72, v20, v86)
	mBase = m.M
	v99 = m.ExcPending
	if v99 != 0 {
		goto L1
	} else {
		goto L31
	}
L30:
	;
	v112 = v84
	goto L27
L31:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v84+v86<<(uint(int32(2))%32)))) = v98
	v102 = v86 + int32(1)
	if v102 != v79 {
		v86 = v102
		goto L29
	} else {
		goto L32
	}
L32:
	;
	goto L30
L33:
	;
	v119 = v43 << (uint(int32(2)) % 32)
	if v119 != 0 {
		goto L34
	} else {
		goto L35
	}
L34:
	;
	base.MemoryCopy(m, v116, v77, v119)
	goto L36
L35:
	;
	goto L36
L36:
	;
	v122 = v79 << (uint(int32(2)) % 32)
	if v122 != 0 {
		goto L37
	} else {
		goto L38
	}
L37:
	;
	base.MemoryCopy(m, v116+v119, v112, v122)
	goto L39
L38:
	;
	goto L39
L39:
	;
	v125 = *(*int32)(unsafe.Add(mBase, uint32(v42)))
	v126 = *(*int32)(unsafe.Add(mBase, uint32(v42)+296))
	v127 = F_make_multirange(m, v125, v126, v113, v116)
	mBase = m.M
	v128 = m.ExcPending
	if v128 != 0 {
		goto L1
	} else {
		goto L40
	}
L40:
	;
	v132 = v127
	goto L5
L41:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v12))) = v28
	F_errmsg_internal(m, int32(_a_F_multirange_union_1), v12)
	mBase = m.M
	v150 = m.ExcPending
	if v150 != 0 {
		goto L1
	} else {
		goto L42
	}
L42:
	;
	F_errfinish(m, int32(_a_F_multirange_union_2), int32(561), int32(_a_F_multirange_union_3))
	mBase = m.M
	v155 = m.ExcPending
	if v155 != 0 {
		goto L1
	} else {
		goto L43
	}
L43:
	;
	base.Wasm_trap_unreachable()
	for {
	}
}
func F_multirange_unnest(m *base.Module, l0 int32) int64 {
	mBase := m.M
	_ = mBase
	var v6 int32
	_ = v6
	var v7 int32
	_ = v7
	var v10 int32
	_ = v10
	var v13 int32
	_ = v13
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
	var v21 int32
	_ = v21
	var v23 int32
	_ = v23
	var v24 int32
	_ = v24
	var v28 int32
	_ = v28
	var v30 int32
	_ = v30
	var v31 int32
	_ = v31
	var v40 int32
	_ = v40
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
	var v48 int32
	_ = v48
	var v49 int32
	_ = v49
	var v50 int32
	_ = v50
	var v51 int32
	_ = v51
	var v52 int32
	_ = v52
	var v55 int64
	_ = v55
	var v59 int32
	_ = v59
	var v65 int32
	_ = v65
	var v66 int32
	_ = v66
	var v69 int32
	_ = v69
	v6 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v7 = *(*int32)(unsafe.Add(mBase, uint32(v6)+16))
	if v7 == int32(0) {
		v10 = F_init_MultiFuncCall(m, l0)
		mBase = m.M
		v13 = m.ExcPending
		if v13 != 0 {
			return int64(0)
		} else {
			v14 = int32(_a_F_multirange_unnest_0)
			v15 = *(*int32)(unsafe.Add(mBase, _c_F_multirange_unnest[0]))
			v17 = *(*int32)(unsafe.Add(mBase, uint32(v10)+24))
			*(*int32)(unsafe.Add(mBase, _c_F_multirange_unnest[0])) = v17
			v19 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
			v20 = F_pg_detoast_datum(m, v19)
			mBase = m.M
			v21 = m.ExcPending
			if v21 != 0 {
				return int64(0)
			} else {
				v23 = F_palloc(m, int32(12))
				mBase = m.M
				v24 = m.ExcPending
				if v24 != 0 {
					return int64(0)
				} else {
					*(*int32)(unsafe.Add(mBase, uint32(v23)+8)) = int32(0)
					*(*int32)(unsafe.Add(mBase, uint32(v23))) = v20
					v28 = *(*int32)(unsafe.Add(mBase, uint32(v20)+4))
					v30 = F_lookup_type_cache(m, v28, int32(_a_F_multirange_unnest_1))
					mBase = m.M
					v31 = m.ExcPending
					if v31 != 0 {
						return int64(0)
					} else {
						*(*int32)(unsafe.Add(mBase, uint32(v23)+4)) = v30
						*(*int32)(unsafe.Add(mBase, uint32(v10)+16)) = v23
						*(*int32)(unsafe.Add(mBase, _c_F_multirange_unnest[0])) = v15
						v40 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
						v41 = *(*int32)(unsafe.Add(mBase, uint32(v40)+16))
						v42 = *(*int32)(unsafe.Add(mBase, uint32(v41)+16))
						v43 = *(*int32)(unsafe.Add(mBase, uint32(v42)+8))
						v44 = *(*int32)(unsafe.Add(mBase, uint32(v42)))
						v45 = *(*int32)(unsafe.Add(mBase, uint32(v44)+8))
						if base.Ui32(v43) < base.Ui32(v45) {
							v47 = *(*int32)(unsafe.Add(mBase, uint32(v42)+4))
							v48 = *(*int32)(unsafe.Add(mBase, uint32(v47)+296))
							v49 = F_multirange_get_range(m, v48, v44, v43)
							mBase = m.M
							v50 = m.ExcPending
							if v50 != 0 {
								return int64(0)
							} else {
								v51 = *(*int32)(unsafe.Add(mBase, uint32(v42)+8))
								v52 = int32(1)
								*(*int32)(unsafe.Add(mBase, uint32(v42)+8)) = v51 + v52
								v55 = *(*int64)(unsafe.Add(mBase, uint32(v41)))
								*(*int64)(unsafe.Add(mBase, uint32(v41))) = v55 + int64(1)
								v59 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
								*(*int32)(unsafe.Add(mBase, uint32(v59)+20)) = v52
								return base.I64_extend_i32_u(v49)
							}
						} else {
							F_end_MultiFuncCall(m, l0)
							mBase = m.M
							v65 = m.ExcPending
							if v65 != 0 {
								return int64(0)
							} else {
								v66 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
								*(*int32)(unsafe.Add(mBase, uint32(v66)+20)) = int32(2)
								v69 = int32(1)
								*(*uint8)(unsafe.Add(mBase, uint32(l0)+16)) = uint8(v69)
								return int64(0)
							}
						}
					}
				}
			}
		}
	} else {
		v40 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
		v41 = *(*int32)(unsafe.Add(mBase, uint32(v40)+16))
		v42 = *(*int32)(unsafe.Add(mBase, uint32(v41)+16))
		v43 = *(*int32)(unsafe.Add(mBase, uint32(v42)+8))
		v44 = *(*int32)(unsafe.Add(mBase, uint32(v42)))
		v45 = *(*int32)(unsafe.Add(mBase, uint32(v44)+8))
		if base.Ui32(v43) < base.Ui32(v45) {
			v47 = *(*int32)(unsafe.Add(mBase, uint32(v42)+4))
			v48 = *(*int32)(unsafe.Add(mBase, uint32(v47)+296))
			v49 = F_multirange_get_range(m, v48, v44, v43)
			mBase = m.M
			v50 = m.ExcPending
			if v50 != 0 {
				return int64(0)
			} else {
				v51 = *(*int32)(unsafe.Add(mBase, uint32(v42)+8))
				v52 = int32(1)
				*(*int32)(unsafe.Add(mBase, uint32(v42)+8)) = v51 + v52
				v55 = *(*int64)(unsafe.Add(mBase, uint32(v41)))
				*(*int64)(unsafe.Add(mBase, uint32(v41))) = v55 + int64(1)
				v59 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
				*(*int32)(unsafe.Add(mBase, uint32(v59)+20)) = v52
				return base.I64_extend_i32_u(v49)
			}
		} else {
			F_end_MultiFuncCall(m, l0)
			mBase = m.M
			v65 = m.ExcPending
			if v65 != 0 {
				return int64(0)
			} else {
				v66 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
				*(*int32)(unsafe.Add(mBase, uint32(v66)+20)) = int32(2)
				v69 = int32(1)
				*(*uint8)(unsafe.Add(mBase, uint32(l0)+16)) = uint8(v69)
				return int64(0)
			}
		}
	}
}
func F_multirange_upper_inc(m *base.Module, l0 int32) int64 {
	mBase := m.M
	_ = mBase
	var v8 int32
	_ = v8
	var v10 int32
	_ = v10
	var v12 int32
	_ = v12
	var v13 int32
	_ = v13
	var v16 int32
	_ = v16
	var v17 int32
	_ = v17
	var v18 int32
	_ = v18
	var v19 int32
	_ = v19
	var v20 int32
	_ = v20
	var v21 int32
	_ = v21
	var v24 int32
	_ = v24
	var v25 int32
	_ = v25
	var v26 int32
	_ = v26
	var v29 int32
	_ = v29
	var v31 int32
	_ = v31
	var v32 int32
	_ = v32
	var v33 int32
	_ = v33
	var v34 int32
	_ = v34
	var v42 int32
	_ = v42
	var v43 int64
	_ = v43
	var v48 int64
	_ = v48
	var v56 int32
	_ = v56
	var v60 int32
	_ = v60
	var v65 int32
	_ = v65
	v8 = m.G0
	v10 = v8 - int32(48)
	m.G0 = v10
	v12 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
	v13 = F_pg_detoast_datum(m, v12)
	mBase = m.M
	v16 = m.ExcPending
	if v16 != 0 {
		return int64(0)
	} else {
		v17 = *(*int32)(unsafe.Add(mBase, uint32(v13)+8))
		if v17 != 0 {
			v18 = *(*int32)(unsafe.Add(mBase, uint32(v13)+4))
			v19 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
			v20 = *(*int32)(unsafe.Add(mBase, uint32(v19)+16))
			if v20 != 0 {
				v21 = *(*int32)(unsafe.Add(mBase, uint32(v20)))
				if v21 == v18 {
					v32 = v20
					v33 = v17
					v34 = *(*int32)(unsafe.Add(mBase, uint32(v32)+296))
					F_multirange_get_bounds(m, v34, v13, v33-int32(1), v10+int32(32), v10+int32(16))
					mBase = m.M
					v42 = m.ExcPending
					if v42 != 0 {
						return int64(0)
					} else {
						v43 = int64(*(*uint8)(unsafe.Add(mBase, uint32(v10)+25)))
						v48 = v43
						m.G0 = v10 + int32(48)
						return v48
					}
				} else {
					v24 = F_lookup_type_cache(m, v18, int32(_a_F_multirange_upper_inc_0))
					mBase = m.M
					v25 = m.ExcPending
					if v25 != 0 {
						return int64(0)
					} else {
						v26 = *(*int32)(unsafe.Add(mBase, uint32(v24)+296))
						if v26 == int32(0) {
							F_errstart_cold(m, int32(21), int32(0))
							mBase = m.M
							v56 = m.ExcPending
							if v56 != 0 {
								return int64(0)
							} else {
								*(*int32)(unsafe.Add(mBase, uint32(v10))) = v18
								F_errmsg_internal(m, int32(_a_F_multirange_upper_inc_1), v10)
								mBase = m.M
								v60 = m.ExcPending
								if v60 != 0 {
									return int64(0)
								} else {
									F_errfinish(m, int32(_a_F_multirange_upper_inc_2), int32(561), int32(_a_F_multirange_upper_inc_3))
									mBase = m.M
									v65 = m.ExcPending
									if v65 != 0 {
										return int64(0)
									} else {
										base.Wasm_trap_unreachable()
										for {
										}
									}
								}
							}
						} else {
							v29 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
							*(*int32)(unsafe.Add(mBase, uint32(v29)+16)) = v24
							v31 = *(*int32)(unsafe.Add(mBase, uint32(v13)+8))
							v32 = v24
							v33 = v31
							v34 = *(*int32)(unsafe.Add(mBase, uint32(v32)+296))
							F_multirange_get_bounds(m, v34, v13, v33-int32(1), v10+int32(32), v10+int32(16))
							mBase = m.M
							v42 = m.ExcPending
							if v42 != 0 {
								return int64(0)
							} else {
								v43 = int64(*(*uint8)(unsafe.Add(mBase, uint32(v10)+25)))
								v48 = v43
								m.G0 = v10 + int32(48)
								return v48
							}
						}
					}
				}
			} else {
				v24 = F_lookup_type_cache(m, v18, int32(_a_F_multirange_upper_inc_0))
				mBase = m.M
				v25 = m.ExcPending
				if v25 != 0 {
					return int64(0)
				} else {
					v26 = *(*int32)(unsafe.Add(mBase, uint32(v24)+296))
					if v26 == int32(0) {
						F_errstart_cold(m, int32(21), int32(0))
						mBase = m.M
						v56 = m.ExcPending
						if v56 != 0 {
							return int64(0)
						} else {
							*(*int32)(unsafe.Add(mBase, uint32(v10))) = v18
							F_errmsg_internal(m, int32(_a_F_multirange_upper_inc_1), v10)
							mBase = m.M
							v60 = m.ExcPending
							if v60 != 0 {
								return int64(0)
							} else {
								F_errfinish(m, int32(_a_F_multirange_upper_inc_2), int32(561), int32(_a_F_multirange_upper_inc_3))
								mBase = m.M
								v65 = m.ExcPending
								if v65 != 0 {
									return int64(0)
								} else {
									base.Wasm_trap_unreachable()
									for {
									}
								}
							}
						}
					} else {
						v29 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
						*(*int32)(unsafe.Add(mBase, uint32(v29)+16)) = v24
						v31 = *(*int32)(unsafe.Add(mBase, uint32(v13)+8))
						v32 = v24
						v33 = v31
						v34 = *(*int32)(unsafe.Add(mBase, uint32(v32)+296))
						F_multirange_get_bounds(m, v34, v13, v33-int32(1), v10+int32(32), v10+int32(16))
						mBase = m.M
						v42 = m.ExcPending
						if v42 != 0 {
							return int64(0)
						} else {
							v43 = int64(*(*uint8)(unsafe.Add(mBase, uint32(v10)+25)))
							v48 = v43
							m.G0 = v10 + int32(48)
							return v48
						}
					}
				}
			}
		} else {
			v48 = int64(0)
			m.G0 = v10 + int32(48)
			return v48
		}
	}
}
