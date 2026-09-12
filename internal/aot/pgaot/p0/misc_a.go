package p0

import (
	base "github.com/shibukawa/pgmem/internal/aot/pgaot/base"
	"unsafe"
)

func F_AbortStrongLockAcquire(m *base.Module) {
	mBase := m.M
	_ = mBase
	var v6 int32
	_ = v6
	var v8 int32
	_ = v8
	var v9 int32
	_ = v9
	var v10 int32
	_ = v10
	var v16 int32
	_ = v16
	var v21 int32
	_ = v21
	var v23 int32
	_ = v23
	var v28 int32
	_ = v28
	var v29 int32
	_ = v29
	var v33 int32
	_ = v33
	v6 = *(*int32)(unsafe.Add(mBase, _c_F_AbortStrongLockAcquire[0]))
	if v6 != 0 {
		v8 = *(*int32)(unsafe.Add(mBase, _c_F_AbortStrongLockAcquire[1]))
		v9 = *(*int32)(unsafe.Add(mBase, uint32(v8)))
		v10 = *(*int32)(unsafe.Add(mBase, uint32(v6)+20))
		*(*int32)(unsafe.Add(mBase, uint32(v8))) = int32(1)
		if v9 != 0 {
			v16 = *(*int32)(unsafe.Add(mBase, _c_F_AbortStrongLockAcquire[1]))
			F_s_lock(m, v16, int32(_a_F_AbortStrongLockAcquire_0), int32(1869), int32(_a_F_AbortStrongLockAcquire_1))
			mBase = m.M
			v21 = m.ExcPending
			if v21 != 0 {
				return
			} else {
				v23 = *(*int32)(unsafe.Add(mBase, _c_F_AbortStrongLockAcquire[1]))
				v28 = v23 + v10&int32(1023)<<(uint(int32(2))%32) + int32(4)
				v29 = *(*int32)(unsafe.Add(mBase, uint32(v28)))
				*(*int32)(unsafe.Add(mBase, uint32(v28))) = v29 - int32(1)
				v33 = int32(0)
				*(*uint8)(unsafe.Add(mBase, uint32(v6)+52)) = uint8(v33)
				*(*int32)(unsafe.Add(mBase, _c_F_AbortStrongLockAcquire[0])) = v33
				*(*int32)(unsafe.Add(mBase, uint32(v23))) = v33
				return
			}
		} else {
			v23 = *(*int32)(unsafe.Add(mBase, _c_F_AbortStrongLockAcquire[1]))
			v28 = v23 + v10&int32(1023)<<(uint(int32(2))%32) + int32(4)
			v29 = *(*int32)(unsafe.Add(mBase, uint32(v28)))
			*(*int32)(unsafe.Add(mBase, uint32(v28))) = v29 - int32(1)
			v33 = int32(0)
			*(*uint8)(unsafe.Add(mBase, uint32(v6)+52)) = uint8(v33)
			*(*int32)(unsafe.Add(mBase, _c_F_AbortStrongLockAcquire[0])) = v33
			*(*int32)(unsafe.Add(mBase, uint32(v23))) = v33
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
	var v20 int32
	_ = v20
	var v21 int32
	_ = v21
	var v25 int32
	_ = v25
	var v26 int32
	_ = v26
	var v27 int32
	_ = v27
	var v29 int32
	_ = v29
	var v32 int32
	_ = v32
	var v34 int32
	_ = v34
	var v35 int32
	_ = v35
	var v37 int32
	_ = v37
	var v39 int32
	_ = v39
	var v46 int32
	_ = v46
	var v50 int32
	_ = v50
	var v53 int32
	_ = v53
	var v55 int32
	_ = v55
	var v60 int32
	_ = v60
	var v62 int32
	_ = v62
	var v65 int32
	_ = v65
	var v74 int32
	_ = v74
	var v76 int32
	_ = v76
	var v83 int32
	_ = v83
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
	v20 = int32(0)
	v21 = base.B2i32(v19 <= v20)
	if v21 == v20 {
		goto L5
	} else {
		goto L6
	}
L5:
	;
	v25 = v19 << (uint(int32(5)) % 32)
	v26 = F_palloc(m, v25)
	mBase = m.M
	v27 = m.ExcPending
	if v27 != 0 {
		goto L3
	} else {
		goto L8
	}
L6:
	;
	v34 = v18
	v35 = int32(0)
	goto L7
L7:
	;
	v37 = int32(_a_F_AbsorbSyncRequests_0)
	v39 = *(*int32)(unsafe.Add(mBase, _c_F_AbsorbSyncRequests[3]))
	*(*int32)(unsafe.Add(mBase, _c_F_AbsorbSyncRequests[3])) = v39 + int32(1)
	*(*int32)(unsafe.Add(mBase, uint32(v34)+48)) = int32(0)
	v46 = *(*int32)(unsafe.Add(mBase, _c_F_AbsorbSyncRequests[1]))
	F_LWLockRelease(m, v46+int32(2176))
	mBase = m.M
	v50 = m.ExcPending
	if v50 != 0 {
		goto L3
	} else {
		goto L13
	}
L8:
	;
	v29 = *(*int32)(unsafe.Add(mBase, _c_F_AbsorbSyncRequests[2]))
	if v25 != 0 {
		goto L10
	} else {
		goto L11
	}
L9:
	;
	v34 = v29
	v35 = v26
	goto L7
L10:
	;
	v32 = F__emscripten_memcpy_bulkmem(m, v26, v29+int32(56), v25)
	mBase = m.M
	goto L12
L11:
	;
	goto L12
L12:
	;
	goto L9
L13:
	;
	if v21 == int32(0) {
		goto L14
	} else {
		goto L15
	}
L14:
	;
	v53 = v35
	v55 = v19
	goto L17
L15:
	;
	goto L16
L16:
	;
	v74 = int32(_a_F_AbsorbSyncRequests_0)
	v76 = *(*int32)(unsafe.Add(mBase, _c_F_AbsorbSyncRequests[3]))
	*(*int32)(unsafe.Add(mBase, _c_F_AbsorbSyncRequests[3])) = v76 - int32(1)
	if v35 == int32(0) {
		goto L1
	} else {
		goto L21
	}
L17:
	;
	v60 = *(*int32)(unsafe.Add(mBase, uint32(v53)))
	F_RememberSyncRequest(m, v53+int32(8), v60)
	mBase = m.M
	v62 = m.ExcPending
	if v62 != 0 {
		goto L3
	} else {
		goto L19
	}
L18:
	;
	goto L16
L19:
	;
	v65 = int32(1)
	if base.Ui32(v65) < base.Ui32(v55) {
		v53 = v53 + int32(32)
		v55 = v55 - v65
		goto L17
	} else {
		goto L20
	}
L20:
	;
	goto L18
L21:
	;
	F_pfree(m, v35)
	mBase = m.M
	v83 = m.ExcPending
	if v83 != 0 {
		goto L3
	} else {
		goto L22
	}
L22:
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
	var v28 int32
	_ = v28
	var v29 int32
	_ = v29
	var v31 int32
	_ = v31
	var v32 int32
	_ = v32
	var v35 int32
	_ = v35
	var v36 int32
	_ = v36
	var v39 int32
	_ = v39
	var v40 int32
	_ = v40
	var v41 int32
	_ = v41
	var v44 int32
	_ = v44
	var v45 int32
	_ = v45
	var v48 int32
	_ = v48
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
		v28 = l0 + v23<<(uint(int32(2))%32) + int32(48)
		v29 = *(*int32)(unsafe.Add(mBase, uint32(v28)))
		if v29 != 0 {
			v31 = v29 + int32(8)
			v32 = *(*int32)(unsafe.Add(mBase, uint32(v31)))
			*(*int32)(unsafe.Add(mBase, uint32(v28))) = v32
			return v31
		} else {
			v35 = int32(8)
			v36 = v35 << (uint(v23) % 32)
			v39 = *(*int32)(unsafe.Add(mBase, uint32(l0)+44))
			v40 = *(*int32)(unsafe.Add(mBase, uint32(v39)+16))
			v41 = *(*int32)(unsafe.Add(mBase, uint32(v39)+12))
			if base.Ui32(v40-v41) < base.Ui32(v36+v35) {
				v44 = F_AllocSetAllocFromNewBlock(m, l0, l1, l2, v23)
				mBase = m.M
				v45 = m.ExcPending
				if v45 != 0 {
					return int32(0)
				} else {
					return v44
				}
			} else {
				v48 = int32(8)
				*(*int32)(unsafe.Add(mBase, uint32(v39)+12)) = v41 + v36 + v48
				*(*int64)(unsafe.Add(mBase, uint32(v41))) = base.I64_extend_i32_u(v23<<(uint(int32(5))%32)) | base.I64_extend_i32_u(v41-v39)<<(uint(int64(34))%64) | int64(3)
				return v41 + v48
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
	var v6 int32
	_ = v6
	var v8 int32
	_ = v8
	var v12 int32
	_ = v12
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
	var v33 int32
	_ = v33
	var v40 int32
	_ = v40
	var v41 int32
	_ = v41
	var v50 int32
	_ = v50
	v6 = l0 + int32(48)
	v8 = l0 + int32(92)
	if base.Ui32(v6) < base.Ui32(v8) {
		v12 = l0 + int32(52)
		if base.Ui32(v12) < base.Ui32(v8) {
			v14 = v8
		} else {
			v14 = v12
		}
		v23 = F__emscripten_memset_bulkmem(m, v6, base.I32_extend8_s(int32(0)), (v14-l0-int32(49))&int32(-4)+int32(4))
		mBase = m.M
	} else {
	}
	v25 = *(*int32)(unsafe.Add(mBase, uint32(l0)+44))
	v27 = l0 + int32(112)
	*(*int32)(unsafe.Add(mBase, uint32(l0)+44)) = v27
	if v25 != 0 {
		v30 = v25
		for {
			v33 = *(*int32)(unsafe.Add(mBase, uint32(v30)+8))
			if v30 == v27 {
				*(*int64)(unsafe.Add(mBase, uint32(v30)+4)) = int64(0)
				*(*int32)(unsafe.Add(mBase, uint32(v30)+12)) = v30 + int32(24)
			} else {
				v40 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
				v41 = *(*int32)(unsafe.Add(mBase, uint32(v30)+16))
				*(*int32)(unsafe.Add(mBase, uint32(l0)+8)) = v40 + (v30 - v41)
				F_emscripten_builtin_free(m, v30)
				mBase = m.M
			}
			if v33 != 0 {
				v30 = v33
				continue
			} else {
				break
			}
			break
		}
	} else {
	}
	v50 = *(*int32)(unsafe.Add(mBase, uint32(l0)+92))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+100)) = v50
	return
}
func F_AllocateDir(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v2 int32
	_ = v2
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
	var v21 int32
	_ = v21
	var v23 int32
	_ = v23
	var v33 int32
	_ = v33
	var v34 int32
	_ = v34
	var v36 int32
	_ = v36
	var v38 int32
	_ = v38
	var v42 int32
	_ = v42
	var v44 int32
	_ = v44
	var v46 int32
	_ = v46
	var v55 int32
	_ = v55
	var v64 int32
	_ = v64
	var v69 int32
	_ = v69
	var v70 int32
	_ = v70
	var v73 int32
	_ = v73
	var v77 int32
	_ = v77
	var v82 int32
	_ = v82
	var v84 int32
	_ = v84
	var v87 int32
	_ = v87
	var v93 int32
	_ = v93
	var v94 int32
	_ = v94
	var v96 int32
	_ = v96
	var v97 int32
	_ = v97
	var v103 int32
	_ = v103
	var v106 int32
	_ = v106
	var v108 int32
	_ = v108
	var v111 int32
	_ = v111
	var v116 int32
	_ = v116
	var v117 int32
	_ = v117
	var v119 int32
	_ = v119
	var v121 int32
	_ = v121
	var v125 int32
	_ = v125
	var v130 int32
	_ = v130
	var v138 int32
	_ = v138
	var v141 int32
	_ = v141
	var v144 int32
	_ = v144
	var v148 int32
	_ = v148
	var v153 int32
	_ = v153
	v2 = int32(0)
	v6 = m.G0
	v8 = v6 - int32(16)
	m.G0 = v8
	v10 = F_reserveAllocatedDesc(m)
	mBase = m.M
	v13 = m.ExcPending
	if v13 != 0 {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	return int32(0)
L2:
	;
	if v10 != 0 {
		goto L3
	} else {
		goto L4
	}
L3:
	;
	v15 = *(*int32)(unsafe.Add(mBase, _c_F_AllocateDir[0]))
	if v15 <= int32(0) {
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
	v138 = m.ExcPending
	if v138 != 0 {
		goto L1
	} else {
		goto L34
	}
L6:
	;
	v55 = F_opendir(m, l0)
	mBase = m.M
	if v55 == int32(0) {
		goto L15
	} else {
		goto L16
	}
L7:
	;
	v19 = *(*int32)(unsafe.Add(mBase, _c_F_AllocateDir[1]))
	v21 = *(*int32)(unsafe.Add(mBase, _c_F_AllocateDir[2]))
	v23 = *(*int32)(unsafe.Add(mBase, _c_F_AllocateDir[3]))
	if v21+(v23+v15) < v19 {
		goto L6
	} else {
		goto L8
	}
L8:
	;
	goto L9
L9:
	;
	v33 = *(*int32)(unsafe.Add(mBase, _c_F_AllocateDir[4]))
	v34 = *(*int32)(unsafe.Add(mBase, uint32(v33)+16))
	F_LruDelete(m, v34)
	mBase = m.M
	v36 = m.ExcPending
	if v36 != 0 {
		goto L1
	} else {
		goto L11
	}
L10:
	;
	goto L6
L11:
	;
	v38 = *(*int32)(unsafe.Add(mBase, _c_F_AllocateDir[0]))
	if v38 <= int32(0) {
		goto L6
	} else {
		goto L12
	}
L12:
	;
	v42 = *(*int32)(unsafe.Add(mBase, _c_F_AllocateDir[1]))
	v44 = *(*int32)(unsafe.Add(mBase, _c_F_AllocateDir[2]))
	v46 = *(*int32)(unsafe.Add(mBase, _c_F_AllocateDir[3]))
	if v42 <= v44+(v46+v38) {
		goto L9
	} else {
		goto L13
	}
L13:
	;
	goto L10
L14:
	;
	m.G0 = v8 + int32(16)
	return v130
L15:
	;
	goto L18
L16:
	;
	v103 = v55
	goto L17
L17:
	;
	v106 = *(*int32)(unsafe.Add(mBase, _c_F_AllocateDir[5]))
	v108 = *(*int32)(unsafe.Add(mBase, _c_F_AllocateDir[3]))
	v111 = v106 + v108*int32(12)
	*(*int32)(unsafe.Add(mBase, uint32(v111)+8)) = v103
	*(*int32)(unsafe.Add(mBase, uint32(v111))) = int32(2)
	v116 = *(*int32)(unsafe.Add(mBase, _c_F_AllocateDir[6]))
	v117 = *(*int32)(unsafe.Add(mBase, uint32(v116)+8))
	goto L33
L18:
	;
	v64 = *(*int32)(unsafe.Add(mBase, _c_F_AllocateDir[7]))
	switch v64 - int32(33) {
	case 0, 8:
		goto L20
	default:
		v130 = v2
		goto L14
	}
L19:
	;
	v103 = v97
	goto L17
L20:
	;
	v69 = F_errstart(m, int32(15), int32(0))
	mBase = m.M
	v70 = m.ExcPending
	if v70 != 0 {
		goto L1
	} else {
		goto L21
	}
L21:
	;
	if v69 != 0 {
		goto L22
	} else {
		goto L23
	}
L22:
	;
	F_errcode(m, int32(197))
	mBase = m.M
	v73 = m.ExcPending
	if v73 != 0 {
		goto L1
	} else {
		goto L25
	}
L23:
	;
	goto L24
L24:
	;
	v84 = int32(0)
	*(*int32)(unsafe.Add(mBase, _c_F_AllocateDir[7])) = v84
	v87 = *(*int32)(unsafe.Add(mBase, _c_F_AllocateDir[0]))
	if v87 <= v84 {
		goto L28
	} else {
		goto L29
	}
L25:
	;
	F_errmsg(m, int32(_a_F_AllocateDir_0), int32(0))
	mBase = m.M
	v77 = m.ExcPending
	if v77 != 0 {
		goto L1
	} else {
		goto L26
	}
L26:
	;
	F_errfinish(m, int32(_a_F_AllocateDir_1), int32(2947), int32(_a_F_AllocateDir_2))
	mBase = m.M
	v82 = m.ExcPending
	if v82 != 0 {
		goto L1
	} else {
		goto L27
	}
L27:
	;
	goto L24
L28:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_AllocateDir[7])) = v64
	v130 = v2
	goto L14
L29:
	;
	goto L30
L30:
	;
	v93 = *(*int32)(unsafe.Add(mBase, _c_F_AllocateDir[4]))
	v94 = *(*int32)(unsafe.Add(mBase, uint32(v93)+16))
	F_LruDelete(m, v94)
	mBase = m.M
	v96 = m.ExcPending
	if v96 != 0 {
		goto L1
	} else {
		goto L31
	}
L31:
	;
	v97 = F_opendir(m, l0)
	mBase = m.M
	if v97 == int32(0) {
		goto L18
	} else {
		goto L32
	}
L32:
	;
	goto L19
L33:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v111)+4)) = v117
	v119 = int32(_a_F_AllocateDir_3)
	v121 = *(*int32)(unsafe.Add(mBase, _c_F_AllocateDir[3]))
	*(*int32)(unsafe.Add(mBase, _c_F_AllocateDir[3])) = v121 + int32(1)
	v125 = *(*int32)(unsafe.Add(mBase, uint32(v111)+8))
	v130 = v125
	goto L14
L34:
	;
	F_errcode(m, int32(197))
	mBase = m.M
	v141 = m.ExcPending
	if v141 != 0 {
		goto L1
	} else {
		goto L35
	}
L35:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v8)+4)) = l0
	v144 = *(*int32)(unsafe.Add(mBase, _c_F_AllocateDir[8]))
	*(*int32)(unsafe.Add(mBase, uint32(v8))) = v144
	F_errmsg(m, int32(_a_F_AllocateDir_4), v8)
	mBase = m.M
	v148 = m.ExcPending
	if v148 != 0 {
		goto L1
	} else {
		goto L36
	}
L36:
	;
	F_errfinish(m, int32(_a_F_AllocateDir_1), int32(2924), int32(_a_F_AllocateDir_2))
	mBase = m.M
	v153 = m.ExcPending
	if v153 != 0 {
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
	var v35 int32
	_ = v35
	var v36 int32
	_ = v36
	var v37 int32
	_ = v37
	var v40 int32
	_ = v40
	var v41 int32
	_ = v41
	var v44 int32
	_ = v44
	var v46 int32
	_ = v46
	var v47 int32
	_ = v47
	var v48 int32
	_ = v48
	var v51 int32
	_ = v51
	var v54 int32
	_ = v54
	var v55 int32
	_ = v55
	var v57 int32
	_ = v57
	var v58 int32
	_ = v58
	var v59 int32
	_ = v59
	var v63 int32
	_ = v63
	var v64 int32
	_ = v64
	var v66 int32
	_ = v66
	var v69 int32
	_ = v69
	var v71 int32
	_ = v71
	var v72 int32
	_ = v72
	var v76 int32
	_ = v76
	var v77 int32
	_ = v77
	var v79 int32
	_ = v79
	var v87 int32
	_ = v87
	var v88 int32
	_ = v88
	var v89 int32
	_ = v89
	var v92 int32
	_ = v92
	var v94 int32
	_ = v94
	var v99 int32
	_ = v99
	var v100 int32
	_ = v100
	var v105 int32
	_ = v105
	var v106 int32
	_ = v106
	var v107 int32
	_ = v107
	var v108 int32
	_ = v108
	var v111 int32
	_ = v111
	var v113 int32
	_ = v113
	var v116 int32
	_ = v116
	var v118 int32
	_ = v118
	var v119 int32
	_ = v119
	var v123 int32
	_ = v123
	var v124 int32
	_ = v124
	var v125 int32
	_ = v125
	var v126 int32
	_ = v126
	var v129 int32
	_ = v129
	var v133 int32
	_ = v133
	var v135 int32
	_ = v135
	var v136 int32
	_ = v136
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
	var v149 int32
	_ = v149
	var v150 int32
	_ = v150
	var v160 int32
	_ = v160
	var v161 int32
	_ = v161
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
	var v172 int32
	_ = v172
	var v173 int32
	_ = v173
	var v177 int32
	_ = v177
	var v184 int32
	_ = v184
	var v188 int32
	_ = v188
	var v200 int32
	_ = v200
	var v201 int32
	_ = v201
	var v208 int32
	_ = v208
	var v213 int32
	_ = v213
	var v217 int32
	_ = v217
	var v220 int32
	_ = v220
	var v221 int32
	_ = v221
	var v222 int32
	_ = v222
	var v228 int32
	_ = v228
	var v233 int32
	_ = v233
	var v237 int32
	_ = v237
	var v243 int32
	_ = v243
	var v248 int32
	_ = v248
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
	v237 = m.ExcPending
	if v237 != 0 {
		goto L1
	} else {
		goto L77
	}
L8:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v217 = m.ExcPending
	if v217 != 0 {
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
	v35 = F_heap_getattr_2(m, v30, v23, v32, v14+int32(47))
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
	v200 = m.ExcPending
	if v200 != 0 {
		goto L1
	} else {
		goto L69
	}
L13:
	;
	v37 = *(*int32)(unsafe.Add(mBase, uint32(l0)+52))
	v40 = F_heap_getattr_2(m, v30, v25, v37, v14+int32(47))
	mBase = m.M
	v41 = m.ExcPending
	if v41 != 0 {
		goto L1
	} else {
		goto L16
	}
L14:
	;
	m.G0 = v14 + int32(48)
	return v40
L15:
	;
	v184 = int32(0)
	F_RunObjectPostAlterHook(m, v16, l1, v184, v184, v184)
	mBase = m.M
	v188 = m.ExcPending
	if v188 != 0 {
		goto L1
	} else {
		goto L68
	}
L16:
	;
	if l2 == v40 {
		goto L17
	} else {
		goto L18
	}
L17:
	;
	v44 = *(*int32)(unsafe.Add(mBase, _c_F_AlterObjectNamespace_internal[0]))
	if v44 != 0 {
		goto L15
	} else {
		goto L20
	}
L18:
	;
	goto L19
L19:
	;
	F_CheckSetNamespace(m, v40, l2)
	mBase = m.M
	v46 = m.ExcPending
	if v46 != 0 {
		goto L1
	} else {
		goto L21
	}
L20:
	;
	goto L14
L21:
	;
	v47 = F_superuser(m)
	mBase = m.M
	v48 = m.ExcPending
	if v48 != 0 {
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
	if v47 != 0 {
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
	v51 = *(*int32)(unsafe.Add(mBase, uint32(l0)+52))
	v54 = F_heap_getattr_2(m, v30, v27, v51, v14+int32(47))
	mBase = m.M
	v55 = m.ExcPending
	if v55 != 0 {
		goto L1
	} else {
		goto L26
	}
L26:
	;
	v57 = *(*int32)(unsafe.Add(mBase, _c_F_AlterObjectNamespace_internal[1]))
	v58 = F_has_privs_of_role(m, v57, v54)
	mBase = m.M
	v59 = m.ExcPending
	if v59 != 0 {
		goto L1
	} else {
		goto L27
	}
L27:
	;
	if v58 == int32(0) {
		goto L28
	} else {
		goto L29
	}
L28:
	;
	v63 = F_get_object_type(m, v16, l1)
	mBase = m.M
	v64 = m.ExcPending
	if v64 != 0 {
		goto L1
	} else {
		goto L31
	}
L29:
	;
	goto L30
L30:
	;
	v69 = *(*int32)(unsafe.Add(mBase, _c_F_AlterObjectNamespace_internal[1]))
	v71 = F_object_aclcheck(m, int32(2615), l2, v69, int64(512))
	mBase = m.M
	v72 = m.ExcPending
	if v72 != 0 {
		goto L1
	} else {
		goto L33
	}
L31:
	;
	F_aclcheck_error(m, int32(2), v63, v35)
	mBase = m.M
	v66 = m.ExcPending
	if v66 != 0 {
		goto L1
	} else {
		goto L32
	}
L32:
	;
	goto L30
L33:
	;
	if v71 == int32(0) {
		goto L22
	} else {
		goto L34
	}
L34:
	;
	v76 = F_get_namespace_name(m, l2)
	mBase = m.M
	v77 = m.ExcPending
	if v77 != 0 {
		goto L1
	} else {
		goto L35
	}
L35:
	;
	F_aclcheck_error(m, v71, int32(36), v76)
	mBase = m.M
	v79 = m.ExcPending
	if v79 != 0 {
		goto L1
	} else {
		goto L36
	}
L36:
	;
	goto L22
L37:
	;
	v135 = *(*int32)(unsafe.Add(mBase, uint32(l0)+48))
	v136 = int32(*(*int16)(unsafe.Add(mBase, uint32(v135)+120)))
	v139 = F_palloc0(m, v136<<(uint(int32(2))%32))
	mBase = m.M
	v140 = m.ExcPending
	if v140 != 0 {
		goto L1
	} else {
		goto L57
	}
L38:
	;
	v124 = *(*int32)(unsafe.Add(mBase, uint32(v30)+16))
	v125 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v124)+22)))
	v126 = v124 + v125
	v129 = int32(*(*int16)(unsafe.Add(mBase, uint32(v126)+104)))
	F_IsThereFunctionInNamespace(m, v126+int32(4), v129, v126+int32(112), l2)
	mBase = m.M
	v133 = m.ExcPending
	if v133 != 0 {
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
	v87 = *(*int32)(unsafe.Add(mBase, uint32(v30)+16))
	v88 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v87)+22)))
	v89 = v87 + v88
	v92 = *(*int32)(unsafe.Add(mBase, uint32(v89)+4))
	F_IsThereOpClassInNamespace(m, v89+int32(8), v92, l2)
	mBase = m.M
	v94 = m.ExcPending
	if v94 != 0 {
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
	v106 = *(*int32)(unsafe.Add(mBase, uint32(v30)+16))
	v107 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v106)+22)))
	v108 = v106 + v107
	v111 = *(*int32)(unsafe.Add(mBase, uint32(v108)+4))
	F_IsThereOpFamilyInNamespace(m, v108+int32(8), v111, l2)
	mBase = m.M
	v113 = m.ExcPending
	if v113 != 0 {
		goto L1
	} else {
		goto L51
	}
L49:
	;
	v99 = *(*int32)(unsafe.Add(mBase, uint32(v30)+16))
	v100 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v99)+22)))
	F_IsThereCollationInNamespace(m, v99+v100+int32(4), l2)
	mBase = m.M
	v105 = m.ExcPending
	if v105 != 0 {
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
	v116 = int32(0)
	v118 = F_SearchSysCacheExists(m, v21, v35, l2, v116, v116)
	mBase = m.M
	v119 = m.ExcPending
	if v119 != 0 {
		goto L1
	} else {
		goto L53
	}
L53:
	;
	if v118 == int32(0) {
		goto L37
	} else {
		goto L54
	}
L54:
	;
	F_report_namespace_conflict(m, v16, v35, l2)
	mBase = m.M
	v123 = m.ExcPending
	if v123 != 0 {
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
	v141 = *(*int32)(unsafe.Add(mBase, uint32(l0)+48))
	v142 = int32(*(*int16)(unsafe.Add(mBase, uint32(v141)+120)))
	v143 = F_palloc0(m, v142)
	mBase = m.M
	v144 = m.ExcPending
	if v144 != 0 {
		goto L1
	} else {
		goto L58
	}
L58:
	;
	v145 = *(*int32)(unsafe.Add(mBase, uint32(l0)+48))
	v146 = int32(*(*int16)(unsafe.Add(mBase, uint32(v145)+120)))
	v147 = F_palloc0(m, v146)
	mBase = m.M
	v148 = m.ExcPending
	if v148 != 0 {
		goto L1
	} else {
		goto L59
	}
L59:
	;
	v149 = int32(1)
	v150 = v25 - v149
	*(*int32)(unsafe.Add(mBase, uint32(v139+v150<<(uint(int32(2))%32)))) = l2
	*(*uint8)(unsafe.Add(mBase, uint32(v147+v150))) = uint8(v149)
	v160 = *(*int32)(unsafe.Add(mBase, uint32(l0)+52))
	v161 = F_heap_modify_tuple(m, v30, v160, v139, v143, v147)
	mBase = m.M
	v162 = m.ExcPending
	if v162 != 0 {
		goto L1
	} else {
		goto L60
	}
L60:
	;
	F_CatalogTupleUpdate(m, l0, v30+int32(4), v161)
	mBase = m.M
	v164 = m.ExcPending
	if v164 != 0 {
		goto L1
	} else {
		goto L61
	}
L61:
	;
	F_pfree(m, v139)
	mBase = m.M
	v166 = m.ExcPending
	if v166 != 0 {
		goto L1
	} else {
		goto L62
	}
L62:
	;
	F_pfree(m, v143)
	mBase = m.M
	v168 = m.ExcPending
	if v168 != 0 {
		goto L1
	} else {
		goto L63
	}
L63:
	;
	F_pfree(m, v147)
	mBase = m.M
	v170 = m.ExcPending
	if v170 != 0 {
		goto L1
	} else {
		goto L64
	}
L64:
	;
	v172 = F_changeDependencyFor(m, v16, l1, int32(2615), v40, l2)
	mBase = m.M
	v173 = m.ExcPending
	if v173 != 0 {
		goto L1
	} else {
		goto L65
	}
L65:
	;
	if v172 != int32(1) {
		goto L7
	} else {
		goto L66
	}
L66:
	;
	v177 = *(*int32)(unsafe.Add(mBase, _c_F_AlterObjectNamespace_internal[0]))
	if v177 == int32(0) {
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
	v201 = *(*int32)(unsafe.Add(mBase, uint32(l0)+48))
	*(*int32)(unsafe.Add(mBase, uint32(v14))) = l1
	*(*int32)(unsafe.Add(mBase, uint32(v14)+4)) = v201 + int32(4)
	F_errmsg_internal(m, int32(_a_F_AlterObjectNamespace_internal_4), v14)
	mBase = m.M
	v208 = m.ExcPending
	if v208 != 0 {
		goto L1
	} else {
		goto L70
	}
L70:
	;
	F_errfinish(m, int32(_a_F_AlterObjectNamespace_internal_1), int32(713), int32(_a_F_AlterObjectNamespace_internal_2))
	mBase = m.M
	v213 = m.ExcPending
	if v213 != 0 {
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
	v220 = m.ExcPending
	if v220 != 0 {
		goto L1
	} else {
		goto L73
	}
L73:
	;
	v221 = F_getObjectDescriptionOids(m, v16, l1)
	mBase = m.M
	v222 = m.ExcPending
	if v222 != 0 {
		goto L1
	} else {
		goto L74
	}
L74:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v14)+32)) = v221
	F_errmsg(m, int32(_a_F_AlterObjectNamespace_internal_3), v14+int32(32))
	mBase = m.M
	v228 = m.ExcPending
	if v228 != 0 {
		goto L1
	} else {
		goto L75
	}
L75:
	;
	F_errfinish(m, int32(_a_F_AlterObjectNamespace_internal_1), int32(747), int32(_a_F_AlterObjectNamespace_internal_2))
	mBase = m.M
	v233 = m.ExcPending
	if v233 != 0 {
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
	v243 = m.ExcPending
	if v243 != 0 {
		goto L1
	} else {
		goto L78
	}
L78:
	;
	F_errfinish(m, int32(_a_F_AlterObjectNamespace_internal_1), int32(825), int32(_a_F_AlterObjectNamespace_internal_2))
	mBase = m.M
	v248 = m.ExcPending
	if v248 != 0 {
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
										F_sequence_close(m, v13, int32(3))
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
									F_sequence_close(m, v13, int32(3))
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
										F_sequence_close(m, v13, int32(3))
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
									F_sequence_close(m, v13, int32(3))
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
	var v355 int32
	_ = v355
	var v357 int32
	_ = v357
	var v358 int32
	_ = v358
	var v361 int32
	_ = v361
	var v362 int32
	_ = v362
	var v365 int32
	_ = v365
	var v370 int32
	_ = v370
	var v380 int32
	_ = v380
	var v385 int32
	_ = v385
	var v386 int32
	_ = v386
	var v392 int32
	_ = v392
	var v393 int32
	_ = v393
	var v396 int32
	_ = v396
	var v400 int32
	_ = v400
	var v407 int32
	_ = v407
	var v453 int32
	_ = v453
	var v454 int32
	_ = v454
	var v456 int32
	_ = v456
	var v457 int32
	_ = v457
	var v458 int32
	_ = v458
	var v460 int32
	_ = v460
	var v461 int32
	_ = v461
	var v462 int32
	_ = v462
	var v464 int32
	_ = v464
	var v465 int32
	_ = v465
	var v467 int32
	_ = v467
	var v469 int32
	_ = v469
	var v473 int32
	_ = v473
	var v474 int32
	_ = v474
	var v475 int32
	_ = v475
	var v476 int32
	_ = v476
	var v480 int32
	_ = v480
	var v484 int32
	_ = v484
	var v488 int32
	_ = v488
	var v489 int32
	_ = v489
	var v490 int32
	_ = v490
	var v491 int32
	_ = v491
	var v495 int32
	_ = v495
	var v496 int32
	_ = v496
	var v497 int32
	_ = v497
	var v499 int32
	_ = v499
	var v513 int32
	_ = v513
	var v514 int32
	_ = v514
	var v516 int32
	_ = v516
	var v517 int32
	_ = v517
	var v518 int32
	_ = v518
	var v520 int32
	_ = v520
	var v521 int32
	_ = v521
	var v522 int32
	_ = v522
	var v524 int32
	_ = v524
	var v525 int32
	_ = v525
	var v527 int32
	_ = v527
	var v529 int32
	_ = v529
	var v533 int32
	_ = v533
	var v534 int32
	_ = v534
	var v535 int32
	_ = v535
	var v536 int32
	_ = v536
	var v540 int32
	_ = v540
	var v544 int32
	_ = v544
	var v548 int32
	_ = v548
	var v549 int32
	_ = v549
	var v550 int32
	_ = v550
	var v551 int32
	_ = v551
	var v555 int32
	_ = v555
	var v556 int32
	_ = v556
	var v557 int32
	_ = v557
	var v559 int32
	_ = v559
	var v571 int32
	_ = v571
	var v575 int32
	_ = v575
	var v576 int32
	_ = v576
	var v580 int32
	_ = v580
	var v581 int32
	_ = v581
	var v585 int32
	_ = v585
	var v586 int32
	_ = v586
	var v588 int32
	_ = v588
	var v590 int32
	_ = v590
	var v594 int32
	_ = v594
	var v595 int32
	_ = v595
	var v599 int32
	_ = v599
	var v600 int32
	_ = v600
	var v602 int32
	_ = v602
	var v603 int32
	_ = v603
	var v605 int32
	_ = v605
	var v609 int32
	_ = v609
	var v610 int32
	_ = v610
	var v614 int32
	_ = v614
	var v615 int32
	_ = v615
	var v617 int32
	_ = v617
	var v621 int32
	_ = v621
	var v622 int32
	_ = v622
	var v626 int32
	_ = v626
	var v627 int32
	_ = v627
	var v631 int32
	_ = v631
	var v632 int32
	_ = v632
	var v636 int32
	_ = v636
	var v637 int32
	_ = v637
	var v638 int32
	_ = v638
	var v642 int32
	_ = v642
	var v643 int32
	_ = v643
	var v644 int32
	_ = v644
	var v648 int32
	_ = v648
	var v649 int32
	_ = v649
	var v650 int32
	_ = v650
	var v652 int32
	_ = v652
	var v653 int32
	_ = v653
	var v654 int32
	_ = v654
	var v658 int32
	_ = v658
	var v659 int32
	_ = v659
	var v660 int32
	_ = v660
	var v661 int32
	_ = v661
	var v665 int32
	_ = v665
	var v666 int32
	_ = v666
	var v667 int32
	_ = v667
	var v668 int32
	_ = v668
	var v672 int32
	_ = v672
	var v673 int32
	_ = v673
	var v674 int32
	_ = v674
	var v675 int32
	_ = v675
	var v680 int32
	_ = v680
	var v681 int32
	_ = v681
	var v682 int32
	_ = v682
	var v685 int32
	_ = v685
	var v687 int32
	_ = v687
	var v691 int32
	_ = v691
	var v695 int32
	_ = v695
	var v699 int32
	_ = v699
	var v703 int32
	_ = v703
	var v707 int32
	_ = v707
	var v717 int32
	_ = v717
	var v718 int32
	_ = v718
	var v719 int64
	_ = v719
	var v722 int32
	_ = v722
	var v725 int32
	_ = v725
	var v727 int32
	_ = v727
	var v728 int32
	_ = v728
	var v731 int32
	_ = v731
	var v732 int32
	_ = v732
	var v735 int32
	_ = v735
	var v737 int32
	_ = v737
	v8 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
	if v8 == int32(0) {
		goto L2
	} else {
		goto L3
	}
L1:
	;
	v385 = int32(4)
	v386 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	if base.Ui32(v380-int32(1021)) < base.Ui32(v385) {
		goto L67
	} else {
		goto L68
	}
L2:
	;
	v11 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v380 = v11
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
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v370
	v380 = v370
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
	v365 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v370 = v365 + int32(4)
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
	v370 = v361
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
	v355 = int32(1024) - v352
	if base.Ui32(v25) < base.Ui32(v355) {
		goto L59
	} else {
		goto L60
	}
L14:
	;
	v347 = F_Int64GetDatum(m, base.I64_extend_i32_u(v337)<<(uint(int64(32))%64)|base.I64_extend_i32_u(v329^v337-base.I32_rotl(v337, int32(24))))
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
	v357 = v25
	goto L61
L60:
	;
	v357 = v355
	goto L61
L61:
	;
	if v357 != 0 {
		goto L63
	} else {
		goto L64
	}
L62:
	;
	v361 = v352 + v357
	v362 = v25 - v357
	if v362 != 0 {
		v23 = v361
		v25 = v362
		v27 = v357 + v27
		goto L9
	} else {
		goto L66
	}
L63:
	;
	v358 = F__emscripten_memcpy_bulkmem(m, v352+v13, v27, v357)
	mBase = m.M
	goto L65
L64:
	;
	goto L65
L65:
	;
	goto L62
L66:
	;
	goto L10
L67:
	;
	v392 = l1
	v393 = v380
	v396 = v385
	goto L70
L68:
	;
	goto L69
L69:
	;
	v735 = *(*int32)(unsafe.Add(mBase, uint32(l1)))
	*(*int32)(unsafe.Add(mBase, uint32(v380+v386))) = v735
	v737 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v737 + int32(4)
	return
L70:
	;
	if base.Ui32(int32(1024)) <= base.Ui32(v393) {
		goto L72
	} else {
		goto L73
	}
L71:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v731
	return
L72:
	;
	v400 = int32(1024)
	v407 = int32(-1636607408)
	goto L77
L73:
	;
	v722 = v393
	goto L74
L74:
	;
	v725 = int32(1024) - v722
	if base.Ui32(v396) < base.Ui32(v725) {
		goto L119
	} else {
		goto L120
	}
L75:
	;
	v717 = F_Int64GetDatum(m, base.I64_extend_i32_u(v707)<<(uint(int64(32))%64)|base.I64_extend_i32_u(v699^v707-base.I32_rotl(v707, int32(24))))
	mBase = m.M
	v718 = m.ExcPending
	if v718 != 0 {
		goto L57
	} else {
		goto L118
	}
L76:
	;
	if v386&int32(3) != 0 {
		goto L92
	} else {
		goto L93
	}
L77:
	;
	goto L76
L80:
	;
	v685 = int32(14)
	v687 = v681 ^ v682 - base.I32_rotl(v681, v685)
	v691 = v687 ^ v680 - base.I32_rotl(v687, int32(11))
	v695 = v691 ^ v681 - base.I32_rotl(v691, int32(25))
	v699 = v695 ^ v687 - base.I32_rotl(v695, int32(16))
	v703 = v699 ^ v691 - base.I32_rotl(v699, int32(4))
	v707 = v703 ^ v695 - base.I32_rotl(v703, v685)
	goto L75
L81:
	;
	v675 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v497))))
	v680 = v672 + v675
	v681 = v673
	v682 = v674
	goto L80
L82:
	;
	v668 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v497)+1)))
	v672 = v668<<(uint(int32(8))%32) + v665
	v673 = v666
	v674 = v667
	goto L81
L83:
	;
	v661 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v497)+2)))
	v665 = v661<<(uint(int32(16))%32) + v658
	v666 = v659
	v667 = v660
	goto L82
L84:
	;
	v654 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v497)+3)))
	v658 = v654<<(uint(int32(24))%32) + v490
	v659 = v652
	v660 = v653
	goto L83
L85:
	;
	v650 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v497)+4)))
	v652 = v648 + v650
	v653 = v649
	goto L84
L86:
	;
	v644 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v497)+5)))
	v648 = v644<<(uint(int32(8))%32) + v642
	v649 = v643
	goto L85
L87:
	;
	v638 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v497)+6)))
	v642 = v638<<(uint(int32(16))%32) + v636
	v643 = v637
	goto L86
L88:
	;
	v632 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v497)+7)))
	v636 = v632<<(uint(int32(24))%32) + v491
	v637 = v631
	goto L87
L89:
	;
	v627 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v497)+8)))
	v631 = v627<<(uint(int32(8))%32) + v626
	goto L88
L90:
	;
	v622 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v497)+9)))
	v626 = v622<<(uint(int32(16))%32) + v621
	goto L89
L91:
	;
	v617 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v497)+10)))
	v621 = v617<<(uint(int32(24))%32) + v495
	goto L90
L92:
	;
	goto L95
L93:
	;
	goto L94
L94:
	;
	goto L101
L95:
	;
	v453 = v386
	v454 = v400
	v456 = v407
	v457 = v407
	v458 = v407
	goto L98
L97:
	;
	switch v499 - int32(1) {
	case 0:
		v672 = v490
		v673 = v491
		v674 = v495
		goto L81
	case 1:
		v665 = v490
		v666 = v491
		v667 = v495
		goto L82
	case 2:
		v658 = v490
		v659 = v491
		v660 = v495
		goto L83
	case 3:
		v652 = v491
		v653 = v495
		goto L84
	case 4:
		v648 = v491
		v649 = v495
		goto L85
	case 5:
		v642 = v491
		v643 = v495
		goto L86
	case 6:
		v636 = v491
		v637 = v495
		goto L87
	case 7:
		v631 = v495
		goto L88
	case 8:
		v626 = v495
		goto L89
	case 9:
		v621 = v495
		goto L90
	case 10:
		goto L91
	default:
		v680 = v490
		v681 = v491
		v682 = v495
		goto L80
	}
L98:
	;
	v460 = *(*int32)(unsafe.Add(mBase, uint32(v453)+4))
	v461 = v460 + v457
	v462 = *(*int32)(unsafe.Add(mBase, uint32(v453)))
	v464 = *(*int32)(unsafe.Add(mBase, uint32(v453)+8))
	v465 = v464 + v458
	v467 = int32(4)
	v469 = v462 + v456 - v465 ^ base.I32_rotl(v465, v467)
	v473 = v461 - v469 ^ base.I32_rotl(v469, int32(6))
	v474 = v465 + v461
	v475 = v469 + v474
	v476 = v473 + v475
	v480 = v474 - v473 ^ base.I32_rotl(v473, int32(8))
	v484 = v475 - v480 ^ base.I32_rotl(v480, int32(16))
	v488 = v476 - v484 ^ base.I32_rotl(v484, int32(19))
	v489 = v480 + v476
	v490 = v484 + v489
	v491 = v488 + v490
	v495 = v489 - v488 ^ base.I32_rotl(v488, v467)
	v496 = int32(12)
	v497 = v453 + v496
	v499 = v454 - v496
	if base.Ui32(int32(11)) < base.Ui32(v499) {
		v453 = v497
		v454 = v499
		v456 = v490
		v457 = v491
		v458 = v495
		goto L98
	} else {
		goto L100
	}
L99:
	;
	goto L97
L100:
	;
	goto L99
L101:
	;
	v513 = v386
	v514 = v400
	v516 = v407
	v517 = v407
	v518 = v407
	goto L104
L103:
	;
	switch v559 - int32(1) {
	case 0:
		v614 = v550
		goto L107
	case 1:
		v609 = v550
		goto L108
	case 2:
		goto L109
	case 3:
		v602 = v551
		goto L110
	case 4:
		v599 = v551
		goto L111
	case 5:
		v594 = v551
		goto L112
	case 6:
		goto L113
	case 7:
		v585 = v555
		goto L114
	case 8:
		v580 = v555
		goto L115
	case 9:
		v575 = v555
		goto L116
	case 10:
		goto L117
	default:
		v680 = v550
		v681 = v551
		v682 = v555
		goto L80
	}
L104:
	;
	v520 = *(*int32)(unsafe.Add(mBase, uint32(v513)+4))
	v521 = v520 + v517
	v522 = *(*int32)(unsafe.Add(mBase, uint32(v513)))
	v524 = *(*int32)(unsafe.Add(mBase, uint32(v513)+8))
	v525 = v524 + v518
	v527 = int32(4)
	v529 = v522 + v516 - v525 ^ base.I32_rotl(v525, v527)
	v533 = v521 - v529 ^ base.I32_rotl(v529, int32(6))
	v534 = v525 + v521
	v535 = v529 + v534
	v536 = v533 + v535
	v540 = v534 - v533 ^ base.I32_rotl(v533, int32(8))
	v544 = v535 - v540 ^ base.I32_rotl(v540, int32(16))
	v548 = v536 - v544 ^ base.I32_rotl(v544, int32(19))
	v549 = v540 + v536
	v550 = v544 + v549
	v551 = v548 + v550
	v555 = v549 - v548 ^ base.I32_rotl(v548, v527)
	v556 = int32(12)
	v557 = v513 + v556
	v559 = v514 - v556
	if base.Ui32(int32(11)) < base.Ui32(v559) {
		v513 = v557
		v514 = v559
		v516 = v550
		v517 = v551
		v518 = v555
		goto L104
	} else {
		goto L106
	}
L105:
	;
	goto L103
L106:
	;
	goto L105
L107:
	;
	v615 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v557))))
	v680 = v614 + v615
	v681 = v551
	v682 = v555
	goto L80
L108:
	;
	v610 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v557)+1)))
	v614 = v610<<(uint(int32(8))%32) + v609
	goto L107
L109:
	;
	v605 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v557)+2)))
	v609 = v605<<(uint(int32(16))%32) + v550
	goto L108
L110:
	;
	v603 = *(*int32)(unsafe.Add(mBase, uint32(v557)))
	v680 = v603 + v550
	v681 = v602
	v682 = v555
	goto L80
L111:
	;
	v600 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v557)+4)))
	v602 = v599 + v600
	goto L110
L112:
	;
	v595 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v557)+5)))
	v599 = v595<<(uint(int32(8))%32) + v594
	goto L111
L113:
	;
	v590 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v557)+6)))
	v594 = v590<<(uint(int32(16))%32) + v551
	goto L112
L114:
	;
	v586 = *(*int32)(unsafe.Add(mBase, uint32(v557)))
	v588 = *(*int32)(unsafe.Add(mBase, uint32(v557)+4))
	v680 = v586 + v550
	v681 = v588 + v551
	v682 = v585
	goto L80
L115:
	;
	v581 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v557)+8)))
	v585 = v581<<(uint(int32(8))%32) + v580
	goto L114
L116:
	;
	v576 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v557)+9)))
	v580 = v576<<(uint(int32(16))%32) + v575
	goto L115
L117:
	;
	v571 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v557)+10)))
	v575 = v571<<(uint(int32(24))%32) + v555
	goto L116
L118:
	;
	v719 = *(*int64)(unsafe.Add(mBase, uint32(v717)))
	*(*int64)(unsafe.Add(mBase, uint32(v386))) = v719
	v722 = int32(8)
	goto L74
L119:
	;
	v727 = v396
	goto L121
L120:
	;
	v727 = v725
	goto L121
L121:
	;
	if v727 != 0 {
		goto L123
	} else {
		goto L124
	}
L122:
	;
	v731 = v722 + v727
	v732 = v396 - v727
	if v732 != 0 {
		v392 = v392 + v727
		v393 = v731
		v396 = v732
		goto L70
	} else {
		goto L126
	}
L123:
	;
	v728 = F__emscripten_memcpy_bulkmem(m, v722+v386, v392, v727)
	mBase = m.M
	goto L125
L124:
	;
	goto L125
L125:
	;
	goto L122
L126:
	;
	goto L71
}
func F_AutoVacWorkerMain(m *base.Module, l0 int32, l1 int32) {
	mBase := m.M
	_ = mBase
	var v3 int32
	_ = v3
	var v10 int32
	_ = v10
	var v12 int32
	_ = v12
	var v18 int32
	_ = v18
	var v19 int32
	_ = v19
	var v20 int32
	_ = v20
	var v21 int32
	_ = v21
	var v22 int32
	_ = v22
	var v27 int32
	_ = v27
	var v30 int32
	_ = v30
	var v33 int32
	_ = v33
	var v37 int32
	_ = v37
	var v50 int32
	_ = v50
	var v51 int32
	_ = v51
	var v55 int32
	_ = v55
	var v57 int32
	_ = v57
	var v59 int32
	_ = v59
	var v69 int32
	_ = v69
	var v81 int32
	_ = v81
	var v92 int32
	_ = v92
	var v99 int32
	_ = v99
	var v101 int32
	_ = v101
	var v103 int32
	_ = v103
	var v113 int32
	_ = v113
	var v125 int32
	_ = v125
	var v136 int32
	_ = v136
	var v143 int32
	_ = v143
	var v145 int32
	_ = v145
	var v147 int32
	_ = v147
	var v157 int32
	_ = v157
	var v169 int32
	_ = v169
	var v180 int32
	_ = v180
	var v186 int32
	_ = v186
	var v196 int32
	_ = v196
	var v198 int32
	_ = v198
	var v199 int32
	_ = v199
	var v202 int32
	_ = v202
	var v209 int64
	_ = v209
	var v228 int32
	_ = v228
	var v230 int32
	_ = v230
	var v259 int32
	_ = v259
	var v261 int32
	_ = v261
	var v294 int32
	_ = v294
	var v296 int32
	_ = v296
	var v299 int32
	_ = v299
	var v306 int64
	_ = v306
	var v327 int32
	_ = v327
	var v335 int32
	_ = v335
	var v337 int32
	_ = v337
	var v339 int32
	_ = v339
	var v349 int32
	_ = v349
	var v361 int32
	_ = v361
	var v372 int32
	_ = v372
	var v379 int32
	_ = v379
	var v381 int32
	_ = v381
	var v383 int32
	_ = v383
	var v393 int32
	_ = v393
	var v405 int32
	_ = v405
	var v416 int32
	_ = v416
	var v423 int32
	_ = v423
	var v425 int32
	_ = v425
	var v427 int32
	_ = v427
	var v437 int32
	_ = v437
	var v449 int32
	_ = v449
	var v460 int32
	_ = v460
	var v467 int32
	_ = v467
	var v469 int32
	_ = v469
	var v471 int32
	_ = v471
	var v481 int32
	_ = v481
	var v493 int32
	_ = v493
	var v504 int32
	_ = v504
	var v511 int32
	_ = v511
	var v513 int32
	_ = v513
	var v515 int32
	_ = v515
	var v525 int32
	_ = v525
	var v537 int32
	_ = v537
	var v548 int32
	_ = v548
	var v555 int32
	_ = v555
	var v559 int32
	_ = v559
	var v566 int32
	_ = v566
	var v567 int32
	_ = v567
	var v569 int32
	_ = v569
	var v570 int32
	_ = v570
	var v571 int32
	_ = v571
	var v573 int32
	_ = v573
	var v583 int32
	_ = v583
	var v588 int32
	_ = v588
	var v596 int32
	_ = v596
	var v604 int32
	_ = v604
	var v612 int32
	_ = v612
	var v620 int32
	_ = v620
	var v628 int32
	_ = v628
	var v636 int32
	_ = v636
	var v644 int32
	_ = v644
	var v652 int32
	_ = v652
	var v654 int32
	_ = v654
	var v664 int32
	_ = v664
	var v672 int32
	_ = v672
	var v676 int32
	_ = v676
	var v680 int32
	_ = v680
	var v681 int32
	_ = v681
	var v683 int32
	_ = v683
	var v684 int32
	_ = v684
	var v688 int32
	_ = v688
	var v691 int32
	_ = v691
	var v692 int32
	_ = v692
	var v693 int32
	_ = v693
	var v697 int32
	_ = v697
	var v707 int32
	_ = v707
	var v711 int32
	_ = v711
	var v717 int32
	_ = v717
	var v719 int32
	_ = v719
	var v720 int32
	_ = v720
	var v724 int32
	_ = v724
	var v725 int32
	_ = v725
	var v733 int32
	_ = v733
	var v734 int32
	_ = v734
	var v735 int32
	_ = v735
	var v739 int32
	_ = v739
	var v740 int32
	_ = v740
	var v741 int32
	_ = v741
	var v744 int64
	_ = v744
	var v745 int64
	_ = v745
	var v756 int32
	_ = v756
	var v759 int32
	_ = v759
	var v764 int32
	_ = v764
	var v770 int32
	_ = v770
	var v777 int32
	_ = v777
	var v778 int32
	_ = v778
	var v784 int32
	_ = v784
	var v791 int32
	_ = v791
	var v793 int32
	_ = v793
	var v799 int32
	_ = v799
	var v803 int64
	_ = v803
	var v804 int32
	_ = v804
	var v809 int32
	_ = v809
	var v810 int32
	_ = v810
	var v815 int32
	_ = v815
	var v820 int32
	_ = v820
	var v821 int32
	_ = v821
	var v827 int32
	_ = v827
	var v834 int32
	_ = v834
	var v838 int32
	_ = v838
	var v842 int32
	_ = v842
	var v851 int32
	_ = v851
	var v858 int32
	_ = v858
	var v860 int32
	_ = v860
	var v861 int64
	_ = v861
	var v865 int32
	_ = v865
	var v867 int32
	_ = v867
	var v868 int32
	_ = v868
	var v872 int32
	_ = v872
	var v874 int32
	_ = v874
	var v875 int32
	_ = v875
	var v876 int32
	_ = v876
	var v877 int32
	_ = v877
	var v879 int32
	_ = v879
	v3 = int32(0)
	v10 = m.G0
	v12 = v10 - int32(16)
	m.G0 = v12
	v18 = v3
	v19 = v3
	v20 = v3
	v21 = int32(-1)
	v22 = v12
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
	if v21 != int32(1) {
		goto L7
	} else {
		goto L8
	}
L5:
	;
	goto L3
L6:
	;
	v860 = int32(m.ExcTag)
	v861 = int64(m.ExcVals[0])
	m.ExcPending = 0
	if v860 == int32(0) {
		goto L192
	} else {
		goto L193
	}
L7:
	;
	v27 = v22 - int32(160)
	m.G0 = v27
	v30 = v27 + int32(-64)
	m.G0 = v30
	v33 = *(*int32)(unsafe.Add(mBase, _c_F_AutoVacWorkerMain[0]))
	if v33 != 0 {
		goto L10
	} else {
		goto L11
	}
L8:
	;
	v566 = v18
	v567 = v19
	v569 = v22
	v570 = v20
	goto L9
L9:
	;
	if v570 != 0 {
		goto L134
	} else {
		goto L135
	}
L10:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v12)+12)) = v27
	*(*int32)(unsafe.Add(mBase, uint32(v12)+8)) = v30
	F_MemoryContextDelete(m, v33)
	mBase = m.M
	v37 = m.ExcPending
	if v37 != 0 {
		v858 = v30
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
	*(*int32)(unsafe.Add(mBase, uint32(v12)+8)) = v30
	*(*int32)(unsafe.Add(mBase, uint32(v12)+12)) = v27
	goto L15
L13:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_AutoVacWorkerMain[0])) = int32(0)
	goto L12
L14:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v12)+12)) = v27
	*(*int32)(unsafe.Add(mBase, uint32(v12)+8)) = v30
	v55 = int32(914)
	v57 = m.G0
	v59 = v57 - int32(144)
	m.G0 = v59
	switch int32(916) {
	case 0, 2:
		v69 = v55
		goto L19
	default:
		goto L20
	}
L15:
	;
	v50 = *(*int32)(unsafe.Add(mBase, _c_F_AutoVacWorkerMain[1]))
	v51 = F_GetBackendTypeDesc(m, v50)
	mBase = m.M
	goto L17
L17:
	;
	goto L14
L18:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v12)+12)) = v27
	*(*int32)(unsafe.Add(mBase, uint32(v12)+8)) = v30
	v99 = int32(915)
	v101 = m.G0
	v103 = v101 - int32(144)
	m.G0 = v103
	switch int32(917) {
	case 0, 2:
		v113 = v99
		goto L32
	default:
		goto L33
	}
L19:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v59)+4)) = v69
	F_sigemptyset(m, v59+int32(8))
	mBase = m.M
	goto L22
L20:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_AutoVacWorkerMain[2])) = v55
	v69 = int32(_a_F_AutoVacWorkerMain_0)
	goto L19
L22:
	;
	goto L23
L23:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v59)+136)) = int32(268435456)
	v81 = v59 + int32(4)
	goto L26
L24:
	;
	m.G0 = v59 + int32(144)
	goto L18
L26:
	;
	goto L27
L27:
	;
	if v81 != 0 {
		goto L28
	} else {
		goto L29
	}
L28:
	;
	v92 = F___memcpy(m, int32(_a_F_AutoVacWorkerMain_1), v81, int32(140))
	mBase = m.M
	goto L30
L29:
	;
	goto L30
L30:
	;
	goto L24
L31:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v12)+12)) = v27
	*(*int32)(unsafe.Add(mBase, uint32(v12)+8)) = v30
	v143 = int32(295)
	v145 = m.G0
	v147 = v145 - int32(144)
	m.G0 = v147
	switch int32(297) {
	case 0, 2:
		v157 = v143
		goto L45
	default:
		goto L46
	}
L32:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v103)+4)) = v113
	F_sigemptyset(m, v103+int32(8))
	mBase = m.M
	goto L35
L33:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_AutoVacWorkerMain[3])) = v99
	v113 = int32(_a_F_AutoVacWorkerMain_0)
	goto L32
L35:
	;
	goto L36
L36:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v103)+136)) = int32(268435456)
	v125 = v103 + int32(4)
	goto L39
L37:
	;
	m.G0 = v103 + int32(144)
	goto L31
L39:
	;
	goto L40
L40:
	;
	if v125 != 0 {
		goto L41
	} else {
		goto L42
	}
L41:
	;
	v136 = F___memcpy(m, int32(_a_F_AutoVacWorkerMain_2), v125, int32(140))
	mBase = m.M
	goto L43
L42:
	;
	goto L43
L43:
	;
	goto L37
L44:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v12)+12)) = v27
	*(*int32)(unsafe.Add(mBase, uint32(v12)+8)) = v30
	v186 = int32(0)
	*(*int32)(unsafe.Add(mBase, _c_F_AutoVacWorkerMain[4])) = v186
	*(*int32)(unsafe.Add(mBase, _c_F_AutoVacWorkerMain[5])) = v186
	v196 = v186
	goto L58
L45:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v147)+4)) = v157
	F_sigemptyset(m, v147+int32(8))
	mBase = m.M
	goto L48
L46:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_AutoVacWorkerMain[6])) = v143
	v157 = int32(_a_F_AutoVacWorkerMain_0)
	goto L45
L48:
	;
	goto L49
L49:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v147)+136)) = int32(268435456)
	v169 = v147 + int32(4)
	goto L52
L50:
	;
	m.G0 = v147 + int32(144)
	goto L44
L52:
	;
	goto L53
L53:
	;
	if v169 != 0 {
		goto L54
	} else {
		goto L55
	}
L54:
	;
	v180 = F___memcpy(m, int32(_a_F_AutoVacWorkerMain_3), v169, int32(140))
	mBase = m.M
	goto L56
L55:
	;
	goto L56
L56:
	;
	goto L50
L57:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v12)+12)) = v27
	*(*int32)(unsafe.Add(mBase, uint32(v12)+8)) = v30
	v335 = int32(-2)
	v337 = m.G0
	v339 = v337 - int32(144)
	m.G0 = v339
	switch int32(0) {
	case 0, 2:
		v349 = v335
		goto L64
	default:
		goto L65
	}
L58:
	;
	v198 = int32(40)
	v199 = v196 * v198
	v202 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v199)+uint32(_c_F_AutoVacWorkerMain[7]))) = uint8(v202)
	*(*int32)(unsafe.Add(mBase, uint32(v199)+uint32(_c_F_AutoVacWorkerMain[8]))) = v196
	v209 = int64(0)
	*(*int64)(unsafe.Add(mBase, uint32(v199)+uint32(_c_F_AutoVacWorkerMain[9]))) = v209
	*(*int32)(unsafe.Add(mBase, uint32(v199)+uint32(_c_F_AutoVacWorkerMain[10]))) = v202
	*(*int64)(unsafe.Add(mBase, uint32(v199)+uint32(_c_F_AutoVacWorkerMain[11]))) = v209
	*(*int32)(unsafe.Add(mBase, uint32(v199)+uint32(_c_F_AutoVacWorkerMain[12]))) = v202
	*(*uint8)(unsafe.Add(mBase, uint32(v199)+uint32(_c_F_AutoVacWorkerMain[13]))) = uint8(v202)
	v228 = v196 | int32(1)
	v230 = v228 * v198
	*(*uint8)(unsafe.Add(mBase, uint32(v230)+uint32(_c_F_AutoVacWorkerMain[7]))) = uint8(v202)
	*(*int32)(unsafe.Add(mBase, uint32(v230)+uint32(_c_F_AutoVacWorkerMain[8]))) = v228
	*(*int64)(unsafe.Add(mBase, uint32(v230)+uint32(_c_F_AutoVacWorkerMain[9]))) = v209
	*(*int32)(unsafe.Add(mBase, uint32(v230)+uint32(_c_F_AutoVacWorkerMain[10]))) = v202
	*(*int64)(unsafe.Add(mBase, uint32(v230)+uint32(_c_F_AutoVacWorkerMain[11]))) = v209
	*(*int32)(unsafe.Add(mBase, uint32(v230)+uint32(_c_F_AutoVacWorkerMain[12]))) = v202
	*(*uint8)(unsafe.Add(mBase, uint32(v230)+uint32(_c_F_AutoVacWorkerMain[13]))) = uint8(v202)
	v259 = v196 | int32(2)
	v261 = v259 * v198
	*(*uint8)(unsafe.Add(mBase, uint32(v261)+uint32(_c_F_AutoVacWorkerMain[7]))) = uint8(v202)
	*(*int32)(unsafe.Add(mBase, uint32(v261)+uint32(_c_F_AutoVacWorkerMain[8]))) = v259
	*(*int64)(unsafe.Add(mBase, uint32(v261)+uint32(_c_F_AutoVacWorkerMain[9]))) = v209
	*(*int32)(unsafe.Add(mBase, uint32(v261)+uint32(_c_F_AutoVacWorkerMain[10]))) = v202
	*(*int64)(unsafe.Add(mBase, uint32(v261)+uint32(_c_F_AutoVacWorkerMain[11]))) = v209
	*(*int32)(unsafe.Add(mBase, uint32(v261)+uint32(_c_F_AutoVacWorkerMain[12]))) = v202
	*(*uint8)(unsafe.Add(mBase, uint32(v261)+uint32(_c_F_AutoVacWorkerMain[13]))) = uint8(v202)
	if base.B2i32(v196 == int32(20)) == v202 {
		goto L60
	} else {
		goto L61
	}
L59:
	;
	v327 = int32(1)
	*(*uint8)(unsafe.Add(mBase, _c_F_AutoVacWorkerMain[14])) = uint8(v327)
	F_pqsignal_be(m, int32(14), int32(1785))
	mBase = m.M
	goto L57
L60:
	;
	v294 = v196 | int32(3)
	v296 = v294 * int32(40)
	v299 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v296)+uint32(_c_F_AutoVacWorkerMain[7]))) = uint8(v299)
	*(*int32)(unsafe.Add(mBase, uint32(v296)+uint32(_c_F_AutoVacWorkerMain[8]))) = v294
	v306 = int64(0)
	*(*int64)(unsafe.Add(mBase, uint32(v296)+uint32(_c_F_AutoVacWorkerMain[9]))) = v306
	*(*int32)(unsafe.Add(mBase, uint32(v296)+uint32(_c_F_AutoVacWorkerMain[10]))) = v299
	*(*int64)(unsafe.Add(mBase, uint32(v296)+uint32(_c_F_AutoVacWorkerMain[11]))) = v306
	*(*int32)(unsafe.Add(mBase, uint32(v296)+uint32(_c_F_AutoVacWorkerMain[12]))) = v299
	*(*uint8)(unsafe.Add(mBase, uint32(v296)+uint32(_c_F_AutoVacWorkerMain[13]))) = uint8(v299)
	v196 = v196 + int32(4)
	goto L58
L61:
	;
	goto L62
L62:
	;
	goto L59
L63:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v12)+12)) = v27
	*(*int32)(unsafe.Add(mBase, uint32(v12)+8)) = v30
	v379 = int32(917)
	v381 = m.G0
	v383 = v381 - int32(144)
	m.G0 = v383
	switch int32(919) {
	case 0, 2:
		v393 = v379
		goto L77
	default:
		goto L78
	}
L64:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v339)+4)) = v349
	F_sigemptyset(m, v339+int32(8))
	mBase = m.M
	goto L67
L65:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_AutoVacWorkerMain[15])) = v335
	v349 = int32(_a_F_AutoVacWorkerMain_0)
	goto L64
L67:
	;
	goto L68
L68:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v339)+136)) = int32(268435456)
	v361 = v339 + int32(4)
	goto L71
L69:
	;
	m.G0 = v339 + int32(144)
	goto L63
L71:
	;
	goto L72
L72:
	;
	if v361 != 0 {
		goto L73
	} else {
		goto L74
	}
L73:
	;
	v372 = F___memcpy(m, int32(_a_F_AutoVacWorkerMain_4), v361, int32(140))
	mBase = m.M
	goto L75
L74:
	;
	goto L75
L75:
	;
	goto L69
L76:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v12)+12)) = v27
	*(*int32)(unsafe.Add(mBase, uint32(v12)+8)) = v30
	v423 = int32(-2)
	v425 = m.G0
	v427 = v425 - int32(144)
	m.G0 = v427
	switch int32(0) {
	case 0, 2:
		v437 = v423
		goto L90
	default:
		goto L91
	}
L77:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v383)+4)) = v393
	F_sigemptyset(m, v383+int32(8))
	mBase = m.M
	goto L80
L78:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_AutoVacWorkerMain[16])) = v379
	v393 = int32(_a_F_AutoVacWorkerMain_0)
	goto L77
L80:
	;
	goto L81
L81:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v383)+136)) = int32(268435456)
	v405 = v383 + int32(4)
	goto L84
L82:
	;
	m.G0 = v383 + int32(144)
	goto L76
L84:
	;
	goto L85
L85:
	;
	if v405 != 0 {
		goto L86
	} else {
		goto L87
	}
L86:
	;
	v416 = F___memcpy(m, int32(_a_F_AutoVacWorkerMain_5), v405, int32(140))
	mBase = m.M
	goto L88
L87:
	;
	goto L88
L88:
	;
	goto L82
L89:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v12)+12)) = v27
	*(*int32)(unsafe.Add(mBase, uint32(v12)+8)) = v30
	v467 = int32(919)
	v469 = m.G0
	v471 = v469 - int32(144)
	m.G0 = v471
	switch int32(921) {
	case 0, 2:
		v481 = v467
		goto L103
	default:
		goto L104
	}
L90:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v427)+4)) = v437
	F_sigemptyset(m, v427+int32(8))
	mBase = m.M
	goto L93
L91:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_AutoVacWorkerMain[17])) = v423
	v437 = int32(_a_F_AutoVacWorkerMain_0)
	goto L90
L93:
	;
	goto L94
L94:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v427)+136)) = int32(268435456)
	v449 = v427 + int32(4)
	goto L97
L95:
	;
	m.G0 = v427 + int32(144)
	goto L89
L97:
	;
	goto L98
L98:
	;
	if v449 != 0 {
		goto L99
	} else {
		goto L100
	}
L99:
	;
	v460 = F___memcpy(m, int32(_a_F_AutoVacWorkerMain_6), v449, int32(140))
	mBase = m.M
	goto L101
L100:
	;
	goto L101
L101:
	;
	goto L95
L102:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v12)+12)) = v27
	*(*int32)(unsafe.Add(mBase, uint32(v12)+8)) = v30
	v511 = int32(0)
	v513 = m.G0
	v515 = v513 - int32(144)
	m.G0 = v515
	switch int32(2) {
	case 0, 2:
		v525 = v511
		goto L116
	default:
		goto L117
	}
L103:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v471)+4)) = v481
	F_sigemptyset(m, v471+int32(8))
	mBase = m.M
	goto L106
L104:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_AutoVacWorkerMain[18])) = v467
	v481 = int32(_a_F_AutoVacWorkerMain_0)
	goto L103
L106:
	;
	goto L107
L107:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v471)+136)) = int32(268435456)
	v493 = v471 + int32(4)
	goto L110
L108:
	;
	m.G0 = v471 + int32(144)
	goto L102
L110:
	;
	goto L111
L111:
	;
	if v493 != 0 {
		goto L112
	} else {
		goto L113
	}
L112:
	;
	v504 = F___memcpy(m, int32(_a_F_AutoVacWorkerMain_7), v493, int32(140))
	mBase = m.M
	goto L114
L113:
	;
	goto L114
L114:
	;
	goto L108
L115:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v12)+12)) = v27
	*(*int32)(unsafe.Add(mBase, uint32(v12)+8)) = v30
	F_InitProcess(m)
	mBase = m.M
	v555 = m.ExcPending
	if v555 != 0 {
		v858 = v30
		goto L6
	} else {
		goto L128
	}
L116:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v515)+4)) = v525
	F_sigemptyset(m, v515+int32(8))
	mBase = m.M
	goto L118
L117:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_AutoVacWorkerMain[19])) = v511
	v525 = int32(_a_F_AutoVacWorkerMain_0)
	goto L116
L118:
	;
	goto L120
L120:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v515)+136)) = int32(268435457)
	v537 = v515 + int32(4)
	goto L123
L121:
	;
	m.G0 = v515 + int32(144)
	goto L115
L123:
	;
	goto L124
L124:
	;
	if v537 != 0 {
		goto L125
	} else {
		goto L126
	}
L125:
	;
	v548 = F___memcpy(m, int32(_a_F_AutoVacWorkerMain_8), v537, int32(140))
	mBase = m.M
	goto L127
L126:
	;
	goto L127
L127:
	;
	goto L121
L128:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v12)+12)) = v27
	*(*int32)(unsafe.Add(mBase, uint32(v12)+8)) = v30
	F_BaseInit(m)
	mBase = m.M
	v559 = m.ExcPending
	if v559 != 0 {
		v858 = v30
		goto L6
	} else {
		goto L129
	}
L129:
	;
	goto L130
L130:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v27)+4)) = int32(1)
	*(*int32)(unsafe.Add(mBase, uint32(v27))) = v12 + int32(4)
	goto L133
L131:
	;
	v566 = v27
	v567 = v30
	v569 = v30
	v570 = int32(0)
	goto L9
L133:
	;
	goto L131
L134:
	;
	v571 = int32(_a_F_AutoVacWorkerMain_9)
	v573 = *(*int32)(unsafe.Add(mBase, _c_F_AutoVacWorkerMain[20]))
	*(*int32)(unsafe.Add(mBase, _c_F_AutoVacWorkerMain[20])) = v573 + int32(1)
	*(*int32)(unsafe.Add(mBase, _c_F_AutoVacWorkerMain[21])) = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v12)+8)) = v567
	*(*int32)(unsafe.Add(mBase, uint32(v12)+12)) = v566
	F_EmitErrorReport(m)
	mBase = m.M
	v583 = m.ExcPending
	if v583 != 0 {
		v858 = v569
		goto L6
	} else {
		goto L137
	}
L135:
	;
	goto L136
L136:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_AutoVacWorkerMain[22])) = v566
	*(*int32)(unsafe.Add(mBase, uint32(v12)+8)) = v567
	*(*int32)(unsafe.Add(mBase, uint32(v12)+12)) = v566
	F_sigprocmask(m, int32(_a_F_AutoVacWorkerMain_10), int32(0))
	mBase = m.M
	v596 = m.ExcPending
	if v596 != 0 {
		v858 = v569
		goto L6
	} else {
		goto L139
	}
L137:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v12)+12)) = v566
	*(*int32)(unsafe.Add(mBase, uint32(v12)+8)) = v567
	F_proc_exit(m, int32(0))
	mBase = m.M
	v588 = m.ExcPending
	if v588 != 0 {
		v858 = v569
		goto L6
	} else {
		goto L138
	}
L138:
	;
	goto L3
L139:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v12)+12)) = v566
	*(*int32)(unsafe.Add(mBase, uint32(v12)+8)) = v567
	F_SetConfigOption(m, int32(_a_F_AutoVacWorkerMain_11), int32(_a_F_AutoVacWorkerMain_12), int32(5), int32(10))
	mBase = m.M
	v604 = m.ExcPending
	if v604 != 0 {
		v858 = v569
		goto L6
	} else {
		goto L140
	}
L140:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v12)+12)) = v566
	*(*int32)(unsafe.Add(mBase, uint32(v12)+8)) = v567
	F_SetConfigOption(m, int32(_a_F_AutoVacWorkerMain_13), int32(_a_F_AutoVacWorkerMain_14), int32(5), int32(10))
	mBase = m.M
	v612 = m.ExcPending
	if v612 != 0 {
		v858 = v569
		goto L6
	} else {
		goto L141
	}
L141:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v12)+12)) = v566
	*(*int32)(unsafe.Add(mBase, uint32(v12)+8)) = v567
	F_SetConfigOption(m, int32(_a_F_AutoVacWorkerMain_15), int32(_a_F_AutoVacWorkerMain_16), int32(5), int32(10))
	mBase = m.M
	v620 = m.ExcPending
	if v620 != 0 {
		v858 = v569
		goto L6
	} else {
		goto L142
	}
L142:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v12)+12)) = v566
	*(*int32)(unsafe.Add(mBase, uint32(v12)+8)) = v567
	F_SetConfigOption(m, int32(_a_F_AutoVacWorkerMain_17), int32(_a_F_AutoVacWorkerMain_16), int32(5), int32(10))
	mBase = m.M
	v628 = m.ExcPending
	if v628 != 0 {
		v858 = v569
		goto L6
	} else {
		goto L143
	}
L143:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v12)+12)) = v566
	*(*int32)(unsafe.Add(mBase, uint32(v12)+8)) = v567
	F_SetConfigOption(m, int32(_a_F_AutoVacWorkerMain_18), int32(_a_F_AutoVacWorkerMain_16), int32(5), int32(10))
	mBase = m.M
	v636 = m.ExcPending
	if v636 != 0 {
		v858 = v569
		goto L6
	} else {
		goto L144
	}
L144:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v12)+12)) = v566
	*(*int32)(unsafe.Add(mBase, uint32(v12)+8)) = v567
	F_SetConfigOption(m, int32(_a_F_AutoVacWorkerMain_19), int32(_a_F_AutoVacWorkerMain_16), int32(5), int32(10))
	mBase = m.M
	v644 = m.ExcPending
	if v644 != 0 {
		v858 = v569
		goto L6
	} else {
		goto L145
	}
L145:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v12)+12)) = v566
	*(*int32)(unsafe.Add(mBase, uint32(v12)+8)) = v567
	F_SetConfigOption(m, int32(_a_F_AutoVacWorkerMain_20), int32(_a_F_AutoVacWorkerMain_21), int32(5), int32(10))
	mBase = m.M
	v652 = m.ExcPending
	if v652 != 0 {
		v858 = v569
		goto L6
	} else {
		goto L146
	}
L146:
	;
	v654 = *(*int32)(unsafe.Add(mBase, _c_F_AutoVacWorkerMain[23]))
	if int32(2) <= v654 {
		goto L147
	} else {
		goto L148
	}
L147:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v12)+12)) = v566
	*(*int32)(unsafe.Add(mBase, uint32(v12)+8)) = v567
	F_SetConfigOption(m, int32(_a_F_AutoVacWorkerMain_22), int32(_a_F_AutoVacWorkerMain_23), int32(5), int32(10))
	mBase = m.M
	v664 = m.ExcPending
	if v664 != 0 {
		v858 = v569
		goto L6
	} else {
		goto L150
	}
L148:
	;
	goto L149
L149:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v12)+12)) = v566
	*(*int32)(unsafe.Add(mBase, uint32(v12)+8)) = v567
	F_SetConfigOption(m, int32(_a_F_AutoVacWorkerMain_24), int32(_a_F_AutoVacWorkerMain_25), int32(5), int32(10))
	mBase = m.M
	v672 = m.ExcPending
	if v672 != 0 {
		v858 = v569
		goto L6
	} else {
		goto L151
	}
L150:
	;
	goto L149
L151:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v12)+8)) = v567
	*(*int32)(unsafe.Add(mBase, uint32(v12)+12)) = v566
	v676 = *(*int32)(unsafe.Add(mBase, _c_F_AutoVacWorkerMain[24]))
	v680 = F_LWLockAcquire(m, v676+int32(2816), int32(0))
	mBase = m.M
	v681 = m.ExcPending
	if v681 != 0 {
		v858 = v569
		goto L6
	} else {
		goto L152
	}
L152:
	;
	v683 = *(*int32)(unsafe.Add(mBase, _c_F_AutoVacWorkerMain[25]))
	v684 = *(*int32)(unsafe.Add(mBase, uint32(v683)+32))
	if v684 != 0 {
		goto L154
	} else {
		goto L155
	}
L153:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v12)+12)) = v566
	*(*int32)(unsafe.Add(mBase, uint32(v12)+8)) = v567
	F_proc_exit(m, int32(0))
	mBase = m.M
	v851 = m.ExcPending
	if v851 != 0 {
		v858 = v569
		goto L6
	} else {
		goto L191
	}
L154:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_AutoVacWorkerMain[26])) = v684
	v688 = *(*int32)(unsafe.Add(mBase, _c_F_AutoVacWorkerMain[27]))
	*(*int32)(unsafe.Add(mBase, uint32(v684)+16)) = v688
	v691 = v683 + int32(24)
	v692 = *(*int32)(unsafe.Add(mBase, uint32(v684)+8))
	v693 = *(*int32)(unsafe.Add(mBase, uint32(v683)+28))
	if v693 == int32(0) {
		goto L157
	} else {
		goto L158
	}
L155:
	;
	goto L156
L156:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v12)+12)) = v566
	*(*int32)(unsafe.Add(mBase, uint32(v12)+8)) = v567
	v820 = F_errstart(m, int32(19), int32(0))
	mBase = m.M
	v821 = m.ExcPending
	if v821 != 0 {
		v858 = v569
		goto L6
	} else {
		goto L184
	}
L157:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v691))) = v691
	v697 = v691
	goto L159
L158:
	;
	v697 = v693
	goto L159
L159:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v684))) = v691
	*(*int32)(unsafe.Add(mBase, uint32(v684)+4)) = v697
	*(*int32)(unsafe.Add(mBase, uint32(v697))) = v684
	*(*int32)(unsafe.Add(mBase, uint32(v683)+32)) = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v683)+28)) = v684
	*(*int32)(unsafe.Add(mBase, uint32(v12)+8)) = v567
	*(*int32)(unsafe.Add(mBase, uint32(v12)+12)) = v566
	v707 = *(*int32)(unsafe.Add(mBase, _c_F_AutoVacWorkerMain[24]))
	F_LWLockRelease(m, v707+int32(2816))
	mBase = m.M
	v711 = m.ExcPending
	if v711 != 0 {
		v858 = v569
		goto L6
	} else {
		goto L160
	}
L160:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v12)+12)) = v566
	*(*int32)(unsafe.Add(mBase, uint32(v12)+8)) = v567
	F_on_shmem_exit(m, int32(921), int32(0))
	mBase = m.M
	v717 = m.ExcPending
	if v717 != 0 {
		v858 = v569
		goto L6
	} else {
		goto L161
	}
L161:
	;
	v719 = *(*int32)(unsafe.Add(mBase, _c_F_AutoVacWorkerMain[25]))
	v720 = *(*int32)(unsafe.Add(mBase, uint32(v719)+8))
	if v720 != 0 {
		goto L162
	} else {
		goto L163
	}
L162:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v12)+12)) = v566
	*(*int32)(unsafe.Add(mBase, uint32(v12)+8)) = v567
	v724 = F_kill(m, v720, int32(12))
	mBase = m.M
	v725 = m.ExcPending
	if v725 != 0 {
		v858 = v569
		goto L6
	} else {
		goto L165
	}
L163:
	;
	goto L164
L164:
	;
	if v692 == int32(0) {
		goto L153
	} else {
		goto L166
	}
L165:
	;
	goto L164
L166:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v12)+12)) = v566
	*(*int32)(unsafe.Add(mBase, uint32(v12)+8)) = v567
	v733 = F_pgstat_get_entry_ref_locked(m, int32(1), v692, int64(0), int32(0))
	mBase = m.M
	v734 = m.ExcPending
	if v734 != 0 {
		v858 = v569
		goto L6
	} else {
		goto L167
	}
L167:
	;
	v735 = *(*int32)(unsafe.Add(mBase, uint32(v733)+4))
	v739 = m.G0
	v740 = int32(16)
	v741 = v739 - v740
	m.G0 = v741
	F___gettimeofday(m, v741)
	mBase = m.M
	v744 = *(*int64)(unsafe.Add(mBase, uint32(v741)))
	v745 = int64(*(*int32)(unsafe.Add(mBase, uint32(v741)+8)))
	m.G0 = v741 + v740
	goto L168
L168:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v735)+96)) = v745 + v744*int64(1000000) - int64(946684800000000)
	F_pgstat_unlock_entry(m, v733)
	mBase = m.M
	v756 = m.ExcPending
	if v756 != 0 {
		v858 = v569
		goto L6
	} else {
		goto L169
	}
L169:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v12)+12)) = v566
	*(*int32)(unsafe.Add(mBase, uint32(v12)+8)) = v567
	v759 = int32(0)
	F_InitPostgres(m, v759, v692, v759, v759, int32(2), v567)
	mBase = m.M
	v764 = m.ExcPending
	if v764 != 0 {
		v858 = v569
		goto L6
	} else {
		goto L170
	}
L170:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_AutoVacWorkerMain[28])) = int32(2)
	*(*int32)(unsafe.Add(mBase, uint32(v12)+8)) = v567
	*(*int32)(unsafe.Add(mBase, uint32(v12)+12)) = v566
	v770 = F_strlen(m, v567)
	mBase = m.M
	*(*int32)(unsafe.Add(mBase, uint32(v12)+12)) = v566
	*(*int32)(unsafe.Add(mBase, uint32(v12)+8)) = v567
	*(*int32)(unsafe.Add(mBase, uint32(v12)+12)) = v566
	*(*int32)(unsafe.Add(mBase, uint32(v12)+8)) = v567
	v777 = F_errstart(m, int32(14), int32(0))
	mBase = m.M
	v778 = m.ExcPending
	if v778 != 0 {
		v858 = v569
		goto L6
	} else {
		goto L171
	}
L171:
	;
	if v777 != 0 {
		goto L172
	} else {
		goto L173
	}
L172:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v12)+12)) = v566
	*(*int32)(unsafe.Add(mBase, uint32(v12)+8)) = v567
	*(*int32)(unsafe.Add(mBase, uint32(v12))) = v567
	F_errmsg_internal(m, int32(_a_F_AutoVacWorkerMain_26), v12)
	mBase = m.M
	v784 = m.ExcPending
	if v784 != 0 {
		v858 = v569
		goto L6
	} else {
		goto L175
	}
L173:
	;
	goto L174
L174:
	;
	v793 = *(*int32)(unsafe.Add(mBase, _c_F_AutoVacWorkerMain[29]))
	if v793 != 0 {
		goto L177
	} else {
		goto L178
	}
L175:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v12)+12)) = v566
	*(*int32)(unsafe.Add(mBase, uint32(v12)+8)) = v567
	F_errfinish(m, int32(_a_F_AutoVacWorkerMain_27), int32(1582), int32(_a_F_AutoVacWorkerMain_28))
	mBase = m.M
	v791 = m.ExcPending
	if v791 != 0 {
		v858 = v569
		goto L6
	} else {
		goto L176
	}
L176:
	;
	goto L174
L177:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v12)+12)) = v566
	*(*int32)(unsafe.Add(mBase, uint32(v12)+8)) = v567
	F_pg_usleep(m, v793*int32(_a_F_AutoVacWorkerMain_29))
	mBase = m.M
	v799 = m.ExcPending
	if v799 != 0 {
		v858 = v569
		goto L6
	} else {
		goto L180
	}
L178:
	;
	goto L179
L179:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v12)+12)) = v566
	*(*int32)(unsafe.Add(mBase, uint32(v12)+8)) = v567
	v803 = F_ReadNextFullTransactionId(m)
	mBase = m.M
	v804 = m.ExcPending
	if v804 != 0 {
		v858 = v569
		goto L6
	} else {
		goto L181
	}
L180:
	;
	goto L179
L181:
	;
	*(*uint32)(unsafe.Add(mBase, _c_F_AutoVacWorkerMain[30])) = uint32(v803)
	*(*int32)(unsafe.Add(mBase, uint32(v12)+8)) = v567
	*(*int32)(unsafe.Add(mBase, uint32(v12)+12)) = v566
	v809 = F_ReadNextMultiXactId(m)
	mBase = m.M
	v810 = m.ExcPending
	if v810 != 0 {
		v858 = v569
		goto L6
	} else {
		goto L182
	}
L182:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_AutoVacWorkerMain[31])) = v809
	*(*int32)(unsafe.Add(mBase, uint32(v12)+8)) = v567
	*(*int32)(unsafe.Add(mBase, uint32(v12)+12)) = v566
	F_do_autovacuum(m)
	mBase = m.M
	v815 = m.ExcPending
	if v815 != 0 {
		v858 = v569
		goto L6
	} else {
		goto L183
	}
L183:
	;
	goto L153
L184:
	;
	if v820 != 0 {
		goto L185
	} else {
		goto L186
	}
L185:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v12)+12)) = v566
	*(*int32)(unsafe.Add(mBase, uint32(v12)+8)) = v567
	F_errmsg_internal(m, int32(_a_F_AutoVacWorkerMain_30), int32(0))
	mBase = m.M
	v827 = m.ExcPending
	if v827 != 0 {
		v858 = v569
		goto L6
	} else {
		goto L188
	}
L186:
	;
	goto L187
L187:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v12)+8)) = v567
	*(*int32)(unsafe.Add(mBase, uint32(v12)+12)) = v566
	v838 = *(*int32)(unsafe.Add(mBase, _c_F_AutoVacWorkerMain[24]))
	F_LWLockRelease(m, v838+int32(2816))
	mBase = m.M
	v842 = m.ExcPending
	if v842 != 0 {
		v858 = v569
		goto L6
	} else {
		goto L190
	}
L188:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v12)+12)) = v566
	*(*int32)(unsafe.Add(mBase, uint32(v12)+8)) = v567
	F_errfinish(m, int32(_a_F_AutoVacWorkerMain_27), int32(1549), int32(_a_F_AutoVacWorkerMain_28))
	mBase = m.M
	v834 = m.ExcPending
	if v834 != 0 {
		v858 = v569
		goto L6
	} else {
		goto L189
	}
L189:
	;
	goto L187
L190:
	;
	goto L153
L191:
	;
	goto L5
L192:
	;
	v865 = int32(v861)
	m.G0 = v858
	v867 = *(*int32)(unsafe.Add(mBase, uint32(v865)+4))
	v868 = *(*int32)(unsafe.Add(mBase, uint32(v865)))
	v872 = *(*int32)(unsafe.Add(mBase, uint32(v868)))
	if v12+int32(4) == v872 {
		goto L195
	} else {
		goto L196
	}
L193:
	;
	m.ExcPending = 1
	goto L201
L194:
	;
	if v875 != 0 {
		goto L198
	} else {
		goto L199
	}
L195:
	;
	v874 = *(*int32)(unsafe.Add(mBase, uint32(v868)+4))
	v875 = v874
	goto L197
L196:
	;
	v875 = int32(0)
	goto L197
L197:
	;
	goto L194
L198:
	;
	v876 = *(*int32)(unsafe.Add(mBase, uint32(v12)+12))
	v877 = *(*int32)(unsafe.Add(mBase, uint32(v12)+8))
	v18 = v876
	v19 = v877
	v20 = v867
	v21 = v875
	v22 = v858
	goto L1
L199:
	;
	goto L200
L200:
	;
	F___wasm_longjmp(m, v868, v867)
	mBase = m.M
	v879 = m.ExcPending
	if v879 != 0 {
		goto L201
	} else {
		goto L202
	}
L201:
	;
	return
L202:
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
	v5 = F_DirectFunctionCall2Coll(m, int32(2460), int32(0), l0, int32(_a_F_abs_interval_0))
	v8 = m.ExcPending
	if v8 != 0 {
		return int32(0)
	} else {
		if v5 != 0 {
			v11 = F_DirectFunctionCall1Coll(m, int32(2464), int32(0), l0)
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
	var v17 int32
	_ = v17
	var v19 int32
	_ = v19
	var v21 int32
	_ = v21
	var v25 int32
	_ = v25
	var v26 int32
	_ = v26
	var v27 int32
	_ = v27
	var v29 int32
	_ = v29
	var v31 int32
	_ = v31
	var v34 int64
	_ = v34
	var v37 int64
	_ = v37
	var v42 int32
	_ = v42
	var v44 int32
	_ = v44
	var v47 int32
	_ = v47
	var v59 int32
	_ = v59
	var v62 int32
	_ = v62
	var v64 int32
	_ = v64
	var v73 int32
	_ = v73
	var v74 int32
	_ = v74
	var v75 int32
	_ = v75
	var v76 int32
	_ = v76
	var v81 int32
	_ = v81
	var v85 int32
	_ = v85
	var v86 int32
	_ = v86
	var v89 int32
	_ = v89
	var v90 int32
	_ = v90
	var v91 int32
	_ = v91
	var v95 int32
	_ = v95
	var v99 int32
	_ = v99
	var v100 int32
	_ = v100
	var v102 int32
	_ = v102
	var v103 int32
	_ = v103
	var v105 int32
	_ = v105
	var v109 int32
	_ = v109
	var v110 int32
	_ = v110
	var v113 int32
	_ = v113
	var v125 int32
	_ = v125
	var v126 int32
	_ = v126
	var v127 int32
	_ = v127
	var v129 int32
	_ = v129
	var v132 int32
	_ = v132
	var v134 int32
	_ = v134
	var v137 int32
	_ = v137
	var v139 int32
	_ = v139
	var v141 int32
	_ = v141
	var v149 int32
	_ = v149
	var v154 int32
	_ = v154
	var v156 int32
	_ = v156
	var v157 int32
	_ = v157
	var v168 int32
	_ = v168
	var v169 int32
	_ = v169
	var v170 int32
	_ = v170
	var v171 int32
	_ = v171
	var v176 int32
	_ = v176
	var v180 int32
	_ = v180
	var v181 int32
	_ = v181
	var v184 int32
	_ = v184
	var v185 int32
	_ = v185
	var v186 int32
	_ = v186
	var v190 int32
	_ = v190
	var v194 int32
	_ = v194
	var v195 int32
	_ = v195
	var v197 int32
	_ = v197
	var v198 int32
	_ = v198
	var v200 int32
	_ = v200
	var v204 int32
	_ = v204
	var v205 int32
	_ = v205
	var v207 int32
	_ = v207
	var v220 int32
	_ = v220
	var v221 int32
	_ = v221
	var v222 int32
	_ = v222
	var v224 int32
	_ = v224
	var v227 int32
	_ = v227
	var v229 int32
	_ = v229
	var v232 int32
	_ = v232
	var v252 int32
	_ = v252
	var v260 int32
	_ = v260
	var v271 int32
	_ = v271
	var v274 int32
	_ = v274
	var v277 int32
	_ = v277
	var v284 int32
	_ = v284
	var v285 int32
	_ = v285
	var v288 int32
	_ = v288
	var v291 int32
	_ = v291
	var v292 int32
	_ = v292
	var v295 int32
	_ = v295
	var v300 int32
	_ = v300
	var v314 int32
	_ = v314
	var v315 int32
	_ = v315
	var v318 int32
	_ = v318
	var v319 int32
	_ = v319
	var v321 int32
	_ = v321
	var v324 int32
	_ = v324
	var v326 int32
	_ = v326
	var v329 int32
	_ = v329
	var v330 int32
	_ = v330
	var v351 int32
	_ = v351
	var v352 int32
	_ = v352
	var v353 int32
	_ = v353
	var v359 int32
	_ = v359
	var v361 int32
	_ = v361
	var v375 int32
	_ = v375
	var v376 int32
	_ = v376
	var v377 int32
	_ = v377
	var v390 int32
	_ = v390
	var v407 int32
	_ = v407
	var v408 int32
	_ = v408
	var v412 int32
	_ = v412
	var v431 int32
	_ = v431
	var v433 int32
	_ = v433
	v3 = int32(0)
	v17 = m.G0
	v19 = v17 - int32(48)
	m.G0 = v19
	v21 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	if v21 == v3 {
		goto L2
	} else {
		goto L3
	}
L1:
	;
	m.G0 = v19 + int32(48)
	return
L2:
	;
	v25 = F_palloc(m, int32(2))
	mBase = m.M
	v26 = m.ExcPending
	if v26 != 0 {
		goto L5
	} else {
		goto L6
	}
L3:
	;
	goto L4
L4:
	;
	v42 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	if v42 != 0 {
		goto L11
	} else {
		goto L12
	}
L5:
	;
	return
L6:
	;
	v27 = int32(0)
	*(*uint16)(unsafe.Add(mBase, uint32(v25))) = uint16(v27)
	v29 = *(*int32)(unsafe.Add(mBase, uint32(l1)+16))
	if v29 != 0 {
		goto L7
	} else {
		goto L8
	}
L7:
	;
	F_pfree(m, v29)
	mBase = m.M
	v31 = m.ExcPending
	if v31 != 0 {
		goto L5
	} else {
		goto L10
	}
L8:
	;
	goto L9
L9:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l1)+16)) = v25
	v34 = *(*int64)(unsafe.Add(mBase, _c_F_accum_sum_final[0]))
	*(*int64)(unsafe.Add(mBase, uint32(l1)+8)) = v34
	v37 = *(*int64)(unsafe.Add(mBase, _c_F_accum_sum_final[1]))
	*(*int64)(unsafe.Add(mBase, uint32(l1))) = v37
	*(*int32)(unsafe.Add(mBase, uint32(l1)+20)) = v25 + int32(2)
	goto L1
L10:
	;
	goto L9
L11:
	;
	v44 = v21 - int32(1)
	if v44 < int32(0) {
		goto L14
	} else {
		goto L15
	}
L12:
	;
	v260 = v21
	goto L13
L13:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v19)+24)) = v260
	*(*int32)(unsafe.Add(mBase, uint32(v19))) = v260
	v271 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	*(*int32)(unsafe.Add(mBase, uint32(v19)+28)) = v271
	*(*int32)(unsafe.Add(mBase, uint32(v19)+4)) = v271
	v274 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	*(*int32)(unsafe.Add(mBase, uint32(v19)+36)) = v274
	*(*int32)(unsafe.Add(mBase, uint32(v19)+12)) = v274
	v277 = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v19)+32)) = v277
	*(*int32)(unsafe.Add(mBase, uint32(v19)+8)) = int32(_a_F_accum_sum_final_0)
	v284 = F_palloc(m, v260<<(uint(int32(1))%32))
	mBase = m.M
	v285 = m.ExcPending
	if v285 != 0 {
		goto L5
	} else {
		goto L60
	}
L14:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+12)) = int32(0)
	v252 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v260 = v252
	goto L13
L15:
	;
	v47 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	if v44 == int32(0) {
		goto L17
	} else {
		goto L18
	}
L16:
	;
	if v21&int32(1) != 0 {
		goto L30
	} else {
		goto L31
	}
L17:
	;
	v109 = v3
	v110 = v44
	v113 = v3
	goto L16
L18:
	;
	goto L19
L19:
	;
	v59 = v44
	v62 = v3
	v64 = v3
	goto L20
L20:
	;
	v73 = v59 << (uint(int32(2)) % 32)
	v74 = v47 + v73
	v75 = *(*int32)(unsafe.Add(mBase, uint32(v74)))
	v76 = v75 + v62
	if v76 < int32(_a_F_accum_sum_final_1) {
		goto L23
	} else {
		goto L24
	}
L21:
	;
	v109 = v99
	v110 = v103
	v113 = v100
	goto L16
L22:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v74))) = v85
	v89 = v73 + (v47 - int32(4))
	v90 = *(*int32)(unsafe.Add(mBase, uint32(v89)))
	v91 = v90 + v86
	if int32(_a_F_accum_sum_final_1) <= v91 {
		goto L26
	} else {
		goto L27
	}
L23:
	;
	v85 = v76
	v86 = int32(0)
	goto L22
L24:
	;
	goto L25
L25:
	;
	v81 = base.I32_div_u_s(v76, int32(_a_F_accum_sum_final_1))
	v85 = v81*int32(-10000) + v76
	v86 = v81
	goto L22
L26:
	;
	v95 = base.I32_div_u_s(v91, int32(_a_F_accum_sum_final_1))
	v99 = v95*int32(-10000) + v91
	v100 = v95
	goto L28
L27:
	;
	v99 = v91
	v100 = int32(0)
	goto L28
L28:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v89))) = v99
	v102 = int32(2)
	v103 = v59 - v102
	v105 = v64 + v102
	if v105 != v21&int32(-2) {
		v59 = v103
		v62 = v100
		v64 = v105
		goto L20
	} else {
		goto L29
	}
L29:
	;
	goto L21
L30:
	;
	v125 = v47 + v110<<(uint(int32(2))%32)
	v126 = *(*int32)(unsafe.Add(mBase, uint32(v125)))
	v127 = v126 + v113
	v129 = base.I32_rem_u_s(v127, int32(_a_F_accum_sum_final_1))
	if int32(_a_F_accum_sum_final_2) < v127 {
		goto L33
	} else {
		goto L34
	}
L31:
	;
	v134 = v109
	goto L32
L32:
	;
	if int32(0) < v134 {
		goto L36
	} else {
		goto L37
	}
L33:
	;
	v132 = v129
	goto L35
L34:
	;
	v132 = v127
	goto L35
L35:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v125))) = v132
	v134 = v132
	goto L32
L36:
	;
	v137 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(l0)+16)) = uint8(v137)
	goto L38
L37:
	;
	goto L38
L38:
	;
	v139 = int32(1)
	v141 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
	if v21 == v139 {
		goto L40
	} else {
		goto L41
	}
L39:
	;
	if v21&v139 != 0 {
		goto L53
	} else {
		goto L54
	}
L40:
	;
	v204 = v134
	v205 = int32(0)
	v207 = v44
	goto L39
L41:
	;
	goto L42
L42:
	;
	v149 = int32(0)
	v154 = v149
	v156 = v44
	v157 = v149
	goto L43
L43:
	;
	v168 = v156 << (uint(int32(2)) % 32)
	v169 = v141 + v168
	v170 = *(*int32)(unsafe.Add(mBase, uint32(v169)))
	v171 = v170 + v154
	if v171 < int32(_a_F_accum_sum_final_1) {
		goto L46
	} else {
		goto L47
	}
L44:
	;
	v204 = v194
	v205 = v195
	v207 = v198
	goto L39
L45:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v169))) = v180
	v184 = v168 + (v141 - int32(4))
	v185 = *(*int32)(unsafe.Add(mBase, uint32(v184)))
	v186 = v185 + v181
	if int32(_a_F_accum_sum_final_1) <= v186 {
		goto L49
	} else {
		goto L50
	}
L46:
	;
	v180 = v171
	v181 = int32(0)
	goto L45
L47:
	;
	goto L48
L48:
	;
	v176 = base.I32_div_u_s(v171, int32(_a_F_accum_sum_final_1))
	v180 = v176*int32(-10000) + v171
	v181 = v176
	goto L45
L49:
	;
	v190 = base.I32_div_u_s(v186, int32(_a_F_accum_sum_final_1))
	v194 = v190*int32(-10000) + v186
	v195 = v190
	goto L51
L50:
	;
	v194 = v186
	v195 = int32(0)
	goto L51
L51:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v184))) = v194
	v197 = int32(2)
	v198 = v156 - v197
	v200 = v157 + v197
	if v200 != v21&int32(-2) {
		v154 = v195
		v156 = v198
		v157 = v200
		goto L43
	} else {
		goto L52
	}
L52:
	;
	goto L44
L53:
	;
	v220 = v141 + v207<<(uint(int32(2))%32)
	v221 = *(*int32)(unsafe.Add(mBase, uint32(v220)))
	v222 = v221 + v205
	v224 = base.I32_rem_u_s(v222, int32(_a_F_accum_sum_final_1))
	if int32(_a_F_accum_sum_final_2) < v222 {
		goto L56
	} else {
		goto L57
	}
L54:
	;
	v229 = v204
	goto L55
L55:
	;
	if v229 <= int32(0) {
		goto L14
	} else {
		goto L59
	}
L56:
	;
	v227 = v224
	goto L58
L57:
	;
	v227 = v222
	goto L58
L58:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v220))) = v227
	v229 = v227
	goto L55
L59:
	;
	v232 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(l0)+16)) = uint8(v232)
	goto L14
L60:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v19)+40)) = v284
	*(*int32)(unsafe.Add(mBase, uint32(v19)+44)) = v284
	v288 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v291 = F_palloc(m, v288<<(uint(int32(1))%32))
	mBase = m.M
	v292 = m.ExcPending
	if v292 != 0 {
		goto L5
	} else {
		goto L61
	}
L61:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v19)+16)) = v291
	*(*int32)(unsafe.Add(mBase, uint32(v19)+20)) = v291
	v295 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	if int32(0) < v295 {
		goto L62
	} else {
		goto L63
	}
L62:
	;
	v300 = v277
	goto L65
L63:
	;
	goto L64
L64:
	;
	F_add_var(m, v19+int32(24), v19, l1)
	mBase = m.M
	v351 = m.ExcPending
	if v351 != 0 {
		goto L5
	} else {
		goto L68
	}
L65:
	;
	v314 = int32(1)
	v315 = v300 << (uint(v314) % 32)
	v318 = v300 << (uint(int32(2)) % 32)
	v319 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v321 = *(*int32)(unsafe.Add(mBase, uint32(v318+v319)))
	*(*uint16)(unsafe.Add(mBase, uint32(v284+v315))) = uint16(v321)
	v324 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
	v326 = *(*int32)(unsafe.Add(mBase, uint32(v324+v318)))
	*(*uint16)(unsafe.Add(mBase, uint32(v315+v291))) = uint16(v326)
	v329 = v300 + v314
	v330 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	if v329 < v330 {
		v300 = v329
		goto L65
	} else {
		goto L67
	}
L66:
	;
	goto L64
L67:
	;
	goto L66
L68:
	;
	v352 = *(*int32)(unsafe.Add(mBase, uint32(l1)+20))
	v353 = *(*int32)(unsafe.Add(mBase, uint32(l1)))
	if int32(0) < v353 {
		goto L71
	} else {
		goto L72
	}
L69:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l1))) = v433
	*(*int32)(unsafe.Add(mBase, uint32(l1)+20)) = v431
	goto L1
L70:
	;
	*(*int64)(unsafe.Add(mBase, uint32(l1)+4)) = int64(0)
	v431 = v412
	v433 = int32(0)
	goto L69
L71:
	;
	v359 = v352
	v361 = v353
	goto L75
L72:
	;
	goto L73
L73:
	;
	if v353 != 0 {
		v431 = v352
		v433 = v353
		goto L69
	} else {
		goto L83
	}
L74:
	;
	v390 = v361
	goto L79
L75:
	;
	v375 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v359))))
	if v375 != 0 {
		goto L74
	} else {
		goto L77
	}
L76:
	;
	v412 = v352 + v353<<(uint(int32(1))%32)
	goto L70
L77:
	;
	v376 = *(*int32)(unsafe.Add(mBase, uint32(l1)+4))
	v377 = int32(1)
	*(*int32)(unsafe.Add(mBase, uint32(l1)+4)) = v376 - v377
	if v377 < v361 {
		v359 = v359 + int32(2)
		v361 = v361 - v377
		goto L75
	} else {
		goto L78
	}
L78:
	;
	goto L76
L79:
	;
	v407 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v359-int32(2)+v390<<(uint(int32(1))%32)))))
	if v407 != 0 {
		v431 = v359
		v433 = v390
		goto L69
	} else {
		goto L81
	}
L80:
	;
	v412 = v359
	goto L70
L81:
	;
	v408 = int32(1)
	if v408 < v390 {
		v390 = v390 - v408
		goto L79
	} else {
		goto L82
	}
L82:
	;
	goto L80
L83:
	;
	v412 = v352
	goto L70
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
	var v8 int64
	_ = v8
	var v12 int32
	_ = v12
	var v14 int32
	_ = v14
	var v16 int32
	_ = v16
	var v17 int32
	_ = v17
	var v19 int32
	_ = v19
	var v22 int32
	_ = v22
	var v25 int32
	_ = v25
	var v26 int32
	_ = v26
	var v29 int32
	_ = v29
	var v33 int32
	_ = v33
	var v34 int32
	_ = v34
	var v35 int32
	_ = v35
	var v43 int32
	_ = v43
	var v44 int32
	_ = v44
	var v45 int32
	_ = v45
	var v48 int32
	_ = v48
	var v49 int32
	_ = v49
	var v51 int32
	_ = v51
	var v52 int32
	_ = v52
	var v54 int32
	_ = v54
	var v56 int32
	_ = v56
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
	var v76 int32
	_ = v76
	var v77 int32
	_ = v77
	var v79 int32
	_ = v79
	var v84 int32
	_ = v84
	var v97 int32
	_ = v97
	var v101 int32
	_ = v101
	var v102 int32
	_ = v102
	var v103 int32
	_ = v103
	var v111 int32
	_ = v111
	var v112 int32
	_ = v112
	var v113 int32
	_ = v113
	var v116 int32
	_ = v116
	var v117 int32
	_ = v117
	var v119 int32
	_ = v119
	var v120 int32
	_ = v120
	var v122 int32
	_ = v122
	var v124 int32
	_ = v124
	var v134 int32
	_ = v134
	var v135 int32
	_ = v135
	var v136 int32
	_ = v136
	var v139 int32
	_ = v139
	var v140 int32
	_ = v140
	var v141 int32
	_ = v141
	var v144 int32
	_ = v144
	var v145 int32
	_ = v145
	var v147 int32
	_ = v147
	var v152 int32
	_ = v152
	var v165 int32
	_ = v165
	var v168 int32
	_ = v168
	var v169 int32
	_ = v169
	var v174 int32
	_ = v174
	var v182 int32
	_ = v182
	var v186 int32
	_ = v186
	var v191 int32
	_ = v191
	var v194 int32
	_ = v194
	var v195 int32
	_ = v195
	var v198 int32
	_ = v198
	var v201 int32
	_ = v201
	var v202 int32
	_ = v202
	var v207 int32
	_ = v207
	var v211 int32
	_ = v211
	var v215 int32
	_ = v215
	var v220 int32
	_ = v220
	var v221 int32
	_ = v221
	var v224 int32
	_ = v224
	var v226 int32
	_ = v226
	var v227 int32
	_ = v227
	var v230 int32
	_ = v230
	var v231 int32
	_ = v231
	var v236 int32
	_ = v236
	var v240 int32
	_ = v240
	var v245 int32
	_ = v245
	var v257 int32
	_ = v257
	var v258 int32
	_ = v258
	var v262 int64
	_ = v262
	var v263 int64
	_ = v263
	var v264 int64
	_ = v264
	var v284 int32
	_ = v284
	var v285 int32
	_ = v285
	var v290 int32
	_ = v290
	var v295 int32
	_ = v295
	var v300 int32
	_ = v300
	var v302 int64
	_ = v302
	var v303 int64
	_ = v303
	var v304 int64
	_ = v304
	var v306 int32
	_ = v306
	var v307 int32
	_ = v307
	var v326 int32
	_ = v326
	var v328 int32
	_ = v328
	var v329 int64
	_ = v329
	var v330 int64
	_ = v330
	var v333 int32
	_ = v333
	var v341 int32
	_ = v341
	var v344 int32
	_ = v344
	var v345 int32
	_ = v345
	var v347 int32
	_ = v347
	var v348 int32
	_ = v348
	var v353 int32
	_ = v353
	var v361 int32
	_ = v361
	var v366 int32
	_ = v366
	var v368 int32
	_ = v368
	var v375 int32
	_ = v375
	var v376 int32
	_ = v376
	var v379 int32
	_ = v379
	var v382 int32
	_ = v382
	var v383 int32
	_ = v383
	var v388 int32
	_ = v388
	var v392 int32
	_ = v392
	var v397 int32
	_ = v397
	var v401 int32
	_ = v401
	var v404 int32
	_ = v404
	var v405 int32
	_ = v405
	var v407 int32
	_ = v407
	var v408 int32
	_ = v408
	var v413 int32
	_ = v413
	var v421 int32
	_ = v421
	var v426 int32
	_ = v426
	var v431 int32
	_ = v431
	var v432 int32
	_ = v432
	var v437 int32
	_ = v437
	var v444 int32
	_ = v444
	var v449 int32
	_ = v449
	var v451 int32
	_ = v451
	var v458 int32
	_ = v458
	var v465 int32
	_ = v465
	var v474 int32
	_ = v474
	var v475 int32
	_ = v475
	var v476 int32
	_ = v476
	var v481 int32
	_ = v481
	var v485 int32
	_ = v485
	var v490 int32
	_ = v490
	var v504 int32
	_ = v504
	var v512 int32
	_ = v512
	v8 = int64(0)
	v12 = m.G0
	v14 = v12 - int32(208)
	m.G0 = v14
	v16 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v17 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v19 = F_palloc(m, int32(16))
	mBase = m.M
	v22 = m.ExcPending
	if v22 != 0 {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	return int32(0)
L2:
	;
	v25 = F_getid(m, v17, v14+int32(144), v16)
	mBase = m.M
	v26 = m.ExcPending
	if v26 != 0 {
		goto L1
	} else {
		goto L5
	}
L3:
	;
	m.G0 = v14 + int32(208)
	return v512
L4:
	;
	v504 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(l0)+16)) = uint8(v504)
	v512 = int32(0)
	goto L3
L5:
	;
	if v25 == int32(0) {
		goto L4
	} else {
		goto L6
	}
L6:
	;
	v29 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v25))))
	if v29 != int32(61) {
		goto L11
	} else {
		goto L12
	}
L7:
	;
	v333 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v14)+144)))
	if v333 == int32(0) {
		goto L105
	} else {
		goto L106
	}
L8:
	;
	v257 = v227
	v258 = v226
	v262 = v8
	v263 = v8
	v264 = v8
	goto L78
L9:
	;
	if base.Ui32(int32(26)) <= base.Ui32((v227|int32(32)-int32(97))&int32(255)) {
		v326 = v226
		v328 = v224
		v329 = v8
		v330 = v8
		goto L7
	} else {
		goto L77
	}
L10:
	;
	v230 = F_errsave_start(m, v16)
	mBase = m.M
	v231 = m.ExcPending
	if v231 != 0 {
		goto L1
	} else {
		goto L72
	}
L11:
	;
	v33 = v14 + int32(144)
	v34 = int32(_a_F_aclitemin_0)
	v35 = int32(6)
	goto L18
L12:
	;
	v224 = v25
	goto L13
L13:
	;
	v226 = v224 + int32(1)
	v227 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v226))))
	if v227 != int32(42) {
		goto L9
	} else {
		goto L71
	}
L14:
	;
	v194 = F_getid(m, v25, v14+int32(144), v16)
	mBase = m.M
	v195 = m.ExcPending
	if v195 != 0 {
		goto L1
	} else {
		goto L59
	}
L15:
	;
	if v97 == int32(0) {
		goto L14
	} else {
		goto L33
	}
L16:
	;
	v97 = int32(0)
	goto L15
L17:
	;
	v71 = v66
	v72 = v67
	v73 = v68
	goto L27
L18:
	;
	if (v33|v34)&int32(3) != 0 {
		v66 = v33
		v67 = v34
		v68 = v35
		goto L17
	} else {
		goto L21
	}
L20:
	;
	if v56 == int32(0) {
		goto L16
	} else {
		goto L26
	}
L21:
	;
	v43 = v33
	v44 = v34
	v45 = v35
	goto L22
L22:
	;
	v48 = *(*int32)(unsafe.Add(mBase, uint32(v43)))
	v49 = *(*int32)(unsafe.Add(mBase, uint32(v44)))
	if v48 != v49 {
		v66 = v43
		v67 = v44
		v68 = v45
		goto L17
	} else {
		goto L24
	}
L23:
	;
	goto L20
L24:
	;
	v51 = int32(4)
	v52 = v44 + v51
	v54 = v43 + v51
	v56 = v45 - v51
	if base.Ui32(int32(3)) < base.Ui32(v56) {
		v43 = v54
		v44 = v52
		v45 = v56
		goto L22
	} else {
		goto L25
	}
L25:
	;
	goto L23
L26:
	;
	v66 = v54
	v67 = v52
	v68 = v56
	goto L17
L27:
	;
	v76 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v71))))
	v77 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v72))))
	if v76 == v77 {
		goto L29
	} else {
		goto L30
	}
L28:
	;
	v97 = v76 - v77
	goto L15
L29:
	;
	v79 = int32(1)
	v84 = v73 - v79
	if v84 != 0 {
		v71 = v71 + v79
		v72 = v72 + v79
		v73 = v84
		goto L27
	} else {
		goto L32
	}
L30:
	;
	goto L31
L31:
	;
	goto L28
L32:
	;
	goto L16
L33:
	;
	v101 = v14 + int32(144)
	v102 = int32(_a_F_aclitemin_1)
	v103 = int32(5)
	goto L37
L34:
	;
	if v165 == int32(0) {
		goto L14
	} else {
		goto L52
	}
L35:
	;
	v165 = int32(0)
	goto L34
L36:
	;
	v139 = v134
	v140 = v135
	v141 = v136
	goto L46
L37:
	;
	if (v101|v102)&int32(3) != 0 {
		v134 = v101
		v135 = v102
		v136 = v103
		goto L36
	} else {
		goto L40
	}
L39:
	;
	if v124 == int32(0) {
		goto L35
	} else {
		goto L45
	}
L40:
	;
	v111 = v101
	v112 = v102
	v113 = v103
	goto L41
L41:
	;
	v116 = *(*int32)(unsafe.Add(mBase, uint32(v111)))
	v117 = *(*int32)(unsafe.Add(mBase, uint32(v112)))
	if v116 != v117 {
		v134 = v111
		v135 = v112
		v136 = v113
		goto L36
	} else {
		goto L43
	}
L42:
	;
	goto L39
L43:
	;
	v119 = int32(4)
	v120 = v112 + v119
	v122 = v111 + v119
	v124 = v113 - v119
	if base.Ui32(int32(3)) < base.Ui32(v124) {
		v111 = v122
		v112 = v120
		v113 = v124
		goto L41
	} else {
		goto L44
	}
L44:
	;
	goto L42
L45:
	;
	v134 = v122
	v135 = v120
	v136 = v124
	goto L36
L46:
	;
	v144 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v139))))
	v145 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v140))))
	if v144 == v145 {
		goto L48
	} else {
		goto L49
	}
L47:
	;
	v165 = v144 - v145
	goto L34
L48:
	;
	v147 = int32(1)
	v152 = v141 - v147
	if v152 != 0 {
		v139 = v139 + v147
		v140 = v140 + v147
		v141 = v152
		goto L46
	} else {
		goto L51
	}
L49:
	;
	goto L50
L50:
	;
	goto L47
L51:
	;
	goto L35
L52:
	;
	v168 = F_errsave_start(m, v16)
	mBase = m.M
	v169 = m.ExcPending
	if v169 != 0 {
		goto L1
	} else {
		goto L53
	}
L53:
	;
	if v168 == int32(0) {
		goto L4
	} else {
		goto L54
	}
L54:
	;
	F_errcode(m, int32(33685634))
	mBase = m.M
	v174 = m.ExcPending
	if v174 != 0 {
		goto L1
	} else {
		goto L55
	}
L55:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v14)+64)) = v14 + int32(144)
	F_errmsg(m, int32(_a_F_aclitemin_2), v14-int32(-64))
	mBase = m.M
	v182 = m.ExcPending
	if v182 != 0 {
		goto L1
	} else {
		goto L56
	}
L56:
	;
	F_errhint(m, int32(_a_F_aclitemin_3), int32(0))
	mBase = m.M
	v186 = m.ExcPending
	if v186 != 0 {
		goto L1
	} else {
		goto L57
	}
L57:
	;
	F_errsave_finish(m, v16, int32(_a_F_aclitemin_4), int32(294), int32(_a_F_aclitemin_5))
	mBase = m.M
	v191 = m.ExcPending
	if v191 != 0 {
		goto L1
	} else {
		goto L58
	}
L58:
	;
	goto L4
L59:
	;
	if v194 == int32(0) {
		goto L4
	} else {
		goto L60
	}
L60:
	;
	v198 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v14)+144)))
	if v198 == int32(0) {
		goto L61
	} else {
		goto L62
	}
L61:
	;
	v201 = F_errsave_start(m, v16)
	mBase = m.M
	v202 = m.ExcPending
	if v202 != 0 {
		goto L1
	} else {
		goto L64
	}
L62:
	;
	goto L63
L63:
	;
	v221 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v194))))
	if v221 != int32(61) {
		goto L10
	} else {
		goto L70
	}
L64:
	;
	if v201 == int32(0) {
		goto L4
	} else {
		goto L65
	}
L65:
	;
	F_errcode(m, int32(33685634))
	mBase = m.M
	v207 = m.ExcPending
	if v207 != 0 {
		goto L1
	} else {
		goto L66
	}
L66:
	;
	F_errmsg(m, int32(_a_F_aclitemin_6), int32(0))
	mBase = m.M
	v211 = m.ExcPending
	if v211 != 0 {
		goto L1
	} else {
		goto L67
	}
L67:
	;
	F_errhint(m, int32(_a_F_aclitemin_7), int32(0))
	mBase = m.M
	v215 = m.ExcPending
	if v215 != 0 {
		goto L1
	} else {
		goto L68
	}
L68:
	;
	F_errsave_finish(m, v16, int32(_a_F_aclitemin_4), int32(303), int32(_a_F_aclitemin_5))
	mBase = m.M
	v220 = m.ExcPending
	if v220 != 0 {
		goto L1
	} else {
		goto L69
	}
L69:
	;
	goto L4
L70:
	;
	v224 = v194
	goto L13
L71:
	;
	goto L8
L72:
	;
	if v230 == int32(0) {
		goto L4
	} else {
		goto L73
	}
L73:
	;
	F_errcode(m, int32(33685634))
	mBase = m.M
	v236 = m.ExcPending
	if v236 != 0 {
		goto L1
	} else {
		goto L74
	}
L74:
	;
	F_errmsg(m, int32(_a_F_aclitemin_8), int32(0))
	mBase = m.M
	v240 = m.ExcPending
	if v240 != 0 {
		goto L1
	} else {
		goto L75
	}
L75:
	;
	F_errsave_finish(m, v16, int32(_a_F_aclitemin_4), int32(309), int32(_a_F_aclitemin_5))
	mBase = m.M
	v245 = m.ExcPending
	if v245 != 0 {
		goto L1
	} else {
		goto L76
	}
L76:
	;
	goto L4
L77:
	;
	goto L8
L78:
	;
	switch v257&int32(255) - int32(42) {
	case 0:
		goto L81
	default:
		goto L82
	case 23:
		goto L84
	case 25:
		goto L88
	case 26:
		goto L93
	case 42:
		goto L87
	case 43:
		goto L89
	case 46:
		goto L90
	case 55:
		v302 = int64(1)
		v303 = v263
		goto L80
	case 57:
		goto L86
	case 58:
		goto L94
	case 67:
		goto L83
	case 72:
		goto L96
	case 73:
		goto L85
	case 74:
		goto L91
	case 77:
		goto L95
	case 78:
		goto L92
	}
L79:
	;
	v326 = v306
	v328 = v258
	v329 = v304 & int64(4294967295)
	v330 = v303 << (uint(int64(32)) % 64)
	goto L7
L80:
	;
	v304 = v302 | v262
	v306 = v258 + int32(1)
	v307 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v306))))
	if v307 == int32(42) {
		v257 = v307
		v258 = v306
		v262 = v304
		v263 = v303
		v264 = v302
		goto L78
	} else {
		goto L102
	}
L81:
	;
	v302 = v264
	v303 = v263 | v264
	goto L80
L82:
	;
	v284 = F_errsave_start(m, v16)
	mBase = m.M
	v285 = m.ExcPending
	if v285 != 0 {
		goto L1
	} else {
		goto L97
	}
L83:
	;
	v302 = int64(16384)
	v303 = v263
	goto L80
L84:
	;
	v302 = int64(8192)
	v303 = v263
	goto L80
L85:
	;
	v302 = int64(4096)
	v303 = v263
	goto L80
L86:
	;
	v302 = int64(2048)
	v303 = v263
	goto L80
L87:
	;
	v302 = int64(1024)
	v303 = v263
	goto L80
L88:
	;
	v302 = int64(512)
	v303 = v263
	goto L80
L89:
	;
	v302 = int64(256)
	v303 = v263
	goto L80
L90:
	;
	v302 = int64(128)
	v303 = v263
	goto L80
L91:
	;
	v302 = int64(64)
	v303 = v263
	goto L80
L92:
	;
	v302 = int64(32)
	v303 = v263
	goto L80
L93:
	;
	v302 = int64(16)
	v303 = v263
	goto L80
L94:
	;
	v302 = int64(8)
	v303 = v263
	goto L80
L95:
	;
	v302 = int64(4)
	v303 = v263
	goto L80
L96:
	;
	v302 = int64(2)
	v303 = v263
	goto L80
L97:
	;
	if v284 == int32(0) {
		goto L4
	} else {
		goto L98
	}
L98:
	;
	F_errcode(m, int32(33685634))
	mBase = m.M
	v290 = m.ExcPending
	if v290 != 0 {
		goto L1
	} else {
		goto L99
	}
L99:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v14))) = int32(_a_F_aclitemin_9)
	F_errmsg(m, int32(_a_F_aclitemin_10), v14)
	mBase = m.M
	v295 = m.ExcPending
	if v295 != 0 {
		goto L1
	} else {
		goto L100
	}
L100:
	;
	F_errsave_finish(m, v16, int32(_a_F_aclitemin_4), int32(369), int32(_a_F_aclitemin_5))
	mBase = m.M
	v300 = m.ExcPending
	if v300 != 0 {
		goto L1
	} else {
		goto L101
	}
L101:
	;
	goto L4
L102:
	;
	if base.Ui32((v307|int32(32)-int32(97))&int32(255)) < base.Ui32(int32(26)) {
		v257 = v307
		v258 = v306
		v262 = v304
		v263 = v303
		v264 = v302
		goto L78
	} else {
		goto L103
	}
L103:
	;
	goto L79
L104:
	;
	v368 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v326))))
	if v368 == int32(47) {
		goto L116
	} else {
		goto L117
	}
L105:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v19))) = int32(0)
	goto L104
L106:
	;
	goto L107
L107:
	;
	v341 = int32(0)
	v344 = F_GetSysCacheOid(m, int32(10), v14+int32(144), v341, v341, v341)
	mBase = m.M
	v345 = m.ExcPending
	if v345 != 0 {
		goto L1
	} else {
		goto L108
	}
L108:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v19))) = v344
	if v344 != 0 {
		goto L104
	} else {
		goto L109
	}
L109:
	;
	v347 = F_errsave_start(m, v16)
	mBase = m.M
	v348 = m.ExcPending
	if v348 != 0 {
		goto L1
	} else {
		goto L110
	}
L110:
	;
	if v347 == int32(0) {
		goto L4
	} else {
		goto L111
	}
L111:
	;
	F_errcode(m, int32(67137668))
	mBase = m.M
	v353 = m.ExcPending
	if v353 != 0 {
		goto L1
	} else {
		goto L112
	}
L112:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v14)+48)) = v14 + int32(144)
	F_errmsg(m, int32(_a_F_aclitemin_11), v14+int32(48))
	mBase = m.M
	v361 = m.ExcPending
	if v361 != 0 {
		goto L1
	} else {
		goto L113
	}
L113:
	;
	F_errsave_finish(m, v16, int32(_a_F_aclitemin_4), int32(383), int32(_a_F_aclitemin_5))
	mBase = m.M
	v366 = m.ExcPending
	if v366 != 0 {
		goto L1
	} else {
		goto L114
	}
L114:
	;
	goto L4
L115:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v19)+8)) = v329 | v330
	v458 = v451
	goto L141
L116:
	;
	v375 = F_getid(m, v328+int32(2), v14+int32(80), v16)
	mBase = m.M
	v376 = m.ExcPending
	if v376 != 0 {
		goto L1
	} else {
		goto L119
	}
L117:
	;
	goto L118
L118:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v19)+4)) = int32(10)
	v431 = F_errstart(m, int32(19), int32(0))
	mBase = m.M
	v432 = m.ExcPending
	if v432 != 0 {
		goto L1
	} else {
		goto L136
	}
L119:
	;
	if v375 == int32(0) {
		goto L4
	} else {
		goto L120
	}
L120:
	;
	v379 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v14)+80)))
	if v379 == int32(0) {
		goto L121
	} else {
		goto L122
	}
L121:
	;
	v382 = F_errsave_start(m, v16)
	mBase = m.M
	v383 = m.ExcPending
	if v383 != 0 {
		goto L1
	} else {
		goto L124
	}
L122:
	;
	goto L123
L123:
	;
	v401 = int32(0)
	v404 = F_GetSysCacheOid(m, int32(10), v14+int32(80), v401, v401, v401)
	mBase = m.M
	v405 = m.ExcPending
	if v405 != 0 {
		goto L1
	} else {
		goto L129
	}
L124:
	;
	if v382 == int32(0) {
		goto L4
	} else {
		goto L125
	}
L125:
	;
	F_errcode(m, int32(33685634))
	mBase = m.M
	v388 = m.ExcPending
	if v388 != 0 {
		goto L1
	} else {
		goto L126
	}
L126:
	;
	F_errmsg(m, int32(_a_F_aclitemin_12), int32(0))
	mBase = m.M
	v392 = m.ExcPending
	if v392 != 0 {
		goto L1
	} else {
		goto L127
	}
L127:
	;
	F_errsave_finish(m, v16, int32(_a_F_aclitemin_4), int32(398), int32(_a_F_aclitemin_5))
	mBase = m.M
	v397 = m.ExcPending
	if v397 != 0 {
		goto L1
	} else {
		goto L128
	}
L128:
	;
	goto L4
L129:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v19)+4)) = v404
	if v404 != 0 {
		v451 = v375
		goto L115
	} else {
		goto L130
	}
L130:
	;
	v407 = F_errsave_start(m, v16)
	mBase = m.M
	v408 = m.ExcPending
	if v408 != 0 {
		goto L1
	} else {
		goto L131
	}
L131:
	;
	if v407 == int32(0) {
		goto L4
	} else {
		goto L132
	}
L132:
	;
	F_errcode(m, int32(67137668))
	mBase = m.M
	v413 = m.ExcPending
	if v413 != 0 {
		goto L1
	} else {
		goto L133
	}
L133:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v14)+16)) = v14 + int32(80)
	F_errmsg(m, int32(_a_F_aclitemin_11), v14+int32(16))
	mBase = m.M
	v421 = m.ExcPending
	if v421 != 0 {
		goto L1
	} else {
		goto L134
	}
L134:
	;
	F_errsave_finish(m, v16, int32(_a_F_aclitemin_4), int32(403), int32(_a_F_aclitemin_5))
	mBase = m.M
	v426 = m.ExcPending
	if v426 != 0 {
		goto L1
	} else {
		goto L135
	}
L135:
	;
	goto L4
L136:
	;
	if v431 == int32(0) {
		v451 = v326
		goto L115
	} else {
		goto L137
	}
L137:
	;
	F_errcode(m, int32(1792))
	mBase = m.M
	v437 = m.ExcPending
	if v437 != 0 {
		goto L1
	} else {
		goto L138
	}
L138:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v14)+32)) = int32(10)
	F_errmsg(m, int32(_a_F_aclitemin_13), v14+int32(32))
	mBase = m.M
	v444 = m.ExcPending
	if v444 != 0 {
		goto L1
	} else {
		goto L139
	}
L139:
	;
	F_errfinish(m, int32(_a_F_aclitemin_4), int32(411), int32(_a_F_aclitemin_5))
	mBase = m.M
	v449 = m.ExcPending
	if v449 != 0 {
		goto L1
	} else {
		goto L140
	}
L140:
	;
	v451 = v326
	goto L115
L141:
	;
	v465 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v458))))
	if base.Ui32(v465-int32(9)) < base.Ui32(int32(5)) {
		goto L143
	} else {
		goto L144
	}
L143:
	;
	v458 = v458 + int32(1)
	goto L141
L144:
	;
	if v465 == int32(32) {
		goto L143
	} else {
		goto L145
	}
L145:
	;
	if v465 == int32(0) {
		v512 = v19
		goto L3
	} else {
		goto L146
	}
L146:
	;
	v474 = int32(0)
	v475 = F_errsave_start(m, v16)
	mBase = m.M
	v476 = m.ExcPending
	if v476 != 0 {
		goto L1
	} else {
		goto L147
	}
L147:
	;
	if v475 == int32(0) {
		v512 = v474
		goto L3
	} else {
		goto L148
	}
L148:
	;
	F_errcode(m, int32(33685634))
	mBase = m.M
	v481 = m.ExcPending
	if v481 != 0 {
		goto L1
	} else {
		goto L149
	}
L149:
	;
	F_errmsg(m, int32(_a_F_aclitemin_14), int32(0))
	mBase = m.M
	v485 = m.ExcPending
	if v485 != 0 {
		goto L1
	} else {
		goto L150
	}
L150:
	;
	F_errsave_finish(m, v16, int32(_a_F_aclitemin_4), int32(633), int32(_a_F_aclitemin_15))
	mBase = m.M
	v490 = m.ExcPending
	if v490 != 0 {
		goto L1
	} else {
		goto L151
	}
L151:
	;
	v512 = v474
	goto L3
}
func F_acos(m *base.Module, l0 float64) float64 {
	var v6 int64
	_ = v6
	var v11 int32
	_ = v11
	var v24 float64
	_ = v24
	var v36 float64
	_ = v36
	var v77 float64
	_ = v77
	var v80 float64
	_ = v80
	var v81 float64
	_ = v81
	var v117 float64
	_ = v117
	var v120 float64
	_ = v120
	var v123 float64
	_ = v123
	var v124 float64
	_ = v124
	var v160 float64
	_ = v160
	var v166 float64
	_ = v166
	var v171 float64
	_ = v171
	v6 = base.I64_reinterpret_f64(l0)
	v11 = base.I32_wrap_i64(int64(base.Ui64(v6)>>(uint(int64(32))%64))) & int32(2147483647)
	if base.Ui32(int32(1072693248)) <= base.Ui32(v11) {
		if base.I32_wrap_i64(v6)|(v11-int32(1072693248)) == int32(0) {
			if int64(0) <= v6 {
				v24 = float64(0)
			} else {
				v24 = float64(3.141592653589793)
			}
			return v24
		} else {
			return base.F64_div(float64(0), base.F64_sub(l0, l0))
		}
	} else {
		if base.Ui32(v11) <= base.Ui32(int32(1071644671)) {
			if base.Ui32(v11) < base.Ui32(int32(1012924417)) {
				v171 = float64(1.5707963267948966)
				return v171
			} else {
				v36 = base.F64_mul(l0, l0)
				return base.F64_add(base.F64_sub(base.F64_sub(float64(6.123233995736766e-17), base.F64_mul(l0, base.F64_div(base.F64_mul(v36, base.F64_add(base.F64_mul(v36, base.F64_add(base.F64_mul(v36, base.F64_add(base.F64_mul(v36, base.F64_add(base.F64_mul(v36, base.F64_add(base.F64_mul(v36, float64(3.479331075960212e-05)), float64(0.0007915349942898145))), float64(-0.04005553450067941))), float64(0.20121253213486293))), float64(-0.3255658186224009))), float64(0.16666666666666666))), base.F64_add(base.F64_mul(v36, base.F64_add(base.F64_mul(v36, base.F64_add(base.F64_mul(v36, base.F64_add(base.F64_mul(v36, float64(0.07703815055590194)), float64(-0.6882839716054533))), float64(2.0209457602335057))), float64(-2.403394911734414))), float64(1))))), l0), float64(1.5707963267948966))
			}
		} else {
			if v6 < int64(0) {
				v77 = float64(1)
				v80 = base.F64_mul(base.F64_add(l0, v77), float64(0.5))
				v81 = base.F64_sqrt(v80)
				v117 = base.F64_sub(float64(1.5707963267948966), base.F64_add(v81, base.F64_add(base.F64_mul(v81, base.F64_div(base.F64_mul(v80, base.F64_add(base.F64_mul(v80, base.F64_add(base.F64_mul(v80, base.F64_add(base.F64_mul(v80, base.F64_add(base.F64_mul(v80, base.F64_add(base.F64_mul(v80, float64(3.479331075960212e-05)), float64(0.0007915349942898145))), float64(-0.04005553450067941))), float64(0.20121253213486293))), float64(-0.3255658186224009))), float64(0.16666666666666666))), base.F64_add(base.F64_mul(v80, base.F64_add(base.F64_mul(v80, base.F64_add(base.F64_mul(v80, base.F64_add(base.F64_mul(v80, float64(0.07703815055590194)), float64(-0.6882839716054533))), float64(2.0209457602335057))), float64(-2.403394911734414))), v77))), float64(-6.123233995736766e-17))))
				return base.F64_add(v117, v117)
			} else {
				v120 = float64(1)
				v123 = base.F64_mul(base.F64_sub(v120, l0), float64(0.5))
				v124 = base.F64_sqrt(v123)
				v160 = base.F64_reinterpret_i64(base.I64_reinterpret_f64(v124) & int64(-4294967296))
				v166 = base.F64_add(base.F64_add(base.F64_mul(v124, base.F64_div(base.F64_mul(v123, base.F64_add(base.F64_mul(v123, base.F64_add(base.F64_mul(v123, base.F64_add(base.F64_mul(v123, base.F64_add(base.F64_mul(v123, base.F64_add(base.F64_mul(v123, float64(3.479331075960212e-05)), float64(0.0007915349942898145))), float64(-0.04005553450067941))), float64(0.20121253213486293))), float64(-0.3255658186224009))), float64(0.16666666666666666))), base.F64_add(base.F64_mul(v123, base.F64_add(base.F64_mul(v123, base.F64_add(base.F64_mul(v123, base.F64_add(base.F64_mul(v123, float64(0.07703815055590194)), float64(-0.6882839716054533))), float64(2.0209457602335057))), float64(-2.403394911734414))), v120))), base.F64_div(base.F64_sub(v123, base.F64_mul(v160, v160)), base.F64_add(v124, v160))), v160)
				v171 = base.F64_add(v166, v166)
				return v171
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
	var v24 int64
	_ = v24
	var v26 int32
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
			v24 = *(*int64)(unsafe.Add(mBase, uint32(l0)))
			*(*int64)(unsafe.Add(mBase, uint32(v23))) = v24
			v26 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
			*(*int32)(unsafe.Add(mBase, uint32(v23)+8)) = v26
			v28 = *(*int32)(unsafe.Add(mBase, uint32(l1)+8))
			*(*int32)(unsafe.Add(mBase, uint32(l1)+8)) = v28 + int32(1)
			return
		}
	} else {
		v19 = v6
		v20 = v7
		v23 = v20*int32(12) + v19
		v24 = *(*int64)(unsafe.Add(mBase, uint32(l0)))
		*(*int64)(unsafe.Add(mBase, uint32(v23))) = v24
		v26 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
		*(*int32)(unsafe.Add(mBase, uint32(v23)+8)) = v26
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
	var v20 int32
	_ = v20
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
	var v43 int32
	_ = v43
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
	var v59 int32
	_ = v59
	var v65 int32
	_ = v65
	var v66 int32
	_ = v66
	var v68 int32
	_ = v68
	var v71 int32
	_ = v71
	var v72 int32
	_ = v72
	var v77 int32
	_ = v77
	var v85 int32
	_ = v85
	var v87 int32
	_ = v87
	var v89 int32
	_ = v89
	var v90 int32
	_ = v90
	var v93 int32
	_ = v93
	var v98 int32
	_ = v98
	var v103 int32
	_ = v103
	var v104 int32
	_ = v104
	var v106 int32
	_ = v106
	var v107 int32
	_ = v107
	var v109 int32
	_ = v109
	var v110 int32
	_ = v110
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
	v20 = v4
	goto L4
L4:
	;
	v24 = *(*int32)(unsafe.Add(mBase, uint32(l1)+12))
	v28 = *(*int32)(unsafe.Add(mBase, uint32(v24+v20<<(uint(int32(2))%32))))
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
	v139 = v20 + int32(1)
	v140 = *(*int32)(unsafe.Add(mBase, uint32(l1)+4))
	if v139 < v140 {
		v20 = v139
		goto L4
	} else {
		goto L37
	}
L7:
	;
	F_distribute_restrictinfo_to_rels(m, l0, v28)
	mBase = m.M
	v128 = m.ExcPending
	if v128 != 0 {
		goto L32
	} else {
		goto L36
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
	v43 = v32
	goto L10
L10:
	;
	v45 = *(*int32)(unsafe.Add(mBase, uint32(v29)+12))
	v49 = *(*int32)(unsafe.Add(mBase, uint32(v45+v43<<(uint(int32(2))%32))))
	v50 = *(*int32)(unsafe.Add(mBase, uint32(v49)+28))
	v51 = *(*int32)(unsafe.Add(mBase, uint32(v28)+28))
	v52 = int32(0)
	v59 = base.B2i32(v50|v51 == v52)
	if v50 == v52 {
		v98 = v59
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
		goto L24
	} else {
		goto L25
	}
L13:
	;
	goto L12
L14:
	;
	if v51 == int32(0) {
		v98 = v59
		goto L13
	} else {
		goto L15
	}
L15:
	;
	v65 = *(*int32)(unsafe.Add(mBase, uint32(v50)+4))
	v66 = *(*int32)(unsafe.Add(mBase, uint32(v51)+4))
	if v65 != v66 {
		v98 = int32(0)
		goto L13
	} else {
		goto L16
	}
L16:
	;
	v68 = int32(1)
	if v65 <= v68 {
		goto L17
	} else {
		goto L18
	}
L17:
	;
	v71 = v68
	goto L19
L18:
	;
	v71 = v65
	goto L19
L19:
	;
	v72 = int32(8)
	v77 = int32(0)
	goto L20
L20:
	;
	v85 = v77 << (uint(int32(2)) % 32)
	v87 = *(*int32)(unsafe.Add(mBase, uint32(v50+v72+v85)))
	v89 = *(*int32)(unsafe.Add(mBase, uint32(v85+(v51+v72))))
	v90 = base.B2i32(v87 == v89)
	if v89 != v87 {
		v98 = v90
		goto L13
	} else {
		goto L22
	}
L21:
	;
	v98 = v90
	goto L13
L22:
	;
	v93 = v77 + int32(1)
	if v93 != v71 {
		v77 = v93
		goto L20
	} else {
		goto L23
	}
L23:
	;
	goto L21
L24:
	;
	if v28 == v49 {
		goto L6
	} else {
		goto L27
	}
L25:
	;
	goto L26
L26:
	;
	v115 = v43 + int32(1)
	v116 = *(*int32)(unsafe.Add(mBase, uint32(v29)+4))
	if v115 < v116 {
		v43 = v115
		goto L10
	} else {
		goto L35
	}
L27:
	;
	v103 = *(*int32)(unsafe.Add(mBase, uint32(v28)+60))
	if v103 != 0 {
		goto L28
	} else {
		goto L29
	}
L28:
	;
	v104 = *(*int32)(unsafe.Add(mBase, uint32(v49)+60))
	if v104 == v103 {
		goto L6
	} else {
		goto L31
	}
L29:
	;
	goto L30
L30:
	;
	v106 = *(*int32)(unsafe.Add(mBase, uint32(v28)+56))
	v107 = *(*int32)(unsafe.Add(mBase, uint32(v49)+56))
	*(*int32)(unsafe.Add(mBase, uint32(v28)+56)) = v107
	v109 = F_equal(m, v28, v49)
	mBase = m.M
	v110 = m.ExcPending
	if v110 != 0 {
		goto L32
	} else {
		goto L33
	}
L31:
	;
	goto L30
L32:
	;
	return
L33:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v28)+56)) = v106
	if v109 != 0 {
		goto L6
	} else {
		goto L34
	}
L34:
	;
	goto L26
L35:
	;
	goto L11
L36:
	;
	goto L6
L37:
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
	var v48 int32
	_ = v48
	var v49 int32
	_ = v49
	var v51 int32
	_ = v51
	var v54 int32
	_ = v54
	var v55 int32
	_ = v55
	var v60 int32
	_ = v60
	var v67 int32
	_ = v67
	var v69 int32
	_ = v69
	var v71 int32
	_ = v71
	var v72 int32
	_ = v72
	var v74 int32
	_ = v74
	var v76 int32
	_ = v76
	var v80 int32
	_ = v80
	var v86 int32
	_ = v86
	var v87 int32
	_ = v87
	var v88 int32
	_ = v88
	var v89 int32
	_ = v89
	var v91 int32
	_ = v91
	var v92 int32
	_ = v92
	var v95 int64
	_ = v95
	var v97 int64
	_ = v97
	var v99 int64
	_ = v99
	var v103 int32
	_ = v103
	var v109 int32
	_ = v109
	var v110 int32
	_ = v110
	var v111 int32
	_ = v111
	var v118 int32
	_ = v118
	var v119 int32
	_ = v119
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
	v118 = F_expression_tree_mutator_impl(m, l0, int32(1052), l1)
	mBase = m.M
	v119 = m.ExcPending
	if v119 != 0 {
		goto L15
	} else {
		goto L42
	}
L5:
	;
	v103 = *(*int32)(unsafe.Add(mBase, uint32(l1)+8))
	*(*int32)(unsafe.Add(mBase, uint32(l1)+8)) = v103 + int32(1)
	v109 = F_query_tree_mutator_impl(m, l0, int32(1052), l1, int32(0))
	mBase = m.M
	v110 = m.ExcPending
	if v110 != 0 {
		goto L15
	} else {
		goto L41
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
	if v38 == v39 {
		v80 = v39
		goto L25
	} else {
		goto L26
	}
L22:
	;
	goto L23
L23:
	;
	v86 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	v87 = *(*int32)(unsafe.Add(mBase, uint32(l1)+4))
	v88 = F_bms_union(m, v86, v87)
	mBase = m.M
	v89 = m.ExcPending
	if v89 != 0 {
		goto L15
	} else {
		goto L39
	}
L24:
	;
	if v80 == int32(0) {
		goto L4
	} else {
		goto L38
	}
L25:
	;
	goto L24
L26:
	;
	if v37 == int32(0) {
		v80 = v39
		goto L25
	} else {
		goto L27
	}
L27:
	;
	v48 = *(*int32)(unsafe.Add(mBase, uint32(v38)+4))
	v49 = *(*int32)(unsafe.Add(mBase, uint32(v37)+4))
	if v48 < v49 {
		goto L28
	} else {
		goto L29
	}
L28:
	;
	v51 = v48
	goto L30
L29:
	;
	v51 = v49
	goto L30
L30:
	;
	if v51 <= int32(1) {
		goto L31
	} else {
		goto L32
	}
L31:
	;
	v54 = int32(1)
	goto L33
L32:
	;
	v54 = v51
	goto L33
L33:
	;
	v55 = int32(8)
	v60 = int32(0)
	goto L34
L34:
	;
	v67 = v60 << (uint(int32(2)) % 32)
	v69 = *(*int32)(unsafe.Add(mBase, uint32(v37+v55+v67)))
	v71 = *(*int32)(unsafe.Add(mBase, uint32(v67+(v38+v55))))
	v72 = v69 & v71
	v74 = base.B2i32(v72 != int32(0))
	if v72 != 0 {
		v80 = v74
		goto L25
	} else {
		goto L36
	}
L35:
	;
	v80 = v74
	goto L25
L36:
	;
	v76 = v60 + int32(1)
	if v76 != v54 {
		v60 = v76
		goto L34
	} else {
		goto L37
	}
L37:
	;
	goto L35
L38:
	;
	goto L23
L39:
	;
	v91 = F_palloc0(m, int32(24))
	mBase = m.M
	v92 = m.ExcPending
	if v92 != 0 {
		goto L15
	} else {
		goto L40
	}
L40:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v91))) = int32(319)
	v95 = *(*int64)(unsafe.Add(mBase, uint32(l0)+8))
	*(*int64)(unsafe.Add(mBase, uint32(v91)+8)) = v95
	v97 = *(*int64)(unsafe.Add(mBase, uint32(l0)+16))
	*(*int64)(unsafe.Add(mBase, uint32(v91)+16)) = v97
	v99 = *(*int64)(unsafe.Add(mBase, uint32(l0)))
	*(*int64)(unsafe.Add(mBase, uint32(v91))) = v99
	*(*int32)(unsafe.Add(mBase, uint32(v91)+12)) = v88
	return v91
L41:
	;
	v111 = *(*int32)(unsafe.Add(mBase, uint32(l1)+8))
	*(*int32)(unsafe.Add(mBase, uint32(l1)+8)) = v111 - int32(1)
	return v109
L42:
	;
	return v118
}
func F_add_paths_to_joinrel(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32, l5 int32, l6 int32) {
	mBase := m.M
	_ = mBase
	var v8 int32
	_ = v8
	var v35 int32
	_ = v35
	var v37 int32
	_ = v37
	var v41 int32
	_ = v41
	var v44 int32
	_ = v44
	var v46 int32
	_ = v46
	var v47 int32
	_ = v47
	var v55 int32
	_ = v55
	var v56 int32
	_ = v56
	var v57 int32
	_ = v57
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
	var v85 int32
	_ = v85
	var v87 int32
	_ = v87
	var v89 int32
	_ = v89
	var v92 int32
	_ = v92
	var v94 int32
	_ = v94
	var v96 int32
	_ = v96
	var v101 int32
	_ = v101
	var v110 int32
	_ = v110
	var v111 int32
	_ = v111
	var v112 int32
	_ = v112
	var v114 int32
	_ = v114
	var v115 int32
	_ = v115
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
	var v124 int32
	_ = v124
	var v126 int32
	_ = v126
	var v135 int32
	_ = v135
	var v138 int32
	_ = v138
	var v156 int32
	_ = v156
	var v158 int32
	_ = v158
	var v159 int32
	_ = v159
	var v179 int32
	_ = v179
	var v183 int32
	_ = v183
	var v184 int32
	_ = v184
	var v185 int32
	_ = v185
	var v186 int32
	_ = v186
	var v187 int32
	_ = v187
	var v196 int32
	_ = v196
	var v197 int32
	_ = v197
	var v199 int32
	_ = v199
	var v202 int32
	_ = v202
	var v203 int32
	_ = v203
	var v208 int32
	_ = v208
	var v215 int32
	_ = v215
	var v217 int32
	_ = v217
	var v219 int32
	_ = v219
	var v222 int32
	_ = v222
	var v224 int32
	_ = v224
	var v226 int32
	_ = v226
	var v231 int32
	_ = v231
	var v240 int32
	_ = v240
	var v243 int32
	_ = v243
	var v246 int32
	_ = v246
	var v247 int32
	_ = v247
	var v248 int32
	_ = v248
	var v252 int32
	_ = v252
	var v253 int32
	_ = v253
	var v254 int32
	_ = v254
	var v255 int32
	_ = v255
	var v264 int32
	_ = v264
	var v265 int32
	_ = v265
	var v267 int32
	_ = v267
	var v270 int32
	_ = v270
	var v271 int32
	_ = v271
	var v276 int32
	_ = v276
	var v283 int32
	_ = v283
	var v285 int32
	_ = v285
	var v287 int32
	_ = v287
	var v290 int32
	_ = v290
	var v292 int32
	_ = v292
	var v294 int32
	_ = v294
	var v299 int32
	_ = v299
	var v308 int32
	_ = v308
	var v311 int32
	_ = v311
	var v312 int32
	_ = v312
	var v321 int32
	_ = v321
	var v322 int32
	_ = v322
	var v324 int32
	_ = v324
	var v327 int32
	_ = v327
	var v328 int32
	_ = v328
	var v333 int32
	_ = v333
	var v340 int32
	_ = v340
	var v342 int32
	_ = v342
	var v344 int32
	_ = v344
	var v347 int32
	_ = v347
	var v349 int32
	_ = v349
	var v351 int32
	_ = v351
	var v356 int32
	_ = v356
	var v365 int32
	_ = v365
	var v368 int32
	_ = v368
	var v370 int32
	_ = v370
	var v371 int32
	_ = v371
	var v380 int32
	_ = v380
	var v381 int32
	_ = v381
	var v383 int32
	_ = v383
	var v386 int32
	_ = v386
	var v387 int32
	_ = v387
	var v392 int32
	_ = v392
	var v399 int32
	_ = v399
	var v401 int32
	_ = v401
	var v403 int32
	_ = v403
	var v406 int32
	_ = v406
	var v408 int32
	_ = v408
	var v410 int32
	_ = v410
	var v415 int32
	_ = v415
	var v424 int32
	_ = v424
	var v428 int32
	_ = v428
	var v429 int32
	_ = v429
	var v438 int32
	_ = v438
	var v439 int32
	_ = v439
	var v441 int32
	_ = v441
	var v444 int32
	_ = v444
	var v445 int32
	_ = v445
	var v450 int32
	_ = v450
	var v457 int32
	_ = v457
	var v459 int32
	_ = v459
	var v461 int32
	_ = v461
	var v464 int32
	_ = v464
	var v466 int32
	_ = v466
	var v468 int32
	_ = v468
	var v473 int32
	_ = v473
	var v482 int32
	_ = v482
	var v486 int32
	_ = v486
	var v488 int32
	_ = v488
	var v489 int32
	_ = v489
	var v490 int32
	_ = v490
	var v491 int32
	_ = v491
	var v493 int32
	_ = v493
	var v494 int32
	_ = v494
	var v513 int32
	_ = v513
	var v530 int32
	_ = v530
	var v565 int32
	_ = v565
	var v566 int32
	_ = v566
	var v585 int32
	_ = v585
	var v602 int32
	_ = v602
	var v637 int32
	_ = v637
	var v638 int32
	_ = v638
	var v640 int32
	_ = v640
	var v641 int32
	_ = v641
	var v643 int32
	_ = v643
	var v644 int32
	_ = v644
	var v656 int32
	_ = v656
	var v658 int32
	_ = v658
	var v680 int32
	_ = v680
	var v681 int32
	_ = v681
	var v698 int32
	_ = v698
	var v700 int32
	_ = v700
	var v726 int32
	_ = v726
	var v741 int32
	_ = v741
	var v743 int32
	_ = v743
	var v778 int32
	_ = v778
	var v805 int32
	_ = v805
	var v806 int32
	_ = v806
	var v807 int32
	_ = v807
	var v809 int32
	_ = v809
	var v814 int32
	_ = v814
	var v819 int32
	_ = v819
	var v833 int32
	_ = v833
	var v842 int32
	_ = v842
	var v856 int32
	_ = v856
	var v860 int32
	_ = v860
	var v861 int32
	_ = v861
	var v862 int32
	_ = v862
	var v863 int32
	_ = v863
	var v864 int32
	_ = v864
	var v873 int32
	_ = v873
	var v874 int32
	_ = v874
	var v876 int32
	_ = v876
	var v879 int32
	_ = v879
	var v880 int32
	_ = v880
	var v885 int32
	_ = v885
	var v892 int32
	_ = v892
	var v894 int32
	_ = v894
	var v896 int32
	_ = v896
	var v899 int32
	_ = v899
	var v901 int32
	_ = v901
	var v903 int32
	_ = v903
	var v908 int32
	_ = v908
	var v917 int32
	_ = v917
	var v920 int32
	_ = v920
	var v921 int32
	_ = v921
	var v922 int32
	_ = v922
	var v924 int32
	_ = v924
	var v925 int32
	_ = v925
	var v938 int32
	_ = v938
	var v962 int32
	_ = v962
	var v966 int32
	_ = v966
	var v967 float64
	_ = v967
	var v968 int32
	_ = v968
	var v970 int32
	_ = v970
	var v971 int32
	_ = v971
	var v972 int32
	_ = v972
	var v973 int64
	_ = v973
	var v989 int32
	_ = v989
	var v993 float64
	_ = v993
	var v994 int32
	_ = v994
	var v996 int32
	_ = v996
	var v999 float64
	_ = v999
	var v1000 float64
	_ = v1000
	var v1002 float64
	_ = v1002
	var v1005 float64
	_ = v1005
	var v1008 float64
	_ = v1008
	var v1048 int32
	_ = v1048
	var v1051 int32
	_ = v1051
	var v1052 int32
	_ = v1052
	var v1061 int32
	_ = v1061
	var v1089 int32
	_ = v1089
	var v1093 int32
	_ = v1093
	var v1094 int32
	_ = v1094
	var v1095 int32
	_ = v1095
	var v1104 int32
	_ = v1104
	var v1105 int32
	_ = v1105
	var v1107 int32
	_ = v1107
	var v1110 int32
	_ = v1110
	var v1111 int32
	_ = v1111
	var v1116 int32
	_ = v1116
	var v1123 int32
	_ = v1123
	var v1125 int32
	_ = v1125
	var v1127 int32
	_ = v1127
	var v1128 int32
	_ = v1128
	var v1130 int32
	_ = v1130
	var v1132 int32
	_ = v1132
	var v1136 int32
	_ = v1136
	var v1142 int32
	_ = v1142
	var v1143 int32
	_ = v1143
	var v1152 int32
	_ = v1152
	var v1153 int32
	_ = v1153
	var v1155 int32
	_ = v1155
	var v1158 int32
	_ = v1158
	var v1159 int32
	_ = v1159
	var v1164 int32
	_ = v1164
	var v1171 int32
	_ = v1171
	var v1173 int32
	_ = v1173
	var v1175 int32
	_ = v1175
	var v1176 int32
	_ = v1176
	var v1178 int32
	_ = v1178
	var v1180 int32
	_ = v1180
	var v1184 int32
	_ = v1184
	var v1188 int32
	_ = v1188
	var v1189 int32
	_ = v1189
	var v1190 int32
	_ = v1190
	var v1191 int32
	_ = v1191
	var v1192 int32
	_ = v1192
	var v1193 int32
	_ = v1193
	var v1194 int32
	_ = v1194
	var v1196 int32
	_ = v1196
	var v1199 int32
	_ = v1199
	var v1200 int32
	_ = v1200
	var v1209 int32
	_ = v1209
	var v1210 int32
	_ = v1210
	var v1212 int32
	_ = v1212
	var v1215 int32
	_ = v1215
	var v1216 int32
	_ = v1216
	var v1221 int32
	_ = v1221
	var v1228 int32
	_ = v1228
	var v1230 int32
	_ = v1230
	var v1232 int32
	_ = v1232
	var v1233 int32
	_ = v1233
	var v1235 int32
	_ = v1235
	var v1237 int32
	_ = v1237
	var v1241 int32
	_ = v1241
	var v1247 int32
	_ = v1247
	var v1248 int32
	_ = v1248
	var v1257 int32
	_ = v1257
	var v1258 int32
	_ = v1258
	var v1260 int32
	_ = v1260
	var v1263 int32
	_ = v1263
	var v1264 int32
	_ = v1264
	var v1269 int32
	_ = v1269
	var v1276 int32
	_ = v1276
	var v1278 int32
	_ = v1278
	var v1280 int32
	_ = v1280
	var v1281 int32
	_ = v1281
	var v1283 int32
	_ = v1283
	var v1285 int32
	_ = v1285
	var v1289 int32
	_ = v1289
	var v1293 int32
	_ = v1293
	var v1294 int32
	_ = v1294
	var v1295 int32
	_ = v1295
	var v1296 int32
	_ = v1296
	var v1297 int32
	_ = v1297
	var v1298 int32
	_ = v1298
	var v1299 int32
	_ = v1299
	var v1302 int32
	_ = v1302
	var v1303 int32
	_ = v1303
	var v1339 int32
	_ = v1339
	var v1340 int32
	_ = v1340
	var v1341 int32
	_ = v1341
	var v1342 int32
	_ = v1342
	var v1348 int32
	_ = v1348
	var v1351 int32
	_ = v1351
	var v1352 int32
	_ = v1352
	var v1353 int32
	_ = v1353
	var v1356 int32
	_ = v1356
	var v1357 int32
	_ = v1357
	var v1358 int32
	_ = v1358
	var v1367 int32
	_ = v1367
	var v1368 int32
	_ = v1368
	var v1370 int32
	_ = v1370
	var v1373 int32
	_ = v1373
	var v1374 int32
	_ = v1374
	var v1379 int32
	_ = v1379
	var v1386 int32
	_ = v1386
	var v1388 int32
	_ = v1388
	var v1390 int32
	_ = v1390
	var v1391 int32
	_ = v1391
	var v1393 int32
	_ = v1393
	var v1395 int32
	_ = v1395
	var v1399 int32
	_ = v1399
	var v1403 int32
	_ = v1403
	var v1406 int32
	_ = v1406
	var v1407 int32
	_ = v1407
	var v1408 int32
	_ = v1408
	var v1417 int32
	_ = v1417
	var v1418 int32
	_ = v1418
	var v1420 int32
	_ = v1420
	var v1423 int32
	_ = v1423
	var v1424 int32
	_ = v1424
	var v1429 int32
	_ = v1429
	var v1436 int32
	_ = v1436
	var v1438 int32
	_ = v1438
	var v1440 int32
	_ = v1440
	var v1441 int32
	_ = v1441
	var v1443 int32
	_ = v1443
	var v1445 int32
	_ = v1445
	var v1449 int32
	_ = v1449
	var v1454 int32
	_ = v1454
	var v1457 int32
	_ = v1457
	var v1458 int32
	_ = v1458
	var v1459 int32
	_ = v1459
	var v1468 int32
	_ = v1468
	var v1469 int32
	_ = v1469
	var v1471 int32
	_ = v1471
	var v1474 int32
	_ = v1474
	var v1475 int32
	_ = v1475
	var v1480 int32
	_ = v1480
	var v1487 int32
	_ = v1487
	var v1489 int32
	_ = v1489
	var v1491 int32
	_ = v1491
	var v1492 int32
	_ = v1492
	var v1494 int32
	_ = v1494
	var v1496 int32
	_ = v1496
	var v1500 int32
	_ = v1500
	var v1504 int32
	_ = v1504
	var v1507 int32
	_ = v1507
	var v1508 int32
	_ = v1508
	var v1509 int32
	_ = v1509
	var v1518 int32
	_ = v1518
	var v1519 int32
	_ = v1519
	var v1521 int32
	_ = v1521
	var v1524 int32
	_ = v1524
	var v1525 int32
	_ = v1525
	var v1530 int32
	_ = v1530
	var v1537 int32
	_ = v1537
	var v1539 int32
	_ = v1539
	var v1541 int32
	_ = v1541
	var v1542 int32
	_ = v1542
	var v1544 int32
	_ = v1544
	var v1546 int32
	_ = v1546
	var v1550 int32
	_ = v1550
	var v1558 int32
	_ = v1558
	var v1559 int32
	_ = v1559
	var v1560 int32
	_ = v1560
	var v1562 int32
	_ = v1562
	var v1563 int32
	_ = v1563
	var v1564 int32
	_ = v1564
	var v1565 int32
	_ = v1565
	var v1566 int32
	_ = v1566
	var v1567 int32
	_ = v1567
	var v1568 int32
	_ = v1568
	var v1569 int32
	_ = v1569
	var v1573 int32
	_ = v1573
	var v1581 int32
	_ = v1581
	var v1584 int32
	_ = v1584
	var v1585 int32
	_ = v1585
	var v1586 int32
	_ = v1586
	var v1587 int32
	_ = v1587
	var v1590 int32
	_ = v1590
	var v1595 int32
	_ = v1595
	var v1598 int32
	_ = v1598
	var v1601 int32
	_ = v1601
	var v1602 int32
	_ = v1602
	var v1604 int32
	_ = v1604
	var v1612 int32
	_ = v1612
	var v1613 int32
	_ = v1613
	var v1616 int32
	_ = v1616
	var v1619 int32
	_ = v1619
	var v1624 int32
	_ = v1624
	var v1633 int32
	_ = v1633
	var v1638 int32
	_ = v1638
	var v1639 int32
	_ = v1639
	var v1640 int32
	_ = v1640
	var v1643 int32
	_ = v1643
	var v1647 int32
	_ = v1647
	var v1648 int32
	_ = v1648
	var v1649 int32
	_ = v1649
	var v1650 int32
	_ = v1650
	var v1651 int32
	_ = v1651
	var v1652 int32
	_ = v1652
	var v1674 int32
	_ = v1674
	var v1679 int32
	_ = v1679
	var v1689 int32
	_ = v1689
	var v1693 int32
	_ = v1693
	var v1694 int32
	_ = v1694
	var v1695 int32
	_ = v1695
	var v1706 int32
	_ = v1706
	var v1731 int32
	_ = v1731
	var v1766 int32
	_ = v1766
	var v1767 int32
	_ = v1767
	var v1778 int32
	_ = v1778
	var v1803 int32
	_ = v1803
	var v1840 int32
	_ = v1840
	var v1841 int32
	_ = v1841
	var v1843 int32
	_ = v1843
	var v1844 int32
	_ = v1844
	var v1857 int32
	_ = v1857
	var v1884 int32
	_ = v1884
	var v1887 int32
	_ = v1887
	var v1923 int32
	_ = v1923
	var v1927 int32
	_ = v1927
	var v1929 int32
	_ = v1929
	var v1938 int32
	_ = v1938
	var v1942 int32
	_ = v1942
	var v1966 int32
	_ = v1966
	var v1970 int32
	_ = v1970
	var v1971 int32
	_ = v1971
	var v1974 int32
	_ = v1974
	var v1975 int32
	_ = v1975
	var v1976 int32
	_ = v1976
	var v1985 int32
	_ = v1985
	var v1986 int32
	_ = v1986
	var v1988 int32
	_ = v1988
	var v1991 int32
	_ = v1991
	var v1992 int32
	_ = v1992
	var v1997 int32
	_ = v1997
	var v2004 int32
	_ = v2004
	var v2006 int32
	_ = v2006
	var v2008 int32
	_ = v2008
	var v2009 int32
	_ = v2009
	var v2011 int32
	_ = v2011
	var v2013 int32
	_ = v2013
	var v2017 int32
	_ = v2017
	var v2024 int32
	_ = v2024
	var v2026 int32
	_ = v2026
	var v2027 int32
	_ = v2027
	var v2035 int32
	_ = v2035
	var v2064 int32
	_ = v2064
	var v2090 int32
	_ = v2090
	var v2106 int32
	_ = v2106
	var v2107 int32
	_ = v2107
	var v2128 int32
	_ = v2128
	var v2143 int32
	_ = v2143
	var v2146 int32
	_ = v2146
	var v2149 int32
	_ = v2149
	var v2150 int32
	_ = v2150
	var v2160 int32
	_ = v2160
	var v2191 int32
	_ = v2191
	var v2192 int32
	_ = v2192
	var v2204 int32
	_ = v2204
	var v2231 int32
	_ = v2231
	var v2234 int32
	_ = v2234
	var v2237 int32
	_ = v2237
	var v2245 int32
	_ = v2245
	var v2274 int32
	_ = v2274
	var v2275 int32
	_ = v2275
	var v2310 int32
	_ = v2310
	var v2311 int32
	_ = v2311
	var v2312 int32
	_ = v2312
	var v2315 int32
	_ = v2315
	var v2318 int32
	_ = v2318
	var v2329 int32
	_ = v2329
	var v2355 int32
	_ = v2355
	var v2359 int32
	_ = v2359
	var v2360 int32
	_ = v2360
	var v2372 int32
	_ = v2372
	var v2397 int32
	_ = v2397
	var v2399 int32
	_ = v2399
	var v2405 int32
	_ = v2405
	var v2442 int32
	_ = v2442
	var v2443 int32
	_ = v2443
	var v2472 int32
	_ = v2472
	var v2480 int32
	_ = v2480
	var v2483 int32
	_ = v2483
	var v2484 int32
	_ = v2484
	var v2485 int32
	_ = v2485
	var v2518 int32
	_ = v2518
	var v2525 int32
	_ = v2525
	var v2527 int32
	_ = v2527
	var v2528 int32
	_ = v2528
	var v2541 int32
	_ = v2541
	var v2543 int32
	_ = v2543
	var v2549 int32
	_ = v2549
	var v2552 int32
	_ = v2552
	var v2568 int32
	_ = v2568
	var v2569 int32
	_ = v2569
	var v2572 int32
	_ = v2572
	var v2574 int32
	_ = v2574
	var v2578 int32
	_ = v2578
	var v2580 int32
	_ = v2580
	var v2584 int32
	_ = v2584
	var v2588 int32
	_ = v2588
	var v2589 int32
	_ = v2589
	var v2590 int32
	_ = v2590
	var v2591 int32
	_ = v2591
	var v2592 int32
	_ = v2592
	var v2593 int32
	_ = v2593
	var v2594 int32
	_ = v2594
	var v2595 int32
	_ = v2595
	var v2596 int32
	_ = v2596
	var v2597 int32
	_ = v2597
	var v2598 int32
	_ = v2598
	var v2599 int32
	_ = v2599
	var v2600 int32
	_ = v2600
	var v2601 int32
	_ = v2601
	var v2602 int32
	_ = v2602
	var v2604 int32
	_ = v2604
	var v2616 int32
	_ = v2616
	var v2622 int32
	_ = v2622
	var v2625 int32
	_ = v2625
	var v2652 int32
	_ = v2652
	var v2658 int32
	_ = v2658
	var v2661 int32
	_ = v2661
	var v2663 int32
	_ = v2663
	var v2679 int32
	_ = v2679
	var v2680 int32
	_ = v2680
	var v2681 int32
	_ = v2681
	var v2682 int32
	_ = v2682
	var v2683 int32
	_ = v2683
	var v2686 int32
	_ = v2686
	var v2704 int32
	_ = v2704
	var v2707 int32
	_ = v2707
	var v2725 int32
	_ = v2725
	var v2727 int32
	_ = v2727
	var v2731 int32
	_ = v2731
	var v2732 int32
	_ = v2732
	var v2733 int32
	_ = v2733
	var v2736 int32
	_ = v2736
	var v2737 int32
	_ = v2737
	var v2738 int32
	_ = v2738
	var v2739 int32
	_ = v2739
	var v2767 int32
	_ = v2767
	var v2775 int32
	_ = v2775
	var v2777 int32
	_ = v2777
	var v2805 int32
	_ = v2805
	var v2814 int32
	_ = v2814
	var v2817 int32
	_ = v2817
	var v2829 int32
	_ = v2829
	var v2859 int32
	_ = v2859
	var v2863 int32
	_ = v2863
	var v2864 int32
	_ = v2864
	var v2865 int32
	_ = v2865
	var v2866 int32
	_ = v2866
	var v2867 int32
	_ = v2867
	var v2868 int32
	_ = v2868
	var v2869 int32
	_ = v2869
	var v2870 int32
	_ = v2870
	var v2871 int32
	_ = v2871
	var v2872 int32
	_ = v2872
	var v2873 int32
	_ = v2873
	var v2874 int32
	_ = v2874
	var v2875 int32
	_ = v2875
	var v2876 int32
	_ = v2876
	var v2877 int32
	_ = v2877
	var v2882 int32
	_ = v2882
	var v2886 int32
	_ = v2886
	var v2888 int32
	_ = v2888
	var v2889 int32
	_ = v2889
	var v2925 int32
	_ = v2925
	var v2926 int32
	_ = v2926
	var v2931 int32
	_ = v2931
	var v2935 int32
	_ = v2935
	var v2940 int32
	_ = v2940
	var v2943 int32
	_ = v2943
	var v2944 int32
	_ = v2944
	var v2945 int32
	_ = v2945
	var v2946 int32
	_ = v2946
	var v2949 int32
	_ = v2949
	var v2950 int32
	_ = v2950
	var v2951 int32
	_ = v2951
	var v2960 int32
	_ = v2960
	var v2961 int32
	_ = v2961
	var v2963 int32
	_ = v2963
	var v2966 int32
	_ = v2966
	var v2967 int32
	_ = v2967
	var v2972 int32
	_ = v2972
	var v2979 int32
	_ = v2979
	var v2981 int32
	_ = v2981
	var v2983 int32
	_ = v2983
	var v2984 int32
	_ = v2984
	var v2986 int32
	_ = v2986
	var v2988 int32
	_ = v2988
	var v2992 int32
	_ = v2992
	var v2996 int32
	_ = v2996
	var v2999 int32
	_ = v2999
	var v3000 int32
	_ = v3000
	var v3001 int32
	_ = v3001
	var v3010 int32
	_ = v3010
	var v3011 int32
	_ = v3011
	var v3013 int32
	_ = v3013
	var v3016 int32
	_ = v3016
	var v3017 int32
	_ = v3017
	var v3022 int32
	_ = v3022
	var v3029 int32
	_ = v3029
	var v3031 int32
	_ = v3031
	var v3033 int32
	_ = v3033
	var v3034 int32
	_ = v3034
	var v3036 int32
	_ = v3036
	var v3038 int32
	_ = v3038
	var v3042 int32
	_ = v3042
	var v3050 int32
	_ = v3050
	var v3051 int32
	_ = v3051
	var v3052 int32
	_ = v3052
	var v3058 int32
	_ = v3058
	var v3059 int32
	_ = v3059
	var v3063 int32
	_ = v3063
	var v3068 int32
	_ = v3068
	var v3070 int32
	_ = v3070
	var v3076 int32
	_ = v3076
	var v3077 int32
	_ = v3077
	var v3079 int32
	_ = v3079
	var v3080 int32
	_ = v3080
	var v3081 int32
	_ = v3081
	var v3082 int32
	_ = v3082
	var v3085 int32
	_ = v3085
	var v3086 int32
	_ = v3086
	var v3091 int32
	_ = v3091
	var v3110 int32
	_ = v3110
	var v3130 int32
	_ = v3130
	var v3134 int32
	_ = v3134
	var v3135 int32
	_ = v3135
	var v3138 int32
	_ = v3138
	var v3139 int32
	_ = v3139
	var v3140 int32
	_ = v3140
	var v3149 int32
	_ = v3149
	var v3150 int32
	_ = v3150
	var v3152 int32
	_ = v3152
	var v3155 int32
	_ = v3155
	var v3156 int32
	_ = v3156
	var v3161 int32
	_ = v3161
	var v3168 int32
	_ = v3168
	var v3170 int32
	_ = v3170
	var v3172 int32
	_ = v3172
	var v3173 int32
	_ = v3173
	var v3175 int32
	_ = v3175
	var v3177 int32
	_ = v3177
	var v3181 int32
	_ = v3181
	var v3185 int32
	_ = v3185
	var v3188 int32
	_ = v3188
	var v3189 int32
	_ = v3189
	var v3190 int32
	_ = v3190
	var v3199 int32
	_ = v3199
	var v3200 int32
	_ = v3200
	var v3202 int32
	_ = v3202
	var v3205 int32
	_ = v3205
	var v3206 int32
	_ = v3206
	var v3211 int32
	_ = v3211
	var v3218 int32
	_ = v3218
	var v3220 int32
	_ = v3220
	var v3222 int32
	_ = v3222
	var v3223 int32
	_ = v3223
	var v3225 int32
	_ = v3225
	var v3227 int32
	_ = v3227
	var v3231 int32
	_ = v3231
	var v3238 int32
	_ = v3238
	var v3240 int32
	_ = v3240
	var v3241 int32
	_ = v3241
	var v3242 int32
	_ = v3242
	var v3243 int32
	_ = v3243
	var v3244 int32
	_ = v3244
	var v3245 int32
	_ = v3245
	var v3246 int32
	_ = v3246
	var v3251 int32
	_ = v3251
	var v3254 int32
	_ = v3254
	var v3255 int32
	_ = v3255
	var v3266 int32
	_ = v3266
	var v3292 int32
	_ = v3292
	var v3296 int32
	_ = v3296
	var v3300 int32
	_ = v3300
	var v3303 int32
	_ = v3303
	var v3304 int32
	_ = v3304
	var v3308 int32
	_ = v3308
	var v3310 int32
	_ = v3310
	var v3311 int32
	_ = v3311
	var v3357 int32
	_ = v3357
	var v3386 int32
	_ = v3386
	var v3425 int32
	_ = v3425
	var v3461 int32
	_ = v3461
	var v3462 int32
	_ = v3462
	var v3498 int32
	_ = v3498
	var v3508 int32
	_ = v3508
	var v3509 int32
	_ = v3509
	var v3512 int32
	_ = v3512
	var v3514 int32
	_ = v3514
	var v3515 int32
	_ = v3515
	var v3518 int32
	_ = v3518
	var v3520 int32
	_ = v3520
	var v3523 int32
	_ = v3523
	var v3524 int32
	_ = v3524
	var v3527 int32
	_ = v3527
	var v3530 int32
	_ = v3530
	var v3531 int32
	_ = v3531
	var v3532 int32
	_ = v3532
	var v3541 int32
	_ = v3541
	var v3542 int32
	_ = v3542
	var v3544 int32
	_ = v3544
	var v3547 int32
	_ = v3547
	var v3548 int32
	_ = v3548
	var v3553 int32
	_ = v3553
	var v3560 int32
	_ = v3560
	var v3562 int32
	_ = v3562
	var v3564 int32
	_ = v3564
	var v3565 int32
	_ = v3565
	var v3567 int32
	_ = v3567
	var v3569 int32
	_ = v3569
	var v3573 int32
	_ = v3573
	var v3577 int32
	_ = v3577
	var v3580 int32
	_ = v3580
	var v3581 int32
	_ = v3581
	var v3582 int32
	_ = v3582
	var v3591 int32
	_ = v3591
	var v3592 int32
	_ = v3592
	var v3594 int32
	_ = v3594
	var v3597 int32
	_ = v3597
	var v3598 int32
	_ = v3598
	var v3603 int32
	_ = v3603
	var v3610 int32
	_ = v3610
	var v3612 int32
	_ = v3612
	var v3614 int32
	_ = v3614
	var v3615 int32
	_ = v3615
	var v3617 int32
	_ = v3617
	var v3619 int32
	_ = v3619
	var v3623 int32
	_ = v3623
	var v3628 int32
	_ = v3628
	var v3630 int32
	_ = v3630
	var v3636 int32
	_ = v3636
	var v3637 int32
	_ = v3637
	var v3638 int32
	_ = v3638
	var v3641 int32
	_ = v3641
	var v3644 int32
	_ = v3644
	var v3648 int32
	_ = v3648
	var v3659 int32
	_ = v3659
	var v3685 int32
	_ = v3685
	var v3689 int32
	_ = v3689
	var v3690 int32
	_ = v3690
	var v3691 int32
	_ = v3691
	var v3692 int32
	_ = v3692
	var v3693 int32
	_ = v3693
	var v3696 int32
	_ = v3696
	var v3697 int32
	_ = v3697
	var v3705 int32
	_ = v3705
	var v3734 int32
	_ = v3734
	var v3738 int32
	_ = v3738
	var v3739 int32
	_ = v3739
	var v3744 int32
	_ = v3744
	var v3746 int32
	_ = v3746
	var v3747 int32
	_ = v3747
	var v3748 int32
	_ = v3748
	var v3749 int32
	_ = v3749
	var v3751 int32
	_ = v3751
	var v3752 int32
	_ = v3752
	var v3753 int32
	_ = v3753
	var v3757 int32
	_ = v3757
	var v3760 int32
	_ = v3760
	var v3761 int32
	_ = v3761
	var v3798 int32
	_ = v3798
	var v3800 int32
	_ = v3800
	var v3801 int32
	_ = v3801
	var v3871 int32
	_ = v3871
	var v3872 int32
	_ = v3872
	var v3877 int32
	_ = v3877
	var v3880 int32
	_ = v3880
	var v3883 int32
	_ = v3883
	var v3884 int32
	_ = v3884
	var v3886 int32
	_ = v3886
	var v3894 int32
	_ = v3894
	var v3895 int32
	_ = v3895
	var v3898 int32
	_ = v3898
	var v3901 int32
	_ = v3901
	var v3906 int32
	_ = v3906
	var v3915 int32
	_ = v3915
	var v3921 int32
	_ = v3921
	var v3924 int32
	_ = v3924
	var v3927 int32
	_ = v3927
	var v3928 int32
	_ = v3928
	var v3936 int32
	_ = v3936
	var v3965 int32
	_ = v3965
	var v3969 int32
	_ = v3969
	var v3971 int32
	_ = v3971
	var v3972 int32
	_ = v3972
	var v3973 int32
	_ = v3973
	var v3976 int32
	_ = v3976
	var v3978 int32
	_ = v3978
	var v3979 int32
	_ = v3979
	var v4052 int32
	_ = v4052
	var v4057 int32
	_ = v4057
	var v4060 int32
	_ = v4060
	var v4061 int32
	_ = v4061
	var v4075 int32
	_ = v4075
	var v4084 int32
	_ = v4084
	var v4103 int32
	_ = v4103
	var v4107 int32
	_ = v4107
	var v4108 int32
	_ = v4108
	var v4109 int32
	_ = v4109
	var v4110 int32
	_ = v4110
	var v4111 int32
	_ = v4111
	var v4120 int32
	_ = v4120
	var v4121 int32
	_ = v4121
	var v4123 int32
	_ = v4123
	var v4126 int32
	_ = v4126
	var v4127 int32
	_ = v4127
	var v4132 int32
	_ = v4132
	var v4139 int32
	_ = v4139
	var v4141 int32
	_ = v4141
	var v4143 int32
	_ = v4143
	var v4146 int32
	_ = v4146
	var v4148 int32
	_ = v4148
	var v4150 int32
	_ = v4150
	var v4155 int32
	_ = v4155
	var v4164 int32
	_ = v4164
	var v4167 int32
	_ = v4167
	var v4170 int32
	_ = v4170
	var v4173 int32
	_ = v4173
	var v4174 int32
	_ = v4174
	var v4175 int32
	_ = v4175
	var v4176 int32
	_ = v4176
	var v4185 int32
	_ = v4185
	var v4186 int32
	_ = v4186
	var v4188 int32
	_ = v4188
	var v4191 int32
	_ = v4191
	var v4192 int32
	_ = v4192
	var v4197 int32
	_ = v4197
	var v4204 int32
	_ = v4204
	var v4206 int32
	_ = v4206
	var v4208 int32
	_ = v4208
	var v4211 int32
	_ = v4211
	var v4213 int32
	_ = v4213
	var v4215 int32
	_ = v4215
	var v4220 int32
	_ = v4220
	var v4229 int32
	_ = v4229
	var v4232 int32
	_ = v4232
	var v4233 int32
	_ = v4233
	var v4242 int32
	_ = v4242
	var v4243 int32
	_ = v4243
	var v4245 int32
	_ = v4245
	var v4248 int32
	_ = v4248
	var v4249 int32
	_ = v4249
	var v4254 int32
	_ = v4254
	var v4261 int32
	_ = v4261
	var v4263 int32
	_ = v4263
	var v4265 int32
	_ = v4265
	var v4268 int32
	_ = v4268
	var v4270 int32
	_ = v4270
	var v4272 int32
	_ = v4272
	var v4277 int32
	_ = v4277
	var v4286 int32
	_ = v4286
	var v4289 int32
	_ = v4289
	var v4291 int32
	_ = v4291
	var v4292 int32
	_ = v4292
	var v4301 int32
	_ = v4301
	var v4302 int32
	_ = v4302
	var v4304 int32
	_ = v4304
	var v4307 int32
	_ = v4307
	var v4308 int32
	_ = v4308
	var v4313 int32
	_ = v4313
	var v4320 int32
	_ = v4320
	var v4322 int32
	_ = v4322
	var v4324 int32
	_ = v4324
	var v4327 int32
	_ = v4327
	var v4329 int32
	_ = v4329
	var v4331 int32
	_ = v4331
	var v4336 int32
	_ = v4336
	var v4345 int32
	_ = v4345
	var v4348 int32
	_ = v4348
	var v4349 int32
	_ = v4349
	var v4358 int32
	_ = v4358
	var v4359 int32
	_ = v4359
	var v4361 int32
	_ = v4361
	var v4364 int32
	_ = v4364
	var v4365 int32
	_ = v4365
	var v4370 int32
	_ = v4370
	var v4377 int32
	_ = v4377
	var v4379 int32
	_ = v4379
	var v4381 int32
	_ = v4381
	var v4384 int32
	_ = v4384
	var v4386 int32
	_ = v4386
	var v4388 int32
	_ = v4388
	var v4393 int32
	_ = v4393
	var v4402 int32
	_ = v4402
	var v4405 int32
	_ = v4405
	var v4407 int32
	_ = v4407
	var v4408 int32
	_ = v4408
	var v4409 int32
	_ = v4409
	var v4410 int32
	_ = v4410
	var v4413 int32
	_ = v4413
	var v4414 int32
	_ = v4414
	var v4417 int32
	_ = v4417
	var v4419 int32
	_ = v4419
	var v4420 int32
	_ = v4420
	var v4424 int32
	_ = v4424
	var v4425 int32
	_ = v4425
	var v4426 int32
	_ = v4426
	var v4427 int32
	_ = v4427
	var v4430 int32
	_ = v4430
	var v4431 int32
	_ = v4431
	var v4432 int32
	_ = v4432
	var v4441 int32
	_ = v4441
	var v4442 int32
	_ = v4442
	var v4444 int32
	_ = v4444
	var v4447 int32
	_ = v4447
	var v4448 int32
	_ = v4448
	var v4453 int32
	_ = v4453
	var v4460 int32
	_ = v4460
	var v4462 int32
	_ = v4462
	var v4464 int32
	_ = v4464
	var v4465 int32
	_ = v4465
	var v4467 int32
	_ = v4467
	var v4469 int32
	_ = v4469
	var v4473 int32
	_ = v4473
	var v4477 int32
	_ = v4477
	var v4480 int32
	_ = v4480
	var v4481 int32
	_ = v4481
	var v4482 int32
	_ = v4482
	var v4491 int32
	_ = v4491
	var v4492 int32
	_ = v4492
	var v4494 int32
	_ = v4494
	var v4497 int32
	_ = v4497
	var v4498 int32
	_ = v4498
	var v4503 int32
	_ = v4503
	var v4510 int32
	_ = v4510
	var v4512 int32
	_ = v4512
	var v4514 int32
	_ = v4514
	var v4515 int32
	_ = v4515
	var v4517 int32
	_ = v4517
	var v4519 int32
	_ = v4519
	var v4523 int32
	_ = v4523
	var v4528 int32
	_ = v4528
	var v4531 int32
	_ = v4531
	var v4532 int32
	_ = v4532
	var v4533 int32
	_ = v4533
	var v4542 int32
	_ = v4542
	var v4543 int32
	_ = v4543
	var v4545 int32
	_ = v4545
	var v4548 int32
	_ = v4548
	var v4549 int32
	_ = v4549
	var v4554 int32
	_ = v4554
	var v4561 int32
	_ = v4561
	var v4563 int32
	_ = v4563
	var v4565 int32
	_ = v4565
	var v4566 int32
	_ = v4566
	var v4568 int32
	_ = v4568
	var v4570 int32
	_ = v4570
	var v4574 int32
	_ = v4574
	var v4578 int32
	_ = v4578
	var v4581 int32
	_ = v4581
	var v4582 int32
	_ = v4582
	var v4583 int32
	_ = v4583
	var v4592 int32
	_ = v4592
	var v4593 int32
	_ = v4593
	var v4595 int32
	_ = v4595
	var v4598 int32
	_ = v4598
	var v4599 int32
	_ = v4599
	var v4604 int32
	_ = v4604
	var v4611 int32
	_ = v4611
	var v4613 int32
	_ = v4613
	var v4615 int32
	_ = v4615
	var v4616 int32
	_ = v4616
	var v4618 int32
	_ = v4618
	var v4620 int32
	_ = v4620
	var v4624 int32
	_ = v4624
	var v4632 int32
	_ = v4632
	var v4633 int32
	_ = v4633
	var v4634 int32
	_ = v4634
	var v4639 int32
	_ = v4639
	var v4640 int32
	_ = v4640
	var v4641 int32
	_ = v4641
	var v4642 int32
	_ = v4642
	var v4643 int32
	_ = v4643
	var v4648 int32
	_ = v4648
	var v4656 int32
	_ = v4656
	var v4660 int32
	_ = v4660
	var v4661 int32
	_ = v4661
	var v4664 int32
	_ = v4664
	var v4679 int32
	_ = v4679
	var v4702 int32
	_ = v4702
	var v4706 int32
	_ = v4706
	var v4707 int32
	_ = v4707
	var v4710 int32
	_ = v4710
	var v4711 int32
	_ = v4711
	var v4712 int32
	_ = v4712
	var v4721 int32
	_ = v4721
	var v4722 int32
	_ = v4722
	var v4724 int32
	_ = v4724
	var v4727 int32
	_ = v4727
	var v4728 int32
	_ = v4728
	var v4733 int32
	_ = v4733
	var v4740 int32
	_ = v4740
	var v4742 int32
	_ = v4742
	var v4744 int32
	_ = v4744
	var v4745 int32
	_ = v4745
	var v4747 int32
	_ = v4747
	var v4749 int32
	_ = v4749
	var v4753 int32
	_ = v4753
	var v4757 int32
	_ = v4757
	var v4760 int32
	_ = v4760
	var v4761 int32
	_ = v4761
	var v4762 int32
	_ = v4762
	var v4771 int32
	_ = v4771
	var v4772 int32
	_ = v4772
	var v4774 int32
	_ = v4774
	var v4777 int32
	_ = v4777
	var v4778 int32
	_ = v4778
	var v4783 int32
	_ = v4783
	var v4790 int32
	_ = v4790
	var v4792 int32
	_ = v4792
	var v4794 int32
	_ = v4794
	var v4795 int32
	_ = v4795
	var v4797 int32
	_ = v4797
	var v4799 int32
	_ = v4799
	var v4803 int32
	_ = v4803
	var v4808 int32
	_ = v4808
	var v4811 int32
	_ = v4811
	var v4812 int32
	_ = v4812
	var v4822 int32
	_ = v4822
	var v4849 int32
	_ = v4849
	var v4853 int32
	_ = v4853
	var v4854 int32
	_ = v4854
	var v4857 int32
	_ = v4857
	var v4858 int32
	_ = v4858
	var v4859 int32
	_ = v4859
	var v4868 int32
	_ = v4868
	var v4869 int32
	_ = v4869
	var v4871 int32
	_ = v4871
	var v4874 int32
	_ = v4874
	var v4875 int32
	_ = v4875
	var v4880 int32
	_ = v4880
	var v4887 int32
	_ = v4887
	var v4889 int32
	_ = v4889
	var v4891 int32
	_ = v4891
	var v4892 int32
	_ = v4892
	var v4894 int32
	_ = v4894
	var v4896 int32
	_ = v4896
	var v4900 int32
	_ = v4900
	var v4904 int32
	_ = v4904
	var v4907 int32
	_ = v4907
	var v4908 int32
	_ = v4908
	var v4909 int32
	_ = v4909
	var v4918 int32
	_ = v4918
	var v4919 int32
	_ = v4919
	var v4921 int32
	_ = v4921
	var v4924 int32
	_ = v4924
	var v4925 int32
	_ = v4925
	var v4930 int32
	_ = v4930
	var v4937 int32
	_ = v4937
	var v4939 int32
	_ = v4939
	var v4941 int32
	_ = v4941
	var v4942 int32
	_ = v4942
	var v4944 int32
	_ = v4944
	var v4946 int32
	_ = v4946
	var v4950 int32
	_ = v4950
	var v4966 int32
	_ = v4966
	var v4969 int32
	_ = v4969
	var v4970 int32
	_ = v4970
	var v5007 int32
	_ = v5007
	var v5008 int32
	_ = v5008
	var v5050 int32
	_ = v5050
	var v5057 int32
	_ = v5057
	var v5078 int32
	_ = v5078
	var v5083 int32
	_ = v5083
	var v5086 int32
	_ = v5086
	var v5087 int32
	_ = v5087
	var v5088 int32
	_ = v5088
	var v5091 int32
	_ = v5091
	var v5095 int32
	_ = v5095
	var v5098 int32
	_ = v5098
	var v5099 int32
	_ = v5099
	var v5104 int32
	_ = v5104
	var v5113 int32
	_ = v5113
	var v5115 int32
	_ = v5115
	var v5116 int32
	_ = v5116
	var v5119 int32
	_ = v5119
	var v5124 int32
	_ = v5124
	var v5127 int32
	_ = v5127
	var v5130 int32
	_ = v5130
	var v5131 int32
	_ = v5131
	var v5133 int32
	_ = v5133
	var v5141 int32
	_ = v5141
	var v5142 int32
	_ = v5142
	var v5145 int32
	_ = v5145
	var v5148 int32
	_ = v5148
	var v5153 int32
	_ = v5153
	var v5162 int32
	_ = v5162
	var v5166 int32
	_ = v5166
	var v5173 int32
	_ = v5173
	var v5208 int32
	_ = v5208
	var v5211 int32
	_ = v5211
	var v5217 int32
	_ = v5217
	var v5220 int32
	_ = v5220
	var v5224 int32
	_ = v5224
	v8 = int32(0)
	v35 = m.G0
	v37 = v35 - int32(48)
	m.G0 = v37
	v41 = *(*int32)(unsafe.Add(mBase, uint32(l1)+4))
	if v41 == int32(3) {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	v44 = int32(228)
	goto L3
L2:
	;
	v44 = int32(8)
	goto L3
L3:
	;
	v46 = *(*int32)(unsafe.Add(mBase, uint32(l1+v44)))
	v47 = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v37)+40)) = v47
	*(*int32)(unsafe.Add(mBase, uint32(v37)+20)) = l5
	*(*int32)(unsafe.Add(mBase, uint32(v37)+12)) = v47
	*(*int32)(unsafe.Add(mBase, uint32(v37)+8)) = l6
	switch l4 - int32(4) {
	case 0, 1:
		v120 = v8
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
	*(*uint8)(unsafe.Add(mBase, uint32(v37)+16)) = uint8(v120)
	if l4 != int32(2) {
		goto L26
	} else {
		goto L27
	}
L5:
	;
	v116 = *(*int32)(unsafe.Add(mBase, uint32(l1)+8))
	v117 = *(*int32)(unsafe.Add(mBase, uint32(l2)+8))
	v118 = F_innerrel_is_unique(m, l0, v116, v117, l3, l4, l6)
	mBase = m.M
	v119 = m.ExcPending
	if v119 != 0 {
		goto L22
	} else {
		goto L24
	}
L6:
	;
	v111 = *(*int32)(unsafe.Add(mBase, uint32(l1)+8))
	v112 = *(*int32)(unsafe.Add(mBase, uint32(l2)+8))
	v114 = F_innerrel_is_unique(m, l0, v111, v112, l3, int32(0), l6)
	mBase = m.M
	v115 = m.ExcPending
	if v115 != 0 {
		goto L22
	} else {
		goto L23
	}
L7:
	;
	v55 = *(*int32)(unsafe.Add(mBase, uint32(l5)+4))
	v56 = *(*int32)(unsafe.Add(mBase, uint32(l2)+8))
	v57 = int32(0)
	if v55 == v57 {
		goto L9
	} else {
		goto L10
	}
L8:
	;
	v120 = v110
	goto L4
L9:
	;
	v110 = int32(1)
	goto L8
L10:
	;
	goto L11
L11:
	;
	if v56 == int32(0) {
		v101 = v57
		goto L12
	} else {
		goto L13
	}
L12:
	;
	v110 = v101
	goto L8
L13:
	;
	v66 = *(*int32)(unsafe.Add(mBase, uint32(v55)+4))
	v67 = *(*int32)(unsafe.Add(mBase, uint32(v56)+4))
	if v67 < v66 {
		v101 = v57
		goto L12
	} else {
		goto L14
	}
L14:
	;
	v69 = int32(1)
	if v66 <= v69 {
		goto L15
	} else {
		goto L16
	}
L15:
	;
	v72 = v69
	goto L17
L16:
	;
	v72 = v66
	goto L17
L17:
	;
	v73 = int32(8)
	v78 = int32(0)
	goto L18
L18:
	;
	v85 = v78 << (uint(int32(2)) % 32)
	v87 = *(*int32)(unsafe.Add(mBase, uint32(v55+v73+v85)))
	v89 = *(*int32)(unsafe.Add(mBase, uint32(v85+(v56+v73))))
	v92 = v87 & (v89 ^ int32(-1))
	v94 = base.B2i32(v92 == int32(0))
	if v92 != 0 {
		v101 = v94
		goto L12
	} else {
		goto L20
	}
L19:
	;
	v101 = v94
	goto L12
L20:
	;
	v96 = v78 + int32(1)
	if v96 != v72 {
		v78 = v96
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
	v120 = v114
	goto L4
L24:
	;
	v120 = v118
	goto L4
L25:
	;
	if (base.B2i32(l4&int32(-2) == int32(4))|v120)&int32(1) != 0 {
		goto L158
	} else {
		goto L159
	}
L26:
	;
	v124 = int32(1)
	v126 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_add_paths_to_joinrel[0])))
	if v126&v124 == int32(0) {
		v778 = v124
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
		v741 = v8
		v743 = int32(0)
		goto L30
	} else {
		goto L31
	}
L29:
	;
	goto L28
L30:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v37)+12)) = v741
	v778 = v743
	goto L25
L31:
	;
	v135 = int32(1)
	if l6 == int32(0) {
		v698 = v8
		v700 = v135
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
	v138 = *(*int32)(unsafe.Add(mBase, uint32(l6)+4))
	if v138 <= int32(0) {
		v698 = v8
		v700 = v135
		goto L32
	} else {
		goto L34
	}
L34:
	;
	v156 = v8
	v158 = v8
	v159 = v8
	goto L35
L35:
	;
	v179 = *(*int32)(unsafe.Add(mBase, uint32(l6)+12))
	v183 = *(*int32)(unsafe.Add(mBase, uint32(v179+v159<<(uint(int32(2))%32))))
	if int32(1)<<(uint(l4)%32)&int32(174) != 0 {
		goto L38
	} else {
		goto L39
	}
L36:
	;
	v698 = v658
	v700 = v656 ^ int32(1)
	goto L32
L37:
	;
	v680 = v159 + int32(1)
	v681 = *(*int32)(unsafe.Add(mBase, uint32(l6)+4))
	if v680 < v681 {
		v156 = v656
		v158 = v658
		v159 = v680
		goto L35
	} else {
		goto L153
	}
L38:
	;
	v184 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v183)+8)))
	if v184 != 0 {
		v656 = v156
		v658 = v158
		goto L37
	} else {
		goto L41
	}
L39:
	;
	goto L40
L40:
	;
	v243 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v183)+9)))
	if v243 == int32(1) {
		goto L58
	} else {
		goto L59
	}
L41:
	;
	v185 = *(*int32)(unsafe.Add(mBase, uint32(v183)+32))
	v186 = *(*int32)(unsafe.Add(mBase, uint32(l1)+8))
	v187 = int32(0)
	if v185 == v187 {
		goto L43
	} else {
		goto L44
	}
L42:
	;
	if v240 == int32(0) {
		v656 = v156
		v658 = v158
		goto L37
	} else {
		goto L56
	}
L43:
	;
	v240 = int32(1)
	goto L42
L44:
	;
	goto L45
L45:
	;
	if v186 == int32(0) {
		v231 = v187
		goto L46
	} else {
		goto L47
	}
L46:
	;
	v240 = v231
	goto L42
L47:
	;
	v196 = *(*int32)(unsafe.Add(mBase, uint32(v185)+4))
	v197 = *(*int32)(unsafe.Add(mBase, uint32(v186)+4))
	if v197 < v196 {
		v231 = v187
		goto L46
	} else {
		goto L48
	}
L48:
	;
	v199 = int32(1)
	if v196 <= v199 {
		goto L49
	} else {
		goto L50
	}
L49:
	;
	v202 = v199
	goto L51
L50:
	;
	v202 = v196
	goto L51
L51:
	;
	v203 = int32(8)
	v208 = int32(0)
	goto L52
L52:
	;
	v215 = v208 << (uint(int32(2)) % 32)
	v217 = *(*int32)(unsafe.Add(mBase, uint32(v185+v203+v215)))
	v219 = *(*int32)(unsafe.Add(mBase, uint32(v215+(v186+v203))))
	v222 = v217 & (v219 ^ int32(-1))
	v224 = base.B2i32(v222 == int32(0))
	if v222 != 0 {
		v231 = v224
		goto L46
	} else {
		goto L54
	}
L53:
	;
	v231 = v224
	goto L46
L54:
	;
	v226 = v208 + int32(1)
	if v226 != v202 {
		v208 = v226
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
	v252 = *(*int32)(unsafe.Add(mBase, uint32(l3)+8))
	v253 = *(*int32)(unsafe.Add(mBase, uint32(v183)+44))
	v254 = *(*int32)(unsafe.Add(mBase, uint32(l2)+8))
	v255 = int32(0)
	if v253 == v255 {
		goto L69
	} else {
		goto L70
	}
L58:
	;
	v246 = *(*int32)(unsafe.Add(mBase, uint32(v183)+96))
	if v246 != 0 {
		goto L57
	} else {
		goto L61
	}
L59:
	;
	goto L60
L60:
	;
	v247 = *(*int32)(unsafe.Add(mBase, uint32(v183)+4))
	if v247 != 0 {
		goto L62
	} else {
		goto L63
	}
L61:
	;
	goto L60
L62:
	;
	v248 = *(*int32)(unsafe.Add(mBase, uint32(v247)))
	if v248 == int32(7) {
		v656 = v156
		v658 = v158
		goto L37
	} else {
		goto L65
	}
L63:
	;
	goto L64
L64:
	;
	v656 = int32(1)
	v658 = v158
	goto L37
L65:
	;
	goto L64
L66:
	;
	v493 = *(*int32)(unsafe.Add(mBase, uint32(v183)+100))
	v494 = *(*int32)(unsafe.Add(mBase, uint32(v493)+56))
	if v494 != 0 {
		goto L134
	} else {
		goto L135
	}
L67:
	;
	v370 = *(*int32)(unsafe.Add(mBase, uint32(v183)+44))
	v371 = int32(0)
	if v370 == v371 {
		goto L99
	} else {
		goto L100
	}
L68:
	;
	if v308 == int32(0) {
		goto L67
	} else {
		goto L82
	}
L69:
	;
	v308 = int32(1)
	goto L68
L70:
	;
	goto L71
L71:
	;
	if v254 == int32(0) {
		v299 = v255
		goto L72
	} else {
		goto L73
	}
L72:
	;
	v308 = v299
	goto L68
L73:
	;
	v264 = *(*int32)(unsafe.Add(mBase, uint32(v253)+4))
	v265 = *(*int32)(unsafe.Add(mBase, uint32(v254)+4))
	if v265 < v264 {
		v299 = v255
		goto L72
	} else {
		goto L74
	}
L74:
	;
	v267 = int32(1)
	if v264 <= v267 {
		goto L75
	} else {
		goto L76
	}
L75:
	;
	v270 = v267
	goto L77
L76:
	;
	v270 = v264
	goto L77
L77:
	;
	v271 = int32(8)
	v276 = int32(0)
	goto L78
L78:
	;
	v283 = v276 << (uint(int32(2)) % 32)
	v285 = *(*int32)(unsafe.Add(mBase, uint32(v253+v271+v283)))
	v287 = *(*int32)(unsafe.Add(mBase, uint32(v283+(v254+v271))))
	v290 = v285 & (v287 ^ int32(-1))
	v292 = base.B2i32(v290 == int32(0))
	if v290 != 0 {
		v299 = v292
		goto L72
	} else {
		goto L80
	}
L79:
	;
	v299 = v292
	goto L72
L80:
	;
	v294 = v276 + int32(1)
	if v294 != v270 {
		v276 = v294
		goto L78
	} else {
		goto L81
	}
L81:
	;
	goto L79
L82:
	;
	v311 = *(*int32)(unsafe.Add(mBase, uint32(v183)+48))
	v312 = int32(0)
	if v311 == v312 {
		goto L84
	} else {
		goto L85
	}
L83:
	;
	if v365 == int32(0) {
		goto L67
	} else {
		goto L97
	}
L84:
	;
	v365 = int32(1)
	goto L83
L85:
	;
	goto L86
L86:
	;
	if v252 == int32(0) {
		v356 = v312
		goto L87
	} else {
		goto L88
	}
L87:
	;
	v365 = v356
	goto L83
L88:
	;
	v321 = *(*int32)(unsafe.Add(mBase, uint32(v311)+4))
	v322 = *(*int32)(unsafe.Add(mBase, uint32(v252)+4))
	if v322 < v321 {
		v356 = v312
		goto L87
	} else {
		goto L89
	}
L89:
	;
	v324 = int32(1)
	if v321 <= v324 {
		goto L90
	} else {
		goto L91
	}
L90:
	;
	v327 = v324
	goto L92
L91:
	;
	v327 = v321
	goto L92
L92:
	;
	v328 = int32(8)
	v333 = int32(0)
	goto L93
L93:
	;
	v340 = v333 << (uint(int32(2)) % 32)
	v342 = *(*int32)(unsafe.Add(mBase, uint32(v311+v328+v340)))
	v344 = *(*int32)(unsafe.Add(mBase, uint32(v340+(v252+v328))))
	v347 = v342 & (v344 ^ int32(-1))
	v349 = base.B2i32(v347 == int32(0))
	if v347 != 0 {
		v356 = v349
		goto L87
	} else {
		goto L95
	}
L94:
	;
	v356 = v349
	goto L87
L95:
	;
	v351 = v333 + int32(1)
	if v351 != v327 {
		v333 = v351
		goto L93
	} else {
		goto L96
	}
L96:
	;
	goto L94
L97:
	;
	v368 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(v183)+120)) = uint8(v368)
	goto L66
L98:
	;
	if v424 == int32(0) {
		goto L112
	} else {
		goto L113
	}
L99:
	;
	v424 = int32(1)
	goto L98
L100:
	;
	goto L101
L101:
	;
	if v252 == int32(0) {
		v415 = v371
		goto L102
	} else {
		goto L103
	}
L102:
	;
	v424 = v415
	goto L98
L103:
	;
	v380 = *(*int32)(unsafe.Add(mBase, uint32(v370)+4))
	v381 = *(*int32)(unsafe.Add(mBase, uint32(v252)+4))
	if v381 < v380 {
		v415 = v371
		goto L102
	} else {
		goto L104
	}
L104:
	;
	v383 = int32(1)
	if v380 <= v383 {
		goto L105
	} else {
		goto L106
	}
L105:
	;
	v386 = v383
	goto L107
L106:
	;
	v386 = v380
	goto L107
L107:
	;
	v387 = int32(8)
	v392 = int32(0)
	goto L108
L108:
	;
	v399 = v392 << (uint(int32(2)) % 32)
	v401 = *(*int32)(unsafe.Add(mBase, uint32(v370+v387+v399)))
	v403 = *(*int32)(unsafe.Add(mBase, uint32(v399+(v252+v387))))
	v406 = v401 & (v403 ^ int32(-1))
	v408 = base.B2i32(v406 == int32(0))
	if v406 != 0 {
		v415 = v408
		goto L102
	} else {
		goto L110
	}
L109:
	;
	v415 = v408
	goto L102
L110:
	;
	v410 = v392 + int32(1)
	if v410 != v386 {
		v392 = v410
		goto L108
	} else {
		goto L111
	}
L111:
	;
	goto L109
L112:
	;
	v656 = int32(1)
	v658 = v158
	goto L37
L113:
	;
	goto L114
L114:
	;
	v428 = *(*int32)(unsafe.Add(mBase, uint32(v183)+48))
	v429 = int32(0)
	if v428 == v429 {
		goto L116
	} else {
		goto L117
	}
L115:
	;
	if v482 == int32(0) {
		goto L129
	} else {
		goto L130
	}
L116:
	;
	v482 = int32(1)
	goto L115
L117:
	;
	goto L118
L118:
	;
	if v254 == int32(0) {
		v473 = v429
		goto L119
	} else {
		goto L120
	}
L119:
	;
	v482 = v473
	goto L115
L120:
	;
	v438 = *(*int32)(unsafe.Add(mBase, uint32(v428)+4))
	v439 = *(*int32)(unsafe.Add(mBase, uint32(v254)+4))
	if v439 < v438 {
		v473 = v429
		goto L119
	} else {
		goto L121
	}
L121:
	;
	v441 = int32(1)
	if v438 <= v441 {
		goto L122
	} else {
		goto L123
	}
L122:
	;
	v444 = v441
	goto L124
L123:
	;
	v444 = v438
	goto L124
L124:
	;
	v445 = int32(8)
	v450 = int32(0)
	goto L125
L125:
	;
	v457 = v450 << (uint(int32(2)) % 32)
	v459 = *(*int32)(unsafe.Add(mBase, uint32(v428+v445+v457)))
	v461 = *(*int32)(unsafe.Add(mBase, uint32(v457+(v254+v445))))
	v464 = v459 & (v461 ^ int32(-1))
	v466 = base.B2i32(v464 == int32(0))
	if v464 != 0 {
		v473 = v466
		goto L119
	} else {
		goto L127
	}
L126:
	;
	v473 = v466
	goto L119
L127:
	;
	v468 = v450 + int32(1)
	if v468 != v444 {
		v450 = v468
		goto L125
	} else {
		goto L128
	}
L128:
	;
	goto L126
L129:
	;
	v656 = int32(1)
	v658 = v158
	goto L37
L130:
	;
	goto L131
L131:
	;
	v486 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v183)+120)) = uint8(v486)
	v488 = *(*int32)(unsafe.Add(mBase, uint32(v183)+4))
	v489 = *(*int32)(unsafe.Add(mBase, uint32(v488)+4))
	v490 = F_get_commutator(m, v489)
	mBase = m.M
	v491 = m.ExcPending
	if v491 != 0 {
		goto L22
	} else {
		goto L132
	}
L132:
	;
	if v490 != 0 {
		goto L66
	} else {
		goto L133
	}
L133:
	;
	v656 = int32(1)
	v658 = v158
	goto L37
L134:
	;
	v513 = v494
	goto L137
L135:
	;
	goto L136
L136:
	;
	v565 = *(*int32)(unsafe.Add(mBase, uint32(v183)+104))
	v566 = *(*int32)(unsafe.Add(mBase, uint32(v565)+56))
	if v566 != 0 {
		goto L140
	} else {
		goto L141
	}
L137:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v183)+100)) = v513
	v530 = *(*int32)(unsafe.Add(mBase, uint32(v513)+56))
	if v530 != 0 {
		v513 = v530
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
	v585 = v566
	goto L143
L141:
	;
	goto L142
L142:
	;
	v637 = *(*int32)(unsafe.Add(mBase, uint32(v183)+100))
	v638 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v637)+40)))
	if v638 != 0 {
		goto L146
	} else {
		goto L147
	}
L143:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v183)+104)) = v585
	v602 = *(*int32)(unsafe.Add(mBase, uint32(v585)+56))
	if v602 != 0 {
		v585 = v602
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
	v656 = int32(1)
	v658 = v158
	goto L37
L147:
	;
	goto L148
L148:
	;
	v640 = *(*int32)(unsafe.Add(mBase, uint32(v183)+104))
	v641 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v640)+40)))
	if v641 != 0 {
		goto L149
	} else {
		goto L150
	}
L149:
	;
	v656 = int32(1)
	v658 = v158
	goto L37
L150:
	;
	goto L151
L151:
	;
	v643 = F_lappend(m, v158, v183)
	mBase = m.M
	v644 = m.ExcPending
	if v644 != 0 {
		goto L22
	} else {
		goto L152
	}
L152:
	;
	v656 = v156
	v658 = v643
	goto L37
L153:
	;
	goto L36
L154:
	;
	v726 = base.B2i32(base.Ui32(l4) <= base.Ui32(int32(7)))
	goto L156
L155:
	;
	v726 = int32(0)
	goto L156
L156:
	;
	if v726 != 0 {
		v741 = v698
		v743 = v700
		goto L30
	} else {
		goto L157
	}
L157:
	;
	v741 = v698
	v743 = int32(1)
	goto L30
L158:
	;
	v805 = v37 + int32(24)
	v806 = int32(0)
	v807 = m.G0
	v809 = v807 + int32(-64)
	m.G0 = v809
	v814 = int32(1) << (uint(l4) % 32) & int32(174)
	if v814 == v806 {
		goto L162
	} else {
		goto L163
	}
L159:
	;
	goto L160
L160:
	;
	v1048 = *(*int32)(unsafe.Add(mBase, uint32(l0)+112))
	if v1048 == int32(0) {
		goto L204
	} else {
		goto L205
	}
L161:
	;
	v962 = int32(5)
	if l4 == v962 {
		goto L188
	} else {
		goto L189
	}
L162:
	;
	v938 = l6
	goto L161
L163:
	;
	goto L164
L164:
	;
	if l6 == int32(0) {
		v938 = v806
		goto L161
	} else {
		goto L165
	}
L165:
	;
	v819 = *(*int32)(unsafe.Add(mBase, uint32(l6)+4))
	if v819 <= int32(0) {
		v938 = v806
		goto L161
	} else {
		goto L166
	}
L166:
	;
	v833 = v806
	v842 = v8
	goto L167
L167:
	;
	v856 = *(*int32)(unsafe.Add(mBase, uint32(l6)+12))
	v860 = *(*int32)(unsafe.Add(mBase, uint32(v856+v842<<(uint(int32(2))%32))))
	v861 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v860)+8)))
	if v861 != 0 {
		v922 = v833
		goto L169
	} else {
		goto L170
	}
L168:
	;
	v938 = v922
	goto L161
L169:
	;
	v924 = v842 + int32(1)
	v925 = *(*int32)(unsafe.Add(mBase, uint32(l6)+4))
	if v924 < v925 {
		v833 = v922
		v842 = v924
		goto L167
	} else {
		goto L187
	}
L170:
	;
	v862 = *(*int32)(unsafe.Add(mBase, uint32(v860)+32))
	v863 = *(*int32)(unsafe.Add(mBase, uint32(l1)+8))
	v864 = int32(0)
	if v862 == v864 {
		goto L172
	} else {
		goto L173
	}
L171:
	;
	if v917 == int32(0) {
		v922 = v833
		goto L169
	} else {
		goto L185
	}
L172:
	;
	v917 = int32(1)
	goto L171
L173:
	;
	goto L174
L174:
	;
	if v863 == int32(0) {
		v908 = v864
		goto L175
	} else {
		goto L176
	}
L175:
	;
	v917 = v908
	goto L171
L176:
	;
	v873 = *(*int32)(unsafe.Add(mBase, uint32(v862)+4))
	v874 = *(*int32)(unsafe.Add(mBase, uint32(v863)+4))
	if v874 < v873 {
		v908 = v864
		goto L175
	} else {
		goto L177
	}
L177:
	;
	v876 = int32(1)
	if v873 <= v876 {
		goto L178
	} else {
		goto L179
	}
L178:
	;
	v879 = v876
	goto L180
L179:
	;
	v879 = v873
	goto L180
L180:
	;
	v880 = int32(8)
	v885 = int32(0)
	goto L181
L181:
	;
	v892 = v885 << (uint(int32(2)) % 32)
	v894 = *(*int32)(unsafe.Add(mBase, uint32(v862+v880+v892)))
	v896 = *(*int32)(unsafe.Add(mBase, uint32(v892+(v863+v880))))
	v899 = v894 & (v896 ^ int32(-1))
	v901 = base.B2i32(v899 == int32(0))
	if v899 != 0 {
		v908 = v901
		goto L175
	} else {
		goto L183
	}
L182:
	;
	v908 = v901
	goto L175
L183:
	;
	v903 = v885 + int32(1)
	if v903 != v879 {
		v885 = v903
		goto L181
	} else {
		goto L184
	}
L184:
	;
	goto L182
L185:
	;
	v920 = F_lappend(m, v833, v860)
	mBase = m.M
	v921 = m.ExcPending
	if v921 != 0 {
		goto L22
	} else {
		goto L186
	}
L186:
	;
	v922 = v920
	goto L169
L187:
	;
	goto L168
L188:
	;
	v966 = v962
	goto L190
L189:
	;
	v966 = int32(4)
	goto L190
L190:
	;
	v967 = F_clauselist_selectivity(m, l0, v938, int32(0), v966, l5)
	mBase = m.M
	v968 = m.ExcPending
	if v968 != 0 {
		goto L22
	} else {
		goto L191
	}
L191:
	;
	v970 = v807 + int32(-56)
	v971 = *(*int32)(unsafe.Add(mBase, uint32(l2)+8))
	v972 = *(*int32)(unsafe.Add(mBase, uint32(l3)+8))
	v973 = int64(0)
	*(*int64)(unsafe.Add(mBase, uint32(v970)+48)) = v973
	*(*int32)(unsafe.Add(mBase, uint32(v970)+16)) = v972
	*(*int32)(unsafe.Add(mBase, uint32(v970)+12)) = v971
	*(*int32)(unsafe.Add(mBase, uint32(v970)+8)) = v972
	*(*int32)(unsafe.Add(mBase, uint32(v970)+4)) = v971
	*(*int32)(unsafe.Add(mBase, uint32(v970))) = int32(320)
	*(*int64)(unsafe.Add(mBase, uint32(v970)+20)) = v973
	*(*int64)(unsafe.Add(mBase, uint32(v970)+28)) = v973
	*(*int64)(unsafe.Add(mBase, uint32(v970)+36)) = v973
	*(*int32)(unsafe.Add(mBase, uint32(v970)+43)) = int32(0)
	goto L192
L192:
	;
	v989 = int32(0)
	v993 = F_clauselist_selectivity(m, l0, v938, v989, v989, v807+int32(-56))
	mBase = m.M
	v994 = m.ExcPending
	if v994 != 0 {
		goto L22
	} else {
		goto L193
	}
L193:
	;
	if v814 != 0 {
		goto L194
	} else {
		goto L195
	}
L194:
	;
	F_list_free(m, v938)
	mBase = m.M
	v996 = m.ExcPending
	if v996 != 0 {
		goto L22
	} else {
		goto L197
	}
L195:
	;
	goto L196
L196:
	;
	if base.F64_gt(v967, float64(0)) != 0 {
		goto L198
	} else {
		goto L199
	}
L197:
	;
	goto L196
L198:
	;
	v999 = float64(1)
	v1000 = *(*float64)(unsafe.Add(mBase, uint32(l3)+16))
	v1002 = base.F64_div(base.F64_mul(v993, v1000), v967)
	if base.F64_lt(v1002, v999) != 0 {
		goto L201
	} else {
		goto L202
	}
L199:
	;
	v1008 = float64(1)
	goto L200
L200:
	;
	*(*float64)(unsafe.Add(mBase, uint32(v805)+8)) = v1008
	*(*float64)(unsafe.Add(mBase, uint32(v805))) = v967
	m.G0 = v809 - int32(-64)
	goto L160
L201:
	;
	v1005 = v999
	goto L203
L202:
	;
	v1005 = v1002
	goto L203
L203:
	;
	v1008 = v1005
	goto L200
L204:
	;
	v1339 = *(*int32)(unsafe.Add(mBase, uint32(v37)+40))
	v1340 = *(*int32)(unsafe.Add(mBase, uint32(l1)+64))
	v1341 = F_bms_add_members(m, v1339, v1340)
	mBase = m.M
	v1342 = m.ExcPending
	if v1342 != 0 {
		goto L22
	} else {
		goto L277
	}
L205:
	;
	v1051 = int32(0)
	v1052 = *(*int32)(unsafe.Add(mBase, uint32(v1048)+4))
	if v1052 <= v1051 {
		goto L204
	} else {
		goto L206
	}
L206:
	;
	v1061 = v1051
	goto L207
L207:
	;
	v1089 = *(*int32)(unsafe.Add(mBase, uint32(v1048)+12))
	v1093 = *(*int32)(unsafe.Add(mBase, uint32(v1089+v1061<<(uint(int32(2))%32))))
	v1094 = *(*int32)(unsafe.Add(mBase, uint32(v1093)+8))
	v1095 = int32(0)
	if v46 == v1095 {
		v1136 = v1095
		goto L211
	} else {
		goto L212
	}
L208:
	;
	goto L204
L209:
	;
	v1196 = *(*int32)(unsafe.Add(mBase, uint32(v1093)+20))
	if v1196 != int32(2) {
		goto L242
	} else {
		goto L243
	}
L210:
	;
	if v1136 == int32(0) {
		goto L209
	} else {
		goto L224
	}
L211:
	;
	goto L210
L212:
	;
	if v1094 == int32(0) {
		v1136 = v1095
		goto L211
	} else {
		goto L213
	}
L213:
	;
	v1104 = *(*int32)(unsafe.Add(mBase, uint32(v46)+4))
	v1105 = *(*int32)(unsafe.Add(mBase, uint32(v1094)+4))
	if v1104 < v1105 {
		goto L214
	} else {
		goto L215
	}
L214:
	;
	v1107 = v1104
	goto L216
L215:
	;
	v1107 = v1105
	goto L216
L216:
	;
	if v1107 <= int32(1) {
		goto L217
	} else {
		goto L218
	}
L217:
	;
	v1110 = int32(1)
	goto L219
L218:
	;
	v1110 = v1107
	goto L219
L219:
	;
	v1111 = int32(8)
	v1116 = int32(0)
	goto L220
L220:
	;
	v1123 = v1116 << (uint(int32(2)) % 32)
	v1125 = *(*int32)(unsafe.Add(mBase, uint32(v1094+v1111+v1123)))
	v1127 = *(*int32)(unsafe.Add(mBase, uint32(v1123+(v46+v1111))))
	v1128 = v1125 & v1127
	v1130 = base.B2i32(v1128 != int32(0))
	if v1128 != 0 {
		v1136 = v1130
		goto L211
	} else {
		goto L222
	}
L221:
	;
	v1136 = v1130
	goto L211
L222:
	;
	v1132 = v1116 + int32(1)
	if v1132 != v1110 {
		v1116 = v1132
		goto L220
	} else {
		goto L223
	}
L223:
	;
	goto L221
L224:
	;
	v1142 = *(*int32)(unsafe.Add(mBase, uint32(v1093)+4))
	v1143 = int32(0)
	if v46 == v1143 {
		v1184 = v1143
		goto L226
	} else {
		goto L227
	}
L225:
	;
	if v1184 != 0 {
		goto L209
	} else {
		goto L239
	}
L226:
	;
	goto L225
L227:
	;
	if v1142 == int32(0) {
		v1184 = v1143
		goto L226
	} else {
		goto L228
	}
L228:
	;
	v1152 = *(*int32)(unsafe.Add(mBase, uint32(v46)+4))
	v1153 = *(*int32)(unsafe.Add(mBase, uint32(v1142)+4))
	if v1152 < v1153 {
		goto L229
	} else {
		goto L230
	}
L229:
	;
	v1155 = v1152
	goto L231
L230:
	;
	v1155 = v1153
	goto L231
L231:
	;
	if v1155 <= int32(1) {
		goto L232
	} else {
		goto L233
	}
L232:
	;
	v1158 = int32(1)
	goto L234
L233:
	;
	v1158 = v1155
	goto L234
L234:
	;
	v1159 = int32(8)
	v1164 = int32(0)
	goto L235
L235:
	;
	v1171 = v1164 << (uint(int32(2)) % 32)
	v1173 = *(*int32)(unsafe.Add(mBase, uint32(v1142+v1159+v1171)))
	v1175 = *(*int32)(unsafe.Add(mBase, uint32(v1171+(v46+v1159))))
	v1176 = v1173 & v1175
	v1178 = base.B2i32(v1176 != int32(0))
	if v1176 != 0 {
		v1184 = v1178
		goto L226
	} else {
		goto L237
	}
L236:
	;
	v1184 = v1178
	goto L226
L237:
	;
	v1180 = v1164 + int32(1)
	if v1180 != v1158 {
		v1164 = v1180
		goto L235
	} else {
		goto L238
	}
L238:
	;
	goto L236
L239:
	;
	v1188 = *(*int32)(unsafe.Add(mBase, uint32(v37)+40))
	v1189 = *(*int32)(unsafe.Add(mBase, uint32(l0)+44))
	v1190 = *(*int32)(unsafe.Add(mBase, uint32(v1093)+8))
	v1191 = F_bms_difference(m, v1189, v1190)
	mBase = m.M
	v1192 = m.ExcPending
	if v1192 != 0 {
		goto L22
	} else {
		goto L240
	}
L240:
	;
	v1193 = F_bms_join(m, v1188, v1191)
	mBase = m.M
	v1194 = m.ExcPending
	if v1194 != 0 {
		goto L22
	} else {
		goto L241
	}
L241:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v37)+40)) = v1193
	goto L209
L242:
	;
	v1302 = v1061 + int32(1)
	v1303 = *(*int32)(unsafe.Add(mBase, uint32(v1048)+4))
	if v1302 < v1303 {
		v1061 = v1302
		goto L207
	} else {
		goto L276
	}
L243:
	;
	v1199 = *(*int32)(unsafe.Add(mBase, uint32(v1093)+4))
	v1200 = int32(0)
	if v46 == v1200 {
		v1241 = v1200
		goto L245
	} else {
		goto L246
	}
L244:
	;
	if v1241 == int32(0) {
		goto L242
	} else {
		goto L258
	}
L245:
	;
	goto L244
L246:
	;
	if v1199 == int32(0) {
		v1241 = v1200
		goto L245
	} else {
		goto L247
	}
L247:
	;
	v1209 = *(*int32)(unsafe.Add(mBase, uint32(v46)+4))
	v1210 = *(*int32)(unsafe.Add(mBase, uint32(v1199)+4))
	if v1209 < v1210 {
		goto L248
	} else {
		goto L249
	}
L248:
	;
	v1212 = v1209
	goto L250
L249:
	;
	v1212 = v1210
	goto L250
L250:
	;
	if v1212 <= int32(1) {
		goto L251
	} else {
		goto L252
	}
L251:
	;
	v1215 = int32(1)
	goto L253
L252:
	;
	v1215 = v1212
	goto L253
L253:
	;
	v1216 = int32(8)
	v1221 = int32(0)
	goto L254
L254:
	;
	v1228 = v1221 << (uint(int32(2)) % 32)
	v1230 = *(*int32)(unsafe.Add(mBase, uint32(v1199+v1216+v1228)))
	v1232 = *(*int32)(unsafe.Add(mBase, uint32(v1228+(v46+v1216))))
	v1233 = v1230 & v1232
	v1235 = base.B2i32(v1233 != int32(0))
	if v1233 != 0 {
		v1241 = v1235
		goto L245
	} else {
		goto L256
	}
L255:
	;
	v1241 = v1235
	goto L245
L256:
	;
	v1237 = v1221 + int32(1)
	if v1237 != v1215 {
		v1221 = v1237
		goto L254
	} else {
		goto L257
	}
L257:
	;
	goto L255
L258:
	;
	v1247 = *(*int32)(unsafe.Add(mBase, uint32(v1093)+8))
	v1248 = int32(0)
	if v46 == v1248 {
		v1289 = v1248
		goto L260
	} else {
		goto L261
	}
L259:
	;
	if v1289 != 0 {
		goto L242
	} else {
		goto L273
	}
L260:
	;
	goto L259
L261:
	;
	if v1247 == int32(0) {
		v1289 = v1248
		goto L260
	} else {
		goto L262
	}
L262:
	;
	v1257 = *(*int32)(unsafe.Add(mBase, uint32(v46)+4))
	v1258 = *(*int32)(unsafe.Add(mBase, uint32(v1247)+4))
	if v1257 < v1258 {
		goto L263
	} else {
		goto L264
	}
L263:
	;
	v1260 = v1257
	goto L265
L264:
	;
	v1260 = v1258
	goto L265
L265:
	;
	if v1260 <= int32(1) {
		goto L266
	} else {
		goto L267
	}
L266:
	;
	v1263 = int32(1)
	goto L268
L267:
	;
	v1263 = v1260
	goto L268
L268:
	;
	v1264 = int32(8)
	v1269 = int32(0)
	goto L269
L269:
	;
	v1276 = v1269 << (uint(int32(2)) % 32)
	v1278 = *(*int32)(unsafe.Add(mBase, uint32(v1247+v1264+v1276)))
	v1280 = *(*int32)(unsafe.Add(mBase, uint32(v1276+(v46+v1264))))
	v1281 = v1278 & v1280
	v1283 = base.B2i32(v1281 != int32(0))
	if v1281 != 0 {
		v1289 = v1283
		goto L260
	} else {
		goto L271
	}
L270:
	;
	v1289 = v1283
	goto L260
L271:
	;
	v1285 = v1269 + int32(1)
	if v1285 != v1263 {
		v1269 = v1285
		goto L269
	} else {
		goto L272
	}
L272:
	;
	goto L270
L273:
	;
	v1293 = *(*int32)(unsafe.Add(mBase, uint32(v37)+40))
	v1294 = *(*int32)(unsafe.Add(mBase, uint32(l0)+44))
	v1295 = *(*int32)(unsafe.Add(mBase, uint32(v1093)+4))
	v1296 = F_bms_difference(m, v1294, v1295)
	mBase = m.M
	v1297 = m.ExcPending
	if v1297 != 0 {
		goto L22
	} else {
		goto L274
	}
L274:
	;
	v1298 = F_bms_join(m, v1293, v1296)
	mBase = m.M
	v1299 = m.ExcPending
	if v1299 != 0 {
		goto L22
	} else {
		goto L275
	}
L275:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v37)+40)) = v1298
	goto L242
L276:
	;
	goto L208
L277:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v37)+40)) = v1341
	if v778&int32(1) == int32(0) {
		goto L278
	} else {
		goto L279
	}
L278:
	;
	if l4 != int32(2) {
		goto L773
	} else {
		goto L774
	}
L279:
	;
	v1348 = *(*int32)(unsafe.Add(mBase, uint32(v37)+12))
	if v1348 == int32(0) {
		goto L280
	} else {
		goto L281
	}
L280:
	;
	v2925 = *(*int32)(unsafe.Add(mBase, uint32(l3)+48))
	v2926 = int32(1)
	switch l4 {
	case 0, 1, 4, 5:
		v2943 = l4
		v2944 = v2926
		v2945 = v8
		goto L540
	case 2, 3, 7:
		goto L541
	case 6:
		goto L278
	case 8, 9:
		goto L543
	default:
		goto L542
	}
L281:
	;
	v1351 = *(*int32)(unsafe.Add(mBase, uint32(l3)+48))
	v1352 = *(*int32)(unsafe.Add(mBase, uint32(l2)+48))
	v1353 = *(*int32)(unsafe.Add(mBase, uint32(v1352)+16))
	if v1353 == int32(0) {
		goto L282
	} else {
		goto L283
	}
L282:
	;
	v1454 = *(*int32)(unsafe.Add(mBase, uint32(v1351)+16))
	if v1454 == int32(0) {
		goto L315
	} else {
		goto L316
	}
L283:
	;
	v1356 = *(*int32)(unsafe.Add(mBase, uint32(v1353)+4))
	v1357 = *(*int32)(unsafe.Add(mBase, uint32(l3)+8))
	v1358 = int32(0)
	if v1356 == v1358 {
		v1399 = v1358
		goto L285
	} else {
		goto L286
	}
L284:
	;
	if v1399 != 0 {
		goto L280
	} else {
		goto L298
	}
L285:
	;
	goto L284
L286:
	;
	if v1357 == int32(0) {
		v1399 = v1358
		goto L285
	} else {
		goto L287
	}
L287:
	;
	v1367 = *(*int32)(unsafe.Add(mBase, uint32(v1356)+4))
	v1368 = *(*int32)(unsafe.Add(mBase, uint32(v1357)+4))
	if v1367 < v1368 {
		goto L288
	} else {
		goto L289
	}
L288:
	;
	v1370 = v1367
	goto L290
L289:
	;
	v1370 = v1368
	goto L290
L290:
	;
	if v1370 <= int32(1) {
		goto L291
	} else {
		goto L292
	}
L291:
	;
	v1373 = int32(1)
	goto L293
L292:
	;
	v1373 = v1370
	goto L293
L293:
	;
	v1374 = int32(8)
	v1379 = int32(0)
	goto L294
L294:
	;
	v1386 = v1379 << (uint(int32(2)) % 32)
	v1388 = *(*int32)(unsafe.Add(mBase, uint32(v1357+v1374+v1386)))
	v1390 = *(*int32)(unsafe.Add(mBase, uint32(v1386+(v1356+v1374))))
	v1391 = v1388 & v1390
	v1393 = base.B2i32(v1391 != int32(0))
	if v1391 != 0 {
		v1399 = v1393
		goto L285
	} else {
		goto L296
	}
L295:
	;
	v1399 = v1393
	goto L285
L296:
	;
	v1395 = v1379 + int32(1)
	if v1395 != v1373 {
		v1379 = v1395
		goto L294
	} else {
		goto L297
	}
L297:
	;
	goto L295
L298:
	;
	v1403 = *(*int32)(unsafe.Add(mBase, uint32(v1352)+16))
	if v1403 == int32(0) {
		goto L282
	} else {
		goto L299
	}
L299:
	;
	v1406 = *(*int32)(unsafe.Add(mBase, uint32(v1403)+4))
	v1407 = *(*int32)(unsafe.Add(mBase, uint32(l3)+228))
	v1408 = int32(0)
	if v1406 == v1408 {
		v1449 = v1408
		goto L301
	} else {
		goto L302
	}
L300:
	;
	if v1449 != 0 {
		goto L280
	} else {
		goto L314
	}
L301:
	;
	goto L300
L302:
	;
	if v1407 == int32(0) {
		v1449 = v1408
		goto L301
	} else {
		goto L303
	}
L303:
	;
	v1417 = *(*int32)(unsafe.Add(mBase, uint32(v1406)+4))
	v1418 = *(*int32)(unsafe.Add(mBase, uint32(v1407)+4))
	if v1417 < v1418 {
		goto L304
	} else {
		goto L305
	}
L304:
	;
	v1420 = v1417
	goto L306
L305:
	;
	v1420 = v1418
	goto L306
L306:
	;
	if v1420 <= int32(1) {
		goto L307
	} else {
		goto L308
	}
L307:
	;
	v1423 = int32(1)
	goto L309
L308:
	;
	v1423 = v1420
	goto L309
L309:
	;
	v1424 = int32(8)
	v1429 = int32(0)
	goto L310
L310:
	;
	v1436 = v1429 << (uint(int32(2)) % 32)
	v1438 = *(*int32)(unsafe.Add(mBase, uint32(v1407+v1424+v1436)))
	v1440 = *(*int32)(unsafe.Add(mBase, uint32(v1436+(v1406+v1424))))
	v1441 = v1438 & v1440
	v1443 = base.B2i32(v1441 != int32(0))
	if v1441 != 0 {
		v1449 = v1443
		goto L301
	} else {
		goto L312
	}
L311:
	;
	v1449 = v1443
	goto L301
L312:
	;
	v1445 = v1429 + int32(1)
	if v1445 != v1423 {
		v1429 = v1445
		goto L310
	} else {
		goto L313
	}
L313:
	;
	goto L311
L314:
	;
	goto L282
L315:
	;
	switch l4 - int32(8) {
	case 0:
		goto L350
	case 1:
		goto L349
	default:
		v1565 = v1352
		v1566 = v1351
		v1567 = l4
		goto L348
	}
L316:
	;
	v1457 = *(*int32)(unsafe.Add(mBase, uint32(v1454)+4))
	v1458 = *(*int32)(unsafe.Add(mBase, uint32(l2)+8))
	v1459 = int32(0)
	if v1457 == v1459 {
		v1500 = v1459
		goto L318
	} else {
		goto L319
	}
L317:
	;
	if v1500 != 0 {
		goto L280
	} else {
		goto L331
	}
L318:
	;
	goto L317
L319:
	;
	if v1458 == int32(0) {
		v1500 = v1459
		goto L318
	} else {
		goto L320
	}
L320:
	;
	v1468 = *(*int32)(unsafe.Add(mBase, uint32(v1457)+4))
	v1469 = *(*int32)(unsafe.Add(mBase, uint32(v1458)+4))
	if v1468 < v1469 {
		goto L321
	} else {
		goto L322
	}
L321:
	;
	v1471 = v1468
	goto L323
L322:
	;
	v1471 = v1469
	goto L323
L323:
	;
	if v1471 <= int32(1) {
		goto L324
	} else {
		goto L325
	}
L324:
	;
	v1474 = int32(1)
	goto L326
L325:
	;
	v1474 = v1471
	goto L326
L326:
	;
	v1475 = int32(8)
	v1480 = int32(0)
	goto L327
L327:
	;
	v1487 = v1480 << (uint(int32(2)) % 32)
	v1489 = *(*int32)(unsafe.Add(mBase, uint32(v1458+v1475+v1487)))
	v1491 = *(*int32)(unsafe.Add(mBase, uint32(v1487+(v1457+v1475))))
	v1492 = v1489 & v1491
	v1494 = base.B2i32(v1492 != int32(0))
	if v1492 != 0 {
		v1500 = v1494
		goto L318
	} else {
		goto L329
	}
L328:
	;
	v1500 = v1494
	goto L318
L329:
	;
	v1496 = v1480 + int32(1)
	if v1496 != v1474 {
		v1480 = v1496
		goto L327
	} else {
		goto L330
	}
L330:
	;
	goto L328
L331:
	;
	v1504 = *(*int32)(unsafe.Add(mBase, uint32(v1351)+16))
	if v1504 == int32(0) {
		goto L315
	} else {
		goto L332
	}
L332:
	;
	v1507 = *(*int32)(unsafe.Add(mBase, uint32(v1504)+4))
	v1508 = *(*int32)(unsafe.Add(mBase, uint32(l2)+228))
	v1509 = int32(0)
	if v1507 == v1509 {
		v1550 = v1509
		goto L334
	} else {
		goto L335
	}
L333:
	;
	if v1550 != 0 {
		goto L280
	} else {
		goto L347
	}
L334:
	;
	goto L333
L335:
	;
	if v1508 == int32(0) {
		v1550 = v1509
		goto L334
	} else {
		goto L336
	}
L336:
	;
	v1518 = *(*int32)(unsafe.Add(mBase, uint32(v1507)+4))
	v1519 = *(*int32)(unsafe.Add(mBase, uint32(v1508)+4))
	if v1518 < v1519 {
		goto L337
	} else {
		goto L338
	}
L337:
	;
	v1521 = v1518
	goto L339
L338:
	;
	v1521 = v1519
	goto L339
L339:
	;
	if v1521 <= int32(1) {
		goto L340
	} else {
		goto L341
	}
L340:
	;
	v1524 = int32(1)
	goto L342
L341:
	;
	v1524 = v1521
	goto L342
L342:
	;
	v1525 = int32(8)
	v1530 = int32(0)
	goto L343
L343:
	;
	v1537 = v1530 << (uint(int32(2)) % 32)
	v1539 = *(*int32)(unsafe.Add(mBase, uint32(v1508+v1525+v1537)))
	v1541 = *(*int32)(unsafe.Add(mBase, uint32(v1537+(v1507+v1525))))
	v1542 = v1539 & v1541
	v1544 = base.B2i32(v1542 != int32(0))
	if v1542 != 0 {
		v1550 = v1544
		goto L334
	} else {
		goto L345
	}
L344:
	;
	v1550 = v1544
	goto L334
L345:
	;
	v1546 = v1530 + int32(1)
	if v1546 != v1524 {
		v1530 = v1546
		goto L343
	} else {
		goto L346
	}
L346:
	;
	goto L344
L347:
	;
	goto L315
L348:
	;
	v1568 = int32(0)
	v1569 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l1)+26)))
	if v1569 != int32(1) {
		v1638 = v1568
		v1639 = v8
		goto L353
	} else {
		goto L354
	}
L349:
	;
	v1562 = *(*int32)(unsafe.Add(mBase, uint32(v37)+20))
	v1563 = F_create_unique_path(m, l0, l3, v1351, v1562)
	mBase = m.M
	v1564 = m.ExcPending
	if v1564 != 0 {
		goto L22
	} else {
		goto L352
	}
L350:
	;
	v1558 = *(*int32)(unsafe.Add(mBase, uint32(v37)+20))
	v1559 = F_create_unique_path(m, l0, l2, v1352, v1558)
	mBase = m.M
	v1560 = m.ExcPending
	if v1560 != 0 {
		goto L22
	} else {
		goto L351
	}
L351:
	;
	v1565 = v1559
	v1566 = v1351
	v1567 = int32(0)
	goto L348
L352:
	;
	v1565 = v1352
	v1566 = v1563
	v1567 = int32(0)
	goto L348
L353:
	;
	v1640 = *(*int32)(unsafe.Add(mBase, uint32(v37)+12))
	if v1640 == int32(0) {
		v2805 = v8
		goto L379
	} else {
		goto L380
	}
L354:
	;
	v1573 = l4 - int32(2)
	if base.B2i32(base.Ui32(v1573) < base.Ui32(int32(7)))&(int32(base.Ui32(int32(99))>>(uint(v1573)%32))&int32(1)) != 0 {
		v1638 = v1568
		v1639 = v8
		goto L353
	} else {
		goto L355
	}
L355:
	;
	v1581 = *(*int32)(unsafe.Add(mBase, uint32(l2)+40))
	if v1581 == int32(0) {
		v1638 = v1568
		v1639 = v8
		goto L353
	} else {
		goto L356
	}
L356:
	;
	v1584 = *(*int32)(unsafe.Add(mBase, uint32(l1)+64))
	if v1584 != 0 {
		v1638 = v1568
		v1639 = v8
		goto L353
	} else {
		goto L357
	}
L357:
	;
	v1585 = *(*int32)(unsafe.Add(mBase, uint32(v1581)+12))
	v1586 = *(*int32)(unsafe.Add(mBase, uint32(v1585)))
	v1587 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1566)+21)))
	if v1587 != 0 {
		goto L358
	} else {
		goto L359
	}
L358:
	;
	v1638 = v1566
	v1639 = v1586
	goto L353
L359:
	;
	goto L360
L360:
	;
	if l4 == int32(9) {
		v1638 = v1568
		v1639 = v1586
		goto L353
	} else {
		goto L361
	}
L361:
	;
	v1590 = *(*int32)(unsafe.Add(mBase, uint32(l3)+32))
	if v1590 != 0 {
		goto L364
	} else {
		goto L365
	}
L362:
	;
	v1638 = v1633
	v1639 = v1586
	goto L353
L363:
	;
	goto L362
L364:
	;
	v1595 = *(*int32)(unsafe.Add(mBase, uint32(v1590)+4))
	if v1595 <= int32(0) {
		v1633 = int32(0)
		goto L363
	} else {
		goto L367
	}
L365:
	;
	goto L366
L366:
	;
	v1633 = int32(0)
	goto L363
L367:
	;
	v1598 = int32(0)
	if v1598 < v1595 {
		goto L368
	} else {
		goto L369
	}
L368:
	;
	v1601 = v1595
	goto L370
L369:
	;
	v1601 = v1598
	goto L370
L370:
	;
	v1602 = *(*int32)(unsafe.Add(mBase, uint32(v1590)+12))
	v1604 = int32(0)
	goto L371
L371:
	;
	v1612 = *(*int32)(unsafe.Add(mBase, uint32(v1602+v1604<<(uint(int32(2))%32))))
	v1613 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1612)+21)))
	if v1613 == int32(1) {
		goto L373
	} else {
		goto L374
	}
L372:
	;
	goto L366
L373:
	;
	v1616 = *(*int32)(unsafe.Add(mBase, uint32(v1612)+16))
	if v1616 == int32(0) {
		v1633 = v1612
		goto L363
	} else {
		goto L376
	}
L374:
	;
	goto L375
L375:
	;
	v1624 = v1604 + int32(1)
	if v1624 != v1601 {
		v1604 = v1624
		goto L371
	} else {
		goto L378
	}
L376:
	;
	v1619 = *(*int32)(unsafe.Add(mBase, uint32(v1616)+4))
	if v1619 == int32(0) {
		v1633 = v1612
		goto L363
	} else {
		goto L377
	}
L377:
	;
	goto L375
L378:
	;
	goto L372
L379:
	;
	if v2805 == int32(0) {
		goto L280
	} else {
		goto L521
	}
L380:
	;
	v1643 = *(*int32)(unsafe.Add(mBase, uint32(v1640)+4))
	if v1643 == int32(0) {
		v2805 = v8
		goto L379
	} else {
		goto L381
	}
L381:
	;
	v1647 = v1643 << (uint(int32(2)) % 32)
	v1648 = F_palloc(m, v1647)
	mBase = m.M
	v1649 = m.ExcPending
	if v1649 != 0 {
		goto L22
	} else {
		goto L382
	}
L382:
	;
	v1650 = F_palloc(m, v1647)
	mBase = m.M
	v1651 = m.ExcPending
	if v1651 != 0 {
		goto L22
	} else {
		goto L383
	}
L383:
	;
	v1652 = *(*int32)(unsafe.Add(mBase, uint32(v1640)+4))
	if int32(0) < v1652 {
		goto L384
	} else {
		goto L385
	}
L384:
	;
	v1674 = v8
	v1679 = v8
	goto L387
L385:
	;
	v2128 = v8
	goto L386
L386:
	;
	v2143 = *(*int32)(unsafe.Add(mBase, uint32(l0)+156))
	if v2143 == int32(0) {
		v2472 = v8
		goto L439
	} else {
		goto L440
	}
L387:
	;
	v1689 = *(*int32)(unsafe.Add(mBase, uint32(v1640)+12))
	v1693 = *(*int32)(unsafe.Add(mBase, uint32(v1689+v1679<<(uint(int32(2))%32))))
	v1694 = *(*int32)(unsafe.Add(mBase, uint32(v1693)+100))
	v1695 = *(*int32)(unsafe.Add(mBase, uint32(v1694)+56))
	if v1695 != 0 {
		goto L389
	} else {
		goto L390
	}
L388:
	;
	v2128 = v2090
	goto L386
L389:
	;
	v1706 = v1695
	goto L392
L390:
	;
	goto L391
L391:
	;
	v1766 = *(*int32)(unsafe.Add(mBase, uint32(v1693)+104))
	v1767 = *(*int32)(unsafe.Add(mBase, uint32(v1766)+56))
	if v1767 != 0 {
		goto L395
	} else {
		goto L396
	}
L392:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v1693)+100)) = v1706
	v1731 = *(*int32)(unsafe.Add(mBase, uint32(v1706)+56))
	if v1731 != 0 {
		v1706 = v1731
		goto L392
	} else {
		goto L394
	}
L393:
	;
	goto L391
L394:
	;
	goto L393
L395:
	;
	v1778 = v1767
	goto L398
L396:
	;
	goto L397
L397:
	;
	v1840 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1693)+120)))
	if v1840 != 0 {
		goto L401
	} else {
		goto L402
	}
L398:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v1693)+104)) = v1778
	v1803 = *(*int32)(unsafe.Add(mBase, uint32(v1778)+56))
	if v1803 != 0 {
		v1778 = v1803
		goto L398
	} else {
		goto L400
	}
L399:
	;
	goto L397
L400:
	;
	goto L399
L401:
	;
	v1841 = int32(100)
	goto L403
L402:
	;
	v1841 = int32(104)
	goto L403
L403:
	;
	v1843 = *(*int32)(unsafe.Add(mBase, uint32(v1693+v1841)))
	v1844 = int32(0)
	if v1844 < v1674 {
		goto L405
	} else {
		goto L406
	}
L404:
	;
	v2106 = v1679 + int32(1)
	v2107 = *(*int32)(unsafe.Add(mBase, uint32(v1640)+4))
	if v2106 < v2107 {
		v1674 = v2090
		v1679 = v2106
		goto L387
	} else {
		goto L437
	}
L405:
	;
	v1857 = v1844
	goto L408
L406:
	;
	goto L407
L407:
	;
	v1923 = *(*int32)(unsafe.Add(mBase, uint32(v1843)+16))
	if v1923 == int32(0) {
		goto L413
	} else {
		goto L414
	}
L408:
	;
	v1884 = *(*int32)(unsafe.Add(mBase, uint32(v1648+v1857<<(uint(int32(2))%32))))
	if v1884 == v1843 {
		v2090 = v1674
		goto L404
	} else {
		goto L410
	}
L409:
	;
	goto L407
L410:
	;
	v1887 = v1857 + int32(1)
	if v1887 != v1674 {
		v1857 = v1887
		goto L408
	} else {
		goto L411
	}
L411:
	;
	goto L409
L412:
	;
	v2064 = v1674 << (uint(int32(2)) % 32)
	*(*int32)(unsafe.Add(mBase, uint32(v1648+v2064))) = v1843
	*(*int32)(unsafe.Add(mBase, uint32(v2064+v1650))) = v2035
	v2090 = v1674 + int32(1)
	goto L404
L413:
	;
	v2035 = int32(0)
	goto L412
L414:
	;
	goto L415
L415:
	;
	v1927 = int32(0)
	v1929 = *(*int32)(unsafe.Add(mBase, uint32(v1923)+4))
	if v1929 <= v1927 {
		v2035 = v1927
		goto L412
	} else {
		goto L416
	}
L416:
	;
	v1938 = v1927
	v1942 = v1927
	goto L417
L417:
	;
	v1966 = *(*int32)(unsafe.Add(mBase, uint32(v1923)+12))
	v1970 = *(*int32)(unsafe.Add(mBase, uint32(v1966+v1942<<(uint(int32(2))%32))))
	v1971 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1970)+12)))
	if v1971 == int32(0) {
		goto L419
	} else {
		goto L420
	}
L418:
	;
	v2035 = v2024
	goto L412
L419:
	;
	v1974 = *(*int32)(unsafe.Add(mBase, uint32(v1970)+8))
	v1975 = *(*int32)(unsafe.Add(mBase, uint32(l1)+8))
	v1976 = int32(0)
	if v1974 == v1976 {
		v2017 = v1976
		goto L423
	} else {
		goto L424
	}
L420:
	;
	v2024 = v1938
	goto L421
L421:
	;
	v2026 = v1942 + int32(1)
	v2027 = *(*int32)(unsafe.Add(mBase, uint32(v1923)+4))
	if v2026 < v2027 {
		v1938 = v2024
		v1942 = v2026
		goto L417
	} else {
		goto L436
	}
L422:
	;
	v2024 = v1938 + (v2017 ^ int32(1))
	goto L421
L423:
	;
	goto L422
L424:
	;
	if v1975 == int32(0) {
		v2017 = v1976
		goto L423
	} else {
		goto L425
	}
L425:
	;
	v1985 = *(*int32)(unsafe.Add(mBase, uint32(v1974)+4))
	v1986 = *(*int32)(unsafe.Add(mBase, uint32(v1975)+4))
	if v1985 < v1986 {
		goto L426
	} else {
		goto L427
	}
L426:
	;
	v1988 = v1985
	goto L428
L427:
	;
	v1988 = v1986
	goto L428
L428:
	;
	if v1988 <= int32(1) {
		goto L429
	} else {
		goto L430
	}
L429:
	;
	v1991 = int32(1)
	goto L431
L430:
	;
	v1991 = v1988
	goto L431
L431:
	;
	v1992 = int32(8)
	v1997 = int32(0)
	goto L432
L432:
	;
	v2004 = v1997 << (uint(int32(2)) % 32)
	v2006 = *(*int32)(unsafe.Add(mBase, uint32(v1975+v1992+v2004)))
	v2008 = *(*int32)(unsafe.Add(mBase, uint32(v2004+(v1974+v1992))))
	v2009 = v2006 & v2008
	v2011 = base.B2i32(v2009 != int32(0))
	if v2009 != 0 {
		v2017 = v2011
		goto L423
	} else {
		goto L434
	}
L433:
	;
	v2017 = v2011
	goto L423
L434:
	;
	v2013 = v1997 + int32(1)
	if v2013 != v1991 {
		v1997 = v2013
		goto L432
	} else {
		goto L435
	}
L435:
	;
	goto L433
L436:
	;
	goto L418
L437:
	;
	goto L388
L438:
	;
	F_pfree(m, v1648)
	mBase = m.M
	v2775 = m.ExcPending
	if v2775 != 0 {
		goto L22
	} else {
		goto L519
	}
L439:
	;
	v2480 = v2128 - int32(1)
	v2483 = int32(3)
	v2484 = v2480 & v2483
	v2485 = int32(2)
	v2518 = v2472
	goto L470
L440:
	;
	v2146 = *(*int32)(unsafe.Add(mBase, uint32(v2143)+4))
	if v2146 <= int32(0) {
		goto L441
	} else {
		goto L442
	}
L441:
	;
	v2310 = F_list_copy(m, v2143)
	mBase = m.M
	v2311 = m.ExcPending
	if v2311 != 0 {
		goto L22
	} else {
		goto L456
	}
L442:
	;
	v2149 = *(*int32)(unsafe.Add(mBase, uint32(v2143)+12))
	v2150 = int32(0)
	v2160 = v2150
	goto L444
L443:
	;
	if v2245 != v1643 {
		v2472 = v8
		goto L439
	} else {
		goto L454
	}
L444:
	;
	if v2128 <= int32(0) {
		v2245 = v2150
		goto L443
	} else {
		goto L446
	}
L445:
	;
	v2245 = v2160
	goto L443
L446:
	;
	v2191 = *(*int32)(unsafe.Add(mBase, uint32(v2149+v2160<<(uint(int32(2))%32))))
	v2192 = *(*int32)(unsafe.Add(mBase, uint32(v2191)+4))
	v2204 = int32(0)
	goto L447
L447:
	;
	v2231 = *(*int32)(unsafe.Add(mBase, uint32(v1648+v2204<<(uint(int32(2))%32))))
	if v2192 == v2231 {
		goto L449
	} else {
		goto L450
	}
L448:
	;
	goto L445
L449:
	;
	v2234 = v2160 + int32(1)
	if v2234 != v2146 {
		v2160 = v2234
		goto L444
	} else {
		goto L452
	}
L450:
	;
	goto L451
L451:
	;
	v2237 = v2204 + int32(1)
	if v2237 != v2128 {
		v2204 = v2237
		goto L447
	} else {
		goto L453
	}
L452:
	;
	goto L441
L453:
	;
	goto L448
L454:
	;
	v2274 = F_list_copy_head(m, v2143, v1643)
	mBase = m.M
	v2275 = m.ExcPending
	if v2275 != 0 {
		goto L22
	} else {
		goto L455
	}
L455:
	;
	v2767 = v2274
	goto L438
L456:
	;
	v2312 = *(*int32)(unsafe.Add(mBase, uint32(l0)+156))
	if v2312 == int32(0) {
		v2472 = v2310
		goto L439
	} else {
		goto L457
	}
L457:
	;
	v2315 = *(*int32)(unsafe.Add(mBase, uint32(v2312)+4))
	if v2315 <= int32(0) {
		v2472 = v2310
		goto L439
	} else {
		goto L458
	}
L458:
	;
	v2318 = int32(0)
	v2329 = v2318
	goto L459
L459:
	;
	if v2128 <= v2318 {
		goto L461
	} else {
		goto L462
	}
L460:
	;
	v2472 = v2310
	goto L439
L461:
	;
	v2442 = v2329 + int32(1)
	v2443 = *(*int32)(unsafe.Add(mBase, uint32(v2312)+4))
	if v2442 < v2443 {
		v2329 = v2442
		goto L459
	} else {
		goto L469
	}
L462:
	;
	v2355 = *(*int32)(unsafe.Add(mBase, uint32(v2312)+12))
	v2359 = *(*int32)(unsafe.Add(mBase, uint32(v2355+v2329<<(uint(int32(2))%32))))
	v2360 = *(*int32)(unsafe.Add(mBase, uint32(v2359)+4))
	v2372 = int32(0)
	goto L463
L463:
	;
	v2397 = v2372 << (uint(int32(2)) % 32)
	v2399 = *(*int32)(unsafe.Add(mBase, uint32(v1648+v2397)))
	if v2360 == v2399 {
		goto L465
	} else {
		goto L466
	}
L464:
	;
	goto L461
L465:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v2397+v1650))) = int32(-1)
	goto L461
L466:
	;
	goto L467
L467:
	;
	v2405 = v2372 + int32(1)
	if v2405 != v2128 {
		v2372 = v2405
		goto L463
	} else {
		goto L468
	}
L468:
	;
	goto L464
L469:
	;
	goto L460
L470:
	;
	v2525 = *(*int32)(unsafe.Add(mBase, uint32(v1650)))
	if v2128 < v2485 {
		goto L473
	} else {
		goto L474
	}
L472:
	;
	if v2704 < int32(0) {
		v2767 = v2518
		goto L438
	} else {
		goto L516
	}
L473:
	;
	v2704 = v2525
	v2707 = int32(0)
	goto L472
L474:
	;
	goto L475
L475:
	;
	v2527 = int32(1)
	v2528 = int32(0)
	if base.B2i32(base.Ui32(v2128-v2485) < base.Ui32(v2483)) == v2528 {
		goto L476
	} else {
		goto L477
	}
L476:
	;
	v2541 = v2528
	v2543 = v2527
	v2549 = v2525
	v2552 = v2528
	goto L479
L477:
	;
	v2616 = v2527
	v2622 = v2525
	v2625 = v2528
	goto L478
L478:
	;
	if v2484 == int32(0) {
		v2704 = v2622
		v2707 = v2625
		goto L472
	} else {
		goto L506
	}
L479:
	;
	v2568 = v2543 + int32(3)
	v2569 = int32(2)
	v2572 = *(*int32)(unsafe.Add(mBase, uint32(v1650+v2568<<(uint(v2569)%32))))
	v2574 = v2543 + v2569
	v2578 = *(*int32)(unsafe.Add(mBase, uint32(v1650+v2574<<(uint(v2569)%32))))
	v2580 = v2543 + int32(1)
	v2584 = *(*int32)(unsafe.Add(mBase, uint32(v1650+v2580<<(uint(v2569)%32))))
	v2588 = *(*int32)(unsafe.Add(mBase, uint32(v1650+v2543<<(uint(v2569)%32))))
	v2589 = base.B2i32(v2549 < v2588)
	if v2549 < v2588 {
		goto L481
	} else {
		goto L482
	}
L480:
	;
	v2616 = v2602
	v2622 = v2596
	v2625 = v2600
	goto L478
L481:
	;
	v2590 = v2588
	goto L483
L482:
	;
	v2590 = v2549
	goto L483
L483:
	;
	v2591 = base.B2i32(v2590 < v2584)
	if v2590 < v2584 {
		goto L484
	} else {
		goto L485
	}
L484:
	;
	v2592 = v2584
	goto L486
L485:
	;
	v2592 = v2590
	goto L486
L486:
	;
	v2593 = base.B2i32(v2592 < v2578)
	if v2592 < v2578 {
		goto L487
	} else {
		goto L488
	}
L487:
	;
	v2594 = v2578
	goto L489
L488:
	;
	v2594 = v2592
	goto L489
L489:
	;
	v2595 = base.B2i32(v2594 < v2572)
	if v2594 < v2572 {
		goto L490
	} else {
		goto L491
	}
L490:
	;
	v2596 = v2572
	goto L492
L491:
	;
	v2596 = v2594
	goto L492
L492:
	;
	if v2549 < v2588 {
		goto L493
	} else {
		goto L494
	}
L493:
	;
	v2597 = v2543
	goto L495
L494:
	;
	v2597 = v2552
	goto L495
L495:
	;
	if v2590 < v2584 {
		goto L496
	} else {
		goto L497
	}
L496:
	;
	v2598 = v2580
	goto L498
L497:
	;
	v2598 = v2597
	goto L498
L498:
	;
	if v2592 < v2578 {
		goto L499
	} else {
		goto L500
	}
L499:
	;
	v2599 = v2574
	goto L501
L500:
	;
	v2599 = v2598
	goto L501
L501:
	;
	if v2594 < v2572 {
		goto L502
	} else {
		goto L503
	}
L502:
	;
	v2600 = v2568
	goto L504
L503:
	;
	v2600 = v2599
	goto L504
L504:
	;
	v2601 = int32(4)
	v2602 = v2543 + v2601
	v2604 = v2541 + v2601
	if v2604 != v2480&int32(-4) {
		v2541 = v2604
		v2543 = v2602
		v2549 = v2596
		v2552 = v2600
		goto L479
	} else {
		goto L505
	}
L505:
	;
	goto L480
L506:
	;
	v2652 = v2616
	v2658 = v2622
	v2661 = v2625
	v2663 = v2528
	goto L507
L507:
	;
	v2679 = *(*int32)(unsafe.Add(mBase, uint32(v1650+v2652<<(uint(int32(2))%32))))
	v2680 = base.B2i32(v2658 < v2679)
	if v2658 < v2679 {
		goto L509
	} else {
		goto L510
	}
L508:
	;
	v2704 = v2681
	v2707 = v2682
	goto L472
L509:
	;
	v2681 = v2679
	goto L511
L510:
	;
	v2681 = v2658
	goto L511
L511:
	;
	if v2658 < v2679 {
		goto L512
	} else {
		goto L513
	}
L512:
	;
	v2682 = v2652
	goto L514
L513:
	;
	v2682 = v2661
	goto L514
L514:
	;
	v2683 = int32(1)
	v2686 = v2663 + v2683
	if v2686 != v2484 {
		v2652 = v2652 + v2683
		v2658 = v2681
		v2661 = v2682
		v2663 = v2686
		goto L507
	} else {
		goto L515
	}
L515:
	;
	goto L508
L516:
	;
	v2725 = v2707 << (uint(int32(2)) % 32)
	v2727 = *(*int32)(unsafe.Add(mBase, uint32(v1648+v2725)))
	*(*int32)(unsafe.Add(mBase, uint32(v2725+v1650))) = int32(-1)
	v2731 = *(*int32)(unsafe.Add(mBase, uint32(v2727)+4))
	v2732 = *(*int32)(unsafe.Add(mBase, uint32(v2731)+12))
	v2733 = *(*int32)(unsafe.Add(mBase, uint32(v2732)))
	v2736 = F_make_canonical_pathkey(m, l0, v2727, v2733, int32(1), int32(0))
	mBase = m.M
	v2737 = m.ExcPending
	if v2737 != 0 {
		goto L22
	} else {
		goto L517
	}
L517:
	;
	v2738 = F_lappend(m, v2518, v2736)
	mBase = m.M
	v2739 = m.ExcPending
	if v2739 != 0 {
		goto L22
	} else {
		goto L518
	}
L518:
	;
	v2518 = v2738
	goto L470
L519:
	;
	F_pfree(m, v1650)
	mBase = m.M
	v2777 = m.ExcPending
	if v2777 != 0 {
		goto L22
	} else {
		goto L520
	}
L520:
	;
	v2805 = v2767
	goto L379
L521:
	;
	v2814 = *(*int32)(unsafe.Add(mBase, uint32(v2805)+4))
	if v2814 <= int32(0) {
		goto L280
	} else {
		goto L522
	}
L522:
	;
	v2817 = int32(0)
	v2829 = v2817
	goto L523
L523:
	;
	if v2829&int32(1073741823) != 0 {
		goto L525
	} else {
		goto L526
	}
L524:
	;
	goto L280
L525:
	;
	v2859 = *(*int32)(unsafe.Add(mBase, uint32(v2805)+12))
	v2863 = *(*int32)(unsafe.Add(mBase, uint32(v2859+v2829<<(uint(int32(2))%32))))
	v2864 = F_list_copy(m, v2805)
	mBase = m.M
	v2865 = m.ExcPending
	if v2865 != 0 {
		goto L22
	} else {
		goto L528
	}
L526:
	;
	v2870 = v2805
	goto L527
L527:
	;
	v2871 = *(*int32)(unsafe.Add(mBase, uint32(v37)+12))
	v2872 = F_find_mergeclauses_for_outer_pathkeys(m, v2870, v2871)
	mBase = m.M
	v2873 = m.ExcPending
	if v2873 != 0 {
		goto L22
	} else {
		goto L531
	}
L528:
	;
	v2866 = F_list_delete_nth_cell(m, v2864, v2829)
	mBase = m.M
	v2867 = m.ExcPending
	if v2867 != 0 {
		goto L22
	} else {
		goto L529
	}
L529:
	;
	v2868 = F_lcons(m, v2863, v2866)
	mBase = m.M
	v2869 = m.ExcPending
	if v2869 != 0 {
		goto L22
	} else {
		goto L530
	}
L530:
	;
	v2870 = v2868
	goto L527
L531:
	;
	v2874 = F_make_inner_pathkeys_for_merge(m, l0, v2872, v2870)
	mBase = m.M
	v2875 = m.ExcPending
	if v2875 != 0 {
		goto L22
	} else {
		goto L532
	}
L532:
	;
	v2876 = F_build_join_pathkeys(m, l0, l1, v1567, v2870)
	mBase = m.M
	v2877 = m.ExcPending
	if v2877 != 0 {
		goto L22
	} else {
		goto L533
	}
L533:
	;
	F_try_mergejoin_path(m, l0, l1, v1565, v1566, v2876, v2872, v2870, v2874, v1567, v37+int32(8), int32(0))
	mBase = m.M
	v2882 = m.ExcPending
	if v2882 != 0 {
		goto L22
	} else {
		goto L534
	}
L534:
	;
	if base.B2i32(v1639 != v2817)&base.B2i32(v1638 != v2817) != 0 {
		goto L535
	} else {
		goto L536
	}
L535:
	;
	F_try_partial_mergejoin_path(m, l0, l1, v1639, v1638, v2876, v2872, v2870, v2874, v1567, v37+int32(8))
	mBase = m.M
	v2886 = m.ExcPending
	if v2886 != 0 {
		goto L22
	} else {
		goto L538
	}
L536:
	;
	goto L537
L537:
	;
	v2888 = v2829 + int32(1)
	v2889 = *(*int32)(unsafe.Add(mBase, uint32(v2805)+4))
	if v2888 < v2889 {
		v2829 = v2888
		goto L523
	} else {
		goto L539
	}
L538:
	;
	goto L537
L539:
	;
	goto L524
L540:
	;
	v2946 = *(*int32)(unsafe.Add(mBase, uint32(v2925)+16))
	if v2946 == int32(0) {
		goto L550
	} else {
		goto L551
	}
L541:
	;
	v2943 = l4
	v2944 = int32(0)
	v2945 = int32(1)
	goto L540
L542:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v2931 = m.ExcPending
	if v2931 != 0 {
		goto L22
	} else {
		goto L544
	}
L543:
	;
	v2943 = int32(0)
	v2944 = v2926
	v2945 = v8
	goto L540
L544:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v37))) = l4
	F_errmsg_internal(m, int32(_a_F_add_paths_to_joinrel_0), v37)
	mBase = m.M
	v2935 = m.ExcPending
	if v2935 != 0 {
		goto L22
	} else {
		goto L545
	}
L545:
	;
	F_errfinish(m, int32(_a_F_add_paths_to_joinrel_1), int32(1863), int32(_a_F_add_paths_to_joinrel_2))
	mBase = m.M
	v2940 = m.ExcPending
	if v2940 != 0 {
		goto L22
	} else {
		goto L546
	}
L546:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L547:
	;
	v3082 = *(*int32)(unsafe.Add(mBase, uint32(l2)+32))
	if v3082 == int32(0) {
		goto L592
	} else {
		goto L593
	}
L548:
	;
	v3059 = int32(0)
	if v2944 == v3059 {
		v3079 = v3058
		v3080 = v3059
		v3081 = v8
		goto L547
	} else {
		goto L586
	}
L549:
	;
	if l4 == int32(9) {
		goto L278
	} else {
		goto L585
	}
L550:
	;
	if l4 != int32(9) {
		v3058 = v2925
		goto L548
	} else {
		goto L583
	}
L551:
	;
	v2949 = *(*int32)(unsafe.Add(mBase, uint32(v2946)+4))
	v2950 = *(*int32)(unsafe.Add(mBase, uint32(l2)+8))
	v2951 = int32(0)
	if v2949 == v2951 {
		v2992 = v2951
		goto L553
	} else {
		goto L554
	}
L552:
	;
	if v2992 != 0 {
		goto L549
	} else {
		goto L566
	}
L553:
	;
	goto L552
L554:
	;
	if v2950 == int32(0) {
		v2992 = v2951
		goto L553
	} else {
		goto L555
	}
L555:
	;
	v2960 = *(*int32)(unsafe.Add(mBase, uint32(v2949)+4))
	v2961 = *(*int32)(unsafe.Add(mBase, uint32(v2950)+4))
	if v2960 < v2961 {
		goto L556
	} else {
		goto L557
	}
L556:
	;
	v2963 = v2960
	goto L558
L557:
	;
	v2963 = v2961
	goto L558
L558:
	;
	if v2963 <= int32(1) {
		goto L559
	} else {
		goto L560
	}
L559:
	;
	v2966 = int32(1)
	goto L561
L560:
	;
	v2966 = v2963
	goto L561
L561:
	;
	v2967 = int32(8)
	v2972 = int32(0)
	goto L562
L562:
	;
	v2979 = v2972 << (uint(int32(2)) % 32)
	v2981 = *(*int32)(unsafe.Add(mBase, uint32(v2950+v2967+v2979)))
	v2983 = *(*int32)(unsafe.Add(mBase, uint32(v2979+(v2949+v2967))))
	v2984 = v2981 & v2983
	v2986 = base.B2i32(v2984 != int32(0))
	if v2984 != 0 {
		v2992 = v2986
		goto L553
	} else {
		goto L564
	}
L563:
	;
	v2992 = v2986
	goto L553
L564:
	;
	v2988 = v2972 + int32(1)
	if v2988 != v2966 {
		v2972 = v2988
		goto L562
	} else {
		goto L565
	}
L565:
	;
	goto L563
L566:
	;
	v2996 = *(*int32)(unsafe.Add(mBase, uint32(v2925)+16))
	if v2996 == int32(0) {
		goto L550
	} else {
		goto L567
	}
L567:
	;
	v2999 = *(*int32)(unsafe.Add(mBase, uint32(v2996)+4))
	v3000 = *(*int32)(unsafe.Add(mBase, uint32(l2)+228))
	v3001 = int32(0)
	if v2999 == v3001 {
		v3042 = v3001
		goto L569
	} else {
		goto L570
	}
L568:
	;
	if v3042 != 0 {
		goto L549
	} else {
		goto L582
	}
L569:
	;
	goto L568
L570:
	;
	if v3000 == int32(0) {
		v3042 = v3001
		goto L569
	} else {
		goto L571
	}
L571:
	;
	v3010 = *(*int32)(unsafe.Add(mBase, uint32(v2999)+4))
	v3011 = *(*int32)(unsafe.Add(mBase, uint32(v3000)+4))
	if v3010 < v3011 {
		goto L572
	} else {
		goto L573
	}
L572:
	;
	v3013 = v3010
	goto L574
L573:
	;
	v3013 = v3011
	goto L574
L574:
	;
	if v3013 <= int32(1) {
		goto L575
	} else {
		goto L576
	}
L575:
	;
	v3016 = int32(1)
	goto L577
L576:
	;
	v3016 = v3013
	goto L577
L577:
	;
	v3017 = int32(8)
	v3022 = int32(0)
	goto L578
L578:
	;
	v3029 = v3022 << (uint(int32(2)) % 32)
	v3031 = *(*int32)(unsafe.Add(mBase, uint32(v3000+v3017+v3029)))
	v3033 = *(*int32)(unsafe.Add(mBase, uint32(v3029+(v2999+v3017))))
	v3034 = v3031 & v3033
	v3036 = base.B2i32(v3034 != int32(0))
	if v3034 != 0 {
		v3042 = v3036
		goto L569
	} else {
		goto L580
	}
L579:
	;
	v3042 = v3036
	goto L569
L580:
	;
	v3038 = v3022 + int32(1)
	if v3038 != v3016 {
		v3022 = v3038
		goto L578
	} else {
		goto L581
	}
L581:
	;
	goto L579
L582:
	;
	goto L550
L583:
	;
	v3050 = *(*int32)(unsafe.Add(mBase, uint32(v37)+20))
	v3051 = F_create_unique_path(m, l0, l3, v2925, v3050)
	mBase = m.M
	v3052 = m.ExcPending
	if v3052 != 0 {
		goto L22
	} else {
		goto L584
	}
L584:
	;
	v3079 = v3051
	v3080 = int32(1)
	v3081 = v8
	goto L547
L585:
	;
	v3058 = int32(0)
	goto L548
L586:
	;
	v3063 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_add_paths_to_joinrel[1])))
	if v3063 != int32(1) {
		v3079 = v3058
		v3080 = v3059
		v3081 = v8
		goto L547
	} else {
		goto L587
	}
L587:
	;
	if v3058 == int32(0) {
		v3079 = v3058
		v3080 = v3059
		v3081 = v8
		goto L547
	} else {
		goto L588
	}
L588:
	;
	v3068 = *(*int32)(unsafe.Add(mBase, uint32(v3058)+4))
	v3070 = v3068 - int32(348)
	goto L589
L589:
	;
	if base.B2i32(base.Ui32(v3070) < base.Ui32(int32(15)))&int32(base.Ui32(int32(_a_F_add_paths_to_joinrel_3))>>(uint(v3070)%32)) != 0 {
		v3079 = v3058
		v3080 = v3059
		v3081 = v8
		goto L547
	} else {
		goto L590
	}
L590:
	;
	v3076 = F_create_material_path(m, l3, v3058)
	mBase = m.M
	v3077 = m.ExcPending
	if v3077 != 0 {
		goto L22
	} else {
		goto L591
	}
L591:
	;
	v3079 = v3058
	v3080 = v3059
	v3081 = v3076
	goto L547
L592:
	;
	v3498 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l1)+26)))
	if v3498 != int32(1) {
		goto L278
	} else {
		goto L659
	}
L593:
	;
	v3085 = int32(0)
	v3086 = *(*int32)(unsafe.Add(mBase, uint32(v3082)+4))
	if v3086 <= v3085 {
		goto L592
	} else {
		goto L594
	}
L594:
	;
	v3091 = int32(8)
	v3110 = v3085
	goto L595
L595:
	;
	v3130 = *(*int32)(unsafe.Add(mBase, uint32(v3082)+12))
	v3134 = *(*int32)(unsafe.Add(mBase, uint32(v3130+v3110<<(uint(int32(2))%32))))
	v3135 = *(*int32)(unsafe.Add(mBase, uint32(v3134)+16))
	if v3135 == int32(0) {
		goto L598
	} else {
		goto L599
	}
L596:
	;
	goto L592
L597:
	;
	v3461 = v3110 + int32(1)
	v3462 = *(*int32)(unsafe.Add(mBase, uint32(v3082)+4))
	if v3461 < v3462 {
		v3110 = v3461
		goto L595
	} else {
		goto L658
	}
L598:
	;
	if base.B2i32(l4 != v3091) == int32(0) {
		goto L631
	} else {
		goto L632
	}
L599:
	;
	v3138 = *(*int32)(unsafe.Add(mBase, uint32(v3135)+4))
	v3139 = *(*int32)(unsafe.Add(mBase, uint32(l3)+8))
	v3140 = int32(0)
	if v3138 == v3140 {
		v3181 = v3140
		goto L601
	} else {
		goto L602
	}
L600:
	;
	if v3181 != 0 {
		goto L597
	} else {
		goto L614
	}
L601:
	;
	goto L600
L602:
	;
	if v3139 == int32(0) {
		v3181 = v3140
		goto L601
	} else {
		goto L603
	}
L603:
	;
	v3149 = *(*int32)(unsafe.Add(mBase, uint32(v3138)+4))
	v3150 = *(*int32)(unsafe.Add(mBase, uint32(v3139)+4))
	if v3149 < v3150 {
		goto L604
	} else {
		goto L605
	}
L604:
	;
	v3152 = v3149
	goto L606
L605:
	;
	v3152 = v3150
	goto L606
L606:
	;
	if v3152 <= int32(1) {
		goto L607
	} else {
		goto L608
	}
L607:
	;
	v3155 = int32(1)
	goto L609
L608:
	;
	v3155 = v3152
	goto L609
L609:
	;
	v3156 = int32(8)
	v3161 = int32(0)
	goto L610
L610:
	;
	v3168 = v3161 << (uint(int32(2)) % 32)
	v3170 = *(*int32)(unsafe.Add(mBase, uint32(v3139+v3156+v3168)))
	v3172 = *(*int32)(unsafe.Add(mBase, uint32(v3168+(v3138+v3156))))
	v3173 = v3170 & v3172
	v3175 = base.B2i32(v3173 != int32(0))
	if v3173 != 0 {
		v3181 = v3175
		goto L601
	} else {
		goto L612
	}
L611:
	;
	v3181 = v3175
	goto L601
L612:
	;
	v3177 = v3161 + int32(1)
	if v3177 != v3155 {
		v3161 = v3177
		goto L610
	} else {
		goto L613
	}
L613:
	;
	goto L611
L614:
	;
	v3185 = *(*int32)(unsafe.Add(mBase, uint32(v3134)+16))
	if v3185 == int32(0) {
		goto L598
	} else {
		goto L615
	}
L615:
	;
	v3188 = *(*int32)(unsafe.Add(mBase, uint32(v3185)+4))
	v3189 = *(*int32)(unsafe.Add(mBase, uint32(l3)+228))
	v3190 = int32(0)
	if v3188 == v3190 {
		v3231 = v3190
		goto L617
	} else {
		goto L618
	}
L616:
	;
	if v3231 != 0 {
		goto L597
	} else {
		goto L630
	}
L617:
	;
	goto L616
L618:
	;
	if v3189 == int32(0) {
		v3231 = v3190
		goto L617
	} else {
		goto L619
	}
L619:
	;
	v3199 = *(*int32)(unsafe.Add(mBase, uint32(v3188)+4))
	v3200 = *(*int32)(unsafe.Add(mBase, uint32(v3189)+4))
	if v3199 < v3200 {
		goto L620
	} else {
		goto L621
	}
L620:
	;
	v3202 = v3199
	goto L622
L621:
	;
	v3202 = v3200
	goto L622
L622:
	;
	if v3202 <= int32(1) {
		goto L623
	} else {
		goto L624
	}
L623:
	;
	v3205 = int32(1)
	goto L625
L624:
	;
	v3205 = v3202
	goto L625
L625:
	;
	v3206 = int32(8)
	v3211 = int32(0)
	goto L626
L626:
	;
	v3218 = v3211 << (uint(int32(2)) % 32)
	v3220 = *(*int32)(unsafe.Add(mBase, uint32(v3189+v3206+v3218)))
	v3222 = *(*int32)(unsafe.Add(mBase, uint32(v3218+(v3188+v3206))))
	v3223 = v3220 & v3222
	v3225 = base.B2i32(v3223 != int32(0))
	if v3223 != 0 {
		v3231 = v3225
		goto L617
	} else {
		goto L628
	}
L627:
	;
	v3231 = v3225
	goto L617
L628:
	;
	v3227 = v3211 + int32(1)
	if v3227 != v3205 {
		v3211 = v3227
		goto L626
	} else {
		goto L629
	}
L629:
	;
	goto L627
L630:
	;
	goto L598
L631:
	;
	v3238 = *(*int32)(unsafe.Add(mBase, uint32(l2)+48))
	if v3134 != v3238 {
		goto L597
	} else {
		goto L634
	}
L632:
	;
	v3243 = v3134
	goto L633
L633:
	;
	v3244 = *(*int32)(unsafe.Add(mBase, uint32(v3243)+64))
	v3245 = F_build_join_pathkeys(m, l0, l1, v2943, v3244)
	mBase = m.M
	v3246 = m.ExcPending
	if v3246 != 0 {
		goto L22
	} else {
		goto L636
	}
L634:
	;
	v3240 = *(*int32)(unsafe.Add(mBase, uint32(v37)+20))
	v3241 = F_create_unique_path(m, l0, l2, v3134, v3240)
	mBase = m.M
	v3242 = m.ExcPending
	if v3242 != 0 {
		goto L22
	} else {
		goto L635
	}
L635:
	;
	v3243 = v3241
	goto L633
L636:
	;
	if v3080 == int32(0) {
		goto L638
	} else {
		goto L639
	}
L637:
	;
	if base.B2i32(v3079 == int32(0))|base.B2i32(l4 == v3091) != 0 {
		goto L597
	} else {
		goto L656
	}
L638:
	;
	if v2944 == int32(0) {
		goto L637
	} else {
		goto L641
	}
L639:
	;
	v3357 = v3079
	goto L640
L640:
	;
	F_try_nestloop_path(m, l0, l1, v3243, v3357, v3245, v2943, v37+int32(8))
	mBase = m.M
	v3386 = m.ExcPending
	if v3386 != 0 {
		goto L22
	} else {
		goto L655
	}
L641:
	;
	v3251 = *(*int32)(unsafe.Add(mBase, uint32(l3)+56))
	if v3251 == int32(0) {
		goto L642
	} else {
		goto L643
	}
L642:
	;
	if v3081 == int32(0) {
		goto L637
	} else {
		goto L654
	}
L643:
	;
	v3254 = int32(0)
	v3255 = *(*int32)(unsafe.Add(mBase, uint32(v3251)+4))
	if v3255 <= v3254 {
		goto L642
	} else {
		goto L644
	}
L644:
	;
	v3266 = v3254
	goto L645
L645:
	;
	v3292 = *(*int32)(unsafe.Add(mBase, uint32(v3251)+12))
	v3296 = *(*int32)(unsafe.Add(mBase, uint32(v3292+v3266<<(uint(int32(2))%32))))
	F_try_nestloop_path(m, l0, l1, v3243, v3296, v3245, v2943, v37+int32(8))
	mBase = m.M
	v3300 = m.ExcPending
	if v3300 != 0 {
		goto L22
	} else {
		goto L647
	}
L646:
	;
	goto L642
L647:
	;
	v3303 = F_get_memoize_path(m, l0, l3, l2, v3296, v3243, v2943, v37+int32(8))
	mBase = m.M
	v3304 = m.ExcPending
	if v3304 != 0 {
		goto L22
	} else {
		goto L648
	}
L648:
	;
	if v3303 != 0 {
		goto L649
	} else {
		goto L650
	}
L649:
	;
	F_try_nestloop_path(m, l0, l1, v3243, v3303, v3245, v2943, v37+int32(8))
	mBase = m.M
	v3308 = m.ExcPending
	if v3308 != 0 {
		goto L22
	} else {
		goto L652
	}
L650:
	;
	goto L651
L651:
	;
	v3310 = v3266 + int32(1)
	v3311 = *(*int32)(unsafe.Add(mBase, uint32(v3251)+4))
	if v3310 < v3311 {
		v3266 = v3310
		goto L645
	} else {
		goto L653
	}
L652:
	;
	goto L651
L653:
	;
	goto L646
L654:
	;
	v3357 = v3081
	goto L640
L655:
	;
	goto L637
L656:
	;
	F_generate_mergejoin_paths(m, l0, l1, l3, v3243, l4, v37+int32(8), v2945, v3079, v3245, int32(0))
	mBase = m.M
	v3425 = m.ExcPending
	if v3425 != 0 {
		goto L22
	} else {
		goto L657
	}
L657:
	;
	goto L597
L658:
	;
	goto L596
L659:
	;
	if int32(1)<<(uint(l4)%32)&int32(396) != 0 {
		goto L660
	} else {
		goto L661
	}
L660:
	;
	v3508 = base.B2i32(base.Ui32(l4) <= base.Ui32(int32(8)))
	goto L662
L661:
	;
	v3508 = int32(0)
	goto L662
L662:
	;
	if v3508 != 0 {
		goto L278
	} else {
		goto L663
	}
L663:
	;
	v3509 = *(*int32)(unsafe.Add(mBase, uint32(l2)+40))
	if v3509 == int32(0) {
		goto L278
	} else {
		goto L664
	}
L664:
	;
	v3512 = *(*int32)(unsafe.Add(mBase, uint32(l1)+64))
	if v3512 != 0 {
		goto L278
	} else {
		goto L665
	}
L665:
	;
	if v2944 != 0 {
		goto L666
	} else {
		goto L667
	}
L666:
	;
	v3514 = v37 + int32(8)
	v3515 = int32(0)
	v3518 = base.B2i32(l4 == int32(9))
	if l4 == int32(9) {
		v3638 = v3515
		goto L669
	} else {
		goto L670
	}
L667:
	;
	goto L668
L668:
	;
	if v3079 != 0 {
		goto L741
	} else {
		goto L742
	}
L669:
	;
	v3641 = *(*int32)(unsafe.Add(mBase, uint32(l2)+40))
	if v3641 == int32(0) {
		goto L709
	} else {
		goto L710
	}
L670:
	;
	v3520 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_add_paths_to_joinrel[1])))
	if v3520 != int32(1) {
		v3638 = v3515
		goto L669
	} else {
		goto L671
	}
L671:
	;
	v3523 = *(*int32)(unsafe.Add(mBase, uint32(l3)+48))
	v3524 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v3523)+21)))
	if v3524 != int32(1) {
		v3638 = v3515
		goto L669
	} else {
		goto L672
	}
L672:
	;
	v3527 = *(*int32)(unsafe.Add(mBase, uint32(v3523)+16))
	if v3527 == int32(0) {
		goto L673
	} else {
		goto L674
	}
L673:
	;
	v3628 = *(*int32)(unsafe.Add(mBase, uint32(v3523)+4))
	v3630 = v3628 - int32(348)
	goto L706
L674:
	;
	v3530 = *(*int32)(unsafe.Add(mBase, uint32(v3527)+4))
	v3531 = *(*int32)(unsafe.Add(mBase, uint32(l2)+8))
	v3532 = int32(0)
	if v3530 == v3532 {
		v3573 = v3532
		goto L676
	} else {
		goto L677
	}
L675:
	;
	if v3573 != 0 {
		v3638 = v3515
		goto L669
	} else {
		goto L689
	}
L676:
	;
	goto L675
L677:
	;
	if v3531 == int32(0) {
		v3573 = v3532
		goto L676
	} else {
		goto L678
	}
L678:
	;
	v3541 = *(*int32)(unsafe.Add(mBase, uint32(v3530)+4))
	v3542 = *(*int32)(unsafe.Add(mBase, uint32(v3531)+4))
	if v3541 < v3542 {
		goto L679
	} else {
		goto L680
	}
L679:
	;
	v3544 = v3541
	goto L681
L680:
	;
	v3544 = v3542
	goto L681
L681:
	;
	if v3544 <= int32(1) {
		goto L682
	} else {
		goto L683
	}
L682:
	;
	v3547 = int32(1)
	goto L684
L683:
	;
	v3547 = v3544
	goto L684
L684:
	;
	v3548 = int32(8)
	v3553 = int32(0)
	goto L685
L685:
	;
	v3560 = v3553 << (uint(int32(2)) % 32)
	v3562 = *(*int32)(unsafe.Add(mBase, uint32(v3531+v3548+v3560)))
	v3564 = *(*int32)(unsafe.Add(mBase, uint32(v3560+(v3530+v3548))))
	v3565 = v3562 & v3564
	v3567 = base.B2i32(v3565 != int32(0))
	if v3565 != 0 {
		v3573 = v3567
		goto L676
	} else {
		goto L687
	}
L686:
	;
	v3573 = v3567
	goto L676
L687:
	;
	v3569 = v3553 + int32(1)
	if v3569 != v3547 {
		v3553 = v3569
		goto L685
	} else {
		goto L688
	}
L688:
	;
	goto L686
L689:
	;
	v3577 = *(*int32)(unsafe.Add(mBase, uint32(v3523)+16))
	if v3577 == int32(0) {
		goto L673
	} else {
		goto L690
	}
L690:
	;
	v3580 = *(*int32)(unsafe.Add(mBase, uint32(v3577)+4))
	v3581 = *(*int32)(unsafe.Add(mBase, uint32(l2)+228))
	v3582 = int32(0)
	if v3580 == v3582 {
		v3623 = v3582
		goto L692
	} else {
		goto L693
	}
L691:
	;
	if v3623 != 0 {
		v3638 = v3515
		goto L669
	} else {
		goto L705
	}
L692:
	;
	goto L691
L693:
	;
	if v3581 == int32(0) {
		v3623 = v3582
		goto L692
	} else {
		goto L694
	}
L694:
	;
	v3591 = *(*int32)(unsafe.Add(mBase, uint32(v3580)+4))
	v3592 = *(*int32)(unsafe.Add(mBase, uint32(v3581)+4))
	if v3591 < v3592 {
		goto L695
	} else {
		goto L696
	}
L695:
	;
	v3594 = v3591
	goto L697
L696:
	;
	v3594 = v3592
	goto L697
L697:
	;
	if v3594 <= int32(1) {
		goto L698
	} else {
		goto L699
	}
L698:
	;
	v3597 = int32(1)
	goto L700
L699:
	;
	v3597 = v3594
	goto L700
L700:
	;
	v3598 = int32(8)
	v3603 = int32(0)
	goto L701
L701:
	;
	v3610 = v3603 << (uint(int32(2)) % 32)
	v3612 = *(*int32)(unsafe.Add(mBase, uint32(v3581+v3598+v3610)))
	v3614 = *(*int32)(unsafe.Add(mBase, uint32(v3610+(v3580+v3598))))
	v3615 = v3612 & v3614
	v3617 = base.B2i32(v3615 != int32(0))
	if v3615 != 0 {
		v3623 = v3617
		goto L692
	} else {
		goto L703
	}
L702:
	;
	v3623 = v3617
	goto L692
L703:
	;
	v3619 = v3603 + int32(1)
	if v3619 != v3597 {
		v3603 = v3619
		goto L701
	} else {
		goto L704
	}
L704:
	;
	goto L702
L705:
	;
	goto L673
L706:
	;
	if base.B2i32(base.Ui32(v3630) < base.Ui32(int32(15)))&int32(base.Ui32(int32(_a_F_add_paths_to_joinrel_3))>>(uint(v3630)%32)) != 0 {
		v3638 = v3515
		goto L669
	} else {
		goto L707
	}
L707:
	;
	v3636 = F_create_material_path(m, l3, v3523)
	mBase = m.M
	v3637 = m.ExcPending
	if v3637 != 0 {
		goto L22
	} else {
		goto L708
	}
L708:
	;
	v3638 = v3636
	goto L669
L709:
	;
	goto L668
L710:
	;
	v3644 = *(*int32)(unsafe.Add(mBase, uint32(v3641)+4))
	if v3644 <= int32(0) {
		goto L709
	} else {
		goto L711
	}
L711:
	;
	if l4 == int32(9) {
		goto L712
	} else {
		goto L713
	}
L712:
	;
	v3648 = int32(0)
	goto L714
L713:
	;
	v3648 = l4
	goto L714
L714:
	;
	v3659 = v3515
	goto L715
L715:
	;
	v3685 = *(*int32)(unsafe.Add(mBase, uint32(v3641)+12))
	v3689 = *(*int32)(unsafe.Add(mBase, uint32(v3685+v3659<<(uint(int32(2))%32))))
	v3690 = *(*int32)(unsafe.Add(mBase, uint32(v3689)+64))
	v3691 = F_build_join_pathkeys(m, l0, l1, v3648, v3690)
	mBase = m.M
	v3692 = m.ExcPending
	if v3692 != 0 {
		goto L22
	} else {
		goto L717
	}
L716:
	;
	goto L709
L717:
	;
	v3693 = *(*int32)(unsafe.Add(mBase, uint32(l3)+56))
	if v3693 == int32(0) {
		goto L718
	} else {
		goto L719
	}
L718:
	;
	if v3638 != 0 {
		goto L735
	} else {
		goto L736
	}
L719:
	;
	v3696 = int32(0)
	v3697 = *(*int32)(unsafe.Add(mBase, uint32(v3693)+4))
	if v3697 <= v3696 {
		goto L718
	} else {
		goto L720
	}
L720:
	;
	v3705 = v3696
	goto L721
L721:
	;
	v3734 = *(*int32)(unsafe.Add(mBase, uint32(v3693)+12))
	v3738 = *(*int32)(unsafe.Add(mBase, uint32(v3734+v3705<<(uint(int32(2))%32))))
	v3739 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v3738)+21)))
	if v3739 != int32(1) {
		goto L723
	} else {
		goto L724
	}
L722:
	;
	goto L718
L723:
	;
	v3760 = v3705 + int32(1)
	v3761 = *(*int32)(unsafe.Add(mBase, uint32(v3693)+4))
	if v3760 < v3761 {
		v3705 = v3760
		goto L721
	} else {
		goto L734
	}
L724:
	;
	if base.B2i32(l4 != int32(9)) == int32(0) {
		goto L725
	} else {
		goto L726
	}
L725:
	;
	v3744 = *(*int32)(unsafe.Add(mBase, uint32(l3)+48))
	if v3738 != v3744 {
		goto L723
	} else {
		goto L728
	}
L726:
	;
	v3749 = v3738
	goto L727
L727:
	;
	F_try_partial_nestloop_path(m, l0, l1, v3689, v3749, v3691, v3648, v3514)
	mBase = m.M
	v3751 = m.ExcPending
	if v3751 != 0 {
		goto L22
	} else {
		goto L730
	}
L728:
	;
	v3746 = *(*int32)(unsafe.Add(mBase, uint32(v3514)+12))
	v3747 = F_create_unique_path(m, l0, l3, v3738, v3746)
	mBase = m.M
	v3748 = m.ExcPending
	if v3748 != 0 {
		goto L22
	} else {
		goto L729
	}
L729:
	;
	v3749 = v3747
	goto L727
L730:
	;
	v3752 = F_get_memoize_path(m, l0, l3, l2, v3749, v3689, v3648, v3514)
	mBase = m.M
	v3753 = m.ExcPending
	if v3753 != 0 {
		goto L22
	} else {
		goto L731
	}
L731:
	;
	if v3752 == int32(0) {
		goto L723
	} else {
		goto L732
	}
L732:
	;
	F_try_partial_nestloop_path(m, l0, l1, v3689, v3752, v3691, v3648, v3514)
	mBase = m.M
	v3757 = m.ExcPending
	if v3757 != 0 {
		goto L22
	} else {
		goto L733
	}
L733:
	;
	goto L723
L734:
	;
	goto L722
L735:
	;
	F_try_partial_nestloop_path(m, l0, l1, v3689, v3638, v3691, v3648, v3514)
	mBase = m.M
	v3798 = m.ExcPending
	if v3798 != 0 {
		goto L22
	} else {
		goto L738
	}
L736:
	;
	goto L737
L737:
	;
	v3800 = v3659 + int32(1)
	v3801 = *(*int32)(unsafe.Add(mBase, uint32(v3641)+4))
	if v3800 < v3801 {
		v3659 = v3800
		goto L715
	} else {
		goto L739
	}
L738:
	;
	goto L737
L739:
	;
	goto L716
L740:
	;
	v3924 = *(*int32)(unsafe.Add(mBase, uint32(l2)+40))
	if v3924 == int32(0) {
		goto L764
	} else {
		goto L765
	}
L741:
	;
	v3871 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v3079)+21)))
	if v3871 != 0 {
		v3921 = v3079
		goto L740
	} else {
		goto L744
	}
L742:
	;
	goto L743
L743:
	;
	if v3080 != 0 {
		goto L278
	} else {
		goto L745
	}
L744:
	;
	goto L743
L745:
	;
	v3872 = *(*int32)(unsafe.Add(mBase, uint32(l3)+32))
	if v3872 != 0 {
		goto L748
	} else {
		goto L749
	}
L746:
	;
	if v3915 == int32(0) {
		goto L278
	} else {
		goto L763
	}
L747:
	;
	goto L746
L748:
	;
	v3877 = *(*int32)(unsafe.Add(mBase, uint32(v3872)+4))
	if v3877 <= int32(0) {
		v3915 = int32(0)
		goto L747
	} else {
		goto L751
	}
L749:
	;
	goto L750
L750:
	;
	v3915 = int32(0)
	goto L747
L751:
	;
	v3880 = int32(0)
	if v3880 < v3877 {
		goto L752
	} else {
		goto L753
	}
L752:
	;
	v3883 = v3877
	goto L754
L753:
	;
	v3883 = v3880
	goto L754
L754:
	;
	v3884 = *(*int32)(unsafe.Add(mBase, uint32(v3872)+12))
	v3886 = int32(0)
	goto L755
L755:
	;
	v3894 = *(*int32)(unsafe.Add(mBase, uint32(v3884+v3886<<(uint(int32(2))%32))))
	v3895 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v3894)+21)))
	if v3895 == int32(1) {
		goto L757
	} else {
		goto L758
	}
L756:
	;
	goto L750
L757:
	;
	v3898 = *(*int32)(unsafe.Add(mBase, uint32(v3894)+16))
	if v3898 == int32(0) {
		v3915 = v3894
		goto L747
	} else {
		goto L760
	}
L758:
	;
	goto L759
L759:
	;
	v3906 = v3886 + int32(1)
	if v3906 != v3883 {
		v3886 = v3906
		goto L755
	} else {
		goto L762
	}
L760:
	;
	v3901 = *(*int32)(unsafe.Add(mBase, uint32(v3898)+4))
	if v3901 == int32(0) {
		v3915 = v3894
		goto L747
	} else {
		goto L761
	}
L761:
	;
	goto L759
L762:
	;
	goto L756
L763:
	;
	v3921 = v3915
	goto L740
L764:
	;
	goto L278
L765:
	;
	v3927 = int32(0)
	v3928 = *(*int32)(unsafe.Add(mBase, uint32(v3924)+4))
	if v3928 <= v3927 {
		goto L764
	} else {
		goto L766
	}
L766:
	;
	v3936 = v3927
	goto L767
L767:
	;
	v3965 = *(*int32)(unsafe.Add(mBase, uint32(v3924)+12))
	v3969 = *(*int32)(unsafe.Add(mBase, uint32(v3965+v3936<<(uint(int32(2))%32))))
	v3971 = *(*int32)(unsafe.Add(mBase, uint32(v3969)+64))
	v3972 = F_build_join_pathkeys(m, l0, l1, l4, v3971)
	mBase = m.M
	v3973 = m.ExcPending
	if v3973 != 0 {
		goto L22
	} else {
		goto L769
	}
L768:
	;
	goto L764
L769:
	;
	F_generate_mergejoin_paths(m, l0, l1, l3, v3969, l4, v37+int32(8), int32(0), v3921, v3972, int32(1))
	mBase = m.M
	v3976 = m.ExcPending
	if v3976 != 0 {
		goto L22
	} else {
		goto L770
	}
L770:
	;
	v3978 = v3936 + int32(1)
	v3979 = *(*int32)(unsafe.Add(mBase, uint32(v3924)+4))
	if v3978 < v3979 {
		v3936 = v3978
		goto L767
	} else {
		goto L771
	}
L771:
	;
	goto L768
L772:
	;
	v5208 = *(*int32)(unsafe.Add(mBase, uint32(l1)+168))
	if v5208 == int32(0) {
		goto L1073
	} else {
		goto L1074
	}
L773:
	;
	v4052 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_add_paths_to_joinrel[2])))
	if v4052&int32(1) == int32(0) {
		goto L772
	} else {
		goto L776
	}
L774:
	;
	goto L775
L775:
	;
	v4057 = *(*int32)(unsafe.Add(mBase, uint32(v37)+8))
	if v4057 == int32(0) {
		goto L772
	} else {
		goto L777
	}
L776:
	;
	goto L775
L777:
	;
	v4060 = int32(0)
	v4061 = *(*int32)(unsafe.Add(mBase, uint32(v4057)+4))
	if v4061 <= v4060 {
		goto L772
	} else {
		goto L778
	}
L778:
	;
	v4075 = v4060
	v4084 = int32(0)
	goto L779
L779:
	;
	v4103 = *(*int32)(unsafe.Add(mBase, uint32(v4057)+12))
	v4107 = *(*int32)(unsafe.Add(mBase, uint32(v4103+v4075<<(uint(int32(2))%32))))
	if int32(1)<<(uint(l4)%32)&int32(174) != 0 {
		goto L782
	} else {
		goto L783
	}
L780:
	;
	if v4417 == int32(0) {
		goto L772
	} else {
		goto L869
	}
L781:
	;
	v4419 = v4075 + int32(1)
	v4420 = *(*int32)(unsafe.Add(mBase, uint32(v4057)+4))
	if v4419 < v4420 {
		v4075 = v4419
		v4084 = v4417
		goto L779
	} else {
		goto L868
	}
L782:
	;
	v4108 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v4107)+8)))
	if v4108 != 0 {
		v4417 = v4084
		goto L781
	} else {
		goto L785
	}
L783:
	;
	goto L784
L784:
	;
	v4167 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v4107)+9)))
	if v4167 != int32(1) {
		v4417 = v4084
		goto L781
	} else {
		goto L801
	}
L785:
	;
	v4109 = *(*int32)(unsafe.Add(mBase, uint32(v4107)+32))
	v4110 = *(*int32)(unsafe.Add(mBase, uint32(l1)+8))
	v4111 = int32(0)
	if v4109 == v4111 {
		goto L787
	} else {
		goto L788
	}
L786:
	;
	if v4164 == int32(0) {
		v4417 = v4084
		goto L781
	} else {
		goto L800
	}
L787:
	;
	v4164 = int32(1)
	goto L786
L788:
	;
	goto L789
L789:
	;
	if v4110 == int32(0) {
		v4155 = v4111
		goto L790
	} else {
		goto L791
	}
L790:
	;
	v4164 = v4155
	goto L786
L791:
	;
	v4120 = *(*int32)(unsafe.Add(mBase, uint32(v4109)+4))
	v4121 = *(*int32)(unsafe.Add(mBase, uint32(v4110)+4))
	if v4121 < v4120 {
		v4155 = v4111
		goto L790
	} else {
		goto L792
	}
L792:
	;
	v4123 = int32(1)
	if v4120 <= v4123 {
		goto L793
	} else {
		goto L794
	}
L793:
	;
	v4126 = v4123
	goto L795
L794:
	;
	v4126 = v4120
	goto L795
L795:
	;
	v4127 = int32(8)
	v4132 = int32(0)
	goto L796
L796:
	;
	v4139 = v4132 << (uint(int32(2)) % 32)
	v4141 = *(*int32)(unsafe.Add(mBase, uint32(v4109+v4127+v4139)))
	v4143 = *(*int32)(unsafe.Add(mBase, uint32(v4139+(v4110+v4127))))
	v4146 = v4141 & (v4143 ^ int32(-1))
	v4148 = base.B2i32(v4146 == int32(0))
	if v4146 != 0 {
		v4155 = v4148
		goto L790
	} else {
		goto L798
	}
L797:
	;
	v4155 = v4148
	goto L790
L798:
	;
	v4150 = v4132 + int32(1)
	if v4150 != v4126 {
		v4132 = v4150
		goto L796
	} else {
		goto L799
	}
L799:
	;
	goto L797
L800:
	;
	goto L784
L801:
	;
	v4170 = *(*int32)(unsafe.Add(mBase, uint32(v4107)+124))
	if v4170 == int32(0) {
		v4417 = v4084
		goto L781
	} else {
		goto L802
	}
L802:
	;
	v4173 = *(*int32)(unsafe.Add(mBase, uint32(l3)+8))
	v4174 = *(*int32)(unsafe.Add(mBase, uint32(v4107)+44))
	v4175 = *(*int32)(unsafe.Add(mBase, uint32(l2)+8))
	v4176 = int32(0)
	if v4174 == v4176 {
		goto L806
	} else {
		goto L807
	}
L803:
	;
	v4413 = F_lappend(m, v4084, v4107)
	mBase = m.M
	v4414 = m.ExcPending
	if v4414 != 0 {
		goto L22
	} else {
		goto L867
	}
L804:
	;
	v4291 = *(*int32)(unsafe.Add(mBase, uint32(v4107)+44))
	v4292 = int32(0)
	if v4291 == v4292 {
		goto L836
	} else {
		goto L837
	}
L805:
	;
	if v4229 == int32(0) {
		goto L804
	} else {
		goto L819
	}
L806:
	;
	v4229 = int32(1)
	goto L805
L807:
	;
	goto L808
L808:
	;
	if v4175 == int32(0) {
		v4220 = v4176
		goto L809
	} else {
		goto L810
	}
L809:
	;
	v4229 = v4220
	goto L805
L810:
	;
	v4185 = *(*int32)(unsafe.Add(mBase, uint32(v4174)+4))
	v4186 = *(*int32)(unsafe.Add(mBase, uint32(v4175)+4))
	if v4186 < v4185 {
		v4220 = v4176
		goto L809
	} else {
		goto L811
	}
L811:
	;
	v4188 = int32(1)
	if v4185 <= v4188 {
		goto L812
	} else {
		goto L813
	}
L812:
	;
	v4191 = v4188
	goto L814
L813:
	;
	v4191 = v4185
	goto L814
L814:
	;
	v4192 = int32(8)
	v4197 = int32(0)
	goto L815
L815:
	;
	v4204 = v4197 << (uint(int32(2)) % 32)
	v4206 = *(*int32)(unsafe.Add(mBase, uint32(v4174+v4192+v4204)))
	v4208 = *(*int32)(unsafe.Add(mBase, uint32(v4204+(v4175+v4192))))
	v4211 = v4206 & (v4208 ^ int32(-1))
	v4213 = base.B2i32(v4211 == int32(0))
	if v4211 != 0 {
		v4220 = v4213
		goto L809
	} else {
		goto L817
	}
L816:
	;
	v4220 = v4213
	goto L809
L817:
	;
	v4215 = v4197 + int32(1)
	if v4215 != v4191 {
		v4197 = v4215
		goto L815
	} else {
		goto L818
	}
L818:
	;
	goto L816
L819:
	;
	v4232 = *(*int32)(unsafe.Add(mBase, uint32(v4107)+48))
	v4233 = int32(0)
	if v4232 == v4233 {
		goto L821
	} else {
		goto L822
	}
L820:
	;
	if v4286 == int32(0) {
		goto L804
	} else {
		goto L834
	}
L821:
	;
	v4286 = int32(1)
	goto L820
L822:
	;
	goto L823
L823:
	;
	if v4173 == int32(0) {
		v4277 = v4233
		goto L824
	} else {
		goto L825
	}
L824:
	;
	v4286 = v4277
	goto L820
L825:
	;
	v4242 = *(*int32)(unsafe.Add(mBase, uint32(v4232)+4))
	v4243 = *(*int32)(unsafe.Add(mBase, uint32(v4173)+4))
	if v4243 < v4242 {
		v4277 = v4233
		goto L824
	} else {
		goto L826
	}
L826:
	;
	v4245 = int32(1)
	if v4242 <= v4245 {
		goto L827
	} else {
		goto L828
	}
L827:
	;
	v4248 = v4245
	goto L829
L828:
	;
	v4248 = v4242
	goto L829
L829:
	;
	v4249 = int32(8)
	v4254 = int32(0)
	goto L830
L830:
	;
	v4261 = v4254 << (uint(int32(2)) % 32)
	v4263 = *(*int32)(unsafe.Add(mBase, uint32(v4232+v4249+v4261)))
	v4265 = *(*int32)(unsafe.Add(mBase, uint32(v4261+(v4173+v4249))))
	v4268 = v4263 & (v4265 ^ int32(-1))
	v4270 = base.B2i32(v4268 == int32(0))
	if v4268 != 0 {
		v4277 = v4270
		goto L824
	} else {
		goto L832
	}
L831:
	;
	v4277 = v4270
	goto L824
L832:
	;
	v4272 = v4254 + int32(1)
	if v4272 != v4248 {
		v4254 = v4272
		goto L830
	} else {
		goto L833
	}
L833:
	;
	goto L831
L834:
	;
	v4289 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(v4107)+120)) = uint8(v4289)
	goto L803
L835:
	;
	if v4345 == int32(0) {
		v4417 = v4084
		goto L781
	} else {
		goto L849
	}
L836:
	;
	v4345 = int32(1)
	goto L835
L837:
	;
	goto L838
L838:
	;
	if v4173 == int32(0) {
		v4336 = v4292
		goto L839
	} else {
		goto L840
	}
L839:
	;
	v4345 = v4336
	goto L835
L840:
	;
	v4301 = *(*int32)(unsafe.Add(mBase, uint32(v4291)+4))
	v4302 = *(*int32)(unsafe.Add(mBase, uint32(v4173)+4))
	if v4302 < v4301 {
		v4336 = v4292
		goto L839
	} else {
		goto L841
	}
L841:
	;
	v4304 = int32(1)
	if v4301 <= v4304 {
		goto L842
	} else {
		goto L843
	}
L842:
	;
	v4307 = v4304
	goto L844
L843:
	;
	v4307 = v4301
	goto L844
L844:
	;
	v4308 = int32(8)
	v4313 = int32(0)
	goto L845
L845:
	;
	v4320 = v4313 << (uint(int32(2)) % 32)
	v4322 = *(*int32)(unsafe.Add(mBase, uint32(v4291+v4308+v4320)))
	v4324 = *(*int32)(unsafe.Add(mBase, uint32(v4320+(v4173+v4308))))
	v4327 = v4322 & (v4324 ^ int32(-1))
	v4329 = base.B2i32(v4327 == int32(0))
	if v4327 != 0 {
		v4336 = v4329
		goto L839
	} else {
		goto L847
	}
L846:
	;
	v4336 = v4329
	goto L839
L847:
	;
	v4331 = v4313 + int32(1)
	if v4331 != v4307 {
		v4313 = v4331
		goto L845
	} else {
		goto L848
	}
L848:
	;
	goto L846
L849:
	;
	v4348 = *(*int32)(unsafe.Add(mBase, uint32(v4107)+48))
	v4349 = int32(0)
	if v4348 == v4349 {
		goto L851
	} else {
		goto L852
	}
L850:
	;
	if v4402 == int32(0) {
		v4417 = v4084
		goto L781
	} else {
		goto L864
	}
L851:
	;
	v4402 = int32(1)
	goto L850
L852:
	;
	goto L853
L853:
	;
	if v4175 == int32(0) {
		v4393 = v4349
		goto L854
	} else {
		goto L855
	}
L854:
	;
	v4402 = v4393
	goto L850
L855:
	;
	v4358 = *(*int32)(unsafe.Add(mBase, uint32(v4348)+4))
	v4359 = *(*int32)(unsafe.Add(mBase, uint32(v4175)+4))
	if v4359 < v4358 {
		v4393 = v4349
		goto L854
	} else {
		goto L856
	}
L856:
	;
	v4361 = int32(1)
	if v4358 <= v4361 {
		goto L857
	} else {
		goto L858
	}
L857:
	;
	v4364 = v4361
	goto L859
L858:
	;
	v4364 = v4358
	goto L859
L859:
	;
	v4365 = int32(8)
	v4370 = int32(0)
	goto L860
L860:
	;
	v4377 = v4370 << (uint(int32(2)) % 32)
	v4379 = *(*int32)(unsafe.Add(mBase, uint32(v4348+v4365+v4377)))
	v4381 = *(*int32)(unsafe.Add(mBase, uint32(v4377+(v4175+v4365))))
	v4384 = v4379 & (v4381 ^ int32(-1))
	v4386 = base.B2i32(v4384 == int32(0))
	if v4384 != 0 {
		v4393 = v4386
		goto L854
	} else {
		goto L862
	}
L861:
	;
	v4393 = v4386
	goto L854
L862:
	;
	v4388 = v4370 + int32(1)
	if v4388 != v4364 {
		v4370 = v4388
		goto L860
	} else {
		goto L863
	}
L863:
	;
	goto L861
L864:
	;
	v4405 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v4107)+120)) = uint8(v4405)
	v4407 = *(*int32)(unsafe.Add(mBase, uint32(v4107)+4))
	v4408 = *(*int32)(unsafe.Add(mBase, uint32(v4407)+4))
	v4409 = F_get_commutator(m, v4408)
	mBase = m.M
	v4410 = m.ExcPending
	if v4410 != 0 {
		goto L22
	} else {
		goto L865
	}
L865:
	;
	if v4409 == int32(0) {
		v4417 = v4084
		goto L781
	} else {
		goto L866
	}
L866:
	;
	goto L803
L867:
	;
	v4417 = v4413
	goto L781
L868:
	;
	goto L780
L869:
	;
	v4424 = *(*int32)(unsafe.Add(mBase, uint32(l3)+48))
	v4425 = *(*int32)(unsafe.Add(mBase, uint32(l2)+44))
	v4426 = *(*int32)(unsafe.Add(mBase, uint32(l2)+48))
	v4427 = *(*int32)(unsafe.Add(mBase, uint32(v4426)+16))
	if v4427 == int32(0) {
		goto L870
	} else {
		goto L871
	}
L870:
	;
	v4528 = *(*int32)(unsafe.Add(mBase, uint32(v4424)+16))
	if v4528 == int32(0) {
		goto L903
	} else {
		goto L904
	}
L871:
	;
	v4430 = *(*int32)(unsafe.Add(mBase, uint32(v4427)+4))
	v4431 = *(*int32)(unsafe.Add(mBase, uint32(l3)+8))
	v4432 = int32(0)
	if v4430 == v4432 {
		v4473 = v4432
		goto L873
	} else {
		goto L874
	}
L872:
	;
	if v4473 != 0 {
		goto L772
	} else {
		goto L886
	}
L873:
	;
	goto L872
L874:
	;
	if v4431 == int32(0) {
		v4473 = v4432
		goto L873
	} else {
		goto L875
	}
L875:
	;
	v4441 = *(*int32)(unsafe.Add(mBase, uint32(v4430)+4))
	v4442 = *(*int32)(unsafe.Add(mBase, uint32(v4431)+4))
	if v4441 < v4442 {
		goto L876
	} else {
		goto L877
	}
L876:
	;
	v4444 = v4441
	goto L878
L877:
	;
	v4444 = v4442
	goto L878
L878:
	;
	if v4444 <= int32(1) {
		goto L879
	} else {
		goto L880
	}
L879:
	;
	v4447 = int32(1)
	goto L881
L880:
	;
	v4447 = v4444
	goto L881
L881:
	;
	v4448 = int32(8)
	v4453 = int32(0)
	goto L882
L882:
	;
	v4460 = v4453 << (uint(int32(2)) % 32)
	v4462 = *(*int32)(unsafe.Add(mBase, uint32(v4431+v4448+v4460)))
	v4464 = *(*int32)(unsafe.Add(mBase, uint32(v4460+(v4430+v4448))))
	v4465 = v4462 & v4464
	v4467 = base.B2i32(v4465 != int32(0))
	if v4465 != 0 {
		v4473 = v4467
		goto L873
	} else {
		goto L884
	}
L883:
	;
	v4473 = v4467
	goto L873
L884:
	;
	v4469 = v4453 + int32(1)
	if v4469 != v4447 {
		v4453 = v4469
		goto L882
	} else {
		goto L885
	}
L885:
	;
	goto L883
L886:
	;
	v4477 = *(*int32)(unsafe.Add(mBase, uint32(v4426)+16))
	if v4477 == int32(0) {
		goto L870
	} else {
		goto L887
	}
L887:
	;
	v4480 = *(*int32)(unsafe.Add(mBase, uint32(v4477)+4))
	v4481 = *(*int32)(unsafe.Add(mBase, uint32(l3)+228))
	v4482 = int32(0)
	if v4480 == v4482 {
		v4523 = v4482
		goto L889
	} else {
		goto L890
	}
L888:
	;
	if v4523 != 0 {
		goto L772
	} else {
		goto L902
	}
L889:
	;
	goto L888
L890:
	;
	if v4481 == int32(0) {
		v4523 = v4482
		goto L889
	} else {
		goto L891
	}
L891:
	;
	v4491 = *(*int32)(unsafe.Add(mBase, uint32(v4480)+4))
	v4492 = *(*int32)(unsafe.Add(mBase, uint32(v4481)+4))
	if v4491 < v4492 {
		goto L892
	} else {
		goto L893
	}
L892:
	;
	v4494 = v4491
	goto L894
L893:
	;
	v4494 = v4492
	goto L894
L894:
	;
	if v4494 <= int32(1) {
		goto L895
	} else {
		goto L896
	}
L895:
	;
	v4497 = int32(1)
	goto L897
L896:
	;
	v4497 = v4494
	goto L897
L897:
	;
	v4498 = int32(8)
	v4503 = int32(0)
	goto L898
L898:
	;
	v4510 = v4503 << (uint(int32(2)) % 32)
	v4512 = *(*int32)(unsafe.Add(mBase, uint32(v4481+v4498+v4510)))
	v4514 = *(*int32)(unsafe.Add(mBase, uint32(v4510+(v4480+v4498))))
	v4515 = v4512 & v4514
	v4517 = base.B2i32(v4515 != int32(0))
	if v4515 != 0 {
		v4523 = v4517
		goto L889
	} else {
		goto L900
	}
L899:
	;
	v4523 = v4517
	goto L889
L900:
	;
	v4519 = v4503 + int32(1)
	if v4519 != v4497 {
		v4503 = v4519
		goto L898
	} else {
		goto L901
	}
L901:
	;
	goto L899
L902:
	;
	goto L870
L903:
	;
	switch l4 - int32(8) {
	case 0:
		goto L939
	case 1:
		goto L938
	default:
		goto L937
	}
L904:
	;
	v4531 = *(*int32)(unsafe.Add(mBase, uint32(v4528)+4))
	v4532 = *(*int32)(unsafe.Add(mBase, uint32(l2)+8))
	v4533 = int32(0)
	if v4531 == v4533 {
		v4574 = v4533
		goto L906
	} else {
		goto L907
	}
L905:
	;
	if v4574 != 0 {
		goto L772
	} else {
		goto L919
	}
L906:
	;
	goto L905
L907:
	;
	if v4532 == int32(0) {
		v4574 = v4533
		goto L906
	} else {
		goto L908
	}
L908:
	;
	v4542 = *(*int32)(unsafe.Add(mBase, uint32(v4531)+4))
	v4543 = *(*int32)(unsafe.Add(mBase, uint32(v4532)+4))
	if v4542 < v4543 {
		goto L909
	} else {
		goto L910
	}
L909:
	;
	v4545 = v4542
	goto L911
L910:
	;
	v4545 = v4543
	goto L911
L911:
	;
	if v4545 <= int32(1) {
		goto L912
	} else {
		goto L913
	}
L912:
	;
	v4548 = int32(1)
	goto L914
L913:
	;
	v4548 = v4545
	goto L914
L914:
	;
	v4549 = int32(8)
	v4554 = int32(0)
	goto L915
L915:
	;
	v4561 = v4554 << (uint(int32(2)) % 32)
	v4563 = *(*int32)(unsafe.Add(mBase, uint32(v4532+v4549+v4561)))
	v4565 = *(*int32)(unsafe.Add(mBase, uint32(v4561+(v4531+v4549))))
	v4566 = v4563 & v4565
	v4568 = base.B2i32(v4566 != int32(0))
	if v4566 != 0 {
		v4574 = v4568
		goto L906
	} else {
		goto L917
	}
L916:
	;
	v4574 = v4568
	goto L906
L917:
	;
	v4570 = v4554 + int32(1)
	if v4570 != v4548 {
		v4554 = v4570
		goto L915
	} else {
		goto L918
	}
L918:
	;
	goto L916
L919:
	;
	v4578 = *(*int32)(unsafe.Add(mBase, uint32(v4424)+16))
	if v4578 == int32(0) {
		goto L903
	} else {
		goto L920
	}
L920:
	;
	v4581 = *(*int32)(unsafe.Add(mBase, uint32(v4578)+4))
	v4582 = *(*int32)(unsafe.Add(mBase, uint32(l2)+228))
	v4583 = int32(0)
	if v4581 == v4583 {
		v4624 = v4583
		goto L922
	} else {
		goto L923
	}
L921:
	;
	if v4624 != 0 {
		goto L772
	} else {
		goto L935
	}
L922:
	;
	goto L921
L923:
	;
	if v4582 == int32(0) {
		v4624 = v4583
		goto L922
	} else {
		goto L924
	}
L924:
	;
	v4592 = *(*int32)(unsafe.Add(mBase, uint32(v4581)+4))
	v4593 = *(*int32)(unsafe.Add(mBase, uint32(v4582)+4))
	if v4592 < v4593 {
		goto L925
	} else {
		goto L926
	}
L925:
	;
	v4595 = v4592
	goto L927
L926:
	;
	v4595 = v4593
	goto L927
L927:
	;
	if v4595 <= int32(1) {
		goto L928
	} else {
		goto L929
	}
L928:
	;
	v4598 = int32(1)
	goto L930
L929:
	;
	v4598 = v4595
	goto L930
L930:
	;
	v4599 = int32(8)
	v4604 = int32(0)
	goto L931
L931:
	;
	v4611 = v4604 << (uint(int32(2)) % 32)
	v4613 = *(*int32)(unsafe.Add(mBase, uint32(v4582+v4599+v4611)))
	v4615 = *(*int32)(unsafe.Add(mBase, uint32(v4611+(v4581+v4599))))
	v4616 = v4613 & v4615
	v4618 = base.B2i32(v4616 != int32(0))
	if v4616 != 0 {
		v4624 = v4618
		goto L922
	} else {
		goto L933
	}
L932:
	;
	v4624 = v4618
	goto L922
L933:
	;
	v4620 = v4604 + int32(1)
	if v4620 != v4598 {
		v4604 = v4620
		goto L931
	} else {
		goto L934
	}
L934:
	;
	goto L932
L935:
	;
	goto L903
L936:
	;
	v5078 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l1)+26)))
	if v5078 != int32(1) {
		goto L772
	} else {
		goto L1035
	}
L937:
	;
	if v4425 != 0 {
		goto L947
	} else {
		goto L948
	}
L938:
	;
	v4640 = int32(0)
	v4641 = *(*int32)(unsafe.Add(mBase, uint32(v37)+20))
	v4642 = F_create_unique_path(m, l0, l3, v4424, v4641)
	mBase = m.M
	v4643 = m.ExcPending
	if v4643 != 0 {
		goto L22
	} else {
		goto L942
	}
L939:
	;
	v4632 = *(*int32)(unsafe.Add(mBase, uint32(v37)+20))
	v4633 = F_create_unique_path(m, l0, l2, v4426, v4632)
	mBase = m.M
	v4634 = m.ExcPending
	if v4634 != 0 {
		goto L22
	} else {
		goto L940
	}
L940:
	;
	F_try_hashjoin_path(m, l0, l1, v4633, v4424, v4417, int32(0), v37+int32(8))
	mBase = m.M
	v4639 = m.ExcPending
	if v4639 != 0 {
		goto L22
	} else {
		goto L941
	}
L941:
	;
	v5050 = int32(0)
	v5057 = v4424
	goto L936
L942:
	;
	F_try_hashjoin_path(m, l0, l1, v4426, v4642, v4417, int32(0), v37+int32(8))
	mBase = m.M
	v4648 = m.ExcPending
	if v4648 != 0 {
		goto L22
	} else {
		goto L943
	}
L943:
	;
	if v4425 == int32(0) {
		v5050 = v4640
		v5057 = v4642
		goto L936
	} else {
		goto L944
	}
L944:
	;
	if v4425 == v4426 {
		v5050 = v4640
		v5057 = v4642
		goto L936
	} else {
		goto L945
	}
L945:
	;
	F_try_hashjoin_path(m, l0, l1, v4425, v4642, v4417, int32(0), v37+int32(8))
	mBase = m.M
	v4656 = m.ExcPending
	if v4656 != 0 {
		goto L22
	} else {
		goto L946
	}
L946:
	;
	v5050 = v4640
	v5057 = v4642
	goto L936
L947:
	;
	F_try_hashjoin_path(m, l0, l1, v4425, v4424, v4417, l4, v37+int32(8))
	mBase = m.M
	v4660 = m.ExcPending
	if v4660 != 0 {
		goto L22
	} else {
		goto L950
	}
L948:
	;
	goto L949
L949:
	;
	v4661 = *(*int32)(unsafe.Add(mBase, uint32(l2)+56))
	if v4661 == int32(0) {
		goto L951
	} else {
		goto L952
	}
L950:
	;
	goto L949
L951:
	;
	v5050 = l4
	v5057 = v4424
	goto L936
L952:
	;
	v4664 = *(*int32)(unsafe.Add(mBase, uint32(v4661)+4))
	if v4664 <= int32(0) {
		goto L951
	} else {
		goto L953
	}
L953:
	;
	v4679 = int32(0)
	goto L954
L954:
	;
	v4702 = *(*int32)(unsafe.Add(mBase, uint32(v4661)+12))
	v4706 = *(*int32)(unsafe.Add(mBase, uint32(v4702+v4679<<(uint(int32(2))%32))))
	v4707 = *(*int32)(unsafe.Add(mBase, uint32(v4706)+16))
	if v4707 == int32(0) {
		goto L957
	} else {
		goto L958
	}
L955:
	;
	goto L951
L956:
	;
	v5007 = v4679 + int32(1)
	v5008 = *(*int32)(unsafe.Add(mBase, uint32(v4661)+4))
	if v5007 < v5008 {
		v4679 = v5007
		goto L954
	} else {
		goto L1034
	}
L957:
	;
	v4808 = *(*int32)(unsafe.Add(mBase, uint32(l3)+56))
	if v4808 == int32(0) {
		goto L956
	} else {
		goto L990
	}
L958:
	;
	v4710 = *(*int32)(unsafe.Add(mBase, uint32(v4707)+4))
	v4711 = *(*int32)(unsafe.Add(mBase, uint32(l3)+8))
	v4712 = int32(0)
	if v4710 == v4712 {
		v4753 = v4712
		goto L960
	} else {
		goto L961
	}
L959:
	;
	if v4753 != 0 {
		goto L956
	} else {
		goto L973
	}
L960:
	;
	goto L959
L961:
	;
	if v4711 == int32(0) {
		v4753 = v4712
		goto L960
	} else {
		goto L962
	}
L962:
	;
	v4721 = *(*int32)(unsafe.Add(mBase, uint32(v4710)+4))
	v4722 = *(*int32)(unsafe.Add(mBase, uint32(v4711)+4))
	if v4721 < v4722 {
		goto L963
	} else {
		goto L964
	}
L963:
	;
	v4724 = v4721
	goto L965
L964:
	;
	v4724 = v4722
	goto L965
L965:
	;
	if v4724 <= int32(1) {
		goto L966
	} else {
		goto L967
	}
L966:
	;
	v4727 = int32(1)
	goto L968
L967:
	;
	v4727 = v4724
	goto L968
L968:
	;
	v4728 = int32(8)
	v4733 = int32(0)
	goto L969
L969:
	;
	v4740 = v4733 << (uint(int32(2)) % 32)
	v4742 = *(*int32)(unsafe.Add(mBase, uint32(v4711+v4728+v4740)))
	v4744 = *(*int32)(unsafe.Add(mBase, uint32(v4740+(v4710+v4728))))
	v4745 = v4742 & v4744
	v4747 = base.B2i32(v4745 != int32(0))
	if v4745 != 0 {
		v4753 = v4747
		goto L960
	} else {
		goto L971
	}
L970:
	;
	v4753 = v4747
	goto L960
L971:
	;
	v4749 = v4733 + int32(1)
	if v4749 != v4727 {
		v4733 = v4749
		goto L969
	} else {
		goto L972
	}
L972:
	;
	goto L970
L973:
	;
	v4757 = *(*int32)(unsafe.Add(mBase, uint32(v4706)+16))
	if v4757 == int32(0) {
		goto L957
	} else {
		goto L974
	}
L974:
	;
	v4760 = *(*int32)(unsafe.Add(mBase, uint32(v4757)+4))
	v4761 = *(*int32)(unsafe.Add(mBase, uint32(l3)+228))
	v4762 = int32(0)
	if v4760 == v4762 {
		v4803 = v4762
		goto L976
	} else {
		goto L977
	}
L975:
	;
	if v4803 != 0 {
		goto L956
	} else {
		goto L989
	}
L976:
	;
	goto L975
L977:
	;
	if v4761 == int32(0) {
		v4803 = v4762
		goto L976
	} else {
		goto L978
	}
L978:
	;
	v4771 = *(*int32)(unsafe.Add(mBase, uint32(v4760)+4))
	v4772 = *(*int32)(unsafe.Add(mBase, uint32(v4761)+4))
	if v4771 < v4772 {
		goto L979
	} else {
		goto L980
	}
L979:
	;
	v4774 = v4771
	goto L981
L980:
	;
	v4774 = v4772
	goto L981
L981:
	;
	if v4774 <= int32(1) {
		goto L982
	} else {
		goto L983
	}
L982:
	;
	v4777 = int32(1)
	goto L984
L983:
	;
	v4777 = v4774
	goto L984
L984:
	;
	v4778 = int32(8)
	v4783 = int32(0)
	goto L985
L985:
	;
	v4790 = v4783 << (uint(int32(2)) % 32)
	v4792 = *(*int32)(unsafe.Add(mBase, uint32(v4761+v4778+v4790)))
	v4794 = *(*int32)(unsafe.Add(mBase, uint32(v4790+(v4760+v4778))))
	v4795 = v4792 & v4794
	v4797 = base.B2i32(v4795 != int32(0))
	if v4795 != 0 {
		v4803 = v4797
		goto L976
	} else {
		goto L987
	}
L986:
	;
	v4803 = v4797
	goto L976
L987:
	;
	v4799 = v4783 + int32(1)
	if v4799 != v4777 {
		v4783 = v4799
		goto L985
	} else {
		goto L988
	}
L988:
	;
	goto L986
L989:
	;
	goto L957
L990:
	;
	v4811 = int32(0)
	v4812 = *(*int32)(unsafe.Add(mBase, uint32(v4808)+4))
	if v4812 <= v4811 {
		goto L956
	} else {
		goto L991
	}
L991:
	;
	v4822 = v4811
	goto L992
L992:
	;
	v4849 = *(*int32)(unsafe.Add(mBase, uint32(v4808)+12))
	v4853 = *(*int32)(unsafe.Add(mBase, uint32(v4849+v4822<<(uint(int32(2))%32))))
	v4854 = *(*int32)(unsafe.Add(mBase, uint32(v4853)+16))
	if v4854 == int32(0) {
		goto L996
	} else {
		goto L997
	}
L993:
	;
	goto L956
L994:
	;
	v4969 = v4822 + int32(1)
	v4970 = *(*int32)(unsafe.Add(mBase, uint32(v4808)+4))
	if v4969 < v4970 {
		v4822 = v4969
		goto L992
	} else {
		goto L1033
	}
L995:
	;
	F_try_hashjoin_path(m, l0, l1, v4706, v4853, v4417, l4, v37+int32(8))
	mBase = m.M
	v4966 = m.ExcPending
	if v4966 != 0 {
		goto L22
	} else {
		goto L1032
	}
L996:
	;
	if v4425 != v4706 {
		goto L995
	} else {
		goto L1030
	}
L997:
	;
	v4857 = *(*int32)(unsafe.Add(mBase, uint32(v4854)+4))
	v4858 = *(*int32)(unsafe.Add(mBase, uint32(l2)+8))
	v4859 = int32(0)
	if v4857 == v4859 {
		v4900 = v4859
		goto L999
	} else {
		goto L1000
	}
L998:
	;
	if v4900 != 0 {
		goto L994
	} else {
		goto L1012
	}
L999:
	;
	goto L998
L1000:
	;
	if v4858 == int32(0) {
		v4900 = v4859
		goto L999
	} else {
		goto L1001
	}
L1001:
	;
	v4868 = *(*int32)(unsafe.Add(mBase, uint32(v4857)+4))
	v4869 = *(*int32)(unsafe.Add(mBase, uint32(v4858)+4))
	if v4868 < v4869 {
		goto L1002
	} else {
		goto L1003
	}
L1002:
	;
	v4871 = v4868
	goto L1004
L1003:
	;
	v4871 = v4869
	goto L1004
L1004:
	;
	if v4871 <= int32(1) {
		goto L1005
	} else {
		goto L1006
	}
L1005:
	;
	v4874 = int32(1)
	goto L1007
L1006:
	;
	v4874 = v4871
	goto L1007
L1007:
	;
	v4875 = int32(8)
	v4880 = int32(0)
	goto L1008
L1008:
	;
	v4887 = v4880 << (uint(int32(2)) % 32)
	v4889 = *(*int32)(unsafe.Add(mBase, uint32(v4858+v4875+v4887)))
	v4891 = *(*int32)(unsafe.Add(mBase, uint32(v4887+(v4857+v4875))))
	v4892 = v4889 & v4891
	v4894 = base.B2i32(v4892 != int32(0))
	if v4892 != 0 {
		v4900 = v4894
		goto L999
	} else {
		goto L1010
	}
L1009:
	;
	v4900 = v4894
	goto L999
L1010:
	;
	v4896 = v4880 + int32(1)
	if v4896 != v4874 {
		v4880 = v4896
		goto L1008
	} else {
		goto L1011
	}
L1011:
	;
	goto L1009
L1012:
	;
	v4904 = *(*int32)(unsafe.Add(mBase, uint32(v4853)+16))
	if v4904 == int32(0) {
		goto L996
	} else {
		goto L1013
	}
L1013:
	;
	v4907 = *(*int32)(unsafe.Add(mBase, uint32(v4904)+4))
	v4908 = *(*int32)(unsafe.Add(mBase, uint32(l2)+228))
	v4909 = int32(0)
	if v4907 == v4909 {
		v4950 = v4909
		goto L1015
	} else {
		goto L1016
	}
L1014:
	;
	if v4950 != 0 {
		goto L994
	} else {
		goto L1028
	}
L1015:
	;
	goto L1014
L1016:
	;
	if v4908 == int32(0) {
		v4950 = v4909
		goto L1015
	} else {
		goto L1017
	}
L1017:
	;
	v4918 = *(*int32)(unsafe.Add(mBase, uint32(v4907)+4))
	v4919 = *(*int32)(unsafe.Add(mBase, uint32(v4908)+4))
	if v4918 < v4919 {
		goto L1018
	} else {
		goto L1019
	}
L1018:
	;
	v4921 = v4918
	goto L1020
L1019:
	;
	v4921 = v4919
	goto L1020
L1020:
	;
	if v4921 <= int32(1) {
		goto L1021
	} else {
		goto L1022
	}
L1021:
	;
	v4924 = int32(1)
	goto L1023
L1022:
	;
	v4924 = v4921
	goto L1023
L1023:
	;
	v4925 = int32(8)
	v4930 = int32(0)
	goto L1024
L1024:
	;
	v4937 = v4930 << (uint(int32(2)) % 32)
	v4939 = *(*int32)(unsafe.Add(mBase, uint32(v4908+v4925+v4937)))
	v4941 = *(*int32)(unsafe.Add(mBase, uint32(v4937+(v4907+v4925))))
	v4942 = v4939 & v4941
	v4944 = base.B2i32(v4942 != int32(0))
	if v4942 != 0 {
		v4950 = v4944
		goto L1015
	} else {
		goto L1026
	}
L1025:
	;
	v4950 = v4944
	goto L1015
L1026:
	;
	v4946 = v4930 + int32(1)
	if v4946 != v4924 {
		v4930 = v4946
		goto L1024
	} else {
		goto L1027
	}
L1027:
	;
	goto L1025
L1028:
	;
	if base.B2i32(v4425 == v4706)&base.B2i32(v4853 == v4424) == int32(0) {
		goto L995
	} else {
		goto L1029
	}
L1029:
	;
	goto L994
L1030:
	;
	if v4853 == v4424 {
		goto L994
	} else {
		goto L1031
	}
L1031:
	;
	goto L995
L1032:
	;
	goto L994
L1033:
	;
	goto L993
L1034:
	;
	goto L955
L1035:
	;
	switch l4 - int32(6) {
	case 0, 2:
		goto L772
	default:
		goto L1036
	}
L1036:
	;
	v5083 = *(*int32)(unsafe.Add(mBase, uint32(l2)+40))
	if v5083 == int32(0) {
		goto L772
	} else {
		goto L1037
	}
L1037:
	;
	v5086 = *(*int32)(unsafe.Add(mBase, uint32(l1)+64))
	if v5086 != 0 {
		goto L772
	} else {
		goto L1038
	}
L1038:
	;
	v5087 = *(*int32)(unsafe.Add(mBase, uint32(v5083)+12))
	v5088 = *(*int32)(unsafe.Add(mBase, uint32(v5087)))
	if l4 == int32(9) {
		goto L1039
	} else {
		goto L1040
	}
L1039:
	;
	if int32(1)<<(uint(l4)%32)&int32(140) != 0 {
		goto L1044
	} else {
		goto L1045
	}
L1040:
	;
	v5091 = *(*int32)(unsafe.Add(mBase, uint32(l3)+40))
	if v5091 == int32(0) {
		goto L1039
	} else {
		goto L1041
	}
L1041:
	;
	v5095 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_add_paths_to_joinrel[3])))
	if v5095 != int32(1) {
		goto L1039
	} else {
		goto L1042
	}
L1042:
	;
	v5098 = *(*int32)(unsafe.Add(mBase, uint32(v5091)+12))
	v5099 = *(*int32)(unsafe.Add(mBase, uint32(v5098)))
	F_try_partial_hashjoin_path(m, l0, l1, v5088, v5099, v4417, v5050, v37+int32(8), int32(1))
	mBase = m.M
	v5104 = m.ExcPending
	if v5104 != 0 {
		goto L22
	} else {
		goto L1043
	}
L1043:
	;
	goto L1039
L1044:
	;
	v5113 = base.B2i32(base.Ui32(l4) <= base.Ui32(int32(7)))
	goto L1046
L1045:
	;
	v5113 = int32(0)
	goto L1046
L1046:
	;
	if v5113 != 0 {
		goto L772
	} else {
		goto L1047
	}
L1047:
	;
	v5115 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v5057)+21)))
	if v5115 != 0 {
		goto L1048
	} else {
		goto L1049
	}
L1048:
	;
	v5116 = v5057
	goto L1050
L1049:
	;
	v5116 = int32(0)
	goto L1050
L1050:
	;
	if l4 == int32(9) {
		v5166 = v5116
		goto L1051
	} else {
		goto L1052
	}
L1051:
	;
	if v5166 == int32(0) {
		goto L772
	} else {
		goto L1071
	}
L1052:
	;
	if v5115 != 0 {
		v5166 = v5116
		goto L1051
	} else {
		goto L1053
	}
L1053:
	;
	v5119 = *(*int32)(unsafe.Add(mBase, uint32(l3)+32))
	if v5119 != 0 {
		goto L1056
	} else {
		goto L1057
	}
L1054:
	;
	v5166 = v5162
	goto L1051
L1055:
	;
	goto L1054
L1056:
	;
	v5124 = *(*int32)(unsafe.Add(mBase, uint32(v5119)+4))
	if v5124 <= int32(0) {
		v5162 = int32(0)
		goto L1055
	} else {
		goto L1059
	}
L1057:
	;
	goto L1058
L1058:
	;
	v5162 = int32(0)
	goto L1055
L1059:
	;
	v5127 = int32(0)
	if v5127 < v5124 {
		goto L1060
	} else {
		goto L1061
	}
L1060:
	;
	v5130 = v5124
	goto L1062
L1061:
	;
	v5130 = v5127
	goto L1062
L1062:
	;
	v5131 = *(*int32)(unsafe.Add(mBase, uint32(v5119)+12))
	v5133 = int32(0)
	goto L1063
L1063:
	;
	v5141 = *(*int32)(unsafe.Add(mBase, uint32(v5131+v5133<<(uint(int32(2))%32))))
	v5142 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v5141)+21)))
	if v5142 == int32(1) {
		goto L1065
	} else {
		goto L1066
	}
L1064:
	;
	goto L1058
L1065:
	;
	v5145 = *(*int32)(unsafe.Add(mBase, uint32(v5141)+16))
	if v5145 == int32(0) {
		v5162 = v5141
		goto L1055
	} else {
		goto L1068
	}
L1066:
	;
	goto L1067
L1067:
	;
	v5153 = v5133 + int32(1)
	if v5153 != v5130 {
		v5133 = v5153
		goto L1063
	} else {
		goto L1070
	}
L1068:
	;
	v5148 = *(*int32)(unsafe.Add(mBase, uint32(v5145)+4))
	if v5148 == int32(0) {
		v5162 = v5141
		goto L1055
	} else {
		goto L1069
	}
L1069:
	;
	goto L1067
L1070:
	;
	goto L1064
L1071:
	;
	F_try_partial_hashjoin_path(m, l0, l1, v5088, v5166, v4417, v5050, v37+int32(8), int32(0))
	mBase = m.M
	v5173 = m.ExcPending
	if v5173 != 0 {
		goto L22
	} else {
		goto L1072
	}
L1072:
	;
	goto L772
L1073:
	;
	v5220 = *(*int32)(unsafe.Add(mBase, _c_F_add_paths_to_joinrel[4]))
	if v5220 != 0 {
		goto L1077
	} else {
		goto L1078
	}
L1074:
	;
	v5211 = *(*int32)(unsafe.Add(mBase, uint32(v5208)+32))
	if v5211 == int32(0) {
		goto L1073
	} else {
		goto L1075
	}
L1075:
	;
	m.T0[v5211].(func(*base.Module, int32, int32, int32, int32, int32, int32))(m, l0, l1, l2, l3, l4, v37+int32(8))
	mBase = m.M
	v5217 = m.ExcPending
	if v5217 != 0 {
		goto L22
	} else {
		goto L1076
	}
L1076:
	;
	goto L1073
L1077:
	;
	m.T0[v5220].(func(*base.Module, int32, int32, int32, int32, int32, int32))(m, l0, l1, l2, l3, l4, v37+int32(8))
	mBase = m.M
	v5224 = m.ExcPending
	if v5224 != 0 {
		goto L22
	} else {
		goto L1080
	}
L1078:
	;
	goto L1079
L1079:
	;
	m.G0 = v37 + int32(48)
	return
L1080:
	;
	goto L1079
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
	if v68 == v146 {
		goto L25
	} else {
		goto L26
	}
L25:
	;
	v154 = v143
	goto L11
L26:
	;
	goto L27
L27:
	;
	if v146 != v96 {
		v98 = v146
		v99 = v143
		v102 = v143
		goto L16
	} else {
		goto L28
	}
L28:
	;
	goto L17
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
	var v41 int32
	_ = v41
	var v50 int32
	_ = v50
	var v51 int32
	_ = v51
	var v53 int32
	_ = v53
	var v56 int32
	_ = v56
	var v57 int32
	_ = v57
	var v62 int32
	_ = v62
	var v69 int32
	_ = v69
	var v71 int32
	_ = v71
	var v73 int32
	_ = v73
	var v76 int32
	_ = v76
	var v78 int32
	_ = v78
	var v80 int32
	_ = v80
	var v85 int32
	_ = v85
	var v94 int32
	_ = v94
	var v95 int32
	_ = v95
	var v98 int32
	_ = v98
	var v99 int32
	_ = v99
	var v101 int32
	_ = v101
	var v102 int32
	_ = v102
	var v103 int32
	_ = v103
	var v104 int32
	_ = v104
	var v110 int32
	_ = v110
	var v111 int32
	_ = v111
	var v115 int32
	_ = v115
	var v120 int32
	_ = v120
	var v121 int32
	_ = v121
	var v122 int32
	_ = v122
	var v123 int32
	_ = v123
	var v124 int32
	_ = v124
	var v125 int32
	_ = v125
	var v131 int32
	_ = v131
	var v132 int32
	_ = v132
	v4 = int32(0)
	v9 = m.G0
	v11 = v9 - int32(16)
	m.G0 = v11
	if l1 == v4 {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	m.G0 = v11 + int32(16)
	return
L2:
	;
	v15 = *(*int32)(unsafe.Add(mBase, uint32(l1)+4))
	if v15 <= int32(0) {
		goto L1
	} else {
		goto L3
	}
L3:
	;
	v24 = v4
	goto L4
L4:
	;
	v26 = *(*int32)(unsafe.Add(mBase, uint32(l1)+12))
	v30 = *(*int32)(unsafe.Add(mBase, uint32(v26+v24<<(uint(int32(2))%32))))
	v31 = *(*int32)(unsafe.Add(mBase, uint32(v30)))
	if v31 != int32(319) {
		goto L7
	} else {
		goto L8
	}
L5:
	;
	goto L1
L6:
	;
	v131 = v24 + int32(1)
	v132 = *(*int32)(unsafe.Add(mBase, uint32(l1)+4))
	if v131 < v132 {
		v24 = v131
		goto L4
	} else {
		goto L36
	}
L7:
	;
	if v31 == int32(6) {
		goto L10
	} else {
		goto L11
	}
L8:
	;
	goto L9
L9:
	;
	v121 = F_find_placeholder_info(m, l0, v30)
	mBase = m.M
	v122 = m.ExcPending
	if v122 != 0 {
		goto L13
	} else {
		goto L34
	}
L10:
	;
	v36 = *(*int32)(unsafe.Add(mBase, uint32(v30)+4))
	v37 = F_find_base_rel(m, l0, v36)
	mBase = m.M
	v38 = m.ExcPending
	if v38 != 0 {
		goto L13
	} else {
		goto L14
	}
L11:
	;
	goto L12
L12:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v110 = m.ExcPending
	if v110 != 0 {
		goto L13
	} else {
		goto L31
	}
L13:
	;
	return
L14:
	;
	v39 = int32(*(*int16)(unsafe.Add(mBase, uint32(v30)+8)))
	v40 = *(*int32)(unsafe.Add(mBase, uint32(v37)+8))
	v41 = int32(0)
	if l2 == v41 {
		goto L16
	} else {
		goto L17
	}
L15:
	;
	if v94 != 0 {
		goto L6
	} else {
		goto L29
	}
L16:
	;
	v94 = int32(1)
	goto L15
L17:
	;
	goto L18
L18:
	;
	if v40 == int32(0) {
		v85 = v41
		goto L19
	} else {
		goto L20
	}
L19:
	;
	v94 = v85
	goto L15
L20:
	;
	v50 = *(*int32)(unsafe.Add(mBase, uint32(l2)+4))
	v51 = *(*int32)(unsafe.Add(mBase, uint32(v40)+4))
	if v51 < v50 {
		v85 = v41
		goto L19
	} else {
		goto L21
	}
L21:
	;
	v53 = int32(1)
	if v50 <= v53 {
		goto L22
	} else {
		goto L23
	}
L22:
	;
	v56 = v53
	goto L24
L23:
	;
	v56 = v50
	goto L24
L24:
	;
	v57 = int32(8)
	v62 = int32(0)
	goto L25
L25:
	;
	v69 = v62 << (uint(int32(2)) % 32)
	v71 = *(*int32)(unsafe.Add(mBase, uint32(l2+v57+v69)))
	v73 = *(*int32)(unsafe.Add(mBase, uint32(v69+(v40+v57))))
	v76 = v71 & (v73 ^ int32(-1))
	v78 = base.B2i32(v76 == int32(0))
	if v76 != 0 {
		v85 = v78
		goto L19
	} else {
		goto L27
	}
L26:
	;
	v85 = v78
	goto L19
L27:
	;
	v80 = v62 + int32(1)
	if v80 != v56 {
		v62 = v80
		goto L25
	} else {
		goto L28
	}
L28:
	;
	goto L26
L29:
	;
	v95 = int32(*(*int16)(unsafe.Add(mBase, uint32(v37)+80)))
	v98 = (v39 - v95) << (uint(int32(2)) % 32)
	v99 = *(*int32)(unsafe.Add(mBase, uint32(v37)+84))
	v101 = *(*int32)(unsafe.Add(mBase, uint32(v98+v99)))
	v102 = F_bms_add_members(m, v101, l2)
	mBase = m.M
	v103 = m.ExcPending
	if v103 != 0 {
		goto L13
	} else {
		goto L30
	}
L30:
	;
	v104 = *(*int32)(unsafe.Add(mBase, uint32(v37)+84))
	*(*int32)(unsafe.Add(mBase, uint32(v104+v98))) = v102
	goto L6
L31:
	;
	v111 = *(*int32)(unsafe.Add(mBase, uint32(v30)))
	*(*int32)(unsafe.Add(mBase, uint32(v11))) = v111
	F_errmsg_internal(m, int32(_a_F_add_vars_to_attr_needed_0), v11)
	mBase = m.M
	v115 = m.ExcPending
	if v115 != 0 {
		goto L13
	} else {
		goto L32
	}
L32:
	;
	F_errfinish(m, int32(_a_F_add_vars_to_attr_needed_1), int32(386), int32(_a_F_add_vars_to_attr_needed_2))
	mBase = m.M
	v120 = m.ExcPending
	if v120 != 0 {
		goto L13
	} else {
		goto L33
	}
L33:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L34:
	;
	v123 = *(*int32)(unsafe.Add(mBase, uint32(v121)+20))
	v124 = F_bms_add_members(m, v123, l2)
	mBase = m.M
	v125 = m.ExcPending
	if v125 != 0 {
		goto L13
	} else {
		goto L35
	}
L35:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v121)+20)) = v124
	goto L6
L36:
	;
	goto L5
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
	var v63 int32
	_ = v63
	var v64 int32
	_ = v64
	var v66 int32
	_ = v66
	var v68 int32
	_ = v68
	var v69 int32
	_ = v69
	var v77 int32
	_ = v77
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
					v63 = v60
					v64 = *(*int32)(unsafe.Add(mBase, uint32(v61)+24))
					if v38 != 0 {
						v66 = F__emscripten_memcpy_bulkmem(m, v63+v64, v9, v38)
						mBase = m.M
					} else {
					}
					v68 = *(*int32)(unsafe.Add(mBase, uint32(l1)))
					v69 = *(*int32)(unsafe.Add(mBase, uint32(v68)+28))
					*(*int32)(unsafe.Add(mBase, uint32(v68)+28)) = v69 + v38
					m.G0 = v9 + int32(48)
					return
				}
			} else {
				v61 = v36
				v63 = v37
				v64 = *(*int32)(unsafe.Add(mBase, uint32(v61)+24))
				if v38 != 0 {
					v66 = F__emscripten_memcpy_bulkmem(m, v63+v64, v9, v38)
					mBase = m.M
				} else {
				}
				v68 = *(*int32)(unsafe.Add(mBase, uint32(l1)))
				v69 = *(*int32)(unsafe.Add(mBase, uint32(v68)+28))
				*(*int32)(unsafe.Add(mBase, uint32(v68)+28)) = v69 + v38
				m.G0 = v9 + int32(48)
				return
			}
		}
	} else {
		F_scanner_yyerror(m, int32(_a_F_addunicode_2), l1)
		mBase = m.M
		v77 = m.ExcPending
		if v77 != 0 {
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
	var v34 int32
	_ = v34
	var v41 int32
	_ = v41
	var v43 int32
	_ = v43
	var v45 int32
	_ = v45
	var v46 int32
	_ = v46
	var v48 int32
	_ = v48
	var v50 int32
	_ = v50
	var v54 int32
	_ = v54
	var v58 int32
	_ = v58
	var v62 int32
	_ = v62
	var v65 int32
	_ = v65
	var v66 int32
	_ = v66
	var v67 int32
	_ = v67
	var v70 int32
	_ = v70
	var v71 int32
	_ = v71
	var v72 int32
	_ = v72
	var v73 int32
	_ = v73
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
	if l1 == v5 {
		v54 = v5
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
		goto L23
	} else {
		goto L47
	}
L2:
	;
	if v54 != 0 {
		goto L16
	} else {
		goto L17
	}
L3:
	;
	goto L2
L4:
	;
	if v12 == int32(0) {
		v54 = v5
		goto L3
	} else {
		goto L5
	}
L5:
	;
	v22 = *(*int32)(unsafe.Add(mBase, uint32(l1)+4))
	v23 = *(*int32)(unsafe.Add(mBase, uint32(v12)+4))
	if v22 < v23 {
		goto L6
	} else {
		goto L7
	}
L6:
	;
	v25 = v22
	goto L8
L7:
	;
	v25 = v23
	goto L8
L8:
	;
	if v25 <= int32(1) {
		goto L9
	} else {
		goto L10
	}
L9:
	;
	v28 = int32(1)
	goto L11
L10:
	;
	v28 = v25
	goto L11
L11:
	;
	v29 = int32(8)
	v34 = int32(0)
	goto L12
L12:
	;
	v41 = v34 << (uint(int32(2)) % 32)
	v43 = *(*int32)(unsafe.Add(mBase, uint32(v12+v29+v41)))
	v45 = *(*int32)(unsafe.Add(mBase, uint32(v41+(l1+v29))))
	v46 = v43 & v45
	v48 = base.B2i32(v46 != int32(0))
	if v46 != 0 {
		v54 = v48
		goto L3
	} else {
		goto L14
	}
L13:
	;
	v54 = v48
	goto L3
L14:
	;
	v50 = v34 + int32(1)
	if v50 != v28 {
		v34 = v50
		goto L12
	} else {
		goto L15
	}
L15:
	;
	goto L13
L16:
	;
	v58 = *(*int32)(unsafe.Add(mBase, uint32(l2)+220))
	if l3 != v58 {
		goto L19
	} else {
		goto L20
	}
L17:
	;
	v116 = l1
	goto L18
L18:
	;
	m.G0 = v10 + int32(16)
	return v116
L19:
	;
	if v58 == int32(0) {
		goto L1
	} else {
		goto L22
	}
L20:
	;
	v66 = l1
	goto L21
L21:
	;
	v67 = *(*int32)(unsafe.Add(mBase, uint32(l2)+8))
	v70 = F_find_appinfos_by_relids(m, l0, v67, v10+int32(12))
	mBase = m.M
	v71 = m.ExcPending
	if v71 != 0 {
		goto L23
	} else {
		goto L25
	}
L22:
	;
	v62 = F_adjust_child_relids_multilevel(m, l0, l1, v58, l3)
	mBase = m.M
	v65 = m.ExcPending
	if v65 != 0 {
		goto L23
	} else {
		goto L24
	}
L23:
	;
	return int32(0)
L24:
	;
	v66 = v62
	goto L21
L25:
	;
	v72 = int32(0)
	v73 = *(*int32)(unsafe.Add(mBase, uint32(v10)+12))
	if v73 <= v72 {
		goto L27
	} else {
		goto L28
	}
L26:
	;
	F_pfree(m, v70)
	mBase = m.M
	v113 = m.ExcPending
	if v113 != 0 {
		goto L23
	} else {
		goto L43
	}
L27:
	;
	v105 = int32(0)
	goto L26
L28:
	;
	goto L29
L29:
	;
	v78 = int32(0)
	v81 = v72
	goto L30
L30:
	;
	v88 = *(*int32)(unsafe.Add(mBase, uint32(v70+v81<<(uint(int32(2))%32))))
	v89 = *(*int32)(unsafe.Add(mBase, uint32(v88)+4))
	v90 = F_bms_is_member(m, v89, v66)
	mBase = m.M
	v91 = m.ExcPending
	if v91 != 0 {
		goto L23
	} else {
		goto L32
	}
L31:
	;
	v105 = v101
	goto L26
L32:
	;
	if v90 != 0 {
		goto L33
	} else {
		goto L34
	}
L33:
	;
	if v78 != 0 {
		goto L36
	} else {
		goto L37
	}
L34:
	;
	v101 = v78
	goto L35
L35:
	;
	v103 = v81 + int32(1)
	if v103 != v73 {
		v78 = v101
		v81 = v103
		goto L30
	} else {
		goto L42
	}
L36:
	;
	v94 = v78
	goto L38
L37:
	;
	v92 = F_bms_copy(m, v66)
	mBase = m.M
	v93 = m.ExcPending
	if v93 != 0 {
		goto L23
	} else {
		goto L39
	}
L38:
	;
	v95 = *(*int32)(unsafe.Add(mBase, uint32(v88)+4))
	v96 = F_bms_del_member(m, v94, v95)
	mBase = m.M
	v97 = m.ExcPending
	if v97 != 0 {
		goto L23
	} else {
		goto L40
	}
L39:
	;
	v94 = v92
	goto L38
L40:
	;
	v98 = *(*int32)(unsafe.Add(mBase, uint32(v88)+8))
	v99 = F_bms_add_member(m, v96, v98)
	mBase = m.M
	v100 = m.ExcPending
	if v100 != 0 {
		goto L23
	} else {
		goto L41
	}
L41:
	;
	v101 = v99
	goto L35
L42:
	;
	goto L31
L43:
	;
	if v105 != 0 {
		goto L44
	} else {
		goto L45
	}
L44:
	;
	v114 = v105
	goto L46
L45:
	;
	v114 = v66
	goto L46
L46:
	;
	v116 = v114
	goto L18
L47:
	;
	F_errmsg_internal(m, int32(_a_F_adjust_child_relids_multilevel_0), int32(0))
	mBase = m.M
	v133 = m.ExcPending
	if v133 != 0 {
		goto L23
	} else {
		goto L48
	}
L48:
	;
	F_errfinish(m, int32(_a_F_adjust_child_relids_multilevel_1), int32(634), int32(_a_F_adjust_child_relids_multilevel_2))
	mBase = m.M
	v138 = m.ExcPending
	if v138 != 0 {
		goto L23
	} else {
		goto L49
	}
L49:
	;
	base.Wasm_trap_unreachable()
	for {
	}
}
func F_adjust_standard_join_alias_expression(m *base.Module, l0 int32, l1 int32) {
	mBase := m.M
	_ = mBase
	var v4 int32
	_ = v4
	var v7 int32
	_ = v7
	var v10 int32
	_ = v10
	var v11 int32
	_ = v11
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
	var v32 int32
	_ = v32
	var v33 int32
	_ = v33
	var v36 int32
	_ = v36
	var v40 int32
	_ = v40
	var v43 int32
	_ = v43
	var v47 int32
	_ = v47
	var v49 int32
	_ = v49
	var v51 int32
	_ = v51
	var v52 int32
	_ = v52
	v4 = l0
	goto L2
L1:
	;
	return
L2:
	;
	v7 = *(*int32)(unsafe.Add(mBase, uint32(v4)))
	switch v7 - int32(6) {
	case 0:
		goto L11
	case 1, 2, 3, 4, 5, 6, 7, 8, 10, 11, 12, 13, 14, 15, 16, 17, 18, 19, 20, 24, 25, 26, 27, 28, 29, 30, 31:
		goto L1
	case 9:
		goto L8
	case 21:
		goto L7
	case 22:
		goto L6
	case 23:
		goto L5
	case 32:
		goto L4
	default:
		goto L10
	}
L3:
	;
	v33 = *(*int32)(unsafe.Add(mBase, uint32(v4)+12))
	if v33 == int32(0) {
		goto L1
	} else {
		goto L17
	}
L4:
	;
	goto L3
L5:
	;
	v32 = *(*int32)(unsafe.Add(mBase, uint32(v4)+4))
	v4 = v32
	goto L2
L6:
	;
	v31 = *(*int32)(unsafe.Add(mBase, uint32(v4)+4))
	v4 = v31
	goto L2
L7:
	;
	v30 = *(*int32)(unsafe.Add(mBase, uint32(v4)+4))
	v4 = v30
	goto L2
L8:
	;
	v27 = *(*int32)(unsafe.Add(mBase, uint32(v4)+28))
	v28 = *(*int32)(unsafe.Add(mBase, uint32(v27)+12))
	v29 = *(*int32)(unsafe.Add(mBase, uint32(v28)))
	v4 = v29
	goto L2
L9:
	;
	v21 = v20 + v4
	v22 = *(*int32)(unsafe.Add(mBase, uint32(v21)))
	v23 = *(*int32)(unsafe.Add(mBase, uint32(l1)+24))
	v24 = F_bms_add_members(m, v22, v23)
	mBase = m.M
	v25 = m.ExcPending
	if v25 != 0 {
		goto L15
	} else {
		goto L16
	}
L10:
	;
	if v7 != int32(319) {
		goto L1
	} else {
		goto L13
	}
L11:
	;
	v10 = *(*int32)(unsafe.Add(mBase, uint32(v4)+28))
	v11 = *(*int32)(unsafe.Add(mBase, uint32(l1)+28))
	if v10 != v11 {
		goto L1
	} else {
		goto L12
	}
L12:
	;
	v20 = int32(24)
	goto L9
L13:
	;
	v16 = *(*int32)(unsafe.Add(mBase, uint32(v4)+20))
	v17 = *(*int32)(unsafe.Add(mBase, uint32(l1)+28))
	if v16 != v17 {
		goto L1
	} else {
		goto L14
	}
L14:
	;
	v20 = int32(12)
	goto L9
L15:
	;
	return
L16:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v21))) = v24
	return
L17:
	;
	v36 = *(*int32)(unsafe.Add(mBase, uint32(v33)+4))
	if v36 <= int32(0) {
		goto L1
	} else {
		goto L18
	}
L18:
	;
	v40 = int32(0)
	goto L19
L19:
	;
	v43 = *(*int32)(unsafe.Add(mBase, uint32(v33)+12))
	v47 = *(*int32)(unsafe.Add(mBase, uint32(v43+v40<<(uint(int32(2))%32))))
	F_adjust_standard_join_alias_expression(m, v47, l1)
	mBase = m.M
	v49 = m.ExcPending
	if v49 != 0 {
		goto L15
	} else {
		goto L21
	}
L20:
	;
	goto L1
L21:
	;
	v51 = v40 + int32(1)
	v52 = *(*int32)(unsafe.Add(mBase, uint32(v33)+4))
	if v51 < v52 {
		v40 = v51
		goto L19
	} else {
		goto L22
	}
L22:
	;
	goto L20
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
	var v100 int32
	_ = v100
	var v107 int32
	_ = v107
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
						v100 = m.ExcPending
						if v100 != 0 {
							return
						} else {
							F_errfinish(m, int32(_a_F_amcheck_lock_relation_and_check_1), int32(128), int32(_a_F_amcheck_lock_relation_and_check_2))
							mBase = m.M
							v107 = m.ExcPending
							if v107 != 0 {
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
									v100 = m.ExcPending
									if v100 != 0 {
										return
									} else {
										F_errfinish(m, int32(_a_F_amcheck_lock_relation_and_check_1), int32(128), int32(_a_F_amcheck_lock_relation_and_check_2))
										mBase = m.M
										v107 = m.ExcPending
										if v107 != 0 {
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
												F_sequence_close(m, v26, l3)
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
											F_sequence_close(m, v26, l3)
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
	mBase := m.M
	_ = mBase
	var v2 int32
	_ = v2
	var v4 int32
	_ = v4
	var v11 int32
	_ = v11
	var v14 int32
	_ = v14
	var v19 int32
	_ = v19
	var v24 int32
	_ = v24
	v2 = m.G0
	v4 = v2 - int32(16)
	m.G0 = v4
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v11 = m.ExcPending
	if v11 != 0 {
		return int32(0)
	} else {
		F_errcode(m, int32(1088))
		mBase = m.M
		v14 = m.ExcPending
		if v14 != 0 {
			return int32(0)
		} else {
			*(*int32)(unsafe.Add(mBase, uint32(v4))) = int32(_a_F_anyarray_recv_0)
			F_errmsg(m, int32(_a_F_anyarray_recv_1), v4)
			mBase = m.M
			v19 = m.ExcPending
			if v19 != 0 {
				return int32(0)
			} else {
				F_errfinish(m, int32(_a_F_anyarray_recv_2), int32(155), int32(_a_F_anyarray_recv_3))
				mBase = m.M
				v24 = m.ExcPending
				if v24 != 0 {
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
	mBase := m.M
	_ = mBase
	var v2 int32
	_ = v2
	var v4 int32
	_ = v4
	var v11 int32
	_ = v11
	var v14 int32
	_ = v14
	var v19 int32
	_ = v19
	var v24 int32
	_ = v24
	v2 = m.G0
	v4 = v2 - int32(16)
	m.G0 = v4
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v11 = m.ExcPending
	if v11 != 0 {
		return int32(0)
	} else {
		F_errcode(m, int32(1088))
		mBase = m.M
		v14 = m.ExcPending
		if v14 != 0 {
			return int32(0)
		} else {
			*(*int32)(unsafe.Add(mBase, uint32(v4))) = int32(_a_F_anycompatiblemultirange_in_0)
			F_errmsg(m, int32(_a_F_anycompatiblemultirange_in_1), v4)
			mBase = m.M
			v19 = m.ExcPending
			if v19 != 0 {
				return int32(0)
			} else {
				F_errfinish(m, int32(_a_F_anycompatiblemultirange_in_2), int32(246), int32(_a_F_anycompatiblemultirange_in_3))
				mBase = m.M
				v24 = m.ExcPending
				if v24 != 0 {
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
	var v23 int32
	_ = v23
	var v25 int32
	_ = v25
	var v28 int32
	_ = v28
	var v30 int32
	_ = v30
	var v36 int32
	_ = v36
	var v41 int32
	_ = v41
	var v45 int32
	_ = v45
	var v51 int32
	_ = v51
	var v59 int32
	_ = v59
	var v60 int32
	_ = v60
	var v64 int32
	_ = v64
	var v69 int32
	_ = v69
	var v74 int32
	_ = v74
	var v80 int32
	_ = v80
	var v85 int32
	_ = v85
	var v95 int32
	_ = v95
	var v101 int32
	_ = v101
	var v103 int32
	_ = v103
	var v105 int32
	_ = v105
	var v106 int32
	_ = v106
	var v107 int32
	_ = v107
	var v111 int32
	_ = v111
	var v122 int32
	_ = v122
	var v128 int32
	_ = v128
	var v132 int32
	_ = v132
	var v140 int32
	_ = v140
	var v145 int32
	_ = v145
	var v147 int32
	_ = v147
	var v156 int32
	_ = v156
	var v158 int32
	_ = v158
	var v159 int32
	_ = v159
	var v161 int32
	_ = v161
	var v162 int32
	_ = v162
	var v163 int32
	_ = v163
	var v167 int32
	_ = v167
	var v178 int32
	_ = v178
	var v184 int32
	_ = v184
	v8 = m.G0
	v10 = v8 - int32(80)
	m.G0 = v10
	v12 = base.I32_wrap_i64(l1)
	v16 = base.I32_div_u_s(v12&int32(_a_F_append_num_word_0), int32(100))
	if base.Ui64(l1) <= base.Ui64(int64(20)) {
		v23 = *(*int32)(unsafe.Add(mBase, uint32(v12<<(uint(int32(2))%32))+uint32(_c_F_append_num_word[0])))
		F_appendStringInfoString(m, l0, v23)
		mBase = m.M
		v25 = m.ExcPending
		if v25 != 0 {
			return
		} else {
			m.G0 = v10 + int32(80)
			return
		}
	} else {
		v28 = v12 - v16*int32(100)
		v30 = v28 & int32(_a_F_append_num_word_0)
		if v30 == int32(0) {
			v36 = base.I32_div_u_s(v12&int32(_a_F_append_num_word_0), int32(100))
			v41 = *(*int32)(unsafe.Add(mBase, uint32(v36<<(uint(int32(2))%32))+uint32(_c_F_append_num_word[0])))
			*(*int32)(unsafe.Add(mBase, uint32(v10))) = v41
			F_appendStringInfo(m, l0, int32(_a_F_append_num_word_1), v10)
			mBase = m.M
			v45 = m.ExcPending
			if v45 != 0 {
				return
			} else {
				m.G0 = v10 + int32(80)
				return
			}
		} else {
			if base.Ui64(int64(100)) <= base.Ui64(l1) {
				v51 = base.I32_rem_u_s(v12&int32(_a_F_append_num_word_0), int32(10))
				if v51 != 0 {
					v85 = *(*int32)(unsafe.Add(mBase, uint32(v16<<(uint(int32(2))%32))+uint32(_c_F_append_num_word[0])))
					if base.Ui32(v28&int32(_a_F_append_num_word_0)) <= base.Ui32(int32(19)) {
						*(*int32)(unsafe.Add(mBase, uint32(v10)+32)) = v85
						v95 = *(*int32)(unsafe.Add(mBase, uint32(v30<<(uint(int32(2))%32))+uint32(_c_F_append_num_word[0])))
						*(*int32)(unsafe.Add(mBase, uint32(v10)+36)) = v95
						F_appendStringInfo(m, l0, int32(_a_F_append_num_word_2), v10+int32(32))
						mBase = m.M
						v101 = m.ExcPending
						if v101 != 0 {
							return
						} else {
							m.G0 = v10 + int32(80)
							return
						}
					} else {
						*(*int32)(unsafe.Add(mBase, uint32(v10)+48)) = v85
						v103 = int32(255)
						v105 = int32(10)
						v106 = base.I32_div_u_s(v28&v103, v105)
						v107 = int32(2)
						v111 = *(*int32)(unsafe.Add(mBase, uint32(v106<<(uint(v107)%32))+uint32(_c_F_append_num_word[1])))
						*(*int32)(unsafe.Add(mBase, uint32(v10)+52)) = v111
						v122 = *(*int32)(unsafe.Add(mBase, uint32((v28-v106*v105)&v103<<(uint(v107)%32))+uint32(_c_F_append_num_word[0])))
						*(*int32)(unsafe.Add(mBase, uint32(v10)+56)) = v122
						F_appendStringInfo(m, l0, int32(_a_F_append_num_word_3), v10+int32(48))
						mBase = m.M
						v128 = m.ExcPending
						if v128 != 0 {
							return
						} else {
							m.G0 = v10 + int32(80)
							return
						}
					}
				} else {
					if base.Ui32(v28&int32(_a_F_append_num_word_0)) < base.Ui32(int32(11)) {
						v85 = *(*int32)(unsafe.Add(mBase, uint32(v16<<(uint(int32(2))%32))+uint32(_c_F_append_num_word[0])))
						if base.Ui32(v28&int32(_a_F_append_num_word_0)) <= base.Ui32(int32(19)) {
							*(*int32)(unsafe.Add(mBase, uint32(v10)+32)) = v85
							v95 = *(*int32)(unsafe.Add(mBase, uint32(v30<<(uint(int32(2))%32))+uint32(_c_F_append_num_word[0])))
							*(*int32)(unsafe.Add(mBase, uint32(v10)+36)) = v95
							F_appendStringInfo(m, l0, int32(_a_F_append_num_word_2), v10+int32(32))
							mBase = m.M
							v101 = m.ExcPending
							if v101 != 0 {
								return
							} else {
								m.G0 = v10 + int32(80)
								return
							}
						} else {
							*(*int32)(unsafe.Add(mBase, uint32(v10)+48)) = v85
							v103 = int32(255)
							v105 = int32(10)
							v106 = base.I32_div_u_s(v28&v103, v105)
							v107 = int32(2)
							v111 = *(*int32)(unsafe.Add(mBase, uint32(v106<<(uint(v107)%32))+uint32(_c_F_append_num_word[1])))
							*(*int32)(unsafe.Add(mBase, uint32(v10)+52)) = v111
							v122 = *(*int32)(unsafe.Add(mBase, uint32((v28-v106*v105)&v103<<(uint(v107)%32))+uint32(_c_F_append_num_word[0])))
							*(*int32)(unsafe.Add(mBase, uint32(v10)+56)) = v122
							F_appendStringInfo(m, l0, int32(_a_F_append_num_word_3), v10+int32(48))
							mBase = m.M
							v128 = m.ExcPending
							if v128 != 0 {
								return
							} else {
								m.G0 = v10 + int32(80)
								return
							}
						}
					} else {
						v59 = base.I32_div_u_s(v28&int32(255), int32(10))
						v60 = int32(2)
						v64 = *(*int32)(unsafe.Add(mBase, uint32(v59<<(uint(v60)%32))+uint32(_c_F_append_num_word[1])))
						*(*int32)(unsafe.Add(mBase, uint32(v10)+20)) = v64
						v69 = base.I32_div_u_s(v12&int32(_a_F_append_num_word_0), int32(100))
						v74 = *(*int32)(unsafe.Add(mBase, uint32(v69<<(uint(v60)%32))+uint32(_c_F_append_num_word[0])))
						*(*int32)(unsafe.Add(mBase, uint32(v10)+16)) = v74
						F_appendStringInfo(m, l0, int32(_a_F_append_num_word_4), v10+int32(16))
						mBase = m.M
						v80 = m.ExcPending
						if v80 != 0 {
							return
						} else {
							m.G0 = v10 + int32(80)
							return
						}
					}
				}
			} else {
				v132 = base.I32_rem_u_s(v12&int32(255), int32(10))
				if v132 != 0 {
					if base.Ui32(v28&int32(_a_F_append_num_word_0)) <= base.Ui32(int32(19)) {
						v156 = *(*int32)(unsafe.Add(mBase, uint32(v30<<(uint(int32(2))%32))+uint32(_c_F_append_num_word[0])))
						F_appendStringInfoString(m, l0, v156)
						mBase = m.M
						v158 = m.ExcPending
						if v158 != 0 {
							return
						} else {
							m.G0 = v10 + int32(80)
							return
						}
					} else {
						v159 = int32(255)
						v161 = int32(10)
						v162 = base.I32_div_u_s(v28&v159, v161)
						v163 = int32(2)
						v167 = *(*int32)(unsafe.Add(mBase, uint32(v162<<(uint(v163)%32))+uint32(_c_F_append_num_word[1])))
						*(*int32)(unsafe.Add(mBase, uint32(v10)+64)) = v167
						v178 = *(*int32)(unsafe.Add(mBase, uint32((v28-v162*v161)&v159<<(uint(v163)%32))+uint32(_c_F_append_num_word[0])))
						*(*int32)(unsafe.Add(mBase, uint32(v10)+68)) = v178
						F_appendStringInfo(m, l0, int32(_a_F_append_num_word_5), v10-int32(-64))
						mBase = m.M
						v184 = m.ExcPending
						if v184 != 0 {
							return
						} else {
							m.G0 = v10 + int32(80)
							return
						}
					}
				} else {
					if base.Ui32(v28&int32(_a_F_append_num_word_0)) < base.Ui32(int32(11)) {
						if base.Ui32(v28&int32(_a_F_append_num_word_0)) <= base.Ui32(int32(19)) {
							v156 = *(*int32)(unsafe.Add(mBase, uint32(v30<<(uint(int32(2))%32))+uint32(_c_F_append_num_word[0])))
							F_appendStringInfoString(m, l0, v156)
							mBase = m.M
							v158 = m.ExcPending
							if v158 != 0 {
								return
							} else {
								m.G0 = v10 + int32(80)
								return
							}
						} else {
							v159 = int32(255)
							v161 = int32(10)
							v162 = base.I32_div_u_s(v28&v159, v161)
							v163 = int32(2)
							v167 = *(*int32)(unsafe.Add(mBase, uint32(v162<<(uint(v163)%32))+uint32(_c_F_append_num_word[1])))
							*(*int32)(unsafe.Add(mBase, uint32(v10)+64)) = v167
							v178 = *(*int32)(unsafe.Add(mBase, uint32((v28-v162*v161)&v159<<(uint(v163)%32))+uint32(_c_F_append_num_word[0])))
							*(*int32)(unsafe.Add(mBase, uint32(v10)+68)) = v178
							F_appendStringInfo(m, l0, int32(_a_F_append_num_word_5), v10-int32(-64))
							mBase = m.M
							v184 = m.ExcPending
							if v184 != 0 {
								return
							} else {
								m.G0 = v10 + int32(80)
								return
							}
						}
					} else {
						v140 = base.I32_div_u_s(v28&int32(255), int32(10))
						v145 = *(*int32)(unsafe.Add(mBase, uint32(v140<<(uint(int32(2))%32))+uint32(_c_F_append_num_word[1])))
						F_appendStringInfoString(m, l0, v145)
						mBase = m.M
						v147 = m.ExcPending
						if v147 != 0 {
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
	var v11 int32
	_ = v11
	var v13 int32
	_ = v13
	var v15 int32
	_ = v15
	var v18 int32
	_ = v18
	var v23 int32
	_ = v23
	var v30 int32
	_ = v30
	var v33 int32
	_ = v33
	var v42 int32
	_ = v42
	var v44 int32
	_ = v44
	var v48 int32
	_ = v48
	var v49 int32
	_ = v49
	var v59 int32
	_ = v59
	var v62 int32
	_ = v62
	var v63 int32
	_ = v63
	var v66 int32
	_ = v66
	var v67 int32
	_ = v67
	var v68 int32
	_ = v68
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
	var v183 int32
	_ = v183
	var v195 int32
	_ = v195
	var v196 int32
	_ = v196
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
	var v234 int32
	_ = v234
	var v236 int32
	_ = v236
	var v241 int32
	_ = v241
	v11 = m.G0
	v13 = v11 - int32(16)
	m.G0 = v13
	v15 = int32(1)
	if l1 < int32(4) {
		v241 = v15
		goto L1
	} else {
		goto L2
	}
L1:
	;
	m.G0 = v13 + int32(16)
	return v241
L2:
	;
	v18 = int32(21)
	v23 = (l1<<(uint(v18)%32) - int32(_a_F_apply_typmod_0)) >> (uint(v18) % 32)
	*(*int32)(unsafe.Add(mBase, uint32(l0)+12)) = v23
	v30 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v33 = v23 + v30<<(uint(int32(2))%32)
	if v33+int32(4) < int32(0) {
		goto L4
	} else {
		goto L5
	}
L3:
	;
	v152 = int32(base.Ui32(l1-int32(4)) >> (uint(int32(16)) % 32))
	v153 = v152 - v23
	v154 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	if v154 < int32(0) {
		goto L30
	} else {
		goto L31
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
	v42 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v44 = v23 & int32(3)
	v48 = base.I32_div_s(v33+int32(7), int32(4))
	v49 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	if v49 <= v48 {
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
		goto L29
	}
L9:
	;
	v95 = v89
	goto L23
L10:
	;
	v62 = int32(1)
	v63 = v48 - v62
	v66 = v42 + v63<<(uint(v62)%32)
	v67 = int32(*(*int16)(unsafe.Add(mBase, uint32(v66))))
	v68 = int32(2)
	v72 = *(*int32)(unsafe.Add(mBase, uint32(v44<<(uint(v68)%32))+uint32(_c_F_apply_typmod[0])))
	v73 = base.I32_rem_s(v67, v72)
	v74 = v67 - v73
	*(*uint16)(unsafe.Add(mBase, uint32(v66))) = uint16(v74)
	v77 = base.I32_div_s(v72, v68)
	if v73 < v77 {
		v115 = v63
		goto L8
	} else {
		goto L18
	}
L11:
	;
	if v44 == int32(0) {
		goto L7
	} else {
		goto L14
	}
L12:
	;
	goto L13
L13:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0))) = v48
	if v44 != 0 {
		goto L10
	} else {
		goto L16
	}
L14:
	;
	if v48 != v49 {
		goto L7
	} else {
		goto L15
	}
L15:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0))) = v48
	goto L10
L16:
	;
	v59 = int32(*(*int16)(unsafe.Add(mBase, uint32(v42+v48<<(uint(int32(1))%32)))))
	if v59 <= int32(_a_F_apply_typmod_1) {
		v115 = v48
		goto L8
	} else {
		goto L17
	}
L17:
	;
	v89 = v48
	goto L9
L18:
	;
	v80 = v72 + base.I32_extend16_s(v74)
	if int32(_a_F_apply_typmod_2) < v80 {
		goto L19
	} else {
		goto L20
	}
L19:
	;
	v85 = v80 + int32(_a_F_apply_typmod_3)
	goto L21
L20:
	;
	v85 = v80
	goto L21
L21:
	;
	*(*uint16)(unsafe.Add(mBase, uint32(v66))) = uint16(v85)
	if v80 < int32(_a_F_apply_typmod_4) {
		v115 = v63
		goto L8
	} else {
		goto L22
	}
L22:
	;
	v89 = v63
	goto L9
L23:
	;
	v101 = int32(1)
	v102 = v95 - v101
	v105 = v42 + v102<<(uint(v101)%32)
	v108 = int32(*(*int16)(unsafe.Add(mBase, uint32(v105))))
	v110 = base.B2i32(int32(_a_F_apply_typmod_5) < v108)
	if int32(_a_F_apply_typmod_5) < v108 {
		goto L25
	} else {
		goto L26
	}
L24:
	;
	v115 = v102
	goto L8
L25:
	;
	v111 = int32(-9999)
	goto L27
L26:
	;
	v111 = v101
	goto L27
L27:
	;
	v112 = v111 + v108
	*(*uint16)(unsafe.Add(mBase, uint32(v105))) = uint16(v112)
	if int32(_a_F_apply_typmod_5) < v108 {
		v95 = v102
		goto L23
	} else {
		goto L28
	}
L28:
	;
	goto L24
L29:
	;
	v123 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+20)) = v123 - int32(2)
	v127 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v128 = int32(1)
	*(*int32)(unsafe.Add(mBase, uint32(l0))) = v127 + v128
	v131 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v131 + v128
	goto L7
L30:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+12)) = int32(0)
	goto L32
L31:
	;
	goto L32
L32:
	;
	v159 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v163 = v159<<(uint(int32(2))%32) + int32(4)
	if v163 <= v153 {
		v241 = v15
		goto L1
	} else {
		goto L33
	}
L33:
	;
	v165 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	if v165 <= int32(0) {
		v241 = v15
		goto L1
	} else {
		goto L34
	}
L34:
	;
	v168 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v170 = int32(0)
	v171 = v163
	goto L35
L35:
	;
	v183 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v168+v170<<(uint(int32(1))%32)))))
	if v183 != 0 {
		goto L37
	} else {
		goto L38
	}
L36:
	;
	v241 = v234
	goto L1
L37:
	;
	if base.I32_extend16_s(v183) < int32(10) {
		v196 = int32(-3)
		goto L40
	} else {
		goto L41
	}
L38:
	;
	goto L39
L39:
	;
	v234 = int32(1)
	v236 = v170 + v234
	if v236 != v165 {
		v170 = v236
		v171 = v171 - int32(4)
		goto L35
	} else {
		goto L60
	}
L40:
	;
	if v196+v171 <= v153 {
		v241 = int32(1)
		goto L1
	} else {
		goto L46
	}
L41:
	;
	if base.Ui32(v183) < base.Ui32(int32(100)) {
		v196 = int32(-2)
		goto L40
	} else {
		goto L42
	}
L42:
	;
	if base.Ui32(v183) < base.Ui32(int32(1000)) {
		goto L43
	} else {
		goto L44
	}
L43:
	;
	v195 = int32(-1)
	goto L45
L44:
	;
	v195 = int32(0)
	goto L45
L45:
	;
	v196 = v195
	goto L40
L46:
	;
	v200 = int32(0)
	v201 = F_errsave_start(m, l2)
	mBase = m.M
	v204 = m.ExcPending
	if v204 != 0 {
		goto L47
	} else {
		goto L48
	}
L47:
	;
	return int32(0)
L48:
	;
	if v201 == int32(0) {
		v241 = v200
		goto L1
	} else {
		goto L49
	}
L49:
	;
	F_errcode(m, int32(50331778))
	mBase = m.M
	v209 = m.ExcPending
	if v209 != 0 {
		goto L47
	} else {
		goto L50
	}
L50:
	;
	F_errmsg(m, int32(_a_F_apply_typmod_6), int32(0))
	mBase = m.M
	v213 = m.ExcPending
	if v213 != 0 {
		goto L47
	} else {
		goto L51
	}
L51:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v13))) = v152
	*(*int32)(unsafe.Add(mBase, uint32(v13)+4)) = v23
	v217 = base.B2i32(v23 == v152)
	if v23 == v152 {
		goto L52
	} else {
		goto L53
	}
L52:
	;
	v218 = int32(1)
	goto L54
L53:
	;
	v218 = v153
	goto L54
L54:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v13)+12)) = v218
	if v23 == v152 {
		goto L55
	} else {
		goto L56
	}
L55:
	;
	v222 = int32(_a_F_apply_typmod_7)
	goto L57
L56:
	;
	v222 = int32(_a_F_apply_typmod_8)
	goto L57
L57:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v13)+8)) = v222
	F_errdetail(m, int32(_a_F_apply_typmod_9), v13)
	mBase = m.M
	v226 = m.ExcPending
	if v226 != 0 {
		goto L47
	} else {
		goto L58
	}
L58:
	;
	F_errsave_finish(m, l2, int32(_a_F_apply_typmod_10), int32(_a_F_apply_typmod_11), int32(_a_F_apply_typmod_12))
	mBase = m.M
	v231 = m.ExcPending
	if v231 != 0 {
		goto L47
	} else {
		goto L59
	}
L59:
	;
	v241 = v200
	goto L1
L60:
	;
	goto L36
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
			v18 = F_array_contain_compare(m, v6, v11, v13, int32(0), v15+int32(16))
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
func F_assign_application_name(m *base.Module, l0 int32, l1 int32) {
	mBase := m.M
	_ = mBase
	var v5 int32
	_ = v5
	var v6 int32
	_ = v6
	var v8 int32
	_ = v8
	var v9 int32
	_ = v9
	var v10 int32
	_ = v10
	var v12 int32
	_ = v12
	var v13 int32
	_ = v13
	var v16 int32
	_ = v16
	var v20 int32
	_ = v20
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
	v5 = *(*int32)(unsafe.Add(mBase, _c_F_assign_application_name[0]))
	if v5 != 0 {
		v6 = F_strlen(m, l0)
		mBase = m.M
		v8 = F_pg_mbcliplen(m, l0, v6, int32(63))
		mBase = m.M
		v9 = m.ExcPending
		if v9 != 0 {
			return
		} else {
			v10 = int32(_a_F_assign_application_name_0)
			v12 = *(*int32)(unsafe.Add(mBase, _c_F_assign_application_name[1]))
			v13 = int32(1)
			*(*int32)(unsafe.Add(mBase, _c_F_assign_application_name[1])) = v12 + v13
			v16 = *(*int32)(unsafe.Add(mBase, uint32(v5)))
			*(*int32)(unsafe.Add(mBase, uint32(v5))) = v16 + v13
			v20 = *(*int32)(unsafe.Add(mBase, uint32(v5)+212))
			if v8 != 0 {
				v21 = F__emscripten_memcpy_bulkmem(m, v20, l0, v8)
				mBase = m.M
			} else {
			}
			v23 = *(*int32)(unsafe.Add(mBase, uint32(v5)+212))
			v25 = int32(0)
			*(*uint8)(unsafe.Add(mBase, uint32(v8+v23))) = uint8(v25)
			v27 = *(*int32)(unsafe.Add(mBase, uint32(v5)))
			v28 = int32(1)
			*(*int32)(unsafe.Add(mBase, uint32(v5))) = v27 + v28
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
