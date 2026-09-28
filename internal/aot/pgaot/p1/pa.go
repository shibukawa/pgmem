package p1

import (
	base "github.com/shibukawa/pgmem/internal/aot/pgaot/base"
	"sync/atomic"
	"unsafe"
)

func F_pa_decr_and_wait_stream_block(m *base.Module) {
	mBase := m.M
	_ = mBase
	var v4 int32
	_ = v4
	var v5 int32
	_ = v5
	var v10 int32
	_ = v10
	var v13 int32
	_ = v13
	var v15 int32
	_ = v15
	var v16 int32
	_ = v16
	var v17 int32
	_ = v17
	var v23 int32
	_ = v23
	var v27 int32
	_ = v27
	var v32 int32
	_ = v32
	var v33 int32
	_ = v33
	var v35 int32
	_ = v35
	var v39 int32
	_ = v39
	var v40 int32
	_ = v40
	var v42 int32
	_ = v42
	var v43 int32
	_ = v43
	var v47 int32
	_ = v47
	var v49 int32
	_ = v49
	var v50 int32
	_ = v50
	var v52 int32
	_ = v52
	var v53 int32
	_ = v53
	var v57 int32
	_ = v57
	v4 = *(*int32)(unsafe.Add(mBase, _c_F_pa_decr_and_wait_stream_block[0]))
	v5 = *(*int32)(unsafe.Add(mBase, uint32(v4)+20))
	if v5 == int32(0) {
		v10 = base.AtomicRmwXchg32(m, v4, int32(0), int32(1))
		if v10 != 0 {
			F_s_lock(m, v4, int32(_a_F_pa_decr_and_wait_stream_block_0))
			mBase = m.M
			v13 = m.ExcPending
			if v13 != 0 {
				return
			} else {
				v15 = *(*int32)(unsafe.Add(mBase, _c_F_pa_decr_and_wait_stream_block[0]))
				v16 = *(*int32)(unsafe.Add(mBase, uint32(v15)+32))
				v17 = int32(0)
				atomic.StoreUint32((*uint32)(unsafe.Add(mBase, uint32(v15))), uint32(v17))
				if v16 != 0 {
					return
				} else {
					F_errstart_cold(m, int32(21), int32(0))
					mBase = m.M
					v23 = m.ExcPending
					if v23 != 0 {
						return
					} else {
						F_errmsg_internal(m, int32(_a_F_pa_decr_and_wait_stream_block_1), int32(0))
						mBase = m.M
						v27 = m.ExcPending
						if v27 != 0 {
							return
						} else {
							F_errfinish(m, int32(_a_F_pa_decr_and_wait_stream_block_2), int32(1623), int32(_a_F_pa_decr_and_wait_stream_block_3))
							mBase = m.M
							v32 = m.ExcPending
							if v32 != 0 {
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
			v15 = *(*int32)(unsafe.Add(mBase, _c_F_pa_decr_and_wait_stream_block[0]))
			v16 = *(*int32)(unsafe.Add(mBase, uint32(v15)+32))
			v17 = int32(0)
			atomic.StoreUint32((*uint32)(unsafe.Add(mBase, uint32(v15))), uint32(v17))
			if v16 != 0 {
				return
			} else {
				F_errstart_cold(m, int32(21), int32(0))
				mBase = m.M
				v23 = m.ExcPending
				if v23 != 0 {
					return
				} else {
					F_errmsg_internal(m, int32(_a_F_pa_decr_and_wait_stream_block_1), int32(0))
					mBase = m.M
					v27 = m.ExcPending
					if v27 != 0 {
						return
					} else {
						F_errfinish(m, int32(_a_F_pa_decr_and_wait_stream_block_2), int32(1623), int32(_a_F_pa_decr_and_wait_stream_block_3))
						mBase = m.M
						v32 = m.ExcPending
						if v32 != 0 {
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
		v33 = int32(1)
		v35 = base.AtomicRmwSub32(m, v4, int32(20), v33)
		if v35 != v33 {
			return
		} else {
			v39 = *(*int32)(unsafe.Add(mBase, _c_F_pa_decr_and_wait_stream_block[1]))
			v40 = *(*int32)(unsafe.Add(mBase, uint32(v39)+32))
			v42 = *(*int32)(unsafe.Add(mBase, _c_F_pa_decr_and_wait_stream_block[0]))
			v43 = *(*int32)(unsafe.Add(mBase, uint32(v42)+4))
			F_LockApplyTransactionForSession(m, v40, v43, int32(0), int32(1))
			mBase = m.M
			v47 = m.ExcPending
			if v47 != 0 {
				return
			} else {
				v49 = *(*int32)(unsafe.Add(mBase, _c_F_pa_decr_and_wait_stream_block[1]))
				v50 = *(*int32)(unsafe.Add(mBase, uint32(v49)+32))
				v52 = *(*int32)(unsafe.Add(mBase, _c_F_pa_decr_and_wait_stream_block[0]))
				v53 = *(*int32)(unsafe.Add(mBase, uint32(v52)+4))
				F_UnlockApplyTransactionForSession(m, v50, v53, int32(0), int32(1))
				mBase = m.M
				v57 = m.ExcPending
				if v57 != 0 {
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
	var v16 int64
	_ = v16
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
	var v54 int32
	_ = v54
	var v59 int32
	_ = v59
	var v61 int32
	_ = v61
	var v65 int32
	_ = v65
	var v69 int32
	_ = v69
	var v70 int32
	_ = v70
	var v71 int32
	_ = v71
	var v74 int64
	_ = v74
	var v75 int64
	_ = v75
	var v83 int64
	_ = v83
	var v98 int32
	_ = v98
	v8 = *(*int32)(unsafe.Add(mBase, _c_F_pa_send_data[0]))
	if v8 == int32(1) {
		v98 = int32(0)
		goto L1
	} else {
		goto L2
	}
L1:
	;
	return v98
L2:
	;
	v16 = int64(0)
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
	v98 = int32(0)
	goto L1
L5:
	;
	v42 = *(*int32)(unsafe.Add(mBase, _c_F_pa_send_data[1]))
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
		v98 = v17
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
	F_errmsg(m, int32(_a_F_pa_send_data_0), int32(0))
	mBase = m.M
	v35 = m.ExcPending
	if v35 != 0 {
		goto L7
	} else {
		goto L11
	}
L11:
	;
	F_errfinish(m, int32(_a_F_pa_send_data_1), int32(1199), int32(_a_F_pa_send_data_2))
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
	v69 = m.G0
	v70 = int32(16)
	v71 = v69 - v70
	m.G0 = v71
	F_gettimeofday(m, v71)
	mBase = m.M
	v74 = *(*int64)(unsafe.Add(mBase, uint32(v71)))
	v75 = int64(*(*int32)(unsafe.Add(mBase, uint32(v71)+8)))
	m.G0 = v71 + v70
	v83 = v75 + v74*int64(1000000) - int64(946684800000000)
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
	v53 = *(*int32)(unsafe.Add(mBase, _c_F_pa_send_data[1]))
	v54 = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v53))) = v54
	v59 = base.AtomicRmwOr32(m, v54, int32(_a_F_pa_send_data_3), v54)
	goto L16
L16:
	;
	v61 = *(*int32)(unsafe.Add(mBase, _c_F_pa_send_data[2]))
	if v61 == int32(0) {
		goto L13
	} else {
		goto L17
	}
L17:
	;
	F_ProcessInterrupts(m)
	mBase = m.M
	v65 = m.ExcPending
	if v65 != 0 {
		goto L7
	} else {
		goto L18
	}
L18:
	;
	goto L13
L19:
	;
	if v16 == int64(0) {
		goto L20
	} else {
		goto L21
	}
L20:
	;
	v16 = v83
	goto L3
L21:
	;
	goto L22
L22:
	;
	goto L23
L23:
	;
	if base.B2i32(base.I64_extend_i32_s(int32(_a_F_pa_send_data_4))*int64(1000) <= v83-v16) == int32(0) {
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
	v3 = *(*int32)(unsafe.Add(mBase, _c_F_pa_unlock_stream[0]))
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
