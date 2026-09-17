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
	var v162 int32
	_ = v162
	var v170 int32
	_ = v170
	var v173 int32
	_ = v173
	var v176 int32
	_ = v176
	var v177 int32
	_ = v177
	var v178 int32
	_ = v178
	var v182 int32
	_ = v182
	var v186 int32
	_ = v186
	var v193 int64
	_ = v193
	var v196 int32
	_ = v196
	var v197 int32
	_ = v197
	var v198 int32
	_ = v198
	var v200 int64
	_ = v200
	var v202 int64
	_ = v202
	var v204 int64
	_ = v204
	var v207 int64
	_ = v207
	var v209 int64
	_ = v209
	var v211 int64
	_ = v211
	var v213 int64
	_ = v213
	var v236 int32
	_ = v236
	var v239 int64
	_ = v239
	var v242 int32
	_ = v242
	var v243 int32
	_ = v243
	var v244 int32
	_ = v244
	var v246 int64
	_ = v246
	var v248 int64
	_ = v248
	var v250 int64
	_ = v250
	var v253 int64
	_ = v253
	var v255 int64
	_ = v255
	var v257 int64
	_ = v257
	var v259 int64
	_ = v259
	var v283 int32
	_ = v283
	var v286 int32
	_ = v286
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
			F_errmsg_internal(m, int32(_a_F_interval_avg_serialize_0), int32(0))
			mBase = m.M
			v52 = m.ExcPending
			if v52 != 0 {
				return int32(0)
			} else {
				F_errfinish(m, int32(_a_F_interval_avg_serialize_1), int32(_a_F_interval_avg_serialize_2), int32(_a_F_interval_avg_serialize_3))
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
						v162 = int32(16711935)
						*(*int32)(unsafe.Add(mBase, uint32(v157+v158))) = base.I32_rotr(v153, int32(24))&v162 | base.I32_rotr(v153&v162, int32(8))
						v170 = int32(4)
						*(*int32)(unsafe.Add(mBase, uint32(v8)+4)) = v157 + v170
						v173 = *(*int32)(unsafe.Add(mBase, uint32(v58)+20))
						F_enlargeStringInfo(m, v8, v170)
						mBase = m.M
						v176 = m.ExcPending
						if v176 != 0 {
							return int32(0)
						} else {
							v177 = *(*int32)(unsafe.Add(mBase, uint32(v8)+4))
							v178 = *(*int32)(unsafe.Add(mBase, uint32(v8)))
							v182 = int32(16711935)
							v186 = int32(8)
							*(*int32)(unsafe.Add(mBase, uint32(v177+v178))) = base.I32_rotr(v173, int32(24))&v182 | base.I32_rotr(v173&v182, v186)
							*(*int32)(unsafe.Add(mBase, uint32(v8)+4)) = v177 + int32(4)
							v193 = *(*int64)(unsafe.Add(mBase, uint32(v58)+24))
							F_enlargeStringInfo(m, v8, v186)
							mBase = m.M
							v196 = m.ExcPending
							if v196 != 0 {
								return int32(0)
							} else {
								v197 = *(*int32)(unsafe.Add(mBase, uint32(v8)+4))
								v198 = *(*int32)(unsafe.Add(mBase, uint32(v8)))
								v200 = int64(56)
								v202 = int64(65280)
								v204 = int64(40)
								v207 = int64(16711680)
								v209 = int64(24)
								v211 = int64(4278190080)
								v213 = int64(8)
								*(*int64)(unsafe.Add(mBase, uint32(v197+v198))) = v193<<(uint(v200)%64) | v193&v202<<(uint(v204)%64) | (v193&v207<<(uint(v209)%64) | v193&v211<<(uint(v213)%64)) | (int64(base.Ui64(v193)>>(uint(v213)%64))&v211 | int64(base.Ui64(v193)>>(uint(v209)%64))&v207 | (int64(base.Ui64(v193)>>(uint(v204)%64))&v202 | int64(base.Ui64(v193)>>(uint(v200)%64))))
								v236 = int32(8)
								*(*int32)(unsafe.Add(mBase, uint32(v8)+4)) = v197 + v236
								v239 = *(*int64)(unsafe.Add(mBase, uint32(v58)+32))
								F_enlargeStringInfo(m, v8, v236)
								mBase = m.M
								v242 = m.ExcPending
								if v242 != 0 {
									return int32(0)
								} else {
									v243 = *(*int32)(unsafe.Add(mBase, uint32(v8)+4))
									v244 = *(*int32)(unsafe.Add(mBase, uint32(v8)))
									v246 = int64(56)
									v248 = int64(65280)
									v250 = int64(40)
									v253 = int64(16711680)
									v255 = int64(24)
									v257 = int64(4278190080)
									v259 = int64(8)
									*(*int64)(unsafe.Add(mBase, uint32(v243+v244))) = v239<<(uint(v246)%64) | v239&v248<<(uint(v250)%64) | (v239&v253<<(uint(v255)%64) | v239&v257<<(uint(v259)%64)) | (int64(base.Ui64(v239)>>(uint(v259)%64))&v257 | int64(base.Ui64(v239)>>(uint(v255)%64))&v253 | (int64(base.Ui64(v239)>>(uint(v250)%64))&v248 | int64(base.Ui64(v239)>>(uint(v246)%64))))
									v283 = v243 + int32(8)
									*(*int32)(unsafe.Add(mBase, uint32(v8)+4)) = v283
									v286 = *(*int32)(unsafe.Add(mBase, uint32(v8)))
									*(*int32)(unsafe.Add(mBase, uint32(v286))) = v283 << (uint(int32(2)) % 32)
									m.G0 = v8 + int32(16)
									return v286
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
	var v78 int32
	_ = v78
	var v91 int32
	_ = v91
	var v93 int32
	_ = v93
	var v95 float64
	_ = v95
	var v98 int32
	_ = v98
	var v117 float64
	_ = v117
	var v121 float64
	_ = v121
	var v124 int32
	_ = v124
	var v128 int32
	_ = v128
	var v131 float64
	_ = v131
	var v137 float64
	_ = v137
	var v145 int32
	_ = v145
	var v146 int32
	_ = v146
	var v156 int32
	_ = v156
	var v157 float64
	_ = v157
	var v159 int32
	_ = v159
	var v167 int64
	_ = v167
	var v171 float64
	_ = v171
	var v174 int32
	_ = v174
	var v187 int64
	_ = v187
	var v191 int32
	_ = v191
	var v215 int32
	_ = v215
	var v218 int32
	_ = v218
	var v222 int32
	_ = v222
	var v227 int32
	_ = v227
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
	v215 = m.ExcPending
	if v215 != 0 {
		goto L1
	} else {
		goto L43
	}
L5:
	;
	v75 = base.F64_div(base.F64_convert_i32_s(v29), v13)
	v78 = int32(0)
	if base.B2i32(base.F64_lt(v75, float64(2.147483648e+09)) == v78)|base.B2i32(base.F64_ge(v75, float64(-2.147483648e+09)) == v78)|base.B2i32(base.Ui64(int64(9218868437227405312)) < base.Ui64(base.I64_reinterpret_f64(v75)&int64(9223372036854775807))) != 0 {
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
	v54 = *(*int64)(unsafe.Add(mBase, uint32(v14)+8))
	*(*int64)(unsafe.Add(mBase, uint32(v18)+8)) = v54
	v56 = *(*int64)(unsafe.Add(mBase, uint32(v14)))
	*(*int64)(unsafe.Add(mBase, uint32(v18))) = v56
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
	F_errmsg(m, int32(_a_F_interval_div_4), int32(0))
	mBase = m.M
	v68 = m.ExcPending
	if v68 != 0 {
		goto L1
	} else {
		goto L26
	}
L26:
	;
	F_errfinish(m, int32(_a_F_interval_div_1), int32(3772), int32(_a_F_interval_div_2))
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
	v91 = base.I32_trunc_sat_f64_s(v75)
	*(*int32)(unsafe.Add(mBase, uint32(v18)+12)) = v91
	v93 = *(*int32)(unsafe.Add(mBase, uint32(v14)+8))
	v95 = base.F64_div(base.F64_convert_i32_s(v93), v13)
	v98 = int32(0)
	if base.B2i32(base.F64_lt(v95, float64(2.147483648e+09)) == v98)|base.B2i32(base.F64_ge(v95, float64(-2.147483648e+09)) == v98)|base.B2i32(base.Ui64(int64(9218868437227405312)) < base.Ui64(base.I64_reinterpret_f64(v95)&int64(9223372036854775807))) != 0 {
		goto L4
	} else {
		goto L29
	}
L29:
	;
	v117 = float64(1e+06)
	v121 = base.F64_div(base.F64_nearest(base.F64_mul(base.F64_mul(base.F64_sub(base.F64_div(base.F64_convert_i32_s(v16), v13), base.F64_convert_i32_s(v91)), float64(30)), v117)), v117)
	v124 = base.I32_trunc_sat_f64_s(v95)
	v128 = base.I32_trunc_sat_f64_s(v121)
	v131 = float64(86400)
	v137 = base.F64_div(base.F64_nearest(base.F64_mul(base.F64_mul(base.F64_sub(base.F64_add(v121, base.F64_sub(base.F64_div(base.F64_convert_i32_s(v15), v13), base.F64_convert_i32_s(v124))), base.F64_convert_i32_s(v128)), v131), v117)), v117)
	if base.F64_ge(base.F64_abs(v137), v131) == int32(0) {
		goto L31
	} else {
		goto L32
	}
L30:
	;
	v159 = v156 + v128
	*(*int32)(unsafe.Add(mBase, uint32(v18)+8)) = v159
	if base.B2i32(v128 < int32(0))^base.B2i32(v159 < v156) != 0 {
		goto L4
	} else {
		goto L35
	}
L31:
	;
	v156 = v124
	v157 = v137
	goto L30
L32:
	;
	goto L33
L33:
	;
	v145 = base.I32_trunc_sat_f64_s(base.F64_div(v137, float64(86400)))
	v146 = v124 + v145
	*(*int32)(unsafe.Add(mBase, uint32(v18)+8)) = v146
	if base.B2i32(v145 < int32(0))^base.B2i32(v146 < v124) != 0 {
		goto L4
	} else {
		goto L34
	}
L34:
	;
	v156 = v146
	v157 = base.F64_sub(v137, base.F64_convert_i32_s(v145*int32(_a_F_interval_div_3)))
	goto L30
L35:
	;
	v167 = *(*int64)(unsafe.Add(mBase, uint32(v14)))
	v171 = base.F64_nearest(base.F64_add(base.F64_mul(v157, float64(1e+06)), base.F64_div(base.F64_convert_i64_s(v167), v13)))
	v174 = int32(0)
	if base.B2i32(base.F64_lt(v171, float64(9.223372036854776e+18)) == v174)|base.B2i32(base.F64_ge(v171, float64(-9.223372036854776e+18)) == v174)|base.B2i32(base.Ui64(int64(9218868437227405312)) < base.Ui64(base.I64_reinterpret_f64(v171)&int64(9223372036854775807))) != 0 {
		goto L4
	} else {
		goto L36
	}
L36:
	;
	v187 = base.I64_trunc_sat_f64_s(v171)
	*(*int64)(unsafe.Add(mBase, uint32(v18))) = v187
	if v91 != int32(2147483647) {
		goto L37
	} else {
		goto L38
	}
L37:
	;
	v191 = int32(-2147483648)
	if base.B2i32(v91 != v191)|base.B2i32(v159 != v191) != 0 {
		goto L3
	} else {
		goto L40
	}
L38:
	;
	goto L39
L39:
	;
	if base.B2i32(v159 != int32(2147483647))|base.B2i32(v187 != int64(9223372036854775807)) != 0 {
		goto L3
	} else {
		goto L42
	}
L40:
	;
	if v187 == int64(-9223372036854775807-1) {
		goto L4
	} else {
		goto L41
	}
L41:
	;
	goto L3
L42:
	;
	goto L4
L43:
	;
	F_errcode(m, int32(134217858))
	mBase = m.M
	v218 = m.ExcPending
	if v218 != 0 {
		goto L1
	} else {
		goto L44
	}
L44:
	;
	F_errmsg(m, int32(_a_F_interval_div_0), int32(0))
	mBase = m.M
	v222 = m.ExcPending
	if v222 != 0 {
		goto L1
	} else {
		goto L45
	}
L45:
	;
	F_errfinish(m, int32(_a_F_interval_div_1), int32(3841), int32(_a_F_interval_div_2))
	mBase = m.M
	v227 = m.ExcPending
	if v227 != 0 {
		goto L1
	} else {
		goto L46
	}
L46:
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
	v110 = *(*int64)(unsafe.Add(mBase, uint32(v64)))
	v111 = *(*int64)(unsafe.Add(mBase, uint32(v14)+8))
	v112 = *(*int64)(unsafe.Add(mBase, uint32(v14)))
	v113 = *(*int64)(unsafe.Add(mBase, uint32(v18)))
	v114 = *(*int64)(unsafe.Add(mBase, uint32(v14)+24))
	v115 = *(*int64)(unsafe.Add(mBase, uint32(v14)+16))
	m.G0 = v14 + int32(32)
	v119 = v113 + v115
	v120 = v110 + v112
	v124 = int64(63)
	v127 = base.I64_extend_i32_u(base.B2i32(base.Ui64(v119) < base.Ui64(v115))) + (v114 + v113>>(uint(v124)%64))
	v133 = base.I64_extend_i32_u(base.B2i32(base.Ui64(v120) < base.Ui64(v112))) + (v111 + v110>>(uint(v124)%64))
	if v133 == v127 {
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
	v110 = *(*int64)(unsafe.Add(mBase, uint32(v64)))
	v111 = *(*int64)(unsafe.Add(mBase, uint32(v14)+8))
	v112 = *(*int64)(unsafe.Add(mBase, uint32(v14)))
	v113 = *(*int64)(unsafe.Add(mBase, uint32(v18)))
	v114 = *(*int64)(unsafe.Add(mBase, uint32(v14)+24))
	v115 = *(*int64)(unsafe.Add(mBase, uint32(v14)+16))
	m.G0 = v14 + int32(32)
	v119 = v113 + v115
	v120 = v110 + v112
	v124 = int64(63)
	v127 = base.I64_extend_i32_u(base.B2i32(base.Ui64(v119) < base.Ui64(v115))) + (v114 + v113>>(uint(v124)%64))
	v133 = base.I64_extend_i32_u(base.B2i32(base.Ui64(v120) < base.Ui64(v112))) + (v111 + v110>>(uint(v124)%64))
	if v133 == v127 {
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
	var v17 int32
	_ = v17
	var v20 int64
	_ = v20
	var v30 int32
	_ = v30
	var v33 int32
	_ = v33
	var v37 int32
	_ = v37
	var v42 int32
	_ = v42
	var v43 int64
	_ = v43
	var v44 int64
	_ = v44
	var v45 int64
	_ = v45
	var v50 int64
	_ = v50
	var v51 int32
	_ = v51
	var v52 int32
	_ = v52
	v4 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v5 = *(*int32)(unsafe.Add(mBase, uint32(v4)+12))
	if v5 != int32(-2147483648) {
		if v5 == int32(2147483647) {
			v17 = *(*int32)(unsafe.Add(mBase, uint32(v4)+8))
			v20 = *(*int64)(unsafe.Add(mBase, uint32(v4)))
			if base.B2i32(v17 != int32(2147483647))|base.B2i32(v20 != int64(9223372036854775807)) != 0 {
				v43 = v20
				v44 = int64(86400000000)
				v45 = base.I64_rem_s(v43, v44)
				if v45 < int64(0) {
					v50 = v45 + v44
				} else {
					v50 = v45
				}
				v51 = F_Int64GetDatum(m, v50)
				mBase = m.M
				v52 = m.ExcPending
				if v52 != 0 {
					return int32(0)
				} else {
					return v51
				}
			} else {
				F_errstart_cold(m, int32(21), int32(0))
				mBase = m.M
				v30 = m.ExcPending
				if v30 != 0 {
					return int32(0)
				} else {
					F_errcode(m, int32(134217858))
					mBase = m.M
					v33 = m.ExcPending
					if v33 != 0 {
						return int32(0)
					} else {
						F_errmsg(m, int32(_a_F_interval_time_0), int32(0))
						mBase = m.M
						v37 = m.ExcPending
						if v37 != 0 {
							return int32(0)
						} else {
							F_errfinish(m, int32(_a_F_interval_time_1), int32(2085), int32(_a_F_interval_time_2))
							mBase = m.M
							v42 = m.ExcPending
							if v42 != 0 {
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
			v10 = *(*int64)(unsafe.Add(mBase, uint32(v4)))
			v43 = v10
			v44 = int64(86400000000)
			v45 = base.I64_rem_s(v43, v44)
			if v45 < int64(0) {
				v50 = v45 + v44
			} else {
				v50 = v45
			}
			v51 = F_Int64GetDatum(m, v50)
			mBase = m.M
			v52 = m.ExcPending
			if v52 != 0 {
				return int32(0)
			} else {
				return v51
			}
		}
	} else {
		v11 = *(*int64)(unsafe.Add(mBase, uint32(v4)))
		v12 = *(*int32)(unsafe.Add(mBase, uint32(v4)+8))
		if v12 != int32(-2147483648) {
			v43 = v11
			v44 = int64(86400000000)
			v45 = base.I64_rem_s(v43, v44)
			if v45 < int64(0) {
				v50 = v45 + v44
			} else {
				v50 = v45
			}
			v51 = F_Int64GetDatum(m, v50)
			mBase = m.M
			v52 = m.ExcPending
			if v52 != 0 {
				return int32(0)
			} else {
				return v51
			}
		} else {
			if v11 == int64(-9223372036854775807-1) {
				F_errstart_cold(m, int32(21), int32(0))
				mBase = m.M
				v30 = m.ExcPending
				if v30 != 0 {
					return int32(0)
				} else {
					F_errcode(m, int32(134217858))
					mBase = m.M
					v33 = m.ExcPending
					if v33 != 0 {
						return int32(0)
					} else {
						F_errmsg(m, int32(_a_F_interval_time_0), int32(0))
						mBase = m.M
						v37 = m.ExcPending
						if v37 != 0 {
							return int32(0)
						} else {
							F_errfinish(m, int32(_a_F_interval_time_1), int32(2085), int32(_a_F_interval_time_2))
							mBase = m.M
							v42 = m.ExcPending
							if v42 != 0 {
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
				v43 = v11
				v44 = int64(86400000000)
				v45 = base.I64_rem_s(v43, v44)
				if v45 < int64(0) {
					v50 = v45 + v44
				} else {
					v50 = v45
				}
				v51 = F_Int64GetDatum(m, v50)
				mBase = m.M
				v52 = m.ExcPending
				if v52 != 0 {
					return int32(0)
				} else {
					return v51
				}
			}
		}
	}
}
