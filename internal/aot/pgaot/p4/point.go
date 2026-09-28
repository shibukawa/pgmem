package p4

import (
	base "github.com/shibukawa/pgmem/internal/aot/pgaot/base"
	"math"
	"unsafe"
)

func F_point_add(m *base.Module, l0 int32) int64 {
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
	v8 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
	v9 = *(*int32)(unsafe.Add(mBase, uint32(l0)+40))
	v11 = F_palloc(m, int32(16))
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
					return base.I64_extend_i32_u(v11)
				}
			}
		}
	}
}
func F_point_div(m *base.Module, l0 int32) int64 {
	mBase := m.M
	_ = mBase
	var v4 int32
	_ = v4
	var v5 int32
	_ = v5
	var v7 int32
	_ = v7
	var v10 int32
	_ = v10
	var v12 int32
	_ = v12
	v4 = *(*int32)(unsafe.Add(mBase, uint32(l0)+40))
	v5 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
	v7 = F_palloc(m, int32(16))
	mBase = m.M
	v10 = m.ExcPending
	if v10 != 0 {
		return int64(0)
	} else {
		F_point_div_point(m, v7, v5, v4)
		mBase = m.M
		v12 = m.ExcPending
		if v12 != 0 {
			return int64(0)
		} else {
			return base.I64_extend_i32_u(v7)
		}
	}
}
func F_point_sl(m *base.Module, l0 int32, l1 int32) float64 {
	mBase := m.M
	_ = mBase
	var v11 float64
	_ = v11
	var v12 float64
	_ = v12
	var v13 float64
	_ = v13
	var v15 float64
	_ = v15
	var v19 float64
	_ = v19
	var v20 float64
	_ = v20
	var v21 float64
	_ = v21
	var v23 float64
	_ = v23
	var v24 float64
	_ = v24
	var v28 float64
	_ = v28
	var v40 float64
	_ = v40
	var v43 int32
	_ = v43
	var v44 float64
	_ = v44
	var v45 float64
	_ = v45
	var v47 float64
	_ = v47
	var v48 float64
	_ = v48
	var v49 float64
	_ = v49
	var v50 float64
	_ = v50
	var v52 float64
	_ = v52
	var v65 float64
	_ = v65
	var v66 int32
	_ = v66
	var v67 float64
	_ = v67
	var v79 float64
	_ = v79
	var v80 int32
	_ = v80
	var v83 float64
	_ = v83
	var v85 float64
	_ = v85
	var v93 float64
	_ = v93
	var v94 int32
	_ = v94
	var v96 float64
	_ = v96
	var v106 float64
	_ = v106
	var v107 int32
	_ = v107
	var v109 float64
	_ = v109
	v11 = math.Float64frombits(uint64(0x7ff0000000000000))
	v12 = *(*float64)(unsafe.Add(mBase, uint32(l0)))
	v13 = *(*float64)(unsafe.Add(mBase, uint32(l1)))
	if base.F64_eq(v12, v13) != 0 {
		v109 = v11
		return v109
	} else {
		v15 = base.F64_sub(v12, v13)
		if base.F64_le(base.F64_abs(v15), float64(1e-06)) != 0 {
			v109 = v11
			return v109
		} else {
			v19 = float64(0)
			v20 = *(*float64)(unsafe.Add(mBase, uint32(l0)+8))
			v21 = *(*float64)(unsafe.Add(mBase, uint32(l1)+8))
			if base.F64_eq(v20, v21) != 0 {
				v109 = v19
				return v109
			} else {
				v23 = base.F64_sub(v20, v21)
				v24 = base.F64_abs(v23)
				if base.F64_le(v24, float64(1e-06)) != 0 {
					v109 = v19
					return v109
				} else {
					v28 = math.Float64frombits(uint64(0x7ff0000000000000))
					if base.F64_eq(base.F64_abs(v20), v28)|base.F64_ne(v24, v28)|base.F64_eq(base.F64_abs(v21), v28) == int32(0) {
						v40 = F_float_overflow_error_ext(m, int32(0))
						mBase = m.M
						v43 = m.ExcPending
						if v43 != 0 {
							return float64(0)
						} else {
							v44 = *(*float64)(unsafe.Add(mBase, uint32(l0)))
							v45 = *(*float64)(unsafe.Add(mBase, uint32(l1)))
							v47 = base.F64_sub(v44, v45)
							v48 = v40
							v49 = v44
							v50 = v45
							v52 = math.Float64frombits(uint64(0x7ff0000000000000))
							if base.F64_eq(base.F64_abs(v49), v52)|base.F64_ne(base.F64_abs(v47), v52)|base.F64_eq(base.F64_abs(v50), v52) == int32(0) {
								v65 = F_float_overflow_error_ext(m, int32(0))
								mBase = m.M
								v66 = m.ExcPending
								if v66 != 0 {
									return float64(0)
								} else {
									v67 = v65
									if base.B2i32(base.Ui64(int64(9218868437227405312)) < base.Ui64(base.I64_reinterpret_f64(v48)&int64(9223372036854775807)))|base.F64_ne(v67, float64(0)) == int32(0) {
										v79 = F_float_zero_divide_error_ext(m, int32(0))
										mBase = m.M
										v80 = m.ExcPending
										if v80 != 0 {
											return float64(0)
										} else {
											return v79
										}
									} else {
										v83 = math.Float64frombits(uint64(0x7ff0000000000000))
										v85 = base.F64_div(v48, v67)
										if base.F64_eq(base.F64_abs(v48), v83)|base.F64_ne(base.F64_abs(v85), v83) == int32(0) {
											v93 = F_float_overflow_error_ext(m, int32(0))
											mBase = m.M
											v94 = m.ExcPending
											if v94 != 0 {
												return float64(0)
											} else {
												return v93
											}
										} else {
											v96 = float64(0)
											if base.F64_eq(v48, v96)|base.F64_ne(v85, v96)|base.F64_eq(base.F64_abs(v67), math.Float64frombits(uint64(0x7ff0000000000000))) != 0 {
												v109 = v85
												return v109
											} else {
												v106 = F_float_underflow_error_ext(m, int32(0))
												mBase = m.M
												v107 = m.ExcPending
												if v107 != 0 {
													return float64(0)
												} else {
													v109 = v106
													return v109
												}
											}
										}
									}
								}
							} else {
								v67 = v47
								if base.B2i32(base.Ui64(int64(9218868437227405312)) < base.Ui64(base.I64_reinterpret_f64(v48)&int64(9223372036854775807)))|base.F64_ne(v67, float64(0)) == int32(0) {
									v79 = F_float_zero_divide_error_ext(m, int32(0))
									mBase = m.M
									v80 = m.ExcPending
									if v80 != 0 {
										return float64(0)
									} else {
										return v79
									}
								} else {
									v83 = math.Float64frombits(uint64(0x7ff0000000000000))
									v85 = base.F64_div(v48, v67)
									if base.F64_eq(base.F64_abs(v48), v83)|base.F64_ne(base.F64_abs(v85), v83) == int32(0) {
										v93 = F_float_overflow_error_ext(m, int32(0))
										mBase = m.M
										v94 = m.ExcPending
										if v94 != 0 {
											return float64(0)
										} else {
											return v93
										}
									} else {
										v96 = float64(0)
										if base.F64_eq(v48, v96)|base.F64_ne(v85, v96)|base.F64_eq(base.F64_abs(v67), math.Float64frombits(uint64(0x7ff0000000000000))) != 0 {
											v109 = v85
											return v109
										} else {
											v106 = F_float_underflow_error_ext(m, int32(0))
											mBase = m.M
											v107 = m.ExcPending
											if v107 != 0 {
												return float64(0)
											} else {
												v109 = v106
												return v109
											}
										}
									}
								}
							}
						}
					} else {
						v47 = v15
						v48 = v23
						v49 = v12
						v50 = v13
						v52 = math.Float64frombits(uint64(0x7ff0000000000000))
						if base.F64_eq(base.F64_abs(v49), v52)|base.F64_ne(base.F64_abs(v47), v52)|base.F64_eq(base.F64_abs(v50), v52) == int32(0) {
							v65 = F_float_overflow_error_ext(m, int32(0))
							mBase = m.M
							v66 = m.ExcPending
							if v66 != 0 {
								return float64(0)
							} else {
								v67 = v65
								if base.B2i32(base.Ui64(int64(9218868437227405312)) < base.Ui64(base.I64_reinterpret_f64(v48)&int64(9223372036854775807)))|base.F64_ne(v67, float64(0)) == int32(0) {
									v79 = F_float_zero_divide_error_ext(m, int32(0))
									mBase = m.M
									v80 = m.ExcPending
									if v80 != 0 {
										return float64(0)
									} else {
										return v79
									}
								} else {
									v83 = math.Float64frombits(uint64(0x7ff0000000000000))
									v85 = base.F64_div(v48, v67)
									if base.F64_eq(base.F64_abs(v48), v83)|base.F64_ne(base.F64_abs(v85), v83) == int32(0) {
										v93 = F_float_overflow_error_ext(m, int32(0))
										mBase = m.M
										v94 = m.ExcPending
										if v94 != 0 {
											return float64(0)
										} else {
											return v93
										}
									} else {
										v96 = float64(0)
										if base.F64_eq(v48, v96)|base.F64_ne(v85, v96)|base.F64_eq(base.F64_abs(v67), math.Float64frombits(uint64(0x7ff0000000000000))) != 0 {
											v109 = v85
											return v109
										} else {
											v106 = F_float_underflow_error_ext(m, int32(0))
											mBase = m.M
											v107 = m.ExcPending
											if v107 != 0 {
												return float64(0)
											} else {
												v109 = v106
												return v109
											}
										}
									}
								}
							}
						} else {
							v67 = v47
							if base.B2i32(base.Ui64(int64(9218868437227405312)) < base.Ui64(base.I64_reinterpret_f64(v48)&int64(9223372036854775807)))|base.F64_ne(v67, float64(0)) == int32(0) {
								v79 = F_float_zero_divide_error_ext(m, int32(0))
								mBase = m.M
								v80 = m.ExcPending
								if v80 != 0 {
									return float64(0)
								} else {
									return v79
								}
							} else {
								v83 = math.Float64frombits(uint64(0x7ff0000000000000))
								v85 = base.F64_div(v48, v67)
								if base.F64_eq(base.F64_abs(v48), v83)|base.F64_ne(base.F64_abs(v85), v83) == int32(0) {
									v93 = F_float_overflow_error_ext(m, int32(0))
									mBase = m.M
									v94 = m.ExcPending
									if v94 != 0 {
										return float64(0)
									} else {
										return v93
									}
								} else {
									v96 = float64(0)
									if base.F64_eq(v48, v96)|base.F64_ne(v85, v96)|base.F64_eq(base.F64_abs(v67), math.Float64frombits(uint64(0x7ff0000000000000))) != 0 {
										v109 = v85
										return v109
									} else {
										v106 = F_float_underflow_error_ext(m, int32(0))
										mBase = m.M
										v107 = m.ExcPending
										if v107 != 0 {
											return float64(0)
										} else {
											v109 = v106
											return v109
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
func F_point_vert(m *base.Module, l0 int32) int64 {
	mBase := m.M
	_ = mBase
	var v4 int32
	_ = v4
	var v5 float64
	_ = v5
	var v6 int32
	_ = v6
	var v7 float64
	_ = v7
	v4 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
	v5 = *(*float64)(unsafe.Add(mBase, uint32(v4)))
	v6 = *(*int32)(unsafe.Add(mBase, uint32(l0)+40))
	v7 = *(*float64)(unsafe.Add(mBase, uint32(v6)))
	return base.I64_extend_i32_u(base.F64_eq(v5, v7) | base.F64_le(base.F64_abs(base.F64_sub(v5, v7)), float64(1e-06)))
}
