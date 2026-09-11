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
								F_errmsg(m, int32(377941), int32(0))
								v62 = m.ExcPending
								if v62 != 0 {
									return int64(0)
								} else {
									F_errfinish(m, int32(467308), int32(137), int32(516100))
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
									F_errmsg(m, int32(377941), int32(0))
									v62 = m.ExcPending
									if v62 != 0 {
										return int64(0)
									} else {
										F_errfinish(m, int32(467308), int32(137), int32(516100))
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
										F_errmsg(m, int32(377941), int32(0))
										v62 = m.ExcPending
										if v62 != 0 {
											return int64(0)
										} else {
											F_errfinish(m, int32(467308), int32(137), int32(516100))
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
									F_errmsg(m, int32(377941), int32(0))
									v62 = m.ExcPending
									if v62 != 0 {
										return int64(0)
									} else {
										F_errfinish(m, int32(467308), int32(137), int32(516100))
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
										F_errmsg(m, int32(377941), int32(0))
										v62 = m.ExcPending
										if v62 != 0 {
											return int64(0)
										} else {
											F_errfinish(m, int32(467308), int32(137), int32(516100))
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
											F_errmsg(m, int32(377941), int32(0))
											v62 = m.ExcPending
											if v62 != 0 {
												return int64(0)
											} else {
												F_errfinish(m, int32(467308), int32(137), int32(516100))
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
										F_errmsg(m, int32(377941), int32(0))
										v62 = m.ExcPending
										if v62 != 0 {
											return int64(0)
										} else {
											F_errfinish(m, int32(467308), int32(137), int32(516100))
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
											F_errmsg(m, int32(377941), int32(0))
											v62 = m.ExcPending
											if v62 != 0 {
												return int64(0)
											} else {
												F_errfinish(m, int32(467308), int32(137), int32(516100))
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
												F_errmsg(m, int32(377941), int32(0))
												v62 = m.ExcPending
												if v62 != 0 {
													return int64(0)
												} else {
													F_errfinish(m, int32(467308), int32(137), int32(516100))
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
							F_errmsg(m, int32(377941), int32(0))
							v62 = m.ExcPending
							if v62 != 0 {
								return int64(0)
							} else {
								F_errfinish(m, int32(467308), int32(137), int32(516100))
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
								F_errmsg(m, int32(377941), int32(0))
								v62 = m.ExcPending
								if v62 != 0 {
									return int64(0)
								} else {
									F_errfinish(m, int32(467308), int32(137), int32(516100))
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
									F_errmsg(m, int32(377941), int32(0))
									v62 = m.ExcPending
									if v62 != 0 {
										return int64(0)
									} else {
										F_errfinish(m, int32(467308), int32(137), int32(516100))
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
								F_errmsg(m, int32(377941), int32(0))
								v62 = m.ExcPending
								if v62 != 0 {
									return int64(0)
								} else {
									F_errfinish(m, int32(467308), int32(137), int32(516100))
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
									F_errmsg(m, int32(377941), int32(0))
									v62 = m.ExcPending
									if v62 != 0 {
										return int64(0)
									} else {
										F_errfinish(m, int32(467308), int32(137), int32(516100))
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
										F_errmsg(m, int32(377941), int32(0))
										v62 = m.ExcPending
										if v62 != 0 {
											return int64(0)
										} else {
											F_errfinish(m, int32(467308), int32(137), int32(516100))
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
									F_errmsg(m, int32(377941), int32(0))
									v62 = m.ExcPending
									if v62 != 0 {
										return int64(0)
									} else {
										F_errfinish(m, int32(467308), int32(137), int32(516100))
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
										F_errmsg(m, int32(377941), int32(0))
										v62 = m.ExcPending
										if v62 != 0 {
											return int64(0)
										} else {
											F_errfinish(m, int32(467308), int32(137), int32(516100))
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
											F_errmsg(m, int32(377941), int32(0))
											v62 = m.ExcPending
											if v62 != 0 {
												return int64(0)
											} else {
												F_errfinish(m, int32(467308), int32(137), int32(516100))
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
				F_errmsg(m, int32(377941), int32(0))
				mBase = m.M
				v25 = m.ExcPending
				if v25 != 0 {
					return int32(0)
				} else {
					F_errfinish(m, int32(467308), int32(111), int32(303866))
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
				F_errmsg(m, int32(377941), int32(0))
				mBase = m.M
				v68 = m.ExcPending
				if v68 != 0 {
					return int32(0)
				} else {
					F_errfinish(m, int32(467308), int32(150), int32(519000))
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
				F_errmsg(m, int32(377941), int32(0))
				mBase = m.M
				v69 = m.ExcPending
				if v69 != 0 {
					return int32(0)
				} else {
					F_errfinish(m, int32(467308), int32(150), int32(519000))
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
	var v132 int32
	_ = v132
	var v137 int32
	_ = v137
	var v141 int32
	_ = v141
	var v146 int32
	_ = v146
	var v148 int32
	_ = v148
	var v152 int32
	_ = v152
	var v158 int32
	_ = v158
	var v161 int32
	_ = v161
	var v167 int32
	_ = v167
	var v171 int32
	_ = v171
	var v173 int32
	_ = v173
	var v181 int32
	_ = v181
	var v182 int32
	_ = v182
	var v183 int32
	_ = v183
	var v185 int32
	_ = v185
	var v187 int32
	_ = v187
	var v188 int32
	_ = v188
	var v189 int64
	_ = v189
	var v190 int64
	_ = v190
	var v196 int32
	_ = v196
	var v199 int32
	_ = v199
	var v210 int32
	_ = v210
	var v217 int32
	_ = v217
	var v218 int32
	_ = v218
	var v225 int32
	_ = v225
	var v226 int32
	_ = v226
	var v230 int32
	_ = v230
	var v233 int32
	_ = v233
	var v238 int32
	_ = v238
	var v247 int32
	_ = v247
	var v248 int32
	_ = v248
	var v255 int32
	_ = v255
	var v256 int32
	_ = v256
	var v260 int32
	_ = v260
	var v263 int32
	_ = v263
	var v268 int32
	_ = v268
	var v277 int32
	_ = v277
	var v278 int32
	_ = v278
	var v287 int32
	_ = v287
	var v288 int32
	_ = v288
	var v295 int32
	_ = v295
	var v298 int32
	_ = v298
	var v304 int32
	_ = v304
	var v309 int32
	_ = v309
	var v310 int32
	_ = v310
	var v317 int32
	_ = v317
	var v320 int32
	_ = v320
	var v326 int32
	_ = v326
	var v331 int32
	_ = v331
	var v332 int32
	_ = v332
	var v339 int32
	_ = v339
	var v342 int32
	_ = v342
	var v348 int32
	_ = v348
	var v353 int32
	_ = v353
	var v354 int32
	_ = v354
	var v361 int32
	_ = v361
	var v364 int32
	_ = v364
	var v370 int32
	_ = v370
	var v375 int32
	_ = v375
	var v376 int32
	_ = v376
	var v379 int32
	_ = v379
	v18 = m.G0
	v20 = v18 - int32(416)
	m.G0 = v20
	v22 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v23 = *(*int64)(unsafe.Add(mBase, uint32(v22)))
	v24 = F_PGLC_localeconv(m)
	mBase = m.M
	v27 = m.ExcPending
	if v27 != 0 {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	return int32(0)
L2:
	;
	v28 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v24)+41)))
	v31 = *(*int32)(unsafe.Add(mBase, uint32(v24)+28))
	v32 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v31))))
	v39 = int32(620157)
	v40 = int32(46)
	v41 = *(*int32)(unsafe.Add(mBase, uint32(v24)+20))
	v42 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v41))))
	if v42 == int32(0) {
		v53 = v39
		v54 = v40
		goto L3
	} else {
		goto L4
	}
L3:
	;
	if base.Ui32(int32(10)) < base.Ui32(v28) {
		goto L9
	} else {
		goto L10
	}
L4:
	;
	v45 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v41)+1)))
	if v45 != 0 {
		v53 = v39
		v54 = v40
		goto L3
	} else {
		goto L5
	}
L5:
	;
	if v42&int32(255) == int32(44) {
		goto L6
	} else {
		goto L7
	}
L6:
	;
	v52 = int32(620121)
	goto L8
L7:
	;
	v52 = int32(620157)
	goto L8
L8:
	;
	v53 = v52
	v54 = v42
	goto L3
L9:
	;
	v56 = int32(2)
	goto L11
L10:
	;
	v56 = v28
	goto L11
L11:
	;
	if base.Ui32((v32-int32(7))&int32(255)) < base.Ui32(int32(250)) {
		goto L12
	} else {
		goto L13
	}
L12:
	;
	v58 = int32(3)
	goto L14
L13:
	;
	v58 = v32
	goto L14
L14:
	;
	v59 = *(*int32)(unsafe.Add(mBase, uint32(v24)+16))
	v60 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v59))))
	v61 = *(*int32)(unsafe.Add(mBase, uint32(v24)+24))
	v62 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v61))))
	if v23 < int64(0) {
		goto L16
	} else {
		goto L17
	}
L15:
	;
	if v60 != 0 {
		goto L22
	} else {
		goto L23
	}
L16:
	;
	v68 = *(*int32)(unsafe.Add(mBase, uint32(v24)+36))
	v70 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v68))))
	if v70 != 0 {
		goto L19
	} else {
		goto L20
	}
L17:
	;
	goto L18
L18:
	;
	v75 = *(*int32)(unsafe.Add(mBase, uint32(v24)+32))
	v77 = int32(42)
	v78 = int32(43)
	v79 = int32(46)
	v80 = v75
	goto L15
L19:
	;
	v71 = v68
	goto L21
L20:
	;
	v71 = int32(620147)
	goto L21
L21:
	;
	v77 = int32(44)
	v78 = int32(45)
	v79 = int32(47)
	v80 = v71
	goto L15
L22:
	;
	v84 = v59
	goto L24
L23:
	;
	v84 = int32(639591)
	goto L24
L24:
	;
	if v62 != 0 {
		goto L25
	} else {
		goto L26
	}
L25:
	;
	v85 = v61
	goto L27
L26:
	;
	v85 = v53
	goto L27
L27:
	;
	v87 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v24+v79))))
	v89 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v24+v77))))
	v91 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v24+v78))))
	v92 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v20)+415)) = uint8(v92)
	v95 = v23 >> (uint(int64(63)) % 64)
	v100 = base.I32_extend8_s(v56)
	v102 = v20 + int32(415)
	v116 = v23 ^ v95 - v95
	goto L28
L28:
	;
	if v56 == int32(0) {
		goto L31
	} else {
		goto L32
	}
L29:
	;
	switch v87 {
	case 0:
		goto L64
	default:
		goto L63
	case 2:
		goto L62
	case 3:
		goto L61
	case 4:
		goto L60
	}
L30:
	;
	v187 = int32(1)
	v188 = v185 - v187
	v189 = int64(10)
	v190 = base.I64_div_u_s(v116, v189)
	v196 = base.I32_wrap_i64(v116-v190*v189) | int32(48)
	*(*uint8)(unsafe.Add(mBase, uint32(v188))) = uint8(v196)
	v199 = v100 - v187
	if base.Ui64(int64(9)) < base.Ui64(v116) {
		v100 = v199
		v102 = v188
		v116 = v190
		goto L28
	} else {
		goto L57
	}
L31:
	;
	if int32(0) <= v100 {
		v185 = v102
		goto L30
	} else {
		goto L34
	}
L32:
	;
	if v100 != 0 {
		goto L31
	} else {
		goto L33
	}
L33:
	;
	v120 = v102 - int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(v120))) = uint8(v54)
	v185 = v120
	goto L30
L34:
	;
	v124 = base.I32_rem_s(v100, base.I32_extend8_s(v58))
	if v124 != 0 {
		v185 = v102
		goto L30
	} else {
		goto L35
	}
L35:
	;
	if v85&int32(3) == int32(0) {
		v148 = v85
		goto L38
	} else {
		goto L39
	}
L36:
	;
	v182 = v102 - v181
	if v181 != 0 {
		goto L54
	} else {
		goto L55
	}
L37:
	;
	v181 = v173 - v85
	goto L36
L38:
	;
	v152 = v148
	goto L47
L39:
	;
	v132 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v85))))
	if v132 == int32(0) {
		goto L40
	} else {
		goto L41
	}
L40:
	;
	v181 = int32(0)
	goto L36
L41:
	;
	goto L42
L42:
	;
	v137 = v85
	goto L43
L43:
	;
	v141 = v137 + int32(1)
	if v141&int32(3) == int32(0) {
		v148 = v141
		goto L38
	} else {
		goto L45
	}
L44:
	;
	v173 = v141
	goto L37
L45:
	;
	v146 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v141))))
	if v146 != 0 {
		v137 = v141
		goto L43
	} else {
		goto L46
	}
L46:
	;
	goto L44
L47:
	;
	v158 = *(*int32)(unsafe.Add(mBase, uint32(v152)))
	v161 = int32(-2139062144)
	if (int32(16843008)-v158|v158)&v161 == v161 {
		v152 = v152 + int32(4)
		goto L47
	} else {
		goto L49
	}
L48:
	;
	v167 = v152
	goto L50
L49:
	;
	goto L48
L50:
	;
	v171 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v167))))
	if v171 != 0 {
		v167 = v167 + int32(1)
		goto L50
	} else {
		goto L52
	}
L51:
	;
	v173 = v167
	goto L37
L52:
	;
	goto L51
L53:
	;
	v185 = v182
	goto L30
L54:
	;
	v183 = F__emscripten_memcpy_bulkmem(m, v182, v85, v181)
	mBase = m.M
	goto L56
L55:
	;
	goto L56
L56:
	;
	goto L53
L57:
	;
	if int32(0) <= v199 {
		v100 = v199
		v102 = v188
		v116 = v190
		goto L28
	} else {
		goto L58
	}
L58:
	;
	goto L29
L59:
	;
	m.G0 = v20 + int32(416)
	return v379
L60:
	;
	if v89 != 0 {
		goto L112
	} else {
		goto L113
	}
L61:
	;
	if v89 != 0 {
		goto L95
	} else {
		goto L96
	}
L62:
	;
	v260 = v91 & int32(255)
	if v260 == int32(2) {
		goto L84
	} else {
		goto L85
	}
L63:
	;
	v230 = v91 & int32(255)
	if v230 == int32(1) {
		goto L73
	} else {
		goto L74
	}
L64:
	;
	if v91&int32(255) == int32(1) {
		goto L65
	} else {
		goto L66
	}
L65:
	;
	v210 = int32(695486)
	goto L67
L66:
	;
	v210 = int32(706478)
	goto L67
L67:
	;
	if v89 != 0 {
		goto L68
	} else {
		goto L69
	}
L68:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v20)+88)) = v188
	*(*int32)(unsafe.Add(mBase, uint32(v20)+84)) = v210
	*(*int32)(unsafe.Add(mBase, uint32(v20)+80)) = v84
	v217 = F_psprintf(m, int32(624706), v20+int32(80))
	mBase = m.M
	v218 = m.ExcPending
	if v218 != 0 {
		goto L1
	} else {
		goto L71
	}
L69:
	;
	goto L70
L70:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v20)+72)) = v84
	*(*int32)(unsafe.Add(mBase, uint32(v20)+68)) = v210
	*(*int32)(unsafe.Add(mBase, uint32(v20)+64)) = v188
	v225 = F_psprintf(m, int32(624706), v20-int32(-64))
	mBase = m.M
	v226 = m.ExcPending
	if v226 != 0 {
		goto L1
	} else {
		goto L72
	}
L71:
	;
	v379 = v217
	goto L59
L72:
	;
	v379 = v225
	goto L59
L73:
	;
	v233 = int32(695486)
	goto L75
L74:
	;
	v233 = int32(706478)
	goto L75
L75:
	;
	if v230 == int32(2) {
		goto L76
	} else {
		goto L77
	}
L76:
	;
	v238 = int32(695486)
	goto L78
L77:
	;
	v238 = int32(706478)
	goto L78
L78:
	;
	if v89 != 0 {
		goto L79
	} else {
		goto L80
	}
L79:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v20)+48)) = v188
	*(*int32)(unsafe.Add(mBase, uint32(v20)+44)) = v233
	*(*int32)(unsafe.Add(mBase, uint32(v20)+40)) = v84
	*(*int32)(unsafe.Add(mBase, uint32(v20)+36)) = v238
	*(*int32)(unsafe.Add(mBase, uint32(v20)+32)) = v80
	v247 = F_psprintf(m, int32(163685), v20+int32(32))
	mBase = m.M
	v248 = m.ExcPending
	if v248 != 0 {
		goto L1
	} else {
		goto L82
	}
L80:
	;
	goto L81
L81:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v20)+16)) = v84
	*(*int32)(unsafe.Add(mBase, uint32(v20)+12)) = v233
	*(*int32)(unsafe.Add(mBase, uint32(v20)+8)) = v188
	*(*int32)(unsafe.Add(mBase, uint32(v20)+4)) = v238
	*(*int32)(unsafe.Add(mBase, uint32(v20))) = v80
	v255 = F_psprintf(m, int32(163685), v20)
	mBase = m.M
	v256 = m.ExcPending
	if v256 != 0 {
		goto L1
	} else {
		goto L83
	}
L82:
	;
	v379 = v247
	goto L59
L83:
	;
	v379 = v255
	goto L59
L84:
	;
	v263 = int32(695486)
	goto L86
L85:
	;
	v263 = int32(706478)
	goto L86
L86:
	;
	if v260 == int32(1) {
		goto L87
	} else {
		goto L88
	}
L87:
	;
	v268 = int32(695486)
	goto L89
L88:
	;
	v268 = int32(706478)
	goto L89
L89:
	;
	if v89 != 0 {
		goto L90
	} else {
		goto L91
	}
L90:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v20)+208)) = v80
	*(*int32)(unsafe.Add(mBase, uint32(v20)+204)) = v263
	*(*int32)(unsafe.Add(mBase, uint32(v20)+200)) = v188
	*(*int32)(unsafe.Add(mBase, uint32(v20)+196)) = v268
	*(*int32)(unsafe.Add(mBase, uint32(v20)+192)) = v84
	v277 = F_psprintf(m, int32(163685), v20+int32(192))
	mBase = m.M
	v278 = m.ExcPending
	if v278 != 0 {
		goto L1
	} else {
		goto L93
	}
L91:
	;
	goto L92
L92:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v20)+176)) = v80
	*(*int32)(unsafe.Add(mBase, uint32(v20)+172)) = v263
	*(*int32)(unsafe.Add(mBase, uint32(v20)+168)) = v84
	*(*int32)(unsafe.Add(mBase, uint32(v20)+164)) = v268
	*(*int32)(unsafe.Add(mBase, uint32(v20)+160)) = v188
	v287 = F_psprintf(m, int32(163685), v20+int32(160))
	mBase = m.M
	v288 = m.ExcPending
	if v288 != 0 {
		goto L1
	} else {
		goto L94
	}
L93:
	;
	v379 = v277
	goto L59
L94:
	;
	v379 = v287
	goto L59
L95:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v20)+272)) = v188
	*(*int32)(unsafe.Add(mBase, uint32(v20)+264)) = v84
	*(*int32)(unsafe.Add(mBase, uint32(v20)+256)) = v80
	v295 = v91 & int32(255)
	if v295 == int32(1) {
		goto L98
	} else {
		goto L99
	}
L96:
	;
	goto L97
L97:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v20)+240)) = v84
	*(*int32)(unsafe.Add(mBase, uint32(v20)+232)) = v80
	*(*int32)(unsafe.Add(mBase, uint32(v20)+224)) = v188
	v317 = v91 & int32(255)
	if v317 == int32(2) {
		goto L105
	} else {
		goto L106
	}
L98:
	;
	v298 = int32(695486)
	goto L100
L99:
	;
	v298 = int32(706478)
	goto L100
L100:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v20)+268)) = v298
	if v295 == int32(2) {
		goto L101
	} else {
		goto L102
	}
L101:
	;
	v304 = int32(695486)
	goto L103
L102:
	;
	v304 = int32(706478)
	goto L103
L103:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v20)+260)) = v304
	v309 = F_psprintf(m, int32(163685), v20+int32(256))
	mBase = m.M
	v310 = m.ExcPending
	if v310 != 0 {
		goto L1
	} else {
		goto L104
	}
L104:
	;
	v379 = v309
	goto L59
L105:
	;
	v320 = int32(695486)
	goto L107
L106:
	;
	v320 = int32(706478)
	goto L107
L107:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v20)+236)) = v320
	if v317 == int32(1) {
		goto L108
	} else {
		goto L109
	}
L108:
	;
	v326 = int32(695486)
	goto L110
L109:
	;
	v326 = int32(706478)
	goto L110
L110:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v20)+228)) = v326
	v331 = F_psprintf(m, int32(163685), v20+int32(224))
	mBase = m.M
	v332 = m.ExcPending
	if v332 != 0 {
		goto L1
	} else {
		goto L111
	}
L111:
	;
	v379 = v331
	goto L59
L112:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v20)+144)) = v188
	*(*int32)(unsafe.Add(mBase, uint32(v20)+136)) = v80
	*(*int32)(unsafe.Add(mBase, uint32(v20)+128)) = v84
	v339 = v91 & int32(255)
	if v339 == int32(1) {
		goto L115
	} else {
		goto L116
	}
L113:
	;
	goto L114
L114:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v20)+112)) = v80
	*(*int32)(unsafe.Add(mBase, uint32(v20)+104)) = v84
	*(*int32)(unsafe.Add(mBase, uint32(v20)+96)) = v188
	v361 = v91 & int32(255)
	if v361 == int32(2) {
		goto L122
	} else {
		goto L123
	}
L115:
	;
	v342 = int32(695486)
	goto L117
L116:
	;
	v342 = int32(706478)
	goto L117
L117:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v20)+140)) = v342
	if v339 == int32(2) {
		goto L118
	} else {
		goto L119
	}
L118:
	;
	v348 = int32(695486)
	goto L120
L119:
	;
	v348 = int32(706478)
	goto L120
L120:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v20)+132)) = v348
	v353 = F_psprintf(m, int32(163685), v20+int32(128))
	mBase = m.M
	v354 = m.ExcPending
	if v354 != 0 {
		goto L1
	} else {
		goto L121
	}
L121:
	;
	v379 = v353
	goto L59
L122:
	;
	v364 = int32(695486)
	goto L124
L123:
	;
	v364 = int32(706478)
	goto L124
L124:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v20)+108)) = v364
	if v361 == int32(1) {
		goto L125
	} else {
		goto L126
	}
L125:
	;
	v370 = int32(695486)
	goto L127
L126:
	;
	v370 = int32(706478)
	goto L127
L127:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v20)+100)) = v370
	v375 = F_psprintf(m, int32(163685), v20+int32(96))
	mBase = m.M
	v376 = m.ExcPending
	if v376 != 0 {
		goto L1
	} else {
		goto L128
	}
L128:
	;
	v379 = v375
	goto L59
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
				F_errmsg(m, int32(377941), int32(0))
				mBase = m.M
				v25 = m.ExcPending
				if v25 != 0 {
					return int32(0)
				} else {
					F_errfinish(m, int32(467308), int32(98), int32(303853))
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
