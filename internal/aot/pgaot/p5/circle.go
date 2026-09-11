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
	var v21 float64
	_ = v21
	var v22 float64
	_ = v22
	var v23 float64
	_ = v23
	var v43 int32
	_ = v43
	v7 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
	v8 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v9 = *(*float64)(unsafe.Add(mBase, uint32(v8)))
	v10 = *(*float64)(unsafe.Add(mBase, uint32(v8)+16))
	v11 = base.F64_sub(v9, v10)
	if base.F64_ne(base.F64_abs(v11), math.Float64frombits(uint64(0x7ff0000000000000))) != 0 {
		v21 = *(*float64)(unsafe.Add(mBase, uint32(v7)))
		v22 = *(*float64)(unsafe.Add(mBase, uint32(v7)+16))
		v23 = base.F64_sub(v21, v22)
		if base.F64_ne(base.F64_abs(v23), math.Float64frombits(uint64(0x7ff0000000000000))) != 0 {
			return base.F64_ge(base.F64_add(v11, float64(1e-06)), v23)
		} else {
			if base.F64_eq(base.F64_abs(v21), math.Float64frombits(uint64(0x7ff0000000000000))) != 0 {
				return base.F64_ge(base.F64_add(v11, float64(1e-06)), v23)
			} else {
				if base.F64_ne(base.F64_abs(v22), math.Float64frombits(uint64(0x7ff0000000000000))) != 0 {
					F_float_overflow_error(m)
					mBase = m.M
					v43 = m.ExcPending
					if v43 != 0 {
						return int32(0)
					} else {
						base.Wasm_trap_unreachable()
						for {
						}
					}
				} else {
					return base.F64_ge(base.F64_add(v11, float64(1e-06)), v23)
				}
			}
		}
	} else {
		if base.F64_eq(base.F64_abs(v9), math.Float64frombits(uint64(0x7ff0000000000000))) != 0 {
			v21 = *(*float64)(unsafe.Add(mBase, uint32(v7)))
			v22 = *(*float64)(unsafe.Add(mBase, uint32(v7)+16))
			v23 = base.F64_sub(v21, v22)
			if base.F64_ne(base.F64_abs(v23), math.Float64frombits(uint64(0x7ff0000000000000))) != 0 {
				return base.F64_ge(base.F64_add(v11, float64(1e-06)), v23)
			} else {
				if base.F64_eq(base.F64_abs(v21), math.Float64frombits(uint64(0x7ff0000000000000))) != 0 {
					return base.F64_ge(base.F64_add(v11, float64(1e-06)), v23)
				} else {
					if base.F64_ne(base.F64_abs(v22), math.Float64frombits(uint64(0x7ff0000000000000))) != 0 {
						F_float_overflow_error(m)
						mBase = m.M
						v43 = m.ExcPending
						if v43 != 0 {
							return int32(0)
						} else {
							base.Wasm_trap_unreachable()
							for {
							}
						}
					} else {
						return base.F64_ge(base.F64_add(v11, float64(1e-06)), v23)
					}
				}
			}
		} else {
			if base.F64_ne(base.F64_abs(v10), math.Float64frombits(uint64(0x7ff0000000000000))) != 0 {
				F_float_overflow_error(m)
				mBase = m.M
				v43 = m.ExcPending
				if v43 != 0 {
					return int32(0)
				} else {
					base.Wasm_trap_unreachable()
					for {
					}
				}
			} else {
				v21 = *(*float64)(unsafe.Add(mBase, uint32(v7)))
				v22 = *(*float64)(unsafe.Add(mBase, uint32(v7)+16))
				v23 = base.F64_sub(v21, v22)
				if base.F64_ne(base.F64_abs(v23), math.Float64frombits(uint64(0x7ff0000000000000))) != 0 {
					return base.F64_ge(base.F64_add(v11, float64(1e-06)), v23)
				} else {
					if base.F64_eq(base.F64_abs(v21), math.Float64frombits(uint64(0x7ff0000000000000))) != 0 {
						return base.F64_ge(base.F64_add(v11, float64(1e-06)), v23)
					} else {
						if base.F64_ne(base.F64_abs(v22), math.Float64frombits(uint64(0x7ff0000000000000))) != 0 {
							F_float_overflow_error(m)
							mBase = m.M
							v43 = m.ExcPending
							if v43 != 0 {
								return int32(0)
							} else {
								base.Wasm_trap_unreachable()
								for {
								}
							}
						} else {
							return base.F64_ge(base.F64_add(v11, float64(1e-06)), v23)
						}
					}
				}
			}
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
	var v36 float64
	_ = v36
	var v42 float64
	_ = v42
	var v44 int64
	_ = v44
	var v45 int64
	_ = v45
	var v46 float64
	_ = v46
	var v49 int64
	_ = v49
	var v54 int32
	_ = v54
	var v57 float64
	_ = v57
	var v64 int64
	_ = v64
	var v71 float64
	_ = v71
	var v90 int32
	_ = v90
	var v91 int32
	_ = v91
	var v96 float64
	_ = v96
	var v101 int64
	_ = v101
	var v102 float64
	_ = v102
	var v105 int64
	_ = v105
	var v116 int32
	_ = v116
	v10 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
	v11 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v12 = *(*float64)(unsafe.Add(mBase, uint32(v11)+16))
	if base.Ui64(base.I64_reinterpret_f64(v12)&int64(9223372036854775807)) <= base.Ui64(int64(9218868437227405312)) {
		v18 = *(*float64)(unsafe.Add(mBase, uint32(v10)+16))
		v26 = v18
		if base.F64_eq(v12, v26) != 0 {
			v36 = *(*float64)(unsafe.Add(mBase, uint32(v11)))
			if base.Ui64(base.I64_reinterpret_f64(v36)&int64(9223372036854775807)) <= base.Ui64(int64(9218868437227405312)) {
				v42 = *(*float64)(unsafe.Add(mBase, uint32(v10)))
				v44 = int64(9223372036854775807)
				v45 = base.I64_reinterpret_f64(v42) & v44
				v46 = *(*float64)(unsafe.Add(mBase, uint32(v11)+8))
				v49 = base.I64_reinterpret_f64(v46) & v44
				if base.Ui64(int64(9218868437227405313)) <= base.Ui64(v49) {
					v90 = base.B2i32(base.Ui64(v45) < base.Ui64(int64(9218868437227405313)))
					v91 = int32(0)
					if base.F64_ne(v36, v42) != 0 {
						v116 = v91
						return v116
					} else {
						if v90 == int32(0) {
							v116 = v91
							return v116
						} else {
							v96 = v46
							v101 = v49
							v102 = *(*float64)(unsafe.Add(mBase, uint32(v10)+8))
							v105 = base.I64_reinterpret_f64(v102) & int64(9223372036854775807)
							if base.Ui64(int64(9218868437227405313)) <= base.Ui64(v101) {
								return base.B2i32(base.Ui64(int64(9218868437227405312)) < base.Ui64(v105))
							} else {
								return base.B2i32(base.Ui64(v105) < base.Ui64(int64(9218868437227405313))) & base.F64_eq(v102, v96)
							}
						}
					}
				} else {
					v54 = int32(0)
					if base.Ui64(int64(9218868437227405312)) < base.Ui64(v45) {
						v116 = v54
						return v116
					} else {
						v57 = *(*float64)(unsafe.Add(mBase, uint32(v10)+8))
						if base.Ui64(base.I64_reinterpret_f64(v57)&int64(9223372036854775807)) <= base.Ui64(int64(9218868437227405312)) {
							if base.F64_ne(v36, v42) != 0 {
								if base.F64_le(base.F64_abs(base.F64_sub(v36, v42)), float64(1e-06)) == int32(0) {
									v116 = v54
								} else {
									v116 = base.F64_eq(v46, v57) | base.F64_le(base.F64_abs(base.F64_sub(v46, v57)), float64(1e-06))
								}
							} else {
								v116 = base.F64_eq(v46, v57) | base.F64_le(base.F64_abs(base.F64_sub(v46, v57)), float64(1e-06))
							}
							return v116
						} else {
							v90 = int32(1)
							v91 = int32(0)
							if base.F64_ne(v36, v42) != 0 {
								v116 = v91
								return v116
							} else {
								if v90 == int32(0) {
									v116 = v91
									return v116
								} else {
									v96 = v46
									v101 = v49
									v102 = *(*float64)(unsafe.Add(mBase, uint32(v10)+8))
									v105 = base.I64_reinterpret_f64(v102) & int64(9223372036854775807)
									if base.Ui64(int64(9218868437227405313)) <= base.Ui64(v101) {
										return base.B2i32(base.Ui64(int64(9218868437227405312)) < base.Ui64(v105))
									} else {
										return base.B2i32(base.Ui64(v105) < base.Ui64(int64(9218868437227405313))) & base.F64_eq(v102, v96)
									}
								}
							}
						}
					}
				}
			} else {
				v64 = *(*int64)(unsafe.Add(mBase, uint32(v10)))
				if base.Ui64(v64&int64(9223372036854775807)) < base.Ui64(int64(9218868437227405313)) {
					return int32(0)
				} else {
					v71 = *(*float64)(unsafe.Add(mBase, uint32(v11)+8))
					v96 = v71
					v101 = base.I64_reinterpret_f64(v71) & int64(9223372036854775807)
					v102 = *(*float64)(unsafe.Add(mBase, uint32(v10)+8))
					v105 = base.I64_reinterpret_f64(v102) & int64(9223372036854775807)
					if base.Ui64(int64(9218868437227405313)) <= base.Ui64(v101) {
						return base.B2i32(base.Ui64(int64(9218868437227405312)) < base.Ui64(v105))
					} else {
						return base.B2i32(base.Ui64(v105) < base.Ui64(int64(9218868437227405313))) & base.F64_eq(v102, v96)
					}
				}
			}
		} else {
			if base.F64_le(base.F64_abs(base.F64_sub(v12, v26)), float64(1e-06)) != 0 {
				v36 = *(*float64)(unsafe.Add(mBase, uint32(v11)))
				if base.Ui64(base.I64_reinterpret_f64(v36)&int64(9223372036854775807)) <= base.Ui64(int64(9218868437227405312)) {
					v42 = *(*float64)(unsafe.Add(mBase, uint32(v10)))
					v44 = int64(9223372036854775807)
					v45 = base.I64_reinterpret_f64(v42) & v44
					v46 = *(*float64)(unsafe.Add(mBase, uint32(v11)+8))
					v49 = base.I64_reinterpret_f64(v46) & v44
					if base.Ui64(int64(9218868437227405313)) <= base.Ui64(v49) {
						v90 = base.B2i32(base.Ui64(v45) < base.Ui64(int64(9218868437227405313)))
						v91 = int32(0)
						if base.F64_ne(v36, v42) != 0 {
							v116 = v91
							return v116
						} else {
							if v90 == int32(0) {
								v116 = v91
								return v116
							} else {
								v96 = v46
								v101 = v49
								v102 = *(*float64)(unsafe.Add(mBase, uint32(v10)+8))
								v105 = base.I64_reinterpret_f64(v102) & int64(9223372036854775807)
								if base.Ui64(int64(9218868437227405313)) <= base.Ui64(v101) {
									return base.B2i32(base.Ui64(int64(9218868437227405312)) < base.Ui64(v105))
								} else {
									return base.B2i32(base.Ui64(v105) < base.Ui64(int64(9218868437227405313))) & base.F64_eq(v102, v96)
								}
							}
						}
					} else {
						v54 = int32(0)
						if base.Ui64(int64(9218868437227405312)) < base.Ui64(v45) {
							v116 = v54
							return v116
						} else {
							v57 = *(*float64)(unsafe.Add(mBase, uint32(v10)+8))
							if base.Ui64(base.I64_reinterpret_f64(v57)&int64(9223372036854775807)) <= base.Ui64(int64(9218868437227405312)) {
								if base.F64_ne(v36, v42) != 0 {
									if base.F64_le(base.F64_abs(base.F64_sub(v36, v42)), float64(1e-06)) == int32(0) {
										v116 = v54
									} else {
										v116 = base.F64_eq(v46, v57) | base.F64_le(base.F64_abs(base.F64_sub(v46, v57)), float64(1e-06))
									}
								} else {
									v116 = base.F64_eq(v46, v57) | base.F64_le(base.F64_abs(base.F64_sub(v46, v57)), float64(1e-06))
								}
								return v116
							} else {
								v90 = int32(1)
								v91 = int32(0)
								if base.F64_ne(v36, v42) != 0 {
									v116 = v91
									return v116
								} else {
									if v90 == int32(0) {
										v116 = v91
										return v116
									} else {
										v96 = v46
										v101 = v49
										v102 = *(*float64)(unsafe.Add(mBase, uint32(v10)+8))
										v105 = base.I64_reinterpret_f64(v102) & int64(9223372036854775807)
										if base.Ui64(int64(9218868437227405313)) <= base.Ui64(v101) {
											return base.B2i32(base.Ui64(int64(9218868437227405312)) < base.Ui64(v105))
										} else {
											return base.B2i32(base.Ui64(v105) < base.Ui64(int64(9218868437227405313))) & base.F64_eq(v102, v96)
										}
									}
								}
							}
						}
					}
				} else {
					v64 = *(*int64)(unsafe.Add(mBase, uint32(v10)))
					if base.Ui64(v64&int64(9223372036854775807)) < base.Ui64(int64(9218868437227405313)) {
						return int32(0)
					} else {
						v71 = *(*float64)(unsafe.Add(mBase, uint32(v11)+8))
						v96 = v71
						v101 = base.I64_reinterpret_f64(v71) & int64(9223372036854775807)
						v102 = *(*float64)(unsafe.Add(mBase, uint32(v10)+8))
						v105 = base.I64_reinterpret_f64(v102) & int64(9223372036854775807)
						if base.Ui64(int64(9218868437227405313)) <= base.Ui64(v101) {
							return base.B2i32(base.Ui64(int64(9218868437227405312)) < base.Ui64(v105))
						} else {
							return base.B2i32(base.Ui64(v105) < base.Ui64(int64(9218868437227405313))) & base.F64_eq(v102, v96)
						}
					}
				}
			} else {
				return int32(0)
			}
		}
	} else {
		v19 = *(*int64)(unsafe.Add(mBase, uint32(v10)+16))
		if base.Ui64(int64(9218868437227405312)) < base.Ui64(v19&int64(9223372036854775807)) {
			v36 = *(*float64)(unsafe.Add(mBase, uint32(v11)))
			if base.Ui64(base.I64_reinterpret_f64(v36)&int64(9223372036854775807)) <= base.Ui64(int64(9218868437227405312)) {
				v42 = *(*float64)(unsafe.Add(mBase, uint32(v10)))
				v44 = int64(9223372036854775807)
				v45 = base.I64_reinterpret_f64(v42) & v44
				v46 = *(*float64)(unsafe.Add(mBase, uint32(v11)+8))
				v49 = base.I64_reinterpret_f64(v46) & v44
				if base.Ui64(int64(9218868437227405313)) <= base.Ui64(v49) {
					v90 = base.B2i32(base.Ui64(v45) < base.Ui64(int64(9218868437227405313)))
					v91 = int32(0)
					if base.F64_ne(v36, v42) != 0 {
						v116 = v91
						return v116
					} else {
						if v90 == int32(0) {
							v116 = v91
							return v116
						} else {
							v96 = v46
							v101 = v49
							v102 = *(*float64)(unsafe.Add(mBase, uint32(v10)+8))
							v105 = base.I64_reinterpret_f64(v102) & int64(9223372036854775807)
							if base.Ui64(int64(9218868437227405313)) <= base.Ui64(v101) {
								return base.B2i32(base.Ui64(int64(9218868437227405312)) < base.Ui64(v105))
							} else {
								return base.B2i32(base.Ui64(v105) < base.Ui64(int64(9218868437227405313))) & base.F64_eq(v102, v96)
							}
						}
					}
				} else {
					v54 = int32(0)
					if base.Ui64(int64(9218868437227405312)) < base.Ui64(v45) {
						v116 = v54
						return v116
					} else {
						v57 = *(*float64)(unsafe.Add(mBase, uint32(v10)+8))
						if base.Ui64(base.I64_reinterpret_f64(v57)&int64(9223372036854775807)) <= base.Ui64(int64(9218868437227405312)) {
							if base.F64_ne(v36, v42) != 0 {
								if base.F64_le(base.F64_abs(base.F64_sub(v36, v42)), float64(1e-06)) == int32(0) {
									v116 = v54
								} else {
									v116 = base.F64_eq(v46, v57) | base.F64_le(base.F64_abs(base.F64_sub(v46, v57)), float64(1e-06))
								}
							} else {
								v116 = base.F64_eq(v46, v57) | base.F64_le(base.F64_abs(base.F64_sub(v46, v57)), float64(1e-06))
							}
							return v116
						} else {
							v90 = int32(1)
							v91 = int32(0)
							if base.F64_ne(v36, v42) != 0 {
								v116 = v91
								return v116
							} else {
								if v90 == int32(0) {
									v116 = v91
									return v116
								} else {
									v96 = v46
									v101 = v49
									v102 = *(*float64)(unsafe.Add(mBase, uint32(v10)+8))
									v105 = base.I64_reinterpret_f64(v102) & int64(9223372036854775807)
									if base.Ui64(int64(9218868437227405313)) <= base.Ui64(v101) {
										return base.B2i32(base.Ui64(int64(9218868437227405312)) < base.Ui64(v105))
									} else {
										return base.B2i32(base.Ui64(v105) < base.Ui64(int64(9218868437227405313))) & base.F64_eq(v102, v96)
									}
								}
							}
						}
					}
				}
			} else {
				v64 = *(*int64)(unsafe.Add(mBase, uint32(v10)))
				if base.Ui64(v64&int64(9223372036854775807)) < base.Ui64(int64(9218868437227405313)) {
					return int32(0)
				} else {
					v71 = *(*float64)(unsafe.Add(mBase, uint32(v11)+8))
					v96 = v71
					v101 = base.I64_reinterpret_f64(v71) & int64(9223372036854775807)
					v102 = *(*float64)(unsafe.Add(mBase, uint32(v10)+8))
					v105 = base.I64_reinterpret_f64(v102) & int64(9223372036854775807)
					if base.Ui64(int64(9218868437227405313)) <= base.Ui64(v101) {
						return base.B2i32(base.Ui64(int64(9218868437227405312)) < base.Ui64(v105))
					} else {
						return base.B2i32(base.Ui64(v105) < base.Ui64(int64(9218868437227405313))) & base.F64_eq(v102, v96)
					}
				}
			}
		} else {
			v26 = base.F64_reinterpret_i64(v19)
			if base.F64_eq(v12, v26) != 0 {
				v36 = *(*float64)(unsafe.Add(mBase, uint32(v11)))
				if base.Ui64(base.I64_reinterpret_f64(v36)&int64(9223372036854775807)) <= base.Ui64(int64(9218868437227405312)) {
					v42 = *(*float64)(unsafe.Add(mBase, uint32(v10)))
					v44 = int64(9223372036854775807)
					v45 = base.I64_reinterpret_f64(v42) & v44
					v46 = *(*float64)(unsafe.Add(mBase, uint32(v11)+8))
					v49 = base.I64_reinterpret_f64(v46) & v44
					if base.Ui64(int64(9218868437227405313)) <= base.Ui64(v49) {
						v90 = base.B2i32(base.Ui64(v45) < base.Ui64(int64(9218868437227405313)))
						v91 = int32(0)
						if base.F64_ne(v36, v42) != 0 {
							v116 = v91
							return v116
						} else {
							if v90 == int32(0) {
								v116 = v91
								return v116
							} else {
								v96 = v46
								v101 = v49
								v102 = *(*float64)(unsafe.Add(mBase, uint32(v10)+8))
								v105 = base.I64_reinterpret_f64(v102) & int64(9223372036854775807)
								if base.Ui64(int64(9218868437227405313)) <= base.Ui64(v101) {
									return base.B2i32(base.Ui64(int64(9218868437227405312)) < base.Ui64(v105))
								} else {
									return base.B2i32(base.Ui64(v105) < base.Ui64(int64(9218868437227405313))) & base.F64_eq(v102, v96)
								}
							}
						}
					} else {
						v54 = int32(0)
						if base.Ui64(int64(9218868437227405312)) < base.Ui64(v45) {
							v116 = v54
							return v116
						} else {
							v57 = *(*float64)(unsafe.Add(mBase, uint32(v10)+8))
							if base.Ui64(base.I64_reinterpret_f64(v57)&int64(9223372036854775807)) <= base.Ui64(int64(9218868437227405312)) {
								if base.F64_ne(v36, v42) != 0 {
									if base.F64_le(base.F64_abs(base.F64_sub(v36, v42)), float64(1e-06)) == int32(0) {
										v116 = v54
									} else {
										v116 = base.F64_eq(v46, v57) | base.F64_le(base.F64_abs(base.F64_sub(v46, v57)), float64(1e-06))
									}
								} else {
									v116 = base.F64_eq(v46, v57) | base.F64_le(base.F64_abs(base.F64_sub(v46, v57)), float64(1e-06))
								}
								return v116
							} else {
								v90 = int32(1)
								v91 = int32(0)
								if base.F64_ne(v36, v42) != 0 {
									v116 = v91
									return v116
								} else {
									if v90 == int32(0) {
										v116 = v91
										return v116
									} else {
										v96 = v46
										v101 = v49
										v102 = *(*float64)(unsafe.Add(mBase, uint32(v10)+8))
										v105 = base.I64_reinterpret_f64(v102) & int64(9223372036854775807)
										if base.Ui64(int64(9218868437227405313)) <= base.Ui64(v101) {
											return base.B2i32(base.Ui64(int64(9218868437227405312)) < base.Ui64(v105))
										} else {
											return base.B2i32(base.Ui64(v105) < base.Ui64(int64(9218868437227405313))) & base.F64_eq(v102, v96)
										}
									}
								}
							}
						}
					}
				} else {
					v64 = *(*int64)(unsafe.Add(mBase, uint32(v10)))
					if base.Ui64(v64&int64(9223372036854775807)) < base.Ui64(int64(9218868437227405313)) {
						return int32(0)
					} else {
						v71 = *(*float64)(unsafe.Add(mBase, uint32(v11)+8))
						v96 = v71
						v101 = base.I64_reinterpret_f64(v71) & int64(9223372036854775807)
						v102 = *(*float64)(unsafe.Add(mBase, uint32(v10)+8))
						v105 = base.I64_reinterpret_f64(v102) & int64(9223372036854775807)
						if base.Ui64(int64(9218868437227405313)) <= base.Ui64(v101) {
							return base.B2i32(base.Ui64(int64(9218868437227405312)) < base.Ui64(v105))
						} else {
							return base.B2i32(base.Ui64(v105) < base.Ui64(int64(9218868437227405313))) & base.F64_eq(v102, v96)
						}
					}
				}
			} else {
				if base.F64_le(base.F64_abs(base.F64_sub(v12, v26)), float64(1e-06)) != 0 {
					v36 = *(*float64)(unsafe.Add(mBase, uint32(v11)))
					if base.Ui64(base.I64_reinterpret_f64(v36)&int64(9223372036854775807)) <= base.Ui64(int64(9218868437227405312)) {
						v42 = *(*float64)(unsafe.Add(mBase, uint32(v10)))
						v44 = int64(9223372036854775807)
						v45 = base.I64_reinterpret_f64(v42) & v44
						v46 = *(*float64)(unsafe.Add(mBase, uint32(v11)+8))
						v49 = base.I64_reinterpret_f64(v46) & v44
						if base.Ui64(int64(9218868437227405313)) <= base.Ui64(v49) {
							v90 = base.B2i32(base.Ui64(v45) < base.Ui64(int64(9218868437227405313)))
							v91 = int32(0)
							if base.F64_ne(v36, v42) != 0 {
								v116 = v91
								return v116
							} else {
								if v90 == int32(0) {
									v116 = v91
									return v116
								} else {
									v96 = v46
									v101 = v49
									v102 = *(*float64)(unsafe.Add(mBase, uint32(v10)+8))
									v105 = base.I64_reinterpret_f64(v102) & int64(9223372036854775807)
									if base.Ui64(int64(9218868437227405313)) <= base.Ui64(v101) {
										return base.B2i32(base.Ui64(int64(9218868437227405312)) < base.Ui64(v105))
									} else {
										return base.B2i32(base.Ui64(v105) < base.Ui64(int64(9218868437227405313))) & base.F64_eq(v102, v96)
									}
								}
							}
						} else {
							v54 = int32(0)
							if base.Ui64(int64(9218868437227405312)) < base.Ui64(v45) {
								v116 = v54
								return v116
							} else {
								v57 = *(*float64)(unsafe.Add(mBase, uint32(v10)+8))
								if base.Ui64(base.I64_reinterpret_f64(v57)&int64(9223372036854775807)) <= base.Ui64(int64(9218868437227405312)) {
									if base.F64_ne(v36, v42) != 0 {
										if base.F64_le(base.F64_abs(base.F64_sub(v36, v42)), float64(1e-06)) == int32(0) {
											v116 = v54
										} else {
											v116 = base.F64_eq(v46, v57) | base.F64_le(base.F64_abs(base.F64_sub(v46, v57)), float64(1e-06))
										}
									} else {
										v116 = base.F64_eq(v46, v57) | base.F64_le(base.F64_abs(base.F64_sub(v46, v57)), float64(1e-06))
									}
									return v116
								} else {
									v90 = int32(1)
									v91 = int32(0)
									if base.F64_ne(v36, v42) != 0 {
										v116 = v91
										return v116
									} else {
										if v90 == int32(0) {
											v116 = v91
											return v116
										} else {
											v96 = v46
											v101 = v49
											v102 = *(*float64)(unsafe.Add(mBase, uint32(v10)+8))
											v105 = base.I64_reinterpret_f64(v102) & int64(9223372036854775807)
											if base.Ui64(int64(9218868437227405313)) <= base.Ui64(v101) {
												return base.B2i32(base.Ui64(int64(9218868437227405312)) < base.Ui64(v105))
											} else {
												return base.B2i32(base.Ui64(v105) < base.Ui64(int64(9218868437227405313))) & base.F64_eq(v102, v96)
											}
										}
									}
								}
							}
						}
					} else {
						v64 = *(*int64)(unsafe.Add(mBase, uint32(v10)))
						if base.Ui64(v64&int64(9223372036854775807)) < base.Ui64(int64(9218868437227405313)) {
							return int32(0)
						} else {
							v71 = *(*float64)(unsafe.Add(mBase, uint32(v11)+8))
							v96 = v71
							v101 = base.I64_reinterpret_f64(v71) & int64(9223372036854775807)
							v102 = *(*float64)(unsafe.Add(mBase, uint32(v10)+8))
							v105 = base.I64_reinterpret_f64(v102) & int64(9223372036854775807)
							if base.Ui64(int64(9218868437227405313)) <= base.Ui64(v101) {
								return base.B2i32(base.Ui64(int64(9218868437227405312)) < base.Ui64(v105))
							} else {
								return base.B2i32(base.Ui64(v105) < base.Ui64(int64(9218868437227405313))) & base.F64_eq(v102, v96)
							}
						}
					}
				} else {
					return int32(0)
				}
			}
		}
	}
}
