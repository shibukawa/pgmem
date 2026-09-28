package p1

import (
	base "github.com/shibukawa/pgmem/internal/aot/pgaot/base"
	"unsafe"
)

func F_build_regexp_split_result(m *base.Module, l0 int32) int64 {
	mBase := m.M
	_ = mBase
	var v6 int32
	_ = v6
	var v8 int32
	_ = v8
	var v9 int32
	_ = v9
	var v10 int32
	_ = v10
	var v17 int32
	_ = v17
	var v20 int32
	_ = v20
	var v22 int32
	_ = v22
	var v24 int32
	_ = v24
	var v29 int32
	_ = v29
	var v32 int32
	_ = v32
	var v33 int32
	_ = v33
	var v34 int32
	_ = v34
	var v39 int64
	_ = v39
	var v45 int64
	_ = v45
	var v46 int32
	_ = v46
	var v51 int32
	_ = v51
	var v55 int32
	_ = v55
	var v60 int32
	_ = v60
	var v64 int32
	_ = v64
	var v68 int32
	_ = v68
	var v73 int32
	_ = v73
	v6 = *(*int32)(unsafe.Add(mBase, uint32(l0)+16))
	v8 = v6 << (uint(int32(3)) % 32)
	v9 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	v10 = *(*int32)(unsafe.Add(mBase, uint32(l0)+32))
	if v6 <= int32(0) {
		v20 = int32(0)
		v22 = *(*int32)(unsafe.Add(mBase, uint32(v8+v9)))
		if v22 < v20 {
			F_errstart_cold(m, int32(21), int32(0))
			mBase = m.M
			v64 = m.ExcPending
			if v64 != 0 {
				return int64(0)
			} else {
				F_errmsg_internal(m, int32(_a_F_build_regexp_split_result_0), int32(0))
				mBase = m.M
				v68 = m.ExcPending
				if v68 != 0 {
					return int64(0)
				} else {
					F_errfinish(m, int32(_a_F_build_regexp_split_result_1), int32(1879), int32(_a_F_build_regexp_split_result_2))
					mBase = m.M
					v73 = m.ExcPending
					if v73 != 0 {
						return int64(0)
					} else {
						base.Wasm_trap_unreachable()
						for {
						}
					}
				}
			}
		} else {
			if v10 != 0 {
				v24 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
				v29 = F_pg_wchar2mb_with_len(m, v24+v20<<(uint(int32(2))%32), v10, v22-v20)
				mBase = m.M
				v32 = m.ExcPending
				if v32 != 0 {
					return int64(0)
				} else {
					v33 = F_cstring_to_text_with_len(m, v10, v29)
					mBase = m.M
					v34 = m.ExcPending
					if v34 != 0 {
						return int64(0)
					} else {
						return base.I64_extend_i32_u(v33)
					}
				}
			} else {
				v39 = int64(*(*uint32)(unsafe.Add(mBase, uint32(l0))))
				v45 = F_DirectFunctionCall3Coll(m, int32(1689), int32(0), v39, base.I64_extend_i32_s(v20+int32(1)), base.I64_extend_i32_s(v22-v20))
				mBase = m.M
				v46 = m.ExcPending
				if v46 != 0 {
					return int64(0)
				} else {
					return v45
				}
			}
		}
	} else {
		v17 = *(*int32)(unsafe.Add(mBase, uint32(v8+v9-int32(4))))
		if v17 < int32(0) {
			F_errstart_cold(m, int32(21), int32(0))
			mBase = m.M
			v51 = m.ExcPending
			if v51 != 0 {
				return int64(0)
			} else {
				F_errmsg_internal(m, int32(_a_F_build_regexp_split_result_3), int32(0))
				mBase = m.M
				v55 = m.ExcPending
				if v55 != 0 {
					return int64(0)
				} else {
					F_errfinish(m, int32(_a_F_build_regexp_split_result_1), int32(1875), int32(_a_F_build_regexp_split_result_2))
					mBase = m.M
					v60 = m.ExcPending
					if v60 != 0 {
						return int64(0)
					} else {
						base.Wasm_trap_unreachable()
						for {
						}
					}
				}
			}
		} else {
			v20 = v17
			v22 = *(*int32)(unsafe.Add(mBase, uint32(v8+v9)))
			if v22 < v20 {
				F_errstart_cold(m, int32(21), int32(0))
				mBase = m.M
				v64 = m.ExcPending
				if v64 != 0 {
					return int64(0)
				} else {
					F_errmsg_internal(m, int32(_a_F_build_regexp_split_result_0), int32(0))
					mBase = m.M
					v68 = m.ExcPending
					if v68 != 0 {
						return int64(0)
					} else {
						F_errfinish(m, int32(_a_F_build_regexp_split_result_1), int32(1879), int32(_a_F_build_regexp_split_result_2))
						mBase = m.M
						v73 = m.ExcPending
						if v73 != 0 {
							return int64(0)
						} else {
							base.Wasm_trap_unreachable()
							for {
							}
						}
					}
				}
			} else {
				if v10 != 0 {
					v24 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
					v29 = F_pg_wchar2mb_with_len(m, v24+v20<<(uint(int32(2))%32), v10, v22-v20)
					mBase = m.M
					v32 = m.ExcPending
					if v32 != 0 {
						return int64(0)
					} else {
						v33 = F_cstring_to_text_with_len(m, v10, v29)
						mBase = m.M
						v34 = m.ExcPending
						if v34 != 0 {
							return int64(0)
						} else {
							return base.I64_extend_i32_u(v33)
						}
					}
				} else {
					v39 = int64(*(*uint32)(unsafe.Add(mBase, uint32(l0))))
					v45 = F_DirectFunctionCall3Coll(m, int32(1689), int32(0), v39, base.I64_extend_i32_s(v20+int32(1)), base.I64_extend_i32_s(v22-v20))
					mBase = m.M
					v46 = m.ExcPending
					if v46 != 0 {
						return int64(0)
					} else {
						return v45
					}
				}
			}
		}
	}
}
func F_regexp_instr_no_start(m *base.Module, l0 int32) int64 {
	var v2 int64
	_ = v2
	var v5 int32
	_ = v5
	v2 = F_regexp_instr(m, l0)
	v5 = m.ExcPending
	if v5 != 0 {
		return int64(0)
	} else {
		return v2
	}
}
func F_regexp_like(m *base.Module, l0 int32) int64 {
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
	var v21 int32
	_ = v21
	var v24 int32
	_ = v24
	var v25 int32
	_ = v25
	var v26 int32
	_ = v26
	var v28 int32
	_ = v28
	var v30 int32
	_ = v30
	var v31 int32
	_ = v31
	var v34 int32
	_ = v34
	var v40 int32
	_ = v40
	var v43 int32
	_ = v43
	var v50 int32
	_ = v50
	var v51 int32
	_ = v51
	var v57 int32
	_ = v57
	var v63 int32
	_ = v63
	var v64 int32
	_ = v64
	var v67 int32
	_ = v67
	var v68 int32
	_ = v68
	var v69 int32
	_ = v69
	var v73 int32
	_ = v73
	var v74 int32
	_ = v74
	var v75 int32
	_ = v75
	var v79 int32
	_ = v79
	var v81 int32
	_ = v81
	var v82 int32
	_ = v82
	var v83 int32
	_ = v83
	var v86 int32
	_ = v86
	var v87 int32
	_ = v87
	var v89 int32
	_ = v89
	var v98 int32
	_ = v98
	var v101 int32
	_ = v101
	var v106 int32
	_ = v106
	var v111 int32
	_ = v111
	v7 = m.G0
	v9 = v7 - int32(16)
	m.G0 = v9
	v11 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
	v12 = F_pg_detoast_datum_packed(m, v11)
	mBase = m.M
	v15 = m.ExcPending
	if v15 != 0 {
		return int64(0)
	} else {
		v16 = *(*int32)(unsafe.Add(mBase, uint32(l0)+40))
		v17 = F_pg_detoast_datum_packed(m, v16)
		mBase = m.M
		v18 = m.ExcPending
		if v18 != 0 {
			return int64(0)
		} else {
			v21 = int32(*(*int16)(unsafe.Add(mBase, uint32(l0)+18)))
			if int32(3) <= v21 {
				v24 = *(*int32)(unsafe.Add(mBase, uint32(l0)+56))
				v25 = F_pg_detoast_datum_packed(m, v24)
				mBase = m.M
				v26 = m.ExcPending
				if v26 != 0 {
					return int64(0)
				} else {
					v28 = v25
					F_parse_re_flags(m, v9+int32(8), v28)
					mBase = m.M
					v30 = m.ExcPending
					if v30 != 0 {
						return int64(0)
					} else {
						v31 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v9)+12)))
						if v31 != int32(1) {
							v34 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v12))))
							if v34 == int32(1) {
								v40 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v12)+1)))
								if v40 == int32(18) {
									v43 = int32(16)
								} else {
									v43 = int32(0)
								}
								if base.Ui32((v40-int32(1))&int32(255)) < base.Ui32(int32(3)) {
									v50 = int32(4)
								} else {
									v50 = v43
								}
								v63 = v50
							} else {
								v51 = int32(1)
								if v34&v51 != 0 {
									v63 = int32(base.Ui32(v34)>>(uint(v51)%32)) - v51
								} else {
									v57 = *(*int32)(unsafe.Add(mBase, uint32(v12)))
									v63 = int32(base.Ui32(v57)>>(uint(int32(2))%32)) - int32(4)
								}
							}
							v64 = *(*int32)(unsafe.Add(mBase, uint32(v9)+8))
							v67 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
							v68 = F_RE_compile_and_cache(m, v17, v64|int32(16), v67)
							mBase = m.M
							v69 = m.ExcPending
							if v69 != 0 {
								return int64(0)
							} else {
								v73 = F_palloc_mul(m, int32(4), v63+int32(1))
								mBase = m.M
								v74 = m.ExcPending
								if v74 != 0 {
									return int64(0)
								} else {
									v75 = int32(1)
									if v34&v75 != 0 {
										v79 = v75
									} else {
										v79 = int32(4)
									}
									v81 = F_pg_mb2wchar_with_len(m, v12+v79, v73, v63)
									mBase = m.M
									v82 = m.ExcPending
									if v82 != 0 {
										return int64(0)
									} else {
										v83 = int32(0)
										v86 = F_RE_wchar_execute(m, v73, v81, v83, v83, v83)
										mBase = m.M
										v87 = m.ExcPending
										if v87 != 0 {
											return int64(0)
										} else {
											F_pfree(m, v73)
											mBase = m.M
											v89 = m.ExcPending
											if v89 != 0 {
												return int64(0)
											} else {
												m.G0 = v9 + int32(16)
												return base.I64_extend_i32_u(v86)
											}
										}
									}
								}
							}
						} else {
							F_errstart_cold(m, int32(21), int32(0))
							mBase = m.M
							v98 = m.ExcPending
							if v98 != 0 {
								return int64(0)
							} else {
								F_errcode(m, int32(50856066))
								mBase = m.M
								v101 = m.ExcPending
								if v101 != 0 {
									return int64(0)
								} else {
									*(*int32)(unsafe.Add(mBase, uint32(v9))) = int32(_a_F_regexp_like_0)
									F_errmsg(m, int32(_a_F_regexp_like_1), v9)
									mBase = m.M
									v106 = m.ExcPending
									if v106 != 0 {
										return int64(0)
									} else {
										F_errfinish(m, int32(_a_F_regexp_like_2), int32(1344), int32(_a_F_regexp_like_3))
										mBase = m.M
										v111 = m.ExcPending
										if v111 != 0 {
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
				}
			} else {
				v28 = int32(0)
				F_parse_re_flags(m, v9+int32(8), v28)
				mBase = m.M
				v30 = m.ExcPending
				if v30 != 0 {
					return int64(0)
				} else {
					v31 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v9)+12)))
					if v31 != int32(1) {
						v34 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v12))))
						if v34 == int32(1) {
							v40 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v12)+1)))
							if v40 == int32(18) {
								v43 = int32(16)
							} else {
								v43 = int32(0)
							}
							if base.Ui32((v40-int32(1))&int32(255)) < base.Ui32(int32(3)) {
								v50 = int32(4)
							} else {
								v50 = v43
							}
							v63 = v50
						} else {
							v51 = int32(1)
							if v34&v51 != 0 {
								v63 = int32(base.Ui32(v34)>>(uint(v51)%32)) - v51
							} else {
								v57 = *(*int32)(unsafe.Add(mBase, uint32(v12)))
								v63 = int32(base.Ui32(v57)>>(uint(int32(2))%32)) - int32(4)
							}
						}
						v64 = *(*int32)(unsafe.Add(mBase, uint32(v9)+8))
						v67 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
						v68 = F_RE_compile_and_cache(m, v17, v64|int32(16), v67)
						mBase = m.M
						v69 = m.ExcPending
						if v69 != 0 {
							return int64(0)
						} else {
							v73 = F_palloc_mul(m, int32(4), v63+int32(1))
							mBase = m.M
							v74 = m.ExcPending
							if v74 != 0 {
								return int64(0)
							} else {
								v75 = int32(1)
								if v34&v75 != 0 {
									v79 = v75
								} else {
									v79 = int32(4)
								}
								v81 = F_pg_mb2wchar_with_len(m, v12+v79, v73, v63)
								mBase = m.M
								v82 = m.ExcPending
								if v82 != 0 {
									return int64(0)
								} else {
									v83 = int32(0)
									v86 = F_RE_wchar_execute(m, v73, v81, v83, v83, v83)
									mBase = m.M
									v87 = m.ExcPending
									if v87 != 0 {
										return int64(0)
									} else {
										F_pfree(m, v73)
										mBase = m.M
										v89 = m.ExcPending
										if v89 != 0 {
											return int64(0)
										} else {
											m.G0 = v9 + int32(16)
											return base.I64_extend_i32_u(v86)
										}
									}
								}
							}
						}
					} else {
						F_errstart_cold(m, int32(21), int32(0))
						mBase = m.M
						v98 = m.ExcPending
						if v98 != 0 {
							return int64(0)
						} else {
							F_errcode(m, int32(50856066))
							mBase = m.M
							v101 = m.ExcPending
							if v101 != 0 {
								return int64(0)
							} else {
								*(*int32)(unsafe.Add(mBase, uint32(v9))) = int32(_a_F_regexp_like_0)
								F_errmsg(m, int32(_a_F_regexp_like_1), v9)
								mBase = m.M
								v106 = m.ExcPending
								if v106 != 0 {
									return int64(0)
								} else {
									F_errfinish(m, int32(_a_F_regexp_like_2), int32(1344), int32(_a_F_regexp_like_3))
									mBase = m.M
									v111 = m.ExcPending
									if v111 != 0 {
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
			}
		}
	}
}
func F_regexp_like_no_flags(m *base.Module, l0 int32) int64 {
	var v2 int64
	_ = v2
	var v5 int32
	_ = v5
	v2 = F_regexp_like(m, l0)
	v5 = m.ExcPending
	if v5 != 0 {
		return int64(0)
	} else {
		return v2
	}
}
func F_regexp_matches_no_flags(m *base.Module, l0 int32) int64 {
	var v2 int64
	_ = v2
	var v5 int32
	_ = v5
	v2 = F_regexp_matches(m, l0)
	v5 = m.ExcPending
	if v5 != 0 {
		return int64(0)
	} else {
		return v2
	}
}
func F_regexp_split_to_table_no_flags(m *base.Module, l0 int32) int64 {
	var v2 int64
	_ = v2
	var v5 int32
	_ = v5
	v2 = F_regexp_split_to_table(m, l0)
	v5 = m.ExcPending
	if v5 != 0 {
		return int64(0)
	} else {
		return v2
	}
}
func F_setup_regexp_matches(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32, l5 int32, l6 int32, l7 int32) int32 {
	mBase := m.M
	_ = mBase
	var v20 int32
	_ = v20
	var v23 int32
	_ = v23
	var v25 int32
	_ = v25
	var v26 int32
	_ = v26
	var v31 int32
	_ = v31
	var v33 int32
	_ = v33
	var v35 int32
	_ = v35
	var v41 int32
	_ = v41
	var v44 int32
	_ = v44
	var v51 int32
	_ = v51
	var v52 int32
	_ = v52
	var v58 int32
	_ = v58
	var v64 int32
	_ = v64
	var v67 int32
	_ = v67
	var v68 int32
	_ = v68
	var v69 int32
	_ = v69
	var v71 int32
	_ = v71
	var v74 int32
	_ = v74
	var v76 int32
	_ = v76
	var v77 int32
	_ = v77
	var v78 int32
	_ = v78
	var v81 int32
	_ = v81
	var v82 int32
	_ = v82
	var v83 int32
	_ = v83
	var v87 int32
	_ = v87
	var v92 int32
	_ = v92
	var v93 int32
	_ = v93
	var v101 int32
	_ = v101
	var v102 int32
	_ = v102
	var v104 int32
	_ = v104
	var v105 int32
	_ = v105
	var v109 int32
	_ = v109
	var v110 int32
	_ = v110
	var v111 int32
	_ = v111
	var v112 int32
	_ = v112
	var v114 int32
	_ = v114
	var v117 int32
	_ = v117
	var v120 int32
	_ = v120
	var v121 int32
	_ = v121
	var v122 int32
	_ = v122
	var v126 int32
	_ = v126
	var v133 int32
	_ = v133
	var v135 int32
	_ = v135
	var v136 int32
	_ = v136
	var v139 int32
	_ = v139
	var v141 int32
	_ = v141
	var v143 int32
	_ = v143
	var v144 int32
	_ = v144
	var v145 int32
	_ = v145
	var v150 int32
	_ = v150
	var v169 int32
	_ = v169
	var v172 int32
	_ = v172
	var v174 int32
	_ = v174
	var v177 int32
	_ = v177
	var v178 int32
	_ = v178
	var v180 int32
	_ = v180
	var v185 int32
	_ = v185
	var v194 int32
	_ = v194
	var v209 int32
	_ = v209
	var v210 int32
	_ = v210
	var v211 int32
	_ = v211
	var v226 int32
	_ = v226
	var v227 int32
	_ = v227
	var v229 int32
	_ = v229
	var v230 int32
	_ = v230
	var v232 int32
	_ = v232
	var v234 int32
	_ = v234
	var v237 int32
	_ = v237
	var v239 int32
	_ = v239
	var v243 int32
	_ = v243
	var v245 int32
	_ = v245
	var v247 int32
	_ = v247
	var v248 int32
	_ = v248
	var v250 int32
	_ = v250
	var v251 int32
	_ = v251
	var v252 int32
	_ = v252
	var v253 int32
	_ = v253
	var v255 int32
	_ = v255
	var v257 int32
	_ = v257
	var v261 int32
	_ = v261
	var v265 int32
	_ = v265
	var v267 int32
	_ = v267
	var v272 int32
	_ = v272
	var v273 int32
	_ = v273
	var v286 int32
	_ = v286
	var v292 int32
	_ = v292
	var v295 int32
	_ = v295
	var v297 int32
	_ = v297
	var v299 int32
	_ = v299
	var v300 int32
	_ = v300
	var v301 int32
	_ = v301
	var v305 int32
	_ = v305
	var v306 int32
	_ = v306
	var v317 int32
	_ = v317
	var v319 int32
	_ = v319
	var v322 int32
	_ = v322
	var v323 int32
	_ = v323
	var v325 int32
	_ = v325
	var v331 int32
	_ = v331
	var v332 int32
	_ = v332
	var v343 int32
	_ = v343
	var v345 int32
	_ = v345
	var v346 int32
	_ = v346
	var v352 int32
	_ = v352
	var v354 int32
	_ = v354
	var v355 int32
	_ = v355
	var v358 int32
	_ = v358
	var v359 int32
	_ = v359
	var v360 int32
	_ = v360
	var v362 int32
	_ = v362
	var v363 int32
	_ = v363
	var v366 int32
	_ = v366
	var v367 int32
	_ = v367
	var v368 int32
	_ = v368
	var v373 int32
	_ = v373
	var v378 int32
	_ = v378
	var v381 int32
	_ = v381
	var v385 int32
	_ = v385
	var v390 int32
	_ = v390
	v20 = F_palloc0(m, int32(40))
	mBase = m.M
	v23 = m.ExcPending
	if v23 != 0 {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	return int32(0)
L2:
	;
	v25 = *(*int32)(unsafe.Add(mBase, _c_F_setup_regexp_matches[0]))
	v26 = *(*int32)(unsafe.Add(mBase, uint32(v25)+4))
	v31 = *(*int32)(unsafe.Add(mBase, uint32(v26*int32(28))+uint32(_c_F_setup_regexp_matches[1])))
	goto L3
L3:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v20))) = l0
	v33 = int32(1)
	v35 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0))))
	if v35 == v33 {
		goto L5
	} else {
		goto L6
	}
L4:
	;
	v67 = F_palloc_mul(m, int32(4), v64+int32(1))
	mBase = m.M
	v68 = m.ExcPending
	if v68 != 0 {
		goto L1
	} else {
		goto L15
	}
L5:
	;
	v41 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+1)))
	if v41 == int32(18) {
		goto L8
	} else {
		goto L9
	}
L6:
	;
	goto L7
L7:
	;
	v52 = int32(1)
	if v35&v52 != 0 {
		v64 = int32(base.Ui32(v35)>>(uint(v52)%32)) - v52
		goto L4
	} else {
		goto L14
	}
L8:
	;
	v44 = int32(16)
	goto L10
L9:
	;
	v44 = int32(0)
	goto L10
L10:
	;
	if base.Ui32((v41-int32(1))&int32(255)) < base.Ui32(int32(3)) {
		goto L11
	} else {
		goto L12
	}
L11:
	;
	v51 = int32(4)
	goto L13
L12:
	;
	v51 = v44
	goto L13
L13:
	;
	v64 = v51
	goto L4
L14:
	;
	v58 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v64 = int32(base.Ui32(v58)>>(uint(int32(2))%32)) - int32(4)
	goto L4
L15:
	;
	v69 = int32(1)
	v71 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0))))
	if v71&v69 != 0 {
		goto L16
	} else {
		goto L17
	}
L16:
	;
	v74 = v69
	goto L18
L17:
	;
	v74 = int32(4)
	goto L18
L18:
	;
	v76 = F_pg_mb2wchar_with_len(m, l0+v74, v67, v64)
	mBase = m.M
	v77 = m.ExcPending
	if v77 != 0 {
		goto L1
	} else {
		goto L19
	}
L19:
	;
	v78 = *(*int32)(unsafe.Add(mBase, uint32(l2)))
	if l5 != 0 {
		goto L20
	} else {
		goto L21
	}
L20:
	;
	v81 = v78
	goto L22
L21:
	;
	v81 = v78 | int32(16)
	goto L22
L22:
	;
	v82 = F_RE_compile_and_cache(m, l1, v81, l4)
	mBase = m.M
	v83 = m.ExcPending
	if v83 != 0 {
		goto L1
	} else {
		goto L23
	}
L23:
	;
	if l5 == int32(0) {
		goto L25
	} else {
		goto L26
	}
L24:
	;
	v104 = F_palloc_mul(m, int32(8), v101)
	mBase = m.M
	v105 = m.ExcPending
	if v105 != 0 {
		goto L1
	} else {
		goto L28
	}
L25:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v20)+8)) = int32(1)
	v101 = v33
	v102 = int32(0)
	goto L24
L26:
	;
	v87 = *(*int32)(unsafe.Add(mBase, _c_F_setup_regexp_matches[2]))
	if v87 == int32(0) {
		goto L25
	} else {
		goto L27
	}
L27:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v20)+8)) = v87
	v92 = *(*int32)(unsafe.Add(mBase, _c_F_setup_regexp_matches[2]))
	v93 = int32(1)
	v101 = v92 + v93
	v102 = v93
	goto L24
L28:
	;
	v109 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l2)+4)))
	if v109 != 0 {
		goto L29
	} else {
		goto L30
	}
L29:
	;
	v110 = int32(255)
	goto L31
L30:
	;
	v110 = int32(31)
	goto L31
L31:
	;
	v111 = F_palloc_mul(m, int32(4), v110)
	mBase = m.M
	v112 = m.ExcPending
	if v112 != 0 {
		goto L1
	} else {
		goto L32
	}
L32:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v20)+12)) = v111
	v114 = int32(0)
	v117 = v110
	v120 = l3
	v121 = v114
	v122 = v114
	v126 = v114
	v133 = int32(0)
	goto L34
L33:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v378 = m.ExcPending
	if v378 != 0 {
		goto L1
	} else {
		goto L92
	}
L34:
	;
	v135 = F_RE_wchar_execute(m, v67, v76, v120, v101, v104)
	mBase = m.M
	v136 = m.ExcPending
	if v136 != 0 {
		goto L1
	} else {
		goto L37
	}
L35:
	;
	v345 = *(*int32)(unsafe.Add(mBase, uint32(v20)+12))
	v346 = int32(2)
	*(*int32)(unsafe.Add(mBase, uint32(v345+v331<<(uint(v346)%32)))) = v76
	if v346 <= v31 {
		goto L80
	} else {
		goto L81
	}
L36:
	;
	goto L35
L37:
	;
	if v135 == int32(0) {
		v331 = v121
		v332 = v122
		v343 = v133
		goto L36
	} else {
		goto L38
	}
L38:
	;
	if l6 != 0 {
		goto L40
	} else {
		goto L41
	}
L39:
	;
	v319 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l2)+4)))
	if v319 != int32(1) {
		v331 = v305
		v332 = v306
		v343 = v317
		goto L36
	} else {
		goto L77
	}
L40:
	;
	v139 = *(*int32)(unsafe.Add(mBase, uint32(v104)))
	if v76 <= v139 {
		v301 = v117
		v305 = v121
		v306 = v122
		v317 = v133
		goto L39
	} else {
		goto L43
	}
L41:
	;
	goto L42
L42:
	;
	v143 = int32(1)
	v144 = v121 + v143
	v145 = *(*int32)(unsafe.Add(mBase, uint32(v20)+8))
	if v117 < v144+v145<<(uint(v143)%32) {
		goto L45
	} else {
		goto L46
	}
L43:
	;
	v141 = *(*int32)(unsafe.Add(mBase, uint32(v104)+4))
	if v141 <= v126 {
		v301 = v117
		v305 = v121
		v306 = v122
		v317 = v133
		goto L39
	} else {
		goto L44
	}
L44:
	;
	goto L42
L45:
	;
	v150 = v117
	goto L48
L46:
	;
	v185 = v117
	v194 = v145
	goto L47
L47:
	;
	if v102 != 0 {
		goto L54
	} else {
		goto L55
	}
L48:
	;
	v169 = v150 << (uint(int32(1)) % 32)
	if base.Ui32(int32(268435456)) <= base.Ui32(v169) {
		goto L33
	} else {
		goto L50
	}
L49:
	;
	v185 = v174
	v194 = v180
	goto L47
L50:
	;
	v172 = *(*int32)(unsafe.Add(mBase, uint32(v20)+12))
	v174 = v169 | int32(1)
	v177 = F_repalloc(m, v172, v174<<(uint(int32(2))%32))
	mBase = m.M
	v178 = m.ExcPending
	if v178 != 0 {
		goto L1
	} else {
		goto L51
	}
L51:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v20)+12)) = v177
	v180 = *(*int32)(unsafe.Add(mBase, uint32(v20)+8))
	if v174 < v144+v180<<(uint(int32(1))%32) {
		v150 = v174
		goto L48
	} else {
		goto L52
	}
L52:
	;
	goto L49
L53:
	;
	v286 = *(*int32)(unsafe.Add(mBase, uint32(v20)+4))
	*(*int32)(unsafe.Add(mBase, uint32(v20)+4)) = v286 + int32(1)
	if l7 == int32(0) {
		v299 = v273
		goto L71
	} else {
		goto L72
	}
L54:
	;
	if v194 <= int32(0) {
		v272 = v121
		v273 = v122
		goto L53
	} else {
		goto L57
	}
L55:
	;
	goto L56
L56:
	;
	v250 = *(*int32)(unsafe.Add(mBase, uint32(v104)+4))
	v251 = int32(2)
	v252 = v121 << (uint(v251) % 32)
	v253 = *(*int32)(unsafe.Add(mBase, uint32(v20)+12))
	v255 = *(*int32)(unsafe.Add(mBase, uint32(v104)))
	*(*int32)(unsafe.Add(mBase, uint32(v252+v253))) = v255
	v257 = *(*int32)(unsafe.Add(mBase, uint32(v20)+12))
	*(*int32)(unsafe.Add(mBase, uint32(v257+v252)+4)) = v250
	v261 = v121 + v251
	if v250|v255 < int32(0) {
		v272 = v261
		v273 = v122
		goto L53
	} else {
		goto L67
	}
L57:
	;
	v209 = int32(1)
	v210 = v121
	v211 = v122
	goto L58
L58:
	;
	v226 = v104 + v209<<(uint(int32(3))%32)
	v227 = *(*int32)(unsafe.Add(mBase, uint32(v226)+4))
	v229 = v210 << (uint(int32(2)) % 32)
	v230 = *(*int32)(unsafe.Add(mBase, uint32(v20)+12))
	v232 = *(*int32)(unsafe.Add(mBase, uint32(v226)))
	*(*int32)(unsafe.Add(mBase, uint32(v229+v230))) = v232
	v234 = *(*int32)(unsafe.Add(mBase, uint32(v20)+12))
	*(*int32)(unsafe.Add(mBase, uint32(v234+v229)+4)) = v227
	v237 = v227 - v232
	if v211 < v237 {
		goto L60
	} else {
		goto L61
	}
L59:
	;
	v272 = v245
	v273 = v243
	goto L53
L60:
	;
	v239 = v237
	goto L62
L61:
	;
	v239 = v211
	goto L62
L62:
	;
	if int32(0) <= v227|v232 {
		goto L63
	} else {
		goto L64
	}
L63:
	;
	v243 = v239
	goto L65
L64:
	;
	v243 = v211
	goto L65
L65:
	;
	v245 = v210 + int32(2)
	v247 = v209 + int32(1)
	v248 = *(*int32)(unsafe.Add(mBase, uint32(v20)+8))
	if v247 <= v248 {
		v209 = v247
		v210 = v245
		v211 = v243
		goto L58
	} else {
		goto L66
	}
L66:
	;
	goto L59
L67:
	;
	v265 = v250 - v255
	if v122 < v265 {
		goto L68
	} else {
		goto L69
	}
L68:
	;
	v267 = v265
	goto L70
L69:
	;
	v267 = v122
	goto L70
L70:
	;
	v272 = v261
	v273 = v267
	goto L53
L71:
	;
	v300 = *(*int32)(unsafe.Add(mBase, uint32(v104)+4))
	v301 = v185
	v305 = v272
	v306 = v299
	v317 = v300
	goto L39
L72:
	;
	v292 = *(*int32)(unsafe.Add(mBase, uint32(v104)))
	if v292 < int32(0) {
		v299 = v273
		goto L71
	} else {
		goto L73
	}
L73:
	;
	v295 = v292 - v133
	if v273 < v295 {
		goto L74
	} else {
		goto L75
	}
L74:
	;
	v297 = v295
	goto L76
L75:
	;
	v297 = v273
	goto L76
L76:
	;
	v299 = v297
	goto L71
L77:
	;
	v322 = *(*int32)(unsafe.Add(mBase, uint32(v104)+4))
	v323 = *(*int32)(unsafe.Add(mBase, uint32(v104)))
	v325 = v322 + base.B2i32(v322 == v323)
	if v325 <= v76 {
		v117 = v301
		v120 = v325
		v121 = v305
		v122 = v306
		v126 = v322
		v133 = v317
		goto L34
	} else {
		goto L78
	}
L78:
	;
	v331 = v305
	v332 = v306
	v343 = v317
	goto L36
L79:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v20)+36)) = v366
	*(*int32)(unsafe.Add(mBase, uint32(v20)+32)) = v368
	*(*int32)(unsafe.Add(mBase, uint32(v20)+28)) = v367
	F_pfree(m, v104)
	mBase = m.M
	v373 = m.ExcPending
	if v373 != 0 {
		goto L1
	} else {
		goto L91
	}
L80:
	;
	v352 = v76 - v343
	if v332 < v352 {
		goto L83
	} else {
		goto L84
	}
L81:
	;
	goto L82
L82:
	;
	F_pfree(m, v67)
	mBase = m.M
	v362 = m.ExcPending
	if v362 != 0 {
		goto L1
	} else {
		goto L90
	}
L83:
	;
	v354 = v352
	goto L85
L84:
	;
	v354 = v332
	goto L85
L85:
	;
	if l7 != 0 {
		goto L86
	} else {
		goto L87
	}
L86:
	;
	v355 = v354
	goto L88
L87:
	;
	v355 = v332
	goto L88
L88:
	;
	v358 = v355*v31 + int32(1)
	v359 = F_palloc(m, v358)
	mBase = m.M
	v360 = m.ExcPending
	if v360 != 0 {
		goto L1
	} else {
		goto L89
	}
L89:
	;
	v366 = v358
	v367 = v67
	v368 = v359
	goto L79
L90:
	;
	v363 = int32(0)
	v366 = v363
	v367 = v363
	v368 = v363
	goto L79
L91:
	;
	return v20
L92:
	;
	F_errcode(m, int32(261))
	mBase = m.M
	v381 = m.ExcPending
	if v381 != 0 {
		goto L1
	} else {
		goto L93
	}
L93:
	;
	F_errmsg(m, int32(_a_F_setup_regexp_matches_0), int32(0))
	mBase = m.M
	v385 = m.ExcPending
	if v385 != 0 {
		goto L1
	} else {
		goto L94
	}
L94:
	;
	F_errfinish(m, int32(_a_F_setup_regexp_matches_1), int32(1572), int32(_a_F_setup_regexp_matches_2))
	mBase = m.M
	v390 = m.ExcPending
	if v390 != 0 {
		goto L1
	} else {
		goto L95
	}
L95:
	;
	base.Wasm_trap_unreachable()
	for {
	}
}
