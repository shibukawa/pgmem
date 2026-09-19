package p0

import (
	base "github.com/shibukawa/pgmem/internal/aot/pgaot/base"
	"sync/atomic"
	"unsafe"
)

func F_AbortStrongLockAcquire(m *base.Module) {
	mBase := m.M
	_ = mBase
	var v5 int32
	_ = v5
	var v6 int32
	_ = v6
	var v10 int32
	_ = v10
	var v13 int32
	_ = v13
	var v15 int32
	_ = v15
	var v20 int32
	_ = v20
	var v22 int32
	_ = v22
	var v25 int32
	_ = v25
	var v26 int32
	_ = v26
	var v30 int32
	_ = v30
	v5 = *(*int32)(unsafe.Add(mBase, _c_F_AbortStrongLockAcquire[0]))
	if v5 != 0 {
		v6 = *(*int32)(unsafe.Add(mBase, uint32(v5)+20))
		v10 = *(*int32)(unsafe.Add(mBase, _c_F_AbortStrongLockAcquire[1]))
		v13 = base.AtomicRmwXchg32(m, v10, int32(0), int32(1))
		if v13 != 0 {
			v15 = *(*int32)(unsafe.Add(mBase, _c_F_AbortStrongLockAcquire[1]))
			F_s_lock(m, v15, int32(_a_F_AbortStrongLockAcquire_0), int32(1869), int32(_a_F_AbortStrongLockAcquire_1))
			mBase = m.M
			v20 = m.ExcPending
			if v20 != 0 {
				return
			} else {
				v22 = *(*int32)(unsafe.Add(mBase, _c_F_AbortStrongLockAcquire[1]))
				v25 = v22 + v6&int32(1023)<<(uint(int32(2))%32)
				v26 = *(*int32)(unsafe.Add(mBase, uint32(v25)+4))
				*(*int32)(unsafe.Add(mBase, uint32(v25)+4)) = v26 - int32(1)
				v30 = int32(0)
				*(*uint8)(unsafe.Add(mBase, uint32(v5)+52)) = uint8(v30)
				*(*int32)(unsafe.Add(mBase, _c_F_AbortStrongLockAcquire[0])) = v30
				atomic.StoreUint32((*uint32)(unsafe.Add(mBase, uint32(v22))), uint32(v30))
				return
			}
		} else {
			v22 = *(*int32)(unsafe.Add(mBase, _c_F_AbortStrongLockAcquire[1]))
			v25 = v22 + v6&int32(1023)<<(uint(int32(2))%32)
			v26 = *(*int32)(unsafe.Add(mBase, uint32(v25)+4))
			*(*int32)(unsafe.Add(mBase, uint32(v25)+4)) = v26 - int32(1)
			v30 = int32(0)
			*(*uint8)(unsafe.Add(mBase, uint32(v5)+52)) = uint8(v30)
			*(*int32)(unsafe.Add(mBase, _c_F_AbortStrongLockAcquire[0])) = v30
			atomic.StoreUint32((*uint32)(unsafe.Add(mBase, uint32(v22))), uint32(v30))
			return
		}
	} else {
		return
	}
}
func F_AbsorbSyncRequests(m *base.Module) {
	mBase := m.M
	_ = mBase
	var v7 int32
	_ = v7
	var v11 int32
	_ = v11
	var v15 int32
	_ = v15
	var v16 int32
	_ = v16
	var v18 int32
	_ = v18
	var v19 int32
	_ = v19
	var v21 int32
	_ = v21
	var v23 int32
	_ = v23
	var v24 int32
	_ = v24
	var v25 int32
	_ = v25
	var v27 int32
	_ = v27
	var v33 int32
	_ = v33
	var v34 int32
	_ = v34
	var v36 int32
	_ = v36
	var v38 int32
	_ = v38
	var v45 int32
	_ = v45
	var v49 int32
	_ = v49
	var v52 int32
	_ = v52
	var v53 int32
	_ = v53
	var v59 int32
	_ = v59
	var v61 int32
	_ = v61
	var v64 int32
	_ = v64
	var v73 int32
	_ = v73
	var v75 int32
	_ = v75
	var v82 int32
	_ = v82
	v7 = *(*int32)(unsafe.Add(mBase, _c_F_AbsorbSyncRequests[0]))
	if v7 != int32(11) {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	return
L2:
	;
	v11 = *(*int32)(unsafe.Add(mBase, _c_F_AbsorbSyncRequests[1]))
	v15 = F_LWLockAcquire(m, v11+int32(2176), int32(0))
	mBase = m.M
	v16 = m.ExcPending
	if v16 != 0 {
		goto L3
	} else {
		goto L4
	}
L3:
	;
	return
L4:
	;
	v18 = *(*int32)(unsafe.Add(mBase, _c_F_AbsorbSyncRequests[2]))
	v19 = *(*int32)(unsafe.Add(mBase, uint32(v18)+48))
	v21 = base.B2i32(v19 <= int32(0))
	if v19 <= int32(0) {
		v33 = v18
		v34 = int32(0)
		goto L5
	} else {
		goto L6
	}
L5:
	;
	v36 = int32(_a_F_AbsorbSyncRequests_0)
	v38 = *(*int32)(unsafe.Add(mBase, _c_F_AbsorbSyncRequests[3]))
	*(*int32)(unsafe.Add(mBase, _c_F_AbsorbSyncRequests[3])) = v38 + int32(1)
	*(*int32)(unsafe.Add(mBase, uint32(v33)+48)) = int32(0)
	v45 = *(*int32)(unsafe.Add(mBase, _c_F_AbsorbSyncRequests[1]))
	F_LWLockRelease(m, v45+int32(2176))
	mBase = m.M
	v49 = m.ExcPending
	if v49 != 0 {
		goto L3
	} else {
		goto L9
	}
L6:
	;
	v23 = v19 << (uint(int32(5)) % 32)
	v24 = F_palloc(m, v23)
	mBase = m.M
	v25 = m.ExcPending
	if v25 != 0 {
		goto L3
	} else {
		goto L7
	}
L7:
	;
	v27 = *(*int32)(unsafe.Add(mBase, _c_F_AbsorbSyncRequests[2]))
	if v23 == int32(0) {
		v33 = v27
		v34 = v24
		goto L5
	} else {
		goto L8
	}
L8:
	;
	base.MemoryCopy(m, v24, v27+int32(56), v23)
	v33 = v27
	v34 = v24
	goto L5
L9:
	;
	if v21 == int32(0) {
		goto L10
	} else {
		goto L11
	}
L10:
	;
	v52 = v34
	v53 = v19
	goto L13
L11:
	;
	goto L12
L12:
	;
	v73 = int32(_a_F_AbsorbSyncRequests_0)
	v75 = *(*int32)(unsafe.Add(mBase, _c_F_AbsorbSyncRequests[3]))
	*(*int32)(unsafe.Add(mBase, _c_F_AbsorbSyncRequests[3])) = v75 - int32(1)
	if v34 == int32(0) {
		goto L1
	} else {
		goto L17
	}
L13:
	;
	v59 = *(*int32)(unsafe.Add(mBase, uint32(v52)))
	F_RememberSyncRequest(m, v52+int32(8), v59)
	mBase = m.M
	v61 = m.ExcPending
	if v61 != 0 {
		goto L3
	} else {
		goto L15
	}
L14:
	;
	goto L12
L15:
	;
	v64 = int32(1)
	if base.Ui32(v64) < base.Ui32(v53) {
		v52 = v52 + int32(32)
		v53 = v53 - v64
		goto L13
	} else {
		goto L16
	}
L16:
	;
	goto L14
L17:
	;
	F_pfree(m, v34)
	mBase = m.M
	v82 = m.ExcPending
	if v82 != 0 {
		goto L3
	} else {
		goto L18
	}
L18:
	;
	goto L1
}
func F_AlignedAllocGetChunkSpace(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v3 int32
	_ = v3
	var v4 int64
	_ = v4
	var v11 int32
	_ = v11
	var v14 int32
	_ = v14
	v3 = l0 - int32(8)
	v4 = *(*int64)(unsafe.Add(mBase, uint32(v3)))
	v11 = F_GetMemoryChunkSpace(m, v3-base.I32_wrap_i64(int64(base.Ui64(v4)>>(uint(int64(34))%64)))&int32(1073741822))
	mBase = m.M
	v14 = m.ExcPending
	if v14 != 0 {
		return int32(0)
	} else {
		return v11
	}
}
func F_AllocSetAlloc(m *base.Module, l0 int32, l1 int32, l2 int32) int32 {
	mBase := m.M
	_ = mBase
	var v8 int32
	_ = v8
	var v10 int32
	_ = v10
	var v13 int32
	_ = v13
	var v23 int32
	_ = v23
	var v26 int32
	_ = v26
	var v27 int32
	_ = v27
	var v28 int32
	_ = v28
	var v33 int32
	_ = v33
	var v34 int32
	_ = v34
	var v37 int32
	_ = v37
	var v38 int32
	_ = v38
	var v39 int32
	_ = v39
	var v42 int32
	_ = v42
	var v43 int32
	_ = v43
	var v46 int32
	_ = v46
	v8 = *(*int32)(unsafe.Add(mBase, uint32(l0)+104))
	if base.Ui32(v8) < base.Ui32(l1) {
		v10 = F_AllocSetAllocLarge(m, l0, l1, l2)
		mBase = m.M
		v13 = m.ExcPending
		if v13 != 0 {
			return int32(0)
		} else {
			return v10
		}
	} else {
		if base.Ui32(int32(8)) < base.Ui32(l1) {
			v23 = int32(29) - base.I32_clz(l1-int32(1))
		} else {
			v23 = int32(0)
		}
		v26 = l0 + v23<<(uint(int32(2))%32)
		v27 = *(*int32)(unsafe.Add(mBase, uint32(v26)+48))
		if v27 != 0 {
			v28 = *(*int32)(unsafe.Add(mBase, uint32(v27)+8))
			*(*int32)(unsafe.Add(mBase, uint32(v26)+48)) = v28
			return v27 + int32(8)
		} else {
			v33 = int32(8)
			v34 = v33 << (uint(v23) % 32)
			v37 = *(*int32)(unsafe.Add(mBase, uint32(l0)+44))
			v38 = *(*int32)(unsafe.Add(mBase, uint32(v37)+16))
			v39 = *(*int32)(unsafe.Add(mBase, uint32(v37)+12))
			if base.Ui32(v38-v39) < base.Ui32(v34+v33) {
				v42 = F_AllocSetAllocFromNewBlock(m, l0, l1, l2, v23)
				mBase = m.M
				v43 = m.ExcPending
				if v43 != 0 {
					return int32(0)
				} else {
					return v42
				}
			} else {
				v46 = int32(8)
				*(*int32)(unsafe.Add(mBase, uint32(v37)+12)) = v39 + v34 + v46
				*(*int64)(unsafe.Add(mBase, uint32(v39))) = base.I64_extend_i32_u(v23<<(uint(int32(5))%32)) | base.I64_extend_i32_u(v39-v37)<<(uint(int64(34))%64) | int64(3)
				return v39 + v46
			}
		}
	}
}
func F_AllocSetIsEmpty(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v2 int32
	_ = v2
	v2 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+4)))
	return v2
}
func F_AllocSetReset(m *base.Module, l0 int32) {
	mBase := m.M
	_ = mBase
	var v7 int64
	_ = v7
	var v17 int32
	_ = v17
	var v19 int32
	_ = v19
	var v22 int32
	_ = v22
	var v25 int32
	_ = v25
	var v32 int32
	_ = v32
	var v33 int32
	_ = v33
	var v42 int32
	_ = v42
	*(*int32)(unsafe.Add(mBase, uint32(l0)+88)) = int32(0)
	v7 = int64(0)
	*(*int64)(unsafe.Add(mBase, uint32(l0)+80)) = v7
	*(*int64)(unsafe.Add(mBase, uint32(l0)+72)) = v7
	*(*int64)(unsafe.Add(mBase, uint32(l0)+64)) = v7
	*(*int64)(unsafe.Add(mBase, uint32(l0)+56)) = v7
	*(*int64)(unsafe.Add(mBase, uint32(l0)+48)) = v7
	v17 = *(*int32)(unsafe.Add(mBase, uint32(l0)+44))
	v19 = l0 + int32(112)
	*(*int32)(unsafe.Add(mBase, uint32(l0)+44)) = v19
	if v17 != 0 {
		v22 = v17
		for {
			v25 = *(*int32)(unsafe.Add(mBase, uint32(v22)+8))
			if v22 == v19 {
				*(*int64)(unsafe.Add(mBase, uint32(v22)+4)) = int64(0)
				*(*int32)(unsafe.Add(mBase, uint32(v22)+12)) = v22 + int32(24)
			} else {
				v32 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
				v33 = *(*int32)(unsafe.Add(mBase, uint32(v22)+16))
				*(*int32)(unsafe.Add(mBase, uint32(l0)+8)) = v32 + (v22 - v33)
				F_emscripten_builtin_free(m, v22)
				mBase = m.M
			}
			if v25 != 0 {
				v22 = v25
				continue
			} else {
				break
			}
			break
		}
	} else {
	}
	v42 = *(*int32)(unsafe.Add(mBase, uint32(l0)+92))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+100)) = v42
	return
}
func F_AllocateDir(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v2 int32
	_ = v2
	var v5 int32
	_ = v5
	var v7 int32
	_ = v7
	var v9 int32
	_ = v9
	var v12 int32
	_ = v12
	var v14 int32
	_ = v14
	var v18 int32
	_ = v18
	var v20 int32
	_ = v20
	var v22 int32
	_ = v22
	var v31 int32
	_ = v31
	var v32 int32
	_ = v32
	var v34 int32
	_ = v34
	var v36 int32
	_ = v36
	var v40 int32
	_ = v40
	var v42 int32
	_ = v42
	var v44 int32
	_ = v44
	var v52 int32
	_ = v52
	var v60 int32
	_ = v60
	var v65 int32
	_ = v65
	var v66 int32
	_ = v66
	var v69 int32
	_ = v69
	var v73 int32
	_ = v73
	var v78 int32
	_ = v78
	var v80 int32
	_ = v80
	var v83 int32
	_ = v83
	var v89 int32
	_ = v89
	var v90 int32
	_ = v90
	var v92 int32
	_ = v92
	var v93 int32
	_ = v93
	var v97 int32
	_ = v97
	var v101 int32
	_ = v101
	var v103 int32
	_ = v103
	var v106 int32
	_ = v106
	var v111 int32
	_ = v111
	var v112 int32
	_ = v112
	var v114 int32
	_ = v114
	var v116 int32
	_ = v116
	var v120 int32
	_ = v120
	var v124 int32
	_ = v124
	var v132 int32
	_ = v132
	var v135 int32
	_ = v135
	var v138 int32
	_ = v138
	var v142 int32
	_ = v142
	var v147 int32
	_ = v147
	v2 = int32(0)
	v5 = m.G0
	v7 = v5 - int32(16)
	m.G0 = v7
	v9 = F_reserveAllocatedDesc(m)
	mBase = m.M
	v12 = m.ExcPending
	if v12 != 0 {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	return int32(0)
L2:
	;
	if v9 != 0 {
		goto L3
	} else {
		goto L4
	}
L3:
	;
	v14 = *(*int32)(unsafe.Add(mBase, _c_F_AllocateDir[0]))
	if v14 <= int32(0) {
		goto L6
	} else {
		goto L7
	}
L4:
	;
	goto L5
L5:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v132 = m.ExcPending
	if v132 != 0 {
		goto L1
	} else {
		goto L34
	}
L6:
	;
	v52 = F_opendir(m, l0)
	mBase = m.M
	if v52 == int32(0) {
		goto L15
	} else {
		goto L16
	}
L7:
	;
	v18 = *(*int32)(unsafe.Add(mBase, _c_F_AllocateDir[1]))
	v20 = *(*int32)(unsafe.Add(mBase, _c_F_AllocateDir[2]))
	v22 = *(*int32)(unsafe.Add(mBase, _c_F_AllocateDir[3]))
	if v20+(v22+v14) < v18 {
		goto L6
	} else {
		goto L8
	}
L8:
	;
	goto L9
L9:
	;
	v31 = *(*int32)(unsafe.Add(mBase, _c_F_AllocateDir[4]))
	v32 = *(*int32)(unsafe.Add(mBase, uint32(v31)+16))
	F_LruDelete(m, v32)
	mBase = m.M
	v34 = m.ExcPending
	if v34 != 0 {
		goto L1
	} else {
		goto L11
	}
L10:
	;
	goto L6
L11:
	;
	v36 = *(*int32)(unsafe.Add(mBase, _c_F_AllocateDir[0]))
	if v36 <= int32(0) {
		goto L6
	} else {
		goto L12
	}
L12:
	;
	v40 = *(*int32)(unsafe.Add(mBase, _c_F_AllocateDir[1]))
	v42 = *(*int32)(unsafe.Add(mBase, _c_F_AllocateDir[2]))
	v44 = *(*int32)(unsafe.Add(mBase, _c_F_AllocateDir[3]))
	if v40 <= v42+(v44+v36) {
		goto L9
	} else {
		goto L13
	}
L13:
	;
	goto L10
L14:
	;
	m.G0 = v7 + int32(16)
	return v124
L15:
	;
	goto L18
L16:
	;
	v97 = v52
	goto L17
L17:
	;
	v101 = *(*int32)(unsafe.Add(mBase, _c_F_AllocateDir[5]))
	v103 = *(*int32)(unsafe.Add(mBase, _c_F_AllocateDir[3]))
	v106 = v101 + v103*int32(12)
	*(*int32)(unsafe.Add(mBase, uint32(v106)+8)) = v97
	*(*int32)(unsafe.Add(mBase, uint32(v106))) = int32(2)
	v111 = *(*int32)(unsafe.Add(mBase, _c_F_AllocateDir[6]))
	v112 = *(*int32)(unsafe.Add(mBase, uint32(v111)+8))
	goto L33
L18:
	;
	v60 = *(*int32)(unsafe.Add(mBase, _c_F_AllocateDir[7]))
	switch v60 - int32(33) {
	case 0, 8:
		goto L20
	default:
		v124 = v2
		goto L14
	}
L19:
	;
	v97 = v93
	goto L17
L20:
	;
	v65 = F_errstart(m, int32(15), int32(0))
	mBase = m.M
	v66 = m.ExcPending
	if v66 != 0 {
		goto L1
	} else {
		goto L21
	}
L21:
	;
	if v65 != 0 {
		goto L22
	} else {
		goto L23
	}
L22:
	;
	F_errcode(m, int32(197))
	mBase = m.M
	v69 = m.ExcPending
	if v69 != 0 {
		goto L1
	} else {
		goto L25
	}
L23:
	;
	goto L24
L24:
	;
	v80 = int32(0)
	*(*int32)(unsafe.Add(mBase, _c_F_AllocateDir[7])) = v80
	v83 = *(*int32)(unsafe.Add(mBase, _c_F_AllocateDir[0]))
	if v83 <= v80 {
		goto L28
	} else {
		goto L29
	}
L25:
	;
	F_errmsg(m, int32(_a_F_AllocateDir_0), int32(0))
	mBase = m.M
	v73 = m.ExcPending
	if v73 != 0 {
		goto L1
	} else {
		goto L26
	}
L26:
	;
	F_errfinish(m, int32(_a_F_AllocateDir_1), int32(2947), int32(_a_F_AllocateDir_2))
	mBase = m.M
	v78 = m.ExcPending
	if v78 != 0 {
		goto L1
	} else {
		goto L27
	}
L27:
	;
	goto L24
L28:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_AllocateDir[7])) = v60
	v124 = v2
	goto L14
L29:
	;
	goto L30
L30:
	;
	v89 = *(*int32)(unsafe.Add(mBase, _c_F_AllocateDir[4]))
	v90 = *(*int32)(unsafe.Add(mBase, uint32(v89)+16))
	F_LruDelete(m, v90)
	mBase = m.M
	v92 = m.ExcPending
	if v92 != 0 {
		goto L1
	} else {
		goto L31
	}
L31:
	;
	v93 = F_opendir(m, l0)
	mBase = m.M
	if v93 == int32(0) {
		goto L18
	} else {
		goto L32
	}
L32:
	;
	goto L19
L33:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v106)+4)) = v112
	v114 = int32(_a_F_AllocateDir_3)
	v116 = *(*int32)(unsafe.Add(mBase, _c_F_AllocateDir[3]))
	*(*int32)(unsafe.Add(mBase, _c_F_AllocateDir[3])) = v116 + int32(1)
	v120 = *(*int32)(unsafe.Add(mBase, uint32(v106)+8))
	v124 = v120
	goto L14
L34:
	;
	F_errcode(m, int32(197))
	mBase = m.M
	v135 = m.ExcPending
	if v135 != 0 {
		goto L1
	} else {
		goto L35
	}
L35:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v7)+4)) = l0
	v138 = *(*int32)(unsafe.Add(mBase, _c_F_AllocateDir[8]))
	*(*int32)(unsafe.Add(mBase, uint32(v7))) = v138
	F_errmsg(m, int32(_a_F_AllocateDir_4), v7)
	mBase = m.M
	v142 = m.ExcPending
	if v142 != 0 {
		goto L1
	} else {
		goto L36
	}
L36:
	;
	F_errfinish(m, int32(_a_F_AllocateDir_1), int32(2924), int32(_a_F_AllocateDir_2))
	mBase = m.M
	v147 = m.ExcPending
	if v147 != 0 {
		goto L1
	} else {
		goto L37
	}
L37:
	;
	base.Wasm_trap_unreachable()
	for {
	}
}
func F_AlterForeignDataWrapperOwner_internal(m *base.Module, l0 int32, l1 int32, l2 int32) {
	mBase := m.M
	_ = mBase
	var v7 int32
	_ = v7
	var v9 int32
	_ = v9
	var v11 int32
	_ = v11
	var v12 int32
	_ = v12
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
	var v20 int32
	_ = v20
	var v22 int32
	_ = v22
	var v31 int32
	_ = v31
	var v34 int32
	_ = v34
	var v37 int32
	_ = v37
	var v38 int32
	_ = v38
	var v39 int32
	_ = v39
	var v42 int32
	_ = v42
	var v43 int32
	_ = v43
	var v44 int32
	_ = v44
	var v45 int32
	_ = v45
	var v46 int32
	_ = v46
	var v48 int32
	_ = v48
	var v50 int32
	_ = v50
	var v57 int32
	_ = v57
	var v58 int32
	_ = v58
	var v62 int32
	_ = v62
	var v64 int32
	_ = v64
	var v66 int32
	_ = v66
	var v70 int32
	_ = v70
	var v72 int32
	_ = v72
	var v73 int32
	_ = v73
	var v77 int32
	_ = v77
	var v84 int32
	_ = v84
	var v87 int32
	_ = v87
	var v95 int32
	_ = v95
	var v99 int32
	_ = v99
	var v104 int32
	_ = v104
	var v108 int32
	_ = v108
	var v111 int32
	_ = v111
	var v117 int32
	_ = v117
	var v121 int32
	_ = v121
	var v126 int32
	_ = v126
	v7 = m.G0
	v9 = v7 - int32(80)
	m.G0 = v9
	v11 = *(*int32)(unsafe.Add(mBase, uint32(l1)+16))
	v12 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v11)+22)))
	v13 = v11 + v12
	v14 = F_superuser(m)
	mBase = m.M
	v15 = m.ExcPending
	if v15 != 0 {
		return
	} else {
		if v14 != 0 {
			v16 = F_superuser_arg(m, l2)
			mBase = m.M
			v17 = m.ExcPending
			if v17 != 0 {
				return
			} else {
				if v16 == int32(0) {
					F_errstart_cold(m, int32(21), int32(0))
					mBase = m.M
					v108 = m.ExcPending
					if v108 != 0 {
						return
					} else {
						F_errcode(m, int32(16797828))
						mBase = m.M
						v111 = m.ExcPending
						if v111 != 0 {
							return
						} else {
							*(*int32)(unsafe.Add(mBase, uint32(v9))) = v13 + int32(4)
							F_errmsg(m, int32(_a_F_AlterForeignDataWrapperOwner_internal_0), v9)
							mBase = m.M
							v117 = m.ExcPending
							if v117 != 0 {
								return
							} else {
								F_errhint(m, int32(_a_F_AlterForeignDataWrapperOwner_internal_1), int32(0))
								mBase = m.M
								v121 = m.ExcPending
								if v121 != 0 {
									return
								} else {
									F_errfinish(m, int32(_a_F_AlterForeignDataWrapperOwner_internal_2), int32(242), int32(_a_F_AlterForeignDataWrapperOwner_internal_3))
									mBase = m.M
									v126 = m.ExcPending
									if v126 != 0 {
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
					v20 = *(*int32)(unsafe.Add(mBase, uint32(v13)+68))
					if l2 != v20 {
						v22 = int32(0)
						*(*int32)(unsafe.Add(mBase, uint32(v9)+32)) = v22
						*(*int32)(unsafe.Add(mBase, uint32(v9)+43)) = v22
						*(*int32)(unsafe.Add(mBase, uint32(v9)+40)) = v22
						*(*int32)(unsafe.Add(mBase, uint32(v9)+35)) = v22
						*(*int32)(unsafe.Add(mBase, uint32(v9)+56)) = l2
						v31 = int32(1)
						*(*uint8)(unsafe.Add(mBase, uint32(v9)+34)) = uint8(v31)
						v34 = *(*int32)(unsafe.Add(mBase, uint32(l0)+52))
						v37 = F_heap_getattr_7(m, l1, int32(6), v34, v9+int32(31))
						mBase = m.M
						v38 = m.ExcPending
						if v38 != 0 {
							return
						} else {
							v39 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v9)+31)))
							if v39 == int32(0) {
								v42 = F_pg_detoast_datum(m, v37)
								mBase = m.M
								v43 = m.ExcPending
								if v43 != 0 {
									return
								} else {
									v44 = *(*int32)(unsafe.Add(mBase, uint32(v13)+68))
									v45 = F_aclnewowner(m, v42, v44, l2)
									mBase = m.M
									v46 = m.ExcPending
									if v46 != 0 {
										return
									} else {
										*(*int32)(unsafe.Add(mBase, uint32(v9)+68)) = v45
										v48 = int32(1)
										*(*uint8)(unsafe.Add(mBase, uint32(v9)+37)) = uint8(v48)
										v50 = *(*int32)(unsafe.Add(mBase, uint32(l0)+52))
										v57 = F_heap_modify_tuple(m, l1, v50, v9+int32(48), v9+int32(40), v9+int32(32))
										mBase = m.M
										v58 = m.ExcPending
										if v58 != 0 {
											return
										} else {
											F_CatalogTupleUpdate(m, l0, v57+int32(4), v57)
											mBase = m.M
											v62 = m.ExcPending
											if v62 != 0 {
												return
											} else {
												v64 = *(*int32)(unsafe.Add(mBase, uint32(v13)))
												F_changeDependencyOnOwner(m, int32(2328), v64, l2)
												mBase = m.M
												v66 = m.ExcPending
												if v66 != 0 {
													return
												} else {
													v70 = *(*int32)(unsafe.Add(mBase, _c_F_AlterForeignDataWrapperOwner_internal[0]))
													if v70 != 0 {
														v72 = *(*int32)(unsafe.Add(mBase, uint32(v13)))
														v73 = int32(0)
														F_RunObjectPostAlterHook(m, int32(2328), v72, v73, v73, v73)
														mBase = m.M
														v77 = m.ExcPending
														if v77 != 0 {
															return
														} else {
															m.G0 = v9 + int32(80)
															return
														}
													} else {
														m.G0 = v9 + int32(80)
														return
													}
												}
											}
										}
									}
								}
							} else {
								v50 = *(*int32)(unsafe.Add(mBase, uint32(l0)+52))
								v57 = F_heap_modify_tuple(m, l1, v50, v9+int32(48), v9+int32(40), v9+int32(32))
								mBase = m.M
								v58 = m.ExcPending
								if v58 != 0 {
									return
								} else {
									F_CatalogTupleUpdate(m, l0, v57+int32(4), v57)
									mBase = m.M
									v62 = m.ExcPending
									if v62 != 0 {
										return
									} else {
										v64 = *(*int32)(unsafe.Add(mBase, uint32(v13)))
										F_changeDependencyOnOwner(m, int32(2328), v64, l2)
										mBase = m.M
										v66 = m.ExcPending
										if v66 != 0 {
											return
										} else {
											v70 = *(*int32)(unsafe.Add(mBase, _c_F_AlterForeignDataWrapperOwner_internal[0]))
											if v70 != 0 {
												v72 = *(*int32)(unsafe.Add(mBase, uint32(v13)))
												v73 = int32(0)
												F_RunObjectPostAlterHook(m, int32(2328), v72, v73, v73, v73)
												mBase = m.M
												v77 = m.ExcPending
												if v77 != 0 {
													return
												} else {
													m.G0 = v9 + int32(80)
													return
												}
											} else {
												m.G0 = v9 + int32(80)
												return
											}
										}
									}
								}
							}
						}
					} else {
						v70 = *(*int32)(unsafe.Add(mBase, _c_F_AlterForeignDataWrapperOwner_internal[0]))
						if v70 != 0 {
							v72 = *(*int32)(unsafe.Add(mBase, uint32(v13)))
							v73 = int32(0)
							F_RunObjectPostAlterHook(m, int32(2328), v72, v73, v73, v73)
							mBase = m.M
							v77 = m.ExcPending
							if v77 != 0 {
								return
							} else {
								m.G0 = v9 + int32(80)
								return
							}
						} else {
							m.G0 = v9 + int32(80)
							return
						}
					}
				}
			}
		} else {
			F_errstart_cold(m, int32(21), int32(0))
			mBase = m.M
			v84 = m.ExcPending
			if v84 != 0 {
				return
			} else {
				F_errcode(m, int32(16797828))
				mBase = m.M
				v87 = m.ExcPending
				if v87 != 0 {
					return
				} else {
					*(*int32)(unsafe.Add(mBase, uint32(v9)+16)) = v13 + int32(4)
					F_errmsg(m, int32(_a_F_AlterForeignDataWrapperOwner_internal_0), v9+int32(16))
					mBase = m.M
					v95 = m.ExcPending
					if v95 != 0 {
						return
					} else {
						F_errhint(m, int32(_a_F_AlterForeignDataWrapperOwner_internal_4), int32(0))
						mBase = m.M
						v99 = m.ExcPending
						if v99 != 0 {
							return
						} else {
							F_errfinish(m, int32(_a_F_AlterForeignDataWrapperOwner_internal_2), int32(234), int32(_a_F_AlterForeignDataWrapperOwner_internal_3))
							mBase = m.M
							v104 = m.ExcPending
							if v104 != 0 {
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
		}
	}
}
func F_AlterObjectNamespace_internal(m *base.Module, l0 int32, l1 int32, l2 int32) int32 {
	mBase := m.M
	_ = mBase
	var v12 int32
	_ = v12
	var v14 int32
	_ = v14
	var v16 int32
	_ = v16
	var v17 int32
	_ = v17
	var v20 int32
	_ = v20
	var v21 int32
	_ = v21
	var v22 int32
	_ = v22
	var v23 int32
	_ = v23
	var v24 int32
	_ = v24
	var v25 int32
	_ = v25
	var v26 int32
	_ = v26
	var v27 int32
	_ = v27
	var v28 int32
	_ = v28
	var v30 int32
	_ = v30
	var v31 int32
	_ = v31
	var v32 int32
	_ = v32
	var v34 int32
	_ = v34
	var v35 int32
	_ = v35
	var v36 int32
	_ = v36
	var v37 int32
	_ = v37
	var v38 int32
	_ = v38
	var v39 int32
	_ = v39
	var v42 int32
	_ = v42
	var v44 int32
	_ = v44
	var v45 int32
	_ = v45
	var v46 int32
	_ = v46
	var v49 int32
	_ = v49
	var v52 int32
	_ = v52
	var v53 int32
	_ = v53
	var v55 int32
	_ = v55
	var v56 int32
	_ = v56
	var v57 int32
	_ = v57
	var v61 int32
	_ = v61
	var v62 int32
	_ = v62
	var v64 int32
	_ = v64
	var v67 int32
	_ = v67
	var v69 int32
	_ = v69
	var v70 int32
	_ = v70
	var v74 int32
	_ = v74
	var v75 int32
	_ = v75
	var v77 int32
	_ = v77
	var v85 int32
	_ = v85
	var v86 int32
	_ = v86
	var v87 int32
	_ = v87
	var v90 int32
	_ = v90
	var v92 int32
	_ = v92
	var v97 int32
	_ = v97
	var v98 int32
	_ = v98
	var v103 int32
	_ = v103
	var v104 int32
	_ = v104
	var v105 int32
	_ = v105
	var v106 int32
	_ = v106
	var v109 int32
	_ = v109
	var v111 int32
	_ = v111
	var v114 int32
	_ = v114
	var v116 int32
	_ = v116
	var v117 int32
	_ = v117
	var v121 int32
	_ = v121
	var v122 int32
	_ = v122
	var v123 int32
	_ = v123
	var v124 int32
	_ = v124
	var v127 int32
	_ = v127
	var v131 int32
	_ = v131
	var v133 int32
	_ = v133
	var v134 int32
	_ = v134
	var v137 int32
	_ = v137
	var v138 int32
	_ = v138
	var v139 int32
	_ = v139
	var v140 int32
	_ = v140
	var v141 int32
	_ = v141
	var v142 int32
	_ = v142
	var v143 int32
	_ = v143
	var v144 int32
	_ = v144
	var v145 int32
	_ = v145
	var v146 int32
	_ = v146
	var v147 int32
	_ = v147
	var v148 int32
	_ = v148
	var v158 int32
	_ = v158
	var v159 int32
	_ = v159
	var v160 int32
	_ = v160
	var v162 int32
	_ = v162
	var v164 int32
	_ = v164
	var v166 int32
	_ = v166
	var v168 int32
	_ = v168
	var v170 int32
	_ = v170
	var v171 int32
	_ = v171
	var v175 int32
	_ = v175
	var v182 int32
	_ = v182
	var v186 int32
	_ = v186
	var v198 int32
	_ = v198
	var v199 int32
	_ = v199
	var v206 int32
	_ = v206
	var v211 int32
	_ = v211
	var v215 int32
	_ = v215
	var v218 int32
	_ = v218
	var v219 int32
	_ = v219
	var v220 int32
	_ = v220
	var v226 int32
	_ = v226
	var v231 int32
	_ = v231
	var v235 int32
	_ = v235
	var v241 int32
	_ = v241
	var v246 int32
	_ = v246
	v12 = m.G0
	v14 = v12 - int32(48)
	m.G0 = v14
	v16 = *(*int32)(unsafe.Add(mBase, uint32(l0)+56))
	v17 = F_get_object_catcache_oid(m, v16)
	mBase = m.M
	v20 = m.ExcPending
	if v20 != 0 {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	return int32(0)
L2:
	;
	v21 = F_get_object_catcache_name(m, v16)
	mBase = m.M
	v22 = m.ExcPending
	if v22 != 0 {
		goto L1
	} else {
		goto L3
	}
L3:
	;
	v23 = F_get_object_attnum_name(m, v16)
	mBase = m.M
	v24 = m.ExcPending
	if v24 != 0 {
		goto L1
	} else {
		goto L4
	}
L4:
	;
	v25 = F_get_object_attnum_namespace(m, v16)
	mBase = m.M
	v26 = m.ExcPending
	if v26 != 0 {
		goto L1
	} else {
		goto L5
	}
L5:
	;
	v27 = F_get_object_attnum_owner(m, v16)
	mBase = m.M
	v28 = m.ExcPending
	if v28 != 0 {
		goto L1
	} else {
		goto L6
	}
L6:
	;
	v30 = F_SearchSysCacheCopy(m, v17, l1, int32(0))
	mBase = m.M
	v31 = m.ExcPending
	if v31 != 0 {
		goto L1
	} else {
		goto L9
	}
L7:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v235 = m.ExcPending
	if v235 != 0 {
		goto L1
	} else {
		goto L77
	}
L8:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v215 = m.ExcPending
	if v215 != 0 {
		goto L1
	} else {
		goto L72
	}
L9:
	;
	if v30 != 0 {
		goto L10
	} else {
		goto L11
	}
L10:
	;
	v32 = *(*int32)(unsafe.Add(mBase, uint32(l0)+52))
	v34 = v14 + int32(47)
	v35 = F_heap_getattr_2(m, v30, v23, v32, v34)
	mBase = m.M
	v36 = m.ExcPending
	if v36 != 0 {
		goto L1
	} else {
		goto L13
	}
L11:
	;
	goto L12
L12:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v198 = m.ExcPending
	if v198 != 0 {
		goto L1
	} else {
		goto L69
	}
L13:
	;
	v37 = *(*int32)(unsafe.Add(mBase, uint32(l0)+52))
	v38 = F_heap_getattr_2(m, v30, v25, v37, v34)
	mBase = m.M
	v39 = m.ExcPending
	if v39 != 0 {
		goto L1
	} else {
		goto L16
	}
L14:
	;
	m.G0 = v14 + int32(48)
	return v38
L15:
	;
	v182 = int32(0)
	F_RunObjectPostAlterHook(m, v16, l1, v182, v182, v182)
	mBase = m.M
	v186 = m.ExcPending
	if v186 != 0 {
		goto L1
	} else {
		goto L68
	}
L16:
	;
	if l2 == v38 {
		goto L17
	} else {
		goto L18
	}
L17:
	;
	v42 = *(*int32)(unsafe.Add(mBase, _c_F_AlterObjectNamespace_internal[0]))
	if v42 != 0 {
		goto L15
	} else {
		goto L20
	}
L18:
	;
	goto L19
L19:
	;
	F_CheckSetNamespace(m, v38, l2)
	mBase = m.M
	v44 = m.ExcPending
	if v44 != 0 {
		goto L1
	} else {
		goto L21
	}
L20:
	;
	goto L14
L21:
	;
	v45 = F_superuser(m)
	mBase = m.M
	v46 = m.ExcPending
	if v46 != 0 {
		goto L1
	} else {
		goto L23
	}
L22:
	;
	if v16 <= int32(2752) {
		goto L40
	} else {
		goto L41
	}
L23:
	;
	if v45 != 0 {
		goto L22
	} else {
		goto L24
	}
L24:
	;
	if v27 <= int32(0) {
		goto L8
	} else {
		goto L25
	}
L25:
	;
	v49 = *(*int32)(unsafe.Add(mBase, uint32(l0)+52))
	v52 = F_heap_getattr_2(m, v30, v27, v49, v14+int32(47))
	mBase = m.M
	v53 = m.ExcPending
	if v53 != 0 {
		goto L1
	} else {
		goto L26
	}
L26:
	;
	v55 = *(*int32)(unsafe.Add(mBase, _c_F_AlterObjectNamespace_internal[1]))
	v56 = F_has_privs_of_role(m, v55, v52)
	mBase = m.M
	v57 = m.ExcPending
	if v57 != 0 {
		goto L1
	} else {
		goto L27
	}
L27:
	;
	if v56 == int32(0) {
		goto L28
	} else {
		goto L29
	}
L28:
	;
	v61 = F_get_object_type(m, v16, l1)
	mBase = m.M
	v62 = m.ExcPending
	if v62 != 0 {
		goto L1
	} else {
		goto L31
	}
L29:
	;
	goto L30
L30:
	;
	v67 = *(*int32)(unsafe.Add(mBase, _c_F_AlterObjectNamespace_internal[1]))
	v69 = F_object_aclcheck(m, int32(2615), l2, v67, int64(512))
	mBase = m.M
	v70 = m.ExcPending
	if v70 != 0 {
		goto L1
	} else {
		goto L33
	}
L31:
	;
	F_aclcheck_error(m, int32(2), v61, v35)
	mBase = m.M
	v64 = m.ExcPending
	if v64 != 0 {
		goto L1
	} else {
		goto L32
	}
L32:
	;
	goto L30
L33:
	;
	if v69 == int32(0) {
		goto L22
	} else {
		goto L34
	}
L34:
	;
	v74 = F_get_namespace_name(m, l2)
	mBase = m.M
	v75 = m.ExcPending
	if v75 != 0 {
		goto L1
	} else {
		goto L35
	}
L35:
	;
	F_aclcheck_error(m, v69, int32(36), v74)
	mBase = m.M
	v77 = m.ExcPending
	if v77 != 0 {
		goto L1
	} else {
		goto L36
	}
L36:
	;
	goto L22
L37:
	;
	v133 = *(*int32)(unsafe.Add(mBase, uint32(l0)+48))
	v134 = int32(*(*int16)(unsafe.Add(mBase, uint32(v133)+120)))
	v137 = F_palloc0(m, v134<<(uint(int32(2))%32))
	mBase = m.M
	v138 = m.ExcPending
	if v138 != 0 {
		goto L1
	} else {
		goto L57
	}
L38:
	;
	v122 = *(*int32)(unsafe.Add(mBase, uint32(v30)+16))
	v123 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v122)+22)))
	v124 = v122 + v123
	v127 = int32(*(*int16)(unsafe.Add(mBase, uint32(v124)+104)))
	F_IsThereFunctionInNamespace(m, v124+int32(4), v127, v124+int32(112), l2)
	mBase = m.M
	v131 = m.ExcPending
	if v131 != 0 {
		goto L1
	} else {
		goto L56
	}
L39:
	;
	if v21 < int32(0) {
		goto L37
	} else {
		goto L52
	}
L40:
	;
	if v16 == int32(1255) {
		goto L38
	} else {
		goto L43
	}
L41:
	;
	goto L42
L42:
	;
	if v16 != int32(2753) {
		goto L46
	} else {
		goto L47
	}
L43:
	;
	if v16 != int32(2616) {
		goto L39
	} else {
		goto L44
	}
L44:
	;
	v85 = *(*int32)(unsafe.Add(mBase, uint32(v30)+16))
	v86 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v85)+22)))
	v87 = v85 + v86
	v90 = *(*int32)(unsafe.Add(mBase, uint32(v87)+4))
	F_IsThereOpClassInNamespace(m, v87+int32(8), v90, l2)
	mBase = m.M
	v92 = m.ExcPending
	if v92 != 0 {
		goto L1
	} else {
		goto L45
	}
L45:
	;
	goto L37
L46:
	;
	if v16 != int32(3456) {
		goto L39
	} else {
		goto L49
	}
L47:
	;
	goto L48
L48:
	;
	v104 = *(*int32)(unsafe.Add(mBase, uint32(v30)+16))
	v105 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v104)+22)))
	v106 = v104 + v105
	v109 = *(*int32)(unsafe.Add(mBase, uint32(v106)+4))
	F_IsThereOpFamilyInNamespace(m, v106+int32(8), v109, l2)
	mBase = m.M
	v111 = m.ExcPending
	if v111 != 0 {
		goto L1
	} else {
		goto L51
	}
L49:
	;
	v97 = *(*int32)(unsafe.Add(mBase, uint32(v30)+16))
	v98 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v97)+22)))
	F_IsThereCollationInNamespace(m, v97+v98+int32(4), l2)
	mBase = m.M
	v103 = m.ExcPending
	if v103 != 0 {
		goto L1
	} else {
		goto L50
	}
L50:
	;
	goto L37
L51:
	;
	goto L37
L52:
	;
	v114 = int32(0)
	v116 = F_SearchSysCacheExists(m, v21, v35, l2, v114, v114)
	mBase = m.M
	v117 = m.ExcPending
	if v117 != 0 {
		goto L1
	} else {
		goto L53
	}
L53:
	;
	if v116 == int32(0) {
		goto L37
	} else {
		goto L54
	}
L54:
	;
	F_report_namespace_conflict(m, v16, v35, l2)
	mBase = m.M
	v121 = m.ExcPending
	if v121 != 0 {
		goto L1
	} else {
		goto L55
	}
L55:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L56:
	;
	goto L37
L57:
	;
	v139 = *(*int32)(unsafe.Add(mBase, uint32(l0)+48))
	v140 = int32(*(*int16)(unsafe.Add(mBase, uint32(v139)+120)))
	v141 = F_palloc0(m, v140)
	mBase = m.M
	v142 = m.ExcPending
	if v142 != 0 {
		goto L1
	} else {
		goto L58
	}
L58:
	;
	v143 = *(*int32)(unsafe.Add(mBase, uint32(l0)+48))
	v144 = int32(*(*int16)(unsafe.Add(mBase, uint32(v143)+120)))
	v145 = F_palloc0(m, v144)
	mBase = m.M
	v146 = m.ExcPending
	if v146 != 0 {
		goto L1
	} else {
		goto L59
	}
L59:
	;
	v147 = int32(1)
	v148 = v25 - v147
	*(*int32)(unsafe.Add(mBase, uint32(v137+v148<<(uint(int32(2))%32)))) = l2
	*(*uint8)(unsafe.Add(mBase, uint32(v145+v148))) = uint8(v147)
	v158 = *(*int32)(unsafe.Add(mBase, uint32(l0)+52))
	v159 = F_heap_modify_tuple(m, v30, v158, v137, v141, v145)
	mBase = m.M
	v160 = m.ExcPending
	if v160 != 0 {
		goto L1
	} else {
		goto L60
	}
L60:
	;
	F_CatalogTupleUpdate(m, l0, v30+int32(4), v159)
	mBase = m.M
	v162 = m.ExcPending
	if v162 != 0 {
		goto L1
	} else {
		goto L61
	}
L61:
	;
	F_pfree(m, v137)
	mBase = m.M
	v164 = m.ExcPending
	if v164 != 0 {
		goto L1
	} else {
		goto L62
	}
L62:
	;
	F_pfree(m, v141)
	mBase = m.M
	v166 = m.ExcPending
	if v166 != 0 {
		goto L1
	} else {
		goto L63
	}
L63:
	;
	F_pfree(m, v145)
	mBase = m.M
	v168 = m.ExcPending
	if v168 != 0 {
		goto L1
	} else {
		goto L64
	}
L64:
	;
	v170 = F_changeDependencyFor(m, v16, l1, int32(2615), v38, l2)
	mBase = m.M
	v171 = m.ExcPending
	if v171 != 0 {
		goto L1
	} else {
		goto L65
	}
L65:
	;
	if v170 != int32(1) {
		goto L7
	} else {
		goto L66
	}
L66:
	;
	v175 = *(*int32)(unsafe.Add(mBase, _c_F_AlterObjectNamespace_internal[0]))
	if v175 == int32(0) {
		goto L14
	} else {
		goto L67
	}
L67:
	;
	goto L15
L68:
	;
	goto L14
L69:
	;
	v199 = *(*int32)(unsafe.Add(mBase, uint32(l0)+48))
	*(*int32)(unsafe.Add(mBase, uint32(v14))) = l1
	*(*int32)(unsafe.Add(mBase, uint32(v14)+4)) = v199 + int32(4)
	F_errmsg_internal(m, int32(_a_F_AlterObjectNamespace_internal_4), v14)
	mBase = m.M
	v206 = m.ExcPending
	if v206 != 0 {
		goto L1
	} else {
		goto L70
	}
L70:
	;
	F_errfinish(m, int32(_a_F_AlterObjectNamespace_internal_1), int32(713), int32(_a_F_AlterObjectNamespace_internal_2))
	mBase = m.M
	v211 = m.ExcPending
	if v211 != 0 {
		goto L1
	} else {
		goto L71
	}
L71:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L72:
	;
	F_errcode(m, int32(16797828))
	mBase = m.M
	v218 = m.ExcPending
	if v218 != 0 {
		goto L1
	} else {
		goto L73
	}
L73:
	;
	v219 = F_getObjectDescriptionOids(m, v16, l1)
	mBase = m.M
	v220 = m.ExcPending
	if v220 != 0 {
		goto L1
	} else {
		goto L74
	}
L74:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v14)+32)) = v219
	F_errmsg(m, int32(_a_F_AlterObjectNamespace_internal_3), v14+int32(32))
	mBase = m.M
	v226 = m.ExcPending
	if v226 != 0 {
		goto L1
	} else {
		goto L75
	}
L75:
	;
	F_errfinish(m, int32(_a_F_AlterObjectNamespace_internal_1), int32(747), int32(_a_F_AlterObjectNamespace_internal_2))
	mBase = m.M
	v231 = m.ExcPending
	if v231 != 0 {
		goto L1
	} else {
		goto L76
	}
L76:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L77:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v14)+16)) = l1
	F_errmsg_internal(m, int32(_a_F_AlterObjectNamespace_internal_0), v14+int32(16))
	mBase = m.M
	v241 = m.ExcPending
	if v241 != 0 {
		goto L1
	} else {
		goto L78
	}
L78:
	;
	F_errfinish(m, int32(_a_F_AlterObjectNamespace_internal_1), int32(825), int32(_a_F_AlterObjectNamespace_internal_2))
	mBase = m.M
	v246 = m.ExcPending
	if v246 != 0 {
		goto L1
	} else {
		goto L79
	}
L79:
	;
	base.Wasm_trap_unreachable()
	for {
	}
}
func F_AlterSchemaOwner_internal(m *base.Module, l0 int32, l1 int32, l2 int32) {
	mBase := m.M
	_ = mBase
	var v7 int32
	_ = v7
	var v9 int32
	_ = v9
	var v11 int32
	_ = v11
	var v12 int32
	_ = v12
	var v13 int32
	_ = v13
	var v14 int32
	_ = v14
	var v17 int32
	_ = v17
	var v19 int32
	_ = v19
	var v20 int32
	_ = v20
	var v21 int32
	_ = v21
	var v29 int32
	_ = v29
	var v31 int32
	_ = v31
	var v33 int32
	_ = v33
	var v36 int32
	_ = v36
	var v38 int32
	_ = v38
	var v40 int32
	_ = v40
	var v41 int32
	_ = v41
	var v44 int32
	_ = v44
	var v45 int32
	_ = v45
	var v46 int32
	_ = v46
	var v48 int32
	_ = v48
	var v58 int32
	_ = v58
	var v59 int32
	_ = v59
	var v60 int32
	_ = v60
	var v63 int32
	_ = v63
	var v64 int32
	_ = v64
	var v65 int32
	_ = v65
	var v66 int32
	_ = v66
	var v67 int32
	_ = v67
	var v69 int32
	_ = v69
	var v71 int32
	_ = v71
	var v78 int32
	_ = v78
	var v79 int32
	_ = v79
	var v83 int32
	_ = v83
	var v85 int32
	_ = v85
	var v87 int32
	_ = v87
	var v89 int32
	_ = v89
	var v93 int32
	_ = v93
	var v95 int32
	_ = v95
	var v96 int32
	_ = v96
	var v100 int32
	_ = v100
	v7 = m.G0
	v9 = v7 - int32(32)
	m.G0 = v9
	v11 = *(*int32)(unsafe.Add(mBase, uint32(l0)+16))
	v12 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v11)+22)))
	v13 = v11 + v12
	v14 = *(*int32)(unsafe.Add(mBase, uint32(v13)+68))
	if l2 != v14 {
		v17 = *(*int32)(unsafe.Add(mBase, uint32(v13)))
		v19 = *(*int32)(unsafe.Add(mBase, _c_F_AlterSchemaOwner_internal[0]))
		v20 = F_object_ownercheck(m, int32(2615), v17, v19)
		mBase = m.M
		v21 = m.ExcPending
		if v21 != 0 {
			return
		} else {
			if v20 == int32(0) {
				F_aclcheck_error(m, int32(2), int32(36), v13+int32(4))
				mBase = m.M
				v29 = m.ExcPending
				if v29 != 0 {
					return
				} else {
					v31 = *(*int32)(unsafe.Add(mBase, _c_F_AlterSchemaOwner_internal[0]))
					F_check_can_set_role(m, v31, l2)
					mBase = m.M
					v33 = m.ExcPending
					if v33 != 0 {
						return
					} else {
						v36 = *(*int32)(unsafe.Add(mBase, _c_F_AlterSchemaOwner_internal[1]))
						v38 = *(*int32)(unsafe.Add(mBase, _c_F_AlterSchemaOwner_internal[0]))
						v40 = F_object_aclcheck(m, int32(1262), v36, v38, int64(512))
						mBase = m.M
						v41 = m.ExcPending
						if v41 != 0 {
							return
						} else {
							if v40 != 0 {
								v44 = *(*int32)(unsafe.Add(mBase, _c_F_AlterSchemaOwner_internal[1]))
								v45 = F_get_database_name(m, v44)
								mBase = m.M
								v46 = m.ExcPending
								if v46 != 0 {
									return
								} else {
									F_aclcheck_error(m, v40, int32(9), v45)
									mBase = m.M
									v48 = m.ExcPending
									if v48 != 0 {
										return
									} else {
										*(*int32)(unsafe.Add(mBase, uint32(v9)+8)) = int32(_a_F_AlterSchemaOwner_internal_0)
										*(*int32)(unsafe.Add(mBase, uint32(v9)+12)) = int32(0)
										*(*int32)(unsafe.Add(mBase, uint32(v9)+24)) = l2
										v58 = F_SysCacheGetAttr(m, int32(37), l0, int32(4), v9+int32(7))
										mBase = m.M
										v59 = m.ExcPending
										if v59 != 0 {
											return
										} else {
											v60 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v9)+7)))
											if v60 == int32(0) {
												v63 = F_pg_detoast_datum(m, v58)
												mBase = m.M
												v64 = m.ExcPending
												if v64 != 0 {
													return
												} else {
													v65 = *(*int32)(unsafe.Add(mBase, uint32(v13)+68))
													v66 = F_aclnewowner(m, v63, v65, l2)
													mBase = m.M
													v67 = m.ExcPending
													if v67 != 0 {
														return
													} else {
														*(*int32)(unsafe.Add(mBase, uint32(v9)+28)) = v66
														v69 = int32(1)
														*(*uint8)(unsafe.Add(mBase, uint32(v9)+11)) = uint8(v69)
														v71 = *(*int32)(unsafe.Add(mBase, uint32(l1)+52))
														v78 = F_heap_modify_tuple(m, l0, v71, v9+int32(16), v9+int32(12), v9+int32(8))
														mBase = m.M
														v79 = m.ExcPending
														if v79 != 0 {
															return
														} else {
															F_CatalogTupleUpdate(m, l1, v78+int32(4), v78)
															mBase = m.M
															v83 = m.ExcPending
															if v83 != 0 {
																return
															} else {
																F_pfree(m, v78)
																mBase = m.M
																v85 = m.ExcPending
																if v85 != 0 {
																	return
																} else {
																	v87 = *(*int32)(unsafe.Add(mBase, uint32(v13)))
																	F_changeDependencyOnOwner(m, int32(2615), v87, l2)
																	mBase = m.M
																	v89 = m.ExcPending
																	if v89 != 0 {
																		return
																	} else {
																		v93 = *(*int32)(unsafe.Add(mBase, _c_F_AlterSchemaOwner_internal[2]))
																		if v93 != 0 {
																			v95 = *(*int32)(unsafe.Add(mBase, uint32(v13)))
																			v96 = int32(0)
																			F_RunObjectPostAlterHook(m, int32(2615), v95, v96, v96, v96)
																			mBase = m.M
																			v100 = m.ExcPending
																			if v100 != 0 {
																				return
																			} else {
																				m.G0 = v9 + int32(32)
																				return
																			}
																		} else {
																			m.G0 = v9 + int32(32)
																			return
																		}
																	}
																}
															}
														}
													}
												}
											} else {
												v71 = *(*int32)(unsafe.Add(mBase, uint32(l1)+52))
												v78 = F_heap_modify_tuple(m, l0, v71, v9+int32(16), v9+int32(12), v9+int32(8))
												mBase = m.M
												v79 = m.ExcPending
												if v79 != 0 {
													return
												} else {
													F_CatalogTupleUpdate(m, l1, v78+int32(4), v78)
													mBase = m.M
													v83 = m.ExcPending
													if v83 != 0 {
														return
													} else {
														F_pfree(m, v78)
														mBase = m.M
														v85 = m.ExcPending
														if v85 != 0 {
															return
														} else {
															v87 = *(*int32)(unsafe.Add(mBase, uint32(v13)))
															F_changeDependencyOnOwner(m, int32(2615), v87, l2)
															mBase = m.M
															v89 = m.ExcPending
															if v89 != 0 {
																return
															} else {
																v93 = *(*int32)(unsafe.Add(mBase, _c_F_AlterSchemaOwner_internal[2]))
																if v93 != 0 {
																	v95 = *(*int32)(unsafe.Add(mBase, uint32(v13)))
																	v96 = int32(0)
																	F_RunObjectPostAlterHook(m, int32(2615), v95, v96, v96, v96)
																	mBase = m.M
																	v100 = m.ExcPending
																	if v100 != 0 {
																		return
																	} else {
																		m.G0 = v9 + int32(32)
																		return
																	}
																} else {
																	m.G0 = v9 + int32(32)
																	return
																}
															}
														}
													}
												}
											}
										}
									}
								}
							} else {
								*(*int32)(unsafe.Add(mBase, uint32(v9)+8)) = int32(_a_F_AlterSchemaOwner_internal_0)
								*(*int32)(unsafe.Add(mBase, uint32(v9)+12)) = int32(0)
								*(*int32)(unsafe.Add(mBase, uint32(v9)+24)) = l2
								v58 = F_SysCacheGetAttr(m, int32(37), l0, int32(4), v9+int32(7))
								mBase = m.M
								v59 = m.ExcPending
								if v59 != 0 {
									return
								} else {
									v60 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v9)+7)))
									if v60 == int32(0) {
										v63 = F_pg_detoast_datum(m, v58)
										mBase = m.M
										v64 = m.ExcPending
										if v64 != 0 {
											return
										} else {
											v65 = *(*int32)(unsafe.Add(mBase, uint32(v13)+68))
											v66 = F_aclnewowner(m, v63, v65, l2)
											mBase = m.M
											v67 = m.ExcPending
											if v67 != 0 {
												return
											} else {
												*(*int32)(unsafe.Add(mBase, uint32(v9)+28)) = v66
												v69 = int32(1)
												*(*uint8)(unsafe.Add(mBase, uint32(v9)+11)) = uint8(v69)
												v71 = *(*int32)(unsafe.Add(mBase, uint32(l1)+52))
												v78 = F_heap_modify_tuple(m, l0, v71, v9+int32(16), v9+int32(12), v9+int32(8))
												mBase = m.M
												v79 = m.ExcPending
												if v79 != 0 {
													return
												} else {
													F_CatalogTupleUpdate(m, l1, v78+int32(4), v78)
													mBase = m.M
													v83 = m.ExcPending
													if v83 != 0 {
														return
													} else {
														F_pfree(m, v78)
														mBase = m.M
														v85 = m.ExcPending
														if v85 != 0 {
															return
														} else {
															v87 = *(*int32)(unsafe.Add(mBase, uint32(v13)))
															F_changeDependencyOnOwner(m, int32(2615), v87, l2)
															mBase = m.M
															v89 = m.ExcPending
															if v89 != 0 {
																return
															} else {
																v93 = *(*int32)(unsafe.Add(mBase, _c_F_AlterSchemaOwner_internal[2]))
																if v93 != 0 {
																	v95 = *(*int32)(unsafe.Add(mBase, uint32(v13)))
																	v96 = int32(0)
																	F_RunObjectPostAlterHook(m, int32(2615), v95, v96, v96, v96)
																	mBase = m.M
																	v100 = m.ExcPending
																	if v100 != 0 {
																		return
																	} else {
																		m.G0 = v9 + int32(32)
																		return
																	}
																} else {
																	m.G0 = v9 + int32(32)
																	return
																}
															}
														}
													}
												}
											}
										}
									} else {
										v71 = *(*int32)(unsafe.Add(mBase, uint32(l1)+52))
										v78 = F_heap_modify_tuple(m, l0, v71, v9+int32(16), v9+int32(12), v9+int32(8))
										mBase = m.M
										v79 = m.ExcPending
										if v79 != 0 {
											return
										} else {
											F_CatalogTupleUpdate(m, l1, v78+int32(4), v78)
											mBase = m.M
											v83 = m.ExcPending
											if v83 != 0 {
												return
											} else {
												F_pfree(m, v78)
												mBase = m.M
												v85 = m.ExcPending
												if v85 != 0 {
													return
												} else {
													v87 = *(*int32)(unsafe.Add(mBase, uint32(v13)))
													F_changeDependencyOnOwner(m, int32(2615), v87, l2)
													mBase = m.M
													v89 = m.ExcPending
													if v89 != 0 {
														return
													} else {
														v93 = *(*int32)(unsafe.Add(mBase, _c_F_AlterSchemaOwner_internal[2]))
														if v93 != 0 {
															v95 = *(*int32)(unsafe.Add(mBase, uint32(v13)))
															v96 = int32(0)
															F_RunObjectPostAlterHook(m, int32(2615), v95, v96, v96, v96)
															mBase = m.M
															v100 = m.ExcPending
															if v100 != 0 {
																return
															} else {
																m.G0 = v9 + int32(32)
																return
															}
														} else {
															m.G0 = v9 + int32(32)
															return
														}
													}
												}
											}
										}
									}
								}
							}
						}
					}
				}
			} else {
				v31 = *(*int32)(unsafe.Add(mBase, _c_F_AlterSchemaOwner_internal[0]))
				F_check_can_set_role(m, v31, l2)
				mBase = m.M
				v33 = m.ExcPending
				if v33 != 0 {
					return
				} else {
					v36 = *(*int32)(unsafe.Add(mBase, _c_F_AlterSchemaOwner_internal[1]))
					v38 = *(*int32)(unsafe.Add(mBase, _c_F_AlterSchemaOwner_internal[0]))
					v40 = F_object_aclcheck(m, int32(1262), v36, v38, int64(512))
					mBase = m.M
					v41 = m.ExcPending
					if v41 != 0 {
						return
					} else {
						if v40 != 0 {
							v44 = *(*int32)(unsafe.Add(mBase, _c_F_AlterSchemaOwner_internal[1]))
							v45 = F_get_database_name(m, v44)
							mBase = m.M
							v46 = m.ExcPending
							if v46 != 0 {
								return
							} else {
								F_aclcheck_error(m, v40, int32(9), v45)
								mBase = m.M
								v48 = m.ExcPending
								if v48 != 0 {
									return
								} else {
									*(*int32)(unsafe.Add(mBase, uint32(v9)+8)) = int32(_a_F_AlterSchemaOwner_internal_0)
									*(*int32)(unsafe.Add(mBase, uint32(v9)+12)) = int32(0)
									*(*int32)(unsafe.Add(mBase, uint32(v9)+24)) = l2
									v58 = F_SysCacheGetAttr(m, int32(37), l0, int32(4), v9+int32(7))
									mBase = m.M
									v59 = m.ExcPending
									if v59 != 0 {
										return
									} else {
										v60 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v9)+7)))
										if v60 == int32(0) {
											v63 = F_pg_detoast_datum(m, v58)
											mBase = m.M
											v64 = m.ExcPending
											if v64 != 0 {
												return
											} else {
												v65 = *(*int32)(unsafe.Add(mBase, uint32(v13)+68))
												v66 = F_aclnewowner(m, v63, v65, l2)
												mBase = m.M
												v67 = m.ExcPending
												if v67 != 0 {
													return
												} else {
													*(*int32)(unsafe.Add(mBase, uint32(v9)+28)) = v66
													v69 = int32(1)
													*(*uint8)(unsafe.Add(mBase, uint32(v9)+11)) = uint8(v69)
													v71 = *(*int32)(unsafe.Add(mBase, uint32(l1)+52))
													v78 = F_heap_modify_tuple(m, l0, v71, v9+int32(16), v9+int32(12), v9+int32(8))
													mBase = m.M
													v79 = m.ExcPending
													if v79 != 0 {
														return
													} else {
														F_CatalogTupleUpdate(m, l1, v78+int32(4), v78)
														mBase = m.M
														v83 = m.ExcPending
														if v83 != 0 {
															return
														} else {
															F_pfree(m, v78)
															mBase = m.M
															v85 = m.ExcPending
															if v85 != 0 {
																return
															} else {
																v87 = *(*int32)(unsafe.Add(mBase, uint32(v13)))
																F_changeDependencyOnOwner(m, int32(2615), v87, l2)
																mBase = m.M
																v89 = m.ExcPending
																if v89 != 0 {
																	return
																} else {
																	v93 = *(*int32)(unsafe.Add(mBase, _c_F_AlterSchemaOwner_internal[2]))
																	if v93 != 0 {
																		v95 = *(*int32)(unsafe.Add(mBase, uint32(v13)))
																		v96 = int32(0)
																		F_RunObjectPostAlterHook(m, int32(2615), v95, v96, v96, v96)
																		mBase = m.M
																		v100 = m.ExcPending
																		if v100 != 0 {
																			return
																		} else {
																			m.G0 = v9 + int32(32)
																			return
																		}
																	} else {
																		m.G0 = v9 + int32(32)
																		return
																	}
																}
															}
														}
													}
												}
											}
										} else {
											v71 = *(*int32)(unsafe.Add(mBase, uint32(l1)+52))
											v78 = F_heap_modify_tuple(m, l0, v71, v9+int32(16), v9+int32(12), v9+int32(8))
											mBase = m.M
											v79 = m.ExcPending
											if v79 != 0 {
												return
											} else {
												F_CatalogTupleUpdate(m, l1, v78+int32(4), v78)
												mBase = m.M
												v83 = m.ExcPending
												if v83 != 0 {
													return
												} else {
													F_pfree(m, v78)
													mBase = m.M
													v85 = m.ExcPending
													if v85 != 0 {
														return
													} else {
														v87 = *(*int32)(unsafe.Add(mBase, uint32(v13)))
														F_changeDependencyOnOwner(m, int32(2615), v87, l2)
														mBase = m.M
														v89 = m.ExcPending
														if v89 != 0 {
															return
														} else {
															v93 = *(*int32)(unsafe.Add(mBase, _c_F_AlterSchemaOwner_internal[2]))
															if v93 != 0 {
																v95 = *(*int32)(unsafe.Add(mBase, uint32(v13)))
																v96 = int32(0)
																F_RunObjectPostAlterHook(m, int32(2615), v95, v96, v96, v96)
																mBase = m.M
																v100 = m.ExcPending
																if v100 != 0 {
																	return
																} else {
																	m.G0 = v9 + int32(32)
																	return
																}
															} else {
																m.G0 = v9 + int32(32)
																return
															}
														}
													}
												}
											}
										}
									}
								}
							}
						} else {
							*(*int32)(unsafe.Add(mBase, uint32(v9)+8)) = int32(_a_F_AlterSchemaOwner_internal_0)
							*(*int32)(unsafe.Add(mBase, uint32(v9)+12)) = int32(0)
							*(*int32)(unsafe.Add(mBase, uint32(v9)+24)) = l2
							v58 = F_SysCacheGetAttr(m, int32(37), l0, int32(4), v9+int32(7))
							mBase = m.M
							v59 = m.ExcPending
							if v59 != 0 {
								return
							} else {
								v60 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v9)+7)))
								if v60 == int32(0) {
									v63 = F_pg_detoast_datum(m, v58)
									mBase = m.M
									v64 = m.ExcPending
									if v64 != 0 {
										return
									} else {
										v65 = *(*int32)(unsafe.Add(mBase, uint32(v13)+68))
										v66 = F_aclnewowner(m, v63, v65, l2)
										mBase = m.M
										v67 = m.ExcPending
										if v67 != 0 {
											return
										} else {
											*(*int32)(unsafe.Add(mBase, uint32(v9)+28)) = v66
											v69 = int32(1)
											*(*uint8)(unsafe.Add(mBase, uint32(v9)+11)) = uint8(v69)
											v71 = *(*int32)(unsafe.Add(mBase, uint32(l1)+52))
											v78 = F_heap_modify_tuple(m, l0, v71, v9+int32(16), v9+int32(12), v9+int32(8))
											mBase = m.M
											v79 = m.ExcPending
											if v79 != 0 {
												return
											} else {
												F_CatalogTupleUpdate(m, l1, v78+int32(4), v78)
												mBase = m.M
												v83 = m.ExcPending
												if v83 != 0 {
													return
												} else {
													F_pfree(m, v78)
													mBase = m.M
													v85 = m.ExcPending
													if v85 != 0 {
														return
													} else {
														v87 = *(*int32)(unsafe.Add(mBase, uint32(v13)))
														F_changeDependencyOnOwner(m, int32(2615), v87, l2)
														mBase = m.M
														v89 = m.ExcPending
														if v89 != 0 {
															return
														} else {
															v93 = *(*int32)(unsafe.Add(mBase, _c_F_AlterSchemaOwner_internal[2]))
															if v93 != 0 {
																v95 = *(*int32)(unsafe.Add(mBase, uint32(v13)))
																v96 = int32(0)
																F_RunObjectPostAlterHook(m, int32(2615), v95, v96, v96, v96)
																mBase = m.M
																v100 = m.ExcPending
																if v100 != 0 {
																	return
																} else {
																	m.G0 = v9 + int32(32)
																	return
																}
															} else {
																m.G0 = v9 + int32(32)
																return
															}
														}
													}
												}
											}
										}
									}
								} else {
									v71 = *(*int32)(unsafe.Add(mBase, uint32(l1)+52))
									v78 = F_heap_modify_tuple(m, l0, v71, v9+int32(16), v9+int32(12), v9+int32(8))
									mBase = m.M
									v79 = m.ExcPending
									if v79 != 0 {
										return
									} else {
										F_CatalogTupleUpdate(m, l1, v78+int32(4), v78)
										mBase = m.M
										v83 = m.ExcPending
										if v83 != 0 {
											return
										} else {
											F_pfree(m, v78)
											mBase = m.M
											v85 = m.ExcPending
											if v85 != 0 {
												return
											} else {
												v87 = *(*int32)(unsafe.Add(mBase, uint32(v13)))
												F_changeDependencyOnOwner(m, int32(2615), v87, l2)
												mBase = m.M
												v89 = m.ExcPending
												if v89 != 0 {
													return
												} else {
													v93 = *(*int32)(unsafe.Add(mBase, _c_F_AlterSchemaOwner_internal[2]))
													if v93 != 0 {
														v95 = *(*int32)(unsafe.Add(mBase, uint32(v13)))
														v96 = int32(0)
														F_RunObjectPostAlterHook(m, int32(2615), v95, v96, v96, v96)
														mBase = m.M
														v100 = m.ExcPending
														if v100 != 0 {
															return
														} else {
															m.G0 = v9 + int32(32)
															return
														}
													} else {
														m.G0 = v9 + int32(32)
														return
													}
												}
											}
										}
									}
								}
							}
						}
					}
				}
			}
		}
	} else {
		v93 = *(*int32)(unsafe.Add(mBase, _c_F_AlterSchemaOwner_internal[2]))
		if v93 != 0 {
			v95 = *(*int32)(unsafe.Add(mBase, uint32(v13)))
			v96 = int32(0)
			F_RunObjectPostAlterHook(m, int32(2615), v95, v96, v96, v96)
			mBase = m.M
			v100 = m.ExcPending
			if v100 != 0 {
				return
			} else {
				m.G0 = v9 + int32(32)
				return
			}
		} else {
			m.G0 = v9 + int32(32)
			return
		}
	}
}
func F_AlterSubscriptionOwner_internal(m *base.Module, l0 int32, l1 int32, l2 int32) {
	mBase := m.M
	_ = mBase
	var v6 int32
	_ = v6
	var v7 int32
	_ = v7
	var v8 int32
	_ = v8
	var v9 int32
	_ = v9
	var v12 int32
	_ = v12
	var v14 int32
	_ = v14
	var v15 int32
	_ = v15
	var v16 int32
	_ = v16
	var v24 int32
	_ = v24
	var v25 int32
	_ = v25
	var v28 int32
	_ = v28
	var v29 int32
	_ = v29
	var v33 int32
	_ = v33
	var v35 int32
	_ = v35
	var v38 int32
	_ = v38
	var v40 int32
	_ = v40
	var v42 int32
	_ = v42
	var v43 int32
	_ = v43
	var v46 int32
	_ = v46
	var v47 int32
	_ = v47
	var v48 int32
	_ = v48
	var v50 int32
	_ = v50
	var v55 int32
	_ = v55
	var v57 int32
	_ = v57
	var v59 int32
	_ = v59
	var v61 int32
	_ = v61
	var v63 int32
	_ = v63
	var v64 int32
	_ = v64
	var v68 int32
	_ = v68
	var v70 int32
	_ = v70
	var v74 int32
	_ = v74
	var v76 int32
	_ = v76
	var v78 int32
	_ = v78
	var v83 int32
	_ = v83
	var v86 int32
	_ = v86
	var v90 int32
	_ = v90
	var v94 int32
	_ = v94
	var v99 int32
	_ = v99
	v6 = *(*int32)(unsafe.Add(mBase, uint32(l1)+16))
	v7 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v6)+22)))
	v8 = v6 + v7
	v9 = *(*int32)(unsafe.Add(mBase, uint32(v8)+80))
	if l2 != v9 {
		v12 = *(*int32)(unsafe.Add(mBase, uint32(v8)))
		v14 = *(*int32)(unsafe.Add(mBase, _c_F_AlterSubscriptionOwner_internal[0]))
		v15 = F_object_ownercheck(m, int32(_a_F_AlterSubscriptionOwner_internal_0), v12, v14)
		mBase = m.M
		v16 = m.ExcPending
		if v16 != 0 {
			return
		} else {
			if v15 == int32(0) {
				F_aclcheck_error(m, int32(2), int32(38), v8+int32(16))
				mBase = m.M
				v24 = m.ExcPending
				if v24 != 0 {
					return
				} else {
					v25 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v8)+89)))
					if v25 == int32(0) {
						v28 = F_superuser(m)
						mBase = m.M
						v29 = m.ExcPending
						if v29 != 0 {
							return
						} else {
							if v28 == int32(0) {
								F_errstart_cold(m, int32(21), int32(0))
								mBase = m.M
								v83 = m.ExcPending
								if v83 != 0 {
									return
								} else {
									F_errcode(m, int32(16797828))
									mBase = m.M
									v86 = m.ExcPending
									if v86 != 0 {
										return
									} else {
										F_errmsg(m, int32(_a_F_AlterSubscriptionOwner_internal_1), int32(0))
										mBase = m.M
										v90 = m.ExcPending
										if v90 != 0 {
											return
										} else {
											F_errhint(m, int32(_a_F_AlterSubscriptionOwner_internal_2), int32(0))
											mBase = m.M
											v94 = m.ExcPending
											if v94 != 0 {
												return
											} else {
												F_errfinish(m, int32(_a_F_AlterSubscriptionOwner_internal_3), int32(1992), int32(_a_F_AlterSubscriptionOwner_internal_4))
												mBase = m.M
												v99 = m.ExcPending
												if v99 != 0 {
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
								v33 = *(*int32)(unsafe.Add(mBase, _c_F_AlterSubscriptionOwner_internal[0]))
								F_check_can_set_role(m, v33, l2)
								mBase = m.M
								v35 = m.ExcPending
								if v35 != 0 {
									return
								} else {
									v38 = *(*int32)(unsafe.Add(mBase, _c_F_AlterSubscriptionOwner_internal[1]))
									v40 = *(*int32)(unsafe.Add(mBase, _c_F_AlterSubscriptionOwner_internal[0]))
									v42 = F_object_aclcheck(m, int32(1262), v38, v40, int64(512))
									mBase = m.M
									v43 = m.ExcPending
									if v43 != 0 {
										return
									} else {
										if v42 != 0 {
											v46 = *(*int32)(unsafe.Add(mBase, _c_F_AlterSubscriptionOwner_internal[1]))
											v47 = F_get_database_name(m, v46)
											mBase = m.M
											v48 = m.ExcPending
											if v48 != 0 {
												return
											} else {
												F_aclcheck_error(m, v42, int32(9), v47)
												mBase = m.M
												v50 = m.ExcPending
												if v50 != 0 {
													return
												} else {
													*(*int32)(unsafe.Add(mBase, uint32(v8)+80)) = l2
													F_CatalogTupleUpdate(m, l0, l1+int32(4), l1)
													mBase = m.M
													v55 = m.ExcPending
													if v55 != 0 {
														return
													} else {
														v57 = *(*int32)(unsafe.Add(mBase, uint32(v8)))
														F_changeDependencyOnOwner(m, int32(_a_F_AlterSubscriptionOwner_internal_0), v57, l2)
														mBase = m.M
														v59 = m.ExcPending
														if v59 != 0 {
															return
														} else {
															v61 = *(*int32)(unsafe.Add(mBase, _c_F_AlterSubscriptionOwner_internal[2]))
															if v61 != 0 {
																v63 = *(*int32)(unsafe.Add(mBase, uint32(v8)))
																v64 = int32(0)
																F_RunObjectPostAlterHook(m, int32(_a_F_AlterSubscriptionOwner_internal_0), v63, v64, v64, v64)
																mBase = m.M
																v68 = m.ExcPending
																if v68 != 0 {
																	return
																} else {
																	v70 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_AlterSubscriptionOwner_internal[3])))
																	if v70 == int32(0) {
																		v74 = int32(1)
																		*(*uint8)(unsafe.Add(mBase, _c_F_AlterSubscriptionOwner_internal[3])) = uint8(v74)
																	} else {
																	}
																	v76 = *(*int32)(unsafe.Add(mBase, uint32(v8)))
																	F_LogicalRepWorkersWakeupAtCommit(m, v76)
																	mBase = m.M
																	v78 = m.ExcPending
																	if v78 != 0 {
																		return
																	} else {
																		return
																	}
																}
															} else {
																v70 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_AlterSubscriptionOwner_internal[3])))
																if v70 == int32(0) {
																	v74 = int32(1)
																	*(*uint8)(unsafe.Add(mBase, _c_F_AlterSubscriptionOwner_internal[3])) = uint8(v74)
																} else {
																}
																v76 = *(*int32)(unsafe.Add(mBase, uint32(v8)))
																F_LogicalRepWorkersWakeupAtCommit(m, v76)
																mBase = m.M
																v78 = m.ExcPending
																if v78 != 0 {
																	return
																} else {
																	return
																}
															}
														}
													}
												}
											}
										} else {
											*(*int32)(unsafe.Add(mBase, uint32(v8)+80)) = l2
											F_CatalogTupleUpdate(m, l0, l1+int32(4), l1)
											mBase = m.M
											v55 = m.ExcPending
											if v55 != 0 {
												return
											} else {
												v57 = *(*int32)(unsafe.Add(mBase, uint32(v8)))
												F_changeDependencyOnOwner(m, int32(_a_F_AlterSubscriptionOwner_internal_0), v57, l2)
												mBase = m.M
												v59 = m.ExcPending
												if v59 != 0 {
													return
												} else {
													v61 = *(*int32)(unsafe.Add(mBase, _c_F_AlterSubscriptionOwner_internal[2]))
													if v61 != 0 {
														v63 = *(*int32)(unsafe.Add(mBase, uint32(v8)))
														v64 = int32(0)
														F_RunObjectPostAlterHook(m, int32(_a_F_AlterSubscriptionOwner_internal_0), v63, v64, v64, v64)
														mBase = m.M
														v68 = m.ExcPending
														if v68 != 0 {
															return
														} else {
															v70 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_AlterSubscriptionOwner_internal[3])))
															if v70 == int32(0) {
																v74 = int32(1)
																*(*uint8)(unsafe.Add(mBase, _c_F_AlterSubscriptionOwner_internal[3])) = uint8(v74)
															} else {
															}
															v76 = *(*int32)(unsafe.Add(mBase, uint32(v8)))
															F_LogicalRepWorkersWakeupAtCommit(m, v76)
															mBase = m.M
															v78 = m.ExcPending
															if v78 != 0 {
																return
															} else {
																return
															}
														}
													} else {
														v70 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_AlterSubscriptionOwner_internal[3])))
														if v70 == int32(0) {
															v74 = int32(1)
															*(*uint8)(unsafe.Add(mBase, _c_F_AlterSubscriptionOwner_internal[3])) = uint8(v74)
														} else {
														}
														v76 = *(*int32)(unsafe.Add(mBase, uint32(v8)))
														F_LogicalRepWorkersWakeupAtCommit(m, v76)
														mBase = m.M
														v78 = m.ExcPending
														if v78 != 0 {
															return
														} else {
															return
														}
													}
												}
											}
										}
									}
								}
							}
						}
					} else {
						v33 = *(*int32)(unsafe.Add(mBase, _c_F_AlterSubscriptionOwner_internal[0]))
						F_check_can_set_role(m, v33, l2)
						mBase = m.M
						v35 = m.ExcPending
						if v35 != 0 {
							return
						} else {
							v38 = *(*int32)(unsafe.Add(mBase, _c_F_AlterSubscriptionOwner_internal[1]))
							v40 = *(*int32)(unsafe.Add(mBase, _c_F_AlterSubscriptionOwner_internal[0]))
							v42 = F_object_aclcheck(m, int32(1262), v38, v40, int64(512))
							mBase = m.M
							v43 = m.ExcPending
							if v43 != 0 {
								return
							} else {
								if v42 != 0 {
									v46 = *(*int32)(unsafe.Add(mBase, _c_F_AlterSubscriptionOwner_internal[1]))
									v47 = F_get_database_name(m, v46)
									mBase = m.M
									v48 = m.ExcPending
									if v48 != 0 {
										return
									} else {
										F_aclcheck_error(m, v42, int32(9), v47)
										mBase = m.M
										v50 = m.ExcPending
										if v50 != 0 {
											return
										} else {
											*(*int32)(unsafe.Add(mBase, uint32(v8)+80)) = l2
											F_CatalogTupleUpdate(m, l0, l1+int32(4), l1)
											mBase = m.M
											v55 = m.ExcPending
											if v55 != 0 {
												return
											} else {
												v57 = *(*int32)(unsafe.Add(mBase, uint32(v8)))
												F_changeDependencyOnOwner(m, int32(_a_F_AlterSubscriptionOwner_internal_0), v57, l2)
												mBase = m.M
												v59 = m.ExcPending
												if v59 != 0 {
													return
												} else {
													v61 = *(*int32)(unsafe.Add(mBase, _c_F_AlterSubscriptionOwner_internal[2]))
													if v61 != 0 {
														v63 = *(*int32)(unsafe.Add(mBase, uint32(v8)))
														v64 = int32(0)
														F_RunObjectPostAlterHook(m, int32(_a_F_AlterSubscriptionOwner_internal_0), v63, v64, v64, v64)
														mBase = m.M
														v68 = m.ExcPending
														if v68 != 0 {
															return
														} else {
															v70 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_AlterSubscriptionOwner_internal[3])))
															if v70 == int32(0) {
																v74 = int32(1)
																*(*uint8)(unsafe.Add(mBase, _c_F_AlterSubscriptionOwner_internal[3])) = uint8(v74)
															} else {
															}
															v76 = *(*int32)(unsafe.Add(mBase, uint32(v8)))
															F_LogicalRepWorkersWakeupAtCommit(m, v76)
															mBase = m.M
															v78 = m.ExcPending
															if v78 != 0 {
																return
															} else {
																return
															}
														}
													} else {
														v70 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_AlterSubscriptionOwner_internal[3])))
														if v70 == int32(0) {
															v74 = int32(1)
															*(*uint8)(unsafe.Add(mBase, _c_F_AlterSubscriptionOwner_internal[3])) = uint8(v74)
														} else {
														}
														v76 = *(*int32)(unsafe.Add(mBase, uint32(v8)))
														F_LogicalRepWorkersWakeupAtCommit(m, v76)
														mBase = m.M
														v78 = m.ExcPending
														if v78 != 0 {
															return
														} else {
															return
														}
													}
												}
											}
										}
									}
								} else {
									*(*int32)(unsafe.Add(mBase, uint32(v8)+80)) = l2
									F_CatalogTupleUpdate(m, l0, l1+int32(4), l1)
									mBase = m.M
									v55 = m.ExcPending
									if v55 != 0 {
										return
									} else {
										v57 = *(*int32)(unsafe.Add(mBase, uint32(v8)))
										F_changeDependencyOnOwner(m, int32(_a_F_AlterSubscriptionOwner_internal_0), v57, l2)
										mBase = m.M
										v59 = m.ExcPending
										if v59 != 0 {
											return
										} else {
											v61 = *(*int32)(unsafe.Add(mBase, _c_F_AlterSubscriptionOwner_internal[2]))
											if v61 != 0 {
												v63 = *(*int32)(unsafe.Add(mBase, uint32(v8)))
												v64 = int32(0)
												F_RunObjectPostAlterHook(m, int32(_a_F_AlterSubscriptionOwner_internal_0), v63, v64, v64, v64)
												mBase = m.M
												v68 = m.ExcPending
												if v68 != 0 {
													return
												} else {
													v70 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_AlterSubscriptionOwner_internal[3])))
													if v70 == int32(0) {
														v74 = int32(1)
														*(*uint8)(unsafe.Add(mBase, _c_F_AlterSubscriptionOwner_internal[3])) = uint8(v74)
													} else {
													}
													v76 = *(*int32)(unsafe.Add(mBase, uint32(v8)))
													F_LogicalRepWorkersWakeupAtCommit(m, v76)
													mBase = m.M
													v78 = m.ExcPending
													if v78 != 0 {
														return
													} else {
														return
													}
												}
											} else {
												v70 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_AlterSubscriptionOwner_internal[3])))
												if v70 == int32(0) {
													v74 = int32(1)
													*(*uint8)(unsafe.Add(mBase, _c_F_AlterSubscriptionOwner_internal[3])) = uint8(v74)
												} else {
												}
												v76 = *(*int32)(unsafe.Add(mBase, uint32(v8)))
												F_LogicalRepWorkersWakeupAtCommit(m, v76)
												mBase = m.M
												v78 = m.ExcPending
												if v78 != 0 {
													return
												} else {
													return
												}
											}
										}
									}
								}
							}
						}
					}
				}
			} else {
				v25 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v8)+89)))
				if v25 == int32(0) {
					v28 = F_superuser(m)
					mBase = m.M
					v29 = m.ExcPending
					if v29 != 0 {
						return
					} else {
						if v28 == int32(0) {
							F_errstart_cold(m, int32(21), int32(0))
							mBase = m.M
							v83 = m.ExcPending
							if v83 != 0 {
								return
							} else {
								F_errcode(m, int32(16797828))
								mBase = m.M
								v86 = m.ExcPending
								if v86 != 0 {
									return
								} else {
									F_errmsg(m, int32(_a_F_AlterSubscriptionOwner_internal_1), int32(0))
									mBase = m.M
									v90 = m.ExcPending
									if v90 != 0 {
										return
									} else {
										F_errhint(m, int32(_a_F_AlterSubscriptionOwner_internal_2), int32(0))
										mBase = m.M
										v94 = m.ExcPending
										if v94 != 0 {
											return
										} else {
											F_errfinish(m, int32(_a_F_AlterSubscriptionOwner_internal_3), int32(1992), int32(_a_F_AlterSubscriptionOwner_internal_4))
											mBase = m.M
											v99 = m.ExcPending
											if v99 != 0 {
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
							v33 = *(*int32)(unsafe.Add(mBase, _c_F_AlterSubscriptionOwner_internal[0]))
							F_check_can_set_role(m, v33, l2)
							mBase = m.M
							v35 = m.ExcPending
							if v35 != 0 {
								return
							} else {
								v38 = *(*int32)(unsafe.Add(mBase, _c_F_AlterSubscriptionOwner_internal[1]))
								v40 = *(*int32)(unsafe.Add(mBase, _c_F_AlterSubscriptionOwner_internal[0]))
								v42 = F_object_aclcheck(m, int32(1262), v38, v40, int64(512))
								mBase = m.M
								v43 = m.ExcPending
								if v43 != 0 {
									return
								} else {
									if v42 != 0 {
										v46 = *(*int32)(unsafe.Add(mBase, _c_F_AlterSubscriptionOwner_internal[1]))
										v47 = F_get_database_name(m, v46)
										mBase = m.M
										v48 = m.ExcPending
										if v48 != 0 {
											return
										} else {
											F_aclcheck_error(m, v42, int32(9), v47)
											mBase = m.M
											v50 = m.ExcPending
											if v50 != 0 {
												return
											} else {
												*(*int32)(unsafe.Add(mBase, uint32(v8)+80)) = l2
												F_CatalogTupleUpdate(m, l0, l1+int32(4), l1)
												mBase = m.M
												v55 = m.ExcPending
												if v55 != 0 {
													return
												} else {
													v57 = *(*int32)(unsafe.Add(mBase, uint32(v8)))
													F_changeDependencyOnOwner(m, int32(_a_F_AlterSubscriptionOwner_internal_0), v57, l2)
													mBase = m.M
													v59 = m.ExcPending
													if v59 != 0 {
														return
													} else {
														v61 = *(*int32)(unsafe.Add(mBase, _c_F_AlterSubscriptionOwner_internal[2]))
														if v61 != 0 {
															v63 = *(*int32)(unsafe.Add(mBase, uint32(v8)))
															v64 = int32(0)
															F_RunObjectPostAlterHook(m, int32(_a_F_AlterSubscriptionOwner_internal_0), v63, v64, v64, v64)
															mBase = m.M
															v68 = m.ExcPending
															if v68 != 0 {
																return
															} else {
																v70 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_AlterSubscriptionOwner_internal[3])))
																if v70 == int32(0) {
																	v74 = int32(1)
																	*(*uint8)(unsafe.Add(mBase, _c_F_AlterSubscriptionOwner_internal[3])) = uint8(v74)
																} else {
																}
																v76 = *(*int32)(unsafe.Add(mBase, uint32(v8)))
																F_LogicalRepWorkersWakeupAtCommit(m, v76)
																mBase = m.M
																v78 = m.ExcPending
																if v78 != 0 {
																	return
																} else {
																	return
																}
															}
														} else {
															v70 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_AlterSubscriptionOwner_internal[3])))
															if v70 == int32(0) {
																v74 = int32(1)
																*(*uint8)(unsafe.Add(mBase, _c_F_AlterSubscriptionOwner_internal[3])) = uint8(v74)
															} else {
															}
															v76 = *(*int32)(unsafe.Add(mBase, uint32(v8)))
															F_LogicalRepWorkersWakeupAtCommit(m, v76)
															mBase = m.M
															v78 = m.ExcPending
															if v78 != 0 {
																return
															} else {
																return
															}
														}
													}
												}
											}
										}
									} else {
										*(*int32)(unsafe.Add(mBase, uint32(v8)+80)) = l2
										F_CatalogTupleUpdate(m, l0, l1+int32(4), l1)
										mBase = m.M
										v55 = m.ExcPending
										if v55 != 0 {
											return
										} else {
											v57 = *(*int32)(unsafe.Add(mBase, uint32(v8)))
											F_changeDependencyOnOwner(m, int32(_a_F_AlterSubscriptionOwner_internal_0), v57, l2)
											mBase = m.M
											v59 = m.ExcPending
											if v59 != 0 {
												return
											} else {
												v61 = *(*int32)(unsafe.Add(mBase, _c_F_AlterSubscriptionOwner_internal[2]))
												if v61 != 0 {
													v63 = *(*int32)(unsafe.Add(mBase, uint32(v8)))
													v64 = int32(0)
													F_RunObjectPostAlterHook(m, int32(_a_F_AlterSubscriptionOwner_internal_0), v63, v64, v64, v64)
													mBase = m.M
													v68 = m.ExcPending
													if v68 != 0 {
														return
													} else {
														v70 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_AlterSubscriptionOwner_internal[3])))
														if v70 == int32(0) {
															v74 = int32(1)
															*(*uint8)(unsafe.Add(mBase, _c_F_AlterSubscriptionOwner_internal[3])) = uint8(v74)
														} else {
														}
														v76 = *(*int32)(unsafe.Add(mBase, uint32(v8)))
														F_LogicalRepWorkersWakeupAtCommit(m, v76)
														mBase = m.M
														v78 = m.ExcPending
														if v78 != 0 {
															return
														} else {
															return
														}
													}
												} else {
													v70 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_AlterSubscriptionOwner_internal[3])))
													if v70 == int32(0) {
														v74 = int32(1)
														*(*uint8)(unsafe.Add(mBase, _c_F_AlterSubscriptionOwner_internal[3])) = uint8(v74)
													} else {
													}
													v76 = *(*int32)(unsafe.Add(mBase, uint32(v8)))
													F_LogicalRepWorkersWakeupAtCommit(m, v76)
													mBase = m.M
													v78 = m.ExcPending
													if v78 != 0 {
														return
													} else {
														return
													}
												}
											}
										}
									}
								}
							}
						}
					}
				} else {
					v33 = *(*int32)(unsafe.Add(mBase, _c_F_AlterSubscriptionOwner_internal[0]))
					F_check_can_set_role(m, v33, l2)
					mBase = m.M
					v35 = m.ExcPending
					if v35 != 0 {
						return
					} else {
						v38 = *(*int32)(unsafe.Add(mBase, _c_F_AlterSubscriptionOwner_internal[1]))
						v40 = *(*int32)(unsafe.Add(mBase, _c_F_AlterSubscriptionOwner_internal[0]))
						v42 = F_object_aclcheck(m, int32(1262), v38, v40, int64(512))
						mBase = m.M
						v43 = m.ExcPending
						if v43 != 0 {
							return
						} else {
							if v42 != 0 {
								v46 = *(*int32)(unsafe.Add(mBase, _c_F_AlterSubscriptionOwner_internal[1]))
								v47 = F_get_database_name(m, v46)
								mBase = m.M
								v48 = m.ExcPending
								if v48 != 0 {
									return
								} else {
									F_aclcheck_error(m, v42, int32(9), v47)
									mBase = m.M
									v50 = m.ExcPending
									if v50 != 0 {
										return
									} else {
										*(*int32)(unsafe.Add(mBase, uint32(v8)+80)) = l2
										F_CatalogTupleUpdate(m, l0, l1+int32(4), l1)
										mBase = m.M
										v55 = m.ExcPending
										if v55 != 0 {
											return
										} else {
											v57 = *(*int32)(unsafe.Add(mBase, uint32(v8)))
											F_changeDependencyOnOwner(m, int32(_a_F_AlterSubscriptionOwner_internal_0), v57, l2)
											mBase = m.M
											v59 = m.ExcPending
											if v59 != 0 {
												return
											} else {
												v61 = *(*int32)(unsafe.Add(mBase, _c_F_AlterSubscriptionOwner_internal[2]))
												if v61 != 0 {
													v63 = *(*int32)(unsafe.Add(mBase, uint32(v8)))
													v64 = int32(0)
													F_RunObjectPostAlterHook(m, int32(_a_F_AlterSubscriptionOwner_internal_0), v63, v64, v64, v64)
													mBase = m.M
													v68 = m.ExcPending
													if v68 != 0 {
														return
													} else {
														v70 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_AlterSubscriptionOwner_internal[3])))
														if v70 == int32(0) {
															v74 = int32(1)
															*(*uint8)(unsafe.Add(mBase, _c_F_AlterSubscriptionOwner_internal[3])) = uint8(v74)
														} else {
														}
														v76 = *(*int32)(unsafe.Add(mBase, uint32(v8)))
														F_LogicalRepWorkersWakeupAtCommit(m, v76)
														mBase = m.M
														v78 = m.ExcPending
														if v78 != 0 {
															return
														} else {
															return
														}
													}
												} else {
													v70 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_AlterSubscriptionOwner_internal[3])))
													if v70 == int32(0) {
														v74 = int32(1)
														*(*uint8)(unsafe.Add(mBase, _c_F_AlterSubscriptionOwner_internal[3])) = uint8(v74)
													} else {
													}
													v76 = *(*int32)(unsafe.Add(mBase, uint32(v8)))
													F_LogicalRepWorkersWakeupAtCommit(m, v76)
													mBase = m.M
													v78 = m.ExcPending
													if v78 != 0 {
														return
													} else {
														return
													}
												}
											}
										}
									}
								}
							} else {
								*(*int32)(unsafe.Add(mBase, uint32(v8)+80)) = l2
								F_CatalogTupleUpdate(m, l0, l1+int32(4), l1)
								mBase = m.M
								v55 = m.ExcPending
								if v55 != 0 {
									return
								} else {
									v57 = *(*int32)(unsafe.Add(mBase, uint32(v8)))
									F_changeDependencyOnOwner(m, int32(_a_F_AlterSubscriptionOwner_internal_0), v57, l2)
									mBase = m.M
									v59 = m.ExcPending
									if v59 != 0 {
										return
									} else {
										v61 = *(*int32)(unsafe.Add(mBase, _c_F_AlterSubscriptionOwner_internal[2]))
										if v61 != 0 {
											v63 = *(*int32)(unsafe.Add(mBase, uint32(v8)))
											v64 = int32(0)
											F_RunObjectPostAlterHook(m, int32(_a_F_AlterSubscriptionOwner_internal_0), v63, v64, v64, v64)
											mBase = m.M
											v68 = m.ExcPending
											if v68 != 0 {
												return
											} else {
												v70 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_AlterSubscriptionOwner_internal[3])))
												if v70 == int32(0) {
													v74 = int32(1)
													*(*uint8)(unsafe.Add(mBase, _c_F_AlterSubscriptionOwner_internal[3])) = uint8(v74)
												} else {
												}
												v76 = *(*int32)(unsafe.Add(mBase, uint32(v8)))
												F_LogicalRepWorkersWakeupAtCommit(m, v76)
												mBase = m.M
												v78 = m.ExcPending
												if v78 != 0 {
													return
												} else {
													return
												}
											}
										} else {
											v70 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_AlterSubscriptionOwner_internal[3])))
											if v70 == int32(0) {
												v74 = int32(1)
												*(*uint8)(unsafe.Add(mBase, _c_F_AlterSubscriptionOwner_internal[3])) = uint8(v74)
											} else {
											}
											v76 = *(*int32)(unsafe.Add(mBase, uint32(v8)))
											F_LogicalRepWorkersWakeupAtCommit(m, v76)
											mBase = m.M
											v78 = m.ExcPending
											if v78 != 0 {
												return
											} else {
												return
											}
										}
									}
								}
							}
						}
					}
				}
			}
		}
	} else {
		return
	}
}
func F_AlterTypeNamespace_oid(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) int32 {
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
	var v17 int32
	_ = v17
	var v22 int32
	_ = v22
	var v23 int32
	_ = v23
	var v24 int32
	_ = v24
	var v27 int32
	_ = v27
	var v28 int32
	_ = v28
	var v34 int32
	_ = v34
	var v37 int32
	_ = v37
	var v38 int32
	_ = v38
	var v39 int32
	_ = v39
	var v45 int32
	_ = v45
	var v46 int32
	_ = v46
	var v47 int32
	_ = v47
	var v51 int32
	_ = v51
	var v56 int32
	_ = v56
	var v59 int32
	_ = v59
	var v60 int32
	_ = v60
	var v61 int32
	_ = v61
	v7 = m.G0
	v9 = v7 - int32(32)
	m.G0 = v9
	v13 = *(*int32)(unsafe.Add(mBase, _c_F_AlterTypeNamespace_oid[0]))
	v14 = F_object_ownercheck(m, int32(1247), l0, v13)
	mBase = m.M
	v17 = m.ExcPending
	if v17 != 0 {
		return int32(0)
	} else {
		if v14 == int32(0) {
			F_aclcheck_error_type(m, int32(2), l0)
			mBase = m.M
			v22 = m.ExcPending
			if v22 != 0 {
				return int32(0)
			} else {
				v23 = F_get_element_type(m, l0)
				mBase = m.M
				v24 = m.ExcPending
				if v24 != 0 {
					return int32(0)
				} else {
					if v23 == int32(0) {
						v59 = F_AlterTypeNamespaceInternal(m, l0, l1, int32(0), l2, int32(1), l3)
						mBase = m.M
						v60 = m.ExcPending
						if v60 != 0 {
							return int32(0)
						} else {
							v61 = v59
							m.G0 = v9 + int32(32)
							return v61
						}
					} else {
						v27 = F_get_array_type(m, v23)
						mBase = m.M
						v28 = m.ExcPending
						if v28 != 0 {
							return int32(0)
						} else {
							if v27 != l0 {
								v59 = F_AlterTypeNamespaceInternal(m, l0, l1, int32(0), l2, int32(1), l3)
								mBase = m.M
								v60 = m.ExcPending
								if v60 != 0 {
									return int32(0)
								} else {
									v61 = v59
									m.G0 = v9 + int32(32)
									return v61
								}
							} else {
								if l2 != 0 {
									v61 = int32(0)
									m.G0 = v9 + int32(32)
									return v61
								} else {
									F_errstart_cold(m, int32(21), int32(0))
									mBase = m.M
									v34 = m.ExcPending
									if v34 != 0 {
										return int32(0)
									} else {
										F_errcode(m, int32(151027844))
										mBase = m.M
										v37 = m.ExcPending
										if v37 != 0 {
											return int32(0)
										} else {
											v38 = F_format_type_be(m, l0)
											mBase = m.M
											v39 = m.ExcPending
											if v39 != 0 {
												return int32(0)
											} else {
												*(*int32)(unsafe.Add(mBase, uint32(v9)+16)) = v38
												F_errmsg(m, int32(_a_F_AlterTypeNamespace_oid_0), v9+int32(16))
												mBase = m.M
												v45 = m.ExcPending
												if v45 != 0 {
													return int32(0)
												} else {
													v46 = F_format_type_be(m, v23)
													mBase = m.M
													v47 = m.ExcPending
													if v47 != 0 {
														return int32(0)
													} else {
														*(*int32)(unsafe.Add(mBase, uint32(v9))) = v46
														F_errhint(m, int32(_a_F_AlterTypeNamespace_oid_1), v9)
														mBase = m.M
														v51 = m.ExcPending
														if v51 != 0 {
															return int32(0)
														} else {
															F_errfinish(m, int32(_a_F_AlterTypeNamespace_oid_2), int32(_a_F_AlterTypeNamespace_oid_3), int32(_a_F_AlterTypeNamespace_oid_4))
															mBase = m.M
															v56 = m.ExcPending
															if v56 != 0 {
																return int32(0)
															} else {
																base.Wasm_trap_unreachable()
																for {
																}
															}
														}
													}
												}
											}
										}
									}
								}
							}
						}
					}
				}
			}
		} else {
			v23 = F_get_element_type(m, l0)
			mBase = m.M
			v24 = m.ExcPending
			if v24 != 0 {
				return int32(0)
			} else {
				if v23 == int32(0) {
					v59 = F_AlterTypeNamespaceInternal(m, l0, l1, int32(0), l2, int32(1), l3)
					mBase = m.M
					v60 = m.ExcPending
					if v60 != 0 {
						return int32(0)
					} else {
						v61 = v59
						m.G0 = v9 + int32(32)
						return v61
					}
				} else {
					v27 = F_get_array_type(m, v23)
					mBase = m.M
					v28 = m.ExcPending
					if v28 != 0 {
						return int32(0)
					} else {
						if v27 != l0 {
							v59 = F_AlterTypeNamespaceInternal(m, l0, l1, int32(0), l2, int32(1), l3)
							mBase = m.M
							v60 = m.ExcPending
							if v60 != 0 {
								return int32(0)
							} else {
								v61 = v59
								m.G0 = v9 + int32(32)
								return v61
							}
						} else {
							if l2 != 0 {
								v61 = int32(0)
								m.G0 = v9 + int32(32)
								return v61
							} else {
								F_errstart_cold(m, int32(21), int32(0))
								mBase = m.M
								v34 = m.ExcPending
								if v34 != 0 {
									return int32(0)
								} else {
									F_errcode(m, int32(151027844))
									mBase = m.M
									v37 = m.ExcPending
									if v37 != 0 {
										return int32(0)
									} else {
										v38 = F_format_type_be(m, l0)
										mBase = m.M
										v39 = m.ExcPending
										if v39 != 0 {
											return int32(0)
										} else {
											*(*int32)(unsafe.Add(mBase, uint32(v9)+16)) = v38
											F_errmsg(m, int32(_a_F_AlterTypeNamespace_oid_0), v9+int32(16))
											mBase = m.M
											v45 = m.ExcPending
											if v45 != 0 {
												return int32(0)
											} else {
												v46 = F_format_type_be(m, v23)
												mBase = m.M
												v47 = m.ExcPending
												if v47 != 0 {
													return int32(0)
												} else {
													*(*int32)(unsafe.Add(mBase, uint32(v9))) = v46
													F_errhint(m, int32(_a_F_AlterTypeNamespace_oid_1), v9)
													mBase = m.M
													v51 = m.ExcPending
													if v51 != 0 {
														return int32(0)
													} else {
														F_errfinish(m, int32(_a_F_AlterTypeNamespace_oid_2), int32(_a_F_AlterTypeNamespace_oid_3), int32(_a_F_AlterTypeNamespace_oid_4))
														mBase = m.M
														v56 = m.ExcPending
														if v56 != 0 {
															return int32(0)
														} else {
															base.Wasm_trap_unreachable()
															for {
															}
														}
													}
												}
											}
										}
									}
								}
							}
						}
					}
				}
			}
		}
	}
}
func F_AlterTypeOwner_oid(m *base.Module, l0 int32, l1 int32) {
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
	var v16 int32
	_ = v16
	var v17 int32
	_ = v17
	var v18 int32
	_ = v18
	var v19 int32
	_ = v19
	var v20 int32
	_ = v20
	var v21 int32
	_ = v21
	var v24 int32
	_ = v24
	var v28 int32
	_ = v28
	var v30 int32
	_ = v30
	var v33 int32
	_ = v33
	var v35 int32
	_ = v35
	var v37 int32
	_ = v37
	var v41 int32
	_ = v41
	var v43 int32
	_ = v43
	var v46 int32
	_ = v46
	var v53 int32
	_ = v53
	var v57 int32
	_ = v57
	var v62 int32
	_ = v62
	v7 = m.G0
	v9 = v7 - int32(16)
	m.G0 = v9
	v13 = F_table_open(m, int32(1247), int32(3))
	mBase = m.M
	v14 = m.ExcPending
	if v14 != 0 {
		return
	} else {
		v16 = F_SearchSysCache1(m, int32(82), l0)
		mBase = m.M
		v17 = m.ExcPending
		if v17 != 0 {
			return
		} else {
			if v16 != 0 {
				v18 = *(*int32)(unsafe.Add(mBase, uint32(v16)+16))
				v19 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v18)+22)))
				v20 = v18 + v19
				v21 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v20)+79)))
				if v21 == int32(99) {
					v24 = *(*int32)(unsafe.Add(mBase, uint32(v20)+84))
					F_ATExecChangeOwner(m, v24, l1, int32(1), int32(8))
					mBase = m.M
					v28 = m.ExcPending
					if v28 != 0 {
						return
					} else {
						F_changeDependencyOnOwner(m, int32(1247), l0, l1)
						mBase = m.M
						v33 = m.ExcPending
						if v33 != 0 {
							return
						} else {
							v35 = *(*int32)(unsafe.Add(mBase, _c_F_AlterTypeOwner_oid[0]))
							if v35 != 0 {
								v37 = int32(0)
								F_RunObjectPostAlterHook(m, int32(1247), l0, v37, v37, v37)
								mBase = m.M
								v41 = m.ExcPending
								if v41 != 0 {
									return
								} else {
									F_ReleaseCatCache(m, v16)
									mBase = m.M
									v43 = m.ExcPending
									if v43 != 0 {
										return
									} else {
										F_relation_close(m, v13, int32(3))
										mBase = m.M
										v46 = m.ExcPending
										if v46 != 0 {
											return
										} else {
											m.G0 = v9 + int32(16)
											return
										}
									}
								}
							} else {
								F_ReleaseCatCache(m, v16)
								mBase = m.M
								v43 = m.ExcPending
								if v43 != 0 {
									return
								} else {
									F_relation_close(m, v13, int32(3))
									mBase = m.M
									v46 = m.ExcPending
									if v46 != 0 {
										return
									} else {
										m.G0 = v9 + int32(16)
										return
									}
								}
							}
						}
					}
				} else {
					F_AlterTypeOwnerInternal(m, l0, l1)
					mBase = m.M
					v30 = m.ExcPending
					if v30 != 0 {
						return
					} else {
						F_changeDependencyOnOwner(m, int32(1247), l0, l1)
						mBase = m.M
						v33 = m.ExcPending
						if v33 != 0 {
							return
						} else {
							v35 = *(*int32)(unsafe.Add(mBase, _c_F_AlterTypeOwner_oid[0]))
							if v35 != 0 {
								v37 = int32(0)
								F_RunObjectPostAlterHook(m, int32(1247), l0, v37, v37, v37)
								mBase = m.M
								v41 = m.ExcPending
								if v41 != 0 {
									return
								} else {
									F_ReleaseCatCache(m, v16)
									mBase = m.M
									v43 = m.ExcPending
									if v43 != 0 {
										return
									} else {
										F_relation_close(m, v13, int32(3))
										mBase = m.M
										v46 = m.ExcPending
										if v46 != 0 {
											return
										} else {
											m.G0 = v9 + int32(16)
											return
										}
									}
								}
							} else {
								F_ReleaseCatCache(m, v16)
								mBase = m.M
								v43 = m.ExcPending
								if v43 != 0 {
									return
								} else {
									F_relation_close(m, v13, int32(3))
									mBase = m.M
									v46 = m.ExcPending
									if v46 != 0 {
										return
									} else {
										m.G0 = v9 + int32(16)
										return
									}
								}
							}
						}
					}
				}
			} else {
				F_errstart_cold(m, int32(21), int32(0))
				mBase = m.M
				v53 = m.ExcPending
				if v53 != 0 {
					return
				} else {
					*(*int32)(unsafe.Add(mBase, uint32(v9))) = l0
					F_errmsg_internal(m, int32(_a_F_AlterTypeOwner_oid_0), v9)
					mBase = m.M
					v57 = m.ExcPending
					if v57 != 0 {
						return
					} else {
						F_errfinish(m, int32(_a_F_AlterTypeOwner_oid_1), int32(3956), int32(_a_F_AlterTypeOwner_oid_2))
						mBase = m.M
						v62 = m.ExcPending
						if v62 != 0 {
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
	}
}
func F_AnonymousShmemDetach(m *base.Module, l0 int32, l1 int32) {
	mBase := m.M
	_ = mBase
	var v3 int32
	_ = v3
	var v5 int32
	_ = v5
	var v8 int32
	_ = v8
	v3 = m.G0
	v5 = v3 - int32(16)
	m.G0 = v5
	v8 = *(*int32)(unsafe.Add(mBase, _c_F_AnonymousShmemDetach[0]))
	if v8 != 0 {
		*(*int32)(unsafe.Add(mBase, _c_F_AnonymousShmemDetach[0])) = int32(0)
	} else {
	}
	m.G0 = v5 + int32(16)
	return
}
func F_AppendJumble32(m *base.Module, l0 int32, l1 int32) {
	mBase := m.M
	_ = mBase
	var v8 int32
	_ = v8
	var v11 int32
	_ = v11
	var v12 int32
	_ = v12
	var v13 int32
	_ = v13
	var v14 int32
	_ = v14
	var v23 int32
	_ = v23
	var v25 int32
	_ = v25
	var v27 int32
	_ = v27
	var v30 int32
	_ = v30
	var v37 int32
	_ = v37
	var v83 int32
	_ = v83
	var v84 int32
	_ = v84
	var v86 int32
	_ = v86
	var v87 int32
	_ = v87
	var v88 int32
	_ = v88
	var v90 int32
	_ = v90
	var v91 int32
	_ = v91
	var v92 int32
	_ = v92
	var v94 int32
	_ = v94
	var v95 int32
	_ = v95
	var v97 int32
	_ = v97
	var v99 int32
	_ = v99
	var v103 int32
	_ = v103
	var v104 int32
	_ = v104
	var v105 int32
	_ = v105
	var v106 int32
	_ = v106
	var v110 int32
	_ = v110
	var v114 int32
	_ = v114
	var v118 int32
	_ = v118
	var v119 int32
	_ = v119
	var v120 int32
	_ = v120
	var v121 int32
	_ = v121
	var v125 int32
	_ = v125
	var v126 int32
	_ = v126
	var v127 int32
	_ = v127
	var v129 int32
	_ = v129
	var v143 int32
	_ = v143
	var v144 int32
	_ = v144
	var v146 int32
	_ = v146
	var v147 int32
	_ = v147
	var v148 int32
	_ = v148
	var v150 int32
	_ = v150
	var v151 int32
	_ = v151
	var v152 int32
	_ = v152
	var v154 int32
	_ = v154
	var v155 int32
	_ = v155
	var v157 int32
	_ = v157
	var v159 int32
	_ = v159
	var v163 int32
	_ = v163
	var v164 int32
	_ = v164
	var v165 int32
	_ = v165
	var v166 int32
	_ = v166
	var v170 int32
	_ = v170
	var v174 int32
	_ = v174
	var v178 int32
	_ = v178
	var v179 int32
	_ = v179
	var v180 int32
	_ = v180
	var v181 int32
	_ = v181
	var v185 int32
	_ = v185
	var v186 int32
	_ = v186
	var v187 int32
	_ = v187
	var v189 int32
	_ = v189
	var v201 int32
	_ = v201
	var v205 int32
	_ = v205
	var v206 int32
	_ = v206
	var v210 int32
	_ = v210
	var v211 int32
	_ = v211
	var v215 int32
	_ = v215
	var v216 int32
	_ = v216
	var v218 int32
	_ = v218
	var v220 int32
	_ = v220
	var v224 int32
	_ = v224
	var v225 int32
	_ = v225
	var v229 int32
	_ = v229
	var v230 int32
	_ = v230
	var v232 int32
	_ = v232
	var v233 int32
	_ = v233
	var v235 int32
	_ = v235
	var v239 int32
	_ = v239
	var v240 int32
	_ = v240
	var v244 int32
	_ = v244
	var v245 int32
	_ = v245
	var v247 int32
	_ = v247
	var v251 int32
	_ = v251
	var v252 int32
	_ = v252
	var v256 int32
	_ = v256
	var v257 int32
	_ = v257
	var v261 int32
	_ = v261
	var v262 int32
	_ = v262
	var v266 int32
	_ = v266
	var v267 int32
	_ = v267
	var v268 int32
	_ = v268
	var v272 int32
	_ = v272
	var v273 int32
	_ = v273
	var v274 int32
	_ = v274
	var v278 int32
	_ = v278
	var v279 int32
	_ = v279
	var v280 int32
	_ = v280
	var v282 int32
	_ = v282
	var v283 int32
	_ = v283
	var v284 int32
	_ = v284
	var v288 int32
	_ = v288
	var v289 int32
	_ = v289
	var v290 int32
	_ = v290
	var v291 int32
	_ = v291
	var v295 int32
	_ = v295
	var v296 int32
	_ = v296
	var v297 int32
	_ = v297
	var v298 int32
	_ = v298
	var v302 int32
	_ = v302
	var v303 int32
	_ = v303
	var v304 int32
	_ = v304
	var v305 int32
	_ = v305
	var v310 int32
	_ = v310
	var v311 int32
	_ = v311
	var v312 int32
	_ = v312
	var v315 int32
	_ = v315
	var v317 int32
	_ = v317
	var v321 int32
	_ = v321
	var v325 int32
	_ = v325
	var v329 int32
	_ = v329
	var v333 int32
	_ = v333
	var v337 int32
	_ = v337
	var v347 int32
	_ = v347
	var v348 int32
	_ = v348
	var v349 int64
	_ = v349
	var v352 int32
	_ = v352
	var v354 int32
	_ = v354
	var v356 int32
	_ = v356
	var v360 int32
	_ = v360
	var v361 int32
	_ = v361
	var v364 int32
	_ = v364
	var v369 int32
	_ = v369
	var v379 int32
	_ = v379
	var v384 int32
	_ = v384
	var v385 int32
	_ = v385
	var v391 int32
	_ = v391
	var v392 int32
	_ = v392
	var v394 int32
	_ = v394
	var v399 int32
	_ = v399
	var v406 int32
	_ = v406
	var v452 int32
	_ = v452
	var v453 int32
	_ = v453
	var v455 int32
	_ = v455
	var v456 int32
	_ = v456
	var v457 int32
	_ = v457
	var v459 int32
	_ = v459
	var v460 int32
	_ = v460
	var v461 int32
	_ = v461
	var v463 int32
	_ = v463
	var v464 int32
	_ = v464
	var v466 int32
	_ = v466
	var v468 int32
	_ = v468
	var v472 int32
	_ = v472
	var v473 int32
	_ = v473
	var v474 int32
	_ = v474
	var v475 int32
	_ = v475
	var v479 int32
	_ = v479
	var v483 int32
	_ = v483
	var v487 int32
	_ = v487
	var v488 int32
	_ = v488
	var v489 int32
	_ = v489
	var v490 int32
	_ = v490
	var v494 int32
	_ = v494
	var v495 int32
	_ = v495
	var v496 int32
	_ = v496
	var v498 int32
	_ = v498
	var v512 int32
	_ = v512
	var v513 int32
	_ = v513
	var v515 int32
	_ = v515
	var v516 int32
	_ = v516
	var v517 int32
	_ = v517
	var v519 int32
	_ = v519
	var v520 int32
	_ = v520
	var v521 int32
	_ = v521
	var v523 int32
	_ = v523
	var v524 int32
	_ = v524
	var v526 int32
	_ = v526
	var v528 int32
	_ = v528
	var v532 int32
	_ = v532
	var v533 int32
	_ = v533
	var v534 int32
	_ = v534
	var v535 int32
	_ = v535
	var v539 int32
	_ = v539
	var v543 int32
	_ = v543
	var v547 int32
	_ = v547
	var v548 int32
	_ = v548
	var v549 int32
	_ = v549
	var v550 int32
	_ = v550
	var v554 int32
	_ = v554
	var v555 int32
	_ = v555
	var v556 int32
	_ = v556
	var v558 int32
	_ = v558
	var v570 int32
	_ = v570
	var v574 int32
	_ = v574
	var v575 int32
	_ = v575
	var v579 int32
	_ = v579
	var v580 int32
	_ = v580
	var v584 int32
	_ = v584
	var v585 int32
	_ = v585
	var v587 int32
	_ = v587
	var v589 int32
	_ = v589
	var v593 int32
	_ = v593
	var v594 int32
	_ = v594
	var v598 int32
	_ = v598
	var v599 int32
	_ = v599
	var v601 int32
	_ = v601
	var v602 int32
	_ = v602
	var v604 int32
	_ = v604
	var v608 int32
	_ = v608
	var v609 int32
	_ = v609
	var v613 int32
	_ = v613
	var v614 int32
	_ = v614
	var v616 int32
	_ = v616
	var v620 int32
	_ = v620
	var v621 int32
	_ = v621
	var v625 int32
	_ = v625
	var v626 int32
	_ = v626
	var v630 int32
	_ = v630
	var v631 int32
	_ = v631
	var v635 int32
	_ = v635
	var v636 int32
	_ = v636
	var v637 int32
	_ = v637
	var v641 int32
	_ = v641
	var v642 int32
	_ = v642
	var v643 int32
	_ = v643
	var v647 int32
	_ = v647
	var v648 int32
	_ = v648
	var v649 int32
	_ = v649
	var v651 int32
	_ = v651
	var v652 int32
	_ = v652
	var v653 int32
	_ = v653
	var v657 int32
	_ = v657
	var v658 int32
	_ = v658
	var v659 int32
	_ = v659
	var v660 int32
	_ = v660
	var v664 int32
	_ = v664
	var v665 int32
	_ = v665
	var v666 int32
	_ = v666
	var v667 int32
	_ = v667
	var v671 int32
	_ = v671
	var v672 int32
	_ = v672
	var v673 int32
	_ = v673
	var v674 int32
	_ = v674
	var v679 int32
	_ = v679
	var v680 int32
	_ = v680
	var v681 int32
	_ = v681
	var v684 int32
	_ = v684
	var v686 int32
	_ = v686
	var v690 int32
	_ = v690
	var v694 int32
	_ = v694
	var v698 int32
	_ = v698
	var v702 int32
	_ = v702
	var v706 int32
	_ = v706
	var v716 int32
	_ = v716
	var v717 int32
	_ = v717
	var v718 int64
	_ = v718
	var v721 int32
	_ = v721
	var v723 int32
	_ = v723
	var v725 int32
	_ = v725
	var v729 int32
	_ = v729
	var v730 int32
	_ = v730
	var v733 int32
	_ = v733
	var v735 int32
	_ = v735
	v8 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
	if v8 == int32(0) {
		goto L2
	} else {
		goto L3
	}
L1:
	;
	v384 = int32(4)
	v385 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	if base.Ui32(v379-int32(1021)) < base.Ui32(v384) {
		goto L66
	} else {
		goto L67
	}
L2:
	;
	v11 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v379 = v11
	goto L1
L3:
	;
	goto L4
L4:
	;
	v12 = int32(4)
	v13 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v14 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	if base.Ui32(v14-int32(1021)) < base.Ui32(v12) {
		goto L6
	} else {
		goto L7
	}
L5:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+28)) = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v369
	v379 = v369
	goto L1
L6:
	;
	v23 = v14
	v25 = v12
	v27 = l0 + int32(28)
	goto L9
L7:
	;
	goto L8
L8:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v14+v13))) = v8
	v364 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v369 = v364 + int32(4)
	goto L5
L9:
	;
	if base.Ui32(int32(1024)) <= base.Ui32(v23) {
		goto L11
	} else {
		goto L12
	}
L10:
	;
	v369 = v360
	goto L5
L11:
	;
	v30 = int32(1024)
	v37 = int32(-1636607408)
	goto L16
L12:
	;
	v352 = v23
	goto L13
L13:
	;
	v354 = int32(1024) - v352
	if base.Ui32(v25) < base.Ui32(v354) {
		goto L59
	} else {
		goto L60
	}
L14:
	;
	v347 = F_Int64GetDatum(m, base.I64_extend_i32_u(v337)<<(uint(int64(32))%64)|base.I64_extend_i32_u(v337^v329-base.I32_rotl(v337, int32(24))))
	mBase = m.M
	v348 = m.ExcPending
	if v348 != 0 {
		goto L57
	} else {
		goto L58
	}
L15:
	;
	if v13&int32(3) != 0 {
		goto L31
	} else {
		goto L32
	}
L16:
	;
	goto L15
L19:
	;
	v315 = int32(14)
	v317 = v311 ^ v312 - base.I32_rotl(v311, v315)
	v321 = v317 ^ v310 - base.I32_rotl(v317, int32(11))
	v325 = v321 ^ v311 - base.I32_rotl(v321, int32(25))
	v329 = v325 ^ v317 - base.I32_rotl(v325, int32(16))
	v333 = v329 ^ v321 - base.I32_rotl(v329, int32(4))
	v337 = v333 ^ v325 - base.I32_rotl(v333, v315)
	goto L14
L20:
	;
	v305 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v127))))
	v310 = v302 + v305
	v311 = v303
	v312 = v304
	goto L19
L21:
	;
	v298 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v127)+1)))
	v302 = v298<<(uint(int32(8))%32) + v295
	v303 = v296
	v304 = v297
	goto L20
L22:
	;
	v291 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v127)+2)))
	v295 = v291<<(uint(int32(16))%32) + v288
	v296 = v289
	v297 = v290
	goto L21
L23:
	;
	v284 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v127)+3)))
	v288 = v284<<(uint(int32(24))%32) + v120
	v289 = v282
	v290 = v283
	goto L22
L24:
	;
	v280 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v127)+4)))
	v282 = v278 + v280
	v283 = v279
	goto L23
L25:
	;
	v274 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v127)+5)))
	v278 = v274<<(uint(int32(8))%32) + v272
	v279 = v273
	goto L24
L26:
	;
	v268 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v127)+6)))
	v272 = v268<<(uint(int32(16))%32) + v266
	v273 = v267
	goto L25
L27:
	;
	v262 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v127)+7)))
	v266 = v262<<(uint(int32(24))%32) + v121
	v267 = v261
	goto L26
L28:
	;
	v257 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v127)+8)))
	v261 = v257<<(uint(int32(8))%32) + v256
	goto L27
L29:
	;
	v252 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v127)+9)))
	v256 = v252<<(uint(int32(16))%32) + v251
	goto L28
L30:
	;
	v247 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v127)+10)))
	v251 = v247<<(uint(int32(24))%32) + v125
	goto L29
L31:
	;
	goto L34
L32:
	;
	goto L33
L33:
	;
	goto L40
L34:
	;
	v83 = v13
	v84 = v30
	v86 = v37
	v87 = v37
	v88 = v37
	goto L37
L36:
	;
	switch v129 - int32(1) {
	case 0:
		v302 = v120
		v303 = v121
		v304 = v125
		goto L20
	case 1:
		v295 = v120
		v296 = v121
		v297 = v125
		goto L21
	case 2:
		v288 = v120
		v289 = v121
		v290 = v125
		goto L22
	case 3:
		v282 = v121
		v283 = v125
		goto L23
	case 4:
		v278 = v121
		v279 = v125
		goto L24
	case 5:
		v272 = v121
		v273 = v125
		goto L25
	case 6:
		v266 = v121
		v267 = v125
		goto L26
	case 7:
		v261 = v125
		goto L27
	case 8:
		v256 = v125
		goto L28
	case 9:
		v251 = v125
		goto L29
	case 10:
		goto L30
	default:
		v310 = v120
		v311 = v121
		v312 = v125
		goto L19
	}
L37:
	;
	v90 = *(*int32)(unsafe.Add(mBase, uint32(v83)+4))
	v91 = v90 + v87
	v92 = *(*int32)(unsafe.Add(mBase, uint32(v83)))
	v94 = *(*int32)(unsafe.Add(mBase, uint32(v83)+8))
	v95 = v94 + v88
	v97 = int32(4)
	v99 = v92 + v86 - v95 ^ base.I32_rotl(v95, v97)
	v103 = v91 - v99 ^ base.I32_rotl(v99, int32(6))
	v104 = v95 + v91
	v105 = v99 + v104
	v106 = v103 + v105
	v110 = v104 - v103 ^ base.I32_rotl(v103, int32(8))
	v114 = v105 - v110 ^ base.I32_rotl(v110, int32(16))
	v118 = v106 - v114 ^ base.I32_rotl(v114, int32(19))
	v119 = v110 + v106
	v120 = v114 + v119
	v121 = v118 + v120
	v125 = v119 - v118 ^ base.I32_rotl(v118, v97)
	v126 = int32(12)
	v127 = v83 + v126
	v129 = v84 - v126
	if base.Ui32(int32(11)) < base.Ui32(v129) {
		v83 = v127
		v84 = v129
		v86 = v120
		v87 = v121
		v88 = v125
		goto L37
	} else {
		goto L39
	}
L38:
	;
	goto L36
L39:
	;
	goto L38
L40:
	;
	v143 = v13
	v144 = v30
	v146 = v37
	v147 = v37
	v148 = v37
	goto L43
L42:
	;
	switch v189 - int32(1) {
	case 0:
		v244 = v180
		goto L46
	case 1:
		v239 = v180
		goto L47
	case 2:
		goto L48
	case 3:
		v232 = v181
		goto L49
	case 4:
		v229 = v181
		goto L50
	case 5:
		v224 = v181
		goto L51
	case 6:
		goto L52
	case 7:
		v215 = v185
		goto L53
	case 8:
		v210 = v185
		goto L54
	case 9:
		v205 = v185
		goto L55
	case 10:
		goto L56
	default:
		v310 = v180
		v311 = v181
		v312 = v185
		goto L19
	}
L43:
	;
	v150 = *(*int32)(unsafe.Add(mBase, uint32(v143)+4))
	v151 = v150 + v147
	v152 = *(*int32)(unsafe.Add(mBase, uint32(v143)))
	v154 = *(*int32)(unsafe.Add(mBase, uint32(v143)+8))
	v155 = v154 + v148
	v157 = int32(4)
	v159 = v152 + v146 - v155 ^ base.I32_rotl(v155, v157)
	v163 = v151 - v159 ^ base.I32_rotl(v159, int32(6))
	v164 = v155 + v151
	v165 = v159 + v164
	v166 = v163 + v165
	v170 = v164 - v163 ^ base.I32_rotl(v163, int32(8))
	v174 = v165 - v170 ^ base.I32_rotl(v170, int32(16))
	v178 = v166 - v174 ^ base.I32_rotl(v174, int32(19))
	v179 = v170 + v166
	v180 = v174 + v179
	v181 = v178 + v180
	v185 = v179 - v178 ^ base.I32_rotl(v178, v157)
	v186 = int32(12)
	v187 = v143 + v186
	v189 = v144 - v186
	if base.Ui32(int32(11)) < base.Ui32(v189) {
		v143 = v187
		v144 = v189
		v146 = v180
		v147 = v181
		v148 = v185
		goto L43
	} else {
		goto L45
	}
L44:
	;
	goto L42
L45:
	;
	goto L44
L46:
	;
	v245 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v187))))
	v310 = v244 + v245
	v311 = v181
	v312 = v185
	goto L19
L47:
	;
	v240 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v187)+1)))
	v244 = v240<<(uint(int32(8))%32) + v239
	goto L46
L48:
	;
	v235 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v187)+2)))
	v239 = v235<<(uint(int32(16))%32) + v180
	goto L47
L49:
	;
	v233 = *(*int32)(unsafe.Add(mBase, uint32(v187)))
	v310 = v233 + v180
	v311 = v232
	v312 = v185
	goto L19
L50:
	;
	v230 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v187)+4)))
	v232 = v229 + v230
	goto L49
L51:
	;
	v225 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v187)+5)))
	v229 = v225<<(uint(int32(8))%32) + v224
	goto L50
L52:
	;
	v220 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v187)+6)))
	v224 = v220<<(uint(int32(16))%32) + v181
	goto L51
L53:
	;
	v216 = *(*int32)(unsafe.Add(mBase, uint32(v187)))
	v218 = *(*int32)(unsafe.Add(mBase, uint32(v187)+4))
	v310 = v216 + v180
	v311 = v218 + v181
	v312 = v215
	goto L19
L54:
	;
	v211 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v187)+8)))
	v215 = v211<<(uint(int32(8))%32) + v210
	goto L53
L55:
	;
	v206 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v187)+9)))
	v210 = v206<<(uint(int32(16))%32) + v205
	goto L54
L56:
	;
	v201 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v187)+10)))
	v205 = v201<<(uint(int32(24))%32) + v185
	goto L55
L57:
	;
	return
L58:
	;
	v349 = *(*int64)(unsafe.Add(mBase, uint32(v347)))
	*(*int64)(unsafe.Add(mBase, uint32(v13))) = v349
	v352 = int32(8)
	goto L13
L59:
	;
	v356 = v25
	goto L61
L60:
	;
	v356 = v354
	goto L61
L61:
	;
	if v356 != 0 {
		goto L62
	} else {
		goto L63
	}
L62:
	;
	base.MemoryCopy(m, v352+v13, v27, v356)
	goto L64
L63:
	;
	goto L64
L64:
	;
	v360 = v352 + v356
	v361 = v25 - v356
	if v361 != 0 {
		v23 = v360
		v25 = v361
		v27 = v356 + v27
		goto L9
	} else {
		goto L65
	}
L65:
	;
	goto L10
L66:
	;
	v391 = l1
	v392 = v379
	v394 = v384
	goto L69
L67:
	;
	goto L68
L68:
	;
	v733 = *(*int32)(unsafe.Add(mBase, uint32(l1)))
	*(*int32)(unsafe.Add(mBase, uint32(v379+v385))) = v733
	v735 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v735 + int32(4)
	return
L69:
	;
	if base.Ui32(int32(1024)) <= base.Ui32(v392) {
		goto L71
	} else {
		goto L72
	}
L70:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v729
	return
L71:
	;
	v399 = int32(1024)
	v406 = int32(-1636607408)
	goto L76
L72:
	;
	v721 = v392
	goto L73
L73:
	;
	v723 = int32(1024) - v721
	if base.Ui32(v394) < base.Ui32(v723) {
		goto L118
	} else {
		goto L119
	}
L74:
	;
	v716 = F_Int64GetDatum(m, base.I64_extend_i32_u(v706)<<(uint(int64(32))%64)|base.I64_extend_i32_u(v706^v698-base.I32_rotl(v706, int32(24))))
	mBase = m.M
	v717 = m.ExcPending
	if v717 != 0 {
		goto L57
	} else {
		goto L117
	}
L75:
	;
	if v385&int32(3) != 0 {
		goto L91
	} else {
		goto L92
	}
L76:
	;
	goto L75
L79:
	;
	v684 = int32(14)
	v686 = v680 ^ v681 - base.I32_rotl(v680, v684)
	v690 = v686 ^ v679 - base.I32_rotl(v686, int32(11))
	v694 = v690 ^ v680 - base.I32_rotl(v690, int32(25))
	v698 = v694 ^ v686 - base.I32_rotl(v694, int32(16))
	v702 = v698 ^ v690 - base.I32_rotl(v698, int32(4))
	v706 = v702 ^ v694 - base.I32_rotl(v702, v684)
	goto L74
L80:
	;
	v674 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v496))))
	v679 = v671 + v674
	v680 = v672
	v681 = v673
	goto L79
L81:
	;
	v667 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v496)+1)))
	v671 = v667<<(uint(int32(8))%32) + v664
	v672 = v665
	v673 = v666
	goto L80
L82:
	;
	v660 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v496)+2)))
	v664 = v660<<(uint(int32(16))%32) + v657
	v665 = v658
	v666 = v659
	goto L81
L83:
	;
	v653 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v496)+3)))
	v657 = v653<<(uint(int32(24))%32) + v489
	v658 = v651
	v659 = v652
	goto L82
L84:
	;
	v649 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v496)+4)))
	v651 = v647 + v649
	v652 = v648
	goto L83
L85:
	;
	v643 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v496)+5)))
	v647 = v643<<(uint(int32(8))%32) + v641
	v648 = v642
	goto L84
L86:
	;
	v637 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v496)+6)))
	v641 = v637<<(uint(int32(16))%32) + v635
	v642 = v636
	goto L85
L87:
	;
	v631 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v496)+7)))
	v635 = v631<<(uint(int32(24))%32) + v490
	v636 = v630
	goto L86
L88:
	;
	v626 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v496)+8)))
	v630 = v626<<(uint(int32(8))%32) + v625
	goto L87
L89:
	;
	v621 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v496)+9)))
	v625 = v621<<(uint(int32(16))%32) + v620
	goto L88
L90:
	;
	v616 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v496)+10)))
	v620 = v616<<(uint(int32(24))%32) + v494
	goto L89
L91:
	;
	goto L94
L92:
	;
	goto L93
L93:
	;
	goto L100
L94:
	;
	v452 = v385
	v453 = v399
	v455 = v406
	v456 = v406
	v457 = v406
	goto L97
L96:
	;
	switch v498 - int32(1) {
	case 0:
		v671 = v489
		v672 = v490
		v673 = v494
		goto L80
	case 1:
		v664 = v489
		v665 = v490
		v666 = v494
		goto L81
	case 2:
		v657 = v489
		v658 = v490
		v659 = v494
		goto L82
	case 3:
		v651 = v490
		v652 = v494
		goto L83
	case 4:
		v647 = v490
		v648 = v494
		goto L84
	case 5:
		v641 = v490
		v642 = v494
		goto L85
	case 6:
		v635 = v490
		v636 = v494
		goto L86
	case 7:
		v630 = v494
		goto L87
	case 8:
		v625 = v494
		goto L88
	case 9:
		v620 = v494
		goto L89
	case 10:
		goto L90
	default:
		v679 = v489
		v680 = v490
		v681 = v494
		goto L79
	}
L97:
	;
	v459 = *(*int32)(unsafe.Add(mBase, uint32(v452)+4))
	v460 = v459 + v456
	v461 = *(*int32)(unsafe.Add(mBase, uint32(v452)))
	v463 = *(*int32)(unsafe.Add(mBase, uint32(v452)+8))
	v464 = v463 + v457
	v466 = int32(4)
	v468 = v461 + v455 - v464 ^ base.I32_rotl(v464, v466)
	v472 = v460 - v468 ^ base.I32_rotl(v468, int32(6))
	v473 = v464 + v460
	v474 = v468 + v473
	v475 = v472 + v474
	v479 = v473 - v472 ^ base.I32_rotl(v472, int32(8))
	v483 = v474 - v479 ^ base.I32_rotl(v479, int32(16))
	v487 = v475 - v483 ^ base.I32_rotl(v483, int32(19))
	v488 = v479 + v475
	v489 = v483 + v488
	v490 = v487 + v489
	v494 = v488 - v487 ^ base.I32_rotl(v487, v466)
	v495 = int32(12)
	v496 = v452 + v495
	v498 = v453 - v495
	if base.Ui32(int32(11)) < base.Ui32(v498) {
		v452 = v496
		v453 = v498
		v455 = v489
		v456 = v490
		v457 = v494
		goto L97
	} else {
		goto L99
	}
L98:
	;
	goto L96
L99:
	;
	goto L98
L100:
	;
	v512 = v385
	v513 = v399
	v515 = v406
	v516 = v406
	v517 = v406
	goto L103
L102:
	;
	switch v558 - int32(1) {
	case 0:
		v613 = v549
		goto L106
	case 1:
		v608 = v549
		goto L107
	case 2:
		goto L108
	case 3:
		v601 = v550
		goto L109
	case 4:
		v598 = v550
		goto L110
	case 5:
		v593 = v550
		goto L111
	case 6:
		goto L112
	case 7:
		v584 = v554
		goto L113
	case 8:
		v579 = v554
		goto L114
	case 9:
		v574 = v554
		goto L115
	case 10:
		goto L116
	default:
		v679 = v549
		v680 = v550
		v681 = v554
		goto L79
	}
L103:
	;
	v519 = *(*int32)(unsafe.Add(mBase, uint32(v512)+4))
	v520 = v519 + v516
	v521 = *(*int32)(unsafe.Add(mBase, uint32(v512)))
	v523 = *(*int32)(unsafe.Add(mBase, uint32(v512)+8))
	v524 = v523 + v517
	v526 = int32(4)
	v528 = v521 + v515 - v524 ^ base.I32_rotl(v524, v526)
	v532 = v520 - v528 ^ base.I32_rotl(v528, int32(6))
	v533 = v524 + v520
	v534 = v528 + v533
	v535 = v532 + v534
	v539 = v533 - v532 ^ base.I32_rotl(v532, int32(8))
	v543 = v534 - v539 ^ base.I32_rotl(v539, int32(16))
	v547 = v535 - v543 ^ base.I32_rotl(v543, int32(19))
	v548 = v539 + v535
	v549 = v543 + v548
	v550 = v547 + v549
	v554 = v548 - v547 ^ base.I32_rotl(v547, v526)
	v555 = int32(12)
	v556 = v512 + v555
	v558 = v513 - v555
	if base.Ui32(int32(11)) < base.Ui32(v558) {
		v512 = v556
		v513 = v558
		v515 = v549
		v516 = v550
		v517 = v554
		goto L103
	} else {
		goto L105
	}
L104:
	;
	goto L102
L105:
	;
	goto L104
L106:
	;
	v614 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v556))))
	v679 = v613 + v614
	v680 = v550
	v681 = v554
	goto L79
L107:
	;
	v609 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v556)+1)))
	v613 = v609<<(uint(int32(8))%32) + v608
	goto L106
L108:
	;
	v604 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v556)+2)))
	v608 = v604<<(uint(int32(16))%32) + v549
	goto L107
L109:
	;
	v602 = *(*int32)(unsafe.Add(mBase, uint32(v556)))
	v679 = v602 + v549
	v680 = v601
	v681 = v554
	goto L79
L110:
	;
	v599 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v556)+4)))
	v601 = v598 + v599
	goto L109
L111:
	;
	v594 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v556)+5)))
	v598 = v594<<(uint(int32(8))%32) + v593
	goto L110
L112:
	;
	v589 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v556)+6)))
	v593 = v589<<(uint(int32(16))%32) + v550
	goto L111
L113:
	;
	v585 = *(*int32)(unsafe.Add(mBase, uint32(v556)))
	v587 = *(*int32)(unsafe.Add(mBase, uint32(v556)+4))
	v679 = v585 + v549
	v680 = v587 + v550
	v681 = v584
	goto L79
L114:
	;
	v580 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v556)+8)))
	v584 = v580<<(uint(int32(8))%32) + v579
	goto L113
L115:
	;
	v575 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v556)+9)))
	v579 = v575<<(uint(int32(16))%32) + v574
	goto L114
L116:
	;
	v570 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v556)+10)))
	v574 = v570<<(uint(int32(24))%32) + v554
	goto L115
L117:
	;
	v718 = *(*int64)(unsafe.Add(mBase, uint32(v716)))
	*(*int64)(unsafe.Add(mBase, uint32(v385))) = v718
	v721 = int32(8)
	goto L73
L118:
	;
	v725 = v394
	goto L120
L119:
	;
	v725 = v723
	goto L120
L120:
	;
	if v725 != 0 {
		goto L121
	} else {
		goto L122
	}
L121:
	;
	base.MemoryCopy(m, v721+v385, v391, v725)
	goto L123
L122:
	;
	goto L123
L123:
	;
	v729 = v721 + v725
	v730 = v394 - v725
	if v730 != 0 {
		v391 = v391 + v725
		v392 = v729
		v394 = v730
		goto L69
	} else {
		goto L124
	}
L124:
	;
	goto L70
}
func F_AutoVacWorkerMain(m *base.Module, l0 int32, l1 int32) {
	mBase := m.M
	_ = mBase
	var v7 int32
	_ = v7
	var v9 int32
	_ = v9
	var v12 int32
	_ = v12
	var v14 int32
	_ = v14
	var v21 int32
	_ = v21
	var v23 int32
	_ = v23
	var v34 int32
	_ = v34
	var v35 int32
	_ = v35
	var v37 int32
	_ = v37
	var v39 int32
	_ = v39
	var v41 int32
	_ = v41
	var v51 int32
	_ = v51
	var v65 int32
	_ = v65
	var v70 int32
	_ = v70
	var v72 int32
	_ = v72
	var v74 int32
	_ = v74
	var v84 int32
	_ = v84
	var v98 int32
	_ = v98
	var v103 int32
	_ = v103
	var v105 int32
	_ = v105
	var v107 int32
	_ = v107
	var v117 int32
	_ = v117
	var v131 int32
	_ = v131
	var v135 int32
	_ = v135
	var v145 int32
	_ = v145
	var v147 int32
	_ = v147
	var v148 int32
	_ = v148
	var v149 int32
	_ = v149
	var v152 int64
	_ = v152
	var v163 int32
	_ = v163
	var v165 int32
	_ = v165
	var v180 int32
	_ = v180
	var v182 int32
	_ = v182
	var v199 int32
	_ = v199
	var v201 int32
	_ = v201
	var v202 int32
	_ = v202
	var v205 int64
	_ = v205
	var v218 int32
	_ = v218
	var v224 int32
	_ = v224
	var v226 int32
	_ = v226
	var v228 int32
	_ = v228
	var v238 int32
	_ = v238
	var v252 int32
	_ = v252
	var v257 int32
	_ = v257
	var v259 int32
	_ = v259
	var v261 int32
	_ = v261
	var v271 int32
	_ = v271
	var v285 int32
	_ = v285
	var v290 int32
	_ = v290
	var v292 int32
	_ = v292
	var v294 int32
	_ = v294
	var v304 int32
	_ = v304
	var v318 int32
	_ = v318
	var v323 int32
	_ = v323
	var v325 int32
	_ = v325
	var v327 int32
	_ = v327
	var v337 int32
	_ = v337
	var v351 int32
	_ = v351
	var v356 int32
	_ = v356
	var v358 int32
	_ = v358
	var v360 int32
	_ = v360
	var v370 int32
	_ = v370
	var v384 int32
	_ = v384
	var v389 int32
	_ = v389
	var v391 int32
	_ = v391
	var v393 int32
	_ = v393
	var v401 int32
	_ = v401
	var v402 int32
	_ = v402
	var v404 int32
	_ = v404
	var v412 int32
	_ = v412
	var v415 int32
	_ = v415
	var v423 int32
	_ = v423
	var v429 int32
	_ = v429
	var v435 int32
	_ = v435
	var v441 int32
	_ = v441
	var v447 int32
	_ = v447
	var v453 int32
	_ = v453
	var v459 int32
	_ = v459
	var v465 int32
	_ = v465
	var v467 int32
	_ = v467
	var v475 int32
	_ = v475
	var v481 int32
	_ = v481
	var v483 int32
	_ = v483
	var v487 int32
	_ = v487
	var v488 int32
	_ = v488
	var v490 int32
	_ = v490
	var v491 int32
	_ = v491
	var v495 int32
	_ = v495
	var v498 int32
	_ = v498
	var v499 int32
	_ = v499
	var v500 int32
	_ = v500
	var v504 int32
	_ = v504
	var v512 int32
	_ = v512
	var v516 int32
	_ = v516
	var v520 int32
	_ = v520
	var v522 int32
	_ = v522
	var v523 int32
	_ = v523
	var v525 int32
	_ = v525
	var v531 int32
	_ = v531
	var v532 int32
	_ = v532
	var v533 int32
	_ = v533
	var v537 int32
	_ = v537
	var v538 int32
	_ = v538
	var v539 int32
	_ = v539
	var v542 int64
	_ = v542
	var v543 int64
	_ = v543
	var v554 int32
	_ = v554
	var v555 int32
	_ = v555
	var v560 int32
	_ = v560
	var v562 int32
	_ = v562
	var v566 int32
	_ = v566
	var v569 int32
	_ = v569
	var v570 int32
	_ = v570
	var v574 int32
	_ = v574
	var v579 int32
	_ = v579
	var v581 int32
	_ = v581
	var v586 int64
	_ = v586
	var v587 int32
	_ = v587
	var v590 int32
	_ = v590
	var v591 int32
	_ = v591
	var v594 int32
	_ = v594
	var v597 int32
	_ = v597
	var v598 int32
	_ = v598
	var v602 int32
	_ = v602
	var v607 int32
	_ = v607
	var v609 int32
	_ = v609
	var v613 int32
	_ = v613
	var v619 int32
	_ = v619
	var v625 int32
	_ = v625
	var v626 int64
	_ = v626
	var v630 int32
	_ = v630
	var v632 int32
	_ = v632
	var v633 int32
	_ = v633
	var v636 int32
	_ = v636
	var v638 int32
	_ = v638
	var v640 int32
	_ = v640
	var v644 int32
	_ = v644
	v7 = m.G0
	v9 = v7 - int32(240)
	m.G0 = v9
	v12 = int32(-1)
	v14 = int32(0)
	goto L1
L1:
	;
	goto L4
L2:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L3:
	;
	goto L2
L4:
	;
	if v12 != int32(1) {
		goto L7
	} else {
		goto L8
	}
L5:
	;
	goto L3
L6:
	;
	v625 = int32(m.ExcTag)
	v626 = int64(m.ExcVals[0])
	m.ExcPending = 0
	if v625 == int32(0) {
		goto L134
	} else {
		goto L135
	}
L7:
	;
	v21 = *(*int32)(unsafe.Add(mBase, _c_F_AutoVacWorkerMain[0]))
	if v21 != 0 {
		goto L10
	} else {
		goto L11
	}
L8:
	;
	v401 = v14
	goto L9
L9:
	;
	if v401 != 0 {
		goto L78
	} else {
		goto L79
	}
L10:
	;
	F_MemoryContextDelete(m, v21)
	mBase = m.M
	v23 = m.ExcPending
	if v23 != 0 {
		goto L6
	} else {
		goto L13
	}
L11:
	;
	goto L12
L12:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_AutoVacWorkerMain[1])) = int32(4)
	goto L15
L13:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_AutoVacWorkerMain[0])) = int32(0)
	goto L12
L14:
	;
	v37 = int32(914)
	v39 = m.G0
	v41 = v39 - int32(32)
	m.G0 = v41
	switch int32(916) {
	case 0, 2:
		v51 = v37
		goto L19
	default:
		goto L20
	}
L15:
	;
	v34 = *(*int32)(unsafe.Add(mBase, _c_F_AutoVacWorkerMain[1]))
	v35 = F_GetBackendTypeDesc(m, v34)
	mBase = m.M
	goto L17
L17:
	;
	goto L14
L18:
	;
	v70 = int32(915)
	v72 = m.G0
	v74 = v72 - int32(32)
	m.G0 = v74
	switch int32(917) {
	case 0, 2:
		v84 = v70
		goto L25
	default:
		goto L26
	}
L19:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v41)+12)) = v51
	F_sigemptyset(m, v41+int32(16))
	mBase = m.M
	goto L22
L20:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_AutoVacWorkerMain[2])) = v37
	v51 = int32(_a_F_AutoVacWorkerMain_0)
	goto L19
L22:
	;
	goto L23
L23:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v41)+24)) = int32(268435456)
	v65 = F___sigaction(m, int32(1), v41+int32(12), int32(0))
	mBase = m.M
	m.G0 = v41 + int32(32)
	goto L18
L24:
	;
	v103 = int32(295)
	v105 = m.G0
	v107 = v105 - int32(32)
	m.G0 = v107
	switch int32(297) {
	case 0, 2:
		v117 = v103
		goto L31
	default:
		goto L32
	}
L25:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v74)+12)) = v84
	F_sigemptyset(m, v74+int32(16))
	mBase = m.M
	goto L28
L26:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_AutoVacWorkerMain[3])) = v70
	v84 = int32(_a_F_AutoVacWorkerMain_0)
	goto L25
L28:
	;
	goto L29
L29:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v74)+24)) = int32(268435456)
	v98 = F___sigaction(m, int32(2), v74+int32(12), int32(0))
	mBase = m.M
	m.G0 = v74 + int32(32)
	goto L24
L30:
	;
	v135 = int32(0)
	*(*int32)(unsafe.Add(mBase, _c_F_AutoVacWorkerMain[4])) = v135
	*(*int32)(unsafe.Add(mBase, _c_F_AutoVacWorkerMain[5])) = v135
	v145 = v135
	goto L37
L31:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v107)+12)) = v117
	F_sigemptyset(m, v107+int32(16))
	mBase = m.M
	goto L34
L32:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_AutoVacWorkerMain[6])) = v103
	v117 = int32(_a_F_AutoVacWorkerMain_0)
	goto L31
L34:
	;
	goto L35
L35:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v107)+24)) = int32(268435456)
	v131 = F___sigaction(m, int32(15), v107+int32(12), int32(0))
	mBase = m.M
	m.G0 = v107 + int32(32)
	goto L30
L36:
	;
	v224 = int32(-2)
	v226 = m.G0
	v228 = v226 - int32(32)
	m.G0 = v228
	switch int32(0) {
	case 0, 2:
		v238 = v224
		goto L43
	default:
		goto L44
	}
L37:
	;
	v147 = int32(40)
	v148 = v145 * v147
	v149 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v148)+uint32(_c_F_AutoVacWorkerMain[7]))) = uint8(v149)
	*(*int32)(unsafe.Add(mBase, uint32(v148)+uint32(_c_F_AutoVacWorkerMain[8]))) = v145
	v152 = int64(0)
	*(*int64)(unsafe.Add(mBase, uint32(v148)+uint32(_c_F_AutoVacWorkerMain[9]))) = v152
	*(*int32)(unsafe.Add(mBase, uint32(v148)+uint32(_c_F_AutoVacWorkerMain[10]))) = v149
	*(*int64)(unsafe.Add(mBase, uint32(v148)+uint32(_c_F_AutoVacWorkerMain[11]))) = v152
	*(*int32)(unsafe.Add(mBase, uint32(v148)+uint32(_c_F_AutoVacWorkerMain[12]))) = v149
	*(*uint8)(unsafe.Add(mBase, uint32(v148)+uint32(_c_F_AutoVacWorkerMain[13]))) = uint8(v149)
	v163 = v145 | int32(1)
	v165 = v163 * v147
	*(*uint8)(unsafe.Add(mBase, uint32(v165)+uint32(_c_F_AutoVacWorkerMain[7]))) = uint8(v149)
	*(*int32)(unsafe.Add(mBase, uint32(v165)+uint32(_c_F_AutoVacWorkerMain[8]))) = v163
	*(*int64)(unsafe.Add(mBase, uint32(v165)+uint32(_c_F_AutoVacWorkerMain[9]))) = v152
	*(*int32)(unsafe.Add(mBase, uint32(v165)+uint32(_c_F_AutoVacWorkerMain[10]))) = v149
	*(*int64)(unsafe.Add(mBase, uint32(v165)+uint32(_c_F_AutoVacWorkerMain[11]))) = v152
	*(*int32)(unsafe.Add(mBase, uint32(v165)+uint32(_c_F_AutoVacWorkerMain[12]))) = v149
	*(*uint8)(unsafe.Add(mBase, uint32(v165)+uint32(_c_F_AutoVacWorkerMain[13]))) = uint8(v149)
	v180 = v145 | int32(2)
	v182 = v180 * v147
	*(*uint8)(unsafe.Add(mBase, uint32(v182)+uint32(_c_F_AutoVacWorkerMain[7]))) = uint8(v149)
	*(*int32)(unsafe.Add(mBase, uint32(v182)+uint32(_c_F_AutoVacWorkerMain[8]))) = v180
	*(*int64)(unsafe.Add(mBase, uint32(v182)+uint32(_c_F_AutoVacWorkerMain[9]))) = v152
	*(*int32)(unsafe.Add(mBase, uint32(v182)+uint32(_c_F_AutoVacWorkerMain[10]))) = v149
	*(*int64)(unsafe.Add(mBase, uint32(v182)+uint32(_c_F_AutoVacWorkerMain[11]))) = v152
	*(*int32)(unsafe.Add(mBase, uint32(v182)+uint32(_c_F_AutoVacWorkerMain[12]))) = v149
	*(*uint8)(unsafe.Add(mBase, uint32(v182)+uint32(_c_F_AutoVacWorkerMain[13]))) = uint8(v149)
	if v145 != int32(20) {
		goto L39
	} else {
		goto L40
	}
L38:
	;
	v218 = int32(1)
	*(*uint8)(unsafe.Add(mBase, _c_F_AutoVacWorkerMain[14])) = uint8(v218)
	F_pqsignal_be(m, int32(14), int32(1769))
	mBase = m.M
	goto L36
L39:
	;
	v199 = v145 | int32(3)
	v201 = v199 * int32(40)
	v202 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v201)+uint32(_c_F_AutoVacWorkerMain[7]))) = uint8(v202)
	*(*int32)(unsafe.Add(mBase, uint32(v201)+uint32(_c_F_AutoVacWorkerMain[8]))) = v199
	v205 = int64(0)
	*(*int64)(unsafe.Add(mBase, uint32(v201)+uint32(_c_F_AutoVacWorkerMain[9]))) = v205
	*(*int32)(unsafe.Add(mBase, uint32(v201)+uint32(_c_F_AutoVacWorkerMain[10]))) = v202
	*(*int64)(unsafe.Add(mBase, uint32(v201)+uint32(_c_F_AutoVacWorkerMain[11]))) = v205
	*(*int32)(unsafe.Add(mBase, uint32(v201)+uint32(_c_F_AutoVacWorkerMain[12]))) = v202
	*(*uint8)(unsafe.Add(mBase, uint32(v201)+uint32(_c_F_AutoVacWorkerMain[13]))) = uint8(v202)
	v145 = v145 + int32(4)
	goto L37
L40:
	;
	goto L41
L41:
	;
	goto L38
L42:
	;
	v257 = int32(917)
	v259 = m.G0
	v261 = v259 - int32(32)
	m.G0 = v261
	switch int32(919) {
	case 0, 2:
		v271 = v257
		goto L49
	default:
		goto L50
	}
L43:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v228)+12)) = v238
	F_sigemptyset(m, v228+int32(16))
	mBase = m.M
	goto L46
L44:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_AutoVacWorkerMain[15])) = v224
	v238 = int32(_a_F_AutoVacWorkerMain_0)
	goto L43
L46:
	;
	goto L47
L47:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v228)+24)) = int32(268435456)
	v252 = F___sigaction(m, int32(13), v228+int32(12), int32(0))
	mBase = m.M
	m.G0 = v228 + int32(32)
	goto L42
L48:
	;
	v290 = int32(-2)
	v292 = m.G0
	v294 = v292 - int32(32)
	m.G0 = v294
	switch int32(0) {
	case 0, 2:
		v304 = v290
		goto L55
	default:
		goto L56
	}
L49:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v261)+12)) = v271
	F_sigemptyset(m, v261+int32(16))
	mBase = m.M
	goto L52
L50:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_AutoVacWorkerMain[16])) = v257
	v271 = int32(_a_F_AutoVacWorkerMain_0)
	goto L49
L52:
	;
	goto L53
L53:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v261)+24)) = int32(268435456)
	v285 = F___sigaction(m, int32(10), v261+int32(12), int32(0))
	mBase = m.M
	m.G0 = v261 + int32(32)
	goto L48
L54:
	;
	v323 = int32(919)
	v325 = m.G0
	v327 = v325 - int32(32)
	m.G0 = v327
	switch int32(921) {
	case 0, 2:
		v337 = v323
		goto L61
	default:
		goto L62
	}
L55:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v294)+12)) = v304
	F_sigemptyset(m, v294+int32(16))
	mBase = m.M
	goto L58
L56:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_AutoVacWorkerMain[17])) = v290
	v304 = int32(_a_F_AutoVacWorkerMain_0)
	goto L55
L58:
	;
	goto L59
L59:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v294)+24)) = int32(268435456)
	v318 = F___sigaction(m, int32(12), v294+int32(12), int32(0))
	mBase = m.M
	m.G0 = v294 + int32(32)
	goto L54
L60:
	;
	v356 = int32(0)
	v358 = m.G0
	v360 = v358 - int32(32)
	m.G0 = v360
	switch int32(2) {
	case 0, 2:
		v370 = v356
		goto L67
	default:
		goto L68
	}
L61:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v327)+12)) = v337
	F_sigemptyset(m, v327+int32(16))
	mBase = m.M
	goto L64
L62:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_AutoVacWorkerMain[18])) = v323
	v337 = int32(_a_F_AutoVacWorkerMain_0)
	goto L61
L64:
	;
	goto L65
L65:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v327)+24)) = int32(268435456)
	v351 = F___sigaction(m, int32(8), v327+int32(12), int32(0))
	mBase = m.M
	m.G0 = v327 + int32(32)
	goto L60
L66:
	;
	F_InitProcess(m)
	mBase = m.M
	v389 = m.ExcPending
	if v389 != 0 {
		goto L6
	} else {
		goto L72
	}
L67:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v360)+12)) = v370
	F_sigemptyset(m, v360+int32(16))
	mBase = m.M
	goto L69
L68:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_AutoVacWorkerMain[19])) = v356
	v370 = int32(_a_F_AutoVacWorkerMain_0)
	goto L67
L69:
	;
	goto L71
L71:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v360)+24)) = int32(268435457)
	v384 = F___sigaction(m, int32(17), v360+int32(12), int32(0))
	mBase = m.M
	m.G0 = v360 + int32(32)
	goto L66
L72:
	;
	F_BaseInit(m)
	mBase = m.M
	v391 = m.ExcPending
	if v391 != 0 {
		goto L6
	} else {
		goto L73
	}
L73:
	;
	goto L74
L74:
	;
	v393 = v9 + int32(80)
	*(*int32)(unsafe.Add(mBase, uint32(v393)+4)) = int32(1)
	*(*int32)(unsafe.Add(mBase, uint32(v393))) = v9 + int32(12)
	goto L77
L75:
	;
	v401 = int32(0)
	goto L9
L77:
	;
	goto L75
L78:
	;
	v402 = int32(_a_F_AutoVacWorkerMain_1)
	v404 = *(*int32)(unsafe.Add(mBase, _c_F_AutoVacWorkerMain[20]))
	*(*int32)(unsafe.Add(mBase, _c_F_AutoVacWorkerMain[20])) = v404 + int32(1)
	*(*int32)(unsafe.Add(mBase, _c_F_AutoVacWorkerMain[21])) = int32(0)
	F_EmitErrorReport(m)
	mBase = m.M
	v412 = m.ExcPending
	if v412 != 0 {
		goto L6
	} else {
		goto L81
	}
L79:
	;
	goto L80
L80:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_AutoVacWorkerMain[22])) = v9 + int32(80)
	F_pgmem_sigprocmask(m, int32(_a_F_AutoVacWorkerMain_2), int32(0))
	mBase = m.M
	v423 = m.ExcPending
	if v423 != 0 {
		goto L6
	} else {
		goto L83
	}
L81:
	;
	F_proc_exit(m, int32(0))
	mBase = m.M
	v415 = m.ExcPending
	if v415 != 0 {
		goto L6
	} else {
		goto L82
	}
L82:
	;
	goto L3
L83:
	;
	F_SetConfigOption(m, int32(_a_F_AutoVacWorkerMain_3), int32(_a_F_AutoVacWorkerMain_4), int32(5), int32(10))
	mBase = m.M
	v429 = m.ExcPending
	if v429 != 0 {
		goto L6
	} else {
		goto L84
	}
L84:
	;
	F_SetConfigOption(m, int32(_a_F_AutoVacWorkerMain_5), int32(_a_F_AutoVacWorkerMain_6), int32(5), int32(10))
	mBase = m.M
	v435 = m.ExcPending
	if v435 != 0 {
		goto L6
	} else {
		goto L85
	}
L85:
	;
	F_SetConfigOption(m, int32(_a_F_AutoVacWorkerMain_7), int32(_a_F_AutoVacWorkerMain_8), int32(5), int32(10))
	mBase = m.M
	v441 = m.ExcPending
	if v441 != 0 {
		goto L6
	} else {
		goto L86
	}
L86:
	;
	F_SetConfigOption(m, int32(_a_F_AutoVacWorkerMain_9), int32(_a_F_AutoVacWorkerMain_8), int32(5), int32(10))
	mBase = m.M
	v447 = m.ExcPending
	if v447 != 0 {
		goto L6
	} else {
		goto L87
	}
L87:
	;
	F_SetConfigOption(m, int32(_a_F_AutoVacWorkerMain_10), int32(_a_F_AutoVacWorkerMain_8), int32(5), int32(10))
	mBase = m.M
	v453 = m.ExcPending
	if v453 != 0 {
		goto L6
	} else {
		goto L88
	}
L88:
	;
	F_SetConfigOption(m, int32(_a_F_AutoVacWorkerMain_11), int32(_a_F_AutoVacWorkerMain_8), int32(5), int32(10))
	mBase = m.M
	v459 = m.ExcPending
	if v459 != 0 {
		goto L6
	} else {
		goto L89
	}
L89:
	;
	F_SetConfigOption(m, int32(_a_F_AutoVacWorkerMain_12), int32(_a_F_AutoVacWorkerMain_13), int32(5), int32(10))
	mBase = m.M
	v465 = m.ExcPending
	if v465 != 0 {
		goto L6
	} else {
		goto L90
	}
L90:
	;
	v467 = *(*int32)(unsafe.Add(mBase, _c_F_AutoVacWorkerMain[23]))
	if int32(2) <= v467 {
		goto L91
	} else {
		goto L92
	}
L91:
	;
	F_SetConfigOption(m, int32(_a_F_AutoVacWorkerMain_14), int32(_a_F_AutoVacWorkerMain_15), int32(5), int32(10))
	mBase = m.M
	v475 = m.ExcPending
	if v475 != 0 {
		goto L6
	} else {
		goto L94
	}
L92:
	;
	goto L93
L93:
	;
	F_SetConfigOption(m, int32(_a_F_AutoVacWorkerMain_16), int32(_a_F_AutoVacWorkerMain_17), int32(5), int32(10))
	mBase = m.M
	v481 = m.ExcPending
	if v481 != 0 {
		goto L6
	} else {
		goto L95
	}
L94:
	;
	goto L93
L95:
	;
	v483 = *(*int32)(unsafe.Add(mBase, _c_F_AutoVacWorkerMain[24]))
	v487 = F_LWLockAcquire(m, v483+int32(2816), int32(0))
	mBase = m.M
	v488 = m.ExcPending
	if v488 != 0 {
		goto L6
	} else {
		goto L96
	}
L96:
	;
	v490 = *(*int32)(unsafe.Add(mBase, _c_F_AutoVacWorkerMain[25]))
	v491 = *(*int32)(unsafe.Add(mBase, uint32(v490)+32))
	if v491 != 0 {
		goto L98
	} else {
		goto L99
	}
L97:
	;
	F_proc_exit(m, int32(0))
	mBase = m.M
	v619 = m.ExcPending
	if v619 != 0 {
		goto L6
	} else {
		goto L133
	}
L98:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_AutoVacWorkerMain[26])) = v491
	v495 = *(*int32)(unsafe.Add(mBase, _c_F_AutoVacWorkerMain[27]))
	*(*int32)(unsafe.Add(mBase, uint32(v491)+16)) = v495
	v498 = v490 + int32(24)
	v499 = *(*int32)(unsafe.Add(mBase, uint32(v491)+8))
	v500 = *(*int32)(unsafe.Add(mBase, uint32(v490)+28))
	if v500 == int32(0) {
		goto L101
	} else {
		goto L102
	}
L99:
	;
	goto L100
L100:
	;
	v597 = F_errstart(m, int32(19), int32(0))
	mBase = m.M
	v598 = m.ExcPending
	if v598 != 0 {
		goto L6
	} else {
		goto L126
	}
L101:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v498))) = v498
	v504 = v498
	goto L103
L102:
	;
	v504 = v500
	goto L103
L103:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v491))) = v498
	*(*int32)(unsafe.Add(mBase, uint32(v491)+4)) = v504
	*(*int32)(unsafe.Add(mBase, uint32(v504))) = v491
	*(*int32)(unsafe.Add(mBase, uint32(v490)+32)) = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v490)+28)) = v491
	v512 = *(*int32)(unsafe.Add(mBase, _c_F_AutoVacWorkerMain[24]))
	F_LWLockRelease(m, v512+int32(2816))
	mBase = m.M
	v516 = m.ExcPending
	if v516 != 0 {
		goto L6
	} else {
		goto L104
	}
L104:
	;
	F_on_shmem_exit(m, int32(921), int32(0))
	mBase = m.M
	v520 = m.ExcPending
	if v520 != 0 {
		goto L6
	} else {
		goto L105
	}
L105:
	;
	v522 = *(*int32)(unsafe.Add(mBase, _c_F_AutoVacWorkerMain[25]))
	v523 = *(*int32)(unsafe.Add(mBase, uint32(v522)+8))
	if v523 != 0 {
		goto L106
	} else {
		goto L107
	}
L106:
	;
	v525 = F_pgmem_kill(m, v523, int32(12))
	mBase = m.M
	goto L108
L107:
	;
	goto L108
L108:
	;
	if v499 == int32(0) {
		goto L97
	} else {
		goto L109
	}
L109:
	;
	v531 = F_pgstat_get_entry_ref_locked(m, int32(1), v499, int64(0), int32(0))
	mBase = m.M
	v532 = m.ExcPending
	if v532 != 0 {
		goto L6
	} else {
		goto L110
	}
L110:
	;
	v533 = *(*int32)(unsafe.Add(mBase, uint32(v531)+4))
	v537 = m.G0
	v538 = int32(16)
	v539 = v537 - v538
	m.G0 = v539
	F_gettimeofday(m, v539)
	mBase = m.M
	v542 = *(*int64)(unsafe.Add(mBase, uint32(v539)))
	v543 = int64(*(*int32)(unsafe.Add(mBase, uint32(v539)+8)))
	m.G0 = v539 + v538
	goto L111
L111:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v533)+96)) = v543 + v542*int64(1000000) - int64(946684800000000)
	F_pgstat_unlock_entry(m, v531)
	mBase = m.M
	v554 = m.ExcPending
	if v554 != 0 {
		goto L6
	} else {
		goto L112
	}
L112:
	;
	v555 = int32(0)
	v560 = v9 + int32(16)
	F_InitPostgres(m, v555, v499, v555, v555, int32(2), v560)
	mBase = m.M
	v562 = m.ExcPending
	if v562 != 0 {
		goto L6
	} else {
		goto L113
	}
L113:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_AutoVacWorkerMain[28])) = int32(2)
	v566 = F_strlen(m, v560)
	mBase = m.M
	v569 = F_errstart(m, int32(14), int32(0))
	mBase = m.M
	v570 = m.ExcPending
	if v570 != 0 {
		goto L6
	} else {
		goto L114
	}
L114:
	;
	if v569 != 0 {
		goto L115
	} else {
		goto L116
	}
L115:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v9))) = v560
	F_errmsg_internal(m, int32(_a_F_AutoVacWorkerMain_18), v9)
	mBase = m.M
	v574 = m.ExcPending
	if v574 != 0 {
		goto L6
	} else {
		goto L118
	}
L116:
	;
	goto L117
L117:
	;
	v581 = *(*int32)(unsafe.Add(mBase, _c_F_AutoVacWorkerMain[29]))
	if v581 != 0 {
		goto L120
	} else {
		goto L121
	}
L118:
	;
	F_errfinish(m, int32(_a_F_AutoVacWorkerMain_19), int32(1582), int32(_a_F_AutoVacWorkerMain_20))
	mBase = m.M
	v579 = m.ExcPending
	if v579 != 0 {
		goto L6
	} else {
		goto L119
	}
L119:
	;
	goto L117
L120:
	;
	F_pg_usleep(m, v581*int32(_a_F_AutoVacWorkerMain_21))
	mBase = m.M
	goto L122
L121:
	;
	goto L122
L122:
	;
	v586 = F_ReadNextFullTransactionId(m)
	mBase = m.M
	v587 = m.ExcPending
	if v587 != 0 {
		goto L6
	} else {
		goto L123
	}
L123:
	;
	*(*uint32)(unsafe.Add(mBase, _c_F_AutoVacWorkerMain[30])) = uint32(v586)
	v590 = F_ReadNextMultiXactId(m)
	mBase = m.M
	v591 = m.ExcPending
	if v591 != 0 {
		goto L6
	} else {
		goto L124
	}
L124:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_AutoVacWorkerMain[31])) = v590
	F_do_autovacuum(m)
	mBase = m.M
	v594 = m.ExcPending
	if v594 != 0 {
		goto L6
	} else {
		goto L125
	}
L125:
	;
	goto L97
L126:
	;
	if v597 != 0 {
		goto L127
	} else {
		goto L128
	}
L127:
	;
	F_errmsg_internal(m, int32(_a_F_AutoVacWorkerMain_22), int32(0))
	mBase = m.M
	v602 = m.ExcPending
	if v602 != 0 {
		goto L6
	} else {
		goto L130
	}
L128:
	;
	goto L129
L129:
	;
	v609 = *(*int32)(unsafe.Add(mBase, _c_F_AutoVacWorkerMain[24]))
	F_LWLockRelease(m, v609+int32(2816))
	mBase = m.M
	v613 = m.ExcPending
	if v613 != 0 {
		goto L6
	} else {
		goto L132
	}
L130:
	;
	F_errfinish(m, int32(_a_F_AutoVacWorkerMain_19), int32(1549), int32(_a_F_AutoVacWorkerMain_20))
	mBase = m.M
	v607 = m.ExcPending
	if v607 != 0 {
		goto L6
	} else {
		goto L131
	}
L131:
	;
	goto L129
L132:
	;
	goto L97
L133:
	;
	goto L5
L134:
	;
	v630 = int32(v626)
	m.G0 = v9
	v632 = *(*int32)(unsafe.Add(mBase, uint32(v630)+4))
	v633 = *(*int32)(unsafe.Add(mBase, uint32(v630)))
	v636 = *(*int32)(unsafe.Add(mBase, uint32(v633)))
	if v9+int32(12) == v636 {
		goto L137
	} else {
		goto L138
	}
L135:
	;
	m.ExcPending = 1
	goto L143
L136:
	;
	if v640 == int32(0) {
		goto L140
	} else {
		goto L141
	}
L137:
	;
	v638 = *(*int32)(unsafe.Add(mBase, uint32(v633)+4))
	v640 = v638
	goto L139
L138:
	;
	v640 = int32(0)
	goto L139
L139:
	;
	goto L136
L140:
	;
	F___wasm_longjmp(m, v633, v632)
	mBase = m.M
	v644 = m.ExcPending
	if v644 != 0 {
		goto L143
	} else {
		goto L144
	}
L141:
	;
	goto L142
L142:
	;
	v12 = v640
	v14 = v632
	goto L1
L143:
	;
	return
L144:
	;
	base.Wasm_trap_unreachable()
	for {
	}
}
func F_AutoVacuumingActive(m *base.Module) int32 {
	mBase := m.M
	_ = mBase
	var v2 int32
	_ = v2
	var v4 int32
	_ = v4
	v2 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_AutoVacuumingActive[0])))
	v4 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_AutoVacuumingActive[1])))
	return v2 & v4 & int32(1)
}
func F_abs_interval(m *base.Module, l0 int32) int32 {
	var v5 int32
	_ = v5
	var v8 int32
	_ = v8
	var v11 int32
	_ = v11
	var v12 int32
	_ = v12
	var v13 int32
	_ = v13
	v5 = F_DirectFunctionCall2Coll(m, int32(2444), int32(0), l0, int32(_a_F_abs_interval_0))
	v8 = m.ExcPending
	if v8 != 0 {
		return int32(0)
	} else {
		if v5 != 0 {
			v11 = F_DirectFunctionCall1Coll(m, int32(2448), int32(0), l0)
			v12 = m.ExcPending
			if v12 != 0 {
				return int32(0)
			} else {
				v13 = v11
				return v13
			}
		} else {
			v13 = l0
			return v13
		}
	}
}
func F_access(m *base.Module, l0 int32, l1 int32) int32 {
	mBase := m.M
	_ = mBase
	var v5 int32
	_ = v5
	var v13 int32
	_ = v13
	v5 = m.Env.X__syscall_faccessat(m, int32(-100), l0, l1, int32(0))
	mBase = m.M
	if base.Ui32(int32(-4095)) <= base.Ui32(v5) {
		*(*int32)(unsafe.Add(mBase, _c_F_access[0])) = int32(0) - v5
		v13 = int32(-1)
	} else {
		v13 = v5
	}
	return v13
}
func F_accum_sum_final(m *base.Module, l0 int32, l1 int32) {
	mBase := m.M
	_ = mBase
	var v3 int32
	_ = v3
	var v15 int32
	_ = v15
	var v17 int32
	_ = v17
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
	var v32 int64
	_ = v32
	var v35 int64
	_ = v35
	var v40 int32
	_ = v40
	var v42 int32
	_ = v42
	var v45 int32
	_ = v45
	var v54 int32
	_ = v54
	var v57 int32
	_ = v57
	var v61 int32
	_ = v61
	var v66 int32
	_ = v66
	var v67 int32
	_ = v67
	var v68 int32
	_ = v68
	var v73 int32
	_ = v73
	var v77 int32
	_ = v77
	var v78 int32
	_ = v78
	var v82 int32
	_ = v82
	var v83 int32
	_ = v83
	var v84 int32
	_ = v84
	var v88 int32
	_ = v88
	var v92 int32
	_ = v92
	var v93 int32
	_ = v93
	var v95 int32
	_ = v95
	var v96 int32
	_ = v96
	var v98 int32
	_ = v98
	var v106 int32
	_ = v106
	var v109 int32
	_ = v109
	var v118 int32
	_ = v118
	var v119 int32
	_ = v119
	var v120 int32
	_ = v120
	var v122 int32
	_ = v122
	var v125 int32
	_ = v125
	var v129 int32
	_ = v129
	var v143 int32
	_ = v143
	var v145 int32
	_ = v145
	var v146 int32
	_ = v146
	var v157 int32
	_ = v157
	var v158 int32
	_ = v158
	var v161 int32
	_ = v161
	var v170 int32
	_ = v170
	var v171 int32
	_ = v171
	var v172 int32
	_ = v172
	var v177 int32
	_ = v177
	var v181 int32
	_ = v181
	var v182 int32
	_ = v182
	var v186 int32
	_ = v186
	var v187 int32
	_ = v187
	var v188 int32
	_ = v188
	var v192 int32
	_ = v192
	var v196 int32
	_ = v196
	var v197 int32
	_ = v197
	var v199 int32
	_ = v199
	var v200 int32
	_ = v200
	var v202 int32
	_ = v202
	var v209 int32
	_ = v209
	var v210 int32
	_ = v210
	var v222 int32
	_ = v222
	var v223 int32
	_ = v223
	var v224 int32
	_ = v224
	var v226 int32
	_ = v226
	var v229 int32
	_ = v229
	var v233 int32
	_ = v233
	var v247 int32
	_ = v247
	var v265 int32
	_ = v265
	var v271 int32
	_ = v271
	var v282 int32
	_ = v282
	var v285 int32
	_ = v285
	var v288 int32
	_ = v288
	var v295 int32
	_ = v295
	var v296 int32
	_ = v296
	var v299 int32
	_ = v299
	var v302 int32
	_ = v302
	var v303 int32
	_ = v303
	var v306 int32
	_ = v306
	var v311 int32
	_ = v311
	var v323 int32
	_ = v323
	var v324 int32
	_ = v324
	var v327 int32
	_ = v327
	var v328 int32
	_ = v328
	var v330 int32
	_ = v330
	var v333 int32
	_ = v333
	var v335 int32
	_ = v335
	var v338 int32
	_ = v338
	var v339 int32
	_ = v339
	var v358 int32
	_ = v358
	var v359 int32
	_ = v359
	var v360 int32
	_ = v360
	var v366 int32
	_ = v366
	var v368 int32
	_ = v368
	var v380 int32
	_ = v380
	var v383 int32
	_ = v383
	var v400 int32
	_ = v400
	var v401 int32
	_ = v401
	var v405 int32
	_ = v405
	var v406 int32
	_ = v406
	var v415 int32
	_ = v415
	var v432 int32
	_ = v432
	var v434 int32
	_ = v434
	v3 = int32(0)
	v15 = m.G0
	v17 = v15 - int32(48)
	m.G0 = v17
	v19 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	if v19 == v3 {
		goto L2
	} else {
		goto L3
	}
L1:
	;
	m.G0 = v17 + int32(48)
	return
L2:
	;
	v23 = F_palloc(m, int32(2))
	mBase = m.M
	v24 = m.ExcPending
	if v24 != 0 {
		goto L5
	} else {
		goto L6
	}
L3:
	;
	goto L4
L4:
	;
	v40 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	if v40 != 0 {
		goto L11
	} else {
		goto L12
	}
L5:
	;
	return
L6:
	;
	v25 = int32(0)
	*(*uint16)(unsafe.Add(mBase, uint32(v23))) = uint16(v25)
	v27 = *(*int32)(unsafe.Add(mBase, uint32(l1)+16))
	if v27 != 0 {
		goto L7
	} else {
		goto L8
	}
L7:
	;
	F_pfree(m, v27)
	mBase = m.M
	v29 = m.ExcPending
	if v29 != 0 {
		goto L5
	} else {
		goto L10
	}
L8:
	;
	goto L9
L9:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l1)+16)) = v23
	v32 = *(*int64)(unsafe.Add(mBase, _c_F_accum_sum_final[0]))
	*(*int64)(unsafe.Add(mBase, uint32(l1)+8)) = v32
	v35 = *(*int64)(unsafe.Add(mBase, _c_F_accum_sum_final[1]))
	*(*int64)(unsafe.Add(mBase, uint32(l1))) = v35
	*(*int32)(unsafe.Add(mBase, uint32(l1)+20)) = v23 + int32(2)
	goto L1
L10:
	;
	goto L9
L11:
	;
	v42 = v19 - int32(1)
	if v42 < int32(0) {
		goto L14
	} else {
		goto L15
	}
L12:
	;
	v271 = v19
	goto L13
L13:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v17)+24)) = v271
	*(*int32)(unsafe.Add(mBase, uint32(v17))) = v271
	v282 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	*(*int32)(unsafe.Add(mBase, uint32(v17)+28)) = v282
	*(*int32)(unsafe.Add(mBase, uint32(v17)+4)) = v282
	v285 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	*(*int32)(unsafe.Add(mBase, uint32(v17)+36)) = v285
	*(*int32)(unsafe.Add(mBase, uint32(v17)+12)) = v285
	v288 = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v17)+32)) = v288
	*(*int32)(unsafe.Add(mBase, uint32(v17)+8)) = int32(_a_F_accum_sum_final_0)
	v295 = F_palloc(m, v271<<(uint(int32(1))%32))
	mBase = m.M
	v296 = m.ExcPending
	if v296 != 0 {
		goto L5
	} else {
		goto L56
	}
L14:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+12)) = int32(0)
	v265 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v271 = v265
	goto L13
L15:
	;
	v45 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	if v42 != 0 {
		goto L17
	} else {
		goto L18
	}
L16:
	;
	if int32(0) < v129 {
		goto L34
	} else {
		goto L35
	}
L17:
	;
	v54 = v42
	v57 = v3
	v61 = v3
	goto L20
L18:
	;
	v106 = v42
	v109 = v3
	goto L19
L19:
	;
	v118 = v45 + v106<<(uint(int32(2))%32)
	v119 = *(*int32)(unsafe.Add(mBase, uint32(v118)))
	v120 = v119 + v109
	v122 = base.I32_rem_u_s(v120, int32(_a_F_accum_sum_final_1))
	if int32(_a_F_accum_sum_final_2) < v120 {
		goto L31
	} else {
		goto L32
	}
L20:
	;
	v66 = v45 + v54<<(uint(int32(2))%32)
	v67 = *(*int32)(unsafe.Add(mBase, uint32(v66)))
	v68 = v67 + v57
	if v68 < int32(_a_F_accum_sum_final_1) {
		goto L23
	} else {
		goto L24
	}
L21:
	;
	if v19&int32(1) == int32(0) {
		v129 = v92
		goto L16
	} else {
		goto L30
	}
L22:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v66))) = v77
	v82 = v66 - int32(4)
	v83 = *(*int32)(unsafe.Add(mBase, uint32(v82)))
	v84 = v83 + v78
	if int32(_a_F_accum_sum_final_1) <= v84 {
		goto L26
	} else {
		goto L27
	}
L23:
	;
	v77 = v68
	v78 = int32(0)
	goto L22
L24:
	;
	goto L25
L25:
	;
	v73 = base.I32_div_u_s(v68, int32(_a_F_accum_sum_final_1))
	v77 = v73*int32(-10000) + v68
	v78 = v73
	goto L22
L26:
	;
	v88 = base.I32_div_u_s(v84, int32(_a_F_accum_sum_final_1))
	v92 = v88*int32(-10000) + v84
	v93 = v88
	goto L28
L27:
	;
	v92 = v84
	v93 = int32(0)
	goto L28
L28:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v82))) = v92
	v95 = int32(2)
	v96 = v54 - v95
	v98 = v61 + v95
	if v98 != v19&int32(-2) {
		v54 = v96
		v57 = v93
		v61 = v98
		goto L20
	} else {
		goto L29
	}
L29:
	;
	goto L21
L30:
	;
	v106 = v96
	v109 = v93
	goto L19
L31:
	;
	v125 = v122
	goto L33
L32:
	;
	v125 = v120
	goto L33
L33:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v118))) = v125
	v129 = v125
	goto L16
L34:
	;
	v143 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(l0)+16)) = uint8(v143)
	goto L36
L35:
	;
	goto L36
L36:
	;
	v145 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
	v146 = int32(0)
	if v19 != int32(1) {
		goto L38
	} else {
		goto L39
	}
L37:
	;
	if v233 <= int32(0) {
		goto L14
	} else {
		goto L55
	}
L38:
	;
	v157 = v42
	v158 = v146
	v161 = int32(0)
	goto L41
L39:
	;
	v209 = v42
	v210 = v146
	goto L40
L40:
	;
	v222 = v145 + v209<<(uint(int32(2))%32)
	v223 = *(*int32)(unsafe.Add(mBase, uint32(v222)))
	v224 = v223 + v210
	v226 = base.I32_rem_u_s(v224, int32(_a_F_accum_sum_final_1))
	if int32(_a_F_accum_sum_final_2) < v224 {
		goto L52
	} else {
		goto L53
	}
L41:
	;
	v170 = v145 + v157<<(uint(int32(2))%32)
	v171 = *(*int32)(unsafe.Add(mBase, uint32(v170)))
	v172 = v171 + v158
	if v172 < int32(_a_F_accum_sum_final_1) {
		goto L44
	} else {
		goto L45
	}
L42:
	;
	if v19&int32(1) == int32(0) {
		v233 = v196
		goto L37
	} else {
		goto L51
	}
L43:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v170))) = v181
	v186 = v170 - int32(4)
	v187 = *(*int32)(unsafe.Add(mBase, uint32(v186)))
	v188 = v187 + v182
	if int32(_a_F_accum_sum_final_1) <= v188 {
		goto L47
	} else {
		goto L48
	}
L44:
	;
	v181 = v172
	v182 = int32(0)
	goto L43
L45:
	;
	goto L46
L46:
	;
	v177 = base.I32_div_u_s(v172, int32(_a_F_accum_sum_final_1))
	v181 = v177*int32(-10000) + v172
	v182 = v177
	goto L43
L47:
	;
	v192 = base.I32_div_u_s(v188, int32(_a_F_accum_sum_final_1))
	v196 = v192*int32(-10000) + v188
	v197 = v192
	goto L49
L48:
	;
	v196 = v188
	v197 = int32(0)
	goto L49
L49:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v186))) = v196
	v199 = int32(2)
	v200 = v157 - v199
	v202 = v161 + v199
	if v202 != v19&int32(-2) {
		v157 = v200
		v158 = v197
		v161 = v202
		goto L41
	} else {
		goto L50
	}
L50:
	;
	goto L42
L51:
	;
	v209 = v200
	v210 = v197
	goto L40
L52:
	;
	v229 = v226
	goto L54
L53:
	;
	v229 = v224
	goto L54
L54:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v222))) = v229
	v233 = v229
	goto L37
L55:
	;
	v247 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(l0)+16)) = uint8(v247)
	goto L14
L56:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v17)+40)) = v295
	*(*int32)(unsafe.Add(mBase, uint32(v17)+44)) = v295
	v299 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v302 = F_palloc(m, v299<<(uint(int32(1))%32))
	mBase = m.M
	v303 = m.ExcPending
	if v303 != 0 {
		goto L5
	} else {
		goto L57
	}
L57:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v17)+16)) = v302
	*(*int32)(unsafe.Add(mBase, uint32(v17)+20)) = v302
	v306 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	if int32(0) < v306 {
		goto L58
	} else {
		goto L59
	}
L58:
	;
	v311 = v288
	goto L61
L59:
	;
	goto L60
L60:
	;
	F_add_var(m, v17+int32(24), v17, l1)
	mBase = m.M
	v358 = m.ExcPending
	if v358 != 0 {
		goto L5
	} else {
		goto L64
	}
L61:
	;
	v323 = int32(1)
	v324 = v311 << (uint(v323) % 32)
	v327 = v311 << (uint(int32(2)) % 32)
	v328 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v330 = *(*int32)(unsafe.Add(mBase, uint32(v327+v328)))
	*(*uint16)(unsafe.Add(mBase, uint32(v295+v324))) = uint16(v330)
	v333 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
	v335 = *(*int32)(unsafe.Add(mBase, uint32(v333+v327)))
	*(*uint16)(unsafe.Add(mBase, uint32(v302+v324))) = uint16(v335)
	v338 = v311 + v323
	v339 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	if v338 < v339 {
		v311 = v338
		goto L61
	} else {
		goto L63
	}
L62:
	;
	goto L60
L63:
	;
	goto L62
L64:
	;
	v359 = *(*int32)(unsafe.Add(mBase, uint32(l1)+20))
	v360 = *(*int32)(unsafe.Add(mBase, uint32(l1)))
	if int32(0) < v360 {
		goto L67
	} else {
		goto L68
	}
L65:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l1))) = v434
	*(*int32)(unsafe.Add(mBase, uint32(l1)+20)) = v432
	goto L1
L66:
	;
	*(*int64)(unsafe.Add(mBase, uint32(l1)+4)) = int64(0)
	v432 = v415
	v434 = int32(0)
	goto L65
L67:
	;
	v366 = v359
	v368 = v360
	goto L70
L68:
	;
	goto L69
L69:
	;
	if v360 != 0 {
		v432 = v359
		v434 = v360
		goto L65
	} else {
		goto L80
	}
L70:
	;
	v380 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v366))))
	if v380 != 0 {
		goto L72
	} else {
		goto L73
	}
L71:
	;
	v415 = v359 + v360<<(uint(int32(1))%32)
	goto L66
L72:
	;
	v383 = v368
	goto L75
L73:
	;
	goto L74
L74:
	;
	v405 = *(*int32)(unsafe.Add(mBase, uint32(l1)+4))
	v406 = int32(1)
	*(*int32)(unsafe.Add(mBase, uint32(l1)+4)) = v405 - v406
	if v406 < v368 {
		v366 = v366 + int32(2)
		v368 = v368 - v406
		goto L70
	} else {
		goto L79
	}
L75:
	;
	v400 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v366+v383<<(uint(int32(1))%32)-int32(2)))))
	if v400 != 0 {
		v432 = v366
		v434 = v383
		goto L65
	} else {
		goto L77
	}
L77:
	;
	v401 = int32(1)
	if v401 < v383 {
		v383 = v383 - v401
		goto L75
	} else {
		goto L78
	}
L78:
	;
	v415 = v366
	goto L66
L79:
	;
	goto L71
L80:
	;
	v415 = v359
	goto L66
}
func F_aclcheck_error_type(m *base.Module, l0 int32, l1 int32) {
	var v4 int32
	_ = v4
	var v5 int32
	_ = v5
	var v6 int32
	_ = v6
	var v7 int32
	_ = v7
	var v8 int32
	_ = v8
	var v10 int32
	_ = v10
	v4 = F_get_element_type(m, l1)
	v5 = m.ExcPending
	if v5 != 0 {
		return
	} else {
		if v4 != 0 {
			v6 = v4
		} else {
			v6 = l1
		}
		v7 = F_format_type_be(m, v6)
		v8 = m.ExcPending
		if v8 != 0 {
			return
		} else {
			F_aclcheck_error(m, l0, int32(49), v7)
			v10 = m.ExcPending
			if v10 != 0 {
				return
			} else {
				return
			}
		}
	}
}
func F_acldefault_sql(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v5 int32
	_ = v5
	var v7 int32
	_ = v7
	var v9 int32
	_ = v9
	var v11 int32
	_ = v11
	var v32 int32
	_ = v32
	var v37 int32
	_ = v37
	var v42 int32
	_ = v42
	var v44 int32
	_ = v44
	var v45 int32
	_ = v45
	var v46 int32
	_ = v46
	v5 = m.G0
	v7 = v5 - int32(16)
	m.G0 = v7
	v9 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
	v11 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	switch v11&int32(255) - int32(70) {
	case 0:
		v44 = int32(16)
		v45 = F_acldefault(m, v44, v9)
		mBase = m.M
		v46 = m.ExcPending
		if v46 != 0 {
			return int32(0)
		} else {
			m.G0 = v7 + int32(16)
			return v45
		}
	default:
		F_errstart_cold(m, int32(21), int32(0))
		mBase = m.M
		v32 = m.ExcPending
		if v32 != 0 {
			return int32(0)
		} else {
			*(*int32)(unsafe.Add(mBase, uint32(v7))) = base.I32_extend8_s(v11)
			F_errmsg_internal(m, int32(_a_F_acldefault_sql_0), v7)
			mBase = m.M
			v37 = m.ExcPending
			if v37 != 0 {
				return int32(0)
			} else {
				F_errfinish(m, int32(_a_F_acldefault_sql_1), int32(969), int32(_a_F_acldefault_sql_2))
				mBase = m.M
				v42 = m.ExcPending
				if v42 != 0 {
					return int32(0)
				} else {
					base.Wasm_trap_unreachable()
					for {
					}
				}
			}
		}
	case 6:
		v44 = int32(22)
		v45 = F_acldefault(m, v44, v9)
		mBase = m.M
		v46 = m.ExcPending
		if v46 != 0 {
			return int32(0)
		} else {
			m.G0 = v7 + int32(16)
			return v45
		}
	case 13:
		v44 = int32(17)
		v45 = F_acldefault(m, v44, v9)
		mBase = m.M
		v46 = m.ExcPending
		if v46 != 0 {
			return int32(0)
		} else {
			m.G0 = v7 + int32(16)
			return v45
		}
	case 14:
		v44 = int32(49)
		v45 = F_acldefault(m, v44, v9)
		mBase = m.M
		v46 = m.ExcPending
		if v46 != 0 {
			return int32(0)
		} else {
			m.G0 = v7 + int32(16)
			return v45
		}
	case 29:
		v44 = int32(6)
		v45 = F_acldefault(m, v44, v9)
		mBase = m.M
		v46 = m.ExcPending
		if v46 != 0 {
			return int32(0)
		} else {
			m.G0 = v7 + int32(16)
			return v45
		}
	case 30:
		v44 = int32(9)
		v45 = F_acldefault(m, v44, v9)
		mBase = m.M
		v46 = m.ExcPending
		if v46 != 0 {
			return int32(0)
		} else {
			m.G0 = v7 + int32(16)
			return v45
		}
	case 32:
		v44 = int32(19)
		v45 = F_acldefault(m, v44, v9)
		mBase = m.M
		v46 = m.ExcPending
		if v46 != 0 {
			return int32(0)
		} else {
			m.G0 = v7 + int32(16)
			return v45
		}
	case 38:
		v44 = int32(21)
		v45 = F_acldefault(m, v44, v9)
		mBase = m.M
		v46 = m.ExcPending
		if v46 != 0 {
			return int32(0)
		} else {
			m.G0 = v7 + int32(16)
			return v45
		}
	case 40:
		v44 = int32(36)
		v45 = F_acldefault(m, v44, v9)
		mBase = m.M
		v46 = m.ExcPending
		if v46 != 0 {
			return int32(0)
		} else {
			m.G0 = v7 + int32(16)
			return v45
		}
	case 42:
		v44 = int32(27)
		v45 = F_acldefault(m, v44, v9)
		mBase = m.M
		v46 = m.ExcPending
		if v46 != 0 {
			return int32(0)
		} else {
			m.G0 = v7 + int32(16)
			return v45
		}
	case 44:
		v44 = int32(41)
		v45 = F_acldefault(m, v44, v9)
		mBase = m.M
		v46 = m.ExcPending
		if v46 != 0 {
			return int32(0)
		} else {
			m.G0 = v7 + int32(16)
			return v45
		}
	case 45:
		v44 = int32(37)
		v45 = F_acldefault(m, v44, v9)
		mBase = m.M
		v46 = m.ExcPending
		if v46 != 0 {
			return int32(0)
		} else {
			m.G0 = v7 + int32(16)
			return v45
		}
	case 46:
		v44 = int32(42)
		v45 = F_acldefault(m, v44, v9)
		mBase = m.M
		v46 = m.ExcPending
		if v46 != 0 {
			return int32(0)
		} else {
			m.G0 = v7 + int32(16)
			return v45
		}
	}
}
func F_aclitemin(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v9 int64
	_ = v9
	var v13 int32
	_ = v13
	var v15 int32
	_ = v15
	var v17 int32
	_ = v17
	var v18 int32
	_ = v18
	var v20 int32
	_ = v20
	var v23 int32
	_ = v23
	var v25 int32
	_ = v25
	var v26 int32
	_ = v26
	var v27 int32
	_ = v27
	var v30 int32
	_ = v30
	var v33 int32
	_ = v33
	var v36 int32
	_ = v36
	var v42 int32
	_ = v42
	var v43 int32
	_ = v43
	var v49 int32
	_ = v49
	var v50 int32
	_ = v50
	var v55 int32
	_ = v55
	var v61 int32
	_ = v61
	var v65 int32
	_ = v65
	var v70 int32
	_ = v70
	var v73 int32
	_ = v73
	var v74 int32
	_ = v74
	var v77 int32
	_ = v77
	var v80 int32
	_ = v80
	var v81 int32
	_ = v81
	var v86 int32
	_ = v86
	var v90 int32
	_ = v90
	var v94 int32
	_ = v94
	var v99 int32
	_ = v99
	var v100 int32
	_ = v100
	var v103 int32
	_ = v103
	var v105 int32
	_ = v105
	var v106 int32
	_ = v106
	var v109 int32
	_ = v109
	var v110 int32
	_ = v110
	var v115 int32
	_ = v115
	var v119 int32
	_ = v119
	var v124 int32
	_ = v124
	var v137 int32
	_ = v137
	var v139 int32
	_ = v139
	var v142 int64
	_ = v142
	var v143 int64
	_ = v143
	var v144 int64
	_ = v144
	var v162 int32
	_ = v162
	var v163 int32
	_ = v163
	var v168 int32
	_ = v168
	var v173 int32
	_ = v173
	var v178 int32
	_ = v178
	var v180 int64
	_ = v180
	var v181 int64
	_ = v181
	var v183 int32
	_ = v183
	var v184 int64
	_ = v184
	var v185 int32
	_ = v185
	var v203 int32
	_ = v203
	var v204 int32
	_ = v204
	var v209 int64
	_ = v209
	var v212 int32
	_ = v212
	var v219 int32
	_ = v219
	var v220 int32
	_ = v220
	var v223 int32
	_ = v223
	var v224 int32
	_ = v224
	var v226 int32
	_ = v226
	var v227 int32
	_ = v227
	var v232 int32
	_ = v232
	var v238 int32
	_ = v238
	var v243 int32
	_ = v243
	var v246 int32
	_ = v246
	var v253 int32
	_ = v253
	var v254 int32
	_ = v254
	var v257 int32
	_ = v257
	var v260 int32
	_ = v260
	var v261 int32
	_ = v261
	var v266 int32
	_ = v266
	var v270 int32
	_ = v270
	var v275 int32
	_ = v275
	var v278 int32
	_ = v278
	var v279 int32
	_ = v279
	var v282 int32
	_ = v282
	var v283 int32
	_ = v283
	var v285 int32
	_ = v285
	var v286 int32
	_ = v286
	var v291 int32
	_ = v291
	var v297 int32
	_ = v297
	var v302 int32
	_ = v302
	var v307 int32
	_ = v307
	var v308 int32
	_ = v308
	var v313 int32
	_ = v313
	var v320 int32
	_ = v320
	var v325 int32
	_ = v325
	var v327 int32
	_ = v327
	var v334 int32
	_ = v334
	var v342 int32
	_ = v342
	var v354 int32
	_ = v354
	var v355 int32
	_ = v355
	var v356 int32
	_ = v356
	var v361 int32
	_ = v361
	var v365 int32
	_ = v365
	var v370 int32
	_ = v370
	var v383 int32
	_ = v383
	var v391 int32
	_ = v391
	v9 = int64(0)
	v13 = m.G0
	v15 = v13 - int32(208)
	m.G0 = v15
	v17 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v18 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v20 = F_palloc(m, int32(16))
	mBase = m.M
	v23 = m.ExcPending
	if v23 != 0 {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	return int32(0)
L2:
	;
	v25 = v15 + int32(144)
	v26 = F_getid(m, v18, v25, v17)
	mBase = m.M
	v27 = m.ExcPending
	if v27 != 0 {
		goto L1
	} else {
		goto L5
	}
L3:
	;
	m.G0 = v15 + int32(208)
	return v391
L4:
	;
	v383 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(l0)+16)) = uint8(v383)
	v391 = int32(0)
	goto L3
L5:
	;
	if v26 == int32(0) {
		goto L4
	} else {
		goto L6
	}
L6:
	;
	v30 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v26))))
	if v30 != int32(61) {
		goto L11
	} else {
		goto L12
	}
L7:
	;
	v212 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v15)+144)))
	if v212 == int32(0) {
		goto L68
	} else {
		goto L69
	}
L8:
	;
	v137 = v105
	v139 = v106
	v142 = v9
	v143 = v9
	v144 = v9
	goto L42
L9:
	;
	if base.Ui32(int32(26)) <= base.Ui32((v106|int32(32)-int32(97))&int32(255)) {
		v203 = v103
		v204 = v105
		v209 = v9
		goto L7
	} else {
		goto L41
	}
L10:
	;
	v109 = F_errsave_start(m, v17)
	mBase = m.M
	v110 = m.ExcPending
	if v110 != 0 {
		goto L1
	} else {
		goto L36
	}
L11:
	;
	v33 = *(*int32)(unsafe.Add(mBase, uint32(v15)+144))
	v36 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v15)+148)))
	if v33^int32(1970238055)|(v36^int32(112)) == int32(0) {
		goto L14
	} else {
		goto L15
	}
L12:
	;
	v103 = v26
	goto L13
L13:
	;
	v105 = v103 + int32(1)
	v106 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v103)+1)))
	if v106 != int32(42) {
		goto L9
	} else {
		goto L35
	}
L14:
	;
	v73 = F_getid(m, v26, v15+int32(144), v17)
	mBase = m.M
	v74 = m.ExcPending
	if v74 != 0 {
		goto L1
	} else {
		goto L23
	}
L15:
	;
	v42 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v15)+148)))
	v43 = *(*int32)(unsafe.Add(mBase, uint32(v15)+144))
	if v42|(v43^int32(1919251317)) == int32(0) {
		goto L14
	} else {
		goto L16
	}
L16:
	;
	v49 = F_errsave_start(m, v17)
	mBase = m.M
	v50 = m.ExcPending
	if v50 != 0 {
		goto L1
	} else {
		goto L17
	}
L17:
	;
	if v49 == int32(0) {
		goto L4
	} else {
		goto L18
	}
L18:
	;
	F_errcode(m, int32(33685634))
	mBase = m.M
	v55 = m.ExcPending
	if v55 != 0 {
		goto L1
	} else {
		goto L19
	}
L19:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v15)+64)) = v25
	F_errmsg(m, int32(_a_F_aclitemin_0), v15-int32(-64))
	mBase = m.M
	v61 = m.ExcPending
	if v61 != 0 {
		goto L1
	} else {
		goto L20
	}
L20:
	;
	F_errhint(m, int32(_a_F_aclitemin_1), int32(0))
	mBase = m.M
	v65 = m.ExcPending
	if v65 != 0 {
		goto L1
	} else {
		goto L21
	}
L21:
	;
	F_errsave_finish(m, v17, int32(_a_F_aclitemin_2), int32(294), int32(_a_F_aclitemin_3))
	mBase = m.M
	v70 = m.ExcPending
	if v70 != 0 {
		goto L1
	} else {
		goto L22
	}
L22:
	;
	goto L4
L23:
	;
	if v73 == int32(0) {
		goto L4
	} else {
		goto L24
	}
L24:
	;
	v77 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v15)+144)))
	if v77 == int32(0) {
		goto L25
	} else {
		goto L26
	}
L25:
	;
	v80 = F_errsave_start(m, v17)
	mBase = m.M
	v81 = m.ExcPending
	if v81 != 0 {
		goto L1
	} else {
		goto L28
	}
L26:
	;
	goto L27
L27:
	;
	v100 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v73))))
	if v100 != int32(61) {
		goto L10
	} else {
		goto L34
	}
L28:
	;
	if v80 == int32(0) {
		goto L4
	} else {
		goto L29
	}
L29:
	;
	F_errcode(m, int32(33685634))
	mBase = m.M
	v86 = m.ExcPending
	if v86 != 0 {
		goto L1
	} else {
		goto L30
	}
L30:
	;
	F_errmsg(m, int32(_a_F_aclitemin_4), int32(0))
	mBase = m.M
	v90 = m.ExcPending
	if v90 != 0 {
		goto L1
	} else {
		goto L31
	}
L31:
	;
	F_errhint(m, int32(_a_F_aclitemin_5), int32(0))
	mBase = m.M
	v94 = m.ExcPending
	if v94 != 0 {
		goto L1
	} else {
		goto L32
	}
L32:
	;
	F_errsave_finish(m, v17, int32(_a_F_aclitemin_2), int32(303), int32(_a_F_aclitemin_3))
	mBase = m.M
	v99 = m.ExcPending
	if v99 != 0 {
		goto L1
	} else {
		goto L33
	}
L33:
	;
	goto L4
L34:
	;
	v103 = v73
	goto L13
L35:
	;
	goto L8
L36:
	;
	if v109 == int32(0) {
		goto L4
	} else {
		goto L37
	}
L37:
	;
	F_errcode(m, int32(33685634))
	mBase = m.M
	v115 = m.ExcPending
	if v115 != 0 {
		goto L1
	} else {
		goto L38
	}
L38:
	;
	F_errmsg(m, int32(_a_F_aclitemin_6), int32(0))
	mBase = m.M
	v119 = m.ExcPending
	if v119 != 0 {
		goto L1
	} else {
		goto L39
	}
L39:
	;
	F_errsave_finish(m, v17, int32(_a_F_aclitemin_2), int32(309), int32(_a_F_aclitemin_3))
	mBase = m.M
	v124 = m.ExcPending
	if v124 != 0 {
		goto L1
	} else {
		goto L40
	}
L40:
	;
	goto L4
L41:
	;
	goto L8
L42:
	;
	switch v139 - int32(42) {
	case 0:
		goto L45
	default:
		goto L46
	case 23:
		goto L48
	case 25:
		goto L52
	case 26:
		goto L57
	case 42:
		goto L51
	case 43:
		goto L53
	case 46:
		goto L54
	case 55:
		v180 = int64(1)
		v181 = v143
		goto L44
	case 57:
		goto L50
	case 58:
		goto L58
	case 67:
		goto L47
	case 72:
		goto L60
	case 73:
		goto L49
	case 74:
		goto L55
	case 77:
		goto L59
	case 78:
		goto L56
	}
L43:
	;
	v203 = v137
	v204 = v183
	v209 = v181<<(uint(int64(32))%64) | v184
	goto L7
L44:
	;
	v183 = v137 + int32(1)
	v184 = v180 | v142
	v185 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v137)+1)))
	if base.B2i32(v185 == int32(42))|base.B2i32(base.Ui32((v185|int32(32)-int32(97))&int32(255)) < base.Ui32(int32(26))) != 0 {
		v137 = v183
		v139 = v185
		v142 = v184
		v143 = v181
		v144 = v180
		goto L42
	} else {
		goto L66
	}
L45:
	;
	v180 = v144
	v181 = v143 | v144
	goto L44
L46:
	;
	v162 = F_errsave_start(m, v17)
	mBase = m.M
	v163 = m.ExcPending
	if v163 != 0 {
		goto L1
	} else {
		goto L61
	}
L47:
	;
	v180 = int64(16384)
	v181 = v143
	goto L44
L48:
	;
	v180 = int64(8192)
	v181 = v143
	goto L44
L49:
	;
	v180 = int64(4096)
	v181 = v143
	goto L44
L50:
	;
	v180 = int64(2048)
	v181 = v143
	goto L44
L51:
	;
	v180 = int64(1024)
	v181 = v143
	goto L44
L52:
	;
	v180 = int64(512)
	v181 = v143
	goto L44
L53:
	;
	v180 = int64(256)
	v181 = v143
	goto L44
L54:
	;
	v180 = int64(128)
	v181 = v143
	goto L44
L55:
	;
	v180 = int64(64)
	v181 = v143
	goto L44
L56:
	;
	v180 = int64(32)
	v181 = v143
	goto L44
L57:
	;
	v180 = int64(16)
	v181 = v143
	goto L44
L58:
	;
	v180 = int64(8)
	v181 = v143
	goto L44
L59:
	;
	v180 = int64(4)
	v181 = v143
	goto L44
L60:
	;
	v180 = int64(2)
	v181 = v143
	goto L44
L61:
	;
	if v162 == int32(0) {
		goto L4
	} else {
		goto L62
	}
L62:
	;
	F_errcode(m, int32(33685634))
	mBase = m.M
	v168 = m.ExcPending
	if v168 != 0 {
		goto L1
	} else {
		goto L63
	}
L63:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v15))) = int32(_a_F_aclitemin_7)
	F_errmsg(m, int32(_a_F_aclitemin_8), v15)
	mBase = m.M
	v173 = m.ExcPending
	if v173 != 0 {
		goto L1
	} else {
		goto L64
	}
L64:
	;
	F_errsave_finish(m, v17, int32(_a_F_aclitemin_2), int32(369), int32(_a_F_aclitemin_3))
	mBase = m.M
	v178 = m.ExcPending
	if v178 != 0 {
		goto L1
	} else {
		goto L65
	}
L65:
	;
	goto L4
L66:
	;
	goto L43
L67:
	;
	v246 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v204))))
	if v246 == int32(47) {
		goto L79
	} else {
		goto L80
	}
L68:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v20))) = int32(0)
	goto L67
L69:
	;
	goto L70
L70:
	;
	v219 = v15 + int32(144)
	v220 = int32(0)
	v223 = F_GetSysCacheOid(m, int32(10), v219, v220, v220, v220)
	mBase = m.M
	v224 = m.ExcPending
	if v224 != 0 {
		goto L1
	} else {
		goto L71
	}
L71:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v20))) = v223
	if v223 != 0 {
		goto L67
	} else {
		goto L72
	}
L72:
	;
	v226 = F_errsave_start(m, v17)
	mBase = m.M
	v227 = m.ExcPending
	if v227 != 0 {
		goto L1
	} else {
		goto L73
	}
L73:
	;
	if v226 == int32(0) {
		goto L4
	} else {
		goto L74
	}
L74:
	;
	F_errcode(m, int32(67137668))
	mBase = m.M
	v232 = m.ExcPending
	if v232 != 0 {
		goto L1
	} else {
		goto L75
	}
L75:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v15)+48)) = v219
	F_errmsg(m, int32(_a_F_aclitemin_9), v15+int32(48))
	mBase = m.M
	v238 = m.ExcPending
	if v238 != 0 {
		goto L1
	} else {
		goto L76
	}
L76:
	;
	F_errsave_finish(m, v17, int32(_a_F_aclitemin_2), int32(383), int32(_a_F_aclitemin_3))
	mBase = m.M
	v243 = m.ExcPending
	if v243 != 0 {
		goto L1
	} else {
		goto L77
	}
L77:
	;
	goto L4
L78:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v20)+8)) = v209
	v334 = v327
	goto L104
L79:
	;
	v253 = F_getid(m, v203+int32(2), v15+int32(80), v17)
	mBase = m.M
	v254 = m.ExcPending
	if v254 != 0 {
		goto L1
	} else {
		goto L82
	}
L80:
	;
	goto L81
L81:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v20)+4)) = int32(10)
	v307 = F_errstart(m, int32(19), int32(0))
	mBase = m.M
	v308 = m.ExcPending
	if v308 != 0 {
		goto L1
	} else {
		goto L99
	}
L82:
	;
	if v253 == int32(0) {
		goto L4
	} else {
		goto L83
	}
L83:
	;
	v257 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v15)+80)))
	if v257 == int32(0) {
		goto L84
	} else {
		goto L85
	}
L84:
	;
	v260 = F_errsave_start(m, v17)
	mBase = m.M
	v261 = m.ExcPending
	if v261 != 0 {
		goto L1
	} else {
		goto L87
	}
L85:
	;
	goto L86
L86:
	;
	v278 = v15 + int32(80)
	v279 = int32(0)
	v282 = F_GetSysCacheOid(m, int32(10), v278, v279, v279, v279)
	mBase = m.M
	v283 = m.ExcPending
	if v283 != 0 {
		goto L1
	} else {
		goto L92
	}
L87:
	;
	if v260 == int32(0) {
		goto L4
	} else {
		goto L88
	}
L88:
	;
	F_errcode(m, int32(33685634))
	mBase = m.M
	v266 = m.ExcPending
	if v266 != 0 {
		goto L1
	} else {
		goto L89
	}
L89:
	;
	F_errmsg(m, int32(_a_F_aclitemin_10), int32(0))
	mBase = m.M
	v270 = m.ExcPending
	if v270 != 0 {
		goto L1
	} else {
		goto L90
	}
L90:
	;
	F_errsave_finish(m, v17, int32(_a_F_aclitemin_2), int32(398), int32(_a_F_aclitemin_3))
	mBase = m.M
	v275 = m.ExcPending
	if v275 != 0 {
		goto L1
	} else {
		goto L91
	}
L91:
	;
	goto L4
L92:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v20)+4)) = v282
	if v282 != 0 {
		v327 = v253
		goto L78
	} else {
		goto L93
	}
L93:
	;
	v285 = F_errsave_start(m, v17)
	mBase = m.M
	v286 = m.ExcPending
	if v286 != 0 {
		goto L1
	} else {
		goto L94
	}
L94:
	;
	if v285 == int32(0) {
		goto L4
	} else {
		goto L95
	}
L95:
	;
	F_errcode(m, int32(67137668))
	mBase = m.M
	v291 = m.ExcPending
	if v291 != 0 {
		goto L1
	} else {
		goto L96
	}
L96:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v15)+16)) = v278
	F_errmsg(m, int32(_a_F_aclitemin_9), v15+int32(16))
	mBase = m.M
	v297 = m.ExcPending
	if v297 != 0 {
		goto L1
	} else {
		goto L97
	}
L97:
	;
	F_errsave_finish(m, v17, int32(_a_F_aclitemin_2), int32(403), int32(_a_F_aclitemin_3))
	mBase = m.M
	v302 = m.ExcPending
	if v302 != 0 {
		goto L1
	} else {
		goto L98
	}
L98:
	;
	goto L4
L99:
	;
	if v307 == int32(0) {
		v327 = v204
		goto L78
	} else {
		goto L100
	}
L100:
	;
	F_errcode(m, int32(1792))
	mBase = m.M
	v313 = m.ExcPending
	if v313 != 0 {
		goto L1
	} else {
		goto L101
	}
L101:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v15)+32)) = int32(10)
	F_errmsg(m, int32(_a_F_aclitemin_11), v15+int32(32))
	mBase = m.M
	v320 = m.ExcPending
	if v320 != 0 {
		goto L1
	} else {
		goto L102
	}
L102:
	;
	F_errfinish(m, int32(_a_F_aclitemin_2), int32(411), int32(_a_F_aclitemin_3))
	mBase = m.M
	v325 = m.ExcPending
	if v325 != 0 {
		goto L1
	} else {
		goto L103
	}
L103:
	;
	v327 = v204
	goto L78
L104:
	;
	v342 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v334))))
	if base.B2i32(base.Ui32(v342-int32(9)) < base.Ui32(int32(5)))|base.B2i32(v342 == int32(32)) != 0 {
		goto L106
	} else {
		goto L107
	}
L106:
	;
	v334 = v334 + int32(1)
	goto L104
L107:
	;
	if v342 == int32(0) {
		v391 = v20
		goto L3
	} else {
		goto L109
	}
L109:
	;
	v354 = int32(0)
	v355 = F_errsave_start(m, v17)
	mBase = m.M
	v356 = m.ExcPending
	if v356 != 0 {
		goto L1
	} else {
		goto L110
	}
L110:
	;
	if v355 == int32(0) {
		v391 = v354
		goto L3
	} else {
		goto L111
	}
L111:
	;
	F_errcode(m, int32(33685634))
	mBase = m.M
	v361 = m.ExcPending
	if v361 != 0 {
		goto L1
	} else {
		goto L112
	}
L112:
	;
	F_errmsg(m, int32(_a_F_aclitemin_12), int32(0))
	mBase = m.M
	v365 = m.ExcPending
	if v365 != 0 {
		goto L1
	} else {
		goto L113
	}
L113:
	;
	F_errsave_finish(m, v17, int32(_a_F_aclitemin_2), int32(633), int32(_a_F_aclitemin_13))
	mBase = m.M
	v370 = m.ExcPending
	if v370 != 0 {
		goto L1
	} else {
		goto L114
	}
L114:
	;
	v391 = v354
	goto L3
}
func F_acos(m *base.Module, l0 float64) float64 {
	var v5 int64
	_ = v5
	var v10 int32
	_ = v10
	var v23 float64
	_ = v23
	var v35 float64
	_ = v35
	var v76 float64
	_ = v76
	var v79 float64
	_ = v79
	var v80 float64
	_ = v80
	var v116 float64
	_ = v116
	var v119 float64
	_ = v119
	var v122 float64
	_ = v122
	var v123 float64
	_ = v123
	var v159 float64
	_ = v159
	var v165 float64
	_ = v165
	var v169 float64
	_ = v169
	v5 = base.I64_reinterpret_f64(l0)
	v10 = base.I32_wrap_i64(int64(base.Ui64(v5)>>(uint(int64(32))%64))) & int32(2147483647)
	if base.Ui32(int32(1072693248)) <= base.Ui32(v10) {
		if base.I32_wrap_i64(v5)|(v10-int32(1072693248)) == int32(0) {
			if int64(0) <= v5 {
				v23 = float64(0)
			} else {
				v23 = float64(3.141592653589793)
			}
			return v23
		} else {
			return base.F64_div(float64(0), base.F64_sub(l0, l0))
		}
	} else {
		if base.Ui32(v10) <= base.Ui32(int32(1071644671)) {
			if base.Ui32(v10) < base.Ui32(int32(1012924417)) {
				v169 = float64(1.5707963267948966)
				return v169
			} else {
				v35 = base.F64_mul(l0, l0)
				return base.F64_add(base.F64_sub(base.F64_sub(float64(6.123233995736766e-17), base.F64_mul(l0, base.F64_div(base.F64_mul(v35, base.F64_add(base.F64_mul(v35, base.F64_add(base.F64_mul(v35, base.F64_add(base.F64_mul(v35, base.F64_add(base.F64_mul(v35, base.F64_add(base.F64_mul(v35, float64(3.479331075960212e-05)), float64(0.0007915349942898145))), float64(-0.04005553450067941))), float64(0.20121253213486293))), float64(-0.3255658186224009))), float64(0.16666666666666666))), base.F64_add(base.F64_mul(v35, base.F64_add(base.F64_mul(v35, base.F64_add(base.F64_mul(v35, base.F64_add(base.F64_mul(v35, float64(0.07703815055590194)), float64(-0.6882839716054533))), float64(2.0209457602335057))), float64(-2.403394911734414))), float64(1))))), l0), float64(1.5707963267948966))
			}
		} else {
			if v5 < int64(0) {
				v76 = float64(1)
				v79 = base.F64_mul(base.F64_add(l0, v76), float64(0.5))
				v80 = base.F64_sqrt(v79)
				v116 = base.F64_sub(float64(1.5707963267948966), base.F64_add(v80, base.F64_add(base.F64_mul(v80, base.F64_div(base.F64_mul(v79, base.F64_add(base.F64_mul(v79, base.F64_add(base.F64_mul(v79, base.F64_add(base.F64_mul(v79, base.F64_add(base.F64_mul(v79, base.F64_add(base.F64_mul(v79, float64(3.479331075960212e-05)), float64(0.0007915349942898145))), float64(-0.04005553450067941))), float64(0.20121253213486293))), float64(-0.3255658186224009))), float64(0.16666666666666666))), base.F64_add(base.F64_mul(v79, base.F64_add(base.F64_mul(v79, base.F64_add(base.F64_mul(v79, base.F64_add(base.F64_mul(v79, float64(0.07703815055590194)), float64(-0.6882839716054533))), float64(2.0209457602335057))), float64(-2.403394911734414))), v76))), float64(-6.123233995736766e-17))))
				return base.F64_add(v116, v116)
			} else {
				v119 = float64(1)
				v122 = base.F64_mul(base.F64_sub(v119, l0), float64(0.5))
				v123 = base.F64_sqrt(v122)
				v159 = base.F64_reinterpret_i64(base.I64_reinterpret_f64(v123) & int64(-4294967296))
				v165 = base.F64_add(base.F64_add(base.F64_mul(v123, base.F64_div(base.F64_mul(v122, base.F64_add(base.F64_mul(v122, base.F64_add(base.F64_mul(v122, base.F64_add(base.F64_mul(v122, base.F64_add(base.F64_mul(v122, base.F64_add(base.F64_mul(v122, float64(3.479331075960212e-05)), float64(0.0007915349942898145))), float64(-0.04005553450067941))), float64(0.20121253213486293))), float64(-0.3255658186224009))), float64(0.16666666666666666))), base.F64_add(base.F64_mul(v122, base.F64_add(base.F64_mul(v122, base.F64_add(base.F64_mul(v122, base.F64_add(base.F64_mul(v122, float64(0.07703815055590194)), float64(-0.6882839716054533))), float64(2.0209457602335057))), float64(-2.403394911734414))), v119))), base.F64_div(base.F64_sub(v122, base.F64_mul(v159, v159)), base.F64_add(v123, v159))), v159)
				v169 = base.F64_add(v165, v165)
				return v169
			}
		}
	}
}
func F_acquireLocksOnSubLinks(m *base.Module, l0 int32, l1 int32) int32 {
	mBase := m.M
	_ = mBase
	var v7 int32
	_ = v7
	var v10 int32
	_ = v10
	var v11 int32
	_ = v11
	var v16 int32
	_ = v16
	var v18 int32
	_ = v18
	var v19 int32
	_ = v19
	if l0 == int32(0) {
		return int32(0)
	} else {
		v7 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
		if v7 == int32(22) {
			v10 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
			v11 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l1))))
			F_AcquireRewriteLocks(m, v10, v11, int32(0))
			mBase = m.M
			v16 = m.ExcPending
			if v16 != 0 {
				return int32(0)
			} else {
				v18 = F_expression_tree_walker_impl(m, l0, int32(1041), l1)
				mBase = m.M
				v19 = m.ExcPending
				if v19 != 0 {
					return int32(0)
				} else {
					return v18
				}
			}
		} else {
			v18 = F_expression_tree_walker_impl(m, l0, int32(1041), l1)
			mBase = m.M
			v19 = m.ExcPending
			if v19 != 0 {
				return int32(0)
			} else {
				return v18
			}
		}
	}
}
func F_addNSItemToQuery(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32) {
	mBase := m.M
	_ = mBase
	var v4 int32
	_ = v4
	var v5 int32
	_ = v5
	var v7 int32
	_ = v7
	var v8 int32
	_ = v8
	var v11 int32
	_ = v11
	var v13 int32
	_ = v13
	var v14 int32
	_ = v14
	var v15 int32
	_ = v15
	var v19 int32
	_ = v19
	var v23 int32
	_ = v23
	var v24 int32
	_ = v24
	var v25 int32
	_ = v25
	v4 = l3
	v5 = l4
	if l2 != 0 {
		v7 = F_palloc0(m, int32(8))
		mBase = m.M
		v8 = m.ExcPending
		if v8 != 0 {
			return
		} else {
			*(*int32)(unsafe.Add(mBase, uint32(v7))) = int32(63)
			v11 = *(*int32)(unsafe.Add(mBase, uint32(l1)+8))
			*(*int32)(unsafe.Add(mBase, uint32(v7)+4)) = v11
			v13 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
			v14 = F_lappend(m, v13, v7)
			mBase = m.M
			v15 = m.ExcPending
			if v15 != 0 {
				return
			} else {
				*(*int32)(unsafe.Add(mBase, uint32(l0)+24)) = v14
				if v4|v5 != 0 {
					v19 = int32(256)
					*(*uint16)(unsafe.Add(mBase, uint32(l1)+22)) = uint16(v19)
					*(*uint8)(unsafe.Add(mBase, uint32(l1)+21)) = uint8(v5)
					*(*uint8)(unsafe.Add(mBase, uint32(l1)+20)) = uint8(v4)
					v23 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
					v24 = F_lappend(m, v23, l1)
					mBase = m.M
					v25 = m.ExcPending
					if v25 != 0 {
						return
					} else {
						*(*int32)(unsafe.Add(mBase, uint32(l0)+28)) = v24
						return
					}
				} else {
					return
				}
			}
		}
	} else {
		if v4|v5 != 0 {
			v19 = int32(256)
			*(*uint16)(unsafe.Add(mBase, uint32(l1)+22)) = uint16(v19)
			*(*uint8)(unsafe.Add(mBase, uint32(l1)+21)) = uint8(v5)
			*(*uint8)(unsafe.Add(mBase, uint32(l1)+20)) = uint8(v4)
			v23 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
			v24 = F_lappend(m, v23, l1)
			mBase = m.M
			v25 = m.ExcPending
			if v25 != 0 {
				return
			} else {
				*(*int32)(unsafe.Add(mBase, uint32(l0)+28)) = v24
				return
			}
		} else {
			return
		}
	}
}
func F_add_exact_object_address(m *base.Module, l0 int32, l1 int32) {
	mBase := m.M
	_ = mBase
	var v6 int32
	_ = v6
	var v7 int32
	_ = v7
	var v8 int32
	_ = v8
	var v15 int32
	_ = v15
	var v16 int32
	_ = v16
	var v18 int32
	_ = v18
	var v19 int32
	_ = v19
	var v20 int32
	_ = v20
	var v23 int32
	_ = v23
	var v24 int32
	_ = v24
	var v26 int64
	_ = v26
	var v28 int32
	_ = v28
	v6 = *(*int32)(unsafe.Add(mBase, uint32(l1)))
	v7 = *(*int32)(unsafe.Add(mBase, uint32(l1)+8))
	v8 = *(*int32)(unsafe.Add(mBase, uint32(l1)+12))
	if v8 <= v7 {
		*(*int32)(unsafe.Add(mBase, uint32(l1)+12)) = v8 << (uint(int32(1)) % 32)
		v15 = F_repalloc(m, v6, v8*int32(24))
		mBase = m.M
		v16 = m.ExcPending
		if v16 != 0 {
			return
		} else {
			*(*int32)(unsafe.Add(mBase, uint32(l1))) = v15
			v18 = *(*int32)(unsafe.Add(mBase, uint32(l1)+8))
			v19 = v15
			v20 = v18
			v23 = v20*int32(12) + v19
			v24 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
			*(*int32)(unsafe.Add(mBase, uint32(v23)+8)) = v24
			v26 = *(*int64)(unsafe.Add(mBase, uint32(l0)))
			*(*int64)(unsafe.Add(mBase, uint32(v23))) = v26
			v28 = *(*int32)(unsafe.Add(mBase, uint32(l1)+8))
			*(*int32)(unsafe.Add(mBase, uint32(l1)+8)) = v28 + int32(1)
			return
		}
	} else {
		v19 = v6
		v20 = v7
		v23 = v20*int32(12) + v19
		v24 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
		*(*int32)(unsafe.Add(mBase, uint32(v23)+8)) = v24
		v26 = *(*int64)(unsafe.Add(mBase, uint32(l0)))
		*(*int64)(unsafe.Add(mBase, uint32(v23))) = v26
		v28 = *(*int32)(unsafe.Add(mBase, uint32(l1)+8))
		*(*int32)(unsafe.Add(mBase, uint32(l1)+8)) = v28 + int32(1)
		return
	}
}
func F_add_non_redundant_clauses(m *base.Module, l0 int32, l1 int32, l2 int32) {
	mBase := m.M
	_ = mBase
	var v4 int32
	_ = v4
	var v12 int32
	_ = v12
	var v23 int32
	_ = v23
	var v24 int32
	_ = v24
	var v28 int32
	_ = v28
	var v29 int32
	_ = v29
	var v32 int32
	_ = v32
	var v33 int32
	_ = v33
	var v42 int32
	_ = v42
	var v45 int32
	_ = v45
	var v49 int32
	_ = v49
	var v50 int32
	_ = v50
	var v51 int32
	_ = v51
	var v52 int32
	_ = v52
	var v66 int32
	_ = v66
	var v67 int32
	_ = v67
	var v69 int32
	_ = v69
	var v72 int32
	_ = v72
	var v73 int32
	_ = v73
	var v78 int32
	_ = v78
	var v86 int32
	_ = v86
	var v88 int32
	_ = v88
	var v90 int32
	_ = v90
	var v91 int32
	_ = v91
	var v94 int32
	_ = v94
	var v98 int32
	_ = v98
	var v104 int32
	_ = v104
	var v105 int32
	_ = v105
	var v107 int32
	_ = v107
	var v108 int32
	_ = v108
	var v110 int32
	_ = v110
	var v111 int32
	_ = v111
	var v115 int32
	_ = v115
	var v116 int32
	_ = v116
	var v128 int32
	_ = v128
	var v139 int32
	_ = v139
	var v140 int32
	_ = v140
	v4 = int32(0)
	if l1 == v4 {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	return
L2:
	;
	v12 = *(*int32)(unsafe.Add(mBase, uint32(l1)+4))
	if v12 <= int32(0) {
		goto L1
	} else {
		goto L3
	}
L3:
	;
	v23 = v4
	goto L4
L4:
	;
	v24 = *(*int32)(unsafe.Add(mBase, uint32(l1)+12))
	v28 = *(*int32)(unsafe.Add(mBase, uint32(v24+v23<<(uint(int32(2))%32))))
	v29 = *(*int32)(unsafe.Add(mBase, uint32(l2)))
	if v29 == int32(0) {
		goto L7
	} else {
		goto L8
	}
L5:
	;
	goto L1
L6:
	;
	v139 = v23 + int32(1)
	v140 = *(*int32)(unsafe.Add(mBase, uint32(l1)+4))
	if v139 < v140 {
		v23 = v139
		goto L4
	} else {
		goto L36
	}
L7:
	;
	F_distribute_restrictinfo_to_rels(m, l0, v28)
	mBase = m.M
	v128 = m.ExcPending
	if v128 != 0 {
		goto L31
	} else {
		goto L35
	}
L8:
	;
	v32 = int32(0)
	v33 = *(*int32)(unsafe.Add(mBase, uint32(v29)+4))
	if v33 <= v32 {
		goto L7
	} else {
		goto L9
	}
L9:
	;
	v42 = v32
	goto L10
L10:
	;
	v45 = *(*int32)(unsafe.Add(mBase, uint32(v29)+12))
	v49 = *(*int32)(unsafe.Add(mBase, uint32(v45+v42<<(uint(int32(2))%32))))
	v50 = *(*int32)(unsafe.Add(mBase, uint32(v49)+28))
	v51 = *(*int32)(unsafe.Add(mBase, uint32(v28)+28))
	v52 = int32(0)
	if base.B2i32(v50 == v52)|base.B2i32(v51 == v52) != 0 {
		v98 = base.B2i32(v50|v51 == v52)
		goto L13
	} else {
		goto L14
	}
L11:
	;
	goto L7
L12:
	;
	if v98 != 0 {
		goto L23
	} else {
		goto L24
	}
L13:
	;
	goto L12
L14:
	;
	v66 = *(*int32)(unsafe.Add(mBase, uint32(v50)+4))
	v67 = *(*int32)(unsafe.Add(mBase, uint32(v51)+4))
	if v66 != v67 {
		v98 = int32(0)
		goto L13
	} else {
		goto L15
	}
L15:
	;
	v69 = int32(1)
	if v66 <= v69 {
		goto L16
	} else {
		goto L17
	}
L16:
	;
	v72 = v69
	goto L18
L17:
	;
	v72 = v66
	goto L18
L18:
	;
	v73 = int32(8)
	v78 = int32(0)
	goto L19
L19:
	;
	v86 = v78 << (uint(int32(2)) % 32)
	v88 = *(*int32)(unsafe.Add(mBase, uint32(v50+v73+v86)))
	v90 = *(*int32)(unsafe.Add(mBase, uint32(v51+v73+v86)))
	v91 = base.B2i32(v88 == v90)
	if v88 != v90 {
		v98 = v91
		goto L13
	} else {
		goto L21
	}
L20:
	;
	v98 = v91
	goto L13
L21:
	;
	v94 = v78 + int32(1)
	if v94 != v72 {
		v78 = v94
		goto L19
	} else {
		goto L22
	}
L22:
	;
	goto L20
L23:
	;
	if v28 == v49 {
		goto L6
	} else {
		goto L26
	}
L24:
	;
	goto L25
L25:
	;
	v115 = v42 + int32(1)
	v116 = *(*int32)(unsafe.Add(mBase, uint32(v29)+4))
	if v115 < v116 {
		v42 = v115
		goto L10
	} else {
		goto L34
	}
L26:
	;
	v104 = *(*int32)(unsafe.Add(mBase, uint32(v28)+60))
	if v104 != 0 {
		goto L27
	} else {
		goto L28
	}
L27:
	;
	v105 = *(*int32)(unsafe.Add(mBase, uint32(v49)+60))
	if v105 == v104 {
		goto L6
	} else {
		goto L30
	}
L28:
	;
	goto L29
L29:
	;
	v107 = *(*int32)(unsafe.Add(mBase, uint32(v28)+56))
	v108 = *(*int32)(unsafe.Add(mBase, uint32(v49)+56))
	*(*int32)(unsafe.Add(mBase, uint32(v28)+56)) = v108
	v110 = F_equal(m, v28, v49)
	mBase = m.M
	v111 = m.ExcPending
	if v111 != 0 {
		goto L31
	} else {
		goto L32
	}
L30:
	;
	goto L29
L31:
	;
	return
L32:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v28)+56)) = v107
	if v110 != 0 {
		goto L6
	} else {
		goto L33
	}
L33:
	;
	goto L25
L34:
	;
	goto L11
L35:
	;
	goto L6
L36:
	;
	goto L5
}
func F_add_nulling_relids_mutator(m *base.Module, l0 int32, l1 int32) int32 {
	mBase := m.M
	_ = mBase
	var v8 int32
	_ = v8
	var v15 int32
	_ = v15
	var v16 int32
	_ = v16
	var v18 int32
	_ = v18
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
	var v28 int32
	_ = v28
	var v29 int32
	_ = v29
	var v30 int32
	_ = v30
	var v31 int32
	_ = v31
	var v34 int32
	_ = v34
	var v35 int32
	_ = v35
	var v37 int32
	_ = v37
	var v38 int32
	_ = v38
	var v39 int32
	_ = v39
	var v49 int32
	_ = v49
	var v50 int32
	_ = v50
	var v52 int32
	_ = v52
	var v55 int32
	_ = v55
	var v56 int32
	_ = v56
	var v61 int32
	_ = v61
	var v68 int32
	_ = v68
	var v70 int32
	_ = v70
	var v72 int32
	_ = v72
	var v73 int32
	_ = v73
	var v75 int32
	_ = v75
	var v77 int32
	_ = v77
	var v84 int32
	_ = v84
	var v87 int32
	_ = v87
	var v88 int32
	_ = v88
	var v89 int32
	_ = v89
	var v90 int32
	_ = v90
	var v92 int32
	_ = v92
	var v93 int32
	_ = v93
	var v96 int64
	_ = v96
	var v98 int64
	_ = v98
	var v100 int64
	_ = v100
	var v104 int32
	_ = v104
	var v110 int32
	_ = v110
	var v111 int32
	_ = v111
	var v112 int32
	_ = v112
	var v119 int32
	_ = v119
	var v120 int32
	_ = v120
	if l0 == int32(0) {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	return int32(0)
L2:
	;
	goto L3
L3:
	;
	v8 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	if v8 != int32(319) {
		goto L6
	} else {
		goto L7
	}
L4:
	;
	v119 = F_expression_tree_mutator_impl(m, l0, int32(1052), l1)
	mBase = m.M
	v120 = m.ExcPending
	if v120 != 0 {
		goto L15
	} else {
		goto L41
	}
L5:
	;
	v104 = *(*int32)(unsafe.Add(mBase, uint32(l1)+8))
	*(*int32)(unsafe.Add(mBase, uint32(l1)+8)) = v104 + int32(1)
	v110 = F_query_tree_mutator_impl(m, l0, int32(1052), l1, int32(0))
	mBase = m.M
	v111 = m.ExcPending
	if v111 != 0 {
		goto L15
	} else {
		goto L40
	}
L6:
	;
	if v8 == int32(67) {
		goto L5
	} else {
		goto L9
	}
L7:
	;
	goto L8
L8:
	;
	v34 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v35 = *(*int32)(unsafe.Add(mBase, uint32(l1)+8))
	if v34 != v35 {
		goto L4
	} else {
		goto L20
	}
L9:
	;
	if v8 != int32(6) {
		goto L4
	} else {
		goto L10
	}
L10:
	;
	v15 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
	v16 = *(*int32)(unsafe.Add(mBase, uint32(l1)+8))
	if v15 != v16 {
		goto L4
	} else {
		goto L11
	}
L11:
	;
	v18 = *(*int32)(unsafe.Add(mBase, uint32(l1)))
	if v18 != 0 {
		goto L12
	} else {
		goto L13
	}
L12:
	;
	v19 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v20 = F_bms_is_member(m, v19, v18)
	mBase = m.M
	v23 = m.ExcPending
	if v23 != 0 {
		goto L15
	} else {
		goto L16
	}
L13:
	;
	goto L14
L14:
	;
	v26 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
	v27 = *(*int32)(unsafe.Add(mBase, uint32(l1)+4))
	v28 = F_bms_union(m, v26, v27)
	mBase = m.M
	v29 = m.ExcPending
	if v29 != 0 {
		goto L15
	} else {
		goto L18
	}
L15:
	;
	return int32(0)
L16:
	;
	if v20 == int32(0) {
		goto L4
	} else {
		goto L17
	}
L17:
	;
	goto L14
L18:
	;
	v30 = F_copyObjectImpl(m, l0)
	mBase = m.M
	v31 = m.ExcPending
	if v31 != 0 {
		goto L15
	} else {
		goto L19
	}
L19:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v30)+24)) = v28
	return v30
L20:
	;
	v37 = *(*int32)(unsafe.Add(mBase, uint32(l1)))
	if v37 != 0 {
		goto L21
	} else {
		goto L22
	}
L21:
	;
	v38 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	v39 = int32(0)
	if base.B2i32(v38 == v39)|base.B2i32(v37 == v39) != 0 {
		v84 = v39
		goto L25
	} else {
		goto L26
	}
L22:
	;
	goto L23
L23:
	;
	v87 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	v88 = *(*int32)(unsafe.Add(mBase, uint32(l1)+4))
	v89 = F_bms_union(m, v87, v88)
	mBase = m.M
	v90 = m.ExcPending
	if v90 != 0 {
		goto L15
	} else {
		goto L38
	}
L24:
	;
	if v84 == int32(0) {
		goto L4
	} else {
		goto L37
	}
L25:
	;
	goto L24
L26:
	;
	v49 = *(*int32)(unsafe.Add(mBase, uint32(v38)+4))
	v50 = *(*int32)(unsafe.Add(mBase, uint32(v37)+4))
	if v49 < v50 {
		goto L27
	} else {
		goto L28
	}
L27:
	;
	v52 = v49
	goto L29
L28:
	;
	v52 = v50
	goto L29
L29:
	;
	if v52 <= int32(1) {
		goto L30
	} else {
		goto L31
	}
L30:
	;
	v55 = int32(1)
	goto L32
L31:
	;
	v55 = v52
	goto L32
L32:
	;
	v56 = int32(8)
	v61 = int32(0)
	goto L33
L33:
	;
	v68 = v61 << (uint(int32(2)) % 32)
	v70 = *(*int32)(unsafe.Add(mBase, uint32(v37+v56+v68)))
	v72 = *(*int32)(unsafe.Add(mBase, uint32(v38+v56+v68)))
	v73 = v70 & v72
	v75 = base.B2i32(v73 != int32(0))
	if v73 != 0 {
		v84 = v75
		goto L25
	} else {
		goto L35
	}
L34:
	;
	v84 = v75
	goto L25
L35:
	;
	v77 = v61 + int32(1)
	if v77 != v55 {
		v61 = v77
		goto L33
	} else {
		goto L36
	}
L36:
	;
	goto L34
L37:
	;
	goto L23
L38:
	;
	v92 = F_palloc0(m, int32(24))
	mBase = m.M
	v93 = m.ExcPending
	if v93 != 0 {
		goto L15
	} else {
		goto L39
	}
L39:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v92))) = int32(319)
	v96 = *(*int64)(unsafe.Add(mBase, uint32(l0)+8))
	*(*int64)(unsafe.Add(mBase, uint32(v92)+8)) = v96
	v98 = *(*int64)(unsafe.Add(mBase, uint32(l0)+16))
	*(*int64)(unsafe.Add(mBase, uint32(v92)+16)) = v98
	v100 = *(*int64)(unsafe.Add(mBase, uint32(l0)))
	*(*int64)(unsafe.Add(mBase, uint32(v92))) = v100
	*(*int32)(unsafe.Add(mBase, uint32(v92)+12)) = v89
	return v92
L40:
	;
	v112 = *(*int32)(unsafe.Add(mBase, uint32(l1)+8))
	*(*int32)(unsafe.Add(mBase, uint32(l1)+8)) = v112 - int32(1)
	return v110
L41:
	;
	return v119
}
func F_add_paths_to_joinrel(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32, l5 int32, l6 int32) {
	mBase := m.M
	_ = mBase
	var v8 int32
	_ = v8
	var v37 int32
	_ = v37
	var v39 int32
	_ = v39
	var v43 int32
	_ = v43
	var v46 int32
	_ = v46
	var v48 int32
	_ = v48
	var v49 int32
	_ = v49
	var v57 int32
	_ = v57
	var v58 int32
	_ = v58
	var v59 int32
	_ = v59
	var v68 int32
	_ = v68
	var v69 int32
	_ = v69
	var v71 int32
	_ = v71
	var v74 int32
	_ = v74
	var v75 int32
	_ = v75
	var v80 int32
	_ = v80
	var v87 int32
	_ = v87
	var v89 int32
	_ = v89
	var v91 int32
	_ = v91
	var v94 int32
	_ = v94
	var v96 int32
	_ = v96
	var v98 int32
	_ = v98
	var v105 int32
	_ = v105
	var v112 int32
	_ = v112
	var v113 int32
	_ = v113
	var v114 int32
	_ = v114
	var v116 int32
	_ = v116
	var v117 int32
	_ = v117
	var v118 int32
	_ = v118
	var v119 int32
	_ = v119
	var v120 int32
	_ = v120
	var v121 int32
	_ = v121
	var v122 int32
	_ = v122
	var v126 int32
	_ = v126
	var v128 int32
	_ = v128
	var v137 int32
	_ = v137
	var v140 int32
	_ = v140
	var v156 int32
	_ = v156
	var v158 int32
	_ = v158
	var v166 int32
	_ = v166
	var v183 int32
	_ = v183
	var v187 int32
	_ = v187
	var v188 int32
	_ = v188
	var v189 int32
	_ = v189
	var v190 int32
	_ = v190
	var v191 int32
	_ = v191
	var v200 int32
	_ = v200
	var v201 int32
	_ = v201
	var v203 int32
	_ = v203
	var v206 int32
	_ = v206
	var v207 int32
	_ = v207
	var v212 int32
	_ = v212
	var v219 int32
	_ = v219
	var v221 int32
	_ = v221
	var v223 int32
	_ = v223
	var v226 int32
	_ = v226
	var v228 int32
	_ = v228
	var v230 int32
	_ = v230
	var v237 int32
	_ = v237
	var v244 int32
	_ = v244
	var v247 int32
	_ = v247
	var v250 int32
	_ = v250
	var v251 int32
	_ = v251
	var v252 int32
	_ = v252
	var v256 int32
	_ = v256
	var v257 int32
	_ = v257
	var v258 int32
	_ = v258
	var v259 int32
	_ = v259
	var v268 int32
	_ = v268
	var v269 int32
	_ = v269
	var v271 int32
	_ = v271
	var v274 int32
	_ = v274
	var v275 int32
	_ = v275
	var v280 int32
	_ = v280
	var v287 int32
	_ = v287
	var v289 int32
	_ = v289
	var v291 int32
	_ = v291
	var v294 int32
	_ = v294
	var v296 int32
	_ = v296
	var v298 int32
	_ = v298
	var v305 int32
	_ = v305
	var v312 int32
	_ = v312
	var v315 int32
	_ = v315
	var v316 int32
	_ = v316
	var v325 int32
	_ = v325
	var v326 int32
	_ = v326
	var v328 int32
	_ = v328
	var v331 int32
	_ = v331
	var v332 int32
	_ = v332
	var v337 int32
	_ = v337
	var v344 int32
	_ = v344
	var v346 int32
	_ = v346
	var v348 int32
	_ = v348
	var v351 int32
	_ = v351
	var v353 int32
	_ = v353
	var v355 int32
	_ = v355
	var v362 int32
	_ = v362
	var v369 int32
	_ = v369
	var v372 int32
	_ = v372
	var v374 int32
	_ = v374
	var v375 int32
	_ = v375
	var v384 int32
	_ = v384
	var v385 int32
	_ = v385
	var v387 int32
	_ = v387
	var v390 int32
	_ = v390
	var v391 int32
	_ = v391
	var v396 int32
	_ = v396
	var v403 int32
	_ = v403
	var v405 int32
	_ = v405
	var v407 int32
	_ = v407
	var v410 int32
	_ = v410
	var v412 int32
	_ = v412
	var v414 int32
	_ = v414
	var v421 int32
	_ = v421
	var v428 int32
	_ = v428
	var v432 int32
	_ = v432
	var v433 int32
	_ = v433
	var v442 int32
	_ = v442
	var v443 int32
	_ = v443
	var v445 int32
	_ = v445
	var v448 int32
	_ = v448
	var v449 int32
	_ = v449
	var v454 int32
	_ = v454
	var v461 int32
	_ = v461
	var v463 int32
	_ = v463
	var v465 int32
	_ = v465
	var v468 int32
	_ = v468
	var v470 int32
	_ = v470
	var v472 int32
	_ = v472
	var v479 int32
	_ = v479
	var v486 int32
	_ = v486
	var v490 int32
	_ = v490
	var v492 int32
	_ = v492
	var v493 int32
	_ = v493
	var v494 int32
	_ = v494
	var v495 int32
	_ = v495
	var v497 int32
	_ = v497
	var v498 int32
	_ = v498
	var v516 int32
	_ = v516
	var v536 int32
	_ = v536
	var v573 int32
	_ = v573
	var v574 int32
	_ = v574
	var v592 int32
	_ = v592
	var v612 int32
	_ = v612
	var v649 int32
	_ = v649
	var v650 int32
	_ = v650
	var v652 int32
	_ = v652
	var v653 int32
	_ = v653
	var v655 int32
	_ = v655
	var v656 int32
	_ = v656
	var v666 int32
	_ = v666
	var v676 int32
	_ = v676
	var v694 int32
	_ = v694
	var v695 int32
	_ = v695
	var v713 int32
	_ = v713
	var v718 int32
	_ = v718
	var v742 int32
	_ = v742
	var v758 int32
	_ = v758
	var v763 int32
	_ = v763
	var v795 int32
	_ = v795
	var v825 int32
	_ = v825
	var v826 int32
	_ = v826
	var v827 int32
	_ = v827
	var v829 int32
	_ = v829
	var v834 int32
	_ = v834
	var v839 int32
	_ = v839
	var v847 int32
	_ = v847
	var v850 int32
	_ = v850
	var v878 int32
	_ = v878
	var v882 int32
	_ = v882
	var v883 int32
	_ = v883
	var v884 int32
	_ = v884
	var v885 int32
	_ = v885
	var v886 int32
	_ = v886
	var v895 int32
	_ = v895
	var v896 int32
	_ = v896
	var v898 int32
	_ = v898
	var v901 int32
	_ = v901
	var v902 int32
	_ = v902
	var v907 int32
	_ = v907
	var v914 int32
	_ = v914
	var v916 int32
	_ = v916
	var v918 int32
	_ = v918
	var v921 int32
	_ = v921
	var v923 int32
	_ = v923
	var v925 int32
	_ = v925
	var v932 int32
	_ = v932
	var v939 int32
	_ = v939
	var v942 int32
	_ = v942
	var v943 int32
	_ = v943
	var v944 int32
	_ = v944
	var v946 int32
	_ = v946
	var v947 int32
	_ = v947
	var v954 int32
	_ = v954
	var v986 int32
	_ = v986
	var v990 int32
	_ = v990
	var v991 float64
	_ = v991
	var v992 int32
	_ = v992
	var v994 int32
	_ = v994
	var v995 int32
	_ = v995
	var v996 int32
	_ = v996
	var v997 int64
	_ = v997
	var v1013 int32
	_ = v1013
	var v1015 float64
	_ = v1015
	var v1016 int32
	_ = v1016
	var v1018 int32
	_ = v1018
	var v1021 float64
	_ = v1021
	var v1022 float64
	_ = v1022
	var v1024 float64
	_ = v1024
	var v1027 float64
	_ = v1027
	var v1030 float64
	_ = v1030
	var v1072 int32
	_ = v1072
	var v1075 int32
	_ = v1075
	var v1085 int32
	_ = v1085
	var v1115 int32
	_ = v1115
	var v1119 int32
	_ = v1119
	var v1120 int32
	_ = v1120
	var v1121 int32
	_ = v1121
	var v1131 int32
	_ = v1131
	var v1132 int32
	_ = v1132
	var v1134 int32
	_ = v1134
	var v1137 int32
	_ = v1137
	var v1138 int32
	_ = v1138
	var v1143 int32
	_ = v1143
	var v1150 int32
	_ = v1150
	var v1152 int32
	_ = v1152
	var v1154 int32
	_ = v1154
	var v1155 int32
	_ = v1155
	var v1157 int32
	_ = v1157
	var v1159 int32
	_ = v1159
	var v1166 int32
	_ = v1166
	var v1169 int32
	_ = v1169
	var v1170 int32
	_ = v1170
	var v1180 int32
	_ = v1180
	var v1181 int32
	_ = v1181
	var v1183 int32
	_ = v1183
	var v1186 int32
	_ = v1186
	var v1187 int32
	_ = v1187
	var v1192 int32
	_ = v1192
	var v1199 int32
	_ = v1199
	var v1201 int32
	_ = v1201
	var v1203 int32
	_ = v1203
	var v1204 int32
	_ = v1204
	var v1206 int32
	_ = v1206
	var v1208 int32
	_ = v1208
	var v1215 int32
	_ = v1215
	var v1216 int32
	_ = v1216
	var v1217 int32
	_ = v1217
	var v1218 int32
	_ = v1218
	var v1219 int32
	_ = v1219
	var v1220 int32
	_ = v1220
	var v1221 int32
	_ = v1221
	var v1222 int32
	_ = v1222
	var v1224 int32
	_ = v1224
	var v1227 int32
	_ = v1227
	var v1228 int32
	_ = v1228
	var v1238 int32
	_ = v1238
	var v1239 int32
	_ = v1239
	var v1241 int32
	_ = v1241
	var v1244 int32
	_ = v1244
	var v1245 int32
	_ = v1245
	var v1250 int32
	_ = v1250
	var v1257 int32
	_ = v1257
	var v1259 int32
	_ = v1259
	var v1261 int32
	_ = v1261
	var v1262 int32
	_ = v1262
	var v1264 int32
	_ = v1264
	var v1266 int32
	_ = v1266
	var v1273 int32
	_ = v1273
	var v1276 int32
	_ = v1276
	var v1277 int32
	_ = v1277
	var v1287 int32
	_ = v1287
	var v1288 int32
	_ = v1288
	var v1290 int32
	_ = v1290
	var v1293 int32
	_ = v1293
	var v1294 int32
	_ = v1294
	var v1299 int32
	_ = v1299
	var v1306 int32
	_ = v1306
	var v1308 int32
	_ = v1308
	var v1310 int32
	_ = v1310
	var v1311 int32
	_ = v1311
	var v1313 int32
	_ = v1313
	var v1315 int32
	_ = v1315
	var v1322 int32
	_ = v1322
	var v1323 int32
	_ = v1323
	var v1324 int32
	_ = v1324
	var v1325 int32
	_ = v1325
	var v1326 int32
	_ = v1326
	var v1327 int32
	_ = v1327
	var v1328 int32
	_ = v1328
	var v1329 int32
	_ = v1329
	var v1332 int32
	_ = v1332
	var v1333 int32
	_ = v1333
	var v1371 int32
	_ = v1371
	var v1372 int32
	_ = v1372
	var v1373 int32
	_ = v1373
	var v1374 int32
	_ = v1374
	var v1380 int32
	_ = v1380
	var v1383 int32
	_ = v1383
	var v1384 int32
	_ = v1384
	var v1385 int32
	_ = v1385
	var v1388 int32
	_ = v1388
	var v1389 int32
	_ = v1389
	var v1390 int32
	_ = v1390
	var v1400 int32
	_ = v1400
	var v1401 int32
	_ = v1401
	var v1403 int32
	_ = v1403
	var v1406 int32
	_ = v1406
	var v1407 int32
	_ = v1407
	var v1412 int32
	_ = v1412
	var v1419 int32
	_ = v1419
	var v1421 int32
	_ = v1421
	var v1423 int32
	_ = v1423
	var v1424 int32
	_ = v1424
	var v1426 int32
	_ = v1426
	var v1428 int32
	_ = v1428
	var v1435 int32
	_ = v1435
	var v1436 int32
	_ = v1436
	var v1439 int32
	_ = v1439
	var v1440 int32
	_ = v1440
	var v1441 int32
	_ = v1441
	var v1451 int32
	_ = v1451
	var v1452 int32
	_ = v1452
	var v1454 int32
	_ = v1454
	var v1457 int32
	_ = v1457
	var v1458 int32
	_ = v1458
	var v1463 int32
	_ = v1463
	var v1470 int32
	_ = v1470
	var v1472 int32
	_ = v1472
	var v1474 int32
	_ = v1474
	var v1475 int32
	_ = v1475
	var v1477 int32
	_ = v1477
	var v1479 int32
	_ = v1479
	var v1486 int32
	_ = v1486
	var v1488 int32
	_ = v1488
	var v1491 int32
	_ = v1491
	var v1492 int32
	_ = v1492
	var v1493 int32
	_ = v1493
	var v1503 int32
	_ = v1503
	var v1504 int32
	_ = v1504
	var v1506 int32
	_ = v1506
	var v1509 int32
	_ = v1509
	var v1510 int32
	_ = v1510
	var v1515 int32
	_ = v1515
	var v1522 int32
	_ = v1522
	var v1524 int32
	_ = v1524
	var v1526 int32
	_ = v1526
	var v1527 int32
	_ = v1527
	var v1529 int32
	_ = v1529
	var v1531 int32
	_ = v1531
	var v1538 int32
	_ = v1538
	var v1539 int32
	_ = v1539
	var v1542 int32
	_ = v1542
	var v1543 int32
	_ = v1543
	var v1544 int32
	_ = v1544
	var v1554 int32
	_ = v1554
	var v1555 int32
	_ = v1555
	var v1557 int32
	_ = v1557
	var v1560 int32
	_ = v1560
	var v1561 int32
	_ = v1561
	var v1566 int32
	_ = v1566
	var v1573 int32
	_ = v1573
	var v1575 int32
	_ = v1575
	var v1577 int32
	_ = v1577
	var v1578 int32
	_ = v1578
	var v1580 int32
	_ = v1580
	var v1582 int32
	_ = v1582
	var v1589 int32
	_ = v1589
	var v1594 int32
	_ = v1594
	var v1595 int32
	_ = v1595
	var v1596 int32
	_ = v1596
	var v1598 int32
	_ = v1598
	var v1599 int32
	_ = v1599
	var v1600 int32
	_ = v1600
	var v1601 int32
	_ = v1601
	var v1602 int32
	_ = v1602
	var v1603 int32
	_ = v1603
	var v1604 int32
	_ = v1604
	var v1605 int32
	_ = v1605
	var v1609 int32
	_ = v1609
	var v1616 int32
	_ = v1616
	var v1621 int32
	_ = v1621
	var v1624 int32
	_ = v1624
	var v1625 int32
	_ = v1625
	var v1626 int32
	_ = v1626
	var v1627 int32
	_ = v1627
	var v1630 int32
	_ = v1630
	var v1635 int32
	_ = v1635
	var v1638 int32
	_ = v1638
	var v1641 int32
	_ = v1641
	var v1642 int32
	_ = v1642
	var v1644 int32
	_ = v1644
	var v1652 int32
	_ = v1652
	var v1653 int32
	_ = v1653
	var v1656 int32
	_ = v1656
	var v1659 int32
	_ = v1659
	var v1664 int32
	_ = v1664
	var v1674 int32
	_ = v1674
	var v1678 int32
	_ = v1678
	var v1679 int32
	_ = v1679
	var v1680 int32
	_ = v1680
	var v1681 int32
	_ = v1681
	var v1684 int32
	_ = v1684
	var v1688 int32
	_ = v1688
	var v1689 int32
	_ = v1689
	var v1690 int32
	_ = v1690
	var v1691 int32
	_ = v1691
	var v1692 int32
	_ = v1692
	var v1693 int32
	_ = v1693
	var v1702 int32
	_ = v1702
	var v1712 int32
	_ = v1712
	var v1732 int32
	_ = v1732
	var v1736 int32
	_ = v1736
	var v1737 int32
	_ = v1737
	var v1738 int32
	_ = v1738
	var v1747 int32
	_ = v1747
	var v1776 int32
	_ = v1776
	var v1813 int32
	_ = v1813
	var v1814 int32
	_ = v1814
	var v1823 int32
	_ = v1823
	var v1852 int32
	_ = v1852
	var v1891 int32
	_ = v1891
	var v1892 int32
	_ = v1892
	var v1894 int32
	_ = v1894
	var v1895 int32
	_ = v1895
	var v1906 int32
	_ = v1906
	var v1937 int32
	_ = v1937
	var v1940 int32
	_ = v1940
	var v1978 int32
	_ = v1978
	var v1982 int32
	_ = v1982
	var v1984 int32
	_ = v1984
	var v1994 int32
	_ = v1994
	var v1995 int32
	_ = v1995
	var v2023 int32
	_ = v2023
	var v2027 int32
	_ = v2027
	var v2028 int32
	_ = v2028
	var v2031 int32
	_ = v2031
	var v2032 int32
	_ = v2032
	var v2033 int32
	_ = v2033
	var v2043 int32
	_ = v2043
	var v2044 int32
	_ = v2044
	var v2046 int32
	_ = v2046
	var v2049 int32
	_ = v2049
	var v2050 int32
	_ = v2050
	var v2055 int32
	_ = v2055
	var v2062 int32
	_ = v2062
	var v2064 int32
	_ = v2064
	var v2066 int32
	_ = v2066
	var v2067 int32
	_ = v2067
	var v2069 int32
	_ = v2069
	var v2071 int32
	_ = v2071
	var v2078 int32
	_ = v2078
	var v2082 int32
	_ = v2082
	var v2084 int32
	_ = v2084
	var v2085 int32
	_ = v2085
	var v2094 int32
	_ = v2094
	var v2124 int32
	_ = v2124
	var v2137 int32
	_ = v2137
	var v2168 int32
	_ = v2168
	var v2169 int32
	_ = v2169
	var v2177 int32
	_ = v2177
	var v2207 int32
	_ = v2207
	var v2210 int32
	_ = v2210
	var v2213 int32
	_ = v2213
	var v2233 int32
	_ = v2233
	var v2256 int32
	_ = v2256
	var v2257 int32
	_ = v2257
	var v2267 int32
	_ = v2267
	var v2298 int32
	_ = v2298
	var v2301 int32
	_ = v2301
	var v2304 int32
	_ = v2304
	var v2342 int32
	_ = v2342
	var v2343 int32
	_ = v2343
	var v2344 int32
	_ = v2344
	var v2347 int32
	_ = v2347
	var v2350 int32
	_ = v2350
	var v2371 int32
	_ = v2371
	var v2389 int32
	_ = v2389
	var v2393 int32
	_ = v2393
	var v2394 int32
	_ = v2394
	var v2404 int32
	_ = v2404
	var v2433 int32
	_ = v2433
	var v2435 int32
	_ = v2435
	var v2441 int32
	_ = v2441
	var v2480 int32
	_ = v2480
	var v2481 int32
	_ = v2481
	var v2484 int32
	_ = v2484
	var v2485 int32
	_ = v2485
	var v2499 int32
	_ = v2499
	var v2523 int32
	_ = v2523
	var v2526 int32
	_ = v2526
	var v2527 int32
	_ = v2527
	var v2528 int32
	_ = v2528
	var v2547 int32
	_ = v2547
	var v2570 int32
	_ = v2570
	var v2572 int32
	_ = v2572
	var v2573 int32
	_ = v2573
	var v2585 int32
	_ = v2585
	var v2587 int32
	_ = v2587
	var v2593 int32
	_ = v2593
	var v2597 int32
	_ = v2597
	var v2616 int32
	_ = v2616
	var v2617 int32
	_ = v2617
	var v2620 int32
	_ = v2620
	var v2622 int32
	_ = v2622
	var v2626 int32
	_ = v2626
	var v2628 int32
	_ = v2628
	var v2632 int32
	_ = v2632
	var v2636 int32
	_ = v2636
	var v2637 int32
	_ = v2637
	var v2638 int32
	_ = v2638
	var v2639 int32
	_ = v2639
	var v2640 int32
	_ = v2640
	var v2641 int32
	_ = v2641
	var v2642 int32
	_ = v2642
	var v2643 int32
	_ = v2643
	var v2644 int32
	_ = v2644
	var v2645 int32
	_ = v2645
	var v2646 int32
	_ = v2646
	var v2647 int32
	_ = v2647
	var v2648 int32
	_ = v2648
	var v2649 int32
	_ = v2649
	var v2650 int32
	_ = v2650
	var v2652 int32
	_ = v2652
	var v2663 int32
	_ = v2663
	var v2664 int32
	_ = v2664
	var v2670 int32
	_ = v2670
	var v2698 int32
	_ = v2698
	var v2700 int32
	_ = v2700
	var v2706 int32
	_ = v2706
	var v2718 int32
	_ = v2718
	var v2731 int32
	_ = v2731
	var v2732 int32
	_ = v2732
	var v2733 int32
	_ = v2733
	var v2734 int32
	_ = v2734
	var v2735 int32
	_ = v2735
	var v2738 int32
	_ = v2738
	var v2746 int32
	_ = v2746
	var v2754 int32
	_ = v2754
	var v2779 int32
	_ = v2779
	var v2781 int32
	_ = v2781
	var v2785 int32
	_ = v2785
	var v2786 int32
	_ = v2786
	var v2787 int32
	_ = v2787
	var v2790 int32
	_ = v2790
	var v2791 int32
	_ = v2791
	var v2792 int32
	_ = v2792
	var v2793 int32
	_ = v2793
	var v2807 int32
	_ = v2807
	var v2831 int32
	_ = v2831
	var v2833 int32
	_ = v2833
	var v2847 int32
	_ = v2847
	var v2872 int32
	_ = v2872
	var v2875 int32
	_ = v2875
	var v2876 int32
	_ = v2876
	var v2877 int32
	_ = v2877
	var v2878 int32
	_ = v2878
	var v2879 int32
	_ = v2879
	var v2880 int32
	_ = v2880
	var v2881 int32
	_ = v2881
	var v2883 int32
	_ = v2883
	var v2886 int32
	_ = v2886
	var v2887 int32
	_ = v2887
	var v2891 int32
	_ = v2891
	var v2893 int32
	_ = v2893
	var v2894 int32
	_ = v2894
	var v2904 int32
	_ = v2904
	var v2934 int32
	_ = v2934
	var v2938 int32
	_ = v2938
	var v2939 int32
	_ = v2939
	var v2940 int32
	_ = v2940
	var v2941 int32
	_ = v2941
	var v2942 int32
	_ = v2942
	var v2943 int32
	_ = v2943
	var v2944 int32
	_ = v2944
	var v2945 int32
	_ = v2945
	var v2946 int32
	_ = v2946
	var v2947 int32
	_ = v2947
	var v2948 int32
	_ = v2948
	var v2949 int32
	_ = v2949
	var v2950 int32
	_ = v2950
	var v2951 int32
	_ = v2951
	var v2953 int32
	_ = v2953
	var v2956 int32
	_ = v2956
	var v2958 int32
	_ = v2958
	var v2960 int32
	_ = v2960
	var v2961 int32
	_ = v2961
	var v2999 int32
	_ = v2999
	var v3000 int32
	_ = v3000
	var v3005 int32
	_ = v3005
	var v3009 int32
	_ = v3009
	var v3014 int32
	_ = v3014
	var v3017 int32
	_ = v3017
	var v3018 int32
	_ = v3018
	var v3019 int32
	_ = v3019
	var v3020 int32
	_ = v3020
	var v3023 int32
	_ = v3023
	var v3024 int32
	_ = v3024
	var v3025 int32
	_ = v3025
	var v3035 int32
	_ = v3035
	var v3036 int32
	_ = v3036
	var v3038 int32
	_ = v3038
	var v3041 int32
	_ = v3041
	var v3042 int32
	_ = v3042
	var v3047 int32
	_ = v3047
	var v3054 int32
	_ = v3054
	var v3056 int32
	_ = v3056
	var v3058 int32
	_ = v3058
	var v3059 int32
	_ = v3059
	var v3061 int32
	_ = v3061
	var v3063 int32
	_ = v3063
	var v3070 int32
	_ = v3070
	var v3071 int32
	_ = v3071
	var v3074 int32
	_ = v3074
	var v3075 int32
	_ = v3075
	var v3076 int32
	_ = v3076
	var v3086 int32
	_ = v3086
	var v3087 int32
	_ = v3087
	var v3089 int32
	_ = v3089
	var v3092 int32
	_ = v3092
	var v3093 int32
	_ = v3093
	var v3098 int32
	_ = v3098
	var v3105 int32
	_ = v3105
	var v3107 int32
	_ = v3107
	var v3109 int32
	_ = v3109
	var v3110 int32
	_ = v3110
	var v3112 int32
	_ = v3112
	var v3114 int32
	_ = v3114
	var v3121 int32
	_ = v3121
	var v3126 int32
	_ = v3126
	var v3127 int32
	_ = v3127
	var v3128 int32
	_ = v3128
	var v3134 int32
	_ = v3134
	var v3135 int32
	_ = v3135
	var v3141 int32
	_ = v3141
	var v3145 int32
	_ = v3145
	var v3147 int32
	_ = v3147
	var v3153 int32
	_ = v3153
	var v3154 int32
	_ = v3154
	var v3156 int32
	_ = v3156
	var v3157 int32
	_ = v3157
	var v3158 int32
	_ = v3158
	var v3159 int32
	_ = v3159
	var v3162 int32
	_ = v3162
	var v3165 int32
	_ = v3165
	var v3167 int32
	_ = v3167
	var v3184 int32
	_ = v3184
	var v3209 int32
	_ = v3209
	var v3213 int32
	_ = v3213
	var v3214 int32
	_ = v3214
	var v3217 int32
	_ = v3217
	var v3218 int32
	_ = v3218
	var v3219 int32
	_ = v3219
	var v3229 int32
	_ = v3229
	var v3230 int32
	_ = v3230
	var v3232 int32
	_ = v3232
	var v3235 int32
	_ = v3235
	var v3236 int32
	_ = v3236
	var v3241 int32
	_ = v3241
	var v3248 int32
	_ = v3248
	var v3250 int32
	_ = v3250
	var v3252 int32
	_ = v3252
	var v3253 int32
	_ = v3253
	var v3255 int32
	_ = v3255
	var v3257 int32
	_ = v3257
	var v3264 int32
	_ = v3264
	var v3265 int32
	_ = v3265
	var v3268 int32
	_ = v3268
	var v3269 int32
	_ = v3269
	var v3270 int32
	_ = v3270
	var v3280 int32
	_ = v3280
	var v3281 int32
	_ = v3281
	var v3283 int32
	_ = v3283
	var v3286 int32
	_ = v3286
	var v3287 int32
	_ = v3287
	var v3292 int32
	_ = v3292
	var v3299 int32
	_ = v3299
	var v3301 int32
	_ = v3301
	var v3303 int32
	_ = v3303
	var v3304 int32
	_ = v3304
	var v3306 int32
	_ = v3306
	var v3308 int32
	_ = v3308
	var v3315 int32
	_ = v3315
	var v3319 int32
	_ = v3319
	var v3321 int32
	_ = v3321
	var v3322 int32
	_ = v3322
	var v3323 int32
	_ = v3323
	var v3324 int32
	_ = v3324
	var v3325 int32
	_ = v3325
	var v3326 int32
	_ = v3326
	var v3327 int32
	_ = v3327
	var v3332 int32
	_ = v3332
	var v3335 int32
	_ = v3335
	var v3336 int32
	_ = v3336
	var v3346 int32
	_ = v3346
	var v3375 int32
	_ = v3375
	var v3379 int32
	_ = v3379
	var v3381 int32
	_ = v3381
	var v3383 int32
	_ = v3383
	var v3384 int32
	_ = v3384
	var v3385 int32
	_ = v3385
	var v3387 int32
	_ = v3387
	var v3389 int32
	_ = v3389
	var v3390 int32
	_ = v3390
	var v3437 int32
	_ = v3437
	var v3469 int32
	_ = v3469
	var v3510 int32
	_ = v3510
	var v3548 int32
	_ = v3548
	var v3549 int32
	_ = v3549
	var v3587 int32
	_ = v3587
	var v3588 int32
	_ = v3588
	var v3597 int32
	_ = v3597
	var v3599 int32
	_ = v3599
	var v3602 int32
	_ = v3602
	var v3604 int32
	_ = v3604
	var v3605 int32
	_ = v3605
	var v3608 int32
	_ = v3608
	var v3610 int32
	_ = v3610
	var v3615 int32
	_ = v3615
	var v3616 int32
	_ = v3616
	var v3619 int32
	_ = v3619
	var v3622 int32
	_ = v3622
	var v3623 int32
	_ = v3623
	var v3624 int32
	_ = v3624
	var v3634 int32
	_ = v3634
	var v3635 int32
	_ = v3635
	var v3637 int32
	_ = v3637
	var v3640 int32
	_ = v3640
	var v3641 int32
	_ = v3641
	var v3646 int32
	_ = v3646
	var v3653 int32
	_ = v3653
	var v3655 int32
	_ = v3655
	var v3657 int32
	_ = v3657
	var v3658 int32
	_ = v3658
	var v3660 int32
	_ = v3660
	var v3662 int32
	_ = v3662
	var v3669 int32
	_ = v3669
	var v3670 int32
	_ = v3670
	var v3673 int32
	_ = v3673
	var v3674 int32
	_ = v3674
	var v3675 int32
	_ = v3675
	var v3685 int32
	_ = v3685
	var v3686 int32
	_ = v3686
	var v3688 int32
	_ = v3688
	var v3691 int32
	_ = v3691
	var v3692 int32
	_ = v3692
	var v3697 int32
	_ = v3697
	var v3704 int32
	_ = v3704
	var v3706 int32
	_ = v3706
	var v3708 int32
	_ = v3708
	var v3709 int32
	_ = v3709
	var v3711 int32
	_ = v3711
	var v3713 int32
	_ = v3713
	var v3720 int32
	_ = v3720
	var v3722 int32
	_ = v3722
	var v3724 int32
	_ = v3724
	var v3730 int32
	_ = v3730
	var v3731 int32
	_ = v3731
	var v3734 int32
	_ = v3734
	var v3735 int32
	_ = v3735
	var v3738 int32
	_ = v3738
	var v3742 int32
	_ = v3742
	var v3771 int32
	_ = v3771
	var v3781 int32
	_ = v3781
	var v3785 int32
	_ = v3785
	var v3786 int32
	_ = v3786
	var v3787 int32
	_ = v3787
	var v3788 int32
	_ = v3788
	var v3789 int32
	_ = v3789
	var v3792 int32
	_ = v3792
	var v3793 int32
	_ = v3793
	var v3801 int32
	_ = v3801
	var v3832 int32
	_ = v3832
	var v3836 int32
	_ = v3836
	var v3837 int32
	_ = v3837
	var v3842 int32
	_ = v3842
	var v3844 int32
	_ = v3844
	var v3845 int32
	_ = v3845
	var v3846 int32
	_ = v3846
	var v3847 int32
	_ = v3847
	var v3849 int32
	_ = v3849
	var v3850 int32
	_ = v3850
	var v3851 int32
	_ = v3851
	var v3855 int32
	_ = v3855
	var v3858 int32
	_ = v3858
	var v3859 int32
	_ = v3859
	var v3898 int32
	_ = v3898
	var v3900 int32
	_ = v3900
	var v3901 int32
	_ = v3901
	var v3975 int32
	_ = v3975
	var v3976 int32
	_ = v3976
	var v3981 int32
	_ = v3981
	var v3984 int32
	_ = v3984
	var v3987 int32
	_ = v3987
	var v3988 int32
	_ = v3988
	var v3990 int32
	_ = v3990
	var v3998 int32
	_ = v3998
	var v3999 int32
	_ = v3999
	var v4002 int32
	_ = v4002
	var v4005 int32
	_ = v4005
	var v4010 int32
	_ = v4010
	var v4020 int32
	_ = v4020
	var v4025 int32
	_ = v4025
	var v4028 int32
	_ = v4028
	var v4031 int32
	_ = v4031
	var v4040 int32
	_ = v4040
	var v4071 int32
	_ = v4071
	var v4075 int32
	_ = v4075
	var v4077 int32
	_ = v4077
	var v4078 int32
	_ = v4078
	var v4079 int32
	_ = v4079
	var v4082 int32
	_ = v4082
	var v4084 int32
	_ = v4084
	var v4085 int32
	_ = v4085
	var v4162 int32
	_ = v4162
	var v4167 int32
	_ = v4167
	var v4170 int32
	_ = v4170
	var v4177 int32
	_ = v4177
	var v4185 int32
	_ = v4185
	var v4193 int32
	_ = v4193
	var v4215 int32
	_ = v4215
	var v4219 int32
	_ = v4219
	var v4220 int32
	_ = v4220
	var v4221 int32
	_ = v4221
	var v4222 int32
	_ = v4222
	var v4223 int32
	_ = v4223
	var v4232 int32
	_ = v4232
	var v4233 int32
	_ = v4233
	var v4235 int32
	_ = v4235
	var v4238 int32
	_ = v4238
	var v4239 int32
	_ = v4239
	var v4244 int32
	_ = v4244
	var v4251 int32
	_ = v4251
	var v4253 int32
	_ = v4253
	var v4255 int32
	_ = v4255
	var v4258 int32
	_ = v4258
	var v4260 int32
	_ = v4260
	var v4262 int32
	_ = v4262
	var v4269 int32
	_ = v4269
	var v4276 int32
	_ = v4276
	var v4279 int32
	_ = v4279
	var v4282 int32
	_ = v4282
	var v4285 int32
	_ = v4285
	var v4286 int32
	_ = v4286
	var v4287 int32
	_ = v4287
	var v4288 int32
	_ = v4288
	var v4297 int32
	_ = v4297
	var v4298 int32
	_ = v4298
	var v4300 int32
	_ = v4300
	var v4303 int32
	_ = v4303
	var v4304 int32
	_ = v4304
	var v4309 int32
	_ = v4309
	var v4316 int32
	_ = v4316
	var v4318 int32
	_ = v4318
	var v4320 int32
	_ = v4320
	var v4323 int32
	_ = v4323
	var v4325 int32
	_ = v4325
	var v4327 int32
	_ = v4327
	var v4334 int32
	_ = v4334
	var v4341 int32
	_ = v4341
	var v4342 int32
	_ = v4342
	var v4343 int32
	_ = v4343
	var v4352 int32
	_ = v4352
	var v4353 int32
	_ = v4353
	var v4355 int32
	_ = v4355
	var v4358 int32
	_ = v4358
	var v4359 int32
	_ = v4359
	var v4364 int32
	_ = v4364
	var v4371 int32
	_ = v4371
	var v4373 int32
	_ = v4373
	var v4375 int32
	_ = v4375
	var v4378 int32
	_ = v4378
	var v4380 int32
	_ = v4380
	var v4382 int32
	_ = v4382
	var v4389 int32
	_ = v4389
	var v4396 int32
	_ = v4396
	var v4397 int32
	_ = v4397
	var v4398 int32
	_ = v4398
	var v4407 int32
	_ = v4407
	var v4408 int32
	_ = v4408
	var v4410 int32
	_ = v4410
	var v4413 int32
	_ = v4413
	var v4414 int32
	_ = v4414
	var v4419 int32
	_ = v4419
	var v4426 int32
	_ = v4426
	var v4428 int32
	_ = v4428
	var v4430 int32
	_ = v4430
	var v4433 int32
	_ = v4433
	var v4435 int32
	_ = v4435
	var v4437 int32
	_ = v4437
	var v4444 int32
	_ = v4444
	var v4451 int32
	_ = v4451
	var v4454 int32
	_ = v4454
	var v4455 int32
	_ = v4455
	var v4464 int32
	_ = v4464
	var v4465 int32
	_ = v4465
	var v4467 int32
	_ = v4467
	var v4470 int32
	_ = v4470
	var v4471 int32
	_ = v4471
	var v4476 int32
	_ = v4476
	var v4483 int32
	_ = v4483
	var v4485 int32
	_ = v4485
	var v4487 int32
	_ = v4487
	var v4490 int32
	_ = v4490
	var v4492 int32
	_ = v4492
	var v4494 int32
	_ = v4494
	var v4501 int32
	_ = v4501
	var v4508 int32
	_ = v4508
	var v4511 int32
	_ = v4511
	var v4513 int32
	_ = v4513
	var v4514 int32
	_ = v4514
	var v4515 int32
	_ = v4515
	var v4516 int32
	_ = v4516
	var v4519 int32
	_ = v4519
	var v4521 int32
	_ = v4521
	var v4522 int32
	_ = v4522
	var v4524 int32
	_ = v4524
	var v4527 int32
	_ = v4527
	var v4528 int32
	_ = v4528
	var v4532 int32
	_ = v4532
	var v4533 int32
	_ = v4533
	var v4534 int32
	_ = v4534
	var v4535 int32
	_ = v4535
	var v4538 int32
	_ = v4538
	var v4539 int32
	_ = v4539
	var v4540 int32
	_ = v4540
	var v4550 int32
	_ = v4550
	var v4551 int32
	_ = v4551
	var v4553 int32
	_ = v4553
	var v4556 int32
	_ = v4556
	var v4557 int32
	_ = v4557
	var v4562 int32
	_ = v4562
	var v4569 int32
	_ = v4569
	var v4571 int32
	_ = v4571
	var v4573 int32
	_ = v4573
	var v4574 int32
	_ = v4574
	var v4576 int32
	_ = v4576
	var v4578 int32
	_ = v4578
	var v4585 int32
	_ = v4585
	var v4586 int32
	_ = v4586
	var v4589 int32
	_ = v4589
	var v4590 int32
	_ = v4590
	var v4591 int32
	_ = v4591
	var v4601 int32
	_ = v4601
	var v4602 int32
	_ = v4602
	var v4604 int32
	_ = v4604
	var v4607 int32
	_ = v4607
	var v4608 int32
	_ = v4608
	var v4613 int32
	_ = v4613
	var v4620 int32
	_ = v4620
	var v4622 int32
	_ = v4622
	var v4624 int32
	_ = v4624
	var v4625 int32
	_ = v4625
	var v4627 int32
	_ = v4627
	var v4629 int32
	_ = v4629
	var v4636 int32
	_ = v4636
	var v4638 int32
	_ = v4638
	var v4641 int32
	_ = v4641
	var v4642 int32
	_ = v4642
	var v4643 int32
	_ = v4643
	var v4653 int32
	_ = v4653
	var v4654 int32
	_ = v4654
	var v4656 int32
	_ = v4656
	var v4659 int32
	_ = v4659
	var v4660 int32
	_ = v4660
	var v4665 int32
	_ = v4665
	var v4672 int32
	_ = v4672
	var v4674 int32
	_ = v4674
	var v4676 int32
	_ = v4676
	var v4677 int32
	_ = v4677
	var v4679 int32
	_ = v4679
	var v4681 int32
	_ = v4681
	var v4688 int32
	_ = v4688
	var v4689 int32
	_ = v4689
	var v4692 int32
	_ = v4692
	var v4693 int32
	_ = v4693
	var v4694 int32
	_ = v4694
	var v4704 int32
	_ = v4704
	var v4705 int32
	_ = v4705
	var v4707 int32
	_ = v4707
	var v4710 int32
	_ = v4710
	var v4711 int32
	_ = v4711
	var v4716 int32
	_ = v4716
	var v4723 int32
	_ = v4723
	var v4725 int32
	_ = v4725
	var v4727 int32
	_ = v4727
	var v4728 int32
	_ = v4728
	var v4730 int32
	_ = v4730
	var v4732 int32
	_ = v4732
	var v4739 int32
	_ = v4739
	var v4744 int32
	_ = v4744
	var v4745 int32
	_ = v4745
	var v4746 int32
	_ = v4746
	var v4751 int32
	_ = v4751
	var v4752 int32
	_ = v4752
	var v4753 int32
	_ = v4753
	var v4754 int32
	_ = v4754
	var v4755 int32
	_ = v4755
	var v4758 int32
	_ = v4758
	var v4760 int32
	_ = v4760
	var v4767 int32
	_ = v4767
	var v4771 int32
	_ = v4771
	var v4772 int32
	_ = v4772
	var v4775 int32
	_ = v4775
	var v4788 int32
	_ = v4788
	var v4815 int32
	_ = v4815
	var v4819 int32
	_ = v4819
	var v4820 int32
	_ = v4820
	var v4823 int32
	_ = v4823
	var v4824 int32
	_ = v4824
	var v4825 int32
	_ = v4825
	var v4835 int32
	_ = v4835
	var v4836 int32
	_ = v4836
	var v4838 int32
	_ = v4838
	var v4841 int32
	_ = v4841
	var v4842 int32
	_ = v4842
	var v4847 int32
	_ = v4847
	var v4854 int32
	_ = v4854
	var v4856 int32
	_ = v4856
	var v4858 int32
	_ = v4858
	var v4859 int32
	_ = v4859
	var v4861 int32
	_ = v4861
	var v4863 int32
	_ = v4863
	var v4870 int32
	_ = v4870
	var v4871 int32
	_ = v4871
	var v4874 int32
	_ = v4874
	var v4875 int32
	_ = v4875
	var v4876 int32
	_ = v4876
	var v4886 int32
	_ = v4886
	var v4887 int32
	_ = v4887
	var v4889 int32
	_ = v4889
	var v4892 int32
	_ = v4892
	var v4893 int32
	_ = v4893
	var v4898 int32
	_ = v4898
	var v4905 int32
	_ = v4905
	var v4907 int32
	_ = v4907
	var v4909 int32
	_ = v4909
	var v4910 int32
	_ = v4910
	var v4912 int32
	_ = v4912
	var v4914 int32
	_ = v4914
	var v4921 int32
	_ = v4921
	var v4923 int32
	_ = v4923
	var v4926 int32
	_ = v4926
	var v4927 int32
	_ = v4927
	var v4938 int32
	_ = v4938
	var v4966 int32
	_ = v4966
	var v4970 int32
	_ = v4970
	var v4971 int32
	_ = v4971
	var v4974 int32
	_ = v4974
	var v4975 int32
	_ = v4975
	var v4976 int32
	_ = v4976
	var v4986 int32
	_ = v4986
	var v4987 int32
	_ = v4987
	var v4989 int32
	_ = v4989
	var v4992 int32
	_ = v4992
	var v4993 int32
	_ = v4993
	var v4998 int32
	_ = v4998
	var v5005 int32
	_ = v5005
	var v5007 int32
	_ = v5007
	var v5009 int32
	_ = v5009
	var v5010 int32
	_ = v5010
	var v5012 int32
	_ = v5012
	var v5014 int32
	_ = v5014
	var v5021 int32
	_ = v5021
	var v5022 int32
	_ = v5022
	var v5025 int32
	_ = v5025
	var v5026 int32
	_ = v5026
	var v5027 int32
	_ = v5027
	var v5037 int32
	_ = v5037
	var v5038 int32
	_ = v5038
	var v5040 int32
	_ = v5040
	var v5043 int32
	_ = v5043
	var v5044 int32
	_ = v5044
	var v5049 int32
	_ = v5049
	var v5056 int32
	_ = v5056
	var v5058 int32
	_ = v5058
	var v5060 int32
	_ = v5060
	var v5061 int32
	_ = v5061
	var v5063 int32
	_ = v5063
	var v5065 int32
	_ = v5065
	var v5072 int32
	_ = v5072
	var v5085 int32
	_ = v5085
	var v5088 int32
	_ = v5088
	var v5089 int32
	_ = v5089
	var v5128 int32
	_ = v5128
	var v5129 int32
	_ = v5129
	var v5173 int32
	_ = v5173
	var v5186 int32
	_ = v5186
	var v5203 int32
	_ = v5203
	var v5208 int32
	_ = v5208
	var v5211 int32
	_ = v5211
	var v5212 int32
	_ = v5212
	var v5213 int32
	_ = v5213
	var v5216 int32
	_ = v5216
	var v5220 int32
	_ = v5220
	var v5225 int32
	_ = v5225
	var v5226 int32
	_ = v5226
	var v5231 int32
	_ = v5231
	var v5240 int32
	_ = v5240
	var v5242 int32
	_ = v5242
	var v5243 int32
	_ = v5243
	var v5249 int32
	_ = v5249
	var v5254 int32
	_ = v5254
	var v5257 int32
	_ = v5257
	var v5260 int32
	_ = v5260
	var v5261 int32
	_ = v5261
	var v5263 int32
	_ = v5263
	var v5271 int32
	_ = v5271
	var v5272 int32
	_ = v5272
	var v5275 int32
	_ = v5275
	var v5278 int32
	_ = v5278
	var v5283 int32
	_ = v5283
	var v5293 int32
	_ = v5293
	var v5296 int32
	_ = v5296
	var v5303 int32
	_ = v5303
	var v5340 int32
	_ = v5340
	var v5343 int32
	_ = v5343
	var v5349 int32
	_ = v5349
	var v5352 int32
	_ = v5352
	var v5356 int32
	_ = v5356
	v8 = int32(0)
	v37 = m.G0
	v39 = v37 - int32(48)
	m.G0 = v39
	v43 = *(*int32)(unsafe.Add(mBase, uint32(l1)+4))
	if v43 == int32(3) {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	v46 = int32(228)
	goto L3
L2:
	;
	v46 = int32(8)
	goto L3
L3:
	;
	v48 = *(*int32)(unsafe.Add(mBase, uint32(l1+v46)))
	v49 = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v39)+40)) = v49
	*(*int32)(unsafe.Add(mBase, uint32(v39)+20)) = l5
	*(*int32)(unsafe.Add(mBase, uint32(v39)+12)) = v49
	*(*int32)(unsafe.Add(mBase, uint32(v39)+8)) = l6
	switch l4 - int32(4) {
	case 0, 1:
		v122 = v8
		goto L4
	default:
		goto L5
	case 4:
		goto L6
	case 5:
		goto L7
	}
L4:
	;
	*(*uint8)(unsafe.Add(mBase, uint32(v39)+16)) = uint8(v122)
	if l4 != int32(2) {
		goto L26
	} else {
		goto L27
	}
L5:
	;
	v118 = *(*int32)(unsafe.Add(mBase, uint32(l1)+8))
	v119 = *(*int32)(unsafe.Add(mBase, uint32(l2)+8))
	v120 = F_innerrel_is_unique(m, l0, v118, v119, l3, l4, l6)
	mBase = m.M
	v121 = m.ExcPending
	if v121 != 0 {
		goto L22
	} else {
		goto L24
	}
L6:
	;
	v113 = *(*int32)(unsafe.Add(mBase, uint32(l1)+8))
	v114 = *(*int32)(unsafe.Add(mBase, uint32(l2)+8))
	v116 = F_innerrel_is_unique(m, l0, v113, v114, l3, int32(0), l6)
	mBase = m.M
	v117 = m.ExcPending
	if v117 != 0 {
		goto L22
	} else {
		goto L23
	}
L7:
	;
	v57 = *(*int32)(unsafe.Add(mBase, uint32(l5)+4))
	v58 = *(*int32)(unsafe.Add(mBase, uint32(l2)+8))
	v59 = int32(0)
	if v57 == v59 {
		goto L9
	} else {
		goto L10
	}
L8:
	;
	v122 = v112
	goto L4
L9:
	;
	v112 = int32(1)
	goto L8
L10:
	;
	goto L11
L11:
	;
	if v58 == int32(0) {
		v105 = v59
		goto L12
	} else {
		goto L13
	}
L12:
	;
	v112 = v105
	goto L8
L13:
	;
	v68 = *(*int32)(unsafe.Add(mBase, uint32(v57)+4))
	v69 = *(*int32)(unsafe.Add(mBase, uint32(v58)+4))
	if v69 < v68 {
		v105 = v59
		goto L12
	} else {
		goto L14
	}
L14:
	;
	v71 = int32(1)
	if v68 <= v71 {
		goto L15
	} else {
		goto L16
	}
L15:
	;
	v74 = v71
	goto L17
L16:
	;
	v74 = v68
	goto L17
L17:
	;
	v75 = int32(8)
	v80 = int32(0)
	goto L18
L18:
	;
	v87 = v80 << (uint(int32(2)) % 32)
	v89 = *(*int32)(unsafe.Add(mBase, uint32(v57+v75+v87)))
	v91 = *(*int32)(unsafe.Add(mBase, uint32(v58+v75+v87)))
	v94 = v89 & (v91 ^ int32(-1))
	v96 = base.B2i32(v94 == int32(0))
	if v94 != 0 {
		v105 = v96
		goto L12
	} else {
		goto L20
	}
L19:
	;
	v105 = v96
	goto L12
L20:
	;
	v98 = v80 + int32(1)
	if v98 != v74 {
		v80 = v98
		goto L18
	} else {
		goto L21
	}
L21:
	;
	goto L19
L22:
	;
	return
L23:
	;
	v122 = v116
	goto L4
L24:
	;
	v122 = v120
	goto L4
L25:
	;
	if (base.B2i32(l4&int32(-2) == int32(4))|v122)&int32(1) != 0 {
		goto L158
	} else {
		goto L159
	}
L26:
	;
	v126 = int32(1)
	v128 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_add_paths_to_joinrel[0])))
	if v128&v126 == int32(0) {
		v795 = v126
		goto L25
	} else {
		goto L29
	}
L27:
	;
	goto L28
L28:
	;
	if l4 == int32(6) {
		v758 = int32(0)
		v763 = v8
		goto L30
	} else {
		goto L31
	}
L29:
	;
	goto L28
L30:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v39)+12)) = v763
	v795 = v758
	goto L25
L31:
	;
	v137 = int32(1)
	if l6 == int32(0) {
		v713 = v137
		v718 = v8
		goto L32
	} else {
		goto L33
	}
L32:
	;
	if int32(1)<<(uint(l4)%32)&int32(140) != 0 {
		goto L154
	} else {
		goto L155
	}
L33:
	;
	v140 = *(*int32)(unsafe.Add(mBase, uint32(l6)+4))
	if v140 <= int32(0) {
		v713 = v137
		v718 = v8
		goto L32
	} else {
		goto L34
	}
L34:
	;
	v156 = v8
	v158 = v8
	v166 = v8
	goto L35
L35:
	;
	v183 = *(*int32)(unsafe.Add(mBase, uint32(l6)+12))
	v187 = *(*int32)(unsafe.Add(mBase, uint32(v183+v158<<(uint(int32(2))%32))))
	if int32(1)<<(uint(l4)%32)&int32(174) != 0 {
		goto L38
	} else {
		goto L39
	}
L36:
	;
	v713 = v666 ^ int32(1)
	v718 = v676
	goto L32
L37:
	;
	v694 = v158 + int32(1)
	v695 = *(*int32)(unsafe.Add(mBase, uint32(l6)+4))
	if v694 < v695 {
		v156 = v666
		v158 = v694
		v166 = v676
		goto L35
	} else {
		goto L153
	}
L38:
	;
	v188 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v187)+8)))
	if v188 != 0 {
		v666 = v156
		v676 = v166
		goto L37
	} else {
		goto L41
	}
L39:
	;
	goto L40
L40:
	;
	v247 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v187)+9)))
	if v247 == int32(1) {
		goto L58
	} else {
		goto L59
	}
L41:
	;
	v189 = *(*int32)(unsafe.Add(mBase, uint32(v187)+32))
	v190 = *(*int32)(unsafe.Add(mBase, uint32(l1)+8))
	v191 = int32(0)
	if v189 == v191 {
		goto L43
	} else {
		goto L44
	}
L42:
	;
	if v244 == int32(0) {
		v666 = v156
		v676 = v166
		goto L37
	} else {
		goto L56
	}
L43:
	;
	v244 = int32(1)
	goto L42
L44:
	;
	goto L45
L45:
	;
	if v190 == int32(0) {
		v237 = v191
		goto L46
	} else {
		goto L47
	}
L46:
	;
	v244 = v237
	goto L42
L47:
	;
	v200 = *(*int32)(unsafe.Add(mBase, uint32(v189)+4))
	v201 = *(*int32)(unsafe.Add(mBase, uint32(v190)+4))
	if v201 < v200 {
		v237 = v191
		goto L46
	} else {
		goto L48
	}
L48:
	;
	v203 = int32(1)
	if v200 <= v203 {
		goto L49
	} else {
		goto L50
	}
L49:
	;
	v206 = v203
	goto L51
L50:
	;
	v206 = v200
	goto L51
L51:
	;
	v207 = int32(8)
	v212 = int32(0)
	goto L52
L52:
	;
	v219 = v212 << (uint(int32(2)) % 32)
	v221 = *(*int32)(unsafe.Add(mBase, uint32(v189+v207+v219)))
	v223 = *(*int32)(unsafe.Add(mBase, uint32(v190+v207+v219)))
	v226 = v221 & (v223 ^ int32(-1))
	v228 = base.B2i32(v226 == int32(0))
	if v226 != 0 {
		v237 = v228
		goto L46
	} else {
		goto L54
	}
L53:
	;
	v237 = v228
	goto L46
L54:
	;
	v230 = v212 + int32(1)
	if v230 != v206 {
		v212 = v230
		goto L52
	} else {
		goto L55
	}
L55:
	;
	goto L53
L56:
	;
	goto L40
L57:
	;
	v256 = *(*int32)(unsafe.Add(mBase, uint32(l3)+8))
	v257 = *(*int32)(unsafe.Add(mBase, uint32(v187)+44))
	v258 = *(*int32)(unsafe.Add(mBase, uint32(l2)+8))
	v259 = int32(0)
	if v257 == v259 {
		goto L69
	} else {
		goto L70
	}
L58:
	;
	v250 = *(*int32)(unsafe.Add(mBase, uint32(v187)+96))
	if v250 != 0 {
		goto L57
	} else {
		goto L61
	}
L59:
	;
	goto L60
L60:
	;
	v251 = *(*int32)(unsafe.Add(mBase, uint32(v187)+4))
	if v251 != 0 {
		goto L62
	} else {
		goto L63
	}
L61:
	;
	goto L60
L62:
	;
	v252 = *(*int32)(unsafe.Add(mBase, uint32(v251)))
	if v252 == int32(7) {
		v666 = v156
		v676 = v166
		goto L37
	} else {
		goto L65
	}
L63:
	;
	goto L64
L64:
	;
	v666 = int32(1)
	v676 = v166
	goto L37
L65:
	;
	goto L64
L66:
	;
	v497 = *(*int32)(unsafe.Add(mBase, uint32(v187)+100))
	v498 = *(*int32)(unsafe.Add(mBase, uint32(v497)+56))
	if v498 != 0 {
		goto L134
	} else {
		goto L135
	}
L67:
	;
	v374 = *(*int32)(unsafe.Add(mBase, uint32(v187)+44))
	v375 = int32(0)
	if v374 == v375 {
		goto L99
	} else {
		goto L100
	}
L68:
	;
	if v312 == int32(0) {
		goto L67
	} else {
		goto L82
	}
L69:
	;
	v312 = int32(1)
	goto L68
L70:
	;
	goto L71
L71:
	;
	if v258 == int32(0) {
		v305 = v259
		goto L72
	} else {
		goto L73
	}
L72:
	;
	v312 = v305
	goto L68
L73:
	;
	v268 = *(*int32)(unsafe.Add(mBase, uint32(v257)+4))
	v269 = *(*int32)(unsafe.Add(mBase, uint32(v258)+4))
	if v269 < v268 {
		v305 = v259
		goto L72
	} else {
		goto L74
	}
L74:
	;
	v271 = int32(1)
	if v268 <= v271 {
		goto L75
	} else {
		goto L76
	}
L75:
	;
	v274 = v271
	goto L77
L76:
	;
	v274 = v268
	goto L77
L77:
	;
	v275 = int32(8)
	v280 = int32(0)
	goto L78
L78:
	;
	v287 = v280 << (uint(int32(2)) % 32)
	v289 = *(*int32)(unsafe.Add(mBase, uint32(v257+v275+v287)))
	v291 = *(*int32)(unsafe.Add(mBase, uint32(v258+v275+v287)))
	v294 = v289 & (v291 ^ int32(-1))
	v296 = base.B2i32(v294 == int32(0))
	if v294 != 0 {
		v305 = v296
		goto L72
	} else {
		goto L80
	}
L79:
	;
	v305 = v296
	goto L72
L80:
	;
	v298 = v280 + int32(1)
	if v298 != v274 {
		v280 = v298
		goto L78
	} else {
		goto L81
	}
L81:
	;
	goto L79
L82:
	;
	v315 = *(*int32)(unsafe.Add(mBase, uint32(v187)+48))
	v316 = int32(0)
	if v315 == v316 {
		goto L84
	} else {
		goto L85
	}
L83:
	;
	if v369 == int32(0) {
		goto L67
	} else {
		goto L97
	}
L84:
	;
	v369 = int32(1)
	goto L83
L85:
	;
	goto L86
L86:
	;
	if v256 == int32(0) {
		v362 = v316
		goto L87
	} else {
		goto L88
	}
L87:
	;
	v369 = v362
	goto L83
L88:
	;
	v325 = *(*int32)(unsafe.Add(mBase, uint32(v315)+4))
	v326 = *(*int32)(unsafe.Add(mBase, uint32(v256)+4))
	if v326 < v325 {
		v362 = v316
		goto L87
	} else {
		goto L89
	}
L89:
	;
	v328 = int32(1)
	if v325 <= v328 {
		goto L90
	} else {
		goto L91
	}
L90:
	;
	v331 = v328
	goto L92
L91:
	;
	v331 = v325
	goto L92
L92:
	;
	v332 = int32(8)
	v337 = int32(0)
	goto L93
L93:
	;
	v344 = v337 << (uint(int32(2)) % 32)
	v346 = *(*int32)(unsafe.Add(mBase, uint32(v315+v332+v344)))
	v348 = *(*int32)(unsafe.Add(mBase, uint32(v256+v332+v344)))
	v351 = v346 & (v348 ^ int32(-1))
	v353 = base.B2i32(v351 == int32(0))
	if v351 != 0 {
		v362 = v353
		goto L87
	} else {
		goto L95
	}
L94:
	;
	v362 = v353
	goto L87
L95:
	;
	v355 = v337 + int32(1)
	if v355 != v331 {
		v337 = v355
		goto L93
	} else {
		goto L96
	}
L96:
	;
	goto L94
L97:
	;
	v372 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(v187)+120)) = uint8(v372)
	goto L66
L98:
	;
	if v428 == int32(0) {
		goto L112
	} else {
		goto L113
	}
L99:
	;
	v428 = int32(1)
	goto L98
L100:
	;
	goto L101
L101:
	;
	if v256 == int32(0) {
		v421 = v375
		goto L102
	} else {
		goto L103
	}
L102:
	;
	v428 = v421
	goto L98
L103:
	;
	v384 = *(*int32)(unsafe.Add(mBase, uint32(v374)+4))
	v385 = *(*int32)(unsafe.Add(mBase, uint32(v256)+4))
	if v385 < v384 {
		v421 = v375
		goto L102
	} else {
		goto L104
	}
L104:
	;
	v387 = int32(1)
	if v384 <= v387 {
		goto L105
	} else {
		goto L106
	}
L105:
	;
	v390 = v387
	goto L107
L106:
	;
	v390 = v384
	goto L107
L107:
	;
	v391 = int32(8)
	v396 = int32(0)
	goto L108
L108:
	;
	v403 = v396 << (uint(int32(2)) % 32)
	v405 = *(*int32)(unsafe.Add(mBase, uint32(v374+v391+v403)))
	v407 = *(*int32)(unsafe.Add(mBase, uint32(v256+v391+v403)))
	v410 = v405 & (v407 ^ int32(-1))
	v412 = base.B2i32(v410 == int32(0))
	if v410 != 0 {
		v421 = v412
		goto L102
	} else {
		goto L110
	}
L109:
	;
	v421 = v412
	goto L102
L110:
	;
	v414 = v396 + int32(1)
	if v414 != v390 {
		v396 = v414
		goto L108
	} else {
		goto L111
	}
L111:
	;
	goto L109
L112:
	;
	v666 = int32(1)
	v676 = v166
	goto L37
L113:
	;
	goto L114
L114:
	;
	v432 = *(*int32)(unsafe.Add(mBase, uint32(v187)+48))
	v433 = int32(0)
	if v432 == v433 {
		goto L116
	} else {
		goto L117
	}
L115:
	;
	if v486 == int32(0) {
		goto L129
	} else {
		goto L130
	}
L116:
	;
	v486 = int32(1)
	goto L115
L117:
	;
	goto L118
L118:
	;
	if v258 == int32(0) {
		v479 = v433
		goto L119
	} else {
		goto L120
	}
L119:
	;
	v486 = v479
	goto L115
L120:
	;
	v442 = *(*int32)(unsafe.Add(mBase, uint32(v432)+4))
	v443 = *(*int32)(unsafe.Add(mBase, uint32(v258)+4))
	if v443 < v442 {
		v479 = v433
		goto L119
	} else {
		goto L121
	}
L121:
	;
	v445 = int32(1)
	if v442 <= v445 {
		goto L122
	} else {
		goto L123
	}
L122:
	;
	v448 = v445
	goto L124
L123:
	;
	v448 = v442
	goto L124
L124:
	;
	v449 = int32(8)
	v454 = int32(0)
	goto L125
L125:
	;
	v461 = v454 << (uint(int32(2)) % 32)
	v463 = *(*int32)(unsafe.Add(mBase, uint32(v432+v449+v461)))
	v465 = *(*int32)(unsafe.Add(mBase, uint32(v258+v449+v461)))
	v468 = v463 & (v465 ^ int32(-1))
	v470 = base.B2i32(v468 == int32(0))
	if v468 != 0 {
		v479 = v470
		goto L119
	} else {
		goto L127
	}
L126:
	;
	v479 = v470
	goto L119
L127:
	;
	v472 = v454 + int32(1)
	if v472 != v448 {
		v454 = v472
		goto L125
	} else {
		goto L128
	}
L128:
	;
	goto L126
L129:
	;
	v666 = int32(1)
	v676 = v166
	goto L37
L130:
	;
	goto L131
L131:
	;
	v490 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v187)+120)) = uint8(v490)
	v492 = *(*int32)(unsafe.Add(mBase, uint32(v187)+4))
	v493 = *(*int32)(unsafe.Add(mBase, uint32(v492)+4))
	v494 = F_get_commutator(m, v493)
	mBase = m.M
	v495 = m.ExcPending
	if v495 != 0 {
		goto L22
	} else {
		goto L132
	}
L132:
	;
	if v494 != 0 {
		goto L66
	} else {
		goto L133
	}
L133:
	;
	v666 = int32(1)
	v676 = v166
	goto L37
L134:
	;
	v516 = v498
	goto L137
L135:
	;
	goto L136
L136:
	;
	v573 = *(*int32)(unsafe.Add(mBase, uint32(v187)+104))
	v574 = *(*int32)(unsafe.Add(mBase, uint32(v573)+56))
	if v574 != 0 {
		goto L140
	} else {
		goto L141
	}
L137:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v187)+100)) = v516
	v536 = *(*int32)(unsafe.Add(mBase, uint32(v516)+56))
	if v536 != 0 {
		v516 = v536
		goto L137
	} else {
		goto L139
	}
L138:
	;
	goto L136
L139:
	;
	goto L138
L140:
	;
	v592 = v574
	goto L143
L141:
	;
	goto L142
L142:
	;
	v649 = *(*int32)(unsafe.Add(mBase, uint32(v187)+100))
	v650 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v649)+40)))
	if v650 != 0 {
		goto L146
	} else {
		goto L147
	}
L143:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v187)+104)) = v592
	v612 = *(*int32)(unsafe.Add(mBase, uint32(v592)+56))
	if v612 != 0 {
		v592 = v612
		goto L143
	} else {
		goto L145
	}
L144:
	;
	goto L142
L145:
	;
	goto L144
L146:
	;
	v666 = int32(1)
	v676 = v166
	goto L37
L147:
	;
	goto L148
L148:
	;
	v652 = *(*int32)(unsafe.Add(mBase, uint32(v187)+104))
	v653 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v652)+40)))
	if v653 != 0 {
		goto L149
	} else {
		goto L150
	}
L149:
	;
	v666 = int32(1)
	v676 = v166
	goto L37
L150:
	;
	goto L151
L151:
	;
	v655 = F_lappend(m, v166, v187)
	mBase = m.M
	v656 = m.ExcPending
	if v656 != 0 {
		goto L22
	} else {
		goto L152
	}
L152:
	;
	v666 = v156
	v676 = v655
	goto L37
L153:
	;
	goto L36
L154:
	;
	v742 = base.B2i32(base.Ui32(l4) <= base.Ui32(int32(7)))
	goto L156
L155:
	;
	v742 = int32(0)
	goto L156
L156:
	;
	if v742 != 0 {
		v758 = v713
		v763 = v718
		goto L30
	} else {
		goto L157
	}
L157:
	;
	v758 = int32(1)
	v763 = v718
	goto L30
L158:
	;
	v825 = v39 + int32(24)
	v826 = int32(0)
	v827 = m.G0
	v829 = v827 + int32(-64)
	m.G0 = v829
	v834 = int32(1) << (uint(l4) % 32) & int32(174)
	if v834 == v826 {
		goto L162
	} else {
		goto L163
	}
L159:
	;
	goto L160
L160:
	;
	v1072 = *(*int32)(unsafe.Add(mBase, uint32(l0)+112))
	if v1072 == int32(0) {
		goto L204
	} else {
		goto L205
	}
L161:
	;
	v986 = int32(5)
	if l4 == v986 {
		goto L188
	} else {
		goto L189
	}
L162:
	;
	v954 = l6
	goto L161
L163:
	;
	goto L164
L164:
	;
	if l6 == int32(0) {
		v954 = v826
		goto L161
	} else {
		goto L165
	}
L165:
	;
	v839 = *(*int32)(unsafe.Add(mBase, uint32(l6)+4))
	if v839 <= int32(0) {
		v954 = v826
		goto L161
	} else {
		goto L166
	}
L166:
	;
	v847 = v826
	v850 = v8
	goto L167
L167:
	;
	v878 = *(*int32)(unsafe.Add(mBase, uint32(l6)+12))
	v882 = *(*int32)(unsafe.Add(mBase, uint32(v878+v850<<(uint(int32(2))%32))))
	v883 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v882)+8)))
	if v883 != 0 {
		v944 = v847
		goto L169
	} else {
		goto L170
	}
L168:
	;
	v954 = v944
	goto L161
L169:
	;
	v946 = v850 + int32(1)
	v947 = *(*int32)(unsafe.Add(mBase, uint32(l6)+4))
	if v946 < v947 {
		v847 = v944
		v850 = v946
		goto L167
	} else {
		goto L187
	}
L170:
	;
	v884 = *(*int32)(unsafe.Add(mBase, uint32(v882)+32))
	v885 = *(*int32)(unsafe.Add(mBase, uint32(l1)+8))
	v886 = int32(0)
	if v884 == v886 {
		goto L172
	} else {
		goto L173
	}
L171:
	;
	if v939 == int32(0) {
		v944 = v847
		goto L169
	} else {
		goto L185
	}
L172:
	;
	v939 = int32(1)
	goto L171
L173:
	;
	goto L174
L174:
	;
	if v885 == int32(0) {
		v932 = v886
		goto L175
	} else {
		goto L176
	}
L175:
	;
	v939 = v932
	goto L171
L176:
	;
	v895 = *(*int32)(unsafe.Add(mBase, uint32(v884)+4))
	v896 = *(*int32)(unsafe.Add(mBase, uint32(v885)+4))
	if v896 < v895 {
		v932 = v886
		goto L175
	} else {
		goto L177
	}
L177:
	;
	v898 = int32(1)
	if v895 <= v898 {
		goto L178
	} else {
		goto L179
	}
L178:
	;
	v901 = v898
	goto L180
L179:
	;
	v901 = v895
	goto L180
L180:
	;
	v902 = int32(8)
	v907 = int32(0)
	goto L181
L181:
	;
	v914 = v907 << (uint(int32(2)) % 32)
	v916 = *(*int32)(unsafe.Add(mBase, uint32(v884+v902+v914)))
	v918 = *(*int32)(unsafe.Add(mBase, uint32(v885+v902+v914)))
	v921 = v916 & (v918 ^ int32(-1))
	v923 = base.B2i32(v921 == int32(0))
	if v921 != 0 {
		v932 = v923
		goto L175
	} else {
		goto L183
	}
L182:
	;
	v932 = v923
	goto L175
L183:
	;
	v925 = v907 + int32(1)
	if v925 != v901 {
		v907 = v925
		goto L181
	} else {
		goto L184
	}
L184:
	;
	goto L182
L185:
	;
	v942 = F_lappend(m, v847, v882)
	mBase = m.M
	v943 = m.ExcPending
	if v943 != 0 {
		goto L22
	} else {
		goto L186
	}
L186:
	;
	v944 = v942
	goto L169
L187:
	;
	goto L168
L188:
	;
	v990 = v986
	goto L190
L189:
	;
	v990 = int32(4)
	goto L190
L190:
	;
	v991 = F_clauselist_selectivity(m, l0, v954, int32(0), v990, l5)
	mBase = m.M
	v992 = m.ExcPending
	if v992 != 0 {
		goto L22
	} else {
		goto L191
	}
L191:
	;
	v994 = v827 + int32(-56)
	v995 = *(*int32)(unsafe.Add(mBase, uint32(l2)+8))
	v996 = *(*int32)(unsafe.Add(mBase, uint32(l3)+8))
	v997 = int64(0)
	*(*int64)(unsafe.Add(mBase, uint32(v994)+48)) = v997
	*(*int32)(unsafe.Add(mBase, uint32(v994)+16)) = v996
	*(*int32)(unsafe.Add(mBase, uint32(v994)+12)) = v995
	*(*int32)(unsafe.Add(mBase, uint32(v994)+8)) = v996
	*(*int32)(unsafe.Add(mBase, uint32(v994)+4)) = v995
	*(*int32)(unsafe.Add(mBase, uint32(v994))) = int32(320)
	*(*int64)(unsafe.Add(mBase, uint32(v994)+20)) = v997
	*(*int64)(unsafe.Add(mBase, uint32(v994)+28)) = v997
	*(*int64)(unsafe.Add(mBase, uint32(v994)+36)) = v997
	*(*int32)(unsafe.Add(mBase, uint32(v994)+43)) = int32(0)
	goto L192
L192:
	;
	v1013 = int32(0)
	v1015 = F_clauselist_selectivity(m, l0, v954, v1013, v1013, v994)
	mBase = m.M
	v1016 = m.ExcPending
	if v1016 != 0 {
		goto L22
	} else {
		goto L193
	}
L193:
	;
	if v834 != 0 {
		goto L194
	} else {
		goto L195
	}
L194:
	;
	F_list_free(m, v954)
	mBase = m.M
	v1018 = m.ExcPending
	if v1018 != 0 {
		goto L22
	} else {
		goto L197
	}
L195:
	;
	goto L196
L196:
	;
	if base.F64_gt(v991, float64(0)) != 0 {
		goto L198
	} else {
		goto L199
	}
L197:
	;
	goto L196
L198:
	;
	v1021 = float64(1)
	v1022 = *(*float64)(unsafe.Add(mBase, uint32(l3)+16))
	v1024 = base.F64_div(base.F64_mul(v1015, v1022), v991)
	if base.F64_lt(v1024, v1021) != 0 {
		goto L201
	} else {
		goto L202
	}
L199:
	;
	v1030 = float64(1)
	goto L200
L200:
	;
	*(*float64)(unsafe.Add(mBase, uint32(v825)+8)) = v1030
	*(*float64)(unsafe.Add(mBase, uint32(v825))) = v991
	m.G0 = v829 - int32(-64)
	goto L160
L201:
	;
	v1027 = v1021
	goto L203
L202:
	;
	v1027 = v1024
	goto L203
L203:
	;
	v1030 = v1027
	goto L200
L204:
	;
	v1371 = *(*int32)(unsafe.Add(mBase, uint32(v39)+40))
	v1372 = *(*int32)(unsafe.Add(mBase, uint32(l1)+64))
	v1373 = F_bms_add_members(m, v1371, v1372)
	mBase = m.M
	v1374 = m.ExcPending
	if v1374 != 0 {
		goto L22
	} else {
		goto L273
	}
L205:
	;
	v1075 = *(*int32)(unsafe.Add(mBase, uint32(v1072)+4))
	if v1075 <= int32(0) {
		goto L204
	} else {
		goto L206
	}
L206:
	;
	v1085 = int32(0)
	goto L207
L207:
	;
	v1115 = *(*int32)(unsafe.Add(mBase, uint32(v1072)+12))
	v1119 = *(*int32)(unsafe.Add(mBase, uint32(v1115+v1085<<(uint(int32(2))%32))))
	v1120 = *(*int32)(unsafe.Add(mBase, uint32(v1119)+8))
	v1121 = int32(0)
	if base.B2i32(v48 == v1121)|base.B2i32(v1120 == v1121) != 0 {
		v1166 = v1121
		goto L211
	} else {
		goto L212
	}
L208:
	;
	goto L204
L209:
	;
	v1224 = *(*int32)(unsafe.Add(mBase, uint32(v1119)+20))
	if v1224 != int32(2) {
		goto L240
	} else {
		goto L241
	}
L210:
	;
	if v1166 == int32(0) {
		goto L209
	} else {
		goto L223
	}
L211:
	;
	goto L210
L212:
	;
	v1131 = *(*int32)(unsafe.Add(mBase, uint32(v48)+4))
	v1132 = *(*int32)(unsafe.Add(mBase, uint32(v1120)+4))
	if v1131 < v1132 {
		goto L213
	} else {
		goto L214
	}
L213:
	;
	v1134 = v1131
	goto L215
L214:
	;
	v1134 = v1132
	goto L215
L215:
	;
	if v1134 <= int32(1) {
		goto L216
	} else {
		goto L217
	}
L216:
	;
	v1137 = int32(1)
	goto L218
L217:
	;
	v1137 = v1134
	goto L218
L218:
	;
	v1138 = int32(8)
	v1143 = int32(0)
	goto L219
L219:
	;
	v1150 = v1143 << (uint(int32(2)) % 32)
	v1152 = *(*int32)(unsafe.Add(mBase, uint32(v1120+v1138+v1150)))
	v1154 = *(*int32)(unsafe.Add(mBase, uint32(v48+v1138+v1150)))
	v1155 = v1152 & v1154
	v1157 = base.B2i32(v1155 != int32(0))
	if v1155 != 0 {
		v1166 = v1157
		goto L211
	} else {
		goto L221
	}
L220:
	;
	v1166 = v1157
	goto L211
L221:
	;
	v1159 = v1143 + int32(1)
	if v1159 != v1137 {
		v1143 = v1159
		goto L219
	} else {
		goto L222
	}
L222:
	;
	goto L220
L223:
	;
	v1169 = *(*int32)(unsafe.Add(mBase, uint32(v1119)+4))
	v1170 = int32(0)
	if base.B2i32(v48 == v1170)|base.B2i32(v1169 == v1170) != 0 {
		v1215 = v1170
		goto L225
	} else {
		goto L226
	}
L224:
	;
	if v1215 != 0 {
		goto L209
	} else {
		goto L237
	}
L225:
	;
	goto L224
L226:
	;
	v1180 = *(*int32)(unsafe.Add(mBase, uint32(v48)+4))
	v1181 = *(*int32)(unsafe.Add(mBase, uint32(v1169)+4))
	if v1180 < v1181 {
		goto L227
	} else {
		goto L228
	}
L227:
	;
	v1183 = v1180
	goto L229
L228:
	;
	v1183 = v1181
	goto L229
L229:
	;
	if v1183 <= int32(1) {
		goto L230
	} else {
		goto L231
	}
L230:
	;
	v1186 = int32(1)
	goto L232
L231:
	;
	v1186 = v1183
	goto L232
L232:
	;
	v1187 = int32(8)
	v1192 = int32(0)
	goto L233
L233:
	;
	v1199 = v1192 << (uint(int32(2)) % 32)
	v1201 = *(*int32)(unsafe.Add(mBase, uint32(v1169+v1187+v1199)))
	v1203 = *(*int32)(unsafe.Add(mBase, uint32(v48+v1187+v1199)))
	v1204 = v1201 & v1203
	v1206 = base.B2i32(v1204 != int32(0))
	if v1204 != 0 {
		v1215 = v1206
		goto L225
	} else {
		goto L235
	}
L234:
	;
	v1215 = v1206
	goto L225
L235:
	;
	v1208 = v1192 + int32(1)
	if v1208 != v1186 {
		v1192 = v1208
		goto L233
	} else {
		goto L236
	}
L236:
	;
	goto L234
L237:
	;
	v1216 = *(*int32)(unsafe.Add(mBase, uint32(v39)+40))
	v1217 = *(*int32)(unsafe.Add(mBase, uint32(l0)+44))
	v1218 = *(*int32)(unsafe.Add(mBase, uint32(v1119)+8))
	v1219 = F_bms_difference(m, v1217, v1218)
	mBase = m.M
	v1220 = m.ExcPending
	if v1220 != 0 {
		goto L22
	} else {
		goto L238
	}
L238:
	;
	v1221 = F_bms_join(m, v1216, v1219)
	mBase = m.M
	v1222 = m.ExcPending
	if v1222 != 0 {
		goto L22
	} else {
		goto L239
	}
L239:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v39)+40)) = v1221
	goto L209
L240:
	;
	v1332 = v1085 + int32(1)
	v1333 = *(*int32)(unsafe.Add(mBase, uint32(v1072)+4))
	if v1332 < v1333 {
		v1085 = v1332
		goto L207
	} else {
		goto L272
	}
L241:
	;
	v1227 = *(*int32)(unsafe.Add(mBase, uint32(v1119)+4))
	v1228 = int32(0)
	if base.B2i32(v48 == v1228)|base.B2i32(v1227 == v1228) != 0 {
		v1273 = v1228
		goto L243
	} else {
		goto L244
	}
L242:
	;
	if v1273 == int32(0) {
		goto L240
	} else {
		goto L255
	}
L243:
	;
	goto L242
L244:
	;
	v1238 = *(*int32)(unsafe.Add(mBase, uint32(v48)+4))
	v1239 = *(*int32)(unsafe.Add(mBase, uint32(v1227)+4))
	if v1238 < v1239 {
		goto L245
	} else {
		goto L246
	}
L245:
	;
	v1241 = v1238
	goto L247
L246:
	;
	v1241 = v1239
	goto L247
L247:
	;
	if v1241 <= int32(1) {
		goto L248
	} else {
		goto L249
	}
L248:
	;
	v1244 = int32(1)
	goto L250
L249:
	;
	v1244 = v1241
	goto L250
L250:
	;
	v1245 = int32(8)
	v1250 = int32(0)
	goto L251
L251:
	;
	v1257 = v1250 << (uint(int32(2)) % 32)
	v1259 = *(*int32)(unsafe.Add(mBase, uint32(v1227+v1245+v1257)))
	v1261 = *(*int32)(unsafe.Add(mBase, uint32(v48+v1245+v1257)))
	v1262 = v1259 & v1261
	v1264 = base.B2i32(v1262 != int32(0))
	if v1262 != 0 {
		v1273 = v1264
		goto L243
	} else {
		goto L253
	}
L252:
	;
	v1273 = v1264
	goto L243
L253:
	;
	v1266 = v1250 + int32(1)
	if v1266 != v1244 {
		v1250 = v1266
		goto L251
	} else {
		goto L254
	}
L254:
	;
	goto L252
L255:
	;
	v1276 = *(*int32)(unsafe.Add(mBase, uint32(v1119)+8))
	v1277 = int32(0)
	if base.B2i32(v48 == v1277)|base.B2i32(v1276 == v1277) != 0 {
		v1322 = v1277
		goto L257
	} else {
		goto L258
	}
L256:
	;
	if v1322 != 0 {
		goto L240
	} else {
		goto L269
	}
L257:
	;
	goto L256
L258:
	;
	v1287 = *(*int32)(unsafe.Add(mBase, uint32(v48)+4))
	v1288 = *(*int32)(unsafe.Add(mBase, uint32(v1276)+4))
	if v1287 < v1288 {
		goto L259
	} else {
		goto L260
	}
L259:
	;
	v1290 = v1287
	goto L261
L260:
	;
	v1290 = v1288
	goto L261
L261:
	;
	if v1290 <= int32(1) {
		goto L262
	} else {
		goto L263
	}
L262:
	;
	v1293 = int32(1)
	goto L264
L263:
	;
	v1293 = v1290
	goto L264
L264:
	;
	v1294 = int32(8)
	v1299 = int32(0)
	goto L265
L265:
	;
	v1306 = v1299 << (uint(int32(2)) % 32)
	v1308 = *(*int32)(unsafe.Add(mBase, uint32(v1276+v1294+v1306)))
	v1310 = *(*int32)(unsafe.Add(mBase, uint32(v48+v1294+v1306)))
	v1311 = v1308 & v1310
	v1313 = base.B2i32(v1311 != int32(0))
	if v1311 != 0 {
		v1322 = v1313
		goto L257
	} else {
		goto L267
	}
L266:
	;
	v1322 = v1313
	goto L257
L267:
	;
	v1315 = v1299 + int32(1)
	if v1315 != v1293 {
		v1299 = v1315
		goto L265
	} else {
		goto L268
	}
L268:
	;
	goto L266
L269:
	;
	v1323 = *(*int32)(unsafe.Add(mBase, uint32(v39)+40))
	v1324 = *(*int32)(unsafe.Add(mBase, uint32(l0)+44))
	v1325 = *(*int32)(unsafe.Add(mBase, uint32(v1119)+4))
	v1326 = F_bms_difference(m, v1324, v1325)
	mBase = m.M
	v1327 = m.ExcPending
	if v1327 != 0 {
		goto L22
	} else {
		goto L270
	}
L270:
	;
	v1328 = F_bms_join(m, v1323, v1326)
	mBase = m.M
	v1329 = m.ExcPending
	if v1329 != 0 {
		goto L22
	} else {
		goto L271
	}
L271:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v39)+40)) = v1328
	goto L240
L272:
	;
	goto L208
L273:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v39)+40)) = v1373
	if v795&int32(1) == int32(0) {
		goto L274
	} else {
		goto L275
	}
L274:
	;
	if l4 != int32(2) {
		goto L763
	} else {
		goto L764
	}
L275:
	;
	v1380 = *(*int32)(unsafe.Add(mBase, uint32(v39)+12))
	if v1380 == int32(0) {
		goto L276
	} else {
		goto L277
	}
L276:
	;
	v2999 = *(*int32)(unsafe.Add(mBase, uint32(l3)+48))
	v3000 = int32(1)
	switch l4 {
	case 0, 1, 4, 5:
		v3017 = l4
		v3018 = v3000
		v3019 = v8
		goto L538
	case 2, 3, 7:
		goto L539
	case 6:
		goto L274
	case 8, 9:
		goto L541
	default:
		goto L540
	}
L277:
	;
	v1383 = *(*int32)(unsafe.Add(mBase, uint32(l3)+48))
	v1384 = *(*int32)(unsafe.Add(mBase, uint32(l2)+48))
	v1385 = *(*int32)(unsafe.Add(mBase, uint32(v1384)+16))
	if v1385 == int32(0) {
		goto L278
	} else {
		goto L279
	}
L278:
	;
	v1488 = *(*int32)(unsafe.Add(mBase, uint32(v1383)+16))
	if v1488 == int32(0) {
		goto L309
	} else {
		goto L310
	}
L279:
	;
	v1388 = *(*int32)(unsafe.Add(mBase, uint32(v1385)+4))
	v1389 = *(*int32)(unsafe.Add(mBase, uint32(l3)+8))
	v1390 = int32(0)
	if base.B2i32(v1388 == v1390)|base.B2i32(v1389 == v1390) != 0 {
		v1435 = v1390
		goto L281
	} else {
		goto L282
	}
L280:
	;
	if v1435 != 0 {
		goto L276
	} else {
		goto L293
	}
L281:
	;
	goto L280
L282:
	;
	v1400 = *(*int32)(unsafe.Add(mBase, uint32(v1388)+4))
	v1401 = *(*int32)(unsafe.Add(mBase, uint32(v1389)+4))
	if v1400 < v1401 {
		goto L283
	} else {
		goto L284
	}
L283:
	;
	v1403 = v1400
	goto L285
L284:
	;
	v1403 = v1401
	goto L285
L285:
	;
	if v1403 <= int32(1) {
		goto L286
	} else {
		goto L287
	}
L286:
	;
	v1406 = int32(1)
	goto L288
L287:
	;
	v1406 = v1403
	goto L288
L288:
	;
	v1407 = int32(8)
	v1412 = int32(0)
	goto L289
L289:
	;
	v1419 = v1412 << (uint(int32(2)) % 32)
	v1421 = *(*int32)(unsafe.Add(mBase, uint32(v1389+v1407+v1419)))
	v1423 = *(*int32)(unsafe.Add(mBase, uint32(v1388+v1407+v1419)))
	v1424 = v1421 & v1423
	v1426 = base.B2i32(v1424 != int32(0))
	if v1424 != 0 {
		v1435 = v1426
		goto L281
	} else {
		goto L291
	}
L290:
	;
	v1435 = v1426
	goto L281
L291:
	;
	v1428 = v1412 + int32(1)
	if v1428 != v1406 {
		v1412 = v1428
		goto L289
	} else {
		goto L292
	}
L292:
	;
	goto L290
L293:
	;
	v1436 = *(*int32)(unsafe.Add(mBase, uint32(v1384)+16))
	if v1436 == int32(0) {
		goto L278
	} else {
		goto L294
	}
L294:
	;
	v1439 = *(*int32)(unsafe.Add(mBase, uint32(v1436)+4))
	v1440 = *(*int32)(unsafe.Add(mBase, uint32(l3)+228))
	v1441 = int32(0)
	if base.B2i32(v1439 == v1441)|base.B2i32(v1440 == v1441) != 0 {
		v1486 = v1441
		goto L296
	} else {
		goto L297
	}
L295:
	;
	if v1486 != 0 {
		goto L276
	} else {
		goto L308
	}
L296:
	;
	goto L295
L297:
	;
	v1451 = *(*int32)(unsafe.Add(mBase, uint32(v1439)+4))
	v1452 = *(*int32)(unsafe.Add(mBase, uint32(v1440)+4))
	if v1451 < v1452 {
		goto L298
	} else {
		goto L299
	}
L298:
	;
	v1454 = v1451
	goto L300
L299:
	;
	v1454 = v1452
	goto L300
L300:
	;
	if v1454 <= int32(1) {
		goto L301
	} else {
		goto L302
	}
L301:
	;
	v1457 = int32(1)
	goto L303
L302:
	;
	v1457 = v1454
	goto L303
L303:
	;
	v1458 = int32(8)
	v1463 = int32(0)
	goto L304
L304:
	;
	v1470 = v1463 << (uint(int32(2)) % 32)
	v1472 = *(*int32)(unsafe.Add(mBase, uint32(v1440+v1458+v1470)))
	v1474 = *(*int32)(unsafe.Add(mBase, uint32(v1439+v1458+v1470)))
	v1475 = v1472 & v1474
	v1477 = base.B2i32(v1475 != int32(0))
	if v1475 != 0 {
		v1486 = v1477
		goto L296
	} else {
		goto L306
	}
L305:
	;
	v1486 = v1477
	goto L296
L306:
	;
	v1479 = v1463 + int32(1)
	if v1479 != v1457 {
		v1463 = v1479
		goto L304
	} else {
		goto L307
	}
L307:
	;
	goto L305
L308:
	;
	goto L278
L309:
	;
	switch l4 - int32(8) {
	case 0:
		goto L342
	case 1:
		goto L341
	default:
		v1601 = v1383
		v1602 = l4
		v1603 = v1384
		goto L340
	}
L310:
	;
	v1491 = *(*int32)(unsafe.Add(mBase, uint32(v1488)+4))
	v1492 = *(*int32)(unsafe.Add(mBase, uint32(l2)+8))
	v1493 = int32(0)
	if base.B2i32(v1491 == v1493)|base.B2i32(v1492 == v1493) != 0 {
		v1538 = v1493
		goto L312
	} else {
		goto L313
	}
L311:
	;
	if v1538 != 0 {
		goto L276
	} else {
		goto L324
	}
L312:
	;
	goto L311
L313:
	;
	v1503 = *(*int32)(unsafe.Add(mBase, uint32(v1491)+4))
	v1504 = *(*int32)(unsafe.Add(mBase, uint32(v1492)+4))
	if v1503 < v1504 {
		goto L314
	} else {
		goto L315
	}
L314:
	;
	v1506 = v1503
	goto L316
L315:
	;
	v1506 = v1504
	goto L316
L316:
	;
	if v1506 <= int32(1) {
		goto L317
	} else {
		goto L318
	}
L317:
	;
	v1509 = int32(1)
	goto L319
L318:
	;
	v1509 = v1506
	goto L319
L319:
	;
	v1510 = int32(8)
	v1515 = int32(0)
	goto L320
L320:
	;
	v1522 = v1515 << (uint(int32(2)) % 32)
	v1524 = *(*int32)(unsafe.Add(mBase, uint32(v1492+v1510+v1522)))
	v1526 = *(*int32)(unsafe.Add(mBase, uint32(v1491+v1510+v1522)))
	v1527 = v1524 & v1526
	v1529 = base.B2i32(v1527 != int32(0))
	if v1527 != 0 {
		v1538 = v1529
		goto L312
	} else {
		goto L322
	}
L321:
	;
	v1538 = v1529
	goto L312
L322:
	;
	v1531 = v1515 + int32(1)
	if v1531 != v1509 {
		v1515 = v1531
		goto L320
	} else {
		goto L323
	}
L323:
	;
	goto L321
L324:
	;
	v1539 = *(*int32)(unsafe.Add(mBase, uint32(v1383)+16))
	if v1539 == int32(0) {
		goto L309
	} else {
		goto L325
	}
L325:
	;
	v1542 = *(*int32)(unsafe.Add(mBase, uint32(v1539)+4))
	v1543 = *(*int32)(unsafe.Add(mBase, uint32(l2)+228))
	v1544 = int32(0)
	if base.B2i32(v1542 == v1544)|base.B2i32(v1543 == v1544) != 0 {
		v1589 = v1544
		goto L327
	} else {
		goto L328
	}
L326:
	;
	if v1589 != 0 {
		goto L276
	} else {
		goto L339
	}
L327:
	;
	goto L326
L328:
	;
	v1554 = *(*int32)(unsafe.Add(mBase, uint32(v1542)+4))
	v1555 = *(*int32)(unsafe.Add(mBase, uint32(v1543)+4))
	if v1554 < v1555 {
		goto L329
	} else {
		goto L330
	}
L329:
	;
	v1557 = v1554
	goto L331
L330:
	;
	v1557 = v1555
	goto L331
L331:
	;
	if v1557 <= int32(1) {
		goto L332
	} else {
		goto L333
	}
L332:
	;
	v1560 = int32(1)
	goto L334
L333:
	;
	v1560 = v1557
	goto L334
L334:
	;
	v1561 = int32(8)
	v1566 = int32(0)
	goto L335
L335:
	;
	v1573 = v1566 << (uint(int32(2)) % 32)
	v1575 = *(*int32)(unsafe.Add(mBase, uint32(v1543+v1561+v1573)))
	v1577 = *(*int32)(unsafe.Add(mBase, uint32(v1542+v1561+v1573)))
	v1578 = v1575 & v1577
	v1580 = base.B2i32(v1578 != int32(0))
	if v1578 != 0 {
		v1589 = v1580
		goto L327
	} else {
		goto L337
	}
L336:
	;
	v1589 = v1580
	goto L327
L337:
	;
	v1582 = v1566 + int32(1)
	if v1582 != v1560 {
		v1566 = v1582
		goto L335
	} else {
		goto L338
	}
L338:
	;
	goto L336
L339:
	;
	goto L309
L340:
	;
	v1604 = int32(0)
	v1605 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l1)+26)))
	if v1605 != int32(1) {
		v1678 = v1604
		v1679 = v8
		goto L345
	} else {
		goto L346
	}
L341:
	;
	v1598 = *(*int32)(unsafe.Add(mBase, uint32(v39)+20))
	v1599 = F_create_unique_path(m, l0, l3, v1383, v1598)
	mBase = m.M
	v1600 = m.ExcPending
	if v1600 != 0 {
		goto L22
	} else {
		goto L344
	}
L342:
	;
	v1594 = *(*int32)(unsafe.Add(mBase, uint32(v39)+20))
	v1595 = F_create_unique_path(m, l0, l2, v1384, v1594)
	mBase = m.M
	v1596 = m.ExcPending
	if v1596 != 0 {
		goto L22
	} else {
		goto L343
	}
L343:
	;
	v1601 = v1383
	v1602 = int32(0)
	v1603 = v1595
	goto L340
L344:
	;
	v1601 = v1599
	v1602 = int32(0)
	v1603 = v1384
	goto L340
L345:
	;
	v1680 = int32(0)
	v1681 = *(*int32)(unsafe.Add(mBase, uint32(v39)+12))
	if v1681 == v1680 {
		v2847 = v8
		goto L371
	} else {
		goto L372
	}
L346:
	;
	v1609 = l4 - int32(2)
	v1616 = int32(0)
	if base.B2i32(base.Ui32(int32(6)) < base.Ui32(v1609))|base.B2i32(int32(base.Ui32(int32(99))>>(uint(v1609)%32))&int32(1) == v1616) == v1616 {
		v1678 = v1604
		v1679 = v8
		goto L345
	} else {
		goto L347
	}
L347:
	;
	v1621 = *(*int32)(unsafe.Add(mBase, uint32(l2)+40))
	if v1621 == int32(0) {
		v1678 = v1604
		v1679 = v8
		goto L345
	} else {
		goto L348
	}
L348:
	;
	v1624 = *(*int32)(unsafe.Add(mBase, uint32(l1)+64))
	if v1624 != 0 {
		v1678 = v1604
		v1679 = v8
		goto L345
	} else {
		goto L349
	}
L349:
	;
	v1625 = *(*int32)(unsafe.Add(mBase, uint32(v1621)+12))
	v1626 = *(*int32)(unsafe.Add(mBase, uint32(v1625)))
	v1627 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1601)+21)))
	if v1627 != 0 {
		goto L350
	} else {
		goto L351
	}
L350:
	;
	v1678 = v1601
	v1679 = v1626
	goto L345
L351:
	;
	goto L352
L352:
	;
	if l4 == int32(9) {
		v1678 = v1604
		v1679 = v1626
		goto L345
	} else {
		goto L353
	}
L353:
	;
	v1630 = *(*int32)(unsafe.Add(mBase, uint32(l3)+32))
	if v1630 != 0 {
		goto L356
	} else {
		goto L357
	}
L354:
	;
	v1678 = v1674
	v1679 = v1626
	goto L345
L355:
	;
	goto L354
L356:
	;
	v1635 = *(*int32)(unsafe.Add(mBase, uint32(v1630)+4))
	if v1635 <= int32(0) {
		v1674 = int32(0)
		goto L355
	} else {
		goto L359
	}
L357:
	;
	goto L358
L358:
	;
	v1674 = int32(0)
	goto L355
L359:
	;
	v1638 = int32(0)
	if v1638 < v1635 {
		goto L360
	} else {
		goto L361
	}
L360:
	;
	v1641 = v1635
	goto L362
L361:
	;
	v1641 = v1638
	goto L362
L362:
	;
	v1642 = *(*int32)(unsafe.Add(mBase, uint32(v1630)+12))
	v1644 = int32(0)
	goto L363
L363:
	;
	v1652 = *(*int32)(unsafe.Add(mBase, uint32(v1642+v1644<<(uint(int32(2))%32))))
	v1653 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1652)+21)))
	if v1653 == int32(1) {
		goto L365
	} else {
		goto L366
	}
L364:
	;
	goto L358
L365:
	;
	v1656 = *(*int32)(unsafe.Add(mBase, uint32(v1652)+16))
	if v1656 == int32(0) {
		v1674 = v1652
		goto L355
	} else {
		goto L368
	}
L366:
	;
	goto L367
L367:
	;
	v1664 = v1644 + int32(1)
	if v1664 != v1641 {
		v1644 = v1664
		goto L363
	} else {
		goto L370
	}
L368:
	;
	v1659 = *(*int32)(unsafe.Add(mBase, uint32(v1656)+4))
	if v1659 == int32(0) {
		v1674 = v1652
		goto L355
	} else {
		goto L369
	}
L369:
	;
	goto L367
L370:
	;
	goto L364
L371:
	;
	if v2847 == int32(0) {
		goto L276
	} else {
		goto L513
	}
L372:
	;
	v1684 = *(*int32)(unsafe.Add(mBase, uint32(v1681)+4))
	if v1684 == int32(0) {
		v2847 = v8
		goto L371
	} else {
		goto L373
	}
L373:
	;
	v1688 = v1684 << (uint(int32(2)) % 32)
	v1689 = F_palloc(m, v1688)
	mBase = m.M
	v1690 = m.ExcPending
	if v1690 != 0 {
		goto L22
	} else {
		goto L374
	}
L374:
	;
	v1691 = F_palloc(m, v1688)
	mBase = m.M
	v1692 = m.ExcPending
	if v1692 != 0 {
		goto L22
	} else {
		goto L375
	}
L375:
	;
	v1693 = *(*int32)(unsafe.Add(mBase, uint32(v1681)+4))
	if int32(0) < v1693 {
		goto L376
	} else {
		goto L377
	}
L376:
	;
	v1702 = v1680
	v1712 = v8
	goto L379
L377:
	;
	v2177 = v1680
	goto L378
L378:
	;
	v2207 = *(*int32)(unsafe.Add(mBase, uint32(l0)+156))
	if v2207 == int32(0) {
		v2499 = v8
		goto L430
	} else {
		goto L431
	}
L379:
	;
	v1732 = *(*int32)(unsafe.Add(mBase, uint32(v1681)+12))
	v1736 = *(*int32)(unsafe.Add(mBase, uint32(v1732+v1712<<(uint(int32(2))%32))))
	v1737 = *(*int32)(unsafe.Add(mBase, uint32(v1736)+100))
	v1738 = *(*int32)(unsafe.Add(mBase, uint32(v1737)+56))
	if v1738 != 0 {
		goto L381
	} else {
		goto L382
	}
L380:
	;
	v2177 = v2137
	goto L378
L381:
	;
	v1747 = v1738
	goto L384
L382:
	;
	goto L383
L383:
	;
	v1813 = *(*int32)(unsafe.Add(mBase, uint32(v1736)+104))
	v1814 = *(*int32)(unsafe.Add(mBase, uint32(v1813)+56))
	if v1814 != 0 {
		goto L387
	} else {
		goto L388
	}
L384:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v1736)+100)) = v1747
	v1776 = *(*int32)(unsafe.Add(mBase, uint32(v1747)+56))
	if v1776 != 0 {
		v1747 = v1776
		goto L384
	} else {
		goto L386
	}
L385:
	;
	goto L383
L386:
	;
	goto L385
L387:
	;
	v1823 = v1814
	goto L390
L388:
	;
	goto L389
L389:
	;
	v1891 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1736)+120)))
	if v1891 != 0 {
		goto L393
	} else {
		goto L394
	}
L390:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v1736)+104)) = v1823
	v1852 = *(*int32)(unsafe.Add(mBase, uint32(v1823)+56))
	if v1852 != 0 {
		v1823 = v1852
		goto L390
	} else {
		goto L392
	}
L391:
	;
	goto L389
L392:
	;
	goto L391
L393:
	;
	v1892 = int32(100)
	goto L395
L394:
	;
	v1892 = int32(104)
	goto L395
L395:
	;
	v1894 = *(*int32)(unsafe.Add(mBase, uint32(v1736+v1892)))
	v1895 = int32(0)
	if v1895 < v1702 {
		goto L397
	} else {
		goto L398
	}
L396:
	;
	v2168 = v1712 + int32(1)
	v2169 = *(*int32)(unsafe.Add(mBase, uint32(v1681)+4))
	if v2168 < v2169 {
		v1702 = v2137
		v1712 = v2168
		goto L379
	} else {
		goto L428
	}
L397:
	;
	v1906 = v1895
	goto L400
L398:
	;
	goto L399
L399:
	;
	v1978 = *(*int32)(unsafe.Add(mBase, uint32(v1894)+16))
	if v1978 == int32(0) {
		goto L405
	} else {
		goto L406
	}
L400:
	;
	v1937 = *(*int32)(unsafe.Add(mBase, uint32(v1689+v1906<<(uint(int32(2))%32))))
	if v1937 == v1894 {
		v2137 = v1702
		goto L396
	} else {
		goto L402
	}
L401:
	;
	goto L399
L402:
	;
	v1940 = v1906 + int32(1)
	if v1940 != v1702 {
		v1906 = v1940
		goto L400
	} else {
		goto L403
	}
L403:
	;
	goto L401
L404:
	;
	v2124 = v1702 << (uint(int32(2)) % 32)
	*(*int32)(unsafe.Add(mBase, uint32(v1689+v2124))) = v1894
	*(*int32)(unsafe.Add(mBase, uint32(v2124+v1691))) = v2094
	v2137 = v1702 + int32(1)
	goto L396
L405:
	;
	v2094 = int32(0)
	goto L404
L406:
	;
	goto L407
L407:
	;
	v1982 = int32(0)
	v1984 = *(*int32)(unsafe.Add(mBase, uint32(v1978)+4))
	if v1984 <= v1982 {
		v2094 = v1982
		goto L404
	} else {
		goto L408
	}
L408:
	;
	v1994 = v1982
	v1995 = v1982
	goto L409
L409:
	;
	v2023 = *(*int32)(unsafe.Add(mBase, uint32(v1978)+12))
	v2027 = *(*int32)(unsafe.Add(mBase, uint32(v2023+v1995<<(uint(int32(2))%32))))
	v2028 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v2027)+12)))
	if v2028 == int32(0) {
		goto L411
	} else {
		goto L412
	}
L410:
	;
	v2094 = v2082
	goto L404
L411:
	;
	v2031 = *(*int32)(unsafe.Add(mBase, uint32(v2027)+8))
	v2032 = *(*int32)(unsafe.Add(mBase, uint32(l1)+8))
	v2033 = int32(0)
	if base.B2i32(v2031 == v2033)|base.B2i32(v2032 == v2033) != 0 {
		v2078 = v2033
		goto L415
	} else {
		goto L416
	}
L412:
	;
	v2082 = v1994
	goto L413
L413:
	;
	v2084 = v1995 + int32(1)
	v2085 = *(*int32)(unsafe.Add(mBase, uint32(v1978)+4))
	if v2084 < v2085 {
		v1994 = v2082
		v1995 = v2084
		goto L409
	} else {
		goto L427
	}
L414:
	;
	v2082 = v1994 + (v2078 ^ int32(1))
	goto L413
L415:
	;
	goto L414
L416:
	;
	v2043 = *(*int32)(unsafe.Add(mBase, uint32(v2031)+4))
	v2044 = *(*int32)(unsafe.Add(mBase, uint32(v2032)+4))
	if v2043 < v2044 {
		goto L417
	} else {
		goto L418
	}
L417:
	;
	v2046 = v2043
	goto L419
L418:
	;
	v2046 = v2044
	goto L419
L419:
	;
	if v2046 <= int32(1) {
		goto L420
	} else {
		goto L421
	}
L420:
	;
	v2049 = int32(1)
	goto L422
L421:
	;
	v2049 = v2046
	goto L422
L422:
	;
	v2050 = int32(8)
	v2055 = int32(0)
	goto L423
L423:
	;
	v2062 = v2055 << (uint(int32(2)) % 32)
	v2064 = *(*int32)(unsafe.Add(mBase, uint32(v2032+v2050+v2062)))
	v2066 = *(*int32)(unsafe.Add(mBase, uint32(v2031+v2050+v2062)))
	v2067 = v2064 & v2066
	v2069 = base.B2i32(v2067 != int32(0))
	if v2067 != 0 {
		v2078 = v2069
		goto L415
	} else {
		goto L425
	}
L424:
	;
	v2078 = v2069
	goto L415
L425:
	;
	v2071 = v2055 + int32(1)
	if v2071 != v2049 {
		v2055 = v2071
		goto L423
	} else {
		goto L426
	}
L426:
	;
	goto L424
L427:
	;
	goto L410
L428:
	;
	goto L380
L429:
	;
	F_pfree(m, v1689)
	mBase = m.M
	v2831 = m.ExcPending
	if v2831 != 0 {
		goto L22
	} else {
		goto L511
	}
L430:
	;
	v2523 = v2177 - int32(1)
	v2526 = int32(3)
	v2527 = v2523 & v2526
	v2528 = int32(2)
	v2547 = v2499
	goto L462
L431:
	;
	v2210 = *(*int32)(unsafe.Add(mBase, uint32(v2207)+4))
	if int32(0) < v2210 {
		goto L433
	} else {
		goto L434
	}
L432:
	;
	if v1684 != v2233 {
		v2499 = v8
		goto L430
	} else {
		goto L460
	}
L433:
	;
	v2213 = *(*int32)(unsafe.Add(mBase, uint32(v2207)+12))
	v2233 = int32(0)
	goto L436
L434:
	;
	goto L435
L435:
	;
	v2342 = F_list_copy(m, v2207)
	mBase = m.M
	v2343 = m.ExcPending
	if v2343 != 0 {
		goto L22
	} else {
		goto L446
	}
L436:
	;
	if v2177 <= int32(0) {
		v2499 = v8
		goto L430
	} else {
		goto L438
	}
L437:
	;
	goto L435
L438:
	;
	v2256 = *(*int32)(unsafe.Add(mBase, uint32(v2213+v2233<<(uint(int32(2))%32))))
	v2257 = *(*int32)(unsafe.Add(mBase, uint32(v2256)+4))
	v2267 = int32(0)
	goto L439
L439:
	;
	v2298 = *(*int32)(unsafe.Add(mBase, uint32(v1689+v2267<<(uint(int32(2))%32))))
	if v2257 != v2298 {
		goto L441
	} else {
		goto L442
	}
L440:
	;
	v2304 = v2233 + int32(1)
	if v2304 != v2210 {
		v2233 = v2304
		goto L436
	} else {
		goto L445
	}
L441:
	;
	v2301 = v2267 + int32(1)
	if v2177 != v2301 {
		v2267 = v2301
		goto L439
	} else {
		goto L444
	}
L442:
	;
	goto L443
L443:
	;
	goto L440
L444:
	;
	goto L432
L445:
	;
	goto L437
L446:
	;
	v2344 = *(*int32)(unsafe.Add(mBase, uint32(l0)+156))
	if v2344 == int32(0) {
		v2499 = v2342
		goto L430
	} else {
		goto L447
	}
L447:
	;
	v2347 = *(*int32)(unsafe.Add(mBase, uint32(v2344)+4))
	if v2347 <= int32(0) {
		v2499 = v2342
		goto L430
	} else {
		goto L448
	}
L448:
	;
	v2350 = int32(0)
	v2371 = v2350
	goto L449
L449:
	;
	if v2177 <= v2350 {
		goto L451
	} else {
		goto L452
	}
L450:
	;
	v2499 = v2342
	goto L430
L451:
	;
	v2480 = v2371 + int32(1)
	v2481 = *(*int32)(unsafe.Add(mBase, uint32(v2344)+4))
	if v2480 < v2481 {
		v2371 = v2480
		goto L449
	} else {
		goto L459
	}
L452:
	;
	v2389 = *(*int32)(unsafe.Add(mBase, uint32(v2344)+12))
	v2393 = *(*int32)(unsafe.Add(mBase, uint32(v2389+v2371<<(uint(int32(2))%32))))
	v2394 = *(*int32)(unsafe.Add(mBase, uint32(v2393)+4))
	v2404 = int32(0)
	goto L453
L453:
	;
	v2433 = v2404 << (uint(int32(2)) % 32)
	v2435 = *(*int32)(unsafe.Add(mBase, uint32(v1689+v2433)))
	if v2394 == v2435 {
		goto L455
	} else {
		goto L456
	}
L454:
	;
	goto L451
L455:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v2433+v1691))) = int32(-1)
	goto L451
L456:
	;
	goto L457
L457:
	;
	v2441 = v2404 + int32(1)
	if v2441 != v2177 {
		v2404 = v2441
		goto L453
	} else {
		goto L458
	}
L458:
	;
	goto L454
L459:
	;
	goto L450
L460:
	;
	v2484 = F_list_copy_head(m, v2207, v1684)
	mBase = m.M
	v2485 = m.ExcPending
	if v2485 != 0 {
		goto L22
	} else {
		goto L461
	}
L461:
	;
	v2807 = v2484
	goto L429
L462:
	;
	v2570 = *(*int32)(unsafe.Add(mBase, uint32(v1691)))
	if v2177 < v2528 {
		goto L465
	} else {
		goto L466
	}
L464:
	;
	if v2754 < int32(0) {
		v2807 = v2547
		goto L429
	} else {
		goto L508
	}
L465:
	;
	v2746 = int32(0)
	v2754 = v2570
	goto L464
L466:
	;
	goto L467
L467:
	;
	v2572 = int32(1)
	v2573 = int32(0)
	if base.B2i32(base.Ui32(v2177-v2528) < base.Ui32(v2526)) == v2573 {
		goto L468
	} else {
		goto L469
	}
L468:
	;
	v2585 = v2573
	v2587 = v2572
	v2593 = v2570
	v2597 = v2573
	goto L471
L469:
	;
	v2663 = v2573
	v2664 = v2572
	v2670 = v2570
	goto L470
L470:
	;
	v2698 = v2663
	v2700 = v2664
	v2706 = v2670
	v2718 = v2573
	goto L499
L471:
	;
	v2616 = v2587 + int32(3)
	v2617 = int32(2)
	v2620 = *(*int32)(unsafe.Add(mBase, uint32(v1691+v2616<<(uint(v2617)%32))))
	v2622 = v2587 + v2617
	v2626 = *(*int32)(unsafe.Add(mBase, uint32(v1691+v2622<<(uint(v2617)%32))))
	v2628 = v2587 + int32(1)
	v2632 = *(*int32)(unsafe.Add(mBase, uint32(v1691+v2628<<(uint(v2617)%32))))
	v2636 = *(*int32)(unsafe.Add(mBase, uint32(v1691+v2587<<(uint(v2617)%32))))
	v2637 = base.B2i32(v2593 < v2636)
	if v2593 < v2636 {
		goto L473
	} else {
		goto L474
	}
L472:
	;
	if v2527 == int32(0) {
		v2746 = v2648
		v2754 = v2644
		goto L464
	} else {
		goto L498
	}
L473:
	;
	v2638 = v2636
	goto L475
L474:
	;
	v2638 = v2593
	goto L475
L475:
	;
	v2639 = base.B2i32(v2638 < v2632)
	if v2638 < v2632 {
		goto L476
	} else {
		goto L477
	}
L476:
	;
	v2640 = v2632
	goto L478
L477:
	;
	v2640 = v2638
	goto L478
L478:
	;
	v2641 = base.B2i32(v2640 < v2626)
	if v2640 < v2626 {
		goto L479
	} else {
		goto L480
	}
L479:
	;
	v2642 = v2626
	goto L481
L480:
	;
	v2642 = v2640
	goto L481
L481:
	;
	v2643 = base.B2i32(v2642 < v2620)
	if v2642 < v2620 {
		goto L482
	} else {
		goto L483
	}
L482:
	;
	v2644 = v2620
	goto L484
L483:
	;
	v2644 = v2642
	goto L484
L484:
	;
	if v2593 < v2636 {
		goto L485
	} else {
		goto L486
	}
L485:
	;
	v2645 = v2587
	goto L487
L486:
	;
	v2645 = v2585
	goto L487
L487:
	;
	if v2638 < v2632 {
		goto L488
	} else {
		goto L489
	}
L488:
	;
	v2646 = v2628
	goto L490
L489:
	;
	v2646 = v2645
	goto L490
L490:
	;
	if v2640 < v2626 {
		goto L491
	} else {
		goto L492
	}
L491:
	;
	v2647 = v2622
	goto L493
L492:
	;
	v2647 = v2646
	goto L493
L493:
	;
	if v2642 < v2620 {
		goto L494
	} else {
		goto L495
	}
L494:
	;
	v2648 = v2616
	goto L496
L495:
	;
	v2648 = v2647
	goto L496
L496:
	;
	v2649 = int32(4)
	v2650 = v2587 + v2649
	v2652 = v2597 + v2649
	if v2652 != v2523&int32(-4) {
		v2585 = v2648
		v2587 = v2650
		v2593 = v2644
		v2597 = v2652
		goto L471
	} else {
		goto L497
	}
L497:
	;
	goto L472
L498:
	;
	v2663 = v2648
	v2664 = v2650
	v2670 = v2644
	goto L470
L499:
	;
	v2731 = *(*int32)(unsafe.Add(mBase, uint32(v1691+v2700<<(uint(int32(2))%32))))
	v2732 = base.B2i32(v2706 < v2731)
	if v2706 < v2731 {
		goto L501
	} else {
		goto L502
	}
L500:
	;
	v2746 = v2734
	v2754 = v2733
	goto L464
L501:
	;
	v2733 = v2731
	goto L503
L502:
	;
	v2733 = v2706
	goto L503
L503:
	;
	if v2706 < v2731 {
		goto L504
	} else {
		goto L505
	}
L504:
	;
	v2734 = v2700
	goto L506
L505:
	;
	v2734 = v2698
	goto L506
L506:
	;
	v2735 = int32(1)
	v2738 = v2718 + v2735
	if v2738 != v2527 {
		v2698 = v2734
		v2700 = v2700 + v2735
		v2706 = v2733
		v2718 = v2738
		goto L499
	} else {
		goto L507
	}
L507:
	;
	goto L500
L508:
	;
	v2779 = v2746 << (uint(int32(2)) % 32)
	v2781 = *(*int32)(unsafe.Add(mBase, uint32(v1689+v2779)))
	*(*int32)(unsafe.Add(mBase, uint32(v2779+v1691))) = int32(-1)
	v2785 = *(*int32)(unsafe.Add(mBase, uint32(v2781)+4))
	v2786 = *(*int32)(unsafe.Add(mBase, uint32(v2785)+12))
	v2787 = *(*int32)(unsafe.Add(mBase, uint32(v2786)))
	v2790 = F_make_canonical_pathkey(m, l0, v2781, v2787, int32(1), int32(0))
	mBase = m.M
	v2791 = m.ExcPending
	if v2791 != 0 {
		goto L22
	} else {
		goto L509
	}
L509:
	;
	v2792 = F_lappend(m, v2547, v2790)
	mBase = m.M
	v2793 = m.ExcPending
	if v2793 != 0 {
		goto L22
	} else {
		goto L510
	}
L510:
	;
	v2547 = v2792
	goto L462
L511:
	;
	F_pfree(m, v1691)
	mBase = m.M
	v2833 = m.ExcPending
	if v2833 != 0 {
		goto L22
	} else {
		goto L512
	}
L512:
	;
	v2847 = v2807
	goto L371
L513:
	;
	v2872 = *(*int32)(unsafe.Add(mBase, uint32(v2847)+4))
	if v2872 <= int32(0) {
		goto L276
	} else {
		goto L514
	}
L514:
	;
	v2875 = *(*int32)(unsafe.Add(mBase, uint32(v39)+12))
	v2876 = F_find_mergeclauses_for_outer_pathkeys(m, v2847, v2875)
	mBase = m.M
	v2877 = m.ExcPending
	if v2877 != 0 {
		goto L22
	} else {
		goto L515
	}
L515:
	;
	v2878 = F_make_inner_pathkeys_for_merge(m, l0, v2876, v2847)
	mBase = m.M
	v2879 = m.ExcPending
	if v2879 != 0 {
		goto L22
	} else {
		goto L516
	}
L516:
	;
	v2880 = F_build_join_pathkeys(m, l0, l1, v1602, v2847)
	mBase = m.M
	v2881 = m.ExcPending
	if v2881 != 0 {
		goto L22
	} else {
		goto L517
	}
L517:
	;
	v2883 = v39 + int32(8)
	F_try_mergejoin_path(m, l0, l1, v1603, v1601, v2880, v2876, v2847, v2878, v1602, v2883, int32(0))
	mBase = m.M
	v2886 = m.ExcPending
	if v2886 != 0 {
		goto L22
	} else {
		goto L518
	}
L518:
	;
	v2887 = int32(0)
	v2891 = base.B2i32(v1679 != v2887) & base.B2i32(v1678 != v2887)
	if v2891 != 0 {
		goto L519
	} else {
		goto L520
	}
L519:
	;
	F_try_partial_mergejoin_path(m, l0, l1, v1679, v1678, v2880, v2876, v2847, v2878, v1602, v2883)
	mBase = m.M
	v2893 = m.ExcPending
	if v2893 != 0 {
		goto L22
	} else {
		goto L522
	}
L520:
	;
	goto L521
L521:
	;
	v2894 = *(*int32)(unsafe.Add(mBase, uint32(v2847)+4))
	if v2894 < int32(2) {
		goto L276
	} else {
		goto L523
	}
L522:
	;
	goto L521
L523:
	;
	v2904 = int32(1)
	goto L524
L524:
	;
	v2934 = *(*int32)(unsafe.Add(mBase, uint32(v2847)+12))
	v2938 = *(*int32)(unsafe.Add(mBase, uint32(v2934+v2904<<(uint(int32(2))%32))))
	v2939 = F_list_copy(m, v2847)
	mBase = m.M
	v2940 = m.ExcPending
	if v2940 != 0 {
		goto L22
	} else {
		goto L526
	}
L525:
	;
	goto L276
L526:
	;
	v2941 = F_list_delete_nth_cell(m, v2939, v2904)
	mBase = m.M
	v2942 = m.ExcPending
	if v2942 != 0 {
		goto L22
	} else {
		goto L527
	}
L527:
	;
	v2943 = F_lcons(m, v2938, v2941)
	mBase = m.M
	v2944 = m.ExcPending
	if v2944 != 0 {
		goto L22
	} else {
		goto L528
	}
L528:
	;
	v2945 = *(*int32)(unsafe.Add(mBase, uint32(v39)+12))
	v2946 = F_find_mergeclauses_for_outer_pathkeys(m, v2943, v2945)
	mBase = m.M
	v2947 = m.ExcPending
	if v2947 != 0 {
		goto L22
	} else {
		goto L529
	}
L529:
	;
	v2948 = F_make_inner_pathkeys_for_merge(m, l0, v2946, v2943)
	mBase = m.M
	v2949 = m.ExcPending
	if v2949 != 0 {
		goto L22
	} else {
		goto L530
	}
L530:
	;
	v2950 = F_build_join_pathkeys(m, l0, l1, v1602, v2943)
	mBase = m.M
	v2951 = m.ExcPending
	if v2951 != 0 {
		goto L22
	} else {
		goto L531
	}
L531:
	;
	v2953 = v39 + int32(8)
	F_try_mergejoin_path(m, l0, l1, v1603, v1601, v2950, v2946, v2943, v2948, v1602, v2953, int32(0))
	mBase = m.M
	v2956 = m.ExcPending
	if v2956 != 0 {
		goto L22
	} else {
		goto L532
	}
L532:
	;
	if v2891 != 0 {
		goto L533
	} else {
		goto L534
	}
L533:
	;
	F_try_partial_mergejoin_path(m, l0, l1, v1679, v1678, v2950, v2946, v2943, v2948, v1602, v2953)
	mBase = m.M
	v2958 = m.ExcPending
	if v2958 != 0 {
		goto L22
	} else {
		goto L536
	}
L534:
	;
	goto L535
L535:
	;
	v2960 = v2904 + int32(1)
	v2961 = *(*int32)(unsafe.Add(mBase, uint32(v2847)+4))
	if v2960 < v2961 {
		v2904 = v2960
		goto L524
	} else {
		goto L537
	}
L536:
	;
	goto L535
L537:
	;
	goto L525
L538:
	;
	v3020 = *(*int32)(unsafe.Add(mBase, uint32(v2999)+16))
	if v3020 == int32(0) {
		goto L548
	} else {
		goto L549
	}
L539:
	;
	v3017 = l4
	v3018 = int32(0)
	v3019 = int32(1)
	goto L538
L540:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v3005 = m.ExcPending
	if v3005 != 0 {
		goto L22
	} else {
		goto L542
	}
L541:
	;
	v3017 = int32(0)
	v3018 = v3000
	v3019 = v8
	goto L538
L542:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v39))) = l4
	F_errmsg_internal(m, int32(_a_F_add_paths_to_joinrel_0), v39)
	mBase = m.M
	v3009 = m.ExcPending
	if v3009 != 0 {
		goto L22
	} else {
		goto L543
	}
L543:
	;
	F_errfinish(m, int32(_a_F_add_paths_to_joinrel_1), int32(1863), int32(_a_F_add_paths_to_joinrel_2))
	mBase = m.M
	v3014 = m.ExcPending
	if v3014 != 0 {
		goto L22
	} else {
		goto L544
	}
L544:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L545:
	;
	v3159 = *(*int32)(unsafe.Add(mBase, uint32(l2)+32))
	if v3159 == int32(0) {
		goto L587
	} else {
		goto L588
	}
L546:
	;
	v3135 = int32(0)
	if v3018 == v3135 {
		v3156 = v3134
		v3157 = v3135
		v3158 = v8
		goto L545
	} else {
		goto L582
	}
L547:
	;
	if l4 == int32(9) {
		goto L274
	} else {
		goto L581
	}
L548:
	;
	if l4 != int32(9) {
		v3134 = v2999
		goto L546
	} else {
		goto L579
	}
L549:
	;
	v3023 = *(*int32)(unsafe.Add(mBase, uint32(v3020)+4))
	v3024 = *(*int32)(unsafe.Add(mBase, uint32(l2)+8))
	v3025 = int32(0)
	if base.B2i32(v3023 == v3025)|base.B2i32(v3024 == v3025) != 0 {
		v3070 = v3025
		goto L551
	} else {
		goto L552
	}
L550:
	;
	if v3070 != 0 {
		goto L547
	} else {
		goto L563
	}
L551:
	;
	goto L550
L552:
	;
	v3035 = *(*int32)(unsafe.Add(mBase, uint32(v3023)+4))
	v3036 = *(*int32)(unsafe.Add(mBase, uint32(v3024)+4))
	if v3035 < v3036 {
		goto L553
	} else {
		goto L554
	}
L553:
	;
	v3038 = v3035
	goto L555
L554:
	;
	v3038 = v3036
	goto L555
L555:
	;
	if v3038 <= int32(1) {
		goto L556
	} else {
		goto L557
	}
L556:
	;
	v3041 = int32(1)
	goto L558
L557:
	;
	v3041 = v3038
	goto L558
L558:
	;
	v3042 = int32(8)
	v3047 = int32(0)
	goto L559
L559:
	;
	v3054 = v3047 << (uint(int32(2)) % 32)
	v3056 = *(*int32)(unsafe.Add(mBase, uint32(v3024+v3042+v3054)))
	v3058 = *(*int32)(unsafe.Add(mBase, uint32(v3023+v3042+v3054)))
	v3059 = v3056 & v3058
	v3061 = base.B2i32(v3059 != int32(0))
	if v3059 != 0 {
		v3070 = v3061
		goto L551
	} else {
		goto L561
	}
L560:
	;
	v3070 = v3061
	goto L551
L561:
	;
	v3063 = v3047 + int32(1)
	if v3063 != v3041 {
		v3047 = v3063
		goto L559
	} else {
		goto L562
	}
L562:
	;
	goto L560
L563:
	;
	v3071 = *(*int32)(unsafe.Add(mBase, uint32(v2999)+16))
	if v3071 == int32(0) {
		goto L548
	} else {
		goto L564
	}
L564:
	;
	v3074 = *(*int32)(unsafe.Add(mBase, uint32(v3071)+4))
	v3075 = *(*int32)(unsafe.Add(mBase, uint32(l2)+228))
	v3076 = int32(0)
	if base.B2i32(v3074 == v3076)|base.B2i32(v3075 == v3076) != 0 {
		v3121 = v3076
		goto L566
	} else {
		goto L567
	}
L565:
	;
	if v3121 != 0 {
		goto L547
	} else {
		goto L578
	}
L566:
	;
	goto L565
L567:
	;
	v3086 = *(*int32)(unsafe.Add(mBase, uint32(v3074)+4))
	v3087 = *(*int32)(unsafe.Add(mBase, uint32(v3075)+4))
	if v3086 < v3087 {
		goto L568
	} else {
		goto L569
	}
L568:
	;
	v3089 = v3086
	goto L570
L569:
	;
	v3089 = v3087
	goto L570
L570:
	;
	if v3089 <= int32(1) {
		goto L571
	} else {
		goto L572
	}
L571:
	;
	v3092 = int32(1)
	goto L573
L572:
	;
	v3092 = v3089
	goto L573
L573:
	;
	v3093 = int32(8)
	v3098 = int32(0)
	goto L574
L574:
	;
	v3105 = v3098 << (uint(int32(2)) % 32)
	v3107 = *(*int32)(unsafe.Add(mBase, uint32(v3075+v3093+v3105)))
	v3109 = *(*int32)(unsafe.Add(mBase, uint32(v3074+v3093+v3105)))
	v3110 = v3107 & v3109
	v3112 = base.B2i32(v3110 != int32(0))
	if v3110 != 0 {
		v3121 = v3112
		goto L566
	} else {
		goto L576
	}
L575:
	;
	v3121 = v3112
	goto L566
L576:
	;
	v3114 = v3098 + int32(1)
	if v3114 != v3092 {
		v3098 = v3114
		goto L574
	} else {
		goto L577
	}
L577:
	;
	goto L575
L578:
	;
	goto L548
L579:
	;
	v3126 = *(*int32)(unsafe.Add(mBase, uint32(v39)+20))
	v3127 = F_create_unique_path(m, l0, l3, v2999, v3126)
	mBase = m.M
	v3128 = m.ExcPending
	if v3128 != 0 {
		goto L22
	} else {
		goto L580
	}
L580:
	;
	v3156 = v3127
	v3157 = int32(1)
	v3158 = v8
	goto L545
L581:
	;
	v3134 = int32(0)
	goto L546
L582:
	;
	v3141 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_add_paths_to_joinrel[1])))
	if base.B2i32(v3134 == int32(0))|base.B2i32(v3141 != int32(1)) != 0 {
		v3156 = v3134
		v3157 = v3135
		v3158 = v8
		goto L545
	} else {
		goto L583
	}
L583:
	;
	v3145 = *(*int32)(unsafe.Add(mBase, uint32(v3134)+4))
	v3147 = v3145 - int32(348)
	goto L584
L584:
	;
	if base.B2i32(base.Ui32(v3147) < base.Ui32(int32(15)))&int32(base.Ui32(int32(_a_F_add_paths_to_joinrel_3))>>(uint(v3147)%32)) != 0 {
		v3156 = v3134
		v3157 = v3135
		v3158 = v8
		goto L545
	} else {
		goto L585
	}
L585:
	;
	v3153 = F_create_material_path(m, l3, v3134)
	mBase = m.M
	v3154 = m.ExcPending
	if v3154 != 0 {
		goto L22
	} else {
		goto L586
	}
L586:
	;
	v3156 = v3134
	v3157 = v3135
	v3158 = v3153
	goto L545
L587:
	;
	v3587 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l1)+26)))
	v3588 = int32(1)
	if v3588<<(uint(l4)%32)&int32(396) != 0 {
		goto L652
	} else {
		goto L653
	}
L588:
	;
	v3162 = *(*int32)(unsafe.Add(mBase, uint32(v3159)+4))
	if v3162 <= int32(0) {
		goto L587
	} else {
		goto L589
	}
L589:
	;
	v3165 = int32(0)
	v3167 = int32(8)
	v3184 = v3165
	goto L590
L590:
	;
	v3209 = *(*int32)(unsafe.Add(mBase, uint32(v3159)+12))
	v3213 = *(*int32)(unsafe.Add(mBase, uint32(v3209+v3184<<(uint(int32(2))%32))))
	v3214 = *(*int32)(unsafe.Add(mBase, uint32(v3213)+16))
	if v3214 == int32(0) {
		goto L593
	} else {
		goto L594
	}
L591:
	;
	goto L587
L592:
	;
	v3548 = v3184 + int32(1)
	v3549 = *(*int32)(unsafe.Add(mBase, uint32(v3159)+4))
	if v3548 < v3549 {
		v3184 = v3548
		goto L590
	} else {
		goto L651
	}
L593:
	;
	if base.B2i32(l4 != v3167) == int32(0) {
		goto L624
	} else {
		goto L625
	}
L594:
	;
	v3217 = *(*int32)(unsafe.Add(mBase, uint32(v3214)+4))
	v3218 = *(*int32)(unsafe.Add(mBase, uint32(l3)+8))
	v3219 = int32(0)
	if base.B2i32(v3217 == v3219)|base.B2i32(v3218 == v3219) != 0 {
		v3264 = v3219
		goto L596
	} else {
		goto L597
	}
L595:
	;
	if v3264 != 0 {
		goto L592
	} else {
		goto L608
	}
L596:
	;
	goto L595
L597:
	;
	v3229 = *(*int32)(unsafe.Add(mBase, uint32(v3217)+4))
	v3230 = *(*int32)(unsafe.Add(mBase, uint32(v3218)+4))
	if v3229 < v3230 {
		goto L598
	} else {
		goto L599
	}
L598:
	;
	v3232 = v3229
	goto L600
L599:
	;
	v3232 = v3230
	goto L600
L600:
	;
	if v3232 <= int32(1) {
		goto L601
	} else {
		goto L602
	}
L601:
	;
	v3235 = int32(1)
	goto L603
L602:
	;
	v3235 = v3232
	goto L603
L603:
	;
	v3236 = int32(8)
	v3241 = int32(0)
	goto L604
L604:
	;
	v3248 = v3241 << (uint(int32(2)) % 32)
	v3250 = *(*int32)(unsafe.Add(mBase, uint32(v3218+v3236+v3248)))
	v3252 = *(*int32)(unsafe.Add(mBase, uint32(v3217+v3236+v3248)))
	v3253 = v3250 & v3252
	v3255 = base.B2i32(v3253 != int32(0))
	if v3253 != 0 {
		v3264 = v3255
		goto L596
	} else {
		goto L606
	}
L605:
	;
	v3264 = v3255
	goto L596
L606:
	;
	v3257 = v3241 + int32(1)
	if v3257 != v3235 {
		v3241 = v3257
		goto L604
	} else {
		goto L607
	}
L607:
	;
	goto L605
L608:
	;
	v3265 = *(*int32)(unsafe.Add(mBase, uint32(v3213)+16))
	if v3265 == int32(0) {
		goto L593
	} else {
		goto L609
	}
L609:
	;
	v3268 = *(*int32)(unsafe.Add(mBase, uint32(v3265)+4))
	v3269 = *(*int32)(unsafe.Add(mBase, uint32(l3)+228))
	v3270 = int32(0)
	if base.B2i32(v3268 == v3270)|base.B2i32(v3269 == v3270) != 0 {
		v3315 = v3270
		goto L611
	} else {
		goto L612
	}
L610:
	;
	if v3315 != 0 {
		goto L592
	} else {
		goto L623
	}
L611:
	;
	goto L610
L612:
	;
	v3280 = *(*int32)(unsafe.Add(mBase, uint32(v3268)+4))
	v3281 = *(*int32)(unsafe.Add(mBase, uint32(v3269)+4))
	if v3280 < v3281 {
		goto L613
	} else {
		goto L614
	}
L613:
	;
	v3283 = v3280
	goto L615
L614:
	;
	v3283 = v3281
	goto L615
L615:
	;
	if v3283 <= int32(1) {
		goto L616
	} else {
		goto L617
	}
L616:
	;
	v3286 = int32(1)
	goto L618
L617:
	;
	v3286 = v3283
	goto L618
L618:
	;
	v3287 = int32(8)
	v3292 = int32(0)
	goto L619
L619:
	;
	v3299 = v3292 << (uint(int32(2)) % 32)
	v3301 = *(*int32)(unsafe.Add(mBase, uint32(v3269+v3287+v3299)))
	v3303 = *(*int32)(unsafe.Add(mBase, uint32(v3268+v3287+v3299)))
	v3304 = v3301 & v3303
	v3306 = base.B2i32(v3304 != int32(0))
	if v3304 != 0 {
		v3315 = v3306
		goto L611
	} else {
		goto L621
	}
L620:
	;
	v3315 = v3306
	goto L611
L621:
	;
	v3308 = v3292 + int32(1)
	if v3308 != v3286 {
		v3292 = v3308
		goto L619
	} else {
		goto L622
	}
L622:
	;
	goto L620
L623:
	;
	goto L593
L624:
	;
	v3319 = *(*int32)(unsafe.Add(mBase, uint32(l2)+48))
	if v3213 != v3319 {
		goto L592
	} else {
		goto L627
	}
L625:
	;
	v3324 = v3213
	goto L626
L626:
	;
	v3325 = *(*int32)(unsafe.Add(mBase, uint32(v3324)+64))
	v3326 = F_build_join_pathkeys(m, l0, l1, v3017, v3325)
	mBase = m.M
	v3327 = m.ExcPending
	if v3327 != 0 {
		goto L22
	} else {
		goto L629
	}
L627:
	;
	v3321 = *(*int32)(unsafe.Add(mBase, uint32(v39)+20))
	v3322 = F_create_unique_path(m, l0, l2, v3213, v3321)
	mBase = m.M
	v3323 = m.ExcPending
	if v3323 != 0 {
		goto L22
	} else {
		goto L628
	}
L628:
	;
	v3324 = v3322
	goto L626
L629:
	;
	if v3157 == int32(0) {
		goto L631
	} else {
		goto L632
	}
L630:
	;
	if base.B2i32(v3156 == v3165)|base.B2i32(l4 == v3167) != 0 {
		goto L592
	} else {
		goto L649
	}
L631:
	;
	if v3018 == int32(0) {
		goto L630
	} else {
		goto L634
	}
L632:
	;
	v3437 = v3156
	goto L633
L633:
	;
	F_try_nestloop_path(m, l0, l1, v3324, v3437, v3326, v3017, v39+int32(8))
	mBase = m.M
	v3469 = m.ExcPending
	if v3469 != 0 {
		goto L22
	} else {
		goto L648
	}
L634:
	;
	v3332 = *(*int32)(unsafe.Add(mBase, uint32(l3)+56))
	if v3332 == int32(0) {
		goto L635
	} else {
		goto L636
	}
L635:
	;
	if v3158 == int32(0) {
		goto L630
	} else {
		goto L647
	}
L636:
	;
	v3335 = int32(0)
	v3336 = *(*int32)(unsafe.Add(mBase, uint32(v3332)+4))
	if v3336 <= v3335 {
		goto L635
	} else {
		goto L637
	}
L637:
	;
	v3346 = v3335
	goto L638
L638:
	;
	v3375 = *(*int32)(unsafe.Add(mBase, uint32(v3332)+12))
	v3379 = *(*int32)(unsafe.Add(mBase, uint32(v3375+v3346<<(uint(int32(2))%32))))
	v3381 = v39 + int32(8)
	F_try_nestloop_path(m, l0, l1, v3324, v3379, v3326, v3017, v3381)
	mBase = m.M
	v3383 = m.ExcPending
	if v3383 != 0 {
		goto L22
	} else {
		goto L640
	}
L639:
	;
	goto L635
L640:
	;
	v3384 = F_get_memoize_path(m, l0, l3, l2, v3379, v3324, v3017, v3381)
	mBase = m.M
	v3385 = m.ExcPending
	if v3385 != 0 {
		goto L22
	} else {
		goto L641
	}
L641:
	;
	if v3384 != 0 {
		goto L642
	} else {
		goto L643
	}
L642:
	;
	F_try_nestloop_path(m, l0, l1, v3324, v3384, v3326, v3017, v3381)
	mBase = m.M
	v3387 = m.ExcPending
	if v3387 != 0 {
		goto L22
	} else {
		goto L645
	}
L643:
	;
	goto L644
L644:
	;
	v3389 = v3346 + int32(1)
	v3390 = *(*int32)(unsafe.Add(mBase, uint32(v3332)+4))
	if v3389 < v3390 {
		v3346 = v3389
		goto L638
	} else {
		goto L646
	}
L645:
	;
	goto L644
L646:
	;
	goto L639
L647:
	;
	v3437 = v3158
	goto L633
L648:
	;
	goto L630
L649:
	;
	F_generate_mergejoin_paths(m, l0, l1, l3, v3324, l4, v39+int32(8), v3019, v3156, v3326, int32(0))
	mBase = m.M
	v3510 = m.ExcPending
	if v3510 != 0 {
		goto L22
	} else {
		goto L650
	}
L650:
	;
	goto L592
L651:
	;
	goto L591
L652:
	;
	v3597 = base.B2i32(base.Ui32(l4) <= base.Ui32(int32(8)))
	goto L654
L653:
	;
	v3597 = int32(0)
	goto L654
L654:
	;
	if base.B2i32(v3587 != v3588)|v3597 != 0 {
		goto L274
	} else {
		goto L655
	}
L655:
	;
	v3599 = *(*int32)(unsafe.Add(mBase, uint32(l2)+40))
	if v3599 == int32(0) {
		goto L274
	} else {
		goto L656
	}
L656:
	;
	v3602 = *(*int32)(unsafe.Add(mBase, uint32(l1)+64))
	if v3602 != 0 {
		goto L274
	} else {
		goto L657
	}
L657:
	;
	if v3018 != 0 {
		goto L658
	} else {
		goto L659
	}
L658:
	;
	v3604 = v39 + int32(8)
	v3605 = int32(0)
	v3608 = base.B2i32(l4 == int32(9))
	if l4 == int32(9) {
		v3734 = v3605
		goto L661
	} else {
		goto L662
	}
L659:
	;
	goto L660
L660:
	;
	if v3156 != 0 {
		goto L731
	} else {
		goto L732
	}
L661:
	;
	v3735 = *(*int32)(unsafe.Add(mBase, uint32(l2)+40))
	if v3735 == int32(0) {
		goto L699
	} else {
		goto L700
	}
L662:
	;
	v3610 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_add_paths_to_joinrel[1])))
	if v3610&int32(1) == int32(0) {
		v3734 = v3605
		goto L661
	} else {
		goto L663
	}
L663:
	;
	v3615 = *(*int32)(unsafe.Add(mBase, uint32(l3)+48))
	v3616 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v3615)+21)))
	if v3616 != int32(1) {
		v3734 = v3605
		goto L661
	} else {
		goto L664
	}
L664:
	;
	v3619 = *(*int32)(unsafe.Add(mBase, uint32(v3615)+16))
	if v3619 == int32(0) {
		goto L665
	} else {
		goto L666
	}
L665:
	;
	v3722 = *(*int32)(unsafe.Add(mBase, uint32(v3615)+4))
	v3724 = v3722 - int32(348)
	goto L696
L666:
	;
	v3622 = *(*int32)(unsafe.Add(mBase, uint32(v3619)+4))
	v3623 = *(*int32)(unsafe.Add(mBase, uint32(l2)+8))
	v3624 = int32(0)
	if base.B2i32(v3622 == v3624)|base.B2i32(v3623 == v3624) != 0 {
		v3669 = v3624
		goto L668
	} else {
		goto L669
	}
L667:
	;
	if v3669 != 0 {
		v3734 = v3605
		goto L661
	} else {
		goto L680
	}
L668:
	;
	goto L667
L669:
	;
	v3634 = *(*int32)(unsafe.Add(mBase, uint32(v3622)+4))
	v3635 = *(*int32)(unsafe.Add(mBase, uint32(v3623)+4))
	if v3634 < v3635 {
		goto L670
	} else {
		goto L671
	}
L670:
	;
	v3637 = v3634
	goto L672
L671:
	;
	v3637 = v3635
	goto L672
L672:
	;
	if v3637 <= int32(1) {
		goto L673
	} else {
		goto L674
	}
L673:
	;
	v3640 = int32(1)
	goto L675
L674:
	;
	v3640 = v3637
	goto L675
L675:
	;
	v3641 = int32(8)
	v3646 = int32(0)
	goto L676
L676:
	;
	v3653 = v3646 << (uint(int32(2)) % 32)
	v3655 = *(*int32)(unsafe.Add(mBase, uint32(v3623+v3641+v3653)))
	v3657 = *(*int32)(unsafe.Add(mBase, uint32(v3622+v3641+v3653)))
	v3658 = v3655 & v3657
	v3660 = base.B2i32(v3658 != int32(0))
	if v3658 != 0 {
		v3669 = v3660
		goto L668
	} else {
		goto L678
	}
L677:
	;
	v3669 = v3660
	goto L668
L678:
	;
	v3662 = v3646 + int32(1)
	if v3662 != v3640 {
		v3646 = v3662
		goto L676
	} else {
		goto L679
	}
L679:
	;
	goto L677
L680:
	;
	v3670 = *(*int32)(unsafe.Add(mBase, uint32(v3615)+16))
	if v3670 == int32(0) {
		goto L665
	} else {
		goto L681
	}
L681:
	;
	v3673 = *(*int32)(unsafe.Add(mBase, uint32(v3670)+4))
	v3674 = *(*int32)(unsafe.Add(mBase, uint32(l2)+228))
	v3675 = int32(0)
	if base.B2i32(v3673 == v3675)|base.B2i32(v3674 == v3675) != 0 {
		v3720 = v3675
		goto L683
	} else {
		goto L684
	}
L682:
	;
	if v3720 != 0 {
		v3734 = v3605
		goto L661
	} else {
		goto L695
	}
L683:
	;
	goto L682
L684:
	;
	v3685 = *(*int32)(unsafe.Add(mBase, uint32(v3673)+4))
	v3686 = *(*int32)(unsafe.Add(mBase, uint32(v3674)+4))
	if v3685 < v3686 {
		goto L685
	} else {
		goto L686
	}
L685:
	;
	v3688 = v3685
	goto L687
L686:
	;
	v3688 = v3686
	goto L687
L687:
	;
	if v3688 <= int32(1) {
		goto L688
	} else {
		goto L689
	}
L688:
	;
	v3691 = int32(1)
	goto L690
L689:
	;
	v3691 = v3688
	goto L690
L690:
	;
	v3692 = int32(8)
	v3697 = int32(0)
	goto L691
L691:
	;
	v3704 = v3697 << (uint(int32(2)) % 32)
	v3706 = *(*int32)(unsafe.Add(mBase, uint32(v3674+v3692+v3704)))
	v3708 = *(*int32)(unsafe.Add(mBase, uint32(v3673+v3692+v3704)))
	v3709 = v3706 & v3708
	v3711 = base.B2i32(v3709 != int32(0))
	if v3709 != 0 {
		v3720 = v3711
		goto L683
	} else {
		goto L693
	}
L692:
	;
	v3720 = v3711
	goto L683
L693:
	;
	v3713 = v3697 + int32(1)
	if v3713 != v3691 {
		v3697 = v3713
		goto L691
	} else {
		goto L694
	}
L694:
	;
	goto L692
L695:
	;
	goto L665
L696:
	;
	if base.B2i32(base.Ui32(v3724) < base.Ui32(int32(15)))&int32(base.Ui32(int32(_a_F_add_paths_to_joinrel_3))>>(uint(v3724)%32)) != 0 {
		v3734 = v3605
		goto L661
	} else {
		goto L697
	}
L697:
	;
	v3730 = F_create_material_path(m, l3, v3615)
	mBase = m.M
	v3731 = m.ExcPending
	if v3731 != 0 {
		goto L22
	} else {
		goto L698
	}
L698:
	;
	v3734 = v3730
	goto L661
L699:
	;
	goto L660
L700:
	;
	v3738 = *(*int32)(unsafe.Add(mBase, uint32(v3735)+4))
	if v3738 <= int32(0) {
		goto L699
	} else {
		goto L701
	}
L701:
	;
	if l4 == int32(9) {
		goto L702
	} else {
		goto L703
	}
L702:
	;
	v3742 = int32(0)
	goto L704
L703:
	;
	v3742 = l4
	goto L704
L704:
	;
	v3771 = v3605
	goto L705
L705:
	;
	v3781 = *(*int32)(unsafe.Add(mBase, uint32(v3735)+12))
	v3785 = *(*int32)(unsafe.Add(mBase, uint32(v3781+v3771<<(uint(int32(2))%32))))
	v3786 = *(*int32)(unsafe.Add(mBase, uint32(v3785)+64))
	v3787 = F_build_join_pathkeys(m, l0, l1, v3742, v3786)
	mBase = m.M
	v3788 = m.ExcPending
	if v3788 != 0 {
		goto L22
	} else {
		goto L707
	}
L706:
	;
	goto L699
L707:
	;
	v3789 = *(*int32)(unsafe.Add(mBase, uint32(l3)+56))
	if v3789 == int32(0) {
		goto L708
	} else {
		goto L709
	}
L708:
	;
	if v3734 != 0 {
		goto L725
	} else {
		goto L726
	}
L709:
	;
	v3792 = int32(0)
	v3793 = *(*int32)(unsafe.Add(mBase, uint32(v3789)+4))
	if v3793 <= v3792 {
		goto L708
	} else {
		goto L710
	}
L710:
	;
	v3801 = v3792
	goto L711
L711:
	;
	v3832 = *(*int32)(unsafe.Add(mBase, uint32(v3789)+12))
	v3836 = *(*int32)(unsafe.Add(mBase, uint32(v3832+v3801<<(uint(int32(2))%32))))
	v3837 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v3836)+21)))
	if v3837 == int32(0) {
		goto L713
	} else {
		goto L714
	}
L712:
	;
	goto L708
L713:
	;
	v3858 = v3801 + int32(1)
	v3859 = *(*int32)(unsafe.Add(mBase, uint32(v3789)+4))
	if v3858 < v3859 {
		v3801 = v3858
		goto L711
	} else {
		goto L724
	}
L714:
	;
	if base.B2i32(l4 != int32(9)) == int32(0) {
		goto L715
	} else {
		goto L716
	}
L715:
	;
	v3842 = *(*int32)(unsafe.Add(mBase, uint32(l3)+48))
	if v3836 != v3842 {
		goto L713
	} else {
		goto L718
	}
L716:
	;
	v3847 = v3836
	goto L717
L717:
	;
	F_try_partial_nestloop_path(m, l0, l1, v3785, v3847, v3787, v3742, v3604)
	mBase = m.M
	v3849 = m.ExcPending
	if v3849 != 0 {
		goto L22
	} else {
		goto L720
	}
L718:
	;
	v3844 = *(*int32)(unsafe.Add(mBase, uint32(v3604)+12))
	v3845 = F_create_unique_path(m, l0, l3, v3836, v3844)
	mBase = m.M
	v3846 = m.ExcPending
	if v3846 != 0 {
		goto L22
	} else {
		goto L719
	}
L719:
	;
	v3847 = v3845
	goto L717
L720:
	;
	v3850 = F_get_memoize_path(m, l0, l3, l2, v3847, v3785, v3742, v3604)
	mBase = m.M
	v3851 = m.ExcPending
	if v3851 != 0 {
		goto L22
	} else {
		goto L721
	}
L721:
	;
	if v3850 == int32(0) {
		goto L713
	} else {
		goto L722
	}
L722:
	;
	F_try_partial_nestloop_path(m, l0, l1, v3785, v3850, v3787, v3742, v3604)
	mBase = m.M
	v3855 = m.ExcPending
	if v3855 != 0 {
		goto L22
	} else {
		goto L723
	}
L723:
	;
	goto L713
L724:
	;
	goto L712
L725:
	;
	F_try_partial_nestloop_path(m, l0, l1, v3785, v3734, v3787, v3742, v3604)
	mBase = m.M
	v3898 = m.ExcPending
	if v3898 != 0 {
		goto L22
	} else {
		goto L728
	}
L726:
	;
	goto L727
L727:
	;
	v3900 = v3771 + int32(1)
	v3901 = *(*int32)(unsafe.Add(mBase, uint32(v3735)+4))
	if v3900 < v3901 {
		v3771 = v3900
		goto L705
	} else {
		goto L729
	}
L728:
	;
	goto L727
L729:
	;
	goto L706
L730:
	;
	v4028 = *(*int32)(unsafe.Add(mBase, uint32(l2)+40))
	if v4028 == int32(0) {
		goto L754
	} else {
		goto L755
	}
L731:
	;
	v3975 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v3156)+21)))
	if v3975 != 0 {
		v4025 = v3156
		goto L730
	} else {
		goto L734
	}
L732:
	;
	goto L733
L733:
	;
	if v3157 != 0 {
		goto L274
	} else {
		goto L735
	}
L734:
	;
	goto L733
L735:
	;
	v3976 = *(*int32)(unsafe.Add(mBase, uint32(l3)+32))
	if v3976 != 0 {
		goto L738
	} else {
		goto L739
	}
L736:
	;
	if v4020 == int32(0) {
		goto L274
	} else {
		goto L753
	}
L737:
	;
	goto L736
L738:
	;
	v3981 = *(*int32)(unsafe.Add(mBase, uint32(v3976)+4))
	if v3981 <= int32(0) {
		v4020 = int32(0)
		goto L737
	} else {
		goto L741
	}
L739:
	;
	goto L740
L740:
	;
	v4020 = int32(0)
	goto L737
L741:
	;
	v3984 = int32(0)
	if v3984 < v3981 {
		goto L742
	} else {
		goto L743
	}
L742:
	;
	v3987 = v3981
	goto L744
L743:
	;
	v3987 = v3984
	goto L744
L744:
	;
	v3988 = *(*int32)(unsafe.Add(mBase, uint32(v3976)+12))
	v3990 = int32(0)
	goto L745
L745:
	;
	v3998 = *(*int32)(unsafe.Add(mBase, uint32(v3988+v3990<<(uint(int32(2))%32))))
	v3999 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v3998)+21)))
	if v3999 == int32(1) {
		goto L747
	} else {
		goto L748
	}
L746:
	;
	goto L740
L747:
	;
	v4002 = *(*int32)(unsafe.Add(mBase, uint32(v3998)+16))
	if v4002 == int32(0) {
		v4020 = v3998
		goto L737
	} else {
		goto L750
	}
L748:
	;
	goto L749
L749:
	;
	v4010 = v3990 + int32(1)
	if v4010 != v3987 {
		v3990 = v4010
		goto L745
	} else {
		goto L752
	}
L750:
	;
	v4005 = *(*int32)(unsafe.Add(mBase, uint32(v4002)+4))
	if v4005 == int32(0) {
		v4020 = v3998
		goto L737
	} else {
		goto L751
	}
L751:
	;
	goto L749
L752:
	;
	goto L746
L753:
	;
	v4025 = v4020
	goto L730
L754:
	;
	goto L274
L755:
	;
	v4031 = *(*int32)(unsafe.Add(mBase, uint32(v4028)+4))
	if v4031 <= int32(0) {
		goto L754
	} else {
		goto L756
	}
L756:
	;
	v4040 = int32(0)
	goto L757
L757:
	;
	v4071 = *(*int32)(unsafe.Add(mBase, uint32(v4028)+12))
	v4075 = *(*int32)(unsafe.Add(mBase, uint32(v4071+v4040<<(uint(int32(2))%32))))
	v4077 = *(*int32)(unsafe.Add(mBase, uint32(v4075)+64))
	v4078 = F_build_join_pathkeys(m, l0, l1, l4, v4077)
	mBase = m.M
	v4079 = m.ExcPending
	if v4079 != 0 {
		goto L22
	} else {
		goto L759
	}
L758:
	;
	goto L754
L759:
	;
	F_generate_mergejoin_paths(m, l0, l1, l3, v4075, l4, v39+int32(8), int32(0), v4025, v4078, int32(1))
	mBase = m.M
	v4082 = m.ExcPending
	if v4082 != 0 {
		goto L22
	} else {
		goto L760
	}
L760:
	;
	v4084 = v4040 + int32(1)
	v4085 = *(*int32)(unsafe.Add(mBase, uint32(v4028)+4))
	if v4084 < v4085 {
		v4040 = v4084
		goto L757
	} else {
		goto L761
	}
L761:
	;
	goto L758
L762:
	;
	v5340 = *(*int32)(unsafe.Add(mBase, uint32(l1)+168))
	if v5340 == int32(0) {
		goto L1056
	} else {
		goto L1057
	}
L763:
	;
	v4162 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_add_paths_to_joinrel[2])))
	if v4162&int32(1) == int32(0) {
		goto L762
	} else {
		goto L766
	}
L764:
	;
	goto L765
L765:
	;
	v4167 = *(*int32)(unsafe.Add(mBase, uint32(v39)+8))
	if v4167 == int32(0) {
		goto L762
	} else {
		goto L767
	}
L766:
	;
	goto L765
L767:
	;
	v4170 = *(*int32)(unsafe.Add(mBase, uint32(v4167)+4))
	if v4170 <= int32(0) {
		goto L762
	} else {
		goto L768
	}
L768:
	;
	v4177 = int32(0)
	v4185 = v4177
	v4193 = v4177
	goto L769
L769:
	;
	v4215 = *(*int32)(unsafe.Add(mBase, uint32(v4167)+12))
	v4219 = *(*int32)(unsafe.Add(mBase, uint32(v4215+v4185<<(uint(int32(2))%32))))
	if int32(1)<<(uint(l4)%32)&int32(174) != 0 {
		goto L772
	} else {
		goto L773
	}
L770:
	;
	if v4524 == int32(0) {
		goto L762
	} else {
		goto L861
	}
L771:
	;
	v4527 = v4185 + int32(1)
	v4528 = *(*int32)(unsafe.Add(mBase, uint32(v4167)+4))
	if v4527 < v4528 {
		v4185 = v4527
		v4193 = v4524
		goto L769
	} else {
		goto L860
	}
L772:
	;
	v4220 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v4219)+8)))
	if v4220 != 0 {
		v4524 = v4193
		goto L771
	} else {
		goto L775
	}
L773:
	;
	goto L774
L774:
	;
	v4279 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v4219)+9)))
	if v4279 != int32(1) {
		v4524 = v4193
		goto L771
	} else {
		goto L791
	}
L775:
	;
	v4221 = *(*int32)(unsafe.Add(mBase, uint32(v4219)+32))
	v4222 = *(*int32)(unsafe.Add(mBase, uint32(l1)+8))
	v4223 = int32(0)
	if v4221 == v4223 {
		goto L777
	} else {
		goto L778
	}
L776:
	;
	if v4276 == int32(0) {
		v4524 = v4193
		goto L771
	} else {
		goto L790
	}
L777:
	;
	v4276 = int32(1)
	goto L776
L778:
	;
	goto L779
L779:
	;
	if v4222 == int32(0) {
		v4269 = v4223
		goto L780
	} else {
		goto L781
	}
L780:
	;
	v4276 = v4269
	goto L776
L781:
	;
	v4232 = *(*int32)(unsafe.Add(mBase, uint32(v4221)+4))
	v4233 = *(*int32)(unsafe.Add(mBase, uint32(v4222)+4))
	if v4233 < v4232 {
		v4269 = v4223
		goto L780
	} else {
		goto L782
	}
L782:
	;
	v4235 = int32(1)
	if v4232 <= v4235 {
		goto L783
	} else {
		goto L784
	}
L783:
	;
	v4238 = v4235
	goto L785
L784:
	;
	v4238 = v4232
	goto L785
L785:
	;
	v4239 = int32(8)
	v4244 = int32(0)
	goto L786
L786:
	;
	v4251 = v4244 << (uint(int32(2)) % 32)
	v4253 = *(*int32)(unsafe.Add(mBase, uint32(v4221+v4239+v4251)))
	v4255 = *(*int32)(unsafe.Add(mBase, uint32(v4222+v4239+v4251)))
	v4258 = v4253 & (v4255 ^ int32(-1))
	v4260 = base.B2i32(v4258 == int32(0))
	if v4258 != 0 {
		v4269 = v4260
		goto L780
	} else {
		goto L788
	}
L787:
	;
	v4269 = v4260
	goto L780
L788:
	;
	v4262 = v4244 + int32(1)
	if v4262 != v4238 {
		v4244 = v4262
		goto L786
	} else {
		goto L789
	}
L789:
	;
	goto L787
L790:
	;
	goto L774
L791:
	;
	v4282 = *(*int32)(unsafe.Add(mBase, uint32(v4219)+124))
	if v4282 == int32(0) {
		v4524 = v4193
		goto L771
	} else {
		goto L792
	}
L792:
	;
	v4285 = *(*int32)(unsafe.Add(mBase, uint32(l3)+8))
	v4286 = *(*int32)(unsafe.Add(mBase, uint32(v4219)+44))
	v4287 = *(*int32)(unsafe.Add(mBase, uint32(l2)+8))
	v4288 = int32(0)
	if v4286 == v4288 {
		goto L796
	} else {
		goto L797
	}
L793:
	;
	v4521 = F_lappend(m, v4193, v4219)
	mBase = m.M
	v4522 = m.ExcPending
	if v4522 != 0 {
		goto L22
	} else {
		goto L859
	}
L794:
	;
	v4519 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(v4219)+120)) = uint8(v4519)
	goto L793
L795:
	;
	if v4341 != 0 {
		goto L809
	} else {
		goto L810
	}
L796:
	;
	v4341 = int32(1)
	goto L795
L797:
	;
	goto L798
L798:
	;
	if v4287 == int32(0) {
		v4334 = v4288
		goto L799
	} else {
		goto L800
	}
L799:
	;
	v4341 = v4334
	goto L795
L800:
	;
	v4297 = *(*int32)(unsafe.Add(mBase, uint32(v4286)+4))
	v4298 = *(*int32)(unsafe.Add(mBase, uint32(v4287)+4))
	if v4298 < v4297 {
		v4334 = v4288
		goto L799
	} else {
		goto L801
	}
L801:
	;
	v4300 = int32(1)
	if v4297 <= v4300 {
		goto L802
	} else {
		goto L803
	}
L802:
	;
	v4303 = v4300
	goto L804
L803:
	;
	v4303 = v4297
	goto L804
L804:
	;
	v4304 = int32(8)
	v4309 = int32(0)
	goto L805
L805:
	;
	v4316 = v4309 << (uint(int32(2)) % 32)
	v4318 = *(*int32)(unsafe.Add(mBase, uint32(v4286+v4304+v4316)))
	v4320 = *(*int32)(unsafe.Add(mBase, uint32(v4287+v4304+v4316)))
	v4323 = v4318 & (v4320 ^ int32(-1))
	v4325 = base.B2i32(v4323 == int32(0))
	if v4323 != 0 {
		v4334 = v4325
		goto L799
	} else {
		goto L807
	}
L806:
	;
	v4334 = v4325
	goto L799
L807:
	;
	v4327 = v4309 + int32(1)
	if v4327 != v4303 {
		v4309 = v4327
		goto L805
	} else {
		goto L808
	}
L808:
	;
	goto L806
L809:
	;
	v4342 = *(*int32)(unsafe.Add(mBase, uint32(v4219)+48))
	v4343 = int32(0)
	if v4342 == v4343 {
		goto L813
	} else {
		goto L814
	}
L810:
	;
	goto L811
L811:
	;
	v4397 = *(*int32)(unsafe.Add(mBase, uint32(v4219)+44))
	v4398 = int32(0)
	if v4397 == v4398 {
		goto L828
	} else {
		goto L829
	}
L812:
	;
	if v4396 != 0 {
		goto L794
	} else {
		goto L826
	}
L813:
	;
	v4396 = int32(1)
	goto L812
L814:
	;
	goto L815
L815:
	;
	if v4285 == int32(0) {
		v4389 = v4343
		goto L816
	} else {
		goto L817
	}
L816:
	;
	v4396 = v4389
	goto L812
L817:
	;
	v4352 = *(*int32)(unsafe.Add(mBase, uint32(v4342)+4))
	v4353 = *(*int32)(unsafe.Add(mBase, uint32(v4285)+4))
	if v4353 < v4352 {
		v4389 = v4343
		goto L816
	} else {
		goto L818
	}
L818:
	;
	v4355 = int32(1)
	if v4352 <= v4355 {
		goto L819
	} else {
		goto L820
	}
L819:
	;
	v4358 = v4355
	goto L821
L820:
	;
	v4358 = v4352
	goto L821
L821:
	;
	v4359 = int32(8)
	v4364 = int32(0)
	goto L822
L822:
	;
	v4371 = v4364 << (uint(int32(2)) % 32)
	v4373 = *(*int32)(unsafe.Add(mBase, uint32(v4342+v4359+v4371)))
	v4375 = *(*int32)(unsafe.Add(mBase, uint32(v4285+v4359+v4371)))
	v4378 = v4373 & (v4375 ^ int32(-1))
	v4380 = base.B2i32(v4378 == int32(0))
	if v4378 != 0 {
		v4389 = v4380
		goto L816
	} else {
		goto L824
	}
L823:
	;
	v4389 = v4380
	goto L816
L824:
	;
	v4382 = v4364 + int32(1)
	if v4382 != v4358 {
		v4364 = v4382
		goto L822
	} else {
		goto L825
	}
L825:
	;
	goto L823
L826:
	;
	goto L811
L827:
	;
	if v4451 == int32(0) {
		v4524 = v4193
		goto L771
	} else {
		goto L841
	}
L828:
	;
	v4451 = int32(1)
	goto L827
L829:
	;
	goto L830
L830:
	;
	if v4285 == int32(0) {
		v4444 = v4398
		goto L831
	} else {
		goto L832
	}
L831:
	;
	v4451 = v4444
	goto L827
L832:
	;
	v4407 = *(*int32)(unsafe.Add(mBase, uint32(v4397)+4))
	v4408 = *(*int32)(unsafe.Add(mBase, uint32(v4285)+4))
	if v4408 < v4407 {
		v4444 = v4398
		goto L831
	} else {
		goto L833
	}
L833:
	;
	v4410 = int32(1)
	if v4407 <= v4410 {
		goto L834
	} else {
		goto L835
	}
L834:
	;
	v4413 = v4410
	goto L836
L835:
	;
	v4413 = v4407
	goto L836
L836:
	;
	v4414 = int32(8)
	v4419 = int32(0)
	goto L837
L837:
	;
	v4426 = v4419 << (uint(int32(2)) % 32)
	v4428 = *(*int32)(unsafe.Add(mBase, uint32(v4397+v4414+v4426)))
	v4430 = *(*int32)(unsafe.Add(mBase, uint32(v4285+v4414+v4426)))
	v4433 = v4428 & (v4430 ^ int32(-1))
	v4435 = base.B2i32(v4433 == int32(0))
	if v4433 != 0 {
		v4444 = v4435
		goto L831
	} else {
		goto L839
	}
L838:
	;
	v4444 = v4435
	goto L831
L839:
	;
	v4437 = v4419 + int32(1)
	if v4437 != v4413 {
		v4419 = v4437
		goto L837
	} else {
		goto L840
	}
L840:
	;
	goto L838
L841:
	;
	v4454 = *(*int32)(unsafe.Add(mBase, uint32(v4219)+48))
	v4455 = int32(0)
	if v4454 == v4455 {
		goto L843
	} else {
		goto L844
	}
L842:
	;
	if v4508 == int32(0) {
		v4524 = v4193
		goto L771
	} else {
		goto L856
	}
L843:
	;
	v4508 = int32(1)
	goto L842
L844:
	;
	goto L845
L845:
	;
	if v4287 == int32(0) {
		v4501 = v4455
		goto L846
	} else {
		goto L847
	}
L846:
	;
	v4508 = v4501
	goto L842
L847:
	;
	v4464 = *(*int32)(unsafe.Add(mBase, uint32(v4454)+4))
	v4465 = *(*int32)(unsafe.Add(mBase, uint32(v4287)+4))
	if v4465 < v4464 {
		v4501 = v4455
		goto L846
	} else {
		goto L848
	}
L848:
	;
	v4467 = int32(1)
	if v4464 <= v4467 {
		goto L849
	} else {
		goto L850
	}
L849:
	;
	v4470 = v4467
	goto L851
L850:
	;
	v4470 = v4464
	goto L851
L851:
	;
	v4471 = int32(8)
	v4476 = int32(0)
	goto L852
L852:
	;
	v4483 = v4476 << (uint(int32(2)) % 32)
	v4485 = *(*int32)(unsafe.Add(mBase, uint32(v4454+v4471+v4483)))
	v4487 = *(*int32)(unsafe.Add(mBase, uint32(v4287+v4471+v4483)))
	v4490 = v4485 & (v4487 ^ int32(-1))
	v4492 = base.B2i32(v4490 == int32(0))
	if v4490 != 0 {
		v4501 = v4492
		goto L846
	} else {
		goto L854
	}
L853:
	;
	v4501 = v4492
	goto L846
L854:
	;
	v4494 = v4476 + int32(1)
	if v4494 != v4470 {
		v4476 = v4494
		goto L852
	} else {
		goto L855
	}
L855:
	;
	goto L853
L856:
	;
	v4511 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v4219)+120)) = uint8(v4511)
	v4513 = *(*int32)(unsafe.Add(mBase, uint32(v4219)+4))
	v4514 = *(*int32)(unsafe.Add(mBase, uint32(v4513)+4))
	v4515 = F_get_commutator(m, v4514)
	mBase = m.M
	v4516 = m.ExcPending
	if v4516 != 0 {
		goto L22
	} else {
		goto L857
	}
L857:
	;
	if v4515 == int32(0) {
		v4524 = v4193
		goto L771
	} else {
		goto L858
	}
L858:
	;
	goto L793
L859:
	;
	v4524 = v4521
	goto L771
L860:
	;
	goto L770
L861:
	;
	v4532 = *(*int32)(unsafe.Add(mBase, uint32(l3)+48))
	v4533 = *(*int32)(unsafe.Add(mBase, uint32(l2)+44))
	v4534 = *(*int32)(unsafe.Add(mBase, uint32(l2)+48))
	v4535 = *(*int32)(unsafe.Add(mBase, uint32(v4534)+16))
	if v4535 == int32(0) {
		goto L862
	} else {
		goto L863
	}
L862:
	;
	v4638 = *(*int32)(unsafe.Add(mBase, uint32(v4532)+16))
	if v4638 == int32(0) {
		goto L893
	} else {
		goto L894
	}
L863:
	;
	v4538 = *(*int32)(unsafe.Add(mBase, uint32(v4535)+4))
	v4539 = *(*int32)(unsafe.Add(mBase, uint32(l3)+8))
	v4540 = int32(0)
	if base.B2i32(v4538 == v4540)|base.B2i32(v4539 == v4540) != 0 {
		v4585 = v4540
		goto L865
	} else {
		goto L866
	}
L864:
	;
	if v4585 != 0 {
		goto L762
	} else {
		goto L877
	}
L865:
	;
	goto L864
L866:
	;
	v4550 = *(*int32)(unsafe.Add(mBase, uint32(v4538)+4))
	v4551 = *(*int32)(unsafe.Add(mBase, uint32(v4539)+4))
	if v4550 < v4551 {
		goto L867
	} else {
		goto L868
	}
L867:
	;
	v4553 = v4550
	goto L869
L868:
	;
	v4553 = v4551
	goto L869
L869:
	;
	if v4553 <= int32(1) {
		goto L870
	} else {
		goto L871
	}
L870:
	;
	v4556 = int32(1)
	goto L872
L871:
	;
	v4556 = v4553
	goto L872
L872:
	;
	v4557 = int32(8)
	v4562 = int32(0)
	goto L873
L873:
	;
	v4569 = v4562 << (uint(int32(2)) % 32)
	v4571 = *(*int32)(unsafe.Add(mBase, uint32(v4539+v4557+v4569)))
	v4573 = *(*int32)(unsafe.Add(mBase, uint32(v4538+v4557+v4569)))
	v4574 = v4571 & v4573
	v4576 = base.B2i32(v4574 != int32(0))
	if v4574 != 0 {
		v4585 = v4576
		goto L865
	} else {
		goto L875
	}
L874:
	;
	v4585 = v4576
	goto L865
L875:
	;
	v4578 = v4562 + int32(1)
	if v4578 != v4556 {
		v4562 = v4578
		goto L873
	} else {
		goto L876
	}
L876:
	;
	goto L874
L877:
	;
	v4586 = *(*int32)(unsafe.Add(mBase, uint32(v4534)+16))
	if v4586 == int32(0) {
		goto L862
	} else {
		goto L878
	}
L878:
	;
	v4589 = *(*int32)(unsafe.Add(mBase, uint32(v4586)+4))
	v4590 = *(*int32)(unsafe.Add(mBase, uint32(l3)+228))
	v4591 = int32(0)
	if base.B2i32(v4589 == v4591)|base.B2i32(v4590 == v4591) != 0 {
		v4636 = v4591
		goto L880
	} else {
		goto L881
	}
L879:
	;
	if v4636 != 0 {
		goto L762
	} else {
		goto L892
	}
L880:
	;
	goto L879
L881:
	;
	v4601 = *(*int32)(unsafe.Add(mBase, uint32(v4589)+4))
	v4602 = *(*int32)(unsafe.Add(mBase, uint32(v4590)+4))
	if v4601 < v4602 {
		goto L882
	} else {
		goto L883
	}
L882:
	;
	v4604 = v4601
	goto L884
L883:
	;
	v4604 = v4602
	goto L884
L884:
	;
	if v4604 <= int32(1) {
		goto L885
	} else {
		goto L886
	}
L885:
	;
	v4607 = int32(1)
	goto L887
L886:
	;
	v4607 = v4604
	goto L887
L887:
	;
	v4608 = int32(8)
	v4613 = int32(0)
	goto L888
L888:
	;
	v4620 = v4613 << (uint(int32(2)) % 32)
	v4622 = *(*int32)(unsafe.Add(mBase, uint32(v4590+v4608+v4620)))
	v4624 = *(*int32)(unsafe.Add(mBase, uint32(v4589+v4608+v4620)))
	v4625 = v4622 & v4624
	v4627 = base.B2i32(v4625 != int32(0))
	if v4625 != 0 {
		v4636 = v4627
		goto L880
	} else {
		goto L890
	}
L889:
	;
	v4636 = v4627
	goto L880
L890:
	;
	v4629 = v4613 + int32(1)
	if v4629 != v4607 {
		v4613 = v4629
		goto L888
	} else {
		goto L891
	}
L891:
	;
	goto L889
L892:
	;
	goto L862
L893:
	;
	switch l4 - int32(8) {
	case 0:
		goto L927
	case 1:
		goto L926
	default:
		goto L925
	}
L894:
	;
	v4641 = *(*int32)(unsafe.Add(mBase, uint32(v4638)+4))
	v4642 = *(*int32)(unsafe.Add(mBase, uint32(l2)+8))
	v4643 = int32(0)
	if base.B2i32(v4641 == v4643)|base.B2i32(v4642 == v4643) != 0 {
		v4688 = v4643
		goto L896
	} else {
		goto L897
	}
L895:
	;
	if v4688 != 0 {
		goto L762
	} else {
		goto L908
	}
L896:
	;
	goto L895
L897:
	;
	v4653 = *(*int32)(unsafe.Add(mBase, uint32(v4641)+4))
	v4654 = *(*int32)(unsafe.Add(mBase, uint32(v4642)+4))
	if v4653 < v4654 {
		goto L898
	} else {
		goto L899
	}
L898:
	;
	v4656 = v4653
	goto L900
L899:
	;
	v4656 = v4654
	goto L900
L900:
	;
	if v4656 <= int32(1) {
		goto L901
	} else {
		goto L902
	}
L901:
	;
	v4659 = int32(1)
	goto L903
L902:
	;
	v4659 = v4656
	goto L903
L903:
	;
	v4660 = int32(8)
	v4665 = int32(0)
	goto L904
L904:
	;
	v4672 = v4665 << (uint(int32(2)) % 32)
	v4674 = *(*int32)(unsafe.Add(mBase, uint32(v4642+v4660+v4672)))
	v4676 = *(*int32)(unsafe.Add(mBase, uint32(v4641+v4660+v4672)))
	v4677 = v4674 & v4676
	v4679 = base.B2i32(v4677 != int32(0))
	if v4677 != 0 {
		v4688 = v4679
		goto L896
	} else {
		goto L906
	}
L905:
	;
	v4688 = v4679
	goto L896
L906:
	;
	v4681 = v4665 + int32(1)
	if v4681 != v4659 {
		v4665 = v4681
		goto L904
	} else {
		goto L907
	}
L907:
	;
	goto L905
L908:
	;
	v4689 = *(*int32)(unsafe.Add(mBase, uint32(v4532)+16))
	if v4689 == int32(0) {
		goto L893
	} else {
		goto L909
	}
L909:
	;
	v4692 = *(*int32)(unsafe.Add(mBase, uint32(v4689)+4))
	v4693 = *(*int32)(unsafe.Add(mBase, uint32(l2)+228))
	v4694 = int32(0)
	if base.B2i32(v4692 == v4694)|base.B2i32(v4693 == v4694) != 0 {
		v4739 = v4694
		goto L911
	} else {
		goto L912
	}
L910:
	;
	if v4739 != 0 {
		goto L762
	} else {
		goto L923
	}
L911:
	;
	goto L910
L912:
	;
	v4704 = *(*int32)(unsafe.Add(mBase, uint32(v4692)+4))
	v4705 = *(*int32)(unsafe.Add(mBase, uint32(v4693)+4))
	if v4704 < v4705 {
		goto L913
	} else {
		goto L914
	}
L913:
	;
	v4707 = v4704
	goto L915
L914:
	;
	v4707 = v4705
	goto L915
L915:
	;
	if v4707 <= int32(1) {
		goto L916
	} else {
		goto L917
	}
L916:
	;
	v4710 = int32(1)
	goto L918
L917:
	;
	v4710 = v4707
	goto L918
L918:
	;
	v4711 = int32(8)
	v4716 = int32(0)
	goto L919
L919:
	;
	v4723 = v4716 << (uint(int32(2)) % 32)
	v4725 = *(*int32)(unsafe.Add(mBase, uint32(v4693+v4711+v4723)))
	v4727 = *(*int32)(unsafe.Add(mBase, uint32(v4692+v4711+v4723)))
	v4728 = v4725 & v4727
	v4730 = base.B2i32(v4728 != int32(0))
	if v4728 != 0 {
		v4739 = v4730
		goto L911
	} else {
		goto L921
	}
L920:
	;
	v4739 = v4730
	goto L911
L921:
	;
	v4732 = v4716 + int32(1)
	if v4732 != v4710 {
		v4716 = v4732
		goto L919
	} else {
		goto L922
	}
L922:
	;
	goto L920
L923:
	;
	goto L893
L924:
	;
	v5203 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l1)+26)))
	if v5203 != int32(1) {
		goto L762
	} else {
		goto L1018
	}
L925:
	;
	if v4533 != 0 {
		goto L934
	} else {
		goto L935
	}
L926:
	;
	v4752 = int32(0)
	v4753 = *(*int32)(unsafe.Add(mBase, uint32(v39)+20))
	v4754 = F_create_unique_path(m, l0, l3, v4532, v4753)
	mBase = m.M
	v4755 = m.ExcPending
	if v4755 != 0 {
		goto L22
	} else {
		goto L930
	}
L927:
	;
	v4744 = *(*int32)(unsafe.Add(mBase, uint32(v39)+20))
	v4745 = F_create_unique_path(m, l0, l2, v4534, v4744)
	mBase = m.M
	v4746 = m.ExcPending
	if v4746 != 0 {
		goto L22
	} else {
		goto L928
	}
L928:
	;
	F_try_hashjoin_path(m, l0, l1, v4745, v4532, v4524, int32(0), v39+int32(8))
	mBase = m.M
	v4751 = m.ExcPending
	if v4751 != 0 {
		goto L22
	} else {
		goto L929
	}
L929:
	;
	v5173 = int32(0)
	v5186 = v4532
	goto L924
L930:
	;
	v4758 = v39 + int32(8)
	F_try_hashjoin_path(m, l0, l1, v4534, v4754, v4524, int32(0), v4758)
	mBase = m.M
	v4760 = m.ExcPending
	if v4760 != 0 {
		goto L22
	} else {
		goto L931
	}
L931:
	;
	if base.B2i32(v4533 == int32(0))|base.B2i32(v4534 == v4533) != 0 {
		v5173 = v4752
		v5186 = v4754
		goto L924
	} else {
		goto L932
	}
L932:
	;
	F_try_hashjoin_path(m, l0, l1, v4533, v4754, v4524, int32(0), v4758)
	mBase = m.M
	v4767 = m.ExcPending
	if v4767 != 0 {
		goto L22
	} else {
		goto L933
	}
L933:
	;
	v5173 = v4752
	v5186 = v4754
	goto L924
L934:
	;
	F_try_hashjoin_path(m, l0, l1, v4533, v4532, v4524, l4, v39+int32(8))
	mBase = m.M
	v4771 = m.ExcPending
	if v4771 != 0 {
		goto L22
	} else {
		goto L937
	}
L935:
	;
	goto L936
L936:
	;
	v4772 = *(*int32)(unsafe.Add(mBase, uint32(l2)+56))
	if v4772 == int32(0) {
		goto L938
	} else {
		goto L939
	}
L937:
	;
	goto L936
L938:
	;
	v5173 = l4
	v5186 = v4532
	goto L924
L939:
	;
	v4775 = *(*int32)(unsafe.Add(mBase, uint32(v4772)+4))
	if v4775 <= int32(0) {
		goto L938
	} else {
		goto L940
	}
L940:
	;
	v4788 = int32(0)
	goto L941
L941:
	;
	v4815 = *(*int32)(unsafe.Add(mBase, uint32(v4772)+12))
	v4819 = *(*int32)(unsafe.Add(mBase, uint32(v4815+v4788<<(uint(int32(2))%32))))
	v4820 = *(*int32)(unsafe.Add(mBase, uint32(v4819)+16))
	if v4820 == int32(0) {
		goto L944
	} else {
		goto L945
	}
L942:
	;
	goto L938
L943:
	;
	v5128 = v4788 + int32(1)
	v5129 = *(*int32)(unsafe.Add(mBase, uint32(v4772)+4))
	if v5128 < v5129 {
		v4788 = v5128
		goto L941
	} else {
		goto L1017
	}
L944:
	;
	v4923 = *(*int32)(unsafe.Add(mBase, uint32(l3)+56))
	if v4923 == int32(0) {
		goto L943
	} else {
		goto L975
	}
L945:
	;
	v4823 = *(*int32)(unsafe.Add(mBase, uint32(v4820)+4))
	v4824 = *(*int32)(unsafe.Add(mBase, uint32(l3)+8))
	v4825 = int32(0)
	if base.B2i32(v4823 == v4825)|base.B2i32(v4824 == v4825) != 0 {
		v4870 = v4825
		goto L947
	} else {
		goto L948
	}
L946:
	;
	if v4870 != 0 {
		goto L943
	} else {
		goto L959
	}
L947:
	;
	goto L946
L948:
	;
	v4835 = *(*int32)(unsafe.Add(mBase, uint32(v4823)+4))
	v4836 = *(*int32)(unsafe.Add(mBase, uint32(v4824)+4))
	if v4835 < v4836 {
		goto L949
	} else {
		goto L950
	}
L949:
	;
	v4838 = v4835
	goto L951
L950:
	;
	v4838 = v4836
	goto L951
L951:
	;
	if v4838 <= int32(1) {
		goto L952
	} else {
		goto L953
	}
L952:
	;
	v4841 = int32(1)
	goto L954
L953:
	;
	v4841 = v4838
	goto L954
L954:
	;
	v4842 = int32(8)
	v4847 = int32(0)
	goto L955
L955:
	;
	v4854 = v4847 << (uint(int32(2)) % 32)
	v4856 = *(*int32)(unsafe.Add(mBase, uint32(v4824+v4842+v4854)))
	v4858 = *(*int32)(unsafe.Add(mBase, uint32(v4823+v4842+v4854)))
	v4859 = v4856 & v4858
	v4861 = base.B2i32(v4859 != int32(0))
	if v4859 != 0 {
		v4870 = v4861
		goto L947
	} else {
		goto L957
	}
L956:
	;
	v4870 = v4861
	goto L947
L957:
	;
	v4863 = v4847 + int32(1)
	if v4863 != v4841 {
		v4847 = v4863
		goto L955
	} else {
		goto L958
	}
L958:
	;
	goto L956
L959:
	;
	v4871 = *(*int32)(unsafe.Add(mBase, uint32(v4819)+16))
	if v4871 == int32(0) {
		goto L944
	} else {
		goto L960
	}
L960:
	;
	v4874 = *(*int32)(unsafe.Add(mBase, uint32(v4871)+4))
	v4875 = *(*int32)(unsafe.Add(mBase, uint32(l3)+228))
	v4876 = int32(0)
	if base.B2i32(v4874 == v4876)|base.B2i32(v4875 == v4876) != 0 {
		v4921 = v4876
		goto L962
	} else {
		goto L963
	}
L961:
	;
	if v4921 != 0 {
		goto L943
	} else {
		goto L974
	}
L962:
	;
	goto L961
L963:
	;
	v4886 = *(*int32)(unsafe.Add(mBase, uint32(v4874)+4))
	v4887 = *(*int32)(unsafe.Add(mBase, uint32(v4875)+4))
	if v4886 < v4887 {
		goto L964
	} else {
		goto L965
	}
L964:
	;
	v4889 = v4886
	goto L966
L965:
	;
	v4889 = v4887
	goto L966
L966:
	;
	if v4889 <= int32(1) {
		goto L967
	} else {
		goto L968
	}
L967:
	;
	v4892 = int32(1)
	goto L969
L968:
	;
	v4892 = v4889
	goto L969
L969:
	;
	v4893 = int32(8)
	v4898 = int32(0)
	goto L970
L970:
	;
	v4905 = v4898 << (uint(int32(2)) % 32)
	v4907 = *(*int32)(unsafe.Add(mBase, uint32(v4875+v4893+v4905)))
	v4909 = *(*int32)(unsafe.Add(mBase, uint32(v4874+v4893+v4905)))
	v4910 = v4907 & v4909
	v4912 = base.B2i32(v4910 != int32(0))
	if v4910 != 0 {
		v4921 = v4912
		goto L962
	} else {
		goto L972
	}
L971:
	;
	v4921 = v4912
	goto L962
L972:
	;
	v4914 = v4898 + int32(1)
	if v4914 != v4892 {
		v4898 = v4914
		goto L970
	} else {
		goto L973
	}
L973:
	;
	goto L971
L974:
	;
	goto L944
L975:
	;
	v4926 = int32(0)
	v4927 = *(*int32)(unsafe.Add(mBase, uint32(v4923)+4))
	if v4927 <= v4926 {
		goto L943
	} else {
		goto L976
	}
L976:
	;
	v4938 = v4926
	goto L977
L977:
	;
	v4966 = *(*int32)(unsafe.Add(mBase, uint32(v4923)+12))
	v4970 = *(*int32)(unsafe.Add(mBase, uint32(v4966+v4938<<(uint(int32(2))%32))))
	v4971 = *(*int32)(unsafe.Add(mBase, uint32(v4970)+16))
	if v4971 == int32(0) {
		goto L981
	} else {
		goto L982
	}
L978:
	;
	goto L943
L979:
	;
	v5088 = v4938 + int32(1)
	v5089 = *(*int32)(unsafe.Add(mBase, uint32(v4923)+4))
	if v5088 < v5089 {
		v4938 = v5088
		goto L977
	} else {
		goto L1016
	}
L980:
	;
	F_try_hashjoin_path(m, l0, l1, v4819, v4970, v4524, l4, v39+int32(8))
	mBase = m.M
	v5085 = m.ExcPending
	if v5085 != 0 {
		goto L22
	} else {
		goto L1015
	}
L981:
	;
	if v4533 != v4819 {
		goto L980
	} else {
		goto L1013
	}
L982:
	;
	v4974 = *(*int32)(unsafe.Add(mBase, uint32(v4971)+4))
	v4975 = *(*int32)(unsafe.Add(mBase, uint32(l2)+8))
	v4976 = int32(0)
	if base.B2i32(v4974 == v4976)|base.B2i32(v4975 == v4976) != 0 {
		v5021 = v4976
		goto L984
	} else {
		goto L985
	}
L983:
	;
	if v5021 != 0 {
		goto L979
	} else {
		goto L996
	}
L984:
	;
	goto L983
L985:
	;
	v4986 = *(*int32)(unsafe.Add(mBase, uint32(v4974)+4))
	v4987 = *(*int32)(unsafe.Add(mBase, uint32(v4975)+4))
	if v4986 < v4987 {
		goto L986
	} else {
		goto L987
	}
L986:
	;
	v4989 = v4986
	goto L988
L987:
	;
	v4989 = v4987
	goto L988
L988:
	;
	if v4989 <= int32(1) {
		goto L989
	} else {
		goto L990
	}
L989:
	;
	v4992 = int32(1)
	goto L991
L990:
	;
	v4992 = v4989
	goto L991
L991:
	;
	v4993 = int32(8)
	v4998 = int32(0)
	goto L992
L992:
	;
	v5005 = v4998 << (uint(int32(2)) % 32)
	v5007 = *(*int32)(unsafe.Add(mBase, uint32(v4975+v4993+v5005)))
	v5009 = *(*int32)(unsafe.Add(mBase, uint32(v4974+v4993+v5005)))
	v5010 = v5007 & v5009
	v5012 = base.B2i32(v5010 != int32(0))
	if v5010 != 0 {
		v5021 = v5012
		goto L984
	} else {
		goto L994
	}
L993:
	;
	v5021 = v5012
	goto L984
L994:
	;
	v5014 = v4998 + int32(1)
	if v5014 != v4992 {
		v4998 = v5014
		goto L992
	} else {
		goto L995
	}
L995:
	;
	goto L993
L996:
	;
	v5022 = *(*int32)(unsafe.Add(mBase, uint32(v4970)+16))
	if v5022 == int32(0) {
		goto L981
	} else {
		goto L997
	}
L997:
	;
	v5025 = *(*int32)(unsafe.Add(mBase, uint32(v5022)+4))
	v5026 = *(*int32)(unsafe.Add(mBase, uint32(l2)+228))
	v5027 = int32(0)
	if base.B2i32(v5025 == v5027)|base.B2i32(v5026 == v5027) != 0 {
		v5072 = v5027
		goto L999
	} else {
		goto L1000
	}
L998:
	;
	if v5072 != 0 {
		goto L979
	} else {
		goto L1011
	}
L999:
	;
	goto L998
L1000:
	;
	v5037 = *(*int32)(unsafe.Add(mBase, uint32(v5025)+4))
	v5038 = *(*int32)(unsafe.Add(mBase, uint32(v5026)+4))
	if v5037 < v5038 {
		goto L1001
	} else {
		goto L1002
	}
L1001:
	;
	v5040 = v5037
	goto L1003
L1002:
	;
	v5040 = v5038
	goto L1003
L1003:
	;
	if v5040 <= int32(1) {
		goto L1004
	} else {
		goto L1005
	}
L1004:
	;
	v5043 = int32(1)
	goto L1006
L1005:
	;
	v5043 = v5040
	goto L1006
L1006:
	;
	v5044 = int32(8)
	v5049 = int32(0)
	goto L1007
L1007:
	;
	v5056 = v5049 << (uint(int32(2)) % 32)
	v5058 = *(*int32)(unsafe.Add(mBase, uint32(v5026+v5044+v5056)))
	v5060 = *(*int32)(unsafe.Add(mBase, uint32(v5025+v5044+v5056)))
	v5061 = v5058 & v5060
	v5063 = base.B2i32(v5061 != int32(0))
	if v5061 != 0 {
		v5072 = v5063
		goto L999
	} else {
		goto L1009
	}
L1008:
	;
	v5072 = v5063
	goto L999
L1009:
	;
	v5065 = v5049 + int32(1)
	if v5065 != v5043 {
		v5049 = v5065
		goto L1007
	} else {
		goto L1010
	}
L1010:
	;
	goto L1008
L1011:
	;
	if base.B2i32(v4533 == v4819)&base.B2i32(v4532 == v4970) == int32(0) {
		goto L980
	} else {
		goto L1012
	}
L1012:
	;
	goto L979
L1013:
	;
	if v4532 == v4970 {
		goto L979
	} else {
		goto L1014
	}
L1014:
	;
	goto L980
L1015:
	;
	goto L979
L1016:
	;
	goto L978
L1017:
	;
	goto L942
L1018:
	;
	switch l4 - int32(6) {
	case 0, 2:
		goto L762
	default:
		goto L1019
	}
L1019:
	;
	v5208 = *(*int32)(unsafe.Add(mBase, uint32(l2)+40))
	if v5208 == int32(0) {
		goto L762
	} else {
		goto L1020
	}
L1020:
	;
	v5211 = *(*int32)(unsafe.Add(mBase, uint32(l1)+64))
	if v5211 != 0 {
		goto L762
	} else {
		goto L1021
	}
L1021:
	;
	v5212 = *(*int32)(unsafe.Add(mBase, uint32(v5208)+12))
	v5213 = *(*int32)(unsafe.Add(mBase, uint32(v5212)))
	if l4 == int32(9) {
		goto L1022
	} else {
		goto L1023
	}
L1022:
	;
	if int32(1)<<(uint(l4)%32)&int32(140) != 0 {
		goto L1027
	} else {
		goto L1028
	}
L1023:
	;
	v5216 = *(*int32)(unsafe.Add(mBase, uint32(l3)+40))
	if v5216 == int32(0) {
		goto L1022
	} else {
		goto L1024
	}
L1024:
	;
	v5220 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_add_paths_to_joinrel[3])))
	if v5220&int32(1) == int32(0) {
		goto L1022
	} else {
		goto L1025
	}
L1025:
	;
	v5225 = *(*int32)(unsafe.Add(mBase, uint32(v5216)+12))
	v5226 = *(*int32)(unsafe.Add(mBase, uint32(v5225)))
	F_try_partial_hashjoin_path(m, l0, l1, v5213, v5226, v4524, v5173, v39+int32(8), int32(1))
	mBase = m.M
	v5231 = m.ExcPending
	if v5231 != 0 {
		goto L22
	} else {
		goto L1026
	}
L1026:
	;
	goto L1022
L1027:
	;
	v5240 = base.B2i32(base.Ui32(l4) <= base.Ui32(int32(7)))
	goto L1029
L1028:
	;
	v5240 = int32(0)
	goto L1029
L1029:
	;
	if v5240 != 0 {
		goto L762
	} else {
		goto L1030
	}
L1030:
	;
	v5242 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v5186)+21)))
	if v5242 != 0 {
		goto L1031
	} else {
		goto L1032
	}
L1031:
	;
	v5243 = v5186
	goto L1033
L1032:
	;
	v5243 = int32(0)
	goto L1033
L1033:
	;
	if base.B2i32(l4 == int32(9))|v5242 == int32(0) {
		goto L1034
	} else {
		goto L1035
	}
L1034:
	;
	v5249 = *(*int32)(unsafe.Add(mBase, uint32(l3)+32))
	if v5249 != 0 {
		goto L1039
	} else {
		goto L1040
	}
L1035:
	;
	v5296 = v5243
	goto L1036
L1036:
	;
	if v5296 == int32(0) {
		goto L762
	} else {
		goto L1054
	}
L1037:
	;
	v5296 = v5293
	goto L1036
L1038:
	;
	goto L1037
L1039:
	;
	v5254 = *(*int32)(unsafe.Add(mBase, uint32(v5249)+4))
	if v5254 <= int32(0) {
		v5293 = int32(0)
		goto L1038
	} else {
		goto L1042
	}
L1040:
	;
	goto L1041
L1041:
	;
	v5293 = int32(0)
	goto L1038
L1042:
	;
	v5257 = int32(0)
	if v5257 < v5254 {
		goto L1043
	} else {
		goto L1044
	}
L1043:
	;
	v5260 = v5254
	goto L1045
L1044:
	;
	v5260 = v5257
	goto L1045
L1045:
	;
	v5261 = *(*int32)(unsafe.Add(mBase, uint32(v5249)+12))
	v5263 = int32(0)
	goto L1046
L1046:
	;
	v5271 = *(*int32)(unsafe.Add(mBase, uint32(v5261+v5263<<(uint(int32(2))%32))))
	v5272 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v5271)+21)))
	if v5272 == int32(1) {
		goto L1048
	} else {
		goto L1049
	}
L1047:
	;
	goto L1041
L1048:
	;
	v5275 = *(*int32)(unsafe.Add(mBase, uint32(v5271)+16))
	if v5275 == int32(0) {
		v5293 = v5271
		goto L1038
	} else {
		goto L1051
	}
L1049:
	;
	goto L1050
L1050:
	;
	v5283 = v5263 + int32(1)
	if v5283 != v5260 {
		v5263 = v5283
		goto L1046
	} else {
		goto L1053
	}
L1051:
	;
	v5278 = *(*int32)(unsafe.Add(mBase, uint32(v5275)+4))
	if v5278 == int32(0) {
		v5293 = v5271
		goto L1038
	} else {
		goto L1052
	}
L1052:
	;
	goto L1050
L1053:
	;
	goto L1047
L1054:
	;
	F_try_partial_hashjoin_path(m, l0, l1, v5213, v5296, v4524, v5173, v39+int32(8), int32(0))
	mBase = m.M
	v5303 = m.ExcPending
	if v5303 != 0 {
		goto L22
	} else {
		goto L1055
	}
L1055:
	;
	goto L762
L1056:
	;
	v5352 = *(*int32)(unsafe.Add(mBase, _c_F_add_paths_to_joinrel[4]))
	if v5352 != 0 {
		goto L1060
	} else {
		goto L1061
	}
L1057:
	;
	v5343 = *(*int32)(unsafe.Add(mBase, uint32(v5340)+32))
	if v5343 == int32(0) {
		goto L1056
	} else {
		goto L1058
	}
L1058:
	;
	m.T0[v5343].(func(*base.Module, int32, int32, int32, int32, int32, int32))(m, l0, l1, l2, l3, l4, v39+int32(8))
	mBase = m.M
	v5349 = m.ExcPending
	if v5349 != 0 {
		goto L22
	} else {
		goto L1059
	}
L1059:
	;
	goto L1056
L1060:
	;
	m.T0[v5352].(func(*base.Module, int32, int32, int32, int32, int32, int32))(m, l0, l1, l2, l3, l4, v39+int32(8))
	mBase = m.M
	v5356 = m.ExcPending
	if v5356 != 0 {
		goto L22
	} else {
		goto L1063
	}
L1061:
	;
	goto L1062
L1062:
	;
	m.G0 = v39 + int32(48)
	return
L1063:
	;
	goto L1062
}
func F_add_pos(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32) int32 {
	mBase := m.M
	_ = mBase
	var v14 int32
	_ = v14
	var v18 int32
	_ = v18
	var v21 int32
	_ = v21
	var v30 int32
	_ = v30
	var v31 int32
	_ = v31
	var v36 int32
	_ = v36
	var v47 int32
	_ = v47
	var v49 int32
	_ = v49
	var v59 int32
	_ = v59
	var v60 int32
	_ = v60
	var v65 int32
	_ = v65
	var v66 int32
	_ = v66
	var v67 int32
	_ = v67
	var v68 int32
	_ = v68
	var v70 int32
	_ = v70
	var v73 int32
	_ = v73
	var v74 int32
	_ = v74
	var v77 int32
	_ = v77
	var v90 int32
	_ = v90
	var v91 int32
	_ = v91
	var v92 int32
	_ = v92
	var v96 int32
	_ = v96
	var v98 int32
	_ = v98
	var v99 int32
	_ = v99
	var v102 int32
	_ = v102
	var v113 int32
	_ = v113
	var v114 int32
	_ = v114
	var v118 int32
	_ = v118
	var v120 int32
	_ = v120
	var v123 int32
	_ = v123
	var v124 int32
	_ = v124
	var v126 int32
	_ = v126
	var v127 int32
	_ = v127
	var v128 int32
	_ = v128
	var v130 int32
	_ = v130
	var v133 int32
	_ = v133
	var v136 int32
	_ = v136
	var v139 int32
	_ = v139
	var v140 int32
	_ = v140
	var v142 int32
	_ = v142
	var v143 int32
	_ = v143
	var v146 int32
	_ = v146
	var v154 int32
	_ = v154
	var v165 int32
	_ = v165
	var v169 int32
	_ = v169
	var v172 int32
	_ = v172
	var v187 int32
	_ = v187
	v14 = *(*int32)(unsafe.Add(mBase, uint32(l2)+4))
	v18 = *(*int32)(unsafe.Add(mBase, uint32(l3)))
	v21 = int32(1)
	v30 = l2 + v14<<(uint(int32(2))%32) + (int32(base.Ui32(v18)>>(uint(int32(12))%32))+int32(base.Ui32(v18)>>(uint(v21)%32))&int32(2047)+v21)&int32(_a_F_add_pos_0)
	v31 = *(*int32)(unsafe.Add(mBase, uint32(l1)))
	if v31&v21 == int32(0) {
		goto L2
	} else {
		goto L3
	}
L1:
	;
	v70 = v30 + int32(8)
	if v18&int32(1) != 0 {
		goto L6
	} else {
		goto L7
	}
L2:
	;
	v36 = int32(1)
	v47 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v66 = (int32(base.Ui32(v31)>>(uint(v36)%32))&int32(2047) + int32(base.Ui32(v31)>>(uint(int32(12))%32)) + v36) & int32(_a_F_add_pos_0)
	v67 = v47
	v68 = int32(0)
	goto L1
L3:
	;
	goto L4
L4:
	;
	v49 = int32(1)
	v59 = (int32(base.Ui32(v31)>>(uint(v49)%32))&int32(2047) + int32(base.Ui32(v31)>>(uint(int32(12))%32)) + v49) & int32(_a_F_add_pos_0)
	v60 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v65 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v59+(l0+v60<<(uint(int32(2))%32)))+8)))
	v66 = v59
	v67 = v60
	v68 = v65
	goto L1
L5:
	;
	if v68 == int32(0) {
		v172 = v77
		goto L9
	} else {
		goto L10
	}
L6:
	;
	v73 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v70))))
	v77 = v73
	goto L5
L7:
	;
	goto L8
L8:
	;
	v74 = int32(0)
	*(*uint16)(unsafe.Add(mBase, uint32(v70))) = uint16(v74)
	v77 = v74
	goto L5
L9:
	;
	return v172&int32(_a_F_add_pos_1) - v77
L10:
	;
	if base.Ui32(int32(255)) < base.Ui32(v77) {
		v154 = v77
		goto L11
	} else {
		goto L12
	}
L11:
	;
	if v77 == v154&int32(_a_F_add_pos_1) {
		v172 = v77
		goto L9
	} else {
		goto L29
	}
L12:
	;
	v187 = int32(10)
	v90 = int32(0)
	v91 = int32(256)
	v92 = v91 - v77
	if base.Ui32(v92) <= base.Ui32(v91) {
		goto L13
	} else {
		goto L14
	}
L13:
	;
	v96 = v92
	goto L15
L14:
	;
	v96 = v90
	goto L15
L15:
	;
	v98 = v90
	v99 = v77
	v102 = v77
	goto L16
L16:
	;
	if v99 != 0 {
		goto L18
	} else {
		goto L19
	}
L17:
	;
	v154 = v143
	goto L11
L18:
	;
	v113 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v70+v99<<(uint(int32(1))%32)))))
	v114 = int32(_a_F_add_pos_2)
	if v113&v114 == v114 {
		v154 = v102
		goto L11
	} else {
		goto L21
	}
L19:
	;
	goto L20
L20:
	;
	v118 = int32(1)
	v120 = v30 + v187 + v99<<(uint(v118)%32)
	v123 = l0 + v67<<(uint(int32(2))%32) + v66 + v187 + v98<<(uint(v118)%32)
	v124 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v123))))
	v126 = v124 & int32(-16384)
	v127 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v120))))
	v128 = int32(_a_F_add_pos_2)
	v130 = v126 | v127&v128
	*(*uint16)(unsafe.Add(mBase, uint32(v120))) = uint16(v130)
	v133 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v123))))
	v136 = l4 + v133&v128
	if base.Ui32(v128) <= base.Ui32(v136) {
		goto L22
	} else {
		goto L23
	}
L21:
	;
	goto L20
L22:
	;
	v139 = v128
	goto L24
L23:
	;
	v139 = v136
	goto L24
L24:
	;
	v140 = v126 | v139
	*(*uint16)(unsafe.Add(mBase, uint32(v120))) = uint16(v140)
	v142 = int32(1)
	v143 = v99 + v142
	*(*uint16)(unsafe.Add(mBase, uint32(v70))) = uint16(v143)
	v146 = v98 + v142
	if v68 != v146 {
		goto L25
	} else {
		goto L26
	}
L25:
	;
	if v146 == v96 {
		v154 = v143
		goto L11
	} else {
		goto L28
	}
L26:
	;
	goto L27
L27:
	;
	goto L17
L28:
	;
	v98 = v146
	v99 = v143
	v102 = v143
	goto L16
L29:
	;
	v165 = *(*int32)(unsafe.Add(mBase, uint32(l3)))
	*(*int32)(unsafe.Add(mBase, uint32(l3))) = v165 | int32(1)
	v169 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v70))))
	v172 = v169
	goto L9
}
func F_add_values_to_range(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32) int32 {
	mBase := m.M
	_ = mBase
	var v6 int32
	_ = v6
	var v13 int32
	_ = v13
	var v14 int32
	_ = v14
	var v15 int32
	_ = v15
	var v28 int32
	_ = v28
	var v29 int32
	_ = v29
	var v36 int32
	_ = v36
	var v38 int32
	_ = v38
	var v40 int32
	_ = v40
	var v41 int32
	_ = v41
	var v42 int32
	_ = v42
	var v44 int32
	_ = v44
	var v46 int32
	_ = v46
	var v47 int32
	_ = v47
	var v51 int32
	_ = v51
	var v54 int32
	_ = v54
	var v55 int32
	_ = v55
	var v62 int32
	_ = v62
	var v65 int32
	_ = v65
	var v66 int32
	_ = v66
	var v68 int32
	_ = v68
	var v70 int32
	_ = v70
	var v72 int32
	_ = v72
	var v73 int32
	_ = v73
	var v74 int32
	_ = v74
	var v75 int32
	_ = v75
	var v77 int32
	_ = v77
	var v82 int32
	_ = v82
	var v83 int32
	_ = v83
	var v84 int32
	_ = v84
	var v86 int32
	_ = v86
	var v88 int32
	_ = v88
	var v89 int32
	_ = v89
	var v90 int32
	_ = v90
	var v99 int32
	_ = v99
	var v104 int32
	_ = v104
	v6 = int32(0)
	v13 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l2)+1)))
	v14 = *(*int32)(unsafe.Add(mBase, uint32(l1)+8))
	v15 = *(*int32)(unsafe.Add(mBase, uint32(v14)))
	if v6 < v15 {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	v28 = v6
	v29 = v13
	goto L4
L2:
	;
	v99 = v13
	goto L3
L3:
	;
	v104 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(l2)+1)) = uint8(v104)
	return v99 & int32(1)
L4:
	;
	v36 = l2 + int32(24) + v28*int32(20)
	v38 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l2)+1)))
	if v38 != 0 {
		v42 = int32(0)
		goto L6
	} else {
		goto L7
	}
L5:
	;
	v99 = v86
	goto L3
L6:
	;
	v44 = v28 << (uint(int32(2)) % 32)
	v46 = *(*int32)(unsafe.Add(mBase, uint32(l1+int32(20)+v44)))
	v47 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v46)+2)))
	if v47 != int32(1) {
		goto L10
	} else {
		goto L11
	}
L7:
	;
	v40 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v36)+2)))
	if v40 != 0 {
		v42 = int32(1)
		goto L6
	} else {
		goto L8
	}
L8:
	;
	v41 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v36)+3)))
	v42 = v41
	goto L6
L9:
	;
	v88 = v28 + int32(1)
	v89 = *(*int32)(unsafe.Add(mBase, uint32(l1)+8))
	v90 = *(*int32)(unsafe.Add(mBase, uint32(v89)))
	if v88 < v90 {
		v28 = v88
		v29 = v86
		goto L4
	} else {
		goto L20
	}
L10:
	;
	v62 = F_index_getprocinfo(m, l0, base.I32_extend16_s(v28+int32(1)), int32(2))
	mBase = m.M
	v65 = m.ExcPending
	if v65 != 0 {
		goto L14
	} else {
		goto L15
	}
L11:
	;
	v51 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l4+v28))))
	if v51 != int32(1) {
		goto L10
	} else {
		goto L12
	}
L12:
	;
	v54 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v36)+2)))
	if v54 != 0 {
		v86 = v29
		goto L9
	} else {
		goto L13
	}
L13:
	;
	v55 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(v36)+2)) = uint8(v55)
	v86 = v55
	goto L9
L14:
	;
	return int32(0)
L15:
	;
	v66 = *(*int32)(unsafe.Add(mBase, uint32(l0)+248))
	v68 = *(*int32)(unsafe.Add(mBase, uint32(v66+v44)))
	v70 = *(*int32)(unsafe.Add(mBase, uint32(l3+v44)))
	v72 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l4+v28))))
	v73 = F_FunctionCall4Coll(m, v62, v68, l1, v36, v70, v72)
	mBase = m.M
	v74 = m.ExcPending
	if v74 != 0 {
		goto L14
	} else {
		goto L16
	}
L16:
	;
	v75 = int32(0)
	v77 = v29 | base.B2i32(v73 != v75)
	if v42&int32(1) == v75 {
		v86 = v77
		goto L9
	} else {
		goto L17
	}
L17:
	;
	v82 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v36)+2)))
	if v82 != 0 {
		v86 = v77
		goto L9
	} else {
		goto L18
	}
L18:
	;
	v83 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v36)+3)))
	if v83 != 0 {
		v86 = v77
		goto L9
	} else {
		goto L19
	}
L19:
	;
	v84 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(v36)+2)) = uint8(v84)
	v86 = v77
	goto L9
L20:
	;
	goto L5
}
func F_add_vars_to_attr_needed(m *base.Module, l0 int32, l1 int32, l2 int32) {
	mBase := m.M
	_ = mBase
	var v4 int32
	_ = v4
	var v9 int32
	_ = v9
	var v11 int32
	_ = v11
	var v15 int32
	_ = v15
	var v24 int32
	_ = v24
	var v26 int32
	_ = v26
	var v30 int32
	_ = v30
	var v31 int32
	_ = v31
	var v36 int32
	_ = v36
	var v37 int32
	_ = v37
	var v38 int32
	_ = v38
	var v39 int32
	_ = v39
	var v40 int32
	_ = v40
	var v42 int32
	_ = v42
	var v43 int32
	_ = v43
	var v44 int32
	_ = v44
	var v45 int32
	_ = v45
	var v46 int32
	_ = v46
	var v47 int32
	_ = v47
	var v56 int32
	_ = v56
	var v57 int32
	_ = v57
	var v59 int32
	_ = v59
	var v62 int32
	_ = v62
	var v63 int32
	_ = v63
	var v68 int32
	_ = v68
	var v75 int32
	_ = v75
	var v77 int32
	_ = v77
	var v79 int32
	_ = v79
	var v82 int32
	_ = v82
	var v84 int32
	_ = v84
	var v86 int32
	_ = v86
	var v93 int32
	_ = v93
	var v100 int32
	_ = v100
	var v101 int32
	_ = v101
	var v104 int32
	_ = v104
	var v105 int32
	_ = v105
	var v107 int32
	_ = v107
	var v108 int32
	_ = v108
	var v109 int32
	_ = v109
	var v110 int32
	_ = v110
	var v117 int32
	_ = v117
	var v118 int32
	_ = v118
	var v134 int32
	_ = v134
	var v135 int32
	_ = v135
	var v139 int32
	_ = v139
	var v144 int32
	_ = v144
	v4 = int32(0)
	v9 = m.G0
	v11 = v9 - int32(16)
	m.G0 = v11
	if l1 == v4 {
		goto L2
	} else {
		goto L3
	}
L1:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v134 = m.ExcPending
	if v134 != 0 {
		goto L12
	} else {
		goto L33
	}
L2:
	;
	m.G0 = v11 + int32(16)
	return
L3:
	;
	v15 = *(*int32)(unsafe.Add(mBase, uint32(l1)+4))
	if v15 <= int32(0) {
		goto L2
	} else {
		goto L4
	}
L4:
	;
	v24 = v4
	goto L5
L5:
	;
	v26 = *(*int32)(unsafe.Add(mBase, uint32(l1)+12))
	v30 = *(*int32)(unsafe.Add(mBase, uint32(v26+v24<<(uint(int32(2))%32))))
	v31 = *(*int32)(unsafe.Add(mBase, uint32(v30)))
	if v31 != int32(6) {
		goto L8
	} else {
		goto L9
	}
L6:
	;
	goto L2
L7:
	;
	v117 = v24 + int32(1)
	v118 = *(*int32)(unsafe.Add(mBase, uint32(l1)+4))
	if v117 < v118 {
		v24 = v117
		goto L5
	} else {
		goto L32
	}
L8:
	;
	if v31 != int32(319) {
		goto L1
	} else {
		goto L11
	}
L9:
	;
	goto L10
L10:
	;
	v42 = *(*int32)(unsafe.Add(mBase, uint32(v30)+4))
	v43 = F_find_base_rel(m, l0, v42)
	mBase = m.M
	v44 = m.ExcPending
	if v44 != 0 {
		goto L12
	} else {
		goto L15
	}
L11:
	;
	v36 = F_find_placeholder_info(m, l0, v30)
	mBase = m.M
	v37 = m.ExcPending
	if v37 != 0 {
		goto L12
	} else {
		goto L13
	}
L12:
	;
	return
L13:
	;
	v38 = *(*int32)(unsafe.Add(mBase, uint32(v36)+20))
	v39 = F_bms_add_members(m, v38, l2)
	mBase = m.M
	v40 = m.ExcPending
	if v40 != 0 {
		goto L12
	} else {
		goto L14
	}
L14:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v36)+20)) = v39
	goto L7
L15:
	;
	v45 = int32(*(*int16)(unsafe.Add(mBase, uint32(v30)+8)))
	v46 = *(*int32)(unsafe.Add(mBase, uint32(v43)+8))
	v47 = int32(0)
	if l2 == v47 {
		goto L17
	} else {
		goto L18
	}
L16:
	;
	if v100 != 0 {
		goto L7
	} else {
		goto L30
	}
L17:
	;
	v100 = int32(1)
	goto L16
L18:
	;
	goto L19
L19:
	;
	if v46 == int32(0) {
		v93 = v47
		goto L20
	} else {
		goto L21
	}
L20:
	;
	v100 = v93
	goto L16
L21:
	;
	v56 = *(*int32)(unsafe.Add(mBase, uint32(l2)+4))
	v57 = *(*int32)(unsafe.Add(mBase, uint32(v46)+4))
	if v57 < v56 {
		v93 = v47
		goto L20
	} else {
		goto L22
	}
L22:
	;
	v59 = int32(1)
	if v56 <= v59 {
		goto L23
	} else {
		goto L24
	}
L23:
	;
	v62 = v59
	goto L25
L24:
	;
	v62 = v56
	goto L25
L25:
	;
	v63 = int32(8)
	v68 = int32(0)
	goto L26
L26:
	;
	v75 = v68 << (uint(int32(2)) % 32)
	v77 = *(*int32)(unsafe.Add(mBase, uint32(l2+v63+v75)))
	v79 = *(*int32)(unsafe.Add(mBase, uint32(v46+v63+v75)))
	v82 = v77 & (v79 ^ int32(-1))
	v84 = base.B2i32(v82 == int32(0))
	if v82 != 0 {
		v93 = v84
		goto L20
	} else {
		goto L28
	}
L27:
	;
	v93 = v84
	goto L20
L28:
	;
	v86 = v68 + int32(1)
	if v86 != v62 {
		v68 = v86
		goto L26
	} else {
		goto L29
	}
L29:
	;
	goto L27
L30:
	;
	v101 = int32(*(*int16)(unsafe.Add(mBase, uint32(v43)+80)))
	v104 = (v45 - v101) << (uint(int32(2)) % 32)
	v105 = *(*int32)(unsafe.Add(mBase, uint32(v43)+84))
	v107 = *(*int32)(unsafe.Add(mBase, uint32(v104+v105)))
	v108 = F_bms_add_members(m, v107, l2)
	mBase = m.M
	v109 = m.ExcPending
	if v109 != 0 {
		goto L12
	} else {
		goto L31
	}
L31:
	;
	v110 = *(*int32)(unsafe.Add(mBase, uint32(v43)+84))
	*(*int32)(unsafe.Add(mBase, uint32(v110+v104))) = v108
	goto L7
L32:
	;
	goto L6
L33:
	;
	v135 = *(*int32)(unsafe.Add(mBase, uint32(v30)))
	*(*int32)(unsafe.Add(mBase, uint32(v11))) = v135
	F_errmsg_internal(m, int32(_a_F_add_vars_to_attr_needed_0), v11)
	mBase = m.M
	v139 = m.ExcPending
	if v139 != 0 {
		goto L12
	} else {
		goto L34
	}
L34:
	;
	F_errfinish(m, int32(_a_F_add_vars_to_attr_needed_1), int32(386), int32(_a_F_add_vars_to_attr_needed_2))
	mBase = m.M
	v144 = m.ExcPending
	if v144 != 0 {
		goto L12
	} else {
		goto L35
	}
L35:
	;
	base.Wasm_trap_unreachable()
	for {
	}
}
func F_addunicode(m *base.Module, l0 int32, l1 int32) {
	mBase := m.M
	_ = mBase
	var v7 int32
	_ = v7
	var v9 int32
	_ = v9
	var v15 int32
	_ = v15
	var v16 int32
	_ = v16
	var v17 int32
	_ = v17
	var v18 int32
	_ = v18
	var v32 int32
	_ = v32
	var v34 int32
	_ = v34
	var v36 int32
	_ = v36
	var v37 int32
	_ = v37
	var v38 int32
	_ = v38
	var v39 int32
	_ = v39
	var v40 int32
	_ = v40
	var v42 int32
	_ = v42
	var v45 int32
	_ = v45
	var v50 int32
	_ = v50
	var v52 int32
	_ = v52
	var v53 int32
	_ = v53
	var v54 int32
	_ = v54
	var v55 int32
	_ = v55
	var v56 int32
	_ = v56
	var v57 int32
	_ = v57
	var v59 int32
	_ = v59
	var v60 int32
	_ = v60
	var v61 int32
	_ = v61
	var v62 int32
	_ = v62
	var v63 int32
	_ = v63
	var v66 int32
	_ = v66
	var v67 int32
	_ = v67
	var v75 int32
	_ = v75
	v7 = m.G0
	v9 = v7 - int32(48)
	m.G0 = v9
	if base.Ui32(l0-int32(1)) < base.Ui32(int32(_a_F_addunicode_0)) {
		v15 = int32(_a_F_addunicode_1)
		v16 = *(*int32)(unsafe.Add(mBase, _c_F_addunicode[0]))
		v17 = *(*int32)(unsafe.Add(mBase, uint32(l1)+96))
		v18 = *(*int32)(unsafe.Add(mBase, uint32(v17)))
		*(*int32)(unsafe.Add(mBase, _c_F_addunicode[0])) = v9 + int32(36)
		*(*int32)(unsafe.Add(mBase, uint32(v9)+40)) = int32(499)
		*(*int32)(unsafe.Add(mBase, uint32(v9)+32)) = v18
		*(*int32)(unsafe.Add(mBase, uint32(v9)+28)) = l1
		*(*int32)(unsafe.Add(mBase, uint32(v9)+36)) = v16
		*(*int32)(unsafe.Add(mBase, uint32(v9)+44)) = v9 + int32(28)
		F_pg_unicode_to_server(m, l0, v9)
		mBase = m.M
		v32 = m.ExcPending
		if v32 != 0 {
			return
		} else {
			v34 = *(*int32)(unsafe.Add(mBase, uint32(v9)+36))
			*(*int32)(unsafe.Add(mBase, _c_F_addunicode[0])) = v34
			v36 = *(*int32)(unsafe.Add(mBase, uint32(l1)))
			v37 = *(*int32)(unsafe.Add(mBase, uint32(v36)+28))
			v38 = F_strlen(m, v9)
			mBase = m.M
			v39 = v37 + v38
			v40 = *(*int32)(unsafe.Add(mBase, uint32(v36)+32))
			if v40 <= v39 {
				v42 = int32(1)
				v45 = v39 + v42
				if v45&v39 != 0 {
					v50 = v42 << (uint(int32(32)-base.I32_clz(v45)) % 32)
				} else {
					v50 = v45
				}
				*(*int32)(unsafe.Add(mBase, uint32(v36)+32)) = v50
				v52 = *(*int32)(unsafe.Add(mBase, uint32(l1)))
				v53 = *(*int32)(unsafe.Add(mBase, uint32(v52)+24))
				v54 = *(*int32)(unsafe.Add(mBase, uint32(v52)+32))
				v55 = F_repalloc(m, v53, v54)
				mBase = m.M
				v56 = m.ExcPending
				if v56 != 0 {
					return
				} else {
					v57 = *(*int32)(unsafe.Add(mBase, uint32(l1)))
					*(*int32)(unsafe.Add(mBase, uint32(v57)+24)) = v55
					v59 = *(*int32)(unsafe.Add(mBase, uint32(l1)))
					v60 = *(*int32)(unsafe.Add(mBase, uint32(v59)+28))
					v61 = v59
					v62 = v60
					if v38 != 0 {
						v63 = *(*int32)(unsafe.Add(mBase, uint32(v61)+24))
						base.MemoryCopy(m, v63+v62, v9, v38)
					} else {
					}
					v66 = *(*int32)(unsafe.Add(mBase, uint32(l1)))
					v67 = *(*int32)(unsafe.Add(mBase, uint32(v66)+28))
					*(*int32)(unsafe.Add(mBase, uint32(v66)+28)) = v67 + v38
					m.G0 = v9 + int32(48)
					return
				}
			} else {
				v61 = v36
				v62 = v37
				if v38 != 0 {
					v63 = *(*int32)(unsafe.Add(mBase, uint32(v61)+24))
					base.MemoryCopy(m, v63+v62, v9, v38)
				} else {
				}
				v66 = *(*int32)(unsafe.Add(mBase, uint32(l1)))
				v67 = *(*int32)(unsafe.Add(mBase, uint32(v66)+28))
				*(*int32)(unsafe.Add(mBase, uint32(v66)+28)) = v67 + v38
				m.G0 = v9 + int32(48)
				return
			}
		}
	} else {
		F_scanner_yyerror(m, int32(_a_F_addunicode_2), l1)
		mBase = m.M
		v75 = m.ExcPending
		if v75 != 0 {
			return
		} else {
			base.Wasm_trap_unreachable()
			for {
			}
		}
	}
}
func F_adjust_appendrel_attrs_multilevel(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) int32 {
	mBase := m.M
	_ = mBase
	var v7 int32
	_ = v7
	var v9 int32
	_ = v9
	var v11 int32
	_ = v11
	var v15 int32
	_ = v15
	var v18 int32
	_ = v18
	var v19 int32
	_ = v19
	var v20 int32
	_ = v20
	var v23 int32
	_ = v23
	var v24 int32
	_ = v24
	var v29 int32
	_ = v29
	var v30 int32
	_ = v30
	var v32 int32
	_ = v32
	var v40 int32
	_ = v40
	var v44 int32
	_ = v44
	var v49 int32
	_ = v49
	v7 = m.G0
	v9 = v7 - int32(16)
	m.G0 = v9
	v11 = *(*int32)(unsafe.Add(mBase, uint32(l2)+220))
	if l3 != v11 {
		if v11 == int32(0) {
			F_errstart_cold(m, int32(21), int32(0))
			mBase = m.M
			v40 = m.ExcPending
			if v40 != 0 {
				return int32(0)
			} else {
				F_errmsg_internal(m, int32(_a_F_adjust_appendrel_attrs_multilevel_0), int32(0))
				mBase = m.M
				v44 = m.ExcPending
				if v44 != 0 {
					return int32(0)
				} else {
					F_errfinish(m, int32(_a_F_adjust_appendrel_attrs_multilevel_1), int32(560), int32(_a_F_adjust_appendrel_attrs_multilevel_2))
					mBase = m.M
					v49 = m.ExcPending
					if v49 != 0 {
						return int32(0)
					} else {
						base.Wasm_trap_unreachable()
						for {
						}
					}
				}
			}
		} else {
			v15 = F_adjust_appendrel_attrs_multilevel(m, l0, l1, v11, l3)
			mBase = m.M
			v18 = m.ExcPending
			if v18 != 0 {
				return int32(0)
			} else {
				v19 = v15
				v20 = *(*int32)(unsafe.Add(mBase, uint32(l2)+8))
				v23 = F_find_appinfos_by_relids(m, l0, v20, v9+int32(8))
				mBase = m.M
				v24 = m.ExcPending
				if v24 != 0 {
					return int32(0)
				} else {
					*(*int32)(unsafe.Add(mBase, uint32(v9)+12)) = v23
					*(*int32)(unsafe.Add(mBase, uint32(v9)+4)) = l0
					v29 = F_adjust_appendrel_attrs_mutator(m, v19, v9+int32(4))
					mBase = m.M
					v30 = m.ExcPending
					if v30 != 0 {
						return int32(0)
					} else {
						F_pfree(m, v23)
						mBase = m.M
						v32 = m.ExcPending
						if v32 != 0 {
							return int32(0)
						} else {
							m.G0 = v9 + int32(16)
							return v29
						}
					}
				}
			}
		}
	} else {
		v19 = l1
		v20 = *(*int32)(unsafe.Add(mBase, uint32(l2)+8))
		v23 = F_find_appinfos_by_relids(m, l0, v20, v9+int32(8))
		mBase = m.M
		v24 = m.ExcPending
		if v24 != 0 {
			return int32(0)
		} else {
			*(*int32)(unsafe.Add(mBase, uint32(v9)+12)) = v23
			*(*int32)(unsafe.Add(mBase, uint32(v9)+4)) = l0
			v29 = F_adjust_appendrel_attrs_mutator(m, v19, v9+int32(4))
			mBase = m.M
			v30 = m.ExcPending
			if v30 != 0 {
				return int32(0)
			} else {
				F_pfree(m, v23)
				mBase = m.M
				v32 = m.ExcPending
				if v32 != 0 {
					return int32(0)
				} else {
					m.G0 = v9 + int32(16)
					return v29
				}
			}
		}
	}
}
func F_adjust_child_relids_multilevel(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) int32 {
	mBase := m.M
	_ = mBase
	var v5 int32
	_ = v5
	var v8 int32
	_ = v8
	var v10 int32
	_ = v10
	var v12 int32
	_ = v12
	var v23 int32
	_ = v23
	var v24 int32
	_ = v24
	var v26 int32
	_ = v26
	var v29 int32
	_ = v29
	var v30 int32
	_ = v30
	var v35 int32
	_ = v35
	var v42 int32
	_ = v42
	var v44 int32
	_ = v44
	var v46 int32
	_ = v46
	var v47 int32
	_ = v47
	var v49 int32
	_ = v49
	var v51 int32
	_ = v51
	var v58 int32
	_ = v58
	var v59 int32
	_ = v59
	var v63 int32
	_ = v63
	var v66 int32
	_ = v66
	var v67 int32
	_ = v67
	var v68 int32
	_ = v68
	var v71 int32
	_ = v71
	var v72 int32
	_ = v72
	var v73 int32
	_ = v73
	var v74 int32
	_ = v74
	var v78 int32
	_ = v78
	var v81 int32
	_ = v81
	var v88 int32
	_ = v88
	var v89 int32
	_ = v89
	var v90 int32
	_ = v90
	var v91 int32
	_ = v91
	var v92 int32
	_ = v92
	var v93 int32
	_ = v93
	var v94 int32
	_ = v94
	var v95 int32
	_ = v95
	var v96 int32
	_ = v96
	var v97 int32
	_ = v97
	var v98 int32
	_ = v98
	var v99 int32
	_ = v99
	var v100 int32
	_ = v100
	var v101 int32
	_ = v101
	var v103 int32
	_ = v103
	var v105 int32
	_ = v105
	var v113 int32
	_ = v113
	var v114 int32
	_ = v114
	var v116 int32
	_ = v116
	var v129 int32
	_ = v129
	var v133 int32
	_ = v133
	var v138 int32
	_ = v138
	v5 = int32(0)
	v8 = m.G0
	v10 = v8 - int32(16)
	m.G0 = v10
	v12 = *(*int32)(unsafe.Add(mBase, uint32(l3)+8))
	if base.B2i32(l1 == v5)|base.B2i32(v12 == v5) != 0 {
		v58 = v5
		goto L3
	} else {
		goto L4
	}
L1:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v129 = m.ExcPending
	if v129 != 0 {
		goto L22
	} else {
		goto L45
	}
L2:
	;
	if v58 != 0 {
		goto L15
	} else {
		goto L16
	}
L3:
	;
	goto L2
L4:
	;
	v23 = *(*int32)(unsafe.Add(mBase, uint32(l1)+4))
	v24 = *(*int32)(unsafe.Add(mBase, uint32(v12)+4))
	if v23 < v24 {
		goto L5
	} else {
		goto L6
	}
L5:
	;
	v26 = v23
	goto L7
L6:
	;
	v26 = v24
	goto L7
L7:
	;
	if v26 <= int32(1) {
		goto L8
	} else {
		goto L9
	}
L8:
	;
	v29 = int32(1)
	goto L10
L9:
	;
	v29 = v26
	goto L10
L10:
	;
	v30 = int32(8)
	v35 = int32(0)
	goto L11
L11:
	;
	v42 = v35 << (uint(int32(2)) % 32)
	v44 = *(*int32)(unsafe.Add(mBase, uint32(v12+v30+v42)))
	v46 = *(*int32)(unsafe.Add(mBase, uint32(l1+v30+v42)))
	v47 = v44 & v46
	v49 = base.B2i32(v47 != int32(0))
	if v47 != 0 {
		v58 = v49
		goto L3
	} else {
		goto L13
	}
L12:
	;
	v58 = v49
	goto L3
L13:
	;
	v51 = v35 + int32(1)
	if v51 != v29 {
		v35 = v51
		goto L11
	} else {
		goto L14
	}
L14:
	;
	goto L12
L15:
	;
	v59 = *(*int32)(unsafe.Add(mBase, uint32(l2)+220))
	if l3 != v59 {
		goto L18
	} else {
		goto L19
	}
L16:
	;
	v116 = l1
	goto L17
L17:
	;
	m.G0 = v10 + int32(16)
	return v116
L18:
	;
	if v59 == int32(0) {
		goto L1
	} else {
		goto L21
	}
L19:
	;
	v67 = l1
	goto L20
L20:
	;
	v68 = *(*int32)(unsafe.Add(mBase, uint32(l2)+8))
	v71 = F_find_appinfos_by_relids(m, l0, v68, v10+int32(12))
	mBase = m.M
	v72 = m.ExcPending
	if v72 != 0 {
		goto L22
	} else {
		goto L24
	}
L21:
	;
	v63 = F_adjust_child_relids_multilevel(m, l0, l1, v59, l3)
	mBase = m.M
	v66 = m.ExcPending
	if v66 != 0 {
		goto L22
	} else {
		goto L23
	}
L22:
	;
	return int32(0)
L23:
	;
	v67 = v63
	goto L20
L24:
	;
	v73 = int32(0)
	v74 = *(*int32)(unsafe.Add(mBase, uint32(v10)+12))
	if v73 < v74 {
		goto L25
	} else {
		goto L26
	}
L25:
	;
	v78 = v73
	v81 = int32(0)
	goto L28
L26:
	;
	v105 = v73
	goto L27
L27:
	;
	F_pfree(m, v71)
	mBase = m.M
	v113 = m.ExcPending
	if v113 != 0 {
		goto L22
	} else {
		goto L41
	}
L28:
	;
	v88 = *(*int32)(unsafe.Add(mBase, uint32(v71+v81<<(uint(int32(2))%32))))
	v89 = *(*int32)(unsafe.Add(mBase, uint32(v88)+4))
	v90 = F_bms_is_member(m, v89, v67)
	mBase = m.M
	v91 = m.ExcPending
	if v91 != 0 {
		goto L22
	} else {
		goto L30
	}
L29:
	;
	v105 = v101
	goto L27
L30:
	;
	if v90 != 0 {
		goto L31
	} else {
		goto L32
	}
L31:
	;
	if v78 != 0 {
		goto L34
	} else {
		goto L35
	}
L32:
	;
	v101 = v78
	goto L33
L33:
	;
	v103 = v81 + int32(1)
	if v103 != v74 {
		v78 = v101
		v81 = v103
		goto L28
	} else {
		goto L40
	}
L34:
	;
	v94 = v78
	goto L36
L35:
	;
	v92 = F_bms_copy(m, v67)
	mBase = m.M
	v93 = m.ExcPending
	if v93 != 0 {
		goto L22
	} else {
		goto L37
	}
L36:
	;
	v95 = *(*int32)(unsafe.Add(mBase, uint32(v88)+4))
	v96 = F_bms_del_member(m, v94, v95)
	mBase = m.M
	v97 = m.ExcPending
	if v97 != 0 {
		goto L22
	} else {
		goto L38
	}
L37:
	;
	v94 = v92
	goto L36
L38:
	;
	v98 = *(*int32)(unsafe.Add(mBase, uint32(v88)+8))
	v99 = F_bms_add_member(m, v96, v98)
	mBase = m.M
	v100 = m.ExcPending
	if v100 != 0 {
		goto L22
	} else {
		goto L39
	}
L39:
	;
	v101 = v99
	goto L33
L40:
	;
	goto L29
L41:
	;
	if v105 != 0 {
		goto L42
	} else {
		goto L43
	}
L42:
	;
	v114 = v105
	goto L44
L43:
	;
	v114 = v67
	goto L44
L44:
	;
	v116 = v114
	goto L17
L45:
	;
	F_errmsg_internal(m, int32(_a_F_adjust_child_relids_multilevel_0), int32(0))
	mBase = m.M
	v133 = m.ExcPending
	if v133 != 0 {
		goto L22
	} else {
		goto L46
	}
L46:
	;
	F_errfinish(m, int32(_a_F_adjust_child_relids_multilevel_1), int32(634), int32(_a_F_adjust_child_relids_multilevel_2))
	mBase = m.M
	v138 = m.ExcPending
	if v138 != 0 {
		goto L22
	} else {
		goto L47
	}
L47:
	;
	base.Wasm_trap_unreachable()
	for {
	}
}
func F_alen_array_element_start(m *base.Module, l0 int32, l1 int32) int32 {
	mBase := m.M
	_ = mBase
	var v3 int32
	_ = v3
	var v4 int32
	_ = v4
	var v7 int32
	_ = v7
	v3 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v4 = *(*int32)(unsafe.Add(mBase, uint32(v3)+32))
	if v4 == int32(1) {
		v7 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
		*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v7 + int32(1)
	} else {
	}
	return int32(0)
}
func F_allcases(m *base.Module, l0 int32, l1 int32) int32 {
	mBase := m.M
	_ = mBase
	var v7 int32
	_ = v7
	var v12 int32
	_ = v12
	var v19 int32
	_ = v19
	var v24 int32
	_ = v24
	var v25 int32
	_ = v25
	var v26 int32
	_ = v26
	var v28 int32
	_ = v28
	var v31 int32
	_ = v31
	var v36 int32
	_ = v36
	var v38 int32
	_ = v38
	var v40 int32
	_ = v40
	var v43 int32
	_ = v43
	var v48 int32
	_ = v48
	var v52 int32
	_ = v52
	var v53 int32
	_ = v53
	var v57 int32
	_ = v57
	var v60 int32
	_ = v60
	var v65 int32
	_ = v65
	var v72 int32
	_ = v72
	var v77 int32
	_ = v77
	var v78 int32
	_ = v78
	var v79 int32
	_ = v79
	var v81 int32
	_ = v81
	var v84 int32
	_ = v84
	var v89 int32
	_ = v89
	var v91 int32
	_ = v91
	var v93 int32
	_ = v93
	var v96 int32
	_ = v96
	var v101 int32
	_ = v101
	var v105 int32
	_ = v105
	var v106 int32
	_ = v106
	var v110 int32
	_ = v110
	var v111 int32
	_ = v111
	var v112 int32
	_ = v112
	var v115 int32
	_ = v115
	var v120 int32
	_ = v120
	var v127 int32
	_ = v127
	var v130 int32
	_ = v130
	var v131 int32
	_ = v131
	var v147 int32
	_ = v147
	var v150 int32
	_ = v150
	var v152 int32
	_ = v152
	var v155 int32
	_ = v155
	var v156 int32
	_ = v156
	var v160 int32
	_ = v160
	var v166 int32
	_ = v166
	var v170 int32
	_ = v170
	v7 = *(*int32)(unsafe.Add(mBase, _c_F_allcases[0]))
	switch v7 - int32(1) {
	case 0:
		if base.Ui32(l1) <= base.Ui32(int32(127)) {
			v24 = l1<<(uint(int32(2))%32) + int32(_a_F_allcases_0)
		} else {
			v19 = F_case_index(m, l1)
			mBase = m.M
			v24 = v19<<(uint(int32(2))%32) + int32(_a_F_allcases_1)
		}
		v25 = *(*int32)(unsafe.Add(mBase, uint32(v24)))
		if v25 != 0 {
			v26 = v25
		} else {
			v26 = l1
		}
		v57 = v26
	case 1:
		v28 = *(*int32)(unsafe.Add(mBase, _c_F_allcases[1]))
		if base.Ui32(int32(127)) < base.Ui32(l1) {
			v38 = F_towlower(m, l1)
			mBase = m.M
			v57 = v38
		} else {
			v31 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v28)+4)))
			if v31&int32(1) == int32(0) {
				v38 = F_towlower(m, l1)
				mBase = m.M
				v57 = v38
			} else {
				v36 = F_pg_tolower(m, l1)
				mBase = m.M
				v57 = v36
			}
		}
	case 2:
		v40 = *(*int32)(unsafe.Add(mBase, _c_F_allcases[1]))
		if base.Ui32(int32(127)) < base.Ui32(l1) {
			if base.Ui32(int32(255)) < base.Ui32(l1) {
				v53 = l1
			} else {
				v52 = F_tolower(m, l1)
				mBase = m.M
				v53 = v52
			}
			v57 = v53
		} else {
			v43 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v40)+4)))
			if v43&int32(1) == int32(0) {
				if base.Ui32(int32(255)) < base.Ui32(l1) {
					v53 = l1
				} else {
					v52 = F_tolower(m, l1)
					mBase = m.M
					v53 = v52
				}
				v57 = v53
			} else {
				v48 = F_pg_tolower(m, l1)
				mBase = m.M
				v57 = v48
			}
		}
	default:
		if base.Ui32(int32(127)) < base.Ui32(l1) {
			v53 = l1
			v57 = v53
		} else {
			v12 = F_pg_tolower(m, l1)
			mBase = m.M
			v57 = v12
		}
	}
	v60 = *(*int32)(unsafe.Add(mBase, _c_F_allcases[0]))
	switch v60 - int32(1) {
	case 0:
		if base.Ui32(l1) <= base.Ui32(int32(127)) {
			v77 = l1<<(uint(int32(2))%32) + int32(_a_F_allcases_2)
		} else {
			v72 = F_case_index(m, l1)
			mBase = m.M
			v77 = v72<<(uint(int32(2))%32) + int32(_a_F_allcases_3)
		}
		v78 = *(*int32)(unsafe.Add(mBase, uint32(v77)))
		if v78 != 0 {
			v79 = v78
		} else {
			v79 = l1
		}
		v110 = v79
	case 1:
		v81 = *(*int32)(unsafe.Add(mBase, _c_F_allcases[1]))
		if base.Ui32(int32(127)) < base.Ui32(l1) {
			v91 = F_towupper(m, l1)
			mBase = m.M
			v110 = v91
		} else {
			v84 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v81)+4)))
			if v84&int32(1) == int32(0) {
				v91 = F_towupper(m, l1)
				mBase = m.M
				v110 = v91
			} else {
				v89 = F_pg_toupper(m, l1)
				mBase = m.M
				v110 = v89
			}
		}
	case 2:
		v93 = *(*int32)(unsafe.Add(mBase, _c_F_allcases[1]))
		if base.Ui32(int32(127)) < base.Ui32(l1) {
			if base.Ui32(int32(255)) < base.Ui32(l1) {
				v106 = l1
			} else {
				v105 = F_toupper(m, l1)
				mBase = m.M
				v106 = v105
			}
			v110 = v106
		} else {
			v96 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v93)+4)))
			if v96&int32(1) == int32(0) {
				if base.Ui32(int32(255)) < base.Ui32(l1) {
					v106 = l1
				} else {
					v105 = F_toupper(m, l1)
					mBase = m.M
					v106 = v105
				}
				v110 = v106
			} else {
				v101 = F_pg_toupper(m, l1)
				mBase = m.M
				v110 = v101
			}
		}
	default:
		if base.Ui32(int32(127)) < base.Ui32(l1) {
			v106 = l1
			v110 = v106
		} else {
			v65 = F_pg_toupper(m, l1)
			mBase = m.M
			v110 = v65
		}
	}
	v111 = *(*int32)(unsafe.Add(mBase, uint32(l0)+120))
	if v111 != 0 {
		v112 = *(*int32)(unsafe.Add(mBase, uint32(v111)+4))
		if v112 < int32(2) {
			F_pfree(m, v111)
			mBase = m.M
			v127 = m.ExcPending
			if v127 != 0 {
				return int32(0)
			} else {
				v130 = F_palloc_extended(m, int32(36), int32(2))
				mBase = m.M
				v131 = m.ExcPending
				if v131 != 0 {
					return int32(0)
				} else {
					if v130 != 0 {
						*(*int64)(unsafe.Add(mBase, uint32(v130))) = int64(8589934592)
						*(*int32)(unsafe.Add(mBase, uint32(v130)+24)) = int32(-1)
						*(*int64)(unsafe.Add(mBase, uint32(v130)+12)) = int64(0)
						*(*int32)(unsafe.Add(mBase, uint32(v130)+20)) = v130 + int32(36)
						*(*int32)(unsafe.Add(mBase, uint32(v130)+8)) = v130 + int32(28)
						*(*int32)(unsafe.Add(mBase, uint32(l0)+120)) = v130
						v155 = v130
					} else {
						*(*int32)(unsafe.Add(mBase, uint32(l0)+24)) = int32(101)
						v147 = int32(0)
						*(*int32)(unsafe.Add(mBase, uint32(l0)+120)) = v147
						v150 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
						if v150 != 0 {
							v152 = v150
						} else {
							v152 = int32(12)
						}
						*(*int32)(unsafe.Add(mBase, uint32(l0)+12)) = v152
						v155 = v147
					}
					v156 = *(*int32)(unsafe.Add(mBase, uint32(v155)))
					*(*int32)(unsafe.Add(mBase, uint32(v155))) = v156 + int32(1)
					v160 = *(*int32)(unsafe.Add(mBase, uint32(v155)+8))
					*(*int32)(unsafe.Add(mBase, uint32(v160+v156<<(uint(int32(2))%32)))) = v57
					if v57 != v110 {
						v166 = *(*int32)(unsafe.Add(mBase, uint32(v155)))
						*(*int32)(unsafe.Add(mBase, uint32(v155))) = v166 + int32(1)
						v170 = *(*int32)(unsafe.Add(mBase, uint32(v155)+8))
						*(*int32)(unsafe.Add(mBase, uint32(v170+v166<<(uint(int32(2))%32)))) = v110
					} else {
					}
					return v155
				}
			}
		} else {
			v115 = *(*int32)(unsafe.Add(mBase, uint32(v111)+16))
			if v115 < int32(0) {
				F_pfree(m, v111)
				mBase = m.M
				v127 = m.ExcPending
				if v127 != 0 {
					return int32(0)
				} else {
					v130 = F_palloc_extended(m, int32(36), int32(2))
					mBase = m.M
					v131 = m.ExcPending
					if v131 != 0 {
						return int32(0)
					} else {
						if v130 != 0 {
							*(*int64)(unsafe.Add(mBase, uint32(v130))) = int64(8589934592)
							*(*int32)(unsafe.Add(mBase, uint32(v130)+24)) = int32(-1)
							*(*int64)(unsafe.Add(mBase, uint32(v130)+12)) = int64(0)
							*(*int32)(unsafe.Add(mBase, uint32(v130)+20)) = v130 + int32(36)
							*(*int32)(unsafe.Add(mBase, uint32(v130)+8)) = v130 + int32(28)
							*(*int32)(unsafe.Add(mBase, uint32(l0)+120)) = v130
							v155 = v130
						} else {
							*(*int32)(unsafe.Add(mBase, uint32(l0)+24)) = int32(101)
							v147 = int32(0)
							*(*int32)(unsafe.Add(mBase, uint32(l0)+120)) = v147
							v150 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
							if v150 != 0 {
								v152 = v150
							} else {
								v152 = int32(12)
							}
							*(*int32)(unsafe.Add(mBase, uint32(l0)+12)) = v152
							v155 = v147
						}
						v156 = *(*int32)(unsafe.Add(mBase, uint32(v155)))
						*(*int32)(unsafe.Add(mBase, uint32(v155))) = v156 + int32(1)
						v160 = *(*int32)(unsafe.Add(mBase, uint32(v155)+8))
						*(*int32)(unsafe.Add(mBase, uint32(v160+v156<<(uint(int32(2))%32)))) = v57
						if v57 != v110 {
							v166 = *(*int32)(unsafe.Add(mBase, uint32(v155)))
							*(*int32)(unsafe.Add(mBase, uint32(v155))) = v166 + int32(1)
							v170 = *(*int32)(unsafe.Add(mBase, uint32(v155)+8))
							*(*int32)(unsafe.Add(mBase, uint32(v170+v166<<(uint(int32(2))%32)))) = v110
						} else {
						}
						return v155
					}
				}
			} else {
				*(*int32)(unsafe.Add(mBase, uint32(v111)+24)) = int32(-1)
				v120 = int32(0)
				*(*int32)(unsafe.Add(mBase, uint32(v111)+12)) = v120
				*(*int32)(unsafe.Add(mBase, uint32(v111))) = v120
				v155 = v111
				v156 = *(*int32)(unsafe.Add(mBase, uint32(v155)))
				*(*int32)(unsafe.Add(mBase, uint32(v155))) = v156 + int32(1)
				v160 = *(*int32)(unsafe.Add(mBase, uint32(v155)+8))
				*(*int32)(unsafe.Add(mBase, uint32(v160+v156<<(uint(int32(2))%32)))) = v57
				if v57 != v110 {
					v166 = *(*int32)(unsafe.Add(mBase, uint32(v155)))
					*(*int32)(unsafe.Add(mBase, uint32(v155))) = v166 + int32(1)
					v170 = *(*int32)(unsafe.Add(mBase, uint32(v155)+8))
					*(*int32)(unsafe.Add(mBase, uint32(v170+v166<<(uint(int32(2))%32)))) = v110
				} else {
				}
				return v155
			}
		}
	} else {
		v130 = F_palloc_extended(m, int32(36), int32(2))
		mBase = m.M
		v131 = m.ExcPending
		if v131 != 0 {
			return int32(0)
		} else {
			if v130 != 0 {
				*(*int64)(unsafe.Add(mBase, uint32(v130))) = int64(8589934592)
				*(*int32)(unsafe.Add(mBase, uint32(v130)+24)) = int32(-1)
				*(*int64)(unsafe.Add(mBase, uint32(v130)+12)) = int64(0)
				*(*int32)(unsafe.Add(mBase, uint32(v130)+20)) = v130 + int32(36)
				*(*int32)(unsafe.Add(mBase, uint32(v130)+8)) = v130 + int32(28)
				*(*int32)(unsafe.Add(mBase, uint32(l0)+120)) = v130
				v155 = v130
			} else {
				*(*int32)(unsafe.Add(mBase, uint32(l0)+24)) = int32(101)
				v147 = int32(0)
				*(*int32)(unsafe.Add(mBase, uint32(l0)+120)) = v147
				v150 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
				if v150 != 0 {
					v152 = v150
				} else {
					v152 = int32(12)
				}
				*(*int32)(unsafe.Add(mBase, uint32(l0)+12)) = v152
				v155 = v147
			}
			v156 = *(*int32)(unsafe.Add(mBase, uint32(v155)))
			*(*int32)(unsafe.Add(mBase, uint32(v155))) = v156 + int32(1)
			v160 = *(*int32)(unsafe.Add(mBase, uint32(v155)+8))
			*(*int32)(unsafe.Add(mBase, uint32(v160+v156<<(uint(int32(2))%32)))) = v57
			if v57 != v110 {
				v166 = *(*int32)(unsafe.Add(mBase, uint32(v155)))
				*(*int32)(unsafe.Add(mBase, uint32(v155))) = v166 + int32(1)
				v170 = *(*int32)(unsafe.Add(mBase, uint32(v155)+8))
				*(*int32)(unsafe.Add(mBase, uint32(v170+v166<<(uint(int32(2))%32)))) = v110
			} else {
			}
			return v155
		}
	}
}
func F_amcheck_lock_relation_and_check(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32) {
	mBase := m.M
	_ = mBase
	var v11 int32
	_ = v11
	var v13 int32
	_ = v13
	var v16 int32
	_ = v16
	var v17 int32
	_ = v17
	var v24 int32
	_ = v24
	var v25 int32
	_ = v25
	var v26 int32
	_ = v26
	var v27 int32
	_ = v27
	var v33 int32
	_ = v33
	var v36 int32
	_ = v36
	var v38 int32
	_ = v38
	var v39 int32
	_ = v39
	var v40 int32
	_ = v40
	var v48 int32
	_ = v48
	var v50 int32
	_ = v50
	var v52 int32
	_ = v52
	var v54 int32
	_ = v54
	var v55 int32
	_ = v55
	var v57 int32
	_ = v57
	var v58 int32
	_ = v58
	var v60 int32
	_ = v60
	var v61 int32
	_ = v61
	var v65 int32
	_ = v65
	var v68 int32
	_ = v68
	var v69 int32
	_ = v69
	var v70 int32
	_ = v70
	var v76 int32
	_ = v76
	var v78 int32
	_ = v78
	var v82 int32
	_ = v82
	var v88 int32
	_ = v88
	var v91 int32
	_ = v91
	var v92 int32
	_ = v92
	var v98 int32
	_ = v98
	var v103 int32
	_ = v103
	v11 = m.G0
	v13 = v11 - int32(16)
	m.G0 = v13
	v16 = F_IndexGetRelation(m, l0, int32(1))
	mBase = m.M
	v17 = m.ExcPending
	if v17 != 0 {
		return
	} else {
		if v16 == int32(0) {
			*(*int32)(unsafe.Add(mBase, uint32(v13)+8)) = int32(-1)
			*(*int32)(unsafe.Add(mBase, uint32(v13)+12)) = int32(0)
			v24 = F_index_open(m, l0, l3)
			mBase = m.M
			v25 = m.ExcPending
			if v25 != 0 {
				return
			} else {
				v82 = v24
				F_errstart_cold(m, int32(21), int32(0))
				mBase = m.M
				v88 = m.ExcPending
				if v88 != 0 {
					return
				} else {
					F_errcode(m, int32(16908420))
					mBase = m.M
					v91 = m.ExcPending
					if v91 != 0 {
						return
					} else {
						v92 = *(*int32)(unsafe.Add(mBase, uint32(v82)+48))
						*(*int32)(unsafe.Add(mBase, uint32(v13))) = v92 + int32(4)
						F_errmsg(m, int32(_a_F_amcheck_lock_relation_and_check_0), v13)
						mBase = m.M
						v98 = m.ExcPending
						if v98 != 0 {
							return
						} else {
							F_errfinish(m, int32(_a_F_amcheck_lock_relation_and_check_1), int32(128), int32(_a_F_amcheck_lock_relation_and_check_2))
							mBase = m.M
							v103 = m.ExcPending
							if v103 != 0 {
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
			v26 = F_table_open(m, v16, l3)
			mBase = m.M
			v27 = m.ExcPending
			if v27 != 0 {
				return
			} else {
				v33 = *(*int32)(unsafe.Add(mBase, _c_F_amcheck_lock_relation_and_check[0]))
				*(*int32)(unsafe.Add(mBase, uint32(v13+int32(12)))) = v33
				v36 = *(*int32)(unsafe.Add(mBase, _c_F_amcheck_lock_relation_and_check[1]))
				*(*int32)(unsafe.Add(mBase, uint32(v13+int32(8)))) = v36
				v38 = *(*int32)(unsafe.Add(mBase, uint32(v26)+48))
				v39 = *(*int32)(unsafe.Add(mBase, uint32(v38)+80))
				v40 = *(*int32)(unsafe.Add(mBase, uint32(v13)+8))
				*(*int32)(unsafe.Add(mBase, _c_F_amcheck_lock_relation_and_check[1])) = v40 | int32(2)
				*(*int32)(unsafe.Add(mBase, _c_F_amcheck_lock_relation_and_check[0])) = v39
				v48 = int32(_a_F_amcheck_lock_relation_and_check_3)
				v50 = *(*int32)(unsafe.Add(mBase, _c_F_amcheck_lock_relation_and_check[2]))
				v52 = v50 + int32(1)
				*(*int32)(unsafe.Add(mBase, _c_F_amcheck_lock_relation_and_check[2])) = v52
				v54 = F_index_open(m, l0, l3)
				mBase = m.M
				v55 = m.ExcPending
				if v55 != 0 {
					return
				} else {
					v57 = F_IndexGetRelation(m, l0, int32(0))
					mBase = m.M
					v58 = m.ExcPending
					if v58 != 0 {
						return
					} else {
						if v57 != v16 {
							v82 = v54
							F_errstart_cold(m, int32(21), int32(0))
							mBase = m.M
							v88 = m.ExcPending
							if v88 != 0 {
								return
							} else {
								F_errcode(m, int32(16908420))
								mBase = m.M
								v91 = m.ExcPending
								if v91 != 0 {
									return
								} else {
									v92 = *(*int32)(unsafe.Add(mBase, uint32(v82)+48))
									*(*int32)(unsafe.Add(mBase, uint32(v13))) = v92 + int32(4)
									F_errmsg(m, int32(_a_F_amcheck_lock_relation_and_check_0), v13)
									mBase = m.M
									v98 = m.ExcPending
									if v98 != 0 {
										return
									} else {
										F_errfinish(m, int32(_a_F_amcheck_lock_relation_and_check_1), int32(128), int32(_a_F_amcheck_lock_relation_and_check_2))
										mBase = m.M
										v103 = m.ExcPending
										if v103 != 0 {
											return
										} else {
											base.Wasm_trap_unreachable()
											for {
											}
										}
									}
								}
							}
						} else {
							v60 = F_index_checkable(m, v54, l1)
							mBase = m.M
							v61 = m.ExcPending
							if v61 != 0 {
								return
							} else {
								if v60 != 0 {
									m.T0[l2].(func(*base.Module, int32, int32, int32, int32))(m, v54, v26, l4, base.B2i32(l3 == int32(5)))
									mBase = m.M
									v65 = m.ExcPending
									if v65 != 0 {
										return
									} else {
										F_AtEOXact_GUC(m, int32(0), v52)
										mBase = m.M
										v68 = m.ExcPending
										if v68 != 0 {
											return
										} else {
											v69 = *(*int32)(unsafe.Add(mBase, uint32(v13)+12))
											v70 = *(*int32)(unsafe.Add(mBase, uint32(v13)+8))
											*(*int32)(unsafe.Add(mBase, _c_F_amcheck_lock_relation_and_check[1])) = v70
											*(*int32)(unsafe.Add(mBase, _c_F_amcheck_lock_relation_and_check[0])) = v69
											F_relation_close(m, v54, l3)
											mBase = m.M
											v76 = m.ExcPending
											if v76 != 0 {
												return
											} else {
												F_relation_close(m, v26, l3)
												mBase = m.M
												v78 = m.ExcPending
												if v78 != 0 {
													return
												} else {
													m.G0 = v13 + int32(16)
													return
												}
											}
										}
									}
								} else {
									F_AtEOXact_GUC(m, int32(0), v52)
									mBase = m.M
									v68 = m.ExcPending
									if v68 != 0 {
										return
									} else {
										v69 = *(*int32)(unsafe.Add(mBase, uint32(v13)+12))
										v70 = *(*int32)(unsafe.Add(mBase, uint32(v13)+8))
										*(*int32)(unsafe.Add(mBase, _c_F_amcheck_lock_relation_and_check[1])) = v70
										*(*int32)(unsafe.Add(mBase, _c_F_amcheck_lock_relation_and_check[0])) = v69
										F_relation_close(m, v54, l3)
										mBase = m.M
										v76 = m.ExcPending
										if v76 != 0 {
											return
										} else {
											F_relation_close(m, v26, l3)
											mBase = m.M
											v78 = m.ExcPending
											if v78 != 0 {
												return
											} else {
												m.G0 = v13 + int32(16)
												return
											}
										}
									}
								}
							}
						}
					}
				}
			}
		}
	}
}
func F_anyarray_recv(m *base.Module, l0 int32) int32 {
	var v7 int32
	_ = v7
	var v10 int32
	_ = v10
	v7 = Fn13874(m, l0, int32(_a_F_anyarray_recv_0), int32(155), int32(_a_F_anyarray_recv_1), int32(_a_F_anyarray_recv_2), int32(_a_F_anyarray_recv_3))
	v10 = m.ExcPending
	if v10 != 0 {
		return int32(0)
	} else {
		return v7
	}
}
func F_anybit_typmodin(m *base.Module, l0 int32, l1 int32) int32 {
	mBase := m.M
	_ = mBase
	var v4 int32
	_ = v4
	var v6 int32
	_ = v6
	var v10 int32
	_ = v10
	var v13 int32
	_ = v13
	var v14 int32
	_ = v14
	var v17 int32
	_ = v17
	var v29 int32
	_ = v29
	var v32 int32
	_ = v32
	var v36 int32
	_ = v36
	var v41 int32
	_ = v41
	var v45 int32
	_ = v45
	var v48 int32
	_ = v48
	var v52 int32
	_ = v52
	var v57 int32
	_ = v57
	var v61 int32
	_ = v61
	var v64 int32
	_ = v64
	var v72 int32
	_ = v72
	var v77 int32
	_ = v77
	v4 = m.G0
	v6 = v4 - int32(32)
	m.G0 = v6
	v10 = F_ArrayGetIntegerTypmods(m, l0, v6+int32(28))
	mBase = m.M
	v13 = m.ExcPending
	if v13 != 0 {
		return int32(0)
	} else {
		v14 = *(*int32)(unsafe.Add(mBase, uint32(v6)+28))
		if v14 == int32(1) {
			v17 = *(*int32)(unsafe.Add(mBase, uint32(v10)))
			if v17 <= int32(0) {
				F_errstart_cold(m, int32(21), int32(0))
				mBase = m.M
				v45 = m.ExcPending
				if v45 != 0 {
					return int32(0)
				} else {
					F_errcode(m, int32(50856066))
					mBase = m.M
					v48 = m.ExcPending
					if v48 != 0 {
						return int32(0)
					} else {
						*(*int32)(unsafe.Add(mBase, uint32(v6))) = l1
						F_errmsg(m, int32(_a_F_anybit_typmodin_0), v6)
						mBase = m.M
						v52 = m.ExcPending
						if v52 != 0 {
							return int32(0)
						} else {
							F_errfinish(m, int32(_a_F_anybit_typmodin_1), int32(111), int32(_a_F_anybit_typmodin_2))
							mBase = m.M
							v57 = m.ExcPending
							if v57 != 0 {
								return int32(0)
							} else {
								base.Wasm_trap_unreachable()
								for {
								}
							}
						}
					}
				}
			} else {
				if base.Ui32(int32(83886081)) <= base.Ui32(v17) {
					F_errstart_cold(m, int32(21), int32(0))
					mBase = m.M
					v61 = m.ExcPending
					if v61 != 0 {
						return int32(0)
					} else {
						F_errcode(m, int32(50856066))
						mBase = m.M
						v64 = m.ExcPending
						if v64 != 0 {
							return int32(0)
						} else {
							*(*int32)(unsafe.Add(mBase, uint32(v6)+20)) = int32(83886080)
							*(*int32)(unsafe.Add(mBase, uint32(v6)+16)) = l1
							F_errmsg(m, int32(_a_F_anybit_typmodin_3), v6+int32(16))
							mBase = m.M
							v72 = m.ExcPending
							if v72 != 0 {
								return int32(0)
							} else {
								F_errfinish(m, int32(_a_F_anybit_typmodin_1), int32(116), int32(_a_F_anybit_typmodin_2))
								mBase = m.M
								v77 = m.ExcPending
								if v77 != 0 {
									return int32(0)
								} else {
									base.Wasm_trap_unreachable()
									for {
									}
								}
							}
						}
					}
				} else {
					m.G0 = v6 + int32(32)
					return v17
				}
			}
		} else {
			F_errstart_cold(m, int32(21), int32(0))
			mBase = m.M
			v29 = m.ExcPending
			if v29 != 0 {
				return int32(0)
			} else {
				F_errcode(m, int32(50856066))
				mBase = m.M
				v32 = m.ExcPending
				if v32 != 0 {
					return int32(0)
				} else {
					F_errmsg(m, int32(_a_F_anybit_typmodin_4), int32(0))
					mBase = m.M
					v36 = m.ExcPending
					if v36 != 0 {
						return int32(0)
					} else {
						F_errfinish(m, int32(_a_F_anybit_typmodin_1), int32(105), int32(_a_F_anybit_typmodin_2))
						mBase = m.M
						v41 = m.ExcPending
						if v41 != 0 {
							return int32(0)
						} else {
							base.Wasm_trap_unreachable()
							for {
							}
						}
					}
				}
			}
		}
	}
}
func F_anycompatiblemultirange_in(m *base.Module, l0 int32) int32 {
	var v7 int32
	_ = v7
	var v10 int32
	_ = v10
	v7 = Fn13874(m, l0, int32(_a_F_anycompatiblemultirange_in_0), int32(246), int32(_a_F_anycompatiblemultirange_in_1), int32(_a_F_anycompatiblemultirange_in_2), int32(_a_F_anycompatiblemultirange_in_3))
	v10 = m.ExcPending
	if v10 != 0 {
		return int32(0)
	} else {
		return v7
	}
}
func F_anyenum_out(m *base.Module, l0 int32) int32 {
	var v2 int32
	_ = v2
	var v5 int32
	_ = v5
	v2 = F_enum_out(m, l0)
	v5 = m.ExcPending
	if v5 != 0 {
		return int32(0)
	} else {
		return v2
	}
}
func F_anyrange_out(m *base.Module, l0 int32) int32 {
	var v2 int32
	_ = v2
	var v5 int32
	_ = v5
	v2 = F_range_out(m, l0)
	v5 = m.ExcPending
	if v5 != 0 {
		return int32(0)
	} else {
		return v2
	}
}
func F_appendContextKeyword(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32) {
	mBase := m.M
	_ = mBase
	var v8 int32
	_ = v8
	var v9 int32
	_ = v9
	var v12 int32
	_ = v12
	var v15 int32
	_ = v15
	var v20 int32
	_ = v20
	var v25 int32
	_ = v25
	var v29 int32
	_ = v29
	var v33 int32
	_ = v33
	var v36 int32
	_ = v36
	var v38 int32
	_ = v38
	var v50 int32
	_ = v50
	var v51 int32
	_ = v51
	var v54 int32
	_ = v54
	var v57 int32
	_ = v57
	var v58 int32
	_ = v58
	var v65 int32
	_ = v65
	var v66 int32
	_ = v66
	var v69 int32
	_ = v69
	var v71 int32
	_ = v71
	var v72 int32
	_ = v72
	var v73 int32
	_ = v73
	var v74 int32
	_ = v74
	var v77 int32
	_ = v77
	var v80 int32
	_ = v80
	v8 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v9 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+20)))
	if v9&int32(2) != 0 {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	v12 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+28)) = v12 + l2
	v15 = *(*int32)(unsafe.Add(mBase, uint32(v8)+4))
	if v15 <= int32(0) {
		goto L4
	} else {
		goto L5
	}
L2:
	;
	goto L3
L3:
	;
	F_appendStringInfoString(m, v8, l1)
	mBase = m.M
	v80 = m.ExcPending
	if v80 != 0 {
		goto L10
	} else {
		goto L24
	}
L4:
	;
	F_appendStringInfoChar(m, v8, int32(10))
	mBase = m.M
	v50 = m.ExcPending
	if v50 != 0 {
		goto L10
	} else {
		goto L11
	}
L5:
	;
	v20 = v15
	goto L6
L6:
	;
	v25 = *(*int32)(unsafe.Add(mBase, uint32(v8)))
	v29 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v25+v20-int32(1)))))
	if v29 != int32(32) {
		goto L4
	} else {
		goto L8
	}
L7:
	;
	goto L4
L8:
	;
	v33 = v20 - int32(1)
	*(*int32)(unsafe.Add(mBase, uint32(v8)+4)) = v33
	v36 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v33+v25))) = uint8(v36)
	v38 = *(*int32)(unsafe.Add(mBase, uint32(v8)+4))
	if v36 < v38 {
		v20 = v38
		goto L6
	} else {
		goto L9
	}
L9:
	;
	goto L7
L10:
	;
	return
L11:
	;
	v51 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
	if v51 <= int32(39) {
		goto L13
	} else {
		goto L14
	}
L12:
	;
	F_appendStringInfoSpaces(m, v8, v66+l4)
	mBase = m.M
	v69 = m.ExcPending
	if v69 != 0 {
		goto L10
	} else {
		goto L19
	}
L13:
	;
	v54 = int32(0)
	if v54 < v51 {
		goto L16
	} else {
		goto L17
	}
L14:
	;
	goto L15
L15:
	;
	v58 = int32(40)
	v65 = base.I32_rem_u_s(int32(base.Ui32(v51-v58)>>(uint(int32(2))%32))+v58, v58)
	v66 = v65
	goto L12
L16:
	;
	v57 = v51
	goto L18
L17:
	;
	v57 = v54
	goto L18
L18:
	;
	v66 = v57
	goto L12
L19:
	;
	F_appendStringInfoString(m, v8, l1)
	mBase = m.M
	v71 = m.ExcPending
	if v71 != 0 {
		goto L10
	} else {
		goto L20
	}
L20:
	;
	v72 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
	v73 = v72 + l3
	v74 = int32(0)
	if v74 < v73 {
		goto L21
	} else {
		goto L22
	}
L21:
	;
	v77 = v73
	goto L23
L22:
	;
	v77 = v74
	goto L23
L23:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+28)) = v77
	return
L24:
	;
	return
}
func F_appendElement(m *base.Module, l0 int32, l1 int32) {
	mBase := m.M
	_ = mBase
	var v6 int32
	_ = v6
	var v8 int32
	_ = v8
	var v10 int32
	_ = v10
	var v13 int32
	_ = v13
	var v15 int32
	_ = v15
	var v19 int32
	_ = v19
	var v22 int32
	_ = v22
	var v23 int32
	_ = v23
	var v25 int32
	_ = v25
	var v26 int32
	_ = v26
	var v27 int32
	_ = v27
	var v33 int32
	_ = v33
	var v34 int32
	_ = v34
	var v36 int64
	_ = v36
	var v38 int64
	_ = v38
	var v46 int32
	_ = v46
	var v49 int32
	_ = v49
	var v54 int32
	_ = v54
	var v59 int32
	_ = v59
	v6 = m.G0
	v8 = v6 - int32(16)
	m.G0 = v8
	v10 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	if base.Ui32(v10) < base.Ui32(int32(53687091)) {
		v13 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
		if base.Ui32(v10) < base.Ui32(v13) {
			v15 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
			v26 = v15
			v27 = v10
			*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v27 + int32(1)
			v33 = v26 + v27*int32(20)
			v34 = *(*int32)(unsafe.Add(mBase, uint32(l1)+16))
			*(*int32)(unsafe.Add(mBase, uint32(v33)+16)) = v34
			v36 = *(*int64)(unsafe.Add(mBase, uint32(l1)+8))
			*(*int64)(unsafe.Add(mBase, uint32(v33)+8)) = v36
			v38 = *(*int64)(unsafe.Add(mBase, uint32(l1)))
			*(*int64)(unsafe.Add(mBase, uint32(v33))) = v38
			m.G0 = v8 + int32(16)
			return
		} else {
			*(*int32)(unsafe.Add(mBase, uint32(l0)+20)) = v13 << (uint(int32(1)) % 32)
			v19 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
			v22 = F_repalloc(m, v19, v13*int32(40))
			mBase = m.M
			v23 = m.ExcPending
			if v23 != 0 {
				return
			} else {
				*(*int32)(unsafe.Add(mBase, uint32(l0)+8)) = v22
				v25 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
				v26 = v22
				v27 = v25
				*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v27 + int32(1)
				v33 = v26 + v27*int32(20)
				v34 = *(*int32)(unsafe.Add(mBase, uint32(l1)+16))
				*(*int32)(unsafe.Add(mBase, uint32(v33)+16)) = v34
				v36 = *(*int64)(unsafe.Add(mBase, uint32(l1)+8))
				*(*int64)(unsafe.Add(mBase, uint32(v33)+8)) = v36
				v38 = *(*int64)(unsafe.Add(mBase, uint32(l1)))
				*(*int64)(unsafe.Add(mBase, uint32(v33))) = v38
				m.G0 = v8 + int32(16)
				return
			}
		}
	} else {
		F_errstart_cold(m, int32(21), int32(0))
		mBase = m.M
		v46 = m.ExcPending
		if v46 != 0 {
			return
		} else {
			F_errcode(m, int32(261))
			mBase = m.M
			v49 = m.ExcPending
			if v49 != 0 {
				return
			} else {
				*(*int32)(unsafe.Add(mBase, uint32(v8))) = int32(53687091)
				F_errmsg(m, int32(_a_F_appendElement_0), v8)
				mBase = m.M
				v54 = m.ExcPending
				if v54 != 0 {
					return
				} else {
					F_errfinish(m, int32(_a_F_appendElement_1), int32(802), int32(_a_F_appendElement_2))
					mBase = m.M
					v59 = m.ExcPending
					if v59 != 0 {
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
}
func F_append_num_word(m *base.Module, l0 int32, l1 int64) {
	mBase := m.M
	_ = mBase
	var v8 int32
	_ = v8
	var v10 int32
	_ = v10
	var v12 int32
	_ = v12
	var v16 int32
	_ = v16
	var v21 int32
	_ = v21
	var v23 int32
	_ = v23
	var v26 int32
	_ = v26
	var v28 int32
	_ = v28
	var v34 int32
	_ = v34
	var v37 int32
	_ = v37
	var v41 int32
	_ = v41
	var v44 int32
	_ = v44
	var v45 int32
	_ = v45
	var v47 int32
	_ = v47
	var v58 int32
	_ = v58
	var v59 int32
	_ = v59
	var v61 int32
	_ = v61
	var v64 int32
	_ = v64
	var v67 int32
	_ = v67
	var v73 int32
	_ = v73
	var v76 int32
	_ = v76
	var v84 int32
	_ = v84
	var v90 int32
	_ = v90
	var v92 int32
	_ = v92
	var v94 int32
	_ = v94
	var v95 int32
	_ = v95
	var v96 int32
	_ = v96
	var v98 int32
	_ = v98
	var v107 int32
	_ = v107
	var v113 int32
	_ = v113
	var v117 int32
	_ = v117
	var v128 int32
	_ = v128
	var v131 int32
	_ = v131
	var v133 int32
	_ = v133
	var v140 int32
	_ = v140
	var v142 int32
	_ = v142
	var v143 int32
	_ = v143
	var v145 int32
	_ = v145
	var v146 int32
	_ = v146
	var v147 int32
	_ = v147
	var v149 int32
	_ = v149
	var v158 int32
	_ = v158
	var v164 int32
	_ = v164
	v8 = m.G0
	v10 = v8 - int32(80)
	m.G0 = v10
	v12 = base.I32_wrap_i64(l1)
	v16 = base.I32_div_u_s(v12&int32(_a_F_append_num_word_0), int32(100))
	if base.Ui64(l1) <= base.Ui64(int64(20)) {
		v21 = *(*int32)(unsafe.Add(mBase, uint32(v12<<(uint(int32(2))%32))+uint32(_c_F_append_num_word[0])))
		F_appendStringInfoString(m, l0, v21)
		mBase = m.M
		v23 = m.ExcPending
		if v23 != 0 {
			return
		} else {
			m.G0 = v10 + int32(80)
			return
		}
	} else {
		v26 = v12 - v16*int32(100)
		v28 = v26 & int32(_a_F_append_num_word_0)
		if v28 == int32(0) {
			v34 = base.I32_div_u_s(v12&int32(_a_F_append_num_word_0), int32(100))
			v37 = *(*int32)(unsafe.Add(mBase, uint32(v34<<(uint(int32(2))%32))+uint32(_c_F_append_num_word[0])))
			*(*int32)(unsafe.Add(mBase, uint32(v10))) = v37
			F_appendStringInfo(m, l0, int32(_a_F_append_num_word_1), v10)
			mBase = m.M
			v41 = m.ExcPending
			if v41 != 0 {
				return
			} else {
				m.G0 = v10 + int32(80)
				return
			}
		} else {
			if base.Ui64(int64(100)) <= base.Ui64(l1) {
				v44 = int32(_a_F_append_num_word_0)
				v45 = v12 & v44
				v47 = base.I32_rem_u_s(v45, int32(10))
				if v47|base.B2i32(base.Ui32(v26&v44) < base.Ui32(int32(11))) == int32(0) {
					v58 = base.I32_div_u_s(v26&int32(255), int32(10))
					v59 = int32(2)
					v61 = *(*int32)(unsafe.Add(mBase, uint32(v58<<(uint(v59)%32))+uint32(_c_F_append_num_word[1])))
					*(*int32)(unsafe.Add(mBase, uint32(v10)+20)) = v61
					v64 = base.I32_div_u_s(v45, int32(100))
					v67 = *(*int32)(unsafe.Add(mBase, uint32(v64<<(uint(v59)%32))+uint32(_c_F_append_num_word[0])))
					*(*int32)(unsafe.Add(mBase, uint32(v10)+16)) = v67
					F_appendStringInfo(m, l0, int32(_a_F_append_num_word_2), v10+int32(16))
					mBase = m.M
					v73 = m.ExcPending
					if v73 != 0 {
						return
					} else {
						m.G0 = v10 + int32(80)
						return
					}
				} else {
					v76 = *(*int32)(unsafe.Add(mBase, uint32(v16<<(uint(int32(2))%32))+uint32(_c_F_append_num_word[0])))
					if base.Ui32(v26&int32(_a_F_append_num_word_0)) <= base.Ui32(int32(19)) {
						*(*int32)(unsafe.Add(mBase, uint32(v10)+32)) = v76
						v84 = *(*int32)(unsafe.Add(mBase, uint32(v28<<(uint(int32(2))%32))+uint32(_c_F_append_num_word[0])))
						*(*int32)(unsafe.Add(mBase, uint32(v10)+36)) = v84
						F_appendStringInfo(m, l0, int32(_a_F_append_num_word_3), v10+int32(32))
						mBase = m.M
						v90 = m.ExcPending
						if v90 != 0 {
							return
						} else {
							m.G0 = v10 + int32(80)
							return
						}
					} else {
						*(*int32)(unsafe.Add(mBase, uint32(v10)+48)) = v76
						v92 = int32(255)
						v94 = int32(10)
						v95 = base.I32_div_u_s(v26&v92, v94)
						v96 = int32(2)
						v98 = *(*int32)(unsafe.Add(mBase, uint32(v95<<(uint(v96)%32))+uint32(_c_F_append_num_word[1])))
						*(*int32)(unsafe.Add(mBase, uint32(v10)+52)) = v98
						v107 = *(*int32)(unsafe.Add(mBase, uint32((v26-v95*v94)&v92<<(uint(v96)%32))+uint32(_c_F_append_num_word[0])))
						*(*int32)(unsafe.Add(mBase, uint32(v10)+56)) = v107
						F_appendStringInfo(m, l0, int32(_a_F_append_num_word_4), v10+int32(48))
						mBase = m.M
						v113 = m.ExcPending
						if v113 != 0 {
							return
						} else {
							m.G0 = v10 + int32(80)
							return
						}
					}
				}
			} else {
				v117 = base.I32_rem_u_s(v12&int32(255), int32(10))
				if v117|base.B2i32(base.Ui32(v26&int32(_a_F_append_num_word_0)) < base.Ui32(int32(11))) == int32(0) {
					v128 = base.I32_div_u_s(v26&int32(255), int32(10))
					v131 = *(*int32)(unsafe.Add(mBase, uint32(v128<<(uint(int32(2))%32))+uint32(_c_F_append_num_word[1])))
					F_appendStringInfoString(m, l0, v131)
					mBase = m.M
					v133 = m.ExcPending
					if v133 != 0 {
						return
					} else {
						m.G0 = v10 + int32(80)
						return
					}
				} else {
					if base.Ui32(v26&int32(_a_F_append_num_word_0)) <= base.Ui32(int32(19)) {
						v140 = *(*int32)(unsafe.Add(mBase, uint32(v28<<(uint(int32(2))%32))+uint32(_c_F_append_num_word[0])))
						F_appendStringInfoString(m, l0, v140)
						mBase = m.M
						v142 = m.ExcPending
						if v142 != 0 {
							return
						} else {
							m.G0 = v10 + int32(80)
							return
						}
					} else {
						v143 = int32(255)
						v145 = int32(10)
						v146 = base.I32_div_u_s(v26&v143, v145)
						v147 = int32(2)
						v149 = *(*int32)(unsafe.Add(mBase, uint32(v146<<(uint(v147)%32))+uint32(_c_F_append_num_word[1])))
						*(*int32)(unsafe.Add(mBase, uint32(v10)+64)) = v149
						v158 = *(*int32)(unsafe.Add(mBase, uint32((v26-v146*v145)&v143<<(uint(v147)%32))+uint32(_c_F_append_num_word[0])))
						*(*int32)(unsafe.Add(mBase, uint32(v10)+68)) = v158
						F_appendStringInfo(m, l0, int32(_a_F_append_num_word_5), v10-int32(-64))
						mBase = m.M
						v164 = m.ExcPending
						if v164 != 0 {
							return
						} else {
							m.G0 = v10 + int32(80)
							return
						}
					}
				}
			}
		}
	}
}
func F_applyLockingClause(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) {
	mBase := m.M
	_ = mBase
	var v9 int32
	_ = v9
	var v10 int32
	_ = v10
	var v13 int32
	_ = v13
	var v16 int32
	_ = v16
	var v17 int32
	_ = v17
	var v19 int32
	_ = v19
	var v27 int32
	_ = v27
	var v28 int32
	_ = v28
	var v31 int32
	_ = v31
	var v42 int32
	_ = v42
	var v44 int32
	_ = v44
	var v46 int32
	_ = v46
	var v48 int32
	_ = v48
	var v50 int32
	_ = v50
	var v52 int32
	_ = v52
	var v54 int32
	_ = v54
	var v57 int32
	_ = v57
	var v58 int32
	_ = v58
	var v59 int32
	_ = v59
	var v66 int32
	_ = v66
	var v67 int32
	_ = v67
	var v68 int32
	_ = v68
	v9 = *(*int32)(unsafe.Add(mBase, uint32(l0)+140))
	if v9 != 0 {
		goto L3
	} else {
		goto L4
	}
L1:
	;
	if v42 != 0 {
		goto L14
	} else {
		goto L15
	}
L2:
	;
	goto L1
L3:
	;
	v10 = *(*int32)(unsafe.Add(mBase, uint32(v9)+4))
	if v10 <= int32(0) {
		v42 = int32(0)
		goto L2
	} else {
		goto L6
	}
L4:
	;
	goto L5
L5:
	;
	v42 = int32(0)
	goto L2
L6:
	;
	v13 = int32(0)
	if v13 < v10 {
		goto L7
	} else {
		goto L8
	}
L7:
	;
	v16 = v10
	goto L9
L8:
	;
	v16 = v13
	goto L9
L9:
	;
	v17 = *(*int32)(unsafe.Add(mBase, uint32(v9)+12))
	v19 = int32(0)
	goto L10
L10:
	;
	v27 = *(*int32)(unsafe.Add(mBase, uint32(v17+v19<<(uint(int32(2))%32))))
	v28 = *(*int32)(unsafe.Add(mBase, uint32(v27)+4))
	if v28 == l1 {
		v42 = v27
		goto L2
	} else {
		goto L12
	}
L11:
	;
	goto L5
L12:
	;
	v31 = v19 + int32(1)
	if v31 != v16 {
		v19 = v31
		goto L10
	} else {
		goto L13
	}
L13:
	;
	goto L11
L14:
	;
	v44 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v42)+16)))
	v46 = v44 & int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(v42)+16)) = uint8(v46)
	v48 = *(*int32)(unsafe.Add(mBase, uint32(v42)+8))
	if base.Ui32(l2) < base.Ui32(v48) {
		goto L17
	} else {
		goto L18
	}
L15:
	;
	goto L16
L16:
	;
	v57 = F_palloc0(m, int32(20))
	mBase = m.M
	v58 = m.ExcPending
	if v58 != 0 {
		goto L23
	} else {
		goto L24
	}
L17:
	;
	v50 = v48
	goto L19
L18:
	;
	v50 = l2
	goto L19
L19:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v42)+8)) = v50
	v52 = *(*int32)(unsafe.Add(mBase, uint32(v42)+12))
	if base.Ui32(l3) < base.Ui32(v52) {
		goto L20
	} else {
		goto L21
	}
L20:
	;
	v54 = v52
	goto L22
L21:
	;
	v54 = l3
	goto L22
L22:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v42)+12)) = v54
	return
L23:
	;
	return
L24:
	;
	v59 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(v57)+16)) = uint8(v59)
	*(*int32)(unsafe.Add(mBase, uint32(v57)+12)) = l3
	*(*int32)(unsafe.Add(mBase, uint32(v57)+8)) = l2
	*(*int32)(unsafe.Add(mBase, uint32(v57)+4)) = l1
	*(*int32)(unsafe.Add(mBase, uint32(v57))) = int32(109)
	v66 = *(*int32)(unsafe.Add(mBase, uint32(l0)+140))
	v67 = F_lappend(m, v66, v57)
	mBase = m.M
	v68 = m.ExcPending
	if v68 != 0 {
		goto L23
	} else {
		goto L25
	}
L25:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+140)) = v67
	return
}
func F_apply_typmod(m *base.Module, l0 int32, l1 int32, l2 int32) int32 {
	mBase := m.M
	_ = mBase
	var v12 int32
	_ = v12
	var v14 int32
	_ = v14
	var v16 int32
	_ = v16
	var v19 int32
	_ = v19
	var v24 int32
	_ = v24
	var v31 int32
	_ = v31
	var v34 int32
	_ = v34
	var v43 int32
	_ = v43
	var v45 int32
	_ = v45
	var v49 int32
	_ = v49
	var v50 int32
	_ = v50
	var v61 int32
	_ = v61
	var v64 int32
	_ = v64
	var v65 int32
	_ = v65
	var v68 int32
	_ = v68
	var v69 int32
	_ = v69
	var v70 int32
	_ = v70
	var v72 int32
	_ = v72
	var v73 int32
	_ = v73
	var v74 int32
	_ = v74
	var v77 int32
	_ = v77
	var v80 int32
	_ = v80
	var v85 int32
	_ = v85
	var v89 int32
	_ = v89
	var v95 int32
	_ = v95
	var v101 int32
	_ = v101
	var v102 int32
	_ = v102
	var v105 int32
	_ = v105
	var v108 int32
	_ = v108
	var v110 int32
	_ = v110
	var v111 int32
	_ = v111
	var v112 int32
	_ = v112
	var v115 int32
	_ = v115
	var v123 int32
	_ = v123
	var v127 int32
	_ = v127
	var v128 int32
	_ = v128
	var v131 int32
	_ = v131
	var v152 int32
	_ = v152
	var v153 int32
	_ = v153
	var v154 int32
	_ = v154
	var v159 int32
	_ = v159
	var v163 int32
	_ = v163
	var v165 int32
	_ = v165
	var v168 int32
	_ = v168
	var v170 int32
	_ = v170
	var v171 int32
	_ = v171
	var v184 int32
	_ = v184
	var v196 int32
	_ = v196
	var v197 int32
	_ = v197
	var v200 int32
	_ = v200
	var v201 int32
	_ = v201
	var v204 int32
	_ = v204
	var v209 int32
	_ = v209
	var v213 int32
	_ = v213
	var v217 int32
	_ = v217
	var v218 int32
	_ = v218
	var v222 int32
	_ = v222
	var v226 int32
	_ = v226
	var v231 int32
	_ = v231
	var v235 int32
	_ = v235
	var v241 int32
	_ = v241
	v12 = m.G0
	v14 = v12 - int32(16)
	m.G0 = v14
	v16 = int32(1)
	if l1 < int32(4) {
		v241 = v16
		goto L1
	} else {
		goto L2
	}
L1:
	;
	m.G0 = v14 + int32(16)
	return v241
L2:
	;
	v19 = int32(21)
	v24 = (l1<<(uint(v19)%32) - int32(_a_F_apply_typmod_0)) >> (uint(v19) % 32)
	*(*int32)(unsafe.Add(mBase, uint32(l0)+12)) = v24
	v31 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v34 = v24 + v31<<(uint(int32(2))%32)
	if v34+int32(4) < int32(0) {
		goto L4
	} else {
		goto L5
	}
L3:
	;
	v152 = int32(base.Ui32(l1-int32(4)) >> (uint(int32(16)) % 32))
	v153 = v152 - v24
	v154 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	if v154 < int32(0) {
		goto L29
	} else {
		goto L30
	}
L4:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+8)) = int32(0)
	*(*int64)(unsafe.Add(mBase, uint32(l0))) = int64(0)
	goto L3
L5:
	;
	goto L6
L6:
	;
	v43 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v45 = v24 & int32(3)
	v49 = base.I32_div_s(v34+int32(7), int32(4))
	v50 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	if v50 <= v49 {
		goto L11
	} else {
		goto L12
	}
L7:
	;
	goto L3
L8:
	;
	if int32(0) <= v115 {
		goto L7
	} else {
		goto L28
	}
L9:
	;
	v95 = v89
	goto L22
L10:
	;
	v64 = int32(1)
	v65 = v49 - v64
	v68 = v43 + v65<<(uint(v64)%32)
	v69 = int32(*(*int16)(unsafe.Add(mBase, uint32(v68))))
	v70 = int32(2)
	v72 = *(*int32)(unsafe.Add(mBase, uint32(v45<<(uint(v70)%32))+uint32(_c_F_apply_typmod[0])))
	v73 = base.I32_rem_s(v69, v72)
	v74 = v69 - v73
	*(*uint16)(unsafe.Add(mBase, uint32(v68))) = uint16(v74)
	v77 = base.I32_div_s(v72, v70)
	if v73 < v77 {
		v115 = v65
		goto L8
	} else {
		goto L17
	}
L11:
	;
	if base.B2i32(v45 == int32(0))|base.B2i32(v49 != v50) != 0 {
		goto L7
	} else {
		goto L14
	}
L12:
	;
	goto L13
L13:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0))) = v49
	if v45 != 0 {
		goto L10
	} else {
		goto L15
	}
L14:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0))) = v49
	goto L10
L15:
	;
	v61 = int32(*(*int16)(unsafe.Add(mBase, uint32(v43+v49<<(uint(int32(1))%32)))))
	if v61 <= int32(_a_F_apply_typmod_1) {
		v115 = v49
		goto L8
	} else {
		goto L16
	}
L16:
	;
	v89 = v49
	goto L9
L17:
	;
	v80 = v72 + base.I32_extend16_s(v74)
	if int32(_a_F_apply_typmod_2) < v80 {
		goto L18
	} else {
		goto L19
	}
L18:
	;
	v85 = v80 + int32(_a_F_apply_typmod_3)
	goto L20
L19:
	;
	v85 = v80
	goto L20
L20:
	;
	*(*uint16)(unsafe.Add(mBase, uint32(v68))) = uint16(v85)
	if v80 < int32(_a_F_apply_typmod_4) {
		v115 = v65
		goto L8
	} else {
		goto L21
	}
L21:
	;
	v89 = v65
	goto L9
L22:
	;
	v101 = int32(1)
	v102 = v95 - v101
	v105 = v43 + v102<<(uint(v101)%32)
	v108 = int32(*(*int16)(unsafe.Add(mBase, uint32(v105))))
	v110 = base.B2i32(int32(_a_F_apply_typmod_5) < v108)
	if int32(_a_F_apply_typmod_5) < v108 {
		goto L24
	} else {
		goto L25
	}
L23:
	;
	v115 = v102
	goto L8
L24:
	;
	v111 = int32(-9999)
	goto L26
L25:
	;
	v111 = v101
	goto L26
L26:
	;
	v112 = v111 + v108
	*(*uint16)(unsafe.Add(mBase, uint32(v105))) = uint16(v112)
	if int32(_a_F_apply_typmod_5) < v108 {
		v95 = v102
		goto L22
	} else {
		goto L27
	}
L27:
	;
	goto L23
L28:
	;
	v123 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+20)) = v123 - int32(2)
	v127 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v128 = int32(1)
	*(*int32)(unsafe.Add(mBase, uint32(l0))) = v127 + v128
	v131 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v131 + v128
	goto L7
L29:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+12)) = int32(0)
	goto L31
L30:
	;
	goto L31
L31:
	;
	v159 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v163 = v159<<(uint(int32(2))%32) + int32(4)
	if v163 <= v153 {
		v241 = v16
		goto L1
	} else {
		goto L32
	}
L32:
	;
	v165 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	if v165 <= int32(0) {
		v241 = v16
		goto L1
	} else {
		goto L33
	}
L33:
	;
	v168 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v170 = int32(0)
	v171 = v163
	goto L34
L34:
	;
	v184 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v168+v170<<(uint(int32(1))%32)))))
	if v184 != 0 {
		goto L36
	} else {
		goto L37
	}
L35:
	;
	v241 = v16
	goto L1
L36:
	;
	if base.I32_extend16_s(v184) < int32(10) {
		v197 = int32(-3)
		goto L39
	} else {
		goto L40
	}
L37:
	;
	goto L38
L38:
	;
	v235 = v170 + int32(1)
	if v235 != v165 {
		v170 = v235
		v171 = v171 - int32(4)
		goto L34
	} else {
		goto L59
	}
L39:
	;
	if v197+v171 <= v153 {
		v241 = v16
		goto L1
	} else {
		goto L45
	}
L40:
	;
	if base.Ui32(v184) < base.Ui32(int32(100)) {
		v197 = int32(-2)
		goto L39
	} else {
		goto L41
	}
L41:
	;
	if base.Ui32(v184) < base.Ui32(int32(1000)) {
		goto L42
	} else {
		goto L43
	}
L42:
	;
	v196 = int32(-1)
	goto L44
L43:
	;
	v196 = int32(0)
	goto L44
L44:
	;
	v197 = v196
	goto L39
L45:
	;
	v200 = int32(0)
	v201 = F_errsave_start(m, l2)
	mBase = m.M
	v204 = m.ExcPending
	if v204 != 0 {
		goto L46
	} else {
		goto L47
	}
L46:
	;
	return int32(0)
L47:
	;
	if v201 == int32(0) {
		v241 = v200
		goto L1
	} else {
		goto L48
	}
L48:
	;
	F_errcode(m, int32(50331778))
	mBase = m.M
	v209 = m.ExcPending
	if v209 != 0 {
		goto L46
	} else {
		goto L49
	}
L49:
	;
	F_errmsg(m, int32(_a_F_apply_typmod_6), int32(0))
	mBase = m.M
	v213 = m.ExcPending
	if v213 != 0 {
		goto L46
	} else {
		goto L50
	}
L50:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v14))) = v152
	*(*int32)(unsafe.Add(mBase, uint32(v14)+4)) = v24
	v217 = base.B2i32(v24 == v152)
	if v24 == v152 {
		goto L51
	} else {
		goto L52
	}
L51:
	;
	v218 = int32(1)
	goto L53
L52:
	;
	v218 = v153
	goto L53
L53:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v14)+12)) = v218
	if v24 == v152 {
		goto L54
	} else {
		goto L55
	}
L54:
	;
	v222 = int32(_a_F_apply_typmod_7)
	goto L56
L55:
	;
	v222 = int32(_a_F_apply_typmod_8)
	goto L56
L56:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v14)+8)) = v222
	F_errdetail(m, int32(_a_F_apply_typmod_9), v14)
	mBase = m.M
	v226 = m.ExcPending
	if v226 != 0 {
		goto L46
	} else {
		goto L57
	}
L57:
	;
	F_errsave_finish(m, l2, int32(_a_F_apply_typmod_10), int32(_a_F_apply_typmod_11), int32(_a_F_apply_typmod_12))
	mBase = m.M
	v231 = m.ExcPending
	if v231 != 0 {
		goto L46
	} else {
		goto L58
	}
L58:
	;
	v241 = v200
	goto L1
L59:
	;
	goto L35
}
func F_apw_compare_blockinfo(m *base.Module, l0 int32, l1 int32) int32 {
	mBase := m.M
	_ = mBase
	var v6 int32
	_ = v6
	var v7 int32
	_ = v7
	var v11 int32
	_ = v11
	var v13 int32
	_ = v13
	var v14 int32
	_ = v14
	var v19 int32
	_ = v19
	var v20 int32
	_ = v20
	var v25 int32
	_ = v25
	var v26 int32
	_ = v26
	var v32 int32
	_ = v32
	var v33 int32
	_ = v33
	var v40 int32
	_ = v40
	v6 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v7 = *(*int32)(unsafe.Add(mBase, uint32(l1)))
	if base.Ui32(v6) < base.Ui32(v7) {
		return int32(-1)
	} else {
		v11 = int32(1)
		if base.Ui32(v7) < base.Ui32(v6) {
			v40 = v11
			return v40
		} else {
			v13 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
			v14 = *(*int32)(unsafe.Add(mBase, uint32(l1)+4))
			if base.Ui32(v13) < base.Ui32(v14) {
				return int32(-1)
			} else {
				if base.Ui32(v14) < base.Ui32(v13) {
					v40 = v11
					return v40
				} else {
					v19 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
					v20 = *(*int32)(unsafe.Add(mBase, uint32(l1)+8))
					if base.Ui32(v19) < base.Ui32(v20) {
						return int32(-1)
					} else {
						if base.Ui32(v20) < base.Ui32(v19) {
							v40 = v11
							return v40
						} else {
							v25 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
							v26 = *(*int32)(unsafe.Add(mBase, uint32(l1)+12))
							if v25 < v26 {
								return int32(-1)
							} else {
								if v26 < v25 {
									v40 = v11
								} else {
									v32 = *(*int32)(unsafe.Add(mBase, uint32(l0)+16))
									v33 = *(*int32)(unsafe.Add(mBase, uint32(l1)+16))
									if base.Ui32(v32) < base.Ui32(v33) {
										v40 = int32(-1)
									} else {
										v40 = base.B2i32(base.Ui32(v33) < base.Ui32(v32))
									}
								}
								return v40
							}
						}
					}
				}
			}
		}
	}
}
func F_apw_read_stream_next_block(m *base.Module, l0 int32, l1 int32, l2 int32) int32 {
	mBase := m.M
	_ = mBase
	var v8 int32
	_ = v8
	var v12 int32
	_ = v12
	var v13 int32
	_ = v13
	var v17 int32
	_ = v17
	var v20 int32
	_ = v20
	var v22 int32
	_ = v22
	var v23 int32
	_ = v23
	var v25 int32
	_ = v25
	var v28 int32
	_ = v28
	var v29 int32
	_ = v29
	var v30 int32
	_ = v30
	var v31 int32
	_ = v31
	var v32 int32
	_ = v32
	var v34 int32
	_ = v34
	var v35 int32
	_ = v35
	var v43 int32
	_ = v43
	var v44 int32
	_ = v44
	var v48 int32
	_ = v48
	var v50 int32
	_ = v50
	var v52 int32
	_ = v52
	var v54 int32
	_ = v54
	var v56 int32
	_ = v56
	var v58 int32
	_ = v58
	var v60 int32
	_ = v60
	v8 = *(*int32)(unsafe.Add(mBase, _c_F_apw_read_stream_next_block[0]))
	if v8 != 0 {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	F_ProcessInterrupts(m)
	mBase = m.M
	v12 = m.ExcPending
	if v12 != 0 {
		goto L4
	} else {
		goto L5
	}
L2:
	;
	goto L3
L3:
	;
	v13 = *(*int32)(unsafe.Add(mBase, uint32(l1)+4))
	v17 = v13
	goto L6
L4:
	;
	return int32(0)
L5:
	;
	goto L3
L6:
	;
	v20 = int32(-1)
	v22 = *(*int32)(unsafe.Add(mBase, _c_F_apw_read_stream_next_block[1]))
	v23 = *(*int32)(unsafe.Add(mBase, uint32(v22)+36))
	if v23 <= v17 {
		v60 = v20
		goto L8
	} else {
		goto L9
	}
L7:
	;
	return v60
L8:
	;
	goto L7
L9:
	;
	v25 = *(*int32)(unsafe.Add(mBase, uint32(l1)))
	v28 = v25 + v17*int32(20)
	v29 = *(*int32)(unsafe.Add(mBase, uint32(v28)+16))
	v30 = *(*int32)(unsafe.Add(mBase, uint32(v28)+12))
	v31 = *(*int32)(unsafe.Add(mBase, uint32(v28)+8))
	v32 = *(*int32)(unsafe.Add(mBase, uint32(v28)+4))
	v34 = *(*int32)(unsafe.Add(mBase, _c_F_apw_read_stream_next_block[2]))
	v35 = *(*int32)(unsafe.Add(mBase, uint32(v34)+8))
	goto L10
L10:
	;
	if int32(base.Ui32(v35^int32(-1))>>(uint(int32(31))%32)) == int32(0) {
		goto L11
	} else {
		goto L12
	}
L11:
	;
	v43 = *(*int32)(unsafe.Add(mBase, _c_F_apw_read_stream_next_block[1]))
	v44 = *(*int32)(unsafe.Add(mBase, uint32(v43)+36))
	*(*int32)(unsafe.Add(mBase, uint32(l1)+4)) = v44
	return int32(-1)
L12:
	;
	goto L13
L13:
	;
	v48 = *(*int32)(unsafe.Add(mBase, uint32(l1)+8))
	if v32 != v48 {
		v60 = v20
		goto L8
	} else {
		goto L14
	}
L14:
	;
	v50 = *(*int32)(unsafe.Add(mBase, uint32(l1)+12))
	if v31 != v50 {
		v60 = v20
		goto L8
	} else {
		goto L15
	}
L15:
	;
	v52 = *(*int32)(unsafe.Add(mBase, uint32(l1)+16))
	if v30 != v52 {
		v60 = v20
		goto L8
	} else {
		goto L16
	}
L16:
	;
	v54 = *(*int32)(unsafe.Add(mBase, uint32(l1)+4))
	v56 = v54 + int32(1)
	*(*int32)(unsafe.Add(mBase, uint32(l1)+4)) = v56
	v58 = *(*int32)(unsafe.Add(mBase, uint32(l1)+20))
	if base.Ui32(v58) <= base.Ui32(v29) {
		v17 = v56
		goto L6
	} else {
		goto L17
	}
L17:
	;
	v60 = v29
	goto L8
}
func F_arraycontains(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v5 int32
	_ = v5
	var v6 int32
	_ = v6
	var v9 int32
	_ = v9
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
	var v18 int32
	_ = v18
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
	var v33 int32
	_ = v33
	v5 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v6 = F_DatumGetAnyArrayP(m, v5)
	mBase = m.M
	v9 = m.ExcPending
	if v9 != 0 {
		return int32(0)
	} else {
		v10 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
		v11 = F_DatumGetAnyArrayP(m, v10)
		mBase = m.M
		v12 = m.ExcPending
		if v12 != 0 {
			return int32(0)
		} else {
			v13 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
			v15 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
			v18 = F_array_contain_compare(m, v11, v6, v13, int32(1), v15+int32(16))
			mBase = m.M
			v19 = m.ExcPending
			if v19 != 0 {
				return int32(0)
			} else {
				v20 = *(*int32)(unsafe.Add(mBase, uint32(v6)))
				if v20 == int32(-1) {
					v27 = *(*int32)(unsafe.Add(mBase, uint32(v11)))
					if v27 == int32(-1) {
						return v18
					} else {
						v30 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
						if v11 == v30 {
							return v18
						} else {
							F_pfree(m, v11)
							mBase = m.M
							v33 = m.ExcPending
							if v33 != 0 {
								return int32(0)
							} else {
								return v18
							}
						}
					}
				} else {
					v23 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
					if v6 == v23 {
						v27 = *(*int32)(unsafe.Add(mBase, uint32(v11)))
						if v27 == int32(-1) {
							return v18
						} else {
							v30 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
							if v11 == v30 {
								return v18
							} else {
								F_pfree(m, v11)
								mBase = m.M
								v33 = m.ExcPending
								if v33 != 0 {
									return int32(0)
								} else {
									return v18
								}
							}
						}
					} else {
						F_pfree(m, v6)
						mBase = m.M
						v26 = m.ExcPending
						if v26 != 0 {
							return int32(0)
						} else {
							v27 = *(*int32)(unsafe.Add(mBase, uint32(v11)))
							if v27 == int32(-1) {
								return v18
							} else {
								v30 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
								if v11 == v30 {
									return v18
								} else {
									F_pfree(m, v11)
									mBase = m.M
									v33 = m.ExcPending
									if v33 != 0 {
										return int32(0)
									} else {
										return v18
									}
								}
							}
						}
					}
				}
			}
		}
	}
}
func F_arrayoverlap(m *base.Module, l0 int32) int32 {
	var v3 int32
	_ = v3
	var v6 int32
	_ = v6
	v3 = Fn13868(m, l0, int32(0))
	v6 = m.ExcPending
	if v6 != 0 {
		return int32(0)
	} else {
		return v3
	}
}
func F_assign_application_name(m *base.Module, l0 int32, l1 int32) {
	mBase := m.M
	_ = mBase
	var v6 int32
	_ = v6
	var v7 int32
	_ = v7
	var v9 int32
	_ = v9
	var v10 int32
	_ = v10
	var v11 int32
	_ = v11
	var v13 int32
	_ = v13
	var v14 int32
	_ = v14
	var v17 int32
	_ = v17
	var v21 int32
	_ = v21
	var v23 int32
	_ = v23
	var v25 int32
	_ = v25
	var v27 int32
	_ = v27
	var v28 int32
	_ = v28
	var v31 int32
	_ = v31
	var v33 int32
	_ = v33
	v6 = *(*int32)(unsafe.Add(mBase, _c_F_assign_application_name[0]))
	if v6 != 0 {
		v7 = F_strlen(m, l0)
		mBase = m.M
		v9 = F_pg_mbcliplen(m, l0, v7, int32(63))
		mBase = m.M
		v10 = m.ExcPending
		if v10 != 0 {
			return
		} else {
			v11 = int32(_a_F_assign_application_name_0)
			v13 = *(*int32)(unsafe.Add(mBase, _c_F_assign_application_name[1]))
			v14 = int32(1)
			*(*int32)(unsafe.Add(mBase, _c_F_assign_application_name[1])) = v13 + v14
			v17 = *(*int32)(unsafe.Add(mBase, uint32(v6)))
			*(*int32)(unsafe.Add(mBase, uint32(v6))) = v17 + v14
			v21 = *(*int32)(unsafe.Add(mBase, uint32(v6)+212))
			if v9 != 0 {
				base.MemoryCopy(m, v21, l0, v9)
			} else {
			}
			v23 = *(*int32)(unsafe.Add(mBase, uint32(v6)+212))
			v25 = int32(0)
			*(*uint8)(unsafe.Add(mBase, uint32(v23+v9))) = uint8(v25)
			v27 = *(*int32)(unsafe.Add(mBase, uint32(v6)))
			v28 = int32(1)
			*(*int32)(unsafe.Add(mBase, uint32(v6))) = v27 + v28
			v31 = int32(_a_F_assign_application_name_0)
			v33 = *(*int32)(unsafe.Add(mBase, _c_F_assign_application_name[1]))
			*(*int32)(unsafe.Add(mBase, _c_F_assign_application_name[1])) = v33 - v28
			return
		}
	} else {
		return
	}
}
func F_assign_backtrace_functions(m *base.Module, l0 int32, l1 int32) {
	mBase := m.M
	_ = mBase
	*(*int32)(unsafe.Add(mBase, _c_F_assign_backtrace_functions[0])) = l1
	return
}
func F_assign_session_authorization(m *base.Module, l0 int32, l1 int32) {
	mBase := m.M
	_ = mBase
	var v3 int32
	_ = v3
	var v4 int32
	_ = v4
	var v6 int32
	_ = v6
	if l1 != 0 {
		v3 = *(*int32)(unsafe.Add(mBase, uint32(l1)))
		v4 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l1)+4)))
		F_SetSessionAuthorization(m, v3, v4)
		mBase = m.M
		v6 = m.ExcPending
		if v6 != 0 {
			return
		} else {
			return
		}
	} else {
		return
	}
}
func F_assign_synchronized_standby_slots(m *base.Module, l0 int32, l1 int32) {
	mBase := m.M
	_ = mBase
	*(*int32)(unsafe.Add(mBase, _c_F_assign_synchronized_standby_slots[0])) = l1
	*(*int64)(unsafe.Add(mBase, _c_F_assign_synchronized_standby_slots[1])) = int64(0)
	return
}
func F_atof(m *base.Module, l0 int32) float64 {
	var v3 float64
	_ = v3
	var v6 int32
	_ = v6
	v3 = F_strtod(m, l0, int32(0))
	v6 = m.ExcPending
	if v6 != 0 {
		return float64(0)
	} else {
		return v3
	}
}
func F_attnumTypeId(m *base.Module, l0 int32, l1 int32) int32 {
	mBase := m.M
	_ = mBase
	var v5 int32
	_ = v5
	var v7 int32
	_ = v7
	var v12 int32
	_ = v12
	var v15 int32
	_ = v15
	var v16 int32
	_ = v16
	var v17 int32
	_ = v17
	var v29 int32
	_ = v29
	var v30 int32
	_ = v30
	var v38 int32
	_ = v38
	var v42 int32
	_ = v42
	var v47 int32
	_ = v47
	v5 = m.G0
	v7 = v5 - int32(16)
	m.G0 = v7
	if l1 <= int32(0) {
		v12 = F_SystemAttributeDefinition(m, base.I32_extend16_s(l1))
		mBase = m.M
		v15 = m.ExcPending
		if v15 != 0 {
			return int32(0)
		} else {
			v29 = v12
			v30 = *(*int32)(unsafe.Add(mBase, uint32(v29)+68))
			m.G0 = v7 + int32(16)
			return v30
		}
	} else {
		v16 = *(*int32)(unsafe.Add(mBase, uint32(l0)+52))
		v17 = *(*int32)(unsafe.Add(mBase, uint32(v16)))
		if v17 < l1 {
			F_errstart_cold(m, int32(21), int32(0))
			mBase = m.M
			v38 = m.ExcPending
			if v38 != 0 {
				return int32(0)
			} else {
				*(*int32)(unsafe.Add(mBase, uint32(v7))) = l1
				F_errmsg_internal(m, int32(_a_F_attnumTypeId_0), v7)
				mBase = m.M
				v42 = m.ExcPending
				if v42 != 0 {
					return int32(0)
				} else {
					F_errfinish(m, int32(_a_F_attnumTypeId_1), int32(3671), int32(_a_F_attnumTypeId_2))
					mBase = m.M
					v47 = m.ExcPending
					if v47 != 0 {
						return int32(0)
					} else {
						base.Wasm_trap_unreachable()
						for {
						}
					}
				}
			}
		} else {
			v29 = v16 + v17<<(uint(int32(4))%32) + l1*int32(100) - int32(80)
			v30 = *(*int32)(unsafe.Add(mBase, uint32(v29)+68))
			m.G0 = v7 + int32(16)
			return v30
		}
	}
}
