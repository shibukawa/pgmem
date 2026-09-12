package p1

import (
	base "github.com/shibukawa/pgmem/internal/aot/pgaot/base"
	"unsafe"
)

func F_pa_decr_and_wait_stream_block(m *base.Module) {
	mBase := m.M
	_ = mBase
	var v4 int32
	_ = v4
	var v5 int32
	_ = v5
	var v8 int32
	_ = v8
	var v12 int32
	_ = v12
	var v17 int32
	_ = v17
	var v19 int32
	_ = v19
	var v22 int32
	_ = v22
	var v26 int32
	_ = v26
	var v30 int32
	_ = v30
	var v35 int32
	_ = v35
	var v36 int32
	_ = v36
	var v37 int32
	_ = v37
	var v43 int32
	_ = v43
	var v44 int32
	_ = v44
	var v46 int32
	_ = v46
	var v47 int32
	_ = v47
	var v51 int32
	_ = v51
	var v53 int32
	_ = v53
	var v54 int32
	_ = v54
	var v56 int32
	_ = v56
	var v57 int32
	_ = v57
	var v61 int32
	_ = v61
	v4 = *(*int32)(unsafe.Add(mBase, _consts[652]))
	v5 = *(*int32)(unsafe.Add(mBase, uint32(v4)+20))
	if v5 == int32(0) {
		v8 = *(*int32)(unsafe.Add(mBase, uint32(v4)))
		*(*int32)(unsafe.Add(mBase, uint32(v4))) = int32(1)
		if v8 != 0 {
			v12 = *(*int32)(unsafe.Add(mBase, _consts[652]))
			F_s_lock(m, v12, int32(518419), int32(1531), int32(368390))
			mBase = m.M
			v17 = m.ExcPending
			if v17 != 0 {
				return
			} else {
				v19 = *(*int32)(unsafe.Add(mBase, _consts[652]))
				*(*int32)(unsafe.Add(mBase, uint32(v19))) = int32(0)
				v22 = *(*int32)(unsafe.Add(mBase, uint32(v19)+32))
				if v22 != 0 {
					return
				} else {
					F_errstart_cold(m, int32(21), int32(0))
					mBase = m.M
					v26 = m.ExcPending
					if v26 != 0 {
						return
					} else {
						F_errmsg_internal(m, int32(596620), int32(0))
						mBase = m.M
						v30 = m.ExcPending
						if v30 != 0 {
							return
						} else {
							F_errfinish(m, int32(518419), int32(1611), int32(331213))
							mBase = m.M
							v35 = m.ExcPending
							if v35 != 0 {
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
		} else {
			v19 = *(*int32)(unsafe.Add(mBase, _consts[652]))
			*(*int32)(unsafe.Add(mBase, uint32(v19))) = int32(0)
			v22 = *(*int32)(unsafe.Add(mBase, uint32(v19)+32))
			if v22 != 0 {
				return
			} else {
				F_errstart_cold(m, int32(21), int32(0))
				mBase = m.M
				v26 = m.ExcPending
				if v26 != 0 {
					return
				} else {
					F_errmsg_internal(m, int32(596620), int32(0))
					mBase = m.M
					v30 = m.ExcPending
					if v30 != 0 {
						return
					} else {
						F_errfinish(m, int32(518419), int32(1611), int32(331213))
						mBase = m.M
						v35 = m.ExcPending
						if v35 != 0 {
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
	} else {
		v36 = *(*int32)(unsafe.Add(mBase, uint32(v4)+20))
		v37 = int32(1)
		*(*int32)(unsafe.Add(mBase, uint32(v4)+20)) = v36 - v37
		if v36 != v37 {
			return
		} else {
			v43 = *(*int32)(unsafe.Add(mBase, _consts[651]))
			v44 = *(*int32)(unsafe.Add(mBase, uint32(v43)+32))
			v46 = *(*int32)(unsafe.Add(mBase, _consts[652]))
			v47 = *(*int32)(unsafe.Add(mBase, uint32(v46)+4))
			F_LockApplyTransactionForSession(m, v44, v47, int32(0), int32(1))
			mBase = m.M
			v51 = m.ExcPending
			if v51 != 0 {
				return
			} else {
				v53 = *(*int32)(unsafe.Add(mBase, _consts[651]))
				v54 = *(*int32)(unsafe.Add(mBase, uint32(v53)+32))
				v56 = *(*int32)(unsafe.Add(mBase, _consts[652]))
				v57 = *(*int32)(unsafe.Add(mBase, uint32(v56)+4))
				F_UnlockApplyTransactionForSession(m, v54, v57, int32(0), int32(1))
				mBase = m.M
				v61 = m.ExcPending
				if v61 != 0 {
					return
				} else {
					return
				}
			}
		}
	}
}
func F_pa_send_data(m *base.Module, l0 int32, l1 int32, l2 int32) int32 {
	mBase := m.M
	_ = mBase
	var v8 int32
	_ = v8
	var v15 int64
	_ = v15
	var v17 int32
	_ = v17
	var v18 int32
	_ = v18
	var v21 int32
	_ = v21
	var v24 int32
	_ = v24
	var v28 int32
	_ = v28
	var v31 int32
	_ = v31
	var v35 int32
	_ = v35
	var v40 int32
	_ = v40
	var v42 int32
	_ = v42
	var v46 int32
	_ = v46
	var v47 int32
	_ = v47
	var v53 int32
	_ = v53
	var v57 int32
	_ = v57
	var v61 int32
	_ = v61
	var v65 int32
	_ = v65
	var v66 int32
	_ = v66
	var v67 int32
	_ = v67
	var v70 int64
	_ = v70
	var v71 int64
	_ = v71
	var v79 int64
	_ = v79
	var v94 int32
	_ = v94
	v8 = *(*int32)(unsafe.Add(mBase, _consts[650]))
	if v8 == int32(1) {
		v94 = int32(0)
		goto L1
	} else {
		goto L2
	}
L1:
	;
	return v94
L2:
	;
	v15 = int64(0)
	goto L3
L3:
	;
	v17 = int32(1)
	v18 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v21 = F_shm_mq_send(m, v18, l1, l2, v17, v17)
	mBase = m.M
	v24 = m.ExcPending
	if v24 != 0 {
		goto L7
	} else {
		goto L8
	}
L4:
	;
	v94 = int32(0)
	goto L1
L5:
	;
	v42 = *(*int32)(unsafe.Add(mBase, _consts[516]))
	v46 = F_WaitLatch(m, v42, int32(41), int32(1000), int32(134217757))
	mBase = m.M
	v47 = m.ExcPending
	if v47 != 0 {
		goto L7
	} else {
		goto L14
	}
L6:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v28 = m.ExcPending
	if v28 != 0 {
		goto L7
	} else {
		goto L9
	}
L7:
	;
	return int32(0)
L8:
	;
	switch v21 {
	case 0:
		v94 = v17
		goto L1
	default:
		goto L5
	case 2:
		goto L6
	}
L9:
	;
	F_errcode(m, int32(325))
	mBase = m.M
	v31 = m.ExcPending
	if v31 != 0 {
		goto L7
	} else {
		goto L10
	}
L10:
	;
	F_errmsg(m, int32(364490), int32(0))
	mBase = m.M
	v35 = m.ExcPending
	if v35 != 0 {
		goto L7
	} else {
		goto L11
	}
L11:
	;
	F_errfinish(m, int32(518419), int32(1187), int32(528563))
	mBase = m.M
	v40 = m.ExcPending
	if v40 != 0 {
		goto L7
	} else {
		goto L12
	}
L12:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L13:
	;
	v65 = m.G0
	v66 = int32(16)
	v67 = v65 - v66
	m.G0 = v67
	F___gettimeofday(m, v67)
	mBase = m.M
	v70 = *(*int64)(unsafe.Add(mBase, uint32(v67)))
	v71 = int64(*(*int32)(unsafe.Add(mBase, uint32(v67)+8)))
	m.G0 = v67 + v66
	v79 = v71 + v70*int64(1000000) - int64(946684800000000)
	goto L19
L14:
	;
	if v46&int32(1) == int32(0) {
		goto L13
	} else {
		goto L15
	}
L15:
	;
	v53 = *(*int32)(unsafe.Add(mBase, _consts[516]))
	*(*int32)(unsafe.Add(mBase, uint32(v53))) = int32(0)
	goto L16
L16:
	;
	v57 = *(*int32)(unsafe.Add(mBase, _consts[48]))
	if v57 == int32(0) {
		goto L13
	} else {
		goto L17
	}
L17:
	;
	F_ProcessInterrupts(m)
	mBase = m.M
	v61 = m.ExcPending
	if v61 != 0 {
		goto L7
	} else {
		goto L18
	}
L18:
	;
	goto L13
L19:
	;
	if v15 == int64(0) {
		goto L20
	} else {
		goto L21
	}
L20:
	;
	v15 = v79
	goto L3
L21:
	;
	goto L22
L22:
	;
	goto L23
L23:
	;
	if base.B2i32(base.I64_extend_i32_s(int32(9000))*int64(1000) <= v79-v15) == int32(0) {
		goto L3
	} else {
		goto L24
	}
L24:
	;
	goto L4
}
func F_pa_unlock_stream(m *base.Module, l0 int32) {
	mBase := m.M
	_ = mBase
	var v3 int32
	_ = v3
	var v4 int32
	_ = v4
	var v8 int32
	_ = v8
	v3 = *(*int32)(unsafe.Add(mBase, _consts[651]))
	v4 = *(*int32)(unsafe.Add(mBase, uint32(v3)+32))
	F_UnlockApplyTransactionForSession(m, v4, l0, int32(0), int32(8))
	mBase = m.M
	v8 = m.ExcPending
	if v8 != 0 {
		return
	} else {
		return
	}
}
