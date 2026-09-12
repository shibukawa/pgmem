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
				F_errmsg_internal(m, int32(249660), int32(0))
				mBase = m.M
				v64 = m.ExcPending
				if v64 != 0 {
					return int32(0)
				} else {
					F_errfinish(m, int32(495763), int32(1878), int32(97800))
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
				v41 = F_DirectFunctionCall3Coll(m, int32(1493), int32(0), v38, v20+int32(1), v24)
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
				F_errmsg_internal(m, int32(249692), int32(0))
				mBase = m.M
				v51 = m.ExcPending
				if v51 != 0 {
					return int32(0)
				} else {
					F_errfinish(m, int32(495763), int32(1874), int32(97800))
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
					F_errmsg_internal(m, int32(249660), int32(0))
					mBase = m.M
					v64 = m.ExcPending
					if v64 != 0 {
						return int32(0)
					} else {
						F_errfinish(m, int32(495763), int32(1878), int32(97800))
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
					v41 = F_DirectFunctionCall3Coll(m, int32(1493), int32(0), v38, v20+int32(1), v24)
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
	var v28 int32
	_ = v28
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
	var v38 int32
	_ = v38
	var v41 int32
	_ = v41
	var v43 int32
	_ = v43
	var v52 int32
	_ = v52
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
					v28 = v26
					F_parse_re_flags(m, v10+int32(8), v28)
					mBase = m.M
					v30 = m.ExcPending
					if v30 != 0 {
						return int32(0)
					} else {
						v31 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v10)+12)))
						if v31 != int32(1) {
							v34 = int32(1)
							v35 = v13 + v34
							v36 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v13))))
							v38 = v36 & v34
							if v36 == v34 {
								v41 = int32(4)
								v43 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v35))))
								if v43&int32(254) == int32(2) {
									v52 = v41
								} else {
									v52 = base.B2i32(v43 == int32(18)) << (uint(v41) % 32)
								}
								if v43 == int32(1) {
									v55 = v41
								} else {
									v55 = v52
								}
								v66 = v55
							} else {
								v56 = int32(1)
								if v38 != 0 {
									v66 = int32(base.Ui32(v36)>>(uint(v56)%32)) - v56
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
									if v38 != 0 {
										v81 = v35
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
									*(*int32)(unsafe.Add(mBase, uint32(v10))) = int32(684446)
									F_errmsg(m, int32(246789), v10)
									mBase = m.M
									v106 = m.ExcPending
									if v106 != 0 {
										return int32(0)
									} else {
										F_errfinish(m, int32(495763), int32(1344), int32(398749))
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
				v28 = int32(0)
				F_parse_re_flags(m, v10+int32(8), v28)
				mBase = m.M
				v30 = m.ExcPending
				if v30 != 0 {
					return int32(0)
				} else {
					v31 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v10)+12)))
					if v31 != int32(1) {
						v34 = int32(1)
						v35 = v13 + v34
						v36 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v13))))
						v38 = v36 & v34
						if v36 == v34 {
							v41 = int32(4)
							v43 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v35))))
							if v43&int32(254) == int32(2) {
								v52 = v41
							} else {
								v52 = base.B2i32(v43 == int32(18)) << (uint(v41) % 32)
							}
							if v43 == int32(1) {
								v55 = v41
							} else {
								v55 = v52
							}
							v66 = v55
						} else {
							v56 = int32(1)
							if v38 != 0 {
								v66 = int32(base.Ui32(v36)>>(uint(v56)%32)) - v56
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
								if v38 != 0 {
									v81 = v35
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
								*(*int32)(unsafe.Add(mBase, uint32(v10))) = int32(684446)
								F_errmsg(m, int32(246789), v10)
								mBase = m.M
								v106 = m.ExcPending
								if v106 != 0 {
									return int32(0)
								} else {
									F_errfinish(m, int32(495763), int32(1344), int32(398749))
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
	var v40 int32
	_ = v40
	var v42 int32
	_ = v42
	var v51 int32
	_ = v51
	var v54 int32
	_ = v54
	var v55 int32
	_ = v55
	var v61 int32
	_ = v61
	var v67 int32
	_ = v67
	var v69 int32
	_ = v69
	var v72 int32
	_ = v72
	var v73 int32
	_ = v73
	var v74 int32
	_ = v74
	var v76 int32
	_ = v76
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
	var v88 int32
	_ = v88
	var v92 int32
	_ = v92
	var v97 int32
	_ = v97
	var v98 int32
	_ = v98
	var v106 int32
	_ = v106
	var v107 int32
	_ = v107
	var v110 int32
	_ = v110
	var v111 int32
	_ = v111
	var v114 int32
	_ = v114
	var v115 int32
	_ = v115
	var v118 int32
	_ = v118
	var v119 int32
	_ = v119
	var v121 int32
	_ = v121
	var v124 int32
	_ = v124
	var v125 int32
	_ = v125
	var v127 int32
	_ = v127
	var v128 int32
	_ = v128
	var v129 int32
	_ = v129
	var v138 int32
	_ = v138
	var v145 int32
	_ = v145
	var v146 int32
	_ = v146
	var v149 int32
	_ = v149
	var v151 int32
	_ = v151
	var v153 int32
	_ = v153
	var v154 int32
	_ = v154
	var v155 int32
	_ = v155
	var v160 int32
	_ = v160
	var v182 int32
	_ = v182
	var v185 int32
	_ = v185
	var v187 int32
	_ = v187
	var v190 int32
	_ = v190
	var v191 int32
	_ = v191
	var v193 int32
	_ = v193
	var v198 int32
	_ = v198
	var v202 int32
	_ = v202
	var v223 int32
	_ = v223
	var v225 int32
	_ = v225
	var v227 int32
	_ = v227
	var v245 int32
	_ = v245
	var v246 int32
	_ = v246
	var v248 int32
	_ = v248
	var v249 int32
	_ = v249
	var v251 int32
	_ = v251
	var v253 int32
	_ = v253
	var v256 int32
	_ = v256
	var v258 int32
	_ = v258
	var v262 int32
	_ = v262
	var v264 int32
	_ = v264
	var v266 int32
	_ = v266
	var v267 int32
	_ = v267
	var v269 int32
	_ = v269
	var v270 int32
	_ = v270
	var v271 int32
	_ = v271
	var v272 int32
	_ = v272
	var v274 int32
	_ = v274
	var v276 int32
	_ = v276
	var v280 int32
	_ = v280
	var v285 int32
	_ = v285
	var v287 int32
	_ = v287
	var v289 int32
	_ = v289
	var v293 int32
	_ = v293
	var v309 int32
	_ = v309
	var v315 int32
	_ = v315
	var v318 int32
	_ = v318
	var v320 int32
	_ = v320
	var v322 int32
	_ = v322
	var v323 int32
	_ = v323
	var v324 int32
	_ = v324
	var v325 int32
	_ = v325
	var v329 int32
	_ = v329
	var v338 int32
	_ = v338
	var v345 int32
	_ = v345
	var v348 int32
	_ = v348
	var v349 int32
	_ = v349
	var v351 int32
	_ = v351
	var v354 int32
	_ = v354
	var v358 int32
	_ = v358
	var v367 int32
	_ = v367
	var v374 int32
	_ = v374
	var v375 int32
	_ = v375
	var v381 int32
	_ = v381
	var v383 int32
	_ = v383
	var v384 int32
	_ = v384
	var v387 int64
	_ = v387
	var v393 int32
	_ = v393
	var v394 int32
	_ = v394
	var v395 int32
	_ = v395
	var v397 int32
	_ = v397
	var v398 int32
	_ = v398
	var v401 int32
	_ = v401
	var v402 int32
	_ = v402
	var v404 int32
	_ = v404
	var v409 int32
	_ = v409
	var v414 int32
	_ = v414
	var v417 int32
	_ = v417
	var v421 int32
	_ = v421
	var v426 int32
	_ = v426
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
	v28 = *(*int32)(unsafe.Add(mBase, _consts[357]))
	v29 = *(*int32)(unsafe.Add(mBase, uint32(v28)+4))
	v34 = *(*int32)(unsafe.Add(mBase, uint32(v29*int32(28))+uint32(_consts[356])))
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
	v69 = v67 + int32(1)
	v72 = F_palloc(m, v69<<(uint(int32(2))%32))
	mBase = m.M
	v73 = m.ExcPending
	if v73 != 0 {
		goto L1
	} else {
		goto L15
	}
L5:
	;
	v40 = int32(4)
	v42 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+1)))
	if v42&int32(254) == int32(2) {
		goto L8
	} else {
		goto L9
	}
L6:
	;
	goto L7
L7:
	;
	v55 = int32(1)
	if v37&v55 != 0 {
		v67 = int32(base.Ui32(v37)>>(uint(v55)%32)) - v55
		goto L4
	} else {
		goto L14
	}
L8:
	;
	v51 = v40
	goto L10
L9:
	;
	v51 = base.B2i32(v42 == int32(18)) << (uint(v40) % 32)
	goto L10
L10:
	;
	if v42 == int32(1) {
		goto L11
	} else {
		goto L12
	}
L11:
	;
	v54 = v40
	goto L13
L12:
	;
	v54 = v51
	goto L13
L13:
	;
	v67 = v54
	goto L4
L14:
	;
	v61 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v67 = int32(base.Ui32(v61)>>(uint(int32(2))%32)) - int32(4)
	goto L4
L15:
	;
	v74 = int32(1)
	v76 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0))))
	if v76&v74 != 0 {
		goto L16
	} else {
		goto L17
	}
L16:
	;
	v79 = v74
	goto L18
L17:
	;
	v79 = int32(4)
	goto L18
L18:
	;
	v81 = F_pg_mb2wchar_with_len(m, l0+v79, v72, v67)
	mBase = m.M
	v82 = m.ExcPending
	if v82 != 0 {
		goto L1
	} else {
		goto L19
	}
L19:
	;
	v83 = *(*int32)(unsafe.Add(mBase, uint32(l2)))
	if l5 != 0 {
		goto L20
	} else {
		goto L21
	}
L20:
	;
	v86 = v83
	goto L22
L21:
	;
	v86 = v83 | int32(16)
	goto L22
L22:
	;
	v87 = F_RE_compile_and_cache(m, l1, v86, l4)
	mBase = m.M
	v88 = m.ExcPending
	if v88 != 0 {
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
	v110 = F_palloc(m, v106<<(uint(int32(3))%32))
	mBase = m.M
	v111 = m.ExcPending
	if v111 != 0 {
		goto L1
	} else {
		goto L28
	}
L25:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v23)+8)) = int32(1)
	v106 = v36
	v107 = int32(0)
	goto L24
L26:
	;
	v92 = *(*int32)(unsafe.Add(mBase, _consts[1125]))
	if v92 == int32(0) {
		goto L25
	} else {
		goto L27
	}
L27:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v23)+8)) = v92
	v97 = *(*int32)(unsafe.Add(mBase, _consts[1125]))
	v98 = int32(1)
	v106 = v97 + v98
	v107 = v98
	goto L24
L28:
	;
	v114 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l2)+4)))
	if v114 != 0 {
		goto L29
	} else {
		goto L30
	}
L29:
	;
	v115 = int32(255)
	goto L31
L30:
	;
	v115 = int32(31)
	goto L31
L31:
	;
	v118 = F_palloc(m, v115<<(uint(int32(2))%32))
	mBase = m.M
	v119 = m.ExcPending
	if v119 != 0 {
		goto L1
	} else {
		goto L32
	}
L32:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v23)+12)) = v118
	v121 = int32(0)
	v124 = v115
	v125 = v121
	v127 = l3
	v128 = v121
	v129 = v121
	v138 = int32(0)
	goto L34
L33:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v414 = m.ExcPending
	if v414 != 0 {
		goto L1
	} else {
		goto L96
	}
L34:
	;
	v145 = F_RE_wchar_execute(m, v72, v81, v127, v106, v110)
	mBase = m.M
	v146 = m.ExcPending
	if v146 != 0 {
		goto L1
	} else {
		goto L37
	}
L35:
	;
	v374 = *(*int32)(unsafe.Add(mBase, uint32(v23)+12))
	v375 = int32(2)
	*(*int32)(unsafe.Add(mBase, uint32(v374+v354<<(uint(v375)%32)))) = v81
	if v375 <= v34 {
		goto L81
	} else {
		goto L82
	}
L36:
	;
	goto L35
L37:
	;
	if v145 == int32(0) {
		v354 = v125
		v358 = v129
		v367 = v138
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
	v345 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l2)+4)))
	if v345 != int32(1) {
		v354 = v325
		v358 = v329
		v367 = v338
		goto L36
	} else {
		goto L78
	}
L40:
	;
	v149 = *(*int32)(unsafe.Add(mBase, uint32(v110)))
	if v81 <= v149 {
		v324 = v124
		v325 = v125
		v329 = v129
		v338 = v138
		goto L39
	} else {
		goto L43
	}
L41:
	;
	goto L42
L42:
	;
	v153 = int32(1)
	v154 = v125 + v153
	v155 = *(*int32)(unsafe.Add(mBase, uint32(v23)+8))
	if v124 < v154+v155<<(uint(v153)%32) {
		goto L45
	} else {
		goto L46
	}
L43:
	;
	v151 = *(*int32)(unsafe.Add(mBase, uint32(v110)+4))
	if v151 <= v128 {
		v324 = v124
		v325 = v125
		v329 = v129
		v338 = v138
		goto L39
	} else {
		goto L44
	}
L44:
	;
	goto L42
L45:
	;
	v160 = v124
	goto L48
L46:
	;
	v198 = v124
	v202 = v155
	goto L47
L47:
	;
	if v107 != 0 {
		goto L54
	} else {
		goto L55
	}
L48:
	;
	v182 = v160 << (uint(int32(1)) % 32)
	if base.Ui32(int32(268435456)) <= base.Ui32(v182) {
		goto L33
	} else {
		goto L50
	}
L49:
	;
	v198 = v187
	v202 = v193
	goto L47
L50:
	;
	v185 = *(*int32)(unsafe.Add(mBase, uint32(v23)+12))
	v187 = v182 | int32(1)
	v190 = F_repalloc(m, v185, v187<<(uint(int32(2))%32))
	mBase = m.M
	v191 = m.ExcPending
	if v191 != 0 {
		goto L1
	} else {
		goto L51
	}
L51:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v23)+12)) = v190
	v193 = *(*int32)(unsafe.Add(mBase, uint32(v23)+8))
	if v187 < v154+v193<<(uint(int32(1))%32) {
		v160 = v187
		goto L48
	} else {
		goto L52
	}
L52:
	;
	goto L49
L53:
	;
	v309 = *(*int32)(unsafe.Add(mBase, uint32(v23)+4))
	*(*int32)(unsafe.Add(mBase, uint32(v23)+4)) = v309 + int32(1)
	if l7 == int32(0) {
		v322 = v293
		goto L72
	} else {
		goto L73
	}
L54:
	;
	if v202 <= int32(0) {
		v289 = v125
		v293 = v129
		goto L53
	} else {
		goto L57
	}
L55:
	;
	goto L56
L56:
	;
	v269 = *(*int32)(unsafe.Add(mBase, uint32(v110)+4))
	v270 = int32(2)
	v271 = v125 << (uint(v270) % 32)
	v272 = *(*int32)(unsafe.Add(mBase, uint32(v23)+12))
	v274 = *(*int32)(unsafe.Add(mBase, uint32(v110)))
	*(*int32)(unsafe.Add(mBase, uint32(v271+v272))) = v274
	v276 = *(*int32)(unsafe.Add(mBase, uint32(v23)+12))
	*(*int32)(unsafe.Add(mBase, uint32(v276+v271)+4)) = v269
	v280 = v125 + v270
	if v274 < int32(0) {
		v289 = v280
		v293 = v129
		goto L53
	} else {
		goto L67
	}
L57:
	;
	v223 = v125
	v225 = int32(1)
	v227 = v129
	goto L58
L58:
	;
	v245 = v110 + v225<<(uint(int32(3))%32)
	v246 = *(*int32)(unsafe.Add(mBase, uint32(v245)+4))
	v248 = v223 << (uint(int32(2)) % 32)
	v249 = *(*int32)(unsafe.Add(mBase, uint32(v23)+12))
	v251 = *(*int32)(unsafe.Add(mBase, uint32(v245)))
	*(*int32)(unsafe.Add(mBase, uint32(v248+v249))) = v251
	v253 = *(*int32)(unsafe.Add(mBase, uint32(v23)+12))
	*(*int32)(unsafe.Add(mBase, uint32(v253+v248)+4)) = v246
	v256 = v246 - v251
	if v227 < v256 {
		goto L60
	} else {
		goto L61
	}
L59:
	;
	v289 = v264
	v293 = v262
	goto L53
L60:
	;
	v258 = v256
	goto L62
L61:
	;
	v258 = v227
	goto L62
L62:
	;
	if int32(0) <= v246|v251 {
		goto L63
	} else {
		goto L64
	}
L63:
	;
	v262 = v258
	goto L65
L64:
	;
	v262 = v227
	goto L65
L65:
	;
	v264 = v223 + int32(2)
	v266 = v225 + int32(1)
	v267 = *(*int32)(unsafe.Add(mBase, uint32(v23)+8))
	if v266 <= v267 {
		v223 = v264
		v225 = v266
		v227 = v262
		goto L58
	} else {
		goto L66
	}
L66:
	;
	goto L59
L67:
	;
	if v269 < int32(0) {
		v289 = v280
		v293 = v129
		goto L53
	} else {
		goto L68
	}
L68:
	;
	v285 = v269 - v274
	if v129 < v285 {
		goto L69
	} else {
		goto L70
	}
L69:
	;
	v287 = v285
	goto L71
L70:
	;
	v287 = v129
	goto L71
L71:
	;
	v289 = v280
	v293 = v287
	goto L53
L72:
	;
	v323 = *(*int32)(unsafe.Add(mBase, uint32(v110)+4))
	v324 = v198
	v325 = v289
	v329 = v322
	v338 = v323
	goto L39
L73:
	;
	v315 = *(*int32)(unsafe.Add(mBase, uint32(v110)))
	if v315 < int32(0) {
		v322 = v293
		goto L72
	} else {
		goto L74
	}
L74:
	;
	v318 = v315 - v138
	if v293 < v318 {
		goto L75
	} else {
		goto L76
	}
L75:
	;
	v320 = v318
	goto L77
L76:
	;
	v320 = v293
	goto L77
L77:
	;
	v322 = v320
	goto L72
L78:
	;
	v348 = *(*int32)(unsafe.Add(mBase, uint32(v110)+4))
	v349 = *(*int32)(unsafe.Add(mBase, uint32(v110)))
	v351 = v348 + base.B2i32(v348 == v349)
	if v351 <= v81 {
		v124 = v324
		v125 = v325
		v127 = v351
		v128 = v348
		v129 = v329
		v138 = v338
		goto L34
	} else {
		goto L79
	}
L79:
	;
	v354 = v325
	v358 = v329
	v367 = v338
	goto L36
L80:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v23)+36)) = v401
	*(*int32)(unsafe.Add(mBase, uint32(v23)+32)) = v404
	*(*int32)(unsafe.Add(mBase, uint32(v23)+28)) = v402
	F_pfree(m, v110)
	mBase = m.M
	v409 = m.ExcPending
	if v409 != 0 {
		goto L1
	} else {
		goto L95
	}
L81:
	;
	v381 = v81 - v367
	if v358 < v381 {
		goto L84
	} else {
		goto L85
	}
L82:
	;
	goto L83
L83:
	;
	F_pfree(m, v72)
	mBase = m.M
	v397 = m.ExcPending
	if v397 != 0 {
		goto L1
	} else {
		goto L94
	}
L84:
	;
	v383 = v381
	goto L86
L85:
	;
	v383 = v358
	goto L86
L86:
	;
	if l7 != 0 {
		goto L87
	} else {
		goto L88
	}
L87:
	;
	v384 = v383
	goto L89
L88:
	;
	v384 = v358
	goto L89
L89:
	;
	v387 = base.I64_extend_i32_u(v384) * base.I64_extend_i32_u(v34)
	if base.I64_extend_i32_s(v67) < v387 {
		goto L90
	} else {
		goto L91
	}
L90:
	;
	v393 = v69
	goto L92
L91:
	;
	v393 = base.I32_wrap_i64(v387) + int32(1)
	goto L92
L92:
	;
	v394 = F_palloc(m, v393)
	mBase = m.M
	v395 = m.ExcPending
	if v395 != 0 {
		goto L1
	} else {
		goto L93
	}
L93:
	;
	v401 = v393
	v402 = v72
	v404 = v394
	goto L80
L94:
	;
	v398 = int32(0)
	v401 = v398
	v402 = v398
	v404 = v398
	goto L80
L95:
	;
	return v23
L96:
	;
	F_errcode(m, int32(261))
	mBase = m.M
	v417 = m.ExcPending
	if v417 != 0 {
		goto L1
	} else {
		goto L97
	}
L97:
	;
	F_errmsg(m, int32(168755), int32(0))
	mBase = m.M
	v421 = m.ExcPending
	if v421 != 0 {
		goto L1
	} else {
		goto L98
	}
L98:
	;
	F_errfinish(m, int32(495763), int32(1572), int32(168704))
	mBase = m.M
	v426 = m.ExcPending
	if v426 != 0 {
		goto L1
	} else {
		goto L99
	}
L99:
	;
	base.Wasm_trap_unreachable()
	for {
	}
}
