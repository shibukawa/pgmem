package p1

import (
	base "github.com/shibukawa/pgmem/internal/aot/pgaot/base"
	"unsafe"
)

func F_pg_get_wal_resource_managers(m *base.Module, l0 int32) int64 {
	mBase := m.M
	_ = mBase
	var v2 int32
	_ = v2
	var v5 int32
	_ = v5
	var v7 int32
	_ = v7
	var v9 int32
	_ = v9
	var v18 int32
	_ = v18
	var v22 int64
	_ = v22
	var v26 int32
	_ = v26
	var v28 int32
	_ = v28
	var v29 int32
	_ = v29
	var v36 int32
	_ = v36
	var v37 int32
	_ = v37
	var v43 int32
	_ = v43
	var v46 int64
	_ = v46
	v2 = int32(0)
	v5 = m.G0
	v7 = v5 - int32(48)
	m.G0 = v7
	v9 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	*(*uint8)(unsafe.Add(mBase, uint32(v7)+14)) = uint8(v2)
	*(*uint16)(unsafe.Add(mBase, uint32(v7)+12)) = uint16(v2)
	F_InitMaterializedSRF(m, l0, v2)
	mBase = m.M
	v18 = m.ExcPending
	if v18 != 0 {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	return int64(0)
L2:
	;
	v22 = int64(0)
	goto L3
L3:
	;
	v26 = *(*int32)(unsafe.Add(mBase, uint32(base.I32_wrap_i64(v22)<<(uint(int32(5))%32))+uint32(_c_F_pg_get_wal_resource_managers[0])))
	if v26 != 0 {
		goto L5
	} else {
		goto L6
	}
L4:
	;
	m.G0 = v7 + int32(48)
	return int64(0)
L5:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v7)+16)) = v22
	v28 = F_cstring_to_text(m, v26)
	mBase = m.M
	v29 = m.ExcPending
	if v29 != 0 {
		goto L1
	} else {
		goto L8
	}
L6:
	;
	goto L7
L7:
	;
	v46 = v22 + int64(1)
	if v46 != int64(256) {
		v22 = v46
		goto L3
	} else {
		goto L10
	}
L8:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v7)+32)) = base.I64_extend_i32_u(base.B2i32(base.Ui64(v22) < base.Ui64(int64(22))))
	*(*int64)(unsafe.Add(mBase, uint32(v7)+24)) = base.I64_extend_i32_u(v28)
	v36 = *(*int32)(unsafe.Add(mBase, uint32(v9)+24))
	v37 = *(*int32)(unsafe.Add(mBase, uint32(v9)+28))
	F_tuplestore_putvalues(m, v36, v37, v7+int32(16), v7+int32(12))
	mBase = m.M
	v43 = m.ExcPending
	if v43 != 0 {
		goto L1
	} else {
		goto L9
	}
L9:
	;
	goto L7
L10:
	;
	goto L4
}
func F_pg_wal_replay_pause(m *base.Module, l0 int32) int64 {
	mBase := m.M
	_ = mBase
	var v2 int32
	_ = v2
	var v4 int32
	_ = v4
	var v8 int32
	_ = v8
	var v13 int32
	_ = v13
	var v14 int32
	_ = v14
	var v16 int32
	_ = v16
	var v18 int32
	_ = v18
	var v19 int32
	_ = v19
	var v22 int32
	_ = v22
	var v25 int32
	_ = v25
	var v27 int32
	_ = v27
	var v39 int32
	_ = v39
	var v42 int32
	_ = v42
	var v46 int32
	_ = v46
	var v50 int32
	_ = v50
	var v55 int32
	_ = v55
	var v59 int32
	_ = v59
	var v62 int32
	_ = v62
	var v66 int32
	_ = v66
	var v71 int32
	_ = v71
	var v76 int32
	_ = v76
	v2 = m.G0
	v4 = v2 - int32(16)
	m.G0 = v4
	v8 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_pg_wal_replay_pause[0])))
	if v8 == int32(1) {
		v13 = *(*int32)(unsafe.Add(mBase, _c_F_pg_wal_replay_pause[1]))
		v14 = *(*int32)(unsafe.Add(mBase, uint32(v13)+308))
		v16 = base.B2i32(v14 != int32(2))
		*(*uint8)(unsafe.Add(mBase, _c_F_pg_wal_replay_pause[0])) = uint8(v16)
		v18 = v16
	} else {
		v18 = int32(0)
	}
	if v18 != 0 {
		v19 = F_PromoteIsTriggered(m)
		mBase = m.M
		v22 = m.ExcPending
		if v22 != 0 {
			return int64(0)
		} else {
			if v19 != 0 {
				F_errstart_cold(m, int32(21), int32(0))
				mBase = m.M
				v59 = m.ExcPending
				if v59 != 0 {
					return int64(0)
				} else {
					F_errcode(m, int32(325))
					mBase = m.M
					v62 = m.ExcPending
					if v62 != 0 {
						return int64(0)
					} else {
						F_errmsg(m, int32(_a_F_pg_wal_replay_pause_0), int32(0))
						mBase = m.M
						v66 = m.ExcPending
						if v66 != 0 {
							return int64(0)
						} else {
							*(*int32)(unsafe.Add(mBase, uint32(v4))) = int32(_a_F_pg_wal_replay_pause_1)
							F_errhint(m, int32(_a_F_pg_wal_replay_pause_2), v4)
							mBase = m.M
							v71 = m.ExcPending
							if v71 != 0 {
								return int64(0)
							} else {
								F_errfinish(m, int32(_a_F_pg_wal_replay_pause_3), int32(562), int32(_a_F_pg_wal_replay_pause_4))
								mBase = m.M
								v76 = m.ExcPending
								if v76 != 0 {
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
				F_SetRecoveryPause(m, int32(1))
				mBase = m.M
				v25 = m.ExcPending
				if v25 != 0 {
					return int64(0)
				} else {
					v27 = *(*int32)(unsafe.Add(mBase, _c_F_pg_wal_replay_pause[2]))
					F_SetLatch(m, v27+int32(4))
					mBase = m.M
					m.G0 = v4 + int32(16)
					return int64(0)
				}
			}
		}
	} else {
		F_errstart_cold(m, int32(21), int32(0))
		mBase = m.M
		v39 = m.ExcPending
		if v39 != 0 {
			return int64(0)
		} else {
			F_errcode(m, int32(325))
			mBase = m.M
			v42 = m.ExcPending
			if v42 != 0 {
				return int64(0)
			} else {
				F_errmsg(m, int32(_a_F_pg_wal_replay_pause_5), int32(0))
				mBase = m.M
				v46 = m.ExcPending
				if v46 != 0 {
					return int64(0)
				} else {
					F_errhint(m, int32(_a_F_pg_wal_replay_pause_6), int32(0))
					mBase = m.M
					v50 = m.ExcPending
					if v50 != 0 {
						return int64(0)
					} else {
						F_errfinish(m, int32(_a_F_pg_wal_replay_pause_3), int32(555), int32(_a_F_pg_wal_replay_pause_4))
						mBase = m.M
						v55 = m.ExcPending
						if v55 != 0 {
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
