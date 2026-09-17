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
	v6 = *(*int32)(unsafe.Add(mBase, _c_F_xid_age[0]))
	v7 = *(*int32)(unsafe.Add(mBase, uint32(v6)+56))
	v9 = *(*int32)(unsafe.Add(mBase, _c_F_xid_age[1]))
	if v7 == v9 {
		v12 = *(*int32)(unsafe.Add(mBase, _c_F_xid_age[2]))
		v30 = v12
		if base.Ui32(v3) <= base.Ui32(int32(2)) {
			v34 = int32(2147483647)
		} else {
			v34 = v30 - v3
		}
		return v34
	} else {
		*(*int32)(unsafe.Add(mBase, _c_F_xid_age[1])) = v7
		v17 = *(*int32)(unsafe.Add(mBase, _c_F_xid_age[3]))
		*(*int32)(unsafe.Add(mBase, _c_F_xid_age[2])) = v17
		if v17 == int32(0) {
			v22 = F_ReadNextFullTransactionId(m)
			mBase = m.M
			v25 = m.ExcPending
			if v25 != 0 {
				return int32(0)
			} else {
				v26 = base.I32_wrap_i64(v22)
				*(*int32)(unsafe.Add(mBase, _c_F_xid_age[2])) = v26
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
	var v4 int32
	_ = v4
	var v7 int32
	_ = v7
	v4 = Fn13952(m, l0, int32(_a_F_xidout_0), int32(16))
	v7 = m.ExcPending
	if v7 != 0 {
		return int32(0)
	} else {
		return v4
	}
}
func F_xmlexists(m *base.Module, l0 int32) int32 {
	var v4 int32
	_ = v4
	var v7 int32
	_ = v7
	v4 = Fn13840(m, l0, int32(_a_F_xmlexists_0), int32(_a_F_xmlexists_1))
	v7 = m.ExcPending
	if v7 != 0 {
		return int32(0)
	} else {
		return v4
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
			F_errmsg(m, int32(_a_F_xmlvalidate_0), int32(0))
			v14 = m.ExcPending
			if v14 != 0 {
				return int32(0)
			} else {
				F_errfinish(m, int32(_a_F_xmlvalidate_1), int32(1123), int32(_a_F_xmlvalidate_2))
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
	var v4 int32
	_ = v4
	var v7 int32
	_ = v7
	v4 = Fn13840(m, l0, int32(_a_F_xpath_0), int32(_a_F_xpath_1))
	v7 = m.ExcPending
	if v7 != 0 {
		return int32(0)
	} else {
		return v4
	}
}
