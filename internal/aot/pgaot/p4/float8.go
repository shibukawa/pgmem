package p4

import (
	base "github.com/shibukawa/pgmem/internal/aot/pgaot/base"
	"math"
	"unsafe"
)

func F_float8_accum(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v10 int32
	_ = v10
	var v12 int32
	_ = v12
	var v14 int32
	_ = v14
	var v15 int32
	_ = v15
	var v18 int32
	_ = v18
	var v19 int32
	_ = v19
	var v22 int32
	_ = v22
	var v25 int32
	_ = v25
	var v26 int32
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
	var v36 float64
	_ = v36
	var v38 float64
	_ = v38
	var v43 float64
	_ = v43
	var v47 float64
	_ = v47
	var v50 float64
	_ = v50
	var v63 int32
	_ = v63
	var v64 float64
	_ = v64
	var v75 float64
	_ = v75
	var v79 int32
	_ = v79
	var v82 int32
	_ = v82
	var v107 int32
	_ = v107
	var v124 int32
	_ = v124
	var v125 int32
	_ = v125
	var v126 int32
	_ = v126
	var v134 int32
	_ = v134
	var v141 int32
	_ = v141
	var v146 int32
	_ = v146
	v10 = m.G0
	v12 = v10 - int32(48)
	m.G0 = v12
	v14 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v15 = F_pg_detoast_datum(m, v14)
	mBase = m.M
	v18 = m.ExcPending
	if v18 != 0 {
		return int32(0)
	} else {
		v19 = *(*int32)(unsafe.Add(mBase, uint32(v15)+4))
		if v19 != int32(1) {
			F_errstart_cold(m, int32(21), int32(0))
			mBase = m.M
			v134 = m.ExcPending
			if v134 != 0 {
				return int32(0)
			} else {
				*(*int32)(unsafe.Add(mBase, uint32(v12)+4)) = int32(3)
				*(*int32)(unsafe.Add(mBase, uint32(v12))) = int32(285780)
				F_errmsg_internal(m, int32(26093), v12)
				mBase = m.M
				v141 = m.ExcPending
				if v141 != 0 {
					return int32(0)
				} else {
					F_errfinish(m, int32(490414), int32(2938), int32(24633))
					mBase = m.M
					v146 = m.ExcPending
					if v146 != 0 {
						return int32(0)
					} else {
						base.Wasm_trap_unreachable()
						for {
						}
					}
				}
			}
		} else {
			v22 = *(*int32)(unsafe.Add(mBase, uint32(v15)+16))
			if v22 != int32(3) {
				F_errstart_cold(m, int32(21), int32(0))
				mBase = m.M
				v134 = m.ExcPending
				if v134 != 0 {
					return int32(0)
				} else {
					*(*int32)(unsafe.Add(mBase, uint32(v12)+4)) = int32(3)
					*(*int32)(unsafe.Add(mBase, uint32(v12))) = int32(285780)
					F_errmsg_internal(m, int32(26093), v12)
					mBase = m.M
					v141 = m.ExcPending
					if v141 != 0 {
						return int32(0)
					} else {
						F_errfinish(m, int32(490414), int32(2938), int32(24633))
						mBase = m.M
						v146 = m.ExcPending
						if v146 != 0 {
							return int32(0)
						} else {
							base.Wasm_trap_unreachable()
							for {
							}
						}
					}
				}
			} else {
				v25 = *(*int32)(unsafe.Add(mBase, uint32(v15)+8))
				if v25 != 0 {
					F_errstart_cold(m, int32(21), int32(0))
					mBase = m.M
					v134 = m.ExcPending
					if v134 != 0 {
						return int32(0)
					} else {
						*(*int32)(unsafe.Add(mBase, uint32(v12)+4)) = int32(3)
						*(*int32)(unsafe.Add(mBase, uint32(v12))) = int32(285780)
						F_errmsg_internal(m, int32(26093), v12)
						mBase = m.M
						v141 = m.ExcPending
						if v141 != 0 {
							return int32(0)
						} else {
							F_errfinish(m, int32(490414), int32(2938), int32(24633))
							mBase = m.M
							v146 = m.ExcPending
							if v146 != 0 {
								return int32(0)
							} else {
								base.Wasm_trap_unreachable()
								for {
								}
							}
						}
					}
				} else {
					v26 = *(*int32)(unsafe.Add(mBase, uint32(v15)+12))
					if v26 != int32(701) {
						F_errstart_cold(m, int32(21), int32(0))
						mBase = m.M
						v134 = m.ExcPending
						if v134 != 0 {
							return int32(0)
						} else {
							*(*int32)(unsafe.Add(mBase, uint32(v12)+4)) = int32(3)
							*(*int32)(unsafe.Add(mBase, uint32(v12))) = int32(285780)
							F_errmsg_internal(m, int32(26093), v12)
							mBase = m.M
							v141 = m.ExcPending
							if v141 != 0 {
								return int32(0)
							} else {
								F_errfinish(m, int32(490414), int32(2938), int32(24633))
								mBase = m.M
								v146 = m.ExcPending
								if v146 != 0 {
									return int32(0)
								} else {
									base.Wasm_trap_unreachable()
									for {
									}
								}
							}
						}
					} else {
						v29 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
						v30 = *(*float64)(unsafe.Add(mBase, uint32(v29)))
						v31 = *(*float64)(unsafe.Add(mBase, uint32(v15)+32))
						v32 = *(*float64)(unsafe.Add(mBase, uint32(v15)+24))
						v33 = *(*float64)(unsafe.Add(mBase, uint32(v15)+40))
						*(*float64)(unsafe.Add(mBase, uint32(v12)+24)) = v33
						v36 = base.F64_add(v32, float64(1))
						*(*float64)(unsafe.Add(mBase, uint32(v12)+40)) = v36
						v38 = base.F64_add(v30, v31)
						*(*float64)(unsafe.Add(mBase, uint32(v12)+32)) = v38
						if base.F64_gt(v32, float64(0)) != 0 {
							v43 = base.F64_sub(base.F64_mul(v30, v36), v38)
							v47 = base.F64_add(v33, base.F64_div(base.F64_mul(v43, v43), base.F64_mul(v36, v32)))
							*(*float64)(unsafe.Add(mBase, uint32(v12)+24)) = v47
							v50 = math.Float64frombits(uint64(0x7ff0000000000000))
							if base.F64_ne(base.F64_abs(v38), v50)&base.F64_ne(base.F64_abs(v47), v50) != 0 {
								v75 = v47
								v79 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
								if v79 == int32(0) {
									v107 = int32(0)
								} else {
									v82 = *(*int32)(unsafe.Add(mBase, uint32(v79)))
									switch v82 - int32(429) {
									case 0:
										v107 = int32(1)
									case 1:
										v107 = int32(2)
									default:
										v107 = int32(0)
									}
								}
								if v107 != 0 {
									*(*float64)(unsafe.Add(mBase, uint32(v15)+40)) = v75
									*(*float64)(unsafe.Add(mBase, uint32(v15)+32)) = v38
									*(*float64)(unsafe.Add(mBase, uint32(v15)+24)) = v36
									v126 = v15
									m.G0 = v12 + int32(48)
									return v126
								} else {
									*(*int32)(unsafe.Add(mBase, uint32(v12)+20)) = v12 + int32(24)
									*(*int32)(unsafe.Add(mBase, uint32(v12)+16)) = v12 + int32(32)
									*(*int32)(unsafe.Add(mBase, uint32(v12)+12)) = v12 + int32(40)
									v124 = F_construct_array_builtin(m, v12+int32(12), int32(3), int32(701))
									mBase = m.M
									v125 = m.ExcPending
									if v125 != 0 {
										return int32(0)
									} else {
										v126 = v124
										m.G0 = v12 + int32(48)
										return v126
									}
								}
							} else {
								if base.F64_eq(base.F64_abs(v31), math.Float64frombits(uint64(0x7ff0000000000000))) != 0 {
									*(*int64)(unsafe.Add(mBase, uint32(v12)+24)) = int64(9221120237041090560)
									v75 = math.Float64frombits(uint64(0x7ff8000000000000))
									v79 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
									if v79 == int32(0) {
										v107 = int32(0)
									} else {
										v82 = *(*int32)(unsafe.Add(mBase, uint32(v79)))
										switch v82 - int32(429) {
										case 0:
											v107 = int32(1)
										case 1:
											v107 = int32(2)
										default:
											v107 = int32(0)
										}
									}
									if v107 != 0 {
										*(*float64)(unsafe.Add(mBase, uint32(v15)+40)) = v75
										*(*float64)(unsafe.Add(mBase, uint32(v15)+32)) = v38
										*(*float64)(unsafe.Add(mBase, uint32(v15)+24)) = v36
										v126 = v15
										m.G0 = v12 + int32(48)
										return v126
									} else {
										*(*int32)(unsafe.Add(mBase, uint32(v12)+20)) = v12 + int32(24)
										*(*int32)(unsafe.Add(mBase, uint32(v12)+16)) = v12 + int32(32)
										*(*int32)(unsafe.Add(mBase, uint32(v12)+12)) = v12 + int32(40)
										v124 = F_construct_array_builtin(m, v12+int32(12), int32(3), int32(701))
										mBase = m.M
										v125 = m.ExcPending
										if v125 != 0 {
											return int32(0)
										} else {
											v126 = v124
											m.G0 = v12 + int32(48)
											return v126
										}
									}
								} else {
									if base.F64_eq(base.F64_abs(v30), math.Float64frombits(uint64(0x7ff0000000000000))) != 0 {
										*(*int64)(unsafe.Add(mBase, uint32(v12)+24)) = int64(9221120237041090560)
										v75 = math.Float64frombits(uint64(0x7ff8000000000000))
										v79 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
										if v79 == int32(0) {
											v107 = int32(0)
										} else {
											v82 = *(*int32)(unsafe.Add(mBase, uint32(v79)))
											switch v82 - int32(429) {
											case 0:
												v107 = int32(1)
											case 1:
												v107 = int32(2)
											default:
												v107 = int32(0)
											}
										}
										if v107 != 0 {
											*(*float64)(unsafe.Add(mBase, uint32(v15)+40)) = v75
											*(*float64)(unsafe.Add(mBase, uint32(v15)+32)) = v38
											*(*float64)(unsafe.Add(mBase, uint32(v15)+24)) = v36
											v126 = v15
											m.G0 = v12 + int32(48)
											return v126
										} else {
											*(*int32)(unsafe.Add(mBase, uint32(v12)+20)) = v12 + int32(24)
											*(*int32)(unsafe.Add(mBase, uint32(v12)+16)) = v12 + int32(32)
											*(*int32)(unsafe.Add(mBase, uint32(v12)+12)) = v12 + int32(40)
											v124 = F_construct_array_builtin(m, v12+int32(12), int32(3), int32(701))
											mBase = m.M
											v125 = m.ExcPending
											if v125 != 0 {
												return int32(0)
											} else {
												v126 = v124
												m.G0 = v12 + int32(48)
												return v126
											}
										}
									} else {
										F_float_overflow_error(m)
										mBase = m.M
										v63 = m.ExcPending
										if v63 != 0 {
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
							v64 = base.F64_abs(v30)
							if base.Ui64(int64(9218868437227405312)) < base.Ui64(base.I64_reinterpret_f64(v64)) {
								*(*int64)(unsafe.Add(mBase, uint32(v12)+24)) = int64(9221120237041090560)
								v75 = math.Float64frombits(uint64(0x7ff8000000000000))
							} else {
								if base.F64_ne(v64, math.Float64frombits(uint64(0x7ff0000000000000))) != 0 {
									v75 = v33
								} else {
									*(*int64)(unsafe.Add(mBase, uint32(v12)+24)) = int64(9221120237041090560)
									v75 = math.Float64frombits(uint64(0x7ff8000000000000))
								}
							}
							v79 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
							if v79 == int32(0) {
								v107 = int32(0)
							} else {
								v82 = *(*int32)(unsafe.Add(mBase, uint32(v79)))
								switch v82 - int32(429) {
								case 0:
									v107 = int32(1)
								case 1:
									v107 = int32(2)
								default:
									v107 = int32(0)
								}
							}
							if v107 != 0 {
								*(*float64)(unsafe.Add(mBase, uint32(v15)+40)) = v75
								*(*float64)(unsafe.Add(mBase, uint32(v15)+32)) = v38
								*(*float64)(unsafe.Add(mBase, uint32(v15)+24)) = v36
								v126 = v15
								m.G0 = v12 + int32(48)
								return v126
							} else {
								*(*int32)(unsafe.Add(mBase, uint32(v12)+20)) = v12 + int32(24)
								*(*int32)(unsafe.Add(mBase, uint32(v12)+16)) = v12 + int32(32)
								*(*int32)(unsafe.Add(mBase, uint32(v12)+12)) = v12 + int32(40)
								v124 = F_construct_array_builtin(m, v12+int32(12), int32(3), int32(701))
								mBase = m.M
								v125 = m.ExcPending
								if v125 != 0 {
									return int32(0)
								} else {
									v126 = v124
									m.G0 = v12 + int32(48)
									return v126
								}
							}
						}
					}
				}
			}
		}
	}
}
func F_float8_covar_pop(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v5 int32
	_ = v5
	var v7 int32
	_ = v7
	var v9 int32
	_ = v9
	var v10 int32
	_ = v10
	var v13 int32
	_ = v13
	var v14 int32
	_ = v14
	var v17 int32
	_ = v17
	var v20 int32
	_ = v20
	var v21 int32
	_ = v21
	var v24 float64
	_ = v24
	var v27 int32
	_ = v27
	var v32 float64
	_ = v32
	var v34 int32
	_ = v34
	var v35 int32
	_ = v35
	var v36 int32
	_ = v36
	var v44 int32
	_ = v44
	var v51 int32
	_ = v51
	var v56 int32
	_ = v56
	v5 = m.G0
	v7 = v5 - int32(16)
	m.G0 = v7
	v9 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v10 = F_pg_detoast_datum(m, v9)
	mBase = m.M
	v13 = m.ExcPending
	if v13 != 0 {
		return int32(0)
	} else {
		v14 = *(*int32)(unsafe.Add(mBase, uint32(v10)+4))
		if v14 != int32(1) {
			F_errstart_cold(m, int32(21), int32(0))
			mBase = m.M
			v44 = m.ExcPending
			if v44 != 0 {
				return int32(0)
			} else {
				*(*int32)(unsafe.Add(mBase, uint32(v7)+4)) = int32(6)
				*(*int32)(unsafe.Add(mBase, uint32(v7))) = int32(233315)
				F_errmsg_internal(m, int32(26093), v7)
				mBase = m.M
				v51 = m.ExcPending
				if v51 != 0 {
					return int32(0)
				} else {
					F_errfinish(m, int32(490414), int32(2938), int32(24633))
					mBase = m.M
					v56 = m.ExcPending
					if v56 != 0 {
						return int32(0)
					} else {
						base.Wasm_trap_unreachable()
						for {
						}
					}
				}
			}
		} else {
			v17 = *(*int32)(unsafe.Add(mBase, uint32(v10)+16))
			if v17 != int32(6) {
				F_errstart_cold(m, int32(21), int32(0))
				mBase = m.M
				v44 = m.ExcPending
				if v44 != 0 {
					return int32(0)
				} else {
					*(*int32)(unsafe.Add(mBase, uint32(v7)+4)) = int32(6)
					*(*int32)(unsafe.Add(mBase, uint32(v7))) = int32(233315)
					F_errmsg_internal(m, int32(26093), v7)
					mBase = m.M
					v51 = m.ExcPending
					if v51 != 0 {
						return int32(0)
					} else {
						F_errfinish(m, int32(490414), int32(2938), int32(24633))
						mBase = m.M
						v56 = m.ExcPending
						if v56 != 0 {
							return int32(0)
						} else {
							base.Wasm_trap_unreachable()
							for {
							}
						}
					}
				}
			} else {
				v20 = *(*int32)(unsafe.Add(mBase, uint32(v10)+8))
				if v20 != 0 {
					F_errstart_cold(m, int32(21), int32(0))
					mBase = m.M
					v44 = m.ExcPending
					if v44 != 0 {
						return int32(0)
					} else {
						*(*int32)(unsafe.Add(mBase, uint32(v7)+4)) = int32(6)
						*(*int32)(unsafe.Add(mBase, uint32(v7))) = int32(233315)
						F_errmsg_internal(m, int32(26093), v7)
						mBase = m.M
						v51 = m.ExcPending
						if v51 != 0 {
							return int32(0)
						} else {
							F_errfinish(m, int32(490414), int32(2938), int32(24633))
							mBase = m.M
							v56 = m.ExcPending
							if v56 != 0 {
								return int32(0)
							} else {
								base.Wasm_trap_unreachable()
								for {
								}
							}
						}
					}
				} else {
					v21 = *(*int32)(unsafe.Add(mBase, uint32(v10)+12))
					if v21 != int32(701) {
						F_errstart_cold(m, int32(21), int32(0))
						mBase = m.M
						v44 = m.ExcPending
						if v44 != 0 {
							return int32(0)
						} else {
							*(*int32)(unsafe.Add(mBase, uint32(v7)+4)) = int32(6)
							*(*int32)(unsafe.Add(mBase, uint32(v7))) = int32(233315)
							F_errmsg_internal(m, int32(26093), v7)
							mBase = m.M
							v51 = m.ExcPending
							if v51 != 0 {
								return int32(0)
							} else {
								F_errfinish(m, int32(490414), int32(2938), int32(24633))
								mBase = m.M
								v56 = m.ExcPending
								if v56 != 0 {
									return int32(0)
								} else {
									base.Wasm_trap_unreachable()
									for {
									}
								}
							}
						}
					} else {
						v24 = *(*float64)(unsafe.Add(mBase, uint32(v10)+24))
						if base.F64_lt(v24, float64(1)) != 0 {
							v27 = int32(1)
							*(*uint8)(unsafe.Add(mBase, uint32(l0)+16)) = uint8(v27)
							v36 = int32(0)
							m.G0 = v7 + int32(16)
							return v36
						} else {
							v32 = *(*float64)(unsafe.Add(mBase, uint32(v10-int32(-64))))
							v34 = F_Float8GetDatum(m, base.F64_div(v32, v24))
							mBase = m.M
							v35 = m.ExcPending
							if v35 != 0 {
								return int32(0)
							} else {
								v36 = v34
								m.G0 = v7 + int32(16)
								return v36
							}
						}
					}
				}
			}
		}
	}
}
func F_float8_div(m *base.Module, l0 float64, l1 float64) float64 {
	var v11 float64
	_ = v11
	var v13 float64
	_ = v13
	var v30 int32
	_ = v30
	var v32 int32
	_ = v32
	var v34 int32
	_ = v34
	if base.F64_eq(l1, float64(0)) != 0 {
		if base.Ui64(base.I64_reinterpret_f64(l0)&int64(9223372036854775807)) <= base.Ui64(int64(9218868437227405312)) {
			F_float_zero_divide_error(m)
			v30 = m.ExcPending
			if v30 != 0 {
				return float64(0)
			} else {
				base.Wasm_trap_unreachable()
				for {
				}
			}
		} else {
			v11 = base.F64_div(l0, l1)
			v13 = math.Float64frombits(uint64(0x7ff0000000000000))
			if base.F64_eq(base.F64_abs(v11), v13)&base.F64_ne(base.F64_abs(l0), v13) != 0 {
				F_float_overflow_error(m)
				v32 = m.ExcPending
				if v32 != 0 {
					return float64(0)
				} else {
					base.Wasm_trap_unreachable()
					for {
					}
				}
			} else {
				if base.F64_ne(v11, float64(0)) != 0 {
					return v11
				} else {
					if base.F64_eq(l0, float64(0)) != 0 {
						return v11
					} else {
						if base.F64_ne(base.F64_abs(l1), math.Float64frombits(uint64(0x7ff0000000000000))) != 0 {
							F_float_underflow_error(m)
							v34 = m.ExcPending
							if v34 != 0 {
								return float64(0)
							} else {
								base.Wasm_trap_unreachable()
								for {
								}
							}
						} else {
							return v11
						}
					}
				}
			}
		}
	} else {
		v11 = base.F64_div(l0, l1)
		v13 = math.Float64frombits(uint64(0x7ff0000000000000))
		if base.F64_eq(base.F64_abs(v11), v13)&base.F64_ne(base.F64_abs(l0), v13) != 0 {
			F_float_overflow_error(m)
			v32 = m.ExcPending
			if v32 != 0 {
				return float64(0)
			} else {
				base.Wasm_trap_unreachable()
				for {
				}
			}
		} else {
			if base.F64_ne(v11, float64(0)) != 0 {
				return v11
			} else {
				if base.F64_eq(l0, float64(0)) != 0 {
					return v11
				} else {
					if base.F64_ne(base.F64_abs(l1), math.Float64frombits(uint64(0x7ff0000000000000))) != 0 {
						F_float_underflow_error(m)
						v34 = m.ExcPending
						if v34 != 0 {
							return float64(0)
						} else {
							base.Wasm_trap_unreachable()
							for {
							}
						}
					} else {
						return v11
					}
				}
			}
		}
	}
}
func F_float8_regr_sxy(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v4 int32
	_ = v4
	var v6 int32
	_ = v6
	var v8 int32
	_ = v8
	var v9 int32
	_ = v9
	var v12 int32
	_ = v12
	var v13 int32
	_ = v13
	var v16 int32
	_ = v16
	var v19 int32
	_ = v19
	var v20 int32
	_ = v20
	var v23 float64
	_ = v23
	var v26 int32
	_ = v26
	var v31 float64
	_ = v31
	var v32 int32
	_ = v32
	var v33 int32
	_ = v33
	var v34 int32
	_ = v34
	var v42 int32
	_ = v42
	var v49 int32
	_ = v49
	var v54 int32
	_ = v54
	v4 = m.G0
	v6 = v4 - int32(16)
	m.G0 = v6
	v8 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v9 = F_pg_detoast_datum(m, v8)
	mBase = m.M
	v12 = m.ExcPending
	if v12 != 0 {
		return int32(0)
	} else {
		v13 = *(*int32)(unsafe.Add(mBase, uint32(v9)+4))
		if v13 != int32(1) {
			F_errstart_cold(m, int32(21), int32(0))
			mBase = m.M
			v42 = m.ExcPending
			if v42 != 0 {
				return int32(0)
			} else {
				*(*int32)(unsafe.Add(mBase, uint32(v6)+4)) = int32(6)
				*(*int32)(unsafe.Add(mBase, uint32(v6))) = int32(8047)
				F_errmsg_internal(m, int32(26093), v6)
				mBase = m.M
				v49 = m.ExcPending
				if v49 != 0 {
					return int32(0)
				} else {
					F_errfinish(m, int32(490414), int32(2938), int32(24633))
					mBase = m.M
					v54 = m.ExcPending
					if v54 != 0 {
						return int32(0)
					} else {
						base.Wasm_trap_unreachable()
						for {
						}
					}
				}
			}
		} else {
			v16 = *(*int32)(unsafe.Add(mBase, uint32(v9)+16))
			if v16 != int32(6) {
				F_errstart_cold(m, int32(21), int32(0))
				mBase = m.M
				v42 = m.ExcPending
				if v42 != 0 {
					return int32(0)
				} else {
					*(*int32)(unsafe.Add(mBase, uint32(v6)+4)) = int32(6)
					*(*int32)(unsafe.Add(mBase, uint32(v6))) = int32(8047)
					F_errmsg_internal(m, int32(26093), v6)
					mBase = m.M
					v49 = m.ExcPending
					if v49 != 0 {
						return int32(0)
					} else {
						F_errfinish(m, int32(490414), int32(2938), int32(24633))
						mBase = m.M
						v54 = m.ExcPending
						if v54 != 0 {
							return int32(0)
						} else {
							base.Wasm_trap_unreachable()
							for {
							}
						}
					}
				}
			} else {
				v19 = *(*int32)(unsafe.Add(mBase, uint32(v9)+8))
				if v19 != 0 {
					F_errstart_cold(m, int32(21), int32(0))
					mBase = m.M
					v42 = m.ExcPending
					if v42 != 0 {
						return int32(0)
					} else {
						*(*int32)(unsafe.Add(mBase, uint32(v6)+4)) = int32(6)
						*(*int32)(unsafe.Add(mBase, uint32(v6))) = int32(8047)
						F_errmsg_internal(m, int32(26093), v6)
						mBase = m.M
						v49 = m.ExcPending
						if v49 != 0 {
							return int32(0)
						} else {
							F_errfinish(m, int32(490414), int32(2938), int32(24633))
							mBase = m.M
							v54 = m.ExcPending
							if v54 != 0 {
								return int32(0)
							} else {
								base.Wasm_trap_unreachable()
								for {
								}
							}
						}
					}
				} else {
					v20 = *(*int32)(unsafe.Add(mBase, uint32(v9)+12))
					if v20 != int32(701) {
						F_errstart_cold(m, int32(21), int32(0))
						mBase = m.M
						v42 = m.ExcPending
						if v42 != 0 {
							return int32(0)
						} else {
							*(*int32)(unsafe.Add(mBase, uint32(v6)+4)) = int32(6)
							*(*int32)(unsafe.Add(mBase, uint32(v6))) = int32(8047)
							F_errmsg_internal(m, int32(26093), v6)
							mBase = m.M
							v49 = m.ExcPending
							if v49 != 0 {
								return int32(0)
							} else {
								F_errfinish(m, int32(490414), int32(2938), int32(24633))
								mBase = m.M
								v54 = m.ExcPending
								if v54 != 0 {
									return int32(0)
								} else {
									base.Wasm_trap_unreachable()
									for {
									}
								}
							}
						}
					} else {
						v23 = *(*float64)(unsafe.Add(mBase, uint32(v9)+24))
						if base.F64_lt(v23, float64(1)) != 0 {
							v26 = int32(1)
							*(*uint8)(unsafe.Add(mBase, uint32(l0)+16)) = uint8(v26)
							v34 = int32(0)
							m.G0 = v6 + int32(16)
							return v34
						} else {
							v31 = *(*float64)(unsafe.Add(mBase, uint32(v9-int32(-64))))
							v32 = F_Float8GetDatum(m, v31)
							mBase = m.M
							v33 = m.ExcPending
							if v33 != 0 {
								return int32(0)
							} else {
								v34 = v32
								m.G0 = v6 + int32(16)
								return v34
							}
						}
					}
				}
			}
		}
	}
}
func F_float8_var_pop(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v5 int32
	_ = v5
	var v7 int32
	_ = v7
	var v9 int32
	_ = v9
	var v10 int32
	_ = v10
	var v13 int32
	_ = v13
	var v14 int32
	_ = v14
	var v17 int32
	_ = v17
	var v20 int32
	_ = v20
	var v21 int32
	_ = v21
	var v24 float64
	_ = v24
	var v27 int32
	_ = v27
	var v30 float64
	_ = v30
	var v32 int32
	_ = v32
	var v33 int32
	_ = v33
	var v34 int32
	_ = v34
	var v42 int32
	_ = v42
	var v49 int32
	_ = v49
	var v54 int32
	_ = v54
	v5 = m.G0
	v7 = v5 - int32(16)
	m.G0 = v7
	v9 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v10 = F_pg_detoast_datum(m, v9)
	mBase = m.M
	v13 = m.ExcPending
	if v13 != 0 {
		return int32(0)
	} else {
		v14 = *(*int32)(unsafe.Add(mBase, uint32(v10)+4))
		if v14 != int32(1) {
			F_errstart_cold(m, int32(21), int32(0))
			mBase = m.M
			v42 = m.ExcPending
			if v42 != 0 {
				return int32(0)
			} else {
				*(*int32)(unsafe.Add(mBase, uint32(v7)+4)) = int32(3)
				*(*int32)(unsafe.Add(mBase, uint32(v7))) = int32(233369)
				F_errmsg_internal(m, int32(26093), v7)
				mBase = m.M
				v49 = m.ExcPending
				if v49 != 0 {
					return int32(0)
				} else {
					F_errfinish(m, int32(490414), int32(2938), int32(24633))
					mBase = m.M
					v54 = m.ExcPending
					if v54 != 0 {
						return int32(0)
					} else {
						base.Wasm_trap_unreachable()
						for {
						}
					}
				}
			}
		} else {
			v17 = *(*int32)(unsafe.Add(mBase, uint32(v10)+16))
			if v17 != int32(3) {
				F_errstart_cold(m, int32(21), int32(0))
				mBase = m.M
				v42 = m.ExcPending
				if v42 != 0 {
					return int32(0)
				} else {
					*(*int32)(unsafe.Add(mBase, uint32(v7)+4)) = int32(3)
					*(*int32)(unsafe.Add(mBase, uint32(v7))) = int32(233369)
					F_errmsg_internal(m, int32(26093), v7)
					mBase = m.M
					v49 = m.ExcPending
					if v49 != 0 {
						return int32(0)
					} else {
						F_errfinish(m, int32(490414), int32(2938), int32(24633))
						mBase = m.M
						v54 = m.ExcPending
						if v54 != 0 {
							return int32(0)
						} else {
							base.Wasm_trap_unreachable()
							for {
							}
						}
					}
				}
			} else {
				v20 = *(*int32)(unsafe.Add(mBase, uint32(v10)+8))
				if v20 != 0 {
					F_errstart_cold(m, int32(21), int32(0))
					mBase = m.M
					v42 = m.ExcPending
					if v42 != 0 {
						return int32(0)
					} else {
						*(*int32)(unsafe.Add(mBase, uint32(v7)+4)) = int32(3)
						*(*int32)(unsafe.Add(mBase, uint32(v7))) = int32(233369)
						F_errmsg_internal(m, int32(26093), v7)
						mBase = m.M
						v49 = m.ExcPending
						if v49 != 0 {
							return int32(0)
						} else {
							F_errfinish(m, int32(490414), int32(2938), int32(24633))
							mBase = m.M
							v54 = m.ExcPending
							if v54 != 0 {
								return int32(0)
							} else {
								base.Wasm_trap_unreachable()
								for {
								}
							}
						}
					}
				} else {
					v21 = *(*int32)(unsafe.Add(mBase, uint32(v10)+12))
					if v21 != int32(701) {
						F_errstart_cold(m, int32(21), int32(0))
						mBase = m.M
						v42 = m.ExcPending
						if v42 != 0 {
							return int32(0)
						} else {
							*(*int32)(unsafe.Add(mBase, uint32(v7)+4)) = int32(3)
							*(*int32)(unsafe.Add(mBase, uint32(v7))) = int32(233369)
							F_errmsg_internal(m, int32(26093), v7)
							mBase = m.M
							v49 = m.ExcPending
							if v49 != 0 {
								return int32(0)
							} else {
								F_errfinish(m, int32(490414), int32(2938), int32(24633))
								mBase = m.M
								v54 = m.ExcPending
								if v54 != 0 {
									return int32(0)
								} else {
									base.Wasm_trap_unreachable()
									for {
									}
								}
							}
						}
					} else {
						v24 = *(*float64)(unsafe.Add(mBase, uint32(v10)+24))
						if base.F64_eq(v24, float64(0)) != 0 {
							v27 = int32(1)
							*(*uint8)(unsafe.Add(mBase, uint32(l0)+16)) = uint8(v27)
							v34 = int32(0)
							m.G0 = v7 + int32(16)
							return v34
						} else {
							v30 = *(*float64)(unsafe.Add(mBase, uint32(v10)+40))
							v32 = F_Float8GetDatum(m, base.F64_div(v30, v24))
							mBase = m.M
							v33 = m.ExcPending
							if v33 != 0 {
								return int32(0)
							} else {
								v34 = v32
								m.G0 = v7 + int32(16)
								return v34
							}
						}
					}
				}
			}
		}
	}
}
