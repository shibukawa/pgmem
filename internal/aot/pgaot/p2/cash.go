package p2

import (
	base "github.com/shibukawa/pgmem/internal/aot/pgaot/base"
	"math"
	"unsafe"
)

func F_cash_div_float8(m *base.Module, l0 int64, l1 float64) int64 {
	var v4 float64
	_ = v4
	var v12 float64
	_ = v12
	var v23 float64
	_ = v23
	var v40 int64
	_ = v40
	var v47 int32
	_ = v47
	var v49 int32
	_ = v49
	var v51 int32
	_ = v51
	var v55 int32
	_ = v55
	var v58 int32
	_ = v58
	var v62 int32
	_ = v62
	var v67 int32
	_ = v67
	v4 = base.F64_convert_i64_s(l0)
	if base.F64_eq(l1, float64(0)) != 0 {
		if base.Ui64(base.I64_reinterpret_f64(v4)&int64(9223372036854775807)) <= base.Ui64(int64(9218868437227405312)) {
			F_float_zero_divide_error(m)
			v47 = m.ExcPending
			if v47 != 0 {
				return int64(0)
			} else {
				base.Wasm_trap_unreachable()
				for {
				}
			}
		} else {
			v12 = base.F64_div(v4, l1)
			if base.F64_eq(base.F64_abs(v12), math.Float64frombits(uint64(0x7ff0000000000000))) != 0 {
				F_float_overflow_error(m)
				v49 = m.ExcPending
				if v49 != 0 {
					return int64(0)
				} else {
					base.Wasm_trap_unreachable()
					for {
					}
				}
			} else {
				if base.F64_ne(v12, float64(0)) != 0 {
					v23 = base.F64_nearest(v12)
					if base.Ui64(int64(9218868437227405312)) < base.Ui64(base.I64_reinterpret_f64(v23)&int64(9223372036854775807)) {
						F_errstart_cold(m, int32(21), int32(0))
						v55 = m.ExcPending
						if v55 != 0 {
							return int64(0)
						} else {
							F_errcode(m, int32(50331778))
							v58 = m.ExcPending
							if v58 != 0 {
								return int64(0)
							} else {
								F_errmsg(m, int32(399601), int32(0))
								v62 = m.ExcPending
								if v62 != 0 {
									return int64(0)
								} else {
									F_errfinish(m, int32(494794), int32(137), int32(545183))
									v67 = m.ExcPending
									if v67 != 0 {
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
						if base.F64_ge(v23, float64(-9.223372036854776e+18)) == int32(0) {
							F_errstart_cold(m, int32(21), int32(0))
							v55 = m.ExcPending
							if v55 != 0 {
								return int64(0)
							} else {
								F_errcode(m, int32(50331778))
								v58 = m.ExcPending
								if v58 != 0 {
									return int64(0)
								} else {
									F_errmsg(m, int32(399601), int32(0))
									v62 = m.ExcPending
									if v62 != 0 {
										return int64(0)
									} else {
										F_errfinish(m, int32(494794), int32(137), int32(545183))
										v67 = m.ExcPending
										if v67 != 0 {
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
							if base.F64_lt(v23, float64(9.223372036854776e+18)) == int32(0) {
								F_errstart_cold(m, int32(21), int32(0))
								v55 = m.ExcPending
								if v55 != 0 {
									return int64(0)
								} else {
									F_errcode(m, int32(50331778))
									v58 = m.ExcPending
									if v58 != 0 {
										return int64(0)
									} else {
										F_errmsg(m, int32(399601), int32(0))
										v62 = m.ExcPending
										if v62 != 0 {
											return int64(0)
										} else {
											F_errfinish(m, int32(494794), int32(137), int32(545183))
											v67 = m.ExcPending
											if v67 != 0 {
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
								if base.F64_lt(base.F64_abs(v23), float64(9.223372036854776e+18)) != 0 {
									v40 = base.I64_trunc_f64_s(v23)
									return v40
								} else {
									return int64(-9223372036854775807 - 1)
								}
							}
						}
					}
				} else {
					if l0 == int64(0) {
						v23 = base.F64_nearest(v12)
						if base.Ui64(int64(9218868437227405312)) < base.Ui64(base.I64_reinterpret_f64(v23)&int64(9223372036854775807)) {
							F_errstart_cold(m, int32(21), int32(0))
							v55 = m.ExcPending
							if v55 != 0 {
								return int64(0)
							} else {
								F_errcode(m, int32(50331778))
								v58 = m.ExcPending
								if v58 != 0 {
									return int64(0)
								} else {
									F_errmsg(m, int32(399601), int32(0))
									v62 = m.ExcPending
									if v62 != 0 {
										return int64(0)
									} else {
										F_errfinish(m, int32(494794), int32(137), int32(545183))
										v67 = m.ExcPending
										if v67 != 0 {
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
							if base.F64_ge(v23, float64(-9.223372036854776e+18)) == int32(0) {
								F_errstart_cold(m, int32(21), int32(0))
								v55 = m.ExcPending
								if v55 != 0 {
									return int64(0)
								} else {
									F_errcode(m, int32(50331778))
									v58 = m.ExcPending
									if v58 != 0 {
										return int64(0)
									} else {
										F_errmsg(m, int32(399601), int32(0))
										v62 = m.ExcPending
										if v62 != 0 {
											return int64(0)
										} else {
											F_errfinish(m, int32(494794), int32(137), int32(545183))
											v67 = m.ExcPending
											if v67 != 0 {
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
								if base.F64_lt(v23, float64(9.223372036854776e+18)) == int32(0) {
									F_errstart_cold(m, int32(21), int32(0))
									v55 = m.ExcPending
									if v55 != 0 {
										return int64(0)
									} else {
										F_errcode(m, int32(50331778))
										v58 = m.ExcPending
										if v58 != 0 {
											return int64(0)
										} else {
											F_errmsg(m, int32(399601), int32(0))
											v62 = m.ExcPending
											if v62 != 0 {
												return int64(0)
											} else {
												F_errfinish(m, int32(494794), int32(137), int32(545183))
												v67 = m.ExcPending
												if v67 != 0 {
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
									if base.F64_lt(base.F64_abs(v23), float64(9.223372036854776e+18)) != 0 {
										v40 = base.I64_trunc_f64_s(v23)
										return v40
									} else {
										return int64(-9223372036854775807 - 1)
									}
								}
							}
						}
					} else {
						if base.F64_ne(base.F64_abs(l1), math.Float64frombits(uint64(0x7ff0000000000000))) != 0 {
							F_float_underflow_error(m)
							v51 = m.ExcPending
							if v51 != 0 {
								return int64(0)
							} else {
								base.Wasm_trap_unreachable()
								for {
								}
							}
						} else {
							v23 = base.F64_nearest(v12)
							if base.Ui64(int64(9218868437227405312)) < base.Ui64(base.I64_reinterpret_f64(v23)&int64(9223372036854775807)) {
								F_errstart_cold(m, int32(21), int32(0))
								v55 = m.ExcPending
								if v55 != 0 {
									return int64(0)
								} else {
									F_errcode(m, int32(50331778))
									v58 = m.ExcPending
									if v58 != 0 {
										return int64(0)
									} else {
										F_errmsg(m, int32(399601), int32(0))
										v62 = m.ExcPending
										if v62 != 0 {
											return int64(0)
										} else {
											F_errfinish(m, int32(494794), int32(137), int32(545183))
											v67 = m.ExcPending
											if v67 != 0 {
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
								if base.F64_ge(v23, float64(-9.223372036854776e+18)) == int32(0) {
									F_errstart_cold(m, int32(21), int32(0))
									v55 = m.ExcPending
									if v55 != 0 {
										return int64(0)
									} else {
										F_errcode(m, int32(50331778))
										v58 = m.ExcPending
										if v58 != 0 {
											return int64(0)
										} else {
											F_errmsg(m, int32(399601), int32(0))
											v62 = m.ExcPending
											if v62 != 0 {
												return int64(0)
											} else {
												F_errfinish(m, int32(494794), int32(137), int32(545183))
												v67 = m.ExcPending
												if v67 != 0 {
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
									if base.F64_lt(v23, float64(9.223372036854776e+18)) == int32(0) {
										F_errstart_cold(m, int32(21), int32(0))
										v55 = m.ExcPending
										if v55 != 0 {
											return int64(0)
										} else {
											F_errcode(m, int32(50331778))
											v58 = m.ExcPending
											if v58 != 0 {
												return int64(0)
											} else {
												F_errmsg(m, int32(399601), int32(0))
												v62 = m.ExcPending
												if v62 != 0 {
													return int64(0)
												} else {
													F_errfinish(m, int32(494794), int32(137), int32(545183))
													v67 = m.ExcPending
													if v67 != 0 {
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
										if base.F64_lt(base.F64_abs(v23), float64(9.223372036854776e+18)) != 0 {
											v40 = base.I64_trunc_f64_s(v23)
											return v40
										} else {
											return int64(-9223372036854775807 - 1)
										}
									}
								}
							}
						}
					}
				}
			}
		}
	} else {
		v12 = base.F64_div(v4, l1)
		if base.F64_eq(base.F64_abs(v12), math.Float64frombits(uint64(0x7ff0000000000000))) != 0 {
			F_float_overflow_error(m)
			v49 = m.ExcPending
			if v49 != 0 {
				return int64(0)
			} else {
				base.Wasm_trap_unreachable()
				for {
				}
			}
		} else {
			if base.F64_ne(v12, float64(0)) != 0 {
				v23 = base.F64_nearest(v12)
				if base.Ui64(int64(9218868437227405312)) < base.Ui64(base.I64_reinterpret_f64(v23)&int64(9223372036854775807)) {
					F_errstart_cold(m, int32(21), int32(0))
					v55 = m.ExcPending
					if v55 != 0 {
						return int64(0)
					} else {
						F_errcode(m, int32(50331778))
						v58 = m.ExcPending
						if v58 != 0 {
							return int64(0)
						} else {
							F_errmsg(m, int32(399601), int32(0))
							v62 = m.ExcPending
							if v62 != 0 {
								return int64(0)
							} else {
								F_errfinish(m, int32(494794), int32(137), int32(545183))
								v67 = m.ExcPending
								if v67 != 0 {
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
					if base.F64_ge(v23, float64(-9.223372036854776e+18)) == int32(0) {
						F_errstart_cold(m, int32(21), int32(0))
						v55 = m.ExcPending
						if v55 != 0 {
							return int64(0)
						} else {
							F_errcode(m, int32(50331778))
							v58 = m.ExcPending
							if v58 != 0 {
								return int64(0)
							} else {
								F_errmsg(m, int32(399601), int32(0))
								v62 = m.ExcPending
								if v62 != 0 {
									return int64(0)
								} else {
									F_errfinish(m, int32(494794), int32(137), int32(545183))
									v67 = m.ExcPending
									if v67 != 0 {
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
						if base.F64_lt(v23, float64(9.223372036854776e+18)) == int32(0) {
							F_errstart_cold(m, int32(21), int32(0))
							v55 = m.ExcPending
							if v55 != 0 {
								return int64(0)
							} else {
								F_errcode(m, int32(50331778))
								v58 = m.ExcPending
								if v58 != 0 {
									return int64(0)
								} else {
									F_errmsg(m, int32(399601), int32(0))
									v62 = m.ExcPending
									if v62 != 0 {
										return int64(0)
									} else {
										F_errfinish(m, int32(494794), int32(137), int32(545183))
										v67 = m.ExcPending
										if v67 != 0 {
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
							if base.F64_lt(base.F64_abs(v23), float64(9.223372036854776e+18)) != 0 {
								v40 = base.I64_trunc_f64_s(v23)
								return v40
							} else {
								return int64(-9223372036854775807 - 1)
							}
						}
					}
				}
			} else {
				if l0 == int64(0) {
					v23 = base.F64_nearest(v12)
					if base.Ui64(int64(9218868437227405312)) < base.Ui64(base.I64_reinterpret_f64(v23)&int64(9223372036854775807)) {
						F_errstart_cold(m, int32(21), int32(0))
						v55 = m.ExcPending
						if v55 != 0 {
							return int64(0)
						} else {
							F_errcode(m, int32(50331778))
							v58 = m.ExcPending
							if v58 != 0 {
								return int64(0)
							} else {
								F_errmsg(m, int32(399601), int32(0))
								v62 = m.ExcPending
								if v62 != 0 {
									return int64(0)
								} else {
									F_errfinish(m, int32(494794), int32(137), int32(545183))
									v67 = m.ExcPending
									if v67 != 0 {
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
						if base.F64_ge(v23, float64(-9.223372036854776e+18)) == int32(0) {
							F_errstart_cold(m, int32(21), int32(0))
							v55 = m.ExcPending
							if v55 != 0 {
								return int64(0)
							} else {
								F_errcode(m, int32(50331778))
								v58 = m.ExcPending
								if v58 != 0 {
									return int64(0)
								} else {
									F_errmsg(m, int32(399601), int32(0))
									v62 = m.ExcPending
									if v62 != 0 {
										return int64(0)
									} else {
										F_errfinish(m, int32(494794), int32(137), int32(545183))
										v67 = m.ExcPending
										if v67 != 0 {
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
							if base.F64_lt(v23, float64(9.223372036854776e+18)) == int32(0) {
								F_errstart_cold(m, int32(21), int32(0))
								v55 = m.ExcPending
								if v55 != 0 {
									return int64(0)
								} else {
									F_errcode(m, int32(50331778))
									v58 = m.ExcPending
									if v58 != 0 {
										return int64(0)
									} else {
										F_errmsg(m, int32(399601), int32(0))
										v62 = m.ExcPending
										if v62 != 0 {
											return int64(0)
										} else {
											F_errfinish(m, int32(494794), int32(137), int32(545183))
											v67 = m.ExcPending
											if v67 != 0 {
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
								if base.F64_lt(base.F64_abs(v23), float64(9.223372036854776e+18)) != 0 {
									v40 = base.I64_trunc_f64_s(v23)
									return v40
								} else {
									return int64(-9223372036854775807 - 1)
								}
							}
						}
					}
				} else {
					if base.F64_ne(base.F64_abs(l1), math.Float64frombits(uint64(0x7ff0000000000000))) != 0 {
						F_float_underflow_error(m)
						v51 = m.ExcPending
						if v51 != 0 {
							return int64(0)
						} else {
							base.Wasm_trap_unreachable()
							for {
							}
						}
					} else {
						v23 = base.F64_nearest(v12)
						if base.Ui64(int64(9218868437227405312)) < base.Ui64(base.I64_reinterpret_f64(v23)&int64(9223372036854775807)) {
							F_errstart_cold(m, int32(21), int32(0))
							v55 = m.ExcPending
							if v55 != 0 {
								return int64(0)
							} else {
								F_errcode(m, int32(50331778))
								v58 = m.ExcPending
								if v58 != 0 {
									return int64(0)
								} else {
									F_errmsg(m, int32(399601), int32(0))
									v62 = m.ExcPending
									if v62 != 0 {
										return int64(0)
									} else {
										F_errfinish(m, int32(494794), int32(137), int32(545183))
										v67 = m.ExcPending
										if v67 != 0 {
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
							if base.F64_ge(v23, float64(-9.223372036854776e+18)) == int32(0) {
								F_errstart_cold(m, int32(21), int32(0))
								v55 = m.ExcPending
								if v55 != 0 {
									return int64(0)
								} else {
									F_errcode(m, int32(50331778))
									v58 = m.ExcPending
									if v58 != 0 {
										return int64(0)
									} else {
										F_errmsg(m, int32(399601), int32(0))
										v62 = m.ExcPending
										if v62 != 0 {
											return int64(0)
										} else {
											F_errfinish(m, int32(494794), int32(137), int32(545183))
											v67 = m.ExcPending
											if v67 != 0 {
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
								if base.F64_lt(v23, float64(9.223372036854776e+18)) == int32(0) {
									F_errstart_cold(m, int32(21), int32(0))
									v55 = m.ExcPending
									if v55 != 0 {
										return int64(0)
									} else {
										F_errcode(m, int32(50331778))
										v58 = m.ExcPending
										if v58 != 0 {
											return int64(0)
										} else {
											F_errmsg(m, int32(399601), int32(0))
											v62 = m.ExcPending
											if v62 != 0 {
												return int64(0)
											} else {
												F_errfinish(m, int32(494794), int32(137), int32(545183))
												v67 = m.ExcPending
												if v67 != 0 {
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
									if base.F64_lt(base.F64_abs(v23), float64(9.223372036854776e+18)) != 0 {
										v40 = base.I64_trunc_f64_s(v23)
										return v40
									} else {
										return int64(-9223372036854775807 - 1)
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
func F_cash_div_flt4(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v2 int32
	_ = v2
	var v3 int64
	_ = v3
	var v4 float32
	_ = v4
	var v6 int64
	_ = v6
	var v9 int32
	_ = v9
	var v10 int32
	_ = v10
	var v11 int32
	_ = v11
	v2 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v3 = *(*int64)(unsafe.Add(mBase, uint32(v2)))
	v4 = *(*float32)(unsafe.Add(mBase, uint32(l0)+28))
	v6 = F_cash_div_float8(m, v3, base.F64_promote_f32(v4))
	mBase = m.M
	v9 = m.ExcPending
	if v9 != 0 {
		return int32(0)
	} else {
		v10 = F_Int64GetDatum(m, v6)
		mBase = m.M
		v11 = m.ExcPending
		if v11 != 0 {
			return int32(0)
		} else {
			return v10
		}
	}
}
func F_cash_lt(m *base.Module, l0 int32) int32 {
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
	return base.B2i32(v3 < v5)
}
func F_cash_mi(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v4 int32
	_ = v4
	var v5 int64
	_ = v5
	var v8 int32
	_ = v8
	var v9 int64
	_ = v9
	var v10 int64
	_ = v10
	var v18 int32
	_ = v18
	var v21 int32
	_ = v21
	var v25 int32
	_ = v25
	var v30 int32
	_ = v30
	var v31 int32
	_ = v31
	var v32 int32
	_ = v32
	v4 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
	v5 = *(*int64)(unsafe.Add(mBase, uint32(v4)))
	v8 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v9 = *(*int64)(unsafe.Add(mBase, uint32(v8)))
	v10 = v9 - v5
	if base.B2i32(int64(0) < v5) != base.B2i32(v10 < v9) {
		F_errstart_cold(m, int32(21), int32(0))
		mBase = m.M
		v18 = m.ExcPending
		if v18 != 0 {
			return int32(0)
		} else {
			F_errcode(m, int32(50331778))
			mBase = m.M
			v21 = m.ExcPending
			if v21 != 0 {
				return int32(0)
			} else {
				F_errmsg(m, int32(399601), int32(0))
				mBase = m.M
				v25 = m.ExcPending
				if v25 != 0 {
					return int32(0)
				} else {
					F_errfinish(m, int32(494794), int32(111), int32(321706))
					mBase = m.M
					v30 = m.ExcPending
					if v30 != 0 {
						return int32(0)
					} else {
						base.Wasm_trap_unreachable()
						for {
						}
					}
				}
			}
		}
	} else {
		v31 = F_Int64GetDatum(m, v10)
		mBase = m.M
		v32 = m.ExcPending
		if v32 != 0 {
			return int32(0)
		} else {
			return v31
		}
	}
}
func F_cash_mul_int2(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v4 int32
	_ = v4
	var v6 int32
	_ = v6
	var v8 int32
	_ = v8
	var v9 int64
	_ = v9
	var v10 int64
	_ = v10
	var v12 int64
	_ = v12
	var v19 int64
	_ = v19
	var v20 int64
	_ = v20
	var v22 int64
	_ = v22
	var v25 int64
	_ = v25
	var v26 int64
	_ = v26
	var v28 int64
	_ = v28
	var v29 int64
	_ = v29
	var v33 int64
	_ = v33
	var v40 int64
	_ = v40
	var v51 int64
	_ = v51
	var v52 int64
	_ = v52
	var v61 int32
	_ = v61
	var v64 int32
	_ = v64
	var v68 int32
	_ = v68
	var v73 int32
	_ = v73
	var v74 int32
	_ = v74
	var v75 int32
	_ = v75
	v4 = m.G0
	v6 = v4 - int32(16)
	m.G0 = v6
	v8 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v9 = *(*int64)(unsafe.Add(mBase, uint32(v8)))
	v10 = int64(63)
	v12 = int64(*(*int16)(unsafe.Add(mBase, uint32(l0)+28)))
	v19 = int64(32)
	v20 = int64(base.Ui64(v12) >> (uint(v19) % 64))
	v22 = int64(base.Ui64(v9) >> (uint(v19) % 64))
	v25 = int64(4294967295)
	v26 = v12 & v25
	v28 = v9 & v25
	v29 = v26 * v28
	v33 = int64(base.Ui64(v29)>>(uint(v19)%64)) + v26*v22
	v40 = v28*v20 + v33&v25
	*(*int64)(unsafe.Add(mBase, uint32(v6)+8)) = v9*(v12>>(uint(v10)%64)) + v9>>(uint(v10)%64)*v12 + v20*v22 + int64(base.Ui64(v33)>>(uint(v19)%64)) + int64(base.Ui64(v40)>>(uint(v19)%64))
	*(*int64)(unsafe.Add(mBase, uint32(v6))) = v29&v25 | v40<<(uint(v19)%64)
	v51 = *(*int64)(unsafe.Add(mBase, uint32(v6)+8))
	v52 = *(*int64)(unsafe.Add(mBase, uint32(v6)))
	if v51 != v52>>(uint(int64(63))%64) {
		F_errstart_cold(m, int32(21), int32(0))
		mBase = m.M
		v61 = m.ExcPending
		if v61 != 0 {
			return int32(0)
		} else {
			F_errcode(m, int32(50331778))
			mBase = m.M
			v64 = m.ExcPending
			if v64 != 0 {
				return int32(0)
			} else {
				F_errmsg(m, int32(399601), int32(0))
				mBase = m.M
				v68 = m.ExcPending
				if v68 != 0 {
					return int32(0)
				} else {
					F_errfinish(m, int32(494794), int32(150), int32(548618))
					mBase = m.M
					v73 = m.ExcPending
					if v73 != 0 {
						return int32(0)
					} else {
						base.Wasm_trap_unreachable()
						for {
						}
					}
				}
			}
		}
	} else {
		v74 = F_Int64GetDatum(m, v52)
		mBase = m.M
		v75 = m.ExcPending
		if v75 != 0 {
			return int32(0)
		} else {
			m.G0 = v6 + int32(16)
			return v74
		}
	}
}
func F_cash_mul_int8(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v4 int32
	_ = v4
	var v6 int32
	_ = v6
	var v8 int32
	_ = v8
	var v9 int64
	_ = v9
	var v10 int64
	_ = v10
	var v12 int32
	_ = v12
	var v13 int64
	_ = v13
	var v20 int64
	_ = v20
	var v21 int64
	_ = v21
	var v23 int64
	_ = v23
	var v26 int64
	_ = v26
	var v27 int64
	_ = v27
	var v29 int64
	_ = v29
	var v30 int64
	_ = v30
	var v34 int64
	_ = v34
	var v41 int64
	_ = v41
	var v52 int64
	_ = v52
	var v53 int64
	_ = v53
	var v62 int32
	_ = v62
	var v65 int32
	_ = v65
	var v69 int32
	_ = v69
	var v74 int32
	_ = v74
	var v75 int32
	_ = v75
	var v76 int32
	_ = v76
	v4 = m.G0
	v6 = v4 - int32(16)
	m.G0 = v6
	v8 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v9 = *(*int64)(unsafe.Add(mBase, uint32(v8)))
	v10 = int64(63)
	v12 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
	v13 = *(*int64)(unsafe.Add(mBase, uint32(v12)))
	v20 = int64(32)
	v21 = int64(base.Ui64(v13) >> (uint(v20) % 64))
	v23 = int64(base.Ui64(v9) >> (uint(v20) % 64))
	v26 = int64(4294967295)
	v27 = v13 & v26
	v29 = v9 & v26
	v30 = v27 * v29
	v34 = int64(base.Ui64(v30)>>(uint(v20)%64)) + v27*v23
	v41 = v29*v21 + v34&v26
	*(*int64)(unsafe.Add(mBase, uint32(v6)+8)) = v9*(v13>>(uint(v10)%64)) + v9>>(uint(v10)%64)*v13 + v21*v23 + int64(base.Ui64(v34)>>(uint(v20)%64)) + int64(base.Ui64(v41)>>(uint(v20)%64))
	*(*int64)(unsafe.Add(mBase, uint32(v6))) = v30&v26 | v41<<(uint(v20)%64)
	v52 = *(*int64)(unsafe.Add(mBase, uint32(v6)+8))
	v53 = *(*int64)(unsafe.Add(mBase, uint32(v6)))
	if v52 != v53>>(uint(int64(63))%64) {
		F_errstart_cold(m, int32(21), int32(0))
		mBase = m.M
		v62 = m.ExcPending
		if v62 != 0 {
			return int32(0)
		} else {
			F_errcode(m, int32(50331778))
			mBase = m.M
			v65 = m.ExcPending
			if v65 != 0 {
				return int32(0)
			} else {
				F_errmsg(m, int32(399601), int32(0))
				mBase = m.M
				v69 = m.ExcPending
				if v69 != 0 {
					return int32(0)
				} else {
					F_errfinish(m, int32(494794), int32(150), int32(548618))
					mBase = m.M
					v74 = m.ExcPending
					if v74 != 0 {
						return int32(0)
					} else {
						base.Wasm_trap_unreachable()
						for {
						}
					}
				}
			}
		}
	} else {
		v75 = F_Int64GetDatum(m, v53)
		mBase = m.M
		v76 = m.ExcPending
		if v76 != 0 {
			return int32(0)
		} else {
			m.G0 = v6 + int32(16)
			return v75
		}
	}
}
func F_cash_out(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v18 int32
	_ = v18
	var v20 int32
	_ = v20
	var v22 int32
	_ = v22
	var v23 int64
	_ = v23
	var v24 int32
	_ = v24
	var v27 int32
	_ = v27
	var v28 int32
	_ = v28
	var v31 int32
	_ = v31
	var v32 int32
	_ = v32
	var v39 int32
	_ = v39
	var v40 int32
	_ = v40
	var v41 int32
	_ = v41
	var v42 int32
	_ = v42
	var v45 int32
	_ = v45
	var v52 int32
	_ = v52
	var v53 int32
	_ = v53
	var v54 int32
	_ = v54
	var v56 int32
	_ = v56
	var v58 int32
	_ = v58
	var v59 int32
	_ = v59
	var v60 int32
	_ = v60
	var v61 int32
	_ = v61
	var v62 int32
	_ = v62
	var v68 int32
	_ = v68
	var v70 int32
	_ = v70
	var v71 int32
	_ = v71
	var v75 int32
	_ = v75
	var v77 int32
	_ = v77
	var v78 int32
	_ = v78
	var v79 int32
	_ = v79
	var v80 int32
	_ = v80
	var v84 int32
	_ = v84
	var v85 int32
	_ = v85
	var v87 int32
	_ = v87
	var v89 int32
	_ = v89
	var v91 int32
	_ = v91
	var v92 int32
	_ = v92
	var v95 int64
	_ = v95
	var v100 int32
	_ = v100
	var v102 int32
	_ = v102
	var v116 int64
	_ = v116
	var v120 int32
	_ = v120
	var v124 int32
	_ = v124
	var v125 int32
	_ = v125
	var v126 int32
	_ = v126
	var v127 int32
	_ = v127
	var v129 int32
	_ = v129
	var v131 int32
	_ = v131
	var v132 int32
	_ = v132
	var v133 int64
	_ = v133
	var v134 int64
	_ = v134
	var v140 int32
	_ = v140
	var v143 int32
	_ = v143
	var v154 int32
	_ = v154
	var v161 int32
	_ = v161
	var v162 int32
	_ = v162
	var v169 int32
	_ = v169
	var v170 int32
	_ = v170
	var v174 int32
	_ = v174
	var v177 int32
	_ = v177
	var v182 int32
	_ = v182
	var v191 int32
	_ = v191
	var v192 int32
	_ = v192
	var v199 int32
	_ = v199
	var v200 int32
	_ = v200
	var v204 int32
	_ = v204
	var v207 int32
	_ = v207
	var v212 int32
	_ = v212
	var v221 int32
	_ = v221
	var v222 int32
	_ = v222
	var v231 int32
	_ = v231
	var v232 int32
	_ = v232
	var v239 int32
	_ = v239
	var v242 int32
	_ = v242
	var v248 int32
	_ = v248
	var v253 int32
	_ = v253
	var v254 int32
	_ = v254
	var v261 int32
	_ = v261
	var v264 int32
	_ = v264
	var v270 int32
	_ = v270
	var v275 int32
	_ = v275
	var v276 int32
	_ = v276
	var v283 int32
	_ = v283
	var v286 int32
	_ = v286
	var v292 int32
	_ = v292
	var v297 int32
	_ = v297
	var v298 int32
	_ = v298
	var v305 int32
	_ = v305
	var v308 int32
	_ = v308
	var v314 int32
	_ = v314
	var v319 int32
	_ = v319
	var v320 int32
	_ = v320
	var v323 int32
	_ = v323
	v18 = m.G0
	v20 = v18 - int32(416)
	m.G0 = v20
	v22 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v23 = *(*int64)(unsafe.Add(mBase, uint32(v22)))
	v24 = F_PGLC_localeconv(m)
	mBase = m.M
	v27 = m.ExcPending
	if v27 != 0 {
		return int32(0)
	} else {
		v28 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v24)+41)))
		v31 = *(*int32)(unsafe.Add(mBase, uint32(v24)+28))
		v32 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v31))))
		v39 = int32(651519)
		v40 = int32(46)
		v41 = *(*int32)(unsafe.Add(mBase, uint32(v24)+20))
		v42 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v41))))
		if v42 == int32(0) {
			v53 = v39
			v54 = v40
		} else {
			v45 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v41)+1)))
			if v45 != 0 {
				v53 = v39
				v54 = v40
			} else {
				if v42&int32(255) == int32(44) {
					v52 = int32(651483)
				} else {
					v52 = int32(651519)
				}
				v53 = v52
				v54 = v42
			}
		}
		if base.Ui32(int32(10)) < base.Ui32(v28) {
			v56 = int32(2)
		} else {
			v56 = v28
		}
		if base.Ui32((v32-int32(7))&int32(255)) < base.Ui32(int32(250)) {
			v58 = int32(3)
		} else {
			v58 = v32
		}
		v59 = *(*int32)(unsafe.Add(mBase, uint32(v24)+16))
		v60 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v59))))
		v61 = *(*int32)(unsafe.Add(mBase, uint32(v24)+24))
		v62 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v61))))
		if v23 < int64(0) {
			v68 = *(*int32)(unsafe.Add(mBase, uint32(v24)+36))
			v70 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v68))))
			if v70 != 0 {
				v71 = v68
			} else {
				v71 = int32(651509)
			}
			v77 = int32(44)
			v78 = int32(45)
			v79 = int32(47)
			v80 = v71
		} else {
			v75 = *(*int32)(unsafe.Add(mBase, uint32(v24)+32))
			v77 = int32(42)
			v78 = int32(43)
			v79 = int32(46)
			v80 = v75
		}
		if v60 != 0 {
			v84 = v59
		} else {
			v84 = int32(671510)
		}
		if v62 != 0 {
			v85 = v61
		} else {
			v85 = v53
		}
		v87 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v24+v79))))
		v89 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v24+v77))))
		v91 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v24+v78))))
		v92 = int32(0)
		*(*uint8)(unsafe.Add(mBase, uint32(v20)+415)) = uint8(v92)
		v95 = v23 >> (uint(int64(63)) % 64)
		v100 = base.I32_extend8_s(v56)
		v102 = v20 + int32(415)
		v116 = v23 ^ v95 - v95
		for {
			if v56 == int32(0) {
				if int32(0) <= v100 {
					v129 = v102
				} else {
					v124 = base.I32_rem_s(v100, base.I32_extend8_s(v58))
					if v124 != 0 {
						v129 = v102
					} else {
						v125 = F_strlen(m, v85)
						mBase = m.M
						v126 = v102 - v125
						if v125 != 0 {
							v127 = F__emscripten_memcpy_bulkmem(m, v126, v85, v125)
							mBase = m.M
						} else {
						}
						v129 = v126
					}
				}
			} else {
				if v100 != 0 {
					if int32(0) <= v100 {
						v129 = v102
					} else {
						v124 = base.I32_rem_s(v100, base.I32_extend8_s(v58))
						if v124 != 0 {
							v129 = v102
						} else {
							v125 = F_strlen(m, v85)
							mBase = m.M
							v126 = v102 - v125
							if v125 != 0 {
								v127 = F__emscripten_memcpy_bulkmem(m, v126, v85, v125)
								mBase = m.M
							} else {
							}
							v129 = v126
						}
					}
				} else {
					v120 = v102 - int32(1)
					*(*uint8)(unsafe.Add(mBase, uint32(v120))) = uint8(v54)
					v129 = v120
				}
			}
			v131 = int32(1)
			v132 = v129 - v131
			v133 = int64(10)
			v134 = base.I64_div_u_s(v116, v133)
			v140 = base.I32_wrap_i64(v116-v134*v133) | int32(48)
			*(*uint8)(unsafe.Add(mBase, uint32(v132))) = uint8(v140)
			v143 = v100 - v131
			if base.Ui64(int64(9)) < base.Ui64(v116) {
				v100 = v143
				v102 = v132
				v116 = v134
				continue
			} else {
			}
			if int32(0) <= v143 {
				v100 = v143
				v102 = v132
				v116 = v134
				continue
			} else {
				break
			}
			break
		}
		switch v87 {
		case 0:
			if v91&int32(255) == int32(1) {
				v154 = int32(727670)
			} else {
				v154 = int32(738731)
			}
			if v89 != 0 {
				*(*int32)(unsafe.Add(mBase, uint32(v20)+88)) = v132
				*(*int32)(unsafe.Add(mBase, uint32(v20)+84)) = v154
				*(*int32)(unsafe.Add(mBase, uint32(v20)+80)) = v84
				v161 = F_psprintf(m, int32(656068), v20+int32(80))
				mBase = m.M
				v162 = m.ExcPending
				if v162 != 0 {
					return int32(0)
				} else {
					v323 = v161
					m.G0 = v20 + int32(416)
					return v323
				}
			} else {
				*(*int32)(unsafe.Add(mBase, uint32(v20)+72)) = v84
				*(*int32)(unsafe.Add(mBase, uint32(v20)+68)) = v154
				*(*int32)(unsafe.Add(mBase, uint32(v20)+64)) = v132
				v169 = F_psprintf(m, int32(656068), v20-int32(-64))
				mBase = m.M
				v170 = m.ExcPending
				if v170 != 0 {
					return int32(0)
				} else {
					v323 = v169
					m.G0 = v20 + int32(416)
					return v323
				}
			}
		default:
			v174 = v91 & int32(255)
			if v174 == int32(1) {
				v177 = int32(727670)
			} else {
				v177 = int32(738731)
			}
			if v174 == int32(2) {
				v182 = int32(727670)
			} else {
				v182 = int32(738731)
			}
			if v89 != 0 {
				*(*int32)(unsafe.Add(mBase, uint32(v20)+48)) = v132
				*(*int32)(unsafe.Add(mBase, uint32(v20)+44)) = v177
				*(*int32)(unsafe.Add(mBase, uint32(v20)+40)) = v84
				*(*int32)(unsafe.Add(mBase, uint32(v20)+36)) = v182
				*(*int32)(unsafe.Add(mBase, uint32(v20)+32)) = v80
				v191 = F_psprintf(m, int32(174692), v20+int32(32))
				mBase = m.M
				v192 = m.ExcPending
				if v192 != 0 {
					return int32(0)
				} else {
					v323 = v191
					m.G0 = v20 + int32(416)
					return v323
				}
			} else {
				*(*int32)(unsafe.Add(mBase, uint32(v20)+16)) = v84
				*(*int32)(unsafe.Add(mBase, uint32(v20)+12)) = v177
				*(*int32)(unsafe.Add(mBase, uint32(v20)+8)) = v132
				*(*int32)(unsafe.Add(mBase, uint32(v20)+4)) = v182
				*(*int32)(unsafe.Add(mBase, uint32(v20))) = v80
				v199 = F_psprintf(m, int32(174692), v20)
				mBase = m.M
				v200 = m.ExcPending
				if v200 != 0 {
					return int32(0)
				} else {
					v323 = v199
					m.G0 = v20 + int32(416)
					return v323
				}
			}
		case 2:
			v204 = v91 & int32(255)
			if v204 == int32(2) {
				v207 = int32(727670)
			} else {
				v207 = int32(738731)
			}
			if v204 == int32(1) {
				v212 = int32(727670)
			} else {
				v212 = int32(738731)
			}
			if v89 != 0 {
				*(*int32)(unsafe.Add(mBase, uint32(v20)+208)) = v80
				*(*int32)(unsafe.Add(mBase, uint32(v20)+204)) = v207
				*(*int32)(unsafe.Add(mBase, uint32(v20)+200)) = v132
				*(*int32)(unsafe.Add(mBase, uint32(v20)+196)) = v212
				*(*int32)(unsafe.Add(mBase, uint32(v20)+192)) = v84
				v221 = F_psprintf(m, int32(174692), v20+int32(192))
				mBase = m.M
				v222 = m.ExcPending
				if v222 != 0 {
					return int32(0)
				} else {
					v323 = v221
					m.G0 = v20 + int32(416)
					return v323
				}
			} else {
				*(*int32)(unsafe.Add(mBase, uint32(v20)+176)) = v80
				*(*int32)(unsafe.Add(mBase, uint32(v20)+172)) = v207
				*(*int32)(unsafe.Add(mBase, uint32(v20)+168)) = v84
				*(*int32)(unsafe.Add(mBase, uint32(v20)+164)) = v212
				*(*int32)(unsafe.Add(mBase, uint32(v20)+160)) = v132
				v231 = F_psprintf(m, int32(174692), v20+int32(160))
				mBase = m.M
				v232 = m.ExcPending
				if v232 != 0 {
					return int32(0)
				} else {
					v323 = v231
					m.G0 = v20 + int32(416)
					return v323
				}
			}
		case 3:
			if v89 != 0 {
				*(*int32)(unsafe.Add(mBase, uint32(v20)+272)) = v132
				*(*int32)(unsafe.Add(mBase, uint32(v20)+264)) = v84
				*(*int32)(unsafe.Add(mBase, uint32(v20)+256)) = v80
				v239 = v91 & int32(255)
				if v239 == int32(1) {
					v242 = int32(727670)
				} else {
					v242 = int32(738731)
				}
				*(*int32)(unsafe.Add(mBase, uint32(v20)+268)) = v242
				if v239 == int32(2) {
					v248 = int32(727670)
				} else {
					v248 = int32(738731)
				}
				*(*int32)(unsafe.Add(mBase, uint32(v20)+260)) = v248
				v253 = F_psprintf(m, int32(174692), v20+int32(256))
				mBase = m.M
				v254 = m.ExcPending
				if v254 != 0 {
					return int32(0)
				} else {
					v323 = v253
					m.G0 = v20 + int32(416)
					return v323
				}
			} else {
				*(*int32)(unsafe.Add(mBase, uint32(v20)+240)) = v84
				*(*int32)(unsafe.Add(mBase, uint32(v20)+232)) = v80
				*(*int32)(unsafe.Add(mBase, uint32(v20)+224)) = v132
				v261 = v91 & int32(255)
				if v261 == int32(2) {
					v264 = int32(727670)
				} else {
					v264 = int32(738731)
				}
				*(*int32)(unsafe.Add(mBase, uint32(v20)+236)) = v264
				if v261 == int32(1) {
					v270 = int32(727670)
				} else {
					v270 = int32(738731)
				}
				*(*int32)(unsafe.Add(mBase, uint32(v20)+228)) = v270
				v275 = F_psprintf(m, int32(174692), v20+int32(224))
				mBase = m.M
				v276 = m.ExcPending
				if v276 != 0 {
					return int32(0)
				} else {
					v323 = v275
					m.G0 = v20 + int32(416)
					return v323
				}
			}
		case 4:
			if v89 != 0 {
				*(*int32)(unsafe.Add(mBase, uint32(v20)+144)) = v132
				*(*int32)(unsafe.Add(mBase, uint32(v20)+136)) = v80
				*(*int32)(unsafe.Add(mBase, uint32(v20)+128)) = v84
				v283 = v91 & int32(255)
				if v283 == int32(1) {
					v286 = int32(727670)
				} else {
					v286 = int32(738731)
				}
				*(*int32)(unsafe.Add(mBase, uint32(v20)+140)) = v286
				if v283 == int32(2) {
					v292 = int32(727670)
				} else {
					v292 = int32(738731)
				}
				*(*int32)(unsafe.Add(mBase, uint32(v20)+132)) = v292
				v297 = F_psprintf(m, int32(174692), v20+int32(128))
				mBase = m.M
				v298 = m.ExcPending
				if v298 != 0 {
					return int32(0)
				} else {
					v323 = v297
					m.G0 = v20 + int32(416)
					return v323
				}
			} else {
				*(*int32)(unsafe.Add(mBase, uint32(v20)+112)) = v80
				*(*int32)(unsafe.Add(mBase, uint32(v20)+104)) = v84
				*(*int32)(unsafe.Add(mBase, uint32(v20)+96)) = v132
				v305 = v91 & int32(255)
				if v305 == int32(2) {
					v308 = int32(727670)
				} else {
					v308 = int32(738731)
				}
				*(*int32)(unsafe.Add(mBase, uint32(v20)+108)) = v308
				if v305 == int32(1) {
					v314 = int32(727670)
				} else {
					v314 = int32(738731)
				}
				*(*int32)(unsafe.Add(mBase, uint32(v20)+100)) = v314
				v319 = F_psprintf(m, int32(174692), v20+int32(96))
				mBase = m.M
				v320 = m.ExcPending
				if v320 != 0 {
					return int32(0)
				} else {
					v323 = v319
					m.G0 = v20 + int32(416)
					return v323
				}
			}
		}
	}
}
func F_cash_pl(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v4 int32
	_ = v4
	var v5 int64
	_ = v5
	var v8 int32
	_ = v8
	var v9 int64
	_ = v9
	var v10 int64
	_ = v10
	var v18 int32
	_ = v18
	var v21 int32
	_ = v21
	var v25 int32
	_ = v25
	var v30 int32
	_ = v30
	var v31 int32
	_ = v31
	var v32 int32
	_ = v32
	v4 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
	v5 = *(*int64)(unsafe.Add(mBase, uint32(v4)))
	v8 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v9 = *(*int64)(unsafe.Add(mBase, uint32(v8)))
	v10 = v9 + v5
	if base.B2i32(v5 < int64(0)) != base.B2i32(v10 < v9) {
		F_errstart_cold(m, int32(21), int32(0))
		mBase = m.M
		v18 = m.ExcPending
		if v18 != 0 {
			return int32(0)
		} else {
			F_errcode(m, int32(50331778))
			mBase = m.M
			v21 = m.ExcPending
			if v21 != 0 {
				return int32(0)
			} else {
				F_errmsg(m, int32(399601), int32(0))
				mBase = m.M
				v25 = m.ExcPending
				if v25 != 0 {
					return int32(0)
				} else {
					F_errfinish(m, int32(494794), int32(98), int32(321693))
					mBase = m.M
					v30 = m.ExcPending
					if v30 != 0 {
						return int32(0)
					} else {
						base.Wasm_trap_unreachable()
						for {
						}
					}
				}
			}
		}
	} else {
		v31 = F_Int64GetDatum(m, v10)
		mBase = m.M
		v32 = m.ExcPending
		if v32 != 0 {
			return int32(0)
		} else {
			return v31
		}
	}
}
