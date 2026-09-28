package p5

import (
	base "github.com/shibukawa/pgmem/internal/aot/pgaot/base"
	"math"
	"unsafe"
)

func F_in_range_float8_float8(m *base.Module, l0 int32) int64 {
	mBase := m.M
	_ = mBase
	var v9 int64
	_ = v9
	var v14 float64
	_ = v14
	var v17 int64
	_ = v17
	var v18 int64
	_ = v18
	var v19 int64
	_ = v19
	var v20 int64
	_ = v20
	var v21 int64
	_ = v21
	var v39 int64
	_ = v39
	var v41 float64
	_ = v41
	var v43 float64
	_ = v43
	var v62 float64
	_ = v62
	var v66 float64
	_ = v66
	var v67 float64
	_ = v67
	var v82 int32
	_ = v82
	var v85 int32
	_ = v85
	var v89 int32
	_ = v89
	var v94 int32
	_ = v94
	v9 = *(*int64)(unsafe.Add(mBase, uint32(l0)+56))
	if base.Ui64(int64(9218868437227405312)) < base.Ui64(v9&int64(9223372036854775807)) {
		F_errstart_cold(m, int32(21), int32(0))
		mBase = m.M
		v82 = m.ExcPending
		if v82 != 0 {
			return int64(0)
		} else {
			F_errcode(m, int32(50593922))
			mBase = m.M
			v85 = m.ExcPending
			if v85 != 0 {
				return int64(0)
			} else {
				F_errmsg(m, int32(_a_F_in_range_float8_float8_0), int32(0))
				mBase = m.M
				v89 = m.ExcPending
				if v89 != 0 {
					return int64(0)
				} else {
					F_errfinish(m, int32(_a_F_in_range_float8_float8_1), int32(1084), int32(_a_F_in_range_float8_float8_2))
					mBase = m.M
					v94 = m.ExcPending
					if v94 != 0 {
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
		v14 = base.F64_reinterpret_i64(v9)
		if base.F64_lt(v14, float64(0)) != 0 {
			F_errstart_cold(m, int32(21), int32(0))
			mBase = m.M
			v82 = m.ExcPending
			if v82 != 0 {
				return int64(0)
			} else {
				F_errcode(m, int32(50593922))
				mBase = m.M
				v85 = m.ExcPending
				if v85 != 0 {
					return int64(0)
				} else {
					F_errmsg(m, int32(_a_F_in_range_float8_float8_0), int32(0))
					mBase = m.M
					v89 = m.ExcPending
					if v89 != 0 {
						return int64(0)
					} else {
						F_errfinish(m, int32(_a_F_in_range_float8_float8_1), int32(1084), int32(_a_F_in_range_float8_float8_2))
						mBase = m.M
						v94 = m.ExcPending
						if v94 != 0 {
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
			v17 = *(*int64)(unsafe.Add(mBase, uint32(l0)+88))
			v18 = *(*int64)(unsafe.Add(mBase, uint32(l0)+40))
			v19 = int64(9223372036854775807)
			v20 = v18 & v19
			v21 = *(*int64)(unsafe.Add(mBase, uint32(l0)+24))
			if base.Ui64(int64(9218868437227405313)) <= base.Ui64(v21&v19) {
				return base.I64_extend_i32_u(base.B2i32(v17 == int64(0)) | base.B2i32(base.Ui64(int64(9218868437227405312)) < base.Ui64(v20)))
			} else {
				if base.Ui64(int64(9218868437227405312)) < base.Ui64(v20) {
					return base.I64_extend_i32_u(base.B2i32(v17 != int64(0)))
				} else {
					v39 = *(*int64)(unsafe.Add(mBase, uint32(l0)+72))
					v41 = math.Float64frombits(uint64(0x7ff0000000000000))
					v43 = base.F64_reinterpret_i64(v18)
					if base.F64_ne(base.F64_abs(v14), v41)|base.F64_ne(base.F64_abs(v43), v41) != 0 {
						v62 = base.F64_reinterpret_i64(v21)
						if v39 == int64(0) {
							v66 = v14
						} else {
							v66 = base.F64_neg(v14)
						}
						v67 = base.F64_add(v66, v43)
						if v17 != int64(0) {
							return base.I64_extend_i32_u(base.F64_ge(v67, v62))
						} else {
							return base.I64_extend_i32_u(base.F64_le(v67, v62))
						}
					} else {
						if v39 != int64(0) {
							if base.F64_gt(v43, float64(0)) == int32(0) {
								v62 = base.F64_reinterpret_i64(v21)
								if v39 == int64(0) {
									v66 = v14
								} else {
									v66 = base.F64_neg(v14)
								}
								v67 = base.F64_add(v66, v43)
								if v17 != int64(0) {
									return base.I64_extend_i32_u(base.F64_ge(v67, v62))
								} else {
									return base.I64_extend_i32_u(base.F64_le(v67, v62))
								}
							} else {
								return int64(1)
							}
						} else {
							if base.F64_lt(v43, float64(0)) == int32(0) {
								v62 = base.F64_reinterpret_i64(v21)
								if v39 == int64(0) {
									v66 = v14
								} else {
									v66 = base.F64_neg(v14)
								}
								v67 = base.F64_add(v66, v43)
								if v17 != int64(0) {
									return base.I64_extend_i32_u(base.F64_ge(v67, v62))
								} else {
									return base.I64_extend_i32_u(base.F64_le(v67, v62))
								}
							} else {
								return int64(1)
							}
						}
					}
				}
			}
		}
	}
}
func F_in_range_int2_int2(m *base.Module, l0 int32) int64 {
	mBase := m.M
	_ = mBase
	var v4 int64
	_ = v4
	var v5 int64
	_ = v5
	var v6 int64
	_ = v6
	var v7 int64
	_ = v7
	var v8 int64
	_ = v8
	var v9 int64
	_ = v9
	var v12 int32
	_ = v12
	v4 = *(*int64)(unsafe.Add(mBase, uint32(l0)+24))
	v5 = *(*int64)(unsafe.Add(mBase, uint32(l0)+40))
	v6 = int64(*(*int16)(unsafe.Add(mBase, uint32(l0)+56)))
	v7 = *(*int64)(unsafe.Add(mBase, uint32(l0)+72))
	v8 = *(*int64)(unsafe.Add(mBase, uint32(l0)+88))
	v9 = F_DirectFunctionCall5Coll(m, int32(1444), int32(0), v4, v5, v6, v7, v8)
	mBase = m.M
	v12 = m.ExcPending
	if v12 != 0 {
		return int64(0)
	} else {
		return v9
	}
}
func F_in_range_int4_int2(m *base.Module, l0 int32) int64 {
	mBase := m.M
	_ = mBase
	var v4 int64
	_ = v4
	var v5 int64
	_ = v5
	var v6 int64
	_ = v6
	var v7 int64
	_ = v7
	var v8 int64
	_ = v8
	var v9 int64
	_ = v9
	var v12 int32
	_ = v12
	v4 = *(*int64)(unsafe.Add(mBase, uint32(l0)+24))
	v5 = *(*int64)(unsafe.Add(mBase, uint32(l0)+40))
	v6 = int64(*(*int16)(unsafe.Add(mBase, uint32(l0)+56)))
	v7 = *(*int64)(unsafe.Add(mBase, uint32(l0)+72))
	v8 = *(*int64)(unsafe.Add(mBase, uint32(l0)+88))
	v9 = F_DirectFunctionCall5Coll(m, int32(1443), int32(0), v4, v5, v6, v7, v8)
	mBase = m.M
	v12 = m.ExcPending
	if v12 != 0 {
		return int64(0)
	} else {
		return v9
	}
}
func F_in_range_int4_int4(m *base.Module, l0 int32) int64 {
	mBase := m.M
	_ = mBase
	var v6 int32
	_ = v6
	var v9 int64
	_ = v9
	var v12 int64
	_ = v12
	var v14 int32
	_ = v14
	var v15 int32
	_ = v15
	var v18 int32
	_ = v18
	var v19 int32
	_ = v19
	var v27 int32
	_ = v27
	var v41 int32
	_ = v41
	var v44 int32
	_ = v44
	var v48 int32
	_ = v48
	var v53 int32
	_ = v53
	v6 = *(*int32)(unsafe.Add(mBase, uint32(l0)+56))
	if int32(0) <= v6 {
		v9 = *(*int64)(unsafe.Add(mBase, uint32(l0)+88))
		v12 = *(*int64)(unsafe.Add(mBase, uint32(l0)+72))
		v14 = base.B2i32(v12 != int64(0))
		if v12 != int64(0) {
			v15 = int32(0) - v6
		} else {
			v15 = v6
		}
		v18 = *(*int32)(unsafe.Add(mBase, uint32(l0)+40))
		v19 = v18 + v15
		if base.B2i32(v15 < int32(0)) != base.B2i32(v19 < v18) {
			return base.I64_extend_i32_u(v14 ^ base.B2i32(v9 != int64(0)))
		} else {
			v27 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
			if v9 != int64(0) {
				return base.I64_extend_i32_u(base.B2i32(v27 <= v19))
			} else {
				return base.I64_extend_i32_u(base.B2i32(v19 <= v27))
			}
		}
	} else {
		F_errstart_cold(m, int32(21), int32(0))
		mBase = m.M
		v41 = m.ExcPending
		if v41 != 0 {
			return int64(0)
		} else {
			F_errcode(m, int32(50593922))
			mBase = m.M
			v44 = m.ExcPending
			if v44 != 0 {
				return int64(0)
			} else {
				F_errmsg(m, int32(_a_F_in_range_int4_int4_0), int32(0))
				mBase = m.M
				v48 = m.ExcPending
				if v48 != 0 {
					return int64(0)
				} else {
					F_errfinish(m, int32(_a_F_in_range_int4_int4_1), int32(664), int32(_a_F_in_range_int4_int4_2))
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
		}
	}
}
func F_in_range_int4_int8(m *base.Module, l0 int32) int64 {
	mBase := m.M
	_ = mBase
	var v6 int64
	_ = v6
	var v9 int64
	_ = v9
	var v10 int64
	_ = v10
	var v12 int64
	_ = v12
	var v14 int32
	_ = v14
	var v15 int64
	_ = v15
	var v18 int64
	_ = v18
	var v19 int64
	_ = v19
	var v27 int64
	_ = v27
	var v41 int32
	_ = v41
	var v44 int32
	_ = v44
	var v48 int32
	_ = v48
	var v53 int32
	_ = v53
	v6 = *(*int64)(unsafe.Add(mBase, uint32(l0)+56))
	if int64(0) <= v6 {
		v9 = *(*int64)(unsafe.Add(mBase, uint32(l0)+88))
		v10 = int64(0)
		v12 = *(*int64)(unsafe.Add(mBase, uint32(l0)+72))
		v14 = base.B2i32(v12 != v10)
		if v12 != v10 {
			v15 = v10 - v6
		} else {
			v15 = v6
		}
		v18 = int64(*(*int32)(unsafe.Add(mBase, uint32(l0)+40)))
		v19 = v18 + v15
		if base.B2i32(v15 < int64(0)) != base.B2i32(v19 < v18) {
			return base.I64_extend_i32_u(v14 ^ base.B2i32(v9 != int64(0)))
		} else {
			v27 = int64(*(*int32)(unsafe.Add(mBase, uint32(l0)+24)))
			if v9 != int64(0) {
				return base.I64_extend_i32_u(base.B2i32(v27 <= v19))
			} else {
				return base.I64_extend_i32_u(base.B2i32(v19 <= v27))
			}
		}
	} else {
		F_errstart_cold(m, int32(21), int32(0))
		mBase = m.M
		v41 = m.ExcPending
		if v41 != 0 {
			return int64(0)
		} else {
			F_errcode(m, int32(50593922))
			mBase = m.M
			v44 = m.ExcPending
			if v44 != 0 {
				return int64(0)
			} else {
				F_errmsg(m, int32(_a_F_in_range_int4_int8_0), int32(0))
				mBase = m.M
				v48 = m.ExcPending
				if v48 != 0 {
					return int64(0)
				} else {
					F_errfinish(m, int32(_a_F_in_range_int4_int8_1), int32(711), int32(_a_F_in_range_int4_int8_2))
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
		}
	}
}
func F_in_range_interval_interval(m *base.Module, l0 int32) int64 {
	mBase := m.M
	_ = mBase
	var v11 int32
	_ = v11
	var v13 int32
	_ = v13
	var v16 int32
	_ = v16
	var v17 int64
	_ = v17
	var v18 int32
	_ = v18
	var v19 int32
	_ = v19
	var v23 int32
	_ = v23
	var v25 int64
	_ = v25
	var v34 int64
	_ = v34
	var v35 int64
	_ = v35
	var v37 int64
	_ = v37
	var v40 int64
	_ = v40
	var v41 int64
	_ = v41
	var v43 int64
	_ = v43
	var v44 int64
	_ = v44
	var v48 int64
	_ = v48
	var v55 int64
	_ = v55
	var v66 int64
	_ = v66
	var v67 int64
	_ = v67
	var v71 int64
	_ = v71
	var v78 int64
	_ = v78
	var v79 int64
	_ = v79
	var v80 int64
	_ = v80
	var v81 int32
	_ = v81
	var v82 int32
	_ = v82
	var v92 int32
	_ = v92
	var v93 int32
	_ = v93
	var v96 int32
	_ = v96
	var v99 int32
	_ = v99
	var v102 int64
	_ = v102
	var v108 int32
	_ = v108
	var v111 int64
	_ = v111
	var v123 int32
	_ = v123
	var v126 int64
	_ = v126
	var v130 int64
	_ = v130
	var v133 int32
	_ = v133
	var v134 int64
	_ = v134
	var v135 int64
	_ = v135
	var v138 int64
	_ = v138
	var v147 int64
	_ = v147
	var v148 int64
	_ = v148
	var v150 int64
	_ = v150
	var v153 int64
	_ = v153
	var v154 int64
	_ = v154
	var v156 int64
	_ = v156
	var v157 int64
	_ = v157
	var v161 int64
	_ = v161
	var v168 int64
	_ = v168
	var v180 int32
	_ = v180
	var v181 int32
	_ = v181
	var v182 int64
	_ = v182
	var v185 int64
	_ = v185
	var v186 int64
	_ = v186
	var v195 int64
	_ = v195
	var v196 int64
	_ = v196
	var v198 int64
	_ = v198
	var v201 int64
	_ = v201
	var v202 int64
	_ = v202
	var v204 int64
	_ = v204
	var v205 int64
	_ = v205
	var v209 int64
	_ = v209
	var v216 int64
	_ = v216
	var v227 int64
	_ = v227
	var v228 int64
	_ = v228
	var v229 int64
	_ = v229
	var v232 int64
	_ = v232
	var v233 int64
	_ = v233
	var v236 int64
	_ = v236
	var v237 int64
	_ = v237
	var v238 int64
	_ = v238
	var v242 int64
	_ = v242
	var v243 int64
	_ = v243
	var v246 int64
	_ = v246
	var v249 int32
	_ = v249
	var v250 int32
	_ = v250
	var v255 int32
	_ = v255
	var v263 int32
	_ = v263
	var v272 int32
	_ = v272
	var v275 int32
	_ = v275
	var v279 int32
	_ = v279
	var v284 int32
	_ = v284
	v11 = m.G0
	v13 = v11 - int32(48)
	m.G0 = v13
	v16 = v13 + int32(32)
	v17 = *(*int64)(unsafe.Add(mBase, uint32(l0)+56))
	v18 = base.I32_wrap_i64(v17)
	v19 = *(*int32)(unsafe.Add(mBase, uint32(v18)+12))
	v23 = *(*int32)(unsafe.Add(mBase, uint32(v18)+8))
	v25 = base.I64_extend_i32_s(v19)*int64(30) + base.I64_extend_i32_s(v23)
	v34 = int64(32)
	v35 = int64(20)
	v37 = int64(base.Ui64(v25) >> (uint(v34) % 64))
	v40 = int64(4294967295)
	v41 = int64(500654080)
	v43 = v25 & v40
	v44 = v41 * v43
	v48 = int64(base.Ui64(v44)>>(uint(v34)%64)) + v41*v37
	v55 = v43*v35 + v48&v40
	*(*int64)(unsafe.Add(mBase, uint32(v16)+8)) = v25*int64(0) + v25>>(uint(int64(63))%64)*int64(86400000000) + v35*v37 + int64(base.Ui64(v48)>>(uint(v34)%64)) + int64(base.Ui64(v55)>>(uint(v34)%64))
	*(*int64)(unsafe.Add(mBase, uint32(v16))) = v44&v40 | v55<<(uint(v34)%64)
	v66 = *(*int64)(unsafe.Add(mBase, uint32(v13)+40))
	v67 = *(*int64)(unsafe.Add(mBase, uint32(v18)))
	v71 = *(*int64)(unsafe.Add(mBase, uint32(v13)+32))
	if int64(0) <= v66+v67>>(uint(int64(63))%64)+base.I64_extend_i32_u(base.B2i32(base.Ui64(v71+v67) < base.Ui64(v71))) {
		v78 = *(*int64)(unsafe.Add(mBase, uint32(l0)+40))
		v79 = *(*int64)(unsafe.Add(mBase, uint32(l0)+88))
		v80 = *(*int64)(unsafe.Add(mBase, uint32(l0)+72))
		v81 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
		v82 = int32(2147483647)
		if base.B2i32(v19 != v82)|base.B2i32(v23 != v82)|base.B2i32(v67 != int64(9223372036854775807)) == int32(0) {
			v92 = base.I32_wrap_i64(v78)
			v93 = *(*int32)(unsafe.Add(mBase, uint32(v92)+12))
			if v80 != int64(0) {
				v96 = int32(1606)
				if v93 != int32(2147483647) {
					v123 = v96
					v126 = int64(4294967295)
					v130 = F_DirectFunctionCall2Coll(m, v123, int32(0), v78&v126, v17&v126)
					mBase = m.M
					v133 = m.ExcPending
					if v133 != 0 {
						return int64(0)
					} else {
						v134 = int64(*(*int32)(unsafe.Add(mBase, uint32(v81)+8)))
						v135 = int64(*(*int32)(unsafe.Add(mBase, uint32(v81)+12)))
						v138 = v134 + v135*int64(30)
						v147 = int64(32)
						v148 = int64(20)
						v150 = int64(base.Ui64(v138) >> (uint(v147) % 64))
						v153 = int64(4294967295)
						v154 = int64(500654080)
						v156 = v138 & v153
						v157 = v154 * v156
						v161 = int64(base.Ui64(v157)>>(uint(v147)%64)) + v154*v150
						v168 = v156*v148 + v161&v153
						*(*int64)(unsafe.Add(mBase, uint32(v13)+8)) = v138*int64(0) + v138>>(uint(int64(63))%64)*int64(86400000000) + v148*v150 + int64(base.Ui64(v161)>>(uint(v147)%64)) + int64(base.Ui64(v168)>>(uint(v147)%64))
						*(*int64)(unsafe.Add(mBase, uint32(v13))) = v157&v153 | v168<<(uint(v147)%64)
						v180 = v13 + int32(16)
						v181 = base.I32_wrap_i64(v130)
						v182 = int64(*(*int32)(unsafe.Add(mBase, uint32(v181)+12)))
						v185 = int64(*(*int32)(unsafe.Add(mBase, uint32(v181)+8)))
						v186 = v182*int64(30) + v185
						v195 = int64(32)
						v196 = int64(20)
						v198 = int64(base.Ui64(v186) >> (uint(v195) % 64))
						v201 = int64(4294967295)
						v202 = int64(500654080)
						v204 = v186 & v201
						v205 = v202 * v204
						v209 = int64(base.Ui64(v205)>>(uint(v195)%64)) + v202*v198
						v216 = v204*v196 + v209&v201
						*(*int64)(unsafe.Add(mBase, uint32(v180)+8)) = v186*int64(0) + v186>>(uint(int64(63))%64)*int64(86400000000) + v196*v198 + int64(base.Ui64(v209)>>(uint(v195)%64)) + int64(base.Ui64(v216)>>(uint(v195)%64))
						*(*int64)(unsafe.Add(mBase, uint32(v180))) = v205&v201 | v216<<(uint(v195)%64)
						v227 = *(*int64)(unsafe.Add(mBase, uint32(v13)+24))
						v228 = *(*int64)(unsafe.Add(mBase, uint32(v181)))
						v229 = int64(63)
						v232 = *(*int64)(unsafe.Add(mBase, uint32(v13)+16))
						v233 = v228 + v232
						v236 = v227 + v228>>(uint(v229)%64) + base.I64_extend_i32_u(base.B2i32(base.Ui64(v233) < base.Ui64(v232)))
						v237 = *(*int64)(unsafe.Add(mBase, uint32(v13)+8))
						v238 = *(*int64)(unsafe.Add(mBase, uint32(v81)))
						v242 = *(*int64)(unsafe.Add(mBase, uint32(v13)))
						v243 = v238 + v242
						v246 = v237 + v238>>(uint(v229)%64) + base.I64_extend_i32_u(base.B2i32(base.Ui64(v243) < base.Ui64(v242)))
						v249 = base.B2i32(v236 == v246)
						if v236 == v246 {
							v250 = base.B2i32(base.Ui64(v243) <= base.Ui64(v233))
						} else {
							v250 = base.B2i32(v246 <= v236)
						}
						if v79 != int64(0) {
							v263 = v250
						} else {
							if v236 == v246 {
								v255 = base.B2i32(base.Ui64(v233) <= base.Ui64(v243))
							} else {
								v255 = base.B2i32(v236 <= v246)
							}
							v263 = v255
						}
						m.G0 = v13 + int32(48)
						return base.I64_extend_i32_u(v263)
					}
				} else {
					v99 = *(*int32)(unsafe.Add(mBase, uint32(v92)+8))
					if v99 != int32(2147483647) {
						v123 = v96
						v126 = int64(4294967295)
						v130 = F_DirectFunctionCall2Coll(m, v123, int32(0), v78&v126, v17&v126)
						mBase = m.M
						v133 = m.ExcPending
						if v133 != 0 {
							return int64(0)
						} else {
							v134 = int64(*(*int32)(unsafe.Add(mBase, uint32(v81)+8)))
							v135 = int64(*(*int32)(unsafe.Add(mBase, uint32(v81)+12)))
							v138 = v134 + v135*int64(30)
							v147 = int64(32)
							v148 = int64(20)
							v150 = int64(base.Ui64(v138) >> (uint(v147) % 64))
							v153 = int64(4294967295)
							v154 = int64(500654080)
							v156 = v138 & v153
							v157 = v154 * v156
							v161 = int64(base.Ui64(v157)>>(uint(v147)%64)) + v154*v150
							v168 = v156*v148 + v161&v153
							*(*int64)(unsafe.Add(mBase, uint32(v13)+8)) = v138*int64(0) + v138>>(uint(int64(63))%64)*int64(86400000000) + v148*v150 + int64(base.Ui64(v161)>>(uint(v147)%64)) + int64(base.Ui64(v168)>>(uint(v147)%64))
							*(*int64)(unsafe.Add(mBase, uint32(v13))) = v157&v153 | v168<<(uint(v147)%64)
							v180 = v13 + int32(16)
							v181 = base.I32_wrap_i64(v130)
							v182 = int64(*(*int32)(unsafe.Add(mBase, uint32(v181)+12)))
							v185 = int64(*(*int32)(unsafe.Add(mBase, uint32(v181)+8)))
							v186 = v182*int64(30) + v185
							v195 = int64(32)
							v196 = int64(20)
							v198 = int64(base.Ui64(v186) >> (uint(v195) % 64))
							v201 = int64(4294967295)
							v202 = int64(500654080)
							v204 = v186 & v201
							v205 = v202 * v204
							v209 = int64(base.Ui64(v205)>>(uint(v195)%64)) + v202*v198
							v216 = v204*v196 + v209&v201
							*(*int64)(unsafe.Add(mBase, uint32(v180)+8)) = v186*int64(0) + v186>>(uint(int64(63))%64)*int64(86400000000) + v196*v198 + int64(base.Ui64(v209)>>(uint(v195)%64)) + int64(base.Ui64(v216)>>(uint(v195)%64))
							*(*int64)(unsafe.Add(mBase, uint32(v180))) = v205&v201 | v216<<(uint(v195)%64)
							v227 = *(*int64)(unsafe.Add(mBase, uint32(v13)+24))
							v228 = *(*int64)(unsafe.Add(mBase, uint32(v181)))
							v229 = int64(63)
							v232 = *(*int64)(unsafe.Add(mBase, uint32(v13)+16))
							v233 = v228 + v232
							v236 = v227 + v228>>(uint(v229)%64) + base.I64_extend_i32_u(base.B2i32(base.Ui64(v233) < base.Ui64(v232)))
							v237 = *(*int64)(unsafe.Add(mBase, uint32(v13)+8))
							v238 = *(*int64)(unsafe.Add(mBase, uint32(v81)))
							v242 = *(*int64)(unsafe.Add(mBase, uint32(v13)))
							v243 = v238 + v242
							v246 = v237 + v238>>(uint(v229)%64) + base.I64_extend_i32_u(base.B2i32(base.Ui64(v243) < base.Ui64(v242)))
							v249 = base.B2i32(v236 == v246)
							if v236 == v246 {
								v250 = base.B2i32(base.Ui64(v243) <= base.Ui64(v233))
							} else {
								v250 = base.B2i32(v246 <= v236)
							}
							if v79 != int64(0) {
								v263 = v250
							} else {
								if v236 == v246 {
									v255 = base.B2i32(base.Ui64(v233) <= base.Ui64(v243))
								} else {
									v255 = base.B2i32(v236 <= v246)
								}
								v263 = v255
							}
							m.G0 = v13 + int32(48)
							return base.I64_extend_i32_u(v263)
						}
					} else {
						v102 = *(*int64)(unsafe.Add(mBase, uint32(v92)))
						if v102 != int64(9223372036854775807) {
							v123 = v96
							v126 = int64(4294967295)
							v130 = F_DirectFunctionCall2Coll(m, v123, int32(0), v78&v126, v17&v126)
							mBase = m.M
							v133 = m.ExcPending
							if v133 != 0 {
								return int64(0)
							} else {
								v134 = int64(*(*int32)(unsafe.Add(mBase, uint32(v81)+8)))
								v135 = int64(*(*int32)(unsafe.Add(mBase, uint32(v81)+12)))
								v138 = v134 + v135*int64(30)
								v147 = int64(32)
								v148 = int64(20)
								v150 = int64(base.Ui64(v138) >> (uint(v147) % 64))
								v153 = int64(4294967295)
								v154 = int64(500654080)
								v156 = v138 & v153
								v157 = v154 * v156
								v161 = int64(base.Ui64(v157)>>(uint(v147)%64)) + v154*v150
								v168 = v156*v148 + v161&v153
								*(*int64)(unsafe.Add(mBase, uint32(v13)+8)) = v138*int64(0) + v138>>(uint(int64(63))%64)*int64(86400000000) + v148*v150 + int64(base.Ui64(v161)>>(uint(v147)%64)) + int64(base.Ui64(v168)>>(uint(v147)%64))
								*(*int64)(unsafe.Add(mBase, uint32(v13))) = v157&v153 | v168<<(uint(v147)%64)
								v180 = v13 + int32(16)
								v181 = base.I32_wrap_i64(v130)
								v182 = int64(*(*int32)(unsafe.Add(mBase, uint32(v181)+12)))
								v185 = int64(*(*int32)(unsafe.Add(mBase, uint32(v181)+8)))
								v186 = v182*int64(30) + v185
								v195 = int64(32)
								v196 = int64(20)
								v198 = int64(base.Ui64(v186) >> (uint(v195) % 64))
								v201 = int64(4294967295)
								v202 = int64(500654080)
								v204 = v186 & v201
								v205 = v202 * v204
								v209 = int64(base.Ui64(v205)>>(uint(v195)%64)) + v202*v198
								v216 = v204*v196 + v209&v201
								*(*int64)(unsafe.Add(mBase, uint32(v180)+8)) = v186*int64(0) + v186>>(uint(int64(63))%64)*int64(86400000000) + v196*v198 + int64(base.Ui64(v209)>>(uint(v195)%64)) + int64(base.Ui64(v216)>>(uint(v195)%64))
								*(*int64)(unsafe.Add(mBase, uint32(v180))) = v205&v201 | v216<<(uint(v195)%64)
								v227 = *(*int64)(unsafe.Add(mBase, uint32(v13)+24))
								v228 = *(*int64)(unsafe.Add(mBase, uint32(v181)))
								v229 = int64(63)
								v232 = *(*int64)(unsafe.Add(mBase, uint32(v13)+16))
								v233 = v228 + v232
								v236 = v227 + v228>>(uint(v229)%64) + base.I64_extend_i32_u(base.B2i32(base.Ui64(v233) < base.Ui64(v232)))
								v237 = *(*int64)(unsafe.Add(mBase, uint32(v13)+8))
								v238 = *(*int64)(unsafe.Add(mBase, uint32(v81)))
								v242 = *(*int64)(unsafe.Add(mBase, uint32(v13)))
								v243 = v238 + v242
								v246 = v237 + v238>>(uint(v229)%64) + base.I64_extend_i32_u(base.B2i32(base.Ui64(v243) < base.Ui64(v242)))
								v249 = base.B2i32(v236 == v246)
								if v236 == v246 {
									v250 = base.B2i32(base.Ui64(v243) <= base.Ui64(v233))
								} else {
									v250 = base.B2i32(v246 <= v236)
								}
								if v79 != int64(0) {
									v263 = v250
								} else {
									if v236 == v246 {
										v255 = base.B2i32(base.Ui64(v233) <= base.Ui64(v243))
									} else {
										v255 = base.B2i32(v236 <= v246)
									}
									v263 = v255
								}
								m.G0 = v13 + int32(48)
								return base.I64_extend_i32_u(v263)
							}
						} else {
							v263 = int32(1)
							m.G0 = v13 + int32(48)
							return base.I64_extend_i32_u(v263)
						}
					}
				}
			} else {
				if v93 != int32(-2147483648) {
					v123 = int32(1604)
					v126 = int64(4294967295)
					v130 = F_DirectFunctionCall2Coll(m, v123, int32(0), v78&v126, v17&v126)
					mBase = m.M
					v133 = m.ExcPending
					if v133 != 0 {
						return int64(0)
					} else {
						v134 = int64(*(*int32)(unsafe.Add(mBase, uint32(v81)+8)))
						v135 = int64(*(*int32)(unsafe.Add(mBase, uint32(v81)+12)))
						v138 = v134 + v135*int64(30)
						v147 = int64(32)
						v148 = int64(20)
						v150 = int64(base.Ui64(v138) >> (uint(v147) % 64))
						v153 = int64(4294967295)
						v154 = int64(500654080)
						v156 = v138 & v153
						v157 = v154 * v156
						v161 = int64(base.Ui64(v157)>>(uint(v147)%64)) + v154*v150
						v168 = v156*v148 + v161&v153
						*(*int64)(unsafe.Add(mBase, uint32(v13)+8)) = v138*int64(0) + v138>>(uint(int64(63))%64)*int64(86400000000) + v148*v150 + int64(base.Ui64(v161)>>(uint(v147)%64)) + int64(base.Ui64(v168)>>(uint(v147)%64))
						*(*int64)(unsafe.Add(mBase, uint32(v13))) = v157&v153 | v168<<(uint(v147)%64)
						v180 = v13 + int32(16)
						v181 = base.I32_wrap_i64(v130)
						v182 = int64(*(*int32)(unsafe.Add(mBase, uint32(v181)+12)))
						v185 = int64(*(*int32)(unsafe.Add(mBase, uint32(v181)+8)))
						v186 = v182*int64(30) + v185
						v195 = int64(32)
						v196 = int64(20)
						v198 = int64(base.Ui64(v186) >> (uint(v195) % 64))
						v201 = int64(4294967295)
						v202 = int64(500654080)
						v204 = v186 & v201
						v205 = v202 * v204
						v209 = int64(base.Ui64(v205)>>(uint(v195)%64)) + v202*v198
						v216 = v204*v196 + v209&v201
						*(*int64)(unsafe.Add(mBase, uint32(v180)+8)) = v186*int64(0) + v186>>(uint(int64(63))%64)*int64(86400000000) + v196*v198 + int64(base.Ui64(v209)>>(uint(v195)%64)) + int64(base.Ui64(v216)>>(uint(v195)%64))
						*(*int64)(unsafe.Add(mBase, uint32(v180))) = v205&v201 | v216<<(uint(v195)%64)
						v227 = *(*int64)(unsafe.Add(mBase, uint32(v13)+24))
						v228 = *(*int64)(unsafe.Add(mBase, uint32(v181)))
						v229 = int64(63)
						v232 = *(*int64)(unsafe.Add(mBase, uint32(v13)+16))
						v233 = v228 + v232
						v236 = v227 + v228>>(uint(v229)%64) + base.I64_extend_i32_u(base.B2i32(base.Ui64(v233) < base.Ui64(v232)))
						v237 = *(*int64)(unsafe.Add(mBase, uint32(v13)+8))
						v238 = *(*int64)(unsafe.Add(mBase, uint32(v81)))
						v242 = *(*int64)(unsafe.Add(mBase, uint32(v13)))
						v243 = v238 + v242
						v246 = v237 + v238>>(uint(v229)%64) + base.I64_extend_i32_u(base.B2i32(base.Ui64(v243) < base.Ui64(v242)))
						v249 = base.B2i32(v236 == v246)
						if v236 == v246 {
							v250 = base.B2i32(base.Ui64(v243) <= base.Ui64(v233))
						} else {
							v250 = base.B2i32(v246 <= v236)
						}
						if v79 != int64(0) {
							v263 = v250
						} else {
							if v236 == v246 {
								v255 = base.B2i32(base.Ui64(v233) <= base.Ui64(v243))
							} else {
								v255 = base.B2i32(v236 <= v246)
							}
							v263 = v255
						}
						m.G0 = v13 + int32(48)
						return base.I64_extend_i32_u(v263)
					}
				} else {
					v108 = *(*int32)(unsafe.Add(mBase, uint32(v92)+8))
					if v108 != int32(-2147483648) {
						v123 = int32(1604)
						v126 = int64(4294967295)
						v130 = F_DirectFunctionCall2Coll(m, v123, int32(0), v78&v126, v17&v126)
						mBase = m.M
						v133 = m.ExcPending
						if v133 != 0 {
							return int64(0)
						} else {
							v134 = int64(*(*int32)(unsafe.Add(mBase, uint32(v81)+8)))
							v135 = int64(*(*int32)(unsafe.Add(mBase, uint32(v81)+12)))
							v138 = v134 + v135*int64(30)
							v147 = int64(32)
							v148 = int64(20)
							v150 = int64(base.Ui64(v138) >> (uint(v147) % 64))
							v153 = int64(4294967295)
							v154 = int64(500654080)
							v156 = v138 & v153
							v157 = v154 * v156
							v161 = int64(base.Ui64(v157)>>(uint(v147)%64)) + v154*v150
							v168 = v156*v148 + v161&v153
							*(*int64)(unsafe.Add(mBase, uint32(v13)+8)) = v138*int64(0) + v138>>(uint(int64(63))%64)*int64(86400000000) + v148*v150 + int64(base.Ui64(v161)>>(uint(v147)%64)) + int64(base.Ui64(v168)>>(uint(v147)%64))
							*(*int64)(unsafe.Add(mBase, uint32(v13))) = v157&v153 | v168<<(uint(v147)%64)
							v180 = v13 + int32(16)
							v181 = base.I32_wrap_i64(v130)
							v182 = int64(*(*int32)(unsafe.Add(mBase, uint32(v181)+12)))
							v185 = int64(*(*int32)(unsafe.Add(mBase, uint32(v181)+8)))
							v186 = v182*int64(30) + v185
							v195 = int64(32)
							v196 = int64(20)
							v198 = int64(base.Ui64(v186) >> (uint(v195) % 64))
							v201 = int64(4294967295)
							v202 = int64(500654080)
							v204 = v186 & v201
							v205 = v202 * v204
							v209 = int64(base.Ui64(v205)>>(uint(v195)%64)) + v202*v198
							v216 = v204*v196 + v209&v201
							*(*int64)(unsafe.Add(mBase, uint32(v180)+8)) = v186*int64(0) + v186>>(uint(int64(63))%64)*int64(86400000000) + v196*v198 + int64(base.Ui64(v209)>>(uint(v195)%64)) + int64(base.Ui64(v216)>>(uint(v195)%64))
							*(*int64)(unsafe.Add(mBase, uint32(v180))) = v205&v201 | v216<<(uint(v195)%64)
							v227 = *(*int64)(unsafe.Add(mBase, uint32(v13)+24))
							v228 = *(*int64)(unsafe.Add(mBase, uint32(v181)))
							v229 = int64(63)
							v232 = *(*int64)(unsafe.Add(mBase, uint32(v13)+16))
							v233 = v228 + v232
							v236 = v227 + v228>>(uint(v229)%64) + base.I64_extend_i32_u(base.B2i32(base.Ui64(v233) < base.Ui64(v232)))
							v237 = *(*int64)(unsafe.Add(mBase, uint32(v13)+8))
							v238 = *(*int64)(unsafe.Add(mBase, uint32(v81)))
							v242 = *(*int64)(unsafe.Add(mBase, uint32(v13)))
							v243 = v238 + v242
							v246 = v237 + v238>>(uint(v229)%64) + base.I64_extend_i32_u(base.B2i32(base.Ui64(v243) < base.Ui64(v242)))
							v249 = base.B2i32(v236 == v246)
							if v236 == v246 {
								v250 = base.B2i32(base.Ui64(v243) <= base.Ui64(v233))
							} else {
								v250 = base.B2i32(v246 <= v236)
							}
							if v79 != int64(0) {
								v263 = v250
							} else {
								if v236 == v246 {
									v255 = base.B2i32(base.Ui64(v233) <= base.Ui64(v243))
								} else {
									v255 = base.B2i32(v236 <= v246)
								}
								v263 = v255
							}
							m.G0 = v13 + int32(48)
							return base.I64_extend_i32_u(v263)
						}
					} else {
						v111 = *(*int64)(unsafe.Add(mBase, uint32(v92)))
						if v111 != int64(-9223372036854775807-1) {
							v123 = int32(1604)
							v126 = int64(4294967295)
							v130 = F_DirectFunctionCall2Coll(m, v123, int32(0), v78&v126, v17&v126)
							mBase = m.M
							v133 = m.ExcPending
							if v133 != 0 {
								return int64(0)
							} else {
								v134 = int64(*(*int32)(unsafe.Add(mBase, uint32(v81)+8)))
								v135 = int64(*(*int32)(unsafe.Add(mBase, uint32(v81)+12)))
								v138 = v134 + v135*int64(30)
								v147 = int64(32)
								v148 = int64(20)
								v150 = int64(base.Ui64(v138) >> (uint(v147) % 64))
								v153 = int64(4294967295)
								v154 = int64(500654080)
								v156 = v138 & v153
								v157 = v154 * v156
								v161 = int64(base.Ui64(v157)>>(uint(v147)%64)) + v154*v150
								v168 = v156*v148 + v161&v153
								*(*int64)(unsafe.Add(mBase, uint32(v13)+8)) = v138*int64(0) + v138>>(uint(int64(63))%64)*int64(86400000000) + v148*v150 + int64(base.Ui64(v161)>>(uint(v147)%64)) + int64(base.Ui64(v168)>>(uint(v147)%64))
								*(*int64)(unsafe.Add(mBase, uint32(v13))) = v157&v153 | v168<<(uint(v147)%64)
								v180 = v13 + int32(16)
								v181 = base.I32_wrap_i64(v130)
								v182 = int64(*(*int32)(unsafe.Add(mBase, uint32(v181)+12)))
								v185 = int64(*(*int32)(unsafe.Add(mBase, uint32(v181)+8)))
								v186 = v182*int64(30) + v185
								v195 = int64(32)
								v196 = int64(20)
								v198 = int64(base.Ui64(v186) >> (uint(v195) % 64))
								v201 = int64(4294967295)
								v202 = int64(500654080)
								v204 = v186 & v201
								v205 = v202 * v204
								v209 = int64(base.Ui64(v205)>>(uint(v195)%64)) + v202*v198
								v216 = v204*v196 + v209&v201
								*(*int64)(unsafe.Add(mBase, uint32(v180)+8)) = v186*int64(0) + v186>>(uint(int64(63))%64)*int64(86400000000) + v196*v198 + int64(base.Ui64(v209)>>(uint(v195)%64)) + int64(base.Ui64(v216)>>(uint(v195)%64))
								*(*int64)(unsafe.Add(mBase, uint32(v180))) = v205&v201 | v216<<(uint(v195)%64)
								v227 = *(*int64)(unsafe.Add(mBase, uint32(v13)+24))
								v228 = *(*int64)(unsafe.Add(mBase, uint32(v181)))
								v229 = int64(63)
								v232 = *(*int64)(unsafe.Add(mBase, uint32(v13)+16))
								v233 = v228 + v232
								v236 = v227 + v228>>(uint(v229)%64) + base.I64_extend_i32_u(base.B2i32(base.Ui64(v233) < base.Ui64(v232)))
								v237 = *(*int64)(unsafe.Add(mBase, uint32(v13)+8))
								v238 = *(*int64)(unsafe.Add(mBase, uint32(v81)))
								v242 = *(*int64)(unsafe.Add(mBase, uint32(v13)))
								v243 = v238 + v242
								v246 = v237 + v238>>(uint(v229)%64) + base.I64_extend_i32_u(base.B2i32(base.Ui64(v243) < base.Ui64(v242)))
								v249 = base.B2i32(v236 == v246)
								if v236 == v246 {
									v250 = base.B2i32(base.Ui64(v243) <= base.Ui64(v233))
								} else {
									v250 = base.B2i32(v246 <= v236)
								}
								if v79 != int64(0) {
									v263 = v250
								} else {
									if v236 == v246 {
										v255 = base.B2i32(base.Ui64(v233) <= base.Ui64(v243))
									} else {
										v255 = base.B2i32(v236 <= v246)
									}
									v263 = v255
								}
								m.G0 = v13 + int32(48)
								return base.I64_extend_i32_u(v263)
							}
						} else {
							v263 = int32(1)
							m.G0 = v13 + int32(48)
							return base.I64_extend_i32_u(v263)
						}
					}
				}
			}
		} else {
			if v80 != int64(0) {
				v123 = int32(1606)
			} else {
				v123 = int32(1604)
			}
			v126 = int64(4294967295)
			v130 = F_DirectFunctionCall2Coll(m, v123, int32(0), v78&v126, v17&v126)
			mBase = m.M
			v133 = m.ExcPending
			if v133 != 0 {
				return int64(0)
			} else {
				v134 = int64(*(*int32)(unsafe.Add(mBase, uint32(v81)+8)))
				v135 = int64(*(*int32)(unsafe.Add(mBase, uint32(v81)+12)))
				v138 = v134 + v135*int64(30)
				v147 = int64(32)
				v148 = int64(20)
				v150 = int64(base.Ui64(v138) >> (uint(v147) % 64))
				v153 = int64(4294967295)
				v154 = int64(500654080)
				v156 = v138 & v153
				v157 = v154 * v156
				v161 = int64(base.Ui64(v157)>>(uint(v147)%64)) + v154*v150
				v168 = v156*v148 + v161&v153
				*(*int64)(unsafe.Add(mBase, uint32(v13)+8)) = v138*int64(0) + v138>>(uint(int64(63))%64)*int64(86400000000) + v148*v150 + int64(base.Ui64(v161)>>(uint(v147)%64)) + int64(base.Ui64(v168)>>(uint(v147)%64))
				*(*int64)(unsafe.Add(mBase, uint32(v13))) = v157&v153 | v168<<(uint(v147)%64)
				v180 = v13 + int32(16)
				v181 = base.I32_wrap_i64(v130)
				v182 = int64(*(*int32)(unsafe.Add(mBase, uint32(v181)+12)))
				v185 = int64(*(*int32)(unsafe.Add(mBase, uint32(v181)+8)))
				v186 = v182*int64(30) + v185
				v195 = int64(32)
				v196 = int64(20)
				v198 = int64(base.Ui64(v186) >> (uint(v195) % 64))
				v201 = int64(4294967295)
				v202 = int64(500654080)
				v204 = v186 & v201
				v205 = v202 * v204
				v209 = int64(base.Ui64(v205)>>(uint(v195)%64)) + v202*v198
				v216 = v204*v196 + v209&v201
				*(*int64)(unsafe.Add(mBase, uint32(v180)+8)) = v186*int64(0) + v186>>(uint(int64(63))%64)*int64(86400000000) + v196*v198 + int64(base.Ui64(v209)>>(uint(v195)%64)) + int64(base.Ui64(v216)>>(uint(v195)%64))
				*(*int64)(unsafe.Add(mBase, uint32(v180))) = v205&v201 | v216<<(uint(v195)%64)
				v227 = *(*int64)(unsafe.Add(mBase, uint32(v13)+24))
				v228 = *(*int64)(unsafe.Add(mBase, uint32(v181)))
				v229 = int64(63)
				v232 = *(*int64)(unsafe.Add(mBase, uint32(v13)+16))
				v233 = v228 + v232
				v236 = v227 + v228>>(uint(v229)%64) + base.I64_extend_i32_u(base.B2i32(base.Ui64(v233) < base.Ui64(v232)))
				v237 = *(*int64)(unsafe.Add(mBase, uint32(v13)+8))
				v238 = *(*int64)(unsafe.Add(mBase, uint32(v81)))
				v242 = *(*int64)(unsafe.Add(mBase, uint32(v13)))
				v243 = v238 + v242
				v246 = v237 + v238>>(uint(v229)%64) + base.I64_extend_i32_u(base.B2i32(base.Ui64(v243) < base.Ui64(v242)))
				v249 = base.B2i32(v236 == v246)
				if v236 == v246 {
					v250 = base.B2i32(base.Ui64(v243) <= base.Ui64(v233))
				} else {
					v250 = base.B2i32(v246 <= v236)
				}
				if v79 != int64(0) {
					v263 = v250
				} else {
					if v236 == v246 {
						v255 = base.B2i32(base.Ui64(v233) <= base.Ui64(v243))
					} else {
						v255 = base.B2i32(v236 <= v246)
					}
					v263 = v255
				}
				m.G0 = v13 + int32(48)
				return base.I64_extend_i32_u(v263)
			}
		}
	} else {
		F_errstart_cold(m, int32(21), int32(0))
		mBase = m.M
		v272 = m.ExcPending
		if v272 != 0 {
			return int64(0)
		} else {
			F_errcode(m, int32(50593922))
			mBase = m.M
			v275 = m.ExcPending
			if v275 != 0 {
				return int64(0)
			} else {
				F_errmsg(m, int32(_a_F_in_range_interval_interval_0), int32(0))
				mBase = m.M
				v279 = m.ExcPending
				if v279 != 0 {
					return int64(0)
				} else {
					F_errfinish(m, int32(_a_F_in_range_interval_interval_1), int32(3948), int32(_a_F_in_range_interval_interval_2))
					mBase = m.M
					v284 = m.ExcPending
					if v284 != 0 {
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
		v9 = *(*int32)(unsafe.Add(mBase, uint32(v8)+308))
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
