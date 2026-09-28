package p4

import (
	base "github.com/shibukawa/pgmem/internal/aot/pgaot/base"
	"math"
	"unsafe"
)

func F_float8_accum(m *base.Module, l0 int32) int64 {
	mBase := m.M
	_ = mBase
	var v11 int32
	_ = v11
	var v13 int32
	_ = v13
	var v15 int32
	_ = v15
	var v16 int32
	_ = v16
	var v19 int32
	_ = v19
	var v20 int32
	_ = v20
	var v23 int32
	_ = v23
	var v26 int32
	_ = v26
	var v27 int32
	_ = v27
	var v30 float64
	_ = v30
	var v31 int64
	_ = v31
	var v32 float64
	_ = v32
	var v33 float64
	_ = v33
	var v34 float64
	_ = v34
	var v36 float64
	_ = v36
	var v37 float64
	_ = v37
	var v44 float64
	_ = v44
	var v48 float64
	_ = v48
	var v56 float64
	_ = v56
	var v63 int32
	_ = v63
	var v64 float64
	_ = v64
	var v69 float64
	_ = v69
	var v74 float64
	_ = v74
	var v75 float64
	_ = v75
	var v79 int32
	_ = v79
	var v82 int32
	_ = v82
	var v107 int32
	_ = v107
	var v118 int32
	_ = v118
	var v119 int32
	_ = v119
	var v120 int32
	_ = v120
	var v129 int32
	_ = v129
	var v136 int32
	_ = v136
	var v141 int32
	_ = v141
	v11 = m.G0
	v13 = v11 - int32(48)
	m.G0 = v13
	v15 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
	v16 = F_pg_detoast_datum(m, v15)
	mBase = m.M
	v19 = m.ExcPending
	if v19 != 0 {
		return int64(0)
	} else {
		v20 = *(*int32)(unsafe.Add(mBase, uint32(v16)+4))
		if v20 != int32(1) {
			F_errstart_cold(m, int32(21), int32(0))
			mBase = m.M
			v129 = m.ExcPending
			if v129 != 0 {
				return int64(0)
			} else {
				*(*int32)(unsafe.Add(mBase, uint32(v13)+4)) = int32(3)
				*(*int32)(unsafe.Add(mBase, uint32(v13))) = int32(_a_F_float8_accum_0)
				F_errmsg_internal(m, int32(_a_F_float8_accum_1), v13)
				mBase = m.M
				v136 = m.ExcPending
				if v136 != 0 {
					return int64(0)
				} else {
					F_errfinish(m, int32(_a_F_float8_accum_2), int32(2985), int32(_a_F_float8_accum_3))
					mBase = m.M
					v141 = m.ExcPending
					if v141 != 0 {
						return int64(0)
					} else {
						base.Wasm_trap_unreachable()
						for {
						}
					}
				}
			}
		} else {
			v23 = *(*int32)(unsafe.Add(mBase, uint32(v16)+16))
			if v23 != int32(3) {
				F_errstart_cold(m, int32(21), int32(0))
				mBase = m.M
				v129 = m.ExcPending
				if v129 != 0 {
					return int64(0)
				} else {
					*(*int32)(unsafe.Add(mBase, uint32(v13)+4)) = int32(3)
					*(*int32)(unsafe.Add(mBase, uint32(v13))) = int32(_a_F_float8_accum_0)
					F_errmsg_internal(m, int32(_a_F_float8_accum_1), v13)
					mBase = m.M
					v136 = m.ExcPending
					if v136 != 0 {
						return int64(0)
					} else {
						F_errfinish(m, int32(_a_F_float8_accum_2), int32(2985), int32(_a_F_float8_accum_3))
						mBase = m.M
						v141 = m.ExcPending
						if v141 != 0 {
							return int64(0)
						} else {
							base.Wasm_trap_unreachable()
							for {
							}
						}
					}
				}
			} else {
				v26 = *(*int32)(unsafe.Add(mBase, uint32(v16)+8))
				if v26 != 0 {
					F_errstart_cold(m, int32(21), int32(0))
					mBase = m.M
					v129 = m.ExcPending
					if v129 != 0 {
						return int64(0)
					} else {
						*(*int32)(unsafe.Add(mBase, uint32(v13)+4)) = int32(3)
						*(*int32)(unsafe.Add(mBase, uint32(v13))) = int32(_a_F_float8_accum_0)
						F_errmsg_internal(m, int32(_a_F_float8_accum_1), v13)
						mBase = m.M
						v136 = m.ExcPending
						if v136 != 0 {
							return int64(0)
						} else {
							F_errfinish(m, int32(_a_F_float8_accum_2), int32(2985), int32(_a_F_float8_accum_3))
							mBase = m.M
							v141 = m.ExcPending
							if v141 != 0 {
								return int64(0)
							} else {
								base.Wasm_trap_unreachable()
								for {
								}
							}
						}
					}
				} else {
					v27 = *(*int32)(unsafe.Add(mBase, uint32(v16)+12))
					if v27 != int32(701) {
						F_errstart_cold(m, int32(21), int32(0))
						mBase = m.M
						v129 = m.ExcPending
						if v129 != 0 {
							return int64(0)
						} else {
							*(*int32)(unsafe.Add(mBase, uint32(v13)+4)) = int32(3)
							*(*int32)(unsafe.Add(mBase, uint32(v13))) = int32(_a_F_float8_accum_0)
							F_errmsg_internal(m, int32(_a_F_float8_accum_1), v13)
							mBase = m.M
							v136 = m.ExcPending
							if v136 != 0 {
								return int64(0)
							} else {
								F_errfinish(m, int32(_a_F_float8_accum_2), int32(2985), int32(_a_F_float8_accum_3))
								mBase = m.M
								v141 = m.ExcPending
								if v141 != 0 {
									return int64(0)
								} else {
									base.Wasm_trap_unreachable()
									for {
									}
								}
							}
						}
					} else {
						v30 = *(*float64)(unsafe.Add(mBase, uint32(v16)+32))
						v31 = *(*int64)(unsafe.Add(mBase, uint32(l0)+40))
						v32 = base.F64_reinterpret_i64(v31)
						v33 = base.F64_add(v30, v32)
						v34 = *(*float64)(unsafe.Add(mBase, uint32(v16)+24))
						v36 = base.F64_add(v34, float64(1))
						v37 = *(*float64)(unsafe.Add(mBase, uint32(v16)+40))
						if base.F64_gt(v34, float64(0)) != 0 {
							if base.F64_ne(base.F64_abs(v33), math.Float64frombits(uint64(0x7ff0000000000000))) != 0 {
								v44 = base.F64_sub(base.F64_mul(v32, v36), v33)
								v48 = base.F64_add(v37, base.F64_div(base.F64_mul(v44, v44), base.F64_mul(v34, v36)))
								if base.F64_ne(base.F64_abs(v48), math.Float64frombits(uint64(0x7ff0000000000000))) != 0 {
									v75 = v48
									v79 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
									if v79 == int32(0) {
										v107 = int32(0)
									} else {
										v82 = *(*int32)(unsafe.Add(mBase, uint32(v79)))
										switch v82 - int32(435) {
										case 0:
											v107 = int32(1)
										case 1:
											v107 = int32(2)
										default:
											v107 = int32(0)
										}
									}
									if v107 != 0 {
										*(*float64)(unsafe.Add(mBase, uint32(v16)+40)) = v75
										*(*float64)(unsafe.Add(mBase, uint32(v16)+32)) = v33
										*(*float64)(unsafe.Add(mBase, uint32(v16)+24)) = v36
										v120 = v16
										m.G0 = v13 + int32(48)
										return base.I64_extend_i32_u(v120)
									} else {
										*(*float64)(unsafe.Add(mBase, uint32(v13)+32)) = v75
										*(*float64)(unsafe.Add(mBase, uint32(v13)+24)) = v33
										*(*float64)(unsafe.Add(mBase, uint32(v13)+16)) = v36
										v118 = F_construct_array_builtin(m, v13+int32(16), int32(3), int32(701))
										mBase = m.M
										v119 = m.ExcPending
										if v119 != 0 {
											return int64(0)
										} else {
											v120 = v118
											m.G0 = v13 + int32(48)
											return base.I64_extend_i32_u(v120)
										}
									}
								} else {
									v56 = math.Float64frombits(uint64(0x7ff0000000000000))
									if base.F64_eq(base.F64_abs(v30), v56)|base.F64_eq(base.F64_abs(v32), v56) != 0 {
										v75 = math.Float64frombits(uint64(0x7ff8000000000000))
										v79 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
										if v79 == int32(0) {
											v107 = int32(0)
										} else {
											v82 = *(*int32)(unsafe.Add(mBase, uint32(v79)))
											switch v82 - int32(435) {
											case 0:
												v107 = int32(1)
											case 1:
												v107 = int32(2)
											default:
												v107 = int32(0)
											}
										}
										if v107 != 0 {
											*(*float64)(unsafe.Add(mBase, uint32(v16)+40)) = v75
											*(*float64)(unsafe.Add(mBase, uint32(v16)+32)) = v33
											*(*float64)(unsafe.Add(mBase, uint32(v16)+24)) = v36
											v120 = v16
											m.G0 = v13 + int32(48)
											return base.I64_extend_i32_u(v120)
										} else {
											*(*float64)(unsafe.Add(mBase, uint32(v13)+32)) = v75
											*(*float64)(unsafe.Add(mBase, uint32(v13)+24)) = v33
											*(*float64)(unsafe.Add(mBase, uint32(v13)+16)) = v36
											v118 = F_construct_array_builtin(m, v13+int32(16), int32(3), int32(701))
											mBase = m.M
											v119 = m.ExcPending
											if v119 != 0 {
												return int64(0)
											} else {
												v120 = v118
												m.G0 = v13 + int32(48)
												return base.I64_extend_i32_u(v120)
											}
										}
									} else {
										F_float_overflow_error(m)
										mBase = m.M
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
							} else {
								v56 = math.Float64frombits(uint64(0x7ff0000000000000))
								if base.F64_eq(base.F64_abs(v30), v56)|base.F64_eq(base.F64_abs(v32), v56) != 0 {
									v75 = math.Float64frombits(uint64(0x7ff8000000000000))
									v79 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
									if v79 == int32(0) {
										v107 = int32(0)
									} else {
										v82 = *(*int32)(unsafe.Add(mBase, uint32(v79)))
										switch v82 - int32(435) {
										case 0:
											v107 = int32(1)
										case 1:
											v107 = int32(2)
										default:
											v107 = int32(0)
										}
									}
									if v107 != 0 {
										*(*float64)(unsafe.Add(mBase, uint32(v16)+40)) = v75
										*(*float64)(unsafe.Add(mBase, uint32(v16)+32)) = v33
										*(*float64)(unsafe.Add(mBase, uint32(v16)+24)) = v36
										v120 = v16
										m.G0 = v13 + int32(48)
										return base.I64_extend_i32_u(v120)
									} else {
										*(*float64)(unsafe.Add(mBase, uint32(v13)+32)) = v75
										*(*float64)(unsafe.Add(mBase, uint32(v13)+24)) = v33
										*(*float64)(unsafe.Add(mBase, uint32(v13)+16)) = v36
										v118 = F_construct_array_builtin(m, v13+int32(16), int32(3), int32(701))
										mBase = m.M
										v119 = m.ExcPending
										if v119 != 0 {
											return int64(0)
										} else {
											v120 = v118
											m.G0 = v13 + int32(48)
											return base.I64_extend_i32_u(v120)
										}
									}
								} else {
									F_float_overflow_error(m)
									mBase = m.M
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
						} else {
							v64 = math.Float64frombits(uint64(0x7ff8000000000000))
							if base.F64_eq(base.F64_abs(v32), math.Float64frombits(uint64(0x7ff0000000000000))) != 0 {
								v69 = v64
							} else {
								v69 = v37
							}
							if base.Ui64(int64(9218868437227405312)) < base.Ui64(v31&int64(9223372036854775807)) {
								v74 = v64
							} else {
								v74 = v69
							}
							v75 = v74
							v79 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
							if v79 == int32(0) {
								v107 = int32(0)
							} else {
								v82 = *(*int32)(unsafe.Add(mBase, uint32(v79)))
								switch v82 - int32(435) {
								case 0:
									v107 = int32(1)
								case 1:
									v107 = int32(2)
								default:
									v107 = int32(0)
								}
							}
							if v107 != 0 {
								*(*float64)(unsafe.Add(mBase, uint32(v16)+40)) = v75
								*(*float64)(unsafe.Add(mBase, uint32(v16)+32)) = v33
								*(*float64)(unsafe.Add(mBase, uint32(v16)+24)) = v36
								v120 = v16
								m.G0 = v13 + int32(48)
								return base.I64_extend_i32_u(v120)
							} else {
								*(*float64)(unsafe.Add(mBase, uint32(v13)+32)) = v75
								*(*float64)(unsafe.Add(mBase, uint32(v13)+24)) = v33
								*(*float64)(unsafe.Add(mBase, uint32(v13)+16)) = v36
								v118 = F_construct_array_builtin(m, v13+int32(16), int32(3), int32(701))
								mBase = m.M
								v119 = m.ExcPending
								if v119 != 0 {
									return int64(0)
								} else {
									v120 = v118
									m.G0 = v13 + int32(48)
									return base.I64_extend_i32_u(v120)
								}
							}
						}
					}
				}
			}
		}
	}
}
func F_float8_covar_pop(m *base.Module, l0 int32) int64 {
	mBase := m.M
	_ = mBase
	var v6 int32
	_ = v6
	var v8 int32
	_ = v8
	var v10 int32
	_ = v10
	var v11 int32
	_ = v11
	var v14 int32
	_ = v14
	var v15 int32
	_ = v15
	var v18 int32
	_ = v18
	var v21 int32
	_ = v21
	var v22 int32
	_ = v22
	var v25 float64
	_ = v25
	var v28 int32
	_ = v28
	var v31 float64
	_ = v31
	var v34 int64
	_ = v34
	var v42 int32
	_ = v42
	var v49 int32
	_ = v49
	var v54 int32
	_ = v54
	v6 = m.G0
	v8 = v6 - int32(16)
	m.G0 = v8
	v10 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
	v11 = F_pg_detoast_datum(m, v10)
	mBase = m.M
	v14 = m.ExcPending
	if v14 != 0 {
		return int64(0)
	} else {
		v15 = *(*int32)(unsafe.Add(mBase, uint32(v11)+4))
		if v15 != int32(1) {
			F_errstart_cold(m, int32(21), int32(0))
			mBase = m.M
			v42 = m.ExcPending
			if v42 != 0 {
				return int64(0)
			} else {
				*(*int32)(unsafe.Add(mBase, uint32(v8)+4)) = int32(8)
				*(*int32)(unsafe.Add(mBase, uint32(v8))) = int32(_a_F_float8_covar_pop_0)
				F_errmsg_internal(m, int32(_a_F_float8_covar_pop_1), v8)
				mBase = m.M
				v49 = m.ExcPending
				if v49 != 0 {
					return int64(0)
				} else {
					F_errfinish(m, int32(_a_F_float8_covar_pop_2), int32(2985), int32(_a_F_float8_covar_pop_3))
					mBase = m.M
					v54 = m.ExcPending
					if v54 != 0 {
						return int64(0)
					} else {
						base.Wasm_trap_unreachable()
						for {
						}
					}
				}
			}
		} else {
			v18 = *(*int32)(unsafe.Add(mBase, uint32(v11)+16))
			if v18 != int32(8) {
				F_errstart_cold(m, int32(21), int32(0))
				mBase = m.M
				v42 = m.ExcPending
				if v42 != 0 {
					return int64(0)
				} else {
					*(*int32)(unsafe.Add(mBase, uint32(v8)+4)) = int32(8)
					*(*int32)(unsafe.Add(mBase, uint32(v8))) = int32(_a_F_float8_covar_pop_0)
					F_errmsg_internal(m, int32(_a_F_float8_covar_pop_1), v8)
					mBase = m.M
					v49 = m.ExcPending
					if v49 != 0 {
						return int64(0)
					} else {
						F_errfinish(m, int32(_a_F_float8_covar_pop_2), int32(2985), int32(_a_F_float8_covar_pop_3))
						mBase = m.M
						v54 = m.ExcPending
						if v54 != 0 {
							return int64(0)
						} else {
							base.Wasm_trap_unreachable()
							for {
							}
						}
					}
				}
			} else {
				v21 = *(*int32)(unsafe.Add(mBase, uint32(v11)+8))
				if v21 != 0 {
					F_errstart_cold(m, int32(21), int32(0))
					mBase = m.M
					v42 = m.ExcPending
					if v42 != 0 {
						return int64(0)
					} else {
						*(*int32)(unsafe.Add(mBase, uint32(v8)+4)) = int32(8)
						*(*int32)(unsafe.Add(mBase, uint32(v8))) = int32(_a_F_float8_covar_pop_0)
						F_errmsg_internal(m, int32(_a_F_float8_covar_pop_1), v8)
						mBase = m.M
						v49 = m.ExcPending
						if v49 != 0 {
							return int64(0)
						} else {
							F_errfinish(m, int32(_a_F_float8_covar_pop_2), int32(2985), int32(_a_F_float8_covar_pop_3))
							mBase = m.M
							v54 = m.ExcPending
							if v54 != 0 {
								return int64(0)
							} else {
								base.Wasm_trap_unreachable()
								for {
								}
							}
						}
					}
				} else {
					v22 = *(*int32)(unsafe.Add(mBase, uint32(v11)+12))
					if v22 != int32(701) {
						F_errstart_cold(m, int32(21), int32(0))
						mBase = m.M
						v42 = m.ExcPending
						if v42 != 0 {
							return int64(0)
						} else {
							*(*int32)(unsafe.Add(mBase, uint32(v8)+4)) = int32(8)
							*(*int32)(unsafe.Add(mBase, uint32(v8))) = int32(_a_F_float8_covar_pop_0)
							F_errmsg_internal(m, int32(_a_F_float8_covar_pop_1), v8)
							mBase = m.M
							v49 = m.ExcPending
							if v49 != 0 {
								return int64(0)
							} else {
								F_errfinish(m, int32(_a_F_float8_covar_pop_2), int32(2985), int32(_a_F_float8_covar_pop_3))
								mBase = m.M
								v54 = m.ExcPending
								if v54 != 0 {
									return int64(0)
								} else {
									base.Wasm_trap_unreachable()
									for {
									}
								}
							}
						}
					} else {
						v25 = *(*float64)(unsafe.Add(mBase, uint32(v11)+24))
						if base.F64_lt(v25, float64(1)) != 0 {
							v28 = int32(1)
							*(*uint8)(unsafe.Add(mBase, uint32(l0)+16)) = uint8(v28)
							v34 = int64(0)
						} else {
							v31 = *(*float64)(unsafe.Add(mBase, uint32(v11)+64))
							v34 = base.I64_reinterpret_f64(base.F64_div(v31, v25))
						}
						m.G0 = v8 + int32(16)
						return v34
					}
				}
			}
		}
	}
}
func F_float8_div(m *base.Module, l0 float64, l1 float64) float64 {
	var v15 float64
	_ = v15
	var v18 int32
	_ = v18
	var v21 float64
	_ = v21
	var v23 float64
	_ = v23
	var v31 float64
	_ = v31
	var v32 int32
	_ = v32
	var v34 float64
	_ = v34
	var v44 float64
	_ = v44
	var v45 int32
	_ = v45
	var v46 float64
	_ = v46
	if base.B2i32(base.Ui64(int64(9218868437227405312)) < base.Ui64(base.I64_reinterpret_f64(l0)&int64(9223372036854775807)))|base.F64_ne(l1, float64(0)) == int32(0) {
		v15 = F_float_zero_divide_error_ext(m, int32(0))
		v18 = m.ExcPending
		if v18 != 0 {
			return float64(0)
		} else {
			return v15
		}
	} else {
		v21 = math.Float64frombits(uint64(0x7ff0000000000000))
		v23 = base.F64_div(l0, l1)
		if base.F64_eq(base.F64_abs(l0), v21)|base.F64_ne(base.F64_abs(v23), v21) == int32(0) {
			v31 = F_float_overflow_error_ext(m, int32(0))
			v32 = m.ExcPending
			if v32 != 0 {
				return float64(0)
			} else {
				return v31
			}
		} else {
			v34 = float64(0)
			if base.F64_eq(l0, v34)|base.F64_ne(v23, v34)|base.F64_eq(base.F64_abs(l1), math.Float64frombits(uint64(0x7ff0000000000000))) != 0 {
				v46 = v23
				return v46
			} else {
				v44 = F_float_underflow_error_ext(m, int32(0))
				v45 = m.ExcPending
				if v45 != 0 {
					return float64(0)
				} else {
					v46 = v44
					return v46
				}
			}
		}
	}
}
func F_float8_regr_sxy(m *base.Module, l0 int32) int64 {
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
	var v30 int64
	_ = v30
	var v31 int64
	_ = v31
	var v39 int32
	_ = v39
	var v46 int32
	_ = v46
	var v51 int32
	_ = v51
	v5 = m.G0
	v7 = v5 - int32(16)
	m.G0 = v7
	v9 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
	v10 = F_pg_detoast_datum(m, v9)
	mBase = m.M
	v13 = m.ExcPending
	if v13 != 0 {
		return int64(0)
	} else {
		v14 = *(*int32)(unsafe.Add(mBase, uint32(v10)+4))
		if v14 != int32(1) {
			F_errstart_cold(m, int32(21), int32(0))
			mBase = m.M
			v39 = m.ExcPending
			if v39 != 0 {
				return int64(0)
			} else {
				*(*int32)(unsafe.Add(mBase, uint32(v7)+4)) = int32(8)
				*(*int32)(unsafe.Add(mBase, uint32(v7))) = int32(_a_F_float8_regr_sxy_0)
				F_errmsg_internal(m, int32(_a_F_float8_regr_sxy_1), v7)
				mBase = m.M
				v46 = m.ExcPending
				if v46 != 0 {
					return int64(0)
				} else {
					F_errfinish(m, int32(_a_F_float8_regr_sxy_2), int32(2985), int32(_a_F_float8_regr_sxy_3))
					mBase = m.M
					v51 = m.ExcPending
					if v51 != 0 {
						return int64(0)
					} else {
						base.Wasm_trap_unreachable()
						for {
						}
					}
				}
			}
		} else {
			v17 = *(*int32)(unsafe.Add(mBase, uint32(v10)+16))
			if v17 != int32(8) {
				F_errstart_cold(m, int32(21), int32(0))
				mBase = m.M
				v39 = m.ExcPending
				if v39 != 0 {
					return int64(0)
				} else {
					*(*int32)(unsafe.Add(mBase, uint32(v7)+4)) = int32(8)
					*(*int32)(unsafe.Add(mBase, uint32(v7))) = int32(_a_F_float8_regr_sxy_0)
					F_errmsg_internal(m, int32(_a_F_float8_regr_sxy_1), v7)
					mBase = m.M
					v46 = m.ExcPending
					if v46 != 0 {
						return int64(0)
					} else {
						F_errfinish(m, int32(_a_F_float8_regr_sxy_2), int32(2985), int32(_a_F_float8_regr_sxy_3))
						mBase = m.M
						v51 = m.ExcPending
						if v51 != 0 {
							return int64(0)
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
					v39 = m.ExcPending
					if v39 != 0 {
						return int64(0)
					} else {
						*(*int32)(unsafe.Add(mBase, uint32(v7)+4)) = int32(8)
						*(*int32)(unsafe.Add(mBase, uint32(v7))) = int32(_a_F_float8_regr_sxy_0)
						F_errmsg_internal(m, int32(_a_F_float8_regr_sxy_1), v7)
						mBase = m.M
						v46 = m.ExcPending
						if v46 != 0 {
							return int64(0)
						} else {
							F_errfinish(m, int32(_a_F_float8_regr_sxy_2), int32(2985), int32(_a_F_float8_regr_sxy_3))
							mBase = m.M
							v51 = m.ExcPending
							if v51 != 0 {
								return int64(0)
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
						v39 = m.ExcPending
						if v39 != 0 {
							return int64(0)
						} else {
							*(*int32)(unsafe.Add(mBase, uint32(v7)+4)) = int32(8)
							*(*int32)(unsafe.Add(mBase, uint32(v7))) = int32(_a_F_float8_regr_sxy_0)
							F_errmsg_internal(m, int32(_a_F_float8_regr_sxy_1), v7)
							mBase = m.M
							v46 = m.ExcPending
							if v46 != 0 {
								return int64(0)
							} else {
								F_errfinish(m, int32(_a_F_float8_regr_sxy_2), int32(2985), int32(_a_F_float8_regr_sxy_3))
								mBase = m.M
								v51 = m.ExcPending
								if v51 != 0 {
									return int64(0)
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
							v31 = int64(0)
						} else {
							v30 = *(*int64)(unsafe.Add(mBase, uint32(v10)+64))
							v31 = v30
						}
						m.G0 = v7 + int32(16)
						return v31
					}
				}
			}
		}
	}
}
func F_float8_var_pop(m *base.Module, l0 int32) int64 {
	mBase := m.M
	_ = mBase
	var v6 int32
	_ = v6
	var v8 int32
	_ = v8
	var v10 int32
	_ = v10
	var v11 int32
	_ = v11
	var v14 int32
	_ = v14
	var v15 int32
	_ = v15
	var v18 int32
	_ = v18
	var v21 int32
	_ = v21
	var v22 int32
	_ = v22
	var v25 float64
	_ = v25
	var v28 int32
	_ = v28
	var v31 float64
	_ = v31
	var v34 int64
	_ = v34
	var v42 int32
	_ = v42
	var v49 int32
	_ = v49
	var v54 int32
	_ = v54
	v6 = m.G0
	v8 = v6 - int32(16)
	m.G0 = v8
	v10 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
	v11 = F_pg_detoast_datum(m, v10)
	mBase = m.M
	v14 = m.ExcPending
	if v14 != 0 {
		return int64(0)
	} else {
		v15 = *(*int32)(unsafe.Add(mBase, uint32(v11)+4))
		if v15 != int32(1) {
			F_errstart_cold(m, int32(21), int32(0))
			mBase = m.M
			v42 = m.ExcPending
			if v42 != 0 {
				return int64(0)
			} else {
				*(*int32)(unsafe.Add(mBase, uint32(v8)+4)) = int32(3)
				*(*int32)(unsafe.Add(mBase, uint32(v8))) = int32(_a_F_float8_var_pop_0)
				F_errmsg_internal(m, int32(_a_F_float8_var_pop_1), v8)
				mBase = m.M
				v49 = m.ExcPending
				if v49 != 0 {
					return int64(0)
				} else {
					F_errfinish(m, int32(_a_F_float8_var_pop_2), int32(2985), int32(_a_F_float8_var_pop_3))
					mBase = m.M
					v54 = m.ExcPending
					if v54 != 0 {
						return int64(0)
					} else {
						base.Wasm_trap_unreachable()
						for {
						}
					}
				}
			}
		} else {
			v18 = *(*int32)(unsafe.Add(mBase, uint32(v11)+16))
			if v18 != int32(3) {
				F_errstart_cold(m, int32(21), int32(0))
				mBase = m.M
				v42 = m.ExcPending
				if v42 != 0 {
					return int64(0)
				} else {
					*(*int32)(unsafe.Add(mBase, uint32(v8)+4)) = int32(3)
					*(*int32)(unsafe.Add(mBase, uint32(v8))) = int32(_a_F_float8_var_pop_0)
					F_errmsg_internal(m, int32(_a_F_float8_var_pop_1), v8)
					mBase = m.M
					v49 = m.ExcPending
					if v49 != 0 {
						return int64(0)
					} else {
						F_errfinish(m, int32(_a_F_float8_var_pop_2), int32(2985), int32(_a_F_float8_var_pop_3))
						mBase = m.M
						v54 = m.ExcPending
						if v54 != 0 {
							return int64(0)
						} else {
							base.Wasm_trap_unreachable()
							for {
							}
						}
					}
				}
			} else {
				v21 = *(*int32)(unsafe.Add(mBase, uint32(v11)+8))
				if v21 != 0 {
					F_errstart_cold(m, int32(21), int32(0))
					mBase = m.M
					v42 = m.ExcPending
					if v42 != 0 {
						return int64(0)
					} else {
						*(*int32)(unsafe.Add(mBase, uint32(v8)+4)) = int32(3)
						*(*int32)(unsafe.Add(mBase, uint32(v8))) = int32(_a_F_float8_var_pop_0)
						F_errmsg_internal(m, int32(_a_F_float8_var_pop_1), v8)
						mBase = m.M
						v49 = m.ExcPending
						if v49 != 0 {
							return int64(0)
						} else {
							F_errfinish(m, int32(_a_F_float8_var_pop_2), int32(2985), int32(_a_F_float8_var_pop_3))
							mBase = m.M
							v54 = m.ExcPending
							if v54 != 0 {
								return int64(0)
							} else {
								base.Wasm_trap_unreachable()
								for {
								}
							}
						}
					}
				} else {
					v22 = *(*int32)(unsafe.Add(mBase, uint32(v11)+12))
					if v22 != int32(701) {
						F_errstart_cold(m, int32(21), int32(0))
						mBase = m.M
						v42 = m.ExcPending
						if v42 != 0 {
							return int64(0)
						} else {
							*(*int32)(unsafe.Add(mBase, uint32(v8)+4)) = int32(3)
							*(*int32)(unsafe.Add(mBase, uint32(v8))) = int32(_a_F_float8_var_pop_0)
							F_errmsg_internal(m, int32(_a_F_float8_var_pop_1), v8)
							mBase = m.M
							v49 = m.ExcPending
							if v49 != 0 {
								return int64(0)
							} else {
								F_errfinish(m, int32(_a_F_float8_var_pop_2), int32(2985), int32(_a_F_float8_var_pop_3))
								mBase = m.M
								v54 = m.ExcPending
								if v54 != 0 {
									return int64(0)
								} else {
									base.Wasm_trap_unreachable()
									for {
									}
								}
							}
						}
					} else {
						v25 = *(*float64)(unsafe.Add(mBase, uint32(v11)+24))
						if base.F64_eq(v25, float64(0)) != 0 {
							v28 = int32(1)
							*(*uint8)(unsafe.Add(mBase, uint32(l0)+16)) = uint8(v28)
							v34 = int64(0)
						} else {
							v31 = *(*float64)(unsafe.Add(mBase, uint32(v11)+40))
							v34 = base.I64_reinterpret_f64(base.F64_div(v31, v25))
						}
						m.G0 = v8 + int32(16)
						return v34
					}
				}
			}
		}
	}
}
