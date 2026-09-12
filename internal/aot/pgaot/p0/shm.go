package p0

import (
	base "github.com/shibukawa/pgmem/internal/aot/pgaot/base"
	"unsafe"
)

func F_shm_mq_set_sender(m *base.Module, l0 int32, l1 int32) {
	mBase := m.M
	_ = mBase
	var v4 int32
	_ = v4
	var v11 int32
	_ = v11
	var v15 int32
	_ = v15
	var v19 int32
	_ = v19
	v4 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	*(*int32)(unsafe.Add(mBase, uint32(l0))) = int32(1)
	if v4 != 0 {
		F_s_lock(m, l0, int32(495018), int32(228), int32(227044))
		mBase = m.M
		v11 = m.ExcPending
		if v11 != 0 {
			return
		} else {
			*(*int32)(unsafe.Add(mBase, uint32(l0)+8)) = l1
			*(*int32)(unsafe.Add(mBase, uint32(l0))) = int32(0)
			v15 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
			if v15 != 0 {
				F_SetLatch(m, v15+int32(20))
				mBase = m.M
				v19 = m.ExcPending
				if v19 != 0 {
					return
				} else {
					return
				}
			} else {
				return
			}
		}
	} else {
		*(*int32)(unsafe.Add(mBase, uint32(l0)+8)) = l1
		*(*int32)(unsafe.Add(mBase, uint32(l0))) = int32(0)
		v15 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
		if v15 != 0 {
			F_SetLatch(m, v15+int32(20))
			mBase = m.M
			v19 = m.ExcPending
			if v19 != 0 {
				return
			} else {
				return
			}
		} else {
			return
		}
	}
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
	var v21 int32
	_ = v21
	var v28 int32
	_ = v28
	var v29 int64
	_ = v29
	var v32 int32
	_ = v32
	var v34 int32
	_ = v34
	var v50 int32
	_ = v50
	var v55 int32
	_ = v55
	var v60 int32
	_ = v60
	var v64 int32
	_ = v64
	v4 = int32(0)
	v9 = m.G0
	v11 = v9 - int32(16)
	m.G0 = v11
	v13 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	if v13 == v4 {
		goto L2
	} else {
		goto L3
	}
L1:
	;
	m.G0 = v11 + int32(16)
	return v64
L2:
	;
	if l2 != 0 {
		v64 = int32(0)
		goto L1
	} else {
		goto L10
	}
L3:
	;
	v21 = v4
	goto L4
L4:
	;
	v28 = l0 + int32(24) + v21<<(uint(int32(4))%32)
	v29 = *(*int64)(unsafe.Add(mBase, uint32(v28)))
	if l1 != v29 {
		goto L6
	} else {
		goto L7
	}
L5:
	;
	v34 = *(*int32)(unsafe.Add(mBase, uint32(v28)+8))
	v64 = l0 + v34
	goto L1
L6:
	;
	v32 = v21 + int32(1)
	if v13 != v32 {
		v21 = v32
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
	v50 = m.ExcPending
	if v50 != 0 {
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
	F_errmsg_internal(m, int32(238811), v11)
	mBase = m.M
	v55 = m.ExcPending
	if v55 != 0 {
		goto L11
	} else {
		goto L13
	}
L13:
	;
	F_errfinish(m, int32(499198), int32(254), int32(232793))
	mBase = m.M
	v60 = m.ExcPending
	if v60 != 0 {
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
		*(*int32)(unsafe.Add(mBase, _consts[166])) = int32(28)
		v53 = int32(0)
	} else {
		v21 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v19))))
		if v21 != 0 {
			*(*int32)(unsafe.Add(mBase, _consts[166])) = int32(28)
			v53 = int32(0)
		} else {
			v22 = v19 - v9
			if int32(2) < v22 {
				if base.Ui32(v22) < base.Ui32(int32(256)) {
					v45 = int32(9)
					v46 = F___memcpy(m, v5, int32(554944), v45)
					mBase = m.M
					v51 = F___memcpy(m, v5+v45, v9, v22+int32(1))
					mBase = m.M
					v53 = v5
				} else {
					*(*int32)(unsafe.Add(mBase, _consts[166])) = int32(37)
					v53 = int32(0)
				}
			} else {
				v25 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v9))))
				if v25 != int32(46) {
					v45 = int32(9)
					v46 = F___memcpy(m, v5, int32(554944), v45)
					mBase = m.M
					v51 = F___memcpy(m, v5+v45, v9, v22+int32(1))
					mBase = m.M
					v53 = v5
				} else {
					v30 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v19-int32(1)))))
					if v30 != int32(46) {
						v45 = int32(9)
						v46 = F___memcpy(m, v5, int32(554944), v45)
						mBase = m.M
						v51 = F___memcpy(m, v5+v45, v9, v22+int32(1))
						mBase = m.M
						v53 = v5
					} else {
						*(*int32)(unsafe.Add(mBase, _consts[166])) = int32(28)
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
