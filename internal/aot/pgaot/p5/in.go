package p5

import (
	base "github.com/shibukawa/pgmem/internal/aot/pgaot/base"
	"math"
	"unsafe"
)

func F_in_range_float8_float8(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v7 int32
	_ = v7
	var v8 float64
	_ = v8
	var v19 int32
	_ = v19
	var v20 int32
	_ = v20
	var v21 float64
	_ = v21
	var v23 int64
	_ = v23
	var v24 int64
	_ = v24
	var v25 int32
	_ = v25
	var v26 float64
	_ = v26
	var v43 int32
	_ = v43
	var v45 float64
	_ = v45
	var v64 float64
	_ = v64
	var v65 float64
	_ = v65
	var v75 int32
	_ = v75
	var v78 int32
	_ = v78
	var v82 int32
	_ = v82
	var v87 int32
	_ = v87
	v7 = *(*int32)(unsafe.Add(mBase, uint32(l0)+36))
	v8 = *(*float64)(unsafe.Add(mBase, uint32(v7)))
	if base.B2i32(base.Ui64(int64(9218868437227405312)) < base.Ui64(base.I64_reinterpret_f64(v8)&int64(9223372036854775807)))|base.F64_lt(v8, float64(0)) == int32(0) {
		v19 = *(*int32)(unsafe.Add(mBase, uint32(l0)+52))
		v20 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
		v21 = *(*float64)(unsafe.Add(mBase, uint32(v20)))
		v23 = int64(9223372036854775807)
		v24 = base.I64_reinterpret_f64(v21) & v23
		v25 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
		v26 = *(*float64)(unsafe.Add(mBase, uint32(v25)))
		if base.Ui64(int64(9218868437227405313)) <= base.Ui64(base.I64_reinterpret_f64(v26)&v23) {
			return base.B2i32(v19 == int32(0)) | base.B2i32(base.Ui64(int64(9218868437227405312)) < base.Ui64(v24))
		} else {
			if base.Ui64(int64(9218868437227405312)) < base.Ui64(v24) {
				return base.B2i32(v19 != int32(0))
			} else {
				v43 = *(*int32)(unsafe.Add(mBase, uint32(l0)+44))
				v45 = math.Float64frombits(uint64(0x7ff0000000000000))
				if base.F64_ne(base.F64_abs(v8), v45)|base.F64_ne(base.F64_abs(v21), v45) != 0 {
					if v43 != 0 {
						v64 = base.F64_neg(v8)
					} else {
						v64 = v8
					}
					v65 = base.F64_add(v21, v64)
					if v19 != 0 {
						return base.F64_ge(v65, v26)
					} else {
						return base.F64_le(v65, v26)
					}
				} else {
					if v43 != 0 {
						if base.F64_gt(v21, float64(0)) == int32(0) {
							if v43 != 0 {
								v64 = base.F64_neg(v8)
							} else {
								v64 = v8
							}
							v65 = base.F64_add(v21, v64)
							if v19 != 0 {
								return base.F64_ge(v65, v26)
							} else {
								return base.F64_le(v65, v26)
							}
						} else {
							return int32(1)
						}
					} else {
						if base.F64_lt(v21, float64(0)) == int32(0) {
							if v43 != 0 {
								v64 = base.F64_neg(v8)
							} else {
								v64 = v8
							}
							v65 = base.F64_add(v21, v64)
							if v19 != 0 {
								return base.F64_ge(v65, v26)
							} else {
								return base.F64_le(v65, v26)
							}
						} else {
							return int32(1)
						}
					}
				}
			}
		}
	} else {
		F_errstart_cold(m, int32(21), int32(0))
		mBase = m.M
		v75 = m.ExcPending
		if v75 != 0 {
			return int32(0)
		} else {
			F_errcode(m, int32(50593922))
			mBase = m.M
			v78 = m.ExcPending
			if v78 != 0 {
				return int32(0)
			} else {
				F_errmsg(m, int32(_a_F_in_range_float8_float8_0), int32(0))
				mBase = m.M
				v82 = m.ExcPending
				if v82 != 0 {
					return int32(0)
				} else {
					F_errfinish(m, int32(_a_F_in_range_float8_float8_1), int32(1043), int32(_a_F_in_range_float8_float8_2))
					mBase = m.M
					v87 = m.ExcPending
					if v87 != 0 {
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
func F_in_range_int2_int2(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v4 int32
	_ = v4
	var v5 int32
	_ = v5
	var v6 int32
	_ = v6
	var v7 int32
	_ = v7
	var v8 int32
	_ = v8
	var v9 int32
	_ = v9
	var v12 int32
	_ = v12
	v4 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v5 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
	v6 = int32(*(*int16)(unsafe.Add(mBase, uint32(l0)+36)))
	v7 = *(*int32)(unsafe.Add(mBase, uint32(l0)+44))
	v8 = *(*int32)(unsafe.Add(mBase, uint32(l0)+52))
	v9 = F_DirectFunctionCall5Coll(m, int32(1303), int32(0), v4, v5, v6, v7, v8)
	mBase = m.M
	v12 = m.ExcPending
	if v12 != 0 {
		return int32(0)
	} else {
		return v9
	}
}
func F_in_range_int4_int2(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v4 int32
	_ = v4
	var v5 int32
	_ = v5
	var v6 int32
	_ = v6
	var v7 int32
	_ = v7
	var v8 int32
	_ = v8
	var v9 int32
	_ = v9
	var v12 int32
	_ = v12
	v4 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v5 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
	v6 = int32(*(*int16)(unsafe.Add(mBase, uint32(l0)+36)))
	v7 = *(*int32)(unsafe.Add(mBase, uint32(l0)+44))
	v8 = *(*int32)(unsafe.Add(mBase, uint32(l0)+52))
	v9 = F_DirectFunctionCall5Coll(m, int32(1302), int32(0), v4, v5, v6, v7, v8)
	mBase = m.M
	v12 = m.ExcPending
	if v12 != 0 {
		return int32(0)
	} else {
		return v9
	}
}
func F_in_range_int4_int4(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v6 int32
	_ = v6
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
	var v20 int32
	_ = v20
	var v26 int32
	_ = v26
	var v36 int32
	_ = v36
	var v39 int32
	_ = v39
	var v43 int32
	_ = v43
	var v48 int32
	_ = v48
	v6 = *(*int32)(unsafe.Add(mBase, uint32(l0)+36))
	if int32(0) <= v6 {
		v9 = *(*int32)(unsafe.Add(mBase, uint32(l0)+52))
		v12 = *(*int32)(unsafe.Add(mBase, uint32(l0)+44))
		if v12 != 0 {
			v13 = int32(0) - v6
		} else {
			v13 = v6
		}
		v16 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
		v17 = v16 + v13
		if base.B2i32(v13 < int32(0)) != base.B2i32(v17 < v16) {
			v20 = int32(0)
			return base.B2i32(v12 != v20) ^ base.B2i32(v9 != v20)
		} else {
			v26 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
			if v9 != 0 {
				return base.B2i32(v26 <= v17)
			} else {
				return base.B2i32(v17 <= v26)
			}
		}
	} else {
		F_errstart_cold(m, int32(21), int32(0))
		mBase = m.M
		v36 = m.ExcPending
		if v36 != 0 {
			return int32(0)
		} else {
			F_errcode(m, int32(50593922))
			mBase = m.M
			v39 = m.ExcPending
			if v39 != 0 {
				return int32(0)
			} else {
				F_errmsg(m, int32(_a_F_in_range_int4_int4_0), int32(0))
				mBase = m.M
				v43 = m.ExcPending
				if v43 != 0 {
					return int32(0)
				} else {
					F_errfinish(m, int32(_a_F_in_range_int4_int4_1), int32(664), int32(_a_F_in_range_int4_int4_2))
					mBase = m.M
					v48 = m.ExcPending
					if v48 != 0 {
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
func F_in_range_int4_int8(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v6 int32
	_ = v6
	var v7 int64
	_ = v7
	var v10 int32
	_ = v10
	var v13 int32
	_ = v13
	var v14 int64
	_ = v14
	var v17 int64
	_ = v17
	var v18 int64
	_ = v18
	var v21 int32
	_ = v21
	var v27 int64
	_ = v27
	var v37 int32
	_ = v37
	var v40 int32
	_ = v40
	var v44 int32
	_ = v44
	var v49 int32
	_ = v49
	v6 = *(*int32)(unsafe.Add(mBase, uint32(l0)+36))
	v7 = *(*int64)(unsafe.Add(mBase, uint32(v6)))
	if int64(0) <= v7 {
		v10 = *(*int32)(unsafe.Add(mBase, uint32(l0)+52))
		v13 = *(*int32)(unsafe.Add(mBase, uint32(l0)+44))
		if v13 != 0 {
			v14 = int64(0) - v7
		} else {
			v14 = v7
		}
		v17 = int64(*(*int32)(unsafe.Add(mBase, uint32(l0)+28)))
		v18 = v17 + v14
		if base.B2i32(v14 < int64(0)) != base.B2i32(v18 < v17) {
			v21 = int32(0)
			return base.B2i32(v13 != v21) ^ base.B2i32(v10 != v21)
		} else {
			v27 = int64(*(*int32)(unsafe.Add(mBase, uint32(l0)+20)))
			if v10 != 0 {
				return base.B2i32(v27 <= v18)
			} else {
				return base.B2i32(v18 <= v27)
			}
		}
	} else {
		F_errstart_cold(m, int32(21), int32(0))
		mBase = m.M
		v37 = m.ExcPending
		if v37 != 0 {
			return int32(0)
		} else {
			F_errcode(m, int32(50593922))
			mBase = m.M
			v40 = m.ExcPending
			if v40 != 0 {
				return int32(0)
			} else {
				F_errmsg(m, int32(_a_F_in_range_int4_int8_0), int32(0))
				mBase = m.M
				v44 = m.ExcPending
				if v44 != 0 {
					return int32(0)
				} else {
					F_errfinish(m, int32(_a_F_in_range_int4_int8_1), int32(711), int32(_a_F_in_range_int4_int8_2))
					mBase = m.M
					v49 = m.ExcPending
					if v49 != 0 {
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
func F_in_range_interval_interval(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v14 int32
	_ = v14
	var v16 int32
	_ = v16
	var v19 int32
	_ = v19
	var v20 int32
	_ = v20
	var v21 int32
	_ = v21
	var v25 int32
	_ = v25
	var v27 int64
	_ = v27
	var v36 int64
	_ = v36
	var v37 int64
	_ = v37
	var v39 int64
	_ = v39
	var v42 int64
	_ = v42
	var v43 int64
	_ = v43
	var v45 int64
	_ = v45
	var v46 int64
	_ = v46
	var v50 int64
	_ = v50
	var v57 int64
	_ = v57
	var v68 int64
	_ = v68
	var v69 int64
	_ = v69
	var v73 int64
	_ = v73
	var v80 int32
	_ = v80
	var v81 int32
	_ = v81
	var v82 int32
	_ = v82
	var v83 int32
	_ = v83
	var v84 int32
	_ = v84
	var v94 int32
	_ = v94
	var v95 int32
	_ = v95
	var v98 int32
	_ = v98
	var v101 int64
	_ = v101
	var v107 int32
	_ = v107
	var v110 int64
	_ = v110
	var v118 int32
	_ = v118
	var v121 int32
	_ = v121
	var v124 int32
	_ = v124
	var v125 int64
	_ = v125
	var v126 int64
	_ = v126
	var v129 int64
	_ = v129
	var v138 int64
	_ = v138
	var v139 int64
	_ = v139
	var v141 int64
	_ = v141
	var v144 int64
	_ = v144
	var v145 int64
	_ = v145
	var v147 int64
	_ = v147
	var v148 int64
	_ = v148
	var v152 int64
	_ = v152
	var v159 int64
	_ = v159
	var v171 int32
	_ = v171
	var v172 int64
	_ = v172
	var v173 int64
	_ = v173
	var v176 int64
	_ = v176
	var v185 int64
	_ = v185
	var v186 int64
	_ = v186
	var v188 int64
	_ = v188
	var v191 int64
	_ = v191
	var v192 int64
	_ = v192
	var v194 int64
	_ = v194
	var v195 int64
	_ = v195
	var v199 int64
	_ = v199
	var v206 int64
	_ = v206
	var v217 int64
	_ = v217
	var v218 int64
	_ = v218
	var v219 int64
	_ = v219
	var v222 int64
	_ = v222
	var v223 int64
	_ = v223
	var v226 int64
	_ = v226
	var v227 int64
	_ = v227
	var v228 int64
	_ = v228
	var v232 int64
	_ = v232
	var v233 int64
	_ = v233
	var v236 int64
	_ = v236
	var v239 int32
	_ = v239
	var v240 int32
	_ = v240
	var v243 int32
	_ = v243
	var v253 int32
	_ = v253
	var v261 int32
	_ = v261
	var v264 int32
	_ = v264
	var v268 int32
	_ = v268
	var v273 int32
	_ = v273
	v14 = m.G0
	v16 = v14 - int32(48)
	m.G0 = v16
	v19 = v16 + int32(32)
	v20 = *(*int32)(unsafe.Add(mBase, uint32(l0)+36))
	v21 = *(*int32)(unsafe.Add(mBase, uint32(v20)+12))
	v25 = *(*int32)(unsafe.Add(mBase, uint32(v20)+8))
	v27 = base.I64_extend_i32_s(v21)*int64(30) + base.I64_extend_i32_s(v25)
	v36 = int64(32)
	v37 = int64(20)
	v39 = int64(base.Ui64(v27) >> (uint(v36) % 64))
	v42 = int64(4294967295)
	v43 = int64(500654080)
	v45 = v27 & v42
	v46 = v43 * v45
	v50 = int64(base.Ui64(v46)>>(uint(v36)%64)) + v43*v39
	v57 = v45*v37 + v50&v42
	*(*int64)(unsafe.Add(mBase, uint32(v19)+8)) = v27*int64(0) + v27>>(uint(int64(63))%64)*int64(86400000000) + v37*v39 + int64(base.Ui64(v50)>>(uint(v36)%64)) + int64(base.Ui64(v57)>>(uint(v36)%64))
	*(*int64)(unsafe.Add(mBase, uint32(v19))) = v46&v42 | v57<<(uint(v36)%64)
	v68 = *(*int64)(unsafe.Add(mBase, uint32(v16)+40))
	v69 = *(*int64)(unsafe.Add(mBase, uint32(v20)))
	v73 = *(*int64)(unsafe.Add(mBase, uint32(v16)+32))
	if int64(0) <= v68+v69>>(uint(int64(63))%64)+base.I64_extend_i32_u(base.B2i32(base.Ui64(v73+v69) < base.Ui64(v73))) {
		v80 = *(*int32)(unsafe.Add(mBase, uint32(l0)+52))
		v81 = *(*int32)(unsafe.Add(mBase, uint32(l0)+44))
		v82 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
		v83 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
		v84 = int32(2147483647)
		if base.B2i32(v21 != v84)|base.B2i32(v25 != v84)|base.B2i32(v69 != int64(9223372036854775807)) == int32(0) {
			v94 = *(*int32)(unsafe.Add(mBase, uint32(v82)+12))
			if v81 != 0 {
				v95 = int32(1460)
				if v94 != int32(2147483647) {
					v118 = v95
					v121 = F_DirectFunctionCall2Coll(m, v118, int32(0), v82, v20)
					mBase = m.M
					v124 = m.ExcPending
					if v124 != 0 {
						return int32(0)
					} else {
						v125 = int64(*(*int32)(unsafe.Add(mBase, uint32(v83)+8)))
						v126 = int64(*(*int32)(unsafe.Add(mBase, uint32(v83)+12)))
						v129 = v125 + v126*int64(30)
						v138 = int64(32)
						v139 = int64(20)
						v141 = int64(base.Ui64(v129) >> (uint(v138) % 64))
						v144 = int64(4294967295)
						v145 = int64(500654080)
						v147 = v129 & v144
						v148 = v145 * v147
						v152 = int64(base.Ui64(v148)>>(uint(v138)%64)) + v145*v141
						v159 = v147*v139 + v152&v144
						*(*int64)(unsafe.Add(mBase, uint32(v16)+8)) = v129*int64(0) + v129>>(uint(int64(63))%64)*int64(86400000000) + v139*v141 + int64(base.Ui64(v152)>>(uint(v138)%64)) + int64(base.Ui64(v159)>>(uint(v138)%64))
						*(*int64)(unsafe.Add(mBase, uint32(v16))) = v148&v144 | v159<<(uint(v138)%64)
						v171 = v16 + int32(16)
						v172 = int64(*(*int32)(unsafe.Add(mBase, uint32(v121)+8)))
						v173 = int64(*(*int32)(unsafe.Add(mBase, uint32(v121)+12)))
						v176 = v172 + v173*int64(30)
						v185 = int64(32)
						v186 = int64(20)
						v188 = int64(base.Ui64(v176) >> (uint(v185) % 64))
						v191 = int64(4294967295)
						v192 = int64(500654080)
						v194 = v176 & v191
						v195 = v192 * v194
						v199 = int64(base.Ui64(v195)>>(uint(v185)%64)) + v192*v188
						v206 = v194*v186 + v199&v191
						*(*int64)(unsafe.Add(mBase, uint32(v171)+8)) = v176*int64(0) + v176>>(uint(int64(63))%64)*int64(86400000000) + v186*v188 + int64(base.Ui64(v199)>>(uint(v185)%64)) + int64(base.Ui64(v206)>>(uint(v185)%64))
						*(*int64)(unsafe.Add(mBase, uint32(v171))) = v195&v191 | v206<<(uint(v185)%64)
						v217 = *(*int64)(unsafe.Add(mBase, uint32(v16)+24))
						v218 = *(*int64)(unsafe.Add(mBase, uint32(v121)))
						v219 = int64(63)
						v222 = *(*int64)(unsafe.Add(mBase, uint32(v16)+16))
						v223 = v218 + v222
						v226 = v217 + v218>>(uint(v219)%64) + base.I64_extend_i32_u(base.B2i32(base.Ui64(v223) < base.Ui64(v222)))
						v227 = *(*int64)(unsafe.Add(mBase, uint32(v16)+8))
						v228 = *(*int64)(unsafe.Add(mBase, uint32(v83)))
						v232 = *(*int64)(unsafe.Add(mBase, uint32(v16)))
						v233 = v228 + v232
						v236 = v227 + v228>>(uint(v219)%64) + base.I64_extend_i32_u(base.B2i32(base.Ui64(v233) < base.Ui64(v232)))
						v239 = base.B2i32(v226 == v236)
						if v226 == v236 {
							v240 = base.B2i32(base.Ui64(v233) <= base.Ui64(v223))
						} else {
							v240 = base.B2i32(v236 <= v226)
						}
						if v80 != 0 {
							v253 = v240
						} else {
							if v226 == v236 {
								v243 = base.B2i32(base.Ui64(v223) <= base.Ui64(v233))
							} else {
								v243 = base.B2i32(v226 <= v236)
							}
							v253 = v243
						}
						m.G0 = v16 + int32(48)
						return v253
					}
				} else {
					v98 = *(*int32)(unsafe.Add(mBase, uint32(v82)+8))
					if v98 != int32(2147483647) {
						v118 = v95
						v121 = F_DirectFunctionCall2Coll(m, v118, int32(0), v82, v20)
						mBase = m.M
						v124 = m.ExcPending
						if v124 != 0 {
							return int32(0)
						} else {
							v125 = int64(*(*int32)(unsafe.Add(mBase, uint32(v83)+8)))
							v126 = int64(*(*int32)(unsafe.Add(mBase, uint32(v83)+12)))
							v129 = v125 + v126*int64(30)
							v138 = int64(32)
							v139 = int64(20)
							v141 = int64(base.Ui64(v129) >> (uint(v138) % 64))
							v144 = int64(4294967295)
							v145 = int64(500654080)
							v147 = v129 & v144
							v148 = v145 * v147
							v152 = int64(base.Ui64(v148)>>(uint(v138)%64)) + v145*v141
							v159 = v147*v139 + v152&v144
							*(*int64)(unsafe.Add(mBase, uint32(v16)+8)) = v129*int64(0) + v129>>(uint(int64(63))%64)*int64(86400000000) + v139*v141 + int64(base.Ui64(v152)>>(uint(v138)%64)) + int64(base.Ui64(v159)>>(uint(v138)%64))
							*(*int64)(unsafe.Add(mBase, uint32(v16))) = v148&v144 | v159<<(uint(v138)%64)
							v171 = v16 + int32(16)
							v172 = int64(*(*int32)(unsafe.Add(mBase, uint32(v121)+8)))
							v173 = int64(*(*int32)(unsafe.Add(mBase, uint32(v121)+12)))
							v176 = v172 + v173*int64(30)
							v185 = int64(32)
							v186 = int64(20)
							v188 = int64(base.Ui64(v176) >> (uint(v185) % 64))
							v191 = int64(4294967295)
							v192 = int64(500654080)
							v194 = v176 & v191
							v195 = v192 * v194
							v199 = int64(base.Ui64(v195)>>(uint(v185)%64)) + v192*v188
							v206 = v194*v186 + v199&v191
							*(*int64)(unsafe.Add(mBase, uint32(v171)+8)) = v176*int64(0) + v176>>(uint(int64(63))%64)*int64(86400000000) + v186*v188 + int64(base.Ui64(v199)>>(uint(v185)%64)) + int64(base.Ui64(v206)>>(uint(v185)%64))
							*(*int64)(unsafe.Add(mBase, uint32(v171))) = v195&v191 | v206<<(uint(v185)%64)
							v217 = *(*int64)(unsafe.Add(mBase, uint32(v16)+24))
							v218 = *(*int64)(unsafe.Add(mBase, uint32(v121)))
							v219 = int64(63)
							v222 = *(*int64)(unsafe.Add(mBase, uint32(v16)+16))
							v223 = v218 + v222
							v226 = v217 + v218>>(uint(v219)%64) + base.I64_extend_i32_u(base.B2i32(base.Ui64(v223) < base.Ui64(v222)))
							v227 = *(*int64)(unsafe.Add(mBase, uint32(v16)+8))
							v228 = *(*int64)(unsafe.Add(mBase, uint32(v83)))
							v232 = *(*int64)(unsafe.Add(mBase, uint32(v16)))
							v233 = v228 + v232
							v236 = v227 + v228>>(uint(v219)%64) + base.I64_extend_i32_u(base.B2i32(base.Ui64(v233) < base.Ui64(v232)))
							v239 = base.B2i32(v226 == v236)
							if v226 == v236 {
								v240 = base.B2i32(base.Ui64(v233) <= base.Ui64(v223))
							} else {
								v240 = base.B2i32(v236 <= v226)
							}
							if v80 != 0 {
								v253 = v240
							} else {
								if v226 == v236 {
									v243 = base.B2i32(base.Ui64(v223) <= base.Ui64(v233))
								} else {
									v243 = base.B2i32(v226 <= v236)
								}
								v253 = v243
							}
							m.G0 = v16 + int32(48)
							return v253
						}
					} else {
						v101 = *(*int64)(unsafe.Add(mBase, uint32(v82)))
						if v101 != int64(9223372036854775807) {
							v118 = v95
							v121 = F_DirectFunctionCall2Coll(m, v118, int32(0), v82, v20)
							mBase = m.M
							v124 = m.ExcPending
							if v124 != 0 {
								return int32(0)
							} else {
								v125 = int64(*(*int32)(unsafe.Add(mBase, uint32(v83)+8)))
								v126 = int64(*(*int32)(unsafe.Add(mBase, uint32(v83)+12)))
								v129 = v125 + v126*int64(30)
								v138 = int64(32)
								v139 = int64(20)
								v141 = int64(base.Ui64(v129) >> (uint(v138) % 64))
								v144 = int64(4294967295)
								v145 = int64(500654080)
								v147 = v129 & v144
								v148 = v145 * v147
								v152 = int64(base.Ui64(v148)>>(uint(v138)%64)) + v145*v141
								v159 = v147*v139 + v152&v144
								*(*int64)(unsafe.Add(mBase, uint32(v16)+8)) = v129*int64(0) + v129>>(uint(int64(63))%64)*int64(86400000000) + v139*v141 + int64(base.Ui64(v152)>>(uint(v138)%64)) + int64(base.Ui64(v159)>>(uint(v138)%64))
								*(*int64)(unsafe.Add(mBase, uint32(v16))) = v148&v144 | v159<<(uint(v138)%64)
								v171 = v16 + int32(16)
								v172 = int64(*(*int32)(unsafe.Add(mBase, uint32(v121)+8)))
								v173 = int64(*(*int32)(unsafe.Add(mBase, uint32(v121)+12)))
								v176 = v172 + v173*int64(30)
								v185 = int64(32)
								v186 = int64(20)
								v188 = int64(base.Ui64(v176) >> (uint(v185) % 64))
								v191 = int64(4294967295)
								v192 = int64(500654080)
								v194 = v176 & v191
								v195 = v192 * v194
								v199 = int64(base.Ui64(v195)>>(uint(v185)%64)) + v192*v188
								v206 = v194*v186 + v199&v191
								*(*int64)(unsafe.Add(mBase, uint32(v171)+8)) = v176*int64(0) + v176>>(uint(int64(63))%64)*int64(86400000000) + v186*v188 + int64(base.Ui64(v199)>>(uint(v185)%64)) + int64(base.Ui64(v206)>>(uint(v185)%64))
								*(*int64)(unsafe.Add(mBase, uint32(v171))) = v195&v191 | v206<<(uint(v185)%64)
								v217 = *(*int64)(unsafe.Add(mBase, uint32(v16)+24))
								v218 = *(*int64)(unsafe.Add(mBase, uint32(v121)))
								v219 = int64(63)
								v222 = *(*int64)(unsafe.Add(mBase, uint32(v16)+16))
								v223 = v218 + v222
								v226 = v217 + v218>>(uint(v219)%64) + base.I64_extend_i32_u(base.B2i32(base.Ui64(v223) < base.Ui64(v222)))
								v227 = *(*int64)(unsafe.Add(mBase, uint32(v16)+8))
								v228 = *(*int64)(unsafe.Add(mBase, uint32(v83)))
								v232 = *(*int64)(unsafe.Add(mBase, uint32(v16)))
								v233 = v228 + v232
								v236 = v227 + v228>>(uint(v219)%64) + base.I64_extend_i32_u(base.B2i32(base.Ui64(v233) < base.Ui64(v232)))
								v239 = base.B2i32(v226 == v236)
								if v226 == v236 {
									v240 = base.B2i32(base.Ui64(v233) <= base.Ui64(v223))
								} else {
									v240 = base.B2i32(v236 <= v226)
								}
								if v80 != 0 {
									v253 = v240
								} else {
									if v226 == v236 {
										v243 = base.B2i32(base.Ui64(v223) <= base.Ui64(v233))
									} else {
										v243 = base.B2i32(v226 <= v236)
									}
									v253 = v243
								}
								m.G0 = v16 + int32(48)
								return v253
							}
						} else {
							v253 = int32(1)
							m.G0 = v16 + int32(48)
							return v253
						}
					}
				}
			} else {
				if v94 != int32(-2147483648) {
					v118 = int32(1458)
					v121 = F_DirectFunctionCall2Coll(m, v118, int32(0), v82, v20)
					mBase = m.M
					v124 = m.ExcPending
					if v124 != 0 {
						return int32(0)
					} else {
						v125 = int64(*(*int32)(unsafe.Add(mBase, uint32(v83)+8)))
						v126 = int64(*(*int32)(unsafe.Add(mBase, uint32(v83)+12)))
						v129 = v125 + v126*int64(30)
						v138 = int64(32)
						v139 = int64(20)
						v141 = int64(base.Ui64(v129) >> (uint(v138) % 64))
						v144 = int64(4294967295)
						v145 = int64(500654080)
						v147 = v129 & v144
						v148 = v145 * v147
						v152 = int64(base.Ui64(v148)>>(uint(v138)%64)) + v145*v141
						v159 = v147*v139 + v152&v144
						*(*int64)(unsafe.Add(mBase, uint32(v16)+8)) = v129*int64(0) + v129>>(uint(int64(63))%64)*int64(86400000000) + v139*v141 + int64(base.Ui64(v152)>>(uint(v138)%64)) + int64(base.Ui64(v159)>>(uint(v138)%64))
						*(*int64)(unsafe.Add(mBase, uint32(v16))) = v148&v144 | v159<<(uint(v138)%64)
						v171 = v16 + int32(16)
						v172 = int64(*(*int32)(unsafe.Add(mBase, uint32(v121)+8)))
						v173 = int64(*(*int32)(unsafe.Add(mBase, uint32(v121)+12)))
						v176 = v172 + v173*int64(30)
						v185 = int64(32)
						v186 = int64(20)
						v188 = int64(base.Ui64(v176) >> (uint(v185) % 64))
						v191 = int64(4294967295)
						v192 = int64(500654080)
						v194 = v176 & v191
						v195 = v192 * v194
						v199 = int64(base.Ui64(v195)>>(uint(v185)%64)) + v192*v188
						v206 = v194*v186 + v199&v191
						*(*int64)(unsafe.Add(mBase, uint32(v171)+8)) = v176*int64(0) + v176>>(uint(int64(63))%64)*int64(86400000000) + v186*v188 + int64(base.Ui64(v199)>>(uint(v185)%64)) + int64(base.Ui64(v206)>>(uint(v185)%64))
						*(*int64)(unsafe.Add(mBase, uint32(v171))) = v195&v191 | v206<<(uint(v185)%64)
						v217 = *(*int64)(unsafe.Add(mBase, uint32(v16)+24))
						v218 = *(*int64)(unsafe.Add(mBase, uint32(v121)))
						v219 = int64(63)
						v222 = *(*int64)(unsafe.Add(mBase, uint32(v16)+16))
						v223 = v218 + v222
						v226 = v217 + v218>>(uint(v219)%64) + base.I64_extend_i32_u(base.B2i32(base.Ui64(v223) < base.Ui64(v222)))
						v227 = *(*int64)(unsafe.Add(mBase, uint32(v16)+8))
						v228 = *(*int64)(unsafe.Add(mBase, uint32(v83)))
						v232 = *(*int64)(unsafe.Add(mBase, uint32(v16)))
						v233 = v228 + v232
						v236 = v227 + v228>>(uint(v219)%64) + base.I64_extend_i32_u(base.B2i32(base.Ui64(v233) < base.Ui64(v232)))
						v239 = base.B2i32(v226 == v236)
						if v226 == v236 {
							v240 = base.B2i32(base.Ui64(v233) <= base.Ui64(v223))
						} else {
							v240 = base.B2i32(v236 <= v226)
						}
						if v80 != 0 {
							v253 = v240
						} else {
							if v226 == v236 {
								v243 = base.B2i32(base.Ui64(v223) <= base.Ui64(v233))
							} else {
								v243 = base.B2i32(v226 <= v236)
							}
							v253 = v243
						}
						m.G0 = v16 + int32(48)
						return v253
					}
				} else {
					v107 = *(*int32)(unsafe.Add(mBase, uint32(v82)+8))
					if v107 != int32(-2147483648) {
						v118 = int32(1458)
						v121 = F_DirectFunctionCall2Coll(m, v118, int32(0), v82, v20)
						mBase = m.M
						v124 = m.ExcPending
						if v124 != 0 {
							return int32(0)
						} else {
							v125 = int64(*(*int32)(unsafe.Add(mBase, uint32(v83)+8)))
							v126 = int64(*(*int32)(unsafe.Add(mBase, uint32(v83)+12)))
							v129 = v125 + v126*int64(30)
							v138 = int64(32)
							v139 = int64(20)
							v141 = int64(base.Ui64(v129) >> (uint(v138) % 64))
							v144 = int64(4294967295)
							v145 = int64(500654080)
							v147 = v129 & v144
							v148 = v145 * v147
							v152 = int64(base.Ui64(v148)>>(uint(v138)%64)) + v145*v141
							v159 = v147*v139 + v152&v144
							*(*int64)(unsafe.Add(mBase, uint32(v16)+8)) = v129*int64(0) + v129>>(uint(int64(63))%64)*int64(86400000000) + v139*v141 + int64(base.Ui64(v152)>>(uint(v138)%64)) + int64(base.Ui64(v159)>>(uint(v138)%64))
							*(*int64)(unsafe.Add(mBase, uint32(v16))) = v148&v144 | v159<<(uint(v138)%64)
							v171 = v16 + int32(16)
							v172 = int64(*(*int32)(unsafe.Add(mBase, uint32(v121)+8)))
							v173 = int64(*(*int32)(unsafe.Add(mBase, uint32(v121)+12)))
							v176 = v172 + v173*int64(30)
							v185 = int64(32)
							v186 = int64(20)
							v188 = int64(base.Ui64(v176) >> (uint(v185) % 64))
							v191 = int64(4294967295)
							v192 = int64(500654080)
							v194 = v176 & v191
							v195 = v192 * v194
							v199 = int64(base.Ui64(v195)>>(uint(v185)%64)) + v192*v188
							v206 = v194*v186 + v199&v191
							*(*int64)(unsafe.Add(mBase, uint32(v171)+8)) = v176*int64(0) + v176>>(uint(int64(63))%64)*int64(86400000000) + v186*v188 + int64(base.Ui64(v199)>>(uint(v185)%64)) + int64(base.Ui64(v206)>>(uint(v185)%64))
							*(*int64)(unsafe.Add(mBase, uint32(v171))) = v195&v191 | v206<<(uint(v185)%64)
							v217 = *(*int64)(unsafe.Add(mBase, uint32(v16)+24))
							v218 = *(*int64)(unsafe.Add(mBase, uint32(v121)))
							v219 = int64(63)
							v222 = *(*int64)(unsafe.Add(mBase, uint32(v16)+16))
							v223 = v218 + v222
							v226 = v217 + v218>>(uint(v219)%64) + base.I64_extend_i32_u(base.B2i32(base.Ui64(v223) < base.Ui64(v222)))
							v227 = *(*int64)(unsafe.Add(mBase, uint32(v16)+8))
							v228 = *(*int64)(unsafe.Add(mBase, uint32(v83)))
							v232 = *(*int64)(unsafe.Add(mBase, uint32(v16)))
							v233 = v228 + v232
							v236 = v227 + v228>>(uint(v219)%64) + base.I64_extend_i32_u(base.B2i32(base.Ui64(v233) < base.Ui64(v232)))
							v239 = base.B2i32(v226 == v236)
							if v226 == v236 {
								v240 = base.B2i32(base.Ui64(v233) <= base.Ui64(v223))
							} else {
								v240 = base.B2i32(v236 <= v226)
							}
							if v80 != 0 {
								v253 = v240
							} else {
								if v226 == v236 {
									v243 = base.B2i32(base.Ui64(v223) <= base.Ui64(v233))
								} else {
									v243 = base.B2i32(v226 <= v236)
								}
								v253 = v243
							}
							m.G0 = v16 + int32(48)
							return v253
						}
					} else {
						v110 = *(*int64)(unsafe.Add(mBase, uint32(v82)))
						if v110 != int64(-9223372036854775807-1) {
							v118 = int32(1458)
							v121 = F_DirectFunctionCall2Coll(m, v118, int32(0), v82, v20)
							mBase = m.M
							v124 = m.ExcPending
							if v124 != 0 {
								return int32(0)
							} else {
								v125 = int64(*(*int32)(unsafe.Add(mBase, uint32(v83)+8)))
								v126 = int64(*(*int32)(unsafe.Add(mBase, uint32(v83)+12)))
								v129 = v125 + v126*int64(30)
								v138 = int64(32)
								v139 = int64(20)
								v141 = int64(base.Ui64(v129) >> (uint(v138) % 64))
								v144 = int64(4294967295)
								v145 = int64(500654080)
								v147 = v129 & v144
								v148 = v145 * v147
								v152 = int64(base.Ui64(v148)>>(uint(v138)%64)) + v145*v141
								v159 = v147*v139 + v152&v144
								*(*int64)(unsafe.Add(mBase, uint32(v16)+8)) = v129*int64(0) + v129>>(uint(int64(63))%64)*int64(86400000000) + v139*v141 + int64(base.Ui64(v152)>>(uint(v138)%64)) + int64(base.Ui64(v159)>>(uint(v138)%64))
								*(*int64)(unsafe.Add(mBase, uint32(v16))) = v148&v144 | v159<<(uint(v138)%64)
								v171 = v16 + int32(16)
								v172 = int64(*(*int32)(unsafe.Add(mBase, uint32(v121)+8)))
								v173 = int64(*(*int32)(unsafe.Add(mBase, uint32(v121)+12)))
								v176 = v172 + v173*int64(30)
								v185 = int64(32)
								v186 = int64(20)
								v188 = int64(base.Ui64(v176) >> (uint(v185) % 64))
								v191 = int64(4294967295)
								v192 = int64(500654080)
								v194 = v176 & v191
								v195 = v192 * v194
								v199 = int64(base.Ui64(v195)>>(uint(v185)%64)) + v192*v188
								v206 = v194*v186 + v199&v191
								*(*int64)(unsafe.Add(mBase, uint32(v171)+8)) = v176*int64(0) + v176>>(uint(int64(63))%64)*int64(86400000000) + v186*v188 + int64(base.Ui64(v199)>>(uint(v185)%64)) + int64(base.Ui64(v206)>>(uint(v185)%64))
								*(*int64)(unsafe.Add(mBase, uint32(v171))) = v195&v191 | v206<<(uint(v185)%64)
								v217 = *(*int64)(unsafe.Add(mBase, uint32(v16)+24))
								v218 = *(*int64)(unsafe.Add(mBase, uint32(v121)))
								v219 = int64(63)
								v222 = *(*int64)(unsafe.Add(mBase, uint32(v16)+16))
								v223 = v218 + v222
								v226 = v217 + v218>>(uint(v219)%64) + base.I64_extend_i32_u(base.B2i32(base.Ui64(v223) < base.Ui64(v222)))
								v227 = *(*int64)(unsafe.Add(mBase, uint32(v16)+8))
								v228 = *(*int64)(unsafe.Add(mBase, uint32(v83)))
								v232 = *(*int64)(unsafe.Add(mBase, uint32(v16)))
								v233 = v228 + v232
								v236 = v227 + v228>>(uint(v219)%64) + base.I64_extend_i32_u(base.B2i32(base.Ui64(v233) < base.Ui64(v232)))
								v239 = base.B2i32(v226 == v236)
								if v226 == v236 {
									v240 = base.B2i32(base.Ui64(v233) <= base.Ui64(v223))
								} else {
									v240 = base.B2i32(v236 <= v226)
								}
								if v80 != 0 {
									v253 = v240
								} else {
									if v226 == v236 {
										v243 = base.B2i32(base.Ui64(v223) <= base.Ui64(v233))
									} else {
										v243 = base.B2i32(v226 <= v236)
									}
									v253 = v243
								}
								m.G0 = v16 + int32(48)
								return v253
							}
						} else {
							v253 = int32(1)
							m.G0 = v16 + int32(48)
							return v253
						}
					}
				}
			}
		} else {
			if v81 != 0 {
				v118 = int32(1460)
			} else {
				v118 = int32(1458)
			}
			v121 = F_DirectFunctionCall2Coll(m, v118, int32(0), v82, v20)
			mBase = m.M
			v124 = m.ExcPending
			if v124 != 0 {
				return int32(0)
			} else {
				v125 = int64(*(*int32)(unsafe.Add(mBase, uint32(v83)+8)))
				v126 = int64(*(*int32)(unsafe.Add(mBase, uint32(v83)+12)))
				v129 = v125 + v126*int64(30)
				v138 = int64(32)
				v139 = int64(20)
				v141 = int64(base.Ui64(v129) >> (uint(v138) % 64))
				v144 = int64(4294967295)
				v145 = int64(500654080)
				v147 = v129 & v144
				v148 = v145 * v147
				v152 = int64(base.Ui64(v148)>>(uint(v138)%64)) + v145*v141
				v159 = v147*v139 + v152&v144
				*(*int64)(unsafe.Add(mBase, uint32(v16)+8)) = v129*int64(0) + v129>>(uint(int64(63))%64)*int64(86400000000) + v139*v141 + int64(base.Ui64(v152)>>(uint(v138)%64)) + int64(base.Ui64(v159)>>(uint(v138)%64))
				*(*int64)(unsafe.Add(mBase, uint32(v16))) = v148&v144 | v159<<(uint(v138)%64)
				v171 = v16 + int32(16)
				v172 = int64(*(*int32)(unsafe.Add(mBase, uint32(v121)+8)))
				v173 = int64(*(*int32)(unsafe.Add(mBase, uint32(v121)+12)))
				v176 = v172 + v173*int64(30)
				v185 = int64(32)
				v186 = int64(20)
				v188 = int64(base.Ui64(v176) >> (uint(v185) % 64))
				v191 = int64(4294967295)
				v192 = int64(500654080)
				v194 = v176 & v191
				v195 = v192 * v194
				v199 = int64(base.Ui64(v195)>>(uint(v185)%64)) + v192*v188
				v206 = v194*v186 + v199&v191
				*(*int64)(unsafe.Add(mBase, uint32(v171)+8)) = v176*int64(0) + v176>>(uint(int64(63))%64)*int64(86400000000) + v186*v188 + int64(base.Ui64(v199)>>(uint(v185)%64)) + int64(base.Ui64(v206)>>(uint(v185)%64))
				*(*int64)(unsafe.Add(mBase, uint32(v171))) = v195&v191 | v206<<(uint(v185)%64)
				v217 = *(*int64)(unsafe.Add(mBase, uint32(v16)+24))
				v218 = *(*int64)(unsafe.Add(mBase, uint32(v121)))
				v219 = int64(63)
				v222 = *(*int64)(unsafe.Add(mBase, uint32(v16)+16))
				v223 = v218 + v222
				v226 = v217 + v218>>(uint(v219)%64) + base.I64_extend_i32_u(base.B2i32(base.Ui64(v223) < base.Ui64(v222)))
				v227 = *(*int64)(unsafe.Add(mBase, uint32(v16)+8))
				v228 = *(*int64)(unsafe.Add(mBase, uint32(v83)))
				v232 = *(*int64)(unsafe.Add(mBase, uint32(v16)))
				v233 = v228 + v232
				v236 = v227 + v228>>(uint(v219)%64) + base.I64_extend_i32_u(base.B2i32(base.Ui64(v233) < base.Ui64(v232)))
				v239 = base.B2i32(v226 == v236)
				if v226 == v236 {
					v240 = base.B2i32(base.Ui64(v233) <= base.Ui64(v223))
				} else {
					v240 = base.B2i32(v236 <= v226)
				}
				if v80 != 0 {
					v253 = v240
				} else {
					if v226 == v236 {
						v243 = base.B2i32(base.Ui64(v223) <= base.Ui64(v233))
					} else {
						v243 = base.B2i32(v226 <= v236)
					}
					v253 = v243
				}
				m.G0 = v16 + int32(48)
				return v253
			}
		}
	} else {
		F_errstart_cold(m, int32(21), int32(0))
		mBase = m.M
		v261 = m.ExcPending
		if v261 != 0 {
			return int32(0)
		} else {
			F_errcode(m, int32(50593922))
			mBase = m.M
			v264 = m.ExcPending
			if v264 != 0 {
				return int32(0)
			} else {
				F_errmsg(m, int32(_a_F_in_range_interval_interval_0), int32(0))
				mBase = m.M
				v268 = m.ExcPending
				if v268 != 0 {
					return int32(0)
				} else {
					F_errfinish(m, int32(_a_F_in_range_interval_interval_1), int32(3947), int32(_a_F_in_range_interval_interval_2))
					mBase = m.M
					v273 = m.ExcPending
					if v273 != 0 {
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
func F_show_in_hot_standby(m *base.Module) int32 {
	mBase := m.M
	_ = mBase
	var v3 int32
	_ = v3
	var v8 int32
	_ = v8
	var v9 int32
	_ = v9
	var v11 int32
	_ = v11
	var v15 int32
	_ = v15
	var v18 int32
	_ = v18
	v3 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_show_in_hot_standby[0])))
	if v3 == int32(1) {
		v8 = *(*int32)(unsafe.Add(mBase, _c_F_show_in_hot_standby[1]))
		v9 = *(*int32)(unsafe.Add(mBase, uint32(v8)+316))
		v11 = base.B2i32(v9 != int32(2))
		*(*uint8)(unsafe.Add(mBase, _c_F_show_in_hot_standby[0])) = uint8(v11)
		if v9 != int32(2) {
			v15 = int32(_a_F_show_in_hot_standby_0)
		} else {
			v15 = int32(_a_F_show_in_hot_standby_1)
		}
		v18 = v15
	} else {
		v18 = int32(_a_F_show_in_hot_standby_1)
	}
	return v18
}
