package p5

import (
	base "github.com/shibukawa/pgmem/internal/aot/pgaot/base"
	"math"
	"unsafe"
)

func F_box_above(m *base.Module, l0 int32) int64 {
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
	v2 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
	v3 = *(*float64)(unsafe.Add(mBase, uint32(v2)+24))
	v4 = *(*int32)(unsafe.Add(mBase, uint32(l0)+40))
	v5 = *(*float64)(unsafe.Add(mBase, uint32(v4)+8))
	return base.I64_extend_i32_u(base.F64_gt(v3, base.F64_add(v5, float64(1e-06))))
}
func F_box_add(m *base.Module, l0 int32) int64 {
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
	var v56 float64
	_ = v56
	var v57 float64
	_ = v57
	var v59 float64
	_ = v59
	var v72 float64
	_ = v72
	var v73 int32
	_ = v73
	var v74 float64
	_ = v74
	var v75 float64
	_ = v75
	var v76 float64
	_ = v76
	var v77 float64
	_ = v77
	var v79 float64
	_ = v79
	var v90 float64
	_ = v90
	var v91 int32
	_ = v91
	var v92 float64
	_ = v92
	v8 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
	v9 = *(*int32)(unsafe.Add(mBase, uint32(l0)+40))
	v11 = F_palloc(m, int32(32))
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
					v56 = *(*float64)(unsafe.Add(mBase, uint32(v9)))
					v57 = base.F64_add(v55, v56)
					v59 = math.Float64frombits(uint64(0x7ff0000000000000))
					if base.F64_ne(base.F64_abs(v57), v59)|base.F64_eq(base.F64_abs(v55), v59)|base.F64_eq(base.F64_abs(v56), v59) == int32(0) {
						v72 = F_float_overflow_error_ext(m, int32(0))
						mBase = m.M
						v73 = m.ExcPending
						if v73 != 0 {
							return int64(0)
						} else {
							v74 = v72
							v75 = *(*float64)(unsafe.Add(mBase, uint32(v8)+24))
							v76 = *(*float64)(unsafe.Add(mBase, uint32(v9)+8))
							v77 = base.F64_add(v75, v76)
							v79 = math.Float64frombits(uint64(0x7ff0000000000000))
							if base.F64_ne(base.F64_abs(v77), v79)|base.F64_eq(base.F64_abs(v75), v79)|base.F64_eq(base.F64_abs(v76), v79) != 0 {
								v92 = v77
								*(*float64)(unsafe.Add(mBase, uint32(v11)+24)) = v92
								*(*float64)(unsafe.Add(mBase, uint32(v11)+16)) = v74
								return base.I64_extend_i32_u(v11)
							} else {
								v90 = F_float_overflow_error_ext(m, int32(0))
								mBase = m.M
								v91 = m.ExcPending
								if v91 != 0 {
									return int64(0)
								} else {
									v92 = v90
									*(*float64)(unsafe.Add(mBase, uint32(v11)+24)) = v92
									*(*float64)(unsafe.Add(mBase, uint32(v11)+16)) = v74
									return base.I64_extend_i32_u(v11)
								}
							}
						}
					} else {
						v74 = v57
						v75 = *(*float64)(unsafe.Add(mBase, uint32(v8)+24))
						v76 = *(*float64)(unsafe.Add(mBase, uint32(v9)+8))
						v77 = base.F64_add(v75, v76)
						v79 = math.Float64frombits(uint64(0x7ff0000000000000))
						if base.F64_ne(base.F64_abs(v77), v79)|base.F64_eq(base.F64_abs(v75), v79)|base.F64_eq(base.F64_abs(v76), v79) != 0 {
							v92 = v77
							*(*float64)(unsafe.Add(mBase, uint32(v11)+24)) = v92
							*(*float64)(unsafe.Add(mBase, uint32(v11)+16)) = v74
							return base.I64_extend_i32_u(v11)
						} else {
							v90 = F_float_overflow_error_ext(m, int32(0))
							mBase = m.M
							v91 = m.ExcPending
							if v91 != 0 {
								return int64(0)
							} else {
								v92 = v90
								*(*float64)(unsafe.Add(mBase, uint32(v11)+24)) = v92
								*(*float64)(unsafe.Add(mBase, uint32(v11)+16)) = v74
								return base.I64_extend_i32_u(v11)
							}
						}
					}
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
						v56 = *(*float64)(unsafe.Add(mBase, uint32(v9)))
						v57 = base.F64_add(v55, v56)
						v59 = math.Float64frombits(uint64(0x7ff0000000000000))
						if base.F64_ne(base.F64_abs(v57), v59)|base.F64_eq(base.F64_abs(v55), v59)|base.F64_eq(base.F64_abs(v56), v59) == int32(0) {
							v72 = F_float_overflow_error_ext(m, int32(0))
							mBase = m.M
							v73 = m.ExcPending
							if v73 != 0 {
								return int64(0)
							} else {
								v74 = v72
								v75 = *(*float64)(unsafe.Add(mBase, uint32(v8)+24))
								v76 = *(*float64)(unsafe.Add(mBase, uint32(v9)+8))
								v77 = base.F64_add(v75, v76)
								v79 = math.Float64frombits(uint64(0x7ff0000000000000))
								if base.F64_ne(base.F64_abs(v77), v79)|base.F64_eq(base.F64_abs(v75), v79)|base.F64_eq(base.F64_abs(v76), v79) != 0 {
									v92 = v77
									*(*float64)(unsafe.Add(mBase, uint32(v11)+24)) = v92
									*(*float64)(unsafe.Add(mBase, uint32(v11)+16)) = v74
									return base.I64_extend_i32_u(v11)
								} else {
									v90 = F_float_overflow_error_ext(m, int32(0))
									mBase = m.M
									v91 = m.ExcPending
									if v91 != 0 {
										return int64(0)
									} else {
										v92 = v90
										*(*float64)(unsafe.Add(mBase, uint32(v11)+24)) = v92
										*(*float64)(unsafe.Add(mBase, uint32(v11)+16)) = v74
										return base.I64_extend_i32_u(v11)
									}
								}
							}
						} else {
							v74 = v57
							v75 = *(*float64)(unsafe.Add(mBase, uint32(v8)+24))
							v76 = *(*float64)(unsafe.Add(mBase, uint32(v9)+8))
							v77 = base.F64_add(v75, v76)
							v79 = math.Float64frombits(uint64(0x7ff0000000000000))
							if base.F64_ne(base.F64_abs(v77), v79)|base.F64_eq(base.F64_abs(v75), v79)|base.F64_eq(base.F64_abs(v76), v79) != 0 {
								v92 = v77
								*(*float64)(unsafe.Add(mBase, uint32(v11)+24)) = v92
								*(*float64)(unsafe.Add(mBase, uint32(v11)+16)) = v74
								return base.I64_extend_i32_u(v11)
							} else {
								v90 = F_float_overflow_error_ext(m, int32(0))
								mBase = m.M
								v91 = m.ExcPending
								if v91 != 0 {
									return int64(0)
								} else {
									v92 = v90
									*(*float64)(unsafe.Add(mBase, uint32(v11)+24)) = v92
									*(*float64)(unsafe.Add(mBase, uint32(v11)+16)) = v74
									return base.I64_extend_i32_u(v11)
								}
							}
						}
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
				v56 = *(*float64)(unsafe.Add(mBase, uint32(v9)))
				v57 = base.F64_add(v55, v56)
				v59 = math.Float64frombits(uint64(0x7ff0000000000000))
				if base.F64_ne(base.F64_abs(v57), v59)|base.F64_eq(base.F64_abs(v55), v59)|base.F64_eq(base.F64_abs(v56), v59) == int32(0) {
					v72 = F_float_overflow_error_ext(m, int32(0))
					mBase = m.M
					v73 = m.ExcPending
					if v73 != 0 {
						return int64(0)
					} else {
						v74 = v72
						v75 = *(*float64)(unsafe.Add(mBase, uint32(v8)+24))
						v76 = *(*float64)(unsafe.Add(mBase, uint32(v9)+8))
						v77 = base.F64_add(v75, v76)
						v79 = math.Float64frombits(uint64(0x7ff0000000000000))
						if base.F64_ne(base.F64_abs(v77), v79)|base.F64_eq(base.F64_abs(v75), v79)|base.F64_eq(base.F64_abs(v76), v79) != 0 {
							v92 = v77
							*(*float64)(unsafe.Add(mBase, uint32(v11)+24)) = v92
							*(*float64)(unsafe.Add(mBase, uint32(v11)+16)) = v74
							return base.I64_extend_i32_u(v11)
						} else {
							v90 = F_float_overflow_error_ext(m, int32(0))
							mBase = m.M
							v91 = m.ExcPending
							if v91 != 0 {
								return int64(0)
							} else {
								v92 = v90
								*(*float64)(unsafe.Add(mBase, uint32(v11)+24)) = v92
								*(*float64)(unsafe.Add(mBase, uint32(v11)+16)) = v74
								return base.I64_extend_i32_u(v11)
							}
						}
					}
				} else {
					v74 = v57
					v75 = *(*float64)(unsafe.Add(mBase, uint32(v8)+24))
					v76 = *(*float64)(unsafe.Add(mBase, uint32(v9)+8))
					v77 = base.F64_add(v75, v76)
					v79 = math.Float64frombits(uint64(0x7ff0000000000000))
					if base.F64_ne(base.F64_abs(v77), v79)|base.F64_eq(base.F64_abs(v75), v79)|base.F64_eq(base.F64_abs(v76), v79) != 0 {
						v92 = v77
						*(*float64)(unsafe.Add(mBase, uint32(v11)+24)) = v92
						*(*float64)(unsafe.Add(mBase, uint32(v11)+16)) = v74
						return base.I64_extend_i32_u(v11)
					} else {
						v90 = F_float_overflow_error_ext(m, int32(0))
						mBase = m.M
						v91 = m.ExcPending
						if v91 != 0 {
							return int64(0)
						} else {
							v92 = v90
							*(*float64)(unsafe.Add(mBase, uint32(v11)+24)) = v92
							*(*float64)(unsafe.Add(mBase, uint32(v11)+16)) = v74
							return base.I64_extend_i32_u(v11)
						}
					}
				}
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
					v56 = *(*float64)(unsafe.Add(mBase, uint32(v9)))
					v57 = base.F64_add(v55, v56)
					v59 = math.Float64frombits(uint64(0x7ff0000000000000))
					if base.F64_ne(base.F64_abs(v57), v59)|base.F64_eq(base.F64_abs(v55), v59)|base.F64_eq(base.F64_abs(v56), v59) == int32(0) {
						v72 = F_float_overflow_error_ext(m, int32(0))
						mBase = m.M
						v73 = m.ExcPending
						if v73 != 0 {
							return int64(0)
						} else {
							v74 = v72
							v75 = *(*float64)(unsafe.Add(mBase, uint32(v8)+24))
							v76 = *(*float64)(unsafe.Add(mBase, uint32(v9)+8))
							v77 = base.F64_add(v75, v76)
							v79 = math.Float64frombits(uint64(0x7ff0000000000000))
							if base.F64_ne(base.F64_abs(v77), v79)|base.F64_eq(base.F64_abs(v75), v79)|base.F64_eq(base.F64_abs(v76), v79) != 0 {
								v92 = v77
								*(*float64)(unsafe.Add(mBase, uint32(v11)+24)) = v92
								*(*float64)(unsafe.Add(mBase, uint32(v11)+16)) = v74
								return base.I64_extend_i32_u(v11)
							} else {
								v90 = F_float_overflow_error_ext(m, int32(0))
								mBase = m.M
								v91 = m.ExcPending
								if v91 != 0 {
									return int64(0)
								} else {
									v92 = v90
									*(*float64)(unsafe.Add(mBase, uint32(v11)+24)) = v92
									*(*float64)(unsafe.Add(mBase, uint32(v11)+16)) = v74
									return base.I64_extend_i32_u(v11)
								}
							}
						}
					} else {
						v74 = v57
						v75 = *(*float64)(unsafe.Add(mBase, uint32(v8)+24))
						v76 = *(*float64)(unsafe.Add(mBase, uint32(v9)+8))
						v77 = base.F64_add(v75, v76)
						v79 = math.Float64frombits(uint64(0x7ff0000000000000))
						if base.F64_ne(base.F64_abs(v77), v79)|base.F64_eq(base.F64_abs(v75), v79)|base.F64_eq(base.F64_abs(v76), v79) != 0 {
							v92 = v77
							*(*float64)(unsafe.Add(mBase, uint32(v11)+24)) = v92
							*(*float64)(unsafe.Add(mBase, uint32(v11)+16)) = v74
							return base.I64_extend_i32_u(v11)
						} else {
							v90 = F_float_overflow_error_ext(m, int32(0))
							mBase = m.M
							v91 = m.ExcPending
							if v91 != 0 {
								return int64(0)
							} else {
								v92 = v90
								*(*float64)(unsafe.Add(mBase, uint32(v11)+24)) = v92
								*(*float64)(unsafe.Add(mBase, uint32(v11)+16)) = v74
								return base.I64_extend_i32_u(v11)
							}
						}
					}
				}
			}
		}
	}
}
func F_box_below(m *base.Module, l0 int32) int64 {
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
	v3 = *(*float64)(unsafe.Add(mBase, uint32(v2)+24))
	v4 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
	v5 = *(*float64)(unsafe.Add(mBase, uint32(v4)+8))
	return base.I64_extend_i32_u(base.F64_gt(v3, base.F64_add(v5, float64(1e-06))))
}
func F_box_contained(m *base.Module, l0 int32) int64 {
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
	v5 = *(*float64)(unsafe.Add(mBase, uint32(v4)))
	v6 = *(*int32)(unsafe.Add(mBase, uint32(l0)+40))
	v7 = *(*float64)(unsafe.Add(mBase, uint32(v6)))
	if base.F64_le(v5, base.F64_add(v7, float64(1e-06))) == int32(0) {
		v33 = v3
	} else {
		v13 = *(*float64)(unsafe.Add(mBase, uint32(v6)+16))
		v14 = *(*float64)(unsafe.Add(mBase, uint32(v4)+16))
		if base.F64_le(v13, base.F64_add(v14, float64(1e-06))) == int32(0) {
			v33 = v3
		} else {
			v20 = *(*float64)(unsafe.Add(mBase, uint32(v4)+8))
			v21 = *(*float64)(unsafe.Add(mBase, uint32(v6)+8))
			if base.F64_le(v20, base.F64_add(v21, float64(1e-06))) == int32(0) {
				v33 = v3
			} else {
				v27 = *(*float64)(unsafe.Add(mBase, uint32(v6)+24))
				v28 = *(*float64)(unsafe.Add(mBase, uint32(v4)+24))
				v33 = base.I64_extend_i32_u(base.F64_le(v27, base.F64_add(v28, float64(1e-06))))
			}
		}
	}
	return v33
}
func F_box_gt(m *base.Module, l0 int32) int64 {
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
			return base.I64_extend_i32_u(base.F64_gt(v5, base.F64_add(v9, float64(1e-06))))
		}
	}
}
func F_box_height(m *base.Module, l0 int32) int64 {
	mBase := m.M
	_ = mBase
	var v5 int32
	_ = v5
	var v6 float64
	_ = v6
	var v7 float64
	_ = v7
	var v8 float64
	_ = v8
	var v10 float64
	_ = v10
	var v21 float64
	_ = v21
	var v24 int32
	_ = v24
	var v25 float64
	_ = v25
	v5 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
	v6 = *(*float64)(unsafe.Add(mBase, uint32(v5)+8))
	v7 = *(*float64)(unsafe.Add(mBase, uint32(v5)+24))
	v8 = base.F64_sub(v6, v7)
	v10 = math.Float64frombits(uint64(0x7ff0000000000000))
	if base.F64_ne(base.F64_abs(v8), v10)|base.F64_eq(base.F64_abs(v6), v10)|base.F64_eq(base.F64_abs(v7), v10) != 0 {
		v25 = v8
		return base.I64_reinterpret_f64(v25)
	} else {
		v21 = F_float_overflow_error_ext(m, int32(0))
		mBase = m.M
		v24 = m.ExcPending
		if v24 != 0 {
			return int64(0)
		} else {
			v25 = v21
			return base.I64_reinterpret_f64(v25)
		}
	}
}
func F_box_in(m *base.Module, l0 int32) int64 {
	mBase := m.M
	_ = mBase
	var v9 int32
	_ = v9
	var v11 int32
	_ = v11
	var v13 int32
	_ = v13
	var v14 int32
	_ = v14
	var v18 int32
	_ = v18
	var v21 int32
	_ = v21
	var v26 int32
	_ = v26
	var v27 int32
	_ = v27
	var v30 int32
	_ = v30
	var v33 float64
	_ = v33
	var v39 float64
	_ = v39
	var v52 float64
	_ = v52
	var v58 float64
	_ = v58
	var v74 int64
	_ = v74
	v9 = m.G0
	v11 = v9 - int32(16)
	m.G0 = v11
	v13 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v14 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
	v18 = F_palloc(m, int32(32))
	mBase = m.M
	v21 = m.ExcPending
	if v21 != 0 {
		return int64(0)
	} else {
		v26 = F_path_decode(m, v14, int32(0), int32(2), v18, v11+int32(15), int32(0), int32(_a_F_box_in_0), v14, v13)
		mBase = m.M
		v27 = m.ExcPending
		if v27 != 0 {
			return int64(0)
		} else {
			if v26 == int32(0) {
				v30 = int32(1)
				*(*uint8)(unsafe.Add(mBase, uint32(l0)+16)) = uint8(v30)
				v74 = int64(0)
			} else {
				v33 = *(*float64)(unsafe.Add(mBase, uint32(v18)))
				if base.Ui64(int64(9218868437227405312)) < base.Ui64(base.I64_reinterpret_f64(v33)&int64(9223372036854775807)) {
				} else {
					v39 = *(*float64)(unsafe.Add(mBase, uint32(v18)+16))
					if base.B2i32(base.F64_lt(v33, v39) == int32(0))&base.B2i32(base.Ui64(base.I64_reinterpret_f64(v39)&int64(9223372036854775807)) < base.Ui64(int64(9218868437227405313))) != 0 {
					} else {
						*(*float64)(unsafe.Add(mBase, uint32(v18)+16)) = v33
						*(*float64)(unsafe.Add(mBase, uint32(v18))) = v39
					}
				}
				v52 = *(*float64)(unsafe.Add(mBase, uint32(v18)+8))
				if base.Ui64(int64(9218868437227405312)) < base.Ui64(base.I64_reinterpret_f64(v52)&int64(9223372036854775807)) {
				} else {
					v58 = *(*float64)(unsafe.Add(mBase, uint32(v18)+24))
					if base.B2i32(base.F64_lt(v52, v58) == int32(0))&base.B2i32(base.Ui64(base.I64_reinterpret_f64(v58)&int64(9223372036854775807)) < base.Ui64(int64(9218868437227405313))) != 0 {
					} else {
						*(*float64)(unsafe.Add(mBase, uint32(v18)+24)) = v52
						*(*float64)(unsafe.Add(mBase, uint32(v18)+8)) = v58
					}
				}
				v74 = base.I64_extend_i32_u(v18)
			}
			m.G0 = v11 + int32(16)
			return v74
		}
	}
}
func F_box_intersect(m *base.Module, l0 int32) int64 {
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
	var v15 float64
	_ = v15
	var v16 float64
	_ = v16
	var v22 float64
	_ = v22
	var v23 float64
	_ = v23
	var v29 float64
	_ = v29
	var v30 float64
	_ = v30
	var v34 int32
	_ = v34
	var v39 int32
	_ = v39
	var v42 int32
	_ = v42
	var v43 float64
	_ = v43
	var v44 float64
	_ = v44
	var v55 float64
	_ = v55
	var v57 float64
	_ = v57
	var v58 float64
	_ = v58
	var v60 float64
	_ = v60
	var v66 float64
	_ = v66
	var v72 float64
	_ = v72
	var v74 float64
	_ = v74
	var v76 float64
	_ = v76
	var v78 float64
	_ = v78
	var v79 float64
	_ = v79
	var v90 float64
	_ = v90
	var v92 float64
	_ = v92
	var v93 float64
	_ = v93
	var v95 float64
	_ = v95
	var v101 float64
	_ = v101
	var v107 float64
	_ = v107
	var v109 float64
	_ = v109
	var v111 float64
	_ = v111
	v6 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
	v7 = *(*float64)(unsafe.Add(mBase, uint32(v6)+16))
	v8 = *(*int32)(unsafe.Add(mBase, uint32(l0)+40))
	v9 = *(*float64)(unsafe.Add(mBase, uint32(v8)))
	if base.F64_le(v7, base.F64_add(v9, float64(1e-06))) == int32(0) {
		v34 = int32(1)
		*(*uint8)(unsafe.Add(mBase, uint32(l0)+16)) = uint8(v34)
		return int64(0)
	} else {
		v15 = *(*float64)(unsafe.Add(mBase, uint32(v8)+16))
		v16 = *(*float64)(unsafe.Add(mBase, uint32(v6)))
		if base.F64_le(v15, base.F64_add(v16, float64(1e-06))) == int32(0) {
			v34 = int32(1)
			*(*uint8)(unsafe.Add(mBase, uint32(l0)+16)) = uint8(v34)
			return int64(0)
		} else {
			v22 = *(*float64)(unsafe.Add(mBase, uint32(v6)+24))
			v23 = *(*float64)(unsafe.Add(mBase, uint32(v8)+8))
			if base.F64_le(v22, base.F64_add(v23, float64(1e-06))) == int32(0) {
				v34 = int32(1)
				*(*uint8)(unsafe.Add(mBase, uint32(l0)+16)) = uint8(v34)
				return int64(0)
			} else {
				v29 = *(*float64)(unsafe.Add(mBase, uint32(v8)+24))
				v30 = *(*float64)(unsafe.Add(mBase, uint32(v6)+8))
				if base.F64_le(v29, base.F64_add(v30, float64(1e-06))) != 0 {
					v39 = F_palloc(m, int32(32))
					mBase = m.M
					v42 = m.ExcPending
					if v42 != 0 {
						return int64(0)
					} else {
						v43 = *(*float64)(unsafe.Add(mBase, uint32(v8)))
						v44 = *(*float64)(unsafe.Add(mBase, uint32(v6)))
						if base.Ui64(base.I64_reinterpret_f64(v44)&int64(9223372036854775807)) <= base.Ui64(int64(9218868437227405312)) {
							if base.Ui64(int64(9218868437227405312)) < base.Ui64(base.I64_reinterpret_f64(v43)&int64(9223372036854775807)) {
								v55 = v44
							} else {
								v55 = v43
							}
							if base.F64_gt(v43, v44) != 0 {
								v57 = v44
							} else {
								v57 = v55
							}
							v58 = v57
						} else {
							v58 = v43
						}
						*(*float64)(unsafe.Add(mBase, uint32(v39))) = v58
						v60 = *(*float64)(unsafe.Add(mBase, uint32(v8)+16))
						if base.Ui64(base.I64_reinterpret_f64(v60)&int64(9223372036854775807)) <= base.Ui64(int64(9218868437227405312)) {
							v66 = *(*float64)(unsafe.Add(mBase, uint32(v6)+16))
							if base.Ui64(int64(9218868437227405312)) < base.Ui64(base.I64_reinterpret_f64(v66)&int64(9223372036854775807)) {
								v72 = v66
							} else {
								v72 = v60
							}
							if base.F64_lt(v60, v66) != 0 {
								v74 = v66
							} else {
								v74 = v72
							}
							v76 = v74
						} else {
							v76 = v60
						}
						*(*float64)(unsafe.Add(mBase, uint32(v39)+16)) = v76
						v78 = *(*float64)(unsafe.Add(mBase, uint32(v8)+8))
						v79 = *(*float64)(unsafe.Add(mBase, uint32(v6)+8))
						if base.Ui64(base.I64_reinterpret_f64(v79)&int64(9223372036854775807)) <= base.Ui64(int64(9218868437227405312)) {
							if base.Ui64(int64(9218868437227405312)) < base.Ui64(base.I64_reinterpret_f64(v78)&int64(9223372036854775807)) {
								v90 = v79
							} else {
								v90 = v78
							}
							if base.F64_gt(v78, v79) != 0 {
								v92 = v79
							} else {
								v92 = v90
							}
							v93 = v92
						} else {
							v93 = v78
						}
						*(*float64)(unsafe.Add(mBase, uint32(v39)+8)) = v93
						v95 = *(*float64)(unsafe.Add(mBase, uint32(v8)+24))
						if base.Ui64(base.I64_reinterpret_f64(v95)&int64(9223372036854775807)) <= base.Ui64(int64(9218868437227405312)) {
							v101 = *(*float64)(unsafe.Add(mBase, uint32(v6)+24))
							if base.Ui64(int64(9218868437227405312)) < base.Ui64(base.I64_reinterpret_f64(v101)&int64(9223372036854775807)) {
								v107 = v101
							} else {
								v107 = v95
							}
							if base.F64_lt(v95, v101) != 0 {
								v109 = v101
							} else {
								v109 = v107
							}
							v111 = v109
						} else {
							v111 = v95
						}
						*(*float64)(unsafe.Add(mBase, uint32(v39)+24)) = v111
						return base.I64_extend_i32_u(v39)
					}
				} else {
					v34 = int32(1)
					*(*uint8)(unsafe.Add(mBase, uint32(l0)+16)) = uint8(v34)
					return int64(0)
				}
			}
		}
	}
}
func F_box_overleft(m *base.Module, l0 int32) int64 {
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
	v2 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
	v3 = *(*float64)(unsafe.Add(mBase, uint32(v2)))
	v4 = *(*int32)(unsafe.Add(mBase, uint32(l0)+40))
	v5 = *(*float64)(unsafe.Add(mBase, uint32(v4)))
	return base.I64_extend_i32_u(base.F64_le(v3, base.F64_add(v5, float64(1e-06))))
}
