package p5

import (
	base "github.com/shibukawa/pgmem/internal/aot/pgaot/base"
	"math"
	"unsafe"
)

func F_circle_lt(m *base.Module, l0 int32) int64 {
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
	return base.I64_extend_i32_u(base.F64_lt(base.F64_add(v55, float64(1e-06)), v101))
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
func F_circle_overright(m *base.Module, l0 int32) int64 {
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
			v33 = base.F64_sub(v31, v32)
			v35 = math.Float64frombits(uint64(0x7ff0000000000000))
			if base.F64_ne(base.F64_abs(v33), v35)|base.F64_eq(base.F64_abs(v31), v35)|base.F64_eq(base.F64_abs(v32), v35) != 0 {
				v48 = v33
				return base.I64_extend_i32_u(base.F64_le(v48, base.F64_add(v30, float64(1e-06))))
			} else {
				v46 = F_float_overflow_error_ext(m, int32(0))
				mBase = m.M
				v47 = m.ExcPending
				if v47 != 0 {
					return int64(0)
				} else {
					v48 = v46
					return base.I64_extend_i32_u(base.F64_le(v48, base.F64_add(v30, float64(1e-06))))
				}
			}
		}
	} else {
		v30 = v11
		v31 = *(*float64)(unsafe.Add(mBase, uint32(v7)))
		v32 = *(*float64)(unsafe.Add(mBase, uint32(v7)+16))
		v33 = base.F64_sub(v31, v32)
		v35 = math.Float64frombits(uint64(0x7ff0000000000000))
		if base.F64_ne(base.F64_abs(v33), v35)|base.F64_eq(base.F64_abs(v31), v35)|base.F64_eq(base.F64_abs(v32), v35) != 0 {
			v48 = v33
			return base.I64_extend_i32_u(base.F64_le(v48, base.F64_add(v30, float64(1e-06))))
		} else {
			v46 = F_float_overflow_error_ext(m, int32(0))
			mBase = m.M
			v47 = m.ExcPending
			if v47 != 0 {
				return int64(0)
			} else {
				v48 = v46
				return base.I64_extend_i32_u(base.F64_le(v48, base.F64_add(v30, float64(1e-06))))
			}
		}
	}
}
func F_circle_same(m *base.Module, l0 int32) int64 {
	mBase := m.M
	_ = mBase
	var v10 int32
	_ = v10
	var v11 int32
	_ = v11
	var v12 float64
	_ = v12
	var v18 float64
	_ = v18
	var v19 int64
	_ = v19
	var v26 float64
	_ = v26
	var v37 float64
	_ = v37
	var v43 float64
	_ = v43
	var v45 int64
	_ = v45
	var v46 int64
	_ = v46
	var v47 float64
	_ = v47
	var v50 int64
	_ = v50
	var v55 int32
	_ = v55
	var v58 float64
	_ = v58
	var v65 int64
	_ = v65
	var v72 float64
	_ = v72
	var v92 int32
	_ = v92
	var v93 int32
	_ = v93
	var v99 float64
	_ = v99
	var v103 int64
	_ = v103
	var v105 float64
	_ = v105
	var v108 int64
	_ = v108
	var v121 int32
	_ = v121
	v10 = *(*int32)(unsafe.Add(mBase, uint32(l0)+40))
	v11 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
	v12 = *(*float64)(unsafe.Add(mBase, uint32(v11)+16))
	if base.Ui64(base.I64_reinterpret_f64(v12)&int64(9223372036854775807)) <= base.Ui64(int64(9218868437227405312)) {
		v18 = *(*float64)(unsafe.Add(mBase, uint32(v10)+16))
		v26 = v18
		if base.F64_eq(v26, v12)|base.F64_le(base.F64_abs(base.F64_sub(v12, v26)), float64(1e-06)) != 0 {
			v37 = *(*float64)(unsafe.Add(mBase, uint32(v11)))
			if base.Ui64(base.I64_reinterpret_f64(v37)&int64(9223372036854775807)) <= base.Ui64(int64(9218868437227405312)) {
				v43 = *(*float64)(unsafe.Add(mBase, uint32(v10)))
				v45 = int64(9223372036854775807)
				v46 = base.I64_reinterpret_f64(v43) & v45
				v47 = *(*float64)(unsafe.Add(mBase, uint32(v11)+8))
				v50 = base.I64_reinterpret_f64(v47) & v45
				if base.Ui64(int64(9218868437227405313)) <= base.Ui64(v50) {
					v92 = base.B2i32(base.Ui64(v46) < base.Ui64(int64(9218868437227405313)))
					v93 = int32(0)
					if base.B2i32(v92 == v93)|base.F64_ne(v37, v43) != 0 {
						v121 = v93
						return base.I64_extend_i32_u(v121)
					} else {
						v99 = v47
						v103 = v50
						v105 = *(*float64)(unsafe.Add(mBase, uint32(v10)+8))
						v108 = base.I64_reinterpret_f64(v105) & int64(9223372036854775807)
						if base.Ui64(int64(9218868437227405313)) <= base.Ui64(v103) {
							return base.I64_extend_i32_u(base.B2i32(base.Ui64(int64(9218868437227405312)) < base.Ui64(v108)))
						} else {
							return base.I64_extend_i32_u(base.B2i32(base.Ui64(v108) < base.Ui64(int64(9218868437227405313))) & base.F64_eq(v105, v99))
						}
					}
				} else {
					v55 = int32(0)
					if base.Ui64(int64(9218868437227405312)) < base.Ui64(v46) {
						v121 = v55
						return base.I64_extend_i32_u(v121)
					} else {
						v58 = *(*float64)(unsafe.Add(mBase, uint32(v10)+8))
						if base.Ui64(base.I64_reinterpret_f64(v58)&int64(9223372036854775807)) <= base.Ui64(int64(9218868437227405312)) {
							if base.B2i32(base.F64_le(base.F64_abs(base.F64_sub(v37, v43)), float64(1e-06)) == int32(0))&base.F64_ne(v37, v43) != 0 {
								v121 = v55
							} else {
								v121 = base.F64_eq(v47, v58) | base.F64_le(base.F64_abs(base.F64_sub(v47, v58)), float64(1e-06))
							}
							return base.I64_extend_i32_u(v121)
						} else {
							v92 = int32(1)
							v93 = int32(0)
							if base.B2i32(v92 == v93)|base.F64_ne(v37, v43) != 0 {
								v121 = v93
								return base.I64_extend_i32_u(v121)
							} else {
								v99 = v47
								v103 = v50
								v105 = *(*float64)(unsafe.Add(mBase, uint32(v10)+8))
								v108 = base.I64_reinterpret_f64(v105) & int64(9223372036854775807)
								if base.Ui64(int64(9218868437227405313)) <= base.Ui64(v103) {
									return base.I64_extend_i32_u(base.B2i32(base.Ui64(int64(9218868437227405312)) < base.Ui64(v108)))
								} else {
									return base.I64_extend_i32_u(base.B2i32(base.Ui64(v108) < base.Ui64(int64(9218868437227405313))) & base.F64_eq(v105, v99))
								}
							}
						}
					}
				}
			} else {
				v65 = *(*int64)(unsafe.Add(mBase, uint32(v10)))
				if base.Ui64(v65&int64(9223372036854775807)) < base.Ui64(int64(9218868437227405313)) {
					return int64(0)
				} else {
					v72 = *(*float64)(unsafe.Add(mBase, uint32(v11)+8))
					v99 = v72
					v103 = base.I64_reinterpret_f64(v72) & int64(9223372036854775807)
					v105 = *(*float64)(unsafe.Add(mBase, uint32(v10)+8))
					v108 = base.I64_reinterpret_f64(v105) & int64(9223372036854775807)
					if base.Ui64(int64(9218868437227405313)) <= base.Ui64(v103) {
						return base.I64_extend_i32_u(base.B2i32(base.Ui64(int64(9218868437227405312)) < base.Ui64(v108)))
					} else {
						return base.I64_extend_i32_u(base.B2i32(base.Ui64(v108) < base.Ui64(int64(9218868437227405313))) & base.F64_eq(v105, v99))
					}
				}
			}
		} else {
			return int64(0)
		}
	} else {
		v19 = *(*int64)(unsafe.Add(mBase, uint32(v10)+16))
		if base.Ui64(int64(9218868437227405312)) < base.Ui64(v19&int64(9223372036854775807)) {
			v37 = *(*float64)(unsafe.Add(mBase, uint32(v11)))
			if base.Ui64(base.I64_reinterpret_f64(v37)&int64(9223372036854775807)) <= base.Ui64(int64(9218868437227405312)) {
				v43 = *(*float64)(unsafe.Add(mBase, uint32(v10)))
				v45 = int64(9223372036854775807)
				v46 = base.I64_reinterpret_f64(v43) & v45
				v47 = *(*float64)(unsafe.Add(mBase, uint32(v11)+8))
				v50 = base.I64_reinterpret_f64(v47) & v45
				if base.Ui64(int64(9218868437227405313)) <= base.Ui64(v50) {
					v92 = base.B2i32(base.Ui64(v46) < base.Ui64(int64(9218868437227405313)))
					v93 = int32(0)
					if base.B2i32(v92 == v93)|base.F64_ne(v37, v43) != 0 {
						v121 = v93
						return base.I64_extend_i32_u(v121)
					} else {
						v99 = v47
						v103 = v50
						v105 = *(*float64)(unsafe.Add(mBase, uint32(v10)+8))
						v108 = base.I64_reinterpret_f64(v105) & int64(9223372036854775807)
						if base.Ui64(int64(9218868437227405313)) <= base.Ui64(v103) {
							return base.I64_extend_i32_u(base.B2i32(base.Ui64(int64(9218868437227405312)) < base.Ui64(v108)))
						} else {
							return base.I64_extend_i32_u(base.B2i32(base.Ui64(v108) < base.Ui64(int64(9218868437227405313))) & base.F64_eq(v105, v99))
						}
					}
				} else {
					v55 = int32(0)
					if base.Ui64(int64(9218868437227405312)) < base.Ui64(v46) {
						v121 = v55
						return base.I64_extend_i32_u(v121)
					} else {
						v58 = *(*float64)(unsafe.Add(mBase, uint32(v10)+8))
						if base.Ui64(base.I64_reinterpret_f64(v58)&int64(9223372036854775807)) <= base.Ui64(int64(9218868437227405312)) {
							if base.B2i32(base.F64_le(base.F64_abs(base.F64_sub(v37, v43)), float64(1e-06)) == int32(0))&base.F64_ne(v37, v43) != 0 {
								v121 = v55
							} else {
								v121 = base.F64_eq(v47, v58) | base.F64_le(base.F64_abs(base.F64_sub(v47, v58)), float64(1e-06))
							}
							return base.I64_extend_i32_u(v121)
						} else {
							v92 = int32(1)
							v93 = int32(0)
							if base.B2i32(v92 == v93)|base.F64_ne(v37, v43) != 0 {
								v121 = v93
								return base.I64_extend_i32_u(v121)
							} else {
								v99 = v47
								v103 = v50
								v105 = *(*float64)(unsafe.Add(mBase, uint32(v10)+8))
								v108 = base.I64_reinterpret_f64(v105) & int64(9223372036854775807)
								if base.Ui64(int64(9218868437227405313)) <= base.Ui64(v103) {
									return base.I64_extend_i32_u(base.B2i32(base.Ui64(int64(9218868437227405312)) < base.Ui64(v108)))
								} else {
									return base.I64_extend_i32_u(base.B2i32(base.Ui64(v108) < base.Ui64(int64(9218868437227405313))) & base.F64_eq(v105, v99))
								}
							}
						}
					}
				}
			} else {
				v65 = *(*int64)(unsafe.Add(mBase, uint32(v10)))
				if base.Ui64(v65&int64(9223372036854775807)) < base.Ui64(int64(9218868437227405313)) {
					return int64(0)
				} else {
					v72 = *(*float64)(unsafe.Add(mBase, uint32(v11)+8))
					v99 = v72
					v103 = base.I64_reinterpret_f64(v72) & int64(9223372036854775807)
					v105 = *(*float64)(unsafe.Add(mBase, uint32(v10)+8))
					v108 = base.I64_reinterpret_f64(v105) & int64(9223372036854775807)
					if base.Ui64(int64(9218868437227405313)) <= base.Ui64(v103) {
						return base.I64_extend_i32_u(base.B2i32(base.Ui64(int64(9218868437227405312)) < base.Ui64(v108)))
					} else {
						return base.I64_extend_i32_u(base.B2i32(base.Ui64(v108) < base.Ui64(int64(9218868437227405313))) & base.F64_eq(v105, v99))
					}
				}
			}
		} else {
			v26 = base.F64_reinterpret_i64(v19)
			if base.F64_eq(v26, v12)|base.F64_le(base.F64_abs(base.F64_sub(v12, v26)), float64(1e-06)) != 0 {
				v37 = *(*float64)(unsafe.Add(mBase, uint32(v11)))
				if base.Ui64(base.I64_reinterpret_f64(v37)&int64(9223372036854775807)) <= base.Ui64(int64(9218868437227405312)) {
					v43 = *(*float64)(unsafe.Add(mBase, uint32(v10)))
					v45 = int64(9223372036854775807)
					v46 = base.I64_reinterpret_f64(v43) & v45
					v47 = *(*float64)(unsafe.Add(mBase, uint32(v11)+8))
					v50 = base.I64_reinterpret_f64(v47) & v45
					if base.Ui64(int64(9218868437227405313)) <= base.Ui64(v50) {
						v92 = base.B2i32(base.Ui64(v46) < base.Ui64(int64(9218868437227405313)))
						v93 = int32(0)
						if base.B2i32(v92 == v93)|base.F64_ne(v37, v43) != 0 {
							v121 = v93
							return base.I64_extend_i32_u(v121)
						} else {
							v99 = v47
							v103 = v50
							v105 = *(*float64)(unsafe.Add(mBase, uint32(v10)+8))
							v108 = base.I64_reinterpret_f64(v105) & int64(9223372036854775807)
							if base.Ui64(int64(9218868437227405313)) <= base.Ui64(v103) {
								return base.I64_extend_i32_u(base.B2i32(base.Ui64(int64(9218868437227405312)) < base.Ui64(v108)))
							} else {
								return base.I64_extend_i32_u(base.B2i32(base.Ui64(v108) < base.Ui64(int64(9218868437227405313))) & base.F64_eq(v105, v99))
							}
						}
					} else {
						v55 = int32(0)
						if base.Ui64(int64(9218868437227405312)) < base.Ui64(v46) {
							v121 = v55
							return base.I64_extend_i32_u(v121)
						} else {
							v58 = *(*float64)(unsafe.Add(mBase, uint32(v10)+8))
							if base.Ui64(base.I64_reinterpret_f64(v58)&int64(9223372036854775807)) <= base.Ui64(int64(9218868437227405312)) {
								if base.B2i32(base.F64_le(base.F64_abs(base.F64_sub(v37, v43)), float64(1e-06)) == int32(0))&base.F64_ne(v37, v43) != 0 {
									v121 = v55
								} else {
									v121 = base.F64_eq(v47, v58) | base.F64_le(base.F64_abs(base.F64_sub(v47, v58)), float64(1e-06))
								}
								return base.I64_extend_i32_u(v121)
							} else {
								v92 = int32(1)
								v93 = int32(0)
								if base.B2i32(v92 == v93)|base.F64_ne(v37, v43) != 0 {
									v121 = v93
									return base.I64_extend_i32_u(v121)
								} else {
									v99 = v47
									v103 = v50
									v105 = *(*float64)(unsafe.Add(mBase, uint32(v10)+8))
									v108 = base.I64_reinterpret_f64(v105) & int64(9223372036854775807)
									if base.Ui64(int64(9218868437227405313)) <= base.Ui64(v103) {
										return base.I64_extend_i32_u(base.B2i32(base.Ui64(int64(9218868437227405312)) < base.Ui64(v108)))
									} else {
										return base.I64_extend_i32_u(base.B2i32(base.Ui64(v108) < base.Ui64(int64(9218868437227405313))) & base.F64_eq(v105, v99))
									}
								}
							}
						}
					}
				} else {
					v65 = *(*int64)(unsafe.Add(mBase, uint32(v10)))
					if base.Ui64(v65&int64(9223372036854775807)) < base.Ui64(int64(9218868437227405313)) {
						return int64(0)
					} else {
						v72 = *(*float64)(unsafe.Add(mBase, uint32(v11)+8))
						v99 = v72
						v103 = base.I64_reinterpret_f64(v72) & int64(9223372036854775807)
						v105 = *(*float64)(unsafe.Add(mBase, uint32(v10)+8))
						v108 = base.I64_reinterpret_f64(v105) & int64(9223372036854775807)
						if base.Ui64(int64(9218868437227405313)) <= base.Ui64(v103) {
							return base.I64_extend_i32_u(base.B2i32(base.Ui64(int64(9218868437227405312)) < base.Ui64(v108)))
						} else {
							return base.I64_extend_i32_u(base.B2i32(base.Ui64(v108) < base.Ui64(int64(9218868437227405313))) & base.F64_eq(v105, v99))
						}
					}
				}
			} else {
				return int64(0)
			}
		}
	}
}
func F_circle_to_poly(m *base.Module, l0 int32) int64 {
	mBase := m.M
	_ = mBase
	var v3 int32
	_ = v3
	var v4 int32
	_ = v4
	var v7 int32
	_ = v7
	v3 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
	v4 = F_circle_poly_internal(m, int32(12), v3, l0)
	mBase = m.M
	v7 = m.ExcPending
	if v7 != 0 {
		return int64(0)
	} else {
		return base.I64_extend_i32_u(v4)
	}
}
