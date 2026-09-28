package p3

import (
	base "github.com/shibukawa/pgmem/internal/aot/pgaot/base"
	"unsafe"
)

func F_AsyncShmemRequest(m *base.Module, l0 int32) {
	mBase := m.M
	_ = mBase
	var v3 int32
	_ = v3
	var v5 int32
	_ = v5
	var v8 int32
	_ = v8
	var v10 int32
	_ = v10
	var v11 int32
	_ = v11
	var v13 int32
	_ = v13
	var v14 int32
	_ = v14
	var v25 int32
	_ = v25
	var v26 int64
	_ = v26
	var v34 int32
	_ = v34
	var v36 int32
	_ = v36
	var v51 int32
	_ = v51
	var v54 int32
	_ = v54
	v3 = m.G0
	v5 = v3 - int32(80)
	m.G0 = v5
	v8 = *(*int32)(unsafe.Add(mBase, _c_F_AsyncShmemRequest[0]))
	v10 = F_mul_size(m, v8, int32(40))
	mBase = m.M
	v11 = m.ExcPending
	if v11 != 0 {
		return
	} else {
		v13 = F_add_size(m, v10, int32(64))
		mBase = m.M
		v14 = m.ExcPending
		if v14 != 0 {
			return
		} else {
			*(*int32)(unsafe.Add(mBase, uint32(v5)+76)) = int32(_a_F_AsyncShmemRequest_0)
			*(*int32)(unsafe.Add(mBase, uint32(v5)+72)) = int32(0)
			*(*int32)(unsafe.Add(mBase, uint32(v5)+68)) = v13
			*(*int32)(unsafe.Add(mBase, uint32(v5)+64)) = int32(_a_F_AsyncShmemRequest_1)
			F_ShmemRequestStructWithOpts(m, v5-int32(-64))
			mBase = m.M
			v25 = m.ExcPending
			if v25 != 0 {
				return
			} else {
				v26 = int64(0)
				*(*int64)(unsafe.Add(mBase, uint32(v5)+8)) = v26
				*(*int64)(unsafe.Add(mBase, uint32(v5))) = v26
				*(*int32)(unsafe.Add(mBase, uint32(v5)+20)) = int32(_a_F_AsyncShmemRequest_2)
				*(*int32)(unsafe.Add(mBase, uint32(v5)+16)) = int32(_a_F_AsyncShmemRequest_3)
				v34 = int32(0)
				*(*uint16)(unsafe.Add(mBase, uint32(v5)+41)) = uint16(v34)
				v36 = int32(1)
				*(*uint8)(unsafe.Add(mBase, uint32(v5)+40)) = uint8(v36)
				*(*int32)(unsafe.Add(mBase, uint32(v5)+36)) = int32(_a_F_AsyncShmemRequest_4)
				*(*int64)(unsafe.Add(mBase, uint32(v5)+28)) = int64(21474836480)
				*(*uint8)(unsafe.Add(mBase, uint32(v5)+43)) = uint8(v34)
				*(*int64)(unsafe.Add(mBase, uint32(v5)+52)) = int64(399431958591)
				*(*int32)(unsafe.Add(mBase, uint32(v5)+48)) = int32(545)
				*(*int32)(unsafe.Add(mBase, uint32(v5)+44)) = int32(546)
				v51 = *(*int32)(unsafe.Add(mBase, _c_F_AsyncShmemRequest[1]))
				*(*int32)(unsafe.Add(mBase, uint32(v5)+24)) = v51
				F_SimpleLruRequestWithOpts(m, v5)
				mBase = m.M
				v54 = m.ExcPending
				if v54 != 0 {
					return
				} else {
					m.G0 = v5 + int32(80)
					return
				}
			}
		}
	}
}
func F_asyncQueueUnregister(m *base.Module) {
	mBase := m.M
	_ = mBase
	var v7 int32
	_ = v7
	var v9 int32
	_ = v9
	var v13 int32
	_ = v13
	var v14 int32
	_ = v14
	var v15 int32
	_ = v15
	var v16 int32
	_ = v16
	var v17 int32
	_ = v17
	var v18 int32
	_ = v18
	var v19 int32
	_ = v19
	var v27 int32
	_ = v27
	var v29 int32
	_ = v29
	var v33 int32
	_ = v33
	var v36 int32
	_ = v36
	var v38 int32
	_ = v38
	var v45 int32
	_ = v45
	var v47 int32
	_ = v47
	var v54 int32
	_ = v54
	var v55 int32
	_ = v55
	var v57 int32
	_ = v57
	var v60 int32
	_ = v60
	var v61 int32
	_ = v61
	var v70 int32
	_ = v70
	var v71 int32
	_ = v71
	var v76 int32
	_ = v76
	var v79 int32
	_ = v79
	var v81 int32
	_ = v81
	var v85 int32
	_ = v85
	var v87 int32
	_ = v87
	var v88 int32
	_ = v88
	var v98 int32
	_ = v98
	var v102 int32
	_ = v102
	var v104 int32
	_ = v104
	v7 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_asyncQueueUnregister[0])))
	if v7 != 0 {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	v9 = *(*int32)(unsafe.Add(mBase, _c_F_asyncQueueUnregister[1]))
	v13 = F_LWLockAcquire(m, v9+int32(3456), int32(0))
	mBase = m.M
	v14 = m.ExcPending
	if v14 != 0 {
		goto L4
	} else {
		goto L5
	}
L2:
	;
	goto L3
L3:
	;
	return
L4:
	;
	return
L5:
	;
	v15 = int32(_a_F_asyncQueueUnregister_0)
	v16 = *(*int32)(unsafe.Add(mBase, _c_F_asyncQueueUnregister[2]))
	v17 = int32(_a_F_asyncQueueUnregister_1)
	v18 = *(*int32)(unsafe.Add(mBase, _c_F_asyncQueueUnregister[3]))
	v19 = int32(40)
	*(*int32)(unsafe.Add(mBase, uint32(v16+v18*v19-int32(-64)))) = int32(-1)
	v27 = *(*int32)(unsafe.Add(mBase, _c_F_asyncQueueUnregister[2]))
	v29 = *(*int32)(unsafe.Add(mBase, _c_F_asyncQueueUnregister[3]))
	v33 = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v27+v29*v19)+68)) = v33
	v36 = *(*int32)(unsafe.Add(mBase, _c_F_asyncQueueUnregister[2]))
	v38 = *(*int32)(unsafe.Add(mBase, _c_F_asyncQueueUnregister[3]))
	*(*uint8)(unsafe.Add(mBase, uint32(v36+v38*v19)+96)) = uint8(v33)
	v45 = *(*int32)(unsafe.Add(mBase, _c_F_asyncQueueUnregister[2]))
	v47 = *(*int32)(unsafe.Add(mBase, _c_F_asyncQueueUnregister[3]))
	*(*uint8)(unsafe.Add(mBase, uint32(v45+v47*v19)+97)) = uint8(v33)
	v54 = *(*int32)(unsafe.Add(mBase, _c_F_asyncQueueUnregister[2]))
	v55 = *(*int32)(unsafe.Add(mBase, uint32(v54)+40))
	v57 = *(*int32)(unsafe.Add(mBase, _c_F_asyncQueueUnregister[3]))
	if v55 != v57 {
		goto L7
	} else {
		goto L8
	}
L6:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v88+v87*int32(40))+72)) = int32(-1)
	v98 = *(*int32)(unsafe.Add(mBase, _c_F_asyncQueueUnregister[1]))
	F_LWLockRelease(m, v98+int32(3456))
	mBase = m.M
	v102 = m.ExcPending
	if v102 != 0 {
		goto L4
	} else {
		goto L16
	}
L7:
	;
	v60 = v54 - int32(-64)
	v61 = v55
	goto L10
L8:
	;
	goto L9
L9:
	;
	v85 = *(*int32)(unsafe.Add(mBase, uint32(v54+v55*int32(40))+72))
	*(*int32)(unsafe.Add(mBase, uint32(v54)+40)) = v85
	v87 = v55
	v88 = v54
	goto L6
L10:
	;
	if v61 == int32(-1) {
		goto L12
	} else {
		goto L13
	}
L11:
	;
	v76 = *(*int32)(unsafe.Add(mBase, uint32(v60+v57*int32(40))+8))
	*(*int32)(unsafe.Add(mBase, uint32(v70)+8)) = v76
	v79 = *(*int32)(unsafe.Add(mBase, _c_F_asyncQueueUnregister[3]))
	v81 = *(*int32)(unsafe.Add(mBase, _c_F_asyncQueueUnregister[2]))
	v87 = v79
	v88 = v81
	goto L6
L12:
	;
	v87 = v57
	v88 = v54
	goto L6
L13:
	;
	goto L14
L14:
	;
	v70 = v60 + v61*int32(40)
	v71 = *(*int32)(unsafe.Add(mBase, uint32(v70)+8))
	if v71 != v57 {
		v61 = v71
		goto L10
	} else {
		goto L15
	}
L15:
	;
	goto L11
L16:
	;
	v104 = int32(0)
	*(*uint8)(unsafe.Add(mBase, _c_F_asyncQueueUnregister[0])) = uint8(v104)
	goto L3
}
