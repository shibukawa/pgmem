package p1

import (
	base "github.com/shibukawa/pgmem/internal/aot/pgaot/base"
	"math"
	"unsafe"
)

func F_extract_interval(m *base.Module, l0 int32) int32 {
	var v3 int32
	_ = v3
	var v6 int32
	_ = v6
	v3 = F_interval_part_common(m, l0, int32(1))
	v6 = m.ExcPending
	if v6 != 0 {
		return int32(0)
	} else {
		return v3
	}
}
func F_interval_avg_serialize(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v6 int32
	_ = v6
	var v8 int32
	_ = v8
	var v12 int32
	_ = v12
	var v15 int32
	_ = v15
	var v40 int32
	_ = v40
	var v48 int32
	_ = v48
	var v52 int32
	_ = v52
	var v57 int32
	_ = v57
	var v58 int32
	_ = v58
	var v60 int32
	_ = v60
	var v61 int64
	_ = v61
	var v64 int32
	_ = v64
	var v65 int32
	_ = v65
	var v66 int32
	_ = v66
	var v68 int64
	_ = v68
	var v70 int64
	_ = v70
	var v72 int64
	_ = v72
	var v75 int64
	_ = v75
	var v77 int64
	_ = v77
	var v79 int64
	_ = v79
	var v81 int64
	_ = v81
	var v104 int32
	_ = v104
	var v107 int64
	_ = v107
	var v110 int32
	_ = v110
	var v111 int32
	_ = v111
	var v112 int32
	_ = v112
	var v114 int64
	_ = v114
	var v116 int64
	_ = v116
	var v118 int64
	_ = v118
	var v121 int64
	_ = v121
	var v123 int64
	_ = v123
	var v125 int64
	_ = v125
	var v127 int64
	_ = v127
	var v153 int32
	_ = v153
	var v156 int32
	_ = v156
	var v157 int32
	_ = v157
	var v158 int32
	_ = v158
	var v160 int32
	_ = v160
	var v162 int32
	_ = v162
	var v164 int32
	_ = v164
	var v176 int32
	_ = v176
	var v179 int32
	_ = v179
	var v182 int32
	_ = v182
	var v183 int32
	_ = v183
	var v184 int32
	_ = v184
	var v186 int32
	_ = v186
	var v188 int32
	_ = v188
	var v190 int32
	_ = v190
	var v205 int64
	_ = v205
	var v208 int32
	_ = v208
	var v209 int32
	_ = v209
	var v210 int32
	_ = v210
	var v212 int64
	_ = v212
	var v214 int64
	_ = v214
	var v216 int64
	_ = v216
	var v219 int64
	_ = v219
	var v221 int64
	_ = v221
	var v223 int64
	_ = v223
	var v225 int64
	_ = v225
	var v248 int32
	_ = v248
	var v251 int64
	_ = v251
	var v254 int32
	_ = v254
	var v255 int32
	_ = v255
	var v256 int32
	_ = v256
	var v258 int64
	_ = v258
	var v260 int64
	_ = v260
	var v262 int64
	_ = v262
	var v265 int64
	_ = v265
	var v267 int64
	_ = v267
	var v269 int64
	_ = v269
	var v271 int64
	_ = v271
	var v295 int32
	_ = v295
	var v298 int32
	_ = v298
	v6 = m.G0
	v8 = v6 - int32(16)
	m.G0 = v8
	v12 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	if v12 == int32(0) {
		v40 = int32(0)
	} else {
		v15 = *(*int32)(unsafe.Add(mBase, uint32(v12)))
		switch v15 - int32(429) {
		case 0:
			v40 = int32(1)
		case 1:
			v40 = int32(2)
		default:
			v40 = int32(0)
		}
	}
	if v40 == int32(0) {
		F_errstart_cold(m, int32(21), int32(0))
		mBase = m.M
		v48 = m.ExcPending
		if v48 != 0 {
			return int32(0)
		} else {
			F_errmsg_internal(m, int32(61060), int32(0))
			mBase = m.M
			v52 = m.ExcPending
			if v52 != 0 {
				return int32(0)
			} else {
				F_errfinish(m, int32(495596), int32(4135), int32(341915))
				mBase = m.M
				v57 = m.ExcPending
				if v57 != 0 {
					return int32(0)
				} else {
					base.Wasm_trap_unreachable()
					for {
					}
				}
			}
		}
	} else {
		v58 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
		F_pq_begintypsend(m, v8)
		mBase = m.M
		v60 = m.ExcPending
		if v60 != 0 {
			return int32(0)
		} else {
			v61 = *(*int64)(unsafe.Add(mBase, uint32(v58)))
			F_enlargeStringInfo(m, v8, int32(8))
			mBase = m.M
			v64 = m.ExcPending
			if v64 != 0 {
				return int32(0)
			} else {
				v65 = *(*int32)(unsafe.Add(mBase, uint32(v8)+4))
				v66 = *(*int32)(unsafe.Add(mBase, uint32(v8)))
				v68 = int64(56)
				v70 = int64(65280)
				v72 = int64(40)
				v75 = int64(16711680)
				v77 = int64(24)
				v79 = int64(4278190080)
				v81 = int64(8)
				*(*int64)(unsafe.Add(mBase, uint32(v65+v66))) = v61<<(uint(v68)%64) | v61&v70<<(uint(v72)%64) | (v61&v75<<(uint(v77)%64) | v61&v79<<(uint(v81)%64)) | (int64(base.Ui64(v61)>>(uint(v81)%64))&v79 | int64(base.Ui64(v61)>>(uint(v77)%64))&v75 | (int64(base.Ui64(v61)>>(uint(v72)%64))&v70 | int64(base.Ui64(v61)>>(uint(v68)%64))))
				v104 = int32(8)
				*(*int32)(unsafe.Add(mBase, uint32(v8)+4)) = v65 + v104
				v107 = *(*int64)(unsafe.Add(mBase, uint32(v58)+8))
				F_enlargeStringInfo(m, v8, v104)
				mBase = m.M
				v110 = m.ExcPending
				if v110 != 0 {
					return int32(0)
				} else {
					v111 = *(*int32)(unsafe.Add(mBase, uint32(v8)+4))
					v112 = *(*int32)(unsafe.Add(mBase, uint32(v8)))
					v114 = int64(56)
					v116 = int64(65280)
					v118 = int64(40)
					v121 = int64(16711680)
					v123 = int64(24)
					v125 = int64(4278190080)
					v127 = int64(8)
					*(*int64)(unsafe.Add(mBase, uint32(v111+v112))) = v107<<(uint(v114)%64) | v107&v116<<(uint(v118)%64) | (v107&v121<<(uint(v123)%64) | v107&v125<<(uint(v127)%64)) | (int64(base.Ui64(v107)>>(uint(v127)%64))&v125 | int64(base.Ui64(v107)>>(uint(v123)%64))&v121 | (int64(base.Ui64(v107)>>(uint(v118)%64))&v116 | int64(base.Ui64(v107)>>(uint(v114)%64))))
					*(*int32)(unsafe.Add(mBase, uint32(v8)+4)) = v111 + int32(8)
					v153 = *(*int32)(unsafe.Add(mBase, uint32(v58)+16))
					F_enlargeStringInfo(m, v8, int32(4))
					mBase = m.M
					v156 = m.ExcPending
					if v156 != 0 {
						return int32(0)
					} else {
						v157 = *(*int32)(unsafe.Add(mBase, uint32(v8)+4))
						v158 = *(*int32)(unsafe.Add(mBase, uint32(v8)))
						v160 = int32(24)
						v162 = int32(65280)
						v164 = int32(8)
						*(*int32)(unsafe.Add(mBase, uint32(v157+v158))) = v153<<(uint(v160)%32) | v153&v162<<(uint(v164)%32) | (int32(base.Ui32(v153)>>(uint(v164)%32))&v162 | int32(base.Ui32(v153)>>(uint(v160)%32)))
						v176 = int32(4)
						*(*int32)(unsafe.Add(mBase, uint32(v8)+4)) = v157 + v176
						v179 = *(*int32)(unsafe.Add(mBase, uint32(v58)+20))
						F_enlargeStringInfo(m, v8, v176)
						mBase = m.M
						v182 = m.ExcPending
						if v182 != 0 {
							return int32(0)
						} else {
							v183 = *(*int32)(unsafe.Add(mBase, uint32(v8)+4))
							v184 = *(*int32)(unsafe.Add(mBase, uint32(v8)))
							v186 = int32(24)
							v188 = int32(65280)
							v190 = int32(8)
							*(*int32)(unsafe.Add(mBase, uint32(v183+v184))) = v179<<(uint(v186)%32) | v179&v188<<(uint(v190)%32) | (int32(base.Ui32(v179)>>(uint(v190)%32))&v188 | int32(base.Ui32(v179)>>(uint(v186)%32)))
							*(*int32)(unsafe.Add(mBase, uint32(v8)+4)) = v183 + int32(4)
							v205 = *(*int64)(unsafe.Add(mBase, uint32(v58)+24))
							F_enlargeStringInfo(m, v8, v190)
							mBase = m.M
							v208 = m.ExcPending
							if v208 != 0 {
								return int32(0)
							} else {
								v209 = *(*int32)(unsafe.Add(mBase, uint32(v8)+4))
								v210 = *(*int32)(unsafe.Add(mBase, uint32(v8)))
								v212 = int64(56)
								v214 = int64(65280)
								v216 = int64(40)
								v219 = int64(16711680)
								v221 = int64(24)
								v223 = int64(4278190080)
								v225 = int64(8)
								*(*int64)(unsafe.Add(mBase, uint32(v209+v210))) = v205<<(uint(v212)%64) | v205&v214<<(uint(v216)%64) | (v205&v219<<(uint(v221)%64) | v205&v223<<(uint(v225)%64)) | (int64(base.Ui64(v205)>>(uint(v225)%64))&v223 | int64(base.Ui64(v205)>>(uint(v221)%64))&v219 | (int64(base.Ui64(v205)>>(uint(v216)%64))&v214 | int64(base.Ui64(v205)>>(uint(v212)%64))))
								v248 = int32(8)
								*(*int32)(unsafe.Add(mBase, uint32(v8)+4)) = v209 + v248
								v251 = *(*int64)(unsafe.Add(mBase, uint32(v58)+32))
								F_enlargeStringInfo(m, v8, v248)
								mBase = m.M
								v254 = m.ExcPending
								if v254 != 0 {
									return int32(0)
								} else {
									v255 = *(*int32)(unsafe.Add(mBase, uint32(v8)+4))
									v256 = *(*int32)(unsafe.Add(mBase, uint32(v8)))
									v258 = int64(56)
									v260 = int64(65280)
									v262 = int64(40)
									v265 = int64(16711680)
									v267 = int64(24)
									v269 = int64(4278190080)
									v271 = int64(8)
									*(*int64)(unsafe.Add(mBase, uint32(v255+v256))) = v251<<(uint(v258)%64) | v251&v260<<(uint(v262)%64) | (v251&v265<<(uint(v267)%64) | v251&v269<<(uint(v271)%64)) | (int64(base.Ui64(v251)>>(uint(v271)%64))&v269 | int64(base.Ui64(v251)>>(uint(v267)%64))&v265 | (int64(base.Ui64(v251)>>(uint(v262)%64))&v260 | int64(base.Ui64(v251)>>(uint(v258)%64))))
									v295 = v255 + int32(8)
									*(*int32)(unsafe.Add(mBase, uint32(v8)+4)) = v295
									v298 = *(*int32)(unsafe.Add(mBase, uint32(v8)))
									*(*int32)(unsafe.Add(mBase, uint32(v298))) = v295 << (uint(int32(2)) % 32)
									m.G0 = v8 + int32(16)
									return v298
								}
							}
						}
					}
				}
			}
		}
	}
}
func F_interval_div(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v12 int32
	_ = v12
	var v13 float64
	_ = v13
	var v14 int32
	_ = v14
	var v15 int32
	_ = v15
	var v16 int32
	_ = v16
	var v18 int32
	_ = v18
	var v21 int32
	_ = v21
	var v29 int32
	_ = v29
	var v34 int32
	_ = v34
	var v37 int64
	_ = v37
	var v40 int32
	_ = v40
	var v43 int64
	_ = v43
	var v52 int32
	_ = v52
	var v54 int64
	_ = v54
	var v56 int64
	_ = v56
	var v61 int32
	_ = v61
	var v64 int32
	_ = v64
	var v68 int32
	_ = v68
	var v73 int32
	_ = v73
	var v75 float64
	_ = v75
	var v92 int32
	_ = v92
	var v94 int32
	_ = v94
	var v96 int32
	_ = v96
	var v98 float64
	_ = v98
	var v117 int32
	_ = v117
	var v119 int32
	_ = v119
	var v126 float64
	_ = v126
	var v130 float64
	_ = v130
	var v137 int32
	_ = v137
	var v139 int32
	_ = v139
	var v142 float64
	_ = v142
	var v144 float64
	_ = v144
	var v148 float64
	_ = v148
	var v155 float64
	_ = v155
	var v159 int32
	_ = v159
	var v161 int32
	_ = v161
	var v162 int32
	_ = v162
	var v172 float64
	_ = v172
	var v174 int32
	_ = v174
	var v176 int32
	_ = v176
	var v184 int64
	_ = v184
	var v188 float64
	_ = v188
	var v205 int64
	_ = v205
	var v207 int64
	_ = v207
	var v233 int32
	_ = v233
	var v236 int32
	_ = v236
	var v240 int32
	_ = v240
	var v245 int32
	_ = v245
	v12 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
	v13 = *(*float64)(unsafe.Add(mBase, uint32(v12)))
	v14 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v15 = *(*int32)(unsafe.Add(mBase, uint32(v14)+8))
	v16 = *(*int32)(unsafe.Add(mBase, uint32(v14)+12))
	v18 = F_palloc(m, int32(16))
	mBase = m.M
	v21 = m.ExcPending
	if v21 != 0 {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	return int32(0)
L2:
	;
	if base.F64_ne(v13, float64(0)) != 0 {
		goto L6
	} else {
		goto L7
	}
L3:
	;
	return v18
L4:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v233 = m.ExcPending
	if v233 != 0 {
		goto L1
	} else {
		goto L71
	}
L5:
	;
	v75 = base.F64_div(base.F64_convert_i32_s(v29), v13)
	if base.F64_lt(v75, float64(2.147483648e+09)) == int32(0) {
		goto L4
	} else {
		goto L28
	}
L6:
	;
	if base.Ui64(int64(9218868437227405312)) < base.Ui64(base.I64_reinterpret_f64(v13)&int64(9223372036854775807)) {
		goto L4
	} else {
		goto L9
	}
L7:
	;
	goto L8
L8:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v61 = m.ExcPending
	if v61 != 0 {
		goto L1
	} else {
		goto L24
	}
L9:
	;
	v29 = *(*int32)(unsafe.Add(mBase, uint32(v14)+12))
	if v29 != int32(2147483647) {
		goto L11
	} else {
		goto L12
	}
L10:
	;
	if base.F64_eq(base.F64_abs(v13), math.Float64frombits(uint64(0x7ff0000000000000))) != 0 {
		goto L4
	} else {
		goto L19
	}
L11:
	;
	if v29 != int32(-2147483648) {
		goto L5
	} else {
		goto L14
	}
L12:
	;
	goto L13
L13:
	;
	v40 = *(*int32)(unsafe.Add(mBase, uint32(v14)+8))
	if v40 != int32(2147483647) {
		goto L5
	} else {
		goto L17
	}
L14:
	;
	v34 = *(*int32)(unsafe.Add(mBase, uint32(v14)+8))
	if v34 != int32(-2147483648) {
		goto L5
	} else {
		goto L15
	}
L15:
	;
	v37 = *(*int64)(unsafe.Add(mBase, uint32(v14)))
	if v37 == int64(-9223372036854775807-1) {
		goto L10
	} else {
		goto L16
	}
L16:
	;
	goto L5
L17:
	;
	v43 = *(*int64)(unsafe.Add(mBase, uint32(v14)))
	if v43 != int64(9223372036854775807) {
		goto L5
	} else {
		goto L18
	}
L18:
	;
	goto L10
L19:
	;
	if base.F64_lt(v13, float64(0)) != 0 {
		goto L20
	} else {
		goto L21
	}
L20:
	;
	F_interval_um_internal(m, v14, v18)
	mBase = m.M
	v52 = m.ExcPending
	if v52 != 0 {
		goto L1
	} else {
		goto L23
	}
L21:
	;
	goto L22
L22:
	;
	v54 = *(*int64)(unsafe.Add(mBase, uint32(v14)))
	*(*int64)(unsafe.Add(mBase, uint32(v18))) = v54
	v56 = *(*int64)(unsafe.Add(mBase, uint32(v14)+8))
	*(*int64)(unsafe.Add(mBase, uint32(v18)+8)) = v56
	goto L3
L23:
	;
	return v18
L24:
	;
	F_errcode(m, int32(33816706))
	mBase = m.M
	v64 = m.ExcPending
	if v64 != 0 {
		goto L1
	} else {
		goto L25
	}
L25:
	;
	F_errmsg(m, int32(239436), int32(0))
	mBase = m.M
	v68 = m.ExcPending
	if v68 != 0 {
		goto L1
	} else {
		goto L26
	}
L26:
	;
	F_errfinish(m, int32(495596), int32(3772), int32(35442))
	mBase = m.M
	v73 = m.ExcPending
	if v73 != 0 {
		goto L1
	} else {
		goto L27
	}
L27:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L28:
	;
	if base.F64_ge(v75, float64(-2.147483648e+09)) == int32(0) {
		goto L4
	} else {
		goto L29
	}
L29:
	;
	if base.Ui64(int64(9218868437227405312)) < base.Ui64(base.I64_reinterpret_f64(v75)&int64(9223372036854775807)) {
		goto L4
	} else {
		goto L30
	}
L30:
	;
	if base.F64_lt(base.F64_abs(v75), float64(2.147483648e+09)) != 0 {
		goto L32
	} else {
		goto L33
	}
L31:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v18)+12)) = v94
	v96 = *(*int32)(unsafe.Add(mBase, uint32(v14)+8))
	v98 = base.F64_div(base.F64_convert_i32_s(v96), v13)
	if base.F64_lt(v98, float64(2.147483648e+09)) == int32(0) {
		goto L4
	} else {
		goto L35
	}
L32:
	;
	v92 = base.I32_trunc_f64_s(v75)
	v94 = v92
	goto L31
L33:
	;
	goto L34
L34:
	;
	v94 = int32(-2147483648)
	goto L31
L35:
	;
	if base.F64_ge(v98, float64(-2.147483648e+09)) == int32(0) {
		goto L4
	} else {
		goto L36
	}
L36:
	;
	if base.Ui64(int64(9218868437227405312)) < base.Ui64(base.I64_reinterpret_f64(v98)&int64(9223372036854775807)) {
		goto L4
	} else {
		goto L37
	}
L37:
	;
	if base.F64_lt(base.F64_abs(v98), float64(2.147483648e+09)) != 0 {
		goto L39
	} else {
		goto L40
	}
L38:
	;
	v126 = float64(1e+06)
	v130 = base.F64_div(base.F64_nearest(base.F64_mul(base.F64_mul(base.F64_sub(base.F64_div(base.F64_convert_i32_s(v16), v13), base.F64_convert_i32_s(v94)), float64(30)), v126)), v126)
	if base.F64_lt(base.F64_abs(v130), float64(2.147483648e+09)) != 0 {
		goto L44
	} else {
		goto L45
	}
L39:
	;
	v117 = base.I32_trunc_f64_s(v98)
	v119 = v117
	goto L38
L40:
	;
	goto L41
L41:
	;
	v119 = int32(-2147483648)
	goto L38
L42:
	;
	v176 = v174 + v139
	*(*int32)(unsafe.Add(mBase, uint32(v18)+8)) = v176
	if base.B2i32(v139 < int32(0))^base.B2i32(v176 < v174) != 0 {
		goto L4
	} else {
		goto L55
	}
L43:
	;
	v142 = float64(86400)
	v144 = float64(1e+06)
	v148 = base.F64_div(base.F64_nearest(base.F64_mul(base.F64_mul(base.F64_sub(base.F64_add(v130, base.F64_sub(base.F64_div(base.F64_convert_i32_s(v15), v13), base.F64_convert_i32_s(v119))), base.F64_convert_i32_s(v139)), v142), v144)), v144)
	if base.F64_ge(base.F64_abs(v148), v142) == int32(0) {
		goto L47
	} else {
		goto L48
	}
L44:
	;
	v137 = base.I32_trunc_f64_s(v130)
	v139 = v137
	goto L43
L45:
	;
	goto L46
L46:
	;
	v139 = int32(-2147483648)
	goto L43
L47:
	;
	v172 = v148
	v174 = v119
	goto L42
L48:
	;
	goto L49
L49:
	;
	v155 = base.F64_div(v148, float64(86400))
	if base.F64_lt(base.F64_abs(v155), float64(2.147483648e+09)) != 0 {
		goto L51
	} else {
		goto L52
	}
L50:
	;
	v162 = v161 + v119
	*(*int32)(unsafe.Add(mBase, uint32(v18)+8)) = v162
	if base.B2i32(v161 < int32(0))^base.B2i32(v162 < v119) != 0 {
		goto L4
	} else {
		goto L54
	}
L51:
	;
	v159 = base.I32_trunc_f64_s(v155)
	v161 = v159
	goto L50
L52:
	;
	goto L53
L53:
	;
	v161 = int32(-2147483648)
	goto L50
L54:
	;
	v172 = base.F64_sub(v148, base.F64_convert_i32_s(v161*int32(86400)))
	v174 = v162
	goto L42
L55:
	;
	v184 = *(*int64)(unsafe.Add(mBase, uint32(v14)))
	v188 = base.F64_nearest(base.F64_add(base.F64_mul(v172, float64(1e+06)), base.F64_div(base.F64_convert_i64_s(v184), v13)))
	if base.F64_lt(v188, float64(9.223372036854776e+18)) == int32(0) {
		goto L4
	} else {
		goto L56
	}
L56:
	;
	if base.F64_ge(v188, float64(-9.223372036854776e+18)) == int32(0) {
		goto L4
	} else {
		goto L57
	}
L57:
	;
	if base.Ui64(int64(9218868437227405312)) < base.Ui64(base.I64_reinterpret_f64(v188)&int64(9223372036854775807)) {
		goto L4
	} else {
		goto L58
	}
L58:
	;
	if base.F64_lt(base.F64_abs(v188), float64(9.223372036854776e+18)) != 0 {
		goto L60
	} else {
		goto L61
	}
L59:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v18))) = v207
	if v94 != int32(2147483647) {
		goto L63
	} else {
		goto L64
	}
L60:
	;
	v205 = base.I64_trunc_f64_s(v188)
	v207 = v205
	goto L59
L61:
	;
	goto L62
L62:
	;
	v207 = int64(-9223372036854775807 - 1)
	goto L59
L63:
	;
	if v94 != int32(-2147483648) {
		goto L3
	} else {
		goto L66
	}
L64:
	;
	goto L65
L65:
	;
	if v176 != int32(2147483647) {
		goto L3
	} else {
		goto L69
	}
L66:
	;
	if v176 != int32(-2147483648) {
		goto L3
	} else {
		goto L67
	}
L67:
	;
	if v207 == int64(-9223372036854775807-1) {
		goto L4
	} else {
		goto L68
	}
L68:
	;
	goto L3
L69:
	;
	if v207 != int64(9223372036854775807) {
		goto L3
	} else {
		goto L70
	}
L70:
	;
	goto L4
L71:
	;
	F_errcode(m, int32(134217858))
	mBase = m.M
	v236 = m.ExcPending
	if v236 != 0 {
		goto L1
	} else {
		goto L72
	}
L72:
	;
	F_errmsg(m, int32(402232), int32(0))
	mBase = m.M
	v240 = m.ExcPending
	if v240 != 0 {
		goto L1
	} else {
		goto L73
	}
L73:
	;
	F_errfinish(m, int32(495596), int32(3841), int32(35442))
	mBase = m.M
	v245 = m.ExcPending
	if v245 != 0 {
		goto L1
	} else {
		goto L74
	}
L74:
	;
	base.Wasm_trap_unreachable()
	for {
	}
}
func F_interval_gt(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v12 int32
	_ = v12
	var v14 int32
	_ = v14
	var v17 int32
	_ = v17
	var v18 int32
	_ = v18
	var v19 int64
	_ = v19
	var v22 int64
	_ = v22
	var v23 int64
	_ = v23
	var v32 int64
	_ = v32
	var v33 int64
	_ = v33
	var v35 int64
	_ = v35
	var v38 int64
	_ = v38
	var v39 int64
	_ = v39
	var v41 int64
	_ = v41
	var v42 int64
	_ = v42
	var v46 int64
	_ = v46
	var v53 int64
	_ = v53
	var v64 int32
	_ = v64
	var v65 int64
	_ = v65
	var v68 int64
	_ = v68
	var v69 int64
	_ = v69
	var v78 int64
	_ = v78
	var v79 int64
	_ = v79
	var v81 int64
	_ = v81
	var v84 int64
	_ = v84
	var v85 int64
	_ = v85
	var v87 int64
	_ = v87
	var v88 int64
	_ = v88
	var v92 int64
	_ = v92
	var v99 int64
	_ = v99
	var v110 int64
	_ = v110
	var v111 int64
	_ = v111
	var v112 int64
	_ = v112
	var v113 int64
	_ = v113
	var v114 int64
	_ = v114
	var v115 int64
	_ = v115
	var v119 int64
	_ = v119
	var v120 int64
	_ = v120
	var v124 int64
	_ = v124
	var v127 int64
	_ = v127
	var v133 int64
	_ = v133
	var v136 int32
	_ = v136
	v12 = m.G0
	v14 = v12 - int32(32)
	m.G0 = v14
	v17 = v14 + int32(16)
	v18 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v19 = int64(*(*int32)(unsafe.Add(mBase, uint32(v18)+12)))
	v22 = int64(*(*int32)(unsafe.Add(mBase, uint32(v18)+8)))
	v23 = v19*int64(30) + v22
	v32 = int64(32)
	v33 = int64(20)
	v35 = int64(base.Ui64(v23) >> (uint(v32) % 64))
	v38 = int64(4294967295)
	v39 = int64(500654080)
	v41 = v23 & v38
	v42 = v39 * v41
	v46 = int64(base.Ui64(v42)>>(uint(v32)%64)) + v39*v35
	v53 = v41*v33 + v46&v38
	*(*int64)(unsafe.Add(mBase, uint32(v17)+8)) = v23*int64(0) + v23>>(uint(int64(63))%64)*int64(86400000000) + v33*v35 + int64(base.Ui64(v46)>>(uint(v32)%64)) + int64(base.Ui64(v53)>>(uint(v32)%64))
	*(*int64)(unsafe.Add(mBase, uint32(v17))) = v42&v38 | v53<<(uint(v32)%64)
	v64 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
	v65 = int64(*(*int32)(unsafe.Add(mBase, uint32(v64)+12)))
	v68 = int64(*(*int32)(unsafe.Add(mBase, uint32(v64)+8)))
	v69 = v65*int64(30) + v68
	v78 = int64(32)
	v79 = int64(20)
	v81 = int64(base.Ui64(v69) >> (uint(v78) % 64))
	v84 = int64(4294967295)
	v85 = int64(500654080)
	v87 = v69 & v84
	v88 = v85 * v87
	v92 = int64(base.Ui64(v88)>>(uint(v78)%64)) + v85*v81
	v99 = v87*v79 + v92&v84
	*(*int64)(unsafe.Add(mBase, uint32(v14)+8)) = v69*int64(0) + v69>>(uint(int64(63))%64)*int64(86400000000) + v79*v81 + int64(base.Ui64(v92)>>(uint(v78)%64)) + int64(base.Ui64(v99)>>(uint(v78)%64))
	*(*int64)(unsafe.Add(mBase, uint32(v14))) = v88&v84 | v99<<(uint(v78)%64)
	v110 = *(*int64)(unsafe.Add(mBase, uint32(v18)))
	v111 = *(*int64)(unsafe.Add(mBase, uint32(v14)+24))
	v112 = *(*int64)(unsafe.Add(mBase, uint32(v14)+16))
	v113 = *(*int64)(unsafe.Add(mBase, uint32(v64)))
	v114 = *(*int64)(unsafe.Add(mBase, uint32(v14)+8))
	v115 = *(*int64)(unsafe.Add(mBase, uint32(v14)))
	m.G0 = v14 + int32(32)
	v119 = v110 + v112
	v120 = v113 + v115
	v124 = int64(63)
	v127 = base.I64_extend_i32_u(base.B2i32(base.Ui64(v119) < base.Ui64(v112))) + (v111 + v110>>(uint(v124)%64))
	v133 = base.I64_extend_i32_u(base.B2i32(base.Ui64(v120) < base.Ui64(v115))) + (v114 + v113>>(uint(v124)%64))
	if v127 == v133 {
		v136 = base.B2i32(base.Ui64(v120) < base.Ui64(v119))
	} else {
		v136 = base.B2i32(v133 < v127)
	}
	return v136
}
func F_interval_le(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v12 int32
	_ = v12
	var v14 int32
	_ = v14
	var v17 int32
	_ = v17
	var v18 int32
	_ = v18
	var v19 int64
	_ = v19
	var v22 int64
	_ = v22
	var v23 int64
	_ = v23
	var v32 int64
	_ = v32
	var v33 int64
	_ = v33
	var v35 int64
	_ = v35
	var v38 int64
	_ = v38
	var v39 int64
	_ = v39
	var v41 int64
	_ = v41
	var v42 int64
	_ = v42
	var v46 int64
	_ = v46
	var v53 int64
	_ = v53
	var v64 int32
	_ = v64
	var v65 int64
	_ = v65
	var v68 int64
	_ = v68
	var v69 int64
	_ = v69
	var v78 int64
	_ = v78
	var v79 int64
	_ = v79
	var v81 int64
	_ = v81
	var v84 int64
	_ = v84
	var v85 int64
	_ = v85
	var v87 int64
	_ = v87
	var v88 int64
	_ = v88
	var v92 int64
	_ = v92
	var v99 int64
	_ = v99
	var v110 int64
	_ = v110
	var v111 int64
	_ = v111
	var v112 int64
	_ = v112
	var v113 int64
	_ = v113
	var v114 int64
	_ = v114
	var v115 int64
	_ = v115
	var v119 int64
	_ = v119
	var v120 int64
	_ = v120
	var v124 int64
	_ = v124
	var v127 int64
	_ = v127
	var v133 int64
	_ = v133
	var v136 int32
	_ = v136
	v12 = m.G0
	v14 = v12 - int32(32)
	m.G0 = v14
	v17 = v14 + int32(16)
	v18 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v19 = int64(*(*int32)(unsafe.Add(mBase, uint32(v18)+12)))
	v22 = int64(*(*int32)(unsafe.Add(mBase, uint32(v18)+8)))
	v23 = v19*int64(30) + v22
	v32 = int64(32)
	v33 = int64(20)
	v35 = int64(base.Ui64(v23) >> (uint(v32) % 64))
	v38 = int64(4294967295)
	v39 = int64(500654080)
	v41 = v23 & v38
	v42 = v39 * v41
	v46 = int64(base.Ui64(v42)>>(uint(v32)%64)) + v39*v35
	v53 = v41*v33 + v46&v38
	*(*int64)(unsafe.Add(mBase, uint32(v17)+8)) = v23*int64(0) + v23>>(uint(int64(63))%64)*int64(86400000000) + v33*v35 + int64(base.Ui64(v46)>>(uint(v32)%64)) + int64(base.Ui64(v53)>>(uint(v32)%64))
	*(*int64)(unsafe.Add(mBase, uint32(v17))) = v42&v38 | v53<<(uint(v32)%64)
	v64 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
	v65 = int64(*(*int32)(unsafe.Add(mBase, uint32(v64)+12)))
	v68 = int64(*(*int32)(unsafe.Add(mBase, uint32(v64)+8)))
	v69 = v65*int64(30) + v68
	v78 = int64(32)
	v79 = int64(20)
	v81 = int64(base.Ui64(v69) >> (uint(v78) % 64))
	v84 = int64(4294967295)
	v85 = int64(500654080)
	v87 = v69 & v84
	v88 = v85 * v87
	v92 = int64(base.Ui64(v88)>>(uint(v78)%64)) + v85*v81
	v99 = v87*v79 + v92&v84
	*(*int64)(unsafe.Add(mBase, uint32(v14)+8)) = v69*int64(0) + v69>>(uint(int64(63))%64)*int64(86400000000) + v79*v81 + int64(base.Ui64(v92)>>(uint(v78)%64)) + int64(base.Ui64(v99)>>(uint(v78)%64))
	*(*int64)(unsafe.Add(mBase, uint32(v14))) = v88&v84 | v99<<(uint(v78)%64)
	v110 = *(*int64)(unsafe.Add(mBase, uint32(v18)))
	v111 = *(*int64)(unsafe.Add(mBase, uint32(v14)+24))
	v112 = *(*int64)(unsafe.Add(mBase, uint32(v14)+16))
	v113 = *(*int64)(unsafe.Add(mBase, uint32(v64)))
	v114 = *(*int64)(unsafe.Add(mBase, uint32(v14)+8))
	v115 = *(*int64)(unsafe.Add(mBase, uint32(v14)))
	m.G0 = v14 + int32(32)
	v119 = v110 + v112
	v120 = v113 + v115
	v124 = int64(63)
	v127 = base.I64_extend_i32_u(base.B2i32(base.Ui64(v119) < base.Ui64(v112))) + (v111 + v110>>(uint(v124)%64))
	v133 = base.I64_extend_i32_u(base.B2i32(base.Ui64(v120) < base.Ui64(v115))) + (v114 + v113>>(uint(v124)%64))
	if v127 == v133 {
		v136 = base.B2i32(base.Ui64(v119) <= base.Ui64(v120))
	} else {
		v136 = base.B2i32(v127 <= v133)
	}
	return v136
}
func F_interval_part(m *base.Module, l0 int32) int32 {
	var v3 int32
	_ = v3
	var v6 int32
	_ = v6
	v3 = F_interval_part_common(m, l0, int32(0))
	v6 = m.ExcPending
	if v6 != 0 {
		return int32(0)
	} else {
		return v3
	}
}
func F_interval_time(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v4 int32
	_ = v4
	var v5 int32
	_ = v5
	var v10 int64
	_ = v10
	var v11 int64
	_ = v11
	var v12 int32
	_ = v12
	var v17 int64
	_ = v17
	var v18 int32
	_ = v18
	var v29 int32
	_ = v29
	var v32 int32
	_ = v32
	var v36 int32
	_ = v36
	var v41 int32
	_ = v41
	var v42 int64
	_ = v42
	var v43 int64
	_ = v43
	var v44 int64
	_ = v44
	var v49 int64
	_ = v49
	var v50 int32
	_ = v50
	var v51 int32
	_ = v51
	v4 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v5 = *(*int32)(unsafe.Add(mBase, uint32(v4)+12))
	if v5 != int32(-2147483648) {
		if v5 == int32(2147483647) {
			v17 = *(*int64)(unsafe.Add(mBase, uint32(v4)))
			v18 = *(*int32)(unsafe.Add(mBase, uint32(v4)+8))
			if v18 != int32(2147483647) {
				v42 = v17
				v43 = int64(86400000000)
				v44 = base.I64_rem_s(v42, v43)
				if v44 < int64(0) {
					v49 = v44 + v43
				} else {
					v49 = v44
				}
				v50 = F_Int64GetDatum(m, v49)
				mBase = m.M
				v51 = m.ExcPending
				if v51 != 0 {
					return int32(0)
				} else {
					return v50
				}
			} else {
				if v17 != int64(9223372036854775807) {
					v42 = v17
					v43 = int64(86400000000)
					v44 = base.I64_rem_s(v42, v43)
					if v44 < int64(0) {
						v49 = v44 + v43
					} else {
						v49 = v44
					}
					v50 = F_Int64GetDatum(m, v49)
					mBase = m.M
					v51 = m.ExcPending
					if v51 != 0 {
						return int32(0)
					} else {
						return v50
					}
				} else {
					F_errstart_cold(m, int32(21), int32(0))
					mBase = m.M
					v29 = m.ExcPending
					if v29 != 0 {
						return int32(0)
					} else {
						F_errcode(m, int32(134217858))
						mBase = m.M
						v32 = m.ExcPending
						if v32 != 0 {
							return int32(0)
						} else {
							F_errmsg(m, int32(375451), int32(0))
							mBase = m.M
							v36 = m.ExcPending
							if v36 != 0 {
								return int32(0)
							} else {
								F_errfinish(m, int32(498380), int32(2085), int32(374847))
								mBase = m.M
								v41 = m.ExcPending
								if v41 != 0 {
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
		} else {
			v10 = *(*int64)(unsafe.Add(mBase, uint32(v4)))
			v42 = v10
			v43 = int64(86400000000)
			v44 = base.I64_rem_s(v42, v43)
			if v44 < int64(0) {
				v49 = v44 + v43
			} else {
				v49 = v44
			}
			v50 = F_Int64GetDatum(m, v49)
			mBase = m.M
			v51 = m.ExcPending
			if v51 != 0 {
				return int32(0)
			} else {
				return v50
			}
		}
	} else {
		v11 = *(*int64)(unsafe.Add(mBase, uint32(v4)))
		v12 = *(*int32)(unsafe.Add(mBase, uint32(v4)+8))
		if v12 != int32(-2147483648) {
			v42 = v11
			v43 = int64(86400000000)
			v44 = base.I64_rem_s(v42, v43)
			if v44 < int64(0) {
				v49 = v44 + v43
			} else {
				v49 = v44
			}
			v50 = F_Int64GetDatum(m, v49)
			mBase = m.M
			v51 = m.ExcPending
			if v51 != 0 {
				return int32(0)
			} else {
				return v50
			}
		} else {
			if v11 == int64(-9223372036854775807-1) {
				F_errstart_cold(m, int32(21), int32(0))
				mBase = m.M
				v29 = m.ExcPending
				if v29 != 0 {
					return int32(0)
				} else {
					F_errcode(m, int32(134217858))
					mBase = m.M
					v32 = m.ExcPending
					if v32 != 0 {
						return int32(0)
					} else {
						F_errmsg(m, int32(375451), int32(0))
						mBase = m.M
						v36 = m.ExcPending
						if v36 != 0 {
							return int32(0)
						} else {
							F_errfinish(m, int32(498380), int32(2085), int32(374847))
							mBase = m.M
							v41 = m.ExcPending
							if v41 != 0 {
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
				v42 = v11
				v43 = int64(86400000000)
				v44 = base.I64_rem_s(v42, v43)
				if v44 < int64(0) {
					v49 = v44 + v43
				} else {
					v49 = v44
				}
				v50 = F_Int64GetDatum(m, v49)
				mBase = m.M
				v51 = m.ExcPending
				if v51 != 0 {
					return int32(0)
				} else {
					return v50
				}
			}
		}
	}
}
