package p1

import (
	base "github.com/shibukawa/pgmem/internal/aot/pgaot/base"
	"unsafe"
)

func F_get_multirange_range(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v4 int32
	_ = v4
	var v7 int32
	_ = v7
	var v12 int32
	_ = v12
	var v13 int32
	_ = v13
	var v15 int32
	_ = v15
	var v17 int32
	_ = v17
	v4 = F_SearchSysCache1(m, int32(54), l0)
	mBase = m.M
	v7 = m.ExcPending
	if v7 != 0 {
		return int32(0)
	} else {
		if v4 == int32(0) {
			return int32(0)
		} else {
			v12 = *(*int32)(unsafe.Add(mBase, uint32(v4)+16))
			v13 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v12)+22)))
			v15 = *(*int32)(unsafe.Add(mBase, uint32(v12+v13)))
			F_ReleaseCatCache(m, v4)
			mBase = m.M
			v17 = m.ExcPending
			if v17 != 0 {
				return int32(0)
			} else {
				return v15
			}
		}
	}
}
func F_make_multirange(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) int32 {
	mBase := m.M
	_ = mBase
	var v5 int32
	_ = v5
	var v20 int32
	_ = v20
	var v27 int32
	_ = v27
	var v29 int32
	_ = v29
	var v30 int32
	_ = v30
	var v37 int32
	_ = v37
	var v40 int32
	_ = v40
	var v41 int32
	_ = v41
	var v47 int32
	_ = v47
	var v52 int32
	_ = v52
	var v53 int32
	_ = v53
	var v60 int32
	_ = v60
	var v61 int32
	_ = v61
	var v63 int32
	_ = v63
	var v64 int32
	_ = v64
	var v71 int32
	_ = v71
	var v72 int32
	_ = v72
	var v80 int32
	_ = v80
	var v81 int32
	_ = v81
	var v83 int32
	_ = v83
	var v91 int32
	_ = v91
	var v100 int32
	_ = v100
	var v101 int32
	_ = v101
	var v104 int32
	_ = v104
	var v107 int32
	_ = v107
	var v112 int32
	_ = v112
	var v113 int32
	_ = v113
	var v128 int32
	_ = v128
	var v141 int32
	_ = v141
	var v143 int32
	_ = v143
	var v150 int32
	_ = v150
	var v153 int32
	_ = v153
	var v154 int32
	_ = v154
	var v156 int32
	_ = v156
	var v173 int32
	_ = v173
	var v174 int32
	_ = v174
	var v176 int32
	_ = v176
	var v183 int32
	_ = v183
	var v192 int32
	_ = v192
	var v193 int32
	_ = v193
	var v196 int32
	_ = v196
	var v200 int32
	_ = v200
	var v207 int32
	_ = v207
	var v208 int32
	_ = v208
	var v228 int32
	_ = v228
	var v232 int32
	_ = v232
	var v235 int32
	_ = v235
	var v236 int32
	_ = v236
	var v237 int32
	_ = v237
	var v238 int32
	_ = v238
	var v243 int32
	_ = v243
	var v245 int32
	_ = v245
	var v246 int32
	_ = v246
	var v247 int32
	_ = v247
	var v249 int32
	_ = v249
	var v251 int32
	_ = v251
	var v271 int32
	_ = v271
	var v283 int32
	_ = v283
	var v286 int32
	_ = v286
	var v288 int32
	_ = v288
	var v296 int32
	_ = v296
	var v300 int32
	_ = v300
	var v306 int32
	_ = v306
	var v309 int32
	_ = v309
	var v310 int32
	_ = v310
	var v311 int32
	_ = v311
	var v312 int32
	_ = v312
	var v317 int32
	_ = v317
	var v319 int32
	_ = v319
	var v320 int32
	_ = v320
	var v322 int32
	_ = v322
	var v324 int32
	_ = v324
	var v342 int32
	_ = v342
	var v345 int32
	_ = v345
	v5 = int32(0)
	if l3 != 0 {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	F_qsort_arg(m, l3, l2, int32(4), int32(1442), l1)
	mBase = m.M
	v20 = m.ExcPending
	if v20 != 0 {
		goto L4
	} else {
		goto L5
	}
L2:
	;
	goto L3
L3:
	;
	if int32(0) < l2 {
		goto L6
	} else {
		goto L7
	}
L4:
	;
	return int32(0)
L5:
	;
	goto L3
L6:
	;
	v27 = v5
	v29 = v5
	v30 = v5
	goto L9
L7:
	;
	v91 = v5
	goto L8
L8:
	;
	v100 = v91 - int32(1)
	v101 = int32(0)
	if v101 < v100 {
		goto L30
	} else {
		goto L31
	}
L9:
	;
	v37 = int32(2)
	v40 = *(*int32)(unsafe.Add(mBase, uint32(l3+v30<<(uint(v37)%32))))
	v41 = *(*int32)(unsafe.Add(mBase, uint32(v40)))
	v47 = int32(*(*int8)(unsafe.Add(mBase, uint32(v40+int32(base.Ui32(v41)>>(uint(v37)%32))-int32(1)))))
	goto L12
L10:
	;
	v91 = v81
	goto L8
L11:
	;
	v83 = v30 + int32(1)
	if v83 != l2 {
		v27 = v80
		v29 = v81
		v30 = v83
		goto L9
	} else {
		goto L28
	}
L12:
	;
	if v47&int32(1) != 0 {
		v80 = v27
		v81 = v29
		goto L11
	} else {
		goto L13
	}
L13:
	;
	if v27 == int32(0) {
		goto L15
	} else {
		goto L16
	}
L14:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l3+v29<<(uint(int32(2))%32)))) = v40
	v80 = v40
	v81 = v29 + int32(1)
	goto L11
L15:
	;
	goto L14
L16:
	;
	goto L17
L17:
	;
	v52 = F_range_adjacent_internal(m, l1, v27, v40)
	mBase = m.M
	v53 = m.ExcPending
	if v53 != 0 {
		goto L4
	} else {
		goto L18
	}
L18:
	;
	if v52 != 0 {
		goto L19
	} else {
		goto L20
	}
L19:
	;
	v60 = F_range_union_internal(m, l1, v27, v40, int32(0))
	mBase = m.M
	v61 = m.ExcPending
	if v61 != 0 {
		goto L4
	} else {
		goto L22
	}
L20:
	;
	goto L21
L21:
	;
	v63 = F_range_before_internal(m, l1, v27, v40)
	mBase = m.M
	v64 = m.ExcPending
	if v64 != 0 {
		goto L4
	} else {
		goto L23
	}
L22:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l3+v29<<(uint(int32(2))%32)-int32(4)))) = v60
	v80 = v60
	v81 = v29
	goto L11
L23:
	;
	if v63 != 0 {
		goto L24
	} else {
		goto L25
	}
L24:
	;
	goto L14
L25:
	;
	goto L26
L26:
	;
	v71 = F_range_union_internal(m, l1, v27, v40, int32(1))
	mBase = m.M
	v72 = m.ExcPending
	if v72 != 0 {
		goto L4
	} else {
		goto L27
	}
L27:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l3+v29<<(uint(int32(2))%32)-int32(4)))) = v71
	v80 = v71
	v81 = v29
	goto L11
L28:
	;
	goto L10
L29:
	;
	if int32(0) < v91 {
		goto L37
	} else {
		goto L38
	}
L30:
	;
	v104 = v100
	goto L32
L31:
	;
	v104 = v101
	goto L32
L32:
	;
	v107 = v104<<(uint(int32(2))%32) + v91
	v112 = *(*int32)(unsafe.Add(mBase, uint32(l1)+200))
	v113 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v112)+11)))
	if v113 == int32(105) {
		v128 = (v107 + int32(15)) & int32(-4)
		goto L29
	} else {
		goto L33
	}
L33:
	;
	switch v113 - int32(99) {
	case 0:
		goto L36
	case 1:
		goto L35
	default:
		goto L34
	}
L34:
	;
	v128 = (v107 + int32(13)) & int32(-2)
	goto L29
L35:
	;
	v128 = (v107 + int32(19)) & int32(-8)
	goto L29
L36:
	;
	v128 = v107 + int32(12)
	goto L29
L37:
	;
	v141 = v128
	v143 = int32(0)
	goto L40
L38:
	;
	v183 = v128
	goto L39
L39:
	;
	v192 = F_palloc0(m, v183)
	mBase = m.M
	v193 = m.ExcPending
	if v193 != 0 {
		goto L4
	} else {
		goto L48
	}
L40:
	;
	v150 = int32(2)
	v153 = *(*int32)(unsafe.Add(mBase, uint32(l3+v143<<(uint(v150)%32))))
	v154 = *(*int32)(unsafe.Add(mBase, uint32(v153)))
	v156 = int32(base.Ui32(v154) >> (uint(v150) % 32))
	if base.B2i32(v113 != int32(105)) == int32(0) {
		v173 = (v156 - int32(6)) & int32(-4)
		goto L42
	} else {
		goto L43
	}
L41:
	;
	v183 = v174
	goto L39
L42:
	;
	v174 = v173 + v141
	v176 = v143 + int32(1)
	if v176 != v91 {
		v141 = v174
		v143 = v176
		goto L40
	} else {
		goto L47
	}
L43:
	;
	switch v113 - int32(99) {
	case 0:
		goto L46
	case 1:
		goto L45
	default:
		goto L44
	}
L44:
	;
	v173 = v156&int32(1073741822) - int32(8)
	goto L42
L45:
	;
	v173 = (v156 - int32(2)) & int32(-8)
	goto L42
L46:
	;
	v173 = v156 - int32(9)
	goto L42
L47:
	;
	goto L41
L48:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v192)+8)) = v91
	*(*int32)(unsafe.Add(mBase, uint32(v192)+4)) = l0
	v196 = int32(2)
	*(*int32)(unsafe.Add(mBase, uint32(v192))) = v183 << (uint(v196) % 32)
	v200 = v91 << (uint(v196) % 32)
	v207 = *(*int32)(unsafe.Add(mBase, uint32(l1)+200))
	v208 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v207)+11)))
	if v208 == int32(105) {
		v228 = (v91*int32(5) + int32(11)) & int32(-4)
		goto L49
	} else {
		goto L50
	}
L49:
	;
	if v91 <= int32(0) {
		goto L54
	} else {
		goto L55
	}
L50:
	;
	switch v208 - int32(99) {
	case 0:
		goto L53
	case 1:
		goto L52
	default:
		goto L51
	}
L51:
	;
	v228 = (v91*int32(5) + int32(9)) & int32(-2)
	goto L49
L52:
	;
	v228 = (v91*int32(5) + int32(15)) & int32(-8)
	goto L49
L53:
	;
	v228 = v200 + v91 + int32(8)
	goto L49
L54:
	;
	return v192
L55:
	;
	v232 = v192 + int32(12)
	v235 = v232 + v200 - int32(4)
	v236 = *(*int32)(unsafe.Add(mBase, uint32(l3)))
	v237 = *(*int32)(unsafe.Add(mBase, uint32(v236)))
	v238 = int32(2)
	v243 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v236+int32(base.Ui32(v237)>>(uint(v238)%32))-int32(1)))))
	*(*uint8)(unsafe.Add(mBase, uint32(v235))) = uint8(v243)
	v245 = v228 + v192
	v246 = *(*int32)(unsafe.Add(mBase, uint32(l3)))
	v247 = *(*int32)(unsafe.Add(mBase, uint32(v246)))
	v249 = int32(base.Ui32(v247) >> (uint(v238) % 32))
	v251 = v249 - int32(9)
	if v251 != 0 {
		goto L56
	} else {
		goto L57
	}
L56:
	;
	base.MemoryCopy(m, v245, v246+int32(8), v251)
	goto L58
L57:
	;
	goto L58
L58:
	;
	if v208 != int32(105) {
		goto L60
	} else {
		goto L61
	}
L59:
	;
	if v91 == int32(1) {
		goto L54
	} else {
		goto L65
	}
L60:
	;
	switch v208 - int32(99) {
	case 0:
		v271 = v251
		goto L59
	case 1:
		goto L64
	default:
		goto L63
	}
L61:
	;
	goto L62
L62:
	;
	v271 = (v249 - int32(6)) & int32(-4)
	goto L59
L63:
	;
	v271 = v249&int32(1073741822) - int32(8)
	goto L59
L64:
	;
	v271 = (v249 - int32(2)) & int32(-8)
	goto L59
L65:
	;
	v283 = int32(0)
	v286 = v245 + v271
	v288 = int32(1)
	goto L66
L66:
	;
	v296 = v288 << (uint(int32(2)) % 32)
	v300 = v286 - v245
	if v288&int32(3) != 0 {
		goto L68
	} else {
		goto L69
	}
L67:
	;
	goto L54
L68:
	;
	v306 = v300 - v283
	goto L70
L69:
	;
	v306 = v300 | int32(-2147483648)
	goto L70
L70:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v232+v296-int32(4)))) = v306
	v309 = l3 + v296
	v310 = *(*int32)(unsafe.Add(mBase, uint32(v309)))
	v311 = *(*int32)(unsafe.Add(mBase, uint32(v310)))
	v312 = int32(2)
	v317 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v310+int32(base.Ui32(v311)>>(uint(v312)%32))-int32(1)))))
	*(*uint8)(unsafe.Add(mBase, uint32(v288+v235))) = uint8(v317)
	v319 = *(*int32)(unsafe.Add(mBase, uint32(v309)))
	v320 = *(*int32)(unsafe.Add(mBase, uint32(v319)))
	v322 = int32(base.Ui32(v320) >> (uint(v312) % 32))
	v324 = v322 - int32(9)
	if v324 != 0 {
		goto L71
	} else {
		goto L72
	}
L71:
	;
	base.MemoryCopy(m, v286, v319+int32(8), v324)
	goto L73
L72:
	;
	goto L73
L73:
	;
	if base.B2i32(v208 != int32(105)) == int32(0) {
		goto L75
	} else {
		goto L76
	}
L74:
	;
	v345 = v288 + int32(1)
	if v345 != v91 {
		v283 = v300
		v286 = v342 + v286
		v288 = v345
		goto L66
	} else {
		goto L80
	}
L75:
	;
	v342 = (v322 - int32(6)) & int32(-4)
	goto L74
L76:
	;
	goto L77
L77:
	;
	switch v208 - int32(99) {
	case 0:
		v342 = v324
		goto L74
	case 1:
		goto L79
	default:
		goto L78
	}
L78:
	;
	v342 = v322&int32(1073741822) - int32(8)
	goto L74
L79:
	;
	v342 = (v322 - int32(2)) & int32(-8)
	goto L74
L80:
	;
	goto L67
}
func F_multirange_before_multirange(m *base.Module, l0 int32) int32 {
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
					v34 = *(*int32)(unsafe.Add(mBase, uint32(v12)+8))
					if v34 == v33 {
						v60 = v33
						m.G0 = v9 + int32(48)
						return v60
					} else {
						v37 = *(*int32)(unsafe.Add(mBase, uint32(v17)+8))
						if v37 == int32(0) {
							v60 = v33
							m.G0 = v9 + int32(48)
							return v60
						} else {
							v40 = *(*int32)(unsafe.Add(mBase, uint32(v32)+296))
							v46 = v9 + int32(32)
							F_multirange_get_bounds(m, v40, v12, v34-int32(1), v9+int32(40), v46)
							mBase = m.M
							v48 = m.ExcPending
							if v48 != 0 {
								return int32(0)
							} else {
								v51 = v9 + int32(24)
								F_multirange_get_bounds(m, v40, v17, int32(0), v51, v9+int32(16))
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
					v25 = F_lookup_type_cache(m, v19, int32(_a_F_multirange_before_multirange_0))
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
								F_errmsg_internal(m, int32(_a_F_multirange_before_multirange_1), v9)
								mBase = m.M
								v74 = m.ExcPending
								if v74 != 0 {
									return int32(0)
								} else {
									F_errfinish(m, int32(_a_F_multirange_before_multirange_2), int32(558), int32(_a_F_multirange_before_multirange_3))
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
							v34 = *(*int32)(unsafe.Add(mBase, uint32(v12)+8))
							if v34 == v33 {
								v60 = v33
								m.G0 = v9 + int32(48)
								return v60
							} else {
								v37 = *(*int32)(unsafe.Add(mBase, uint32(v17)+8))
								if v37 == int32(0) {
									v60 = v33
									m.G0 = v9 + int32(48)
									return v60
								} else {
									v40 = *(*int32)(unsafe.Add(mBase, uint32(v32)+296))
									v46 = v9 + int32(32)
									F_multirange_get_bounds(m, v40, v12, v34-int32(1), v9+int32(40), v46)
									mBase = m.M
									v48 = m.ExcPending
									if v48 != 0 {
										return int32(0)
									} else {
										v51 = v9 + int32(24)
										F_multirange_get_bounds(m, v40, v17, int32(0), v51, v9+int32(16))
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
				v25 = F_lookup_type_cache(m, v19, int32(_a_F_multirange_before_multirange_0))
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
							F_errmsg_internal(m, int32(_a_F_multirange_before_multirange_1), v9)
							mBase = m.M
							v74 = m.ExcPending
							if v74 != 0 {
								return int32(0)
							} else {
								F_errfinish(m, int32(_a_F_multirange_before_multirange_2), int32(558), int32(_a_F_multirange_before_multirange_3))
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
						v34 = *(*int32)(unsafe.Add(mBase, uint32(v12)+8))
						if v34 == v33 {
							v60 = v33
							m.G0 = v9 + int32(48)
							return v60
						} else {
							v37 = *(*int32)(unsafe.Add(mBase, uint32(v17)+8))
							if v37 == int32(0) {
								v60 = v33
								m.G0 = v9 + int32(48)
								return v60
							} else {
								v40 = *(*int32)(unsafe.Add(mBase, uint32(v32)+296))
								v46 = v9 + int32(32)
								F_multirange_get_bounds(m, v40, v12, v34-int32(1), v9+int32(40), v46)
								mBase = m.M
								v48 = m.ExcPending
								if v48 != 0 {
									return int32(0)
								} else {
									v51 = v9 + int32(24)
									F_multirange_get_bounds(m, v40, v17, int32(0), v51, v9+int32(16))
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
func F_multirange_constructor0(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v5 int32
	_ = v5
	var v7 int32
	_ = v7
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
	var v18 int32
	_ = v18
	var v19 int32
	_ = v19
	var v22 int32
	_ = v22
	var v23 int32
	_ = v23
	var v24 int32
	_ = v24
	var v27 int32
	_ = v27
	var v29 int32
	_ = v29
	var v30 int32
	_ = v30
	var v31 int32
	_ = v31
	var v33 int32
	_ = v33
	var v34 int32
	_ = v34
	var v42 int32
	_ = v42
	var v46 int32
	_ = v46
	var v51 int32
	_ = v51
	var v55 int32
	_ = v55
	var v59 int32
	_ = v59
	var v64 int32
	_ = v64
	v5 = m.G0
	v7 = v5 - int32(16)
	m.G0 = v7
	v9 = int32(*(*uint16)(unsafe.Add(mBase, uint32(l0)+18)))
	if v9 == int32(0) {
		v12 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
		v13 = F_get_fn_expr_rettype(m, v12)
		mBase = m.M
		v16 = m.ExcPending
		if v16 != 0 {
			return int32(0)
		} else {
			v17 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
			v18 = *(*int32)(unsafe.Add(mBase, uint32(v17)+16))
			if v18 != 0 {
				v19 = *(*int32)(unsafe.Add(mBase, uint32(v18)))
				if v19 == v13 {
					v29 = v18
					v30 = *(*int32)(unsafe.Add(mBase, uint32(v29)+296))
					v31 = int32(0)
					v33 = F_make_multirange(m, v13, v30, v31, v31)
					mBase = m.M
					v34 = m.ExcPending
					if v34 != 0 {
						return int32(0)
					} else {
						m.G0 = v7 + int32(16)
						return v33
					}
				} else {
					v22 = F_lookup_type_cache(m, v13, int32(_a_F_multirange_constructor0_0))
					mBase = m.M
					v23 = m.ExcPending
					if v23 != 0 {
						return int32(0)
					} else {
						v24 = *(*int32)(unsafe.Add(mBase, uint32(v22)+296))
						if v24 == int32(0) {
							F_errstart_cold(m, int32(21), int32(0))
							mBase = m.M
							v55 = m.ExcPending
							if v55 != 0 {
								return int32(0)
							} else {
								*(*int32)(unsafe.Add(mBase, uint32(v7))) = v13
								F_errmsg_internal(m, int32(_a_F_multirange_constructor0_1), v7)
								mBase = m.M
								v59 = m.ExcPending
								if v59 != 0 {
									return int32(0)
								} else {
									F_errfinish(m, int32(_a_F_multirange_constructor0_2), int32(558), int32(_a_F_multirange_constructor0_3))
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
							v27 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
							*(*int32)(unsafe.Add(mBase, uint32(v27)+16)) = v22
							v29 = v22
							v30 = *(*int32)(unsafe.Add(mBase, uint32(v29)+296))
							v31 = int32(0)
							v33 = F_make_multirange(m, v13, v30, v31, v31)
							mBase = m.M
							v34 = m.ExcPending
							if v34 != 0 {
								return int32(0)
							} else {
								m.G0 = v7 + int32(16)
								return v33
							}
						}
					}
				}
			} else {
				v22 = F_lookup_type_cache(m, v13, int32(_a_F_multirange_constructor0_0))
				mBase = m.M
				v23 = m.ExcPending
				if v23 != 0 {
					return int32(0)
				} else {
					v24 = *(*int32)(unsafe.Add(mBase, uint32(v22)+296))
					if v24 == int32(0) {
						F_errstart_cold(m, int32(21), int32(0))
						mBase = m.M
						v55 = m.ExcPending
						if v55 != 0 {
							return int32(0)
						} else {
							*(*int32)(unsafe.Add(mBase, uint32(v7))) = v13
							F_errmsg_internal(m, int32(_a_F_multirange_constructor0_1), v7)
							mBase = m.M
							v59 = m.ExcPending
							if v59 != 0 {
								return int32(0)
							} else {
								F_errfinish(m, int32(_a_F_multirange_constructor0_2), int32(558), int32(_a_F_multirange_constructor0_3))
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
						v27 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
						*(*int32)(unsafe.Add(mBase, uint32(v27)+16)) = v22
						v29 = v22
						v30 = *(*int32)(unsafe.Add(mBase, uint32(v29)+296))
						v31 = int32(0)
						v33 = F_make_multirange(m, v13, v30, v31, v31)
						mBase = m.M
						v34 = m.ExcPending
						if v34 != 0 {
							return int32(0)
						} else {
							m.G0 = v7 + int32(16)
							return v33
						}
					}
				}
			}
		}
	} else {
		F_errstart_cold(m, int32(21), int32(0))
		mBase = m.M
		v42 = m.ExcPending
		if v42 != 0 {
			return int32(0)
		} else {
			F_errmsg_internal(m, int32(_a_F_multirange_constructor0_4), int32(0))
			mBase = m.M
			v46 = m.ExcPending
			if v46 != 0 {
				return int32(0)
			} else {
				F_errfinish(m, int32(_a_F_multirange_constructor0_2), int32(1069), int32(_a_F_multirange_constructor0_5))
				mBase = m.M
				v51 = m.ExcPending
				if v51 != 0 {
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
func F_multirange_contained_by_multirange(m *base.Module, l0 int32) int32 {
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
	var v43 int32
	_ = v43
	var v47 int32
	_ = v47
	var v52 int32
	_ = v52
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
					v34 = F_multirange_contains_multirange_internal(m, v33, v17, v12)
					mBase = m.M
					v35 = m.ExcPending
					if v35 != 0 {
						return int32(0)
					} else {
						m.G0 = v9 + int32(16)
						return v34
					}
				} else {
					v25 = F_lookup_type_cache(m, v19, int32(_a_F_multirange_contained_by_multirange_0))
					mBase = m.M
					v26 = m.ExcPending
					if v26 != 0 {
						return int32(0)
					} else {
						v27 = *(*int32)(unsafe.Add(mBase, uint32(v25)+296))
						if v27 == int32(0) {
							F_errstart_cold(m, int32(21), int32(0))
							mBase = m.M
							v43 = m.ExcPending
							if v43 != 0 {
								return int32(0)
							} else {
								*(*int32)(unsafe.Add(mBase, uint32(v9))) = v19
								F_errmsg_internal(m, int32(_a_F_multirange_contained_by_multirange_1), v9)
								mBase = m.M
								v47 = m.ExcPending
								if v47 != 0 {
									return int32(0)
								} else {
									F_errfinish(m, int32(_a_F_multirange_contained_by_multirange_2), int32(558), int32(_a_F_multirange_contained_by_multirange_3))
									mBase = m.M
									v52 = m.ExcPending
									if v52 != 0 {
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
							v34 = F_multirange_contains_multirange_internal(m, v33, v17, v12)
							mBase = m.M
							v35 = m.ExcPending
							if v35 != 0 {
								return int32(0)
							} else {
								m.G0 = v9 + int32(16)
								return v34
							}
						}
					}
				}
			} else {
				v25 = F_lookup_type_cache(m, v19, int32(_a_F_multirange_contained_by_multirange_0))
				mBase = m.M
				v26 = m.ExcPending
				if v26 != 0 {
					return int32(0)
				} else {
					v27 = *(*int32)(unsafe.Add(mBase, uint32(v25)+296))
					if v27 == int32(0) {
						F_errstart_cold(m, int32(21), int32(0))
						mBase = m.M
						v43 = m.ExcPending
						if v43 != 0 {
							return int32(0)
						} else {
							*(*int32)(unsafe.Add(mBase, uint32(v9))) = v19
							F_errmsg_internal(m, int32(_a_F_multirange_contained_by_multirange_1), v9)
							mBase = m.M
							v47 = m.ExcPending
							if v47 != 0 {
								return int32(0)
							} else {
								F_errfinish(m, int32(_a_F_multirange_contained_by_multirange_2), int32(558), int32(_a_F_multirange_contained_by_multirange_3))
								mBase = m.M
								v52 = m.ExcPending
								if v52 != 0 {
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
						v34 = F_multirange_contains_multirange_internal(m, v33, v17, v12)
						mBase = m.M
						v35 = m.ExcPending
						if v35 != 0 {
							return int32(0)
						} else {
							m.G0 = v9 + int32(16)
							return v34
						}
					}
				}
			}
		}
	}
}
func F_multirange_ge(m *base.Module, l0 int32) int32 {
	var v2 int32
	_ = v2
	var v5 int32
	_ = v5
	v2 = F_multirange_cmp(m, l0)
	v5 = m.ExcPending
	if v5 != 0 {
		return int32(0)
	} else {
		return int32(base.Ui32(v2^int32(-1)) >> (uint(int32(31)) % 32))
	}
}
func F_multirange_gist_compress(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v8 int32
	_ = v8
	var v9 int32
	_ = v9
	var v13 int32
	_ = v13
	var v14 int32
	_ = v14
	var v17 int32
	_ = v17
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
	var v24 int32
	_ = v24
	var v25 int32
	_ = v25
	var v27 int32
	_ = v27
	var v29 int32
	_ = v29
	var v32 int32
	_ = v32
	var v33 int32
	_ = v33
	var v36 int32
	_ = v36
	var v38 int32
	_ = v38
	var v40 int32
	_ = v40
	var v41 int32
	_ = v41
	var v45 int32
	_ = v45
	var v47 int32
	_ = v47
	var v48 int32
	_ = v48
	var v50 int32
	_ = v50
	var v51 int32
	_ = v51
	var v55 int32
	_ = v55
	var v60 int32
	_ = v60
	var v62 int32
	_ = v62
	var v64 int32
	_ = v64
	var v65 int32
	_ = v65
	v8 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v9 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v8)+14)))
	if v9 != int32(1) {
		return v8
	} else {
		v13 = *(*int32)(unsafe.Add(mBase, uint32(v8)))
		v14 = F_pg_detoast_datum(m, v13)
		mBase = m.M
		v17 = m.ExcPending
		if v17 != 0 {
			return int32(0)
		} else {
			v19 = F_palloc(m, int32(16))
			mBase = m.M
			v20 = m.ExcPending
			if v20 != 0 {
				return int32(0)
			} else {
				v21 = *(*int32)(unsafe.Add(mBase, uint32(v14)+4))
				v22 = F_multirange_get_typcache(m, l0, v21)
				mBase = m.M
				v23 = m.ExcPending
				if v23 != 0 {
					return int32(0)
				} else {
					v24 = *(*int32)(unsafe.Add(mBase, uint32(v22)+296))
					v25 = m.G0
					v27 = v25 - int32(32)
					m.G0 = v27
					v29 = *(*int32)(unsafe.Add(mBase, uint32(v14)+8))
					if v29 == int32(0) {
						v32 = F_make_empty_range(m, v24)
						mBase = m.M
						v33 = m.ExcPending
						if v33 != 0 {
							return int32(0)
						} else {
							v55 = v32
							m.G0 = v27 + int32(32)
							*(*int32)(unsafe.Add(mBase, uint32(v19))) = v55
							v60 = *(*int32)(unsafe.Add(mBase, uint32(v8)+4))
							*(*int32)(unsafe.Add(mBase, uint32(v19)+4)) = v60
							v62 = *(*int32)(unsafe.Add(mBase, uint32(v8)+8))
							*(*int32)(unsafe.Add(mBase, uint32(v19)+8)) = v62
							v64 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v8)+12)))
							v65 = int32(0)
							*(*uint8)(unsafe.Add(mBase, uint32(v19)+14)) = uint8(v65)
							*(*uint16)(unsafe.Add(mBase, uint32(v19)+12)) = uint16(v64)
							return v19
						}
					} else {
						v36 = v27 + int32(24)
						v38 = v27 + int32(8)
						F_multirange_get_bounds(m, v24, v14, int32(0), v36, v38)
						mBase = m.M
						v40 = m.ExcPending
						if v40 != 0 {
							return int32(0)
						} else {
							v41 = *(*int32)(unsafe.Add(mBase, uint32(v14)+8))
							v45 = v27 + int32(16)
							F_multirange_get_bounds(m, v24, v14, v41-int32(1), v38, v45)
							mBase = m.M
							v47 = m.ExcPending
							if v47 != 0 {
								return int32(0)
							} else {
								v48 = int32(0)
								v50 = F_make_range(m, v24, v36, v45, v48, v48)
								mBase = m.M
								v51 = m.ExcPending
								if v51 != 0 {
									return int32(0)
								} else {
									v55 = v50
									m.G0 = v27 + int32(32)
									*(*int32)(unsafe.Add(mBase, uint32(v19))) = v55
									v60 = *(*int32)(unsafe.Add(mBase, uint32(v8)+4))
									*(*int32)(unsafe.Add(mBase, uint32(v19)+4)) = v60
									v62 = *(*int32)(unsafe.Add(mBase, uint32(v8)+8))
									*(*int32)(unsafe.Add(mBase, uint32(v19)+8)) = v62
									v64 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v8)+12)))
									v65 = int32(0)
									*(*uint8)(unsafe.Add(mBase, uint32(v19)+14)) = uint8(v65)
									*(*uint16)(unsafe.Add(mBase, uint32(v19)+12)) = uint16(v64)
									return v19
								}
							}
						}
					}
				}
			}
		}
	}
}
func F_multirange_gist_consistent(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v9 int32
	_ = v9
	var v11 int32
	_ = v11
	var v13 int32
	_ = v13
	var v14 int32
	_ = v14
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
	var v29 int32
	_ = v29
	var v31 int32
	_ = v31
	var v38 int32
	_ = v38
	var v39 int32
	_ = v39
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
	var v51 int32
	_ = v51
	var v57 int32
	_ = v57
	var v62 int32
	_ = v62
	var v67 int32
	_ = v67
	var v68 int32
	_ = v68
	var v69 int32
	_ = v69
	var v70 int32
	_ = v70
	var v71 int32
	_ = v71
	var v72 int32
	_ = v72
	var v73 int32
	_ = v73
	var v74 int32
	_ = v74
	var v77 int32
	_ = v77
	var v78 int32
	_ = v78
	var v79 int32
	_ = v79
	var v87 int32
	_ = v87
	var v91 int32
	_ = v91
	var v96 int32
	_ = v96
	v9 = m.G0
	v11 = v9 - int32(32)
	m.G0 = v11
	v13 = *(*int32)(unsafe.Add(mBase, uint32(l0)+44))
	v14 = int32(*(*uint16)(unsafe.Add(mBase, uint32(l0)+36)))
	v15 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
	v16 = *(*int32)(unsafe.Add(mBase, uint32(l0)+52))
	v17 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v18 = *(*int32)(unsafe.Add(mBase, uint32(v17)))
	v19 = F_pg_detoast_datum(m, v18)
	mBase = m.M
	v22 = m.ExcPending
	if v22 != 0 {
		return int32(0)
	} else {
		v23 = int32(1)
		*(*uint8)(unsafe.Add(mBase, uint32(v16))) = uint8(v23)
		v25 = *(*int32)(unsafe.Add(mBase, uint32(v19)+4))
		v26 = F_range_get_typcache(m, l0, v25)
		mBase = m.M
		v27 = m.ExcPending
		if v27 != 0 {
			return int32(0)
		} else {
			v28 = *(*int32)(unsafe.Add(mBase, uint32(v17)+8))
			v29 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v28)+16)))
			v31 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v28+v29)+12)))
			if v31&int32(1) != 0 {
				if v13 != int32(_a_F_multirange_gist_consistent_0) {
					if v13 == int32(3831) {
						v42 = F_pg_detoast_datum(m, v15)
						mBase = m.M
						v43 = m.ExcPending
						if v43 != 0 {
							return int32(0)
						} else {
							v44 = F_range_gist_consistent_leaf_range(m, v26, v14, v19, v42)
							mBase = m.M
							v45 = m.ExcPending
							if v45 != 0 {
								return int32(0)
							} else {
								v79 = v44
								m.G0 = v11 + int32(32)
								return v79
							}
						}
					} else {
						if v13 != 0 {
							if v14 == int32(16) {
								v77 = F_range_contains_elem_internal(m, v26, v19, v15)
								mBase = m.M
								v78 = m.ExcPending
								if v78 != 0 {
									return int32(0)
								} else {
									v79 = v77
									m.G0 = v11 + int32(32)
									return v79
								}
							} else {
								F_errstart_cold(m, int32(21), int32(0))
								mBase = m.M
								v51 = m.ExcPending
								if v51 != 0 {
									return int32(0)
								} else {
									*(*int32)(unsafe.Add(mBase, uint32(v11)+16)) = v14
									F_errmsg_internal(m, int32(_a_F_multirange_gist_consistent_1), v11+int32(16))
									mBase = m.M
									v57 = m.ExcPending
									if v57 != 0 {
										return int32(0)
									} else {
										F_errfinish(m, int32(_a_F_multirange_gist_consistent_2), int32(1138), int32(_a_F_multirange_gist_consistent_3))
										mBase = m.M
										v62 = m.ExcPending
										if v62 != 0 {
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
							v38 = F_pg_detoast_datum(m, v15)
							mBase = m.M
							v39 = m.ExcPending
							if v39 != 0 {
								return int32(0)
							} else {
								v40 = F_range_gist_consistent_leaf_multirange(m, v26, v14, v19, v38)
								mBase = m.M
								v41 = m.ExcPending
								if v41 != 0 {
									return int32(0)
								} else {
									v79 = v40
									m.G0 = v11 + int32(32)
									return v79
								}
							}
						}
					}
				} else {
					v38 = F_pg_detoast_datum(m, v15)
					mBase = m.M
					v39 = m.ExcPending
					if v39 != 0 {
						return int32(0)
					} else {
						v40 = F_range_gist_consistent_leaf_multirange(m, v26, v14, v19, v38)
						mBase = m.M
						v41 = m.ExcPending
						if v41 != 0 {
							return int32(0)
						} else {
							v79 = v40
							m.G0 = v11 + int32(32)
							return v79
						}
					}
				}
			} else {
				if v13 != int32(_a_F_multirange_gist_consistent_0) {
					if v13 == int32(3831) {
						v71 = F_pg_detoast_datum(m, v15)
						mBase = m.M
						v72 = m.ExcPending
						if v72 != 0 {
							return int32(0)
						} else {
							v73 = F_range_gist_consistent_int_range(m, v26, v14, v19, v71)
							mBase = m.M
							v74 = m.ExcPending
							if v74 != 0 {
								return int32(0)
							} else {
								v79 = v73
								m.G0 = v11 + int32(32)
								return v79
							}
						}
					} else {
						if v13 != 0 {
							if v14 != int32(16) {
								F_errstart_cold(m, int32(21), int32(0))
								mBase = m.M
								v87 = m.ExcPending
								if v87 != 0 {
									return int32(0)
								} else {
									*(*int32)(unsafe.Add(mBase, uint32(v11))) = v14
									F_errmsg_internal(m, int32(_a_F_multirange_gist_consistent_1), v11)
									mBase = m.M
									v91 = m.ExcPending
									if v91 != 0 {
										return int32(0)
									} else {
										F_errfinish(m, int32(_a_F_multirange_gist_consistent_2), int32(1049), int32(_a_F_multirange_gist_consistent_4))
										mBase = m.M
										v96 = m.ExcPending
										if v96 != 0 {
											return int32(0)
										} else {
											base.Wasm_trap_unreachable()
											for {
											}
										}
									}
								}
							} else {
								v77 = F_range_contains_elem_internal(m, v26, v19, v15)
								mBase = m.M
								v78 = m.ExcPending
								if v78 != 0 {
									return int32(0)
								} else {
									v79 = v77
									m.G0 = v11 + int32(32)
									return v79
								}
							}
						} else {
							v67 = F_pg_detoast_datum(m, v15)
							mBase = m.M
							v68 = m.ExcPending
							if v68 != 0 {
								return int32(0)
							} else {
								v69 = F_range_gist_consistent_int_multirange(m, v26, v14, v19, v67)
								mBase = m.M
								v70 = m.ExcPending
								if v70 != 0 {
									return int32(0)
								} else {
									v79 = v69
									m.G0 = v11 + int32(32)
									return v79
								}
							}
						}
					}
				} else {
					v67 = F_pg_detoast_datum(m, v15)
					mBase = m.M
					v68 = m.ExcPending
					if v68 != 0 {
						return int32(0)
					} else {
						v69 = F_range_gist_consistent_int_multirange(m, v26, v14, v19, v67)
						mBase = m.M
						v70 = m.ExcPending
						if v70 != 0 {
							return int32(0)
						} else {
							v79 = v69
							m.G0 = v11 + int32(32)
							return v79
						}
					}
				}
			}
		}
	}
}
func F_multirange_intersect_agg_transfn(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v2 int32
	_ = v2
	var v11 int32
	_ = v11
	var v13 int32
	_ = v13
	var v16 int32
	_ = v16
	var v18 int32
	_ = v18
	var v21 int32
	_ = v21
	var v27 int32
	_ = v27
	var v28 int32
	_ = v28
	var v33 int32
	_ = v33
	var v35 int32
	_ = v35
	var v38 int32
	_ = v38
	var v39 int32
	_ = v39
	var v43 int32
	_ = v43
	var v46 int32
	_ = v46
	var v47 int32
	_ = v47
	var v49 int32
	_ = v49
	var v52 int32
	_ = v52
	var v53 int32
	_ = v53
	var v54 int32
	_ = v54
	var v57 int32
	_ = v57
	var v58 int32
	_ = v58
	var v59 int32
	_ = v59
	var v62 int32
	_ = v62
	var v63 int32
	_ = v63
	var v64 int32
	_ = v64
	var v67 int32
	_ = v67
	var v69 int32
	_ = v69
	var v70 int32
	_ = v70
	var v71 int32
	_ = v71
	var v72 int32
	_ = v72
	var v73 int32
	_ = v73
	var v74 int32
	_ = v74
	var v75 int32
	_ = v75
	var v76 int32
	_ = v76
	var v77 int32
	_ = v77
	var v83 int32
	_ = v83
	var v84 int32
	_ = v84
	var v85 int32
	_ = v85
	var v98 int32
	_ = v98
	var v99 int32
	_ = v99
	var v102 int32
	_ = v102
	var v104 int32
	_ = v104
	var v108 int32
	_ = v108
	var v113 int32
	_ = v113
	var v115 int32
	_ = v115
	var v121 int32
	_ = v121
	var v122 int32
	_ = v122
	var v123 int32
	_ = v123
	var v136 int32
	_ = v136
	var v137 int32
	_ = v137
	var v140 int32
	_ = v140
	var v142 int32
	_ = v142
	var v152 int32
	_ = v152
	var v153 int32
	_ = v153
	var v154 int32
	_ = v154
	var v155 int32
	_ = v155
	var v163 int32
	_ = v163
	var v167 int32
	_ = v167
	var v172 int32
	_ = v172
	var v176 int32
	_ = v176
	var v180 int32
	_ = v180
	var v185 int32
	_ = v185
	var v189 int32
	_ = v189
	var v193 int32
	_ = v193
	var v198 int32
	_ = v198
	v2 = int32(0)
	v11 = m.G0
	v13 = v11 - int32(16)
	m.G0 = v13
	v16 = v13 + int32(12)
	v18 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	if v18 == v2 {
		goto L6
	} else {
		goto L7
	}
L1:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v189 = m.ExcPending
	if v189 != 0 {
		goto L20
	} else {
		goto L56
	}
L2:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v176 = m.ExcPending
	if v176 != 0 {
		goto L20
	} else {
		goto L53
	}
L3:
	;
	if v46 != 0 {
		goto L17
	} else {
		goto L18
	}
L4:
	;
	v46 = v43
	goto L3
L5:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v16))) = v38
	v43 = v39
	goto L4
L6:
	;
	v35 = int32(0)
	if v16 == v35 {
		v43 = v35
		goto L4
	} else {
		goto L16
	}
L7:
	;
	v21 = *(*int32)(unsafe.Add(mBase, uint32(v18)))
	switch v21 - int32(429) {
	case 0:
		goto L9
	case 1:
		goto L8
	default:
		goto L6
	}
L8:
	;
	if v16 == int32(0) {
		goto L13
	} else {
		goto L14
	}
L9:
	;
	if v16 == int32(0) {
		goto L10
	} else {
		goto L11
	}
L10:
	;
	v46 = int32(1)
	goto L3
L11:
	;
	goto L12
L12:
	;
	v27 = *(*int32)(unsafe.Add(mBase, uint32(v18)+168))
	v28 = *(*int32)(unsafe.Add(mBase, uint32(v27)+20))
	v38 = v28
	v39 = int32(1)
	goto L5
L13:
	;
	v46 = int32(2)
	goto L3
L14:
	;
	goto L15
L15:
	;
	v33 = *(*int32)(unsafe.Add(mBase, uint32(v18)+368))
	v38 = v33
	v39 = int32(2)
	goto L5
L16:
	;
	v38 = v35
	v39 = v2
	goto L5
L17:
	;
	v47 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v49 = F_get_fn_expr_argtype(m, v47, int32(1))
	mBase = m.M
	v52 = m.ExcPending
	if v52 != 0 {
		goto L20
	} else {
		goto L21
	}
L18:
	;
	goto L19
L19:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v163 = m.ExcPending
	if v163 != 0 {
		goto L20
	} else {
		goto L50
	}
L20:
	;
	return int32(0)
L21:
	;
	v53 = F_type_is_multirange(m, v49)
	mBase = m.M
	v54 = m.ExcPending
	if v54 != 0 {
		goto L20
	} else {
		goto L22
	}
L22:
	;
	if v53 == int32(0) {
		goto L2
	} else {
		goto L23
	}
L23:
	;
	v57 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v58 = *(*int32)(unsafe.Add(mBase, uint32(v57)+16))
	if v58 != 0 {
		goto L25
	} else {
		goto L26
	}
L24:
	;
	v70 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v71 = F_pg_detoast_datum(m, v70)
	mBase = m.M
	v72 = m.ExcPending
	if v72 != 0 {
		goto L20
	} else {
		goto L31
	}
L25:
	;
	v59 = *(*int32)(unsafe.Add(mBase, uint32(v58)))
	if v59 == v49 {
		v69 = v58
		goto L24
	} else {
		goto L28
	}
L26:
	;
	goto L27
L27:
	;
	v62 = F_lookup_type_cache(m, v49, int32(_a_F_multirange_intersect_agg_transfn_0))
	mBase = m.M
	v63 = m.ExcPending
	if v63 != 0 {
		goto L20
	} else {
		goto L29
	}
L28:
	;
	goto L27
L29:
	;
	v64 = *(*int32)(unsafe.Add(mBase, uint32(v62)+296))
	if v64 == int32(0) {
		goto L1
	} else {
		goto L30
	}
L30:
	;
	v67 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	*(*int32)(unsafe.Add(mBase, uint32(v67)+16)) = v62
	v69 = v62
	goto L24
L31:
	;
	v73 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
	v74 = F_pg_detoast_datum(m, v73)
	mBase = m.M
	v75 = m.ExcPending
	if v75 != 0 {
		goto L20
	} else {
		goto L32
	}
L32:
	;
	v76 = *(*int32)(unsafe.Add(mBase, uint32(v69)+296))
	v77 = *(*int32)(unsafe.Add(mBase, uint32(v71)+8))
	if int32(0) < v77 {
		goto L33
	} else {
		goto L34
	}
L33:
	;
	v83 = F_palloc(m, v77<<(uint(int32(2))%32))
	mBase = m.M
	v84 = m.ExcPending
	if v84 != 0 {
		goto L20
	} else {
		goto L36
	}
L34:
	;
	v108 = v76
	v113 = v2
	goto L35
L35:
	;
	v115 = *(*int32)(unsafe.Add(mBase, uint32(v74)+8))
	if int32(0) < v115 {
		goto L41
	} else {
		goto L42
	}
L36:
	;
	v85 = int32(0)
	goto L37
L37:
	;
	v98 = F_multirange_get_range(m, v76, v71, v85)
	mBase = m.M
	v99 = m.ExcPending
	if v99 != 0 {
		goto L20
	} else {
		goto L39
	}
L38:
	;
	v104 = *(*int32)(unsafe.Add(mBase, uint32(v69)+296))
	v108 = v104
	v113 = v83
	goto L35
L39:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v83+v85<<(uint(int32(2))%32)))) = v98
	v102 = v85 + int32(1)
	if v102 != v77 {
		v85 = v102
		goto L37
	} else {
		goto L40
	}
L40:
	;
	goto L38
L41:
	;
	v121 = F_palloc(m, v115<<(uint(int32(2))%32))
	mBase = m.M
	v122 = m.ExcPending
	if v122 != 0 {
		goto L20
	} else {
		goto L44
	}
L42:
	;
	v152 = v2
	v153 = v108
	goto L43
L43:
	;
	v154 = F_multirange_intersect_internal(m, v49, v153, v77, v113, v115, v152)
	mBase = m.M
	v155 = m.ExcPending
	if v155 != 0 {
		goto L20
	} else {
		goto L49
	}
L44:
	;
	v123 = int32(0)
	goto L45
L45:
	;
	v136 = F_multirange_get_range(m, v108, v74, v123)
	mBase = m.M
	v137 = m.ExcPending
	if v137 != 0 {
		goto L20
	} else {
		goto L47
	}
L46:
	;
	v142 = *(*int32)(unsafe.Add(mBase, uint32(v69)+296))
	v152 = v121
	v153 = v142
	goto L43
L47:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v121+v123<<(uint(int32(2))%32)))) = v136
	v140 = v123 + int32(1)
	if v140 != v115 {
		v123 = v140
		goto L45
	} else {
		goto L48
	}
L48:
	;
	goto L46
L49:
	;
	m.G0 = v13 + int32(16)
	return v154
L50:
	;
	F_errmsg_internal(m, int32(_a_F_multirange_intersect_agg_transfn_1), int32(0))
	mBase = m.M
	v167 = m.ExcPending
	if v167 != 0 {
		goto L20
	} else {
		goto L51
	}
L51:
	;
	F_errfinish(m, int32(_a_F_multirange_intersect_agg_transfn_2), int32(1479), int32(_a_F_multirange_intersect_agg_transfn_3))
	mBase = m.M
	v172 = m.ExcPending
	if v172 != 0 {
		goto L20
	} else {
		goto L52
	}
L52:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L53:
	;
	F_errmsg_internal(m, int32(_a_F_multirange_intersect_agg_transfn_4), int32(0))
	mBase = m.M
	v180 = m.ExcPending
	if v180 != 0 {
		goto L20
	} else {
		goto L54
	}
L54:
	;
	F_errfinish(m, int32(_a_F_multirange_intersect_agg_transfn_2), int32(1483), int32(_a_F_multirange_intersect_agg_transfn_3))
	mBase = m.M
	v185 = m.ExcPending
	if v185 != 0 {
		goto L20
	} else {
		goto L55
	}
L55:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L56:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v13))) = v49
	F_errmsg_internal(m, int32(_a_F_multirange_intersect_agg_transfn_5), v13)
	mBase = m.M
	v193 = m.ExcPending
	if v193 != 0 {
		goto L20
	} else {
		goto L57
	}
L57:
	;
	F_errfinish(m, int32(_a_F_multirange_intersect_agg_transfn_2), int32(558), int32(_a_F_multirange_intersect_agg_transfn_6))
	mBase = m.M
	v198 = m.ExcPending
	if v198 != 0 {
		goto L20
	} else {
		goto L58
	}
L58:
	;
	base.Wasm_trap_unreachable()
	for {
	}
}
func F_multirange_upper(m *base.Module, l0 int32) int32 {
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
	var v43 int32
	_ = v43
	var v44 int32
	_ = v44
	var v47 int32
	_ = v47
	var v51 int32
	_ = v51
	var v57 int32
	_ = v57
	var v65 int32
	_ = v65
	var v69 int32
	_ = v69
	var v74 int32
	_ = v74
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
		if v16 == int32(0) {
			v51 = int32(1)
			*(*uint8)(unsafe.Add(mBase, uint32(l0)+16)) = uint8(v51)
			v57 = int32(0)
			m.G0 = v9 + int32(32)
			return v57
		} else {
			v19 = *(*int32)(unsafe.Add(mBase, uint32(v12)+4))
			v20 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
			v21 = *(*int32)(unsafe.Add(mBase, uint32(v20)+16))
			if v21 != 0 {
				v22 = *(*int32)(unsafe.Add(mBase, uint32(v21)))
				if v22 == v19 {
					v33 = v21
					v34 = v16
					v35 = *(*int32)(unsafe.Add(mBase, uint32(v33)+296))
					F_multirange_get_bounds(m, v35, v12, v34-int32(1), v9+int32(24), v9+int32(16))
					mBase = m.M
					v43 = m.ExcPending
					if v43 != 0 {
						return int32(0)
					} else {
						v44 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v9)+20)))
						if v44 == int32(0) {
							v47 = *(*int32)(unsafe.Add(mBase, uint32(v9)+16))
							v57 = v47
						} else {
							v51 = int32(1)
							*(*uint8)(unsafe.Add(mBase, uint32(l0)+16)) = uint8(v51)
							v57 = int32(0)
						}
						m.G0 = v9 + int32(32)
						return v57
					}
				} else {
					v25 = F_lookup_type_cache(m, v19, int32(_a_F_multirange_upper_0))
					mBase = m.M
					v26 = m.ExcPending
					if v26 != 0 {
						return int32(0)
					} else {
						v27 = *(*int32)(unsafe.Add(mBase, uint32(v25)+296))
						if v27 == int32(0) {
							F_errstart_cold(m, int32(21), int32(0))
							mBase = m.M
							v65 = m.ExcPending
							if v65 != 0 {
								return int32(0)
							} else {
								*(*int32)(unsafe.Add(mBase, uint32(v9))) = v19
								F_errmsg_internal(m, int32(_a_F_multirange_upper_1), v9)
								mBase = m.M
								v69 = m.ExcPending
								if v69 != 0 {
									return int32(0)
								} else {
									F_errfinish(m, int32(_a_F_multirange_upper_2), int32(558), int32(_a_F_multirange_upper_3))
									mBase = m.M
									v74 = m.ExcPending
									if v74 != 0 {
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
							v32 = *(*int32)(unsafe.Add(mBase, uint32(v12)+8))
							v33 = v25
							v34 = v32
							v35 = *(*int32)(unsafe.Add(mBase, uint32(v33)+296))
							F_multirange_get_bounds(m, v35, v12, v34-int32(1), v9+int32(24), v9+int32(16))
							mBase = m.M
							v43 = m.ExcPending
							if v43 != 0 {
								return int32(0)
							} else {
								v44 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v9)+20)))
								if v44 == int32(0) {
									v47 = *(*int32)(unsafe.Add(mBase, uint32(v9)+16))
									v57 = v47
								} else {
									v51 = int32(1)
									*(*uint8)(unsafe.Add(mBase, uint32(l0)+16)) = uint8(v51)
									v57 = int32(0)
								}
								m.G0 = v9 + int32(32)
								return v57
							}
						}
					}
				}
			} else {
				v25 = F_lookup_type_cache(m, v19, int32(_a_F_multirange_upper_0))
				mBase = m.M
				v26 = m.ExcPending
				if v26 != 0 {
					return int32(0)
				} else {
					v27 = *(*int32)(unsafe.Add(mBase, uint32(v25)+296))
					if v27 == int32(0) {
						F_errstart_cold(m, int32(21), int32(0))
						mBase = m.M
						v65 = m.ExcPending
						if v65 != 0 {
							return int32(0)
						} else {
							*(*int32)(unsafe.Add(mBase, uint32(v9))) = v19
							F_errmsg_internal(m, int32(_a_F_multirange_upper_1), v9)
							mBase = m.M
							v69 = m.ExcPending
							if v69 != 0 {
								return int32(0)
							} else {
								F_errfinish(m, int32(_a_F_multirange_upper_2), int32(558), int32(_a_F_multirange_upper_3))
								mBase = m.M
								v74 = m.ExcPending
								if v74 != 0 {
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
						v32 = *(*int32)(unsafe.Add(mBase, uint32(v12)+8))
						v33 = v25
						v34 = v32
						v35 = *(*int32)(unsafe.Add(mBase, uint32(v33)+296))
						F_multirange_get_bounds(m, v35, v12, v34-int32(1), v9+int32(24), v9+int32(16))
						mBase = m.M
						v43 = m.ExcPending
						if v43 != 0 {
							return int32(0)
						} else {
							v44 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v9)+20)))
							if v44 == int32(0) {
								v47 = *(*int32)(unsafe.Add(mBase, uint32(v9)+16))
								v57 = v47
							} else {
								v51 = int32(1)
								*(*uint8)(unsafe.Add(mBase, uint32(l0)+16)) = uint8(v51)
								v57 = int32(0)
							}
							m.G0 = v9 + int32(32)
							return v57
						}
					}
				}
			}
		}
	}
}
