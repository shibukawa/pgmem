package p3

import (
	base "github.com/shibukawa/pgmem/internal/aot/pgaot/base"
	"unsafe"
)

func F_cash_div_int2(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v3 int64
	_ = v3
	var v5 int64
	_ = v5
	var v13 int32
	_ = v13
	var v16 int32
	_ = v16
	var v20 int32
	_ = v20
	var v25 int32
	_ = v25
	var v26 int32
	_ = v26
	var v27 int64
	_ = v27
	var v30 int64
	_ = v30
	var v31 int32
	_ = v31
	var v32 int32
	_ = v32
	v3 = int64(*(*uint32)(unsafe.Add(mBase, uint32(l0)+28)))
	v5 = v3 << (uint(int64(48)) % 64)
	if v5 == int64(0) {
		F_errstart_cold(m, int32(21), int32(0))
		mBase = m.M
		v13 = m.ExcPending
		if v13 != 0 {
			return int32(0)
		} else {
			F_errcode(m, int32(33816706))
			mBase = m.M
			v16 = m.ExcPending
			if v16 != 0 {
				return int32(0)
			} else {
				F_errmsg(m, int32(_a_F_cash_div_int2_0), int32(0))
				mBase = m.M
				v20 = m.ExcPending
				if v20 != 0 {
					return int32(0)
				} else {
					F_errfinish(m, int32(_a_F_cash_div_int2_1), int32(161), int32(_a_F_cash_div_int2_2))
					mBase = m.M
					v25 = m.ExcPending
					if v25 != 0 {
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
		v26 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
		v27 = *(*int64)(unsafe.Add(mBase, uint32(v26)))
		v30 = base.I64_div_s(v27, v5>>(uint(int64(48))%64))
		v31 = F_Int64GetDatum(m, v30)
		mBase = m.M
		v32 = m.ExcPending
		if v32 != 0 {
			return int32(0)
		} else {
			return v31
		}
	}
}
func F_cash_div_int4(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v3 int32
	_ = v3
	var v11 int32
	_ = v11
	var v14 int32
	_ = v14
	var v18 int32
	_ = v18
	var v23 int32
	_ = v23
	var v24 int32
	_ = v24
	var v25 int64
	_ = v25
	var v27 int64
	_ = v27
	var v28 int32
	_ = v28
	var v29 int32
	_ = v29
	v3 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
	if v3 == int32(0) {
		F_errstart_cold(m, int32(21), int32(0))
		mBase = m.M
		v11 = m.ExcPending
		if v11 != 0 {
			return int32(0)
		} else {
			F_errcode(m, int32(33816706))
			mBase = m.M
			v14 = m.ExcPending
			if v14 != 0 {
				return int32(0)
			} else {
				F_errmsg(m, int32(_a_F_cash_div_int4_0), int32(0))
				mBase = m.M
				v18 = m.ExcPending
				if v18 != 0 {
					return int32(0)
				} else {
					F_errfinish(m, int32(_a_F_cash_div_int4_1), int32(161), int32(_a_F_cash_div_int4_2))
					mBase = m.M
					v23 = m.ExcPending
					if v23 != 0 {
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
		v24 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
		v25 = *(*int64)(unsafe.Add(mBase, uint32(v24)))
		v27 = base.I64_div_s(v25, base.I64_extend_i32_s(v3))
		v28 = F_Int64GetDatum(m, v27)
		mBase = m.M
		v29 = m.ExcPending
		if v29 != 0 {
			return int32(0)
		} else {
			return v28
		}
	}
}
func F_cash_div_int8(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v3 int32
	_ = v3
	var v4 int64
	_ = v4
	var v12 int32
	_ = v12
	var v15 int32
	_ = v15
	var v19 int32
	_ = v19
	var v24 int32
	_ = v24
	var v25 int32
	_ = v25
	var v26 int64
	_ = v26
	var v27 int64
	_ = v27
	var v28 int32
	_ = v28
	var v29 int32
	_ = v29
	v3 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
	v4 = *(*int64)(unsafe.Add(mBase, uint32(v3)))
	if v4 == int64(0) {
		F_errstart_cold(m, int32(21), int32(0))
		mBase = m.M
		v12 = m.ExcPending
		if v12 != 0 {
			return int32(0)
		} else {
			F_errcode(m, int32(33816706))
			mBase = m.M
			v15 = m.ExcPending
			if v15 != 0 {
				return int32(0)
			} else {
				F_errmsg(m, int32(_a_F_cash_div_int8_0), int32(0))
				mBase = m.M
				v19 = m.ExcPending
				if v19 != 0 {
					return int32(0)
				} else {
					F_errfinish(m, int32(_a_F_cash_div_int8_1), int32(161), int32(_a_F_cash_div_int8_2))
					mBase = m.M
					v24 = m.ExcPending
					if v24 != 0 {
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
		v25 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
		v26 = *(*int64)(unsafe.Add(mBase, uint32(v25)))
		v27 = base.I64_div_s(v26, v4)
		v28 = F_Int64GetDatum(m, v27)
		mBase = m.M
		v29 = m.ExcPending
		if v29 != 0 {
			return int32(0)
		} else {
			return v28
		}
	}
}
func F_cash_mul_flt4(m *base.Module, l0 int32) int32 {
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
	v6 = F_cash_mul_float8(m, v3, base.F64_promote_f32(v4))
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
func F_cash_numeric(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v7 int32
	_ = v7
	var v8 int64
	_ = v8
	var v9 int32
	_ = v9
	var v12 int32
	_ = v12
	var v13 int32
	_ = v13
	var v14 int32
	_ = v14
	var v15 int32
	_ = v15
	var v19 int32
	_ = v19
	var v20 int32
	_ = v20
	var v23 int64
	_ = v23
	var v29 int32
	_ = v29
	var v30 int64
	_ = v30
	var v36 int64
	_ = v36
	var v38 int32
	_ = v38
	var v45 int64
	_ = v45
	var v53 int32
	_ = v53
	var v54 int64
	_ = v54
	var v60 int64
	_ = v60
	var v62 int32
	_ = v62
	var v65 int64
	_ = v65
	var v70 int32
	_ = v70
	var v71 int32
	_ = v71
	var v76 int32
	_ = v76
	var v77 int32
	_ = v77
	var v78 int32
	_ = v78
	var v79 int32
	_ = v79
	var v80 int32
	_ = v80
	var v81 int32
	_ = v81
	var v82 int32
	_ = v82
	var v83 int32
	_ = v83
	var v90 int32
	_ = v90
	v7 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v8 = *(*int64)(unsafe.Add(mBase, uint32(v7)))
	v9 = F_PGLC_localeconv(m)
	mBase = m.M
	v12 = m.ExcPending
	if v12 != 0 {
		return int32(0)
	} else {
		v13 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v9)+41)))
		v14 = F_int64_to_numeric(m, v8)
		mBase = m.M
		v15 = m.ExcPending
		if v15 != 0 {
			return int32(0)
		} else {
			if base.Ui32(int32(10)) < base.Ui32(v13) {
				v19 = int32(2)
			} else {
				v19 = v13
			}
			v20 = base.I32_extend8_s(v19)
			if int32(0) < v20 {
				v23 = int64(1)
				if base.Ui32(int32(8)) <= base.Ui32(v19) {
					v29 = int32(0)
					v30 = v23
					for {
						v36 = v30 * int64(100000000)
						v38 = v29 + int32(8)
						if v38 != v20&int32(120) {
							v29 = v38
							v30 = v36
							continue
						} else {
							break
						}
						break
					}
					if v19&int32(7) == int32(0) {
						v65 = v36
					} else {
						v45 = v36
						v53 = int32(0)
						v54 = v45
						for {
							v60 = v54 * int64(10)
							v62 = v53 + int32(1)
							if v62 != v20&int32(7) {
								v53 = v62
								v54 = v60
								continue
							} else {
								break
							}
							break
						}
						v65 = v60
					}
				} else {
					v45 = v23
					v53 = int32(0)
					v54 = v45
					for {
						v60 = v54 * int64(10)
						v62 = v53 + int32(1)
						if v62 != v20&int32(7) {
							v53 = v62
							v54 = v60
							continue
						} else {
							break
						}
						break
					}
					v65 = v60
				}
				v70 = int32(1259)
				v71 = int32(0)
				v76 = F_int64_to_numeric(m, v65)
				mBase = m.M
				v77 = m.ExcPending
				if v77 != 0 {
					return int32(0)
				} else {
					v78 = F_DirectFunctionCall2Coll(m, v70, v71, v76, v20)
					mBase = m.M
					v79 = m.ExcPending
					if v79 != 0 {
						return int32(0)
					} else {
						v80 = F_DirectFunctionCall2Coll(m, int32(1260), v71, v14, v78)
						mBase = m.M
						v81 = m.ExcPending
						if v81 != 0 {
							return int32(0)
						} else {
							v82 = F_DirectFunctionCall2Coll(m, v70, v71, v80, v20)
							mBase = m.M
							v83 = m.ExcPending
							if v83 != 0 {
								return int32(0)
							} else {
								v90 = v82
								return v90
							}
						}
					}
				}
			} else {
				v90 = v14
				return v90
			}
		}
	}
}
