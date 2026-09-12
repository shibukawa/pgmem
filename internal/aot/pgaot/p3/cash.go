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
				F_errmsg(m, int32(237220), int32(0))
				mBase = m.M
				v20 = m.ExcPending
				if v20 != 0 {
					return int32(0)
				} else {
					F_errfinish(m, int32(492480), int32(161), int32(545691))
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
				F_errmsg(m, int32(237220), int32(0))
				mBase = m.M
				v18 = m.ExcPending
				if v18 != 0 {
					return int32(0)
				} else {
					F_errfinish(m, int32(492480), int32(161), int32(545691))
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
				F_errmsg(m, int32(237220), int32(0))
				mBase = m.M
				v19 = m.ExcPending
				if v19 != 0 {
					return int32(0)
				} else {
					F_errfinish(m, int32(492480), int32(161), int32(545691))
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
	var v30 int32
	_ = v30
	var v31 int64
	_ = v31
	var v37 int64
	_ = v37
	var v39 int32
	_ = v39
	var v42 int64
	_ = v42
	var v52 int32
	_ = v52
	var v53 int64
	_ = v53
	var v59 int64
	_ = v59
	var v61 int32
	_ = v61
	var v64 int64
	_ = v64
	var v69 int32
	_ = v69
	var v70 int32
	_ = v70
	var v75 int32
	_ = v75
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
	var v89 int32
	_ = v89
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
				if base.Ui32(v19) < base.Ui32(int32(8)) {
					v42 = int64(1)
				} else {
					v30 = int32(0)
					v31 = int64(1)
					for {
						v37 = v31 * int64(100000000)
						v39 = v30 + int32(8)
						if v39 != v20&int32(120) {
							v30 = v39
							v31 = v37
							continue
						} else {
							break
						}
						break
					}
					v42 = v37
				}
				if v19&int32(7) != 0 {
					v52 = int32(0)
					v53 = v42
					for {
						v59 = v53 * int64(10)
						v61 = v52 + int32(1)
						if v61 != v20&int32(7) {
							v52 = v61
							v53 = v59
							continue
						} else {
							break
						}
						break
					}
					v64 = v59
				} else {
					v64 = v42
				}
				v69 = int32(1275)
				v70 = int32(0)
				v75 = F_int64_to_numeric(m, v64)
				mBase = m.M
				v76 = m.ExcPending
				if v76 != 0 {
					return int32(0)
				} else {
					v77 = F_DirectFunctionCall2Coll(m, v69, v70, v75, v20)
					mBase = m.M
					v78 = m.ExcPending
					if v78 != 0 {
						return int32(0)
					} else {
						v79 = F_DirectFunctionCall2Coll(m, int32(1276), v70, v14, v77)
						mBase = m.M
						v80 = m.ExcPending
						if v80 != 0 {
							return int32(0)
						} else {
							v81 = F_DirectFunctionCall2Coll(m, v69, v70, v79, v20)
							mBase = m.M
							v82 = m.ExcPending
							if v82 != 0 {
								return int32(0)
							} else {
								v89 = v81
								return v89
							}
						}
					}
				}
			} else {
				v89 = v14
				return v89
			}
		}
	}
}
