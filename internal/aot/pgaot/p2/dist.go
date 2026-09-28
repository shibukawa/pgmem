package p2

import (
	base "github.com/shibukawa/pgmem/internal/aot/pgaot/base"
	"math"
	"unsafe"
)

func F_dist_cpoint(m *base.Module, l0 int32) int64 {
	mBase := m.M
	_ = mBase
	var v6 int32
	_ = v6
	var v7 float64
	_ = v7
	var v8 int32
	_ = v8
	var v9 float64
	_ = v9
	var v10 float64
	_ = v10
	var v12 float64
	_ = v12
	var v23 float64
	_ = v23
	var v26 int32
	_ = v26
	var v27 float64
	_ = v27
	var v28 float64
	_ = v28
	var v29 float64
	_ = v29
	var v30 float64
	_ = v30
	var v32 float64
	_ = v32
	var v43 float64
	_ = v43
	var v44 int32
	_ = v44
	var v45 float64
	_ = v45
	var v54 int32
	_ = v54
	var v56 int32
	_ = v56
	var v58 float64
	_ = v58
	var v59 float64
	_ = v59
	var v62 int32
	_ = v62
	var v63 float64
	_ = v63
	var v64 int64
	_ = v64
	var v66 int64
	_ = v66
	var v69 float64
	_ = v69
	var v72 int64
	_ = v72
	var v74 int64
	_ = v74
	var v85 float64
	_ = v85
	var v93 float64
	_ = v93
	var v98 float64
	_ = v98
	var v99 float64
	_ = v99
	var v100 float64
	_ = v100
	var v109 float64
	_ = v109
	var v110 float64
	_ = v110
	var v112 float64
	_ = v112
	var v114 float64
	_ = v114
	var v121 float64
	_ = v121
	var v127 float64
	_ = v127
	var v128 float64
	_ = v128
	var v130 float64
	_ = v130
	var v143 float64
	_ = v143
	var v144 int32
	_ = v144
	var v145 float64
	_ = v145
	var v146 float64
	_ = v146
	var v149 float64
	_ = v149
	v6 = *(*int32)(unsafe.Add(mBase, uint32(l0)+40))
	v7 = *(*float64)(unsafe.Add(mBase, uint32(v6)))
	v8 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
	v9 = *(*float64)(unsafe.Add(mBase, uint32(v8)))
	v10 = base.F64_sub(v7, v9)
	v12 = math.Float64frombits(uint64(0x7ff0000000000000))
	if base.F64_ne(base.F64_abs(v10), v12)|base.F64_eq(base.F64_abs(v7), v12)|base.F64_eq(base.F64_abs(v9), v12) != 0 {
		v27 = v10
		v28 = *(*float64)(unsafe.Add(mBase, uint32(v6)+8))
		v29 = *(*float64)(unsafe.Add(mBase, uint32(v8)+8))
		v30 = base.F64_sub(v28, v29)
		v32 = math.Float64frombits(uint64(0x7ff0000000000000))
		if base.F64_ne(base.F64_abs(v30), v32)|base.F64_eq(base.F64_abs(v28), v32)|base.F64_eq(base.F64_abs(v29), v32) != 0 {
			v45 = v30
			v54 = m.G0
			v56 = v54 - int32(32)
			m.G0 = v56
			v58 = base.F64_abs(v27)
			v59 = base.F64_abs(v45)
			v62 = base.B2i32(base.Ui64(base.I64_reinterpret_f64(v58)) < base.Ui64(base.I64_reinterpret_f64(v59)))
			if base.Ui64(base.I64_reinterpret_f64(v58)) < base.Ui64(base.I64_reinterpret_f64(v59)) {
				v63 = v58
			} else {
				v63 = v59
			}
			v64 = base.I64_reinterpret_f64(v63)
			v66 = int64(base.Ui64(v64) >> (uint(int64(52)) % 64))
			if v66 == int64(2047) {
				v121 = v63
			} else {
				if base.Ui64(base.I64_reinterpret_f64(v58)) < base.Ui64(base.I64_reinterpret_f64(v59)) {
					v69 = v59
				} else {
					v69 = v58
				}
				if v64 == int64(0) {
					v121 = v69
				} else {
					v72 = base.I64_reinterpret_f64(v69)
					v74 = int64(base.Ui64(v72) >> (uint(int64(52)) % 64))
					if v74 == int64(2047) {
						v121 = v69
					} else {
						if int32(65) <= base.I32_wrap_i64(v74)-base.I32_wrap_i64(v66) {
							v121 = base.F64_add(v58, v59)
						} else {
							if base.Ui64(int64(6908521828386340864)) <= base.Ui64(v72) {
								v85 = float64(1.90109156629516e-211)
								v98 = base.F64_mul(v69, v85)
								v99 = base.F64_mul(v63, v85)
								v100 = float64(5.260135901548374e+210)
							} else {
								if base.Ui64(int64(2580562586483294207)) < base.Ui64(v64) {
									v98 = v69
									v99 = v63
									v100 = float64(1)
								} else {
									v93 = float64(5.260135901548374e+210)
									v98 = base.F64_mul(v69, v93)
									v99 = base.F64_mul(v63, v93)
									v100 = float64(1.90109156629516e-211)
								}
							}
							F_sq(m, v56+int32(24), v56+int32(16), v98)
							mBase = m.M
							F_sq(m, v56+int32(8), v56, v99)
							mBase = m.M
							v109 = *(*float64)(unsafe.Add(mBase, uint32(v56)))
							v110 = *(*float64)(unsafe.Add(mBase, uint32(v56)+16))
							v112 = *(*float64)(unsafe.Add(mBase, uint32(v56)+8))
							v114 = *(*float64)(unsafe.Add(mBase, uint32(v56)+24))
							v121 = base.F64_mul(v100, base.F64_sqrt(base.F64_add(base.F64_add(base.F64_add(v109, v110), v112), v114)))
						}
					}
				}
			}
			m.G0 = v56 + int32(32)
			v127 = *(*float64)(unsafe.Add(mBase, uint32(v8)+16))
			v128 = base.F64_sub(v121, v127)
			v130 = math.Float64frombits(uint64(0x7ff0000000000000))
			if base.F64_ne(base.F64_abs(v128), v130)|base.F64_eq(base.F64_abs(v121), v130)|base.F64_eq(base.F64_abs(v127), v130) == int32(0) {
				v143 = F_float_overflow_error_ext(m, int32(0))
				mBase = m.M
				v144 = m.ExcPending
				if v144 != 0 {
					return int64(0)
				} else {
					v145 = v143
					v146 = float64(0)
					if base.F64_lt(v145, v146) != 0 {
						v149 = v146
					} else {
						v149 = v145
					}
					return base.I64_reinterpret_f64(v149)
				}
			} else {
				v145 = v128
				v146 = float64(0)
				if base.F64_lt(v145, v146) != 0 {
					v149 = v146
				} else {
					v149 = v145
				}
				return base.I64_reinterpret_f64(v149)
			}
		} else {
			v43 = F_float_overflow_error_ext(m, int32(0))
			mBase = m.M
			v44 = m.ExcPending
			if v44 != 0 {
				return int64(0)
			} else {
				v45 = v43
				v54 = m.G0
				v56 = v54 - int32(32)
				m.G0 = v56
				v58 = base.F64_abs(v27)
				v59 = base.F64_abs(v45)
				v62 = base.B2i32(base.Ui64(base.I64_reinterpret_f64(v58)) < base.Ui64(base.I64_reinterpret_f64(v59)))
				if base.Ui64(base.I64_reinterpret_f64(v58)) < base.Ui64(base.I64_reinterpret_f64(v59)) {
					v63 = v58
				} else {
					v63 = v59
				}
				v64 = base.I64_reinterpret_f64(v63)
				v66 = int64(base.Ui64(v64) >> (uint(int64(52)) % 64))
				if v66 == int64(2047) {
					v121 = v63
				} else {
					if base.Ui64(base.I64_reinterpret_f64(v58)) < base.Ui64(base.I64_reinterpret_f64(v59)) {
						v69 = v59
					} else {
						v69 = v58
					}
					if v64 == int64(0) {
						v121 = v69
					} else {
						v72 = base.I64_reinterpret_f64(v69)
						v74 = int64(base.Ui64(v72) >> (uint(int64(52)) % 64))
						if v74 == int64(2047) {
							v121 = v69
						} else {
							if int32(65) <= base.I32_wrap_i64(v74)-base.I32_wrap_i64(v66) {
								v121 = base.F64_add(v58, v59)
							} else {
								if base.Ui64(int64(6908521828386340864)) <= base.Ui64(v72) {
									v85 = float64(1.90109156629516e-211)
									v98 = base.F64_mul(v69, v85)
									v99 = base.F64_mul(v63, v85)
									v100 = float64(5.260135901548374e+210)
								} else {
									if base.Ui64(int64(2580562586483294207)) < base.Ui64(v64) {
										v98 = v69
										v99 = v63
										v100 = float64(1)
									} else {
										v93 = float64(5.260135901548374e+210)
										v98 = base.F64_mul(v69, v93)
										v99 = base.F64_mul(v63, v93)
										v100 = float64(1.90109156629516e-211)
									}
								}
								F_sq(m, v56+int32(24), v56+int32(16), v98)
								mBase = m.M
								F_sq(m, v56+int32(8), v56, v99)
								mBase = m.M
								v109 = *(*float64)(unsafe.Add(mBase, uint32(v56)))
								v110 = *(*float64)(unsafe.Add(mBase, uint32(v56)+16))
								v112 = *(*float64)(unsafe.Add(mBase, uint32(v56)+8))
								v114 = *(*float64)(unsafe.Add(mBase, uint32(v56)+24))
								v121 = base.F64_mul(v100, base.F64_sqrt(base.F64_add(base.F64_add(base.F64_add(v109, v110), v112), v114)))
							}
						}
					}
				}
				m.G0 = v56 + int32(32)
				v127 = *(*float64)(unsafe.Add(mBase, uint32(v8)+16))
				v128 = base.F64_sub(v121, v127)
				v130 = math.Float64frombits(uint64(0x7ff0000000000000))
				if base.F64_ne(base.F64_abs(v128), v130)|base.F64_eq(base.F64_abs(v121), v130)|base.F64_eq(base.F64_abs(v127), v130) == int32(0) {
					v143 = F_float_overflow_error_ext(m, int32(0))
					mBase = m.M
					v144 = m.ExcPending
					if v144 != 0 {
						return int64(0)
					} else {
						v145 = v143
						v146 = float64(0)
						if base.F64_lt(v145, v146) != 0 {
							v149 = v146
						} else {
							v149 = v145
						}
						return base.I64_reinterpret_f64(v149)
					}
				} else {
					v145 = v128
					v146 = float64(0)
					if base.F64_lt(v145, v146) != 0 {
						v149 = v146
					} else {
						v149 = v145
					}
					return base.I64_reinterpret_f64(v149)
				}
			}
		}
	} else {
		v23 = F_float_overflow_error_ext(m, int32(0))
		mBase = m.M
		v26 = m.ExcPending
		if v26 != 0 {
			return int64(0)
		} else {
			v27 = v23
			v28 = *(*float64)(unsafe.Add(mBase, uint32(v6)+8))
			v29 = *(*float64)(unsafe.Add(mBase, uint32(v8)+8))
			v30 = base.F64_sub(v28, v29)
			v32 = math.Float64frombits(uint64(0x7ff0000000000000))
			if base.F64_ne(base.F64_abs(v30), v32)|base.F64_eq(base.F64_abs(v28), v32)|base.F64_eq(base.F64_abs(v29), v32) != 0 {
				v45 = v30
				v54 = m.G0
				v56 = v54 - int32(32)
				m.G0 = v56
				v58 = base.F64_abs(v27)
				v59 = base.F64_abs(v45)
				v62 = base.B2i32(base.Ui64(base.I64_reinterpret_f64(v58)) < base.Ui64(base.I64_reinterpret_f64(v59)))
				if base.Ui64(base.I64_reinterpret_f64(v58)) < base.Ui64(base.I64_reinterpret_f64(v59)) {
					v63 = v58
				} else {
					v63 = v59
				}
				v64 = base.I64_reinterpret_f64(v63)
				v66 = int64(base.Ui64(v64) >> (uint(int64(52)) % 64))
				if v66 == int64(2047) {
					v121 = v63
				} else {
					if base.Ui64(base.I64_reinterpret_f64(v58)) < base.Ui64(base.I64_reinterpret_f64(v59)) {
						v69 = v59
					} else {
						v69 = v58
					}
					if v64 == int64(0) {
						v121 = v69
					} else {
						v72 = base.I64_reinterpret_f64(v69)
						v74 = int64(base.Ui64(v72) >> (uint(int64(52)) % 64))
						if v74 == int64(2047) {
							v121 = v69
						} else {
							if int32(65) <= base.I32_wrap_i64(v74)-base.I32_wrap_i64(v66) {
								v121 = base.F64_add(v58, v59)
							} else {
								if base.Ui64(int64(6908521828386340864)) <= base.Ui64(v72) {
									v85 = float64(1.90109156629516e-211)
									v98 = base.F64_mul(v69, v85)
									v99 = base.F64_mul(v63, v85)
									v100 = float64(5.260135901548374e+210)
								} else {
									if base.Ui64(int64(2580562586483294207)) < base.Ui64(v64) {
										v98 = v69
										v99 = v63
										v100 = float64(1)
									} else {
										v93 = float64(5.260135901548374e+210)
										v98 = base.F64_mul(v69, v93)
										v99 = base.F64_mul(v63, v93)
										v100 = float64(1.90109156629516e-211)
									}
								}
								F_sq(m, v56+int32(24), v56+int32(16), v98)
								mBase = m.M
								F_sq(m, v56+int32(8), v56, v99)
								mBase = m.M
								v109 = *(*float64)(unsafe.Add(mBase, uint32(v56)))
								v110 = *(*float64)(unsafe.Add(mBase, uint32(v56)+16))
								v112 = *(*float64)(unsafe.Add(mBase, uint32(v56)+8))
								v114 = *(*float64)(unsafe.Add(mBase, uint32(v56)+24))
								v121 = base.F64_mul(v100, base.F64_sqrt(base.F64_add(base.F64_add(base.F64_add(v109, v110), v112), v114)))
							}
						}
					}
				}
				m.G0 = v56 + int32(32)
				v127 = *(*float64)(unsafe.Add(mBase, uint32(v8)+16))
				v128 = base.F64_sub(v121, v127)
				v130 = math.Float64frombits(uint64(0x7ff0000000000000))
				if base.F64_ne(base.F64_abs(v128), v130)|base.F64_eq(base.F64_abs(v121), v130)|base.F64_eq(base.F64_abs(v127), v130) == int32(0) {
					v143 = F_float_overflow_error_ext(m, int32(0))
					mBase = m.M
					v144 = m.ExcPending
					if v144 != 0 {
						return int64(0)
					} else {
						v145 = v143
						v146 = float64(0)
						if base.F64_lt(v145, v146) != 0 {
							v149 = v146
						} else {
							v149 = v145
						}
						return base.I64_reinterpret_f64(v149)
					}
				} else {
					v145 = v128
					v146 = float64(0)
					if base.F64_lt(v145, v146) != 0 {
						v149 = v146
					} else {
						v149 = v145
					}
					return base.I64_reinterpret_f64(v149)
				}
			} else {
				v43 = F_float_overflow_error_ext(m, int32(0))
				mBase = m.M
				v44 = m.ExcPending
				if v44 != 0 {
					return int64(0)
				} else {
					v45 = v43
					v54 = m.G0
					v56 = v54 - int32(32)
					m.G0 = v56
					v58 = base.F64_abs(v27)
					v59 = base.F64_abs(v45)
					v62 = base.B2i32(base.Ui64(base.I64_reinterpret_f64(v58)) < base.Ui64(base.I64_reinterpret_f64(v59)))
					if base.Ui64(base.I64_reinterpret_f64(v58)) < base.Ui64(base.I64_reinterpret_f64(v59)) {
						v63 = v58
					} else {
						v63 = v59
					}
					v64 = base.I64_reinterpret_f64(v63)
					v66 = int64(base.Ui64(v64) >> (uint(int64(52)) % 64))
					if v66 == int64(2047) {
						v121 = v63
					} else {
						if base.Ui64(base.I64_reinterpret_f64(v58)) < base.Ui64(base.I64_reinterpret_f64(v59)) {
							v69 = v59
						} else {
							v69 = v58
						}
						if v64 == int64(0) {
							v121 = v69
						} else {
							v72 = base.I64_reinterpret_f64(v69)
							v74 = int64(base.Ui64(v72) >> (uint(int64(52)) % 64))
							if v74 == int64(2047) {
								v121 = v69
							} else {
								if int32(65) <= base.I32_wrap_i64(v74)-base.I32_wrap_i64(v66) {
									v121 = base.F64_add(v58, v59)
								} else {
									if base.Ui64(int64(6908521828386340864)) <= base.Ui64(v72) {
										v85 = float64(1.90109156629516e-211)
										v98 = base.F64_mul(v69, v85)
										v99 = base.F64_mul(v63, v85)
										v100 = float64(5.260135901548374e+210)
									} else {
										if base.Ui64(int64(2580562586483294207)) < base.Ui64(v64) {
											v98 = v69
											v99 = v63
											v100 = float64(1)
										} else {
											v93 = float64(5.260135901548374e+210)
											v98 = base.F64_mul(v69, v93)
											v99 = base.F64_mul(v63, v93)
											v100 = float64(1.90109156629516e-211)
										}
									}
									F_sq(m, v56+int32(24), v56+int32(16), v98)
									mBase = m.M
									F_sq(m, v56+int32(8), v56, v99)
									mBase = m.M
									v109 = *(*float64)(unsafe.Add(mBase, uint32(v56)))
									v110 = *(*float64)(unsafe.Add(mBase, uint32(v56)+16))
									v112 = *(*float64)(unsafe.Add(mBase, uint32(v56)+8))
									v114 = *(*float64)(unsafe.Add(mBase, uint32(v56)+24))
									v121 = base.F64_mul(v100, base.F64_sqrt(base.F64_add(base.F64_add(base.F64_add(v109, v110), v112), v114)))
								}
							}
						}
					}
					m.G0 = v56 + int32(32)
					v127 = *(*float64)(unsafe.Add(mBase, uint32(v8)+16))
					v128 = base.F64_sub(v121, v127)
					v130 = math.Float64frombits(uint64(0x7ff0000000000000))
					if base.F64_ne(base.F64_abs(v128), v130)|base.F64_eq(base.F64_abs(v121), v130)|base.F64_eq(base.F64_abs(v127), v130) == int32(0) {
						v143 = F_float_overflow_error_ext(m, int32(0))
						mBase = m.M
						v144 = m.ExcPending
						if v144 != 0 {
							return int64(0)
						} else {
							v145 = v143
							v146 = float64(0)
							if base.F64_lt(v145, v146) != 0 {
								v149 = v146
							} else {
								v149 = v145
							}
							return base.I64_reinterpret_f64(v149)
						}
					} else {
						v145 = v128
						v146 = float64(0)
						if base.F64_lt(v145, v146) != 0 {
							v149 = v146
						} else {
							v149 = v145
						}
						return base.I64_reinterpret_f64(v149)
					}
				}
			}
		}
	}
}
func F_dist_pc(m *base.Module, l0 int32) int64 {
	mBase := m.M
	_ = mBase
	var v6 int32
	_ = v6
	var v7 float64
	_ = v7
	var v8 int32
	_ = v8
	var v9 float64
	_ = v9
	var v10 float64
	_ = v10
	var v12 float64
	_ = v12
	var v23 float64
	_ = v23
	var v26 int32
	_ = v26
	var v27 float64
	_ = v27
	var v28 float64
	_ = v28
	var v29 float64
	_ = v29
	var v30 float64
	_ = v30
	var v32 float64
	_ = v32
	var v43 float64
	_ = v43
	var v44 int32
	_ = v44
	var v45 float64
	_ = v45
	var v54 int32
	_ = v54
	var v56 int32
	_ = v56
	var v58 float64
	_ = v58
	var v59 float64
	_ = v59
	var v62 int32
	_ = v62
	var v63 float64
	_ = v63
	var v64 int64
	_ = v64
	var v66 int64
	_ = v66
	var v69 float64
	_ = v69
	var v72 int64
	_ = v72
	var v74 int64
	_ = v74
	var v85 float64
	_ = v85
	var v93 float64
	_ = v93
	var v98 float64
	_ = v98
	var v99 float64
	_ = v99
	var v100 float64
	_ = v100
	var v109 float64
	_ = v109
	var v110 float64
	_ = v110
	var v112 float64
	_ = v112
	var v114 float64
	_ = v114
	var v121 float64
	_ = v121
	var v127 float64
	_ = v127
	var v128 float64
	_ = v128
	var v130 float64
	_ = v130
	var v143 float64
	_ = v143
	var v144 int32
	_ = v144
	var v145 float64
	_ = v145
	var v146 float64
	_ = v146
	var v149 float64
	_ = v149
	v6 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
	v7 = *(*float64)(unsafe.Add(mBase, uint32(v6)))
	v8 = *(*int32)(unsafe.Add(mBase, uint32(l0)+40))
	v9 = *(*float64)(unsafe.Add(mBase, uint32(v8)))
	v10 = base.F64_sub(v7, v9)
	v12 = math.Float64frombits(uint64(0x7ff0000000000000))
	if base.F64_ne(base.F64_abs(v10), v12)|base.F64_eq(base.F64_abs(v7), v12)|base.F64_eq(base.F64_abs(v9), v12) != 0 {
		v27 = v10
		v28 = *(*float64)(unsafe.Add(mBase, uint32(v6)+8))
		v29 = *(*float64)(unsafe.Add(mBase, uint32(v8)+8))
		v30 = base.F64_sub(v28, v29)
		v32 = math.Float64frombits(uint64(0x7ff0000000000000))
		if base.F64_ne(base.F64_abs(v30), v32)|base.F64_eq(base.F64_abs(v28), v32)|base.F64_eq(base.F64_abs(v29), v32) != 0 {
			v45 = v30
			v54 = m.G0
			v56 = v54 - int32(32)
			m.G0 = v56
			v58 = base.F64_abs(v27)
			v59 = base.F64_abs(v45)
			v62 = base.B2i32(base.Ui64(base.I64_reinterpret_f64(v58)) < base.Ui64(base.I64_reinterpret_f64(v59)))
			if base.Ui64(base.I64_reinterpret_f64(v58)) < base.Ui64(base.I64_reinterpret_f64(v59)) {
				v63 = v58
			} else {
				v63 = v59
			}
			v64 = base.I64_reinterpret_f64(v63)
			v66 = int64(base.Ui64(v64) >> (uint(int64(52)) % 64))
			if v66 == int64(2047) {
				v121 = v63
			} else {
				if base.Ui64(base.I64_reinterpret_f64(v58)) < base.Ui64(base.I64_reinterpret_f64(v59)) {
					v69 = v59
				} else {
					v69 = v58
				}
				if v64 == int64(0) {
					v121 = v69
				} else {
					v72 = base.I64_reinterpret_f64(v69)
					v74 = int64(base.Ui64(v72) >> (uint(int64(52)) % 64))
					if v74 == int64(2047) {
						v121 = v69
					} else {
						if int32(65) <= base.I32_wrap_i64(v74)-base.I32_wrap_i64(v66) {
							v121 = base.F64_add(v58, v59)
						} else {
							if base.Ui64(int64(6908521828386340864)) <= base.Ui64(v72) {
								v85 = float64(1.90109156629516e-211)
								v98 = base.F64_mul(v69, v85)
								v99 = base.F64_mul(v63, v85)
								v100 = float64(5.260135901548374e+210)
							} else {
								if base.Ui64(int64(2580562586483294207)) < base.Ui64(v64) {
									v98 = v69
									v99 = v63
									v100 = float64(1)
								} else {
									v93 = float64(5.260135901548374e+210)
									v98 = base.F64_mul(v69, v93)
									v99 = base.F64_mul(v63, v93)
									v100 = float64(1.90109156629516e-211)
								}
							}
							F_sq(m, v56+int32(24), v56+int32(16), v98)
							mBase = m.M
							F_sq(m, v56+int32(8), v56, v99)
							mBase = m.M
							v109 = *(*float64)(unsafe.Add(mBase, uint32(v56)))
							v110 = *(*float64)(unsafe.Add(mBase, uint32(v56)+16))
							v112 = *(*float64)(unsafe.Add(mBase, uint32(v56)+8))
							v114 = *(*float64)(unsafe.Add(mBase, uint32(v56)+24))
							v121 = base.F64_mul(v100, base.F64_sqrt(base.F64_add(base.F64_add(base.F64_add(v109, v110), v112), v114)))
						}
					}
				}
			}
			m.G0 = v56 + int32(32)
			v127 = *(*float64)(unsafe.Add(mBase, uint32(v8)+16))
			v128 = base.F64_sub(v121, v127)
			v130 = math.Float64frombits(uint64(0x7ff0000000000000))
			if base.F64_ne(base.F64_abs(v128), v130)|base.F64_eq(base.F64_abs(v121), v130)|base.F64_eq(base.F64_abs(v127), v130) == int32(0) {
				v143 = F_float_overflow_error_ext(m, int32(0))
				mBase = m.M
				v144 = m.ExcPending
				if v144 != 0 {
					return int64(0)
				} else {
					v145 = v143
					v146 = float64(0)
					if base.F64_lt(v145, v146) != 0 {
						v149 = v146
					} else {
						v149 = v145
					}
					return base.I64_reinterpret_f64(v149)
				}
			} else {
				v145 = v128
				v146 = float64(0)
				if base.F64_lt(v145, v146) != 0 {
					v149 = v146
				} else {
					v149 = v145
				}
				return base.I64_reinterpret_f64(v149)
			}
		} else {
			v43 = F_float_overflow_error_ext(m, int32(0))
			mBase = m.M
			v44 = m.ExcPending
			if v44 != 0 {
				return int64(0)
			} else {
				v45 = v43
				v54 = m.G0
				v56 = v54 - int32(32)
				m.G0 = v56
				v58 = base.F64_abs(v27)
				v59 = base.F64_abs(v45)
				v62 = base.B2i32(base.Ui64(base.I64_reinterpret_f64(v58)) < base.Ui64(base.I64_reinterpret_f64(v59)))
				if base.Ui64(base.I64_reinterpret_f64(v58)) < base.Ui64(base.I64_reinterpret_f64(v59)) {
					v63 = v58
				} else {
					v63 = v59
				}
				v64 = base.I64_reinterpret_f64(v63)
				v66 = int64(base.Ui64(v64) >> (uint(int64(52)) % 64))
				if v66 == int64(2047) {
					v121 = v63
				} else {
					if base.Ui64(base.I64_reinterpret_f64(v58)) < base.Ui64(base.I64_reinterpret_f64(v59)) {
						v69 = v59
					} else {
						v69 = v58
					}
					if v64 == int64(0) {
						v121 = v69
					} else {
						v72 = base.I64_reinterpret_f64(v69)
						v74 = int64(base.Ui64(v72) >> (uint(int64(52)) % 64))
						if v74 == int64(2047) {
							v121 = v69
						} else {
							if int32(65) <= base.I32_wrap_i64(v74)-base.I32_wrap_i64(v66) {
								v121 = base.F64_add(v58, v59)
							} else {
								if base.Ui64(int64(6908521828386340864)) <= base.Ui64(v72) {
									v85 = float64(1.90109156629516e-211)
									v98 = base.F64_mul(v69, v85)
									v99 = base.F64_mul(v63, v85)
									v100 = float64(5.260135901548374e+210)
								} else {
									if base.Ui64(int64(2580562586483294207)) < base.Ui64(v64) {
										v98 = v69
										v99 = v63
										v100 = float64(1)
									} else {
										v93 = float64(5.260135901548374e+210)
										v98 = base.F64_mul(v69, v93)
										v99 = base.F64_mul(v63, v93)
										v100 = float64(1.90109156629516e-211)
									}
								}
								F_sq(m, v56+int32(24), v56+int32(16), v98)
								mBase = m.M
								F_sq(m, v56+int32(8), v56, v99)
								mBase = m.M
								v109 = *(*float64)(unsafe.Add(mBase, uint32(v56)))
								v110 = *(*float64)(unsafe.Add(mBase, uint32(v56)+16))
								v112 = *(*float64)(unsafe.Add(mBase, uint32(v56)+8))
								v114 = *(*float64)(unsafe.Add(mBase, uint32(v56)+24))
								v121 = base.F64_mul(v100, base.F64_sqrt(base.F64_add(base.F64_add(base.F64_add(v109, v110), v112), v114)))
							}
						}
					}
				}
				m.G0 = v56 + int32(32)
				v127 = *(*float64)(unsafe.Add(mBase, uint32(v8)+16))
				v128 = base.F64_sub(v121, v127)
				v130 = math.Float64frombits(uint64(0x7ff0000000000000))
				if base.F64_ne(base.F64_abs(v128), v130)|base.F64_eq(base.F64_abs(v121), v130)|base.F64_eq(base.F64_abs(v127), v130) == int32(0) {
					v143 = F_float_overflow_error_ext(m, int32(0))
					mBase = m.M
					v144 = m.ExcPending
					if v144 != 0 {
						return int64(0)
					} else {
						v145 = v143
						v146 = float64(0)
						if base.F64_lt(v145, v146) != 0 {
							v149 = v146
						} else {
							v149 = v145
						}
						return base.I64_reinterpret_f64(v149)
					}
				} else {
					v145 = v128
					v146 = float64(0)
					if base.F64_lt(v145, v146) != 0 {
						v149 = v146
					} else {
						v149 = v145
					}
					return base.I64_reinterpret_f64(v149)
				}
			}
		}
	} else {
		v23 = F_float_overflow_error_ext(m, int32(0))
		mBase = m.M
		v26 = m.ExcPending
		if v26 != 0 {
			return int64(0)
		} else {
			v27 = v23
			v28 = *(*float64)(unsafe.Add(mBase, uint32(v6)+8))
			v29 = *(*float64)(unsafe.Add(mBase, uint32(v8)+8))
			v30 = base.F64_sub(v28, v29)
			v32 = math.Float64frombits(uint64(0x7ff0000000000000))
			if base.F64_ne(base.F64_abs(v30), v32)|base.F64_eq(base.F64_abs(v28), v32)|base.F64_eq(base.F64_abs(v29), v32) != 0 {
				v45 = v30
				v54 = m.G0
				v56 = v54 - int32(32)
				m.G0 = v56
				v58 = base.F64_abs(v27)
				v59 = base.F64_abs(v45)
				v62 = base.B2i32(base.Ui64(base.I64_reinterpret_f64(v58)) < base.Ui64(base.I64_reinterpret_f64(v59)))
				if base.Ui64(base.I64_reinterpret_f64(v58)) < base.Ui64(base.I64_reinterpret_f64(v59)) {
					v63 = v58
				} else {
					v63 = v59
				}
				v64 = base.I64_reinterpret_f64(v63)
				v66 = int64(base.Ui64(v64) >> (uint(int64(52)) % 64))
				if v66 == int64(2047) {
					v121 = v63
				} else {
					if base.Ui64(base.I64_reinterpret_f64(v58)) < base.Ui64(base.I64_reinterpret_f64(v59)) {
						v69 = v59
					} else {
						v69 = v58
					}
					if v64 == int64(0) {
						v121 = v69
					} else {
						v72 = base.I64_reinterpret_f64(v69)
						v74 = int64(base.Ui64(v72) >> (uint(int64(52)) % 64))
						if v74 == int64(2047) {
							v121 = v69
						} else {
							if int32(65) <= base.I32_wrap_i64(v74)-base.I32_wrap_i64(v66) {
								v121 = base.F64_add(v58, v59)
							} else {
								if base.Ui64(int64(6908521828386340864)) <= base.Ui64(v72) {
									v85 = float64(1.90109156629516e-211)
									v98 = base.F64_mul(v69, v85)
									v99 = base.F64_mul(v63, v85)
									v100 = float64(5.260135901548374e+210)
								} else {
									if base.Ui64(int64(2580562586483294207)) < base.Ui64(v64) {
										v98 = v69
										v99 = v63
										v100 = float64(1)
									} else {
										v93 = float64(5.260135901548374e+210)
										v98 = base.F64_mul(v69, v93)
										v99 = base.F64_mul(v63, v93)
										v100 = float64(1.90109156629516e-211)
									}
								}
								F_sq(m, v56+int32(24), v56+int32(16), v98)
								mBase = m.M
								F_sq(m, v56+int32(8), v56, v99)
								mBase = m.M
								v109 = *(*float64)(unsafe.Add(mBase, uint32(v56)))
								v110 = *(*float64)(unsafe.Add(mBase, uint32(v56)+16))
								v112 = *(*float64)(unsafe.Add(mBase, uint32(v56)+8))
								v114 = *(*float64)(unsafe.Add(mBase, uint32(v56)+24))
								v121 = base.F64_mul(v100, base.F64_sqrt(base.F64_add(base.F64_add(base.F64_add(v109, v110), v112), v114)))
							}
						}
					}
				}
				m.G0 = v56 + int32(32)
				v127 = *(*float64)(unsafe.Add(mBase, uint32(v8)+16))
				v128 = base.F64_sub(v121, v127)
				v130 = math.Float64frombits(uint64(0x7ff0000000000000))
				if base.F64_ne(base.F64_abs(v128), v130)|base.F64_eq(base.F64_abs(v121), v130)|base.F64_eq(base.F64_abs(v127), v130) == int32(0) {
					v143 = F_float_overflow_error_ext(m, int32(0))
					mBase = m.M
					v144 = m.ExcPending
					if v144 != 0 {
						return int64(0)
					} else {
						v145 = v143
						v146 = float64(0)
						if base.F64_lt(v145, v146) != 0 {
							v149 = v146
						} else {
							v149 = v145
						}
						return base.I64_reinterpret_f64(v149)
					}
				} else {
					v145 = v128
					v146 = float64(0)
					if base.F64_lt(v145, v146) != 0 {
						v149 = v146
					} else {
						v149 = v145
					}
					return base.I64_reinterpret_f64(v149)
				}
			} else {
				v43 = F_float_overflow_error_ext(m, int32(0))
				mBase = m.M
				v44 = m.ExcPending
				if v44 != 0 {
					return int64(0)
				} else {
					v45 = v43
					v54 = m.G0
					v56 = v54 - int32(32)
					m.G0 = v56
					v58 = base.F64_abs(v27)
					v59 = base.F64_abs(v45)
					v62 = base.B2i32(base.Ui64(base.I64_reinterpret_f64(v58)) < base.Ui64(base.I64_reinterpret_f64(v59)))
					if base.Ui64(base.I64_reinterpret_f64(v58)) < base.Ui64(base.I64_reinterpret_f64(v59)) {
						v63 = v58
					} else {
						v63 = v59
					}
					v64 = base.I64_reinterpret_f64(v63)
					v66 = int64(base.Ui64(v64) >> (uint(int64(52)) % 64))
					if v66 == int64(2047) {
						v121 = v63
					} else {
						if base.Ui64(base.I64_reinterpret_f64(v58)) < base.Ui64(base.I64_reinterpret_f64(v59)) {
							v69 = v59
						} else {
							v69 = v58
						}
						if v64 == int64(0) {
							v121 = v69
						} else {
							v72 = base.I64_reinterpret_f64(v69)
							v74 = int64(base.Ui64(v72) >> (uint(int64(52)) % 64))
							if v74 == int64(2047) {
								v121 = v69
							} else {
								if int32(65) <= base.I32_wrap_i64(v74)-base.I32_wrap_i64(v66) {
									v121 = base.F64_add(v58, v59)
								} else {
									if base.Ui64(int64(6908521828386340864)) <= base.Ui64(v72) {
										v85 = float64(1.90109156629516e-211)
										v98 = base.F64_mul(v69, v85)
										v99 = base.F64_mul(v63, v85)
										v100 = float64(5.260135901548374e+210)
									} else {
										if base.Ui64(int64(2580562586483294207)) < base.Ui64(v64) {
											v98 = v69
											v99 = v63
											v100 = float64(1)
										} else {
											v93 = float64(5.260135901548374e+210)
											v98 = base.F64_mul(v69, v93)
											v99 = base.F64_mul(v63, v93)
											v100 = float64(1.90109156629516e-211)
										}
									}
									F_sq(m, v56+int32(24), v56+int32(16), v98)
									mBase = m.M
									F_sq(m, v56+int32(8), v56, v99)
									mBase = m.M
									v109 = *(*float64)(unsafe.Add(mBase, uint32(v56)))
									v110 = *(*float64)(unsafe.Add(mBase, uint32(v56)+16))
									v112 = *(*float64)(unsafe.Add(mBase, uint32(v56)+8))
									v114 = *(*float64)(unsafe.Add(mBase, uint32(v56)+24))
									v121 = base.F64_mul(v100, base.F64_sqrt(base.F64_add(base.F64_add(base.F64_add(v109, v110), v112), v114)))
								}
							}
						}
					}
					m.G0 = v56 + int32(32)
					v127 = *(*float64)(unsafe.Add(mBase, uint32(v8)+16))
					v128 = base.F64_sub(v121, v127)
					v130 = math.Float64frombits(uint64(0x7ff0000000000000))
					if base.F64_ne(base.F64_abs(v128), v130)|base.F64_eq(base.F64_abs(v121), v130)|base.F64_eq(base.F64_abs(v127), v130) == int32(0) {
						v143 = F_float_overflow_error_ext(m, int32(0))
						mBase = m.M
						v144 = m.ExcPending
						if v144 != 0 {
							return int64(0)
						} else {
							v145 = v143
							v146 = float64(0)
							if base.F64_lt(v145, v146) != 0 {
								v149 = v146
							} else {
								v149 = v145
							}
							return base.I64_reinterpret_f64(v149)
						}
					} else {
						v145 = v128
						v146 = float64(0)
						if base.F64_lt(v145, v146) != 0 {
							v149 = v146
						} else {
							v149 = v145
						}
						return base.I64_reinterpret_f64(v149)
					}
				}
			}
		}
	}
}
func F_dist_polyc(m *base.Module, l0 int32) int64 {
	mBase := m.M
	_ = mBase
	var v6 int32
	_ = v6
	var v7 int32
	_ = v7
	var v10 int32
	_ = v10
	var v11 int32
	_ = v11
	var v12 float64
	_ = v12
	var v13 int32
	_ = v13
	var v14 float64
	_ = v14
	var v15 float64
	_ = v15
	var v17 float64
	_ = v17
	var v30 float64
	_ = v30
	var v31 int32
	_ = v31
	var v32 float64
	_ = v32
	var v33 float64
	_ = v33
	var v36 float64
	_ = v36
	v6 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
	v7 = F_pg_detoast_datum(m, v6)
	mBase = m.M
	v10 = m.ExcPending
	if v10 != 0 {
		return int64(0)
	} else {
		v11 = *(*int32)(unsafe.Add(mBase, uint32(l0)+40))
		v12 = F_dist_ppoly_internal(m, v11, v7)
		mBase = m.M
		v13 = m.ExcPending
		if v13 != 0 {
			return int64(0)
		} else {
			v14 = *(*float64)(unsafe.Add(mBase, uint32(v11)+16))
			v15 = base.F64_sub(v12, v14)
			v17 = math.Float64frombits(uint64(0x7ff0000000000000))
			if base.F64_ne(base.F64_abs(v15), v17)|base.F64_eq(base.F64_abs(v12), v17)|base.F64_eq(base.F64_abs(v14), v17) == int32(0) {
				v30 = F_float_overflow_error_ext(m, int32(0))
				mBase = m.M
				v31 = m.ExcPending
				if v31 != 0 {
					return int64(0)
				} else {
					v32 = v30
					v33 = float64(0)
					if base.F64_lt(v32, v33) != 0 {
						v36 = v33
					} else {
						v36 = v32
					}
					return base.I64_reinterpret_f64(v36)
				}
			} else {
				v32 = v15
				v33 = float64(0)
				if base.F64_lt(v32, v33) != 0 {
					v36 = v33
				} else {
					v36 = v32
				}
				return base.I64_reinterpret_f64(v36)
			}
		}
	}
}
func F_dist_sb(m *base.Module, l0 int32) int64 {
	mBase := m.M
	_ = mBase
	var v3 int32
	_ = v3
	var v4 int32
	_ = v4
	var v5 float64
	_ = v5
	var v8 int32
	_ = v8
	v3 = *(*int32)(unsafe.Add(mBase, uint32(l0)+40))
	v4 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
	v5 = F_box_closept_lseg(m, int32(0), v3, v4)
	mBase = m.M
	v8 = m.ExcPending
	if v8 != 0 {
		return int64(0)
	} else {
		return base.I64_reinterpret_f64(v5)
	}
}
func F_dist_sl(m *base.Module, l0 int32) int64 {
	mBase := m.M
	_ = mBase
	var v6 int32
	_ = v6
	var v7 int32
	_ = v7
	var v8 int32
	_ = v8
	var v11 int32
	_ = v11
	var v14 float64
	_ = v14
	var v15 int32
	_ = v15
	var v19 float64
	_ = v19
	var v20 int32
	_ = v20
	var v22 float64
	_ = v22
	var v26 int64
	_ = v26
	v6 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
	v7 = *(*int32)(unsafe.Add(mBase, uint32(l0)+40))
	v8 = F_lseg_interpt_line(m, int32(0), v6, v7)
	mBase = m.M
	v11 = m.ExcPending
	if v11 != 0 {
		return int64(0)
	} else {
		if v8 != 0 {
			v26 = int64(0)
			return v26
		} else {
			v14 = F_line_closept_point(m, int32(0), v7, v6)
			mBase = m.M
			v15 = m.ExcPending
			if v15 != 0 {
				return int64(0)
			} else {
				v19 = F_line_closept_point(m, int32(0), v7, v6+int32(16))
				mBase = m.M
				v20 = m.ExcPending
				if v20 != 0 {
					return int64(0)
				} else {
					if base.F64_lt(v14, v19) != 0 {
						v22 = v14
					} else {
						v22 = v19
					}
					v26 = base.I64_reinterpret_f64(v22)
					return v26
				}
			}
		}
	}
}
