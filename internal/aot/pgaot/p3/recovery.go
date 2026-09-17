package p3

import (
	base "github.com/shibukawa/pgmem/internal/aot/pgaot/base"
	"unsafe"
)

func F_ResolveRecoveryConflictWithSnapshot(m *base.Module, l0 int32, l1 int32, l2 int32) {
	mBase := m.M
	_ = mBase
	var v6 int32
	_ = v6
	var v7 int32
	_ = v7
	var v8 int32
	_ = v8
	var v13 int32
	_ = v13
	var v17 int32
	_ = v17
	var v22 int32
	_ = v22
	var v23 int32
	_ = v23
	if l0 == int32(0) {
		return
	} else {
		v6 = *(*int32)(unsafe.Add(mBase, uint32(l2)+4))
		v7 = F_GetConflictingVirtualXIDs(m, l0, v6)
		mBase = m.M
		v8 = m.ExcPending
		if v8 != 0 {
			return
		} else {
			F_ResolveRecoveryConflictWithVirtualXIDs(m, v7, int32(10), int32(134217772), int32(1))
			mBase = m.M
			v13 = m.ExcPending
			if v13 != 0 {
				return
			} else {
				if l1 == int32(0) {
					return
				} else {
					v17 = *(*int32)(unsafe.Add(mBase, _c_F_ResolveRecoveryConflictWithSnapshot[0]))
					if v17 < int32(2) {
						return
					} else {
						v22 = F_InvalidateObsoleteReplicationSlots(m, int32(2), int64(0), v6, l0)
						mBase = m.M
						v23 = m.ExcPending
						if v23 != 0 {
							return
						} else {
							return
						}
					}
				}
			}
		}
	}
}
func F_assign_recovery_target_lsn(m *base.Module, l0 int32, l1 int32) {
	mBase := m.M
	_ = mBase
	var v4 int32
	_ = v4
	var v11 int32
	_ = v11
	var v18 int64
	_ = v18
	var v24 int32
	_ = v24
	v4 = *(*int32)(unsafe.Add(mBase, _c_F_assign_recovery_target_lsn[0]))
	if v4&int32(-5) == int32(0) {
		if l0 == int32(0) {
			*(*int32)(unsafe.Add(mBase, _c_F_assign_recovery_target_lsn[0])) = int32(0)
			return
		} else {
			v11 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0))))
			if v11 == int32(0) {
				*(*int32)(unsafe.Add(mBase, _c_F_assign_recovery_target_lsn[0])) = int32(0)
				return
			} else {
				*(*int32)(unsafe.Add(mBase, _c_F_assign_recovery_target_lsn[0])) = int32(4)
				v18 = *(*int64)(unsafe.Add(mBase, uint32(l1)))
				*(*int64)(unsafe.Add(mBase, _c_F_assign_recovery_target_lsn[1])) = v18
				return
			}
		}
	} else {
		F_error_multiple_recovery_targets(m)
		mBase = m.M
		v24 = m.ExcPending
		if v24 != 0 {
			return
		} else {
			base.Wasm_trap_unreachable()
			for {
			}
		}
	}
}
func F_assign_recovery_target_time(m *base.Module, l0 int32, l1 int32) {
	mBase := m.M
	_ = mBase
	var v4 int32
	_ = v4
	var v11 int32
	_ = v11
	var v13 int32
	_ = v13
	var v16 int32
	_ = v16
	v4 = *(*int32)(unsafe.Add(mBase, _c_F_assign_recovery_target_time[0]))
	if v4&int32(-3) == int32(0) {
		if l0 != 0 {
			v11 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0))))
			if v11 != 0 {
				v13 = int32(2)
			} else {
				v13 = int32(0)
			}
		} else {
			v13 = int32(0)
		}
		*(*int32)(unsafe.Add(mBase, _c_F_assign_recovery_target_time[0])) = v13
		return
	} else {
		F_error_multiple_recovery_targets(m)
		mBase = m.M
		v16 = m.ExcPending
		if v16 != 0 {
			return
		} else {
			base.Wasm_trap_unreachable()
			for {
			}
		}
	}
}
func F_assign_recovery_target_timeline(m *base.Module, l0 int32, l1 int32) {
	mBase := m.M
	_ = mBase
	var v4 int32
	_ = v4
	var v9 int32
	_ = v9
	var v12 int64
	_ = v12
	var v15 int32
	_ = v15
	v4 = *(*int32)(unsafe.Add(mBase, uint32(l1)))
	*(*int32)(unsafe.Add(mBase, _c_F_assign_recovery_target_timeline[0])) = v4
	if v4 == int32(2) {
		v9 = int32(0)
		v12 = F_strtox_2(m, l0, v9, v9, int64(4294967295))
		mBase = m.M
		v15 = base.I32_wrap_i64(v12)
	} else {
		v15 = int32(0)
	}
	*(*int32)(unsafe.Add(mBase, _c_F_assign_recovery_target_timeline[1])) = v15
	return
}
