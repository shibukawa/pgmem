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
	var v48 int32
	_ = v48
	var v55 int32
	_ = v55
	var v60 int32
	_ = v60
	var v61 int32
	_ = v61
	var v64 int32
	_ = v64
	var v72 int32
	_ = v72
	var v76 int32
	_ = v76
	var v81 int32
	_ = v81
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
						v64 = v33
						m.G0 = v9 + int32(48)
						return v64
					} else {
						v37 = *(*int32)(unsafe.Add(mBase, uint32(v12)+8))
						if v37 == int32(0) {
							v64 = v33
							m.G0 = v9 + int32(48)
							return v64
						} else {
							v40 = *(*int32)(unsafe.Add(mBase, uint32(v32)+296))
							F_multirange_get_bounds(m, v40, v17, v34-int32(1), v9+int32(40), v9+int32(32))
							mBase = m.M
							v48 = m.ExcPending
							if v48 != 0 {
								return int32(0)
							} else {
								F_multirange_get_bounds(m, v40, v12, int32(0), v9+int32(24), v9+int32(16))
								mBase = m.M
								v55 = m.ExcPending
								if v55 != 0 {
									return int32(0)
								} else {
									v60 = F_range_cmp_bounds(m, v40, v9+int32(32), v9+int32(24))
									mBase = m.M
									v61 = m.ExcPending
									if v61 != 0 {
										return int32(0)
									} else {
										v64 = int32(base.Ui32(v60) >> (uint(int32(31)) % 32))
										m.G0 = v9 + int32(48)
										return v64
									}
								}
							}
						}
					}
				} else {
					v25 = F_lookup_type_cache(m, v19, int32(65536))
					mBase = m.M
					v26 = m.ExcPending
					if v26 != 0 {
						return int32(0)
					} else {
						v27 = *(*int32)(unsafe.Add(mBase, uint32(v25)+296))
						if v27 == int32(0) {
							F_errstart_cold(m, int32(21), int32(0))
							mBase = m.M
							v72 = m.ExcPending
							if v72 != 0 {
								return int32(0)
							} else {
								*(*int32)(unsafe.Add(mBase, uint32(v9))) = v19
								F_errmsg_internal(m, int32(370140), v9)
								mBase = m.M
								v76 = m.ExcPending
								if v76 != 0 {
									return int32(0)
								} else {
									F_errfinish(m, int32(493986), int32(558), int32(398748))
									mBase = m.M
									v81 = m.ExcPending
									if v81 != 0 {
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
								v64 = v33
								m.G0 = v9 + int32(48)
								return v64
							} else {
								v37 = *(*int32)(unsafe.Add(mBase, uint32(v12)+8))
								if v37 == int32(0) {
									v64 = v33
									m.G0 = v9 + int32(48)
									return v64
								} else {
									v40 = *(*int32)(unsafe.Add(mBase, uint32(v32)+296))
									F_multirange_get_bounds(m, v40, v17, v34-int32(1), v9+int32(40), v9+int32(32))
									mBase = m.M
									v48 = m.ExcPending
									if v48 != 0 {
										return int32(0)
									} else {
										F_multirange_get_bounds(m, v40, v12, int32(0), v9+int32(24), v9+int32(16))
										mBase = m.M
										v55 = m.ExcPending
										if v55 != 0 {
											return int32(0)
										} else {
											v60 = F_range_cmp_bounds(m, v40, v9+int32(32), v9+int32(24))
											mBase = m.M
											v61 = m.ExcPending
											if v61 != 0 {
												return int32(0)
											} else {
												v64 = int32(base.Ui32(v60) >> (uint(int32(31)) % 32))
												m.G0 = v9 + int32(48)
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
				v25 = F_lookup_type_cache(m, v19, int32(65536))
				mBase = m.M
				v26 = m.ExcPending
				if v26 != 0 {
					return int32(0)
				} else {
					v27 = *(*int32)(unsafe.Add(mBase, uint32(v25)+296))
					if v27 == int32(0) {
						F_errstart_cold(m, int32(21), int32(0))
						mBase = m.M
						v72 = m.ExcPending
						if v72 != 0 {
							return int32(0)
						} else {
							*(*int32)(unsafe.Add(mBase, uint32(v9))) = v19
							F_errmsg_internal(m, int32(370140), v9)
							mBase = m.M
							v76 = m.ExcPending
							if v76 != 0 {
								return int32(0)
							} else {
								F_errfinish(m, int32(493986), int32(558), int32(398748))
								mBase = m.M
								v81 = m.ExcPending
								if v81 != 0 {
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
							v64 = v33
							m.G0 = v9 + int32(48)
							return v64
						} else {
							v37 = *(*int32)(unsafe.Add(mBase, uint32(v12)+8))
							if v37 == int32(0) {
								v64 = v33
								m.G0 = v9 + int32(48)
								return v64
							} else {
								v40 = *(*int32)(unsafe.Add(mBase, uint32(v32)+296))
								F_multirange_get_bounds(m, v40, v17, v34-int32(1), v9+int32(40), v9+int32(32))
								mBase = m.M
								v48 = m.ExcPending
								if v48 != 0 {
									return int32(0)
								} else {
									F_multirange_get_bounds(m, v40, v12, int32(0), v9+int32(24), v9+int32(16))
									mBase = m.M
									v55 = m.ExcPending
									if v55 != 0 {
										return int32(0)
									} else {
										v60 = F_range_cmp_bounds(m, v40, v9+int32(32), v9+int32(24))
										mBase = m.M
										v61 = m.ExcPending
										if v61 != 0 {
											return int32(0)
										} else {
											v64 = int32(base.Ui32(v60) >> (uint(int32(31)) % 32))
											m.G0 = v9 + int32(48)
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
	var v56 int32
	_ = v56
	var v57 int32
	_ = v57
	var v81 int32
	_ = v81
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
	return v81
L2:
	;
	v81 = int32(1)
	goto L1
L3:
	;
	v16 = l0 + int32(212)
	v21 = int32(0)
	v22 = v14
	goto L6
L4:
	;
	goto L5
L5:
	;
	v81 = int32(0)
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
	if base.Ui32(v56) < base.Ui32(v57) {
		v21 = v56
		v22 = v57
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
	v56 = v21
	v57 = v28
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
	if v42 == int32(1) {
		goto L11
	} else {
		goto L18
	}
L18:
	;
	v56 = v21
	v57 = v28
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
	v56 = v28 + int32(1)
	v57 = v22
	goto L10
L24:
	;
	v53 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v12)+5)))
	if v53 != 0 {
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
	var v97 int32
	_ = v97
	var v98 int32
	_ = v98
	var v103 int32
	_ = v103
	var v105 int32
	_ = v105
	var v106 int32
	_ = v106
	var v109 int32
	_ = v109
	var v116 int32
	_ = v116
	var v117 int32
	_ = v117
	var v126 int32
	_ = v126
	var v127 int32
	_ = v127
	var v128 int32
	_ = v128
	var v133 int32
	_ = v133
	var v139 int32
	_ = v139
	var v143 int32
	_ = v143
	var v148 int32
	_ = v148
	var v149 int32
	_ = v149
	var v165 int32
	_ = v165
	var v166 int32
	_ = v166
	var v167 int32
	_ = v167
	var v170 int32
	_ = v170
	var v171 int32
	_ = v171
	var v181 int32
	_ = v181
	var v190 int32
	_ = v190
	var v193 int32
	_ = v193
	var v195 int32
	_ = v195
	var v212 int32
	_ = v212
	var v215 int32
	_ = v215
	var v220 int32
	_ = v220
	var v221 int32
	_ = v221
	var v222 int32
	_ = v222
	var v227 int32
	_ = v227
	var v233 int32
	_ = v233
	var v237 int32
	_ = v237
	var v242 int32
	_ = v242
	var v252 int32
	_ = v252
	var v253 int32
	_ = v253
	var v259 int32
	_ = v259
	var v260 int32
	_ = v260
	var v261 int32
	_ = v261
	var v262 int32
	_ = v262
	var v263 int32
	_ = v263
	var v266 int32
	_ = v266
	var v267 int32
	_ = v267
	var v270 int32
	_ = v270
	var v274 int32
	_ = v274
	var v275 int32
	_ = v275
	var v276 int32
	_ = v276
	var v277 int32
	_ = v277
	var v278 int32
	_ = v278
	var v279 int32
	_ = v279
	var v285 int32
	_ = v285
	var v300 int32
	_ = v300
	var v301 int32
	_ = v301
	var v303 int32
	_ = v303
	var v304 int32
	_ = v304
	var v307 int32
	_ = v307
	var v314 int32
	_ = v314
	var v315 int32
	_ = v315
	var v316 int32
	_ = v316
	var v321 int32
	_ = v321
	var v327 int32
	_ = v327
	var v331 int32
	_ = v331
	var v336 int32
	_ = v336
	var v339 int32
	_ = v339
	var v341 int32
	_ = v341
	var v343 int32
	_ = v343
	var v344 int32
	_ = v344
	var v345 int32
	_ = v345
	var v346 int32
	_ = v346
	var v350 int32
	_ = v350
	var v370 int32
	_ = v370
	var v371 int32
	_ = v371
	var v378 int32
	_ = v378
	var v379 int32
	_ = v379
	var v380 int32
	_ = v380
	var v385 int32
	_ = v385
	var v391 int32
	_ = v391
	var v395 int32
	_ = v395
	var v400 int32
	_ = v400
	var v401 int32
	_ = v401
	var v402 int32
	_ = v402
	var v404 int32
	_ = v404
	v2 = int32(0)
	v21 = m.G0
	v23 = v21 - int32(80)
	m.G0 = v23
	v25 = *(*int32)(unsafe.Add(mBase, uint32(l0)+36))
	v26 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
	v27 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v28 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
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
	v35 = F_get_multirange_io_data(m, l0, v26, int32(0))
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
	v40 = v28
	goto L6
L4:
	;
	m.G0 = v23 + int32(80)
	return v404
L5:
	;
	v96 = v40
	v97 = int32(0)
	v98 = v2
	v103 = v2
	v105 = v30
	v106 = v2
	v109 = int32(8)
	goto L19
L6:
	;
	v59 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v40))))
	if base.Ui32(v59-int32(9)) < base.Ui32(int32(5)) {
		goto L9
	} else {
		goto L10
	}
L7:
	;
	v74 = int32(0)
	v75 = F_errsave_start(m, v27)
	mBase = m.M
	v76 = m.ExcPending
	if v76 != 0 {
		goto L1
	} else {
		goto L13
	}
L8:
	;
	goto L7
L9:
	;
	v40 = v40 + int32(1)
	goto L6
L10:
	;
	if v59 == int32(32) {
		goto L9
	} else {
		goto L11
	}
L11:
	;
	if v59 != int32(123) {
		goto L8
	} else {
		goto L12
	}
L12:
	;
	goto L5
L13:
	;
	if v75 == int32(0) {
		v404 = v74
		goto L4
	} else {
		goto L14
	}
L14:
	;
	F_errcode(m, int32(33685634))
	mBase = m.M
	v81 = m.ExcPending
	if v81 != 0 {
		goto L1
	} else {
		goto L15
	}
L15:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v23))) = v28
	F_errmsg(m, int32(726157), v23)
	mBase = m.M
	v85 = m.ExcPending
	if v85 != 0 {
		goto L1
	} else {
		goto L16
	}
L16:
	;
	F_errdetail(m, int32(643957), int32(0))
	mBase = m.M
	v89 = m.ExcPending
	if v89 != 0 {
		goto L1
	} else {
		goto L17
	}
L17:
	;
	F_errsave_finish(m, v27, int32(493986), int32(152), int32(279701))
	mBase = m.M
	v94 = m.ExcPending
	if v94 != 0 {
		goto L1
	} else {
		goto L18
	}
L18:
	;
	v404 = v74
	goto L4
L19:
	;
	v116 = v96 + int32(1)
	v117 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v116))))
	if base.Ui32(v117-int32(9)) < base.Ui32(int32(5)) {
		v96 = v116
		goto L19
	} else {
		goto L21
	}
L20:
	;
	v350 = v116
	goto L101
L21:
	;
	if v117 == int32(32) {
		v96 = v116
		goto L19
	} else {
		goto L22
	}
L22:
	;
	if v117 == int32(0) {
		goto L23
	} else {
		goto L24
	}
L23:
	;
	v126 = int32(0)
	v127 = F_errsave_start(m, v27)
	mBase = m.M
	v128 = m.ExcPending
	if v128 != 0 {
		goto L1
	} else {
		goto L26
	}
L24:
	;
	goto L25
L25:
	;
	v149 = int32(1)
	switch v98 - v149 {
	case 0:
		goto L38
	case 1:
		v96 = v116
		v98 = v149
		goto L19
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
	if v127 == int32(0) {
		v404 = v126
		goto L4
	} else {
		goto L27
	}
L27:
	;
	F_errcode(m, int32(33685634))
	mBase = m.M
	v133 = m.ExcPending
	if v133 != 0 {
		goto L1
	} else {
		goto L28
	}
L28:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v23)+64)) = v28
	F_errmsg(m, int32(726157), v23-int32(-64))
	mBase = m.M
	v139 = m.ExcPending
	if v139 != 0 {
		goto L1
	} else {
		goto L29
	}
L29:
	;
	F_errdetail(m, int32(578011), int32(0))
	mBase = m.M
	v143 = m.ExcPending
	if v143 != 0 {
		goto L1
	} else {
		goto L30
	}
L30:
	;
	F_errsave_finish(m, v27, int32(493986), int32(165), int32(279701))
	mBase = m.M
	v148 = m.ExcPending
	if v148 != 0 {
		goto L1
	} else {
		goto L31
	}
L31:
	;
	v404 = v126
	goto L4
L32:
	;
	goto L20
L33:
	;
	v96 = v339
	v98 = v341
	v103 = v343
	v105 = v344
	v106 = v345
	v109 = v346
	goto L19
L34:
	;
	v339 = v116
	v341 = int32(1)
	v343 = v103
	v344 = v105
	v345 = v106
	v346 = v109
	goto L33
L35:
	;
	v339 = v116
	v341 = int32(3)
	v343 = v103
	v344 = v105
	v345 = v106
	v346 = v109
	goto L33
L36:
	;
	if v117 == int32(44) {
		v96 = v116
		v98 = int32(0)
		goto L19
	} else {
		goto L93
	}
L37:
	;
	if v117 != int32(92) {
		goto L83
	} else {
		goto L84
	}
L38:
	;
	switch v117 - int32(34) {
	case 0:
		v96 = v116
		v98 = int32(3)
		goto L19
	case 1, 2, 3, 4, 5, 6:
		goto L34
	case 7:
		goto L68
	default:
		goto L69
	}
L39:
	;
	if v117 == int32(40) {
		v96 = v116
		v97 = v116
		v98 = v149
		goto L19
	} else {
		goto L40
	}
L40:
	;
	if v117 == int32(91) {
		v96 = v116
		v97 = v116
		v98 = v149
		goto L19
	} else {
		goto L41
	}
L41:
	;
	if base.B2i32(v106 == int32(0))&base.B2i32(v117 == int32(125)) != 0 {
		goto L32
	} else {
		goto L42
	}
L42:
	;
	v165 = v116
	v166 = int32(8965)
	v167 = int32(5)
	goto L44
L43:
	;
	if v212 == int32(0) {
		goto L59
	} else {
		goto L60
	}
L44:
	;
	if v167 != 0 {
		goto L46
	} else {
		goto L47
	}
L45:
	;
	v212 = int32(0)
	goto L43
L46:
	;
	v170 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v165))))
	v171 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v166))))
	if v170 == v171 {
		v193 = v170
		goto L49
	} else {
		goto L50
	}
L47:
	;
	goto L48
L48:
	;
	goto L45
L49:
	;
	v195 = int32(1)
	if v193 != 0 {
		v165 = v165 + v195
		v166 = v166 + v195
		v167 = v167 - v195
		goto L44
	} else {
		goto L58
	}
L50:
	;
	if base.Ui32((v170-int32(65))&int32(255)) < base.Ui32(int32(26)) {
		goto L51
	} else {
		goto L52
	}
L51:
	;
	v181 = v170 | int32(32)
	goto L53
L52:
	;
	v181 = v170
	goto L53
L53:
	;
	if base.Ui32((v171-int32(65))&int32(255)) < base.Ui32(int32(26)) {
		goto L54
	} else {
		goto L55
	}
L54:
	;
	v190 = v171 | int32(32)
	goto L56
L55:
	;
	v190 = v171
	goto L56
L56:
	;
	if v181 == v190 {
		v193 = v181
		goto L49
	} else {
		goto L57
	}
L57:
	;
	v212 = v181 - v190
	goto L43
L58:
	;
	goto L48
L59:
	;
	v215 = int32(5)
	v339 = v96 + v215
	v341 = v215
	v343 = v103
	v344 = v105
	v345 = v106 + int32(1)
	v346 = v109
	goto L33
L60:
	;
	goto L61
L61:
	;
	v220 = int32(0)
	v221 = F_errsave_start(m, v27)
	mBase = m.M
	v222 = m.ExcPending
	if v222 != 0 {
		goto L1
	} else {
		goto L62
	}
L62:
	;
	if v221 == int32(0) {
		v404 = v220
		goto L4
	} else {
		goto L63
	}
L63:
	;
	F_errcode(m, int32(33685634))
	mBase = m.M
	v227 = m.ExcPending
	if v227 != 0 {
		goto L1
	} else {
		goto L64
	}
L64:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v23)+32)) = v28
	F_errmsg(m, int32(726157), v23+int32(32))
	mBase = m.M
	v233 = m.ExcPending
	if v233 != 0 {
		goto L1
	} else {
		goto L65
	}
L65:
	;
	F_errdetail(m, int32(578861), int32(0))
	mBase = m.M
	v237 = m.ExcPending
	if v237 != 0 {
		goto L1
	} else {
		goto L66
	}
L66:
	;
	F_errsave_finish(m, v27, int32(493986), int32(194), int32(279701))
	mBase = m.M
	v242 = m.ExcPending
	if v242 != 0 {
		goto L1
	} else {
		goto L67
	}
L67:
	;
	v404 = v220
	goto L4
L68:
	;
	v252 = F_pnstrdup(m, v97, v116-v97+int32(1))
	mBase = m.M
	v253 = m.ExcPending
	if v253 != 0 {
		goto L1
	} else {
		goto L71
	}
L69:
	;
	switch v117 - int32(92) {
	case 0:
		goto L70
	case 1:
		goto L68
	default:
		goto L34
	}
L70:
	;
	v339 = v116
	v341 = int32(2)
	v343 = v103
	v344 = v105
	v345 = v106
	v346 = v109
	goto L33
L71:
	;
	if v103 == v109 {
		goto L72
	} else {
		goto L73
	}
L72:
	;
	v259 = F_repalloc(m, v105, v103<<(uint(int32(3))%32))
	mBase = m.M
	v260 = m.ExcPending
	if v260 != 0 {
		goto L1
	} else {
		goto L75
	}
L73:
	;
	v261 = v105
	v262 = v109
	goto L74
L74:
	;
	v263 = *(*int32)(unsafe.Add(mBase, uint32(v35)+32))
	v266 = F_InputFunctionCallSafe(m, v35+int32(4), v252, v263, v25, v27, v23+int32(76))
	mBase = m.M
	v267 = m.ExcPending
	if v267 != 0 {
		goto L1
	} else {
		goto L76
	}
L75:
	;
	v261 = v259
	v262 = v103 << (uint(int32(1)) % 32)
	goto L74
L76:
	;
	if v266 == int32(0) {
		goto L77
	} else {
		goto L78
	}
L77:
	;
	v270 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(l0)+16)) = uint8(v270)
	v404 = int32(0)
	goto L4
L78:
	;
	goto L79
L79:
	;
	v274 = v106 + int32(1)
	v275 = int32(5)
	v276 = *(*int32)(unsafe.Add(mBase, uint32(v23)+76))
	v277 = F_pg_detoast_datum(m, v276)
	mBase = m.M
	v278 = m.ExcPending
	if v278 != 0 {
		goto L1
	} else {
		goto L80
	}
L80:
	;
	v279 = *(*int32)(unsafe.Add(mBase, uint32(v277)))
	v285 = int32(*(*int8)(unsafe.Add(mBase, uint32(v277+int32(base.Ui32(v279)>>(uint(int32(2))%32))-int32(1)))))
	goto L81
L81:
	;
	if v285&int32(1) != 0 {
		v96 = v116
		v98 = v275
		v105 = v261
		v106 = v274
		v109 = v262
		goto L19
	} else {
		goto L82
	}
L82:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v261+v103<<(uint(int32(2))%32)))) = v277
	v339 = v116
	v341 = v275
	v343 = v103 + int32(1)
	v344 = v261
	v345 = v274
	v346 = v262
	goto L33
L83:
	;
	if v117 != int32(34) {
		v96 = v116
		v98 = int32(3)
		goto L19
	} else {
		goto L86
	}
L84:
	;
	goto L85
L85:
	;
	v339 = v116
	v341 = int32(4)
	v343 = v103
	v344 = v105
	v345 = v106
	v346 = v109
	goto L33
L86:
	;
	v300 = v96 + int32(2)
	v301 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v300))))
	v303 = base.B2i32(v301 == int32(34))
	if v301 == int32(34) {
		goto L87
	} else {
		goto L88
	}
L87:
	;
	v304 = v300
	goto L89
L88:
	;
	v304 = v116
	goto L89
L89:
	;
	if v301 == int32(34) {
		goto L90
	} else {
		goto L91
	}
L90:
	;
	v307 = int32(3)
	goto L92
L91:
	;
	v307 = int32(1)
	goto L92
L92:
	;
	v339 = v304
	v341 = v307
	v343 = v103
	v344 = v105
	v345 = v106
	v346 = v109
	goto L33
L93:
	;
	if v117 == int32(125) {
		goto L32
	} else {
		goto L94
	}
L94:
	;
	v314 = int32(0)
	v315 = F_errsave_start(m, v27)
	mBase = m.M
	v316 = m.ExcPending
	if v316 != 0 {
		goto L1
	} else {
		goto L95
	}
L95:
	;
	if v315 == int32(0) {
		v404 = v314
		goto L4
	} else {
		goto L96
	}
L96:
	;
	F_errcode(m, int32(33685634))
	mBase = m.M
	v321 = m.ExcPending
	if v321 != 0 {
		goto L1
	} else {
		goto L97
	}
L97:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v23)+48)) = v28
	F_errmsg(m, int32(726157), v23+int32(48))
	mBase = m.M
	v327 = m.ExcPending
	if v327 != 0 {
		goto L1
	} else {
		goto L98
	}
L98:
	;
	F_errdetail(m, int32(641519), int32(0))
	mBase = m.M
	v331 = m.ExcPending
	if v331 != 0 {
		goto L1
	} else {
		goto L99
	}
L99:
	;
	F_errsave_finish(m, v27, int32(493986), int32(268), int32(279701))
	mBase = m.M
	v336 = m.ExcPending
	if v336 != 0 {
		goto L1
	} else {
		goto L100
	}
L100:
	;
	v404 = v314
	goto L4
L101:
	;
	v370 = v350 + int32(1)
	v371 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v370))))
	if base.Ui32(v371-int32(9)) < base.Ui32(int32(5)) {
		v350 = v370
		goto L101
	} else {
		goto L103
	}
L102:
	;
	if v371 != 0 {
		goto L105
	} else {
		goto L106
	}
L103:
	;
	if v371 == int32(32) {
		v350 = v370
		goto L101
	} else {
		goto L104
	}
L104:
	;
	goto L102
L105:
	;
	v378 = int32(0)
	v379 = F_errsave_start(m, v27)
	mBase = m.M
	v380 = m.ExcPending
	if v380 != 0 {
		goto L1
	} else {
		goto L108
	}
L106:
	;
	goto L107
L107:
	;
	v401 = F_make_multirange(m, v26, v38, v103, v105)
	mBase = m.M
	v402 = m.ExcPending
	if v402 != 0 {
		goto L1
	} else {
		goto L114
	}
L108:
	;
	if v379 == int32(0) {
		v404 = v378
		goto L4
	} else {
		goto L109
	}
L109:
	;
	F_errcode(m, int32(33685634))
	mBase = m.M
	v385 = m.ExcPending
	if v385 != 0 {
		goto L1
	} else {
		goto L110
	}
L110:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v23)+16)) = v28
	F_errmsg(m, int32(726157), v23+int32(16))
	mBase = m.M
	v391 = m.ExcPending
	if v391 != 0 {
		goto L1
	} else {
		goto L111
	}
L111:
	;
	F_errdetail(m, int32(643925), int32(0))
	mBase = m.M
	v395 = m.ExcPending
	if v395 != 0 {
		goto L1
	} else {
		goto L112
	}
L112:
	;
	F_errsave_finish(m, v27, int32(493986), int32(292), int32(279701))
	mBase = m.M
	v400 = m.ExcPending
	if v400 != 0 {
		goto L1
	} else {
		goto L113
	}
L113:
	;
	v404 = v378
	goto L4
L114:
	;
	v404 = v401
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
					v25 = F_lookup_type_cache(m, v19, int32(65536))
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
								F_errmsg_internal(m, int32(370140), v9)
								mBase = m.M
								v49 = m.ExcPending
								if v49 != 0 {
									return int32(0)
								} else {
									F_errfinish(m, int32(493986), int32(558), int32(398748))
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
				v25 = F_lookup_type_cache(m, v19, int32(65536))
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
							F_errmsg_internal(m, int32(370140), v9)
							mBase = m.M
							v49 = m.ExcPending
							if v49 != 0 {
								return int32(0)
							} else {
								F_errfinish(m, int32(493986), int32(558), int32(398748))
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
	var v54 int32
	_ = v54
	var v55 int32
	_ = v55
	var v63 int32
	_ = v63
	var v64 int32
	_ = v64
	var v69 int32
	_ = v69
	var v70 int32
	_ = v70
	var v73 int32
	_ = v73
	var v82 int32
	_ = v82
	var v86 int32
	_ = v86
	var v91 int32
	_ = v91
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
				v73 = v2
				m.G0 = v9 + int32(48)
				return v73
			} else {
				v22 = *(*int32)(unsafe.Add(mBase, uint32(v17)))
				v28 = int32(*(*int8)(unsafe.Add(mBase, uint32(v17+int32(base.Ui32(v22)>>(uint(int32(2))%32))-int32(1)))))
				if v28&int32(1) != 0 {
					v73 = v2
					m.G0 = v9 + int32(48)
					return v73
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
							F_multirange_get_bounds(m, v45, v12, v46-int32(1), v9+int32(40), v9+int32(32))
							mBase = m.M
							v54 = m.ExcPending
							if v54 != 0 {
								return int32(0)
							} else {
								v55 = *(*int32)(unsafe.Add(mBase, uint32(v44)+296))
								F_range_deserialize(m, v55, v17, v9+int32(24), v9+int32(16), v9+int32(15))
								mBase = m.M
								v63 = m.ExcPending
								if v63 != 0 {
									return int32(0)
								} else {
									v64 = *(*int32)(unsafe.Add(mBase, uint32(v44)+296))
									v69 = F_range_cmp_bounds(m, v64, v9+int32(32), v9+int32(16))
									mBase = m.M
									v70 = m.ExcPending
									if v70 != 0 {
										return int32(0)
									} else {
										v73 = base.B2i32(v69 <= int32(0))
										m.G0 = v9 + int32(48)
										return v73
									}
								}
							}
						} else {
							v37 = F_lookup_type_cache(m, v31, int32(65536))
							mBase = m.M
							v38 = m.ExcPending
							if v38 != 0 {
								return int32(0)
							} else {
								v39 = *(*int32)(unsafe.Add(mBase, uint32(v37)+296))
								if v39 == int32(0) {
									F_errstart_cold(m, int32(21), int32(0))
									mBase = m.M
									v82 = m.ExcPending
									if v82 != 0 {
										return int32(0)
									} else {
										*(*int32)(unsafe.Add(mBase, uint32(v9))) = v31
										F_errmsg_internal(m, int32(370140), v9)
										mBase = m.M
										v86 = m.ExcPending
										if v86 != 0 {
											return int32(0)
										} else {
											F_errfinish(m, int32(493986), int32(558), int32(398748))
											mBase = m.M
											v91 = m.ExcPending
											if v91 != 0 {
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
									F_multirange_get_bounds(m, v45, v12, v46-int32(1), v9+int32(40), v9+int32(32))
									mBase = m.M
									v54 = m.ExcPending
									if v54 != 0 {
										return int32(0)
									} else {
										v55 = *(*int32)(unsafe.Add(mBase, uint32(v44)+296))
										F_range_deserialize(m, v55, v17, v9+int32(24), v9+int32(16), v9+int32(15))
										mBase = m.M
										v63 = m.ExcPending
										if v63 != 0 {
											return int32(0)
										} else {
											v64 = *(*int32)(unsafe.Add(mBase, uint32(v44)+296))
											v69 = F_range_cmp_bounds(m, v64, v9+int32(32), v9+int32(16))
											mBase = m.M
											v70 = m.ExcPending
											if v70 != 0 {
												return int32(0)
											} else {
												v73 = base.B2i32(v69 <= int32(0))
												m.G0 = v9 + int32(48)
												return v73
											}
										}
									}
								}
							}
						}
					} else {
						v37 = F_lookup_type_cache(m, v31, int32(65536))
						mBase = m.M
						v38 = m.ExcPending
						if v38 != 0 {
							return int32(0)
						} else {
							v39 = *(*int32)(unsafe.Add(mBase, uint32(v37)+296))
							if v39 == int32(0) {
								F_errstart_cold(m, int32(21), int32(0))
								mBase = m.M
								v82 = m.ExcPending
								if v82 != 0 {
									return int32(0)
								} else {
									*(*int32)(unsafe.Add(mBase, uint32(v9))) = v31
									F_errmsg_internal(m, int32(370140), v9)
									mBase = m.M
									v86 = m.ExcPending
									if v86 != 0 {
										return int32(0)
									} else {
										F_errfinish(m, int32(493986), int32(558), int32(398748))
										mBase = m.M
										v91 = m.ExcPending
										if v91 != 0 {
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
								F_multirange_get_bounds(m, v45, v12, v46-int32(1), v9+int32(40), v9+int32(32))
								mBase = m.M
								v54 = m.ExcPending
								if v54 != 0 {
									return int32(0)
								} else {
									v55 = *(*int32)(unsafe.Add(mBase, uint32(v44)+296))
									F_range_deserialize(m, v55, v17, v9+int32(24), v9+int32(16), v9+int32(15))
									mBase = m.M
									v63 = m.ExcPending
									if v63 != 0 {
										return int32(0)
									} else {
										v64 = *(*int32)(unsafe.Add(mBase, uint32(v44)+296))
										v69 = F_range_cmp_bounds(m, v64, v9+int32(32), v9+int32(16))
										mBase = m.M
										v70 = m.ExcPending
										if v70 != 0 {
											return int32(0)
										} else {
											v73 = base.B2i32(v69 <= int32(0))
											m.G0 = v9 + int32(48)
											return v73
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
	var v46 int32
	_ = v46
	var v47 int32
	_ = v47
	var v54 int32
	_ = v54
	var v55 int32
	_ = v55
	var v60 int32
	_ = v60
	var v61 int32
	_ = v61
	var v66 int32
	_ = v66
	var v75 int32
	_ = v75
	var v79 int32
	_ = v79
	var v84 int32
	_ = v84
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
				v66 = v2
				m.G0 = v9 + int32(48)
				return v66
			} else {
				v22 = *(*int32)(unsafe.Add(mBase, uint32(v17)+8))
				if v22 == int32(0) {
					v66 = v2
					m.G0 = v9 + int32(48)
					return v66
				} else {
					v25 = *(*int32)(unsafe.Add(mBase, uint32(v12)+4))
					v26 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
					v27 = *(*int32)(unsafe.Add(mBase, uint32(v26)+16))
					if v27 != 0 {
						v28 = *(*int32)(unsafe.Add(mBase, uint32(v27)))
						if v28 == v25 {
							v38 = v27
							v39 = *(*int32)(unsafe.Add(mBase, uint32(v38)+296))
							F_multirange_get_bounds(m, v39, v12, int32(0), v9+int32(40), v9+int32(32))
							mBase = m.M
							v46 = m.ExcPending
							if v46 != 0 {
								return int32(0)
							} else {
								v47 = *(*int32)(unsafe.Add(mBase, uint32(v38)+296))
								F_multirange_get_bounds(m, v47, v17, int32(0), v9+int32(24), v9+int32(16))
								mBase = m.M
								v54 = m.ExcPending
								if v54 != 0 {
									return int32(0)
								} else {
									v55 = *(*int32)(unsafe.Add(mBase, uint32(v38)+296))
									v60 = F_range_cmp_bounds(m, v55, v9+int32(40), v9+int32(24))
									mBase = m.M
									v61 = m.ExcPending
									if v61 != 0 {
										return int32(0)
									} else {
										v66 = int32(base.Ui32(v60^int32(-1)) >> (uint(int32(31)) % 32))
										m.G0 = v9 + int32(48)
										return v66
									}
								}
							}
						} else {
							v31 = F_lookup_type_cache(m, v25, int32(65536))
							mBase = m.M
							v32 = m.ExcPending
							if v32 != 0 {
								return int32(0)
							} else {
								v33 = *(*int32)(unsafe.Add(mBase, uint32(v31)+296))
								if v33 == int32(0) {
									F_errstart_cold(m, int32(21), int32(0))
									mBase = m.M
									v75 = m.ExcPending
									if v75 != 0 {
										return int32(0)
									} else {
										*(*int32)(unsafe.Add(mBase, uint32(v9))) = v25
										F_errmsg_internal(m, int32(370140), v9)
										mBase = m.M
										v79 = m.ExcPending
										if v79 != 0 {
											return int32(0)
										} else {
											F_errfinish(m, int32(493986), int32(558), int32(398748))
											mBase = m.M
											v84 = m.ExcPending
											if v84 != 0 {
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
									F_multirange_get_bounds(m, v39, v12, int32(0), v9+int32(40), v9+int32(32))
									mBase = m.M
									v46 = m.ExcPending
									if v46 != 0 {
										return int32(0)
									} else {
										v47 = *(*int32)(unsafe.Add(mBase, uint32(v38)+296))
										F_multirange_get_bounds(m, v47, v17, int32(0), v9+int32(24), v9+int32(16))
										mBase = m.M
										v54 = m.ExcPending
										if v54 != 0 {
											return int32(0)
										} else {
											v55 = *(*int32)(unsafe.Add(mBase, uint32(v38)+296))
											v60 = F_range_cmp_bounds(m, v55, v9+int32(40), v9+int32(24))
											mBase = m.M
											v61 = m.ExcPending
											if v61 != 0 {
												return int32(0)
											} else {
												v66 = int32(base.Ui32(v60^int32(-1)) >> (uint(int32(31)) % 32))
												m.G0 = v9 + int32(48)
												return v66
											}
										}
									}
								}
							}
						}
					} else {
						v31 = F_lookup_type_cache(m, v25, int32(65536))
						mBase = m.M
						v32 = m.ExcPending
						if v32 != 0 {
							return int32(0)
						} else {
							v33 = *(*int32)(unsafe.Add(mBase, uint32(v31)+296))
							if v33 == int32(0) {
								F_errstart_cold(m, int32(21), int32(0))
								mBase = m.M
								v75 = m.ExcPending
								if v75 != 0 {
									return int32(0)
								} else {
									*(*int32)(unsafe.Add(mBase, uint32(v9))) = v25
									F_errmsg_internal(m, int32(370140), v9)
									mBase = m.M
									v79 = m.ExcPending
									if v79 != 0 {
										return int32(0)
									} else {
										F_errfinish(m, int32(493986), int32(558), int32(398748))
										mBase = m.M
										v84 = m.ExcPending
										if v84 != 0 {
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
								F_multirange_get_bounds(m, v39, v12, int32(0), v9+int32(40), v9+int32(32))
								mBase = m.M
								v46 = m.ExcPending
								if v46 != 0 {
									return int32(0)
								} else {
									v47 = *(*int32)(unsafe.Add(mBase, uint32(v38)+296))
									F_multirange_get_bounds(m, v47, v17, int32(0), v9+int32(24), v9+int32(16))
									mBase = m.M
									v54 = m.ExcPending
									if v54 != 0 {
										return int32(0)
									} else {
										v55 = *(*int32)(unsafe.Add(mBase, uint32(v38)+296))
										v60 = F_range_cmp_bounds(m, v55, v9+int32(40), v9+int32(24))
										mBase = m.M
										v61 = m.ExcPending
										if v61 != 0 {
											return int32(0)
										} else {
											v66 = int32(base.Ui32(v60^int32(-1)) >> (uint(int32(31)) % 32))
											m.G0 = v9 + int32(48)
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
				v16 = *(*int32)(unsafe.Add(mBase, _consts[831]))
				*(*int32)(unsafe.Add(mBase, uint32(v4))) = v16
				v18 = v16
			} else {
				v18 = v12
			}
			*(*int32)(unsafe.Add(mBase, uint32(v4)+32)) = v10
			*(*int32)(unsafe.Add(mBase, uint32(v4)+24)) = int32(1490)
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
	var v75 int32
	_ = v75
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
	var v117 int32
	_ = v117
	var v120 int32
	_ = v120
	var v123 int32
	_ = v123
	var v124 int32
	_ = v124
	var v126 int32
	_ = v126
	var v127 int32
	_ = v127
	var v128 int32
	_ = v128
	var v130 int32
	_ = v130
	var v132 int32
	_ = v132
	var v133 int32
	_ = v133
	var v134 int32
	_ = v134
	var v135 int32
	_ = v135
	var v139 int32
	_ = v139
	var v153 int32
	_ = v153
	var v157 int32
	_ = v157
	var v162 int32
	_ = v162
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
		v139 = v21
		goto L5
	} else {
		goto L6
	}
L4:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v153 = m.ExcPending
	if v153 != 0 {
		goto L1
	} else {
		goto L44
	}
L5:
	;
	m.G0 = v13 + int32(16)
	return v139
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
	v139 = v16
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
	if v44 <= int32(0) {
		goto L18
	} else {
		goto L19
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
	v35 = F_lookup_type_cache(m, v29, int32(65536))
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
	v83 = *(*int32)(unsafe.Add(mBase, uint32(v21)+8))
	v85 = v83 << (uint(int32(2)) % 32)
	if int32(0) < v83 {
		goto L26
	} else {
		goto L27
	}
L18:
	;
	v75 = v45
	v81 = v2
	goto L17
L19:
	;
	goto L20
L20:
	;
	v51 = F_palloc(m, v44<<(uint(int32(2))%32))
	mBase = m.M
	v52 = m.ExcPending
	if v52 != 0 {
		goto L1
	} else {
		goto L21
	}
L21:
	;
	v53 = int32(0)
	goto L22
L22:
	;
	v66 = F_multirange_get_range(m, v45, v16, v53)
	mBase = m.M
	v67 = m.ExcPending
	if v67 != 0 {
		goto L1
	} else {
		goto L24
	}
L23:
	;
	v72 = *(*int32)(unsafe.Add(mBase, uint32(v43)+296))
	v75 = v72
	v81 = v51
	goto L17
L24:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v51+v53<<(uint(int32(2))%32)))) = v66
	v70 = v53 + int32(1)
	if v70 != v44 {
		v53 = v70
		goto L22
	} else {
		goto L25
	}
L25:
	;
	goto L23
L26:
	;
	v89 = F_palloc(m, v85)
	mBase = m.M
	v90 = m.ExcPending
	if v90 != 0 {
		goto L1
	} else {
		goto L29
	}
L27:
	;
	v117 = v2
	goto L28
L28:
	;
	v120 = v83 + v44
	v123 = F_palloc0(m, v120<<(uint(int32(2))%32))
	mBase = m.M
	v124 = m.ExcPending
	if v124 != 0 {
		goto L1
	} else {
		goto L34
	}
L29:
	;
	v91 = int32(0)
	goto L30
L30:
	;
	v104 = F_multirange_get_range(m, v75, v21, v91)
	mBase = m.M
	v105 = m.ExcPending
	if v105 != 0 {
		goto L1
	} else {
		goto L32
	}
L31:
	;
	v117 = v89
	goto L28
L32:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v89+v91<<(uint(int32(2))%32)))) = v104
	v108 = v91 + int32(1)
	if v108 != v83 {
		v91 = v108
		goto L30
	} else {
		goto L33
	}
L33:
	;
	goto L31
L34:
	;
	v126 = v44 << (uint(int32(2)) % 32)
	if v126 != 0 {
		goto L36
	} else {
		goto L37
	}
L35:
	;
	if v85 != 0 {
		goto L40
	} else {
		goto L41
	}
L36:
	;
	v127 = F__emscripten_memcpy_bulkmem(m, v123, v81, v126)
	mBase = m.M
	v128 = v127
	goto L38
L37:
	;
	v128 = v123
	goto L38
L38:
	;
	goto L35
L39:
	;
	v132 = *(*int32)(unsafe.Add(mBase, uint32(v43)))
	v133 = *(*int32)(unsafe.Add(mBase, uint32(v43)+296))
	v134 = F_make_multirange(m, v132, v133, v120, v128)
	mBase = m.M
	v135 = m.ExcPending
	if v135 != 0 {
		goto L1
	} else {
		goto L43
	}
L40:
	;
	v130 = F__emscripten_memcpy_bulkmem(m, v128+v126, v117, v85)
	mBase = m.M
	goto L42
L41:
	;
	goto L42
L42:
	;
	goto L39
L43:
	;
	v139 = v134
	goto L5
L44:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v13))) = v29
	F_errmsg_internal(m, int32(370140), v13)
	mBase = m.M
	v157 = m.ExcPending
	if v157 != 0 {
		goto L1
	} else {
		goto L45
	}
L45:
	;
	F_errfinish(m, int32(493986), int32(558), int32(398748))
	mBase = m.M
	v162 = m.ExcPending
	if v162 != 0 {
		goto L1
	} else {
		goto L46
	}
L46:
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
			v14 = int32(4515248)
			v15 = *(*int32)(unsafe.Add(mBase, _consts[0]))
			v17 = *(*int32)(unsafe.Add(mBase, uint32(v10)+24))
			*(*int32)(unsafe.Add(mBase, _consts[0])) = v17
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
					v30 = F_lookup_type_cache(m, v28, int32(65536))
					mBase = m.M
					v31 = m.ExcPending
					if v31 != 0 {
						return int32(0)
					} else {
						*(*int32)(unsafe.Add(mBase, uint32(v23)+4)) = v30
						*(*int32)(unsafe.Add(mBase, uint32(v10)+16)) = v23
						*(*int32)(unsafe.Add(mBase, _consts[0])) = v15
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
					v23 = F_lookup_type_cache(m, v17, int32(65536))
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
								F_errmsg_internal(m, int32(370140), v9)
								mBase = m.M
								v59 = m.ExcPending
								if v59 != 0 {
									return int32(0)
								} else {
									F_errfinish(m, int32(493986), int32(558), int32(398748))
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
				v23 = F_lookup_type_cache(m, v17, int32(65536))
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
							F_errmsg_internal(m, int32(370140), v9)
							mBase = m.M
							v59 = m.ExcPending
							if v59 != 0 {
								return int32(0)
							} else {
								F_errfinish(m, int32(493986), int32(558), int32(398748))
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
