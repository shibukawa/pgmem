package p5

import (
	base "github.com/shibukawa/pgmem/internal/aot/pgaot/base"
	"unsafe"
)

func F_ShmemCallRequestCallbacks(m *base.Module) {
	mBase := m.M
	_ = mBase
	var v1 int32
	_ = v1
	var v9 int32
	_ = v9
	var v12 int32
	_ = v12
	var v16 int32
	_ = v16
	var v19 int32
	_ = v19
	var v23 int32
	_ = v23
	var v24 int32
	_ = v24
	var v25 int32
	_ = v25
	var v27 int32
	_ = v27
	var v29 int32
	_ = v29
	var v30 int32
	_ = v30
	v1 = int32(0)
	*(*int32)(unsafe.Add(mBase, _c_F_ShmemCallRequestCallbacks[0])) = int32(1)
	v9 = *(*int32)(unsafe.Add(mBase, _c_F_ShmemCallRequestCallbacks[1]))
	if v9 == v1 {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	return
L2:
	;
	v12 = *(*int32)(unsafe.Add(mBase, uint32(v9)+4))
	if v12 <= int32(0) {
		goto L1
	} else {
		goto L3
	}
L3:
	;
	v16 = v1
	goto L4
L4:
	;
	v19 = *(*int32)(unsafe.Add(mBase, uint32(v9)+12))
	v23 = *(*int32)(unsafe.Add(mBase, uint32(v19+v16<<(uint(int32(2))%32))))
	v24 = *(*int32)(unsafe.Add(mBase, uint32(v23)+4))
	if v24 != 0 {
		goto L6
	} else {
		goto L7
	}
L5:
	;
	goto L1
L6:
	;
	v25 = *(*int32)(unsafe.Add(mBase, uint32(v23)+16))
	m.T0[v24].(func(*base.Module, int32))(m, v25)
	mBase = m.M
	v27 = m.ExcPending
	if v27 != 0 {
		goto L9
	} else {
		goto L10
	}
L7:
	;
	goto L8
L8:
	;
	v29 = v16 + int32(1)
	v30 = *(*int32)(unsafe.Add(mBase, uint32(v9)+4))
	if v29 < v30 {
		v16 = v29
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
}
func F_ShmemInitStruct(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v2 int32
	_ = v2
	var v4 int32
	_ = v4
	var v6 int32
	_ = v6
	var v22 int32
	_ = v22
	var v29 int32
	_ = v29
	var v32 int32
	_ = v32
	var v33 int32
	_ = v33
	var v36 int32
	_ = v36
	var v40 int32
	_ = v40
	var v41 int32
	_ = v41
	var v45 int32
	_ = v45
	var v46 int32
	_ = v46
	var v47 int32
	_ = v47
	var v51 int32
	_ = v51
	var v55 int32
	_ = v55
	var v56 int32
	_ = v56
	v2 = int32(0)
	v4 = m.G0
	v6 = v4 - int32(32)
	m.G0 = v6
	*(*int32)(unsafe.Add(mBase, uint32(v6)+28)) = v2
	*(*int32)(unsafe.Add(mBase, uint32(v6)+20)) = v2
	*(*int32)(unsafe.Add(mBase, uint32(v6)+16)) = int32(4)
	*(*int32)(unsafe.Add(mBase, uint32(v6)+12)) = int32(_a_F_ShmemInitStruct_0)
	*(*int32)(unsafe.Add(mBase, uint32(v6)+24)) = v6 + int32(28)
	*(*int64)(unsafe.Add(mBase, uint32(v6)+4)) = int64(0)
	v22 = *(*int32)(unsafe.Add(mBase, _c_F_ShmemInitStruct[0]))
	*(*int32)(unsafe.Add(mBase, uint32(v6))) = v6 + int32(12)
	v29 = F_LWLockAcquire(m, v22+int32(16), v2)
	mBase = m.M
	v32 = m.ExcPending
	if v32 != 0 {
		return int32(0)
	} else {
		v33 = int32(0)
		*(*uint8)(unsafe.Add(mBase, uint32(l0))) = uint8(v33)
		v36 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_ShmemInitStruct[1])))
		if v36 == int32(1) {
			v40 = F_AttachShmemIndexEntry(m, v6, int32(1))
			mBase = m.M
			v41 = m.ExcPending
			if v41 != 0 {
				return int32(0)
			} else {
				*(*uint8)(unsafe.Add(mBase, uint32(l0))) = uint8(v40)
				if v40 != 0 {
					v51 = *(*int32)(unsafe.Add(mBase, _c_F_ShmemInitStruct[0]))
					F_LWLockRelease(m, v51+int32(16))
					mBase = m.M
					v55 = m.ExcPending
					if v55 != 0 {
						return int32(0)
					} else {
						v56 = *(*int32)(unsafe.Add(mBase, uint32(v6)+28))
						m.G0 = v6 + int32(32)
						return v56
					}
				} else {
					F_InitShmemIndexEntry(m, v6)
					mBase = m.M
					v45 = m.ExcPending
					if v45 != 0 {
						return int32(0)
					} else {
						v46 = *(*int32)(unsafe.Add(mBase, uint32(v6)+8))
						v47 = int32(1)
						*(*uint8)(unsafe.Add(mBase, uint32(v46)+60)) = uint8(v47)
						v51 = *(*int32)(unsafe.Add(mBase, _c_F_ShmemInitStruct[0]))
						F_LWLockRelease(m, v51+int32(16))
						mBase = m.M
						v55 = m.ExcPending
						if v55 != 0 {
							return int32(0)
						} else {
							v56 = *(*int32)(unsafe.Add(mBase, uint32(v6)+28))
							m.G0 = v6 + int32(32)
							return v56
						}
					}
				}
			}
		} else {
			F_InitShmemIndexEntry(m, v6)
			mBase = m.M
			v45 = m.ExcPending
			if v45 != 0 {
				return int32(0)
			} else {
				v46 = *(*int32)(unsafe.Add(mBase, uint32(v6)+8))
				v47 = int32(1)
				*(*uint8)(unsafe.Add(mBase, uint32(v46)+60)) = uint8(v47)
				v51 = *(*int32)(unsafe.Add(mBase, _c_F_ShmemInitStruct[0]))
				F_LWLockRelease(m, v51+int32(16))
				mBase = m.M
				v55 = m.ExcPending
				if v55 != 0 {
					return int32(0)
				} else {
					v56 = *(*int32)(unsafe.Add(mBase, uint32(v6)+28))
					m.G0 = v6 + int32(32)
					return v56
				}
			}
		}
	}
}
func F_ShmemRequestHashWithOpts(m *base.Module, l0 int32) {
	mBase := m.M
	_ = mBase
	var v4 int32
	_ = v4
	var v6 int32
	_ = v6
	var v7 int32
	_ = v7
	var v10 int32
	_ = v10
	var v12 int64
	_ = v12
	var v13 int32
	_ = v13
	var v14 int32
	_ = v14
	var v15 int32
	_ = v15
	var v19 int32
	_ = v19
	v4 = *(*int32)(unsafe.Add(mBase, _c_F_ShmemRequestHashWithOpts[0]))
	v6 = F_MemoryContextAlloc(m, v4, int32(88))
	mBase = m.M
	v7 = m.ExcPending
	if v7 != 0 {
		return
	} else {
		base.MemoryCopy(m, v6, l0, int32(88))
		v10 = *(*int32)(unsafe.Add(mBase, uint32(l0)+16))
		*(*int32)(unsafe.Add(mBase, uint32(v6))) = v10
		v12 = *(*int64)(unsafe.Add(mBase, uint32(v6)+24))
		v13 = *(*int32)(unsafe.Add(mBase, uint32(v6)+44))
		v14 = F_hash_estimate_size(m, v12, v13)
		mBase = m.M
		v15 = m.ExcPending
		if v15 != 0 {
			return
		} else {
			*(*int32)(unsafe.Add(mBase, uint32(v6)+4)) = v14
			F_ShmemRequestInternal(m, v6, int32(1))
			mBase = m.M
			v19 = m.ExcPending
			if v19 != 0 {
				return
			} else {
				return
			}
		}
	}
}
func F_process_shmem_requests(m *base.Module) {
	mBase := m.M
	_ = mBase
	var v3 int32
	_ = v3
	var v6 int32
	_ = v6
	var v8 int32
	_ = v8
	var v10 int32
	_ = v10
	v3 = int32(1)
	*(*uint8)(unsafe.Add(mBase, _c_F_process_shmem_requests[0])) = uint8(v3)
	v6 = *(*int32)(unsafe.Add(mBase, _c_F_process_shmem_requests[1]))
	if v6 != 0 {
		m.T0[v6].(func(*base.Module))(m)
		mBase = m.M
		v8 = m.ExcPending
		if v8 != 0 {
			return
		} else {
			v10 = int32(0)
			*(*uint8)(unsafe.Add(mBase, _c_F_process_shmem_requests[0])) = uint8(v10)
			return
		}
	} else {
		v10 = int32(0)
		*(*uint8)(unsafe.Add(mBase, _c_F_process_shmem_requests[0])) = uint8(v10)
		return
	}
}
