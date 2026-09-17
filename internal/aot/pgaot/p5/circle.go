package p5

import (
	base "github.com/shibukawa/pgmem/internal/aot/pgaot/base"
	"math"
	"unsafe"
)

func F_circle_lt(m *base.Module, l0 int32) int32 {
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
	var v12 float64
	_ = v12
	var v18 float64
	_ = v18
	var v24 float64
	_ = v24
	var v26 float64
	_ = v26
	var v31 float64
	_ = v31
	var v36 float64
	_ = v36
	var v37 float64
	_ = v37
	var v38 float64
	_ = v38
	var v39 float64
	_ = v39
	var v45 float64
	_ = v45
	var v51 float64
	_ = v51
	var v53 float64
	_ = v53
	var v58 float64
	_ = v58
	var v74 int32
	_ = v74
	var v80 int32
	_ = v80
	v7 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
	v8 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v9 = *(*float64)(unsafe.Add(mBase, uint32(v8)+16))
	v10 = base.F64_mul(v9, v9)
	v11 = base.F64_abs(v10)
	v12 = math.Float64frombits(uint64(0x7ff0000000000000))
	if base.F64_eq(v11, v12)&base.F64_ne(base.F64_abs(v9), v12) != 0 {
		F_float_overflow_error(m)
		mBase = m.M
		v74 = m.ExcPending
		if v74 != 0 {
			return int32(0)
		} else {
			base.Wasm_trap_unreachable()
			for {
			}
		}
	} else {
		v18 = float64(0)
		if base.F64_eq(v10, v18)&base.F64_ne(v9, v18) != 0 {
			F_float_underflow_error(m)
			mBase = m.M
			v80 = m.ExcPending
			if v80 != 0 {
				return int32(0)
			} else {
				base.Wasm_trap_unreachable()
				for {
				}
			}
		} else {
			v24 = base.F64_mul(v10, float64(3.141592653589793))
			v26 = math.Float64frombits(uint64(0x7ff0000000000000))
			if base.F64_eq(base.F64_abs(v24), v26)&base.F64_ne(v11, v26) != 0 {
				F_float_overflow_error(m)
				mBase = m.M
				v74 = m.ExcPending
				if v74 != 0 {
					return int32(0)
				} else {
					base.Wasm_trap_unreachable()
					for {
					}
				}
			} else {
				v31 = float64(0)
				if base.F64_eq(v24, v31)&base.F64_ne(v10, v31) != 0 {
					F_float_underflow_error(m)
					mBase = m.M
					v80 = m.ExcPending
					if v80 != 0 {
						return int32(0)
					} else {
						base.Wasm_trap_unreachable()
						for {
						}
					}
				} else {
					v36 = *(*float64)(unsafe.Add(mBase, uint32(v7)+16))
					v37 = base.F64_mul(v36, v36)
					v38 = base.F64_abs(v37)
					v39 = math.Float64frombits(uint64(0x7ff0000000000000))
					if base.F64_eq(v38, v39)&base.F64_ne(base.F64_abs(v36), v39) != 0 {
						F_float_overflow_error(m)
						mBase = m.M
						v74 = m.ExcPending
						if v74 != 0 {
							return int32(0)
						} else {
							base.Wasm_trap_unreachable()
							for {
							}
						}
					} else {
						v45 = float64(0)
						if base.F64_eq(v37, v45)&base.F64_ne(v36, v45) != 0 {
							F_float_underflow_error(m)
							mBase = m.M
							v80 = m.ExcPending
							if v80 != 0 {
								return int32(0)
							} else {
								base.Wasm_trap_unreachable()
								for {
								}
							}
						} else {
							v51 = base.F64_mul(v37, float64(3.141592653589793))
							v53 = math.Float64frombits(uint64(0x7ff0000000000000))
							if base.F64_eq(base.F64_abs(v51), v53)&base.F64_ne(v38, v53) != 0 {
								F_float_overflow_error(m)
								mBase = m.M
								v74 = m.ExcPending
								if v74 != 0 {
									return int32(0)
								} else {
									base.Wasm_trap_unreachable()
									for {
									}
								}
							} else {
								v58 = float64(0)
								if base.F64_eq(v51, v58)&base.F64_ne(v37, v58) != 0 {
									F_float_underflow_error(m)
									mBase = m.M
									v80 = m.ExcPending
									if v80 != 0 {
										return int32(0)
									} else {
										base.Wasm_trap_unreachable()
										for {
										}
									}
								} else {
									return base.F64_lt(base.F64_add(v24, float64(1e-06)), v51)
								}
							}
						}
					}
				}
			}
		}
	}
}
func F_circle_overright(m *base.Module, l0 int32) int32 {
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
	var v25 float64
	_ = v25
	var v26 float64
	_ = v26
	var v27 float64
	_ = v27
	var v29 float64
	_ = v29
	var v51 int32
	_ = v51
	v7 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
	v8 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v9 = *(*float64)(unsafe.Add(mBase, uint32(v8)))
	v10 = *(*float64)(unsafe.Add(mBase, uint32(v8)+16))
	v11 = base.F64_sub(v9, v10)
	v13 = math.Float64frombits(uint64(0x7ff0000000000000))
	if base.B2i32(base.F64_ne(base.F64_abs(v11), v13)|base.F64_eq(base.F64_abs(v9), v13) == int32(0))&base.F64_ne(base.F64_abs(v10), v13) != 0 {
		F_float_overflow_error(m)
		mBase = m.M
		v51 = m.ExcPending
		if v51 != 0 {
			return int32(0)
		} else {
			base.Wasm_trap_unreachable()
			for {
			}
		}
	} else {
		v25 = *(*float64)(unsafe.Add(mBase, uint32(v7)))
		v26 = *(*float64)(unsafe.Add(mBase, uint32(v7)+16))
		v27 = base.F64_sub(v25, v26)
		v29 = math.Float64frombits(uint64(0x7ff0000000000000))
		if base.B2i32(base.F64_ne(base.F64_abs(v27), v29)|base.F64_eq(base.F64_abs(v25), v29) == int32(0))&base.F64_ne(base.F64_abs(v26), v29) != 0 {
			F_float_overflow_error(m)
			mBase = m.M
			v51 = m.ExcPending
			if v51 != 0 {
				return int32(0)
			} else {
				base.Wasm_trap_unreachable()
				for {
				}
			}
		} else {
			return base.F64_ge(base.F64_add(v11, float64(1e-06)), v27)
		}
	}
}
func F_circle_same(m *base.Module, l0 int32) int32 {
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
	var v119 int32
	_ = v119
	v10 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
	v11 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
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
						v119 = v93
						return v119
					} else {
						v99 = v47
						v103 = v50
						v105 = *(*float64)(unsafe.Add(mBase, uint32(v10)+8))
						v108 = base.I64_reinterpret_f64(v105) & int64(9223372036854775807)
						if base.Ui64(int64(9218868437227405313)) <= base.Ui64(v103) {
							return base.B2i32(base.Ui64(int64(9218868437227405312)) < base.Ui64(v108))
						} else {
							return base.B2i32(base.Ui64(v108) < base.Ui64(int64(9218868437227405313))) & base.F64_eq(v105, v99)
						}
					}
				} else {
					v55 = int32(0)
					if base.Ui64(int64(9218868437227405312)) < base.Ui64(v46) {
						v119 = v55
						return v119
					} else {
						v58 = *(*float64)(unsafe.Add(mBase, uint32(v10)+8))
						if base.Ui64(base.I64_reinterpret_f64(v58)&int64(9223372036854775807)) <= base.Ui64(int64(9218868437227405312)) {
							if base.B2i32(base.F64_le(base.F64_abs(base.F64_sub(v37, v43)), float64(1e-06)) == int32(0))&base.F64_ne(v37, v43) != 0 {
								v119 = v55
							} else {
								v119 = base.F64_eq(v47, v58) | base.F64_le(base.F64_abs(base.F64_sub(v47, v58)), float64(1e-06))
							}
							return v119
						} else {
							v92 = int32(1)
							v93 = int32(0)
							if base.B2i32(v92 == v93)|base.F64_ne(v37, v43) != 0 {
								v119 = v93
								return v119
							} else {
								v99 = v47
								v103 = v50
								v105 = *(*float64)(unsafe.Add(mBase, uint32(v10)+8))
								v108 = base.I64_reinterpret_f64(v105) & int64(9223372036854775807)
								if base.Ui64(int64(9218868437227405313)) <= base.Ui64(v103) {
									return base.B2i32(base.Ui64(int64(9218868437227405312)) < base.Ui64(v108))
								} else {
									return base.B2i32(base.Ui64(v108) < base.Ui64(int64(9218868437227405313))) & base.F64_eq(v105, v99)
								}
							}
						}
					}
				}
			} else {
				v65 = *(*int64)(unsafe.Add(mBase, uint32(v10)))
				if base.Ui64(v65&int64(9223372036854775807)) < base.Ui64(int64(9218868437227405313)) {
					return int32(0)
				} else {
					v72 = *(*float64)(unsafe.Add(mBase, uint32(v11)+8))
					v99 = v72
					v103 = base.I64_reinterpret_f64(v72) & int64(9223372036854775807)
					v105 = *(*float64)(unsafe.Add(mBase, uint32(v10)+8))
					v108 = base.I64_reinterpret_f64(v105) & int64(9223372036854775807)
					if base.Ui64(int64(9218868437227405313)) <= base.Ui64(v103) {
						return base.B2i32(base.Ui64(int64(9218868437227405312)) < base.Ui64(v108))
					} else {
						return base.B2i32(base.Ui64(v108) < base.Ui64(int64(9218868437227405313))) & base.F64_eq(v105, v99)
					}
				}
			}
		} else {
			return int32(0)
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
						v119 = v93
						return v119
					} else {
						v99 = v47
						v103 = v50
						v105 = *(*float64)(unsafe.Add(mBase, uint32(v10)+8))
						v108 = base.I64_reinterpret_f64(v105) & int64(9223372036854775807)
						if base.Ui64(int64(9218868437227405313)) <= base.Ui64(v103) {
							return base.B2i32(base.Ui64(int64(9218868437227405312)) < base.Ui64(v108))
						} else {
							return base.B2i32(base.Ui64(v108) < base.Ui64(int64(9218868437227405313))) & base.F64_eq(v105, v99)
						}
					}
				} else {
					v55 = int32(0)
					if base.Ui64(int64(9218868437227405312)) < base.Ui64(v46) {
						v119 = v55
						return v119
					} else {
						v58 = *(*float64)(unsafe.Add(mBase, uint32(v10)+8))
						if base.Ui64(base.I64_reinterpret_f64(v58)&int64(9223372036854775807)) <= base.Ui64(int64(9218868437227405312)) {
							if base.B2i32(base.F64_le(base.F64_abs(base.F64_sub(v37, v43)), float64(1e-06)) == int32(0))&base.F64_ne(v37, v43) != 0 {
								v119 = v55
							} else {
								v119 = base.F64_eq(v47, v58) | base.F64_le(base.F64_abs(base.F64_sub(v47, v58)), float64(1e-06))
							}
							return v119
						} else {
							v92 = int32(1)
							v93 = int32(0)
							if base.B2i32(v92 == v93)|base.F64_ne(v37, v43) != 0 {
								v119 = v93
								return v119
							} else {
								v99 = v47
								v103 = v50
								v105 = *(*float64)(unsafe.Add(mBase, uint32(v10)+8))
								v108 = base.I64_reinterpret_f64(v105) & int64(9223372036854775807)
								if base.Ui64(int64(9218868437227405313)) <= base.Ui64(v103) {
									return base.B2i32(base.Ui64(int64(9218868437227405312)) < base.Ui64(v108))
								} else {
									return base.B2i32(base.Ui64(v108) < base.Ui64(int64(9218868437227405313))) & base.F64_eq(v105, v99)
								}
							}
						}
					}
				}
			} else {
				v65 = *(*int64)(unsafe.Add(mBase, uint32(v10)))
				if base.Ui64(v65&int64(9223372036854775807)) < base.Ui64(int64(9218868437227405313)) {
					return int32(0)
				} else {
					v72 = *(*float64)(unsafe.Add(mBase, uint32(v11)+8))
					v99 = v72
					v103 = base.I64_reinterpret_f64(v72) & int64(9223372036854775807)
					v105 = *(*float64)(unsafe.Add(mBase, uint32(v10)+8))
					v108 = base.I64_reinterpret_f64(v105) & int64(9223372036854775807)
					if base.Ui64(int64(9218868437227405313)) <= base.Ui64(v103) {
						return base.B2i32(base.Ui64(int64(9218868437227405312)) < base.Ui64(v108))
					} else {
						return base.B2i32(base.Ui64(v108) < base.Ui64(int64(9218868437227405313))) & base.F64_eq(v105, v99)
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
							v119 = v93
							return v119
						} else {
							v99 = v47
							v103 = v50
							v105 = *(*float64)(unsafe.Add(mBase, uint32(v10)+8))
							v108 = base.I64_reinterpret_f64(v105) & int64(9223372036854775807)
							if base.Ui64(int64(9218868437227405313)) <= base.Ui64(v103) {
								return base.B2i32(base.Ui64(int64(9218868437227405312)) < base.Ui64(v108))
							} else {
								return base.B2i32(base.Ui64(v108) < base.Ui64(int64(9218868437227405313))) & base.F64_eq(v105, v99)
							}
						}
					} else {
						v55 = int32(0)
						if base.Ui64(int64(9218868437227405312)) < base.Ui64(v46) {
							v119 = v55
							return v119
						} else {
							v58 = *(*float64)(unsafe.Add(mBase, uint32(v10)+8))
							if base.Ui64(base.I64_reinterpret_f64(v58)&int64(9223372036854775807)) <= base.Ui64(int64(9218868437227405312)) {
								if base.B2i32(base.F64_le(base.F64_abs(base.F64_sub(v37, v43)), float64(1e-06)) == int32(0))&base.F64_ne(v37, v43) != 0 {
									v119 = v55
								} else {
									v119 = base.F64_eq(v47, v58) | base.F64_le(base.F64_abs(base.F64_sub(v47, v58)), float64(1e-06))
								}
								return v119
							} else {
								v92 = int32(1)
								v93 = int32(0)
								if base.B2i32(v92 == v93)|base.F64_ne(v37, v43) != 0 {
									v119 = v93
									return v119
								} else {
									v99 = v47
									v103 = v50
									v105 = *(*float64)(unsafe.Add(mBase, uint32(v10)+8))
									v108 = base.I64_reinterpret_f64(v105) & int64(9223372036854775807)
									if base.Ui64(int64(9218868437227405313)) <= base.Ui64(v103) {
										return base.B2i32(base.Ui64(int64(9218868437227405312)) < base.Ui64(v108))
									} else {
										return base.B2i32(base.Ui64(v108) < base.Ui64(int64(9218868437227405313))) & base.F64_eq(v105, v99)
									}
								}
							}
						}
					}
				} else {
					v65 = *(*int64)(unsafe.Add(mBase, uint32(v10)))
					if base.Ui64(v65&int64(9223372036854775807)) < base.Ui64(int64(9218868437227405313)) {
						return int32(0)
					} else {
						v72 = *(*float64)(unsafe.Add(mBase, uint32(v11)+8))
						v99 = v72
						v103 = base.I64_reinterpret_f64(v72) & int64(9223372036854775807)
						v105 = *(*float64)(unsafe.Add(mBase, uint32(v10)+8))
						v108 = base.I64_reinterpret_f64(v105) & int64(9223372036854775807)
						if base.Ui64(int64(9218868437227405313)) <= base.Ui64(v103) {
							return base.B2i32(base.Ui64(int64(9218868437227405312)) < base.Ui64(v108))
						} else {
							return base.B2i32(base.Ui64(v108) < base.Ui64(int64(9218868437227405313))) & base.F64_eq(v105, v99)
						}
					}
				}
			} else {
				return int32(0)
			}
		}
	}
}
