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
	var v19 int32
	_ = v19
	var v23 int32
	_ = v23
	var v29 int32
	_ = v29
	var v30 int32
	_ = v30
	var v31 int32
	_ = v31
	var v37 int32
	_ = v37
	var v40 int32
	_ = v40
	var v41 int32
	_ = v41
	var v47 int32
	_ = v47
	var v56 int32
	_ = v56
	var v57 int32
	_ = v57
	var v62 int32
	_ = v62
	var v63 int32
	_ = v63
	var v65 int32
	_ = v65
	var v66 int32
	_ = v66
	var v75 int32
	_ = v75
	var v76 int32
	_ = v76
	var v80 int32
	_ = v80
	var v81 int32
	_ = v81
	var v83 int32
	_ = v83
	var v92 int32
	_ = v92
	var v99 int32
	_ = v99
	var v100 int32
	_ = v100
	var v103 int32
	_ = v103
	var v106 int32
	_ = v106
	var v111 int32
	_ = v111
	var v112 int32
	_ = v112
	var v127 int32
	_ = v127
	var v139 int32
	_ = v139
	var v141 int32
	_ = v141
	var v148 int32
	_ = v148
	var v151 int32
	_ = v151
	var v152 int32
	_ = v152
	var v154 int32
	_ = v154
	var v171 int32
	_ = v171
	var v172 int32
	_ = v172
	var v174 int32
	_ = v174
	var v180 int32
	_ = v180
	var v189 int32
	_ = v189
	var v190 int32
	_ = v190
	var v192 int32
	_ = v192
	var v196 int32
	_ = v196
	var v199 int32
	_ = v199
	var v206 int32
	_ = v206
	var v207 int32
	_ = v207
	var v227 int32
	_ = v227
	var v230 int32
	_ = v230
	var v231 int32
	_ = v231
	var v232 int32
	_ = v232
	var v233 int32
	_ = v233
	var v238 int32
	_ = v238
	var v240 int32
	_ = v240
	var v241 int32
	_ = v241
	var v244 int32
	_ = v244
	var v246 int32
	_ = v246
	var v248 int32
	_ = v248
	var v249 int32
	_ = v249
	var v250 int32
	_ = v250
	var v267 int32
	_ = v267
	var v281 int32
	_ = v281
	var v282 int32
	_ = v282
	var v283 int32
	_ = v283
	var v291 int32
	_ = v291
	var v293 int32
	_ = v293
	var v299 int32
	_ = v299
	var v302 int32
	_ = v302
	var v303 int32
	_ = v303
	var v304 int32
	_ = v304
	var v305 int32
	_ = v305
	var v310 int32
	_ = v310
	var v312 int32
	_ = v312
	var v315 int32
	_ = v315
	var v317 int32
	_ = v317
	var v319 int32
	_ = v319
	var v320 int32
	_ = v320
	var v321 int32
	_ = v321
	var v336 int32
	_ = v336
	var v339 int32
	_ = v339
	v5 = int32(0)
	if l3 != 0 {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	F_qsort_arg(m, l3, l2, int32(4), int32(1458), l1)
	mBase = m.M
	v19 = m.ExcPending
	if v19 != 0 {
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
	v23 = l3 - int32(4)
	v29 = v5
	v30 = v5
	v31 = v5
	goto L9
L7:
	;
	v92 = v5
	goto L8
L8:
	;
	v99 = v92 - int32(1)
	v100 = int32(0)
	if v100 < v99 {
		goto L29
	} else {
		goto L30
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
	v92 = v81
	goto L8
L11:
	;
	v83 = v30 + int32(1)
	if v83 != l2 {
		v29 = v80
		v30 = v83
		v31 = v81
		goto L9
	} else {
		goto L28
	}
L12:
	;
	if v47&int32(1) != 0 {
		v80 = v29
		v81 = v31
		goto L11
	} else {
		goto L13
	}
L13:
	;
	if v29 == int32(0) {
		goto L15
	} else {
		goto L16
	}
L14:
	;
	v80 = v40
	v81 = v31 + int32(1)
	goto L11
L15:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l3+v31<<(uint(int32(2))%32)))) = v40
	goto L14
L16:
	;
	goto L17
L17:
	;
	v56 = F_range_adjacent_internal(m, l1, v29, v40)
	mBase = m.M
	v57 = m.ExcPending
	if v57 != 0 {
		goto L4
	} else {
		goto L18
	}
L18:
	;
	if v56 != 0 {
		goto L19
	} else {
		goto L20
	}
L19:
	;
	v62 = F_range_union_internal(m, l1, v29, v40, int32(0))
	mBase = m.M
	v63 = m.ExcPending
	if v63 != 0 {
		goto L4
	} else {
		goto L22
	}
L20:
	;
	goto L21
L21:
	;
	v65 = F_range_before_internal(m, l1, v29, v40)
	mBase = m.M
	v66 = m.ExcPending
	if v66 != 0 {
		goto L4
	} else {
		goto L23
	}
L22:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v23+v31<<(uint(int32(2))%32)))) = v62
	v80 = v62
	v81 = v31
	goto L11
L23:
	;
	if v65 != 0 {
		goto L24
	} else {
		goto L25
	}
L24:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l3+v31<<(uint(int32(2))%32)))) = v40
	goto L14
L25:
	;
	goto L26
L26:
	;
	v75 = F_range_union_internal(m, l1, v29, v40, int32(1))
	mBase = m.M
	v76 = m.ExcPending
	if v76 != 0 {
		goto L4
	} else {
		goto L27
	}
L27:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v23+v31<<(uint(int32(2))%32)))) = v75
	v80 = v75
	v81 = v31
	goto L11
L28:
	;
	goto L10
L29:
	;
	v103 = v99
	goto L31
L30:
	;
	v103 = v100
	goto L31
L31:
	;
	v106 = v103<<(uint(int32(2))%32) + v92
	v111 = *(*int32)(unsafe.Add(mBase, uint32(l1)+200))
	v112 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v111)+11)))
	if v112 == int32(105) {
		v127 = (v106 + int32(15)) & int32(-4)
		goto L32
	} else {
		goto L33
	}
L32:
	;
	if int32(0) < v92 {
		goto L37
	} else {
		goto L38
	}
L33:
	;
	switch v112 - int32(99) {
	case 0:
		goto L36
	case 1:
		goto L35
	default:
		goto L34
	}
L34:
	;
	v127 = (v106 + int32(13)) & int32(-2)
	goto L32
L35:
	;
	v127 = (v106 + int32(19)) & int32(-8)
	goto L32
L36:
	;
	v127 = v106 + int32(12)
	goto L32
L37:
	;
	v139 = v127
	v141 = int32(0)
	goto L40
L38:
	;
	v180 = v127
	goto L39
L39:
	;
	v189 = F_palloc0(m, v180)
	mBase = m.M
	v190 = m.ExcPending
	if v190 != 0 {
		goto L4
	} else {
		goto L48
	}
L40:
	;
	v148 = int32(2)
	v151 = *(*int32)(unsafe.Add(mBase, uint32(l3+v141<<(uint(v148)%32))))
	v152 = *(*int32)(unsafe.Add(mBase, uint32(v151)))
	v154 = int32(base.Ui32(v152) >> (uint(v148) % 32))
	if base.B2i32(v112 != int32(105)) == int32(0) {
		v171 = (v154 - int32(6)) & int32(-4)
		goto L42
	} else {
		goto L43
	}
L41:
	;
	v180 = v172
	goto L39
L42:
	;
	v172 = v171 + v139
	v174 = v141 + int32(1)
	if v174 != v92 {
		v139 = v172
		v141 = v174
		goto L40
	} else {
		goto L47
	}
L43:
	;
	switch v112 - int32(99) {
	case 0:
		goto L46
	case 1:
		goto L45
	default:
		goto L44
	}
L44:
	;
	v171 = v154&int32(1073741822) - int32(8)
	goto L42
L45:
	;
	v171 = (v154 - int32(2)) & int32(-8)
	goto L42
L46:
	;
	v171 = v154 - int32(9)
	goto L42
L47:
	;
	goto L41
L48:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v189)+4)) = l0
	v192 = int32(2)
	*(*int32)(unsafe.Add(mBase, uint32(v189))) = v180 << (uint(v192) % 32)
	v196 = v189 + int32(8)
	*(*int32)(unsafe.Add(mBase, uint32(v196))) = v92
	v199 = v92 << (uint(v192) % 32)
	v206 = *(*int32)(unsafe.Add(mBase, uint32(l1)+200))
	v207 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v206)+11)))
	if v207 == int32(105) {
		v227 = (v92*int32(5) + int32(11)) & int32(-4)
		goto L49
	} else {
		goto L50
	}
L49:
	;
	if v92 <= int32(0) {
		goto L54
	} else {
		goto L55
	}
L50:
	;
	switch v207 - int32(99) {
	case 0:
		goto L53
	case 1:
		goto L52
	default:
		goto L51
	}
L51:
	;
	v227 = (v92*int32(5) + int32(9)) & int32(-2)
	goto L49
L52:
	;
	v227 = (v92*int32(5) + int32(15)) & int32(-8)
	goto L49
L53:
	;
	v227 = v199 + v92 + int32(8)
	goto L49
L54:
	;
	return v189
L55:
	;
	v230 = v199 + v196
	v231 = *(*int32)(unsafe.Add(mBase, uint32(l3)))
	v232 = *(*int32)(unsafe.Add(mBase, uint32(v231)))
	v233 = int32(2)
	v238 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v231+int32(base.Ui32(v232)>>(uint(v233)%32))-int32(1)))))
	*(*uint8)(unsafe.Add(mBase, uint32(v230))) = uint8(v238)
	v240 = v227 + v189
	v241 = *(*int32)(unsafe.Add(mBase, uint32(l3)))
	v244 = *(*int32)(unsafe.Add(mBase, uint32(v241)))
	v246 = int32(base.Ui32(v244) >> (uint(v233) % 32))
	v248 = v246 - int32(9)
	if v248 != 0 {
		goto L57
	} else {
		goto L58
	}
L56:
	;
	if v207 != int32(105) {
		goto L61
	} else {
		goto L62
	}
L57:
	;
	v249 = F__emscripten_memcpy_bulkmem(m, v240, v241+int32(8), v248)
	mBase = m.M
	v250 = v249
	goto L59
L58:
	;
	v250 = v240
	goto L59
L59:
	;
	goto L56
L60:
	;
	if v92 == int32(1) {
		goto L54
	} else {
		goto L66
	}
L61:
	;
	switch v207 - int32(99) {
	case 0:
		v267 = v248
		goto L60
	case 1:
		goto L65
	default:
		goto L64
	}
L62:
	;
	goto L63
L63:
	;
	v267 = (v246 - int32(6)) & int32(-4)
	goto L60
L64:
	;
	v267 = v246&int32(1073741822) - int32(8)
	goto L60
L65:
	;
	v267 = (v246 - int32(2)) & int32(-8)
	goto L60
L66:
	;
	v281 = v250 + v267
	v282 = int32(0)
	v283 = int32(1)
	goto L67
L67:
	;
	v291 = v283 << (uint(int32(2)) % 32)
	v293 = v281 - v250
	if v283&int32(3) != 0 {
		goto L69
	} else {
		goto L70
	}
L68:
	;
	goto L54
L69:
	;
	v299 = v293 - v282
	goto L71
L70:
	;
	v299 = v293 | int32(-2147483648)
	goto L71
L71:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v196+v291))) = v299
	v302 = v291 + l3
	v303 = *(*int32)(unsafe.Add(mBase, uint32(v302)))
	v304 = *(*int32)(unsafe.Add(mBase, uint32(v303)))
	v305 = int32(2)
	v310 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v303+int32(base.Ui32(v304)>>(uint(v305)%32))-int32(1)))))
	*(*uint8)(unsafe.Add(mBase, uint32(v283+v230))) = uint8(v310)
	v312 = *(*int32)(unsafe.Add(mBase, uint32(v302)))
	v315 = *(*int32)(unsafe.Add(mBase, uint32(v312)))
	v317 = int32(base.Ui32(v315) >> (uint(v305) % 32))
	v319 = v317 - int32(9)
	if v319 != 0 {
		goto L73
	} else {
		goto L74
	}
L72:
	;
	if base.B2i32(v207 != int32(105)) == int32(0) {
		goto L77
	} else {
		goto L78
	}
L73:
	;
	v320 = F__emscripten_memcpy_bulkmem(m, v281, v312+int32(8), v319)
	mBase = m.M
	v321 = v320
	goto L75
L74:
	;
	v321 = v281
	goto L75
L75:
	;
	goto L72
L76:
	;
	v339 = v283 + int32(1)
	if v339 != v92 {
		v281 = v336 + v321
		v282 = v293
		v283 = v339
		goto L67
	} else {
		goto L82
	}
L77:
	;
	v336 = (v317 - int32(6)) & int32(-4)
	goto L76
L78:
	;
	goto L79
L79:
	;
	switch v207 - int32(99) {
	case 0:
		v336 = v319
		goto L76
	case 1:
		goto L81
	default:
		goto L80
	}
L80:
	;
	v336 = v317&int32(1073741822) - int32(8)
	goto L76
L81:
	;
	v336 = (v317 - int32(2)) & int32(-8)
	goto L76
L82:
	;
	goto L68
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
					v34 = *(*int32)(unsafe.Add(mBase, uint32(v12)+8))
					if v34 == v33 {
						v64 = v33
						m.G0 = v9 + int32(48)
						return v64
					} else {
						v37 = *(*int32)(unsafe.Add(mBase, uint32(v17)+8))
						if v37 == int32(0) {
							v64 = v33
							m.G0 = v9 + int32(48)
							return v64
						} else {
							v40 = *(*int32)(unsafe.Add(mBase, uint32(v32)+296))
							F_multirange_get_bounds(m, v40, v12, v34-int32(1), v9+int32(40), v9+int32(32))
							mBase = m.M
							v48 = m.ExcPending
							if v48 != 0 {
								return int32(0)
							} else {
								F_multirange_get_bounds(m, v40, v17, int32(0), v9+int32(24), v9+int32(16))
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
								F_errmsg_internal(m, int32(385310), v9)
								mBase = m.M
								v76 = m.ExcPending
								if v76 != 0 {
									return int32(0)
								} else {
									F_errfinish(m, int32(513647), int32(558), int32(414618))
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
							v34 = *(*int32)(unsafe.Add(mBase, uint32(v12)+8))
							if v34 == v33 {
								v64 = v33
								m.G0 = v9 + int32(48)
								return v64
							} else {
								v37 = *(*int32)(unsafe.Add(mBase, uint32(v17)+8))
								if v37 == int32(0) {
									v64 = v33
									m.G0 = v9 + int32(48)
									return v64
								} else {
									v40 = *(*int32)(unsafe.Add(mBase, uint32(v32)+296))
									F_multirange_get_bounds(m, v40, v12, v34-int32(1), v9+int32(40), v9+int32(32))
									mBase = m.M
									v48 = m.ExcPending
									if v48 != 0 {
										return int32(0)
									} else {
										F_multirange_get_bounds(m, v40, v17, int32(0), v9+int32(24), v9+int32(16))
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
							F_errmsg_internal(m, int32(385310), v9)
							mBase = m.M
							v76 = m.ExcPending
							if v76 != 0 {
								return int32(0)
							} else {
								F_errfinish(m, int32(513647), int32(558), int32(414618))
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
						v34 = *(*int32)(unsafe.Add(mBase, uint32(v12)+8))
						if v34 == v33 {
							v64 = v33
							m.G0 = v9 + int32(48)
							return v64
						} else {
							v37 = *(*int32)(unsafe.Add(mBase, uint32(v17)+8))
							if v37 == int32(0) {
								v64 = v33
								m.G0 = v9 + int32(48)
								return v64
							} else {
								v40 = *(*int32)(unsafe.Add(mBase, uint32(v32)+296))
								F_multirange_get_bounds(m, v40, v12, v34-int32(1), v9+int32(40), v9+int32(32))
								mBase = m.M
								v48 = m.ExcPending
								if v48 != 0 {
									return int32(0)
								} else {
									F_multirange_get_bounds(m, v40, v17, int32(0), v9+int32(24), v9+int32(16))
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
					v22 = F_lookup_type_cache(m, v13, int32(65536))
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
								F_errmsg_internal(m, int32(385310), v7)
								mBase = m.M
								v59 = m.ExcPending
								if v59 != 0 {
									return int32(0)
								} else {
									F_errfinish(m, int32(513647), int32(558), int32(414618))
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
				v22 = F_lookup_type_cache(m, v13, int32(65536))
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
							F_errmsg_internal(m, int32(385310), v7)
							mBase = m.M
							v59 = m.ExcPending
							if v59 != 0 {
								return int32(0)
							} else {
								F_errfinish(m, int32(513647), int32(558), int32(414618))
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
			F_errmsg_internal(m, int32(128683), int32(0))
			mBase = m.M
			v46 = m.ExcPending
			if v46 != 0 {
				return int32(0)
			} else {
				F_errfinish(m, int32(513647), int32(1069), int32(584530))
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
							v43 = m.ExcPending
							if v43 != 0 {
								return int32(0)
							} else {
								*(*int32)(unsafe.Add(mBase, uint32(v9))) = v19
								F_errmsg_internal(m, int32(385310), v9)
								mBase = m.M
								v47 = m.ExcPending
								if v47 != 0 {
									return int32(0)
								} else {
									F_errfinish(m, int32(513647), int32(558), int32(414618))
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
						v43 = m.ExcPending
						if v43 != 0 {
							return int32(0)
						} else {
							*(*int32)(unsafe.Add(mBase, uint32(v9))) = v19
							F_errmsg_internal(m, int32(385310), v9)
							mBase = m.M
							v47 = m.ExcPending
							if v47 != 0 {
								return int32(0)
							} else {
								F_errfinish(m, int32(513647), int32(558), int32(414618))
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
	var v6 int32
	_ = v6
	var v7 int32
	_ = v7
	var v11 int32
	_ = v11
	var v12 int32
	_ = v12
	var v15 int32
	_ = v15
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
	var v25 int32
	_ = v25
	var v27 int32
	_ = v27
	var v30 int32
	_ = v30
	var v31 int32
	_ = v31
	var v38 int32
	_ = v38
	var v39 int32
	_ = v39
	var v47 int32
	_ = v47
	var v52 int32
	_ = v52
	var v54 int32
	_ = v54
	var v55 int32
	_ = v55
	var v56 int32
	_ = v56
	var v61 int32
	_ = v61
	var v63 int32
	_ = v63
	var v65 int32
	_ = v65
	var v66 int32
	_ = v66
	v6 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v7 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v6)+14)))
	if v7 != int32(1) {
		return v6
	} else {
		v11 = *(*int32)(unsafe.Add(mBase, uint32(v6)))
		v12 = F_pg_detoast_datum(m, v11)
		mBase = m.M
		v15 = m.ExcPending
		if v15 != 0 {
			return int32(0)
		} else {
			v17 = F_palloc(m, int32(16))
			mBase = m.M
			v18 = m.ExcPending
			if v18 != 0 {
				return int32(0)
			} else {
				v19 = *(*int32)(unsafe.Add(mBase, uint32(v12)+4))
				v20 = F_multirange_get_typcache(m, l0, v19)
				mBase = m.M
				v21 = m.ExcPending
				if v21 != 0 {
					return int32(0)
				} else {
					v22 = *(*int32)(unsafe.Add(mBase, uint32(v20)+296))
					v23 = m.G0
					v25 = v23 - int32(32)
					m.G0 = v25
					v27 = *(*int32)(unsafe.Add(mBase, uint32(v12)+8))
					if v27 == int32(0) {
						v30 = F_make_empty_range(m, v22)
						mBase = m.M
						v31 = m.ExcPending
						if v31 != 0 {
							return int32(0)
						} else {
							v56 = v30
							m.G0 = v25 + int32(32)
							*(*int32)(unsafe.Add(mBase, uint32(v17))) = v56
							v61 = *(*int32)(unsafe.Add(mBase, uint32(v6)+4))
							*(*int32)(unsafe.Add(mBase, uint32(v17)+4)) = v61
							v63 = *(*int32)(unsafe.Add(mBase, uint32(v6)+8))
							*(*int32)(unsafe.Add(mBase, uint32(v17)+8)) = v63
							v65 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v6)+12)))
							v66 = int32(0)
							*(*uint8)(unsafe.Add(mBase, uint32(v17)+14)) = uint8(v66)
							*(*uint16)(unsafe.Add(mBase, uint32(v17)+12)) = uint16(v65)
							return v17
						}
					} else {
						F_multirange_get_bounds(m, v22, v12, int32(0), v25+int32(24), v25+int32(8))
						mBase = m.M
						v38 = m.ExcPending
						if v38 != 0 {
							return int32(0)
						} else {
							v39 = *(*int32)(unsafe.Add(mBase, uint32(v12)+8))
							F_multirange_get_bounds(m, v22, v12, v39-int32(1), v25+int32(8), v25+int32(16))
							mBase = m.M
							v47 = m.ExcPending
							if v47 != 0 {
								return int32(0)
							} else {
								v52 = int32(0)
								v54 = F_make_range(m, v22, v25+int32(24), v25+int32(16), v52, v52)
								mBase = m.M
								v55 = m.ExcPending
								if v55 != 0 {
									return int32(0)
								} else {
									v56 = v54
									m.G0 = v25 + int32(32)
									*(*int32)(unsafe.Add(mBase, uint32(v17))) = v56
									v61 = *(*int32)(unsafe.Add(mBase, uint32(v6)+4))
									*(*int32)(unsafe.Add(mBase, uint32(v17)+4)) = v61
									v63 = *(*int32)(unsafe.Add(mBase, uint32(v6)+8))
									*(*int32)(unsafe.Add(mBase, uint32(v17)+8)) = v63
									v65 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v6)+12)))
									v66 = int32(0)
									*(*uint8)(unsafe.Add(mBase, uint32(v17)+14)) = uint8(v66)
									*(*uint16)(unsafe.Add(mBase, uint32(v17)+12)) = uint16(v65)
									return v17
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
				if v13 != int32(4537) {
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
									F_errmsg_internal(m, int32(498947), v11+int32(16))
									mBase = m.M
									v57 = m.ExcPending
									if v57 != 0 {
										return int32(0)
									} else {
										F_errfinish(m, int32(511463), int32(1138), int32(101916))
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
				if v13 != int32(4537) {
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
									F_errmsg_internal(m, int32(498947), v11)
									mBase = m.M
									v91 = m.ExcPending
									if v91 != 0 {
										return int32(0)
									} else {
										F_errfinish(m, int32(511463), int32(1049), int32(101846))
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
	var v110 int32
	_ = v110
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
	var v149 int32
	_ = v149
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
		goto L57
	}
L2:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v176 = m.ExcPending
	if v176 != 0 {
		goto L20
	} else {
		goto L54
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
		goto L51
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
	v62 = F_lookup_type_cache(m, v49, int32(65536))
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
	if v77 <= int32(0) {
		goto L34
	} else {
		goto L35
	}
L33:
	;
	v115 = *(*int32)(unsafe.Add(mBase, uint32(v74)+8))
	if int32(0) < v115 {
		goto L42
	} else {
		goto L43
	}
L34:
	;
	v110 = v76
	v113 = v2
	goto L33
L35:
	;
	goto L36
L36:
	;
	v83 = F_palloc(m, v77<<(uint(int32(2))%32))
	mBase = m.M
	v84 = m.ExcPending
	if v84 != 0 {
		goto L20
	} else {
		goto L37
	}
L37:
	;
	v85 = int32(0)
	goto L38
L38:
	;
	v98 = F_multirange_get_range(m, v76, v71, v85)
	mBase = m.M
	v99 = m.ExcPending
	if v99 != 0 {
		goto L20
	} else {
		goto L40
	}
L39:
	;
	v104 = *(*int32)(unsafe.Add(mBase, uint32(v69)+296))
	v110 = v104
	v113 = v83
	goto L33
L40:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v83+v85<<(uint(int32(2))%32)))) = v98
	v102 = v85 + int32(1)
	if v102 != v77 {
		v85 = v102
		goto L38
	} else {
		goto L41
	}
L41:
	;
	goto L39
L42:
	;
	v121 = F_palloc(m, v115<<(uint(int32(2))%32))
	mBase = m.M
	v122 = m.ExcPending
	if v122 != 0 {
		goto L20
	} else {
		goto L45
	}
L43:
	;
	v149 = v2
	v153 = v110
	goto L44
L44:
	;
	v154 = F_multirange_intersect_internal(m, v49, v153, v77, v113, v115, v149)
	mBase = m.M
	v155 = m.ExcPending
	if v155 != 0 {
		goto L20
	} else {
		goto L50
	}
L45:
	;
	v123 = int32(0)
	goto L46
L46:
	;
	v136 = F_multirange_get_range(m, v110, v74, v123)
	mBase = m.M
	v137 = m.ExcPending
	if v137 != 0 {
		goto L20
	} else {
		goto L48
	}
L47:
	;
	v142 = *(*int32)(unsafe.Add(mBase, uint32(v69)+296))
	v149 = v121
	v153 = v142
	goto L44
L48:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v121+v123<<(uint(int32(2))%32)))) = v136
	v140 = v123 + int32(1)
	if v140 != v115 {
		v123 = v140
		goto L46
	} else {
		goto L49
	}
L49:
	;
	goto L47
L50:
	;
	m.G0 = v13 + int32(16)
	return v154
L51:
	;
	F_errmsg_internal(m, int32(66442), int32(0))
	mBase = m.M
	v167 = m.ExcPending
	if v167 != 0 {
		goto L20
	} else {
		goto L52
	}
L52:
	;
	F_errfinish(m, int32(513647), int32(1479), int32(291663))
	mBase = m.M
	v172 = m.ExcPending
	if v172 != 0 {
		goto L20
	} else {
		goto L53
	}
L53:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L54:
	;
	F_errmsg_internal(m, int32(416944), int32(0))
	mBase = m.M
	v180 = m.ExcPending
	if v180 != 0 {
		goto L20
	} else {
		goto L55
	}
L55:
	;
	F_errfinish(m, int32(513647), int32(1483), int32(291663))
	mBase = m.M
	v185 = m.ExcPending
	if v185 != 0 {
		goto L20
	} else {
		goto L56
	}
L56:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L57:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v13))) = v49
	F_errmsg_internal(m, int32(385310), v13)
	mBase = m.M
	v193 = m.ExcPending
	if v193 != 0 {
		goto L20
	} else {
		goto L58
	}
L58:
	;
	F_errfinish(m, int32(513647), int32(558), int32(414618))
	mBase = m.M
	v198 = m.ExcPending
	if v198 != 0 {
		goto L20
	} else {
		goto L59
	}
L59:
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
	var v21 int32
	_ = v21
	var v22 int32
	_ = v22
	var v23 int32
	_ = v23
	var v24 int32
	_ = v24
	var v27 int32
	_ = v27
	var v28 int32
	_ = v28
	var v29 int32
	_ = v29
	var v32 int32
	_ = v32
	var v34 int32
	_ = v34
	var v35 int32
	_ = v35
	var v36 int32
	_ = v36
	var v37 int32
	_ = v37
	var v45 int32
	_ = v45
	var v46 int32
	_ = v46
	var v49 int32
	_ = v49
	var v50 int32
	_ = v50
	var v59 int32
	_ = v59
	var v67 int32
	_ = v67
	var v71 int32
	_ = v71
	var v76 int32
	_ = v76
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
			v19 = int32(1)
			*(*uint8)(unsafe.Add(mBase, uint32(l0)+16)) = uint8(v19)
			v59 = int32(0)
			m.G0 = v9 + int32(32)
			return v59
		} else {
			v21 = *(*int32)(unsafe.Add(mBase, uint32(v12)+4))
			v22 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
			v23 = *(*int32)(unsafe.Add(mBase, uint32(v22)+16))
			if v23 != 0 {
				v24 = *(*int32)(unsafe.Add(mBase, uint32(v23)))
				if v24 == v21 {
					v35 = v23
					v36 = v16
					v37 = *(*int32)(unsafe.Add(mBase, uint32(v35)+296))
					F_multirange_get_bounds(m, v37, v12, v36-int32(1), v9+int32(24), v9+int32(16))
					mBase = m.M
					v45 = m.ExcPending
					if v45 != 0 {
						return int32(0)
					} else {
						v46 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v9)+20)))
						if v46 == int32(0) {
							v49 = *(*int32)(unsafe.Add(mBase, uint32(v9)+16))
							v59 = v49
						} else {
							v50 = int32(1)
							*(*uint8)(unsafe.Add(mBase, uint32(l0)+16)) = uint8(v50)
							v59 = int32(0)
						}
						m.G0 = v9 + int32(32)
						return v59
					}
				} else {
					v27 = F_lookup_type_cache(m, v21, int32(65536))
					mBase = m.M
					v28 = m.ExcPending
					if v28 != 0 {
						return int32(0)
					} else {
						v29 = *(*int32)(unsafe.Add(mBase, uint32(v27)+296))
						if v29 == int32(0) {
							F_errstart_cold(m, int32(21), int32(0))
							mBase = m.M
							v67 = m.ExcPending
							if v67 != 0 {
								return int32(0)
							} else {
								*(*int32)(unsafe.Add(mBase, uint32(v9))) = v21
								F_errmsg_internal(m, int32(385310), v9)
								mBase = m.M
								v71 = m.ExcPending
								if v71 != 0 {
									return int32(0)
								} else {
									F_errfinish(m, int32(513647), int32(558), int32(414618))
									mBase = m.M
									v76 = m.ExcPending
									if v76 != 0 {
										return int32(0)
									} else {
										base.Wasm_trap_unreachable()
										for {
										}
									}
								}
							}
						} else {
							v32 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
							*(*int32)(unsafe.Add(mBase, uint32(v32)+16)) = v27
							v34 = *(*int32)(unsafe.Add(mBase, uint32(v12)+8))
							v35 = v27
							v36 = v34
							v37 = *(*int32)(unsafe.Add(mBase, uint32(v35)+296))
							F_multirange_get_bounds(m, v37, v12, v36-int32(1), v9+int32(24), v9+int32(16))
							mBase = m.M
							v45 = m.ExcPending
							if v45 != 0 {
								return int32(0)
							} else {
								v46 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v9)+20)))
								if v46 == int32(0) {
									v49 = *(*int32)(unsafe.Add(mBase, uint32(v9)+16))
									v59 = v49
								} else {
									v50 = int32(1)
									*(*uint8)(unsafe.Add(mBase, uint32(l0)+16)) = uint8(v50)
									v59 = int32(0)
								}
								m.G0 = v9 + int32(32)
								return v59
							}
						}
					}
				}
			} else {
				v27 = F_lookup_type_cache(m, v21, int32(65536))
				mBase = m.M
				v28 = m.ExcPending
				if v28 != 0 {
					return int32(0)
				} else {
					v29 = *(*int32)(unsafe.Add(mBase, uint32(v27)+296))
					if v29 == int32(0) {
						F_errstart_cold(m, int32(21), int32(0))
						mBase = m.M
						v67 = m.ExcPending
						if v67 != 0 {
							return int32(0)
						} else {
							*(*int32)(unsafe.Add(mBase, uint32(v9))) = v21
							F_errmsg_internal(m, int32(385310), v9)
							mBase = m.M
							v71 = m.ExcPending
							if v71 != 0 {
								return int32(0)
							} else {
								F_errfinish(m, int32(513647), int32(558), int32(414618))
								mBase = m.M
								v76 = m.ExcPending
								if v76 != 0 {
									return int32(0)
								} else {
									base.Wasm_trap_unreachable()
									for {
									}
								}
							}
						}
					} else {
						v32 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
						*(*int32)(unsafe.Add(mBase, uint32(v32)+16)) = v27
						v34 = *(*int32)(unsafe.Add(mBase, uint32(v12)+8))
						v35 = v27
						v36 = v34
						v37 = *(*int32)(unsafe.Add(mBase, uint32(v35)+296))
						F_multirange_get_bounds(m, v37, v12, v36-int32(1), v9+int32(24), v9+int32(16))
						mBase = m.M
						v45 = m.ExcPending
						if v45 != 0 {
							return int32(0)
						} else {
							v46 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v9)+20)))
							if v46 == int32(0) {
								v49 = *(*int32)(unsafe.Add(mBase, uint32(v9)+16))
								v59 = v49
							} else {
								v50 = int32(1)
								*(*uint8)(unsafe.Add(mBase, uint32(l0)+16)) = uint8(v50)
								v59 = int32(0)
							}
							m.G0 = v9 + int32(32)
							return v59
						}
					}
				}
			}
		}
	}
}
