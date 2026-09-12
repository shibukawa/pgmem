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
	var v11 int32
	_ = v11
	var v12 int32
	_ = v12
	var v17 int32
	_ = v17
	var v19 int64
	_ = v19
	v4 = *(*int32)(unsafe.Add(mBase, _consts[8]))
	v11 = v4 + (l0^int32(-1))<<(uint(int32(6))%32) + int32(24)
	v12 = *(*int32)(unsafe.Add(mBase, uint32(v11)))
	if v12&int32(8388608) == int32(0) {
		v17 = int32(4387752)
		v19 = *(*int64)(unsafe.Add(mBase, _consts[110]))
		*(*int64)(unsafe.Add(mBase, _consts[110])) = v19 + int64(1)
	} else {
	}
	*(*int32)(unsafe.Add(mBase, uint32(v11))) = v12 | int32(8388608)
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
	v8 = *(*int32)(unsafe.Add(mBase, _consts[753]))
	v10 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v14 = v8 + (int32(-2)-v10)<<(uint(int32(2))%32)
	v15 = *(*int32)(unsafe.Add(mBase, uint32(v14)))
	if v15 == int32(0) {
		v18 = int32(4405384)
		v20 = *(*int32)(unsafe.Add(mBase, _consts[754]))
		v21 = int32(1)
		*(*int32)(unsafe.Add(mBase, _consts[754])) = v20 + v21
		v27 = v6 + v21
		if base.Ui32(v27&int32(3932160)) < base.Ui32(int32(1310720)) {
			v32 = v6 + int32(262145)
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
	v42 = *(*int32)(unsafe.Add(mBase, _consts[262]))
	v43 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	F_ResourceOwnerRemember(m, v42, v43+v38, int32(1612256))
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
	var v48 int32
	_ = v48
	var v53 int32
	_ = v53
	var v55 int32
	_ = v55
	var v60 int32
	_ = v60
	var v61 int32
	_ = v61
	var v65 int32
	_ = v65
	var v73 int32
	_ = v73
	var v76 int32
	_ = v76
	var v77 int32
	_ = v77
	var v80 int32
	_ = v80
	var v81 int32
	_ = v81
	var v87 int32
	_ = v87
	var v92 int32
	_ = v92
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
	v40 = *(*int32)(unsafe.Add(mBase, _consts[803]))
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
	v73 = *(*int32)(unsafe.Add(mBase, _consts[102]))
	v76 = F_hash_search(m, v73, l0, int32(2), int32(0))
	mBase = m.M
	v77 = m.ExcPending
	if v77 != 0 {
		goto L9
	} else {
		goto L24
	}
L19:
	;
	v48 = *(*int32)(unsafe.Add(mBase, _consts[803]))
	F_s_lock(m, v48, int32(497060), int32(1495), int32(317142))
	mBase = m.M
	v53 = m.ExcPending
	if v53 != 0 {
		goto L9
	} else {
		goto L22
	}
L20:
	;
	goto L21
L21:
	;
	v55 = *(*int32)(unsafe.Add(mBase, _consts[803]))
	v60 = v55 + v42&int32(1023)<<(uint(int32(2))%32) + int32(4)
	v61 = *(*int32)(unsafe.Add(mBase, uint32(v60)))
	*(*int32)(unsafe.Add(mBase, uint32(v60))) = v61 - int32(1)
	v65 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(l0)+52)) = uint8(v65)
	*(*int32)(unsafe.Add(mBase, uint32(v55))) = v65
	goto L18
L22:
	;
	goto L21
L23:
	;
	return
L24:
	;
	if v76 != 0 {
		goto L23
	} else {
		goto L25
	}
L25:
	;
	v80 = F_errstart(m, int32(19), int32(0))
	mBase = m.M
	v81 = m.ExcPending
	if v81 != 0 {
		goto L9
	} else {
		goto L26
	}
L26:
	;
	if v80 == int32(0) {
		goto L23
	} else {
		goto L27
	}
L27:
	;
	F_errmsg_internal(m, int32(444334), int32(0))
	mBase = m.M
	v87 = m.ExcPending
	if v87 != 0 {
		goto L9
	} else {
		goto L28
	}
L28:
	;
	F_errfinish(m, int32(497060), int32(1505), int32(317142))
	mBase = m.M
	v92 = m.ExcPending
	if v92 != 0 {
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
