package p2

import (
	base "github.com/shibukawa/pgmem/internal/aot/pgaot/base"
	"math"
	"unsafe"
)

func F_circle_le(m *base.Module, l0 int32) int32 {
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
									return base.F64_le(v24, base.F64_add(v51, float64(1e-06)))
								}
							}
						}
					}
				}
			}
		}
	}
}
func F_circle_mul_pt(m *base.Module, l0 int32) int32 {
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
	var v16 int32
	_ = v16
	var v17 float64
	_ = v17
	var v18 float64
	_ = v18
	var v19 float64
	_ = v19
	var v20 float64
	_ = v20
	var v23 float64
	_ = v23
	var v24 float64
	_ = v24
	var v29 int64
	_ = v29
	var v35 int32
	_ = v35
	var v36 float64
	_ = v36
	var v37 float64
	_ = v37
	var v40 float64
	_ = v40
	var v45 float64
	_ = v45
	var v52 float64
	_ = v52
	var v56 float64
	_ = v56
	var v58 float64
	_ = v58
	var v69 float64
	_ = v69
	var v86 int32
	_ = v86
	var v92 int32
	_ = v92
	v8 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
	v9 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v11 = F_palloc(m, int32(24))
	mBase = m.M
	v14 = m.ExcPending
	if v14 != 0 {
		return int32(0)
	} else {
		F_point_mul_point(m, v11, v9, v8)
		mBase = m.M
		v16 = m.ExcPending
		if v16 != 0 {
			return int32(0)
		} else {
			v17 = *(*float64)(unsafe.Add(mBase, uint32(v9)+16))
			v18 = math.Float64frombits(uint64(0x7ff0000000000000))
			v19 = *(*float64)(unsafe.Add(mBase, uint32(v8)))
			v20 = base.F64_abs(v19)
			if base.F64_eq(v20, v18) != 0 {
				v52 = v18
				v56 = math.Float64frombits(uint64(0x7ff0000000000000))
				v58 = base.F64_mul(v17, v52)
				if base.B2i32(base.F64_eq(base.F64_abs(v17), v56)|base.F64_ne(base.F64_abs(v58), v56) == int32(0))&base.F64_ne(base.F64_abs(v52), v56) != 0 {
					F_float_overflow_error(m)
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
					v69 = float64(0)
					if base.B2i32(base.F64_eq(v17, v69)|base.F64_ne(v58, v69) == int32(0))&base.F64_ne(v52, v69) != 0 {
						F_float_underflow_error(m)
						mBase = m.M
						v92 = m.ExcPending
						if v92 != 0 {
							return int32(0)
						} else {
							base.Wasm_trap_unreachable()
							for {
							}
						}
					} else {
						*(*float64)(unsafe.Add(mBase, uint32(v11)+16)) = v58
						return v11
					}
				}
			} else {
				v23 = *(*float64)(unsafe.Add(mBase, uint32(v8)+8))
				v24 = base.F64_abs(v23)
				if base.F64_eq(v24, math.Float64frombits(uint64(0x7ff0000000000000))) != 0 {
					v52 = v18
					v56 = math.Float64frombits(uint64(0x7ff0000000000000))
					v58 = base.F64_mul(v17, v52)
					if base.B2i32(base.F64_eq(base.F64_abs(v17), v56)|base.F64_ne(base.F64_abs(v58), v56) == int32(0))&base.F64_ne(base.F64_abs(v52), v56) != 0 {
						F_float_overflow_error(m)
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
						v69 = float64(0)
						if base.B2i32(base.F64_eq(v17, v69)|base.F64_ne(v58, v69) == int32(0))&base.F64_ne(v52, v69) != 0 {
							F_float_underflow_error(m)
							mBase = m.M
							v92 = m.ExcPending
							if v92 != 0 {
								return int32(0)
							} else {
								base.Wasm_trap_unreachable()
								for {
								}
							}
						} else {
							*(*float64)(unsafe.Add(mBase, uint32(v11)+16)) = v58
							return v11
						}
					}
				} else {
					v29 = int64(9218868437227405312)
					if base.B2i32(base.Ui64(v29) < base.Ui64(base.I64_reinterpret_f64(v20)))|base.B2i32(base.Ui64(v29) < base.Ui64(base.I64_reinterpret_f64(v24))) != 0 {
						v52 = math.Float64frombits(uint64(0x7ff8000000000000))
						v56 = math.Float64frombits(uint64(0x7ff0000000000000))
						v58 = base.F64_mul(v17, v52)
						if base.B2i32(base.F64_eq(base.F64_abs(v17), v56)|base.F64_ne(base.F64_abs(v58), v56) == int32(0))&base.F64_ne(base.F64_abs(v52), v56) != 0 {
							F_float_overflow_error(m)
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
							v69 = float64(0)
							if base.B2i32(base.F64_eq(v17, v69)|base.F64_ne(v58, v69) == int32(0))&base.F64_ne(v52, v69) != 0 {
								F_float_underflow_error(m)
								mBase = m.M
								v92 = m.ExcPending
								if v92 != 0 {
									return int32(0)
								} else {
									base.Wasm_trap_unreachable()
									for {
									}
								}
							} else {
								*(*float64)(unsafe.Add(mBase, uint32(v11)+16)) = v58
								return v11
							}
						}
					} else {
						v35 = base.F64_lt(v20, v24)
						if v35 != 0 {
							v36 = v24
						} else {
							v36 = v20
						}
						if v35 != 0 {
							v37 = v20
						} else {
							v37 = v24
						}
						if base.F64_eq(v37, float64(0)) != 0 {
							v52 = v36
							v56 = math.Float64frombits(uint64(0x7ff0000000000000))
							v58 = base.F64_mul(v17, v52)
							if base.B2i32(base.F64_eq(base.F64_abs(v17), v56)|base.F64_ne(base.F64_abs(v58), v56) == int32(0))&base.F64_ne(base.F64_abs(v52), v56) != 0 {
								F_float_overflow_error(m)
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
								v69 = float64(0)
								if base.B2i32(base.F64_eq(v17, v69)|base.F64_ne(v58, v69) == int32(0))&base.F64_ne(v52, v69) != 0 {
									F_float_underflow_error(m)
									mBase = m.M
									v92 = m.ExcPending
									if v92 != 0 {
										return int32(0)
									} else {
										base.Wasm_trap_unreachable()
										for {
										}
									}
								} else {
									*(*float64)(unsafe.Add(mBase, uint32(v11)+16)) = v58
									return v11
								}
							}
						} else {
							v40 = base.F64_div(v37, v36)
							v45 = base.F64_mul(v36, base.F64_sqrt(base.F64_add(base.F64_mul(v40, v40), float64(1))))
							if base.F64_eq(base.F64_abs(v45), math.Float64frombits(uint64(0x7ff0000000000000))) != 0 {
								F_float_overflow_error(m)
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
								if base.F64_eq(v45, float64(0)) != 0 {
									F_float_underflow_error(m)
									mBase = m.M
									v92 = m.ExcPending
									if v92 != 0 {
										return int32(0)
									} else {
										base.Wasm_trap_unreachable()
										for {
										}
									}
								} else {
									v52 = v45
									v56 = math.Float64frombits(uint64(0x7ff0000000000000))
									v58 = base.F64_mul(v17, v52)
									if base.B2i32(base.F64_eq(base.F64_abs(v17), v56)|base.F64_ne(base.F64_abs(v58), v56) == int32(0))&base.F64_ne(base.F64_abs(v52), v56) != 0 {
										F_float_overflow_error(m)
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
										v69 = float64(0)
										if base.B2i32(base.F64_eq(v17, v69)|base.F64_ne(v58, v69) == int32(0))&base.F64_ne(v52, v69) != 0 {
											F_float_underflow_error(m)
											mBase = m.M
											v92 = m.ExcPending
											if v92 != 0 {
												return int32(0)
											} else {
												base.Wasm_trap_unreachable()
												for {
												}
											}
										} else {
											*(*float64)(unsafe.Add(mBase, uint32(v11)+16)) = v58
											return v11
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
