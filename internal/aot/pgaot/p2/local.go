package p2

import (
	base "github.com/shibukawa/pgmem/internal/aot/pgaot/base"
	"sync/atomic"
	"unsafe"
)

func F_MarkLocalBufferDirty(m *base.Module, l0 int32) {
	mBase := m.M
	_ = mBase
	var v4 int32
	_ = v4
	var v9 int32
	_ = v9
	var v10 int32
	_ = v10
	var v15 int32
	_ = v15
	var v17 int64
	_ = v17
	v4 = *(*int32)(unsafe.Add(mBase, _c_F_MarkLocalBufferDirty[0]))
	v9 = v4 + (l0^int32(-1))<<(uint(int32(6))%32)
	v10 = *(*int32)(unsafe.Add(mBase, uint32(v9)+24))
	if v10&int32(_a_F_MarkLocalBufferDirty_0) == int32(0) {
		v15 = int32(_a_F_MarkLocalBufferDirty_1)
		v17 = *(*int64)(unsafe.Add(mBase, _c_F_MarkLocalBufferDirty[1]))
		*(*int64)(unsafe.Add(mBase, _c_F_MarkLocalBufferDirty[1])) = v17 + int64(1)
	} else {
	}
	*(*int32)(unsafe.Add(mBase, uint32(v9)+24)) = v10 | int32(_a_F_MarkLocalBufferDirty_0)
	return
}
func F_PinLocalBuffer(m *base.Module, l0 int32, l1 int32) {
	mBase := m.M
	_ = mBase
	var v6 int32
	_ = v6
	var v8 int32
	_ = v8
	var v10 int32
	_ = v10
	var v14 int32
	_ = v14
	var v15 int32
	_ = v15
	var v18 int32
	_ = v18
	var v20 int32
	_ = v20
	var v21 int32
	_ = v21
	var v27 int32
	_ = v27
	var v32 int32
	_ = v32
	var v33 int32
	_ = v33
	var v35 int32
	_ = v35
	var v37 int32
	_ = v37
	var v38 int32
	_ = v38
	var v42 int32
	_ = v42
	var v43 int32
	_ = v43
	var v48 int32
	_ = v48
	v6 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
	v8 = *(*int32)(unsafe.Add(mBase, _c_F_PinLocalBuffer[0]))
	v10 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v14 = v8 + (int32(-2)-v10)<<(uint(int32(2))%32)
	v15 = *(*int32)(unsafe.Add(mBase, uint32(v14)))
	if v15 == int32(0) {
		v18 = int32(_a_F_PinLocalBuffer_0)
		v20 = *(*int32)(unsafe.Add(mBase, _c_F_PinLocalBuffer[1]))
		v21 = int32(1)
		*(*int32)(unsafe.Add(mBase, _c_F_PinLocalBuffer[1])) = v20 + v21
		v27 = v6 + v21
		if base.Ui32(v27&int32(_a_F_PinLocalBuffer_1)) < base.Ui32(int32(_a_F_PinLocalBuffer_2)) {
			v32 = v6 + int32(_a_F_PinLocalBuffer_3)
		} else {
			v32 = v27
		}
		if l1 != 0 {
			v33 = v32
		} else {
			v33 = v27
		}
		*(*int32)(unsafe.Add(mBase, uint32(l0)+24)) = v33
		v35 = *(*int32)(unsafe.Add(mBase, uint32(v14)))
		v37 = v35
	} else {
		v37 = v15
	}
	v38 = int32(1)
	*(*int32)(unsafe.Add(mBase, uint32(v14))) = v37 + v38
	v42 = *(*int32)(unsafe.Add(mBase, _c_F_PinLocalBuffer[2]))
	v43 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	F_ResourceOwnerRemember(m, v42, v43+v38, int32(_a_F_PinLocalBuffer_4))
	mBase = m.M
	v48 = m.ExcPending
	if v48 != 0 {
		return
	} else {
		return
	}
}
func F_RemoveLocalLock(m *base.Module, l0 int32) {
	mBase := m.M
	_ = mBase
	var v4 int32
	_ = v4
	var v6 int32
	_ = v6
	var v10 int32
	_ = v10
	var v12 int32
	_ = v12
	var v16 int32
	_ = v16
	var v18 int32
	_ = v18
	var v28 int32
	_ = v28
	var v30 int32
	_ = v30
	var v33 int32
	_ = v33
	var v36 int32
	_ = v36
	var v40 int32
	_ = v40
	var v43 int32
	_ = v43
	var v45 int32
	_ = v45
	var v50 int32
	_ = v50
	var v52 int32
	_ = v52
	var v55 int32
	_ = v55
	var v56 int32
	_ = v56
	var v60 int32
	_ = v60
	var v68 int32
	_ = v68
	var v71 int32
	_ = v71
	var v72 int32
	_ = v72
	var v75 int32
	_ = v75
	var v76 int32
	_ = v76
	var v82 int32
	_ = v82
	var v87 int32
	_ = v87
	v4 = *(*int32)(unsafe.Add(mBase, uint32(l0)+40))
	v6 = v4 - int32(1)
	if int32(0) <= v6 {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	v10 = v6
	goto L4
L2:
	;
	goto L3
L3:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+40)) = int32(0)
	v28 = *(*int32)(unsafe.Add(mBase, uint32(l0)+48))
	if v28 != 0 {
		goto L12
	} else {
		goto L13
	}
L4:
	;
	v12 = *(*int32)(unsafe.Add(mBase, uint32(l0)+48))
	v16 = *(*int32)(unsafe.Add(mBase, uint32(v12+v10<<(uint(int32(4))%32))))
	if v16 != 0 {
		goto L6
	} else {
		goto L7
	}
L5:
	;
	goto L3
L6:
	;
	F_ResourceOwnerForgetLock(m, v16, l0)
	mBase = m.M
	v18 = m.ExcPending
	if v18 != 0 {
		goto L9
	} else {
		goto L10
	}
L7:
	;
	goto L8
L8:
	;
	if int32(0) < v10 {
		v10 = v10 - int32(1)
		goto L4
	} else {
		goto L11
	}
L9:
	;
	return
L10:
	;
	goto L8
L11:
	;
	goto L5
L12:
	;
	F_pfree(m, v28)
	mBase = m.M
	v30 = m.ExcPending
	if v30 != 0 {
		goto L9
	} else {
		goto L15
	}
L13:
	;
	goto L14
L14:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+48)) = int32(0)
	v33 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+52)))
	if v33 == int32(1) {
		goto L16
	} else {
		goto L17
	}
L15:
	;
	goto L14
L16:
	;
	v36 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v40 = *(*int32)(unsafe.Add(mBase, _c_F_RemoveLocalLock[0]))
	v43 = base.AtomicRmwXchg32(m, v40, int32(0), int32(1))
	if v43 != 0 {
		goto L19
	} else {
		goto L20
	}
L17:
	;
	goto L18
L18:
	;
	v68 = *(*int32)(unsafe.Add(mBase, _c_F_RemoveLocalLock[1]))
	v71 = F_hash_search(m, v68, l0, int32(2), int32(0))
	mBase = m.M
	v72 = m.ExcPending
	if v72 != 0 {
		goto L9
	} else {
		goto L24
	}
L19:
	;
	v45 = *(*int32)(unsafe.Add(mBase, _c_F_RemoveLocalLock[0]))
	F_s_lock(m, v45, int32(_a_F_RemoveLocalLock_0), int32(1495), int32(_a_F_RemoveLocalLock_1))
	mBase = m.M
	v50 = m.ExcPending
	if v50 != 0 {
		goto L9
	} else {
		goto L22
	}
L20:
	;
	goto L21
L21:
	;
	v52 = *(*int32)(unsafe.Add(mBase, _c_F_RemoveLocalLock[0]))
	v55 = v52 + v36&int32(1023)<<(uint(int32(2))%32)
	v56 = *(*int32)(unsafe.Add(mBase, uint32(v55)+4))
	*(*int32)(unsafe.Add(mBase, uint32(v55)+4)) = v56 - int32(1)
	v60 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(l0)+52)) = uint8(v60)
	atomic.StoreUint32((*uint32)(unsafe.Add(mBase, uint32(v52))), uint32(v60))
	goto L18
L22:
	;
	goto L21
L23:
	;
	return
L24:
	;
	if v71 != 0 {
		goto L23
	} else {
		goto L25
	}
L25:
	;
	v75 = F_errstart(m, int32(19), int32(0))
	mBase = m.M
	v76 = m.ExcPending
	if v76 != 0 {
		goto L9
	} else {
		goto L26
	}
L26:
	;
	if v75 == int32(0) {
		goto L23
	} else {
		goto L27
	}
L27:
	;
	F_errmsg_internal(m, int32(_a_F_RemoveLocalLock_2), int32(0))
	mBase = m.M
	v82 = m.ExcPending
	if v82 != 0 {
		goto L9
	} else {
		goto L28
	}
L28:
	;
	F_errfinish(m, int32(_a_F_RemoveLocalLock_0), int32(1505), int32(_a_F_RemoveLocalLock_1))
	mBase = m.M
	v87 = m.ExcPending
	if v87 != 0 {
		goto L9
	} else {
		goto L29
	}
L29:
	;
	goto L23
}
func F_add_local_int_reloption(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32, l5 int32) {
	mBase := m.M
	_ = mBase
	var v10 int32
	_ = v10
	var v11 int32
	_ = v11
	var v12 int32
	_ = v12
	var v13 int32
	_ = v13
	var v15 int32
	_ = v15
	var v16 int32
	_ = v16
	var v17 int32
	_ = v17
	var v18 int32
	_ = v18
	var v21 int32
	_ = v21
	var v31 int32
	_ = v31
	var v32 int32
	_ = v32
	var v36 int32
	_ = v36
	var v37 int32
	_ = v37
	var v38 int32
	_ = v38
	v10 = F_palloc(m, int32(36))
	mBase = m.M
	v11 = m.ExcPending
	if v11 != 0 {
		return
	} else {
		v12 = F_pstrdup(m, l1)
		mBase = m.M
		v13 = m.ExcPending
		if v13 != 0 {
			return
		} else {
			*(*int32)(unsafe.Add(mBase, uint32(v10))) = v12
			if l2 != 0 {
				v15 = F_pstrdup(m, l2)
				mBase = m.M
				v16 = m.ExcPending
				if v16 != 0 {
					return
				} else {
					v17 = v15
					v18 = int32(0)
					*(*int32)(unsafe.Add(mBase, uint32(v10)+8)) = v18
					*(*int32)(unsafe.Add(mBase, uint32(v10)+4)) = v17
					v21 = F_strlen(m, l1)
					mBase = m.M
					*(*int32)(unsafe.Add(mBase, uint32(v10)+20)) = int32(1)
					*(*int32)(unsafe.Add(mBase, uint32(v10)+16)) = v21
					*(*int32)(unsafe.Add(mBase, uint32(v10)+32)) = l5
					*(*int32)(unsafe.Add(mBase, uint32(v10)+28)) = l4
					*(*int32)(unsafe.Add(mBase, uint32(v10)+24)) = l3
					*(*int32)(unsafe.Add(mBase, uint32(v10)+12)) = v18
					v31 = F_palloc(m, int32(8))
					mBase = m.M
					v32 = m.ExcPending
					if v32 != 0 {
						return
					} else {
						*(*int32)(unsafe.Add(mBase, uint32(v31)+4)) = int32(4)
						*(*int32)(unsafe.Add(mBase, uint32(v31))) = v10
						v36 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
						v37 = F_lappend(m, v36, v31)
						mBase = m.M
						v38 = m.ExcPending
						if v38 != 0 {
							return
						} else {
							*(*int32)(unsafe.Add(mBase, uint32(l0))) = v37
							return
						}
					}
				}
			} else {
				v17 = int32(0)
				v18 = int32(0)
				*(*int32)(unsafe.Add(mBase, uint32(v10)+8)) = v18
				*(*int32)(unsafe.Add(mBase, uint32(v10)+4)) = v17
				v21 = F_strlen(m, l1)
				mBase = m.M
				*(*int32)(unsafe.Add(mBase, uint32(v10)+20)) = int32(1)
				*(*int32)(unsafe.Add(mBase, uint32(v10)+16)) = v21
				*(*int32)(unsafe.Add(mBase, uint32(v10)+32)) = l5
				*(*int32)(unsafe.Add(mBase, uint32(v10)+28)) = l4
				*(*int32)(unsafe.Add(mBase, uint32(v10)+24)) = l3
				*(*int32)(unsafe.Add(mBase, uint32(v10)+12)) = v18
				v31 = F_palloc(m, int32(8))
				mBase = m.M
				v32 = m.ExcPending
				if v32 != 0 {
					return
				} else {
					*(*int32)(unsafe.Add(mBase, uint32(v31)+4)) = int32(4)
					*(*int32)(unsafe.Add(mBase, uint32(v31))) = v10
					v36 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
					v37 = F_lappend(m, v36, v31)
					mBase = m.M
					v38 = m.ExcPending
					if v38 != 0 {
						return
					} else {
						*(*int32)(unsafe.Add(mBase, uint32(l0))) = v37
						return
					}
				}
			}
		}
	}
}
