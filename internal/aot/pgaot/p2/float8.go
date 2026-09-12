package p2

import (
	base "github.com/shibukawa/pgmem/internal/aot/pgaot/base"
	"math"
	"unsafe"
)

func F_float8_lerp(m *base.Module, l0 int32, l1 int32, l2 float64) int32 {
	mBase := m.M
	_ = mBase
	var v4 float64
	_ = v4
	var v5 float64
	_ = v5
	var v9 int32
	_ = v9
	var v12 int32
	_ = v12
	v4 = *(*float64)(unsafe.Add(mBase, uint32(l1)))
	v5 = *(*float64)(unsafe.Add(mBase, uint32(l0)))
	v9 = F_Float8GetDatum(m, base.F64_add(base.F64_mul(l2, base.F64_sub(v4, v5)), v5))
	mBase = m.M
	v12 = m.ExcPending
	if v12 != 0 {
		return int32(0)
	} else {
		return v9
	}
}
func F_float8_numeric(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v5 int32
	_ = v5
	var v7 int32
	_ = v7
	var v9 int32
	_ = v9
	var v10 float64
	_ = v10
	var v18 int32
	_ = v18
	var v21 int32
	_ = v21
	var v29 int32
	_ = v29
	var v30 int32
	_ = v30
	var v33 int32
	_ = v33
	var v34 int32
	_ = v34
	var v42 int32
	_ = v42
	var v43 int32
	_ = v43
	var v45 int32
	_ = v45
	var v46 int64
	_ = v46
	var v53 int32
	_ = v53
	var v61 int32
	_ = v61
	var v62 int32
	_ = v62
	var v66 int32
	_ = v66
	var v67 int32
	_ = v67
	var v68 int32
	_ = v68
	var v72 int32
	_ = v72
	var v73 int32
	_ = v73
	v5 = m.G0
	v7 = v5 - int32(176)
	m.G0 = v7
	v9 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v10 = *(*float64)(unsafe.Add(mBase, uint32(v9)))
	if base.Ui64(int64(9218868437227405313)) <= base.Ui64(base.I64_reinterpret_f64(v10)&int64(9223372036854775807)) {
		v18 = F_make_result_opt_error(m, int32(1737708), int32(0))
		mBase = m.M
		v21 = m.ExcPending
		if v21 != 0 {
			return int32(0)
		} else {
			v73 = v18
			m.G0 = v7 + int32(176)
			return v73
		}
	} else {
		if base.F64_eq(base.F64_abs(v10), math.Float64frombits(uint64(0x7ff0000000000000))) != 0 {
			if base.F64_lt(v10, float64(0)) != 0 {
				v29 = F_make_result_opt_error(m, int32(1737756), int32(0))
				mBase = m.M
				v30 = m.ExcPending
				if v30 != 0 {
					return int32(0)
				} else {
					v73 = v29
					m.G0 = v7 + int32(176)
					return v73
				}
			} else {
				v33 = F_make_result_opt_error(m, int32(1737732), int32(0))
				mBase = m.M
				v34 = m.ExcPending
				if v34 != 0 {
					return int32(0)
				} else {
					v73 = v33
					m.G0 = v7 + int32(176)
					return v73
				}
			}
		} else {
			*(*float64)(unsafe.Add(mBase, uint32(v7)+8)) = v10
			*(*int32)(unsafe.Add(mBase, uint32(v7))) = int32(15)
			v42 = F_pg_snprintf(m, v7+int32(32), int32(115), int32(337978), v7)
			mBase = m.M
			v43 = m.ExcPending
			if v43 != 0 {
				return int32(0)
			} else {
				v45 = v7 + int32(168)
				v46 = int64(0)
				*(*int64)(unsafe.Add(mBase, uint32(v45))) = v46
				*(*int64)(unsafe.Add(mBase, uint32(v7)+160)) = v46
				*(*int64)(unsafe.Add(mBase, uint32(v7)+152)) = v46
				v53 = v7 + int32(32)
				v61 = F_set_var_from_str(m, v53, v53, v7+int32(152), v7+int32(28), int32(0))
				mBase = m.M
				v62 = m.ExcPending
				if v62 != 0 {
					return int32(0)
				} else {
					v66 = F_make_result_opt_error(m, v7+int32(152), int32(0))
					mBase = m.M
					v67 = m.ExcPending
					if v67 != 0 {
						return int32(0)
					} else {
						v68 = *(*int32)(unsafe.Add(mBase, uint32(v45)))
						if v68 == int32(0) {
							v73 = v66
							m.G0 = v7 + int32(176)
							return v73
						} else {
							F_pfree(m, v68)
							mBase = m.M
							v72 = m.ExcPending
							if v72 != 0 {
								return int32(0)
							} else {
								v73 = v66
								m.G0 = v7 + int32(176)
								return v73
							}
						}
					}
				}
			}
		}
	}
}
func F_float8_regr_intercept(m *base.Module, l0 int32) int32 {
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
	var v34 int32
	_ = v34
	var v37 float64
	_ = v37
	var v38 float64
	_ = v38
	var v41 float64
	_ = v41
	var v46 int32
	_ = v46
	var v47 int32
	_ = v47
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
				*(*int32)(unsafe.Add(mBase, uint32(v8))) = int32(84140)
				F_errmsg_internal(m, int32(26162), v8)
				mBase = m.M
				v64 = m.ExcPending
				if v64 != 0 {
					return int32(0)
				} else {
					F_errfinish(m, int32(493466), int32(2938), int32(24702))
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
					*(*int32)(unsafe.Add(mBase, uint32(v8))) = int32(84140)
					F_errmsg_internal(m, int32(26162), v8)
					mBase = m.M
					v64 = m.ExcPending
					if v64 != 0 {
						return int32(0)
					} else {
						F_errfinish(m, int32(493466), int32(2938), int32(24702))
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
						*(*int32)(unsafe.Add(mBase, uint32(v8))) = int32(84140)
						F_errmsg_internal(m, int32(26162), v8)
						mBase = m.M
						v64 = m.ExcPending
						if v64 != 0 {
							return int32(0)
						} else {
							F_errfinish(m, int32(493466), int32(2938), int32(24702))
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
							*(*int32)(unsafe.Add(mBase, uint32(v8))) = int32(84140)
							F_errmsg_internal(m, int32(26162), v8)
							mBase = m.M
							v64 = m.ExcPending
							if v64 != 0 {
								return int32(0)
							} else {
								F_errfinish(m, int32(493466), int32(2938), int32(24702))
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
							if base.F64_eq(v31, float64(0)) != 0 {
								v34 = int32(1)
								*(*uint8)(unsafe.Add(mBase, uint32(l0)+16)) = uint8(v34)
								v49 = int32(0)
								m.G0 = v8 + int32(16)
								return v49
							} else {
								v37 = *(*float64)(unsafe.Add(mBase, uint32(v11)+48))
								v38 = *(*float64)(unsafe.Add(mBase, uint32(v11)+32))
								v41 = *(*float64)(unsafe.Add(mBase, uint32(v11-int32(-64))))
								v46 = F_Float8GetDatum(m, base.F64_div(base.F64_sub(v37, base.F64_div(base.F64_mul(v38, v41), v31)), v25))
								mBase = m.M
								v47 = m.ExcPending
								if v47 != 0 {
									return int32(0)
								} else {
									v49 = v46
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
}
func F_float8_regr_r2(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v7 int32
	_ = v7
	var v9 int32
	_ = v9
	var v11 int32
	_ = v11
	var v12 int32
	_ = v12
	var v15 int32
	_ = v15
	var v16 int32
	_ = v16
	var v19 int32
	_ = v19
	var v22 int32
	_ = v22
	var v23 int32
	_ = v23
	var v26 float64
	_ = v26
	var v29 int32
	_ = v29
	var v32 float64
	_ = v32
	var v35 int32
	_ = v35
	var v38 float64
	_ = v38
	var v42 int32
	_ = v42
	var v43 int32
	_ = v43
	var v46 float64
	_ = v46
	var v50 int32
	_ = v50
	var v51 int32
	_ = v51
	var v55 int32
	_ = v55
	var v63 int32
	_ = v63
	var v70 int32
	_ = v70
	var v75 int32
	_ = v75
	v7 = m.G0
	v9 = v7 - int32(16)
	m.G0 = v9
	v11 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v12 = F_pg_detoast_datum(m, v11)
	mBase = m.M
	v15 = m.ExcPending
	if v15 != 0 {
		return int32(0)
	} else {
		v16 = *(*int32)(unsafe.Add(mBase, uint32(v12)+4))
		if v16 != int32(1) {
			F_errstart_cold(m, int32(21), int32(0))
			mBase = m.M
			v63 = m.ExcPending
			if v63 != 0 {
				return int32(0)
			} else {
				*(*int32)(unsafe.Add(mBase, uint32(v9)+4)) = int32(6)
				*(*int32)(unsafe.Add(mBase, uint32(v9))) = int32(560419)
				F_errmsg_internal(m, int32(26162), v9)
				mBase = m.M
				v70 = m.ExcPending
				if v70 != 0 {
					return int32(0)
				} else {
					F_errfinish(m, int32(493466), int32(2938), int32(24702))
					mBase = m.M
					v75 = m.ExcPending
					if v75 != 0 {
						return int32(0)
					} else {
						base.Wasm_trap_unreachable()
						for {
						}
					}
				}
			}
		} else {
			v19 = *(*int32)(unsafe.Add(mBase, uint32(v12)+16))
			if v19 != int32(6) {
				F_errstart_cold(m, int32(21), int32(0))
				mBase = m.M
				v63 = m.ExcPending
				if v63 != 0 {
					return int32(0)
				} else {
					*(*int32)(unsafe.Add(mBase, uint32(v9)+4)) = int32(6)
					*(*int32)(unsafe.Add(mBase, uint32(v9))) = int32(560419)
					F_errmsg_internal(m, int32(26162), v9)
					mBase = m.M
					v70 = m.ExcPending
					if v70 != 0 {
						return int32(0)
					} else {
						F_errfinish(m, int32(493466), int32(2938), int32(24702))
						mBase = m.M
						v75 = m.ExcPending
						if v75 != 0 {
							return int32(0)
						} else {
							base.Wasm_trap_unreachable()
							for {
							}
						}
					}
				}
			} else {
				v22 = *(*int32)(unsafe.Add(mBase, uint32(v12)+8))
				if v22 != 0 {
					F_errstart_cold(m, int32(21), int32(0))
					mBase = m.M
					v63 = m.ExcPending
					if v63 != 0 {
						return int32(0)
					} else {
						*(*int32)(unsafe.Add(mBase, uint32(v9)+4)) = int32(6)
						*(*int32)(unsafe.Add(mBase, uint32(v9))) = int32(560419)
						F_errmsg_internal(m, int32(26162), v9)
						mBase = m.M
						v70 = m.ExcPending
						if v70 != 0 {
							return int32(0)
						} else {
							F_errfinish(m, int32(493466), int32(2938), int32(24702))
							mBase = m.M
							v75 = m.ExcPending
							if v75 != 0 {
								return int32(0)
							} else {
								base.Wasm_trap_unreachable()
								for {
								}
							}
						}
					}
				} else {
					v23 = *(*int32)(unsafe.Add(mBase, uint32(v12)+12))
					if v23 != int32(701) {
						F_errstart_cold(m, int32(21), int32(0))
						mBase = m.M
						v63 = m.ExcPending
						if v63 != 0 {
							return int32(0)
						} else {
							*(*int32)(unsafe.Add(mBase, uint32(v9)+4)) = int32(6)
							*(*int32)(unsafe.Add(mBase, uint32(v9))) = int32(560419)
							F_errmsg_internal(m, int32(26162), v9)
							mBase = m.M
							v70 = m.ExcPending
							if v70 != 0 {
								return int32(0)
							} else {
								F_errfinish(m, int32(493466), int32(2938), int32(24702))
								mBase = m.M
								v75 = m.ExcPending
								if v75 != 0 {
									return int32(0)
								} else {
									base.Wasm_trap_unreachable()
									for {
									}
								}
							}
						}
					} else {
						v26 = *(*float64)(unsafe.Add(mBase, uint32(v12)+24))
						if base.F64_lt(v26, float64(1)) != 0 {
							v29 = int32(1)
							*(*uint8)(unsafe.Add(mBase, uint32(l0)+16)) = uint8(v29)
							v55 = int32(0)
							m.G0 = v9 + int32(16)
							return v55
						} else {
							v32 = *(*float64)(unsafe.Add(mBase, uint32(v12)+40))
							if base.F64_eq(v32, float64(0)) != 0 {
								v35 = int32(1)
								*(*uint8)(unsafe.Add(mBase, uint32(l0)+16)) = uint8(v35)
								v55 = int32(0)
								m.G0 = v9 + int32(16)
								return v55
							} else {
								v38 = *(*float64)(unsafe.Add(mBase, uint32(v12)+56))
								if base.F64_eq(v38, float64(0)) != 0 {
									v42 = F_Float8GetDatum(m, float64(1))
									mBase = m.M
									v43 = m.ExcPending
									if v43 != 0 {
										return int32(0)
									} else {
										v55 = v42
										m.G0 = v9 + int32(16)
										return v55
									}
								} else {
									v46 = *(*float64)(unsafe.Add(mBase, uint32(v12-int32(-64))))
									v50 = F_Float8GetDatum(m, base.F64_div(base.F64_mul(v46, v46), base.F64_mul(v32, v38)))
									mBase = m.M
									v51 = m.ExcPending
									if v51 != 0 {
										return int32(0)
									} else {
										v55 = v50
										m.G0 = v9 + int32(16)
										return v55
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
func F_float8_var_samp(m *base.Module, l0 int32) int32 {
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
				*(*int32)(unsafe.Add(mBase, uint32(v7)+4)) = int32(3)
				*(*int32)(unsafe.Add(mBase, uint32(v7))) = int32(237380)
				F_errmsg_internal(m, int32(26162), v7)
				mBase = m.M
				v51 = m.ExcPending
				if v51 != 0 {
					return int32(0)
				} else {
					F_errfinish(m, int32(493466), int32(2938), int32(24702))
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
			if v17 != int32(3) {
				F_errstart_cold(m, int32(21), int32(0))
				mBase = m.M
				v44 = m.ExcPending
				if v44 != 0 {
					return int32(0)
				} else {
					*(*int32)(unsafe.Add(mBase, uint32(v7)+4)) = int32(3)
					*(*int32)(unsafe.Add(mBase, uint32(v7))) = int32(237380)
					F_errmsg_internal(m, int32(26162), v7)
					mBase = m.M
					v51 = m.ExcPending
					if v51 != 0 {
						return int32(0)
					} else {
						F_errfinish(m, int32(493466), int32(2938), int32(24702))
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
						*(*int32)(unsafe.Add(mBase, uint32(v7)+4)) = int32(3)
						*(*int32)(unsafe.Add(mBase, uint32(v7))) = int32(237380)
						F_errmsg_internal(m, int32(26162), v7)
						mBase = m.M
						v51 = m.ExcPending
						if v51 != 0 {
							return int32(0)
						} else {
							F_errfinish(m, int32(493466), int32(2938), int32(24702))
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
							*(*int32)(unsafe.Add(mBase, uint32(v7)+4)) = int32(3)
							*(*int32)(unsafe.Add(mBase, uint32(v7))) = int32(237380)
							F_errmsg_internal(m, int32(26162), v7)
							mBase = m.M
							v51 = m.ExcPending
							if v51 != 0 {
								return int32(0)
							} else {
								F_errfinish(m, int32(493466), int32(2938), int32(24702))
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
						if base.F64_le(v24, float64(1)) != 0 {
							v27 = int32(1)
							*(*uint8)(unsafe.Add(mBase, uint32(l0)+16)) = uint8(v27)
							v36 = int32(0)
							m.G0 = v7 + int32(16)
							return v36
						} else {
							v30 = *(*float64)(unsafe.Add(mBase, uint32(v10)+40))
							v34 = F_Float8GetDatum(m, base.F64_div(v30, base.F64_add(v24, float64(-1))))
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
