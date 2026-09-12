package p3

import (
	base "github.com/shibukawa/pgmem/internal/aot/pgaot/base"
	"unsafe"
)

func F_xid8toxid(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v2 int32
	_ = v2
	var v3 int32
	_ = v3
	v2 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v3 = *(*int32)(unsafe.Add(mBase, uint32(v2)))
	return v3
}
func F_xid_age(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v3 int32
	_ = v3
	var v6 int32
	_ = v6
	var v7 int32
	_ = v7
	var v9 int32
	_ = v9
	var v12 int32
	_ = v12
	var v17 int32
	_ = v17
	var v22 int64
	_ = v22
	var v25 int32
	_ = v25
	var v26 int32
	_ = v26
	var v28 int32
	_ = v28
	var v30 int32
	_ = v30
	var v34 int32
	_ = v34
	v3 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v6 = *(*int32)(unsafe.Add(mBase, _consts[137]))
	v7 = *(*int32)(unsafe.Add(mBase, uint32(v6)+56))
	v9 = *(*int32)(unsafe.Add(mBase, _consts[1060]))
	if v7 == v9 {
		v12 = *(*int32)(unsafe.Add(mBase, _consts[1061]))
		v30 = v12
		if base.Ui32(v3) <= base.Ui32(int32(2)) {
			v34 = int32(2147483647)
		} else {
			v34 = v30 - v3
		}
		return v34
	} else {
		*(*int32)(unsafe.Add(mBase, _consts[1060])) = v7
		v17 = *(*int32)(unsafe.Add(mBase, _consts[80]))
		*(*int32)(unsafe.Add(mBase, _consts[1061])) = v17
		if v17 == int32(0) {
			v22 = F_ReadNextFullTransactionId(m)
			mBase = m.M
			v25 = m.ExcPending
			if v25 != 0 {
				return int32(0)
			} else {
				v26 = base.I32_wrap_i64(v22)
				*(*int32)(unsafe.Add(mBase, _consts[1061])) = v26
				v28 = v26
				v30 = v28
				if base.Ui32(v3) <= base.Ui32(int32(2)) {
					v34 = int32(2147483647)
				} else {
					v34 = v30 - v3
				}
				return v34
			}
		} else {
			v28 = v17
			v30 = v28
			if base.Ui32(v3) <= base.Ui32(int32(2)) {
				v34 = int32(2147483647)
			} else {
				v34 = v30 - v3
			}
			return v34
		}
	}
}
func F_xidout(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v4 int32
	_ = v4
	var v5 int32
	_ = v5
	var v6 int32
	_ = v6
	var v8 int32
	_ = v8
	var v10 int32
	_ = v10
	var v13 int32
	_ = v13
	var v17 int32
	_ = v17
	var v18 int32
	_ = v18
	v4 = m.G0
	v5 = int32(16)
	v6 = v4 - v5
	m.G0 = v6
	v8 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v10 = F_palloc(m, v5)
	mBase = m.M
	v13 = m.ExcPending
	if v13 != 0 {
		return int32(0)
	} else {
		*(*int32)(unsafe.Add(mBase, uint32(v6))) = v8
		v17 = F_pg_snprintf(m, v10, int32(16), int32(38619), v6)
		mBase = m.M
		v18 = m.ExcPending
		if v18 != 0 {
			return int32(0)
		} else {
			m.G0 = v6 + int32(16)
			return v10
		}
	}
}
func F_xmlexists(m *base.Module, l0 int32) int32 {
	var v7 int32
	_ = v7
	var v10 int32
	_ = v10
	var v14 int32
	_ = v14
	var v18 int32
	_ = v18
	var v23 int32
	_ = v23
	F_errstart_cold(m, int32(21), int32(0))
	v7 = m.ExcPending
	if v7 != 0 {
		return int32(0)
	} else {
		F_errcode(m, int32(1088))
		v10 = m.ExcPending
		if v10 != 0 {
			return int32(0)
		} else {
			F_errmsg(m, int32(364009), int32(0))
			v14 = m.ExcPending
			if v14 != 0 {
				return int32(0)
			} else {
				F_errdetail(m, int32(580218), int32(0))
				v18 = m.ExcPending
				if v18 != 0 {
					return int32(0)
				} else {
					F_errfinish(m, int32(498523), int32(4555), int32(115952))
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
	}
}
func F_xmlvalidate(m *base.Module, l0 int32) int32 {
	var v7 int32
	_ = v7
	var v10 int32
	_ = v10
	var v14 int32
	_ = v14
	var v19 int32
	_ = v19
	F_errstart_cold(m, int32(21), int32(0))
	v7 = m.ExcPending
	if v7 != 0 {
		return int32(0)
	} else {
		F_errcode(m, int32(1088))
		v10 = m.ExcPending
		if v10 != 0 {
			return int32(0)
		} else {
			F_errmsg(m, int32(446798), int32(0))
			v14 = m.ExcPending
			if v14 != 0 {
				return int32(0)
			} else {
				F_errfinish(m, int32(498523), int32(1123), int32(357520))
				v19 = m.ExcPending
				if v19 != 0 {
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
func F_xpath(m *base.Module, l0 int32) int32 {
	var v7 int32
	_ = v7
	var v10 int32
	_ = v10
	var v14 int32
	_ = v14
	var v18 int32
	_ = v18
	var v23 int32
	_ = v23
	F_errstart_cold(m, int32(21), int32(0))
	v7 = m.ExcPending
	if v7 != 0 {
		return int32(0)
	} else {
		F_errcode(m, int32(1088))
		v10 = m.ExcPending
		if v10 != 0 {
			return int32(0)
		} else {
			F_errmsg(m, int32(364009), int32(0))
			v14 = m.ExcPending
			if v14 != 0 {
				return int32(0)
			} else {
				F_errdetail(m, int32(580218), int32(0))
				v18 = m.ExcPending
				if v18 != 0 {
					return int32(0)
				} else {
					F_errfinish(m, int32(498523), int32(4533), int32(321992))
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
	}
}
