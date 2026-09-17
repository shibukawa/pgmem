package p0

import (
	base "github.com/shibukawa/pgmem/internal/aot/pgaot/base"
	"unsafe"
)

func F_SearchSysCache1(m *base.Module, l0 int32, l1 int32) int32 {
	mBase := m.M
	_ = mBase
	var v7 int32
	_ = v7
	var v8 int32
	_ = v8
	var v11 int32
	_ = v11
	v7 = *(*int32)(unsafe.Add(mBase, uint32(l0<<(uint(int32(2))%32))+uint32(_c_F_SearchSysCache1[0])))
	v8 = F_SearchCatCache1(m, v7, l1)
	mBase = m.M
	v11 = m.ExcPending
	if v11 != 0 {
		return int32(0)
	} else {
		return v8
	}
}
func F_SearchSysCacheCopy(m *base.Module, l0 int32, l1 int32, l2 int32) int32 {
	mBase := m.M
	_ = mBase
	var v8 int32
	_ = v8
	var v9 int32
	_ = v9
	var v11 int32
	_ = v11
	var v14 int32
	_ = v14
	var v19 int32
	_ = v19
	var v20 int32
	_ = v20
	var v22 int32
	_ = v22
	v8 = *(*int32)(unsafe.Add(mBase, uint32(l0<<(uint(int32(2))%32))+uint32(_c_F_SearchSysCacheCopy[0])))
	v9 = int32(0)
	v11 = F_SearchCatCache(m, v8, l1, l2, v9, v9)
	mBase = m.M
	v14 = m.ExcPending
	if v14 != 0 {
		return int32(0)
	} else {
		if v11 == int32(0) {
			return int32(0)
		} else {
			v19 = F_heap_copytuple(m, v11)
			mBase = m.M
			v20 = m.ExcPending
			if v20 != 0 {
				return int32(0)
			} else {
				F_ReleaseCatCache(m, v11)
				mBase = m.M
				v22 = m.ExcPending
				if v22 != 0 {
					return int32(0)
				} else {
					return v19
				}
			}
		}
	}
}
func F_SearchSysCacheLocked1(m *base.Module, l0 int32, l1 int32) int32 {
	mBase := m.M
	_ = mBase
	var v3 int32
	_ = v3
	var v6 int32
	_ = v6
	var v8 int32
	_ = v8
	var v11 int32
	_ = v11
	var v14 int32
	_ = v14
	var v19 int32
	_ = v19
	var v24 int32
	_ = v24
	var v27 int32
	_ = v27
	var v28 int32
	_ = v28
	var v31 int32
	_ = v31
	var v36 int32
	_ = v36
	var v37 int32
	_ = v37
	var v39 int32
	_ = v39
	var v41 int32
	_ = v41
	var v42 int32
	_ = v42
	var v43 int32
	_ = v43
	var v44 int32
	_ = v44
	var v47 int32
	_ = v47
	var v48 int32
	_ = v48
	var v54 int32
	_ = v54
	var v55 int32
	_ = v55
	var v58 int32
	_ = v58
	var v63 int32
	_ = v63
	var v64 int32
	_ = v64
	var v66 int32
	_ = v66
	var v71 int32
	_ = v71
	var v73 int32
	_ = v73
	var v76 int32
	_ = v76
	var v79 int32
	_ = v79
	var v80 int32
	_ = v80
	var v81 int32
	_ = v81
	var v83 int32
	_ = v83
	var v85 int32
	_ = v85
	var v87 int32
	_ = v87
	var v89 int32
	_ = v89
	var v90 int32
	_ = v90
	var v98 int32
	_ = v98
	var v100 int32
	_ = v100
	var v101 int32
	_ = v101
	var v103 int32
	_ = v103
	var v104 int32
	_ = v104
	v3 = int32(0)
	v6 = m.G0
	v8 = v6 - int32(32)
	m.G0 = v8
	v11 = l0 << (uint(int32(2)) % 32)
	v14 = *(*int32)(unsafe.Add(mBase, uint32(v11)+uint32(_c_F_SearchSysCacheLocked1[0])))
	*(*uint16)(unsafe.Add(mBase, uint32(v8)+28)) = uint16(v3)
	*(*int32)(unsafe.Add(mBase, uint32(v8)+24)) = int32(-1)
	v19 = v14
	goto L1
L1:
	;
	v24 = F_SearchCatCache1(m, v19, l1)
	mBase = m.M
	v27 = m.ExcPending
	if v27 != 0 {
		goto L3
	} else {
		goto L4
	}
L3:
	;
	return int32(0)
L4:
	;
	v28 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v8)+28)))
	if v28 != 0 {
		goto L7
	} else {
		goto L8
	}
L5:
	;
	v71 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v24)+8)))
	*(*uint16)(unsafe.Add(mBase, uint32(v8)+28)) = uint16(v71)
	v73 = *(*int32)(unsafe.Add(mBase, uint32(v24)+4))
	*(*int32)(unsafe.Add(mBase, uint32(v8)+24)) = v73
	F_ReleaseCatCache(m, v24)
	mBase = m.M
	v76 = m.ExcPending
	if v76 != 0 {
		goto L3
	} else {
		goto L23
	}
L6:
	;
	m.G0 = v8 + int32(32)
	return v66
L7:
	;
	if v24 == int32(0) {
		goto L10
	} else {
		goto L11
	}
L8:
	;
	goto L9
L9:
	;
	if v24 != 0 {
		goto L5
	} else {
		goto L22
	}
L10:
	;
	v31 = int32(0)
	v36 = F_LockRelease(m, v8+int32(8), int32(7), v31)
	mBase = m.M
	v37 = m.ExcPending
	if v37 != 0 {
		goto L3
	} else {
		goto L13
	}
L11:
	;
	goto L12
L12:
	;
	v39 = v8 + int32(24)
	v41 = v24 + int32(4)
	v42 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v39)+2)))
	v43 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v39))))
	v44 = int32(16)
	v47 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v41)+2)))
	v48 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v41))))
	if v42|v43<<(uint(v44)%32) == v47|v48<<(uint(v44)%32) {
		goto L16
	} else {
		goto L17
	}
L13:
	;
	v66 = v31
	goto L6
L14:
	;
	if v58 != 0 {
		v66 = v24
		goto L6
	} else {
		goto L20
	}
L15:
	;
	goto L14
L16:
	;
	v54 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v39)+4)))
	v55 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v41)+4)))
	if v54 == v55 {
		v58 = int32(1)
		goto L15
	} else {
		goto L19
	}
L17:
	;
	goto L18
L18:
	;
	v58 = int32(0)
	goto L15
L19:
	;
	goto L18
L20:
	;
	v63 = F_LockRelease(m, v8+int32(8), int32(7), int32(0))
	mBase = m.M
	v64 = m.ExcPending
	if v64 != 0 {
		goto L3
	} else {
		goto L21
	}
L21:
	;
	goto L5
L22:
	;
	v66 = int32(0)
	goto L6
L23:
	;
	v79 = *(*int32)(unsafe.Add(mBase, _c_F_SearchSysCacheLocked1[1]))
	v80 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v14)+96)))
	if v80 != 0 {
		goto L24
	} else {
		goto L25
	}
L24:
	;
	v81 = int32(0)
	goto L26
L25:
	;
	v81 = v79
	goto L26
L26:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v8)+8)) = v81
	v83 = *(*int32)(unsafe.Add(mBase, uint32(v14)+88))
	*(*int32)(unsafe.Add(mBase, uint32(v8)+12)) = v83
	v85 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v8)+28)))
	*(*uint16)(unsafe.Add(mBase, uint32(v8)+20)) = uint16(v85)
	v87 = int32(260)
	*(*uint16)(unsafe.Add(mBase, uint32(v8)+22)) = uint16(v87)
	v89 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v8)+26)))
	v90 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v8)+24)))
	*(*int32)(unsafe.Add(mBase, uint32(v8)+16)) = v89 | v90<<(uint(int32(16))%32)
	v98 = int32(0)
	v100 = F_LockAcquire(m, v8+int32(8), int32(7), v98, v98)
	mBase = m.M
	v101 = m.ExcPending
	if v101 != 0 {
		goto L3
	} else {
		goto L27
	}
L27:
	;
	F_ReceiveSharedInvalidMessages(m)
	mBase = m.M
	v103 = m.ExcPending
	if v103 != 0 {
		goto L3
	} else {
		goto L28
	}
L28:
	;
	v104 = *(*int32)(unsafe.Add(mBase, uint32(v11)+uint32(_c_F_SearchSysCacheLocked1[0])))
	v19 = v104
	goto L1
}
func F_SysCacheGetAttr(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) int32 {
	mBase := m.M
	_ = mBase
	var v8 int32
	_ = v8
	var v10 int32
	_ = v10
	var v15 int32
	_ = v15
	var v16 int32
	_ = v16
	var v19 int32
	_ = v19
	var v24 int32
	_ = v24
	var v27 int32
	_ = v27
	var v28 int32
	_ = v28
	var v29 int32
	_ = v29
	var v30 int32
	_ = v30
	var v31 int32
	_ = v31
	var v41 int32
	_ = v41
	var v45 int32
	_ = v45
	var v50 int32
	_ = v50
	v8 = m.G0
	v10 = v8 - int32(16)
	m.G0 = v10
	if base.Ui32(int32(84)) < base.Ui32(l0) {
		F_errstart_cold(m, int32(21), int32(0))
		mBase = m.M
		v41 = m.ExcPending
		if v41 != 0 {
			return int32(0)
		} else {
			*(*int32)(unsafe.Add(mBase, uint32(v10))) = l0
			F_errmsg_internal(m, int32(_a_F_SysCacheGetAttr_0), v10)
			mBase = m.M
			v45 = m.ExcPending
			if v45 != 0 {
				return int32(0)
			} else {
				F_errfinish(m, int32(_a_F_SysCacheGetAttr_1), int32(612), int32(_a_F_SysCacheGetAttr_2))
				mBase = m.M
				v50 = m.ExcPending
				if v50 != 0 {
					return int32(0)
				} else {
					base.Wasm_trap_unreachable()
					for {
					}
				}
			}
		}
	} else {
		v15 = l0 << (uint(int32(2)) % 32)
		v16 = *(*int32)(unsafe.Add(mBase, uint32(v15)+uint32(_c_F_SysCacheGetAttr[0])))
		if v16 == int32(0) {
			F_errstart_cold(m, int32(21), int32(0))
			mBase = m.M
			v41 = m.ExcPending
			if v41 != 0 {
				return int32(0)
			} else {
				*(*int32)(unsafe.Add(mBase, uint32(v10))) = l0
				F_errmsg_internal(m, int32(_a_F_SysCacheGetAttr_0), v10)
				mBase = m.M
				v45 = m.ExcPending
				if v45 != 0 {
					return int32(0)
				} else {
					F_errfinish(m, int32(_a_F_SysCacheGetAttr_1), int32(612), int32(_a_F_SysCacheGetAttr_2))
					mBase = m.M
					v50 = m.ExcPending
					if v50 != 0 {
						return int32(0)
					} else {
						base.Wasm_trap_unreachable()
						for {
						}
					}
				}
			}
		} else {
			v19 = *(*int32)(unsafe.Add(mBase, uint32(v16)+8))
			if v19 != 0 {
				v29 = v19
				v30 = F_heap_getattr_1(m, l1, l2, v29, l3)
				mBase = m.M
				v31 = m.ExcPending
				if v31 != 0 {
					return int32(0)
				} else {
					m.G0 = v10 + int32(16)
					return v30
				}
			} else {
				F_InitCatCachePhase2(m, v16, int32(0))
				mBase = m.M
				v24 = m.ExcPending
				if v24 != 0 {
					return int32(0)
				} else {
					v27 = *(*int32)(unsafe.Add(mBase, uint32(v15)+uint32(_c_F_SysCacheGetAttr[0])))
					v28 = *(*int32)(unsafe.Add(mBase, uint32(v27)+8))
					v29 = v28
					v30 = F_heap_getattr_1(m, l1, l2, v29, l3)
					mBase = m.M
					v31 = m.ExcPending
					if v31 != 0 {
						return int32(0)
					} else {
						m.G0 = v10 + int32(16)
						return v30
					}
				}
			}
		}
	}
}
