package p5

import (
	base "github.com/shibukawa/pgmem/internal/aot/pgaot/base"
	"math"
	"unsafe"
)

func F_Float8GetDatum(m *base.Module, l0 float64) int32 {
	mBase := m.M
	_ = mBase
	var v4 int32
	_ = v4
	var v7 int32
	_ = v7
	v4 = F_palloc(m, int32(8))
	mBase = m.M
	v7 = m.ExcPending
	if v7 != 0 {
		return int32(0)
	} else {
		*(*float64)(unsafe.Add(mBase, uint32(v4))) = l0
		return v4
	}
}
func F_float8_corr(m *base.Module, l0 int32) int32 {
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
	var v34 float64
	_ = v34
	var v38 int32
	_ = v38
	var v41 float64
	_ = v41
	var v45 int32
	_ = v45
	var v46 int32
	_ = v46
	var v49 int32
	_ = v49
	var v57 int32
	_ = v57
	var v64 int32
	_ = v64
	var v69 int32
	_ = v69
	v6 = m.G0
	v8 = v6 - int32(16)
	m.G0 = v8
	v10 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v11 = F_pg_detoast_datum(m, v10)
	mBase = m.M
	v14 = m.ExcPending
	if v14 != 0 {
		return int32(0)
	} else {
		v15 = *(*int32)(unsafe.Add(mBase, uint32(v11)+4))
		if v15 != int32(1) {
			F_errstart_cold(m, int32(21), int32(0))
			mBase = m.M
			v57 = m.ExcPending
			if v57 != 0 {
				return int32(0)
			} else {
				*(*int32)(unsafe.Add(mBase, uint32(v8)+4)) = int32(6)
				*(*int32)(unsafe.Add(mBase, uint32(v8))) = int32(_a_F_float8_corr_0)
				F_errmsg_internal(m, int32(_a_F_float8_corr_1), v8)
				mBase = m.M
				v64 = m.ExcPending
				if v64 != 0 {
					return int32(0)
				} else {
					F_errfinish(m, int32(_a_F_float8_corr_2), int32(2938), int32(_a_F_float8_corr_3))
					mBase = m.M
					v69 = m.ExcPending
					if v69 != 0 {
						return int32(0)
					} else {
						base.Wasm_trap_unreachable()
						for {
						}
					}
				}
			}
		} else {
			v18 = *(*int32)(unsafe.Add(mBase, uint32(v11)+16))
			if v18 != int32(6) {
				F_errstart_cold(m, int32(21), int32(0))
				mBase = m.M
				v57 = m.ExcPending
				if v57 != 0 {
					return int32(0)
				} else {
					*(*int32)(unsafe.Add(mBase, uint32(v8)+4)) = int32(6)
					*(*int32)(unsafe.Add(mBase, uint32(v8))) = int32(_a_F_float8_corr_0)
					F_errmsg_internal(m, int32(_a_F_float8_corr_1), v8)
					mBase = m.M
					v64 = m.ExcPending
					if v64 != 0 {
						return int32(0)
					} else {
						F_errfinish(m, int32(_a_F_float8_corr_2), int32(2938), int32(_a_F_float8_corr_3))
						mBase = m.M
						v69 = m.ExcPending
						if v69 != 0 {
							return int32(0)
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
					v57 = m.ExcPending
					if v57 != 0 {
						return int32(0)
					} else {
						*(*int32)(unsafe.Add(mBase, uint32(v8)+4)) = int32(6)
						*(*int32)(unsafe.Add(mBase, uint32(v8))) = int32(_a_F_float8_corr_0)
						F_errmsg_internal(m, int32(_a_F_float8_corr_1), v8)
						mBase = m.M
						v64 = m.ExcPending
						if v64 != 0 {
							return int32(0)
						} else {
							F_errfinish(m, int32(_a_F_float8_corr_2), int32(2938), int32(_a_F_float8_corr_3))
							mBase = m.M
							v69 = m.ExcPending
							if v69 != 0 {
								return int32(0)
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
						v57 = m.ExcPending
						if v57 != 0 {
							return int32(0)
						} else {
							*(*int32)(unsafe.Add(mBase, uint32(v8)+4)) = int32(6)
							*(*int32)(unsafe.Add(mBase, uint32(v8))) = int32(_a_F_float8_corr_0)
							F_errmsg_internal(m, int32(_a_F_float8_corr_1), v8)
							mBase = m.M
							v64 = m.ExcPending
							if v64 != 0 {
								return int32(0)
							} else {
								F_errfinish(m, int32(_a_F_float8_corr_2), int32(2938), int32(_a_F_float8_corr_3))
								mBase = m.M
								v69 = m.ExcPending
								if v69 != 0 {
									return int32(0)
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
							v49 = int32(0)
							m.G0 = v8 + int32(16)
							return v49
						} else {
							v31 = *(*float64)(unsafe.Add(mBase, uint32(v11)+40))
							if base.F64_ne(v31, float64(0)) != 0 {
								v34 = *(*float64)(unsafe.Add(mBase, uint32(v11)+56))
								if base.F64_ne(v34, float64(0)) != 0 {
									v41 = *(*float64)(unsafe.Add(mBase, uint32(v11)+64))
									v45 = F_Float8GetDatum(m, base.F64_div(v41, base.F64_sqrt(base.F64_mul(v31, v34))))
									mBase = m.M
									v46 = m.ExcPending
									if v46 != 0 {
										return int32(0)
									} else {
										v49 = v45
										m.G0 = v8 + int32(16)
										return v49
									}
								} else {
									v38 = int32(1)
									*(*uint8)(unsafe.Add(mBase, uint32(l0)+16)) = uint8(v38)
									v49 = int32(0)
									m.G0 = v8 + int32(16)
									return v49
								}
							} else {
								v38 = int32(1)
								*(*uint8)(unsafe.Add(mBase, uint32(l0)+16)) = uint8(v38)
								v49 = int32(0)
								m.G0 = v8 + int32(16)
								return v49
							}
						}
					}
				}
			}
		}
	}
}
func F_float8_mul(m *base.Module, l0 float64, l1 float64) float64 {
	var v5 float64
	_ = v5
	var v7 float64
	_ = v7
	var v12 int32
	_ = v12
	var v20 float64
	_ = v20
	var v34 int32
	_ = v34
	var v36 int32
	_ = v36
	v5 = math.Float64frombits(uint64(0x7ff0000000000000))
	v7 = base.F64_mul(l0, l1)
	v12 = int32(0)
	if base.B2i32(base.F64_eq(base.F64_abs(l0), v5)|base.F64_ne(base.F64_abs(v7), v5) == v12)&base.F64_ne(base.F64_abs(l1), v5) == v12 {
		v20 = float64(0)
		if base.B2i32(base.F64_eq(l0, v20)|base.F64_ne(v7, v20) == int32(0))&base.F64_ne(l1, v20) != 0 {
			F_float_underflow_error(m)
			v36 = m.ExcPending
			if v36 != 0 {
				return float64(0)
			} else {
				base.Wasm_trap_unreachable()
				for {
				}
			}
		} else {
			return v7
		}
	} else {
		F_float_overflow_error(m)
		v34 = m.ExcPending
		if v34 != 0 {
			return float64(0)
		} else {
			base.Wasm_trap_unreachable()
			for {
			}
		}
	}
}
func F_float8_qsort_cmp(m *base.Module, l0 int32, l1 int32, l2 int32) int32 {
	mBase := m.M
	_ = mBase
	var v7 float64
	_ = v7
	var v8 float64
	_ = v8
	var v11 int32
	_ = v11
	v7 = *(*float64)(unsafe.Add(mBase, uint32(l0)))
	v8 = *(*float64)(unsafe.Add(mBase, uint32(l1)))
	if base.F64_lt(v7, v8) != 0 {
		v11 = int32(-1)
	} else {
		v11 = base.F64_ne(v7, v8)
	}
	return v11
}
func F_float8_regr_accum(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v19 int32
	_ = v19
	var v21 int32
	_ = v21
	var v23 int32
	_ = v23
	var v24 int32
	_ = v24
	var v27 int32
	_ = v27
	var v28 int32
	_ = v28
	var v31 int32
	_ = v31
	var v34 int32
	_ = v34
	var v35 int32
	_ = v35
	var v38 int32
	_ = v38
	var v39 float64
	_ = v39
	var v40 int32
	_ = v40
	var v41 float64
	_ = v41
	var v42 float64
	_ = v42
	var v43 float64
	_ = v43
	var v44 float64
	_ = v44
	var v46 float64
	_ = v46
	var v47 float64
	_ = v47
	var v49 float64
	_ = v49
	var v51 float64
	_ = v51
	var v54 float64
	_ = v54
	var v56 float64
	_ = v56
	var v61 float64
	_ = v61
	var v65 float64
	_ = v65
	var v67 float64
	_ = v67
	var v70 float64
	_ = v70
	var v73 float64
	_ = v73
	var v77 float64
	_ = v77
	var v82 float64
	_ = v82
	var v83 float64
	_ = v83
	var v85 float64
	_ = v85
	var v104 float64
	_ = v104
	var v111 float64
	_ = v111
	var v114 float64
	_ = v114
	var v115 int32
	_ = v115
	var v118 int32
	_ = v118
	var v121 float64
	_ = v121
	var v130 float64
	_ = v130
	var v132 float64
	_ = v132
	var v150 float64
	_ = v150
	var v156 float64
	_ = v156
	var v161 float64
	_ = v161
	var v170 int64
	_ = v170
	var v174 float64
	_ = v174
	var v176 float64
	_ = v176
	var v177 float64
	_ = v177
	var v178 float64
	_ = v178
	var v195 float64
	_ = v195
	var v196 float64
	_ = v196
	var v200 int32
	_ = v200
	var v208 float64
	_ = v208
	var v209 float64
	_ = v209
	var v210 float64
	_ = v210
	var v216 int32
	_ = v216
	var v219 int32
	_ = v219
	var v244 int32
	_ = v244
	var v273 int32
	_ = v273
	var v274 int32
	_ = v274
	var v275 int32
	_ = v275
	var v283 int32
	_ = v283
	var v290 int32
	_ = v290
	var v295 int32
	_ = v295
	var v302 int32
	_ = v302
	v19 = m.G0
	v21 = v19 - int32(96)
	m.G0 = v21
	v23 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v24 = F_pg_detoast_datum(m, v23)
	mBase = m.M
	v27 = m.ExcPending
	if v27 != 0 {
		return int32(0)
	} else {
		v28 = *(*int32)(unsafe.Add(mBase, uint32(v24)+4))
		if v28 != int32(1) {
			F_errstart_cold(m, int32(21), int32(0))
			mBase = m.M
			v283 = m.ExcPending
			if v283 != 0 {
				return int32(0)
			} else {
				*(*int32)(unsafe.Add(mBase, uint32(v21)+4)) = int32(6)
				*(*int32)(unsafe.Add(mBase, uint32(v21))) = int32(_a_F_float8_regr_accum_0)
				F_errmsg_internal(m, int32(_a_F_float8_regr_accum_1), v21)
				mBase = m.M
				v290 = m.ExcPending
				if v290 != 0 {
					return int32(0)
				} else {
					F_errfinish(m, int32(_a_F_float8_regr_accum_2), int32(2938), int32(_a_F_float8_regr_accum_3))
					mBase = m.M
					v295 = m.ExcPending
					if v295 != 0 {
						return int32(0)
					} else {
						base.Wasm_trap_unreachable()
						for {
						}
					}
				}
			}
		} else {
			v31 = *(*int32)(unsafe.Add(mBase, uint32(v24)+16))
			if v31 != int32(6) {
				F_errstart_cold(m, int32(21), int32(0))
				mBase = m.M
				v283 = m.ExcPending
				if v283 != 0 {
					return int32(0)
				} else {
					*(*int32)(unsafe.Add(mBase, uint32(v21)+4)) = int32(6)
					*(*int32)(unsafe.Add(mBase, uint32(v21))) = int32(_a_F_float8_regr_accum_0)
					F_errmsg_internal(m, int32(_a_F_float8_regr_accum_1), v21)
					mBase = m.M
					v290 = m.ExcPending
					if v290 != 0 {
						return int32(0)
					} else {
						F_errfinish(m, int32(_a_F_float8_regr_accum_2), int32(2938), int32(_a_F_float8_regr_accum_3))
						mBase = m.M
						v295 = m.ExcPending
						if v295 != 0 {
							return int32(0)
						} else {
							base.Wasm_trap_unreachable()
							for {
							}
						}
					}
				}
			} else {
				v34 = *(*int32)(unsafe.Add(mBase, uint32(v24)+8))
				if v34 != 0 {
					F_errstart_cold(m, int32(21), int32(0))
					mBase = m.M
					v283 = m.ExcPending
					if v283 != 0 {
						return int32(0)
					} else {
						*(*int32)(unsafe.Add(mBase, uint32(v21)+4)) = int32(6)
						*(*int32)(unsafe.Add(mBase, uint32(v21))) = int32(_a_F_float8_regr_accum_0)
						F_errmsg_internal(m, int32(_a_F_float8_regr_accum_1), v21)
						mBase = m.M
						v290 = m.ExcPending
						if v290 != 0 {
							return int32(0)
						} else {
							F_errfinish(m, int32(_a_F_float8_regr_accum_2), int32(2938), int32(_a_F_float8_regr_accum_3))
							mBase = m.M
							v295 = m.ExcPending
							if v295 != 0 {
								return int32(0)
							} else {
								base.Wasm_trap_unreachable()
								for {
								}
							}
						}
					}
				} else {
					v35 = *(*int32)(unsafe.Add(mBase, uint32(v24)+12))
					if v35 != int32(701) {
						F_errstart_cold(m, int32(21), int32(0))
						mBase = m.M
						v283 = m.ExcPending
						if v283 != 0 {
							return int32(0)
						} else {
							*(*int32)(unsafe.Add(mBase, uint32(v21)+4)) = int32(6)
							*(*int32)(unsafe.Add(mBase, uint32(v21))) = int32(_a_F_float8_regr_accum_0)
							F_errmsg_internal(m, int32(_a_F_float8_regr_accum_1), v21)
							mBase = m.M
							v290 = m.ExcPending
							if v290 != 0 {
								return int32(0)
							} else {
								F_errfinish(m, int32(_a_F_float8_regr_accum_2), int32(2938), int32(_a_F_float8_regr_accum_3))
								mBase = m.M
								v295 = m.ExcPending
								if v295 != 0 {
									return int32(0)
								} else {
									base.Wasm_trap_unreachable()
									for {
									}
								}
							}
						}
					} else {
						v38 = *(*int32)(unsafe.Add(mBase, uint32(l0)+36))
						v39 = *(*float64)(unsafe.Add(mBase, uint32(v38)))
						v40 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
						v41 = *(*float64)(unsafe.Add(mBase, uint32(v40)))
						v42 = *(*float64)(unsafe.Add(mBase, uint32(v24)+32))
						v43 = *(*float64)(unsafe.Add(mBase, uint32(v24)+24))
						v44 = *(*float64)(unsafe.Add(mBase, uint32(v24)+40))
						*(*float64)(unsafe.Add(mBase, uint32(v21)+72)) = v44
						v46 = *(*float64)(unsafe.Add(mBase, uint32(v24)+48))
						v47 = *(*float64)(unsafe.Add(mBase, uint32(v24)+56))
						*(*float64)(unsafe.Add(mBase, uint32(v21)+56)) = v47
						v49 = *(*float64)(unsafe.Add(mBase, uint32(v24)+64))
						v51 = base.F64_add(v43, float64(1))
						*(*float64)(unsafe.Add(mBase, uint32(v21)+88)) = v51
						*(*float64)(unsafe.Add(mBase, uint32(v21)+48)) = v49
						v54 = base.F64_add(v39, v42)
						*(*float64)(unsafe.Add(mBase, uint32(v21)+80)) = v54
						v56 = base.F64_add(v41, v46)
						*(*float64)(unsafe.Add(mBase, uint32(v21)+64)) = v56
						if base.F64_gt(v43, float64(0)) != 0 {
							v61 = base.F64_sub(base.F64_mul(v41, v51), v56)
							v65 = base.F64_div(float64(1), base.F64_mul(v43, v51))
							v67 = base.F64_add(base.F64_mul(base.F64_mul(v61, v61), v65), v47)
							*(*float64)(unsafe.Add(mBase, uint32(v21)+56)) = v67
							v70 = base.F64_sub(base.F64_mul(v39, v51), v54)
							v73 = base.F64_add(base.F64_mul(base.F64_mul(v70, v70), v65), v44)
							*(*float64)(unsafe.Add(mBase, uint32(v21)+72)) = v73
							v77 = base.F64_add(base.F64_mul(base.F64_mul(v70, v61), v65), v49)
							*(*float64)(unsafe.Add(mBase, uint32(v21)+48)) = v77
							if base.F64_ne(base.F64_abs(v54), math.Float64frombits(uint64(0x7ff0000000000000))) != 0 {
								v82 = base.F64_abs(v73)
								v83 = math.Float64frombits(uint64(0x7ff0000000000000))
								v85 = base.F64_abs(v56)
								if base.B2i32(base.F64_eq(v82, v83)|base.F64_eq(v85, v83)|base.F64_eq(base.F64_abs(v67), v83) == int32(0))&base.F64_ne(base.F64_abs(v77), v83) != 0 {
									v208 = v77
									v209 = v73
									v210 = v67
									v216 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
									if v216 == int32(0) {
										v244 = int32(0)
									} else {
										v219 = *(*int32)(unsafe.Add(mBase, uint32(v216)))
										switch v219 - int32(429) {
										case 0:
											v244 = int32(1)
										case 1:
											v244 = int32(2)
										default:
											v244 = int32(0)
										}
									}
									if v244 != 0 {
										*(*float64)(unsafe.Add(mBase, uint32(v24)+64)) = v208
										*(*float64)(unsafe.Add(mBase, uint32(v24)+56)) = v210
										*(*float64)(unsafe.Add(mBase, uint32(v24)+48)) = v56
										*(*float64)(unsafe.Add(mBase, uint32(v24)+40)) = v209
										*(*float64)(unsafe.Add(mBase, uint32(v24)+32)) = v54
										*(*float64)(unsafe.Add(mBase, uint32(v24)+24)) = v51
										v275 = v24
										m.G0 = v21 + int32(96)
										return v275
									} else {
										*(*int32)(unsafe.Add(mBase, uint32(v21)+36)) = v21 + int32(48)
										*(*int32)(unsafe.Add(mBase, uint32(v21)+32)) = v21 + int32(56)
										*(*int32)(unsafe.Add(mBase, uint32(v21)+28)) = v21 - int32(-64)
										*(*int32)(unsafe.Add(mBase, uint32(v21)+24)) = v21 + int32(72)
										*(*int32)(unsafe.Add(mBase, uint32(v21)+20)) = v21 + int32(80)
										*(*int32)(unsafe.Add(mBase, uint32(v21)+16)) = v21 + int32(88)
										v273 = F_construct_array_builtin(m, v21+int32(16), int32(6), int32(701))
										mBase = m.M
										v274 = m.ExcPending
										if v274 != 0 {
											return int32(0)
										} else {
											v275 = v273
											m.G0 = v21 + int32(96)
											return v275
										}
									}
								} else {
									if base.F64_ne(v82, math.Float64frombits(uint64(0x7ff0000000000000))) != 0 {
										v111 = v85
										v114 = math.Float64frombits(uint64(0x7ff0000000000000))
										v115 = base.F64_ne(base.F64_abs(v41), v114)
										v118 = base.F64_eq(base.F64_abs(v46), v114)
										v121 = base.F64_abs(v67)
										if v115&base.B2i32(v118|base.F64_ne(v111, v114)&base.F64_ne(v121, v114) == int32(0)) != 0 {
											F_float_overflow_error(m)
											mBase = m.M
											v302 = m.ExcPending
											if v302 != 0 {
												return int32(0)
											} else {
												base.Wasm_trap_unreachable()
												for {
												}
											}
										} else {
											v130 = math.Float64frombits(uint64(0x7ff0000000000000))
											v132 = base.F64_abs(v77)
											if v115&base.B2i32(base.F64_eq(base.F64_abs(v42), v130)|base.F64_ne(v132, v130)|(v118|base.F64_eq(base.F64_abs(v39), v130)) == int32(0)) != 0 {
												F_float_overflow_error(m)
												mBase = m.M
												v302 = m.ExcPending
												if v302 != 0 {
													return int32(0)
												} else {
													base.Wasm_trap_unreachable()
													for {
													}
												}
											} else {
												if base.F64_eq(base.F64_abs(v73), math.Float64frombits(uint64(0x7ff0000000000000))) != 0 {
													*(*int64)(unsafe.Add(mBase, uint32(v21)+72)) = int64(9221120237041090560)
													v150 = math.Float64frombits(uint64(0x7ff8000000000000))
												} else {
													v150 = v73
												}
												if base.F64_eq(v121, math.Float64frombits(uint64(0x7ff0000000000000))) != 0 {
													*(*int64)(unsafe.Add(mBase, uint32(v21)+56)) = int64(9221120237041090560)
													v156 = math.Float64frombits(uint64(0x7ff8000000000000))
												} else {
													v156 = v67
												}
												if base.F64_eq(v132, math.Float64frombits(uint64(0x7ff0000000000000))) != 0 {
													v195 = v150
													v196 = v156
													v200 = v21 + int32(48)
													*(*int64)(unsafe.Add(mBase, uint32(v200))) = int64(9221120237041090560)
													v208 = math.Float64frombits(uint64(0x7ff8000000000000))
													v209 = v195
													v210 = v196
												} else {
													v208 = v77
													v209 = v150
													v210 = v156
												}
												v216 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
												if v216 == int32(0) {
													v244 = int32(0)
												} else {
													v219 = *(*int32)(unsafe.Add(mBase, uint32(v216)))
													switch v219 - int32(429) {
													case 0:
														v244 = int32(1)
													case 1:
														v244 = int32(2)
													default:
														v244 = int32(0)
													}
												}
												if v244 != 0 {
													*(*float64)(unsafe.Add(mBase, uint32(v24)+64)) = v208
													*(*float64)(unsafe.Add(mBase, uint32(v24)+56)) = v210
													*(*float64)(unsafe.Add(mBase, uint32(v24)+48)) = v56
													*(*float64)(unsafe.Add(mBase, uint32(v24)+40)) = v209
													*(*float64)(unsafe.Add(mBase, uint32(v24)+32)) = v54
													*(*float64)(unsafe.Add(mBase, uint32(v24)+24)) = v51
													v275 = v24
													m.G0 = v21 + int32(96)
													return v275
												} else {
													*(*int32)(unsafe.Add(mBase, uint32(v21)+36)) = v21 + int32(48)
													*(*int32)(unsafe.Add(mBase, uint32(v21)+32)) = v21 + int32(56)
													*(*int32)(unsafe.Add(mBase, uint32(v21)+28)) = v21 - int32(-64)
													*(*int32)(unsafe.Add(mBase, uint32(v21)+24)) = v21 + int32(72)
													*(*int32)(unsafe.Add(mBase, uint32(v21)+20)) = v21 + int32(80)
													*(*int32)(unsafe.Add(mBase, uint32(v21)+16)) = v21 + int32(88)
													v273 = F_construct_array_builtin(m, v21+int32(16), int32(6), int32(701))
													mBase = m.M
													v274 = m.ExcPending
													if v274 != 0 {
														return int32(0)
													} else {
														v275 = v273
														m.G0 = v21 + int32(96)
														return v275
													}
												}
											}
										}
									} else {
										v104 = math.Float64frombits(uint64(0x7ff0000000000000))
										if base.F64_ne(base.F64_abs(v42), v104)&base.F64_ne(base.F64_abs(v39), v104) != 0 {
											F_float_overflow_error(m)
											mBase = m.M
											v302 = m.ExcPending
											if v302 != 0 {
												return int32(0)
											} else {
												base.Wasm_trap_unreachable()
												for {
												}
											}
										} else {
											v111 = base.F64_abs(v56)
											v114 = math.Float64frombits(uint64(0x7ff0000000000000))
											v115 = base.F64_ne(base.F64_abs(v41), v114)
											v118 = base.F64_eq(base.F64_abs(v46), v114)
											v121 = base.F64_abs(v67)
											if v115&base.B2i32(v118|base.F64_ne(v111, v114)&base.F64_ne(v121, v114) == int32(0)) != 0 {
												F_float_overflow_error(m)
												mBase = m.M
												v302 = m.ExcPending
												if v302 != 0 {
													return int32(0)
												} else {
													base.Wasm_trap_unreachable()
													for {
													}
												}
											} else {
												v130 = math.Float64frombits(uint64(0x7ff0000000000000))
												v132 = base.F64_abs(v77)
												if v115&base.B2i32(base.F64_eq(base.F64_abs(v42), v130)|base.F64_ne(v132, v130)|(v118|base.F64_eq(base.F64_abs(v39), v130)) == int32(0)) != 0 {
													F_float_overflow_error(m)
													mBase = m.M
													v302 = m.ExcPending
													if v302 != 0 {
														return int32(0)
													} else {
														base.Wasm_trap_unreachable()
														for {
														}
													}
												} else {
													if base.F64_eq(base.F64_abs(v73), math.Float64frombits(uint64(0x7ff0000000000000))) != 0 {
														*(*int64)(unsafe.Add(mBase, uint32(v21)+72)) = int64(9221120237041090560)
														v150 = math.Float64frombits(uint64(0x7ff8000000000000))
													} else {
														v150 = v73
													}
													if base.F64_eq(v121, math.Float64frombits(uint64(0x7ff0000000000000))) != 0 {
														*(*int64)(unsafe.Add(mBase, uint32(v21)+56)) = int64(9221120237041090560)
														v156 = math.Float64frombits(uint64(0x7ff8000000000000))
													} else {
														v156 = v67
													}
													if base.F64_eq(v132, math.Float64frombits(uint64(0x7ff0000000000000))) != 0 {
														v195 = v150
														v196 = v156
														v200 = v21 + int32(48)
														*(*int64)(unsafe.Add(mBase, uint32(v200))) = int64(9221120237041090560)
														v208 = math.Float64frombits(uint64(0x7ff8000000000000))
														v209 = v195
														v210 = v196
													} else {
														v208 = v77
														v209 = v150
														v210 = v156
													}
													v216 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
													if v216 == int32(0) {
														v244 = int32(0)
													} else {
														v219 = *(*int32)(unsafe.Add(mBase, uint32(v216)))
														switch v219 - int32(429) {
														case 0:
															v244 = int32(1)
														case 1:
															v244 = int32(2)
														default:
															v244 = int32(0)
														}
													}
													if v244 != 0 {
														*(*float64)(unsafe.Add(mBase, uint32(v24)+64)) = v208
														*(*float64)(unsafe.Add(mBase, uint32(v24)+56)) = v210
														*(*float64)(unsafe.Add(mBase, uint32(v24)+48)) = v56
														*(*float64)(unsafe.Add(mBase, uint32(v24)+40)) = v209
														*(*float64)(unsafe.Add(mBase, uint32(v24)+32)) = v54
														*(*float64)(unsafe.Add(mBase, uint32(v24)+24)) = v51
														v275 = v24
														m.G0 = v21 + int32(96)
														return v275
													} else {
														*(*int32)(unsafe.Add(mBase, uint32(v21)+36)) = v21 + int32(48)
														*(*int32)(unsafe.Add(mBase, uint32(v21)+32)) = v21 + int32(56)
														*(*int32)(unsafe.Add(mBase, uint32(v21)+28)) = v21 - int32(-64)
														*(*int32)(unsafe.Add(mBase, uint32(v21)+24)) = v21 + int32(72)
														*(*int32)(unsafe.Add(mBase, uint32(v21)+20)) = v21 + int32(80)
														*(*int32)(unsafe.Add(mBase, uint32(v21)+16)) = v21 + int32(88)
														v273 = F_construct_array_builtin(m, v21+int32(16), int32(6), int32(701))
														mBase = m.M
														v274 = m.ExcPending
														if v274 != 0 {
															return int32(0)
														} else {
															v275 = v273
															m.G0 = v21 + int32(96)
															return v275
														}
													}
												}
											}
										}
									}
								}
							} else {
								v104 = math.Float64frombits(uint64(0x7ff0000000000000))
								if base.F64_ne(base.F64_abs(v42), v104)&base.F64_ne(base.F64_abs(v39), v104) != 0 {
									F_float_overflow_error(m)
									mBase = m.M
									v302 = m.ExcPending
									if v302 != 0 {
										return int32(0)
									} else {
										base.Wasm_trap_unreachable()
										for {
										}
									}
								} else {
									v111 = base.F64_abs(v56)
									v114 = math.Float64frombits(uint64(0x7ff0000000000000))
									v115 = base.F64_ne(base.F64_abs(v41), v114)
									v118 = base.F64_eq(base.F64_abs(v46), v114)
									v121 = base.F64_abs(v67)
									if v115&base.B2i32(v118|base.F64_ne(v111, v114)&base.F64_ne(v121, v114) == int32(0)) != 0 {
										F_float_overflow_error(m)
										mBase = m.M
										v302 = m.ExcPending
										if v302 != 0 {
											return int32(0)
										} else {
											base.Wasm_trap_unreachable()
											for {
											}
										}
									} else {
										v130 = math.Float64frombits(uint64(0x7ff0000000000000))
										v132 = base.F64_abs(v77)
										if v115&base.B2i32(base.F64_eq(base.F64_abs(v42), v130)|base.F64_ne(v132, v130)|(v118|base.F64_eq(base.F64_abs(v39), v130)) == int32(0)) != 0 {
											F_float_overflow_error(m)
											mBase = m.M
											v302 = m.ExcPending
											if v302 != 0 {
												return int32(0)
											} else {
												base.Wasm_trap_unreachable()
												for {
												}
											}
										} else {
											if base.F64_eq(base.F64_abs(v73), math.Float64frombits(uint64(0x7ff0000000000000))) != 0 {
												*(*int64)(unsafe.Add(mBase, uint32(v21)+72)) = int64(9221120237041090560)
												v150 = math.Float64frombits(uint64(0x7ff8000000000000))
											} else {
												v150 = v73
											}
											if base.F64_eq(v121, math.Float64frombits(uint64(0x7ff0000000000000))) != 0 {
												*(*int64)(unsafe.Add(mBase, uint32(v21)+56)) = int64(9221120237041090560)
												v156 = math.Float64frombits(uint64(0x7ff8000000000000))
											} else {
												v156 = v67
											}
											if base.F64_eq(v132, math.Float64frombits(uint64(0x7ff0000000000000))) != 0 {
												v195 = v150
												v196 = v156
												v200 = v21 + int32(48)
												*(*int64)(unsafe.Add(mBase, uint32(v200))) = int64(9221120237041090560)
												v208 = math.Float64frombits(uint64(0x7ff8000000000000))
												v209 = v195
												v210 = v196
											} else {
												v208 = v77
												v209 = v150
												v210 = v156
											}
											v216 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
											if v216 == int32(0) {
												v244 = int32(0)
											} else {
												v219 = *(*int32)(unsafe.Add(mBase, uint32(v216)))
												switch v219 - int32(429) {
												case 0:
													v244 = int32(1)
												case 1:
													v244 = int32(2)
												default:
													v244 = int32(0)
												}
											}
											if v244 != 0 {
												*(*float64)(unsafe.Add(mBase, uint32(v24)+64)) = v208
												*(*float64)(unsafe.Add(mBase, uint32(v24)+56)) = v210
												*(*float64)(unsafe.Add(mBase, uint32(v24)+48)) = v56
												*(*float64)(unsafe.Add(mBase, uint32(v24)+40)) = v209
												*(*float64)(unsafe.Add(mBase, uint32(v24)+32)) = v54
												*(*float64)(unsafe.Add(mBase, uint32(v24)+24)) = v51
												v275 = v24
												m.G0 = v21 + int32(96)
												return v275
											} else {
												*(*int32)(unsafe.Add(mBase, uint32(v21)+36)) = v21 + int32(48)
												*(*int32)(unsafe.Add(mBase, uint32(v21)+32)) = v21 + int32(56)
												*(*int32)(unsafe.Add(mBase, uint32(v21)+28)) = v21 - int32(-64)
												*(*int32)(unsafe.Add(mBase, uint32(v21)+24)) = v21 + int32(72)
												*(*int32)(unsafe.Add(mBase, uint32(v21)+20)) = v21 + int32(80)
												*(*int32)(unsafe.Add(mBase, uint32(v21)+16)) = v21 + int32(88)
												v273 = F_construct_array_builtin(m, v21+int32(16), int32(6), int32(701))
												mBase = m.M
												v274 = m.ExcPending
												if v274 != 0 {
													return int32(0)
												} else {
													v275 = v273
													m.G0 = v21 + int32(96)
													return v275
												}
											}
										}
									}
								}
							}
						} else {
							v161 = base.F64_abs(v39)
							if base.B2i32(base.Ui64(base.I64_reinterpret_f64(v161)) <= base.Ui64(int64(9218868437227405312)))&base.F64_ne(v161, math.Float64frombits(uint64(0x7ff0000000000000))) == int32(0) {
								v170 = int64(9221120237041090560)
								*(*int64)(unsafe.Add(mBase, uint32(v21)+72)) = v170
								*(*int64)(unsafe.Add(mBase, uint32(v21)+48)) = v170
								v174 = math.Float64frombits(uint64(0x7ff8000000000000))
								v176 = v174
								v177 = v174
							} else {
								v176 = v49
								v177 = v44
							}
							v178 = base.F64_abs(v41)
							if base.B2i32(base.Ui64(base.I64_reinterpret_f64(v178)) <= base.Ui64(int64(9218868437227405312)))&base.F64_ne(v178, math.Float64frombits(uint64(0x7ff0000000000000))) != 0 {
								v208 = v176
								v209 = v177
								v210 = v47
							} else {
								*(*int64)(unsafe.Add(mBase, uint32(v21)+48)) = int64(9221120237041090560)
								v195 = v177
								v196 = math.Float64frombits(uint64(0x7ff8000000000000))
								v200 = v21 + int32(56)
								*(*int64)(unsafe.Add(mBase, uint32(v200))) = int64(9221120237041090560)
								v208 = math.Float64frombits(uint64(0x7ff8000000000000))
								v209 = v195
								v210 = v196
							}
							v216 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
							if v216 == int32(0) {
								v244 = int32(0)
							} else {
								v219 = *(*int32)(unsafe.Add(mBase, uint32(v216)))
								switch v219 - int32(429) {
								case 0:
									v244 = int32(1)
								case 1:
									v244 = int32(2)
								default:
									v244 = int32(0)
								}
							}
							if v244 != 0 {
								*(*float64)(unsafe.Add(mBase, uint32(v24)+64)) = v208
								*(*float64)(unsafe.Add(mBase, uint32(v24)+56)) = v210
								*(*float64)(unsafe.Add(mBase, uint32(v24)+48)) = v56
								*(*float64)(unsafe.Add(mBase, uint32(v24)+40)) = v209
								*(*float64)(unsafe.Add(mBase, uint32(v24)+32)) = v54
								*(*float64)(unsafe.Add(mBase, uint32(v24)+24)) = v51
								v275 = v24
								m.G0 = v21 + int32(96)
								return v275
							} else {
								*(*int32)(unsafe.Add(mBase, uint32(v21)+36)) = v21 + int32(48)
								*(*int32)(unsafe.Add(mBase, uint32(v21)+32)) = v21 + int32(56)
								*(*int32)(unsafe.Add(mBase, uint32(v21)+28)) = v21 - int32(-64)
								*(*int32)(unsafe.Add(mBase, uint32(v21)+24)) = v21 + int32(72)
								*(*int32)(unsafe.Add(mBase, uint32(v21)+20)) = v21 + int32(80)
								*(*int32)(unsafe.Add(mBase, uint32(v21)+16)) = v21 + int32(88)
								v273 = F_construct_array_builtin(m, v21+int32(16), int32(6), int32(701))
								mBase = m.M
								v274 = m.ExcPending
								if v274 != 0 {
									return int32(0)
								} else {
									v275 = v273
									m.G0 = v21 + int32(96)
									return v275
								}
							}
						}
					}
				}
			}
		}
	}
}
func F_float8_regr_avgy(m *base.Module, l0 int32) int32 {
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
				*(*int32)(unsafe.Add(mBase, uint32(v7)+4)) = int32(6)
				*(*int32)(unsafe.Add(mBase, uint32(v7))) = int32(_a_F_float8_regr_avgy_0)
				F_errmsg_internal(m, int32(_a_F_float8_regr_avgy_1), v7)
				mBase = m.M
				v49 = m.ExcPending
				if v49 != 0 {
					return int32(0)
				} else {
					F_errfinish(m, int32(_a_F_float8_regr_avgy_2), int32(2938), int32(_a_F_float8_regr_avgy_3))
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
			if v17 != int32(6) {
				F_errstart_cold(m, int32(21), int32(0))
				mBase = m.M
				v42 = m.ExcPending
				if v42 != 0 {
					return int32(0)
				} else {
					*(*int32)(unsafe.Add(mBase, uint32(v7)+4)) = int32(6)
					*(*int32)(unsafe.Add(mBase, uint32(v7))) = int32(_a_F_float8_regr_avgy_0)
					F_errmsg_internal(m, int32(_a_F_float8_regr_avgy_1), v7)
					mBase = m.M
					v49 = m.ExcPending
					if v49 != 0 {
						return int32(0)
					} else {
						F_errfinish(m, int32(_a_F_float8_regr_avgy_2), int32(2938), int32(_a_F_float8_regr_avgy_3))
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
						*(*int32)(unsafe.Add(mBase, uint32(v7)+4)) = int32(6)
						*(*int32)(unsafe.Add(mBase, uint32(v7))) = int32(_a_F_float8_regr_avgy_0)
						F_errmsg_internal(m, int32(_a_F_float8_regr_avgy_1), v7)
						mBase = m.M
						v49 = m.ExcPending
						if v49 != 0 {
							return int32(0)
						} else {
							F_errfinish(m, int32(_a_F_float8_regr_avgy_2), int32(2938), int32(_a_F_float8_regr_avgy_3))
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
							*(*int32)(unsafe.Add(mBase, uint32(v7)+4)) = int32(6)
							*(*int32)(unsafe.Add(mBase, uint32(v7))) = int32(_a_F_float8_regr_avgy_0)
							F_errmsg_internal(m, int32(_a_F_float8_regr_avgy_1), v7)
							mBase = m.M
							v49 = m.ExcPending
							if v49 != 0 {
								return int32(0)
							} else {
								F_errfinish(m, int32(_a_F_float8_regr_avgy_2), int32(2938), int32(_a_F_float8_regr_avgy_3))
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
						if base.F64_lt(v24, float64(1)) != 0 {
							v27 = int32(1)
							*(*uint8)(unsafe.Add(mBase, uint32(l0)+16)) = uint8(v27)
							v34 = int32(0)
							m.G0 = v7 + int32(16)
							return v34
						} else {
							v30 = *(*float64)(unsafe.Add(mBase, uint32(v10)+48))
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
func F_float8_regr_slope(m *base.Module, l0 int32) int32 {
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
	var v33 int32
	_ = v33
	var v36 float64
	_ = v36
	var v38 int32
	_ = v38
	var v39 int32
	_ = v39
	var v41 int32
	_ = v41
	var v49 int32
	_ = v49
	var v56 int32
	_ = v56
	var v61 int32
	_ = v61
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
			v49 = m.ExcPending
			if v49 != 0 {
				return int32(0)
			} else {
				*(*int32)(unsafe.Add(mBase, uint32(v7)+4)) = int32(6)
				*(*int32)(unsafe.Add(mBase, uint32(v7))) = int32(_a_F_float8_regr_slope_0)
				F_errmsg_internal(m, int32(_a_F_float8_regr_slope_1), v7)
				mBase = m.M
				v56 = m.ExcPending
				if v56 != 0 {
					return int32(0)
				} else {
					F_errfinish(m, int32(_a_F_float8_regr_slope_2), int32(2938), int32(_a_F_float8_regr_slope_3))
					mBase = m.M
					v61 = m.ExcPending
					if v61 != 0 {
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
				v49 = m.ExcPending
				if v49 != 0 {
					return int32(0)
				} else {
					*(*int32)(unsafe.Add(mBase, uint32(v7)+4)) = int32(6)
					*(*int32)(unsafe.Add(mBase, uint32(v7))) = int32(_a_F_float8_regr_slope_0)
					F_errmsg_internal(m, int32(_a_F_float8_regr_slope_1), v7)
					mBase = m.M
					v56 = m.ExcPending
					if v56 != 0 {
						return int32(0)
					} else {
						F_errfinish(m, int32(_a_F_float8_regr_slope_2), int32(2938), int32(_a_F_float8_regr_slope_3))
						mBase = m.M
						v61 = m.ExcPending
						if v61 != 0 {
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
					v49 = m.ExcPending
					if v49 != 0 {
						return int32(0)
					} else {
						*(*int32)(unsafe.Add(mBase, uint32(v7)+4)) = int32(6)
						*(*int32)(unsafe.Add(mBase, uint32(v7))) = int32(_a_F_float8_regr_slope_0)
						F_errmsg_internal(m, int32(_a_F_float8_regr_slope_1), v7)
						mBase = m.M
						v56 = m.ExcPending
						if v56 != 0 {
							return int32(0)
						} else {
							F_errfinish(m, int32(_a_F_float8_regr_slope_2), int32(2938), int32(_a_F_float8_regr_slope_3))
							mBase = m.M
							v61 = m.ExcPending
							if v61 != 0 {
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
						v49 = m.ExcPending
						if v49 != 0 {
							return int32(0)
						} else {
							*(*int32)(unsafe.Add(mBase, uint32(v7)+4)) = int32(6)
							*(*int32)(unsafe.Add(mBase, uint32(v7))) = int32(_a_F_float8_regr_slope_0)
							F_errmsg_internal(m, int32(_a_F_float8_regr_slope_1), v7)
							mBase = m.M
							v56 = m.ExcPending
							if v56 != 0 {
								return int32(0)
							} else {
								F_errfinish(m, int32(_a_F_float8_regr_slope_2), int32(2938), int32(_a_F_float8_regr_slope_3))
								mBase = m.M
								v61 = m.ExcPending
								if v61 != 0 {
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
							v41 = int32(0)
							m.G0 = v7 + int32(16)
							return v41
						} else {
							v30 = *(*float64)(unsafe.Add(mBase, uint32(v10)+40))
							if base.F64_eq(v30, float64(0)) != 0 {
								v33 = int32(1)
								*(*uint8)(unsafe.Add(mBase, uint32(l0)+16)) = uint8(v33)
								v41 = int32(0)
								m.G0 = v7 + int32(16)
								return v41
							} else {
								v36 = *(*float64)(unsafe.Add(mBase, uint32(v10)+64))
								v38 = F_Float8GetDatum(m, base.F64_div(v36, v30))
								mBase = m.M
								v39 = m.ExcPending
								if v39 != 0 {
									return int32(0)
								} else {
									v41 = v38
									m.G0 = v7 + int32(16)
									return v41
								}
							}
						}
					}
				}
			}
		}
	}
}
