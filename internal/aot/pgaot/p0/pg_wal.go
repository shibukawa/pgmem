package p0

import (
	base "github.com/shibukawa/pgmem/internal/aot/pgaot/base"
	"unsafe"
)

func F_pg_get_wal_replay_pause_state(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v5 int32
	_ = v5
	var v10 int32
	_ = v10
	var v11 int32
	_ = v11
	var v13 int32
	_ = v13
	var v15 int32
	_ = v15
	var v17 int32
	_ = v17
	var v20 int32
	_ = v20
	var v27 int32
	_ = v27
	var v28 int32
	_ = v28
	var v29 int32
	_ = v29
	var v30 int32
	_ = v30
	var v35 int32
	_ = v35
	var v38 int32
	_ = v38
	var v42 int32
	_ = v42
	var v46 int32
	_ = v46
	var v51 int32
	_ = v51
	v5 = int32(*(*uint8)(unsafe.Add(mBase, _consts[2])))
	if v5 == int32(1) {
		v10 = *(*int32)(unsafe.Add(mBase, _consts[3]))
		v11 = *(*int32)(unsafe.Add(mBase, uint32(v10)+316))
		v13 = base.B2i32(v11 != int32(2))
		*(*uint8)(unsafe.Add(mBase, _consts[2])) = uint8(v13)
		v15 = v13
	} else {
		v15 = int32(0)
	}
	if v15 != 0 {
		v17 = F_GetRecoveryPauseState(m)
		mBase = m.M
		v20 = m.ExcPending
		if v20 != 0 {
			return int32(0)
		} else {
			if base.Ui32(v17) <= base.Ui32(int32(2)) {
				v27 = *(*int32)(unsafe.Add(mBase, uint32(v17<<(uint(int32(2))%32))+uint32(_consts[178])))
				v28 = v27
			} else {
				v28 = int32(0)
			}
			v29 = F_cstring_to_text(m, v28)
			mBase = m.M
			v30 = m.ExcPending
			if v30 != 0 {
				return int32(0)
			} else {
				return v29
			}
		}
	} else {
		F_errstart_cold(m, int32(21), int32(0))
		mBase = m.M
		v35 = m.ExcPending
		if v35 != 0 {
			return int32(0)
		} else {
			F_errcode(m, int32(325))
			mBase = m.M
			v38 = m.ExcPending
			if v38 != 0 {
				return int32(0)
			} else {
				F_errmsg(m, int32(136128), int32(0))
				mBase = m.M
				v42 = m.ExcPending
				if v42 != 0 {
					return int32(0)
				} else {
					F_errhint(m, int32(600672), int32(0))
					mBase = m.M
					v46 = m.ExcPending
					if v46 != 0 {
						return int32(0)
					} else {
						F_errfinish(m, int32(519185), int32(601), int32(369647))
						mBase = m.M
						v51 = m.ExcPending
						if v51 != 0 {
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
func F_pg_is_wal_replay_paused(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v4 int32
	_ = v4
	var v9 int32
	_ = v9
	var v10 int32
	_ = v10
	var v12 int32
	_ = v12
	var v14 int32
	_ = v14
	var v22 int32
	_ = v22
	var v25 int32
	_ = v25
	var v29 int32
	_ = v29
	var v33 int32
	_ = v33
	var v38 int32
	_ = v38
	var v39 int32
	_ = v39
	var v40 int32
	_ = v40
	v4 = int32(*(*uint8)(unsafe.Add(mBase, _consts[2])))
	if v4 == int32(1) {
		v9 = *(*int32)(unsafe.Add(mBase, _consts[3]))
		v10 = *(*int32)(unsafe.Add(mBase, uint32(v9)+316))
		v12 = base.B2i32(v10 != int32(2))
		*(*uint8)(unsafe.Add(mBase, _consts[2])) = uint8(v12)
		v14 = v12
	} else {
		v14 = int32(0)
	}
	if v14 == int32(0) {
		F_errstart_cold(m, int32(21), int32(0))
		mBase = m.M
		v22 = m.ExcPending
		if v22 != 0 {
			return int32(0)
		} else {
			F_errcode(m, int32(325))
			mBase = m.M
			v25 = m.ExcPending
			if v25 != 0 {
				return int32(0)
			} else {
				F_errmsg(m, int32(136128), int32(0))
				mBase = m.M
				v29 = m.ExcPending
				if v29 != 0 {
					return int32(0)
				} else {
					F_errhint(m, int32(600672), int32(0))
					mBase = m.M
					v33 = m.ExcPending
					if v33 != 0 {
						return int32(0)
					} else {
						F_errfinish(m, int32(519185), int32(578), int32(470314))
						mBase = m.M
						v38 = m.ExcPending
						if v38 != 0 {
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
	} else {
		v39 = F_GetRecoveryPauseState(m)
		mBase = m.M
		v40 = m.ExcPending
		if v40 != 0 {
			return int32(0)
		} else {
			return base.B2i32(v39 != int32(0))
		}
	}
}
func F_pg_wal_lsn_diff(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v4 int32
	_ = v4
	var v5 int32
	_ = v5
	var v6 int32
	_ = v6
	var v9 int32
	_ = v9
	v4 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v5 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
	v6 = F_DirectFunctionCall2Coll(m, int32(409), int32(0), v4, v5)
	mBase = m.M
	v9 = m.ExcPending
	if v9 != 0 {
		return int32(0)
	} else {
		return v6
	}
}
