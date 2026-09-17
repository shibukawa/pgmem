package p2

import (
	base "github.com/shibukawa/pgmem/internal/aot/pgaot/base"
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
	var v5 int32
	_ = v5
	var v7 int32
	_ = v7
	var v11 int32
	_ = v11
	var v14 int32
	_ = v14
	var v18 int32
	_ = v18
	var v20 int32
	_ = v20
	var v31 int32
	_ = v31
	var v33 int32
	_ = v33
	var v36 int32
	_ = v36
	var v40 int32
	_ = v40
	var v41 int32
	_ = v41
	var v42 int32
	_ = v42
	var v46 int32
	_ = v46
	var v51 int32
	_ = v51
	var v53 int32
	_ = v53
	var v58 int32
	_ = v58
	var v59 int32
	_ = v59
	var v63 int32
	_ = v63
	var v71 int32
	_ = v71
	var v74 int32
	_ = v74
	var v75 int32
	_ = v75
	var v78 int32
	_ = v78
	var v79 int32
	_ = v79
	var v85 int32
	_ = v85
	var v90 int32
	_ = v90
	v5 = *(*int32)(unsafe.Add(mBase, uint32(l0)+40))
	v7 = v5 - int32(1)
	if int32(0) <= v7 {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	v11 = v7
	goto L4
L2:
	;
	goto L3
L3:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+40)) = int32(0)
	v31 = *(*int32)(unsafe.Add(mBase, uint32(l0)+48))
	if v31 != 0 {
		goto L12
	} else {
		goto L13
	}
L4:
	;
	v14 = *(*int32)(unsafe.Add(mBase, uint32(l0)+48))
	v18 = *(*int32)(unsafe.Add(mBase, uint32(v14+v11<<(uint(int32(4))%32))))
	if v18 != 0 {
		goto L6
	} else {
		goto L7
	}
L5:
	;
	goto L3
L6:
	;
	F_ResourceOwnerForgetLock(m, v18, l0)
	mBase = m.M
	v20 = m.ExcPending
	if v20 != 0 {
		goto L9
	} else {
		goto L10
	}
L7:
	;
	goto L8
L8:
	;
	if int32(0) < v11 {
		v11 = v11 - int32(1)
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
	F_pfree(m, v31)
	mBase = m.M
	v33 = m.ExcPending
	if v33 != 0 {
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
	v36 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+52)))
	if v36 == int32(1) {
		goto L16
	} else {
		goto L17
	}
L15:
	;
	goto L14
L16:
	;
	v40 = *(*int32)(unsafe.Add(mBase, _c_F_RemoveLocalLock[0]))
	v41 = *(*int32)(unsafe.Add(mBase, uint32(v40)))
	v42 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	*(*int32)(unsafe.Add(mBase, uint32(v40))) = int32(1)
	if v41 != 0 {
		goto L19
	} else {
		goto L20
	}
L17:
	;
	goto L18
L18:
	;
	v71 = *(*int32)(unsafe.Add(mBase, _c_F_RemoveLocalLock[1]))
	v74 = F_hash_search(m, v71, l0, int32(2), int32(0))
	mBase = m.M
	v75 = m.ExcPending
	if v75 != 0 {
		goto L9
	} else {
		goto L24
	}
L19:
	;
	v46 = *(*int32)(unsafe.Add(mBase, _c_F_RemoveLocalLock[0]))
	F_s_lock(m, v46, int32(_a_F_RemoveLocalLock_0), int32(1495), int32(_a_F_RemoveLocalLock_1))
	mBase = m.M
	v51 = m.ExcPending
	if v51 != 0 {
		goto L9
	} else {
		goto L22
	}
L20:
	;
	goto L21
L21:
	;
	v53 = *(*int32)(unsafe.Add(mBase, _c_F_RemoveLocalLock[0]))
	v58 = v53 + v42&int32(1023)<<(uint(int32(2))%32)
	v59 = *(*int32)(unsafe.Add(mBase, uint32(v58)+4))
	*(*int32)(unsafe.Add(mBase, uint32(v58)+4)) = v59 - int32(1)
	v63 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(l0)+52)) = uint8(v63)
	*(*int32)(unsafe.Add(mBase, uint32(v53))) = v63
	goto L18
L22:
	;
	goto L21
L23:
	;
	return
L24:
	;
	if v74 != 0 {
		goto L23
	} else {
		goto L25
	}
L25:
	;
	v78 = F_errstart(m, int32(19), int32(0))
	mBase = m.M
	v79 = m.ExcPending
	if v79 != 0 {
		goto L9
	} else {
		goto L26
	}
L26:
	;
	if v78 == int32(0) {
		goto L23
	} else {
		goto L27
	}
L27:
	;
	F_errmsg_internal(m, int32(_a_F_RemoveLocalLock_2), int32(0))
	mBase = m.M
	v85 = m.ExcPending
	if v85 != 0 {
		goto L9
	} else {
		goto L28
	}
L28:
	;
	F_errfinish(m, int32(_a_F_RemoveLocalLock_0), int32(1505), int32(_a_F_RemoveLocalLock_1))
	mBase = m.M
	v90 = m.ExcPending
	if v90 != 0 {
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
