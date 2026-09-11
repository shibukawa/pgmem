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
	v2 = *(*int32)(unsafe.Add(mBase, _consts[543]))
	if v2 != 0 {
		v4 = *(*int32)(unsafe.Add(mBase, _consts[8]))
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
					F_errmsg(m, int32(679250), int32(0))
					mBase = m.M
					v33 = m.ExcPending
					if v33 != 0 {
						return
					} else {
						F_errfinish(m, int32(462786), int32(1539), int32(113667))
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
				F_errmsg(m, int32(522168), int32(0))
				mBase = m.M
				v17 = m.ExcPending
				if v17 != 0 {
					return
				} else {
					F_errfinish(m, int32(462786), int32(1534), int32(113667))
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
	var v19 int32
	_ = v19
	var v24 int32
	_ = v24
	var v25 int32
	_ = v25
	var v29 int32
	_ = v29
	var v30 int32
	_ = v30
	var v33 int32
	_ = v33
	var v34 int32
	_ = v34
	var v37 int32
	_ = v37
	var v44 int32
	_ = v44
	var v45 int32
	_ = v45
	var v48 int32
	_ = v48
	var v58 int32
	_ = v58
	var v63 int32
	_ = v63
	var v67 int32
	_ = v67
	var v72 int32
	_ = v72
	var v74 int32
	_ = v74
	var v78 int32
	_ = v78
	var v84 int32
	_ = v84
	var v87 int32
	_ = v87
	var v93 int32
	_ = v93
	var v97 int32
	_ = v97
	var v99 int32
	_ = v99
	var v107 int32
	_ = v107
	var v109 int32
	_ = v109
	var v112 int32
	_ = v112
	var v116 int32
	_ = v116
	v2 = int32(0)
	v8 = *(*int32)(unsafe.Add(mBase, _consts[545]))
	if v8 == v2 {
		v116 = v2
		goto L1
	} else {
		goto L2
	}
L1:
	;
	return v116
L2:
	;
	v11 = *(*int32)(unsafe.Add(mBase, uint32(v8)))
	if v11 <= int32(0) {
		v116 = v2
		goto L1
	} else {
		goto L3
	}
L3:
	;
	v17 = v8 + int32(4)
	v19 = v2
	goto L4
L4:
	;
	v24 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0))))
	v25 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v17))))
	if v25 == int32(0) {
		v44 = v24
		v45 = v25
		goto L7
	} else {
		goto L8
	}
L5:
	;
	v116 = v48
	goto L1
L6:
	;
	v48 = base.B2i32(v45-v44 == int32(0))
	if v45-v44 == int32(0) {
		v116 = v48
		goto L1
	} else {
		goto L14
	}
L7:
	;
	goto L6
L8:
	;
	if v24 != v25 {
		v44 = v24
		v45 = v25
		goto L7
	} else {
		goto L9
	}
L9:
	;
	v29 = v17
	v30 = l0
	goto L10
L10:
	;
	v33 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v30)+1)))
	v34 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v29)+1)))
	if v34 == int32(0) {
		v44 = v33
		v45 = v34
		goto L7
	} else {
		goto L12
	}
L11:
	;
	v44 = v33
	v45 = v34
	goto L7
L12:
	;
	v37 = int32(1)
	if v33 == v34 {
		v29 = v29 + v37
		v30 = v30 + v37
		goto L10
	} else {
		goto L13
	}
L13:
	;
	goto L11
L14:
	;
	if v17&int32(3) == int32(0) {
		v74 = v17
		goto L17
	} else {
		goto L18
	}
L15:
	;
	v109 = int32(1)
	v112 = v19 + v109
	if v112 != v11 {
		v17 = v107 + v17 + v109
		v19 = v112
		goto L4
	} else {
		goto L32
	}
L16:
	;
	v107 = v99 - v17
	goto L15
L17:
	;
	v78 = v74
	goto L26
L18:
	;
	v58 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v17))))
	if v58 == int32(0) {
		goto L19
	} else {
		goto L20
	}
L19:
	;
	v107 = int32(0)
	goto L15
L20:
	;
	goto L21
L21:
	;
	v63 = v17
	goto L22
L22:
	;
	v67 = v63 + int32(1)
	if v67&int32(3) == int32(0) {
		v74 = v67
		goto L17
	} else {
		goto L24
	}
L23:
	;
	v99 = v67
	goto L16
L24:
	;
	v72 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v67))))
	if v72 != 0 {
		v63 = v67
		goto L22
	} else {
		goto L25
	}
L25:
	;
	goto L23
L26:
	;
	v84 = *(*int32)(unsafe.Add(mBase, uint32(v78)))
	v87 = int32(-2139062144)
	if (int32(16843008)-v84|v84)&v87 == v87 {
		v78 = v78 + int32(4)
		goto L26
	} else {
		goto L28
	}
L27:
	;
	v93 = v78
	goto L29
L28:
	;
	goto L27
L29:
	;
	v97 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v93))))
	if v97 != 0 {
		v93 = v93 + int32(1)
		goto L29
	} else {
		goto L31
	}
L30:
	;
	v99 = v93
	goto L16
L31:
	;
	goto L30
L32:
	;
	goto L5
}
