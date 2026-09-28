package p1

import (
	base "github.com/shibukawa/pgmem/internal/aot/pgaot/base"
	"unsafe"
)

func F_get_multirange_range(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v5 int32
	_ = v5
	var v8 int32
	_ = v8
	var v13 int32
	_ = v13
	var v14 int32
	_ = v14
	var v16 int32
	_ = v16
	var v18 int32
	_ = v18
	v5 = F_SearchSysCache1(m, int32(54), base.I64_extend_i32_u(l0))
	mBase = m.M
	v8 = m.ExcPending
	if v8 != 0 {
		return int32(0)
	} else {
		if v5 == int32(0) {
			return int32(0)
		} else {
			v13 = *(*int32)(unsafe.Add(mBase, uint32(v5)+16))
			v14 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v13)+22)))
			v16 = *(*int32)(unsafe.Add(mBase, uint32(v13+v14)))
			F_ReleaseCatCache(m, v5)
			mBase = m.M
			v18 = m.ExcPending
			if v18 != 0 {
				return int32(0)
			} else {
				return v16
			}
		}
	}
}
func F_make_multirange(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) int32 {
	mBase := m.M
	_ = mBase
	var v5 int32
	_ = v5
	var v15 int32
	_ = v15
	var v17 int32
	_ = v17
	var v24 int32
	_ = v24
	var v33 int32
	_ = v33
	var v34 int32
	_ = v34
	var v35 int32
	_ = v35
	var v42 int32
	_ = v42
	var v45 int32
	_ = v45
	var v46 int32
	_ = v46
	var v52 int32
	_ = v52
	var v57 int32
	_ = v57
	var v58 int32
	_ = v58
	var v65 int32
	_ = v65
	var v66 int32
	_ = v66
	var v68 int32
	_ = v68
	var v69 int32
	_ = v69
	var v76 int32
	_ = v76
	var v77 int32
	_ = v77
	var v85 int32
	_ = v85
	var v86 int32
	_ = v86
	var v88 int32
	_ = v88
	var v97 int32
	_ = v97
	var v104 int32
	_ = v104
	var v105 int32
	_ = v105
	var v113 int32
	_ = v113
	var v118 int32
	_ = v118
	var v123 int32
	_ = v123
	var v125 int32
	_ = v125
	var v126 int32
	_ = v126
	var v128 int32
	_ = v128
	var v130 int32
	_ = v130
	var v134 int32
	_ = v134
	var v141 int32
	_ = v141
	var v146 int32
	_ = v146
	var v158 int32
	_ = v158
	var v159 int32
	_ = v159
	var v162 int32
	_ = v162
	var v168 int32
	_ = v168
	var v170 int32
	_ = v170
	var v171 int32
	_ = v171
	var v172 int32
	_ = v172
	var v177 int32
	_ = v177
	var v178 int32
	_ = v178
	var v184 int32
	_ = v184
	var v186 int32
	_ = v186
	var v188 int32
	_ = v188
	var v196 int32
	_ = v196
	var v197 int32
	_ = v197
	var v206 int32
	_ = v206
	var v209 int32
	_ = v209
	var v210 int32
	_ = v210
	var v220 int32
	_ = v220
	var v230 int32
	_ = v230
	var v231 int32
	_ = v231
	var v238 int32
	_ = v238
	var v239 int32
	_ = v239
	var v245 int32
	_ = v245
	var v252 int32
	_ = v252
	var v257 int32
	_ = v257
	var v264 int32
	_ = v264
	var v265 int32
	_ = v265
	var v268 int32
	_ = v268
	var v272 int32
	_ = v272
	var v273 int32
	_ = v273
	var v274 int32
	_ = v274
	var v280 int32
	_ = v280
	var v283 int32
	_ = v283
	var v288 int32
	_ = v288
	var v289 int32
	_ = v289
	var v290 int32
	_ = v290
	var v294 int32
	_ = v294
	var v300 int32
	_ = v300
	var v301 int32
	_ = v301
	var v309 int32
	_ = v309
	var v311 int32
	_ = v311
	var v312 int32
	_ = v312
	var v322 int32
	_ = v322
	var v324 int32
	_ = v324
	var v330 int32
	_ = v330
	var v333 int32
	_ = v333
	var v334 int32
	_ = v334
	var v335 int32
	_ = v335
	var v336 int32
	_ = v336
	var v341 int32
	_ = v341
	var v343 int32
	_ = v343
	var v344 int32
	_ = v344
	var v348 int32
	_ = v348
	var v356 int32
	_ = v356
	v5 = int32(0)
	v15 = m.G0
	v17 = v15 - int32(32)
	m.G0 = v17
	if l3 != 0 {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	F_qsort_arg(m, l3, l2, int32(4), int32(1589), l1)
	mBase = m.M
	v24 = m.ExcPending
	if v24 != 0 {
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
	v33 = v5
	v34 = v5
	v35 = v5
	goto L9
L7:
	;
	v97 = v5
	goto L8
L8:
	;
	v104 = *(*int32)(unsafe.Add(mBase, uint32(l1)+200))
	v105 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v104)+11)))
	switch v105 - int32(99) {
	case 0:
		v125 = int32(1)
		goto L29
	case 1:
		goto L32
	default:
		goto L31
	case 6:
		goto L33
	case 16:
		goto L30
	}
L9:
	;
	v42 = int32(2)
	v45 = *(*int32)(unsafe.Add(mBase, uint32(l3+v33<<(uint(v42)%32))))
	v46 = *(*int32)(unsafe.Add(mBase, uint32(v45)))
	v52 = int32(*(*int8)(unsafe.Add(mBase, uint32(v45+int32(base.Ui32(v46)>>(uint(v42)%32))-int32(1)))))
	goto L12
L10:
	;
	v97 = v86
	goto L8
L11:
	;
	v88 = v33 + int32(1)
	if v88 != l2 {
		v33 = v88
		v34 = v85
		v35 = v86
		goto L9
	} else {
		goto L28
	}
L12:
	;
	if v52&int32(1) != 0 {
		v85 = v34
		v86 = v35
		goto L11
	} else {
		goto L13
	}
L13:
	;
	if v34 == int32(0) {
		goto L15
	} else {
		goto L16
	}
L14:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l3+v35<<(uint(int32(2))%32)))) = v45
	v85 = v45
	v86 = v35 + int32(1)
	goto L11
L15:
	;
	goto L14
L16:
	;
	goto L17
L17:
	;
	v57 = F_range_adjacent_internal(m, l1, v34, v45)
	mBase = m.M
	v58 = m.ExcPending
	if v58 != 0 {
		goto L4
	} else {
		goto L18
	}
L18:
	;
	if v57 != 0 {
		goto L19
	} else {
		goto L20
	}
L19:
	;
	v65 = F_range_union_internal(m, l1, v34, v45, int32(0))
	mBase = m.M
	v66 = m.ExcPending
	if v66 != 0 {
		goto L4
	} else {
		goto L22
	}
L20:
	;
	goto L21
L21:
	;
	v68 = F_range_before_internal(m, l1, v34, v45)
	mBase = m.M
	v69 = m.ExcPending
	if v69 != 0 {
		goto L4
	} else {
		goto L23
	}
L22:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l3+v35<<(uint(int32(2))%32)-int32(4)))) = v65
	v85 = v65
	v86 = v35
	goto L11
L23:
	;
	if v68 != 0 {
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
	v76 = F_range_union_internal(m, l1, v34, v45, int32(1))
	mBase = m.M
	v77 = m.ExcPending
	if v77 != 0 {
		goto L4
	} else {
		goto L27
	}
L27:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l3+v35<<(uint(int32(2))%32)-int32(4)))) = v76
	v85 = v76
	v86 = v35
	goto L11
L28:
	;
	goto L10
L29:
	;
	v126 = int32(0)
	v128 = v126 - v125
	v130 = v97 - int32(1)
	if v126 < v130 {
		goto L37
	} else {
		goto L38
	}
L30:
	;
	v125 = int32(2)
	goto L29
L31:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v113 = m.ExcPending
	if v113 != 0 {
		goto L4
	} else {
		goto L34
	}
L32:
	;
	v125 = int32(8)
	goto L29
L33:
	;
	v125 = int32(4)
	goto L29
L34:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v17))) = base.I32_extend8_s(v105)
	F_errmsg_internal(m, int32(_a_F_make_multirange_0), v17)
	mBase = m.M
	v118 = m.ExcPending
	if v118 != 0 {
		goto L4
	} else {
		goto L35
	}
L35:
	;
	F_errfinish(m, int32(_a_F_make_multirange_1), int32(322), int32(_a_F_make_multirange_2))
	mBase = m.M
	v123 = m.ExcPending
	if v123 != 0 {
		goto L4
	} else {
		goto L36
	}
L36:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L37:
	;
	v134 = v130
	goto L39
L38:
	;
	v134 = v126
	goto L39
L39:
	;
	v141 = v128 & (v97 + v134<<(uint(int32(2))%32) + v125 + int32(11))
	if v97 <= int32(0) {
		v220 = v141
		goto L40
	} else {
		goto L41
	}
L40:
	;
	v230 = F_palloc0(m, v220)
	mBase = m.M
	v231 = m.ExcPending
	if v231 != 0 {
		goto L4
	} else {
		goto L49
	}
L41:
	;
	v146 = v125 - int32(10)
	if v97 != int32(1) {
		goto L42
	} else {
		goto L43
	}
L42:
	;
	v158 = v141
	v159 = v126
	v162 = int32(0)
	goto L45
L43:
	;
	v196 = v141
	v197 = v126
	goto L44
L44:
	;
	v206 = int32(2)
	v209 = *(*int32)(unsafe.Add(mBase, uint32(l3+v197<<(uint(v206)%32))))
	v210 = *(*int32)(unsafe.Add(mBase, uint32(v209)))
	v220 = (v146+int32(base.Ui32(v210)>>(uint(v206)%32)))&v128 + v196
	goto L40
L45:
	;
	v168 = int32(2)
	v170 = l3 + v159<<(uint(v168)%32)
	v171 = *(*int32)(unsafe.Add(mBase, uint32(v170)+4))
	v172 = *(*int32)(unsafe.Add(mBase, uint32(v171)))
	v177 = *(*int32)(unsafe.Add(mBase, uint32(v170)))
	v178 = *(*int32)(unsafe.Add(mBase, uint32(v177)))
	v184 = (v146+int32(base.Ui32(v172)>>(uint(v168)%32)))&v128 + ((v146+int32(base.Ui32(v178)>>(uint(v168)%32)))&v128 + v158)
	v186 = v159 + v168
	v188 = v162 + v168
	if v188 != v97&int32(2147483646) {
		v158 = v184
		v159 = v186
		v162 = v188
		goto L45
	} else {
		goto L47
	}
L46:
	;
	if v97&int32(1) == int32(0) {
		v220 = v184
		goto L40
	} else {
		goto L48
	}
L47:
	;
	goto L46
L48:
	;
	v196 = v184
	v197 = v186
	goto L44
L49:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v230)+8)) = v97
	*(*int32)(unsafe.Add(mBase, uint32(v230)+4)) = l0
	*(*int32)(unsafe.Add(mBase, uint32(v230))) = v220 << (uint(int32(2)) % 32)
	v238 = *(*int32)(unsafe.Add(mBase, uint32(l1)+200))
	v239 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v238)+11)))
	switch v239 - int32(99) {
	case 0:
		v264 = int32(1)
		v265 = int32(8)
		goto L50
	case 1:
		goto L51
	default:
		goto L54
	case 6:
		goto L52
	case 16:
		goto L53
	}
L50:
	;
	if v97 <= int32(0) {
		goto L58
	} else {
		goto L59
	}
L51:
	;
	v264 = int32(8)
	v265 = int32(15)
	goto L50
L52:
	;
	v264 = int32(4)
	v265 = int32(11)
	goto L50
L53:
	;
	v264 = int32(2)
	v265 = int32(9)
	goto L50
L54:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v245 = m.ExcPending
	if v245 != 0 {
		goto L4
	} else {
		goto L55
	}
L55:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v17)+16)) = base.I32_extend8_s(v239)
	F_errmsg_internal(m, int32(_a_F_make_multirange_0), v17+int32(16))
	mBase = m.M
	v252 = m.ExcPending
	if v252 != 0 {
		goto L4
	} else {
		goto L56
	}
L56:
	;
	F_errfinish(m, int32(_a_F_make_multirange_1), int32(322), int32(_a_F_make_multirange_2))
	mBase = m.M
	v257 = m.ExcPending
	if v257 != 0 {
		goto L4
	} else {
		goto L57
	}
L57:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L58:
	;
	m.G0 = v17 + int32(32)
	return v230
L59:
	;
	v268 = int32(2)
	v272 = v230 + v97<<(uint(v268)%32) + int32(8)
	v273 = *(*int32)(unsafe.Add(mBase, uint32(l3)))
	v274 = *(*int32)(unsafe.Add(mBase, uint32(v273)))
	v280 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v273+int32(base.Ui32(v274)>>(uint(v268)%32))-int32(1)))))
	*(*uint8)(unsafe.Add(mBase, uint32(v272))) = uint8(v280)
	v283 = int32(0) - v264
	v288 = v230 + v283&(v97*int32(5)+v265)
	v289 = *(*int32)(unsafe.Add(mBase, uint32(l3)))
	v290 = *(*int32)(unsafe.Add(mBase, uint32(v289)))
	v294 = int32(base.Ui32(v290)>>(uint(v268)%32)) - int32(9)
	if v294 != 0 {
		goto L60
	} else {
		goto L61
	}
L60:
	;
	base.MemoryCopy(m, v288, v289+int32(8), v294)
	goto L62
L61:
	;
	goto L62
L62:
	;
	if v97 == int32(1) {
		goto L58
	} else {
		goto L63
	}
L63:
	;
	v300 = int32(1)
	v301 = v264 - v300
	v309 = int32(0)
	v311 = v288 + (v294+v301)&v283
	v312 = v300
	goto L64
L64:
	;
	v322 = v312 << (uint(int32(2)) % 32)
	v324 = v311 - v288
	if v312&int32(3) != 0 {
		goto L66
	} else {
		goto L67
	}
L65:
	;
	goto L58
L66:
	;
	v330 = v324 - v309
	goto L68
L67:
	;
	v330 = v324 | int32(-2147483648)
	goto L68
L68:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v230+v322)+8)) = v330
	v333 = l3 + v322
	v334 = *(*int32)(unsafe.Add(mBase, uint32(v333)))
	v335 = *(*int32)(unsafe.Add(mBase, uint32(v334)))
	v336 = int32(2)
	v341 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v334+int32(base.Ui32(v335)>>(uint(v336)%32))-int32(1)))))
	*(*uint8)(unsafe.Add(mBase, uint32(v312+v272))) = uint8(v341)
	v343 = *(*int32)(unsafe.Add(mBase, uint32(v333)))
	v344 = *(*int32)(unsafe.Add(mBase, uint32(v343)))
	v348 = int32(base.Ui32(v344)>>(uint(v336)%32)) - int32(9)
	if v348 != 0 {
		goto L69
	} else {
		goto L70
	}
L69:
	;
	base.MemoryCopy(m, v311, v343+int32(8), v348)
	goto L71
L70:
	;
	goto L71
L71:
	;
	v356 = v312 + int32(1)
	if v356 != v97 {
		v309 = v324
		v311 = v311 + (v348+v301)&v283
		v312 = v356
		goto L64
	} else {
		goto L72
	}
L72:
	;
	goto L65
}
func F_multirange_before_multirange(m *base.Module, l0 int32) int64 {
	mBase := m.M
	_ = mBase
	var v7 int64
	_ = v7
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
	var v64 int64
	_ = v64
	var v72 int32
	_ = v72
	var v76 int32
	_ = v76
	var v81 int32
	_ = v81
	v7 = int64(0)
	v8 = m.G0
	v10 = v8 - int32(80)
	m.G0 = v10
	v12 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
	v13 = F_pg_detoast_datum(m, v12)
	mBase = m.M
	v16 = m.ExcPending
	if v16 != 0 {
		return int64(0)
	} else {
		v17 = *(*int32)(unsafe.Add(mBase, uint32(l0)+40))
		v18 = F_pg_detoast_datum(m, v17)
		mBase = m.M
		v19 = m.ExcPending
		if v19 != 0 {
			return int64(0)
		} else {
			v20 = *(*int32)(unsafe.Add(mBase, uint32(v13)+4))
			v21 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
			v22 = *(*int32)(unsafe.Add(mBase, uint32(v21)+16))
			if v22 != 0 {
				v23 = *(*int32)(unsafe.Add(mBase, uint32(v22)))
				if v23 == v20 {
					v33 = v22
					v34 = *(*int32)(unsafe.Add(mBase, uint32(v13)+8))
					if v34 == int32(0) {
						v64 = v7
						m.G0 = v10 + int32(80)
						return v64
					} else {
						v37 = *(*int32)(unsafe.Add(mBase, uint32(v18)+8))
						if v37 == int32(0) {
							v64 = v7
							m.G0 = v10 + int32(80)
							return v64
						} else {
							v40 = *(*int32)(unsafe.Add(mBase, uint32(v33)+296))
							v46 = v10 + int32(48)
							F_multirange_get_bounds(m, v40, v13, v34-int32(1), v10-int32(-64), v46)
							mBase = m.M
							v48 = m.ExcPending
							if v48 != 0 {
								return int64(0)
							} else {
								v51 = v10 + int32(32)
								F_multirange_get_bounds(m, v40, v18, int32(0), v51, v10+int32(16))
								mBase = m.M
								v55 = m.ExcPending
								if v55 != 0 {
									return int64(0)
								} else {
									v56 = F_range_cmp_bounds(m, v40, v46, v51)
									mBase = m.M
									v57 = m.ExcPending
									if v57 != 0 {
										return int64(0)
									} else {
										v64 = base.I64_extend_i32_u(int32(base.Ui32(v56) >> (uint(int32(31)) % 32)))
										m.G0 = v10 + int32(80)
										return v64
									}
								}
							}
						}
					}
				} else {
					v26 = F_lookup_type_cache(m, v20, int32(_a_F_multirange_before_multirange_0))
					mBase = m.M
					v27 = m.ExcPending
					if v27 != 0 {
						return int64(0)
					} else {
						v28 = *(*int32)(unsafe.Add(mBase, uint32(v26)+296))
						if v28 == int32(0) {
							F_errstart_cold(m, int32(21), int32(0))
							mBase = m.M
							v72 = m.ExcPending
							if v72 != 0 {
								return int64(0)
							} else {
								*(*int32)(unsafe.Add(mBase, uint32(v10))) = v20
								F_errmsg_internal(m, int32(_a_F_multirange_before_multirange_1), v10)
								mBase = m.M
								v76 = m.ExcPending
								if v76 != 0 {
									return int64(0)
								} else {
									F_errfinish(m, int32(_a_F_multirange_before_multirange_2), int32(561), int32(_a_F_multirange_before_multirange_3))
									mBase = m.M
									v81 = m.ExcPending
									if v81 != 0 {
										return int64(0)
									} else {
										base.Wasm_trap_unreachable()
										for {
										}
									}
								}
							}
						} else {
							v31 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
							*(*int32)(unsafe.Add(mBase, uint32(v31)+16)) = v26
							v33 = v26
							v34 = *(*int32)(unsafe.Add(mBase, uint32(v13)+8))
							if v34 == int32(0) {
								v64 = v7
								m.G0 = v10 + int32(80)
								return v64
							} else {
								v37 = *(*int32)(unsafe.Add(mBase, uint32(v18)+8))
								if v37 == int32(0) {
									v64 = v7
									m.G0 = v10 + int32(80)
									return v64
								} else {
									v40 = *(*int32)(unsafe.Add(mBase, uint32(v33)+296))
									v46 = v10 + int32(48)
									F_multirange_get_bounds(m, v40, v13, v34-int32(1), v10-int32(-64), v46)
									mBase = m.M
									v48 = m.ExcPending
									if v48 != 0 {
										return int64(0)
									} else {
										v51 = v10 + int32(32)
										F_multirange_get_bounds(m, v40, v18, int32(0), v51, v10+int32(16))
										mBase = m.M
										v55 = m.ExcPending
										if v55 != 0 {
											return int64(0)
										} else {
											v56 = F_range_cmp_bounds(m, v40, v46, v51)
											mBase = m.M
											v57 = m.ExcPending
											if v57 != 0 {
												return int64(0)
											} else {
												v64 = base.I64_extend_i32_u(int32(base.Ui32(v56) >> (uint(int32(31)) % 32)))
												m.G0 = v10 + int32(80)
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
				v26 = F_lookup_type_cache(m, v20, int32(_a_F_multirange_before_multirange_0))
				mBase = m.M
				v27 = m.ExcPending
				if v27 != 0 {
					return int64(0)
				} else {
					v28 = *(*int32)(unsafe.Add(mBase, uint32(v26)+296))
					if v28 == int32(0) {
						F_errstart_cold(m, int32(21), int32(0))
						mBase = m.M
						v72 = m.ExcPending
						if v72 != 0 {
							return int64(0)
						} else {
							*(*int32)(unsafe.Add(mBase, uint32(v10))) = v20
							F_errmsg_internal(m, int32(_a_F_multirange_before_multirange_1), v10)
							mBase = m.M
							v76 = m.ExcPending
							if v76 != 0 {
								return int64(0)
							} else {
								F_errfinish(m, int32(_a_F_multirange_before_multirange_2), int32(561), int32(_a_F_multirange_before_multirange_3))
								mBase = m.M
								v81 = m.ExcPending
								if v81 != 0 {
									return int64(0)
								} else {
									base.Wasm_trap_unreachable()
									for {
									}
								}
							}
						}
					} else {
						v31 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
						*(*int32)(unsafe.Add(mBase, uint32(v31)+16)) = v26
						v33 = v26
						v34 = *(*int32)(unsafe.Add(mBase, uint32(v13)+8))
						if v34 == int32(0) {
							v64 = v7
							m.G0 = v10 + int32(80)
							return v64
						} else {
							v37 = *(*int32)(unsafe.Add(mBase, uint32(v18)+8))
							if v37 == int32(0) {
								v64 = v7
								m.G0 = v10 + int32(80)
								return v64
							} else {
								v40 = *(*int32)(unsafe.Add(mBase, uint32(v33)+296))
								v46 = v10 + int32(48)
								F_multirange_get_bounds(m, v40, v13, v34-int32(1), v10-int32(-64), v46)
								mBase = m.M
								v48 = m.ExcPending
								if v48 != 0 {
									return int64(0)
								} else {
									v51 = v10 + int32(32)
									F_multirange_get_bounds(m, v40, v18, int32(0), v51, v10+int32(16))
									mBase = m.M
									v55 = m.ExcPending
									if v55 != 0 {
										return int64(0)
									} else {
										v56 = F_range_cmp_bounds(m, v40, v46, v51)
										mBase = m.M
										v57 = m.ExcPending
										if v57 != 0 {
											return int64(0)
										} else {
											v64 = base.I64_extend_i32_u(int32(base.Ui32(v56) >> (uint(int32(31)) % 32)))
											m.G0 = v10 + int32(80)
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
func F_multirange_constructor0(m *base.Module, l0 int32) int64 {
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
	var v43 int32
	_ = v43
	var v47 int32
	_ = v47
	var v52 int32
	_ = v52
	var v56 int32
	_ = v56
	var v60 int32
	_ = v60
	var v65 int32
	_ = v65
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
			return int64(0)
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
						return int64(0)
					} else {
						m.G0 = v7 + int32(16)
						return base.I64_extend_i32_u(v33)
					}
				} else {
					v22 = F_lookup_type_cache(m, v13, int32(_a_F_multirange_constructor0_0))
					mBase = m.M
					v23 = m.ExcPending
					if v23 != 0 {
						return int64(0)
					} else {
						v24 = *(*int32)(unsafe.Add(mBase, uint32(v22)+296))
						if v24 == int32(0) {
							F_errstart_cold(m, int32(21), int32(0))
							mBase = m.M
							v56 = m.ExcPending
							if v56 != 0 {
								return int64(0)
							} else {
								*(*int32)(unsafe.Add(mBase, uint32(v7))) = v13
								F_errmsg_internal(m, int32(_a_F_multirange_constructor0_1), v7)
								mBase = m.M
								v60 = m.ExcPending
								if v60 != 0 {
									return int64(0)
								} else {
									F_errfinish(m, int32(_a_F_multirange_constructor0_2), int32(561), int32(_a_F_multirange_constructor0_3))
									mBase = m.M
									v65 = m.ExcPending
									if v65 != 0 {
										return int64(0)
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
								return int64(0)
							} else {
								m.G0 = v7 + int32(16)
								return base.I64_extend_i32_u(v33)
							}
						}
					}
				}
			} else {
				v22 = F_lookup_type_cache(m, v13, int32(_a_F_multirange_constructor0_0))
				mBase = m.M
				v23 = m.ExcPending
				if v23 != 0 {
					return int64(0)
				} else {
					v24 = *(*int32)(unsafe.Add(mBase, uint32(v22)+296))
					if v24 == int32(0) {
						F_errstart_cold(m, int32(21), int32(0))
						mBase = m.M
						v56 = m.ExcPending
						if v56 != 0 {
							return int64(0)
						} else {
							*(*int32)(unsafe.Add(mBase, uint32(v7))) = v13
							F_errmsg_internal(m, int32(_a_F_multirange_constructor0_1), v7)
							mBase = m.M
							v60 = m.ExcPending
							if v60 != 0 {
								return int64(0)
							} else {
								F_errfinish(m, int32(_a_F_multirange_constructor0_2), int32(561), int32(_a_F_multirange_constructor0_3))
								mBase = m.M
								v65 = m.ExcPending
								if v65 != 0 {
									return int64(0)
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
							return int64(0)
						} else {
							m.G0 = v7 + int32(16)
							return base.I64_extend_i32_u(v33)
						}
					}
				}
			}
		}
	} else {
		F_errstart_cold(m, int32(21), int32(0))
		mBase = m.M
		v43 = m.ExcPending
		if v43 != 0 {
			return int64(0)
		} else {
			F_errmsg_internal(m, int32(_a_F_multirange_constructor0_4), int32(0))
			mBase = m.M
			v47 = m.ExcPending
			if v47 != 0 {
				return int64(0)
			} else {
				F_errfinish(m, int32(_a_F_multirange_constructor0_2), int32(1074), int32(_a_F_multirange_constructor0_5))
				mBase = m.M
				v52 = m.ExcPending
				if v52 != 0 {
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
func F_multirange_contained_by_multirange(m *base.Module, l0 int32) int64 {
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
	var v44 int32
	_ = v44
	var v48 int32
	_ = v48
	var v53 int32
	_ = v53
	v7 = m.G0
	v9 = v7 - int32(16)
	m.G0 = v9
	v11 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
	v12 = F_pg_detoast_datum(m, v11)
	mBase = m.M
	v15 = m.ExcPending
	if v15 != 0 {
		return int64(0)
	} else {
		v16 = *(*int32)(unsafe.Add(mBase, uint32(l0)+40))
		v17 = F_pg_detoast_datum(m, v16)
		mBase = m.M
		v18 = m.ExcPending
		if v18 != 0 {
			return int64(0)
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
						return int64(0)
					} else {
						m.G0 = v9 + int32(16)
						return base.I64_extend_i32_u(v34)
					}
				} else {
					v25 = F_lookup_type_cache(m, v19, int32(_a_F_multirange_contained_by_multirange_0))
					mBase = m.M
					v26 = m.ExcPending
					if v26 != 0 {
						return int64(0)
					} else {
						v27 = *(*int32)(unsafe.Add(mBase, uint32(v25)+296))
						if v27 == int32(0) {
							F_errstart_cold(m, int32(21), int32(0))
							mBase = m.M
							v44 = m.ExcPending
							if v44 != 0 {
								return int64(0)
							} else {
								*(*int32)(unsafe.Add(mBase, uint32(v9))) = v19
								F_errmsg_internal(m, int32(_a_F_multirange_contained_by_multirange_1), v9)
								mBase = m.M
								v48 = m.ExcPending
								if v48 != 0 {
									return int64(0)
								} else {
									F_errfinish(m, int32(_a_F_multirange_contained_by_multirange_2), int32(561), int32(_a_F_multirange_contained_by_multirange_3))
									mBase = m.M
									v53 = m.ExcPending
									if v53 != 0 {
										return int64(0)
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
								return int64(0)
							} else {
								m.G0 = v9 + int32(16)
								return base.I64_extend_i32_u(v34)
							}
						}
					}
				}
			} else {
				v25 = F_lookup_type_cache(m, v19, int32(_a_F_multirange_contained_by_multirange_0))
				mBase = m.M
				v26 = m.ExcPending
				if v26 != 0 {
					return int64(0)
				} else {
					v27 = *(*int32)(unsafe.Add(mBase, uint32(v25)+296))
					if v27 == int32(0) {
						F_errstart_cold(m, int32(21), int32(0))
						mBase = m.M
						v44 = m.ExcPending
						if v44 != 0 {
							return int64(0)
						} else {
							*(*int32)(unsafe.Add(mBase, uint32(v9))) = v19
							F_errmsg_internal(m, int32(_a_F_multirange_contained_by_multirange_1), v9)
							mBase = m.M
							v48 = m.ExcPending
							if v48 != 0 {
								return int64(0)
							} else {
								F_errfinish(m, int32(_a_F_multirange_contained_by_multirange_2), int32(561), int32(_a_F_multirange_contained_by_multirange_3))
								mBase = m.M
								v53 = m.ExcPending
								if v53 != 0 {
									return int64(0)
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
							return int64(0)
						} else {
							m.G0 = v9 + int32(16)
							return base.I64_extend_i32_u(v34)
						}
					}
				}
			}
		}
	}
}
func F_multirange_ge(m *base.Module, l0 int32) int64 {
	var v2 int64
	_ = v2
	var v5 int32
	_ = v5
	v2 = F_multirange_cmp(m, l0)
	v5 = m.ExcPending
	if v5 != 0 {
		return int64(0)
	} else {
		return int64(base.Ui64(v2^int64(-1)) >> (uint(int64(63)) % 64))
	}
}
func F_multirange_gist_compress(m *base.Module, l0 int32) int64 {
	mBase := m.M
	_ = mBase
	var v7 int32
	_ = v7
	var v8 int32
	_ = v8
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
	var v39 int32
	_ = v39
	var v43 int32
	_ = v43
	var v45 int32
	_ = v45
	var v46 int32
	_ = v46
	var v48 int32
	_ = v48
	var v49 int32
	_ = v49
	var v52 int32
	_ = v52
	var v58 int32
	_ = v58
	var v60 int32
	_ = v60
	var v62 int32
	_ = v62
	var v63 int32
	_ = v63
	v7 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
	v8 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v7)+18)))
	if v8 != int32(1) {
		return base.I64_extend_i32_u(v7)
	} else {
		v13 = *(*int32)(unsafe.Add(mBase, uint32(v7)))
		v14 = F_pg_detoast_datum(m, v13)
		mBase = m.M
		v17 = m.ExcPending
		if v17 != 0 {
			return int64(0)
		} else {
			v19 = F_palloc(m, int32(24))
			mBase = m.M
			v20 = m.ExcPending
			if v20 != 0 {
				return int64(0)
			} else {
				v21 = *(*int32)(unsafe.Add(mBase, uint32(v14)+4))
				v22 = F_multirange_get_typcache(m, l0, v21)
				mBase = m.M
				v23 = m.ExcPending
				if v23 != 0 {
					return int64(0)
				} else {
					v24 = *(*int32)(unsafe.Add(mBase, uint32(v22)+296))
					v25 = m.G0
					v27 = v25 - int32(48)
					m.G0 = v27
					v29 = *(*int32)(unsafe.Add(mBase, uint32(v14)+8))
					if v29 == int32(0) {
						v32 = F_make_empty_range(m, v24)
						mBase = m.M
						v33 = m.ExcPending
						if v33 != 0 {
							return int64(0)
						} else {
							v52 = v32
							m.G0 = v27 + int32(48)
							*(*int64)(unsafe.Add(mBase, uint32(v19))) = base.I64_extend_i32_u(v52)
							v58 = *(*int32)(unsafe.Add(mBase, uint32(v7)+8))
							*(*int32)(unsafe.Add(mBase, uint32(v19)+8)) = v58
							v60 = *(*int32)(unsafe.Add(mBase, uint32(v7)+12))
							*(*int32)(unsafe.Add(mBase, uint32(v19)+12)) = v60
							v62 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v7)+16)))
							v63 = int32(0)
							*(*uint8)(unsafe.Add(mBase, uint32(v19)+18)) = uint8(v63)
							*(*uint16)(unsafe.Add(mBase, uint32(v19)+16)) = uint16(v62)
							return base.I64_extend_i32_u(v19)
						}
					} else {
						v36 = v27 + int32(32)
						F_multirange_get_bounds(m, v24, v14, int32(0), v36, v27)
						mBase = m.M
						v38 = m.ExcPending
						if v38 != 0 {
							return int64(0)
						} else {
							v39 = *(*int32)(unsafe.Add(mBase, uint32(v14)+8))
							v43 = v27 + int32(16)
							F_multirange_get_bounds(m, v24, v14, v39-int32(1), v27, v43)
							mBase = m.M
							v45 = m.ExcPending
							if v45 != 0 {
								return int64(0)
							} else {
								v46 = int32(0)
								v48 = F_make_range(m, v24, v36, v43, v46, v46)
								mBase = m.M
								v49 = m.ExcPending
								if v49 != 0 {
									return int64(0)
								} else {
									v52 = v48
									m.G0 = v27 + int32(48)
									*(*int64)(unsafe.Add(mBase, uint32(v19))) = base.I64_extend_i32_u(v52)
									v58 = *(*int32)(unsafe.Add(mBase, uint32(v7)+8))
									*(*int32)(unsafe.Add(mBase, uint32(v19)+8)) = v58
									v60 = *(*int32)(unsafe.Add(mBase, uint32(v7)+12))
									*(*int32)(unsafe.Add(mBase, uint32(v19)+12)) = v60
									v62 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v7)+16)))
									v63 = int32(0)
									*(*uint8)(unsafe.Add(mBase, uint32(v19)+18)) = uint8(v63)
									*(*uint16)(unsafe.Add(mBase, uint32(v19)+16)) = uint16(v62)
									return base.I64_extend_i32_u(v19)
								}
							}
						}
					}
				}
			}
		}
	}
}
func F_multirange_gist_consistent(m *base.Module, l0 int32) int64 {
	mBase := m.M
	_ = mBase
	var v9 int32
	_ = v9
	var v11 int32
	_ = v11
	var v13 int32
	_ = v13
	var v14 int64
	_ = v14
	var v15 int64
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
	var v30 int32
	_ = v30
	var v32 int32
	_ = v32
	var v42 int32
	_ = v42
	var v43 int32
	_ = v43
	var v44 int32
	_ = v44
	var v45 int32
	_ = v45
	var v49 int32
	_ = v49
	var v50 int32
	_ = v50
	var v51 int32
	_ = v51
	var v52 int32
	_ = v52
	var v54 int32
	_ = v54
	var v60 int32
	_ = v60
	var v66 int32
	_ = v66
	var v71 int32
	_ = v71
	var v79 int32
	_ = v79
	var v80 int32
	_ = v80
	var v81 int32
	_ = v81
	var v82 int32
	_ = v82
	var v86 int32
	_ = v86
	var v87 int32
	_ = v87
	var v88 int32
	_ = v88
	var v89 int32
	_ = v89
	var v91 int32
	_ = v91
	var v95 int32
	_ = v95
	var v96 int32
	_ = v96
	var v98 int32
	_ = v98
	var v107 int32
	_ = v107
	var v111 int32
	_ = v111
	var v116 int32
	_ = v116
	v9 = m.G0
	v11 = v9 - int32(32)
	m.G0 = v11
	v13 = *(*int32)(unsafe.Add(mBase, uint32(l0)+72))
	v14 = *(*int64)(unsafe.Add(mBase, uint32(l0)+40))
	v15 = *(*int64)(unsafe.Add(mBase, uint32(l0)+56))
	v16 = *(*int32)(unsafe.Add(mBase, uint32(l0)+88))
	v17 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
	v18 = *(*int32)(unsafe.Add(mBase, uint32(v17)))
	v19 = F_pg_detoast_datum(m, v18)
	mBase = m.M
	v22 = m.ExcPending
	if v22 != 0 {
		return int64(0)
	} else {
		v23 = int32(1)
		*(*uint8)(unsafe.Add(mBase, uint32(v16))) = uint8(v23)
		v25 = *(*int32)(unsafe.Add(mBase, uint32(v19)+4))
		v26 = F_range_get_typcache(m, l0, v25)
		mBase = m.M
		v27 = m.ExcPending
		if v27 != 0 {
			return int64(0)
		} else {
			v28 = base.I32_wrap_i64(v15)
			v29 = *(*int32)(unsafe.Add(mBase, uint32(v17)+12))
			v30 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v29)+16)))
			v32 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v29+v30)+12)))
			if v32&int32(1) != 0 {
				if v13 != int32(_a_F_multirange_gist_consistent_0) {
					if v13 == int32(3831) {
						v49 = F_pg_detoast_datum(m, base.I32_wrap_i64(v14))
						mBase = m.M
						v50 = m.ExcPending
						if v50 != 0 {
							return int64(0)
						} else {
							v51 = F_range_gist_consistent_leaf_range(m, v26, v28&int32(_a_F_multirange_gist_consistent_1), v19, v49)
							mBase = m.M
							v52 = m.ExcPending
							if v52 != 0 {
								return int64(0)
							} else {
								v98 = v51
								m.G0 = v11 + int32(32)
								return base.I64_extend_i32_u(v98)
							}
						}
					} else {
						if v13 != 0 {
							v54 = v28 & int32(_a_F_multirange_gist_consistent_1)
							if v54 == int32(16) {
								v95 = F_range_contains_elem_internal(m, v26, v19, v14)
								mBase = m.M
								v96 = m.ExcPending
								if v96 != 0 {
									return int64(0)
								} else {
									v98 = v95
									m.G0 = v11 + int32(32)
									return base.I64_extend_i32_u(v98)
								}
							} else {
								F_errstart_cold(m, int32(21), int32(0))
								mBase = m.M
								v60 = m.ExcPending
								if v60 != 0 {
									return int64(0)
								} else {
									*(*int32)(unsafe.Add(mBase, uint32(v11)+16)) = v54
									F_errmsg_internal(m, int32(_a_F_multirange_gist_consistent_2), v11+int32(16))
									mBase = m.M
									v66 = m.ExcPending
									if v66 != 0 {
										return int64(0)
									} else {
										F_errfinish(m, int32(_a_F_multirange_gist_consistent_3), int32(1138), int32(_a_F_multirange_gist_consistent_4))
										mBase = m.M
										v71 = m.ExcPending
										if v71 != 0 {
											return int64(0)
										} else {
											base.Wasm_trap_unreachable()
											for {
											}
										}
									}
								}
							}
						} else {
							v42 = F_pg_detoast_datum(m, base.I32_wrap_i64(v14))
							mBase = m.M
							v43 = m.ExcPending
							if v43 != 0 {
								return int64(0)
							} else {
								v44 = F_range_gist_consistent_leaf_multirange(m, v26, v28&int32(_a_F_multirange_gist_consistent_1), v19, v42)
								mBase = m.M
								v45 = m.ExcPending
								if v45 != 0 {
									return int64(0)
								} else {
									v98 = v44
									m.G0 = v11 + int32(32)
									return base.I64_extend_i32_u(v98)
								}
							}
						}
					}
				} else {
					v42 = F_pg_detoast_datum(m, base.I32_wrap_i64(v14))
					mBase = m.M
					v43 = m.ExcPending
					if v43 != 0 {
						return int64(0)
					} else {
						v44 = F_range_gist_consistent_leaf_multirange(m, v26, v28&int32(_a_F_multirange_gist_consistent_1), v19, v42)
						mBase = m.M
						v45 = m.ExcPending
						if v45 != 0 {
							return int64(0)
						} else {
							v98 = v44
							m.G0 = v11 + int32(32)
							return base.I64_extend_i32_u(v98)
						}
					}
				}
			} else {
				if v13 != int32(_a_F_multirange_gist_consistent_0) {
					if v13 == int32(3831) {
						v86 = F_pg_detoast_datum(m, base.I32_wrap_i64(v14))
						mBase = m.M
						v87 = m.ExcPending
						if v87 != 0 {
							return int64(0)
						} else {
							v88 = F_range_gist_consistent_int_range(m, v26, v28&int32(_a_F_multirange_gist_consistent_1), v19, v86)
							mBase = m.M
							v89 = m.ExcPending
							if v89 != 0 {
								return int64(0)
							} else {
								v98 = v88
								m.G0 = v11 + int32(32)
								return base.I64_extend_i32_u(v98)
							}
						}
					} else {
						if v13 != 0 {
							v91 = v28 & int32(_a_F_multirange_gist_consistent_1)
							if v91 != int32(16) {
								F_errstart_cold(m, int32(21), int32(0))
								mBase = m.M
								v107 = m.ExcPending
								if v107 != 0 {
									return int64(0)
								} else {
									*(*int32)(unsafe.Add(mBase, uint32(v11))) = v91
									F_errmsg_internal(m, int32(_a_F_multirange_gist_consistent_2), v11)
									mBase = m.M
									v111 = m.ExcPending
									if v111 != 0 {
										return int64(0)
									} else {
										F_errfinish(m, int32(_a_F_multirange_gist_consistent_3), int32(1049), int32(_a_F_multirange_gist_consistent_5))
										mBase = m.M
										v116 = m.ExcPending
										if v116 != 0 {
											return int64(0)
										} else {
											base.Wasm_trap_unreachable()
											for {
											}
										}
									}
								}
							} else {
								v95 = F_range_contains_elem_internal(m, v26, v19, v14)
								mBase = m.M
								v96 = m.ExcPending
								if v96 != 0 {
									return int64(0)
								} else {
									v98 = v95
									m.G0 = v11 + int32(32)
									return base.I64_extend_i32_u(v98)
								}
							}
						} else {
							v79 = F_pg_detoast_datum(m, base.I32_wrap_i64(v14))
							mBase = m.M
							v80 = m.ExcPending
							if v80 != 0 {
								return int64(0)
							} else {
								v81 = F_range_gist_consistent_int_multirange(m, v26, v28&int32(_a_F_multirange_gist_consistent_1), v19, v79)
								mBase = m.M
								v82 = m.ExcPending
								if v82 != 0 {
									return int64(0)
								} else {
									v98 = v81
									m.G0 = v11 + int32(32)
									return base.I64_extend_i32_u(v98)
								}
							}
						}
					}
				} else {
					v79 = F_pg_detoast_datum(m, base.I32_wrap_i64(v14))
					mBase = m.M
					v80 = m.ExcPending
					if v80 != 0 {
						return int64(0)
					} else {
						v81 = F_range_gist_consistent_int_multirange(m, v26, v28&int32(_a_F_multirange_gist_consistent_1), v19, v79)
						mBase = m.M
						v82 = m.ExcPending
						if v82 != 0 {
							return int64(0)
						} else {
							v98 = v81
							m.G0 = v11 + int32(32)
							return base.I64_extend_i32_u(v98)
						}
					}
				}
			}
		}
	}
}
func F_multirange_intersect_agg_transfn(m *base.Module, l0 int32) int64 {
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
	var v82 int32
	_ = v82
	var v83 int32
	_ = v83
	var v84 int32
	_ = v84
	var v97 int32
	_ = v97
	var v98 int32
	_ = v98
	var v101 int32
	_ = v101
	var v103 int32
	_ = v103
	var v107 int32
	_ = v107
	var v112 int32
	_ = v112
	var v114 int32
	_ = v114
	var v119 int32
	_ = v119
	var v120 int32
	_ = v120
	var v121 int32
	_ = v121
	var v134 int32
	_ = v134
	var v135 int32
	_ = v135
	var v138 int32
	_ = v138
	var v140 int32
	_ = v140
	var v150 int32
	_ = v150
	var v151 int32
	_ = v151
	var v152 int32
	_ = v152
	var v153 int32
	_ = v153
	var v162 int32
	_ = v162
	var v166 int32
	_ = v166
	var v171 int32
	_ = v171
	var v175 int32
	_ = v175
	var v179 int32
	_ = v179
	var v184 int32
	_ = v184
	var v188 int32
	_ = v188
	var v192 int32
	_ = v192
	var v197 int32
	_ = v197
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
	v188 = m.ExcPending
	if v188 != 0 {
		goto L20
	} else {
		goto L56
	}
L2:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v175 = m.ExcPending
	if v175 != 0 {
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
	switch v21 - int32(435) {
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
	v33 = *(*int32)(unsafe.Add(mBase, uint32(v18)+376))
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
	v162 = m.ExcPending
	if v162 != 0 {
		goto L20
	} else {
		goto L50
	}
L20:
	;
	return int64(0)
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
	v70 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
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
	v73 = *(*int32)(unsafe.Add(mBase, uint32(l0)+40))
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
	v82 = F_palloc_mul(m, int32(4), v77)
	mBase = m.M
	v83 = m.ExcPending
	if v83 != 0 {
		goto L20
	} else {
		goto L36
	}
L34:
	;
	v107 = v76
	v112 = v2
	goto L35
L35:
	;
	v114 = *(*int32)(unsafe.Add(mBase, uint32(v74)+8))
	if int32(0) < v114 {
		goto L41
	} else {
		goto L42
	}
L36:
	;
	v84 = int32(0)
	goto L37
L37:
	;
	v97 = F_multirange_get_range(m, v76, v71, v84)
	mBase = m.M
	v98 = m.ExcPending
	if v98 != 0 {
		goto L20
	} else {
		goto L39
	}
L38:
	;
	v103 = *(*int32)(unsafe.Add(mBase, uint32(v69)+296))
	v107 = v103
	v112 = v82
	goto L35
L39:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v82+v84<<(uint(int32(2))%32)))) = v97
	v101 = v84 + int32(1)
	if v101 != v77 {
		v84 = v101
		goto L37
	} else {
		goto L40
	}
L40:
	;
	goto L38
L41:
	;
	v119 = F_palloc_mul(m, int32(4), v114)
	mBase = m.M
	v120 = m.ExcPending
	if v120 != 0 {
		goto L20
	} else {
		goto L44
	}
L42:
	;
	v150 = v2
	v151 = v107
	goto L43
L43:
	;
	v152 = F_multirange_intersect_internal(m, v49, v151, v77, v112, v114, v150)
	mBase = m.M
	v153 = m.ExcPending
	if v153 != 0 {
		goto L20
	} else {
		goto L49
	}
L44:
	;
	v121 = int32(0)
	goto L45
L45:
	;
	v134 = F_multirange_get_range(m, v107, v74, v121)
	mBase = m.M
	v135 = m.ExcPending
	if v135 != 0 {
		goto L20
	} else {
		goto L47
	}
L46:
	;
	v140 = *(*int32)(unsafe.Add(mBase, uint32(v69)+296))
	v150 = v119
	v151 = v140
	goto L43
L47:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v119+v121<<(uint(int32(2))%32)))) = v134
	v138 = v121 + int32(1)
	if v138 != v114 {
		v121 = v138
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
	return base.I64_extend_i32_u(v152)
L50:
	;
	F_errmsg_internal(m, int32(_a_F_multirange_intersect_agg_transfn_1), int32(0))
	mBase = m.M
	v166 = m.ExcPending
	if v166 != 0 {
		goto L20
	} else {
		goto L51
	}
L51:
	;
	F_errfinish(m, int32(_a_F_multirange_intersect_agg_transfn_2), int32(1555), int32(_a_F_multirange_intersect_agg_transfn_3))
	mBase = m.M
	v171 = m.ExcPending
	if v171 != 0 {
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
	v179 = m.ExcPending
	if v179 != 0 {
		goto L20
	} else {
		goto L54
	}
L54:
	;
	F_errfinish(m, int32(_a_F_multirange_intersect_agg_transfn_2), int32(1559), int32(_a_F_multirange_intersect_agg_transfn_3))
	mBase = m.M
	v184 = m.ExcPending
	if v184 != 0 {
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
	v192 = m.ExcPending
	if v192 != 0 {
		goto L20
	} else {
		goto L57
	}
L57:
	;
	F_errfinish(m, int32(_a_F_multirange_intersect_agg_transfn_2), int32(561), int32(_a_F_multirange_intersect_agg_transfn_6))
	mBase = m.M
	v197 = m.ExcPending
	if v197 != 0 {
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
func F_multirange_upper(m *base.Module, l0 int32) int64 {
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
	var v44 int32
	_ = v44
	var v45 int32
	_ = v45
	var v48 int64
	_ = v48
	var v52 int32
	_ = v52
	var v58 int64
	_ = v58
	var v66 int32
	_ = v66
	var v70 int32
	_ = v70
	var v75 int32
	_ = v75
	v8 = m.G0
	v10 = v8 - int32(48)
	m.G0 = v10
	v12 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
	v13 = F_pg_detoast_datum(m, v12)
	mBase = m.M
	v16 = m.ExcPending
	if v16 != 0 {
		return int64(0)
	} else {
		v17 = *(*int32)(unsafe.Add(mBase, uint32(v13)+8))
		if v17 == int32(0) {
			v52 = int32(1)
			*(*uint8)(unsafe.Add(mBase, uint32(l0)+16)) = uint8(v52)
			v58 = int64(0)
			m.G0 = v10 + int32(48)
			return v58
		} else {
			v20 = *(*int32)(unsafe.Add(mBase, uint32(v13)+4))
			v21 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
			v22 = *(*int32)(unsafe.Add(mBase, uint32(v21)+16))
			if v22 != 0 {
				v23 = *(*int32)(unsafe.Add(mBase, uint32(v22)))
				if v23 == v20 {
					v34 = v22
					v35 = v17
					v36 = *(*int32)(unsafe.Add(mBase, uint32(v34)+296))
					F_multirange_get_bounds(m, v36, v13, v35-int32(1), v10+int32(32), v10+int32(16))
					mBase = m.M
					v44 = m.ExcPending
					if v44 != 0 {
						return int64(0)
					} else {
						v45 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v10)+24)))
						if v45 == int32(0) {
							v48 = *(*int64)(unsafe.Add(mBase, uint32(v10)+16))
							v58 = v48
						} else {
							v52 = int32(1)
							*(*uint8)(unsafe.Add(mBase, uint32(l0)+16)) = uint8(v52)
							v58 = int64(0)
						}
						m.G0 = v10 + int32(48)
						return v58
					}
				} else {
					v26 = F_lookup_type_cache(m, v20, int32(_a_F_multirange_upper_0))
					mBase = m.M
					v27 = m.ExcPending
					if v27 != 0 {
						return int64(0)
					} else {
						v28 = *(*int32)(unsafe.Add(mBase, uint32(v26)+296))
						if v28 == int32(0) {
							F_errstart_cold(m, int32(21), int32(0))
							mBase = m.M
							v66 = m.ExcPending
							if v66 != 0 {
								return int64(0)
							} else {
								*(*int32)(unsafe.Add(mBase, uint32(v10))) = v20
								F_errmsg_internal(m, int32(_a_F_multirange_upper_1), v10)
								mBase = m.M
								v70 = m.ExcPending
								if v70 != 0 {
									return int64(0)
								} else {
									F_errfinish(m, int32(_a_F_multirange_upper_2), int32(561), int32(_a_F_multirange_upper_3))
									mBase = m.M
									v75 = m.ExcPending
									if v75 != 0 {
										return int64(0)
									} else {
										base.Wasm_trap_unreachable()
										for {
										}
									}
								}
							}
						} else {
							v31 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
							*(*int32)(unsafe.Add(mBase, uint32(v31)+16)) = v26
							v33 = *(*int32)(unsafe.Add(mBase, uint32(v13)+8))
							v34 = v26
							v35 = v33
							v36 = *(*int32)(unsafe.Add(mBase, uint32(v34)+296))
							F_multirange_get_bounds(m, v36, v13, v35-int32(1), v10+int32(32), v10+int32(16))
							mBase = m.M
							v44 = m.ExcPending
							if v44 != 0 {
								return int64(0)
							} else {
								v45 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v10)+24)))
								if v45 == int32(0) {
									v48 = *(*int64)(unsafe.Add(mBase, uint32(v10)+16))
									v58 = v48
								} else {
									v52 = int32(1)
									*(*uint8)(unsafe.Add(mBase, uint32(l0)+16)) = uint8(v52)
									v58 = int64(0)
								}
								m.G0 = v10 + int32(48)
								return v58
							}
						}
					}
				}
			} else {
				v26 = F_lookup_type_cache(m, v20, int32(_a_F_multirange_upper_0))
				mBase = m.M
				v27 = m.ExcPending
				if v27 != 0 {
					return int64(0)
				} else {
					v28 = *(*int32)(unsafe.Add(mBase, uint32(v26)+296))
					if v28 == int32(0) {
						F_errstart_cold(m, int32(21), int32(0))
						mBase = m.M
						v66 = m.ExcPending
						if v66 != 0 {
							return int64(0)
						} else {
							*(*int32)(unsafe.Add(mBase, uint32(v10))) = v20
							F_errmsg_internal(m, int32(_a_F_multirange_upper_1), v10)
							mBase = m.M
							v70 = m.ExcPending
							if v70 != 0 {
								return int64(0)
							} else {
								F_errfinish(m, int32(_a_F_multirange_upper_2), int32(561), int32(_a_F_multirange_upper_3))
								mBase = m.M
								v75 = m.ExcPending
								if v75 != 0 {
									return int64(0)
								} else {
									base.Wasm_trap_unreachable()
									for {
									}
								}
							}
						}
					} else {
						v31 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
						*(*int32)(unsafe.Add(mBase, uint32(v31)+16)) = v26
						v33 = *(*int32)(unsafe.Add(mBase, uint32(v13)+8))
						v34 = v26
						v35 = v33
						v36 = *(*int32)(unsafe.Add(mBase, uint32(v34)+296))
						F_multirange_get_bounds(m, v36, v13, v35-int32(1), v10+int32(32), v10+int32(16))
						mBase = m.M
						v44 = m.ExcPending
						if v44 != 0 {
							return int64(0)
						} else {
							v45 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v10)+24)))
							if v45 == int32(0) {
								v48 = *(*int64)(unsafe.Add(mBase, uint32(v10)+16))
								v58 = v48
							} else {
								v52 = int32(1)
								*(*uint8)(unsafe.Add(mBase, uint32(l0)+16)) = uint8(v52)
								v58 = int64(0)
							}
							m.G0 = v10 + int32(48)
							return v58
						}
					}
				}
			}
		}
	}
}
