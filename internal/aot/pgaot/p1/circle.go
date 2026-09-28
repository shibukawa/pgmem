package p1

import (
	base "github.com/shibukawa/pgmem/internal/aot/pgaot/base"
	"math"
	"unsafe"
)

func F_circle_add_pt(m *base.Module, l0 int32) int64 {
	mBase := m.M
	_ = mBase
	var v8 int32
	_ = v8
	var v9 int32
	_ = v9
	var v11 int32
	_ = v11
	var v14 int32
	_ = v14
	var v15 float64
	_ = v15
	var v16 float64
	_ = v16
	var v17 float64
	_ = v17
	var v19 float64
	_ = v19
	var v32 float64
	_ = v32
	var v33 int32
	_ = v33
	var v34 float64
	_ = v34
	var v35 float64
	_ = v35
	var v36 float64
	_ = v36
	var v37 float64
	_ = v37
	var v39 float64
	_ = v39
	var v50 float64
	_ = v50
	var v51 int32
	_ = v51
	var v52 float64
	_ = v52
	var v55 float64
	_ = v55
	v8 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
	v9 = *(*int32)(unsafe.Add(mBase, uint32(l0)+40))
	v11 = F_palloc(m, int32(24))
	mBase = m.M
	v14 = m.ExcPending
	if v14 != 0 {
		return int64(0)
	} else {
		v15 = *(*float64)(unsafe.Add(mBase, uint32(v8)))
		v16 = *(*float64)(unsafe.Add(mBase, uint32(v9)))
		v17 = base.F64_add(v15, v16)
		v19 = math.Float64frombits(uint64(0x7ff0000000000000))
		if base.F64_ne(base.F64_abs(v17), v19)|base.F64_eq(base.F64_abs(v15), v19)|base.F64_eq(base.F64_abs(v16), v19) == int32(0) {
			v32 = F_float_overflow_error_ext(m, int32(0))
			mBase = m.M
			v33 = m.ExcPending
			if v33 != 0 {
				return int64(0)
			} else {
				v34 = v32
				v35 = *(*float64)(unsafe.Add(mBase, uint32(v8)+8))
				v36 = *(*float64)(unsafe.Add(mBase, uint32(v9)+8))
				v37 = base.F64_add(v35, v36)
				v39 = math.Float64frombits(uint64(0x7ff0000000000000))
				if base.F64_ne(base.F64_abs(v37), v39)|base.F64_eq(base.F64_abs(v35), v39)|base.F64_eq(base.F64_abs(v36), v39) != 0 {
					v52 = v37
					*(*float64)(unsafe.Add(mBase, uint32(v11)+8)) = v52
					*(*float64)(unsafe.Add(mBase, uint32(v11))) = v34
					v55 = *(*float64)(unsafe.Add(mBase, uint32(v8)+16))
					*(*float64)(unsafe.Add(mBase, uint32(v11)+16)) = v55
					return base.I64_extend_i32_u(v11)
				} else {
					v50 = F_float_overflow_error_ext(m, int32(0))
					mBase = m.M
					v51 = m.ExcPending
					if v51 != 0 {
						return int64(0)
					} else {
						v52 = v50
						*(*float64)(unsafe.Add(mBase, uint32(v11)+8)) = v52
						*(*float64)(unsafe.Add(mBase, uint32(v11))) = v34
						v55 = *(*float64)(unsafe.Add(mBase, uint32(v8)+16))
						*(*float64)(unsafe.Add(mBase, uint32(v11)+16)) = v55
						return base.I64_extend_i32_u(v11)
					}
				}
			}
		} else {
			v34 = v17
			v35 = *(*float64)(unsafe.Add(mBase, uint32(v8)+8))
			v36 = *(*float64)(unsafe.Add(mBase, uint32(v9)+8))
			v37 = base.F64_add(v35, v36)
			v39 = math.Float64frombits(uint64(0x7ff0000000000000))
			if base.F64_ne(base.F64_abs(v37), v39)|base.F64_eq(base.F64_abs(v35), v39)|base.F64_eq(base.F64_abs(v36), v39) != 0 {
				v52 = v37
				*(*float64)(unsafe.Add(mBase, uint32(v11)+8)) = v52
				*(*float64)(unsafe.Add(mBase, uint32(v11))) = v34
				v55 = *(*float64)(unsafe.Add(mBase, uint32(v8)+16))
				*(*float64)(unsafe.Add(mBase, uint32(v11)+16)) = v55
				return base.I64_extend_i32_u(v11)
			} else {
				v50 = F_float_overflow_error_ext(m, int32(0))
				mBase = m.M
				v51 = m.ExcPending
				if v51 != 0 {
					return int64(0)
				} else {
					v52 = v50
					*(*float64)(unsafe.Add(mBase, uint32(v11)+8)) = v52
					*(*float64)(unsafe.Add(mBase, uint32(v11))) = v34
					v55 = *(*float64)(unsafe.Add(mBase, uint32(v8)+16))
					*(*float64)(unsafe.Add(mBase, uint32(v11)+16)) = v55
					return base.I64_extend_i32_u(v11)
				}
			}
		}
	}
}
func F_circle_contain(m *base.Module, l0 int32) int64 {
	mBase := m.M
	_ = mBase
	var v8 int32
	_ = v8
	var v9 float64
	_ = v9
	var v10 int32
	_ = v10
	var v11 float64
	_ = v11
	var v12 float64
	_ = v12
	var v14 float64
	_ = v14
	var v27 float64
	_ = v27
	var v30 int32
	_ = v30
	var v31 float64
	_ = v31
	var v32 float64
	_ = v32
	var v33 float64
	_ = v33
	var v34 float64
	_ = v34
	var v36 float64
	_ = v36
	var v49 float64
	_ = v49
	var v50 int32
	_ = v50
	var v51 float64
	_ = v51
	var v52 float64
	_ = v52
	var v53 float64
	_ = v53
	var v54 float64
	_ = v54
	var v56 float64
	_ = v56
	var v69 float64
	_ = v69
	var v70 int32
	_ = v70
	var v71 float64
	_ = v71
	var v80 int32
	_ = v80
	var v82 int32
	_ = v82
	var v84 float64
	_ = v84
	var v85 float64
	_ = v85
	var v88 int32
	_ = v88
	var v89 float64
	_ = v89
	var v90 int64
	_ = v90
	var v92 int64
	_ = v92
	var v95 float64
	_ = v95
	var v98 int64
	_ = v98
	var v100 int64
	_ = v100
	var v111 float64
	_ = v111
	var v119 float64
	_ = v119
	var v124 float64
	_ = v124
	var v125 float64
	_ = v125
	var v126 float64
	_ = v126
	var v135 float64
	_ = v135
	var v136 float64
	_ = v136
	var v138 float64
	_ = v138
	var v140 float64
	_ = v140
	var v147 float64
	_ = v147
	v8 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
	v9 = *(*float64)(unsafe.Add(mBase, uint32(v8)))
	v10 = *(*int32)(unsafe.Add(mBase, uint32(l0)+40))
	v11 = *(*float64)(unsafe.Add(mBase, uint32(v10)))
	v12 = base.F64_sub(v9, v11)
	v14 = math.Float64frombits(uint64(0x7ff0000000000000))
	if base.F64_ne(base.F64_abs(v12), v14)|base.F64_eq(base.F64_abs(v9), v14)|base.F64_eq(base.F64_abs(v11), v14) == int32(0) {
		v27 = F_float_overflow_error_ext(m, int32(0))
		mBase = m.M
		v30 = m.ExcPending
		if v30 != 0 {
			return int64(0)
		} else {
			v31 = v27
			v32 = *(*float64)(unsafe.Add(mBase, uint32(v8)+8))
			v33 = *(*float64)(unsafe.Add(mBase, uint32(v10)+8))
			v34 = base.F64_sub(v32, v33)
			v36 = math.Float64frombits(uint64(0x7ff0000000000000))
			if base.F64_ne(base.F64_abs(v34), v36)|base.F64_eq(base.F64_abs(v32), v36)|base.F64_eq(base.F64_abs(v33), v36) == int32(0) {
				v49 = F_float_overflow_error_ext(m, int32(0))
				mBase = m.M
				v50 = m.ExcPending
				if v50 != 0 {
					return int64(0)
				} else {
					v51 = v49
					v52 = *(*float64)(unsafe.Add(mBase, uint32(v8)+16))
					v53 = *(*float64)(unsafe.Add(mBase, uint32(v10)+16))
					v54 = base.F64_sub(v52, v53)
					v56 = math.Float64frombits(uint64(0x7ff0000000000000))
					if base.F64_ne(base.F64_abs(v54), v56)|base.F64_eq(base.F64_abs(v52), v56)|base.F64_eq(base.F64_abs(v53), v56) == int32(0) {
						v69 = F_float_overflow_error_ext(m, int32(0))
						mBase = m.M
						v70 = m.ExcPending
						if v70 != 0 {
							return int64(0)
						} else {
							v71 = v69
							v80 = m.G0
							v82 = v80 - int32(32)
							m.G0 = v82
							v84 = base.F64_abs(v31)
							v85 = base.F64_abs(v51)
							v88 = base.B2i32(base.Ui64(base.I64_reinterpret_f64(v84)) < base.Ui64(base.I64_reinterpret_f64(v85)))
							if base.Ui64(base.I64_reinterpret_f64(v84)) < base.Ui64(base.I64_reinterpret_f64(v85)) {
								v89 = v84
							} else {
								v89 = v85
							}
							v90 = base.I64_reinterpret_f64(v89)
							v92 = int64(base.Ui64(v90) >> (uint(int64(52)) % 64))
							if v92 == int64(2047) {
								v147 = v89
							} else {
								if base.Ui64(base.I64_reinterpret_f64(v84)) < base.Ui64(base.I64_reinterpret_f64(v85)) {
									v95 = v85
								} else {
									v95 = v84
								}
								if v90 == int64(0) {
									v147 = v95
								} else {
									v98 = base.I64_reinterpret_f64(v95)
									v100 = int64(base.Ui64(v98) >> (uint(int64(52)) % 64))
									if v100 == int64(2047) {
										v147 = v95
									} else {
										if int32(65) <= base.I32_wrap_i64(v100)-base.I32_wrap_i64(v92) {
											v147 = base.F64_add(v84, v85)
										} else {
											if base.Ui64(int64(6908521828386340864)) <= base.Ui64(v98) {
												v111 = float64(1.90109156629516e-211)
												v124 = base.F64_mul(v95, v111)
												v125 = base.F64_mul(v89, v111)
												v126 = float64(5.260135901548374e+210)
											} else {
												if base.Ui64(int64(2580562586483294207)) < base.Ui64(v90) {
													v124 = v95
													v125 = v89
													v126 = float64(1)
												} else {
													v119 = float64(5.260135901548374e+210)
													v124 = base.F64_mul(v95, v119)
													v125 = base.F64_mul(v89, v119)
													v126 = float64(1.90109156629516e-211)
												}
											}
											F_sq(m, v82+int32(24), v82+int32(16), v124)
											mBase = m.M
											F_sq(m, v82+int32(8), v82, v125)
											mBase = m.M
											v135 = *(*float64)(unsafe.Add(mBase, uint32(v82)))
											v136 = *(*float64)(unsafe.Add(mBase, uint32(v82)+16))
											v138 = *(*float64)(unsafe.Add(mBase, uint32(v82)+8))
											v140 = *(*float64)(unsafe.Add(mBase, uint32(v82)+24))
											v147 = base.F64_mul(v126, base.F64_sqrt(base.F64_add(base.F64_add(base.F64_add(v135, v136), v138), v140)))
										}
									}
								}
							}
							m.G0 = v82 + int32(32)
							return base.I64_extend_i32_u(base.F64_le(v147, base.F64_add(v71, float64(1e-06))))
						}
					} else {
						v71 = v54
						v80 = m.G0
						v82 = v80 - int32(32)
						m.G0 = v82
						v84 = base.F64_abs(v31)
						v85 = base.F64_abs(v51)
						v88 = base.B2i32(base.Ui64(base.I64_reinterpret_f64(v84)) < base.Ui64(base.I64_reinterpret_f64(v85)))
						if base.Ui64(base.I64_reinterpret_f64(v84)) < base.Ui64(base.I64_reinterpret_f64(v85)) {
							v89 = v84
						} else {
							v89 = v85
						}
						v90 = base.I64_reinterpret_f64(v89)
						v92 = int64(base.Ui64(v90) >> (uint(int64(52)) % 64))
						if v92 == int64(2047) {
							v147 = v89
						} else {
							if base.Ui64(base.I64_reinterpret_f64(v84)) < base.Ui64(base.I64_reinterpret_f64(v85)) {
								v95 = v85
							} else {
								v95 = v84
							}
							if v90 == int64(0) {
								v147 = v95
							} else {
								v98 = base.I64_reinterpret_f64(v95)
								v100 = int64(base.Ui64(v98) >> (uint(int64(52)) % 64))
								if v100 == int64(2047) {
									v147 = v95
								} else {
									if int32(65) <= base.I32_wrap_i64(v100)-base.I32_wrap_i64(v92) {
										v147 = base.F64_add(v84, v85)
									} else {
										if base.Ui64(int64(6908521828386340864)) <= base.Ui64(v98) {
											v111 = float64(1.90109156629516e-211)
											v124 = base.F64_mul(v95, v111)
											v125 = base.F64_mul(v89, v111)
											v126 = float64(5.260135901548374e+210)
										} else {
											if base.Ui64(int64(2580562586483294207)) < base.Ui64(v90) {
												v124 = v95
												v125 = v89
												v126 = float64(1)
											} else {
												v119 = float64(5.260135901548374e+210)
												v124 = base.F64_mul(v95, v119)
												v125 = base.F64_mul(v89, v119)
												v126 = float64(1.90109156629516e-211)
											}
										}
										F_sq(m, v82+int32(24), v82+int32(16), v124)
										mBase = m.M
										F_sq(m, v82+int32(8), v82, v125)
										mBase = m.M
										v135 = *(*float64)(unsafe.Add(mBase, uint32(v82)))
										v136 = *(*float64)(unsafe.Add(mBase, uint32(v82)+16))
										v138 = *(*float64)(unsafe.Add(mBase, uint32(v82)+8))
										v140 = *(*float64)(unsafe.Add(mBase, uint32(v82)+24))
										v147 = base.F64_mul(v126, base.F64_sqrt(base.F64_add(base.F64_add(base.F64_add(v135, v136), v138), v140)))
									}
								}
							}
						}
						m.G0 = v82 + int32(32)
						return base.I64_extend_i32_u(base.F64_le(v147, base.F64_add(v71, float64(1e-06))))
					}
				}
			} else {
				v51 = v34
				v52 = *(*float64)(unsafe.Add(mBase, uint32(v8)+16))
				v53 = *(*float64)(unsafe.Add(mBase, uint32(v10)+16))
				v54 = base.F64_sub(v52, v53)
				v56 = math.Float64frombits(uint64(0x7ff0000000000000))
				if base.F64_ne(base.F64_abs(v54), v56)|base.F64_eq(base.F64_abs(v52), v56)|base.F64_eq(base.F64_abs(v53), v56) == int32(0) {
					v69 = F_float_overflow_error_ext(m, int32(0))
					mBase = m.M
					v70 = m.ExcPending
					if v70 != 0 {
						return int64(0)
					} else {
						v71 = v69
						v80 = m.G0
						v82 = v80 - int32(32)
						m.G0 = v82
						v84 = base.F64_abs(v31)
						v85 = base.F64_abs(v51)
						v88 = base.B2i32(base.Ui64(base.I64_reinterpret_f64(v84)) < base.Ui64(base.I64_reinterpret_f64(v85)))
						if base.Ui64(base.I64_reinterpret_f64(v84)) < base.Ui64(base.I64_reinterpret_f64(v85)) {
							v89 = v84
						} else {
							v89 = v85
						}
						v90 = base.I64_reinterpret_f64(v89)
						v92 = int64(base.Ui64(v90) >> (uint(int64(52)) % 64))
						if v92 == int64(2047) {
							v147 = v89
						} else {
							if base.Ui64(base.I64_reinterpret_f64(v84)) < base.Ui64(base.I64_reinterpret_f64(v85)) {
								v95 = v85
							} else {
								v95 = v84
							}
							if v90 == int64(0) {
								v147 = v95
							} else {
								v98 = base.I64_reinterpret_f64(v95)
								v100 = int64(base.Ui64(v98) >> (uint(int64(52)) % 64))
								if v100 == int64(2047) {
									v147 = v95
								} else {
									if int32(65) <= base.I32_wrap_i64(v100)-base.I32_wrap_i64(v92) {
										v147 = base.F64_add(v84, v85)
									} else {
										if base.Ui64(int64(6908521828386340864)) <= base.Ui64(v98) {
											v111 = float64(1.90109156629516e-211)
											v124 = base.F64_mul(v95, v111)
											v125 = base.F64_mul(v89, v111)
											v126 = float64(5.260135901548374e+210)
										} else {
											if base.Ui64(int64(2580562586483294207)) < base.Ui64(v90) {
												v124 = v95
												v125 = v89
												v126 = float64(1)
											} else {
												v119 = float64(5.260135901548374e+210)
												v124 = base.F64_mul(v95, v119)
												v125 = base.F64_mul(v89, v119)
												v126 = float64(1.90109156629516e-211)
											}
										}
										F_sq(m, v82+int32(24), v82+int32(16), v124)
										mBase = m.M
										F_sq(m, v82+int32(8), v82, v125)
										mBase = m.M
										v135 = *(*float64)(unsafe.Add(mBase, uint32(v82)))
										v136 = *(*float64)(unsafe.Add(mBase, uint32(v82)+16))
										v138 = *(*float64)(unsafe.Add(mBase, uint32(v82)+8))
										v140 = *(*float64)(unsafe.Add(mBase, uint32(v82)+24))
										v147 = base.F64_mul(v126, base.F64_sqrt(base.F64_add(base.F64_add(base.F64_add(v135, v136), v138), v140)))
									}
								}
							}
						}
						m.G0 = v82 + int32(32)
						return base.I64_extend_i32_u(base.F64_le(v147, base.F64_add(v71, float64(1e-06))))
					}
				} else {
					v71 = v54
					v80 = m.G0
					v82 = v80 - int32(32)
					m.G0 = v82
					v84 = base.F64_abs(v31)
					v85 = base.F64_abs(v51)
					v88 = base.B2i32(base.Ui64(base.I64_reinterpret_f64(v84)) < base.Ui64(base.I64_reinterpret_f64(v85)))
					if base.Ui64(base.I64_reinterpret_f64(v84)) < base.Ui64(base.I64_reinterpret_f64(v85)) {
						v89 = v84
					} else {
						v89 = v85
					}
					v90 = base.I64_reinterpret_f64(v89)
					v92 = int64(base.Ui64(v90) >> (uint(int64(52)) % 64))
					if v92 == int64(2047) {
						v147 = v89
					} else {
						if base.Ui64(base.I64_reinterpret_f64(v84)) < base.Ui64(base.I64_reinterpret_f64(v85)) {
							v95 = v85
						} else {
							v95 = v84
						}
						if v90 == int64(0) {
							v147 = v95
						} else {
							v98 = base.I64_reinterpret_f64(v95)
							v100 = int64(base.Ui64(v98) >> (uint(int64(52)) % 64))
							if v100 == int64(2047) {
								v147 = v95
							} else {
								if int32(65) <= base.I32_wrap_i64(v100)-base.I32_wrap_i64(v92) {
									v147 = base.F64_add(v84, v85)
								} else {
									if base.Ui64(int64(6908521828386340864)) <= base.Ui64(v98) {
										v111 = float64(1.90109156629516e-211)
										v124 = base.F64_mul(v95, v111)
										v125 = base.F64_mul(v89, v111)
										v126 = float64(5.260135901548374e+210)
									} else {
										if base.Ui64(int64(2580562586483294207)) < base.Ui64(v90) {
											v124 = v95
											v125 = v89
											v126 = float64(1)
										} else {
											v119 = float64(5.260135901548374e+210)
											v124 = base.F64_mul(v95, v119)
											v125 = base.F64_mul(v89, v119)
											v126 = float64(1.90109156629516e-211)
										}
									}
									F_sq(m, v82+int32(24), v82+int32(16), v124)
									mBase = m.M
									F_sq(m, v82+int32(8), v82, v125)
									mBase = m.M
									v135 = *(*float64)(unsafe.Add(mBase, uint32(v82)))
									v136 = *(*float64)(unsafe.Add(mBase, uint32(v82)+16))
									v138 = *(*float64)(unsafe.Add(mBase, uint32(v82)+8))
									v140 = *(*float64)(unsafe.Add(mBase, uint32(v82)+24))
									v147 = base.F64_mul(v126, base.F64_sqrt(base.F64_add(base.F64_add(base.F64_add(v135, v136), v138), v140)))
								}
							}
						}
					}
					m.G0 = v82 + int32(32)
					return base.I64_extend_i32_u(base.F64_le(v147, base.F64_add(v71, float64(1e-06))))
				}
			}
		}
	} else {
		v31 = v12
		v32 = *(*float64)(unsafe.Add(mBase, uint32(v8)+8))
		v33 = *(*float64)(unsafe.Add(mBase, uint32(v10)+8))
		v34 = base.F64_sub(v32, v33)
		v36 = math.Float64frombits(uint64(0x7ff0000000000000))
		if base.F64_ne(base.F64_abs(v34), v36)|base.F64_eq(base.F64_abs(v32), v36)|base.F64_eq(base.F64_abs(v33), v36) == int32(0) {
			v49 = F_float_overflow_error_ext(m, int32(0))
			mBase = m.M
			v50 = m.ExcPending
			if v50 != 0 {
				return int64(0)
			} else {
				v51 = v49
				v52 = *(*float64)(unsafe.Add(mBase, uint32(v8)+16))
				v53 = *(*float64)(unsafe.Add(mBase, uint32(v10)+16))
				v54 = base.F64_sub(v52, v53)
				v56 = math.Float64frombits(uint64(0x7ff0000000000000))
				if base.F64_ne(base.F64_abs(v54), v56)|base.F64_eq(base.F64_abs(v52), v56)|base.F64_eq(base.F64_abs(v53), v56) == int32(0) {
					v69 = F_float_overflow_error_ext(m, int32(0))
					mBase = m.M
					v70 = m.ExcPending
					if v70 != 0 {
						return int64(0)
					} else {
						v71 = v69
						v80 = m.G0
						v82 = v80 - int32(32)
						m.G0 = v82
						v84 = base.F64_abs(v31)
						v85 = base.F64_abs(v51)
						v88 = base.B2i32(base.Ui64(base.I64_reinterpret_f64(v84)) < base.Ui64(base.I64_reinterpret_f64(v85)))
						if base.Ui64(base.I64_reinterpret_f64(v84)) < base.Ui64(base.I64_reinterpret_f64(v85)) {
							v89 = v84
						} else {
							v89 = v85
						}
						v90 = base.I64_reinterpret_f64(v89)
						v92 = int64(base.Ui64(v90) >> (uint(int64(52)) % 64))
						if v92 == int64(2047) {
							v147 = v89
						} else {
							if base.Ui64(base.I64_reinterpret_f64(v84)) < base.Ui64(base.I64_reinterpret_f64(v85)) {
								v95 = v85
							} else {
								v95 = v84
							}
							if v90 == int64(0) {
								v147 = v95
							} else {
								v98 = base.I64_reinterpret_f64(v95)
								v100 = int64(base.Ui64(v98) >> (uint(int64(52)) % 64))
								if v100 == int64(2047) {
									v147 = v95
								} else {
									if int32(65) <= base.I32_wrap_i64(v100)-base.I32_wrap_i64(v92) {
										v147 = base.F64_add(v84, v85)
									} else {
										if base.Ui64(int64(6908521828386340864)) <= base.Ui64(v98) {
											v111 = float64(1.90109156629516e-211)
											v124 = base.F64_mul(v95, v111)
											v125 = base.F64_mul(v89, v111)
											v126 = float64(5.260135901548374e+210)
										} else {
											if base.Ui64(int64(2580562586483294207)) < base.Ui64(v90) {
												v124 = v95
												v125 = v89
												v126 = float64(1)
											} else {
												v119 = float64(5.260135901548374e+210)
												v124 = base.F64_mul(v95, v119)
												v125 = base.F64_mul(v89, v119)
												v126 = float64(1.90109156629516e-211)
											}
										}
										F_sq(m, v82+int32(24), v82+int32(16), v124)
										mBase = m.M
										F_sq(m, v82+int32(8), v82, v125)
										mBase = m.M
										v135 = *(*float64)(unsafe.Add(mBase, uint32(v82)))
										v136 = *(*float64)(unsafe.Add(mBase, uint32(v82)+16))
										v138 = *(*float64)(unsafe.Add(mBase, uint32(v82)+8))
										v140 = *(*float64)(unsafe.Add(mBase, uint32(v82)+24))
										v147 = base.F64_mul(v126, base.F64_sqrt(base.F64_add(base.F64_add(base.F64_add(v135, v136), v138), v140)))
									}
								}
							}
						}
						m.G0 = v82 + int32(32)
						return base.I64_extend_i32_u(base.F64_le(v147, base.F64_add(v71, float64(1e-06))))
					}
				} else {
					v71 = v54
					v80 = m.G0
					v82 = v80 - int32(32)
					m.G0 = v82
					v84 = base.F64_abs(v31)
					v85 = base.F64_abs(v51)
					v88 = base.B2i32(base.Ui64(base.I64_reinterpret_f64(v84)) < base.Ui64(base.I64_reinterpret_f64(v85)))
					if base.Ui64(base.I64_reinterpret_f64(v84)) < base.Ui64(base.I64_reinterpret_f64(v85)) {
						v89 = v84
					} else {
						v89 = v85
					}
					v90 = base.I64_reinterpret_f64(v89)
					v92 = int64(base.Ui64(v90) >> (uint(int64(52)) % 64))
					if v92 == int64(2047) {
						v147 = v89
					} else {
						if base.Ui64(base.I64_reinterpret_f64(v84)) < base.Ui64(base.I64_reinterpret_f64(v85)) {
							v95 = v85
						} else {
							v95 = v84
						}
						if v90 == int64(0) {
							v147 = v95
						} else {
							v98 = base.I64_reinterpret_f64(v95)
							v100 = int64(base.Ui64(v98) >> (uint(int64(52)) % 64))
							if v100 == int64(2047) {
								v147 = v95
							} else {
								if int32(65) <= base.I32_wrap_i64(v100)-base.I32_wrap_i64(v92) {
									v147 = base.F64_add(v84, v85)
								} else {
									if base.Ui64(int64(6908521828386340864)) <= base.Ui64(v98) {
										v111 = float64(1.90109156629516e-211)
										v124 = base.F64_mul(v95, v111)
										v125 = base.F64_mul(v89, v111)
										v126 = float64(5.260135901548374e+210)
									} else {
										if base.Ui64(int64(2580562586483294207)) < base.Ui64(v90) {
											v124 = v95
											v125 = v89
											v126 = float64(1)
										} else {
											v119 = float64(5.260135901548374e+210)
											v124 = base.F64_mul(v95, v119)
											v125 = base.F64_mul(v89, v119)
											v126 = float64(1.90109156629516e-211)
										}
									}
									F_sq(m, v82+int32(24), v82+int32(16), v124)
									mBase = m.M
									F_sq(m, v82+int32(8), v82, v125)
									mBase = m.M
									v135 = *(*float64)(unsafe.Add(mBase, uint32(v82)))
									v136 = *(*float64)(unsafe.Add(mBase, uint32(v82)+16))
									v138 = *(*float64)(unsafe.Add(mBase, uint32(v82)+8))
									v140 = *(*float64)(unsafe.Add(mBase, uint32(v82)+24))
									v147 = base.F64_mul(v126, base.F64_sqrt(base.F64_add(base.F64_add(base.F64_add(v135, v136), v138), v140)))
								}
							}
						}
					}
					m.G0 = v82 + int32(32)
					return base.I64_extend_i32_u(base.F64_le(v147, base.F64_add(v71, float64(1e-06))))
				}
			}
		} else {
			v51 = v34
			v52 = *(*float64)(unsafe.Add(mBase, uint32(v8)+16))
			v53 = *(*float64)(unsafe.Add(mBase, uint32(v10)+16))
			v54 = base.F64_sub(v52, v53)
			v56 = math.Float64frombits(uint64(0x7ff0000000000000))
			if base.F64_ne(base.F64_abs(v54), v56)|base.F64_eq(base.F64_abs(v52), v56)|base.F64_eq(base.F64_abs(v53), v56) == int32(0) {
				v69 = F_float_overflow_error_ext(m, int32(0))
				mBase = m.M
				v70 = m.ExcPending
				if v70 != 0 {
					return int64(0)
				} else {
					v71 = v69
					v80 = m.G0
					v82 = v80 - int32(32)
					m.G0 = v82
					v84 = base.F64_abs(v31)
					v85 = base.F64_abs(v51)
					v88 = base.B2i32(base.Ui64(base.I64_reinterpret_f64(v84)) < base.Ui64(base.I64_reinterpret_f64(v85)))
					if base.Ui64(base.I64_reinterpret_f64(v84)) < base.Ui64(base.I64_reinterpret_f64(v85)) {
						v89 = v84
					} else {
						v89 = v85
					}
					v90 = base.I64_reinterpret_f64(v89)
					v92 = int64(base.Ui64(v90) >> (uint(int64(52)) % 64))
					if v92 == int64(2047) {
						v147 = v89
					} else {
						if base.Ui64(base.I64_reinterpret_f64(v84)) < base.Ui64(base.I64_reinterpret_f64(v85)) {
							v95 = v85
						} else {
							v95 = v84
						}
						if v90 == int64(0) {
							v147 = v95
						} else {
							v98 = base.I64_reinterpret_f64(v95)
							v100 = int64(base.Ui64(v98) >> (uint(int64(52)) % 64))
							if v100 == int64(2047) {
								v147 = v95
							} else {
								if int32(65) <= base.I32_wrap_i64(v100)-base.I32_wrap_i64(v92) {
									v147 = base.F64_add(v84, v85)
								} else {
									if base.Ui64(int64(6908521828386340864)) <= base.Ui64(v98) {
										v111 = float64(1.90109156629516e-211)
										v124 = base.F64_mul(v95, v111)
										v125 = base.F64_mul(v89, v111)
										v126 = float64(5.260135901548374e+210)
									} else {
										if base.Ui64(int64(2580562586483294207)) < base.Ui64(v90) {
											v124 = v95
											v125 = v89
											v126 = float64(1)
										} else {
											v119 = float64(5.260135901548374e+210)
											v124 = base.F64_mul(v95, v119)
											v125 = base.F64_mul(v89, v119)
											v126 = float64(1.90109156629516e-211)
										}
									}
									F_sq(m, v82+int32(24), v82+int32(16), v124)
									mBase = m.M
									F_sq(m, v82+int32(8), v82, v125)
									mBase = m.M
									v135 = *(*float64)(unsafe.Add(mBase, uint32(v82)))
									v136 = *(*float64)(unsafe.Add(mBase, uint32(v82)+16))
									v138 = *(*float64)(unsafe.Add(mBase, uint32(v82)+8))
									v140 = *(*float64)(unsafe.Add(mBase, uint32(v82)+24))
									v147 = base.F64_mul(v126, base.F64_sqrt(base.F64_add(base.F64_add(base.F64_add(v135, v136), v138), v140)))
								}
							}
						}
					}
					m.G0 = v82 + int32(32)
					return base.I64_extend_i32_u(base.F64_le(v147, base.F64_add(v71, float64(1e-06))))
				}
			} else {
				v71 = v54
				v80 = m.G0
				v82 = v80 - int32(32)
				m.G0 = v82
				v84 = base.F64_abs(v31)
				v85 = base.F64_abs(v51)
				v88 = base.B2i32(base.Ui64(base.I64_reinterpret_f64(v84)) < base.Ui64(base.I64_reinterpret_f64(v85)))
				if base.Ui64(base.I64_reinterpret_f64(v84)) < base.Ui64(base.I64_reinterpret_f64(v85)) {
					v89 = v84
				} else {
					v89 = v85
				}
				v90 = base.I64_reinterpret_f64(v89)
				v92 = int64(base.Ui64(v90) >> (uint(int64(52)) % 64))
				if v92 == int64(2047) {
					v147 = v89
				} else {
					if base.Ui64(base.I64_reinterpret_f64(v84)) < base.Ui64(base.I64_reinterpret_f64(v85)) {
						v95 = v85
					} else {
						v95 = v84
					}
					if v90 == int64(0) {
						v147 = v95
					} else {
						v98 = base.I64_reinterpret_f64(v95)
						v100 = int64(base.Ui64(v98) >> (uint(int64(52)) % 64))
						if v100 == int64(2047) {
							v147 = v95
						} else {
							if int32(65) <= base.I32_wrap_i64(v100)-base.I32_wrap_i64(v92) {
								v147 = base.F64_add(v84, v85)
							} else {
								if base.Ui64(int64(6908521828386340864)) <= base.Ui64(v98) {
									v111 = float64(1.90109156629516e-211)
									v124 = base.F64_mul(v95, v111)
									v125 = base.F64_mul(v89, v111)
									v126 = float64(5.260135901548374e+210)
								} else {
									if base.Ui64(int64(2580562586483294207)) < base.Ui64(v90) {
										v124 = v95
										v125 = v89
										v126 = float64(1)
									} else {
										v119 = float64(5.260135901548374e+210)
										v124 = base.F64_mul(v95, v119)
										v125 = base.F64_mul(v89, v119)
										v126 = float64(1.90109156629516e-211)
									}
								}
								F_sq(m, v82+int32(24), v82+int32(16), v124)
								mBase = m.M
								F_sq(m, v82+int32(8), v82, v125)
								mBase = m.M
								v135 = *(*float64)(unsafe.Add(mBase, uint32(v82)))
								v136 = *(*float64)(unsafe.Add(mBase, uint32(v82)+16))
								v138 = *(*float64)(unsafe.Add(mBase, uint32(v82)+8))
								v140 = *(*float64)(unsafe.Add(mBase, uint32(v82)+24))
								v147 = base.F64_mul(v126, base.F64_sqrt(base.F64_add(base.F64_add(base.F64_add(v135, v136), v138), v140)))
							}
						}
					}
				}
				m.G0 = v82 + int32(32)
				return base.I64_extend_i32_u(base.F64_le(v147, base.F64_add(v71, float64(1e-06))))
			}
		}
	}
}
func F_circle_div_pt(m *base.Module, l0 int32) int64 {
	mBase := m.M
	_ = mBase
	var v7 int32
	_ = v7
	var v8 int32
	_ = v8
	var v10 int32
	_ = v10
	var v13 int32
	_ = v13
	var v15 int32
	_ = v15
	var v16 float64
	_ = v16
	var v22 float64
	_ = v22
	var v23 float64
	_ = v23
	var v32 int32
	_ = v32
	var v34 int32
	_ = v34
	var v36 float64
	_ = v36
	var v37 float64
	_ = v37
	var v40 int32
	_ = v40
	var v41 float64
	_ = v41
	var v42 int64
	_ = v42
	var v44 int64
	_ = v44
	var v47 float64
	_ = v47
	var v50 int64
	_ = v50
	var v52 int64
	_ = v52
	var v63 float64
	_ = v63
	var v71 float64
	_ = v71
	var v76 float64
	_ = v76
	var v77 float64
	_ = v77
	var v78 float64
	_ = v78
	var v87 float64
	_ = v87
	var v88 float64
	_ = v88
	var v90 float64
	_ = v90
	var v92 float64
	_ = v92
	var v99 float64
	_ = v99
	var v111 float64
	_ = v111
	var v112 int32
	_ = v112
	var v114 float64
	_ = v114
	var v116 float64
	_ = v116
	var v124 float64
	_ = v124
	var v125 int32
	_ = v125
	var v126 float64
	_ = v126
	var v136 float64
	_ = v136
	var v137 int32
	_ = v137
	var v138 float64
	_ = v138
	v7 = *(*int32)(unsafe.Add(mBase, uint32(l0)+40))
	v8 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
	v10 = F_palloc(m, int32(24))
	mBase = m.M
	v13 = m.ExcPending
	if v13 != 0 {
		return int64(0)
	} else {
		F_point_div_point(m, v10, v8, v7)
		mBase = m.M
		v15 = m.ExcPending
		if v15 != 0 {
			return int64(0)
		} else {
			v16 = *(*float64)(unsafe.Add(mBase, uint32(v8)+16))
			v22 = *(*float64)(unsafe.Add(mBase, uint32(v7)))
			v23 = *(*float64)(unsafe.Add(mBase, uint32(v7)+8))
			v32 = m.G0
			v34 = v32 - int32(32)
			m.G0 = v34
			v36 = base.F64_abs(v22)
			v37 = base.F64_abs(v23)
			v40 = base.B2i32(base.Ui64(base.I64_reinterpret_f64(v36)) < base.Ui64(base.I64_reinterpret_f64(v37)))
			if base.Ui64(base.I64_reinterpret_f64(v36)) < base.Ui64(base.I64_reinterpret_f64(v37)) {
				v41 = v36
			} else {
				v41 = v37
			}
			v42 = base.I64_reinterpret_f64(v41)
			v44 = int64(base.Ui64(v42) >> (uint(int64(52)) % 64))
			if v44 == int64(2047) {
				v99 = v41
			} else {
				if base.Ui64(base.I64_reinterpret_f64(v36)) < base.Ui64(base.I64_reinterpret_f64(v37)) {
					v47 = v37
				} else {
					v47 = v36
				}
				if v42 == int64(0) {
					v99 = v47
				} else {
					v50 = base.I64_reinterpret_f64(v47)
					v52 = int64(base.Ui64(v50) >> (uint(int64(52)) % 64))
					if v52 == int64(2047) {
						v99 = v47
					} else {
						if int32(65) <= base.I32_wrap_i64(v52)-base.I32_wrap_i64(v44) {
							v99 = base.F64_add(v36, v37)
						} else {
							if base.Ui64(int64(6908521828386340864)) <= base.Ui64(v50) {
								v63 = float64(1.90109156629516e-211)
								v76 = base.F64_mul(v47, v63)
								v77 = base.F64_mul(v41, v63)
								v78 = float64(5.260135901548374e+210)
							} else {
								if base.Ui64(int64(2580562586483294207)) < base.Ui64(v42) {
									v76 = v47
									v77 = v41
									v78 = float64(1)
								} else {
									v71 = float64(5.260135901548374e+210)
									v76 = base.F64_mul(v47, v71)
									v77 = base.F64_mul(v41, v71)
									v78 = float64(1.90109156629516e-211)
								}
							}
							F_sq(m, v34+int32(24), v34+int32(16), v76)
							mBase = m.M
							F_sq(m, v34+int32(8), v34, v77)
							mBase = m.M
							v87 = *(*float64)(unsafe.Add(mBase, uint32(v34)))
							v88 = *(*float64)(unsafe.Add(mBase, uint32(v34)+16))
							v90 = *(*float64)(unsafe.Add(mBase, uint32(v34)+8))
							v92 = *(*float64)(unsafe.Add(mBase, uint32(v34)+24))
							v99 = base.F64_mul(v78, base.F64_sqrt(base.F64_add(base.F64_add(base.F64_add(v87, v88), v90), v92)))
						}
					}
				}
			}
			m.G0 = v34 + int32(32)
			if base.B2i32(base.Ui64(int64(9218868437227405312)) < base.Ui64(base.I64_reinterpret_f64(v16)&int64(9223372036854775807)))|base.F64_ne(v99, float64(0)) == int32(0) {
				v111 = F_float_zero_divide_error_ext(m, int32(0))
				mBase = m.M
				v112 = m.ExcPending
				if v112 != 0 {
					return int64(0)
				} else {
					v138 = v111
					*(*float64)(unsafe.Add(mBase, uint32(v10)+16)) = v138
					return base.I64_extend_i32_u(v10)
				}
			} else {
				v114 = math.Float64frombits(uint64(0x7ff0000000000000))
				v116 = base.F64_div(v16, v99)
				if base.F64_eq(base.F64_abs(v16), v114)|base.F64_ne(base.F64_abs(v116), v114) == int32(0) {
					v124 = F_float_overflow_error_ext(m, int32(0))
					mBase = m.M
					v125 = m.ExcPending
					if v125 != 0 {
						return int64(0)
					} else {
						v138 = v124
						*(*float64)(unsafe.Add(mBase, uint32(v10)+16)) = v138
						return base.I64_extend_i32_u(v10)
					}
				} else {
					v126 = float64(0)
					if base.F64_eq(v16, v126)|base.F64_ne(v116, v126)|base.F64_eq(base.F64_abs(v99), math.Float64frombits(uint64(0x7ff0000000000000))) != 0 {
						v138 = v116
						*(*float64)(unsafe.Add(mBase, uint32(v10)+16)) = v138
						return base.I64_extend_i32_u(v10)
					} else {
						v136 = F_float_underflow_error_ext(m, int32(0))
						mBase = m.M
						v137 = m.ExcPending
						if v137 != 0 {
							return int64(0)
						} else {
							v138 = v136
							*(*float64)(unsafe.Add(mBase, uint32(v10)+16)) = v138
							return base.I64_extend_i32_u(v10)
						}
					}
				}
			}
		}
	}
}
func F_circle_eq(m *base.Module, l0 int32) int64 {
	mBase := m.M
	_ = mBase
	var v6 int32
	_ = v6
	var v7 int32
	_ = v7
	var v8 float64
	_ = v8
	var v9 float64
	_ = v9
	var v11 float64
	_ = v11
	var v20 float64
	_ = v20
	var v23 int32
	_ = v23
	var v24 float64
	_ = v24
	var v30 float64
	_ = v30
	var v31 int32
	_ = v31
	var v32 float64
	_ = v32
	var v34 float64
	_ = v34
	var v37 float64
	_ = v37
	var v45 float64
	_ = v45
	var v46 int32
	_ = v46
	var v47 float64
	_ = v47
	var v53 float64
	_ = v53
	var v54 int32
	_ = v54
	var v55 float64
	_ = v55
	var v56 float64
	_ = v56
	var v57 float64
	_ = v57
	var v59 float64
	_ = v59
	var v68 float64
	_ = v68
	var v69 int32
	_ = v69
	var v70 float64
	_ = v70
	var v76 float64
	_ = v76
	var v77 int32
	_ = v77
	var v78 float64
	_ = v78
	var v80 float64
	_ = v80
	var v83 float64
	_ = v83
	var v91 float64
	_ = v91
	var v92 int32
	_ = v92
	var v93 float64
	_ = v93
	var v99 float64
	_ = v99
	var v100 int32
	_ = v100
	var v101 float64
	_ = v101
	v6 = *(*int32)(unsafe.Add(mBase, uint32(l0)+40))
	v7 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
	v8 = *(*float64)(unsafe.Add(mBase, uint32(v7)+16))
	v9 = base.F64_mul(v8, v8)
	v11 = math.Float64frombits(uint64(0x7ff0000000000000))
	if base.F64_ne(base.F64_abs(v9), v11)|base.F64_eq(base.F64_abs(v8), v11) == int32(0) {
		goto L2
	} else {
		goto L3
	}
L1:
	;
	v34 = math.Float64frombits(uint64(0x7ff0000000000000))
	v37 = base.F64_mul(v32, float64(3.141592653589793))
	if base.F64_eq(base.F64_abs(v32), v34)|base.F64_ne(base.F64_abs(v37), v34) == int32(0) {
		goto L10
	} else {
		goto L11
	}
L2:
	;
	v20 = F_float_overflow_error_ext(m, int32(0))
	mBase = m.M
	v23 = m.ExcPending
	if v23 != 0 {
		goto L5
	} else {
		goto L6
	}
L3:
	;
	goto L4
L4:
	;
	v24 = float64(0)
	if base.F64_eq(v8, v24)|base.F64_ne(v9, v24) != 0 {
		v32 = v9
		goto L1
	} else {
		goto L7
	}
L5:
	;
	return int64(0)
L6:
	;
	v32 = v20
	goto L1
L7:
	;
	v30 = F_float_underflow_error_ext(m, int32(0))
	mBase = m.M
	v31 = m.ExcPending
	if v31 != 0 {
		goto L5
	} else {
		goto L8
	}
L8:
	;
	v32 = v30
	goto L1
L9:
	;
	v56 = *(*float64)(unsafe.Add(mBase, uint32(v6)+16))
	v57 = base.F64_mul(v56, v56)
	v59 = math.Float64frombits(uint64(0x7ff0000000000000))
	if base.F64_ne(base.F64_abs(v57), v59)|base.F64_eq(base.F64_abs(v56), v59) == int32(0) {
		goto L17
	} else {
		goto L18
	}
L10:
	;
	v45 = F_float_overflow_error_ext(m, int32(0))
	mBase = m.M
	v46 = m.ExcPending
	if v46 != 0 {
		goto L5
	} else {
		goto L13
	}
L11:
	;
	goto L12
L12:
	;
	v47 = float64(0)
	if base.F64_eq(v32, v47)|base.F64_ne(v37, v47) != 0 {
		v55 = v37
		goto L9
	} else {
		goto L14
	}
L13:
	;
	v55 = v45
	goto L9
L14:
	;
	v53 = F_float_underflow_error_ext(m, int32(0))
	mBase = m.M
	v54 = m.ExcPending
	if v54 != 0 {
		goto L5
	} else {
		goto L15
	}
L15:
	;
	v55 = v53
	goto L9
L16:
	;
	v80 = math.Float64frombits(uint64(0x7ff0000000000000))
	v83 = base.F64_mul(v78, float64(3.141592653589793))
	if base.F64_eq(base.F64_abs(v78), v80)|base.F64_ne(base.F64_abs(v83), v80) == int32(0) {
		goto L24
	} else {
		goto L25
	}
L17:
	;
	v68 = F_float_overflow_error_ext(m, int32(0))
	mBase = m.M
	v69 = m.ExcPending
	if v69 != 0 {
		goto L5
	} else {
		goto L20
	}
L18:
	;
	goto L19
L19:
	;
	v70 = float64(0)
	if base.F64_eq(v56, v70)|base.F64_ne(v57, v70) != 0 {
		v78 = v57
		goto L16
	} else {
		goto L21
	}
L20:
	;
	v78 = v68
	goto L16
L21:
	;
	v76 = F_float_underflow_error_ext(m, int32(0))
	mBase = m.M
	v77 = m.ExcPending
	if v77 != 0 {
		goto L5
	} else {
		goto L22
	}
L22:
	;
	v78 = v76
	goto L16
L23:
	;
	return base.I64_extend_i32_u(base.F64_eq(v101, v55) | base.F64_le(base.F64_abs(base.F64_sub(v55, v101)), float64(1e-06)))
L24:
	;
	v91 = F_float_overflow_error_ext(m, int32(0))
	mBase = m.M
	v92 = m.ExcPending
	if v92 != 0 {
		goto L5
	} else {
		goto L27
	}
L25:
	;
	goto L26
L26:
	;
	v93 = float64(0)
	if base.F64_eq(v78, v93)|base.F64_ne(v83, v93) != 0 {
		v101 = v83
		goto L23
	} else {
		goto L28
	}
L27:
	;
	v101 = v91
	goto L23
L28:
	;
	v99 = F_float_underflow_error_ext(m, int32(0))
	mBase = m.M
	v100 = m.ExcPending
	if v100 != 0 {
		goto L5
	} else {
		goto L29
	}
L29:
	;
	v101 = v99
	goto L23
}
func F_circle_ge(m *base.Module, l0 int32) int64 {
	mBase := m.M
	_ = mBase
	var v6 int32
	_ = v6
	var v7 int32
	_ = v7
	var v8 float64
	_ = v8
	var v9 float64
	_ = v9
	var v11 float64
	_ = v11
	var v20 float64
	_ = v20
	var v23 int32
	_ = v23
	var v24 float64
	_ = v24
	var v30 float64
	_ = v30
	var v31 int32
	_ = v31
	var v32 float64
	_ = v32
	var v34 float64
	_ = v34
	var v37 float64
	_ = v37
	var v45 float64
	_ = v45
	var v46 int32
	_ = v46
	var v47 float64
	_ = v47
	var v53 float64
	_ = v53
	var v54 int32
	_ = v54
	var v55 float64
	_ = v55
	var v56 float64
	_ = v56
	var v57 float64
	_ = v57
	var v59 float64
	_ = v59
	var v68 float64
	_ = v68
	var v69 int32
	_ = v69
	var v70 float64
	_ = v70
	var v76 float64
	_ = v76
	var v77 int32
	_ = v77
	var v78 float64
	_ = v78
	var v80 float64
	_ = v80
	var v83 float64
	_ = v83
	var v91 float64
	_ = v91
	var v92 int32
	_ = v92
	var v93 float64
	_ = v93
	var v99 float64
	_ = v99
	var v100 int32
	_ = v100
	var v101 float64
	_ = v101
	v6 = *(*int32)(unsafe.Add(mBase, uint32(l0)+40))
	v7 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
	v8 = *(*float64)(unsafe.Add(mBase, uint32(v7)+16))
	v9 = base.F64_mul(v8, v8)
	v11 = math.Float64frombits(uint64(0x7ff0000000000000))
	if base.F64_ne(base.F64_abs(v9), v11)|base.F64_eq(base.F64_abs(v8), v11) == int32(0) {
		goto L2
	} else {
		goto L3
	}
L1:
	;
	v34 = math.Float64frombits(uint64(0x7ff0000000000000))
	v37 = base.F64_mul(v32, float64(3.141592653589793))
	if base.F64_eq(base.F64_abs(v32), v34)|base.F64_ne(base.F64_abs(v37), v34) == int32(0) {
		goto L10
	} else {
		goto L11
	}
L2:
	;
	v20 = F_float_overflow_error_ext(m, int32(0))
	mBase = m.M
	v23 = m.ExcPending
	if v23 != 0 {
		goto L5
	} else {
		goto L6
	}
L3:
	;
	goto L4
L4:
	;
	v24 = float64(0)
	if base.F64_eq(v8, v24)|base.F64_ne(v9, v24) != 0 {
		v32 = v9
		goto L1
	} else {
		goto L7
	}
L5:
	;
	return int64(0)
L6:
	;
	v32 = v20
	goto L1
L7:
	;
	v30 = F_float_underflow_error_ext(m, int32(0))
	mBase = m.M
	v31 = m.ExcPending
	if v31 != 0 {
		goto L5
	} else {
		goto L8
	}
L8:
	;
	v32 = v30
	goto L1
L9:
	;
	v56 = *(*float64)(unsafe.Add(mBase, uint32(v6)+16))
	v57 = base.F64_mul(v56, v56)
	v59 = math.Float64frombits(uint64(0x7ff0000000000000))
	if base.F64_ne(base.F64_abs(v57), v59)|base.F64_eq(base.F64_abs(v56), v59) == int32(0) {
		goto L17
	} else {
		goto L18
	}
L10:
	;
	v45 = F_float_overflow_error_ext(m, int32(0))
	mBase = m.M
	v46 = m.ExcPending
	if v46 != 0 {
		goto L5
	} else {
		goto L13
	}
L11:
	;
	goto L12
L12:
	;
	v47 = float64(0)
	if base.F64_eq(v32, v47)|base.F64_ne(v37, v47) != 0 {
		v55 = v37
		goto L9
	} else {
		goto L14
	}
L13:
	;
	v55 = v45
	goto L9
L14:
	;
	v53 = F_float_underflow_error_ext(m, int32(0))
	mBase = m.M
	v54 = m.ExcPending
	if v54 != 0 {
		goto L5
	} else {
		goto L15
	}
L15:
	;
	v55 = v53
	goto L9
L16:
	;
	v80 = math.Float64frombits(uint64(0x7ff0000000000000))
	v83 = base.F64_mul(v78, float64(3.141592653589793))
	if base.F64_eq(base.F64_abs(v78), v80)|base.F64_ne(base.F64_abs(v83), v80) == int32(0) {
		goto L24
	} else {
		goto L25
	}
L17:
	;
	v68 = F_float_overflow_error_ext(m, int32(0))
	mBase = m.M
	v69 = m.ExcPending
	if v69 != 0 {
		goto L5
	} else {
		goto L20
	}
L18:
	;
	goto L19
L19:
	;
	v70 = float64(0)
	if base.F64_eq(v56, v70)|base.F64_ne(v57, v70) != 0 {
		v78 = v57
		goto L16
	} else {
		goto L21
	}
L20:
	;
	v78 = v68
	goto L16
L21:
	;
	v76 = F_float_underflow_error_ext(m, int32(0))
	mBase = m.M
	v77 = m.ExcPending
	if v77 != 0 {
		goto L5
	} else {
		goto L22
	}
L22:
	;
	v78 = v76
	goto L16
L23:
	;
	return base.I64_extend_i32_u(base.F64_ge(base.F64_add(v55, float64(1e-06)), v101))
L24:
	;
	v91 = F_float_overflow_error_ext(m, int32(0))
	mBase = m.M
	v92 = m.ExcPending
	if v92 != 0 {
		goto L5
	} else {
		goto L27
	}
L25:
	;
	goto L26
L26:
	;
	v93 = float64(0)
	if base.F64_eq(v78, v93)|base.F64_ne(v83, v93) != 0 {
		v101 = v83
		goto L23
	} else {
		goto L28
	}
L27:
	;
	v101 = v91
	goto L23
L28:
	;
	v99 = F_float_underflow_error_ext(m, int32(0))
	mBase = m.M
	v100 = m.ExcPending
	if v100 != 0 {
		goto L5
	} else {
		goto L29
	}
L29:
	;
	v101 = v99
	goto L23
}
func F_circle_out(m *base.Module, l0 int32) int64 {
	mBase := m.M
	_ = mBase
	var v8 int32
	_ = v8
	var v10 int32
	_ = v10
	var v12 int32
	_ = v12
	var v14 int32
	_ = v14
	var v18 int32
	_ = v18
	var v21 int32
	_ = v21
	var v24 int32
	_ = v24
	var v25 float64
	_ = v25
	var v26 float64
	_ = v26
	var v27 int32
	_ = v27
	var v28 int32
	_ = v28
	var v29 int32
	_ = v29
	var v30 int32
	_ = v30
	var v35 int32
	_ = v35
	var v37 int32
	_ = v37
	var v39 int32
	_ = v39
	var v42 int32
	_ = v42
	var v45 int32
	_ = v45
	var v46 float64
	_ = v46
	var v47 int32
	_ = v47
	var v48 int32
	_ = v48
	var v50 int32
	_ = v50
	var v52 int32
	_ = v52
	var v55 int32
	_ = v55
	var v56 int64
	_ = v56
	v8 = m.G0
	v10 = v8 - int32(32)
	m.G0 = v10
	v12 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
	v14 = v10 + int32(16)
	F_initStringInfo(m, v14)
	mBase = m.M
	v18 = m.ExcPending
	if v18 != 0 {
		return int64(0)
	} else {
		F_appendStringInfoChar(m, v14, int32(60))
		mBase = m.M
		v21 = m.ExcPending
		if v21 != 0 {
			return int64(0)
		} else {
			F_appendStringInfoChar(m, v14, int32(40))
			mBase = m.M
			v24 = m.ExcPending
			if v24 != 0 {
				return int64(0)
			} else {
				v25 = *(*float64)(unsafe.Add(mBase, uint32(v12)+8))
				v26 = *(*float64)(unsafe.Add(mBase, uint32(v12)))
				v27 = F_float8out_internal(m, v26)
				mBase = m.M
				v28 = m.ExcPending
				if v28 != 0 {
					return int64(0)
				} else {
					v29 = F_float8out_internal(m, v25)
					mBase = m.M
					v30 = m.ExcPending
					if v30 != 0 {
						return int64(0)
					} else {
						*(*int32)(unsafe.Add(mBase, uint32(v10)+4)) = v29
						*(*int32)(unsafe.Add(mBase, uint32(v10))) = v27
						F_appendStringInfo(m, v14, int32(_a_F_circle_out_0), v10)
						mBase = m.M
						v35 = m.ExcPending
						if v35 != 0 {
							return int64(0)
						} else {
							F_pfree(m, v27)
							mBase = m.M
							v37 = m.ExcPending
							if v37 != 0 {
								return int64(0)
							} else {
								F_pfree(m, v29)
								mBase = m.M
								v39 = m.ExcPending
								if v39 != 0 {
									return int64(0)
								} else {
									F_appendStringInfoChar(m, v14, int32(41))
									mBase = m.M
									v42 = m.ExcPending
									if v42 != 0 {
										return int64(0)
									} else {
										F_appendStringInfoChar(m, v14, int32(44))
										mBase = m.M
										v45 = m.ExcPending
										if v45 != 0 {
											return int64(0)
										} else {
											v46 = *(*float64)(unsafe.Add(mBase, uint32(v12)+16))
											v47 = F_float8out_internal(m, v46)
											mBase = m.M
											v48 = m.ExcPending
											if v48 != 0 {
												return int64(0)
											} else {
												F_appendStringInfoString(m, v14, v47)
												mBase = m.M
												v50 = m.ExcPending
												if v50 != 0 {
													return int64(0)
												} else {
													F_pfree(m, v47)
													mBase = m.M
													v52 = m.ExcPending
													if v52 != 0 {
														return int64(0)
													} else {
														F_appendStringInfoChar(m, v14, int32(62))
														mBase = m.M
														v55 = m.ExcPending
														if v55 != 0 {
															return int64(0)
														} else {
															v56 = int64(*(*uint32)(unsafe.Add(mBase, uint32(v10)+16)))
															m.G0 = v10 + int32(32)
															return v56
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
			}
		}
	}
}
func F_circle_poly(m *base.Module, l0 int32) int64 {
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
	v2 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
	v3 = *(*int32)(unsafe.Add(mBase, uint32(l0)+40))
	v4 = F_circle_poly_internal(m, v2, v3, l0)
	mBase = m.M
	v7 = m.ExcPending
	if v7 != 0 {
		return int64(0)
	} else {
		return base.I64_extend_i32_u(v4)
	}
}
func F_circle_recv(m *base.Module, l0 int32) int64 {
	mBase := m.M
	_ = mBase
	var v4 int32
	_ = v4
	var v6 int32
	_ = v6
	var v9 int32
	_ = v9
	var v10 float64
	_ = v10
	var v11 int32
	_ = v11
	var v13 float64
	_ = v13
	var v14 int32
	_ = v14
	var v16 float64
	_ = v16
	var v17 int32
	_ = v17
	var v24 int32
	_ = v24
	var v27 int32
	_ = v27
	var v31 int32
	_ = v31
	var v36 int32
	_ = v36
	v4 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
	v6 = F_palloc(m, int32(24))
	mBase = m.M
	v9 = m.ExcPending
	if v9 != 0 {
		return int64(0)
	} else {
		v10 = F_pq_getmsgfloat8(m, v4)
		mBase = m.M
		v11 = m.ExcPending
		if v11 != 0 {
			return int64(0)
		} else {
			*(*float64)(unsafe.Add(mBase, uint32(v6))) = v10
			v13 = F_pq_getmsgfloat8(m, v4)
			mBase = m.M
			v14 = m.ExcPending
			if v14 != 0 {
				return int64(0)
			} else {
				*(*float64)(unsafe.Add(mBase, uint32(v6)+8)) = v13
				v16 = F_pq_getmsgfloat8(m, v4)
				mBase = m.M
				v17 = m.ExcPending
				if v17 != 0 {
					return int64(0)
				} else {
					*(*float64)(unsafe.Add(mBase, uint32(v6)+16)) = v16
					if base.F64_lt(v16, float64(0)) != 0 {
						F_errstart_cold(m, int32(21), int32(0))
						mBase = m.M
						v24 = m.ExcPending
						if v24 != 0 {
							return int64(0)
						} else {
							F_errcode(m, int32(50462850))
							mBase = m.M
							v27 = m.ExcPending
							if v27 != 0 {
								return int64(0)
							} else {
								F_errmsg(m, int32(_a_F_circle_recv_0), int32(0))
								mBase = m.M
								v31 = m.ExcPending
								if v31 != 0 {
									return int64(0)
								} else {
									F_errfinish(m, int32(_a_F_circle_recv_1), int32(_a_F_circle_recv_2), int32(_a_F_circle_recv_3))
									mBase = m.M
									v36 = m.ExcPending
									if v36 != 0 {
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
						return base.I64_extend_i32_u(v6)
					}
				}
			}
		}
	}
}
func F_circle_right(m *base.Module, l0 int32) int64 {
	mBase := m.M
	_ = mBase
	var v7 int32
	_ = v7
	var v8 int32
	_ = v8
	var v9 float64
	_ = v9
	var v10 float64
	_ = v10
	var v11 float64
	_ = v11
	var v13 float64
	_ = v13
	var v26 float64
	_ = v26
	var v29 int32
	_ = v29
	var v30 float64
	_ = v30
	var v31 float64
	_ = v31
	var v32 float64
	_ = v32
	var v33 float64
	_ = v33
	var v35 float64
	_ = v35
	var v46 float64
	_ = v46
	var v47 int32
	_ = v47
	var v48 float64
	_ = v48
	v7 = *(*int32)(unsafe.Add(mBase, uint32(l0)+40))
	v8 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
	v9 = *(*float64)(unsafe.Add(mBase, uint32(v8)))
	v10 = *(*float64)(unsafe.Add(mBase, uint32(v8)+16))
	v11 = base.F64_sub(v9, v10)
	v13 = math.Float64frombits(uint64(0x7ff0000000000000))
	if base.F64_ne(base.F64_abs(v11), v13)|base.F64_eq(base.F64_abs(v9), v13)|base.F64_eq(base.F64_abs(v10), v13) == int32(0) {
		v26 = F_float_overflow_error_ext(m, int32(0))
		mBase = m.M
		v29 = m.ExcPending
		if v29 != 0 {
			return int64(0)
		} else {
			v30 = v26
			v31 = *(*float64)(unsafe.Add(mBase, uint32(v7)))
			v32 = *(*float64)(unsafe.Add(mBase, uint32(v7)+16))
			v33 = base.F64_add(v31, v32)
			v35 = math.Float64frombits(uint64(0x7ff0000000000000))
			if base.F64_ne(base.F64_abs(v33), v35)|base.F64_eq(base.F64_abs(v31), v35)|base.F64_eq(base.F64_abs(v32), v35) != 0 {
				v48 = v33
				return base.I64_extend_i32_u(base.F64_gt(v30, base.F64_add(v48, float64(1e-06))))
			} else {
				v46 = F_float_overflow_error_ext(m, int32(0))
				mBase = m.M
				v47 = m.ExcPending
				if v47 != 0 {
					return int64(0)
				} else {
					v48 = v46
					return base.I64_extend_i32_u(base.F64_gt(v30, base.F64_add(v48, float64(1e-06))))
				}
			}
		}
	} else {
		v30 = v11
		v31 = *(*float64)(unsafe.Add(mBase, uint32(v7)))
		v32 = *(*float64)(unsafe.Add(mBase, uint32(v7)+16))
		v33 = base.F64_add(v31, v32)
		v35 = math.Float64frombits(uint64(0x7ff0000000000000))
		if base.F64_ne(base.F64_abs(v33), v35)|base.F64_eq(base.F64_abs(v31), v35)|base.F64_eq(base.F64_abs(v32), v35) != 0 {
			v48 = v33
			return base.I64_extend_i32_u(base.F64_gt(v30, base.F64_add(v48, float64(1e-06))))
		} else {
			v46 = F_float_overflow_error_ext(m, int32(0))
			mBase = m.M
			v47 = m.ExcPending
			if v47 != 0 {
				return int64(0)
			} else {
				v48 = v46
				return base.I64_extend_i32_u(base.F64_gt(v30, base.F64_add(v48, float64(1e-06))))
			}
		}
	}
}
