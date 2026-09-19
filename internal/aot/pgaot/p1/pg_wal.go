package p1

import (
	base "github.com/shibukawa/pgmem/internal/aot/pgaot/base"
	"unsafe"
)

func F_pg_get_wal_resource_managers(m *base.Module, l0 int32) int32 {
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
	var v21 int32
	_ = v21
	var v25 int32
	_ = v25
	var v27 int32
	_ = v27
	var v28 int32
	_ = v28
	var v33 int32
	_ = v33
	var v34 int32
	_ = v34
	var v38 int32
	_ = v38
	var v41 int32
	_ = v41
	v2 = int32(0)
	v5 = m.G0
	v7 = v5 - int32(16)
	m.G0 = v7
	v9 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	*(*uint8)(unsafe.Add(mBase, uint32(v7)+2)) = uint8(v2)
	*(*uint16)(unsafe.Add(mBase, uint32(v7))) = uint16(v2)
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
	return int32(0)
L2:
	;
	v21 = v2
	goto L3
L3:
	;
	v25 = *(*int32)(unsafe.Add(mBase, uint32(v21<<(uint(int32(5))%32))+uint32(_c_F_pg_get_wal_resource_managers[0])))
	if v25 != 0 {
		goto L5
	} else {
		goto L6
	}
L4:
	;
	m.G0 = v7 + int32(16)
	return int32(0)
L5:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v7)+4)) = v21
	v27 = F_cstring_to_text(m, v25)
	mBase = m.M
	v28 = m.ExcPending
	if v28 != 0 {
		goto L1
	} else {
		goto L8
	}
L6:
	;
	goto L7
L7:
	;
	v41 = v21 + int32(1)
	if v41 != int32(256) {
		v21 = v41
		goto L3
	} else {
		goto L10
	}
L8:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v7)+12)) = base.B2i32(base.Ui32(v21) < base.Ui32(int32(22)))
	*(*int32)(unsafe.Add(mBase, uint32(v7)+8)) = v27
	v33 = *(*int32)(unsafe.Add(mBase, uint32(v9)+24))
	v34 = *(*int32)(unsafe.Add(mBase, uint32(v9)+28))
	F_tuplestore_putvalues(m, v33, v34, v7+int32(4), v7)
	mBase = m.M
	v38 = m.ExcPending
	if v38 != 0 {
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
func F_pg_wal_replay_pause(m *base.Module, l0 int32) int32 {
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
		v14 = *(*int32)(unsafe.Add(mBase, uint32(v13)+316))
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
			return int32(0)
		} else {
			if v19 != 0 {
				F_errstart_cold(m, int32(21), int32(0))
				mBase = m.M
				v59 = m.ExcPending
				if v59 != 0 {
					return int32(0)
				} else {
					F_errcode(m, int32(325))
					mBase = m.M
					v62 = m.ExcPending
					if v62 != 0 {
						return int32(0)
					} else {
						F_errmsg(m, int32(_a_F_pg_wal_replay_pause_0), int32(0))
						mBase = m.M
						v66 = m.ExcPending
						if v66 != 0 {
							return int32(0)
						} else {
							*(*int32)(unsafe.Add(mBase, uint32(v4))) = int32(_a_F_pg_wal_replay_pause_1)
							F_errhint(m, int32(_a_F_pg_wal_replay_pause_2), v4)
							mBase = m.M
							v71 = m.ExcPending
							if v71 != 0 {
								return int32(0)
							} else {
								F_errfinish(m, int32(_a_F_pg_wal_replay_pause_3), int32(531), int32(_a_F_pg_wal_replay_pause_4))
								mBase = m.M
								v76 = m.ExcPending
								if v76 != 0 {
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
				F_SetRecoveryPause(m, int32(1))
				mBase = m.M
				v25 = m.ExcPending
				if v25 != 0 {
					return int32(0)
				} else {
					v27 = *(*int32)(unsafe.Add(mBase, _c_F_pg_wal_replay_pause[2]))
					F_SetLatch(m, v27+int32(4))
					mBase = m.M
					m.G0 = v4 + int32(16)
					return int32(0)
				}
			}
		}
	} else {
		F_errstart_cold(m, int32(21), int32(0))
		mBase = m.M
		v39 = m.ExcPending
		if v39 != 0 {
			return int32(0)
		} else {
			F_errcode(m, int32(325))
			mBase = m.M
			v42 = m.ExcPending
			if v42 != 0 {
				return int32(0)
			} else {
				F_errmsg(m, int32(_a_F_pg_wal_replay_pause_5), int32(0))
				mBase = m.M
				v46 = m.ExcPending
				if v46 != 0 {
					return int32(0)
				} else {
					F_errhint(m, int32(_a_F_pg_wal_replay_pause_6), int32(0))
					mBase = m.M
					v50 = m.ExcPending
					if v50 != 0 {
						return int32(0)
					} else {
						F_errfinish(m, int32(_a_F_pg_wal_replay_pause_3), int32(524), int32(_a_F_pg_wal_replay_pause_4))
						mBase = m.M
						v55 = m.ExcPending
						if v55 != 0 {
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
