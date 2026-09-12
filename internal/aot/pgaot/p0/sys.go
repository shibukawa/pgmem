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
	v7 = *(*int32)(unsafe.Add(mBase, uint32(l0<<(uint(int32(2))%32))+uint32(_consts[885])))
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
	v8 = *(*int32)(unsafe.Add(mBase, uint32(l0<<(uint(int32(2))%32))+uint32(_consts[885])))
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
	var v7 int32
	_ = v7
	var v9 int32
	_ = v9
	var v12 int32
	_ = v12
	var v15 int32
	_ = v15
	var v21 int32
	_ = v21
	var v22 int32
	_ = v22
	var v28 int32
	_ = v28
	var v31 int32
	_ = v31
	var v32 int32
	_ = v32
	var v35 int32
	_ = v35
	var v40 int32
	_ = v40
	var v41 int32
	_ = v41
	var v43 int32
	_ = v43
	var v45 int32
	_ = v45
	var v46 int32
	_ = v46
	var v47 int32
	_ = v47
	var v48 int32
	_ = v48
	var v51 int32
	_ = v51
	var v52 int32
	_ = v52
	var v58 int32
	_ = v58
	var v59 int32
	_ = v59
	var v62 int32
	_ = v62
	var v67 int32
	_ = v67
	var v68 int32
	_ = v68
	var v70 int32
	_ = v70
	var v75 int32
	_ = v75
	var v77 int32
	_ = v77
	var v80 int32
	_ = v80
	var v83 int32
	_ = v83
	var v84 int32
	_ = v84
	var v85 int32
	_ = v85
	var v87 int32
	_ = v87
	var v89 int32
	_ = v89
	var v91 int32
	_ = v91
	var v93 int32
	_ = v93
	var v94 int32
	_ = v94
	var v102 int32
	_ = v102
	var v104 int32
	_ = v104
	var v105 int32
	_ = v105
	var v107 int32
	_ = v107
	var v108 int32
	_ = v108
	v3 = int32(0)
	v7 = m.G0
	v9 = v7 - int32(32)
	m.G0 = v9
	v12 = l0 << (uint(int32(2)) % 32)
	v15 = *(*int32)(unsafe.Add(mBase, uint32(v12)+uint32(_consts[885])))
	*(*uint16)(unsafe.Add(mBase, uint32(v9)+28)) = uint16(v3)
	*(*int32)(unsafe.Add(mBase, uint32(v9)+24)) = int32(-1)
	v21 = v9 + int32(28)
	v22 = v15
	goto L1
L1:
	;
	v28 = F_SearchCatCache1(m, v22, l1)
	mBase = m.M
	v31 = m.ExcPending
	if v31 != 0 {
		goto L3
	} else {
		goto L4
	}
L3:
	;
	return int32(0)
L4:
	;
	v32 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v9)+28)))
	if v32 != 0 {
		goto L7
	} else {
		goto L8
	}
L5:
	;
	v75 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v28)+8)))
	*(*uint16)(unsafe.Add(mBase, uint32(v21))) = uint16(v75)
	v77 = *(*int32)(unsafe.Add(mBase, uint32(v28)+4))
	*(*int32)(unsafe.Add(mBase, uint32(v9)+24)) = v77
	F_ReleaseCatCache(m, v28)
	mBase = m.M
	v80 = m.ExcPending
	if v80 != 0 {
		goto L3
	} else {
		goto L23
	}
L6:
	;
	m.G0 = v9 + int32(32)
	return v70
L7:
	;
	if v28 == int32(0) {
		goto L10
	} else {
		goto L11
	}
L8:
	;
	goto L9
L9:
	;
	if v28 != 0 {
		goto L5
	} else {
		goto L22
	}
L10:
	;
	v35 = int32(0)
	v40 = F_LockRelease(m, v9+int32(8), int32(7), v35)
	mBase = m.M
	v41 = m.ExcPending
	if v41 != 0 {
		goto L3
	} else {
		goto L13
	}
L11:
	;
	goto L12
L12:
	;
	v43 = v9 + int32(24)
	v45 = v28 + int32(4)
	v46 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v43)+2)))
	v47 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v43))))
	v48 = int32(16)
	v51 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v45)+2)))
	v52 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v45))))
	if v46|v47<<(uint(v48)%32) == v51|v52<<(uint(v48)%32) {
		goto L16
	} else {
		goto L17
	}
L13:
	;
	v70 = v35
	goto L6
L14:
	;
	if v62 != 0 {
		v70 = v28
		goto L6
	} else {
		goto L20
	}
L15:
	;
	goto L14
L16:
	;
	v58 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v43)+4)))
	v59 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v45)+4)))
	if v58 == v59 {
		v62 = int32(1)
		goto L15
	} else {
		goto L19
	}
L17:
	;
	goto L18
L18:
	;
	v62 = int32(0)
	goto L15
L19:
	;
	goto L18
L20:
	;
	v67 = F_LockRelease(m, v9+int32(8), int32(7), int32(0))
	mBase = m.M
	v68 = m.ExcPending
	if v68 != 0 {
		goto L3
	} else {
		goto L21
	}
L21:
	;
	goto L5
L22:
	;
	v70 = int32(0)
	goto L6
L23:
	;
	v83 = *(*int32)(unsafe.Add(mBase, _consts[226]))
	v84 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v15)+96)))
	if v84 != 0 {
		goto L24
	} else {
		goto L25
	}
L24:
	;
	v85 = int32(0)
	goto L26
L25:
	;
	v85 = v83
	goto L26
L26:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v9)+8)) = v85
	v87 = *(*int32)(unsafe.Add(mBase, uint32(v15)+88))
	*(*int32)(unsafe.Add(mBase, uint32(v9)+12)) = v87
	v89 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v21))))
	*(*uint16)(unsafe.Add(mBase, uint32(v9)+20)) = uint16(v89)
	v91 = int32(260)
	*(*uint16)(unsafe.Add(mBase, uint32(v9)+22)) = uint16(v91)
	v93 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v9)+26)))
	v94 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v9)+24)))
	*(*int32)(unsafe.Add(mBase, uint32(v9)+16)) = v93 | v94<<(uint(int32(16))%32)
	v102 = int32(0)
	v104 = F_LockAcquire(m, v9+int32(8), int32(7), v102, v102)
	mBase = m.M
	v105 = m.ExcPending
	if v105 != 0 {
		goto L3
	} else {
		goto L27
	}
L27:
	;
	F_ReceiveSharedInvalidMessages(m)
	mBase = m.M
	v107 = m.ExcPending
	if v107 != 0 {
		goto L3
	} else {
		goto L28
	}
L28:
	;
	v108 = *(*int32)(unsafe.Add(mBase, uint32(v12)+uint32(_consts[885])))
	v22 = v108
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
	var v18 int32
	_ = v18
	var v21 int32
	_ = v21
	var v26 int32
	_ = v26
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
			F_errmsg_internal(m, int32(465570), v10)
			mBase = m.M
			v45 = m.ExcPending
			if v45 != 0 {
				return int32(0)
			} else {
				F_errfinish(m, int32(476073), int32(612), int32(196045))
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
		v18 = *(*int32)(unsafe.Add(mBase, uint32(v15)+uint32(_consts[885])))
		if v18 == int32(0) {
			F_errstart_cold(m, int32(21), int32(0))
			mBase = m.M
			v41 = m.ExcPending
			if v41 != 0 {
				return int32(0)
			} else {
				*(*int32)(unsafe.Add(mBase, uint32(v10))) = l0
				F_errmsg_internal(m, int32(465570), v10)
				mBase = m.M
				v45 = m.ExcPending
				if v45 != 0 {
					return int32(0)
				} else {
					F_errfinish(m, int32(476073), int32(612), int32(196045))
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
			v21 = *(*int32)(unsafe.Add(mBase, uint32(v18)+8))
			if v21 != 0 {
				v29 = v21
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
				F_InitCatCachePhase2(m, v18, int32(0))
				mBase = m.M
				v26 = m.ExcPending
				if v26 != 0 {
					return int32(0)
				} else {
					v27 = *(*int32)(unsafe.Add(mBase, uint32(v15)+uint32(_consts[885])))
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
