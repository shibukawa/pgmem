package p4

import (
	base "github.com/shibukawa/pgmem/internal/aot/pgaot/base"
	"unsafe"
)

func F_enum_cmp(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v2 int32
	_ = v2
	var v3 int32
	_ = v3
	var v4 int32
	_ = v4
	var v7 int32
	_ = v7
	v2 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v3 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
	v4 = F_enum_cmp_internal(m, v2, v3, l0)
	mBase = m.M
	v7 = m.ExcPending
	if v7 != 0 {
		return int32(0)
	} else {
		return v4
	}
}
func F_enum_first(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v4 int32
	_ = v4
	var v6 int32
	_ = v6
	var v8 int32
	_ = v8
	var v10 int32
	_ = v10
	var v13 int32
	_ = v13
	var v15 int32
	_ = v15
	var v16 int32
	_ = v16
	var v26 int32
	_ = v26
	var v29 int32
	_ = v29
	var v33 int32
	_ = v33
	var v38 int32
	_ = v38
	var v42 int32
	_ = v42
	var v45 int32
	_ = v45
	var v46 int32
	_ = v46
	var v47 int32
	_ = v47
	var v51 int32
	_ = v51
	var v56 int32
	_ = v56
	v4 = m.G0
	v6 = v4 - int32(16)
	m.G0 = v6
	v8 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v10 = F_get_fn_expr_argtype(m, v8, int32(0))
	mBase = m.M
	v13 = m.ExcPending
	if v13 != 0 {
		return int32(0)
	} else {
		if v10 != 0 {
			v15 = F_enum_endpoint(m, v10, int32(1))
			mBase = m.M
			v16 = m.ExcPending
			if v16 != 0 {
				return int32(0)
			} else {
				if v15 == int32(0) {
					F_errstart_cold(m, int32(21), int32(0))
					mBase = m.M
					v42 = m.ExcPending
					if v42 != 0 {
						return int32(0)
					} else {
						F_errcode(m, int32(325))
						mBase = m.M
						v45 = m.ExcPending
						if v45 != 0 {
							return int32(0)
						} else {
							v46 = F_format_type_be(m, v10)
							mBase = m.M
							v47 = m.ExcPending
							if v47 != 0 {
								return int32(0)
							} else {
								*(*int32)(unsafe.Add(mBase, uint32(v6))) = v46
								F_errmsg(m, int32(158377), v6)
								mBase = m.M
								v51 = m.ExcPending
								if v51 != 0 {
									return int32(0)
								} else {
									F_errfinish(m, int32(498240), int32(460), int32(68222))
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
						}
					}
				} else {
					m.G0 = v6 + int32(16)
					return v15
				}
			}
		} else {
			F_errstart_cold(m, int32(21), int32(0))
			mBase = m.M
			v26 = m.ExcPending
			if v26 != 0 {
				return int32(0)
			} else {
				F_errcode(m, int32(1088))
				mBase = m.M
				v29 = m.ExcPending
				if v29 != 0 {
					return int32(0)
				} else {
					F_errmsg(m, int32(369914), int32(0))
					mBase = m.M
					v33 = m.ExcPending
					if v33 != 0 {
						return int32(0)
					} else {
						F_errfinish(m, int32(498240), int32(451), int32(68222))
						mBase = m.M
						v38 = m.ExcPending
						if v38 != 0 {
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
func F_enum_last(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v4 int32
	_ = v4
	var v6 int32
	_ = v6
	var v8 int32
	_ = v8
	var v10 int32
	_ = v10
	var v13 int32
	_ = v13
	var v15 int32
	_ = v15
	var v16 int32
	_ = v16
	var v26 int32
	_ = v26
	var v29 int32
	_ = v29
	var v33 int32
	_ = v33
	var v38 int32
	_ = v38
	var v42 int32
	_ = v42
	var v45 int32
	_ = v45
	var v46 int32
	_ = v46
	var v47 int32
	_ = v47
	var v51 int32
	_ = v51
	var v56 int32
	_ = v56
	v4 = m.G0
	v6 = v4 - int32(16)
	m.G0 = v6
	v8 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v10 = F_get_fn_expr_argtype(m, v8, int32(0))
	mBase = m.M
	v13 = m.ExcPending
	if v13 != 0 {
		return int32(0)
	} else {
		if v10 != 0 {
			v15 = F_enum_endpoint(m, v10, int32(-1))
			mBase = m.M
			v16 = m.ExcPending
			if v16 != 0 {
				return int32(0)
			} else {
				if v15 == int32(0) {
					F_errstart_cold(m, int32(21), int32(0))
					mBase = m.M
					v42 = m.ExcPending
					if v42 != 0 {
						return int32(0)
					} else {
						F_errcode(m, int32(325))
						mBase = m.M
						v45 = m.ExcPending
						if v45 != 0 {
							return int32(0)
						} else {
							v46 = F_format_type_be(m, v10)
							mBase = m.M
							v47 = m.ExcPending
							if v47 != 0 {
								return int32(0)
							} else {
								*(*int32)(unsafe.Add(mBase, uint32(v6))) = v46
								F_errmsg(m, int32(158377), v6)
								mBase = m.M
								v51 = m.ExcPending
								if v51 != 0 {
									return int32(0)
								} else {
									F_errfinish(m, int32(498240), int32(489), int32(78333))
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
						}
					}
				} else {
					m.G0 = v6 + int32(16)
					return v15
				}
			}
		} else {
			F_errstart_cold(m, int32(21), int32(0))
			mBase = m.M
			v26 = m.ExcPending
			if v26 != 0 {
				return int32(0)
			} else {
				F_errcode(m, int32(1088))
				mBase = m.M
				v29 = m.ExcPending
				if v29 != 0 {
					return int32(0)
				} else {
					F_errmsg(m, int32(369914), int32(0))
					mBase = m.M
					v33 = m.ExcPending
					if v33 != 0 {
						return int32(0)
					} else {
						F_errfinish(m, int32(498240), int32(480), int32(78333))
						mBase = m.M
						v38 = m.ExcPending
						if v38 != 0 {
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
func F_enum_lt(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v2 int32
	_ = v2
	var v3 int32
	_ = v3
	var v4 int32
	_ = v4
	var v7 int32
	_ = v7
	v2 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v3 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
	v4 = F_enum_cmp_internal(m, v2, v3, l0)
	mBase = m.M
	v7 = m.ExcPending
	if v7 != 0 {
		return int32(0)
	} else {
		return int32(base.Ui32(v4) >> (uint(int32(31)) % 32))
	}
}
func F_enum_smaller(m *base.Module, l0 int32) int32 {
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
	var v12 int32
	_ = v12
	v4 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v5 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
	v6 = F_enum_cmp_internal(m, v4, v5, l0)
	mBase = m.M
	v9 = m.ExcPending
	if v9 != 0 {
		return int32(0)
	} else {
		if v6 < int32(0) {
			v12 = v4
		} else {
			v12 = v5
		}
		return v12
	}
}
func F_load_enum_cache_data(m *base.Module, l0 int32) {
	mBase := m.M
	_ = mBase
	var v2 int32
	_ = v2
	var v19 int32
	_ = v19
	var v21 int32
	_ = v21
	var v23 int32
	_ = v23
	var v27 int32
	_ = v27
	var v28 int32
	_ = v28
	var v34 int32
	_ = v34
	var v36 int32
	_ = v36
	var v39 int32
	_ = v39
	var v40 int32
	_ = v40
	var v42 int32
	_ = v42
	var v47 int32
	_ = v47
	var v48 int32
	_ = v48
	var v49 int32
	_ = v49
	var v50 int32
	_ = v50
	var v55 int32
	_ = v55
	var v56 int32
	_ = v56
	var v57 int32
	_ = v57
	var v59 int32
	_ = v59
	var v72 int32
	_ = v72
	var v73 int32
	_ = v73
	var v74 int32
	_ = v74
	var v78 int32
	_ = v78
	var v79 int32
	_ = v79
	var v82 int32
	_ = v82
	var v83 int32
	_ = v83
	var v86 int32
	_ = v86
	var v87 int32
	_ = v87
	var v89 float32
	_ = v89
	var v92 int32
	_ = v92
	var v93 int32
	_ = v93
	var v94 int32
	_ = v94
	var v97 int32
	_ = v97
	var v100 int32
	_ = v100
	var v114 int32
	_ = v114
	var v117 int32
	_ = v117
	var v121 int32
	_ = v121
	var v123 int32
	_ = v123
	var v124 int32
	_ = v124
	var v127 int32
	_ = v127
	var v138 int32
	_ = v138
	var v139 int32
	_ = v139
	var v140 int32
	_ = v140
	var v141 int32
	_ = v141
	var v148 int32
	_ = v148
	var v150 int32
	_ = v150
	var v151 int32
	_ = v151
	var v154 int32
	_ = v154
	var v155 int32
	_ = v155
	var v157 int32
	_ = v157
	var v159 float32
	_ = v159
	var v161 int32
	_ = v161
	var v164 int32
	_ = v164
	var v166 int32
	_ = v166
	var v176 float32
	_ = v176
	var v180 int32
	_ = v180
	var v181 int32
	_ = v181
	var v182 int32
	_ = v182
	var v185 float32
	_ = v185
	var v189 int32
	_ = v189
	var v190 int32
	_ = v190
	var v191 int32
	_ = v191
	var v192 int32
	_ = v192
	var v193 float32
	_ = v193
	var v195 int32
	_ = v195
	var v201 int32
	_ = v201
	var v203 int32
	_ = v203
	var v215 int32
	_ = v215
	var v216 int32
	_ = v216
	var v218 int32
	_ = v218
	var v219 int32
	_ = v219
	var v220 int32
	_ = v220
	var v221 int32
	_ = v221
	var v235 int32
	_ = v235
	var v236 int32
	_ = v236
	var v244 int32
	_ = v244
	var v245 int32
	_ = v245
	var v248 int32
	_ = v248
	var v251 int32
	_ = v251
	var v254 int32
	_ = v254
	var v255 int32
	_ = v255
	var v257 int32
	_ = v257
	var v258 int32
	_ = v258
	var v263 int32
	_ = v263
	var v268 int32
	_ = v268
	var v270 int32
	_ = v270
	var v271 int32
	_ = v271
	var v273 int32
	_ = v273
	var v281 int32
	_ = v281
	var v284 int32
	_ = v284
	var v285 int32
	_ = v285
	var v286 int32
	_ = v286
	var v287 int32
	_ = v287
	var v291 int32
	_ = v291
	var v296 int32
	_ = v296
	v2 = int32(0)
	v19 = m.G0
	v21 = v19 + int32(-64)
	m.G0 = v21
	v23 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+13)))
	if v23 == int32(101) {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	v27 = F_palloc(m, int32(512))
	mBase = m.M
	v28 = m.ExcPending
	if v28 != 0 {
		goto L4
	} else {
		goto L5
	}
L2:
	;
	goto L3
L3:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v281 = m.ExcPending
	if v281 != 0 {
		goto L4
	} else {
		goto L70
	}
L4:
	;
	return
L5:
	;
	v34 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	F_ScanKeyInit(m, v19+int32(-48), int32(2), int32(3), int32(184), v34)
	mBase = m.M
	v36 = m.ExcPending
	if v36 != 0 {
		goto L4
	} else {
		goto L6
	}
L6:
	;
	v39 = F_table_open(m, int32(3501), int32(1))
	mBase = m.M
	v40 = m.ExcPending
	if v40 != 0 {
		goto L4
	} else {
		goto L8
	}
L7:
	;
	F_systable_endscan(m, v47)
	mBase = m.M
	v114 = m.ExcPending
	if v114 != 0 {
		goto L4
	} else {
		goto L22
	}
L8:
	;
	v42 = int32(1)
	v47 = F_systable_beginscan(m, v39, int32(3503), v42, int32(0), v42, v19+int32(-48))
	mBase = m.M
	v48 = m.ExcPending
	if v48 != 0 {
		goto L4
	} else {
		goto L9
	}
L9:
	;
	v49 = F_systable_getnext(m, v47)
	mBase = m.M
	v50 = m.ExcPending
	if v50 != 0 {
		goto L4
	} else {
		goto L10
	}
L10:
	;
	if v49 == int32(0) {
		goto L11
	} else {
		goto L12
	}
L11:
	;
	v97 = v2
	v100 = v27
	goto L7
L12:
	;
	goto L13
L13:
	;
	v55 = v49
	v56 = v2
	v57 = int32(64)
	v59 = v27
	goto L14
L14:
	;
	v72 = *(*int32)(unsafe.Add(mBase, uint32(v55)+16))
	v73 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v72)+22)))
	v74 = v72 + v73
	if v57 <= v56 {
		goto L16
	} else {
		goto L17
	}
L15:
	;
	v97 = v92
	v100 = v83
	goto L7
L16:
	;
	v78 = F_repalloc(m, v59, v57<<(uint(int32(4))%32))
	mBase = m.M
	v79 = m.ExcPending
	if v79 != 0 {
		goto L4
	} else {
		goto L19
	}
L17:
	;
	v82 = v57
	v83 = v59
	goto L18
L18:
	;
	v86 = v83 + v56<<(uint(int32(3))%32)
	v87 = *(*int32)(unsafe.Add(mBase, uint32(v74)))
	*(*int32)(unsafe.Add(mBase, uint32(v86))) = v87
	v89 = *(*float32)(unsafe.Add(mBase, uint32(v74)+8))
	*(*float32)(unsafe.Add(mBase, uint32(v86)+4)) = v89
	v92 = v56 + int32(1)
	v93 = F_systable_getnext(m, v47)
	mBase = m.M
	v94 = m.ExcPending
	if v94 != 0 {
		goto L4
	} else {
		goto L20
	}
L19:
	;
	v82 = v57 << (uint(int32(1)) % 32)
	v83 = v78
	goto L18
L20:
	;
	if v93 != 0 {
		v55 = v93
		v56 = v92
		v57 = v82
		v59 = v83
		goto L14
	} else {
		goto L21
	}
L21:
	;
	goto L15
L22:
	;
	F_sequence_close(m, v39, int32(1))
	mBase = m.M
	v117 = m.ExcPending
	if v117 != 0 {
		goto L4
	} else {
		goto L23
	}
L23:
	;
	F_pg_qsort(m, v100, v97, int32(8), int32(1624))
	mBase = m.M
	v121 = m.ExcPending
	if v121 != 0 {
		goto L4
	} else {
		goto L24
	}
L24:
	;
	v123 = v97 - int32(1)
	v124 = int32(0)
	if v124 < v123 {
		goto L25
	} else {
		goto L26
	}
L25:
	;
	v127 = v123
	goto L27
L26:
	;
	v127 = v124
	goto L27
L27:
	;
	v138 = v2
	v139 = v2
	v140 = int32(1)
	v141 = v2
	goto L28
L28:
	;
	if v127 != v141 {
		goto L30
	} else {
		goto L31
	}
L29:
	;
	v244 = int32(4520272)
	v245 = *(*int32)(unsafe.Add(mBase, _consts[28]))
	v248 = *(*int32)(unsafe.Add(mBase, _consts[382]))
	*(*int32)(unsafe.Add(mBase, _consts[28])) = v248
	v251 = v97 << (uint(int32(3)) % 32)
	v254 = F_palloc(m, v251+int32(12))
	mBase = m.M
	v255 = m.ExcPending
	if v255 != 0 {
		goto L4
	} else {
		goto L58
	}
L30:
	;
	v148 = int32(1)
	v150 = F_bms_make_singleton(m, int32(0))
	mBase = m.M
	v151 = m.ExcPending
	if v151 != 0 {
		goto L4
	} else {
		goto L33
	}
L31:
	;
	v235 = v138
	v236 = v139
	goto L32
L32:
	;
	goto L29
L33:
	;
	v154 = v100 + v141<<(uint(int32(3))%32)
	v155 = *(*int32)(unsafe.Add(mBase, uint32(v154)))
	v157 = v141 + int32(1)
	if v97 <= v157 {
		v201 = v148
		v203 = v150
		goto L34
	} else {
		goto L35
	}
L34:
	;
	v215 = base.B2i32(v140 < v201)
	if v140 < v201 {
		goto L44
	} else {
		goto L45
	}
L35:
	;
	v159 = *(*float32)(unsafe.Add(mBase, uint32(v154)+4))
	v161 = v157
	v164 = v148
	v166 = v150
	v176 = v159
	goto L36
L36:
	;
	v180 = v100 + v161<<(uint(int32(3))%32)
	v181 = *(*int32)(unsafe.Add(mBase, uint32(v180)))
	v182 = v181 - v155
	if base.Ui32(int32(8191)) < base.Ui32(v182) {
		v201 = v164
		v203 = v166
		goto L34
	} else {
		goto L38
	}
L37:
	;
	v201 = v191
	v203 = v192
	goto L34
L38:
	;
	v185 = *(*float32)(unsafe.Add(mBase, uint32(v180)+4))
	if base.F32_lt(v176, v185) != 0 {
		goto L39
	} else {
		goto L40
	}
L39:
	;
	v189 = F_bms_add_member(m, v166, v182)
	mBase = m.M
	v190 = m.ExcPending
	if v190 != 0 {
		goto L4
	} else {
		goto L42
	}
L40:
	;
	v191 = v164
	v192 = v166
	v193 = v176
	goto L41
L41:
	;
	v195 = v161 + int32(1)
	if v195 != v97 {
		v161 = v195
		v164 = v191
		v166 = v192
		v176 = v193
		goto L36
	} else {
		goto L43
	}
L42:
	;
	v191 = v164 + int32(1)
	v192 = v189
	v193 = v185
	goto L41
L43:
	;
	goto L37
L44:
	;
	v216 = v138
	goto L46
L45:
	;
	v216 = v203
	goto L46
L46:
	;
	F_bms_free(m, v216)
	mBase = m.M
	v218 = m.ExcPending
	if v218 != 0 {
		goto L4
	} else {
		goto L47
	}
L47:
	;
	if v140 < v201 {
		goto L48
	} else {
		goto L49
	}
L48:
	;
	v219 = v155
	goto L50
L49:
	;
	v219 = v139
	goto L50
L50:
	;
	if v140 < v201 {
		goto L51
	} else {
		goto L52
	}
L51:
	;
	v220 = v203
	goto L53
L52:
	;
	v220 = v138
	goto L53
L53:
	;
	if v140 < v201 {
		goto L54
	} else {
		goto L55
	}
L54:
	;
	v221 = v201
	goto L56
L55:
	;
	v221 = v140
	goto L56
L56:
	;
	if v221 < v97+(v141^int32(-1)) {
		v138 = v220
		v139 = v219
		v140 = v221
		v141 = v157
		goto L28
	} else {
		goto L57
	}
L57:
	;
	v235 = v220
	v236 = v219
	goto L32
L58:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v254))) = v236
	v257 = F_bms_copy(m, v235)
	mBase = m.M
	v258 = m.ExcPending
	if v258 != 0 {
		goto L4
	} else {
		goto L59
	}
L59:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v254)+8)) = v97
	*(*int32)(unsafe.Add(mBase, uint32(v254)+4)) = v257
	if v251 != 0 {
		goto L61
	} else {
		goto L62
	}
L60:
	;
	*(*int32)(unsafe.Add(mBase, _consts[28])) = v245
	F_pfree(m, v100)
	mBase = m.M
	v268 = m.ExcPending
	if v268 != 0 {
		goto L4
	} else {
		goto L64
	}
L61:
	;
	v263 = F__emscripten_memcpy_bulkmem(m, v254+int32(12), v100, v251)
	mBase = m.M
	goto L63
L62:
	;
	goto L63
L63:
	;
	goto L60
L64:
	;
	F_bms_free(m, v235)
	mBase = m.M
	v270 = m.ExcPending
	if v270 != 0 {
		goto L4
	} else {
		goto L65
	}
L65:
	;
	v271 = *(*int32)(unsafe.Add(mBase, uint32(l0)+316))
	if v271 != 0 {
		goto L66
	} else {
		goto L67
	}
L66:
	;
	F_pfree(m, v271)
	mBase = m.M
	v273 = m.ExcPending
	if v273 != 0 {
		goto L4
	} else {
		goto L69
	}
L67:
	;
	goto L68
L68:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+316)) = v254
	m.G0 = v21 - int32(-64)
	return
L69:
	;
	goto L68
L70:
	;
	F_errcode(m, int32(151027844))
	mBase = m.M
	v284 = m.ExcPending
	if v284 != 0 {
		goto L4
	} else {
		goto L71
	}
L71:
	;
	v285 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v286 = F_format_type_be(m, v285)
	mBase = m.M
	v287 = m.ExcPending
	if v287 != 0 {
		goto L4
	} else {
		goto L72
	}
L72:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v21))) = v286
	F_errmsg(m, int32(287991), v21)
	mBase = m.M
	v291 = m.ExcPending
	if v291 != 0 {
		goto L4
	} else {
		goto L73
	}
L73:
	;
	F_errfinish(m, int32(500408), int32(2758), int32(505397))
	mBase = m.M
	v296 = m.ExcPending
	if v296 != 0 {
		goto L4
	} else {
		goto L74
	}
L74:
	;
	base.Wasm_trap_unreachable()
	for {
	}
}
