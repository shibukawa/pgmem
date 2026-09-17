package p1

import (
	base "github.com/shibukawa/pgmem/internal/aot/pgaot/base"
	"unsafe"
)

func F_build_regexp_split_result(m *base.Module, l0 int32) int32 {
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
	var v25 int32
	_ = v25
	var v29 int32
	_ = v29
	var v32 int32
	_ = v32
	var v33 int32
	_ = v33
	var v34 int32
	_ = v34
	var v38 int32
	_ = v38
	var v41 int32
	_ = v41
	var v42 int32
	_ = v42
	var v47 int32
	_ = v47
	var v51 int32
	_ = v51
	var v56 int32
	_ = v56
	var v60 int32
	_ = v60
	var v64 int32
	_ = v64
	var v69 int32
	_ = v69
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
			v60 = m.ExcPending
			if v60 != 0 {
				return int32(0)
			} else {
				F_errmsg_internal(m, int32(_a_F_build_regexp_split_result_0), int32(0))
				mBase = m.M
				v64 = m.ExcPending
				if v64 != 0 {
					return int32(0)
				} else {
					F_errfinish(m, int32(_a_F_build_regexp_split_result_1), int32(1878), int32(_a_F_build_regexp_split_result_2))
					mBase = m.M
					v69 = m.ExcPending
					if v69 != 0 {
						return int32(0)
					} else {
						base.Wasm_trap_unreachable()
						for {
						}
					}
				}
			}
		} else {
			v24 = v22 - v20
			if v10 != 0 {
				v25 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
				v29 = F_pg_wchar2mb_with_len(m, v25+v20<<(uint(int32(2))%32), v10, v24)
				mBase = m.M
				v32 = m.ExcPending
				if v32 != 0 {
					return int32(0)
				} else {
					v33 = F_cstring_to_text_with_len(m, v10, v29)
					mBase = m.M
					v34 = m.ExcPending
					if v34 != 0 {
						return int32(0)
					} else {
						return v33
					}
				}
			} else {
				v38 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
				v41 = F_DirectFunctionCall3Coll(m, int32(1477), int32(0), v38, v20+int32(1), v24)
				mBase = m.M
				v42 = m.ExcPending
				if v42 != 0 {
					return int32(0)
				} else {
					return v41
				}
			}
		}
	} else {
		v17 = *(*int32)(unsafe.Add(mBase, uint32(v8+v9-int32(4))))
		if v17 < int32(0) {
			F_errstart_cold(m, int32(21), int32(0))
			mBase = m.M
			v47 = m.ExcPending
			if v47 != 0 {
				return int32(0)
			} else {
				F_errmsg_internal(m, int32(_a_F_build_regexp_split_result_3), int32(0))
				mBase = m.M
				v51 = m.ExcPending
				if v51 != 0 {
					return int32(0)
				} else {
					F_errfinish(m, int32(_a_F_build_regexp_split_result_1), int32(1874), int32(_a_F_build_regexp_split_result_2))
					mBase = m.M
					v56 = m.ExcPending
					if v56 != 0 {
						return int32(0)
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
				v60 = m.ExcPending
				if v60 != 0 {
					return int32(0)
				} else {
					F_errmsg_internal(m, int32(_a_F_build_regexp_split_result_0), int32(0))
					mBase = m.M
					v64 = m.ExcPending
					if v64 != 0 {
						return int32(0)
					} else {
						F_errfinish(m, int32(_a_F_build_regexp_split_result_1), int32(1878), int32(_a_F_build_regexp_split_result_2))
						mBase = m.M
						v69 = m.ExcPending
						if v69 != 0 {
							return int32(0)
						} else {
							base.Wasm_trap_unreachable()
							for {
							}
						}
					}
				}
			} else {
				v24 = v22 - v20
				if v10 != 0 {
					v25 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
					v29 = F_pg_wchar2mb_with_len(m, v25+v20<<(uint(int32(2))%32), v10, v24)
					mBase = m.M
					v32 = m.ExcPending
					if v32 != 0 {
						return int32(0)
					} else {
						v33 = F_cstring_to_text_with_len(m, v10, v29)
						mBase = m.M
						v34 = m.ExcPending
						if v34 != 0 {
							return int32(0)
						} else {
							return v33
						}
					}
				} else {
					v38 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
					v41 = F_DirectFunctionCall3Coll(m, int32(1477), int32(0), v38, v20+int32(1), v24)
					mBase = m.M
					v42 = m.ExcPending
					if v42 != 0 {
						return int32(0)
					} else {
						return v41
					}
				}
			}
		}
	}
}
func F_regexp_instr_no_start(m *base.Module, l0 int32) int32 {
	var v2 int32
	_ = v2
	var v5 int32
	_ = v5
	v2 = F_regexp_instr(m, l0)
	v5 = m.ExcPending
	if v5 != 0 {
		return int32(0)
	} else {
		return v2
	}
}
func F_regexp_like(m *base.Module, l0 int32) int32 {
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
	var v22 int32
	_ = v22
	var v25 int32
	_ = v25
	var v26 int32
	_ = v26
	var v27 int32
	_ = v27
	var v29 int32
	_ = v29
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
	var v39 int32
	_ = v39
	var v45 int32
	_ = v45
	var v48 int32
	_ = v48
	var v55 int32
	_ = v55
	var v56 int32
	_ = v56
	var v60 int32
	_ = v60
	var v66 int32
	_ = v66
	var v67 int32
	_ = v67
	var v70 int32
	_ = v70
	var v71 int32
	_ = v71
	var v72 int32
	_ = v72
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
	var v84 int32
	_ = v84
	var v87 int32
	_ = v87
	var v88 int32
	_ = v88
	var v90 int32
	_ = v90
	var v98 int32
	_ = v98
	var v101 int32
	_ = v101
	var v106 int32
	_ = v106
	var v111 int32
	_ = v111
	v8 = m.G0
	v10 = v8 - int32(16)
	m.G0 = v10
	v12 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v13 = F_pg_detoast_datum_packed(m, v12)
	mBase = m.M
	v16 = m.ExcPending
	if v16 != 0 {
		return int32(0)
	} else {
		v17 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
		v18 = F_pg_detoast_datum_packed(m, v17)
		mBase = m.M
		v19 = m.ExcPending
		if v19 != 0 {
			return int32(0)
		} else {
			v22 = int32(*(*int16)(unsafe.Add(mBase, uint32(l0)+18)))
			if int32(3) <= v22 {
				v25 = *(*int32)(unsafe.Add(mBase, uint32(l0)+36))
				v26 = F_pg_detoast_datum_packed(m, v25)
				mBase = m.M
				v27 = m.ExcPending
				if v27 != 0 {
					return int32(0)
				} else {
					v29 = v26
					F_parse_re_flags(m, v10+int32(8), v29)
					mBase = m.M
					v31 = m.ExcPending
					if v31 != 0 {
						return int32(0)
					} else {
						v32 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v10)+12)))
						if v32 != int32(1) {
							v35 = int32(1)
							v36 = v13 + v35
							v37 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v13))))
							v39 = v37 & v35
							if v37 == v35 {
								v45 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v36))))
								if v45 == int32(18) {
									v48 = int32(16)
								} else {
									v48 = int32(0)
								}
								if base.Ui32((v45-int32(1))&int32(255)) < base.Ui32(int32(3)) {
									v55 = int32(4)
								} else {
									v55 = v48
								}
								v66 = v55
							} else {
								v56 = int32(1)
								if v39 != 0 {
									v66 = int32(base.Ui32(v37)>>(uint(v56)%32)) - v56
								} else {
									v60 = *(*int32)(unsafe.Add(mBase, uint32(v13)))
									v66 = int32(base.Ui32(v60)>>(uint(int32(2))%32)) - int32(4)
								}
							}
							v67 = *(*int32)(unsafe.Add(mBase, uint32(v10)+8))
							v70 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
							v71 = F_RE_compile_and_cache(m, v18, v67|int32(16), v70)
							mBase = m.M
							v72 = m.ExcPending
							if v72 != 0 {
								return int32(0)
							} else {
								v77 = F_palloc(m, v66<<(uint(int32(2))%32)+int32(4))
								mBase = m.M
								v78 = m.ExcPending
								if v78 != 0 {
									return int32(0)
								} else {
									if v39 != 0 {
										v81 = v36
									} else {
										v81 = v13 + int32(4)
									}
									v82 = F_pg_mb2wchar_with_len(m, v81, v77, v66)
									mBase = m.M
									v83 = m.ExcPending
									if v83 != 0 {
										return int32(0)
									} else {
										v84 = int32(0)
										v87 = F_RE_wchar_execute(m, v77, v82, v84, v84, v84)
										mBase = m.M
										v88 = m.ExcPending
										if v88 != 0 {
											return int32(0)
										} else {
											F_pfree(m, v77)
											mBase = m.M
											v90 = m.ExcPending
											if v90 != 0 {
												return int32(0)
											} else {
												m.G0 = v10 + int32(16)
												return v87
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
								return int32(0)
							} else {
								F_errcode(m, int32(50856066))
								mBase = m.M
								v101 = m.ExcPending
								if v101 != 0 {
									return int32(0)
								} else {
									*(*int32)(unsafe.Add(mBase, uint32(v10))) = int32(_a_F_regexp_like_0)
									F_errmsg(m, int32(_a_F_regexp_like_1), v10)
									mBase = m.M
									v106 = m.ExcPending
									if v106 != 0 {
										return int32(0)
									} else {
										F_errfinish(m, int32(_a_F_regexp_like_2), int32(1344), int32(_a_F_regexp_like_3))
										mBase = m.M
										v111 = m.ExcPending
										if v111 != 0 {
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
				}
			} else {
				v29 = int32(0)
				F_parse_re_flags(m, v10+int32(8), v29)
				mBase = m.M
				v31 = m.ExcPending
				if v31 != 0 {
					return int32(0)
				} else {
					v32 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v10)+12)))
					if v32 != int32(1) {
						v35 = int32(1)
						v36 = v13 + v35
						v37 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v13))))
						v39 = v37 & v35
						if v37 == v35 {
							v45 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v36))))
							if v45 == int32(18) {
								v48 = int32(16)
							} else {
								v48 = int32(0)
							}
							if base.Ui32((v45-int32(1))&int32(255)) < base.Ui32(int32(3)) {
								v55 = int32(4)
							} else {
								v55 = v48
							}
							v66 = v55
						} else {
							v56 = int32(1)
							if v39 != 0 {
								v66 = int32(base.Ui32(v37)>>(uint(v56)%32)) - v56
							} else {
								v60 = *(*int32)(unsafe.Add(mBase, uint32(v13)))
								v66 = int32(base.Ui32(v60)>>(uint(int32(2))%32)) - int32(4)
							}
						}
						v67 = *(*int32)(unsafe.Add(mBase, uint32(v10)+8))
						v70 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
						v71 = F_RE_compile_and_cache(m, v18, v67|int32(16), v70)
						mBase = m.M
						v72 = m.ExcPending
						if v72 != 0 {
							return int32(0)
						} else {
							v77 = F_palloc(m, v66<<(uint(int32(2))%32)+int32(4))
							mBase = m.M
							v78 = m.ExcPending
							if v78 != 0 {
								return int32(0)
							} else {
								if v39 != 0 {
									v81 = v36
								} else {
									v81 = v13 + int32(4)
								}
								v82 = F_pg_mb2wchar_with_len(m, v81, v77, v66)
								mBase = m.M
								v83 = m.ExcPending
								if v83 != 0 {
									return int32(0)
								} else {
									v84 = int32(0)
									v87 = F_RE_wchar_execute(m, v77, v82, v84, v84, v84)
									mBase = m.M
									v88 = m.ExcPending
									if v88 != 0 {
										return int32(0)
									} else {
										F_pfree(m, v77)
										mBase = m.M
										v90 = m.ExcPending
										if v90 != 0 {
											return int32(0)
										} else {
											m.G0 = v10 + int32(16)
											return v87
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
							return int32(0)
						} else {
							F_errcode(m, int32(50856066))
							mBase = m.M
							v101 = m.ExcPending
							if v101 != 0 {
								return int32(0)
							} else {
								*(*int32)(unsafe.Add(mBase, uint32(v10))) = int32(_a_F_regexp_like_0)
								F_errmsg(m, int32(_a_F_regexp_like_1), v10)
								mBase = m.M
								v106 = m.ExcPending
								if v106 != 0 {
									return int32(0)
								} else {
									F_errfinish(m, int32(_a_F_regexp_like_2), int32(1344), int32(_a_F_regexp_like_3))
									mBase = m.M
									v111 = m.ExcPending
									if v111 != 0 {
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
			}
		}
	}
}
func F_regexp_like_no_flags(m *base.Module, l0 int32) int32 {
	var v2 int32
	_ = v2
	var v5 int32
	_ = v5
	v2 = F_regexp_like(m, l0)
	v5 = m.ExcPending
	if v5 != 0 {
		return int32(0)
	} else {
		return v2
	}
}
func F_regexp_matches_no_flags(m *base.Module, l0 int32) int32 {
	var v2 int32
	_ = v2
	var v5 int32
	_ = v5
	v2 = F_regexp_matches(m, l0)
	v5 = m.ExcPending
	if v5 != 0 {
		return int32(0)
	} else {
		return v2
	}
}
func F_regexp_split_to_table_no_flags(m *base.Module, l0 int32) int32 {
	var v2 int32
	_ = v2
	var v5 int32
	_ = v5
	v2 = F_regexp_split_to_table(m, l0)
	v5 = m.ExcPending
	if v5 != 0 {
		return int32(0)
	} else {
		return v2
	}
}
func F_setup_regexp_matches(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32, l5 int32, l6 int32, l7 int32) int32 {
	mBase := m.M
	_ = mBase
	var v23 int32
	_ = v23
	var v26 int32
	_ = v26
	var v28 int32
	_ = v28
	var v29 int32
	_ = v29
	var v34 int32
	_ = v34
	var v36 int32
	_ = v36
	var v37 int32
	_ = v37
	var v43 int32
	_ = v43
	var v46 int32
	_ = v46
	var v53 int32
	_ = v53
	var v54 int32
	_ = v54
	var v60 int32
	_ = v60
	var v66 int32
	_ = v66
	var v68 int32
	_ = v68
	var v71 int32
	_ = v71
	var v72 int32
	_ = v72
	var v73 int32
	_ = v73
	var v75 int32
	_ = v75
	var v78 int32
	_ = v78
	var v80 int32
	_ = v80
	var v81 int32
	_ = v81
	var v82 int32
	_ = v82
	var v85 int32
	_ = v85
	var v86 int32
	_ = v86
	var v87 int32
	_ = v87
	var v91 int32
	_ = v91
	var v96 int32
	_ = v96
	var v97 int32
	_ = v97
	var v105 int32
	_ = v105
	var v106 int32
	_ = v106
	var v109 int32
	_ = v109
	var v110 int32
	_ = v110
	var v113 int32
	_ = v113
	var v114 int32
	_ = v114
	var v117 int32
	_ = v117
	var v118 int32
	_ = v118
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
	var v140 int32
	_ = v140
	var v144 int32
	_ = v144
	var v145 int32
	_ = v145
	var v148 int32
	_ = v148
	var v150 int32
	_ = v150
	var v152 int32
	_ = v152
	var v153 int32
	_ = v153
	var v154 int32
	_ = v154
	var v159 int32
	_ = v159
	var v181 int32
	_ = v181
	var v184 int32
	_ = v184
	var v186 int32
	_ = v186
	var v189 int32
	_ = v189
	var v190 int32
	_ = v190
	var v192 int32
	_ = v192
	var v197 int32
	_ = v197
	var v201 int32
	_ = v201
	var v222 int32
	_ = v222
	var v224 int32
	_ = v224
	var v226 int32
	_ = v226
	var v244 int32
	_ = v244
	var v245 int32
	_ = v245
	var v247 int32
	_ = v247
	var v248 int32
	_ = v248
	var v250 int32
	_ = v250
	var v252 int32
	_ = v252
	var v255 int32
	_ = v255
	var v257 int32
	_ = v257
	var v261 int32
	_ = v261
	var v263 int32
	_ = v263
	var v265 int32
	_ = v265
	var v266 int32
	_ = v266
	var v268 int32
	_ = v268
	var v269 int32
	_ = v269
	var v270 int32
	_ = v270
	var v271 int32
	_ = v271
	var v273 int32
	_ = v273
	var v275 int32
	_ = v275
	var v279 int32
	_ = v279
	var v283 int32
	_ = v283
	var v285 int32
	_ = v285
	var v287 int32
	_ = v287
	var v291 int32
	_ = v291
	var v307 int32
	_ = v307
	var v313 int32
	_ = v313
	var v316 int32
	_ = v316
	var v318 int32
	_ = v318
	var v320 int32
	_ = v320
	var v321 int32
	_ = v321
	var v322 int32
	_ = v322
	var v323 int32
	_ = v323
	var v327 int32
	_ = v327
	var v339 int32
	_ = v339
	var v343 int32
	_ = v343
	var v346 int32
	_ = v346
	var v347 int32
	_ = v347
	var v349 int32
	_ = v349
	var v352 int32
	_ = v352
	var v356 int32
	_ = v356
	var v368 int32
	_ = v368
	var v372 int32
	_ = v372
	var v373 int32
	_ = v373
	var v379 int32
	_ = v379
	var v381 int32
	_ = v381
	var v382 int32
	_ = v382
	var v385 int64
	_ = v385
	var v391 int32
	_ = v391
	var v392 int32
	_ = v392
	var v393 int32
	_ = v393
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
	var v412 int32
	_ = v412
	var v415 int32
	_ = v415
	var v419 int32
	_ = v419
	var v424 int32
	_ = v424
	v23 = F_palloc0(m, int32(40))
	mBase = m.M
	v26 = m.ExcPending
	if v26 != 0 {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	return int32(0)
L2:
	;
	v28 = *(*int32)(unsafe.Add(mBase, _c_F_setup_regexp_matches[0]))
	v29 = *(*int32)(unsafe.Add(mBase, uint32(v28)+4))
	v34 = *(*int32)(unsafe.Add(mBase, uint32(v29*int32(28))+uint32(_c_F_setup_regexp_matches[1])))
	goto L3
L3:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v23))) = l0
	v36 = int32(1)
	v37 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0))))
	if v37 == v36 {
		goto L5
	} else {
		goto L6
	}
L4:
	;
	v68 = v66 + int32(1)
	v71 = F_palloc(m, v68<<(uint(int32(2))%32))
	mBase = m.M
	v72 = m.ExcPending
	if v72 != 0 {
		goto L1
	} else {
		goto L15
	}
L5:
	;
	v43 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+1)))
	if v43 == int32(18) {
		goto L8
	} else {
		goto L9
	}
L6:
	;
	goto L7
L7:
	;
	v54 = int32(1)
	if v37&v54 != 0 {
		v66 = int32(base.Ui32(v37)>>(uint(v54)%32)) - v54
		goto L4
	} else {
		goto L14
	}
L8:
	;
	v46 = int32(16)
	goto L10
L9:
	;
	v46 = int32(0)
	goto L10
L10:
	;
	if base.Ui32((v43-int32(1))&int32(255)) < base.Ui32(int32(3)) {
		goto L11
	} else {
		goto L12
	}
L11:
	;
	v53 = int32(4)
	goto L13
L12:
	;
	v53 = v46
	goto L13
L13:
	;
	v66 = v53
	goto L4
L14:
	;
	v60 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v66 = int32(base.Ui32(v60)>>(uint(int32(2))%32)) - int32(4)
	goto L4
L15:
	;
	v73 = int32(1)
	v75 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0))))
	if v75&v73 != 0 {
		goto L16
	} else {
		goto L17
	}
L16:
	;
	v78 = v73
	goto L18
L17:
	;
	v78 = int32(4)
	goto L18
L18:
	;
	v80 = F_pg_mb2wchar_with_len(m, l0+v78, v71, v66)
	mBase = m.M
	v81 = m.ExcPending
	if v81 != 0 {
		goto L1
	} else {
		goto L19
	}
L19:
	;
	v82 = *(*int32)(unsafe.Add(mBase, uint32(l2)))
	if l5 != 0 {
		goto L20
	} else {
		goto L21
	}
L20:
	;
	v85 = v82
	goto L22
L21:
	;
	v85 = v82 | int32(16)
	goto L22
L22:
	;
	v86 = F_RE_compile_and_cache(m, l1, v85, l4)
	mBase = m.M
	v87 = m.ExcPending
	if v87 != 0 {
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
	v109 = F_palloc(m, v105<<(uint(int32(3))%32))
	mBase = m.M
	v110 = m.ExcPending
	if v110 != 0 {
		goto L1
	} else {
		goto L28
	}
L25:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v23)+8)) = int32(1)
	v105 = v36
	v106 = int32(0)
	goto L24
L26:
	;
	v91 = *(*int32)(unsafe.Add(mBase, _c_F_setup_regexp_matches[2]))
	if v91 == int32(0) {
		goto L25
	} else {
		goto L27
	}
L27:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v23)+8)) = v91
	v96 = *(*int32)(unsafe.Add(mBase, _c_F_setup_regexp_matches[2]))
	v97 = int32(1)
	v105 = v96 + v97
	v106 = v97
	goto L24
L28:
	;
	v113 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l2)+4)))
	if v113 != 0 {
		goto L29
	} else {
		goto L30
	}
L29:
	;
	v114 = int32(255)
	goto L31
L30:
	;
	v114 = int32(31)
	goto L31
L31:
	;
	v117 = F_palloc(m, v114<<(uint(int32(2))%32))
	mBase = m.M
	v118 = m.ExcPending
	if v118 != 0 {
		goto L1
	} else {
		goto L32
	}
L32:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v23)+12)) = v117
	v120 = int32(0)
	v123 = v114
	v124 = v120
	v126 = l3
	v127 = v120
	v128 = v120
	v140 = int32(0)
	goto L34
L33:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v412 = m.ExcPending
	if v412 != 0 {
		goto L1
	} else {
		goto L95
	}
L34:
	;
	v144 = F_RE_wchar_execute(m, v71, v80, v126, v105, v109)
	mBase = m.M
	v145 = m.ExcPending
	if v145 != 0 {
		goto L1
	} else {
		goto L37
	}
L35:
	;
	v372 = *(*int32)(unsafe.Add(mBase, uint32(v23)+12))
	v373 = int32(2)
	*(*int32)(unsafe.Add(mBase, uint32(v372+v352<<(uint(v373)%32)))) = v80
	if v373 <= v34 {
		goto L80
	} else {
		goto L81
	}
L36:
	;
	goto L35
L37:
	;
	if v144 == int32(0) {
		v352 = v124
		v356 = v128
		v368 = v140
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
	v343 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l2)+4)))
	if v343 != int32(1) {
		v352 = v323
		v356 = v327
		v368 = v339
		goto L36
	} else {
		goto L77
	}
L40:
	;
	v148 = *(*int32)(unsafe.Add(mBase, uint32(v109)))
	if v80 <= v148 {
		v322 = v123
		v323 = v124
		v327 = v128
		v339 = v140
		goto L39
	} else {
		goto L43
	}
L41:
	;
	goto L42
L42:
	;
	v152 = int32(1)
	v153 = v124 + v152
	v154 = *(*int32)(unsafe.Add(mBase, uint32(v23)+8))
	if v123 < v153+v154<<(uint(v152)%32) {
		goto L45
	} else {
		goto L46
	}
L43:
	;
	v150 = *(*int32)(unsafe.Add(mBase, uint32(v109)+4))
	if v150 <= v127 {
		v322 = v123
		v323 = v124
		v327 = v128
		v339 = v140
		goto L39
	} else {
		goto L44
	}
L44:
	;
	goto L42
L45:
	;
	v159 = v123
	goto L48
L46:
	;
	v197 = v123
	v201 = v154
	goto L47
L47:
	;
	if v106 != 0 {
		goto L54
	} else {
		goto L55
	}
L48:
	;
	v181 = v159 << (uint(int32(1)) % 32)
	if base.Ui32(int32(268435456)) <= base.Ui32(v181) {
		goto L33
	} else {
		goto L50
	}
L49:
	;
	v197 = v186
	v201 = v192
	goto L47
L50:
	;
	v184 = *(*int32)(unsafe.Add(mBase, uint32(v23)+12))
	v186 = v181 | int32(1)
	v189 = F_repalloc(m, v184, v186<<(uint(int32(2))%32))
	mBase = m.M
	v190 = m.ExcPending
	if v190 != 0 {
		goto L1
	} else {
		goto L51
	}
L51:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v23)+12)) = v189
	v192 = *(*int32)(unsafe.Add(mBase, uint32(v23)+8))
	if v186 < v153+v192<<(uint(int32(1))%32) {
		v159 = v186
		goto L48
	} else {
		goto L52
	}
L52:
	;
	goto L49
L53:
	;
	v307 = *(*int32)(unsafe.Add(mBase, uint32(v23)+4))
	*(*int32)(unsafe.Add(mBase, uint32(v23)+4)) = v307 + int32(1)
	if l7 == int32(0) {
		v320 = v291
		goto L71
	} else {
		goto L72
	}
L54:
	;
	if v201 <= int32(0) {
		v287 = v124
		v291 = v128
		goto L53
	} else {
		goto L57
	}
L55:
	;
	goto L56
L56:
	;
	v268 = *(*int32)(unsafe.Add(mBase, uint32(v109)+4))
	v269 = int32(2)
	v270 = v124 << (uint(v269) % 32)
	v271 = *(*int32)(unsafe.Add(mBase, uint32(v23)+12))
	v273 = *(*int32)(unsafe.Add(mBase, uint32(v109)))
	*(*int32)(unsafe.Add(mBase, uint32(v270+v271))) = v273
	v275 = *(*int32)(unsafe.Add(mBase, uint32(v23)+12))
	*(*int32)(unsafe.Add(mBase, uint32(v275+v270)+4)) = v268
	v279 = v124 + v269
	if v268|v273 < int32(0) {
		v287 = v279
		v291 = v128
		goto L53
	} else {
		goto L67
	}
L57:
	;
	v222 = v124
	v224 = int32(1)
	v226 = v128
	goto L58
L58:
	;
	v244 = v109 + v224<<(uint(int32(3))%32)
	v245 = *(*int32)(unsafe.Add(mBase, uint32(v244)+4))
	v247 = v222 << (uint(int32(2)) % 32)
	v248 = *(*int32)(unsafe.Add(mBase, uint32(v23)+12))
	v250 = *(*int32)(unsafe.Add(mBase, uint32(v244)))
	*(*int32)(unsafe.Add(mBase, uint32(v247+v248))) = v250
	v252 = *(*int32)(unsafe.Add(mBase, uint32(v23)+12))
	*(*int32)(unsafe.Add(mBase, uint32(v252+v247)+4)) = v245
	v255 = v245 - v250
	if v226 < v255 {
		goto L60
	} else {
		goto L61
	}
L59:
	;
	v287 = v263
	v291 = v261
	goto L53
L60:
	;
	v257 = v255
	goto L62
L61:
	;
	v257 = v226
	goto L62
L62:
	;
	if int32(0) <= v245|v250 {
		goto L63
	} else {
		goto L64
	}
L63:
	;
	v261 = v257
	goto L65
L64:
	;
	v261 = v226
	goto L65
L65:
	;
	v263 = v222 + int32(2)
	v265 = v224 + int32(1)
	v266 = *(*int32)(unsafe.Add(mBase, uint32(v23)+8))
	if v265 <= v266 {
		v222 = v263
		v224 = v265
		v226 = v261
		goto L58
	} else {
		goto L66
	}
L66:
	;
	goto L59
L67:
	;
	v283 = v268 - v273
	if v128 < v283 {
		goto L68
	} else {
		goto L69
	}
L68:
	;
	v285 = v283
	goto L70
L69:
	;
	v285 = v128
	goto L70
L70:
	;
	v287 = v279
	v291 = v285
	goto L53
L71:
	;
	v321 = *(*int32)(unsafe.Add(mBase, uint32(v109)+4))
	v322 = v197
	v323 = v287
	v327 = v320
	v339 = v321
	goto L39
L72:
	;
	v313 = *(*int32)(unsafe.Add(mBase, uint32(v109)))
	if v313 < int32(0) {
		v320 = v291
		goto L71
	} else {
		goto L73
	}
L73:
	;
	v316 = v313 - v140
	if v291 < v316 {
		goto L74
	} else {
		goto L75
	}
L74:
	;
	v318 = v316
	goto L76
L75:
	;
	v318 = v291
	goto L76
L76:
	;
	v320 = v318
	goto L71
L77:
	;
	v346 = *(*int32)(unsafe.Add(mBase, uint32(v109)+4))
	v347 = *(*int32)(unsafe.Add(mBase, uint32(v109)))
	v349 = v346 + base.B2i32(v346 == v347)
	if v349 <= v80 {
		v123 = v322
		v124 = v323
		v126 = v349
		v127 = v346
		v128 = v327
		v140 = v339
		goto L34
	} else {
		goto L78
	}
L78:
	;
	v352 = v323
	v356 = v327
	v368 = v339
	goto L36
L79:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v23)+36)) = v399
	*(*int32)(unsafe.Add(mBase, uint32(v23)+32)) = v402
	*(*int32)(unsafe.Add(mBase, uint32(v23)+28)) = v400
	F_pfree(m, v109)
	mBase = m.M
	v407 = m.ExcPending
	if v407 != 0 {
		goto L1
	} else {
		goto L94
	}
L80:
	;
	v379 = v80 - v368
	if v356 < v379 {
		goto L83
	} else {
		goto L84
	}
L81:
	;
	goto L82
L82:
	;
	F_pfree(m, v71)
	mBase = m.M
	v395 = m.ExcPending
	if v395 != 0 {
		goto L1
	} else {
		goto L93
	}
L83:
	;
	v381 = v379
	goto L85
L84:
	;
	v381 = v356
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
	v382 = v381
	goto L88
L87:
	;
	v382 = v356
	goto L88
L88:
	;
	v385 = base.I64_extend_i32_u(v382) * base.I64_extend_i32_u(v34)
	if base.I64_extend_i32_s(v66) < v385 {
		goto L89
	} else {
		goto L90
	}
L89:
	;
	v391 = v68
	goto L91
L90:
	;
	v391 = base.I32_wrap_i64(v385) + int32(1)
	goto L91
L91:
	;
	v392 = F_palloc(m, v391)
	mBase = m.M
	v393 = m.ExcPending
	if v393 != 0 {
		goto L1
	} else {
		goto L92
	}
L92:
	;
	v399 = v391
	v400 = v71
	v402 = v392
	goto L79
L93:
	;
	v396 = int32(0)
	v399 = v396
	v400 = v396
	v402 = v396
	goto L79
L94:
	;
	return v23
L95:
	;
	F_errcode(m, int32(261))
	mBase = m.M
	v415 = m.ExcPending
	if v415 != 0 {
		goto L1
	} else {
		goto L96
	}
L96:
	;
	F_errmsg(m, int32(_a_F_setup_regexp_matches_0), int32(0))
	mBase = m.M
	v419 = m.ExcPending
	if v419 != 0 {
		goto L1
	} else {
		goto L97
	}
L97:
	;
	F_errfinish(m, int32(_a_F_setup_regexp_matches_1), int32(1572), int32(_a_F_setup_regexp_matches_2))
	mBase = m.M
	v424 = m.ExcPending
	if v424 != 0 {
		goto L1
	} else {
		goto L98
	}
L98:
	;
	base.Wasm_trap_unreachable()
	for {
	}
}
