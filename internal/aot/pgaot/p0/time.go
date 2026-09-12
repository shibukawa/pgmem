package p0

import (
	base "github.com/shibukawa/pgmem/internal/aot/pgaot/base"
	"unsafe"
)

func F_AdjustTimeForTypmod(m *base.Module, l0 int32, l1 int32) {
	mBase := m.M
	_ = mBase
	var v9 int32
	_ = v9
	var v12 int64
	_ = v12
	var v15 int64
	_ = v15
	var v16 int64
	_ = v16
	var v19 int64
	_ = v19
	var v20 int64
	_ = v20
	var v22 int64
	_ = v22
	var v23 int64
	_ = v23
	var v26 int64
	_ = v26
	if base.Ui32(l1) <= base.Ui32(int32(6)) {
		v9 = l1 << (uint(int32(3)) % 32)
		v12 = *(*int64)(unsafe.Add(mBase, uint32(v9)+uint32(_consts[771])))
		v15 = *(*int64)(unsafe.Add(mBase, uint32(v9)+uint32(_consts[772])))
		v16 = *(*int64)(unsafe.Add(mBase, uint32(l0)))
		if int64(0) <= v16 {
			v19 = v15 + v16
			v20 = base.I64_rem_s(v19, v12)
			v26 = v19 - v20
		} else {
			v22 = v15 - v16
			v23 = base.I64_rem_s(v22, v12)
			v26 = v23 - v22
		}
		*(*int64)(unsafe.Add(mBase, uint32(l0))) = v26
	} else {
	}
	return
}
func F_make_time(m *base.Module, l0 int32) int32 {
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
	var v13 float64
	_ = v13
	var v14 int32
	_ = v14
	var v26 float64
	_ = v26
	var v34 int64
	_ = v34
	var v36 int64
	_ = v36
	var v37 int32
	_ = v37
	var v45 int64
	_ = v45
	var v55 int32
	_ = v55
	var v58 int32
	_ = v58
	var v64 int32
	_ = v64
	var v69 int32
	_ = v69
	var v70 int32
	_ = v70
	var v71 int32
	_ = v71
	v7 = m.G0
	v9 = v7 - int32(16)
	m.G0 = v9
	v11 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
	v12 = *(*int32)(unsafe.Add(mBase, uint32(l0)+36))
	v13 = *(*float64)(unsafe.Add(mBase, uint32(v12)))
	v14 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	if base.Ui32(int32(24)) < base.Ui32(v14) {
		F_errstart_cold(m, int32(21), int32(0))
		mBase = m.M
		v55 = m.ExcPending
		if v55 != 0 {
			return int32(0)
		} else {
			F_errcode(m, int32(134217858))
			mBase = m.M
			v58 = m.ExcPending
			if v58 != 0 {
				return int32(0)
			} else {
				*(*float64)(unsafe.Add(mBase, uint32(v9)+8)) = v13
				*(*int32)(unsafe.Add(mBase, uint32(v9)+4)) = v11
				*(*int32)(unsafe.Add(mBase, uint32(v9))) = v14
				F_errmsg(m, int32(356120), v9)
				mBase = m.M
				v64 = m.ExcPending
				if v64 != 0 {
					return int32(0)
				} else {
					F_errfinish(m, int32(524949), int32(1654), int32(395199))
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
		}
	} else {
		if base.Ui32(int32(59)) < base.Ui32(v11) {
			F_errstart_cold(m, int32(21), int32(0))
			mBase = m.M
			v55 = m.ExcPending
			if v55 != 0 {
				return int32(0)
			} else {
				F_errcode(m, int32(134217858))
				mBase = m.M
				v58 = m.ExcPending
				if v58 != 0 {
					return int32(0)
				} else {
					*(*float64)(unsafe.Add(mBase, uint32(v9)+8)) = v13
					*(*int32)(unsafe.Add(mBase, uint32(v9)+4)) = v11
					*(*int32)(unsafe.Add(mBase, uint32(v9))) = v14
					F_errmsg(m, int32(356120), v9)
					mBase = m.M
					v64 = m.ExcPending
					if v64 != 0 {
						return int32(0)
					} else {
						F_errfinish(m, int32(524949), int32(1654), int32(395199))
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
			}
		} else {
			if base.Ui64(int64(9218868437227405312)) < base.Ui64(base.I64_reinterpret_f64(v13)&int64(9223372036854775807)) {
				F_errstart_cold(m, int32(21), int32(0))
				mBase = m.M
				v55 = m.ExcPending
				if v55 != 0 {
					return int32(0)
				} else {
					F_errcode(m, int32(134217858))
					mBase = m.M
					v58 = m.ExcPending
					if v58 != 0 {
						return int32(0)
					} else {
						*(*float64)(unsafe.Add(mBase, uint32(v9)+8)) = v13
						*(*int32)(unsafe.Add(mBase, uint32(v9)+4)) = v11
						*(*int32)(unsafe.Add(mBase, uint32(v9))) = v14
						F_errmsg(m, int32(356120), v9)
						mBase = m.M
						v64 = m.ExcPending
						if v64 != 0 {
							return int32(0)
						} else {
							F_errfinish(m, int32(524949), int32(1654), int32(395199))
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
				}
			} else {
				v26 = base.F64_nearest(base.F64_mul(v13, float64(1e+06)))
				if base.F64_lt(v26, float64(0)) != 0 {
					F_errstart_cold(m, int32(21), int32(0))
					mBase = m.M
					v55 = m.ExcPending
					if v55 != 0 {
						return int32(0)
					} else {
						F_errcode(m, int32(134217858))
						mBase = m.M
						v58 = m.ExcPending
						if v58 != 0 {
							return int32(0)
						} else {
							*(*float64)(unsafe.Add(mBase, uint32(v9)+8)) = v13
							*(*int32)(unsafe.Add(mBase, uint32(v9)+4)) = v11
							*(*int32)(unsafe.Add(mBase, uint32(v9))) = v14
							F_errmsg(m, int32(356120), v9)
							mBase = m.M
							v64 = m.ExcPending
							if v64 != 0 {
								return int32(0)
							} else {
								F_errfinish(m, int32(524949), int32(1654), int32(395199))
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
					}
				} else {
					if base.F64_gt(v26, float64(6e+07)) != 0 {
						F_errstart_cold(m, int32(21), int32(0))
						mBase = m.M
						v55 = m.ExcPending
						if v55 != 0 {
							return int32(0)
						} else {
							F_errcode(m, int32(134217858))
							mBase = m.M
							v58 = m.ExcPending
							if v58 != 0 {
								return int32(0)
							} else {
								*(*float64)(unsafe.Add(mBase, uint32(v9)+8)) = v13
								*(*int32)(unsafe.Add(mBase, uint32(v9)+4)) = v11
								*(*int32)(unsafe.Add(mBase, uint32(v9))) = v14
								F_errmsg(m, int32(356120), v9)
								mBase = m.M
								v64 = m.ExcPending
								if v64 != 0 {
									return int32(0)
								} else {
									F_errfinish(m, int32(524949), int32(1654), int32(395199))
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
						}
					} else {
						if base.F64_lt(base.F64_abs(v26), float64(9.223372036854776e+18)) != 0 {
							v34 = base.I64_trunc_f64_s(v26)
							v36 = v34
						} else {
							v36 = int64(-9223372036854775807 - 1)
						}
						v37 = int32(60)
						v45 = v36 + base.I64_extend_i32_u((v14*v37+v11)*v37)*int64(1000000)
						if v45 < int64(86400000001) {
							v70 = F_Int64GetDatum(m, v45)
							mBase = m.M
							v71 = m.ExcPending
							if v71 != 0 {
								return int32(0)
							} else {
								m.G0 = v9 + int32(16)
								return v70
							}
						} else {
							F_errstart_cold(m, int32(21), int32(0))
							mBase = m.M
							v55 = m.ExcPending
							if v55 != 0 {
								return int32(0)
							} else {
								F_errcode(m, int32(134217858))
								mBase = m.M
								v58 = m.ExcPending
								if v58 != 0 {
									return int32(0)
								} else {
									*(*float64)(unsafe.Add(mBase, uint32(v9)+8)) = v13
									*(*int32)(unsafe.Add(mBase, uint32(v9)+4)) = v11
									*(*int32)(unsafe.Add(mBase, uint32(v9))) = v14
									F_errmsg(m, int32(356120), v9)
									mBase = m.M
									v64 = m.ExcPending
									if v64 != 0 {
										return int32(0)
									} else {
										F_errfinish(m, int32(524949), int32(1654), int32(395199))
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
							}
						}
					}
				}
			}
		}
	}
}
func F_time_mi_interval(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v5 int32
	_ = v5
	var v6 int64
	_ = v6
	var v7 int32
	_ = v7
	var v8 int32
	_ = v8
	var v13 int64
	_ = v13
	var v14 int64
	_ = v14
	var v15 int32
	_ = v15
	var v20 int64
	_ = v20
	var v21 int32
	_ = v21
	var v32 int32
	_ = v32
	var v35 int32
	_ = v35
	var v39 int32
	_ = v39
	var v44 int32
	_ = v44
	var v45 int64
	_ = v45
	var v47 int64
	_ = v47
	var v48 int64
	_ = v48
	var v53 int64
	_ = v53
	var v54 int32
	_ = v54
	var v55 int32
	_ = v55
	v5 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v6 = *(*int64)(unsafe.Add(mBase, uint32(v5)))
	v7 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
	v8 = *(*int32)(unsafe.Add(mBase, uint32(v7)+12))
	if v8 != int32(-2147483648) {
		if v8 == int32(2147483647) {
			v20 = *(*int64)(unsafe.Add(mBase, uint32(v7)))
			v21 = *(*int32)(unsafe.Add(mBase, uint32(v7)+8))
			if v21 != int32(2147483647) {
				v45 = v20
				v47 = int64(86400000000)
				v48 = base.I64_rem_s(v6-v45, v47)
				if v48 < int64(0) {
					v53 = v48 + v47
				} else {
					v53 = v48
				}
				v54 = F_Int64GetDatum(m, v53)
				mBase = m.M
				v55 = m.ExcPending
				if v55 != 0 {
					return int32(0)
				} else {
					return v54
				}
			} else {
				if v20 != int64(9223372036854775807) {
					v45 = v20
					v47 = int64(86400000000)
					v48 = base.I64_rem_s(v6-v45, v47)
					if v48 < int64(0) {
						v53 = v48 + v47
					} else {
						v53 = v48
					}
					v54 = F_Int64GetDatum(m, v53)
					mBase = m.M
					v55 = m.ExcPending
					if v55 != 0 {
						return int32(0)
					} else {
						return v54
					}
				} else {
					F_errstart_cold(m, int32(21), int32(0))
					mBase = m.M
					v32 = m.ExcPending
					if v32 != 0 {
						return int32(0)
					} else {
						F_errcode(m, int32(134217858))
						mBase = m.M
						v35 = m.ExcPending
						if v35 != 0 {
							return int32(0)
						} else {
							F_errmsg(m, int32(395445), int32(0))
							mBase = m.M
							v39 = m.ExcPending
							if v39 != 0 {
								return int32(0)
							} else {
								F_errfinish(m, int32(524949), int32(2149), int32(325193))
								mBase = m.M
								v44 = m.ExcPending
								if v44 != 0 {
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
		} else {
			v13 = *(*int64)(unsafe.Add(mBase, uint32(v7)))
			v45 = v13
			v47 = int64(86400000000)
			v48 = base.I64_rem_s(v6-v45, v47)
			if v48 < int64(0) {
				v53 = v48 + v47
			} else {
				v53 = v48
			}
			v54 = F_Int64GetDatum(m, v53)
			mBase = m.M
			v55 = m.ExcPending
			if v55 != 0 {
				return int32(0)
			} else {
				return v54
			}
		}
	} else {
		v14 = *(*int64)(unsafe.Add(mBase, uint32(v7)))
		v15 = *(*int32)(unsafe.Add(mBase, uint32(v7)+8))
		if v15 != int32(-2147483648) {
			v45 = v14
			v47 = int64(86400000000)
			v48 = base.I64_rem_s(v6-v45, v47)
			if v48 < int64(0) {
				v53 = v48 + v47
			} else {
				v53 = v48
			}
			v54 = F_Int64GetDatum(m, v53)
			mBase = m.M
			v55 = m.ExcPending
			if v55 != 0 {
				return int32(0)
			} else {
				return v54
			}
		} else {
			if v14 == int64(-9223372036854775807-1) {
				F_errstart_cold(m, int32(21), int32(0))
				mBase = m.M
				v32 = m.ExcPending
				if v32 != 0 {
					return int32(0)
				} else {
					F_errcode(m, int32(134217858))
					mBase = m.M
					v35 = m.ExcPending
					if v35 != 0 {
						return int32(0)
					} else {
						F_errmsg(m, int32(395445), int32(0))
						mBase = m.M
						v39 = m.ExcPending
						if v39 != 0 {
							return int32(0)
						} else {
							F_errfinish(m, int32(524949), int32(2149), int32(325193))
							mBase = m.M
							v44 = m.ExcPending
							if v44 != 0 {
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
				v45 = v14
				v47 = int64(86400000000)
				v48 = base.I64_rem_s(v6-v45, v47)
				if v48 < int64(0) {
					v53 = v48 + v47
				} else {
					v53 = v48
				}
				v54 = F_Int64GetDatum(m, v53)
				mBase = m.M
				v55 = m.ExcPending
				if v55 != 0 {
					return int32(0)
				} else {
					return v54
				}
			}
		}
	}
}
func F_time_out(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v5 int32
	_ = v5
	var v7 int32
	_ = v7
	var v9 int32
	_ = v9
	var v10 int64
	_ = v10
	var v12 int64
	_ = v12
	var v17 int64
	_ = v17
	var v19 int64
	_ = v19
	var v24 int64
	_ = v24
	var v26 int64
	_ = v26
	var v29 int32
	_ = v29
	var v38 int32
	_ = v38
	var v39 int32
	_ = v39
	var v40 int32
	_ = v40
	var v41 int32
	_ = v41
	var v43 int32
	_ = v43
	var v45 int32
	_ = v45
	var v47 int32
	_ = v47
	var v52 int32
	_ = v52
	var v54 int32
	_ = v54
	var v57 int32
	_ = v57
	var v59 int32
	_ = v59
	var v62 int32
	_ = v62
	v5 = m.G0
	v7 = v5 - int32(176)
	m.G0 = v7
	v9 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v10 = *(*int64)(unsafe.Add(mBase, uint32(v9)))
	v12 = base.I64_div_s(v10, int64(3600000000))
	*(*uint32)(unsafe.Add(mBase, uint32(v7)+140)) = uint32(v12)
	v17 = v10 + base.I64_extend32_s(v12)*int64(-3600000000)
	v19 = base.I64_div_s(v17, int64(60000000))
	*(*uint32)(unsafe.Add(mBase, uint32(v7)+136)) = uint32(v19)
	v24 = base.I64_extend32_s(v19)*int64(-60000000) + v17
	v26 = base.I64_div_s(v24, int64(1000000))
	*(*uint32)(unsafe.Add(mBase, uint32(v7)+132)) = uint32(v26)
	v29 = v7 + int32(132)
	v38 = *(*int32)(unsafe.Add(mBase, uint32(v29)+8))
	v39 = int32(2)
	v40 = F_pg_ultostr_zeropad(m, v7, v38, v39)
	mBase = m.M
	v41 = int32(58)
	*(*uint8)(unsafe.Add(mBase, uint32(v40))) = uint8(v41)
	v43 = int32(1)
	v45 = *(*int32)(unsafe.Add(mBase, uint32(v29)+4))
	v47 = F_pg_ultostr_zeropad(m, v40+v43, v45, v39)
	mBase = m.M
	*(*uint8)(unsafe.Add(mBase, uint32(v47))) = uint8(v41)
	v52 = *(*int32)(unsafe.Add(mBase, uint32(v29)))
	v54 = F_AppendSeconds(m, v47+v43, v52, base.I32_wrap_i64(v26*int64(4293967296)+v24), v43)
	mBase = m.M
	v57 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v54))) = uint8(v57)
	v59 = F_pstrdup(m, v7)
	mBase = m.M
	v62 = m.ExcPending
	if v62 != 0 {
		return int32(0)
	} else {
		m.G0 = v7 + int32(176)
		return v59
	}
}
