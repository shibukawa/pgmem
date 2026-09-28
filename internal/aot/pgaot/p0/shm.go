package p0

import (
	base "github.com/shibukawa/pgmem/internal/aot/pgaot/base"
	"sync/atomic"
	"unsafe"
)

func F_shm_mq_set_sender(m *base.Module, l0 int32, l1 int32) {
	mBase := m.M
	_ = mBase
	var v5 int32
	_ = v5
	var v8 int32
	_ = v8
	var v10 int32
	_ = v10
	var v11 int32
	_ = v11
	var v15 int32
	_ = v15
	var v16 int32
	_ = v16
	var v19 int32
	_ = v19
	var v20 int32
	_ = v20
	var v23 int32
	_ = v23
	var v26 int32
	_ = v26
	var v27 int32
	_ = v27
	var v30 int32
	_ = v30
	var v34 int32
	_ = v34
	var v36 int32
	_ = v36
	var v38 int32
	_ = v38
	var v41 int32
	_ = v41
	var v44 int32
	_ = v44
	var v48 int32
	_ = v48
	var v52 int32
	_ = v52
	var v56 int32
	_ = v56
	var v64 int32
	_ = v64
	v5 = base.AtomicRmwXchg32(m, l0, int32(0), int32(1))
	if v5 != 0 {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	F_s_lock(m, l0, int32(_a_F_shm_mq_set_sender_0))
	mBase = m.M
	v8 = m.ExcPending
	if v8 != 0 {
		goto L4
	} else {
		goto L5
	}
L2:
	;
	goto L3
L3:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+8)) = l1
	v10 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v11 = int32(0)
	atomic.StoreUint32((*uint32)(unsafe.Add(mBase, uint32(l0))), uint32(v11))
	if v10 != 0 {
		goto L6
	} else {
		goto L7
	}
L4:
	;
	return
L5:
	;
	goto L3
L6:
	;
	v15 = v10 + int32(316)
	v16 = int32(0)
	v19 = base.AtomicRmwOr32(m, v16, int32(_a_F_shm_mq_set_sender_1), v16)
	v20 = *(*int32)(unsafe.Add(mBase, uint32(v15)))
	if v20 != 0 {
		goto L10
	} else {
		goto L11
	}
L7:
	;
	goto L8
L8:
	;
	return
L9:
	;
	goto L8
L10:
	;
	goto L9
L11:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v15))) = int32(1)
	v23 = int32(0)
	v26 = base.AtomicRmwOr32(m, v23, int32(_a_F_shm_mq_set_sender_1), v23)
	v27 = *(*int32)(unsafe.Add(mBase, uint32(v15)+4))
	if v27 == v23 {
		goto L10
	} else {
		goto L12
	}
L12:
	;
	v30 = *(*int32)(unsafe.Add(mBase, uint32(v15)+12))
	if v30 == int32(0) {
		goto L10
	} else {
		goto L13
	}
L13:
	;
	v34 = *(*int32)(unsafe.Add(mBase, _c_F_shm_mq_set_sender[0]))
	if v34 == v30 {
		goto L14
	} else {
		goto L15
	}
L14:
	;
	v36 = m.G0
	v38 = v36 - int32(16)
	m.G0 = v38
	v41 = *(*int32)(unsafe.Add(mBase, _c_F_shm_mq_set_sender[1]))
	if v41 == int32(0) {
		goto L17
	} else {
		goto L18
	}
L15:
	;
	goto L16
L16:
	;
	v64 = F_pgmem_kill(m, v30, int32(23))
	mBase = m.M
	goto L10
L17:
	;
	m.G0 = v38 + int32(16)
	goto L9
L18:
	;
	v44 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v38)+15)) = uint8(v44)
	goto L19
L19:
	;
	v48 = *(*int32)(unsafe.Add(mBase, _c_F_shm_mq_set_sender[2]))
	v52 = F_write(m, v48, v38+int32(15), int32(1))
	mBase = m.M
	if int32(0) <= v52 {
		goto L17
	} else {
		goto L21
	}
L20:
	;
	goto L17
L21:
	;
	v56 = *(*int32)(unsafe.Add(mBase, _c_F_shm_mq_set_sender[3]))
	if v56 == int32(27) {
		goto L19
	} else {
		goto L22
	}
L22:
	;
	goto L20
}
func F_shm_toc_lookup(m *base.Module, l0 int32, l1 int64, l2 int32) int32 {
	mBase := m.M
	_ = mBase
	var v4 int32
	_ = v4
	var v9 int32
	_ = v9
	var v11 int32
	_ = v11
	var v13 int32
	_ = v13
	var v17 int32
	_ = v17
	var v25 int32
	_ = v25
	var v32 int32
	_ = v32
	var v33 int64
	_ = v33
	var v36 int32
	_ = v36
	var v38 int32
	_ = v38
	var v54 int32
	_ = v54
	var v59 int32
	_ = v59
	var v64 int32
	_ = v64
	var v68 int32
	_ = v68
	v4 = int32(0)
	v9 = m.G0
	v11 = v9 - int32(16)
	m.G0 = v11
	v13 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v17 = base.AtomicRmwOr32(m, v4, int32(_a_F_shm_toc_lookup_0), v4)
	if v13 == v4 {
		goto L2
	} else {
		goto L3
	}
L1:
	;
	m.G0 = v11 + int32(16)
	return v68
L2:
	;
	if l2 != 0 {
		v68 = int32(0)
		goto L1
	} else {
		goto L10
	}
L3:
	;
	v25 = v4
	goto L4
L4:
	;
	v32 = l0 + int32(24) + v25<<(uint(int32(4))%32)
	v33 = *(*int64)(unsafe.Add(mBase, uint32(v32)))
	if l1 != v33 {
		goto L6
	} else {
		goto L7
	}
L5:
	;
	v38 = *(*int32)(unsafe.Add(mBase, uint32(v32)+8))
	v68 = l0 + v38
	goto L1
L6:
	;
	v36 = v25 + int32(1)
	if v13 != v36 {
		v25 = v36
		goto L4
	} else {
		goto L9
	}
L7:
	;
	goto L8
L8:
	;
	goto L5
L9:
	;
	goto L2
L10:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v54 = m.ExcPending
	if v54 != 0 {
		goto L11
	} else {
		goto L12
	}
L11:
	;
	return int32(0)
L12:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v11)+8)) = l0
	*(*int64)(unsafe.Add(mBase, uint32(v11))) = l1
	F_errmsg_internal(m, int32(_a_F_shm_toc_lookup_1), v11)
	mBase = m.M
	v59 = m.ExcPending
	if v59 != 0 {
		goto L11
	} else {
		goto L13
	}
L13:
	;
	F_errfinish(m, int32(_a_F_shm_toc_lookup_2), int32(261), int32(_a_F_shm_toc_lookup_3))
	mBase = m.M
	v64 = m.ExcPending
	if v64 != 0 {
		goto L11
	} else {
		goto L14
	}
L14:
	;
	base.Wasm_trap_unreachable()
	for {
	}
}
func F_shm_unlink(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v3 int32
	_ = v3
	var v5 int32
	_ = v5
	var v9 int32
	_ = v9
	var v15 int32
	_ = v15
	var v19 int32
	_ = v19
	var v21 int32
	_ = v21
	var v22 int32
	_ = v22
	var v25 int32
	_ = v25
	var v30 int32
	_ = v30
	var v45 int32
	_ = v45
	var v46 int32
	_ = v46
	var v51 int32
	_ = v51
	var v53 int32
	_ = v53
	var v54 int32
	_ = v54
	var v56 int32
	_ = v56
	v3 = m.G0
	v5 = v3 - int32(272)
	m.G0 = v5
	v9 = l0
	for {
		v15 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v9))))
		if v15 == int32(47) {
			v9 = v9 + int32(1)
			continue
		} else {
			break
		}
		break
	}
	v19 = F___strchrnul(m, v9, int32(47))
	mBase = m.M
	if v19 == v9 {
		*(*int32)(unsafe.Add(mBase, _c_F_shm_unlink[0])) = int32(28)
		v53 = int32(0)
	} else {
		v21 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v19))))
		if v21 != 0 {
			*(*int32)(unsafe.Add(mBase, _c_F_shm_unlink[0])) = int32(28)
			v53 = int32(0)
		} else {
			v22 = v19 - v9
			if int32(2) < v22 {
				if base.Ui32(v22) < base.Ui32(int32(256)) {
					v45 = int32(9)
					v46 = F___memcpy(m, v5, int32(_a_F_shm_unlink_0), v45)
					mBase = m.M
					v51 = F___memcpy(m, v5+v45, v9, v22+int32(1))
					mBase = m.M
					v53 = v5
				} else {
					*(*int32)(unsafe.Add(mBase, _c_F_shm_unlink[0])) = int32(37)
					v53 = int32(0)
				}
			} else {
				v25 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v9))))
				if v25 != int32(46) {
					v45 = int32(9)
					v46 = F___memcpy(m, v5, int32(_a_F_shm_unlink_0), v45)
					mBase = m.M
					v51 = F___memcpy(m, v5+v45, v9, v22+int32(1))
					mBase = m.M
					v53 = v5
				} else {
					v30 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v19-int32(1)))))
					if v30 != int32(46) {
						v45 = int32(9)
						v46 = F___memcpy(m, v5, int32(_a_F_shm_unlink_0), v45)
						mBase = m.M
						v51 = F___memcpy(m, v5+v45, v9, v22+int32(1))
						mBase = m.M
						v53 = v5
					} else {
						*(*int32)(unsafe.Add(mBase, _c_F_shm_unlink[0])) = int32(28)
						v53 = int32(0)
					}
				}
			}
		}
	}
	if v53 != 0 {
		v54 = F_unlink(m, v53)
		mBase = m.M
		v56 = v54
	} else {
		v56 = int32(-1)
	}
	m.G0 = v5 + int32(272)
	return v56
}
