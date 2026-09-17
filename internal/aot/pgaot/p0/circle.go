package p0

import (
	base "github.com/shibukawa/pgmem/internal/aot/pgaot/base"
	"math"
	"unsafe"
)

func F_circle_below(m *base.Module, l0 int32) int32 {
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
	v9 = *(*float64)(unsafe.Add(mBase, uint32(v8)+8))
	v10 = *(*float64)(unsafe.Add(mBase, uint32(v8)+16))
	v11 = base.F64_add(v9, v10)
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
		v25 = *(*float64)(unsafe.Add(mBase, uint32(v7)+8))
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
			return base.F64_lt(base.F64_add(v11, float64(1e-06)), v27)
		}
	}
}
func F_circle_box(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v7 int32
	_ = v7
	var v9 int32
	_ = v9
	var v12 int32
	_ = v12
	var v13 float64
	_ = v13
	var v15 float64
	_ = v15
	var v16 float64
	_ = v16
	var v17 float64
	_ = v17
	var v23 float64
	_ = v23
	var v28 float64
	_ = v28
	var v29 int32
	_ = v29
	var v30 float64
	_ = v30
	var v31 float64
	_ = v31
	var v43 float64
	_ = v43
	var v44 float64
	_ = v44
	var v46 float64
	_ = v46
	var v56 float64
	_ = v56
	var v57 float64
	_ = v57
	var v59 float64
	_ = v59
	var v69 float64
	_ = v69
	var v70 float64
	_ = v70
	var v72 float64
	_ = v72
	var v84 int32
	_ = v84
	var v86 int32
	_ = v86
	v7 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v9 = F_palloc(m, int32(32))
	mBase = m.M
	v12 = m.ExcPending
	if v12 != 0 {
		return int32(0)
	} else {
		v13 = *(*float64)(unsafe.Add(mBase, uint32(v7)+16))
		v15 = base.F64_div(v13, float64(1.4142135623730951))
		v16 = base.F64_abs(v15)
		v17 = math.Float64frombits(uint64(0x7ff0000000000000))
		if base.F64_eq(v16, v17)&base.F64_ne(base.F64_abs(v13), v17) != 0 {
			F_float_overflow_error(m)
			mBase = m.M
			v84 = m.ExcPending
			if v84 != 0 {
				return int32(0)
			} else {
				base.Wasm_trap_unreachable()
				for {
				}
			}
		} else {
			v23 = float64(0)
			if base.F64_eq(v15, v23)&base.F64_ne(v13, v23) != 0 {
				F_float_underflow_error(m)
				mBase = m.M
				v86 = m.ExcPending
				if v86 != 0 {
					return int32(0)
				} else {
					base.Wasm_trap_unreachable()
					for {
					}
				}
			} else {
				v28 = math.Float64frombits(uint64(0x7ff0000000000000))
				v29 = base.F64_eq(v16, v28)
				v30 = *(*float64)(unsafe.Add(mBase, uint32(v7)))
				v31 = base.F64_add(v15, v30)
				if base.B2i32(v29|base.F64_ne(base.F64_abs(v31), v28) == int32(0))&base.F64_ne(base.F64_abs(v30), v28) != 0 {
					F_float_overflow_error(m)
					mBase = m.M
					v84 = m.ExcPending
					if v84 != 0 {
						return int32(0)
					} else {
						base.Wasm_trap_unreachable()
						for {
						}
					}
				} else {
					*(*float64)(unsafe.Add(mBase, uint32(v9))) = v31
					v43 = *(*float64)(unsafe.Add(mBase, uint32(v7)))
					v44 = base.F64_sub(v43, v15)
					v46 = math.Float64frombits(uint64(0x7ff0000000000000))
					if base.B2i32(base.F64_ne(base.F64_abs(v44), v46)|v29 == int32(0))&base.F64_ne(base.F64_abs(v43), v46) != 0 {
						F_float_overflow_error(m)
						mBase = m.M
						v84 = m.ExcPending
						if v84 != 0 {
							return int32(0)
						} else {
							base.Wasm_trap_unreachable()
							for {
							}
						}
					} else {
						*(*float64)(unsafe.Add(mBase, uint32(v9)+16)) = v44
						v56 = *(*float64)(unsafe.Add(mBase, uint32(v7)+8))
						v57 = base.F64_add(v15, v56)
						v59 = math.Float64frombits(uint64(0x7ff0000000000000))
						if base.B2i32(base.F64_ne(base.F64_abs(v57), v59)|v29 == int32(0))&base.F64_ne(base.F64_abs(v56), v59) != 0 {
							F_float_overflow_error(m)
							mBase = m.M
							v84 = m.ExcPending
							if v84 != 0 {
								return int32(0)
							} else {
								base.Wasm_trap_unreachable()
								for {
								}
							}
						} else {
							*(*float64)(unsafe.Add(mBase, uint32(v9)+8)) = v57
							v69 = *(*float64)(unsafe.Add(mBase, uint32(v7)+8))
							v70 = base.F64_sub(v69, v15)
							v72 = math.Float64frombits(uint64(0x7ff0000000000000))
							if base.F64_ne(base.F64_abs(v70), v72)|v29|base.F64_eq(base.F64_abs(v69), v72) != 0 {
								*(*float64)(unsafe.Add(mBase, uint32(v9)+24)) = v70
								return v9
							} else {
								F_float_overflow_error(m)
								mBase = m.M
								v84 = m.ExcPending
								if v84 != 0 {
									return int32(0)
								} else {
									base.Wasm_trap_unreachable()
									for {
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
func F_circle_contain_pt(m *base.Module, l0 int32) int32 {
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
	v3 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v4 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
	v5 = F_point_dt(m, v3, v4)
	mBase = m.M
	v8 = m.ExcPending
	if v8 != 0 {
		return int32(0)
	} else {
		v9 = *(*float64)(unsafe.Add(mBase, uint32(v3)+16))
		return base.F64_le(v5, v9)
	}
}
func F_circle_gt(m *base.Module, l0 int32) int32 {
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
									return base.F64_gt(v24, base.F64_add(v51, float64(1e-06)))
								}
							}
						}
					}
				}
			}
		}
	}
}
func F_circle_left(m *base.Module, l0 int32) int32 {
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
	v11 = base.F64_add(v9, v10)
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
			return base.F64_lt(base.F64_add(v11, float64(1e-06)), v27)
		}
	}
}
func F_circle_overabove(m *base.Module, l0 int32) int32 {
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
	v9 = *(*float64)(unsafe.Add(mBase, uint32(v8)+8))
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
		v25 = *(*float64)(unsafe.Add(mBase, uint32(v7)+8))
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
func F_circle_overbelow(m *base.Module, l0 int32) int32 {
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
	v9 = *(*float64)(unsafe.Add(mBase, uint32(v8)+8))
	v10 = *(*float64)(unsafe.Add(mBase, uint32(v8)+16))
	v11 = base.F64_add(v9, v10)
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
		v25 = *(*float64)(unsafe.Add(mBase, uint32(v7)+8))
		v26 = *(*float64)(unsafe.Add(mBase, uint32(v7)+16))
		v27 = base.F64_add(v25, v26)
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
			return base.F64_le(v11, base.F64_add(v27, float64(1e-06)))
		}
	}
}
func F_circle_overlap(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v7 int32
	_ = v7
	var v8 int32
	_ = v8
	var v9 float64
	_ = v9
	var v12 int32
	_ = v12
	var v13 float64
	_ = v13
	var v14 float64
	_ = v14
	var v15 float64
	_ = v15
	var v17 float64
	_ = v17
	var v30 int32
	_ = v30
	v7 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v8 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
	v9 = F_point_dt(m, v7, v8)
	mBase = m.M
	v12 = m.ExcPending
	if v12 != 0 {
		return int32(0)
	} else {
		v13 = *(*float64)(unsafe.Add(mBase, uint32(v7)+16))
		v14 = *(*float64)(unsafe.Add(mBase, uint32(v8)+16))
		v15 = base.F64_add(v13, v14)
		v17 = math.Float64frombits(uint64(0x7ff0000000000000))
		if base.F64_ne(base.F64_abs(v15), v17)|base.F64_eq(base.F64_abs(v13), v17)|base.F64_eq(base.F64_abs(v14), v17) == int32(0) {
			F_float_overflow_error(m)
			mBase = m.M
			v30 = m.ExcPending
			if v30 != 0 {
				return int32(0)
			} else {
				base.Wasm_trap_unreachable()
				for {
				}
			}
		} else {
			return base.F64_le(v9, base.F64_add(v15, float64(1e-06)))
		}
	}
}
func F_circle_sub_pt(m *base.Module, l0 int32) int32 {
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
	var v31 float64
	_ = v31
	var v32 float64
	_ = v32
	var v33 float64
	_ = v33
	var v35 float64
	_ = v35
	var v49 float64
	_ = v49
	var v56 int32
	_ = v56
	v8 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v9 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
	v11 = F_palloc(m, int32(24))
	mBase = m.M
	v14 = m.ExcPending
	if v14 != 0 {
		return int32(0)
	} else {
		v15 = *(*float64)(unsafe.Add(mBase, uint32(v8)))
		v16 = *(*float64)(unsafe.Add(mBase, uint32(v9)))
		v17 = base.F64_sub(v15, v16)
		v19 = math.Float64frombits(uint64(0x7ff0000000000000))
		if base.B2i32(base.F64_ne(base.F64_abs(v17), v19)|base.F64_eq(base.F64_abs(v15), v19) == int32(0))&base.F64_ne(base.F64_abs(v16), v19) != 0 {
			F_float_overflow_error(m)
			mBase = m.M
			v56 = m.ExcPending
			if v56 != 0 {
				return int32(0)
			} else {
				base.Wasm_trap_unreachable()
				for {
				}
			}
		} else {
			v31 = *(*float64)(unsafe.Add(mBase, uint32(v8)+8))
			v32 = *(*float64)(unsafe.Add(mBase, uint32(v9)+8))
			v33 = base.F64_sub(v31, v32)
			v35 = math.Float64frombits(uint64(0x7ff0000000000000))
			if base.B2i32(base.F64_ne(base.F64_abs(v33), v35)|base.F64_eq(base.F64_abs(v31), v35) == int32(0))&base.F64_ne(base.F64_abs(v32), v35) != 0 {
				F_float_overflow_error(m)
				mBase = m.M
				v56 = m.ExcPending
				if v56 != 0 {
					return int32(0)
				} else {
					base.Wasm_trap_unreachable()
					for {
					}
				}
			} else {
				*(*float64)(unsafe.Add(mBase, uint32(v11)+8)) = v33
				*(*float64)(unsafe.Add(mBase, uint32(v11))) = v17
				v49 = *(*float64)(unsafe.Add(mBase, uint32(v8)+16))
				*(*float64)(unsafe.Add(mBase, uint32(v11)+16)) = v49
				return v11
			}
		}
	}
}
