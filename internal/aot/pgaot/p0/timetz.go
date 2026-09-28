package p0

import (
	base "github.com/shibukawa/pgmem/internal/aot/pgaot/base"
	"unsafe"
)

func F_timetz_ge(m *base.Module, l0 int32) int64 {
	mBase := m.M
	_ = mBase
	var v7 int32
	_ = v7
	var v8 int32
	_ = v8
	var v10 int64
	_ = v10
	var v12 int64
	_ = v12
	var v13 int64
	_ = v13
	var v14 int32
	_ = v14
	var v15 int32
	_ = v15
	var v19 int64
	_ = v19
	var v20 int64
	_ = v20
	var v28 int64
	_ = v28
	v7 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
	v8 = *(*int32)(unsafe.Add(mBase, uint32(v7)+8))
	v10 = int64(1000000)
	v12 = *(*int64)(unsafe.Add(mBase, uint32(v7)))
	v13 = base.I64_extend_i32_s(v8)*v10 + v12
	v14 = *(*int32)(unsafe.Add(mBase, uint32(l0)+40))
	v15 = *(*int32)(unsafe.Add(mBase, uint32(v14)+8))
	v19 = *(*int64)(unsafe.Add(mBase, uint32(v14)))
	v20 = base.I64_extend_i32_s(v15)*v10 + v19
	if v20 < v13 {
		v28 = int64(1)
	} else {
		if v13 < v20 {
			v28 = int64(0)
		} else {
			if v15 < v8 {
				v28 = int64(1)
			} else {
				v28 = base.I64_extend_i32_u(base.B2i32(v15 <= v8))
			}
		}
	}
	return v28
}
func F_timetz_lt(m *base.Module, l0 int32) int64 {
	mBase := m.M
	_ = mBase
	var v7 int32
	_ = v7
	var v8 int32
	_ = v8
	var v10 int64
	_ = v10
	var v12 int64
	_ = v12
	var v13 int64
	_ = v13
	var v14 int32
	_ = v14
	var v15 int32
	_ = v15
	var v19 int64
	_ = v19
	var v20 int64
	_ = v20
	var v28 int64
	_ = v28
	v7 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
	v8 = *(*int32)(unsafe.Add(mBase, uint32(v7)+8))
	v10 = int64(1000000)
	v12 = *(*int64)(unsafe.Add(mBase, uint32(v7)))
	v13 = base.I64_extend_i32_s(v8)*v10 + v12
	v14 = *(*int32)(unsafe.Add(mBase, uint32(l0)+40))
	v15 = *(*int32)(unsafe.Add(mBase, uint32(v14)+8))
	v19 = *(*int64)(unsafe.Add(mBase, uint32(v14)))
	v20 = base.I64_extend_i32_s(v15)*v10 + v19
	if v20 < v13 {
		v28 = int64(0)
	} else {
		if v13 < v20 {
			v28 = int64(1)
		} else {
			if v15 < v8 {
				v28 = int64(0)
			} else {
				v28 = base.I64_extend_i32_u(base.B2i32(v8 < v15))
			}
		}
	}
	return v28
}
func F_timetz_mi_interval(m *base.Module, l0 int32) int64 {
	mBase := m.M
	_ = mBase
	var v5 int64
	_ = v5
	var v6 int32
	_ = v6
	var v7 int32
	_ = v7
	var v12 int32
	_ = v12
	var v15 int64
	_ = v15
	var v18 int32
	_ = v18
	var v21 int64
	_ = v21
	var v29 int32
	_ = v29
	var v32 int32
	_ = v32
	var v36 int32
	_ = v36
	var v41 int32
	_ = v41
	var v43 int32
	_ = v43
	var v44 int32
	_ = v44
	var v45 int32
	_ = v45
	var v46 int64
	_ = v46
	var v47 int64
	_ = v47
	var v49 int64
	_ = v49
	var v50 int64
	_ = v50
	var v55 int64
	_ = v55
	var v57 int32
	_ = v57
	v5 = *(*int64)(unsafe.Add(mBase, uint32(l0)+24))
	v6 = *(*int32)(unsafe.Add(mBase, uint32(l0)+40))
	v7 = *(*int32)(unsafe.Add(mBase, uint32(v6)+12))
	if v7 != int32(2147483647) {
		if v7 != int32(-2147483648) {
			v43 = F_palloc(m, int32(16))
			mBase = m.M
			v44 = m.ExcPending
			if v44 != 0 {
				return int64(0)
			} else {
				v45 = base.I32_wrap_i64(v5)
				v46 = *(*int64)(unsafe.Add(mBase, uint32(v45)))
				v47 = *(*int64)(unsafe.Add(mBase, uint32(v6)))
				v49 = int64(86400000000)
				v50 = base.I64_rem_s(v46-v47, v49)
				if v50 < int64(0) {
					v55 = v50 + v49
				} else {
					v55 = v50
				}
				*(*int64)(unsafe.Add(mBase, uint32(v43))) = v55
				v57 = *(*int32)(unsafe.Add(mBase, uint32(v45)+8))
				*(*int32)(unsafe.Add(mBase, uint32(v43)+8)) = v57
				return base.I64_extend_i32_u(v43)
			}
		} else {
			v12 = *(*int32)(unsafe.Add(mBase, uint32(v6)+8))
			if v12 != int32(-2147483648) {
				v43 = F_palloc(m, int32(16))
				mBase = m.M
				v44 = m.ExcPending
				if v44 != 0 {
					return int64(0)
				} else {
					v45 = base.I32_wrap_i64(v5)
					v46 = *(*int64)(unsafe.Add(mBase, uint32(v45)))
					v47 = *(*int64)(unsafe.Add(mBase, uint32(v6)))
					v49 = int64(86400000000)
					v50 = base.I64_rem_s(v46-v47, v49)
					if v50 < int64(0) {
						v55 = v50 + v49
					} else {
						v55 = v50
					}
					*(*int64)(unsafe.Add(mBase, uint32(v43))) = v55
					v57 = *(*int32)(unsafe.Add(mBase, uint32(v45)+8))
					*(*int32)(unsafe.Add(mBase, uint32(v43)+8)) = v57
					return base.I64_extend_i32_u(v43)
				}
			} else {
				v15 = *(*int64)(unsafe.Add(mBase, uint32(v6)))
				if v15 == int64(-9223372036854775807-1) {
					F_errstart_cold(m, int32(21), int32(0))
					mBase = m.M
					v29 = m.ExcPending
					if v29 != 0 {
						return int64(0)
					} else {
						F_errcode(m, int32(134217858))
						mBase = m.M
						v32 = m.ExcPending
						if v32 != 0 {
							return int64(0)
						} else {
							F_errmsg(m, int32(_a_F_timetz_mi_interval_0), int32(0))
							mBase = m.M
							v36 = m.ExcPending
							if v36 != 0 {
								return int64(0)
							} else {
								F_errfinish(m, int32(_a_F_timetz_mi_interval_1), int32(2755), int32(_a_F_timetz_mi_interval_2))
								mBase = m.M
								v41 = m.ExcPending
								if v41 != 0 {
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
					v43 = F_palloc(m, int32(16))
					mBase = m.M
					v44 = m.ExcPending
					if v44 != 0 {
						return int64(0)
					} else {
						v45 = base.I32_wrap_i64(v5)
						v46 = *(*int64)(unsafe.Add(mBase, uint32(v45)))
						v47 = *(*int64)(unsafe.Add(mBase, uint32(v6)))
						v49 = int64(86400000000)
						v50 = base.I64_rem_s(v46-v47, v49)
						if v50 < int64(0) {
							v55 = v50 + v49
						} else {
							v55 = v50
						}
						*(*int64)(unsafe.Add(mBase, uint32(v43))) = v55
						v57 = *(*int32)(unsafe.Add(mBase, uint32(v45)+8))
						*(*int32)(unsafe.Add(mBase, uint32(v43)+8)) = v57
						return base.I64_extend_i32_u(v43)
					}
				}
			}
		}
	} else {
		v18 = *(*int32)(unsafe.Add(mBase, uint32(v6)+8))
		if v18 != int32(2147483647) {
			v43 = F_palloc(m, int32(16))
			mBase = m.M
			v44 = m.ExcPending
			if v44 != 0 {
				return int64(0)
			} else {
				v45 = base.I32_wrap_i64(v5)
				v46 = *(*int64)(unsafe.Add(mBase, uint32(v45)))
				v47 = *(*int64)(unsafe.Add(mBase, uint32(v6)))
				v49 = int64(86400000000)
				v50 = base.I64_rem_s(v46-v47, v49)
				if v50 < int64(0) {
					v55 = v50 + v49
				} else {
					v55 = v50
				}
				*(*int64)(unsafe.Add(mBase, uint32(v43))) = v55
				v57 = *(*int32)(unsafe.Add(mBase, uint32(v45)+8))
				*(*int32)(unsafe.Add(mBase, uint32(v43)+8)) = v57
				return base.I64_extend_i32_u(v43)
			}
		} else {
			v21 = *(*int64)(unsafe.Add(mBase, uint32(v6)))
			if v21 != int64(9223372036854775807) {
				v43 = F_palloc(m, int32(16))
				mBase = m.M
				v44 = m.ExcPending
				if v44 != 0 {
					return int64(0)
				} else {
					v45 = base.I32_wrap_i64(v5)
					v46 = *(*int64)(unsafe.Add(mBase, uint32(v45)))
					v47 = *(*int64)(unsafe.Add(mBase, uint32(v6)))
					v49 = int64(86400000000)
					v50 = base.I64_rem_s(v46-v47, v49)
					if v50 < int64(0) {
						v55 = v50 + v49
					} else {
						v55 = v50
					}
					*(*int64)(unsafe.Add(mBase, uint32(v43))) = v55
					v57 = *(*int32)(unsafe.Add(mBase, uint32(v45)+8))
					*(*int32)(unsafe.Add(mBase, uint32(v43)+8)) = v57
					return base.I64_extend_i32_u(v43)
				}
			} else {
				F_errstart_cold(m, int32(21), int32(0))
				mBase = m.M
				v29 = m.ExcPending
				if v29 != 0 {
					return int64(0)
				} else {
					F_errcode(m, int32(134217858))
					mBase = m.M
					v32 = m.ExcPending
					if v32 != 0 {
						return int64(0)
					} else {
						F_errmsg(m, int32(_a_F_timetz_mi_interval_0), int32(0))
						mBase = m.M
						v36 = m.ExcPending
						if v36 != 0 {
							return int64(0)
						} else {
							F_errfinish(m, int32(_a_F_timetz_mi_interval_1), int32(2755), int32(_a_F_timetz_mi_interval_2))
							mBase = m.M
							v41 = m.ExcPending
							if v41 != 0 {
								return int64(0)
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
func F_timetz_out(m *base.Module, l0 int32) int64 {
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
	var v34 int32
	_ = v34
	var v35 int32
	_ = v35
	var v37 int32
	_ = v37
	var v38 int32
	_ = v38
	var v39 int32
	_ = v39
	var v40 int32
	_ = v40
	var v41 int32
	_ = v41
	var v45 int32
	_ = v45
	var v47 int32
	_ = v47
	var v52 int32
	_ = v52
	var v54 int32
	_ = v54
	var v55 int32
	_ = v55
	var v57 int32
	_ = v57
	var v59 int32
	_ = v59
	var v62 int32
	_ = v62
	v5 = m.G0
	v7 = v5 - int32(176)
	m.G0 = v7
	v9 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
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
	v34 = int32(1)
	v35 = *(*int32)(unsafe.Add(mBase, uint32(v9)+8))
	v37 = *(*int32)(unsafe.Add(mBase, _c_F_timetz_out[0]))
	v38 = *(*int32)(unsafe.Add(mBase, uint32(v29)+8))
	v39 = int32(2)
	v40 = F_pg_ultostr_zeropad(m, v7, v38, v39)
	mBase = m.M
	v41 = int32(58)
	*(*uint8)(unsafe.Add(mBase, uint32(v40))) = uint8(v41)
	v45 = *(*int32)(unsafe.Add(mBase, uint32(v29)+4))
	v47 = F_pg_ultostr_zeropad(m, v40+v34, v45, v39)
	mBase = m.M
	*(*uint8)(unsafe.Add(mBase, uint32(v47))) = uint8(v41)
	v52 = *(*int32)(unsafe.Add(mBase, uint32(v29)))
	v54 = F_AppendSeconds(m, v47+v34, v52, base.I32_wrap_i64(v26*int64(4293967296)+v24), v34)
	mBase = m.M
	v55 = F_EncodeTimezone(m, v54, v35, v37)
	mBase = m.M
	v57 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v55))) = uint8(v57)
	v59 = F_pstrdup(m, v7)
	mBase = m.M
	v62 = m.ExcPending
	if v62 != 0 {
		return int64(0)
	} else {
		m.G0 = v7 + int32(176)
		return base.I64_extend_i32_u(v59)
	}
}
func F_timetz_recv(m *base.Module, l0 int32) int64 {
	mBase := m.M
	_ = mBase
	var v7 int32
	_ = v7
	var v8 int32
	_ = v8
	var v10 int32
	_ = v10
	var v13 int32
	_ = v13
	var v14 int64
	_ = v14
	var v15 int32
	_ = v15
	var v20 int32
	_ = v20
	var v21 int32
	_ = v21
	var v30 int32
	_ = v30
	var v31 int64
	_ = v31
	var v32 int64
	_ = v32
	var v33 int64
	_ = v33
	var v36 int64
	_ = v36
	var v37 int64
	_ = v37
	var v39 int64
	_ = v39
	var v40 int64
	_ = v40
	var v43 int64
	_ = v43
	var v54 int32
	_ = v54
	var v57 int32
	_ = v57
	var v61 int32
	_ = v61
	var v66 int32
	_ = v66
	var v70 int32
	_ = v70
	var v73 int32
	_ = v73
	var v77 int32
	_ = v77
	var v82 int32
	_ = v82
	v7 = *(*int32)(unsafe.Add(mBase, uint32(l0)+56))
	v8 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
	v10 = F_palloc(m, int32(16))
	mBase = m.M
	v13 = m.ExcPending
	if v13 != 0 {
		return int64(0)
	} else {
		v14 = F_pq_getmsgint64(m, v8)
		mBase = m.M
		v15 = m.ExcPending
		if v15 != 0 {
			return int64(0)
		} else {
			*(*int64)(unsafe.Add(mBase, uint32(v10))) = v14
			if base.Ui64(v14) < base.Ui64(int64(86400000001)) {
				v20 = F_pq_getmsgint(m, v8, int32(4))
				mBase = m.M
				v21 = m.ExcPending
				if v21 != 0 {
					return int64(0)
				} else {
					*(*int32)(unsafe.Add(mBase, uint32(v10)+8)) = v20
					if base.Ui32(v20-int32(_a_F_timetz_recv_0)) <= base.Ui32(int32(-115200)) {
						F_errstart_cold(m, int32(21), int32(0))
						mBase = m.M
						v70 = m.ExcPending
						if v70 != 0 {
							return int64(0)
						} else {
							F_errcode(m, int32(150995074))
							mBase = m.M
							v73 = m.ExcPending
							if v73 != 0 {
								return int64(0)
							} else {
								F_errmsg(m, int32(_a_F_timetz_recv_1), int32(0))
								mBase = m.M
								v77 = m.ExcPending
								if v77 != 0 {
									return int64(0)
								} else {
									F_errfinish(m, int32(_a_F_timetz_recv_2), int32(2479), int32(_a_F_timetz_recv_3))
									mBase = m.M
									v82 = m.ExcPending
									if v82 != 0 {
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
						if base.Ui32(v7) <= base.Ui32(int32(6)) {
							v30 = v7 << (uint(int32(3)) % 32)
							v31 = *(*int64)(unsafe.Add(mBase, uint32(v30)+uint32(_c_F_timetz_recv[0])))
							v32 = *(*int64)(unsafe.Add(mBase, uint32(v30)+uint32(_c_F_timetz_recv[1])))
							v33 = *(*int64)(unsafe.Add(mBase, uint32(v10)))
							if int64(0) <= v33 {
								v36 = v32 + v33
								v37 = base.I64_rem_s(v36, v31)
								v43 = v36 - v37
							} else {
								v39 = v32 - v33
								v40 = base.I64_rem_s(v39, v31)
								v43 = v40 - v39
							}
							*(*int64)(unsafe.Add(mBase, uint32(v10))) = v43
						} else {
						}
						return base.I64_extend_i32_u(v10)
					}
				}
			} else {
				F_errstart_cold(m, int32(21), int32(0))
				mBase = m.M
				v54 = m.ExcPending
				if v54 != 0 {
					return int64(0)
				} else {
					F_errcode(m, int32(134217858))
					mBase = m.M
					v57 = m.ExcPending
					if v57 != 0 {
						return int64(0)
					} else {
						F_errmsg(m, int32(_a_F_timetz_recv_4), int32(0))
						mBase = m.M
						v61 = m.ExcPending
						if v61 != 0 {
							return int64(0)
						} else {
							F_errfinish(m, int32(_a_F_timetz_recv_2), int32(2471), int32(_a_F_timetz_recv_3))
							mBase = m.M
							v66 = m.ExcPending
							if v66 != 0 {
								return int64(0)
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
func F_timetz_scale(m *base.Module, l0 int32) int64 {
	mBase := m.M
	_ = mBase
	var v7 int32
	_ = v7
	var v8 int32
	_ = v8
	var v10 int32
	_ = v10
	var v13 int32
	_ = v13
	var v14 int64
	_ = v14
	var v16 int32
	_ = v16
	var v21 int32
	_ = v21
	var v22 int64
	_ = v22
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
	var v33 int64
	_ = v33
	v7 = *(*int32)(unsafe.Add(mBase, uint32(l0)+40))
	v8 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
	v10 = F_palloc(m, int32(16))
	mBase = m.M
	v13 = m.ExcPending
	if v13 != 0 {
		return int64(0)
	} else {
		v14 = *(*int64)(unsafe.Add(mBase, uint32(v8)))
		*(*int64)(unsafe.Add(mBase, uint32(v10))) = v14
		v16 = *(*int32)(unsafe.Add(mBase, uint32(v8)+8))
		*(*int32)(unsafe.Add(mBase, uint32(v10)+8)) = v16
		if base.Ui32(v7) <= base.Ui32(int32(6)) {
			v21 = v7 << (uint(int32(3)) % 32)
			v22 = *(*int64)(unsafe.Add(mBase, uint32(v21)+uint32(_c_F_timetz_scale[0])))
			v23 = *(*int64)(unsafe.Add(mBase, uint32(v21)+uint32(_c_F_timetz_scale[1])))
			if int64(0) <= v14 {
				v26 = v14 + v23
				v27 = base.I64_rem_s(v26, v22)
				v33 = v26 - v27
			} else {
				v29 = v23 - v14
				v30 = base.I64_rem_s(v29, v22)
				v33 = v30 - v29
			}
			*(*int64)(unsafe.Add(mBase, uint32(v10))) = v33
		} else {
		}
		return base.I64_extend_i32_u(v10)
	}
}
func F_timetz_smaller(m *base.Module, l0 int32) int64 {
	mBase := m.M
	_ = mBase
	var v8 int64
	_ = v8
	var v11 int64
	_ = v11
	var v12 int32
	_ = v12
	var v13 int32
	_ = v13
	var v15 int64
	_ = v15
	var v17 int64
	_ = v17
	var v18 int64
	_ = v18
	var v19 int32
	_ = v19
	var v20 int32
	_ = v20
	var v24 int64
	_ = v24
	var v25 int64
	_ = v25
	var v35 int64
	_ = v35
	v8 = *(*int64)(unsafe.Add(mBase, uint32(l0)+40))
	v11 = *(*int64)(unsafe.Add(mBase, uint32(l0)+24))
	v12 = base.I32_wrap_i64(v11)
	v13 = *(*int32)(unsafe.Add(mBase, uint32(v12)+8))
	v15 = int64(1000000)
	v17 = *(*int64)(unsafe.Add(mBase, uint32(v12)))
	v18 = base.I64_extend_i32_s(v13)*v15 + v17
	v19 = base.I32_wrap_i64(v8)
	v20 = *(*int32)(unsafe.Add(mBase, uint32(v19)+8))
	v24 = *(*int64)(unsafe.Add(mBase, uint32(v19)))
	v25 = base.I64_extend_i32_s(v20)*v15 + v24
	if base.B2i32(v25 < v18)|(base.B2i32(v20 < v13)|base.B2i32(v20 <= v13))&base.B2i32(v25 <= v18) != 0 {
		v35 = v8 & int64(4294967295)
	} else {
		v35 = v11 & int64(4294967295)
	}
	return v35
}
