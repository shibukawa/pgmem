package p1

import (
	base "github.com/shibukawa/pgmem/internal/aot/pgaot/base"
	"math"
	"unsafe"
)

func F_box_ar(m *base.Module, l0 int32) float64 {
	mBase := m.M
	_ = mBase
	var v6 float64
	_ = v6
	var v7 float64
	_ = v7
	var v8 float64
	_ = v8
	var v10 float64
	_ = v10
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
	var v45 float64
	_ = v45
	var v46 int32
	_ = v46
	var v47 float64
	_ = v47
	var v49 float64
	_ = v49
	var v51 float64
	_ = v51
	var v63 float64
	_ = v63
	var v64 int32
	_ = v64
	var v66 float64
	_ = v66
	var v75 float64
	_ = v75
	var v76 int32
	_ = v76
	var v77 float64
	_ = v77
	v6 = *(*float64)(unsafe.Add(mBase, uint32(l0)))
	v7 = *(*float64)(unsafe.Add(mBase, uint32(l0)+16))
	v8 = base.F64_sub(v6, v7)
	v10 = math.Float64frombits(uint64(0x7ff0000000000000))
	if base.F64_ne(base.F64_abs(v8), v10)|base.F64_eq(base.F64_abs(v6), v10)|base.F64_eq(base.F64_abs(v7), v10) == int32(0) {
		v23 = F_float_overflow_error_ext(m, int32(0))
		mBase = m.M
		v26 = m.ExcPending
		if v26 != 0 {
			return float64(0)
		} else {
			v27 = v23
			v28 = *(*float64)(unsafe.Add(mBase, uint32(l0)+8))
			v29 = *(*float64)(unsafe.Add(mBase, uint32(l0)+24))
			v30 = base.F64_sub(v28, v29)
			v32 = math.Float64frombits(uint64(0x7ff0000000000000))
			if base.F64_ne(base.F64_abs(v30), v32)|base.F64_eq(base.F64_abs(v28), v32)|base.F64_eq(base.F64_abs(v29), v32) == int32(0) {
				v45 = F_float_overflow_error_ext(m, int32(0))
				mBase = m.M
				v46 = m.ExcPending
				if v46 != 0 {
					return float64(0)
				} else {
					v47 = v45
					v49 = math.Float64frombits(uint64(0x7ff0000000000000))
					v51 = base.F64_mul(v27, v47)
					if base.F64_eq(base.F64_abs(v27), v49)|base.F64_ne(base.F64_abs(v51), v49)|base.F64_eq(base.F64_abs(v47), v49) == int32(0) {
						v63 = F_float_overflow_error_ext(m, int32(0))
						mBase = m.M
						v64 = m.ExcPending
						if v64 != 0 {
							return float64(0)
						} else {
							return v63
						}
					} else {
						v66 = float64(0)
						if base.F64_eq(v27, v66)|base.F64_ne(v51, v66)|base.F64_eq(v47, v66) != 0 {
							v77 = v51
							return v77
						} else {
							v75 = F_float_underflow_error_ext(m, int32(0))
							mBase = m.M
							v76 = m.ExcPending
							if v76 != 0 {
								return float64(0)
							} else {
								v77 = v75
								return v77
							}
						}
					}
				}
			} else {
				v47 = v30
				v49 = math.Float64frombits(uint64(0x7ff0000000000000))
				v51 = base.F64_mul(v27, v47)
				if base.F64_eq(base.F64_abs(v27), v49)|base.F64_ne(base.F64_abs(v51), v49)|base.F64_eq(base.F64_abs(v47), v49) == int32(0) {
					v63 = F_float_overflow_error_ext(m, int32(0))
					mBase = m.M
					v64 = m.ExcPending
					if v64 != 0 {
						return float64(0)
					} else {
						return v63
					}
				} else {
					v66 = float64(0)
					if base.F64_eq(v27, v66)|base.F64_ne(v51, v66)|base.F64_eq(v47, v66) != 0 {
						v77 = v51
						return v77
					} else {
						v75 = F_float_underflow_error_ext(m, int32(0))
						mBase = m.M
						v76 = m.ExcPending
						if v76 != 0 {
							return float64(0)
						} else {
							v77 = v75
							return v77
						}
					}
				}
			}
		}
	} else {
		v27 = v8
		v28 = *(*float64)(unsafe.Add(mBase, uint32(l0)+8))
		v29 = *(*float64)(unsafe.Add(mBase, uint32(l0)+24))
		v30 = base.F64_sub(v28, v29)
		v32 = math.Float64frombits(uint64(0x7ff0000000000000))
		if base.F64_ne(base.F64_abs(v30), v32)|base.F64_eq(base.F64_abs(v28), v32)|base.F64_eq(base.F64_abs(v29), v32) == int32(0) {
			v45 = F_float_overflow_error_ext(m, int32(0))
			mBase = m.M
			v46 = m.ExcPending
			if v46 != 0 {
				return float64(0)
			} else {
				v47 = v45
				v49 = math.Float64frombits(uint64(0x7ff0000000000000))
				v51 = base.F64_mul(v27, v47)
				if base.F64_eq(base.F64_abs(v27), v49)|base.F64_ne(base.F64_abs(v51), v49)|base.F64_eq(base.F64_abs(v47), v49) == int32(0) {
					v63 = F_float_overflow_error_ext(m, int32(0))
					mBase = m.M
					v64 = m.ExcPending
					if v64 != 0 {
						return float64(0)
					} else {
						return v63
					}
				} else {
					v66 = float64(0)
					if base.F64_eq(v27, v66)|base.F64_ne(v51, v66)|base.F64_eq(v47, v66) != 0 {
						v77 = v51
						return v77
					} else {
						v75 = F_float_underflow_error_ext(m, int32(0))
						mBase = m.M
						v76 = m.ExcPending
						if v76 != 0 {
							return float64(0)
						} else {
							v77 = v75
							return v77
						}
					}
				}
			}
		} else {
			v47 = v30
			v49 = math.Float64frombits(uint64(0x7ff0000000000000))
			v51 = base.F64_mul(v27, v47)
			if base.F64_eq(base.F64_abs(v27), v49)|base.F64_ne(base.F64_abs(v51), v49)|base.F64_eq(base.F64_abs(v47), v49) == int32(0) {
				v63 = F_float_overflow_error_ext(m, int32(0))
				mBase = m.M
				v64 = m.ExcPending
				if v64 != 0 {
					return float64(0)
				} else {
					return v63
				}
			} else {
				v66 = float64(0)
				if base.F64_eq(v27, v66)|base.F64_ne(v51, v66)|base.F64_eq(v47, v66) != 0 {
					v77 = v51
					return v77
				} else {
					v75 = F_float_underflow_error_ext(m, int32(0))
					mBase = m.M
					v76 = m.ExcPending
					if v76 != 0 {
						return float64(0)
					} else {
						v77 = v75
						return v77
					}
				}
			}
		}
	}
}
func F_box_center(m *base.Module, l0 int32) int64 {
	mBase := m.M
	_ = mBase
	var v4 int32
	_ = v4
	var v6 int32
	_ = v6
	var v9 int32
	_ = v9
	var v10 int32
	_ = v10
	var v12 int32
	_ = v12
	var v13 int32
	_ = v13
	var v16 int32
	_ = v16
	var v19 int32
	_ = v19
	var v22 int32
	_ = v22
	v4 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
	v6 = F_palloc(m, int32(16))
	mBase = m.M
	v9 = m.ExcPending
	if v9 != 0 {
		return int64(0)
	} else {
		v10 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
		F_box_cn(m, v6, v4, v10)
		mBase = m.M
		v12 = m.ExcPending
		if v12 != 0 {
			return int64(0)
		} else {
			v13 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
			if v13 == int32(0) {
				return base.I64_extend_i32_u(v6)
			} else {
				v16 = *(*int32)(unsafe.Add(mBase, uint32(v13)))
				if v16 != int32(453) {
					return base.I64_extend_i32_u(v6)
				} else {
					v19 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v13)+4)))
					if v19 != int32(1) {
						return base.I64_extend_i32_u(v6)
					} else {
						v22 = int32(1)
						*(*uint8)(unsafe.Add(mBase, uint32(l0)+16)) = uint8(v22)
						return int64(0)
					}
				}
			}
		}
	}
}
func F_box_distance(m *base.Module, l0 int32) int64 {
	mBase := m.M
	_ = mBase
	var v7 int32
	_ = v7
	var v9 int32
	_ = v9
	var v11 int32
	_ = v11
	var v14 int32
	_ = v14
	var v19 int32
	_ = v19
	var v22 int32
	_ = v22
	var v23 float64
	_ = v23
	var v24 float64
	_ = v24
	var v25 float64
	_ = v25
	var v27 float64
	_ = v27
	var v38 float64
	_ = v38
	var v39 int32
	_ = v39
	var v40 float64
	_ = v40
	var v41 float64
	_ = v41
	var v42 float64
	_ = v42
	var v43 float64
	_ = v43
	var v45 float64
	_ = v45
	var v56 float64
	_ = v56
	var v57 int32
	_ = v57
	var v58 float64
	_ = v58
	var v67 int32
	_ = v67
	var v69 int32
	_ = v69
	var v71 float64
	_ = v71
	var v72 float64
	_ = v72
	var v75 int32
	_ = v75
	var v76 float64
	_ = v76
	var v77 int64
	_ = v77
	var v79 int64
	_ = v79
	var v82 float64
	_ = v82
	var v85 int64
	_ = v85
	var v87 int64
	_ = v87
	var v98 float64
	_ = v98
	var v106 float64
	_ = v106
	var v111 float64
	_ = v111
	var v112 float64
	_ = v112
	var v113 float64
	_ = v113
	var v122 float64
	_ = v122
	var v123 float64
	_ = v123
	var v125 float64
	_ = v125
	var v127 float64
	_ = v127
	var v134 float64
	_ = v134
	v7 = m.G0
	v9 = v7 - int32(32)
	m.G0 = v9
	v11 = *(*int32)(unsafe.Add(mBase, uint32(l0)+40))
	v14 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
	F_box_cn(m, v9+int32(16), v14, int32(0))
	mBase = m.M
	v19 = m.ExcPending
	if v19 != 0 {
		return int64(0)
	} else {
		F_box_cn(m, v9, v11, int32(0))
		mBase = m.M
		v22 = m.ExcPending
		if v22 != 0 {
			return int64(0)
		} else {
			v23 = *(*float64)(unsafe.Add(mBase, uint32(v9)+16))
			v24 = *(*float64)(unsafe.Add(mBase, uint32(v9)))
			v25 = base.F64_sub(v23, v24)
			v27 = math.Float64frombits(uint64(0x7ff0000000000000))
			if base.F64_ne(base.F64_abs(v25), v27)|base.F64_eq(base.F64_abs(v23), v27)|base.F64_eq(base.F64_abs(v24), v27) != 0 {
				v40 = v25
				v41 = *(*float64)(unsafe.Add(mBase, uint32(v9)+24))
				v42 = *(*float64)(unsafe.Add(mBase, uint32(v9)+8))
				v43 = base.F64_sub(v41, v42)
				v45 = math.Float64frombits(uint64(0x7ff0000000000000))
				if base.F64_ne(base.F64_abs(v43), v45)|base.F64_eq(base.F64_abs(v41), v45)|base.F64_eq(base.F64_abs(v42), v45) != 0 {
					v58 = v43
					v67 = m.G0
					v69 = v67 - int32(32)
					m.G0 = v69
					v71 = base.F64_abs(v40)
					v72 = base.F64_abs(v58)
					v75 = base.B2i32(base.Ui64(base.I64_reinterpret_f64(v71)) < base.Ui64(base.I64_reinterpret_f64(v72)))
					if base.Ui64(base.I64_reinterpret_f64(v71)) < base.Ui64(base.I64_reinterpret_f64(v72)) {
						v76 = v71
					} else {
						v76 = v72
					}
					v77 = base.I64_reinterpret_f64(v76)
					v79 = int64(base.Ui64(v77) >> (uint(int64(52)) % 64))
					if v79 == int64(2047) {
						v134 = v76
					} else {
						if base.Ui64(base.I64_reinterpret_f64(v71)) < base.Ui64(base.I64_reinterpret_f64(v72)) {
							v82 = v72
						} else {
							v82 = v71
						}
						if v77 == int64(0) {
							v134 = v82
						} else {
							v85 = base.I64_reinterpret_f64(v82)
							v87 = int64(base.Ui64(v85) >> (uint(int64(52)) % 64))
							if v87 == int64(2047) {
								v134 = v82
							} else {
								if int32(65) <= base.I32_wrap_i64(v87)-base.I32_wrap_i64(v79) {
									v134 = base.F64_add(v71, v72)
								} else {
									if base.Ui64(int64(6908521828386340864)) <= base.Ui64(v85) {
										v98 = float64(1.90109156629516e-211)
										v111 = base.F64_mul(v82, v98)
										v112 = base.F64_mul(v76, v98)
										v113 = float64(5.260135901548374e+210)
									} else {
										if base.Ui64(int64(2580562586483294207)) < base.Ui64(v77) {
											v111 = v82
											v112 = v76
											v113 = float64(1)
										} else {
											v106 = float64(5.260135901548374e+210)
											v111 = base.F64_mul(v82, v106)
											v112 = base.F64_mul(v76, v106)
											v113 = float64(1.90109156629516e-211)
										}
									}
									F_sq(m, v69+int32(24), v69+int32(16), v111)
									mBase = m.M
									F_sq(m, v69+int32(8), v69, v112)
									mBase = m.M
									v122 = *(*float64)(unsafe.Add(mBase, uint32(v69)))
									v123 = *(*float64)(unsafe.Add(mBase, uint32(v69)+16))
									v125 = *(*float64)(unsafe.Add(mBase, uint32(v69)+8))
									v127 = *(*float64)(unsafe.Add(mBase, uint32(v69)+24))
									v134 = base.F64_mul(v113, base.F64_sqrt(base.F64_add(base.F64_add(base.F64_add(v122, v123), v125), v127)))
								}
							}
						}
					}
					m.G0 = v69 + int32(32)
					m.G0 = v9 + int32(32)
					return base.I64_reinterpret_f64(v134)
				} else {
					v56 = F_float_overflow_error_ext(m, int32(0))
					mBase = m.M
					v57 = m.ExcPending
					if v57 != 0 {
						return int64(0)
					} else {
						v58 = v56
						v67 = m.G0
						v69 = v67 - int32(32)
						m.G0 = v69
						v71 = base.F64_abs(v40)
						v72 = base.F64_abs(v58)
						v75 = base.B2i32(base.Ui64(base.I64_reinterpret_f64(v71)) < base.Ui64(base.I64_reinterpret_f64(v72)))
						if base.Ui64(base.I64_reinterpret_f64(v71)) < base.Ui64(base.I64_reinterpret_f64(v72)) {
							v76 = v71
						} else {
							v76 = v72
						}
						v77 = base.I64_reinterpret_f64(v76)
						v79 = int64(base.Ui64(v77) >> (uint(int64(52)) % 64))
						if v79 == int64(2047) {
							v134 = v76
						} else {
							if base.Ui64(base.I64_reinterpret_f64(v71)) < base.Ui64(base.I64_reinterpret_f64(v72)) {
								v82 = v72
							} else {
								v82 = v71
							}
							if v77 == int64(0) {
								v134 = v82
							} else {
								v85 = base.I64_reinterpret_f64(v82)
								v87 = int64(base.Ui64(v85) >> (uint(int64(52)) % 64))
								if v87 == int64(2047) {
									v134 = v82
								} else {
									if int32(65) <= base.I32_wrap_i64(v87)-base.I32_wrap_i64(v79) {
										v134 = base.F64_add(v71, v72)
									} else {
										if base.Ui64(int64(6908521828386340864)) <= base.Ui64(v85) {
											v98 = float64(1.90109156629516e-211)
											v111 = base.F64_mul(v82, v98)
											v112 = base.F64_mul(v76, v98)
											v113 = float64(5.260135901548374e+210)
										} else {
											if base.Ui64(int64(2580562586483294207)) < base.Ui64(v77) {
												v111 = v82
												v112 = v76
												v113 = float64(1)
											} else {
												v106 = float64(5.260135901548374e+210)
												v111 = base.F64_mul(v82, v106)
												v112 = base.F64_mul(v76, v106)
												v113 = float64(1.90109156629516e-211)
											}
										}
										F_sq(m, v69+int32(24), v69+int32(16), v111)
										mBase = m.M
										F_sq(m, v69+int32(8), v69, v112)
										mBase = m.M
										v122 = *(*float64)(unsafe.Add(mBase, uint32(v69)))
										v123 = *(*float64)(unsafe.Add(mBase, uint32(v69)+16))
										v125 = *(*float64)(unsafe.Add(mBase, uint32(v69)+8))
										v127 = *(*float64)(unsafe.Add(mBase, uint32(v69)+24))
										v134 = base.F64_mul(v113, base.F64_sqrt(base.F64_add(base.F64_add(base.F64_add(v122, v123), v125), v127)))
									}
								}
							}
						}
						m.G0 = v69 + int32(32)
						m.G0 = v9 + int32(32)
						return base.I64_reinterpret_f64(v134)
					}
				}
			} else {
				v38 = F_float_overflow_error_ext(m, int32(0))
				mBase = m.M
				v39 = m.ExcPending
				if v39 != 0 {
					return int64(0)
				} else {
					v40 = v38
					v41 = *(*float64)(unsafe.Add(mBase, uint32(v9)+24))
					v42 = *(*float64)(unsafe.Add(mBase, uint32(v9)+8))
					v43 = base.F64_sub(v41, v42)
					v45 = math.Float64frombits(uint64(0x7ff0000000000000))
					if base.F64_ne(base.F64_abs(v43), v45)|base.F64_eq(base.F64_abs(v41), v45)|base.F64_eq(base.F64_abs(v42), v45) != 0 {
						v58 = v43
						v67 = m.G0
						v69 = v67 - int32(32)
						m.G0 = v69
						v71 = base.F64_abs(v40)
						v72 = base.F64_abs(v58)
						v75 = base.B2i32(base.Ui64(base.I64_reinterpret_f64(v71)) < base.Ui64(base.I64_reinterpret_f64(v72)))
						if base.Ui64(base.I64_reinterpret_f64(v71)) < base.Ui64(base.I64_reinterpret_f64(v72)) {
							v76 = v71
						} else {
							v76 = v72
						}
						v77 = base.I64_reinterpret_f64(v76)
						v79 = int64(base.Ui64(v77) >> (uint(int64(52)) % 64))
						if v79 == int64(2047) {
							v134 = v76
						} else {
							if base.Ui64(base.I64_reinterpret_f64(v71)) < base.Ui64(base.I64_reinterpret_f64(v72)) {
								v82 = v72
							} else {
								v82 = v71
							}
							if v77 == int64(0) {
								v134 = v82
							} else {
								v85 = base.I64_reinterpret_f64(v82)
								v87 = int64(base.Ui64(v85) >> (uint(int64(52)) % 64))
								if v87 == int64(2047) {
									v134 = v82
								} else {
									if int32(65) <= base.I32_wrap_i64(v87)-base.I32_wrap_i64(v79) {
										v134 = base.F64_add(v71, v72)
									} else {
										if base.Ui64(int64(6908521828386340864)) <= base.Ui64(v85) {
											v98 = float64(1.90109156629516e-211)
											v111 = base.F64_mul(v82, v98)
											v112 = base.F64_mul(v76, v98)
											v113 = float64(5.260135901548374e+210)
										} else {
											if base.Ui64(int64(2580562586483294207)) < base.Ui64(v77) {
												v111 = v82
												v112 = v76
												v113 = float64(1)
											} else {
												v106 = float64(5.260135901548374e+210)
												v111 = base.F64_mul(v82, v106)
												v112 = base.F64_mul(v76, v106)
												v113 = float64(1.90109156629516e-211)
											}
										}
										F_sq(m, v69+int32(24), v69+int32(16), v111)
										mBase = m.M
										F_sq(m, v69+int32(8), v69, v112)
										mBase = m.M
										v122 = *(*float64)(unsafe.Add(mBase, uint32(v69)))
										v123 = *(*float64)(unsafe.Add(mBase, uint32(v69)+16))
										v125 = *(*float64)(unsafe.Add(mBase, uint32(v69)+8))
										v127 = *(*float64)(unsafe.Add(mBase, uint32(v69)+24))
										v134 = base.F64_mul(v113, base.F64_sqrt(base.F64_add(base.F64_add(base.F64_add(v122, v123), v125), v127)))
									}
								}
							}
						}
						m.G0 = v69 + int32(32)
						m.G0 = v9 + int32(32)
						return base.I64_reinterpret_f64(v134)
					} else {
						v56 = F_float_overflow_error_ext(m, int32(0))
						mBase = m.M
						v57 = m.ExcPending
						if v57 != 0 {
							return int64(0)
						} else {
							v58 = v56
							v67 = m.G0
							v69 = v67 - int32(32)
							m.G0 = v69
							v71 = base.F64_abs(v40)
							v72 = base.F64_abs(v58)
							v75 = base.B2i32(base.Ui64(base.I64_reinterpret_f64(v71)) < base.Ui64(base.I64_reinterpret_f64(v72)))
							if base.Ui64(base.I64_reinterpret_f64(v71)) < base.Ui64(base.I64_reinterpret_f64(v72)) {
								v76 = v71
							} else {
								v76 = v72
							}
							v77 = base.I64_reinterpret_f64(v76)
							v79 = int64(base.Ui64(v77) >> (uint(int64(52)) % 64))
							if v79 == int64(2047) {
								v134 = v76
							} else {
								if base.Ui64(base.I64_reinterpret_f64(v71)) < base.Ui64(base.I64_reinterpret_f64(v72)) {
									v82 = v72
								} else {
									v82 = v71
								}
								if v77 == int64(0) {
									v134 = v82
								} else {
									v85 = base.I64_reinterpret_f64(v82)
									v87 = int64(base.Ui64(v85) >> (uint(int64(52)) % 64))
									if v87 == int64(2047) {
										v134 = v82
									} else {
										if int32(65) <= base.I32_wrap_i64(v87)-base.I32_wrap_i64(v79) {
											v134 = base.F64_add(v71, v72)
										} else {
											if base.Ui64(int64(6908521828386340864)) <= base.Ui64(v85) {
												v98 = float64(1.90109156629516e-211)
												v111 = base.F64_mul(v82, v98)
												v112 = base.F64_mul(v76, v98)
												v113 = float64(5.260135901548374e+210)
											} else {
												if base.Ui64(int64(2580562586483294207)) < base.Ui64(v77) {
													v111 = v82
													v112 = v76
													v113 = float64(1)
												} else {
													v106 = float64(5.260135901548374e+210)
													v111 = base.F64_mul(v82, v106)
													v112 = base.F64_mul(v76, v106)
													v113 = float64(1.90109156629516e-211)
												}
											}
											F_sq(m, v69+int32(24), v69+int32(16), v111)
											mBase = m.M
											F_sq(m, v69+int32(8), v69, v112)
											mBase = m.M
											v122 = *(*float64)(unsafe.Add(mBase, uint32(v69)))
											v123 = *(*float64)(unsafe.Add(mBase, uint32(v69)+16))
											v125 = *(*float64)(unsafe.Add(mBase, uint32(v69)+8))
											v127 = *(*float64)(unsafe.Add(mBase, uint32(v69)+24))
											v134 = base.F64_mul(v113, base.F64_sqrt(base.F64_add(base.F64_add(base.F64_add(v122, v123), v125), v127)))
										}
									}
								}
							}
							m.G0 = v69 + int32(32)
							m.G0 = v9 + int32(32)
							return base.I64_reinterpret_f64(v134)
						}
					}
				}
			}
		}
	}
}
func F_box_le(m *base.Module, l0 int32) int64 {
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
	var v9 float64
	_ = v9
	var v10 int32
	_ = v10
	v3 = *(*int32)(unsafe.Add(mBase, uint32(l0)+40))
	v4 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
	v5 = F_box_ar(m, v4)
	mBase = m.M
	v8 = m.ExcPending
	if v8 != 0 {
		return int64(0)
	} else {
		v9 = F_box_ar(m, v3)
		mBase = m.M
		v10 = m.ExcPending
		if v10 != 0 {
			return int64(0)
		} else {
			return base.I64_extend_i32_u(base.F64_le(v5, base.F64_add(v9, float64(1e-06))))
		}
	}
}
func F_box_left(m *base.Module, l0 int32) int64 {
	mBase := m.M
	_ = mBase
	var v2 int32
	_ = v2
	var v3 float64
	_ = v3
	var v4 int32
	_ = v4
	var v5 float64
	_ = v5
	v2 = *(*int32)(unsafe.Add(mBase, uint32(l0)+40))
	v3 = *(*float64)(unsafe.Add(mBase, uint32(v2)+16))
	v4 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
	v5 = *(*float64)(unsafe.Add(mBase, uint32(v4)))
	return base.I64_extend_i32_u(base.F64_gt(v3, base.F64_add(v5, float64(1e-06))))
}
func F_box_overlap(m *base.Module, l0 int32) int64 {
	mBase := m.M
	_ = mBase
	var v3 int64
	_ = v3
	var v4 int32
	_ = v4
	var v5 float64
	_ = v5
	var v6 int32
	_ = v6
	var v7 float64
	_ = v7
	var v13 float64
	_ = v13
	var v14 float64
	_ = v14
	var v20 float64
	_ = v20
	var v21 float64
	_ = v21
	var v27 float64
	_ = v27
	var v28 float64
	_ = v28
	var v33 int64
	_ = v33
	v3 = int64(0)
	v4 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
	v5 = *(*float64)(unsafe.Add(mBase, uint32(v4)+16))
	v6 = *(*int32)(unsafe.Add(mBase, uint32(l0)+40))
	v7 = *(*float64)(unsafe.Add(mBase, uint32(v6)))
	if base.F64_le(v5, base.F64_add(v7, float64(1e-06))) == int32(0) {
		v33 = v3
	} else {
		v13 = *(*float64)(unsafe.Add(mBase, uint32(v6)+16))
		v14 = *(*float64)(unsafe.Add(mBase, uint32(v4)))
		if base.F64_le(v13, base.F64_add(v14, float64(1e-06))) == int32(0) {
			v33 = v3
		} else {
			v20 = *(*float64)(unsafe.Add(mBase, uint32(v4)+24))
			v21 = *(*float64)(unsafe.Add(mBase, uint32(v6)+8))
			if base.F64_le(v20, base.F64_add(v21, float64(1e-06))) == int32(0) {
				v33 = v3
			} else {
				v27 = *(*float64)(unsafe.Add(mBase, uint32(v6)+24))
				v28 = *(*float64)(unsafe.Add(mBase, uint32(v4)+8))
				v33 = base.I64_extend_i32_u(base.F64_le(v27, base.F64_add(v28, float64(1e-06))))
			}
		}
	}
	return v33
}
func F_write_box(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32) int32 {
	mBase := m.M
	_ = mBase
	var v13 int32
	_ = v13
	var v15 int32
	_ = v15
	var v20 int32
	_ = v20
	var v21 int32
	_ = v21
	var v24 int32
	_ = v24
	var v28 int32
	_ = v28
	var v33 int32
	_ = v33
	var v40 float64
	_ = v40
	var v41 int32
	_ = v41
	var v45 int32
	_ = v45
	var v48 int32
	_ = v48
	var v52 int32
	_ = v52
	var v53 int32
	_ = v53
	var v54 int32
	_ = v54
	var v56 int32
	_ = v56
	var v60 int32
	_ = v60
	var v68 int32
	_ = v68
	var v70 int32
	_ = v70
	var v79 int32
	_ = v79
	var v83 float64
	_ = v83
	var v84 int32
	_ = v84
	var v88 int32
	_ = v88
	var v91 int32
	_ = v91
	var v96 int32
	_ = v96
	var v97 int32
	_ = v97
	var v98 int32
	_ = v98
	var v100 int32
	_ = v100
	var v104 int32
	_ = v104
	var v110 int32
	_ = v110
	var v121 int32
	_ = v121
	var v128 float64
	_ = v128
	var v129 int32
	_ = v129
	var v133 int32
	_ = v133
	var v136 int32
	_ = v136
	var v142 float64
	_ = v142
	var v144 int32
	_ = v144
	var v145 int32
	_ = v145
	var v147 int32
	_ = v147
	var v148 int32
	_ = v148
	var v150 int32
	_ = v150
	var v154 int32
	_ = v154
	var v156 int32
	_ = v156
	var v158 int32
	_ = v158
	var v162 int32
	_ = v162
	var v165 int32
	_ = v165
	var v173 int32
	_ = v173
	var v177 float64
	_ = v177
	var v178 int32
	_ = v178
	var v182 int32
	_ = v182
	var v185 int32
	_ = v185
	var v195 float64
	_ = v195
	var v197 int32
	_ = v197
	var v198 int32
	_ = v198
	var v199 int32
	_ = v199
	var v201 int32
	_ = v201
	var v205 int32
	_ = v205
	var v207 int32
	_ = v207
	var v223 int32
	_ = v223
	var v241 int32
	_ = v241
	v13 = m.G0
	v15 = v13 - int32(16)
	m.G0 = v15
	v20 = l0<<(uint(int32(4))%32) | int32(8)
	v21 = F_palloc0(m, v20)
	mBase = m.M
	v24 = m.ExcPending
	if v24 != 0 {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	return int32(0)
L2:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v21))) = v20 << (uint(int32(2)) % 32)
	v28 = *(*int32)(unsafe.Add(mBase, uint32(v21)+4))
	*(*int32)(unsafe.Add(mBase, uint32(v21)+4)) = v28&int32(-2147483648) | l0
	v33 = int32(1)
	if l0 <= int32(0) {
		goto L5
	} else {
		goto L6
	}
L3:
	;
	m.G0 = v15 + int32(16)
	return v241
L4:
	;
	v53 = int32(44)
	v54 = F___strchrnul(m, l1, v53)
	mBase = m.M
	v56 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v54))))
	if v56 == v53 {
		goto L13
	} else {
		goto L14
	}
L5:
	;
	v52 = int32(0)
	goto L4
L6:
	;
	goto L7
L7:
	;
	v40 = F_float8in_internal(m, l1, v15+int32(12), int32(_a_F_write_box_0), l1, l4)
	mBase = m.M
	v41 = m.ExcPending
	if v41 != 0 {
		goto L1
	} else {
		goto L8
	}
L8:
	;
	*(*float64)(unsafe.Add(mBase, uint32(v21)+8)) = v40
	if l4 == int32(0) {
		v52 = v33
		goto L4
	} else {
		goto L9
	}
L9:
	;
	v45 = *(*int32)(unsafe.Add(mBase, uint32(l4)))
	if v45 != int32(453) {
		v52 = v33
		goto L4
	} else {
		goto L10
	}
L10:
	;
	v48 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l4)+4)))
	if v48 == int32(0) {
		v52 = v33
		goto L4
	} else {
		goto L11
	}
L11:
	;
	v241 = int32(0)
	goto L3
L12:
	;
	if v60 != 0 {
		goto L16
	} else {
		goto L17
	}
L13:
	;
	v60 = v54
	goto L15
L14:
	;
	v60 = int32(0)
	goto L15
L15:
	;
	goto L12
L16:
	;
	v68 = v52
	v70 = v60
	goto L19
L17:
	;
	v110 = v52
	goto L18
L18:
	;
	if int32(0) < l0 {
		goto L31
	} else {
		goto L32
	}
L19:
	;
	v79 = v70 + int32(1)
	v83 = F_float8in_internal(m, v79, v15+int32(12), int32(_a_F_write_box_0), l1, l4)
	mBase = m.M
	v84 = m.ExcPending
	if v84 != 0 {
		goto L1
	} else {
		goto L21
	}
L20:
	;
	v110 = v96
	goto L18
L21:
	;
	*(*float64)(unsafe.Add(mBase, uint32(v21+int32(8)+v68<<(uint(int32(3))%32)))) = v83
	if l4 == int32(0) {
		goto L22
	} else {
		goto L23
	}
L22:
	;
	v96 = v68 + int32(1)
	v97 = int32(44)
	v98 = F___strchrnul(m, v79, v97)
	mBase = m.M
	v100 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v98))))
	if v100 == v97 {
		goto L27
	} else {
		goto L28
	}
L23:
	;
	v88 = *(*int32)(unsafe.Add(mBase, uint32(l4)))
	if v88 != int32(453) {
		goto L22
	} else {
		goto L24
	}
L24:
	;
	v91 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l4)+4)))
	if v91 == int32(0) {
		goto L22
	} else {
		goto L25
	}
L25:
	;
	v241 = int32(0)
	goto L3
L26:
	;
	if v104 != 0 {
		v68 = v96
		v70 = v104
		goto L19
	} else {
		goto L30
	}
L27:
	;
	v104 = v98
	goto L29
L28:
	;
	v104 = int32(0)
	goto L29
L29:
	;
	goto L26
L30:
	;
	goto L20
L31:
	;
	v121 = v21 + int32(8)
	v128 = F_float8in_internal(m, l2, v15+int32(12), int32(_a_F_write_box_0), l2, l4)
	mBase = m.M
	v129 = m.ExcPending
	if v129 != 0 {
		goto L1
	} else {
		goto L34
	}
L32:
	;
	v144 = int32(1)
	v145 = v110
	goto L33
L33:
	;
	v147 = int32(44)
	v148 = F___strchrnul(m, l2, v147)
	mBase = m.M
	v150 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v148))))
	if v150 == v147 {
		goto L40
	} else {
		goto L41
	}
L34:
	;
	*(*float64)(unsafe.Add(mBase, uint32(v121+v110<<(uint(int32(3))%32)))) = v128
	if l4 == int32(0) {
		goto L35
	} else {
		goto L36
	}
L35:
	;
	v142 = *(*float64)(unsafe.Add(mBase, uint32(v121)))
	v144 = base.F64_eq(v128, v142)
	v145 = v110 + int32(1)
	goto L33
L36:
	;
	v133 = *(*int32)(unsafe.Add(mBase, uint32(l4)))
	if v133 != int32(453) {
		goto L35
	} else {
		goto L37
	}
L37:
	;
	v136 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l4)+4)))
	if v136 == int32(0) {
		goto L35
	} else {
		goto L38
	}
L38:
	;
	v241 = int32(0)
	goto L3
L39:
	;
	if v154 != 0 {
		goto L43
	} else {
		goto L44
	}
L40:
	;
	v154 = v148
	goto L42
L41:
	;
	v154 = int32(0)
	goto L42
L42:
	;
	goto L39
L43:
	;
	v156 = v21 + int32(8)
	v158 = v144
	v162 = v145
	v165 = v154
	goto L46
L44:
	;
	v207 = v144
	goto L45
L45:
	;
	if v207 != 0 {
		goto L58
	} else {
		goto L59
	}
L46:
	;
	v173 = v165 + int32(1)
	v177 = F_float8in_internal(m, v173, v15+int32(12), int32(_a_F_write_box_0), l2, l4)
	mBase = m.M
	v178 = m.ExcPending
	if v178 != 0 {
		goto L1
	} else {
		goto L48
	}
L47:
	;
	v207 = v197
	goto L45
L48:
	;
	*(*float64)(unsafe.Add(mBase, uint32(v156+v162<<(uint(int32(3))%32)))) = v177
	if l4 == int32(0) {
		goto L49
	} else {
		goto L50
	}
L49:
	;
	v195 = *(*float64)(unsafe.Add(mBase, uint32(v156+(v162-l0)<<(uint(int32(3))%32))))
	v197 = v158 & base.F64_eq(v177, v195)
	v198 = int32(44)
	v199 = F___strchrnul(m, v173, v198)
	mBase = m.M
	v201 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v199))))
	if v201 == v198 {
		goto L54
	} else {
		goto L55
	}
L50:
	;
	v182 = *(*int32)(unsafe.Add(mBase, uint32(l4)))
	if v182 != int32(453) {
		goto L49
	} else {
		goto L51
	}
L51:
	;
	v185 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l4)+4)))
	if v185 == int32(0) {
		goto L49
	} else {
		goto L52
	}
L52:
	;
	v241 = int32(0)
	goto L3
L53:
	;
	if v205 != 0 {
		v158 = v197
		v162 = v162 + int32(1)
		v165 = v205
		goto L46
	} else {
		goto L57
	}
L54:
	;
	v205 = v199
	goto L56
L55:
	;
	v205 = int32(0)
	goto L56
L56:
	;
	goto L53
L57:
	;
	goto L47
L58:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v21))) = l0<<(uint(int32(5))%32) + int32(32)
	v223 = *(*int32)(unsafe.Add(mBase, uint32(v21)+4))
	*(*int32)(unsafe.Add(mBase, uint32(v21)+4)) = v223 | int32(-2147483648)
	goto L60
L59:
	;
	goto L60
L60:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l3))) = v21
	v241 = int32(1)
	goto L3
}
