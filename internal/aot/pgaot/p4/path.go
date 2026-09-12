package p4

import (
	base "github.com/shibukawa/pgmem/internal/aot/pgaot/base"
	"math"
	"unsafe"
)

func F_PathNameDeleteTemporaryFile(m *base.Module, l0 int32, l1 int32) int32 {
	mBase := m.M
	_ = mBase
	var v7 int32
	_ = v7
	var v9 int32
	_ = v9
	var v15 int32
	_ = v15
	var v17 int32
	_ = v17
	var v21 int32
	_ = v21
	var v24 int32
	_ = v24
	var v29 int32
	_ = v29
	var v33 int32
	_ = v33
	var v36 int32
	_ = v36
	var v38 int32
	_ = v38
	var v41 int32
	_ = v41
	var v45 int32
	_ = v45
	var v49 int32
	_ = v49
	var v54 int32
	_ = v54
	var v57 int32
	_ = v57
	var v60 int32
	_ = v60
	var v61 int32
	_ = v61
	var v65 int32
	_ = v65
	var v71 int32
	_ = v71
	var v76 int32
	_ = v76
	var v78 int64
	_ = v78
	var v79 int32
	_ = v79
	var v81 int32
	_ = v81
	var v83 int64
	_ = v83
	var v84 int32
	_ = v84
	var v86 int64
	_ = v86
	var v92 int32
	_ = v92
	var v93 int32
	_ = v93
	var v102 int32
	_ = v102
	var v107 int32
	_ = v107
	var v109 int32
	_ = v109
	v7 = m.G0
	v9 = v7 - int32(144)
	m.G0 = v9
	v15 = F___fstatat(m, int32(-100), l0, v9+int32(48), int32(0))
	mBase = m.M
	if v15 != 0 {
		v17 = *(*int32)(unsafe.Add(mBase, _consts[140]))
		if v17 == int32(44) {
			v109 = int32(0)
			m.G0 = v9 + int32(144)
			return v109
		} else {
			v21 = F_unlink(m, l0)
			mBase = m.M
			if v21 < int32(0) {
				v29 = *(*int32)(unsafe.Add(mBase, _consts[140]))
				if v29 == int32(44) {
					v109 = int32(0)
					m.G0 = v9 + int32(144)
					return v109
				} else {
					v33 = int32(0)
					if l1 != 0 {
						v36 = int32(21)
					} else {
						v36 = int32(15)
					}
					v38 = F_errstart(m, v36, int32(0))
					mBase = m.M
					v41 = m.ExcPending
					if v41 != 0 {
						return int32(0)
					} else {
						if v38 == int32(0) {
							v109 = v33
							m.G0 = v9 + int32(144)
							return v109
						} else {
							F_errcode_for_file_access(m)
							mBase = m.M
							v45 = m.ExcPending
							if v45 != 0 {
								return int32(0)
							} else {
								*(*int32)(unsafe.Add(mBase, uint32(v9))) = l0
								F_errmsg(m, int32(297379), v9)
								mBase = m.M
								v49 = m.ExcPending
								if v49 != 0 {
									return int32(0)
								} else {
									F_errfinish(m, int32(499457), int32(1966), int32(389315))
									mBase = m.M
									v54 = m.ExcPending
									if v54 != 0 {
										return int32(0)
									} else {
										v109 = v33
										m.G0 = v9 + int32(144)
										return v109
									}
								}
							}
						}
					}
				}
			} else {
				if v17 != 0 {
					*(*int32)(unsafe.Add(mBase, _consts[140])) = v17
					v57 = int32(1)
					v60 = F_errstart(m, int32(15), int32(0))
					mBase = m.M
					v61 = m.ExcPending
					if v61 != 0 {
						return int32(0)
					} else {
						if v60 == int32(0) {
							v109 = v57
							m.G0 = v9 + int32(144)
							return v109
						} else {
							F_errcode_for_file_access(m)
							mBase = m.M
							v65 = m.ExcPending
							if v65 != 0 {
								return int32(0)
							} else {
								*(*int32)(unsafe.Add(mBase, uint32(v9)+32)) = l0
								F_errmsg(m, int32(297574), v9+int32(32))
								mBase = m.M
								v71 = m.ExcPending
								if v71 != 0 {
									return int32(0)
								} else {
									F_errfinish(m, int32(499457), int32(1977), int32(389315))
									mBase = m.M
									v76 = m.ExcPending
									if v76 != 0 {
										return int32(0)
									} else {
										v109 = v57
										m.G0 = v9 + int32(144)
										return v109
									}
								}
							}
						}
					}
				} else {
					v78 = *(*int64)(unsafe.Add(mBase, uint32(v9)+72))
					v79 = base.I32_wrap_i64(v78)
					F_pgstat_report_tempfile(m, v79)
					mBase = m.M
					v81 = m.ExcPending
					if v81 != 0 {
						return int32(0)
					} else {
						v83 = base.I64_div_s(v78, int64(1024))
						v84 = int32(1)
						v86 = int64(*(*int32)(unsafe.Add(mBase, _consts[948])))
						if v86 < int64(0) {
							v109 = v84
							m.G0 = v9 + int32(144)
							return v109
						} else {
							if v83 < v86 {
								v109 = v84
								m.G0 = v9 + int32(144)
								return v109
							} else {
								v92 = F_errstart(m, int32(15), int32(0))
								mBase = m.M
								v93 = m.ExcPending
								if v93 != 0 {
									return int32(0)
								} else {
									if v92 == int32(0) {
										v109 = v84
										m.G0 = v9 + int32(144)
										return v109
									} else {
										*(*int32)(unsafe.Add(mBase, uint32(v9)+20)) = v79
										*(*int32)(unsafe.Add(mBase, uint32(v9)+16)) = l0
										F_errmsg(m, int32(38321), v9+int32(16))
										mBase = m.M
										v102 = m.ExcPending
										if v102 != 0 {
											return int32(0)
										} else {
											F_errfinish(m, int32(499457), int32(1546), int32(405209))
											mBase = m.M
											v107 = m.ExcPending
											if v107 != 0 {
												return int32(0)
											} else {
												v109 = v84
												m.G0 = v9 + int32(144)
												return v109
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
	} else {
		v24 = F_unlink(m, l0)
		mBase = m.M
		if int32(0) <= v24 {
			v78 = *(*int64)(unsafe.Add(mBase, uint32(v9)+72))
			v79 = base.I32_wrap_i64(v78)
			F_pgstat_report_tempfile(m, v79)
			mBase = m.M
			v81 = m.ExcPending
			if v81 != 0 {
				return int32(0)
			} else {
				v83 = base.I64_div_s(v78, int64(1024))
				v84 = int32(1)
				v86 = int64(*(*int32)(unsafe.Add(mBase, _consts[948])))
				if v86 < int64(0) {
					v109 = v84
					m.G0 = v9 + int32(144)
					return v109
				} else {
					if v83 < v86 {
						v109 = v84
						m.G0 = v9 + int32(144)
						return v109
					} else {
						v92 = F_errstart(m, int32(15), int32(0))
						mBase = m.M
						v93 = m.ExcPending
						if v93 != 0 {
							return int32(0)
						} else {
							if v92 == int32(0) {
								v109 = v84
								m.G0 = v9 + int32(144)
								return v109
							} else {
								*(*int32)(unsafe.Add(mBase, uint32(v9)+20)) = v79
								*(*int32)(unsafe.Add(mBase, uint32(v9)+16)) = l0
								F_errmsg(m, int32(38321), v9+int32(16))
								mBase = m.M
								v102 = m.ExcPending
								if v102 != 0 {
									return int32(0)
								} else {
									F_errfinish(m, int32(499457), int32(1546), int32(405209))
									mBase = m.M
									v107 = m.ExcPending
									if v107 != 0 {
										return int32(0)
									} else {
										v109 = v84
										m.G0 = v9 + int32(144)
										return v109
									}
								}
							}
						}
					}
				}
			}
		} else {
			v29 = *(*int32)(unsafe.Add(mBase, _consts[140]))
			if v29 == int32(44) {
				v109 = int32(0)
				m.G0 = v9 + int32(144)
				return v109
			} else {
				v33 = int32(0)
				if l1 != 0 {
					v36 = int32(21)
				} else {
					v36 = int32(15)
				}
				v38 = F_errstart(m, v36, int32(0))
				mBase = m.M
				v41 = m.ExcPending
				if v41 != 0 {
					return int32(0)
				} else {
					if v38 == int32(0) {
						v109 = v33
						m.G0 = v9 + int32(144)
						return v109
					} else {
						F_errcode_for_file_access(m)
						mBase = m.M
						v45 = m.ExcPending
						if v45 != 0 {
							return int32(0)
						} else {
							*(*int32)(unsafe.Add(mBase, uint32(v9))) = l0
							F_errmsg(m, int32(297379), v9)
							mBase = m.M
							v49 = m.ExcPending
							if v49 != 0 {
								return int32(0)
							} else {
								F_errfinish(m, int32(499457), int32(1966), int32(389315))
								mBase = m.M
								v54 = m.ExcPending
								if v54 != 0 {
									return int32(0)
								} else {
									v109 = v33
									m.G0 = v9 + int32(144)
									return v109
								}
							}
						}
					}
				}
			}
		}
	}
}
func F_path_add_pt(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v10 int32
	_ = v10
	var v11 int32
	_ = v11
	var v14 int32
	_ = v14
	var v15 int32
	_ = v15
	var v18 int32
	_ = v18
	var v23 int32
	_ = v23
	var v32 int32
	_ = v32
	var v33 float64
	_ = v33
	var v34 float64
	_ = v34
	var v35 float64
	_ = v35
	var v45 float64
	_ = v45
	var v46 float64
	_ = v46
	var v47 float64
	_ = v47
	var v60 int32
	_ = v60
	var v61 int32
	_ = v61
	var v77 int32
	_ = v77
	v10 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v11 = F_pg_detoast_datum_copy(m, v10)
	mBase = m.M
	v14 = m.ExcPending
	if v14 != 0 {
		goto L2
	} else {
		goto L3
	}
L1:
	;
	F_float_overflow_error(m)
	mBase = m.M
	v77 = m.ExcPending
	if v77 != 0 {
		goto L2
	} else {
		goto L18
	}
L2:
	;
	return int32(0)
L3:
	;
	v15 = *(*int32)(unsafe.Add(mBase, uint32(v11)+4))
	if int32(0) < v15 {
		goto L4
	} else {
		goto L5
	}
L4:
	;
	v18 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
	v23 = int32(0)
	goto L7
L5:
	;
	goto L6
L6:
	;
	return v11
L7:
	;
	v32 = v11 + int32(16) + v23<<(uint(int32(4))%32)
	v33 = *(*float64)(unsafe.Add(mBase, uint32(v32)))
	v34 = *(*float64)(unsafe.Add(mBase, uint32(v18)))
	v35 = base.F64_add(v33, v34)
	if base.F64_ne(base.F64_abs(v35), math.Float64frombits(uint64(0x7ff0000000000000))) != 0 {
		goto L9
	} else {
		goto L10
	}
L8:
	;
	goto L6
L9:
	;
	v45 = *(*float64)(unsafe.Add(mBase, uint32(v32)+8))
	v46 = *(*float64)(unsafe.Add(mBase, uint32(v18)+8))
	v47 = base.F64_add(v45, v46)
	if base.F64_ne(base.F64_abs(v47), math.Float64frombits(uint64(0x7ff0000000000000))) != 0 {
		goto L13
	} else {
		goto L14
	}
L10:
	;
	if base.F64_eq(base.F64_abs(v33), math.Float64frombits(uint64(0x7ff0000000000000))) != 0 {
		goto L9
	} else {
		goto L11
	}
L11:
	;
	if base.F64_ne(base.F64_abs(v34), math.Float64frombits(uint64(0x7ff0000000000000))) != 0 {
		goto L1
	} else {
		goto L12
	}
L12:
	;
	goto L9
L13:
	;
	*(*float64)(unsafe.Add(mBase, uint32(v32)+8)) = v47
	*(*float64)(unsafe.Add(mBase, uint32(v32))) = v35
	v60 = v23 + int32(1)
	v61 = *(*int32)(unsafe.Add(mBase, uint32(v11)+4))
	if v60 < v61 {
		v23 = v60
		goto L7
	} else {
		goto L17
	}
L14:
	;
	if base.F64_eq(base.F64_abs(v45), math.Float64frombits(uint64(0x7ff0000000000000))) != 0 {
		goto L13
	} else {
		goto L15
	}
L15:
	;
	if base.F64_ne(base.F64_abs(v46), math.Float64frombits(uint64(0x7ff0000000000000))) != 0 {
		goto L1
	} else {
		goto L16
	}
L16:
	;
	goto L13
L17:
	;
	goto L8
L18:
	;
	base.Wasm_trap_unreachable()
	for {
	}
}
func F_path_inter(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v2 int32
	_ = v2
	var v23 int32
	_ = v23
	var v25 int32
	_ = v25
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
	var v35 int32
	_ = v35
	var v36 int32
	_ = v36
	var v37 float64
	_ = v37
	var v38 float64
	_ = v38
	var v39 int32
	_ = v39
	var v43 int32
	_ = v43
	var v57 float64
	_ = v57
	var v58 float64
	_ = v58
	var v61 float64
	_ = v61
	var v62 float64
	_ = v62
	var v67 int32
	_ = v67
	var v68 float64
	_ = v68
	var v79 float64
	_ = v79
	var v81 float64
	_ = v81
	var v82 float64
	_ = v82
	var v83 float64
	_ = v83
	var v94 float64
	_ = v94
	var v96 float64
	_ = v96
	var v97 float64
	_ = v97
	var v104 float64
	_ = v104
	var v110 float64
	_ = v110
	var v111 float64
	_ = v111
	var v118 float64
	_ = v118
	var v124 float64
	_ = v124
	var v125 float64
	_ = v125
	var v127 int32
	_ = v127
	var v143 float64
	_ = v143
	var v144 float64
	_ = v144
	var v147 float64
	_ = v147
	var v148 float64
	_ = v148
	var v152 int32
	_ = v152
	var v153 float64
	_ = v153
	var v154 float64
	_ = v154
	var v155 int32
	_ = v155
	var v159 int32
	_ = v159
	var v175 float64
	_ = v175
	var v176 float64
	_ = v176
	var v179 float64
	_ = v179
	var v180 float64
	_ = v180
	var v183 int32
	_ = v183
	var v184 float64
	_ = v184
	var v195 float64
	_ = v195
	var v197 float64
	_ = v197
	var v198 float64
	_ = v198
	var v199 float64
	_ = v199
	var v210 float64
	_ = v210
	var v212 float64
	_ = v212
	var v213 float64
	_ = v213
	var v220 float64
	_ = v220
	var v226 float64
	_ = v226
	var v227 float64
	_ = v227
	var v234 float64
	_ = v234
	var v240 float64
	_ = v240
	var v241 float64
	_ = v241
	var v243 int32
	_ = v243
	var v261 float64
	_ = v261
	var v262 float64
	_ = v262
	var v265 float64
	_ = v265
	var v266 float64
	_ = v266
	var v291 int32
	_ = v291
	var v292 int32
	_ = v292
	var v297 int32
	_ = v297
	var v313 int32
	_ = v313
	var v316 int32
	_ = v316
	var v319 int32
	_ = v319
	var v321 int32
	_ = v321
	var v326 int32
	_ = v326
	var v327 int32
	_ = v327
	var v328 float64
	_ = v328
	var v330 float64
	_ = v330
	var v332 float64
	_ = v332
	var v334 float64
	_ = v334
	var v340 int32
	_ = v340
	var v341 float64
	_ = v341
	var v343 float64
	_ = v343
	var v345 float64
	_ = v345
	var v347 float64
	_ = v347
	var v352 int32
	_ = v352
	var v353 int32
	_ = v353
	var v355 int32
	_ = v355
	var v357 int32
	_ = v357
	var v358 int32
	_ = v358
	var v362 int32
	_ = v362
	var v384 float64
	_ = v384
	var v386 float64
	_ = v386
	var v388 float64
	_ = v388
	var v390 float64
	_ = v390
	var v394 int32
	_ = v394
	var v396 int32
	_ = v396
	var v397 float64
	_ = v397
	var v399 float64
	_ = v399
	var v401 float64
	_ = v401
	var v403 float64
	_ = v403
	var v408 int32
	_ = v408
	var v409 int32
	_ = v409
	var v411 int32
	_ = v411
	var v412 int32
	_ = v412
	var v416 int32
	_ = v416
	var v436 int32
	_ = v436
	var v439 int32
	_ = v439
	var v440 int32
	_ = v440
	var v460 int32
	_ = v460
	var v470 int32
	_ = v470
	v2 = int32(0)
	v23 = m.G0
	v25 = v23 + int32(-64)
	m.G0 = v25
	v27 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v28 = F_pg_detoast_datum(m, v27)
	mBase = m.M
	v31 = m.ExcPending
	if v31 != 0 {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	return int32(0)
L2:
	;
	v33 = v28 + int32(16)
	v34 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
	v35 = F_pg_detoast_datum(m, v34)
	mBase = m.M
	v36 = m.ExcPending
	if v36 != 0 {
		goto L1
	} else {
		goto L3
	}
L3:
	;
	v37 = *(*float64)(unsafe.Add(mBase, uint32(v28)+24))
	v38 = *(*float64)(unsafe.Add(mBase, uint32(v28)+16))
	v39 = *(*int32)(unsafe.Add(mBase, uint32(v28)+4))
	if v39 < int32(2) {
		goto L5
	} else {
		goto L6
	}
L4:
	;
	v152 = v35 + int32(16)
	v153 = *(*float64)(unsafe.Add(mBase, uint32(v35)+24))
	v154 = *(*float64)(unsafe.Add(mBase, uint32(v35)+16))
	v155 = *(*int32)(unsafe.Add(mBase, uint32(v35)+4))
	if v155 < int32(2) {
		goto L48
	} else {
		goto L49
	}
L5:
	;
	v143 = v37
	v144 = v38
	v147 = v37
	v148 = v38
	goto L4
L6:
	;
	goto L7
L7:
	;
	v43 = int32(1)
	v57 = v37
	v58 = v38
	v61 = v37
	v62 = v38
	goto L8
L8:
	;
	v67 = v33 + v43<<(uint(int32(4))%32)
	v68 = *(*float64)(unsafe.Add(mBase, uint32(v67)))
	if base.Ui64(base.I64_reinterpret_f64(v58)&int64(9223372036854775807)) <= base.Ui64(int64(9218868437227405312)) {
		goto L10
	} else {
		goto L11
	}
L9:
	;
	v143 = v97
	v144 = v82
	v147 = v125
	v148 = v111
	goto L4
L10:
	;
	if base.Ui64(int64(9218868437227405312)) < base.Ui64(base.I64_reinterpret_f64(v68)&int64(9223372036854775807)) {
		goto L13
	} else {
		goto L14
	}
L11:
	;
	v82 = v58
	goto L12
L12:
	;
	v83 = *(*float64)(unsafe.Add(mBase, uint32(v67)+8))
	if base.Ui64(base.I64_reinterpret_f64(v57)&int64(9223372036854775807)) <= base.Ui64(int64(9218868437227405312)) {
		goto L19
	} else {
		goto L20
	}
L13:
	;
	v79 = v68
	goto L15
L14:
	;
	v79 = v58
	goto L15
L15:
	;
	if base.F64_gt(v68, v58) != 0 {
		goto L16
	} else {
		goto L17
	}
L16:
	;
	v81 = v68
	goto L18
L17:
	;
	v81 = v79
	goto L18
L18:
	;
	v82 = v81
	goto L12
L19:
	;
	if base.Ui64(int64(9218868437227405312)) < base.Ui64(base.I64_reinterpret_f64(v83)&int64(9223372036854775807)) {
		goto L22
	} else {
		goto L23
	}
L20:
	;
	v97 = v57
	goto L21
L21:
	;
	if base.Ui64(base.I64_reinterpret_f64(v68)&int64(9223372036854775807)) <= base.Ui64(int64(9218868437227405312)) {
		goto L28
	} else {
		goto L29
	}
L22:
	;
	v94 = v83
	goto L24
L23:
	;
	v94 = v57
	goto L24
L24:
	;
	if base.F64_gt(v83, v57) != 0 {
		goto L25
	} else {
		goto L26
	}
L25:
	;
	v96 = v83
	goto L27
L26:
	;
	v96 = v94
	goto L27
L27:
	;
	v97 = v96
	goto L21
L28:
	;
	if base.F64_lt(v68, v62) != 0 {
		goto L31
	} else {
		goto L32
	}
L29:
	;
	v111 = v62
	goto L30
L30:
	;
	if base.Ui64(base.I64_reinterpret_f64(v83)&int64(9223372036854775807)) <= base.Ui64(int64(9218868437227405312)) {
		goto L37
	} else {
		goto L38
	}
L31:
	;
	v104 = v68
	goto L33
L32:
	;
	v104 = v62
	goto L33
L33:
	;
	if base.Ui64(int64(9218868437227405312)) < base.Ui64(base.I64_reinterpret_f64(v62)&int64(9223372036854775807)) {
		goto L34
	} else {
		goto L35
	}
L34:
	;
	v110 = v68
	goto L36
L35:
	;
	v110 = v104
	goto L36
L36:
	;
	v111 = v110
	goto L30
L37:
	;
	if base.F64_lt(v83, v61) != 0 {
		goto L40
	} else {
		goto L41
	}
L38:
	;
	v125 = v61
	goto L39
L39:
	;
	v127 = v43 + int32(1)
	if v127 != v39 {
		v43 = v127
		v57 = v97
		v58 = v82
		v61 = v125
		v62 = v111
		goto L8
	} else {
		goto L46
	}
L40:
	;
	v118 = v83
	goto L42
L41:
	;
	v118 = v61
	goto L42
L42:
	;
	if base.Ui64(int64(9218868437227405312)) < base.Ui64(base.I64_reinterpret_f64(v61)&int64(9223372036854775807)) {
		goto L43
	} else {
		goto L44
	}
L43:
	;
	v124 = v83
	goto L45
L44:
	;
	v124 = v118
	goto L45
L45:
	;
	v125 = v124
	goto L39
L46:
	;
	goto L9
L47:
	;
	if base.F64_le(v148, base.F64_add(v262, float64(1e-06))) == int32(0) {
		v470 = v2
		goto L90
	} else {
		goto L91
	}
L48:
	;
	v261 = v153
	v262 = v154
	v265 = v153
	v266 = v154
	goto L47
L49:
	;
	goto L50
L50:
	;
	v159 = int32(1)
	v175 = v153
	v176 = v154
	v179 = v153
	v180 = v154
	goto L51
L51:
	;
	v183 = v152 + v159<<(uint(int32(4))%32)
	v184 = *(*float64)(unsafe.Add(mBase, uint32(v183)))
	if base.Ui64(base.I64_reinterpret_f64(v176)&int64(9223372036854775807)) <= base.Ui64(int64(9218868437227405312)) {
		goto L53
	} else {
		goto L54
	}
L52:
	;
	v261 = v213
	v262 = v198
	v265 = v241
	v266 = v227
	goto L47
L53:
	;
	if base.Ui64(int64(9218868437227405312)) < base.Ui64(base.I64_reinterpret_f64(v184)&int64(9223372036854775807)) {
		goto L56
	} else {
		goto L57
	}
L54:
	;
	v198 = v176
	goto L55
L55:
	;
	v199 = *(*float64)(unsafe.Add(mBase, uint32(v183)+8))
	if base.Ui64(base.I64_reinterpret_f64(v175)&int64(9223372036854775807)) <= base.Ui64(int64(9218868437227405312)) {
		goto L62
	} else {
		goto L63
	}
L56:
	;
	v195 = v184
	goto L58
L57:
	;
	v195 = v176
	goto L58
L58:
	;
	if base.F64_gt(v184, v176) != 0 {
		goto L59
	} else {
		goto L60
	}
L59:
	;
	v197 = v184
	goto L61
L60:
	;
	v197 = v195
	goto L61
L61:
	;
	v198 = v197
	goto L55
L62:
	;
	if base.Ui64(int64(9218868437227405312)) < base.Ui64(base.I64_reinterpret_f64(v199)&int64(9223372036854775807)) {
		goto L65
	} else {
		goto L66
	}
L63:
	;
	v213 = v175
	goto L64
L64:
	;
	if base.Ui64(base.I64_reinterpret_f64(v184)&int64(9223372036854775807)) <= base.Ui64(int64(9218868437227405312)) {
		goto L71
	} else {
		goto L72
	}
L65:
	;
	v210 = v199
	goto L67
L66:
	;
	v210 = v175
	goto L67
L67:
	;
	if base.F64_gt(v199, v175) != 0 {
		goto L68
	} else {
		goto L69
	}
L68:
	;
	v212 = v199
	goto L70
L69:
	;
	v212 = v210
	goto L70
L70:
	;
	v213 = v212
	goto L64
L71:
	;
	if base.F64_lt(v184, v180) != 0 {
		goto L74
	} else {
		goto L75
	}
L72:
	;
	v227 = v180
	goto L73
L73:
	;
	if base.Ui64(base.I64_reinterpret_f64(v199)&int64(9223372036854775807)) <= base.Ui64(int64(9218868437227405312)) {
		goto L80
	} else {
		goto L81
	}
L74:
	;
	v220 = v184
	goto L76
L75:
	;
	v220 = v180
	goto L76
L76:
	;
	if base.Ui64(int64(9218868437227405312)) < base.Ui64(base.I64_reinterpret_f64(v180)&int64(9223372036854775807)) {
		goto L77
	} else {
		goto L78
	}
L77:
	;
	v226 = v184
	goto L79
L78:
	;
	v226 = v220
	goto L79
L79:
	;
	v227 = v226
	goto L73
L80:
	;
	if base.F64_lt(v199, v179) != 0 {
		goto L83
	} else {
		goto L84
	}
L81:
	;
	v241 = v179
	goto L82
L82:
	;
	v243 = v159 + int32(1)
	if v243 != v155 {
		v159 = v243
		v175 = v213
		v176 = v198
		v179 = v241
		v180 = v227
		goto L51
	} else {
		goto L89
	}
L83:
	;
	v234 = v199
	goto L85
L84:
	;
	v234 = v179
	goto L85
L85:
	;
	if base.Ui64(int64(9218868437227405312)) < base.Ui64(base.I64_reinterpret_f64(v179)&int64(9223372036854775807)) {
		goto L86
	} else {
		goto L87
	}
L86:
	;
	v240 = v199
	goto L88
L87:
	;
	v240 = v234
	goto L88
L88:
	;
	v241 = v240
	goto L82
L89:
	;
	goto L52
L90:
	;
	m.G0 = v25 - int32(-64)
	return v470
L91:
	;
	if base.F64_le(v266, base.F64_add(v144, float64(1e-06))) == int32(0) {
		v470 = v2
		goto L90
	} else {
		goto L92
	}
L92:
	;
	if base.F64_le(v147, base.F64_add(v261, float64(1e-06))) == int32(0) {
		v470 = v2
		goto L90
	} else {
		goto L93
	}
L93:
	;
	if base.F64_le(v265, base.F64_add(v143, float64(1e-06))) == int32(0) {
		v470 = v2
		goto L90
	} else {
		goto L94
	}
L94:
	;
	if v39 <= int32(0) {
		v470 = v2
		goto L90
	} else {
		goto L95
	}
L95:
	;
	v291 = v155
	v292 = v39
	v297 = v2
	goto L96
L96:
	;
	if v297 == int32(0) {
		goto L99
	} else {
		goto L100
	}
L97:
	;
	v470 = int32(0)
	goto L90
L98:
	;
	v460 = v297 + int32(1)
	if v460 < v440 {
		v291 = v439
		v292 = v440
		v297 = v460
		goto L96
	} else {
		goto L119
	}
L99:
	;
	v313 = *(*int32)(unsafe.Add(mBase, uint32(v28)+8))
	if v313 == int32(0) {
		v439 = v291
		v440 = v292
		goto L98
	} else {
		goto L102
	}
L100:
	;
	v316 = v297
	goto L101
L101:
	;
	if v291 <= int32(0) {
		v439 = v291
		v440 = v292
		goto L98
	} else {
		goto L103
	}
L102:
	;
	v316 = v292
	goto L101
L103:
	;
	v319 = int32(4)
	v321 = v33 + v297<<(uint(v319)%32)
	v326 = v316<<(uint(v319)%32) + v33 - int32(16)
	v327 = *(*int32)(unsafe.Add(mBase, uint32(v35)+8))
	if v327 != 0 {
		goto L104
	} else {
		goto L105
	}
L104:
	;
	v328 = *(*float64)(unsafe.Add(mBase, uint32(v326)))
	*(*float64)(unsafe.Add(mBase, uint32(v25)+32)) = v328
	v330 = *(*float64)(unsafe.Add(mBase, uint32(v326)+8))
	*(*float64)(unsafe.Add(mBase, uint32(v25)+40)) = v330
	v332 = *(*float64)(unsafe.Add(mBase, uint32(v321)))
	*(*float64)(unsafe.Add(mBase, uint32(v25)+48)) = v332
	v334 = *(*float64)(unsafe.Add(mBase, uint32(v321)+8))
	*(*float64)(unsafe.Add(mBase, uint32(v25)+56)) = v334
	v340 = v291<<(uint(int32(4))%32) + v152 - int32(16)
	v341 = *(*float64)(unsafe.Add(mBase, uint32(v340)))
	*(*float64)(unsafe.Add(mBase, uint32(v25))) = v341
	v343 = *(*float64)(unsafe.Add(mBase, uint32(v340)+8))
	*(*float64)(unsafe.Add(mBase, uint32(v25)+8)) = v343
	v345 = *(*float64)(unsafe.Add(mBase, uint32(v35)+16))
	*(*float64)(unsafe.Add(mBase, uint32(v25)+16)) = v345
	v347 = *(*float64)(unsafe.Add(mBase, uint32(v35)+24))
	*(*float64)(unsafe.Add(mBase, uint32(v25)+24)) = v347
	v352 = F_lseg_interpt_lseg(m, int32(0), v23+int32(-32), v25)
	mBase = m.M
	v353 = m.ExcPending
	if v353 != 0 {
		goto L1
	} else {
		goto L107
	}
L105:
	;
	v357 = v291
	goto L106
L106:
	;
	v358 = int32(1)
	if int32(2) <= v357 {
		goto L111
	} else {
		goto L112
	}
L107:
	;
	if v352 != 0 {
		goto L108
	} else {
		goto L109
	}
L108:
	;
	v470 = int32(1)
	goto L90
L109:
	;
	goto L110
L110:
	;
	v355 = *(*int32)(unsafe.Add(mBase, uint32(v35)+4))
	v357 = v355
	goto L106
L111:
	;
	v362 = v358
	goto L114
L112:
	;
	v416 = v357
	goto L113
L113:
	;
	v436 = *(*int32)(unsafe.Add(mBase, uint32(v28)+4))
	v439 = v416
	v440 = v436
	goto L98
L114:
	;
	v384 = *(*float64)(unsafe.Add(mBase, uint32(v326)))
	*(*float64)(unsafe.Add(mBase, uint32(v25)+32)) = v384
	v386 = *(*float64)(unsafe.Add(mBase, uint32(v326)+8))
	*(*float64)(unsafe.Add(mBase, uint32(v25)+40)) = v386
	v388 = *(*float64)(unsafe.Add(mBase, uint32(v321)))
	*(*float64)(unsafe.Add(mBase, uint32(v25)+48)) = v388
	v390 = *(*float64)(unsafe.Add(mBase, uint32(v321)+8))
	*(*float64)(unsafe.Add(mBase, uint32(v25)+56)) = v390
	v394 = v152 + v362<<(uint(int32(4))%32)
	v396 = v394 - int32(16)
	v397 = *(*float64)(unsafe.Add(mBase, uint32(v396)))
	*(*float64)(unsafe.Add(mBase, uint32(v25))) = v397
	v399 = *(*float64)(unsafe.Add(mBase, uint32(v396)+8))
	*(*float64)(unsafe.Add(mBase, uint32(v25)+8)) = v399
	v401 = *(*float64)(unsafe.Add(mBase, uint32(v394)))
	*(*float64)(unsafe.Add(mBase, uint32(v25)+16)) = v401
	v403 = *(*float64)(unsafe.Add(mBase, uint32(v394)+8))
	*(*float64)(unsafe.Add(mBase, uint32(v25)+24)) = v403
	v408 = F_lseg_interpt_lseg(m, int32(0), v23+int32(-32), v25)
	mBase = m.M
	v409 = m.ExcPending
	if v409 != 0 {
		goto L1
	} else {
		goto L116
	}
L115:
	;
	v416 = v412
	goto L113
L116:
	;
	if v408 != 0 {
		v470 = v358
		goto L90
	} else {
		goto L117
	}
L117:
	;
	v411 = v362 + int32(1)
	v412 = *(*int32)(unsafe.Add(mBase, uint32(v35)+4))
	if v411 < v412 {
		v362 = v411
		goto L114
	} else {
		goto L118
	}
L118:
	;
	goto L115
L119:
	;
	goto L97
}
func F_path_n_gt(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v3 int32
	_ = v3
	var v4 int32
	_ = v4
	var v7 int32
	_ = v7
	var v8 int32
	_ = v8
	var v9 int32
	_ = v9
	var v10 int32
	_ = v10
	var v11 int32
	_ = v11
	var v12 int32
	_ = v12
	v3 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v4 = F_pg_detoast_datum(m, v3)
	mBase = m.M
	v7 = m.ExcPending
	if v7 != 0 {
		return int32(0)
	} else {
		v8 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
		v9 = F_pg_detoast_datum(m, v8)
		mBase = m.M
		v10 = m.ExcPending
		if v10 != 0 {
			return int32(0)
		} else {
			v11 = *(*int32)(unsafe.Add(mBase, uint32(v4)+4))
			v12 = *(*int32)(unsafe.Add(mBase, uint32(v9)+4))
			return base.B2i32(v12 < v11)
		}
	}
}
func F_path_n_le(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v3 int32
	_ = v3
	var v4 int32
	_ = v4
	var v7 int32
	_ = v7
	var v8 int32
	_ = v8
	var v9 int32
	_ = v9
	var v10 int32
	_ = v10
	var v11 int32
	_ = v11
	var v12 int32
	_ = v12
	v3 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v4 = F_pg_detoast_datum(m, v3)
	mBase = m.M
	v7 = m.ExcPending
	if v7 != 0 {
		return int32(0)
	} else {
		v8 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
		v9 = F_pg_detoast_datum(m, v8)
		mBase = m.M
		v10 = m.ExcPending
		if v10 != 0 {
			return int32(0)
		} else {
			v11 = *(*int32)(unsafe.Add(mBase, uint32(v4)+4))
			v12 = *(*int32)(unsafe.Add(mBase, uint32(v9)+4))
			return base.B2i32(v11 <= v12)
		}
	}
}
func F_path_open(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v2 int32
	_ = v2
	var v3 int32
	_ = v3
	var v6 int32
	_ = v6
	v2 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v3 = F_pg_detoast_datum_copy(m, v2)
	mBase = m.M
	v6 = m.ExcPending
	if v6 != 0 {
		return int32(0)
	} else {
		*(*int32)(unsafe.Add(mBase, uint32(v3)+8)) = int32(0)
		return v3
	}
}
func F_path_out(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v4 int32
	_ = v4
	var v5 int32
	_ = v5
	var v8 int32
	_ = v8
	var v9 int32
	_ = v9
	var v10 int32
	_ = v10
	var v11 int32
	_ = v11
	var v14 int32
	_ = v14
	var v15 int32
	_ = v15
	v4 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v5 = F_pg_detoast_datum(m, v4)
	mBase = m.M
	v8 = m.ExcPending
	if v8 != 0 {
		return int32(0)
	} else {
		v9 = *(*int32)(unsafe.Add(mBase, uint32(v5)+8))
		if v9 != 0 {
			v10 = int32(2)
		} else {
			v10 = int32(1)
		}
		v11 = *(*int32)(unsafe.Add(mBase, uint32(v5)+4))
		v14 = F_path_encode(m, v10, v11, v5+int32(16))
		mBase = m.M
		v15 = m.ExcPending
		if v15 != 0 {
			return int32(0)
		} else {
			return v14
		}
	}
}
func F_path_send(m *base.Module, l0 int32) int32 {
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
	var v14 int32
	_ = v14
	var v16 int32
	_ = v16
	var v17 int32
	_ = v17
	var v21 int32
	_ = v21
	var v22 int32
	_ = v22
	var v23 int32
	_ = v23
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
	var v40 int32
	_ = v40
	var v42 int32
	_ = v42
	var v57 int32
	_ = v57
	var v62 int32
	_ = v62
	var v69 int32
	_ = v69
	var v70 float64
	_ = v70
	var v72 int32
	_ = v72
	var v73 float64
	_ = v73
	var v75 int32
	_ = v75
	var v77 int32
	_ = v77
	var v78 int32
	_ = v78
	var v86 int32
	_ = v86
	var v87 int32
	_ = v87
	v6 = m.G0
	v8 = v6 - int32(16)
	m.G0 = v8
	v10 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v11 = F_pg_detoast_datum(m, v10)
	mBase = m.M
	v14 = m.ExcPending
	if v14 != 0 {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	return int32(0)
L2:
	;
	F_pq_begintypsend(m, v8)
	mBase = m.M
	v16 = m.ExcPending
	if v16 != 0 {
		goto L1
	} else {
		goto L3
	}
L3:
	;
	v17 = *(*int32)(unsafe.Add(mBase, uint32(v11)+8))
	F_enlargeStringInfo(m, v8, int32(1))
	mBase = m.M
	v21 = m.ExcPending
	if v21 != 0 {
		goto L1
	} else {
		goto L4
	}
L4:
	;
	v22 = *(*int32)(unsafe.Add(mBase, uint32(v8)+4))
	v23 = *(*int32)(unsafe.Add(mBase, uint32(v8)))
	*(*uint8)(unsafe.Add(mBase, uint32(v22+v23))) = uint8(base.B2i32(v17 != int32(0)))
	*(*int32)(unsafe.Add(mBase, uint32(v8)+4)) = v22 + int32(1)
	v31 = *(*int32)(unsafe.Add(mBase, uint32(v11)+4))
	F_enlargeStringInfo(m, v8, int32(4))
	mBase = m.M
	v34 = m.ExcPending
	if v34 != 0 {
		goto L1
	} else {
		goto L5
	}
L5:
	;
	v35 = *(*int32)(unsafe.Add(mBase, uint32(v8)+4))
	v36 = *(*int32)(unsafe.Add(mBase, uint32(v8)))
	v38 = int32(24)
	v40 = int32(65280)
	v42 = int32(8)
	*(*int32)(unsafe.Add(mBase, uint32(v35+v36))) = v31<<(uint(v38)%32) | v31&v40<<(uint(v42)%32) | (int32(base.Ui32(v31)>>(uint(v42)%32))&v40 | int32(base.Ui32(v31)>>(uint(v38)%32)))
	*(*int32)(unsafe.Add(mBase, uint32(v8)+4)) = v35 + int32(4)
	v57 = *(*int32)(unsafe.Add(mBase, uint32(v11)+4))
	if int32(0) < v57 {
		goto L6
	} else {
		goto L7
	}
L6:
	;
	v62 = int32(0)
	goto L9
L7:
	;
	goto L8
L8:
	;
	v86 = *(*int32)(unsafe.Add(mBase, uint32(v8)))
	v87 = *(*int32)(unsafe.Add(mBase, uint32(v8)+4))
	*(*int32)(unsafe.Add(mBase, uint32(v86))) = v87 << (uint(int32(2)) % 32)
	goto L14
L9:
	;
	v69 = v11 + int32(16) + v62<<(uint(int32(4))%32)
	v70 = *(*float64)(unsafe.Add(mBase, uint32(v69)))
	F_pq_sendfloat8(m, v8, v70)
	mBase = m.M
	v72 = m.ExcPending
	if v72 != 0 {
		goto L1
	} else {
		goto L11
	}
L10:
	;
	goto L8
L11:
	;
	v73 = *(*float64)(unsafe.Add(mBase, uint32(v69)+8))
	F_pq_sendfloat8(m, v8, v73)
	mBase = m.M
	v75 = m.ExcPending
	if v75 != 0 {
		goto L1
	} else {
		goto L12
	}
L12:
	;
	v77 = v62 + int32(1)
	v78 = *(*int32)(unsafe.Add(mBase, uint32(v11)+4))
	if v77 < v78 {
		v62 = v77
		goto L9
	} else {
		goto L13
	}
L13:
	;
	goto L10
L14:
	;
	m.G0 = v8 + int32(16)
	return v86
}
func F_path_usage_comparator(m *base.Module, l0 int32, l1 int32) int32 {
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
	var v12 int32
	_ = v12
	var v20 int32
	_ = v20
	var v21 int32
	_ = v21
	var v25 int32
	_ = v25
	var v27 float64
	_ = v27
	var v28 float64
	_ = v28
	var v33 float64
	_ = v33
	var v34 float64
	_ = v34
	var v39 int32
	_ = v39
	v6 = m.G0
	v8 = v6 - int32(32)
	m.G0 = v8
	v10 = *(*int32)(unsafe.Add(mBase, uint32(l1)))
	v11 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v12 = *(*int32)(unsafe.Add(mBase, uint32(v11)))
	F_cost_bitmap_tree_node(m, v12, v8+int32(24), v8+int32(8))
	mBase = m.M
	v20 = m.ExcPending
	if v20 != 0 {
		return int32(0)
	} else {
		v21 = *(*int32)(unsafe.Add(mBase, uint32(v10)))
		F_cost_bitmap_tree_node(m, v21, v8+int32(16), v8)
		mBase = m.M
		v25 = m.ExcPending
		if v25 != 0 {
			return int32(0)
		} else {
			v27 = *(*float64)(unsafe.Add(mBase, uint32(v8)+24))
			v28 = *(*float64)(unsafe.Add(mBase, uint32(v8)+16))
			if base.F64_lt(v27, v28) != 0 {
				v39 = int32(-1)
			} else {
				if base.F64_gt(v27, v28) != 0 {
					v39 = int32(1)
				} else {
					v33 = *(*float64)(unsafe.Add(mBase, uint32(v8)+8))
					v34 = *(*float64)(unsafe.Add(mBase, uint32(v8)))
					if base.F64_lt(v33, v34) != 0 {
						v39 = int32(-1)
					} else {
						v39 = base.F64_gt(v33, v34)
					}
				}
			}
			m.G0 = v8 + int32(32)
			return v39
		}
	}
}
