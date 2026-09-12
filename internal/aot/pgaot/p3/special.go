package p3

import (
	base "github.com/shibukawa/pgmem/internal/aot/pgaot/base"
	"unsafe"
)

func F_EncodeSpecialTimestamp(m *base.Module, l0 int64, l1 int32) {
	mBase := m.M
	_ = mBase
	var v8 int32
	_ = v8
	var v11 int64
	_ = v11
	var v16 int32
	_ = v16
	var v20 int32
	_ = v20
	var v25 int32
	_ = v25
	var v27 int32
	_ = v27
	var v30 int64
	_ = v30
	if l0 != int64(9223372036854775807) {
		if l0 == int64(-9223372036854775807-1) {
			v8 = int32(*(*uint16)(unsafe.Add(mBase, _consts[1037])))
			*(*uint16)(unsafe.Add(mBase, uint32(l1)+8)) = uint16(v8)
			v11 = *(*int64)(unsafe.Add(mBase, _consts[1038]))
			*(*int64)(unsafe.Add(mBase, uint32(l1))) = v11
			return
		} else {
			F_errstart_cold(m, int32(21), int32(0))
			mBase = m.M
			v16 = m.ExcPending
			if v16 != 0 {
				return
			} else {
				F_errmsg_internal(m, int32(234954), int32(0))
				mBase = m.M
				v20 = m.ExcPending
				if v20 != 0 {
					return
				} else {
					F_errfinish(m, int32(490442), int32(1594), int32(234975))
					mBase = m.M
					v25 = m.ExcPending
					if v25 != 0 {
						return
					} else {
						base.Wasm_trap_unreachable()
						for {
						}
					}
				}
			}
		}
	} else {
		v27 = int32(*(*uint8)(unsafe.Add(mBase, _consts[1039])))
		*(*uint8)(unsafe.Add(mBase, uint32(l1)+8)) = uint8(v27)
		v30 = *(*int64)(unsafe.Add(mBase, _consts[1040]))
		*(*int64)(unsafe.Add(mBase, uint32(l1))) = v30
		return
	}
}
func F_assign_special_exec_param(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v4 int32
	_ = v4
	var v5 int32
	_ = v5
	var v6 int32
	_ = v6
	var v7 int32
	_ = v7
	var v9 int32
	_ = v9
	var v12 int32
	_ = v12
	var v13 int32
	_ = v13
	v4 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	v5 = *(*int32)(unsafe.Add(mBase, uint32(v4)+64))
	if v5 != 0 {
		v6 = *(*int32)(unsafe.Add(mBase, uint32(v5)+4))
		v7 = v6
	} else {
		v7 = int32(0)
	}
	v9 = F_lappend_oid(m, v5, int32(0))
	mBase = m.M
	v12 = m.ExcPending
	if v12 != 0 {
		return int32(0)
	} else {
		v13 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
		*(*int32)(unsafe.Add(mBase, uint32(v13)+64)) = v9
		return v7
	}
}
func F_get_special_variable(m *base.Module, l0 int32, l1 int32, l2 int32) {
	mBase := m.M
	_ = mBase
	var v4 int32
	_ = v4
	var v5 int32
	_ = v5
	var v10 int32
	_ = v10
	var v13 int32
	_ = v13
	var v14 int32
	_ = v14
	var v19 int32
	_ = v19
	v4 = *(*int32)(unsafe.Add(mBase, uint32(l1)))
	v5 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	if v5 != int32(6) {
		F_appendStringInfoChar(m, v4, int32(40))
		mBase = m.M
		v10 = m.ExcPending
		if v10 != 0 {
			return
		} else {
			F_get_rule_expr(m, l0, l1, int32(1))
			mBase = m.M
			v13 = m.ExcPending
			if v13 != 0 {
				return
			} else {
				v14 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
				if v14 != int32(6) {
					F_appendStringInfoChar(m, v4, int32(41))
					mBase = m.M
					v19 = m.ExcPending
					if v19 != 0 {
						return
					} else {
						return
					}
				} else {
					return
				}
			}
		}
	} else {
		F_get_rule_expr(m, l0, l1, int32(1))
		mBase = m.M
		v13 = m.ExcPending
		if v13 != 0 {
			return
		} else {
			v14 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
			if v14 != int32(6) {
				F_appendStringInfoChar(m, v4, int32(41))
				mBase = m.M
				v19 = m.ExcPending
				if v19 != 0 {
					return
				} else {
					return
				}
			} else {
				return
			}
		}
	}
}
