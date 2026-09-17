package p0

import (
	base "github.com/shibukawa/pgmem/internal/aot/pgaot/base"
	"unsafe"
)

func F_multirange_after_multirange(m *base.Module, l0 int32) int32 {
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
	var v60 int32
	_ = v60
	var v70 int32
	_ = v70
	var v74 int32
	_ = v74
	var v79 int32
	_ = v79
	v7 = m.G0
	v9 = v7 - int32(48)
	m.G0 = v9
	v11 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v12 = F_pg_detoast_datum(m, v11)
	mBase = m.M
	v15 = m.ExcPending
	if v15 != 0 {
		return int32(0)
	} else {
		v16 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
		v17 = F_pg_detoast_datum(m, v16)
		mBase = m.M
		v18 = m.ExcPending
		if v18 != 0 {
			return int32(0)
		} else {
			v19 = *(*int32)(unsafe.Add(mBase, uint32(v12)+4))
			v20 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
			v21 = *(*int32)(unsafe.Add(mBase, uint32(v20)+16))
			if v21 != 0 {
				v22 = *(*int32)(unsafe.Add(mBase, uint32(v21)))
				if v22 == v19 {
					v32 = v21
					v33 = int32(0)
					v34 = *(*int32)(unsafe.Add(mBase, uint32(v17)+8))
					if v34 == v33 {
						v60 = v33
						m.G0 = v9 + int32(48)
						return v60
					} else {
						v37 = *(*int32)(unsafe.Add(mBase, uint32(v12)+8))
						if v37 == int32(0) {
							v60 = v33
							m.G0 = v9 + int32(48)
							return v60
						} else {
							v40 = *(*int32)(unsafe.Add(mBase, uint32(v32)+296))
							v46 = v9 + int32(32)
							F_multirange_get_bounds(m, v40, v17, v34-int32(1), v9+int32(40), v46)
							mBase = m.M
							v48 = m.ExcPending
							if v48 != 0 {
								return int32(0)
							} else {
								v51 = v9 + int32(24)
								F_multirange_get_bounds(m, v40, v12, int32(0), v51, v9+int32(16))
								mBase = m.M
								v55 = m.ExcPending
								if v55 != 0 {
									return int32(0)
								} else {
									v56 = F_range_cmp_bounds(m, v40, v46, v51)
									mBase = m.M
									v57 = m.ExcPending
									if v57 != 0 {
										return int32(0)
									} else {
										v60 = int32(base.Ui32(v56) >> (uint(int32(31)) % 32))
										m.G0 = v9 + int32(48)
										return v60
									}
								}
							}
						}
					}
				} else {
					v25 = F_lookup_type_cache(m, v19, int32(_a_F_multirange_after_multirange_0))
					mBase = m.M
					v26 = m.ExcPending
					if v26 != 0 {
						return int32(0)
					} else {
						v27 = *(*int32)(unsafe.Add(mBase, uint32(v25)+296))
						if v27 == int32(0) {
							F_errstart_cold(m, int32(21), int32(0))
							mBase = m.M
							v70 = m.ExcPending
							if v70 != 0 {
								return int32(0)
							} else {
								*(*int32)(unsafe.Add(mBase, uint32(v9))) = v19
								F_errmsg_internal(m, int32(_a_F_multirange_after_multirange_1), v9)
								mBase = m.M
								v74 = m.ExcPending
								if v74 != 0 {
									return int32(0)
								} else {
									F_errfinish(m, int32(_a_F_multirange_after_multirange_2), int32(558), int32(_a_F_multirange_after_multirange_3))
									mBase = m.M
									v79 = m.ExcPending
									if v79 != 0 {
										return int32(0)
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
							v33 = int32(0)
							v34 = *(*int32)(unsafe.Add(mBase, uint32(v17)+8))
							if v34 == v33 {
								v60 = v33
								m.G0 = v9 + int32(48)
								return v60
							} else {
								v37 = *(*int32)(unsafe.Add(mBase, uint32(v12)+8))
								if v37 == int32(0) {
									v60 = v33
									m.G0 = v9 + int32(48)
									return v60
								} else {
									v40 = *(*int32)(unsafe.Add(mBase, uint32(v32)+296))
									v46 = v9 + int32(32)
									F_multirange_get_bounds(m, v40, v17, v34-int32(1), v9+int32(40), v46)
									mBase = m.M
									v48 = m.ExcPending
									if v48 != 0 {
										return int32(0)
									} else {
										v51 = v9 + int32(24)
										F_multirange_get_bounds(m, v40, v12, int32(0), v51, v9+int32(16))
										mBase = m.M
										v55 = m.ExcPending
										if v55 != 0 {
											return int32(0)
										} else {
											v56 = F_range_cmp_bounds(m, v40, v46, v51)
											mBase = m.M
											v57 = m.ExcPending
											if v57 != 0 {
												return int32(0)
											} else {
												v60 = int32(base.Ui32(v56) >> (uint(int32(31)) % 32))
												m.G0 = v9 + int32(48)
												return v60
											}
										}
									}
								}
							}
						}
					}
				}
			} else {
				v25 = F_lookup_type_cache(m, v19, int32(_a_F_multirange_after_multirange_0))
				mBase = m.M
				v26 = m.ExcPending
				if v26 != 0 {
					return int32(0)
				} else {
					v27 = *(*int32)(unsafe.Add(mBase, uint32(v25)+296))
					if v27 == int32(0) {
						F_errstart_cold(m, int32(21), int32(0))
						mBase = m.M
						v70 = m.ExcPending
						if v70 != 0 {
							return int32(0)
						} else {
							*(*int32)(unsafe.Add(mBase, uint32(v9))) = v19
							F_errmsg_internal(m, int32(_a_F_multirange_after_multirange_1), v9)
							mBase = m.M
							v74 = m.ExcPending
							if v74 != 0 {
								return int32(0)
							} else {
								F_errfinish(m, int32(_a_F_multirange_after_multirange_2), int32(558), int32(_a_F_multirange_after_multirange_3))
								mBase = m.M
								v79 = m.ExcPending
								if v79 != 0 {
									return int32(0)
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
						v33 = int32(0)
						v34 = *(*int32)(unsafe.Add(mBase, uint32(v17)+8))
						if v34 == v33 {
							v60 = v33
							m.G0 = v9 + int32(48)
							return v60
						} else {
							v37 = *(*int32)(unsafe.Add(mBase, uint32(v12)+8))
							if v37 == int32(0) {
								v60 = v33
								m.G0 = v9 + int32(48)
								return v60
							} else {
								v40 = *(*int32)(unsafe.Add(mBase, uint32(v32)+296))
								v46 = v9 + int32(32)
								F_multirange_get_bounds(m, v40, v17, v34-int32(1), v9+int32(40), v46)
								mBase = m.M
								v48 = m.ExcPending
								if v48 != 0 {
									return int32(0)
								} else {
									v51 = v9 + int32(24)
									F_multirange_get_bounds(m, v40, v12, int32(0), v51, v9+int32(16))
									mBase = m.M
									v55 = m.ExcPending
									if v55 != 0 {
										return int32(0)
									} else {
										v56 = F_range_cmp_bounds(m, v40, v46, v51)
										mBase = m.M
										v57 = m.ExcPending
										if v57 != 0 {
											return int32(0)
										} else {
											v60 = int32(base.Ui32(v56) >> (uint(int32(31)) % 32))
											m.G0 = v9 + int32(48)
											return v60
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
func F_multirange_contains_elem_internal(m *base.Module, l0 int32, l1 int32, l2 int32) int32 {
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
	var v37 int32
	_ = v37
	var v38 int32
	_ = v38
	var v39 int32
	_ = v39
	var v42 int32
	_ = v42
	var v46 int32
	_ = v46
	var v47 int32
	_ = v47
	var v48 int32
	_ = v48
	var v49 int32
	_ = v49
	var v50 int32
	_ = v50
	var v53 int32
	_ = v53
	var v58 int32
	_ = v58
	var v59 int32
	_ = v59
	var v83 int32
	_ = v83
	v10 = m.G0
	v12 = v10 - int32(16)
	m.G0 = v12
	v14 = *(*int32)(unsafe.Add(mBase, uint32(l1)+8))
	if v14 != 0 {
		goto L3
	} else {
		goto L4
	}
L1:
	;
	m.G0 = v12 + int32(16)
	return v83
L2:
	;
	v83 = int32(1)
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
	v83 = int32(0)
	goto L1
L6:
	;
	v28 = int32(base.Ui32(v21+v22) >> (uint(int32(1)) % 32))
	F_multirange_get_bounds(m, l0, l1, v28, v12+int32(8), v12)
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
	v35 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v12)+12)))
	if v35 != 0 {
		goto L11
	} else {
		goto L12
	}
L10:
	;
	if base.Ui32(v59) < base.Ui32(v58) {
		v21 = v58
		v22 = v59
		goto L6
	} else {
		goto L26
	}
L11:
	;
	v46 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v12)+4)))
	if v46 != 0 {
		goto L2
	} else {
		goto L19
	}
L12:
	;
	v36 = *(*int32)(unsafe.Add(mBase, uint32(l0)+208))
	v37 = *(*int32)(unsafe.Add(mBase, uint32(v12)+8))
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
	if int32(0) < v38 {
		goto L14
	} else {
		goto L15
	}
L14:
	;
	v58 = v28
	v59 = v22
	goto L10
L15:
	;
	goto L16
L16:
	;
	if v38 != 0 {
		goto L11
	} else {
		goto L17
	}
L17:
	;
	v42 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v12)+13)))
	if v42&int32(1) != 0 {
		goto L11
	} else {
		goto L18
	}
L18:
	;
	v58 = v28
	v59 = v22
	goto L10
L19:
	;
	v47 = *(*int32)(unsafe.Add(mBase, uint32(l0)+208))
	v48 = *(*int32)(unsafe.Add(mBase, uint32(v12)))
	v49 = F_FunctionCall2Coll(m, v16, v47, v48, l2)
	mBase = m.M
	v50 = m.ExcPending
	if v50 != 0 {
		goto L8
	} else {
		goto L20
	}
L20:
	;
	if int32(0) <= v49 {
		goto L21
	} else {
		goto L22
	}
L21:
	;
	if v49 != 0 {
		goto L2
	} else {
		goto L24
	}
L22:
	;
	goto L23
L23:
	;
	v58 = v21
	v59 = v28 + int32(1)
	goto L10
L24:
	;
	v53 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v12)+5)))
	if v53&int32(1) != 0 {
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
func F_multirange_gt(m *base.Module, l0 int32) int32 {
	var v2 int32
	_ = v2
	var v5 int32
	_ = v5
	v2 = F_multirange_cmp(m, l0)
	v5 = m.ExcPending
	if v5 != 0 {
		return int32(0)
	} else {
		return base.B2i32(int32(0) < v2)
	}
}
func F_multirange_in(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v2 int32
	_ = v2
	var v21 int32
	_ = v21
	var v23 int32
	_ = v23
	var v25 int32
	_ = v25
	var v26 int32
	_ = v26
	var v27 int32
	_ = v27
	var v28 int32
	_ = v28
	var v30 int32
	_ = v30
	var v33 int32
	_ = v33
	var v35 int32
	_ = v35
	var v36 int32
	_ = v36
	var v37 int32
	_ = v37
	var v38 int32
	_ = v38
	var v40 int32
	_ = v40
	var v59 int32
	_ = v59
	var v74 int32
	_ = v74
	var v75 int32
	_ = v75
	var v76 int32
	_ = v76
	var v81 int32
	_ = v81
	var v85 int32
	_ = v85
	var v89 int32
	_ = v89
	var v94 int32
	_ = v94
	var v96 int32
	_ = v96
	var v98 int32
	_ = v98
	var v103 int32
	_ = v103
	var v104 int32
	_ = v104
	var v106 int32
	_ = v106
	var v107 int32
	_ = v107
	var v110 int32
	_ = v110
	var v116 int32
	_ = v116
	var v117 int32
	_ = v117
	var v127 int32
	_ = v127
	var v128 int32
	_ = v128
	var v129 int32
	_ = v129
	var v134 int32
	_ = v134
	var v140 int32
	_ = v140
	var v144 int32
	_ = v144
	var v149 int32
	_ = v149
	var v150 int32
	_ = v150
	var v167 int32
	_ = v167
	var v168 int32
	_ = v168
	var v169 int32
	_ = v169
	var v172 int32
	_ = v172
	var v173 int32
	_ = v173
	var v183 int32
	_ = v183
	var v192 int32
	_ = v192
	var v195 int32
	_ = v195
	var v197 int32
	_ = v197
	var v214 int32
	_ = v214
	var v217 int32
	_ = v217
	var v222 int32
	_ = v222
	var v223 int32
	_ = v223
	var v224 int32
	_ = v224
	var v229 int32
	_ = v229
	var v235 int32
	_ = v235
	var v239 int32
	_ = v239
	var v244 int32
	_ = v244
	var v254 int32
	_ = v254
	var v255 int32
	_ = v255
	var v261 int32
	_ = v261
	var v262 int32
	_ = v262
	var v263 int32
	_ = v263
	var v264 int32
	_ = v264
	var v265 int32
	_ = v265
	var v268 int32
	_ = v268
	var v269 int32
	_ = v269
	var v272 int32
	_ = v272
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
	var v318 int32
	_ = v318
	var v323 int32
	_ = v323
	var v329 int32
	_ = v329
	var v333 int32
	_ = v333
	var v338 int32
	_ = v338
	var v341 int32
	_ = v341
	var v342 int32
	_ = v342
	var v344 int32
	_ = v344
	var v345 int32
	_ = v345
	var v347 int32
	_ = v347
	var v348 int32
	_ = v348
	var v352 int32
	_ = v352
	var v373 int32
	_ = v373
	var v381 int32
	_ = v381
	var v382 int32
	_ = v382
	var v383 int32
	_ = v383
	var v388 int32
	_ = v388
	var v394 int32
	_ = v394
	var v398 int32
	_ = v398
	var v403 int32
	_ = v403
	var v404 int32
	_ = v404
	var v405 int32
	_ = v405
	var v407 int32
	_ = v407
	v2 = int32(0)
	v21 = m.G0
	v23 = v21 - int32(80)
	m.G0 = v23
	v25 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v26 = *(*int32)(unsafe.Add(mBase, uint32(l0)+36))
	v27 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v28 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
	v30 = F_palloc(m, int32(32))
	mBase = m.M
	v33 = m.ExcPending
	if v33 != 0 {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	return int32(0)
L2:
	;
	v35 = F_get_multirange_io_data(m, l0, v28, int32(0))
	mBase = m.M
	v36 = m.ExcPending
	if v36 != 0 {
		goto L1
	} else {
		goto L3
	}
L3:
	;
	v37 = *(*int32)(unsafe.Add(mBase, uint32(v35)))
	v38 = *(*int32)(unsafe.Add(mBase, uint32(v37)+296))
	v40 = v27
	goto L6
L4:
	;
	m.G0 = v23 + int32(80)
	return v407
L5:
	;
	v96 = v40
	v98 = v2
	v103 = v2
	v104 = v30
	v106 = v2
	v107 = v2
	v110 = int32(8)
	goto L20
L6:
	;
	v59 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v40))))
	if base.B2i32(base.Ui32(v59-int32(9)) < base.Ui32(int32(5)))|base.B2i32(v59 == int32(32)) != 0 {
		goto L8
	} else {
		goto L9
	}
L7:
	;
	v74 = int32(0)
	v75 = F_errsave_start(m, v25)
	mBase = m.M
	v76 = m.ExcPending
	if v76 != 0 {
		goto L1
	} else {
		goto L14
	}
L8:
	;
	v40 = v40 + int32(1)
	goto L6
L9:
	;
	if v59 == int32(123) {
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
	if v75 == int32(0) {
		v407 = v74
		goto L4
	} else {
		goto L15
	}
L15:
	;
	F_errcode(m, int32(33685634))
	mBase = m.M
	v81 = m.ExcPending
	if v81 != 0 {
		goto L1
	} else {
		goto L16
	}
L16:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v23))) = v27
	F_errmsg(m, int32(_a_F_multirange_in_0), v23)
	mBase = m.M
	v85 = m.ExcPending
	if v85 != 0 {
		goto L1
	} else {
		goto L17
	}
L17:
	;
	F_errdetail(m, int32(_a_F_multirange_in_1), int32(0))
	mBase = m.M
	v89 = m.ExcPending
	if v89 != 0 {
		goto L1
	} else {
		goto L18
	}
L18:
	;
	F_errsave_finish(m, v25, int32(_a_F_multirange_in_2), int32(152), int32(_a_F_multirange_in_3))
	mBase = m.M
	v94 = m.ExcPending
	if v94 != 0 {
		goto L1
	} else {
		goto L19
	}
L19:
	;
	v407 = v74
	goto L4
L20:
	;
	v116 = v96 + int32(1)
	v117 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v96)+1)))
	if base.B2i32(base.Ui32(v117-int32(9)) < base.Ui32(int32(5)))|base.B2i32(v117 == int32(32)) != 0 {
		v96 = v116
		goto L20
	} else {
		goto L22
	}
L21:
	;
	v352 = v116
	goto L100
L22:
	;
	if v117 == int32(0) {
		goto L23
	} else {
		goto L24
	}
L23:
	;
	v127 = int32(0)
	v128 = F_errsave_start(m, v25)
	mBase = m.M
	v129 = m.ExcPending
	if v129 != 0 {
		goto L1
	} else {
		goto L26
	}
L24:
	;
	goto L25
L25:
	;
	v150 = int32(1)
	switch v98 - v150 {
	case 0:
		goto L38
	case 1:
		v96 = v116
		v98 = v150
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
	if v128 == int32(0) {
		v407 = v127
		goto L4
	} else {
		goto L27
	}
L27:
	;
	F_errcode(m, int32(33685634))
	mBase = m.M
	v134 = m.ExcPending
	if v134 != 0 {
		goto L1
	} else {
		goto L28
	}
L28:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v23)+64)) = v27
	F_errmsg(m, int32(_a_F_multirange_in_0), v23-int32(-64))
	mBase = m.M
	v140 = m.ExcPending
	if v140 != 0 {
		goto L1
	} else {
		goto L29
	}
L29:
	;
	F_errdetail(m, int32(_a_F_multirange_in_4), int32(0))
	mBase = m.M
	v144 = m.ExcPending
	if v144 != 0 {
		goto L1
	} else {
		goto L30
	}
L30:
	;
	F_errsave_finish(m, v25, int32(_a_F_multirange_in_2), int32(165), int32(_a_F_multirange_in_3))
	mBase = m.M
	v149 = m.ExcPending
	if v149 != 0 {
		goto L1
	} else {
		goto L31
	}
L31:
	;
	v407 = v127
	goto L4
L32:
	;
	goto L21
L33:
	;
	v96 = v341
	v98 = v342
	v103 = v344
	v104 = v345
	v107 = v347
	v110 = v348
	goto L20
L34:
	;
	v341 = v116
	v342 = int32(1)
	v344 = v103
	v345 = v104
	v347 = v107
	v348 = v110
	goto L33
L35:
	;
	v341 = v116
	v342 = int32(3)
	v344 = v103
	v345 = v104
	v347 = v107
	v348 = v110
	goto L33
L36:
	;
	if v117 == int32(44) {
		v96 = v116
		v98 = int32(0)
		goto L20
	} else {
		goto L92
	}
L37:
	;
	if v117 != int32(92) {
		goto L82
	} else {
		goto L83
	}
L38:
	;
	switch v117 - int32(34) {
	case 0:
		v96 = v116
		v98 = int32(3)
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
	if base.B2i32(v117 == int32(40))|base.B2i32(v117 == int32(91)) != 0 {
		v96 = v116
		v98 = v150
		v106 = v116
		goto L20
	} else {
		goto L40
	}
L40:
	;
	if base.B2i32(v107 == int32(0))&base.B2i32(v117 == int32(125)) != 0 {
		goto L32
	} else {
		goto L41
	}
L41:
	;
	v167 = v116
	v168 = int32(_a_F_multirange_in_5)
	v169 = int32(5)
	goto L43
L42:
	;
	if v214 == int32(0) {
		goto L58
	} else {
		goto L59
	}
L43:
	;
	if v169 != 0 {
		goto L45
	} else {
		goto L46
	}
L44:
	;
	v214 = int32(0)
	goto L42
L45:
	;
	v172 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v167))))
	v173 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v168))))
	if v172 == v173 {
		v195 = v172
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
	v197 = int32(1)
	if v195 != 0 {
		v167 = v167 + v197
		v168 = v168 + v197
		v169 = v169 - v197
		goto L43
	} else {
		goto L57
	}
L49:
	;
	if base.Ui32((v172-int32(65))&int32(255)) < base.Ui32(int32(26)) {
		goto L50
	} else {
		goto L51
	}
L50:
	;
	v183 = v172 | int32(32)
	goto L52
L51:
	;
	v183 = v172
	goto L52
L52:
	;
	if base.Ui32((v173-int32(65))&int32(255)) < base.Ui32(int32(26)) {
		goto L53
	} else {
		goto L54
	}
L53:
	;
	v192 = v173 | int32(32)
	goto L55
L54:
	;
	v192 = v173
	goto L55
L55:
	;
	if v183 == v192 {
		v195 = v183
		goto L48
	} else {
		goto L56
	}
L56:
	;
	v214 = v183 - v192
	goto L42
L57:
	;
	goto L47
L58:
	;
	v217 = int32(5)
	v341 = v96 + v217
	v342 = v217
	v344 = v103
	v345 = v104
	v347 = v107 + int32(1)
	v348 = v110
	goto L33
L59:
	;
	goto L60
L60:
	;
	v222 = int32(0)
	v223 = F_errsave_start(m, v25)
	mBase = m.M
	v224 = m.ExcPending
	if v224 != 0 {
		goto L1
	} else {
		goto L61
	}
L61:
	;
	if v223 == int32(0) {
		v407 = v222
		goto L4
	} else {
		goto L62
	}
L62:
	;
	F_errcode(m, int32(33685634))
	mBase = m.M
	v229 = m.ExcPending
	if v229 != 0 {
		goto L1
	} else {
		goto L63
	}
L63:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v23)+32)) = v27
	F_errmsg(m, int32(_a_F_multirange_in_0), v23+int32(32))
	mBase = m.M
	v235 = m.ExcPending
	if v235 != 0 {
		goto L1
	} else {
		goto L64
	}
L64:
	;
	F_errdetail(m, int32(_a_F_multirange_in_6), int32(0))
	mBase = m.M
	v239 = m.ExcPending
	if v239 != 0 {
		goto L1
	} else {
		goto L65
	}
L65:
	;
	F_errsave_finish(m, v25, int32(_a_F_multirange_in_2), int32(194), int32(_a_F_multirange_in_3))
	mBase = m.M
	v244 = m.ExcPending
	if v244 != 0 {
		goto L1
	} else {
		goto L66
	}
L66:
	;
	v407 = v222
	goto L4
L67:
	;
	v254 = F_pnstrdup(m, v106, v116-v106+int32(1))
	mBase = m.M
	v255 = m.ExcPending
	if v255 != 0 {
		goto L1
	} else {
		goto L70
	}
L68:
	;
	switch v117 - int32(92) {
	case 0:
		goto L69
	case 1:
		goto L67
	default:
		goto L34
	}
L69:
	;
	v341 = v116
	v342 = int32(2)
	v344 = v103
	v345 = v104
	v347 = v107
	v348 = v110
	goto L33
L70:
	;
	if v103 == v110 {
		goto L71
	} else {
		goto L72
	}
L71:
	;
	v261 = F_repalloc(m, v104, v103<<(uint(int32(3))%32))
	mBase = m.M
	v262 = m.ExcPending
	if v262 != 0 {
		goto L1
	} else {
		goto L74
	}
L72:
	;
	v263 = v104
	v264 = v110
	goto L73
L73:
	;
	v265 = *(*int32)(unsafe.Add(mBase, uint32(v35)+32))
	v268 = F_InputFunctionCallSafe(m, v35+int32(4), v254, v265, v26, v25, v23+int32(76))
	mBase = m.M
	v269 = m.ExcPending
	if v269 != 0 {
		goto L1
	} else {
		goto L75
	}
L74:
	;
	v263 = v261
	v264 = v103 << (uint(int32(1)) % 32)
	goto L73
L75:
	;
	if v268 == int32(0) {
		goto L76
	} else {
		goto L77
	}
L76:
	;
	v272 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(l0)+16)) = uint8(v272)
	v407 = int32(0)
	goto L4
L77:
	;
	goto L78
L78:
	;
	v276 = v107 + int32(1)
	v277 = int32(5)
	v278 = *(*int32)(unsafe.Add(mBase, uint32(v23)+76))
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
		v96 = v116
		v98 = v277
		v104 = v263
		v107 = v276
		v110 = v264
		goto L20
	} else {
		goto L81
	}
L81:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v263+v103<<(uint(int32(2))%32)))) = v279
	v341 = v116
	v342 = v277
	v344 = v103 + int32(1)
	v345 = v263
	v347 = v276
	v348 = v264
	goto L33
L82:
	;
	if v117 != int32(34) {
		v96 = v116
		v98 = int32(3)
		goto L20
	} else {
		goto L85
	}
L83:
	;
	goto L84
L84:
	;
	v341 = v116
	v342 = int32(4)
	v344 = v103
	v345 = v104
	v347 = v107
	v348 = v110
	goto L33
L85:
	;
	v303 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v96)+2)))
	v305 = base.B2i32(v303 == int32(34))
	if v303 == int32(34) {
		goto L86
	} else {
		goto L87
	}
L86:
	;
	v306 = v96 + int32(2)
	goto L88
L87:
	;
	v306 = v116
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
	v341 = v306
	v342 = v309
	v344 = v103
	v345 = v104
	v347 = v107
	v348 = v110
	goto L33
L92:
	;
	if v117 == int32(125) {
		goto L32
	} else {
		goto L93
	}
L93:
	;
	v316 = int32(0)
	v317 = F_errsave_start(m, v25)
	mBase = m.M
	v318 = m.ExcPending
	if v318 != 0 {
		goto L1
	} else {
		goto L94
	}
L94:
	;
	if v317 == int32(0) {
		v407 = v316
		goto L4
	} else {
		goto L95
	}
L95:
	;
	F_errcode(m, int32(33685634))
	mBase = m.M
	v323 = m.ExcPending
	if v323 != 0 {
		goto L1
	} else {
		goto L96
	}
L96:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v23)+48)) = v27
	F_errmsg(m, int32(_a_F_multirange_in_0), v23+int32(48))
	mBase = m.M
	v329 = m.ExcPending
	if v329 != 0 {
		goto L1
	} else {
		goto L97
	}
L97:
	;
	F_errdetail(m, int32(_a_F_multirange_in_7), int32(0))
	mBase = m.M
	v333 = m.ExcPending
	if v333 != 0 {
		goto L1
	} else {
		goto L98
	}
L98:
	;
	F_errsave_finish(m, v25, int32(_a_F_multirange_in_2), int32(268), int32(_a_F_multirange_in_3))
	mBase = m.M
	v338 = m.ExcPending
	if v338 != 0 {
		goto L1
	} else {
		goto L99
	}
L99:
	;
	v407 = v316
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
	v381 = int32(0)
	v382 = F_errsave_start(m, v25)
	mBase = m.M
	v383 = m.ExcPending
	if v383 != 0 {
		goto L1
	} else {
		goto L106
	}
L104:
	;
	goto L105
L105:
	;
	v404 = F_make_multirange(m, v28, v38, v103, v104)
	mBase = m.M
	v405 = m.ExcPending
	if v405 != 0 {
		goto L1
	} else {
		goto L112
	}
L106:
	;
	if v382 == int32(0) {
		v407 = v381
		goto L4
	} else {
		goto L107
	}
L107:
	;
	F_errcode(m, int32(33685634))
	mBase = m.M
	v388 = m.ExcPending
	if v388 != 0 {
		goto L1
	} else {
		goto L108
	}
L108:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v23)+16)) = v27
	F_errmsg(m, int32(_a_F_multirange_in_0), v23+int32(16))
	mBase = m.M
	v394 = m.ExcPending
	if v394 != 0 {
		goto L1
	} else {
		goto L109
	}
L109:
	;
	F_errdetail(m, int32(_a_F_multirange_in_8), int32(0))
	mBase = m.M
	v398 = m.ExcPending
	if v398 != 0 {
		goto L1
	} else {
		goto L110
	}
L110:
	;
	F_errsave_finish(m, v25, int32(_a_F_multirange_in_2), int32(292), int32(_a_F_multirange_in_3))
	mBase = m.M
	v403 = m.ExcPending
	if v403 != 0 {
		goto L1
	} else {
		goto L111
	}
L111:
	;
	v407 = v381
	goto L4
L112:
	;
	v407 = v404
	goto L4
}
func F_multirange_le(m *base.Module, l0 int32) int32 {
	var v2 int32
	_ = v2
	var v5 int32
	_ = v5
	v2 = F_multirange_cmp(m, l0)
	v5 = m.ExcPending
	if v5 != 0 {
		return int32(0)
	} else {
		return base.B2i32(v2 <= int32(0))
	}
}
func F_multirange_ne(m *base.Module, l0 int32) int32 {
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
	var v45 int32
	_ = v45
	var v49 int32
	_ = v49
	var v54 int32
	_ = v54
	v7 = m.G0
	v9 = v7 - int32(16)
	m.G0 = v9
	v11 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v12 = F_pg_detoast_datum(m, v11)
	mBase = m.M
	v15 = m.ExcPending
	if v15 != 0 {
		return int32(0)
	} else {
		v16 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
		v17 = F_pg_detoast_datum(m, v16)
		mBase = m.M
		v18 = m.ExcPending
		if v18 != 0 {
			return int32(0)
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
						return int32(0)
					} else {
						m.G0 = v9 + int32(16)
						return v34 ^ int32(1)
					}
				} else {
					v25 = F_lookup_type_cache(m, v19, int32(_a_F_multirange_ne_0))
					mBase = m.M
					v26 = m.ExcPending
					if v26 != 0 {
						return int32(0)
					} else {
						v27 = *(*int32)(unsafe.Add(mBase, uint32(v25)+296))
						if v27 == int32(0) {
							F_errstart_cold(m, int32(21), int32(0))
							mBase = m.M
							v45 = m.ExcPending
							if v45 != 0 {
								return int32(0)
							} else {
								*(*int32)(unsafe.Add(mBase, uint32(v9))) = v19
								F_errmsg_internal(m, int32(_a_F_multirange_ne_1), v9)
								mBase = m.M
								v49 = m.ExcPending
								if v49 != 0 {
									return int32(0)
								} else {
									F_errfinish(m, int32(_a_F_multirange_ne_2), int32(558), int32(_a_F_multirange_ne_3))
									mBase = m.M
									v54 = m.ExcPending
									if v54 != 0 {
										return int32(0)
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
								return int32(0)
							} else {
								m.G0 = v9 + int32(16)
								return v34 ^ int32(1)
							}
						}
					}
				}
			} else {
				v25 = F_lookup_type_cache(m, v19, int32(_a_F_multirange_ne_0))
				mBase = m.M
				v26 = m.ExcPending
				if v26 != 0 {
					return int32(0)
				} else {
					v27 = *(*int32)(unsafe.Add(mBase, uint32(v25)+296))
					if v27 == int32(0) {
						F_errstart_cold(m, int32(21), int32(0))
						mBase = m.M
						v45 = m.ExcPending
						if v45 != 0 {
							return int32(0)
						} else {
							*(*int32)(unsafe.Add(mBase, uint32(v9))) = v19
							F_errmsg_internal(m, int32(_a_F_multirange_ne_1), v9)
							mBase = m.M
							v49 = m.ExcPending
							if v49 != 0 {
								return int32(0)
							} else {
								F_errfinish(m, int32(_a_F_multirange_ne_2), int32(558), int32(_a_F_multirange_ne_3))
								mBase = m.M
								v54 = m.ExcPending
								if v54 != 0 {
									return int32(0)
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
							return int32(0)
						} else {
							m.G0 = v9 + int32(16)
							return v34 ^ int32(1)
						}
					}
				}
			}
		}
	}
}
func F_multirange_overleft_range(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v2 int32
	_ = v2
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
	var v22 int32
	_ = v22
	var v28 int32
	_ = v28
	var v31 int32
	_ = v31
	var v32 int32
	_ = v32
	var v33 int32
	_ = v33
	var v34 int32
	_ = v34
	var v37 int32
	_ = v37
	var v38 int32
	_ = v38
	var v39 int32
	_ = v39
	var v42 int32
	_ = v42
	var v44 int32
	_ = v44
	var v45 int32
	_ = v45
	var v46 int32
	_ = v46
	var v52 int32
	_ = v52
	var v54 int32
	_ = v54
	var v55 int32
	_ = v55
	var v59 int32
	_ = v59
	var v63 int32
	_ = v63
	var v64 int32
	_ = v64
	var v65 int32
	_ = v65
	var v66 int32
	_ = v66
	var v70 int32
	_ = v70
	var v80 int32
	_ = v80
	var v84 int32
	_ = v84
	var v89 int32
	_ = v89
	v2 = int32(0)
	v7 = m.G0
	v9 = v7 - int32(48)
	m.G0 = v9
	v11 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v12 = F_pg_detoast_datum(m, v11)
	mBase = m.M
	v15 = m.ExcPending
	if v15 != 0 {
		return int32(0)
	} else {
		v16 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
		v17 = F_pg_detoast_datum(m, v16)
		mBase = m.M
		v18 = m.ExcPending
		if v18 != 0 {
			return int32(0)
		} else {
			v19 = *(*int32)(unsafe.Add(mBase, uint32(v12)+8))
			if v19 == int32(0) {
				v70 = v2
				m.G0 = v9 + int32(48)
				return v70
			} else {
				v22 = *(*int32)(unsafe.Add(mBase, uint32(v17)))
				v28 = int32(*(*int8)(unsafe.Add(mBase, uint32(v17+int32(base.Ui32(v22)>>(uint(int32(2))%32))-int32(1)))))
				if v28&int32(1) != 0 {
					v70 = v2
					m.G0 = v9 + int32(48)
					return v70
				} else {
					v31 = *(*int32)(unsafe.Add(mBase, uint32(v12)+4))
					v32 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
					v33 = *(*int32)(unsafe.Add(mBase, uint32(v32)+16))
					if v33 != 0 {
						v34 = *(*int32)(unsafe.Add(mBase, uint32(v33)))
						if v34 == v31 {
							v44 = v33
							v45 = *(*int32)(unsafe.Add(mBase, uint32(v44)+296))
							v46 = *(*int32)(unsafe.Add(mBase, uint32(v12)+8))
							v52 = v9 + int32(32)
							F_multirange_get_bounds(m, v45, v12, v46-int32(1), v9+int32(40), v52)
							mBase = m.M
							v54 = m.ExcPending
							if v54 != 0 {
								return int32(0)
							} else {
								v55 = *(*int32)(unsafe.Add(mBase, uint32(v44)+296))
								v59 = v9 + int32(16)
								F_range_deserialize(m, v55, v17, v9+int32(24), v59, v9+int32(15))
								mBase = m.M
								v63 = m.ExcPending
								if v63 != 0 {
									return int32(0)
								} else {
									v64 = *(*int32)(unsafe.Add(mBase, uint32(v44)+296))
									v65 = F_range_cmp_bounds(m, v64, v52, v59)
									mBase = m.M
									v66 = m.ExcPending
									if v66 != 0 {
										return int32(0)
									} else {
										v70 = base.B2i32(v65 <= int32(0))
										m.G0 = v9 + int32(48)
										return v70
									}
								}
							}
						} else {
							v37 = F_lookup_type_cache(m, v31, int32(_a_F_multirange_overleft_range_0))
							mBase = m.M
							v38 = m.ExcPending
							if v38 != 0 {
								return int32(0)
							} else {
								v39 = *(*int32)(unsafe.Add(mBase, uint32(v37)+296))
								if v39 == int32(0) {
									F_errstart_cold(m, int32(21), int32(0))
									mBase = m.M
									v80 = m.ExcPending
									if v80 != 0 {
										return int32(0)
									} else {
										*(*int32)(unsafe.Add(mBase, uint32(v9))) = v31
										F_errmsg_internal(m, int32(_a_F_multirange_overleft_range_1), v9)
										mBase = m.M
										v84 = m.ExcPending
										if v84 != 0 {
											return int32(0)
										} else {
											F_errfinish(m, int32(_a_F_multirange_overleft_range_2), int32(558), int32(_a_F_multirange_overleft_range_3))
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
								} else {
									v42 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
									*(*int32)(unsafe.Add(mBase, uint32(v42)+16)) = v37
									v44 = v37
									v45 = *(*int32)(unsafe.Add(mBase, uint32(v44)+296))
									v46 = *(*int32)(unsafe.Add(mBase, uint32(v12)+8))
									v52 = v9 + int32(32)
									F_multirange_get_bounds(m, v45, v12, v46-int32(1), v9+int32(40), v52)
									mBase = m.M
									v54 = m.ExcPending
									if v54 != 0 {
										return int32(0)
									} else {
										v55 = *(*int32)(unsafe.Add(mBase, uint32(v44)+296))
										v59 = v9 + int32(16)
										F_range_deserialize(m, v55, v17, v9+int32(24), v59, v9+int32(15))
										mBase = m.M
										v63 = m.ExcPending
										if v63 != 0 {
											return int32(0)
										} else {
											v64 = *(*int32)(unsafe.Add(mBase, uint32(v44)+296))
											v65 = F_range_cmp_bounds(m, v64, v52, v59)
											mBase = m.M
											v66 = m.ExcPending
											if v66 != 0 {
												return int32(0)
											} else {
												v70 = base.B2i32(v65 <= int32(0))
												m.G0 = v9 + int32(48)
												return v70
											}
										}
									}
								}
							}
						}
					} else {
						v37 = F_lookup_type_cache(m, v31, int32(_a_F_multirange_overleft_range_0))
						mBase = m.M
						v38 = m.ExcPending
						if v38 != 0 {
							return int32(0)
						} else {
							v39 = *(*int32)(unsafe.Add(mBase, uint32(v37)+296))
							if v39 == int32(0) {
								F_errstart_cold(m, int32(21), int32(0))
								mBase = m.M
								v80 = m.ExcPending
								if v80 != 0 {
									return int32(0)
								} else {
									*(*int32)(unsafe.Add(mBase, uint32(v9))) = v31
									F_errmsg_internal(m, int32(_a_F_multirange_overleft_range_1), v9)
									mBase = m.M
									v84 = m.ExcPending
									if v84 != 0 {
										return int32(0)
									} else {
										F_errfinish(m, int32(_a_F_multirange_overleft_range_2), int32(558), int32(_a_F_multirange_overleft_range_3))
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
							} else {
								v42 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
								*(*int32)(unsafe.Add(mBase, uint32(v42)+16)) = v37
								v44 = v37
								v45 = *(*int32)(unsafe.Add(mBase, uint32(v44)+296))
								v46 = *(*int32)(unsafe.Add(mBase, uint32(v12)+8))
								v52 = v9 + int32(32)
								F_multirange_get_bounds(m, v45, v12, v46-int32(1), v9+int32(40), v52)
								mBase = m.M
								v54 = m.ExcPending
								if v54 != 0 {
									return int32(0)
								} else {
									v55 = *(*int32)(unsafe.Add(mBase, uint32(v44)+296))
									v59 = v9 + int32(16)
									F_range_deserialize(m, v55, v17, v9+int32(24), v59, v9+int32(15))
									mBase = m.M
									v63 = m.ExcPending
									if v63 != 0 {
										return int32(0)
									} else {
										v64 = *(*int32)(unsafe.Add(mBase, uint32(v44)+296))
										v65 = F_range_cmp_bounds(m, v64, v52, v59)
										mBase = m.M
										v66 = m.ExcPending
										if v66 != 0 {
											return int32(0)
										} else {
											v70 = base.B2i32(v65 <= int32(0))
											m.G0 = v9 + int32(48)
											return v70
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
func F_multirange_overright_multirange(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v2 int32
	_ = v2
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
	var v22 int32
	_ = v22
	var v25 int32
	_ = v25
	var v26 int32
	_ = v26
	var v27 int32
	_ = v27
	var v28 int32
	_ = v28
	var v31 int32
	_ = v31
	var v32 int32
	_ = v32
	var v33 int32
	_ = v33
	var v36 int32
	_ = v36
	var v38 int32
	_ = v38
	var v39 int32
	_ = v39
	var v42 int32
	_ = v42
	var v46 int32
	_ = v46
	var v47 int32
	_ = v47
	var v50 int32
	_ = v50
	var v54 int32
	_ = v54
	var v55 int32
	_ = v55
	var v56 int32
	_ = v56
	var v57 int32
	_ = v57
	var v63 int32
	_ = v63
	var v73 int32
	_ = v73
	var v77 int32
	_ = v77
	var v82 int32
	_ = v82
	v2 = int32(0)
	v7 = m.G0
	v9 = v7 - int32(48)
	m.G0 = v9
	v11 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v12 = F_pg_detoast_datum(m, v11)
	mBase = m.M
	v15 = m.ExcPending
	if v15 != 0 {
		return int32(0)
	} else {
		v16 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
		v17 = F_pg_detoast_datum(m, v16)
		mBase = m.M
		v18 = m.ExcPending
		if v18 != 0 {
			return int32(0)
		} else {
			v19 = *(*int32)(unsafe.Add(mBase, uint32(v12)+8))
			if v19 == int32(0) {
				v63 = v2
				m.G0 = v9 + int32(48)
				return v63
			} else {
				v22 = *(*int32)(unsafe.Add(mBase, uint32(v17)+8))
				if v22 == int32(0) {
					v63 = v2
					m.G0 = v9 + int32(48)
					return v63
				} else {
					v25 = *(*int32)(unsafe.Add(mBase, uint32(v12)+4))
					v26 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
					v27 = *(*int32)(unsafe.Add(mBase, uint32(v26)+16))
					if v27 != 0 {
						v28 = *(*int32)(unsafe.Add(mBase, uint32(v27)))
						if v28 == v25 {
							v38 = v27
							v39 = *(*int32)(unsafe.Add(mBase, uint32(v38)+296))
							v42 = v9 + int32(40)
							F_multirange_get_bounds(m, v39, v12, int32(0), v42, v9+int32(32))
							mBase = m.M
							v46 = m.ExcPending
							if v46 != 0 {
								return int32(0)
							} else {
								v47 = *(*int32)(unsafe.Add(mBase, uint32(v38)+296))
								v50 = v9 + int32(24)
								F_multirange_get_bounds(m, v47, v17, int32(0), v50, v9+int32(16))
								mBase = m.M
								v54 = m.ExcPending
								if v54 != 0 {
									return int32(0)
								} else {
									v55 = *(*int32)(unsafe.Add(mBase, uint32(v38)+296))
									v56 = F_range_cmp_bounds(m, v55, v42, v50)
									mBase = m.M
									v57 = m.ExcPending
									if v57 != 0 {
										return int32(0)
									} else {
										v63 = int32(base.Ui32(v56^int32(-1)) >> (uint(int32(31)) % 32))
										m.G0 = v9 + int32(48)
										return v63
									}
								}
							}
						} else {
							v31 = F_lookup_type_cache(m, v25, int32(_a_F_multirange_overright_multirange_0))
							mBase = m.M
							v32 = m.ExcPending
							if v32 != 0 {
								return int32(0)
							} else {
								v33 = *(*int32)(unsafe.Add(mBase, uint32(v31)+296))
								if v33 == int32(0) {
									F_errstart_cold(m, int32(21), int32(0))
									mBase = m.M
									v73 = m.ExcPending
									if v73 != 0 {
										return int32(0)
									} else {
										*(*int32)(unsafe.Add(mBase, uint32(v9))) = v25
										F_errmsg_internal(m, int32(_a_F_multirange_overright_multirange_1), v9)
										mBase = m.M
										v77 = m.ExcPending
										if v77 != 0 {
											return int32(0)
										} else {
											F_errfinish(m, int32(_a_F_multirange_overright_multirange_2), int32(558), int32(_a_F_multirange_overright_multirange_3))
											mBase = m.M
											v82 = m.ExcPending
											if v82 != 0 {
												return int32(0)
											} else {
												base.Wasm_trap_unreachable()
												for {
												}
											}
										}
									}
								} else {
									v36 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
									*(*int32)(unsafe.Add(mBase, uint32(v36)+16)) = v31
									v38 = v31
									v39 = *(*int32)(unsafe.Add(mBase, uint32(v38)+296))
									v42 = v9 + int32(40)
									F_multirange_get_bounds(m, v39, v12, int32(0), v42, v9+int32(32))
									mBase = m.M
									v46 = m.ExcPending
									if v46 != 0 {
										return int32(0)
									} else {
										v47 = *(*int32)(unsafe.Add(mBase, uint32(v38)+296))
										v50 = v9 + int32(24)
										F_multirange_get_bounds(m, v47, v17, int32(0), v50, v9+int32(16))
										mBase = m.M
										v54 = m.ExcPending
										if v54 != 0 {
											return int32(0)
										} else {
											v55 = *(*int32)(unsafe.Add(mBase, uint32(v38)+296))
											v56 = F_range_cmp_bounds(m, v55, v42, v50)
											mBase = m.M
											v57 = m.ExcPending
											if v57 != 0 {
												return int32(0)
											} else {
												v63 = int32(base.Ui32(v56^int32(-1)) >> (uint(int32(31)) % 32))
												m.G0 = v9 + int32(48)
												return v63
											}
										}
									}
								}
							}
						}
					} else {
						v31 = F_lookup_type_cache(m, v25, int32(_a_F_multirange_overright_multirange_0))
						mBase = m.M
						v32 = m.ExcPending
						if v32 != 0 {
							return int32(0)
						} else {
							v33 = *(*int32)(unsafe.Add(mBase, uint32(v31)+296))
							if v33 == int32(0) {
								F_errstart_cold(m, int32(21), int32(0))
								mBase = m.M
								v73 = m.ExcPending
								if v73 != 0 {
									return int32(0)
								} else {
									*(*int32)(unsafe.Add(mBase, uint32(v9))) = v25
									F_errmsg_internal(m, int32(_a_F_multirange_overright_multirange_1), v9)
									mBase = m.M
									v77 = m.ExcPending
									if v77 != 0 {
										return int32(0)
									} else {
										F_errfinish(m, int32(_a_F_multirange_overright_multirange_2), int32(558), int32(_a_F_multirange_overright_multirange_3))
										mBase = m.M
										v82 = m.ExcPending
										if v82 != 0 {
											return int32(0)
										} else {
											base.Wasm_trap_unreachable()
											for {
											}
										}
									}
								}
							} else {
								v36 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
								*(*int32)(unsafe.Add(mBase, uint32(v36)+16)) = v31
								v38 = v31
								v39 = *(*int32)(unsafe.Add(mBase, uint32(v38)+296))
								v42 = v9 + int32(40)
								F_multirange_get_bounds(m, v39, v12, int32(0), v42, v9+int32(32))
								mBase = m.M
								v46 = m.ExcPending
								if v46 != 0 {
									return int32(0)
								} else {
									v47 = *(*int32)(unsafe.Add(mBase, uint32(v38)+296))
									v50 = v9 + int32(24)
									F_multirange_get_bounds(m, v47, v17, int32(0), v50, v9+int32(16))
									mBase = m.M
									v54 = m.ExcPending
									if v54 != 0 {
										return int32(0)
									} else {
										v55 = *(*int32)(unsafe.Add(mBase, uint32(v38)+296))
										v56 = F_range_cmp_bounds(m, v55, v42, v50)
										mBase = m.M
										v57 = m.ExcPending
										if v57 != 0 {
											return int32(0)
										} else {
											v63 = int32(base.Ui32(v56^int32(-1)) >> (uint(int32(31)) % 32))
											m.G0 = v9 + int32(48)
											return v63
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
func F_multirange_typanalyze(m *base.Module, l0 int32) int32 {
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
	v4 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v5 = *(*int32)(unsafe.Add(mBase, uint32(v4)+4))
	v6 = F_getBaseType(m, v5)
	mBase = m.M
	v9 = m.ExcPending
	if v9 != 0 {
		return int32(0)
	} else {
		v10 = F_multirange_get_typcache(m, l0, v6)
		mBase = m.M
		v11 = m.ExcPending
		if v11 != 0 {
			return int32(0)
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
			*(*int32)(unsafe.Add(mBase, uint32(v4)+24)) = int32(1474)
			*(*int32)(unsafe.Add(mBase, uint32(v4)+28)) = v18 * int32(300)
			return int32(1)
		}
	}
}
func F_multirange_union(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v2 int32
	_ = v2
	var v11 int32
	_ = v11
	var v13 int32
	_ = v13
	var v15 int32
	_ = v15
	var v16 int32
	_ = v16
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
	var v29 int32
	_ = v29
	var v30 int32
	_ = v30
	var v31 int32
	_ = v31
	var v32 int32
	_ = v32
	var v35 int32
	_ = v35
	var v36 int32
	_ = v36
	var v37 int32
	_ = v37
	var v40 int32
	_ = v40
	var v42 int32
	_ = v42
	var v43 int32
	_ = v43
	var v44 int32
	_ = v44
	var v45 int32
	_ = v45
	var v51 int32
	_ = v51
	var v52 int32
	_ = v52
	var v53 int32
	_ = v53
	var v66 int32
	_ = v66
	var v67 int32
	_ = v67
	var v70 int32
	_ = v70
	var v72 int32
	_ = v72
	var v76 int32
	_ = v76
	var v81 int32
	_ = v81
	var v83 int32
	_ = v83
	var v85 int32
	_ = v85
	var v89 int32
	_ = v89
	var v90 int32
	_ = v90
	var v91 int32
	_ = v91
	var v104 int32
	_ = v104
	var v105 int32
	_ = v105
	var v108 int32
	_ = v108
	var v119 int32
	_ = v119
	var v120 int32
	_ = v120
	var v123 int32
	_ = v123
	var v124 int32
	_ = v124
	var v126 int32
	_ = v126
	var v130 int32
	_ = v130
	var v131 int32
	_ = v131
	var v132 int32
	_ = v132
	var v133 int32
	_ = v133
	var v136 int32
	_ = v136
	var v151 int32
	_ = v151
	var v155 int32
	_ = v155
	var v160 int32
	_ = v160
	v2 = int32(0)
	v11 = m.G0
	v13 = v11 - int32(16)
	m.G0 = v13
	v15 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v16 = F_pg_detoast_datum(m, v15)
	mBase = m.M
	v19 = m.ExcPending
	if v19 != 0 {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	return int32(0)
L2:
	;
	v20 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
	v21 = F_pg_detoast_datum(m, v20)
	mBase = m.M
	v22 = m.ExcPending
	if v22 != 0 {
		goto L1
	} else {
		goto L3
	}
L3:
	;
	v23 = *(*int32)(unsafe.Add(mBase, uint32(v16)+8))
	if v23 == int32(0) {
		v136 = v21
		goto L5
	} else {
		goto L6
	}
L4:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v151 = m.ExcPending
	if v151 != 0 {
		goto L1
	} else {
		goto L41
	}
L5:
	;
	m.G0 = v13 + int32(16)
	return v136
L6:
	;
	v26 = *(*int32)(unsafe.Add(mBase, uint32(v21)+8))
	if v26 == int32(0) {
		goto L7
	} else {
		goto L8
	}
L7:
	;
	v136 = v16
	goto L5
L8:
	;
	goto L9
L9:
	;
	v29 = *(*int32)(unsafe.Add(mBase, uint32(v16)+4))
	v30 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v31 = *(*int32)(unsafe.Add(mBase, uint32(v30)+16))
	if v31 != 0 {
		goto L11
	} else {
		goto L12
	}
L10:
	;
	v45 = *(*int32)(unsafe.Add(mBase, uint32(v43)+296))
	if int32(0) < v44 {
		goto L17
	} else {
		goto L18
	}
L11:
	;
	v32 = *(*int32)(unsafe.Add(mBase, uint32(v31)))
	if v32 == v29 {
		v43 = v31
		v44 = v23
		goto L10
	} else {
		goto L14
	}
L12:
	;
	goto L13
L13:
	;
	v35 = F_lookup_type_cache(m, v29, int32(_a_F_multirange_union_0))
	mBase = m.M
	v36 = m.ExcPending
	if v36 != 0 {
		goto L1
	} else {
		goto L15
	}
L14:
	;
	goto L13
L15:
	;
	v37 = *(*int32)(unsafe.Add(mBase, uint32(v35)+296))
	if v37 == int32(0) {
		goto L4
	} else {
		goto L16
	}
L16:
	;
	v40 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	*(*int32)(unsafe.Add(mBase, uint32(v40)+16)) = v35
	v42 = *(*int32)(unsafe.Add(mBase, uint32(v16)+8))
	v43 = v35
	v44 = v42
	goto L10
L17:
	;
	v51 = F_palloc(m, v44<<(uint(int32(2))%32))
	mBase = m.M
	v52 = m.ExcPending
	if v52 != 0 {
		goto L1
	} else {
		goto L20
	}
L18:
	;
	v76 = v45
	v81 = v2
	goto L19
L19:
	;
	v83 = *(*int32)(unsafe.Add(mBase, uint32(v21)+8))
	v85 = v83 << (uint(int32(2)) % 32)
	if int32(0) < v83 {
		goto L25
	} else {
		goto L26
	}
L20:
	;
	v53 = int32(0)
	goto L21
L21:
	;
	v66 = F_multirange_get_range(m, v45, v16, v53)
	mBase = m.M
	v67 = m.ExcPending
	if v67 != 0 {
		goto L1
	} else {
		goto L23
	}
L22:
	;
	v72 = *(*int32)(unsafe.Add(mBase, uint32(v43)+296))
	v76 = v72
	v81 = v51
	goto L19
L23:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v51+v53<<(uint(int32(2))%32)))) = v66
	v70 = v53 + int32(1)
	if v70 != v44 {
		v53 = v70
		goto L21
	} else {
		goto L24
	}
L24:
	;
	goto L22
L25:
	;
	v89 = F_palloc(m, v85)
	mBase = m.M
	v90 = m.ExcPending
	if v90 != 0 {
		goto L1
	} else {
		goto L28
	}
L26:
	;
	v119 = v2
	goto L27
L27:
	;
	v120 = v83 + v44
	v123 = F_palloc0(m, v120<<(uint(int32(2))%32))
	mBase = m.M
	v124 = m.ExcPending
	if v124 != 0 {
		goto L1
	} else {
		goto L33
	}
L28:
	;
	v91 = int32(0)
	goto L29
L29:
	;
	v104 = F_multirange_get_range(m, v76, v21, v91)
	mBase = m.M
	v105 = m.ExcPending
	if v105 != 0 {
		goto L1
	} else {
		goto L31
	}
L30:
	;
	v119 = v89
	goto L27
L31:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v89+v91<<(uint(int32(2))%32)))) = v104
	v108 = v91 + int32(1)
	if v108 != v83 {
		v91 = v108
		goto L29
	} else {
		goto L32
	}
L32:
	;
	goto L30
L33:
	;
	v126 = v44 << (uint(int32(2)) % 32)
	if v126 != 0 {
		goto L34
	} else {
		goto L35
	}
L34:
	;
	base.MemoryCopy(m, v123, v81, v126)
	goto L36
L35:
	;
	goto L36
L36:
	;
	if v85 != 0 {
		goto L37
	} else {
		goto L38
	}
L37:
	;
	base.MemoryCopy(m, v123+v126, v119, v85)
	goto L39
L38:
	;
	goto L39
L39:
	;
	v130 = *(*int32)(unsafe.Add(mBase, uint32(v43)))
	v131 = *(*int32)(unsafe.Add(mBase, uint32(v43)+296))
	v132 = F_make_multirange(m, v130, v131, v120, v123)
	mBase = m.M
	v133 = m.ExcPending
	if v133 != 0 {
		goto L1
	} else {
		goto L40
	}
L40:
	;
	v136 = v132
	goto L5
L41:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v13))) = v29
	F_errmsg_internal(m, int32(_a_F_multirange_union_1), v13)
	mBase = m.M
	v155 = m.ExcPending
	if v155 != 0 {
		goto L1
	} else {
		goto L42
	}
L42:
	;
	F_errfinish(m, int32(_a_F_multirange_union_2), int32(558), int32(_a_F_multirange_union_3))
	mBase = m.M
	v160 = m.ExcPending
	if v160 != 0 {
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
func F_multirange_unnest(m *base.Module, l0 int32) int32 {
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
	var v64 int32
	_ = v64
	var v65 int32
	_ = v65
	var v68 int32
	_ = v68
	v6 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v7 = *(*int32)(unsafe.Add(mBase, uint32(v6)+16))
	if v7 == int32(0) {
		v10 = F_init_MultiFuncCall(m, l0)
		mBase = m.M
		v13 = m.ExcPending
		if v13 != 0 {
			return int32(0)
		} else {
			v14 = int32(_a_F_multirange_unnest_0)
			v15 = *(*int32)(unsafe.Add(mBase, _c_F_multirange_unnest[0]))
			v17 = *(*int32)(unsafe.Add(mBase, uint32(v10)+24))
			*(*int32)(unsafe.Add(mBase, _c_F_multirange_unnest[0])) = v17
			v19 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
			v20 = F_pg_detoast_datum(m, v19)
			mBase = m.M
			v21 = m.ExcPending
			if v21 != 0 {
				return int32(0)
			} else {
				v23 = F_palloc(m, int32(12))
				mBase = m.M
				v24 = m.ExcPending
				if v24 != 0 {
					return int32(0)
				} else {
					*(*int32)(unsafe.Add(mBase, uint32(v23)+8)) = int32(0)
					*(*int32)(unsafe.Add(mBase, uint32(v23))) = v20
					v28 = *(*int32)(unsafe.Add(mBase, uint32(v20)+4))
					v30 = F_lookup_type_cache(m, v28, int32(_a_F_multirange_unnest_1))
					mBase = m.M
					v31 = m.ExcPending
					if v31 != 0 {
						return int32(0)
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
								return int32(0)
							} else {
								v51 = *(*int32)(unsafe.Add(mBase, uint32(v42)+8))
								v52 = int32(1)
								*(*int32)(unsafe.Add(mBase, uint32(v42)+8)) = v51 + v52
								v55 = *(*int64)(unsafe.Add(mBase, uint32(v41)))
								*(*int64)(unsafe.Add(mBase, uint32(v41))) = v55 + int64(1)
								v59 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
								*(*int32)(unsafe.Add(mBase, uint32(v59)+20)) = v52
								return v49
							}
						} else {
							F_end_MultiFuncCall(m, l0)
							mBase = m.M
							v64 = m.ExcPending
							if v64 != 0 {
								return int32(0)
							} else {
								v65 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
								*(*int32)(unsafe.Add(mBase, uint32(v65)+20)) = int32(2)
								v68 = int32(1)
								*(*uint8)(unsafe.Add(mBase, uint32(l0)+16)) = uint8(v68)
								return int32(0)
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
				return int32(0)
			} else {
				v51 = *(*int32)(unsafe.Add(mBase, uint32(v42)+8))
				v52 = int32(1)
				*(*int32)(unsafe.Add(mBase, uint32(v42)+8)) = v51 + v52
				v55 = *(*int64)(unsafe.Add(mBase, uint32(v41)))
				*(*int64)(unsafe.Add(mBase, uint32(v41))) = v55 + int64(1)
				v59 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
				*(*int32)(unsafe.Add(mBase, uint32(v59)+20)) = v52
				return v49
			}
		} else {
			F_end_MultiFuncCall(m, l0)
			mBase = m.M
			v64 = m.ExcPending
			if v64 != 0 {
				return int32(0)
			} else {
				v65 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
				*(*int32)(unsafe.Add(mBase, uint32(v65)+20)) = int32(2)
				v68 = int32(1)
				*(*uint8)(unsafe.Add(mBase, uint32(l0)+16)) = uint8(v68)
				return int32(0)
			}
		}
	}
}
func F_multirange_upper_inc(m *base.Module, l0 int32) int32 {
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
	var v23 int32
	_ = v23
	var v24 int32
	_ = v24
	var v25 int32
	_ = v25
	var v28 int32
	_ = v28
	var v30 int32
	_ = v30
	var v31 int32
	_ = v31
	var v32 int32
	_ = v32
	var v33 int32
	_ = v33
	var v41 int32
	_ = v41
	var v42 int32
	_ = v42
	var v47 int32
	_ = v47
	var v55 int32
	_ = v55
	var v59 int32
	_ = v59
	var v64 int32
	_ = v64
	v7 = m.G0
	v9 = v7 - int32(32)
	m.G0 = v9
	v11 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v12 = F_pg_detoast_datum(m, v11)
	mBase = m.M
	v15 = m.ExcPending
	if v15 != 0 {
		return int32(0)
	} else {
		v16 = *(*int32)(unsafe.Add(mBase, uint32(v12)+8))
		if v16 != 0 {
			v17 = *(*int32)(unsafe.Add(mBase, uint32(v12)+4))
			v18 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
			v19 = *(*int32)(unsafe.Add(mBase, uint32(v18)+16))
			if v19 != 0 {
				v20 = *(*int32)(unsafe.Add(mBase, uint32(v19)))
				if v20 == v17 {
					v31 = v19
					v32 = v16
					v33 = *(*int32)(unsafe.Add(mBase, uint32(v31)+296))
					F_multirange_get_bounds(m, v33, v12, v32-int32(1), v9+int32(24), v9+int32(16))
					mBase = m.M
					v41 = m.ExcPending
					if v41 != 0 {
						return int32(0)
					} else {
						v42 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v9)+21)))
						v47 = v42
						m.G0 = v9 + int32(32)
						return v47
					}
				} else {
					v23 = F_lookup_type_cache(m, v17, int32(_a_F_multirange_upper_inc_0))
					mBase = m.M
					v24 = m.ExcPending
					if v24 != 0 {
						return int32(0)
					} else {
						v25 = *(*int32)(unsafe.Add(mBase, uint32(v23)+296))
						if v25 == int32(0) {
							F_errstart_cold(m, int32(21), int32(0))
							mBase = m.M
							v55 = m.ExcPending
							if v55 != 0 {
								return int32(0)
							} else {
								*(*int32)(unsafe.Add(mBase, uint32(v9))) = v17
								F_errmsg_internal(m, int32(_a_F_multirange_upper_inc_1), v9)
								mBase = m.M
								v59 = m.ExcPending
								if v59 != 0 {
									return int32(0)
								} else {
									F_errfinish(m, int32(_a_F_multirange_upper_inc_2), int32(558), int32(_a_F_multirange_upper_inc_3))
									mBase = m.M
									v64 = m.ExcPending
									if v64 != 0 {
										return int32(0)
									} else {
										base.Wasm_trap_unreachable()
										for {
										}
									}
								}
							}
						} else {
							v28 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
							*(*int32)(unsafe.Add(mBase, uint32(v28)+16)) = v23
							v30 = *(*int32)(unsafe.Add(mBase, uint32(v12)+8))
							v31 = v23
							v32 = v30
							v33 = *(*int32)(unsafe.Add(mBase, uint32(v31)+296))
							F_multirange_get_bounds(m, v33, v12, v32-int32(1), v9+int32(24), v9+int32(16))
							mBase = m.M
							v41 = m.ExcPending
							if v41 != 0 {
								return int32(0)
							} else {
								v42 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v9)+21)))
								v47 = v42
								m.G0 = v9 + int32(32)
								return v47
							}
						}
					}
				}
			} else {
				v23 = F_lookup_type_cache(m, v17, int32(_a_F_multirange_upper_inc_0))
				mBase = m.M
				v24 = m.ExcPending
				if v24 != 0 {
					return int32(0)
				} else {
					v25 = *(*int32)(unsafe.Add(mBase, uint32(v23)+296))
					if v25 == int32(0) {
						F_errstart_cold(m, int32(21), int32(0))
						mBase = m.M
						v55 = m.ExcPending
						if v55 != 0 {
							return int32(0)
						} else {
							*(*int32)(unsafe.Add(mBase, uint32(v9))) = v17
							F_errmsg_internal(m, int32(_a_F_multirange_upper_inc_1), v9)
							mBase = m.M
							v59 = m.ExcPending
							if v59 != 0 {
								return int32(0)
							} else {
								F_errfinish(m, int32(_a_F_multirange_upper_inc_2), int32(558), int32(_a_F_multirange_upper_inc_3))
								mBase = m.M
								v64 = m.ExcPending
								if v64 != 0 {
									return int32(0)
								} else {
									base.Wasm_trap_unreachable()
									for {
									}
								}
							}
						}
					} else {
						v28 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
						*(*int32)(unsafe.Add(mBase, uint32(v28)+16)) = v23
						v30 = *(*int32)(unsafe.Add(mBase, uint32(v12)+8))
						v31 = v23
						v32 = v30
						v33 = *(*int32)(unsafe.Add(mBase, uint32(v31)+296))
						F_multirange_get_bounds(m, v33, v12, v32-int32(1), v9+int32(24), v9+int32(16))
						mBase = m.M
						v41 = m.ExcPending
						if v41 != 0 {
							return int32(0)
						} else {
							v42 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v9)+21)))
							v47 = v42
							m.G0 = v9 + int32(32)
							return v47
						}
					}
				}
			}
		} else {
			v47 = int32(0)
			m.G0 = v9 + int32(32)
			return v47
		}
	}
}
