package p2

import (
	base "github.com/shibukawa/pgmem/internal/aot/pgaot/base"
	"math"
	"unsafe"
)

func F_box_above_eq(m *base.Module, l0 int32) int32 {
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
	v2 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
	v3 = *(*float64)(unsafe.Add(mBase, uint32(v2)+8))
	v4 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v5 = *(*float64)(unsafe.Add(mBase, uint32(v4)+24))
	return base.F64_le(v3, base.F64_add(v5, float64(1e-06)))
}
func F_box_below_eq(m *base.Module, l0 int32) int32 {
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
	v2 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v3 = *(*float64)(unsafe.Add(mBase, uint32(v2)+8))
	v4 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
	v5 = *(*float64)(unsafe.Add(mBase, uint32(v4)+24))
	return base.F64_le(v3, base.F64_add(v5, float64(1e-06)))
}
func F_box_circle(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v6 int32
	_ = v6
	var v8 int32
	_ = v8
	var v11 int32
	_ = v11
	var v12 float64
	_ = v12
	var v13 float64
	_ = v13
	var v14 float64
	_ = v14
	var v19 float64
	_ = v19
	var v28 float64
	_ = v28
	var v32 float64
	_ = v32
	var v33 float64
	_ = v33
	var v39 float64
	_ = v39
	var v40 float64
	_ = v40
	var v41 float64
	_ = v41
	var v46 float64
	_ = v46
	var v55 float64
	_ = v55
	var v59 float64
	_ = v59
	var v60 float64
	_ = v60
	var v66 float64
	_ = v66
	var v67 int32
	_ = v67
	var v74 int32
	_ = v74
	var v79 int32
	_ = v79
	v6 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v8 = F_palloc(m, int32(24))
	mBase = m.M
	v11 = m.ExcPending
	if v11 != 0 {
		return int32(0)
	} else {
		v12 = *(*float64)(unsafe.Add(mBase, uint32(v6)))
		v13 = *(*float64)(unsafe.Add(mBase, uint32(v6)+16))
		v14 = base.F64_add(v12, v13)
		if base.F64_eq(base.F64_abs(v14), math.Float64frombits(uint64(0x7ff0000000000000))) != 0 {
			v19 = math.Float64frombits(uint64(0x7ff0000000000000))
			if base.F64_ne(base.F64_abs(v12), v19)&base.F64_ne(base.F64_abs(v13), v19) != 0 {
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
				v32 = base.F64_mul(v14, float64(0.5))
				v33 = float64(0)
				if base.F64_eq(v32, v33)&base.F64_ne(v14, v33) != 0 {
					F_float_underflow_error(m)
					mBase = m.M
					v79 = m.ExcPending
					if v79 != 0 {
						return int32(0)
					} else {
						base.Wasm_trap_unreachable()
						for {
						}
					}
				} else {
					*(*float64)(unsafe.Add(mBase, uint32(v8))) = v32
					v39 = *(*float64)(unsafe.Add(mBase, uint32(v6)+8))
					v40 = *(*float64)(unsafe.Add(mBase, uint32(v6)+24))
					v41 = base.F64_add(v39, v40)
					if base.F64_eq(base.F64_abs(v41), math.Float64frombits(uint64(0x7ff0000000000000))) != 0 {
						v46 = math.Float64frombits(uint64(0x7ff0000000000000))
						if base.F64_ne(base.F64_abs(v39), v46)&base.F64_ne(base.F64_abs(v40), v46) != 0 {
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
							v59 = base.F64_mul(v41, float64(0.5))
							v60 = float64(0)
							if base.F64_eq(v59, v60)&base.F64_ne(v41, v60) != 0 {
								F_float_underflow_error(m)
								mBase = m.M
								v79 = m.ExcPending
								if v79 != 0 {
									return int32(0)
								} else {
									base.Wasm_trap_unreachable()
									for {
									}
								}
							} else {
								*(*float64)(unsafe.Add(mBase, uint32(v8)+8)) = v59
								v66 = F_point_dt(m, v8, v6)
								mBase = m.M
								v67 = m.ExcPending
								if v67 != 0 {
									return int32(0)
								} else {
									*(*float64)(unsafe.Add(mBase, uint32(v8)+16)) = v66
									return v8
								}
							}
						}
					} else {
						v55 = base.F64_mul(v41, float64(0.5))
						if base.F64_eq(base.F64_abs(v55), math.Float64frombits(uint64(0x7ff0000000000000))) != 0 {
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
							v59 = v55
							v60 = float64(0)
							if base.F64_eq(v59, v60)&base.F64_ne(v41, v60) != 0 {
								F_float_underflow_error(m)
								mBase = m.M
								v79 = m.ExcPending
								if v79 != 0 {
									return int32(0)
								} else {
									base.Wasm_trap_unreachable()
									for {
									}
								}
							} else {
								*(*float64)(unsafe.Add(mBase, uint32(v8)+8)) = v59
								v66 = F_point_dt(m, v8, v6)
								mBase = m.M
								v67 = m.ExcPending
								if v67 != 0 {
									return int32(0)
								} else {
									*(*float64)(unsafe.Add(mBase, uint32(v8)+16)) = v66
									return v8
								}
							}
						}
					}
				}
			}
		} else {
			v28 = base.F64_mul(v14, float64(0.5))
			if base.F64_eq(base.F64_abs(v28), math.Float64frombits(uint64(0x7ff0000000000000))) != 0 {
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
				v32 = v28
				v33 = float64(0)
				if base.F64_eq(v32, v33)&base.F64_ne(v14, v33) != 0 {
					F_float_underflow_error(m)
					mBase = m.M
					v79 = m.ExcPending
					if v79 != 0 {
						return int32(0)
					} else {
						base.Wasm_trap_unreachable()
						for {
						}
					}
				} else {
					*(*float64)(unsafe.Add(mBase, uint32(v8))) = v32
					v39 = *(*float64)(unsafe.Add(mBase, uint32(v6)+8))
					v40 = *(*float64)(unsafe.Add(mBase, uint32(v6)+24))
					v41 = base.F64_add(v39, v40)
					if base.F64_eq(base.F64_abs(v41), math.Float64frombits(uint64(0x7ff0000000000000))) != 0 {
						v46 = math.Float64frombits(uint64(0x7ff0000000000000))
						if base.F64_ne(base.F64_abs(v39), v46)&base.F64_ne(base.F64_abs(v40), v46) != 0 {
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
							v59 = base.F64_mul(v41, float64(0.5))
							v60 = float64(0)
							if base.F64_eq(v59, v60)&base.F64_ne(v41, v60) != 0 {
								F_float_underflow_error(m)
								mBase = m.M
								v79 = m.ExcPending
								if v79 != 0 {
									return int32(0)
								} else {
									base.Wasm_trap_unreachable()
									for {
									}
								}
							} else {
								*(*float64)(unsafe.Add(mBase, uint32(v8)+8)) = v59
								v66 = F_point_dt(m, v8, v6)
								mBase = m.M
								v67 = m.ExcPending
								if v67 != 0 {
									return int32(0)
								} else {
									*(*float64)(unsafe.Add(mBase, uint32(v8)+16)) = v66
									return v8
								}
							}
						}
					} else {
						v55 = base.F64_mul(v41, float64(0.5))
						if base.F64_eq(base.F64_abs(v55), math.Float64frombits(uint64(0x7ff0000000000000))) != 0 {
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
							v59 = v55
							v60 = float64(0)
							if base.F64_eq(v59, v60)&base.F64_ne(v41, v60) != 0 {
								F_float_underflow_error(m)
								mBase = m.M
								v79 = m.ExcPending
								if v79 != 0 {
									return int32(0)
								} else {
									base.Wasm_trap_unreachable()
									for {
									}
								}
							} else {
								*(*float64)(unsafe.Add(mBase, uint32(v8)+8)) = v59
								v66 = F_point_dt(m, v8, v6)
								mBase = m.M
								v67 = m.ExcPending
								if v67 != 0 {
									return int32(0)
								} else {
									*(*float64)(unsafe.Add(mBase, uint32(v8)+16)) = v66
									return v8
								}
							}
						}
					}
				}
			}
		}
	}
}
func F_box_closept_lseg(m *base.Module, l0 int32, l1 int32, l2 int32) float64 {
	mBase := m.M
	_ = mBase
	var v9 int32
	_ = v9
	var v11 int32
	_ = v11
	var v13 int32
	_ = v13
	var v16 int32
	_ = v16
	var v17 float64
	_ = v17
	var v18 float64
	_ = v18
	var v20 float64
	_ = v20
	var v24 float64
	_ = v24
	var v25 int32
	_ = v25
	var v26 float64
	_ = v26
	var v28 float64
	_ = v28
	var v34 float64
	_ = v34
	var v35 int32
	_ = v35
	var v37 int64
	_ = v37
	var v39 int64
	_ = v39
	var v42 int32
	_ = v42
	var v53 int64
	_ = v53
	var v55 int64
	_ = v55
	var v57 float64
	_ = v57
	var v58 float64
	_ = v58
	var v59 float64
	_ = v59
	var v60 float64
	_ = v60
	var v67 float64
	_ = v67
	var v68 int32
	_ = v68
	var v70 int64
	_ = v70
	var v72 int64
	_ = v72
	var v75 int32
	_ = v75
	var v86 int64
	_ = v86
	var v88 int64
	_ = v88
	var v90 float64
	_ = v90
	var v91 float64
	_ = v91
	var v93 float64
	_ = v93
	var v99 float64
	_ = v99
	var v100 int32
	_ = v100
	var v102 int64
	_ = v102
	var v104 int64
	_ = v104
	var v116 int64
	_ = v116
	var v118 int64
	_ = v118
	var v121 float64
	_ = v121
	v9 = m.G0
	v11 = v9 - int32(48)
	m.G0 = v11
	v13 = F_box_interpt_lseg(m, l0, l1, l2)
	mBase = m.M
	v16 = m.ExcPending
	if v16 != 0 {
		return float64(0)
	} else {
		if v13 != 0 {
			v121 = float64(0)
			m.G0 = v11 + int32(48)
			return v121
		} else {
			v17 = *(*float64)(unsafe.Add(mBase, uint32(l1)+8))
			v18 = *(*float64)(unsafe.Add(mBase, uint32(l1)+16))
			*(*float64)(unsafe.Add(mBase, uint32(v11))) = v18
			v20 = *(*float64)(unsafe.Add(mBase, uint32(l1)+24))
			*(*float64)(unsafe.Add(mBase, uint32(v11)+24)) = v17
			*(*float64)(unsafe.Add(mBase, uint32(v11)+16)) = v18
			*(*float64)(unsafe.Add(mBase, uint32(v11)+8)) = v20
			v24 = F_lseg_closept_lseg(m, l0, v11, l2)
			mBase = m.M
			v25 = m.ExcPending
			if v25 != 0 {
				return float64(0)
			} else {
				v26 = *(*float64)(unsafe.Add(mBase, uint32(l1)))
				*(*float64)(unsafe.Add(mBase, uint32(v11))) = v26
				v28 = *(*float64)(unsafe.Add(mBase, uint32(l1)+8))
				*(*float64)(unsafe.Add(mBase, uint32(v11)+24)) = v17
				*(*float64)(unsafe.Add(mBase, uint32(v11)+16)) = v18
				*(*float64)(unsafe.Add(mBase, uint32(v11)+8)) = v28
				v34 = F_lseg_closept_lseg(m, v11+int32(32), v11, l2)
				mBase = m.M
				v35 = m.ExcPending
				if v35 != 0 {
					return float64(0)
				} else {
					v37 = int64(9223372036854775807)
					v39 = int64(9218868437227405312)
					v42 = int32(0)
					if base.B2i32(base.Ui64(v39) < base.Ui64(base.I64_reinterpret_f64(v34)&v37))|base.B2i32(base.F64_lt(v34, v24) == v42)&base.B2i32(base.Ui64(base.I64_reinterpret_f64(v24)&v37) <= base.Ui64(v39)) == v42 {
						if l0 != 0 {
							v53 = *(*int64)(unsafe.Add(mBase, uint32(v11)+40))
							*(*int64)(unsafe.Add(mBase, uint32(l0)+8)) = v53
							v55 = *(*int64)(unsafe.Add(mBase, uint32(v11)+32))
							*(*int64)(unsafe.Add(mBase, uint32(l0))) = v55
						} else {
						}
						v57 = v34
					} else {
						v57 = v24
					}
					v58 = *(*float64)(unsafe.Add(mBase, uint32(l1)+16))
					v59 = *(*float64)(unsafe.Add(mBase, uint32(l1)))
					v60 = *(*float64)(unsafe.Add(mBase, uint32(l1)+24))
					*(*float64)(unsafe.Add(mBase, uint32(v11)+24)) = v60
					*(*float64)(unsafe.Add(mBase, uint32(v11)+16)) = v59
					*(*float64)(unsafe.Add(mBase, uint32(v11)+8)) = v60
					*(*float64)(unsafe.Add(mBase, uint32(v11))) = v58
					v67 = F_lseg_closept_lseg(m, v11+int32(32), v11, l2)
					mBase = m.M
					v68 = m.ExcPending
					if v68 != 0 {
						return float64(0)
					} else {
						v70 = int64(9223372036854775807)
						v72 = int64(9218868437227405312)
						v75 = int32(0)
						if base.B2i32(base.Ui64(v72) < base.Ui64(base.I64_reinterpret_f64(v67)&v70))|base.B2i32(base.F64_lt(v67, v57) == v75)&base.B2i32(base.Ui64(base.I64_reinterpret_f64(v57)&v70) <= base.Ui64(v72)) == v75 {
							if l0 != 0 {
								v86 = *(*int64)(unsafe.Add(mBase, uint32(v11)+40))
								*(*int64)(unsafe.Add(mBase, uint32(l0)+8)) = v86
								v88 = *(*int64)(unsafe.Add(mBase, uint32(v11)+32))
								*(*int64)(unsafe.Add(mBase, uint32(l0))) = v88
							} else {
							}
							v90 = v67
						} else {
							v90 = v57
						}
						v91 = *(*float64)(unsafe.Add(mBase, uint32(l1)))
						*(*float64)(unsafe.Add(mBase, uint32(v11))) = v91
						v93 = *(*float64)(unsafe.Add(mBase, uint32(l1)+8))
						*(*float64)(unsafe.Add(mBase, uint32(v11)+24)) = v60
						*(*float64)(unsafe.Add(mBase, uint32(v11)+16)) = v59
						*(*float64)(unsafe.Add(mBase, uint32(v11)+8)) = v93
						v99 = F_lseg_closept_lseg(m, v11+int32(32), v11, l2)
						mBase = m.M
						v100 = m.ExcPending
						if v100 != 0 {
							return float64(0)
						} else {
							v102 = int64(9223372036854775807)
							v104 = int64(9218868437227405312)
							if base.B2i32(base.Ui64(v104) < base.Ui64(base.I64_reinterpret_f64(v99)&v102))|base.B2i32(base.F64_lt(v99, v90) == int32(0))&base.B2i32(base.Ui64(base.I64_reinterpret_f64(v90)&v102) <= base.Ui64(v104)) != 0 {
								v121 = v90
							} else {
								if l0 != 0 {
									v116 = *(*int64)(unsafe.Add(mBase, uint32(v11)+40))
									*(*int64)(unsafe.Add(mBase, uint32(l0)+8)) = v116
									v118 = *(*int64)(unsafe.Add(mBase, uint32(v11)+32))
									*(*int64)(unsafe.Add(mBase, uint32(l0))) = v118
								} else {
								}
								v121 = v99
							}
							m.G0 = v11 + int32(48)
							return v121
						}
					}
				}
			}
		}
	}
}
func F_box_eq(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v5 int32
	_ = v5
	var v6 int32
	_ = v6
	var v7 float64
	_ = v7
	var v10 int32
	_ = v10
	var v11 float64
	_ = v11
	var v12 int32
	_ = v12
	v5 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
	v6 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v7 = F_box_ar(m, v6)
	mBase = m.M
	v10 = m.ExcPending
	if v10 != 0 {
		return int32(0)
	} else {
		v11 = F_box_ar(m, v5)
		mBase = m.M
		v12 = m.ExcPending
		if v12 != 0 {
			return int32(0)
		} else {
			return base.F64_eq(v7, v11) | base.F64_le(base.F64_abs(base.F64_sub(v7, v11)), float64(1e-06))
		}
	}
}
func F_box_ge(m *base.Module, l0 int32) int32 {
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
	var v11 float64
	_ = v11
	var v12 int32
	_ = v12
	v3 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
	v4 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v5 = F_box_ar(m, v4)
	mBase = m.M
	v8 = m.ExcPending
	if v8 != 0 {
		return int32(0)
	} else {
		v11 = F_box_ar(m, v3)
		mBase = m.M
		v12 = m.ExcPending
		if v12 != 0 {
			return int32(0)
		} else {
			return base.F64_ge(base.F64_add(v5, float64(1e-06)), v11)
		}
	}
}
func F_box_mul(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v8 int32
	_ = v8
	var v9 int32
	_ = v9
	var v10 int32
	_ = v10
	var v12 int32
	_ = v12
	var v13 int32
	_ = v13
	var v15 int32
	_ = v15
	var v18 int32
	_ = v18
	var v22 int32
	_ = v22
	var v26 int32
	_ = v26
	var v27 float64
	_ = v27
	var v29 int64
	_ = v29
	var v31 int64
	_ = v31
	var v33 float64
	_ = v33
	var v44 float64
	_ = v44
	var v45 float64
	_ = v45
	var v48 float64
	_ = v48
	var v50 int64
	_ = v50
	var v52 int64
	_ = v52
	var v54 float64
	_ = v54
	var v65 float64
	_ = v65
	var v66 float64
	_ = v66
	v8 = m.G0
	v9 = int32(32)
	v10 = v8 - v9
	m.G0 = v10
	v12 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
	v13 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v15 = F_palloc(m, v9)
	mBase = m.M
	v18 = m.ExcPending
	if v18 != 0 {
		return int32(0)
	} else {
		F_point_mul_point(m, v10+int32(16), v13, v12)
		mBase = m.M
		v22 = m.ExcPending
		if v22 != 0 {
			return int32(0)
		} else {
			F_point_mul_point(m, v10, v13+int32(16), v12)
			mBase = m.M
			v26 = m.ExcPending
			if v26 != 0 {
				return int32(0)
			} else {
				v27 = *(*float64)(unsafe.Add(mBase, uint32(v10)))
				v29 = int64(9223372036854775807)
				v31 = int64(9218868437227405312)
				v33 = *(*float64)(unsafe.Add(mBase, uint32(v10)+16))
				if base.B2i32(base.Ui64(base.I64_reinterpret_f64(v27)&v29) <= base.Ui64(v31))&(base.B2i32(base.Ui64(v31) < base.Ui64(base.I64_reinterpret_f64(v33)&v29))|base.F64_lt(v27, v33)) == int32(0) {
					v44 = v27
					v45 = v33
				} else {
					v44 = v33
					v45 = v27
				}
				*(*float64)(unsafe.Add(mBase, uint32(v15)+16)) = v45
				*(*float64)(unsafe.Add(mBase, uint32(v15))) = v44
				v48 = *(*float64)(unsafe.Add(mBase, uint32(v10)+8))
				v50 = int64(9223372036854775807)
				v52 = int64(9218868437227405312)
				v54 = *(*float64)(unsafe.Add(mBase, uint32(v10)+24))
				if base.B2i32(base.Ui64(base.I64_reinterpret_f64(v48)&v50) <= base.Ui64(v52))&(base.B2i32(base.Ui64(v52) < base.Ui64(base.I64_reinterpret_f64(v54)&v50))|base.F64_gt(v54, v48)) == int32(0) {
					v65 = v54
					v66 = v48
				} else {
					v65 = v48
					v66 = v54
				}
				*(*float64)(unsafe.Add(mBase, uint32(v15)+24)) = v65
				*(*float64)(unsafe.Add(mBase, uint32(v15)+8)) = v66
				m.G0 = v10 + int32(32)
				return v15
			}
		}
	}
}
func F_box_overright(m *base.Module, l0 int32) int32 {
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
	v2 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
	v3 = *(*float64)(unsafe.Add(mBase, uint32(v2)+16))
	v4 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v5 = *(*float64)(unsafe.Add(mBase, uint32(v4)+16))
	return base.F64_le(v3, base.F64_add(v5, float64(1e-06)))
}
func F_box_recv(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v6 int32
	_ = v6
	var v8 int32
	_ = v8
	var v11 int32
	_ = v11
	var v12 float64
	_ = v12
	var v13 int32
	_ = v13
	var v15 float64
	_ = v15
	var v16 int32
	_ = v16
	var v18 float64
	_ = v18
	var v19 int32
	_ = v19
	var v21 float64
	_ = v21
	var v22 int32
	_ = v22
	var v24 float64
	_ = v24
	var v30 float64
	_ = v30
	var v43 float64
	_ = v43
	var v45 int64
	_ = v45
	var v47 int64
	_ = v47
	var v50 int32
	_ = v50
	v6 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v8 = F_palloc(m, int32(32))
	mBase = m.M
	v11 = m.ExcPending
	if v11 != 0 {
		return int32(0)
	} else {
		v12 = F_pq_getmsgfloat8(m, v6)
		mBase = m.M
		v13 = m.ExcPending
		if v13 != 0 {
			return int32(0)
		} else {
			*(*float64)(unsafe.Add(mBase, uint32(v8))) = v12
			v15 = F_pq_getmsgfloat8(m, v6)
			mBase = m.M
			v16 = m.ExcPending
			if v16 != 0 {
				return int32(0)
			} else {
				*(*float64)(unsafe.Add(mBase, uint32(v8)+8)) = v15
				v18 = F_pq_getmsgfloat8(m, v6)
				mBase = m.M
				v19 = m.ExcPending
				if v19 != 0 {
					return int32(0)
				} else {
					*(*float64)(unsafe.Add(mBase, uint32(v8)+16)) = v18
					v21 = F_pq_getmsgfloat8(m, v6)
					mBase = m.M
					v22 = m.ExcPending
					if v22 != 0 {
						return int32(0)
					} else {
						*(*float64)(unsafe.Add(mBase, uint32(v8)+24)) = v21
						v24 = *(*float64)(unsafe.Add(mBase, uint32(v8)))
						if base.Ui64(int64(9218868437227405312)) < base.Ui64(base.I64_reinterpret_f64(v24)&int64(9223372036854775807)) {
						} else {
							v30 = *(*float64)(unsafe.Add(mBase, uint32(v8)+16))
							if base.B2i32(base.F64_lt(v24, v30) == int32(0))&base.B2i32(base.Ui64(base.I64_reinterpret_f64(v30)&int64(9223372036854775807)) < base.Ui64(int64(9218868437227405313))) != 0 {
							} else {
								*(*float64)(unsafe.Add(mBase, uint32(v8)+16)) = v24
								*(*float64)(unsafe.Add(mBase, uint32(v8))) = v30
							}
						}
						v43 = *(*float64)(unsafe.Add(mBase, uint32(v8)+8))
						v45 = int64(9223372036854775807)
						v47 = int64(9218868437227405312)
						v50 = int32(0)
						if base.B2i32(base.Ui64(v47) < base.Ui64(base.I64_reinterpret_f64(v43)&v45))|base.B2i32(base.F64_lt(v43, v21) == v50)&base.B2i32(base.Ui64(base.I64_reinterpret_f64(v21)&v45) <= base.Ui64(v47)) == v50 {
							*(*float64)(unsafe.Add(mBase, uint32(v8)+24)) = v43
							*(*float64)(unsafe.Add(mBase, uint32(v8)+8)) = v21
						} else {
						}
						return v8
					}
				}
			}
		}
	}
}
func F_box_same(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v8 int32
	_ = v8
	var v11 int32
	_ = v11
	var v12 int32
	_ = v12
	var v13 float64
	_ = v13
	var v19 float64
	_ = v19
	var v21 int64
	_ = v21
	var v22 int64
	_ = v22
	var v23 float64
	_ = v23
	var v26 int64
	_ = v26
	var v33 float64
	_ = v33
	var v40 int64
	_ = v40
	var v47 float64
	_ = v47
	var v67 int32
	_ = v67
	var v72 float64
	_ = v72
	var v76 int64
	_ = v76
	var v78 float64
	_ = v78
	var v81 int64
	_ = v81
	var v98 float64
	_ = v98
	var v104 float64
	_ = v104
	var v106 int64
	_ = v106
	var v107 int64
	_ = v107
	var v108 float64
	_ = v108
	var v111 int64
	_ = v111
	var v118 float64
	_ = v118
	var v125 int64
	_ = v125
	var v132 float64
	_ = v132
	var v150 int32
	_ = v150
	var v157 float64
	_ = v157
	var v161 int64
	_ = v161
	var v162 float64
	_ = v162
	var v165 int64
	_ = v165
	var v184 int32
	_ = v184
	v8 = int32(0)
	v11 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
	v12 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v13 = *(*float64)(unsafe.Add(mBase, uint32(v12)))
	if base.Ui64(base.I64_reinterpret_f64(v13)&int64(9223372036854775807)) <= base.Ui64(int64(9218868437227405312)) {
		v19 = *(*float64)(unsafe.Add(mBase, uint32(v11)))
		v21 = int64(9223372036854775807)
		v22 = base.I64_reinterpret_f64(v19) & v21
		v23 = *(*float64)(unsafe.Add(mBase, uint32(v12)+8))
		v26 = base.I64_reinterpret_f64(v23) & v21
		if base.Ui64(int64(9218868437227405313)) <= base.Ui64(v26) {
			v67 = base.B2i32(base.Ui64(v22) < base.Ui64(int64(9218868437227405313)))
			if base.B2i32(v67 == int32(0))|base.F64_ne(v13, v19) != 0 {
				v184 = v8
				return v184
			} else {
				v72 = v23
				v76 = v26
				v78 = *(*float64)(unsafe.Add(mBase, uint32(v11)+8))
				v81 = base.I64_reinterpret_f64(v78) & int64(9223372036854775807)
				if base.Ui64(v76) <= base.Ui64(int64(9218868437227405312)) {
					if base.F64_ne(v78, v72) != 0 {
						v184 = v8
						return v184
					} else {
						if base.Ui64(v81) < base.Ui64(int64(9218868437227405313)) {
							v98 = *(*float64)(unsafe.Add(mBase, uint32(v12)+16))
							if base.Ui64(base.I64_reinterpret_f64(v98)&int64(9223372036854775807)) <= base.Ui64(int64(9218868437227405312)) {
								v104 = *(*float64)(unsafe.Add(mBase, uint32(v11)+16))
								v106 = int64(9223372036854775807)
								v107 = base.I64_reinterpret_f64(v104) & v106
								v108 = *(*float64)(unsafe.Add(mBase, uint32(v12)+24))
								v111 = base.I64_reinterpret_f64(v108) & v106
								if base.Ui64(int64(9218868437227405313)) <= base.Ui64(v111) {
									v150 = base.B2i32(base.Ui64(v107) < base.Ui64(int64(9218868437227405313)))
									if base.B2i32(v150 == int32(0))|base.F64_ne(v98, v104) != 0 {
										v184 = v8
										return v184
									} else {
										v157 = v108
										v161 = v111
										v162 = *(*float64)(unsafe.Add(mBase, uint32(v11)+24))
										v165 = base.I64_reinterpret_f64(v162) & int64(9223372036854775807)
										if base.Ui64(int64(9218868437227405313)) <= base.Ui64(v161) {
											return base.B2i32(base.Ui64(int64(9218868437227405312)) < base.Ui64(v165))
										} else {
											return base.B2i32(base.Ui64(v165) < base.Ui64(int64(9218868437227405313))) & base.F64_eq(v162, v157)
										}
									}
								} else {
									if base.Ui64(int64(9218868437227405312)) < base.Ui64(v107) {
										v184 = v8
										return v184
									} else {
										v118 = *(*float64)(unsafe.Add(mBase, uint32(v11)+24))
										if base.Ui64(base.I64_reinterpret_f64(v118)&int64(9223372036854775807)) <= base.Ui64(int64(9218868437227405312)) {
											if base.B2i32(base.F64_le(base.F64_abs(base.F64_sub(v98, v104)), float64(1e-06)) == int32(0))&base.F64_ne(v98, v104) != 0 {
												v184 = v8
											} else {
												v184 = base.F64_eq(v108, v118) | base.F64_le(base.F64_abs(base.F64_sub(v108, v118)), float64(1e-06))
											}
											return v184
										} else {
											v150 = int32(1)
											if base.B2i32(v150 == int32(0))|base.F64_ne(v98, v104) != 0 {
												v184 = v8
												return v184
											} else {
												v157 = v108
												v161 = v111
												v162 = *(*float64)(unsafe.Add(mBase, uint32(v11)+24))
												v165 = base.I64_reinterpret_f64(v162) & int64(9223372036854775807)
												if base.Ui64(int64(9218868437227405313)) <= base.Ui64(v161) {
													return base.B2i32(base.Ui64(int64(9218868437227405312)) < base.Ui64(v165))
												} else {
													return base.B2i32(base.Ui64(v165) < base.Ui64(int64(9218868437227405313))) & base.F64_eq(v162, v157)
												}
											}
										}
									}
								}
							} else {
								v125 = *(*int64)(unsafe.Add(mBase, uint32(v11)+16))
								if base.Ui64(v125&int64(9223372036854775807)) < base.Ui64(int64(9218868437227405313)) {
									return int32(0)
								} else {
									v132 = *(*float64)(unsafe.Add(mBase, uint32(v12)+24))
									v157 = v132
									v161 = base.I64_reinterpret_f64(v132) & int64(9223372036854775807)
									v162 = *(*float64)(unsafe.Add(mBase, uint32(v11)+24))
									v165 = base.I64_reinterpret_f64(v162) & int64(9223372036854775807)
									if base.Ui64(int64(9218868437227405313)) <= base.Ui64(v161) {
										return base.B2i32(base.Ui64(int64(9218868437227405312)) < base.Ui64(v165))
									} else {
										return base.B2i32(base.Ui64(v165) < base.Ui64(int64(9218868437227405313))) & base.F64_eq(v162, v157)
									}
								}
							}
						} else {
							v184 = v8
							return v184
						}
					}
				} else {
					if base.Ui64(int64(9218868437227405313)) <= base.Ui64(v81) {
						v98 = *(*float64)(unsafe.Add(mBase, uint32(v12)+16))
						if base.Ui64(base.I64_reinterpret_f64(v98)&int64(9223372036854775807)) <= base.Ui64(int64(9218868437227405312)) {
							v104 = *(*float64)(unsafe.Add(mBase, uint32(v11)+16))
							v106 = int64(9223372036854775807)
							v107 = base.I64_reinterpret_f64(v104) & v106
							v108 = *(*float64)(unsafe.Add(mBase, uint32(v12)+24))
							v111 = base.I64_reinterpret_f64(v108) & v106
							if base.Ui64(int64(9218868437227405313)) <= base.Ui64(v111) {
								v150 = base.B2i32(base.Ui64(v107) < base.Ui64(int64(9218868437227405313)))
								if base.B2i32(v150 == int32(0))|base.F64_ne(v98, v104) != 0 {
									v184 = v8
									return v184
								} else {
									v157 = v108
									v161 = v111
									v162 = *(*float64)(unsafe.Add(mBase, uint32(v11)+24))
									v165 = base.I64_reinterpret_f64(v162) & int64(9223372036854775807)
									if base.Ui64(int64(9218868437227405313)) <= base.Ui64(v161) {
										return base.B2i32(base.Ui64(int64(9218868437227405312)) < base.Ui64(v165))
									} else {
										return base.B2i32(base.Ui64(v165) < base.Ui64(int64(9218868437227405313))) & base.F64_eq(v162, v157)
									}
								}
							} else {
								if base.Ui64(int64(9218868437227405312)) < base.Ui64(v107) {
									v184 = v8
									return v184
								} else {
									v118 = *(*float64)(unsafe.Add(mBase, uint32(v11)+24))
									if base.Ui64(base.I64_reinterpret_f64(v118)&int64(9223372036854775807)) <= base.Ui64(int64(9218868437227405312)) {
										if base.B2i32(base.F64_le(base.F64_abs(base.F64_sub(v98, v104)), float64(1e-06)) == int32(0))&base.F64_ne(v98, v104) != 0 {
											v184 = v8
										} else {
											v184 = base.F64_eq(v108, v118) | base.F64_le(base.F64_abs(base.F64_sub(v108, v118)), float64(1e-06))
										}
										return v184
									} else {
										v150 = int32(1)
										if base.B2i32(v150 == int32(0))|base.F64_ne(v98, v104) != 0 {
											v184 = v8
											return v184
										} else {
											v157 = v108
											v161 = v111
											v162 = *(*float64)(unsafe.Add(mBase, uint32(v11)+24))
											v165 = base.I64_reinterpret_f64(v162) & int64(9223372036854775807)
											if base.Ui64(int64(9218868437227405313)) <= base.Ui64(v161) {
												return base.B2i32(base.Ui64(int64(9218868437227405312)) < base.Ui64(v165))
											} else {
												return base.B2i32(base.Ui64(v165) < base.Ui64(int64(9218868437227405313))) & base.F64_eq(v162, v157)
											}
										}
									}
								}
							}
						} else {
							v125 = *(*int64)(unsafe.Add(mBase, uint32(v11)+16))
							if base.Ui64(v125&int64(9223372036854775807)) < base.Ui64(int64(9218868437227405313)) {
								return int32(0)
							} else {
								v132 = *(*float64)(unsafe.Add(mBase, uint32(v12)+24))
								v157 = v132
								v161 = base.I64_reinterpret_f64(v132) & int64(9223372036854775807)
								v162 = *(*float64)(unsafe.Add(mBase, uint32(v11)+24))
								v165 = base.I64_reinterpret_f64(v162) & int64(9223372036854775807)
								if base.Ui64(int64(9218868437227405313)) <= base.Ui64(v161) {
									return base.B2i32(base.Ui64(int64(9218868437227405312)) < base.Ui64(v165))
								} else {
									return base.B2i32(base.Ui64(v165) < base.Ui64(int64(9218868437227405313))) & base.F64_eq(v162, v157)
								}
							}
						}
					} else {
						return int32(0)
					}
				}
			}
		} else {
			if base.Ui64(int64(9218868437227405312)) < base.Ui64(v22) {
				v184 = v8
				return v184
			} else {
				v33 = *(*float64)(unsafe.Add(mBase, uint32(v11)+8))
				if base.Ui64(base.I64_reinterpret_f64(v33)&int64(9223372036854775807)) <= base.Ui64(int64(9218868437227405312)) {
					if base.B2i32(base.F64_le(base.F64_abs(base.F64_sub(v13, v19)), float64(1e-06)) == int32(0))&base.F64_ne(v13, v19) != 0 {
						v184 = v8
						return v184
					} else {
						if base.F64_eq(v23, v33) != 0 {
							v98 = *(*float64)(unsafe.Add(mBase, uint32(v12)+16))
							if base.Ui64(base.I64_reinterpret_f64(v98)&int64(9223372036854775807)) <= base.Ui64(int64(9218868437227405312)) {
								v104 = *(*float64)(unsafe.Add(mBase, uint32(v11)+16))
								v106 = int64(9223372036854775807)
								v107 = base.I64_reinterpret_f64(v104) & v106
								v108 = *(*float64)(unsafe.Add(mBase, uint32(v12)+24))
								v111 = base.I64_reinterpret_f64(v108) & v106
								if base.Ui64(int64(9218868437227405313)) <= base.Ui64(v111) {
									v150 = base.B2i32(base.Ui64(v107) < base.Ui64(int64(9218868437227405313)))
									if base.B2i32(v150 == int32(0))|base.F64_ne(v98, v104) != 0 {
										v184 = v8
										return v184
									} else {
										v157 = v108
										v161 = v111
										v162 = *(*float64)(unsafe.Add(mBase, uint32(v11)+24))
										v165 = base.I64_reinterpret_f64(v162) & int64(9223372036854775807)
										if base.Ui64(int64(9218868437227405313)) <= base.Ui64(v161) {
											return base.B2i32(base.Ui64(int64(9218868437227405312)) < base.Ui64(v165))
										} else {
											return base.B2i32(base.Ui64(v165) < base.Ui64(int64(9218868437227405313))) & base.F64_eq(v162, v157)
										}
									}
								} else {
									if base.Ui64(int64(9218868437227405312)) < base.Ui64(v107) {
										v184 = v8
										return v184
									} else {
										v118 = *(*float64)(unsafe.Add(mBase, uint32(v11)+24))
										if base.Ui64(base.I64_reinterpret_f64(v118)&int64(9223372036854775807)) <= base.Ui64(int64(9218868437227405312)) {
											if base.B2i32(base.F64_le(base.F64_abs(base.F64_sub(v98, v104)), float64(1e-06)) == int32(0))&base.F64_ne(v98, v104) != 0 {
												v184 = v8
											} else {
												v184 = base.F64_eq(v108, v118) | base.F64_le(base.F64_abs(base.F64_sub(v108, v118)), float64(1e-06))
											}
											return v184
										} else {
											v150 = int32(1)
											if base.B2i32(v150 == int32(0))|base.F64_ne(v98, v104) != 0 {
												v184 = v8
												return v184
											} else {
												v157 = v108
												v161 = v111
												v162 = *(*float64)(unsafe.Add(mBase, uint32(v11)+24))
												v165 = base.I64_reinterpret_f64(v162) & int64(9223372036854775807)
												if base.Ui64(int64(9218868437227405313)) <= base.Ui64(v161) {
													return base.B2i32(base.Ui64(int64(9218868437227405312)) < base.Ui64(v165))
												} else {
													return base.B2i32(base.Ui64(v165) < base.Ui64(int64(9218868437227405313))) & base.F64_eq(v162, v157)
												}
											}
										}
									}
								}
							} else {
								v125 = *(*int64)(unsafe.Add(mBase, uint32(v11)+16))
								if base.Ui64(v125&int64(9223372036854775807)) < base.Ui64(int64(9218868437227405313)) {
									return int32(0)
								} else {
									v132 = *(*float64)(unsafe.Add(mBase, uint32(v12)+24))
									v157 = v132
									v161 = base.I64_reinterpret_f64(v132) & int64(9223372036854775807)
									v162 = *(*float64)(unsafe.Add(mBase, uint32(v11)+24))
									v165 = base.I64_reinterpret_f64(v162) & int64(9223372036854775807)
									if base.Ui64(int64(9218868437227405313)) <= base.Ui64(v161) {
										return base.B2i32(base.Ui64(int64(9218868437227405312)) < base.Ui64(v165))
									} else {
										return base.B2i32(base.Ui64(v165) < base.Ui64(int64(9218868437227405313))) & base.F64_eq(v162, v157)
									}
								}
							}
						} else {
							if base.F64_le(base.F64_abs(base.F64_sub(v23, v33)), float64(1e-06)) == int32(0) {
								v184 = v8
								return v184
							} else {
								v98 = *(*float64)(unsafe.Add(mBase, uint32(v12)+16))
								if base.Ui64(base.I64_reinterpret_f64(v98)&int64(9223372036854775807)) <= base.Ui64(int64(9218868437227405312)) {
									v104 = *(*float64)(unsafe.Add(mBase, uint32(v11)+16))
									v106 = int64(9223372036854775807)
									v107 = base.I64_reinterpret_f64(v104) & v106
									v108 = *(*float64)(unsafe.Add(mBase, uint32(v12)+24))
									v111 = base.I64_reinterpret_f64(v108) & v106
									if base.Ui64(int64(9218868437227405313)) <= base.Ui64(v111) {
										v150 = base.B2i32(base.Ui64(v107) < base.Ui64(int64(9218868437227405313)))
										if base.B2i32(v150 == int32(0))|base.F64_ne(v98, v104) != 0 {
											v184 = v8
											return v184
										} else {
											v157 = v108
											v161 = v111
											v162 = *(*float64)(unsafe.Add(mBase, uint32(v11)+24))
											v165 = base.I64_reinterpret_f64(v162) & int64(9223372036854775807)
											if base.Ui64(int64(9218868437227405313)) <= base.Ui64(v161) {
												return base.B2i32(base.Ui64(int64(9218868437227405312)) < base.Ui64(v165))
											} else {
												return base.B2i32(base.Ui64(v165) < base.Ui64(int64(9218868437227405313))) & base.F64_eq(v162, v157)
											}
										}
									} else {
										if base.Ui64(int64(9218868437227405312)) < base.Ui64(v107) {
											v184 = v8
											return v184
										} else {
											v118 = *(*float64)(unsafe.Add(mBase, uint32(v11)+24))
											if base.Ui64(base.I64_reinterpret_f64(v118)&int64(9223372036854775807)) <= base.Ui64(int64(9218868437227405312)) {
												if base.B2i32(base.F64_le(base.F64_abs(base.F64_sub(v98, v104)), float64(1e-06)) == int32(0))&base.F64_ne(v98, v104) != 0 {
													v184 = v8
												} else {
													v184 = base.F64_eq(v108, v118) | base.F64_le(base.F64_abs(base.F64_sub(v108, v118)), float64(1e-06))
												}
												return v184
											} else {
												v150 = int32(1)
												if base.B2i32(v150 == int32(0))|base.F64_ne(v98, v104) != 0 {
													v184 = v8
													return v184
												} else {
													v157 = v108
													v161 = v111
													v162 = *(*float64)(unsafe.Add(mBase, uint32(v11)+24))
													v165 = base.I64_reinterpret_f64(v162) & int64(9223372036854775807)
													if base.Ui64(int64(9218868437227405313)) <= base.Ui64(v161) {
														return base.B2i32(base.Ui64(int64(9218868437227405312)) < base.Ui64(v165))
													} else {
														return base.B2i32(base.Ui64(v165) < base.Ui64(int64(9218868437227405313))) & base.F64_eq(v162, v157)
													}
												}
											}
										}
									}
								} else {
									v125 = *(*int64)(unsafe.Add(mBase, uint32(v11)+16))
									if base.Ui64(v125&int64(9223372036854775807)) < base.Ui64(int64(9218868437227405313)) {
										return int32(0)
									} else {
										v132 = *(*float64)(unsafe.Add(mBase, uint32(v12)+24))
										v157 = v132
										v161 = base.I64_reinterpret_f64(v132) & int64(9223372036854775807)
										v162 = *(*float64)(unsafe.Add(mBase, uint32(v11)+24))
										v165 = base.I64_reinterpret_f64(v162) & int64(9223372036854775807)
										if base.Ui64(int64(9218868437227405313)) <= base.Ui64(v161) {
											return base.B2i32(base.Ui64(int64(9218868437227405312)) < base.Ui64(v165))
										} else {
											return base.B2i32(base.Ui64(v165) < base.Ui64(int64(9218868437227405313))) & base.F64_eq(v162, v157)
										}
									}
								}
							}
						}
					}
				} else {
					v67 = int32(1)
					if base.B2i32(v67 == int32(0))|base.F64_ne(v13, v19) != 0 {
						v184 = v8
						return v184
					} else {
						v72 = v23
						v76 = v26
						v78 = *(*float64)(unsafe.Add(mBase, uint32(v11)+8))
						v81 = base.I64_reinterpret_f64(v78) & int64(9223372036854775807)
						if base.Ui64(v76) <= base.Ui64(int64(9218868437227405312)) {
							if base.F64_ne(v78, v72) != 0 {
								v184 = v8
								return v184
							} else {
								if base.Ui64(v81) < base.Ui64(int64(9218868437227405313)) {
									v98 = *(*float64)(unsafe.Add(mBase, uint32(v12)+16))
									if base.Ui64(base.I64_reinterpret_f64(v98)&int64(9223372036854775807)) <= base.Ui64(int64(9218868437227405312)) {
										v104 = *(*float64)(unsafe.Add(mBase, uint32(v11)+16))
										v106 = int64(9223372036854775807)
										v107 = base.I64_reinterpret_f64(v104) & v106
										v108 = *(*float64)(unsafe.Add(mBase, uint32(v12)+24))
										v111 = base.I64_reinterpret_f64(v108) & v106
										if base.Ui64(int64(9218868437227405313)) <= base.Ui64(v111) {
											v150 = base.B2i32(base.Ui64(v107) < base.Ui64(int64(9218868437227405313)))
											if base.B2i32(v150 == int32(0))|base.F64_ne(v98, v104) != 0 {
												v184 = v8
												return v184
											} else {
												v157 = v108
												v161 = v111
												v162 = *(*float64)(unsafe.Add(mBase, uint32(v11)+24))
												v165 = base.I64_reinterpret_f64(v162) & int64(9223372036854775807)
												if base.Ui64(int64(9218868437227405313)) <= base.Ui64(v161) {
													return base.B2i32(base.Ui64(int64(9218868437227405312)) < base.Ui64(v165))
												} else {
													return base.B2i32(base.Ui64(v165) < base.Ui64(int64(9218868437227405313))) & base.F64_eq(v162, v157)
												}
											}
										} else {
											if base.Ui64(int64(9218868437227405312)) < base.Ui64(v107) {
												v184 = v8
												return v184
											} else {
												v118 = *(*float64)(unsafe.Add(mBase, uint32(v11)+24))
												if base.Ui64(base.I64_reinterpret_f64(v118)&int64(9223372036854775807)) <= base.Ui64(int64(9218868437227405312)) {
													if base.B2i32(base.F64_le(base.F64_abs(base.F64_sub(v98, v104)), float64(1e-06)) == int32(0))&base.F64_ne(v98, v104) != 0 {
														v184 = v8
													} else {
														v184 = base.F64_eq(v108, v118) | base.F64_le(base.F64_abs(base.F64_sub(v108, v118)), float64(1e-06))
													}
													return v184
												} else {
													v150 = int32(1)
													if base.B2i32(v150 == int32(0))|base.F64_ne(v98, v104) != 0 {
														v184 = v8
														return v184
													} else {
														v157 = v108
														v161 = v111
														v162 = *(*float64)(unsafe.Add(mBase, uint32(v11)+24))
														v165 = base.I64_reinterpret_f64(v162) & int64(9223372036854775807)
														if base.Ui64(int64(9218868437227405313)) <= base.Ui64(v161) {
															return base.B2i32(base.Ui64(int64(9218868437227405312)) < base.Ui64(v165))
														} else {
															return base.B2i32(base.Ui64(v165) < base.Ui64(int64(9218868437227405313))) & base.F64_eq(v162, v157)
														}
													}
												}
											}
										}
									} else {
										v125 = *(*int64)(unsafe.Add(mBase, uint32(v11)+16))
										if base.Ui64(v125&int64(9223372036854775807)) < base.Ui64(int64(9218868437227405313)) {
											return int32(0)
										} else {
											v132 = *(*float64)(unsafe.Add(mBase, uint32(v12)+24))
											v157 = v132
											v161 = base.I64_reinterpret_f64(v132) & int64(9223372036854775807)
											v162 = *(*float64)(unsafe.Add(mBase, uint32(v11)+24))
											v165 = base.I64_reinterpret_f64(v162) & int64(9223372036854775807)
											if base.Ui64(int64(9218868437227405313)) <= base.Ui64(v161) {
												return base.B2i32(base.Ui64(int64(9218868437227405312)) < base.Ui64(v165))
											} else {
												return base.B2i32(base.Ui64(v165) < base.Ui64(int64(9218868437227405313))) & base.F64_eq(v162, v157)
											}
										}
									}
								} else {
									v184 = v8
									return v184
								}
							}
						} else {
							if base.Ui64(int64(9218868437227405313)) <= base.Ui64(v81) {
								v98 = *(*float64)(unsafe.Add(mBase, uint32(v12)+16))
								if base.Ui64(base.I64_reinterpret_f64(v98)&int64(9223372036854775807)) <= base.Ui64(int64(9218868437227405312)) {
									v104 = *(*float64)(unsafe.Add(mBase, uint32(v11)+16))
									v106 = int64(9223372036854775807)
									v107 = base.I64_reinterpret_f64(v104) & v106
									v108 = *(*float64)(unsafe.Add(mBase, uint32(v12)+24))
									v111 = base.I64_reinterpret_f64(v108) & v106
									if base.Ui64(int64(9218868437227405313)) <= base.Ui64(v111) {
										v150 = base.B2i32(base.Ui64(v107) < base.Ui64(int64(9218868437227405313)))
										if base.B2i32(v150 == int32(0))|base.F64_ne(v98, v104) != 0 {
											v184 = v8
											return v184
										} else {
											v157 = v108
											v161 = v111
											v162 = *(*float64)(unsafe.Add(mBase, uint32(v11)+24))
											v165 = base.I64_reinterpret_f64(v162) & int64(9223372036854775807)
											if base.Ui64(int64(9218868437227405313)) <= base.Ui64(v161) {
												return base.B2i32(base.Ui64(int64(9218868437227405312)) < base.Ui64(v165))
											} else {
												return base.B2i32(base.Ui64(v165) < base.Ui64(int64(9218868437227405313))) & base.F64_eq(v162, v157)
											}
										}
									} else {
										if base.Ui64(int64(9218868437227405312)) < base.Ui64(v107) {
											v184 = v8
											return v184
										} else {
											v118 = *(*float64)(unsafe.Add(mBase, uint32(v11)+24))
											if base.Ui64(base.I64_reinterpret_f64(v118)&int64(9223372036854775807)) <= base.Ui64(int64(9218868437227405312)) {
												if base.B2i32(base.F64_le(base.F64_abs(base.F64_sub(v98, v104)), float64(1e-06)) == int32(0))&base.F64_ne(v98, v104) != 0 {
													v184 = v8
												} else {
													v184 = base.F64_eq(v108, v118) | base.F64_le(base.F64_abs(base.F64_sub(v108, v118)), float64(1e-06))
												}
												return v184
											} else {
												v150 = int32(1)
												if base.B2i32(v150 == int32(0))|base.F64_ne(v98, v104) != 0 {
													v184 = v8
													return v184
												} else {
													v157 = v108
													v161 = v111
													v162 = *(*float64)(unsafe.Add(mBase, uint32(v11)+24))
													v165 = base.I64_reinterpret_f64(v162) & int64(9223372036854775807)
													if base.Ui64(int64(9218868437227405313)) <= base.Ui64(v161) {
														return base.B2i32(base.Ui64(int64(9218868437227405312)) < base.Ui64(v165))
													} else {
														return base.B2i32(base.Ui64(v165) < base.Ui64(int64(9218868437227405313))) & base.F64_eq(v162, v157)
													}
												}
											}
										}
									}
								} else {
									v125 = *(*int64)(unsafe.Add(mBase, uint32(v11)+16))
									if base.Ui64(v125&int64(9223372036854775807)) < base.Ui64(int64(9218868437227405313)) {
										return int32(0)
									} else {
										v132 = *(*float64)(unsafe.Add(mBase, uint32(v12)+24))
										v157 = v132
										v161 = base.I64_reinterpret_f64(v132) & int64(9223372036854775807)
										v162 = *(*float64)(unsafe.Add(mBase, uint32(v11)+24))
										v165 = base.I64_reinterpret_f64(v162) & int64(9223372036854775807)
										if base.Ui64(int64(9218868437227405313)) <= base.Ui64(v161) {
											return base.B2i32(base.Ui64(int64(9218868437227405312)) < base.Ui64(v165))
										} else {
											return base.B2i32(base.Ui64(v165) < base.Ui64(int64(9218868437227405313))) & base.F64_eq(v162, v157)
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
		}
	} else {
		v40 = *(*int64)(unsafe.Add(mBase, uint32(v11)))
		if base.Ui64(v40&int64(9223372036854775807)) < base.Ui64(int64(9218868437227405313)) {
			return int32(0)
		} else {
			v47 = *(*float64)(unsafe.Add(mBase, uint32(v12)+8))
			v72 = v47
			v76 = base.I64_reinterpret_f64(v47) & int64(9223372036854775807)
			v78 = *(*float64)(unsafe.Add(mBase, uint32(v11)+8))
			v81 = base.I64_reinterpret_f64(v78) & int64(9223372036854775807)
			if base.Ui64(v76) <= base.Ui64(int64(9218868437227405312)) {
				if base.F64_ne(v78, v72) != 0 {
					v184 = v8
					return v184
				} else {
					if base.Ui64(v81) < base.Ui64(int64(9218868437227405313)) {
						v98 = *(*float64)(unsafe.Add(mBase, uint32(v12)+16))
						if base.Ui64(base.I64_reinterpret_f64(v98)&int64(9223372036854775807)) <= base.Ui64(int64(9218868437227405312)) {
							v104 = *(*float64)(unsafe.Add(mBase, uint32(v11)+16))
							v106 = int64(9223372036854775807)
							v107 = base.I64_reinterpret_f64(v104) & v106
							v108 = *(*float64)(unsafe.Add(mBase, uint32(v12)+24))
							v111 = base.I64_reinterpret_f64(v108) & v106
							if base.Ui64(int64(9218868437227405313)) <= base.Ui64(v111) {
								v150 = base.B2i32(base.Ui64(v107) < base.Ui64(int64(9218868437227405313)))
								if base.B2i32(v150 == int32(0))|base.F64_ne(v98, v104) != 0 {
									v184 = v8
									return v184
								} else {
									v157 = v108
									v161 = v111
									v162 = *(*float64)(unsafe.Add(mBase, uint32(v11)+24))
									v165 = base.I64_reinterpret_f64(v162) & int64(9223372036854775807)
									if base.Ui64(int64(9218868437227405313)) <= base.Ui64(v161) {
										return base.B2i32(base.Ui64(int64(9218868437227405312)) < base.Ui64(v165))
									} else {
										return base.B2i32(base.Ui64(v165) < base.Ui64(int64(9218868437227405313))) & base.F64_eq(v162, v157)
									}
								}
							} else {
								if base.Ui64(int64(9218868437227405312)) < base.Ui64(v107) {
									v184 = v8
									return v184
								} else {
									v118 = *(*float64)(unsafe.Add(mBase, uint32(v11)+24))
									if base.Ui64(base.I64_reinterpret_f64(v118)&int64(9223372036854775807)) <= base.Ui64(int64(9218868437227405312)) {
										if base.B2i32(base.F64_le(base.F64_abs(base.F64_sub(v98, v104)), float64(1e-06)) == int32(0))&base.F64_ne(v98, v104) != 0 {
											v184 = v8
										} else {
											v184 = base.F64_eq(v108, v118) | base.F64_le(base.F64_abs(base.F64_sub(v108, v118)), float64(1e-06))
										}
										return v184
									} else {
										v150 = int32(1)
										if base.B2i32(v150 == int32(0))|base.F64_ne(v98, v104) != 0 {
											v184 = v8
											return v184
										} else {
											v157 = v108
											v161 = v111
											v162 = *(*float64)(unsafe.Add(mBase, uint32(v11)+24))
											v165 = base.I64_reinterpret_f64(v162) & int64(9223372036854775807)
											if base.Ui64(int64(9218868437227405313)) <= base.Ui64(v161) {
												return base.B2i32(base.Ui64(int64(9218868437227405312)) < base.Ui64(v165))
											} else {
												return base.B2i32(base.Ui64(v165) < base.Ui64(int64(9218868437227405313))) & base.F64_eq(v162, v157)
											}
										}
									}
								}
							}
						} else {
							v125 = *(*int64)(unsafe.Add(mBase, uint32(v11)+16))
							if base.Ui64(v125&int64(9223372036854775807)) < base.Ui64(int64(9218868437227405313)) {
								return int32(0)
							} else {
								v132 = *(*float64)(unsafe.Add(mBase, uint32(v12)+24))
								v157 = v132
								v161 = base.I64_reinterpret_f64(v132) & int64(9223372036854775807)
								v162 = *(*float64)(unsafe.Add(mBase, uint32(v11)+24))
								v165 = base.I64_reinterpret_f64(v162) & int64(9223372036854775807)
								if base.Ui64(int64(9218868437227405313)) <= base.Ui64(v161) {
									return base.B2i32(base.Ui64(int64(9218868437227405312)) < base.Ui64(v165))
								} else {
									return base.B2i32(base.Ui64(v165) < base.Ui64(int64(9218868437227405313))) & base.F64_eq(v162, v157)
								}
							}
						}
					} else {
						v184 = v8
						return v184
					}
				}
			} else {
				if base.Ui64(int64(9218868437227405313)) <= base.Ui64(v81) {
					v98 = *(*float64)(unsafe.Add(mBase, uint32(v12)+16))
					if base.Ui64(base.I64_reinterpret_f64(v98)&int64(9223372036854775807)) <= base.Ui64(int64(9218868437227405312)) {
						v104 = *(*float64)(unsafe.Add(mBase, uint32(v11)+16))
						v106 = int64(9223372036854775807)
						v107 = base.I64_reinterpret_f64(v104) & v106
						v108 = *(*float64)(unsafe.Add(mBase, uint32(v12)+24))
						v111 = base.I64_reinterpret_f64(v108) & v106
						if base.Ui64(int64(9218868437227405313)) <= base.Ui64(v111) {
							v150 = base.B2i32(base.Ui64(v107) < base.Ui64(int64(9218868437227405313)))
							if base.B2i32(v150 == int32(0))|base.F64_ne(v98, v104) != 0 {
								v184 = v8
								return v184
							} else {
								v157 = v108
								v161 = v111
								v162 = *(*float64)(unsafe.Add(mBase, uint32(v11)+24))
								v165 = base.I64_reinterpret_f64(v162) & int64(9223372036854775807)
								if base.Ui64(int64(9218868437227405313)) <= base.Ui64(v161) {
									return base.B2i32(base.Ui64(int64(9218868437227405312)) < base.Ui64(v165))
								} else {
									return base.B2i32(base.Ui64(v165) < base.Ui64(int64(9218868437227405313))) & base.F64_eq(v162, v157)
								}
							}
						} else {
							if base.Ui64(int64(9218868437227405312)) < base.Ui64(v107) {
								v184 = v8
								return v184
							} else {
								v118 = *(*float64)(unsafe.Add(mBase, uint32(v11)+24))
								if base.Ui64(base.I64_reinterpret_f64(v118)&int64(9223372036854775807)) <= base.Ui64(int64(9218868437227405312)) {
									if base.B2i32(base.F64_le(base.F64_abs(base.F64_sub(v98, v104)), float64(1e-06)) == int32(0))&base.F64_ne(v98, v104) != 0 {
										v184 = v8
									} else {
										v184 = base.F64_eq(v108, v118) | base.F64_le(base.F64_abs(base.F64_sub(v108, v118)), float64(1e-06))
									}
									return v184
								} else {
									v150 = int32(1)
									if base.B2i32(v150 == int32(0))|base.F64_ne(v98, v104) != 0 {
										v184 = v8
										return v184
									} else {
										v157 = v108
										v161 = v111
										v162 = *(*float64)(unsafe.Add(mBase, uint32(v11)+24))
										v165 = base.I64_reinterpret_f64(v162) & int64(9223372036854775807)
										if base.Ui64(int64(9218868437227405313)) <= base.Ui64(v161) {
											return base.B2i32(base.Ui64(int64(9218868437227405312)) < base.Ui64(v165))
										} else {
											return base.B2i32(base.Ui64(v165) < base.Ui64(int64(9218868437227405313))) & base.F64_eq(v162, v157)
										}
									}
								}
							}
						}
					} else {
						v125 = *(*int64)(unsafe.Add(mBase, uint32(v11)+16))
						if base.Ui64(v125&int64(9223372036854775807)) < base.Ui64(int64(9218868437227405313)) {
							return int32(0)
						} else {
							v132 = *(*float64)(unsafe.Add(mBase, uint32(v12)+24))
							v157 = v132
							v161 = base.I64_reinterpret_f64(v132) & int64(9223372036854775807)
							v162 = *(*float64)(unsafe.Add(mBase, uint32(v11)+24))
							v165 = base.I64_reinterpret_f64(v162) & int64(9223372036854775807)
							if base.Ui64(int64(9218868437227405313)) <= base.Ui64(v161) {
								return base.B2i32(base.Ui64(int64(9218868437227405312)) < base.Ui64(v165))
							} else {
								return base.B2i32(base.Ui64(v165) < base.Ui64(int64(9218868437227405313))) & base.F64_eq(v162, v157)
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
