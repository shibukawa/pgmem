package p0

import (
	base "github.com/shibukawa/pgmem/internal/aot/pgaot/base"
	"unsafe"
)

func F_CreateParallelContext(m *base.Module, l0 int32, l1 int32, l2 int32) int32 {
	mBase := m.M
	_ = mBase
	var v7 int32
	_ = v7
	var v8 int32
	_ = v8
	var v11 int32
	_ = v11
	var v14 int32
	_ = v14
	var v17 int32
	_ = v17
	var v19 int32
	_ = v19
	var v20 int32
	_ = v20
	var v24 int32
	_ = v24
	var v25 int32
	_ = v25
	var v27 int32
	_ = v27
	var v28 int32
	_ = v28
	var v31 int32
	_ = v31
	var v36 int32
	_ = v36
	var v39 int32
	_ = v39
	var v43 int32
	_ = v43
	v7 = int32(_a_F_CreateParallelContext_0)
	v8 = *(*int32)(unsafe.Add(mBase, _c_F_CreateParallelContext[0]))
	v11 = *(*int32)(unsafe.Add(mBase, _c_F_CreateParallelContext[1]))
	*(*int32)(unsafe.Add(mBase, _c_F_CreateParallelContext[0])) = v11
	v14 = F_palloc0(m, int32(68))
	mBase = m.M
	v17 = m.ExcPending
	if v17 != 0 {
		return int32(0)
	} else {
		v19 = *(*int32)(unsafe.Add(mBase, _c_F_CreateParallelContext[2]))
		v20 = *(*int32)(unsafe.Add(mBase, uint32(v19)+8))
		*(*int32)(unsafe.Add(mBase, uint32(v14)+16)) = l2
		*(*int32)(unsafe.Add(mBase, uint32(v14)+12)) = l2
		*(*int32)(unsafe.Add(mBase, uint32(v14)+8)) = v20
		v24 = F_pstrdup(m, l0)
		mBase = m.M
		v25 = m.ExcPending
		if v25 != 0 {
			return int32(0)
		} else {
			*(*int32)(unsafe.Add(mBase, uint32(v14)+24)) = v24
			v27 = F_pstrdup(m, l1)
			mBase = m.M
			v28 = m.ExcPending
			if v28 != 0 {
				return int32(0)
			} else {
				*(*int32)(unsafe.Add(mBase, uint32(v14)+28)) = v27
				v31 = *(*int32)(unsafe.Add(mBase, _c_F_CreateParallelContext[3]))
				*(*int64)(unsafe.Add(mBase, uint32(v14)+36)) = int64(0)
				*(*int32)(unsafe.Add(mBase, uint32(v14)+32)) = v31
				v36 = *(*int32)(unsafe.Add(mBase, _c_F_CreateParallelContext[4]))
				if v36 == int32(0) {
					v39 = int32(_a_F_CreateParallelContext_1)
					*(*int32)(unsafe.Add(mBase, _c_F_CreateParallelContext[5])) = v39
					v43 = v39
				} else {
					v43 = v36
				}
				*(*int32)(unsafe.Add(mBase, uint32(v14))) = int32(_a_F_CreateParallelContext_1)
				*(*int32)(unsafe.Add(mBase, uint32(v14)+4)) = v43
				*(*int32)(unsafe.Add(mBase, uint32(v43))) = v14
				*(*int32)(unsafe.Add(mBase, _c_F_CreateParallelContext[0])) = v8
				*(*int32)(unsafe.Add(mBase, _c_F_CreateParallelContext[4])) = v14
				return v14
			}
		}
	}
}
func F_DestroyParallelContext(m *base.Module, l0 int32) {
	mBase := m.M
	_ = mBase
	var v7 int32
	_ = v7
	var v8 int32
	_ = v8
	var v10 int32
	_ = v10
	var v12 int32
	_ = v12
	var v15 int32
	_ = v15
	var v20 int32
	_ = v20
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
	var v31 int32
	_ = v31
	var v32 int32
	_ = v32
	var v33 int32
	_ = v33
	var v35 int32
	_ = v35
	var v39 int32
	_ = v39
	var v40 int32
	_ = v40
	var v41 int64
	_ = v41
	var v44 int32
	_ = v44
	var v45 int64
	_ = v45
	var v47 int32
	_ = v47
	var v50 int32
	_ = v50
	var v54 int32
	_ = v54
	var v57 int32
	_ = v57
	var v59 int32
	_ = v59
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
	var v73 int32
	_ = v73
	var v74 int32
	_ = v74
	var v78 int32
	_ = v78
	var v86 int32
	_ = v86
	var v88 int32
	_ = v88
	var v91 int32
	_ = v91
	var v93 int32
	_ = v93
	var v96 int32
	_ = v96
	var v98 int32
	_ = v98
	var v103 int32
	_ = v103
	var v104 int32
	_ = v104
	var v106 int32
	_ = v106
	var v110 int32
	_ = v110
	var v112 int32
	_ = v112
	var v115 int32
	_ = v115
	var v117 int32
	_ = v117
	var v118 int32
	_ = v118
	var v120 int32
	_ = v120
	var v122 int32
	_ = v122
	v7 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v8 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	*(*int32)(unsafe.Add(mBase, uint32(v7)+4)) = v8
	v10 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	*(*int32)(unsafe.Add(mBase, uint32(v8))) = v10
	v12 = *(*int32)(unsafe.Add(mBase, uint32(l0)+56))
	if v12 == int32(0) {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	v86 = *(*int32)(unsafe.Add(mBase, uint32(l0)+44))
	if v86 != 0 {
		goto L20
	} else {
		goto L21
	}
L2:
	;
	v15 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	if v15 <= int32(0) {
		goto L1
	} else {
		goto L3
	}
L3:
	;
	v20 = int32(0)
	v21 = v15
	goto L4
L4:
	;
	v26 = v20 << (uint(int32(3)) % 32)
	v27 = *(*int32)(unsafe.Add(mBase, uint32(l0)+56))
	v28 = v26 + v27
	v29 = *(*int32)(unsafe.Add(mBase, uint32(v28)+4))
	if v29 != 0 {
		goto L6
	} else {
		goto L7
	}
L5:
	;
	goto L1
L6:
	;
	v31 = *(*int32)(unsafe.Add(mBase, _c_F_DestroyParallelContext[0]))
	v32 = *(*int32)(unsafe.Add(mBase, uint32(v28)))
	v33 = *(*int32)(unsafe.Add(mBase, uint32(v32)))
	v35 = *(*int32)(unsafe.Add(mBase, _c_F_DestroyParallelContext[1]))
	v39 = F_LWLockAcquire(m, v35+int32(_a_F_DestroyParallelContext_0), int32(0))
	mBase = m.M
	v40 = m.ExcPending
	if v40 != 0 {
		goto L9
	} else {
		goto L10
	}
L7:
	;
	v74 = v21
	goto L8
L8:
	;
	v78 = v20 + int32(1)
	if v78 < v74 {
		v20 = v78
		v21 = v74
		goto L4
	} else {
		goto L19
	}
L9:
	;
	return
L10:
	;
	v41 = *(*int64)(unsafe.Add(mBase, uint32(v32)+8))
	v44 = v31 + v33*int32(1480)
	v45 = *(*int64)(unsafe.Add(mBase, uint32(v44)+24))
	if v41 == v45 {
		goto L12
	} else {
		goto L13
	}
L11:
	;
	v64 = *(*int32)(unsafe.Add(mBase, uint32(l0)+56))
	v66 = *(*int32)(unsafe.Add(mBase, uint32(v64+v26)+4))
	F_shm_mq_detach(m, v66)
	mBase = m.M
	v68 = m.ExcPending
	if v68 != 0 {
		goto L9
	} else {
		goto L18
	}
L12:
	;
	v47 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(v44)+17)) = uint8(v47)
	v50 = *(*int32)(unsafe.Add(mBase, _c_F_DestroyParallelContext[1]))
	F_LWLockRelease(m, v50+int32(_a_F_DestroyParallelContext_0))
	mBase = m.M
	v54 = m.ExcPending
	if v54 != 0 {
		goto L9
	} else {
		goto L15
	}
L13:
	;
	goto L14
L14:
	;
	v59 = *(*int32)(unsafe.Add(mBase, _c_F_DestroyParallelContext[1]))
	F_LWLockRelease(m, v59+int32(_a_F_DestroyParallelContext_0))
	mBase = m.M
	v63 = m.ExcPending
	if v63 != 0 {
		goto L9
	} else {
		goto L17
	}
L15:
	;
	F_SendPostmasterSignal(m, int32(6))
	mBase = m.M
	v57 = m.ExcPending
	if v57 != 0 {
		goto L9
	} else {
		goto L16
	}
L16:
	;
	goto L11
L17:
	;
	goto L11
L18:
	;
	v69 = *(*int32)(unsafe.Add(mBase, uint32(l0)+56))
	*(*int32)(unsafe.Add(mBase, uint32(v69+v26)+4)) = int32(0)
	v73 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v74 = v73
	goto L8
L19:
	;
	goto L5
L20:
	;
	F_dsm_detach(m, v86)
	mBase = m.M
	v88 = m.ExcPending
	if v88 != 0 {
		goto L9
	} else {
		goto L23
	}
L21:
	;
	goto L22
L22:
	;
	v91 = *(*int32)(unsafe.Add(mBase, uint32(l0)+48))
	if v91 != 0 {
		goto L24
	} else {
		goto L25
	}
L23:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+44)) = int32(0)
	goto L22
L24:
	;
	F_pfree(m, v91)
	mBase = m.M
	v93 = m.ExcPending
	if v93 != 0 {
		goto L9
	} else {
		goto L27
	}
L25:
	;
	goto L26
L26:
	;
	v96 = int32(_a_F_DestroyParallelContext_1)
	v98 = *(*int32)(unsafe.Add(mBase, _c_F_DestroyParallelContext[2]))
	*(*int32)(unsafe.Add(mBase, _c_F_DestroyParallelContext[2])) = v98 + int32(1)
	F_WaitForParallelWorkersToExit(m, l0)
	mBase = m.M
	v103 = m.ExcPending
	if v103 != 0 {
		goto L9
	} else {
		goto L28
	}
L27:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+48)) = int32(0)
	goto L26
L28:
	;
	v104 = int32(_a_F_DestroyParallelContext_1)
	v106 = *(*int32)(unsafe.Add(mBase, _c_F_DestroyParallelContext[2]))
	*(*int32)(unsafe.Add(mBase, _c_F_DestroyParallelContext[2])) = v106 - int32(1)
	v110 = *(*int32)(unsafe.Add(mBase, uint32(l0)+56))
	if v110 != 0 {
		goto L29
	} else {
		goto L30
	}
L29:
	;
	F_pfree(m, v110)
	mBase = m.M
	v112 = m.ExcPending
	if v112 != 0 {
		goto L9
	} else {
		goto L32
	}
L30:
	;
	goto L31
L31:
	;
	v115 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
	F_pfree(m, v115)
	mBase = m.M
	v117 = m.ExcPending
	if v117 != 0 {
		goto L9
	} else {
		goto L33
	}
L32:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+56)) = int32(0)
	goto L31
L33:
	;
	v118 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
	F_pfree(m, v118)
	mBase = m.M
	v120 = m.ExcPending
	if v120 != 0 {
		goto L9
	} else {
		goto L34
	}
L34:
	;
	F_pfree(m, l0)
	mBase = m.M
	v122 = m.ExcPending
	if v122 != 0 {
		goto L9
	} else {
		goto L35
	}
L35:
	;
	return
}
func F_ExecParallelCleanup(m *base.Module, l0 int32) {
	mBase := m.M
	_ = mBase
	var v8 int32
	_ = v8
	var v9 int32
	_ = v9
	var v10 int32
	_ = v10
	var v11 int32
	_ = v11
	var v12 int32
	_ = v12
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
	var v23 int32
	_ = v23
	var v24 int32
	_ = v24
	var v26 int32
	_ = v26
	var v27 int32
	_ = v27
	var v29 int32
	_ = v29
	var v30 int32
	_ = v30
	var v38 int32
	_ = v38
	var v45 int32
	_ = v45
	var v46 int32
	_ = v46
	var v47 int32
	_ = v47
	var v50 int64
	_ = v50
	var v51 int64
	_ = v51
	var v54 int64
	_ = v54
	var v55 int64
	_ = v55
	var v58 int64
	_ = v58
	var v59 int64
	_ = v59
	var v62 int64
	_ = v62
	var v63 int64
	_ = v63
	var v66 int64
	_ = v66
	var v67 int64
	_ = v67
	var v71 int32
	_ = v71
	var v72 int32
	_ = v72
	var v79 int32
	_ = v79
	var v82 int32
	_ = v82
	var v83 int32
	_ = v83
	var v84 int32
	_ = v84
	var v85 int32
	_ = v85
	var v87 int32
	_ = v87
	var v88 int32
	_ = v88
	var v89 int32
	_ = v89
	var v101 int32
	_ = v101
	var v102 int32
	_ = v102
	var v104 int32
	_ = v104
	var v107 int32
	_ = v107
	var v109 int32
	_ = v109
	var v112 int32
	_ = v112
	var v114 int32
	_ = v114
	var v118 int32
	_ = v118
	v8 = *(*int32)(unsafe.Add(mBase, uint32(l0)+16))
	if v8 != 0 {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	v9 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v10 = F_ExecParallelRetrieveInstrumentation(m, v9, v8)
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
	v12 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	if v12 == int32(0) {
		goto L6
	} else {
		goto L7
	}
L4:
	;
	return
L5:
	;
	goto L3
L6:
	;
	v101 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
	if v101 != 0 {
		goto L22
	} else {
		goto L23
	}
L7:
	;
	v15 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v16 = *(*int32)(unsafe.Add(mBase, uint32(v15)+8))
	v17 = *(*int32)(unsafe.Add(mBase, uint32(v16)+184))
	if v17 == int32(0) {
		goto L8
	} else {
		goto L9
	}
L8:
	;
	v20 = *(*int32)(unsafe.Add(mBase, uint32(v16)+100))
	v22 = F_MemoryContextAllocZero(m, v20, int32(48))
	mBase = m.M
	v23 = m.ExcPending
	if v23 != 0 {
		goto L4
	} else {
		goto L11
	}
L9:
	;
	v29 = v17
	goto L10
L10:
	;
	v30 = *(*int32)(unsafe.Add(mBase, uint32(v12)))
	if int32(0) < v30 {
		goto L12
	} else {
		goto L13
	}
L11:
	;
	v24 = *(*int32)(unsafe.Add(mBase, uint32(v15)+8))
	*(*int32)(unsafe.Add(mBase, uint32(v24)+184)) = v22
	v26 = *(*int32)(unsafe.Add(mBase, uint32(v15)+8))
	v27 = *(*int32)(unsafe.Add(mBase, uint32(v26)+184))
	v29 = v27
	goto L10
L12:
	;
	v38 = int32(0)
	goto L15
L13:
	;
	v79 = v30
	goto L14
L14:
	;
	v82 = F_mul_size(m, v79, int32(48))
	mBase = m.M
	v83 = m.ExcPending
	if v83 != 0 {
		goto L4
	} else {
		goto L19
	}
L15:
	;
	v45 = v12 + int32(8) + v38*int32(48)
	v46 = *(*int32)(unsafe.Add(mBase, uint32(v29)))
	v47 = *(*int32)(unsafe.Add(mBase, uint32(v45)))
	*(*int32)(unsafe.Add(mBase, uint32(v29))) = v46 + v47
	v50 = *(*int64)(unsafe.Add(mBase, uint32(v29)+8))
	v51 = *(*int64)(unsafe.Add(mBase, uint32(v45)+8))
	*(*int64)(unsafe.Add(mBase, uint32(v29)+8)) = v50 + v51
	v54 = *(*int64)(unsafe.Add(mBase, uint32(v29)+16))
	v55 = *(*int64)(unsafe.Add(mBase, uint32(v45)+16))
	*(*int64)(unsafe.Add(mBase, uint32(v29)+16)) = v54 + v55
	v58 = *(*int64)(unsafe.Add(mBase, uint32(v29)+24))
	v59 = *(*int64)(unsafe.Add(mBase, uint32(v45)+24))
	*(*int64)(unsafe.Add(mBase, uint32(v29)+24)) = v58 + v59
	v62 = *(*int64)(unsafe.Add(mBase, uint32(v29)+32))
	v63 = *(*int64)(unsafe.Add(mBase, uint32(v45)+32))
	*(*int64)(unsafe.Add(mBase, uint32(v29)+32)) = v62 + v63
	v66 = *(*int64)(unsafe.Add(mBase, uint32(v29)+40))
	v67 = *(*int64)(unsafe.Add(mBase, uint32(v45)+40))
	*(*int64)(unsafe.Add(mBase, uint32(v29)+40)) = v66 + v67
	goto L17
L16:
	;
	v79 = v72
	goto L14
L17:
	;
	v71 = v38 + int32(1)
	v72 = *(*int32)(unsafe.Add(mBase, uint32(v12)))
	if v71 < v72 {
		v38 = v71
		goto L15
	} else {
		goto L18
	}
L18:
	;
	goto L16
L19:
	;
	v84 = *(*int32)(unsafe.Add(mBase, uint32(v15)+8))
	v85 = *(*int32)(unsafe.Add(mBase, uint32(v84)+100))
	v87 = v82 + int32(8)
	v88 = F_MemoryContextAlloc(m, v85, v87)
	mBase = m.M
	v89 = m.ExcPending
	if v89 != 0 {
		goto L4
	} else {
		goto L20
	}
L20:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v15)+28)) = v88
	if v87 == int32(0) {
		goto L6
	} else {
		goto L21
	}
L21:
	;
	base.MemoryCopy(m, v88, v12, v87)
	goto L6
L22:
	;
	v102 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
	F_dsa_free(m, v102, v101)
	mBase = m.M
	v104 = m.ExcPending
	if v104 != 0 {
		goto L4
	} else {
		goto L25
	}
L23:
	;
	goto L24
L24:
	;
	v107 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
	if v107 != 0 {
		goto L26
	} else {
		goto L27
	}
L25:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+28)) = int32(0)
	goto L24
L26:
	;
	F_dsa_detach(m, v107)
	mBase = m.M
	v109 = m.ExcPending
	if v109 != 0 {
		goto L4
	} else {
		goto L29
	}
L27:
	;
	goto L28
L28:
	;
	v112 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	if v112 != 0 {
		goto L30
	} else {
		goto L31
	}
L29:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+24)) = int32(0)
	goto L28
L30:
	;
	F_DestroyParallelContext(m, v112)
	mBase = m.M
	v114 = m.ExcPending
	if v114 != 0 {
		goto L4
	} else {
		goto L33
	}
L31:
	;
	goto L32
L32:
	;
	F_pfree(m, l0)
	mBase = m.M
	v118 = m.ExcPending
	if v118 != 0 {
		goto L4
	} else {
		goto L34
	}
L33:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = int32(0)
	goto L32
L34:
	;
	return
}
func F_ExecParallelHashCloseBatchAccessors(m *base.Module, l0 int32) {
	mBase := m.M
	_ = mBase
	var v2 int32
	_ = v2
	var v4 int32
	_ = v4
	var v9 int32
	_ = v9
	var v11 int32
	_ = v11
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
	var v21 int32
	_ = v21
	var v22 int32
	_ = v22
	var v24 int32
	_ = v24
	var v26 int32
	_ = v26
	var v27 int32
	_ = v27
	var v29 int32
	_ = v29
	var v31 int32
	_ = v31
	var v33 int32
	_ = v33
	var v34 int32
	_ = v34
	var v39 int32
	_ = v39
	var v41 int32
	_ = v41
	v2 = int32(0)
	v4 = *(*int32)(unsafe.Add(mBase, uint32(l0)+44))
	if v2 < v4 {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	v9 = v2
	goto L4
L2:
	;
	goto L3
L3:
	;
	v39 = *(*int32)(unsafe.Add(mBase, uint32(l0)+144))
	F_pfree(m, v39)
	mBase = m.M
	v41 = m.ExcPending
	if v41 != 0 {
		goto L6
	} else {
		goto L12
	}
L4:
	;
	v11 = v9 * int32(36)
	v12 = *(*int32)(unsafe.Add(mBase, uint32(l0)+144))
	v14 = *(*int32)(unsafe.Add(mBase, uint32(v11+v12)+28))
	F_sts_end_write(m, v14)
	mBase = m.M
	v16 = m.ExcPending
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
	return
L7:
	;
	v17 = *(*int32)(unsafe.Add(mBase, uint32(l0)+144))
	v19 = *(*int32)(unsafe.Add(mBase, uint32(v17+v11)+32))
	F_sts_end_write(m, v19)
	mBase = m.M
	v21 = m.ExcPending
	if v21 != 0 {
		goto L6
	} else {
		goto L8
	}
L8:
	;
	v22 = *(*int32)(unsafe.Add(mBase, uint32(l0)+144))
	v24 = *(*int32)(unsafe.Add(mBase, uint32(v22+v11)+28))
	F_sts_end_parallel_scan(m, v24)
	mBase = m.M
	v26 = m.ExcPending
	if v26 != 0 {
		goto L6
	} else {
		goto L9
	}
L9:
	;
	v27 = *(*int32)(unsafe.Add(mBase, uint32(l0)+144))
	v29 = *(*int32)(unsafe.Add(mBase, uint32(v27+v11)+32))
	F_sts_end_parallel_scan(m, v29)
	mBase = m.M
	v31 = m.ExcPending
	if v31 != 0 {
		goto L6
	} else {
		goto L10
	}
L10:
	;
	v33 = v9 + int32(1)
	v34 = *(*int32)(unsafe.Add(mBase, uint32(l0)+44))
	if v33 < v34 {
		v9 = v33
		goto L4
	} else {
		goto L11
	}
L11:
	;
	goto L5
L12:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+144)) = int32(0)
	return
}
func F_ExecParallelHashEnsureBatchAccessors(m *base.Module, l0 int32) {
	mBase := m.M
	_ = mBase
	var v9 int32
	_ = v9
	var v10 int32
	_ = v10
	var v11 int32
	_ = v11
	var v12 int32
	_ = v12
	var v15 int32
	_ = v15
	var v16 int32
	_ = v16
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
	var v36 int32
	_ = v36
	var v41 int32
	_ = v41
	var v45 int32
	_ = v45
	var v46 int32
	_ = v46
	var v47 int32
	_ = v47
	var v54 int32
	_ = v54
	var v55 int32
	_ = v55
	var v59 int32
	_ = v59
	var v65 int32
	_ = v65
	var v68 int32
	_ = v68
	var v71 int32
	_ = v71
	var v73 int32
	_ = v73
	var v76 int32
	_ = v76
	var v77 int32
	_ = v77
	var v79 int32
	_ = v79
	var v90 int32
	_ = v90
	var v93 int32
	_ = v93
	var v94 int32
	_ = v94
	var v97 int32
	_ = v97
	var v98 int32
	_ = v98
	v9 = *(*int32)(unsafe.Add(mBase, uint32(l0)+140))
	v10 = *(*int32)(unsafe.Add(mBase, uint32(l0)+144))
	if v10 != 0 {
		goto L2
	} else {
		goto L3
	}
L1:
	;
	return
L2:
	;
	v11 = *(*int32)(unsafe.Add(mBase, uint32(l0)+44))
	v12 = *(*int32)(unsafe.Add(mBase, uint32(v9)+8))
	if v11 == v12 {
		goto L1
	} else {
		goto L5
	}
L3:
	;
	goto L4
L4:
	;
	v16 = int32(_a_F_ExecParallelHashEnsureBatchAccessors_0)
	v17 = *(*int32)(unsafe.Add(mBase, _c_F_ExecParallelHashEnsureBatchAccessors[0]))
	v19 = *(*int32)(unsafe.Add(mBase, uint32(l0)+124))
	*(*int32)(unsafe.Add(mBase, _c_F_ExecParallelHashEnsureBatchAccessors[0])) = v19
	v21 = *(*int32)(unsafe.Add(mBase, uint32(v9)+8))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+44)) = v21
	v25 = F_palloc0(m, v21*int32(36))
	mBase = m.M
	v26 = m.ExcPending
	if v26 != 0 {
		goto L6
	} else {
		goto L8
	}
L5:
	;
	F_ExecParallelHashCloseBatchAccessors(m, l0)
	mBase = m.M
	v15 = m.ExcPending
	if v15 != 0 {
		goto L6
	} else {
		goto L7
	}
L6:
	;
	return
L7:
	;
	goto L4
L8:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+144)) = v25
	v28 = *(*int32)(unsafe.Add(mBase, uint32(l0)+136))
	v29 = *(*int32)(unsafe.Add(mBase, uint32(v9)))
	v30 = F_dsa_get_address(m, v28, v29)
	mBase = m.M
	v31 = m.ExcPending
	if v31 != 0 {
		goto L6
	} else {
		goto L9
	}
L9:
	;
	v32 = *(*int32)(unsafe.Add(mBase, uint32(l0)+44))
	if int32(0) < v32 {
		goto L10
	} else {
		goto L11
	}
L10:
	;
	v36 = v9 + int32(168)
	v41 = int32(0)
	goto L13
L11:
	;
	goto L12
L12:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_ExecParallelHashEnsureBatchAccessors[0])) = v17
	goto L1
L13:
	;
	v45 = *(*int32)(unsafe.Add(mBase, uint32(l0)+144))
	v46 = *(*int32)(unsafe.Add(mBase, uint32(l0)+140))
	v47 = *(*int32)(unsafe.Add(mBase, uint32(v46)+28))
	goto L15
L14:
	;
	goto L12
L15:
	;
	v54 = v45 + v41*int32(36)
	v55 = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v54)+4)) = v55
	*(*uint16)(unsafe.Add(mBase, uint32(v54)+25)) = uint16(v55)
	v59 = int32(1)
	v65 = int32(-64)
	v68 = v30 + (((v47*int32(28)+int32(76))<<(uint(v59)%32)+int32(14))&int32(-16)-v65)*v41
	*(*int32)(unsafe.Add(mBase, uint32(v54))) = v68
	v71 = v68 - v65
	v73 = *(*int32)(unsafe.Add(mBase, _c_F_ExecParallelHashEnsureBatchAccessors[1]))
	v76 = F_sts_attach(m, v71, v73+v59, v36)
	mBase = m.M
	v77 = m.ExcPending
	if v77 != 0 {
		goto L6
	} else {
		goto L16
	}
L16:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v54)+28)) = v76
	v79 = *(*int32)(unsafe.Add(mBase, uint32(v9)+28))
	goto L17
L17:
	;
	v90 = *(*int32)(unsafe.Add(mBase, _c_F_ExecParallelHashEnsureBatchAccessors[1]))
	v93 = F_sts_attach(m, v71+(v79*int32(28)+int32(83))&int32(-8), v90+int32(1), v36)
	mBase = m.M
	v94 = m.ExcPending
	if v94 != 0 {
		goto L6
	} else {
		goto L18
	}
L18:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v54)+32)) = v93
	v97 = v41 + int32(1)
	v98 = *(*int32)(unsafe.Add(mBase, uint32(l0)+44))
	if v97 < v98 {
		v41 = v97
		goto L13
	} else {
		goto L19
	}
L19:
	;
	goto L14
}
func F_ExecParallelHashIncreaseNumBatches(m *base.Module, l0 int32) {
	mBase := m.M
	_ = mBase
	var v17 int32
	_ = v17
	var v19 int32
	_ = v19
	var v21 int32
	_ = v21
	var v23 int32
	_ = v23
	var v24 int32
	_ = v24
	var v26 int32
	_ = v26
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
	var v39 int32
	_ = v39
	var v40 int32
	_ = v40
	var v44 float64
	_ = v44
	var v46 int32
	_ = v46
	var v50 float64
	_ = v50
	var v51 float64
	_ = v51
	var v54 float64
	_ = v54
	var v57 int32
	_ = v57
	var v59 int32
	_ = v59
	var v61 int32
	_ = v61
	var v68 int32
	_ = v68
	var v73 int32
	_ = v73
	var v75 int32
	_ = v75
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
	var v84 int32
	_ = v84
	var v85 int32
	_ = v85
	var v89 float64
	_ = v89
	var v93 float64
	_ = v93
	var v94 float64
	_ = v94
	var v97 float64
	_ = v97
	var v98 int32
	_ = v98
	var v101 int32
	_ = v101
	var v108 int32
	_ = v108
	var v112 int32
	_ = v112
	var v113 int32
	_ = v113
	var v114 int32
	_ = v114
	var v115 int32
	_ = v115
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
	var v132 int32
	_ = v132
	var v136 int32
	_ = v136
	var v149 int32
	_ = v149
	var v150 int32
	_ = v150
	var v166 int32
	_ = v166
	var v167 int32
	_ = v167
	var v169 int32
	_ = v169
	var v174 int32
	_ = v174
	var v191 int32
	_ = v191
	var v192 int32
	_ = v192
	var v211 int32
	_ = v211
	var v214 int32
	_ = v214
	var v233 int32
	_ = v233
	var v234 int32
	_ = v234
	var v235 int32
	_ = v235
	var v237 int32
	_ = v237
	var v238 int32
	_ = v238
	var v239 int32
	_ = v239
	var v240 int32
	_ = v240
	var v245 int32
	_ = v245
	var v266 int32
	_ = v266
	var v267 int32
	_ = v267
	var v285 int32
	_ = v285
	var v290 int32
	_ = v290
	var v308 int32
	_ = v308
	var v309 int32
	_ = v309
	var v327 int32
	_ = v327
	var v330 int32
	_ = v330
	var v331 int32
	_ = v331
	var v332 int32
	_ = v332
	var v333 int32
	_ = v333
	var v334 int32
	_ = v334
	var v335 int32
	_ = v335
	var v337 int32
	_ = v337
	var v338 int32
	_ = v338
	var v341 int32
	_ = v341
	var v344 int32
	_ = v344
	var v352 int32
	_ = v352
	var v353 int32
	_ = v353
	var v358 int32
	_ = v358
	var v361 int32
	_ = v361
	var v363 int32
	_ = v363
	var v365 int32
	_ = v365
	var v366 int32
	_ = v366
	var v367 int32
	_ = v367
	var v371 int32
	_ = v371
	var v372 int32
	_ = v372
	var v378 int32
	_ = v378
	var v386 int32
	_ = v386
	var v387 int32
	_ = v387
	var v388 int32
	_ = v388
	var v389 int32
	_ = v389
	var v392 int32
	_ = v392
	var v393 int32
	_ = v393
	var v400 int32
	_ = v400
	var v413 int32
	_ = v413
	var v414 int32
	_ = v414
	var v415 int32
	_ = v415
	var v417 int32
	_ = v417
	var v419 int32
	_ = v419
	var v420 int32
	_ = v420
	var v425 int32
	_ = v425
	var v427 int32
	_ = v427
	var v432 int32
	_ = v432
	var v433 int32
	_ = v433
	var v437 int32
	_ = v437
	var v438 int32
	_ = v438
	var v439 int32
	_ = v439
	var v441 int32
	_ = v441
	var v445 int32
	_ = v445
	var v446 int32
	_ = v446
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
	var v459 int32
	_ = v459
	var v474 int32
	_ = v474
	var v476 int32
	_ = v476
	var v477 int32
	_ = v477
	var v478 int32
	_ = v478
	var v483 int32
	_ = v483
	var v484 int32
	_ = v484
	var v485 int32
	_ = v485
	var v486 int32
	_ = v486
	var v487 int32
	_ = v487
	var v494 int32
	_ = v494
	var v496 int32
	_ = v496
	var v498 int32
	_ = v498
	var v508 int32
	_ = v508
	var v515 int32
	_ = v515
	var v516 int32
	_ = v516
	var v517 int32
	_ = v517
	var v520 int32
	_ = v520
	var v523 int32
	_ = v523
	var v524 int32
	_ = v524
	var v528 int32
	_ = v528
	var v533 int32
	_ = v533
	var v534 int32
	_ = v534
	var v552 int32
	_ = v552
	var v554 int32
	_ = v554
	var v556 int32
	_ = v556
	var v558 int32
	_ = v558
	var v559 int32
	_ = v559
	var v561 int32
	_ = v561
	var v563 int32
	_ = v563
	var v564 int32
	_ = v564
	var v567 int32
	_ = v567
	var v570 int32
	_ = v570
	var v585 int32
	_ = v585
	var v586 int32
	_ = v586
	var v587 int32
	_ = v587
	var v588 int32
	_ = v588
	var v589 int32
	_ = v589
	var v590 int32
	_ = v590
	var v591 int32
	_ = v591
	var v594 int32
	_ = v594
	var v595 int32
	_ = v595
	var v602 int32
	_ = v602
	var v620 int32
	_ = v620
	var v621 int32
	_ = v621
	var v626 int32
	_ = v626
	var v632 int32
	_ = v632
	var v639 int32
	_ = v639
	var v642 int32
	_ = v642
	var v643 int32
	_ = v643
	var v646 int32
	_ = v646
	var v652 int32
	_ = v652
	var v667 int32
	_ = v667
	var v668 int32
	_ = v668
	var v670 int32
	_ = v670
	var v671 int32
	_ = v671
	var v674 int32
	_ = v674
	var v675 int32
	_ = v675
	var v677 int32
	_ = v677
	var v692 int32
	_ = v692
	var v697 int32
	_ = v697
	var v702 int32
	_ = v702
	var v703 int32
	_ = v703
	var v707 int32
	_ = v707
	var v708 int32
	_ = v708
	var v709 int32
	_ = v709
	var v710 int32
	_ = v710
	var v711 int32
	_ = v711
	var v712 int32
	_ = v712
	var v715 int32
	_ = v715
	var v716 int32
	_ = v716
	var v717 int32
	_ = v717
	var v718 int32
	_ = v718
	var v721 int32
	_ = v721
	var v724 int32
	_ = v724
	var v725 int32
	_ = v725
	var v729 int32
	_ = v729
	var v731 int32
	_ = v731
	var v735 int32
	_ = v735
	var v737 int32
	_ = v737
	var v739 int32
	_ = v739
	var v740 int32
	_ = v740
	var v743 int32
	_ = v743
	var v744 int32
	_ = v744
	var v761 int32
	_ = v761
	var v763 int32
	_ = v763
	var v765 int32
	_ = v765
	var v784 int32
	_ = v784
	var v786 int32
	_ = v786
	var v788 int32
	_ = v788
	var v789 int32
	_ = v789
	var v807 int32
	_ = v807
	var v808 int32
	_ = v808
	var v812 int32
	_ = v812
	var v813 int32
	_ = v813
	var v816 int32
	_ = v816
	var v817 int32
	_ = v817
	var v818 int32
	_ = v818
	var v819 int32
	_ = v819
	var v820 int32
	_ = v820
	var v821 int32
	_ = v821
	var v823 int32
	_ = v823
	var v824 int32
	_ = v824
	var v827 int32
	_ = v827
	var v830 int32
	_ = v830
	var v838 int32
	_ = v838
	var v839 int32
	_ = v839
	var v844 int32
	_ = v844
	var v847 int32
	_ = v847
	var v848 int32
	_ = v848
	var v849 int32
	_ = v849
	var v850 int32
	_ = v850
	var v851 int32
	_ = v851
	var v854 int32
	_ = v854
	var v858 int32
	_ = v858
	var v861 int32
	_ = v861
	var v862 int32
	_ = v862
	var v873 int32
	_ = v873
	var v877 int32
	_ = v877
	var v878 int32
	_ = v878
	var v881 int32
	_ = v881
	var v882 int32
	_ = v882
	var v885 int32
	_ = v885
	var v886 int32
	_ = v886
	var v887 int32
	_ = v887
	var v888 int32
	_ = v888
	var v901 int32
	_ = v901
	var v904 int32
	_ = v904
	var v907 int32
	_ = v907
	var v908 int32
	_ = v908
	var v910 int32
	_ = v910
	var v911 int32
	_ = v911
	var v915 int32
	_ = v915
	var v916 int32
	_ = v916
	var v919 int32
	_ = v919
	var v921 int32
	_ = v921
	var v922 int32
	_ = v922
	var v926 int32
	_ = v926
	var v936 int32
	_ = v936
	var v940 int32
	_ = v940
	var v954 int32
	_ = v954
	var v955 int32
	_ = v955
	var v957 int32
	_ = v957
	var v977 int32
	_ = v977
	var v978 int32
	_ = v978
	v17 = m.G0
	v19 = v17 - int32(16)
	m.G0 = v19
	v21 = *(*int32)(unsafe.Add(mBase, uint32(l0)+140))
	v23 = v21 + int32(92)
	v24 = *(*int32)(unsafe.Add(mBase, uint32(v23)+4))
	v26 = base.I32_rem_s(v24, int32(5))
	switch v26 {
	case 0:
		goto L6
	case 1:
		goto L5
	case 2:
		goto L4
	case 3:
		goto L3
	case 4:
		goto L2
	default:
		goto L1
	}
L1:
	;
	m.G0 = v19 + int32(16)
	return
L2:
	;
	v977 = F_BarrierArriveAndWait(m, v23, int32(134217751))
	mBase = m.M
	v978 = m.ExcPending
	if v978 != 0 {
		goto L7
	} else {
		goto L178
	}
L3:
	;
	v807 = F_BarrierArriveAndWait(m, v23, int32(134217749))
	mBase = m.M
	v808 = m.ExcPending
	if v808 != 0 {
		goto L7
	} else {
		goto L145
	}
L4:
	;
	F_ExecParallelHashEnsureBatchAccessors(m, l0)
	mBase = m.M
	v327 = m.ExcPending
	if v327 != 0 {
		goto L7
	} else {
		goto L59
	}
L5:
	;
	v308 = F_BarrierArriveAndWait(m, v23, int32(134217752))
	mBase = m.M
	v309 = m.ExcPending
	if v309 != 0 {
		goto L7
	} else {
		goto L58
	}
L6:
	;
	v28 = F_BarrierArriveAndWait(m, v23, int32(134217750))
	mBase = m.M
	v29 = m.ExcPending
	if v29 != 0 {
		goto L7
	} else {
		goto L8
	}
L7:
	;
	return
L8:
	;
	if v28 != 0 {
		goto L9
	} else {
		goto L10
	}
L9:
	;
	v30 = *(*int32)(unsafe.Add(mBase, uint32(l0)+144))
	v31 = *(*int32)(unsafe.Add(mBase, uint32(v30)))
	v32 = *(*int32)(unsafe.Add(mBase, uint32(v21)))
	*(*int32)(unsafe.Add(mBase, uint32(v21)+4)) = v32
	v34 = *(*int32)(unsafe.Add(mBase, uint32(l0)+44))
	*(*int32)(unsafe.Add(mBase, uint32(v21)+12)) = v34
	*(*int32)(unsafe.Add(mBase, uint32(v21))) = int32(0)
	F_ExecParallelHashCloseBatchAccessors(m, l0)
	mBase = m.M
	v39 = m.ExcPending
	if v39 != 0 {
		goto L7
	} else {
		goto L12
	}
L10:
	;
	goto L11
L11:
	;
	F_ExecParallelHashCloseBatchAccessors(m, l0)
	mBase = m.M
	v290 = m.ExcPending
	if v290 != 0 {
		goto L7
	} else {
		goto L57
	}
L12:
	;
	v40 = *(*int32)(unsafe.Add(mBase, uint32(l0)+44))
	if v40 == int32(1) {
		goto L14
	} else {
		goto L15
	}
L13:
	;
	F_ExecParallelHashJoinSetUpBatches(m, l0, v73)
	mBase = m.M
	v75 = m.ExcPending
	if v75 != 0 {
		goto L7
	} else {
		goto L23
	}
L14:
	;
	v44 = *(*float64)(unsafe.Add(mBase, _c_F_ExecParallelHashIncreaseNumBatches[0]))
	v46 = *(*int32)(unsafe.Add(mBase, _c_F_ExecParallelHashIncreaseNumBatches[1]))
	v50 = base.F64_mul(base.F64_mul(v44, base.F64_convert_i32_s(v46)), float64(1024))
	v51 = float64(4.294967295e+09)
	if base.F64_lt(v50, v51) != 0 {
		goto L17
	} else {
		goto L18
	}
L15:
	;
	goto L16
L16:
	;
	v73 = v40 << (uint(int32(1)) % 32)
	goto L13
L17:
	;
	v54 = v50
	goto L19
L18:
	;
	v54 = v51
	goto L19
L19:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v21)+32)) = base.I32_trunc_sat_f64_u(v54)
	v57 = int32(1)
	v59 = *(*int32)(unsafe.Add(mBase, uint32(v21)+28))
	v61 = v59 << (uint(v57) % 32)
	if v61&(v61-v57) != 0 {
		goto L20
	} else {
		goto L21
	}
L20:
	;
	v68 = v57 << (uint(int32(32)-base.I32_clz(v61)) % 32)
	goto L22
L21:
	;
	v68 = v61
	goto L22
L22:
	;
	v73 = v68
	goto L13
L23:
	;
	v76 = *(*int32)(unsafe.Add(mBase, uint32(v21)+12))
	if v76 == int32(1) {
		goto L25
	} else {
		goto L26
	}
L24:
	;
	v285 = *(*int32)(unsafe.Add(mBase, uint32(v31)+40))
	*(*int32)(unsafe.Add(mBase, uint32(v21)+20)) = int32(3)
	*(*int32)(unsafe.Add(mBase, uint32(v21)+24)) = v285
	goto L5
L25:
	;
	v79 = *(*int32)(unsafe.Add(mBase, uint32(v31)+52))
	v80 = *(*int32)(unsafe.Add(mBase, uint32(l0)+136))
	v81 = *(*int32)(unsafe.Add(mBase, uint32(v31)))
	F_dsa_free(m, v80, v81)
	mBase = m.M
	v83 = m.ExcPending
	if v83 != 0 {
		goto L7
	} else {
		goto L28
	}
L26:
	;
	goto L27
L27:
	;
	v233 = *(*int32)(unsafe.Add(mBase, uint32(l0)+144))
	v234 = *(*int32)(unsafe.Add(mBase, uint32(v233)))
	v235 = *(*int32)(unsafe.Add(mBase, uint32(v31)))
	*(*int32)(unsafe.Add(mBase, uint32(v234))) = v235
	v237 = *(*int32)(unsafe.Add(mBase, uint32(l0)+136))
	v238 = F_dsa_get_address(m, v237, v235)
	mBase = m.M
	v239 = m.ExcPending
	if v239 != 0 {
		goto L7
	} else {
		goto L52
	}
L28:
	;
	v84 = int32(0)
	v85 = *(*int32)(unsafe.Add(mBase, uint32(l0)+136))
	v89 = base.F64_convert_i32_u(v79)
	v93 = base.F64_ceil(base.F64_div(base.F64_add(v89, v89), base.F64_convert_i32_s(v73)))
	v94 = float64(1.34217728e+08)
	if base.F64_lt(v93, v94) != 0 {
		goto L29
	} else {
		goto L30
	}
L29:
	;
	v97 = v93
	goto L31
L30:
	;
	v97 = v94
	goto L31
L31:
	;
	v98 = base.I32_trunc_sat_f64_s(v97)
	if v98 <= int32(1024) {
		goto L32
	} else {
		goto L33
	}
L32:
	;
	v101 = int32(1024)
	goto L34
L33:
	;
	v101 = v98
	goto L34
L34:
	;
	if v101&(v101-int32(1)) != 0 {
		goto L35
	} else {
		goto L36
	}
L35:
	;
	v108 = int32(1) << (uint(int32(32)-base.I32_clz(v101)) % 32)
	goto L37
L36:
	;
	v108 = v101
	goto L37
L37:
	;
	v112 = F_dsa_allocate_extended(m, v85, v108<<(uint(int32(2))%32), int32(0))
	mBase = m.M
	v113 = m.ExcPending
	if v113 != 0 {
		goto L7
	} else {
		goto L38
	}
L38:
	;
	v114 = *(*int32)(unsafe.Add(mBase, uint32(l0)+144))
	v115 = *(*int32)(unsafe.Add(mBase, uint32(v114)))
	*(*int32)(unsafe.Add(mBase, uint32(v115))) = v112
	v117 = *(*int32)(unsafe.Add(mBase, uint32(l0)+136))
	v118 = *(*int32)(unsafe.Add(mBase, uint32(l0)+144))
	v119 = *(*int32)(unsafe.Add(mBase, uint32(v118)))
	v120 = *(*int32)(unsafe.Add(mBase, uint32(v119)))
	v121 = F_dsa_get_address(m, v117, v120)
	mBase = m.M
	v122 = m.ExcPending
	if v122 != 0 {
		goto L7
	} else {
		goto L39
	}
L39:
	;
	if v108 <= int32(0) {
		goto L40
	} else {
		goto L41
	}
L40:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v21)+16)) = v108
	goto L24
L41:
	;
	v126 = v108 & int32(7)
	if base.Ui32(int32(8)) <= base.Ui32(v108) {
		goto L42
	} else {
		goto L43
	}
L42:
	;
	v132 = v84
	v136 = int32(0)
	goto L45
L43:
	;
	v174 = v84
	goto L44
L44:
	;
	v191 = v174
	v192 = int32(0)
	goto L49
L45:
	;
	v149 = v121 + v132<<(uint(int32(2))%32)
	v150 = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v149))) = v150
	*(*int32)(unsafe.Add(mBase, uint32(v149)+4)) = v150
	*(*int32)(unsafe.Add(mBase, uint32(v149)+8)) = v150
	*(*int32)(unsafe.Add(mBase, uint32(v149)+12)) = v150
	*(*int32)(unsafe.Add(mBase, uint32(v149)+16)) = v150
	*(*int32)(unsafe.Add(mBase, uint32(v149)+20)) = v150
	*(*int32)(unsafe.Add(mBase, uint32(v149)+24)) = v150
	*(*int32)(unsafe.Add(mBase, uint32(v149)+28)) = v150
	v166 = int32(8)
	v167 = v132 + v166
	v169 = v136 + v166
	if v169 != v108&int32(-8) {
		v132 = v167
		v136 = v169
		goto L45
	} else {
		goto L47
	}
L46:
	;
	if v126 == int32(0) {
		goto L40
	} else {
		goto L48
	}
L47:
	;
	goto L46
L48:
	;
	v174 = v167
	goto L44
L49:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v121+v191<<(uint(int32(2))%32)))) = int32(0)
	v211 = int32(1)
	v214 = v192 + v211
	if v214 != v126 {
		v191 = v191 + v211
		v192 = v214
		goto L49
	} else {
		goto L51
	}
L50:
	;
	goto L40
L51:
	;
	goto L50
L52:
	;
	v240 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	if v240 <= int32(0) {
		goto L24
	} else {
		goto L53
	}
L53:
	;
	v245 = int32(0)
	goto L54
L54:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v238+v245<<(uint(int32(2))%32)))) = int32(0)
	v266 = v245 + int32(1)
	v267 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	if v266 < v267 {
		v245 = v266
		goto L54
	} else {
		goto L56
	}
L55:
	;
	goto L24
L56:
	;
	goto L55
L57:
	;
	goto L5
L58:
	;
	goto L4
L59:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+48)) = int32(0)
	v330 = *(*int32)(unsafe.Add(mBase, uint32(l0)+136))
	v331 = *(*int32)(unsafe.Add(mBase, uint32(l0)+144))
	v332 = *(*int32)(unsafe.Add(mBase, uint32(v331)))
	v333 = *(*int32)(unsafe.Add(mBase, uint32(v332)))
	v334 = F_dsa_get_address(m, v330, v333)
	mBase = m.M
	v335 = m.ExcPending
	if v335 != 0 {
		goto L7
	} else {
		goto L60
	}
L60:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+20)) = v334
	v337 = *(*int32)(unsafe.Add(mBase, uint32(l0)+140))
	v338 = *(*int32)(unsafe.Add(mBase, uint32(v337)+16))
	*(*int32)(unsafe.Add(mBase, uint32(l0))) = v338
	v341 = int32(1073741823)
	if v341 <= v338 {
		goto L62
	} else {
		goto L63
	}
L61:
	;
	v353 = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(l0)+148)) = v353
	*(*int32)(unsafe.Add(mBase, uint32(l0)+132)) = v353
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v352
	v358 = *(*int32)(unsafe.Add(mBase, uint32(l0)+144))
	*(*uint8)(unsafe.Add(mBase, uint32(v358)+24)) = uint8(v353)
	v361 = *(*int32)(unsafe.Add(mBase, uint32(l0)+140))
	v363 = v361 + int32(40)
	v365 = F_LWLockAcquire(m, v363, v353)
	mBase = m.M
	v366 = m.ExcPending
	if v366 != 0 {
		goto L7
	} else {
		goto L68
	}
L62:
	;
	v344 = v341
	goto L64
L63:
	;
	v344 = v338
	goto L64
L64:
	;
	if base.Ui32(int32(2)) <= base.Ui32(v344) {
		goto L65
	} else {
		goto L66
	}
L65:
	;
	v352 = int32(32) - base.I32_clz(v344-int32(1))
	goto L67
L66:
	;
	v352 = int32(0)
	goto L67
L67:
	;
	goto L61
L68:
	;
	v367 = *(*int32)(unsafe.Add(mBase, uint32(v361)+24))
	if v367 != 0 {
		goto L69
	} else {
		goto L70
	}
L69:
	;
	v371 = v361 + int32(24)
	v372 = v363
	v378 = v367
	goto L72
L70:
	;
	v570 = v363
	goto L71
L71:
	;
	F_LWLockRelease(m, v570)
	mBase = m.M
	v585 = m.ExcPending
	if v585 != 0 {
		goto L7
	} else {
		goto L110
	}
L72:
	;
	v386 = *(*int32)(unsafe.Add(mBase, uint32(l0)+136))
	v387 = F_dsa_get_address(m, v386, v378)
	mBase = m.M
	v388 = m.ExcPending
	if v388 != 0 {
		goto L7
	} else {
		goto L74
	}
L73:
	;
	v570 = v561
	goto L71
L74:
	;
	v389 = *(*int32)(unsafe.Add(mBase, uint32(v387)+12))
	*(*int32)(unsafe.Add(mBase, uint32(v371))) = v389
	F_LWLockRelease(m, v372)
	mBase = m.M
	v392 = m.ExcPending
	if v392 != 0 {
		goto L7
	} else {
		goto L75
	}
L75:
	;
	v393 = *(*int32)(unsafe.Add(mBase, uint32(v387)+8))
	if v393 != 0 {
		goto L76
	} else {
		goto L77
	}
L76:
	;
	v400 = int32(0)
	goto L79
L77:
	;
	goto L78
L78:
	;
	v552 = *(*int32)(unsafe.Add(mBase, uint32(l0)+136))
	F_dsa_free(m, v552, v378)
	mBase = m.M
	v554 = m.ExcPending
	if v554 != 0 {
		goto L7
	} else {
		goto L103
	}
L79:
	;
	v413 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v414 = v400 + (v387 + int32(16))
	v415 = *(*int32)(unsafe.Add(mBase, uint32(v414)+4))
	v417 = v414 + int32(4)
	v419 = v414 + int32(8)
	v420 = *(*int32)(unsafe.Add(mBase, uint32(l0)+44))
	if base.Ui32(int32(2)) <= base.Ui32(v420) {
		goto L83
	} else {
		goto L84
	}
L80:
	;
	goto L78
L81:
	;
	v515 = *(*int32)(unsafe.Add(mBase, uint32(l0)+144))
	v516 = *(*int32)(unsafe.Add(mBase, uint32(v515)+20))
	v517 = int32(1)
	*(*int32)(unsafe.Add(mBase, uint32(v515)+20)) = v516 + v517
	v520 = *(*int32)(unsafe.Add(mBase, uint32(l0)+144))
	v523 = v520 + v508*int32(36)
	v524 = *(*int32)(unsafe.Add(mBase, uint32(v523)+8))
	*(*int32)(unsafe.Add(mBase, uint32(v523)+8)) = v524 + v517
	v528 = *(*int32)(unsafe.Add(mBase, uint32(v419)))
	v533 = (v528+int32(15))&int32(-8) + v400
	v534 = *(*int32)(unsafe.Add(mBase, uint32(v387)+8))
	if base.Ui32(v533) < base.Ui32(v534) {
		v400 = v533
		goto L79
	} else {
		goto L102
	}
L82:
	;
	v483 = v427 * int32(36)
	v484 = *(*int32)(unsafe.Add(mBase, uint32(l0)+144))
	v485 = v483 + v484
	v486 = *(*int32)(unsafe.Add(mBase, uint32(v485)+16))
	v487 = *(*int32)(unsafe.Add(mBase, uint32(v419)))
	*(*int32)(unsafe.Add(mBase, uint32(v485)+16)) = v486 + (v487+int32(15))&int32(-8)
	v494 = *(*int32)(unsafe.Add(mBase, uint32(l0)+144))
	v496 = *(*int32)(unsafe.Add(mBase, uint32(v494+v483)+28))
	F_sts_puttuple(m, v496, v417, v419)
	mBase = m.M
	v498 = m.ExcPending
	if v498 != 0 {
		goto L7
	} else {
		goto L101
	}
L83:
	;
	v425 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v427 = (v420 - int32(1)) & base.I32_rotr(v415, v425)
	if v427 != 0 {
		goto L82
	} else {
		goto L86
	}
L84:
	;
	goto L85
L85:
	;
	v432 = *(*int32)(unsafe.Add(mBase, uint32(v419)))
	v433 = int32(8)
	v437 = F_ExecParallelHashTupleAlloc(m, l0, v432+v433, v19+v433)
	mBase = m.M
	v438 = m.ExcPending
	if v438 != 0 {
		goto L7
	} else {
		goto L87
	}
L86:
	;
	goto L85
L87:
	;
	v439 = *(*int32)(unsafe.Add(mBase, uint32(v417)))
	*(*int32)(unsafe.Add(mBase, uint32(v437)+4)) = v439
	v441 = *(*int32)(unsafe.Add(mBase, uint32(v419)))
	if v441 != 0 {
		goto L88
	} else {
		goto L89
	}
L88:
	;
	base.MemoryCopy(m, v437+int32(8), v419, v441)
	goto L90
L89:
	;
	goto L90
L90:
	;
	v445 = *(*int32)(unsafe.Add(mBase, uint32(v19)+8))
	v446 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v449 = v446 + (v413-int32(1))&v415<<(uint(int32(2))%32)
	v450 = *(*int32)(unsafe.Add(mBase, uint32(v449)))
	*(*int32)(unsafe.Add(mBase, uint32(v437))) = v450
	v452 = *(*int32)(unsafe.Add(mBase, uint32(v449)))
	v453 = base.B2i32(v452 == v450)
	if v452 == v450 {
		goto L91
	} else {
		goto L92
	}
L91:
	;
	v454 = v445
	goto L93
L92:
	;
	v454 = v452
	goto L93
L93:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v449))) = v454
	v456 = int32(0)
	if v452 == v450 {
		v508 = v456
		goto L81
	} else {
		goto L94
	}
L94:
	;
	v459 = v452
	goto L95
L95:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v437))) = v459
	v474 = *(*int32)(unsafe.Add(mBase, uint32(v449)))
	*(*int32)(unsafe.Add(mBase, uint32(v437))) = v474
	v476 = *(*int32)(unsafe.Add(mBase, uint32(v449)))
	v477 = base.B2i32(v476 == v474)
	if v476 == v474 {
		goto L97
	} else {
		goto L98
	}
L96:
	;
	v508 = v456
	goto L81
L97:
	;
	v478 = v445
	goto L99
L98:
	;
	v478 = v476
	goto L99
L99:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v449))) = v478
	if v477 == int32(0) {
		v459 = v476
		goto L95
	} else {
		goto L100
	}
L100:
	;
	goto L96
L101:
	;
	v508 = v427
	goto L81
L102:
	;
	goto L80
L103:
	;
	v556 = *(*int32)(unsafe.Add(mBase, _c_F_ExecParallelHashIncreaseNumBatches[2]))
	if v556 != 0 {
		goto L104
	} else {
		goto L105
	}
L104:
	;
	F_ProcessInterrupts(m)
	mBase = m.M
	v558 = m.ExcPending
	if v558 != 0 {
		goto L7
	} else {
		goto L107
	}
L105:
	;
	goto L106
L106:
	;
	v559 = *(*int32)(unsafe.Add(mBase, uint32(l0)+140))
	v561 = v559 + int32(40)
	v563 = F_LWLockAcquire(m, v561, int32(0))
	mBase = m.M
	v564 = m.ExcPending
	if v564 != 0 {
		goto L7
	} else {
		goto L108
	}
L107:
	;
	goto L106
L108:
	;
	v567 = *(*int32)(unsafe.Add(mBase, uint32(v559)+24))
	if v567 != 0 {
		v371 = v559 + int32(24)
		v372 = v561
		v378 = v567
		goto L72
	} else {
		goto L109
	}
L109:
	;
	goto L73
L110:
	;
	v586 = *(*int32)(unsafe.Add(mBase, uint32(l0)+140))
	v587 = *(*int32)(unsafe.Add(mBase, uint32(v586)+12))
	v588 = *(*int32)(unsafe.Add(mBase, uint32(l0)+136))
	v589 = *(*int32)(unsafe.Add(mBase, uint32(v586)+4))
	v590 = F_dsa_get_address(m, v588, v589)
	mBase = m.M
	v591 = m.ExcPending
	if v591 != 0 {
		goto L7
	} else {
		goto L111
	}
L111:
	;
	v594 = F_palloc0(m, v587<<(uint(int32(2))%32))
	mBase = m.M
	v595 = m.ExcPending
	if v595 != 0 {
		goto L7
	} else {
		goto L112
	}
L112:
	;
	if int32(2) <= v587 {
		goto L113
	} else {
		goto L114
	}
L113:
	;
	v602 = int32(1)
	goto L116
L114:
	;
	goto L115
L115:
	;
	F_pfree(m, v594)
	mBase = m.M
	v784 = m.ExcPending
	if v784 != 0 {
		goto L7
	} else {
		goto L142
	}
L116:
	;
	v620 = *(*int32)(unsafe.Add(mBase, uint32(l0)+140))
	v621 = *(*int32)(unsafe.Add(mBase, uint32(v620)+28))
	goto L118
L117:
	;
	v652 = int32(1)
	goto L121
L118:
	;
	v626 = int32(1)
	v632 = int32(-64)
	v639 = *(*int32)(unsafe.Add(mBase, _c_F_ExecParallelHashIncreaseNumBatches[3]))
	v642 = F_sts_attach(m, v590+(((v621*int32(28)+int32(76))<<(uint(v626)%32)+int32(14))&int32(-16)-v632)*v602-v632, v639+v626, v586+int32(168))
	mBase = m.M
	v643 = m.ExcPending
	if v643 != 0 {
		goto L7
	} else {
		goto L119
	}
L119:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v594+v602<<(uint(int32(2))%32)))) = v642
	v646 = v602 + int32(1)
	if v646 != v587 {
		v602 = v646
		goto L116
	} else {
		goto L120
	}
L120:
	;
	goto L117
L121:
	;
	v667 = v594 + v652<<(uint(int32(2))%32)
	v668 = *(*int32)(unsafe.Add(mBase, uint32(v667)))
	F_sts_begin_parallel_scan(m, v668)
	mBase = m.M
	v670 = m.ExcPending
	if v670 != 0 {
		goto L7
	} else {
		goto L123
	}
L122:
	;
	goto L115
L123:
	;
	v671 = *(*int32)(unsafe.Add(mBase, uint32(v667)))
	v674 = F_sts_parallel_scan_next(m, v671, v19+int32(12))
	mBase = m.M
	v675 = m.ExcPending
	if v675 != 0 {
		goto L7
	} else {
		goto L124
	}
L124:
	;
	if v674 != 0 {
		goto L125
	} else {
		goto L126
	}
L125:
	;
	v677 = v674
	goto L128
L126:
	;
	goto L127
L127:
	;
	v761 = *(*int32)(unsafe.Add(mBase, uint32(v667)))
	F_sts_end_parallel_scan(m, v761)
	mBase = m.M
	v763 = m.ExcPending
	if v763 != 0 {
		goto L7
	} else {
		goto L140
	}
L128:
	;
	v692 = *(*int32)(unsafe.Add(mBase, uint32(v677)))
	v697 = *(*int32)(unsafe.Add(mBase, uint32(l0)+44))
	if base.Ui32(int32(2)) <= base.Ui32(v697) {
		goto L130
	} else {
		goto L131
	}
L129:
	;
	goto L127
L130:
	;
	v702 = *(*int32)(unsafe.Add(mBase, uint32(v19)+12))
	v703 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v707 = (v697 - int32(1)) & base.I32_rotr(v702, v703)
	goto L132
L131:
	;
	v707 = int32(0)
	goto L132
L132:
	;
	v708 = int32(36)
	v709 = v707 * v708
	v710 = *(*int32)(unsafe.Add(mBase, uint32(l0)+144))
	v711 = v709 + v710
	v712 = *(*int32)(unsafe.Add(mBase, uint32(v711)+16))
	*(*int32)(unsafe.Add(mBase, uint32(v711)+16)) = v712 + (v692+int32(15))&int32(-8)
	v715 = *(*int32)(unsafe.Add(mBase, uint32(l0)+144))
	v716 = v715 + v709
	v717 = *(*int32)(unsafe.Add(mBase, uint32(v716)+8))
	v718 = int32(1)
	*(*int32)(unsafe.Add(mBase, uint32(v716)+8)) = v717 + v718
	v721 = *(*int32)(unsafe.Add(mBase, uint32(l0)+144))
	v724 = v721 + v652*v708
	v725 = *(*int32)(unsafe.Add(mBase, uint32(v724)+20))
	*(*int32)(unsafe.Add(mBase, uint32(v724)+20)) = v725 + v718
	v729 = *(*int32)(unsafe.Add(mBase, uint32(l0)+144))
	v731 = *(*int32)(unsafe.Add(mBase, uint32(v729+v709)+28))
	F_sts_puttuple(m, v731, v19+int32(12), v677)
	mBase = m.M
	v735 = m.ExcPending
	if v735 != 0 {
		goto L7
	} else {
		goto L133
	}
L133:
	;
	v737 = *(*int32)(unsafe.Add(mBase, _c_F_ExecParallelHashIncreaseNumBatches[2]))
	if v737 != 0 {
		goto L134
	} else {
		goto L135
	}
L134:
	;
	F_ProcessInterrupts(m)
	mBase = m.M
	v739 = m.ExcPending
	if v739 != 0 {
		goto L7
	} else {
		goto L137
	}
L135:
	;
	goto L136
L136:
	;
	v740 = *(*int32)(unsafe.Add(mBase, uint32(v667)))
	v743 = F_sts_parallel_scan_next(m, v740, v19+int32(12))
	mBase = m.M
	v744 = m.ExcPending
	if v744 != 0 {
		goto L7
	} else {
		goto L138
	}
L137:
	;
	goto L136
L138:
	;
	if v743 != 0 {
		v677 = v743
		goto L128
	} else {
		goto L139
	}
L139:
	;
	goto L129
L140:
	;
	v765 = v652 + int32(1)
	if v765 != v587 {
		v652 = v765
		goto L121
	} else {
		goto L141
	}
L141:
	;
	goto L122
L142:
	;
	F_ExecParallelHashMergeCounters(m, l0)
	mBase = m.M
	v786 = m.ExcPending
	if v786 != 0 {
		goto L7
	} else {
		goto L143
	}
L143:
	;
	v788 = F_BarrierArriveAndWait(m, v23, int32(134217753))
	mBase = m.M
	v789 = m.ExcPending
	if v789 != 0 {
		goto L7
	} else {
		goto L144
	}
L144:
	;
	goto L3
L145:
	;
	if v807 == int32(0) {
		goto L2
	} else {
		goto L146
	}
L146:
	;
	F_ExecParallelHashEnsureBatchAccessors(m, l0)
	mBase = m.M
	v812 = m.ExcPending
	if v812 != 0 {
		goto L7
	} else {
		goto L147
	}
L147:
	;
	v813 = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(l0)+48)) = v813
	v816 = *(*int32)(unsafe.Add(mBase, uint32(l0)+136))
	v817 = *(*int32)(unsafe.Add(mBase, uint32(l0)+144))
	v818 = *(*int32)(unsafe.Add(mBase, uint32(v817)))
	v819 = *(*int32)(unsafe.Add(mBase, uint32(v818)))
	v820 = F_dsa_get_address(m, v816, v819)
	mBase = m.M
	v821 = m.ExcPending
	if v821 != 0 {
		goto L7
	} else {
		goto L148
	}
L148:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+20)) = v820
	v823 = *(*int32)(unsafe.Add(mBase, uint32(l0)+140))
	v824 = *(*int32)(unsafe.Add(mBase, uint32(v823)+16))
	*(*int32)(unsafe.Add(mBase, uint32(l0))) = v824
	v827 = int32(1073741823)
	if v827 <= v824 {
		goto L150
	} else {
		goto L151
	}
L149:
	;
	v839 = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(l0)+148)) = v839
	*(*int32)(unsafe.Add(mBase, uint32(l0)+132)) = v839
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v838
	v844 = *(*int32)(unsafe.Add(mBase, uint32(l0)+144))
	*(*uint8)(unsafe.Add(mBase, uint32(v844)+24)) = uint8(v839)
	v847 = *(*int32)(unsafe.Add(mBase, uint32(l0)+136))
	v848 = *(*int32)(unsafe.Add(mBase, uint32(v21)+4))
	v849 = F_dsa_get_address(m, v847, v848)
	mBase = m.M
	v850 = m.ExcPending
	if v850 != 0 {
		goto L7
	} else {
		goto L156
	}
L150:
	;
	v830 = v827
	goto L152
L151:
	;
	v830 = v824
	goto L152
L152:
	;
	if base.Ui32(int32(2)) <= base.Ui32(v830) {
		goto L153
	} else {
		goto L154
	}
L153:
	;
	v838 = int32(32) - base.I32_clz(v830-int32(1))
	goto L155
L154:
	;
	v838 = int32(0)
	goto L155
L155:
	;
	goto L149
L156:
	;
	v851 = *(*int32)(unsafe.Add(mBase, uint32(l0)+44))
	if v851 <= int32(0) {
		v940 = v813
		goto L157
	} else {
		goto L158
	}
L157:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v21)+20)) = v940
	v954 = *(*int32)(unsafe.Add(mBase, uint32(l0)+136))
	v955 = *(*int32)(unsafe.Add(mBase, uint32(v21)+4))
	F_dsa_free(m, v954, v955)
	mBase = m.M
	v957 = m.ExcPending
	if v957 != 0 {
		goto L7
	} else {
		goto L177
	}
L158:
	;
	v854 = int32(0)
	v858 = v854
	v861 = v854
	v862 = v854
	goto L159
L159:
	;
	v873 = *(*int32)(unsafe.Add(mBase, uint32(l0)+144))
	v877 = *(*int32)(unsafe.Add(mBase, uint32(v873+v858*int32(36))))
	v878 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v877)+60)))
	if v878 == int32(0) {
		goto L162
	} else {
		goto L163
	}
L160:
	;
	v926 = v919 | base.B2i32(int32(1073741822) < v922)
	if (v926|v885)&int32(1) == int32(0) {
		v940 = v813
		goto L157
	} else {
		goto L173
	}
L161:
	;
	v886 = *(*int32)(unsafe.Add(mBase, uint32(v21)+12))
	v887 = *(*int32)(unsafe.Add(mBase, uint32(l0)+140))
	v888 = *(*int32)(unsafe.Add(mBase, uint32(v887)+28))
	goto L167
L162:
	;
	v881 = *(*int32)(unsafe.Add(mBase, uint32(v877)+48))
	v882 = *(*int32)(unsafe.Add(mBase, uint32(v21)+32))
	if base.Ui32(v881) <= base.Ui32(v882) {
		v885 = v861
		goto L161
	} else {
		goto L165
	}
L163:
	;
	goto L164
L164:
	;
	v885 = int32(1)
	goto L161
L165:
	;
	goto L164
L166:
	;
	v921 = v858 + int32(1)
	v922 = *(*int32)(unsafe.Add(mBase, uint32(l0)+44))
	if v921 < v922 {
		v858 = v921
		v861 = v885
		v862 = v919
		goto L159
	} else {
		goto L172
	}
L167:
	;
	v901 = base.I32_rem_s(v858, v886)
	v904 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v849+(((v888*int32(28)+int32(76))<<(uint(int32(1))%32)+int32(14))&int32(-16)-int32(-64))*v901)+60)))
	if v904 == int32(0) {
		goto L168
	} else {
		goto L169
	}
L168:
	;
	v907 = *(*int32)(unsafe.Add(mBase, uint32(v877)+48))
	v908 = *(*int32)(unsafe.Add(mBase, uint32(v21)+32))
	if base.Ui32(v907) <= base.Ui32(v908) {
		v919 = v862
		goto L166
	} else {
		goto L171
	}
L169:
	;
	goto L170
L170:
	;
	v910 = *(*int32)(unsafe.Add(mBase, uint32(v877)+52))
	v911 = *(*int32)(unsafe.Add(mBase, uint32(l0)+144))
	v915 = *(*int32)(unsafe.Add(mBase, uint32(v911+v901*int32(36))))
	v916 = *(*int32)(unsafe.Add(mBase, uint32(v915)+56))
	v919 = base.B2i32(v910 == v916) | v862
	goto L166
L171:
	;
	goto L170
L172:
	;
	goto L160
L173:
	;
	if v926&int32(1) != 0 {
		goto L174
	} else {
		goto L175
	}
L174:
	;
	v936 = int32(3)
	goto L176
L175:
	;
	v936 = int32(2)
	goto L176
L176:
	;
	v940 = v936
	goto L157
L177:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v21)+4)) = int32(0)
	goto L2
L178:
	;
	goto L1
}
func F_ExecParallelInitializeDSM(m *base.Module, l0 int32, l1 int32) int32 {
	mBase := m.M
	_ = mBase
	var v3 int32
	_ = v3
	var v13 int32
	_ = v13
	var v14 int32
	_ = v14
	var v18 int32
	_ = v18
	var v19 int32
	_ = v19
	var v21 int32
	_ = v21
	var v25 int32
	_ = v25
	var v28 int32
	_ = v28
	var v29 int32
	_ = v29
	var v32 int32
	_ = v32
	var v33 int32
	_ = v33
	var v34 int32
	_ = v34
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
	var v43 int32
	_ = v43
	var v44 int32
	_ = v44
	var v45 int32
	_ = v45
	var v46 int64
	_ = v46
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
	var v56 int32
	_ = v56
	var v62 int32
	_ = v62
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
	var v68 int32
	_ = v68
	var v69 int32
	_ = v69
	var v73 int32
	_ = v73
	var v74 int32
	_ = v74
	var v78 int32
	_ = v78
	var v79 int32
	_ = v79
	var v80 int32
	_ = v80
	var v81 int64
	_ = v81
	var v83 int32
	_ = v83
	var v86 int32
	_ = v86
	var v87 int32
	_ = v87
	var v90 int32
	_ = v90
	var v91 int32
	_ = v91
	var v92 int32
	_ = v92
	var v93 int32
	_ = v93
	var v95 int32
	_ = v95
	var v96 int32
	_ = v96
	var v99 int32
	_ = v99
	var v100 int32
	_ = v100
	var v101 int32
	_ = v101
	var v102 int32
	_ = v102
	var v104 int32
	_ = v104
	var v109 int32
	_ = v109
	var v110 int32
	_ = v110
	var v111 int32
	_ = v111
	var v112 int32
	_ = v112
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
	var v123 int32
	_ = v123
	var v124 int32
	_ = v124
	var v125 int32
	_ = v125
	var v129 int32
	_ = v129
	var v130 int32
	_ = v130
	var v134 int32
	_ = v134
	var v135 int32
	_ = v135
	var v136 int32
	_ = v136
	var v137 int64
	_ = v137
	var v139 int32
	_ = v139
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
	var v151 int32
	_ = v151
	var v155 int32
	_ = v155
	var v156 int32
	_ = v156
	var v159 int32
	_ = v159
	var v160 int32
	_ = v160
	var v161 int32
	_ = v161
	var v162 int32
	_ = v162
	var v163 int32
	_ = v163
	var v165 int32
	_ = v165
	var v170 int32
	_ = v170
	var v171 int32
	_ = v171
	var v174 int32
	_ = v174
	var v177 int32
	_ = v177
	var v181 int32
	_ = v181
	var v182 int32
	_ = v182
	var v183 int32
	_ = v183
	var v185 int32
	_ = v185
	var v186 int32
	_ = v186
	var v187 int64
	_ = v187
	var v189 int32
	_ = v189
	var v190 int32
	_ = v190
	var v193 int32
	_ = v193
	var v194 int32
	_ = v194
	var v198 int32
	_ = v198
	var v199 int32
	_ = v199
	var v202 int32
	_ = v202
	var v203 int32
	_ = v203
	var v204 int32
	_ = v204
	var v205 int32
	_ = v205
	var v206 int64
	_ = v206
	var v207 int32
	_ = v207
	var v208 int32
	_ = v208
	var v209 int32
	_ = v209
	var v210 int32
	_ = v210
	var v211 int32
	_ = v211
	var v213 int32
	_ = v213
	var v214 int32
	_ = v214
	var v216 int32
	_ = v216
	var v219 int32
	_ = v219
	var v220 int32
	_ = v220
	var v223 int32
	_ = v223
	var v224 int32
	_ = v224
	var v225 int32
	_ = v225
	var v226 int32
	_ = v226
	var v227 int32
	_ = v227
	var v228 int32
	_ = v228
	var v231 int32
	_ = v231
	var v237 int32
	_ = v237
	var v238 int32
	_ = v238
	var v239 int64
	_ = v239
	var v241 int32
	_ = v241
	var v245 int32
	_ = v245
	var v246 int32
	_ = v246
	var v249 int32
	_ = v249
	var v250 int32
	_ = v250
	var v251 int32
	_ = v251
	var v252 int32
	_ = v252
	var v253 int64
	_ = v253
	var v254 int32
	_ = v254
	var v255 int32
	_ = v255
	var v256 int32
	_ = v256
	var v257 int32
	_ = v257
	var v258 int32
	_ = v258
	var v260 int32
	_ = v260
	var v261 int32
	_ = v261
	var v263 int32
	_ = v263
	var v266 int32
	_ = v266
	var v267 int32
	_ = v267
	var v270 int32
	_ = v270
	var v271 int32
	_ = v271
	var v272 int32
	_ = v272
	var v273 int32
	_ = v273
	var v274 int32
	_ = v274
	var v277 int32
	_ = v277
	var v282 int32
	_ = v282
	var v283 int32
	_ = v283
	var v284 int32
	_ = v284
	var v286 int32
	_ = v286
	var v287 int32
	_ = v287
	var v288 int32
	_ = v288
	var v289 int32
	_ = v289
	var v290 int32
	_ = v290
	var v291 int32
	_ = v291
	var v292 int32
	_ = v292
	var v293 int32
	_ = v293
	var v294 int32
	_ = v294
	var v297 int32
	_ = v297
	var v298 int32
	_ = v298
	var v301 int32
	_ = v301
	var v302 int32
	_ = v302
	var v308 int32
	_ = v308
	var v315 int32
	_ = v315
	var v318 int32
	_ = v318
	var v326 int32
	_ = v326
	var v327 int32
	_ = v327
	var v328 int64
	_ = v328
	var v330 int32
	_ = v330
	var v336 int32
	_ = v336
	var v337 int32
	_ = v337
	var v340 int32
	_ = v340
	var v341 int32
	_ = v341
	var v342 int32
	_ = v342
	var v343 int64
	_ = v343
	var v348 int32
	_ = v348
	var v350 int32
	_ = v350
	var v351 int32
	_ = v351
	var v352 int32
	_ = v352
	var v354 int32
	_ = v354
	var v355 int32
	_ = v355
	var v363 int64
	_ = v363
	var v367 int32
	_ = v367
	var v374 int32
	_ = v374
	var v375 int32
	_ = v375
	var v382 int32
	_ = v382
	var v383 int32
	_ = v383
	var v385 int64
	_ = v385
	var v395 int32
	_ = v395
	var v396 int32
	_ = v396
	var v398 int64
	_ = v398
	var v408 int32
	_ = v408
	var v409 int32
	_ = v409
	var v411 int64
	_ = v411
	var v422 int32
	_ = v422
	var v424 int32
	_ = v424
	var v425 int32
	_ = v425
	var v430 int32
	_ = v430
	var v431 int32
	_ = v431
	var v434 int32
	_ = v434
	var v437 int32
	_ = v437
	var v441 int32
	_ = v441
	var v442 int32
	_ = v442
	var v443 int32
	_ = v443
	var v447 int32
	_ = v447
	var v448 int32
	_ = v448
	var v450 int32
	_ = v450
	var v451 int32
	_ = v451
	var v452 int64
	_ = v452
	var v453 int32
	_ = v453
	var v455 int32
	_ = v455
	var v458 int32
	_ = v458
	var v459 int32
	_ = v459
	var v462 int32
	_ = v462
	var v465 int32
	_ = v465
	var v469 int32
	_ = v469
	var v470 int32
	_ = v470
	var v471 int32
	_ = v471
	var v475 int32
	_ = v475
	var v476 int32
	_ = v476
	var v478 int32
	_ = v478
	var v479 int32
	_ = v479
	var v480 int64
	_ = v480
	var v481 int32
	_ = v481
	var v483 int32
	_ = v483
	var v486 int32
	_ = v486
	var v487 int32
	_ = v487
	var v490 int32
	_ = v490
	var v493 int32
	_ = v493
	var v497 int32
	_ = v497
	var v498 int32
	_ = v498
	var v499 int32
	_ = v499
	var v503 int32
	_ = v503
	var v504 int32
	_ = v504
	var v506 int32
	_ = v506
	var v507 int32
	_ = v507
	var v508 int64
	_ = v508
	var v509 int32
	_ = v509
	var v511 int32
	_ = v511
	var v514 int32
	_ = v514
	var v515 int32
	_ = v515
	var v518 int32
	_ = v518
	var v521 int32
	_ = v521
	var v525 int32
	_ = v525
	var v526 int32
	_ = v526
	var v527 int32
	_ = v527
	var v531 int32
	_ = v531
	var v532 int32
	_ = v532
	var v534 int32
	_ = v534
	var v535 int32
	_ = v535
	var v536 int64
	_ = v536
	var v537 int32
	_ = v537
	var v539 int32
	_ = v539
	var v542 int32
	_ = v542
	var v543 int32
	_ = v543
	var v546 int32
	_ = v546
	var v549 int32
	_ = v549
	var v553 int32
	_ = v553
	var v554 int32
	_ = v554
	var v555 int32
	_ = v555
	var v559 int32
	_ = v559
	var v560 int32
	_ = v560
	var v562 int32
	_ = v562
	var v563 int32
	_ = v563
	var v564 int64
	_ = v564
	var v565 int32
	_ = v565
	var v567 int32
	_ = v567
	var v577 int32
	_ = v577
	var v578 int32
	_ = v578
	v3 = int32(0)
	if l0 == v3 {
		return int32(0)
	} else {
		v13 = *(*int32)(unsafe.Add(mBase, uint32(l1)+4))
		if v13 != 0 {
			v14 = *(*int32)(unsafe.Add(mBase, uint32(l1)+8))
			v18 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
			v19 = *(*int32)(unsafe.Add(mBase, uint32(v18)+40))
			*(*int32)(unsafe.Add(mBase, uint32(v13+v14<<(uint(int32(2))%32))+16)) = v19
		} else {
		}
		v21 = *(*int32)(unsafe.Add(mBase, uint32(l1)+8))
		*(*int32)(unsafe.Add(mBase, uint32(l1)+8)) = v21 + int32(1)
		v25 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
		switch v25 - int32(397) {
		case 0:
			v219 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
			v220 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v219)+36)))
			if v220 != int32(1) {
				v577 = F_planstate_tree_walker_impl(m, l0, int32(627), l1)
				mBase = m.M
				v578 = m.ExcPending
				if v578 != 0 {
					return int32(0)
				} else {
					return v577
				}
			} else {
				v223 = *(*int32)(unsafe.Add(mBase, uint32(l1)))
				v224 = *(*int32)(unsafe.Add(mBase, uint32(v223)+52))
				v225 = *(*int32)(unsafe.Add(mBase, uint32(l0)+164))
				v226 = F_shm_toc_allocate(m, v224, v225)
				mBase = m.M
				v227 = m.ExcPending
				if v227 != 0 {
					return int32(0)
				} else {
					v228 = *(*int32)(unsafe.Add(mBase, uint32(l0)+164))
					if v228 != 0 {
						base.MemoryFill(m, v226, int32(0), v228)
					} else {
					}
					v231 = int32(77)
					*(*uint16)(unsafe.Add(mBase, uint32(v226))) = uint16(v231)
					*(*int32)(unsafe.Add(mBase, uint32(v226)+4)) = int32(1073741824)
					*(*int64)(unsafe.Add(mBase, uint32(v226)+8)) = int64(-1)
					v237 = *(*int32)(unsafe.Add(mBase, uint32(v223)+52))
					v238 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
					v239 = int64(*(*int32)(unsafe.Add(mBase, uint32(v238)+40)))
					F_shm_toc_insert(m, v237, v239, v226)
					mBase = m.M
					v241 = m.ExcPending
					if v241 != 0 {
						return int32(0)
					} else {
						*(*int32)(unsafe.Add(mBase, uint32(l0)+184)) = int32(696)
						*(*int32)(unsafe.Add(mBase, uint32(l0)+160)) = v226
						v577 = F_planstate_tree_walker_impl(m, l0, int32(627), l1)
						mBase = m.M
						v578 = m.ExcPending
						if v578 != 0 {
							return int32(0)
						} else {
							return v577
						}
					}
				}
			}
		default:
			v577 = F_planstate_tree_walker_impl(m, l0, int32(627), l1)
			mBase = m.M
			v578 = m.ExcPending
			if v578 != 0 {
				return int32(0)
			} else {
				return v577
			}
		case 6:
			v28 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
			v29 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v28)+36)))
			if v29 != int32(1) {
				v577 = F_planstate_tree_walker_impl(m, l0, int32(627), l1)
				mBase = m.M
				v578 = m.ExcPending
				if v578 != 0 {
					return int32(0)
				} else {
					return v577
				}
			} else {
				v32 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
				v33 = *(*int32)(unsafe.Add(mBase, uint32(l1)))
				v34 = *(*int32)(unsafe.Add(mBase, uint32(v33)+52))
				v35 = *(*int32)(unsafe.Add(mBase, uint32(l0)+116))
				v36 = F_shm_toc_allocate(m, v34, v35)
				mBase = m.M
				v39 = m.ExcPending
				if v39 != 0 {
					return int32(0)
				} else {
					v40 = *(*int32)(unsafe.Add(mBase, uint32(l0)+104))
					v41 = *(*int32)(unsafe.Add(mBase, uint32(v32)+8))
					F_table_parallelscan_initialize(m, v40, v36, v41)
					mBase = m.M
					v43 = m.ExcPending
					if v43 != 0 {
						return int32(0)
					} else {
						v44 = *(*int32)(unsafe.Add(mBase, uint32(v33)+52))
						v45 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
						v46 = int64(*(*int32)(unsafe.Add(mBase, uint32(v45)+40)))
						F_shm_toc_insert(m, v44, v46, v36)
						mBase = m.M
						v48 = m.ExcPending
						if v48 != 0 {
							return int32(0)
						} else {
							v49 = *(*int32)(unsafe.Add(mBase, uint32(l0)+104))
							v50 = F_table_beginscan_parallel(m, v49, v36)
							mBase = m.M
							v51 = m.ExcPending
							if v51 != 0 {
								return int32(0)
							} else {
								*(*int32)(unsafe.Add(mBase, uint32(l0)+108)) = v50
								v577 = F_planstate_tree_walker_impl(m, l0, int32(627), l1)
								mBase = m.M
								v578 = m.ExcPending
								if v578 != 0 {
									return int32(0)
								} else {
									return v577
								}
							}
						}
					}
				}
			}
		case 8:
			v53 = *(*int32)(unsafe.Add(mBase, uint32(l1)))
			v54 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
			v55 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
			v56 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v55)+36)))
			if v54|v56&int32(1) == int32(0) {
				v577 = F_planstate_tree_walker_impl(m, l0, int32(627), l1)
				mBase = m.M
				v578 = m.ExcPending
				if v578 != 0 {
					return int32(0)
				} else {
					return v577
				}
			} else {
				v62 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
				v63 = *(*int32)(unsafe.Add(mBase, uint32(v53)+52))
				v64 = *(*int32)(unsafe.Add(mBase, uint32(l0)+208))
				v65 = F_shm_toc_allocate(m, v63, v64)
				mBase = m.M
				v66 = m.ExcPending
				if v66 != 0 {
					return int32(0)
				} else {
					v67 = *(*int32)(unsafe.Add(mBase, uint32(l0)+104))
					v68 = *(*int32)(unsafe.Add(mBase, uint32(l0)+156))
					v69 = *(*int32)(unsafe.Add(mBase, uint32(v62)+8))
					v73 = v56 & int32(1)
					v74 = *(*int32)(unsafe.Add(mBase, uint32(v53)+12))
					F_index_parallelscan_initialize(m, v67, v68, v69, base.B2i32(v54 != int32(0)), v73, v74, l0+int32(176), v65)
					mBase = m.M
					v78 = m.ExcPending
					if v78 != 0 {
						return int32(0)
					} else {
						v79 = *(*int32)(unsafe.Add(mBase, uint32(v53)+52))
						v80 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
						v81 = int64(*(*int32)(unsafe.Add(mBase, uint32(v80)+40)))
						F_shm_toc_insert(m, v79, v81, v65)
						mBase = m.M
						v83 = m.ExcPending
						if v83 != 0 {
							return int32(0)
						} else {
							if v73 == int32(0) {
								v577 = F_planstate_tree_walker_impl(m, l0, int32(627), l1)
								mBase = m.M
								v578 = m.ExcPending
								if v578 != 0 {
									return int32(0)
								} else {
									return v577
								}
							} else {
								v86 = *(*int32)(unsafe.Add(mBase, uint32(l0)+104))
								v87 = *(*int32)(unsafe.Add(mBase, uint32(l0)+156))
								v90 = *(*int32)(unsafe.Add(mBase, uint32(l0)+128))
								v91 = *(*int32)(unsafe.Add(mBase, uint32(l0)+136))
								v92 = F_index_beginscan_parallel(m, v86, v87, l0+int32(168), v90, v91, v65)
								mBase = m.M
								v93 = m.ExcPending
								if v93 != 0 {
									return int32(0)
								} else {
									*(*int32)(unsafe.Add(mBase, uint32(l0)+160)) = v92
									v95 = *(*int32)(unsafe.Add(mBase, uint32(l0)+144))
									if v95 != 0 {
										v96 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+148)))
										if v96 != int32(1) {
											v577 = F_planstate_tree_walker_impl(m, l0, int32(627), l1)
											mBase = m.M
											v578 = m.ExcPending
											if v578 != 0 {
												return int32(0)
											} else {
												return v577
											}
										} else {
											v99 = *(*int32)(unsafe.Add(mBase, uint32(l0)+124))
											v100 = *(*int32)(unsafe.Add(mBase, uint32(l0)+128))
											v101 = *(*int32)(unsafe.Add(mBase, uint32(l0)+132))
											v102 = *(*int32)(unsafe.Add(mBase, uint32(l0)+136))
											F_index_rescan(m, v92, v99, v100, v101, v102)
											mBase = m.M
											v104 = m.ExcPending
											if v104 != 0 {
												return int32(0)
											} else {
												v577 = F_planstate_tree_walker_impl(m, l0, int32(627), l1)
												mBase = m.M
												v578 = m.ExcPending
												if v578 != 0 {
													return int32(0)
												} else {
													return v577
												}
											}
										}
									} else {
										v99 = *(*int32)(unsafe.Add(mBase, uint32(l0)+124))
										v100 = *(*int32)(unsafe.Add(mBase, uint32(l0)+128))
										v101 = *(*int32)(unsafe.Add(mBase, uint32(l0)+132))
										v102 = *(*int32)(unsafe.Add(mBase, uint32(l0)+136))
										F_index_rescan(m, v92, v99, v100, v101, v102)
										mBase = m.M
										v104 = m.ExcPending
										if v104 != 0 {
											return int32(0)
										} else {
											v577 = F_planstate_tree_walker_impl(m, l0, int32(627), l1)
											mBase = m.M
											v578 = m.ExcPending
											if v578 != 0 {
												return int32(0)
											} else {
												return v577
											}
										}
									}
								}
							}
						}
					}
				}
			}
		case 9:
			v109 = *(*int32)(unsafe.Add(mBase, uint32(l1)))
			v110 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
			v111 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
			v112 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v111)+36)))
			if v110|v112&int32(1) == int32(0) {
				v577 = F_planstate_tree_walker_impl(m, l0, int32(627), l1)
				mBase = m.M
				v578 = m.ExcPending
				if v578 != 0 {
					return int32(0)
				} else {
					return v577
				}
			} else {
				v118 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
				v119 = *(*int32)(unsafe.Add(mBase, uint32(v109)+52))
				v120 = *(*int32)(unsafe.Add(mBase, uint32(l0)+180))
				v121 = F_shm_toc_allocate(m, v119, v120)
				mBase = m.M
				v122 = m.ExcPending
				if v122 != 0 {
					return int32(0)
				} else {
					v123 = *(*int32)(unsafe.Add(mBase, uint32(l0)+104))
					v124 = *(*int32)(unsafe.Add(mBase, uint32(l0)+152))
					v125 = *(*int32)(unsafe.Add(mBase, uint32(v118)+8))
					v129 = v112 & int32(1)
					v130 = *(*int32)(unsafe.Add(mBase, uint32(v109)+12))
					F_index_parallelscan_initialize(m, v123, v124, v125, base.B2i32(v110 != int32(0)), v129, v130, l0+int32(168), v121)
					mBase = m.M
					v134 = m.ExcPending
					if v134 != 0 {
						return int32(0)
					} else {
						v135 = *(*int32)(unsafe.Add(mBase, uint32(v109)+52))
						v136 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
						v137 = int64(*(*int32)(unsafe.Add(mBase, uint32(v136)+40)))
						F_shm_toc_insert(m, v135, v137, v121)
						mBase = m.M
						v139 = m.ExcPending
						if v139 != 0 {
							return int32(0)
						} else {
							if v129 == int32(0) {
								v577 = F_planstate_tree_walker_impl(m, l0, int32(627), l1)
								mBase = m.M
								v578 = m.ExcPending
								if v578 != 0 {
									return int32(0)
								} else {
									return v577
								}
							} else {
								v142 = *(*int32)(unsafe.Add(mBase, uint32(l0)+104))
								v143 = *(*int32)(unsafe.Add(mBase, uint32(l0)+152))
								v146 = *(*int32)(unsafe.Add(mBase, uint32(l0)+124))
								v147 = *(*int32)(unsafe.Add(mBase, uint32(l0)+132))
								v148 = F_index_beginscan_parallel(m, v142, v143, l0+int32(160), v146, v147, v121)
								mBase = m.M
								v149 = m.ExcPending
								if v149 != 0 {
									return int32(0)
								} else {
									*(*int32)(unsafe.Add(mBase, uint32(l0)+156)) = v148
									v151 = int32(1)
									*(*uint8)(unsafe.Add(mBase, uint32(v148)+28)) = uint8(v151)
									*(*int32)(unsafe.Add(mBase, uint32(l0)+176)) = int32(0)
									v155 = *(*int32)(unsafe.Add(mBase, uint32(l0)+140))
									if v155 != 0 {
										v156 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+144)))
										if v156 != int32(1) {
											v577 = F_planstate_tree_walker_impl(m, l0, int32(627), l1)
											mBase = m.M
											v578 = m.ExcPending
											if v578 != 0 {
												return int32(0)
											} else {
												return v577
											}
										} else {
											v159 = *(*int32)(unsafe.Add(mBase, uint32(l0)+156))
											v160 = *(*int32)(unsafe.Add(mBase, uint32(l0)+120))
											v161 = *(*int32)(unsafe.Add(mBase, uint32(l0)+124))
											v162 = *(*int32)(unsafe.Add(mBase, uint32(l0)+128))
											v163 = *(*int32)(unsafe.Add(mBase, uint32(l0)+132))
											F_index_rescan(m, v159, v160, v161, v162, v163)
											mBase = m.M
											v165 = m.ExcPending
											if v165 != 0 {
												return int32(0)
											} else {
												v577 = F_planstate_tree_walker_impl(m, l0, int32(627), l1)
												mBase = m.M
												v578 = m.ExcPending
												if v578 != 0 {
													return int32(0)
												} else {
													return v577
												}
											}
										}
									} else {
										v159 = *(*int32)(unsafe.Add(mBase, uint32(l0)+156))
										v160 = *(*int32)(unsafe.Add(mBase, uint32(l0)+120))
										v161 = *(*int32)(unsafe.Add(mBase, uint32(l0)+124))
										v162 = *(*int32)(unsafe.Add(mBase, uint32(l0)+128))
										v163 = *(*int32)(unsafe.Add(mBase, uint32(l0)+132))
										F_index_rescan(m, v159, v160, v161, v162, v163)
										mBase = m.M
										v165 = m.ExcPending
										if v165 != 0 {
											return int32(0)
										} else {
											v577 = F_planstate_tree_walker_impl(m, l0, int32(627), l1)
											mBase = m.M
											v578 = m.ExcPending
											if v578 != 0 {
												return int32(0)
											} else {
												return v577
											}
										}
									}
								}
							}
						}
					}
				}
			}
		case 10:
			v170 = *(*int32)(unsafe.Add(mBase, uint32(l1)))
			v171 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
			if v171 == int32(0) {
				v577 = F_planstate_tree_walker_impl(m, l0, int32(627), l1)
				mBase = m.M
				v578 = m.ExcPending
				if v578 != 0 {
					return int32(0)
				} else {
					return v577
				}
			} else {
				v174 = *(*int32)(unsafe.Add(mBase, uint32(v170)+12))
				if v174 == int32(0) {
					v577 = F_planstate_tree_walker_impl(m, l0, int32(627), l1)
					mBase = m.M
					v578 = m.ExcPending
					if v578 != 0 {
						return int32(0)
					} else {
						return v577
					}
				} else {
					v177 = *(*int32)(unsafe.Add(mBase, uint32(v170)+52))
					v181 = v174<<(uint(int32(3))%32) + int32(8)
					v182 = F_shm_toc_allocate(m, v177, v181)
					mBase = m.M
					v183 = m.ExcPending
					if v183 != 0 {
						return int32(0)
					} else {
						*(*int32)(unsafe.Add(mBase, uint32(l0)+168)) = v182
						v185 = *(*int32)(unsafe.Add(mBase, uint32(v170)+52))
						v186 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
						v187 = int64(*(*int32)(unsafe.Add(mBase, uint32(v186)+40)))
						F_shm_toc_insert(m, v185, v187, v182)
						mBase = m.M
						v189 = m.ExcPending
						if v189 != 0 {
							return int32(0)
						} else {
							if v181 != 0 {
								v190 = *(*int32)(unsafe.Add(mBase, uint32(l0)+168))
								base.MemoryFill(m, v190, int32(0), v181)
							} else {
							}
							v193 = *(*int32)(unsafe.Add(mBase, uint32(l0)+168))
							v194 = *(*int32)(unsafe.Add(mBase, uint32(v170)+12))
							*(*int32)(unsafe.Add(mBase, uint32(v193))) = v194
							v577 = F_planstate_tree_walker_impl(m, l0, int32(627), l1)
							mBase = m.M
							v578 = m.ExcPending
							if v578 != 0 {
								return int32(0)
							} else {
								return v577
							}
						}
					}
				}
			}
		case 11:
			v266 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
			v267 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v266)+36)))
			if v267 != int32(1) {
				v577 = F_planstate_tree_walker_impl(m, l0, int32(627), l1)
				mBase = m.M
				v578 = m.ExcPending
				if v578 != 0 {
					return int32(0)
				} else {
					return v577
				}
			} else {
				v270 = *(*int32)(unsafe.Add(mBase, uint32(l1)))
				v271 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
				v272 = *(*int32)(unsafe.Add(mBase, uint32(v271)+172))
				if v272 != 0 {
					v273 = int32(24)
					v274 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
					if v274 == int32(0) {
						v290 = v273
						v291 = *(*int32)(unsafe.Add(mBase, uint32(v270)+52))
						v292 = F_shm_toc_allocate(m, v291, v290)
						mBase = m.M
						v293 = m.ExcPending
						if v293 != 0 {
							return int32(0)
						} else {
							v294 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
							if v294 != 0 {
								v297 = int32(0)
								v298 = *(*int32)(unsafe.Add(mBase, uint32(v270)+12))
								if v297 < v298 {
									v301 = v292 + int32(24)
								} else {
									v301 = v297
								}
								v302 = v301
							} else {
								v302 = v3
							}
							*(*int32)(unsafe.Add(mBase, uint32(v292)+8)) = int32(0)
							*(*int64)(unsafe.Add(mBase, uint32(v292))) = int64(0)
							v308 = v292 + int32(12)
							*(*int32)(unsafe.Add(mBase, uint32(v308)+8)) = int32(-1)
							*(*int64)(unsafe.Add(mBase, uint32(v308))) = int64(-4294967296)
							if v302 == int32(0) {
							} else {
								v315 = *(*int32)(unsafe.Add(mBase, uint32(v270)+12))
								*(*int32)(unsafe.Add(mBase, uint32(v302))) = v315
								v318 = v315 << (uint(int32(4)) % 32)
								if v318 == int32(0) {
								} else {
									base.MemoryFill(m, v302+int32(8), int32(0), v318)
								}
							}
							v326 = *(*int32)(unsafe.Add(mBase, uint32(v270)+52))
							v327 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
							v328 = int64(*(*int32)(unsafe.Add(mBase, uint32(v327)+40)))
							F_shm_toc_insert(m, v326, v328, v292)
							mBase = m.M
							v330 = m.ExcPending
							if v330 != 0 {
								return int32(0)
							} else {
								*(*int32)(unsafe.Add(mBase, uint32(l0)+152)) = v302
								*(*int32)(unsafe.Add(mBase, uint32(l0)+148)) = v292
								v577 = F_planstate_tree_walker_impl(m, l0, int32(627), l1)
								mBase = m.M
								v578 = m.ExcPending
								if v578 != 0 {
									return int32(0)
								} else {
									return v577
								}
							}
						}
					} else {
						v277 = *(*int32)(unsafe.Add(mBase, uint32(v270)+12))
						if v277 <= int32(0) {
							v290 = v273
							v291 = *(*int32)(unsafe.Add(mBase, uint32(v270)+52))
							v292 = F_shm_toc_allocate(m, v291, v290)
							mBase = m.M
							v293 = m.ExcPending
							if v293 != 0 {
								return int32(0)
							} else {
								v294 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
								if v294 != 0 {
									v297 = int32(0)
									v298 = *(*int32)(unsafe.Add(mBase, uint32(v270)+12))
									if v297 < v298 {
										v301 = v292 + int32(24)
									} else {
										v301 = v297
									}
									v302 = v301
								} else {
									v302 = v3
								}
								*(*int32)(unsafe.Add(mBase, uint32(v292)+8)) = int32(0)
								*(*int64)(unsafe.Add(mBase, uint32(v292))) = int64(0)
								v308 = v292 + int32(12)
								*(*int32)(unsafe.Add(mBase, uint32(v308)+8)) = int32(-1)
								*(*int64)(unsafe.Add(mBase, uint32(v308))) = int64(-4294967296)
								if v302 == int32(0) {
								} else {
									v315 = *(*int32)(unsafe.Add(mBase, uint32(v270)+12))
									*(*int32)(unsafe.Add(mBase, uint32(v302))) = v315
									v318 = v315 << (uint(int32(4)) % 32)
									if v318 == int32(0) {
									} else {
										base.MemoryFill(m, v302+int32(8), int32(0), v318)
									}
								}
								v326 = *(*int32)(unsafe.Add(mBase, uint32(v270)+52))
								v327 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
								v328 = int64(*(*int32)(unsafe.Add(mBase, uint32(v327)+40)))
								F_shm_toc_insert(m, v326, v328, v292)
								mBase = m.M
								v330 = m.ExcPending
								if v330 != 0 {
									return int32(0)
								} else {
									*(*int32)(unsafe.Add(mBase, uint32(l0)+152)) = v302
									*(*int32)(unsafe.Add(mBase, uint32(l0)+148)) = v292
									v577 = F_planstate_tree_walker_impl(m, l0, int32(627), l1)
									mBase = m.M
									v578 = m.ExcPending
									if v578 != 0 {
										return int32(0)
									} else {
										return v577
									}
								}
							}
						} else {
							v282 = F_add_size(m, int32(24), int32(8))
							mBase = m.M
							v283 = m.ExcPending
							if v283 != 0 {
								return int32(0)
							} else {
								v284 = *(*int32)(unsafe.Add(mBase, uint32(v270)+12))
								v286 = F_mul_size(m, v284, int32(16))
								mBase = m.M
								v287 = m.ExcPending
								if v287 != 0 {
									return int32(0)
								} else {
									v288 = F_add_size(m, v282, v286)
									mBase = m.M
									v289 = m.ExcPending
									if v289 != 0 {
										return int32(0)
									} else {
										v290 = v288
										v291 = *(*int32)(unsafe.Add(mBase, uint32(v270)+52))
										v292 = F_shm_toc_allocate(m, v291, v290)
										mBase = m.M
										v293 = m.ExcPending
										if v293 != 0 {
											return int32(0)
										} else {
											v294 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
											if v294 != 0 {
												v297 = int32(0)
												v298 = *(*int32)(unsafe.Add(mBase, uint32(v270)+12))
												if v297 < v298 {
													v301 = v292 + int32(24)
												} else {
													v301 = v297
												}
												v302 = v301
											} else {
												v302 = v3
											}
											*(*int32)(unsafe.Add(mBase, uint32(v292)+8)) = int32(0)
											*(*int64)(unsafe.Add(mBase, uint32(v292))) = int64(0)
											v308 = v292 + int32(12)
											*(*int32)(unsafe.Add(mBase, uint32(v308)+8)) = int32(-1)
											*(*int64)(unsafe.Add(mBase, uint32(v308))) = int64(-4294967296)
											if v302 == int32(0) {
											} else {
												v315 = *(*int32)(unsafe.Add(mBase, uint32(v270)+12))
												*(*int32)(unsafe.Add(mBase, uint32(v302))) = v315
												v318 = v315 << (uint(int32(4)) % 32)
												if v318 == int32(0) {
												} else {
													base.MemoryFill(m, v302+int32(8), int32(0), v318)
												}
											}
											v326 = *(*int32)(unsafe.Add(mBase, uint32(v270)+52))
											v327 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
											v328 = int64(*(*int32)(unsafe.Add(mBase, uint32(v327)+40)))
											F_shm_toc_insert(m, v326, v328, v292)
											mBase = m.M
											v330 = m.ExcPending
											if v330 != 0 {
												return int32(0)
											} else {
												*(*int32)(unsafe.Add(mBase, uint32(l0)+152)) = v302
												*(*int32)(unsafe.Add(mBase, uint32(l0)+148)) = v292
												v577 = F_planstate_tree_walker_impl(m, l0, int32(627), l1)
												mBase = m.M
												v578 = m.ExcPending
												if v578 != 0 {
													return int32(0)
												} else {
													return v577
												}
											}
										}
									}
								}
							}
						}
					}
				} else {
					v577 = F_planstate_tree_walker_impl(m, l0, int32(627), l1)
					mBase = m.M
					v578 = m.ExcPending
					if v578 != 0 {
						return int32(0)
					} else {
						return v577
					}
				}
			}
		case 21:
			v198 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
			v199 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v198)+36)))
			if v199 != int32(1) {
				v577 = F_planstate_tree_walker_impl(m, l0, int32(627), l1)
				mBase = m.M
				v578 = m.ExcPending
				if v578 != 0 {
					return int32(0)
				} else {
					return v577
				}
			} else {
				v202 = *(*int32)(unsafe.Add(mBase, uint32(l1)))
				v203 = *(*int32)(unsafe.Add(mBase, uint32(l0)+128))
				v204 = *(*int32)(unsafe.Add(mBase, uint32(v203)+148))
				if v204 != 0 {
					v205 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
					v206 = int64(*(*int32)(unsafe.Add(mBase, uint32(v205)+40)))
					v207 = *(*int32)(unsafe.Add(mBase, uint32(v202)+52))
					v208 = *(*int32)(unsafe.Add(mBase, uint32(l0)+120))
					v209 = F_shm_toc_allocate(m, v207, v208)
					mBase = m.M
					v210 = m.ExcPending
					if v210 != 0 {
						return int32(0)
					} else {
						v211 = *(*int32)(unsafe.Add(mBase, uint32(v203)+148))
						m.T0[v211].(func(*base.Module, int32, int32, int32))(m, l0, v202, v209)
						mBase = m.M
						v213 = m.ExcPending
						if v213 != 0 {
							return int32(0)
						} else {
							v214 = *(*int32)(unsafe.Add(mBase, uint32(v202)+52))
							F_shm_toc_insert(m, v214, v206, v209)
							mBase = m.M
							v216 = m.ExcPending
							if v216 != 0 {
								return int32(0)
							} else {
								v577 = F_planstate_tree_walker_impl(m, l0, int32(627), l1)
								mBase = m.M
								v578 = m.ExcPending
								if v578 != 0 {
									return int32(0)
								} else {
									return v577
								}
							}
						}
					}
				} else {
					v577 = F_planstate_tree_walker_impl(m, l0, int32(627), l1)
					mBase = m.M
					v578 = m.ExcPending
					if v578 != 0 {
						return int32(0)
					} else {
						return v577
					}
				}
			}
		case 22:
			v245 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
			v246 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v245)+36)))
			if v246 != int32(1) {
				v577 = F_planstate_tree_walker_impl(m, l0, int32(627), l1)
				mBase = m.M
				v578 = m.ExcPending
				if v578 != 0 {
					return int32(0)
				} else {
					return v577
				}
			} else {
				v249 = *(*int32)(unsafe.Add(mBase, uint32(l1)))
				v250 = *(*int32)(unsafe.Add(mBase, uint32(l0)+128))
				v251 = *(*int32)(unsafe.Add(mBase, uint32(v250)+32))
				if v251 != 0 {
					v252 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
					v253 = int64(*(*int32)(unsafe.Add(mBase, uint32(v252)+40)))
					v254 = *(*int32)(unsafe.Add(mBase, uint32(v249)+52))
					v255 = *(*int32)(unsafe.Add(mBase, uint32(l0)+124))
					v256 = F_shm_toc_allocate(m, v254, v255)
					mBase = m.M
					v257 = m.ExcPending
					if v257 != 0 {
						return int32(0)
					} else {
						v258 = *(*int32)(unsafe.Add(mBase, uint32(v250)+32))
						m.T0[v258].(func(*base.Module, int32, int32, int32))(m, l0, v249, v256)
						mBase = m.M
						v260 = m.ExcPending
						if v260 != 0 {
							return int32(0)
						} else {
							v261 = *(*int32)(unsafe.Add(mBase, uint32(v249)+52))
							F_shm_toc_insert(m, v261, v253, v256)
							mBase = m.M
							v263 = m.ExcPending
							if v263 != 0 {
								return int32(0)
							} else {
								v577 = F_planstate_tree_walker_impl(m, l0, int32(627), l1)
								mBase = m.M
								v578 = m.ExcPending
								if v578 != 0 {
									return int32(0)
								} else {
									return v577
								}
							}
						}
					}
				} else {
					v577 = F_planstate_tree_walker_impl(m, l0, int32(627), l1)
					mBase = m.M
					v578 = m.ExcPending
					if v578 != 0 {
						return int32(0)
					} else {
						return v577
					}
				}
			}
		case 26:
			v336 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
			v337 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v336)+36)))
			if v337 != int32(1) {
				v577 = F_planstate_tree_walker_impl(m, l0, int32(627), l1)
				mBase = m.M
				v578 = m.ExcPending
				if v578 != 0 {
					return int32(0)
				} else {
					return v577
				}
			} else {
				v340 = *(*int32)(unsafe.Add(mBase, uint32(l1)))
				v341 = *(*int32)(unsafe.Add(mBase, uint32(v340)+44))
				if v341 != 0 {
					v342 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
					v343 = int64(*(*int32)(unsafe.Add(mBase, uint32(v342)+40)))
					*(*int32)(unsafe.Add(mBase, uint32(l0)+12)) = int32(632)
					*(*int32)(unsafe.Add(mBase, uint32(l0)+16)) = int32(720)
					v348 = *(*int32)(unsafe.Add(mBase, uint32(v340)+52))
					v350 = F_shm_toc_allocate(m, v348, int32(220))
					mBase = m.M
					v351 = m.ExcPending
					if v351 != 0 {
						return int32(0)
					} else {
						v352 = *(*int32)(unsafe.Add(mBase, uint32(v340)+52))
						F_shm_toc_insert(m, v352, v343, v350)
						mBase = m.M
						v354 = m.ExcPending
						if v354 != 0 {
							return int32(0)
						} else {
							v355 = int32(0)
							*(*int32)(unsafe.Add(mBase, uint32(v350)+32)) = v355
							*(*int32)(unsafe.Add(mBase, uint32(v350)+8)) = v355
							*(*int32)(unsafe.Add(mBase, uint32(v350)+164)) = v355
							*(*int32)(unsafe.Add(mBase, uint32(v350)+24)) = v355
							v363 = int64(0)
							*(*int64)(unsafe.Add(mBase, uint32(v350)+16)) = v363
							*(*int64)(unsafe.Add(mBase, uint32(v350))) = v363
							v367 = *(*int32)(unsafe.Add(mBase, uint32(v340)+12))
							*(*int32)(unsafe.Add(mBase, uint32(v350)+36)) = v355
							*(*int32)(unsafe.Add(mBase, uint32(v350)+28)) = v367 + int32(1)
							v374 = v350 + int32(40)
							v375 = int32(69)
							*(*uint16)(unsafe.Add(mBase, uint32(v374))) = uint16(v375)
							*(*int32)(unsafe.Add(mBase, uint32(v374)+4)) = int32(1073741824)
							*(*int64)(unsafe.Add(mBase, uint32(v374)+8)) = int64(-1)
							v382 = v350 + int32(56)
							v383 = int32(0)
							*(*int32)(unsafe.Add(mBase, uint32(v382)+8)) = v383
							v385 = int64(0)
							*(*int64)(unsafe.Add(mBase, uint32(v382))) = v385
							*(*int64)(unsafe.Add(mBase, uint32(v382)+12)) = v385
							*(*uint8)(unsafe.Add(mBase, uint32(v382)+20)) = uint8(v383)
							F_ConditionVariableInit(m, v350+int32(80))
							mBase = m.M
							v395 = v350 + int32(92)
							v396 = int32(0)
							*(*int32)(unsafe.Add(mBase, uint32(v395)+8)) = v396
							v398 = int64(0)
							*(*int64)(unsafe.Add(mBase, uint32(v395))) = v398
							*(*int64)(unsafe.Add(mBase, uint32(v395)+12)) = v398
							*(*uint8)(unsafe.Add(mBase, uint32(v395)+20)) = uint8(v396)
							F_ConditionVariableInit(m, v350+int32(116))
							mBase = m.M
							v408 = v350 + int32(128)
							v409 = int32(0)
							*(*int32)(unsafe.Add(mBase, uint32(v408)+8)) = v409
							v411 = int64(0)
							*(*int64)(unsafe.Add(mBase, uint32(v408))) = v411
							*(*int64)(unsafe.Add(mBase, uint32(v408)+12)) = v411
							*(*uint8)(unsafe.Add(mBase, uint32(v408)+20)) = uint8(v409)
							F_ConditionVariableInit(m, v350+int32(152))
							mBase = m.M
							v422 = *(*int32)(unsafe.Add(mBase, uint32(v340)+44))
							F_SharedFileSetInit(m, v350+int32(168), v422)
							mBase = m.M
							v424 = m.ExcPending
							if v424 != 0 {
								return int32(0)
							} else {
								v425 = *(*int32)(unsafe.Add(mBase, uint32(l0)+40))
								*(*int32)(unsafe.Add(mBase, uint32(v425)+128)) = v350
								v577 = F_planstate_tree_walker_impl(m, l0, int32(627), l1)
								mBase = m.M
								v578 = m.ExcPending
								if v578 != 0 {
									return int32(0)
								} else {
									return v577
								}
							}
						}
					}
				} else {
					v577 = F_planstate_tree_walker_impl(m, l0, int32(627), l1)
					mBase = m.M
					v578 = m.ExcPending
					if v578 != 0 {
						return int32(0)
					} else {
						return v577
					}
				}
			}
		case 28:
			v542 = *(*int32)(unsafe.Add(mBase, uint32(l1)))
			v543 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
			if v543 == int32(0) {
				v577 = F_planstate_tree_walker_impl(m, l0, int32(627), l1)
				mBase = m.M
				v578 = m.ExcPending
				if v578 != 0 {
					return int32(0)
				} else {
					return v577
				}
			} else {
				v546 = *(*int32)(unsafe.Add(mBase, uint32(v542)+12))
				if v546 == int32(0) {
					v577 = F_planstate_tree_walker_impl(m, l0, int32(627), l1)
					mBase = m.M
					v578 = m.ExcPending
					if v578 != 0 {
						return int32(0)
					} else {
						return v577
					}
				} else {
					v549 = *(*int32)(unsafe.Add(mBase, uint32(v542)+52))
					v553 = v546*int32(40) + int32(8)
					v554 = F_shm_toc_allocate(m, v549, v553)
					mBase = m.M
					v555 = m.ExcPending
					if v555 != 0 {
						return int32(0)
					} else {
						*(*int32)(unsafe.Add(mBase, uint32(l0)+240)) = v554
						if v553 != 0 {
							base.MemoryFill(m, v554, int32(0), v553)
						} else {
						}
						v559 = *(*int32)(unsafe.Add(mBase, uint32(l0)+240))
						v560 = *(*int32)(unsafe.Add(mBase, uint32(v542)+12))
						*(*int32)(unsafe.Add(mBase, uint32(v559))) = v560
						v562 = *(*int32)(unsafe.Add(mBase, uint32(v542)+52))
						v563 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
						v564 = int64(*(*int32)(unsafe.Add(mBase, uint32(v563)+40)))
						v565 = *(*int32)(unsafe.Add(mBase, uint32(l0)+240))
						F_shm_toc_insert(m, v562, v564, v565)
						mBase = m.M
						v567 = m.ExcPending
						if v567 != 0 {
							return int32(0)
						} else {
							v577 = F_planstate_tree_walker_impl(m, l0, int32(627), l1)
							mBase = m.M
							v578 = m.ExcPending
							if v578 != 0 {
								return int32(0)
							} else {
								return v577
							}
						}
					}
				}
			}
		case 29:
			v458 = *(*int32)(unsafe.Add(mBase, uint32(l1)))
			v459 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
			if v459 == int32(0) {
				v577 = F_planstate_tree_walker_impl(m, l0, int32(627), l1)
				mBase = m.M
				v578 = m.ExcPending
				if v578 != 0 {
					return int32(0)
				} else {
					return v577
				}
			} else {
				v462 = *(*int32)(unsafe.Add(mBase, uint32(v458)+12))
				if v462 == int32(0) {
					v577 = F_planstate_tree_walker_impl(m, l0, int32(627), l1)
					mBase = m.M
					v578 = m.ExcPending
					if v578 != 0 {
						return int32(0)
					} else {
						return v577
					}
				} else {
					v465 = *(*int32)(unsafe.Add(mBase, uint32(v458)+52))
					v469 = v462<<(uint(int32(4))%32) | int32(8)
					v470 = F_shm_toc_allocate(m, v465, v469)
					mBase = m.M
					v471 = m.ExcPending
					if v471 != 0 {
						return int32(0)
					} else {
						*(*int32)(unsafe.Add(mBase, uint32(l0)+152)) = v470
						if v469 != 0 {
							base.MemoryFill(m, v470, int32(0), v469)
						} else {
						}
						v475 = *(*int32)(unsafe.Add(mBase, uint32(l0)+152))
						v476 = *(*int32)(unsafe.Add(mBase, uint32(v458)+12))
						*(*int32)(unsafe.Add(mBase, uint32(v475))) = v476
						v478 = *(*int32)(unsafe.Add(mBase, uint32(v458)+52))
						v479 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
						v480 = int64(*(*int32)(unsafe.Add(mBase, uint32(v479)+40)))
						v481 = *(*int32)(unsafe.Add(mBase, uint32(l0)+152))
						F_shm_toc_insert(m, v478, v480, v481)
						mBase = m.M
						v483 = m.ExcPending
						if v483 != 0 {
							return int32(0)
						} else {
							v577 = F_planstate_tree_walker_impl(m, l0, int32(627), l1)
							mBase = m.M
							v578 = m.ExcPending
							if v578 != 0 {
								return int32(0)
							} else {
								return v577
							}
						}
					}
				}
			}
		case 30:
			v486 = *(*int32)(unsafe.Add(mBase, uint32(l1)))
			v487 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
			if v487 == int32(0) {
				v577 = F_planstate_tree_walker_impl(m, l0, int32(627), l1)
				mBase = m.M
				v578 = m.ExcPending
				if v578 != 0 {
					return int32(0)
				} else {
					return v577
				}
			} else {
				v490 = *(*int32)(unsafe.Add(mBase, uint32(v486)+12))
				if v490 == int32(0) {
					v577 = F_planstate_tree_walker_impl(m, l0, int32(627), l1)
					mBase = m.M
					v578 = m.ExcPending
					if v578 != 0 {
						return int32(0)
					} else {
						return v577
					}
				} else {
					v493 = *(*int32)(unsafe.Add(mBase, uint32(v486)+52))
					v497 = v490*int32(96) | int32(8)
					v498 = F_shm_toc_allocate(m, v493, v497)
					mBase = m.M
					v499 = m.ExcPending
					if v499 != 0 {
						return int32(0)
					} else {
						*(*int32)(unsafe.Add(mBase, uint32(l0)+284)) = v498
						if v497 != 0 {
							base.MemoryFill(m, v498, int32(0), v497)
						} else {
						}
						v503 = *(*int32)(unsafe.Add(mBase, uint32(l0)+284))
						v504 = *(*int32)(unsafe.Add(mBase, uint32(v486)+12))
						*(*int32)(unsafe.Add(mBase, uint32(v503))) = v504
						v506 = *(*int32)(unsafe.Add(mBase, uint32(v486)+52))
						v507 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
						v508 = int64(*(*int32)(unsafe.Add(mBase, uint32(v507)+40)))
						v509 = *(*int32)(unsafe.Add(mBase, uint32(l0)+284))
						F_shm_toc_insert(m, v506, v508, v509)
						mBase = m.M
						v511 = m.ExcPending
						if v511 != 0 {
							return int32(0)
						} else {
							v577 = F_planstate_tree_walker_impl(m, l0, int32(627), l1)
							mBase = m.M
							v578 = m.ExcPending
							if v578 != 0 {
								return int32(0)
							} else {
								return v577
							}
						}
					}
				}
			}
		case 32:
			v514 = *(*int32)(unsafe.Add(mBase, uint32(l1)))
			v515 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
			if v515 == int32(0) {
				v577 = F_planstate_tree_walker_impl(m, l0, int32(627), l1)
				mBase = m.M
				v578 = m.ExcPending
				if v578 != 0 {
					return int32(0)
				} else {
					return v577
				}
			} else {
				v518 = *(*int32)(unsafe.Add(mBase, uint32(v514)+12))
				if v518 == int32(0) {
					v577 = F_planstate_tree_walker_impl(m, l0, int32(627), l1)
					mBase = m.M
					v578 = m.ExcPending
					if v578 != 0 {
						return int32(0)
					} else {
						return v577
					}
				} else {
					v521 = *(*int32)(unsafe.Add(mBase, uint32(v514)+52))
					v525 = v518*int32(24) + int32(8)
					v526 = F_shm_toc_allocate(m, v521, v525)
					mBase = m.M
					v527 = m.ExcPending
					if v527 != 0 {
						return int32(0)
					} else {
						*(*int32)(unsafe.Add(mBase, uint32(l0)+352)) = v526
						if v525 != 0 {
							base.MemoryFill(m, v526, int32(0), v525)
						} else {
						}
						v531 = *(*int32)(unsafe.Add(mBase, uint32(l0)+352))
						v532 = *(*int32)(unsafe.Add(mBase, uint32(v514)+12))
						*(*int32)(unsafe.Add(mBase, uint32(v531))) = v532
						v534 = *(*int32)(unsafe.Add(mBase, uint32(v514)+52))
						v535 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
						v536 = int64(*(*int32)(unsafe.Add(mBase, uint32(v535)+40)))
						v537 = *(*int32)(unsafe.Add(mBase, uint32(l0)+352))
						F_shm_toc_insert(m, v534, v536, v537)
						mBase = m.M
						v539 = m.ExcPending
						if v539 != 0 {
							return int32(0)
						} else {
							v577 = F_planstate_tree_walker_impl(m, l0, int32(627), l1)
							mBase = m.M
							v578 = m.ExcPending
							if v578 != 0 {
								return int32(0)
							} else {
								return v577
							}
						}
					}
				}
			}
		case 37:
			v430 = *(*int32)(unsafe.Add(mBase, uint32(l1)))
			v431 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
			if v431 == int32(0) {
				v577 = F_planstate_tree_walker_impl(m, l0, int32(627), l1)
				mBase = m.M
				v578 = m.ExcPending
				if v578 != 0 {
					return int32(0)
				} else {
					return v577
				}
			} else {
				v434 = *(*int32)(unsafe.Add(mBase, uint32(v430)+12))
				if v434 == int32(0) {
					v577 = F_planstate_tree_walker_impl(m, l0, int32(627), l1)
					mBase = m.M
					v578 = m.ExcPending
					if v578 != 0 {
						return int32(0)
					} else {
						return v577
					}
				} else {
					v437 = *(*int32)(unsafe.Add(mBase, uint32(v430)+52))
					v441 = v434*int32(20) + int32(4)
					v442 = F_shm_toc_allocate(m, v437, v441)
					mBase = m.M
					v443 = m.ExcPending
					if v443 != 0 {
						return int32(0)
					} else {
						*(*int32)(unsafe.Add(mBase, uint32(l0)+120)) = v442
						if v441 != 0 {
							base.MemoryFill(m, v442, int32(0), v441)
						} else {
						}
						v447 = *(*int32)(unsafe.Add(mBase, uint32(l0)+120))
						v448 = *(*int32)(unsafe.Add(mBase, uint32(v430)+12))
						*(*int32)(unsafe.Add(mBase, uint32(v447))) = v448
						v450 = *(*int32)(unsafe.Add(mBase, uint32(v430)+52))
						v451 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
						v452 = int64(*(*int32)(unsafe.Add(mBase, uint32(v451)+40)))
						v453 = *(*int32)(unsafe.Add(mBase, uint32(l0)+120))
						F_shm_toc_insert(m, v450, v452, v453)
						mBase = m.M
						v455 = m.ExcPending
						if v455 != 0 {
							return int32(0)
						} else {
							v577 = F_planstate_tree_walker_impl(m, l0, int32(627), l1)
							mBase = m.M
							v578 = m.ExcPending
							if v578 != 0 {
								return int32(0)
							} else {
								return v577
							}
						}
					}
				}
			}
		}
	}
}
func F_ParallelWorkerMain(m *base.Module, l0 int32) {
	mBase := m.M
	_ = mBase
	var v2 int32
	_ = v2
	var v17 int32
	_ = v17
	var v18 int32
	_ = v18
	var v19 int32
	_ = v19
	var v22 int32
	_ = v22
	var v25 int32
	_ = v25
	var v27 int32
	_ = v27
	var v29 int32
	_ = v29
	var v39 int32
	_ = v39
	var v51 int32
	_ = v51
	var v58 int32
	_ = v58
	var v59 int32
	_ = v59
	var v61 int64
	_ = v61
	var v63 int64
	_ = v63
	var v71 int32
	_ = v71
	var v74 int32
	_ = v74
	var v75 int32
	_ = v75
	var v79 int32
	_ = v79
	var v84 int32
	_ = v84
	var v85 int32
	_ = v85
	var v87 int32
	_ = v87
	var v88 int32
	_ = v88
	var v90 int32
	_ = v90
	var v92 int64
	_ = v92
	var v94 int32
	_ = v94
	var v100 int32
	_ = v100
	var v101 int32
	_ = v101
	var v104 int32
	_ = v104
	var v107 int32
	_ = v107
	var v111 int32
	_ = v111
	var v114 int32
	_ = v114
	var v115 int32
	_ = v115
	var v117 int32
	_ = v117
	var v120 int32
	_ = v120
	var v122 int32
	_ = v122
	var v124 int32
	_ = v124
	var v125 int32
	_ = v125
	var v126 int32
	_ = v126
	var v128 int32
	_ = v128
	var v129 int32
	_ = v129
	var v130 int32
	_ = v130
	var v135 int32
	_ = v135
	var v137 int32
	_ = v137
	var v138 int32
	_ = v138
	var v140 int32
	_ = v140
	var v141 int32
	_ = v141
	var v144 int32
	_ = v144
	var v146 int32
	_ = v146
	var v151 int32
	_ = v151
	var v153 int32
	_ = v153
	var v154 int32
	_ = v154
	var v155 int32
	_ = v155
	var v157 int32
	_ = v157
	var v160 int32
	_ = v160
	var v163 int32
	_ = v163
	var v164 int32
	_ = v164
	var v170 int32
	_ = v170
	var v173 int32
	_ = v173
	var v178 int32
	_ = v178
	var v181 int32
	_ = v181
	var v182 int64
	_ = v182
	var v184 int64
	_ = v184
	var v190 int32
	_ = v190
	var v191 int32
	_ = v191
	var v192 int32
	_ = v192
	var v195 int32
	_ = v195
	var v196 int32
	_ = v196
	var v199 int32
	_ = v199
	var v202 int32
	_ = v202
	var v205 int32
	_ = v205
	var v206 int32
	_ = v206
	var v209 int32
	_ = v209
	var v210 int32
	_ = v210
	var v213 int32
	_ = v213
	var v220 int32
	_ = v220
	var v221 int32
	_ = v221
	var v225 int32
	_ = v225
	var v228 int32
	_ = v228
	var v231 int32
	_ = v231
	var v234 int32
	_ = v234
	var v235 int32
	_ = v235
	var v238 int32
	_ = v238
	var v239 int32
	_ = v239
	var v242 int32
	_ = v242
	var v249 int32
	_ = v249
	var v250 int32
	_ = v250
	var v255 int32
	_ = v255
	var v256 int32
	_ = v256
	var v259 int32
	_ = v259
	var v262 int32
	_ = v262
	var v265 int32
	_ = v265
	var v266 int32
	_ = v266
	var v269 int32
	_ = v269
	var v270 int32
	_ = v270
	var v273 int32
	_ = v273
	var v280 int32
	_ = v280
	var v281 int32
	_ = v281
	var v286 int32
	_ = v286
	var v287 int32
	_ = v287
	var v290 int32
	_ = v290
	var v293 int32
	_ = v293
	var v296 int32
	_ = v296
	var v297 int32
	_ = v297
	var v300 int32
	_ = v300
	var v301 int32
	_ = v301
	var v304 int32
	_ = v304
	var v311 int32
	_ = v311
	var v312 int32
	_ = v312
	var v317 int32
	_ = v317
	var v318 int32
	_ = v318
	var v321 int32
	_ = v321
	var v324 int32
	_ = v324
	var v327 int32
	_ = v327
	var v328 int32
	_ = v328
	var v331 int32
	_ = v331
	var v332 int32
	_ = v332
	var v335 int32
	_ = v335
	var v342 int32
	_ = v342
	var v343 int32
	_ = v343
	var v348 int32
	_ = v348
	var v349 int32
	_ = v349
	var v352 int32
	_ = v352
	var v355 int32
	_ = v355
	var v358 int32
	_ = v358
	var v359 int32
	_ = v359
	var v362 int32
	_ = v362
	var v363 int32
	_ = v363
	var v366 int32
	_ = v366
	var v373 int32
	_ = v373
	var v374 int32
	_ = v374
	var v377 int32
	_ = v377
	var v380 int32
	_ = v380
	var v381 int32
	_ = v381
	var v382 int32
	_ = v382
	var v384 int32
	_ = v384
	var v387 int32
	_ = v387
	var v389 int32
	_ = v389
	var v390 int32
	_ = v390
	var v392 int32
	_ = v392
	var v393 int32
	_ = v393
	var v394 int32
	_ = v394
	var v396 int32
	_ = v396
	var v397 int32
	_ = v397
	var v398 int32
	_ = v398
	var v401 int32
	_ = v401
	var v403 int32
	_ = v403
	var v404 int32
	_ = v404
	var v405 int32
	_ = v405
	var v406 int32
	_ = v406
	var v411 int32
	_ = v411
	var v412 int32
	_ = v412
	var v414 int32
	_ = v414
	var v415 int32
	_ = v415
	var v416 int32
	_ = v416
	var v432 int32
	_ = v432
	var v433 int32
	_ = v433
	var v434 int32
	_ = v434
	var v437 int32
	_ = v437
	var v438 int32
	_ = v438
	var v456 int32
	_ = v456
	var v459 int32
	_ = v459
	var v460 int32
	_ = v460
	var v462 int32
	_ = v462
	var v464 int32
	_ = v464
	var v467 int32
	_ = v467
	var v470 int64
	_ = v470
	var v473 int32
	_ = v473
	var v474 int64
	_ = v474
	var v477 int32
	_ = v477
	var v479 int32
	_ = v479
	var v490 int32
	_ = v490
	var v491 int32
	_ = v491
	var v492 int32
	_ = v492
	var v494 int32
	_ = v494
	var v496 int32
	_ = v496
	var v497 int32
	_ = v497
	var v514 int32
	_ = v514
	var v520 int32
	_ = v520
	var v526 int32
	_ = v526
	var v527 int32
	_ = v527
	var v529 int32
	_ = v529
	var v531 int32
	_ = v531
	var v532 int32
	_ = v532
	var v533 int32
	_ = v533
	var v535 int32
	_ = v535
	var v559 int32
	_ = v559
	var v560 int32
	_ = v560
	var v562 int32
	_ = v562
	var v564 int32
	_ = v564
	var v566 int32
	_ = v566
	var v568 int32
	_ = v568
	var v574 int32
	_ = v574
	var v578 int32
	_ = v578
	var v583 int32
	_ = v583
	var v585 int32
	_ = v585
	var v594 int32
	_ = v594
	var v595 int32
	_ = v595
	var v596 int32
	_ = v596
	var v598 int32
	_ = v598
	var v601 int32
	_ = v601
	var v603 int32
	_ = v603
	var v604 int32
	_ = v604
	var v607 int32
	_ = v607
	var v609 int32
	_ = v609
	var v615 int32
	_ = v615
	var v617 int32
	_ = v617
	var v618 int32
	_ = v618
	var v636 int32
	_ = v636
	var v637 int32
	_ = v637
	var v638 int32
	_ = v638
	var v641 int32
	_ = v641
	var v642 int32
	_ = v642
	var v664 int32
	_ = v664
	var v665 int32
	_ = v665
	var v669 int32
	_ = v669
	var v670 int32
	_ = v670
	var v671 int32
	_ = v671
	var v672 int32
	_ = v672
	var v678 int32
	_ = v678
	var v695 int32
	_ = v695
	var v696 int32
	_ = v696
	var v697 int32
	_ = v697
	var v698 int32
	_ = v698
	var v699 int32
	_ = v699
	var v702 int32
	_ = v702
	var v707 int32
	_ = v707
	var v711 int32
	_ = v711
	var v716 int32
	_ = v716
	var v735 int32
	_ = v735
	var v736 int32
	_ = v736
	var v737 int32
	_ = v737
	var v738 int32
	_ = v738
	var v739 int32
	_ = v739
	var v742 int32
	_ = v742
	var v744 int32
	_ = v744
	var v745 int32
	_ = v745
	var v751 int32
	_ = v751
	var v755 int32
	_ = v755
	var v760 int32
	_ = v760
	var v762 int32
	_ = v762
	var v764 int64
	_ = v764
	var v766 int32
	_ = v766
	var v769 int32
	_ = v769
	var v770 int32
	_ = v770
	var v771 int32
	_ = v771
	var v772 int32
	_ = v772
	var v773 int32
	_ = v773
	var v774 int32
	_ = v774
	var v777 int32
	_ = v777
	var v781 int32
	_ = v781
	var v782 int32
	_ = v782
	var v783 int32
	_ = v783
	var v784 int32
	_ = v784
	var v787 int32
	_ = v787
	var v790 int32
	_ = v790
	var v791 int32
	_ = v791
	var v793 int32
	_ = v793
	var v794 int32
	_ = v794
	var v795 int32
	_ = v795
	var v797 int32
	_ = v797
	var v798 int32
	_ = v798
	var v800 int32
	_ = v800
	var v802 int32
	_ = v802
	var v803 int32
	_ = v803
	var v807 int32
	_ = v807
	var v808 int32
	_ = v808
	var v811 int32
	_ = v811
	var v813 int32
	_ = v813
	var v818 int32
	_ = v818
	var v820 int32
	_ = v820
	var v825 int32
	_ = v825
	var v826 int32
	_ = v826
	var v829 int32
	_ = v829
	var v830 int32
	_ = v830
	var v831 int32
	_ = v831
	var v832 int32
	_ = v832
	var v833 int32
	_ = v833
	var v834 int32
	_ = v834
	var v835 int32
	_ = v835
	var v836 int32
	_ = v836
	var v838 int32
	_ = v838
	var v840 int32
	_ = v840
	var v842 int32
	_ = v842
	var v845 int32
	_ = v845
	var v846 int32
	_ = v846
	var v847 int32
	_ = v847
	var v849 int32
	_ = v849
	var v852 int32
	_ = v852
	var v853 int32
	_ = v853
	var v860 int32
	_ = v860
	var v876 int32
	_ = v876
	var v879 int32
	_ = v879
	var v883 int32
	_ = v883
	var v884 int32
	_ = v884
	var v888 int32
	_ = v888
	var v889 int32
	_ = v889
	var v891 int32
	_ = v891
	var v892 int32
	_ = v892
	var v894 int32
	_ = v894
	var v895 int32
	_ = v895
	var v897 int32
	_ = v897
	var v900 int32
	_ = v900
	var v901 int32
	_ = v901
	var v904 int32
	_ = v904
	var v906 int32
	_ = v906
	var v909 int32
	_ = v909
	var v911 int32
	_ = v911
	var v914 int32
	_ = v914
	var v916 int32
	_ = v916
	var v917 int32
	_ = v917
	var v919 int32
	_ = v919
	var v920 int32
	_ = v920
	var v923 int32
	_ = v923
	var v924 int32
	_ = v924
	var v927 int32
	_ = v927
	var v928 int32
	_ = v928
	var v931 int32
	_ = v931
	var v933 int32
	_ = v933
	var v936 int32
	_ = v936
	var v938 int32
	_ = v938
	var v940 int32
	_ = v940
	var v942 int32
	_ = v942
	var v943 int32
	_ = v943
	var v944 int32
	_ = v944
	var v946 int32
	_ = v946
	var v951 int32
	_ = v951
	var v956 int32
	_ = v956
	var v959 int32
	_ = v959
	var v963 int32
	_ = v963
	var v967 int32
	_ = v967
	var v974 int32
	_ = v974
	var v977 int32
	_ = v977
	var v981 int32
	_ = v981
	var v986 int32
	_ = v986
	var v1008 int32
	_ = v1008
	var v1011 int32
	_ = v1011
	var v1012 int32
	_ = v1012
	var v1021 int32
	_ = v1021
	var v1024 int32
	_ = v1024
	var v1042 int32
	_ = v1042
	var v1043 int32
	_ = v1043
	var v1044 int32
	_ = v1044
	var v1045 int32
	_ = v1045
	var v1046 int32
	_ = v1046
	var v1047 int32
	_ = v1047
	var v1048 int32
	_ = v1048
	var v1049 int32
	_ = v1049
	var v1050 int32
	_ = v1050
	var v1055 int32
	_ = v1055
	var v1057 int32
	_ = v1057
	var v1058 int32
	_ = v1058
	var v1059 int32
	_ = v1059
	var v1061 int32
	_ = v1061
	var v1064 int32
	_ = v1064
	var v1067 int32
	_ = v1067
	var v1069 int32
	_ = v1069
	var v1070 int32
	_ = v1070
	var v1071 int32
	_ = v1071
	var v1078 int32
	_ = v1078
	var v1080 int32
	_ = v1080
	var v1083 int32
	_ = v1083
	var v1084 int32
	_ = v1084
	var v1087 int32
	_ = v1087
	var v1095 int32
	_ = v1095
	var v1096 int32
	_ = v1096
	var v1097 int32
	_ = v1097
	var v1098 int32
	_ = v1098
	var v1101 int32
	_ = v1101
	var v1102 int32
	_ = v1102
	var v1103 int32
	_ = v1103
	var v1105 int32
	_ = v1105
	var v1114 int32
	_ = v1114
	var v1131 int32
	_ = v1131
	var v1139 int32
	_ = v1139
	var v1142 int32
	_ = v1142
	var v1146 int32
	_ = v1146
	var v1151 int32
	_ = v1151
	var v1152 int32
	_ = v1152
	var v1153 int32
	_ = v1153
	var v1158 int32
	_ = v1158
	var v1160 int32
	_ = v1160
	var v1165 int32
	_ = v1165
	var v1168 int32
	_ = v1168
	var v1172 int32
	_ = v1172
	var v1173 int32
	_ = v1173
	var v1174 int32
	_ = v1174
	var v1176 int32
	_ = v1176
	var v1178 int32
	_ = v1178
	var v1182 int32
	_ = v1182
	var v1188 int32
	_ = v1188
	var v1189 int32
	_ = v1189
	var v1191 int32
	_ = v1191
	var v1208 int32
	_ = v1208
	var v1211 int32
	_ = v1211
	var v1212 int32
	_ = v1212
	var v1214 int32
	_ = v1214
	var v1215 int32
	_ = v1215
	var v1216 int32
	_ = v1216
	var v1232 int32
	_ = v1232
	var v1238 int32
	_ = v1238
	var v1244 int32
	_ = v1244
	var v1245 int32
	_ = v1245
	var v1248 int32
	_ = v1248
	var v1264 int32
	_ = v1264
	var v1267 int32
	_ = v1267
	var v1268 int32
	_ = v1268
	var v1269 int32
	_ = v1269
	var v1293 int32
	_ = v1293
	var v1294 int32
	_ = v1294
	var v1295 int32
	_ = v1295
	var v1297 int32
	_ = v1297
	var v1300 int32
	_ = v1300
	var v1306 int32
	_ = v1306
	var v1309 int32
	_ = v1309
	var v1310 int32
	_ = v1310
	var v1313 int32
	_ = v1313
	var v1315 int32
	_ = v1315
	var v1318 int32
	_ = v1318
	var v1320 int32
	_ = v1320
	var v1321 int32
	_ = v1321
	var v1322 int32
	_ = v1322
	var v1324 int32
	_ = v1324
	var v1333 int32
	_ = v1333
	var v1335 int32
	_ = v1335
	var v1336 int32
	_ = v1336
	var v1342 int32
	_ = v1342
	var v1346 int32
	_ = v1346
	var v1347 int32
	_ = v1347
	var v1352 int32
	_ = v1352
	var v1355 int32
	_ = v1355
	var v1356 int32
	_ = v1356
	var v1361 int32
	_ = v1361
	var v1363 int32
	_ = v1363
	var v1365 int32
	_ = v1365
	var v1369 int32
	_ = v1369
	var v1370 int32
	_ = v1370
	var v1372 int32
	_ = v1372
	var v1373 int32
	_ = v1373
	var v1374 int32
	_ = v1374
	var v1378 int32
	_ = v1378
	var v1379 int32
	_ = v1379
	var v1381 int32
	_ = v1381
	var v1383 int32
	_ = v1383
	var v1384 int32
	_ = v1384
	var v1390 int32
	_ = v1390
	var v1391 int32
	_ = v1391
	var v1392 int32
	_ = v1392
	var v1393 int32
	_ = v1393
	var v1416 int32
	_ = v1416
	var v1419 int32
	_ = v1419
	var v1423 int32
	_ = v1423
	var v1428 int32
	_ = v1428
	var v1432 int32
	_ = v1432
	var v1435 int32
	_ = v1435
	var v1439 int32
	_ = v1439
	var v1444 int32
	_ = v1444
	var v1448 int32
	_ = v1448
	var v1454 int32
	_ = v1454
	var v1459 int32
	_ = v1459
	var v1463 int32
	_ = v1463
	var v1465 int32
	_ = v1465
	var v1466 int32
	_ = v1466
	var v1470 int32
	_ = v1470
	var v1475 int32
	_ = v1475
	var v1484 int32
	_ = v1484
	var v1488 int32
	_ = v1488
	var v1493 int32
	_ = v1493
	v2 = int32(0)
	v17 = m.G0
	v18 = int32(32)
	v19 = v17 - v18
	m.G0 = v19
	v22 = int32(1)
	*(*uint8)(unsafe.Add(mBase, _c_F_ParallelWorkerMain[0])) = uint8(v22)
	v25 = int32(295)
	v27 = m.G0
	v29 = v27 - v18
	m.G0 = v29
	switch int32(297) {
	case 0, 2:
		v39 = v25
		goto L2
	default:
		goto L3
	}
L1:
	;
	F_BackgroundWorkerUnblockSignals(m)
	mBase = m.M
	v71 = m.ExcPending
	if v71 != 0 {
		goto L14
	} else {
		goto L15
	}
L2:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v29)+12)) = v39
	F_sigemptyset(m, v29+int32(16))
	mBase = m.M
	goto L5
L3:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_ParallelWorkerMain[1])) = v25
	v39 = int32(_a_F_ParallelWorkerMain_0)
	goto L2
L5:
	;
	goto L6
L6:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v29)+24)) = int32(268435456)
	v51 = v29 + int32(12)
	goto L9
L7:
	;
	m.G0 = v29 + int32(32)
	goto L1
L9:
	;
	goto L10
L10:
	;
	if v51 != 0 {
		goto L11
	} else {
		goto L12
	}
L11:
	;
	v58 = int32(300)
	v59 = *(*int32)(unsafe.Add(mBase, uint32(v51)+16))
	*(*int32)(unsafe.Add(mBase, _c_F_ParallelWorkerMain[2])) = v59
	v61 = *(*int64)(unsafe.Add(mBase, uint32(v51)+8))
	*(*int64)(unsafe.Add(mBase, _c_F_ParallelWorkerMain[3])) = v61
	v63 = *(*int64)(unsafe.Add(mBase, uint32(v51)))
	*(*int64)(unsafe.Add(mBase, _c_F_ParallelWorkerMain[4])) = v63
	goto L13
L12:
	;
	goto L13
L13:
	;
	goto L7
L14:
	;
	return
L15:
	;
	v74 = *(*int32)(unsafe.Add(mBase, _c_F_ParallelWorkerMain[5]))
	v75 = *(*int32)(unsafe.Add(mBase, uint32(v74)+1328))
	*(*int32)(unsafe.Add(mBase, _c_F_ParallelWorkerMain[6])) = v75
	v79 = *(*int32)(unsafe.Add(mBase, _c_F_ParallelWorkerMain[7]))
	v84 = F_AllocSetContextCreateInternal(m, v79, int32(_a_F_ParallelWorkerMain_1), int32(0), int32(_a_F_ParallelWorkerMain_2), int32(_a_F_ParallelWorkerMain_3))
	mBase = m.M
	v85 = m.ExcPending
	if v85 != 0 {
		goto L14
	} else {
		goto L16
	}
L16:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_ParallelWorkerMain[8])) = v84
	v87 = F_dsm_attach(m, l0)
	mBase = m.M
	v88 = m.ExcPending
	if v88 != 0 {
		goto L14
	} else {
		goto L21
	}
L17:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v1484 = m.ExcPending
	if v1484 != 0 {
		goto L14
	} else {
		goto L363
	}
L18:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v1463 = m.ExcPending
	if v1463 != 0 {
		goto L14
	} else {
		goto L359
	}
L19:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v1448 = m.ExcPending
	if v1448 != 0 {
		goto L14
	} else {
		goto L356
	}
L20:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v1432 = m.ExcPending
	if v1432 != 0 {
		goto L14
	} else {
		goto L352
	}
L21:
	;
	if v87 != 0 {
		goto L22
	} else {
		goto L23
	}
L22:
	;
	v90 = *(*int32)(unsafe.Add(mBase, uint32(v87)+24))
	v92 = *(*int64)(unsafe.Add(mBase, uint32(v90)))
	if v92 == int64(1346862204) {
		goto L26
	} else {
		goto L27
	}
L23:
	;
	goto L24
L24:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v1416 = m.ExcPending
	if v1416 != 0 {
		goto L14
	} else {
		goto L348
	}
L25:
	;
	if v94 == int32(0) {
		goto L20
	} else {
		goto L29
	}
L26:
	;
	v94 = v90
	goto L28
L27:
	;
	v94 = int32(0)
	goto L28
L28:
	;
	goto L25
L29:
	;
	v100 = F_shm_toc_lookup(m, v94, int64(-65535), int32(0))
	mBase = m.M
	v101 = m.ExcPending
	if v101 != 0 {
		goto L14
	} else {
		goto L30
	}
L30:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_ParallelWorkerMain[9])) = v100
	v104 = *(*int32)(unsafe.Add(mBase, uint32(v100)+40))
	*(*int32)(unsafe.Add(mBase, _c_F_ParallelWorkerMain[10])) = v104
	v107 = *(*int32)(unsafe.Add(mBase, uint32(v100)+44))
	*(*int32)(unsafe.Add(mBase, _c_F_ParallelWorkerMain[11])) = v107
	F_before_shmem_exit(m, int32(296), v87)
	mBase = m.M
	v111 = m.ExcPending
	if v111 != 0 {
		goto L14
	} else {
		goto L31
	}
L31:
	;
	v114 = F_shm_toc_lookup(m, v94, int64(-65534), int32(0))
	mBase = m.M
	v115 = m.ExcPending
	if v115 != 0 {
		goto L14
	} else {
		goto L32
	}
L32:
	;
	v117 = *(*int32)(unsafe.Add(mBase, _c_F_ParallelWorkerMain[6]))
	v120 = v114 + v117<<(uint(int32(14))%32)
	v122 = *(*int32)(unsafe.Add(mBase, _c_F_ParallelWorkerMain[12]))
	F_shm_mq_set_sender(m, v120, v122)
	mBase = m.M
	v124 = m.ExcPending
	if v124 != 0 {
		goto L14
	} else {
		goto L33
	}
L33:
	;
	v125 = F_shm_mq_attach(m, v120, v87)
	mBase = m.M
	v126 = m.ExcPending
	if v126 != 0 {
		goto L14
	} else {
		goto L34
	}
L34:
	;
	F_pq_redirect_to_shm_mq(m, v87, v125)
	mBase = m.M
	v128 = m.ExcPending
	if v128 != 0 {
		goto L14
	} else {
		goto L35
	}
L35:
	;
	v129 = *(*int32)(unsafe.Add(mBase, uint32(v100)+40))
	v130 = *(*int32)(unsafe.Add(mBase, uint32(v100)+44))
	*(*int32)(unsafe.Add(mBase, _c_F_ParallelWorkerMain[13])) = v130
	*(*int32)(unsafe.Add(mBase, _c_F_ParallelWorkerMain[14])) = v129
	goto L36
L36:
	;
	v135 = *(*int32)(unsafe.Add(mBase, uint32(v100)+40))
	v137 = *(*int32)(unsafe.Add(mBase, _c_F_ParallelWorkerMain[15]))
	v138 = *(*int32)(unsafe.Add(mBase, uint32(v100)+36))
	v140 = *(*int32)(unsafe.Add(mBase, _c_F_ParallelWorkerMain[16]))
	v141 = *(*int32)(unsafe.Add(mBase, uint32(v140)))
	v144 = base.I32_div_s(v138-v141, int32(640))
	v146 = base.I32_rem_s(v144, int32(16))
	v151 = v137 + v146<<(uint(int32(7))%32) + int32(_a_F_ParallelWorkerMain_4)
	v153 = F_LWLockAcquire(m, v151, int32(0))
	mBase = m.M
	v154 = m.ExcPending
	if v154 != 0 {
		goto L14
	} else {
		goto L37
	}
L37:
	;
	v155 = *(*int32)(unsafe.Add(mBase, uint32(v138)+44))
	if v155 != v135 {
		v178 = v2
		goto L38
	} else {
		goto L39
	}
L38:
	;
	F_LWLockRelease(m, v151)
	mBase = m.M
	v181 = m.ExcPending
	if v181 != 0 {
		goto L14
	} else {
		goto L44
	}
L39:
	;
	v157 = *(*int32)(unsafe.Add(mBase, uint32(v138)+616))
	if v157 != v138 {
		v178 = v2
		goto L38
	} else {
		goto L40
	}
L40:
	;
	v160 = *(*int32)(unsafe.Add(mBase, _c_F_ParallelWorkerMain[12]))
	*(*int32)(unsafe.Add(mBase, uint32(v160)+616)) = v138
	v163 = v138 + int32(620)
	v164 = *(*int32)(unsafe.Add(mBase, uint32(v138)+624))
	if v164 == int32(0) {
		goto L41
	} else {
		goto L42
	}
L41:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v138)+624)) = v163
	*(*int32)(unsafe.Add(mBase, uint32(v138)+620)) = v163
	goto L43
L42:
	;
	goto L43
L43:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v160)+632)) = v163
	v170 = *(*int32)(unsafe.Add(mBase, uint32(v163)))
	*(*int32)(unsafe.Add(mBase, uint32(v160)+628)) = v170
	v173 = v160 + int32(628)
	*(*int32)(unsafe.Add(mBase, uint32(v170)+4)) = v173
	*(*int32)(unsafe.Add(mBase, uint32(v163))) = v173
	v178 = int32(1)
	goto L38
L44:
	;
	if v178 != 0 {
		goto L45
	} else {
		goto L46
	}
L45:
	;
	v182 = *(*int64)(unsafe.Add(mBase, uint32(v100)+48))
	v184 = *(*int64)(unsafe.Add(mBase, uint32(v100)+56))
	*(*int64)(unsafe.Add(mBase, _c_F_ParallelWorkerMain[17])) = v184
	*(*int64)(unsafe.Add(mBase, _c_F_ParallelWorkerMain[18])) = v182
	v190 = F_shm_toc_lookup(m, v94, int64(-65527), int32(0))
	mBase = m.M
	v191 = m.ExcPending
	if v191 != 0 {
		goto L14
	} else {
		goto L48
	}
L46:
	;
	goto L47
L47:
	;
	m.G0 = v19 + int32(32)
	return
L48:
	;
	v192 = F_strlen(m, v190)
	mBase = m.M
	v195 = v192 + v190 + int32(1)
	v196 = int32(_a_F_ParallelWorkerMain_5)
	v199 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v190))))
	v202 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_ParallelWorkerMain[19])))
	if base.B2i32(v199 == int32(0))|base.B2i32(v199 != v202) != 0 {
		v220 = v199
		v221 = v202
		goto L51
	} else {
		goto L52
	}
L49:
	;
	v384 = *(*int32)(unsafe.Add(mBase, uint32(v100)+4))
	*(*int32)(unsafe.Add(mBase, _c_F_ParallelWorkerMain[20])) = v384
	v387 = *(*int32)(unsafe.Add(mBase, _c_F_ParallelWorkerMain[12]))
	*(*int32)(unsafe.Add(mBase, uint32(v387)+64)) = v384
	v389 = *(*int32)(unsafe.Add(mBase, uint32(v100)+8))
	v390 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v100)+32)))
	F_SetSessionAuthorization(m, v389, v390)
	mBase = m.M
	v392 = m.ExcPending
	if v392 != 0 {
		goto L14
	} else {
		goto L109
	}
L50:
	;
	if v220-v221 == int32(0) {
		goto L57
	} else {
		goto L58
	}
L51:
	;
	goto L50
L52:
	;
	v205 = v190
	v206 = v196
	goto L53
L53:
	;
	v209 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v206)+1)))
	v210 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v205)+1)))
	if v210 == int32(0) {
		v220 = v210
		v221 = v209
		goto L51
	} else {
		goto L55
	}
L54:
	;
	v220 = v210
	v221 = v209
	goto L51
L55:
	;
	v213 = int32(1)
	if v210 == v209 {
		v205 = v205 + v213
		v206 = v206 + v213
		goto L53
	} else {
		goto L56
	}
L56:
	;
	goto L54
L57:
	;
	v225 = int32(_a_F_ParallelWorkerMain_6)
	v228 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_ParallelWorkerMain[21])))
	v231 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v195))))
	if base.B2i32(v228 == int32(0))|base.B2i32(v228 != v231) != 0 {
		v249 = v228
		v250 = v231
		goto L61
	} else {
		goto L62
	}
L58:
	;
	goto L59
L59:
	;
	v380 = F_load_external_function(m, v190, v195, int32(1), int32(0))
	mBase = m.M
	v381 = m.ExcPending
	if v381 != 0 {
		goto L14
	} else {
		goto L108
	}
L60:
	;
	if v249-v250 == int32(0) {
		goto L67
	} else {
		goto L68
	}
L61:
	;
	goto L60
L62:
	;
	v234 = v225
	v235 = v195
	goto L63
L63:
	;
	v238 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v235)+1)))
	v239 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v234)+1)))
	if v239 == int32(0) {
		v249 = v239
		v250 = v238
		goto L61
	} else {
		goto L65
	}
L64:
	;
	v249 = v239
	v250 = v238
	goto L61
L65:
	;
	v242 = int32(1)
	if v239 == v238 {
		v234 = v234 + v242
		v235 = v235 + v242
		goto L63
	} else {
		goto L66
	}
L66:
	;
	goto L64
L67:
	;
	v255 = *(*int32)(unsafe.Add(mBase, _c_F_ParallelWorkerMain[22]))
	v382 = v255
	goto L49
L68:
	;
	goto L69
L69:
	;
	v256 = int32(_a_F_ParallelWorkerMain_7)
	v259 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_ParallelWorkerMain[23])))
	v262 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v195))))
	if base.B2i32(v259 == int32(0))|base.B2i32(v259 != v262) != 0 {
		v280 = v259
		v281 = v262
		goto L71
	} else {
		goto L72
	}
L70:
	;
	if v280-v281 == int32(0) {
		goto L77
	} else {
		goto L78
	}
L71:
	;
	goto L70
L72:
	;
	v265 = v256
	v266 = v195
	goto L73
L73:
	;
	v269 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v266)+1)))
	v270 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v265)+1)))
	if v270 == int32(0) {
		v280 = v270
		v281 = v269
		goto L71
	} else {
		goto L75
	}
L74:
	;
	v280 = v270
	v281 = v269
	goto L71
L75:
	;
	v273 = int32(1)
	if v270 == v269 {
		v265 = v265 + v273
		v266 = v266 + v273
		goto L73
	} else {
		goto L76
	}
L76:
	;
	goto L74
L77:
	;
	v286 = *(*int32)(unsafe.Add(mBase, _c_F_ParallelWorkerMain[24]))
	v382 = v286
	goto L49
L78:
	;
	goto L79
L79:
	;
	v287 = int32(_a_F_ParallelWorkerMain_8)
	v290 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_ParallelWorkerMain[25])))
	v293 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v195))))
	if base.B2i32(v290 == int32(0))|base.B2i32(v290 != v293) != 0 {
		v311 = v290
		v312 = v293
		goto L81
	} else {
		goto L82
	}
L80:
	;
	if v311-v312 == int32(0) {
		goto L87
	} else {
		goto L88
	}
L81:
	;
	goto L80
L82:
	;
	v296 = v287
	v297 = v195
	goto L83
L83:
	;
	v300 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v297)+1)))
	v301 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v296)+1)))
	if v301 == int32(0) {
		v311 = v301
		v312 = v300
		goto L81
	} else {
		goto L85
	}
L84:
	;
	v311 = v301
	v312 = v300
	goto L81
L85:
	;
	v304 = int32(1)
	if v301 == v300 {
		v296 = v296 + v304
		v297 = v297 + v304
		goto L83
	} else {
		goto L86
	}
L86:
	;
	goto L84
L87:
	;
	v317 = *(*int32)(unsafe.Add(mBase, _c_F_ParallelWorkerMain[26]))
	v382 = v317
	goto L49
L88:
	;
	goto L89
L89:
	;
	v318 = int32(_a_F_ParallelWorkerMain_9)
	v321 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_ParallelWorkerMain[27])))
	v324 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v195))))
	if base.B2i32(v321 == int32(0))|base.B2i32(v321 != v324) != 0 {
		v342 = v321
		v343 = v324
		goto L91
	} else {
		goto L92
	}
L90:
	;
	if v342-v343 == int32(0) {
		goto L97
	} else {
		goto L98
	}
L91:
	;
	goto L90
L92:
	;
	v327 = v318
	v328 = v195
	goto L93
L93:
	;
	v331 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v328)+1)))
	v332 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v327)+1)))
	if v332 == int32(0) {
		v342 = v332
		v343 = v331
		goto L91
	} else {
		goto L95
	}
L94:
	;
	v342 = v332
	v343 = v331
	goto L91
L95:
	;
	v335 = int32(1)
	if v332 == v331 {
		v327 = v327 + v335
		v328 = v328 + v335
		goto L93
	} else {
		goto L96
	}
L96:
	;
	goto L94
L97:
	;
	v348 = *(*int32)(unsafe.Add(mBase, _c_F_ParallelWorkerMain[28]))
	v382 = v348
	goto L49
L98:
	;
	goto L99
L99:
	;
	v349 = int32(_a_F_ParallelWorkerMain_10)
	v352 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_ParallelWorkerMain[29])))
	v355 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v195))))
	if base.B2i32(v352 == int32(0))|base.B2i32(v352 != v355) != 0 {
		v373 = v352
		v374 = v355
		goto L101
	} else {
		goto L102
	}
L100:
	;
	if v373-v374 != 0 {
		goto L19
	} else {
		goto L107
	}
L101:
	;
	goto L100
L102:
	;
	v358 = v349
	v359 = v195
	goto L103
L103:
	;
	v362 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v359)+1)))
	v363 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v358)+1)))
	if v363 == int32(0) {
		v373 = v363
		v374 = v362
		goto L101
	} else {
		goto L105
	}
L104:
	;
	v373 = v363
	v374 = v362
	goto L101
L105:
	;
	v366 = int32(1)
	if v363 == v362 {
		v358 = v358 + v366
		v359 = v359 + v366
		goto L103
	} else {
		goto L106
	}
L106:
	;
	goto L104
L107:
	;
	v377 = *(*int32)(unsafe.Add(mBase, _c_F_ParallelWorkerMain[30]))
	v382 = v377
	goto L49
L108:
	;
	v382 = v380
	goto L49
L109:
	;
	v393 = *(*int32)(unsafe.Add(mBase, uint32(v100)+12))
	v394 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v100)+33)))
	F_SetCurrentRoleId(m, v393, v394)
	mBase = m.M
	v396 = m.ExcPending
	if v396 != 0 {
		goto L14
	} else {
		goto L110
	}
L110:
	;
	v397 = *(*int32)(unsafe.Add(mBase, uint32(v100)))
	v398 = *(*int32)(unsafe.Add(mBase, uint32(v100)+4))
	F_BackgroundWorkerInitializeConnectionByOid(m, v397, v398, int32(3))
	mBase = m.M
	v401 = m.ExcPending
	if v401 != 0 {
		goto L14
	} else {
		goto L111
	}
L111:
	;
	v403 = *(*int32)(unsafe.Add(mBase, _c_F_ParallelWorkerMain[31]))
	v404 = *(*int32)(unsafe.Add(mBase, uint32(v403)+4))
	goto L112
L112:
	;
	v405 = F_SetClientEncoding(m, v404)
	mBase = m.M
	v406 = m.ExcPending
	if v406 != 0 {
		goto L14
	} else {
		goto L113
	}
L113:
	;
	if v405 < int32(0) {
		goto L18
	} else {
		goto L114
	}
L114:
	;
	v411 = F_shm_toc_lookup(m, v94, int64(-65533), int32(0))
	mBase = m.M
	v412 = m.ExcPending
	if v412 != 0 {
		goto L14
	} else {
		goto L115
	}
L115:
	;
	F_StartTransactionCommand(m)
	mBase = m.M
	v414 = m.ExcPending
	if v414 != 0 {
		goto L14
	} else {
		goto L116
	}
L116:
	;
	v415 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v411))))
	if v415 != 0 {
		goto L117
	} else {
		goto L118
	}
L117:
	;
	v416 = v411
	goto L120
L118:
	;
	goto L119
L119:
	;
	F_CommitTransactionCommand(m)
	mBase = m.M
	v456 = m.ExcPending
	if v456 != 0 {
		goto L14
	} else {
		goto L124
	}
L120:
	;
	v432 = F_internal_load_library(m, v416)
	mBase = m.M
	v433 = m.ExcPending
	if v433 != 0 {
		goto L14
	} else {
		goto L122
	}
L121:
	;
	goto L119
L122:
	;
	v434 = F_strlen(m, v416)
	mBase = m.M
	v437 = v434 + v416 + int32(1)
	v438 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v437))))
	if v438 != 0 {
		v416 = v437
		goto L120
	} else {
		goto L123
	}
L123:
	;
	goto L121
L124:
	;
	v459 = F_shm_toc_lookup(m, v94, int64(-65528), int32(0))
	mBase = m.M
	v460 = m.ExcPending
	if v460 != 0 {
		goto L14
	} else {
		goto L125
	}
L125:
	;
	F_StartTransaction(m)
	mBase = m.M
	v462 = m.ExcPending
	if v462 != 0 {
		goto L14
	} else {
		goto L126
	}
L126:
	;
	v464 = *(*int32)(unsafe.Add(mBase, uint32(v459)))
	*(*int32)(unsafe.Add(mBase, _c_F_ParallelWorkerMain[32])) = v464
	v467 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v459)+4)))
	*(*uint8)(unsafe.Add(mBase, _c_F_ParallelWorkerMain[33])) = uint8(v467)
	v470 = *(*int64)(unsafe.Add(mBase, uint32(v459)+8))
	*(*int64)(unsafe.Add(mBase, _c_F_ParallelWorkerMain[34])) = v470
	v473 = *(*int32)(unsafe.Add(mBase, _c_F_ParallelWorkerMain[35]))
	v474 = *(*int64)(unsafe.Add(mBase, uint32(v459)+16))
	*(*int64)(unsafe.Add(mBase, uint32(v473))) = v474
	v477 = *(*int32)(unsafe.Add(mBase, uint32(v459)+24))
	*(*int32)(unsafe.Add(mBase, _c_F_ParallelWorkerMain[36])) = v477
	v479 = *(*int32)(unsafe.Add(mBase, uint32(v459)+28))
	*(*int32)(unsafe.Add(mBase, _c_F_ParallelWorkerMain[37])) = v459 + int32(32)
	*(*int32)(unsafe.Add(mBase, _c_F_ParallelWorkerMain[38])) = v479
	*(*int32)(unsafe.Add(mBase, uint32(v473)+24)) = int32(5)
	v490 = F_shm_toc_lookup(m, v94, int64(-65525), int32(0))
	mBase = m.M
	v491 = m.ExcPending
	if v491 != 0 {
		goto L14
	} else {
		goto L127
	}
L127:
	;
	v492 = m.G0
	v494 = v492 - int32(48)
	m.G0 = v494
	v496 = *(*int32)(unsafe.Add(mBase, uint32(v490)+8))
	if v496 != 0 {
		goto L128
	} else {
		goto L129
	}
L128:
	;
	v497 = v490
	goto L131
L129:
	;
	goto L130
L130:
	;
	m.G0 = v494 + int32(48)
	v559 = F_shm_toc_lookup(m, v94, int64(-65523), int32(0))
	mBase = m.M
	v560 = m.ExcPending
	if v560 != 0 {
		goto L14
	} else {
		goto L139
	}
L131:
	;
	v514 = *(*int32)(unsafe.Add(mBase, _c_F_ParallelWorkerMain[39]))
	if v514 == int32(0) {
		goto L133
	} else {
		goto L134
	}
L132:
	;
	goto L130
L133:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v494)+16)) = int64(68719476748)
	v520 = *(*int32)(unsafe.Add(mBase, _c_F_ParallelWorkerMain[40]))
	*(*int32)(unsafe.Add(mBase, uint32(v494)+40)) = v520
	v526 = F_hash_create(m, int32(_a_F_ParallelWorkerMain_11), int32(16), v494, int32(1064))
	mBase = m.M
	v527 = m.ExcPending
	if v527 != 0 {
		goto L14
	} else {
		goto L136
	}
L134:
	;
	v529 = v514
	goto L135
L135:
	;
	v531 = F_hash_search(m, v529, v497, int32(1), v494)
	mBase = m.M
	v532 = m.ExcPending
	if v532 != 0 {
		goto L14
	} else {
		goto L137
	}
L136:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_ParallelWorkerMain[39])) = v526
	v529 = v526
	goto L135
L137:
	;
	v533 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v531)+12)) = uint8(v533)
	v535 = *(*int32)(unsafe.Add(mBase, uint32(v497)+20))
	if v535 != 0 {
		v497 = v497 + int32(12)
		goto L131
	} else {
		goto L138
	}
L138:
	;
	goto L132
L139:
	;
	v562 = *(*int32)(unsafe.Add(mBase, _c_F_ParallelWorkerMain[41]))
	if v562 != 0 {
		goto L141
	} else {
		goto L142
	}
L140:
	;
	v585 = int32(524)
	base.MemoryCopy(m, int32(_a_F_ParallelWorkerMain_12), v559, v585)
	base.MemoryCopy(m, int32(_a_F_ParallelWorkerMain_13), v559+v585, v585)
	v594 = F_shm_toc_lookup(m, v94, int64(-65524), int32(0))
	mBase = m.M
	v595 = m.ExcPending
	if v595 != 0 {
		goto L14
	} else {
		goto L149
	}
L141:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v574 = m.ExcPending
	if v574 != 0 {
		goto L14
	} else {
		goto L146
	}
L142:
	;
	v564 = *(*int32)(unsafe.Add(mBase, _c_F_ParallelWorkerMain[42]))
	if v564 != 0 {
		goto L141
	} else {
		goto L143
	}
L143:
	;
	v566 = *(*int32)(unsafe.Add(mBase, _c_F_ParallelWorkerMain[43]))
	if v566 != 0 {
		goto L141
	} else {
		goto L144
	}
L144:
	;
	v568 = *(*int32)(unsafe.Add(mBase, _c_F_ParallelWorkerMain[44]))
	if v568 == int32(0) {
		goto L140
	} else {
		goto L145
	}
L145:
	;
	goto L141
L146:
	;
	F_errmsg_internal(m, int32(_a_F_ParallelWorkerMain_14), int32(0))
	mBase = m.M
	v578 = m.ExcPending
	if v578 != 0 {
		goto L14
	} else {
		goto L147
	}
L147:
	;
	F_errfinish(m, int32(_a_F_ParallelWorkerMain_15), int32(749), int32(_a_F_ParallelWorkerMain_16))
	mBase = m.M
	v583 = m.ExcPending
	if v583 != 0 {
		goto L14
	} else {
		goto L148
	}
L148:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L149:
	;
	v596 = int32(0)
	v598 = *(*int32)(unsafe.Add(mBase, uint32(v594)))
	*(*int32)(unsafe.Add(mBase, _c_F_ParallelWorkerMain[45])) = v598
	v601 = *(*int32)(unsafe.Add(mBase, uint32(v594)+4))
	*(*int32)(unsafe.Add(mBase, _c_F_ParallelWorkerMain[46])) = v601
	v603 = int32(_a_F_ParallelWorkerMain_17)
	v604 = *(*int32)(unsafe.Add(mBase, _c_F_ParallelWorkerMain[8]))
	v607 = *(*int32)(unsafe.Add(mBase, _c_F_ParallelWorkerMain[7]))
	*(*int32)(unsafe.Add(mBase, _c_F_ParallelWorkerMain[8])) = v607
	v609 = *(*int32)(unsafe.Add(mBase, uint32(v594)+8))
	if v596 < v609 {
		goto L150
	} else {
		goto L151
	}
L150:
	;
	v615 = *(*int32)(unsafe.Add(mBase, _c_F_ParallelWorkerMain[47]))
	v617 = v596
	v618 = v615
	goto L153
L151:
	;
	goto L152
L152:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_ParallelWorkerMain[8])) = v604
	v664 = *(*int32)(unsafe.Add(mBase, _c_F_ParallelWorkerMain[35]))
	v665 = *(*int32)(unsafe.Add(mBase, uint32(v664)+28))
	goto L157
L153:
	;
	v636 = *(*int32)(unsafe.Add(mBase, uint32(v594+int32(12)+v617<<(uint(int32(2))%32))))
	v637 = F_lappend_oid(m, v618, v636)
	mBase = m.M
	v638 = m.ExcPending
	if v638 != 0 {
		goto L14
	} else {
		goto L155
	}
L154:
	;
	goto L152
L155:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_ParallelWorkerMain[47])) = v637
	v641 = v617 + int32(1)
	v642 = *(*int32)(unsafe.Add(mBase, uint32(v594)+8))
	if v641 < v642 {
		v617 = v641
		v618 = v637
		goto L153
	} else {
		goto L156
	}
L156:
	;
	goto L154
L157:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_ParallelWorkerMain[48])) = v665
	v669 = F_shm_toc_lookup(m, v94, int64(-65531), int32(0))
	mBase = m.M
	v670 = m.ExcPending
	if v670 != 0 {
		goto L14
	} else {
		goto L158
	}
L158:
	;
	v671 = int32(0)
	v672 = *(*int32)(unsafe.Add(mBase, uint32(v669)))
	if v672 <= v671 {
		goto L159
	} else {
		goto L160
	}
L159:
	;
	v735 = F_shm_toc_lookup(m, v94, int64(-65526), int32(0))
	mBase = m.M
	v736 = m.ExcPending
	if v736 != 0 {
		goto L14
	} else {
		goto L171
	}
L160:
	;
	v678 = v671
	goto L161
L161:
	;
	v695 = v669 + int32(4) + v678<<(uint(int32(3))%32)
	v696 = *(*int32)(unsafe.Add(mBase, uint32(v695)))
	v697 = *(*int32)(unsafe.Add(mBase, uint32(v695)+4))
	v698 = F_GetComboCommandId(m, v696, v697)
	mBase = m.M
	v699 = m.ExcPending
	if v699 != 0 {
		goto L14
	} else {
		goto L163
	}
L162:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v707 = m.ExcPending
	if v707 != 0 {
		goto L14
	} else {
		goto L168
	}
L163:
	;
	if v698 == v678 {
		goto L164
	} else {
		goto L165
	}
L164:
	;
	v702 = v678 + int32(1)
	if v672 != v702 {
		v678 = v702
		goto L161
	} else {
		goto L167
	}
L165:
	;
	goto L166
L166:
	;
	goto L162
L167:
	;
	goto L159
L168:
	;
	F_errmsg_internal(m, int32(_a_F_ParallelWorkerMain_18), int32(0))
	mBase = m.M
	v711 = m.ExcPending
	if v711 != 0 {
		goto L14
	} else {
		goto L169
	}
L169:
	;
	F_errfinish(m, int32(_a_F_ParallelWorkerMain_19), int32(362), int32(_a_F_ParallelWorkerMain_20))
	mBase = m.M
	v716 = m.ExcPending
	if v716 != 0 {
		goto L14
	} else {
		goto L170
	}
L170:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L171:
	;
	v737 = *(*int32)(unsafe.Add(mBase, uint32(v735)))
	v738 = int32(_a_F_ParallelWorkerMain_17)
	v739 = *(*int32)(unsafe.Add(mBase, _c_F_ParallelWorkerMain[8]))
	v742 = *(*int32)(unsafe.Add(mBase, _c_F_ParallelWorkerMain[7]))
	*(*int32)(unsafe.Add(mBase, _c_F_ParallelWorkerMain[8])) = v742
	v744 = F_dsm_attach(m, v737)
	mBase = m.M
	v745 = m.ExcPending
	if v745 != 0 {
		goto L14
	} else {
		goto L172
	}
L172:
	;
	if v744 == int32(0) {
		goto L173
	} else {
		goto L174
	}
L173:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v751 = m.ExcPending
	if v751 != 0 {
		goto L14
	} else {
		goto L176
	}
L174:
	;
	goto L175
L175:
	;
	v762 = *(*int32)(unsafe.Add(mBase, uint32(v744)+24))
	v764 = *(*int64)(unsafe.Add(mBase, uint32(v762)))
	if v764 == int64(2880502729) {
		goto L180
	} else {
		goto L181
	}
L176:
	;
	F_errmsg_internal(m, int32(_a_F_ParallelWorkerMain_21), int32(0))
	mBase = m.M
	v755 = m.ExcPending
	if v755 != 0 {
		goto L14
	} else {
		goto L177
	}
L177:
	;
	F_errfinish(m, int32(_a_F_ParallelWorkerMain_22), int32(169), int32(_a_F_ParallelWorkerMain_23))
	mBase = m.M
	v760 = m.ExcPending
	if v760 != 0 {
		goto L14
	} else {
		goto L178
	}
L178:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L179:
	;
	v769 = F_shm_toc_lookup(m, v766, int64(-65535), int32(0))
	mBase = m.M
	v770 = m.ExcPending
	if v770 != 0 {
		goto L14
	} else {
		goto L183
	}
L180:
	;
	v766 = v762
	goto L182
L181:
	;
	v766 = int32(0)
	goto L182
L182:
	;
	goto L179
L183:
	;
	v771 = F_dsa_attach_in_place(m, v769, v744)
	mBase = m.M
	v772 = m.ExcPending
	if v772 != 0 {
		goto L14
	} else {
		goto L184
	}
L184:
	;
	v773 = int32(_a_F_ParallelWorkerMain_24)
	v774 = *(*int32)(unsafe.Add(mBase, _c_F_ParallelWorkerMain[49]))
	*(*int32)(unsafe.Add(mBase, uint32(v774))) = v744
	v777 = *(*int32)(unsafe.Add(mBase, _c_F_ParallelWorkerMain[49]))
	*(*int32)(unsafe.Add(mBase, uint32(v777)+4)) = v771
	v781 = F_shm_toc_lookup(m, v766, int64(-65534), int32(0))
	mBase = m.M
	v782 = m.ExcPending
	if v782 != 0 {
		goto L14
	} else {
		goto L185
	}
L185:
	;
	v783 = int32(_a_F_ParallelWorkerMain_17)
	v784 = *(*int32)(unsafe.Add(mBase, _c_F_ParallelWorkerMain[8]))
	v787 = *(*int32)(unsafe.Add(mBase, _c_F_ParallelWorkerMain[7]))
	*(*int32)(unsafe.Add(mBase, _c_F_ParallelWorkerMain[8])) = v787
	v790 = *(*int32)(unsafe.Add(mBase, _c_F_ParallelWorkerMain[49]))
	v791 = *(*int32)(unsafe.Add(mBase, uint32(v790)+4))
	v793 = *(*int32)(unsafe.Add(mBase, uint32(v781)))
	v794 = F_dshash_attach(m, v791, int32(_a_F_ParallelWorkerMain_25), v793, v791)
	mBase = m.M
	v795 = m.ExcPending
	if v795 != 0 {
		goto L14
	} else {
		goto L186
	}
L186:
	;
	v797 = *(*int32)(unsafe.Add(mBase, _c_F_ParallelWorkerMain[49]))
	v798 = *(*int32)(unsafe.Add(mBase, uint32(v797)+4))
	v800 = *(*int32)(unsafe.Add(mBase, uint32(v781)+4))
	v802 = F_dshash_attach(m, v798, int32(_a_F_ParallelWorkerMain_26), v800, int32(0))
	mBase = m.M
	v803 = m.ExcPending
	if v803 != 0 {
		goto L14
	} else {
		goto L187
	}
L187:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_ParallelWorkerMain[8])) = v784
	v807 = *(*int32)(unsafe.Add(mBase, _c_F_ParallelWorkerMain[49]))
	v808 = *(*int32)(unsafe.Add(mBase, uint32(v807)))
	F_on_dsm_detach(m, v808, int32(1607), v781)
	mBase = m.M
	v811 = m.ExcPending
	if v811 != 0 {
		goto L14
	} else {
		goto L188
	}
L188:
	;
	v813 = *(*int32)(unsafe.Add(mBase, _c_F_ParallelWorkerMain[49]))
	*(*int32)(unsafe.Add(mBase, uint32(v813)+16)) = v802
	*(*int32)(unsafe.Add(mBase, uint32(v813)+12)) = v794
	*(*int32)(unsafe.Add(mBase, uint32(v813)+8)) = v781
	F_dsm_pin_mapping(m, v744)
	mBase = m.M
	v818 = m.ExcPending
	if v818 != 0 {
		goto L14
	} else {
		goto L189
	}
L189:
	;
	F_dsa_pin_mapping(m, v771)
	mBase = m.M
	v820 = m.ExcPending
	if v820 != 0 {
		goto L14
	} else {
		goto L190
	}
L190:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_ParallelWorkerMain[8])) = v739
	v825 = F_shm_toc_lookup(m, v94, int64(-65529), int32(0))
	mBase = m.M
	v826 = m.ExcPending
	if v826 != 0 {
		goto L14
	} else {
		goto L191
	}
L191:
	;
	v829 = F_shm_toc_lookup(m, v94, int64(-65530), int32(1))
	mBase = m.M
	v830 = m.ExcPending
	if v830 != 0 {
		goto L14
	} else {
		goto L192
	}
L192:
	;
	v831 = F_RestoreSnapshot(m, v825)
	mBase = m.M
	v832 = m.ExcPending
	if v832 != 0 {
		goto L14
	} else {
		goto L193
	}
L193:
	;
	if v829 != 0 {
		goto L194
	} else {
		goto L195
	}
L194:
	;
	v833 = F_RestoreSnapshot(m, v829)
	mBase = m.M
	v834 = m.ExcPending
	if v834 != 0 {
		goto L14
	} else {
		goto L197
	}
L195:
	;
	v835 = v831
	goto L196
L196:
	;
	v836 = *(*int32)(unsafe.Add(mBase, uint32(v100)+36))
	F_RestoreTransactionSnapshot(m, v835, v836)
	mBase = m.M
	v838 = m.ExcPending
	if v838 != 0 {
		goto L14
	} else {
		goto L198
	}
L197:
	;
	v835 = v833
	goto L196
L198:
	;
	F_PushActiveSnapshot(m, v831)
	mBase = m.M
	v840 = m.ExcPending
	if v840 != 0 {
		goto L14
	} else {
		goto L199
	}
L199:
	;
	F_InvalidateSystemCachesExtended(m)
	mBase = m.M
	v842 = m.ExcPending
	if v842 != 0 {
		goto L14
	} else {
		goto L200
	}
L200:
	;
	v845 = F_shm_toc_lookup(m, v94, int64(-65532), int32(0))
	mBase = m.M
	v846 = m.ExcPending
	if v846 != 0 {
		goto L14
	} else {
		goto L201
	}
L201:
	;
	v847 = m.G0
	v849 = v847 - int32(32)
	m.G0 = v849
	v852 = *(*int32)(unsafe.Add(mBase, _c_F_ParallelWorkerMain[50]))
	v853 = int32(0)
	if base.B2i32(v852 == v853)|base.B2i32(v852 == int32(_a_F_ParallelWorkerMain_27)) == v853 {
		goto L202
	} else {
		goto L203
	}
L202:
	;
	v860 = v852
	goto L205
L203:
	;
	goto L204
L204:
	;
	v1008 = *(*int32)(unsafe.Add(mBase, uint32(v845)))
	*(*int32)(unsafe.Add(mBase, uint32(v849)+20)) = int32(1643)
	v1011 = int32(_a_F_ParallelWorkerMain_28)
	v1012 = *(*int32)(unsafe.Add(mBase, _c_F_ParallelWorkerMain[51]))
	*(*int32)(unsafe.Add(mBase, _c_F_ParallelWorkerMain[51])) = v849 + int32(16)
	*(*int32)(unsafe.Add(mBase, uint32(v849)+24)) = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v849)+16)) = v1012
	v1021 = v845 + int32(4)
	*(*int32)(unsafe.Add(mBase, uint32(v849)+28)) = v1021
	if v1008 != 0 {
		goto L273
	} else {
		goto L274
	}
L205:
	;
	v876 = *(*int32)(unsafe.Add(mBase, uint32(v860)+4))
	v879 = *(*int32)(unsafe.Add(mBase, uint32(v860-int32(60))))
	if base.Ui32(v879) < base.Ui32(int32(2)) {
		goto L207
	} else {
		goto L208
	}
L206:
	;
	goto L204
L207:
	;
	if v876 != int32(_a_F_ParallelWorkerMain_27) {
		v860 = v876
		goto L205
	} else {
		goto L270
	}
L208:
	;
	v883 = v860 - int32(32)
	v884 = *(*int32)(unsafe.Add(mBase, uint32(v883)))
	if v884 == int32(0) {
		goto L207
	} else {
		goto L209
	}
L209:
	;
	v888 = v860 - int32(4)
	v889 = *(*int32)(unsafe.Add(mBase, uint32(v888)))
	if v889 != 0 {
		goto L210
	} else {
		goto L211
	}
L210:
	;
	F_pfree(m, v889)
	mBase = m.M
	v891 = m.ExcPending
	if v891 != 0 {
		goto L14
	} else {
		goto L213
	}
L211:
	;
	goto L212
L212:
	;
	v892 = *(*int32)(unsafe.Add(mBase, uint32(v860)+16))
	if v892 != 0 {
		goto L214
	} else {
		goto L215
	}
L213:
	;
	goto L212
L214:
	;
	F_pfree(m, v892)
	mBase = m.M
	v894 = m.ExcPending
	if v894 != 0 {
		goto L14
	} else {
		goto L217
	}
L215:
	;
	goto L216
L216:
	;
	v895 = *(*int32)(unsafe.Add(mBase, uint32(v860)+20))
	if v895 != 0 {
		goto L218
	} else {
		goto L219
	}
L217:
	;
	goto L216
L218:
	;
	F_pfree(m, v895)
	mBase = m.M
	v897 = m.ExcPending
	if v897 != 0 {
		goto L14
	} else {
		goto L221
	}
L219:
	;
	goto L220
L220:
	;
	v900 = *(*int32)(unsafe.Add(mBase, uint32(v860-int32(40))))
	switch v900 {
	case 0:
		goto L228
	case 1:
		goto L227
	case 2:
		goto L226
	case 3:
		goto L225
	case 4:
		goto L224
	default:
		goto L222
	}
L221:
	;
	goto L220
L222:
	;
	v942 = *(*int32)(unsafe.Add(mBase, uint32(v883)))
	if v942 != 0 {
		goto L248
	} else {
		goto L249
	}
L223:
	;
	F_pfree(m, v938)
	mBase = m.M
	v940 = m.ExcPending
	if v940 != 0 {
		goto L14
	} else {
		goto L247
	}
L224:
	;
	v933 = *(*int32)(unsafe.Add(mBase, uint32(v860)+56))
	if v933 == int32(0) {
		goto L222
	} else {
		goto L245
	}
L225:
	;
	v916 = *(*int32)(unsafe.Add(mBase, uint32(v860)+28))
	v917 = *(*int32)(unsafe.Add(mBase, uint32(v916)))
	if v917 != 0 {
		goto L235
	} else {
		goto L236
	}
L226:
	;
	v911 = *(*int32)(unsafe.Add(mBase, uint32(v860)+80))
	if v911 == int32(0) {
		goto L222
	} else {
		goto L233
	}
L227:
	;
	v906 = *(*int32)(unsafe.Add(mBase, uint32(v860)+60))
	if v906 == int32(0) {
		goto L222
	} else {
		goto L231
	}
L228:
	;
	v901 = *(*int32)(unsafe.Add(mBase, uint32(v860)+52))
	if v901 == int32(0) {
		goto L222
	} else {
		goto L229
	}
L229:
	;
	v904 = *(*int32)(unsafe.Add(mBase, uint32(v888)))
	if v901 != v904 {
		v938 = v901
		goto L223
	} else {
		goto L230
	}
L230:
	;
	goto L222
L231:
	;
	v909 = *(*int32)(unsafe.Add(mBase, uint32(v888)))
	if v906 != v909 {
		v938 = v906
		goto L223
	} else {
		goto L232
	}
L232:
	;
	goto L222
L233:
	;
	v914 = *(*int32)(unsafe.Add(mBase, uint32(v888)))
	if v911 != v914 {
		v938 = v911
		goto L223
	} else {
		goto L234
	}
L234:
	;
	goto L222
L235:
	;
	F_pfree(m, v917)
	mBase = m.M
	v919 = m.ExcPending
	if v919 != 0 {
		goto L14
	} else {
		goto L238
	}
L236:
	;
	goto L237
L237:
	;
	v920 = *(*int32)(unsafe.Add(mBase, uint32(v860)+48))
	if v920 == int32(0) {
		goto L239
	} else {
		goto L240
	}
L238:
	;
	goto L237
L239:
	;
	v928 = *(*int32)(unsafe.Add(mBase, uint32(v860)+52))
	if v928 == int32(0) {
		goto L222
	} else {
		goto L243
	}
L240:
	;
	v923 = *(*int32)(unsafe.Add(mBase, uint32(v860)+28))
	v924 = *(*int32)(unsafe.Add(mBase, uint32(v923)))
	if v920 == v924 {
		goto L239
	} else {
		goto L241
	}
L241:
	;
	F_pfree(m, v920)
	mBase = m.M
	v927 = m.ExcPending
	if v927 != 0 {
		goto L14
	} else {
		goto L242
	}
L242:
	;
	goto L239
L243:
	;
	v931 = *(*int32)(unsafe.Add(mBase, uint32(v888)))
	if v928 != v931 {
		v938 = v928
		goto L223
	} else {
		goto L244
	}
L244:
	;
	goto L222
L245:
	;
	v936 = *(*int32)(unsafe.Add(mBase, uint32(v888)))
	if v933 == v936 {
		goto L222
	} else {
		goto L246
	}
L246:
	;
	v938 = v933
	goto L223
L247:
	;
	goto L222
L248:
	;
	v943 = *(*int32)(unsafe.Add(mBase, uint32(v860)))
	v944 = *(*int32)(unsafe.Add(mBase, uint32(v860)+4))
	*(*int32)(unsafe.Add(mBase, uint32(v943)+4)) = v944
	v946 = *(*int32)(unsafe.Add(mBase, uint32(v860)))
	*(*int32)(unsafe.Add(mBase, uint32(v944))) = v946
	goto L250
L249:
	;
	goto L250
L250:
	;
	v951 = *(*int32)(unsafe.Add(mBase, uint32(v860-int32(8))))
	if v951 != 0 {
		goto L251
	} else {
		goto L252
	}
L251:
	;
	v956 = int32(_a_F_ParallelWorkerMain_29)
	goto L256
L252:
	;
	goto L253
L253:
	;
	v967 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v860-int32(36)))))
	if v967&int32(4) != 0 {
		goto L260
	} else {
		goto L261
	}
L254:
	;
	goto L253
L255:
	;
	goto L254
L256:
	;
	v959 = *(*int32)(unsafe.Add(mBase, uint32(v956)))
	if v959 == int32(0) {
		goto L255
	} else {
		goto L258
	}
L257:
	;
	v963 = *(*int32)(unsafe.Add(mBase, uint32(v959)))
	*(*int32)(unsafe.Add(mBase, uint32(v956))) = v963
	goto L255
L258:
	;
	if v959 != v860+int32(8) {
		v956 = v959
		goto L256
	} else {
		goto L259
	}
L259:
	;
	goto L257
L260:
	;
	v974 = int32(_a_F_ParallelWorkerMain_30)
	goto L265
L261:
	;
	goto L262
L262:
	;
	F_InitializeOneGUCOption(m, v860+int32(-64))
	mBase = m.M
	v986 = m.ExcPending
	if v986 != 0 {
		goto L14
	} else {
		goto L269
	}
L263:
	;
	goto L262
L264:
	;
	goto L263
L265:
	;
	v977 = *(*int32)(unsafe.Add(mBase, uint32(v974)))
	if v977 == int32(0) {
		goto L264
	} else {
		goto L267
	}
L266:
	;
	v981 = *(*int32)(unsafe.Add(mBase, uint32(v977)))
	*(*int32)(unsafe.Add(mBase, uint32(v974))) = v981
	goto L264
L267:
	;
	if v977 != v860+int32(12) {
		v974 = v977
		goto L265
	} else {
		goto L268
	}
L268:
	;
	goto L266
L269:
	;
	goto L207
L270:
	;
	goto L206
L271:
	;
	v1152 = *(*int32)(unsafe.Add(mBase, uint32(v100)+16))
	v1153 = *(*int32)(unsafe.Add(mBase, uint32(v100)+28))
	*(*int32)(unsafe.Add(mBase, _c_F_ParallelWorkerMain[52])) = v1153
	*(*int32)(unsafe.Add(mBase, _c_F_ParallelWorkerMain[53])) = v1152
	goto L308
L272:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v1139 = m.ExcPending
	if v1139 != 0 {
		goto L14
	} else {
		goto L304
	}
L273:
	;
	v1024 = v1008 + v1021
	goto L276
L274:
	;
	v1131 = v1012
	goto L275
L275:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_ParallelWorkerMain[51])) = v1131
	m.G0 = v849 + int32(32)
	goto L271
L276:
	;
	v1042 = v849 + int32(28)
	v1043 = F_read_gucstate(m, v1042, v1024)
	mBase = m.M
	v1044 = m.ExcPending
	if v1044 != 0 {
		goto L14
	} else {
		goto L278
	}
L277:
	;
	v1114 = *(*int32)(unsafe.Add(mBase, uint32(v849)+16))
	v1131 = v1114
	goto L275
L278:
	;
	v1045 = F_read_gucstate(m, v1042, v1024)
	mBase = m.M
	v1046 = m.ExcPending
	if v1046 != 0 {
		goto L14
	} else {
		goto L279
	}
L279:
	;
	v1047 = F_read_gucstate(m, v1042, v1024)
	mBase = m.M
	v1048 = m.ExcPending
	if v1048 != 0 {
		goto L14
	} else {
		goto L280
	}
L280:
	;
	v1049 = *(*int32)(unsafe.Add(mBase, uint32(v849)+28))
	v1050 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1047))))
	if v1050 == int32(0) {
		goto L282
	} else {
		goto L283
	}
L281:
	;
	v1061 = v1058 + int32(4)
	if base.Ui32(v1024) < base.Ui32(v1061) {
		goto L17
	} else {
		goto L286
	}
L282:
	;
	v1058 = v1049
	v1059 = int32(0)
	goto L281
L283:
	;
	goto L284
L284:
	;
	v1055 = v1049 + int32(4)
	if base.Ui32(v1024) < base.Ui32(v1055) {
		goto L17
	} else {
		goto L285
	}
L285:
	;
	v1057 = *(*int32)(unsafe.Add(mBase, uint32(v1049)))
	v1058 = v1055
	v1059 = v1057
	goto L281
L286:
	;
	v1064 = v1058 + int32(8)
	if base.Ui32(v1024) < base.Ui32(v1064) {
		goto L17
	} else {
		goto L287
	}
L287:
	;
	v1067 = v1058 + int32(12)
	if base.Ui32(v1024) < base.Ui32(v1067) {
		goto L17
	} else {
		goto L288
	}
L288:
	;
	v1069 = *(*int32)(unsafe.Add(mBase, uint32(v1058)))
	v1070 = *(*int32)(unsafe.Add(mBase, uint32(v1064)))
	v1071 = *(*int32)(unsafe.Add(mBase, uint32(v1061)))
	*(*int32)(unsafe.Add(mBase, uint32(v849)+12)) = v1045
	*(*int32)(unsafe.Add(mBase, uint32(v849)+8)) = v1043
	*(*int32)(unsafe.Add(mBase, uint32(v849)+28)) = v1067
	*(*int32)(unsafe.Add(mBase, uint32(v849)+24)) = v849 + int32(8)
	v1078 = int32(0)
	v1080 = int32(1)
	v1083 = F_set_config_with_handle(m, v1043, v1078, v1045, v1071, v1069, v1070, v1078, v1080, int32(21), v1080)
	mBase = m.M
	v1084 = m.ExcPending
	if v1084 != 0 {
		goto L14
	} else {
		goto L289
	}
L289:
	;
	if v1083 <= int32(0) {
		goto L272
	} else {
		goto L290
	}
L290:
	;
	v1087 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1047))))
	if v1087 == int32(0) {
		goto L291
	} else {
		goto L292
	}
L291:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v849)+24)) = int32(0)
	if base.Ui32(v1067) < base.Ui32(v1024) {
		goto L276
	} else {
		goto L303
	}
L292:
	;
	v1095 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_ParallelWorkerMain[54])))
	if v1095 != 0 {
		goto L293
	} else {
		goto L294
	}
L293:
	;
	v1096 = int32(12)
	goto L295
L294:
	;
	v1096 = int32(15)
	goto L295
L295:
	;
	v1097 = F_find_option(m, v1043, int32(1), int32(0), v1096)
	mBase = m.M
	v1098 = m.ExcPending
	if v1098 != 0 {
		goto L14
	} else {
		goto L296
	}
L296:
	;
	if v1097 == int32(0) {
		goto L291
	} else {
		goto L297
	}
L297:
	;
	v1101 = F_guc_strdup(m, v1096, v1047)
	mBase = m.M
	v1102 = m.ExcPending
	if v1102 != 0 {
		goto L14
	} else {
		goto L298
	}
L298:
	;
	v1103 = *(*int32)(unsafe.Add(mBase, uint32(v1097)+84))
	if v1103 != 0 {
		goto L299
	} else {
		goto L300
	}
L299:
	;
	F_pfree(m, v1103)
	mBase = m.M
	v1105 = m.ExcPending
	if v1105 != 0 {
		goto L14
	} else {
		goto L302
	}
L300:
	;
	goto L301
L301:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v1097)+88)) = v1059
	*(*int32)(unsafe.Add(mBase, uint32(v1097)+84)) = v1101
	goto L291
L302:
	;
	goto L301
L303:
	;
	goto L277
L304:
	;
	F_errcode(m, int32(2600))
	mBase = m.M
	v1142 = m.ExcPending
	if v1142 != 0 {
		goto L14
	} else {
		goto L305
	}
L305:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v849))) = v1043
	F_errmsg(m, int32(_a_F_ParallelWorkerMain_31), v849)
	mBase = m.M
	v1146 = m.ExcPending
	if v1146 != 0 {
		goto L14
	} else {
		goto L306
	}
L306:
	;
	F_errfinish(m, int32(_a_F_ParallelWorkerMain_32), int32(_a_F_ParallelWorkerMain_33), int32(_a_F_ParallelWorkerMain_34))
	mBase = m.M
	v1151 = m.ExcPending
	if v1151 != 0 {
		goto L14
	} else {
		goto L307
	}
L307:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L308:
	;
	v1158 = *(*int32)(unsafe.Add(mBase, uint32(v100)+20))
	v1160 = *(*int32)(unsafe.Add(mBase, uint32(v100)+24))
	*(*int32)(unsafe.Add(mBase, _c_F_ParallelWorkerMain[55])) = v1160
	*(*int32)(unsafe.Add(mBase, _c_F_ParallelWorkerMain[56])) = v1158
	v1165 = int32(1)
	*(*uint8)(unsafe.Add(mBase, _c_F_ParallelWorkerMain[57])) = uint8(v1165)
	v1168 = int32(0)
	*(*uint8)(unsafe.Add(mBase, _c_F_ParallelWorkerMain[58])) = uint8(v1168)
	v1172 = F_shm_toc_lookup(m, v94, int64(-65522), v1168)
	mBase = m.M
	v1173 = m.ExcPending
	if v1173 != 0 {
		goto L14
	} else {
		goto L309
	}
L309:
	;
	v1174 = m.G0
	v1176 = v1174 - int32(48)
	m.G0 = v1176
	v1178 = *(*int32)(unsafe.Add(mBase, uint32(v1172)))
	if v1178 != 0 {
		goto L310
	} else {
		goto L311
	}
L310:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v1176)+16)) = int64(17179869188)
	v1182 = *(*int32)(unsafe.Add(mBase, _c_F_ParallelWorkerMain[40]))
	*(*int32)(unsafe.Add(mBase, uint32(v1176)+40)) = v1182
	v1188 = F_hash_create(m, int32(_a_F_ParallelWorkerMain_35), int32(32), v1176, int32(1064))
	mBase = m.M
	v1189 = m.ExcPending
	if v1189 != 0 {
		goto L14
	} else {
		goto L313
	}
L311:
	;
	v1216 = v1172
	goto L312
L312:
	;
	v1232 = *(*int32)(unsafe.Add(mBase, uint32(v1216)+4))
	if v1232 != 0 {
		goto L318
	} else {
		goto L319
	}
L313:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_ParallelWorkerMain[59])) = v1188
	v1191 = v1172
	goto L314
L314:
	;
	v1208 = *(*int32)(unsafe.Add(mBase, _c_F_ParallelWorkerMain[59]))
	v1211 = F_hash_search(m, v1208, v1191, int32(1), int32(0))
	mBase = m.M
	v1212 = m.ExcPending
	if v1212 != 0 {
		goto L14
	} else {
		goto L316
	}
L315:
	;
	v1216 = v1214
	goto L312
L316:
	;
	v1214 = v1191 + int32(4)
	v1215 = *(*int32)(unsafe.Add(mBase, uint32(v1191)+4))
	if v1215 != 0 {
		v1191 = v1214
		goto L314
	} else {
		goto L317
	}
L317:
	;
	goto L315
L318:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v1176)+16)) = int64(17179869188)
	v1238 = *(*int32)(unsafe.Add(mBase, _c_F_ParallelWorkerMain[40]))
	*(*int32)(unsafe.Add(mBase, uint32(v1176)+40)) = v1238
	v1244 = F_hash_create(m, int32(_a_F_ParallelWorkerMain_36), int32(32), v1176, int32(1064))
	mBase = m.M
	v1245 = m.ExcPending
	if v1245 != 0 {
		goto L14
	} else {
		goto L321
	}
L319:
	;
	goto L320
L320:
	;
	m.G0 = v1176 + int32(48)
	v1293 = F_shm_toc_lookup(m, v94, int64(-65521), int32(0))
	mBase = m.M
	v1294 = m.ExcPending
	if v1294 != 0 {
		goto L14
	} else {
		goto L326
	}
L321:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_ParallelWorkerMain[60])) = v1244
	v1248 = v1216 + int32(4)
	goto L322
L322:
	;
	v1264 = *(*int32)(unsafe.Add(mBase, _c_F_ParallelWorkerMain[60]))
	v1267 = F_hash_search(m, v1264, v1248, int32(1), int32(0))
	mBase = m.M
	v1268 = m.ExcPending
	if v1268 != 0 {
		goto L14
	} else {
		goto L324
	}
L323:
	;
	goto L320
L324:
	;
	v1269 = *(*int32)(unsafe.Add(mBase, uint32(v1248)+4))
	if v1269 != 0 {
		v1248 = v1248 + int32(4)
		goto L322
	} else {
		goto L325
	}
L325:
	;
	goto L323
L326:
	;
	v1295 = *(*int32)(unsafe.Add(mBase, uint32(v1293)))
	v1297 = *(*int32)(unsafe.Add(mBase, uint32(v1293)+4))
	*(*int32)(unsafe.Add(mBase, _c_F_ParallelWorkerMain[61])) = v1297
	v1300 = int32(0)
	*(*int32)(unsafe.Add(mBase, _c_F_ParallelWorkerMain[62])) = v1300
	if v1300 <= v1295 {
		goto L327
	} else {
		goto L328
	}
L327:
	;
	v1306 = *(*int32)(unsafe.Add(mBase, _c_F_ParallelWorkerMain[7]))
	v1309 = F_MemoryContextStrdup(m, v1306, v1293+int32(8))
	mBase = m.M
	v1310 = m.ExcPending
	if v1310 != 0 {
		goto L14
	} else {
		goto L330
	}
L328:
	;
	goto L329
L329:
	;
	v1313 = *(*int32)(unsafe.Add(mBase, _c_F_ParallelWorkerMain[62]))
	if v1313 != 0 {
		goto L331
	} else {
		goto L332
	}
L330:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_ParallelWorkerMain[62])) = v1309
	goto L329
L331:
	;
	v1315 = *(*int32)(unsafe.Add(mBase, _c_F_ParallelWorkerMain[61]))
	v1318 = *(*int32)(unsafe.Add(mBase, uint32(v1315<<(uint(int32(2))%32))+uint32(_c_F_ParallelWorkerMain[63])))
	goto L334
L332:
	;
	goto L333
L333:
	;
	v1321 = *(*int32)(unsafe.Add(mBase, uint32(v100)+64))
	v1322 = m.G0
	v1324 = v1322 - int32(48)
	m.G0 = v1324
	*(*int32)(unsafe.Add(mBase, _c_F_ParallelWorkerMain[64])) = v1321
	if v1321 != 0 {
		goto L336
	} else {
		goto L337
	}
L334:
	;
	F_InitializeSystemUser(m, v1313, v1318)
	mBase = m.M
	v1320 = m.ExcPending
	if v1320 != 0 {
		goto L14
	} else {
		goto L335
	}
L335:
	;
	goto L333
L336:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v1324)+16)) = int64(103079215120)
	v1333 = *(*int32)(unsafe.Add(mBase, _c_F_ParallelWorkerMain[65]))
	v1335 = F_hash_create(m, int32(_a_F_ParallelWorkerMain_37), v1333, v1324, int32(40))
	mBase = m.M
	v1336 = m.ExcPending
	if v1336 != 0 {
		goto L14
	} else {
		goto L339
	}
L337:
	;
	goto L338
L338:
	;
	m.G0 = v1324 + int32(48)
	v1342 = int32(0)
	*(*uint8)(unsafe.Add(mBase, _c_F_ParallelWorkerMain[0])) = uint8(v1342)
	v1346 = *(*int32)(unsafe.Add(mBase, _c_F_ParallelWorkerMain[35]))
	v1347 = *(*int32)(unsafe.Add(mBase, uint32(v1346)+72))
	*(*int32)(unsafe.Add(mBase, uint32(v1346)+72)) = v1347 + int32(1)
	goto L340
L339:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_ParallelWorkerMain[66])) = v1335
	goto L338
L340:
	;
	m.T0[v382].(func(*base.Module, int32, int32))(m, v87, v94)
	mBase = m.M
	v1352 = m.ExcPending
	if v1352 != 0 {
		goto L14
	} else {
		goto L341
	}
L341:
	;
	v1355 = *(*int32)(unsafe.Add(mBase, _c_F_ParallelWorkerMain[35]))
	v1356 = *(*int32)(unsafe.Add(mBase, uint32(v1355)+72))
	*(*int32)(unsafe.Add(mBase, uint32(v1355)+72)) = v1356 - int32(1)
	goto L342
L342:
	;
	F_PopActiveSnapshot(m)
	mBase = m.M
	v1361 = m.ExcPending
	if v1361 != 0 {
		goto L14
	} else {
		goto L343
	}
L343:
	;
	F_CommitTransaction(m)
	mBase = m.M
	v1363 = m.ExcPending
	if v1363 != 0 {
		goto L14
	} else {
		goto L344
	}
L344:
	;
	v1365 = *(*int32)(unsafe.Add(mBase, _c_F_ParallelWorkerMain[35]))
	*(*int32)(unsafe.Add(mBase, uint32(v1365)+24)) = int32(0)
	v1369 = *(*int32)(unsafe.Add(mBase, _c_F_ParallelWorkerMain[49]))
	v1370 = *(*int32)(unsafe.Add(mBase, uint32(v1369)))
	F_dsm_detach(m, v1370)
	mBase = m.M
	v1372 = m.ExcPending
	if v1372 != 0 {
		goto L14
	} else {
		goto L345
	}
L345:
	;
	v1373 = int32(_a_F_ParallelWorkerMain_24)
	v1374 = *(*int32)(unsafe.Add(mBase, _c_F_ParallelWorkerMain[49]))
	*(*int32)(unsafe.Add(mBase, uint32(v1374))) = int32(0)
	v1378 = *(*int32)(unsafe.Add(mBase, _c_F_ParallelWorkerMain[49]))
	v1379 = *(*int32)(unsafe.Add(mBase, uint32(v1378)+4))
	F_dsa_detach(m, v1379)
	mBase = m.M
	v1381 = m.ExcPending
	if v1381 != 0 {
		goto L14
	} else {
		goto L346
	}
L346:
	;
	v1383 = *(*int32)(unsafe.Add(mBase, _c_F_ParallelWorkerMain[49]))
	v1384 = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v1383)+4)) = v1384
	v1390 = *(*int32)(unsafe.Add(mBase, _c_F_ParallelWorkerMain[67]))
	v1391 = *(*int32)(unsafe.Add(mBase, uint32(v1390)+16))
	v1392 = m.T0[v1391].(func(*base.Module, int32, int32, int32) int32)(m, int32(88), v1384, v1384)
	mBase = m.M
	v1393 = m.ExcPending
	if v1393 != 0 {
		goto L14
	} else {
		goto L347
	}
L347:
	;
	goto L47
L348:
	;
	F_errcode(m, int32(325))
	mBase = m.M
	v1419 = m.ExcPending
	if v1419 != 0 {
		goto L14
	} else {
		goto L349
	}
L349:
	;
	F_errmsg(m, int32(_a_F_ParallelWorkerMain_38), int32(0))
	mBase = m.M
	v1423 = m.ExcPending
	if v1423 != 0 {
		goto L14
	} else {
		goto L350
	}
L350:
	;
	F_errfinish(m, int32(_a_F_ParallelWorkerMain_39), int32(1355), int32(_a_F_ParallelWorkerMain_40))
	mBase = m.M
	v1428 = m.ExcPending
	if v1428 != 0 {
		goto L14
	} else {
		goto L351
	}
L351:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L352:
	;
	F_errcode(m, int32(325))
	mBase = m.M
	v1435 = m.ExcPending
	if v1435 != 0 {
		goto L14
	} else {
		goto L353
	}
L353:
	;
	F_errmsg(m, int32(_a_F_ParallelWorkerMain_41), int32(0))
	mBase = m.M
	v1439 = m.ExcPending
	if v1439 != 0 {
		goto L14
	} else {
		goto L354
	}
L354:
	;
	F_errfinish(m, int32(_a_F_ParallelWorkerMain_39), int32(1360), int32(_a_F_ParallelWorkerMain_40))
	mBase = m.M
	v1444 = m.ExcPending
	if v1444 != 0 {
		goto L14
	} else {
		goto L355
	}
L355:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L356:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v19)+16)) = v195
	F_errmsg_internal(m, int32(_a_F_ParallelWorkerMain_42), v19+int32(16))
	mBase = m.M
	v1454 = m.ExcPending
	if v1454 != 0 {
		goto L14
	} else {
		goto L357
	}
L357:
	;
	F_errfinish(m, int32(_a_F_ParallelWorkerMain_39), int32(1666), int32(_a_F_ParallelWorkerMain_43))
	mBase = m.M
	v1459 = m.ExcPending
	if v1459 != 0 {
		goto L14
	} else {
		goto L358
	}
L358:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L359:
	;
	v1465 = *(*int32)(unsafe.Add(mBase, _c_F_ParallelWorkerMain[31]))
	v1466 = *(*int32)(unsafe.Add(mBase, uint32(v1465)+4))
	goto L360
L360:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v19))) = v1466
	F_errmsg_internal(m, int32(_a_F_ParallelWorkerMain_44), v19)
	mBase = m.M
	v1470 = m.ExcPending
	if v1470 != 0 {
		goto L14
	} else {
		goto L361
	}
L361:
	;
	F_errfinish(m, int32(_a_F_ParallelWorkerMain_39), int32(1452), int32(_a_F_ParallelWorkerMain_40))
	mBase = m.M
	v1475 = m.ExcPending
	if v1475 != 0 {
		goto L14
	} else {
		goto L362
	}
L362:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L363:
	;
	F_errmsg_internal(m, int32(_a_F_ParallelWorkerMain_45), int32(0))
	mBase = m.M
	v1488 = m.ExcPending
	if v1488 != 0 {
		goto L14
	} else {
		goto L364
	}
L364:
	;
	F_errfinish(m, int32(_a_F_ParallelWorkerMain_32), int32(_a_F_ParallelWorkerMain_46), int32(_a_F_ParallelWorkerMain_47))
	mBase = m.M
	v1493 = m.ExcPending
	if v1493 != 0 {
		goto L14
	} else {
		goto L365
	}
L365:
	;
	base.Wasm_trap_unreachable()
	for {
	}
}
func F_compute_parallel_worker(m *base.Module, l0 int32, l1 float64, l2 float64, l3 int32) int32 {
	mBase := m.M
	_ = mBase
	var v7 int32
	_ = v7
	var v10 int32
	_ = v10
	var v15 int32
	_ = v15
	var v24 int32
	_ = v24
	var v27 int32
	_ = v27
	var v32 int32
	_ = v32
	var v34 int32
	_ = v34
	var v37 int32
	_ = v37
	var v39 int32
	_ = v39
	var v43 int32
	_ = v43
	var v46 int32
	_ = v46
	var v52 int32
	_ = v52
	var v59 int32
	_ = v59
	var v65 int32
	_ = v65
	var v67 int32
	_ = v67
	var v70 int32
	_ = v70
	var v72 int32
	_ = v72
	var v77 int32
	_ = v77
	var v79 int32
	_ = v79
	var v83 int32
	_ = v83
	var v86 int32
	_ = v86
	var v88 int32
	_ = v88
	var v91 int32
	_ = v91
	var v96 int32
	_ = v96
	var v99 int32
	_ = v99
	var v106 int32
	_ = v106
	v7 = *(*int32)(unsafe.Add(mBase, uint32(l0)+148))
	if v7 != int32(-1) {
		v96 = v7
		goto L2
	} else {
		goto L3
	}
L1:
	;
	return v106
L2:
	;
	if v96 < l3 {
		goto L37
	} else {
		goto L38
	}
L3:
	;
	v10 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	if v10 != 0 {
		goto L4
	} else {
		goto L5
	}
L4:
	;
	v27 = int32(0)
	if base.F64_ge(l1, float64(0)) == v27 {
		v59 = v27
		goto L12
	} else {
		goto L13
	}
L5:
	;
	if base.F64_ge(l1, float64(0)) != 0 {
		goto L6
	} else {
		goto L7
	}
L6:
	;
	v15 = *(*int32)(unsafe.Add(mBase, _c_F_compute_parallel_worker[0]))
	if base.F64_lt(l1, base.F64_convert_i32_s(v15)) != 0 {
		v106 = int32(0)
		goto L1
	} else {
		goto L9
	}
L7:
	;
	goto L8
L8:
	;
	if base.F64_ge(l2, float64(0)) == int32(0) {
		goto L4
	} else {
		goto L10
	}
L9:
	;
	goto L8
L10:
	;
	v24 = *(*int32)(unsafe.Add(mBase, _c_F_compute_parallel_worker[1]))
	if base.F64_lt(l2, base.F64_convert_i32_s(v24)) != 0 {
		v106 = int32(0)
		goto L1
	} else {
		goto L11
	}
L11:
	;
	goto L4
L12:
	;
	if base.F64_ge(l2, float64(0)) == int32(0) {
		v96 = v59
		goto L2
	} else {
		goto L21
	}
L13:
	;
	v32 = int32(1)
	v34 = *(*int32)(unsafe.Add(mBase, _c_F_compute_parallel_worker[0]))
	if v34 <= v32 {
		goto L14
	} else {
		goto L15
	}
L14:
	;
	v37 = v32
	goto L16
L15:
	;
	v37 = v34
	goto L16
L16:
	;
	v39 = v37
	v43 = int32(1)
	goto L17
L17:
	;
	v46 = v39 * int32(3)
	if base.F64_ge(l1, base.F64_convert_i32_u(v46)) == int32(0) {
		v59 = v43
		goto L12
	} else {
		goto L19
	}
L18:
	;
	v59 = v52
	goto L12
L19:
	;
	v52 = v43 + int32(1)
	if v46 < int32(715827883) {
		v39 = v46
		v43 = v52
		goto L17
	} else {
		goto L20
	}
L20:
	;
	goto L18
L21:
	;
	v65 = int32(1)
	v67 = *(*int32)(unsafe.Add(mBase, _c_F_compute_parallel_worker[1]))
	if v67 <= v65 {
		goto L22
	} else {
		goto L23
	}
L22:
	;
	v70 = v65
	goto L24
L23:
	;
	v70 = v67
	goto L24
L24:
	;
	v72 = v70
	v77 = int32(1)
	goto L25
L25:
	;
	v79 = v72 * int32(3)
	if base.F64_le(base.F64_convert_i32_u(v79), l2) != 0 {
		goto L27
	} else {
		goto L28
	}
L26:
	;
	if v59 < v86 {
		goto L31
	} else {
		goto L32
	}
L27:
	;
	v83 = v77 + int32(1)
	if v79 < int32(715827883) {
		v72 = v79
		v77 = v83
		goto L25
	} else {
		goto L30
	}
L28:
	;
	v86 = v77
	goto L29
L29:
	;
	goto L26
L30:
	;
	v86 = v83
	goto L29
L31:
	;
	v88 = v59
	goto L33
L32:
	;
	v88 = v86
	goto L33
L33:
	;
	if int32(0) < v59 {
		goto L34
	} else {
		goto L35
	}
L34:
	;
	v91 = v88
	goto L36
L35:
	;
	v91 = v86
	goto L36
L36:
	;
	v96 = v91
	goto L2
L37:
	;
	v99 = v96
	goto L39
L38:
	;
	v99 = l3
	goto L39
L39:
	;
	v106 = v99
	goto L1
}
func F_parallel_vacuum_get_dead_items(m *base.Module, l0 int32, l1 int32) int32 {
	mBase := m.M
	_ = mBase
	var v3 int32
	_ = v3
	var v7 int32
	_ = v7
	v3 = *(*int32)(unsafe.Add(mBase, uint32(l0)+16))
	*(*int32)(unsafe.Add(mBase, uint32(l1))) = v3 + int32(56)
	v7 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
	return v7
}
