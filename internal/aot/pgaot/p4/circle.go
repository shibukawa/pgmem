package p4

import (
	base "github.com/shibukawa/pgmem/internal/aot/pgaot/base"
	"math"
	"unsafe"
)

func F_circle_above(m *base.Module, l0 int32) int64 {
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
	v9 = *(*float64)(unsafe.Add(mBase, uint32(v8)+8))
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
			v31 = *(*float64)(unsafe.Add(mBase, uint32(v7)+8))
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
		v31 = *(*float64)(unsafe.Add(mBase, uint32(v7)+8))
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
func F_circle_area(m *base.Module, l0 int32) int64 {
	mBase := m.M
	_ = mBase
	var v4 int32
	_ = v4
	var v5 float64
	_ = v5
	var v6 float64
	_ = v6
	var v8 float64
	_ = v8
	var v17 float64
	_ = v17
	var v20 int32
	_ = v20
	var v21 float64
	_ = v21
	var v27 float64
	_ = v27
	var v28 int32
	_ = v28
	var v29 float64
	_ = v29
	var v31 float64
	_ = v31
	var v34 float64
	_ = v34
	var v42 float64
	_ = v42
	var v43 int32
	_ = v43
	var v46 float64
	_ = v46
	var v52 float64
	_ = v52
	var v53 int32
	_ = v53
	var v54 float64
	_ = v54
	v4 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
	v5 = *(*float64)(unsafe.Add(mBase, uint32(v4)+16))
	v6 = base.F64_mul(v5, v5)
	v8 = math.Float64frombits(uint64(0x7ff0000000000000))
	if base.F64_ne(base.F64_abs(v6), v8)|base.F64_eq(base.F64_abs(v5), v8) == int32(0) {
		v17 = F_float_overflow_error_ext(m, int32(0))
		mBase = m.M
		v20 = m.ExcPending
		if v20 != 0 {
			return int64(0)
		} else {
			v29 = v17
			v31 = math.Float64frombits(uint64(0x7ff0000000000000))
			v34 = base.F64_mul(v29, float64(3.141592653589793))
			if base.F64_eq(base.F64_abs(v29), v31)|base.F64_ne(base.F64_abs(v34), v31) == int32(0) {
				v42 = F_float_overflow_error_ext(m, int32(0))
				mBase = m.M
				v43 = m.ExcPending
				if v43 != 0 {
					return int64(0)
				} else {
					return base.I64_reinterpret_f64(v42)
				}
			} else {
				v46 = float64(0)
				if base.F64_eq(v29, v46)|base.F64_ne(v34, v46) != 0 {
					v54 = v34
					return base.I64_reinterpret_f64(v54)
				} else {
					v52 = F_float_underflow_error_ext(m, int32(0))
					mBase = m.M
					v53 = m.ExcPending
					if v53 != 0 {
						return int64(0)
					} else {
						v54 = v52
						return base.I64_reinterpret_f64(v54)
					}
				}
			}
		}
	} else {
		v21 = float64(0)
		if base.F64_eq(v5, v21)|base.F64_ne(v6, v21) != 0 {
			v29 = v6
			v31 = math.Float64frombits(uint64(0x7ff0000000000000))
			v34 = base.F64_mul(v29, float64(3.141592653589793))
			if base.F64_eq(base.F64_abs(v29), v31)|base.F64_ne(base.F64_abs(v34), v31) == int32(0) {
				v42 = F_float_overflow_error_ext(m, int32(0))
				mBase = m.M
				v43 = m.ExcPending
				if v43 != 0 {
					return int64(0)
				} else {
					return base.I64_reinterpret_f64(v42)
				}
			} else {
				v46 = float64(0)
				if base.F64_eq(v29, v46)|base.F64_ne(v34, v46) != 0 {
					v54 = v34
					return base.I64_reinterpret_f64(v54)
				} else {
					v52 = F_float_underflow_error_ext(m, int32(0))
					mBase = m.M
					v53 = m.ExcPending
					if v53 != 0 {
						return int64(0)
					} else {
						v54 = v52
						return base.I64_reinterpret_f64(v54)
					}
				}
			}
		} else {
			v27 = F_float_underflow_error_ext(m, int32(0))
			mBase = m.M
			v28 = m.ExcPending
			if v28 != 0 {
				return int64(0)
			} else {
				v29 = v27
				v31 = math.Float64frombits(uint64(0x7ff0000000000000))
				v34 = base.F64_mul(v29, float64(3.141592653589793))
				if base.F64_eq(base.F64_abs(v29), v31)|base.F64_ne(base.F64_abs(v34), v31) == int32(0) {
					v42 = F_float_overflow_error_ext(m, int32(0))
					mBase = m.M
					v43 = m.ExcPending
					if v43 != 0 {
						return int64(0)
					} else {
						return base.I64_reinterpret_f64(v42)
					}
				} else {
					v46 = float64(0)
					if base.F64_eq(v29, v46)|base.F64_ne(v34, v46) != 0 {
						v54 = v34
						return base.I64_reinterpret_f64(v54)
					} else {
						v52 = F_float_underflow_error_ext(m, int32(0))
						mBase = m.M
						v53 = m.ExcPending
						if v53 != 0 {
							return int64(0)
						} else {
							v54 = v52
							return base.I64_reinterpret_f64(v54)
						}
					}
				}
			}
		}
	}
}
func F_circle_center(m *base.Module, l0 int32) int64 {
	mBase := m.M
	_ = mBase
	var v3 int32
	_ = v3
	var v5 int32
	_ = v5
	var v8 int32
	_ = v8
	var v9 float64
	_ = v9
	var v11 float64
	_ = v11
	v3 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
	v5 = F_palloc(m, int32(16))
	mBase = m.M
	v8 = m.ExcPending
	if v8 != 0 {
		return int64(0)
	} else {
		v9 = *(*float64)(unsafe.Add(mBase, uint32(v3)))
		*(*float64)(unsafe.Add(mBase, uint32(v5))) = v9
		v11 = *(*float64)(unsafe.Add(mBase, uint32(v3)+8))
		*(*float64)(unsafe.Add(mBase, uint32(v5)+8)) = v11
		return base.I64_extend_i32_u(v5)
	}
}
func F_circle_contained(m *base.Module, l0 int32) int64 {
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
					v52 = *(*float64)(unsafe.Add(mBase, uint32(v10)+16))
					v53 = *(*float64)(unsafe.Add(mBase, uint32(v8)+16))
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
				v52 = *(*float64)(unsafe.Add(mBase, uint32(v10)+16))
				v53 = *(*float64)(unsafe.Add(mBase, uint32(v8)+16))
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
				v52 = *(*float64)(unsafe.Add(mBase, uint32(v10)+16))
				v53 = *(*float64)(unsafe.Add(mBase, uint32(v8)+16))
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
			v52 = *(*float64)(unsafe.Add(mBase, uint32(v10)+16))
			v53 = *(*float64)(unsafe.Add(mBase, uint32(v8)+16))
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
func F_circle_diameter(m *base.Module, l0 int32) int64 {
	mBase := m.M
	_ = mBase
	var v4 int32
	_ = v4
	var v5 float64
	_ = v5
	var v6 float64
	_ = v6
	var v8 float64
	_ = v8
	var v17 float64
	_ = v17
	var v20 int32
	_ = v20
	var v23 float64
	_ = v23
	var v29 float64
	_ = v29
	var v30 int32
	_ = v30
	var v31 float64
	_ = v31
	v4 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
	v5 = *(*float64)(unsafe.Add(mBase, uint32(v4)+16))
	v6 = base.F64_add(v5, v5)
	v8 = math.Float64frombits(uint64(0x7ff0000000000000))
	if base.F64_ne(base.F64_abs(v6), v8)|base.F64_eq(base.F64_abs(v5), v8) == int32(0) {
		v17 = F_float_overflow_error_ext(m, int32(0))
		mBase = m.M
		v20 = m.ExcPending
		if v20 != 0 {
			return int64(0)
		} else {
			return base.I64_reinterpret_f64(v17)
		}
	} else {
		v23 = float64(0)
		if base.F64_eq(v5, v23)|base.F64_ne(v6, v23) != 0 {
			v31 = v6
			return base.I64_reinterpret_f64(v31)
		} else {
			v29 = F_float_underflow_error_ext(m, int32(0))
			mBase = m.M
			v30 = m.ExcPending
			if v30 != 0 {
				return int64(0)
			} else {
				v31 = v29
				return base.I64_reinterpret_f64(v31)
			}
		}
	}
}
func F_circle_distance(m *base.Module, l0 int32) int64 {
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
	var v128 float64
	_ = v128
	var v130 float64
	_ = v130
	var v131 float64
	_ = v131
	var v132 float64
	_ = v132
	var v147 float64
	_ = v147
	var v148 int32
	_ = v148
	var v149 float64
	_ = v149
	var v150 float64
	_ = v150
	var v152 float64
	_ = v152
	var v162 float64
	_ = v162
	var v163 int32
	_ = v163
	var v164 float64
	_ = v164
	var v165 float64
	_ = v165
	var v168 float64
	_ = v168
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
			v128 = math.Float64frombits(uint64(0x7ff0000000000000))
			v130 = *(*float64)(unsafe.Add(mBase, uint32(v6)+16))
			v131 = *(*float64)(unsafe.Add(mBase, uint32(v8)+16))
			v132 = base.F64_add(v130, v131)
			if base.F64_ne(base.F64_abs(v132), v128)|base.F64_eq(base.F64_abs(v130), v128)|base.F64_eq(base.F64_abs(v131), v128) == int32(0) {
				v147 = F_float_overflow_error_ext(m, int32(0))
				mBase = m.M
				v148 = m.ExcPending
				if v148 != 0 {
					return int64(0)
				} else {
					v149 = v147
					v150 = base.F64_sub(v121, v149)
					v152 = math.Float64frombits(uint64(0x7ff0000000000000))
					if base.F64_eq(base.F64_abs(v121), v128)|base.F64_ne(base.F64_abs(v150), v152)|base.F64_eq(base.F64_abs(v149), v152) == int32(0) {
						v162 = F_float_overflow_error_ext(m, int32(0))
						mBase = m.M
						v163 = m.ExcPending
						if v163 != 0 {
							return int64(0)
						} else {
							v164 = v162
							v165 = float64(0)
							if base.F64_lt(v164, v165) != 0 {
								v168 = v165
							} else {
								v168 = v164
							}
							return base.I64_reinterpret_f64(v168)
						}
					} else {
						v164 = v150
						v165 = float64(0)
						if base.F64_lt(v164, v165) != 0 {
							v168 = v165
						} else {
							v168 = v164
						}
						return base.I64_reinterpret_f64(v168)
					}
				}
			} else {
				v149 = v132
				v150 = base.F64_sub(v121, v149)
				v152 = math.Float64frombits(uint64(0x7ff0000000000000))
				if base.F64_eq(base.F64_abs(v121), v128)|base.F64_ne(base.F64_abs(v150), v152)|base.F64_eq(base.F64_abs(v149), v152) == int32(0) {
					v162 = F_float_overflow_error_ext(m, int32(0))
					mBase = m.M
					v163 = m.ExcPending
					if v163 != 0 {
						return int64(0)
					} else {
						v164 = v162
						v165 = float64(0)
						if base.F64_lt(v164, v165) != 0 {
							v168 = v165
						} else {
							v168 = v164
						}
						return base.I64_reinterpret_f64(v168)
					}
				} else {
					v164 = v150
					v165 = float64(0)
					if base.F64_lt(v164, v165) != 0 {
						v168 = v165
					} else {
						v168 = v164
					}
					return base.I64_reinterpret_f64(v168)
				}
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
				v128 = math.Float64frombits(uint64(0x7ff0000000000000))
				v130 = *(*float64)(unsafe.Add(mBase, uint32(v6)+16))
				v131 = *(*float64)(unsafe.Add(mBase, uint32(v8)+16))
				v132 = base.F64_add(v130, v131)
				if base.F64_ne(base.F64_abs(v132), v128)|base.F64_eq(base.F64_abs(v130), v128)|base.F64_eq(base.F64_abs(v131), v128) == int32(0) {
					v147 = F_float_overflow_error_ext(m, int32(0))
					mBase = m.M
					v148 = m.ExcPending
					if v148 != 0 {
						return int64(0)
					} else {
						v149 = v147
						v150 = base.F64_sub(v121, v149)
						v152 = math.Float64frombits(uint64(0x7ff0000000000000))
						if base.F64_eq(base.F64_abs(v121), v128)|base.F64_ne(base.F64_abs(v150), v152)|base.F64_eq(base.F64_abs(v149), v152) == int32(0) {
							v162 = F_float_overflow_error_ext(m, int32(0))
							mBase = m.M
							v163 = m.ExcPending
							if v163 != 0 {
								return int64(0)
							} else {
								v164 = v162
								v165 = float64(0)
								if base.F64_lt(v164, v165) != 0 {
									v168 = v165
								} else {
									v168 = v164
								}
								return base.I64_reinterpret_f64(v168)
							}
						} else {
							v164 = v150
							v165 = float64(0)
							if base.F64_lt(v164, v165) != 0 {
								v168 = v165
							} else {
								v168 = v164
							}
							return base.I64_reinterpret_f64(v168)
						}
					}
				} else {
					v149 = v132
					v150 = base.F64_sub(v121, v149)
					v152 = math.Float64frombits(uint64(0x7ff0000000000000))
					if base.F64_eq(base.F64_abs(v121), v128)|base.F64_ne(base.F64_abs(v150), v152)|base.F64_eq(base.F64_abs(v149), v152) == int32(0) {
						v162 = F_float_overflow_error_ext(m, int32(0))
						mBase = m.M
						v163 = m.ExcPending
						if v163 != 0 {
							return int64(0)
						} else {
							v164 = v162
							v165 = float64(0)
							if base.F64_lt(v164, v165) != 0 {
								v168 = v165
							} else {
								v168 = v164
							}
							return base.I64_reinterpret_f64(v168)
						}
					} else {
						v164 = v150
						v165 = float64(0)
						if base.F64_lt(v164, v165) != 0 {
							v168 = v165
						} else {
							v168 = v164
						}
						return base.I64_reinterpret_f64(v168)
					}
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
				v128 = math.Float64frombits(uint64(0x7ff0000000000000))
				v130 = *(*float64)(unsafe.Add(mBase, uint32(v6)+16))
				v131 = *(*float64)(unsafe.Add(mBase, uint32(v8)+16))
				v132 = base.F64_add(v130, v131)
				if base.F64_ne(base.F64_abs(v132), v128)|base.F64_eq(base.F64_abs(v130), v128)|base.F64_eq(base.F64_abs(v131), v128) == int32(0) {
					v147 = F_float_overflow_error_ext(m, int32(0))
					mBase = m.M
					v148 = m.ExcPending
					if v148 != 0 {
						return int64(0)
					} else {
						v149 = v147
						v150 = base.F64_sub(v121, v149)
						v152 = math.Float64frombits(uint64(0x7ff0000000000000))
						if base.F64_eq(base.F64_abs(v121), v128)|base.F64_ne(base.F64_abs(v150), v152)|base.F64_eq(base.F64_abs(v149), v152) == int32(0) {
							v162 = F_float_overflow_error_ext(m, int32(0))
							mBase = m.M
							v163 = m.ExcPending
							if v163 != 0 {
								return int64(0)
							} else {
								v164 = v162
								v165 = float64(0)
								if base.F64_lt(v164, v165) != 0 {
									v168 = v165
								} else {
									v168 = v164
								}
								return base.I64_reinterpret_f64(v168)
							}
						} else {
							v164 = v150
							v165 = float64(0)
							if base.F64_lt(v164, v165) != 0 {
								v168 = v165
							} else {
								v168 = v164
							}
							return base.I64_reinterpret_f64(v168)
						}
					}
				} else {
					v149 = v132
					v150 = base.F64_sub(v121, v149)
					v152 = math.Float64frombits(uint64(0x7ff0000000000000))
					if base.F64_eq(base.F64_abs(v121), v128)|base.F64_ne(base.F64_abs(v150), v152)|base.F64_eq(base.F64_abs(v149), v152) == int32(0) {
						v162 = F_float_overflow_error_ext(m, int32(0))
						mBase = m.M
						v163 = m.ExcPending
						if v163 != 0 {
							return int64(0)
						} else {
							v164 = v162
							v165 = float64(0)
							if base.F64_lt(v164, v165) != 0 {
								v168 = v165
							} else {
								v168 = v164
							}
							return base.I64_reinterpret_f64(v168)
						}
					} else {
						v164 = v150
						v165 = float64(0)
						if base.F64_lt(v164, v165) != 0 {
							v168 = v165
						} else {
							v168 = v164
						}
						return base.I64_reinterpret_f64(v168)
					}
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
					v128 = math.Float64frombits(uint64(0x7ff0000000000000))
					v130 = *(*float64)(unsafe.Add(mBase, uint32(v6)+16))
					v131 = *(*float64)(unsafe.Add(mBase, uint32(v8)+16))
					v132 = base.F64_add(v130, v131)
					if base.F64_ne(base.F64_abs(v132), v128)|base.F64_eq(base.F64_abs(v130), v128)|base.F64_eq(base.F64_abs(v131), v128) == int32(0) {
						v147 = F_float_overflow_error_ext(m, int32(0))
						mBase = m.M
						v148 = m.ExcPending
						if v148 != 0 {
							return int64(0)
						} else {
							v149 = v147
							v150 = base.F64_sub(v121, v149)
							v152 = math.Float64frombits(uint64(0x7ff0000000000000))
							if base.F64_eq(base.F64_abs(v121), v128)|base.F64_ne(base.F64_abs(v150), v152)|base.F64_eq(base.F64_abs(v149), v152) == int32(0) {
								v162 = F_float_overflow_error_ext(m, int32(0))
								mBase = m.M
								v163 = m.ExcPending
								if v163 != 0 {
									return int64(0)
								} else {
									v164 = v162
									v165 = float64(0)
									if base.F64_lt(v164, v165) != 0 {
										v168 = v165
									} else {
										v168 = v164
									}
									return base.I64_reinterpret_f64(v168)
								}
							} else {
								v164 = v150
								v165 = float64(0)
								if base.F64_lt(v164, v165) != 0 {
									v168 = v165
								} else {
									v168 = v164
								}
								return base.I64_reinterpret_f64(v168)
							}
						}
					} else {
						v149 = v132
						v150 = base.F64_sub(v121, v149)
						v152 = math.Float64frombits(uint64(0x7ff0000000000000))
						if base.F64_eq(base.F64_abs(v121), v128)|base.F64_ne(base.F64_abs(v150), v152)|base.F64_eq(base.F64_abs(v149), v152) == int32(0) {
							v162 = F_float_overflow_error_ext(m, int32(0))
							mBase = m.M
							v163 = m.ExcPending
							if v163 != 0 {
								return int64(0)
							} else {
								v164 = v162
								v165 = float64(0)
								if base.F64_lt(v164, v165) != 0 {
									v168 = v165
								} else {
									v168 = v164
								}
								return base.I64_reinterpret_f64(v168)
							}
						} else {
							v164 = v150
							v165 = float64(0)
							if base.F64_lt(v164, v165) != 0 {
								v168 = v165
							} else {
								v168 = v164
							}
							return base.I64_reinterpret_f64(v168)
						}
					}
				}
			}
		}
	}
}
func F_circle_in(m *base.Module, l0 int32) int64 {
	mBase := m.M
	_ = mBase
	var v11 int64
	_ = v11
	var v12 int32
	_ = v12
	var v14 int32
	_ = v14
	var v16 int32
	_ = v16
	var v17 int32
	_ = v17
	var v19 int32
	_ = v19
	var v22 int32
	_ = v22
	var v23 int32
	_ = v23
	var v26 int32
	_ = v26
	var v36 int32
	_ = v36
	var v50 int32
	_ = v50
	var v58 int32
	_ = v58
	var v59 int32
	_ = v59
	var v72 int32
	_ = v72
	var v83 int32
	_ = v83
	var v89 int32
	_ = v89
	var v97 int32
	_ = v97
	var v98 int32
	_ = v98
	var v101 int32
	_ = v101
	var v103 int32
	_ = v103
	var v104 int32
	_ = v104
	var v108 int32
	_ = v108
	var v110 int32
	_ = v110
	var v114 float64
	_ = v114
	var v115 int32
	_ = v115
	var v119 int32
	_ = v119
	var v122 int32
	_ = v122
	var v125 int32
	_ = v125
	var v131 int32
	_ = v131
	var v132 int32
	_ = v132
	var v138 int32
	_ = v138
	var v139 int32
	_ = v139
	var v144 int32
	_ = v144
	var v152 int32
	_ = v152
	var v157 int32
	_ = v157
	var v158 int32
	_ = v158
	var v159 int32
	_ = v159
	var v164 int32
	_ = v164
	var v170 int32
	_ = v170
	var v175 int32
	_ = v175
	var v178 int32
	_ = v178
	var v188 int32
	_ = v188
	var v190 int32
	_ = v190
	var v202 int32
	_ = v202
	var v209 int32
	_ = v209
	var v210 int32
	_ = v210
	var v215 int32
	_ = v215
	var v223 int32
	_ = v223
	var v228 int32
	_ = v228
	var v240 int64
	_ = v240
	v11 = int64(0)
	v12 = m.G0
	v14 = v12 - int32(48)
	m.G0 = v14
	v16 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v17 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
	v19 = F_palloc(m, int32(24))
	mBase = m.M
	v22 = m.ExcPending
	if v22 != 0 {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	return int64(0)
L2:
	;
	v23 = int32(1)
	v26 = v17
	goto L6
L3:
	;
	v97 = F_pair_decode(m, v83, v19, v19+int32(8), v14+int32(44), int32(_a_F_circle_in_0), v17, v16)
	mBase = m.M
	v98 = m.ExcPending
	if v98 != 0 {
		goto L1
	} else {
		goto L15
	}
L4:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v14)+44)) = v72
	v83 = v72
	v89 = int32(0)
	goto L3
L5:
	;
	v72 = v26 + int32(1)
	goto L4
L6:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v14)+44)) = v26
	v36 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v26))))
	if base.Ui32(v36-int32(9)) < base.Ui32(int32(5)) {
		goto L9
	} else {
		goto L10
	}
L7:
	;
	v50 = v26
	goto L11
L8:
	;
	goto L7
L9:
	;
	v26 = v26 + int32(1)
	goto L6
L10:
	;
	switch v36 - int32(32) {
	case 0:
		goto L9
	default:
		v83 = v26
		v89 = v23
		goto L3
	case 8:
		goto L8
	case 28:
		goto L5
	}
L11:
	;
	v58 = v50 + int32(1)
	v59 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v50)+1)))
	if base.Ui32(v59-int32(9)) < base.Ui32(int32(5)) {
		v50 = v58
		goto L11
	} else {
		goto L13
	}
L13:
	;
	switch v59 - int32(32) {
	case 0:
		v50 = v58
		goto L11
	default:
		v83 = v26
		v89 = v23
		goto L3
	case 8:
		v72 = v58
		goto L4
	}
L14:
	;
	m.G0 = v14 + int32(48)
	return v240
L15:
	;
	if v97 == int32(0) {
		goto L16
	} else {
		goto L17
	}
L16:
	;
	v101 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(l0)+16)) = uint8(v101)
	v240 = v11
	goto L14
L17:
	;
	goto L18
L18:
	;
	v103 = *(*int32)(unsafe.Add(mBase, uint32(v14)+44))
	v104 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v103))))
	if v104 == int32(44) {
		goto L19
	} else {
		goto L20
	}
L19:
	;
	v108 = v103 + int32(1)
	*(*int32)(unsafe.Add(mBase, uint32(v14)+44)) = v108
	v110 = v108
	goto L21
L20:
	;
	v110 = v103
	goto L21
L21:
	;
	v114 = F_float8in_internal(m, v110, v14+int32(44), int32(_a_F_circle_in_0), v17, v16)
	mBase = m.M
	v115 = m.ExcPending
	if v115 != 0 {
		goto L1
	} else {
		goto L22
	}
L22:
	;
	*(*float64)(unsafe.Add(mBase, uint32(v19)+16)) = v114
	if v16 == int32(0) {
		goto L23
	} else {
		goto L24
	}
L23:
	;
	if base.F64_lt(v114, float64(0)) == int32(0) {
		goto L29
	} else {
		goto L30
	}
L24:
	;
	v119 = *(*int32)(unsafe.Add(mBase, uint32(v16)))
	if v119 != int32(453) {
		goto L23
	} else {
		goto L25
	}
L25:
	;
	v122 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v16)+4)))
	if v122 != int32(1) {
		goto L23
	} else {
		goto L26
	}
L26:
	;
	v125 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(l0)+16)) = uint8(v125)
	v240 = v11
	goto L14
L27:
	;
	if v202 != 0 {
		goto L47
	} else {
		goto L48
	}
L28:
	;
	v178 = v131
	goto L44
L29:
	;
	v131 = *(*int32)(unsafe.Add(mBase, uint32(v14)+44))
	v132 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v131))))
	if v89 != 0 {
		v202 = v132
		goto L27
	} else {
		goto L32
	}
L30:
	;
	goto L31
L31:
	;
	v158 = F_errsave_start(m, v16)
	mBase = m.M
	v159 = m.ExcPending
	if v159 != 0 {
		goto L1
	} else {
		goto L39
	}
L32:
	;
	if base.B2i32(v132 == int32(62))|base.B2i32(v132 == int32(41)) != 0 {
		goto L28
	} else {
		goto L33
	}
L33:
	;
	v138 = F_errsave_start(m, v16)
	mBase = m.M
	v139 = m.ExcPending
	if v139 != 0 {
		goto L1
	} else {
		goto L34
	}
L34:
	;
	if v138 == int32(0) {
		v240 = v11
		goto L14
	} else {
		goto L35
	}
L35:
	;
	F_errcode(m, int32(33685634))
	mBase = m.M
	v144 = m.ExcPending
	if v144 != 0 {
		goto L1
	} else {
		goto L36
	}
L36:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v14)+36)) = v17
	*(*int32)(unsafe.Add(mBase, uint32(v14)+32)) = int32(_a_F_circle_in_0)
	F_errmsg(m, int32(_a_F_circle_in_1), v14+int32(32))
	mBase = m.M
	v152 = m.ExcPending
	if v152 != 0 {
		goto L1
	} else {
		goto L37
	}
L37:
	;
	F_errsave_finish(m, v16, int32(_a_F_circle_in_2), int32(_a_F_circle_in_3), int32(_a_F_circle_in_4))
	mBase = m.M
	v157 = m.ExcPending
	if v157 != 0 {
		goto L1
	} else {
		goto L38
	}
L38:
	;
	v240 = v11
	goto L14
L39:
	;
	if v158 == int32(0) {
		v240 = v11
		goto L14
	} else {
		goto L40
	}
L40:
	;
	F_errcode(m, int32(33685634))
	mBase = m.M
	v164 = m.ExcPending
	if v164 != 0 {
		goto L1
	} else {
		goto L41
	}
L41:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v14)+4)) = v17
	*(*int32)(unsafe.Add(mBase, uint32(v14))) = int32(_a_F_circle_in_0)
	F_errmsg(m, int32(_a_F_circle_in_1), v14)
	mBase = m.M
	v170 = m.ExcPending
	if v170 != 0 {
		goto L1
	} else {
		goto L42
	}
L42:
	;
	F_errsave_finish(m, v16, int32(_a_F_circle_in_2), int32(_a_F_circle_in_5), int32(_a_F_circle_in_4))
	mBase = m.M
	v175 = m.ExcPending
	if v175 != 0 {
		goto L1
	} else {
		goto L43
	}
L43:
	;
	v240 = v11
	goto L14
L44:
	;
	v188 = v178 + int32(1)
	*(*int32)(unsafe.Add(mBase, uint32(v14)+44)) = v188
	v190 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v178)+1)))
	if base.B2i32(base.Ui32(v190-int32(9)) < base.Ui32(int32(5)))|base.B2i32(v190 == int32(32)) != 0 {
		v178 = v188
		goto L44
	} else {
		goto L46
	}
L45:
	;
	v202 = v190
	goto L27
L46:
	;
	goto L45
L47:
	;
	v209 = F_errsave_start(m, v16)
	mBase = m.M
	v210 = m.ExcPending
	if v210 != 0 {
		goto L1
	} else {
		goto L50
	}
L48:
	;
	goto L49
L49:
	;
	v240 = base.I64_extend_i32_u(v19)
	goto L14
L50:
	;
	if v209 == int32(0) {
		v240 = v11
		goto L14
	} else {
		goto L51
	}
L51:
	;
	F_errcode(m, int32(33685634))
	mBase = m.M
	v215 = m.ExcPending
	if v215 != 0 {
		goto L1
	} else {
		goto L52
	}
L52:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v14)+20)) = v17
	*(*int32)(unsafe.Add(mBase, uint32(v14)+16)) = int32(_a_F_circle_in_0)
	F_errmsg(m, int32(_a_F_circle_in_1), v14+int32(16))
	mBase = m.M
	v223 = m.ExcPending
	if v223 != 0 {
		goto L1
	} else {
		goto L53
	}
L53:
	;
	F_errsave_finish(m, v16, int32(_a_F_circle_in_2), int32(_a_F_circle_in_6), int32(_a_F_circle_in_4))
	mBase = m.M
	v228 = m.ExcPending
	if v228 != 0 {
		goto L1
	} else {
		goto L54
	}
L54:
	;
	v240 = v11
	goto L14
}
func F_circle_ne(m *base.Module, l0 int32) int64 {
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
	return base.I64_extend_i32_u(base.F64_ne(v101, v55) & base.F64_gt(base.F64_abs(base.F64_sub(v55, v101)), float64(1e-06)))
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
func F_circle_overleft(m *base.Module, l0 int32) int64 {
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
	v11 = base.F64_add(v9, v10)
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
				return base.I64_extend_i32_u(base.F64_le(v30, base.F64_add(v48, float64(1e-06))))
			} else {
				v46 = F_float_overflow_error_ext(m, int32(0))
				mBase = m.M
				v47 = m.ExcPending
				if v47 != 0 {
					return int64(0)
				} else {
					v48 = v46
					return base.I64_extend_i32_u(base.F64_le(v30, base.F64_add(v48, float64(1e-06))))
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
			return base.I64_extend_i32_u(base.F64_le(v30, base.F64_add(v48, float64(1e-06))))
		} else {
			v46 = F_float_overflow_error_ext(m, int32(0))
			mBase = m.M
			v47 = m.ExcPending
			if v47 != 0 {
				return int64(0)
			} else {
				v48 = v46
				return base.I64_extend_i32_u(base.F64_le(v30, base.F64_add(v48, float64(1e-06))))
			}
		}
	}
}
