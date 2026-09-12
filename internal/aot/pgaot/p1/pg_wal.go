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
	var v27 int32
	_ = v27
	var v29 int32
	_ = v29
	var v30 int32
	_ = v30
	var v35 int32
	_ = v35
	var v36 int32
	_ = v36
	var v40 int32
	_ = v40
	var v43 int32
	_ = v43
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
	v27 = *(*int32)(unsafe.Add(mBase, uint32(v21<<(uint(int32(5))%32))+uint32(_consts[135])))
	if v27 != 0 {
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
	v29 = F_cstring_to_text(m, v27)
	mBase = m.M
	v30 = m.ExcPending
	if v30 != 0 {
		goto L1
	} else {
		goto L8
	}
L6:
	;
	goto L7
L7:
	;
	v43 = v21 + int32(1)
	if v43 != int32(256) {
		v21 = v43
		goto L3
	} else {
		goto L10
	}
L8:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v7)+12)) = base.B2i32(base.Ui32(v21) < base.Ui32(int32(22)))
	*(*int32)(unsafe.Add(mBase, uint32(v7)+8)) = v29
	v35 = *(*int32)(unsafe.Add(mBase, uint32(v9)+24))
	v36 = *(*int32)(unsafe.Add(mBase, uint32(v9)+28))
	F_tuplestore_putvalues(m, v35, v36, v7+int32(4), v7)
	mBase = m.M
	v40 = m.ExcPending
	if v40 != 0 {
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
	var v36 int32
	_ = v36
	var v39 int32
	_ = v39
	var v43 int32
	_ = v43
	var v47 int32
	_ = v47
	var v52 int32
	_ = v52
	var v56 int32
	_ = v56
	var v59 int32
	_ = v59
	var v63 int32
	_ = v63
	var v68 int32
	_ = v68
	var v73 int32
	_ = v73
	v2 = m.G0
	v4 = v2 - int32(16)
	m.G0 = v4
	v8 = int32(*(*uint8)(unsafe.Add(mBase, _consts[29])))
	if v8 == int32(1) {
		v13 = *(*int32)(unsafe.Add(mBase, _consts[30]))
		v14 = *(*int32)(unsafe.Add(mBase, uint32(v13)+316))
		v16 = base.B2i32(v14 != int32(2))
		*(*uint8)(unsafe.Add(mBase, _consts[29])) = uint8(v16)
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
				v56 = m.ExcPending
				if v56 != 0 {
					return int32(0)
				} else {
					F_errcode(m, int32(325))
					mBase = m.M
					v59 = m.ExcPending
					if v59 != 0 {
						return int32(0)
					} else {
						F_errmsg(m, int32(351597), int32(0))
						mBase = m.M
						v63 = m.ExcPending
						if v63 != 0 {
							return int32(0)
						} else {
							*(*int32)(unsafe.Add(mBase, uint32(v4))) = int32(715868)
							F_errhint(m, int32(678042), v4)
							mBase = m.M
							v68 = m.ExcPending
							if v68 != 0 {
								return int32(0)
							} else {
								F_errfinish(m, int32(519185), int32(531), int32(375910))
								mBase = m.M
								v73 = m.ExcPending
								if v73 != 0 {
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
					F_WakeupRecovery(m)
					mBase = m.M
					v27 = m.ExcPending
					if v27 != 0 {
						return int32(0)
					} else {
						m.G0 = v4 + int32(16)
						return int32(0)
					}
				}
			}
		}
	} else {
		F_errstart_cold(m, int32(21), int32(0))
		mBase = m.M
		v36 = m.ExcPending
		if v36 != 0 {
			return int32(0)
		} else {
			F_errcode(m, int32(325))
			mBase = m.M
			v39 = m.ExcPending
			if v39 != 0 {
				return int32(0)
			} else {
				F_errmsg(m, int32(136128), int32(0))
				mBase = m.M
				v43 = m.ExcPending
				if v43 != 0 {
					return int32(0)
				} else {
					F_errhint(m, int32(600672), int32(0))
					mBase = m.M
					v47 = m.ExcPending
					if v47 != 0 {
						return int32(0)
					} else {
						F_errfinish(m, int32(519185), int32(524), int32(375910))
						mBase = m.M
						v52 = m.ExcPending
						if v52 != 0 {
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
