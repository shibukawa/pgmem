package p1

import (
	base "github.com/shibukawa/pgmem/internal/aot/pgaot/base"
	"math"
	"unsafe"
)

func F_cash_mul_float8(m *base.Module, l0 int64, l1 float64) int64 {
	var v5 float64
	_ = v5
	var v7 float64
	_ = v7
	var v21 float64
	_ = v21
	var v38 int64
	_ = v38
	var v45 int32
	_ = v45
	var v47 int32
	_ = v47
	var v51 int32
	_ = v51
	var v54 int32
	_ = v54
	var v58 int32
	_ = v58
	var v63 int32
	_ = v63
	v5 = base.F64_mul(l1, base.F64_convert_i64_s(l0))
	v7 = math.Float64frombits(uint64(0x7ff0000000000000))
	if base.F64_eq(base.F64_abs(v5), v7)&base.F64_ne(base.F64_abs(l1), v7) == int32(0) {
		if base.F64_ne(v5, float64(0)) != 0 {
			v21 = base.F64_nearest(v5)
			if base.Ui64(int64(9218868437227405312)) < base.Ui64(base.I64_reinterpret_f64(v21)&int64(9223372036854775807)) {
				F_errstart_cold(m, int32(21), int32(0))
				v51 = m.ExcPending
				if v51 != 0 {
					return int64(0)
				} else {
					F_errcode(m, int32(50331778))
					v54 = m.ExcPending
					if v54 != 0 {
						return int64(0)
					} else {
						F_errmsg(m, int32(395009), int32(0))
						v58 = m.ExcPending
						if v58 != 0 {
							return int64(0)
						} else {
							F_errfinish(m, int32(488987), int32(124), int32(538917))
							v63 = m.ExcPending
							if v63 != 0 {
								return int64(0)
							} else {
								base.Wasm_trap_unreachable()
								for {
								}
							}
						}
					}
				}
			} else {
				if base.F64_ge(v21, float64(-9.223372036854776e+18)) == int32(0) {
					F_errstart_cold(m, int32(21), int32(0))
					v51 = m.ExcPending
					if v51 != 0 {
						return int64(0)
					} else {
						F_errcode(m, int32(50331778))
						v54 = m.ExcPending
						if v54 != 0 {
							return int64(0)
						} else {
							F_errmsg(m, int32(395009), int32(0))
							v58 = m.ExcPending
							if v58 != 0 {
								return int64(0)
							} else {
								F_errfinish(m, int32(488987), int32(124), int32(538917))
								v63 = m.ExcPending
								if v63 != 0 {
									return int64(0)
								} else {
									base.Wasm_trap_unreachable()
									for {
									}
								}
							}
						}
					}
				} else {
					if base.F64_lt(v21, float64(9.223372036854776e+18)) == int32(0) {
						F_errstart_cold(m, int32(21), int32(0))
						v51 = m.ExcPending
						if v51 != 0 {
							return int64(0)
						} else {
							F_errcode(m, int32(50331778))
							v54 = m.ExcPending
							if v54 != 0 {
								return int64(0)
							} else {
								F_errmsg(m, int32(395009), int32(0))
								v58 = m.ExcPending
								if v58 != 0 {
									return int64(0)
								} else {
									F_errfinish(m, int32(488987), int32(124), int32(538917))
									v63 = m.ExcPending
									if v63 != 0 {
										return int64(0)
									} else {
										base.Wasm_trap_unreachable()
										for {
										}
									}
								}
							}
						}
					} else {
						if base.F64_lt(base.F64_abs(v21), float64(9.223372036854776e+18)) != 0 {
							v38 = base.I64_trunc_f64_s(v21)
							return v38
						} else {
							return int64(-9223372036854775807 - 1)
						}
					}
				}
			}
		} else {
			if l0 == int64(0) {
				v21 = base.F64_nearest(v5)
				if base.Ui64(int64(9218868437227405312)) < base.Ui64(base.I64_reinterpret_f64(v21)&int64(9223372036854775807)) {
					F_errstart_cold(m, int32(21), int32(0))
					v51 = m.ExcPending
					if v51 != 0 {
						return int64(0)
					} else {
						F_errcode(m, int32(50331778))
						v54 = m.ExcPending
						if v54 != 0 {
							return int64(0)
						} else {
							F_errmsg(m, int32(395009), int32(0))
							v58 = m.ExcPending
							if v58 != 0 {
								return int64(0)
							} else {
								F_errfinish(m, int32(488987), int32(124), int32(538917))
								v63 = m.ExcPending
								if v63 != 0 {
									return int64(0)
								} else {
									base.Wasm_trap_unreachable()
									for {
									}
								}
							}
						}
					}
				} else {
					if base.F64_ge(v21, float64(-9.223372036854776e+18)) == int32(0) {
						F_errstart_cold(m, int32(21), int32(0))
						v51 = m.ExcPending
						if v51 != 0 {
							return int64(0)
						} else {
							F_errcode(m, int32(50331778))
							v54 = m.ExcPending
							if v54 != 0 {
								return int64(0)
							} else {
								F_errmsg(m, int32(395009), int32(0))
								v58 = m.ExcPending
								if v58 != 0 {
									return int64(0)
								} else {
									F_errfinish(m, int32(488987), int32(124), int32(538917))
									v63 = m.ExcPending
									if v63 != 0 {
										return int64(0)
									} else {
										base.Wasm_trap_unreachable()
										for {
										}
									}
								}
							}
						}
					} else {
						if base.F64_lt(v21, float64(9.223372036854776e+18)) == int32(0) {
							F_errstart_cold(m, int32(21), int32(0))
							v51 = m.ExcPending
							if v51 != 0 {
								return int64(0)
							} else {
								F_errcode(m, int32(50331778))
								v54 = m.ExcPending
								if v54 != 0 {
									return int64(0)
								} else {
									F_errmsg(m, int32(395009), int32(0))
									v58 = m.ExcPending
									if v58 != 0 {
										return int64(0)
									} else {
										F_errfinish(m, int32(488987), int32(124), int32(538917))
										v63 = m.ExcPending
										if v63 != 0 {
											return int64(0)
										} else {
											base.Wasm_trap_unreachable()
											for {
											}
										}
									}
								}
							}
						} else {
							if base.F64_lt(base.F64_abs(v21), float64(9.223372036854776e+18)) != 0 {
								v38 = base.I64_trunc_f64_s(v21)
								return v38
							} else {
								return int64(-9223372036854775807 - 1)
							}
						}
					}
				}
			} else {
				if base.F64_ne(l1, float64(0)) != 0 {
					F_float_underflow_error(m)
					v47 = m.ExcPending
					if v47 != 0 {
						return int64(0)
					} else {
						base.Wasm_trap_unreachable()
						for {
						}
					}
				} else {
					v21 = base.F64_nearest(v5)
					if base.Ui64(int64(9218868437227405312)) < base.Ui64(base.I64_reinterpret_f64(v21)&int64(9223372036854775807)) {
						F_errstart_cold(m, int32(21), int32(0))
						v51 = m.ExcPending
						if v51 != 0 {
							return int64(0)
						} else {
							F_errcode(m, int32(50331778))
							v54 = m.ExcPending
							if v54 != 0 {
								return int64(0)
							} else {
								F_errmsg(m, int32(395009), int32(0))
								v58 = m.ExcPending
								if v58 != 0 {
									return int64(0)
								} else {
									F_errfinish(m, int32(488987), int32(124), int32(538917))
									v63 = m.ExcPending
									if v63 != 0 {
										return int64(0)
									} else {
										base.Wasm_trap_unreachable()
										for {
										}
									}
								}
							}
						}
					} else {
						if base.F64_ge(v21, float64(-9.223372036854776e+18)) == int32(0) {
							F_errstart_cold(m, int32(21), int32(0))
							v51 = m.ExcPending
							if v51 != 0 {
								return int64(0)
							} else {
								F_errcode(m, int32(50331778))
								v54 = m.ExcPending
								if v54 != 0 {
									return int64(0)
								} else {
									F_errmsg(m, int32(395009), int32(0))
									v58 = m.ExcPending
									if v58 != 0 {
										return int64(0)
									} else {
										F_errfinish(m, int32(488987), int32(124), int32(538917))
										v63 = m.ExcPending
										if v63 != 0 {
											return int64(0)
										} else {
											base.Wasm_trap_unreachable()
											for {
											}
										}
									}
								}
							}
						} else {
							if base.F64_lt(v21, float64(9.223372036854776e+18)) == int32(0) {
								F_errstart_cold(m, int32(21), int32(0))
								v51 = m.ExcPending
								if v51 != 0 {
									return int64(0)
								} else {
									F_errcode(m, int32(50331778))
									v54 = m.ExcPending
									if v54 != 0 {
										return int64(0)
									} else {
										F_errmsg(m, int32(395009), int32(0))
										v58 = m.ExcPending
										if v58 != 0 {
											return int64(0)
										} else {
											F_errfinish(m, int32(488987), int32(124), int32(538917))
											v63 = m.ExcPending
											if v63 != 0 {
												return int64(0)
											} else {
												base.Wasm_trap_unreachable()
												for {
												}
											}
										}
									}
								}
							} else {
								if base.F64_lt(base.F64_abs(v21), float64(9.223372036854776e+18)) != 0 {
									v38 = base.I64_trunc_f64_s(v21)
									return v38
								} else {
									return int64(-9223372036854775807 - 1)
								}
							}
						}
					}
				}
			}
		}
	} else {
		F_float_overflow_error(m)
		v45 = m.ExcPending
		if v45 != 0 {
			return int64(0)
		} else {
			base.Wasm_trap_unreachable()
			for {
			}
		}
	}
}
func F_cash_ne(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v2 int32
	_ = v2
	var v3 int64
	_ = v3
	var v4 int32
	_ = v4
	var v5 int64
	_ = v5
	v2 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v3 = *(*int64)(unsafe.Add(mBase, uint32(v2)))
	v4 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
	v5 = *(*int64)(unsafe.Add(mBase, uint32(v4)))
	return base.B2i32(v3 != v5)
}
