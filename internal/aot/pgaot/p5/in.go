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
	var v16 int32
	_ = v16
	var v17 int32
	_ = v17
	var v18 float64
	_ = v18
	var v20 int64
	_ = v20
	var v21 int64
	_ = v21
	var v22 int32
	_ = v22
	var v23 float64
	_ = v23
	var v40 int32
	_ = v40
	var v60 float64
	_ = v60
	var v61 float64
	_ = v61
	var v71 int32
	_ = v71
	var v74 int32
	_ = v74
	var v78 int32
	_ = v78
	var v83 int32
	_ = v83
	v7 = *(*int32)(unsafe.Add(mBase, uint32(l0)+36))
	v8 = *(*float64)(unsafe.Add(mBase, uint32(v7)))
	if base.Ui64(int64(9218868437227405312)) < base.Ui64(base.I64_reinterpret_f64(v8)&int64(9223372036854775807)) {
		F_errstart_cold(m, int32(21), int32(0))
		mBase = m.M
		v71 = m.ExcPending
		if v71 != 0 {
			return int32(0)
		} else {
			F_errcode(m, int32(50593922))
			mBase = m.M
			v74 = m.ExcPending
			if v74 != 0 {
				return int32(0)
			} else {
				F_errmsg(m, int32(239981), int32(0))
				mBase = m.M
				v78 = m.ExcPending
				if v78 != 0 {
					return int32(0)
				} else {
					F_errfinish(m, int32(472060), int32(1043), int32(525905))
					mBase = m.M
					v83 = m.ExcPending
					if v83 != 0 {
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
		if base.F64_lt(v8, float64(0)) != 0 {
			F_errstart_cold(m, int32(21), int32(0))
			mBase = m.M
			v71 = m.ExcPending
			if v71 != 0 {
				return int32(0)
			} else {
				F_errcode(m, int32(50593922))
				mBase = m.M
				v74 = m.ExcPending
				if v74 != 0 {
					return int32(0)
				} else {
					F_errmsg(m, int32(239981), int32(0))
					mBase = m.M
					v78 = m.ExcPending
					if v78 != 0 {
						return int32(0)
					} else {
						F_errfinish(m, int32(472060), int32(1043), int32(525905))
						mBase = m.M
						v83 = m.ExcPending
						if v83 != 0 {
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
			v16 = *(*int32)(unsafe.Add(mBase, uint32(l0)+52))
			v17 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
			v18 = *(*float64)(unsafe.Add(mBase, uint32(v17)))
			v20 = int64(9223372036854775807)
			v21 = base.I64_reinterpret_f64(v18) & v20
			v22 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
			v23 = *(*float64)(unsafe.Add(mBase, uint32(v22)))
			if base.Ui64(int64(9218868437227405313)) <= base.Ui64(base.I64_reinterpret_f64(v23)&v20) {
				return base.B2i32(v16 == int32(0)) | base.B2i32(base.Ui64(int64(9218868437227405312)) < base.Ui64(v21))
			} else {
				if base.Ui64(int64(9218868437227405312)) < base.Ui64(v21) {
					return base.B2i32(v16 != int32(0))
				} else {
					v40 = *(*int32)(unsafe.Add(mBase, uint32(l0)+44))
					if base.F64_ne(base.F64_abs(v8), math.Float64frombits(uint64(0x7ff0000000000000))) != 0 {
						if v40 != 0 {
							v60 = base.F64_neg(v8)
						} else {
							v60 = v8
						}
						v61 = base.F64_add(v18, v60)
						if v16 != 0 {
							return base.F64_ge(v61, v23)
						} else {
							return base.F64_le(v61, v23)
						}
					} else {
						if base.F64_ne(base.F64_abs(v18), math.Float64frombits(uint64(0x7ff0000000000000))) != 0 {
							if v40 != 0 {
								v60 = base.F64_neg(v8)
							} else {
								v60 = v8
							}
							v61 = base.F64_add(v18, v60)
							if v16 != 0 {
								return base.F64_ge(v61, v23)
							} else {
								return base.F64_le(v61, v23)
							}
						} else {
							if v40 != 0 {
								if base.F64_gt(v18, float64(0)) == int32(0) {
									if v40 != 0 {
										v60 = base.F64_neg(v8)
									} else {
										v60 = v8
									}
									v61 = base.F64_add(v18, v60)
									if v16 != 0 {
										return base.F64_ge(v61, v23)
									} else {
										return base.F64_le(v61, v23)
									}
								} else {
									return int32(1)
								}
							} else {
								if base.F64_lt(v18, float64(0)) == int32(0) {
									if v40 != 0 {
										v60 = base.F64_neg(v8)
									} else {
										v60 = v8
									}
									v61 = base.F64_add(v18, v60)
									if v16 != 0 {
										return base.F64_ge(v61, v23)
									} else {
										return base.F64_le(v61, v23)
									}
								} else {
									return int32(1)
								}
							}
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
	v9 = F_DirectFunctionCall5Coll(m, int32(1318), int32(0), v4, v5, v6, v7, v8)
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
	v9 = F_DirectFunctionCall5Coll(m, int32(1317), int32(0), v4, v5, v6, v7, v8)
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
				F_errmsg(m, int32(239981), int32(0))
				mBase = m.M
				v43 = m.ExcPending
				if v43 != 0 {
					return int32(0)
				} else {
					F_errfinish(m, int32(471646), int32(664), int32(528351))
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
				F_errmsg(m, int32(239981), int32(0))
				mBase = m.M
				v44 = m.ExcPending
				if v44 != 0 {
					return int32(0)
				} else {
					F_errfinish(m, int32(471646), int32(711), int32(525714))
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
	var v13 int32
	_ = v13
	var v15 int32
	_ = v15
	var v18 int32
	_ = v18
	var v19 int32
	_ = v19
	var v20 int32
	_ = v20
	var v24 int32
	_ = v24
	var v26 int64
	_ = v26
	var v35 int64
	_ = v35
	var v36 int64
	_ = v36
	var v38 int64
	_ = v38
	var v41 int64
	_ = v41
	var v42 int64
	_ = v42
	var v44 int64
	_ = v44
	var v45 int64
	_ = v45
	var v49 int64
	_ = v49
	var v56 int64
	_ = v56
	var v67 int64
	_ = v67
	var v68 int64
	_ = v68
	var v72 int64
	_ = v72
	var v79 int32
	_ = v79
	var v80 int32
	_ = v80
	var v81 int32
	_ = v81
	var v82 int32
	_ = v82
	var v89 int32
	_ = v89
	var v90 int32
	_ = v90
	var v93 int32
	_ = v93
	var v96 int64
	_ = v96
	var v102 int32
	_ = v102
	var v105 int64
	_ = v105
	var v113 int32
	_ = v113
	var v116 int32
	_ = v116
	var v119 int32
	_ = v119
	var v120 int64
	_ = v120
	var v121 int64
	_ = v121
	var v124 int64
	_ = v124
	var v133 int64
	_ = v133
	var v134 int64
	_ = v134
	var v136 int64
	_ = v136
	var v139 int64
	_ = v139
	var v140 int64
	_ = v140
	var v142 int64
	_ = v142
	var v143 int64
	_ = v143
	var v147 int64
	_ = v147
	var v154 int64
	_ = v154
	var v166 int32
	_ = v166
	var v167 int64
	_ = v167
	var v168 int64
	_ = v168
	var v171 int64
	_ = v171
	var v180 int64
	_ = v180
	var v181 int64
	_ = v181
	var v183 int64
	_ = v183
	var v186 int64
	_ = v186
	var v187 int64
	_ = v187
	var v189 int64
	_ = v189
	var v190 int64
	_ = v190
	var v194 int64
	_ = v194
	var v201 int64
	_ = v201
	var v212 int64
	_ = v212
	var v213 int64
	_ = v213
	var v214 int64
	_ = v214
	var v217 int64
	_ = v217
	var v218 int64
	_ = v218
	var v221 int64
	_ = v221
	var v222 int64
	_ = v222
	var v223 int64
	_ = v223
	var v227 int64
	_ = v227
	var v228 int64
	_ = v228
	var v231 int64
	_ = v231
	var v235 int32
	_ = v235
	var v239 int32
	_ = v239
	var v247 int32
	_ = v247
	var v255 int32
	_ = v255
	var v258 int32
	_ = v258
	var v262 int32
	_ = v262
	var v267 int32
	_ = v267
	v13 = m.G0
	v15 = v13 - int32(48)
	m.G0 = v15
	v18 = v15 + int32(32)
	v19 = *(*int32)(unsafe.Add(mBase, uint32(l0)+36))
	v20 = *(*int32)(unsafe.Add(mBase, uint32(v19)+12))
	v24 = *(*int32)(unsafe.Add(mBase, uint32(v19)+8))
	v26 = base.I64_extend_i32_s(v20)*int64(30) + base.I64_extend_i32_s(v24)
	v35 = int64(32)
	v36 = int64(20)
	v38 = int64(base.Ui64(v26) >> (uint(v35) % 64))
	v41 = int64(4294967295)
	v42 = int64(500654080)
	v44 = v26 & v41
	v45 = v42 * v44
	v49 = int64(base.Ui64(v45)>>(uint(v35)%64)) + v42*v38
	v56 = v44*v36 + v49&v41
	*(*int64)(unsafe.Add(mBase, uint32(v18)+8)) = v26*int64(0) + v26>>(uint(int64(63))%64)*int64(86400000000) + v36*v38 + int64(base.Ui64(v49)>>(uint(v35)%64)) + int64(base.Ui64(v56)>>(uint(v35)%64))
	*(*int64)(unsafe.Add(mBase, uint32(v18))) = v45&v41 | v56<<(uint(v35)%64)
	v67 = *(*int64)(unsafe.Add(mBase, uint32(v15)+40))
	v68 = *(*int64)(unsafe.Add(mBase, uint32(v19)))
	v72 = *(*int64)(unsafe.Add(mBase, uint32(v15)+32))
	if int64(0) <= v67+v68>>(uint(int64(63))%64)+base.I64_extend_i32_u(base.B2i32(base.Ui64(v72+v68) < base.Ui64(v72))) {
		v79 = *(*int32)(unsafe.Add(mBase, uint32(l0)+52))
		v80 = *(*int32)(unsafe.Add(mBase, uint32(l0)+44))
		v81 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
		v82 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
		if v20 != int32(2147483647) {
			if v80 != 0 {
				v113 = int32(1475)
			} else {
				v113 = int32(1473)
			}
			v116 = F_DirectFunctionCall2Coll(m, v113, int32(0), v81, v19)
			mBase = m.M
			v119 = m.ExcPending
			if v119 != 0 {
				return int32(0)
			} else {
				v120 = int64(*(*int32)(unsafe.Add(mBase, uint32(v82)+8)))
				v121 = int64(*(*int32)(unsafe.Add(mBase, uint32(v82)+12)))
				v124 = v120 + v121*int64(30)
				v133 = int64(32)
				v134 = int64(20)
				v136 = int64(base.Ui64(v124) >> (uint(v133) % 64))
				v139 = int64(4294967295)
				v140 = int64(500654080)
				v142 = v124 & v139
				v143 = v140 * v142
				v147 = int64(base.Ui64(v143)>>(uint(v133)%64)) + v140*v136
				v154 = v142*v134 + v147&v139
				*(*int64)(unsafe.Add(mBase, uint32(v15)+8)) = v124*int64(0) + v124>>(uint(int64(63))%64)*int64(86400000000) + v134*v136 + int64(base.Ui64(v147)>>(uint(v133)%64)) + int64(base.Ui64(v154)>>(uint(v133)%64))
				*(*int64)(unsafe.Add(mBase, uint32(v15))) = v143&v139 | v154<<(uint(v133)%64)
				v166 = v15 + int32(16)
				v167 = int64(*(*int32)(unsafe.Add(mBase, uint32(v116)+8)))
				v168 = int64(*(*int32)(unsafe.Add(mBase, uint32(v116)+12)))
				v171 = v167 + v168*int64(30)
				v180 = int64(32)
				v181 = int64(20)
				v183 = int64(base.Ui64(v171) >> (uint(v180) % 64))
				v186 = int64(4294967295)
				v187 = int64(500654080)
				v189 = v171 & v186
				v190 = v187 * v189
				v194 = int64(base.Ui64(v190)>>(uint(v180)%64)) + v187*v183
				v201 = v189*v181 + v194&v186
				*(*int64)(unsafe.Add(mBase, uint32(v166)+8)) = v171*int64(0) + v171>>(uint(int64(63))%64)*int64(86400000000) + v181*v183 + int64(base.Ui64(v194)>>(uint(v180)%64)) + int64(base.Ui64(v201)>>(uint(v180)%64))
				*(*int64)(unsafe.Add(mBase, uint32(v166))) = v190&v186 | v201<<(uint(v180)%64)
				v212 = *(*int64)(unsafe.Add(mBase, uint32(v15)+24))
				v213 = *(*int64)(unsafe.Add(mBase, uint32(v116)))
				v214 = int64(63)
				v217 = *(*int64)(unsafe.Add(mBase, uint32(v15)+16))
				v218 = v217 + v213
				v221 = v212 + v213>>(uint(v214)%64) + base.I64_extend_i32_u(base.B2i32(base.Ui64(v218) < base.Ui64(v217)))
				v222 = *(*int64)(unsafe.Add(mBase, uint32(v15)+8))
				v223 = *(*int64)(unsafe.Add(mBase, uint32(v82)))
				v227 = *(*int64)(unsafe.Add(mBase, uint32(v15)))
				v228 = v223 + v227
				v231 = v222 + v223>>(uint(v214)%64) + base.I64_extend_i32_u(base.B2i32(base.Ui64(v228) < base.Ui64(v227)))
				if v221 == v231 {
					v235 = base.B2i32(base.Ui64(v228) <= base.Ui64(v218))
				} else {
					v235 = base.B2i32(v231 <= v221)
				}
				if v79 != 0 {
					v247 = v235
				} else {
					if v221 == v231 {
						v239 = base.B2i32(base.Ui64(v218) <= base.Ui64(v228))
					} else {
						v239 = base.B2i32(v221 <= v231)
					}
					v247 = v239
				}
				m.G0 = v15 + int32(48)
				return v247
			}
		} else {
			if v24 != int32(2147483647) {
				if v80 != 0 {
					v113 = int32(1475)
				} else {
					v113 = int32(1473)
				}
				v116 = F_DirectFunctionCall2Coll(m, v113, int32(0), v81, v19)
				mBase = m.M
				v119 = m.ExcPending
				if v119 != 0 {
					return int32(0)
				} else {
					v120 = int64(*(*int32)(unsafe.Add(mBase, uint32(v82)+8)))
					v121 = int64(*(*int32)(unsafe.Add(mBase, uint32(v82)+12)))
					v124 = v120 + v121*int64(30)
					v133 = int64(32)
					v134 = int64(20)
					v136 = int64(base.Ui64(v124) >> (uint(v133) % 64))
					v139 = int64(4294967295)
					v140 = int64(500654080)
					v142 = v124 & v139
					v143 = v140 * v142
					v147 = int64(base.Ui64(v143)>>(uint(v133)%64)) + v140*v136
					v154 = v142*v134 + v147&v139
					*(*int64)(unsafe.Add(mBase, uint32(v15)+8)) = v124*int64(0) + v124>>(uint(int64(63))%64)*int64(86400000000) + v134*v136 + int64(base.Ui64(v147)>>(uint(v133)%64)) + int64(base.Ui64(v154)>>(uint(v133)%64))
					*(*int64)(unsafe.Add(mBase, uint32(v15))) = v143&v139 | v154<<(uint(v133)%64)
					v166 = v15 + int32(16)
					v167 = int64(*(*int32)(unsafe.Add(mBase, uint32(v116)+8)))
					v168 = int64(*(*int32)(unsafe.Add(mBase, uint32(v116)+12)))
					v171 = v167 + v168*int64(30)
					v180 = int64(32)
					v181 = int64(20)
					v183 = int64(base.Ui64(v171) >> (uint(v180) % 64))
					v186 = int64(4294967295)
					v187 = int64(500654080)
					v189 = v171 & v186
					v190 = v187 * v189
					v194 = int64(base.Ui64(v190)>>(uint(v180)%64)) + v187*v183
					v201 = v189*v181 + v194&v186
					*(*int64)(unsafe.Add(mBase, uint32(v166)+8)) = v171*int64(0) + v171>>(uint(int64(63))%64)*int64(86400000000) + v181*v183 + int64(base.Ui64(v194)>>(uint(v180)%64)) + int64(base.Ui64(v201)>>(uint(v180)%64))
					*(*int64)(unsafe.Add(mBase, uint32(v166))) = v190&v186 | v201<<(uint(v180)%64)
					v212 = *(*int64)(unsafe.Add(mBase, uint32(v15)+24))
					v213 = *(*int64)(unsafe.Add(mBase, uint32(v116)))
					v214 = int64(63)
					v217 = *(*int64)(unsafe.Add(mBase, uint32(v15)+16))
					v218 = v217 + v213
					v221 = v212 + v213>>(uint(v214)%64) + base.I64_extend_i32_u(base.B2i32(base.Ui64(v218) < base.Ui64(v217)))
					v222 = *(*int64)(unsafe.Add(mBase, uint32(v15)+8))
					v223 = *(*int64)(unsafe.Add(mBase, uint32(v82)))
					v227 = *(*int64)(unsafe.Add(mBase, uint32(v15)))
					v228 = v223 + v227
					v231 = v222 + v223>>(uint(v214)%64) + base.I64_extend_i32_u(base.B2i32(base.Ui64(v228) < base.Ui64(v227)))
					if v221 == v231 {
						v235 = base.B2i32(base.Ui64(v228) <= base.Ui64(v218))
					} else {
						v235 = base.B2i32(v231 <= v221)
					}
					if v79 != 0 {
						v247 = v235
					} else {
						if v221 == v231 {
							v239 = base.B2i32(base.Ui64(v218) <= base.Ui64(v228))
						} else {
							v239 = base.B2i32(v221 <= v231)
						}
						v247 = v239
					}
					m.G0 = v15 + int32(48)
					return v247
				}
			} else {
				if v68 != int64(9223372036854775807) {
					if v80 != 0 {
						v113 = int32(1475)
					} else {
						v113 = int32(1473)
					}
					v116 = F_DirectFunctionCall2Coll(m, v113, int32(0), v81, v19)
					mBase = m.M
					v119 = m.ExcPending
					if v119 != 0 {
						return int32(0)
					} else {
						v120 = int64(*(*int32)(unsafe.Add(mBase, uint32(v82)+8)))
						v121 = int64(*(*int32)(unsafe.Add(mBase, uint32(v82)+12)))
						v124 = v120 + v121*int64(30)
						v133 = int64(32)
						v134 = int64(20)
						v136 = int64(base.Ui64(v124) >> (uint(v133) % 64))
						v139 = int64(4294967295)
						v140 = int64(500654080)
						v142 = v124 & v139
						v143 = v140 * v142
						v147 = int64(base.Ui64(v143)>>(uint(v133)%64)) + v140*v136
						v154 = v142*v134 + v147&v139
						*(*int64)(unsafe.Add(mBase, uint32(v15)+8)) = v124*int64(0) + v124>>(uint(int64(63))%64)*int64(86400000000) + v134*v136 + int64(base.Ui64(v147)>>(uint(v133)%64)) + int64(base.Ui64(v154)>>(uint(v133)%64))
						*(*int64)(unsafe.Add(mBase, uint32(v15))) = v143&v139 | v154<<(uint(v133)%64)
						v166 = v15 + int32(16)
						v167 = int64(*(*int32)(unsafe.Add(mBase, uint32(v116)+8)))
						v168 = int64(*(*int32)(unsafe.Add(mBase, uint32(v116)+12)))
						v171 = v167 + v168*int64(30)
						v180 = int64(32)
						v181 = int64(20)
						v183 = int64(base.Ui64(v171) >> (uint(v180) % 64))
						v186 = int64(4294967295)
						v187 = int64(500654080)
						v189 = v171 & v186
						v190 = v187 * v189
						v194 = int64(base.Ui64(v190)>>(uint(v180)%64)) + v187*v183
						v201 = v189*v181 + v194&v186
						*(*int64)(unsafe.Add(mBase, uint32(v166)+8)) = v171*int64(0) + v171>>(uint(int64(63))%64)*int64(86400000000) + v181*v183 + int64(base.Ui64(v194)>>(uint(v180)%64)) + int64(base.Ui64(v201)>>(uint(v180)%64))
						*(*int64)(unsafe.Add(mBase, uint32(v166))) = v190&v186 | v201<<(uint(v180)%64)
						v212 = *(*int64)(unsafe.Add(mBase, uint32(v15)+24))
						v213 = *(*int64)(unsafe.Add(mBase, uint32(v116)))
						v214 = int64(63)
						v217 = *(*int64)(unsafe.Add(mBase, uint32(v15)+16))
						v218 = v217 + v213
						v221 = v212 + v213>>(uint(v214)%64) + base.I64_extend_i32_u(base.B2i32(base.Ui64(v218) < base.Ui64(v217)))
						v222 = *(*int64)(unsafe.Add(mBase, uint32(v15)+8))
						v223 = *(*int64)(unsafe.Add(mBase, uint32(v82)))
						v227 = *(*int64)(unsafe.Add(mBase, uint32(v15)))
						v228 = v223 + v227
						v231 = v222 + v223>>(uint(v214)%64) + base.I64_extend_i32_u(base.B2i32(base.Ui64(v228) < base.Ui64(v227)))
						if v221 == v231 {
							v235 = base.B2i32(base.Ui64(v228) <= base.Ui64(v218))
						} else {
							v235 = base.B2i32(v231 <= v221)
						}
						if v79 != 0 {
							v247 = v235
						} else {
							if v221 == v231 {
								v239 = base.B2i32(base.Ui64(v218) <= base.Ui64(v228))
							} else {
								v239 = base.B2i32(v221 <= v231)
							}
							v247 = v239
						}
						m.G0 = v15 + int32(48)
						return v247
					}
				} else {
					v89 = *(*int32)(unsafe.Add(mBase, uint32(v81)+12))
					if v80 != 0 {
						v90 = int32(1475)
						if v89 != int32(2147483647) {
							v113 = v90
							v116 = F_DirectFunctionCall2Coll(m, v113, int32(0), v81, v19)
							mBase = m.M
							v119 = m.ExcPending
							if v119 != 0 {
								return int32(0)
							} else {
								v120 = int64(*(*int32)(unsafe.Add(mBase, uint32(v82)+8)))
								v121 = int64(*(*int32)(unsafe.Add(mBase, uint32(v82)+12)))
								v124 = v120 + v121*int64(30)
								v133 = int64(32)
								v134 = int64(20)
								v136 = int64(base.Ui64(v124) >> (uint(v133) % 64))
								v139 = int64(4294967295)
								v140 = int64(500654080)
								v142 = v124 & v139
								v143 = v140 * v142
								v147 = int64(base.Ui64(v143)>>(uint(v133)%64)) + v140*v136
								v154 = v142*v134 + v147&v139
								*(*int64)(unsafe.Add(mBase, uint32(v15)+8)) = v124*int64(0) + v124>>(uint(int64(63))%64)*int64(86400000000) + v134*v136 + int64(base.Ui64(v147)>>(uint(v133)%64)) + int64(base.Ui64(v154)>>(uint(v133)%64))
								*(*int64)(unsafe.Add(mBase, uint32(v15))) = v143&v139 | v154<<(uint(v133)%64)
								v166 = v15 + int32(16)
								v167 = int64(*(*int32)(unsafe.Add(mBase, uint32(v116)+8)))
								v168 = int64(*(*int32)(unsafe.Add(mBase, uint32(v116)+12)))
								v171 = v167 + v168*int64(30)
								v180 = int64(32)
								v181 = int64(20)
								v183 = int64(base.Ui64(v171) >> (uint(v180) % 64))
								v186 = int64(4294967295)
								v187 = int64(500654080)
								v189 = v171 & v186
								v190 = v187 * v189
								v194 = int64(base.Ui64(v190)>>(uint(v180)%64)) + v187*v183
								v201 = v189*v181 + v194&v186
								*(*int64)(unsafe.Add(mBase, uint32(v166)+8)) = v171*int64(0) + v171>>(uint(int64(63))%64)*int64(86400000000) + v181*v183 + int64(base.Ui64(v194)>>(uint(v180)%64)) + int64(base.Ui64(v201)>>(uint(v180)%64))
								*(*int64)(unsafe.Add(mBase, uint32(v166))) = v190&v186 | v201<<(uint(v180)%64)
								v212 = *(*int64)(unsafe.Add(mBase, uint32(v15)+24))
								v213 = *(*int64)(unsafe.Add(mBase, uint32(v116)))
								v214 = int64(63)
								v217 = *(*int64)(unsafe.Add(mBase, uint32(v15)+16))
								v218 = v217 + v213
								v221 = v212 + v213>>(uint(v214)%64) + base.I64_extend_i32_u(base.B2i32(base.Ui64(v218) < base.Ui64(v217)))
								v222 = *(*int64)(unsafe.Add(mBase, uint32(v15)+8))
								v223 = *(*int64)(unsafe.Add(mBase, uint32(v82)))
								v227 = *(*int64)(unsafe.Add(mBase, uint32(v15)))
								v228 = v223 + v227
								v231 = v222 + v223>>(uint(v214)%64) + base.I64_extend_i32_u(base.B2i32(base.Ui64(v228) < base.Ui64(v227)))
								if v221 == v231 {
									v235 = base.B2i32(base.Ui64(v228) <= base.Ui64(v218))
								} else {
									v235 = base.B2i32(v231 <= v221)
								}
								if v79 != 0 {
									v247 = v235
								} else {
									if v221 == v231 {
										v239 = base.B2i32(base.Ui64(v218) <= base.Ui64(v228))
									} else {
										v239 = base.B2i32(v221 <= v231)
									}
									v247 = v239
								}
								m.G0 = v15 + int32(48)
								return v247
							}
						} else {
							v93 = *(*int32)(unsafe.Add(mBase, uint32(v81)+8))
							if v93 != int32(2147483647) {
								v113 = v90
								v116 = F_DirectFunctionCall2Coll(m, v113, int32(0), v81, v19)
								mBase = m.M
								v119 = m.ExcPending
								if v119 != 0 {
									return int32(0)
								} else {
									v120 = int64(*(*int32)(unsafe.Add(mBase, uint32(v82)+8)))
									v121 = int64(*(*int32)(unsafe.Add(mBase, uint32(v82)+12)))
									v124 = v120 + v121*int64(30)
									v133 = int64(32)
									v134 = int64(20)
									v136 = int64(base.Ui64(v124) >> (uint(v133) % 64))
									v139 = int64(4294967295)
									v140 = int64(500654080)
									v142 = v124 & v139
									v143 = v140 * v142
									v147 = int64(base.Ui64(v143)>>(uint(v133)%64)) + v140*v136
									v154 = v142*v134 + v147&v139
									*(*int64)(unsafe.Add(mBase, uint32(v15)+8)) = v124*int64(0) + v124>>(uint(int64(63))%64)*int64(86400000000) + v134*v136 + int64(base.Ui64(v147)>>(uint(v133)%64)) + int64(base.Ui64(v154)>>(uint(v133)%64))
									*(*int64)(unsafe.Add(mBase, uint32(v15))) = v143&v139 | v154<<(uint(v133)%64)
									v166 = v15 + int32(16)
									v167 = int64(*(*int32)(unsafe.Add(mBase, uint32(v116)+8)))
									v168 = int64(*(*int32)(unsafe.Add(mBase, uint32(v116)+12)))
									v171 = v167 + v168*int64(30)
									v180 = int64(32)
									v181 = int64(20)
									v183 = int64(base.Ui64(v171) >> (uint(v180) % 64))
									v186 = int64(4294967295)
									v187 = int64(500654080)
									v189 = v171 & v186
									v190 = v187 * v189
									v194 = int64(base.Ui64(v190)>>(uint(v180)%64)) + v187*v183
									v201 = v189*v181 + v194&v186
									*(*int64)(unsafe.Add(mBase, uint32(v166)+8)) = v171*int64(0) + v171>>(uint(int64(63))%64)*int64(86400000000) + v181*v183 + int64(base.Ui64(v194)>>(uint(v180)%64)) + int64(base.Ui64(v201)>>(uint(v180)%64))
									*(*int64)(unsafe.Add(mBase, uint32(v166))) = v190&v186 | v201<<(uint(v180)%64)
									v212 = *(*int64)(unsafe.Add(mBase, uint32(v15)+24))
									v213 = *(*int64)(unsafe.Add(mBase, uint32(v116)))
									v214 = int64(63)
									v217 = *(*int64)(unsafe.Add(mBase, uint32(v15)+16))
									v218 = v217 + v213
									v221 = v212 + v213>>(uint(v214)%64) + base.I64_extend_i32_u(base.B2i32(base.Ui64(v218) < base.Ui64(v217)))
									v222 = *(*int64)(unsafe.Add(mBase, uint32(v15)+8))
									v223 = *(*int64)(unsafe.Add(mBase, uint32(v82)))
									v227 = *(*int64)(unsafe.Add(mBase, uint32(v15)))
									v228 = v223 + v227
									v231 = v222 + v223>>(uint(v214)%64) + base.I64_extend_i32_u(base.B2i32(base.Ui64(v228) < base.Ui64(v227)))
									if v221 == v231 {
										v235 = base.B2i32(base.Ui64(v228) <= base.Ui64(v218))
									} else {
										v235 = base.B2i32(v231 <= v221)
									}
									if v79 != 0 {
										v247 = v235
									} else {
										if v221 == v231 {
											v239 = base.B2i32(base.Ui64(v218) <= base.Ui64(v228))
										} else {
											v239 = base.B2i32(v221 <= v231)
										}
										v247 = v239
									}
									m.G0 = v15 + int32(48)
									return v247
								}
							} else {
								v96 = *(*int64)(unsafe.Add(mBase, uint32(v81)))
								if v96 != int64(9223372036854775807) {
									v113 = v90
									v116 = F_DirectFunctionCall2Coll(m, v113, int32(0), v81, v19)
									mBase = m.M
									v119 = m.ExcPending
									if v119 != 0 {
										return int32(0)
									} else {
										v120 = int64(*(*int32)(unsafe.Add(mBase, uint32(v82)+8)))
										v121 = int64(*(*int32)(unsafe.Add(mBase, uint32(v82)+12)))
										v124 = v120 + v121*int64(30)
										v133 = int64(32)
										v134 = int64(20)
										v136 = int64(base.Ui64(v124) >> (uint(v133) % 64))
										v139 = int64(4294967295)
										v140 = int64(500654080)
										v142 = v124 & v139
										v143 = v140 * v142
										v147 = int64(base.Ui64(v143)>>(uint(v133)%64)) + v140*v136
										v154 = v142*v134 + v147&v139
										*(*int64)(unsafe.Add(mBase, uint32(v15)+8)) = v124*int64(0) + v124>>(uint(int64(63))%64)*int64(86400000000) + v134*v136 + int64(base.Ui64(v147)>>(uint(v133)%64)) + int64(base.Ui64(v154)>>(uint(v133)%64))
										*(*int64)(unsafe.Add(mBase, uint32(v15))) = v143&v139 | v154<<(uint(v133)%64)
										v166 = v15 + int32(16)
										v167 = int64(*(*int32)(unsafe.Add(mBase, uint32(v116)+8)))
										v168 = int64(*(*int32)(unsafe.Add(mBase, uint32(v116)+12)))
										v171 = v167 + v168*int64(30)
										v180 = int64(32)
										v181 = int64(20)
										v183 = int64(base.Ui64(v171) >> (uint(v180) % 64))
										v186 = int64(4294967295)
										v187 = int64(500654080)
										v189 = v171 & v186
										v190 = v187 * v189
										v194 = int64(base.Ui64(v190)>>(uint(v180)%64)) + v187*v183
										v201 = v189*v181 + v194&v186
										*(*int64)(unsafe.Add(mBase, uint32(v166)+8)) = v171*int64(0) + v171>>(uint(int64(63))%64)*int64(86400000000) + v181*v183 + int64(base.Ui64(v194)>>(uint(v180)%64)) + int64(base.Ui64(v201)>>(uint(v180)%64))
										*(*int64)(unsafe.Add(mBase, uint32(v166))) = v190&v186 | v201<<(uint(v180)%64)
										v212 = *(*int64)(unsafe.Add(mBase, uint32(v15)+24))
										v213 = *(*int64)(unsafe.Add(mBase, uint32(v116)))
										v214 = int64(63)
										v217 = *(*int64)(unsafe.Add(mBase, uint32(v15)+16))
										v218 = v217 + v213
										v221 = v212 + v213>>(uint(v214)%64) + base.I64_extend_i32_u(base.B2i32(base.Ui64(v218) < base.Ui64(v217)))
										v222 = *(*int64)(unsafe.Add(mBase, uint32(v15)+8))
										v223 = *(*int64)(unsafe.Add(mBase, uint32(v82)))
										v227 = *(*int64)(unsafe.Add(mBase, uint32(v15)))
										v228 = v223 + v227
										v231 = v222 + v223>>(uint(v214)%64) + base.I64_extend_i32_u(base.B2i32(base.Ui64(v228) < base.Ui64(v227)))
										if v221 == v231 {
											v235 = base.B2i32(base.Ui64(v228) <= base.Ui64(v218))
										} else {
											v235 = base.B2i32(v231 <= v221)
										}
										if v79 != 0 {
											v247 = v235
										} else {
											if v221 == v231 {
												v239 = base.B2i32(base.Ui64(v218) <= base.Ui64(v228))
											} else {
												v239 = base.B2i32(v221 <= v231)
											}
											v247 = v239
										}
										m.G0 = v15 + int32(48)
										return v247
									}
								} else {
									v247 = int32(1)
									m.G0 = v15 + int32(48)
									return v247
								}
							}
						}
					} else {
						if v89 != int32(-2147483648) {
							v113 = int32(1473)
							v116 = F_DirectFunctionCall2Coll(m, v113, int32(0), v81, v19)
							mBase = m.M
							v119 = m.ExcPending
							if v119 != 0 {
								return int32(0)
							} else {
								v120 = int64(*(*int32)(unsafe.Add(mBase, uint32(v82)+8)))
								v121 = int64(*(*int32)(unsafe.Add(mBase, uint32(v82)+12)))
								v124 = v120 + v121*int64(30)
								v133 = int64(32)
								v134 = int64(20)
								v136 = int64(base.Ui64(v124) >> (uint(v133) % 64))
								v139 = int64(4294967295)
								v140 = int64(500654080)
								v142 = v124 & v139
								v143 = v140 * v142
								v147 = int64(base.Ui64(v143)>>(uint(v133)%64)) + v140*v136
								v154 = v142*v134 + v147&v139
								*(*int64)(unsafe.Add(mBase, uint32(v15)+8)) = v124*int64(0) + v124>>(uint(int64(63))%64)*int64(86400000000) + v134*v136 + int64(base.Ui64(v147)>>(uint(v133)%64)) + int64(base.Ui64(v154)>>(uint(v133)%64))
								*(*int64)(unsafe.Add(mBase, uint32(v15))) = v143&v139 | v154<<(uint(v133)%64)
								v166 = v15 + int32(16)
								v167 = int64(*(*int32)(unsafe.Add(mBase, uint32(v116)+8)))
								v168 = int64(*(*int32)(unsafe.Add(mBase, uint32(v116)+12)))
								v171 = v167 + v168*int64(30)
								v180 = int64(32)
								v181 = int64(20)
								v183 = int64(base.Ui64(v171) >> (uint(v180) % 64))
								v186 = int64(4294967295)
								v187 = int64(500654080)
								v189 = v171 & v186
								v190 = v187 * v189
								v194 = int64(base.Ui64(v190)>>(uint(v180)%64)) + v187*v183
								v201 = v189*v181 + v194&v186
								*(*int64)(unsafe.Add(mBase, uint32(v166)+8)) = v171*int64(0) + v171>>(uint(int64(63))%64)*int64(86400000000) + v181*v183 + int64(base.Ui64(v194)>>(uint(v180)%64)) + int64(base.Ui64(v201)>>(uint(v180)%64))
								*(*int64)(unsafe.Add(mBase, uint32(v166))) = v190&v186 | v201<<(uint(v180)%64)
								v212 = *(*int64)(unsafe.Add(mBase, uint32(v15)+24))
								v213 = *(*int64)(unsafe.Add(mBase, uint32(v116)))
								v214 = int64(63)
								v217 = *(*int64)(unsafe.Add(mBase, uint32(v15)+16))
								v218 = v217 + v213
								v221 = v212 + v213>>(uint(v214)%64) + base.I64_extend_i32_u(base.B2i32(base.Ui64(v218) < base.Ui64(v217)))
								v222 = *(*int64)(unsafe.Add(mBase, uint32(v15)+8))
								v223 = *(*int64)(unsafe.Add(mBase, uint32(v82)))
								v227 = *(*int64)(unsafe.Add(mBase, uint32(v15)))
								v228 = v223 + v227
								v231 = v222 + v223>>(uint(v214)%64) + base.I64_extend_i32_u(base.B2i32(base.Ui64(v228) < base.Ui64(v227)))
								if v221 == v231 {
									v235 = base.B2i32(base.Ui64(v228) <= base.Ui64(v218))
								} else {
									v235 = base.B2i32(v231 <= v221)
								}
								if v79 != 0 {
									v247 = v235
								} else {
									if v221 == v231 {
										v239 = base.B2i32(base.Ui64(v218) <= base.Ui64(v228))
									} else {
										v239 = base.B2i32(v221 <= v231)
									}
									v247 = v239
								}
								m.G0 = v15 + int32(48)
								return v247
							}
						} else {
							v102 = *(*int32)(unsafe.Add(mBase, uint32(v81)+8))
							if v102 != int32(-2147483648) {
								v113 = int32(1473)
								v116 = F_DirectFunctionCall2Coll(m, v113, int32(0), v81, v19)
								mBase = m.M
								v119 = m.ExcPending
								if v119 != 0 {
									return int32(0)
								} else {
									v120 = int64(*(*int32)(unsafe.Add(mBase, uint32(v82)+8)))
									v121 = int64(*(*int32)(unsafe.Add(mBase, uint32(v82)+12)))
									v124 = v120 + v121*int64(30)
									v133 = int64(32)
									v134 = int64(20)
									v136 = int64(base.Ui64(v124) >> (uint(v133) % 64))
									v139 = int64(4294967295)
									v140 = int64(500654080)
									v142 = v124 & v139
									v143 = v140 * v142
									v147 = int64(base.Ui64(v143)>>(uint(v133)%64)) + v140*v136
									v154 = v142*v134 + v147&v139
									*(*int64)(unsafe.Add(mBase, uint32(v15)+8)) = v124*int64(0) + v124>>(uint(int64(63))%64)*int64(86400000000) + v134*v136 + int64(base.Ui64(v147)>>(uint(v133)%64)) + int64(base.Ui64(v154)>>(uint(v133)%64))
									*(*int64)(unsafe.Add(mBase, uint32(v15))) = v143&v139 | v154<<(uint(v133)%64)
									v166 = v15 + int32(16)
									v167 = int64(*(*int32)(unsafe.Add(mBase, uint32(v116)+8)))
									v168 = int64(*(*int32)(unsafe.Add(mBase, uint32(v116)+12)))
									v171 = v167 + v168*int64(30)
									v180 = int64(32)
									v181 = int64(20)
									v183 = int64(base.Ui64(v171) >> (uint(v180) % 64))
									v186 = int64(4294967295)
									v187 = int64(500654080)
									v189 = v171 & v186
									v190 = v187 * v189
									v194 = int64(base.Ui64(v190)>>(uint(v180)%64)) + v187*v183
									v201 = v189*v181 + v194&v186
									*(*int64)(unsafe.Add(mBase, uint32(v166)+8)) = v171*int64(0) + v171>>(uint(int64(63))%64)*int64(86400000000) + v181*v183 + int64(base.Ui64(v194)>>(uint(v180)%64)) + int64(base.Ui64(v201)>>(uint(v180)%64))
									*(*int64)(unsafe.Add(mBase, uint32(v166))) = v190&v186 | v201<<(uint(v180)%64)
									v212 = *(*int64)(unsafe.Add(mBase, uint32(v15)+24))
									v213 = *(*int64)(unsafe.Add(mBase, uint32(v116)))
									v214 = int64(63)
									v217 = *(*int64)(unsafe.Add(mBase, uint32(v15)+16))
									v218 = v217 + v213
									v221 = v212 + v213>>(uint(v214)%64) + base.I64_extend_i32_u(base.B2i32(base.Ui64(v218) < base.Ui64(v217)))
									v222 = *(*int64)(unsafe.Add(mBase, uint32(v15)+8))
									v223 = *(*int64)(unsafe.Add(mBase, uint32(v82)))
									v227 = *(*int64)(unsafe.Add(mBase, uint32(v15)))
									v228 = v223 + v227
									v231 = v222 + v223>>(uint(v214)%64) + base.I64_extend_i32_u(base.B2i32(base.Ui64(v228) < base.Ui64(v227)))
									if v221 == v231 {
										v235 = base.B2i32(base.Ui64(v228) <= base.Ui64(v218))
									} else {
										v235 = base.B2i32(v231 <= v221)
									}
									if v79 != 0 {
										v247 = v235
									} else {
										if v221 == v231 {
											v239 = base.B2i32(base.Ui64(v218) <= base.Ui64(v228))
										} else {
											v239 = base.B2i32(v221 <= v231)
										}
										v247 = v239
									}
									m.G0 = v15 + int32(48)
									return v247
								}
							} else {
								v105 = *(*int64)(unsafe.Add(mBase, uint32(v81)))
								if v105 != int64(-9223372036854775807-1) {
									v113 = int32(1473)
									v116 = F_DirectFunctionCall2Coll(m, v113, int32(0), v81, v19)
									mBase = m.M
									v119 = m.ExcPending
									if v119 != 0 {
										return int32(0)
									} else {
										v120 = int64(*(*int32)(unsafe.Add(mBase, uint32(v82)+8)))
										v121 = int64(*(*int32)(unsafe.Add(mBase, uint32(v82)+12)))
										v124 = v120 + v121*int64(30)
										v133 = int64(32)
										v134 = int64(20)
										v136 = int64(base.Ui64(v124) >> (uint(v133) % 64))
										v139 = int64(4294967295)
										v140 = int64(500654080)
										v142 = v124 & v139
										v143 = v140 * v142
										v147 = int64(base.Ui64(v143)>>(uint(v133)%64)) + v140*v136
										v154 = v142*v134 + v147&v139
										*(*int64)(unsafe.Add(mBase, uint32(v15)+8)) = v124*int64(0) + v124>>(uint(int64(63))%64)*int64(86400000000) + v134*v136 + int64(base.Ui64(v147)>>(uint(v133)%64)) + int64(base.Ui64(v154)>>(uint(v133)%64))
										*(*int64)(unsafe.Add(mBase, uint32(v15))) = v143&v139 | v154<<(uint(v133)%64)
										v166 = v15 + int32(16)
										v167 = int64(*(*int32)(unsafe.Add(mBase, uint32(v116)+8)))
										v168 = int64(*(*int32)(unsafe.Add(mBase, uint32(v116)+12)))
										v171 = v167 + v168*int64(30)
										v180 = int64(32)
										v181 = int64(20)
										v183 = int64(base.Ui64(v171) >> (uint(v180) % 64))
										v186 = int64(4294967295)
										v187 = int64(500654080)
										v189 = v171 & v186
										v190 = v187 * v189
										v194 = int64(base.Ui64(v190)>>(uint(v180)%64)) + v187*v183
										v201 = v189*v181 + v194&v186
										*(*int64)(unsafe.Add(mBase, uint32(v166)+8)) = v171*int64(0) + v171>>(uint(int64(63))%64)*int64(86400000000) + v181*v183 + int64(base.Ui64(v194)>>(uint(v180)%64)) + int64(base.Ui64(v201)>>(uint(v180)%64))
										*(*int64)(unsafe.Add(mBase, uint32(v166))) = v190&v186 | v201<<(uint(v180)%64)
										v212 = *(*int64)(unsafe.Add(mBase, uint32(v15)+24))
										v213 = *(*int64)(unsafe.Add(mBase, uint32(v116)))
										v214 = int64(63)
										v217 = *(*int64)(unsafe.Add(mBase, uint32(v15)+16))
										v218 = v217 + v213
										v221 = v212 + v213>>(uint(v214)%64) + base.I64_extend_i32_u(base.B2i32(base.Ui64(v218) < base.Ui64(v217)))
										v222 = *(*int64)(unsafe.Add(mBase, uint32(v15)+8))
										v223 = *(*int64)(unsafe.Add(mBase, uint32(v82)))
										v227 = *(*int64)(unsafe.Add(mBase, uint32(v15)))
										v228 = v223 + v227
										v231 = v222 + v223>>(uint(v214)%64) + base.I64_extend_i32_u(base.B2i32(base.Ui64(v228) < base.Ui64(v227)))
										if v221 == v231 {
											v235 = base.B2i32(base.Ui64(v228) <= base.Ui64(v218))
										} else {
											v235 = base.B2i32(v231 <= v221)
										}
										if v79 != 0 {
											v247 = v235
										} else {
											if v221 == v231 {
												v239 = base.B2i32(base.Ui64(v218) <= base.Ui64(v228))
											} else {
												v239 = base.B2i32(v221 <= v231)
											}
											v247 = v239
										}
										m.G0 = v15 + int32(48)
										return v247
									}
								} else {
									v247 = int32(1)
									m.G0 = v15 + int32(48)
									return v247
								}
							}
						}
					}
				}
			}
		}
	} else {
		F_errstart_cold(m, int32(21), int32(0))
		mBase = m.M
		v255 = m.ExcPending
		if v255 != 0 {
			return int32(0)
		} else {
			F_errcode(m, int32(50593922))
			mBase = m.M
			v258 = m.ExcPending
			if v258 != 0 {
				return int32(0)
			} else {
				F_errmsg(m, int32(239981), int32(0))
				mBase = m.M
				v262 = m.ExcPending
				if v262 != 0 {
					return int32(0)
				} else {
					F_errfinish(m, int32(474326), int32(3947), int32(295208))
					mBase = m.M
					v267 = m.ExcPending
					if v267 != 0 {
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
	var v4 int32
	_ = v4
	var v9 int32
	_ = v9
	var v10 int32
	_ = v10
	var v12 int32
	_ = v12
	var v16 int32
	_ = v16
	var v18 int32
	_ = v18
	v4 = int32(*(*uint8)(unsafe.Add(mBase, _consts[112])))
	if v4 == int32(1) {
		v9 = *(*int32)(unsafe.Add(mBase, _consts[113]))
		v10 = *(*int32)(unsafe.Add(mBase, uint32(v9)+316))
		v12 = base.B2i32(v10 != int32(2))
		*(*uint8)(unsafe.Add(mBase, _consts[112])) = uint8(v12)
		if v10 != int32(2) {
			v16 = int32(260641)
		} else {
			v16 = int32(323540)
		}
		v18 = v16
	} else {
		v18 = int32(323540)
	}
	return v18
}
