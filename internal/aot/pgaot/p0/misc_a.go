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
	var v16 int32
	_ = v16
	var v17 int32
	_ = v17
	var v18 int32
	_ = v18
	var v21 int32
	_ = v21
	var v22 int32
	_ = v22
	var v26 int32
	_ = v26
	var v32 int32
	_ = v32
	v5 = *(*int32)(unsafe.Add(mBase, _c_F_AbortStrongLockAcquire[0]))
	if v5 != 0 {
		v6 = *(*int32)(unsafe.Add(mBase, uint32(v5)+20))
		v10 = *(*int32)(unsafe.Add(mBase, _c_F_AbortStrongLockAcquire[1]))
		v13 = base.AtomicRmwXchg32(m, v10, int32(0), int32(1))
		if v13 != 0 {
			F_s_lock(m, v10, int32(_a_F_AbortStrongLockAcquire_0))
			mBase = m.M
			v16 = m.ExcPending
			if v16 != 0 {
				return
			} else {
				v17 = int32(_a_F_AbortStrongLockAcquire_1)
				v18 = *(*int32)(unsafe.Add(mBase, _c_F_AbortStrongLockAcquire[1]))
				v21 = v18 + v6&int32(1023)<<(uint(int32(2))%32)
				v22 = *(*int32)(unsafe.Add(mBase, uint32(v21)+4))
				*(*int32)(unsafe.Add(mBase, uint32(v21)+4)) = v22 - int32(1)
				v26 = int32(0)
				*(*uint8)(unsafe.Add(mBase, uint32(v5)+52)) = uint8(v26)
				*(*int32)(unsafe.Add(mBase, _c_F_AbortStrongLockAcquire[0])) = v26
				v32 = *(*int32)(unsafe.Add(mBase, _c_F_AbortStrongLockAcquire[1]))
				atomic.StoreUint32((*uint32)(unsafe.Add(mBase, uint32(v32))), uint32(v26))
				return
			}
		} else {
			v17 = int32(_a_F_AbortStrongLockAcquire_1)
			v18 = *(*int32)(unsafe.Add(mBase, _c_F_AbortStrongLockAcquire[1]))
			v21 = v18 + v6&int32(1023)<<(uint(int32(2))%32)
			v22 = *(*int32)(unsafe.Add(mBase, uint32(v21)+4))
			*(*int32)(unsafe.Add(mBase, uint32(v21)+4)) = v22 - int32(1)
			v26 = int32(0)
			*(*uint8)(unsafe.Add(mBase, uint32(v5)+52)) = uint8(v26)
			*(*int32)(unsafe.Add(mBase, _c_F_AbortStrongLockAcquire[0])) = v26
			v32 = *(*int32)(unsafe.Add(mBase, _c_F_AbortStrongLockAcquire[1]))
			atomic.StoreUint32((*uint32)(unsafe.Add(mBase, uint32(v32))), uint32(v26))
			return
		}
	} else {
		return
	}
}
func F_AbsorbSyncRequests(m *base.Module) {
	mBase := m.M
	_ = mBase
	var v11 int32
	_ = v11
	var v17 int32
	_ = v17
	var v24 int32
	_ = v24
	var v28 int32
	_ = v28
	var v29 int32
	_ = v29
	var v30 int32
	_ = v30
	var v32 int32
	_ = v32
	var v33 int32
	_ = v33
	var v36 int32
	_ = v36
	var v37 int32
	_ = v37
	var v38 int32
	_ = v38
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
	var v54 int32
	_ = v54
	var v57 int32
	_ = v57
	var v59 int32
	_ = v59
	var v62 int32
	_ = v62
	var v67 int32
	_ = v67
	var v69 int32
	_ = v69
	var v72 int32
	_ = v72
	var v73 int64
	_ = v73
	var v75 int64
	_ = v75
	var v77 int64
	_ = v77
	var v79 int64
	_ = v79
	var v81 int32
	_ = v81
	var v82 int32
	_ = v82
	var v84 int32
	_ = v84
	var v85 int32
	_ = v85
	var v88 int32
	_ = v88
	var v90 int32
	_ = v90
	var v93 int32
	_ = v93
	var v96 int32
	_ = v96
	var v102 int32
	_ = v102
	var v104 int32
	_ = v104
	var v108 int32
	_ = v108
	var v110 int32
	_ = v110
	var v114 int32
	_ = v114
	var v117 int32
	_ = v117
	var v119 int32
	_ = v119
	var v128 int32
	_ = v128
	var v130 int32
	_ = v130
	var v133 int32
	_ = v133
	var v146 int32
	_ = v146
	var v148 int32
	_ = v148
	var v155 int32
	_ = v155
	v11 = *(*int32)(unsafe.Add(mBase, _c_F_AbsorbSyncRequests[0]))
	if v11 != int32(11) {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	return
L2:
	;
	v17 = int32(0)
	goto L3
L3:
	;
	v24 = *(*int32)(unsafe.Add(mBase, _c_F_AbsorbSyncRequests[1]))
	v28 = F_LWLockAcquire(m, v24+int32(2176), int32(0))
	mBase = m.M
	v29 = m.ExcPending
	if v29 != 0 {
		goto L5
	} else {
		goto L6
	}
L4:
	;
	if v96 == int32(0) {
		goto L1
	} else {
		goto L32
	}
L5:
	;
	return
L6:
	;
	v30 = int32(_a_F_AbsorbSyncRequests_0)
	v32 = *(*int32)(unsafe.Add(mBase, _c_F_AbsorbSyncRequests[2]))
	v33 = *(*int32)(unsafe.Add(mBase, uint32(v32)+48))
	if v30 <= v33 {
		goto L7
	} else {
		goto L8
	}
L7:
	;
	v36 = v30
	goto L9
L8:
	;
	v36 = v33
	goto L9
L9:
	;
	v37 = int32(0)
	v38 = base.B2i32(v33 <= v37)
	if v38 == v37 {
		goto L10
	} else {
		goto L11
	}
L10:
	;
	if v17 == int32(0) {
		goto L13
	} else {
		goto L14
	}
L11:
	;
	v93 = v32
	v96 = v17
	goto L12
L12:
	;
	v102 = int32(_a_F_AbsorbSyncRequests_1)
	v104 = *(*int32)(unsafe.Add(mBase, _c_F_AbsorbSyncRequests[3]))
	*(*int32)(unsafe.Add(mBase, _c_F_AbsorbSyncRequests[3])) = v104 + int32(1)
	v108 = *(*int32)(unsafe.Add(mBase, uint32(v93)+48))
	v110 = *(*int32)(unsafe.Add(mBase, _c_F_AbsorbSyncRequests[1]))
	F_LWLockRelease(m, v110+int32(2176))
	mBase = m.M
	v114 = m.ExcPending
	if v114 != 0 {
		goto L5
	} else {
		goto L23
	}
L13:
	;
	v45 = F_palloc(m, v36<<(uint(int32(5))%32))
	mBase = m.M
	v46 = m.ExcPending
	if v46 != 0 {
		goto L5
	} else {
		goto L16
	}
L14:
	;
	v47 = v17
	goto L15
L15:
	;
	v48 = int32(1)
	if v36 <= v48 {
		goto L17
	} else {
		goto L18
	}
L16:
	;
	v47 = v45
	goto L15
L17:
	;
	v51 = v48
	goto L19
L18:
	;
	v51 = v36
	goto L19
L19:
	;
	v54 = *(*int32)(unsafe.Add(mBase, _c_F_AbsorbSyncRequests[2]))
	v57 = *(*int32)(unsafe.Add(mBase, uint32(v54)+56))
	v59 = int32(0)
	v62 = v57
	goto L20
L20:
	;
	v67 = int32(5)
	v69 = v47 + v59<<(uint(v67)%32)
	v72 = v54 - int32(-64) + v62<<(uint(v67)%32)
	v73 = *(*int64)(unsafe.Add(mBase, uint32(v72)+24))
	*(*int64)(unsafe.Add(mBase, uint32(v69)+24)) = v73
	v75 = *(*int64)(unsafe.Add(mBase, uint32(v72)+16))
	*(*int64)(unsafe.Add(mBase, uint32(v69)+16)) = v75
	v77 = *(*int64)(unsafe.Add(mBase, uint32(v72)+8))
	*(*int64)(unsafe.Add(mBase, uint32(v69)+8)) = v77
	v79 = *(*int64)(unsafe.Add(mBase, uint32(v72)))
	*(*int64)(unsafe.Add(mBase, uint32(v69))) = v79
	v81 = *(*int32)(unsafe.Add(mBase, uint32(v54)+56))
	v82 = int32(1)
	v84 = *(*int32)(unsafe.Add(mBase, uint32(v54)+52))
	v85 = base.I32_rem_s(v81+v82, v84)
	*(*int32)(unsafe.Add(mBase, uint32(v54)+56)) = v85
	v88 = v59 + v82
	if v88 != v51 {
		v59 = v88
		v62 = v85
		goto L20
	} else {
		goto L22
	}
L21:
	;
	v90 = *(*int32)(unsafe.Add(mBase, uint32(v54)+48))
	*(*int32)(unsafe.Add(mBase, uint32(v54)+48)) = v90 - v36
	v93 = v54
	v96 = v47
	goto L12
L22:
	;
	goto L21
L23:
	;
	if v38 == int32(0) {
		goto L24
	} else {
		goto L25
	}
L24:
	;
	v117 = v96
	v119 = v36
	goto L27
L25:
	;
	goto L26
L26:
	;
	v146 = int32(_a_F_AbsorbSyncRequests_1)
	v148 = *(*int32)(unsafe.Add(mBase, _c_F_AbsorbSyncRequests[3]))
	*(*int32)(unsafe.Add(mBase, _c_F_AbsorbSyncRequests[3])) = v148 - int32(1)
	if v108 != 0 {
		v17 = v96
		goto L3
	} else {
		goto L31
	}
L27:
	;
	v128 = *(*int32)(unsafe.Add(mBase, uint32(v117)))
	F_RememberSyncRequest(m, v117+int32(8), v128)
	mBase = m.M
	v130 = m.ExcPending
	if v130 != 0 {
		goto L5
	} else {
		goto L29
	}
L28:
	;
	goto L26
L29:
	;
	v133 = int32(1)
	if v133 < v119 {
		v117 = v117 + int32(32)
		v119 = v119 - v133
		goto L27
	} else {
		goto L30
	}
L30:
	;
	goto L28
L31:
	;
	goto L4
L32:
	;
	F_pfree(m, v96)
	mBase = m.M
	v155 = m.ExcPending
	if v155 != 0 {
		goto L5
	} else {
		goto L33
	}
L33:
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
	F_errfinish(m, int32(_a_F_AllocateDir_1), int32(2931), int32(_a_F_AllocateDir_2))
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
	F_errfinish(m, int32(_a_F_AllocateDir_1), int32(2908), int32(_a_F_AllocateDir_2))
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
	var v8 int32
	_ = v8
	var v10 int32
	_ = v10
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
	var v18 int32
	_ = v18
	var v21 int32
	_ = v21
	var v29 int32
	_ = v29
	var v32 int64
	_ = v32
	var v33 int32
	_ = v33
	var v34 int32
	_ = v34
	var v38 int32
	_ = v38
	var v39 int32
	_ = v39
	var v40 int32
	_ = v40
	var v41 int32
	_ = v41
	var v42 int32
	_ = v42
	var v43 int32
	_ = v43
	var v48 int32
	_ = v48
	var v55 int32
	_ = v55
	var v56 int32
	_ = v56
	var v60 int32
	_ = v60
	var v62 int32
	_ = v62
	var v64 int32
	_ = v64
	var v69 int32
	_ = v69
	var v71 int32
	_ = v71
	var v72 int32
	_ = v72
	var v76 int32
	_ = v76
	var v83 int32
	_ = v83
	var v86 int32
	_ = v86
	var v94 int32
	_ = v94
	var v98 int32
	_ = v98
	var v103 int32
	_ = v103
	var v107 int32
	_ = v107
	var v110 int32
	_ = v110
	var v116 int32
	_ = v116
	var v120 int32
	_ = v120
	var v125 int32
	_ = v125
	v8 = m.G0
	v10 = v8 - int32(112)
	m.G0 = v10
	v12 = *(*int32)(unsafe.Add(mBase, uint32(l1)+16))
	v13 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v12)+22)))
	v14 = v12 + v13
	v15 = F_superuser(m)
	mBase = m.M
	v16 = m.ExcPending
	if v16 != 0 {
		return
	} else {
		if v15 != 0 {
			v17 = F_superuser_arg(m, l2)
			mBase = m.M
			v18 = m.ExcPending
			if v18 != 0 {
				return
			} else {
				if v17 == int32(0) {
					F_errstart_cold(m, int32(21), int32(0))
					mBase = m.M
					v107 = m.ExcPending
					if v107 != 0 {
						return
					} else {
						F_errcode(m, int32(16797828))
						mBase = m.M
						v110 = m.ExcPending
						if v110 != 0 {
							return
						} else {
							*(*int32)(unsafe.Add(mBase, uint32(v10))) = v14 + int32(4)
							F_errmsg(m, int32(_a_F_AlterForeignDataWrapperOwner_internal_0), v10)
							mBase = m.M
							v116 = m.ExcPending
							if v116 != 0 {
								return
							} else {
								F_errhint(m, int32(_a_F_AlterForeignDataWrapperOwner_internal_1), int32(0))
								mBase = m.M
								v120 = m.ExcPending
								if v120 != 0 {
									return
								} else {
									F_errfinish(m, int32(_a_F_AlterForeignDataWrapperOwner_internal_2), int32(242), int32(_a_F_AlterForeignDataWrapperOwner_internal_3))
									mBase = m.M
									v125 = m.ExcPending
									if v125 != 0 {
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
					v21 = *(*int32)(unsafe.Add(mBase, uint32(v14)+68))
					if l2 != v21 {
						*(*int64)(unsafe.Add(mBase, uint32(v10)+32)) = int64(65536)
						*(*int64)(unsafe.Add(mBase, uint32(v10)+40)) = int64(0)
						*(*int64)(unsafe.Add(mBase, uint32(v10)+64)) = base.I64_extend_i32_u(l2)
						v29 = *(*int32)(unsafe.Add(mBase, uint32(l0)+52))
						v32 = F_heap_getattr_8(m, l1, v29, v10+int32(31))
						mBase = m.M
						v33 = m.ExcPending
						if v33 != 0 {
							return
						} else {
							v34 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v10)+31)))
							if v34 == int32(0) {
								v38 = F_pg_detoast_datum(m, base.I32_wrap_i64(v32))
								mBase = m.M
								v39 = m.ExcPending
								if v39 != 0 {
									return
								} else {
									v40 = *(*int32)(unsafe.Add(mBase, uint32(v14)+68))
									v41 = F_aclnewowner(m, v38, v40, l2)
									mBase = m.M
									v42 = m.ExcPending
									if v42 != 0 {
										return
									} else {
										v43 = int32(1)
										*(*uint8)(unsafe.Add(mBase, uint32(v10)+38)) = uint8(v43)
										*(*int64)(unsafe.Add(mBase, uint32(v10)+96)) = base.I64_extend_i32_u(v41)
										v48 = *(*int32)(unsafe.Add(mBase, uint32(l0)+52))
										v55 = F_heap_modify_tuple(m, l1, v48, v10+int32(48), v10+int32(40), v10+int32(32))
										mBase = m.M
										v56 = m.ExcPending
										if v56 != 0 {
											return
										} else {
											F_CatalogTupleUpdate(m, l0, v55+int32(4), v55)
											mBase = m.M
											v60 = m.ExcPending
											if v60 != 0 {
												return
											} else {
												v62 = *(*int32)(unsafe.Add(mBase, uint32(v14)))
												F_changeDependencyOnOwner(m, int32(2328), v62, l2)
												mBase = m.M
												v64 = m.ExcPending
												if v64 != 0 {
													return
												} else {
													v69 = *(*int32)(unsafe.Add(mBase, _c_F_AlterForeignDataWrapperOwner_internal[0]))
													if v69 != 0 {
														v71 = *(*int32)(unsafe.Add(mBase, uint32(v14)))
														v72 = int32(0)
														F_RunObjectPostAlterHook(m, int32(2328), v71, v72, v72, v72)
														mBase = m.M
														v76 = m.ExcPending
														if v76 != 0 {
															return
														} else {
															m.G0 = v10 + int32(112)
															return
														}
													} else {
														m.G0 = v10 + int32(112)
														return
													}
												}
											}
										}
									}
								}
							} else {
								v48 = *(*int32)(unsafe.Add(mBase, uint32(l0)+52))
								v55 = F_heap_modify_tuple(m, l1, v48, v10+int32(48), v10+int32(40), v10+int32(32))
								mBase = m.M
								v56 = m.ExcPending
								if v56 != 0 {
									return
								} else {
									F_CatalogTupleUpdate(m, l0, v55+int32(4), v55)
									mBase = m.M
									v60 = m.ExcPending
									if v60 != 0 {
										return
									} else {
										v62 = *(*int32)(unsafe.Add(mBase, uint32(v14)))
										F_changeDependencyOnOwner(m, int32(2328), v62, l2)
										mBase = m.M
										v64 = m.ExcPending
										if v64 != 0 {
											return
										} else {
											v69 = *(*int32)(unsafe.Add(mBase, _c_F_AlterForeignDataWrapperOwner_internal[0]))
											if v69 != 0 {
												v71 = *(*int32)(unsafe.Add(mBase, uint32(v14)))
												v72 = int32(0)
												F_RunObjectPostAlterHook(m, int32(2328), v71, v72, v72, v72)
												mBase = m.M
												v76 = m.ExcPending
												if v76 != 0 {
													return
												} else {
													m.G0 = v10 + int32(112)
													return
												}
											} else {
												m.G0 = v10 + int32(112)
												return
											}
										}
									}
								}
							}
						}
					} else {
						v69 = *(*int32)(unsafe.Add(mBase, _c_F_AlterForeignDataWrapperOwner_internal[0]))
						if v69 != 0 {
							v71 = *(*int32)(unsafe.Add(mBase, uint32(v14)))
							v72 = int32(0)
							F_RunObjectPostAlterHook(m, int32(2328), v71, v72, v72, v72)
							mBase = m.M
							v76 = m.ExcPending
							if v76 != 0 {
								return
							} else {
								m.G0 = v10 + int32(112)
								return
							}
						} else {
							m.G0 = v10 + int32(112)
							return
						}
					}
				}
			}
		} else {
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
					*(*int32)(unsafe.Add(mBase, uint32(v10)+16)) = v14 + int32(4)
					F_errmsg(m, int32(_a_F_AlterForeignDataWrapperOwner_internal_0), v10+int32(16))
					mBase = m.M
					v94 = m.ExcPending
					if v94 != 0 {
						return
					} else {
						F_errhint(m, int32(_a_F_AlterForeignDataWrapperOwner_internal_4), int32(0))
						mBase = m.M
						v98 = m.ExcPending
						if v98 != 0 {
							return
						} else {
							F_errfinish(m, int32(_a_F_AlterForeignDataWrapperOwner_internal_2), int32(234), int32(_a_F_AlterForeignDataWrapperOwner_internal_3))
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
		}
	}
}
func F_AlterObjectNamespace_internal(m *base.Module, l0 int32, l1 int32, l2 int32) int32 {
	mBase := m.M
	_ = mBase
	var v14 int32
	_ = v14
	var v16 int32
	_ = v16
	var v18 int32
	_ = v18
	var v19 int32
	_ = v19
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
	var v29 int32
	_ = v29
	var v30 int32
	_ = v30
	var v33 int32
	_ = v33
	var v34 int32
	_ = v34
	var v35 int32
	_ = v35
	var v37 int32
	_ = v37
	var v38 int64
	_ = v38
	var v39 int32
	_ = v39
	var v40 int32
	_ = v40
	var v41 int64
	_ = v41
	var v42 int32
	_ = v42
	var v43 int32
	_ = v43
	var v46 int32
	_ = v46
	var v48 int32
	_ = v48
	var v49 int32
	_ = v49
	var v50 int32
	_ = v50
	var v53 int32
	_ = v53
	var v56 int64
	_ = v56
	var v57 int32
	_ = v57
	var v59 int32
	_ = v59
	var v61 int32
	_ = v61
	var v62 int32
	_ = v62
	var v66 int32
	_ = v66
	var v67 int32
	_ = v67
	var v70 int32
	_ = v70
	var v73 int32
	_ = v73
	var v75 int32
	_ = v75
	var v76 int32
	_ = v76
	var v80 int32
	_ = v80
	var v81 int32
	_ = v81
	var v83 int32
	_ = v83
	var v92 int32
	_ = v92
	var v93 int32
	_ = v93
	var v94 int32
	_ = v94
	var v97 int32
	_ = v97
	var v99 int32
	_ = v99
	var v104 int32
	_ = v104
	var v105 int32
	_ = v105
	var v110 int32
	_ = v110
	var v111 int32
	_ = v111
	var v112 int32
	_ = v112
	var v113 int32
	_ = v113
	var v116 int32
	_ = v116
	var v118 int32
	_ = v118
	var v122 int64
	_ = v122
	var v124 int32
	_ = v124
	var v125 int32
	_ = v125
	var v130 int32
	_ = v130
	var v131 int32
	_ = v131
	var v132 int32
	_ = v132
	var v133 int32
	_ = v133
	var v136 int32
	_ = v136
	var v140 int32
	_ = v140
	var v142 int32
	_ = v142
	var v143 int32
	_ = v143
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
	var v151 int32
	_ = v151
	var v152 int32
	_ = v152
	var v153 int32
	_ = v153
	var v154 int32
	_ = v154
	var v155 int32
	_ = v155
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
	var v172 int32
	_ = v172
	var v174 int32
	_ = v174
	var v176 int32
	_ = v176
	var v178 int32
	_ = v178
	var v180 int32
	_ = v180
	var v181 int32
	_ = v181
	var v185 int32
	_ = v185
	var v193 int32
	_ = v193
	var v197 int32
	_ = v197
	var v210 int32
	_ = v210
	var v211 int32
	_ = v211
	var v218 int32
	_ = v218
	var v223 int32
	_ = v223
	var v227 int32
	_ = v227
	var v230 int32
	_ = v230
	var v231 int32
	_ = v231
	var v232 int32
	_ = v232
	var v238 int32
	_ = v238
	var v243 int32
	_ = v243
	var v247 int32
	_ = v247
	var v253 int32
	_ = v253
	var v258 int32
	_ = v258
	v14 = m.G0
	v16 = v14 - int32(48)
	m.G0 = v16
	v18 = *(*int32)(unsafe.Add(mBase, uint32(l0)+56))
	v19 = F_get_object_catcache_oid(m, v18)
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
	v23 = F_get_object_catcache_name(m, v18)
	mBase = m.M
	v24 = m.ExcPending
	if v24 != 0 {
		goto L1
	} else {
		goto L3
	}
L3:
	;
	v25 = F_get_object_attnum_name(m, v18)
	mBase = m.M
	v26 = m.ExcPending
	if v26 != 0 {
		goto L1
	} else {
		goto L4
	}
L4:
	;
	v27 = F_get_object_attnum_namespace(m, v18)
	mBase = m.M
	v28 = m.ExcPending
	if v28 != 0 {
		goto L1
	} else {
		goto L5
	}
L5:
	;
	v29 = F_get_object_attnum_owner(m, v18)
	mBase = m.M
	v30 = m.ExcPending
	if v30 != 0 {
		goto L1
	} else {
		goto L6
	}
L6:
	;
	v33 = F_SearchSysCacheCopy(m, v19, base.I64_extend_i32_u(l1), int64(0))
	mBase = m.M
	v34 = m.ExcPending
	if v34 != 0 {
		goto L1
	} else {
		goto L9
	}
L7:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v247 = m.ExcPending
	if v247 != 0 {
		goto L1
	} else {
		goto L77
	}
L8:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v227 = m.ExcPending
	if v227 != 0 {
		goto L1
	} else {
		goto L72
	}
L9:
	;
	if v33 != 0 {
		goto L10
	} else {
		goto L11
	}
L10:
	;
	v35 = *(*int32)(unsafe.Add(mBase, uint32(l0)+52))
	v37 = v16 + int32(47)
	v38 = F_heap_getattr_2(m, v33, v25, v35, v37)
	mBase = m.M
	v39 = m.ExcPending
	if v39 != 0 {
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
	v210 = m.ExcPending
	if v210 != 0 {
		goto L1
	} else {
		goto L69
	}
L13:
	;
	v40 = *(*int32)(unsafe.Add(mBase, uint32(l0)+52))
	v41 = F_heap_getattr_2(m, v33, v27, v40, v37)
	mBase = m.M
	v42 = m.ExcPending
	if v42 != 0 {
		goto L1
	} else {
		goto L16
	}
L14:
	;
	m.G0 = v16 + int32(48)
	return v43
L15:
	;
	v193 = int32(0)
	F_RunObjectPostAlterHook(m, v18, l1, v193, v193, v193)
	mBase = m.M
	v197 = m.ExcPending
	if v197 != 0 {
		goto L1
	} else {
		goto L68
	}
L16:
	;
	v43 = base.I32_wrap_i64(v41)
	if v43 == l2 {
		goto L17
	} else {
		goto L18
	}
L17:
	;
	v46 = *(*int32)(unsafe.Add(mBase, _c_F_AlterObjectNamespace_internal[0]))
	if v46 != 0 {
		goto L15
	} else {
		goto L20
	}
L18:
	;
	goto L19
L19:
	;
	F_CheckSetNamespace(m, v43, l2)
	mBase = m.M
	v48 = m.ExcPending
	if v48 != 0 {
		goto L1
	} else {
		goto L21
	}
L20:
	;
	goto L14
L21:
	;
	v49 = F_superuser(m)
	mBase = m.M
	v50 = m.ExcPending
	if v50 != 0 {
		goto L1
	} else {
		goto L23
	}
L22:
	;
	if v18 <= int32(2752) {
		goto L40
	} else {
		goto L41
	}
L23:
	;
	if v49 != 0 {
		goto L22
	} else {
		goto L24
	}
L24:
	;
	if v29 <= int32(0) {
		goto L8
	} else {
		goto L25
	}
L25:
	;
	v53 = *(*int32)(unsafe.Add(mBase, uint32(l0)+52))
	v56 = F_heap_getattr_2(m, v33, v29, v53, v16+int32(47))
	mBase = m.M
	v57 = m.ExcPending
	if v57 != 0 {
		goto L1
	} else {
		goto L26
	}
L26:
	;
	v59 = *(*int32)(unsafe.Add(mBase, _c_F_AlterObjectNamespace_internal[1]))
	v61 = F_has_privs_of_role(m, v59, base.I32_wrap_i64(v56))
	mBase = m.M
	v62 = m.ExcPending
	if v62 != 0 {
		goto L1
	} else {
		goto L27
	}
L27:
	;
	if v61 == int32(0) {
		goto L28
	} else {
		goto L29
	}
L28:
	;
	v66 = F_get_object_type(m, v18, l1)
	mBase = m.M
	v67 = m.ExcPending
	if v67 != 0 {
		goto L1
	} else {
		goto L31
	}
L29:
	;
	goto L30
L30:
	;
	v73 = *(*int32)(unsafe.Add(mBase, _c_F_AlterObjectNamespace_internal[1]))
	v75 = F_object_aclcheck(m, int32(2615), l2, v73, int64(512))
	mBase = m.M
	v76 = m.ExcPending
	if v76 != 0 {
		goto L1
	} else {
		goto L33
	}
L31:
	;
	F_aclcheck_error(m, int32(2), v66, base.I32_wrap_i64(v38))
	mBase = m.M
	v70 = m.ExcPending
	if v70 != 0 {
		goto L1
	} else {
		goto L32
	}
L32:
	;
	goto L30
L33:
	;
	if v75 == int32(0) {
		goto L22
	} else {
		goto L34
	}
L34:
	;
	v80 = F_get_namespace_name(m, l2)
	mBase = m.M
	v81 = m.ExcPending
	if v81 != 0 {
		goto L1
	} else {
		goto L35
	}
L35:
	;
	F_aclcheck_error(m, v75, int32(37), v80)
	mBase = m.M
	v83 = m.ExcPending
	if v83 != 0 {
		goto L1
	} else {
		goto L36
	}
L36:
	;
	goto L22
L37:
	;
	v142 = *(*int32)(unsafe.Add(mBase, uint32(l0)+48))
	v143 = int32(*(*int16)(unsafe.Add(mBase, uint32(v142)+120)))
	v146 = F_palloc0(m, v143<<(uint(int32(3))%32))
	mBase = m.M
	v147 = m.ExcPending
	if v147 != 0 {
		goto L1
	} else {
		goto L57
	}
L38:
	;
	v131 = *(*int32)(unsafe.Add(mBase, uint32(v33)+16))
	v132 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v131)+22)))
	v133 = v131 + v132
	v136 = int32(*(*int16)(unsafe.Add(mBase, uint32(v133)+104)))
	F_IsThereFunctionInNamespace(m, v133+int32(4), v136, v133+int32(112), l2)
	mBase = m.M
	v140 = m.ExcPending
	if v140 != 0 {
		goto L1
	} else {
		goto L56
	}
L39:
	;
	if v23 < int32(0) {
		goto L37
	} else {
		goto L52
	}
L40:
	;
	if v18 == int32(1255) {
		goto L38
	} else {
		goto L43
	}
L41:
	;
	goto L42
L42:
	;
	if v18 != int32(2753) {
		goto L46
	} else {
		goto L47
	}
L43:
	;
	if v18 != int32(2616) {
		goto L39
	} else {
		goto L44
	}
L44:
	;
	v92 = *(*int32)(unsafe.Add(mBase, uint32(v33)+16))
	v93 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v92)+22)))
	v94 = v92 + v93
	v97 = *(*int32)(unsafe.Add(mBase, uint32(v94)+4))
	F_IsThereOpClassInNamespace(m, v94+int32(8), v97, l2)
	mBase = m.M
	v99 = m.ExcPending
	if v99 != 0 {
		goto L1
	} else {
		goto L45
	}
L45:
	;
	goto L37
L46:
	;
	if v18 != int32(3456) {
		goto L39
	} else {
		goto L49
	}
L47:
	;
	goto L48
L48:
	;
	v111 = *(*int32)(unsafe.Add(mBase, uint32(v33)+16))
	v112 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v111)+22)))
	v113 = v111 + v112
	v116 = *(*int32)(unsafe.Add(mBase, uint32(v113)+4))
	F_IsThereOpFamilyInNamespace(m, v113+int32(8), v116, l2)
	mBase = m.M
	v118 = m.ExcPending
	if v118 != 0 {
		goto L1
	} else {
		goto L51
	}
L49:
	;
	v104 = *(*int32)(unsafe.Add(mBase, uint32(v33)+16))
	v105 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v104)+22)))
	F_IsThereCollationInNamespace(m, v104+v105+int32(4), l2)
	mBase = m.M
	v110 = m.ExcPending
	if v110 != 0 {
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
	v122 = int64(0)
	v124 = F_SearchSysCacheExists(m, v23, v38, base.I64_extend_i32_u(l2), v122, v122)
	mBase = m.M
	v125 = m.ExcPending
	if v125 != 0 {
		goto L1
	} else {
		goto L53
	}
L53:
	;
	if v124 == int32(0) {
		goto L37
	} else {
		goto L54
	}
L54:
	;
	F_report_namespace_conflict(m, v18, base.I32_wrap_i64(v38), l2)
	mBase = m.M
	v130 = m.ExcPending
	if v130 != 0 {
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
	v148 = *(*int32)(unsafe.Add(mBase, uint32(l0)+48))
	v149 = int32(*(*int16)(unsafe.Add(mBase, uint32(v148)+120)))
	v150 = F_palloc0(m, v149)
	mBase = m.M
	v151 = m.ExcPending
	if v151 != 0 {
		goto L1
	} else {
		goto L58
	}
L58:
	;
	v152 = *(*int32)(unsafe.Add(mBase, uint32(l0)+48))
	v153 = int32(*(*int16)(unsafe.Add(mBase, uint32(v152)+120)))
	v154 = F_palloc0(m, v153)
	mBase = m.M
	v155 = m.ExcPending
	if v155 != 0 {
		goto L1
	} else {
		goto L59
	}
L59:
	;
	v156 = int32(1)
	v157 = v27 - v156
	*(*int64)(unsafe.Add(mBase, uint32(v146+v157<<(uint(int32(3))%32)))) = base.I64_extend_i32_u(l2)
	*(*uint8)(unsafe.Add(mBase, uint32(v154+v157))) = uint8(v156)
	v168 = *(*int32)(unsafe.Add(mBase, uint32(l0)+52))
	v169 = F_heap_modify_tuple(m, v33, v168, v146, v150, v154)
	mBase = m.M
	v170 = m.ExcPending
	if v170 != 0 {
		goto L1
	} else {
		goto L60
	}
L60:
	;
	F_CatalogTupleUpdate(m, l0, v33+int32(4), v169)
	mBase = m.M
	v172 = m.ExcPending
	if v172 != 0 {
		goto L1
	} else {
		goto L61
	}
L61:
	;
	F_pfree(m, v146)
	mBase = m.M
	v174 = m.ExcPending
	if v174 != 0 {
		goto L1
	} else {
		goto L62
	}
L62:
	;
	F_pfree(m, v150)
	mBase = m.M
	v176 = m.ExcPending
	if v176 != 0 {
		goto L1
	} else {
		goto L63
	}
L63:
	;
	F_pfree(m, v154)
	mBase = m.M
	v178 = m.ExcPending
	if v178 != 0 {
		goto L1
	} else {
		goto L64
	}
L64:
	;
	v180 = F_changeDependencyFor(m, v18, l1, int32(2615), v43, l2)
	mBase = m.M
	v181 = m.ExcPending
	if v181 != 0 {
		goto L1
	} else {
		goto L65
	}
L65:
	;
	if v180 != int32(1) {
		goto L7
	} else {
		goto L66
	}
L66:
	;
	v185 = *(*int32)(unsafe.Add(mBase, _c_F_AlterObjectNamespace_internal[0]))
	if v185 == int32(0) {
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
	v211 = *(*int32)(unsafe.Add(mBase, uint32(l0)+48))
	*(*int32)(unsafe.Add(mBase, uint32(v16))) = l1
	*(*int32)(unsafe.Add(mBase, uint32(v16)+4)) = v211 + int32(4)
	F_errmsg_internal(m, int32(_a_F_AlterObjectNamespace_internal_4), v16)
	mBase = m.M
	v218 = m.ExcPending
	if v218 != 0 {
		goto L1
	} else {
		goto L70
	}
L70:
	;
	F_errfinish(m, int32(_a_F_AlterObjectNamespace_internal_1), int32(707), int32(_a_F_AlterObjectNamespace_internal_2))
	mBase = m.M
	v223 = m.ExcPending
	if v223 != 0 {
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
	v230 = m.ExcPending
	if v230 != 0 {
		goto L1
	} else {
		goto L73
	}
L73:
	;
	v231 = F_getObjectDescriptionOids(m, v18, l1)
	mBase = m.M
	v232 = m.ExcPending
	if v232 != 0 {
		goto L1
	} else {
		goto L74
	}
L74:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v16)+32)) = v231
	F_errmsg(m, int32(_a_F_AlterObjectNamespace_internal_3), v16+int32(32))
	mBase = m.M
	v238 = m.ExcPending
	if v238 != 0 {
		goto L1
	} else {
		goto L75
	}
L75:
	;
	F_errfinish(m, int32(_a_F_AlterObjectNamespace_internal_1), int32(741), int32(_a_F_AlterObjectNamespace_internal_2))
	mBase = m.M
	v243 = m.ExcPending
	if v243 != 0 {
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
	*(*int32)(unsafe.Add(mBase, uint32(v16)+16)) = l1
	F_errmsg_internal(m, int32(_a_F_AlterObjectNamespace_internal_0), v16+int32(16))
	mBase = m.M
	v253 = m.ExcPending
	if v253 != 0 {
		goto L1
	} else {
		goto L78
	}
L78:
	;
	F_errfinish(m, int32(_a_F_AlterObjectNamespace_internal_1), int32(819), int32(_a_F_AlterObjectNamespace_internal_2))
	mBase = m.M
	v258 = m.ExcPending
	if v258 != 0 {
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
	var v8 int32
	_ = v8
	var v10 int32
	_ = v10
	var v12 int32
	_ = v12
	var v13 int32
	_ = v13
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
	var v22 int32
	_ = v22
	var v30 int32
	_ = v30
	var v32 int32
	_ = v32
	var v34 int32
	_ = v34
	var v37 int32
	_ = v37
	var v39 int32
	_ = v39
	var v41 int32
	_ = v41
	var v42 int32
	_ = v42
	var v45 int32
	_ = v45
	var v46 int32
	_ = v46
	var v47 int32
	_ = v47
	var v49 int32
	_ = v49
	var v60 int64
	_ = v60
	var v61 int32
	_ = v61
	var v62 int32
	_ = v62
	var v66 int32
	_ = v66
	var v67 int32
	_ = v67
	var v68 int32
	_ = v68
	var v69 int32
	_ = v69
	var v70 int32
	_ = v70
	var v71 int32
	_ = v71
	var v76 int32
	_ = v76
	var v83 int32
	_ = v83
	var v84 int32
	_ = v84
	var v88 int32
	_ = v88
	var v90 int32
	_ = v90
	var v92 int32
	_ = v92
	var v94 int32
	_ = v94
	var v99 int32
	_ = v99
	var v101 int32
	_ = v101
	var v102 int32
	_ = v102
	var v106 int32
	_ = v106
	v8 = m.G0
	v10 = v8 - int32(48)
	m.G0 = v10
	v12 = *(*int32)(unsafe.Add(mBase, uint32(l0)+16))
	v13 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v12)+22)))
	v14 = v12 + v13
	v15 = *(*int32)(unsafe.Add(mBase, uint32(v14)+68))
	if l2 != v15 {
		v18 = *(*int32)(unsafe.Add(mBase, uint32(v14)))
		v20 = *(*int32)(unsafe.Add(mBase, _c_F_AlterSchemaOwner_internal[0]))
		v21 = F_object_ownercheck(m, int32(2615), v18, v20)
		mBase = m.M
		v22 = m.ExcPending
		if v22 != 0 {
			return
		} else {
			if v21 == int32(0) {
				F_aclcheck_error(m, int32(2), int32(37), v14+int32(4))
				mBase = m.M
				v30 = m.ExcPending
				if v30 != 0 {
					return
				} else {
					v32 = *(*int32)(unsafe.Add(mBase, _c_F_AlterSchemaOwner_internal[0]))
					F_check_can_set_role(m, v32, l2)
					mBase = m.M
					v34 = m.ExcPending
					if v34 != 0 {
						return
					} else {
						v37 = *(*int32)(unsafe.Add(mBase, _c_F_AlterSchemaOwner_internal[1]))
						v39 = *(*int32)(unsafe.Add(mBase, _c_F_AlterSchemaOwner_internal[0]))
						v41 = F_object_aclcheck(m, int32(1262), v37, v39, int64(512))
						mBase = m.M
						v42 = m.ExcPending
						if v42 != 0 {
							return
						} else {
							if v41 != 0 {
								v45 = *(*int32)(unsafe.Add(mBase, _c_F_AlterSchemaOwner_internal[1]))
								v46 = F_get_database_name(m, v45)
								mBase = m.M
								v47 = m.ExcPending
								if v47 != 0 {
									return
								} else {
									F_aclcheck_error(m, v41, int32(9), v46)
									mBase = m.M
									v49 = m.ExcPending
									if v49 != 0 {
										return
									} else {
										*(*int32)(unsafe.Add(mBase, uint32(v10)+8)) = int32(_a_F_AlterSchemaOwner_internal_0)
										*(*int32)(unsafe.Add(mBase, uint32(v10)+12)) = int32(0)
										*(*int64)(unsafe.Add(mBase, uint32(v10)+32)) = base.I64_extend_i32_u(l2)
										v60 = F_SysCacheGetAttr(m, int32(37), l0, int32(4), v10+int32(7))
										mBase = m.M
										v61 = m.ExcPending
										if v61 != 0 {
											return
										} else {
											v62 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v10)+7)))
											if v62 == int32(0) {
												v66 = F_pg_detoast_datum(m, base.I32_wrap_i64(v60))
												mBase = m.M
												v67 = m.ExcPending
												if v67 != 0 {
													return
												} else {
													v68 = *(*int32)(unsafe.Add(mBase, uint32(v14)+68))
													v69 = F_aclnewowner(m, v66, v68, l2)
													mBase = m.M
													v70 = m.ExcPending
													if v70 != 0 {
														return
													} else {
														v71 = int32(1)
														*(*uint8)(unsafe.Add(mBase, uint32(v10)+11)) = uint8(v71)
														*(*int64)(unsafe.Add(mBase, uint32(v10)+40)) = base.I64_extend_i32_u(v69)
														v76 = *(*int32)(unsafe.Add(mBase, uint32(l1)+52))
														v83 = F_heap_modify_tuple(m, l0, v76, v10+int32(16), v10+int32(12), v10+int32(8))
														mBase = m.M
														v84 = m.ExcPending
														if v84 != 0 {
															return
														} else {
															F_CatalogTupleUpdate(m, l1, v83+int32(4), v83)
															mBase = m.M
															v88 = m.ExcPending
															if v88 != 0 {
																return
															} else {
																F_pfree(m, v83)
																mBase = m.M
																v90 = m.ExcPending
																if v90 != 0 {
																	return
																} else {
																	v92 = *(*int32)(unsafe.Add(mBase, uint32(v14)))
																	F_changeDependencyOnOwner(m, int32(2615), v92, l2)
																	mBase = m.M
																	v94 = m.ExcPending
																	if v94 != 0 {
																		return
																	} else {
																		v99 = *(*int32)(unsafe.Add(mBase, _c_F_AlterSchemaOwner_internal[2]))
																		if v99 != 0 {
																			v101 = *(*int32)(unsafe.Add(mBase, uint32(v14)))
																			v102 = int32(0)
																			F_RunObjectPostAlterHook(m, int32(2615), v101, v102, v102, v102)
																			mBase = m.M
																			v106 = m.ExcPending
																			if v106 != 0 {
																				return
																			} else {
																				m.G0 = v10 + int32(48)
																				return
																			}
																		} else {
																			m.G0 = v10 + int32(48)
																			return
																		}
																	}
																}
															}
														}
													}
												}
											} else {
												v76 = *(*int32)(unsafe.Add(mBase, uint32(l1)+52))
												v83 = F_heap_modify_tuple(m, l0, v76, v10+int32(16), v10+int32(12), v10+int32(8))
												mBase = m.M
												v84 = m.ExcPending
												if v84 != 0 {
													return
												} else {
													F_CatalogTupleUpdate(m, l1, v83+int32(4), v83)
													mBase = m.M
													v88 = m.ExcPending
													if v88 != 0 {
														return
													} else {
														F_pfree(m, v83)
														mBase = m.M
														v90 = m.ExcPending
														if v90 != 0 {
															return
														} else {
															v92 = *(*int32)(unsafe.Add(mBase, uint32(v14)))
															F_changeDependencyOnOwner(m, int32(2615), v92, l2)
															mBase = m.M
															v94 = m.ExcPending
															if v94 != 0 {
																return
															} else {
																v99 = *(*int32)(unsafe.Add(mBase, _c_F_AlterSchemaOwner_internal[2]))
																if v99 != 0 {
																	v101 = *(*int32)(unsafe.Add(mBase, uint32(v14)))
																	v102 = int32(0)
																	F_RunObjectPostAlterHook(m, int32(2615), v101, v102, v102, v102)
																	mBase = m.M
																	v106 = m.ExcPending
																	if v106 != 0 {
																		return
																	} else {
																		m.G0 = v10 + int32(48)
																		return
																	}
																} else {
																	m.G0 = v10 + int32(48)
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
								*(*int32)(unsafe.Add(mBase, uint32(v10)+8)) = int32(_a_F_AlterSchemaOwner_internal_0)
								*(*int32)(unsafe.Add(mBase, uint32(v10)+12)) = int32(0)
								*(*int64)(unsafe.Add(mBase, uint32(v10)+32)) = base.I64_extend_i32_u(l2)
								v60 = F_SysCacheGetAttr(m, int32(37), l0, int32(4), v10+int32(7))
								mBase = m.M
								v61 = m.ExcPending
								if v61 != 0 {
									return
								} else {
									v62 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v10)+7)))
									if v62 == int32(0) {
										v66 = F_pg_detoast_datum(m, base.I32_wrap_i64(v60))
										mBase = m.M
										v67 = m.ExcPending
										if v67 != 0 {
											return
										} else {
											v68 = *(*int32)(unsafe.Add(mBase, uint32(v14)+68))
											v69 = F_aclnewowner(m, v66, v68, l2)
											mBase = m.M
											v70 = m.ExcPending
											if v70 != 0 {
												return
											} else {
												v71 = int32(1)
												*(*uint8)(unsafe.Add(mBase, uint32(v10)+11)) = uint8(v71)
												*(*int64)(unsafe.Add(mBase, uint32(v10)+40)) = base.I64_extend_i32_u(v69)
												v76 = *(*int32)(unsafe.Add(mBase, uint32(l1)+52))
												v83 = F_heap_modify_tuple(m, l0, v76, v10+int32(16), v10+int32(12), v10+int32(8))
												mBase = m.M
												v84 = m.ExcPending
												if v84 != 0 {
													return
												} else {
													F_CatalogTupleUpdate(m, l1, v83+int32(4), v83)
													mBase = m.M
													v88 = m.ExcPending
													if v88 != 0 {
														return
													} else {
														F_pfree(m, v83)
														mBase = m.M
														v90 = m.ExcPending
														if v90 != 0 {
															return
														} else {
															v92 = *(*int32)(unsafe.Add(mBase, uint32(v14)))
															F_changeDependencyOnOwner(m, int32(2615), v92, l2)
															mBase = m.M
															v94 = m.ExcPending
															if v94 != 0 {
																return
															} else {
																v99 = *(*int32)(unsafe.Add(mBase, _c_F_AlterSchemaOwner_internal[2]))
																if v99 != 0 {
																	v101 = *(*int32)(unsafe.Add(mBase, uint32(v14)))
																	v102 = int32(0)
																	F_RunObjectPostAlterHook(m, int32(2615), v101, v102, v102, v102)
																	mBase = m.M
																	v106 = m.ExcPending
																	if v106 != 0 {
																		return
																	} else {
																		m.G0 = v10 + int32(48)
																		return
																	}
																} else {
																	m.G0 = v10 + int32(48)
																	return
																}
															}
														}
													}
												}
											}
										}
									} else {
										v76 = *(*int32)(unsafe.Add(mBase, uint32(l1)+52))
										v83 = F_heap_modify_tuple(m, l0, v76, v10+int32(16), v10+int32(12), v10+int32(8))
										mBase = m.M
										v84 = m.ExcPending
										if v84 != 0 {
											return
										} else {
											F_CatalogTupleUpdate(m, l1, v83+int32(4), v83)
											mBase = m.M
											v88 = m.ExcPending
											if v88 != 0 {
												return
											} else {
												F_pfree(m, v83)
												mBase = m.M
												v90 = m.ExcPending
												if v90 != 0 {
													return
												} else {
													v92 = *(*int32)(unsafe.Add(mBase, uint32(v14)))
													F_changeDependencyOnOwner(m, int32(2615), v92, l2)
													mBase = m.M
													v94 = m.ExcPending
													if v94 != 0 {
														return
													} else {
														v99 = *(*int32)(unsafe.Add(mBase, _c_F_AlterSchemaOwner_internal[2]))
														if v99 != 0 {
															v101 = *(*int32)(unsafe.Add(mBase, uint32(v14)))
															v102 = int32(0)
															F_RunObjectPostAlterHook(m, int32(2615), v101, v102, v102, v102)
															mBase = m.M
															v106 = m.ExcPending
															if v106 != 0 {
																return
															} else {
																m.G0 = v10 + int32(48)
																return
															}
														} else {
															m.G0 = v10 + int32(48)
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
				v32 = *(*int32)(unsafe.Add(mBase, _c_F_AlterSchemaOwner_internal[0]))
				F_check_can_set_role(m, v32, l2)
				mBase = m.M
				v34 = m.ExcPending
				if v34 != 0 {
					return
				} else {
					v37 = *(*int32)(unsafe.Add(mBase, _c_F_AlterSchemaOwner_internal[1]))
					v39 = *(*int32)(unsafe.Add(mBase, _c_F_AlterSchemaOwner_internal[0]))
					v41 = F_object_aclcheck(m, int32(1262), v37, v39, int64(512))
					mBase = m.M
					v42 = m.ExcPending
					if v42 != 0 {
						return
					} else {
						if v41 != 0 {
							v45 = *(*int32)(unsafe.Add(mBase, _c_F_AlterSchemaOwner_internal[1]))
							v46 = F_get_database_name(m, v45)
							mBase = m.M
							v47 = m.ExcPending
							if v47 != 0 {
								return
							} else {
								F_aclcheck_error(m, v41, int32(9), v46)
								mBase = m.M
								v49 = m.ExcPending
								if v49 != 0 {
									return
								} else {
									*(*int32)(unsafe.Add(mBase, uint32(v10)+8)) = int32(_a_F_AlterSchemaOwner_internal_0)
									*(*int32)(unsafe.Add(mBase, uint32(v10)+12)) = int32(0)
									*(*int64)(unsafe.Add(mBase, uint32(v10)+32)) = base.I64_extend_i32_u(l2)
									v60 = F_SysCacheGetAttr(m, int32(37), l0, int32(4), v10+int32(7))
									mBase = m.M
									v61 = m.ExcPending
									if v61 != 0 {
										return
									} else {
										v62 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v10)+7)))
										if v62 == int32(0) {
											v66 = F_pg_detoast_datum(m, base.I32_wrap_i64(v60))
											mBase = m.M
											v67 = m.ExcPending
											if v67 != 0 {
												return
											} else {
												v68 = *(*int32)(unsafe.Add(mBase, uint32(v14)+68))
												v69 = F_aclnewowner(m, v66, v68, l2)
												mBase = m.M
												v70 = m.ExcPending
												if v70 != 0 {
													return
												} else {
													v71 = int32(1)
													*(*uint8)(unsafe.Add(mBase, uint32(v10)+11)) = uint8(v71)
													*(*int64)(unsafe.Add(mBase, uint32(v10)+40)) = base.I64_extend_i32_u(v69)
													v76 = *(*int32)(unsafe.Add(mBase, uint32(l1)+52))
													v83 = F_heap_modify_tuple(m, l0, v76, v10+int32(16), v10+int32(12), v10+int32(8))
													mBase = m.M
													v84 = m.ExcPending
													if v84 != 0 {
														return
													} else {
														F_CatalogTupleUpdate(m, l1, v83+int32(4), v83)
														mBase = m.M
														v88 = m.ExcPending
														if v88 != 0 {
															return
														} else {
															F_pfree(m, v83)
															mBase = m.M
															v90 = m.ExcPending
															if v90 != 0 {
																return
															} else {
																v92 = *(*int32)(unsafe.Add(mBase, uint32(v14)))
																F_changeDependencyOnOwner(m, int32(2615), v92, l2)
																mBase = m.M
																v94 = m.ExcPending
																if v94 != 0 {
																	return
																} else {
																	v99 = *(*int32)(unsafe.Add(mBase, _c_F_AlterSchemaOwner_internal[2]))
																	if v99 != 0 {
																		v101 = *(*int32)(unsafe.Add(mBase, uint32(v14)))
																		v102 = int32(0)
																		F_RunObjectPostAlterHook(m, int32(2615), v101, v102, v102, v102)
																		mBase = m.M
																		v106 = m.ExcPending
																		if v106 != 0 {
																			return
																		} else {
																			m.G0 = v10 + int32(48)
																			return
																		}
																	} else {
																		m.G0 = v10 + int32(48)
																		return
																	}
																}
															}
														}
													}
												}
											}
										} else {
											v76 = *(*int32)(unsafe.Add(mBase, uint32(l1)+52))
											v83 = F_heap_modify_tuple(m, l0, v76, v10+int32(16), v10+int32(12), v10+int32(8))
											mBase = m.M
											v84 = m.ExcPending
											if v84 != 0 {
												return
											} else {
												F_CatalogTupleUpdate(m, l1, v83+int32(4), v83)
												mBase = m.M
												v88 = m.ExcPending
												if v88 != 0 {
													return
												} else {
													F_pfree(m, v83)
													mBase = m.M
													v90 = m.ExcPending
													if v90 != 0 {
														return
													} else {
														v92 = *(*int32)(unsafe.Add(mBase, uint32(v14)))
														F_changeDependencyOnOwner(m, int32(2615), v92, l2)
														mBase = m.M
														v94 = m.ExcPending
														if v94 != 0 {
															return
														} else {
															v99 = *(*int32)(unsafe.Add(mBase, _c_F_AlterSchemaOwner_internal[2]))
															if v99 != 0 {
																v101 = *(*int32)(unsafe.Add(mBase, uint32(v14)))
																v102 = int32(0)
																F_RunObjectPostAlterHook(m, int32(2615), v101, v102, v102, v102)
																mBase = m.M
																v106 = m.ExcPending
																if v106 != 0 {
																	return
																} else {
																	m.G0 = v10 + int32(48)
																	return
																}
															} else {
																m.G0 = v10 + int32(48)
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
							*(*int32)(unsafe.Add(mBase, uint32(v10)+8)) = int32(_a_F_AlterSchemaOwner_internal_0)
							*(*int32)(unsafe.Add(mBase, uint32(v10)+12)) = int32(0)
							*(*int64)(unsafe.Add(mBase, uint32(v10)+32)) = base.I64_extend_i32_u(l2)
							v60 = F_SysCacheGetAttr(m, int32(37), l0, int32(4), v10+int32(7))
							mBase = m.M
							v61 = m.ExcPending
							if v61 != 0 {
								return
							} else {
								v62 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v10)+7)))
								if v62 == int32(0) {
									v66 = F_pg_detoast_datum(m, base.I32_wrap_i64(v60))
									mBase = m.M
									v67 = m.ExcPending
									if v67 != 0 {
										return
									} else {
										v68 = *(*int32)(unsafe.Add(mBase, uint32(v14)+68))
										v69 = F_aclnewowner(m, v66, v68, l2)
										mBase = m.M
										v70 = m.ExcPending
										if v70 != 0 {
											return
										} else {
											v71 = int32(1)
											*(*uint8)(unsafe.Add(mBase, uint32(v10)+11)) = uint8(v71)
											*(*int64)(unsafe.Add(mBase, uint32(v10)+40)) = base.I64_extend_i32_u(v69)
											v76 = *(*int32)(unsafe.Add(mBase, uint32(l1)+52))
											v83 = F_heap_modify_tuple(m, l0, v76, v10+int32(16), v10+int32(12), v10+int32(8))
											mBase = m.M
											v84 = m.ExcPending
											if v84 != 0 {
												return
											} else {
												F_CatalogTupleUpdate(m, l1, v83+int32(4), v83)
												mBase = m.M
												v88 = m.ExcPending
												if v88 != 0 {
													return
												} else {
													F_pfree(m, v83)
													mBase = m.M
													v90 = m.ExcPending
													if v90 != 0 {
														return
													} else {
														v92 = *(*int32)(unsafe.Add(mBase, uint32(v14)))
														F_changeDependencyOnOwner(m, int32(2615), v92, l2)
														mBase = m.M
														v94 = m.ExcPending
														if v94 != 0 {
															return
														} else {
															v99 = *(*int32)(unsafe.Add(mBase, _c_F_AlterSchemaOwner_internal[2]))
															if v99 != 0 {
																v101 = *(*int32)(unsafe.Add(mBase, uint32(v14)))
																v102 = int32(0)
																F_RunObjectPostAlterHook(m, int32(2615), v101, v102, v102, v102)
																mBase = m.M
																v106 = m.ExcPending
																if v106 != 0 {
																	return
																} else {
																	m.G0 = v10 + int32(48)
																	return
																}
															} else {
																m.G0 = v10 + int32(48)
																return
															}
														}
													}
												}
											}
										}
									}
								} else {
									v76 = *(*int32)(unsafe.Add(mBase, uint32(l1)+52))
									v83 = F_heap_modify_tuple(m, l0, v76, v10+int32(16), v10+int32(12), v10+int32(8))
									mBase = m.M
									v84 = m.ExcPending
									if v84 != 0 {
										return
									} else {
										F_CatalogTupleUpdate(m, l1, v83+int32(4), v83)
										mBase = m.M
										v88 = m.ExcPending
										if v88 != 0 {
											return
										} else {
											F_pfree(m, v83)
											mBase = m.M
											v90 = m.ExcPending
											if v90 != 0 {
												return
											} else {
												v92 = *(*int32)(unsafe.Add(mBase, uint32(v14)))
												F_changeDependencyOnOwner(m, int32(2615), v92, l2)
												mBase = m.M
												v94 = m.ExcPending
												if v94 != 0 {
													return
												} else {
													v99 = *(*int32)(unsafe.Add(mBase, _c_F_AlterSchemaOwner_internal[2]))
													if v99 != 0 {
														v101 = *(*int32)(unsafe.Add(mBase, uint32(v14)))
														v102 = int32(0)
														F_RunObjectPostAlterHook(m, int32(2615), v101, v102, v102, v102)
														mBase = m.M
														v106 = m.ExcPending
														if v106 != 0 {
															return
														} else {
															m.G0 = v10 + int32(48)
															return
														}
													} else {
														m.G0 = v10 + int32(48)
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
		v99 = *(*int32)(unsafe.Add(mBase, _c_F_AlterSchemaOwner_internal[2]))
		if v99 != 0 {
			v101 = *(*int32)(unsafe.Add(mBase, uint32(v14)))
			v102 = int32(0)
			F_RunObjectPostAlterHook(m, int32(2615), v101, v102, v102, v102)
			mBase = m.M
			v106 = m.ExcPending
			if v106 != 0 {
				return
			} else {
				m.G0 = v10 + int32(48)
				return
			}
		} else {
			m.G0 = v10 + int32(48)
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
	var v51 int32
	_ = v51
	var v53 int32
	_ = v53
	var v58 int32
	_ = v58
	var v60 int32
	_ = v60
	var v62 int32
	_ = v62
	var v64 int32
	_ = v64
	var v66 int32
	_ = v66
	var v67 int32
	_ = v67
	var v71 int32
	_ = v71
	var v73 int32
	_ = v73
	var v77 int32
	_ = v77
	var v79 int32
	_ = v79
	var v81 int32
	_ = v81
	var v86 int32
	_ = v86
	var v89 int32
	_ = v89
	var v93 int32
	_ = v93
	var v97 int32
	_ = v97
	var v102 int32
	_ = v102
	v6 = *(*int32)(unsafe.Add(mBase, uint32(l1)+16))
	v7 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v6)+22)))
	v8 = v6 + v7
	v9 = *(*int32)(unsafe.Add(mBase, uint32(v8)+80))
	if l2 != v9 {
		goto L2
	} else {
		goto L3
	}
L1:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v86 = m.ExcPending
	if v86 != 0 {
		goto L5
	} else {
		goto L38
	}
L2:
	;
	v12 = *(*int32)(unsafe.Add(mBase, uint32(v8)))
	v14 = *(*int32)(unsafe.Add(mBase, _c_F_AlterSubscriptionOwner_internal[0]))
	v15 = F_object_ownercheck(m, int32(_a_F_AlterSubscriptionOwner_internal_0), v12, v14)
	mBase = m.M
	v16 = m.ExcPending
	if v16 != 0 {
		goto L5
	} else {
		goto L6
	}
L3:
	;
	goto L4
L4:
	;
	return
L5:
	;
	return
L6:
	;
	if v15 == int32(0) {
		goto L7
	} else {
		goto L8
	}
L7:
	;
	F_aclcheck_error(m, int32(2), int32(39), v8+int32(16))
	mBase = m.M
	v24 = m.ExcPending
	if v24 != 0 {
		goto L5
	} else {
		goto L10
	}
L8:
	;
	goto L9
L9:
	;
	v25 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v8)+89)))
	if v25 == int32(0) {
		goto L11
	} else {
		goto L12
	}
L10:
	;
	goto L9
L11:
	;
	v28 = F_superuser(m)
	mBase = m.M
	v29 = m.ExcPending
	if v29 != 0 {
		goto L5
	} else {
		goto L14
	}
L12:
	;
	goto L13
L13:
	;
	v33 = *(*int32)(unsafe.Add(mBase, _c_F_AlterSubscriptionOwner_internal[0]))
	F_check_can_set_role(m, v33, l2)
	mBase = m.M
	v35 = m.ExcPending
	if v35 != 0 {
		goto L5
	} else {
		goto L16
	}
L14:
	;
	if v28 == int32(0) {
		goto L1
	} else {
		goto L15
	}
L15:
	;
	goto L13
L16:
	;
	v38 = *(*int32)(unsafe.Add(mBase, _c_F_AlterSubscriptionOwner_internal[1]))
	v40 = *(*int32)(unsafe.Add(mBase, _c_F_AlterSubscriptionOwner_internal[0]))
	v42 = F_object_aclcheck(m, int32(1262), v38, v40, int64(512))
	mBase = m.M
	v43 = m.ExcPending
	if v43 != 0 {
		goto L5
	} else {
		goto L17
	}
L17:
	;
	if v42 != 0 {
		goto L18
	} else {
		goto L19
	}
L18:
	;
	v46 = *(*int32)(unsafe.Add(mBase, _c_F_AlterSubscriptionOwner_internal[1]))
	v47 = F_get_database_name(m, v46)
	mBase = m.M
	v48 = m.ExcPending
	if v48 != 0 {
		goto L5
	} else {
		goto L21
	}
L19:
	;
	goto L20
L20:
	;
	v51 = *(*int32)(unsafe.Add(mBase, uint32(v8)+104))
	if v51 != 0 {
		goto L23
	} else {
		goto L24
	}
L21:
	;
	F_aclcheck_error(m, v42, int32(9), v47)
	mBase = m.M
	v50 = m.ExcPending
	if v50 != 0 {
		goto L5
	} else {
		goto L22
	}
L22:
	;
	goto L20
L23:
	;
	F_GetUserMappingExtended(m, l2, v51)
	mBase = m.M
	v53 = m.ExcPending
	if v53 != 0 {
		goto L5
	} else {
		goto L26
	}
L24:
	;
	goto L25
L25:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v8)+80)) = l2
	F_CatalogTupleUpdate(m, l0, l1+int32(4), l1)
	mBase = m.M
	v58 = m.ExcPending
	if v58 != 0 {
		goto L5
	} else {
		goto L27
	}
L26:
	;
	goto L25
L27:
	;
	v60 = *(*int32)(unsafe.Add(mBase, uint32(v8)))
	F_changeDependencyOnOwner(m, int32(_a_F_AlterSubscriptionOwner_internal_0), v60, l2)
	mBase = m.M
	v62 = m.ExcPending
	if v62 != 0 {
		goto L5
	} else {
		goto L28
	}
L28:
	;
	v64 = *(*int32)(unsafe.Add(mBase, _c_F_AlterSubscriptionOwner_internal[2]))
	if v64 != 0 {
		goto L29
	} else {
		goto L30
	}
L29:
	;
	v66 = *(*int32)(unsafe.Add(mBase, uint32(v8)))
	v67 = int32(0)
	F_RunObjectPostAlterHook(m, int32(_a_F_AlterSubscriptionOwner_internal_0), v66, v67, v67, v67)
	mBase = m.M
	v71 = m.ExcPending
	if v71 != 0 {
		goto L5
	} else {
		goto L32
	}
L30:
	;
	goto L31
L31:
	;
	v73 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_AlterSubscriptionOwner_internal[3])))
	if v73 == int32(0) {
		goto L34
	} else {
		goto L35
	}
L32:
	;
	goto L31
L33:
	;
	v79 = *(*int32)(unsafe.Add(mBase, uint32(v8)))
	F_LogicalRepWorkersWakeupAtCommit(m, v79)
	mBase = m.M
	v81 = m.ExcPending
	if v81 != 0 {
		goto L5
	} else {
		goto L37
	}
L34:
	;
	v77 = int32(1)
	*(*uint8)(unsafe.Add(mBase, _c_F_AlterSubscriptionOwner_internal[3])) = uint8(v77)
	goto L36
L35:
	;
	goto L36
L36:
	;
	goto L33
L37:
	;
	goto L4
L38:
	;
	F_errcode(m, int32(16797828))
	mBase = m.M
	v89 = m.ExcPending
	if v89 != 0 {
		goto L5
	} else {
		goto L39
	}
L39:
	;
	F_errmsg(m, int32(_a_F_AlterSubscriptionOwner_internal_1), int32(0))
	mBase = m.M
	v93 = m.ExcPending
	if v93 != 0 {
		goto L5
	} else {
		goto L40
	}
L40:
	;
	F_errhint(m, int32(_a_F_AlterSubscriptionOwner_internal_2), int32(0))
	mBase = m.M
	v97 = m.ExcPending
	if v97 != 0 {
		goto L5
	} else {
		goto L41
	}
L41:
	;
	F_errfinish(m, int32(_a_F_AlterSubscriptionOwner_internal_3), int32(2785), int32(_a_F_AlterSubscriptionOwner_internal_4))
	mBase = m.M
	v102 = m.ExcPending
	if v102 != 0 {
		goto L5
	} else {
		goto L42
	}
L42:
	;
	base.Wasm_trap_unreachable()
	for {
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
	var v22 int32
	_ = v22
	var v25 int32
	_ = v25
	var v29 int32
	_ = v29
	var v31 int32
	_ = v31
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
	var v47 int32
	_ = v47
	var v54 int32
	_ = v54
	var v58 int32
	_ = v58
	var v63 int32
	_ = v63
	v7 = m.G0
	v9 = v7 - int32(16)
	m.G0 = v9
	v13 = F_table_open(m, int32(1247), int32(3))
	mBase = m.M
	v14 = m.ExcPending
	if v14 != 0 {
		return
	} else {
		v17 = F_SearchSysCache1(m, int32(82), base.I64_extend_i32_u(l0))
		mBase = m.M
		v18 = m.ExcPending
		if v18 != 0 {
			return
		} else {
			if v17 != 0 {
				v19 = *(*int32)(unsafe.Add(mBase, uint32(v17)+16))
				v20 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v19)+22)))
				v21 = v19 + v20
				v22 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v21)+79)))
				if v22 == int32(99) {
					v25 = *(*int32)(unsafe.Add(mBase, uint32(v21)+84))
					F_ATExecChangeOwner(m, v25, l1, int32(1), int32(8))
					mBase = m.M
					v29 = m.ExcPending
					if v29 != 0 {
						return
					} else {
						F_changeDependencyOnOwner(m, int32(1247), l0, l1)
						mBase = m.M
						v34 = m.ExcPending
						if v34 != 0 {
							return
						} else {
							v36 = *(*int32)(unsafe.Add(mBase, _c_F_AlterTypeOwner_oid[0]))
							if v36 != 0 {
								v38 = int32(0)
								F_RunObjectPostAlterHook(m, int32(1247), l0, v38, v38, v38)
								mBase = m.M
								v42 = m.ExcPending
								if v42 != 0 {
									return
								} else {
									F_ReleaseCatCache(m, v17)
									mBase = m.M
									v44 = m.ExcPending
									if v44 != 0 {
										return
									} else {
										F_relation_close(m, v13, int32(3))
										mBase = m.M
										v47 = m.ExcPending
										if v47 != 0 {
											return
										} else {
											m.G0 = v9 + int32(16)
											return
										}
									}
								}
							} else {
								F_ReleaseCatCache(m, v17)
								mBase = m.M
								v44 = m.ExcPending
								if v44 != 0 {
									return
								} else {
									F_relation_close(m, v13, int32(3))
									mBase = m.M
									v47 = m.ExcPending
									if v47 != 0 {
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
					v31 = m.ExcPending
					if v31 != 0 {
						return
					} else {
						F_changeDependencyOnOwner(m, int32(1247), l0, l1)
						mBase = m.M
						v34 = m.ExcPending
						if v34 != 0 {
							return
						} else {
							v36 = *(*int32)(unsafe.Add(mBase, _c_F_AlterTypeOwner_oid[0]))
							if v36 != 0 {
								v38 = int32(0)
								F_RunObjectPostAlterHook(m, int32(1247), l0, v38, v38, v38)
								mBase = m.M
								v42 = m.ExcPending
								if v42 != 0 {
									return
								} else {
									F_ReleaseCatCache(m, v17)
									mBase = m.M
									v44 = m.ExcPending
									if v44 != 0 {
										return
									} else {
										F_relation_close(m, v13, int32(3))
										mBase = m.M
										v47 = m.ExcPending
										if v47 != 0 {
											return
										} else {
											m.G0 = v9 + int32(16)
											return
										}
									}
								}
							} else {
								F_ReleaseCatCache(m, v17)
								mBase = m.M
								v44 = m.ExcPending
								if v44 != 0 {
									return
								} else {
									F_relation_close(m, v13, int32(3))
									mBase = m.M
									v47 = m.ExcPending
									if v47 != 0 {
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
				v54 = m.ExcPending
				if v54 != 0 {
					return
				} else {
					*(*int32)(unsafe.Add(mBase, uint32(v9))) = l0
					F_errmsg_internal(m, int32(_a_F_AlterTypeOwner_oid_0), v9)
					mBase = m.M
					v58 = m.ExcPending
					if v58 != 0 {
						return
					} else {
						F_errfinish(m, int32(_a_F_AlterTypeOwner_oid_1), int32(4024), int32(_a_F_AlterTypeOwner_oid_2))
						mBase = m.M
						v63 = m.ExcPending
						if v63 != 0 {
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
func F_AnonymousShmemDetach(m *base.Module, l0 int32, l1 int64) {
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
	var v349 int32
	_ = v349
	var v351 int32
	_ = v351
	var v353 int32
	_ = v353
	var v357 int32
	_ = v357
	var v358 int32
	_ = v358
	var v361 int32
	_ = v361
	var v366 int32
	_ = v366
	var v376 int32
	_ = v376
	var v381 int32
	_ = v381
	var v382 int32
	_ = v382
	var v388 int32
	_ = v388
	var v389 int32
	_ = v389
	var v391 int32
	_ = v391
	var v396 int32
	_ = v396
	var v403 int32
	_ = v403
	var v449 int32
	_ = v449
	var v450 int32
	_ = v450
	var v452 int32
	_ = v452
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
	var v463 int32
	_ = v463
	var v465 int32
	_ = v465
	var v469 int32
	_ = v469
	var v470 int32
	_ = v470
	var v471 int32
	_ = v471
	var v472 int32
	_ = v472
	var v476 int32
	_ = v476
	var v480 int32
	_ = v480
	var v484 int32
	_ = v484
	var v485 int32
	_ = v485
	var v486 int32
	_ = v486
	var v487 int32
	_ = v487
	var v491 int32
	_ = v491
	var v492 int32
	_ = v492
	var v493 int32
	_ = v493
	var v495 int32
	_ = v495
	var v509 int32
	_ = v509
	var v510 int32
	_ = v510
	var v512 int32
	_ = v512
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
	var v523 int32
	_ = v523
	var v525 int32
	_ = v525
	var v529 int32
	_ = v529
	var v530 int32
	_ = v530
	var v531 int32
	_ = v531
	var v532 int32
	_ = v532
	var v536 int32
	_ = v536
	var v540 int32
	_ = v540
	var v544 int32
	_ = v544
	var v545 int32
	_ = v545
	var v546 int32
	_ = v546
	var v547 int32
	_ = v547
	var v551 int32
	_ = v551
	var v552 int32
	_ = v552
	var v553 int32
	_ = v553
	var v555 int32
	_ = v555
	var v567 int32
	_ = v567
	var v571 int32
	_ = v571
	var v572 int32
	_ = v572
	var v576 int32
	_ = v576
	var v577 int32
	_ = v577
	var v581 int32
	_ = v581
	var v582 int32
	_ = v582
	var v584 int32
	_ = v584
	var v586 int32
	_ = v586
	var v590 int32
	_ = v590
	var v591 int32
	_ = v591
	var v595 int32
	_ = v595
	var v596 int32
	_ = v596
	var v598 int32
	_ = v598
	var v599 int32
	_ = v599
	var v601 int32
	_ = v601
	var v605 int32
	_ = v605
	var v606 int32
	_ = v606
	var v610 int32
	_ = v610
	var v611 int32
	_ = v611
	var v613 int32
	_ = v613
	var v617 int32
	_ = v617
	var v618 int32
	_ = v618
	var v622 int32
	_ = v622
	var v623 int32
	_ = v623
	var v627 int32
	_ = v627
	var v628 int32
	_ = v628
	var v632 int32
	_ = v632
	var v633 int32
	_ = v633
	var v634 int32
	_ = v634
	var v638 int32
	_ = v638
	var v639 int32
	_ = v639
	var v640 int32
	_ = v640
	var v644 int32
	_ = v644
	var v645 int32
	_ = v645
	var v646 int32
	_ = v646
	var v648 int32
	_ = v648
	var v649 int32
	_ = v649
	var v650 int32
	_ = v650
	var v654 int32
	_ = v654
	var v655 int32
	_ = v655
	var v656 int32
	_ = v656
	var v657 int32
	_ = v657
	var v661 int32
	_ = v661
	var v662 int32
	_ = v662
	var v663 int32
	_ = v663
	var v664 int32
	_ = v664
	var v668 int32
	_ = v668
	var v669 int32
	_ = v669
	var v670 int32
	_ = v670
	var v671 int32
	_ = v671
	var v676 int32
	_ = v676
	var v677 int32
	_ = v677
	var v678 int32
	_ = v678
	var v681 int32
	_ = v681
	var v683 int32
	_ = v683
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
	var v715 int32
	_ = v715
	var v717 int32
	_ = v717
	var v719 int32
	_ = v719
	var v723 int32
	_ = v723
	var v724 int32
	_ = v724
	var v727 int32
	_ = v727
	var v729 int32
	_ = v729
	v8 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
	if v8 == int32(0) {
		v11 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
		v376 = v11
	} else {
		v12 = int32(4)
		v13 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
		v14 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
		if base.Ui32(v14-int32(1021)) < base.Ui32(v12) {
			v23 = v14
			v25 = v12
			v27 = l0 + int32(28)
			for {
				if base.Ui32(int32(1024)) <= base.Ui32(v23) {
					v30 = int32(1024)
					v37 = int32(-1636607408)
					if v13&int32(3) != 0 {
						v83 = v13
						v84 = v30
						v86 = v37
						v87 = v37
						v88 = v37
						for {
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
								continue
							} else {
								break
							}
							break
						}
						switch v129 - int32(1) {
						case 0:
							v302 = v120
							v303 = v121
							v304 = v125
							v305 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v127))))
							v310 = v302 + v305
							v311 = v303
							v312 = v304
						case 1:
							v295 = v120
							v296 = v121
							v297 = v125
							v298 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v127)+1)))
							v302 = v298<<(uint(int32(8))%32) + v295
							v303 = v296
							v304 = v297
							v305 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v127))))
							v310 = v302 + v305
							v311 = v303
							v312 = v304
						case 2:
							v288 = v120
							v289 = v121
							v290 = v125
							v291 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v127)+2)))
							v295 = v291<<(uint(int32(16))%32) + v288
							v296 = v289
							v297 = v290
							v298 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v127)+1)))
							v302 = v298<<(uint(int32(8))%32) + v295
							v303 = v296
							v304 = v297
							v305 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v127))))
							v310 = v302 + v305
							v311 = v303
							v312 = v304
						case 3:
							v282 = v121
							v283 = v125
							v284 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v127)+3)))
							v288 = v284<<(uint(int32(24))%32) + v120
							v289 = v282
							v290 = v283
							v291 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v127)+2)))
							v295 = v291<<(uint(int32(16))%32) + v288
							v296 = v289
							v297 = v290
							v298 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v127)+1)))
							v302 = v298<<(uint(int32(8))%32) + v295
							v303 = v296
							v304 = v297
							v305 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v127))))
							v310 = v302 + v305
							v311 = v303
							v312 = v304
						case 4:
							v278 = v121
							v279 = v125
							v280 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v127)+4)))
							v282 = v278 + v280
							v283 = v279
							v284 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v127)+3)))
							v288 = v284<<(uint(int32(24))%32) + v120
							v289 = v282
							v290 = v283
							v291 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v127)+2)))
							v295 = v291<<(uint(int32(16))%32) + v288
							v296 = v289
							v297 = v290
							v298 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v127)+1)))
							v302 = v298<<(uint(int32(8))%32) + v295
							v303 = v296
							v304 = v297
							v305 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v127))))
							v310 = v302 + v305
							v311 = v303
							v312 = v304
						case 5:
							v272 = v121
							v273 = v125
							v274 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v127)+5)))
							v278 = v274<<(uint(int32(8))%32) + v272
							v279 = v273
							v280 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v127)+4)))
							v282 = v278 + v280
							v283 = v279
							v284 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v127)+3)))
							v288 = v284<<(uint(int32(24))%32) + v120
							v289 = v282
							v290 = v283
							v291 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v127)+2)))
							v295 = v291<<(uint(int32(16))%32) + v288
							v296 = v289
							v297 = v290
							v298 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v127)+1)))
							v302 = v298<<(uint(int32(8))%32) + v295
							v303 = v296
							v304 = v297
							v305 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v127))))
							v310 = v302 + v305
							v311 = v303
							v312 = v304
						case 6:
							v266 = v121
							v267 = v125
							v268 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v127)+6)))
							v272 = v268<<(uint(int32(16))%32) + v266
							v273 = v267
							v274 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v127)+5)))
							v278 = v274<<(uint(int32(8))%32) + v272
							v279 = v273
							v280 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v127)+4)))
							v282 = v278 + v280
							v283 = v279
							v284 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v127)+3)))
							v288 = v284<<(uint(int32(24))%32) + v120
							v289 = v282
							v290 = v283
							v291 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v127)+2)))
							v295 = v291<<(uint(int32(16))%32) + v288
							v296 = v289
							v297 = v290
							v298 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v127)+1)))
							v302 = v298<<(uint(int32(8))%32) + v295
							v303 = v296
							v304 = v297
							v305 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v127))))
							v310 = v302 + v305
							v311 = v303
							v312 = v304
						case 7:
							v261 = v125
							v262 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v127)+7)))
							v266 = v262<<(uint(int32(24))%32) + v121
							v267 = v261
							v268 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v127)+6)))
							v272 = v268<<(uint(int32(16))%32) + v266
							v273 = v267
							v274 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v127)+5)))
							v278 = v274<<(uint(int32(8))%32) + v272
							v279 = v273
							v280 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v127)+4)))
							v282 = v278 + v280
							v283 = v279
							v284 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v127)+3)))
							v288 = v284<<(uint(int32(24))%32) + v120
							v289 = v282
							v290 = v283
							v291 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v127)+2)))
							v295 = v291<<(uint(int32(16))%32) + v288
							v296 = v289
							v297 = v290
							v298 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v127)+1)))
							v302 = v298<<(uint(int32(8))%32) + v295
							v303 = v296
							v304 = v297
							v305 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v127))))
							v310 = v302 + v305
							v311 = v303
							v312 = v304
						case 8:
							v256 = v125
							v257 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v127)+8)))
							v261 = v257<<(uint(int32(8))%32) + v256
							v262 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v127)+7)))
							v266 = v262<<(uint(int32(24))%32) + v121
							v267 = v261
							v268 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v127)+6)))
							v272 = v268<<(uint(int32(16))%32) + v266
							v273 = v267
							v274 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v127)+5)))
							v278 = v274<<(uint(int32(8))%32) + v272
							v279 = v273
							v280 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v127)+4)))
							v282 = v278 + v280
							v283 = v279
							v284 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v127)+3)))
							v288 = v284<<(uint(int32(24))%32) + v120
							v289 = v282
							v290 = v283
							v291 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v127)+2)))
							v295 = v291<<(uint(int32(16))%32) + v288
							v296 = v289
							v297 = v290
							v298 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v127)+1)))
							v302 = v298<<(uint(int32(8))%32) + v295
							v303 = v296
							v304 = v297
							v305 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v127))))
							v310 = v302 + v305
							v311 = v303
							v312 = v304
						case 9:
							v251 = v125
							v252 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v127)+9)))
							v256 = v252<<(uint(int32(16))%32) + v251
							v257 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v127)+8)))
							v261 = v257<<(uint(int32(8))%32) + v256
							v262 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v127)+7)))
							v266 = v262<<(uint(int32(24))%32) + v121
							v267 = v261
							v268 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v127)+6)))
							v272 = v268<<(uint(int32(16))%32) + v266
							v273 = v267
							v274 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v127)+5)))
							v278 = v274<<(uint(int32(8))%32) + v272
							v279 = v273
							v280 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v127)+4)))
							v282 = v278 + v280
							v283 = v279
							v284 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v127)+3)))
							v288 = v284<<(uint(int32(24))%32) + v120
							v289 = v282
							v290 = v283
							v291 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v127)+2)))
							v295 = v291<<(uint(int32(16))%32) + v288
							v296 = v289
							v297 = v290
							v298 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v127)+1)))
							v302 = v298<<(uint(int32(8))%32) + v295
							v303 = v296
							v304 = v297
							v305 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v127))))
							v310 = v302 + v305
							v311 = v303
							v312 = v304
						case 10:
							v247 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v127)+10)))
							v251 = v247<<(uint(int32(24))%32) + v125
							v252 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v127)+9)))
							v256 = v252<<(uint(int32(16))%32) + v251
							v257 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v127)+8)))
							v261 = v257<<(uint(int32(8))%32) + v256
							v262 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v127)+7)))
							v266 = v262<<(uint(int32(24))%32) + v121
							v267 = v261
							v268 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v127)+6)))
							v272 = v268<<(uint(int32(16))%32) + v266
							v273 = v267
							v274 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v127)+5)))
							v278 = v274<<(uint(int32(8))%32) + v272
							v279 = v273
							v280 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v127)+4)))
							v282 = v278 + v280
							v283 = v279
							v284 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v127)+3)))
							v288 = v284<<(uint(int32(24))%32) + v120
							v289 = v282
							v290 = v283
							v291 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v127)+2)))
							v295 = v291<<(uint(int32(16))%32) + v288
							v296 = v289
							v297 = v290
							v298 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v127)+1)))
							v302 = v298<<(uint(int32(8))%32) + v295
							v303 = v296
							v304 = v297
							v305 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v127))))
							v310 = v302 + v305
							v311 = v303
							v312 = v304
						default:
							v310 = v120
							v311 = v121
							v312 = v125
						}
					} else {
						v143 = v13
						v144 = v30
						v146 = v37
						v147 = v37
						v148 = v37
						for {
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
								continue
							} else {
								break
							}
							break
						}
						switch v189 - int32(1) {
						case 0:
							v244 = v180
							v245 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v187))))
							v310 = v244 + v245
							v311 = v181
							v312 = v185
						case 1:
							v239 = v180
							v240 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v187)+1)))
							v244 = v240<<(uint(int32(8))%32) + v239
							v245 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v187))))
							v310 = v244 + v245
							v311 = v181
							v312 = v185
						case 2:
							v235 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v187)+2)))
							v239 = v235<<(uint(int32(16))%32) + v180
							v240 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v187)+1)))
							v244 = v240<<(uint(int32(8))%32) + v239
							v245 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v187))))
							v310 = v244 + v245
							v311 = v181
							v312 = v185
						case 3:
							v232 = v181
							v233 = *(*int32)(unsafe.Add(mBase, uint32(v187)))
							v310 = v233 + v180
							v311 = v232
							v312 = v185
						case 4:
							v229 = v181
							v230 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v187)+4)))
							v232 = v229 + v230
							v233 = *(*int32)(unsafe.Add(mBase, uint32(v187)))
							v310 = v233 + v180
							v311 = v232
							v312 = v185
						case 5:
							v224 = v181
							v225 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v187)+5)))
							v229 = v225<<(uint(int32(8))%32) + v224
							v230 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v187)+4)))
							v232 = v229 + v230
							v233 = *(*int32)(unsafe.Add(mBase, uint32(v187)))
							v310 = v233 + v180
							v311 = v232
							v312 = v185
						case 6:
							v220 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v187)+6)))
							v224 = v220<<(uint(int32(16))%32) + v181
							v225 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v187)+5)))
							v229 = v225<<(uint(int32(8))%32) + v224
							v230 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v187)+4)))
							v232 = v229 + v230
							v233 = *(*int32)(unsafe.Add(mBase, uint32(v187)))
							v310 = v233 + v180
							v311 = v232
							v312 = v185
						case 7:
							v215 = v185
							v216 = *(*int32)(unsafe.Add(mBase, uint32(v187)))
							v218 = *(*int32)(unsafe.Add(mBase, uint32(v187)+4))
							v310 = v216 + v180
							v311 = v218 + v181
							v312 = v215
						case 8:
							v210 = v185
							v211 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v187)+8)))
							v215 = v211<<(uint(int32(8))%32) + v210
							v216 = *(*int32)(unsafe.Add(mBase, uint32(v187)))
							v218 = *(*int32)(unsafe.Add(mBase, uint32(v187)+4))
							v310 = v216 + v180
							v311 = v218 + v181
							v312 = v215
						case 9:
							v205 = v185
							v206 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v187)+9)))
							v210 = v206<<(uint(int32(16))%32) + v205
							v211 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v187)+8)))
							v215 = v211<<(uint(int32(8))%32) + v210
							v216 = *(*int32)(unsafe.Add(mBase, uint32(v187)))
							v218 = *(*int32)(unsafe.Add(mBase, uint32(v187)+4))
							v310 = v216 + v180
							v311 = v218 + v181
							v312 = v215
						case 10:
							v201 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v187)+10)))
							v205 = v201<<(uint(int32(24))%32) + v185
							v206 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v187)+9)))
							v210 = v206<<(uint(int32(16))%32) + v205
							v211 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v187)+8)))
							v215 = v211<<(uint(int32(8))%32) + v210
							v216 = *(*int32)(unsafe.Add(mBase, uint32(v187)))
							v218 = *(*int32)(unsafe.Add(mBase, uint32(v187)+4))
							v310 = v216 + v180
							v311 = v218 + v181
							v312 = v215
						default:
							v310 = v180
							v311 = v181
							v312 = v185
						}
					}
					v315 = int32(14)
					v317 = v311 ^ v312 - base.I32_rotl(v311, v315)
					v321 = v317 ^ v310 - base.I32_rotl(v317, int32(11))
					v325 = v321 ^ v311 - base.I32_rotl(v321, int32(25))
					v329 = v325 ^ v317 - base.I32_rotl(v325, int32(16))
					v333 = v329 ^ v321 - base.I32_rotl(v329, int32(4))
					v337 = v333 ^ v325 - base.I32_rotl(v333, v315)
					*(*int64)(unsafe.Add(mBase, uint32(v13))) = base.I64_extend_i32_u(v337)<<(uint(int64(32))%64) | base.I64_extend_i32_u(v337^v329-base.I32_rotl(v337, int32(24)))
					v349 = int32(8)
				} else {
					v349 = v23
				}
				v351 = int32(1024) - v349
				if base.Ui32(v25) < base.Ui32(v351) {
					v353 = v25
				} else {
					v353 = v351
				}
				if v353 != 0 {
					base.MemoryCopy(m, v349+v13, v27, v353)
				} else {
				}
				v357 = v349 + v353
				v358 = v25 - v353
				if v358 != 0 {
					v23 = v357
					v25 = v358
					v27 = v353 + v27
					continue
				} else {
					break
				}
				break
			}
			v366 = v357
		} else {
			*(*int32)(unsafe.Add(mBase, uint32(v14+v13))) = v8
			v361 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
			v366 = v361 + int32(4)
		}
		*(*int32)(unsafe.Add(mBase, uint32(l0)+28)) = int32(0)
		*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v366
		v376 = v366
	}
	v381 = int32(4)
	v382 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	if base.Ui32(v376-int32(1021)) < base.Ui32(v381) {
		v388 = l1
		v389 = v376
		v391 = v381
		for {
			if base.Ui32(int32(1024)) <= base.Ui32(v389) {
				v396 = int32(1024)
				v403 = int32(-1636607408)
				if v382&int32(3) != 0 {
					v449 = v382
					v450 = v396
					v452 = v403
					v453 = v403
					v454 = v403
					for {
						v456 = *(*int32)(unsafe.Add(mBase, uint32(v449)+4))
						v457 = v456 + v453
						v458 = *(*int32)(unsafe.Add(mBase, uint32(v449)))
						v460 = *(*int32)(unsafe.Add(mBase, uint32(v449)+8))
						v461 = v460 + v454
						v463 = int32(4)
						v465 = v458 + v452 - v461 ^ base.I32_rotl(v461, v463)
						v469 = v457 - v465 ^ base.I32_rotl(v465, int32(6))
						v470 = v461 + v457
						v471 = v465 + v470
						v472 = v469 + v471
						v476 = v470 - v469 ^ base.I32_rotl(v469, int32(8))
						v480 = v471 - v476 ^ base.I32_rotl(v476, int32(16))
						v484 = v472 - v480 ^ base.I32_rotl(v480, int32(19))
						v485 = v476 + v472
						v486 = v480 + v485
						v487 = v484 + v486
						v491 = v485 - v484 ^ base.I32_rotl(v484, v463)
						v492 = int32(12)
						v493 = v449 + v492
						v495 = v450 - v492
						if base.Ui32(int32(11)) < base.Ui32(v495) {
							v449 = v493
							v450 = v495
							v452 = v486
							v453 = v487
							v454 = v491
							continue
						} else {
							break
						}
						break
					}
					switch v495 - int32(1) {
					case 0:
						v668 = v486
						v669 = v487
						v670 = v491
						v671 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v493))))
						v676 = v668 + v671
						v677 = v669
						v678 = v670
					case 1:
						v661 = v486
						v662 = v487
						v663 = v491
						v664 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v493)+1)))
						v668 = v664<<(uint(int32(8))%32) + v661
						v669 = v662
						v670 = v663
						v671 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v493))))
						v676 = v668 + v671
						v677 = v669
						v678 = v670
					case 2:
						v654 = v486
						v655 = v487
						v656 = v491
						v657 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v493)+2)))
						v661 = v657<<(uint(int32(16))%32) + v654
						v662 = v655
						v663 = v656
						v664 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v493)+1)))
						v668 = v664<<(uint(int32(8))%32) + v661
						v669 = v662
						v670 = v663
						v671 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v493))))
						v676 = v668 + v671
						v677 = v669
						v678 = v670
					case 3:
						v648 = v487
						v649 = v491
						v650 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v493)+3)))
						v654 = v650<<(uint(int32(24))%32) + v486
						v655 = v648
						v656 = v649
						v657 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v493)+2)))
						v661 = v657<<(uint(int32(16))%32) + v654
						v662 = v655
						v663 = v656
						v664 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v493)+1)))
						v668 = v664<<(uint(int32(8))%32) + v661
						v669 = v662
						v670 = v663
						v671 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v493))))
						v676 = v668 + v671
						v677 = v669
						v678 = v670
					case 4:
						v644 = v487
						v645 = v491
						v646 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v493)+4)))
						v648 = v644 + v646
						v649 = v645
						v650 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v493)+3)))
						v654 = v650<<(uint(int32(24))%32) + v486
						v655 = v648
						v656 = v649
						v657 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v493)+2)))
						v661 = v657<<(uint(int32(16))%32) + v654
						v662 = v655
						v663 = v656
						v664 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v493)+1)))
						v668 = v664<<(uint(int32(8))%32) + v661
						v669 = v662
						v670 = v663
						v671 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v493))))
						v676 = v668 + v671
						v677 = v669
						v678 = v670
					case 5:
						v638 = v487
						v639 = v491
						v640 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v493)+5)))
						v644 = v640<<(uint(int32(8))%32) + v638
						v645 = v639
						v646 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v493)+4)))
						v648 = v644 + v646
						v649 = v645
						v650 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v493)+3)))
						v654 = v650<<(uint(int32(24))%32) + v486
						v655 = v648
						v656 = v649
						v657 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v493)+2)))
						v661 = v657<<(uint(int32(16))%32) + v654
						v662 = v655
						v663 = v656
						v664 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v493)+1)))
						v668 = v664<<(uint(int32(8))%32) + v661
						v669 = v662
						v670 = v663
						v671 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v493))))
						v676 = v668 + v671
						v677 = v669
						v678 = v670
					case 6:
						v632 = v487
						v633 = v491
						v634 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v493)+6)))
						v638 = v634<<(uint(int32(16))%32) + v632
						v639 = v633
						v640 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v493)+5)))
						v644 = v640<<(uint(int32(8))%32) + v638
						v645 = v639
						v646 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v493)+4)))
						v648 = v644 + v646
						v649 = v645
						v650 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v493)+3)))
						v654 = v650<<(uint(int32(24))%32) + v486
						v655 = v648
						v656 = v649
						v657 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v493)+2)))
						v661 = v657<<(uint(int32(16))%32) + v654
						v662 = v655
						v663 = v656
						v664 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v493)+1)))
						v668 = v664<<(uint(int32(8))%32) + v661
						v669 = v662
						v670 = v663
						v671 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v493))))
						v676 = v668 + v671
						v677 = v669
						v678 = v670
					case 7:
						v627 = v491
						v628 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v493)+7)))
						v632 = v628<<(uint(int32(24))%32) + v487
						v633 = v627
						v634 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v493)+6)))
						v638 = v634<<(uint(int32(16))%32) + v632
						v639 = v633
						v640 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v493)+5)))
						v644 = v640<<(uint(int32(8))%32) + v638
						v645 = v639
						v646 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v493)+4)))
						v648 = v644 + v646
						v649 = v645
						v650 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v493)+3)))
						v654 = v650<<(uint(int32(24))%32) + v486
						v655 = v648
						v656 = v649
						v657 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v493)+2)))
						v661 = v657<<(uint(int32(16))%32) + v654
						v662 = v655
						v663 = v656
						v664 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v493)+1)))
						v668 = v664<<(uint(int32(8))%32) + v661
						v669 = v662
						v670 = v663
						v671 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v493))))
						v676 = v668 + v671
						v677 = v669
						v678 = v670
					case 8:
						v622 = v491
						v623 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v493)+8)))
						v627 = v623<<(uint(int32(8))%32) + v622
						v628 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v493)+7)))
						v632 = v628<<(uint(int32(24))%32) + v487
						v633 = v627
						v634 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v493)+6)))
						v638 = v634<<(uint(int32(16))%32) + v632
						v639 = v633
						v640 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v493)+5)))
						v644 = v640<<(uint(int32(8))%32) + v638
						v645 = v639
						v646 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v493)+4)))
						v648 = v644 + v646
						v649 = v645
						v650 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v493)+3)))
						v654 = v650<<(uint(int32(24))%32) + v486
						v655 = v648
						v656 = v649
						v657 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v493)+2)))
						v661 = v657<<(uint(int32(16))%32) + v654
						v662 = v655
						v663 = v656
						v664 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v493)+1)))
						v668 = v664<<(uint(int32(8))%32) + v661
						v669 = v662
						v670 = v663
						v671 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v493))))
						v676 = v668 + v671
						v677 = v669
						v678 = v670
					case 9:
						v617 = v491
						v618 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v493)+9)))
						v622 = v618<<(uint(int32(16))%32) + v617
						v623 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v493)+8)))
						v627 = v623<<(uint(int32(8))%32) + v622
						v628 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v493)+7)))
						v632 = v628<<(uint(int32(24))%32) + v487
						v633 = v627
						v634 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v493)+6)))
						v638 = v634<<(uint(int32(16))%32) + v632
						v639 = v633
						v640 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v493)+5)))
						v644 = v640<<(uint(int32(8))%32) + v638
						v645 = v639
						v646 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v493)+4)))
						v648 = v644 + v646
						v649 = v645
						v650 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v493)+3)))
						v654 = v650<<(uint(int32(24))%32) + v486
						v655 = v648
						v656 = v649
						v657 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v493)+2)))
						v661 = v657<<(uint(int32(16))%32) + v654
						v662 = v655
						v663 = v656
						v664 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v493)+1)))
						v668 = v664<<(uint(int32(8))%32) + v661
						v669 = v662
						v670 = v663
						v671 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v493))))
						v676 = v668 + v671
						v677 = v669
						v678 = v670
					case 10:
						v613 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v493)+10)))
						v617 = v613<<(uint(int32(24))%32) + v491
						v618 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v493)+9)))
						v622 = v618<<(uint(int32(16))%32) + v617
						v623 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v493)+8)))
						v627 = v623<<(uint(int32(8))%32) + v622
						v628 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v493)+7)))
						v632 = v628<<(uint(int32(24))%32) + v487
						v633 = v627
						v634 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v493)+6)))
						v638 = v634<<(uint(int32(16))%32) + v632
						v639 = v633
						v640 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v493)+5)))
						v644 = v640<<(uint(int32(8))%32) + v638
						v645 = v639
						v646 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v493)+4)))
						v648 = v644 + v646
						v649 = v645
						v650 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v493)+3)))
						v654 = v650<<(uint(int32(24))%32) + v486
						v655 = v648
						v656 = v649
						v657 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v493)+2)))
						v661 = v657<<(uint(int32(16))%32) + v654
						v662 = v655
						v663 = v656
						v664 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v493)+1)))
						v668 = v664<<(uint(int32(8))%32) + v661
						v669 = v662
						v670 = v663
						v671 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v493))))
						v676 = v668 + v671
						v677 = v669
						v678 = v670
					default:
						v676 = v486
						v677 = v487
						v678 = v491
					}
				} else {
					v509 = v382
					v510 = v396
					v512 = v403
					v513 = v403
					v514 = v403
					for {
						v516 = *(*int32)(unsafe.Add(mBase, uint32(v509)+4))
						v517 = v516 + v513
						v518 = *(*int32)(unsafe.Add(mBase, uint32(v509)))
						v520 = *(*int32)(unsafe.Add(mBase, uint32(v509)+8))
						v521 = v520 + v514
						v523 = int32(4)
						v525 = v518 + v512 - v521 ^ base.I32_rotl(v521, v523)
						v529 = v517 - v525 ^ base.I32_rotl(v525, int32(6))
						v530 = v521 + v517
						v531 = v525 + v530
						v532 = v529 + v531
						v536 = v530 - v529 ^ base.I32_rotl(v529, int32(8))
						v540 = v531 - v536 ^ base.I32_rotl(v536, int32(16))
						v544 = v532 - v540 ^ base.I32_rotl(v540, int32(19))
						v545 = v536 + v532
						v546 = v540 + v545
						v547 = v544 + v546
						v551 = v545 - v544 ^ base.I32_rotl(v544, v523)
						v552 = int32(12)
						v553 = v509 + v552
						v555 = v510 - v552
						if base.Ui32(int32(11)) < base.Ui32(v555) {
							v509 = v553
							v510 = v555
							v512 = v546
							v513 = v547
							v514 = v551
							continue
						} else {
							break
						}
						break
					}
					switch v555 - int32(1) {
					case 0:
						v610 = v546
						v611 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v553))))
						v676 = v610 + v611
						v677 = v547
						v678 = v551
					case 1:
						v605 = v546
						v606 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v553)+1)))
						v610 = v606<<(uint(int32(8))%32) + v605
						v611 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v553))))
						v676 = v610 + v611
						v677 = v547
						v678 = v551
					case 2:
						v601 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v553)+2)))
						v605 = v601<<(uint(int32(16))%32) + v546
						v606 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v553)+1)))
						v610 = v606<<(uint(int32(8))%32) + v605
						v611 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v553))))
						v676 = v610 + v611
						v677 = v547
						v678 = v551
					case 3:
						v598 = v547
						v599 = *(*int32)(unsafe.Add(mBase, uint32(v553)))
						v676 = v599 + v546
						v677 = v598
						v678 = v551
					case 4:
						v595 = v547
						v596 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v553)+4)))
						v598 = v595 + v596
						v599 = *(*int32)(unsafe.Add(mBase, uint32(v553)))
						v676 = v599 + v546
						v677 = v598
						v678 = v551
					case 5:
						v590 = v547
						v591 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v553)+5)))
						v595 = v591<<(uint(int32(8))%32) + v590
						v596 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v553)+4)))
						v598 = v595 + v596
						v599 = *(*int32)(unsafe.Add(mBase, uint32(v553)))
						v676 = v599 + v546
						v677 = v598
						v678 = v551
					case 6:
						v586 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v553)+6)))
						v590 = v586<<(uint(int32(16))%32) + v547
						v591 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v553)+5)))
						v595 = v591<<(uint(int32(8))%32) + v590
						v596 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v553)+4)))
						v598 = v595 + v596
						v599 = *(*int32)(unsafe.Add(mBase, uint32(v553)))
						v676 = v599 + v546
						v677 = v598
						v678 = v551
					case 7:
						v581 = v551
						v582 = *(*int32)(unsafe.Add(mBase, uint32(v553)))
						v584 = *(*int32)(unsafe.Add(mBase, uint32(v553)+4))
						v676 = v582 + v546
						v677 = v584 + v547
						v678 = v581
					case 8:
						v576 = v551
						v577 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v553)+8)))
						v581 = v577<<(uint(int32(8))%32) + v576
						v582 = *(*int32)(unsafe.Add(mBase, uint32(v553)))
						v584 = *(*int32)(unsafe.Add(mBase, uint32(v553)+4))
						v676 = v582 + v546
						v677 = v584 + v547
						v678 = v581
					case 9:
						v571 = v551
						v572 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v553)+9)))
						v576 = v572<<(uint(int32(16))%32) + v571
						v577 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v553)+8)))
						v581 = v577<<(uint(int32(8))%32) + v576
						v582 = *(*int32)(unsafe.Add(mBase, uint32(v553)))
						v584 = *(*int32)(unsafe.Add(mBase, uint32(v553)+4))
						v676 = v582 + v546
						v677 = v584 + v547
						v678 = v581
					case 10:
						v567 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v553)+10)))
						v571 = v567<<(uint(int32(24))%32) + v551
						v572 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v553)+9)))
						v576 = v572<<(uint(int32(16))%32) + v571
						v577 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v553)+8)))
						v581 = v577<<(uint(int32(8))%32) + v576
						v582 = *(*int32)(unsafe.Add(mBase, uint32(v553)))
						v584 = *(*int32)(unsafe.Add(mBase, uint32(v553)+4))
						v676 = v582 + v546
						v677 = v584 + v547
						v678 = v581
					default:
						v676 = v546
						v677 = v547
						v678 = v551
					}
				}
				v681 = int32(14)
				v683 = v677 ^ v678 - base.I32_rotl(v677, v681)
				v687 = v683 ^ v676 - base.I32_rotl(v683, int32(11))
				v691 = v687 ^ v677 - base.I32_rotl(v687, int32(25))
				v695 = v691 ^ v683 - base.I32_rotl(v691, int32(16))
				v699 = v695 ^ v687 - base.I32_rotl(v695, int32(4))
				v703 = v699 ^ v691 - base.I32_rotl(v699, v681)
				*(*int64)(unsafe.Add(mBase, uint32(v382))) = base.I64_extend_i32_u(v703)<<(uint(int64(32))%64) | base.I64_extend_i32_u(v703^v695-base.I32_rotl(v703, int32(24)))
				v715 = int32(8)
			} else {
				v715 = v389
			}
			v717 = int32(1024) - v715
			if base.Ui32(v391) < base.Ui32(v717) {
				v719 = v391
			} else {
				v719 = v717
			}
			if v719 != 0 {
				base.MemoryCopy(m, v715+v382, v388, v719)
			} else {
			}
			v723 = v715 + v719
			v724 = v391 - v719
			if v724 != 0 {
				v388 = v388 + v719
				v389 = v723
				v391 = v724
				continue
			} else {
				break
			}
			break
		}
		*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v723
		return
	} else {
		v727 = *(*int32)(unsafe.Add(mBase, uint32(l1)))
		*(*int32)(unsafe.Add(mBase, uint32(v376+v382))) = v727
		v729 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
		*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v729 + int32(4)
		return
	}
}
func F_ApplyLauncherShmemInit(m *base.Module, l0 int32) {
	mBase := m.M
	_ = mBase
	var v2 int32
	_ = v2
	var v5 int32
	_ = v5
	var v9 int32
	_ = v9
	var v12 int32
	_ = v12
	var v15 int32
	_ = v15
	var v18 int32
	_ = v18
	var v21 int32
	_ = v21
	var v28 int32
	_ = v28
	var v30 int32
	_ = v30
	v2 = int32(0)
	v5 = *(*int32)(unsafe.Add(mBase, _c_F_ApplyLauncherShmemInit[0]))
	*(*int64)(unsafe.Add(mBase, uint32(v5)+4)) = int64(0)
	v9 = *(*int32)(unsafe.Add(mBase, _c_F_ApplyLauncherShmemInit[1]))
	if v2 < v9 {
		v12 = v2
		for {
			v15 = *(*int32)(unsafe.Add(mBase, _c_F_ApplyLauncherShmemInit[0]))
			v18 = v15 + v12<<(uint(int32(7))%32)
			v21 = int32(0)
			base.MemoryFill(m, v18+int32(16), v21, int32(128))
			atomic.StoreUint32((*uint32)(unsafe.Add(mBase, uint32(v18)+72)), uint32(v21))
			v28 = v12 + int32(1)
			v30 = *(*int32)(unsafe.Add(mBase, _c_F_ApplyLauncherShmemInit[1]))
			if v28 < v30 {
				v12 = v28
				continue
			} else {
				break
			}
			break
		}
	} else {
	}
	return
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
	var v31 int32
	_ = v31
	var v32 int32
	_ = v32
	var v37 int32
	_ = v37
	var v39 int32
	_ = v39
	var v42 int32
	_ = v42
	var v71 int32
	_ = v71
	var v79 int32
	_ = v79
	var v81 int32
	_ = v81
	var v84 int32
	_ = v84
	var v113 int32
	_ = v113
	var v121 int32
	_ = v121
	var v123 int32
	_ = v123
	var v126 int32
	_ = v126
	var v155 int32
	_ = v155
	var v159 int32
	_ = v159
	var v169 int32
	_ = v169
	var v171 int32
	_ = v171
	var v172 int32
	_ = v172
	var v173 int32
	_ = v173
	var v176 int64
	_ = v176
	var v187 int32
	_ = v187
	var v189 int32
	_ = v189
	var v204 int32
	_ = v204
	var v206 int32
	_ = v206
	var v223 int32
	_ = v223
	var v225 int32
	_ = v225
	var v226 int32
	_ = v226
	var v229 int64
	_ = v229
	var v242 int32
	_ = v242
	var v249 int32
	_ = v249
	var v251 int32
	_ = v251
	var v253 int32
	_ = v253
	var v285 int32
	_ = v285
	var v293 int32
	_ = v293
	var v295 int32
	_ = v295
	var v298 int32
	_ = v298
	var v327 int32
	_ = v327
	var v333 int32
	_ = v333
	var v335 int32
	_ = v335
	var v337 int32
	_ = v337
	var v369 int32
	_ = v369
	var v377 int32
	_ = v377
	var v379 int32
	_ = v379
	var v382 int32
	_ = v382
	var v411 int32
	_ = v411
	var v419 int32
	_ = v419
	var v421 int32
	_ = v421
	var v423 int32
	_ = v423
	var v446 int32
	_ = v446
	var v453 int32
	_ = v453
	var v458 int32
	_ = v458
	var v460 int32
	_ = v460
	var v462 int32
	_ = v462
	var v470 int32
	_ = v470
	var v471 int32
	_ = v471
	var v473 int32
	_ = v473
	var v481 int32
	_ = v481
	var v484 int32
	_ = v484
	var v492 int32
	_ = v492
	var v498 int32
	_ = v498
	var v504 int32
	_ = v504
	var v510 int32
	_ = v510
	var v516 int32
	_ = v516
	var v522 int32
	_ = v522
	var v528 int32
	_ = v528
	var v534 int32
	_ = v534
	var v536 int32
	_ = v536
	var v544 int32
	_ = v544
	var v550 int32
	_ = v550
	var v552 int32
	_ = v552
	var v556 int32
	_ = v556
	var v557 int32
	_ = v557
	var v559 int32
	_ = v559
	var v560 int32
	_ = v560
	var v564 int32
	_ = v564
	var v567 int32
	_ = v567
	var v568 int32
	_ = v568
	var v569 int32
	_ = v569
	var v573 int32
	_ = v573
	var v579 int32
	_ = v579
	var v583 int32
	_ = v583
	var v587 int32
	_ = v587
	var v591 int32
	_ = v591
	var v593 int32
	_ = v593
	var v594 int32
	_ = v594
	var v596 int32
	_ = v596
	var v602 int32
	_ = v602
	var v603 int32
	_ = v603
	var v604 int32
	_ = v604
	var v608 int32
	_ = v608
	var v609 int32
	_ = v609
	var v610 int32
	_ = v610
	var v613 int64
	_ = v613
	var v614 int64
	_ = v614
	var v625 int32
	_ = v625
	var v626 int32
	_ = v626
	var v631 int32
	_ = v631
	var v633 int32
	_ = v633
	var v637 int32
	_ = v637
	var v640 int32
	_ = v640
	var v641 int32
	_ = v641
	var v645 int32
	_ = v645
	var v650 int32
	_ = v650
	var v652 int32
	_ = v652
	var v657 int64
	_ = v657
	var v658 int32
	_ = v658
	var v661 int32
	_ = v661
	var v662 int32
	_ = v662
	var v665 int32
	_ = v665
	var v668 int32
	_ = v668
	var v669 int32
	_ = v669
	var v673 int32
	_ = v673
	var v678 int32
	_ = v678
	var v680 int32
	_ = v680
	var v684 int32
	_ = v684
	var v690 int32
	_ = v690
	var v696 int32
	_ = v696
	var v697 int64
	_ = v697
	var v701 int32
	_ = v701
	var v703 int32
	_ = v703
	var v704 int32
	_ = v704
	var v707 int32
	_ = v707
	var v709 int32
	_ = v709
	var v711 int32
	_ = v711
	var v715 int32
	_ = v715
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
	v696 = int32(m.ExcTag)
	v697 = int64(m.ExcVals[0])
	m.ExcPending = 0
	if v696 == int32(0) {
		goto L166
	} else {
		goto L167
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
	v470 = v14
	goto L9
L9:
	;
	if v470 != 0 {
		goto L110
	} else {
		goto L111
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
	goto L15
L13:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_AutoVacWorkerMain[0])) = int32(0)
	goto L12
L14:
	;
	v37 = m.G0
	v39 = v37 - int32(32)
	m.G0 = v39
	v42 = int32(967)
	switch v42 {
	case 0, 2:
		goto L19
	default:
		goto L20
	}
L15:
	;
	v31 = *(*int32)(unsafe.Add(mBase, _c_F_AutoVacWorkerMain[1]))
	v32 = F_GetBackendTypeDesc(m, v31)
	mBase = m.M
	goto L17
L17:
	;
	goto L14
L18:
	;
	v79 = m.G0
	v81 = v79 - int32(32)
	m.G0 = v81
	v84 = int32(968)
	switch v84 {
	case 0, 2:
		goto L29
	default:
		goto L30
	}
L19:
	;
	F_sigemptyset(m, v39+int32(16))
	mBase = m.M
	*(*int32)(unsafe.Add(mBase, uint32(v39)+24)) = int32(268435456)
	switch v42 {
	case 0:
		goto L24
	default:
		goto L22
	case 2:
		goto L23
	}
L20:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_AutoVacWorkerMain[2])) = int32(965)
	goto L19
L21:
	;
	goto L26
L22:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v39)+24)) = int32(268435460)
	*(*int32)(unsafe.Add(mBase, uint32(v39)+12)) = int32(_a_F_AutoVacWorkerMain_0)
	goto L21
L23:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v39)+12)) = int32(0)
	goto L21
L24:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v39)+12)) = int32(-2)
	goto L21
L26:
	;
	goto L27
L27:
	;
	v71 = F___sigaction(m, int32(1), v39+int32(12), int32(0))
	mBase = m.M
	m.G0 = v39 + int32(32)
	goto L18
L28:
	;
	v121 = m.G0
	v123 = v121 - int32(32)
	m.G0 = v123
	v126 = int32(974)
	switch v126 {
	case 0, 2:
		goto L39
	default:
		goto L40
	}
L29:
	;
	F_sigemptyset(m, v81+int32(16))
	mBase = m.M
	*(*int32)(unsafe.Add(mBase, uint32(v81)+24)) = int32(268435456)
	switch v84 {
	case 0:
		goto L34
	default:
		goto L32
	case 2:
		goto L33
	}
L30:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_AutoVacWorkerMain[3])) = int32(966)
	goto L29
L31:
	;
	goto L36
L32:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v81)+24)) = int32(268435460)
	*(*int32)(unsafe.Add(mBase, uint32(v81)+12)) = int32(_a_F_AutoVacWorkerMain_0)
	goto L31
L33:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v81)+12)) = int32(0)
	goto L31
L34:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v81)+12)) = int32(-2)
	goto L31
L36:
	;
	goto L37
L37:
	;
	v113 = F___sigaction(m, int32(2), v81+int32(12), int32(0))
	mBase = m.M
	m.G0 = v81 + int32(32)
	goto L28
L38:
	;
	v159 = int32(0)
	*(*int32)(unsafe.Add(mBase, _c_F_AutoVacWorkerMain[4])) = v159
	*(*int32)(unsafe.Add(mBase, _c_F_AutoVacWorkerMain[5])) = v159
	v169 = v159
	goto L49
L39:
	;
	F_sigemptyset(m, v123+int32(16))
	mBase = m.M
	*(*int32)(unsafe.Add(mBase, uint32(v123)+24)) = int32(268435456)
	switch v126 {
	case 0:
		goto L44
	default:
		goto L42
	case 2:
		goto L43
	}
L40:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_AutoVacWorkerMain[6])) = int32(972)
	goto L39
L41:
	;
	goto L46
L42:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v123)+24)) = int32(268435460)
	*(*int32)(unsafe.Add(mBase, uint32(v123)+12)) = int32(_a_F_AutoVacWorkerMain_0)
	goto L41
L43:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v123)+12)) = int32(0)
	goto L41
L44:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v123)+12)) = int32(-2)
	goto L41
L46:
	;
	goto L47
L47:
	;
	v155 = F___sigaction(m, int32(15), v123+int32(12), int32(0))
	mBase = m.M
	m.G0 = v123 + int32(32)
	goto L38
L48:
	;
	v249 = int32(0)
	v251 = m.G0
	v253 = v251 - int32(32)
	m.G0 = v253
	switch v249 {
	case 0, 2:
		goto L55
	default:
		goto L56
	}
L49:
	;
	v171 = int32(40)
	v172 = v169 * v171
	v173 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v172)+uint32(_c_F_AutoVacWorkerMain[7]))) = uint8(v173)
	*(*int32)(unsafe.Add(mBase, uint32(v172)+uint32(_c_F_AutoVacWorkerMain[8]))) = v169
	v176 = int64(0)
	*(*int64)(unsafe.Add(mBase, uint32(v172)+uint32(_c_F_AutoVacWorkerMain[9]))) = v176
	*(*int32)(unsafe.Add(mBase, uint32(v172)+uint32(_c_F_AutoVacWorkerMain[10]))) = v173
	*(*int64)(unsafe.Add(mBase, uint32(v172)+uint32(_c_F_AutoVacWorkerMain[11]))) = v176
	*(*int32)(unsafe.Add(mBase, uint32(v172)+uint32(_c_F_AutoVacWorkerMain[12]))) = v173
	*(*uint8)(unsafe.Add(mBase, uint32(v172)+uint32(_c_F_AutoVacWorkerMain[13]))) = uint8(v173)
	v187 = v169 | int32(1)
	v189 = v187 * v171
	*(*uint8)(unsafe.Add(mBase, uint32(v189)+uint32(_c_F_AutoVacWorkerMain[7]))) = uint8(v173)
	*(*int32)(unsafe.Add(mBase, uint32(v189)+uint32(_c_F_AutoVacWorkerMain[8]))) = v187
	*(*int64)(unsafe.Add(mBase, uint32(v189)+uint32(_c_F_AutoVacWorkerMain[9]))) = v176
	*(*int32)(unsafe.Add(mBase, uint32(v189)+uint32(_c_F_AutoVacWorkerMain[10]))) = v173
	*(*int64)(unsafe.Add(mBase, uint32(v189)+uint32(_c_F_AutoVacWorkerMain[11]))) = v176
	*(*int32)(unsafe.Add(mBase, uint32(v189)+uint32(_c_F_AutoVacWorkerMain[12]))) = v173
	*(*uint8)(unsafe.Add(mBase, uint32(v189)+uint32(_c_F_AutoVacWorkerMain[13]))) = uint8(v173)
	v204 = v169 | int32(2)
	v206 = v204 * v171
	*(*uint8)(unsafe.Add(mBase, uint32(v206)+uint32(_c_F_AutoVacWorkerMain[7]))) = uint8(v173)
	*(*int32)(unsafe.Add(mBase, uint32(v206)+uint32(_c_F_AutoVacWorkerMain[8]))) = v204
	*(*int64)(unsafe.Add(mBase, uint32(v206)+uint32(_c_F_AutoVacWorkerMain[9]))) = v176
	*(*int32)(unsafe.Add(mBase, uint32(v206)+uint32(_c_F_AutoVacWorkerMain[10]))) = v173
	*(*int64)(unsafe.Add(mBase, uint32(v206)+uint32(_c_F_AutoVacWorkerMain[11]))) = v176
	*(*int32)(unsafe.Add(mBase, uint32(v206)+uint32(_c_F_AutoVacWorkerMain[12]))) = v173
	*(*uint8)(unsafe.Add(mBase, uint32(v206)+uint32(_c_F_AutoVacWorkerMain[13]))) = uint8(v173)
	if v169 != int32(20) {
		goto L51
	} else {
		goto L52
	}
L50:
	;
	v242 = int32(1)
	*(*uint8)(unsafe.Add(mBase, _c_F_AutoVacWorkerMain[14])) = uint8(v242)
	F_pqsignal_be(m, int32(14), int32(1992))
	mBase = m.M
	goto L48
L51:
	;
	v223 = v169 | int32(3)
	v225 = v223 * int32(40)
	v226 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v225)+uint32(_c_F_AutoVacWorkerMain[7]))) = uint8(v226)
	*(*int32)(unsafe.Add(mBase, uint32(v225)+uint32(_c_F_AutoVacWorkerMain[8]))) = v223
	v229 = int64(0)
	*(*int64)(unsafe.Add(mBase, uint32(v225)+uint32(_c_F_AutoVacWorkerMain[9]))) = v229
	*(*int32)(unsafe.Add(mBase, uint32(v225)+uint32(_c_F_AutoVacWorkerMain[10]))) = v226
	*(*int64)(unsafe.Add(mBase, uint32(v225)+uint32(_c_F_AutoVacWorkerMain[11]))) = v229
	*(*int32)(unsafe.Add(mBase, uint32(v225)+uint32(_c_F_AutoVacWorkerMain[12]))) = v226
	*(*uint8)(unsafe.Add(mBase, uint32(v225)+uint32(_c_F_AutoVacWorkerMain[13]))) = uint8(v226)
	v169 = v169 + int32(4)
	goto L49
L52:
	;
	goto L53
L53:
	;
	goto L50
L54:
	;
	v293 = m.G0
	v295 = v293 - int32(32)
	m.G0 = v295
	v298 = int32(970)
	switch v298 {
	case 0, 2:
		goto L65
	default:
		goto L66
	}
L55:
	;
	F_sigemptyset(m, v253+int32(16))
	mBase = m.M
	*(*int32)(unsafe.Add(mBase, uint32(v253)+24)) = int32(268435456)
	switch v249 {
	case 0:
		goto L60
	default:
		goto L58
	case 2:
		goto L59
	}
L56:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_AutoVacWorkerMain[15])) = int32(-2)
	goto L55
L57:
	;
	goto L62
L58:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v253)+24)) = int32(268435460)
	*(*int32)(unsafe.Add(mBase, uint32(v253)+12)) = int32(_a_F_AutoVacWorkerMain_0)
	goto L57
L59:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v253)+12)) = int32(0)
	goto L57
L60:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v253)+12)) = int32(-2)
	goto L57
L62:
	;
	goto L63
L63:
	;
	v285 = F___sigaction(m, int32(13), v253+int32(12), int32(0))
	mBase = m.M
	m.G0 = v253 + int32(32)
	goto L54
L64:
	;
	v333 = int32(0)
	v335 = m.G0
	v337 = v335 - int32(32)
	m.G0 = v337
	switch v333 {
	case 0, 2:
		goto L75
	default:
		goto L76
	}
L65:
	;
	F_sigemptyset(m, v295+int32(16))
	mBase = m.M
	*(*int32)(unsafe.Add(mBase, uint32(v295)+24)) = int32(268435456)
	switch v298 {
	case 0:
		goto L70
	default:
		goto L68
	case 2:
		goto L69
	}
L66:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_AutoVacWorkerMain[16])) = int32(968)
	goto L65
L67:
	;
	goto L72
L68:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v295)+24)) = int32(268435460)
	*(*int32)(unsafe.Add(mBase, uint32(v295)+12)) = int32(_a_F_AutoVacWorkerMain_0)
	goto L67
L69:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v295)+12)) = int32(0)
	goto L67
L70:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v295)+12)) = int32(-2)
	goto L67
L72:
	;
	goto L73
L73:
	;
	v327 = F___sigaction(m, int32(10), v295+int32(12), int32(0))
	mBase = m.M
	m.G0 = v295 + int32(32)
	goto L64
L74:
	;
	v377 = m.G0
	v379 = v377 - int32(32)
	m.G0 = v379
	v382 = int32(972)
	switch v382 {
	case 0, 2:
		goto L85
	default:
		goto L86
	}
L75:
	;
	F_sigemptyset(m, v337+int32(16))
	mBase = m.M
	*(*int32)(unsafe.Add(mBase, uint32(v337)+24)) = int32(268435456)
	switch v333 {
	case 0:
		goto L80
	default:
		goto L78
	case 2:
		goto L79
	}
L76:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_AutoVacWorkerMain[17])) = int32(-2)
	goto L75
L77:
	;
	goto L82
L78:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v337)+24)) = int32(268435460)
	*(*int32)(unsafe.Add(mBase, uint32(v337)+12)) = int32(_a_F_AutoVacWorkerMain_0)
	goto L77
L79:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v337)+12)) = int32(0)
	goto L77
L80:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v337)+12)) = int32(-2)
	goto L77
L82:
	;
	goto L83
L83:
	;
	v369 = F___sigaction(m, int32(12), v337+int32(12), int32(0))
	mBase = m.M
	m.G0 = v337 + int32(32)
	goto L74
L84:
	;
	v419 = m.G0
	v421 = v419 - int32(32)
	m.G0 = v421
	v423 = int32(2)
	switch v423 {
	case 0, 2:
		goto L95
	default:
		goto L96
	}
L85:
	;
	F_sigemptyset(m, v379+int32(16))
	mBase = m.M
	*(*int32)(unsafe.Add(mBase, uint32(v379)+24)) = int32(268435456)
	switch v382 {
	case 0:
		goto L90
	default:
		goto L88
	case 2:
		goto L89
	}
L86:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_AutoVacWorkerMain[18])) = int32(970)
	goto L85
L87:
	;
	goto L92
L88:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v379)+24)) = int32(268435460)
	*(*int32)(unsafe.Add(mBase, uint32(v379)+12)) = int32(_a_F_AutoVacWorkerMain_0)
	goto L87
L89:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v379)+12)) = int32(0)
	goto L87
L90:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v379)+12)) = int32(-2)
	goto L87
L92:
	;
	goto L93
L93:
	;
	v411 = F___sigaction(m, int32(8), v379+int32(12), int32(0))
	mBase = m.M
	m.G0 = v379 + int32(32)
	goto L84
L94:
	;
	F_InitProcess(m)
	mBase = m.M
	v458 = m.ExcPending
	if v458 != 0 {
		goto L6
	} else {
		goto L104
	}
L95:
	;
	F_sigemptyset(m, v421+int32(16))
	mBase = m.M
	*(*int32)(unsafe.Add(mBase, uint32(v421)+24)) = int32(268435456)
	switch v423 {
	case 0:
		goto L100
	default:
		goto L98
	case 2:
		goto L99
	}
L96:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_AutoVacWorkerMain[19])) = int32(0)
	goto L95
L97:
	;
	goto L101
L98:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v421)+24)) = int32(268435460)
	*(*int32)(unsafe.Add(mBase, uint32(v421)+12)) = int32(_a_F_AutoVacWorkerMain_0)
	v446 = int32(268435461)
	goto L97
L99:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v421)+12)) = int32(0)
	v446 = int32(268435457)
	goto L97
L100:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v421)+12)) = int32(-2)
	v446 = int32(268435457)
	goto L97
L101:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v421)+24)) = v446
	goto L103
L103:
	;
	v453 = F___sigaction(m, int32(17), v421+int32(12), int32(0))
	mBase = m.M
	m.G0 = v421 + int32(32)
	goto L94
L104:
	;
	F_BaseInit(m)
	mBase = m.M
	v460 = m.ExcPending
	if v460 != 0 {
		goto L6
	} else {
		goto L105
	}
L105:
	;
	goto L106
L106:
	;
	v462 = v9 + int32(80)
	*(*int32)(unsafe.Add(mBase, uint32(v462)+4)) = int32(1)
	*(*int32)(unsafe.Add(mBase, uint32(v462))) = v9 + int32(12)
	goto L109
L107:
	;
	v470 = int32(0)
	goto L9
L109:
	;
	goto L107
L110:
	;
	v471 = int32(_a_F_AutoVacWorkerMain_1)
	v473 = *(*int32)(unsafe.Add(mBase, _c_F_AutoVacWorkerMain[20]))
	*(*int32)(unsafe.Add(mBase, _c_F_AutoVacWorkerMain[20])) = v473 + int32(1)
	*(*int32)(unsafe.Add(mBase, _c_F_AutoVacWorkerMain[21])) = int32(0)
	F_EmitErrorReport(m)
	mBase = m.M
	v481 = m.ExcPending
	if v481 != 0 {
		goto L6
	} else {
		goto L113
	}
L111:
	;
	goto L112
L112:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_AutoVacWorkerMain[22])) = v9 + int32(80)
	F_pgmem_sigprocmask(m, int32(_a_F_AutoVacWorkerMain_2), int32(0))
	mBase = m.M
	v492 = m.ExcPending
	if v492 != 0 {
		goto L6
	} else {
		goto L115
	}
L113:
	;
	F_proc_exit(m, int32(0))
	mBase = m.M
	v484 = m.ExcPending
	if v484 != 0 {
		goto L6
	} else {
		goto L114
	}
L114:
	;
	goto L3
L115:
	;
	F_SetConfigOption(m, int32(_a_F_AutoVacWorkerMain_3), int32(_a_F_AutoVacWorkerMain_4), int32(5), int32(10))
	mBase = m.M
	v498 = m.ExcPending
	if v498 != 0 {
		goto L6
	} else {
		goto L116
	}
L116:
	;
	F_SetConfigOption(m, int32(_a_F_AutoVacWorkerMain_5), int32(_a_F_AutoVacWorkerMain_6), int32(5), int32(10))
	mBase = m.M
	v504 = m.ExcPending
	if v504 != 0 {
		goto L6
	} else {
		goto L117
	}
L117:
	;
	F_SetConfigOption(m, int32(_a_F_AutoVacWorkerMain_7), int32(_a_F_AutoVacWorkerMain_8), int32(5), int32(10))
	mBase = m.M
	v510 = m.ExcPending
	if v510 != 0 {
		goto L6
	} else {
		goto L118
	}
L118:
	;
	F_SetConfigOption(m, int32(_a_F_AutoVacWorkerMain_9), int32(_a_F_AutoVacWorkerMain_8), int32(5), int32(10))
	mBase = m.M
	v516 = m.ExcPending
	if v516 != 0 {
		goto L6
	} else {
		goto L119
	}
L119:
	;
	F_SetConfigOption(m, int32(_a_F_AutoVacWorkerMain_10), int32(_a_F_AutoVacWorkerMain_8), int32(5), int32(10))
	mBase = m.M
	v522 = m.ExcPending
	if v522 != 0 {
		goto L6
	} else {
		goto L120
	}
L120:
	;
	F_SetConfigOption(m, int32(_a_F_AutoVacWorkerMain_11), int32(_a_F_AutoVacWorkerMain_8), int32(5), int32(10))
	mBase = m.M
	v528 = m.ExcPending
	if v528 != 0 {
		goto L6
	} else {
		goto L121
	}
L121:
	;
	F_SetConfigOption(m, int32(_a_F_AutoVacWorkerMain_12), int32(_a_F_AutoVacWorkerMain_13), int32(5), int32(10))
	mBase = m.M
	v534 = m.ExcPending
	if v534 != 0 {
		goto L6
	} else {
		goto L122
	}
L122:
	;
	v536 = *(*int32)(unsafe.Add(mBase, _c_F_AutoVacWorkerMain[23]))
	if int32(2) <= v536 {
		goto L123
	} else {
		goto L124
	}
L123:
	;
	F_SetConfigOption(m, int32(_a_F_AutoVacWorkerMain_14), int32(_a_F_AutoVacWorkerMain_15), int32(5), int32(10))
	mBase = m.M
	v544 = m.ExcPending
	if v544 != 0 {
		goto L6
	} else {
		goto L126
	}
L124:
	;
	goto L125
L125:
	;
	F_SetConfigOption(m, int32(_a_F_AutoVacWorkerMain_16), int32(_a_F_AutoVacWorkerMain_17), int32(5), int32(10))
	mBase = m.M
	v550 = m.ExcPending
	if v550 != 0 {
		goto L6
	} else {
		goto L127
	}
L126:
	;
	goto L125
L127:
	;
	v552 = *(*int32)(unsafe.Add(mBase, _c_F_AutoVacWorkerMain[24]))
	v556 = F_LWLockAcquire(m, v552+int32(2816), int32(0))
	mBase = m.M
	v557 = m.ExcPending
	if v557 != 0 {
		goto L6
	} else {
		goto L128
	}
L128:
	;
	v559 = *(*int32)(unsafe.Add(mBase, _c_F_AutoVacWorkerMain[25]))
	v560 = *(*int32)(unsafe.Add(mBase, uint32(v559)+32))
	if v560 != 0 {
		goto L130
	} else {
		goto L131
	}
L129:
	;
	F_proc_exit(m, int32(0))
	mBase = m.M
	v690 = m.ExcPending
	if v690 != 0 {
		goto L6
	} else {
		goto L165
	}
L130:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_AutoVacWorkerMain[26])) = v560
	v564 = *(*int32)(unsafe.Add(mBase, _c_F_AutoVacWorkerMain[27]))
	*(*int32)(unsafe.Add(mBase, uint32(v560)+16)) = v564
	v567 = v559 + int32(24)
	v568 = *(*int32)(unsafe.Add(mBase, uint32(v560)+8))
	v569 = *(*int32)(unsafe.Add(mBase, uint32(v559)+28))
	if v569 == int32(0) {
		goto L133
	} else {
		goto L134
	}
L131:
	;
	goto L132
L132:
	;
	v668 = F_errstart(m, int32(19), int32(0))
	mBase = m.M
	v669 = m.ExcPending
	if v669 != 0 {
		goto L6
	} else {
		goto L158
	}
L133:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v567))) = v567
	v573 = v567
	goto L135
L134:
	;
	v573 = v569
	goto L135
L135:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v560))) = v567
	*(*int32)(unsafe.Add(mBase, uint32(v560)+4)) = v573
	*(*int32)(unsafe.Add(mBase, uint32(v573))) = v560
	*(*int32)(unsafe.Add(mBase, uint32(v559)+28)) = v560
	v579 = *(*int32)(unsafe.Add(mBase, _c_F_AutoVacWorkerMain[25]))
	*(*int32)(unsafe.Add(mBase, uint32(v579)+32)) = int32(0)
	v583 = *(*int32)(unsafe.Add(mBase, _c_F_AutoVacWorkerMain[24]))
	F_LWLockRelease(m, v583+int32(2816))
	mBase = m.M
	v587 = m.ExcPending
	if v587 != 0 {
		goto L6
	} else {
		goto L136
	}
L136:
	;
	F_on_shmem_exit(m, int32(973), int64(0))
	mBase = m.M
	v591 = m.ExcPending
	if v591 != 0 {
		goto L6
	} else {
		goto L137
	}
L137:
	;
	v593 = *(*int32)(unsafe.Add(mBase, _c_F_AutoVacWorkerMain[25]))
	v594 = *(*int32)(unsafe.Add(mBase, uint32(v593)+8))
	if v594 != 0 {
		goto L138
	} else {
		goto L139
	}
L138:
	;
	v596 = F_pgmem_kill(m, v594, int32(12))
	mBase = m.M
	goto L140
L139:
	;
	goto L140
L140:
	;
	if v568 == int32(0) {
		goto L129
	} else {
		goto L141
	}
L141:
	;
	v602 = F_pgstat_get_entry_ref_locked(m, int32(1), v568, int64(0), int32(0))
	mBase = m.M
	v603 = m.ExcPending
	if v603 != 0 {
		goto L6
	} else {
		goto L142
	}
L142:
	;
	v604 = *(*int32)(unsafe.Add(mBase, uint32(v602)+4))
	v608 = m.G0
	v609 = int32(16)
	v610 = v608 - v609
	m.G0 = v610
	F_gettimeofday(m, v610)
	mBase = m.M
	v613 = *(*int64)(unsafe.Add(mBase, uint32(v610)))
	v614 = int64(*(*int32)(unsafe.Add(mBase, uint32(v610)+8)))
	m.G0 = v610 + v609
	goto L143
L143:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v604)+96)) = v614 + v613*int64(1000000) - int64(946684800000000)
	F_pgstat_unlock_entry(m, v602)
	mBase = m.M
	v625 = m.ExcPending
	if v625 != 0 {
		goto L6
	} else {
		goto L144
	}
L144:
	;
	v626 = int32(0)
	v631 = v9 + int32(16)
	F_InitPostgres(m, v626, v568, v626, v626, int32(2), v631)
	mBase = m.M
	v633 = m.ExcPending
	if v633 != 0 {
		goto L6
	} else {
		goto L145
	}
L145:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_AutoVacWorkerMain[28])) = int32(2)
	v637 = F_strlen(m, v631)
	mBase = m.M
	v640 = F_errstart(m, int32(14), int32(0))
	mBase = m.M
	v641 = m.ExcPending
	if v641 != 0 {
		goto L6
	} else {
		goto L146
	}
L146:
	;
	if v640 != 0 {
		goto L147
	} else {
		goto L148
	}
L147:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v9))) = v631
	F_errmsg_internal(m, int32(_a_F_AutoVacWorkerMain_18), v9)
	mBase = m.M
	v645 = m.ExcPending
	if v645 != 0 {
		goto L6
	} else {
		goto L150
	}
L148:
	;
	goto L149
L149:
	;
	v652 = *(*int32)(unsafe.Add(mBase, _c_F_AutoVacWorkerMain[29]))
	if v652 != 0 {
		goto L152
	} else {
		goto L153
	}
L150:
	;
	F_errfinish(m, int32(_a_F_AutoVacWorkerMain_19), int32(1629), int32(_a_F_AutoVacWorkerMain_20))
	mBase = m.M
	v650 = m.ExcPending
	if v650 != 0 {
		goto L6
	} else {
		goto L151
	}
L151:
	;
	goto L149
L152:
	;
	F_pg_usleep(m, v652*int32(_a_F_AutoVacWorkerMain_21))
	mBase = m.M
	goto L154
L153:
	;
	goto L154
L154:
	;
	v657 = F_ReadNextFullTransactionId(m)
	mBase = m.M
	v658 = m.ExcPending
	if v658 != 0 {
		goto L6
	} else {
		goto L155
	}
L155:
	;
	*(*uint32)(unsafe.Add(mBase, _c_F_AutoVacWorkerMain[30])) = uint32(v657)
	v661 = F_ReadNextMultiXactId(m)
	mBase = m.M
	v662 = m.ExcPending
	if v662 != 0 {
		goto L6
	} else {
		goto L156
	}
L156:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_AutoVacWorkerMain[31])) = v661
	F_do_autovacuum(m)
	mBase = m.M
	v665 = m.ExcPending
	if v665 != 0 {
		goto L6
	} else {
		goto L157
	}
L157:
	;
	goto L129
L158:
	;
	if v668 != 0 {
		goto L159
	} else {
		goto L160
	}
L159:
	;
	F_errmsg_internal(m, int32(_a_F_AutoVacWorkerMain_22), int32(0))
	mBase = m.M
	v673 = m.ExcPending
	if v673 != 0 {
		goto L6
	} else {
		goto L162
	}
L160:
	;
	goto L161
L161:
	;
	v680 = *(*int32)(unsafe.Add(mBase, _c_F_AutoVacWorkerMain[24]))
	F_LWLockRelease(m, v680+int32(2816))
	mBase = m.M
	v684 = m.ExcPending
	if v684 != 0 {
		goto L6
	} else {
		goto L164
	}
L162:
	;
	F_errfinish(m, int32(_a_F_AutoVacWorkerMain_19), int32(1596), int32(_a_F_AutoVacWorkerMain_20))
	mBase = m.M
	v678 = m.ExcPending
	if v678 != 0 {
		goto L6
	} else {
		goto L163
	}
L163:
	;
	goto L161
L164:
	;
	goto L129
L165:
	;
	goto L5
L166:
	;
	v701 = int32(v697)
	m.G0 = v9
	v703 = *(*int32)(unsafe.Add(mBase, uint32(v701)+4))
	v704 = *(*int32)(unsafe.Add(mBase, uint32(v701)))
	v707 = *(*int32)(unsafe.Add(mBase, uint32(v704)))
	if v9+int32(12) == v707 {
		goto L169
	} else {
		goto L170
	}
L167:
	;
	m.ExcPending = 1
	goto L175
L168:
	;
	if v711 == int32(0) {
		goto L172
	} else {
		goto L173
	}
L169:
	;
	v709 = *(*int32)(unsafe.Add(mBase, uint32(v704)+4))
	v711 = v709
	goto L171
L170:
	;
	v711 = int32(0)
	goto L171
L171:
	;
	goto L168
L172:
	;
	F___wasm_longjmp(m, v704, v703)
	mBase = m.M
	v715 = m.ExcPending
	if v715 != 0 {
		goto L175
	} else {
		goto L176
	}
L173:
	;
	goto L174
L174:
	;
	v12 = v711
	v14 = v703
	goto L1
L175:
	;
	return
L176:
	;
	base.Wasm_trap_unreachable()
	for {
	}
}
func F_AutoVacuumShmemInit(m *base.Module, l0 int32) {
	mBase := m.M
	_ = mBase
	var v2 int32
	_ = v2
	var v8 int32
	_ = v8
	var v14 int32
	_ = v14
	var v18 int32
	_ = v18
	var v27 int32
	_ = v27
	var v36 int32
	_ = v36
	var v39 int32
	_ = v39
	var v41 int32
	_ = v41
	var v42 int32
	_ = v42
	var v50 int32
	_ = v50
	var v53 int32
	_ = v53
	var v58 int32
	_ = v58
	var v59 int32
	_ = v59
	var v62 int32
	_ = v62
	var v66 int32
	_ = v66
	var v68 int32
	_ = v68
	var v71 int32
	_ = v71
	var v78 int32
	_ = v78
	v2 = int32(0)
	v8 = *(*int32)(unsafe.Add(mBase, _c_F_AutoVacuumShmemInit[0]))
	*(*int32)(unsafe.Add(mBase, uint32(v8)+20)) = v2
	*(*int32)(unsafe.Add(mBase, uint32(v8)+8)) = v2
	v14 = v8 + int32(24)
	*(*int32)(unsafe.Add(mBase, uint32(v8)+28)) = v14
	*(*int32)(unsafe.Add(mBase, uint32(v8)+24)) = v14
	v18 = v8 + int32(12)
	*(*int32)(unsafe.Add(mBase, uint32(v8)+16)) = v18
	*(*int32)(unsafe.Add(mBase, uint32(v8)+12)) = v18
	base.MemoryFill(m, v8+int32(32), v2, int32(_a_F_AutoVacuumShmemInit_0))
	v27 = *(*int32)(unsafe.Add(mBase, _c_F_AutoVacuumShmemInit[1]))
	if v2 < v27 {
		v36 = v2
		for {
			v39 = *(*int32)(unsafe.Add(mBase, _c_F_AutoVacuumShmemInit[0]))
			v41 = v39 + int32(12)
			v42 = *(*int32)(unsafe.Add(mBase, uint32(v39)+16))
			if v42 == int32(0) {
				*(*int32)(unsafe.Add(mBase, uint32(v39)+20)) = int32(0)
				*(*int32)(unsafe.Add(mBase, uint32(v39)+12)) = v39 + int32(12)
				v50 = v41
			} else {
				v50 = v42
			}
			v53 = v8 + int32(_a_F_AutoVacuumShmemInit_1) + v36*int32(40)
			*(*int32)(unsafe.Add(mBase, uint32(v53))) = v41
			*(*int32)(unsafe.Add(mBase, uint32(v53)+4)) = v50
			*(*int32)(unsafe.Add(mBase, uint32(v50))) = v53
			*(*int32)(unsafe.Add(mBase, uint32(v39)+16)) = v53
			v58 = *(*int32)(unsafe.Add(mBase, uint32(v39)+20))
			v59 = int32(1)
			*(*int32)(unsafe.Add(mBase, uint32(v39)+20)) = v58 + v59
			v62 = int32(0)
			atomic.StoreUint32((*uint32)(unsafe.Add(mBase, uint32(v53)+32)), uint32(v62))
			v66 = v36 + v59
			v68 = *(*int32)(unsafe.Add(mBase, _c_F_AutoVacuumShmemInit[1]))
			if v66 < v68 {
				v36 = v66
				continue
			} else {
				break
			}
			break
		}
		v71 = *(*int32)(unsafe.Add(mBase, _c_F_AutoVacuumShmemInit[0]))
		v78 = v71
	} else {
		v78 = v8
	}
	*(*int32)(unsafe.Add(mBase, uint32(v78)+uint32(_c_F_AutoVacuumShmemInit[2]))) = int32(0)
	return
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
	var v5 int64
	_ = v5
	var v7 int64
	_ = v7
	var v10 int32
	_ = v10
	var v15 int64
	_ = v15
	var v16 int32
	_ = v16
	var v18 int32
	_ = v18
	v5 = base.I64_extend_i32_u(l0)
	v7 = F_DirectFunctionCall2Coll(m, int32(2657), int32(0), v5, int64(4165696))
	v10 = m.ExcPending
	if v10 != 0 {
		return int32(0)
	} else {
		if v7 != int64(0) {
			v15 = F_DirectFunctionCall1Coll(m, int32(2661), int32(0), v5)
			v16 = m.ExcPending
			if v16 != 0 {
				return int32(0)
			} else {
				v18 = base.I32_wrap_i64(v15)
				return v18
			}
		} else {
			v18 = l0
			return v18
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
			F_aclcheck_error(m, l0, int32(50), v7)
			v10 = m.ExcPending
			if v10 != 0 {
				return
			} else {
				return
			}
		}
	}
}
func F_acldefault_sql(m *base.Module, l0 int32) int64 {
	mBase := m.M
	_ = mBase
	var v5 int32
	_ = v5
	var v7 int32
	_ = v7
	var v9 int64
	_ = v9
	var v11 int32
	_ = v11
	var v30 int32
	_ = v30
	var v35 int32
	_ = v35
	var v40 int32
	_ = v40
	var v42 int32
	_ = v42
	var v44 int32
	_ = v44
	var v45 int32
	_ = v45
	v5 = m.G0
	v7 = v5 - int32(16)
	m.G0 = v7
	v9 = *(*int64)(unsafe.Add(mBase, uint32(l0)+40))
	v11 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+24)))
	switch v11 - int32(70) {
	case 0:
		v42 = int32(16)
		v44 = F_acldefault(m, v42, base.I32_wrap_i64(v9))
		mBase = m.M
		v45 = m.ExcPending
		if v45 != 0 {
			return int64(0)
		} else {
			m.G0 = v7 + int32(16)
			return base.I64_extend_i32_u(v44)
		}
	default:
		F_errstart_cold(m, int32(21), int32(0))
		mBase = m.M
		v30 = m.ExcPending
		if v30 != 0 {
			return int64(0)
		} else {
			*(*int32)(unsafe.Add(mBase, uint32(v7))) = base.I32_extend8_s(v11)
			F_errmsg_internal(m, int32(_a_F_acldefault_sql_0), v7)
			mBase = m.M
			v35 = m.ExcPending
			if v35 != 0 {
				return int64(0)
			} else {
				F_errfinish(m, int32(_a_F_acldefault_sql_1), int32(992), int32(_a_F_acldefault_sql_2))
				mBase = m.M
				v40 = m.ExcPending
				if v40 != 0 {
					return int64(0)
				} else {
					base.Wasm_trap_unreachable()
					for {
					}
				}
			}
		}
	case 6:
		v42 = int32(22)
		v44 = F_acldefault(m, v42, base.I32_wrap_i64(v9))
		mBase = m.M
		v45 = m.ExcPending
		if v45 != 0 {
			return int64(0)
		} else {
			m.G0 = v7 + int32(16)
			return base.I64_extend_i32_u(v44)
		}
	case 13:
		v42 = int32(17)
		v44 = F_acldefault(m, v42, base.I32_wrap_i64(v9))
		mBase = m.M
		v45 = m.ExcPending
		if v45 != 0 {
			return int64(0)
		} else {
			m.G0 = v7 + int32(16)
			return base.I64_extend_i32_u(v44)
		}
	case 14:
		v42 = int32(50)
		v44 = F_acldefault(m, v42, base.I32_wrap_i64(v9))
		mBase = m.M
		v45 = m.ExcPending
		if v45 != 0 {
			return int64(0)
		} else {
			m.G0 = v7 + int32(16)
			return base.I64_extend_i32_u(v44)
		}
	case 29:
		v42 = int32(6)
		v44 = F_acldefault(m, v42, base.I32_wrap_i64(v9))
		mBase = m.M
		v45 = m.ExcPending
		if v45 != 0 {
			return int64(0)
		} else {
			m.G0 = v7 + int32(16)
			return base.I64_extend_i32_u(v44)
		}
	case 30:
		v42 = int32(9)
		v44 = F_acldefault(m, v42, base.I32_wrap_i64(v9))
		mBase = m.M
		v45 = m.ExcPending
		if v45 != 0 {
			return int64(0)
		} else {
			m.G0 = v7 + int32(16)
			return base.I64_extend_i32_u(v44)
		}
	case 32:
		v42 = int32(19)
		v44 = F_acldefault(m, v42, base.I32_wrap_i64(v9))
		mBase = m.M
		v45 = m.ExcPending
		if v45 != 0 {
			return int64(0)
		} else {
			m.G0 = v7 + int32(16)
			return base.I64_extend_i32_u(v44)
		}
	case 38:
		v42 = int32(21)
		v44 = F_acldefault(m, v42, base.I32_wrap_i64(v9))
		mBase = m.M
		v45 = m.ExcPending
		if v45 != 0 {
			return int64(0)
		} else {
			m.G0 = v7 + int32(16)
			return base.I64_extend_i32_u(v44)
		}
	case 40:
		v42 = int32(37)
		v44 = F_acldefault(m, v42, base.I32_wrap_i64(v9))
		mBase = m.M
		v45 = m.ExcPending
		if v45 != 0 {
			return int64(0)
		} else {
			m.G0 = v7 + int32(16)
			return base.I64_extend_i32_u(v44)
		}
	case 42:
		v42 = int32(27)
		v44 = F_acldefault(m, v42, base.I32_wrap_i64(v9))
		mBase = m.M
		v45 = m.ExcPending
		if v45 != 0 {
			return int64(0)
		} else {
			m.G0 = v7 + int32(16)
			return base.I64_extend_i32_u(v44)
		}
	case 44:
		v42 = int32(42)
		v44 = F_acldefault(m, v42, base.I32_wrap_i64(v9))
		mBase = m.M
		v45 = m.ExcPending
		if v45 != 0 {
			return int64(0)
		} else {
			m.G0 = v7 + int32(16)
			return base.I64_extend_i32_u(v44)
		}
	case 45:
		v42 = int32(38)
		v44 = F_acldefault(m, v42, base.I32_wrap_i64(v9))
		mBase = m.M
		v45 = m.ExcPending
		if v45 != 0 {
			return int64(0)
		} else {
			m.G0 = v7 + int32(16)
			return base.I64_extend_i32_u(v44)
		}
	case 46:
		v42 = int32(43)
		v44 = F_acldefault(m, v42, base.I32_wrap_i64(v9))
		mBase = m.M
		v45 = m.ExcPending
		if v45 != 0 {
			return int64(0)
		} else {
			m.G0 = v7 + int32(16)
			return base.I64_extend_i32_u(v44)
		}
	}
}
func F_aclitemin(m *base.Module, l0 int32) int64 {
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
	var v24 int32
	_ = v24
	var v25 int32
	_ = v25
	var v26 int32
	_ = v26
	var v29 int32
	_ = v29
	var v32 int32
	_ = v32
	var v35 int32
	_ = v35
	var v41 int32
	_ = v41
	var v42 int32
	_ = v42
	var v48 int32
	_ = v48
	var v49 int32
	_ = v49
	var v54 int32
	_ = v54
	var v60 int32
	_ = v60
	var v64 int32
	_ = v64
	var v69 int32
	_ = v69
	var v72 int32
	_ = v72
	var v73 int32
	_ = v73
	var v76 int32
	_ = v76
	var v79 int32
	_ = v79
	var v80 int32
	_ = v80
	var v85 int32
	_ = v85
	var v89 int32
	_ = v89
	var v93 int32
	_ = v93
	var v98 int32
	_ = v98
	var v99 int32
	_ = v99
	var v102 int32
	_ = v102
	var v104 int32
	_ = v104
	var v105 int32
	_ = v105
	var v108 int32
	_ = v108
	var v109 int32
	_ = v109
	var v114 int32
	_ = v114
	var v118 int32
	_ = v118
	var v123 int32
	_ = v123
	var v136 int32
	_ = v136
	var v138 int32
	_ = v138
	var v140 int64
	_ = v140
	var v141 int64
	_ = v141
	var v142 int64
	_ = v142
	var v160 int32
	_ = v160
	var v161 int32
	_ = v161
	var v166 int32
	_ = v166
	var v171 int32
	_ = v171
	var v176 int32
	_ = v176
	var v178 int64
	_ = v178
	var v179 int64
	_ = v179
	var v181 int32
	_ = v181
	var v182 int64
	_ = v182
	var v183 int32
	_ = v183
	var v201 int32
	_ = v201
	var v202 int32
	_ = v202
	var v206 int64
	_ = v206
	var v209 int32
	_ = v209
	var v215 int32
	_ = v215
	var v219 int32
	_ = v219
	var v220 int32
	_ = v220
	var v223 int32
	_ = v223
	var v228 int32
	_ = v228
	var v233 int32
	_ = v233
	var v238 int32
	_ = v238
	var v243 int32
	_ = v243
	var v248 int32
	_ = v248
	var v253 int32
	_ = v253
	var v258 int32
	_ = v258
	var v263 int32
	_ = v263
	var v268 int32
	_ = v268
	var v273 int32
	_ = v273
	var v278 int32
	_ = v278
	var v283 int32
	_ = v283
	var v288 int32
	_ = v288
	var v293 int32
	_ = v293
	var v298 int32
	_ = v298
	var v302 int32
	_ = v302
	var v304 int32
	_ = v304
	var v305 int32
	_ = v305
	var v306 int32
	_ = v306
	var v311 int64
	_ = v311
	var v314 int32
	_ = v314
	var v315 int32
	_ = v315
	var v316 int32
	_ = v316
	var v318 int32
	_ = v318
	var v319 int32
	_ = v319
	var v324 int32
	_ = v324
	var v332 int32
	_ = v332
	var v337 int32
	_ = v337
	var v339 int32
	_ = v339
	var v346 int32
	_ = v346
	var v347 int32
	_ = v347
	var v350 int32
	_ = v350
	var v353 int32
	_ = v353
	var v354 int32
	_ = v354
	var v359 int32
	_ = v359
	var v363 int32
	_ = v363
	var v368 int32
	_ = v368
	var v370 int32
	_ = v370
	var v374 int32
	_ = v374
	var v375 int32
	_ = v375
	var v378 int32
	_ = v378
	var v383 int32
	_ = v383
	var v388 int32
	_ = v388
	var v393 int32
	_ = v393
	var v398 int32
	_ = v398
	var v403 int32
	_ = v403
	var v408 int32
	_ = v408
	var v413 int32
	_ = v413
	var v418 int32
	_ = v418
	var v423 int32
	_ = v423
	var v428 int32
	_ = v428
	var v433 int32
	_ = v433
	var v438 int32
	_ = v438
	var v443 int32
	_ = v443
	var v448 int32
	_ = v448
	var v453 int32
	_ = v453
	var v457 int32
	_ = v457
	var v459 int32
	_ = v459
	var v460 int32
	_ = v460
	var v461 int32
	_ = v461
	var v466 int64
	_ = v466
	var v469 int32
	_ = v469
	var v470 int32
	_ = v470
	var v471 int32
	_ = v471
	var v473 int32
	_ = v473
	var v474 int32
	_ = v474
	var v479 int32
	_ = v479
	var v487 int32
	_ = v487
	var v492 int32
	_ = v492
	var v496 int32
	_ = v496
	var v501 int32
	_ = v501
	var v502 int32
	_ = v502
	var v507 int32
	_ = v507
	var v514 int32
	_ = v514
	var v519 int32
	_ = v519
	var v521 int32
	_ = v521
	var v527 int32
	_ = v527
	var v534 int32
	_ = v534
	var v544 int64
	_ = v544
	var v545 int32
	_ = v545
	var v546 int32
	_ = v546
	var v551 int32
	_ = v551
	var v555 int32
	_ = v555
	var v560 int32
	_ = v560
	var v573 int32
	_ = v573
	var v583 int64
	_ = v583
	v8 = int64(0)
	v12 = m.G0
	v14 = v12 - int32(208)
	m.G0 = v14
	v16 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v17 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
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
	return int64(0)
L2:
	;
	v24 = v14 + int32(144)
	v25 = F_getid(m, v17, v24, v16)
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
	return v583
L4:
	;
	v573 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(l0)+16)) = uint8(v573)
	v583 = int64(0)
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
	v209 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v14)+144)))
	if v209 == int32(0) {
		goto L68
	} else {
		goto L69
	}
L8:
	;
	v136 = v104
	v138 = v105
	v140 = v8
	v141 = v8
	v142 = v8
	goto L42
L9:
	;
	if base.Ui32(int32(26)) <= base.Ui32((v105|int32(32)-int32(97))&int32(255)) {
		v201 = v102
		v202 = v104
		v206 = v8
		goto L7
	} else {
		goto L41
	}
L10:
	;
	v108 = F_errsave_start(m, v16)
	mBase = m.M
	v109 = m.ExcPending
	if v109 != 0 {
		goto L1
	} else {
		goto L36
	}
L11:
	;
	v32 = *(*int32)(unsafe.Add(mBase, uint32(v14)+144))
	v35 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v14)+148)))
	if v32^int32(1970238055)|(v35^int32(112)) == int32(0) {
		goto L14
	} else {
		goto L15
	}
L12:
	;
	v102 = v25
	goto L13
L13:
	;
	v104 = v102 + int32(1)
	v105 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v102)+1)))
	if v105 != int32(42) {
		goto L9
	} else {
		goto L35
	}
L14:
	;
	v72 = F_getid(m, v25, v14+int32(144), v16)
	mBase = m.M
	v73 = m.ExcPending
	if v73 != 0 {
		goto L1
	} else {
		goto L23
	}
L15:
	;
	v41 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v14)+148)))
	v42 = *(*int32)(unsafe.Add(mBase, uint32(v14)+144))
	if v41|(v42^int32(1919251317)) == int32(0) {
		goto L14
	} else {
		goto L16
	}
L16:
	;
	v48 = F_errsave_start(m, v16)
	mBase = m.M
	v49 = m.ExcPending
	if v49 != 0 {
		goto L1
	} else {
		goto L17
	}
L17:
	;
	if v48 == int32(0) {
		goto L4
	} else {
		goto L18
	}
L18:
	;
	F_errcode(m, int32(33685634))
	mBase = m.M
	v54 = m.ExcPending
	if v54 != 0 {
		goto L1
	} else {
		goto L19
	}
L19:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v14)+64)) = v24
	F_errmsg(m, int32(_a_F_aclitemin_0), v14-int32(-64))
	mBase = m.M
	v60 = m.ExcPending
	if v60 != 0 {
		goto L1
	} else {
		goto L20
	}
L20:
	;
	F_errhint(m, int32(_a_F_aclitemin_1), int32(0))
	mBase = m.M
	v64 = m.ExcPending
	if v64 != 0 {
		goto L1
	} else {
		goto L21
	}
L21:
	;
	F_errsave_finish(m, v16, int32(_a_F_aclitemin_2), int32(299), int32(_a_F_aclitemin_3))
	mBase = m.M
	v69 = m.ExcPending
	if v69 != 0 {
		goto L1
	} else {
		goto L22
	}
L22:
	;
	goto L4
L23:
	;
	if v72 == int32(0) {
		goto L4
	} else {
		goto L24
	}
L24:
	;
	v76 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v14)+144)))
	if v76 == int32(0) {
		goto L25
	} else {
		goto L26
	}
L25:
	;
	v79 = F_errsave_start(m, v16)
	mBase = m.M
	v80 = m.ExcPending
	if v80 != 0 {
		goto L1
	} else {
		goto L28
	}
L26:
	;
	goto L27
L27:
	;
	v99 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v72))))
	if v99 != int32(61) {
		goto L10
	} else {
		goto L34
	}
L28:
	;
	if v79 == int32(0) {
		goto L4
	} else {
		goto L29
	}
L29:
	;
	F_errcode(m, int32(33685634))
	mBase = m.M
	v85 = m.ExcPending
	if v85 != 0 {
		goto L1
	} else {
		goto L30
	}
L30:
	;
	F_errmsg(m, int32(_a_F_aclitemin_4), int32(0))
	mBase = m.M
	v89 = m.ExcPending
	if v89 != 0 {
		goto L1
	} else {
		goto L31
	}
L31:
	;
	F_errhint(m, int32(_a_F_aclitemin_5), int32(0))
	mBase = m.M
	v93 = m.ExcPending
	if v93 != 0 {
		goto L1
	} else {
		goto L32
	}
L32:
	;
	F_errsave_finish(m, v16, int32(_a_F_aclitemin_2), int32(308), int32(_a_F_aclitemin_3))
	mBase = m.M
	v98 = m.ExcPending
	if v98 != 0 {
		goto L1
	} else {
		goto L33
	}
L33:
	;
	goto L4
L34:
	;
	v102 = v72
	goto L13
L35:
	;
	goto L8
L36:
	;
	if v108 == int32(0) {
		goto L4
	} else {
		goto L37
	}
L37:
	;
	F_errcode(m, int32(33685634))
	mBase = m.M
	v114 = m.ExcPending
	if v114 != 0 {
		goto L1
	} else {
		goto L38
	}
L38:
	;
	F_errmsg(m, int32(_a_F_aclitemin_6), int32(0))
	mBase = m.M
	v118 = m.ExcPending
	if v118 != 0 {
		goto L1
	} else {
		goto L39
	}
L39:
	;
	F_errsave_finish(m, v16, int32(_a_F_aclitemin_2), int32(314), int32(_a_F_aclitemin_3))
	mBase = m.M
	v123 = m.ExcPending
	if v123 != 0 {
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
	switch v138 - int32(42) {
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
		v178 = int64(1)
		v179 = v141
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
	v201 = v136
	v202 = v181
	v206 = v179<<(uint(int64(32))%64) | v182
	goto L7
L44:
	;
	v181 = v136 + int32(1)
	v182 = v178 | v140
	v183 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v136)+1)))
	if base.B2i32(v183 == int32(42))|base.B2i32(base.Ui32((v183|int32(32)-int32(97))&int32(255)) < base.Ui32(int32(26))) != 0 {
		v136 = v181
		v138 = v183
		v140 = v182
		v141 = v179
		v142 = v178
		goto L42
	} else {
		goto L66
	}
L45:
	;
	v178 = v142
	v179 = v141 | v142
	goto L44
L46:
	;
	v160 = F_errsave_start(m, v16)
	mBase = m.M
	v161 = m.ExcPending
	if v161 != 0 {
		goto L1
	} else {
		goto L61
	}
L47:
	;
	v178 = int64(16384)
	v179 = v141
	goto L44
L48:
	;
	v178 = int64(8192)
	v179 = v141
	goto L44
L49:
	;
	v178 = int64(4096)
	v179 = v141
	goto L44
L50:
	;
	v178 = int64(2048)
	v179 = v141
	goto L44
L51:
	;
	v178 = int64(1024)
	v179 = v141
	goto L44
L52:
	;
	v178 = int64(512)
	v179 = v141
	goto L44
L53:
	;
	v178 = int64(256)
	v179 = v141
	goto L44
L54:
	;
	v178 = int64(128)
	v179 = v141
	goto L44
L55:
	;
	v178 = int64(64)
	v179 = v141
	goto L44
L56:
	;
	v178 = int64(32)
	v179 = v141
	goto L44
L57:
	;
	v178 = int64(16)
	v179 = v141
	goto L44
L58:
	;
	v178 = int64(8)
	v179 = v141
	goto L44
L59:
	;
	v178 = int64(4)
	v179 = v141
	goto L44
L60:
	;
	v178 = int64(2)
	v179 = v141
	goto L44
L61:
	;
	if v160 == int32(0) {
		goto L4
	} else {
		goto L62
	}
L62:
	;
	F_errcode(m, int32(33685634))
	mBase = m.M
	v166 = m.ExcPending
	if v166 != 0 {
		goto L1
	} else {
		goto L63
	}
L63:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v14))) = int32(_a_F_aclitemin_7)
	F_errmsg(m, int32(_a_F_aclitemin_8), v14)
	mBase = m.M
	v171 = m.ExcPending
	if v171 != 0 {
		goto L1
	} else {
		goto L64
	}
L64:
	;
	F_errsave_finish(m, v16, int32(_a_F_aclitemin_2), int32(374), int32(_a_F_aclitemin_3))
	mBase = m.M
	v176 = m.ExcPending
	if v176 != 0 {
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
	v339 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v202))))
	if v339 == int32(47) {
		goto L103
	} else {
		goto L104
	}
L68:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v19))) = int32(0)
	goto L67
L69:
	;
	goto L70
L70:
	;
	v215 = *(*int32)(unsafe.Add(mBase, _c_F_aclitemin[0]))
	if v215 == int32(0) {
		goto L72
	} else {
		goto L73
	}
L71:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v19))) = v316
	if v316 != 0 {
		goto L67
	} else {
		goto L96
	}
L72:
	;
	v219 = v14 + int32(144)
	v220 = int32(0)
	v223 = F_strcmp(m, int32(_a_F_aclitemin_9), v219)
	mBase = m.M
	if v223 == v220 {
		v304 = int32(_a_F_aclitemin_10)
		goto L77
	} else {
		goto L78
	}
L73:
	;
	goto L74
L74:
	;
	v311 = int64(0)
	v314 = F_GetSysCacheOid(m, int32(10), base.I64_extend_i32_u(v14+int32(144)), v311, v311, v311)
	mBase = m.M
	v315 = m.ExcPending
	if v315 != 0 {
		goto L1
	} else {
		goto L95
	}
L75:
	;
	v316 = v306
	goto L71
L76:
	;
	goto L75
L77:
	;
	v305 = *(*int32)(unsafe.Add(mBase, uint32(v304)+4))
	v306 = v305
	goto L76
L78:
	;
	v228 = F_strcmp(m, int32(_a_F_aclitemin_11), v219)
	mBase = m.M
	if v228 == int32(0) {
		v304 = int32(_a_F_aclitemin_12)
		goto L77
	} else {
		goto L79
	}
L79:
	;
	v233 = F_strcmp(m, int32(_a_F_aclitemin_13), v219)
	mBase = m.M
	if v233 == int32(0) {
		v304 = int32(_a_F_aclitemin_14)
		goto L77
	} else {
		goto L80
	}
L80:
	;
	v238 = F_strcmp(m, int32(_a_F_aclitemin_15), v219)
	mBase = m.M
	if v238 == int32(0) {
		v304 = int32(_a_F_aclitemin_16)
		goto L77
	} else {
		goto L81
	}
L81:
	;
	v243 = F_strcmp(m, int32(_a_F_aclitemin_17), v219)
	mBase = m.M
	if v243 == int32(0) {
		v304 = int32(_a_F_aclitemin_18)
		goto L77
	} else {
		goto L82
	}
L82:
	;
	v248 = F_strcmp(m, int32(_a_F_aclitemin_19), v219)
	mBase = m.M
	if v248 == int32(0) {
		v304 = int32(_a_F_aclitemin_20)
		goto L77
	} else {
		goto L83
	}
L83:
	;
	v253 = F_strcmp(m, int32(_a_F_aclitemin_21), v219)
	mBase = m.M
	if v253 == int32(0) {
		v304 = int32(_a_F_aclitemin_22)
		goto L77
	} else {
		goto L84
	}
L84:
	;
	v258 = F_strcmp(m, int32(_a_F_aclitemin_23), v219)
	mBase = m.M
	if v258 == int32(0) {
		v304 = int32(_a_F_aclitemin_24)
		goto L77
	} else {
		goto L85
	}
L85:
	;
	v263 = F_strcmp(m, int32(_a_F_aclitemin_25), v219)
	mBase = m.M
	if v263 == int32(0) {
		v304 = int32(_a_F_aclitemin_26)
		goto L77
	} else {
		goto L86
	}
L86:
	;
	v268 = F_strcmp(m, int32(_a_F_aclitemin_27), v219)
	mBase = m.M
	if v268 == int32(0) {
		v304 = int32(_a_F_aclitemin_28)
		goto L77
	} else {
		goto L87
	}
L87:
	;
	v273 = F_strcmp(m, int32(_a_F_aclitemin_29), v219)
	mBase = m.M
	if v273 == int32(0) {
		v304 = int32(_a_F_aclitemin_30)
		goto L77
	} else {
		goto L88
	}
L88:
	;
	v278 = F_strcmp(m, int32(_a_F_aclitemin_31), v219)
	mBase = m.M
	if v278 == int32(0) {
		v304 = int32(_a_F_aclitemin_32)
		goto L77
	} else {
		goto L89
	}
L89:
	;
	v283 = F_strcmp(m, int32(_a_F_aclitemin_33), v219)
	mBase = m.M
	if v283 == int32(0) {
		v304 = int32(_a_F_aclitemin_34)
		goto L77
	} else {
		goto L90
	}
L90:
	;
	v288 = F_strcmp(m, int32(_a_F_aclitemin_35), v219)
	mBase = m.M
	if v288 == int32(0) {
		v304 = int32(_a_F_aclitemin_36)
		goto L77
	} else {
		goto L91
	}
L91:
	;
	v293 = F_strcmp(m, int32(_a_F_aclitemin_37), v219)
	mBase = m.M
	if v293 == int32(0) {
		v304 = int32(_a_F_aclitemin_38)
		goto L77
	} else {
		goto L92
	}
L92:
	;
	v298 = F_strcmp(m, int32(_a_F_aclitemin_39), v219)
	mBase = m.M
	if v298 == int32(0) {
		v304 = int32(_a_F_aclitemin_40)
		goto L77
	} else {
		goto L93
	}
L93:
	;
	v302 = F_strcmp(m, int32(_a_F_aclitemin_41), v219)
	mBase = m.M
	if v302 != 0 {
		v306 = v220
		goto L76
	} else {
		goto L94
	}
L94:
	;
	v304 = int32(_a_F_aclitemin_42)
	goto L77
L95:
	;
	v316 = v314
	goto L71
L96:
	;
	v318 = F_errsave_start(m, v16)
	mBase = m.M
	v319 = m.ExcPending
	if v319 != 0 {
		goto L1
	} else {
		goto L97
	}
L97:
	;
	if v318 == int32(0) {
		goto L4
	} else {
		goto L98
	}
L98:
	;
	F_errcode(m, int32(67137668))
	mBase = m.M
	v324 = m.ExcPending
	if v324 != 0 {
		goto L1
	} else {
		goto L99
	}
L99:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v14)+48)) = v14 + int32(144)
	F_errmsg(m, int32(_a_F_aclitemin_43), v14+int32(48))
	mBase = m.M
	v332 = m.ExcPending
	if v332 != 0 {
		goto L1
	} else {
		goto L100
	}
L100:
	;
	F_errsave_finish(m, v16, int32(_a_F_aclitemin_2), int32(391), int32(_a_F_aclitemin_3))
	mBase = m.M
	v337 = m.ExcPending
	if v337 != 0 {
		goto L1
	} else {
		goto L101
	}
L101:
	;
	goto L4
L102:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v19)+8)) = v206
	v527 = v521
	goto L153
L103:
	;
	v346 = F_getid(m, v201+int32(2), v14+int32(80), v16)
	mBase = m.M
	v347 = m.ExcPending
	if v347 != 0 {
		goto L1
	} else {
		goto L106
	}
L104:
	;
	goto L105
L105:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v19)+4)) = int32(10)
	v496 = *(*int32)(unsafe.Add(mBase, _c_F_aclitemin[0]))
	if v496 == int32(0) {
		v521 = v202
		goto L102
	} else {
		goto L147
	}
L106:
	;
	if v346 == int32(0) {
		goto L4
	} else {
		goto L107
	}
L107:
	;
	v350 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v14)+80)))
	if v350 == int32(0) {
		goto L108
	} else {
		goto L109
	}
L108:
	;
	v353 = F_errsave_start(m, v16)
	mBase = m.M
	v354 = m.ExcPending
	if v354 != 0 {
		goto L1
	} else {
		goto L111
	}
L109:
	;
	goto L110
L110:
	;
	v370 = *(*int32)(unsafe.Add(mBase, _c_F_aclitemin[0]))
	if v370 == int32(0) {
		goto L117
	} else {
		goto L118
	}
L111:
	;
	if v353 == int32(0) {
		goto L4
	} else {
		goto L112
	}
L112:
	;
	F_errcode(m, int32(33685634))
	mBase = m.M
	v359 = m.ExcPending
	if v359 != 0 {
		goto L1
	} else {
		goto L113
	}
L113:
	;
	F_errmsg(m, int32(_a_F_aclitemin_44), int32(0))
	mBase = m.M
	v363 = m.ExcPending
	if v363 != 0 {
		goto L1
	} else {
		goto L114
	}
L114:
	;
	F_errsave_finish(m, v16, int32(_a_F_aclitemin_2), int32(407), int32(_a_F_aclitemin_3))
	mBase = m.M
	v368 = m.ExcPending
	if v368 != 0 {
		goto L1
	} else {
		goto L115
	}
L115:
	;
	goto L4
L116:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v19)+4)) = v471
	if v471 != 0 {
		v521 = v346
		goto L102
	} else {
		goto L141
	}
L117:
	;
	v374 = v14 + int32(80)
	v375 = int32(0)
	v378 = F_strcmp(m, int32(_a_F_aclitemin_9), v374)
	mBase = m.M
	if v378 == v375 {
		v459 = int32(_a_F_aclitemin_10)
		goto L122
	} else {
		goto L123
	}
L118:
	;
	goto L119
L119:
	;
	v466 = int64(0)
	v469 = F_GetSysCacheOid(m, int32(10), base.I64_extend_i32_u(v14+int32(80)), v466, v466, v466)
	mBase = m.M
	v470 = m.ExcPending
	if v470 != 0 {
		goto L1
	} else {
		goto L140
	}
L120:
	;
	v471 = v461
	goto L116
L121:
	;
	goto L120
L122:
	;
	v460 = *(*int32)(unsafe.Add(mBase, uint32(v459)+4))
	v461 = v460
	goto L121
L123:
	;
	v383 = F_strcmp(m, int32(_a_F_aclitemin_11), v374)
	mBase = m.M
	if v383 == int32(0) {
		v459 = int32(_a_F_aclitemin_12)
		goto L122
	} else {
		goto L124
	}
L124:
	;
	v388 = F_strcmp(m, int32(_a_F_aclitemin_13), v374)
	mBase = m.M
	if v388 == int32(0) {
		v459 = int32(_a_F_aclitemin_14)
		goto L122
	} else {
		goto L125
	}
L125:
	;
	v393 = F_strcmp(m, int32(_a_F_aclitemin_15), v374)
	mBase = m.M
	if v393 == int32(0) {
		v459 = int32(_a_F_aclitemin_16)
		goto L122
	} else {
		goto L126
	}
L126:
	;
	v398 = F_strcmp(m, int32(_a_F_aclitemin_17), v374)
	mBase = m.M
	if v398 == int32(0) {
		v459 = int32(_a_F_aclitemin_18)
		goto L122
	} else {
		goto L127
	}
L127:
	;
	v403 = F_strcmp(m, int32(_a_F_aclitemin_19), v374)
	mBase = m.M
	if v403 == int32(0) {
		v459 = int32(_a_F_aclitemin_20)
		goto L122
	} else {
		goto L128
	}
L128:
	;
	v408 = F_strcmp(m, int32(_a_F_aclitemin_21), v374)
	mBase = m.M
	if v408 == int32(0) {
		v459 = int32(_a_F_aclitemin_22)
		goto L122
	} else {
		goto L129
	}
L129:
	;
	v413 = F_strcmp(m, int32(_a_F_aclitemin_23), v374)
	mBase = m.M
	if v413 == int32(0) {
		v459 = int32(_a_F_aclitemin_24)
		goto L122
	} else {
		goto L130
	}
L130:
	;
	v418 = F_strcmp(m, int32(_a_F_aclitemin_25), v374)
	mBase = m.M
	if v418 == int32(0) {
		v459 = int32(_a_F_aclitemin_26)
		goto L122
	} else {
		goto L131
	}
L131:
	;
	v423 = F_strcmp(m, int32(_a_F_aclitemin_27), v374)
	mBase = m.M
	if v423 == int32(0) {
		v459 = int32(_a_F_aclitemin_28)
		goto L122
	} else {
		goto L132
	}
L132:
	;
	v428 = F_strcmp(m, int32(_a_F_aclitemin_29), v374)
	mBase = m.M
	if v428 == int32(0) {
		v459 = int32(_a_F_aclitemin_30)
		goto L122
	} else {
		goto L133
	}
L133:
	;
	v433 = F_strcmp(m, int32(_a_F_aclitemin_31), v374)
	mBase = m.M
	if v433 == int32(0) {
		v459 = int32(_a_F_aclitemin_32)
		goto L122
	} else {
		goto L134
	}
L134:
	;
	v438 = F_strcmp(m, int32(_a_F_aclitemin_33), v374)
	mBase = m.M
	if v438 == int32(0) {
		v459 = int32(_a_F_aclitemin_34)
		goto L122
	} else {
		goto L135
	}
L135:
	;
	v443 = F_strcmp(m, int32(_a_F_aclitemin_35), v374)
	mBase = m.M
	if v443 == int32(0) {
		v459 = int32(_a_F_aclitemin_36)
		goto L122
	} else {
		goto L136
	}
L136:
	;
	v448 = F_strcmp(m, int32(_a_F_aclitemin_37), v374)
	mBase = m.M
	if v448 == int32(0) {
		v459 = int32(_a_F_aclitemin_38)
		goto L122
	} else {
		goto L137
	}
L137:
	;
	v453 = F_strcmp(m, int32(_a_F_aclitemin_39), v374)
	mBase = m.M
	if v453 == int32(0) {
		v459 = int32(_a_F_aclitemin_40)
		goto L122
	} else {
		goto L138
	}
L138:
	;
	v457 = F_strcmp(m, int32(_a_F_aclitemin_41), v374)
	mBase = m.M
	if v457 != 0 {
		v461 = v375
		goto L121
	} else {
		goto L139
	}
L139:
	;
	v459 = int32(_a_F_aclitemin_42)
	goto L122
L140:
	;
	v471 = v469
	goto L116
L141:
	;
	v473 = F_errsave_start(m, v16)
	mBase = m.M
	v474 = m.ExcPending
	if v474 != 0 {
		goto L1
	} else {
		goto L142
	}
L142:
	;
	if v473 == int32(0) {
		goto L4
	} else {
		goto L143
	}
L143:
	;
	F_errcode(m, int32(67137668))
	mBase = m.M
	v479 = m.ExcPending
	if v479 != 0 {
		goto L1
	} else {
		goto L144
	}
L144:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v14)+16)) = v14 + int32(80)
	F_errmsg(m, int32(_a_F_aclitemin_43), v14+int32(16))
	mBase = m.M
	v487 = m.ExcPending
	if v487 != 0 {
		goto L1
	} else {
		goto L145
	}
L145:
	;
	F_errsave_finish(m, v16, int32(_a_F_aclitemin_2), int32(415), int32(_a_F_aclitemin_3))
	mBase = m.M
	v492 = m.ExcPending
	if v492 != 0 {
		goto L1
	} else {
		goto L146
	}
L146:
	;
	goto L4
L147:
	;
	v501 = F_errstart(m, int32(19), int32(0))
	mBase = m.M
	v502 = m.ExcPending
	if v502 != 0 {
		goto L1
	} else {
		goto L148
	}
L148:
	;
	if v501 == int32(0) {
		v521 = v202
		goto L102
	} else {
		goto L149
	}
L149:
	;
	F_errcode(m, int32(1792))
	mBase = m.M
	v507 = m.ExcPending
	if v507 != 0 {
		goto L1
	} else {
		goto L150
	}
L150:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v14)+32)) = int32(10)
	F_errmsg(m, int32(_a_F_aclitemin_45), v14+int32(32))
	mBase = m.M
	v514 = m.ExcPending
	if v514 != 0 {
		goto L1
	} else {
		goto L151
	}
L151:
	;
	F_errfinish(m, int32(_a_F_aclitemin_2), int32(424), int32(_a_F_aclitemin_3))
	mBase = m.M
	v519 = m.ExcPending
	if v519 != 0 {
		goto L1
	} else {
		goto L152
	}
L152:
	;
	v521 = v202
	goto L102
L153:
	;
	v534 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v527))))
	if base.B2i32(base.Ui32(v534-int32(9)) < base.Ui32(int32(5)))|base.B2i32(v534 == int32(32)) != 0 {
		goto L155
	} else {
		goto L156
	}
L155:
	;
	v527 = v527 + int32(1)
	goto L153
L156:
	;
	if v534 != 0 {
		goto L158
	} else {
		goto L159
	}
L158:
	;
	v544 = int64(0)
	v545 = F_errsave_start(m, v16)
	mBase = m.M
	v546 = m.ExcPending
	if v546 != 0 {
		goto L1
	} else {
		goto L161
	}
L159:
	;
	goto L160
L160:
	;
	v583 = base.I64_extend_i32_u(v19)
	goto L3
L161:
	;
	if v545 == int32(0) {
		v583 = v544
		goto L3
	} else {
		goto L162
	}
L162:
	;
	F_errcode(m, int32(33685634))
	mBase = m.M
	v551 = m.ExcPending
	if v551 != 0 {
		goto L1
	} else {
		goto L163
	}
L163:
	;
	F_errmsg(m, int32(_a_F_aclitemin_46), int32(0))
	mBase = m.M
	v555 = m.ExcPending
	if v555 != 0 {
		goto L1
	} else {
		goto L164
	}
L164:
	;
	F_errsave_finish(m, v16, int32(_a_F_aclitemin_2), int32(646), int32(_a_F_aclitemin_47))
	mBase = m.M
	v560 = m.ExcPending
	if v560 != 0 {
		goto L1
	} else {
		goto L165
	}
L165:
	;
	v583 = v544
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
				v18 = F_expression_tree_walker_impl(m, l0, int32(1119), l1)
				mBase = m.M
				v19 = m.ExcPending
				if v19 != 0 {
					return int32(0)
				} else {
					return v18
				}
			}
		} else {
			v18 = F_expression_tree_walker_impl(m, l0, int32(1119), l1)
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
	if v8 != int32(321) {
		goto L6
	} else {
		goto L7
	}
L4:
	;
	v119 = F_expression_tree_mutator_impl(m, l0, int32(1130), l1)
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
	v110 = F_query_tree_mutator_impl(m, l0, int32(1130), l1, int32(0))
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
	*(*int32)(unsafe.Add(mBase, uint32(v92))) = int32(321)
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
	var v53 int64
	_ = v53
	var v56 int32
	_ = v56
	var v60 int32
	_ = v60
	var v64 int32
	_ = v64
	var v65 int32
	_ = v65
	var v66 int32
	_ = v66
	var v75 int32
	_ = v75
	var v76 int32
	_ = v76
	var v78 int32
	_ = v78
	var v81 int32
	_ = v81
	var v82 int32
	_ = v82
	var v87 int32
	_ = v87
	var v94 int32
	_ = v94
	var v96 int32
	_ = v96
	var v98 int32
	_ = v98
	var v101 int32
	_ = v101
	var v103 int32
	_ = v103
	var v105 int32
	_ = v105
	var v112 int32
	_ = v112
	var v119 int32
	_ = v119
	var v120 int32
	_ = v120
	var v121 int32
	_ = v121
	var v123 int32
	_ = v123
	var v124 int32
	_ = v124
	var v125 int32
	_ = v125
	var v126 int32
	_ = v126
	var v127 int32
	_ = v127
	var v128 int32
	_ = v128
	var v129 int32
	_ = v129
	var v136 int32
	_ = v136
	var v140 int64
	_ = v140
	var v149 int32
	_ = v149
	var v152 int32
	_ = v152
	var v171 int32
	_ = v171
	var v172 int32
	_ = v172
	var v177 int32
	_ = v177
	var v193 int32
	_ = v193
	var v197 int32
	_ = v197
	var v198 int32
	_ = v198
	var v199 int32
	_ = v199
	var v200 int32
	_ = v200
	var v201 int32
	_ = v201
	var v210 int32
	_ = v210
	var v211 int32
	_ = v211
	var v213 int32
	_ = v213
	var v216 int32
	_ = v216
	var v217 int32
	_ = v217
	var v222 int32
	_ = v222
	var v229 int32
	_ = v229
	var v231 int32
	_ = v231
	var v233 int32
	_ = v233
	var v236 int32
	_ = v236
	var v238 int32
	_ = v238
	var v240 int32
	_ = v240
	var v247 int32
	_ = v247
	var v254 int32
	_ = v254
	var v257 int32
	_ = v257
	var v260 int32
	_ = v260
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
	var v269 int32
	_ = v269
	var v278 int32
	_ = v278
	var v279 int32
	_ = v279
	var v281 int32
	_ = v281
	var v284 int32
	_ = v284
	var v285 int32
	_ = v285
	var v290 int32
	_ = v290
	var v297 int32
	_ = v297
	var v299 int32
	_ = v299
	var v301 int32
	_ = v301
	var v304 int32
	_ = v304
	var v306 int32
	_ = v306
	var v308 int32
	_ = v308
	var v315 int32
	_ = v315
	var v322 int32
	_ = v322
	var v325 int32
	_ = v325
	var v326 int32
	_ = v326
	var v335 int32
	_ = v335
	var v336 int32
	_ = v336
	var v338 int32
	_ = v338
	var v341 int32
	_ = v341
	var v342 int32
	_ = v342
	var v347 int32
	_ = v347
	var v354 int32
	_ = v354
	var v356 int32
	_ = v356
	var v358 int32
	_ = v358
	var v361 int32
	_ = v361
	var v363 int32
	_ = v363
	var v365 int32
	_ = v365
	var v372 int32
	_ = v372
	var v379 int32
	_ = v379
	var v382 int32
	_ = v382
	var v384 int32
	_ = v384
	var v385 int32
	_ = v385
	var v394 int32
	_ = v394
	var v395 int32
	_ = v395
	var v397 int32
	_ = v397
	var v400 int32
	_ = v400
	var v401 int32
	_ = v401
	var v406 int32
	_ = v406
	var v413 int32
	_ = v413
	var v415 int32
	_ = v415
	var v417 int32
	_ = v417
	var v420 int32
	_ = v420
	var v422 int32
	_ = v422
	var v424 int32
	_ = v424
	var v431 int32
	_ = v431
	var v438 int32
	_ = v438
	var v442 int32
	_ = v442
	var v443 int32
	_ = v443
	var v452 int32
	_ = v452
	var v453 int32
	_ = v453
	var v455 int32
	_ = v455
	var v458 int32
	_ = v458
	var v459 int32
	_ = v459
	var v464 int32
	_ = v464
	var v471 int32
	_ = v471
	var v473 int32
	_ = v473
	var v475 int32
	_ = v475
	var v478 int32
	_ = v478
	var v480 int32
	_ = v480
	var v482 int32
	_ = v482
	var v489 int32
	_ = v489
	var v496 int32
	_ = v496
	var v500 int32
	_ = v500
	var v502 int32
	_ = v502
	var v503 int32
	_ = v503
	var v504 int32
	_ = v504
	var v505 int32
	_ = v505
	var v507 int32
	_ = v507
	var v508 int32
	_ = v508
	var v524 int32
	_ = v524
	var v544 int32
	_ = v544
	var v579 int32
	_ = v579
	var v580 int32
	_ = v580
	var v596 int32
	_ = v596
	var v616 int32
	_ = v616
	var v651 int32
	_ = v651
	var v652 int32
	_ = v652
	var v654 int32
	_ = v654
	var v655 int32
	_ = v655
	var v657 int32
	_ = v657
	var v658 int32
	_ = v658
	var v671 int32
	_ = v671
	var v672 int32
	_ = v672
	var v694 int32
	_ = v694
	var v695 int32
	_ = v695
	var v712 int32
	_ = v712
	var v718 int32
	_ = v718
	var v740 int32
	_ = v740
	var v755 int32
	_ = v755
	var v761 int32
	_ = v761
	var v777 int32
	_ = v777
	var v785 int32
	_ = v785
	var v797 int32
	_ = v797
	var v814 int32
	_ = v814
	var v824 int32
	_ = v824
	var v825 int32
	_ = v825
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
	var v854 int32
	_ = v854
	var v876 int32
	_ = v876
	var v880 int32
	_ = v880
	var v881 int32
	_ = v881
	var v882 int32
	_ = v882
	var v883 int32
	_ = v883
	var v884 int32
	_ = v884
	var v893 int32
	_ = v893
	var v894 int32
	_ = v894
	var v896 int32
	_ = v896
	var v899 int32
	_ = v899
	var v900 int32
	_ = v900
	var v905 int32
	_ = v905
	var v912 int32
	_ = v912
	var v914 int32
	_ = v914
	var v916 int32
	_ = v916
	var v919 int32
	_ = v919
	var v921 int32
	_ = v921
	var v923 int32
	_ = v923
	var v930 int32
	_ = v930
	var v937 int32
	_ = v937
	var v940 int32
	_ = v940
	var v941 int32
	_ = v941
	var v942 int32
	_ = v942
	var v944 int32
	_ = v944
	var v945 int32
	_ = v945
	var v952 int32
	_ = v952
	var v982 int32
	_ = v982
	var v986 int32
	_ = v986
	var v987 float64
	_ = v987
	var v988 int32
	_ = v988
	var v990 int32
	_ = v990
	var v991 int32
	_ = v991
	var v992 int32
	_ = v992
	var v993 int64
	_ = v993
	var v1009 int32
	_ = v1009
	var v1011 float64
	_ = v1011
	var v1012 int32
	_ = v1012
	var v1014 int32
	_ = v1014
	var v1017 float64
	_ = v1017
	var v1018 float64
	_ = v1018
	var v1020 float64
	_ = v1020
	var v1023 float64
	_ = v1023
	var v1026 float64
	_ = v1026
	var v1066 int32
	_ = v1066
	var v1069 int32
	_ = v1069
	var v1079 int32
	_ = v1079
	var v1107 int32
	_ = v1107
	var v1111 int32
	_ = v1111
	var v1112 int32
	_ = v1112
	var v1113 int32
	_ = v1113
	var v1123 int32
	_ = v1123
	var v1124 int32
	_ = v1124
	var v1126 int32
	_ = v1126
	var v1129 int32
	_ = v1129
	var v1130 int32
	_ = v1130
	var v1135 int32
	_ = v1135
	var v1142 int32
	_ = v1142
	var v1144 int32
	_ = v1144
	var v1146 int32
	_ = v1146
	var v1147 int32
	_ = v1147
	var v1149 int32
	_ = v1149
	var v1151 int32
	_ = v1151
	var v1158 int32
	_ = v1158
	var v1161 int32
	_ = v1161
	var v1162 int32
	_ = v1162
	var v1172 int32
	_ = v1172
	var v1173 int32
	_ = v1173
	var v1175 int32
	_ = v1175
	var v1178 int32
	_ = v1178
	var v1179 int32
	_ = v1179
	var v1184 int32
	_ = v1184
	var v1191 int32
	_ = v1191
	var v1193 int32
	_ = v1193
	var v1195 int32
	_ = v1195
	var v1196 int32
	_ = v1196
	var v1198 int32
	_ = v1198
	var v1200 int32
	_ = v1200
	var v1207 int32
	_ = v1207
	var v1208 int32
	_ = v1208
	var v1209 int32
	_ = v1209
	var v1210 int32
	_ = v1210
	var v1211 int32
	_ = v1211
	var v1212 int32
	_ = v1212
	var v1213 int32
	_ = v1213
	var v1214 int32
	_ = v1214
	var v1216 int32
	_ = v1216
	var v1219 int32
	_ = v1219
	var v1220 int32
	_ = v1220
	var v1230 int32
	_ = v1230
	var v1231 int32
	_ = v1231
	var v1233 int32
	_ = v1233
	var v1236 int32
	_ = v1236
	var v1237 int32
	_ = v1237
	var v1242 int32
	_ = v1242
	var v1249 int32
	_ = v1249
	var v1251 int32
	_ = v1251
	var v1253 int32
	_ = v1253
	var v1254 int32
	_ = v1254
	var v1256 int32
	_ = v1256
	var v1258 int32
	_ = v1258
	var v1265 int32
	_ = v1265
	var v1268 int32
	_ = v1268
	var v1269 int32
	_ = v1269
	var v1279 int32
	_ = v1279
	var v1280 int32
	_ = v1280
	var v1282 int32
	_ = v1282
	var v1285 int32
	_ = v1285
	var v1286 int32
	_ = v1286
	var v1291 int32
	_ = v1291
	var v1298 int32
	_ = v1298
	var v1300 int32
	_ = v1300
	var v1302 int32
	_ = v1302
	var v1303 int32
	_ = v1303
	var v1305 int32
	_ = v1305
	var v1307 int32
	_ = v1307
	var v1314 int32
	_ = v1314
	var v1315 int32
	_ = v1315
	var v1316 int32
	_ = v1316
	var v1317 int32
	_ = v1317
	var v1318 int32
	_ = v1318
	var v1319 int32
	_ = v1319
	var v1320 int32
	_ = v1320
	var v1321 int32
	_ = v1321
	var v1324 int32
	_ = v1324
	var v1325 int32
	_ = v1325
	var v1361 int32
	_ = v1361
	var v1362 int32
	_ = v1362
	var v1363 int32
	_ = v1363
	var v1364 int32
	_ = v1364
	var v1370 int32
	_ = v1370
	var v1373 int32
	_ = v1373
	var v1374 int32
	_ = v1374
	var v1375 int32
	_ = v1375
	var v1378 int32
	_ = v1378
	var v1379 int32
	_ = v1379
	var v1380 int32
	_ = v1380
	var v1390 int32
	_ = v1390
	var v1391 int32
	_ = v1391
	var v1393 int32
	_ = v1393
	var v1396 int32
	_ = v1396
	var v1397 int32
	_ = v1397
	var v1402 int32
	_ = v1402
	var v1409 int32
	_ = v1409
	var v1411 int32
	_ = v1411
	var v1413 int32
	_ = v1413
	var v1414 int32
	_ = v1414
	var v1416 int32
	_ = v1416
	var v1418 int32
	_ = v1418
	var v1425 int32
	_ = v1425
	var v1426 int32
	_ = v1426
	var v1429 int32
	_ = v1429
	var v1430 int32
	_ = v1430
	var v1431 int32
	_ = v1431
	var v1441 int32
	_ = v1441
	var v1442 int32
	_ = v1442
	var v1444 int32
	_ = v1444
	var v1447 int32
	_ = v1447
	var v1448 int32
	_ = v1448
	var v1453 int32
	_ = v1453
	var v1460 int32
	_ = v1460
	var v1462 int32
	_ = v1462
	var v1464 int32
	_ = v1464
	var v1465 int32
	_ = v1465
	var v1467 int32
	_ = v1467
	var v1469 int32
	_ = v1469
	var v1476 int32
	_ = v1476
	var v1478 int32
	_ = v1478
	var v1481 int32
	_ = v1481
	var v1482 int32
	_ = v1482
	var v1483 int32
	_ = v1483
	var v1493 int32
	_ = v1493
	var v1494 int32
	_ = v1494
	var v1496 int32
	_ = v1496
	var v1499 int32
	_ = v1499
	var v1500 int32
	_ = v1500
	var v1505 int32
	_ = v1505
	var v1512 int32
	_ = v1512
	var v1514 int32
	_ = v1514
	var v1516 int32
	_ = v1516
	var v1517 int32
	_ = v1517
	var v1519 int32
	_ = v1519
	var v1521 int32
	_ = v1521
	var v1528 int32
	_ = v1528
	var v1529 int32
	_ = v1529
	var v1532 int32
	_ = v1532
	var v1533 int32
	_ = v1533
	var v1534 int32
	_ = v1534
	var v1544 int32
	_ = v1544
	var v1545 int32
	_ = v1545
	var v1547 int32
	_ = v1547
	var v1550 int32
	_ = v1550
	var v1551 int32
	_ = v1551
	var v1556 int32
	_ = v1556
	var v1563 int32
	_ = v1563
	var v1565 int32
	_ = v1565
	var v1567 int32
	_ = v1567
	var v1568 int32
	_ = v1568
	var v1570 int32
	_ = v1570
	var v1572 int32
	_ = v1572
	var v1579 int32
	_ = v1579
	var v1581 int32
	_ = v1581
	var v1582 int32
	_ = v1582
	var v1593 int32
	_ = v1593
	var v1597 int32
	_ = v1597
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
	var v1609 int32
	_ = v1609
	var v1612 int32
	_ = v1612
	var v1615 int32
	_ = v1615
	var v1616 int32
	_ = v1616
	var v1618 int32
	_ = v1618
	var v1626 int32
	_ = v1626
	var v1627 int32
	_ = v1627
	var v1630 int32
	_ = v1630
	var v1633 int32
	_ = v1633
	var v1638 int32
	_ = v1638
	var v1648 int32
	_ = v1648
	var v1652 int32
	_ = v1652
	var v1653 int32
	_ = v1653
	var v1654 int32
	_ = v1654
	var v1655 int32
	_ = v1655
	var v1658 int32
	_ = v1658
	var v1662 int32
	_ = v1662
	var v1663 int32
	_ = v1663
	var v1664 int32
	_ = v1664
	var v1665 int32
	_ = v1665
	var v1666 int32
	_ = v1666
	var v1667 int32
	_ = v1667
	var v1676 int32
	_ = v1676
	var v1687 int32
	_ = v1687
	var v1704 int32
	_ = v1704
	var v1708 int32
	_ = v1708
	var v1709 int32
	_ = v1709
	var v1710 int32
	_ = v1710
	var v1725 int32
	_ = v1725
	var v1746 int32
	_ = v1746
	var v1781 int32
	_ = v1781
	var v1782 int32
	_ = v1782
	var v1797 int32
	_ = v1797
	var v1818 int32
	_ = v1818
	var v1855 int32
	_ = v1855
	var v1856 int32
	_ = v1856
	var v1858 int32
	_ = v1858
	var v1859 int32
	_ = v1859
	var v1876 int32
	_ = v1876
	var v1899 int32
	_ = v1899
	var v1902 int32
	_ = v1902
	var v1938 int32
	_ = v1938
	var v1942 int32
	_ = v1942
	var v1944 int32
	_ = v1944
	var v1954 int32
	_ = v1954
	var v1961 int32
	_ = v1961
	var v1981 int32
	_ = v1981
	var v1985 int32
	_ = v1985
	var v1986 int32
	_ = v1986
	var v1989 int32
	_ = v1989
	var v1990 int32
	_ = v1990
	var v1991 int32
	_ = v1991
	var v2001 int32
	_ = v2001
	var v2002 int32
	_ = v2002
	var v2004 int32
	_ = v2004
	var v2007 int32
	_ = v2007
	var v2008 int32
	_ = v2008
	var v2013 int32
	_ = v2013
	var v2020 int32
	_ = v2020
	var v2022 int32
	_ = v2022
	var v2024 int32
	_ = v2024
	var v2025 int32
	_ = v2025
	var v2027 int32
	_ = v2027
	var v2029 int32
	_ = v2029
	var v2036 int32
	_ = v2036
	var v2040 int32
	_ = v2040
	var v2042 int32
	_ = v2042
	var v2043 int32
	_ = v2043
	var v2052 int32
	_ = v2052
	var v2080 int32
	_ = v2080
	var v2093 int32
	_ = v2093
	var v2122 int32
	_ = v2122
	var v2123 int32
	_ = v2123
	var v2131 int32
	_ = v2131
	var v2159 int32
	_ = v2159
	var v2162 int32
	_ = v2162
	var v2165 int32
	_ = v2165
	var v2185 int32
	_ = v2185
	var v2206 int32
	_ = v2206
	var v2207 int32
	_ = v2207
	var v2223 int32
	_ = v2223
	var v2246 int32
	_ = v2246
	var v2249 int32
	_ = v2249
	var v2252 int32
	_ = v2252
	var v2288 int32
	_ = v2288
	var v2289 int32
	_ = v2289
	var v2290 int32
	_ = v2290
	var v2293 int32
	_ = v2293
	var v2296 int32
	_ = v2296
	var v2317 int32
	_ = v2317
	var v2333 int32
	_ = v2333
	var v2337 int32
	_ = v2337
	var v2338 int32
	_ = v2338
	var v2354 int32
	_ = v2354
	var v2375 int32
	_ = v2375
	var v2377 int32
	_ = v2377
	var v2383 int32
	_ = v2383
	var v2420 int32
	_ = v2420
	var v2421 int32
	_ = v2421
	var v2424 int32
	_ = v2424
	var v2425 int32
	_ = v2425
	var v2442 int32
	_ = v2442
	var v2461 int32
	_ = v2461
	var v2464 int32
	_ = v2464
	var v2465 int32
	_ = v2465
	var v2466 int32
	_ = v2466
	var v2488 int32
	_ = v2488
	var v2506 int32
	_ = v2506
	var v2508 int32
	_ = v2508
	var v2509 int32
	_ = v2509
	var v2521 int32
	_ = v2521
	var v2529 int32
	_ = v2529
	var v2533 int32
	_ = v2533
	var v2534 int32
	_ = v2534
	var v2550 int32
	_ = v2550
	var v2551 int32
	_ = v2551
	var v2554 int32
	_ = v2554
	var v2556 int32
	_ = v2556
	var v2560 int32
	_ = v2560
	var v2562 int32
	_ = v2562
	var v2566 int32
	_ = v2566
	var v2570 int32
	_ = v2570
	var v2571 int32
	_ = v2571
	var v2572 int32
	_ = v2572
	var v2573 int32
	_ = v2573
	var v2574 int32
	_ = v2574
	var v2575 int32
	_ = v2575
	var v2576 int32
	_ = v2576
	var v2577 int32
	_ = v2577
	var v2578 int32
	_ = v2578
	var v2579 int32
	_ = v2579
	var v2580 int32
	_ = v2580
	var v2581 int32
	_ = v2581
	var v2582 int32
	_ = v2582
	var v2583 int32
	_ = v2583
	var v2584 int32
	_ = v2584
	var v2586 int32
	_ = v2586
	var v2597 int32
	_ = v2597
	var v2604 int32
	_ = v2604
	var v2609 int32
	_ = v2609
	var v2630 int32
	_ = v2630
	var v2638 int32
	_ = v2638
	var v2643 int32
	_ = v2643
	var v2646 int32
	_ = v2646
	var v2661 int32
	_ = v2661
	var v2662 int32
	_ = v2662
	var v2663 int32
	_ = v2663
	var v2664 int32
	_ = v2664
	var v2665 int32
	_ = v2665
	var v2668 int32
	_ = v2668
	var v2676 int32
	_ = v2676
	var v2689 int32
	_ = v2689
	var v2707 int32
	_ = v2707
	var v2709 int32
	_ = v2709
	var v2713 int32
	_ = v2713
	var v2714 int32
	_ = v2714
	var v2715 int32
	_ = v2715
	var v2718 int32
	_ = v2718
	var v2719 int32
	_ = v2719
	var v2720 int32
	_ = v2720
	var v2721 int32
	_ = v2721
	var v2738 int32
	_ = v2738
	var v2757 int32
	_ = v2757
	var v2759 int32
	_ = v2759
	var v2776 int32
	_ = v2776
	var v2796 int32
	_ = v2796
	var v2799 int32
	_ = v2799
	var v2800 int32
	_ = v2800
	var v2801 int32
	_ = v2801
	var v2802 int32
	_ = v2802
	var v2803 int32
	_ = v2803
	var v2804 int32
	_ = v2804
	var v2805 int32
	_ = v2805
	var v2807 int32
	_ = v2807
	var v2810 int32
	_ = v2810
	var v2811 int32
	_ = v2811
	var v2815 int32
	_ = v2815
	var v2817 int32
	_ = v2817
	var v2818 int32
	_ = v2818
	var v2828 int32
	_ = v2828
	var v2856 int32
	_ = v2856
	var v2860 int32
	_ = v2860
	var v2861 int32
	_ = v2861
	var v2862 int32
	_ = v2862
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
	var v2875 int32
	_ = v2875
	var v2878 int32
	_ = v2878
	var v2880 int32
	_ = v2880
	var v2882 int32
	_ = v2882
	var v2883 int32
	_ = v2883
	var v2921 int32
	_ = v2921
	var v2923 int32
	_ = v2923
	var v2934 int32
	_ = v2934
	var v2935 int32
	_ = v2935
	var v2936 int32
	_ = v2936
	var v2939 int32
	_ = v2939
	var v2940 int32
	_ = v2940
	var v2941 int32
	_ = v2941
	var v2951 int32
	_ = v2951
	var v2952 int32
	_ = v2952
	var v2954 int32
	_ = v2954
	var v2957 int32
	_ = v2957
	var v2958 int32
	_ = v2958
	var v2963 int32
	_ = v2963
	var v2970 int32
	_ = v2970
	var v2972 int32
	_ = v2972
	var v2974 int32
	_ = v2974
	var v2975 int32
	_ = v2975
	var v2977 int32
	_ = v2977
	var v2979 int32
	_ = v2979
	var v2986 int32
	_ = v2986
	var v2989 int32
	_ = v2989
	var v2992 int32
	_ = v2992
	var v2993 int32
	_ = v2993
	var v2994 int32
	_ = v2994
	var v3004 int32
	_ = v3004
	var v3005 int32
	_ = v3005
	var v3007 int32
	_ = v3007
	var v3010 int32
	_ = v3010
	var v3011 int32
	_ = v3011
	var v3016 int32
	_ = v3016
	var v3023 int32
	_ = v3023
	var v3025 int32
	_ = v3025
	var v3027 int32
	_ = v3027
	var v3028 int32
	_ = v3028
	var v3030 int32
	_ = v3030
	var v3032 int32
	_ = v3032
	var v3039 int32
	_ = v3039
	var v3043 int32
	_ = v3043
	var v3044 int32
	_ = v3044
	var v3045 int32
	_ = v3045
	var v3048 int32
	_ = v3048
	var v3049 int32
	_ = v3049
	var v3050 int32
	_ = v3050
	var v3064 int32
	_ = v3064
	var v3065 int32
	_ = v3065
	var v3067 int32
	_ = v3067
	var v3070 int32
	_ = v3070
	var v3071 int32
	_ = v3071
	var v3076 int32
	_ = v3076
	var v3084 int32
	_ = v3084
	var v3086 int32
	_ = v3086
	var v3088 int32
	_ = v3088
	var v3089 int32
	_ = v3089
	var v3092 int32
	_ = v3092
	var v3096 int32
	_ = v3096
	var v3102 int32
	_ = v3102
	var v3103 int32
	_ = v3103
	var v3108 int64
	_ = v3108
	var v3109 int64
	_ = v3109
	var v3114 int32
	_ = v3114
	var v3116 int32
	_ = v3116
	var v3122 int32
	_ = v3122
	var v3123 int32
	_ = v3123
	var v3124 int32
	_ = v3124
	var v3125 int32
	_ = v3125
	var v3128 int32
	_ = v3128
	var v3137 int32
	_ = v3137
	var v3166 int32
	_ = v3166
	var v3170 int32
	_ = v3170
	var v3171 int32
	_ = v3171
	var v3174 int32
	_ = v3174
	var v3175 int32
	_ = v3175
	var v3176 int32
	_ = v3176
	var v3186 int32
	_ = v3186
	var v3187 int32
	_ = v3187
	var v3189 int32
	_ = v3189
	var v3192 int32
	_ = v3192
	var v3193 int32
	_ = v3193
	var v3198 int32
	_ = v3198
	var v3205 int32
	_ = v3205
	var v3207 int32
	_ = v3207
	var v3209 int32
	_ = v3209
	var v3210 int32
	_ = v3210
	var v3212 int32
	_ = v3212
	var v3214 int32
	_ = v3214
	var v3221 int32
	_ = v3221
	var v3222 int32
	_ = v3222
	var v3225 int32
	_ = v3225
	var v3226 int32
	_ = v3226
	var v3227 int32
	_ = v3227
	var v3237 int32
	_ = v3237
	var v3238 int32
	_ = v3238
	var v3240 int32
	_ = v3240
	var v3243 int32
	_ = v3243
	var v3244 int32
	_ = v3244
	var v3249 int32
	_ = v3249
	var v3256 int32
	_ = v3256
	var v3258 int32
	_ = v3258
	var v3260 int32
	_ = v3260
	var v3261 int32
	_ = v3261
	var v3263 int32
	_ = v3263
	var v3265 int32
	_ = v3265
	var v3272 int32
	_ = v3272
	var v3274 int32
	_ = v3274
	var v3275 int32
	_ = v3275
	var v3276 int32
	_ = v3276
	var v3279 int32
	_ = v3279
	var v3282 int32
	_ = v3282
	var v3283 int32
	_ = v3283
	var v3292 int32
	_ = v3292
	var v3320 int32
	_ = v3320
	var v3324 int32
	_ = v3324
	var v3327 int32
	_ = v3327
	var v3329 int32
	_ = v3329
	var v3330 int32
	_ = v3330
	var v3331 int32
	_ = v3331
	var v3334 int32
	_ = v3334
	var v3336 int32
	_ = v3336
	var v3337 int32
	_ = v3337
	var v3379 int32
	_ = v3379
	var v3420 int32
	_ = v3420
	var v3456 int32
	_ = v3456
	var v3457 int32
	_ = v3457
	var v3493 int32
	_ = v3493
	var v3494 int32
	_ = v3494
	var v3503 int32
	_ = v3503
	var v3505 int32
	_ = v3505
	var v3508 int32
	_ = v3508
	var v3509 int32
	_ = v3509
	var v3512 int32
	_ = v3512
	var v3513 int32
	_ = v3513
	var v3518 int32
	_ = v3518
	var v3519 int32
	_ = v3519
	var v3522 int32
	_ = v3522
	var v3525 int32
	_ = v3525
	var v3526 int32
	_ = v3526
	var v3527 int32
	_ = v3527
	var v3537 int32
	_ = v3537
	var v3538 int32
	_ = v3538
	var v3540 int32
	_ = v3540
	var v3543 int32
	_ = v3543
	var v3544 int32
	_ = v3544
	var v3549 int32
	_ = v3549
	var v3556 int32
	_ = v3556
	var v3558 int32
	_ = v3558
	var v3560 int32
	_ = v3560
	var v3561 int32
	_ = v3561
	var v3563 int32
	_ = v3563
	var v3565 int32
	_ = v3565
	var v3572 int32
	_ = v3572
	var v3573 int32
	_ = v3573
	var v3576 int32
	_ = v3576
	var v3577 int32
	_ = v3577
	var v3578 int32
	_ = v3578
	var v3588 int32
	_ = v3588
	var v3589 int32
	_ = v3589
	var v3591 int32
	_ = v3591
	var v3594 int32
	_ = v3594
	var v3595 int32
	_ = v3595
	var v3600 int32
	_ = v3600
	var v3607 int32
	_ = v3607
	var v3609 int32
	_ = v3609
	var v3611 int32
	_ = v3611
	var v3612 int32
	_ = v3612
	var v3614 int32
	_ = v3614
	var v3616 int32
	_ = v3616
	var v3623 int32
	_ = v3623
	var v3625 int32
	_ = v3625
	var v3627 int32
	_ = v3627
	var v3633 int32
	_ = v3633
	var v3634 int32
	_ = v3634
	var v3637 int32
	_ = v3637
	var v3638 int32
	_ = v3638
	var v3641 int32
	_ = v3641
	var v3666 int32
	_ = v3666
	var v3678 int32
	_ = v3678
	var v3682 int32
	_ = v3682
	var v3683 int32
	_ = v3683
	var v3684 int32
	_ = v3684
	var v3685 int32
	_ = v3685
	var v3686 int32
	_ = v3686
	var v3689 int32
	_ = v3689
	var v3690 int32
	_ = v3690
	var v3698 int32
	_ = v3698
	var v3727 int32
	_ = v3727
	var v3731 int32
	_ = v3731
	var v3732 int32
	_ = v3732
	var v3737 int32
	_ = v3737
	var v3738 int32
	_ = v3738
	var v3739 int32
	_ = v3739
	var v3744 int32
	_ = v3744
	var v3747 int32
	_ = v3747
	var v3748 int32
	_ = v3748
	var v3786 int32
	_ = v3786
	var v3788 int32
	_ = v3788
	var v3789 int32
	_ = v3789
	var v3859 int32
	_ = v3859
	var v3860 int32
	_ = v3860
	var v3865 int32
	_ = v3865
	var v3868 int32
	_ = v3868
	var v3871 int32
	_ = v3871
	var v3872 int32
	_ = v3872
	var v3874 int32
	_ = v3874
	var v3882 int32
	_ = v3882
	var v3883 int32
	_ = v3883
	var v3886 int32
	_ = v3886
	var v3889 int32
	_ = v3889
	var v3894 int32
	_ = v3894
	var v3904 int32
	_ = v3904
	var v3909 int32
	_ = v3909
	var v3912 int32
	_ = v3912
	var v3915 int32
	_ = v3915
	var v3924 int32
	_ = v3924
	var v3953 int32
	_ = v3953
	var v3957 int32
	_ = v3957
	var v3959 int32
	_ = v3959
	var v3960 int32
	_ = v3960
	var v3961 int32
	_ = v3961
	var v3964 int32
	_ = v3964
	var v3966 int32
	_ = v3966
	var v3967 int32
	_ = v3967
	var v4039 int64
	_ = v4039
	var v4044 int32
	_ = v4044
	var v4047 int32
	_ = v4047
	var v4054 int32
	_ = v4054
	var v4062 int32
	_ = v4062
	var v4069 int32
	_ = v4069
	var v4090 int32
	_ = v4090
	var v4094 int32
	_ = v4094
	var v4095 int32
	_ = v4095
	var v4096 int32
	_ = v4096
	var v4097 int32
	_ = v4097
	var v4098 int32
	_ = v4098
	var v4107 int32
	_ = v4107
	var v4108 int32
	_ = v4108
	var v4110 int32
	_ = v4110
	var v4113 int32
	_ = v4113
	var v4114 int32
	_ = v4114
	var v4119 int32
	_ = v4119
	var v4126 int32
	_ = v4126
	var v4128 int32
	_ = v4128
	var v4130 int32
	_ = v4130
	var v4133 int32
	_ = v4133
	var v4135 int32
	_ = v4135
	var v4137 int32
	_ = v4137
	var v4144 int32
	_ = v4144
	var v4151 int32
	_ = v4151
	var v4154 int32
	_ = v4154
	var v4157 int32
	_ = v4157
	var v4160 int32
	_ = v4160
	var v4161 int32
	_ = v4161
	var v4162 int32
	_ = v4162
	var v4163 int32
	_ = v4163
	var v4172 int32
	_ = v4172
	var v4173 int32
	_ = v4173
	var v4175 int32
	_ = v4175
	var v4178 int32
	_ = v4178
	var v4179 int32
	_ = v4179
	var v4184 int32
	_ = v4184
	var v4191 int32
	_ = v4191
	var v4193 int32
	_ = v4193
	var v4195 int32
	_ = v4195
	var v4198 int32
	_ = v4198
	var v4200 int32
	_ = v4200
	var v4202 int32
	_ = v4202
	var v4209 int32
	_ = v4209
	var v4216 int32
	_ = v4216
	var v4217 int32
	_ = v4217
	var v4218 int32
	_ = v4218
	var v4227 int32
	_ = v4227
	var v4228 int32
	_ = v4228
	var v4230 int32
	_ = v4230
	var v4233 int32
	_ = v4233
	var v4234 int32
	_ = v4234
	var v4239 int32
	_ = v4239
	var v4246 int32
	_ = v4246
	var v4248 int32
	_ = v4248
	var v4250 int32
	_ = v4250
	var v4253 int32
	_ = v4253
	var v4255 int32
	_ = v4255
	var v4257 int32
	_ = v4257
	var v4264 int32
	_ = v4264
	var v4271 int32
	_ = v4271
	var v4272 int32
	_ = v4272
	var v4273 int32
	_ = v4273
	var v4282 int32
	_ = v4282
	var v4283 int32
	_ = v4283
	var v4285 int32
	_ = v4285
	var v4288 int32
	_ = v4288
	var v4289 int32
	_ = v4289
	var v4294 int32
	_ = v4294
	var v4301 int32
	_ = v4301
	var v4303 int32
	_ = v4303
	var v4305 int32
	_ = v4305
	var v4308 int32
	_ = v4308
	var v4310 int32
	_ = v4310
	var v4312 int32
	_ = v4312
	var v4319 int32
	_ = v4319
	var v4326 int32
	_ = v4326
	var v4329 int32
	_ = v4329
	var v4330 int32
	_ = v4330
	var v4339 int32
	_ = v4339
	var v4340 int32
	_ = v4340
	var v4342 int32
	_ = v4342
	var v4345 int32
	_ = v4345
	var v4346 int32
	_ = v4346
	var v4351 int32
	_ = v4351
	var v4358 int32
	_ = v4358
	var v4360 int32
	_ = v4360
	var v4362 int32
	_ = v4362
	var v4365 int32
	_ = v4365
	var v4367 int32
	_ = v4367
	var v4369 int32
	_ = v4369
	var v4376 int32
	_ = v4376
	var v4383 int32
	_ = v4383
	var v4386 int32
	_ = v4386
	var v4388 int32
	_ = v4388
	var v4389 int32
	_ = v4389
	var v4390 int32
	_ = v4390
	var v4391 int32
	_ = v4391
	var v4394 int32
	_ = v4394
	var v4396 int32
	_ = v4396
	var v4397 int32
	_ = v4397
	var v4399 int32
	_ = v4399
	var v4402 int32
	_ = v4402
	var v4403 int32
	_ = v4403
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
	var v4415 int32
	_ = v4415
	var v4425 int32
	_ = v4425
	var v4426 int32
	_ = v4426
	var v4428 int32
	_ = v4428
	var v4431 int32
	_ = v4431
	var v4432 int32
	_ = v4432
	var v4437 int32
	_ = v4437
	var v4444 int32
	_ = v4444
	var v4446 int32
	_ = v4446
	var v4448 int32
	_ = v4448
	var v4449 int32
	_ = v4449
	var v4451 int32
	_ = v4451
	var v4453 int32
	_ = v4453
	var v4460 int32
	_ = v4460
	var v4461 int32
	_ = v4461
	var v4464 int32
	_ = v4464
	var v4465 int32
	_ = v4465
	var v4466 int32
	_ = v4466
	var v4476 int32
	_ = v4476
	var v4477 int32
	_ = v4477
	var v4479 int32
	_ = v4479
	var v4482 int32
	_ = v4482
	var v4483 int32
	_ = v4483
	var v4488 int32
	_ = v4488
	var v4495 int32
	_ = v4495
	var v4497 int32
	_ = v4497
	var v4499 int32
	_ = v4499
	var v4500 int32
	_ = v4500
	var v4502 int32
	_ = v4502
	var v4504 int32
	_ = v4504
	var v4511 int32
	_ = v4511
	var v4513 int32
	_ = v4513
	var v4516 int32
	_ = v4516
	var v4517 int32
	_ = v4517
	var v4518 int32
	_ = v4518
	var v4528 int32
	_ = v4528
	var v4529 int32
	_ = v4529
	var v4531 int32
	_ = v4531
	var v4534 int32
	_ = v4534
	var v4535 int32
	_ = v4535
	var v4540 int32
	_ = v4540
	var v4547 int32
	_ = v4547
	var v4549 int32
	_ = v4549
	var v4551 int32
	_ = v4551
	var v4552 int32
	_ = v4552
	var v4554 int32
	_ = v4554
	var v4556 int32
	_ = v4556
	var v4563 int32
	_ = v4563
	var v4564 int32
	_ = v4564
	var v4567 int32
	_ = v4567
	var v4568 int32
	_ = v4568
	var v4569 int32
	_ = v4569
	var v4579 int32
	_ = v4579
	var v4580 int32
	_ = v4580
	var v4582 int32
	_ = v4582
	var v4585 int32
	_ = v4585
	var v4586 int32
	_ = v4586
	var v4591 int32
	_ = v4591
	var v4598 int32
	_ = v4598
	var v4600 int32
	_ = v4600
	var v4602 int32
	_ = v4602
	var v4603 int32
	_ = v4603
	var v4605 int32
	_ = v4605
	var v4607 int32
	_ = v4607
	var v4614 int32
	_ = v4614
	var v4619 int32
	_ = v4619
	var v4620 int32
	_ = v4620
	var v4623 int32
	_ = v4623
	var v4649 int32
	_ = v4649
	var v4661 int32
	_ = v4661
	var v4665 int32
	_ = v4665
	var v4666 int32
	_ = v4666
	var v4669 int32
	_ = v4669
	var v4670 int32
	_ = v4670
	var v4671 int32
	_ = v4671
	var v4681 int32
	_ = v4681
	var v4682 int32
	_ = v4682
	var v4684 int32
	_ = v4684
	var v4687 int32
	_ = v4687
	var v4688 int32
	_ = v4688
	var v4693 int32
	_ = v4693
	var v4700 int32
	_ = v4700
	var v4702 int32
	_ = v4702
	var v4704 int32
	_ = v4704
	var v4705 int32
	_ = v4705
	var v4707 int32
	_ = v4707
	var v4709 int32
	_ = v4709
	var v4716 int32
	_ = v4716
	var v4717 int32
	_ = v4717
	var v4720 int32
	_ = v4720
	var v4721 int32
	_ = v4721
	var v4722 int32
	_ = v4722
	var v4732 int32
	_ = v4732
	var v4733 int32
	_ = v4733
	var v4735 int32
	_ = v4735
	var v4738 int32
	_ = v4738
	var v4739 int32
	_ = v4739
	var v4744 int32
	_ = v4744
	var v4751 int32
	_ = v4751
	var v4753 int32
	_ = v4753
	var v4755 int32
	_ = v4755
	var v4756 int32
	_ = v4756
	var v4758 int32
	_ = v4758
	var v4760 int32
	_ = v4760
	var v4767 int32
	_ = v4767
	var v4769 int32
	_ = v4769
	var v4772 int32
	_ = v4772
	var v4773 int32
	_ = v4773
	var v4783 int32
	_ = v4783
	var v4810 int32
	_ = v4810
	var v4814 int32
	_ = v4814
	var v4815 int32
	_ = v4815
	var v4818 int32
	_ = v4818
	var v4819 int32
	_ = v4819
	var v4820 int32
	_ = v4820
	var v4830 int32
	_ = v4830
	var v4831 int32
	_ = v4831
	var v4833 int32
	_ = v4833
	var v4836 int32
	_ = v4836
	var v4837 int32
	_ = v4837
	var v4842 int32
	_ = v4842
	var v4849 int32
	_ = v4849
	var v4851 int32
	_ = v4851
	var v4853 int32
	_ = v4853
	var v4854 int32
	_ = v4854
	var v4856 int32
	_ = v4856
	var v4858 int32
	_ = v4858
	var v4865 int32
	_ = v4865
	var v4866 int32
	_ = v4866
	var v4869 int32
	_ = v4869
	var v4870 int32
	_ = v4870
	var v4871 int32
	_ = v4871
	var v4881 int32
	_ = v4881
	var v4882 int32
	_ = v4882
	var v4884 int32
	_ = v4884
	var v4887 int32
	_ = v4887
	var v4888 int32
	_ = v4888
	var v4893 int32
	_ = v4893
	var v4900 int32
	_ = v4900
	var v4902 int32
	_ = v4902
	var v4904 int32
	_ = v4904
	var v4905 int32
	_ = v4905
	var v4907 int32
	_ = v4907
	var v4909 int32
	_ = v4909
	var v4916 int32
	_ = v4916
	var v4929 int32
	_ = v4929
	var v4932 int32
	_ = v4932
	var v4933 int32
	_ = v4933
	var v4970 int32
	_ = v4970
	var v4971 int32
	_ = v4971
	var v4976 int32
	_ = v4976
	var v4980 int32
	_ = v4980
	var v4985 int32
	_ = v4985
	var v5022 int32
	_ = v5022
	var v5027 int32
	_ = v5027
	var v5030 int32
	_ = v5030
	var v5031 int32
	_ = v5031
	var v5032 int32
	_ = v5032
	var v5033 int32
	_ = v5033
	var v5037 int32
	_ = v5037
	var v5042 int32
	_ = v5042
	var v5043 int32
	_ = v5043
	var v5048 int32
	_ = v5048
	var v5056 int32
	_ = v5056
	var v5057 int32
	_ = v5057
	var v5060 int32
	_ = v5060
	var v5065 int32
	_ = v5065
	var v5068 int32
	_ = v5068
	var v5071 int32
	_ = v5071
	var v5072 int32
	_ = v5072
	var v5074 int32
	_ = v5074
	var v5082 int32
	_ = v5082
	var v5083 int32
	_ = v5083
	var v5086 int32
	_ = v5086
	var v5089 int32
	_ = v5089
	var v5094 int32
	_ = v5094
	var v5104 int32
	_ = v5104
	var v5109 int32
	_ = v5109
	var v5114 int32
	_ = v5114
	var v5149 int32
	_ = v5149
	var v5154 int32
	_ = v5154
	var v5157 int32
	_ = v5157
	var v5163 int32
	_ = v5163
	var v5166 int32
	_ = v5166
	var v5170 int32
	_ = v5170
	v8 = int32(0)
	v35 = m.G0
	v37 = v35 + int32(-64)
	m.G0 = v37
	v41 = *(*int32)(unsafe.Add(mBase, uint32(l1)+4))
	if v41 == int32(3) {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	v44 = int32(252)
	goto L3
L2:
	;
	v44 = int32(8)
	goto L3
L3:
	;
	v46 = *(*int32)(unsafe.Add(mBase, uint32(l1+v44)))
	v47 = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v37)+48)) = v47
	*(*int32)(unsafe.Add(mBase, uint32(v37)+28)) = l5
	*(*int32)(unsafe.Add(mBase, uint32(v37)+20)) = v47
	*(*int32)(unsafe.Add(mBase, uint32(v37)+16)) = l6
	v53 = *(*int64)(unsafe.Add(mBase, uint32(l1)+32))
	*(*int64)(unsafe.Add(mBase, uint32(v37)+56)) = v53
	v56 = *(*int32)(unsafe.Add(mBase, _c_F_add_paths_to_joinrel[0]))
	if v56 != 0 {
		goto L4
	} else {
		goto L5
	}
L4:
	;
	m.T0[v56].(func(*base.Module, int32, int32, int32, int32, int32, int32))(m, l0, l1, l2, l3, l4, v35+int32(-48))
	mBase = m.M
	v60 = m.ExcPending
	if v60 != 0 {
		goto L7
	} else {
		goto L8
	}
L5:
	;
	goto L6
L6:
	;
	switch l4 - int32(4) {
	case 0:
		goto L13
	default:
		goto L10
	case 4:
		goto L11
	case 5:
		goto L12
	}
L7:
	;
	return
L8:
	;
	goto L6
L9:
	;
	*(*uint8)(unsafe.Add(mBase, uint32(v37)+24)) = uint8(v129)
	if l4&int32(-2) != int32(8) {
		goto L31
	} else {
		goto L32
	}
L10:
	;
	v125 = *(*int32)(unsafe.Add(mBase, uint32(l1)+8))
	v126 = *(*int32)(unsafe.Add(mBase, uint32(l2)+8))
	v127 = F_innerrel_is_unique(m, l0, v125, v126, l3, l4, l6)
	mBase = m.M
	v128 = m.ExcPending
	if v128 != 0 {
		goto L7
	} else {
		goto L29
	}
L11:
	;
	v120 = *(*int32)(unsafe.Add(mBase, uint32(l1)+8))
	v121 = *(*int32)(unsafe.Add(mBase, uint32(l2)+8))
	v123 = F_innerrel_is_unique(m, l0, v120, v121, l3, int32(0), l6)
	mBase = m.M
	v124 = m.ExcPending
	if v124 != 0 {
		goto L7
	} else {
		goto L28
	}
L12:
	;
	v64 = *(*int32)(unsafe.Add(mBase, uint32(l5)+4))
	v65 = *(*int32)(unsafe.Add(mBase, uint32(l2)+8))
	v66 = int32(0)
	if v64 == v66 {
		goto L15
	} else {
		goto L16
	}
L13:
	;
	v129 = int32(0)
	goto L9
L14:
	;
	v129 = v119
	goto L9
L15:
	;
	v119 = int32(1)
	goto L14
L16:
	;
	goto L17
L17:
	;
	if v65 == int32(0) {
		v112 = v66
		goto L18
	} else {
		goto L19
	}
L18:
	;
	v119 = v112
	goto L14
L19:
	;
	v75 = *(*int32)(unsafe.Add(mBase, uint32(v64)+4))
	v76 = *(*int32)(unsafe.Add(mBase, uint32(v65)+4))
	if v76 < v75 {
		v112 = v66
		goto L18
	} else {
		goto L20
	}
L20:
	;
	v78 = int32(1)
	if v75 <= v78 {
		goto L21
	} else {
		goto L22
	}
L21:
	;
	v81 = v78
	goto L23
L22:
	;
	v81 = v75
	goto L23
L23:
	;
	v82 = int32(8)
	v87 = int32(0)
	goto L24
L24:
	;
	v94 = v87 << (uint(int32(2)) % 32)
	v96 = *(*int32)(unsafe.Add(mBase, uint32(v64+v82+v94)))
	v98 = *(*int32)(unsafe.Add(mBase, uint32(v65+v82+v94)))
	v101 = v96 & (v98 ^ int32(-1))
	v103 = base.B2i32(v101 == int32(0))
	if v101 != 0 {
		v112 = v103
		goto L18
	} else {
		goto L26
	}
L25:
	;
	v112 = v103
	goto L18
L26:
	;
	v105 = v87 + int32(1)
	if v105 != v81 {
		v87 = v105
		goto L24
	} else {
		goto L27
	}
L27:
	;
	goto L25
L28:
	;
	v129 = v123
	goto L9
L29:
	;
	v129 = v127
	goto L9
L30:
	;
	v814 = int32(0)
	if base.B2i32(v785&int32(1) == v814)&base.B2i32(v136&int32(-2) != int32(4)) == v814 {
		goto L166
	} else {
		goto L167
	}
L31:
	;
	v136 = l4
	goto L33
L32:
	;
	v136 = int32(0)
	goto L33
L33:
	;
	if v136 != int32(2) {
		goto L34
	} else {
		goto L35
	}
L34:
	;
	v140 = *(*int64)(unsafe.Add(mBase, uint32(v37)+56))
	if v140&int64(192) == int64(0) {
		v785 = v129
		v797 = int32(1)
		goto L30
	} else {
		goto L37
	}
L35:
	;
	goto L36
L36:
	;
	if v136 == int32(6) {
		v755 = v8
		v761 = int32(0)
		goto L38
	} else {
		goto L39
	}
L37:
	;
	goto L36
L38:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v37)+20)) = v755
	v777 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v37)+24)))
	v785 = v777
	v797 = v761
	goto L30
L39:
	;
	v149 = int32(1)
	if l6 == int32(0) {
		v712 = v8
		v718 = v149
		goto L40
	} else {
		goto L41
	}
L40:
	;
	if int32(1)<<(uint(v136)%32)&int32(140) != 0 {
		goto L162
	} else {
		goto L163
	}
L41:
	;
	v152 = *(*int32)(unsafe.Add(mBase, uint32(l6)+4))
	if v152 <= int32(0) {
		v712 = v8
		v718 = v149
		goto L40
	} else {
		goto L42
	}
L42:
	;
	v171 = v8
	v172 = v8
	v177 = v8
	goto L43
L43:
	;
	v193 = *(*int32)(unsafe.Add(mBase, uint32(l6)+12))
	v197 = *(*int32)(unsafe.Add(mBase, uint32(v193+v177<<(uint(int32(2))%32))))
	if int32(1)<<(uint(v136)%32)&int32(174) != 0 {
		goto L46
	} else {
		goto L47
	}
L44:
	;
	v712 = v672
	v718 = v671 ^ int32(1)
	goto L40
L45:
	;
	v694 = v177 + int32(1)
	v695 = *(*int32)(unsafe.Add(mBase, uint32(l6)+4))
	if v694 < v695 {
		v171 = v671
		v172 = v672
		v177 = v694
		goto L43
	} else {
		goto L161
	}
L46:
	;
	v198 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v197)+8)))
	if v198 != 0 {
		v671 = v171
		v672 = v172
		goto L45
	} else {
		goto L49
	}
L47:
	;
	goto L48
L48:
	;
	v257 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v197)+9)))
	if v257 == int32(1) {
		goto L66
	} else {
		goto L67
	}
L49:
	;
	v199 = *(*int32)(unsafe.Add(mBase, uint32(v197)+32))
	v200 = *(*int32)(unsafe.Add(mBase, uint32(l1)+8))
	v201 = int32(0)
	if v199 == v201 {
		goto L51
	} else {
		goto L52
	}
L50:
	;
	if v254 == int32(0) {
		v671 = v171
		v672 = v172
		goto L45
	} else {
		goto L64
	}
L51:
	;
	v254 = int32(1)
	goto L50
L52:
	;
	goto L53
L53:
	;
	if v200 == int32(0) {
		v247 = v201
		goto L54
	} else {
		goto L55
	}
L54:
	;
	v254 = v247
	goto L50
L55:
	;
	v210 = *(*int32)(unsafe.Add(mBase, uint32(v199)+4))
	v211 = *(*int32)(unsafe.Add(mBase, uint32(v200)+4))
	if v211 < v210 {
		v247 = v201
		goto L54
	} else {
		goto L56
	}
L56:
	;
	v213 = int32(1)
	if v210 <= v213 {
		goto L57
	} else {
		goto L58
	}
L57:
	;
	v216 = v213
	goto L59
L58:
	;
	v216 = v210
	goto L59
L59:
	;
	v217 = int32(8)
	v222 = int32(0)
	goto L60
L60:
	;
	v229 = v222 << (uint(int32(2)) % 32)
	v231 = *(*int32)(unsafe.Add(mBase, uint32(v199+v217+v229)))
	v233 = *(*int32)(unsafe.Add(mBase, uint32(v200+v217+v229)))
	v236 = v231 & (v233 ^ int32(-1))
	v238 = base.B2i32(v236 == int32(0))
	if v236 != 0 {
		v247 = v238
		goto L54
	} else {
		goto L62
	}
L61:
	;
	v247 = v238
	goto L54
L62:
	;
	v240 = v222 + int32(1)
	if v240 != v216 {
		v222 = v240
		goto L60
	} else {
		goto L63
	}
L63:
	;
	goto L61
L64:
	;
	goto L48
L65:
	;
	v266 = *(*int32)(unsafe.Add(mBase, uint32(l3)+8))
	v267 = *(*int32)(unsafe.Add(mBase, uint32(v197)+44))
	v268 = *(*int32)(unsafe.Add(mBase, uint32(l2)+8))
	v269 = int32(0)
	if v267 == v269 {
		goto L77
	} else {
		goto L78
	}
L66:
	;
	v260 = *(*int32)(unsafe.Add(mBase, uint32(v197)+96))
	if v260 != 0 {
		goto L65
	} else {
		goto L69
	}
L67:
	;
	goto L68
L68:
	;
	v261 = *(*int32)(unsafe.Add(mBase, uint32(v197)+4))
	if v261 != 0 {
		goto L70
	} else {
		goto L71
	}
L69:
	;
	goto L68
L70:
	;
	v262 = *(*int32)(unsafe.Add(mBase, uint32(v261)))
	if v262 == int32(7) {
		v671 = v171
		v672 = v172
		goto L45
	} else {
		goto L73
	}
L71:
	;
	goto L72
L72:
	;
	v671 = int32(1)
	v672 = v172
	goto L45
L73:
	;
	goto L72
L74:
	;
	v507 = *(*int32)(unsafe.Add(mBase, uint32(v197)+100))
	v508 = *(*int32)(unsafe.Add(mBase, uint32(v507)+56))
	if v508 != 0 {
		goto L142
	} else {
		goto L143
	}
L75:
	;
	v384 = *(*int32)(unsafe.Add(mBase, uint32(v197)+44))
	v385 = int32(0)
	if v384 == v385 {
		goto L107
	} else {
		goto L108
	}
L76:
	;
	if v322 == int32(0) {
		goto L75
	} else {
		goto L90
	}
L77:
	;
	v322 = int32(1)
	goto L76
L78:
	;
	goto L79
L79:
	;
	if v268 == int32(0) {
		v315 = v269
		goto L80
	} else {
		goto L81
	}
L80:
	;
	v322 = v315
	goto L76
L81:
	;
	v278 = *(*int32)(unsafe.Add(mBase, uint32(v267)+4))
	v279 = *(*int32)(unsafe.Add(mBase, uint32(v268)+4))
	if v279 < v278 {
		v315 = v269
		goto L80
	} else {
		goto L82
	}
L82:
	;
	v281 = int32(1)
	if v278 <= v281 {
		goto L83
	} else {
		goto L84
	}
L83:
	;
	v284 = v281
	goto L85
L84:
	;
	v284 = v278
	goto L85
L85:
	;
	v285 = int32(8)
	v290 = int32(0)
	goto L86
L86:
	;
	v297 = v290 << (uint(int32(2)) % 32)
	v299 = *(*int32)(unsafe.Add(mBase, uint32(v267+v285+v297)))
	v301 = *(*int32)(unsafe.Add(mBase, uint32(v268+v285+v297)))
	v304 = v299 & (v301 ^ int32(-1))
	v306 = base.B2i32(v304 == int32(0))
	if v304 != 0 {
		v315 = v306
		goto L80
	} else {
		goto L88
	}
L87:
	;
	v315 = v306
	goto L80
L88:
	;
	v308 = v290 + int32(1)
	if v308 != v284 {
		v290 = v308
		goto L86
	} else {
		goto L89
	}
L89:
	;
	goto L87
L90:
	;
	v325 = *(*int32)(unsafe.Add(mBase, uint32(v197)+48))
	v326 = int32(0)
	if v325 == v326 {
		goto L92
	} else {
		goto L93
	}
L91:
	;
	if v379 == int32(0) {
		goto L75
	} else {
		goto L105
	}
L92:
	;
	v379 = int32(1)
	goto L91
L93:
	;
	goto L94
L94:
	;
	if v266 == int32(0) {
		v372 = v326
		goto L95
	} else {
		goto L96
	}
L95:
	;
	v379 = v372
	goto L91
L96:
	;
	v335 = *(*int32)(unsafe.Add(mBase, uint32(v325)+4))
	v336 = *(*int32)(unsafe.Add(mBase, uint32(v266)+4))
	if v336 < v335 {
		v372 = v326
		goto L95
	} else {
		goto L97
	}
L97:
	;
	v338 = int32(1)
	if v335 <= v338 {
		goto L98
	} else {
		goto L99
	}
L98:
	;
	v341 = v338
	goto L100
L99:
	;
	v341 = v335
	goto L100
L100:
	;
	v342 = int32(8)
	v347 = int32(0)
	goto L101
L101:
	;
	v354 = v347 << (uint(int32(2)) % 32)
	v356 = *(*int32)(unsafe.Add(mBase, uint32(v325+v342+v354)))
	v358 = *(*int32)(unsafe.Add(mBase, uint32(v266+v342+v354)))
	v361 = v356 & (v358 ^ int32(-1))
	v363 = base.B2i32(v361 == int32(0))
	if v361 != 0 {
		v372 = v363
		goto L95
	} else {
		goto L103
	}
L102:
	;
	v372 = v363
	goto L95
L103:
	;
	v365 = v347 + int32(1)
	if v365 != v341 {
		v347 = v365
		goto L101
	} else {
		goto L104
	}
L104:
	;
	goto L102
L105:
	;
	v382 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(v197)+120)) = uint8(v382)
	goto L74
L106:
	;
	if v438 == int32(0) {
		goto L120
	} else {
		goto L121
	}
L107:
	;
	v438 = int32(1)
	goto L106
L108:
	;
	goto L109
L109:
	;
	if v266 == int32(0) {
		v431 = v385
		goto L110
	} else {
		goto L111
	}
L110:
	;
	v438 = v431
	goto L106
L111:
	;
	v394 = *(*int32)(unsafe.Add(mBase, uint32(v384)+4))
	v395 = *(*int32)(unsafe.Add(mBase, uint32(v266)+4))
	if v395 < v394 {
		v431 = v385
		goto L110
	} else {
		goto L112
	}
L112:
	;
	v397 = int32(1)
	if v394 <= v397 {
		goto L113
	} else {
		goto L114
	}
L113:
	;
	v400 = v397
	goto L115
L114:
	;
	v400 = v394
	goto L115
L115:
	;
	v401 = int32(8)
	v406 = int32(0)
	goto L116
L116:
	;
	v413 = v406 << (uint(int32(2)) % 32)
	v415 = *(*int32)(unsafe.Add(mBase, uint32(v384+v401+v413)))
	v417 = *(*int32)(unsafe.Add(mBase, uint32(v266+v401+v413)))
	v420 = v415 & (v417 ^ int32(-1))
	v422 = base.B2i32(v420 == int32(0))
	if v420 != 0 {
		v431 = v422
		goto L110
	} else {
		goto L118
	}
L117:
	;
	v431 = v422
	goto L110
L118:
	;
	v424 = v406 + int32(1)
	if v424 != v400 {
		v406 = v424
		goto L116
	} else {
		goto L119
	}
L119:
	;
	goto L117
L120:
	;
	v671 = int32(1)
	v672 = v172
	goto L45
L121:
	;
	goto L122
L122:
	;
	v442 = *(*int32)(unsafe.Add(mBase, uint32(v197)+48))
	v443 = int32(0)
	if v442 == v443 {
		goto L124
	} else {
		goto L125
	}
L123:
	;
	if v496 == int32(0) {
		goto L137
	} else {
		goto L138
	}
L124:
	;
	v496 = int32(1)
	goto L123
L125:
	;
	goto L126
L126:
	;
	if v268 == int32(0) {
		v489 = v443
		goto L127
	} else {
		goto L128
	}
L127:
	;
	v496 = v489
	goto L123
L128:
	;
	v452 = *(*int32)(unsafe.Add(mBase, uint32(v442)+4))
	v453 = *(*int32)(unsafe.Add(mBase, uint32(v268)+4))
	if v453 < v452 {
		v489 = v443
		goto L127
	} else {
		goto L129
	}
L129:
	;
	v455 = int32(1)
	if v452 <= v455 {
		goto L130
	} else {
		goto L131
	}
L130:
	;
	v458 = v455
	goto L132
L131:
	;
	v458 = v452
	goto L132
L132:
	;
	v459 = int32(8)
	v464 = int32(0)
	goto L133
L133:
	;
	v471 = v464 << (uint(int32(2)) % 32)
	v473 = *(*int32)(unsafe.Add(mBase, uint32(v442+v459+v471)))
	v475 = *(*int32)(unsafe.Add(mBase, uint32(v268+v459+v471)))
	v478 = v473 & (v475 ^ int32(-1))
	v480 = base.B2i32(v478 == int32(0))
	if v478 != 0 {
		v489 = v480
		goto L127
	} else {
		goto L135
	}
L134:
	;
	v489 = v480
	goto L127
L135:
	;
	v482 = v464 + int32(1)
	if v482 != v458 {
		v464 = v482
		goto L133
	} else {
		goto L136
	}
L136:
	;
	goto L134
L137:
	;
	v671 = int32(1)
	v672 = v172
	goto L45
L138:
	;
	goto L139
L139:
	;
	v500 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v197)+120)) = uint8(v500)
	v502 = *(*int32)(unsafe.Add(mBase, uint32(v197)+4))
	v503 = *(*int32)(unsafe.Add(mBase, uint32(v502)+4))
	v504 = F_get_commutator(m, v503)
	mBase = m.M
	v505 = m.ExcPending
	if v505 != 0 {
		goto L7
	} else {
		goto L140
	}
L140:
	;
	if v504 != 0 {
		goto L74
	} else {
		goto L141
	}
L141:
	;
	v671 = int32(1)
	v672 = v172
	goto L45
L142:
	;
	v524 = v508
	goto L145
L143:
	;
	goto L144
L144:
	;
	v579 = *(*int32)(unsafe.Add(mBase, uint32(v197)+104))
	v580 = *(*int32)(unsafe.Add(mBase, uint32(v579)+56))
	if v580 != 0 {
		goto L148
	} else {
		goto L149
	}
L145:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v197)+100)) = v524
	v544 = *(*int32)(unsafe.Add(mBase, uint32(v524)+56))
	if v544 != 0 {
		v524 = v544
		goto L145
	} else {
		goto L147
	}
L146:
	;
	goto L144
L147:
	;
	goto L146
L148:
	;
	v596 = v580
	goto L151
L149:
	;
	goto L150
L150:
	;
	v651 = *(*int32)(unsafe.Add(mBase, uint32(v197)+100))
	v652 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v651)+40)))
	if v652 != 0 {
		goto L154
	} else {
		goto L155
	}
L151:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v197)+104)) = v596
	v616 = *(*int32)(unsafe.Add(mBase, uint32(v596)+56))
	if v616 != 0 {
		v596 = v616
		goto L151
	} else {
		goto L153
	}
L152:
	;
	goto L150
L153:
	;
	goto L152
L154:
	;
	v671 = int32(1)
	v672 = v172
	goto L45
L155:
	;
	goto L156
L156:
	;
	v654 = *(*int32)(unsafe.Add(mBase, uint32(v197)+104))
	v655 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v654)+40)))
	if v655 != 0 {
		goto L157
	} else {
		goto L158
	}
L157:
	;
	v671 = int32(1)
	v672 = v172
	goto L45
L158:
	;
	goto L159
L159:
	;
	v657 = F_lappend(m, v172, v197)
	mBase = m.M
	v658 = m.ExcPending
	if v658 != 0 {
		goto L7
	} else {
		goto L160
	}
L160:
	;
	v671 = v171
	v672 = v657
	goto L45
L161:
	;
	goto L44
L162:
	;
	v740 = base.B2i32(base.Ui32(v136) <= base.Ui32(int32(7)))
	goto L164
L163:
	;
	v740 = int32(0)
	goto L164
L164:
	;
	if v740 != 0 {
		v755 = v712
		v761 = v718
		goto L38
	} else {
		goto L165
	}
L165:
	;
	v755 = v712
	v761 = int32(1)
	goto L38
L166:
	;
	v824 = v35 + int32(-32)
	v825 = int32(0)
	v827 = m.G0
	v829 = v827 + int32(-64)
	m.G0 = v829
	v834 = int32(1) << (uint(v136) % 32) & int32(174)
	if v834 == v825 {
		goto L170
	} else {
		goto L171
	}
L167:
	;
	goto L168
L168:
	;
	v1066 = *(*int32)(unsafe.Add(mBase, uint32(l0)+120))
	if v1066 == int32(0) {
		goto L212
	} else {
		goto L213
	}
L169:
	;
	v982 = int32(5)
	if v136 == v982 {
		goto L196
	} else {
		goto L197
	}
L170:
	;
	v952 = l6
	goto L169
L171:
	;
	goto L172
L172:
	;
	if l6 == int32(0) {
		v952 = v825
		goto L169
	} else {
		goto L173
	}
L173:
	;
	v839 = *(*int32)(unsafe.Add(mBase, uint32(l6)+4))
	if v839 <= int32(0) {
		v952 = v825
		goto L169
	} else {
		goto L174
	}
L174:
	;
	v847 = v825
	v854 = v825
	goto L175
L175:
	;
	v876 = *(*int32)(unsafe.Add(mBase, uint32(l6)+12))
	v880 = *(*int32)(unsafe.Add(mBase, uint32(v876+v854<<(uint(int32(2))%32))))
	v881 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v880)+8)))
	if v881 != 0 {
		v942 = v847
		goto L177
	} else {
		goto L178
	}
L176:
	;
	v952 = v942
	goto L169
L177:
	;
	v944 = v854 + int32(1)
	v945 = *(*int32)(unsafe.Add(mBase, uint32(l6)+4))
	if v944 < v945 {
		v847 = v942
		v854 = v944
		goto L175
	} else {
		goto L195
	}
L178:
	;
	v882 = *(*int32)(unsafe.Add(mBase, uint32(v880)+32))
	v883 = *(*int32)(unsafe.Add(mBase, uint32(l1)+8))
	v884 = int32(0)
	if v882 == v884 {
		goto L180
	} else {
		goto L181
	}
L179:
	;
	if v937 == int32(0) {
		v942 = v847
		goto L177
	} else {
		goto L193
	}
L180:
	;
	v937 = int32(1)
	goto L179
L181:
	;
	goto L182
L182:
	;
	if v883 == int32(0) {
		v930 = v884
		goto L183
	} else {
		goto L184
	}
L183:
	;
	v937 = v930
	goto L179
L184:
	;
	v893 = *(*int32)(unsafe.Add(mBase, uint32(v882)+4))
	v894 = *(*int32)(unsafe.Add(mBase, uint32(v883)+4))
	if v894 < v893 {
		v930 = v884
		goto L183
	} else {
		goto L185
	}
L185:
	;
	v896 = int32(1)
	if v893 <= v896 {
		goto L186
	} else {
		goto L187
	}
L186:
	;
	v899 = v896
	goto L188
L187:
	;
	v899 = v893
	goto L188
L188:
	;
	v900 = int32(8)
	v905 = int32(0)
	goto L189
L189:
	;
	v912 = v905 << (uint(int32(2)) % 32)
	v914 = *(*int32)(unsafe.Add(mBase, uint32(v882+v900+v912)))
	v916 = *(*int32)(unsafe.Add(mBase, uint32(v883+v900+v912)))
	v919 = v914 & (v916 ^ int32(-1))
	v921 = base.B2i32(v919 == int32(0))
	if v919 != 0 {
		v930 = v921
		goto L183
	} else {
		goto L191
	}
L190:
	;
	v930 = v921
	goto L183
L191:
	;
	v923 = v905 + int32(1)
	if v923 != v899 {
		v905 = v923
		goto L189
	} else {
		goto L192
	}
L192:
	;
	goto L190
L193:
	;
	v940 = F_lappend(m, v847, v880)
	mBase = m.M
	v941 = m.ExcPending
	if v941 != 0 {
		goto L7
	} else {
		goto L194
	}
L194:
	;
	v942 = v940
	goto L177
L195:
	;
	goto L176
L196:
	;
	v986 = v982
	goto L198
L197:
	;
	v986 = int32(4)
	goto L198
L198:
	;
	v987 = F_clauselist_selectivity(m, l0, v952, int32(0), v986, l5)
	mBase = m.M
	v988 = m.ExcPending
	if v988 != 0 {
		goto L7
	} else {
		goto L199
	}
L199:
	;
	v990 = v827 + int32(-56)
	v991 = *(*int32)(unsafe.Add(mBase, uint32(l2)+8))
	v992 = *(*int32)(unsafe.Add(mBase, uint32(l3)+8))
	v993 = int64(0)
	*(*int64)(unsafe.Add(mBase, uint32(v990)+48)) = v993
	*(*int32)(unsafe.Add(mBase, uint32(v990)+16)) = v992
	*(*int32)(unsafe.Add(mBase, uint32(v990)+12)) = v991
	*(*int32)(unsafe.Add(mBase, uint32(v990)+8)) = v992
	*(*int32)(unsafe.Add(mBase, uint32(v990)+4)) = v991
	*(*int32)(unsafe.Add(mBase, uint32(v990))) = int32(322)
	*(*int64)(unsafe.Add(mBase, uint32(v990)+20)) = v993
	*(*int64)(unsafe.Add(mBase, uint32(v990)+28)) = v993
	*(*int64)(unsafe.Add(mBase, uint32(v990)+36)) = v993
	*(*int32)(unsafe.Add(mBase, uint32(v990)+43)) = int32(0)
	goto L200
L200:
	;
	v1009 = int32(0)
	v1011 = F_clauselist_selectivity(m, l0, v952, v1009, v1009, v990)
	mBase = m.M
	v1012 = m.ExcPending
	if v1012 != 0 {
		goto L7
	} else {
		goto L201
	}
L201:
	;
	if v834 != 0 {
		goto L202
	} else {
		goto L203
	}
L202:
	;
	F_list_free(m, v952)
	mBase = m.M
	v1014 = m.ExcPending
	if v1014 != 0 {
		goto L7
	} else {
		goto L205
	}
L203:
	;
	goto L204
L204:
	;
	if base.F64_gt(v987, float64(0)) != 0 {
		goto L206
	} else {
		goto L207
	}
L205:
	;
	goto L204
L206:
	;
	v1017 = float64(1)
	v1018 = *(*float64)(unsafe.Add(mBase, uint32(l3)+16))
	v1020 = base.F64_div(base.F64_mul(v1011, v1018), v987)
	if base.F64_lt(v1020, v1017) != 0 {
		goto L209
	} else {
		goto L210
	}
L207:
	;
	v1026 = float64(1)
	goto L208
L208:
	;
	*(*float64)(unsafe.Add(mBase, uint32(v824)+8)) = v1026
	*(*float64)(unsafe.Add(mBase, uint32(v824))) = v987
	m.G0 = v829 - int32(-64)
	goto L168
L209:
	;
	v1023 = v1017
	goto L211
L210:
	;
	v1023 = v1020
	goto L211
L211:
	;
	v1026 = v1023
	goto L208
L212:
	;
	v1361 = *(*int32)(unsafe.Add(mBase, uint32(v37)+48))
	v1362 = *(*int32)(unsafe.Add(mBase, uint32(l1)+72))
	v1363 = F_bms_add_members(m, v1361, v1362)
	mBase = m.M
	v1364 = m.ExcPending
	if v1364 != 0 {
		goto L7
	} else {
		goto L281
	}
L213:
	;
	v1069 = *(*int32)(unsafe.Add(mBase, uint32(v1066)+4))
	if v1069 <= int32(0) {
		goto L212
	} else {
		goto L214
	}
L214:
	;
	v1079 = int32(0)
	goto L215
L215:
	;
	v1107 = *(*int32)(unsafe.Add(mBase, uint32(v1066)+12))
	v1111 = *(*int32)(unsafe.Add(mBase, uint32(v1107+v1079<<(uint(int32(2))%32))))
	v1112 = *(*int32)(unsafe.Add(mBase, uint32(v1111)+8))
	v1113 = int32(0)
	if base.B2i32(v46 == v1113)|base.B2i32(v1112 == v1113) != 0 {
		v1158 = v1113
		goto L219
	} else {
		goto L220
	}
L216:
	;
	goto L212
L217:
	;
	v1216 = *(*int32)(unsafe.Add(mBase, uint32(v1111)+20))
	if v1216 != int32(2) {
		goto L248
	} else {
		goto L249
	}
L218:
	;
	if v1158 == int32(0) {
		goto L217
	} else {
		goto L231
	}
L219:
	;
	goto L218
L220:
	;
	v1123 = *(*int32)(unsafe.Add(mBase, uint32(v46)+4))
	v1124 = *(*int32)(unsafe.Add(mBase, uint32(v1112)+4))
	if v1123 < v1124 {
		goto L221
	} else {
		goto L222
	}
L221:
	;
	v1126 = v1123
	goto L223
L222:
	;
	v1126 = v1124
	goto L223
L223:
	;
	if v1126 <= int32(1) {
		goto L224
	} else {
		goto L225
	}
L224:
	;
	v1129 = int32(1)
	goto L226
L225:
	;
	v1129 = v1126
	goto L226
L226:
	;
	v1130 = int32(8)
	v1135 = int32(0)
	goto L227
L227:
	;
	v1142 = v1135 << (uint(int32(2)) % 32)
	v1144 = *(*int32)(unsafe.Add(mBase, uint32(v1112+v1130+v1142)))
	v1146 = *(*int32)(unsafe.Add(mBase, uint32(v46+v1130+v1142)))
	v1147 = v1144 & v1146
	v1149 = base.B2i32(v1147 != int32(0))
	if v1147 != 0 {
		v1158 = v1149
		goto L219
	} else {
		goto L229
	}
L228:
	;
	v1158 = v1149
	goto L219
L229:
	;
	v1151 = v1135 + int32(1)
	if v1151 != v1129 {
		v1135 = v1151
		goto L227
	} else {
		goto L230
	}
L230:
	;
	goto L228
L231:
	;
	v1161 = *(*int32)(unsafe.Add(mBase, uint32(v1111)+4))
	v1162 = int32(0)
	if base.B2i32(v46 == v1162)|base.B2i32(v1161 == v1162) != 0 {
		v1207 = v1162
		goto L233
	} else {
		goto L234
	}
L232:
	;
	if v1207 != 0 {
		goto L217
	} else {
		goto L245
	}
L233:
	;
	goto L232
L234:
	;
	v1172 = *(*int32)(unsafe.Add(mBase, uint32(v46)+4))
	v1173 = *(*int32)(unsafe.Add(mBase, uint32(v1161)+4))
	if v1172 < v1173 {
		goto L235
	} else {
		goto L236
	}
L235:
	;
	v1175 = v1172
	goto L237
L236:
	;
	v1175 = v1173
	goto L237
L237:
	;
	if v1175 <= int32(1) {
		goto L238
	} else {
		goto L239
	}
L238:
	;
	v1178 = int32(1)
	goto L240
L239:
	;
	v1178 = v1175
	goto L240
L240:
	;
	v1179 = int32(8)
	v1184 = int32(0)
	goto L241
L241:
	;
	v1191 = v1184 << (uint(int32(2)) % 32)
	v1193 = *(*int32)(unsafe.Add(mBase, uint32(v1161+v1179+v1191)))
	v1195 = *(*int32)(unsafe.Add(mBase, uint32(v46+v1179+v1191)))
	v1196 = v1193 & v1195
	v1198 = base.B2i32(v1196 != int32(0))
	if v1196 != 0 {
		v1207 = v1198
		goto L233
	} else {
		goto L243
	}
L242:
	;
	v1207 = v1198
	goto L233
L243:
	;
	v1200 = v1184 + int32(1)
	if v1200 != v1178 {
		v1184 = v1200
		goto L241
	} else {
		goto L244
	}
L244:
	;
	goto L242
L245:
	;
	v1208 = *(*int32)(unsafe.Add(mBase, uint32(v37)+48))
	v1209 = *(*int32)(unsafe.Add(mBase, uint32(l0)+52))
	v1210 = *(*int32)(unsafe.Add(mBase, uint32(v1111)+8))
	v1211 = F_bms_difference(m, v1209, v1210)
	mBase = m.M
	v1212 = m.ExcPending
	if v1212 != 0 {
		goto L7
	} else {
		goto L246
	}
L246:
	;
	v1213 = F_bms_join(m, v1208, v1211)
	mBase = m.M
	v1214 = m.ExcPending
	if v1214 != 0 {
		goto L7
	} else {
		goto L247
	}
L247:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v37)+48)) = v1213
	goto L217
L248:
	;
	v1324 = v1079 + int32(1)
	v1325 = *(*int32)(unsafe.Add(mBase, uint32(v1066)+4))
	if v1324 < v1325 {
		v1079 = v1324
		goto L215
	} else {
		goto L280
	}
L249:
	;
	v1219 = *(*int32)(unsafe.Add(mBase, uint32(v1111)+4))
	v1220 = int32(0)
	if base.B2i32(v46 == v1220)|base.B2i32(v1219 == v1220) != 0 {
		v1265 = v1220
		goto L251
	} else {
		goto L252
	}
L250:
	;
	if v1265 == int32(0) {
		goto L248
	} else {
		goto L263
	}
L251:
	;
	goto L250
L252:
	;
	v1230 = *(*int32)(unsafe.Add(mBase, uint32(v46)+4))
	v1231 = *(*int32)(unsafe.Add(mBase, uint32(v1219)+4))
	if v1230 < v1231 {
		goto L253
	} else {
		goto L254
	}
L253:
	;
	v1233 = v1230
	goto L255
L254:
	;
	v1233 = v1231
	goto L255
L255:
	;
	if v1233 <= int32(1) {
		goto L256
	} else {
		goto L257
	}
L256:
	;
	v1236 = int32(1)
	goto L258
L257:
	;
	v1236 = v1233
	goto L258
L258:
	;
	v1237 = int32(8)
	v1242 = int32(0)
	goto L259
L259:
	;
	v1249 = v1242 << (uint(int32(2)) % 32)
	v1251 = *(*int32)(unsafe.Add(mBase, uint32(v1219+v1237+v1249)))
	v1253 = *(*int32)(unsafe.Add(mBase, uint32(v46+v1237+v1249)))
	v1254 = v1251 & v1253
	v1256 = base.B2i32(v1254 != int32(0))
	if v1254 != 0 {
		v1265 = v1256
		goto L251
	} else {
		goto L261
	}
L260:
	;
	v1265 = v1256
	goto L251
L261:
	;
	v1258 = v1242 + int32(1)
	if v1258 != v1236 {
		v1242 = v1258
		goto L259
	} else {
		goto L262
	}
L262:
	;
	goto L260
L263:
	;
	v1268 = *(*int32)(unsafe.Add(mBase, uint32(v1111)+8))
	v1269 = int32(0)
	if base.B2i32(v46 == v1269)|base.B2i32(v1268 == v1269) != 0 {
		v1314 = v1269
		goto L265
	} else {
		goto L266
	}
L264:
	;
	if v1314 != 0 {
		goto L248
	} else {
		goto L277
	}
L265:
	;
	goto L264
L266:
	;
	v1279 = *(*int32)(unsafe.Add(mBase, uint32(v46)+4))
	v1280 = *(*int32)(unsafe.Add(mBase, uint32(v1268)+4))
	if v1279 < v1280 {
		goto L267
	} else {
		goto L268
	}
L267:
	;
	v1282 = v1279
	goto L269
L268:
	;
	v1282 = v1280
	goto L269
L269:
	;
	if v1282 <= int32(1) {
		goto L270
	} else {
		goto L271
	}
L270:
	;
	v1285 = int32(1)
	goto L272
L271:
	;
	v1285 = v1282
	goto L272
L272:
	;
	v1286 = int32(8)
	v1291 = int32(0)
	goto L273
L273:
	;
	v1298 = v1291 << (uint(int32(2)) % 32)
	v1300 = *(*int32)(unsafe.Add(mBase, uint32(v1268+v1286+v1298)))
	v1302 = *(*int32)(unsafe.Add(mBase, uint32(v46+v1286+v1298)))
	v1303 = v1300 & v1302
	v1305 = base.B2i32(v1303 != int32(0))
	if v1303 != 0 {
		v1314 = v1305
		goto L265
	} else {
		goto L275
	}
L274:
	;
	v1314 = v1305
	goto L265
L275:
	;
	v1307 = v1291 + int32(1)
	if v1307 != v1285 {
		v1291 = v1307
		goto L273
	} else {
		goto L276
	}
L276:
	;
	goto L274
L277:
	;
	v1315 = *(*int32)(unsafe.Add(mBase, uint32(v37)+48))
	v1316 = *(*int32)(unsafe.Add(mBase, uint32(l0)+52))
	v1317 = *(*int32)(unsafe.Add(mBase, uint32(v1111)+4))
	v1318 = F_bms_difference(m, v1316, v1317)
	mBase = m.M
	v1319 = m.ExcPending
	if v1319 != 0 {
		goto L7
	} else {
		goto L278
	}
L278:
	;
	v1320 = F_bms_join(m, v1315, v1318)
	mBase = m.M
	v1321 = m.ExcPending
	if v1321 != 0 {
		goto L7
	} else {
		goto L279
	}
L279:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v37)+48)) = v1320
	goto L248
L280:
	;
	goto L216
L281:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v37)+48)) = v1363
	if v797&int32(1) == int32(0) {
		goto L285
	} else {
		goto L286
	}
L282:
	;
	v5149 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v37)+56)))
	if v5149&int32(32) == int32(0) {
		goto L1039
	} else {
		goto L1040
	}
L283:
	;
	if v136 == int32(6) {
		goto L282
	} else {
		goto L1005
	}
L284:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v4976 = m.ExcPending
	if v4976 != 0 {
		goto L7
	} else {
		goto L1002
	}
L285:
	;
	if v136 != int32(2) {
		goto L758
	} else {
		goto L759
	}
L286:
	;
	v1370 = *(*int32)(unsafe.Add(mBase, uint32(v37)+20))
	if v1370 == int32(0) {
		goto L287
	} else {
		goto L288
	}
L287:
	;
	if base.Ui32(int32(7)) < base.Ui32(v136) {
		goto L284
	} else {
		goto L542
	}
L288:
	;
	v1373 = *(*int32)(unsafe.Add(mBase, uint32(l3)+60))
	v1374 = *(*int32)(unsafe.Add(mBase, uint32(l2)+60))
	v1375 = *(*int32)(unsafe.Add(mBase, uint32(v1374)+16))
	if v1375 == int32(0) {
		goto L289
	} else {
		goto L290
	}
L289:
	;
	v1478 = *(*int32)(unsafe.Add(mBase, uint32(v1373)+16))
	if v1478 == int32(0) {
		goto L320
	} else {
		goto L321
	}
L290:
	;
	v1378 = *(*int32)(unsafe.Add(mBase, uint32(v1375)+4))
	v1379 = *(*int32)(unsafe.Add(mBase, uint32(l3)+8))
	v1380 = int32(0)
	if base.B2i32(v1378 == v1380)|base.B2i32(v1379 == v1380) != 0 {
		v1425 = v1380
		goto L292
	} else {
		goto L293
	}
L291:
	;
	if v1425 != 0 {
		goto L287
	} else {
		goto L304
	}
L292:
	;
	goto L291
L293:
	;
	v1390 = *(*int32)(unsafe.Add(mBase, uint32(v1378)+4))
	v1391 = *(*int32)(unsafe.Add(mBase, uint32(v1379)+4))
	if v1390 < v1391 {
		goto L294
	} else {
		goto L295
	}
L294:
	;
	v1393 = v1390
	goto L296
L295:
	;
	v1393 = v1391
	goto L296
L296:
	;
	if v1393 <= int32(1) {
		goto L297
	} else {
		goto L298
	}
L297:
	;
	v1396 = int32(1)
	goto L299
L298:
	;
	v1396 = v1393
	goto L299
L299:
	;
	v1397 = int32(8)
	v1402 = int32(0)
	goto L300
L300:
	;
	v1409 = v1402 << (uint(int32(2)) % 32)
	v1411 = *(*int32)(unsafe.Add(mBase, uint32(v1379+v1397+v1409)))
	v1413 = *(*int32)(unsafe.Add(mBase, uint32(v1378+v1397+v1409)))
	v1414 = v1411 & v1413
	v1416 = base.B2i32(v1414 != int32(0))
	if v1414 != 0 {
		v1425 = v1416
		goto L292
	} else {
		goto L302
	}
L301:
	;
	v1425 = v1416
	goto L292
L302:
	;
	v1418 = v1402 + int32(1)
	if v1418 != v1396 {
		v1402 = v1418
		goto L300
	} else {
		goto L303
	}
L303:
	;
	goto L301
L304:
	;
	v1426 = *(*int32)(unsafe.Add(mBase, uint32(v1374)+16))
	if v1426 == int32(0) {
		goto L289
	} else {
		goto L305
	}
L305:
	;
	v1429 = *(*int32)(unsafe.Add(mBase, uint32(v1426)+4))
	v1430 = *(*int32)(unsafe.Add(mBase, uint32(l3)+252))
	v1431 = int32(0)
	if base.B2i32(v1429 == v1431)|base.B2i32(v1430 == v1431) != 0 {
		v1476 = v1431
		goto L307
	} else {
		goto L308
	}
L306:
	;
	if v1476 != 0 {
		goto L287
	} else {
		goto L319
	}
L307:
	;
	goto L306
L308:
	;
	v1441 = *(*int32)(unsafe.Add(mBase, uint32(v1429)+4))
	v1442 = *(*int32)(unsafe.Add(mBase, uint32(v1430)+4))
	if v1441 < v1442 {
		goto L309
	} else {
		goto L310
	}
L309:
	;
	v1444 = v1441
	goto L311
L310:
	;
	v1444 = v1442
	goto L311
L311:
	;
	if v1444 <= int32(1) {
		goto L312
	} else {
		goto L313
	}
L312:
	;
	v1447 = int32(1)
	goto L314
L313:
	;
	v1447 = v1444
	goto L314
L314:
	;
	v1448 = int32(8)
	v1453 = int32(0)
	goto L315
L315:
	;
	v1460 = v1453 << (uint(int32(2)) % 32)
	v1462 = *(*int32)(unsafe.Add(mBase, uint32(v1430+v1448+v1460)))
	v1464 = *(*int32)(unsafe.Add(mBase, uint32(v1429+v1448+v1460)))
	v1465 = v1462 & v1464
	v1467 = base.B2i32(v1465 != int32(0))
	if v1465 != 0 {
		v1476 = v1467
		goto L307
	} else {
		goto L317
	}
L316:
	;
	v1476 = v1467
	goto L307
L317:
	;
	v1469 = v1453 + int32(1)
	if v1469 != v1447 {
		v1453 = v1469
		goto L315
	} else {
		goto L318
	}
L318:
	;
	goto L316
L319:
	;
	goto L289
L320:
	;
	v1581 = int32(0)
	v1582 = int32(1)
	v1593 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l1)+26)))
	if base.B2i32(base.B2i32(v1582<<(uint(v136)%32)&int32(140) == v1581)|base.B2i32(base.Ui32(int32(7)) < base.Ui32(v136)) == v1581)|base.B2i32(v1593 != v1582) != 0 {
		v1652 = v1581
		v1653 = v8
		goto L351
	} else {
		goto L352
	}
L321:
	;
	v1481 = *(*int32)(unsafe.Add(mBase, uint32(v1478)+4))
	v1482 = *(*int32)(unsafe.Add(mBase, uint32(l2)+8))
	v1483 = int32(0)
	if base.B2i32(v1481 == v1483)|base.B2i32(v1482 == v1483) != 0 {
		v1528 = v1483
		goto L323
	} else {
		goto L324
	}
L322:
	;
	if v1528 != 0 {
		goto L287
	} else {
		goto L335
	}
L323:
	;
	goto L322
L324:
	;
	v1493 = *(*int32)(unsafe.Add(mBase, uint32(v1481)+4))
	v1494 = *(*int32)(unsafe.Add(mBase, uint32(v1482)+4))
	if v1493 < v1494 {
		goto L325
	} else {
		goto L326
	}
L325:
	;
	v1496 = v1493
	goto L327
L326:
	;
	v1496 = v1494
	goto L327
L327:
	;
	if v1496 <= int32(1) {
		goto L328
	} else {
		goto L329
	}
L328:
	;
	v1499 = int32(1)
	goto L330
L329:
	;
	v1499 = v1496
	goto L330
L330:
	;
	v1500 = int32(8)
	v1505 = int32(0)
	goto L331
L331:
	;
	v1512 = v1505 << (uint(int32(2)) % 32)
	v1514 = *(*int32)(unsafe.Add(mBase, uint32(v1482+v1500+v1512)))
	v1516 = *(*int32)(unsafe.Add(mBase, uint32(v1481+v1500+v1512)))
	v1517 = v1514 & v1516
	v1519 = base.B2i32(v1517 != int32(0))
	if v1517 != 0 {
		v1528 = v1519
		goto L323
	} else {
		goto L333
	}
L332:
	;
	v1528 = v1519
	goto L323
L333:
	;
	v1521 = v1505 + int32(1)
	if v1521 != v1499 {
		v1505 = v1521
		goto L331
	} else {
		goto L334
	}
L334:
	;
	goto L332
L335:
	;
	v1529 = *(*int32)(unsafe.Add(mBase, uint32(v1373)+16))
	if v1529 == int32(0) {
		goto L320
	} else {
		goto L336
	}
L336:
	;
	v1532 = *(*int32)(unsafe.Add(mBase, uint32(v1529)+4))
	v1533 = *(*int32)(unsafe.Add(mBase, uint32(l2)+252))
	v1534 = int32(0)
	if base.B2i32(v1532 == v1534)|base.B2i32(v1533 == v1534) != 0 {
		v1579 = v1534
		goto L338
	} else {
		goto L339
	}
L337:
	;
	if v1579 != 0 {
		goto L287
	} else {
		goto L350
	}
L338:
	;
	goto L337
L339:
	;
	v1544 = *(*int32)(unsafe.Add(mBase, uint32(v1532)+4))
	v1545 = *(*int32)(unsafe.Add(mBase, uint32(v1533)+4))
	if v1544 < v1545 {
		goto L340
	} else {
		goto L341
	}
L340:
	;
	v1547 = v1544
	goto L342
L341:
	;
	v1547 = v1545
	goto L342
L342:
	;
	if v1547 <= int32(1) {
		goto L343
	} else {
		goto L344
	}
L343:
	;
	v1550 = int32(1)
	goto L345
L344:
	;
	v1550 = v1547
	goto L345
L345:
	;
	v1551 = int32(8)
	v1556 = int32(0)
	goto L346
L346:
	;
	v1563 = v1556 << (uint(int32(2)) % 32)
	v1565 = *(*int32)(unsafe.Add(mBase, uint32(v1533+v1551+v1563)))
	v1567 = *(*int32)(unsafe.Add(mBase, uint32(v1532+v1551+v1563)))
	v1568 = v1565 & v1567
	v1570 = base.B2i32(v1568 != int32(0))
	if v1568 != 0 {
		v1579 = v1570
		goto L338
	} else {
		goto L348
	}
L347:
	;
	v1579 = v1570
	goto L338
L348:
	;
	v1572 = v1556 + int32(1)
	if v1572 != v1550 {
		v1556 = v1572
		goto L346
	} else {
		goto L349
	}
L349:
	;
	goto L347
L350:
	;
	goto L320
L351:
	;
	v1654 = int32(0)
	v1655 = *(*int32)(unsafe.Add(mBase, uint32(v37)+20))
	if v1655 == v1654 {
		v2776 = v8
		goto L375
	} else {
		goto L376
	}
L352:
	;
	v1597 = *(*int32)(unsafe.Add(mBase, uint32(l2)+52))
	if v1597 == int32(0) {
		v1652 = v1581
		v1653 = v8
		goto L351
	} else {
		goto L353
	}
L353:
	;
	v1600 = *(*int32)(unsafe.Add(mBase, uint32(l1)+72))
	if v1600 != 0 {
		v1652 = v1581
		v1653 = v8
		goto L351
	} else {
		goto L354
	}
L354:
	;
	v1601 = *(*int32)(unsafe.Add(mBase, uint32(v1597)+12))
	v1602 = *(*int32)(unsafe.Add(mBase, uint32(v1601)))
	v1603 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1373)+21)))
	if v1603 != 0 {
		goto L355
	} else {
		goto L356
	}
L355:
	;
	v1652 = v1373
	v1653 = v1602
	goto L351
L356:
	;
	goto L357
L357:
	;
	v1604 = *(*int32)(unsafe.Add(mBase, uint32(l3)+44))
	if v1604 != 0 {
		goto L360
	} else {
		goto L361
	}
L358:
	;
	v1652 = v1648
	v1653 = v1602
	goto L351
L359:
	;
	goto L358
L360:
	;
	v1609 = *(*int32)(unsafe.Add(mBase, uint32(v1604)+4))
	if v1609 <= int32(0) {
		v1648 = int32(0)
		goto L359
	} else {
		goto L363
	}
L361:
	;
	goto L362
L362:
	;
	v1648 = int32(0)
	goto L359
L363:
	;
	v1612 = int32(0)
	if v1612 < v1609 {
		goto L364
	} else {
		goto L365
	}
L364:
	;
	v1615 = v1609
	goto L366
L365:
	;
	v1615 = v1612
	goto L366
L366:
	;
	v1616 = *(*int32)(unsafe.Add(mBase, uint32(v1604)+12))
	v1618 = int32(0)
	goto L367
L367:
	;
	v1626 = *(*int32)(unsafe.Add(mBase, uint32(v1616+v1618<<(uint(int32(2))%32))))
	v1627 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1626)+21)))
	if v1627 == int32(1) {
		goto L369
	} else {
		goto L370
	}
L368:
	;
	goto L362
L369:
	;
	v1630 = *(*int32)(unsafe.Add(mBase, uint32(v1626)+16))
	if v1630 == int32(0) {
		v1648 = v1626
		goto L359
	} else {
		goto L372
	}
L370:
	;
	goto L371
L371:
	;
	v1638 = v1618 + int32(1)
	if v1638 != v1615 {
		v1618 = v1638
		goto L367
	} else {
		goto L374
	}
L372:
	;
	v1633 = *(*int32)(unsafe.Add(mBase, uint32(v1630)+4))
	if v1633 == int32(0) {
		v1648 = v1626
		goto L359
	} else {
		goto L373
	}
L373:
	;
	goto L371
L374:
	;
	goto L368
L375:
	;
	if v2776 == int32(0) {
		goto L287
	} else {
		goto L517
	}
L376:
	;
	v1658 = *(*int32)(unsafe.Add(mBase, uint32(v1655)+4))
	if v1658 == int32(0) {
		v2776 = v8
		goto L375
	} else {
		goto L377
	}
L377:
	;
	v1662 = v1658 << (uint(int32(2)) % 32)
	v1663 = F_palloc(m, v1662)
	mBase = m.M
	v1664 = m.ExcPending
	if v1664 != 0 {
		goto L7
	} else {
		goto L378
	}
L378:
	;
	v1665 = F_palloc(m, v1662)
	mBase = m.M
	v1666 = m.ExcPending
	if v1666 != 0 {
		goto L7
	} else {
		goto L379
	}
L379:
	;
	v1667 = *(*int32)(unsafe.Add(mBase, uint32(v1655)+4))
	if int32(0) < v1667 {
		goto L380
	} else {
		goto L381
	}
L380:
	;
	v1676 = v1654
	v1687 = v8
	goto L383
L381:
	;
	v2131 = v1654
	goto L382
L382:
	;
	v2159 = *(*int32)(unsafe.Add(mBase, uint32(l0)+176))
	if v2159 == int32(0) {
		v2442 = v8
		goto L434
	} else {
		goto L435
	}
L383:
	;
	v1704 = *(*int32)(unsafe.Add(mBase, uint32(v1655)+12))
	v1708 = *(*int32)(unsafe.Add(mBase, uint32(v1704+v1687<<(uint(int32(2))%32))))
	v1709 = *(*int32)(unsafe.Add(mBase, uint32(v1708)+100))
	v1710 = *(*int32)(unsafe.Add(mBase, uint32(v1709)+56))
	if v1710 != 0 {
		goto L385
	} else {
		goto L386
	}
L384:
	;
	v2131 = v2093
	goto L382
L385:
	;
	v1725 = v1710
	goto L388
L386:
	;
	goto L387
L387:
	;
	v1781 = *(*int32)(unsafe.Add(mBase, uint32(v1708)+104))
	v1782 = *(*int32)(unsafe.Add(mBase, uint32(v1781)+56))
	if v1782 != 0 {
		goto L391
	} else {
		goto L392
	}
L388:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v1708)+100)) = v1725
	v1746 = *(*int32)(unsafe.Add(mBase, uint32(v1725)+56))
	if v1746 != 0 {
		v1725 = v1746
		goto L388
	} else {
		goto L390
	}
L389:
	;
	goto L387
L390:
	;
	goto L389
L391:
	;
	v1797 = v1782
	goto L394
L392:
	;
	goto L393
L393:
	;
	v1855 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1708)+120)))
	if v1855 != 0 {
		goto L397
	} else {
		goto L398
	}
L394:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v1708)+104)) = v1797
	v1818 = *(*int32)(unsafe.Add(mBase, uint32(v1797)+56))
	if v1818 != 0 {
		v1797 = v1818
		goto L394
	} else {
		goto L396
	}
L395:
	;
	goto L393
L396:
	;
	goto L395
L397:
	;
	v1856 = int32(100)
	goto L399
L398:
	;
	v1856 = int32(104)
	goto L399
L399:
	;
	v1858 = *(*int32)(unsafe.Add(mBase, uint32(v1708+v1856)))
	v1859 = int32(0)
	if v1859 < v1676 {
		goto L401
	} else {
		goto L402
	}
L400:
	;
	v2122 = v1687 + int32(1)
	v2123 = *(*int32)(unsafe.Add(mBase, uint32(v1655)+4))
	if v2122 < v2123 {
		v1676 = v2093
		v1687 = v2122
		goto L383
	} else {
		goto L432
	}
L401:
	;
	v1876 = v1859
	goto L404
L402:
	;
	goto L403
L403:
	;
	v1938 = *(*int32)(unsafe.Add(mBase, uint32(v1858)+16))
	if v1938 == int32(0) {
		goto L409
	} else {
		goto L410
	}
L404:
	;
	v1899 = *(*int32)(unsafe.Add(mBase, uint32(v1663+v1876<<(uint(int32(2))%32))))
	if v1899 == v1858 {
		v2093 = v1676
		goto L400
	} else {
		goto L406
	}
L405:
	;
	goto L403
L406:
	;
	v1902 = v1876 + int32(1)
	if v1902 != v1676 {
		v1876 = v1902
		goto L404
	} else {
		goto L407
	}
L407:
	;
	goto L405
L408:
	;
	v2080 = v1676 << (uint(int32(2)) % 32)
	*(*int32)(unsafe.Add(mBase, uint32(v1663+v2080))) = v1858
	*(*int32)(unsafe.Add(mBase, uint32(v2080+v1665))) = v2052
	v2093 = v1676 + int32(1)
	goto L400
L409:
	;
	v2052 = int32(0)
	goto L408
L410:
	;
	goto L411
L411:
	;
	v1942 = int32(0)
	v1944 = *(*int32)(unsafe.Add(mBase, uint32(v1938)+4))
	if v1944 <= v1942 {
		v2052 = v1942
		goto L408
	} else {
		goto L412
	}
L412:
	;
	v1954 = v1942
	v1961 = v1942
	goto L413
L413:
	;
	v1981 = *(*int32)(unsafe.Add(mBase, uint32(v1938)+12))
	v1985 = *(*int32)(unsafe.Add(mBase, uint32(v1981+v1961<<(uint(int32(2))%32))))
	v1986 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1985)+12)))
	if v1986 == int32(0) {
		goto L415
	} else {
		goto L416
	}
L414:
	;
	v2052 = v2040
	goto L408
L415:
	;
	v1989 = *(*int32)(unsafe.Add(mBase, uint32(v1985)+8))
	v1990 = *(*int32)(unsafe.Add(mBase, uint32(l1)+8))
	v1991 = int32(0)
	if base.B2i32(v1989 == v1991)|base.B2i32(v1990 == v1991) != 0 {
		v2036 = v1991
		goto L419
	} else {
		goto L420
	}
L416:
	;
	v2040 = v1954
	goto L417
L417:
	;
	v2042 = v1961 + int32(1)
	v2043 = *(*int32)(unsafe.Add(mBase, uint32(v1938)+4))
	if v2042 < v2043 {
		v1954 = v2040
		v1961 = v2042
		goto L413
	} else {
		goto L431
	}
L418:
	;
	v2040 = v1954 + (v2036 ^ int32(1))
	goto L417
L419:
	;
	goto L418
L420:
	;
	v2001 = *(*int32)(unsafe.Add(mBase, uint32(v1989)+4))
	v2002 = *(*int32)(unsafe.Add(mBase, uint32(v1990)+4))
	if v2001 < v2002 {
		goto L421
	} else {
		goto L422
	}
L421:
	;
	v2004 = v2001
	goto L423
L422:
	;
	v2004 = v2002
	goto L423
L423:
	;
	if v2004 <= int32(1) {
		goto L424
	} else {
		goto L425
	}
L424:
	;
	v2007 = int32(1)
	goto L426
L425:
	;
	v2007 = v2004
	goto L426
L426:
	;
	v2008 = int32(8)
	v2013 = int32(0)
	goto L427
L427:
	;
	v2020 = v2013 << (uint(int32(2)) % 32)
	v2022 = *(*int32)(unsafe.Add(mBase, uint32(v1990+v2008+v2020)))
	v2024 = *(*int32)(unsafe.Add(mBase, uint32(v1989+v2008+v2020)))
	v2025 = v2022 & v2024
	v2027 = base.B2i32(v2025 != int32(0))
	if v2025 != 0 {
		v2036 = v2027
		goto L419
	} else {
		goto L429
	}
L428:
	;
	v2036 = v2027
	goto L419
L429:
	;
	v2029 = v2013 + int32(1)
	if v2029 != v2007 {
		v2013 = v2029
		goto L427
	} else {
		goto L430
	}
L430:
	;
	goto L428
L431:
	;
	goto L414
L432:
	;
	goto L384
L433:
	;
	F_pfree(m, v1663)
	mBase = m.M
	v2757 = m.ExcPending
	if v2757 != 0 {
		goto L7
	} else {
		goto L515
	}
L434:
	;
	v2461 = v2131 - int32(1)
	v2464 = int32(3)
	v2465 = v2461 & v2464
	v2466 = int32(2)
	v2488 = v2442
	goto L466
L435:
	;
	v2162 = *(*int32)(unsafe.Add(mBase, uint32(v2159)+4))
	if int32(0) < v2162 {
		goto L437
	} else {
		goto L438
	}
L436:
	;
	if v2185 != v1658 {
		v2442 = v8
		goto L434
	} else {
		goto L464
	}
L437:
	;
	v2165 = *(*int32)(unsafe.Add(mBase, uint32(v2159)+12))
	v2185 = int32(0)
	goto L440
L438:
	;
	goto L439
L439:
	;
	v2288 = F_list_copy(m, v2159)
	mBase = m.M
	v2289 = m.ExcPending
	if v2289 != 0 {
		goto L7
	} else {
		goto L450
	}
L440:
	;
	if v2131 <= int32(0) {
		v2442 = v8
		goto L434
	} else {
		goto L442
	}
L441:
	;
	goto L439
L442:
	;
	v2206 = *(*int32)(unsafe.Add(mBase, uint32(v2165+v2185<<(uint(int32(2))%32))))
	v2207 = *(*int32)(unsafe.Add(mBase, uint32(v2206)+4))
	v2223 = int32(0)
	goto L443
L443:
	;
	v2246 = *(*int32)(unsafe.Add(mBase, uint32(v1663+v2223<<(uint(int32(2))%32))))
	if v2207 != v2246 {
		goto L445
	} else {
		goto L446
	}
L444:
	;
	v2252 = v2185 + int32(1)
	if v2252 != v2162 {
		v2185 = v2252
		goto L440
	} else {
		goto L449
	}
L445:
	;
	v2249 = v2223 + int32(1)
	if v2131 != v2249 {
		v2223 = v2249
		goto L443
	} else {
		goto L448
	}
L446:
	;
	goto L447
L447:
	;
	goto L444
L448:
	;
	goto L436
L449:
	;
	goto L441
L450:
	;
	v2290 = *(*int32)(unsafe.Add(mBase, uint32(l0)+176))
	if v2290 == int32(0) {
		v2442 = v2288
		goto L434
	} else {
		goto L451
	}
L451:
	;
	v2293 = *(*int32)(unsafe.Add(mBase, uint32(v2290)+4))
	if v2293 <= int32(0) {
		v2442 = v2288
		goto L434
	} else {
		goto L452
	}
L452:
	;
	v2296 = int32(0)
	v2317 = v2296
	goto L453
L453:
	;
	if v2131 <= v2296 {
		goto L455
	} else {
		goto L456
	}
L454:
	;
	v2442 = v2288
	goto L434
L455:
	;
	v2420 = v2317 + int32(1)
	v2421 = *(*int32)(unsafe.Add(mBase, uint32(v2290)+4))
	if v2420 < v2421 {
		v2317 = v2420
		goto L453
	} else {
		goto L463
	}
L456:
	;
	v2333 = *(*int32)(unsafe.Add(mBase, uint32(v2290)+12))
	v2337 = *(*int32)(unsafe.Add(mBase, uint32(v2333+v2317<<(uint(int32(2))%32))))
	v2338 = *(*int32)(unsafe.Add(mBase, uint32(v2337)+4))
	v2354 = int32(0)
	goto L457
L457:
	;
	v2375 = v2354 << (uint(int32(2)) % 32)
	v2377 = *(*int32)(unsafe.Add(mBase, uint32(v1663+v2375)))
	if v2338 == v2377 {
		goto L459
	} else {
		goto L460
	}
L458:
	;
	goto L455
L459:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v2375+v1665))) = int32(-1)
	goto L455
L460:
	;
	goto L461
L461:
	;
	v2383 = v2354 + int32(1)
	if v2383 != v2131 {
		v2354 = v2383
		goto L457
	} else {
		goto L462
	}
L462:
	;
	goto L458
L463:
	;
	goto L454
L464:
	;
	v2424 = F_list_copy_head(m, v2159, v1658)
	mBase = m.M
	v2425 = m.ExcPending
	if v2425 != 0 {
		goto L7
	} else {
		goto L465
	}
L465:
	;
	v2738 = v2424
	goto L433
L466:
	;
	v2506 = *(*int32)(unsafe.Add(mBase, uint32(v1665)))
	if v2131 < v2466 {
		goto L469
	} else {
		goto L470
	}
L468:
	;
	if v2689 < int32(0) {
		v2738 = v2488
		goto L433
	} else {
		goto L512
	}
L469:
	;
	v2676 = int32(0)
	v2689 = v2506
	goto L468
L470:
	;
	goto L471
L471:
	;
	v2508 = int32(1)
	v2509 = int32(0)
	if base.B2i32(base.Ui32(v2131-v2466) < base.Ui32(v2464)) == v2509 {
		goto L472
	} else {
		goto L473
	}
L472:
	;
	v2521 = v2509
	v2529 = v2508
	v2533 = v2509
	v2534 = v2506
	goto L475
L473:
	;
	v2597 = v2509
	v2604 = v2508
	v2609 = v2506
	goto L474
L474:
	;
	v2630 = v2597
	v2638 = v2604
	v2643 = v2609
	v2646 = v2509
	goto L503
L475:
	;
	v2550 = v2529 + int32(3)
	v2551 = int32(2)
	v2554 = *(*int32)(unsafe.Add(mBase, uint32(v1665+v2550<<(uint(v2551)%32))))
	v2556 = v2529 + v2551
	v2560 = *(*int32)(unsafe.Add(mBase, uint32(v1665+v2556<<(uint(v2551)%32))))
	v2562 = v2529 + int32(1)
	v2566 = *(*int32)(unsafe.Add(mBase, uint32(v1665+v2562<<(uint(v2551)%32))))
	v2570 = *(*int32)(unsafe.Add(mBase, uint32(v1665+v2529<<(uint(v2551)%32))))
	v2571 = base.B2i32(v2534 < v2570)
	if v2534 < v2570 {
		goto L477
	} else {
		goto L478
	}
L476:
	;
	if v2465 == int32(0) {
		v2676 = v2582
		v2689 = v2578
		goto L468
	} else {
		goto L502
	}
L477:
	;
	v2572 = v2570
	goto L479
L478:
	;
	v2572 = v2534
	goto L479
L479:
	;
	v2573 = base.B2i32(v2572 < v2566)
	if v2572 < v2566 {
		goto L480
	} else {
		goto L481
	}
L480:
	;
	v2574 = v2566
	goto L482
L481:
	;
	v2574 = v2572
	goto L482
L482:
	;
	v2575 = base.B2i32(v2574 < v2560)
	if v2574 < v2560 {
		goto L483
	} else {
		goto L484
	}
L483:
	;
	v2576 = v2560
	goto L485
L484:
	;
	v2576 = v2574
	goto L485
L485:
	;
	v2577 = base.B2i32(v2576 < v2554)
	if v2576 < v2554 {
		goto L486
	} else {
		goto L487
	}
L486:
	;
	v2578 = v2554
	goto L488
L487:
	;
	v2578 = v2576
	goto L488
L488:
	;
	if v2534 < v2570 {
		goto L489
	} else {
		goto L490
	}
L489:
	;
	v2579 = v2529
	goto L491
L490:
	;
	v2579 = v2521
	goto L491
L491:
	;
	if v2572 < v2566 {
		goto L492
	} else {
		goto L493
	}
L492:
	;
	v2580 = v2562
	goto L494
L493:
	;
	v2580 = v2579
	goto L494
L494:
	;
	if v2574 < v2560 {
		goto L495
	} else {
		goto L496
	}
L495:
	;
	v2581 = v2556
	goto L497
L496:
	;
	v2581 = v2580
	goto L497
L497:
	;
	if v2576 < v2554 {
		goto L498
	} else {
		goto L499
	}
L498:
	;
	v2582 = v2550
	goto L500
L499:
	;
	v2582 = v2581
	goto L500
L500:
	;
	v2583 = int32(4)
	v2584 = v2529 + v2583
	v2586 = v2533 + v2583
	if v2586 != v2461&int32(-4) {
		v2521 = v2582
		v2529 = v2584
		v2533 = v2586
		v2534 = v2578
		goto L475
	} else {
		goto L501
	}
L501:
	;
	goto L476
L502:
	;
	v2597 = v2582
	v2604 = v2584
	v2609 = v2578
	goto L474
L503:
	;
	v2661 = *(*int32)(unsafe.Add(mBase, uint32(v1665+v2638<<(uint(int32(2))%32))))
	v2662 = base.B2i32(v2643 < v2661)
	if v2643 < v2661 {
		goto L505
	} else {
		goto L506
	}
L504:
	;
	v2676 = v2664
	v2689 = v2663
	goto L468
L505:
	;
	v2663 = v2661
	goto L507
L506:
	;
	v2663 = v2643
	goto L507
L507:
	;
	if v2643 < v2661 {
		goto L508
	} else {
		goto L509
	}
L508:
	;
	v2664 = v2638
	goto L510
L509:
	;
	v2664 = v2630
	goto L510
L510:
	;
	v2665 = int32(1)
	v2668 = v2646 + v2665
	if v2668 != v2465 {
		v2630 = v2664
		v2638 = v2638 + v2665
		v2643 = v2663
		v2646 = v2668
		goto L503
	} else {
		goto L511
	}
L511:
	;
	goto L504
L512:
	;
	v2707 = v2676 << (uint(int32(2)) % 32)
	v2709 = *(*int32)(unsafe.Add(mBase, uint32(v1663+v2707)))
	*(*int32)(unsafe.Add(mBase, uint32(v2707+v1665))) = int32(-1)
	v2713 = *(*int32)(unsafe.Add(mBase, uint32(v2709)+4))
	v2714 = *(*int32)(unsafe.Add(mBase, uint32(v2713)+12))
	v2715 = *(*int32)(unsafe.Add(mBase, uint32(v2714)))
	v2718 = F_make_canonical_pathkey(m, l0, v2709, v2715, int32(1), int32(0))
	mBase = m.M
	v2719 = m.ExcPending
	if v2719 != 0 {
		goto L7
	} else {
		goto L513
	}
L513:
	;
	v2720 = F_lappend(m, v2488, v2718)
	mBase = m.M
	v2721 = m.ExcPending
	if v2721 != 0 {
		goto L7
	} else {
		goto L514
	}
L514:
	;
	v2488 = v2720
	goto L466
L515:
	;
	F_pfree(m, v1665)
	mBase = m.M
	v2759 = m.ExcPending
	if v2759 != 0 {
		goto L7
	} else {
		goto L516
	}
L516:
	;
	v2776 = v2738
	goto L375
L517:
	;
	v2796 = *(*int32)(unsafe.Add(mBase, uint32(v2776)+4))
	if v2796 <= int32(0) {
		goto L287
	} else {
		goto L518
	}
L518:
	;
	v2799 = *(*int32)(unsafe.Add(mBase, uint32(v37)+20))
	v2800 = F_find_mergeclauses_for_outer_pathkeys(m, v2776, v2799)
	mBase = m.M
	v2801 = m.ExcPending
	if v2801 != 0 {
		goto L7
	} else {
		goto L519
	}
L519:
	;
	v2802 = F_make_inner_pathkeys_for_merge(m, l0, v2800, v2776)
	mBase = m.M
	v2803 = m.ExcPending
	if v2803 != 0 {
		goto L7
	} else {
		goto L520
	}
L520:
	;
	v2804 = F_build_join_pathkeys(m, l0, l1, v136, v2776)
	mBase = m.M
	v2805 = m.ExcPending
	if v2805 != 0 {
		goto L7
	} else {
		goto L521
	}
L521:
	;
	v2807 = v35 + int32(-48)
	F_try_mergejoin_path(m, l0, l1, v1374, v1373, v2804, v2800, v2776, v2802, v136, v2807, int32(0))
	mBase = m.M
	v2810 = m.ExcPending
	if v2810 != 0 {
		goto L7
	} else {
		goto L522
	}
L522:
	;
	v2811 = int32(0)
	v2815 = base.B2i32(v1653 != v2811) & base.B2i32(v1652 != v2811)
	if v2815 != 0 {
		goto L523
	} else {
		goto L524
	}
L523:
	;
	F_try_partial_mergejoin_path(m, l0, l1, v1653, v1652, v2804, v2800, v2776, v2802, v136, v2807)
	mBase = m.M
	v2817 = m.ExcPending
	if v2817 != 0 {
		goto L7
	} else {
		goto L526
	}
L524:
	;
	goto L525
L525:
	;
	v2818 = *(*int32)(unsafe.Add(mBase, uint32(v2776)+4))
	if v2818 < int32(2) {
		goto L287
	} else {
		goto L527
	}
L526:
	;
	goto L525
L527:
	;
	v2828 = int32(1)
	goto L528
L528:
	;
	v2856 = *(*int32)(unsafe.Add(mBase, uint32(v2776)+12))
	v2860 = *(*int32)(unsafe.Add(mBase, uint32(v2856+v2828<<(uint(int32(2))%32))))
	v2861 = F_list_copy(m, v2776)
	mBase = m.M
	v2862 = m.ExcPending
	if v2862 != 0 {
		goto L7
	} else {
		goto L530
	}
L529:
	;
	goto L287
L530:
	;
	v2863 = F_list_delete_nth_cell(m, v2861, v2828)
	mBase = m.M
	v2864 = m.ExcPending
	if v2864 != 0 {
		goto L7
	} else {
		goto L531
	}
L531:
	;
	v2865 = F_lcons(m, v2860, v2863)
	mBase = m.M
	v2866 = m.ExcPending
	if v2866 != 0 {
		goto L7
	} else {
		goto L532
	}
L532:
	;
	v2867 = *(*int32)(unsafe.Add(mBase, uint32(v37)+20))
	v2868 = F_find_mergeclauses_for_outer_pathkeys(m, v2865, v2867)
	mBase = m.M
	v2869 = m.ExcPending
	if v2869 != 0 {
		goto L7
	} else {
		goto L533
	}
L533:
	;
	v2870 = F_make_inner_pathkeys_for_merge(m, l0, v2868, v2865)
	mBase = m.M
	v2871 = m.ExcPending
	if v2871 != 0 {
		goto L7
	} else {
		goto L534
	}
L534:
	;
	v2872 = F_build_join_pathkeys(m, l0, l1, v136, v2865)
	mBase = m.M
	v2873 = m.ExcPending
	if v2873 != 0 {
		goto L7
	} else {
		goto L535
	}
L535:
	;
	v2875 = v35 + int32(-48)
	F_try_mergejoin_path(m, l0, l1, v1374, v1373, v2872, v2868, v2865, v2870, v136, v2875, int32(0))
	mBase = m.M
	v2878 = m.ExcPending
	if v2878 != 0 {
		goto L7
	} else {
		goto L536
	}
L536:
	;
	if v2815 != 0 {
		goto L537
	} else {
		goto L538
	}
L537:
	;
	F_try_partial_mergejoin_path(m, l0, l1, v1653, v1652, v2872, v2868, v2865, v2870, v136, v2875)
	mBase = m.M
	v2880 = m.ExcPending
	if v2880 != 0 {
		goto L7
	} else {
		goto L540
	}
L538:
	;
	goto L539
L539:
	;
	v2882 = v2828 + int32(1)
	v2883 = *(*int32)(unsafe.Add(mBase, uint32(v2776)+4))
	if v2882 < v2883 {
		v2828 = v2882
		goto L528
	} else {
		goto L541
	}
L540:
	;
	goto L539
L541:
	;
	goto L529
L542:
	;
	v2921 = *(*int32)(unsafe.Add(mBase, uint32(l3)+60))
	v2923 = int32(1) << (uint(v136) % 32)
	if v2923&int32(51) != 0 {
		goto L544
	} else {
		goto L545
	}
L543:
	;
	v2936 = *(*int32)(unsafe.Add(mBase, uint32(v2921)+16))
	if v2936 == int32(0) {
		v3102 = v2921
		goto L548
	} else {
		goto L549
	}
L544:
	;
	v2934 = int32(0)
	v2935 = int32(1)
	goto L543
L545:
	;
	goto L546
L546:
	;
	if v2923&int32(140) == int32(0) {
		goto L285
	} else {
		goto L547
	}
L547:
	;
	v2934 = int32(1)
	v2935 = int32(0)
	goto L543
L548:
	;
	v3103 = int32(0)
	if v2935 == v3103 {
		v3124 = v3103
		goto L595
	} else {
		goto L596
	}
L549:
	;
	v2939 = *(*int32)(unsafe.Add(mBase, uint32(v2936)+4))
	v2940 = *(*int32)(unsafe.Add(mBase, uint32(l2)+8))
	v2941 = int32(0)
	if base.B2i32(v2939 == v2941)|base.B2i32(v2940 == v2941) != 0 {
		v2986 = v2941
		goto L551
	} else {
		goto L552
	}
L550:
	;
	if v2986 == int32(0) {
		goto L563
	} else {
		goto L564
	}
L551:
	;
	goto L550
L552:
	;
	v2951 = *(*int32)(unsafe.Add(mBase, uint32(v2939)+4))
	v2952 = *(*int32)(unsafe.Add(mBase, uint32(v2940)+4))
	if v2951 < v2952 {
		goto L553
	} else {
		goto L554
	}
L553:
	;
	v2954 = v2951
	goto L555
L554:
	;
	v2954 = v2952
	goto L555
L555:
	;
	if v2954 <= int32(1) {
		goto L556
	} else {
		goto L557
	}
L556:
	;
	v2957 = int32(1)
	goto L558
L557:
	;
	v2957 = v2954
	goto L558
L558:
	;
	v2958 = int32(8)
	v2963 = int32(0)
	goto L559
L559:
	;
	v2970 = v2963 << (uint(int32(2)) % 32)
	v2972 = *(*int32)(unsafe.Add(mBase, uint32(v2940+v2958+v2970)))
	v2974 = *(*int32)(unsafe.Add(mBase, uint32(v2939+v2958+v2970)))
	v2975 = v2972 & v2974
	v2977 = base.B2i32(v2975 != int32(0))
	if v2975 != 0 {
		v2986 = v2977
		goto L551
	} else {
		goto L561
	}
L560:
	;
	v2986 = v2977
	goto L551
L561:
	;
	v2979 = v2963 + int32(1)
	if v2979 != v2957 {
		v2963 = v2979
		goto L559
	} else {
		goto L562
	}
L562:
	;
	goto L560
L563:
	;
	v2989 = *(*int32)(unsafe.Add(mBase, uint32(v2921)+16))
	if v2989 == int32(0) {
		v3102 = v2921
		goto L548
	} else {
		goto L566
	}
L564:
	;
	goto L565
L565:
	;
	v3043 = int32(0)
	if v136 != 0 {
		v3102 = v3043
		goto L548
	} else {
		goto L581
	}
L566:
	;
	v2992 = *(*int32)(unsafe.Add(mBase, uint32(v2989)+4))
	v2993 = *(*int32)(unsafe.Add(mBase, uint32(l2)+252))
	v2994 = int32(0)
	if base.B2i32(v2992 == v2994)|base.B2i32(v2993 == v2994) != 0 {
		v3039 = v2994
		goto L568
	} else {
		goto L569
	}
L567:
	;
	if v3039 == int32(0) {
		v3102 = v2921
		goto L548
	} else {
		goto L580
	}
L568:
	;
	goto L567
L569:
	;
	v3004 = *(*int32)(unsafe.Add(mBase, uint32(v2992)+4))
	v3005 = *(*int32)(unsafe.Add(mBase, uint32(v2993)+4))
	if v3004 < v3005 {
		goto L570
	} else {
		goto L571
	}
L570:
	;
	v3007 = v3004
	goto L572
L571:
	;
	v3007 = v3005
	goto L572
L572:
	;
	if v3007 <= int32(1) {
		goto L573
	} else {
		goto L574
	}
L573:
	;
	v3010 = int32(1)
	goto L575
L574:
	;
	v3010 = v3007
	goto L575
L575:
	;
	v3011 = int32(8)
	v3016 = int32(0)
	goto L576
L576:
	;
	v3023 = v3016 << (uint(int32(2)) % 32)
	v3025 = *(*int32)(unsafe.Add(mBase, uint32(v2993+v3011+v3023)))
	v3027 = *(*int32)(unsafe.Add(mBase, uint32(v2992+v3011+v3023)))
	v3028 = v3025 & v3027
	v3030 = base.B2i32(v3028 != int32(0))
	if v3028 != 0 {
		v3039 = v3030
		goto L568
	} else {
		goto L578
	}
L577:
	;
	v3039 = v3030
	goto L568
L578:
	;
	v3032 = v3016 + int32(1)
	if v3032 != v3010 {
		v3016 = v3032
		goto L576
	} else {
		goto L579
	}
L579:
	;
	goto L577
L580:
	;
	goto L565
L581:
	;
	v3044 = *(*int32)(unsafe.Add(mBase, uint32(v37)+28))
	v3045 = *(*int32)(unsafe.Add(mBase, uint32(v3044)+20))
	if v3045 != int32(4) {
		v3102 = v3043
		goto L548
	} else {
		goto L582
	}
L582:
	;
	v3048 = *(*int32)(unsafe.Add(mBase, uint32(v3044)+16))
	v3049 = *(*int32)(unsafe.Add(mBase, uint32(l3)+8))
	v3050 = int32(0)
	if base.B2i32(v3048 == v3050)|base.B2i32(v3049 == v3050) != 0 {
		v3096 = base.B2i32(v3048|v3049 == v3050)
		goto L584
	} else {
		goto L585
	}
L583:
	;
	if v3096 != 0 {
		goto L285
	} else {
		goto L594
	}
L584:
	;
	goto L583
L585:
	;
	v3064 = *(*int32)(unsafe.Add(mBase, uint32(v3048)+4))
	v3065 = *(*int32)(unsafe.Add(mBase, uint32(v3049)+4))
	if v3064 != v3065 {
		v3096 = int32(0)
		goto L584
	} else {
		goto L586
	}
L586:
	;
	v3067 = int32(1)
	if v3064 <= v3067 {
		goto L587
	} else {
		goto L588
	}
L587:
	;
	v3070 = v3067
	goto L589
L588:
	;
	v3070 = v3064
	goto L589
L589:
	;
	v3071 = int32(8)
	v3076 = int32(0)
	goto L590
L590:
	;
	v3084 = v3076 << (uint(int32(2)) % 32)
	v3086 = *(*int32)(unsafe.Add(mBase, uint32(v3048+v3071+v3084)))
	v3088 = *(*int32)(unsafe.Add(mBase, uint32(v3049+v3071+v3084)))
	v3089 = base.B2i32(v3086 == v3088)
	if v3086 != v3088 {
		v3096 = v3089
		goto L584
	} else {
		goto L592
	}
L591:
	;
	v3096 = v3089
	goto L584
L592:
	;
	v3092 = v3076 + int32(1)
	if v3092 != v3070 {
		v3076 = v3092
		goto L590
	} else {
		goto L593
	}
L593:
	;
	goto L591
L594:
	;
	v3102 = v3043
	goto L548
L595:
	;
	v3125 = *(*int32)(unsafe.Add(mBase, uint32(l2)+44))
	if v3125 == int32(0) {
		goto L601
	} else {
		goto L602
	}
L596:
	;
	v3108 = *(*int64)(unsafe.Add(mBase, uint32(v37)+56))
	v3109 = int64(262656)
	if base.B2i32(v3102 == int32(0))|base.B2i32(v3108&v3109 != v3109) != 0 {
		v3124 = v3103
		goto L595
	} else {
		goto L597
	}
L597:
	;
	v3114 = *(*int32)(unsafe.Add(mBase, uint32(v3102)+4))
	v3116 = v3114 - int32(352)
	goto L598
L598:
	;
	if base.B2i32(base.Ui32(v3116) < base.Ui32(int32(15)))&int32(base.Ui32(int32(_a_F_add_paths_to_joinrel_0))>>(uint(v3116)%32)) != 0 {
		v3124 = v3103
		goto L595
	} else {
		goto L599
	}
L599:
	;
	v3122 = F_create_material_path(m, l3, v3102)
	mBase = m.M
	v3123 = m.ExcPending
	if v3123 != 0 {
		goto L7
	} else {
		goto L600
	}
L600:
	;
	v3124 = v3122
	goto L595
L601:
	;
	v3493 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l1)+26)))
	v3494 = int32(1)
	if v3494<<(uint(v136)%32)&int32(140) != 0 {
		goto L658
	} else {
		goto L659
	}
L602:
	;
	v3128 = *(*int32)(unsafe.Add(mBase, uint32(v3125)+4))
	if v3128 <= int32(0) {
		goto L601
	} else {
		goto L603
	}
L603:
	;
	v3137 = int32(0)
	goto L604
L604:
	;
	v3166 = *(*int32)(unsafe.Add(mBase, uint32(v3125)+12))
	v3170 = *(*int32)(unsafe.Add(mBase, uint32(v3166+v3137<<(uint(int32(2))%32))))
	v3171 = *(*int32)(unsafe.Add(mBase, uint32(v3170)+16))
	if v3171 == int32(0) {
		goto L607
	} else {
		goto L608
	}
L605:
	;
	goto L601
L606:
	;
	v3456 = v3137 + int32(1)
	v3457 = *(*int32)(unsafe.Add(mBase, uint32(v3125)+4))
	if v3456 < v3457 {
		v3137 = v3456
		goto L604
	} else {
		goto L657
	}
L607:
	;
	v3274 = *(*int32)(unsafe.Add(mBase, uint32(v3170)+64))
	v3275 = F_build_join_pathkeys(m, l0, l1, v136, v3274)
	mBase = m.M
	v3276 = m.ExcPending
	if v3276 != 0 {
		goto L7
	} else {
		goto L638
	}
L608:
	;
	v3174 = *(*int32)(unsafe.Add(mBase, uint32(v3171)+4))
	v3175 = *(*int32)(unsafe.Add(mBase, uint32(l3)+8))
	v3176 = int32(0)
	if base.B2i32(v3174 == v3176)|base.B2i32(v3175 == v3176) != 0 {
		v3221 = v3176
		goto L610
	} else {
		goto L611
	}
L609:
	;
	if v3221 != 0 {
		goto L606
	} else {
		goto L622
	}
L610:
	;
	goto L609
L611:
	;
	v3186 = *(*int32)(unsafe.Add(mBase, uint32(v3174)+4))
	v3187 = *(*int32)(unsafe.Add(mBase, uint32(v3175)+4))
	if v3186 < v3187 {
		goto L612
	} else {
		goto L613
	}
L612:
	;
	v3189 = v3186
	goto L614
L613:
	;
	v3189 = v3187
	goto L614
L614:
	;
	if v3189 <= int32(1) {
		goto L615
	} else {
		goto L616
	}
L615:
	;
	v3192 = int32(1)
	goto L617
L616:
	;
	v3192 = v3189
	goto L617
L617:
	;
	v3193 = int32(8)
	v3198 = int32(0)
	goto L618
L618:
	;
	v3205 = v3198 << (uint(int32(2)) % 32)
	v3207 = *(*int32)(unsafe.Add(mBase, uint32(v3175+v3193+v3205)))
	v3209 = *(*int32)(unsafe.Add(mBase, uint32(v3174+v3193+v3205)))
	v3210 = v3207 & v3209
	v3212 = base.B2i32(v3210 != int32(0))
	if v3210 != 0 {
		v3221 = v3212
		goto L610
	} else {
		goto L620
	}
L619:
	;
	v3221 = v3212
	goto L610
L620:
	;
	v3214 = v3198 + int32(1)
	if v3214 != v3192 {
		v3198 = v3214
		goto L618
	} else {
		goto L621
	}
L621:
	;
	goto L619
L622:
	;
	v3222 = *(*int32)(unsafe.Add(mBase, uint32(v3170)+16))
	if v3222 == int32(0) {
		goto L607
	} else {
		goto L623
	}
L623:
	;
	v3225 = *(*int32)(unsafe.Add(mBase, uint32(v3222)+4))
	v3226 = *(*int32)(unsafe.Add(mBase, uint32(l3)+252))
	v3227 = int32(0)
	if base.B2i32(v3225 == v3227)|base.B2i32(v3226 == v3227) != 0 {
		v3272 = v3227
		goto L625
	} else {
		goto L626
	}
L624:
	;
	if v3272 != 0 {
		goto L606
	} else {
		goto L637
	}
L625:
	;
	goto L624
L626:
	;
	v3237 = *(*int32)(unsafe.Add(mBase, uint32(v3225)+4))
	v3238 = *(*int32)(unsafe.Add(mBase, uint32(v3226)+4))
	if v3237 < v3238 {
		goto L627
	} else {
		goto L628
	}
L627:
	;
	v3240 = v3237
	goto L629
L628:
	;
	v3240 = v3238
	goto L629
L629:
	;
	if v3240 <= int32(1) {
		goto L630
	} else {
		goto L631
	}
L630:
	;
	v3243 = int32(1)
	goto L632
L631:
	;
	v3243 = v3240
	goto L632
L632:
	;
	v3244 = int32(8)
	v3249 = int32(0)
	goto L633
L633:
	;
	v3256 = v3249 << (uint(int32(2)) % 32)
	v3258 = *(*int32)(unsafe.Add(mBase, uint32(v3226+v3244+v3256)))
	v3260 = *(*int32)(unsafe.Add(mBase, uint32(v3225+v3244+v3256)))
	v3261 = v3258 & v3260
	v3263 = base.B2i32(v3261 != int32(0))
	if v3261 != 0 {
		v3272 = v3263
		goto L625
	} else {
		goto L635
	}
L634:
	;
	v3272 = v3263
	goto L625
L635:
	;
	v3265 = v3249 + int32(1)
	if v3265 != v3243 {
		v3249 = v3265
		goto L633
	} else {
		goto L636
	}
L636:
	;
	goto L634
L637:
	;
	goto L607
L638:
	;
	if v2935 == int32(0) {
		goto L639
	} else {
		goto L640
	}
L639:
	;
	if v3102 == int32(0) {
		goto L606
	} else {
		goto L655
	}
L640:
	;
	v3279 = *(*int32)(unsafe.Add(mBase, uint32(l3)+64))
	if v3279 == int32(0) {
		goto L641
	} else {
		goto L642
	}
L641:
	;
	if v3124 == int32(0) {
		goto L639
	} else {
		goto L653
	}
L642:
	;
	v3282 = int32(0)
	v3283 = *(*int32)(unsafe.Add(mBase, uint32(v3279)+4))
	if v3283 <= v3282 {
		goto L641
	} else {
		goto L643
	}
L643:
	;
	v3292 = v3282
	goto L644
L644:
	;
	v3320 = *(*int32)(unsafe.Add(mBase, uint32(v3279)+12))
	v3324 = *(*int32)(unsafe.Add(mBase, uint32(v3320+v3292<<(uint(int32(2))%32))))
	v3327 = v35 + int32(-48)
	F_try_nestloop_path(m, l0, l1, v3170, v3324, v3275, v136, int64(256), v3327)
	mBase = m.M
	v3329 = m.ExcPending
	if v3329 != 0 {
		goto L7
	} else {
		goto L646
	}
L645:
	;
	goto L641
L646:
	;
	v3330 = F_get_memoize_path(m, l0, l3, l2, v3324, v3170, v136, v3327)
	mBase = m.M
	v3331 = m.ExcPending
	if v3331 != 0 {
		goto L7
	} else {
		goto L647
	}
L647:
	;
	if v3330 != 0 {
		goto L648
	} else {
		goto L649
	}
L648:
	;
	F_try_nestloop_path(m, l0, l1, v3170, v3330, v3275, v136, int64(1024), v3327)
	mBase = m.M
	v3334 = m.ExcPending
	if v3334 != 0 {
		goto L7
	} else {
		goto L651
	}
L649:
	;
	goto L650
L650:
	;
	v3336 = v3292 + int32(1)
	v3337 = *(*int32)(unsafe.Add(mBase, uint32(v3279)+4))
	if v3336 < v3337 {
		v3292 = v3336
		goto L644
	} else {
		goto L652
	}
L651:
	;
	goto L650
L652:
	;
	goto L645
L653:
	;
	F_try_nestloop_path(m, l0, l1, v3170, v3124, v3275, v136, int64(512), v35+int32(-48))
	mBase = m.M
	v3379 = m.ExcPending
	if v3379 != 0 {
		goto L7
	} else {
		goto L654
	}
L654:
	;
	goto L639
L655:
	;
	F_generate_mergejoin_paths(m, l0, l1, l3, v3170, v136, v35+int32(-48), v2934, v3102, v3275, int32(0))
	mBase = m.M
	v3420 = m.ExcPending
	if v3420 != 0 {
		goto L7
	} else {
		goto L656
	}
L656:
	;
	goto L606
L657:
	;
	goto L605
L658:
	;
	v3503 = base.B2i32(base.Ui32(v136) <= base.Ui32(int32(7)))
	goto L660
L659:
	;
	v3503 = int32(0)
	goto L660
L660:
	;
	if base.B2i32(v3493 != v3494)|v3503 != 0 {
		goto L285
	} else {
		goto L661
	}
L661:
	;
	v3505 = *(*int32)(unsafe.Add(mBase, uint32(l2)+52))
	if v3505 == int32(0) {
		goto L285
	} else {
		goto L662
	}
L662:
	;
	v3508 = *(*int32)(unsafe.Add(mBase, uint32(l1)+72))
	if v3508 != 0 {
		goto L285
	} else {
		goto L663
	}
L663:
	;
	if v2935 != 0 {
		goto L664
	} else {
		goto L665
	}
L664:
	;
	v3509 = int32(0)
	v3512 = v35 + int32(-48)
	v3513 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v3512)+41)))
	if v3513&int32(2) == v3509 {
		v3637 = v3509
		goto L667
	} else {
		goto L668
	}
L665:
	;
	goto L666
L666:
	;
	if v3102 != 0 {
		goto L728
	} else {
		goto L729
	}
L667:
	;
	v3638 = *(*int32)(unsafe.Add(mBase, uint32(l2)+52))
	if v3638 == int32(0) {
		goto L704
	} else {
		goto L705
	}
L668:
	;
	v3518 = *(*int32)(unsafe.Add(mBase, uint32(l3)+60))
	v3519 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v3518)+21)))
	if v3519 != int32(1) {
		v3637 = v3509
		goto L667
	} else {
		goto L669
	}
L669:
	;
	v3522 = *(*int32)(unsafe.Add(mBase, uint32(v3518)+16))
	if v3522 == int32(0) {
		goto L670
	} else {
		goto L671
	}
L670:
	;
	v3625 = *(*int32)(unsafe.Add(mBase, uint32(v3518)+4))
	v3627 = v3625 - int32(352)
	goto L701
L671:
	;
	v3525 = *(*int32)(unsafe.Add(mBase, uint32(v3522)+4))
	v3526 = *(*int32)(unsafe.Add(mBase, uint32(l2)+8))
	v3527 = int32(0)
	if base.B2i32(v3525 == v3527)|base.B2i32(v3526 == v3527) != 0 {
		v3572 = v3527
		goto L673
	} else {
		goto L674
	}
L672:
	;
	if v3572 != 0 {
		v3637 = v3509
		goto L667
	} else {
		goto L685
	}
L673:
	;
	goto L672
L674:
	;
	v3537 = *(*int32)(unsafe.Add(mBase, uint32(v3525)+4))
	v3538 = *(*int32)(unsafe.Add(mBase, uint32(v3526)+4))
	if v3537 < v3538 {
		goto L675
	} else {
		goto L676
	}
L675:
	;
	v3540 = v3537
	goto L677
L676:
	;
	v3540 = v3538
	goto L677
L677:
	;
	if v3540 <= int32(1) {
		goto L678
	} else {
		goto L679
	}
L678:
	;
	v3543 = int32(1)
	goto L680
L679:
	;
	v3543 = v3540
	goto L680
L680:
	;
	v3544 = int32(8)
	v3549 = int32(0)
	goto L681
L681:
	;
	v3556 = v3549 << (uint(int32(2)) % 32)
	v3558 = *(*int32)(unsafe.Add(mBase, uint32(v3526+v3544+v3556)))
	v3560 = *(*int32)(unsafe.Add(mBase, uint32(v3525+v3544+v3556)))
	v3561 = v3558 & v3560
	v3563 = base.B2i32(v3561 != int32(0))
	if v3561 != 0 {
		v3572 = v3563
		goto L673
	} else {
		goto L683
	}
L682:
	;
	v3572 = v3563
	goto L673
L683:
	;
	v3565 = v3549 + int32(1)
	if v3565 != v3543 {
		v3549 = v3565
		goto L681
	} else {
		goto L684
	}
L684:
	;
	goto L682
L685:
	;
	v3573 = *(*int32)(unsafe.Add(mBase, uint32(v3518)+16))
	if v3573 == int32(0) {
		goto L670
	} else {
		goto L686
	}
L686:
	;
	v3576 = *(*int32)(unsafe.Add(mBase, uint32(v3573)+4))
	v3577 = *(*int32)(unsafe.Add(mBase, uint32(l2)+252))
	v3578 = int32(0)
	if base.B2i32(v3576 == v3578)|base.B2i32(v3577 == v3578) != 0 {
		v3623 = v3578
		goto L688
	} else {
		goto L689
	}
L687:
	;
	if v3623 != 0 {
		v3637 = v3509
		goto L667
	} else {
		goto L700
	}
L688:
	;
	goto L687
L689:
	;
	v3588 = *(*int32)(unsafe.Add(mBase, uint32(v3576)+4))
	v3589 = *(*int32)(unsafe.Add(mBase, uint32(v3577)+4))
	if v3588 < v3589 {
		goto L690
	} else {
		goto L691
	}
L690:
	;
	v3591 = v3588
	goto L692
L691:
	;
	v3591 = v3589
	goto L692
L692:
	;
	if v3591 <= int32(1) {
		goto L693
	} else {
		goto L694
	}
L693:
	;
	v3594 = int32(1)
	goto L695
L694:
	;
	v3594 = v3591
	goto L695
L695:
	;
	v3595 = int32(8)
	v3600 = int32(0)
	goto L696
L696:
	;
	v3607 = v3600 << (uint(int32(2)) % 32)
	v3609 = *(*int32)(unsafe.Add(mBase, uint32(v3577+v3595+v3607)))
	v3611 = *(*int32)(unsafe.Add(mBase, uint32(v3576+v3595+v3607)))
	v3612 = v3609 & v3611
	v3614 = base.B2i32(v3612 != int32(0))
	if v3612 != 0 {
		v3623 = v3614
		goto L688
	} else {
		goto L698
	}
L697:
	;
	v3623 = v3614
	goto L688
L698:
	;
	v3616 = v3600 + int32(1)
	if v3616 != v3594 {
		v3600 = v3616
		goto L696
	} else {
		goto L699
	}
L699:
	;
	goto L697
L700:
	;
	goto L670
L701:
	;
	if base.B2i32(base.Ui32(v3627) < base.Ui32(int32(15)))&int32(base.Ui32(int32(_a_F_add_paths_to_joinrel_0))>>(uint(v3627)%32)) != 0 {
		v3637 = v3509
		goto L667
	} else {
		goto L702
	}
L702:
	;
	v3633 = F_create_material_path(m, l3, v3518)
	mBase = m.M
	v3634 = m.ExcPending
	if v3634 != 0 {
		goto L7
	} else {
		goto L703
	}
L703:
	;
	v3637 = v3633
	goto L667
L704:
	;
	goto L666
L705:
	;
	v3641 = *(*int32)(unsafe.Add(mBase, uint32(v3638)+4))
	if v3641 <= int32(0) {
		goto L704
	} else {
		goto L706
	}
L706:
	;
	v3666 = v3509
	goto L707
L707:
	;
	v3678 = *(*int32)(unsafe.Add(mBase, uint32(v3638)+12))
	v3682 = *(*int32)(unsafe.Add(mBase, uint32(v3678+v3666<<(uint(int32(2))%32))))
	v3683 = *(*int32)(unsafe.Add(mBase, uint32(v3682)+64))
	v3684 = F_build_join_pathkeys(m, l0, l1, v136, v3683)
	mBase = m.M
	v3685 = m.ExcPending
	if v3685 != 0 {
		goto L7
	} else {
		goto L709
	}
L708:
	;
	goto L704
L709:
	;
	v3686 = *(*int32)(unsafe.Add(mBase, uint32(l3)+64))
	if v3686 == int32(0) {
		goto L710
	} else {
		goto L711
	}
L710:
	;
	if v3637 != 0 {
		goto L722
	} else {
		goto L723
	}
L711:
	;
	v3689 = int32(0)
	v3690 = *(*int32)(unsafe.Add(mBase, uint32(v3686)+4))
	if v3690 <= v3689 {
		goto L710
	} else {
		goto L712
	}
L712:
	;
	v3698 = v3689
	goto L713
L713:
	;
	v3727 = *(*int32)(unsafe.Add(mBase, uint32(v3686)+12))
	v3731 = *(*int32)(unsafe.Add(mBase, uint32(v3727+v3698<<(uint(int32(2))%32))))
	v3732 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v3731)+21)))
	if v3732 == int32(0) {
		goto L715
	} else {
		goto L716
	}
L714:
	;
	goto L710
L715:
	;
	v3747 = v3698 + int32(1)
	v3748 = *(*int32)(unsafe.Add(mBase, uint32(v3686)+4))
	if v3747 < v3748 {
		v3698 = v3747
		goto L713
	} else {
		goto L721
	}
L716:
	;
	F_try_partial_nestloop_path(m, l0, l1, v3682, v3731, v3684, v136, int64(256), v3512)
	mBase = m.M
	v3737 = m.ExcPending
	if v3737 != 0 {
		goto L7
	} else {
		goto L717
	}
L717:
	;
	v3738 = F_get_memoize_path(m, l0, l3, l2, v3731, v3682, v136, v3512)
	mBase = m.M
	v3739 = m.ExcPending
	if v3739 != 0 {
		goto L7
	} else {
		goto L718
	}
L718:
	;
	if v3738 == int32(0) {
		goto L715
	} else {
		goto L719
	}
L719:
	;
	F_try_partial_nestloop_path(m, l0, l1, v3682, v3738, v3684, v136, int64(1024), v3512)
	mBase = m.M
	v3744 = m.ExcPending
	if v3744 != 0 {
		goto L7
	} else {
		goto L720
	}
L720:
	;
	goto L715
L721:
	;
	goto L714
L722:
	;
	F_try_partial_nestloop_path(m, l0, l1, v3682, v3637, v3684, v136, int64(512), v3512)
	mBase = m.M
	v3786 = m.ExcPending
	if v3786 != 0 {
		goto L7
	} else {
		goto L725
	}
L723:
	;
	goto L724
L724:
	;
	v3788 = v3666 + int32(1)
	v3789 = *(*int32)(unsafe.Add(mBase, uint32(v3638)+4))
	if v3788 < v3789 {
		v3666 = v3788
		goto L707
	} else {
		goto L726
	}
L725:
	;
	goto L724
L726:
	;
	goto L708
L727:
	;
	v3912 = *(*int32)(unsafe.Add(mBase, uint32(l2)+52))
	if v3912 == int32(0) {
		goto L750
	} else {
		goto L751
	}
L728:
	;
	v3859 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v3102)+21)))
	if v3859 != 0 {
		v3909 = v3102
		goto L727
	} else {
		goto L731
	}
L729:
	;
	goto L730
L730:
	;
	v3860 = *(*int32)(unsafe.Add(mBase, uint32(l3)+44))
	if v3860 != 0 {
		goto L734
	} else {
		goto L735
	}
L731:
	;
	goto L730
L732:
	;
	if v3904 == int32(0) {
		goto L285
	} else {
		goto L749
	}
L733:
	;
	goto L732
L734:
	;
	v3865 = *(*int32)(unsafe.Add(mBase, uint32(v3860)+4))
	if v3865 <= int32(0) {
		v3904 = int32(0)
		goto L733
	} else {
		goto L737
	}
L735:
	;
	goto L736
L736:
	;
	v3904 = int32(0)
	goto L733
L737:
	;
	v3868 = int32(0)
	if v3868 < v3865 {
		goto L738
	} else {
		goto L739
	}
L738:
	;
	v3871 = v3865
	goto L740
L739:
	;
	v3871 = v3868
	goto L740
L740:
	;
	v3872 = *(*int32)(unsafe.Add(mBase, uint32(v3860)+12))
	v3874 = int32(0)
	goto L741
L741:
	;
	v3882 = *(*int32)(unsafe.Add(mBase, uint32(v3872+v3874<<(uint(int32(2))%32))))
	v3883 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v3882)+21)))
	if v3883 == int32(1) {
		goto L743
	} else {
		goto L744
	}
L742:
	;
	goto L736
L743:
	;
	v3886 = *(*int32)(unsafe.Add(mBase, uint32(v3882)+16))
	if v3886 == int32(0) {
		v3904 = v3882
		goto L733
	} else {
		goto L746
	}
L744:
	;
	goto L745
L745:
	;
	v3894 = v3874 + int32(1)
	if v3894 != v3871 {
		v3874 = v3894
		goto L741
	} else {
		goto L748
	}
L746:
	;
	v3889 = *(*int32)(unsafe.Add(mBase, uint32(v3886)+4))
	if v3889 == int32(0) {
		v3904 = v3882
		goto L733
	} else {
		goto L747
	}
L747:
	;
	goto L745
L748:
	;
	goto L742
L749:
	;
	v3909 = v3904
	goto L727
L750:
	;
	goto L285
L751:
	;
	v3915 = *(*int32)(unsafe.Add(mBase, uint32(v3912)+4))
	if v3915 <= int32(0) {
		goto L750
	} else {
		goto L752
	}
L752:
	;
	v3924 = int32(0)
	goto L753
L753:
	;
	v3953 = *(*int32)(unsafe.Add(mBase, uint32(v3912)+12))
	v3957 = *(*int32)(unsafe.Add(mBase, uint32(v3953+v3924<<(uint(int32(2))%32))))
	v3959 = *(*int32)(unsafe.Add(mBase, uint32(v3957)+64))
	v3960 = F_build_join_pathkeys(m, l0, l1, v136, v3959)
	mBase = m.M
	v3961 = m.ExcPending
	if v3961 != 0 {
		goto L7
	} else {
		goto L755
	}
L754:
	;
	goto L750
L755:
	;
	F_generate_mergejoin_paths(m, l0, l1, l3, v3957, v136, v35+int32(-48), int32(0), v3909, v3960, int32(1))
	mBase = m.M
	v3964 = m.ExcPending
	if v3964 != 0 {
		goto L7
	} else {
		goto L756
	}
L756:
	;
	v3966 = v3924 + int32(1)
	v3967 = *(*int32)(unsafe.Add(mBase, uint32(v3912)+4))
	if v3966 < v3967 {
		v3924 = v3966
		goto L753
	} else {
		goto L757
	}
L757:
	;
	goto L754
L758:
	;
	v4039 = *(*int64)(unsafe.Add(mBase, uint32(v37)+56))
	if v4039&int64(2048) == int64(0) {
		goto L282
	} else {
		goto L761
	}
L759:
	;
	goto L760
L760:
	;
	v4044 = *(*int32)(unsafe.Add(mBase, uint32(v37)+16))
	if v4044 == int32(0) {
		goto L282
	} else {
		goto L762
	}
L761:
	;
	goto L760
L762:
	;
	v4047 = *(*int32)(unsafe.Add(mBase, uint32(v4044)+4))
	if v4047 <= int32(0) {
		goto L282
	} else {
		goto L763
	}
L763:
	;
	v4054 = int32(0)
	v4062 = v4054
	v4069 = v4054
	goto L764
L764:
	;
	v4090 = *(*int32)(unsafe.Add(mBase, uint32(v4044)+12))
	v4094 = *(*int32)(unsafe.Add(mBase, uint32(v4090+v4062<<(uint(int32(2))%32))))
	if int32(1)<<(uint(v136)%32)&int32(174) != 0 {
		goto L767
	} else {
		goto L768
	}
L765:
	;
	if v4399 == int32(0) {
		goto L282
	} else {
		goto L856
	}
L766:
	;
	v4402 = v4062 + int32(1)
	v4403 = *(*int32)(unsafe.Add(mBase, uint32(v4044)+4))
	if v4402 < v4403 {
		v4062 = v4402
		v4069 = v4399
		goto L764
	} else {
		goto L855
	}
L767:
	;
	v4095 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v4094)+8)))
	if v4095 != 0 {
		v4399 = v4069
		goto L766
	} else {
		goto L770
	}
L768:
	;
	goto L769
L769:
	;
	v4154 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v4094)+9)))
	if v4154 != int32(1) {
		v4399 = v4069
		goto L766
	} else {
		goto L786
	}
L770:
	;
	v4096 = *(*int32)(unsafe.Add(mBase, uint32(v4094)+32))
	v4097 = *(*int32)(unsafe.Add(mBase, uint32(l1)+8))
	v4098 = int32(0)
	if v4096 == v4098 {
		goto L772
	} else {
		goto L773
	}
L771:
	;
	if v4151 == int32(0) {
		v4399 = v4069
		goto L766
	} else {
		goto L785
	}
L772:
	;
	v4151 = int32(1)
	goto L771
L773:
	;
	goto L774
L774:
	;
	if v4097 == int32(0) {
		v4144 = v4098
		goto L775
	} else {
		goto L776
	}
L775:
	;
	v4151 = v4144
	goto L771
L776:
	;
	v4107 = *(*int32)(unsafe.Add(mBase, uint32(v4096)+4))
	v4108 = *(*int32)(unsafe.Add(mBase, uint32(v4097)+4))
	if v4108 < v4107 {
		v4144 = v4098
		goto L775
	} else {
		goto L777
	}
L777:
	;
	v4110 = int32(1)
	if v4107 <= v4110 {
		goto L778
	} else {
		goto L779
	}
L778:
	;
	v4113 = v4110
	goto L780
L779:
	;
	v4113 = v4107
	goto L780
L780:
	;
	v4114 = int32(8)
	v4119 = int32(0)
	goto L781
L781:
	;
	v4126 = v4119 << (uint(int32(2)) % 32)
	v4128 = *(*int32)(unsafe.Add(mBase, uint32(v4096+v4114+v4126)))
	v4130 = *(*int32)(unsafe.Add(mBase, uint32(v4097+v4114+v4126)))
	v4133 = v4128 & (v4130 ^ int32(-1))
	v4135 = base.B2i32(v4133 == int32(0))
	if v4133 != 0 {
		v4144 = v4135
		goto L775
	} else {
		goto L783
	}
L782:
	;
	v4144 = v4135
	goto L775
L783:
	;
	v4137 = v4119 + int32(1)
	if v4137 != v4113 {
		v4119 = v4137
		goto L781
	} else {
		goto L784
	}
L784:
	;
	goto L782
L785:
	;
	goto L769
L786:
	;
	v4157 = *(*int32)(unsafe.Add(mBase, uint32(v4094)+124))
	if v4157 == int32(0) {
		v4399 = v4069
		goto L766
	} else {
		goto L787
	}
L787:
	;
	v4160 = *(*int32)(unsafe.Add(mBase, uint32(l3)+8))
	v4161 = *(*int32)(unsafe.Add(mBase, uint32(v4094)+44))
	v4162 = *(*int32)(unsafe.Add(mBase, uint32(l2)+8))
	v4163 = int32(0)
	if v4161 == v4163 {
		goto L791
	} else {
		goto L792
	}
L788:
	;
	v4396 = F_lappend(m, v4069, v4094)
	mBase = m.M
	v4397 = m.ExcPending
	if v4397 != 0 {
		goto L7
	} else {
		goto L854
	}
L789:
	;
	v4394 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(v4094)+120)) = uint8(v4394)
	goto L788
L790:
	;
	if v4216 != 0 {
		goto L804
	} else {
		goto L805
	}
L791:
	;
	v4216 = int32(1)
	goto L790
L792:
	;
	goto L793
L793:
	;
	if v4162 == int32(0) {
		v4209 = v4163
		goto L794
	} else {
		goto L795
	}
L794:
	;
	v4216 = v4209
	goto L790
L795:
	;
	v4172 = *(*int32)(unsafe.Add(mBase, uint32(v4161)+4))
	v4173 = *(*int32)(unsafe.Add(mBase, uint32(v4162)+4))
	if v4173 < v4172 {
		v4209 = v4163
		goto L794
	} else {
		goto L796
	}
L796:
	;
	v4175 = int32(1)
	if v4172 <= v4175 {
		goto L797
	} else {
		goto L798
	}
L797:
	;
	v4178 = v4175
	goto L799
L798:
	;
	v4178 = v4172
	goto L799
L799:
	;
	v4179 = int32(8)
	v4184 = int32(0)
	goto L800
L800:
	;
	v4191 = v4184 << (uint(int32(2)) % 32)
	v4193 = *(*int32)(unsafe.Add(mBase, uint32(v4161+v4179+v4191)))
	v4195 = *(*int32)(unsafe.Add(mBase, uint32(v4162+v4179+v4191)))
	v4198 = v4193 & (v4195 ^ int32(-1))
	v4200 = base.B2i32(v4198 == int32(0))
	if v4198 != 0 {
		v4209 = v4200
		goto L794
	} else {
		goto L802
	}
L801:
	;
	v4209 = v4200
	goto L794
L802:
	;
	v4202 = v4184 + int32(1)
	if v4202 != v4178 {
		v4184 = v4202
		goto L800
	} else {
		goto L803
	}
L803:
	;
	goto L801
L804:
	;
	v4217 = *(*int32)(unsafe.Add(mBase, uint32(v4094)+48))
	v4218 = int32(0)
	if v4217 == v4218 {
		goto L808
	} else {
		goto L809
	}
L805:
	;
	goto L806
L806:
	;
	v4272 = *(*int32)(unsafe.Add(mBase, uint32(v4094)+44))
	v4273 = int32(0)
	if v4272 == v4273 {
		goto L823
	} else {
		goto L824
	}
L807:
	;
	if v4271 != 0 {
		goto L789
	} else {
		goto L821
	}
L808:
	;
	v4271 = int32(1)
	goto L807
L809:
	;
	goto L810
L810:
	;
	if v4160 == int32(0) {
		v4264 = v4218
		goto L811
	} else {
		goto L812
	}
L811:
	;
	v4271 = v4264
	goto L807
L812:
	;
	v4227 = *(*int32)(unsafe.Add(mBase, uint32(v4217)+4))
	v4228 = *(*int32)(unsafe.Add(mBase, uint32(v4160)+4))
	if v4228 < v4227 {
		v4264 = v4218
		goto L811
	} else {
		goto L813
	}
L813:
	;
	v4230 = int32(1)
	if v4227 <= v4230 {
		goto L814
	} else {
		goto L815
	}
L814:
	;
	v4233 = v4230
	goto L816
L815:
	;
	v4233 = v4227
	goto L816
L816:
	;
	v4234 = int32(8)
	v4239 = int32(0)
	goto L817
L817:
	;
	v4246 = v4239 << (uint(int32(2)) % 32)
	v4248 = *(*int32)(unsafe.Add(mBase, uint32(v4217+v4234+v4246)))
	v4250 = *(*int32)(unsafe.Add(mBase, uint32(v4160+v4234+v4246)))
	v4253 = v4248 & (v4250 ^ int32(-1))
	v4255 = base.B2i32(v4253 == int32(0))
	if v4253 != 0 {
		v4264 = v4255
		goto L811
	} else {
		goto L819
	}
L818:
	;
	v4264 = v4255
	goto L811
L819:
	;
	v4257 = v4239 + int32(1)
	if v4257 != v4233 {
		v4239 = v4257
		goto L817
	} else {
		goto L820
	}
L820:
	;
	goto L818
L821:
	;
	goto L806
L822:
	;
	if v4326 == int32(0) {
		v4399 = v4069
		goto L766
	} else {
		goto L836
	}
L823:
	;
	v4326 = int32(1)
	goto L822
L824:
	;
	goto L825
L825:
	;
	if v4160 == int32(0) {
		v4319 = v4273
		goto L826
	} else {
		goto L827
	}
L826:
	;
	v4326 = v4319
	goto L822
L827:
	;
	v4282 = *(*int32)(unsafe.Add(mBase, uint32(v4272)+4))
	v4283 = *(*int32)(unsafe.Add(mBase, uint32(v4160)+4))
	if v4283 < v4282 {
		v4319 = v4273
		goto L826
	} else {
		goto L828
	}
L828:
	;
	v4285 = int32(1)
	if v4282 <= v4285 {
		goto L829
	} else {
		goto L830
	}
L829:
	;
	v4288 = v4285
	goto L831
L830:
	;
	v4288 = v4282
	goto L831
L831:
	;
	v4289 = int32(8)
	v4294 = int32(0)
	goto L832
L832:
	;
	v4301 = v4294 << (uint(int32(2)) % 32)
	v4303 = *(*int32)(unsafe.Add(mBase, uint32(v4272+v4289+v4301)))
	v4305 = *(*int32)(unsafe.Add(mBase, uint32(v4160+v4289+v4301)))
	v4308 = v4303 & (v4305 ^ int32(-1))
	v4310 = base.B2i32(v4308 == int32(0))
	if v4308 != 0 {
		v4319 = v4310
		goto L826
	} else {
		goto L834
	}
L833:
	;
	v4319 = v4310
	goto L826
L834:
	;
	v4312 = v4294 + int32(1)
	if v4312 != v4288 {
		v4294 = v4312
		goto L832
	} else {
		goto L835
	}
L835:
	;
	goto L833
L836:
	;
	v4329 = *(*int32)(unsafe.Add(mBase, uint32(v4094)+48))
	v4330 = int32(0)
	if v4329 == v4330 {
		goto L838
	} else {
		goto L839
	}
L837:
	;
	if v4383 == int32(0) {
		v4399 = v4069
		goto L766
	} else {
		goto L851
	}
L838:
	;
	v4383 = int32(1)
	goto L837
L839:
	;
	goto L840
L840:
	;
	if v4162 == int32(0) {
		v4376 = v4330
		goto L841
	} else {
		goto L842
	}
L841:
	;
	v4383 = v4376
	goto L837
L842:
	;
	v4339 = *(*int32)(unsafe.Add(mBase, uint32(v4329)+4))
	v4340 = *(*int32)(unsafe.Add(mBase, uint32(v4162)+4))
	if v4340 < v4339 {
		v4376 = v4330
		goto L841
	} else {
		goto L843
	}
L843:
	;
	v4342 = int32(1)
	if v4339 <= v4342 {
		goto L844
	} else {
		goto L845
	}
L844:
	;
	v4345 = v4342
	goto L846
L845:
	;
	v4345 = v4339
	goto L846
L846:
	;
	v4346 = int32(8)
	v4351 = int32(0)
	goto L847
L847:
	;
	v4358 = v4351 << (uint(int32(2)) % 32)
	v4360 = *(*int32)(unsafe.Add(mBase, uint32(v4329+v4346+v4358)))
	v4362 = *(*int32)(unsafe.Add(mBase, uint32(v4162+v4346+v4358)))
	v4365 = v4360 & (v4362 ^ int32(-1))
	v4367 = base.B2i32(v4365 == int32(0))
	if v4365 != 0 {
		v4376 = v4367
		goto L841
	} else {
		goto L849
	}
L848:
	;
	v4376 = v4367
	goto L841
L849:
	;
	v4369 = v4351 + int32(1)
	if v4369 != v4345 {
		v4351 = v4369
		goto L847
	} else {
		goto L850
	}
L850:
	;
	goto L848
L851:
	;
	v4386 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v4094)+120)) = uint8(v4386)
	v4388 = *(*int32)(unsafe.Add(mBase, uint32(v4094)+4))
	v4389 = *(*int32)(unsafe.Add(mBase, uint32(v4388)+4))
	v4390 = F_get_commutator(m, v4389)
	mBase = m.M
	v4391 = m.ExcPending
	if v4391 != 0 {
		goto L7
	} else {
		goto L852
	}
L852:
	;
	if v4390 == int32(0) {
		v4399 = v4069
		goto L766
	} else {
		goto L853
	}
L853:
	;
	goto L788
L854:
	;
	v4399 = v4396
	goto L766
L855:
	;
	goto L765
L856:
	;
	v4407 = *(*int32)(unsafe.Add(mBase, uint32(l3)+60))
	v4408 = *(*int32)(unsafe.Add(mBase, uint32(l2)+56))
	v4409 = *(*int32)(unsafe.Add(mBase, uint32(l2)+60))
	v4410 = *(*int32)(unsafe.Add(mBase, uint32(v4409)+16))
	if v4410 == int32(0) {
		goto L857
	} else {
		goto L858
	}
L857:
	;
	v4513 = *(*int32)(unsafe.Add(mBase, uint32(v4407)+16))
	if v4513 == int32(0) {
		goto L888
	} else {
		goto L889
	}
L858:
	;
	v4413 = *(*int32)(unsafe.Add(mBase, uint32(v4410)+4))
	v4414 = *(*int32)(unsafe.Add(mBase, uint32(l3)+8))
	v4415 = int32(0)
	if base.B2i32(v4413 == v4415)|base.B2i32(v4414 == v4415) != 0 {
		v4460 = v4415
		goto L860
	} else {
		goto L861
	}
L859:
	;
	if v4460 != 0 {
		goto L282
	} else {
		goto L872
	}
L860:
	;
	goto L859
L861:
	;
	v4425 = *(*int32)(unsafe.Add(mBase, uint32(v4413)+4))
	v4426 = *(*int32)(unsafe.Add(mBase, uint32(v4414)+4))
	if v4425 < v4426 {
		goto L862
	} else {
		goto L863
	}
L862:
	;
	v4428 = v4425
	goto L864
L863:
	;
	v4428 = v4426
	goto L864
L864:
	;
	if v4428 <= int32(1) {
		goto L865
	} else {
		goto L866
	}
L865:
	;
	v4431 = int32(1)
	goto L867
L866:
	;
	v4431 = v4428
	goto L867
L867:
	;
	v4432 = int32(8)
	v4437 = int32(0)
	goto L868
L868:
	;
	v4444 = v4437 << (uint(int32(2)) % 32)
	v4446 = *(*int32)(unsafe.Add(mBase, uint32(v4414+v4432+v4444)))
	v4448 = *(*int32)(unsafe.Add(mBase, uint32(v4413+v4432+v4444)))
	v4449 = v4446 & v4448
	v4451 = base.B2i32(v4449 != int32(0))
	if v4449 != 0 {
		v4460 = v4451
		goto L860
	} else {
		goto L870
	}
L869:
	;
	v4460 = v4451
	goto L860
L870:
	;
	v4453 = v4437 + int32(1)
	if v4453 != v4431 {
		v4437 = v4453
		goto L868
	} else {
		goto L871
	}
L871:
	;
	goto L869
L872:
	;
	v4461 = *(*int32)(unsafe.Add(mBase, uint32(v4409)+16))
	if v4461 == int32(0) {
		goto L857
	} else {
		goto L873
	}
L873:
	;
	v4464 = *(*int32)(unsafe.Add(mBase, uint32(v4461)+4))
	v4465 = *(*int32)(unsafe.Add(mBase, uint32(l3)+252))
	v4466 = int32(0)
	if base.B2i32(v4464 == v4466)|base.B2i32(v4465 == v4466) != 0 {
		v4511 = v4466
		goto L875
	} else {
		goto L876
	}
L874:
	;
	if v4511 != 0 {
		goto L282
	} else {
		goto L887
	}
L875:
	;
	goto L874
L876:
	;
	v4476 = *(*int32)(unsafe.Add(mBase, uint32(v4464)+4))
	v4477 = *(*int32)(unsafe.Add(mBase, uint32(v4465)+4))
	if v4476 < v4477 {
		goto L877
	} else {
		goto L878
	}
L877:
	;
	v4479 = v4476
	goto L879
L878:
	;
	v4479 = v4477
	goto L879
L879:
	;
	if v4479 <= int32(1) {
		goto L880
	} else {
		goto L881
	}
L880:
	;
	v4482 = int32(1)
	goto L882
L881:
	;
	v4482 = v4479
	goto L882
L882:
	;
	v4483 = int32(8)
	v4488 = int32(0)
	goto L883
L883:
	;
	v4495 = v4488 << (uint(int32(2)) % 32)
	v4497 = *(*int32)(unsafe.Add(mBase, uint32(v4465+v4483+v4495)))
	v4499 = *(*int32)(unsafe.Add(mBase, uint32(v4464+v4483+v4495)))
	v4500 = v4497 & v4499
	v4502 = base.B2i32(v4500 != int32(0))
	if v4500 != 0 {
		v4511 = v4502
		goto L875
	} else {
		goto L885
	}
L884:
	;
	v4511 = v4502
	goto L875
L885:
	;
	v4504 = v4488 + int32(1)
	if v4504 != v4482 {
		v4488 = v4504
		goto L883
	} else {
		goto L886
	}
L886:
	;
	goto L884
L887:
	;
	goto L857
L888:
	;
	if v4408 != 0 {
		goto L919
	} else {
		goto L920
	}
L889:
	;
	v4516 = *(*int32)(unsafe.Add(mBase, uint32(v4513)+4))
	v4517 = *(*int32)(unsafe.Add(mBase, uint32(l2)+8))
	v4518 = int32(0)
	if base.B2i32(v4516 == v4518)|base.B2i32(v4517 == v4518) != 0 {
		v4563 = v4518
		goto L891
	} else {
		goto L892
	}
L890:
	;
	if v4563 != 0 {
		goto L282
	} else {
		goto L903
	}
L891:
	;
	goto L890
L892:
	;
	v4528 = *(*int32)(unsafe.Add(mBase, uint32(v4516)+4))
	v4529 = *(*int32)(unsafe.Add(mBase, uint32(v4517)+4))
	if v4528 < v4529 {
		goto L893
	} else {
		goto L894
	}
L893:
	;
	v4531 = v4528
	goto L895
L894:
	;
	v4531 = v4529
	goto L895
L895:
	;
	if v4531 <= int32(1) {
		goto L896
	} else {
		goto L897
	}
L896:
	;
	v4534 = int32(1)
	goto L898
L897:
	;
	v4534 = v4531
	goto L898
L898:
	;
	v4535 = int32(8)
	v4540 = int32(0)
	goto L899
L899:
	;
	v4547 = v4540 << (uint(int32(2)) % 32)
	v4549 = *(*int32)(unsafe.Add(mBase, uint32(v4517+v4535+v4547)))
	v4551 = *(*int32)(unsafe.Add(mBase, uint32(v4516+v4535+v4547)))
	v4552 = v4549 & v4551
	v4554 = base.B2i32(v4552 != int32(0))
	if v4552 != 0 {
		v4563 = v4554
		goto L891
	} else {
		goto L901
	}
L900:
	;
	v4563 = v4554
	goto L891
L901:
	;
	v4556 = v4540 + int32(1)
	if v4556 != v4534 {
		v4540 = v4556
		goto L899
	} else {
		goto L902
	}
L902:
	;
	goto L900
L903:
	;
	v4564 = *(*int32)(unsafe.Add(mBase, uint32(v4407)+16))
	if v4564 == int32(0) {
		goto L888
	} else {
		goto L904
	}
L904:
	;
	v4567 = *(*int32)(unsafe.Add(mBase, uint32(v4564)+4))
	v4568 = *(*int32)(unsafe.Add(mBase, uint32(l2)+252))
	v4569 = int32(0)
	if base.B2i32(v4567 == v4569)|base.B2i32(v4568 == v4569) != 0 {
		v4614 = v4569
		goto L906
	} else {
		goto L907
	}
L905:
	;
	if v4614 != 0 {
		goto L282
	} else {
		goto L918
	}
L906:
	;
	goto L905
L907:
	;
	v4579 = *(*int32)(unsafe.Add(mBase, uint32(v4567)+4))
	v4580 = *(*int32)(unsafe.Add(mBase, uint32(v4568)+4))
	if v4579 < v4580 {
		goto L908
	} else {
		goto L909
	}
L908:
	;
	v4582 = v4579
	goto L910
L909:
	;
	v4582 = v4580
	goto L910
L910:
	;
	if v4582 <= int32(1) {
		goto L911
	} else {
		goto L912
	}
L911:
	;
	v4585 = int32(1)
	goto L913
L912:
	;
	v4585 = v4582
	goto L913
L913:
	;
	v4586 = int32(8)
	v4591 = int32(0)
	goto L914
L914:
	;
	v4598 = v4591 << (uint(int32(2)) % 32)
	v4600 = *(*int32)(unsafe.Add(mBase, uint32(v4568+v4586+v4598)))
	v4602 = *(*int32)(unsafe.Add(mBase, uint32(v4567+v4586+v4598)))
	v4603 = v4600 & v4602
	v4605 = base.B2i32(v4603 != int32(0))
	if v4603 != 0 {
		v4614 = v4605
		goto L906
	} else {
		goto L916
	}
L915:
	;
	v4614 = v4605
	goto L906
L916:
	;
	v4607 = v4591 + int32(1)
	if v4607 != v4585 {
		v4591 = v4607
		goto L914
	} else {
		goto L917
	}
L917:
	;
	goto L915
L918:
	;
	goto L888
L919:
	;
	F_try_hashjoin_path(m, l0, l1, v4408, v4407, v4399, v136, v37+int32(16))
	mBase = m.M
	v4619 = m.ExcPending
	if v4619 != 0 {
		goto L7
	} else {
		goto L922
	}
L920:
	;
	goto L921
L921:
	;
	v4620 = *(*int32)(unsafe.Add(mBase, uint32(l2)+64))
	if v4620 == int32(0) {
		goto L283
	} else {
		goto L923
	}
L922:
	;
	goto L921
L923:
	;
	v4623 = *(*int32)(unsafe.Add(mBase, uint32(v4620)+4))
	if v4623 <= int32(0) {
		goto L283
	} else {
		goto L924
	}
L924:
	;
	v4649 = int32(0)
	goto L925
L925:
	;
	v4661 = *(*int32)(unsafe.Add(mBase, uint32(v4620)+12))
	v4665 = *(*int32)(unsafe.Add(mBase, uint32(v4661+v4649<<(uint(int32(2))%32))))
	v4666 = *(*int32)(unsafe.Add(mBase, uint32(v4665)+16))
	if v4666 == int32(0) {
		goto L928
	} else {
		goto L929
	}
L926:
	;
	goto L283
L927:
	;
	v4970 = v4649 + int32(1)
	v4971 = *(*int32)(unsafe.Add(mBase, uint32(v4620)+4))
	if v4970 < v4971 {
		v4649 = v4970
		goto L925
	} else {
		goto L1001
	}
L928:
	;
	v4769 = *(*int32)(unsafe.Add(mBase, uint32(l3)+64))
	if v4769 == int32(0) {
		goto L927
	} else {
		goto L959
	}
L929:
	;
	v4669 = *(*int32)(unsafe.Add(mBase, uint32(v4666)+4))
	v4670 = *(*int32)(unsafe.Add(mBase, uint32(l3)+8))
	v4671 = int32(0)
	if base.B2i32(v4669 == v4671)|base.B2i32(v4670 == v4671) != 0 {
		v4716 = v4671
		goto L931
	} else {
		goto L932
	}
L930:
	;
	if v4716 != 0 {
		goto L927
	} else {
		goto L943
	}
L931:
	;
	goto L930
L932:
	;
	v4681 = *(*int32)(unsafe.Add(mBase, uint32(v4669)+4))
	v4682 = *(*int32)(unsafe.Add(mBase, uint32(v4670)+4))
	if v4681 < v4682 {
		goto L933
	} else {
		goto L934
	}
L933:
	;
	v4684 = v4681
	goto L935
L934:
	;
	v4684 = v4682
	goto L935
L935:
	;
	if v4684 <= int32(1) {
		goto L936
	} else {
		goto L937
	}
L936:
	;
	v4687 = int32(1)
	goto L938
L937:
	;
	v4687 = v4684
	goto L938
L938:
	;
	v4688 = int32(8)
	v4693 = int32(0)
	goto L939
L939:
	;
	v4700 = v4693 << (uint(int32(2)) % 32)
	v4702 = *(*int32)(unsafe.Add(mBase, uint32(v4670+v4688+v4700)))
	v4704 = *(*int32)(unsafe.Add(mBase, uint32(v4669+v4688+v4700)))
	v4705 = v4702 & v4704
	v4707 = base.B2i32(v4705 != int32(0))
	if v4705 != 0 {
		v4716 = v4707
		goto L931
	} else {
		goto L941
	}
L940:
	;
	v4716 = v4707
	goto L931
L941:
	;
	v4709 = v4693 + int32(1)
	if v4709 != v4687 {
		v4693 = v4709
		goto L939
	} else {
		goto L942
	}
L942:
	;
	goto L940
L943:
	;
	v4717 = *(*int32)(unsafe.Add(mBase, uint32(v4665)+16))
	if v4717 == int32(0) {
		goto L928
	} else {
		goto L944
	}
L944:
	;
	v4720 = *(*int32)(unsafe.Add(mBase, uint32(v4717)+4))
	v4721 = *(*int32)(unsafe.Add(mBase, uint32(l3)+252))
	v4722 = int32(0)
	if base.B2i32(v4720 == v4722)|base.B2i32(v4721 == v4722) != 0 {
		v4767 = v4722
		goto L946
	} else {
		goto L947
	}
L945:
	;
	if v4767 != 0 {
		goto L927
	} else {
		goto L958
	}
L946:
	;
	goto L945
L947:
	;
	v4732 = *(*int32)(unsafe.Add(mBase, uint32(v4720)+4))
	v4733 = *(*int32)(unsafe.Add(mBase, uint32(v4721)+4))
	if v4732 < v4733 {
		goto L948
	} else {
		goto L949
	}
L948:
	;
	v4735 = v4732
	goto L950
L949:
	;
	v4735 = v4733
	goto L950
L950:
	;
	if v4735 <= int32(1) {
		goto L951
	} else {
		goto L952
	}
L951:
	;
	v4738 = int32(1)
	goto L953
L952:
	;
	v4738 = v4735
	goto L953
L953:
	;
	v4739 = int32(8)
	v4744 = int32(0)
	goto L954
L954:
	;
	v4751 = v4744 << (uint(int32(2)) % 32)
	v4753 = *(*int32)(unsafe.Add(mBase, uint32(v4721+v4739+v4751)))
	v4755 = *(*int32)(unsafe.Add(mBase, uint32(v4720+v4739+v4751)))
	v4756 = v4753 & v4755
	v4758 = base.B2i32(v4756 != int32(0))
	if v4756 != 0 {
		v4767 = v4758
		goto L946
	} else {
		goto L956
	}
L955:
	;
	v4767 = v4758
	goto L946
L956:
	;
	v4760 = v4744 + int32(1)
	if v4760 != v4738 {
		v4744 = v4760
		goto L954
	} else {
		goto L957
	}
L957:
	;
	goto L955
L958:
	;
	goto L928
L959:
	;
	v4772 = int32(0)
	v4773 = *(*int32)(unsafe.Add(mBase, uint32(v4769)+4))
	if v4773 <= v4772 {
		goto L927
	} else {
		goto L960
	}
L960:
	;
	v4783 = v4772
	goto L961
L961:
	;
	v4810 = *(*int32)(unsafe.Add(mBase, uint32(v4769)+12))
	v4814 = *(*int32)(unsafe.Add(mBase, uint32(v4810+v4783<<(uint(int32(2))%32))))
	v4815 = *(*int32)(unsafe.Add(mBase, uint32(v4814)+16))
	if v4815 == int32(0) {
		goto L965
	} else {
		goto L966
	}
L962:
	;
	goto L927
L963:
	;
	v4932 = v4783 + int32(1)
	v4933 = *(*int32)(unsafe.Add(mBase, uint32(v4769)+4))
	if v4932 < v4933 {
		v4783 = v4932
		goto L961
	} else {
		goto L1000
	}
L964:
	;
	F_try_hashjoin_path(m, l0, l1, v4665, v4814, v4399, v136, v37+int32(16))
	mBase = m.M
	v4929 = m.ExcPending
	if v4929 != 0 {
		goto L7
	} else {
		goto L999
	}
L965:
	;
	if v4665 != v4408 {
		goto L964
	} else {
		goto L997
	}
L966:
	;
	v4818 = *(*int32)(unsafe.Add(mBase, uint32(v4815)+4))
	v4819 = *(*int32)(unsafe.Add(mBase, uint32(l2)+8))
	v4820 = int32(0)
	if base.B2i32(v4818 == v4820)|base.B2i32(v4819 == v4820) != 0 {
		v4865 = v4820
		goto L968
	} else {
		goto L969
	}
L967:
	;
	if v4865 != 0 {
		goto L963
	} else {
		goto L980
	}
L968:
	;
	goto L967
L969:
	;
	v4830 = *(*int32)(unsafe.Add(mBase, uint32(v4818)+4))
	v4831 = *(*int32)(unsafe.Add(mBase, uint32(v4819)+4))
	if v4830 < v4831 {
		goto L970
	} else {
		goto L971
	}
L970:
	;
	v4833 = v4830
	goto L972
L971:
	;
	v4833 = v4831
	goto L972
L972:
	;
	if v4833 <= int32(1) {
		goto L973
	} else {
		goto L974
	}
L973:
	;
	v4836 = int32(1)
	goto L975
L974:
	;
	v4836 = v4833
	goto L975
L975:
	;
	v4837 = int32(8)
	v4842 = int32(0)
	goto L976
L976:
	;
	v4849 = v4842 << (uint(int32(2)) % 32)
	v4851 = *(*int32)(unsafe.Add(mBase, uint32(v4819+v4837+v4849)))
	v4853 = *(*int32)(unsafe.Add(mBase, uint32(v4818+v4837+v4849)))
	v4854 = v4851 & v4853
	v4856 = base.B2i32(v4854 != int32(0))
	if v4854 != 0 {
		v4865 = v4856
		goto L968
	} else {
		goto L978
	}
L977:
	;
	v4865 = v4856
	goto L968
L978:
	;
	v4858 = v4842 + int32(1)
	if v4858 != v4836 {
		v4842 = v4858
		goto L976
	} else {
		goto L979
	}
L979:
	;
	goto L977
L980:
	;
	v4866 = *(*int32)(unsafe.Add(mBase, uint32(v4814)+16))
	if v4866 == int32(0) {
		goto L965
	} else {
		goto L981
	}
L981:
	;
	v4869 = *(*int32)(unsafe.Add(mBase, uint32(v4866)+4))
	v4870 = *(*int32)(unsafe.Add(mBase, uint32(l2)+252))
	v4871 = int32(0)
	if base.B2i32(v4869 == v4871)|base.B2i32(v4870 == v4871) != 0 {
		v4916 = v4871
		goto L983
	} else {
		goto L984
	}
L982:
	;
	if v4916 != 0 {
		goto L963
	} else {
		goto L995
	}
L983:
	;
	goto L982
L984:
	;
	v4881 = *(*int32)(unsafe.Add(mBase, uint32(v4869)+4))
	v4882 = *(*int32)(unsafe.Add(mBase, uint32(v4870)+4))
	if v4881 < v4882 {
		goto L985
	} else {
		goto L986
	}
L985:
	;
	v4884 = v4881
	goto L987
L986:
	;
	v4884 = v4882
	goto L987
L987:
	;
	if v4884 <= int32(1) {
		goto L988
	} else {
		goto L989
	}
L988:
	;
	v4887 = int32(1)
	goto L990
L989:
	;
	v4887 = v4884
	goto L990
L990:
	;
	v4888 = int32(8)
	v4893 = int32(0)
	goto L991
L991:
	;
	v4900 = v4893 << (uint(int32(2)) % 32)
	v4902 = *(*int32)(unsafe.Add(mBase, uint32(v4870+v4888+v4900)))
	v4904 = *(*int32)(unsafe.Add(mBase, uint32(v4869+v4888+v4900)))
	v4905 = v4902 & v4904
	v4907 = base.B2i32(v4905 != int32(0))
	if v4905 != 0 {
		v4916 = v4907
		goto L983
	} else {
		goto L993
	}
L992:
	;
	v4916 = v4907
	goto L983
L993:
	;
	v4909 = v4893 + int32(1)
	if v4909 != v4887 {
		v4893 = v4909
		goto L991
	} else {
		goto L994
	}
L994:
	;
	goto L992
L995:
	;
	if base.B2i32(v4665 == v4408)&base.B2i32(v4407 == v4814) == int32(0) {
		goto L964
	} else {
		goto L996
	}
L996:
	;
	goto L963
L997:
	;
	if v4407 == v4814 {
		goto L963
	} else {
		goto L998
	}
L998:
	;
	goto L964
L999:
	;
	goto L963
L1000:
	;
	goto L962
L1001:
	;
	goto L926
L1002:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v37))) = v136
	F_errmsg_internal(m, int32(_a_F_add_paths_to_joinrel_1), v37)
	mBase = m.M
	v4980 = m.ExcPending
	if v4980 != 0 {
		goto L7
	} else {
		goto L1003
	}
L1003:
	;
	F_errfinish(m, int32(_a_F_add_paths_to_joinrel_2), int32(1878), int32(_a_F_add_paths_to_joinrel_3))
	mBase = m.M
	v4985 = m.ExcPending
	if v4985 != 0 {
		goto L7
	} else {
		goto L1004
	}
L1004:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L1005:
	;
	v5022 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l1)+26)))
	if v5022&int32(1) == int32(0) {
		goto L282
	} else {
		goto L1006
	}
L1006:
	;
	v5027 = *(*int32)(unsafe.Add(mBase, uint32(l2)+52))
	if v5027 == int32(0) {
		goto L282
	} else {
		goto L1007
	}
L1007:
	;
	v5030 = *(*int32)(unsafe.Add(mBase, uint32(l1)+72))
	if v5030 != 0 {
		goto L282
	} else {
		goto L1008
	}
L1008:
	;
	v5031 = *(*int32)(unsafe.Add(mBase, uint32(v5027)+12))
	v5032 = *(*int32)(unsafe.Add(mBase, uint32(v5031)))
	v5033 = *(*int32)(unsafe.Add(mBase, uint32(l3)+52))
	if v5033 == int32(0) {
		goto L1009
	} else {
		goto L1010
	}
L1009:
	;
	if int32(1)<<(uint(v136)%32)&int32(140) != 0 {
		goto L1013
	} else {
		goto L1014
	}
L1010:
	;
	v5037 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_add_paths_to_joinrel[1])))
	if v5037&int32(1) == int32(0) {
		goto L1009
	} else {
		goto L1011
	}
L1011:
	;
	v5042 = *(*int32)(unsafe.Add(mBase, uint32(v5033)+12))
	v5043 = *(*int32)(unsafe.Add(mBase, uint32(v5042)))
	F_try_partial_hashjoin_path(m, l0, l1, v5032, v5043, v4399, v136, v37+int32(16), int32(1))
	mBase = m.M
	v5048 = m.ExcPending
	if v5048 != 0 {
		goto L7
	} else {
		goto L1012
	}
L1012:
	;
	goto L1009
L1013:
	;
	v5056 = base.B2i32(base.Ui32(v136) <= base.Ui32(int32(7)))
	goto L1015
L1014:
	;
	v5056 = int32(0)
	goto L1015
L1015:
	;
	if v5056 != 0 {
		goto L282
	} else {
		goto L1016
	}
L1016:
	;
	v5057 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v4407)+21)))
	if v5057 == int32(0) {
		goto L1017
	} else {
		goto L1018
	}
L1017:
	;
	v5060 = *(*int32)(unsafe.Add(mBase, uint32(l3)+44))
	if v5060 != 0 {
		goto L1022
	} else {
		goto L1023
	}
L1018:
	;
	v5109 = v4407
	goto L1019
L1019:
	;
	F_try_partial_hashjoin_path(m, l0, l1, v5032, v5109, v4399, v136, v37+int32(16), int32(0))
	mBase = m.M
	v5114 = m.ExcPending
	if v5114 != 0 {
		goto L7
	} else {
		goto L1038
	}
L1020:
	;
	if v5104 == int32(0) {
		goto L282
	} else {
		goto L1037
	}
L1021:
	;
	goto L1020
L1022:
	;
	v5065 = *(*int32)(unsafe.Add(mBase, uint32(v5060)+4))
	if v5065 <= int32(0) {
		v5104 = int32(0)
		goto L1021
	} else {
		goto L1025
	}
L1023:
	;
	goto L1024
L1024:
	;
	v5104 = int32(0)
	goto L1021
L1025:
	;
	v5068 = int32(0)
	if v5068 < v5065 {
		goto L1026
	} else {
		goto L1027
	}
L1026:
	;
	v5071 = v5065
	goto L1028
L1027:
	;
	v5071 = v5068
	goto L1028
L1028:
	;
	v5072 = *(*int32)(unsafe.Add(mBase, uint32(v5060)+12))
	v5074 = int32(0)
	goto L1029
L1029:
	;
	v5082 = *(*int32)(unsafe.Add(mBase, uint32(v5072+v5074<<(uint(int32(2))%32))))
	v5083 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v5082)+21)))
	if v5083 == int32(1) {
		goto L1031
	} else {
		goto L1032
	}
L1030:
	;
	goto L1024
L1031:
	;
	v5086 = *(*int32)(unsafe.Add(mBase, uint32(v5082)+16))
	if v5086 == int32(0) {
		v5104 = v5082
		goto L1021
	} else {
		goto L1034
	}
L1032:
	;
	goto L1033
L1033:
	;
	v5094 = v5074 + int32(1)
	if v5094 != v5071 {
		v5074 = v5094
		goto L1029
	} else {
		goto L1036
	}
L1034:
	;
	v5089 = *(*int32)(unsafe.Add(mBase, uint32(v5086)+4))
	if v5089 == int32(0) {
		v5104 = v5082
		goto L1021
	} else {
		goto L1035
	}
L1035:
	;
	goto L1033
L1036:
	;
	goto L1030
L1037:
	;
	v5109 = v5104
	goto L1019
L1038:
	;
	goto L282
L1039:
	;
	v5166 = *(*int32)(unsafe.Add(mBase, _c_F_add_paths_to_joinrel[2]))
	if v5166 != 0 {
		goto L1044
	} else {
		goto L1045
	}
L1040:
	;
	v5154 = *(*int32)(unsafe.Add(mBase, uint32(l1)+176))
	if v5154 == int32(0) {
		goto L1039
	} else {
		goto L1041
	}
L1041:
	;
	v5157 = *(*int32)(unsafe.Add(mBase, uint32(v5154)+32))
	if v5157 == int32(0) {
		goto L1039
	} else {
		goto L1042
	}
L1042:
	;
	m.T0[v5157].(func(*base.Module, int32, int32, int32, int32, int32, int32))(m, l0, l1, l2, l3, l4, v37+int32(16))
	mBase = m.M
	v5163 = m.ExcPending
	if v5163 != 0 {
		goto L7
	} else {
		goto L1043
	}
L1043:
	;
	goto L1039
L1044:
	;
	m.T0[v5166].(func(*base.Module, int32, int32, int32, int32, int32, int32))(m, l0, l1, l2, l3, l4, v37+int32(16))
	mBase = m.M
	v5170 = m.ExcPending
	if v5170 != 0 {
		goto L7
	} else {
		goto L1047
	}
L1045:
	;
	goto L1046
L1046:
	;
	m.G0 = v37 - int32(-64)
	return
L1047:
	;
	goto L1046
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
	var v14 int32
	_ = v14
	var v15 int32
	_ = v15
	var v16 int32
	_ = v16
	var v30 int32
	_ = v30
	var v31 int32
	_ = v31
	var v39 int32
	_ = v39
	var v41 int32
	_ = v41
	var v43 int32
	_ = v43
	var v44 int32
	_ = v44
	var v45 int32
	_ = v45
	var v47 int32
	_ = v47
	var v49 int32
	_ = v49
	var v50 int32
	_ = v50
	var v54 int32
	_ = v54
	var v57 int32
	_ = v57
	var v58 int32
	_ = v58
	var v65 int32
	_ = v65
	var v68 int32
	_ = v68
	var v69 int32
	_ = v69
	var v71 int32
	_ = v71
	var v76 int64
	_ = v76
	var v78 int64
	_ = v78
	var v79 int64
	_ = v79
	var v80 int32
	_ = v80
	var v83 int32
	_ = v83
	var v88 int32
	_ = v88
	var v89 int32
	_ = v89
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
	var v105 int32
	_ = v105
	var v111 int32
	_ = v111
	v6 = int32(0)
	v14 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l2)+1)))
	v15 = *(*int32)(unsafe.Add(mBase, uint32(l1)+8))
	v16 = *(*int32)(unsafe.Add(mBase, uint32(v15)))
	if v6 < v16 {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	v30 = v6
	v31 = v14
	goto L4
L2:
	;
	v105 = v14
	goto L3
L3:
	;
	v111 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(l2)+1)) = uint8(v111)
	return v105 & int32(1)
L4:
	;
	v39 = l2 + int32(24) + v30*int32(24)
	v41 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l2)+1)))
	if v41 != 0 {
		v45 = int32(0)
		goto L6
	} else {
		goto L7
	}
L5:
	;
	v105 = v92
	goto L3
L6:
	;
	v47 = v30 << (uint(int32(2)) % 32)
	v49 = *(*int32)(unsafe.Add(mBase, uint32(l1+int32(20)+v47)))
	v50 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v49)+2)))
	if v50 != int32(1) {
		goto L10
	} else {
		goto L11
	}
L7:
	;
	v43 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v39)+2)))
	if v43 != 0 {
		v45 = int32(1)
		goto L6
	} else {
		goto L8
	}
L8:
	;
	v44 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v39)+3)))
	v45 = v44
	goto L6
L9:
	;
	v94 = v30 + int32(1)
	v95 = *(*int32)(unsafe.Add(mBase, uint32(l1)+8))
	v96 = *(*int32)(unsafe.Add(mBase, uint32(v95)))
	if v94 < v96 {
		v30 = v94
		v31 = v92
		goto L4
	} else {
		goto L20
	}
L10:
	;
	v65 = F_index_getprocinfo(m, l0, base.I32_extend16_s(v30+int32(1)), int32(2))
	mBase = m.M
	v68 = m.ExcPending
	if v68 != 0 {
		goto L14
	} else {
		goto L15
	}
L11:
	;
	v54 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l4+v30))))
	if v54 != int32(1) {
		goto L10
	} else {
		goto L12
	}
L12:
	;
	v57 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v39)+2)))
	if v57 != 0 {
		v92 = v31
		goto L9
	} else {
		goto L13
	}
L13:
	;
	v58 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(v39)+2)) = uint8(v58)
	v92 = v58
	goto L9
L14:
	;
	return int32(0)
L15:
	;
	v69 = *(*int32)(unsafe.Add(mBase, uint32(l0)+248))
	v71 = *(*int32)(unsafe.Add(mBase, uint32(v69+v47)))
	v76 = *(*int64)(unsafe.Add(mBase, uint32(l3+v30<<(uint(int32(3))%32))))
	v78 = int64(*(*uint8)(unsafe.Add(mBase, uint32(l4+v30))))
	v79 = F_FunctionCall4Coll(m, v65, v71, base.I64_extend_i32_u(l1), base.I64_extend_i32_u(v39), v76, v78)
	mBase = m.M
	v80 = m.ExcPending
	if v80 != 0 {
		goto L14
	} else {
		goto L16
	}
L16:
	;
	v83 = v31 | base.B2i32(v79 != int64(0))
	if v45&int32(1) == int32(0) {
		v92 = v83
		goto L9
	} else {
		goto L17
	}
L17:
	;
	v88 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v39)+2)))
	if v88 != 0 {
		v92 = v83
		goto L9
	} else {
		goto L18
	}
L18:
	;
	v89 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v39)+3)))
	if v89 != 0 {
		v92 = v83
		goto L9
	} else {
		goto L19
	}
L19:
	;
	v90 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(v39)+2)) = uint8(v90)
	v92 = v83
	goto L9
L20:
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
		*(*int32)(unsafe.Add(mBase, uint32(v9)+40)) = int32(533)
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
			v37 = *(*int32)(unsafe.Add(mBase, uint32(v36)+24))
			v38 = F_strlen(m, v9)
			mBase = m.M
			v39 = v37 + v38
			v40 = *(*int32)(unsafe.Add(mBase, uint32(v36)+28))
			if v40 <= v39 {
				v42 = int32(1)
				v45 = v39 + v42
				if v45&v39 != 0 {
					v50 = v42 << (uint(int32(32)-base.I32_clz(v45)) % 32)
				} else {
					v50 = v45
				}
				*(*int32)(unsafe.Add(mBase, uint32(v36)+28)) = v50
				v52 = *(*int32)(unsafe.Add(mBase, uint32(l1)))
				v53 = *(*int32)(unsafe.Add(mBase, uint32(v52)+20))
				v54 = *(*int32)(unsafe.Add(mBase, uint32(v52)+28))
				v55 = F_repalloc(m, v53, v54)
				mBase = m.M
				v56 = m.ExcPending
				if v56 != 0 {
					return
				} else {
					v57 = *(*int32)(unsafe.Add(mBase, uint32(l1)))
					*(*int32)(unsafe.Add(mBase, uint32(v57)+20)) = v55
					v59 = *(*int32)(unsafe.Add(mBase, uint32(l1)))
					v60 = *(*int32)(unsafe.Add(mBase, uint32(v59)+24))
					v61 = v59
					v62 = v60
					if v38 != 0 {
						v63 = *(*int32)(unsafe.Add(mBase, uint32(v61)+20))
						base.MemoryCopy(m, v63+v62, v9, v38)
					} else {
					}
					v66 = *(*int32)(unsafe.Add(mBase, uint32(l1)))
					v67 = *(*int32)(unsafe.Add(mBase, uint32(v66)+24))
					*(*int32)(unsafe.Add(mBase, uint32(v66)+24)) = v67 + v38
					m.G0 = v9 + int32(48)
					return
				}
			} else {
				v61 = v36
				v62 = v37
				if v38 != 0 {
					v63 = *(*int32)(unsafe.Add(mBase, uint32(v61)+20))
					base.MemoryCopy(m, v63+v62, v9, v38)
				} else {
				}
				v66 = *(*int32)(unsafe.Add(mBase, uint32(l1)))
				v67 = *(*int32)(unsafe.Add(mBase, uint32(v66)+24))
				*(*int32)(unsafe.Add(mBase, uint32(v66)+24)) = v67 + v38
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
	v11 = *(*int32)(unsafe.Add(mBase, uint32(l2)+244))
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
					F_errfinish(m, int32(_a_F_adjust_appendrel_attrs_multilevel_1), int32(612), int32(_a_F_adjust_appendrel_attrs_multilevel_2))
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
	v59 = *(*int32)(unsafe.Add(mBase, uint32(l2)+244))
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
	F_errfinish(m, int32(_a_F_adjust_child_relids_multilevel_1), int32(686), int32(_a_F_adjust_child_relids_multilevel_2))
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
	var v6 int32
	_ = v6
	var v7 int32
	_ = v7
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
	var v33 int32
	_ = v33
	var v35 int32
	_ = v35
	var v44 int32
	_ = v44
	var v47 int32
	_ = v47
	var v48 int32
	_ = v48
	var v49 int32
	_ = v49
	var v50 int32
	_ = v50
	var v51 int32
	_ = v51
	var v53 int32
	_ = v53
	var v54 int32
	_ = v54
	var v55 int32
	_ = v55
	var v58 int32
	_ = v58
	var v63 int32
	_ = v63
	var v68 int32
	_ = v68
	var v71 int32
	_ = v71
	var v72 int32
	_ = v72
	var v88 int32
	_ = v88
	var v91 int32
	_ = v91
	var v93 int32
	_ = v93
	var v96 int32
	_ = v96
	var v97 int32
	_ = v97
	var v101 int32
	_ = v101
	var v107 int32
	_ = v107
	var v111 int32
	_ = v111
	v6 = *(*int32)(unsafe.Add(mBase, _c_F_allcases[0]))
	v7 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v6)+2)))
	if v7 == int32(1) {
		if base.Ui32(l1) <= base.Ui32(int32(127)) {
			if base.Ui32((l1-int32(65))&int32(255)) < base.Ui32(int32(26)) {
				v33 = l1 | int32(32)
			} else {
				v33 = l1
			}
			v35 = v33
			if base.Ui32((l1-int32(97))&int32(255)) < base.Ui32(int32(26)) {
				v44 = l1 + int32(224)
			} else {
				v44 = l1
			}
			v51 = v44 & int32(255)
			v53 = v35
		} else {
			v51 = l1
			v53 = l1
		}
		v54 = *(*int32)(unsafe.Add(mBase, uint32(l0)+120))
		if v54 != 0 {
			v55 = *(*int32)(unsafe.Add(mBase, uint32(v54)+4))
			if v55 < int32(2) {
				F_pfree(m, v54)
				mBase = m.M
				v68 = m.ExcPending
				if v68 != 0 {
					return int32(0)
				} else {
					v71 = F_palloc_extended(m, int32(36), int32(2))
					mBase = m.M
					v72 = m.ExcPending
					if v72 != 0 {
						return int32(0)
					} else {
						if v71 != 0 {
							*(*int64)(unsafe.Add(mBase, uint32(v71))) = int64(8589934592)
							*(*int32)(unsafe.Add(mBase, uint32(v71)+24)) = int32(-1)
							*(*int64)(unsafe.Add(mBase, uint32(v71)+12)) = int64(0)
							*(*int32)(unsafe.Add(mBase, uint32(v71)+20)) = v71 + int32(36)
							*(*int32)(unsafe.Add(mBase, uint32(v71)+8)) = v71 + int32(28)
							*(*int32)(unsafe.Add(mBase, uint32(l0)+120)) = v71
							v96 = v71
						} else {
							*(*int32)(unsafe.Add(mBase, uint32(l0)+24)) = int32(101)
							v88 = int32(0)
							*(*int32)(unsafe.Add(mBase, uint32(l0)+120)) = v88
							v91 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
							if v91 != 0 {
								v93 = v91
							} else {
								v93 = int32(12)
							}
							*(*int32)(unsafe.Add(mBase, uint32(l0)+12)) = v93
							v96 = v88
						}
						v97 = *(*int32)(unsafe.Add(mBase, uint32(v96)))
						*(*int32)(unsafe.Add(mBase, uint32(v96))) = v97 + int32(1)
						v101 = *(*int32)(unsafe.Add(mBase, uint32(v96)+8))
						*(*int32)(unsafe.Add(mBase, uint32(v101+v97<<(uint(int32(2))%32)))) = v53
						if v51 != v53 {
							v107 = *(*int32)(unsafe.Add(mBase, uint32(v96)))
							*(*int32)(unsafe.Add(mBase, uint32(v96))) = v107 + int32(1)
							v111 = *(*int32)(unsafe.Add(mBase, uint32(v96)+8))
							*(*int32)(unsafe.Add(mBase, uint32(v111+v107<<(uint(int32(2))%32)))) = v51
						} else {
						}
						return v96
					}
				}
			} else {
				v58 = *(*int32)(unsafe.Add(mBase, uint32(v54)+16))
				if v58 < int32(0) {
					F_pfree(m, v54)
					mBase = m.M
					v68 = m.ExcPending
					if v68 != 0 {
						return int32(0)
					} else {
						v71 = F_palloc_extended(m, int32(36), int32(2))
						mBase = m.M
						v72 = m.ExcPending
						if v72 != 0 {
							return int32(0)
						} else {
							if v71 != 0 {
								*(*int64)(unsafe.Add(mBase, uint32(v71))) = int64(8589934592)
								*(*int32)(unsafe.Add(mBase, uint32(v71)+24)) = int32(-1)
								*(*int64)(unsafe.Add(mBase, uint32(v71)+12)) = int64(0)
								*(*int32)(unsafe.Add(mBase, uint32(v71)+20)) = v71 + int32(36)
								*(*int32)(unsafe.Add(mBase, uint32(v71)+8)) = v71 + int32(28)
								*(*int32)(unsafe.Add(mBase, uint32(l0)+120)) = v71
								v96 = v71
							} else {
								*(*int32)(unsafe.Add(mBase, uint32(l0)+24)) = int32(101)
								v88 = int32(0)
								*(*int32)(unsafe.Add(mBase, uint32(l0)+120)) = v88
								v91 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
								if v91 != 0 {
									v93 = v91
								} else {
									v93 = int32(12)
								}
								*(*int32)(unsafe.Add(mBase, uint32(l0)+12)) = v93
								v96 = v88
							}
							v97 = *(*int32)(unsafe.Add(mBase, uint32(v96)))
							*(*int32)(unsafe.Add(mBase, uint32(v96))) = v97 + int32(1)
							v101 = *(*int32)(unsafe.Add(mBase, uint32(v96)+8))
							*(*int32)(unsafe.Add(mBase, uint32(v101+v97<<(uint(int32(2))%32)))) = v53
							if v51 != v53 {
								v107 = *(*int32)(unsafe.Add(mBase, uint32(v96)))
								*(*int32)(unsafe.Add(mBase, uint32(v96))) = v107 + int32(1)
								v111 = *(*int32)(unsafe.Add(mBase, uint32(v96)+8))
								*(*int32)(unsafe.Add(mBase, uint32(v111+v107<<(uint(int32(2))%32)))) = v51
							} else {
							}
							return v96
						}
					}
				} else {
					*(*int32)(unsafe.Add(mBase, uint32(v54)+24)) = int32(-1)
					v63 = int32(0)
					*(*int32)(unsafe.Add(mBase, uint32(v54)+12)) = v63
					*(*int32)(unsafe.Add(mBase, uint32(v54))) = v63
					v96 = v54
					v97 = *(*int32)(unsafe.Add(mBase, uint32(v96)))
					*(*int32)(unsafe.Add(mBase, uint32(v96))) = v97 + int32(1)
					v101 = *(*int32)(unsafe.Add(mBase, uint32(v96)+8))
					*(*int32)(unsafe.Add(mBase, uint32(v101+v97<<(uint(int32(2))%32)))) = v53
					if v51 != v53 {
						v107 = *(*int32)(unsafe.Add(mBase, uint32(v96)))
						*(*int32)(unsafe.Add(mBase, uint32(v96))) = v107 + int32(1)
						v111 = *(*int32)(unsafe.Add(mBase, uint32(v96)+8))
						*(*int32)(unsafe.Add(mBase, uint32(v111+v107<<(uint(int32(2))%32)))) = v51
					} else {
					}
					return v96
				}
			}
		} else {
			v71 = F_palloc_extended(m, int32(36), int32(2))
			mBase = m.M
			v72 = m.ExcPending
			if v72 != 0 {
				return int32(0)
			} else {
				if v71 != 0 {
					*(*int64)(unsafe.Add(mBase, uint32(v71))) = int64(8589934592)
					*(*int32)(unsafe.Add(mBase, uint32(v71)+24)) = int32(-1)
					*(*int64)(unsafe.Add(mBase, uint32(v71)+12)) = int64(0)
					*(*int32)(unsafe.Add(mBase, uint32(v71)+20)) = v71 + int32(36)
					*(*int32)(unsafe.Add(mBase, uint32(v71)+8)) = v71 + int32(28)
					*(*int32)(unsafe.Add(mBase, uint32(l0)+120)) = v71
					v96 = v71
				} else {
					*(*int32)(unsafe.Add(mBase, uint32(l0)+24)) = int32(101)
					v88 = int32(0)
					*(*int32)(unsafe.Add(mBase, uint32(l0)+120)) = v88
					v91 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
					if v91 != 0 {
						v93 = v91
					} else {
						v93 = int32(12)
					}
					*(*int32)(unsafe.Add(mBase, uint32(l0)+12)) = v93
					v96 = v88
				}
				v97 = *(*int32)(unsafe.Add(mBase, uint32(v96)))
				*(*int32)(unsafe.Add(mBase, uint32(v96))) = v97 + int32(1)
				v101 = *(*int32)(unsafe.Add(mBase, uint32(v96)+8))
				*(*int32)(unsafe.Add(mBase, uint32(v101+v97<<(uint(int32(2))%32)))) = v53
				if v51 != v53 {
					v107 = *(*int32)(unsafe.Add(mBase, uint32(v96)))
					*(*int32)(unsafe.Add(mBase, uint32(v96))) = v107 + int32(1)
					v111 = *(*int32)(unsafe.Add(mBase, uint32(v96)+8))
					*(*int32)(unsafe.Add(mBase, uint32(v111+v107<<(uint(int32(2))%32)))) = v51
				} else {
				}
				return v96
			}
		}
	} else {
		v12 = *(*int32)(unsafe.Add(mBase, uint32(v6)+8))
		v13 = *(*int32)(unsafe.Add(mBase, uint32(v12)+64))
		v14 = m.T0[v13].(func(*base.Module, int32, int32) int32)(m, l1, v6)
		mBase = m.M
		v17 = m.ExcPending
		if v17 != 0 {
			return int32(0)
		} else {
			v19 = *(*int32)(unsafe.Add(mBase, _c_F_allcases[0]))
			v20 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v19)+2)))
			if v20 != int32(1) {
				v47 = *(*int32)(unsafe.Add(mBase, uint32(v19)+8))
				v48 = *(*int32)(unsafe.Add(mBase, uint32(v47)+60))
				v49 = m.T0[v48].(func(*base.Module, int32, int32) int32)(m, l1, v19)
				mBase = m.M
				v50 = m.ExcPending
				if v50 != 0 {
					return int32(0)
				} else {
					v51 = v49
					v53 = v14
					v54 = *(*int32)(unsafe.Add(mBase, uint32(l0)+120))
					if v54 != 0 {
						v55 = *(*int32)(unsafe.Add(mBase, uint32(v54)+4))
						if v55 < int32(2) {
							F_pfree(m, v54)
							mBase = m.M
							v68 = m.ExcPending
							if v68 != 0 {
								return int32(0)
							} else {
								v71 = F_palloc_extended(m, int32(36), int32(2))
								mBase = m.M
								v72 = m.ExcPending
								if v72 != 0 {
									return int32(0)
								} else {
									if v71 != 0 {
										*(*int64)(unsafe.Add(mBase, uint32(v71))) = int64(8589934592)
										*(*int32)(unsafe.Add(mBase, uint32(v71)+24)) = int32(-1)
										*(*int64)(unsafe.Add(mBase, uint32(v71)+12)) = int64(0)
										*(*int32)(unsafe.Add(mBase, uint32(v71)+20)) = v71 + int32(36)
										*(*int32)(unsafe.Add(mBase, uint32(v71)+8)) = v71 + int32(28)
										*(*int32)(unsafe.Add(mBase, uint32(l0)+120)) = v71
										v96 = v71
									} else {
										*(*int32)(unsafe.Add(mBase, uint32(l0)+24)) = int32(101)
										v88 = int32(0)
										*(*int32)(unsafe.Add(mBase, uint32(l0)+120)) = v88
										v91 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
										if v91 != 0 {
											v93 = v91
										} else {
											v93 = int32(12)
										}
										*(*int32)(unsafe.Add(mBase, uint32(l0)+12)) = v93
										v96 = v88
									}
									v97 = *(*int32)(unsafe.Add(mBase, uint32(v96)))
									*(*int32)(unsafe.Add(mBase, uint32(v96))) = v97 + int32(1)
									v101 = *(*int32)(unsafe.Add(mBase, uint32(v96)+8))
									*(*int32)(unsafe.Add(mBase, uint32(v101+v97<<(uint(int32(2))%32)))) = v53
									if v51 != v53 {
										v107 = *(*int32)(unsafe.Add(mBase, uint32(v96)))
										*(*int32)(unsafe.Add(mBase, uint32(v96))) = v107 + int32(1)
										v111 = *(*int32)(unsafe.Add(mBase, uint32(v96)+8))
										*(*int32)(unsafe.Add(mBase, uint32(v111+v107<<(uint(int32(2))%32)))) = v51
									} else {
									}
									return v96
								}
							}
						} else {
							v58 = *(*int32)(unsafe.Add(mBase, uint32(v54)+16))
							if v58 < int32(0) {
								F_pfree(m, v54)
								mBase = m.M
								v68 = m.ExcPending
								if v68 != 0 {
									return int32(0)
								} else {
									v71 = F_palloc_extended(m, int32(36), int32(2))
									mBase = m.M
									v72 = m.ExcPending
									if v72 != 0 {
										return int32(0)
									} else {
										if v71 != 0 {
											*(*int64)(unsafe.Add(mBase, uint32(v71))) = int64(8589934592)
											*(*int32)(unsafe.Add(mBase, uint32(v71)+24)) = int32(-1)
											*(*int64)(unsafe.Add(mBase, uint32(v71)+12)) = int64(0)
											*(*int32)(unsafe.Add(mBase, uint32(v71)+20)) = v71 + int32(36)
											*(*int32)(unsafe.Add(mBase, uint32(v71)+8)) = v71 + int32(28)
											*(*int32)(unsafe.Add(mBase, uint32(l0)+120)) = v71
											v96 = v71
										} else {
											*(*int32)(unsafe.Add(mBase, uint32(l0)+24)) = int32(101)
											v88 = int32(0)
											*(*int32)(unsafe.Add(mBase, uint32(l0)+120)) = v88
											v91 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
											if v91 != 0 {
												v93 = v91
											} else {
												v93 = int32(12)
											}
											*(*int32)(unsafe.Add(mBase, uint32(l0)+12)) = v93
											v96 = v88
										}
										v97 = *(*int32)(unsafe.Add(mBase, uint32(v96)))
										*(*int32)(unsafe.Add(mBase, uint32(v96))) = v97 + int32(1)
										v101 = *(*int32)(unsafe.Add(mBase, uint32(v96)+8))
										*(*int32)(unsafe.Add(mBase, uint32(v101+v97<<(uint(int32(2))%32)))) = v53
										if v51 != v53 {
											v107 = *(*int32)(unsafe.Add(mBase, uint32(v96)))
											*(*int32)(unsafe.Add(mBase, uint32(v96))) = v107 + int32(1)
											v111 = *(*int32)(unsafe.Add(mBase, uint32(v96)+8))
											*(*int32)(unsafe.Add(mBase, uint32(v111+v107<<(uint(int32(2))%32)))) = v51
										} else {
										}
										return v96
									}
								}
							} else {
								*(*int32)(unsafe.Add(mBase, uint32(v54)+24)) = int32(-1)
								v63 = int32(0)
								*(*int32)(unsafe.Add(mBase, uint32(v54)+12)) = v63
								*(*int32)(unsafe.Add(mBase, uint32(v54))) = v63
								v96 = v54
								v97 = *(*int32)(unsafe.Add(mBase, uint32(v96)))
								*(*int32)(unsafe.Add(mBase, uint32(v96))) = v97 + int32(1)
								v101 = *(*int32)(unsafe.Add(mBase, uint32(v96)+8))
								*(*int32)(unsafe.Add(mBase, uint32(v101+v97<<(uint(int32(2))%32)))) = v53
								if v51 != v53 {
									v107 = *(*int32)(unsafe.Add(mBase, uint32(v96)))
									*(*int32)(unsafe.Add(mBase, uint32(v96))) = v107 + int32(1)
									v111 = *(*int32)(unsafe.Add(mBase, uint32(v96)+8))
									*(*int32)(unsafe.Add(mBase, uint32(v111+v107<<(uint(int32(2))%32)))) = v51
								} else {
								}
								return v96
							}
						}
					} else {
						v71 = F_palloc_extended(m, int32(36), int32(2))
						mBase = m.M
						v72 = m.ExcPending
						if v72 != 0 {
							return int32(0)
						} else {
							if v71 != 0 {
								*(*int64)(unsafe.Add(mBase, uint32(v71))) = int64(8589934592)
								*(*int32)(unsafe.Add(mBase, uint32(v71)+24)) = int32(-1)
								*(*int64)(unsafe.Add(mBase, uint32(v71)+12)) = int64(0)
								*(*int32)(unsafe.Add(mBase, uint32(v71)+20)) = v71 + int32(36)
								*(*int32)(unsafe.Add(mBase, uint32(v71)+8)) = v71 + int32(28)
								*(*int32)(unsafe.Add(mBase, uint32(l0)+120)) = v71
								v96 = v71
							} else {
								*(*int32)(unsafe.Add(mBase, uint32(l0)+24)) = int32(101)
								v88 = int32(0)
								*(*int32)(unsafe.Add(mBase, uint32(l0)+120)) = v88
								v91 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
								if v91 != 0 {
									v93 = v91
								} else {
									v93 = int32(12)
								}
								*(*int32)(unsafe.Add(mBase, uint32(l0)+12)) = v93
								v96 = v88
							}
							v97 = *(*int32)(unsafe.Add(mBase, uint32(v96)))
							*(*int32)(unsafe.Add(mBase, uint32(v96))) = v97 + int32(1)
							v101 = *(*int32)(unsafe.Add(mBase, uint32(v96)+8))
							*(*int32)(unsafe.Add(mBase, uint32(v101+v97<<(uint(int32(2))%32)))) = v53
							if v51 != v53 {
								v107 = *(*int32)(unsafe.Add(mBase, uint32(v96)))
								*(*int32)(unsafe.Add(mBase, uint32(v96))) = v107 + int32(1)
								v111 = *(*int32)(unsafe.Add(mBase, uint32(v96)+8))
								*(*int32)(unsafe.Add(mBase, uint32(v111+v107<<(uint(int32(2))%32)))) = v51
							} else {
							}
							return v96
						}
					}
				}
			} else {
				if base.Ui32(l1) <= base.Ui32(int32(127)) {
					v35 = v14
					if base.Ui32((l1-int32(97))&int32(255)) < base.Ui32(int32(26)) {
						v44 = l1 + int32(224)
					} else {
						v44 = l1
					}
					v51 = v44 & int32(255)
					v53 = v35
				} else {
					v51 = l1
					v53 = v14
				}
				v54 = *(*int32)(unsafe.Add(mBase, uint32(l0)+120))
				if v54 != 0 {
					v55 = *(*int32)(unsafe.Add(mBase, uint32(v54)+4))
					if v55 < int32(2) {
						F_pfree(m, v54)
						mBase = m.M
						v68 = m.ExcPending
						if v68 != 0 {
							return int32(0)
						} else {
							v71 = F_palloc_extended(m, int32(36), int32(2))
							mBase = m.M
							v72 = m.ExcPending
							if v72 != 0 {
								return int32(0)
							} else {
								if v71 != 0 {
									*(*int64)(unsafe.Add(mBase, uint32(v71))) = int64(8589934592)
									*(*int32)(unsafe.Add(mBase, uint32(v71)+24)) = int32(-1)
									*(*int64)(unsafe.Add(mBase, uint32(v71)+12)) = int64(0)
									*(*int32)(unsafe.Add(mBase, uint32(v71)+20)) = v71 + int32(36)
									*(*int32)(unsafe.Add(mBase, uint32(v71)+8)) = v71 + int32(28)
									*(*int32)(unsafe.Add(mBase, uint32(l0)+120)) = v71
									v96 = v71
								} else {
									*(*int32)(unsafe.Add(mBase, uint32(l0)+24)) = int32(101)
									v88 = int32(0)
									*(*int32)(unsafe.Add(mBase, uint32(l0)+120)) = v88
									v91 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
									if v91 != 0 {
										v93 = v91
									} else {
										v93 = int32(12)
									}
									*(*int32)(unsafe.Add(mBase, uint32(l0)+12)) = v93
									v96 = v88
								}
								v97 = *(*int32)(unsafe.Add(mBase, uint32(v96)))
								*(*int32)(unsafe.Add(mBase, uint32(v96))) = v97 + int32(1)
								v101 = *(*int32)(unsafe.Add(mBase, uint32(v96)+8))
								*(*int32)(unsafe.Add(mBase, uint32(v101+v97<<(uint(int32(2))%32)))) = v53
								if v51 != v53 {
									v107 = *(*int32)(unsafe.Add(mBase, uint32(v96)))
									*(*int32)(unsafe.Add(mBase, uint32(v96))) = v107 + int32(1)
									v111 = *(*int32)(unsafe.Add(mBase, uint32(v96)+8))
									*(*int32)(unsafe.Add(mBase, uint32(v111+v107<<(uint(int32(2))%32)))) = v51
								} else {
								}
								return v96
							}
						}
					} else {
						v58 = *(*int32)(unsafe.Add(mBase, uint32(v54)+16))
						if v58 < int32(0) {
							F_pfree(m, v54)
							mBase = m.M
							v68 = m.ExcPending
							if v68 != 0 {
								return int32(0)
							} else {
								v71 = F_palloc_extended(m, int32(36), int32(2))
								mBase = m.M
								v72 = m.ExcPending
								if v72 != 0 {
									return int32(0)
								} else {
									if v71 != 0 {
										*(*int64)(unsafe.Add(mBase, uint32(v71))) = int64(8589934592)
										*(*int32)(unsafe.Add(mBase, uint32(v71)+24)) = int32(-1)
										*(*int64)(unsafe.Add(mBase, uint32(v71)+12)) = int64(0)
										*(*int32)(unsafe.Add(mBase, uint32(v71)+20)) = v71 + int32(36)
										*(*int32)(unsafe.Add(mBase, uint32(v71)+8)) = v71 + int32(28)
										*(*int32)(unsafe.Add(mBase, uint32(l0)+120)) = v71
										v96 = v71
									} else {
										*(*int32)(unsafe.Add(mBase, uint32(l0)+24)) = int32(101)
										v88 = int32(0)
										*(*int32)(unsafe.Add(mBase, uint32(l0)+120)) = v88
										v91 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
										if v91 != 0 {
											v93 = v91
										} else {
											v93 = int32(12)
										}
										*(*int32)(unsafe.Add(mBase, uint32(l0)+12)) = v93
										v96 = v88
									}
									v97 = *(*int32)(unsafe.Add(mBase, uint32(v96)))
									*(*int32)(unsafe.Add(mBase, uint32(v96))) = v97 + int32(1)
									v101 = *(*int32)(unsafe.Add(mBase, uint32(v96)+8))
									*(*int32)(unsafe.Add(mBase, uint32(v101+v97<<(uint(int32(2))%32)))) = v53
									if v51 != v53 {
										v107 = *(*int32)(unsafe.Add(mBase, uint32(v96)))
										*(*int32)(unsafe.Add(mBase, uint32(v96))) = v107 + int32(1)
										v111 = *(*int32)(unsafe.Add(mBase, uint32(v96)+8))
										*(*int32)(unsafe.Add(mBase, uint32(v111+v107<<(uint(int32(2))%32)))) = v51
									} else {
									}
									return v96
								}
							}
						} else {
							*(*int32)(unsafe.Add(mBase, uint32(v54)+24)) = int32(-1)
							v63 = int32(0)
							*(*int32)(unsafe.Add(mBase, uint32(v54)+12)) = v63
							*(*int32)(unsafe.Add(mBase, uint32(v54))) = v63
							v96 = v54
							v97 = *(*int32)(unsafe.Add(mBase, uint32(v96)))
							*(*int32)(unsafe.Add(mBase, uint32(v96))) = v97 + int32(1)
							v101 = *(*int32)(unsafe.Add(mBase, uint32(v96)+8))
							*(*int32)(unsafe.Add(mBase, uint32(v101+v97<<(uint(int32(2))%32)))) = v53
							if v51 != v53 {
								v107 = *(*int32)(unsafe.Add(mBase, uint32(v96)))
								*(*int32)(unsafe.Add(mBase, uint32(v96))) = v107 + int32(1)
								v111 = *(*int32)(unsafe.Add(mBase, uint32(v96)+8))
								*(*int32)(unsafe.Add(mBase, uint32(v111+v107<<(uint(int32(2))%32)))) = v51
							} else {
							}
							return v96
						}
					}
				} else {
					v71 = F_palloc_extended(m, int32(36), int32(2))
					mBase = m.M
					v72 = m.ExcPending
					if v72 != 0 {
						return int32(0)
					} else {
						if v71 != 0 {
							*(*int64)(unsafe.Add(mBase, uint32(v71))) = int64(8589934592)
							*(*int32)(unsafe.Add(mBase, uint32(v71)+24)) = int32(-1)
							*(*int64)(unsafe.Add(mBase, uint32(v71)+12)) = int64(0)
							*(*int32)(unsafe.Add(mBase, uint32(v71)+20)) = v71 + int32(36)
							*(*int32)(unsafe.Add(mBase, uint32(v71)+8)) = v71 + int32(28)
							*(*int32)(unsafe.Add(mBase, uint32(l0)+120)) = v71
							v96 = v71
						} else {
							*(*int32)(unsafe.Add(mBase, uint32(l0)+24)) = int32(101)
							v88 = int32(0)
							*(*int32)(unsafe.Add(mBase, uint32(l0)+120)) = v88
							v91 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
							if v91 != 0 {
								v93 = v91
							} else {
								v93 = int32(12)
							}
							*(*int32)(unsafe.Add(mBase, uint32(l0)+12)) = v93
							v96 = v88
						}
						v97 = *(*int32)(unsafe.Add(mBase, uint32(v96)))
						*(*int32)(unsafe.Add(mBase, uint32(v96))) = v97 + int32(1)
						v101 = *(*int32)(unsafe.Add(mBase, uint32(v96)+8))
						*(*int32)(unsafe.Add(mBase, uint32(v101+v97<<(uint(int32(2))%32)))) = v53
						if v51 != v53 {
							v107 = *(*int32)(unsafe.Add(mBase, uint32(v96)))
							*(*int32)(unsafe.Add(mBase, uint32(v96))) = v107 + int32(1)
							v111 = *(*int32)(unsafe.Add(mBase, uint32(v96)+8))
							*(*int32)(unsafe.Add(mBase, uint32(v111+v107<<(uint(int32(2))%32)))) = v51
						} else {
						}
						return v96
					}
				}
			}
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
	var v62 int32
	_ = v62
	var v63 int32
	_ = v63
	var v66 int32
	_ = v66
	var v68 int32
	_ = v68
	var v71 int32
	_ = v71
	var v74 int32
	_ = v74
	var v75 int32
	_ = v75
	var v76 int32
	_ = v76
	var v77 int32
	_ = v77
	var v84 int32
	_ = v84
	var v89 int32
	_ = v89
	var v90 int32
	_ = v90
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
	var v114 int32
	_ = v114
	var v119 int32
	_ = v119
	var v123 int32
	_ = v123
	var v126 int32
	_ = v126
	var v127 int32
	_ = v127
	var v128 int32
	_ = v128
	var v134 int32
	_ = v134
	var v136 int32
	_ = v136
	var v143 int32
	_ = v143
	var v146 int32
	_ = v146
	var v150 int32
	_ = v150
	var v151 int32
	_ = v151
	var v152 int32
	_ = v152
	var v154 int32
	_ = v154
	var v159 int32
	_ = v159
	var v163 int32
	_ = v163
	var v166 int32
	_ = v166
	var v167 int32
	_ = v167
	var v168 int32
	_ = v168
	var v174 int32
	_ = v174
	var v175 int32
	_ = v175
	var v176 int32
	_ = v176
	var v177 int32
	_ = v177
	var v178 int32
	_ = v178
	var v186 int32
	_ = v186
	var v187 int32
	_ = v187
	var v192 int32
	_ = v192
	var v196 int32
	_ = v196
	var v199 int32
	_ = v199
	var v203 int32
	_ = v203
	var v204 int32
	_ = v204
	var v211 int32
	_ = v211
	var v212 int32
	_ = v212
	var v217 int32
	_ = v217
	var v221 int32
	_ = v221
	var v224 int32
	_ = v224
	var v225 int32
	_ = v225
	var v233 int32
	_ = v233
	var v236 int32
	_ = v236
	var v237 int32
	_ = v237
	var v242 int32
	_ = v242
	var v243 int32
	_ = v243
	var v249 int32
	_ = v249
	var v252 int32
	_ = v252
	var v253 int32
	_ = v253
	var v259 int32
	_ = v259
	var v264 int32
	_ = v264
	v11 = m.G0
	v13 = v11 - int32(96)
	m.G0 = v13
	v16 = F_IndexGetRelation(m, l0, int32(1))
	mBase = m.M
	v17 = m.ExcPending
	if v17 != 0 {
		return
	} else {
		if v16 == int32(0) {
			*(*int32)(unsafe.Add(mBase, uint32(v13)+88)) = int32(-1)
			*(*int32)(unsafe.Add(mBase, uint32(v13)+92)) = int32(0)
			v24 = F_index_open(m, l0, l3)
			mBase = m.M
			v25 = m.ExcPending
			if v25 != 0 {
				return
			} else {
				v243 = v24
				F_errstart_cold(m, int32(21), int32(0))
				mBase = m.M
				v249 = m.ExcPending
				if v249 != 0 {
					return
				} else {
					F_errcode(m, int32(16908420))
					mBase = m.M
					v252 = m.ExcPending
					if v252 != 0 {
						return
					} else {
						v253 = *(*int32)(unsafe.Add(mBase, uint32(v243)+48))
						*(*int32)(unsafe.Add(mBase, uint32(v13))) = v253 + int32(4)
						F_errmsg(m, int32(_a_F_amcheck_lock_relation_and_check_0), v13)
						mBase = m.M
						v259 = m.ExcPending
						if v259 != 0 {
							return
						} else {
							F_errfinish(m, int32(_a_F_amcheck_lock_relation_and_check_1), int32(131), int32(_a_F_amcheck_lock_relation_and_check_2))
							mBase = m.M
							v264 = m.ExcPending
							if v264 != 0 {
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
				*(*int32)(unsafe.Add(mBase, uint32(v13+int32(92)))) = v33
				v36 = *(*int32)(unsafe.Add(mBase, _c_F_amcheck_lock_relation_and_check[1]))
				*(*int32)(unsafe.Add(mBase, uint32(v13+int32(88)))) = v36
				v38 = *(*int32)(unsafe.Add(mBase, uint32(v26)+48))
				v39 = *(*int32)(unsafe.Add(mBase, uint32(v38)+80))
				v40 = *(*int32)(unsafe.Add(mBase, uint32(v13)+88))
				*(*int32)(unsafe.Add(mBase, _c_F_amcheck_lock_relation_and_check[1])) = v40 | int32(2)
				*(*int32)(unsafe.Add(mBase, _c_F_amcheck_lock_relation_and_check[0])) = v39
				v48 = int32(_a_F_amcheck_lock_relation_and_check_3)
				v50 = *(*int32)(unsafe.Add(mBase, _c_F_amcheck_lock_relation_and_check[2]))
				v52 = v50 + int32(1)
				*(*int32)(unsafe.Add(mBase, _c_F_amcheck_lock_relation_and_check[2])) = v52
				F_RestrictSearchPath(m)
				mBase = m.M
				v55 = m.ExcPending
				if v55 != 0 {
					return
				} else {
					v56 = F_index_open(m, l0, l3)
					mBase = m.M
					v57 = m.ExcPending
					if v57 != 0 {
						return
					} else {
						v59 = F_IndexGetRelation(m, l0, int32(0))
						mBase = m.M
						v60 = m.ExcPending
						if v60 != 0 {
							return
						} else {
							if v59 != v16 {
								v243 = v56
								F_errstart_cold(m, int32(21), int32(0))
								mBase = m.M
								v249 = m.ExcPending
								if v249 != 0 {
									return
								} else {
									F_errcode(m, int32(16908420))
									mBase = m.M
									v252 = m.ExcPending
									if v252 != 0 {
										return
									} else {
										v253 = *(*int32)(unsafe.Add(mBase, uint32(v243)+48))
										*(*int32)(unsafe.Add(mBase, uint32(v13))) = v253 + int32(4)
										F_errmsg(m, int32(_a_F_amcheck_lock_relation_and_check_0), v13)
										mBase = m.M
										v259 = m.ExcPending
										if v259 != 0 {
											return
										} else {
											F_errfinish(m, int32(_a_F_amcheck_lock_relation_and_check_1), int32(131), int32(_a_F_amcheck_lock_relation_and_check_2))
											mBase = m.M
											v264 = m.ExcPending
											if v264 != 0 {
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
								v62 = *(*int32)(unsafe.Add(mBase, uint32(v56)+48))
								v63 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v62)+119)))
								if v63 == int32(105) {
									v66 = *(*int32)(unsafe.Add(mBase, uint32(v62)+84))
									if v66 != l1 {
										F_errstart_cold(m, int32(21), int32(0))
										mBase = m.M
										v163 = m.ExcPending
										if v163 != 0 {
											return
										} else {
											F_errcode(m, int32(1088))
											mBase = m.M
											v166 = m.ExcPending
											if v166 != 0 {
												return
											} else {
												v167 = F_get_am_name(m, l1)
												mBase = m.M
												v168 = m.ExcPending
												if v168 != 0 {
													return
												} else {
													*(*int32)(unsafe.Add(mBase, uint32(v13)+80)) = v167
													F_errmsg(m, int32(_a_F_amcheck_lock_relation_and_check_4), v13+int32(80))
													mBase = m.M
													v174 = m.ExcPending
													if v174 != 0 {
														return
													} else {
														v175 = *(*int32)(unsafe.Add(mBase, uint32(v56)+48))
														v176 = *(*int32)(unsafe.Add(mBase, uint32(v175)+84))
														v177 = F_get_am_name(m, v176)
														mBase = m.M
														v178 = m.ExcPending
														if v178 != 0 {
															return
														} else {
															*(*int32)(unsafe.Add(mBase, uint32(v13)+68)) = v177
															*(*int32)(unsafe.Add(mBase, uint32(v13)+64)) = v175 + int32(4)
															v186 = F_errdetail(m, int32(_a_F_amcheck_lock_relation_and_check_5), v13-int32(-64))
															mBase = m.M
															v187 = m.ExcPending
															if v187 != 0 {
																return
															} else {
																F_errfinish(m, int32(_a_F_amcheck_lock_relation_and_check_1), int32(175), int32(_a_F_amcheck_lock_relation_and_check_6))
																mBase = m.M
																v192 = m.ExcPending
																if v192 != 0 {
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
									} else {
										v68 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v62)+118)))
										if v68 == int32(116) {
											v71 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v56)+24)))
											if v71 == int32(0) {
												F_errstart_cold(m, int32(21), int32(0))
												mBase = m.M
												v196 = m.ExcPending
												if v196 != 0 {
													return
												} else {
													F_errcode(m, int32(1088))
													mBase = m.M
													v199 = m.ExcPending
													if v199 != 0 {
														return
													} else {
														F_errmsg(m, int32(_a_F_amcheck_lock_relation_and_check_7), int32(0))
														mBase = m.M
														v203 = m.ExcPending
														if v203 != 0 {
															return
														} else {
															v204 = *(*int32)(unsafe.Add(mBase, uint32(v56)+48))
															*(*int32)(unsafe.Add(mBase, uint32(v13)+32)) = v204 + int32(4)
															v211 = F_errdetail(m, int32(_a_F_amcheck_lock_relation_and_check_8), v13+int32(32))
															mBase = m.M
															v212 = m.ExcPending
															if v212 != 0 {
																return
															} else {
																F_errfinish(m, int32(_a_F_amcheck_lock_relation_and_check_1), int32(182), int32(_a_F_amcheck_lock_relation_and_check_6))
																mBase = m.M
																v217 = m.ExcPending
																if v217 != 0 {
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
												v74 = *(*int32)(unsafe.Add(mBase, uint32(v56)+192))
												v75 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v74)+18)))
												if v75 != 0 {
													m.T0[l2].(func(*base.Module, int32, int32, int32, int32))(m, v56, v26, l4, base.B2i32(l3 == int32(5)))
													mBase = m.M
													v123 = m.ExcPending
													if v123 != 0 {
														return
													} else {
														F_AtEOXact_GUC(m, int32(0), v52)
														mBase = m.M
														v126 = m.ExcPending
														if v126 != 0 {
															return
														} else {
															v127 = *(*int32)(unsafe.Add(mBase, uint32(v13)+92))
															v128 = *(*int32)(unsafe.Add(mBase, uint32(v13)+88))
															*(*int32)(unsafe.Add(mBase, _c_F_amcheck_lock_relation_and_check[1])) = v128
															*(*int32)(unsafe.Add(mBase, _c_F_amcheck_lock_relation_and_check[0])) = v127
															F_relation_close(m, v56, l3)
															mBase = m.M
															v134 = m.ExcPending
															if v134 != 0 {
																return
															} else {
																F_relation_close(m, v26, l3)
																mBase = m.M
																v136 = m.ExcPending
																if v136 != 0 {
																	return
																} else {
																	m.G0 = v13 + int32(96)
																	return
																}
															}
														}
													}
												} else {
													F_errstart_cold(m, int32(21), int32(0))
													mBase = m.M
													v221 = m.ExcPending
													if v221 != 0 {
														return
													} else {
														F_errcode(m, int32(325))
														mBase = m.M
														v224 = m.ExcPending
														if v224 != 0 {
															return
														} else {
															v225 = *(*int32)(unsafe.Add(mBase, uint32(v56)+48))
															*(*int32)(unsafe.Add(mBase, uint32(v13)+16)) = v225 + int32(4)
															F_errmsg(m, int32(_a_F_amcheck_lock_relation_and_check_9), v13+int32(16))
															mBase = m.M
															v233 = m.ExcPending
															if v233 != 0 {
																return
															} else {
																v236 = F_errdetail(m, int32(_a_F_amcheck_lock_relation_and_check_10), int32(0))
																mBase = m.M
																v237 = m.ExcPending
																if v237 != 0 {
																	return
																} else {
																	F_errfinish(m, int32(_a_F_amcheck_lock_relation_and_check_1), int32(189), int32(_a_F_amcheck_lock_relation_and_check_6))
																	mBase = m.M
																	v242 = m.ExcPending
																	if v242 != 0 {
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
										} else {
											v76 = *(*int32)(unsafe.Add(mBase, uint32(v56)+192))
											v77 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v76)+18)))
											if v77 == int32(0) {
												F_errstart_cold(m, int32(21), int32(0))
												mBase = m.M
												v221 = m.ExcPending
												if v221 != 0 {
													return
												} else {
													F_errcode(m, int32(325))
													mBase = m.M
													v224 = m.ExcPending
													if v224 != 0 {
														return
													} else {
														v225 = *(*int32)(unsafe.Add(mBase, uint32(v56)+48))
														*(*int32)(unsafe.Add(mBase, uint32(v13)+16)) = v225 + int32(4)
														F_errmsg(m, int32(_a_F_amcheck_lock_relation_and_check_9), v13+int32(16))
														mBase = m.M
														v233 = m.ExcPending
														if v233 != 0 {
															return
														} else {
															v236 = F_errdetail(m, int32(_a_F_amcheck_lock_relation_and_check_10), int32(0))
															mBase = m.M
															v237 = m.ExcPending
															if v237 != 0 {
																return
															} else {
																F_errfinish(m, int32(_a_F_amcheck_lock_relation_and_check_1), int32(189), int32(_a_F_amcheck_lock_relation_and_check_6))
																mBase = m.M
																v242 = m.ExcPending
																if v242 != 0 {
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
												if v68 != int32(117) {
													m.T0[l2].(func(*base.Module, int32, int32, int32, int32))(m, v56, v26, l4, base.B2i32(l3 == int32(5)))
													mBase = m.M
													v123 = m.ExcPending
													if v123 != 0 {
														return
													} else {
														F_AtEOXact_GUC(m, int32(0), v52)
														mBase = m.M
														v126 = m.ExcPending
														if v126 != 0 {
															return
														} else {
															v127 = *(*int32)(unsafe.Add(mBase, uint32(v13)+92))
															v128 = *(*int32)(unsafe.Add(mBase, uint32(v13)+88))
															*(*int32)(unsafe.Add(mBase, _c_F_amcheck_lock_relation_and_check[1])) = v128
															*(*int32)(unsafe.Add(mBase, _c_F_amcheck_lock_relation_and_check[0])) = v127
															F_relation_close(m, v56, l3)
															mBase = m.M
															v134 = m.ExcPending
															if v134 != 0 {
																return
															} else {
																F_relation_close(m, v26, l3)
																mBase = m.M
																v136 = m.ExcPending
																if v136 != 0 {
																	return
																} else {
																	m.G0 = v13 + int32(96)
																	return
																}
															}
														}
													}
												} else {
													v84 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_amcheck_lock_relation_and_check[3])))
													if v84 == int32(1) {
														v89 = *(*int32)(unsafe.Add(mBase, _c_F_amcheck_lock_relation_and_check[4]))
														v90 = *(*int32)(unsafe.Add(mBase, uint32(v89)+308))
														v92 = base.B2i32(v90 != int32(2))
														*(*uint8)(unsafe.Add(mBase, _c_F_amcheck_lock_relation_and_check[3])) = uint8(v92)
														v94 = v92
													} else {
														v94 = int32(0)
													}
													if v94 == int32(0) {
														m.T0[l2].(func(*base.Module, int32, int32, int32, int32))(m, v56, v26, l4, base.B2i32(l3 == int32(5)))
														mBase = m.M
														v123 = m.ExcPending
														if v123 != 0 {
															return
														} else {
															F_AtEOXact_GUC(m, int32(0), v52)
															mBase = m.M
															v126 = m.ExcPending
															if v126 != 0 {
																return
															} else {
																v127 = *(*int32)(unsafe.Add(mBase, uint32(v13)+92))
																v128 = *(*int32)(unsafe.Add(mBase, uint32(v13)+88))
																*(*int32)(unsafe.Add(mBase, _c_F_amcheck_lock_relation_and_check[1])) = v128
																*(*int32)(unsafe.Add(mBase, _c_F_amcheck_lock_relation_and_check[0])) = v127
																F_relation_close(m, v56, l3)
																mBase = m.M
																v134 = m.ExcPending
																if v134 != 0 {
																	return
																} else {
																	F_relation_close(m, v26, l3)
																	mBase = m.M
																	v136 = m.ExcPending
																	if v136 != 0 {
																		return
																	} else {
																		m.G0 = v13 + int32(96)
																		return
																	}
																}
															}
														}
													} else {
														v99 = F_errstart(m, int32(18), int32(0))
														mBase = m.M
														v100 = m.ExcPending
														if v100 != 0 {
															return
														} else {
															if v99 == int32(0) {
																F_AtEOXact_GUC(m, int32(0), v52)
																mBase = m.M
																v126 = m.ExcPending
																if v126 != 0 {
																	return
																} else {
																	v127 = *(*int32)(unsafe.Add(mBase, uint32(v13)+92))
																	v128 = *(*int32)(unsafe.Add(mBase, uint32(v13)+88))
																	*(*int32)(unsafe.Add(mBase, _c_F_amcheck_lock_relation_and_check[1])) = v128
																	*(*int32)(unsafe.Add(mBase, _c_F_amcheck_lock_relation_and_check[0])) = v127
																	F_relation_close(m, v56, l3)
																	mBase = m.M
																	v134 = m.ExcPending
																	if v134 != 0 {
																		return
																	} else {
																		F_relation_close(m, v26, l3)
																		mBase = m.M
																		v136 = m.ExcPending
																		if v136 != 0 {
																			return
																		} else {
																			m.G0 = v13 + int32(96)
																			return
																		}
																	}
																}
															} else {
																F_errcode(m, int32(100663618))
																mBase = m.M
																v105 = m.ExcPending
																if v105 != 0 {
																	return
																} else {
																	v106 = *(*int32)(unsafe.Add(mBase, uint32(v56)+48))
																	*(*int32)(unsafe.Add(mBase, uint32(v13)+48)) = v106 + int32(4)
																	F_errmsg(m, int32(_a_F_amcheck_lock_relation_and_check_11), v13+int32(48))
																	mBase = m.M
																	v114 = m.ExcPending
																	if v114 != 0 {
																		return
																	} else {
																		F_errfinish(m, int32(_a_F_amcheck_lock_relation_and_check_1), int32(47), int32(_a_F_amcheck_lock_relation_and_check_12))
																		mBase = m.M
																		v119 = m.ExcPending
																		if v119 != 0 {
																			return
																		} else {
																			F_AtEOXact_GUC(m, int32(0), v52)
																			mBase = m.M
																			v126 = m.ExcPending
																			if v126 != 0 {
																				return
																			} else {
																				v127 = *(*int32)(unsafe.Add(mBase, uint32(v13)+92))
																				v128 = *(*int32)(unsafe.Add(mBase, uint32(v13)+88))
																				*(*int32)(unsafe.Add(mBase, _c_F_amcheck_lock_relation_and_check[1])) = v128
																				*(*int32)(unsafe.Add(mBase, _c_F_amcheck_lock_relation_and_check[0])) = v127
																				F_relation_close(m, v56, l3)
																				mBase = m.M
																				v134 = m.ExcPending
																				if v134 != 0 {
																					return
																				} else {
																					F_relation_close(m, v26, l3)
																					mBase = m.M
																					v136 = m.ExcPending
																					if v136 != 0 {
																						return
																					} else {
																						m.G0 = v13 + int32(96)
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
									}
								} else {
									F_errstart_cold(m, int32(21), int32(0))
									mBase = m.M
									v143 = m.ExcPending
									if v143 != 0 {
										return
									} else {
										F_errcode(m, int32(151027844))
										mBase = m.M
										v146 = m.ExcPending
										if v146 != 0 {
											return
										} else {
											F_errmsg(m, int32(_a_F_amcheck_lock_relation_and_check_13), int32(0))
											mBase = m.M
											v150 = m.ExcPending
											if v150 != 0 {
												return
											} else {
												v151 = *(*int32)(unsafe.Add(mBase, uint32(v56)+48))
												v152 = int32(*(*int8)(unsafe.Add(mBase, uint32(v151)+119)))
												F_errdetail_relkind_not_supported(m, v152)
												mBase = m.M
												v154 = m.ExcPending
												if v154 != 0 {
													return
												} else {
													F_errfinish(m, int32(_a_F_amcheck_lock_relation_and_check_1), int32(168), int32(_a_F_amcheck_lock_relation_and_check_6))
													mBase = m.M
													v159 = m.ExcPending
													if v159 != 0 {
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
					}
				}
			}
		}
	}
}
func F_anyarray_recv(m *base.Module, l0 int32) int64 {
	var v7 int64
	_ = v7
	var v10 int32
	_ = v10
	v7 = Fn14235(m, l0, int32(_a_F_anyarray_recv_0), int32(155), int32(_a_F_anyarray_recv_1), int32(_a_F_anyarray_recv_2), int32(_a_F_anyarray_recv_3))
	v10 = m.ExcPending
	if v10 != 0 {
		return int64(0)
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
func F_anycompatiblemultirange_in(m *base.Module, l0 int32) int64 {
	var v7 int64
	_ = v7
	var v10 int32
	_ = v10
	v7 = Fn14235(m, l0, int32(_a_F_anycompatiblemultirange_in_0), int32(246), int32(_a_F_anycompatiblemultirange_in_1), int32(_a_F_anycompatiblemultirange_in_2), int32(_a_F_anycompatiblemultirange_in_3))
	v10 = m.ExcPending
	if v10 != 0 {
		return int64(0)
	} else {
		return v7
	}
}
func F_anyenum_out(m *base.Module, l0 int32) int64 {
	var v2 int64
	_ = v2
	var v5 int32
	_ = v5
	v2 = F_enum_out(m, l0)
	v5 = m.ExcPending
	if v5 != 0 {
		return int64(0)
	} else {
		return v2
	}
}
func F_anyrange_out(m *base.Module, l0 int32) int64 {
	var v2 int64
	_ = v2
	var v5 int32
	_ = v5
	v2 = F_range_out(m, l0)
	v5 = m.ExcPending
	if v5 != 0 {
		return int64(0)
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
func F_appendElement(m *base.Module, l0 int32, l1 int32, l2 int32) {
	mBase := m.M
	_ = mBase
	var v8 int32
	_ = v8
	var v10 int32
	_ = v10
	var v12 int32
	_ = v12
	var v13 int32
	_ = v13
	var v14 int32
	_ = v14
	var v16 int32
	_ = v16
	var v19 int32
	_ = v19
	var v21 int32
	_ = v21
	var v24 int32
	_ = v24
	var v26 int32
	_ = v26
	var v29 int32
	_ = v29
	var v30 int32
	_ = v30
	var v32 int32
	_ = v32
	var v33 int32
	_ = v33
	var v34 int32
	_ = v34
	var v37 int32
	_ = v37
	var v38 int64
	_ = v38
	var v40 int64
	_ = v40
	var v42 int64
	_ = v42
	var v44 int64
	_ = v44
	var v46 int32
	_ = v46
	var v50 int32
	_ = v50
	var v52 int32
	_ = v52
	var v59 int32
	_ = v59
	var v62 int32
	_ = v62
	var v67 int32
	_ = v67
	var v72 int32
	_ = v72
	v8 = m.G0
	v10 = v8 - int32(16)
	m.G0 = v10
	v12 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	v13 = *(*int32)(unsafe.Add(mBase, uint32(v12)+8))
	v14 = *(*int32)(unsafe.Add(mBase, uint32(v12)+32))
	if base.Ui32(v13) < base.Ui32(v14) {
		v16 = *(*int32)(unsafe.Add(mBase, uint32(v12)+12))
		v33 = v13
		v34 = v16
		v37 = v34 + v33<<(uint(int32(5))%32)
		v38 = *(*int64)(unsafe.Add(mBase, uint32(l1)+24))
		*(*int64)(unsafe.Add(mBase, uint32(v37)+24)) = v38
		v40 = *(*int64)(unsafe.Add(mBase, uint32(l1)+16))
		*(*int64)(unsafe.Add(mBase, uint32(v37)+16)) = v40
		v42 = *(*int64)(unsafe.Add(mBase, uint32(l1)+8))
		*(*int64)(unsafe.Add(mBase, uint32(v37)+8)) = v42
		v44 = *(*int64)(unsafe.Add(mBase, uint32(l1)))
		*(*int64)(unsafe.Add(mBase, uint32(v37))) = v44
		v46 = *(*int32)(unsafe.Add(mBase, uint32(v12)+8))
		*(*int32)(unsafe.Add(mBase, uint32(v12)+8)) = v46 + int32(1)
		if l2 != 0 {
			v50 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
			F_copyScalarSubstructure(m, v37, v50)
			mBase = m.M
			v52 = m.ExcPending
			if v52 != 0 {
				return
			} else {
				m.G0 = v10 + int32(16)
				return
			}
		} else {
			m.G0 = v10 + int32(16)
			return
		}
	} else {
		if base.Ui32(int32(33554431)) <= base.Ui32(v13) {
			F_errstart_cold(m, int32(21), int32(0))
			mBase = m.M
			v59 = m.ExcPending
			if v59 != 0 {
				return
			} else {
				F_errcode(m, int32(261))
				mBase = m.M
				v62 = m.ExcPending
				if v62 != 0 {
					return
				} else {
					*(*int32)(unsafe.Add(mBase, uint32(v10))) = int32(33554431)
					F_errmsg(m, int32(_a_F_appendElement_0), v10)
					mBase = m.M
					v67 = m.ExcPending
					if v67 != 0 {
						return
					} else {
						F_errfinish(m, int32(_a_F_appendElement_1), int32(851), int32(_a_F_appendElement_2))
						mBase = m.M
						v72 = m.ExcPending
						if v72 != 0 {
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
			v19 = int32(33554431)
			v21 = v14 << (uint(int32(1)) % 32)
			if base.Ui32(v19) <= base.Ui32(v21) {
				v24 = v19
			} else {
				v24 = v21
			}
			*(*int32)(unsafe.Add(mBase, uint32(v12)+32)) = v24
			v26 = *(*int32)(unsafe.Add(mBase, uint32(v12)+12))
			v29 = F_repalloc(m, v26, v24<<(uint(int32(5))%32))
			mBase = m.M
			v30 = m.ExcPending
			if v30 != 0 {
				return
			} else {
				*(*int32)(unsafe.Add(mBase, uint32(v12)+12)) = v29
				v32 = *(*int32)(unsafe.Add(mBase, uint32(v12)+8))
				v33 = v32
				v34 = v29
				v37 = v34 + v33<<(uint(int32(5))%32)
				v38 = *(*int64)(unsafe.Add(mBase, uint32(l1)+24))
				*(*int64)(unsafe.Add(mBase, uint32(v37)+24)) = v38
				v40 = *(*int64)(unsafe.Add(mBase, uint32(l1)+16))
				*(*int64)(unsafe.Add(mBase, uint32(v37)+16)) = v40
				v42 = *(*int64)(unsafe.Add(mBase, uint32(l1)+8))
				*(*int64)(unsafe.Add(mBase, uint32(v37)+8)) = v42
				v44 = *(*int64)(unsafe.Add(mBase, uint32(l1)))
				*(*int64)(unsafe.Add(mBase, uint32(v37))) = v44
				v46 = *(*int32)(unsafe.Add(mBase, uint32(v12)+8))
				*(*int32)(unsafe.Add(mBase, uint32(v12)+8)) = v46 + int32(1)
				if l2 != 0 {
					v50 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
					F_copyScalarSubstructure(m, v37, v50)
					mBase = m.M
					v52 = m.ExcPending
					if v52 != 0 {
						return
					} else {
						m.G0 = v10 + int32(16)
						return
					}
				} else {
					m.G0 = v10 + int32(16)
					return
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
	var v225 int32
	_ = v225
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
	v225 = F_errdetail(m, int32(_a_F_apply_typmod_9), v14)
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
	var v7 int32
	_ = v7
	var v11 int32
	_ = v11
	var v12 int32
	_ = v12
	var v14 int32
	_ = v14
	var v15 int32
	_ = v15
	var v17 int32
	_ = v17
	var v20 int32
	_ = v20
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
	var v32 int32
	_ = v32
	var v33 int32
	_ = v33
	var v35 int32
	_ = v35
	var v36 int32
	_ = v36
	var v38 int32
	_ = v38
	var v40 int32
	_ = v40
	var v42 int32
	_ = v42
	var v44 int32
	_ = v44
	v7 = *(*int32)(unsafe.Add(mBase, _c_F_apw_read_stream_next_block[0]))
	if v7 != 0 {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	F_ProcessInterrupts(m)
	mBase = m.M
	v11 = m.ExcPending
	if v11 != 0 {
		goto L4
	} else {
		goto L5
	}
L2:
	;
	goto L3
L3:
	;
	v12 = *(*int32)(unsafe.Add(mBase, uint32(l1)+4))
	v14 = *(*int32)(unsafe.Add(mBase, _c_F_apw_read_stream_next_block[1]))
	v15 = *(*int32)(unsafe.Add(mBase, uint32(v14)+36))
	if v15 < v12 {
		goto L6
	} else {
		goto L7
	}
L4:
	;
	return int32(0)
L5:
	;
	goto L3
L6:
	;
	v17 = v12
	goto L8
L7:
	;
	v17 = v15
	goto L8
L8:
	;
	v20 = v12
	goto L9
L9:
	;
	v23 = int32(-1)
	if v20 == v17 {
		v44 = v23
		goto L11
	} else {
		goto L12
	}
L10:
	;
	return v44
L11:
	;
	goto L10
L12:
	;
	v25 = *(*int32)(unsafe.Add(mBase, uint32(l1)))
	v28 = v25 + v20*int32(20)
	v29 = *(*int32)(unsafe.Add(mBase, uint32(v28)+4))
	v30 = *(*int32)(unsafe.Add(mBase, uint32(l1)+8))
	if v29 != v30 {
		v44 = v23
		goto L11
	} else {
		goto L13
	}
L13:
	;
	v32 = *(*int32)(unsafe.Add(mBase, uint32(v28)+8))
	v33 = *(*int32)(unsafe.Add(mBase, uint32(l1)+12))
	if v32 != v33 {
		v44 = v23
		goto L11
	} else {
		goto L14
	}
L14:
	;
	v35 = *(*int32)(unsafe.Add(mBase, uint32(v28)+12))
	v36 = *(*int32)(unsafe.Add(mBase, uint32(l1)+16))
	if v35 != v36 {
		v44 = v23
		goto L11
	} else {
		goto L15
	}
L15:
	;
	v38 = *(*int32)(unsafe.Add(mBase, uint32(v28)+16))
	v40 = v20 + int32(1)
	*(*int32)(unsafe.Add(mBase, uint32(l1)+4)) = v40
	v42 = *(*int32)(unsafe.Add(mBase, uint32(l1)+20))
	if base.Ui32(v42) <= base.Ui32(v38) {
		v20 = v40
		goto L9
	} else {
		goto L16
	}
L16:
	;
	v44 = v38
	goto L11
}
func F_arraycontains(m *base.Module, l0 int32) int64 {
	mBase := m.M
	_ = mBase
	var v5 int64
	_ = v5
	var v6 int32
	_ = v6
	var v9 int32
	_ = v9
	var v10 int64
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
	v5 = *(*int64)(unsafe.Add(mBase, uint32(l0)+24))
	v6 = F_DatumGetAnyArrayP(m, v5)
	mBase = m.M
	v9 = m.ExcPending
	if v9 != 0 {
		return int64(0)
	} else {
		v10 = *(*int64)(unsafe.Add(mBase, uint32(l0)+40))
		v11 = F_DatumGetAnyArrayP(m, v10)
		mBase = m.M
		v12 = m.ExcPending
		if v12 != 0 {
			return int64(0)
		} else {
			v13 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
			v15 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
			v18 = F_array_contain_compare(m, v11, v6, v13, int32(1), v15+int32(16))
			mBase = m.M
			v19 = m.ExcPending
			if v19 != 0 {
				return int64(0)
			} else {
				v20 = *(*int32)(unsafe.Add(mBase, uint32(v6)))
				if v20 == int32(-1) {
					v27 = *(*int32)(unsafe.Add(mBase, uint32(v11)))
					if v27 == int32(-1) {
						return base.I64_extend_i32_u(v18)
					} else {
						v30 = *(*int32)(unsafe.Add(mBase, uint32(l0)+40))
						if v11 == v30 {
							return base.I64_extend_i32_u(v18)
						} else {
							F_pfree(m, v11)
							mBase = m.M
							v33 = m.ExcPending
							if v33 != 0 {
								return int64(0)
							} else {
								return base.I64_extend_i32_u(v18)
							}
						}
					}
				} else {
					v23 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
					if v6 == v23 {
						v27 = *(*int32)(unsafe.Add(mBase, uint32(v11)))
						if v27 == int32(-1) {
							return base.I64_extend_i32_u(v18)
						} else {
							v30 = *(*int32)(unsafe.Add(mBase, uint32(l0)+40))
							if v11 == v30 {
								return base.I64_extend_i32_u(v18)
							} else {
								F_pfree(m, v11)
								mBase = m.M
								v33 = m.ExcPending
								if v33 != 0 {
									return int64(0)
								} else {
									return base.I64_extend_i32_u(v18)
								}
							}
						}
					} else {
						F_pfree(m, v6)
						mBase = m.M
						v26 = m.ExcPending
						if v26 != 0 {
							return int64(0)
						} else {
							v27 = *(*int32)(unsafe.Add(mBase, uint32(v11)))
							if v27 == int32(-1) {
								return base.I64_extend_i32_u(v18)
							} else {
								v30 = *(*int32)(unsafe.Add(mBase, uint32(l0)+40))
								if v11 == v30 {
									return base.I64_extend_i32_u(v18)
								} else {
									F_pfree(m, v11)
									mBase = m.M
									v33 = m.ExcPending
									if v33 != 0 {
										return int64(0)
									} else {
										return base.I64_extend_i32_u(v18)
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
func F_arrayoverlap(m *base.Module, l0 int32) int64 {
	var v3 int64
	_ = v3
	var v6 int32
	_ = v6
	v3 = Fn14229(m, l0, int32(0))
	v6 = m.ExcPending
	if v6 != 0 {
		return int64(0)
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
	var v24 int32
	_ = v24
	var v25 int32
	_ = v25
	var v27 int32
	_ = v27
	var v29 int32
	_ = v29
	var v34 int32
	_ = v34
	var v35 int32
	_ = v35
	var v36 int32
	_ = v36
	var v39 int32
	_ = v39
	var v41 int32
	_ = v41
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
			v21 = int32(0)
			v24 = base.AtomicRmwOr32(m, v21, int32(_a_F_assign_application_name_1), v21)
			v25 = *(*int32)(unsafe.Add(mBase, uint32(v6)+212))
			if v9 != 0 {
				base.MemoryCopy(m, v25, l0, v9)
			} else {
			}
			v27 = *(*int32)(unsafe.Add(mBase, uint32(v6)+212))
			v29 = int32(0)
			*(*uint8)(unsafe.Add(mBase, uint32(v27+v9))) = uint8(v29)
			v34 = base.AtomicRmwOr32(m, v29, int32(_a_F_assign_application_name_1), v29)
			v35 = *(*int32)(unsafe.Add(mBase, uint32(v6)))
			v36 = int32(1)
			*(*int32)(unsafe.Add(mBase, uint32(v6))) = v35 + v36
			v39 = int32(_a_F_assign_application_name_0)
			v41 = *(*int32)(unsafe.Add(mBase, _c_F_assign_application_name[1]))
			*(*int32)(unsafe.Add(mBase, _c_F_assign_application_name[1])) = v41 - v36
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
					F_errfinish(m, int32(_a_F_attnumTypeId_1), int32(3694), int32(_a_F_attnumTypeId_2))
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
			v29 = v16 + v17<<(uint(int32(3))%32) + l1*int32(100) - int32(72)
			v30 = *(*int32)(unsafe.Add(mBase, uint32(v29)+68))
			m.G0 = v7 + int32(16)
			return v30
		}
	}
}
