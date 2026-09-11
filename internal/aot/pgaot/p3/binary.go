package p3

import (
	base "github.com/shibukawa/pgmem/internal/aot/pgaot/base"
	"unsafe"
)

func F_binary_upgrade_set_next_pg_tablespace_oid(m *base.Module, l0 int32) int32 {
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
	var v25 int32
	_ = v25
	v3 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_binary_upgrade_set_next_pg_tablespace_oid[0])))
	if v3 == int32(0) {
		F_errstart_cold(m, int32(21), int32(0))
		mBase = m.M
		v11 = m.ExcPending
		if v11 != 0 {
			return int32(0)
		} else {
			F_errcode(m, int32(33685829))
			mBase = m.M
			v14 = m.ExcPending
			if v14 != 0 {
				return int32(0)
			} else {
				F_errmsg(m, int32(_a_F_binary_upgrade_set_next_pg_tablespace_oid_0), int32(0))
				mBase = m.M
				v18 = m.ExcPending
				if v18 != 0 {
					return int32(0)
				} else {
					F_errfinish(m, int32(_a_F_binary_upgrade_set_next_pg_tablespace_oid_1), int32(46), int32(_a_F_binary_upgrade_set_next_pg_tablespace_oid_2))
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
		v25 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
		*(*int32)(unsafe.Add(mBase, _c_F_binary_upgrade_set_next_pg_tablespace_oid[1])) = v25
		return int32(0)
	}
}
func F_binary_upgrade_set_next_pg_type_oid(m *base.Module, l0 int32) int32 {
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
	var v25 int32
	_ = v25
	v3 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_binary_upgrade_set_next_pg_type_oid[0])))
	if v3 == int32(0) {
		F_errstart_cold(m, int32(21), int32(0))
		mBase = m.M
		v11 = m.ExcPending
		if v11 != 0 {
			return int32(0)
		} else {
			F_errcode(m, int32(33685829))
			mBase = m.M
			v14 = m.ExcPending
			if v14 != 0 {
				return int32(0)
			} else {
				F_errmsg(m, int32(_a_F_binary_upgrade_set_next_pg_type_oid_0), int32(0))
				mBase = m.M
				v18 = m.ExcPending
				if v18 != 0 {
					return int32(0)
				} else {
					F_errfinish(m, int32(_a_F_binary_upgrade_set_next_pg_type_oid_1), int32(57), int32(_a_F_binary_upgrade_set_next_pg_type_oid_2))
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
		v25 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
		*(*int32)(unsafe.Add(mBase, _c_F_binary_upgrade_set_next_pg_type_oid[1])) = v25
		return int32(0)
	}
}
func F_binary_upgrade_set_next_toast_relfilenode(m *base.Module, l0 int32) int32 {
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
	var v25 int32
	_ = v25
	v3 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_binary_upgrade_set_next_toast_relfilenode[0])))
	if v3 == int32(0) {
		F_errstart_cold(m, int32(21), int32(0))
		mBase = m.M
		v11 = m.ExcPending
		if v11 != 0 {
			return int32(0)
		} else {
			F_errcode(m, int32(33685829))
			mBase = m.M
			v14 = m.ExcPending
			if v14 != 0 {
				return int32(0)
			} else {
				F_errmsg(m, int32(_a_F_binary_upgrade_set_next_toast_relfilenode_0), int32(0))
				mBase = m.M
				v18 = m.ExcPending
				if v18 != 0 {
					return int32(0)
				} else {
					F_errfinish(m, int32(_a_F_binary_upgrade_set_next_toast_relfilenode_1), int32(156), int32(_a_F_binary_upgrade_set_next_toast_relfilenode_2))
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
		v25 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
		*(*int32)(unsafe.Add(mBase, _c_F_binary_upgrade_set_next_toast_relfilenode[1])) = v25
		return int32(0)
	}
}
