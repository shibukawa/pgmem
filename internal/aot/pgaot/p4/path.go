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
	var v82 int32
	_ = v82
	var v84 int64
	_ = v84
	var v88 int64
	_ = v88
	var v93 int32
	_ = v93
	var v94 int32
	_ = v94
	var v103 int32
	_ = v103
	var v108 int32
	_ = v108
	var v110 int32
	_ = v110
	v7 = m.G0
	v9 = v7 - int32(144)
	m.G0 = v9
	v15 = F___fstatat(m, int32(-100), l0, v9+int32(48), int32(0))
	mBase = m.M
	if v15 != 0 {
		v17 = *(*int32)(unsafe.Add(mBase, _c_F_PathNameDeleteTemporaryFile[0]))
		if v17 == int32(44) {
			v110 = int32(0)
			m.G0 = v9 + int32(144)
			return v110
		} else {
			v21 = F_unlink(m, l0)
			mBase = m.M
			if v21 < int32(0) {
				v29 = *(*int32)(unsafe.Add(mBase, _c_F_PathNameDeleteTemporaryFile[0]))
				if v29 == int32(44) {
					v110 = int32(0)
					m.G0 = v9 + int32(144)
					return v110
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
							v110 = v33
							m.G0 = v9 + int32(144)
							return v110
						} else {
							F_errcode_for_file_access(m)
							mBase = m.M
							v45 = m.ExcPending
							if v45 != 0 {
								return int32(0)
							} else {
								*(*int32)(unsafe.Add(mBase, uint32(v9))) = l0
								F_errmsg(m, int32(_a_F_PathNameDeleteTemporaryFile_0), v9)
								mBase = m.M
								v49 = m.ExcPending
								if v49 != 0 {
									return int32(0)
								} else {
									F_errfinish(m, int32(_a_F_PathNameDeleteTemporaryFile_1), int32(1966), int32(_a_F_PathNameDeleteTemporaryFile_2))
									mBase = m.M
									v54 = m.ExcPending
									if v54 != 0 {
										return int32(0)
									} else {
										v110 = v33
										m.G0 = v9 + int32(144)
										return v110
									}
								}
							}
						}
					}
				}
			} else {
				if v17 != 0 {
					*(*int32)(unsafe.Add(mBase, _c_F_PathNameDeleteTemporaryFile[0])) = v17
					v57 = int32(1)
					v60 = F_errstart(m, int32(15), int32(0))
					mBase = m.M
					v61 = m.ExcPending
					if v61 != 0 {
						return int32(0)
					} else {
						if v60 == int32(0) {
							v110 = v57
							m.G0 = v9 + int32(144)
							return v110
						} else {
							F_errcode_for_file_access(m)
							mBase = m.M
							v65 = m.ExcPending
							if v65 != 0 {
								return int32(0)
							} else {
								*(*int32)(unsafe.Add(mBase, uint32(v9)+32)) = l0
								F_errmsg(m, int32(_a_F_PathNameDeleteTemporaryFile_3), v9+int32(32))
								mBase = m.M
								v71 = m.ExcPending
								if v71 != 0 {
									return int32(0)
								} else {
									F_errfinish(m, int32(_a_F_PathNameDeleteTemporaryFile_1), int32(1977), int32(_a_F_PathNameDeleteTemporaryFile_2))
									mBase = m.M
									v76 = m.ExcPending
									if v76 != 0 {
										return int32(0)
									} else {
										v110 = v57
										m.G0 = v9 + int32(144)
										return v110
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
						v82 = int32(1)
						v84 = int64(*(*int32)(unsafe.Add(mBase, _c_F_PathNameDeleteTemporaryFile[1])))
						v88 = base.I64_div_s(v78, int64(1024))
						if base.B2i32(v84 < int64(0))|base.B2i32(v88 < v84) != 0 {
							v110 = v82
							m.G0 = v9 + int32(144)
							return v110
						} else {
							v93 = F_errstart(m, int32(15), int32(0))
							mBase = m.M
							v94 = m.ExcPending
							if v94 != 0 {
								return int32(0)
							} else {
								if v93 == int32(0) {
									v110 = v82
									m.G0 = v9 + int32(144)
									return v110
								} else {
									*(*int32)(unsafe.Add(mBase, uint32(v9)+20)) = v79
									*(*int32)(unsafe.Add(mBase, uint32(v9)+16)) = l0
									F_errmsg(m, int32(_a_F_PathNameDeleteTemporaryFile_4), v9+int32(16))
									mBase = m.M
									v103 = m.ExcPending
									if v103 != 0 {
										return int32(0)
									} else {
										F_errfinish(m, int32(_a_F_PathNameDeleteTemporaryFile_1), int32(1546), int32(_a_F_PathNameDeleteTemporaryFile_5))
										mBase = m.M
										v108 = m.ExcPending
										if v108 != 0 {
											return int32(0)
										} else {
											v110 = v82
											m.G0 = v9 + int32(144)
											return v110
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
				v82 = int32(1)
				v84 = int64(*(*int32)(unsafe.Add(mBase, _c_F_PathNameDeleteTemporaryFile[1])))
				v88 = base.I64_div_s(v78, int64(1024))
				if base.B2i32(v84 < int64(0))|base.B2i32(v88 < v84) != 0 {
					v110 = v82
					m.G0 = v9 + int32(144)
					return v110
				} else {
					v93 = F_errstart(m, int32(15), int32(0))
					mBase = m.M
					v94 = m.ExcPending
					if v94 != 0 {
						return int32(0)
					} else {
						if v93 == int32(0) {
							v110 = v82
							m.G0 = v9 + int32(144)
							return v110
						} else {
							*(*int32)(unsafe.Add(mBase, uint32(v9)+20)) = v79
							*(*int32)(unsafe.Add(mBase, uint32(v9)+16)) = l0
							F_errmsg(m, int32(_a_F_PathNameDeleteTemporaryFile_4), v9+int32(16))
							mBase = m.M
							v103 = m.ExcPending
							if v103 != 0 {
								return int32(0)
							} else {
								F_errfinish(m, int32(_a_F_PathNameDeleteTemporaryFile_1), int32(1546), int32(_a_F_PathNameDeleteTemporaryFile_5))
								mBase = m.M
								v108 = m.ExcPending
								if v108 != 0 {
									return int32(0)
								} else {
									v110 = v82
									m.G0 = v9 + int32(144)
									return v110
								}
							}
						}
					}
				}
			}
		} else {
			v29 = *(*int32)(unsafe.Add(mBase, _c_F_PathNameDeleteTemporaryFile[0]))
			if v29 == int32(44) {
				v110 = int32(0)
				m.G0 = v9 + int32(144)
				return v110
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
						v110 = v33
						m.G0 = v9 + int32(144)
						return v110
					} else {
						F_errcode_for_file_access(m)
						mBase = m.M
						v45 = m.ExcPending
						if v45 != 0 {
							return int32(0)
						} else {
							*(*int32)(unsafe.Add(mBase, uint32(v9))) = l0
							F_errmsg(m, int32(_a_F_PathNameDeleteTemporaryFile_0), v9)
							mBase = m.M
							v49 = m.ExcPending
							if v49 != 0 {
								return int32(0)
							} else {
								F_errfinish(m, int32(_a_F_PathNameDeleteTemporaryFile_1), int32(1966), int32(_a_F_PathNameDeleteTemporaryFile_2))
								mBase = m.M
								v54 = m.ExcPending
								if v54 != 0 {
									return int32(0)
								} else {
									v110 = v33
									m.G0 = v9 + int32(144)
									return v110
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
	var v11 int32
	_ = v11
	var v12 int32
	_ = v12
	var v15 int32
	_ = v15
	var v16 int32
	_ = v16
	var v19 int32
	_ = v19
	var v29 int32
	_ = v29
	var v34 int32
	_ = v34
	var v35 float64
	_ = v35
	var v36 float64
	_ = v36
	var v37 float64
	_ = v37
	var v39 float64
	_ = v39
	var v51 float64
	_ = v51
	var v52 float64
	_ = v52
	var v53 float64
	_ = v53
	var v55 float64
	_ = v55
	var v70 int32
	_ = v70
	var v87 int32
	_ = v87
	v11 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v12 = F_pg_detoast_datum_copy(m, v11)
	mBase = m.M
	v15 = m.ExcPending
	if v15 != 0 {
		goto L2
	} else {
		goto L3
	}
L1:
	;
	F_float_overflow_error(m)
	mBase = m.M
	v87 = m.ExcPending
	if v87 != 0 {
		goto L2
	} else {
		goto L12
	}
L2:
	;
	return int32(0)
L3:
	;
	v16 = *(*int32)(unsafe.Add(mBase, uint32(v12)+4))
	if int32(0) < v16 {
		goto L4
	} else {
		goto L5
	}
L4:
	;
	v19 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
	v29 = int32(0)
	goto L7
L5:
	;
	goto L6
L6:
	;
	return v12
L7:
	;
	v34 = v12 + int32(16) + v29<<(uint(int32(4))%32)
	v35 = *(*float64)(unsafe.Add(mBase, uint32(v34)))
	v36 = *(*float64)(unsafe.Add(mBase, uint32(v19)))
	v37 = base.F64_add(v35, v36)
	v39 = math.Float64frombits(uint64(0x7ff0000000000000))
	if base.B2i32(base.F64_ne(base.F64_abs(v37), v39)|base.F64_eq(base.F64_abs(v35), v39) == int32(0))&base.F64_ne(base.F64_abs(v36), v39) != 0 {
		goto L1
	} else {
		goto L9
	}
L8:
	;
	goto L6
L9:
	;
	v51 = *(*float64)(unsafe.Add(mBase, uint32(v34)+8))
	v52 = *(*float64)(unsafe.Add(mBase, uint32(v19)+8))
	v53 = base.F64_add(v51, v52)
	v55 = math.Float64frombits(uint64(0x7ff0000000000000))
	if base.B2i32(base.F64_ne(base.F64_abs(v53), v55)|base.F64_eq(base.F64_abs(v51), v55) == int32(0))&base.F64_ne(base.F64_abs(v52), v55) != 0 {
		goto L1
	} else {
		goto L10
	}
L10:
	;
	*(*float64)(unsafe.Add(mBase, uint32(v34)+8)) = v53
	*(*float64)(unsafe.Add(mBase, uint32(v34))) = v37
	v70 = v29 + int32(1)
	if v70 != v16 {
		v29 = v70
		goto L7
	} else {
		goto L11
	}
L11:
	;
	goto L8
L12:
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
	var v74 float64
	_ = v74
	var v76 float64
	_ = v76
	var v82 float64
	_ = v82
	var v83 float64
	_ = v83
	var v89 float64
	_ = v89
	var v91 float64
	_ = v91
	var v97 float64
	_ = v97
	var v99 float64
	_ = v99
	var v105 float64
	_ = v105
	var v111 float64
	_ = v111
	var v113 float64
	_ = v113
	var v119 float64
	_ = v119
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
	var v171 float64
	_ = v171
	var v172 float64
	_ = v172
	var v179 float64
	_ = v179
	var v180 float64
	_ = v180
	var v183 int32
	_ = v183
	var v184 float64
	_ = v184
	var v190 float64
	_ = v190
	var v192 float64
	_ = v192
	var v198 float64
	_ = v198
	var v199 float64
	_ = v199
	var v205 float64
	_ = v205
	var v207 float64
	_ = v207
	var v213 float64
	_ = v213
	var v215 float64
	_ = v215
	var v221 float64
	_ = v221
	var v227 float64
	_ = v227
	var v229 float64
	_ = v229
	var v235 float64
	_ = v235
	var v241 float64
	_ = v241
	var v243 int32
	_ = v243
	var v257 float64
	_ = v257
	var v258 float64
	_ = v258
	var v265 float64
	_ = v265
	var v266 float64
	_ = v266
	var v267 float64
	_ = v267
	var v270 int32
	_ = v270
	var v295 int32
	_ = v295
	var v297 int32
	_ = v297
	var v300 int32
	_ = v300
	var v317 int32
	_ = v317
	var v320 int32
	_ = v320
	var v323 int32
	_ = v323
	var v325 int32
	_ = v325
	var v328 int32
	_ = v328
	var v329 int32
	_ = v329
	var v330 float64
	_ = v330
	var v332 float64
	_ = v332
	var v334 float64
	_ = v334
	var v336 float64
	_ = v336
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
	var v393 int32
	_ = v393
	var v394 int32
	_ = v394
	var v395 float64
	_ = v395
	var v397 float64
	_ = v397
	var v399 int32
	_ = v399
	var v400 float64
	_ = v400
	var v402 float64
	_ = v402
	var v407 int32
	_ = v407
	var v408 int32
	_ = v408
	var v410 int32
	_ = v410
	var v411 int32
	_ = v411
	var v415 int32
	_ = v415
	var v435 int32
	_ = v435
	var v438 int32
	_ = v438
	var v440 int32
	_ = v440
	var v459 int32
	_ = v459
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
	if base.Ui64(int64(9218868437227405312)) < base.Ui64(base.I64_reinterpret_f64(v68)&int64(9223372036854775807)) {
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
	v74 = v68
	goto L12
L11:
	;
	v74 = v58
	goto L12
L12:
	;
	if base.F64_gt(v68, v58) != 0 {
		goto L13
	} else {
		goto L14
	}
L13:
	;
	v76 = v68
	goto L15
L14:
	;
	v76 = v74
	goto L15
L15:
	;
	if base.Ui64(base.I64_reinterpret_f64(v58)&int64(9223372036854775807)) <= base.Ui64(int64(9218868437227405312)) {
		goto L16
	} else {
		goto L17
	}
L16:
	;
	v82 = v76
	goto L18
L17:
	;
	v82 = v58
	goto L18
L18:
	;
	v83 = *(*float64)(unsafe.Add(mBase, uint32(v67)+8))
	if base.Ui64(int64(9218868437227405312)) < base.Ui64(base.I64_reinterpret_f64(v83)&int64(9223372036854775807)) {
		goto L19
	} else {
		goto L20
	}
L19:
	;
	v89 = v83
	goto L21
L20:
	;
	v89 = v57
	goto L21
L21:
	;
	if base.F64_gt(v83, v57) != 0 {
		goto L22
	} else {
		goto L23
	}
L22:
	;
	v91 = v83
	goto L24
L23:
	;
	v91 = v89
	goto L24
L24:
	;
	if base.Ui64(base.I64_reinterpret_f64(v57)&int64(9223372036854775807)) <= base.Ui64(int64(9218868437227405312)) {
		goto L25
	} else {
		goto L26
	}
L25:
	;
	v97 = v91
	goto L27
L26:
	;
	v97 = v57
	goto L27
L27:
	;
	if base.F64_lt(v68, v62) != 0 {
		goto L28
	} else {
		goto L29
	}
L28:
	;
	v99 = v68
	goto L30
L29:
	;
	v99 = v62
	goto L30
L30:
	;
	if base.Ui64(int64(9218868437227405312)) < base.Ui64(base.I64_reinterpret_f64(v62)&int64(9223372036854775807)) {
		goto L31
	} else {
		goto L32
	}
L31:
	;
	v105 = v68
	goto L33
L32:
	;
	v105 = v99
	goto L33
L33:
	;
	if base.Ui64(base.I64_reinterpret_f64(v68)&int64(9223372036854775807)) <= base.Ui64(int64(9218868437227405312)) {
		goto L34
	} else {
		goto L35
	}
L34:
	;
	v111 = v105
	goto L36
L35:
	;
	v111 = v62
	goto L36
L36:
	;
	if base.F64_lt(v83, v61) != 0 {
		goto L37
	} else {
		goto L38
	}
L37:
	;
	v113 = v83
	goto L39
L38:
	;
	v113 = v61
	goto L39
L39:
	;
	if base.Ui64(int64(9218868437227405312)) < base.Ui64(base.I64_reinterpret_f64(v61)&int64(9223372036854775807)) {
		goto L40
	} else {
		goto L41
	}
L40:
	;
	v119 = v83
	goto L42
L41:
	;
	v119 = v113
	goto L42
L42:
	;
	if base.Ui64(base.I64_reinterpret_f64(v83)&int64(9223372036854775807)) <= base.Ui64(int64(9218868437227405312)) {
		goto L43
	} else {
		goto L44
	}
L43:
	;
	v125 = v119
	goto L45
L44:
	;
	v125 = v61
	goto L45
L45:
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
L46:
	;
	goto L9
L47:
	;
	v267 = float64(1e-06)
	v270 = int32(0)
	if base.B2i32(base.F64_le(v148, base.F64_add(v258, v267)) == v270)|base.B2i32(base.F64_le(v266, base.F64_add(v144, v267)) == v270)|(base.B2i32(base.F64_le(v147, base.F64_add(v257, v267)) == v270)|base.B2i32(base.F64_le(v265, base.F64_add(v143, v267)) == v270))|base.B2i32(v39 <= v270) != 0 {
		v470 = v2
		goto L90
	} else {
		goto L91
	}
L48:
	;
	v257 = v153
	v258 = v154
	v265 = v153
	v266 = v154
	goto L47
L49:
	;
	goto L50
L50:
	;
	v159 = int32(1)
	v171 = v153
	v172 = v154
	v179 = v153
	v180 = v154
	goto L51
L51:
	;
	v183 = v152 + v159<<(uint(int32(4))%32)
	v184 = *(*float64)(unsafe.Add(mBase, uint32(v183)))
	if base.Ui64(int64(9218868437227405312)) < base.Ui64(base.I64_reinterpret_f64(v184)&int64(9223372036854775807)) {
		goto L53
	} else {
		goto L54
	}
L52:
	;
	v257 = v213
	v258 = v198
	v265 = v241
	v266 = v227
	goto L47
L53:
	;
	v190 = v184
	goto L55
L54:
	;
	v190 = v172
	goto L55
L55:
	;
	if base.F64_lt(v172, v184) != 0 {
		goto L56
	} else {
		goto L57
	}
L56:
	;
	v192 = v184
	goto L58
L57:
	;
	v192 = v190
	goto L58
L58:
	;
	if base.Ui64(base.I64_reinterpret_f64(v172)&int64(9223372036854775807)) <= base.Ui64(int64(9218868437227405312)) {
		goto L59
	} else {
		goto L60
	}
L59:
	;
	v198 = v192
	goto L61
L60:
	;
	v198 = v172
	goto L61
L61:
	;
	v199 = *(*float64)(unsafe.Add(mBase, uint32(v183)+8))
	if base.Ui64(int64(9218868437227405312)) < base.Ui64(base.I64_reinterpret_f64(v199)&int64(9223372036854775807)) {
		goto L62
	} else {
		goto L63
	}
L62:
	;
	v205 = v199
	goto L64
L63:
	;
	v205 = v171
	goto L64
L64:
	;
	if base.F64_lt(v171, v199) != 0 {
		goto L65
	} else {
		goto L66
	}
L65:
	;
	v207 = v199
	goto L67
L66:
	;
	v207 = v205
	goto L67
L67:
	;
	if base.Ui64(base.I64_reinterpret_f64(v171)&int64(9223372036854775807)) <= base.Ui64(int64(9218868437227405312)) {
		goto L68
	} else {
		goto L69
	}
L68:
	;
	v213 = v207
	goto L70
L69:
	;
	v213 = v171
	goto L70
L70:
	;
	if base.F64_lt(v184, v180) != 0 {
		goto L71
	} else {
		goto L72
	}
L71:
	;
	v215 = v184
	goto L73
L72:
	;
	v215 = v180
	goto L73
L73:
	;
	if base.Ui64(int64(9218868437227405312)) < base.Ui64(base.I64_reinterpret_f64(v180)&int64(9223372036854775807)) {
		goto L74
	} else {
		goto L75
	}
L74:
	;
	v221 = v184
	goto L76
L75:
	;
	v221 = v215
	goto L76
L76:
	;
	if base.Ui64(base.I64_reinterpret_f64(v184)&int64(9223372036854775807)) <= base.Ui64(int64(9218868437227405312)) {
		goto L77
	} else {
		goto L78
	}
L77:
	;
	v227 = v221
	goto L79
L78:
	;
	v227 = v180
	goto L79
L79:
	;
	if base.F64_lt(v199, v179) != 0 {
		goto L80
	} else {
		goto L81
	}
L80:
	;
	v229 = v199
	goto L82
L81:
	;
	v229 = v179
	goto L82
L82:
	;
	if base.Ui64(int64(9218868437227405312)) < base.Ui64(base.I64_reinterpret_f64(v179)&int64(9223372036854775807)) {
		goto L83
	} else {
		goto L84
	}
L83:
	;
	v235 = v199
	goto L85
L84:
	;
	v235 = v229
	goto L85
L85:
	;
	if base.Ui64(base.I64_reinterpret_f64(v199)&int64(9223372036854775807)) <= base.Ui64(int64(9218868437227405312)) {
		goto L86
	} else {
		goto L87
	}
L86:
	;
	v241 = v235
	goto L88
L87:
	;
	v241 = v179
	goto L88
L88:
	;
	v243 = v159 + int32(1)
	if v243 != v155 {
		v159 = v243
		v171 = v213
		v172 = v198
		v179 = v241
		v180 = v227
		goto L51
	} else {
		goto L89
	}
L89:
	;
	goto L52
L90:
	;
	m.G0 = v25 - int32(-64)
	return v470
L91:
	;
	v295 = v155
	v297 = v39
	v300 = v2
	goto L92
L92:
	;
	if v300 == int32(0) {
		goto L95
	} else {
		goto L96
	}
L93:
	;
	v470 = int32(0)
	goto L90
L94:
	;
	v459 = v300 + int32(1)
	if v459 < v440 {
		v295 = v438
		v297 = v440
		v300 = v459
		goto L92
	} else {
		goto L115
	}
L95:
	;
	v317 = *(*int32)(unsafe.Add(mBase, uint32(v28)+8))
	if v317 == int32(0) {
		v438 = v295
		v440 = v297
		goto L94
	} else {
		goto L98
	}
L96:
	;
	v320 = v300
	goto L97
L97:
	;
	if v295 <= int32(0) {
		v438 = v295
		v440 = v297
		goto L94
	} else {
		goto L99
	}
L98:
	;
	v320 = v297
	goto L97
L99:
	;
	v323 = int32(4)
	v325 = v33 + v300<<(uint(v323)%32)
	v328 = v28 + v320<<(uint(v323)%32)
	v329 = *(*int32)(unsafe.Add(mBase, uint32(v35)+8))
	if v329 != 0 {
		goto L100
	} else {
		goto L101
	}
L100:
	;
	v330 = *(*float64)(unsafe.Add(mBase, uint32(v328)))
	*(*float64)(unsafe.Add(mBase, uint32(v25)+32)) = v330
	v332 = *(*float64)(unsafe.Add(mBase, uint32(v328)+8))
	*(*float64)(unsafe.Add(mBase, uint32(v25)+40)) = v332
	v334 = *(*float64)(unsafe.Add(mBase, uint32(v325)))
	*(*float64)(unsafe.Add(mBase, uint32(v25)+48)) = v334
	v336 = *(*float64)(unsafe.Add(mBase, uint32(v325)+8))
	*(*float64)(unsafe.Add(mBase, uint32(v25)+56)) = v336
	v340 = v35 + v295<<(uint(int32(4))%32)
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
		goto L103
	}
L101:
	;
	v357 = v295
	goto L102
L102:
	;
	v358 = int32(1)
	if int32(2) <= v357 {
		goto L107
	} else {
		goto L108
	}
L103:
	;
	if v352 != 0 {
		goto L104
	} else {
		goto L105
	}
L104:
	;
	v470 = int32(1)
	goto L90
L105:
	;
	goto L106
L106:
	;
	v355 = *(*int32)(unsafe.Add(mBase, uint32(v35)+4))
	v357 = v355
	goto L102
L107:
	;
	v362 = v358
	goto L110
L108:
	;
	v415 = v357
	goto L109
L109:
	;
	v435 = *(*int32)(unsafe.Add(mBase, uint32(v28)+4))
	v438 = v415
	v440 = v435
	goto L94
L110:
	;
	v384 = *(*float64)(unsafe.Add(mBase, uint32(v328)))
	*(*float64)(unsafe.Add(mBase, uint32(v25)+32)) = v384
	v386 = *(*float64)(unsafe.Add(mBase, uint32(v328)+8))
	*(*float64)(unsafe.Add(mBase, uint32(v25)+40)) = v386
	v388 = *(*float64)(unsafe.Add(mBase, uint32(v325)))
	*(*float64)(unsafe.Add(mBase, uint32(v25)+48)) = v388
	v390 = *(*float64)(unsafe.Add(mBase, uint32(v325)+8))
	*(*float64)(unsafe.Add(mBase, uint32(v25)+56)) = v390
	v393 = v362 << (uint(int32(4)) % 32)
	v394 = v35 + v393
	v395 = *(*float64)(unsafe.Add(mBase, uint32(v394)))
	*(*float64)(unsafe.Add(mBase, uint32(v25))) = v395
	v397 = *(*float64)(unsafe.Add(mBase, uint32(v394)+8))
	*(*float64)(unsafe.Add(mBase, uint32(v25)+8)) = v397
	v399 = v393 + v152
	v400 = *(*float64)(unsafe.Add(mBase, uint32(v399)))
	*(*float64)(unsafe.Add(mBase, uint32(v25)+16)) = v400
	v402 = *(*float64)(unsafe.Add(mBase, uint32(v399)+8))
	*(*float64)(unsafe.Add(mBase, uint32(v25)+24)) = v402
	v407 = F_lseg_interpt_lseg(m, int32(0), v23+int32(-32), v25)
	mBase = m.M
	v408 = m.ExcPending
	if v408 != 0 {
		goto L1
	} else {
		goto L112
	}
L111:
	;
	v415 = v411
	goto L109
L112:
	;
	if v407 != 0 {
		v470 = v358
		goto L90
	} else {
		goto L113
	}
L113:
	;
	v410 = v362 + int32(1)
	v411 = *(*int32)(unsafe.Add(mBase, uint32(v35)+4))
	if v410 < v411 {
		v362 = v410
		goto L110
	} else {
		goto L114
	}
L114:
	;
	goto L111
L115:
	;
	goto L93
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
	var v40 int32
	_ = v40
	var v51 int32
	_ = v51
	var v56 int32
	_ = v56
	var v63 int32
	_ = v63
	var v64 float64
	_ = v64
	var v66 int32
	_ = v66
	var v67 float64
	_ = v67
	var v69 int32
	_ = v69
	var v71 int32
	_ = v71
	var v72 int32
	_ = v72
	var v80 int32
	_ = v80
	var v81 int32
	_ = v81
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
	v40 = int32(16711935)
	*(*int32)(unsafe.Add(mBase, uint32(v35+v36))) = base.I32_rotr(v31, int32(24))&v40 | base.I32_rotr(v31&v40, int32(8))
	*(*int32)(unsafe.Add(mBase, uint32(v8)+4)) = v35 + int32(4)
	v51 = *(*int32)(unsafe.Add(mBase, uint32(v11)+4))
	if int32(0) < v51 {
		goto L6
	} else {
		goto L7
	}
L6:
	;
	v56 = int32(0)
	goto L9
L7:
	;
	goto L8
L8:
	;
	v80 = *(*int32)(unsafe.Add(mBase, uint32(v8)))
	v81 = *(*int32)(unsafe.Add(mBase, uint32(v8)+4))
	*(*int32)(unsafe.Add(mBase, uint32(v80))) = v81 << (uint(int32(2)) % 32)
	goto L14
L9:
	;
	v63 = v11 + int32(16) + v56<<(uint(int32(4))%32)
	v64 = *(*float64)(unsafe.Add(mBase, uint32(v63)))
	F_pq_sendfloat8(m, v8, v64)
	mBase = m.M
	v66 = m.ExcPending
	if v66 != 0 {
		goto L1
	} else {
		goto L11
	}
L10:
	;
	goto L8
L11:
	;
	v67 = *(*float64)(unsafe.Add(mBase, uint32(v63)+8))
	F_pq_sendfloat8(m, v8, v67)
	mBase = m.M
	v69 = m.ExcPending
	if v69 != 0 {
		goto L1
	} else {
		goto L12
	}
L12:
	;
	v71 = v56 + int32(1)
	v72 = *(*int32)(unsafe.Add(mBase, uint32(v11)+4))
	if v71 < v72 {
		v56 = v71
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
	return v80
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
