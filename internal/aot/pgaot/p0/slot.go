package p0

import (
	base "github.com/shibukawa/pgmem/internal/aot/pgaot/base"
	"unsafe"
)

func F_CheckSlotRequirements(m *base.Module) {
	mBase := m.M
	_ = mBase
	var v2 int32
	_ = v2
	var v4 int32
	_ = v4
	var v10 int32
	_ = v10
	var v13 int32
	_ = v13
	var v17 int32
	_ = v17
	var v22 int32
	_ = v22
	var v26 int32
	_ = v26
	var v29 int32
	_ = v29
	var v33 int32
	_ = v33
	var v38 int32
	_ = v38
	v2 = *(*int32)(unsafe.Add(mBase, _c_F_CheckSlotRequirements[0]))
	if v2 != 0 {
		v4 = *(*int32)(unsafe.Add(mBase, _c_F_CheckSlotRequirements[1]))
		if v4 <= int32(0) {
			F_errstart_cold(m, int32(21), int32(0))
			mBase = m.M
			v26 = m.ExcPending
			if v26 != 0 {
				return
			} else {
				F_errcode(m, int32(325))
				mBase = m.M
				v29 = m.ExcPending
				if v29 != 0 {
					return
				} else {
					F_errmsg(m, int32(_a_F_CheckSlotRequirements_0), int32(0))
					mBase = m.M
					v33 = m.ExcPending
					if v33 != 0 {
						return
					} else {
						F_errfinish(m, int32(_a_F_CheckSlotRequirements_1), int32(1539), int32(_a_F_CheckSlotRequirements_2))
						mBase = m.M
						v38 = m.ExcPending
						if v38 != 0 {
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
			return
		}
	} else {
		F_errstart_cold(m, int32(21), int32(0))
		mBase = m.M
		v10 = m.ExcPending
		if v10 != 0 {
			return
		} else {
			F_errcode(m, int32(325))
			mBase = m.M
			v13 = m.ExcPending
			if v13 != 0 {
				return
			} else {
				F_errmsg(m, int32(_a_F_CheckSlotRequirements_3), int32(0))
				mBase = m.M
				v17 = m.ExcPending
				if v17 != 0 {
					return
				} else {
					F_errfinish(m, int32(_a_F_CheckSlotRequirements_1), int32(1534), int32(_a_F_CheckSlotRequirements_2))
					mBase = m.M
					v22 = m.ExcPending
					if v22 != 0 {
						return
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
func F_ExecFetchSlotMinimalTuple(m *base.Module, l0 int32, l1 int32) int32 {
	mBase := m.M
	_ = mBase
	var v5 int32
	_ = v5
	var v6 int32
	_ = v6
	var v7 int32
	_ = v7
	var v9 int32
	_ = v9
	var v10 int32
	_ = v10
	var v11 int32
	_ = v11
	var v12 int32
	_ = v12
	var v15 int32
	_ = v15
	var v18 int32
	_ = v18
	var v20 int32
	_ = v20
	var v21 int32
	_ = v21
	var v22 int32
	_ = v22
	var v23 int32
	_ = v23
	var v24 int32
	_ = v24
	v5 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	v6 = *(*int32)(unsafe.Add(mBase, uint32(v5)+40))
	if v6 != 0 {
		if l1 != 0 {
			v7 = int32(0)
			*(*uint8)(unsafe.Add(mBase, uint32(l1))) = uint8(v7)
			v9 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
			v10 = *(*int32)(unsafe.Add(mBase, uint32(v9)+40))
			v11 = v10
		} else {
			v11 = v6
		}
		v12 = m.T0[v11].(func(*base.Module, int32) int32)(m, l0)
		mBase = m.M
		v15 = m.ExcPending
		if v15 != 0 {
			return int32(0)
		} else {
			return v12
		}
	} else {
		if l1 != 0 {
			v18 = int32(1)
			*(*uint8)(unsafe.Add(mBase, uint32(l1))) = uint8(v18)
			v20 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
			v21 = v20
		} else {
			v21 = v5
		}
		v22 = *(*int32)(unsafe.Add(mBase, uint32(v21)+48))
		v23 = m.T0[v22].(func(*base.Module, int32, int32) int32)(m, l0, int32(0))
		mBase = m.M
		v24 = m.ExcPending
		if v24 != 0 {
			return int32(0)
		} else {
			return v23
		}
	}
}
func F_SlotExistsInSyncStandbySlots(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v2 int32
	_ = v2
	var v8 int32
	_ = v8
	var v11 int32
	_ = v11
	var v17 int32
	_ = v17
	var v21 int32
	_ = v21
	var v24 int32
	_ = v24
	var v27 int32
	_ = v27
	var v30 int32
	_ = v30
	var v31 int32
	_ = v31
	var v34 int32
	_ = v34
	var v35 int32
	_ = v35
	var v38 int32
	_ = v38
	var v45 int32
	_ = v45
	var v46 int32
	_ = v46
	var v49 int32
	_ = v49
	var v52 int32
	_ = v52
	var v54 int32
	_ = v54
	var v57 int32
	_ = v57
	var v63 int32
	_ = v63
	v2 = int32(0)
	v8 = *(*int32)(unsafe.Add(mBase, _c_F_SlotExistsInSyncStandbySlots[0]))
	if v8 == v2 {
		v63 = v2
		goto L1
	} else {
		goto L2
	}
L1:
	;
	return v63
L2:
	;
	v11 = *(*int32)(unsafe.Add(mBase, uint32(v8)))
	if v11 <= int32(0) {
		v63 = v2
		goto L1
	} else {
		goto L3
	}
L3:
	;
	v17 = v8 + int32(4)
	v21 = v2
	goto L4
L4:
	;
	v24 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v17))))
	v27 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0))))
	if base.B2i32(v24 == int32(0))|base.B2i32(v24 != v27) != 0 {
		v45 = v24
		v46 = v27
		goto L7
	} else {
		goto L8
	}
L5:
	;
	v63 = v49
	goto L1
L6:
	;
	v49 = base.B2i32(v45-v46 == int32(0))
	if v45-v46 == int32(0) {
		v63 = v49
		goto L1
	} else {
		goto L13
	}
L7:
	;
	goto L6
L8:
	;
	v30 = v17
	v31 = l0
	goto L9
L9:
	;
	v34 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v31)+1)))
	v35 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v30)+1)))
	if v35 == int32(0) {
		v45 = v35
		v46 = v34
		goto L7
	} else {
		goto L11
	}
L10:
	;
	v45 = v35
	v46 = v34
	goto L7
L11:
	;
	v38 = int32(1)
	if v35 == v34 {
		v30 = v30 + v38
		v31 = v31 + v38
		goto L9
	} else {
		goto L12
	}
L12:
	;
	goto L10
L13:
	;
	v52 = F_strlen(m, v17)
	mBase = m.M
	v54 = int32(1)
	v57 = v21 + v54
	if v57 != v11 {
		v17 = v52 + v17 + v54
		v21 = v57
		goto L4
	} else {
		goto L14
	}
L14:
	;
	goto L5
}
