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
	v9 = *(*float64)(unsafe.Add(mBase, uint32(v8)+8))
	v10 = *(*float64)(unsafe.Add(mBase, uint32(v8)+16))
	v11 = base.F64_add(v9, v10)
	if base.F64_ne(base.F64_abs(v11), math.Float64frombits(uint64(0x7ff0000000000000))) != 0 {
		v21 = *(*float64)(unsafe.Add(mBase, uint32(v7)+8))
		v22 = *(*float64)(unsafe.Add(mBase, uint32(v7)+16))
		v23 = base.F64_sub(v21, v22)
		if base.F64_ne(base.F64_abs(v23), math.Float64frombits(uint64(0x7ff0000000000000))) != 0 {
			return base.F64_lt(base.F64_add(v11, float64(1e-06)), v23)
		} else {
			if base.F64_eq(base.F64_abs(v21), math.Float64frombits(uint64(0x7ff0000000000000))) != 0 {
				return base.F64_lt(base.F64_add(v11, float64(1e-06)), v23)
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
					return base.F64_lt(base.F64_add(v11, float64(1e-06)), v23)
				}
			}
		}
	} else {
		if base.F64_eq(base.F64_abs(v9), math.Float64frombits(uint64(0x7ff0000000000000))) != 0 {
			v21 = *(*float64)(unsafe.Add(mBase, uint32(v7)+8))
			v22 = *(*float64)(unsafe.Add(mBase, uint32(v7)+16))
			v23 = base.F64_sub(v21, v22)
			if base.F64_ne(base.F64_abs(v23), math.Float64frombits(uint64(0x7ff0000000000000))) != 0 {
				return base.F64_lt(base.F64_add(v11, float64(1e-06)), v23)
			} else {
				if base.F64_eq(base.F64_abs(v21), math.Float64frombits(uint64(0x7ff0000000000000))) != 0 {
					return base.F64_lt(base.F64_add(v11, float64(1e-06)), v23)
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
						return base.F64_lt(base.F64_add(v11, float64(1e-06)), v23)
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
				v21 = *(*float64)(unsafe.Add(mBase, uint32(v7)+8))
				v22 = *(*float64)(unsafe.Add(mBase, uint32(v7)+16))
				v23 = base.F64_sub(v21, v22)
				if base.F64_ne(base.F64_abs(v23), math.Float64frombits(uint64(0x7ff0000000000000))) != 0 {
					return base.F64_lt(base.F64_add(v11, float64(1e-06)), v23)
				} else {
					if base.F64_eq(base.F64_abs(v21), math.Float64frombits(uint64(0x7ff0000000000000))) != 0 {
						return base.F64_lt(base.F64_add(v11, float64(1e-06)), v23)
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
							return base.F64_lt(base.F64_add(v11, float64(1e-06)), v23)
						}
					}
				}
			}
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
	var v29 float64
	_ = v29
	var v39 float64
	_ = v39
	var v40 float64
	_ = v40
	var v50 float64
	_ = v50
	var v51 float64
	_ = v51
	var v61 float64
	_ = v61
	var v62 float64
	_ = v62
	var v75 int32
	_ = v75
	var v77 int32
	_ = v77
	v7 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v9 = F_palloc(m, int32(32))
	mBase = m.M
	v12 = m.ExcPending
	if v12 != 0 {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	return int32(0)
L2:
	;
	v13 = *(*float64)(unsafe.Add(mBase, uint32(v7)+16))
	v15 = base.F64_div(v13, float64(1.4142135623730951))
	v16 = base.F64_abs(v15)
	v17 = math.Float64frombits(uint64(0x7ff0000000000000))
	if base.F64_eq(v16, v17)&base.F64_ne(base.F64_abs(v13), v17) != 0 {
		goto L5
	} else {
		goto L6
	}
L3:
	;
	*(*float64)(unsafe.Add(mBase, uint32(v9)+24)) = v62
	return v9
L4:
	;
	F_float_underflow_error(m)
	mBase = m.M
	v77 = m.ExcPending
	if v77 != 0 {
		goto L1
	} else {
		goto L24
	}
L5:
	;
	F_float_overflow_error(m)
	mBase = m.M
	v75 = m.ExcPending
	if v75 != 0 {
		goto L1
	} else {
		goto L23
	}
L6:
	;
	v23 = float64(0)
	if base.F64_eq(v15, v23)&base.F64_ne(v13, v23) != 0 {
		goto L4
	} else {
		goto L7
	}
L7:
	;
	v28 = *(*float64)(unsafe.Add(mBase, uint32(v7)))
	v29 = base.F64_add(v15, v28)
	if base.F64_ne(base.F64_abs(v29), math.Float64frombits(uint64(0x7ff0000000000000))) != 0 {
		goto L8
	} else {
		goto L9
	}
L8:
	;
	*(*float64)(unsafe.Add(mBase, uint32(v9))) = v29
	v39 = *(*float64)(unsafe.Add(mBase, uint32(v7)))
	v40 = base.F64_sub(v39, v15)
	if base.F64_ne(base.F64_abs(v40), math.Float64frombits(uint64(0x7ff0000000000000))) != 0 {
		goto L12
	} else {
		goto L13
	}
L9:
	;
	if base.F64_eq(v16, math.Float64frombits(uint64(0x7ff0000000000000))) != 0 {
		goto L8
	} else {
		goto L10
	}
L10:
	;
	if base.F64_ne(base.F64_abs(v28), math.Float64frombits(uint64(0x7ff0000000000000))) != 0 {
		goto L5
	} else {
		goto L11
	}
L11:
	;
	goto L8
L12:
	;
	*(*float64)(unsafe.Add(mBase, uint32(v9)+16)) = v40
	v50 = *(*float64)(unsafe.Add(mBase, uint32(v7)+8))
	v51 = base.F64_add(v15, v50)
	if base.F64_ne(base.F64_abs(v51), math.Float64frombits(uint64(0x7ff0000000000000))) != 0 {
		goto L16
	} else {
		goto L17
	}
L13:
	;
	if base.F64_eq(v16, math.Float64frombits(uint64(0x7ff0000000000000))) != 0 {
		goto L12
	} else {
		goto L14
	}
L14:
	;
	if base.F64_ne(base.F64_abs(v39), math.Float64frombits(uint64(0x7ff0000000000000))) != 0 {
		goto L5
	} else {
		goto L15
	}
L15:
	;
	goto L12
L16:
	;
	*(*float64)(unsafe.Add(mBase, uint32(v9)+8)) = v51
	v61 = *(*float64)(unsafe.Add(mBase, uint32(v7)+8))
	v62 = base.F64_sub(v61, v15)
	if base.F64_ne(base.F64_abs(v62), math.Float64frombits(uint64(0x7ff0000000000000))) != 0 {
		goto L3
	} else {
		goto L20
	}
L17:
	;
	if base.F64_eq(v16, math.Float64frombits(uint64(0x7ff0000000000000))) != 0 {
		goto L16
	} else {
		goto L18
	}
L18:
	;
	if base.F64_ne(base.F64_abs(v50), math.Float64frombits(uint64(0x7ff0000000000000))) != 0 {
		goto L5
	} else {
		goto L19
	}
L19:
	;
	goto L16
L20:
	;
	if base.F64_eq(v16, math.Float64frombits(uint64(0x7ff0000000000000))) != 0 {
		goto L3
	} else {
		goto L21
	}
L21:
	;
	if base.F64_eq(base.F64_abs(v61), math.Float64frombits(uint64(0x7ff0000000000000))) != 0 {
		goto L3
	} else {
		goto L22
	}
L22:
	;
	goto L5
L23:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L24:
	;
	base.Wasm_trap_unreachable()
	for {
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
	v11 = base.F64_add(v9, v10)
	if base.F64_ne(base.F64_abs(v11), math.Float64frombits(uint64(0x7ff0000000000000))) != 0 {
		v21 = *(*float64)(unsafe.Add(mBase, uint32(v7)))
		v22 = *(*float64)(unsafe.Add(mBase, uint32(v7)+16))
		v23 = base.F64_sub(v21, v22)
		if base.F64_ne(base.F64_abs(v23), math.Float64frombits(uint64(0x7ff0000000000000))) != 0 {
			return base.F64_lt(base.F64_add(v11, float64(1e-06)), v23)
		} else {
			if base.F64_eq(base.F64_abs(v21), math.Float64frombits(uint64(0x7ff0000000000000))) != 0 {
				return base.F64_lt(base.F64_add(v11, float64(1e-06)), v23)
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
					return base.F64_lt(base.F64_add(v11, float64(1e-06)), v23)
				}
			}
		}
	} else {
		if base.F64_eq(base.F64_abs(v9), math.Float64frombits(uint64(0x7ff0000000000000))) != 0 {
			v21 = *(*float64)(unsafe.Add(mBase, uint32(v7)))
			v22 = *(*float64)(unsafe.Add(mBase, uint32(v7)+16))
			v23 = base.F64_sub(v21, v22)
			if base.F64_ne(base.F64_abs(v23), math.Float64frombits(uint64(0x7ff0000000000000))) != 0 {
				return base.F64_lt(base.F64_add(v11, float64(1e-06)), v23)
			} else {
				if base.F64_eq(base.F64_abs(v21), math.Float64frombits(uint64(0x7ff0000000000000))) != 0 {
					return base.F64_lt(base.F64_add(v11, float64(1e-06)), v23)
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
						return base.F64_lt(base.F64_add(v11, float64(1e-06)), v23)
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
					return base.F64_lt(base.F64_add(v11, float64(1e-06)), v23)
				} else {
					if base.F64_eq(base.F64_abs(v21), math.Float64frombits(uint64(0x7ff0000000000000))) != 0 {
						return base.F64_lt(base.F64_add(v11, float64(1e-06)), v23)
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
							return base.F64_lt(base.F64_add(v11, float64(1e-06)), v23)
						}
					}
				}
			}
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
	v9 = *(*float64)(unsafe.Add(mBase, uint32(v8)+8))
	v10 = *(*float64)(unsafe.Add(mBase, uint32(v8)+16))
	v11 = base.F64_sub(v9, v10)
	if base.F64_ne(base.F64_abs(v11), math.Float64frombits(uint64(0x7ff0000000000000))) != 0 {
		v21 = *(*float64)(unsafe.Add(mBase, uint32(v7)+8))
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
			v21 = *(*float64)(unsafe.Add(mBase, uint32(v7)+8))
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
				v21 = *(*float64)(unsafe.Add(mBase, uint32(v7)+8))
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
	v9 = *(*float64)(unsafe.Add(mBase, uint32(v8)+8))
	v10 = *(*float64)(unsafe.Add(mBase, uint32(v8)+16))
	v11 = base.F64_add(v9, v10)
	if base.F64_ne(base.F64_abs(v11), math.Float64frombits(uint64(0x7ff0000000000000))) != 0 {
		v21 = *(*float64)(unsafe.Add(mBase, uint32(v7)+8))
		v22 = *(*float64)(unsafe.Add(mBase, uint32(v7)+16))
		v23 = base.F64_add(v21, v22)
		if base.F64_ne(base.F64_abs(v23), math.Float64frombits(uint64(0x7ff0000000000000))) != 0 {
			return base.F64_le(v11, base.F64_add(v23, float64(1e-06)))
		} else {
			if base.F64_eq(base.F64_abs(v21), math.Float64frombits(uint64(0x7ff0000000000000))) != 0 {
				return base.F64_le(v11, base.F64_add(v23, float64(1e-06)))
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
					return base.F64_le(v11, base.F64_add(v23, float64(1e-06)))
				}
			}
		}
	} else {
		if base.F64_eq(base.F64_abs(v9), math.Float64frombits(uint64(0x7ff0000000000000))) != 0 {
			v21 = *(*float64)(unsafe.Add(mBase, uint32(v7)+8))
			v22 = *(*float64)(unsafe.Add(mBase, uint32(v7)+16))
			v23 = base.F64_add(v21, v22)
			if base.F64_ne(base.F64_abs(v23), math.Float64frombits(uint64(0x7ff0000000000000))) != 0 {
				return base.F64_le(v11, base.F64_add(v23, float64(1e-06)))
			} else {
				if base.F64_eq(base.F64_abs(v21), math.Float64frombits(uint64(0x7ff0000000000000))) != 0 {
					return base.F64_le(v11, base.F64_add(v23, float64(1e-06)))
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
						return base.F64_le(v11, base.F64_add(v23, float64(1e-06)))
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
				v21 = *(*float64)(unsafe.Add(mBase, uint32(v7)+8))
				v22 = *(*float64)(unsafe.Add(mBase, uint32(v7)+16))
				v23 = base.F64_add(v21, v22)
				if base.F64_ne(base.F64_abs(v23), math.Float64frombits(uint64(0x7ff0000000000000))) != 0 {
					return base.F64_le(v11, base.F64_add(v23, float64(1e-06)))
				} else {
					if base.F64_eq(base.F64_abs(v21), math.Float64frombits(uint64(0x7ff0000000000000))) != 0 {
						return base.F64_le(v11, base.F64_add(v23, float64(1e-06)))
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
							return base.F64_le(v11, base.F64_add(v23, float64(1e-06)))
						}
					}
				}
			}
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
	var v26 int32
	_ = v26
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
		if base.F64_ne(base.F64_abs(v15), math.Float64frombits(uint64(0x7ff0000000000000))) != 0 {
			return base.F64_le(v9, base.F64_add(v15, float64(1e-06)))
		} else {
			if base.F64_eq(base.F64_abs(v13), math.Float64frombits(uint64(0x7ff0000000000000))) != 0 {
				return base.F64_le(v9, base.F64_add(v15, float64(1e-06)))
			} else {
				if base.F64_eq(base.F64_abs(v14), math.Float64frombits(uint64(0x7ff0000000000000))) != 0 {
					return base.F64_le(v9, base.F64_add(v15, float64(1e-06)))
				} else {
					F_float_overflow_error(m)
					mBase = m.M
					v26 = m.ExcPending
					if v26 != 0 {
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
	var v27 float64
	_ = v27
	var v28 float64
	_ = v28
	var v29 float64
	_ = v29
	var v41 float64
	_ = v41
	var v48 int32
	_ = v48
	v8 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
	v9 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v11 = F_palloc(m, int32(24))
	mBase = m.M
	v14 = m.ExcPending
	if v14 != 0 {
		return int32(0)
	} else {
		v15 = *(*float64)(unsafe.Add(mBase, uint32(v9)))
		v16 = *(*float64)(unsafe.Add(mBase, uint32(v8)))
		v17 = base.F64_sub(v15, v16)
		if base.F64_ne(base.F64_abs(v17), math.Float64frombits(uint64(0x7ff0000000000000))) != 0 {
			v27 = *(*float64)(unsafe.Add(mBase, uint32(v9)+8))
			v28 = *(*float64)(unsafe.Add(mBase, uint32(v8)+8))
			v29 = base.F64_sub(v27, v28)
			if base.F64_ne(base.F64_abs(v29), math.Float64frombits(uint64(0x7ff0000000000000))) != 0 {
				*(*float64)(unsafe.Add(mBase, uint32(v11)+8)) = v29
				*(*float64)(unsafe.Add(mBase, uint32(v11))) = v17
				v41 = *(*float64)(unsafe.Add(mBase, uint32(v9)+16))
				*(*float64)(unsafe.Add(mBase, uint32(v11)+16)) = v41
				return v11
			} else {
				if base.F64_eq(base.F64_abs(v27), math.Float64frombits(uint64(0x7ff0000000000000))) != 0 {
					*(*float64)(unsafe.Add(mBase, uint32(v11)+8)) = v29
					*(*float64)(unsafe.Add(mBase, uint32(v11))) = v17
					v41 = *(*float64)(unsafe.Add(mBase, uint32(v9)+16))
					*(*float64)(unsafe.Add(mBase, uint32(v11)+16)) = v41
					return v11
				} else {
					if base.F64_ne(base.F64_abs(v28), math.Float64frombits(uint64(0x7ff0000000000000))) != 0 {
						F_float_overflow_error(m)
						mBase = m.M
						v48 = m.ExcPending
						if v48 != 0 {
							return int32(0)
						} else {
							base.Wasm_trap_unreachable()
							for {
							}
						}
					} else {
						*(*float64)(unsafe.Add(mBase, uint32(v11)+8)) = v29
						*(*float64)(unsafe.Add(mBase, uint32(v11))) = v17
						v41 = *(*float64)(unsafe.Add(mBase, uint32(v9)+16))
						*(*float64)(unsafe.Add(mBase, uint32(v11)+16)) = v41
						return v11
					}
				}
			}
		} else {
			if base.F64_eq(base.F64_abs(v15), math.Float64frombits(uint64(0x7ff0000000000000))) != 0 {
				v27 = *(*float64)(unsafe.Add(mBase, uint32(v9)+8))
				v28 = *(*float64)(unsafe.Add(mBase, uint32(v8)+8))
				v29 = base.F64_sub(v27, v28)
				if base.F64_ne(base.F64_abs(v29), math.Float64frombits(uint64(0x7ff0000000000000))) != 0 {
					*(*float64)(unsafe.Add(mBase, uint32(v11)+8)) = v29
					*(*float64)(unsafe.Add(mBase, uint32(v11))) = v17
					v41 = *(*float64)(unsafe.Add(mBase, uint32(v9)+16))
					*(*float64)(unsafe.Add(mBase, uint32(v11)+16)) = v41
					return v11
				} else {
					if base.F64_eq(base.F64_abs(v27), math.Float64frombits(uint64(0x7ff0000000000000))) != 0 {
						*(*float64)(unsafe.Add(mBase, uint32(v11)+8)) = v29
						*(*float64)(unsafe.Add(mBase, uint32(v11))) = v17
						v41 = *(*float64)(unsafe.Add(mBase, uint32(v9)+16))
						*(*float64)(unsafe.Add(mBase, uint32(v11)+16)) = v41
						return v11
					} else {
						if base.F64_ne(base.F64_abs(v28), math.Float64frombits(uint64(0x7ff0000000000000))) != 0 {
							F_float_overflow_error(m)
							mBase = m.M
							v48 = m.ExcPending
							if v48 != 0 {
								return int32(0)
							} else {
								base.Wasm_trap_unreachable()
								for {
								}
							}
						} else {
							*(*float64)(unsafe.Add(mBase, uint32(v11)+8)) = v29
							*(*float64)(unsafe.Add(mBase, uint32(v11))) = v17
							v41 = *(*float64)(unsafe.Add(mBase, uint32(v9)+16))
							*(*float64)(unsafe.Add(mBase, uint32(v11)+16)) = v41
							return v11
						}
					}
				}
			} else {
				if base.F64_ne(base.F64_abs(v16), math.Float64frombits(uint64(0x7ff0000000000000))) != 0 {
					F_float_overflow_error(m)
					mBase = m.M
					v48 = m.ExcPending
					if v48 != 0 {
						return int32(0)
					} else {
						base.Wasm_trap_unreachable()
						for {
						}
					}
				} else {
					v27 = *(*float64)(unsafe.Add(mBase, uint32(v9)+8))
					v28 = *(*float64)(unsafe.Add(mBase, uint32(v8)+8))
					v29 = base.F64_sub(v27, v28)
					if base.F64_ne(base.F64_abs(v29), math.Float64frombits(uint64(0x7ff0000000000000))) != 0 {
						*(*float64)(unsafe.Add(mBase, uint32(v11)+8)) = v29
						*(*float64)(unsafe.Add(mBase, uint32(v11))) = v17
						v41 = *(*float64)(unsafe.Add(mBase, uint32(v9)+16))
						*(*float64)(unsafe.Add(mBase, uint32(v11)+16)) = v41
						return v11
					} else {
						if base.F64_eq(base.F64_abs(v27), math.Float64frombits(uint64(0x7ff0000000000000))) != 0 {
							*(*float64)(unsafe.Add(mBase, uint32(v11)+8)) = v29
							*(*float64)(unsafe.Add(mBase, uint32(v11))) = v17
							v41 = *(*float64)(unsafe.Add(mBase, uint32(v9)+16))
							*(*float64)(unsafe.Add(mBase, uint32(v11)+16)) = v41
							return v11
						} else {
							if base.F64_ne(base.F64_abs(v28), math.Float64frombits(uint64(0x7ff0000000000000))) != 0 {
								F_float_overflow_error(m)
								mBase = m.M
								v48 = m.ExcPending
								if v48 != 0 {
									return int32(0)
								} else {
									base.Wasm_trap_unreachable()
									for {
									}
								}
							} else {
								*(*float64)(unsafe.Add(mBase, uint32(v11)+8)) = v29
								*(*float64)(unsafe.Add(mBase, uint32(v11))) = v17
								v41 = *(*float64)(unsafe.Add(mBase, uint32(v9)+16))
								*(*float64)(unsafe.Add(mBase, uint32(v11)+16)) = v41
								return v11
							}
						}
					}
				}
			}
		}
	}
}
