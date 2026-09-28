package p0

import (
	base "github.com/shibukawa/pgmem/internal/aot/pgaot/base"
	"unsafe"
)

func F_pg_get_wal_replay_pause_state(m *base.Module, l0 int32) int64 {
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
	var v15 int32
	_ = v15
	var v18 int32
	_ = v18
	var v23 int32
	_ = v23
	var v25 int32
	_ = v25
	var v26 int32
	_ = v26
	var v27 int32
	_ = v27
	var v33 int32
	_ = v33
	var v36 int32
	_ = v36
	var v40 int32
	_ = v40
	var v44 int32
	_ = v44
	var v49 int32
	_ = v49
	v4 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_pg_get_wal_replay_pause_state[0])))
	if v4 == int32(1) {
		v9 = *(*int32)(unsafe.Add(mBase, _c_F_pg_get_wal_replay_pause_state[1]))
		v10 = *(*int32)(unsafe.Add(mBase, uint32(v9)+308))
		v12 = base.B2i32(v10 != int32(2))
		*(*uint8)(unsafe.Add(mBase, _c_F_pg_get_wal_replay_pause_state[0])) = uint8(v12)
		v14 = v12
	} else {
		v14 = int32(0)
	}
	if v14 != 0 {
		v15 = F_GetRecoveryPauseState(m)
		mBase = m.M
		v18 = m.ExcPending
		if v18 != 0 {
			return int64(0)
		} else {
			if base.Ui32(v15) <= base.Ui32(int32(2)) {
				v23 = *(*int32)(unsafe.Add(mBase, uint32(v15<<(uint(int32(2))%32))+uint32(_c_F_pg_get_wal_replay_pause_state[2])))
				v25 = v23
			} else {
				v25 = int32(0)
			}
			v26 = F_cstring_to_text(m, v25)
			mBase = m.M
			v27 = m.ExcPending
			if v27 != 0 {
				return int64(0)
			} else {
				return base.I64_extend_i32_u(v26)
			}
		}
	} else {
		F_errstart_cold(m, int32(21), int32(0))
		mBase = m.M
		v33 = m.ExcPending
		if v33 != 0 {
			return int64(0)
		} else {
			F_errcode(m, int32(325))
			mBase = m.M
			v36 = m.ExcPending
			if v36 != 0 {
				return int64(0)
			} else {
				F_errmsg(m, int32(_a_F_pg_get_wal_replay_pause_state_0), int32(0))
				mBase = m.M
				v40 = m.ExcPending
				if v40 != 0 {
					return int64(0)
				} else {
					F_errhint(m, int32(_a_F_pg_get_wal_replay_pause_state_1), int32(0))
					mBase = m.M
					v44 = m.ExcPending
					if v44 != 0 {
						return int64(0)
					} else {
						F_errfinish(m, int32(_a_F_pg_get_wal_replay_pause_state_2), int32(632), int32(_a_F_pg_get_wal_replay_pause_state_3))
						mBase = m.M
						v49 = m.ExcPending
						if v49 != 0 {
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
func F_pg_is_wal_replay_paused(m *base.Module, l0 int32) int64 {
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
	v4 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_pg_is_wal_replay_paused[0])))
	if v4 == int32(1) {
		v9 = *(*int32)(unsafe.Add(mBase, _c_F_pg_is_wal_replay_paused[1]))
		v10 = *(*int32)(unsafe.Add(mBase, uint32(v9)+308))
		v12 = base.B2i32(v10 != int32(2))
		*(*uint8)(unsafe.Add(mBase, _c_F_pg_is_wal_replay_paused[0])) = uint8(v12)
		v14 = v12
	} else {
		v14 = int32(0)
	}
	if v14 == int32(0) {
		F_errstart_cold(m, int32(21), int32(0))
		mBase = m.M
		v22 = m.ExcPending
		if v22 != 0 {
			return int64(0)
		} else {
			F_errcode(m, int32(325))
			mBase = m.M
			v25 = m.ExcPending
			if v25 != 0 {
				return int64(0)
			} else {
				F_errmsg(m, int32(_a_F_pg_is_wal_replay_paused_0), int32(0))
				mBase = m.M
				v29 = m.ExcPending
				if v29 != 0 {
					return int64(0)
				} else {
					F_errhint(m, int32(_a_F_pg_is_wal_replay_paused_1), int32(0))
					mBase = m.M
					v33 = m.ExcPending
					if v33 != 0 {
						return int64(0)
					} else {
						F_errfinish(m, int32(_a_F_pg_is_wal_replay_paused_2), int32(609), int32(_a_F_pg_is_wal_replay_paused_3))
						mBase = m.M
						v38 = m.ExcPending
						if v38 != 0 {
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
	} else {
		v39 = F_GetRecoveryPauseState(m)
		mBase = m.M
		v40 = m.ExcPending
		if v40 != 0 {
			return int64(0)
		} else {
			return base.I64_extend_i32_u(base.B2i32(v39 != int32(0)))
		}
	}
}
func F_pg_wal_lsn_diff(m *base.Module, l0 int32) int64 {
	mBase := m.M
	_ = mBase
	var v4 int64
	_ = v4
	var v5 int64
	_ = v5
	var v6 int64
	_ = v6
	var v9 int32
	_ = v9
	v4 = *(*int64)(unsafe.Add(mBase, uint32(l0)+24))
	v5 = *(*int64)(unsafe.Add(mBase, uint32(l0)+40))
	v6 = F_DirectFunctionCall2Coll(m, int32(435), int32(0), v4, v5)
	mBase = m.M
	v9 = m.ExcPending
	if v9 != 0 {
		return int64(0)
	} else {
		return v6
	}
}
